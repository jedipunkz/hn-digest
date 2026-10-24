---
source: "https://github.com/swrdfish/narro"
hn_url: "https://news.ycombinator.com/item?id=50224373"
title: "Show HN: Narro – Read web pages and Kindle books using local AI voices"
article_title: "GitHub - swrdfish/Narro: Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader. · GitHub"
image: "https://repository-images.githubusercontent.com/1431486431/7c061558-bc38-4131-b45c-5c2cbe5791ca"
author: "ghoshbinayak"
captured_at: "2026-10-24T20:20:03Z"
capture_tool: "hn-digest"
hn_id: 50224373
score: 5
comments: 1
posted_at: "2026-10-24T19:37:32Z"
tags:
  - hacker-news
---

# Show HN: Narro – Read web pages and Kindle books using local AI voices

- HN: [50224373](https://news.ycombinator.com/item?id=50224373)
- Source: [github.com](https://github.com/swrdfish/narro)
- Score: 5
- Comments: 1
- Posted: 2026-10-24T19:37:32Z

## Translation

Title: Show HN: Narro – Read web pages and Kindle books using local AI voices
Article title: GitHub - swrdfish/Narro: Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader. · GitHub
Description: Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader. - swrdfish/Narro
HN text: I wanted a read-aloud extension that could read from the Kindle Online Reader. The ones I tried used voices I did not enjoy listening to or had a UX that was too cluttered or annoying or needed a signup. So I decided to vibe code my own. It is open source and uses open weight models that run in the browser. https://swrdfish.github.io/Narro/

Article text:
GitHub - swrdfish/Narro: Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader. · GitHub
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
swrdfish
Narro Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader.
swrdfish.github.io/Narro/ Topics
Readme MIT license Contributing
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
64 Commits 64 Commits Folders and files
.github/ workflows .github/ workflows docs docs icons icons scripts scripts site site src src store-assets store-assets test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md build.mjs build.mjs manifest.json manifest.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Read any web page aloud with on device AI voices. Supports Kindle online reader.
Features · Privacy · Report an issue
Narro is a Chrome extension that reads the page you're on and highlights each sentence as it's
spoken. Use a cloud voice with your own API key, or a voice that runs entirely in your browser,
with no account and no server. It also supports Kindle Online Reader.
Engine
Notes
Cloud
OpenAI
gpt-4o-mini-tts with style instructions, tts-1 , tts-1-hd
ElevenLabs
Your own voices, listed with your key
Google Gemini
Gemini TTS voices
OpenRouter
One key for speech models from many providers
In your browser
Kokoro-82M
8 voices in English
Supertonic 2
10 voices in English, Spanish, French, Portuguese and Korean
Piper
Popular piper-voices in several languages
Kitten TTS Nano
8 voices in English (Jasper, Bella, Luna, Bruno, Rosie, Hugo, Kiki, Leo); the smallest download
Other
Local server
Any OpenAI-compatible /v1/audio/speech server, e.g. Kokoro-FastAPI
System voices
Your computer's built-in voices
What the in-browser voices need
Each model downloads once and then works offline. Memory is what Narro uses in total at its peak
while reading, as Chrome's Task Manager shows it.
Download sizes are the model files on Hugging Face. Memory was measured in Chromium on an M1 Pro,
reading a 12-paragraph article with the model already downloaded, and rounded to 50 MB; expect some
variation between machines. Most of it is the speech runtime the three engines share, which is why
a small Piper voice needs about as much as Kokoro. Kokoro and Supertonic use WebGPU when Device is
Auto and the browser supports it; the popup shows the figures for your setup.
Sentence highlighting , without changing the page. The page scrolls to follow along.
Natural speech: consecutive sentences are sent as one request, so the voice keeps its intonation. The first sentence is sent on its own, so speech starts quickly.
Play/pause, previous/next sentence , and speed from 0.5× to 3× with pitch preserved.
Start anywhere: Point and select capability. Select text and right-click → Read aloud from here .
Resumes where you left off on each page.
Control it from any tab: the popup shows what's being read, with pause and stop. The toolbar icon shows a red ▶︎ while it plays.
Kindle Cloud Reader: reads book pages with on-device OCR and turns pages automatically ( details ).
Install Narro from the Chrome Web Store , then click the Narro icon on any page,
choose a voice engine and press Read this page .
To try the latest changes before they reach the store:
git clone https://github.com/swrdfish/Narro.git
cd Narro
npm install
npm run build
Open chrome://extensions and turn on Developer mode .
Choose Load unpacked and select the dist/ folder.
Click the Narro icon, choose a voice engine, then Read this page .
Keyboard shortcuts are shown in the popup and can be changed at chrome://extensions/shortcuts .
Text you have read aloud : with a cloud voice, the text is sent to that provider (OpenAI, ElevenLabs or Google) under your account and their terms. With OpenRouter, it goes to OpenRouter and the model's provider. With local models like Kokoro, Supertonic, Piper, Kitten TTS or system voices, nothing leaves your computer. A local server receives the text at the address you configure.
API keys are stored in chrome.storage.local on your computer. They're never synced and never sent anywhere except to the provider they belong to.
Model downloads : in-browser voices download their model files from Hugging Face on first use, then cache them in your browser.
Reading positions : Narro stores a short fingerprint of the last sentence you heard on each page, plus its position, for up to 200 pages. These stay on your computer and don't include the text itself.
Narro has no analytics, accounts or servers of its own .
Kindle Cloud Reader shows book pages as images rather than text, so on read.amazon.com , .co.uk , .in , .de , etc. Narro runs
OCR (text recognition) on the page you can see , entirely on your computer, using Tesseract. It then reads the recognized text with your chosen voice.
At the end of a page, Narro presses Kindle's Next page button and keeps reading. It can turn the page slightly early, while the page's last sentence is still playing, so there's no pause. Both behaviours can be switched off in the settings.
A sentence split across a page turn or a column break is spoken once, in full.
Reading continues if the Kindle tab is hidden behind another window. If Kindle doesn't draw the next page while hidden, Narro waits until the tab is visible again.
OCR currently supports English only. Decorative fonts, very small text and illustrations reduce accuracy.
Narro is an independent project. It isn't affiliated with or endorsed by Amazon. Kindle is a
trademark of Amazon.com, Inc. Narro only reads pages you have open yourself, but automated use of
Kindle Cloud Reader may be restricted by Amazon's terms. Check that your use complies with them.
The Model dropdowns for OpenAI, ElevenLabs and Gemini, and the Piper voice list, come from
src/models.json . Each list maps an ID to the label shown in the dropdown:
{ "openai" : { "tts-1" : " tts-1 · fastest " }, "piper" : { "en_US-lessac-medium" : " Lessac · US English, female " } }
The local section holds the download size and memory of each in-browser model, shown in the popup:
per device ( wasm for CPU, webgpu ) for Kokoro and Supertonic, and per voice for Piper. A Piper
voice added to the list needs an entry there too.
{ "local" : { "piper" : { "en_US-lessac-medium" : { "downloadMB" : 63.2 , "memoryMB" : 800 } } } }
To make a permanent change, edit src/models.json and run npm run build .
To make a quick change without rebuilding, edit dist/models.json and reload the extension. The next build overwrites this copy.
A saved model that's missing from a list stays selectable. The local server's model list comes from
its own /v1/models endpoint.
npm run watch # rebuild the extension on changes
npm test # test suite (Node's test runner + jsdom)
npm run site:preview # preview the landing page (site/) at http://localhost:8000
After a rebuild, reload Narro at chrome://extensions , then reload the tab you're testing on. See
CONTRIBUTING.md for guidelines.
Found a bug or have an idea? Open an issue . You can also use the
Send feedback or report an issue link at the bottom of Narro's popup.
Narro is a Manifest V3 extension with no framework and no build-time code generation: plain ES
modules bundled by esbuild.
flowchart LR
UI["Popup and settings<br/>(options.html)"] -->|start, pause, stop| BG
KEYS["Shortcuts and<br/>context menu"] --> BG
BG["background.js<br/>service worker"] <-->|sentences, playback events| CS
subgraph TAB["The tab being read"]
CS["content.js<br/>text, highlighting, player"]
end
BG <-->|speech requests, audio events| OFF["offscreen.js<br/>synthesizes and plays audio"]
OFF --> LW["local-worker.js<br/>Kokoro, Supertonic, Piper, Kitten"]
OFF --> CLOUD[("Cloud speech APIs<br/>OpenAI, ElevenLabs, Gemini, OpenRouter")]
OFF --> OCR["Tesseract OCR<br/>Kindle pages"]
Loading
Design notes
Audio plays in an offscreen document, so a page's security policy can't block it, and cloud APIs are called from the extension's own origin.
Each speech request carries a sequence number, so the page ignores events from requests it has already replaced.
While reading, the page holds a connection to the background. Refreshing or closing the page stops playback.
Narro's source code is released under the MIT License .
The built extension also includes third-party components under their own licences, including
eSpeak NG (GPL-3.0-or-later) through the phonemizer used by Kokoro, Piper and Kitten TTS. See
THIRD_PARTY_NOTICES.md for details, including the licences of models that
are downloaded at runtime.
Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader.
swrdfish.github.io/Narro/ Topics
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader. - swrdfish/Narro

I wanted a read-aloud extension that could read from the Kindle Online Reader. The ones I tried used voices I did not enjoy listening to or had a UX that was too cluttered or annoying or needed a signup. So I decided to vibe code my own. It is open source and uses open weight models that run in the browser. https://swrdfish.github.io/Narro/

GitHub - swrdfish/Narro: Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader. · GitHub
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
swrdfish
Narro Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader.
swrdfish.github.io/Narro/ Topics
Readme MIT license Contributing
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
64 Commits 64 Commits Folders and files
.github/ workflows .github/ workflows docs docs icons icons scripts scripts site site src src store-assets store-assets test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md build.mjs build.mjs manifest.json manifest.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Read any web page aloud with on device AI voices. Supports Kindle online reader.
Features · Privacy · Report an issue
Narro is a Chrome extension that reads the page you're on and highlights each sentence as it's
spoken. Use a cloud voice with your own API key, or a voice that runs entirely in your browser,
with no account and no server. It also supports Kindle Online Reader.
Engine
Notes
Cloud
OpenAI
gpt-4o-mini-tts with style instructions, tts-1 , tts-1-hd
ElevenLabs
Your own voices, listed with your key
Google Gemini
Gemini TTS voices
OpenRouter
One key for speech models from many providers
In your browser
Kokoro-82M
8 voices in English
Supertonic 2
10 voices in English, Spanish, French, Portuguese and Korean
Piper
Popular piper-voices in several languages
Kitten TTS Nano
8 voices in English (Jasper, Bella, Luna, Bruno, Rosie, Hugo, Kiki, Leo); the smallest download
Other
Local server
Any OpenAI-compatible /v1/audio/speech server, e.g. Kokoro-FastAPI
System voices
Your computer's built-in voices
What the in-browser voices need
Each model downloads once and then works offline. Memory is what Narro uses in total at its peak
while reading, as Chrome's Task Manager shows it.
Download sizes are the model files on Hugging Face. Memory was measured in Chromium on an M1 Pro,
reading a 12-paragraph article with the model already downloaded, and rounded to 50 MB; expect some
variation between machines. Most of it is the speech runtime the three engines share, which is why
a small Piper voice needs about as much as Kokoro. Kokoro and Supertonic use WebGPU when Device is
Auto and the browser supports it; the popup shows the figures for your setup.
Sentence highlighting , without changing the page. The page scrolls to follow along.
Natural speech: consecutive sentences are sent as one request, so the voice keeps its intonation. The first sentence is sent on its own, so speech starts quickly.
Play/pause, previous/next sentence , and speed from 0.5× to 3× with pitch preserved.
Start anywhere: Point and select capability. Select text and right-click → Read aloud from here .
Resumes where you left off on each page.
Control it from any tab: the popup shows what's being read, with pause and stop. The toolbar icon shows a red ▶︎ while it plays.
Kindle Cloud Reader: reads book pages with on-device OCR and turns pages automatically ( details ).
Install Narro from the Chrome Web Store , then click the Narro icon on any page,
choose a voice engine and press Read this page .
To try the latest changes before they reach the store:
git clone https://github.com/swrdfish/Narro.git
cd Narro
npm install
npm run build
Open chrome://extensions and turn on Developer mode .
Choose Load unpacked and select the dist/ folder.
Click the Narro icon, choose a voice engine, then Read this page .
Keyboard shortcuts are shown in the popup and can be changed at chrome://extensions/shortcuts .
Text you have read aloud : with a cloud voice, the text is sent to that provider (OpenAI, ElevenLabs or Google) under your account and their terms. With OpenRouter, it goes to OpenRouter and the model's provider. With local models like Kokoro, Supertonic, Piper, Kitten TTS or system voices, nothing leaves your computer. A local server receives the text at the address you configure.
API keys are stored in chrome.storage.local on your computer. They're never synced and never sent anywhere except to the provider they belong to.
Model downloads : in-browser voices download their model files from Hugging Face on first use, then cache them in your browser.
Reading positions : Narro stores a short fingerprint of the last sentence you heard on each page, plus its position, for up to 200 pages. These stay on your computer and don't include the text itself.
Narro has no analytics, accounts or servers of its own .
Kindle Cloud Reader shows book pages as images rather than text, so on read.amazon.com , .co.uk , .in , .de , etc. Narro runs
OCR (text recognition) on the page you can see , entirely on your computer, using Tesseract. It then reads the recognized text with your chosen voice.
At the end of a page, Narro presses Kindle's Next page button and keeps reading. It can turn the page slightly early, while the page's last sentence is still playing, so there's no pause. Both behaviours can be switched off in the settings.
A sentence split across a page turn or a column break is spoken once, in full.
Reading continues if the Kindle tab is hidden behind another window. If Kindle doesn't draw the next page while hidden, Narro waits until the tab is visible again.
OCR currently supports English only. Decorative fonts, very small text and illustrations reduce accuracy.
Narro is an independent project. It isn't affiliated with or endorsed by Amazon. Kindle is a
trademark of Amazon.com, Inc. Narro only reads pages you have open yourself, but automated use of
Kindle Cloud Reader may be restricted by Amazon's terms. Check that your use complies with them.
The Model dropdowns for OpenAI, ElevenLabs and Gemini, and the Piper voice list, come from
src/models.json . Each list maps an ID to the label shown in the dropdown:
{ "openai" : { "tts-1" : " tts-1 · fastest " }, "piper" : { "en_US-lessac-medium" : " Lessac · US English, female " } }
The local section holds the download size and memory of each in-browser model, shown in the popup:
per device ( wasm for CPU, webgpu ) for Kokoro and Supertonic, and per voice for Piper. A Piper
voice added to the list needs an entry there too.
{ "local" : { "piper" : { "en_US-lessac-medium" : { "downloadMB" : 63.2 , "memoryMB" : 800 } } } }
To make a permanent change, edit src/models.json and run npm run build .
To make a quick change without rebuilding, edit dist/models.json and reload the extension. The next build overwrites this copy.
A saved model that's missing from a list stays selectable. The local server's model list comes from
its own /v1/models endpoint.
npm run watch # rebuild the extension on changes
npm test # test suite (Node's test runner + jsdom)
npm run site:preview # preview the landing page (site/) at http://localhost:8000
After a rebuild, reload Narro at chrome://extensions , then reload the tab you're testing on. See
CONTRIBUTING.md for guidelines.
Found a bug or have an idea? Open an issue . You can also use the
Send feedback or report an issue link at the bottom of Narro's popup.
Narro is a Manifest V3 extension with no framework and no build-time code generation: plain ES
modules bundled by esbuild.
flowchart LR
UI["Popup and settings<br/>(options.html)"] -->|start, pause, stop| BG
KEYS["Shortcuts and<br/>context menu"] --> BG
BG["background.js<br/>service worker"] <-->|sentences, playback events| CS
subgraph TAB["The tab being read"]
CS["content.js<br/>text, highlighting, player"]
end
BG <-->|speech requests, audio events| OFF["offscreen.js<br/>synthesizes and plays audio"]
OFF --> LW["local-worker.js<br/>Kokoro, Supertonic, Piper, Kitten"]
OFF --> CLOUD[("Cloud speech APIs<br/>OpenAI, ElevenLabs, Gemini, OpenRouter")]
OFF --> OCR["Tesseract OCR<br/>Kindle pages"]
Loading
Design notes
Audio plays in an offscreen document, so a page's security policy can't block it, and cloud APIs are called from the extension's own origin.
Each speech request carries a sequence number, so the page ignores events from requests it has already replaced.
While reading, the page holds a connection to the background. Refreshing or closing the page stops playback.
Narro's source code is released under the MIT License .
The built extension also includes third-party components under their own licences, including
eSpeak NG (GPL-3.0-or-later) through the phonemizer used by Kokoro, Piper and Kitten TTS. See
THIRD_PARTY_NOTICES.md for details, including the licences of models that
are downloaded at runtime.
Reads any web page aloud in natural AI voices, from cloud or local models. Supports Kindle Cloud Reader.
swrdfish.github.io/Narro/ Topics
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
