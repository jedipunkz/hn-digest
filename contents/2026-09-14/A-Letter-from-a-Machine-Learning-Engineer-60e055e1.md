---
source: "https://nemin.hu/llm-letter/index.html"
hn_url: "https://news.ycombinator.com/item?id=49704586"
title: "A Letter from a Machine Learning Engineer"
article_title: "A Letter from a Machine Learning Engineer - Nemin's Blog"
image: "https://nemin.hu/logo.png"
author: "sampsn"
captured_at: "2026-09-14T22:17:01Z"
capture_tool: "hn-digest"
hn_id: 49704586
score: 2
comments: 0
posted_at: "2026-09-14T21:47:17Z"
tags:
  - hacker-news
---

# A Letter from a Machine Learning Engineer

- HN: [49704586](https://news.ycombinator.com/item?id=49704586)
- Source: [nemin.hu](https://nemin.hu/llm-letter/index.html)
- Score: 2
- Comments: 0
- Posted: 2026-09-14T21:47:17Z

## Translation

Title: A Letter from a Machine Learning Engineer
Article title: A Letter from a Machine Learning Engineer - Nemin's Blog
Description: My thoughts about an AI-related letter I received recently

Article text:
A Letter from a Machine Learning Engineer
My thoughts about an AI-related letter I received recently
My thoughts on the individual
sections
A couple of days ago, I received an email from a throwaway address,
whose author claims to be an ML Engineer at one of the frontier labs. In
this, they explain their beliefs about the future of the current LLM
craze and how frontier labs are bound to eventually go under.
I have no means of ascertaining whether the author was really who
they say they are, nor was I able to reply to them, because the address
had since been deleted. 1
I'm not one to air out my private mails normally and I deliberated
for a few days whether to make an exception here. However, considering
the importance of the topic in the current times and the fact that the
author wrote it in a way that (per their own words) cannot be used to
identify them, I figured it'd be worth sharing with the greater
community. 2
I will first quote the mail verbatim for those who want to read it as
a coherent essay. Then I'll repeat it while interrupting the text
whenever I have something to say.
Both sections will contain the exact same text, so if you're fine
with reading the cut up text with my annotations, feel free to skip to
the second part.
I work at a frontier AI lab as an MLE. I help train models and do
RLHF for frontier stuff at one of the big companies you probably
know.
Sorry this email is short. I have an NDA, so I need to keep this
general.
I saw your post about worrying about the future of programming.
Honestly, I'm more worried about losing my job than yours. How is that
even possible? Internally, we've hit limits on scaling models for
programming tasks, which is why so many people are switching careers. At
my company, a lot of executives have quit in a year, and there's real
stress about the IPO. There is lots more I cannot share, but it does not
look good internally.
The real threat to labs like ours is that there is no competitive
moat. We have to keep advertising our math-solving abilities and
overselling because things are changing fast. The proprietary diffusion
models got largely replaced by open-weights ones. Image and video
generation can now be done with open models without needing us. Sales
tanked and that's why the frontier labs quietly stopped promoting
image/video gen.
LLMs are next. Even parts of the serious research community like
LeCun have moved on from them. Most of the people who wrote the
attention paper have left Google and scattered across new labs and
startups. If you look at efficiency and size, the trend is clear. We'll
see an open-weights model as good or better than Fable 5.1 within a
year. There is research going on into parameter efficiency, so if it
pans out, something like this could one day run on devices like high-end
laptops.
AGI is not mathematically possible with LLMs. There was a paper a
couple of years ago arguing that, using Cantor's diagonalization
argument. I can't remember the name for the life of me as I read it a
long time ago. This realization changes everything. And a Google
researcher leaked an internal memo in 2023 that said the quiet part out
loud: "We have no moat, and neither does OpenAI." feel free to read if
you're interested.
The online discussion is poisoned. A lot of the voices driving the
hype, the influencers who sold those online courses, are now just being
sponsored by the AI companies themselves. This is why programmers are so
stressed, I think. Don't trust anything you hear or read on mainstream
sites. A few engineers have been deploying agents as "fun projects" to
market models and scrape data (that's how I found your email), and you
should absolutely ignore the forecasts from the CEOs. They are
completely clueless about what's actually happening.
2 My thoughts on the individual
sections
I work at a frontier AI lab as an MLE. I help train models and do
RLHF for frontier stuff at one of the big companies you probably
know.
MLE is Machine Learning Engineer, this I figured out on my own,
because it was pretty obvious. However, I had no idea what RLHF is (only
heard of RTFM and GLHF before :), so I looked it up. It means Reinforcement
Learning from Human Feedback , i.e. a particular variant of agentic
training, where the reward function (what determines how "good" an AI's
responses are) is first manually tuned by human operators.
Sorry this email is short. I have an NDA, so I need to keep this
general.
Hard to fault the author (who I'll refer to as "G" 3
from now on, due to the fake moniker they used in the mail) for this.
While NDA -s
are not
universally enforceable , they are still a sword of Damocles above
one's head of "Talk too much and you'll never work in this field
again".
I saw your post about worrying about the future of programming.
By "my post", G means my commitment
to not using LLMs for hobby projects, due to their ethical issues, the
perceived uselessness of AI, and the love of the craft.
Since writing this declaration my opinion on LLMs has slightly
shifted or, rather, I'd say became more nuanced:
For one, I can no longer say that the output of the
various coding agents is useless and I think anyone who keeps insisting
that they are still outputting pure gibberish has not tried an agent
recently.
I still do not trust vibe-coding, as I believe that unsupervised
development results in write-only apps, that cost exponentially more to
maintain over time and are impossible to comprehend without the constant
assistance of the agent.
This, in my opinion, is tying yourself to a kind of "supply chain" of
its own, where the token prices are in constant flux and aren't based on
any real, measurable metric. It also does nothing to improve your own
understanding and mastery over computers. I'm sure, regardless of this,
you can "get rich quick" by vibing up something, but I'm not very
interested in that.
However, after being strongly recommended at my workplace to try
using AI, I found that agents are great at being a sort of "secretary",
who can fetch you information about the codebase based on free-form text
queries. While grep , find , etc. can get you
far, "vibe searching" takes things a step further in a way, I couldn't
really replicate with conventional tools.
I also found, that using an agent as a pre-reviewer before you send
in your PR for a human to actually take a look at is a nice way of
avoiding some unnecessary back and forth. While I still ultimately defer
to human judgement (including my own), it is nice to catch issues such
as "hey, this line is potentially dereferencing null", "this comment
wasn't updated to what the code actually does", and similar
nitpicks.
For two, and this might sound a little paradoxical,
after I just conceded agentic output being better than I once thought,
but I no longer expect programming as a job to disappear. In fact, I
think people who understand how these mechanical parrots "think" have
more job security than ever.
Back when writing my first post, I was genuinely somewhat worried,
that after almost five years of university (and before that more than a
decade of hobby programming) will just become obsolete once
models become good enough to fully replace us.
However, while I've seen a lot of improvement to the quality of
agentic output, I still experience a ton of flukes as well. My company
allows us to use top of the line models and they are both incredibly
capable at times and terribly, horribly dumb.
Among other things: The model we use still occasionally hallucinates
functions that simply don't exist, despite being plugged into LSP. It
occasionally moves code into conditionals, that should happen in both
branches. And sometimes it just starts using legacy code that's clearly
marked as such as a style to follow. And that's just a small sample.
You simply need a human in the loop to make sure the machine doesn't
do anything stupid. And, as mentioned before, I'm absolutely not the
sort to let the agent loose. I rarely allow it to generate anything more
than 10-20 lines. How could it then replace me as a whole?
Instead, I believe the recent lay-off wave, that was conducted "due
to AI performance optimization" was really just CEO-s buying into hype
and the market experiencing a painful normalization in the post-COVID /
ZIRP
era.
That being said, as convenient as LLMs are, the
moral concerns and the fact, that I want to keep my skills sharp still
makes me want to relegate AI usage to my job only.
Honestly, I'm more worried about losing my job than yours. How is
that even possible? Internally, we've hit limits on scaling models for
programming tasks, which is why so many people are switching careers. At
my company, a lot of executives have quit in a year, and there's real
stress about the IPO. There is lots more I cannot share, but it does not
look good internally.
Funnily enough, I assumed both Anthropic and OpenAI have already
IPO-d, but it turns out neither of them have. I guess the numbers don't look good enough to bring
this stuff onto the stock market.
As for the rest of the claims, I have no real means of checking the
validity of them. There has been that story about a researcher quitting ,
because they fear AI will endanger humanity, but this person wasn't an
executive and they have seemingly left for ethical reasons, not
financial ones.
The real threat to labs like ours is that there is no competitive
moat. We have to keep advertising our math-solving abilities and
overselling because things are changing fast. The proprietary diffusion
models got largely replaced by open-weights ones. Image and video
generation can now be done with open models without needing us. Sales
tanked and that's why the frontier labs quietly stopped promoting
image/video gen.
This part feels particularly striking after the recent controversy
related to progress on the Navier-Stokes equations, which OpenAI
potentially snatched from the hands of two researchers.
For clarity, G's mail arrived OpenAI published its findings, so this
was a reaction, not a forewarning to it and it is no evidence that G may
be from OpenAI.
As for the other part of this quote, it really is interesting how
much AI slop videos became a cheap commodity, that is simply
not worth selling as a singular product. Back in the day we had things
like Sora and then it went belly up, when the math wasn't mathing about
its economics. I wonder if one day selling tokens will be seen as a
pointless venture, because they'll be so commodified.
LLMs are next. Even parts of the serious research community like
LeCun have moved on from them. Most of the people who wrote the
attention paper have left Google and scattered across new labs and
startups. If you look at efficiency and size, the trend is clear. We'll
see an open-weights model as good or better than Fable 5.1 within a
year. There is research going on into parameter efficiency, so if it
pans out, something like this could one day run on devices like high-end
laptops.
G is talking about Yann LeCun , known
for his work on ML and computer vision, and who in April
of this year engaged in a lecture about LLMs being a dead-end and
that a new approach would be necessary, if humanity really wanted to go
for AGI.
Despite this, he is also involved in a
"collaborative foundation for open and sovereign AI," so take that as
you will.
AGI is not mathematically possible with LLMs. There was a paper a
couple of years ago arguing that, using Cantor's diagonalization
argument. I can't remember the name for the life of me as I read it a
long time ago. This realization changes everything. And a Google
researcher leaked an internal memo in 2023 that said the quiet part out
loud: "We have no moat, and neither does OpenAI." feel free to read if
you're interested.
I suspect G is thinking of Hallucination is
Inevitable , a frequently-cited paper from 2024, which proved
that it is impossible to make an LLM, that doesn't hallucinate (i.e. act
as a perfect general problem solver).
As for the other

[truncated]

## Original Extract

My thoughts about an AI-related letter I received recently

A Letter from a Machine Learning Engineer
My thoughts about an AI-related letter I received recently
My thoughts on the individual
sections
A couple of days ago, I received an email from a throwaway address,
whose author claims to be an ML Engineer at one of the frontier labs. In
this, they explain their beliefs about the future of the current LLM
craze and how frontier labs are bound to eventually go under.
I have no means of ascertaining whether the author was really who
they say they are, nor was I able to reply to them, because the address
had since been deleted. 1
I'm not one to air out my private mails normally and I deliberated
for a few days whether to make an exception here. However, considering
the importance of the topic in the current times and the fact that the
author wrote it in a way that (per their own words) cannot be used to
identify them, I figured it'd be worth sharing with the greater
community. 2
I will first quote the mail verbatim for those who want to read it as
a coherent essay. Then I'll repeat it while interrupting the text
whenever I have something to say.
Both sections will contain the exact same text, so if you're fine
with reading the cut up text with my annotations, feel free to skip to
the second part.
I work at a frontier AI lab as an MLE. I help train models and do
RLHF for frontier stuff at one of the big companies you probably
know.
Sorry this email is short. I have an NDA, so I need to keep this
general.
I saw your post about worrying about the future of programming.
Honestly, I'm more worried about losing my job than yours. How is that
even possible? Internally, we've hit limits on scaling models for
programming tasks, which is why so many people are switching careers. At
my company, a lot of executives have quit in a year, and there's real
stress about the IPO. There is lots more I cannot share, but it does not
look good internally.
The real threat to labs like ours is that there is no competitive
moat. We have to keep advertising our math-solving abilities and
overselling because things are changing fast. The proprietary diffusion
models got largely replaced by open-weights ones. Image and video
generation can now be done with open models without needing us. Sales
tanked and that's why the frontier labs quietly stopped promoting
image/video gen.
LLMs are next. Even parts of the serious research community like
LeCun have moved on from them. Most of the people who wrote the
attention paper have left Google and scattered across new labs and
startups. If you look at efficiency and size, the trend is clear. We'll
see an open-weights model as good or better than Fable 5.1 within a
year. There is research going on into parameter efficiency, so if it
pans out, something like this could one day run on devices like high-end
laptops.
AGI is not mathematically possible with LLMs. There was a paper a
couple of years ago arguing that, using Cantor's diagonalization
argument. I can't remember the name for the life of me as I read it a
long time ago. This realization changes everything. And a Google
researcher leaked an internal memo in 2023 that said the quiet part out
loud: "We have no moat, and neither does OpenAI." feel free to read if
you're interested.
The online discussion is poisoned. A lot of the voices driving the
hype, the influencers who sold those online courses, are now just being
sponsored by the AI companies themselves. This is why programmers are so
stressed, I think. Don't trust anything you hear or read on mainstream
sites. A few engineers have been deploying agents as "fun projects" to
market models and scrape data (that's how I found your email), and you
should absolutely ignore the forecasts from the CEOs. They are
completely clueless about what's actually happening.
2 My thoughts on the individual
sections
I work at a frontier AI lab as an MLE. I help train models and do
RLHF for frontier stuff at one of the big companies you probably
know.
MLE is Machine Learning Engineer, this I figured out on my own,
because it was pretty obvious. However, I had no idea what RLHF is (only
heard of RTFM and GLHF before :), so I looked it up. It means Reinforcement
Learning from Human Feedback , i.e. a particular variant of agentic
training, where the reward function (what determines how "good" an AI's
responses are) is first manually tuned by human operators.
Sorry this email is short. I have an NDA, so I need to keep this
general.
Hard to fault the author (who I'll refer to as "G" 3
from now on, due to the fake moniker they used in the mail) for this.
While NDA -s
are not
universally enforceable , they are still a sword of Damocles above
one's head of "Talk too much and you'll never work in this field
again".
I saw your post about worrying about the future of programming.
By "my post", G means my commitment
to not using LLMs for hobby projects, due to their ethical issues, the
perceived uselessness of AI, and the love of the craft.
Since writing this declaration my opinion on LLMs has slightly
shifted or, rather, I'd say became more nuanced:
For one, I can no longer say that the output of the
various coding agents is useless and I think anyone who keeps insisting
that they are still outputting pure gibberish has not tried an agent
recently.
I still do not trust vibe-coding, as I believe that unsupervised
development results in write-only apps, that cost exponentially more to
maintain over time and are impossible to comprehend without the constant
assistance of the agent.
This, in my opinion, is tying yourself to a kind of "supply chain" of
its own, where the token prices are in constant flux and aren't based on
any real, measurable metric. It also does nothing to improve your own
understanding and mastery over computers. I'm sure, regardless of this,
you can "get rich quick" by vibing up something, but I'm not very
interested in that.
However, after being strongly recommended at my workplace to try
using AI, I found that agents are great at being a sort of "secretary",
who can fetch you information about the codebase based on free-form text
queries. While grep , find , etc. can get you
far, "vibe searching" takes things a step further in a way, I couldn't
really replicate with conventional tools.
I also found, that using an agent as a pre-reviewer before you send
in your PR for a human to actually take a look at is a nice way of
avoiding some unnecessary back and forth. While I still ultimately defer
to human judgement (including my own), it is nice to catch issues such
as "hey, this line is potentially dereferencing null", "this comment
wasn't updated to what the code actually does", and similar
nitpicks.
For two, and this might sound a little paradoxical,
after I just conceded agentic output being better than I once thought,
but I no longer expect programming as a job to disappear. In fact, I
think people who understand how these mechanical parrots "think" have
more job security than ever.
Back when writing my first post, I was genuinely somewhat worried,
that after almost five years of university (and before that more than a
decade of hobby programming) will just become obsolete once
models become good enough to fully replace us.
However, while I've seen a lot of improvement to the quality of
agentic output, I still experience a ton of flukes as well. My company
allows us to use top of the line models and they are both incredibly
capable at times and terribly, horribly dumb.
Among other things: The model we use still occasionally hallucinates
functions that simply don't exist, despite being plugged into LSP. It
occasionally moves code into conditionals, that should happen in both
branches. And sometimes it just starts using legacy code that's clearly
marked as such as a style to follow. And that's just a small sample.
You simply need a human in the loop to make sure the machine doesn't
do anything stupid. And, as mentioned before, I'm absolutely not the
sort to let the agent loose. I rarely allow it to generate anything more
than 10-20 lines. How could it then replace me as a whole?
Instead, I believe the recent lay-off wave, that was conducted "due
to AI performance optimization" was really just CEO-s buying into hype
and the market experiencing a painful normalization in the post-COVID /
ZIRP
era.
That being said, as convenient as LLMs are, the
moral concerns and the fact, that I want to keep my skills sharp still
makes me want to relegate AI usage to my job only.
Honestly, I'm more worried about losing my job than yours. How is
that even possible? Internally, we've hit limits on scaling models for
programming tasks, which is why so many people are switching careers. At
my company, a lot of executives have quit in a year, and there's real
stress about the IPO. There is lots more I cannot share, but it does not
look good internally.
Funnily enough, I assumed both Anthropic and OpenAI have already
IPO-d, but it turns out neither of them have. I guess the numbers don't look good enough to bring
this stuff onto the stock market.
As for the rest of the claims, I have no real means of checking the
validity of them. There has been that story about a researcher quitting ,
because they fear AI will endanger humanity, but this person wasn't an
executive and they have seemingly left for ethical reasons, not
financial ones.
The real threat to labs like ours is that there is no competitive
moat. We have to keep advertising our math-solving abilities and
overselling because things are changing fast. The proprietary diffusion
models got largely replaced by open-weights ones. Image and video
generation can now be done with open models without needing us. Sales
tanked and that's why the frontier labs quietly stopped promoting
image/video gen.
This part feels particularly striking after the recent controversy
related to progress on the Navier-Stokes equations, which OpenAI
potentially snatched from the hands of two researchers.
For clarity, G's mail arrived OpenAI published its findings, so this
was a reaction, not a forewarning to it and it is no evidence that G may
be from OpenAI.
As for the other part of this quote, it really is interesting how
much AI slop videos became a cheap commodity, that is simply
not worth selling as a singular product. Back in the day we had things
like Sora and then it went belly up, when the math wasn't mathing about
its economics. I wonder if one day selling tokens will be seen as a
pointless venture, because they'll be so commodified.
LLMs are next. Even parts of the serious research community like
LeCun have moved on from them. Most of the people who wrote the
attention paper have left Google and scattered across new labs and
startups. If you look at efficiency and size, the trend is clear. We'll
see an open-weights model as good or better than Fable 5.1 within a
year. There is research going on into parameter efficiency, so if it
pans out, something like this could one day run on devices like high-end
laptops.
G is talking about Yann LeCun , known
for his work on ML and computer vision, and who in April
of this year engaged in a lecture about LLMs being a dead-end and
that a new approach would be necessary, if humanity really wanted to go
for AGI.
Despite this, he is also involved in a
"collaborative foundation for open and sovereign AI," so take that as
you will.
AGI is not mathematically possible with LLMs. There was a paper a
couple of years ago arguing that, using Cantor's diagonalization
argument. I can't remember the name for the life of me as I read it a
long time ago. This realization changes everything. And a Google
researcher leaked an internal memo in 2023 that said the quiet part out
loud: "We have no moat, and neither does OpenAI." feel free to read if
you're interested.
I suspect G is thinking of Hallucination is
Inevitable , a frequently-cited paper from 2024, which proved
that it is impossible to make an LLM, that doesn't hallucinate (i.e. act
as a perfect general problem solver).
As for the other

[truncated]
