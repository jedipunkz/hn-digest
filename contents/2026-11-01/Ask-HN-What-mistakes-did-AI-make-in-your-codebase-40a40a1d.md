---
source: ""
hn_url: "https://news.ycombinator.com/item?id=50328797"
title: "Ask HN: What mistakes did AI make in your codebase?"
article_title: ""
image: ""
author: "underdeserver"
captured_at: "2026-11-01T20:15:36Z"
capture_tool: "hn-digest"
hn_id: 50328797
score: 1
comments: 0
posted_at: "2026-11-01T19:25:25Z"
tags:
  - hacker-news
---

# Ask HN: What mistakes did AI make in your codebase?

- HN: [50328797](https://news.ycombinator.com/item?id=50328797)
- Score: 1
- Comments: 0
- Posted: 2026-11-01T19:25:25Z

## Translation

Title: Ask HN: What mistakes did AI make in your codebase?
HN text: In the past few weeks I've beeing working with frontier AI - Opus 5.5, Astra, Sol 6.1, Fable 5.1 and Gemini Argon. AI has become really good at coding in some sense, that's undeniable, and there are people that say you don't really need to review their code anymore. But I keep getting stuck in a sensation that's between cognitive dissonance, FOMO, and a vague sense of being gaslit, because when I review the code it writes - and this is all of them - 80% of it is right, and 20% is wrong, but in extremely subtle ways; only when you deeply understand the system and carefully review it do you spot the issues. These are not stylistic choices, the code is incorrect. Of course, the AI fixes it when you prompt it to, but you need to work very hard to know what needs prompting. This is all after the grill-me skills and cross-model adversarial review and whatnot (all these techniques do help somewhat but they didn't get it perfect), For example, I had AI (Opus 5.5 and Gemini both tried, in this particular case) write an integration test for our system, to reproduce an issue we had. It failed because it chose, for simplicity, to combine two backend fakes into one store, having writes covered by the same mutex, but the bug was only happening in production when these two stores went out of sync (and they did because we wrote to both concurrently and one write failed). (I don't want to argue the merits of the design. This is existing code.) The wrong code it wrote to try and find the issue was carefully written, well organized, function names were reasonable, etc. etc. I can't be alone in this, but I can't put my finger on what exactly goes wrong in the process. Can you share examples of changes you wanted to make that Claude and the bunch (frontier versions) got subtly wrong?

## Original Extract

In the past few weeks I've beeing working with frontier AI - Opus 5.5, Astra, Sol 6.1, Fable 5.1 and Gemini Argon. AI has become really good at coding in some sense, that's undeniable, and there are people that say you don't really need to review their code anymore. But I keep getting stuck in a sensation that's between cognitive dissonance, FOMO, and a vague sense of being gaslit, because when I review the code it writes - and this is all of them - 80% of it is right, and 20% is wrong, but in extremely subtle ways; only when you deeply understand the system and carefully review it do you spot the issues. These are not stylistic choices, the code is incorrect. Of course, the AI fixes it when you prompt it to, but you need to work very hard to know what needs prompting. This is all after the grill-me skills and cross-model adversarial review and whatnot (all these techniques do help somewhat but they didn't get it perfect), For example, I had AI (Opus 5.5 and Gemini both tried, in this particular case) write an integration test for our system, to reproduce an issue we had. It failed because it chose, for simplicity, to combine two backend fakes into one store, having writes covered by the same mutex, but the bug was only happening in production when these two stores went out of sync (and they did because we wrote to both concurrently and one write failed). (I don't want to argue the merits of the design. This is existing code.) The wrong code it wrote to try and find the issue was carefully written, well organized, function names were reasonable, etc. etc. I can't be alone in this, but I can't put my finger on what exactly goes wrong in the process. Can you share examples of changes you wanted to make that Claude and the bunch (frontier versions) got subtly wrong?

