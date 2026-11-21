---
source: "https://bfontaine.net/blog/2026/11/21/llms-and-the-illusion-of-competence/"
hn_url: "https://news.ycombinator.com/item?id=50588119"
title: "LLMs and the Illusion of Competence"
article_title: "LLMs and the illusion of competence · Baptiste Fontaine’s Blog"
image: ""
author: "hk__2"
captured_at: "2026-11-21T12:11:04Z"
capture_tool: "hn-digest"
hn_id: 50588119
score: 1
comments: 0
posted_at: "2026-11-21T11:25:08Z"
tags:
  - hacker-news
---

# LLMs and the Illusion of Competence

- HN: [50588119](https://news.ycombinator.com/item?id=50588119)
- Source: [bfontaine.net](https://bfontaine.net/blog/2026/11/21/llms-and-the-illusion-of-competence/)
- Score: 1
- Comments: 0
- Posted: 2026-11-21T11:25:08Z

## Translation

Title: LLMs and the Illusion of Competence
Article title: LLMs and the illusion of competence · Baptiste Fontaine’s Blog
Description: Tips and tricks in Python, Bash, Clojure and other languages/environments.

Article text:
Baptiste Fontaine’s Blog
( back to the website )
LLMs and the illusion of competence
One thing that I worry about with coding in the age of AI is that the uncritical use of LLMs can make us believe we are
far more competent than we really are. In general, the more expert you become in your field, the better you can judge
what someone (or an LLM) says about it, and conversely, the less you know in a field, the less likely you are to
question what people say about it 1 .
This happens to all of us, sometimes in comical ways: while reading a newspaper, you may come across an article
about some technical topic you already know really well. Reading the article, you’ll likely notice things that are
off, oversimplified, dubious. You’ll think “this journalist doesn’t know what they’re talking about”. And then you’ll
turn the page and continue reading the rest of the newspaper as if nothing happened. In a speech in 2002, Michael
Crichton called this the “ Gell-Mann amnesia effect ”.
I see this with my use of LLMs, too: I’m very critical of what they produce for coding or what they say about
architectural decisions, but then when I ask them about the best grammar book to learn Romanian or how to deal with real
estate agencies in Italy, I tend to easily trust them without verifying too much, making excuses for myself like “it’s
just to check”, “this is common knowledge, it can’t really be off”.
The problem with this is that if you are not aware of the phenomenon, you can easily get overconfident in your ability
to produce good work, especially as LLMs are notoriously bad at judging some very important aspects of software
development, such as maintainability or the architecture of large projects. You can make a shiny website that looks
professional but is a mountain of spaghetti code under the hood, full of security holes and hidden bugs. But if
you don’t have the actual competence to judge these problems, you won’t see them and will focus on the appearance,
thinking that with your “friend Claude” you can do the same as more experienced developers in a fraction of the time,
wondering why they take so much time to do the same.
If you work alone, it’s a latent problem because you might become overconfident and then be hit hard by reality when
you apply for a job with technical tests. If you work with others, it’s already a problem because you can degrade the
quality of the codebase by unknowingly introducing the overly complex code LLMs typically produce, just because the LLM
convinced you it was necessary. At $work, I’m sometimes bitten by this: I’m not as critical of what the LLM says about
potential bugs in applications I’m not familiar with as I am of what it says about the code in systems I maintain every
day. I might send a PR to my colleague with what looks to me like a sound fix to the issue the LLM described to me, only
to see my colleague dismiss it, saying it’s garbage and the LLM doesn’t know what it’s talking about.
This is a phenomenon we can’t fix, but can mitigate if we keep it in mind when we work with things we’re not familiar
with. For example, I now force myself to make PRs on systems I don’t know well to have someone review them instead of
trusting the LLM too much.
This is usually described as the Dunning–Kruger effect . ↩

## Original Extract

Tips and tricks in Python, Bash, Clojure and other languages/environments.

Baptiste Fontaine’s Blog
( back to the website )
LLMs and the illusion of competence
One thing that I worry about with coding in the age of AI is that the uncritical use of LLMs can make us believe we are
far more competent than we really are. In general, the more expert you become in your field, the better you can judge
what someone (or an LLM) says about it, and conversely, the less you know in a field, the less likely you are to
question what people say about it 1 .
This happens to all of us, sometimes in comical ways: while reading a newspaper, you may come across an article
about some technical topic you already know really well. Reading the article, you’ll likely notice things that are
off, oversimplified, dubious. You’ll think “this journalist doesn’t know what they’re talking about”. And then you’ll
turn the page and continue reading the rest of the newspaper as if nothing happened. In a speech in 2002, Michael
Crichton called this the “ Gell-Mann amnesia effect ”.
I see this with my use of LLMs, too: I’m very critical of what they produce for coding or what they say about
architectural decisions, but then when I ask them about the best grammar book to learn Romanian or how to deal with real
estate agencies in Italy, I tend to easily trust them without verifying too much, making excuses for myself like “it’s
just to check”, “this is common knowledge, it can’t really be off”.
The problem with this is that if you are not aware of the phenomenon, you can easily get overconfident in your ability
to produce good work, especially as LLMs are notoriously bad at judging some very important aspects of software
development, such as maintainability or the architecture of large projects. You can make a shiny website that looks
professional but is a mountain of spaghetti code under the hood, full of security holes and hidden bugs. But if
you don’t have the actual competence to judge these problems, you won’t see them and will focus on the appearance,
thinking that with your “friend Claude” you can do the same as more experienced developers in a fraction of the time,
wondering why they take so much time to do the same.
If you work alone, it’s a latent problem because you might become overconfident and then be hit hard by reality when
you apply for a job with technical tests. If you work with others, it’s already a problem because you can degrade the
quality of the codebase by unknowingly introducing the overly complex code LLMs typically produce, just because the LLM
convinced you it was necessary. At $work, I’m sometimes bitten by this: I’m not as critical of what the LLM says about
potential bugs in applications I’m not familiar with as I am of what it says about the code in systems I maintain every
day. I might send a PR to my colleague with what looks to me like a sound fix to the issue the LLM described to me, only
to see my colleague dismiss it, saying it’s garbage and the LLM doesn’t know what it’s talking about.
This is a phenomenon we can’t fix, but can mitigate if we keep it in mind when we work with things we’re not familiar
with. For example, I now force myself to make PRs on systems I don’t know well to have someone review them instead of
trusting the LLM too much.
This is usually described as the Dunning–Kruger effect . ↩
