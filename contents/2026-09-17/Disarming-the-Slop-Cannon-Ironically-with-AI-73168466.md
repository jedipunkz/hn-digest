---
source: "https://blog.expectedparrot.com/p/disarming-the-slop-cannon-ironically"
hn_url: "https://news.ycombinator.com/item?id=49742333"
title: "Disarming the Slop Cannon (Ironically, with AI)"
article_title: "Disarming the Slop Cannon (ironically, with AI)"
image: "https://substackcdn.com/image/fetch/$s_!kW41!,w_1200,h_600,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Ff3cd6e11-8c65-47c9-b260-f4a1257ccf9c_1118x590.png"
author: "john_horton"
captured_at: "2026-09-17T16:23:23Z"
capture_tool: "hn-digest"
hn_id: 49742333
score: 2
comments: 1
posted_at: "2026-09-17T15:32:10Z"
tags:
  - hacker-news
---

# Disarming the Slop Cannon (Ironically, with AI)

- HN: [49742333](https://news.ycombinator.com/item?id=49742333)
- Source: [blog.expectedparrot.com](https://blog.expectedparrot.com/p/disarming-the-slop-cannon-ironically)
- Score: 2
- Comments: 1
- Posted: 2026-09-17T15:32:10Z

## Translation

Title: Disarming the Slop Cannon (Ironically, with AI)
Article title: Disarming the Slop Cannon (ironically, with AI)
Description: Using Expected Parrot to spin up bespoke AI interview + analysis workflows to flag “authors” who can’t explain their own work

Article text:
Disarming the Slop Cannon (ironically, with AI)
Expected Parrot
Subscribe Sign in Disarming the Slop Cannon (ironically, with AI)
Using Expected Parrot to spin up bespoke AI interview + analysis workflows to flag “authors” who can’t explain their own work
John Horton Sep 17, 2026 1 Share tl; dr
I created a survey where the respondent uploads a PDF of their paper. Expected Parrot uses a model to prepare an interview guide and then sends it to an AI interviewer that quizzes the author and captures the transcript for analysis. Research conferences are one obvious application, but this approach can be used (far) more generally to combat “slop cannons” of humans uncritically passing off AI work as their own.
Across academia, people are complaining about being inundated with AI-generated research and it’s breaking “classical” academia that depends on a bunch of human-scale evaluation and curation. I am, of course, a huge fan of using AI to accelerate science, but it’s not clear what’s the point of a system where you slap your name on something Claude wrote. I don’t want to read or evaluate a paper written by an “author” who doesn’t even understand what they supposedly did.
Thanks for reading! Subscribe for free to receive new posts and support my work.
Here’s a tweet yesterday about one journal fighting back by having the editor interview the “authors” of submissions:
Transactions on Machine Learning Research @TmlrOrg TMLR has faced a deluge of submissions, necessitating stricter desk rejection policies due to limited reviewer capacity
Co-EiC Nihar Shah reached out to authors of 10 papers slated for desk reject. Could they answer questions about their *own* submission?
medium.com/@TmlrOrg/askin… 8:33 PM · Sep 16, 2026 · 254K Views 21 Replies · 148 Reposts · 757 Likes The results were predictable: authors had no clue on what they did. Possibly some amount of spot-checking on this might sufficiently deter would-be slopsters, but I doubt it. This “works” but is not very scalable.
The slop cannon externality isn’t just a journals problem
What’s happening in academia seems to have parallels across the economy. I was teaching some executives visiting MIT and said “we all have encountered the colleague who just sends us pages of straight copy-and-pasted ChatGPT output” and there was a audible groan across the classroom.
Not to get to Nostradmus-y/back-patting, but I predicted all this: back in December 2022 I wrote a memo on what I thought would happen as AI hit the world. One of my predictions was that we’d have this moment of disequilibrium where AI and human-generated work were mixed and the issues this would cause:
Imagine the world where all clothing was made by hand. Clothes were expensive (price = marginal cost). All of a sudden, there are people on the market who secretly have power looms, unlimited cotton, unlimited automated stitching machines, etc. They are going to try to flood the market, getting the “human-made” price at the machine-made cost. This will take some time to re-equilibrate.
In my memo, I also predicted that what I was calling bullshit - not quite the same as AI slop, really, but that term didn’t exist yet - would spur demand for defensive goods and technologies:
One could argue that Pangram is perhaps the most useful and successful example of this. I suspect we’ll see a lot more from them. But this Pangram only answers something about the qualities of the text without necessarily answering what we often care about: do you, the author, actually understand what you’re communicating to others, even if AI was used?
Real-time AI interviews on Expected Parrot
There are a bunch of companies selling AI Interviewers to enterprises. But we, Expected Parrot, wanted something researchers could use. And we wanted an AI interview to be capable of being “mixed in” with a conventional survey, and to be highly customizable because we think there’s still a ton of research to be done on how to make these work well. Our version allows you to choose voice or text mode, pass any interview guide you want (essentially a system prompt for the AI model powering the conversation), use whatever languages you need, choose settings and controls to present to respondents, mix structured questions within or alongside the interview, etc. We want you to tinker with everything.
What I wanted to happen in the survey:
The author/user uploads their survey
Another vision-capable model reads the paper and prepares an interview guide
The author has a live AI Interview, either voice or text
This is what it actually looks like, on our site:
When a human takes the survey by following the URL, they are first asked to upload a file:
They wait while the “thinking question” runs and then, when ready, brought to an AI interview. I set it up so the user had a choice of voice or text, but this is controllable.
Trying it out / hoisted by my own petard
I tested it on myself by uploaded a recent paper of mine which you really should read. This isn’t the full transcript but rather the part that started to annoy me (this was a voice interview):
Note that I got a bit annoyed with the responses. You might wonder why I selected an example highlighting “bad” performance? Because it related to our philosophy of tooling more generally: you, the researcher, should be able to see how things went and fiddle with the controls, models, approach etc. until you get what you want. We don’t think one size fits all so we make everything adjustable.
It thinks I wrote the paper I wrote (85% confidence)
What does the Research Agent think of my responses? Well, it’s fairly confident that I’m the author:
But it also flags what it didn’t get to dive into. It’s also self-critiquing, finding a gap between what I wanted and what actually happened:
But John, I don’t want to learn Python!
First of all - how dare you? Second - that’s OK, you don’t have to. You can just describe what you want to our conversational research agent and it builds all of this because it knows our system and our APIs. This is how this survey was actually built:
It worked with me to clarify my goals then wrote all the Python in our sandboxed environment. Here’s the upload question:
And here’s the prompt for the thinking question to generate the interview guide and then the thinking question that makes use of it:
and lastly, the actual AI interview question:
The code assembles these all in a survey object and sends it to our servers. In addition to the substance of the survey, we have a way to modify the presentation and some aspects of the survey when it is human-facing:
This is what the research agent gives you:
Thanks for reading! Subscribe for free to receive new posts and support my work.
1 Share Discussion about this post Comments Restacks Top Latest Discussions No posts

## Original Extract

Using Expected Parrot to spin up bespoke AI interview + analysis workflows to flag “authors” who can’t explain their own work

Disarming the Slop Cannon (ironically, with AI)
Expected Parrot
Subscribe Sign in Disarming the Slop Cannon (ironically, with AI)
Using Expected Parrot to spin up bespoke AI interview + analysis workflows to flag “authors” who can’t explain their own work
John Horton Sep 17, 2026 1 Share tl; dr
I created a survey where the respondent uploads a PDF of their paper. Expected Parrot uses a model to prepare an interview guide and then sends it to an AI interviewer that quizzes the author and captures the transcript for analysis. Research conferences are one obvious application, but this approach can be used (far) more generally to combat “slop cannons” of humans uncritically passing off AI work as their own.
Across academia, people are complaining about being inundated with AI-generated research and it’s breaking “classical” academia that depends on a bunch of human-scale evaluation and curation. I am, of course, a huge fan of using AI to accelerate science, but it’s not clear what’s the point of a system where you slap your name on something Claude wrote. I don’t want to read or evaluate a paper written by an “author” who doesn’t even understand what they supposedly did.
Thanks for reading! Subscribe for free to receive new posts and support my work.
Here’s a tweet yesterday about one journal fighting back by having the editor interview the “authors” of submissions:
Transactions on Machine Learning Research @TmlrOrg TMLR has faced a deluge of submissions, necessitating stricter desk rejection policies due to limited reviewer capacity
Co-EiC Nihar Shah reached out to authors of 10 papers slated for desk reject. Could they answer questions about their *own* submission?
medium.com/@TmlrOrg/askin… 8:33 PM · Sep 16, 2026 · 254K Views 21 Replies · 148 Reposts · 757 Likes The results were predictable: authors had no clue on what they did. Possibly some amount of spot-checking on this might sufficiently deter would-be slopsters, but I doubt it. This “works” but is not very scalable.
The slop cannon externality isn’t just a journals problem
What’s happening in academia seems to have parallels across the economy. I was teaching some executives visiting MIT and said “we all have encountered the colleague who just sends us pages of straight copy-and-pasted ChatGPT output” and there was a audible groan across the classroom.
Not to get to Nostradmus-y/back-patting, but I predicted all this: back in December 2022 I wrote a memo on what I thought would happen as AI hit the world. One of my predictions was that we’d have this moment of disequilibrium where AI and human-generated work were mixed and the issues this would cause:
Imagine the world where all clothing was made by hand. Clothes were expensive (price = marginal cost). All of a sudden, there are people on the market who secretly have power looms, unlimited cotton, unlimited automated stitching machines, etc. They are going to try to flood the market, getting the “human-made” price at the machine-made cost. This will take some time to re-equilibrate.
In my memo, I also predicted that what I was calling bullshit - not quite the same as AI slop, really, but that term didn’t exist yet - would spur demand for defensive goods and technologies:
One could argue that Pangram is perhaps the most useful and successful example of this. I suspect we’ll see a lot more from them. But this Pangram only answers something about the qualities of the text without necessarily answering what we often care about: do you, the author, actually understand what you’re communicating to others, even if AI was used?
Real-time AI interviews on Expected Parrot
There are a bunch of companies selling AI Interviewers to enterprises. But we, Expected Parrot, wanted something researchers could use. And we wanted an AI interview to be capable of being “mixed in” with a conventional survey, and to be highly customizable because we think there’s still a ton of research to be done on how to make these work well. Our version allows you to choose voice or text mode, pass any interview guide you want (essentially a system prompt for the AI model powering the conversation), use whatever languages you need, choose settings and controls to present to respondents, mix structured questions within or alongside the interview, etc. We want you to tinker with everything.
What I wanted to happen in the survey:
The author/user uploads their survey
Another vision-capable model reads the paper and prepares an interview guide
The author has a live AI Interview, either voice or text
This is what it actually looks like, on our site:
When a human takes the survey by following the URL, they are first asked to upload a file:
They wait while the “thinking question” runs and then, when ready, brought to an AI interview. I set it up so the user had a choice of voice or text, but this is controllable.
Trying it out / hoisted by my own petard
I tested it on myself by uploaded a recent paper of mine which you really should read. This isn’t the full transcript but rather the part that started to annoy me (this was a voice interview):
Note that I got a bit annoyed with the responses. You might wonder why I selected an example highlighting “bad” performance? Because it related to our philosophy of tooling more generally: you, the researcher, should be able to see how things went and fiddle with the controls, models, approach etc. until you get what you want. We don’t think one size fits all so we make everything adjustable.
It thinks I wrote the paper I wrote (85% confidence)
What does the Research Agent think of my responses? Well, it’s fairly confident that I’m the author:
But it also flags what it didn’t get to dive into. It’s also self-critiquing, finding a gap between what I wanted and what actually happened:
But John, I don’t want to learn Python!
First of all - how dare you? Second - that’s OK, you don’t have to. You can just describe what you want to our conversational research agent and it builds all of this because it knows our system and our APIs. This is how this survey was actually built:
It worked with me to clarify my goals then wrote all the Python in our sandboxed environment. Here’s the upload question:
And here’s the prompt for the thinking question to generate the interview guide and then the thinking question that makes use of it:
and lastly, the actual AI interview question:
The code assembles these all in a survey object and sends it to our servers. In addition to the substance of the survey, we have a way to modify the presentation and some aspects of the survey when it is human-facing:
This is what the research agent gives you:
Thanks for reading! Subscribe for free to receive new posts and support my work.
1 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
