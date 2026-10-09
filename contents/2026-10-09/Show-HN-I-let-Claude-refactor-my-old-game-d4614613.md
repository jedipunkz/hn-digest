---
source: "https://github.com/victorqribeiro/aimAndShoot"
hn_url: "https://news.ycombinator.com/item?id=50018821"
title: "Show HN: I let Claude refactor my old game"
article_title: "GitHub - victorqribeiro/aimAndShoot: A neuroevolution game experiment. · GitHub"
image: "https://repository-images.githubusercontent.com/216132604/1c686080-f370-11e9-82b1-b8ae3c8b8ba1"
author: "atum47"
captured_at: "2026-10-09T11:40:59Z"
capture_tool: "hn-digest"
hn_id: 50018821
score: 1
comments: 0
posted_at: "2026-10-09T11:04:00Z"
tags:
  - hacker-news
---

# Show HN: I let Claude refactor my old game

- HN: [50018821](https://news.ycombinator.com/item?id=50018821)
- Source: [github.com](https://github.com/victorqribeiro/aimAndShoot)
- Score: 1
- Comments: 0
- Posted: 2026-10-09T11:04:00Z

## Translation

Title: Show HN: I let Claude refactor my old game
Article title: GitHub - victorqribeiro/aimAndShoot: A neuroevolution game experiment. · GitHub
Description: A neuroevolution game experiment. Contribute to victorqribeiro/aimAndShoot development by creating an account on GitHub.
HN text: I've been using Claude code to finish up some projects that were on hold for a while. I've been doing this in my laptop, keeping the changes local. In one session I was offered $100 usd credits to let Claude connect and work in one of my GitHub projects directly. So I chose aimAndShoot, a neural evolution game from 2019. It catched some bugs and introduced improvements. I told it I was going to post about this experience on hackernews and lo and behold, the topic was already in it's training data. It knew about the project and the fact that I hadn't addressed the feedback it got. Anyhow, I asked it to address the comments on the topic and here's the result. Live game is here: https://victorribeiro.com/aimAndShoot/

Article text:
GitHub - victorqribeiro/aimAndShoot: A neuroevolution game experiment. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
victorqribeiro
Notifications You must be signed in to change notification settings
Star 234 ( 234 ) You must be signed in to star a repository
A neuroevolution game experiment.
victorribeiro.com/aimAndShoot Topics
Readme MIT license Activity Stars
24 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
52 Commits 52 Commits Folders and files
.github .github api api css css docs docs js js sounds sounds tests tests LICENSE LICENSE README.md README.md artwork.png artwork.png controls.png controls.png favicon.png favicon.png index.html index.html manifest.json manifest.json sw.js sw.js View all files Repository files navigation
You're Nole Ksum (the k is silent), a citizen concern about the uprising of the machine who decided to take matters into your own hands and put an end to all artificial intelligence. You must kill all the evil robots controlled by Neural Networks and stop them from evolving into more dangerous beings. The entire human race counts on you, don't let them down.
w, a, s, d - Move the player up, left, down, right. Arrow keys do the same.
mouse - Aims and shoots (click).
I do not recommend using this on a mobile, but if you must
The big left circle moves the player.
The big right circle aims the player.
The two little circles above them, shoot.
Kill the bots, don't get killed. Also don't touch the borders of the screen, they hurt you. But, feel free to push the bots into them.
Above the bots there are two status bars.
The red one indicates health, if it's empty you die.
The green one is the cool down meter, if it's empty you can't shoot until it regenerates.
I've always wanted to take the time to make a Neuroevolution experiment, so I did.
Each bot is controlled by it's own Neural Network (that I made a while back - here ). When all the bots die, the genetic algorithm evaluates their fitness score (based on how many shots they fired, how many hits the got, how many friends they shot, how much they hurt themselves and how much they moved during the round) and cross the ones with the highest scores.
This goes on forever, until the player dies (which will happen eventually, so Nole can't never save the human race, after all). By the way, the background history is a joke. I don't mean to make fun of anyone. The idea just seems funny and fit the project.
Fun Fact: the artwork was created using my PaintDraw tool.
2026 Update: fixed and extended with Claude Code
The live version is no longer the 2019 original. In 2026 the game was fixed and extended with Claude Code :
The bots can actually learn now. Several bugs in the neuroevolution were fixed: each bot sees every player plus its own state, the aim works, the fitness no longer rewards spraying bullets, and the best bot of each generation is kept.
Shared evolution. The bots are no longer reset when you die. They come from one population kept on the server and evolved by everyone who played before you. Your browser only reports how each bot did in the round; the server scores the bots and breeds the next generations. The HUD shows the shared generation and your round. If the server can't be reached, the game falls back to evolving the bots locally, as before.
Suggestions from the 2019 Hacker News thread . Your health refills every round, the game-over screen shows the round you reached and your best, the gunshots are quieter (press M to mute), bots no longer spawn on top of you, and killing the shooters first no longer breeds pacifists as easily.
Fairer gameplay. Limited fire rate, a short grace period at the start of each round, the same movement speed at any screen refresh rate, a fixed-size arena that scales to fit the screen, and collision fixes.
The backend is PHP with SQLite ( api/ ). See docs/shared-evolution-plan.md for how it works. The GitHub Pages link now redirects to the live version.
A neuroevolution game experiment.
victorribeiro.com/aimAndShoot Topics
Readme MIT license Activity Stars
24 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A neuroevolution game experiment. Contribute to victorqribeiro/aimAndShoot development by creating an account on GitHub.

I've been using Claude code to finish up some projects that were on hold for a while. I've been doing this in my laptop, keeping the changes local. In one session I was offered $100 usd credits to let Claude connect and work in one of my GitHub projects directly. So I chose aimAndShoot, a neural evolution game from 2019. It catched some bugs and introduced improvements. I told it I was going to post about this experience on hackernews and lo and behold, the topic was already in it's training data. It knew about the project and the fact that I hadn't addressed the feedback it got. Anyhow, I asked it to address the comments on the topic and here's the result. Live game is here: https://victorribeiro.com/aimAndShoot/

GitHub - victorqribeiro/aimAndShoot: A neuroevolution game experiment. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
victorqribeiro
Notifications You must be signed in to change notification settings
Star 234 ( 234 ) You must be signed in to star a repository
A neuroevolution game experiment.
victorribeiro.com/aimAndShoot Topics
Readme MIT license Activity Stars
24 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
52 Commits 52 Commits Folders and files
.github .github api api css css docs docs js js sounds sounds tests tests LICENSE LICENSE README.md README.md artwork.png artwork.png controls.png controls.png favicon.png favicon.png index.html index.html manifest.json manifest.json sw.js sw.js View all files Repository files navigation
You're Nole Ksum (the k is silent), a citizen concern about the uprising of the machine who decided to take matters into your own hands and put an end to all artificial intelligence. You must kill all the evil robots controlled by Neural Networks and stop them from evolving into more dangerous beings. The entire human race counts on you, don't let them down.
w, a, s, d - Move the player up, left, down, right. Arrow keys do the same.
mouse - Aims and shoots (click).
I do not recommend using this on a mobile, but if you must
The big left circle moves the player.
The big right circle aims the player.
The two little circles above them, shoot.
Kill the bots, don't get killed. Also don't touch the borders of the screen, they hurt you. But, feel free to push the bots into them.
Above the bots there are two status bars.
The red one indicates health, if it's empty you die.
The green one is the cool down meter, if it's empty you can't shoot until it regenerates.
I've always wanted to take the time to make a Neuroevolution experiment, so I did.
Each bot is controlled by it's own Neural Network (that I made a while back - here ). When all the bots die, the genetic algorithm evaluates their fitness score (based on how many shots they fired, how many hits the got, how many friends they shot, how much they hurt themselves and how much they moved during the round) and cross the ones with the highest scores.
This goes on forever, until the player dies (which will happen eventually, so Nole can't never save the human race, after all). By the way, the background history is a joke. I don't mean to make fun of anyone. The idea just seems funny and fit the project.
Fun Fact: the artwork was created using my PaintDraw tool.
2026 Update: fixed and extended with Claude Code
The live version is no longer the 2019 original. In 2026 the game was fixed and extended with Claude Code :
The bots can actually learn now. Several bugs in the neuroevolution were fixed: each bot sees every player plus its own state, the aim works, the fitness no longer rewards spraying bullets, and the best bot of each generation is kept.
Shared evolution. The bots are no longer reset when you die. They come from one population kept on the server and evolved by everyone who played before you. Your browser only reports how each bot did in the round; the server scores the bots and breeds the next generations. The HUD shows the shared generation and your round. If the server can't be reached, the game falls back to evolving the bots locally, as before.
Suggestions from the 2019 Hacker News thread . Your health refills every round, the game-over screen shows the round you reached and your best, the gunshots are quieter (press M to mute), bots no longer spawn on top of you, and killing the shooters first no longer breeds pacifists as easily.
Fairer gameplay. Limited fire rate, a short grace period at the start of each round, the same movement speed at any screen refresh rate, a fixed-size arena that scales to fit the screen, and collision fixes.
The backend is PHP with SQLite ( api/ ). See docs/shared-evolution-plan.md for how it works. The GitHub Pages link now redirects to the live version.
A neuroevolution game experiment.
victorribeiro.com/aimAndShoot Topics
Readme MIT license Activity Stars
24 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
