---
source: "https://transept.ai/journal/how-human-and-ai-translation-differ"
hn_url: "https://news.ycombinator.com/item?id=50655346"
title: "How do human and AI book translation differ, and what makes readers like both?"
article_title: "How human and AI translation differ, and what readers like"
image: "https://transept.ai/images/journal/covers/how-human-and-ai-translation-differ.png"
author: "sithamet"
captured_at: "2026-11-26T14:45:42Z"
capture_tool: "hn-digest"
hn_id: 50655346
score: 1
comments: 1
posted_at: "2026-11-26T13:56:42Z"
tags:
  - hacker-news
---

# How do human and AI book translation differ, and what makes readers like both?

- HN: [50655346](https://news.ycombinator.com/item?id=50655346)
- Source: [transept.ai](https://transept.ai/journal/how-human-and-ai-translation-differ)
- Score: 1
- Comments: 1
- Posted: 2026-11-26T13:56:42Z

## Translation

Title: How do human and AI book translation differ, and what makes readers like both?
Article title: How human and AI translation differ, and what readers like
Description: Machines keep 89–98 of 100 units of the original, published translators 73–76. What readers and experts reward in each, across four public corpora.

Article text:
How human and AI translation differ, and what readers like “Fifty percent. I counted twice.”
Black Friday 50% off your first month on any plan with code BF26 , until Monday, November 30. See plans →
Transept Product Library Guides How-tos you can follow step by step Journal Experiments, author stories and notes Help How to use Transept Changelog What changed, week by week Pricing About Start translating → Transept Product Library Pricing About Start translating → Product
How do human and AI translation differ, and what makes readers like both? A cross-corpus empirical re-analysis and LLM-as-a-judge meta-evaluation
Using four public corpora (LAIT, ParaFarm, PAR3 and WMT25 literary), we measured what separates AI from human translation and what readers and experts reward. AI keeps the information; humans spend it on voice. Readers punish stumbles; experts who read the original punish changed meaning.
On this page Introduction, or why we can't automate translation quality yet
Theory one: does a translation have to carry every unit of meaning and style?
Why did we expect every unit to matter?
How did we measure meaning and style?
What do translators do with the information?
Which units do translators rework?
Is information density a property of the language?
Do readers reward keeping the information?
Why do experts who read the original judge differently?
How far can the unit tracker be trusted?
Theory two: what breaks a reader's flow?
What do the reading judges check?
What didn't work, and what did we drop?
So, how do AI and human translation differ?
…And what makes a translation good?
What makes human translators still better?
What does this mean if you translate a book with AI?
How did we use AI in this research?
Scope. Using four public corpora ( LAIT : 15 novels, 386 passages, two readers each; plus expert-rated ParaFarm , PAR3 , and WMT25 literary ), we ran new measurements on existing translations, generated new AI translations on LAIT, validated LLM judges against human ratings, and ran diagnostic tests on our own texts.
AI keeps the information; humans spend it on voice. On LAIT, machines keep 89–98 of 100 source units of meaning and style vs. 73–76 for published translators (who rework emphasis, images, and details most). It separates human from machine in 97% of passages.
Readers don't reward preservation . Within a book, nothing about what a translator kept, cut or changed predicts which version readers prefer. Across books, the translators whom readers did not like were the ones who expanded on net.
Readers punish stumbles. Source-blind AI judges counting stumbles (calques, literal phrasing, odd syntax, grammar) match readers as closely as two readers match each other. Flags for register, flatness, or wordiness barely track them.
Seeing the source flips the verdict. Experts with the original punish changed meaning and heavy reworking, not compression or cuts, ranking the same translations almost opposite a source-blind judge. The translations that win both audiences keep the information in a native voice.
AI protects language standards; humans protect authorial artifacts. AI translators and judges (Claude 5.5 and GPT-6 alike) preserve source structure and smooth toward target norms. That is why model judges rank top AI above published translators, a lead that collapses to a near coin flip once stumbles and copy-editing are removed. We will develop the method further to capture what constitutes translation quality.
Where Transept lands. Transept Pro keeps 98 of 100 units of the original, like the study MT, and of every version we measured it draws the fewest flags for the stumbles readers notice most: sentence structure, unnatural wording, grammar and calques. Its copy-editing load is 0.08 fixes per 100 words against 0.26 for the published translations. No reader has compared it with a published translation, and our judges overrate AI, so read this as a stumble count, not a verdict on craft. I co-founded Transept; the same judges scored every version, and their prompts are public.
Introduction, or why we can't automate translation quality yet
In a 2026 study of fifteen translated novels, LAIT (Ferstler et al.), avid readers (12–120 books a year) tried to tell human-translated excerpts from AI ones. They guessed right 34 times out of 60: essentially a coin flip. Asked which version they actually preferred, blind, they picked the human translation in 522 of 772 passage comparisons (67.6%). In our guide to translating a book without losing your voice , we traced that gap to four losses: literal idioms, copied source syntax, drifting character voices, and flattened wordplay.
Other studies find the same split between what machines produce and what people value:
Reviewers rated from-scratch literary translations most creative and raw MT least; post-editing MT constrained translators' creativity (Guerberof-Arenas & Toral 2022).
Post-edited translations stay up to 38% closer to source sentence lengths than from-scratch ones, with flatter vocabulary and source syntax showing through (Toral 2019).
In ParaFarm (seven published Ukrainian Animal Farm translations vs. three AIs), AI used up to 2× fewer particles, 2.6× fewer diminutives, and 18% lower lexical diversity, converging at 0.94 similarity vs. 0.71 for humans (Chaplynskyi et al. 2026).
In LitEval-Corpus, professional raters preferred the human translation 94% of the time; student raters, 42% (Zhang et al. 2025).
LAIT's three automatic metrics (MetricX-QE, GEMBA-MQM, LiTransProQA) predicted readers at chance (0.48–0.50), and two picked the machine almost every time.
Building our own AI editors and judges, we hit the same gap from the inside. The smoothest drafts stripped the most nuance, source-aware judges forgave calques, and models ironed out deliberate quirks. It appeared that natively, LLM-as-Judge approach cannot measure translation quality.
So we set out to measure what separates AI from human translation, and what separates a good translation from a bad one. We tested two theories: first , that a text carries units of meaning and style and breaks if a translation fails to move all of them; second , that what readers actually feel is reading flow, and the places where it stumbles.
Our data combines LAIT (French, Polish, and Japanese into English, with two readers' choices per passage), three public datasets rated by translators who saw the original (ParaFarm, PAR3, WMT25 literary), and our own texts, including my novel's prologue translated into Ukrainian.
Theory one: does a translation have to carry every unit of meaning and style?
In short: No. Machines carry almost every unit of information in the original; good human translators don't, and readers don't hold it against them. Readers who never see the original punish expansion. Experts who read the original punish changed meaning.
How much of the original survives is the clearest difference between human and machine translation we measured.
Why did we expect every unit to matter?
A book invents language artifacts (names, coined terms, rhythm, running jokes) that work on several levels at once, while leaning on target-language standards (idioms, register, punctuation, syntax). Translation pulls both ways. We split text into meaning (what a reader learns) and style (how it reads), betting that dropping pieces of either layer makes a translation break.
How did we measure meaning and style?
How to read: one judge per row: what it sees, what it checks, and the number it gives.
The first two work. The meaning judge catches 51 of 61 real errors on our answer key while flagging 10 of 33 deliberate liberties in published translations. The rubric caught all 61 planted errors (57 of 59 on Gemini) and exposed what smooth prose hides: translating through English lost 12–18 elements per 1,000 words vs. 0–1 directly, with only 0–2.4 hard errors.
Neither predicts readers (rubric losses score −0.05 against LAIT). Only the rubric's source-blind native-reader check tracked them (+0.23 on tuning books): our first clue that readers react to how a text reads, not what it kept.
What do translators do with the information?
Machines keep nearly everything. Published translators keep about three units in four and rework the rest.
How to read: a longer "kept" bar means closer to the original; the other segments are what was reworked, and added units sit outside the 100. Neither end is "better" on its own.
Source: 135 LAIT passages; one Gemini 4 Hydrogen call (high thinking) per passage marks all 14 versions blind against one list of units. Compressed and changed are upper bounds (see "How far can the unit tracker be trusted?" ).
The pattern holds in every corpus we ran:
How to read: units of the original kept per 100. Higher means closer to the original, not a better translation.
How much survives is the best single tell between a human and a machine translation. On LAIT, the published translation keeps fewer units than the study's MT in 374 of 386 passages (97%) and fewer than Transept Pro in 381 (99%). Among the 14 versions in Chart 1, the human one keeps the fewest units in 115 of 135 passages.
The one machine that compresses like a human is GPT-4.1 in WMT25: it keeps 76 and writes about 12% shorter than the others, a concise style rather than truncation.
For comparison, our reading score puts the published translation above the study MT in 66% of passages (the readers: 67.6%), and the published version is longer than the MT in only 59%.
How to read: one passage, four English versions, and what the tracker marked for each.
«А ще воно відкривало її пружні плечі та жіночний живіт — під яким грали м'язи, коли Агнес сміялася чи затягувалася сигаретою. І те, й інше магеса любила.»
“ It also bared her firm shoulders and womanly belly, under which muscles played whenever Agnes laughed or drew on a cigarette. The magess loved both. ”
“ It also bared her firm shoulders and her womanly belly , where muscles played whenever Agnes laughed or drew on a cigarette. The magess loved doing both. ”
Sonnet and Haiku read almost the same. Tracker: Sonnet keeps every unit; Opus too, except that "under which" becoming "where" counts as compressed. A reader judge flagged Opus's "her womanly belly" as hard to follow.
“ The outfit bared her toned shoulders and the soft curve of her stomach , where muscle stirred beneath the skin whenever she laughed or drew on her cigarette, two pleasures the magess rarely denied herself . ”
Tracker: "womanly belly" changed, "the magess loved both" changed. No reader judge flagged these lines.
the models carry the units over word for word; my version says the same things the English way, and the tracker counts that as rework. Across this chapter my version keeps 71–77 units in 100, the three models 87–100.
Which units do translators rework?
The ones that carry voice. Facts and events stay; emphasis, images and details move.
How to read: the share of each unit type a version reworked. Higher means more reworking, not worse.
"Reworked" means compressed, changed, or dropped across Chart 1's 135 passages (Gemini 4 Hydrogen; reference and emphasis are the smallest groups, so their columns are noisiest).
How to read: further right means more of that unit type was reworked. The gap between the big published dot and the cluster of machine dots is the human–machine difference.
Source: one Gemini 4 Hydrogen call (high thinking) per passage marks all 14 versions blind against one list of units; group sizes 214 emphasis, 247 image, 994 detail, 127 reference, 902 feeling, 1,758 fact, 1,706 event, 555 logic, 205 speaker units. Shares are upper bounds (see "How far can the unit tracker be trusted?" ). The cheaper models' higher shares (Haiku, the Flash Lites, 3.8 Flash) are more likely slips than deliberate reworking.
Translators also agree on which units are open. In ParaFarm, units reworked by at least four of the six published translators appear 56 times (against 11 by chance), and un

[truncated]

## Original Extract

Machines keep 89–98 of 100 units of the original, published translators 73–76. What readers and experts reward in each, across four public corpora.

How human and AI translation differ, and what readers like “Fifty percent. I counted twice.”
Black Friday 50% off your first month on any plan with code BF26 , until Monday, November 30. See plans →
Transept Product Library Guides How-tos you can follow step by step Journal Experiments, author stories and notes Help How to use Transept Changelog What changed, week by week Pricing About Start translating → Transept Product Library Pricing About Start translating → Product
How do human and AI translation differ, and what makes readers like both? A cross-corpus empirical re-analysis and LLM-as-a-judge meta-evaluation
Using four public corpora (LAIT, ParaFarm, PAR3 and WMT25 literary), we measured what separates AI from human translation and what readers and experts reward. AI keeps the information; humans spend it on voice. Readers punish stumbles; experts who read the original punish changed meaning.
On this page Introduction, or why we can't automate translation quality yet
Theory one: does a translation have to carry every unit of meaning and style?
Why did we expect every unit to matter?
How did we measure meaning and style?
What do translators do with the information?
Which units do translators rework?
Is information density a property of the language?
Do readers reward keeping the information?
Why do experts who read the original judge differently?
How far can the unit tracker be trusted?
Theory two: what breaks a reader's flow?
What do the reading judges check?
What didn't work, and what did we drop?
So, how do AI and human translation differ?
…And what makes a translation good?
What makes human translators still better?
What does this mean if you translate a book with AI?
How did we use AI in this research?
Scope. Using four public corpora ( LAIT : 15 novels, 386 passages, two readers each; plus expert-rated ParaFarm , PAR3 , and WMT25 literary ), we ran new measurements on existing translations, generated new AI translations on LAIT, validated LLM judges against human ratings, and ran diagnostic tests on our own texts.
AI keeps the information; humans spend it on voice. On LAIT, machines keep 89–98 of 100 source units of meaning and style vs. 73–76 for published translators (who rework emphasis, images, and details most). It separates human from machine in 97% of passages.
Readers don't reward preservation . Within a book, nothing about what a translator kept, cut or changed predicts which version readers prefer. Across books, the translators whom readers did not like were the ones who expanded on net.
Readers punish stumbles. Source-blind AI judges counting stumbles (calques, literal phrasing, odd syntax, grammar) match readers as closely as two readers match each other. Flags for register, flatness, or wordiness barely track them.
Seeing the source flips the verdict. Experts with the original punish changed meaning and heavy reworking, not compression or cuts, ranking the same translations almost opposite a source-blind judge. The translations that win both audiences keep the information in a native voice.
AI protects language standards; humans protect authorial artifacts. AI translators and judges (Claude 5.5 and GPT-6 alike) preserve source structure and smooth toward target norms. That is why model judges rank top AI above published translators, a lead that collapses to a near coin flip once stumbles and copy-editing are removed. We will develop the method further to capture what constitutes translation quality.
Where Transept lands. Transept Pro keeps 98 of 100 units of the original, like the study MT, and of every version we measured it draws the fewest flags for the stumbles readers notice most: sentence structure, unnatural wording, grammar and calques. Its copy-editing load is 0.08 fixes per 100 words against 0.26 for the published translations. No reader has compared it with a published translation, and our judges overrate AI, so read this as a stumble count, not a verdict on craft. I co-founded Transept; the same judges scored every version, and their prompts are public.
Introduction, or why we can't automate translation quality yet
In a 2026 study of fifteen translated novels, LAIT (Ferstler et al.), avid readers (12–120 books a year) tried to tell human-translated excerpts from AI ones. They guessed right 34 times out of 60: essentially a coin flip. Asked which version they actually preferred, blind, they picked the human translation in 522 of 772 passage comparisons (67.6%). In our guide to translating a book without losing your voice , we traced that gap to four losses: literal idioms, copied source syntax, drifting character voices, and flattened wordplay.
Other studies find the same split between what machines produce and what people value:
Reviewers rated from-scratch literary translations most creative and raw MT least; post-editing MT constrained translators' creativity (Guerberof-Arenas & Toral 2022).
Post-edited translations stay up to 38% closer to source sentence lengths than from-scratch ones, with flatter vocabulary and source syntax showing through (Toral 2019).
In ParaFarm (seven published Ukrainian Animal Farm translations vs. three AIs), AI used up to 2× fewer particles, 2.6× fewer diminutives, and 18% lower lexical diversity, converging at 0.94 similarity vs. 0.71 for humans (Chaplynskyi et al. 2026).
In LitEval-Corpus, professional raters preferred the human translation 94% of the time; student raters, 42% (Zhang et al. 2025).
LAIT's three automatic metrics (MetricX-QE, GEMBA-MQM, LiTransProQA) predicted readers at chance (0.48–0.50), and two picked the machine almost every time.
Building our own AI editors and judges, we hit the same gap from the inside. The smoothest drafts stripped the most nuance, source-aware judges forgave calques, and models ironed out deliberate quirks. It appeared that natively, LLM-as-Judge approach cannot measure translation quality.
So we set out to measure what separates AI from human translation, and what separates a good translation from a bad one. We tested two theories: first , that a text carries units of meaning and style and breaks if a translation fails to move all of them; second , that what readers actually feel is reading flow, and the places where it stumbles.
Our data combines LAIT (French, Polish, and Japanese into English, with two readers' choices per passage), three public datasets rated by translators who saw the original (ParaFarm, PAR3, WMT25 literary), and our own texts, including my novel's prologue translated into Ukrainian.
Theory one: does a translation have to carry every unit of meaning and style?
In short: No. Machines carry almost every unit of information in the original; good human translators don't, and readers don't hold it against them. Readers who never see the original punish expansion. Experts who read the original punish changed meaning.
How much of the original survives is the clearest difference between human and machine translation we measured.
Why did we expect every unit to matter?
A book invents language artifacts (names, coined terms, rhythm, running jokes) that work on several levels at once, while leaning on target-language standards (idioms, register, punctuation, syntax). Translation pulls both ways. We split text into meaning (what a reader learns) and style (how it reads), betting that dropping pieces of either layer makes a translation break.
How did we measure meaning and style?
How to read: one judge per row: what it sees, what it checks, and the number it gives.
The first two work. The meaning judge catches 51 of 61 real errors on our answer key while flagging 10 of 33 deliberate liberties in published translations. The rubric caught all 61 planted errors (57 of 59 on Gemini) and exposed what smooth prose hides: translating through English lost 12–18 elements per 1,000 words vs. 0–1 directly, with only 0–2.4 hard errors.
Neither predicts readers (rubric losses score −0.05 against LAIT). Only the rubric's source-blind native-reader check tracked them (+0.23 on tuning books): our first clue that readers react to how a text reads, not what it kept.
What do translators do with the information?
Machines keep nearly everything. Published translators keep about three units in four and rework the rest.
How to read: a longer "kept" bar means closer to the original; the other segments are what was reworked, and added units sit outside the 100. Neither end is "better" on its own.
Source: 135 LAIT passages; one Gemini 4 Hydrogen call (high thinking) per passage marks all 14 versions blind against one list of units. Compressed and changed are upper bounds (see "How far can the unit tracker be trusted?" ).
The pattern holds in every corpus we ran:
How to read: units of the original kept per 100. Higher means closer to the original, not a better translation.
How much survives is the best single tell between a human and a machine translation. On LAIT, the published translation keeps fewer units than the study's MT in 374 of 386 passages (97%) and fewer than Transept Pro in 381 (99%). Among the 14 versions in Chart 1, the human one keeps the fewest units in 115 of 135 passages.
The one machine that compresses like a human is GPT-4.1 in WMT25: it keeps 76 and writes about 12% shorter than the others, a concise style rather than truncation.
For comparison, our reading score puts the published translation above the study MT in 66% of passages (the readers: 67.6%), and the published version is longer than the MT in only 59%.
How to read: one passage, four English versions, and what the tracker marked for each.
«А ще воно відкривало її пружні плечі та жіночний живіт — під яким грали м'язи, коли Агнес сміялася чи затягувалася сигаретою. І те, й інше магеса любила.»
“ It also bared her firm shoulders and womanly belly, under which muscles played whenever Agnes laughed or drew on a cigarette. The magess loved both. ”
“ It also bared her firm shoulders and her womanly belly , where muscles played whenever Agnes laughed or drew on a cigarette. The magess loved doing both. ”
Sonnet and Haiku read almost the same. Tracker: Sonnet keeps every unit; Opus too, except that "under which" becoming "where" counts as compressed. A reader judge flagged Opus's "her womanly belly" as hard to follow.
“ The outfit bared her toned shoulders and the soft curve of her stomach , where muscle stirred beneath the skin whenever she laughed or drew on her cigarette, two pleasures the magess rarely denied herself . ”
Tracker: "womanly belly" changed, "the magess loved both" changed. No reader judge flagged these lines.
the models carry the units over word for word; my version says the same things the English way, and the tracker counts that as rework. Across this chapter my version keeps 71–77 units in 100, the three models 87–100.
Which units do translators rework?
The ones that carry voice. Facts and events stay; emphasis, images and details move.
How to read: the share of each unit type a version reworked. Higher means more reworking, not worse.
"Reworked" means compressed, changed, or dropped across Chart 1's 135 passages (Gemini 4 Hydrogen; reference and emphasis are the smallest groups, so their columns are noisiest).
How to read: further right means more of that unit type was reworked. The gap between the big published dot and the cluster of machine dots is the human–machine difference.
Source: one Gemini 4 Hydrogen call (high thinking) per passage marks all 14 versions blind against one list of units; group sizes 214 emphasis, 247 image, 994 detail, 127 reference, 902 feeling, 1,758 fact, 1,706 event, 555 logic, 205 speaker units. Shares are upper bounds (see "How far can the unit tracker be trusted?" ). The cheaper models' higher shares (Haiku, the Flash Lites, 3.8 Flash) are more likely slips than deliberate reworking.
Translators also agree on which units are open. In ParaFarm, units reworked by at least four of the six published translators appear 56 times (against 11 by chance), and un

[truncated]
