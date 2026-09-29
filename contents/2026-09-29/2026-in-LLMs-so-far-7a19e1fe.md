---
source: "https://simonwillison.net/2026/Sep/27/2026-in-llms-so-far/"
hn_url: "https://news.ycombinator.com/item?id=49888261"
title: "2026 in LLMs (so far)"
article_title: "2026 in LLMs (so far)"
image: "https://static.simonwillison.net/static/2026/2026-in-llms/simon-willison-2026-in-llms-png.001.webp"
author: "matt_d"
captured_at: "2026-09-29T04:34:33Z"
capture_tool: "hn-digest"
hn_id: 49888261
score: 1
comments: 0
posted_at: "2026-09-29T04:24:36Z"
tags:
  - hacker-news
---

# 2026 in LLMs (so far)

- HN: [49888261](https://news.ycombinator.com/item?id=49888261)
- Source: [simonwillison.net](https://simonwillison.net/2026/Sep/27/2026-in-llms-so-far/)
- Score: 1
- Comments: 0
- Posted: 2026-09-29T04:24:36Z

## Translation

Title: 2026 in LLMs (so far)
Description: On Friday I gave the closing keynote at the WeAreDevelopers World Congress North America in San Jose. I tied together the key trends from the past year into a chronological …

Article text:
2026 in LLMs (so far)
Simon Willison’s Weblog
On Friday I gave the closing keynote at the WeAreDevelopers World Congress North America in San Jose. I tied together the key trends from the past year into a chronological exploration of everything that happened in 2026. The video is on YouTube ; here are my annotated slides and notes to accompany the talk.
And as an annotated presentation :
I’m going to give a lightning tour of everything that has happened so far in 2026. The year isn’t over yet!
For me, 2026 started a couple of months earlier in November 2025.
November saw the release of two important models: Claude Opus 4.5 and GPT-5.1.
As is usually the case with new models, these were incremental improvements on the models that came before them.
But every now and then when a model improves, it crosses an invisible line where something that didn’t really work starts working.
In this case, the thing that started working was their coding agents. Claude Code had been around since February 2025; Codex was a little younger.
These two new models, when paired with their respective coding agent harnesses, improved from “often make mistakes” to “reliable enough to use on a day-to-day basis”.
For a couple of years now I’ve been evaluating new models by asking them to “Generate an SVG of a pelican riding a bicycle”. It’s probably the world’s stupidest benchmark—there’s only so much you can learn from it.
But it’s still a challenge for models, because drawing pelicans is difficult, drawing bicycles is difficult, and pelicans can’t ride bicycles in the first place.
Here’s the state of the art for November. Claude still couldn’t really draw a bicycle! The GPT-5.1 bicycle frame is pretty crap too.
Also in November, we had the first commit to an obscure GitHub repository called “Warelay”. We’ll come back to this repository shortly.
And then there were the December holidays, and individual developers took some time off and many started tinkering with these new coding agent model combinations... and it began to dawn on us quite how much they could do that they couldn’t do before.
Come January, a lot of us were quite excited to start putting this stuff into action.
Every year I set myself a New Year’s resolution, and for as long as I can remember it’s been the same thing: stay focused. Take on less new projects. Try to get things done in the projects I already have.
This year I decided that since that had never worked before, I’d go the other way.
We’ve got coding agents now, let’s see what they can do. I’m going to take on as many new projects as I like!
(You can ask me at the end of the year if this turned out to be a good idea or not. I have a lot of plates spinning right now.)
“Be more ambitious” has been something of a theme for the year, because the only way to find the limits of this technology is to keep on pushing them until they don’t work.
I also went on the Oxide and friends podcast with Bryan Cantrill and Adam Leventhal to share predictions for the next year (and three and six years).
With hindsight, my LLM predictions were pretty unambitious.
I said “it will become undeniable that LLMs write good code”—I think we’re there now.
I predicted we would finally solve sandboxing. I counted and around 40 of the 277 sessions at this conference touched on sandboxing or agent security in some way, so we’re at least putting a lot of effort into that!
I predicted “a Challenger disaster” for coding agent security. There’s certainly been a whole lot of noise around agent security this year, though the exact disaster I predicted (with coding agents being hijacked and causing real-world economic damage) hasn’t really played out.
We threw in a joke prediction that the Pope would weigh in on the economic impact of LLMs.
I also predicted that New Zealand’s Kākāpō parrots would have an outstanding breeding season this year.
These are flightless nocturnal parrots. They’re kind of dumpy looking, I think they’re beautiful, and there were only 236 of these parrots in the world at the start of the year.
Kākāpō only breed when the Rimu trees have a big fruiting season, and that hasn’t happened in four years... but this year the Rimu fruit were looking excellent.
Also on that podcast, we coined a term (full credit to Adam) for “that feeling of AI induced ennui where software engineers get listless because the AI can do anything”.
This has been a major theme throughout the year, and was touched on by several speakers at this conference.
As a software engineer, I’ve never had a year of my career where everything has changed so quickly and so dramatically.
A lot of what I’ve been doing this year is trying to come to terms with that and what that means for my own profession.
Also in January, I suffered from what I’m calling AI mania .
This is not the same thing as AI psychosis .
With AI mania, any time your agent isn’t building something for you feels like wasted time. You’re losing sleep because you could be staying up later getting your agents to do stuff.
My AI mania presented itself in some ridiculously over-ambitious projects.
I built a JavaScript interpreter entirely in Python , vibe-ported from MicroQuickJS by Fabrice Bellard.
Then I built a WebAssembly runtime in Python as well .
These projects were quite useful, in that they sort of cured me of my AI mania... because after I built these things, I got to look at them and ask “does the world need a slow, buggy, half-baked Python JavaScript interpreter?”
I did get this out of it: https://simonw.github.io/micro-javascript/playground.html
This page runs my JavaScript interpreter built in Python, running in Python using Pyodide , which is Python compiled to WebAssembly, running in JavaScript, running in a browser.
It’s a beautiful stack of horrors. I’ve been having a lot of fun with WebAssembly this year.
By the end of January, that repository we saw start in November had renamed itself, first to CLAWDIS, then CLAWDBOT, then Moltbot, and finally to OpenClaw.
At this point OpenClaw had 8,300 commits, less than two months after the project had started. I looked today and it’s over 100,000 commits now!
This is the most vibe-coded piece of software in existence.
(Here’s how I generated that list of name changes .)
This kicked off the OpenClaw revolution. It effectively defined a new category of software.
There’s a generic term for this which I really enjoy. We call software like this a “Claw”. There’s OpenClaw, NanoClaw , IronClaw , PicoClaw ...
Today they’re being rebranded as “personal agents” or “general agents”, but I still like to think of them as Claws.
The Apple stores in the Bay Area sold out of Mac Minis because so many people were buying Mac Minis to run OpenClaw!
Drew Breunig said that this is because your OpenClaw is a digital pet, and you buy a Mac mini as an aquarium to keep your claw in, which is kind of delightful.
Also in January, we had this website.
This was MoltBook , a social network for AI agents, where the idea was that you send your Claw to go and talk to all of the other Claws, because what could possibly go wrong if you did that?
The website launched on Thursday. It blew up on Friday . It was profiled by the New York Times on Monday . And by Tuesday, everyone had forgotten it existed as it drowned in a deluge of slop and spam.
Facebook/Meta bought it a month later .
In February, a company called StrongDM described what they called their Software Factory.
They wrote about this in Software Factories and the Agentic Moment . I posted my own notes at the time, having seen their demo in person back in October.
Dan Shapiro called this approach the Dark Factory , after the idea that if your factory is sufficiently automated you can turn the lights out, because you don’t even need to see what’s going on.
StrongDM presented two rules for software development that they’d been following since July last year.
The first was code must not be written by humans .
Any code that you write has to have been routed through a coding agent.
This sounded radical in February, but I imagine there are a lot of people in this room who are pretty much living that today.
Rule number two was code must not be reviewed by humans .
You’re not allowed to read the code!
This continued to be a huge topic for much of this year. Many of the sessions at this event have been about code review and how you can get away with this.
What I found interesting about StrongDM is that they were living six months ahead of the rest of us, and they’d been exploring what it means to build software, not read the code, but still be confident that the software is of high quality. What can you do with these agents to help verify their work?
StrongDM are a security company, and they had people with decades of experience on this project. They were very much exploring the edges of what’s possible and responsible to do with this stuff.
Also in February: First kākāpō chick in four years hatches on Valentine’s Day . Breeding season is off to a good start!
Also in February... Google released Gemini 3.1 Pro . That’s a pretty great pelican riding a bicycle! It’s got the chain in the right place, it’s got feet on both sides. There’s a little fish in the basket.
And then Google’s Jeff Dean tweeted a video comparing Gemini 3 Pro and Gemini 3.1 Pro that featured an animated pelican riding a bicycle, a frog on a penny-farthing, a giraffe driving a tiny car, an ostrich on roller skates, a turtle kickflipping a skateboard, and a dachshund driving a stretch limousine.
This was frustrating, because my protection for the pelican riding the bicycle test was always “if they draw a perfect pelican on a bicycle, I’ll ask for some other animal on something else.”
Google trained for all forms of animals on all forms of transport! They’ve defeated my benchmark at this point.
The other thing that started in February was Tokenmaxxing . We had headlines about Meta making AI adoption a formal part of performance reviews, and Microsoft wanting every employee to use AI, and Uber boasting that 90% of their engineers were using AI workflows.
Then a few months later we have Meta cracking down on token use, Microsoft saying tokenmaxxing is “not what we are optimizing for”, and Uber capping employee AI spending.
So tokenmaxxing went straight up and then straight back down again—because it turns out the agents are expensive .
Last year it was difficult to spend more than $50 on AI tokens, because we didn’t have anything interesting to do with them. Then agents blew up, and now you can actually spend $1,000 in a day doing real work.
This is also the reason that Anthropic’s valuation skyrocketed to maybe a trillion dollars.
AI appears to have hit product market fit in 2026, primarily through coding agents.
In March, we hit peak OpenClaw.
These photographs are from China, where companies hosted OpenClaw install parties which saw non-tech-nerds queueing up around the block for help getting Claws installed on their personal devices.
I think this proved real market demand for this class of Claws, or personal AI agents. It turns out regular people really do want a weird little AI agent that can do useful things on their behalf.
A Claw is really just a coding agent wearing a less threatening hat. Under the hood they work much the same way—writing and then executing code on your computer to get stuff done.
The race was on to be the first to build a safe Claw —a Claw you could give to regular human beings where they wouldn’t instantly shoot themselves in the foot.
Meta’s Muse came out three weeks ago and is currently at the top of the free charts on the iPhone App Store. It appears to be taking off with consumers.
I’m not yet convinced you can’t shoot yourself in the foot with Muse, but I guess we’ll find out for sure pretty soon.
Photos from How the OpenClaw Frenzy Is Testing China’s AI Commitment (March 29th) and The Enthusiasm and Anxiety Behind China’s OpenClaw Craze (April 8th, 2026).
In April, we had a model release where the mode

[truncated]

## Original Extract

On Friday I gave the closing keynote at the WeAreDevelopers World Congress North America in San Jose. I tied together the key trends from the past year into a chronological …

2026 in LLMs (so far)
Simon Willison’s Weblog
On Friday I gave the closing keynote at the WeAreDevelopers World Congress North America in San Jose. I tied together the key trends from the past year into a chronological exploration of everything that happened in 2026. The video is on YouTube ; here are my annotated slides and notes to accompany the talk.
And as an annotated presentation :
I’m going to give a lightning tour of everything that has happened so far in 2026. The year isn’t over yet!
For me, 2026 started a couple of months earlier in November 2025.
November saw the release of two important models: Claude Opus 4.5 and GPT-5.1.
As is usually the case with new models, these were incremental improvements on the models that came before them.
But every now and then when a model improves, it crosses an invisible line where something that didn’t really work starts working.
In this case, the thing that started working was their coding agents. Claude Code had been around since February 2025; Codex was a little younger.
These two new models, when paired with their respective coding agent harnesses, improved from “often make mistakes” to “reliable enough to use on a day-to-day basis”.
For a couple of years now I’ve been evaluating new models by asking them to “Generate an SVG of a pelican riding a bicycle”. It’s probably the world’s stupidest benchmark—there’s only so much you can learn from it.
But it’s still a challenge for models, because drawing pelicans is difficult, drawing bicycles is difficult, and pelicans can’t ride bicycles in the first place.
Here’s the state of the art for November. Claude still couldn’t really draw a bicycle! The GPT-5.1 bicycle frame is pretty crap too.
Also in November, we had the first commit to an obscure GitHub repository called “Warelay”. We’ll come back to this repository shortly.
And then there were the December holidays, and individual developers took some time off and many started tinkering with these new coding agent model combinations... and it began to dawn on us quite how much they could do that they couldn’t do before.
Come January, a lot of us were quite excited to start putting this stuff into action.
Every year I set myself a New Year’s resolution, and for as long as I can remember it’s been the same thing: stay focused. Take on less new projects. Try to get things done in the projects I already have.
This year I decided that since that had never worked before, I’d go the other way.
We’ve got coding agents now, let’s see what they can do. I’m going to take on as many new projects as I like!
(You can ask me at the end of the year if this turned out to be a good idea or not. I have a lot of plates spinning right now.)
“Be more ambitious” has been something of a theme for the year, because the only way to find the limits of this technology is to keep on pushing them until they don’t work.
I also went on the Oxide and friends podcast with Bryan Cantrill and Adam Leventhal to share predictions for the next year (and three and six years).
With hindsight, my LLM predictions were pretty unambitious.
I said “it will become undeniable that LLMs write good code”—I think we’re there now.
I predicted we would finally solve sandboxing. I counted and around 40 of the 277 sessions at this conference touched on sandboxing or agent security in some way, so we’re at least putting a lot of effort into that!
I predicted “a Challenger disaster” for coding agent security. There’s certainly been a whole lot of noise around agent security this year, though the exact disaster I predicted (with coding agents being hijacked and causing real-world economic damage) hasn’t really played out.
We threw in a joke prediction that the Pope would weigh in on the economic impact of LLMs.
I also predicted that New Zealand’s Kākāpō parrots would have an outstanding breeding season this year.
These are flightless nocturnal parrots. They’re kind of dumpy looking, I think they’re beautiful, and there were only 236 of these parrots in the world at the start of the year.
Kākāpō only breed when the Rimu trees have a big fruiting season, and that hasn’t happened in four years... but this year the Rimu fruit were looking excellent.
Also on that podcast, we coined a term (full credit to Adam) for “that feeling of AI induced ennui where software engineers get listless because the AI can do anything”.
This has been a major theme throughout the year, and was touched on by several speakers at this conference.
As a software engineer, I’ve never had a year of my career where everything has changed so quickly and so dramatically.
A lot of what I’ve been doing this year is trying to come to terms with that and what that means for my own profession.
Also in January, I suffered from what I’m calling AI mania .
This is not the same thing as AI psychosis .
With AI mania, any time your agent isn’t building something for you feels like wasted time. You’re losing sleep because you could be staying up later getting your agents to do stuff.
My AI mania presented itself in some ridiculously over-ambitious projects.
I built a JavaScript interpreter entirely in Python , vibe-ported from MicroQuickJS by Fabrice Bellard.
Then I built a WebAssembly runtime in Python as well .
These projects were quite useful, in that they sort of cured me of my AI mania... because after I built these things, I got to look at them and ask “does the world need a slow, buggy, half-baked Python JavaScript interpreter?”
I did get this out of it: https://simonw.github.io/micro-javascript/playground.html
This page runs my JavaScript interpreter built in Python, running in Python using Pyodide , which is Python compiled to WebAssembly, running in JavaScript, running in a browser.
It’s a beautiful stack of horrors. I’ve been having a lot of fun with WebAssembly this year.
By the end of January, that repository we saw start in November had renamed itself, first to CLAWDIS, then CLAWDBOT, then Moltbot, and finally to OpenClaw.
At this point OpenClaw had 8,300 commits, less than two months after the project had started. I looked today and it’s over 100,000 commits now!
This is the most vibe-coded piece of software in existence.
(Here’s how I generated that list of name changes .)
This kicked off the OpenClaw revolution. It effectively defined a new category of software.
There’s a generic term for this which I really enjoy. We call software like this a “Claw”. There’s OpenClaw, NanoClaw , IronClaw , PicoClaw ...
Today they’re being rebranded as “personal agents” or “general agents”, but I still like to think of them as Claws.
The Apple stores in the Bay Area sold out of Mac Minis because so many people were buying Mac Minis to run OpenClaw!
Drew Breunig said that this is because your OpenClaw is a digital pet, and you buy a Mac mini as an aquarium to keep your claw in, which is kind of delightful.
Also in January, we had this website.
This was MoltBook , a social network for AI agents, where the idea was that you send your Claw to go and talk to all of the other Claws, because what could possibly go wrong if you did that?
The website launched on Thursday. It blew up on Friday . It was profiled by the New York Times on Monday . And by Tuesday, everyone had forgotten it existed as it drowned in a deluge of slop and spam.
Facebook/Meta bought it a month later .
In February, a company called StrongDM described what they called their Software Factory.
They wrote about this in Software Factories and the Agentic Moment . I posted my own notes at the time, having seen their demo in person back in October.
Dan Shapiro called this approach the Dark Factory , after the idea that if your factory is sufficiently automated you can turn the lights out, because you don’t even need to see what’s going on.
StrongDM presented two rules for software development that they’d been following since July last year.
The first was code must not be written by humans .
Any code that you write has to have been routed through a coding agent.
This sounded radical in February, but I imagine there are a lot of people in this room who are pretty much living that today.
Rule number two was code must not be reviewed by humans .
You’re not allowed to read the code!
This continued to be a huge topic for much of this year. Many of the sessions at this event have been about code review and how you can get away with this.
What I found interesting about StrongDM is that they were living six months ahead of the rest of us, and they’d been exploring what it means to build software, not read the code, but still be confident that the software is of high quality. What can you do with these agents to help verify their work?
StrongDM are a security company, and they had people with decades of experience on this project. They were very much exploring the edges of what’s possible and responsible to do with this stuff.
Also in February: First kākāpō chick in four years hatches on Valentine’s Day . Breeding season is off to a good start!
Also in February... Google released Gemini 3.1 Pro . That’s a pretty great pelican riding a bicycle! It’s got the chain in the right place, it’s got feet on both sides. There’s a little fish in the basket.
And then Google’s Jeff Dean tweeted a video comparing Gemini 3 Pro and Gemini 3.1 Pro that featured an animated pelican riding a bicycle, a frog on a penny-farthing, a giraffe driving a tiny car, an ostrich on roller skates, a turtle kickflipping a skateboard, and a dachshund driving a stretch limousine.
This was frustrating, because my protection for the pelican riding the bicycle test was always “if they draw a perfect pelican on a bicycle, I’ll ask for some other animal on something else.”
Google trained for all forms of animals on all forms of transport! They’ve defeated my benchmark at this point.
The other thing that started in February was Tokenmaxxing . We had headlines about Meta making AI adoption a formal part of performance reviews, and Microsoft wanting every employee to use AI, and Uber boasting that 90% of their engineers were using AI workflows.
Then a few months later we have Meta cracking down on token use, Microsoft saying tokenmaxxing is “not what we are optimizing for”, and Uber capping employee AI spending.
So tokenmaxxing went straight up and then straight back down again—because it turns out the agents are expensive .
Last year it was difficult to spend more than $50 on AI tokens, because we didn’t have anything interesting to do with them. Then agents blew up, and now you can actually spend $1,000 in a day doing real work.
This is also the reason that Anthropic’s valuation skyrocketed to maybe a trillion dollars.
AI appears to have hit product market fit in 2026, primarily through coding agents.
In March, we hit peak OpenClaw.
These photographs are from China, where companies hosted OpenClaw install parties which saw non-tech-nerds queueing up around the block for help getting Claws installed on their personal devices.
I think this proved real market demand for this class of Claws, or personal AI agents. It turns out regular people really do want a weird little AI agent that can do useful things on their behalf.
A Claw is really just a coding agent wearing a less threatening hat. Under the hood they work much the same way—writing and then executing code on your computer to get stuff done.
The race was on to be the first to build a safe Claw —a Claw you could give to regular human beings where they wouldn’t instantly shoot themselves in the foot.
Meta’s Muse came out three weeks ago and is currently at the top of the free charts on the iPhone App Store. It appears to be taking off with consumers.
I’m not yet convinced you can’t shoot yourself in the foot with Muse, but I guess we’ll find out for sure pretty soon.
Photos from How the OpenClaw Frenzy Is Testing China’s AI Commitment (March 29th) and The Enthusiasm and Anxiety Behind China’s OpenClaw Craze (April 8th, 2026).
In April, we had a model release where the mode

[truncated]
