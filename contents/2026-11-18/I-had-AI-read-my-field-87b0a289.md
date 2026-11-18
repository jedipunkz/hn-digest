---
source: "https://science.caseyhandmer.com/physics/"
hn_url: "https://news.ycombinator.com/item?id=50542217"
title: "I had AI read my field"
article_title: "A field of physics, audited by machine"
image: ""
author: "gcoguiec"
captured_at: "2026-11-18T06:10:42Z"
capture_tool: "hn-digest"
hn_id: 50542217
score: 8
comments: 3
posted_at: "2026-11-18T05:26:41Z"
tags:
  - hacker-news
---

# I had AI read my field

- HN: [50542217](https://news.ycombinator.com/item?id=50542217)
- Source: [science.caseyhandmer.com](https://science.caseyhandmer.com/physics/)
- Score: 8
- Comments: 3
- Posted: 2026-11-18T05:26:41Z

## Translation

Title: I had AI read my field
Article title: A field of physics, audited by machine
Description: 158,059 astrophysics papers mapped in 3D, 659 of the most important checked by AI by recomputing them, and every verified error followed through the papers that cite it.

Article text:
Physics, audited · science.caseyhandmer.com
Story Errors Explore Subfields Review AI opportunities Method
◐
Loading…
Theoretical astrophysics · a whole field · audited by machine
I had AI read my entire field.
158,059 papers and 554,670 citations. Every paper with full text screened, 659 of the most important recomputed and checked, and every verified error followed through the papers that cite it. This is the map, and what it says about whether science corrects itself.
I spent five years in this field doing my PhD at Caltech, and a few months of that went to errors in the literature that were already known. The equations my code was built on came from Bishop, Gómez, Lehner, Maharaj & Winicour, “High-powered gravitational news” (1997), the paper the subfield's characteristic codes descend from. People knew it had mistakes. The fixes were scattered through some 10,000 lines of intersecting email threads that contradicted each other, and the paper itself was never corrected. The best I could do was correct the record in my own papers: an appendix setting out the full set of equations, made consistent, and a parenthetical “with a sign error in s 3 corrected” for anyone reading closely.
That is the limit of human literature review. Mine was not deep: my three thesis papers cite fewer than a hundred papers between them, in a field of tens of thousands. Nobody re-derives every equation they cite, referees check one paper once, and nobody is paid to go back and fix the original. So every student who builds on a broken paper pays the same tax, one at a time, and the correction lives in somebody's inbox. A machine can read the whole field, recompute it, and follow each error to every paper that copied it. So I had one do it. It found errors in my own papers too.
Height is time: the 1960s at the bottom, this year at the top. Each faint line is a citation. The bright dots are the papers everything else leans on, ranked the way Google ranked the early web.
The seed is TAPIR, Caltech's theoretical astrophysics group, where I did my PhD: 4,093 papers by its 326 members since 2004. Add every paper they cite and every paper that cites them and you get 158,059 papers joined by 554,670 citations. The 49,301 drawn here are the ones tied to TAPIR at least three times over.
The map holds 499 papers from the 1970s, 2,035 from the 1990s and 18,976 from the 2010s. The old layers are the survivors, the papers people working today still cite. Everything else from 1975 has already been forgotten, right or wrong.
Most of what this field knows was written in the last twenty years, by people who are still around to answer email. Hold that thought.
Nobody told the layout what numerical relativity is. Papers that cite each other pull together, so a subfield shows up as a column: the same few dozen people writing about the same problem year after year, each layer standing on the last.
There are 32 columns, from fast radio bursts to the Milky Way's halo.
Solving Einstein's equations on a computer until two black holes merge. 278 papers in this column, 71 of them TAPIR's, busiest in 2010.
Every subfield has a page: the questions it is organised around and whether each is open, contested or settled, who wrote the central papers and where, how fast it publishes, and which of its errors we traced.
Each open question also carries a flag for what today's AI could do about it. Of 155, a model judged 8 answerable now from data and codes that already exist, 119 ones AI would speed up, and 28 stuck waiting on a telescope or a detector.
Numerical relativity → All 32 subfields
Reading a paper carefully costs a few dollars of model time. Screening one costs nothing, so every paper I could get the full text of got screened: 50,269 papers, 421 million words, checked for the fingerprints of the worst kinds of trouble. That covers 76% of the papers drawn here. Nothing before 1991 is on arXiv and journal-only papers aren't open, so the rest got only the checks that run on a database record. (Gold marks TAPIR's own papers.)
Chatbot residue and hidden instructions to AI reviewers: none.
Text shared between papers with no author in common: 43 pairs in 34,286 papers, led by successive survey catalogues and data releases. The rest are queued for a human look.
Retractions: 13 among 184,733 papers with a DOI. None is TAPIR's.
TODO and FIXME notes shipped to arXiv: 62 papers.
Astrophysics is relatively clean. Whatever is wrong with this literature, it isn't fraud. The one needle that moves is vocabulary: “delve”, “crucial” and friends are 2.5 times as common in 2025 as in 2016.
The 200 that matter get checked line by line
Citations are brutally unequal, and a few hundred papers carry the field. So the 200 most-cited TAPIR papers of 2006–2019 (gold) each got their own AI reviewer with the LaTeX source, tools to run its own calculations, and orders to recompute everything it could: re-derive the equations, test the limits, rebuild the numbers from the paper's own tables. It worked blind, with no access to errata or later papers.
A flag is not a finding. Only what a second reviewer upholds is called an error here.
Thirteen fixed, fifty-six through
About 136 of the 200 were revised between the first preprint and the published paper, so we read those first preprints the same way. 13 verified errors were in the preprint and gone from the journal version, fixed by the referee or by the authors on a second look (purple).
56 went through review and into the journal (gold). On this evidence peer review removes about one error in five. At least two of the thirteen were copied by later papers anyway, from the stale preprint.
Smith, Kamionkowski & Cooray convert the gravitational waves left over from inflation into an energy density a detector could see. Two factors of one half went missing, so the answer is four times too big: a 6 in the denominator where the textbook has 24. This one was double-checked. The paper's numbers in the paper's formula reproduce its result; the same numbers in the standard formula give a quarter of it.
Watch the years run. Red papers copied the factor of four. The green one said it was wrong, in 2008. The copying carried on until 2014: six copies, one objection.
Five copies for every criticism
Then we cast the net wider. The same recompute-and-verify check, run on 459 more still-cited papers across the map, mostly highly cited ones that a cheaper first read had flagged, upheld 816 more errors. That makes 872 in 433 papers. 159 of them were copied, wrong value and all, into at least one later paper. Counting papers:
367 papers copied an error. 68 criticised one.
Only 47 of the 872 were ever criticised in print, and 21 of those kept being copied afterwards, by 51 more papers. These are floors: we read only the citing papers whose text we hold.
And 872 is a floor for the field too. Only 4% of the highly cited papers on the map got the thorough treatment so far. Scaled up from the papers that did, the likely number of undetected errors in this field's papers cited 100 or more times is about 34,000 (25,000–48,000). How that number is reached →
Half are still being copied 7 years later
How effective is peer review and publishing at detecting and expunging errors from the literature? We can actually measure it.
0% 25% 50% 75% 100% 0 2 4 6 8 10 12 years 42% at year 9
Share of copied errors still being copied, by years since the paper appeared. Band: 95% interval. 33 errors are old enough to be observed at year 10.
The curve is still at 52% at year 7 and crosses one half at year 8; at year 10 it is at 39%. Counting papers that lean on the affected result without showing which value they used, it crosses at year 9. The errors we traced most thoroughly last longer, crossing at year 10: a shallower search simply finds fewer late copies.
What finally stops an error? Of the 56 errors in the 200 papers read most closely, 52 are still in the latest arXiv version, where a fix is free. Of the 200 papers read most closely, the database records an erratum for 2. And of the 87 errors that did die out, 67 were never criticised by anyone. They faded when the subfield moved on, which retires right answers and wrong ones at the same rate.
4% of the papers that matter, so far
Every paper with open full text was screened and skimmed. The expensive check, the one that recomputes, has reached 659 papers, starting with the most cited. Each bar is the papers on the map with that many citations, filled by how far they have been read.
0 2,000 4,000 6,000 8,000 10,000 12,000 14,000 Papers 0 citations: 2,541 papers; Full text screened 1,463 (58%); Haiku skim 27 (1%); Sonnet read 0 (0%); Opus recompute, verified 0 (0%) 2,541 1–2 citations: 2,448 papers; Full text screened 1,618 (66%); Haiku skim 41 (2%); Sonnet read 0 (0%); Opus recompute, verified 0 (0%) 2,448 3–9 citations: 4,968 papers; Full text screened 3,493 (70%); Haiku skim 545 (11%); Sonnet read 0 (0%); Opus recompute, verified 0 (0%) 4,968 10–29 citations: 8,147 papers; Full text screened 6,286 (77%); Haiku skim 3,231 (40%); Sonnet read 26 (0%); Opus recompute, verified 0 (0%) 8,147 30–99 citations: 13,265 papers; Full text screened 10,759 (81%); Haiku skim 9,002 (68%); Sonnet read 742 (6%); Opus recompute, verified 10 (0%) 13,265 100–299 citations: 11,077 papers; Full text screened 8,972 (81%); Haiku skim 8,823 (80%); Sonnet read 1,047 (9%); Opus recompute, verified 324 (3%) 11,077 300–999 citations: 5,051 papers; Full text screened 3,591 (71%); Haiku skim 3,535 (70%); Sonnet read 512 (10%); Opus recompute, verified 256 (5%) 5,051 1,000–2,999 citations: 1,324 papers; Full text screened 718 (54%); Haiku skim 706 (53%); Sonnet read 67 (5%); Opus recompute, verified 60 (5%) 1,324 3,000+ citations: 403 papers; Full text screened 122 (30%); Haiku skim 117 (29%); Sonnet read 10 (2%); Opus recompute, verified 8 (2%) 403 0 1–2 3–9 10–29 30–99 100–299 300–999 1,000–2,999 3,000+ Times the paper has been cited
Coverage in detail →
A field outgrows its errors. It rarely corrects them.
All of this would cost about $2,720 in model time bought at list price, roughly $3 per verified error. In practice it ran within a Claude Max subscription. The whole project, this field and the others, took about 4.9 billion tokens and about 44 hours of my time over 41 days. It is one field, read by a first-generation method, with limits spelled out on the method page. And it is not finished: the deep check has reached only 4% of the highly cited papers on the map, most cited first. The rest of the queue is a budget decision, and so is the next field.
Errors were found by one AI session and re-derived by a second, independent one, and each was confirmed in the version the journal published. Five were also double-checked step by step against the source. No person has checked any of them. As a test, 40 drawn at random went to a third session that saw only the paper and the passage: it found 33 of the 40 wrong, so some named here will not survive scrutiny. If we have your paper wrong, say so and we will fix it and say that we did.

## Original Extract

158,059 astrophysics papers mapped in 3D, 659 of the most important checked by AI by recomputing them, and every verified error followed through the papers that cite it.

Physics, audited · science.caseyhandmer.com
Story Errors Explore Subfields Review AI opportunities Method
◐
Loading…
Theoretical astrophysics · a whole field · audited by machine
I had AI read my entire field.
158,059 papers and 554,670 citations. Every paper with full text screened, 659 of the most important recomputed and checked, and every verified error followed through the papers that cite it. This is the map, and what it says about whether science corrects itself.
I spent five years in this field doing my PhD at Caltech, and a few months of that went to errors in the literature that were already known. The equations my code was built on came from Bishop, Gómez, Lehner, Maharaj & Winicour, “High-powered gravitational news” (1997), the paper the subfield's characteristic codes descend from. People knew it had mistakes. The fixes were scattered through some 10,000 lines of intersecting email threads that contradicted each other, and the paper itself was never corrected. The best I could do was correct the record in my own papers: an appendix setting out the full set of equations, made consistent, and a parenthetical “with a sign error in s 3 corrected” for anyone reading closely.
That is the limit of human literature review. Mine was not deep: my three thesis papers cite fewer than a hundred papers between them, in a field of tens of thousands. Nobody re-derives every equation they cite, referees check one paper once, and nobody is paid to go back and fix the original. So every student who builds on a broken paper pays the same tax, one at a time, and the correction lives in somebody's inbox. A machine can read the whole field, recompute it, and follow each error to every paper that copied it. So I had one do it. It found errors in my own papers too.
Height is time: the 1960s at the bottom, this year at the top. Each faint line is a citation. The bright dots are the papers everything else leans on, ranked the way Google ranked the early web.
The seed is TAPIR, Caltech's theoretical astrophysics group, where I did my PhD: 4,093 papers by its 326 members since 2004. Add every paper they cite and every paper that cites them and you get 158,059 papers joined by 554,670 citations. The 49,301 drawn here are the ones tied to TAPIR at least three times over.
The map holds 499 papers from the 1970s, 2,035 from the 1990s and 18,976 from the 2010s. The old layers are the survivors, the papers people working today still cite. Everything else from 1975 has already been forgotten, right or wrong.
Most of what this field knows was written in the last twenty years, by people who are still around to answer email. Hold that thought.
Nobody told the layout what numerical relativity is. Papers that cite each other pull together, so a subfield shows up as a column: the same few dozen people writing about the same problem year after year, each layer standing on the last.
There are 32 columns, from fast radio bursts to the Milky Way's halo.
Solving Einstein's equations on a computer until two black holes merge. 278 papers in this column, 71 of them TAPIR's, busiest in 2010.
Every subfield has a page: the questions it is organised around and whether each is open, contested or settled, who wrote the central papers and where, how fast it publishes, and which of its errors we traced.
Each open question also carries a flag for what today's AI could do about it. Of 155, a model judged 8 answerable now from data and codes that already exist, 119 ones AI would speed up, and 28 stuck waiting on a telescope or a detector.
Numerical relativity → All 32 subfields
Reading a paper carefully costs a few dollars of model time. Screening one costs nothing, so every paper I could get the full text of got screened: 50,269 papers, 421 million words, checked for the fingerprints of the worst kinds of trouble. That covers 76% of the papers drawn here. Nothing before 1991 is on arXiv and journal-only papers aren't open, so the rest got only the checks that run on a database record. (Gold marks TAPIR's own papers.)
Chatbot residue and hidden instructions to AI reviewers: none.
Text shared between papers with no author in common: 43 pairs in 34,286 papers, led by successive survey catalogues and data releases. The rest are queued for a human look.
Retractions: 13 among 184,733 papers with a DOI. None is TAPIR's.
TODO and FIXME notes shipped to arXiv: 62 papers.
Astrophysics is relatively clean. Whatever is wrong with this literature, it isn't fraud. The one needle that moves is vocabulary: “delve”, “crucial” and friends are 2.5 times as common in 2025 as in 2016.
The 200 that matter get checked line by line
Citations are brutally unequal, and a few hundred papers carry the field. So the 200 most-cited TAPIR papers of 2006–2019 (gold) each got their own AI reviewer with the LaTeX source, tools to run its own calculations, and orders to recompute everything it could: re-derive the equations, test the limits, rebuild the numbers from the paper's own tables. It worked blind, with no access to errata or later papers.
A flag is not a finding. Only what a second reviewer upholds is called an error here.
Thirteen fixed, fifty-six through
About 136 of the 200 were revised between the first preprint and the published paper, so we read those first preprints the same way. 13 verified errors were in the preprint and gone from the journal version, fixed by the referee or by the authors on a second look (purple).
56 went through review and into the journal (gold). On this evidence peer review removes about one error in five. At least two of the thirteen were copied by later papers anyway, from the stale preprint.
Smith, Kamionkowski & Cooray convert the gravitational waves left over from inflation into an energy density a detector could see. Two factors of one half went missing, so the answer is four times too big: a 6 in the denominator where the textbook has 24. This one was double-checked. The paper's numbers in the paper's formula reproduce its result; the same numbers in the standard formula give a quarter of it.
Watch the years run. Red papers copied the factor of four. The green one said it was wrong, in 2008. The copying carried on until 2014: six copies, one objection.
Five copies for every criticism
Then we cast the net wider. The same recompute-and-verify check, run on 459 more still-cited papers across the map, mostly highly cited ones that a cheaper first read had flagged, upheld 816 more errors. That makes 872 in 433 papers. 159 of them were copied, wrong value and all, into at least one later paper. Counting papers:
367 papers copied an error. 68 criticised one.
Only 47 of the 872 were ever criticised in print, and 21 of those kept being copied afterwards, by 51 more papers. These are floors: we read only the citing papers whose text we hold.
And 872 is a floor for the field too. Only 4% of the highly cited papers on the map got the thorough treatment so far. Scaled up from the papers that did, the likely number of undetected errors in this field's papers cited 100 or more times is about 34,000 (25,000–48,000). How that number is reached →
Half are still being copied 7 years later
How effective is peer review and publishing at detecting and expunging errors from the literature? We can actually measure it.
0% 25% 50% 75% 100% 0 2 4 6 8 10 12 years 42% at year 9
Share of copied errors still being copied, by years since the paper appeared. Band: 95% interval. 33 errors are old enough to be observed at year 10.
The curve is still at 52% at year 7 and crosses one half at year 8; at year 10 it is at 39%. Counting papers that lean on the affected result without showing which value they used, it crosses at year 9. The errors we traced most thoroughly last longer, crossing at year 10: a shallower search simply finds fewer late copies.
What finally stops an error? Of the 56 errors in the 200 papers read most closely, 52 are still in the latest arXiv version, where a fix is free. Of the 200 papers read most closely, the database records an erratum for 2. And of the 87 errors that did die out, 67 were never criticised by anyone. They faded when the subfield moved on, which retires right answers and wrong ones at the same rate.
4% of the papers that matter, so far
Every paper with open full text was screened and skimmed. The expensive check, the one that recomputes, has reached 659 papers, starting with the most cited. Each bar is the papers on the map with that many citations, filled by how far they have been read.
0 2,000 4,000 6,000 8,000 10,000 12,000 14,000 Papers 0 citations: 2,541 papers; Full text screened 1,463 (58%); Haiku skim 27 (1%); Sonnet read 0 (0%); Opus recompute, verified 0 (0%) 2,541 1–2 citations: 2,448 papers; Full text screened 1,618 (66%); Haiku skim 41 (2%); Sonnet read 0 (0%); Opus recompute, verified 0 (0%) 2,448 3–9 citations: 4,968 papers; Full text screened 3,493 (70%); Haiku skim 545 (11%); Sonnet read 0 (0%); Opus recompute, verified 0 (0%) 4,968 10–29 citations: 8,147 papers; Full text screened 6,286 (77%); Haiku skim 3,231 (40%); Sonnet read 26 (0%); Opus recompute, verified 0 (0%) 8,147 30–99 citations: 13,265 papers; Full text screened 10,759 (81%); Haiku skim 9,002 (68%); Sonnet read 742 (6%); Opus recompute, verified 10 (0%) 13,265 100–299 citations: 11,077 papers; Full text screened 8,972 (81%); Haiku skim 8,823 (80%); Sonnet read 1,047 (9%); Opus recompute, verified 324 (3%) 11,077 300–999 citations: 5,051 papers; Full text screened 3,591 (71%); Haiku skim 3,535 (70%); Sonnet read 512 (10%); Opus recompute, verified 256 (5%) 5,051 1,000–2,999 citations: 1,324 papers; Full text screened 718 (54%); Haiku skim 706 (53%); Sonnet read 67 (5%); Opus recompute, verified 60 (5%) 1,324 3,000+ citations: 403 papers; Full text screened 122 (30%); Haiku skim 117 (29%); Sonnet read 10 (2%); Opus recompute, verified 8 (2%) 403 0 1–2 3–9 10–29 30–99 100–299 300–999 1,000–2,999 3,000+ Times the paper has been cited
Coverage in detail →
A field outgrows its errors. It rarely corrects them.
All of this would cost about $2,720 in model time bought at list price, roughly $3 per verified error. In practice it ran within a Claude Max subscription. The whole project, this field and the others, took about 4.9 billion tokens and about 44 hours of my time over 41 days. It is one field, read by a first-generation method, with limits spelled out on the method page. And it is not finished: the deep check has reached only 4% of the highly cited papers on the map, most cited first. The rest of the queue is a budget decision, and so is the next field.
Errors were found by one AI session and re-derived by a second, independent one, and each was confirmed in the version the journal published. Five were also double-checked step by step against the source. No person has checked any of them. As a test, 40 drawn at random went to a third session that saw only the paper and the passage: it found 33 of the 40 wrong, so some named here will not survive scrutiny. If we have your paper wrong, say so and we will fix it and say that we did.
