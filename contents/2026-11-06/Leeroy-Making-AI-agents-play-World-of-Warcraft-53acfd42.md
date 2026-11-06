---
source: "https://pthm.dev/leeroy-making-ai-agents-play-world-of-warcraft/"
hn_url: "https://news.ycombinator.com/item?id=50398561"
title: "Leeroy: Making AI agents play World of Warcraft"
article_title: "LEEROY: Making AI agents play World of Warcraft – pthm.dev"
image: "https://bearblog.dev/static/og-image.png"
author: "speckx"
captured_at: "2026-11-06T21:18:39Z"
capture_tool: "hn-digest"
hn_id: 50398561
score: 2
comments: 0
posted_at: "2026-11-06T21:07:44Z"
tags:
  - hacker-news
---

# Leeroy: Making AI agents play World of Warcraft

- HN: [50398561](https://news.ycombinator.com/item?id=50398561)
- Source: [pthm.dev](https://pthm.dev/leeroy-making-ai-agents-play-world-of-warcraft/)
- Score: 2
- Comments: 0
- Posted: 2026-11-06T21:07:44Z

## Translation

Title: Leeroy: Making AI agents play World of Warcraft
Article title: LEEROY: Making AI agents play World of Warcraft – pthm.dev
Description: I’ve been using AI coding agents daily for my work as a software engineer. I typically give them well-defined tasks that take me less than an hour to complet...

Article text:
LEEROY: Making AI agents play World of Warcraft
I’ve been using AI coding agents daily for my work as a software engineer. I typically give them well-defined tasks that take me less than an hour to complete manually. But I don’t think I truly understand these tools. How do they approach a task that takes hours to complete? How do they plan, recover from mistakes, collaborate with other agents? How could I even measure this? So I decided to build a benchmark.
I was inspired by RuneBench , a benchmark that has AI coding agents play RuneScape by writing code against an SDK. I played a lot of World of Warcraft during the Burning Crusade and Wrath of the Lich King expansions, and when I saw RuneBench I knew I wanted to try this in WoW. World of Warcraft has a mature private server implementation (AzerothCore). It’s a game that I know well, and it has content that is explicitly designed to be completed by a group. A dungeon designed for 5 players is a benchmark task that someone else has carefully designed and balanced for you. And you can’t complete it solo.
The outcome of this project is LEEROY .
How do the agents play the game?
World of Warcraft is a real-time game. A fight can be won or lost in less than 5 seconds. But an LLM-based agent works in turns, and one turn can take anywhere from 2 to 60 seconds. So the agents can’t directly control their character. Instead they control a headless game client, and they write Lua scripts that are executed by the game client. The scripts play the game in real time while the agent is thinking.
task . keep ( 'leader_beacon' , function ()
while true do
local p = me . position ()
local msg = string.format ( 'TANK %d: %s at %.0f,%.0f,%.1f. Follow Aldous; next stage call follows.' , me . level (), S . stage or 'regroup' , p . x , p . y , p . z )
if group . info () then chat . party ( msg )
elseif S . buddy_guid then chat . channel ( 'Deadmines' , msg ) end
wait ( S . buddy_guid and 20 or 60 ) -- 20s in group, 60s when waiting to group
end
end )
The client doesn’t have a screen either. It communicates directly with the game server using the network protocol of World of Warcraft. It converts the information sent by the server into structured data that the agent can work with. For example, that there is a kobold 25 yards behind your character (and 3 yards lower), that the kobold can currently see your character, and will attack it if it comes 3 yards closer. When the agent wants to move their character, it requests a path through a navmesh pathfinder, which takes into account the collision models of the world. (This pathfinder was the source by far of the most bugs in the project).
The SDK exposes all the information that a human player can see, and all the actions a human player can take. But nothing else. There is no “pull nearest mob” function, no “don’t let the agents cast any crowd-control abilities until I say so” function, no “position healers behind melee tanks” function. Those things need to be implemented by the agents themselves in Lua, and they are a significant part of what the benchmark measures. More details here .
How to measure the performance of the agents?
This is the hardest part of the benchmark to get right. I ended up doing 6 different iterations of the “judge” program, which scores the performance of the agents. I kept run into the problem that I would measure some metric (number of achievements, or item level, or gold, or number of elites killed), and then realize that the metric measures my choice of what to do, not an optimal strategy. The best metric I found was XP per hour. It worked well enough that all agents that successfully clear a dungeon get the same score (and we can measure in how little time they do it).
Similarly, I gave the early versions of the agents a long prompt full of advice and best practices. Of course then the benchmark measured in part which model followed my advice best. The LEEROY docs now describe the rules of the game, but never include advice on how to play it.
Running the agents for many hours is a great stress test of the game server, client and judge programs. Most of the dozens of bugs I fixed in those programs were only visible after a few hours of running. Like the bug where the judge keeps tracking the time of an agent whose game client has been disconnected for 30 minutes. Or the bug where the supervisor program that launches agents crashes, but the agents keep playing (not getting judged).
It’s really cool to watch them play
When you see a group of 5 agents, that were not told to work together, decide to form a group together (29 seconds into their first run), and then after a few hours of playing, writing new scripts, etc, pull 1 mob at a time, call for crowd-control on demand in party chat (a protocol they invented), it is exciting. And then they win the game (clear the Deadmines) (in 17 hours 41 minutes).
All game runs are recorded, and you can go back in time. When there is a wipe, you can see what each agent thought was going on in the last seconds before the wipe, and what they were making their character do.
The source code is on GitHub , and you can find the results of many hundreds of experiments on leeroy.dev .

## Original Extract

I’ve been using AI coding agents daily for my work as a software engineer. I typically give them well-defined tasks that take me less than an hour to complet...

LEEROY: Making AI agents play World of Warcraft
I’ve been using AI coding agents daily for my work as a software engineer. I typically give them well-defined tasks that take me less than an hour to complete manually. But I don’t think I truly understand these tools. How do they approach a task that takes hours to complete? How do they plan, recover from mistakes, collaborate with other agents? How could I even measure this? So I decided to build a benchmark.
I was inspired by RuneBench , a benchmark that has AI coding agents play RuneScape by writing code against an SDK. I played a lot of World of Warcraft during the Burning Crusade and Wrath of the Lich King expansions, and when I saw RuneBench I knew I wanted to try this in WoW. World of Warcraft has a mature private server implementation (AzerothCore). It’s a game that I know well, and it has content that is explicitly designed to be completed by a group. A dungeon designed for 5 players is a benchmark task that someone else has carefully designed and balanced for you. And you can’t complete it solo.
The outcome of this project is LEEROY .
How do the agents play the game?
World of Warcraft is a real-time game. A fight can be won or lost in less than 5 seconds. But an LLM-based agent works in turns, and one turn can take anywhere from 2 to 60 seconds. So the agents can’t directly control their character. Instead they control a headless game client, and they write Lua scripts that are executed by the game client. The scripts play the game in real time while the agent is thinking.
task . keep ( 'leader_beacon' , function ()
while true do
local p = me . position ()
local msg = string.format ( 'TANK %d: %s at %.0f,%.0f,%.1f. Follow Aldous; next stage call follows.' , me . level (), S . stage or 'regroup' , p . x , p . y , p . z )
if group . info () then chat . party ( msg )
elseif S . buddy_guid then chat . channel ( 'Deadmines' , msg ) end
wait ( S . buddy_guid and 20 or 60 ) -- 20s in group, 60s when waiting to group
end
end )
The client doesn’t have a screen either. It communicates directly with the game server using the network protocol of World of Warcraft. It converts the information sent by the server into structured data that the agent can work with. For example, that there is a kobold 25 yards behind your character (and 3 yards lower), that the kobold can currently see your character, and will attack it if it comes 3 yards closer. When the agent wants to move their character, it requests a path through a navmesh pathfinder, which takes into account the collision models of the world. (This pathfinder was the source by far of the most bugs in the project).
The SDK exposes all the information that a human player can see, and all the actions a human player can take. But nothing else. There is no “pull nearest mob” function, no “don’t let the agents cast any crowd-control abilities until I say so” function, no “position healers behind melee tanks” function. Those things need to be implemented by the agents themselves in Lua, and they are a significant part of what the benchmark measures. More details here .
How to measure the performance of the agents?
This is the hardest part of the benchmark to get right. I ended up doing 6 different iterations of the “judge” program, which scores the performance of the agents. I kept run into the problem that I would measure some metric (number of achievements, or item level, or gold, or number of elites killed), and then realize that the metric measures my choice of what to do, not an optimal strategy. The best metric I found was XP per hour. It worked well enough that all agents that successfully clear a dungeon get the same score (and we can measure in how little time they do it).
Similarly, I gave the early versions of the agents a long prompt full of advice and best practices. Of course then the benchmark measured in part which model followed my advice best. The LEEROY docs now describe the rules of the game, but never include advice on how to play it.
Running the agents for many hours is a great stress test of the game server, client and judge programs. Most of the dozens of bugs I fixed in those programs were only visible after a few hours of running. Like the bug where the judge keeps tracking the time of an agent whose game client has been disconnected for 30 minutes. Or the bug where the supervisor program that launches agents crashes, but the agents keep playing (not getting judged).
It’s really cool to watch them play
When you see a group of 5 agents, that were not told to work together, decide to form a group together (29 seconds into their first run), and then after a few hours of playing, writing new scripts, etc, pull 1 mob at a time, call for crowd-control on demand in party chat (a protocol they invented), it is exciting. And then they win the game (clear the Deadmines) (in 17 hours 41 minutes).
All game runs are recorded, and you can go back in time. When there is a wipe, you can see what each agent thought was going on in the last seconds before the wipe, and what they were making their character do.
The source code is on GitHub , and you can find the results of many hundreds of experiments on leeroy.dev .
