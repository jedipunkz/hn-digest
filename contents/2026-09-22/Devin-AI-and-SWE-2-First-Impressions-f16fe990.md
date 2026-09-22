---
source: "https://catalins.tech/devin-ai-swe-2-review/"
hn_url: "https://news.ycombinator.com/item?id=49800923"
title: "Devin AI and SWE-2 First Impressions"
article_title: "Devin AI & SWE-2 First Impressions"
image: "https://catalins.tech/content/images/size/w1200/2026/09/devinai-thumbnail-image.webp"
author: "cmpit"
captured_at: "2026-09-22T14:24:00Z"
capture_tool: "hn-digest"
hn_id: 49800923
score: 1
comments: 0
posted_at: "2026-09-22T13:25:16Z"
tags:
  - hacker-news
---

# Devin AI and SWE-2 First Impressions

- HN: [49800923](https://news.ycombinator.com/item?id=49800923)
- Source: [catalins.tech](https://catalins.tech/devin-ai-swe-2-review/)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T13:25:16Z

## Translation

Title: Devin AI and SWE-2 First Impressions
Article title: Devin AI & SWE-2 First Impressions
Description: First impressions of Devin AI, SWE-2 and Fusion on coding tasks from a real project. The good and the bad parts and why the CLI won me over.

Article text:
Sign in
Subscribe
Menu
Article
AI
Devin AI & SWE-2 First Impressions
Get your brand in front of our audience of developers. Book your ad directly via the button. No sales calls, no back-and-forth.
I had the opportunity of trying Devin and its suite for the first time. By suite I mean its desktop app , CLI tool , SWE-2 model and the Fusion mode.
In this post, I'll share my thoughts and observations after using it for approximately 1 week. I noted them down as I went along and I kept them raw for genuine feedback.
I'm not sure about the history behind it, but I think the old WindSurf got merged with Devin.
As a result, I see my old WindSurf chats in Devin under the "Recent sessions" section. When I click "View all" it takes me to a new page, but I don't see the old WindSurf chats anymore. So, I can't delete them.
This has happened only when I used the desktop app for the first time. Once I had more sessions in Devin, the old WindSurf chats vanished (which is good).
0:00
/ 0:17
1×
The desktop app is very busy. It feels like it's trying to be both an IDE and an agent harness. I can't pinpoint the exact parts that give me this impression, but the app feels crowded.
That made me use the CLI more, even though I like AI desktop apps.
A new session opens the chat interface and an empty panel on the right side. The right side panel clutters the view. It always opens up with a new session and you need to manually close it. The "New session" option is cool, though, because it starts a new chat with all the context from the existing space. When you start a new session, the agent can start working right away on the task without having to collect all the context again.
Why are the "Open file" and "Open customizations" options in this sidebar? I feel that they should be on the left sidebar, since they seem like general options, not necessarily tied to spaces.
I'd make the right sidebar less intrusive. Instead of opening automatically all the time, let the users open it manually if they need it.
Would be nice to be able to send existing spaces/sessions to the cloud and vice-versa. For example, I start working on something locally, but then I want to move it to the cloud. This seems possible from the CLI, but not from the desktop app. You can run the /cloud command to switch a session to run in Devin Cloud.
There's also the /handoff command that hands off a task to a remote Devin session. It starts a remote session in a separate VM with its shell, browser and full repo access. This way, it can work on a given task while your away from the keyboard.
I'm a bit confused by the difference between /cloud and /handoff . According to their docs AI assistant, the former starts a new cloud session, while the latter transfers an in-progress local session to the cloud.
Devin CLI setup took quite a bit of time. I stopped the installation session a few times thinking that it got stuck. After 3 tries or so, it successfully got installed. Maybe it was just an isolated incident with my internet and/or machine.
If you start a task in the CLI, you can see its name in the "Spaces" section, but that's all about it. You can't see the conversation and you can't perform any action. I don't know how feasible it is, but it would be nice if you could interact with a chat from both, even if not at at the same time.
I don't care much about benchmarks when trying LLMs. All I care about is the cost, how well it integrates with my stack and how well it handles my tasks.
But if you're someone that cares about benchmarks, let's look at some from Cognition's studies .
I wouldn’t read too much into these benchmarks, but they're good to get an overall idea. SWE-2 is not on par with Fable 5.1 and GPT-6 Astra, but it offers really good performance for a lot less money.
Mind you that I tested SWE-2 on real tasks while working on Documenso , not dummy, irrelevant examples and it did really well.
The model talks the closest to how a human talks. No unnecessary information, over-the-top phrasing and verbosity. You have no idea how important this is. It doesn't wear you down after a day of talking to it. In contrast, Fable 5.1's output is so verbose that your brain hurts at the end of the day.
It's a smart model. I kept running the same tasks with both SWE-2 Max and Fable 5.1 Max in Cursor and the results were very similar. I'm not saying that they're on par. Fable 5.1 is one of the top models. The idea is that I didn't have to cross-check SWE-2's code with Fable very often. It was extremely rare when I asked Fable for a second opinion. To me, that looks like a success for SWE-2.
It knows how to use the browser flawlessly. Gave it a GitHub issue and asked it to check if the issue still persists in the codebase. I just sat and watched how it used the browser because I was curious if it can reproduce the issue. It did it, confirmed the issue still persists and fixed it.
Devin also comes with an interesting feature called Fusion. Based on my understanding, it's a hybrid mode with 2 models: a lead model and an assistant. The lead usually uses a top of the line model like Fable 5.1 or GPT-6 Astra for planning, decisions, and other critical stuff.
The assistant uses an efficient model like SWE-2 for the actual work like code implementation, tests and verifications.
The main idea of Fusion is to allow you to get the most bang out of your money by using frontier models for less.
For some sessions, I used SWE-2 Max alone and for others I used Fusion on Medium with Fable 5.1 as the lead and SWE-2 High as the sidekick. Both configurations yielded good results.
I'd show you some stats of my usage, but it's impossible to find anything except the weekly quota or session-based usage, aka the usage for only one session. If I find a way of seeing my daily usage or something, I'll update the post.
Very little information about the usage. At least in the editor and account web page. You have the progress bar for the weekly quota and that's mostly it.
I'd like to see more info like the tokens spent, remaining tokens, daily usage, monthly usage, etc.
Every page I could find about usage shows "0 (zero)" or "No data for this period".
In the CLI you have the /session-stats commands that show the information I was looking for. However, it only shows the stats for the current sessions.
Unfortunately, the /usage command isn't too helpful, since it only shows the percentage of your weekly quota usage.
It absolutely drives me nuts with the Devin/WindSurf mix up. I click on "Devin Usage" and it takes me to the old WindSurf account page, for example.
Overall, I really like Devin and its suite, with the exception of the desktop app. It feels busy and less powerful than the CLI.
If you're looking for a good balance between cost and capability, the combo of Devin CLI with SWE-2 or Fusion is great.
Everything I wrote here is based on my first impressions. That means, some of the things I mentioned may be just my skill issues. If that's the case, I stand corrected and will update the post accordingly.
For the sake of transparency, I want to mention that I received the Max plan after replying to this X post. As you can see from the replies, I'm not the only that received a plan.
I'm giving away five $200 Max plans for @DevinAI
If you're using tools like Claude Code, Codex, or Cursor and haven't tried @DevinAI , comment below with what you're building to be eligible. ⚡️ https://t.co/E2i77n9mxi
However, there are no strings attached. The Cognition team didn't ask for anything in return and they don't even know that I'm posting this. I simply did it because I like to write and share stuff.
Get your brand in front of our audience of developers. Book your ad directly via the button. No sales calls, no back-and-forth.
Topic
AI Devin Developer Tools
Codex vs Claude Code Desktop Apps
I compared the CLIs on video earlier in Codex vs Warp vs…
Get your brand in front of our audience of developers. Book your ad directly via the button. No sales calls, no back-and-forth.

## Original Extract

First impressions of Devin AI, SWE-2 and Fusion on coding tasks from a real project. The good and the bad parts and why the CLI won me over.

Sign in
Subscribe
Menu
Article
AI
Devin AI & SWE-2 First Impressions
Get your brand in front of our audience of developers. Book your ad directly via the button. No sales calls, no back-and-forth.
I had the opportunity of trying Devin and its suite for the first time. By suite I mean its desktop app , CLI tool , SWE-2 model and the Fusion mode.
In this post, I'll share my thoughts and observations after using it for approximately 1 week. I noted them down as I went along and I kept them raw for genuine feedback.
I'm not sure about the history behind it, but I think the old WindSurf got merged with Devin.
As a result, I see my old WindSurf chats in Devin under the "Recent sessions" section. When I click "View all" it takes me to a new page, but I don't see the old WindSurf chats anymore. So, I can't delete them.
This has happened only when I used the desktop app for the first time. Once I had more sessions in Devin, the old WindSurf chats vanished (which is good).
0:00
/ 0:17
1×
The desktop app is very busy. It feels like it's trying to be both an IDE and an agent harness. I can't pinpoint the exact parts that give me this impression, but the app feels crowded.
That made me use the CLI more, even though I like AI desktop apps.
A new session opens the chat interface and an empty panel on the right side. The right side panel clutters the view. It always opens up with a new session and you need to manually close it. The "New session" option is cool, though, because it starts a new chat with all the context from the existing space. When you start a new session, the agent can start working right away on the task without having to collect all the context again.
Why are the "Open file" and "Open customizations" options in this sidebar? I feel that they should be on the left sidebar, since they seem like general options, not necessarily tied to spaces.
I'd make the right sidebar less intrusive. Instead of opening automatically all the time, let the users open it manually if they need it.
Would be nice to be able to send existing spaces/sessions to the cloud and vice-versa. For example, I start working on something locally, but then I want to move it to the cloud. This seems possible from the CLI, but not from the desktop app. You can run the /cloud command to switch a session to run in Devin Cloud.
There's also the /handoff command that hands off a task to a remote Devin session. It starts a remote session in a separate VM with its shell, browser and full repo access. This way, it can work on a given task while your away from the keyboard.
I'm a bit confused by the difference between /cloud and /handoff . According to their docs AI assistant, the former starts a new cloud session, while the latter transfers an in-progress local session to the cloud.
Devin CLI setup took quite a bit of time. I stopped the installation session a few times thinking that it got stuck. After 3 tries or so, it successfully got installed. Maybe it was just an isolated incident with my internet and/or machine.
If you start a task in the CLI, you can see its name in the "Spaces" section, but that's all about it. You can't see the conversation and you can't perform any action. I don't know how feasible it is, but it would be nice if you could interact with a chat from both, even if not at at the same time.
I don't care much about benchmarks when trying LLMs. All I care about is the cost, how well it integrates with my stack and how well it handles my tasks.
But if you're someone that cares about benchmarks, let's look at some from Cognition's studies .
I wouldn’t read too much into these benchmarks, but they're good to get an overall idea. SWE-2 is not on par with Fable 5.1 and GPT-6 Astra, but it offers really good performance for a lot less money.
Mind you that I tested SWE-2 on real tasks while working on Documenso , not dummy, irrelevant examples and it did really well.
The model talks the closest to how a human talks. No unnecessary information, over-the-top phrasing and verbosity. You have no idea how important this is. It doesn't wear you down after a day of talking to it. In contrast, Fable 5.1's output is so verbose that your brain hurts at the end of the day.
It's a smart model. I kept running the same tasks with both SWE-2 Max and Fable 5.1 Max in Cursor and the results were very similar. I'm not saying that they're on par. Fable 5.1 is one of the top models. The idea is that I didn't have to cross-check SWE-2's code with Fable very often. It was extremely rare when I asked Fable for a second opinion. To me, that looks like a success for SWE-2.
It knows how to use the browser flawlessly. Gave it a GitHub issue and asked it to check if the issue still persists in the codebase. I just sat and watched how it used the browser because I was curious if it can reproduce the issue. It did it, confirmed the issue still persists and fixed it.
Devin also comes with an interesting feature called Fusion. Based on my understanding, it's a hybrid mode with 2 models: a lead model and an assistant. The lead usually uses a top of the line model like Fable 5.1 or GPT-6 Astra for planning, decisions, and other critical stuff.
The assistant uses an efficient model like SWE-2 for the actual work like code implementation, tests and verifications.
The main idea of Fusion is to allow you to get the most bang out of your money by using frontier models for less.
For some sessions, I used SWE-2 Max alone and for others I used Fusion on Medium with Fable 5.1 as the lead and SWE-2 High as the sidekick. Both configurations yielded good results.
I'd show you some stats of my usage, but it's impossible to find anything except the weekly quota or session-based usage, aka the usage for only one session. If I find a way of seeing my daily usage or something, I'll update the post.
Very little information about the usage. At least in the editor and account web page. You have the progress bar for the weekly quota and that's mostly it.
I'd like to see more info like the tokens spent, remaining tokens, daily usage, monthly usage, etc.
Every page I could find about usage shows "0 (zero)" or "No data for this period".
In the CLI you have the /session-stats commands that show the information I was looking for. However, it only shows the stats for the current sessions.
Unfortunately, the /usage command isn't too helpful, since it only shows the percentage of your weekly quota usage.
It absolutely drives me nuts with the Devin/WindSurf mix up. I click on "Devin Usage" and it takes me to the old WindSurf account page, for example.
Overall, I really like Devin and its suite, with the exception of the desktop app. It feels busy and less powerful than the CLI.
If you're looking for a good balance between cost and capability, the combo of Devin CLI with SWE-2 or Fusion is great.
Everything I wrote here is based on my first impressions. That means, some of the things I mentioned may be just my skill issues. If that's the case, I stand corrected and will update the post accordingly.
For the sake of transparency, I want to mention that I received the Max plan after replying to this X post. As you can see from the replies, I'm not the only that received a plan.
I'm giving away five $200 Max plans for @DevinAI
If you're using tools like Claude Code, Codex, or Cursor and haven't tried @DevinAI , comment below with what you're building to be eligible. ⚡️ https://t.co/E2i77n9mxi
However, there are no strings attached. The Cognition team didn't ask for anything in return and they don't even know that I'm posting this. I simply did it because I like to write and share stuff.
Get your brand in front of our audience of developers. Book your ad directly via the button. No sales calls, no back-and-forth.
Topic
AI Devin Developer Tools
Codex vs Claude Code Desktop Apps
I compared the CLIs on video earlier in Codex vs Warp vs…
Get your brand in front of our audience of developers. Book your ad directly via the button. No sales calls, no back-and-forth.
