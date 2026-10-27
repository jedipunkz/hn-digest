---
source: "https://randcomp.io/mission-control-site/"
hn_url: "https://news.ycombinator.com/item?id=50255427"
title: "RandComp Studio – run a team of AI coding agents on your own machine"
article_title: "RandComp Studio — say one sentence, get a team of experts"
image: "https://randcomp.io/mission-control-site/screenshots/project-diagram.png"
author: "tyrelltle"
captured_at: "2026-10-27T07:23:10Z"
capture_tool: "hn-digest"
hn_id: 50255427
score: 1
comments: 0
posted_at: "2026-10-27T06:48:55Z"
tags:
  - hacker-news
---

# RandComp Studio – run a team of AI coding agents on your own machine

- HN: [50255427](https://news.ycombinator.com/item?id=50255427)
- Source: [randcomp.io](https://randcomp.io/mission-control-site/)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T06:48:55Z

## Translation

Title: RandComp Studio – run a team of AI coding agents on your own machine
Article title: RandComp Studio — say one sentence, get a team of experts
Description: RandComp Studio plans the work, builds it, and has an independent checker verify it — all on your own computer, using the AI tools you already have. Free early preview, no account.

Article text:
Early preview — expect a few bugs. Share feedback on GitHub →
RandComp Studio
Local only Three ways Documentation About 中文 Download for macOS
Hi there! I live on your computer — and so does everything your team makes.
Say one sentence.
Get a team of experts.
RandComp Studio plans the work, builds it, and has an independent checker verify it — all on your own computer, using the AI tools you already have: Claude Code, Codex, Kimi, or a local model.
Local only The app, your projects, your agents’ terminals and their files all live on your own computer. No account, no RandComp cloud, no server of ours in the middle.
Your chosen AI tools still talk to their own providers, and the app sends us anonymous usage counts you can switch off. Exactly what leaves your computer →
Free right now · Runs on your computer · No account
One team on your computer. Three ways to reach it.
The app, the randcomp command and MCP all drive the same local company: the same projects, the same files, never a second copy. Use whichever suits the moment.
Watch the plan grow as a diagram, open any task to see what it produced, and find everything that needs you in one place.
The plan, one click from the diagram. Click any task to see what it made.
Drive the same company from your terminal or your scripts — over SSH from another machine, too. Each command prints JSON.
randcomp project create "Build a reading-list site"
randcomp project list
randcomp project ask PROJECT_ID "Add a search box"
Install it from the app menu: RandComp Studio → Install the randcomp command in PATH . Every command →
Ask Claude Code, Claude Desktop, Codex or another MCP client how the team is doing, and hand it work — without leaving the chat you are in.
# Claude Code, once:
claude mcp add -s user randcomp-studio -- randcomp mcp
Your assistant gets tools such as project_list , task_create , helper_list and result_list ; the ones that start paid agents say so. Other clients add the same randcomp mcp in their MCP settings. Setup →
All three talk only to the RandComp process on your computer — no extra cloud in between.
📱 Connect from your phone. Check on your team and answer what needs you while you are away from your desk. The work itself stays on your computer.
☁️ Deploy it to your own cloud. Paid Run the same team in a cloud account you own, so it keeps working while your computer sleeps. Your cloud, not ours — and this one will have a price.
💻 Windows and Linux. After the one-person army has had some sleep.
No dates yet — they ship when they work. Tell us which one you want first: Feedback →
It all happens on your own computer.
RandComp Studio is a desktop app, not a service. There is no account to make and no RandComp server between you and your agents: the app, your projects, every agent’s terminal and every file it writes stay on your computer. It drives the AI command-line tools you already have — Claude Code, Codex, Kimi Code or Ollama — just as you would from a terminal.
RandComp Studio Claude Code Codex Kimi A model on this machine (Ollama)
📄 plan.md 📄 DESIGN.md 📁 your-project/ 📄 qa-report.md ⌨️ each agent’s terminal
What does leave your computer — all of it
Your prompts and code → the AI tool you picked. Claude Code talks to Anthropic, Codex to OpenAI, Kimi to Moonshot AI — on your own account, exactly as when you run them yourself. With a model on your computer (Ollama), nothing leaves.
Anonymous usage counts → RandComp. Four events: app opened (at most once a day), project created, MCP client connected, message sent. Each is only the event’s name, a UTC time and a random install ID — no content, no project names, no file paths. It is on by default; Settings → Usage statistics turns it off and deletes anything not yet sent.
Nothing else. RandComp account: none · RandComp copy of your work: none · Your prompts and files sent to RandComp: none. The app’s own server listens on 127.0.0.1 only.
Unplug and see. A role on a model on your computer (Ollama) keeps working without the internet. Roles on Claude Code, Codex or Kimi get no answers until you are back online, because their provider is on the internet. RandComp itself does not need it.
Honest limit: “local” means RandComp adds no cloud of its own — it does not make Claude Code, Codex or Kimi work offline. And small local models suit small, bounded work.
Not one AI grading its own homework.
The checker never saw the builder's reasoning, and may run on a different AI entirely. Mix freely — a strong model for planning, a cheaper one for routine work, a private one for sensitive files.
Planning on Claude Code, building on Codex, a quick check on Kimi, drafts on a model on your machine. Change it any time, even mid-project.
The checker works from the result, not from the builder's story, so blind spots are not shared. It is a second look, not a promise that everything is correct.
More agents cost more tokens. Small jobs stay a single session; teams are for work that earns it.
# everything is a plain file on your disk
reading-list/
plan.md coordinator
DESIGN.md designer
server/ web/ developers
qa-report.md independent QA
Open them in any editor. Move them, back them up, delete them. No export step, because there is nothing to export from.
Free during early preview. macOS only for now — Windows and Linux come later, once the one-person army has slept.
Download for macOS ↓ Read the docs
A Mac with Apple silicon (M1 or newer) on macOS 11 or later.
At least one AI command-line tool, installed and signed in — Claude Code, Codex or Kimi Code, on your own plan or API key with that company. Or Ollama with Qwen Code, for a free model that runs on your computer — fine for small, bounded work.
RandComp does not include AI usage. Each tool bills you exactly as it does today, and more helpers use more of it.
Apple silicon · macOS 11+ · 33 MB
SHA-256 88c7f816503e42b6a022be0ca4ca946ba948a03b955101bf140188c96c6503ab
It started while I was traveling. I wanted agents working on my Mac mini back home — but one long AI session goes out of shape fast: it drifts, loses the plot, and no one is there to catch it. So instead of one agent doing everything, I built a company: a coordinator plans the work, each role owns a piece, and I can attach to any of them from wherever I am. Then I kept using it myself — several small apps later, it had earned its place as my default way of working. That's when I decided it was time to share it. As for the name — I couldn't think of one, so Random Company it is. RandComp is just the shortcut.

## Original Extract

RandComp Studio plans the work, builds it, and has an independent checker verify it — all on your own computer, using the AI tools you already have. Free early preview, no account.

Early preview — expect a few bugs. Share feedback on GitHub →
RandComp Studio
Local only Three ways Documentation About 中文 Download for macOS
Hi there! I live on your computer — and so does everything your team makes.
Say one sentence.
Get a team of experts.
RandComp Studio plans the work, builds it, and has an independent checker verify it — all on your own computer, using the AI tools you already have: Claude Code, Codex, Kimi, or a local model.
Local only The app, your projects, your agents’ terminals and their files all live on your own computer. No account, no RandComp cloud, no server of ours in the middle.
Your chosen AI tools still talk to their own providers, and the app sends us anonymous usage counts you can switch off. Exactly what leaves your computer →
Free right now · Runs on your computer · No account
One team on your computer. Three ways to reach it.
The app, the randcomp command and MCP all drive the same local company: the same projects, the same files, never a second copy. Use whichever suits the moment.
Watch the plan grow as a diagram, open any task to see what it produced, and find everything that needs you in one place.
The plan, one click from the diagram. Click any task to see what it made.
Drive the same company from your terminal or your scripts — over SSH from another machine, too. Each command prints JSON.
randcomp project create "Build a reading-list site"
randcomp project list
randcomp project ask PROJECT_ID "Add a search box"
Install it from the app menu: RandComp Studio → Install the randcomp command in PATH . Every command →
Ask Claude Code, Claude Desktop, Codex or another MCP client how the team is doing, and hand it work — without leaving the chat you are in.
# Claude Code, once:
claude mcp add -s user randcomp-studio -- randcomp mcp
Your assistant gets tools such as project_list , task_create , helper_list and result_list ; the ones that start paid agents say so. Other clients add the same randcomp mcp in their MCP settings. Setup →
All three talk only to the RandComp process on your computer — no extra cloud in between.
📱 Connect from your phone. Check on your team and answer what needs you while you are away from your desk. The work itself stays on your computer.
☁️ Deploy it to your own cloud. Paid Run the same team in a cloud account you own, so it keeps working while your computer sleeps. Your cloud, not ours — and this one will have a price.
💻 Windows and Linux. After the one-person army has had some sleep.
No dates yet — they ship when they work. Tell us which one you want first: Feedback →
It all happens on your own computer.
RandComp Studio is a desktop app, not a service. There is no account to make and no RandComp server between you and your agents: the app, your projects, every agent’s terminal and every file it writes stay on your computer. It drives the AI command-line tools you already have — Claude Code, Codex, Kimi Code or Ollama — just as you would from a terminal.
RandComp Studio Claude Code Codex Kimi A model on this machine (Ollama)
📄 plan.md 📄 DESIGN.md 📁 your-project/ 📄 qa-report.md ⌨️ each agent’s terminal
What does leave your computer — all of it
Your prompts and code → the AI tool you picked. Claude Code talks to Anthropic, Codex to OpenAI, Kimi to Moonshot AI — on your own account, exactly as when you run them yourself. With a model on your computer (Ollama), nothing leaves.
Anonymous usage counts → RandComp. Four events: app opened (at most once a day), project created, MCP client connected, message sent. Each is only the event’s name, a UTC time and a random install ID — no content, no project names, no file paths. It is on by default; Settings → Usage statistics turns it off and deletes anything not yet sent.
Nothing else. RandComp account: none · RandComp copy of your work: none · Your prompts and files sent to RandComp: none. The app’s own server listens on 127.0.0.1 only.
Unplug and see. A role on a model on your computer (Ollama) keeps working without the internet. Roles on Claude Code, Codex or Kimi get no answers until you are back online, because their provider is on the internet. RandComp itself does not need it.
Honest limit: “local” means RandComp adds no cloud of its own — it does not make Claude Code, Codex or Kimi work offline. And small local models suit small, bounded work.
Not one AI grading its own homework.
The checker never saw the builder's reasoning, and may run on a different AI entirely. Mix freely — a strong model for planning, a cheaper one for routine work, a private one for sensitive files.
Planning on Claude Code, building on Codex, a quick check on Kimi, drafts on a model on your machine. Change it any time, even mid-project.
The checker works from the result, not from the builder's story, so blind spots are not shared. It is a second look, not a promise that everything is correct.
More agents cost more tokens. Small jobs stay a single session; teams are for work that earns it.
# everything is a plain file on your disk
reading-list/
plan.md coordinator
DESIGN.md designer
server/ web/ developers
qa-report.md independent QA
Open them in any editor. Move them, back them up, delete them. No export step, because there is nothing to export from.
Free during early preview. macOS only for now — Windows and Linux come later, once the one-person army has slept.
Download for macOS ↓ Read the docs
A Mac with Apple silicon (M1 or newer) on macOS 11 or later.
At least one AI command-line tool, installed and signed in — Claude Code, Codex or Kimi Code, on your own plan or API key with that company. Or Ollama with Qwen Code, for a free model that runs on your computer — fine for small, bounded work.
RandComp does not include AI usage. Each tool bills you exactly as it does today, and more helpers use more of it.
Apple silicon · macOS 11+ · 33 MB
SHA-256 88c7f816503e42b6a022be0ca4ca946ba948a03b955101bf140188c96c6503ab
It started while I was traveling. I wanted agents working on my Mac mini back home — but one long AI session goes out of shape fast: it drifts, loses the plot, and no one is there to catch it. So instead of one agent doing everything, I built a company: a coordinator plans the work, each role owns a piece, and I can attach to any of them from wherever I am. Then I kept using it myself — several small apps later, it had earned its place as my default way of working. That's when I decided it was time to share it. As for the name — I couldn't think of one, so Random Company it is. RandComp is just the shortcut.
