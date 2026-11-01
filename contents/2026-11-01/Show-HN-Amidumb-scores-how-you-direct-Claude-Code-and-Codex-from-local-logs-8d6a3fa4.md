---
source: "https://github.com/jjinendra3/amidumb"
hn_url: "https://news.ycombinator.com/item?id=50326432"
title: "Show HN: Amidumb – scores how you direct Claude Code and Codex, from local logs"
article_title: "GitHub - jjinendra3/amidumb · GitHub"
image: "https://repository-images.githubusercontent.com/1440623275/7731c013-8016-46d9-8617-fdc3bf894330"
author: "jjinendra3"
captured_at: "2026-11-01T15:49:35Z"
capture_tool: "hn-digest"
hn_id: 50326432
score: 1
comments: 0
posted_at: "2026-11-01T15:15:03Z"
tags:
  - hacker-news
---

# Show HN: Amidumb – scores how you direct Claude Code and Codex, from local logs

- HN: [50326432](https://news.ycombinator.com/item?id=50326432)
- Source: [github.com](https://github.com/jjinendra3/amidumb)
- Score: 1
- Comments: 0
- Posted: 2026-11-01T15:15:03Z

## Translation

Title: Show HN: Amidumb – scores how you direct Claude Code and Codex, from local logs
Article title: GitHub - jjinendra3/amidumb · GitHub
Description: Contribute to jjinendra3/amidumb development by creating an account on GitHub.

Article text:
GitHub - jjinendra3/amidumb · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
jjinendra3
amidumb Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
www.npmjs.com/package/amidumb Topics
Readme MIT license Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows assets/ fonts assets/ fonts docs docs scripts scripts src src test test .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts View all files Repository files navigation
Your AI coding transcripts, judged.
Do you direct your coding agent like an engineer, or do you just vibecode?
amidumb reads your Claude Code and Codex history on your machine, scores how well you direct the agent from 0 to 100, gives you a title for it, gets your own claude to roast you, and saves a card you can post. One command, no sign-up, no API key, and the score never leaves your machine.
A direction score from 0 to 100, with five sub-scores: specification, steering, verification, discipline and leverage.
A title for every 5 points of your score, from Hostage to Agent's Worst Nightmare. All twenty of them are roasts.
The roast : a one-line verdict and three reasons for your score, written by your own claude (or codex ). Every reason has to cite prompts that were actually sent; a reason that cites a prompt that doesn't exist is thrown out.
A card : amidumb-YYYY-MM-DD.png (1080×1350), saved where you ran it and opened for you.
A rage meter : prompts with swearing, CAPS LOCK or !!! .
A secret leak count : API keys, tokens and passwords you've pasted into prompts.
Three habits to fix , written as rules you can paste into CLAUDE.md or AGENTS.md .
The title comes from your score alone, one for every 5 points. It's a ladder of who's in charge: at the bottom the agent runs you, at the top you run it too hard. None of them is a compliment, the top one included.
The score is computed on your machine. amidumb reads ~/.claude/projects and ~/.codex/sessions . It has no servers, no telemetry and makes no network calls of its own.
The roast runs on your own CLI, on your own plan. It sends a sample of about 150 prompts to your claude CLI (or codex ), and prints one line saying exactly what it's sending before it does. If that CLI isn't logged in, it says so and skips the roast.
Only what you typed and what the agent said. Each prompt is cut to 500 characters and goes with the agent's message before it, cut to 300. Tool output and the files the agent read or wrote are never sent.
Secrets are redacted before anything is sent : private keys, AWS, GitHub, Anthropic, OpenAI, Stripe, Slack, Google and npm keys, JWTs, and password= -style assignments.
--dry-run prints exactly what would be sent, and sends nothing. --no-roast gives you the score and the card with nothing sent anywhere.
The judge leaves no trace. It runs ephemeral ( claude --no-session-persistence with no tools, codex exec --ephemeral in a read-only sandbox) from an empty folder, so it never ends up in the transcripts it judges.
The card never quotes your prompts. It shows scores and what the judge wrote. The judge is told to leave out paths, code and names, and its text is masked again before rendering.
Does it send my code?
Not your files. If you pasted code into a prompt, the first 500 characters of that prompt can be sent, after secrets are redacted. Run --dry-run to see every character that would go.
What does it cost?
About 20–30k tokens from your own Claude or Codex plan, once per run. Nothing else.
Can I rerun it until I get a better score?
Rerun all you like: the same transcripts always get the same score. The score comes from a fixed rubric computed locally, never from the LLM. Only the roast's wording changes.
I use Codex, not Claude Code.
amidumb reads Codex sessions too. The roast uses claude by default; pass --codex to use Codex instead. If Claude Code isn't installed, it uses Codex on its own. Sub-scores built on signals Codex doesn't record are left out, not counted against you.
My history looks short.
Claude Code deletes transcripts after 30 days unless you raise cleanupPeriodDays in its settings. amidumb reads everything still on disk; --days 30 scores just the last month.
My score is wrong.
Open an issue with your npx amidumb --json output. It's numbers only: no prompt text, titles or paths. If the numbers themselves look off, npx amidumb survey prints the record types and keys in your transcripts, with no content.
How the score works (rubric v2)
The score never comes from the LLM. An LLM score would change on every run and go easy on whoever is paying for it. The score is a fixed, versioned rubric over signals computed locally. The LLM only writes the words.
Every prompt is classified in the context of the agent message right before it, because a short prompt isn't a dumb prompt: "yes, go" after a good plan is good direction, and "fix it" on repeat is vibecoding. Those classes feed five sub-scores, weighted into the overall score.
Every prompt is judged in context
class
what it is
specify
names a file or symbol, sets a constraint, or says what done looks like
instruct
asks for work without any of those
question
asks for information
approve
"yes", "go", "continue" right after the agent proposed a plan or asked a question
nudge
the same words when nothing was proposed
correct
pushes back with new information: "no, use the cache instead", "you missed the mobile layout"
verify
asks for tests, a review or the diff
retry
"fix it", "still broken", "try again", with nothing new
error_paste
a pasted error or log, with barely a word of your own
Sub-scores
Each sub-score is computed per session, then averaged with every session weighted by its prompt count, capped at 10. Short sessions count for less, and no single marathon session decides your score.
A sub-score that doesn't apply is left out and the other weights are rescaled. You need at least 5 sessions with two or more prompts, and 30 prompts in total.
The honest answer is "better at the extremes than in the middle".
Reference personas are synthetic users at the extremes everyone agrees on, and they're tests. Someone who mostly pastes errors must score under 30, a retry-looper under 35, a careful spec-writer over 75, and a spec-writer who also verifies and steers over 80. Any rubric change has to keep them in their bands.
Calibration: the sub-scores were checked against an independent rater that never saw the rubric, on 68 sessions from one heavy user, split into a tuning set and a held-out set.
On the held-out set, the opening score ranked sessions the way the rater did 70% of the time; 50% is a coin flip.
Verification was about 55%. Most checking (running the app, reading the diff) never shows up in a transcript.
Steering couldn't be tested, because that user never got stuck in a retry loop.
Scores probably skew high for anyone who doesn't retry-loop or paste errors.
Any change to the rubric bumps its version.
amidumb [options] your direction score, from Claude Code and Codex sessions
amidumb label [options] label your own sessions (blind) to sanity-check the rubric
amidumb calibrate how well the rubric ranks your labelled sessions
amidumb survey [options] record kinds and keys in your transcripts (no content)
--days <n|all> window to read (default: all history on disk)
--raw the full table of raw signals instead of the score
--sessions one row per session, with titles (stays local)
--sample <n> print n random detected prompts to check detection (stays local)
--json numbers only, no text, titles or paths: safe to share
--no-roast score and card only, nothing is sent anywhere
--dry-run print exactly what the roast would send, and send nothing
--codex roast with your codex CLI instead of claude
--no-open save the card image without opening it
--claude-dir <p> Claude Code transcripts (default ~/.claude/projects)
--codex-dir <p> Codex transcripts (default ~/.codex/sessions)
Supported
Claude Code in the terminal, VS Code and the desktop app, and Codex in the CLI, VS Code and the desktop app. Cursor comes later.
Node 18.3 or newer on macOS, Linux and Windows. CI runs the built CLI end to end on all three, with Node 18, 20 and 22, against a fake judge.
Transcript formats aren't documented and change between versions, so amidumb parses defensively and skips what it doesn't recognise.
pnpm install
pnpm typecheck && pnpm lint && pnpm test
pnpm build && pnpm smoke # end to end with a fake judge, no tokens
node dist/cli.js --dry-run
License
MIT. The package bundles third-party code, including satori (MPL-2.0) and zod (MIT); dist/THIRD_PARTY_LICENSES.txt lists every bundled package with its license and source. The Inter font is under the SIL Open Font License ( assets/fonts/OFL.txt ).
www.npmjs.com/package/amidumb Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to jjinendra3/amidumb development by creating an account on GitHub.

GitHub - jjinendra3/amidumb · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
jjinendra3
amidumb Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
www.npmjs.com/package/amidumb Topics
Readme MIT license Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows assets/ fonts assets/ fonts docs docs scripts scripts src src test test .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts View all files Repository files navigation
Your AI coding transcripts, judged.
Do you direct your coding agent like an engineer, or do you just vibecode?
amidumb reads your Claude Code and Codex history on your machine, scores how well you direct the agent from 0 to 100, gives you a title for it, gets your own claude to roast you, and saves a card you can post. One command, no sign-up, no API key, and the score never leaves your machine.
A direction score from 0 to 100, with five sub-scores: specification, steering, verification, discipline and leverage.
A title for every 5 points of your score, from Hostage to Agent's Worst Nightmare. All twenty of them are roasts.
The roast : a one-line verdict and three reasons for your score, written by your own claude (or codex ). Every reason has to cite prompts that were actually sent; a reason that cites a prompt that doesn't exist is thrown out.
A card : amidumb-YYYY-MM-DD.png (1080×1350), saved where you ran it and opened for you.
A rage meter : prompts with swearing, CAPS LOCK or !!! .
A secret leak count : API keys, tokens and passwords you've pasted into prompts.
Three habits to fix , written as rules you can paste into CLAUDE.md or AGENTS.md .
The title comes from your score alone, one for every 5 points. It's a ladder of who's in charge: at the bottom the agent runs you, at the top you run it too hard. None of them is a compliment, the top one included.
The score is computed on your machine. amidumb reads ~/.claude/projects and ~/.codex/sessions . It has no servers, no telemetry and makes no network calls of its own.
The roast runs on your own CLI, on your own plan. It sends a sample of about 150 prompts to your claude CLI (or codex ), and prints one line saying exactly what it's sending before it does. If that CLI isn't logged in, it says so and skips the roast.
Only what you typed and what the agent said. Each prompt is cut to 500 characters and goes with the agent's message before it, cut to 300. Tool output and the files the agent read or wrote are never sent.
Secrets are redacted before anything is sent : private keys, AWS, GitHub, Anthropic, OpenAI, Stripe, Slack, Google and npm keys, JWTs, and password= -style assignments.
--dry-run prints exactly what would be sent, and sends nothing. --no-roast gives you the score and the card with nothing sent anywhere.
The judge leaves no trace. It runs ephemeral ( claude --no-session-persistence with no tools, codex exec --ephemeral in a read-only sandbox) from an empty folder, so it never ends up in the transcripts it judges.
The card never quotes your prompts. It shows scores and what the judge wrote. The judge is told to leave out paths, code and names, and its text is masked again before rendering.
Does it send my code?
Not your files. If you pasted code into a prompt, the first 500 characters of that prompt can be sent, after secrets are redacted. Run --dry-run to see every character that would go.
What does it cost?
About 20–30k tokens from your own Claude or Codex plan, once per run. Nothing else.
Can I rerun it until I get a better score?
Rerun all you like: the same transcripts always get the same score. The score comes from a fixed rubric computed locally, never from the LLM. Only the roast's wording changes.
I use Codex, not Claude Code.
amidumb reads Codex sessions too. The roast uses claude by default; pass --codex to use Codex instead. If Claude Code isn't installed, it uses Codex on its own. Sub-scores built on signals Codex doesn't record are left out, not counted against you.
My history looks short.
Claude Code deletes transcripts after 30 days unless you raise cleanupPeriodDays in its settings. amidumb reads everything still on disk; --days 30 scores just the last month.
My score is wrong.
Open an issue with your npx amidumb --json output. It's numbers only: no prompt text, titles or paths. If the numbers themselves look off, npx amidumb survey prints the record types and keys in your transcripts, with no content.
How the score works (rubric v2)
The score never comes from the LLM. An LLM score would change on every run and go easy on whoever is paying for it. The score is a fixed, versioned rubric over signals computed locally. The LLM only writes the words.
Every prompt is classified in the context of the agent message right before it, because a short prompt isn't a dumb prompt: "yes, go" after a good plan is good direction, and "fix it" on repeat is vibecoding. Those classes feed five sub-scores, weighted into the overall score.
Every prompt is judged in context
class
what it is
specify
names a file or symbol, sets a constraint, or says what done looks like
instruct
asks for work without any of those
question
asks for information
approve
"yes", "go", "continue" right after the agent proposed a plan or asked a question
nudge
the same words when nothing was proposed
correct
pushes back with new information: "no, use the cache instead", "you missed the mobile layout"
verify
asks for tests, a review or the diff
retry
"fix it", "still broken", "try again", with nothing new
error_paste
a pasted error or log, with barely a word of your own
Sub-scores
Each sub-score is computed per session, then averaged with every session weighted by its prompt count, capped at 10. Short sessions count for less, and no single marathon session decides your score.
A sub-score that doesn't apply is left out and the other weights are rescaled. You need at least 5 sessions with two or more prompts, and 30 prompts in total.
The honest answer is "better at the extremes than in the middle".
Reference personas are synthetic users at the extremes everyone agrees on, and they're tests. Someone who mostly pastes errors must score under 30, a retry-looper under 35, a careful spec-writer over 75, and a spec-writer who also verifies and steers over 80. Any rubric change has to keep them in their bands.
Calibration: the sub-scores were checked against an independent rater that never saw the rubric, on 68 sessions from one heavy user, split into a tuning set and a held-out set.
On the held-out set, the opening score ranked sessions the way the rater did 70% of the time; 50% is a coin flip.
Verification was about 55%. Most checking (running the app, reading the diff) never shows up in a transcript.
Steering couldn't be tested, because that user never got stuck in a retry loop.
Scores probably skew high for anyone who doesn't retry-loop or paste errors.
Any change to the rubric bumps its version.
amidumb [options] your direction score, from Claude Code and Codex sessions
amidumb label [options] label your own sessions (blind) to sanity-check the rubric
amidumb calibrate how well the rubric ranks your labelled sessions
amidumb survey [options] record kinds and keys in your transcripts (no content)
--days <n|all> window to read (default: all history on disk)
--raw the full table of raw signals instead of the score
--sessions one row per session, with titles (stays local)
--sample <n> print n random detected prompts to check detection (stays local)
--json numbers only, no text, titles or paths: safe to share
--no-roast score and card only, nothing is sent anywhere
--dry-run print exactly what the roast would send, and send nothing
--codex roast with your codex CLI instead of claude
--no-open save the card image without opening it
--claude-dir <p> Claude Code transcripts (default ~/.claude/projects)
--codex-dir <p> Codex transcripts (default ~/.codex/sessions)
Supported
Claude Code in the terminal, VS Code and the desktop app, and Codex in the CLI, VS Code and the desktop app. Cursor comes later.
Node 18.3 or newer on macOS, Linux and Windows. CI runs the built CLI end to end on all three, with Node 18, 20 and 22, against a fake judge.
Transcript formats aren't documented and change between versions, so amidumb parses defensively and skips what it doesn't recognise.
pnpm install
pnpm typecheck && pnpm lint && pnpm test
pnpm build && pnpm smoke # end to end with a fake judge, no tokens
node dist/cli.js --dry-run
License
MIT. The package bundles third-party code, including satori (MPL-2.0) and zod (MIT); dist/THIRD_PARTY_LICENSES.txt lists every bundled package with its license and source. The Inter font is under the SIL Open Font License ( assets/fonts/OFL.txt ).
www.npmjs.com/package/amidumb Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
