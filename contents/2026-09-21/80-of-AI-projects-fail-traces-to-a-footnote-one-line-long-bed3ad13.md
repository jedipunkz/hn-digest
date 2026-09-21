---
source: "https://jamiewatters.work/journey/80-percent-ai-projects-fail"
hn_url: "https://news.ycombinator.com/item?id=49792600"
title: "\"80% of AI projects fail\" traces to a footnote one line long"
article_title: "\"80% of AI projects fail\" traces to a footnote one line long | Jamie Watters"
image: "https://pub-4f2aa5e351b44f67ba6dd0bc32fe1bb2.r2.dev/blog/80-percent-social.png"
author: "TheWayWithin"
captured_at: "2026-09-21T20:49:17Z"
capture_tool: "hn-digest"
hn_id: 49792600
score: 2
comments: 0
posted_at: "2026-09-21T20:02:47Z"
tags:
  - hacker-news
---

# "80% of AI projects fail" traces to a footnote one line long

- HN: [49792600](https://news.ycombinator.com/item?id=49792600)
- Source: [jamiewatters.work](https://jamiewatters.work/journey/80-percent-ai-projects-fail)
- Score: 2
- Comments: 0
- Posted: 2026-09-21T20:02:47Z

## Translation

Title: "80% of AI projects fail" traces to a footnote one line long
Article title: "80% of AI projects fail" traces to a footnote one line long | Jamie Watters
Description: The most repeated number in enterprise AI turns out to be a paraphrase of unnamed opinion surveys, quoted in a magazine profile of a vendor.

Article text:
"80% of AI projects fail" traces to a footnote one line long | Jamie Watters Skip to main content JW Jamie Watters Home Research Portfolio Proof The Journey Videos Books About Contact "80% of AI projects fail" traces to a footnote one line long
Operational resilience and AI delivery practitioner . Technology since 1985 .
You have seen the number. Eighty per cent of AI projects fail, roughly twice the rate of conventional IT projects. It is in the pitch decks, the board papers, the conference keynotes and the LinkedIn posts. I have seen it used to justify buying software and to justify not buying software, which should have been the first clue.
I went looking for the study behind it. There isn't one.
The citation almost everyone reaches for is RAND. Their 2024 report, The Root Causes of Failure for Artificial Intelligence Projects and How They Can Succeed , says it twice: "By some estimates, more than 80 percent of AI projects fail."
Read that sentence again, because RAND wrote it carefully. By some estimates. That hedge is theirs, and it is doing a lot of work.
The footnote attached to it is complete. It reads, in full:
Kahn, "Want Your Company's AI Project to Succeed?"
That is Jeremy Kahn, writing in Fortune on 26 July 2022. Not a study. A magazine article, and specifically a profile of the chief executive of a company that sells AI software.
So I read Kahn. Here is the sentence the whole edifice rests on:
That's borne out in a slew of recent surveys, where business leaders have put the failure rate of A.I. projects at between 83% and 92%.
A slew of recent surveys. Unnamed. And look at what they measured: not projects, but business leaders' opinions about projects. Not a rate, but a range, 83 to 92 per cent.
That is the bottom of the chain. There is nothing underneath it.
Now watch the number climb back.
Unnamed surveys of executive opinion become "between 83% and 92%" in a vendor profile. RAND rounds that down to "more than 80 percent" and attaches a hedge. And then every downstream citation strips the hedge and reattributes the whole thing, so that what circulates is "RAND found that 80% of AI projects fail".
RAND found no such thing. RAND generated no failure-rate estimate at all, and their study design could not have produced one. It is an exploratory interview study about causes, built on 65 interviews, in which failure is defined perceptually as "a project that was perceived to be a failure by the organization". There is no denominator of projects anywhere in it. You cannot compute a rate without a denominator, and they never claimed to.
RAND's handling of this figure is entirely defensible. What happened to it afterwards is not. The error is not in the research. It is in the transmission, and every time I traced one of these, it had moved in the direction that made the sentence tidier.
"Twice the rate of IT projects" has its own provenance, and it is thinner still.
It traces to Iavor Bojinov in Harvard Business Review , November-December 2023:
Some estimates place the failure rate as high as 80% — almost double the rate of corporate IT project failures from a decade ago.
No source is given for either number. Not for the 80, not for the double.
And notice the clause that never survives the retelling: from a decade ago . Bojinov is comparing a 2023 AI figure against IT failures from around 2013. A ten-year gap, stated plainly by the author, and it vanishes in every restatement I found. The comparison you have heard is between a number with no source and a different number with no source, measured a decade apart.
There is no dataset. There is no consistent definition. There are no two comparable samples. The numerator and the denominator were never measured by the same instrument, or, as far as I can establish, by any instrument at all.
The most repeated empirical claim in this field is not an empirical claim.
What this does and does not show
I want to be careful here, because the tempting conclusion is the wrong one.
This does not prove that AI projects succeed. It does not prove the real rate is low. I have no idea what the real rate is, and after several weeks inside this literature I am fairly confident nobody else does either.
What it proves is narrower and, I think, more useful: the number everyone repeats is not evidence about anything, and nobody checked it for four years.
That last part is the bit I keep returning to. This was not hard. It took me an afternoon, a copy of the RAND PDF and the ability to read a footnote. Anyone quoting the figure could have done the same at any point since 2022. The chain was never hidden. It was simply never followed.
The bit that made me laugh, then stop laughing
While researching this, I used several AI-generated syntheses of the failure literature as a prior-art survey, to see what the consensus looked like before I went at the primary sources.
Two of them independently cited the same study: a RAND review of "over 2,400 enterprise initiatives", complete with a percentage breakdown of causes.
RAND conducted 65 interviews. It reviewed no enterprise initiatives. The study does not exist.
Two separate models produced a citable-sounding artefact with the same fictitious sample size, which is a fairly good illustration of how a number enters circulation and then gets repeated by people who assume somebody upstream checked. The machines are now doing to this literature exactly what the literature has been doing to itself.
I am not going to tell you to be sceptical, which is advice nobody has ever acted on. Here is something narrower you can actually do.
The next time you meet a headline failure statistic, before you repeat it, ask four things. What does it count: projects, companies, use cases, or opinions? On what unit, and is that the same unit as the thing you are comparing it to? Over what window? And is it a measurement of something that happened, or a forecast of something that might?
They take about a minute. The 80% does not survive the first one: it counts opinions.
This is the first of several pieces drawn from a working paper I published this week, which traces nineteen headline AI failure statistics to their primary sources and classifies each one. Seven of the nineteen turn out not to be measurements of anything that happened. The full paper, with every source stated and every claim checked, is here .
This article is part of a series off the paper Why AI projects fail: 19 headline statistics, and what each one actually counts . The rest of the series, in reading order:
Zillow wrote off $407.9m. The fix everyone recommends would have changed nothing.
AI project failure rates: four questions that tell you whether the number means anything
Epic's sepsis model missed two thirds of cases. Its accuracy claim was never published.
Checking the source is not enough: I traced a number to a scientific paper and found a Teams meeting
Amazon's AI did exactly what it was told. Deleting words for two years did not help.
Anthropic's safety policy has never stopped a release. Yours probably hasn't either.
AI vs IT project failure rate: the comparison has never been made
Jamie Watters has been programming since 1985, starting in mainframe assembler. He wrote the best-selling book on business continuity in the world for its first few years, teaches The Headless Way , and now builds AI products solo with the code open on GitHub .
No paywall, no sponsors. If this saved you some time, you can buy me a coffee.
Get the next one in your inbox
I build with AI in the open and write up what held and what didn't. Real numbers, the failures before the wins.
Share on Twitter Share on LinkedIn Copy Link ← Previous Post Why AI projects fail: 19 headline statistics, and what each one actually counts Nineteen headline AI failure statistics, traced to source. Seven are not measurements of anything that happened.
The story is that an algorithm mispriced houses and a company lost half a billion dollars. A federal court order describes something else: a deliberate plan to bid above what the algorithm said.
New posts and what actually held up, straight to your inbox. No hype.
Open code, real numbers, the failures before the wins.
How this site works: I write every post myself, with AI assisting the research and the drafting, and writing most of the code. Where a post reports on something, I check the reported version against the primary sources, and the gap between the two is usually the point of the piece. No sponsors, no paid placements, and no affiliate links unless a post says so in plain sight. The product code is open on GitHub .
© 2026 Jamie Watters. All rights reserved.

## Original Extract

The most repeated number in enterprise AI turns out to be a paraphrase of unnamed opinion surveys, quoted in a magazine profile of a vendor.

"80% of AI projects fail" traces to a footnote one line long | Jamie Watters Skip to main content JW Jamie Watters Home Research Portfolio Proof The Journey Videos Books About Contact "80% of AI projects fail" traces to a footnote one line long
Operational resilience and AI delivery practitioner . Technology since 1985 .
You have seen the number. Eighty per cent of AI projects fail, roughly twice the rate of conventional IT projects. It is in the pitch decks, the board papers, the conference keynotes and the LinkedIn posts. I have seen it used to justify buying software and to justify not buying software, which should have been the first clue.
I went looking for the study behind it. There isn't one.
The citation almost everyone reaches for is RAND. Their 2024 report, The Root Causes of Failure for Artificial Intelligence Projects and How They Can Succeed , says it twice: "By some estimates, more than 80 percent of AI projects fail."
Read that sentence again, because RAND wrote it carefully. By some estimates. That hedge is theirs, and it is doing a lot of work.
The footnote attached to it is complete. It reads, in full:
Kahn, "Want Your Company's AI Project to Succeed?"
That is Jeremy Kahn, writing in Fortune on 26 July 2022. Not a study. A magazine article, and specifically a profile of the chief executive of a company that sells AI software.
So I read Kahn. Here is the sentence the whole edifice rests on:
That's borne out in a slew of recent surveys, where business leaders have put the failure rate of A.I. projects at between 83% and 92%.
A slew of recent surveys. Unnamed. And look at what they measured: not projects, but business leaders' opinions about projects. Not a rate, but a range, 83 to 92 per cent.
That is the bottom of the chain. There is nothing underneath it.
Now watch the number climb back.
Unnamed surveys of executive opinion become "between 83% and 92%" in a vendor profile. RAND rounds that down to "more than 80 percent" and attaches a hedge. And then every downstream citation strips the hedge and reattributes the whole thing, so that what circulates is "RAND found that 80% of AI projects fail".
RAND found no such thing. RAND generated no failure-rate estimate at all, and their study design could not have produced one. It is an exploratory interview study about causes, built on 65 interviews, in which failure is defined perceptually as "a project that was perceived to be a failure by the organization". There is no denominator of projects anywhere in it. You cannot compute a rate without a denominator, and they never claimed to.
RAND's handling of this figure is entirely defensible. What happened to it afterwards is not. The error is not in the research. It is in the transmission, and every time I traced one of these, it had moved in the direction that made the sentence tidier.
"Twice the rate of IT projects" has its own provenance, and it is thinner still.
It traces to Iavor Bojinov in Harvard Business Review , November-December 2023:
Some estimates place the failure rate as high as 80% — almost double the rate of corporate IT project failures from a decade ago.
No source is given for either number. Not for the 80, not for the double.
And notice the clause that never survives the retelling: from a decade ago . Bojinov is comparing a 2023 AI figure against IT failures from around 2013. A ten-year gap, stated plainly by the author, and it vanishes in every restatement I found. The comparison you have heard is between a number with no source and a different number with no source, measured a decade apart.
There is no dataset. There is no consistent definition. There are no two comparable samples. The numerator and the denominator were never measured by the same instrument, or, as far as I can establish, by any instrument at all.
The most repeated empirical claim in this field is not an empirical claim.
What this does and does not show
I want to be careful here, because the tempting conclusion is the wrong one.
This does not prove that AI projects succeed. It does not prove the real rate is low. I have no idea what the real rate is, and after several weeks inside this literature I am fairly confident nobody else does either.
What it proves is narrower and, I think, more useful: the number everyone repeats is not evidence about anything, and nobody checked it for four years.
That last part is the bit I keep returning to. This was not hard. It took me an afternoon, a copy of the RAND PDF and the ability to read a footnote. Anyone quoting the figure could have done the same at any point since 2022. The chain was never hidden. It was simply never followed.
The bit that made me laugh, then stop laughing
While researching this, I used several AI-generated syntheses of the failure literature as a prior-art survey, to see what the consensus looked like before I went at the primary sources.
Two of them independently cited the same study: a RAND review of "over 2,400 enterprise initiatives", complete with a percentage breakdown of causes.
RAND conducted 65 interviews. It reviewed no enterprise initiatives. The study does not exist.
Two separate models produced a citable-sounding artefact with the same fictitious sample size, which is a fairly good illustration of how a number enters circulation and then gets repeated by people who assume somebody upstream checked. The machines are now doing to this literature exactly what the literature has been doing to itself.
I am not going to tell you to be sceptical, which is advice nobody has ever acted on. Here is something narrower you can actually do.
The next time you meet a headline failure statistic, before you repeat it, ask four things. What does it count: projects, companies, use cases, or opinions? On what unit, and is that the same unit as the thing you are comparing it to? Over what window? And is it a measurement of something that happened, or a forecast of something that might?
They take about a minute. The 80% does not survive the first one: it counts opinions.
This is the first of several pieces drawn from a working paper I published this week, which traces nineteen headline AI failure statistics to their primary sources and classifies each one. Seven of the nineteen turn out not to be measurements of anything that happened. The full paper, with every source stated and every claim checked, is here .
This article is part of a series off the paper Why AI projects fail: 19 headline statistics, and what each one actually counts . The rest of the series, in reading order:
Zillow wrote off $407.9m. The fix everyone recommends would have changed nothing.
AI project failure rates: four questions that tell you whether the number means anything
Epic's sepsis model missed two thirds of cases. Its accuracy claim was never published.
Checking the source is not enough: I traced a number to a scientific paper and found a Teams meeting
Amazon's AI did exactly what it was told. Deleting words for two years did not help.
Anthropic's safety policy has never stopped a release. Yours probably hasn't either.
AI vs IT project failure rate: the comparison has never been made
Jamie Watters has been programming since 1985, starting in mainframe assembler. He wrote the best-selling book on business continuity in the world for its first few years, teaches The Headless Way , and now builds AI products solo with the code open on GitHub .
No paywall, no sponsors. If this saved you some time, you can buy me a coffee.
Get the next one in your inbox
I build with AI in the open and write up what held and what didn't. Real numbers, the failures before the wins.
Share on Twitter Share on LinkedIn Copy Link ← Previous Post Why AI projects fail: 19 headline statistics, and what each one actually counts Nineteen headline AI failure statistics, traced to source. Seven are not measurements of anything that happened.
The story is that an algorithm mispriced houses and a company lost half a billion dollars. A federal court order describes something else: a deliberate plan to bid above what the algorithm said.
New posts and what actually held up, straight to your inbox. No hype.
Open code, real numbers, the failures before the wins.
How this site works: I write every post myself, with AI assisting the research and the drafting, and writing most of the code. Where a post reports on something, I check the reported version against the primary sources, and the gap between the two is usually the point of the piece. No sponsors, no paid placements, and no affiliate links unless a post says so in plain sight. The product code is open on GitHub .
© 2026 Jamie Watters. All rights reserved.
