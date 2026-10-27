---
source: "https://github.com/yudhanjaya/Zest"
hn_url: "https://news.ycombinator.com/item?id=50258702"
title: "Friendly UI and server and agent for local LLMs"
article_title: "GitHub - yudhanjaya/Zest: A friendly console for local AI: chat, agent, models, GPU selection and backends · GitHub"
image: "https://opengraph.githubassets.com/497b3f84f61670f41079d3a45cdd8ad2e004a2d589bb52cd0ba907ff49cc560c/yudhanjaya/Zest"
author: "icaruswept"
captured_at: "2026-10-27T14:02:50Z"
capture_tool: "hn-digest"
hn_id: 50258702
score: 1
comments: 1
posted_at: "2026-10-27T13:13:05Z"
tags:
  - hacker-news
---

# Friendly UI and server and agent for local LLMs

- HN: [50258702](https://news.ycombinator.com/item?id=50258702)
- Source: [github.com](https://github.com/yudhanjaya/Zest)
- Score: 1
- Comments: 1
- Posted: 2026-10-27T13:13:05Z

## Translation

Title: Friendly UI and server and agent for local LLMs
Article title: GitHub - yudhanjaya/Zest: A friendly console for local AI: chat, agent, models, GPU selection and backends · GitHub
Description: A friendly console for local AI: chat, agent, models, GPU selection and backends - yudhanjaya/Zest

Article text:
GitHub - yudhanjaya/Zest: A friendly console for local AI: chat, agent, models, GPU selection and backends · GitHub
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
yudhanjaya
Zest Public Notifications You must be signed in to change notification settings
Star 11 ( 11 ) You must be signed in to star a repository
A friendly console for local AI: chat, agent, models, GPU selection and backends
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
34 Commits 34 Commits Folders and files
.github/ workflows .github/ workflows docs docs packaging packaging zest zest .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Run AI models on your own computer, and put them to work, with one download. Zest is a friendly interface for Lemonade , the open-source engine that runs language models locally, and for Pi , a coding-agent harness. The Zest AppImage includes both . Linux-only for now.
The idea is to provide a user interface that people who aren't necessarily steeped in technical know-how can still navigate to get open-weights models running on their own hardware, and then do real work with them: chat, draw diagrams, write documents, or let a model read and change the files in a folder while they watch. Download Zest, and you're good to go.
Chat, working in a folder: it read the code, fixed a bug with a one-line change (shown as a diff) and ran the tests again, step by step.
Contents: What you can do · Install and run · Finding your way around · Your first conversation · Choosing and managing models · Skills · Use your models in other apps · Keeping everything up to date · Advanced tools · Themes · Your data and privacy · For developers
Chat with an AI that runs on your own computer. Nothing you type leaves it, unless you let the model search the web: it asks first, every time.
Let it work on your files. Choose a folder and the model can read, search, edit and run things there while you watch. It asks before it changes anything.
Hand it a bigger job. A job keeps working in the background, round after round, and checks its own progress. You review what it changed before anything touches your folder.
Draw diagrams and write documents. Ask for a picture or a README and get one, shown right in the chat.
Teach it new tricks with skills. Short instruction sheets that make small models far more dependable. Add more by pasting a link.
Pick and manage models. Zest tells you which ones will run comfortably on your computer and downloads them for you.
Make it look the way you like. Fifteen themes, from soft plastic to Dracula and Nord, under Advanced > Themes .
Use your models in other apps. Zest can share them with tools such as opencode or any app that talks to an OpenAI-style server.
Zest runs on Linux. It needs no technical setup: you download one file and follow a few short steps. The Docs tab inside the app explains everything in plain language (its text lives in zest/docs.md ).
Download Zest-<version>-x86_64.AppImage from the latest release .
chmod +x Zest- * -x86_64.AppImage
./Zest- * -x86_64.AppImage
Follow the welcome steps. Zest opens in your browser and walks you through four short steps in plain language, one screen at a time:
The welcome steps finish with a button that opens the chat.
People who already have models, or run their own Lemonade, can skip to Advanced > Setup guide , which keeps the full set of choices: where models are stored, models you already have (LM Studio, Ollama, llama.cpp and others are found for you), and Lemonade updates.
Linux, x86_64. The built-in Lemonade needs a reasonably recent system (glibc 2.38 or newer: Ubuntu 24.04, Debian 13, Fedora 39 or later). On an older system Zest tells you so and you can connect it to a Lemonade you install yourself instead.
A few common system libraries that almost every desktop Linux already has (OpenSSL 3, zstd, libdrm). If one is missing, Zest names it.
For GPU use, your graphics driver. Zest and Lemonade download everything else.
Internet access for the first-run downloads (the backend and your models).
If your system has no FUSE support, run ./Zest-*-x86_64.AppImage --appimage-extract-and-run instead. Closing the browser tab quits Zest (and its Lemonade) after about three minutes; launching it again while it's running just reopens the page.
The left-hand menu has six places. Everything a newcomer needs is in the first five; everything technical sits behind Advanced , and so does the choice of theme.
The menu on the left, with the Docs tab open.
Open Chat and type. The bar under the title shows which model will answer and whether it is in memory. If it isn't, Zest loads it when you send your first message, so that answer takes a little longer.
Chat runs on Pi , a coding-agent harness that ships inside Zest. Every chat is a co-working session: the model reads, searches, edits and runs commands (it has a terminal) in a folder while you watch. Pick the folder, or leave it and it works in your home folder, and says so at the top of the chat; there it starts on Ask first , and jobs need a project folder. It answers questions like a normal chat too, and asks before it changes anything (tool cards show each step; edits show a diff). Choose how careful it is per chat ( Ask first , Edit freely , Do everything , Look only ), add more folders it can look at or change, switch models, watch how much of the context window is used, and add notes while it works.
For bigger work, press Start a job… . Say what should get done and how Zest can tell it is done, such as a test command that passes. The job runs in the background, round after round, and stops when the check passes, when it hits its round or time limit, or when it stops making progress. By default it works on a copy of your folder; the Changes tab shows what it changed and you apply what you want. Jobs are listed under their chat, and you can keep chatting while they work. Chats and jobs are saved and pick up where they left off. See docs/COWORK.md .
Replies stream in as they are written. The bar under the title shows which model answers, how much context it has (128k by default) and whether it is in memory, with Load , Unload and Reload . A reasoning model's thinking appears live above its answer and folds away when it is done. Markdown and code blocks come with copy buttons, and a small strip shows what the model is doing: 📖 reading, 🤔 thinking, ✍️ writing, 🛠️ working.
Every chat works in a folder (your home folder unless you choose one) and can read, search, edit and run things there. Tool cards show each step, edits show a diff, and in Ask first and Edit freely modes it stops and asks before it runs a command. Here it was asked to fix a folder whose tests were failing: it read the tests and the code, made a one-line change and ran the tests again to confirm.
Lemonade handles models; Zest gives you the choices.
Where they live. The default is the Hugging Face cache ( ~/.cache/huggingface/hub , or wherever HF_HOME points), so downloads are shared with other AI tools. Or choose any folder as the place new downloads go.
Models you already have. Look on this computer checks the usual places (LM Studio, llama.cpp's cache, GPT4All, Jan, Ollama, ~/models , ~/Downloads ) and lists what it finds. Tick what you want, or add any folder yourself. Nothing is copied : Zest gives Lemonade a folder of links to your files, so there's no extra disk use and your originals stay where they are.
Downloading. Press Get it on any model, or Get from Hugging Face to paste a link to a GGUF model and choose a version. It runs in the background with a progress bar, and you can keep using Zest. Zest's suggestions are the models that fit your hardware's free memory.
Models : shows you what models you can run and one sentence on whether it will run comfortably on your computer. The fit comes from the memory that is actually free on your graphics cards right now, so you can tell before loading whether another program is in the way. The default page, My picks , leads with the models I've tested and recommend. More models shows the rest of what Lemonade offers that will run well on your computer. A Details fold on each card holds the technical names and Settings for this model : context window, backend, devices and the llama-server options (flash attention, cache precision, expert layers in system RAM, GPU splitting, thinking, speculative decoding) as tick boxes and dropdowns with plain explanations, checked against your own llama-server before you load.
Context window: 128k by default. Chat always loads a model with a 128k (131,072-token) context window, a good threshold for usefulness in long chats and agentic work. The context menu in the chat's model bar changes it for every chat. A size saved under Settings for this model is used when other apps load the model through the Server page. A bigger window needs more graphics memory.
How do I choose a model? A button on Models (and an entry in Help) explains how to pick: will it fit, what is it for, how much it needs to remember. It then links to places to compare and learn about various models: Luke's Dev Lab, Artificial Analysis, LMArena, Aider, the Berkeley Function-Calling Leaderboard, SWE-bench, LiveBench, BenchLM, LLM Stats, r/LocalLLaMA and the Level1Techs forum.
Load, unload and reload live in the chat. A bar under the chat's title shows the model, the context size and whether the model is in memory, with Load , Unload and Reload next to them. Pick a different model there, or change the context size (128k by default; the choice is remembered) and a loaded model is reloaded at the new size. Sending a message to a model that isn't in memory loads it for you, at that size. The Models page is for downloading models and editing their settings: each downloaded card has Chat with it and Settings . Delete from this computer is under Details ; it asks first, says how much disk it frees, unloads the model if it is running, and you can download it again later.
Get from Hugging Face takes a link or org/name , lists the GGUF versions in the repository with their sizes and a plain label (for example "A good balance"), marks which ones fit your computer, and hands the download to Lemonade.
The Models page is for downloading models and editing how they run. My picks leads with the models I've tested and recommend; On this computer lists what you have. The meter on each card shows what will fit in the memory that is free on your graphics cards right now; if another program is using it, Zest warns you and names it.
Start here: LFM2.5 8B A1B. Very fast, runs on almost anything, calls tools well, and keeps answers short.
My general pick: Qwen3.6 35

[truncated]

## Original Extract

A friendly console for local AI: chat, agent, models, GPU selection and backends - yudhanjaya/Zest

GitHub - yudhanjaya/Zest: A friendly console for local AI: chat, agent, models, GPU selection and backends · GitHub
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
yudhanjaya
Zest Public Notifications You must be signed in to change notification settings
Star 11 ( 11 ) You must be signed in to star a repository
A friendly console for local AI: chat, agent, models, GPU selection and backends
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
34 Commits 34 Commits Folders and files
.github/ workflows .github/ workflows docs docs packaging packaging zest zest .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Run AI models on your own computer, and put them to work, with one download. Zest is a friendly interface for Lemonade , the open-source engine that runs language models locally, and for Pi , a coding-agent harness. The Zest AppImage includes both . Linux-only for now.
The idea is to provide a user interface that people who aren't necessarily steeped in technical know-how can still navigate to get open-weights models running on their own hardware, and then do real work with them: chat, draw diagrams, write documents, or let a model read and change the files in a folder while they watch. Download Zest, and you're good to go.
Chat, working in a folder: it read the code, fixed a bug with a one-line change (shown as a diff) and ran the tests again, step by step.
Contents: What you can do · Install and run · Finding your way around · Your first conversation · Choosing and managing models · Skills · Use your models in other apps · Keeping everything up to date · Advanced tools · Themes · Your data and privacy · For developers
Chat with an AI that runs on your own computer. Nothing you type leaves it, unless you let the model search the web: it asks first, every time.
Let it work on your files. Choose a folder and the model can read, search, edit and run things there while you watch. It asks before it changes anything.
Hand it a bigger job. A job keeps working in the background, round after round, and checks its own progress. You review what it changed before anything touches your folder.
Draw diagrams and write documents. Ask for a picture or a README and get one, shown right in the chat.
Teach it new tricks with skills. Short instruction sheets that make small models far more dependable. Add more by pasting a link.
Pick and manage models. Zest tells you which ones will run comfortably on your computer and downloads them for you.
Make it look the way you like. Fifteen themes, from soft plastic to Dracula and Nord, under Advanced > Themes .
Use your models in other apps. Zest can share them with tools such as opencode or any app that talks to an OpenAI-style server.
Zest runs on Linux. It needs no technical setup: you download one file and follow a few short steps. The Docs tab inside the app explains everything in plain language (its text lives in zest/docs.md ).
Download Zest-<version>-x86_64.AppImage from the latest release .
chmod +x Zest- * -x86_64.AppImage
./Zest- * -x86_64.AppImage
Follow the welcome steps. Zest opens in your browser and walks you through four short steps in plain language, one screen at a time:
The welcome steps finish with a button that opens the chat.
People who already have models, or run their own Lemonade, can skip to Advanced > Setup guide , which keeps the full set of choices: where models are stored, models you already have (LM Studio, Ollama, llama.cpp and others are found for you), and Lemonade updates.
Linux, x86_64. The built-in Lemonade needs a reasonably recent system (glibc 2.38 or newer: Ubuntu 24.04, Debian 13, Fedora 39 or later). On an older system Zest tells you so and you can connect it to a Lemonade you install yourself instead.
A few common system libraries that almost every desktop Linux already has (OpenSSL 3, zstd, libdrm). If one is missing, Zest names it.
For GPU use, your graphics driver. Zest and Lemonade download everything else.
Internet access for the first-run downloads (the backend and your models).
If your system has no FUSE support, run ./Zest-*-x86_64.AppImage --appimage-extract-and-run instead. Closing the browser tab quits Zest (and its Lemonade) after about three minutes; launching it again while it's running just reopens the page.
The left-hand menu has six places. Everything a newcomer needs is in the first five; everything technical sits behind Advanced , and so does the choice of theme.
The menu on the left, with the Docs tab open.
Open Chat and type. The bar under the title shows which model will answer and whether it is in memory. If it isn't, Zest loads it when you send your first message, so that answer takes a little longer.
Chat runs on Pi , a coding-agent harness that ships inside Zest. Every chat is a co-working session: the model reads, searches, edits and runs commands (it has a terminal) in a folder while you watch. Pick the folder, or leave it and it works in your home folder, and says so at the top of the chat; there it starts on Ask first , and jobs need a project folder. It answers questions like a normal chat too, and asks before it changes anything (tool cards show each step; edits show a diff). Choose how careful it is per chat ( Ask first , Edit freely , Do everything , Look only ), add more folders it can look at or change, switch models, watch how much of the context window is used, and add notes while it works.
For bigger work, press Start a job… . Say what should get done and how Zest can tell it is done, such as a test command that passes. The job runs in the background, round after round, and stops when the check passes, when it hits its round or time limit, or when it stops making progress. By default it works on a copy of your folder; the Changes tab shows what it changed and you apply what you want. Jobs are listed under their chat, and you can keep chatting while they work. Chats and jobs are saved and pick up where they left off. See docs/COWORK.md .
Replies stream in as they are written. The bar under the title shows which model answers, how much context it has (128k by default) and whether it is in memory, with Load , Unload and Reload . A reasoning model's thinking appears live above its answer and folds away when it is done. Markdown and code blocks come with copy buttons, and a small strip shows what the model is doing: 📖 reading, 🤔 thinking, ✍️ writing, 🛠️ working.
Every chat works in a folder (your home folder unless you choose one) and can read, search, edit and run things there. Tool cards show each step, edits show a diff, and in Ask first and Edit freely modes it stops and asks before it runs a command. Here it was asked to fix a folder whose tests were failing: it read the tests and the code, made a one-line change and ran the tests again to confirm.
Lemonade handles models; Zest gives you the choices.
Where they live. The default is the Hugging Face cache ( ~/.cache/huggingface/hub , or wherever HF_HOME points), so downloads are shared with other AI tools. Or choose any folder as the place new downloads go.
Models you already have. Look on this computer checks the usual places (LM Studio, llama.cpp's cache, GPT4All, Jan, Ollama, ~/models , ~/Downloads ) and lists what it finds. Tick what you want, or add any folder yourself. Nothing is copied : Zest gives Lemonade a folder of links to your files, so there's no extra disk use and your originals stay where they are.
Downloading. Press Get it on any model, or Get from Hugging Face to paste a link to a GGUF model and choose a version. It runs in the background with a progress bar, and you can keep using Zest. Zest's suggestions are the models that fit your hardware's free memory.
Models : shows you what models you can run and one sentence on whether it will run comfortably on your computer. The fit comes from the memory that is actually free on your graphics cards right now, so you can tell before loading whether another program is in the way. The default page, My picks , leads with the models I've tested and recommend. More models shows the rest of what Lemonade offers that will run well on your computer. A Details fold on each card holds the technical names and Settings for this model : context window, backend, devices and the llama-server options (flash attention, cache precision, expert layers in system RAM, GPU splitting, thinking, speculative decoding) as tick boxes and dropdowns with plain explanations, checked against your own llama-server before you load.
Context window: 128k by default. Chat always loads a model with a 128k (131,072-token) context window, a good threshold for usefulness in long chats and agentic work. The context menu in the chat's model bar changes it for every chat. A size saved under Settings for this model is used when other apps load the model through the Server page. A bigger window needs more graphics memory.
How do I choose a model? A button on Models (and an entry in Help) explains how to pick: will it fit, what is it for, how much it needs to remember. It then links to places to compare and learn about various models: Luke's Dev Lab, Artificial Analysis, LMArena, Aider, the Berkeley Function-Calling Leaderboard, SWE-bench, LiveBench, BenchLM, LLM Stats, r/LocalLLaMA and the Level1Techs forum.
Load, unload and reload live in the chat. A bar under the chat's title shows the model, the context size and whether the model is in memory, with Load , Unload and Reload next to them. Pick a different model there, or change the context size (128k by default; the choice is remembered) and a loaded model is reloaded at the new size. Sending a message to a model that isn't in memory loads it for you, at that size. The Models page is for downloading models and editing their settings: each downloaded card has Chat with it and Settings . Delete from this computer is under Details ; it asks first, says how much disk it frees, unloads the model if it is running, and you can download it again later.
Get from Hugging Face takes a link or org/name , lists the GGUF versions in the repository with their sizes and a plain label (for example "A good balance"), marks which ones fit your computer, and hands the download to Lemonade.
The Models page is for downloading models and editing how they run. My picks leads with the models I've tested and recommend; On this computer lists what you have. The meter on each card shows what will fit in the memory that is free on your graphics cards right now; if another program is using it, Zest warns you and names it.
Start here: LFM2.5 8B A1B. Very fast, runs on almost anything, calls tools well, and keeps answers short.
My general pick: Qwen3.6 35

[truncated]
