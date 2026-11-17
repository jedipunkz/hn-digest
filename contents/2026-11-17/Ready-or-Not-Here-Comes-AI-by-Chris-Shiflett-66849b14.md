---
source: "https://shiflett.org/blog/2026/ready-or-not-here-comes-ai"
hn_url: "https://news.ycombinator.com/item?id=50533000"
title: "Ready or Not, Here Comes AI, by Chris Shiflett"
article_title: "Ready or Not, Here Comes AI, by Chris Shiflett"
image: "https://files.shiflett.org/images/ready-or-not.png"
author: "speckx"
captured_at: "2026-11-17T16:44:10Z"
capture_tool: "hn-digest"
hn_id: 50533000
score: 1
comments: 0
posted_at: "2026-11-17T15:46:17Z"
tags:
  - hacker-news
---

# Ready or Not, Here Comes AI, by Chris Shiflett

- HN: [50533000](https://news.ycombinator.com/item?id=50533000)
- Source: [shiflett.org](https://shiflett.org/blog/2026/ready-or-not-here-comes-ai)
- Score: 1
- Comments: 0
- Posted: 2026-11-17T15:46:17Z

## Translation

Title: Ready or Not, Here Comes AI, by Chris Shiflett
Description: On being found by AI: what doing things right gets you, and what it doesn’t.

Article text:
Maroon Bells, CO
39.0508° N
107.0208° W
A lot of robots read my website. Probably yours, too .
It feels weird, doesn’t it? But maybe it shouldn’t. Bots have been reading our websites for a long time, and for most of that time, we’ve been making sure they can.
In fact, bots — particularly those from search engines like Google — were the justification for doing things right for many years. We didn’t do it for the bots, though. We wanted Google to recommend our websites to people who might be interested in them.
For its first fifteen years or so, Google’s bot (aptly named Googlebot) was functionally a text browser. Anything not in the HTML didn’t exist. Honestly, it was glorious . It might not have been Google’s intention, but I believe they enforced a standard of discipline that is sorely missed.
When Ajax and JS-only developers first arrived, Google answered with awkward half measures, asking developers to serve pre-rendered HTML snapshots. You might remember the hashbang hell era of the web. Ugly URLs aside, at least the story remained clear: bots consume HTML; rendering it is your job.
Then, in 2014, Google caved . They announced that Googlebot would execute JavaScript and render pages like a browser. Combine this with an emerging supermajority of web developers who knew nothing but JavaScript, and what followed wasn’t hard to predict: SPAs , CSR , etc. The web became a shell that browsers had to assemble to get the real content. Broken and slow replaced reliable and fast. We’re still suffering the consequences of this era.
There was a period between 2014 and 2019 during which Googlebot lagged Chromium significantly. The evergreen Googlebot (always-current Chromium) didn’t arrive until 2019. Five years is a long time on the internet, and “Google executes JavaScript now” became a catch-all excuse. If it worked in Chrome, it must be fine. That’s how a generation of websites were hiding by accident, long before AI.
GPTBot, ClaudeBot, and their peers are a lot like 2004 Googlebot, and that’s a good thing. They fetch HTML and largely don’t execute JavaScript.
I’d hoped to tell you this was the second coming of text-only bots. It might be, but recently OpenAI’s crawlers appeared to start running scripts on some sites, without a word about it. Sad trombone.
We’ve seen this movie, and the dangerous part is coming soon. A bot that sometimes executes JavaScript is all the excuse some people need to stop caring whether it has to.
HTML works for everyone. It always has.
More and more people ask AI before they search, or they search and get an answer from AI anyway. Whether AI can find you matters regardless of how you feel about it.
Some people hide on purpose . Some hide behind a JavaScript wall without knowing it. Some simply relaxed their standards when it seemed like no one had any.
None of those describe me. I build with care. My websites are fast and reliable. I obsess over URLs and HTML. I’ve spent thirty years trying to do things right, and I assumed that was enough. I have multiple websites I’d like people to find, and as far as I could tell, AI wasn’t recommending them.
I decided to build some tools to figure it all out: one to check whether AI could find me, and one to check whether it did.
Soon after I finished writing this, OpenAI started letting businesses claim a profile in ChatGPT and see how often they’re mentioned, which is great. It still doesn’t tell you why.
I assumed there was a lot to learn, so I went looking for tips and tricks, because plenty of people are selling them. There really aren’t any. The closest thing is llms.txt , a proposal with no evidence that any bot reads it. Everything else is what Google recommended twenty years ago, and for the same reason.
The companies building these tools want to give people the best answer, so anything easy to game stops working the moment enough people game it.
They’re also operating at massive scale, which means the easier you can make it for the bots to find what they’re looking for, the more likely they will.
Don’t try to outsmart them. Help them.
So the first tool checks the basics:
Can AI get in? A robots.txt that allows AI bots.
Can it find everything? A good sitemap.
Can it read what it finds? Real content in real HTML.
Can it tell it’s you? A little structured data that says who you are.
That last one might sound strange, but I share my name with a guitarist from the Foo Fighters (a good reminder that names are not identifiers). Plenty of people, and a few AI answers, claim structured data matters a lot. No one has proven it, so call it a hunch. It’s cheap, it makes sense to me, and over time I’ll find out.
The usual way to write structured data is JSON-LD using names defined by schema.org , like jobTitle and worksFor . A few lines identify you, link to your website, and list your profiles elsewhere. You can view source to see mine. If you’re a fan of microformats , an h-card can do the same job by marking up what’s already on the page.
Whether AI can find you is the easy part, but it’s also one a lot of people get wrong, because failures are largely invisible.
This is where the rubber meets the road. Now that AI can find you, does it?
I started by asking the questions I want to be the answer to:
What’s the best proposal and invoicing software for design studios? ( Studioworks )
What invoicing app doesn’t take a cut of every payment? ( Studioworks )
Who should I hire to turn my idea into an MVP? ( Faculty )
What should I use to keep track of all my kids’ school emails? ( Schoolcase )
What’s the best independent coworking space in Boulder? ( Roost )
At first, I asked by hand. Then I built the second tool to ask every week, and to ask more than one model, and to ask more than once, because the same question rarely gets the same answer twice. A single answer is like a weather report. What I want is the trend.
When I asked about keeping track of school emails, I got a lot of advice about folders and filters, and no mention of Schoolcase. Then I added six words that changed everything:
Schoolcase was the top recommendation.
Those six words made AI search instead of answering from memory. Ask for advice, and it tells you what it already knows. Ask for a tool, and it goes looking. That seemed promising, but nobody asks for a tool they don’t know exists.
So “does AI find you” is really three questions:
When AI searches, does it find you? This is the part you control.
Does it choose to search? That depends on the question being asked, not on you.
Does it already know you? That’s memory, and memory is slow.
A model doesn’t learn as it goes. It learns in batches, and a new batch ships every few months with a later cutoff. Whatever the web says about you when a batch is collected is what that model knows until the next one. Pawel Jozefiak tracks this , and most models are already several months behind on the day they’re released.
There’s a big gap between can AI find you and does AI find you. That gap is authority: what others say about you counts more than what you say about yourself.
I deliberately don’t use an equals sign, because these are ingredients, not a formula. Your own website makes you findable when AI searches. Other people’s websites give you credibility, and both influence how you show up in results.
If that sounds familiar, it should. Google’s big idea was that the best measure of a page is what other pages say about it, counted in links. AI appears to have landed in a similar place.
Naturally, I built a third tool, and this one was a doozy. Measuring authority means leaning on research where it exists, hunches where it doesn’t, and a lot of testing. I’ve spent months on it. Every judgment the tool makes gets checked against pages I’ve judged by hand, and I’m constantly making adjustments. Only time will tell if I’ve got it right.
Here’s my biggest hunch: authority belongs to a question . A link from a friend’s blog tells AI you exist. It doesn’t tell AI you’re the answer to anything. To be the answer, you have to show up alongside the question, in the places AI reads when someone asks it.
I could be wrong. I’m measuring it, and I’ll share what I find either way.
If I’m right, the path is clear, but integrity is still key. Asking to be included where you belong is fair. A coworking space missing from a list of coworking spaces in its own town is an omission, and asking to be listed is honest. Planting your own name in a forum thread is gaming, and we know how that ends. Comment spam is the reason links in comments stopped counting.
A lot of robots read my website. Now they can get in, find everything, read it, and tell it’s me. What they say about me next depends on what everyone else says about me. That takes time, so I’m starting now.
Chris Shiflett
Boulder-based founder, designer, and developer. Co-founder of Studioworks and Schoolcase , and founder of Faculty , a 0→1 product studio. Maker of Hide and Seek , which shows whether AI can find your website. More about Chris →

## Original Extract

On being found by AI: what doing things right gets you, and what it doesn’t.

Maroon Bells, CO
39.0508° N
107.0208° W
A lot of robots read my website. Probably yours, too .
It feels weird, doesn’t it? But maybe it shouldn’t. Bots have been reading our websites for a long time, and for most of that time, we’ve been making sure they can.
In fact, bots — particularly those from search engines like Google — were the justification for doing things right for many years. We didn’t do it for the bots, though. We wanted Google to recommend our websites to people who might be interested in them.
For its first fifteen years or so, Google’s bot (aptly named Googlebot) was functionally a text browser. Anything not in the HTML didn’t exist. Honestly, it was glorious . It might not have been Google’s intention, but I believe they enforced a standard of discipline that is sorely missed.
When Ajax and JS-only developers first arrived, Google answered with awkward half measures, asking developers to serve pre-rendered HTML snapshots. You might remember the hashbang hell era of the web. Ugly URLs aside, at least the story remained clear: bots consume HTML; rendering it is your job.
Then, in 2014, Google caved . They announced that Googlebot would execute JavaScript and render pages like a browser. Combine this with an emerging supermajority of web developers who knew nothing but JavaScript, and what followed wasn’t hard to predict: SPAs , CSR , etc. The web became a shell that browsers had to assemble to get the real content. Broken and slow replaced reliable and fast. We’re still suffering the consequences of this era.
There was a period between 2014 and 2019 during which Googlebot lagged Chromium significantly. The evergreen Googlebot (always-current Chromium) didn’t arrive until 2019. Five years is a long time on the internet, and “Google executes JavaScript now” became a catch-all excuse. If it worked in Chrome, it must be fine. That’s how a generation of websites were hiding by accident, long before AI.
GPTBot, ClaudeBot, and their peers are a lot like 2004 Googlebot, and that’s a good thing. They fetch HTML and largely don’t execute JavaScript.
I’d hoped to tell you this was the second coming of text-only bots. It might be, but recently OpenAI’s crawlers appeared to start running scripts on some sites, without a word about it. Sad trombone.
We’ve seen this movie, and the dangerous part is coming soon. A bot that sometimes executes JavaScript is all the excuse some people need to stop caring whether it has to.
HTML works for everyone. It always has.
More and more people ask AI before they search, or they search and get an answer from AI anyway. Whether AI can find you matters regardless of how you feel about it.
Some people hide on purpose . Some hide behind a JavaScript wall without knowing it. Some simply relaxed their standards when it seemed like no one had any.
None of those describe me. I build with care. My websites are fast and reliable. I obsess over URLs and HTML. I’ve spent thirty years trying to do things right, and I assumed that was enough. I have multiple websites I’d like people to find, and as far as I could tell, AI wasn’t recommending them.
I decided to build some tools to figure it all out: one to check whether AI could find me, and one to check whether it did.
Soon after I finished writing this, OpenAI started letting businesses claim a profile in ChatGPT and see how often they’re mentioned, which is great. It still doesn’t tell you why.
I assumed there was a lot to learn, so I went looking for tips and tricks, because plenty of people are selling them. There really aren’t any. The closest thing is llms.txt , a proposal with no evidence that any bot reads it. Everything else is what Google recommended twenty years ago, and for the same reason.
The companies building these tools want to give people the best answer, so anything easy to game stops working the moment enough people game it.
They’re also operating at massive scale, which means the easier you can make it for the bots to find what they’re looking for, the more likely they will.
Don’t try to outsmart them. Help them.
So the first tool checks the basics:
Can AI get in? A robots.txt that allows AI bots.
Can it find everything? A good sitemap.
Can it read what it finds? Real content in real HTML.
Can it tell it’s you? A little structured data that says who you are.
That last one might sound strange, but I share my name with a guitarist from the Foo Fighters (a good reminder that names are not identifiers). Plenty of people, and a few AI answers, claim structured data matters a lot. No one has proven it, so call it a hunch. It’s cheap, it makes sense to me, and over time I’ll find out.
The usual way to write structured data is JSON-LD using names defined by schema.org , like jobTitle and worksFor . A few lines identify you, link to your website, and list your profiles elsewhere. You can view source to see mine. If you’re a fan of microformats , an h-card can do the same job by marking up what’s already on the page.
Whether AI can find you is the easy part, but it’s also one a lot of people get wrong, because failures are largely invisible.
This is where the rubber meets the road. Now that AI can find you, does it?
I started by asking the questions I want to be the answer to:
What’s the best proposal and invoicing software for design studios? ( Studioworks )
What invoicing app doesn’t take a cut of every payment? ( Studioworks )
Who should I hire to turn my idea into an MVP? ( Faculty )
What should I use to keep track of all my kids’ school emails? ( Schoolcase )
What’s the best independent coworking space in Boulder? ( Roost )
At first, I asked by hand. Then I built the second tool to ask every week, and to ask more than one model, and to ask more than once, because the same question rarely gets the same answer twice. A single answer is like a weather report. What I want is the trend.
When I asked about keeping track of school emails, I got a lot of advice about folders and filters, and no mention of Schoolcase. Then I added six words that changed everything:
Schoolcase was the top recommendation.
Those six words made AI search instead of answering from memory. Ask for advice, and it tells you what it already knows. Ask for a tool, and it goes looking. That seemed promising, but nobody asks for a tool they don’t know exists.
So “does AI find you” is really three questions:
When AI searches, does it find you? This is the part you control.
Does it choose to search? That depends on the question being asked, not on you.
Does it already know you? That’s memory, and memory is slow.
A model doesn’t learn as it goes. It learns in batches, and a new batch ships every few months with a later cutoff. Whatever the web says about you when a batch is collected is what that model knows until the next one. Pawel Jozefiak tracks this , and most models are already several months behind on the day they’re released.
There’s a big gap between can AI find you and does AI find you. That gap is authority: what others say about you counts more than what you say about yourself.
I deliberately don’t use an equals sign, because these are ingredients, not a formula. Your own website makes you findable when AI searches. Other people’s websites give you credibility, and both influence how you show up in results.
If that sounds familiar, it should. Google’s big idea was that the best measure of a page is what other pages say about it, counted in links. AI appears to have landed in a similar place.
Naturally, I built a third tool, and this one was a doozy. Measuring authority means leaning on research where it exists, hunches where it doesn’t, and a lot of testing. I’ve spent months on it. Every judgment the tool makes gets checked against pages I’ve judged by hand, and I’m constantly making adjustments. Only time will tell if I’ve got it right.
Here’s my biggest hunch: authority belongs to a question . A link from a friend’s blog tells AI you exist. It doesn’t tell AI you’re the answer to anything. To be the answer, you have to show up alongside the question, in the places AI reads when someone asks it.
I could be wrong. I’m measuring it, and I’ll share what I find either way.
If I’m right, the path is clear, but integrity is still key. Asking to be included where you belong is fair. A coworking space missing from a list of coworking spaces in its own town is an omission, and asking to be listed is honest. Planting your own name in a forum thread is gaming, and we know how that ends. Comment spam is the reason links in comments stopped counting.
A lot of robots read my website. Now they can get in, find everything, read it, and tell it’s me. What they say about me next depends on what everyone else says about me. That takes time, so I’m starting now.
Chris Shiflett
Boulder-based founder, designer, and developer. Co-founder of Studioworks and Schoolcase , and founder of Faculty , a 0→1 product studio. Maker of Hide and Seek , which shows whether AI can find your website. More about Chris →
