package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pkg/browser"
	"gitlab.com/mokytis/goreader/feed"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func customUsage() {
	fmt.Fprintf(os.Stderr,
		`Usage: %s FILE [options]

Select and open an item from RSS feeds in FILE

Options:
`,
		os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = customUsage

	var workers int
	flag.IntVar(&workers, "w", 20, "the ammount of workers")

	flag.Parse()

	feedsFile := flag.Arg(0)
	if _, err := os.Stat(feedsFile); os.IsNotExist(err) {
		fmt.Println("File " + feedsFile + " does not exists.")
		os.Exit(1)
	}

	feed_urls := make(chan string)
	output := make(chan feed.Item)

	// feed workers
	var feedWG sync.WaitGroup
	for i := 0; i < workers; i++ {
		feedWG.Add(1)

		go func() {
			for feed_url := range feed_urls {
				xml, err := feed.Fetch(feed_url)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to get feed: %s\n", err)
				}
				rss := feed.Parse(xml)
				for _, channel := range rss.Channels {
					for _, item := range channel.Items {
						output <- item
					}
				}
			}

			feedWG.Done()
		}()
	}

	// output workder
	items := []feed.Item{}
	var outputWG sync.WaitGroup
	outputWG.Add(1)
	go func() {
		i := 1
		for o := range output {
			fmt.Printf("%2d) %s\n", i, o.Title)
			items = append(items, o)
			i += 1
		}
		outputWG.Done()
	}()

	// close the output channel when all feed workers are done
	go func() {
		feedWG.Wait()
		close(output)
	}()

	file, err := os.Open(feedsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		feed_url := sc.Text()
		feed_urls <- feed_url
	}

	close(feed_urls)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %s\n", err)
	}
	outputWG.Wait()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose Article: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("failed to read stdin", err)
	}
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		fmt.Println("Invalid item.", err)
		os.Exit(1)
	}
	if i > len(items) || i <= 0 {
		fmt.Println("index out of range")
		os.Exit(1)
	}
	browser.OpenURL(items[i-1].Link)
}
