---
source: "https://github.com/eremes81/game-design-ai-practice-en"
hn_url: "https://news.ycombinator.com/item?id=49671342"
title: "Show HN: I wrote a book on using AI in game design production (free on GitHub)"
article_title: "GitHub - eremes81/game-design-ai-practice-en: Free book: AI Workflow for Game Designers — No Fabricated Numbers. Full text free in this repo (CC BY-NC-SA). English edition of the Korean print original (ISBN 979-11-12-21479-9). · GitHub"
image: "https://repository-images.githubusercontent.com/1267465560/6949e6b9-bc9d-4504-a2da-b68fa1c29d2d"
author: "eremes81"
captured_at: "2026-09-12T12:39:25Z"
capture_tool: "hn-digest"
hn_id: 49671342
score: 3
comments: 0
posted_at: "2026-09-12T11:44:34Z"
tags:
  - hacker-news
---

# Show HN: I wrote a book on using AI in game design production (free on GitHub)

- HN: [49671342](https://news.ycombinator.com/item?id=49671342)
- Source: [github.com](https://github.com/eremes81/game-design-ai-practice-en)
- Score: 3
- Comments: 0
- Posted: 2026-09-12T11:44:34Z

## Translation

Title: Show HN: I wrote a book on using AI in game design production (free on GitHub)
Article title: GitHub - eremes81/game-design-ai-practice-en: Free book: AI Workflow for Game Designers — No Fabricated Numbers. Full text free in this repo (CC BY-NC-SA). English edition of the Korean print original (ISBN 979-11-12-21479-9). · GitHub
Description: Free book: AI Workflow for Game Designers — No Fabricated Numbers. Full text free in this repo (CC BY-NC-SA). English edition of the Korean print original (ISBN 979-11-12-21479-9). - eremes81/game-design-ai-practice-en

Article text:
GitHub - eremes81/game-design-ai-practice-en: Free book: AI Workflow for Game Designers — No Fabricated Numbers. Full text free in this repo (CC BY-NC-SA). English edition of the Korean print original (ISBN 979-11-12-21479-9). · GitHub
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
eremes81
/
game-design-ai-practice-en
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
21 Commits 21 Commits Folders and files
assets assets manuscript manuscript .gitignore .gitignore .nojekyll .nojekyll LICENSE LICENSE README.md README.md index.html index.html robots.txt robots.txt sitemap.xml sitemap.xml View all files Repository files navigation
AI Workflow for Game Designers
No Fabricated Numbers — A Six-Month Field Manual for Claude Code, Prompts, Validation, and Production Memory
🌐 Editions: 한국어 — Original · English · 日本語 · ไทย · Bahasa Indonesia · 简体中文 · 繁體中文
🧰 Companion tooling — Harness Starter Kit
A minimal, runnable version of the memory, hook and retrospective structure this book builds in 1.3, Part 21 and Part 24. Download the folder, open it in an agent, and it works.
Built as the appendix to the author's other book (for general office work), but the skeleton is the same one described here. MIT licensed.
A hands-on field manual by Minsoo Lee — a game designer of 24 years, former design team lead on Ragnarok Online (2002–2007), now a Design Director shipping an MMORPG in Korea — on bringing generative AI (Claude Code) into daily production work . Not theory or forecasts — it walks one task at a time from the very first screen (install, accounts, pricing) through systems design, combat, narrative, level design, balance, UX, and live ops, all the way to turning meeting notes into decisions, validation gates, cost management, and copyright.
The subtitle — "No Fabricated Numbers" — is this book's promise. Every figure and case in the text comes from real work, not invented examples, and most of the code runs as is on the Python standard library alone.
Schema-first data design — the table schema and legal value ranges come before the AI writes a single row
Validation gates that catch AI-invented numbers, IDs, and dangling references before they reach the data sheets
Meeting notes → tracked decisions — decision, owner, and rationale extracted and greppable
AI cost management in code — token budgets enforced by scripts, not by good intentions
Balance simulation, procedural content pipelines, UX lint, live-ops workflows, and a team adoption strategy for leads
The book is open to readers outside games, too. Workflows like turning meeting notes into decisions, tracing a decision's ripple effects, and guarding quality with validation gates work regardless of your job. The "Beyond Games" box in each chapter is the bridge, and every chapter also carries a "Solo Scale-Down" for people building alone, with no team.
This is the English edition of the Korean original, 게임 기획 실무에서 바로 쓰는 AI·클로드 코드 활용법 (BOOKK, 2026 · ISBN 979-11-12-21479-9).
In the spirit of this book's honesty-first principle, here is exactly how this edition was made: it was translated from the Korean original using the book's own AI workflow (Claude-assisted translation under a fixed terminology sheet), then reviewed by the author — who is not a native English speaker. All worked transcripts are translations of the original Korean sessions, not re-runs performed in English ; the untouched Korean originals live in the Korean repository . Code syntax, identifiers, numbers, and verification values are untouched.
Native-speaker corrections are very welcome — if a sentence reads wrong, please open an Issue or a Pull Request.
If you only read one chapter: 22.2 The Colleague Who Lies with Confidence — Stopping Hallucinations with a Verification Gate — the book's method in one sitting
Preface · How to Read This Book — start here for the reading routes
1.0 Before You Start — Install, Account, Pricing, and the Terminal Survival Kit — follow along from the very first screen
Diagrams are written as ```mermaid code blocks and render natively on GitHub . Click any chapter below and read.
Minsoo Lee (이민수) has been a game designer since 2002 — design team lead on Ragnarok Online at Gravity (2002–2007), and today a Design Director on an MMORPG in production in Korea. This book is the field record of bringing Claude Code into that team's daily design work. · LinkedIn
24 parts + appendices A~N + epilogue · 100 chapters. Appendices live in manuscript/part99-appendix/ , the colophon in manuscript/_colophon.md .
1.0 Before You Start — Install, Account, Pricing, and the Terminal Survival Kit
1.1 A Game Designer's First Encounter with Claude Code
1.2 Model, Token, Harness — The Path a Task's Tokens Travel
1.3 Memory, Permissions, and Settings Infrastructure
Part 2 · Information Architecture
2.1 YAML Frontmatter — Every Document as Data
2.2 Per-Page Atoms — The Anatomy of One Decision per Document
2.3 Layer Design — Abstracting Game Systems
2.4 Ontology and the Wikilink Graph — Verifying the Semantic Arrows
3.1 The Systems Designer's Work and Layer Coordinates
3.2 Schema First — The $schema Sheet Comes Before the Data
3.3 Relation Map Visualization — Seeing Dependencies with Your Own Eyes
3.4 Prompt Patterns for AI-Assisted Systems Design
4.1 The Combat Designer and the Layer Stack — Which Cell Does Game Feel Go In
4.2 Combat Look & Feel — Pinning Game Feel Down in Data
4.3 Combos, Cancels, and Input Queues — Enumerate the Paths and Verify Them
4.4 AI-Assisted Combat Simulation and Verification
5.1 The NarrativeDocs Layer 0–4 Structure
5.2 Worldview → Character → Quest Consistency Verification
5.3 AI-Assisted Narrative Writing
5.4 Dialogue and Voice Consistency
6.1 Procedural Content Generation and AI — The One Cell Where the Two Axes Cross
6.2 city_hunting_generator — 30 Cities in 4 Weeks
6.3 NPC Persona and Squad — From a Mannequin Museum to a Small Society
6.4 Content Production Workflow — Tying Multiple Generators into One Production Line
7.1 Procedural Level Design Master
7.2 The Behavior Tree Editor — A Worked Transcript of a Human and AI Editing and Verifying BT json Together
7.3 The Dungeon and Field Pattern Library
8.1 The Combat Balance Formula — The Seat of the Rulebook Called Determinism
8.2 The Economy Model in Machinations — Catching Inflation with Simulation, Not Meetings
8.3 Damage Simulator — The Day Spec DPS and Sim Output Diverged
8.4 AI-Assisted Balance Simulation
8.5 PvP and Competitive Balance — Win-Rate Matrix, Matchmaking, and Server Authority
9.1 Running the HUD Screenshot Through Lint — Where AI Catches Out-of-Gaze Placement and Failing Contrast
9.2 Skill Button Layout — AI Drafts Three Layouts, lint Rejects Them
9.3 ArtGuide/06_UI Collaboration — Designers Write md; the Art Team Sees Only html
10.1 The Integrity-Check Atom — A Cascade That Guards FKs Across 30 Sheets
10.2 The Decision Verification 3-Layer Sensor — Where Human Review Evidence Lives
10.3 The Alpha Gap Report — Gaps Classified in Natural Language, Prioritized by Humans
Part 11 · Characters, Pets, and Mounts
11.1 Naming Conventions and Skill-to-Art Mapping
11.2 Pet and Mount Systems — From 1 Template to 50 Instances
12.1 The AI Art Asset Pipeline — Mass-Generate in Reversible Stages, Stop Before the Irreversible Gate
12.2 The Seven Areas of ArtGuide (Character, Animation, Monster, NPC, VFX, UI, Environment)
12.3 Design Doc → Concept → In-Game Asset Flow
13.1 Hundreds of Free-Text Responses into Topics — AI Does the Clustering, People Do the Diagnosis
13.2 KPI Definition and Tracking — Humans Define, AI Diagnoses Anomaly Signals
13.3 From Anomalous Metrics to Decisions — AI Proposes Hypotheses, Humans Decide
14.1 From 30 PC HUD Elements to 10 on Mobile — Constraints as a Rulebook, Compression by AI
14.2 Platform Differences (iOS / Android / PC)
14.3 Touch / Mouse Input Design
15.1 Live Ops Overview — AI Combines Event Candidates, the Rulebook Filters Them, and a Human Chooses
15.2 Event and Season Ops — From One Template to Ten Variation Candidates, Only the Review Is Human
15.3 100 Feedback Items into Topics — Clustering Goes to the LLM, Priorities Stay with People
16.1 Combat TF Operations — Only Decisions Leave the Isolated Workspace as Canon
16.2 Collaborating with Other Disciplines — Sorting External Requests into 3 Tracks
16.3 One Decision, Three Packages — Framing Deliverables by Discipline
17.1 Why Meeting Notes Are the Biggest Pain
17.2 An Extraction Pipeline That Mines Decisions from Meeting Notes
17.3 Meeting Categories, Captions, and Sync — The Three Axes That Turn Meeting Notes into Assets
17.4 Turning Meeting Notes into a Decision Database — Five AI Automation Points
Part 18 · Decisions and Impact
18.1 The Decision-Tracking System
18.2 Impact Propagation and Tier Classification
18.3 Pre- and Post-Decision Impact Tracking Workflow — From Pre-Assessment to Post-Verification
18.4 The Document Impact Grep Workflow — Pulling the Impact Scope with impact
Part 19 · Leads and Team Leadership
19.1 Turning the Vision into a Scorecard for Decisions — Running 26 decisions/ Atoms Through an LLM
19.2 Classify Conflicts and Don't Let Meeting Decisions Slip Away — AI Assistance for Meeting Leadership
19.3 AI Adoption Strategy and Executive Buy-In — From Conservative to Progressive, and No Doctored ROI
20.1 One DD Runs Five People's Worth of Collaboration Memory — The team_memory System
20.2 Per-Member Memory — Separating User Compartments and the Shared Compartment
20.3 The Design Portal — Where the Team Comes In Through a Browser
20.4 MCP Project Management — Connecting Collaboration Tools and Documents to the LLM
Part 21 · Chapter 1. The Retrospective as the Starting Point of Everything
Part 21 · Chapter 2. The Retrospective System and Atom Promotion — Turning Discoveries into Permanent Assets
Part 21 · Chapter 3. Closing the Self-Improving Loop
22.1 Prompt Engineering — The Game Designer's One-Page Work Order
22.2 The Colleague Who Lies with Confidence — Stopping Hallucinations with a Verification Gate
22.3 AI Cost Management — Enforcing the Token Budget in Code
22.4 Copyright and Ethics — Closing an Output's Rights, Disclosure, and Agreement in One Procedure
Part 23 · Extensions and What's Next
Part 23 · Chapter 1. The Wrapper, Cascade, and Junction Patterns
Part 23 · Chapter 2. Adopting Hermes Agent
Part 23 · Chapter 3. Tool Curation — Cutting Unused Tools with Data
Part 23 · Chapter 4. The Puzzle Game I Built Alone — A Critter Sort Field Report
24.1 The Verification System — Catching Consistency, Links, and Staleness in Code
24.2 Mermaid Diagram Automation — Letting Documents Draw Their Own Diagrams
24.3 Wikilinks and Document Hierarchy

[truncated]

## Original Extract

Free book: AI Workflow for Game Designers — No Fabricated Numbers. Full text free in this repo (CC BY-NC-SA). English edition of the Korean print original (ISBN 979-11-12-21479-9). - eremes81/game-design-ai-practice-en

GitHub - eremes81/game-design-ai-practice-en: Free book: AI Workflow for Game Designers — No Fabricated Numbers. Full text free in this repo (CC BY-NC-SA). English edition of the Korean print original (ISBN 979-11-12-21479-9). · GitHub
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
eremes81
/
game-design-ai-practice-en
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
21 Commits 21 Commits Folders and files
assets assets manuscript manuscript .gitignore .gitignore .nojekyll .nojekyll LICENSE LICENSE README.md README.md index.html index.html robots.txt robots.txt sitemap.xml sitemap.xml View all files Repository files navigation
AI Workflow for Game Designers
No Fabricated Numbers — A Six-Month Field Manual for Claude Code, Prompts, Validation, and Production Memory
🌐 Editions: 한국어 — Original · English · 日本語 · ไทย · Bahasa Indonesia · 简体中文 · 繁體中文
🧰 Companion tooling — Harness Starter Kit
A minimal, runnable version of the memory, hook and retrospective structure this book builds in 1.3, Part 21 and Part 24. Download the folder, open it in an agent, and it works.
Built as the appendix to the author's other book (for general office work), but the skeleton is the same one described here. MIT licensed.
A hands-on field manual by Minsoo Lee — a game designer of 24 years, former design team lead on Ragnarok Online (2002–2007), now a Design Director shipping an MMORPG in Korea — on bringing generative AI (Claude Code) into daily production work . Not theory or forecasts — it walks one task at a time from the very first screen (install, accounts, pricing) through systems design, combat, narrative, level design, balance, UX, and live ops, all the way to turning meeting notes into decisions, validation gates, cost management, and copyright.
The subtitle — "No Fabricated Numbers" — is this book's promise. Every figure and case in the text comes from real work, not invented examples, and most of the code runs as is on the Python standard library alone.
Schema-first data design — the table schema and legal value ranges come before the AI writes a single row
Validation gates that catch AI-invented numbers, IDs, and dangling references before they reach the data sheets
Meeting notes → tracked decisions — decision, owner, and rationale extracted and greppable
AI cost management in code — token budgets enforced by scripts, not by good intentions
Balance simulation, procedural content pipelines, UX lint, live-ops workflows, and a team adoption strategy for leads
The book is open to readers outside games, too. Workflows like turning meeting notes into decisions, tracing a decision's ripple effects, and guarding quality with validation gates work regardless of your job. The "Beyond Games" box in each chapter is the bridge, and every chapter also carries a "Solo Scale-Down" for people building alone, with no team.
This is the English edition of the Korean original, 게임 기획 실무에서 바로 쓰는 AI·클로드 코드 활용법 (BOOKK, 2026 · ISBN 979-11-12-21479-9).
In the spirit of this book's honesty-first principle, here is exactly how this edition was made: it was translated from the Korean original using the book's own AI workflow (Claude-assisted translation under a fixed terminology sheet), then reviewed by the author — who is not a native English speaker. All worked transcripts are translations of the original Korean sessions, not re-runs performed in English ; the untouched Korean originals live in the Korean repository . Code syntax, identifiers, numbers, and verification values are untouched.
Native-speaker corrections are very welcome — if a sentence reads wrong, please open an Issue or a Pull Request.
If you only read one chapter: 22.2 The Colleague Who Lies with Confidence — Stopping Hallucinations with a Verification Gate — the book's method in one sitting
Preface · How to Read This Book — start here for the reading routes
1.0 Before You Start — Install, Account, Pricing, and the Terminal Survival Kit — follow along from the very first screen
Diagrams are written as ```mermaid code blocks and render natively on GitHub . Click any chapter below and read.
Minsoo Lee (이민수) has been a game designer since 2002 — design team lead on Ragnarok Online at Gravity (2002–2007), and today a Design Director on an MMORPG in production in Korea. This book is the field record of bringing Claude Code into that team's daily design work. · LinkedIn
24 parts + appendices A~N + epilogue · 100 chapters. Appendices live in manuscript/part99-appendix/ , the colophon in manuscript/_colophon.md .
1.0 Before You Start — Install, Account, Pricing, and the Terminal Survival Kit
1.1 A Game Designer's First Encounter with Claude Code
1.2 Model, Token, Harness — The Path a Task's Tokens Travel
1.3 Memory, Permissions, and Settings Infrastructure
Part 2 · Information Architecture
2.1 YAML Frontmatter — Every Document as Data
2.2 Per-Page Atoms — The Anatomy of One Decision per Document
2.3 Layer Design — Abstracting Game Systems
2.4 Ontology and the Wikilink Graph — Verifying the Semantic Arrows
3.1 The Systems Designer's Work and Layer Coordinates
3.2 Schema First — The $schema Sheet Comes Before the Data
3.3 Relation Map Visualization — Seeing Dependencies with Your Own Eyes
3.4 Prompt Patterns for AI-Assisted Systems Design
4.1 The Combat Designer and the Layer Stack — Which Cell Does Game Feel Go In
4.2 Combat Look & Feel — Pinning Game Feel Down in Data
4.3 Combos, Cancels, and Input Queues — Enumerate the Paths and Verify Them
4.4 AI-Assisted Combat Simulation and Verification
5.1 The NarrativeDocs Layer 0–4 Structure
5.2 Worldview → Character → Quest Consistency Verification
5.3 AI-Assisted Narrative Writing
5.4 Dialogue and Voice Consistency
6.1 Procedural Content Generation and AI — The One Cell Where the Two Axes Cross
6.2 city_hunting_generator — 30 Cities in 4 Weeks
6.3 NPC Persona and Squad — From a Mannequin Museum to a Small Society
6.4 Content Production Workflow — Tying Multiple Generators into One Production Line
7.1 Procedural Level Design Master
7.2 The Behavior Tree Editor — A Worked Transcript of a Human and AI Editing and Verifying BT json Together
7.3 The Dungeon and Field Pattern Library
8.1 The Combat Balance Formula — The Seat of the Rulebook Called Determinism
8.2 The Economy Model in Machinations — Catching Inflation with Simulation, Not Meetings
8.3 Damage Simulator — The Day Spec DPS and Sim Output Diverged
8.4 AI-Assisted Balance Simulation
8.5 PvP and Competitive Balance — Win-Rate Matrix, Matchmaking, and Server Authority
9.1 Running the HUD Screenshot Through Lint — Where AI Catches Out-of-Gaze Placement and Failing Contrast
9.2 Skill Button Layout — AI Drafts Three Layouts, lint Rejects Them
9.3 ArtGuide/06_UI Collaboration — Designers Write md; the Art Team Sees Only html
10.1 The Integrity-Check Atom — A Cascade That Guards FKs Across 30 Sheets
10.2 The Decision Verification 3-Layer Sensor — Where Human Review Evidence Lives
10.3 The Alpha Gap Report — Gaps Classified in Natural Language, Prioritized by Humans
Part 11 · Characters, Pets, and Mounts
11.1 Naming Conventions and Skill-to-Art Mapping
11.2 Pet and Mount Systems — From 1 Template to 50 Instances
12.1 The AI Art Asset Pipeline — Mass-Generate in Reversible Stages, Stop Before the Irreversible Gate
12.2 The Seven Areas of ArtGuide (Character, Animation, Monster, NPC, VFX, UI, Environment)
12.3 Design Doc → Concept → In-Game Asset Flow
13.1 Hundreds of Free-Text Responses into Topics — AI Does the Clustering, People Do the Diagnosis
13.2 KPI Definition and Tracking — Humans Define, AI Diagnoses Anomaly Signals
13.3 From Anomalous Metrics to Decisions — AI Proposes Hypotheses, Humans Decide
14.1 From 30 PC HUD Elements to 10 on Mobile — Constraints as a Rulebook, Compression by AI
14.2 Platform Differences (iOS / Android / PC)
14.3 Touch / Mouse Input Design
15.1 Live Ops Overview — AI Combines Event Candidates, the Rulebook Filters Them, and a Human Chooses
15.2 Event and Season Ops — From One Template to Ten Variation Candidates, Only the Review Is Human
15.3 100 Feedback Items into Topics — Clustering Goes to the LLM, Priorities Stay with People
16.1 Combat TF Operations — Only Decisions Leave the Isolated Workspace as Canon
16.2 Collaborating with Other Disciplines — Sorting External Requests into 3 Tracks
16.3 One Decision, Three Packages — Framing Deliverables by Discipline
17.1 Why Meeting Notes Are the Biggest Pain
17.2 An Extraction Pipeline That Mines Decisions from Meeting Notes
17.3 Meeting Categories, Captions, and Sync — The Three Axes That Turn Meeting Notes into Assets
17.4 Turning Meeting Notes into a Decision Database — Five AI Automation Points
Part 18 · Decisions and Impact
18.1 The Decision-Tracking System
18.2 Impact Propagation and Tier Classification
18.3 Pre- and Post-Decision Impact Tracking Workflow — From Pre-Assessment to Post-Verification
18.4 The Document Impact Grep Workflow — Pulling the Impact Scope with impact
Part 19 · Leads and Team Leadership
19.1 Turning the Vision into a Scorecard for Decisions — Running 26 decisions/ Atoms Through an LLM
19.2 Classify Conflicts and Don't Let Meeting Decisions Slip Away — AI Assistance for Meeting Leadership
19.3 AI Adoption Strategy and Executive Buy-In — From Conservative to Progressive, and No Doctored ROI
20.1 One DD Runs Five People's Worth of Collaboration Memory — The team_memory System
20.2 Per-Member Memory — Separating User Compartments and the Shared Compartment
20.3 The Design Portal — Where the Team Comes In Through a Browser
20.4 MCP Project Management — Connecting Collaboration Tools and Documents to the LLM
Part 21 · Chapter 1. The Retrospective as the Starting Point of Everything
Part 21 · Chapter 2. The Retrospective System and Atom Promotion — Turning Discoveries into Permanent Assets
Part 21 · Chapter 3. Closing the Self-Improving Loop
22.1 Prompt Engineering — The Game Designer's One-Page Work Order
22.2 The Colleague Who Lies with Confidence — Stopping Hallucinations with a Verification Gate
22.3 AI Cost Management — Enforcing the Token Budget in Code
22.4 Copyright and Ethics — Closing an Output's Rights, Disclosure, and Agreement in One Procedure
Part 23 · Extensions and What's Next
Part 23 · Chapter 1. The Wrapper, Cascade, and Junction Patterns
Part 23 · Chapter 2. Adopting Hermes Agent
Part 23 · Chapter 3. Tool Curation — Cutting Unused Tools with Data
Part 23 · Chapter 4. The Puzzle Game I Built Alone — A Critter Sort Field Report
24.1 The Verification System — Catching Consistency, Links, and Staleness in Code
24.2 Mermaid Diagram Automation — Letting Documents Draw Their Own Diagrams
24.3 Wikilinks and Document Hierarchy

[truncated]
