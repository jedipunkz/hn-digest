---
source: "https://www.compile.la/ai-changed-the-team-not-the-process/"
hn_url: "https://news.ycombinator.com/item?id=50001749"
title: "AI Changed the Team, Not the Process"
article_title: "AI Changed the Team, Not the Process"
image: "https://storage.ghost.io/c/a4/b1/a4b18877-561a-494f-b192-85940e5b9e31/content/images/size/w1200/2026/09/throughline_meal_slot_trace_202609.png"
author: "dalfonso"
captured_at: "2026-10-08T04:48:54Z"
capture_tool: "hn-digest"
hn_id: 50001749
score: 1
comments: 0
posted_at: "2026-10-08T04:10:21Z"
tags:
  - hacker-news
---

# AI Changed the Team, Not the Process

- HN: [50001749](https://news.ycombinator.com/item?id=50001749)
- Source: [www.compile.la](https://www.compile.la/ai-changed-the-team-not-the-process/)
- Score: 1
- Comments: 0
- Posted: 2026-10-08T04:10:21Z

## Translation

Title: AI Changed the Team, Not the Process
Description: When Guy Guyadeen left LA in 2009, the tech industry was a tiny niche. He moved to the East Coast and eventually landed in Mountain View as a Product Manager at Google. He decided to move back to LA and thought, “This is going to be career suicide.” But after

Article text:
Sign in
Subscribe
AI Changed the Team, Not the Process
When Guy Guyadeen left LA in 2009, the tech industry was a tiny niche. He moved to the East Coast and eventually landed in Mountain View as a Product Manager at Google. He decided to move back to LA and thought, “This is going to be career suicide.” But after arriving in 2018, Guy was blown away, “What’s going on here? There is so much happening!”
Guy joined Act One Ventures as Entrepreneur in Residence, followed by product roles at Quibi and Tapcart. He eventually rejoined Google. This time, working out of the LA office. Guy spent these years on teams building high-quality complex software at scale.
Now, Guy is building litFit, an AI nutrition coach that keeps track of your eating habits and learns them over time. litFit’s nutrition data comes from food databases like Open Food Facts and the USDA’s Food Data Central rather than the model’s guess. While training for a triathlon, Guy saw a gap in the market, “There are apps that try to do part of this, but none of them feel like I have a $120-an-hour coach on speed dial. If you try to do this with bare ChatGPT, it can’t do the math reliably. It forgets one day to the next. It cannot maintain a ledger reliably.”
Guy uses coding agents, but not the way most people do. This is because “what LLMs are trained on is not what we do internally at companies doing software engineering.” The question for him became: “How do you build something in a way that would meet the same quality bar at the places I’ve worked in the past, like Google?”
His thesis? The way to reach the quality bar is by using the same software engineering processes that have worked for years: “I, as a product manager, write a product requirement document [PRD] with engineers who write technical architecture and technical design documents [TDD], and then other engineers implement them.” On his current team-of-one, the “other engineers” are Claude and Codex.
“I have a corpus of five or six PRDs and nine or ten TDDs that I've put my blood and sweat and love into,” says Guy. He then built a tool called “Throughline” to keep litFit’s code connected to the product decisions behind it. The AI agents don't have to understand the entire project in their limited context windows. Throughline “enables you to run a shell command that says, ‘Given this line of code, how does it connect to all the documentation?’ So it goes up to the technical docs. It also goes up to the PRDs.”
Guy adds, “Agents love to declare victory when the tests go green. Throughline won’t let them. Done means the promise to the user is kept.”
Guy gave me a concrete example, which I’m paraphrasing. A user says, “I had two eggs for breakfast.” The product promise is that litFit records the meal as stated by the user, and the technical design reflects that promise. The AI shouldn’t hallucinate a different meal.
But passing tests still doesn’t mean the feature is done. If the user then quickly says, “Actually, make that lunch,” litFit still has to honor the original promise and record the meal the way the user ultimately stated it. Until that behavior is implemented, Throughline considers the feature incomplete.
But he understands this might not be the only way to build: “I'm rebuilding the whole playbook – specs, reviews, traceability, how you steer a team – for teams that wield coding agents. I'm curious who else is converging on the same ideas, or has better ones.” Guy believes that we’re at a new frontier with AI, “Compilers didn't replace programmers in the 70s, they raised the level of abstraction you could work at. That's how I see coding agents: they raise the level you can create at, but we haven't figured out the processes yet. That’s what’s really exciting and fun about this.”
Building with coding agents too? Guy is interested in comparing notes with other teams doing the same. Reach him on LinkedIn .
Podcast Episode 1: Snap, Specs, and Missed Opportunities
Snap is the home team, and I want it to win. But being a fan of the company hasn’t made me a regular Snapchat user, or sold me on its new Specs glasses.
In the first episode of Compile LA, I talk about why I’m rooting for Snap,
From Corporate VP to Founder, Rooted in LA
Ari Wilson joined VideoAmp as Vice President of Engineering in April 2026. After four months, he was involved in the planning of a major layoff driven by AI efficiencies, a common theme in 2026. When the time came to execute the layoff, he was informed that his role was being
I moved to Los Angeles County (Hermosa Beach) on November 11, 2015. I’ve founded companies here, attended TechWeek, attended meetups, worked in Santa Monica, worked from home, worked from coworking spaces. I’ve lived in two apartments in Hermosa and finally bought a townhome in North Redondo Beach. I
Los Angeles tech news and stories about the founders, engineers, and companies making hardware and software in LA.

## Original Extract

When Guy Guyadeen left LA in 2009, the tech industry was a tiny niche. He moved to the East Coast and eventually landed in Mountain View as a Product Manager at Google. He decided to move back to LA and thought, “This is going to be career suicide.” But after

Sign in
Subscribe
AI Changed the Team, Not the Process
When Guy Guyadeen left LA in 2009, the tech industry was a tiny niche. He moved to the East Coast and eventually landed in Mountain View as a Product Manager at Google. He decided to move back to LA and thought, “This is going to be career suicide.” But after arriving in 2018, Guy was blown away, “What’s going on here? There is so much happening!”
Guy joined Act One Ventures as Entrepreneur in Residence, followed by product roles at Quibi and Tapcart. He eventually rejoined Google. This time, working out of the LA office. Guy spent these years on teams building high-quality complex software at scale.
Now, Guy is building litFit, an AI nutrition coach that keeps track of your eating habits and learns them over time. litFit’s nutrition data comes from food databases like Open Food Facts and the USDA’s Food Data Central rather than the model’s guess. While training for a triathlon, Guy saw a gap in the market, “There are apps that try to do part of this, but none of them feel like I have a $120-an-hour coach on speed dial. If you try to do this with bare ChatGPT, it can’t do the math reliably. It forgets one day to the next. It cannot maintain a ledger reliably.”
Guy uses coding agents, but not the way most people do. This is because “what LLMs are trained on is not what we do internally at companies doing software engineering.” The question for him became: “How do you build something in a way that would meet the same quality bar at the places I’ve worked in the past, like Google?”
His thesis? The way to reach the quality bar is by using the same software engineering processes that have worked for years: “I, as a product manager, write a product requirement document [PRD] with engineers who write technical architecture and technical design documents [TDD], and then other engineers implement them.” On his current team-of-one, the “other engineers” are Claude and Codex.
“I have a corpus of five or six PRDs and nine or ten TDDs that I've put my blood and sweat and love into,” says Guy. He then built a tool called “Throughline” to keep litFit’s code connected to the product decisions behind it. The AI agents don't have to understand the entire project in their limited context windows. Throughline “enables you to run a shell command that says, ‘Given this line of code, how does it connect to all the documentation?’ So it goes up to the technical docs. It also goes up to the PRDs.”
Guy adds, “Agents love to declare victory when the tests go green. Throughline won’t let them. Done means the promise to the user is kept.”
Guy gave me a concrete example, which I’m paraphrasing. A user says, “I had two eggs for breakfast.” The product promise is that litFit records the meal as stated by the user, and the technical design reflects that promise. The AI shouldn’t hallucinate a different meal.
But passing tests still doesn’t mean the feature is done. If the user then quickly says, “Actually, make that lunch,” litFit still has to honor the original promise and record the meal the way the user ultimately stated it. Until that behavior is implemented, Throughline considers the feature incomplete.
But he understands this might not be the only way to build: “I'm rebuilding the whole playbook – specs, reviews, traceability, how you steer a team – for teams that wield coding agents. I'm curious who else is converging on the same ideas, or has better ones.” Guy believes that we’re at a new frontier with AI, “Compilers didn't replace programmers in the 70s, they raised the level of abstraction you could work at. That's how I see coding agents: they raise the level you can create at, but we haven't figured out the processes yet. That’s what’s really exciting and fun about this.”
Building with coding agents too? Guy is interested in comparing notes with other teams doing the same. Reach him on LinkedIn .
Podcast Episode 1: Snap, Specs, and Missed Opportunities
Snap is the home team, and I want it to win. But being a fan of the company hasn’t made me a regular Snapchat user, or sold me on its new Specs glasses.
In the first episode of Compile LA, I talk about why I’m rooting for Snap,
From Corporate VP to Founder, Rooted in LA
Ari Wilson joined VideoAmp as Vice President of Engineering in April 2026. After four months, he was involved in the planning of a major layoff driven by AI efficiencies, a common theme in 2026. When the time came to execute the layoff, he was informed that his role was being
I moved to Los Angeles County (Hermosa Beach) on November 11, 2015. I’ve founded companies here, attended TechWeek, attended meetups, worked in Santa Monica, worked from home, worked from coworking spaces. I’ve lived in two apartments in Hermosa and finally bought a townhome in North Redondo Beach. I
Los Angeles tech news and stories about the founders, engineers, and companies making hardware and software in LA.
