---
source: "https://github.com/OddlyGoods/peep-note"
hn_url: "https://news.ycombinator.com/item?id=50448161"
title: "Show HN: PeepNote – Local desktop notes for saving and reusing LLM responses"
article_title: "GitHub - OddlyGoods/peep-note: Save and recall useful LLM responses, Markdown notes, and screenshots. · GitHub"
image: "https://opengraph.githubassets.com/e84eee1fc40cc2899b63acb9b443c9b6825c77ba684874768fd1c68df2065c91/OddlyGoods/peep-note"
author: "oddlyHoon"
captured_at: "2026-11-11T01:03:09Z"
capture_tool: "hn-digest"
hn_id: 50448161
score: 1
comments: 0
posted_at: "2026-11-11T00:14:18Z"
tags:
  - hacker-news
---

# Show HN: PeepNote – Local desktop notes for saving and reusing LLM responses

- HN: [50448161](https://news.ycombinator.com/item?id=50448161)
- Source: [github.com](https://github.com/OddlyGoods/peep-note)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T00:14:18Z

## Translation

Title: Show HN: PeepNote – Local desktop notes for saving and reusing LLM responses
Article title: GitHub - OddlyGoods/peep-note: Save and recall useful LLM responses, Markdown notes, and screenshots. · GitHub
Description: Save and recall useful LLM responses, Markdown notes, and screenshots. - GitHub - OddlyGoods/peep-note: Save and recall useful LLM responses, Markdown notes, and screenshots.

Article text:
GitHub - OddlyGoods/peep-note: Save and recall useful LLM responses, Markdown notes, and screenshots. · GitHub
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
OddlyGoods
peep-note Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Save and recall useful LLM responses, Markdown notes, and screenshots.
peep-note-site.web.app Resources
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
263 Commits 263 Commits Folders and files
.github/ workflows .github/ workflows docs docs electron electron public public resources resources scripts scripts server server shared shared src src tests tests .env.example .env.example .gitignore .gitignore LICENSE LICENSE PRIVACY-AUDIT.md PRIVACY-AUDIT.md README-LINUX.md README-LINUX.md README.md README.md SECURITY.md SECURITY.md capture-analysis.html capture-analysis.html index.html index.html overlay.html overlay.html package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js quick-switch.html quick-switch.html tailwind.config.js tailwind.config.js tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
PeepNote is a local-first desktop app for saving useful LLM responses, recording your own thoughts, and finding them again while you work. Keep Markdown answers and screenshots together, recall them with a shortcut, and organize related records into AI summaries.
PeepNote keeps its application data on the local computer. AI summaries are optional and require an API key configured by the user.
2026-11-09 — v1.3.0 — Added Map, recently viewed Quick List ordering, and Alt+X to toggle global shortcuts.
2026-11-05 — v1.2.0 — Added Magic Pen and data folder switching.
2026-11-03 — v1.1.8 — Added remote connection guides, configurable context length, and automatic batched AI summaries.
2026-11-03 — v1.1.7 — Fixed Linux update errors and added HTTP API endpoint support.
2026-10-30 — v1.1.5 — Added MCP support, save notifications, and AI connection testing.
1. Save LLM responses as Markdown notes
Keep the answers worth revisiting without searching through a long conversation. Copy an LLM response, press Alt+1 , choose Markdown , and paste it into a note. Save with Ctrl+Enter , then open the note to review its headings, tables, and code blocks.
2. Capture diagrams and screenshots
Save the visual parts of an answer alongside your notes. Press Alt+2 , drag to select a screen region, and capture a diagram, code example, or reference image. Open the saved screenshot to inspect the details.
3. Recall notes with Quick List
Bring saved answers back into your workflow without searching through the main window. Press Alt+Q to open Quick List and move to the next item; use Alt+E to go back. Open an item when you need its full content. Quick List uses the currently selected date range.
4. Turn related notes into an AI summary
Bring answers from multiple LLM conversations together. Use Manage to select related records, create a group, and click AI Summary to organize the material into one structured note. Summaries use the provider and API key configured in Settings → AI settings . Response language defaults to your PeepNote language; choose a specific language to use it for both summaries and Magic Pen replies. AI settings, Magic Pen and MCP each have their own Save button.
5. Keep note previews private with Secret
Enable Secret when creating a note to hide its content in list previews. Opening the saved note prompts for your four-digit PIN before showing its contents. Set your PIN in Settings → Secret note lock .
6. Explore screen content with Magic Pen
Press Alt+S and draw around text, a diagram, or anything you want to understand on your screen. Magic Pen explains the selected content in a movable response window. Ask follow-up questions by typing or selecting another area, then save the conversation as a Markdown note in your chosen workspace. Choose your AI API or an installed, signed-in AI tool in Settings → AI settings → Magic Pen .
PeepNote opens directly to the main screen. There are no app accounts, login forms, Google sign-in, session tokens or authentication server. All windows use one local data space for notes, workspaces, settings and optional AI keys. OS user accounts still have their own application data directories.
The optional VITE_LOCAL_SERVER_PORT setting in .env.example defaults to 3848 , separate from the original PeepNote app. It is read at build time. The API listens only on 127.0.0.1 and rejects unrelated browser origins and hosts. No API key is bundled: enter your own provider key in Settings only if you use AI summaries.
The package and Linux executable are named peep-note . The app keeps its existing open-peep-note data directory and internal application ID so installed copies continue to read the same notes and settings after updates.
Create text, Markdown, code, and HTML notes
Capture a selected screen region and record a selected screen region
Draw directly over the live screen with Alt+S to ask AI about a selected area. In Settings → AI settings → Magic Pen , choose AI API to use the provider and API key from AI settings, or AI tools to choose Codex, Claude Code, Gemini CLI, or Qwen Code (the selected CLI must be installed and signed in). Magic Pen remembers its own model and custom input for each provider and tool, separately from summaries. API models must support image input. A movable response window shows the provider, model and duration, supports follow-up questions and captures, and saves the conversation with its images as a Markdown note. The window closes after a successful save. Codex analysis does not save a Codex session; change or disable the shortcut in Settings → Global shortcuts .
Select and remember a preferred monitor for capture and recording
Organize records with workspaces, tags, date ranges, search, and groups
Restore deleted records from Trash during a configurable retention period (7, 10, 30, or 90 days; default 10), or permanently delete them. Individually deleted AI summaries are removed immediately; deleting a whole group keeps its summary in Trash with the group.
Search Korean and English text in screenshots and attached images using local OCR
Generate AI summaries with OpenAI, Claude, Gemini, or Qwen
Connect Codex and other MCP clients to search, save, edit and group non-secret notes
Use global shortcuts and a recent-record Quick List
Protect secret notes with a four-digit PIN and masked list previews
English, Korean, Japanese, and Simplified Chinese UI
Windows 10/11 or a current Linux distribution
Build installers on the target operating system when possible: Windows for .exe , Linux for .deb
For Ubuntu 22.04 or newer, install Electron runtime dependencies:
sudo apt update
sudo apt install -y build-essential libnss3 libatk1.0-0 libatk-bridge2.0-0 \
libcups2 libdrm2 libgtk-3-0 libgbm1 libasound2t64 libxss1 \
xdg-utils xdg-desktop-portal xdg-desktop-portal-gtk
On Ubuntu 22.04, the audio package can be named libasound2 instead of libasound2t64 .
Wayland screen capture relies on PipeWire and an xdg-desktop-portal implementation for your desktop environment. Some Wayland compositors restrict global shortcuts; an Xorg session is the fallback if a shortcut cannot be registered.
git clone https://github.com/OddlyGoods/peep-note.git
cd peep-note
npm ci --onnxruntime-node-install-cuda=skip
npm run dev
Stop the development process with Ctrl+C . If the app remains in the system tray, use Quit from the tray menu to terminate the Electron process completely.
npm run build:linux
The package is written to release/ :
sudo apt install ./release/PeepNote- * -linux-amd64.deb
To produce an unpacked build for a quick smoke test:
npm run build:linux:dir
./release/linux-unpacked/peep-note
To update an existing PeepNote installation manually, install a newer .deb with the same command. Existing local records and settings are preserved.
Build the Windows installer on Windows:
npm run build:win
The NSIS installer is written to release/ . Run the generated PeepNote-*-Setup.exe and choose the normal install flow. Running the installer again updates an existing installation while preserving local records and settings.
Windows supports the same core features as Linux, including notes, local storage, comments, themes, AI summaries, capture, recording, tray controls, and global shortcuts. Capture behavior, multi-monitor DPI scaling, and shortcut availability should still be smoke-tested on the target Windows setup.
Command
Purpose
npm run dev
Run the development Electron app with Vite.
npm run build
Build using electron-builder's default target.
npm run build:linux
Build the Linux x64 Debian package.
npm run build:linux:dir
Build an unpacked Linux x64 app in release/linux-unpacked/ .
npm run build:win
Build the Windows NSIS installer.
npm run release
Create a release build and publish through the configured electron-builder provider.
npm run test:updates
Test update state, retries, renderer recovery, installation and release validation without installing anything.
npm run verify:release
Check the package/lockfile version and optional RELEASE_TAG .
App updates
Installed Windows and Ubuntu .deb builds check for updates automatically. To check manually, go to Settings → App information → App updates → Check for updates .
Choose Download , then Restart and install when you are ready. Ubuntu may ask for administrator authentication. Notes, attachments, and settings are preserved. Development and unpacked builds do not support in-app updates.
In Settings → General , Automatically advance dates updates the main list's end date to today when you reopen or refocus the app. By default, the start date moves with it to preserve the selected range length. Enable the separate Fix start date option to keep the selected start date and advance only the end date. Both options are saved independently; with automatic advancement off, the selected dates remain unchanged. You can still change dates manually. The separate Search screen uses its own date range.
Trash is in the left navigation, below Favorites . Deleting a record, selected records, or a group moves its records to Trash and removes them from normal lists, search, Quick List, and calendar counts. Deleted groups appear as a single group, including their AI summary, and can be restored or permanently deleted together. Select records to restore them or permanently delete them. Empty permanently deletes all of the current user's trashed records after confirmation.
Open an image record and click

[truncated]

## Original Extract

Save and recall useful LLM responses, Markdown notes, and screenshots. - GitHub - OddlyGoods/peep-note: Save and recall useful LLM responses, Markdown notes, and screenshots.

GitHub - OddlyGoods/peep-note: Save and recall useful LLM responses, Markdown notes, and screenshots. · GitHub
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
OddlyGoods
peep-note Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Save and recall useful LLM responses, Markdown notes, and screenshots.
peep-note-site.web.app Resources
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
263 Commits 263 Commits Folders and files
.github/ workflows .github/ workflows docs docs electron electron public public resources resources scripts scripts server server shared shared src src tests tests .env.example .env.example .gitignore .gitignore LICENSE LICENSE PRIVACY-AUDIT.md PRIVACY-AUDIT.md README-LINUX.md README-LINUX.md README.md README.md SECURITY.md SECURITY.md capture-analysis.html capture-analysis.html index.html index.html overlay.html overlay.html package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js quick-switch.html quick-switch.html tailwind.config.js tailwind.config.js tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
PeepNote is a local-first desktop app for saving useful LLM responses, recording your own thoughts, and finding them again while you work. Keep Markdown answers and screenshots together, recall them with a shortcut, and organize related records into AI summaries.
PeepNote keeps its application data on the local computer. AI summaries are optional and require an API key configured by the user.
2026-11-09 — v1.3.0 — Added Map, recently viewed Quick List ordering, and Alt+X to toggle global shortcuts.
2026-11-05 — v1.2.0 — Added Magic Pen and data folder switching.
2026-11-03 — v1.1.8 — Added remote connection guides, configurable context length, and automatic batched AI summaries.
2026-11-03 — v1.1.7 — Fixed Linux update errors and added HTTP API endpoint support.
2026-10-30 — v1.1.5 — Added MCP support, save notifications, and AI connection testing.
1. Save LLM responses as Markdown notes
Keep the answers worth revisiting without searching through a long conversation. Copy an LLM response, press Alt+1 , choose Markdown , and paste it into a note. Save with Ctrl+Enter , then open the note to review its headings, tables, and code blocks.
2. Capture diagrams and screenshots
Save the visual parts of an answer alongside your notes. Press Alt+2 , drag to select a screen region, and capture a diagram, code example, or reference image. Open the saved screenshot to inspect the details.
3. Recall notes with Quick List
Bring saved answers back into your workflow without searching through the main window. Press Alt+Q to open Quick List and move to the next item; use Alt+E to go back. Open an item when you need its full content. Quick List uses the currently selected date range.
4. Turn related notes into an AI summary
Bring answers from multiple LLM conversations together. Use Manage to select related records, create a group, and click AI Summary to organize the material into one structured note. Summaries use the provider and API key configured in Settings → AI settings . Response language defaults to your PeepNote language; choose a specific language to use it for both summaries and Magic Pen replies. AI settings, Magic Pen and MCP each have their own Save button.
5. Keep note previews private with Secret
Enable Secret when creating a note to hide its content in list previews. Opening the saved note prompts for your four-digit PIN before showing its contents. Set your PIN in Settings → Secret note lock .
6. Explore screen content with Magic Pen
Press Alt+S and draw around text, a diagram, or anything you want to understand on your screen. Magic Pen explains the selected content in a movable response window. Ask follow-up questions by typing or selecting another area, then save the conversation as a Markdown note in your chosen workspace. Choose your AI API or an installed, signed-in AI tool in Settings → AI settings → Magic Pen .
PeepNote opens directly to the main screen. There are no app accounts, login forms, Google sign-in, session tokens or authentication server. All windows use one local data space for notes, workspaces, settings and optional AI keys. OS user accounts still have their own application data directories.
The optional VITE_LOCAL_SERVER_PORT setting in .env.example defaults to 3848 , separate from the original PeepNote app. It is read at build time. The API listens only on 127.0.0.1 and rejects unrelated browser origins and hosts. No API key is bundled: enter your own provider key in Settings only if you use AI summaries.
The package and Linux executable are named peep-note . The app keeps its existing open-peep-note data directory and internal application ID so installed copies continue to read the same notes and settings after updates.
Create text, Markdown, code, and HTML notes
Capture a selected screen region and record a selected screen region
Draw directly over the live screen with Alt+S to ask AI about a selected area. In Settings → AI settings → Magic Pen , choose AI API to use the provider and API key from AI settings, or AI tools to choose Codex, Claude Code, Gemini CLI, or Qwen Code (the selected CLI must be installed and signed in). Magic Pen remembers its own model and custom input for each provider and tool, separately from summaries. API models must support image input. A movable response window shows the provider, model and duration, supports follow-up questions and captures, and saves the conversation with its images as a Markdown note. The window closes after a successful save. Codex analysis does not save a Codex session; change or disable the shortcut in Settings → Global shortcuts .
Select and remember a preferred monitor for capture and recording
Organize records with workspaces, tags, date ranges, search, and groups
Restore deleted records from Trash during a configurable retention period (7, 10, 30, or 90 days; default 10), or permanently delete them. Individually deleted AI summaries are removed immediately; deleting a whole group keeps its summary in Trash with the group.
Search Korean and English text in screenshots and attached images using local OCR
Generate AI summaries with OpenAI, Claude, Gemini, or Qwen
Connect Codex and other MCP clients to search, save, edit and group non-secret notes
Use global shortcuts and a recent-record Quick List
Protect secret notes with a four-digit PIN and masked list previews
English, Korean, Japanese, and Simplified Chinese UI
Windows 10/11 or a current Linux distribution
Build installers on the target operating system when possible: Windows for .exe , Linux for .deb
For Ubuntu 22.04 or newer, install Electron runtime dependencies:
sudo apt update
sudo apt install -y build-essential libnss3 libatk1.0-0 libatk-bridge2.0-0 \
libcups2 libdrm2 libgtk-3-0 libgbm1 libasound2t64 libxss1 \
xdg-utils xdg-desktop-portal xdg-desktop-portal-gtk
On Ubuntu 22.04, the audio package can be named libasound2 instead of libasound2t64 .
Wayland screen capture relies on PipeWire and an xdg-desktop-portal implementation for your desktop environment. Some Wayland compositors restrict global shortcuts; an Xorg session is the fallback if a shortcut cannot be registered.
git clone https://github.com/OddlyGoods/peep-note.git
cd peep-note
npm ci --onnxruntime-node-install-cuda=skip
npm run dev
Stop the development process with Ctrl+C . If the app remains in the system tray, use Quit from the tray menu to terminate the Electron process completely.
npm run build:linux
The package is written to release/ :
sudo apt install ./release/PeepNote- * -linux-amd64.deb
To produce an unpacked build for a quick smoke test:
npm run build:linux:dir
./release/linux-unpacked/peep-note
To update an existing PeepNote installation manually, install a newer .deb with the same command. Existing local records and settings are preserved.
Build the Windows installer on Windows:
npm run build:win
The NSIS installer is written to release/ . Run the generated PeepNote-*-Setup.exe and choose the normal install flow. Running the installer again updates an existing installation while preserving local records and settings.
Windows supports the same core features as Linux, including notes, local storage, comments, themes, AI summaries, capture, recording, tray controls, and global shortcuts. Capture behavior, multi-monitor DPI scaling, and shortcut availability should still be smoke-tested on the target Windows setup.
Command
Purpose
npm run dev
Run the development Electron app with Vite.
npm run build
Build using electron-builder's default target.
npm run build:linux
Build the Linux x64 Debian package.
npm run build:linux:dir
Build an unpacked Linux x64 app in release/linux-unpacked/ .
npm run build:win
Build the Windows NSIS installer.
npm run release
Create a release build and publish through the configured electron-builder provider.
npm run test:updates
Test update state, retries, renderer recovery, installation and release validation without installing anything.
npm run verify:release
Check the package/lockfile version and optional RELEASE_TAG .
App updates
Installed Windows and Ubuntu .deb builds check for updates automatically. To check manually, go to Settings → App information → App updates → Check for updates .
Choose Download , then Restart and install when you are ready. Ubuntu may ask for administrator authentication. Notes, attachments, and settings are preserved. Development and unpacked builds do not support in-app updates.
In Settings → General , Automatically advance dates updates the main list's end date to today when you reopen or refocus the app. By default, the start date moves with it to preserve the selected range length. Enable the separate Fix start date option to keep the selected start date and advance only the end date. Both options are saved independently; with automatic advancement off, the selected dates remain unchanged. You can still change dates manually. The separate Search screen uses its own date range.
Trash is in the left navigation, below Favorites . Deleting a record, selected records, or a group moves its records to Trash and removes them from normal lists, search, Quick List, and calendar counts. Deleted groups appear as a single group, including their AI summary, and can be restored or permanently deleted together. Select records to restore them or permanently delete them. Empty permanently deletes all of the current user's trashed records after confirmation.
Open an image record and click

[truncated]
