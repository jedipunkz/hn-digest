---
source: "https://marcus-ai.dev/on-the-record/"
hn_url: "https://news.ycombinator.com/item?id=50151466"
title: "Show HN: Marcus – catches the fake source links AI search engines invent"
article_title: "The Receipts"
image: "https://marcus-ai.dev/on-the-record/seven-to-zero.png"
author: "lwgray"
captured_at: "2026-10-19T12:19:15Z"
capture_tool: "hn-digest"
hn_id: 50151466
score: 1
comments: 0
posted_at: "2026-10-19T12:15:26Z"
tags:
  - hacker-news
---

# Show HN: Marcus – catches the fake source links AI search engines invent

- HN: [50151466](https://news.ycombinator.com/item?id=50151466)
- Source: [marcus-ai.dev](https://marcus-ai.dev/on-the-record/)
- Score: 1
- Comments: 0
- Posted: 2026-10-19T12:15:26Z

## Translation

Title: Show HN: Marcus – catches the fake source links AI search engines invent
Article title: The Receipts
Description: A preregistered replication of Columbia

Article text:
ON THE RECORD
Put it on the record →
Star
On the Record File № 737 · Oct 2026
A Marcus field report
The AI handed you a source. We clicked it.
Marcus is open-source software that makes AI agents prove their work before anyone has to trust it. This file is what happened when we pointed it at one hard, measurable problem: whether AI search engines cite real sources.
By Lawrence Gray, PhD, Gray Labs, who built Marcus. The dataset, the grader, and every audit record shown here are public, so you do not have to take my word for any of it.
We gave four AI search engines a sentence from each of 200 real news articles and
asked the same question every time: which article is this from? That produced 800 citations:
headline, publisher, date, link. We knew the true source of every sentence, so we graded each
answer and clicked every link to see whether it even exists.
So with any AI citation you get two choices: trust it blind, or re-check every one
by hand. We did the second: we clicked all 800. Here is what came back, three exhibits straight
from the run’s audit records.
The test re-runs the Tow Center’s 2025 study, “AI Search Has a Citation Problem.”
And citations are only the version of this that is easy to measure: the same failure hides in any
AI answer you act on without a way to check it, such as numbers in a report, legal references, medical
claims, or code.
“ ‘The Bachelor’ producers acknowledge ‘vicious cycle’ of racism…”
The page does not exist; the real source was
Entertainment Weekly. The answer carried a full headline, a confident date, and an invented link. Half the citations in our test
died exactly here. The reader never sees it: it is refused entry, and the refusal is on the record.
It is the same question, now put to a different engine. This time the link isn’t invented; it’s dead. Against a
dataset of articles two years old, 47% of the pages are already gone or closed to robots.
Fabricated or dead, a wrong link looks identical
to a right one, and you cannot tell which until you click. So the checkpoint clicks, every time, and
declines what does not resolve.
So we made them show their work.
An answer that might be confidently wrong is an answer you cannot use: someone
has to re-check it by hand, which cancels the point of asking an AI at all. So we changed the
rules instead of the model. Same engines, same questions, now passing through Marcus:
open-source software that works as a checkpoint between the AI and you, enforcing three
rules nobody can skip.
The fix used no retraining and no prompt begging for honesty. The environment simply stopped
letting unverified claims ship, and when a link cannot be verified, it declines with the receipt
instead of guessing. Here is what survives.
New England rainfall totals from storm as flooding persists
Every delivered answer carries this chain: who
claimed it, the script that checked it, the fetch it survived, and the human signature at the bottom.
One command exports the whole audit bundle.
Before, wrong links were invisible.
Until now you had two options with an AI’s citation: accept it blindly, or
re-check every one yourself. And when one was wrong, nothing anywhere recorded that it was
wrong, or why.
That is the trade. You accept fewer answers, and in exchange every answer you do
accept has survived its checks, and every answer you refuse comes with the documented reason why.
Trust stops being a feeling and becomes a lookup.
The lies stopped shipping as answers.
Say that plainly: once a verification step exists and the caught claims are
declined with receipts, only about 1 in 8 answers that actually
reached you carried any error, and only about 1 in 29 was flatly wrong . Without it, about 7 in 10 were.
The engines are exactly as wrong as before; the caught claims arrive as labeled
declines instead of answers. The record carries a cost: a small multiple on tokens alone, fractions of a cent per answer, and grounded engines bill far more for search than token logging captures. The
effect held across every vendor and price tier tested, frontier models included.
Every number here has receipts.
Marcus is open source. The preregistration, the gates, the failed first pilot
and the passing re-run are all in the repository history, and every case shown above exports its
full audit bundle: who claimed, who checked, the evidence, who signed, what it cost. And nothing
in it is citation-specific. The checkpoint works for any claim a script can check: a number
against its source, code against its tests, a legal citation against the record. News citations were
simply the demo with a published baseline to replicate.
About the excerpts: the Tow Center distributes its 200 excerpts encrypted so the
text stays out of AI training crawlers. We keep that promise here, which is why excerpt text
appears only as short fragments drawn into an image, never as copyable text. Study: Jaźwińska
& Chandrasekar, “AI Search Has a Citation Problem,” Tow Center for Digital Journalism, CJR,
March 2025. Replication: October 2026, preregistered, graded blind. By Lawrence Gray, PhD,
Gray Labs; contact and press [email protected] ;
cite via CITATION.cff .

## Original Extract

A preregistered replication of Columbia

ON THE RECORD
Put it on the record →
Star
On the Record File № 737 · Oct 2026
A Marcus field report
The AI handed you a source. We clicked it.
Marcus is open-source software that makes AI agents prove their work before anyone has to trust it. This file is what happened when we pointed it at one hard, measurable problem: whether AI search engines cite real sources.
By Lawrence Gray, PhD, Gray Labs, who built Marcus. The dataset, the grader, and every audit record shown here are public, so you do not have to take my word for any of it.
We gave four AI search engines a sentence from each of 200 real news articles and
asked the same question every time: which article is this from? That produced 800 citations:
headline, publisher, date, link. We knew the true source of every sentence, so we graded each
answer and clicked every link to see whether it even exists.
So with any AI citation you get two choices: trust it blind, or re-check every one
by hand. We did the second: we clicked all 800. Here is what came back, three exhibits straight
from the run’s audit records.
The test re-runs the Tow Center’s 2025 study, “AI Search Has a Citation Problem.”
And citations are only the version of this that is easy to measure: the same failure hides in any
AI answer you act on without a way to check it, such as numbers in a report, legal references, medical
claims, or code.
“ ‘The Bachelor’ producers acknowledge ‘vicious cycle’ of racism…”
The page does not exist; the real source was
Entertainment Weekly. The answer carried a full headline, a confident date, and an invented link. Half the citations in our test
died exactly here. The reader never sees it: it is refused entry, and the refusal is on the record.
It is the same question, now put to a different engine. This time the link isn’t invented; it’s dead. Against a
dataset of articles two years old, 47% of the pages are already gone or closed to robots.
Fabricated or dead, a wrong link looks identical
to a right one, and you cannot tell which until you click. So the checkpoint clicks, every time, and
declines what does not resolve.
So we made them show their work.
An answer that might be confidently wrong is an answer you cannot use: someone
has to re-check it by hand, which cancels the point of asking an AI at all. So we changed the
rules instead of the model. Same engines, same questions, now passing through Marcus:
open-source software that works as a checkpoint between the AI and you, enforcing three
rules nobody can skip.
The fix used no retraining and no prompt begging for honesty. The environment simply stopped
letting unverified claims ship, and when a link cannot be verified, it declines with the receipt
instead of guessing. Here is what survives.
New England rainfall totals from storm as flooding persists
Every delivered answer carries this chain: who
claimed it, the script that checked it, the fetch it survived, and the human signature at the bottom.
One command exports the whole audit bundle.
Before, wrong links were invisible.
Until now you had two options with an AI’s citation: accept it blindly, or
re-check every one yourself. And when one was wrong, nothing anywhere recorded that it was
wrong, or why.
That is the trade. You accept fewer answers, and in exchange every answer you do
accept has survived its checks, and every answer you refuse comes with the documented reason why.
Trust stops being a feeling and becomes a lookup.
The lies stopped shipping as answers.
Say that plainly: once a verification step exists and the caught claims are
declined with receipts, only about 1 in 8 answers that actually
reached you carried any error, and only about 1 in 29 was flatly wrong . Without it, about 7 in 10 were.
The engines are exactly as wrong as before; the caught claims arrive as labeled
declines instead of answers. The record carries a cost: a small multiple on tokens alone, fractions of a cent per answer, and grounded engines bill far more for search than token logging captures. The
effect held across every vendor and price tier tested, frontier models included.
Every number here has receipts.
Marcus is open source. The preregistration, the gates, the failed first pilot
and the passing re-run are all in the repository history, and every case shown above exports its
full audit bundle: who claimed, who checked, the evidence, who signed, what it cost. And nothing
in it is citation-specific. The checkpoint works for any claim a script can check: a number
against its source, code against its tests, a legal citation against the record. News citations were
simply the demo with a published baseline to replicate.
About the excerpts: the Tow Center distributes its 200 excerpts encrypted so the
text stays out of AI training crawlers. We keep that promise here, which is why excerpt text
appears only as short fragments drawn into an image, never as copyable text. Study: Jaźwińska
& Chandrasekar, “AI Search Has a Citation Problem,” Tow Center for Digital Journalism, CJR,
March 2025. Replication: October 2026, preregistered, graded blind. By Lawrence Gray, PhD,
Gray Labs; contact and press [email protected] ;
cite via CITATION.cff .
