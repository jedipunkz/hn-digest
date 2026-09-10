---
source: "https://terriblesoftware.org/2026/09/10/ai-is-breaking-this-thing-we-call-trust/"
hn_url: "https://news.ycombinator.com/item?id=49644179"
title: "AI Is Breaking This Thing We Call Trust"
article_title: "AI Is Breaking This Thing We Call Trust – Terrible Software"
image: "https://i0.wp.com/terriblesoftware.org/wp-content/uploads/2026/09/cHJpdmF0ZS9sci9pbWFnZXMvd2Vic2l0ZS8yMDIyLTA1L2lzMTkzODMtaW1hZ2Uta3d2d3MwemwuanBn.webp?fit=1024%2C683&ssl=1"
author: "matheusml"
captured_at: "2026-09-10T14:52:50Z"
capture_tool: "hn-digest"
hn_id: 49644179
score: 3
comments: 1
posted_at: "2026-09-10T14:21:20Z"
tags:
  - hacker-news
---

# AI Is Breaking This Thing We Call Trust

- HN: [49644179](https://news.ycombinator.com/item?id=49644179)
- Source: [terriblesoftware.org](https://terriblesoftware.org/2026/09/10/ai-is-breaking-this-thing-we-call-trust/)
- Score: 3
- Comments: 1
- Posted: 2026-09-10T14:21:20Z

## Translation

Title: AI Is Breaking This Thing We Call Trust
Article title: AI Is Breaking This Thing We Call Trust – Terrible Software
Description: AI-generated work can look finished without being understood. That uncertainty erodes trust, slows reviews, and leaves colleagues doing the missing work.

Article text:
Skip to content
Terrible Software
AI Is Breaking This Thing We Call Trust
I want you to picture yourself in the early 1900s (in America), real quick. People have been using horses for their entire lives, and they are pros at it. Entire cities are actually built with the premise that horses are how we move from point A to B.
Then this Henry Ford guy shows up and turns the world upside down.
Over the next few years, people had to adapt to cars becoming part of everyday life. Habits built around the transportation they knew had to change.
I think that’s where we are right now. Most of us can see that AI is turning how we work upside down… But, just like in the early 1900s, we have to adapt to it. And it’s painful.
Here’s the thing: most of our habits at work still assume that producing something means you have to understand it. When you send me a PR, I assume you’ve at least read the code yourself, right? Or when you send me a brief, I assume you’ve done your research, etc.
Of course, these assumptions aren’t 100% reliable: engineers are able to copy/paste from Stack Overflow without understanding anything for years. But even in this case, it does take more effort, and reviewers are able to identify the gaps more easily.
AI makes it easy to generate code, explanations, answers to review questions, and entire docs without the author understanding any of it. Everything can look finished. Actually, more often than not, it can even be correct (!) without the author having checked it.
And that is the part that I’ve been struggling with recently, and what most of this post is about… I used to start reviewing work with the assumption that the other person did their homework. Now that may or may not be true, so instead of starting the review with a “is this a good change?” Now I often start with, “does the author understand what they’re sending me?”
All this to say that, at least right now (while we adapt), AI is breaking trust for me. And trust, especially at work, allows you to skip a bunch of steps (and save time!). When I trust a colleague, I can focus on just the parts that need another pair of eyes. I don’t have to repeat their entire investigation to feel comfortable with the output.
Sometimes the phrasing alone ( “Yes — and the real unlock is…” ) makes me suspicious. Which isn’t necessarily fair. But when I ask a specific question and get another (clearly AI-generated) explanation that doesn’t answer it, I have a reason to worry. The next thing they send me will get more scrutiny, even if they were more careful this time. A shortcut on one task can make working together harder for much longer.
Let me be clear and say that I’ve been that person myself . Yes, I merged an AI-written fix without understanding it. The fix was correct, but I hadn’t done the work to know that.
Maybe part of adapting to these tools is being more explicit about what putting our name on something means? It should mean you understand it and stand behind it. “Claude wrote it” doesn’t tell me what you checked or whether you agree with the result.
If you’re sending an early idea (or a proof of concept) that you want help with, say that! I mean, I can work with uncertainty… but it’s harder for me to work with something that’s not finished, but it’s being presented as such.
And when I ask you a question, I’m asking you a question (I have Claude and ChatGPT too, if I wanted their thoughts, I’d simply ask them). Of course, use whatever tools that might help with the answer, but take the time to decide whether the answer makes sense, or if it’s too long or if there are too many em dashes, before sending it back.
But let’s make this more actionable, with a few suggestions…
Read before sending. Review your own diff, read your own doc. Cut the explanation down to what the reviewer needs and nothing else, please! It’s totally fine to spend a bit more time trimming down and editing something so others don’t have to spend their time reading fluff.
Understand what you’re submitting. Do you really know why the bug happened and why this fix will solve it? If not, go back and understand before asking me to review it.
Push back when reviewing, too. Asking things like “Which cases did you test?” or “What led you to this recommendation?” are totally fine (and probably even expected!).
If you’re an engineering leader ,
“Ready for review” should mean it’s ready to be reviewed . Ask authors to explain what they checked and what still needs attention. But keep this proportional, because a risk-free change should not require pages and pages of evidence.
Make it safe to share unfinished work . “I haven’t verified all this yet” is useful information. A proof of concept can deserve feedback just like finished work, provided that everyone knows what they’re reviewing.
Apply the same standard to yourself . If you’re sending slop, of course your team will send each other slop. Be the role model !
Long story short, that’s the expectation I want us to keep as the tools change. Your name on the work should still mean I can trust you to understand it and stand behind it, especially before sending it to me.
Otherwise, every time you send me something, I have to do your part as well before I can do mine.
Share on Facebook (Opens in new window)
Facebook
Share on LinkedIn (Opens in new window)
LinkedIn
Share on Threads (Opens in new window)
Threads
Share on Bluesky (Opens in new window)
Bluesky
Share on Mastodon (Opens in new window)
Mastodon
Share on X (Opens in new window)
X
Loading…
opinion
ai , collaboration , teams , trust
Discover more from Terrible Software
Subscribe to get the latest posts sent to your email.
The Strange Case of Engineers Who Dismiss AI Date December 19, 2025
Making AI Actually Work on Your Team Date April 7, 2025
What Doesn’t Change Date July 14, 2025
Bluesky , GitHub , LinkedIn , X
Discover more from Terrible Software
Subscribe now to keep reading and get access to the full archive.

## Original Extract

AI-generated work can look finished without being understood. That uncertainty erodes trust, slows reviews, and leaves colleagues doing the missing work.

Skip to content
Terrible Software
AI Is Breaking This Thing We Call Trust
I want you to picture yourself in the early 1900s (in America), real quick. People have been using horses for their entire lives, and they are pros at it. Entire cities are actually built with the premise that horses are how we move from point A to B.
Then this Henry Ford guy shows up and turns the world upside down.
Over the next few years, people had to adapt to cars becoming part of everyday life. Habits built around the transportation they knew had to change.
I think that’s where we are right now. Most of us can see that AI is turning how we work upside down… But, just like in the early 1900s, we have to adapt to it. And it’s painful.
Here’s the thing: most of our habits at work still assume that producing something means you have to understand it. When you send me a PR, I assume you’ve at least read the code yourself, right? Or when you send me a brief, I assume you’ve done your research, etc.
Of course, these assumptions aren’t 100% reliable: engineers are able to copy/paste from Stack Overflow without understanding anything for years. But even in this case, it does take more effort, and reviewers are able to identify the gaps more easily.
AI makes it easy to generate code, explanations, answers to review questions, and entire docs without the author understanding any of it. Everything can look finished. Actually, more often than not, it can even be correct (!) without the author having checked it.
And that is the part that I’ve been struggling with recently, and what most of this post is about… I used to start reviewing work with the assumption that the other person did their homework. Now that may or may not be true, so instead of starting the review with a “is this a good change?” Now I often start with, “does the author understand what they’re sending me?”
All this to say that, at least right now (while we adapt), AI is breaking trust for me. And trust, especially at work, allows you to skip a bunch of steps (and save time!). When I trust a colleague, I can focus on just the parts that need another pair of eyes. I don’t have to repeat their entire investigation to feel comfortable with the output.
Sometimes the phrasing alone ( “Yes — and the real unlock is…” ) makes me suspicious. Which isn’t necessarily fair. But when I ask a specific question and get another (clearly AI-generated) explanation that doesn’t answer it, I have a reason to worry. The next thing they send me will get more scrutiny, even if they were more careful this time. A shortcut on one task can make working together harder for much longer.
Let me be clear and say that I’ve been that person myself . Yes, I merged an AI-written fix without understanding it. The fix was correct, but I hadn’t done the work to know that.
Maybe part of adapting to these tools is being more explicit about what putting our name on something means? It should mean you understand it and stand behind it. “Claude wrote it” doesn’t tell me what you checked or whether you agree with the result.
If you’re sending an early idea (or a proof of concept) that you want help with, say that! I mean, I can work with uncertainty… but it’s harder for me to work with something that’s not finished, but it’s being presented as such.
And when I ask you a question, I’m asking you a question (I have Claude and ChatGPT too, if I wanted their thoughts, I’d simply ask them). Of course, use whatever tools that might help with the answer, but take the time to decide whether the answer makes sense, or if it’s too long or if there are too many em dashes, before sending it back.
But let’s make this more actionable, with a few suggestions…
Read before sending. Review your own diff, read your own doc. Cut the explanation down to what the reviewer needs and nothing else, please! It’s totally fine to spend a bit more time trimming down and editing something so others don’t have to spend their time reading fluff.
Understand what you’re submitting. Do you really know why the bug happened and why this fix will solve it? If not, go back and understand before asking me to review it.
Push back when reviewing, too. Asking things like “Which cases did you test?” or “What led you to this recommendation?” are totally fine (and probably even expected!).
If you’re an engineering leader ,
“Ready for review” should mean it’s ready to be reviewed . Ask authors to explain what they checked and what still needs attention. But keep this proportional, because a risk-free change should not require pages and pages of evidence.
Make it safe to share unfinished work . “I haven’t verified all this yet” is useful information. A proof of concept can deserve feedback just like finished work, provided that everyone knows what they’re reviewing.
Apply the same standard to yourself . If you’re sending slop, of course your team will send each other slop. Be the role model !
Long story short, that’s the expectation I want us to keep as the tools change. Your name on the work should still mean I can trust you to understand it and stand behind it, especially before sending it to me.
Otherwise, every time you send me something, I have to do your part as well before I can do mine.
Share on Facebook (Opens in new window)
Facebook
Share on LinkedIn (Opens in new window)
LinkedIn
Share on Threads (Opens in new window)
Threads
Share on Bluesky (Opens in new window)
Bluesky
Share on Mastodon (Opens in new window)
Mastodon
Share on X (Opens in new window)
X
Loading…
opinion
ai , collaboration , teams , trust
Discover more from Terrible Software
Subscribe to get the latest posts sent to your email.
The Strange Case of Engineers Who Dismiss AI Date December 19, 2025
Making AI Actually Work on Your Team Date April 7, 2025
What Doesn’t Change Date July 14, 2025
Bluesky , GitHub , LinkedIn , X
Discover more from Terrible Software
Subscribe now to keep reading and get access to the full archive.
