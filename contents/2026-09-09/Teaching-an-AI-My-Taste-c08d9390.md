---
source: "https://lonriesberg.com/posts/teaching-an-ai-my-taste/"
hn_url: "https://news.ycombinator.com/item?id=49629118"
title: "Teaching an AI My Taste"
article_title: "Teaching an AI My Taste – Lon Riesberg"
image: ""
author: "speckx"
captured_at: "2026-09-09T17:03:57Z"
capture_tool: "hn-digest"
hn_id: 49629118
score: 1
comments: 0
posted_at: "2026-09-09T16:27:49Z"
tags:
  - hacker-news
---

# Teaching an AI My Taste

- HN: [49629118](https://news.ycombinator.com/item?id=49629118)
- Source: [lonriesberg.com](https://lonriesberg.com/posts/teaching-an-ai-my-taste/)
- Score: 1
- Comments: 0
- Posted: 2026-09-09T16:27:49Z

## Translation

Title: Teaching an AI My Taste
Article title: Teaching an AI My Taste – Lon Riesberg
Description: What it takes to teach a machine judgment, and how I checked whether it learned.

Article text:
What it takes to teach a machine judgment, and how I checked whether it learned.
When putting together an issue of the Data Elixir newsletter , there are hundreds of relevant posts, tools, and resources to consider. Most of it gets to me through ~200 RSS feeds, 806 social media accounts, Hacker News, and Reddit.
I’ve written Python scripts to fetch and rank posts from the social media accounts and those scripts have become essential for catching what’s resonating around the web. But resonating isn’t the same as good , and it’s not the same as being right for Data Elixir either. That judgment still comes from me, reading. And even with the scripts helping to filter everything, I still end up with 100-200 posts to consider each week.
It would be great to automate some of that last-mile curation, but that kind of reading and judgment isn’t something traditional code can do.
I still make the final calls, but some of the links in the Data Elixir issue you read this week were surfaced by an LLM that I taught to think like me.
This post walks through how I built it, including the judge at its center, and the part I care about most: how I taught it my taste, and how I checked whether it actually learned. Along the way are the dead ends that taught me the most, and a prompt you can drop into a coding agent to start building something like this yourself.
Every morning, a Python script pulls new links from a collection of RSS feeds, an email inbox, and Hacker News. Then each link is followed and simplified to a common structure that includes a title, the link, where it came from, and a short snippet.
The script skips links that are stale or have already been seen, and then it passes what’s left to an LLM judge that scores each link against my taste. Whatever clears the bar gets ranked and sent to my Inbox.
That’s the entire application, and most all of it is plumbing.
The judge is the most important part. If I added more sources but didn’t have a good judge, I’d just be drowning faster with things to read. And the judge is only as good as the rubric it judges against. For this project, the rubric is defined in a plain markdown file I call taste.md.
taste.md is 111 lines of short descriptions of what I’m looking for. It’s organized into categories that cover who the audience is, positive signals that raise the score, and things that disqualify a post.
Those three categories cover everything I need, including specific traits like:
“Substantive and about working with data (or directly useful to people who do)”
“Teaches, demonstrates, or argues something”
“Reasonably fresh (roughly within the last month) and durable”
Writing taste.md is easy. Figuring out whether I actually got it right is not. And it didn’t start out as 111 lines, either.
My first attempt at a taste.md file was just me, describing what I look for and what I skip. “Teach me something. Posts should target a smart, data-savvy audience. No hype.”
I ran it against a week of candidates and read the results. They looked fine. Good stuff near the top, junk near the bottom. Great.
But that’s the trap. The results were plausible , but plausible isn’t the same as right. It’s the same trap as ranking by engagement. It’s a useful first filter, but a high number doesn’t prove that something is good. I needed a test that an LLM couldn’t talk its way through.
So I built a test using a spreadsheet of my own curation decisions. I started with 20 links that I had featured in Data Elixir and another 15 that I had passed on, and included a one-line reason for each.
The rejects, and especially the near misses, turned out to matter more than the picks. That boundary between “good” and “good-but-not-for-me” is where taste really lives. A slam-dunk “yes” doesn’t teach the judge anything it didn’t already know.
I did two things with that spreadsheet. The first was to write a script that reads each labeled article and works out the traits that separate the yeses from the nos, so I could fold them back into taste.md. The first time I ran it, I used only the featured examples and it confidently concluded that my audience was professional software engineers. For Data Elixir, that’s not true. Without the rejects there was no contrast, and everything that makes my taste specific rather than generic lives in what I say no to. Negatives carry the signal.
The second thing I did with that spreadsheet was the one that hurt. I used it for backtesting. Score all 35 links blind. No labels, no notes, just the judge, and then measure how often it agreed with me. The first number was 51%. That’s essentially a coin flip. And when I looked closer it was even worse than that. The median score of the things I had featured was 62, and the median of the things I’d rejected was also 62. The judge didn’t disagree with me so much as it had no opinion. Everything landed in the same undifferentiated pile.
But now I had a number to move instead of a feeling to satisfy. I tightened the traits, re-ran the backtest, and watched the results climb. First, 51%, then 57%, then 71%, and the medians finally pulled apart: 72 for featured, 48 for rejected. The judge had learned to discriminate.
Of course, it’s more complicated than that because the traits interact. At one point I softened a rule about engineering depth, because it was sinking a post that had been a popular pick in Data Elixir. The softening fixed that post but then floated a dense, math-heavy theory piece I probably wouldn’t feature. It was whack-a-mole.
The reality is, my labels are taste too, and my taste isn’t ground truth. What I think is a great link and what my readers are actually interested in are not always the same. So I added real engagement as a final judge.
I use a platform called beehiiv to publish Data Elixir, and it logs every time an actual person clicks a link in the newsletter. “Actual person” is important here because there are a LOT of bots that click links in the newsletter as they try to determine if something is spam. For me, that’s noise, but beehiiv uses a variety of mechanisms to identify and filter out bot clicks.
For the 35 links I had selected to help define the taste.md file, I pulled the engagement numbers from beehiiv and lined them up against the judge’s scores for the same posts. Two stories came out of that, and they pointed in opposite directions.
The first story was a “build attention from scratch” post that I had featured. The judge scored it as a 38, which was well below my bar. I would have been tempted to modify the taste.md file to push it up, but based on reader engagement, it was one of the weakest links in that issue. The judge wasn’t missing something. It was seeing something my own enthusiasm had talked me past.
The second story was a post from a well-known AI writer about what it feels like to work with a new model. The judge scored it 38 and buried it. That particluar post was the single most-clicked link in that issue. The judge was being hard on first-person posts, but readers love those posts.
I kept reworking the taste.md file and running backtests until the judge consistently scored like me. Or scored like my readers, actually.
The individual verdicts aren’t the point. The ladder is: vibes → my labels → real outcomes. Each step corrects the one before it. I started out by trusting my gut, but that wasn’t measurable, so I replaced it with my documented decisions. But those decisions didn’t always line up with what my readers actually clicked. Every step made the judge’s taste a little less mine-in-theory and a little more mine-in-practice. And sometimes the judge’s scores were better than mine.
Every morning before I’m awake, the whole thing runs and drops a ranked shortlist in my inbox:
The reason is the part that matters. The score tells me where something landed. The reason tells me why, and that’s what lets me accept or overrule it in about two seconds. “Clear, hands-on Bayesian workflow in Python” is a yes. “Package release announcement dressed up as an article” is not. I’m not reading 200 posts anymore. I’m reading 15 one-line arguments and then clicking through to read the 3 or 4 posts that are actually worth my time.
Nothing about this is specific to Data Elixir. The judge doesn’t know or care where a candidate came from. It could be from an RSS feed, a newsletter in an inbox, or a social media feed. Every source gets normalized into the same shape before it reaches the taste test. Which means the whole system can be aimed at something completely different by just changing the sources it reads and the taste.md it judges against. This same scheme will work to search job listings, find useful social media discussions, find relevant academic papers, etc.
To verify that, I stood up a second instance. It’s the same code, untouched and pointed at a list of the most popular blogs among Hacker News readers. After developing a new taste.md file, I had a completely different curator. This one reads a couple hundred of the web’s most-loved technical blogs and then hands me the posts that it thinks I’m most likely to be interested in. It’s super useful!
If you’re interested in building this, a Hacker News curator is a great place to start. To track what’s new, there’s a ranked list of popular Hacker News writers at HN Popularity Contest .
It’s a bit of work to pull out the URLs and then track down the RSS feeds, but here’s an OPML file that I created recently for this exact use-case. Just drop this file into your own code:
<?xml version= "1.0" ?>
< opml version= "2.0" >
< head >
< title >Top HN Writers: RSS Feeds</ title >
< dateCreated >Fri, 26 Jun 2026 16:22:56 +0000</ dateCreated >
</ head >
< body >
< outline type= "rss" text= "Paul Graham: Essays" title= "Paul Graham: Essays" xmlUrl= "http://www.aaronsw.com/2002/feeds/pgessays.rss" htmlUrl= "https://paulgraham.com" />
< outline type= "rss" text= "Krebs on Security" title= "Krebs on Security" xmlUrl= "https://krebsonsecurity.com/feed/" htmlUrl= "https://krebsonsecurity.com" />
< outline type= "rss" text= "Simon Willison's Weblog" title= "Simon Willison's Weblog" xmlUrl= "https://simonwillison.net/atom/everything/" htmlUrl= "https://simonwillison.net" />
< outline type= "rss" text= "Daring Fireball" title= "Daring Fireball" xmlUrl= "https://daringfireball.net/feeds/main" htmlUrl= "https://daringfireball.net" />
< outline type= "rss" text= "Julia Evans" title= "Julia Evans" xmlUrl= "https://jvns.ca/atom.xml" htmlUrl= "https://jvns.ca" />
< outline type= "rss" text= "danluu.com" title= "danluu.com" xmlUrl= "https://danluu.com/atom.xml" htmlUrl= "https://danluu.com" />
< outline type= "rss" text= "Ken Shirriff's blog" title= "Ken Shirriff's blog" xmlUrl= "https://www.righto.com/feeds/posts/default" htmlUrl= "https://righto.com" />
< outline type= "rss" text= "Stratechery by Ben Thompson" title= "Stratechery by Ben Thompson" xmlUrl= "https://stratechery.com/feed" htmlUrl= "https://stratechery.com" />
< outline type= "rss" text= "Troy Hunt" title= "Troy Hunt" xmlUrl= "https://www.troyhunt.com/rss/" htmlUrl= "https://troyhunt.com" />
< outline type= "rss" text= "Kalzumeus Software" title= "Kalzumeus Software" xmlUrl= "https://kalzumeus.com/feed/articles/" htmlUrl= "https://kalzumeus.com" />
< outline type= "rss" text= "Schneier on Security" title= "Schneier on Security" xmlUrl= "https://www.schneier.com/feed/" htmlUrl= "https://schneier.com" />
< outline type= "rss" text= "Terence Eden’s Blog" title= "Terence Eden’s Blog" xmlUrl= "https:
[truncated]
I’m not releasing all the code because it’s wired to my specific setup, but the code is the easy part anyway. A coding agent can build the whole scaffold in a couple hours. The two things that actually make it work are a taste.md that captures your judgment and the discipline to test it.
Here are a couple tips… First, once you have a simple scaffold, start with just a single source and the judge. Let that be the whole app initially, and once you like the judge’s selections,

[truncated]

## Original Extract

What it takes to teach a machine judgment, and how I checked whether it learned.

What it takes to teach a machine judgment, and how I checked whether it learned.
When putting together an issue of the Data Elixir newsletter , there are hundreds of relevant posts, tools, and resources to consider. Most of it gets to me through ~200 RSS feeds, 806 social media accounts, Hacker News, and Reddit.
I’ve written Python scripts to fetch and rank posts from the social media accounts and those scripts have become essential for catching what’s resonating around the web. But resonating isn’t the same as good , and it’s not the same as being right for Data Elixir either. That judgment still comes from me, reading. And even with the scripts helping to filter everything, I still end up with 100-200 posts to consider each week.
It would be great to automate some of that last-mile curation, but that kind of reading and judgment isn’t something traditional code can do.
I still make the final calls, but some of the links in the Data Elixir issue you read this week were surfaced by an LLM that I taught to think like me.
This post walks through how I built it, including the judge at its center, and the part I care about most: how I taught it my taste, and how I checked whether it actually learned. Along the way are the dead ends that taught me the most, and a prompt you can drop into a coding agent to start building something like this yourself.
Every morning, a Python script pulls new links from a collection of RSS feeds, an email inbox, and Hacker News. Then each link is followed and simplified to a common structure that includes a title, the link, where it came from, and a short snippet.
The script skips links that are stale or have already been seen, and then it passes what’s left to an LLM judge that scores each link against my taste. Whatever clears the bar gets ranked and sent to my Inbox.
That’s the entire application, and most all of it is plumbing.
The judge is the most important part. If I added more sources but didn’t have a good judge, I’d just be drowning faster with things to read. And the judge is only as good as the rubric it judges against. For this project, the rubric is defined in a plain markdown file I call taste.md.
taste.md is 111 lines of short descriptions of what I’m looking for. It’s organized into categories that cover who the audience is, positive signals that raise the score, and things that disqualify a post.
Those three categories cover everything I need, including specific traits like:
“Substantive and about working with data (or directly useful to people who do)”
“Teaches, demonstrates, or argues something”
“Reasonably fresh (roughly within the last month) and durable”
Writing taste.md is easy. Figuring out whether I actually got it right is not. And it didn’t start out as 111 lines, either.
My first attempt at a taste.md file was just me, describing what I look for and what I skip. “Teach me something. Posts should target a smart, data-savvy audience. No hype.”
I ran it against a week of candidates and read the results. They looked fine. Good stuff near the top, junk near the bottom. Great.
But that’s the trap. The results were plausible , but plausible isn’t the same as right. It’s the same trap as ranking by engagement. It’s a useful first filter, but a high number doesn’t prove that something is good. I needed a test that an LLM couldn’t talk its way through.
So I built a test using a spreadsheet of my own curation decisions. I started with 20 links that I had featured in Data Elixir and another 15 that I had passed on, and included a one-line reason for each.
The rejects, and especially the near misses, turned out to matter more than the picks. That boundary between “good” and “good-but-not-for-me” is where taste really lives. A slam-dunk “yes” doesn’t teach the judge anything it didn’t already know.
I did two things with that spreadsheet. The first was to write a script that reads each labeled article and works out the traits that separate the yeses from the nos, so I could fold them back into taste.md. The first time I ran it, I used only the featured examples and it confidently concluded that my audience was professional software engineers. For Data Elixir, that’s not true. Without the rejects there was no contrast, and everything that makes my taste specific rather than generic lives in what I say no to. Negatives carry the signal.
The second thing I did with that spreadsheet was the one that hurt. I used it for backtesting. Score all 35 links blind. No labels, no notes, just the judge, and then measure how often it agreed with me. The first number was 51%. That’s essentially a coin flip. And when I looked closer it was even worse than that. The median score of the things I had featured was 62, and the median of the things I’d rejected was also 62. The judge didn’t disagree with me so much as it had no opinion. Everything landed in the same undifferentiated pile.
But now I had a number to move instead of a feeling to satisfy. I tightened the traits, re-ran the backtest, and watched the results climb. First, 51%, then 57%, then 71%, and the medians finally pulled apart: 72 for featured, 48 for rejected. The judge had learned to discriminate.
Of course, it’s more complicated than that because the traits interact. At one point I softened a rule about engineering depth, because it was sinking a post that had been a popular pick in Data Elixir. The softening fixed that post but then floated a dense, math-heavy theory piece I probably wouldn’t feature. It was whack-a-mole.
The reality is, my labels are taste too, and my taste isn’t ground truth. What I think is a great link and what my readers are actually interested in are not always the same. So I added real engagement as a final judge.
I use a platform called beehiiv to publish Data Elixir, and it logs every time an actual person clicks a link in the newsletter. “Actual person” is important here because there are a LOT of bots that click links in the newsletter as they try to determine if something is spam. For me, that’s noise, but beehiiv uses a variety of mechanisms to identify and filter out bot clicks.
For the 35 links I had selected to help define the taste.md file, I pulled the engagement numbers from beehiiv and lined them up against the judge’s scores for the same posts. Two stories came out of that, and they pointed in opposite directions.
The first story was a “build attention from scratch” post that I had featured. The judge scored it as a 38, which was well below my bar. I would have been tempted to modify the taste.md file to push it up, but based on reader engagement, it was one of the weakest links in that issue. The judge wasn’t missing something. It was seeing something my own enthusiasm had talked me past.
The second story was a post from a well-known AI writer about what it feels like to work with a new model. The judge scored it 38 and buried it. That particluar post was the single most-clicked link in that issue. The judge was being hard on first-person posts, but readers love those posts.
I kept reworking the taste.md file and running backtests until the judge consistently scored like me. Or scored like my readers, actually.
The individual verdicts aren’t the point. The ladder is: vibes → my labels → real outcomes. Each step corrects the one before it. I started out by trusting my gut, but that wasn’t measurable, so I replaced it with my documented decisions. But those decisions didn’t always line up with what my readers actually clicked. Every step made the judge’s taste a little less mine-in-theory and a little more mine-in-practice. And sometimes the judge’s scores were better than mine.
Every morning before I’m awake, the whole thing runs and drops a ranked shortlist in my inbox:
The reason is the part that matters. The score tells me where something landed. The reason tells me why, and that’s what lets me accept or overrule it in about two seconds. “Clear, hands-on Bayesian workflow in Python” is a yes. “Package release announcement dressed up as an article” is not. I’m not reading 200 posts anymore. I’m reading 15 one-line arguments and then clicking through to read the 3 or 4 posts that are actually worth my time.
Nothing about this is specific to Data Elixir. The judge doesn’t know or care where a candidate came from. It could be from an RSS feed, a newsletter in an inbox, or a social media feed. Every source gets normalized into the same shape before it reaches the taste test. Which means the whole system can be aimed at something completely different by just changing the sources it reads and the taste.md it judges against. This same scheme will work to search job listings, find useful social media discussions, find relevant academic papers, etc.
To verify that, I stood up a second instance. It’s the same code, untouched and pointed at a list of the most popular blogs among Hacker News readers. After developing a new taste.md file, I had a completely different curator. This one reads a couple hundred of the web’s most-loved technical blogs and then hands me the posts that it thinks I’m most likely to be interested in. It’s super useful!
If you’re interested in building this, a Hacker News curator is a great place to start. To track what’s new, there’s a ranked list of popular Hacker News writers at HN Popularity Contest .
It’s a bit of work to pull out the URLs and then track down the RSS feeds, but here’s an OPML file that I created recently for this exact use-case. Just drop this file into your own code:
<?xml version= "1.0" ?>
< opml version= "2.0" >
< head >
< title >Top HN Writers: RSS Feeds</ title >
< dateCreated >Fri, 26 Jun 2026 16:22:56 +0000</ dateCreated >
</ head >
< body >
< outline type= "rss" text= "Paul Graham: Essays" title= "Paul Graham: Essays" xmlUrl= "http://www.aaronsw.com/2002/feeds/pgessays.rss" htmlUrl= "https://paulgraham.com" />
< outline type= "rss" text= "Krebs on Security" title= "Krebs on Security" xmlUrl= "https://krebsonsecurity.com/feed/" htmlUrl= "https://krebsonsecurity.com" />
< outline type= "rss" text= "Simon Willison's Weblog" title= "Simon Willison's Weblog" xmlUrl= "https://simonwillison.net/atom/everything/" htmlUrl= "https://simonwillison.net" />
< outline type= "rss" text= "Daring Fireball" title= "Daring Fireball" xmlUrl= "https://daringfireball.net/feeds/main" htmlUrl= "https://daringfireball.net" />
< outline type= "rss" text= "Julia Evans" title= "Julia Evans" xmlUrl= "https://jvns.ca/atom.xml" htmlUrl= "https://jvns.ca" />
< outline type= "rss" text= "danluu.com" title= "danluu.com" xmlUrl= "https://danluu.com/atom.xml" htmlUrl= "https://danluu.com" />
< outline type= "rss" text= "Ken Shirriff's blog" title= "Ken Shirriff's blog" xmlUrl= "https://www.righto.com/feeds/posts/default" htmlUrl= "https://righto.com" />
< outline type= "rss" text= "Stratechery by Ben Thompson" title= "Stratechery by Ben Thompson" xmlUrl= "https://stratechery.com/feed" htmlUrl= "https://stratechery.com" />
< outline type= "rss" text= "Troy Hunt" title= "Troy Hunt" xmlUrl= "https://www.troyhunt.com/rss/" htmlUrl= "https://troyhunt.com" />
< outline type= "rss" text= "Kalzumeus Software" title= "Kalzumeus Software" xmlUrl= "https://kalzumeus.com/feed/articles/" htmlUrl= "https://kalzumeus.com" />
< outline type= "rss" text= "Schneier on Security" title= "Schneier on Security" xmlUrl= "https://www.schneier.com/feed/" htmlUrl= "https://schneier.com" />
< outline type= "rss" text= "Terence Eden’s Blog" title= "Terence Eden’s Blog" xmlUrl= "https:
[truncated]
I’m not releasing all the code because it’s wired to my specific setup, but the code is the easy part anyway. A coding agent can build the whole scaffold in a couple hours. The two things that actually make it work are a taste.md that captures your judgment and the discipline to test it.
Here are a couple tips… First, once you have a simple scaffold, start with just a single source and the judge. Let that be the whole app initially, and once you like the judge’s selections,

[truncated]
