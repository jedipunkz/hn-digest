---
source: "https://morethanamachine.com/posts/handy-context/"
hn_url: "https://news.ycombinator.com/item?id=50417052"
title: "Offline context-aware dictation without screenshots, LLMs or fine-tuning"
article_title: "Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning | More Than a Machine"
image: "https://morethanamachine.com/_astro/handy-context-og.BUn5bGfs.png"
author: "reddmachine"
captured_at: "2026-11-08T21:44:51Z"
capture_tool: "hn-digest"
hn_id: 50417052
score: 1
comments: 0
posted_at: "2026-11-08T20:47:19Z"
tags:
  - hacker-news
---

# Offline context-aware dictation without screenshots, LLMs or fine-tuning

- HN: [50417052](https://news.ycombinator.com/item?id=50417052)
- Source: [morethanamachine.com](https://morethanamachine.com/posts/handy-context/)
- Score: 1
- Comments: 0
- Posted: 2026-11-08T20:47:19Z

## Translation

Title: Offline context-aware dictation without screenshots, LLMs or fine-tuning
Article title: Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning | More Than a Machine
Description: Wispr Flow reads names off your screen in its cloud. I forked Handy, an existing open-source dictation app, and taught it to do the same on my laptop. It got 35% fewer on-screen words wrong.

Article text:
Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning | More Than a Machine ~/ more-than-a-machine [home] [about] [posts] [contribute] [pitch] [rss] Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning
By Nishaanth Reddy · November 8, 2026
Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning
Playing Contra with decision models
Twine: Coding Agents That Challenge Each Other
StonkTrace: 100 Days of Reddit Stock Chatter: Hype, Hindsight and How LLMs Pick Stock
BubblegumTree: A Family Story Press for Kobo
Containing Coding Agents: Why I Run Them in a Jail
Ubiklaw: An Old Android Tablet Became My Local Home Assistant
Intro to Wafer Scale Compute: The Tokens Per Second Race
An Intro to Speculative Decoding
Discover Agents, Not Tools: MCP at the Wrong Layer
Anthropic Academy Courses: Reviews and Recommendations
Wispr Flow’s best feature is that it spells names right, because it reads them off your screen . It also sends your audio and that text to its cloud. Unlimited dictation on Pro costs $15 a month, or $12 billed annually (as of 8 Nov 2026).
I didn’t want to pay for that when the pieces are already free. Handy is an open-source dictation app by cjpais that runs offline. Its recommended model, Parakeet , has open weights. The operating system already lets an app with accessibility permission read the text on screen.
So I forked Handy and taught my copy to read the front window while I dictate. I called it Handy Context . The app is cjpais’s work. Mine is the screen reading and the transcription pipeline behind it. I measured it against VoiceInk , an open-source app that fixes transcripts by swapping in close matches from a word list.
Here’s the same audio with and without the slide’s text.
“Queue” and “cue” sound the same, so only the slide can settle it. The same goes for jargon like “gut microbiota” and names like Proscia. The last clip is the failure I worried about most. The slide overwrote a word the speaker said correctly. A lot of the work went into preventing that.
There’s no new model. Handy Context changes what happens around Parakeet.
Handy Context audio · custom words · screen text
The screen text never goes into the model. It only changes which word sequences the decoder tries and which one gets picked. Handy Context runs the exact checkpoint NVIDIA published (the file hash matches Hugging Face), with nothing fine-tuned. Handy uses an 8-bit copy of the same weights, which scored within 0.002 points of the original on LibriSpeech.
Each candidate gets Parakeet’s own score for the audio plus a small bonus for every screen word in it. So a screen word can win a close call, but it can’t override audio that clearly says something else. VoiceInk swaps any word that looks close to a list word without checking.
Real slides needed two extra rules. Matching ignores plurals and spelling variants, so a slide that says “Careers” doesn’t turn a spoken “career” into “Careers”. No fix can change only a word’s ending either.
Nothing leaves the laptop. It needs the same microphone and accessibility permissions as Wispr Flow. It reads text, so it never captures the screen. Words from Handy’s custom-words list always go to the front of the list.
I wrote down what I expected to see before downloading the test data, then ran the test once. SlideSpeech is a set of recorded talks with their slides. The word list comes from the slide on screen, or from the whole talk’s slides to simulate a busy screen.
With the screen, Parakeet gets about a third fewer on-screen words wrong and about two thirds fewer rare words on the benchmark. The cost is small. The overall error rate goes up 0.04 points with one slide and 0.11 with a whole talk, and barely moves when the screen is unrelated. It does up to nine times the decoding work, which section 04 covers.
Handy Context got 35% fewer on-screen words wrong than plain Parakeet and 26% fewer than VoiceInk’s method. VoiceInk’s method also raised the overall error rate as the screen filled up, from Parakeet’s 8.76% to 10.70% with a whole talk.
One of my own checks failed. With a whole talk on screen, Handy Context’s overall error rate was 0.11 points above Parakeet’s. The upper end of its confidence interval passed the +0.15 limit I’d set.
VoiceInk ran locally and unmodified, without its optional cloud LLM step. “Parakeet alone” means the same checkpoint decoded once with no screen text.
The gap is wider on the standard benchmark for this, Meta’s LibriSpeech biasing lists . They hide each sentence’s rare words among up to 1,000 distractors.
These numbers are from the core pipeline, before the two extra rules for real slides. All four claims I made for this test passed.
I expected to train my way to this and spent a good part of a day trying. None of it beat the untrained pipeline.
First I tried a pointer layer on top of the frozen model that learns to copy list words (TCPGen-style). The first version got far more list words wrong than plain Parakeet. The second was better but still trailed untrained boosting and hurt accuracy when the list was unrelated.
Then I trained small keyword spotters that guess which screen words were actually spoken. They worked, cutting a 1,000-word list to about 3 words while keeping 98% of the right ones. The untrained pipeline still scored a bit better, so they didn’t ship. They’re my best bet for busy screens.
A small model to pick among the candidate decodes came out too cautious or too eager, depending on what I trained it on. I never fine-tuned Parakeet itself.
Giving the word list to a speech LLM didn’t help. Qwen3-ASR-1.7B with 1,000 words missed list words as often as with none and ran about 10× slower. A vocabulary prompt made Whisper large-v3-turbo worse overall, from 3.53% to 4.8–5.6% WER.
VoiceInk-style swapping raised the overall error rate on real slides from 8.25% to 8.79%. About half of its bad swaps only changed a word’s ending, like “career” to “Careers”. That’s where the two extra rules came from.
A 2–14 second dictation comes back in 0.2–0.5 s. Memory is the real cost. The server runs Parakeet at full precision in PyTorch, while Handy runs an 8-bit build of the same model. About 2.7 GiB of the total is PyTorch’s GPU cache. Freeing it after each dictation, and loading Handy’s fallback model only when needed, should get it under 4 GiB. Porting the pipeline into Handy’s own engine would bring it back to Handy’s size.
It’s English only and runs on one desktop OS so far. There’s no OCR yet, so apps that draw their text as images give it nothing to read. Busy screens are the main open problem. Ranking terms by how close they are to the cursor is next.
The code is private, so this is the full recipe. Every setting below is what the app runs.
When the shortcut goes down, a helper walks the front window’s accessibility tree breadth first. It collects text from static text, text fields and areas, headings, links, table cells and rows, outlines, buttons, menu items, tabs, groups and web areas. It stops at 4,000 elements, 30,000 characters or 350 ms, whichever comes first. It also reads the window title, the selected text and the last 8,000 characters of the focused field. Chromium and Electron apps only expose their tree after the helper asks them to. When the shortcut comes up, the app waits at most 600 ms for the capture.
The title and screen text are split into words. Hyphenated words are split into parts. Apostrophes are trimmed from the ends. A word becomes a term if it’s alphabetic, at least 3 letters long and not among the 5,000 common words in Meta’s is21_deep_bias word list. Terms keep the order they first appear in, up to 1,000. Custom words from Handy’s settings go first. A single-word custom entry skips the common-word filter.
Each term keeps its most frequent on-screen spelling. An all-caps word becomes lowercase if it’s longer than 5 letters or in the LibriSpeech vocabulary, so short acronyms stay in capitals.
For scoring, the list is expanded with each term’s inflections that appear in the LibriSpeech vocabulary. That covers plural and singular forms, possessives and the spelling pairs -ise/-ize, -isation/-ization, -our/-or, -re/-er and -mme/-m. Possessives always count.
The audio is encoded once. Greedy RNN-T decoding, at most 10 symbols per frame, then runs up to nine times.
Biasing uses NeMo’s prefix-tree booster ( GPUBoostingTreeModel ) with context score 1.0, depth scaling 2.0 and unknown score 0. Its phrases are each term and its capitalized form. At a frame where the model’s top choice is a token rather than blank, α times the booster’s score is added to every token’s log-probability and blank is ruled out. Frames where blank is the top choice are left alone. A gated decode also leaves a frame alone when its top token holds at least τ of the probability among non-blank tokens.
A two-pass list keeps the terms that resemble something in the plain transcript. Each term is compared with every 1-, 2- and 3-word span of the plain text, joined without spaces and at least 3 letters long. Spans whose length differs from the term’s by more than 3 letters, or by more than half the term’s length if that’s larger, are skipped. Similarity is 1 minus the Levenshtein distance divided by the longer length. It’s computed on the letters and on the Metaphone codes. The higher of the two counts. Terms scoring at least θ are kept, the top k by score. A two-pass decode is skipped when its list comes out empty.
Duplicate candidates are merged. Each remaining candidate gets its exact log-likelihood (natural log) under the unbiased model, computed with the RNN-T forward algorithm, plus 8 for every word in it that’s on the expanded term list. The highest score wins.
Swaps are proposed from any 1-, 2- or 3-word span of the winning text to a term on the expanded list. They use the same similarity and length rule with a threshold of 0.6. The 30 most similar are kept. A single-word swap is dropped if it only changes an ending, meaning the two words share a prefix of at least 5 letters that reaches within 3 letters of the end of the shorter one.
Each remaining swap scores the change in log-likelihood plus 8 times the change in term count. Swaps with a positive score are applied from highest to lowest, skipping any that overlap one already applied. A swap keeps the capitalization of the text it replaces. Finally, any word matching a term with capitals after its first letter, like PyTorch or NASA, takes the on-screen spelling. Handy’s usual filler-word cleanup runs last.
I wrote down the claims and pass bars before downloading each test set, then ran each test set once. All the settings above were chosen on development data only. That was two LibriSpeech dev splits with Meta-style lists, plus SlideSpeech dev (21 talks, 1,798 segments, 4.56 hours).
SlideSpeech test has 25 talks, 3,176 segments and 7.95 hours. Its lists come from the dataset’s slide OCR and KeyBERT keywords for each segment, filtered like the term list. The four conditions are the segment’s keywords, its current slide, every slide in the talk and a slide from another talk. On-screen words are the reference words that appear in the segment’s own slide or keyword list.
LibriSpeech test-clean (2,618 utterances) and test-other (2,939) use Meta’s official lists at 100, 500 and 1,000 words. The null list is the 100-word list minus the utterance’s own rare words. Two utterances I’d looked at during early exploration were excluded. These runs used the settings above without the inflection expansion and the ending rule, which came later.
List-word error rate follows Meta’s benchmark. It counts substitutions and deletions of list words in the reference, plus inserted list words, over the number of list words in the reference. Text is normalized to Meta’s format. Intervals are 95%, from a 10,000-resample paired bootstrap over utterances.
The VoiceInk baseline passes the

[truncated]

## Original Extract

Wispr Flow reads names off your screen in its cloud. I forked Handy, an existing open-source dictation app, and taught it to do the same on my laptop. It got 35% fewer on-screen words wrong.

Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning | More Than a Machine ~/ more-than-a-machine [home] [about] [posts] [contribute] [pitch] [rss] Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning
By Nishaanth Reddy · November 8, 2026
Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning
Playing Contra with decision models
Twine: Coding Agents That Challenge Each Other
StonkTrace: 100 Days of Reddit Stock Chatter: Hype, Hindsight and How LLMs Pick Stock
BubblegumTree: A Family Story Press for Kobo
Containing Coding Agents: Why I Run Them in a Jail
Ubiklaw: An Old Android Tablet Became My Local Home Assistant
Intro to Wafer Scale Compute: The Tokens Per Second Race
An Intro to Speculative Decoding
Discover Agents, Not Tools: MCP at the Wrong Layer
Anthropic Academy Courses: Reviews and Recommendations
Wispr Flow’s best feature is that it spells names right, because it reads them off your screen . It also sends your audio and that text to its cloud. Unlimited dictation on Pro costs $15 a month, or $12 billed annually (as of 8 Nov 2026).
I didn’t want to pay for that when the pieces are already free. Handy is an open-source dictation app by cjpais that runs offline. Its recommended model, Parakeet , has open weights. The operating system already lets an app with accessibility permission read the text on screen.
So I forked Handy and taught my copy to read the front window while I dictate. I called it Handy Context . The app is cjpais’s work. Mine is the screen reading and the transcription pipeline behind it. I measured it against VoiceInk , an open-source app that fixes transcripts by swapping in close matches from a word list.
Here’s the same audio with and without the slide’s text.
“Queue” and “cue” sound the same, so only the slide can settle it. The same goes for jargon like “gut microbiota” and names like Proscia. The last clip is the failure I worried about most. The slide overwrote a word the speaker said correctly. A lot of the work went into preventing that.
There’s no new model. Handy Context changes what happens around Parakeet.
Handy Context audio · custom words · screen text
The screen text never goes into the model. It only changes which word sequences the decoder tries and which one gets picked. Handy Context runs the exact checkpoint NVIDIA published (the file hash matches Hugging Face), with nothing fine-tuned. Handy uses an 8-bit copy of the same weights, which scored within 0.002 points of the original on LibriSpeech.
Each candidate gets Parakeet’s own score for the audio plus a small bonus for every screen word in it. So a screen word can win a close call, but it can’t override audio that clearly says something else. VoiceInk swaps any word that looks close to a list word without checking.
Real slides needed two extra rules. Matching ignores plurals and spelling variants, so a slide that says “Careers” doesn’t turn a spoken “career” into “Careers”. No fix can change only a word’s ending either.
Nothing leaves the laptop. It needs the same microphone and accessibility permissions as Wispr Flow. It reads text, so it never captures the screen. Words from Handy’s custom-words list always go to the front of the list.
I wrote down what I expected to see before downloading the test data, then ran the test once. SlideSpeech is a set of recorded talks with their slides. The word list comes from the slide on screen, or from the whole talk’s slides to simulate a busy screen.
With the screen, Parakeet gets about a third fewer on-screen words wrong and about two thirds fewer rare words on the benchmark. The cost is small. The overall error rate goes up 0.04 points with one slide and 0.11 with a whole talk, and barely moves when the screen is unrelated. It does up to nine times the decoding work, which section 04 covers.
Handy Context got 35% fewer on-screen words wrong than plain Parakeet and 26% fewer than VoiceInk’s method. VoiceInk’s method also raised the overall error rate as the screen filled up, from Parakeet’s 8.76% to 10.70% with a whole talk.
One of my own checks failed. With a whole talk on screen, Handy Context’s overall error rate was 0.11 points above Parakeet’s. The upper end of its confidence interval passed the +0.15 limit I’d set.
VoiceInk ran locally and unmodified, without its optional cloud LLM step. “Parakeet alone” means the same checkpoint decoded once with no screen text.
The gap is wider on the standard benchmark for this, Meta’s LibriSpeech biasing lists . They hide each sentence’s rare words among up to 1,000 distractors.
These numbers are from the core pipeline, before the two extra rules for real slides. All four claims I made for this test passed.
I expected to train my way to this and spent a good part of a day trying. None of it beat the untrained pipeline.
First I tried a pointer layer on top of the frozen model that learns to copy list words (TCPGen-style). The first version got far more list words wrong than plain Parakeet. The second was better but still trailed untrained boosting and hurt accuracy when the list was unrelated.
Then I trained small keyword spotters that guess which screen words were actually spoken. They worked, cutting a 1,000-word list to about 3 words while keeping 98% of the right ones. The untrained pipeline still scored a bit better, so they didn’t ship. They’re my best bet for busy screens.
A small model to pick among the candidate decodes came out too cautious or too eager, depending on what I trained it on. I never fine-tuned Parakeet itself.
Giving the word list to a speech LLM didn’t help. Qwen3-ASR-1.7B with 1,000 words missed list words as often as with none and ran about 10× slower. A vocabulary prompt made Whisper large-v3-turbo worse overall, from 3.53% to 4.8–5.6% WER.
VoiceInk-style swapping raised the overall error rate on real slides from 8.25% to 8.79%. About half of its bad swaps only changed a word’s ending, like “career” to “Careers”. That’s where the two extra rules came from.
A 2–14 second dictation comes back in 0.2–0.5 s. Memory is the real cost. The server runs Parakeet at full precision in PyTorch, while Handy runs an 8-bit build of the same model. About 2.7 GiB of the total is PyTorch’s GPU cache. Freeing it after each dictation, and loading Handy’s fallback model only when needed, should get it under 4 GiB. Porting the pipeline into Handy’s own engine would bring it back to Handy’s size.
It’s English only and runs on one desktop OS so far. There’s no OCR yet, so apps that draw their text as images give it nothing to read. Busy screens are the main open problem. Ranking terms by how close they are to the cursor is next.
The code is private, so this is the full recipe. Every setting below is what the app runs.
When the shortcut goes down, a helper walks the front window’s accessibility tree breadth first. It collects text from static text, text fields and areas, headings, links, table cells and rows, outlines, buttons, menu items, tabs, groups and web areas. It stops at 4,000 elements, 30,000 characters or 350 ms, whichever comes first. It also reads the window title, the selected text and the last 8,000 characters of the focused field. Chromium and Electron apps only expose their tree after the helper asks them to. When the shortcut comes up, the app waits at most 600 ms for the capture.
The title and screen text are split into words. Hyphenated words are split into parts. Apostrophes are trimmed from the ends. A word becomes a term if it’s alphabetic, at least 3 letters long and not among the 5,000 common words in Meta’s is21_deep_bias word list. Terms keep the order they first appear in, up to 1,000. Custom words from Handy’s settings go first. A single-word custom entry skips the common-word filter.
Each term keeps its most frequent on-screen spelling. An all-caps word becomes lowercase if it’s longer than 5 letters or in the LibriSpeech vocabulary, so short acronyms stay in capitals.
For scoring, the list is expanded with each term’s inflections that appear in the LibriSpeech vocabulary. That covers plural and singular forms, possessives and the spelling pairs -ise/-ize, -isation/-ization, -our/-or, -re/-er and -mme/-m. Possessives always count.
The audio is encoded once. Greedy RNN-T decoding, at most 10 symbols per frame, then runs up to nine times.
Biasing uses NeMo’s prefix-tree booster ( GPUBoostingTreeModel ) with context score 1.0, depth scaling 2.0 and unknown score 0. Its phrases are each term and its capitalized form. At a frame where the model’s top choice is a token rather than blank, α times the booster’s score is added to every token’s log-probability and blank is ruled out. Frames where blank is the top choice are left alone. A gated decode also leaves a frame alone when its top token holds at least τ of the probability among non-blank tokens.
A two-pass list keeps the terms that resemble something in the plain transcript. Each term is compared with every 1-, 2- and 3-word span of the plain text, joined without spaces and at least 3 letters long. Spans whose length differs from the term’s by more than 3 letters, or by more than half the term’s length if that’s larger, are skipped. Similarity is 1 minus the Levenshtein distance divided by the longer length. It’s computed on the letters and on the Metaphone codes. The higher of the two counts. Terms scoring at least θ are kept, the top k by score. A two-pass decode is skipped when its list comes out empty.
Duplicate candidates are merged. Each remaining candidate gets its exact log-likelihood (natural log) under the unbiased model, computed with the RNN-T forward algorithm, plus 8 for every word in it that’s on the expanded term list. The highest score wins.
Swaps are proposed from any 1-, 2- or 3-word span of the winning text to a term on the expanded list. They use the same similarity and length rule with a threshold of 0.6. The 30 most similar are kept. A single-word swap is dropped if it only changes an ending, meaning the two words share a prefix of at least 5 letters that reaches within 3 letters of the end of the shorter one.
Each remaining swap scores the change in log-likelihood plus 8 times the change in term count. Swaps with a positive score are applied from highest to lowest, skipping any that overlap one already applied. A swap keeps the capitalization of the text it replaces. Finally, any word matching a term with capitals after its first letter, like PyTorch or NASA, takes the on-screen spelling. Handy’s usual filler-word cleanup runs last.
I wrote down the claims and pass bars before downloading each test set, then ran each test set once. All the settings above were chosen on development data only. That was two LibriSpeech dev splits with Meta-style lists, plus SlideSpeech dev (21 talks, 1,798 segments, 4.56 hours).
SlideSpeech test has 25 talks, 3,176 segments and 7.95 hours. Its lists come from the dataset’s slide OCR and KeyBERT keywords for each segment, filtered like the term list. The four conditions are the segment’s keywords, its current slide, every slide in the talk and a slide from another talk. On-screen words are the reference words that appear in the segment’s own slide or keyword list.
LibriSpeech test-clean (2,618 utterances) and test-other (2,939) use Meta’s official lists at 100, 500 and 1,000 words. The null list is the 100-word list minus the utterance’s own rare words. Two utterances I’d looked at during early exploration were excluded. These runs used the settings above without the inflection expansion and the ending rule, which came later.
List-word error rate follows Meta’s benchmark. It counts substitutions and deletions of list words in the reference, plus inserted list words, over the number of list words in the reference. Text is normalized to Meta’s format. Intervals are 95%, from a 10,000-resample paired bootstrap over utterances.
The VoiceInk baseline passes the

[truncated]
