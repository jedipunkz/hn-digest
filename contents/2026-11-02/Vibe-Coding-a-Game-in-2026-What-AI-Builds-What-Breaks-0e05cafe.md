---
source: "https://vlstudio.app/blog/vibe-coding-a-game-2026/"
hn_url: "https://news.ycombinator.com/item?id=50340109"
title: "Vibe-Coding a Game in 2026: What AI Builds, What Breaks"
article_title: "Vibe-Coding a Game in 2026: What AI Builds, What Breaks"
image: "https://vlstudio.app/img/lp/band_tg_game.og.jpg"
author: "vladosskins"
captured_at: "2026-11-02T17:53:27Z"
capture_tool: "hn-digest"
hn_id: 50340109
score: 5
comments: 1
posted_at: "2026-11-02T17:17:41Z"
tags:
  - hacker-news
---

# Vibe-Coding a Game in 2026: What AI Builds, What Breaks

- HN: [50340109](https://news.ycombinator.com/item?id=50340109)
- Source: [vlstudio.app](https://vlstudio.app/blog/vibe-coding-a-game-2026/)
- Score: 5
- Comments: 1
- Posted: 2026-11-02T17:17:41Z

## Translation

Title: Vibe-Coding a Game in 2026: What AI Builds, What Breaks
Description: AI can now build a playable game from a prompt. What still breaks, the Steam, Google Play and App Store rules, and when a brand game needs a studio.

Article text:
Skip to article
VL Studio
Services Work Pricing Guides Blog FAQ
Start your project
Services Work Pricing Guides Blog FAQ Start your project
VL Studio / Blog / Games
Vibe-coding a game in 2026 What AI builds well, what still breaks, and when to hire a studio
AI can now turn a short prompt into a playable browser game. Getting that game in front of real players, with prizes, payments or a store listing, is still engineering work, and the platform rules are changing.
By VL Studio Updated November 2, 2026 10 min read
In late October 2026 a well-known developer wrote that building games that play well is "shockingly easy now" , and Roblox told shareholders that people had published over 70,000 games with its prompt-to-game tool.
AI is strong at the first playable version: small browser games in three.js or Phaser, placeholder art and even code-generated music.
It is still weak at the parts players and brands pay for: game logic that holds up over a long session, consistent art and animation, multiplayer, cheat-proof leaderboards, payments and performance on ordinary phones.
Platform rules apply to AI games: Steam asks for an AI disclosure , Google Play requires in-app reporting for AI-generated content, and Apple rejects copycats and spam.
A prototype you made yourself is a strong brief for a studio. Once a game carries prizes, personal data or payments, it needs engineering and testing.
The short answer: yes, in 2026 you can vibe-code a small, playable game in an evening by describing it to an AI model, and a prototype like that is now a perfectly sensible first step. What you can't skip is the work between "playable on my phone" and "live in front of customers": fixing logic bugs that only appear in long sessions, making the art consistent, testing on mid-range Android phones, checking scores and payments on a server, and meeting Steam, Google Play, App Store or Telegram rules . This guide covers what changed this autumn, what AI handles well, where it breaks, and how to decide between doing it yourself and hiring a studio.
Methodology: this is an editorial guide. Every quote, figure and platform rule below was checked against the linked source on November 2, 2026. Where a claim has only been reported by the press and not confirmed by the company, we say so.
Three things happened within two weeks of each other.
Developers said games got easy. On October 28, Simon Willison, co-creator of the Django web framework and a widely read writer on AI tools, called October 2026 "a new inflection point" for vibe coding computer games. "Building games that play well is shockingly easy now," he wrote. None of his four example games took more than four hours of his own work, "plus many more hours of Claude Code churning away by itself" ( Simon Willison ).
Platforms reported a flood. PC Gamer reported that Roblox's shareholder letter says more than 500,000 users have tried Build, its prompt-to-game tool, and published over 70,000 games, and that more than 80% of the top 100 Roblox games used AI tools in the past month ( PC Gamer ). On itch.io, new game pages per day rose from around 3,000 to regularly over 6,000 between August 3 and October 30 ( PC Gamer ).
Players pushed back. During October's Steam Next Fest, players on the Steam subreddit asked for a filter to hide games with AI-generated content ( IGN ). Tom's Hardware reported that Claude had started refusing to recreate specific existing games, citing a rule that reads "Don't build a recreation of a specific existing game, even if it is renamed or called original." The publication noted there was no official confirmation of that text and no statement from Anthropic ( Tom's Hardware ).
For a business, the takeaway is not "games are free now". It is that a playable prototype is cheap, while attention, trust and polish are harder to come by than ever.
The first playable version. Willison's "Tree Appreciator", a spot-the-difference game, "required no iteration or follow-up prompting at all". His suggested starting prompt is simply "Build a mobile-friendly three.js game where..." ( Simon Willison ).
Browser games on known libraries. His games are HTML and JavaScript on static hosting, and he notes that frontier models "already know libraries like three.js and Phaser".
Systems that used to need a specialist. In his strategy game Battle Drones, the soundtrack "is entirely generated through code" with no audio files, and the AI built replays and AI-versus-AI battles ( Simon Willison ).
Non-programmers can get surprisingly far. A writer who set out to test whether someone who can't model, animate or code could make a 3D game built a playable browser shooter with cutscenes in about 10 days. The only paid tool was the AI coding assistant included in a $20 ChatGPT subscription, and the writer put the cost, "rounding very generously", at about $15 ( Build to Understand ).
Where vibe-coded games still break
The same sources that show how fast AI is also show where the time goes. These are the gaps we check first when someone sends us an AI-built game.
Sources for the table rows: Willison (October 28) , Willison (October 20) and Build to Understand . The last two rows are standard engineering practice rather than findings from those posts.
Willison makes the key point himself: personal games are "the perfect application for vibe coding because it's so low stakes". A brand game is not low stakes. A broken quest in a hobby project costs nothing, while a leaderboard someone cheats in a contest with real prizes costs money and goodwill.
Platform rules for AI-made games
Since January 2024, Steam's Content Survey has included an AI disclosure. It separates Pre-Generated content (art, code, sound and so on created with AI during development) from Live-Generated content (created while the game runs). For live generation you must also describe the guardrails that stop it producing illegal content, and Valve shows much of the disclosure on the store page ( Valve, AI Content on Steam ). The current Steamworks documentation adds that AI tools built into development environments, used for efficiency, are not the focus; the section is about AI content "that ships with your game, and is consumed by players" ( Steamworks: Content Survey ).
Google Play: in-app reporting for AI content
If your game generates content with AI while it runs (an AI character that chats with players, generated images or levels), Google Play requires "in-app user reporting or flagging features" so players can report offensive content without leaving the app ( Google Play Console Help ).
App Store: no copycats, no spam, in-app purchase for digital items
Apple's App Review Guidelines tell developers not to "simply copy the latest popular app on the App Store, or make some minor changes to another app's name or UI" (4.1 Copycats), not to flood the store with near-identical apps (4.3 Spam), and to offer "lasting entertainment value" (4.2). Game currencies, levels and premium content must be sold through in-app purchase (3.1.1) ( Apple App Review Guidelines ). A fast AI clone of a hit game runs into these rules even before any legal claim does.
Telegram: Stars for digital goods
Games inside Telegram must sell digital items in Telegram Stars ( Telegram Stars payments ). We cover the payment flow and launch options in our Telegram mini games guide .
Who owns an AI-generated game?
For a brand, this matters as much as the code. In January 2025 the US Copyright Office concluded that generative AI outputs "can be protected by copyright only where a human author has determined sufficient expressive elements", and that this does not include "the mere provision of prompts". Using AI to assist, or including AI material in a larger human-made work, does not by itself block protection ( U.S. Copyright Office ).
In practice: the more of a game's characters, writing and art comes from human creative decisions, and the better you document them, the stronger your claim to the parts that make the game yours. This is general information, not legal advice; for a campaign with a mascot or licensed brand assets, ask your lawyer.
Clones are the other side of ownership. Speaking to developers about AI-made copies, Tom Giardino of Valve's Steam business team said Valve does not want to dictate which tools studios use, and advised building a community early and continuing to improve the original, since individual mechanics are relatively easy to copy, while new features and a distinct visual style are harder to copy. He also acknowledged that building a community takes time and is not enough protection on its own ( ComputerBase, in German ).
Do it yourself, or hire a studio?
A prototype is never wasted. A version you have already played is a better specification than any document, and it makes a studio quote faster and more accurate.
How we use AI in our own games
We make two games of our own, Kosti, a 3D action game, and Block Nook, a block puzzle, both built in Godot 4 for Android and both still in development (see our Kosti and Block Nook project page ). AI does a lot of production work in them: 3D models and rigs from text and images, concept and interface art, and voice lines in three languages. What it does not do is decide. We set the style, fix what the tools get wrong and check every asset in the engine before it ships, because the gaps in the table above are the same ones we run into.
What a studio-built game costs
Our published starting prices for games are:
Branded Mini-Game from $8,000: advergames and playable campaigns.
Telegram Mini App Game from $10,000: games inside Telegram with leaderboards.
Mobile Game from $25,000: Android and iOS games, from concept to store.
"From" means the smallest sensible version; we send a fixed written quote before work starts, and new customers get 20% off the first project. If you already have a vibe-coded prototype, send it with your brief: we'll tell you what can be kept and what needs rebuilding. The full list is on our pricing page .
Can AI make a whole game from a prompt?
For small browser games, often yes. In October 2026 Simon Willison described getting a playable spot-the-difference game from a single prompt with no follow-up. Bigger games still need iteration: his fantasy RPG needed extra playtesting and prompting to fix quest bugs, and his strategy game took 12 major iterations and 321 commits by coding agents.
Do I have to disclose AI use when I publish a game on Steam?
Yes, for AI-made content that ships with the game. Steam's Content Survey asks you to describe AI content created during development (Pre-Generated) and AI content created while the game runs (Live-Generated), and Valve shows much of that disclosure on the store page. Valve says AI coding tools used for efficiency are not the focus of that section.
Do I own the copyright in an AI-generated game?
In the US, only partly. The Copyright Office concluded in January 2025 that AI outputs are protected only where a human author determined sufficient expressive elements, and that prompts alone are not enough. Human-made art, writing and creative changes to AI output can be protected. For a brand game, get legal advice and keep a record of the human work.
When is a vibe-coded game good enough for a brand campaign?
When nothing valuable rides on it: an internal demo, a pitch prototype or a short single-player toy with no prizes, no personal data and no payments. Once there are real prizes, a leaderboard people will try to cheat, sign-ups, purchases or a store release, the parts AI tools skip (server checks, testing on real phones, store rules and support) need engineering.
Building inside Telegram? Read our Telegram mini games guide . Built an app (not a game) with AI? See how to make an AI-built app production-ready .
Have a game idea or an AI prototype? Get a fixed quote.
Send us the prototype or a short brief. We reply within one business day with a fixed written quote: Branded Mini-Game from $8,000,

[truncated]

## Original Extract

AI can now build a playable game from a prompt. What still breaks, the Steam, Google Play and App Store rules, and when a brand game needs a studio.

Skip to article
VL Studio
Services Work Pricing Guides Blog FAQ
Start your project
Services Work Pricing Guides Blog FAQ Start your project
VL Studio / Blog / Games
Vibe-coding a game in 2026 What AI builds well, what still breaks, and when to hire a studio
AI can now turn a short prompt into a playable browser game. Getting that game in front of real players, with prizes, payments or a store listing, is still engineering work, and the platform rules are changing.
By VL Studio Updated November 2, 2026 10 min read
In late October 2026 a well-known developer wrote that building games that play well is "shockingly easy now" , and Roblox told shareholders that people had published over 70,000 games with its prompt-to-game tool.
AI is strong at the first playable version: small browser games in three.js or Phaser, placeholder art and even code-generated music.
It is still weak at the parts players and brands pay for: game logic that holds up over a long session, consistent art and animation, multiplayer, cheat-proof leaderboards, payments and performance on ordinary phones.
Platform rules apply to AI games: Steam asks for an AI disclosure , Google Play requires in-app reporting for AI-generated content, and Apple rejects copycats and spam.
A prototype you made yourself is a strong brief for a studio. Once a game carries prizes, personal data or payments, it needs engineering and testing.
The short answer: yes, in 2026 you can vibe-code a small, playable game in an evening by describing it to an AI model, and a prototype like that is now a perfectly sensible first step. What you can't skip is the work between "playable on my phone" and "live in front of customers": fixing logic bugs that only appear in long sessions, making the art consistent, testing on mid-range Android phones, checking scores and payments on a server, and meeting Steam, Google Play, App Store or Telegram rules . This guide covers what changed this autumn, what AI handles well, where it breaks, and how to decide between doing it yourself and hiring a studio.
Methodology: this is an editorial guide. Every quote, figure and platform rule below was checked against the linked source on November 2, 2026. Where a claim has only been reported by the press and not confirmed by the company, we say so.
Three things happened within two weeks of each other.
Developers said games got easy. On October 28, Simon Willison, co-creator of the Django web framework and a widely read writer on AI tools, called October 2026 "a new inflection point" for vibe coding computer games. "Building games that play well is shockingly easy now," he wrote. None of his four example games took more than four hours of his own work, "plus many more hours of Claude Code churning away by itself" ( Simon Willison ).
Platforms reported a flood. PC Gamer reported that Roblox's shareholder letter says more than 500,000 users have tried Build, its prompt-to-game tool, and published over 70,000 games, and that more than 80% of the top 100 Roblox games used AI tools in the past month ( PC Gamer ). On itch.io, new game pages per day rose from around 3,000 to regularly over 6,000 between August 3 and October 30 ( PC Gamer ).
Players pushed back. During October's Steam Next Fest, players on the Steam subreddit asked for a filter to hide games with AI-generated content ( IGN ). Tom's Hardware reported that Claude had started refusing to recreate specific existing games, citing a rule that reads "Don't build a recreation of a specific existing game, even if it is renamed or called original." The publication noted there was no official confirmation of that text and no statement from Anthropic ( Tom's Hardware ).
For a business, the takeaway is not "games are free now". It is that a playable prototype is cheap, while attention, trust and polish are harder to come by than ever.
The first playable version. Willison's "Tree Appreciator", a spot-the-difference game, "required no iteration or follow-up prompting at all". His suggested starting prompt is simply "Build a mobile-friendly three.js game where..." ( Simon Willison ).
Browser games on known libraries. His games are HTML and JavaScript on static hosting, and he notes that frontier models "already know libraries like three.js and Phaser".
Systems that used to need a specialist. In his strategy game Battle Drones, the soundtrack "is entirely generated through code" with no audio files, and the AI built replays and AI-versus-AI battles ( Simon Willison ).
Non-programmers can get surprisingly far. A writer who set out to test whether someone who can't model, animate or code could make a 3D game built a playable browser shooter with cutscenes in about 10 days. The only paid tool was the AI coding assistant included in a $20 ChatGPT subscription, and the writer put the cost, "rounding very generously", at about $15 ( Build to Understand ).
Where vibe-coded games still break
The same sources that show how fast AI is also show where the time goes. These are the gaps we check first when someone sends us an AI-built game.
Sources for the table rows: Willison (October 28) , Willison (October 20) and Build to Understand . The last two rows are standard engineering practice rather than findings from those posts.
Willison makes the key point himself: personal games are "the perfect application for vibe coding because it's so low stakes". A brand game is not low stakes. A broken quest in a hobby project costs nothing, while a leaderboard someone cheats in a contest with real prizes costs money and goodwill.
Platform rules for AI-made games
Since January 2024, Steam's Content Survey has included an AI disclosure. It separates Pre-Generated content (art, code, sound and so on created with AI during development) from Live-Generated content (created while the game runs). For live generation you must also describe the guardrails that stop it producing illegal content, and Valve shows much of the disclosure on the store page ( Valve, AI Content on Steam ). The current Steamworks documentation adds that AI tools built into development environments, used for efficiency, are not the focus; the section is about AI content "that ships with your game, and is consumed by players" ( Steamworks: Content Survey ).
Google Play: in-app reporting for AI content
If your game generates content with AI while it runs (an AI character that chats with players, generated images or levels), Google Play requires "in-app user reporting or flagging features" so players can report offensive content without leaving the app ( Google Play Console Help ).
App Store: no copycats, no spam, in-app purchase for digital items
Apple's App Review Guidelines tell developers not to "simply copy the latest popular app on the App Store, or make some minor changes to another app's name or UI" (4.1 Copycats), not to flood the store with near-identical apps (4.3 Spam), and to offer "lasting entertainment value" (4.2). Game currencies, levels and premium content must be sold through in-app purchase (3.1.1) ( Apple App Review Guidelines ). A fast AI clone of a hit game runs into these rules even before any legal claim does.
Telegram: Stars for digital goods
Games inside Telegram must sell digital items in Telegram Stars ( Telegram Stars payments ). We cover the payment flow and launch options in our Telegram mini games guide .
Who owns an AI-generated game?
For a brand, this matters as much as the code. In January 2025 the US Copyright Office concluded that generative AI outputs "can be protected by copyright only where a human author has determined sufficient expressive elements", and that this does not include "the mere provision of prompts". Using AI to assist, or including AI material in a larger human-made work, does not by itself block protection ( U.S. Copyright Office ).
In practice: the more of a game's characters, writing and art comes from human creative decisions, and the better you document them, the stronger your claim to the parts that make the game yours. This is general information, not legal advice; for a campaign with a mascot or licensed brand assets, ask your lawyer.
Clones are the other side of ownership. Speaking to developers about AI-made copies, Tom Giardino of Valve's Steam business team said Valve does not want to dictate which tools studios use, and advised building a community early and continuing to improve the original, since individual mechanics are relatively easy to copy, while new features and a distinct visual style are harder to copy. He also acknowledged that building a community takes time and is not enough protection on its own ( ComputerBase, in German ).
Do it yourself, or hire a studio?
A prototype is never wasted. A version you have already played is a better specification than any document, and it makes a studio quote faster and more accurate.
How we use AI in our own games
We make two games of our own, Kosti, a 3D action game, and Block Nook, a block puzzle, both built in Godot 4 for Android and both still in development (see our Kosti and Block Nook project page ). AI does a lot of production work in them: 3D models and rigs from text and images, concept and interface art, and voice lines in three languages. What it does not do is decide. We set the style, fix what the tools get wrong and check every asset in the engine before it ships, because the gaps in the table above are the same ones we run into.
What a studio-built game costs
Our published starting prices for games are:
Branded Mini-Game from $8,000: advergames and playable campaigns.
Telegram Mini App Game from $10,000: games inside Telegram with leaderboards.
Mobile Game from $25,000: Android and iOS games, from concept to store.
"From" means the smallest sensible version; we send a fixed written quote before work starts, and new customers get 20% off the first project. If you already have a vibe-coded prototype, send it with your brief: we'll tell you what can be kept and what needs rebuilding. The full list is on our pricing page .
Can AI make a whole game from a prompt?
For small browser games, often yes. In October 2026 Simon Willison described getting a playable spot-the-difference game from a single prompt with no follow-up. Bigger games still need iteration: his fantasy RPG needed extra playtesting and prompting to fix quest bugs, and his strategy game took 12 major iterations and 321 commits by coding agents.
Do I have to disclose AI use when I publish a game on Steam?
Yes, for AI-made content that ships with the game. Steam's Content Survey asks you to describe AI content created during development (Pre-Generated) and AI content created while the game runs (Live-Generated), and Valve shows much of that disclosure on the store page. Valve says AI coding tools used for efficiency are not the focus of that section.
Do I own the copyright in an AI-generated game?
In the US, only partly. The Copyright Office concluded in January 2025 that AI outputs are protected only where a human author determined sufficient expressive elements, and that prompts alone are not enough. Human-made art, writing and creative changes to AI output can be protected. For a brand game, get legal advice and keep a record of the human work.
When is a vibe-coded game good enough for a brand campaign?
When nothing valuable rides on it: an internal demo, a pitch prototype or a short single-player toy with no prizes, no personal data and no payments. Once there are real prizes, a leaderboard people will try to cheat, sign-ups, purchases or a store release, the parts AI tools skip (server checks, testing on real phones, store rules and support) need engineering.
Building inside Telegram? Read our Telegram mini games guide . Built an app (not a game) with AI? See how to make an AI-built app production-ready .
Have a game idea or an AI prototype? Get a fixed quote.
Send us the prototype or a short brief. We reply within one business day with a fixed written quote: Branded Mini-Game from $8,000,

[truncated]
