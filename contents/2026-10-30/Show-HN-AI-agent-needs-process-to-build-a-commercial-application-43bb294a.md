---
source: "https://nomorehappypath.com/"
hn_url: "https://news.ycombinator.com/item?id=50304765"
title: "Show HN: AI agent needs process to build a commercial application"
article_title: "NoMoreHappyPath: be a software company of one"
image: "https://nomorehappypath.com/assets/og-1200x630.png?v=2"
author: "abalgir"
captured_at: "2026-10-30T15:36:28Z"
capture_tool: "hn-digest"
hn_id: 50304765
score: 1
comments: 0
posted_at: "2026-10-30T14:49:17Z"
tags:
  - hacker-news
---

# Show HN: AI agent needs process to build a commercial application

- HN: [50304765](https://news.ycombinator.com/item?id=50304765)
- Source: [nomorehappypath.com](https://nomorehappypath.com/)
- Score: 1
- Comments: 0
- Posted: 2026-10-30T14:49:17Z

## Translation

Title: Show HN: AI agent needs process to build a commercial application
Article title: NoMoreHappyPath: be a software company of one
Description: Make ‘done’ mean done. An open-source delivery system for building commercial software with AI agents: 21 quality steps from your idea to accepted work. Free and open source.
HN text: I decided to test this claim with my app, NoMoreHappyPath (NMHP), against Claude on a real project. Here is the test: I built Acremuse, a commercial real estate analysis platform, with Claude Desktop and Codex as a reviewer. It is deployed and used today by a couple of investors. Then I wondered what I'd missed that might bite me later. ~ 374,000 Lines of program code - 1,070 Source files (Python and TypeScript) - ~ 11,400 Functions and methods - 3,067 Automated tests So I gave the same 374k-line code snapshot and the same brief to one Claude session that can use Codex as a reviewer and to my own tool, NMHP. My directive was the same for both: Analyze the code, find any errors/bugs, and list them by severity. Findings must be reviewed before writing the final report. The Results (Verified by an impartial, blind Codex judge): NMHP: 33 real issues found. 7 false claims. Claude: 10 real issues found. 2 false claims. Only 6 issues overlapped. Time & Cost:
From the owner's go-ahead to the reviewer's pass, NoMoreHappyPath took about 74 minutes, including time lost while it found and worked around a defect in the product itself. The Python version NMHP was working on was 3.9, the OS version, and it didn't have pytest installed, which was needed to verify some claims. After 15 minutes of blockage, the CTO stepped in and told the agent to create a venv with Python 3.13, install the corresponding pytest, and move the project along.
NMHP CTO reacts when 15 minutes pass without movement from agents. To be fair to Claude, Claude did create a venv from the get-go. Time & Cost: NMHP took 74 minutes and cost $24 in API spend. Claude finished in a few minutes for a fraction of that cost.
And yes, NMHP takes over an hour and costs twenty-four bucks. But for a 374k-line codebase, finding 23 more real issues overall that vanilla Claude missed makes $24 a massive win compared to human audit rates. I know this isn't a benchmark; it's one app and one run each. I wrote a detailed white paper on it. But it confirmed my idea: an AI agent without a real process is not yet capable of building a commercially ready application, even with a frontier model and reviews that are initiated by the same agent.

Article text:
NoMoreHappyPath
How it works
Platform
Proof
Install
About
GitHub →
Be a software company of one.
Make ‘done’ mean done.
NoMoreHappyPath is an open-source delivery system for building commercial software with AI agents. A builder, an independent reviewer from a different AI company and a CTO take your idea through 21 quality steps, in a control room that speaks plain English. Nothing is done until you accept it.
Free and open source (Apache 2.0) · macOS and Linux · runs on your machine, with your own AI accounts
A real build, sped up: a weather app from one plain-English request. 48 minutes, 4 independent reviews, 1 failed. Narration by AI.
Most AI says “done” in one step. This takes 21.
In plain English. Your exact words are captured, never altered.
Up to 4 multiple-choice questions, best answer pre-picked. Skip any time.
In plain words, with a “what will be tested” list.
From the contract, before any code exists.
Each test is labelled with the requirement(s) it covers; each requirement lists the failure cases it tries. A requirement shows “covered” or “gap”; a gap goes back once, and you are never held.
Stored with a fingerprint (SHA-256). Any change is recorded.
Run on the project as it was before the work began. A test that passes there is flagged “did not distinguish” and does not count.
A Scenario Ledger and an evidence contract. Independent parts run in parallel.
One failure and the review request is refused. Passing evidence is stored tamper-proof.
Records the real results. “Done” counts only if every final-check line says passed.
Walks the real product, in a real browser when it has a screen. A PASS is refused if any fails.
Every blocker, with its harm and the steps to reproduce it.
All at once. Tests are read-only during repair.
Only four kinds of problem can block. New nitpicks become notes.
If the review can't be completed, it retries once on a higher model set in Settings.
If repairs keep failing, a fresh builder retries once on a higher model set in Settings.
Shown with the reason. Nothing ships.
Every part settled, the full test suite on the combined work. The CTO release gate blocks if evidence is missing or tampered.
Accept (Windows: Get the build, with its SHA-256) or Not accepted with your reason; it goes straight back into repair.
It reaches the project's main.
You Describe the task In plain English. Your exact words are captured, never altered.
Reviewer Frames the task Up to 4 multiple-choice questions, best answer pre-picked. Skip any time.
You Approve the contract In plain words, with a “what will be tested” list.
Reviewer Writes the tests first From the contract, before any code exists.
Reviewer Coverage check Each test is labelled with the requirement(s) it covers; each requirement lists the failure cases it tries. A requirement shows “covered” or “gap”; a gap goes back once, and you are never held.
Board Locks the tests Stored with a fingerprint (SHA-256). Any change is recorded.
Board Red proof Run on the project as it was before the work began. A test that passes there is flagged “did not distinguish” and does not count.
Builder Plans the work A Scenario Ledger and an evidence contract. Independent parts run in parallel.
Builder Builds To the locked tests.
Builder Runs every scenario itself One failure and the review request is refused. Passing evidence is stored tamper-proof.
Board Runs the locked tests Records the real results. “Done” counts only if every final-check line says passed.
Reviewer Runs its own challenges Walks the real product, in a real browser when it has a screen. A PASS is refused if any fails.
Reviewer Review, round 1 Every blocker, with its harm and the steps to reproduce it.
Builder Repairs the whole list All at once. Tests are read-only during repair.
Reviewer Rounds 2 and 3 Only four kinds of problem can block. New nitpicks become notes.
Reviewer Reviewer net If the review can't be completed, it retries once on a higher model set in Settings.
Builder Builder net If repairs keep failing, a fresh builder retries once on a higher model set in Settings.
Board Could not finish Shown with the reason. Nothing ships.
Board Final gate Every part settled, the full test suite on the combined work. The CTO release gate blocks if evidence is missing or tampered.
You Review and decide Accept (Windows: Get the build, with its SHA-256) or Not accepted with your reason; it goes straight back into repair.
Board Accepted work lands It reaches the project's main.
What it looks like in Mission Control
A real task, running: what is happening now, which step it is on, what it has cost so far, and what each agent is doing.
Quality that is proven, not promised
A reviewer from a different AI company reads your requirements and writes the tests before any code exists. A coverage check makes sure every requirement has a test, including what happens when things go wrong.
You don’t trust a smoke alarm until you’ve tested it with smoke. The board runs each test before the work: a test that passes on nothing proves nothing, so it doesn’t count.
Round one lists every real blocker, with its harm. Rounds two and three only block on four kinds of problem. Then a retry on a higher model, or an honest “could not finish”. Never an endless loop.
Everything a software company needs, on your machine
Underneath it all: a board the agents can’t edit, an operating-system sandbox around every agent, and self-healing when an agent stalls.
A friendly interface to talk to the agents and see where each task is, step by step, explained in plain English.
It doesn’t only monitor: it clears the way for the builder and the reviewer so nothing stays stuck, and never asks you a technical question.
The right model, within your quota
Claude, Codex or Gemini in any role. Specs on higher models, code on lower ones, an automatic switch to a higher model when needed, and a check on your quota.
Define virtual machines of different operating systems as test targets. A Windows program is built and tested on a real Windows machine, from your Mac.
Ask questions about your project, look at its full history, and pause and restart without losing anything.
Every app gets a built-in security baseline. Switch on Enhanced security for a task and it adds a threat model, security tests and a controls register, ready for a security assessment or audit.
Real screens from the products above, built with NoMoreHappyPath as it grew, so not every one went through every step shown here. And real review records: passing tests didn’t stop the reviewer from failing the work.
PropDesk, built with NoMoreHappyPath: parcel, utilities, flood, zoning, traffic and more in one workspace. Narration by AI.
It fails its own work, and says so
What you need, and two commands
The platform is free. The AI agents it runs are not: they use your own accounts, and a serious build costs real money on them.
git clone https://github.com/nomorehappypath/nomorehappypath.git
cd nomorehappypath && bash install.sh
The installer checks everything above and tells you plainly what’s missing. Then open Mission Control, point it at a project folder, and give your first direction.
Built by someone who has shipped commercial software for 35 years
I’m Gerard Rafie. For more than 35 years I led product management, marketing and sales in fintech: I conceived, built, marketed and sold technology solutions for banks, and I advise banks and private equity firms on treasury, capital markets and AI.
I built NoMoreHappyPath because AI kept telling me “done” when it wasn’t. I use it for my own products, and so do my clients.
Many harnesses help you run agents, and leave the quality process for you to design. NoMoreHappyPath ships the whole route by default, and it's built for a solo builder who wants to ship commercial software, not demos:
Agreed before built. The reviewer asks you up to four multiple-choice questions, and you approve a plain-words contract before any code exists.
Tests that are real. A reviewer from a different AI company writes the tests first; a coverage check makes sure every requirement has one; every test must fail before the work, or it doesn't count.
A board no agent can edit. The board runs the tests itself and records the real results, so a pass can't be faked.
Reviews that end. Three bounded rounds, then a retry on a higher model, or an honest “could not finish”. Never an endless loop.
A control room you can understand. Mission Control shows every step in plain English, and a CTO clears the path so nothing stays stuck.
Built for real targets. Test on virtual machines of different operating systems, pick the model for each role within your quota, and pause or restart without losing anything.
Security built in, and enhanced on demand. A security baseline on every app, and an Enhanced security switch per task that adds a threat model, security tests and a controls register.
You decide. Nothing is done until you review and accept it.
Proven on real products. A land due-diligence platform, a pre-purchase analysis app and a marketing agency, each used by a client.
The reviewer isn’t reading a diff and nodding. It writes its own tests, the harness runs them and keeps the output, and the verdict has to match that output. A pass with no recorded execution behind it is refused. It also has no reason to agree: it comes from a different company than the builder.
The software is free. The agents run on your own subscriptions. A serious build uses real usage. Using a top model for the plan and lighter models for building and review keeps that down.
Because the harness builds for release, not for a demo. It agrees the requirements first, has a different AI review each part, tries the failure cases, checks security and quality basics, and finishes with a review that uses the whole app in a real browser. A weather app took 48 minutes and four reviews, and one of them failed. A demo takes five minutes and skips all of that.
The harness runs on your machine and doesn’t send your code to us. Your agents talk to their own providers, the same as when you use those tools directly.
No, and nothing does. What it guarantees is that nothing ships unchecked: the tests are written before the code and must fail first, the board runs them itself, an independent reviewer from a different AI company runs its own tests on the real product, a final gate runs the full test suite, and you review and accept the result before it ships. Every check leaves evidence you can read.
Free, open source, and built under its own rules.

## Original Extract

Make ‘done’ mean done. An open-source delivery system for building commercial software with AI agents: 21 quality steps from your idea to accepted work. Free and open source.

I decided to test this claim with my app, NoMoreHappyPath (NMHP), against Claude on a real project. Here is the test: I built Acremuse, a commercial real estate analysis platform, with Claude Desktop and Codex as a reviewer. It is deployed and used today by a couple of investors. Then I wondered what I'd missed that might bite me later. ~ 374,000 Lines of program code - 1,070 Source files (Python and TypeScript) - ~ 11,400 Functions and methods - 3,067 Automated tests So I gave the same 374k-line code snapshot and the same brief to one Claude session that can use Codex as a reviewer and to my own tool, NMHP. My directive was the same for both: Analyze the code, find any errors/bugs, and list them by severity. Findings must be reviewed before writing the final report. The Results (Verified by an impartial, blind Codex judge): NMHP: 33 real issues found. 7 false claims. Claude: 10 real issues found. 2 false claims. Only 6 issues overlapped. Time & Cost:
From the owner's go-ahead to the reviewer's pass, NoMoreHappyPath took about 74 minutes, including time lost while it found and worked around a defect in the product itself. The Python version NMHP was working on was 3.9, the OS version, and it didn't have pytest installed, which was needed to verify some claims. After 15 minutes of blockage, the CTO stepped in and told the agent to create a venv with Python 3.13, install the corresponding pytest, and move the project along.
NMHP CTO reacts when 15 minutes pass without movement from agents. To be fair to Claude, Claude did create a venv from the get-go. Time & Cost: NMHP took 74 minutes and cost $24 in API spend. Claude finished in a few minutes for a fraction of that cost.
And yes, NMHP takes over an hour and costs twenty-four bucks. But for a 374k-line codebase, finding 23 more real issues overall that vanilla Claude missed makes $24 a massive win compared to human audit rates. I know this isn't a benchmark; it's one app and one run each. I wrote a detailed white paper on it. But it confirmed my idea: an AI agent without a real process is not yet capable of building a commercially ready application, even with a frontier model and reviews that are initiated by the same agent.

NoMoreHappyPath
How it works
Platform
Proof
Install
About
GitHub →
Be a software company of one.
Make ‘done’ mean done.
NoMoreHappyPath is an open-source delivery system for building commercial software with AI agents. A builder, an independent reviewer from a different AI company and a CTO take your idea through 21 quality steps, in a control room that speaks plain English. Nothing is done until you accept it.
Free and open source (Apache 2.0) · macOS and Linux · runs on your machine, with your own AI accounts
A real build, sped up: a weather app from one plain-English request. 48 minutes, 4 independent reviews, 1 failed. Narration by AI.
Most AI says “done” in one step. This takes 21.
In plain English. Your exact words are captured, never altered.
Up to 4 multiple-choice questions, best answer pre-picked. Skip any time.
In plain words, with a “what will be tested” list.
From the contract, before any code exists.
Each test is labelled with the requirement(s) it covers; each requirement lists the failure cases it tries. A requirement shows “covered” or “gap”; a gap goes back once, and you are never held.
Stored with a fingerprint (SHA-256). Any change is recorded.
Run on the project as it was before the work began. A test that passes there is flagged “did not distinguish” and does not count.
A Scenario Ledger and an evidence contract. Independent parts run in parallel.
One failure and the review request is refused. Passing evidence is stored tamper-proof.
Records the real results. “Done” counts only if every final-check line says passed.
Walks the real product, in a real browser when it has a screen. A PASS is refused if any fails.
Every blocker, with its harm and the steps to reproduce it.
All at once. Tests are read-only during repair.
Only four kinds of problem can block. New nitpicks become notes.
If the review can't be completed, it retries once on a higher model set in Settings.
If repairs keep failing, a fresh builder retries once on a higher model set in Settings.
Shown with the reason. Nothing ships.
Every part settled, the full test suite on the combined work. The CTO release gate blocks if evidence is missing or tampered.
Accept (Windows: Get the build, with its SHA-256) or Not accepted with your reason; it goes straight back into repair.
It reaches the project's main.
You Describe the task In plain English. Your exact words are captured, never altered.
Reviewer Frames the task Up to 4 multiple-choice questions, best answer pre-picked. Skip any time.
You Approve the contract In plain words, with a “what will be tested” list.
Reviewer Writes the tests first From the contract, before any code exists.
Reviewer Coverage check Each test is labelled with the requirement(s) it covers; each requirement lists the failure cases it tries. A requirement shows “covered” or “gap”; a gap goes back once, and you are never held.
Board Locks the tests Stored with a fingerprint (SHA-256). Any change is recorded.
Board Red proof Run on the project as it was before the work began. A test that passes there is flagged “did not distinguish” and does not count.
Builder Plans the work A Scenario Ledger and an evidence contract. Independent parts run in parallel.
Builder Builds To the locked tests.
Builder Runs every scenario itself One failure and the review request is refused. Passing evidence is stored tamper-proof.
Board Runs the locked tests Records the real results. “Done” counts only if every final-check line says passed.
Reviewer Runs its own challenges Walks the real product, in a real browser when it has a screen. A PASS is refused if any fails.
Reviewer Review, round 1 Every blocker, with its harm and the steps to reproduce it.
Builder Repairs the whole list All at once. Tests are read-only during repair.
Reviewer Rounds 2 and 3 Only four kinds of problem can block. New nitpicks become notes.
Reviewer Reviewer net If the review can't be completed, it retries once on a higher model set in Settings.
Builder Builder net If repairs keep failing, a fresh builder retries once on a higher model set in Settings.
Board Could not finish Shown with the reason. Nothing ships.
Board Final gate Every part settled, the full test suite on the combined work. The CTO release gate blocks if evidence is missing or tampered.
You Review and decide Accept (Windows: Get the build, with its SHA-256) or Not accepted with your reason; it goes straight back into repair.
Board Accepted work lands It reaches the project's main.
What it looks like in Mission Control
A real task, running: what is happening now, which step it is on, what it has cost so far, and what each agent is doing.
Quality that is proven, not promised
A reviewer from a different AI company reads your requirements and writes the tests before any code exists. A coverage check makes sure every requirement has a test, including what happens when things go wrong.
You don’t trust a smoke alarm until you’ve tested it with smoke. The board runs each test before the work: a test that passes on nothing proves nothing, so it doesn’t count.
Round one lists every real blocker, with its harm. Rounds two and three only block on four kinds of problem. Then a retry on a higher model, or an honest “could not finish”. Never an endless loop.
Everything a software company needs, on your machine
Underneath it all: a board the agents can’t edit, an operating-system sandbox around every agent, and self-healing when an agent stalls.
A friendly interface to talk to the agents and see where each task is, step by step, explained in plain English.
It doesn’t only monitor: it clears the way for the builder and the reviewer so nothing stays stuck, and never asks you a technical question.
The right model, within your quota
Claude, Codex or Gemini in any role. Specs on higher models, code on lower ones, an automatic switch to a higher model when needed, and a check on your quota.
Define virtual machines of different operating systems as test targets. A Windows program is built and tested on a real Windows machine, from your Mac.
Ask questions about your project, look at its full history, and pause and restart without losing anything.
Every app gets a built-in security baseline. Switch on Enhanced security for a task and it adds a threat model, security tests and a controls register, ready for a security assessment or audit.
Real screens from the products above, built with NoMoreHappyPath as it grew, so not every one went through every step shown here. And real review records: passing tests didn’t stop the reviewer from failing the work.
PropDesk, built with NoMoreHappyPath: parcel, utilities, flood, zoning, traffic and more in one workspace. Narration by AI.
It fails its own work, and says so
What you need, and two commands
The platform is free. The AI agents it runs are not: they use your own accounts, and a serious build costs real money on them.
git clone https://github.com/nomorehappypath/nomorehappypath.git
cd nomorehappypath && bash install.sh
The installer checks everything above and tells you plainly what’s missing. Then open Mission Control, point it at a project folder, and give your first direction.
Built by someone who has shipped commercial software for 35 years
I’m Gerard Rafie. For more than 35 years I led product management, marketing and sales in fintech: I conceived, built, marketed and sold technology solutions for banks, and I advise banks and private equity firms on treasury, capital markets and AI.
I built NoMoreHappyPath because AI kept telling me “done” when it wasn’t. I use it for my own products, and so do my clients.
Many harnesses help you run agents, and leave the quality process for you to design. NoMoreHappyPath ships the whole route by default, and it's built for a solo builder who wants to ship commercial software, not demos:
Agreed before built. The reviewer asks you up to four multiple-choice questions, and you approve a plain-words contract before any code exists.
Tests that are real. A reviewer from a different AI company writes the tests first; a coverage check makes sure every requirement has one; every test must fail before the work, or it doesn't count.
A board no agent can edit. The board runs the tests itself and records the real results, so a pass can't be faked.
Reviews that end. Three bounded rounds, then a retry on a higher model, or an honest “could not finish”. Never an endless loop.
A control room you can understand. Mission Control shows every step in plain English, and a CTO clears the path so nothing stays stuck.
Built for real targets. Test on virtual machines of different operating systems, pick the model for each role within your quota, and pause or restart without losing anything.
Security built in, and enhanced on demand. A security baseline on every app, and an Enhanced security switch per task that adds a threat model, security tests and a controls register.
You decide. Nothing is done until you review and accept it.
Proven on real products. A land due-diligence platform, a pre-purchase analysis app and a marketing agency, each used by a client.
The reviewer isn’t reading a diff and nodding. It writes its own tests, the harness runs them and keeps the output, and the verdict has to match that output. A pass with no recorded execution behind it is refused. It also has no reason to agree: it comes from a different company than the builder.
The software is free. The agents run on your own subscriptions. A serious build uses real usage. Using a top model for the plan and lighter models for building and review keeps that down.
Because the harness builds for release, not for a demo. It agrees the requirements first, has a different AI review each part, tries the failure cases, checks security and quality basics, and finishes with a review that uses the whole app in a real browser. A weather app took 48 minutes and four reviews, and one of them failed. A demo takes five minutes and skips all of that.
The harness runs on your machine and doesn’t send your code to us. Your agents talk to their own providers, the same as when you use those tools directly.
No, and nothing does. What it guarantees is that nothing ships unchecked: the tests are written before the code and must fail first, the board runs them itself, an independent reviewer from a different AI company runs its own tests on the real product, a final gate runs the full test suite, and you review and accept the result before it ships. Every check leaves evidence you can read.
Free, open source, and built under its own rules.
