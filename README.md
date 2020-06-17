# goreader

A go tool for playing with RSS feeds.

## goreader

`goreader` is an RSS reader written in go.

### Install

```bash
go get gitlab.com/mokytis/goreader
```

### Usage

```bash
$ cat feeds.txt
http://news.mit.edu/rss/feed
http://news.mit.edu/rss/research

$ goreader feeds.txt
 1) Astronomers detect regular rhythm of radio waves, with origins unknown
 2) Why the Mediterranean is a climate change hotspot
 3) 3 questions with James H. Williams, Jr., professor of mechanical engineering
 4) Study sheds light on a classic visual illusion
 5) Steady demand for PPE might encourage local businesses to start production
 6) What moves people?
 7) Study finds path for addressing Alzheimer’s blood-brain barrier impairment
 8) Startup with MIT roots develops lightweight solar panels
 9) From delayed deceleration to Zooming
10) Capturing stardust
11) Team 2020 charts a course for MIT
12) Tiny sand grains trigger massive glacial surges
13) Letter to the MIT faculty: A moment of moral urgency
14) Newly observed phenomenon could lead to new quantum devices
15) State-level R&amp;D tax credits spur growth of new businesses
16) Astronomers detect regular rhythm of radio waves, with origins unknown
17) Why the Mediterranean is a climate change hotspot
18) Study sheds light on a classic visual illusion
19) Study finds path for addressing Alzheimer’s blood-brain barrier impairment
20) From delayed deceleration to Zooming
21) Capturing stardust
22) Tiny sand grains trigger massive glacial surges
23) Newly observed phenomenon could lead to new quantum devices
24) State-level R&amp;D tax credits spur growth of new businesses
25) A layered approach to safety
Choose Article: 10
# opens article in web browser
```

## getrssitems

`getrssitems` is a commandline utility written in go for getting urls for items in RSS feeds.

### Install

```bash
go get gitlab.com/mokytis/goreader/cmd/getrssitems
```

### Usage

```bash
$ cat feeds.txt
http://news.mit.edu/rss/feed
http://news.mit.edu/rss/research
```
```bash
$ getrssitems < feeds.txt
http://news.mit.edu/2020/astronomers-rhythm-radio-waves-0617
http://news.mit.edu/2020/why-mediterranean-climate-change-hotspot-0617
http://news.mit.edu/2020/study-visual-illusion-brightness-0617
http://news.mit.edu/2020/study-finds-path-addressing-alzheimers-blood-brain-barrier-impairment-0615
http://news.mit.edu/2020/delayed-deceleration-zooming-jacqueline-thomas-mit-0614
http://news.mit.edu/2020/capturing-stardust-danielle-frostig-mit-kavli-0612
http://news.mit.edu/2020/sand-grains-massive-glacial-surges-0612
http://news.mit.edu/2020/kohn-anomaly-quantum-devices-0612
http://news.mit.edu/2020/state-rd-tax-credits-growth-new-businesses-0612
http://news.mit.edu/2020/layered-approach-safety-koroush-shirvan-0611
http://news.mit.edu/2020/astronomers-rhythm-radio-waves-0617
http://news.mit.edu/2020/why-mediterranean-climate-change-hotspot-0617
http://news.mit.edu/2020/3-questions-james-h-williams-jr-professor-mechanical-engineering
http://news.mit.edu/2020/study-visual-illusion-brightness-0617
http://news.mit.edu/2020/steady-demand-ppe-may-encourage-local-businesses-start-production-0616
http://news.mit.edu/2020/professor-jinhua-zhao-0616
http://news.mit.edu/2020/study-finds-path-addressing-alzheimers-blood-brain-barrier-impairment-0615
http://news.mit.edu/2020/swift-solar-startup-mit-roots-develops-lightweight-solar-panels-0615
http://news.mit.edu/2020/delayed-deceleration-zooming-jacqueline-thomas-mit-0614
http://news.mit.edu/2020/capturing-stardust-danielle-frostig-mit-kavli-0612
http://news.mit.edu/2020/team-2020-charts-course-mit-covid-19-0612
http://news.mit.edu/2020/sand-grains-massive-glacial-surges-0612
http://news.mit.edu/2020/letter-faculty-moral-urgency-inclusion-0612
http://news.mit.edu/2020/kohn-anomaly-quantum-devices-0612
http://news.mit.edu/2020/state-rd-tax-credits-growth-new-businesses-0612
```

```bash
$  getrssitems -f "%u %d" < feeds.txt
http://news.mit.edu/2020/astronomers-rhythm-radio-waves-0617 Wed, 17 Jun 2020 11:00:00 -0400
http://news.mit.edu/2020/why-mediterranean-climate-change-hotspot-0617 Wed, 17 Jun 2020 09:55:48 -0400
http://news.mit.edu/2020/study-visual-illusion-brightness-0617 Tue, 16 Jun 2020 23:59:59 -0400
http://news.mit.edu/2020/study-finds-path-addressing-alzheimers-blood-brain-barrier-impairment-0615 Mon, 15 Jun 2020 14:35:01 -0400
http://news.mit.edu/2020/delayed-deceleration-zooming-jacqueline-thomas-mit-0614 Sun, 14 Jun 2020 00:00:00 -0400
http://news.mit.edu/2020/capturing-stardust-danielle-frostig-mit-kavli-0612 Fri, 12 Jun 2020 15:40:01 -0400
http://news.mit.edu/2020/sand-grains-massive-glacial-surges-0612 Fri, 12 Jun 2020 15:17:33 -0400
http://news.mit.edu/2020/kohn-anomaly-quantum-devices-0612 Fri, 12 Jun 2020 11:20:32 -0400
http://news.mit.edu/2020/state-rd-tax-credits-growth-new-businesses-0612 Fri, 12 Jun 2020 00:00:00 -0400
http://news.mit.edu/2020/layered-approach-safety-koroush-shirvan-0611 Thu, 11 Jun 2020 15:50:01 -0400
http://news.mit.edu/2020/astronomers-rhythm-radio-waves-0617 Wed, 17 Jun 2020 11:00:00 -0400
http://news.mit.edu/2020/why-mediterranean-climate-change-hotspot-0617 Wed, 17 Jun 2020 09:55:48 -0400
http://news.mit.edu/2020/3-questions-james-h-williams-jr-professor-mechanical-engineering Wed, 17 Jun 2020 08:00:00 -0400
http://news.mit.edu/2020/study-visual-illusion-brightness-0617 Tue, 16 Jun 2020 23:59:59 -0400
http://news.mit.edu/2020/steady-demand-ppe-may-encourage-local-businesses-start-production-0616 Tue, 16 Jun 2020 16:00:01 -0400
http://news.mit.edu/2020/professor-jinhua-zhao-0616 Mon, 15 Jun 2020 23:59:59 -0400
http://news.mit.edu/2020/study-finds-path-addressing-alzheimers-blood-brain-barrier-impairment-0615 Mon, 15 Jun 2020 14:35:01 -0400
http://news.mit.edu/2020/swift-solar-startup-mit-roots-develops-lightweight-solar-panels-0615 Mon, 15 Jun 2020 14:10:01 -0400
http://news.mit.edu/2020/delayed-deceleration-zooming-jacqueline-thomas-mit-0614 Sun, 14 Jun 2020 00:00:00 -0400
http://news.mit.edu/2020/capturing-stardust-danielle-frostig-mit-kavli-0612 Fri, 12 Jun 2020 15:40:01 -0400
http://news.mit.edu/2020/team-2020-charts-course-mit-covid-19-0612 Fri, 12 Jun 2020 15:20:01 -0400
http://news.mit.edu/2020/sand-grains-massive-glacial-surges-0612 Fri, 12 Jun 2020 15:17:33 -0400
http://news.mit.edu/2020/letter-faculty-moral-urgency-inclusion-0612 Fri, 12 Jun 2020 14:55:25 -0400
http://news.mit.edu/2020/kohn-anomaly-quantum-devices-0612 Fri, 12 Jun 2020 11:20:32 -0400
http://news.mit.edu/2020/state-rd-tax-credits-growth-new-businesses-0612 Fri, 12 Jun 2020 00:00:00 -0400
```

```bash
$ getrssitems -f "%t %u" < feeds.txt
Astronomers detect regular rhythm of radio waves, with origins unknown  http://news.mit.edu/2020/astronomers-rhythm-radio-waves-0617
Why the Mediterranean is a climate change hotspot http://news.mit.edu/2020/why-mediterranean-climate-change-hotspot-0617
Study sheds light on a classic visual illusion http://news.mit.edu/2020/study-visual-illusion-brightness-0617
Study finds path for addressing Alzheimer’s blood-brain barrier impairment http://news.mit.edu/2020/study-finds-path-addressing-alzheimers-blood-brain-barrier-impairment-0615
From delayed deceleration to Zooming http://news.mit.edu/2020/delayed-deceleration-zooming-jacqueline-thomas-mit-0614
Capturing stardust http://news.mit.edu/2020/capturing-stardust-danielle-frostig-mit-kavli-0612
Tiny sand grains trigger massive glacial surges http://news.mit.edu/2020/sand-grains-massive-glacial-surges-0612
Newly observed phenomenon could lead to new quantum devices http://news.mit.edu/2020/kohn-anomaly-quantum-devices-0612
State-level R&amp;D tax credits spur growth of new businesses http://news.mit.edu/2020/state-rd-tax-credits-growth-new-businesses-0612
A layered approach to safety http://news.mit.edu/2020/layered-approach-safety-koroush-shirvan-0611
Astronomers detect regular rhythm of radio waves, with origins unknown  http://news.mit.edu/2020/astronomers-rhythm-radio-waves-0617
Why the Mediterranean is a climate change hotspot http://news.mit.edu/2020/why-mediterranean-climate-change-hotspot-0617
3 questions with James H. Williams, Jr., professor of mechanical engineering  http://news.mit.edu/2020/3-questions-james-h-williams-jr-professor-mechanical-engineering
Study sheds light on a classic visual illusion http://news.mit.edu/2020/study-visual-illusion-brightness-0617
Steady demand for PPE might encourage local businesses to start production http://news.mit.edu/2020/steady-demand-ppe-may-encourage-local-businesses-start-production-0616
What moves people? http://news.mit.edu/2020/professor-jinhua-zhao-0616
Study finds path for addressing Alzheimer’s blood-brain barrier impairment http://news.mit.edu/2020/study-finds-path-addressing-alzheimers-blood-brain-barrier-impairment-0615
Startup with MIT roots develops lightweight solar panels http://news.mit.edu/2020/swift-solar-startup-mit-roots-develops-lightweight-solar-panels-0615
From delayed deceleration to Zooming http://news.mit.edu/2020/delayed-deceleration-zooming-jacqueline-thomas-mit-0614
Capturing stardust http://news.mit.edu/2020/capturing-stardust-danielle-frostig-mit-kavli-0612
Team 2020 charts a course for MIT http://news.mit.edu/2020/team-2020-charts-course-mit-covid-19-0612
Tiny sand grains trigger massive glacial surges http://news.mit.edu/2020/sand-grains-massive-glacial-surges-0612
Letter to the MIT faculty: A moment of moral urgency http://news.mit.edu/2020/letter-faculty-moral-urgency-inclusion-0612
Newly observed phenomenon could lead to new quantum devices http://news.mit.edu/2020/kohn-anomaly-quantum-devices-0612
State-level R&amp;D tax credits spur growth of new businesses http://news.mit.edu/2020/state-rd-tax-credits-growth-new-businesses-0612
```



