package main

import (
	"bufio"
	"flag"
	"fmt"
	"gitlab.com/mokytis/goreader/feed"
	"os"
	"strings"
	"sync"
)

func customUsage() {
	fmt.Fprintf(os.Stderr,
		`Usage: %s [options]

Gets URLS for items from RSS feeds provided from stdin


Options:
`,
		os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `Format String Options:
  %%u - item url
  %%t - item title
  %%d - item publish date

Examples:
  %s < feeds.txt
      output all urls for each RSS feed in a file
  %s -f "%%u - %%d" < feeds.txt
      use custom formating (in this case "<URL> - <PUBLISH DATE>"
`,
		os.Args[0], os.Args[0])
}

func formatOutput(formatString string, item feed.Item) string {
	r := strings.NewReplacer(
		"%u", item.Link,
		"%t", item.Title,
		"%d", item.PubDate,
	)
	result := r.Replace(formatString)
	return result
}

func main() {
	flag.Usage = customUsage

	var workers int
	flag.IntVar(&workers, "w", 20, "the ammount of workers")

	var outputFmt string
	flag.StringVar(&outputFmt, "f", "%u", "the format string for output")

	flag.Parse()

	feed_urls := make(chan string)
	output := make(chan string)

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
						output <- formatOutput(outputFmt, item)
					}
				}
			}

			feedWG.Done()
		}()
	}

	// output workder
	var outputWG sync.WaitGroup
	outputWG.Add(1)
	go func() {
		for o := range output {
			fmt.Println(o)
		}
		outputWG.Done()
	}()

	// close the output channel when all feed workers are done
	go func() {
		feedWG.Wait()
		close(output)
	}()

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		feed_url := sc.Text()
		feed_urls <- feed_url
	}

	close(feed_urls)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}
	outputWG.Wait()

}
