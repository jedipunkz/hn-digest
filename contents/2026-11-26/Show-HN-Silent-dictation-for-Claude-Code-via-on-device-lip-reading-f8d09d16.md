---
source: "https://github.com/batanus/lip-dictation"
hn_url: "https://news.ycombinator.com/item?id=50655639"
title: "Show HN: Silent dictation for Claude Code via on-device lip reading"
article_title: "GitHub - batanus/lip-dictation: Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt · GitHub"
image: "https://opengraph.githubassets.com/28ceb63ed591f32738ff15d3988b9751b4ded835996fe8a7c2fae97de7223948/batanus/lip-dictation"
author: "dima7711"
captured_at: "2026-11-26T14:45:28Z"
capture_tool: "hn-digest"
hn_id: 50655639
score: 4
comments: 0
posted_at: "2026-11-26T14:25:08Z"
tags:
  - hacker-news
---

# Show HN: Silent dictation for Claude Code via on-device lip reading

- HN: [50655639](https://news.ycombinator.com/item?id=50655639)
- Source: [github.com](https://github.com/batanus/lip-dictation)
- Score: 4
- Comments: 0
- Posted: 2026-11-26T14:25:08Z

## Translation

Title: Show HN: Silent dictation for Claude Code via on-device lip reading
Article title: GitHub - batanus/lip-dictation: Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt · GitHub
Description: Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt - batanus/lip-dictation

Article text:
GitHub - batanus/lip-dictation: Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
batanus
lip-dictation Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
.claude-plugin .claude-plugin .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE docs docs hooks hooks scripts scripts sidecar sidecar tests tests types types .gitignore .gitignore GUIDE.ru.md GUIDE.ru.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Talk to Claude Code without making a sound. Press a button above the prompt, mouth a phrase at your webcam, press again, and the words appear in the prompt box. A lip-reading model ( Auto-AVSR ) runs locally on your Mac: no audio is recorded and no video leaves your machine.
This is a prototype, and accuracy is modest. Lip reading is hard, and silent mouthing is harder than speech: expect word errors, especially on technical terms, so read the text before you send it. Claude is told which words were lip-read (with the model's other readings), so it reads them in context and asks before acting on anything destructive.
Recognition uses Auto-AVSR's visual-only model vsr_trlrs2lrs3vox2avsp_base : 250M parameters, trained on 3,291 hours of video (LRS2, LRS3, VoxCeleb2, AVSpeech).
¹ Auto-AVSR model zoo : 20.3% word error rate on LRS3 (audio models: 1–2%).
Why it is lower here than on the benchmark:
Silent mouthing moves the lips less than speaking aloud, and the model learned from people speaking aloud.
A laptop webcam at an angle, in office light, is harder than a TV camera on a well-lit face.
Vocabulary: TED talks rarely say "refactor" or "pull request".
Speed over accuracy: the plugin decodes with a beam of 10 and no external language model, to show the text about half a second after Stop.
What helps: short phrases, facing the camera in even light, mouthing a little more clearly than you would speak, and training it on your lips . Training checks how many words it gets wrong on held-out phrases, before and after, and keeps the personal model only when it gets fewer wrong, and ⚙ Settings → 🕘 History shows what the model read and what else it considered, so you can see your own accuracy.
Voice dictation needs you to speak out loud. Typing needs your hands and attention. Silent dictation fits the moments in between:
Open-plan office, coworking, café, library: ask Claude to "run the tests" or "show me the diff" without talking over your neighbours.
On a call or in a meeting: you are muted or listening, and can still send Claude a request.
At home when others are asleep: late-night work without waking anyone.
Noisy places: a train, a busy office. Lips don't care about background noise.
Privacy: nobody nearby hears what you ask, and recognition never leaves your Mac.
Resting your hands or voice: short commands without typing, or when speaking aloud is hard or tiring.
Install the plugin. In a terminal session of Claude Code:
/plugin install lip-dictation --marketplace batanus/lip-dictation
In the desktop app: click + next to the prompt box → Plugins → Add marketplace batanus/lip-dictation → Add plugin .
Press 📹 Set up lip dictation above the prompt (or run /lip-setup ). It shows the download size and the model's license terms, and starts only when you press Install. Nothing needs to be installed first: setup fetches what is missing (Python, its packages, the model) into ~/.cache/lip-dictation , about 1.5 GB, with a progress bar above the prompt. If your Mac lacks Apple's Command Line Tools, their installer opens; finish it and press Set up again.
Press 📹 Lip dictation , allow camera access when macOS asks (once), mouth a phrase, press ■ Stop .
Requirements: macOS on Apple Silicon, a webcam, about 2.5 GB of disk, and Claude Code 2.1.287 or newer (terminal or desktop app).
Press 📹 Lip dictation above the prompt. The camera turns on (the camera light shows it).
When the row shows ■ Stop ● 0:03 ▁▁▃▆█▅▂▁▁ , mouth one phrase clearly, facing the camera in good light. The timer counts up and the waveform follows your lips: bars rise when your mouth moves and settle to a flat line when it is still, and a warning if your face leaves the frame. Short phrases work best; recording stops by itself after 20 seconds, with a countdown for the last 5.
Press ■ Stop . A "Reading lips" progress bar runs for a few seconds, then the text appears at the cursor. Nothing is sent until you press Enter.
The first dictation of a session takes about 2 seconds to turn the camera on, while Python and the model load. After it, the next dictation is loaded ahead and waits (about 1 GB of memory, for up to 10 minutes), so the camera starts as soon as you press. The model's encoder runs on the Mac's GPU, so the text appears about half a second after you press Stop.
In a terminal session, focus the button with ctrl+x then tab and press l , or bind a key to the /lip command in ~/.claude/keybindings.json :
{ "bindings" : [{ "context" : " Chat " , "bindings" : { "meta+l" : " command:lip " } }] }
meta+l needs "Use Option as Meta key" turned on in Terminal.app or iTerm2; otherwise pick a combination that is free in your terminal. Typing /lip also works, but it replaces whatever is in the prompt box, so use the button or a key binding to add to a draft.
⚙ Settings , at the right of the row, holds everything else:
🧠 Personal model: On / Off — whether dictation uses the model trained on your lips (shown once you have trained one).
🎯 Train on my lips — record phrases and train your personal model (below).
🕘 History — what was read, what else the model considered, and what you sent.
🗑 Delete my data — training recordings, personal model and history.
The installed version, then ✕ Close .
Why: the stock model learned from people speaking aloud on TV and in talks. Silent mouthing moves the lips less, and everyone's lips move differently, so out of the box it makes more mistakes on you than on its benchmarks. Ten to fifteen minutes of your own phrases adapt it to you.
How: press ⚙ Settings → 🎯 Train on my lips (or run /lip-train ):
The row shows a phrase to mouth, such as «run the tests», with the usual ■ Stop and waveform. Mouth it and press Stop. Only the mouth area (96×96 px) of each phrase is kept, in ~/.cache/lip-dictation/personal/ .
Press ● Next phrase for the next one, as many as you like: Train appears after 20, and about 100 (10 to 15 minutes) works best. Done stops; the next /lip-train continues where you left off.
Press Train . A progress bar runs for a few minutes (about 1 second per phrase per pass, 8 passes). It adds small LoRA adapters (about 2M parameters) to the model's encoder, the part that reads the lips, and leaves the decoder, its language model, as it was, holds one phrase in five out, and counts the words it gets wrong on those, before and after. A toast reports both; the personal model replaces the stock one only when it makes fewer errors.
The phrases mix everyday English with short requests to a coding assistant, so the adapter learns your lips rather than one topic.
Once trained, ⚙ Settings shows 🧠 Personal model: On . Press it to switch to the stock model ( Off ) and back, for example to compare the two in the history; training a new model switches it on.
Each dictation also sends Claude the recognizer's other readings of the phrase (with the note below), since the right words are often among them.
⚙ Settings → 🕘 History opens a pane with the last dictations, newest first: the text that was inserted, the other readings the model considered, the clip length, the time it took to read, whether your personal model was used, and whether you sent the text as read or edited it first (with the prompt you sent). It is the quickest way to see what the model gets wrong for you. The same data is in ~/.cache/lip-dictation/dictations.jsonl , one JSON line per dictation or sent prompt, also printed by scripts/lipread.sh --history .
⚙ Settings → 🗑 Delete my data (or scripts/lipread.sh forget ) deletes the training phrases, the personal model and the history.
The button is only shown on this computer's terminal and desktop app, never on remote surfaces such as the phone app, and /lip only starts the camera when typed here. Settings in /config : hide the button (the status line then shows the state) and turn the note to Claude off.
button / /lip ──► mod (hooks/register.tsx) ──spawn──► scripts/lipread.sh ──► Lip Dictation.app ──► python -m lipread dictate
│ press: go file (standby) or a new process ──────────────► camera, resampled to 25 fps
│ 2nd press writes a stop file (model loads meanwhile)
◄── JSON lines: started / recording / level ×10/s / processing / result (+ other readings)
(level = lip opening from MediaPipe Face Mesh, drawn as the waveform)
result: $.prompt.fill inserts at the cursor; on Enter a prompt.submit hook adds a note for Claude
after it: the next dictation loads ahead (`dictate --wait`) and waits for the go file
sidecar/lipread/model.py : mouth crops (MediaPipe) → Auto-AVSR encoder (on the GPU) → beam search, plus the personal adapter when it is on.
sidecar/lipread/personal.py : the training phrases, LoRA training of the encoder and the before/after accuracy check.
scripts/camera-app/ : the helper app macOS grants camera access to (see the comment in main.c ).
Each dictation is a line in ~/.cache/lip-dictation/dictations.jsonl , the History pane's source.
Camera permission is granted to a small "Lip Dictation" helper app that setup builds from scripts/camera-app into ~/.cache/lip-dictation , not to your terminal or Claude. The sidecar runs through it because macOS refuses the camera without asking when the app running Claude Code declares no camera use, as under Claude Desktop.
Dictation frames are kept in memory for the length of one recording and discarded; nothing is written to disk unless you pass --save to the command-line tool. Training phrases ( /lip-train ) are the exception: their mouth crops stay in ~/.cache/lip-dictation/personal/ . The history keeps the text of each dictation and of prompts that contained one in ~/.cache/lip-dictation/dictations.jsonl . ⚙ Settings → 🗑 Delete my data removes all of it.
The only network access is the one-time setup: GitHub (uv, when it is not installed, and the Auto-AVSR code), PyPI (Python packages) and Google Drive (model weights).
Ideas, bug reports and pull requests are all we

[truncated]

## Original Extract

Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt - batanus/lip-dictation

GitHub - batanus/lip-dictation: Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
batanus
lip-dictation Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Silent dictation for Claude Code: mouth words at your webcam, a local lip-reading model types them into the prompt
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
.claude-plugin .claude-plugin .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE docs docs hooks hooks scripts scripts sidecar sidecar tests tests types types .gitignore .gitignore GUIDE.ru.md GUIDE.ru.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Talk to Claude Code without making a sound. Press a button above the prompt, mouth a phrase at your webcam, press again, and the words appear in the prompt box. A lip-reading model ( Auto-AVSR ) runs locally on your Mac: no audio is recorded and no video leaves your machine.
This is a prototype, and accuracy is modest. Lip reading is hard, and silent mouthing is harder than speech: expect word errors, especially on technical terms, so read the text before you send it. Claude is told which words were lip-read (with the model's other readings), so it reads them in context and asks before acting on anything destructive.
Recognition uses Auto-AVSR's visual-only model vsr_trlrs2lrs3vox2avsp_base : 250M parameters, trained on 3,291 hours of video (LRS2, LRS3, VoxCeleb2, AVSpeech).
¹ Auto-AVSR model zoo : 20.3% word error rate on LRS3 (audio models: 1–2%).
Why it is lower here than on the benchmark:
Silent mouthing moves the lips less than speaking aloud, and the model learned from people speaking aloud.
A laptop webcam at an angle, in office light, is harder than a TV camera on a well-lit face.
Vocabulary: TED talks rarely say "refactor" or "pull request".
Speed over accuracy: the plugin decodes with a beam of 10 and no external language model, to show the text about half a second after Stop.
What helps: short phrases, facing the camera in even light, mouthing a little more clearly than you would speak, and training it on your lips . Training checks how many words it gets wrong on held-out phrases, before and after, and keeps the personal model only when it gets fewer wrong, and ⚙ Settings → 🕘 History shows what the model read and what else it considered, so you can see your own accuracy.
Voice dictation needs you to speak out loud. Typing needs your hands and attention. Silent dictation fits the moments in between:
Open-plan office, coworking, café, library: ask Claude to "run the tests" or "show me the diff" without talking over your neighbours.
On a call or in a meeting: you are muted or listening, and can still send Claude a request.
At home when others are asleep: late-night work without waking anyone.
Noisy places: a train, a busy office. Lips don't care about background noise.
Privacy: nobody nearby hears what you ask, and recognition never leaves your Mac.
Resting your hands or voice: short commands without typing, or when speaking aloud is hard or tiring.
Install the plugin. In a terminal session of Claude Code:
/plugin install lip-dictation --marketplace batanus/lip-dictation
In the desktop app: click + next to the prompt box → Plugins → Add marketplace batanus/lip-dictation → Add plugin .
Press 📹 Set up lip dictation above the prompt (or run /lip-setup ). It shows the download size and the model's license terms, and starts only when you press Install. Nothing needs to be installed first: setup fetches what is missing (Python, its packages, the model) into ~/.cache/lip-dictation , about 1.5 GB, with a progress bar above the prompt. If your Mac lacks Apple's Command Line Tools, their installer opens; finish it and press Set up again.
Press 📹 Lip dictation , allow camera access when macOS asks (once), mouth a phrase, press ■ Stop .
Requirements: macOS on Apple Silicon, a webcam, about 2.5 GB of disk, and Claude Code 2.1.287 or newer (terminal or desktop app).
Press 📹 Lip dictation above the prompt. The camera turns on (the camera light shows it).
When the row shows ■ Stop ● 0:03 ▁▁▃▆█▅▂▁▁ , mouth one phrase clearly, facing the camera in good light. The timer counts up and the waveform follows your lips: bars rise when your mouth moves and settle to a flat line when it is still, and a warning if your face leaves the frame. Short phrases work best; recording stops by itself after 20 seconds, with a countdown for the last 5.
Press ■ Stop . A "Reading lips" progress bar runs for a few seconds, then the text appears at the cursor. Nothing is sent until you press Enter.
The first dictation of a session takes about 2 seconds to turn the camera on, while Python and the model load. After it, the next dictation is loaded ahead and waits (about 1 GB of memory, for up to 10 minutes), so the camera starts as soon as you press. The model's encoder runs on the Mac's GPU, so the text appears about half a second after you press Stop.
In a terminal session, focus the button with ctrl+x then tab and press l , or bind a key to the /lip command in ~/.claude/keybindings.json :
{ "bindings" : [{ "context" : " Chat " , "bindings" : { "meta+l" : " command:lip " } }] }
meta+l needs "Use Option as Meta key" turned on in Terminal.app or iTerm2; otherwise pick a combination that is free in your terminal. Typing /lip also works, but it replaces whatever is in the prompt box, so use the button or a key binding to add to a draft.
⚙ Settings , at the right of the row, holds everything else:
🧠 Personal model: On / Off — whether dictation uses the model trained on your lips (shown once you have trained one).
🎯 Train on my lips — record phrases and train your personal model (below).
🕘 History — what was read, what else the model considered, and what you sent.
🗑 Delete my data — training recordings, personal model and history.
The installed version, then ✕ Close .
Why: the stock model learned from people speaking aloud on TV and in talks. Silent mouthing moves the lips less, and everyone's lips move differently, so out of the box it makes more mistakes on you than on its benchmarks. Ten to fifteen minutes of your own phrases adapt it to you.
How: press ⚙ Settings → 🎯 Train on my lips (or run /lip-train ):
The row shows a phrase to mouth, such as «run the tests», with the usual ■ Stop and waveform. Mouth it and press Stop. Only the mouth area (96×96 px) of each phrase is kept, in ~/.cache/lip-dictation/personal/ .
Press ● Next phrase for the next one, as many as you like: Train appears after 20, and about 100 (10 to 15 minutes) works best. Done stops; the next /lip-train continues where you left off.
Press Train . A progress bar runs for a few minutes (about 1 second per phrase per pass, 8 passes). It adds small LoRA adapters (about 2M parameters) to the model's encoder, the part that reads the lips, and leaves the decoder, its language model, as it was, holds one phrase in five out, and counts the words it gets wrong on those, before and after. A toast reports both; the personal model replaces the stock one only when it makes fewer errors.
The phrases mix everyday English with short requests to a coding assistant, so the adapter learns your lips rather than one topic.
Once trained, ⚙ Settings shows 🧠 Personal model: On . Press it to switch to the stock model ( Off ) and back, for example to compare the two in the history; training a new model switches it on.
Each dictation also sends Claude the recognizer's other readings of the phrase (with the note below), since the right words are often among them.
⚙ Settings → 🕘 History opens a pane with the last dictations, newest first: the text that was inserted, the other readings the model considered, the clip length, the time it took to read, whether your personal model was used, and whether you sent the text as read or edited it first (with the prompt you sent). It is the quickest way to see what the model gets wrong for you. The same data is in ~/.cache/lip-dictation/dictations.jsonl , one JSON line per dictation or sent prompt, also printed by scripts/lipread.sh --history .
⚙ Settings → 🗑 Delete my data (or scripts/lipread.sh forget ) deletes the training phrases, the personal model and the history.
The button is only shown on this computer's terminal and desktop app, never on remote surfaces such as the phone app, and /lip only starts the camera when typed here. Settings in /config : hide the button (the status line then shows the state) and turn the note to Claude off.
button / /lip ──► mod (hooks/register.tsx) ──spawn──► scripts/lipread.sh ──► Lip Dictation.app ──► python -m lipread dictate
│ press: go file (standby) or a new process ──────────────► camera, resampled to 25 fps
│ 2nd press writes a stop file (model loads meanwhile)
◄── JSON lines: started / recording / level ×10/s / processing / result (+ other readings)
(level = lip opening from MediaPipe Face Mesh, drawn as the waveform)
result: $.prompt.fill inserts at the cursor; on Enter a prompt.submit hook adds a note for Claude
after it: the next dictation loads ahead (`dictate --wait`) and waits for the go file
sidecar/lipread/model.py : mouth crops (MediaPipe) → Auto-AVSR encoder (on the GPU) → beam search, plus the personal adapter when it is on.
sidecar/lipread/personal.py : the training phrases, LoRA training of the encoder and the before/after accuracy check.
scripts/camera-app/ : the helper app macOS grants camera access to (see the comment in main.c ).
Each dictation is a line in ~/.cache/lip-dictation/dictations.jsonl , the History pane's source.
Camera permission is granted to a small "Lip Dictation" helper app that setup builds from scripts/camera-app into ~/.cache/lip-dictation , not to your terminal or Claude. The sidecar runs through it because macOS refuses the camera without asking when the app running Claude Code declares no camera use, as under Claude Desktop.
Dictation frames are kept in memory for the length of one recording and discarded; nothing is written to disk unless you pass --save to the command-line tool. Training phrases ( /lip-train ) are the exception: their mouth crops stay in ~/.cache/lip-dictation/personal/ . The history keeps the text of each dictation and of prompts that contained one in ~/.cache/lip-dictation/dictations.jsonl . ⚙ Settings → 🗑 Delete my data removes all of it.
The only network access is the one-time setup: GitHub (uv, when it is not installed, and the Auto-AVSR code), PyPI (Python packages) and Google Drive (model weights).
Ideas, bug reports and pull requests are all we

[truncated]
