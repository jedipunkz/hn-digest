---
source: "https://www.xda-developers.com/switched-to-claude-code-open-source-rival-and-stopped-worrying-about-rate-limits/"
hn_url: "https://news.ycombinator.com/item?id=49727345"
title: "Switched to Claude Code's open-source rival, stopped worrying about rate limits"
article_title: "I switched to Claude Code's open-source rival and stopped worrying about rate limits"
image: "https://static0.xdaimages.com/wordpress/wp-content/uploads/wm/2026/09/opencode-on-desktop-pc-clock-and-lamp-in-view.jpeg?w=1600&h=900&fit=crop"
author: "fork-bomber"
captured_at: "2026-09-16T14:39:29Z"
capture_tool: "hn-digest"
hn_id: 49727345
score: 1
comments: 0
posted_at: "2026-09-16T14:12:17Z"
tags:
  - hacker-news
---

# Switched to Claude Code's open-source rival, stopped worrying about rate limits

- HN: [49727345](https://news.ycombinator.com/item?id=49727345)
- Source: [www.xda-developers.com](https://www.xda-developers.com/switched-to-claude-code-open-source-rival-and-stopped-worrying-about-rate-limits/)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T14:12:17Z

## Translation

Title: Switched to Claude Code's open-source rival, stopped worrying about rate limits
Article title: I switched to Claude Code's open-source rival and stopped worrying about rate limits
Description: It's not about escaping Claude, it's about escaping counting

Article text:
I switched to Claude Code's open-source rival and stopped worrying about rate limits
Menu
Sign in
Sign in now
Close
News
Operating Systems
Submenu
Windows
Devices
Submenu
Single-Board Computers
Entertainment
Submenu
Android Auto
Like
Follow
Followed
16
16
More Action
Sign in
Sign in now
Claude
ESP32
AI Tools
Entertainment
Forums
Close
I switched to Claude Code's open-source rival and stopped worrying about rate limits
By
Nolen Jonker
Published Sep 15, 2026, 5:31 PM EDT
Nolen began their writing career in 2019, with three years dedicated to editing the Creative section at MakeUseOf . Their expertise lies at the crossroads of technology and creativity, covering areas like photography, video editing, and graphic design.
Outside of work, you'll often find Nolen diving into a good book, writing their own stories, or playing video games.
Sign in to your XDA account
Add Us
Add
on Google
Preferred Source
Google News
Like
Like
follow
Follow
followed
Followed
Thread
16
Log in
Here is a fact-based summary of the story contents:
Try something different:
Show me the facts
Explain it like I’m 5
Give me a lighthearted recap
Every paid Claude plan meters usage in two ways: a rolling 5-hour window and a weekly cap across all models, and that quota is shared with Claude chat and whatever else is running on the same account. So if you're a heavy Cowork or Claude Design user in addition to Claude Code, then you're out of luck or onto the Max plan.
The 5-hour window is the one that keeps catching me out , because it resets from your first prompt, not from midnight or any fixed time, so I'd lose track of when the clock actually started. A heavy morning in Claude chat could shrink my afternoon coding window without me realizing. Before rushing to Max, I started looking for alternatives that might reduce the heavy rates, and that's when I decided to try OpenCode…
Want to stay in the loop with the latest in AI? The XDA AI Insider newsletter drops weekly with deep dives, tool recommendations, and hands-on coverage you won't find anywhere else on the site. Subscribe by modifying your newsletter preferences !
The terminal scared me off at first
I just assumed a command-line tool wouldn't be for me
I don't code apart from adjusting some html from time to time, so every time I saw "terminal-based AI coding agent" in a headline, I assumed it wasn't for me and scrolled. Which, looking back, is a bit silly, because OpenCode is in the same category as Claude Code, and I'd been using Claude Code without a problem - but in the desktop GUI, whereas OpenCode is terminal-only.
The install was just one line pasted into the terminal, and opening OpenCode from my folder of choice. And then it's just a matter of navigating the keyboard commands to set yourself up with a provider, a model, enter different modes, and access the commands. In a way, this is actually much simpler than Claude's Code's GUI since there aren't a million-and-one things on the screen to distract me. I could definitely get used to this.
I installed it through npm with npm install -g opencode-ai , which needs Node.js on your machine. There's also a one-line curl install if you're on Mac or Linux ( curl -fsSL https://opencode.ai/install | bash ) but on Windows the npm route is cleaner.
Claude might not be the only coding agent worth using.
OpenCode's billing is different
And the model selection is wild
OpenCode is open-source , built by a small team that also runs the open-source dev framework SST, and it works with dozens of AI providers rather than being locked to one. So which model powers it is entirely up to you. Before getting started, you'll have to hit /connect to pick a provider, and the options here are kind of overwhelming. I went with OpenCode's Zen provider since I'm still new here and it was the recommendation. As for models, all the big names from Claude and Google are there, but since my goal was to reduce my rate limits, I went with OpenCode's own model entry, Big Pickle, to start with.
Claude Code Pro is a subscription, but it doesn't give you unlimited use of that subscription. It gives you roughly 90 prompts per rolling 5-hour window, plus a weekly cap across all models, and that pool is shared with everything else you do on Claude. Hit the wall and you have two choices: wait for the window to reset, or opt in to usage credits and keep working at standard API rates on top of what you're already paying monthly.
OpenCode doesn't have that mechanism at all; it's just a client. Whatever provider you connect through /connect sets the terms, and Zen works on a prepaid balance instead of a subscription window. You send prompts until your balance runs out, so there's no 5-hour timer, weekly cap, or annoying "you're about to run out" pop-ups. Zen will auto-reload $20 when you drop below $5, which I turned off because I wanted a hard cap.
OpenCode calls Big Pickle a "stealth model" which is free during its testing period. There is a catch, though. During the free period the data collected may be used to improve the model, which is a different tradeoff than paid Zen models where zero-retention is the norm. I'm fine with it given how I use it, which is simply experimenting with vibe-designing and also doing some computer admin. Others might not be, though, so it might be worth hooking it up to something else instead.
The question is what happens when Big Pickle isn't free anymore, because it won't be forever, and neither will other free-tier Zen models like DeepSeek V4 Flash Free or MiMo-V2.5 Free. The answer is that even on paid models, there's still no subscription window. You pay per token; prices range from cents to hundreds of dollars per million tokens depending on what you pick, and the worry moves to running out of balance rather than hitting a wall the app decides you should hit.
On Claude Code, rate limits are just a thing you're always sort of aware of. On OpenCode with a free model, they're not really a factor, and on a paid one it becomes more about watching your balance than getting cut off. There's no 5-hour window or weekly cap either way, you're just paying token rates directly, same as Claude Pro's extra usage but without the subscription on top of it.
One week, two tools, a lot of opinions.
Build mode and Plan mode are the reason exploration doesn't feel expensive anymore
Exploration used to cost me, now it doesn't
OpenCode has two built-in agents you swap between with the Tab key. Build is the default with full access, and it does what it sounds like - working directly with your files. Plan is read-only, it can look at your files but not touch them, and that restriction is actually enforced at the permission level rather than just in the system prompt, so the edit and write tools are hard-denied and the model can't modify a file even if I ask it to. This is basically read vs. write that you'll see in most GUI-based agentic tools.
Bash commands aren't fully locked down in Plan mode though. A read-only-looking command can still change remote state, something like gh issue comment will happily post to GitHub without touching a local file. So Plan mode isn't a total safety net, just a very good one for the specific class of thing it covers.
On Claude Code I'd think twice before asking "walk me through what this whole project does," because that's a big context read and I could feel the quota ticking. On OpenCode, with Big Pickle in particular, Plan mode is now my default entry point for anything unfamiliar. This is what I use to have scripts explained to me before I run them, what would happen if I change a config, and just general auditing of folders so I know exactly what to touch or not touch in them.
For vibe-coding specifically, Plan mode ends up functioning like a free consultation before every Build session starts. I get the AI to describe what it would change and why, then Tab into Build and let it execute the plan I already agreed to. For computer admin, which is my other main use, Plan mode basically just keeps me from breaking the system.
On a metered plan I'd never have spent prompts on exploratory questions like these. On OpenCode I front-load the understanding, then only Build once I know what I want.
Beat-for-beat, feature-for-feature.
The Claude rival I've been ignoring
I'm still in Claude nearly every day for my general work, so this isn't about breaking up with it. But for anything where I know I'm about to explore or get stuck, OpenCode is honestly the better option because its caps aren't nearly as restrictive (or expensive). Guess I never had reason to be scared of terminal-based tools after all.
Google is updating how content is shown. Don't miss our industry-leading content, written by humans, by setting XDA as a preferred source .
on Google
Preferred Source
Google News
Close
See more XDA stories on Google.
Add us on Google
Today's best deals
The Google Pixel 10a drops to $424, proving you don't need to spend $1,000 on a flagship
This $36 compact projector gives you a 200-inch display with autofocus and a three-year warranty
Garmin's Vivoactive 5 smartwatch just hit $182, nearly matching its all-time low
See More
Trending Now
One UI is quietly killing your Samsung battery — here's which 6 settings to disable
I stopped paying for ChatGPT after putting my local LLM on Tailscale
In the rush from Plex to Jellyfin, everyone forgot the media server that sits right in the middle
Thread
16
Sign in to your XDA account
We want to hear from you. Share your perspective in the comments below, and please keep the conversation respectful.
Attachment(s)
Please respect our community guidelines . No links, inappropriate language, or spam.
Your comment has not been saved
Jason
Jason
Jason
#DS966311
Member since 2026-06-23
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
OpenCode has a desktop app also
Christian
Christian
Christian
#WK796718
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Uh.. OpenCode isn't terminal only. There is OpenCode Desktop which is your traditional GUI too.
J
J
J
#CG088173
Member since 2026-03-05
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Opencode's free models have 100rpm / do have overall limits , theyre just generous
Mohammad Mesum
Mohammad Mesum
Mohammad Mesum
#ZZ406566
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I use Opencode with Gemini 3.1 Flash Lite (through API key no Zen required here) set to High Performance Mode.
With 1M token context window and 1000 requests per day limit (practically impossible to hit if you are a solo developer using it for personal use) it's basically a free model. The only drawback being that it's not as powerful as the latest and greatest coding agents (but that's fine if you are a student learning to code or a professional who knows his job but just requires some assistance).
Gemini 3.1 Flash Lite performs even better if you explicitly add some Coding Skills to Opencode to optimise the workflow.
The only issue with this setup is privacy..... So I wouldn't use this for projects which are supposed to be closed source (I usually use this for open source projects).
Eventually if some local model becomes good enough to replace Gemini 3.1 Flash Lite, I might go totally local with this Opencode setup.
medisano
medisano
medisano
#ZI329207
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Opencode is not only terminal.
Paul
Paul
Paul
#YE718337
Member since 2024-06-23
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
LM Studio is way better, you can also use the LocalMind app on your phone to access you LM server.
medisano
medisano
medisano
#ZI329207
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
No is not,unless you have some dgx or equivalent.cloud based llm are much faster and cheaper than local dgx.
Matt
Matt
Matt
#ZQ953615
Member since 2026-09-16
Following
0
Topics
0
Users
Foll

[truncated]

## Original Extract

It's not about escaping Claude, it's about escaping counting

I switched to Claude Code's open-source rival and stopped worrying about rate limits
Menu
Sign in
Sign in now
Close
News
Operating Systems
Submenu
Windows
Devices
Submenu
Single-Board Computers
Entertainment
Submenu
Android Auto
Like
Follow
Followed
16
16
More Action
Sign in
Sign in now
Claude
ESP32
AI Tools
Entertainment
Forums
Close
I switched to Claude Code's open-source rival and stopped worrying about rate limits
By
Nolen Jonker
Published Sep 15, 2026, 5:31 PM EDT
Nolen began their writing career in 2019, with three years dedicated to editing the Creative section at MakeUseOf . Their expertise lies at the crossroads of technology and creativity, covering areas like photography, video editing, and graphic design.
Outside of work, you'll often find Nolen diving into a good book, writing their own stories, or playing video games.
Sign in to your XDA account
Add Us
Add
on Google
Preferred Source
Google News
Like
Like
follow
Follow
followed
Followed
Thread
16
Log in
Here is a fact-based summary of the story contents:
Try something different:
Show me the facts
Explain it like I’m 5
Give me a lighthearted recap
Every paid Claude plan meters usage in two ways: a rolling 5-hour window and a weekly cap across all models, and that quota is shared with Claude chat and whatever else is running on the same account. So if you're a heavy Cowork or Claude Design user in addition to Claude Code, then you're out of luck or onto the Max plan.
The 5-hour window is the one that keeps catching me out , because it resets from your first prompt, not from midnight or any fixed time, so I'd lose track of when the clock actually started. A heavy morning in Claude chat could shrink my afternoon coding window without me realizing. Before rushing to Max, I started looking for alternatives that might reduce the heavy rates, and that's when I decided to try OpenCode…
Want to stay in the loop with the latest in AI? The XDA AI Insider newsletter drops weekly with deep dives, tool recommendations, and hands-on coverage you won't find anywhere else on the site. Subscribe by modifying your newsletter preferences !
The terminal scared me off at first
I just assumed a command-line tool wouldn't be for me
I don't code apart from adjusting some html from time to time, so every time I saw "terminal-based AI coding agent" in a headline, I assumed it wasn't for me and scrolled. Which, looking back, is a bit silly, because OpenCode is in the same category as Claude Code, and I'd been using Claude Code without a problem - but in the desktop GUI, whereas OpenCode is terminal-only.
The install was just one line pasted into the terminal, and opening OpenCode from my folder of choice. And then it's just a matter of navigating the keyboard commands to set yourself up with a provider, a model, enter different modes, and access the commands. In a way, this is actually much simpler than Claude's Code's GUI since there aren't a million-and-one things on the screen to distract me. I could definitely get used to this.
I installed it through npm with npm install -g opencode-ai , which needs Node.js on your machine. There's also a one-line curl install if you're on Mac or Linux ( curl -fsSL https://opencode.ai/install | bash ) but on Windows the npm route is cleaner.
Claude might not be the only coding agent worth using.
OpenCode's billing is different
And the model selection is wild
OpenCode is open-source , built by a small team that also runs the open-source dev framework SST, and it works with dozens of AI providers rather than being locked to one. So which model powers it is entirely up to you. Before getting started, you'll have to hit /connect to pick a provider, and the options here are kind of overwhelming. I went with OpenCode's Zen provider since I'm still new here and it was the recommendation. As for models, all the big names from Claude and Google are there, but since my goal was to reduce my rate limits, I went with OpenCode's own model entry, Big Pickle, to start with.
Claude Code Pro is a subscription, but it doesn't give you unlimited use of that subscription. It gives you roughly 90 prompts per rolling 5-hour window, plus a weekly cap across all models, and that pool is shared with everything else you do on Claude. Hit the wall and you have two choices: wait for the window to reset, or opt in to usage credits and keep working at standard API rates on top of what you're already paying monthly.
OpenCode doesn't have that mechanism at all; it's just a client. Whatever provider you connect through /connect sets the terms, and Zen works on a prepaid balance instead of a subscription window. You send prompts until your balance runs out, so there's no 5-hour timer, weekly cap, or annoying "you're about to run out" pop-ups. Zen will auto-reload $20 when you drop below $5, which I turned off because I wanted a hard cap.
OpenCode calls Big Pickle a "stealth model" which is free during its testing period. There is a catch, though. During the free period the data collected may be used to improve the model, which is a different tradeoff than paid Zen models where zero-retention is the norm. I'm fine with it given how I use it, which is simply experimenting with vibe-designing and also doing some computer admin. Others might not be, though, so it might be worth hooking it up to something else instead.
The question is what happens when Big Pickle isn't free anymore, because it won't be forever, and neither will other free-tier Zen models like DeepSeek V4 Flash Free or MiMo-V2.5 Free. The answer is that even on paid models, there's still no subscription window. You pay per token; prices range from cents to hundreds of dollars per million tokens depending on what you pick, and the worry moves to running out of balance rather than hitting a wall the app decides you should hit.
On Claude Code, rate limits are just a thing you're always sort of aware of. On OpenCode with a free model, they're not really a factor, and on a paid one it becomes more about watching your balance than getting cut off. There's no 5-hour window or weekly cap either way, you're just paying token rates directly, same as Claude Pro's extra usage but without the subscription on top of it.
One week, two tools, a lot of opinions.
Build mode and Plan mode are the reason exploration doesn't feel expensive anymore
Exploration used to cost me, now it doesn't
OpenCode has two built-in agents you swap between with the Tab key. Build is the default with full access, and it does what it sounds like - working directly with your files. Plan is read-only, it can look at your files but not touch them, and that restriction is actually enforced at the permission level rather than just in the system prompt, so the edit and write tools are hard-denied and the model can't modify a file even if I ask it to. This is basically read vs. write that you'll see in most GUI-based agentic tools.
Bash commands aren't fully locked down in Plan mode though. A read-only-looking command can still change remote state, something like gh issue comment will happily post to GitHub without touching a local file. So Plan mode isn't a total safety net, just a very good one for the specific class of thing it covers.
On Claude Code I'd think twice before asking "walk me through what this whole project does," because that's a big context read and I could feel the quota ticking. On OpenCode, with Big Pickle in particular, Plan mode is now my default entry point for anything unfamiliar. This is what I use to have scripts explained to me before I run them, what would happen if I change a config, and just general auditing of folders so I know exactly what to touch or not touch in them.
For vibe-coding specifically, Plan mode ends up functioning like a free consultation before every Build session starts. I get the AI to describe what it would change and why, then Tab into Build and let it execute the plan I already agreed to. For computer admin, which is my other main use, Plan mode basically just keeps me from breaking the system.
On a metered plan I'd never have spent prompts on exploratory questions like these. On OpenCode I front-load the understanding, then only Build once I know what I want.
Beat-for-beat, feature-for-feature.
The Claude rival I've been ignoring
I'm still in Claude nearly every day for my general work, so this isn't about breaking up with it. But for anything where I know I'm about to explore or get stuck, OpenCode is honestly the better option because its caps aren't nearly as restrictive (or expensive). Guess I never had reason to be scared of terminal-based tools after all.
Google is updating how content is shown. Don't miss our industry-leading content, written by humans, by setting XDA as a preferred source .
on Google
Preferred Source
Google News
Close
See more XDA stories on Google.
Add us on Google
Today's best deals
The Google Pixel 10a drops to $424, proving you don't need to spend $1,000 on a flagship
This $36 compact projector gives you a 200-inch display with autofocus and a three-year warranty
Garmin's Vivoactive 5 smartwatch just hit $182, nearly matching its all-time low
See More
Trending Now
One UI is quietly killing your Samsung battery — here's which 6 settings to disable
I stopped paying for ChatGPT after putting my local LLM on Tailscale
In the rush from Plex to Jellyfin, everyone forgot the media server that sits right in the middle
Thread
16
Sign in to your XDA account
We want to hear from you. Share your perspective in the comments below, and please keep the conversation respectful.
Attachment(s)
Please respect our community guidelines . No links, inappropriate language, or spam.
Your comment has not been saved
Jason
Jason
Jason
#DS966311
Member since 2026-06-23
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
OpenCode has a desktop app also
Christian
Christian
Christian
#WK796718
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Uh.. OpenCode isn't terminal only. There is OpenCode Desktop which is your traditional GUI too.
J
J
J
#CG088173
Member since 2026-03-05
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Opencode's free models have 100rpm / do have overall limits , theyre just generous
Mohammad Mesum
Mohammad Mesum
Mohammad Mesum
#ZZ406566
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I use Opencode with Gemini 3.1 Flash Lite (through API key no Zen required here) set to High Performance Mode.
With 1M token context window and 1000 requests per day limit (practically impossible to hit if you are a solo developer using it for personal use) it's basically a free model. The only drawback being that it's not as powerful as the latest and greatest coding agents (but that's fine if you are a student learning to code or a professional who knows his job but just requires some assistance).
Gemini 3.1 Flash Lite performs even better if you explicitly add some Coding Skills to Opencode to optimise the workflow.
The only issue with this setup is privacy..... So I wouldn't use this for projects which are supposed to be closed source (I usually use this for open source projects).
Eventually if some local model becomes good enough to replace Gemini 3.1 Flash Lite, I might go totally local with this Opencode setup.
medisano
medisano
medisano
#ZI329207
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Opencode is not only terminal.
Paul
Paul
Paul
#YE718337
Member since 2024-06-23
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
LM Studio is way better, you can also use the LocalMind app on your phone to access you LM server.
medisano
medisano
medisano
#ZI329207
Member since 2026-09-16
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
No is not,unless you have some dgx or equivalent.cloud based llm are much faster and cheaper than local dgx.
Matt
Matt
Matt
#ZQ953615
Member since 2026-09-16
Following
0
Topics
0
Users
Foll

[truncated]
