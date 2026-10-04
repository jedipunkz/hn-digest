---
source: "https://github.com/vincentsch/explainroo"
hn_url: "https://news.ycombinator.com/item?id=49956470"
title: "Explainer videos and product demos made by your AI agent. Free and open source"
article_title: "GitHub - vincentsch/explainroo: Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4. · GitHub"
image: "https://opengraph.githubassets.com/24f884b46703889d312d3cb56fb55e1be6854f0446e22db4465906552720936d/vincentsch/explainroo"
author: "thunderbong"
captured_at: "2026-10-04T19:14:57Z"
capture_tool: "hn-digest"
hn_id: 49956470
score: 1
comments: 0
posted_at: "2026-10-04T18:24:39Z"
tags:
  - hacker-news
---

# Explainer videos and product demos made by your AI agent. Free and open source

- HN: [49956470](https://news.ycombinator.com/item?id=49956470)
- Source: [github.com](https://github.com/vincentsch/explainroo)
- Score: 1
- Comments: 0
- Posted: 2026-10-04T18:24:39Z

## Translation

Title: Explainer videos and product demos made by your AI agent. Free and open source
Article title: GitHub - vincentsch/explainroo: Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4. · GitHub
Description: Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4. - vincentsch/explainroo

Article text:
GitHub - vincentsch/explainroo: Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4. · GitHub
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
vincentsch
/
explainroo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
30 Commits 30 Commits Folders and files
bin bin dev dev docs/ media docs/ media engine engine examples examples fonts fonts scripts scripts skills/ explainroo skills/ explainroo src src templates/ starter templates/ starter test test .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Explainer videos made by your AI agent.
Free and open source. The voice, the timing and the rendering run on your computer.
Website ·
Example videos ·
Docs ·
AGENTS.md
explainroo-intro-readme.mp4
A coding agent made this video with explainroo. You can also watch it on explainroo.com .
Give your coding agent this repo and tell it what the video should explain.
It works with Claude Code, Codex, Pi and other coding agents. Right now it
works best with Claude Code and Opus 5.5.
Copy this into your agent and put your topic in place of the brackets:
Make me a short explainer video about [your topic].
Use explainroo for it: clone https://github.com/vincentsch/explainroo,
read its AGENTS.md and follow the steps.
The agent sets up explainroo, makes the video and checks it. You get an MP4
file.
An explainer video is a short video where a voice explains a topic and
drawings appear while it speaks. For explainroo, the agent writes two files.
script.md has the words the voice says. scenes.js draws the pictures with
a bit of JavaScript, and each drawing can appear on a word from the script.
explainroo does the rest:
Voice. Kokoro , an open
voice model, reads the script aloud. It has 28 voices and needs no account
or API key.
Timing. Whisper listens to the
recording and notes when each word is spoken.
Pictures. Chrome runs in the background and draws the frames. The lines
can look hand drawn ( Rough.js ), and there are 1,800
icons from Lucide . A scene can also show charts, code
or your own screenshots.
Sound. explainroo makes its own background music for each video and
adds small sound effects. The music gets quieter while the voice speaks.
File. ffmpeg puts it all together into an MP4.
An agent can't watch a video, so explainroo gives it other ways to check its
work. It saves stills of the scenes and a sheet of small frames for the whole
video. A layout check finds text that is cut off or overlaps, and a speech
check finds words the voice got wrong. For a feed-sized player,
explainroo check videos/<name> --view-width 854 also flags small text.
None of this leaves your computer, and it costs nothing. Your coding agent is
a separate service with its own terms and prices. If you want, the agent can
also make illustrations with an AI image model through OpenRouter, and you pay
for each image.
There are five looks: paper, clean, chalk, blueprint and midnight. You change
a video's look with one setting in video.json . Here is one frame in each
look, from the example videos.
You pick the size for the place the video goes. YouTube videos are wide, and
Shorts, TikTok and Reels are tall. Instagram and LinkedIn posts use 4:5, and
there is a square size too. Shorts, TikTok and Reels put their own buttons
over the video, and explainroo keeps your text out of those spots.
Tall, 4:5 and square videos get captions that light up word by word.
"pace": 1.2 in a video's video.json makes the voice, the pauses and the
animations 20% quicker. The music gets a little quicker too. Pace goes from
0.7 to 1.6, and 1 is normal.
explainroo can also make product demos, the videos software companies make to
show their app. The agent rebuilds the app's screens from its code or its
website, with the same colors, fonts and button labels. A mouse pointer then
clicks through the screens and types into the fields. Tell your agent which
product it is and where to find its code or website.
Here are the demos for Unspar
and Vroni , two of my
own products. The files for the Unspar demo are in
examples/unspar-demo . explainroo.com also has
unofficial demos of Gmail, ChatGPT and Claude .
You need Node.js 20.11 or newer, ffmpeg, and Chrome or Chromium. The first
setup downloads the voice and timing models once, about 400 MB together. You
don't need a graphics card. I develop and test explainroo on Linux. It should
work on macOS and Windows, but I have tested it less there.
git clone https://github.com/vincentsch/explainroo.git
cd explainroo
npm install
node bin/explainroo.js doctor --fetch
Then start your agent in the explainroo folder and ask for a video, for
example "Make a 60 second video about how HTTPS keeps a password secret." The
agent follows AGENTS.md and saves the video as
videos/<name>/out/video.mp4 . The example videos are in
examples/ , and the full documentation is on
explainroo.com .
Each video has a small "explainroo.com" in one corner. "watermark": false in
video.json turns it off. Please keep it if you can, because that is how
other people find explainroo.
explainroo is MIT licensed. The voice comes from
Kokoro , and the word timing from
Whisper through
Transformers.js . The
drawings use Rough.js and Lucide
icons. Playwright runs Chrome, and
ffmpeg makes the video file. The fonts are under the SIL
Open Font License.
Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4.
Readme MIT license Activity Stars
11 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4. - vincentsch/explainroo

GitHub - vincentsch/explainroo: Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4. · GitHub
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
vincentsch
/
explainroo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
30 Commits 30 Commits Folders and files
bin bin dev dev docs/ media docs/ media engine engine examples examples fonts fonts scripts scripts skills/ explainroo skills/ explainroo src src templates/ starter templates/ starter test test .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Explainer videos made by your AI agent.
Free and open source. The voice, the timing and the rendering run on your computer.
Website ·
Example videos ·
Docs ·
AGENTS.md
explainroo-intro-readme.mp4
A coding agent made this video with explainroo. You can also watch it on explainroo.com .
Give your coding agent this repo and tell it what the video should explain.
It works with Claude Code, Codex, Pi and other coding agents. Right now it
works best with Claude Code and Opus 5.5.
Copy this into your agent and put your topic in place of the brackets:
Make me a short explainer video about [your topic].
Use explainroo for it: clone https://github.com/vincentsch/explainroo,
read its AGENTS.md and follow the steps.
The agent sets up explainroo, makes the video and checks it. You get an MP4
file.
An explainer video is a short video where a voice explains a topic and
drawings appear while it speaks. For explainroo, the agent writes two files.
script.md has the words the voice says. scenes.js draws the pictures with
a bit of JavaScript, and each drawing can appear on a word from the script.
explainroo does the rest:
Voice. Kokoro , an open
voice model, reads the script aloud. It has 28 voices and needs no account
or API key.
Timing. Whisper listens to the
recording and notes when each word is spoken.
Pictures. Chrome runs in the background and draws the frames. The lines
can look hand drawn ( Rough.js ), and there are 1,800
icons from Lucide . A scene can also show charts, code
or your own screenshots.
Sound. explainroo makes its own background music for each video and
adds small sound effects. The music gets quieter while the voice speaks.
File. ffmpeg puts it all together into an MP4.
An agent can't watch a video, so explainroo gives it other ways to check its
work. It saves stills of the scenes and a sheet of small frames for the whole
video. A layout check finds text that is cut off or overlaps, and a speech
check finds words the voice got wrong. For a feed-sized player,
explainroo check videos/<name> --view-width 854 also flags small text.
None of this leaves your computer, and it costs nothing. Your coding agent is
a separate service with its own terms and prices. If you want, the agent can
also make illustrations with an AI image model through OpenRouter, and you pay
for each image.
There are five looks: paper, clean, chalk, blueprint and midnight. You change
a video's look with one setting in video.json . Here is one frame in each
look, from the example videos.
You pick the size for the place the video goes. YouTube videos are wide, and
Shorts, TikTok and Reels are tall. Instagram and LinkedIn posts use 4:5, and
there is a square size too. Shorts, TikTok and Reels put their own buttons
over the video, and explainroo keeps your text out of those spots.
Tall, 4:5 and square videos get captions that light up word by word.
"pace": 1.2 in a video's video.json makes the voice, the pauses and the
animations 20% quicker. The music gets a little quicker too. Pace goes from
0.7 to 1.6, and 1 is normal.
explainroo can also make product demos, the videos software companies make to
show their app. The agent rebuilds the app's screens from its code or its
website, with the same colors, fonts and button labels. A mouse pointer then
clicks through the screens and types into the fields. Tell your agent which
product it is and where to find its code or website.
Here are the demos for Unspar
and Vroni , two of my
own products. The files for the Unspar demo are in
examples/unspar-demo . explainroo.com also has
unofficial demos of Gmail, ChatGPT and Claude .
You need Node.js 20.11 or newer, ffmpeg, and Chrome or Chromium. The first
setup downloads the voice and timing models once, about 400 MB together. You
don't need a graphics card. I develop and test explainroo on Linux. It should
work on macOS and Windows, but I have tested it less there.
git clone https://github.com/vincentsch/explainroo.git
cd explainroo
npm install
node bin/explainroo.js doctor --fetch
Then start your agent in the explainroo folder and ask for a video, for
example "Make a 60 second video about how HTTPS keeps a password secret." The
agent follows AGENTS.md and saves the video as
videos/<name>/out/video.mp4 . The example videos are in
examples/ , and the full documentation is on
explainroo.com .
Each video has a small "explainroo.com" in one corner. "watermark": false in
video.json turns it off. Please keep it if you can, because that is how
other people find explainroo.
explainroo is MIT licensed. The voice comes from
Kokoro , and the word timing from
Whisper through
Transformers.js . The
drawings use Rough.js and Lucide
icons. Playwright runs Chrome, and
ffmpeg makes the video file. The fonts are under the SIL
Open Font License.
Explainer videos and product demos made by your AI agent. Free and open source: a local voice (Kokoro), word timing (Whisper) and a canvas renderer turn a script into a narrated MP4.
Readme MIT license Activity Stars
11 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
