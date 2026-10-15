---
source: "https://dgjones.info/moat.html"
hn_url: "https://news.ycombinator.com/item?id=50102562"
title: "Open source for 27 years, unread: what LLMs changed about the Xanadu code"
article_title: "The Moat Was Never the Code"
image: ""
author: "jgdgj"
captured_at: "2026-10-15T13:07:18Z"
capture_tool: "hn-digest"
hn_id: 50102562
score: 1
comments: 0
posted_at: "2026-10-15T12:15:22Z"
tags:
  - hacker-news
---

# Open source for 27 years, unread: what LLMs changed about the Xanadu code

- HN: [50102562](https://news.ycombinator.com/item?id=50102562)
- Source: [dgjones.info](https://dgjones.info/moat.html)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T12:15:22Z

## Translation

Title: Open source for 27 years, unread: what LLMs changed about the Xanadu code
Article title: The Moat Was Never the Code
Description: Why the 1999 Udanax Gold release sat unread for 27 years — and what large language models changed about reading old code.

Article text:
In 1999, the Xanadu project released the source code to Udanax Gold — Ted Nelson's second hypertext engine, the system meant to give literature permanent, visible, two-way connections — under a permissive license. Free to read, free to use, free to build on. By then the Web was eight years old and had famously shipped a fraction of the vision: one-way links to whole pages, copies everywhere, nothing permanent. The deep version was in the box. Anyone could have it.
For twenty-seven years, almost nobody touched it.
That is the puzzle, and the easy explanation is wrong. The easy explanation says the code was broken — vaporware that deserved its reputation. This year I went to find out. The release, it turns out, contains the complete back end of the system — written in a custom Smalltalk dialect, with its cross-compiled C++ sitting alongside.
Udanax Gold was written once, in an unusual way. The source is Smalltalk — 732 classes filed out from Objectworks for Smalltalk-80, version 2.5 of July 1989, last touched in March 1995, organized under categories like Xanadu-Be-Canopy : the canopy trees, the tumblers, the crums, the link machinery, the space algebra the whole model runs on. But it is not quite the Smalltalk you know. The methods carry type annotations — a constrained dialect built for one purpose: the code could be developed in a live Smalltalk environment and then cross-compiled .
{MuSet of: CanopyCrum} pathFor: canopyCrum {CanopyCrum}
"Return the set of all crums from canopyCrum
(inclusive) to the top of canopyCrum's canopy."
canopy.st, line 37 — typed Smalltalk, written for cross-compilation (fileOut dated March 1995)
The C++ in the release is that compilation's output. The headers still carry the translator's banner:
***************************************************************************
Output from Objectworks for Smalltalk-80(tm), Version 2.5 of 29 July 1989
*/
#ifndef CANOPYX_HXX
#define CANOPYX_HXX
canopyx.hxx, line 13 — the generated C++ opens by naming the tool that made it
The same class names ( Heaper , CanopyCrum , MuSet ) appear throughout both languages.
The release even ships the base Smalltalk class library alongside the application classes, so the source is loadable, not just readable — though good luck finding a 1989 Objectworks to load it into.
So what was missing? Not the model, and not the code — which exists twice over, as source and as generated C++. The team had dispersed when the money ended, mid-productization. There was no public documentation to speak of — the team's knowledge lived in heads and in internal recordings that never made it into any release — and no map. The path in ran through a custom dialect, a vanished cross-compiler, and generated C++ that no human ever wrote by hand. The system was orphaned, not broken.
Here is the thing about open source: it drops the price of code to zero. It does nothing to the price of understanding.
To enter that codebase in, say, 2010, you needed to hold a sixty-year-old design in your head (Nelson's model, which the documentation mostly doesn't explain, because the authors assumed you already spoke Xanadu), plus thirty-year-old code in a dead dialect nobody speaks, plus generated C++ no human ever wrote by hand, plus the vanished cross-compiler that connected them. The literature consistently estimates onboarding costs of months for codebases far friendlier than this. Against that, the license being permissive was almost a joke. The moat was comprehension, and comprehension held the door for twenty-seven years. I know it held me off for years.
This is not a story about Xanadu specifically. It's a pattern: software dies not when it's broken but when the cost of understanding it exceeds anyone's reason to pay. Open source promised to fix this and mostly fixed the wrong half. The right to read is not the ability to read.
What changed in the last few years is not that large language models learned to write code. The interesting change is what they did to reading.
Large language models are, whatever else they are, comprehension accelerants for old code. They have ingested most of the world's programming languages, including dead ones, and they can do the three things an archaeologist actually needs — read an unfamiliar idiom, locate where a mechanism lives, explain what a module is doing — in seconds instead of days.
From a year of using them on exactly this codebase, my estimate of the speedup is a spectrum, not a number:
Reading, locating, explaining: roughly 100× .
Translation between languages: roughly 50× .
Building a new system around an understood core: maybe 5–10× .
Design judgment: 1× — no change. The models are worse than useless at deciding what should exist, which turns out to be the job.
The spectrum shape matters more than the values. The cheap end is exactly what held the door shut for twenty-seven years. The expensive end — judgment — is exactly what the original designers had and the moat was protecting.
Evidence beats adjectives, so here is mine.
Product scale. I started with the 1992 C++ engine and migrated it to Rust. With model assistance, that took weeks. Alone it would have taken years: the engine was full of algorithms I had never seen, and there was a lot of it. And translation, remember, was the middle of the spectrum — not the cheap end. Then I kept building, because the point was never the port. Five months after starting, there was a live, multi-user system implementing the full model — transclusion at character level, typed multi-ended links, deep versioning, per-passage provenance — built by one developer, carrying 3,800+ Rust tests on the engine and another 1,200 on the TypeScript client. I am not a team. I am not funded. I had the model doing the reading.
Verification keeps the accelerant honest
There is a trap in all this, and it's the interesting one. The build caught bugs that reading could not have found — latent, silent, wrapped in assumptions nobody thought to question. The splay machinery I had carried over — a design the original needed because every access descends the tree — held a content-loss bug in its terminal arms. No reader would ever have seen it, because reading checks what the code says, not what the model guarantees. Measuring the implementation against those guarantees surfaced it; it is fixed and regression-tested now — and the same measurement said the splay wasn't even needed for my access pattern. The same discipline — a dual-write shadow, running the new engine beside the old one on every operation — caught merge defects the primary implementation hid entirely.
That is the lesson, and it's backwards from how code review usually works. When reading is cheap, verification has to move from "check the lineage" to "check against the spec." Trust re-derivation, not ancestry — the ancestry is exactly what a fast translation contaminates.
The thesis of all this is falsifiable: the Xanadu model was buildable all along, and the barrier was never the model's merits but the cost of comprehending the artifacts that carried it. The way you test such a claim is not to argue about it. You build the thing.
So, as evidence rather than advertisement: a complete, independent, open-source implementation of the Gold model exists and is running — with the parts the original never had the years for, like realtime collaborative editing on a CRDT engine and per-passage cryptographic attribution. You can try it at xudanu.com ; the code is at github.com/jonesd/xudanu ; and the documentation — including a map from every Gold concept to its Rust descendant — is at dgjones.info/xudanu . It is an independent re-implementation of the model from the 1999 release, not affiliated with Project Xanadu. The model is Nelson's and his team's. The reading is what got cheap.
Now ask the question this story is actually about. The Xanadu release sat open and comprehensible-to-no-one for 27 years. How much else is sitting in the literature the same way — designed, finished on paper, sometimes even implemented, and unreadable at a price anyone would pay? Bush's associative trails. Engelbart's view specs. Intermedia's linkbases. Half the good ideas in the 1980s and 1990s hypertext literature.
We tell ourselves the bottleneck was that these ideas were impractical. For some, maybe. But the Xanadu case says something less comfortable: at least sometimes, the ideas were practical and the entry cost was just mispriced — and nobody re-priced it for thirty years, because repackaging old code for new readers was a thing nobody could sell.
Some of those designs just became readable again. The question is no longer only "what should we build?" It's "what was already designed, and waiting on a cost curve?"
Disclosure. This work — the migration, this essay's drafts — was done with LLM assistance throughout. The design decisions, the mistakes, and any remaining wrong ideas are mine.

## Original Extract

Why the 1999 Udanax Gold release sat unread for 27 years — and what large language models changed about reading old code.

In 1999, the Xanadu project released the source code to Udanax Gold — Ted Nelson's second hypertext engine, the system meant to give literature permanent, visible, two-way connections — under a permissive license. Free to read, free to use, free to build on. By then the Web was eight years old and had famously shipped a fraction of the vision: one-way links to whole pages, copies everywhere, nothing permanent. The deep version was in the box. Anyone could have it.
For twenty-seven years, almost nobody touched it.
That is the puzzle, and the easy explanation is wrong. The easy explanation says the code was broken — vaporware that deserved its reputation. This year I went to find out. The release, it turns out, contains the complete back end of the system — written in a custom Smalltalk dialect, with its cross-compiled C++ sitting alongside.
Udanax Gold was written once, in an unusual way. The source is Smalltalk — 732 classes filed out from Objectworks for Smalltalk-80, version 2.5 of July 1989, last touched in March 1995, organized under categories like Xanadu-Be-Canopy : the canopy trees, the tumblers, the crums, the link machinery, the space algebra the whole model runs on. But it is not quite the Smalltalk you know. The methods carry type annotations — a constrained dialect built for one purpose: the code could be developed in a live Smalltalk environment and then cross-compiled .
{MuSet of: CanopyCrum} pathFor: canopyCrum {CanopyCrum}
"Return the set of all crums from canopyCrum
(inclusive) to the top of canopyCrum's canopy."
canopy.st, line 37 — typed Smalltalk, written for cross-compilation (fileOut dated March 1995)
The C++ in the release is that compilation's output. The headers still carry the translator's banner:
***************************************************************************
Output from Objectworks for Smalltalk-80(tm), Version 2.5 of 29 July 1989
*/
#ifndef CANOPYX_HXX
#define CANOPYX_HXX
canopyx.hxx, line 13 — the generated C++ opens by naming the tool that made it
The same class names ( Heaper , CanopyCrum , MuSet ) appear throughout both languages.
The release even ships the base Smalltalk class library alongside the application classes, so the source is loadable, not just readable — though good luck finding a 1989 Objectworks to load it into.
So what was missing? Not the model, and not the code — which exists twice over, as source and as generated C++. The team had dispersed when the money ended, mid-productization. There was no public documentation to speak of — the team's knowledge lived in heads and in internal recordings that never made it into any release — and no map. The path in ran through a custom dialect, a vanished cross-compiler, and generated C++ that no human ever wrote by hand. The system was orphaned, not broken.
Here is the thing about open source: it drops the price of code to zero. It does nothing to the price of understanding.
To enter that codebase in, say, 2010, you needed to hold a sixty-year-old design in your head (Nelson's model, which the documentation mostly doesn't explain, because the authors assumed you already spoke Xanadu), plus thirty-year-old code in a dead dialect nobody speaks, plus generated C++ no human ever wrote by hand, plus the vanished cross-compiler that connected them. The literature consistently estimates onboarding costs of months for codebases far friendlier than this. Against that, the license being permissive was almost a joke. The moat was comprehension, and comprehension held the door for twenty-seven years. I know it held me off for years.
This is not a story about Xanadu specifically. It's a pattern: software dies not when it's broken but when the cost of understanding it exceeds anyone's reason to pay. Open source promised to fix this and mostly fixed the wrong half. The right to read is not the ability to read.
What changed in the last few years is not that large language models learned to write code. The interesting change is what they did to reading.
Large language models are, whatever else they are, comprehension accelerants for old code. They have ingested most of the world's programming languages, including dead ones, and they can do the three things an archaeologist actually needs — read an unfamiliar idiom, locate where a mechanism lives, explain what a module is doing — in seconds instead of days.
From a year of using them on exactly this codebase, my estimate of the speedup is a spectrum, not a number:
Reading, locating, explaining: roughly 100× .
Translation between languages: roughly 50× .
Building a new system around an understood core: maybe 5–10× .
Design judgment: 1× — no change. The models are worse than useless at deciding what should exist, which turns out to be the job.
The spectrum shape matters more than the values. The cheap end is exactly what held the door shut for twenty-seven years. The expensive end — judgment — is exactly what the original designers had and the moat was protecting.
Evidence beats adjectives, so here is mine.
Product scale. I started with the 1992 C++ engine and migrated it to Rust. With model assistance, that took weeks. Alone it would have taken years: the engine was full of algorithms I had never seen, and there was a lot of it. And translation, remember, was the middle of the spectrum — not the cheap end. Then I kept building, because the point was never the port. Five months after starting, there was a live, multi-user system implementing the full model — transclusion at character level, typed multi-ended links, deep versioning, per-passage provenance — built by one developer, carrying 3,800+ Rust tests on the engine and another 1,200 on the TypeScript client. I am not a team. I am not funded. I had the model doing the reading.
Verification keeps the accelerant honest
There is a trap in all this, and it's the interesting one. The build caught bugs that reading could not have found — latent, silent, wrapped in assumptions nobody thought to question. The splay machinery I had carried over — a design the original needed because every access descends the tree — held a content-loss bug in its terminal arms. No reader would ever have seen it, because reading checks what the code says, not what the model guarantees. Measuring the implementation against those guarantees surfaced it; it is fixed and regression-tested now — and the same measurement said the splay wasn't even needed for my access pattern. The same discipline — a dual-write shadow, running the new engine beside the old one on every operation — caught merge defects the primary implementation hid entirely.
That is the lesson, and it's backwards from how code review usually works. When reading is cheap, verification has to move from "check the lineage" to "check against the spec." Trust re-derivation, not ancestry — the ancestry is exactly what a fast translation contaminates.
The thesis of all this is falsifiable: the Xanadu model was buildable all along, and the barrier was never the model's merits but the cost of comprehending the artifacts that carried it. The way you test such a claim is not to argue about it. You build the thing.
So, as evidence rather than advertisement: a complete, independent, open-source implementation of the Gold model exists and is running — with the parts the original never had the years for, like realtime collaborative editing on a CRDT engine and per-passage cryptographic attribution. You can try it at xudanu.com ; the code is at github.com/jonesd/xudanu ; and the documentation — including a map from every Gold concept to its Rust descendant — is at dgjones.info/xudanu . It is an independent re-implementation of the model from the 1999 release, not affiliated with Project Xanadu. The model is Nelson's and his team's. The reading is what got cheap.
Now ask the question this story is actually about. The Xanadu release sat open and comprehensible-to-no-one for 27 years. How much else is sitting in the literature the same way — designed, finished on paper, sometimes even implemented, and unreadable at a price anyone would pay? Bush's associative trails. Engelbart's view specs. Intermedia's linkbases. Half the good ideas in the 1980s and 1990s hypertext literature.
We tell ourselves the bottleneck was that these ideas were impractical. For some, maybe. But the Xanadu case says something less comfortable: at least sometimes, the ideas were practical and the entry cost was just mispriced — and nobody re-priced it for thirty years, because repackaging old code for new readers was a thing nobody could sell.
Some of those designs just became readable again. The question is no longer only "what should we build?" It's "what was already designed, and waiting on a cost curve?"
Disclosure. This work — the migration, this essay's drafts — was done with LLM assistance throughout. The design decisions, the mistakes, and any remaining wrong ideas are mine.
