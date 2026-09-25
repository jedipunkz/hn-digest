---
source: "https://github.com/Atri10/executor"
hn_url: "https://news.ycombinator.com/item?id=49847953"
title: "Show HN: Stop shipping AI slop, Executor makes agents prove it works"
article_title: "GitHub - Atri10/executor: Stop shipping AI slop, Executor makes agents prove it works · GitHub"
image: "https://opengraph.githubassets.com/0c6a254f3d4aee7926e74813f49169fbbd4519cd0a6bec4a3316ff29236f3838/Atri10/executor"
author: "LambdaLogic"
captured_at: "2026-09-25T18:15:21Z"
capture_tool: "hn-digest"
hn_id: 49847953
score: 1
comments: 0
posted_at: "2026-09-25T18:05:31Z"
tags:
  - hacker-news
---

# Show HN: Stop shipping AI slop, Executor makes agents prove it works

- HN: [49847953](https://news.ycombinator.com/item?id=49847953)
- Source: [github.com](https://github.com/Atri10/executor)
- Score: 1
- Comments: 0
- Posted: 2026-09-25T18:05:31Z

## Translation

Title: Show HN: Stop shipping AI slop, Executor makes agents prove it works
Article title: GitHub - Atri10/executor: Stop shipping AI slop, Executor makes agents prove it works · GitHub
Description: Stop shipping AI slop, Executor makes agents prove it works - Atri10/executor

Article text:
GitHub - Atri10/executor: Stop shipping AI slop, Executor makes agents prove it works · GitHub
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
Atri10
/
executor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
52 Commits 52 Commits Folders and files
.github .github scripts scripts skills skills .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
An initiative-scoped workflow system for coding agents: it takes a major
idea from intake through architecture, spec, plan, execution, review, and
verification — with a strict per-initiative ID namespace, separated thinking
and execution stores, a script-enforced contract at every state transition,
and evidence-backed completion.
Nothing floats. Every document, task, review, verdict, and ruling
carries an ID that names the initiative it belongs to. A body of work gets
one Initiative ; the initiative owns a folder, an ID namespace, and every
document produced about it.
Nothing is enforced by hope. Where a workflow could silently drift — a
run row lying about the ledger, a task completed without review, a plan
naming paths only scripts may resolve — a script checks it and exits
non-zero. Contracts enforced by prose are hopes; contracts enforced by
scripts are contracts.
Works with any agent that can read skill files from a directory —
Claude Code, Codex, omp , or anything
similar. The skills are markdown; the scripts are POSIX bash.
Clone the repo and copy the skills/ directories into whatever directory
your agent loads skills from:
git clone git@github.com:Atri10/executor.git
cp -R executor/skills/ * < your-agents-skills-dir > /
Pin to a release tag instead of main for stability:
git clone --branch v0.3.0 git@github.com:Atri10/executor.git
One-paste install for any LLM agent
Paste this into your agent — Claude Code, Cursor, Aider, Codex, omp, or any
other harness. It discovers the right skills directory itself and verifies
the install:
Install The Executor skill library for me:
1. Clone https://github.com/Atri10/executor.git into a temp directory
(use --branch v0.3.0 for the latest release, or default branch for main).
2. Find my agent's skills directory. Candidates, in order — use the first
that exists, or ask me if none do:
- ~/.omp/agent/skills/ (omp)
- ~/.claude/skills/ (Claude Code)
- ~/.cursor/skills/ (Cursor)
- .claude/skills/ (repo-local Claude Code)
- ~/.aider/skills/ or as my harness documents
3. Copy every directory from the clone's skills/ folder into that skills
directory (each is one skill: skills/executor, skills/executor-spec, ...).
4. Verify: run bash <skills-dir>/executor/scripts/exec-run with no arguments
— it must print a usage line and exit non-zero. Then confirm the ten
SKILL.md files exist under the skills directory.
5. Tell me which directory you installed into, and how to invoke the
router in my harness (usually /skill:executor or just asking for
"the executor").
Do not modify any file inside the clone or the skills directory other
than the copy operation itself.
Then invoke the root router:
/skill:executor
or just say "start an initiative" — normal requests route by each skill's
frontmatter description.
Skill
Phase
Output
executor
Router + contract
Loads the right phase skill, defines the ID namespace
executor-initiative
Intake
Initiative folder, charter, registry entry, initiative branch
executor-discovery
Discovery
Research notes, options comparison
executor-architecture
Architecture, Design
Architecture, ADRs, interfaces, component designs
executor-spec
Specification
Spec, risks, verification strategy (one row per requirement)
executor-planning
Planning
Plans with tasks, linted before the gate
executor-execution
Execution
Task dispatch, ledger, reports, run registry
executor-review
Review
Per-task and whole-branch verdicts, findings, fix loops
executor-verification
Verification
Evidence-backed proof each requirement holds
executor-handoff
Handoff
Human decision menu: merge, PR, or keep the branch
Phases compress, they never vanish. A small initiative can produce a
charter and a spec in one exchange and skip discovery — but skipping is a
stated decision recorded in the charter, not an omission.
Every task gets a brief, a context file, and a fresh reviewer
exec-brief extracts one task's text into a self-contained brief — the
implementer reads requirements in one call, and task text never passes
through the controller's context.
exec-context assembles everything the brief cannot know: the exact
signatures earlier tasks provide, the current surface of the files being
modified, the binding global constraints, and the rulings that touch the
task's files. Implementers start working without exploring.
Two-verdict reviews (spec compliance + code quality) from a reviewer
who never trusted the implementer's report, writing a verdict file —
not a chat message that vanishes on the next summarization.
Non-code tasks still get reviewed : a docs-only or evidence-capture
task is judged on its report vs its brief, with the same mandatory
verdict file.
Findings are severity-graded with worked calibration examples, fixed in
rounds (1–3 resume the original implementer; 4–5 escalate to a fresh,
more-capable model), re-reviewed scoped to the fix diff, and at the cap
adjudicated by recorded ruling — never silently dropped. Re-reviews check
the fix addressed the root cause , and whether any test was weakened.
Evidence is capability-aware, not ceremony
The implementer contract asks for the strongest feasible evidence
for every behavior change: a watched failing test where a test harness
exists, a named alternative instrument where it does not (CLI fixture
run, parse/render check, exercised UI), and an explicit NOT-RUN/UNAVAILABLE
record where nothing feasible exists. Reviewers verify evidence, and a
reviewer who cannot name the failure a demanded test would catch does
not get to demand it.
Script-enforced state, everywhere
Script
Owns
exec-initiative
Allocate initiative IDs, scaffold folders, phase log, initiative branch ( branch INIT-0004 )
exec-id
Next free ID of any type — allocation never guesses
exec-plan-lint
Planning gate : rejects literal store paths in plans, task headings without IDs, missing or empty spec / interfaces / tasks / execution_mode , task-count mismatch, and over-specified task bodies (impl-language fences >40 lines or >60% of a body)
exec-workspace
Resolve and seed a plan's execution workspace: ledger, rulings, preflight scan, dispatch log
exec-brief / exec-context
Task brief and context files, generated, never hand-built — briefs carry contract verbatim and mark embedded code as advisory
exec-review-package
Review diffs with commit list + stat + -U10 diff in one file, per round
exec-fix-package
Fix-round dispatch package: verdict findings + implementer report + brief + context verbatim, so fix agents see the contract, not a paraphrase
exec-run
Run lifecycle in the registry: start / task / complete / pause / blocked / check
exec-run check
Drift + semantic audit : registry row vs ledger, verdict CONTENT (a FAIL verdict blocks), exact task set with latest-state reduction, final verdict lineage (a failing final re-review supersedes an earlier clean one), completed tasks present in the Task status table — exit 1 names the failure
exec-branch
Plan-branch lifecycle: fork from the initiative branch, merge refused unless the review audit passes
exec-evidence
Per-criterion evidence files in the initiative's tracked verification/evidence/PNN/ , immutable per round ( -attempt2 on same-round reruns), atomic publish, per-round state stamp (branch, commit, dirtiness)
exec-store-check
Thinking-store integrity gate : registry ↔ folders ↔ Documents table ↔ frontmatter statuses ↔ cross-links ↔ evidence citations ↔ phase-log chronology, plus per-kind document contracts (architecture must carry a Mermaid diagram, specs need requirement headings + a resolving verification link, acti
[truncated]
One branch per initiative ( initiative/INIT-NNNN , forked from wherever the
human currently is, fork point recorded), one branch per plan
( plan/INIT-NNNN-Pnn , forked from the initiative branch). Task commits land
with detailed messages on the plan branch; the plan branch merges back with
--no-ff — only after its final review verdict exists and the audit
passes. Merging the initiative branch onward is always the human's
explicit decision at handoff.
flowchart LR
BASE["base branch"] --> INIT["initiative/INIT-0004"]
INIT --> P1["plan/INIT-0004-P01"]
INIT --> P2["plan/INIT-0004-P02"]
P1 -->|"merge: review-gated"| INIT
P2 -->|"merge: review-gated"| INIT
INIT --> HUMAN["human decides at handoff"]
Loading
Reviews that survive the session
Every review is a round with an ID ( INIT-0004-P01-T03-R02 ), a diff
file, and a verdict file carrying YAML frontmatter. Findings are labelled
( C1 , I2 , M1 ), cited to the spec requirement they violate
( INIT-0004-SPEC-01-R07 ), and live in files a fixer reads directly — the
controller transcribes nothing. The final whole-branch review walks every
declared cross-task seam and triages every deferred or parked finding.
Re-reviews do two jobs: impact review of the fix (following what it
actually affects — including unchanged callers) before finding closure,
so a regression the fix introduced in untouched code is still caught,
and a fix-only lens never hides it.
Every generated artifact carries frontmatter
Briefs, contexts, ledger, rulings, preflight scan, dispatch log, reports,
verdicts, evidence files — all carry the same YAML identity block
( kind , id , initiative , plan , created_at , …), defined in the
frontmatter contract . An agent
reading any file cold knows exactly what it is holding.
Verification converts claims into evidence
The spec's verification strategy names one criterion per requirement with
its exact command. The verification phase runs each row fresh against the
current commit and reports four honest statuses: PROVEN , FAILED ,
NOT-RUN , UNAVAILABLE . A single NOT-RUN blocks the word "complete" —
and nothing upgrades it by inference. Raw observed output lands in
per-criterion evidence files the outcomes table cites.
After a context loss or model switch, the controller reads the ledger —
not its recollection. Completed tasks are not re-dispatched; the ledger's
identity block refuses a ledger that belongs to another plan; live subagent
identities are recorded so a fix round can resume rather than replace.
Nothing in either store is ever deleted by a skill — pruning is a human
decision.
flowchart LR
s

[truncated]

## Original Extract

Stop shipping AI slop, Executor makes agents prove it works - Atri10/executor

GitHub - Atri10/executor: Stop shipping AI slop, Executor makes agents prove it works · GitHub
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
Atri10
/
executor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
52 Commits 52 Commits Folders and files
.github .github scripts scripts skills skills .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
An initiative-scoped workflow system for coding agents: it takes a major
idea from intake through architecture, spec, plan, execution, review, and
verification — with a strict per-initiative ID namespace, separated thinking
and execution stores, a script-enforced contract at every state transition,
and evidence-backed completion.
Nothing floats. Every document, task, review, verdict, and ruling
carries an ID that names the initiative it belongs to. A body of work gets
one Initiative ; the initiative owns a folder, an ID namespace, and every
document produced about it.
Nothing is enforced by hope. Where a workflow could silently drift — a
run row lying about the ledger, a task completed without review, a plan
naming paths only scripts may resolve — a script checks it and exits
non-zero. Contracts enforced by prose are hopes; contracts enforced by
scripts are contracts.
Works with any agent that can read skill files from a directory —
Claude Code, Codex, omp , or anything
similar. The skills are markdown; the scripts are POSIX bash.
Clone the repo and copy the skills/ directories into whatever directory
your agent loads skills from:
git clone git@github.com:Atri10/executor.git
cp -R executor/skills/ * < your-agents-skills-dir > /
Pin to a release tag instead of main for stability:
git clone --branch v0.3.0 git@github.com:Atri10/executor.git
One-paste install for any LLM agent
Paste this into your agent — Claude Code, Cursor, Aider, Codex, omp, or any
other harness. It discovers the right skills directory itself and verifies
the install:
Install The Executor skill library for me:
1. Clone https://github.com/Atri10/executor.git into a temp directory
(use --branch v0.3.0 for the latest release, or default branch for main).
2. Find my agent's skills directory. Candidates, in order — use the first
that exists, or ask me if none do:
- ~/.omp/agent/skills/ (omp)
- ~/.claude/skills/ (Claude Code)
- ~/.cursor/skills/ (Cursor)
- .claude/skills/ (repo-local Claude Code)
- ~/.aider/skills/ or as my harness documents
3. Copy every directory from the clone's skills/ folder into that skills
directory (each is one skill: skills/executor, skills/executor-spec, ...).
4. Verify: run bash <skills-dir>/executor/scripts/exec-run with no arguments
— it must print a usage line and exit non-zero. Then confirm the ten
SKILL.md files exist under the skills directory.
5. Tell me which directory you installed into, and how to invoke the
router in my harness (usually /skill:executor or just asking for
"the executor").
Do not modify any file inside the clone or the skills directory other
than the copy operation itself.
Then invoke the root router:
/skill:executor
or just say "start an initiative" — normal requests route by each skill's
frontmatter description.
Skill
Phase
Output
executor
Router + contract
Loads the right phase skill, defines the ID namespace
executor-initiative
Intake
Initiative folder, charter, registry entry, initiative branch
executor-discovery
Discovery
Research notes, options comparison
executor-architecture
Architecture, Design
Architecture, ADRs, interfaces, component designs
executor-spec
Specification
Spec, risks, verification strategy (one row per requirement)
executor-planning
Planning
Plans with tasks, linted before the gate
executor-execution
Execution
Task dispatch, ledger, reports, run registry
executor-review
Review
Per-task and whole-branch verdicts, findings, fix loops
executor-verification
Verification
Evidence-backed proof each requirement holds
executor-handoff
Handoff
Human decision menu: merge, PR, or keep the branch
Phases compress, they never vanish. A small initiative can produce a
charter and a spec in one exchange and skip discovery — but skipping is a
stated decision recorded in the charter, not an omission.
Every task gets a brief, a context file, and a fresh reviewer
exec-brief extracts one task's text into a self-contained brief — the
implementer reads requirements in one call, and task text never passes
through the controller's context.
exec-context assembles everything the brief cannot know: the exact
signatures earlier tasks provide, the current surface of the files being
modified, the binding global constraints, and the rulings that touch the
task's files. Implementers start working without exploring.
Two-verdict reviews (spec compliance + code quality) from a reviewer
who never trusted the implementer's report, writing a verdict file —
not a chat message that vanishes on the next summarization.
Non-code tasks still get reviewed : a docs-only or evidence-capture
task is judged on its report vs its brief, with the same mandatory
verdict file.
Findings are severity-graded with worked calibration examples, fixed in
rounds (1–3 resume the original implementer; 4–5 escalate to a fresh,
more-capable model), re-reviewed scoped to the fix diff, and at the cap
adjudicated by recorded ruling — never silently dropped. Re-reviews check
the fix addressed the root cause , and whether any test was weakened.
Evidence is capability-aware, not ceremony
The implementer contract asks for the strongest feasible evidence
for every behavior change: a watched failing test where a test harness
exists, a named alternative instrument where it does not (CLI fixture
run, parse/render check, exercised UI), and an explicit NOT-RUN/UNAVAILABLE
record where nothing feasible exists. Reviewers verify evidence, and a
reviewer who cannot name the failure a demanded test would catch does
not get to demand it.
Script-enforced state, everywhere
Script
Owns
exec-initiative
Allocate initiative IDs, scaffold folders, phase log, initiative branch ( branch INIT-0004 )
exec-id
Next free ID of any type — allocation never guesses
exec-plan-lint
Planning gate : rejects literal store paths in plans, task headings without IDs, missing or empty spec / interfaces / tasks / execution_mode , task-count mismatch, and over-specified task bodies (impl-language fences >40 lines or >60% of a body)
exec-workspace
Resolve and seed a plan's execution workspace: ledger, rulings, preflight scan, dispatch log
exec-brief / exec-context
Task brief and context files, generated, never hand-built — briefs carry contract verbatim and mark embedded code as advisory
exec-review-package
Review diffs with commit list + stat + -U10 diff in one file, per round
exec-fix-package
Fix-round dispatch package: verdict findings + implementer report + brief + context verbatim, so fix agents see the contract, not a paraphrase
exec-run
Run lifecycle in the registry: start / task / complete / pause / blocked / check
exec-run check
Drift + semantic audit : registry row vs ledger, verdict CONTENT (a FAIL verdict blocks), exact task set with latest-state reduction, final verdict lineage (a failing final re-review supersedes an earlier clean one), completed tasks present in the Task status table — exit 1 names the failure
exec-branch
Plan-branch lifecycle: fork from the initiative branch, merge refused unless the review audit passes
exec-evidence
Per-criterion evidence files in the initiative's tracked verification/evidence/PNN/ , immutable per round ( -attempt2 on same-round reruns), atomic publish, per-round state stamp (branch, commit, dirtiness)
exec-store-check
Thinking-store integrity gate : registry ↔ folders ↔ Documents table ↔ frontmatter statuses ↔ cross-links ↔ evidence citations ↔ phase-log chronology, plus per-kind document contracts (architecture must carry a Mermaid diagram, specs need requirement headings + a resolving verification link, acti
[truncated]
One branch per initiative ( initiative/INIT-NNNN , forked from wherever the
human currently is, fork point recorded), one branch per plan
( plan/INIT-NNNN-Pnn , forked from the initiative branch). Task commits land
with detailed messages on the plan branch; the plan branch merges back with
--no-ff — only after its final review verdict exists and the audit
passes. Merging the initiative branch onward is always the human's
explicit decision at handoff.
flowchart LR
BASE["base branch"] --> INIT["initiative/INIT-0004"]
INIT --> P1["plan/INIT-0004-P01"]
INIT --> P2["plan/INIT-0004-P02"]
P1 -->|"merge: review-gated"| INIT
P2 -->|"merge: review-gated"| INIT
INIT --> HUMAN["human decides at handoff"]
Loading
Reviews that survive the session
Every review is a round with an ID ( INIT-0004-P01-T03-R02 ), a diff
file, and a verdict file carrying YAML frontmatter. Findings are labelled
( C1 , I2 , M1 ), cited to the spec requirement they violate
( INIT-0004-SPEC-01-R07 ), and live in files a fixer reads directly — the
controller transcribes nothing. The final whole-branch review walks every
declared cross-task seam and triages every deferred or parked finding.
Re-reviews do two jobs: impact review of the fix (following what it
actually affects — including unchanged callers) before finding closure,
so a regression the fix introduced in untouched code is still caught,
and a fix-only lens never hides it.
Every generated artifact carries frontmatter
Briefs, contexts, ledger, rulings, preflight scan, dispatch log, reports,
verdicts, evidence files — all carry the same YAML identity block
( kind , id , initiative , plan , created_at , …), defined in the
frontmatter contract . An agent
reading any file cold knows exactly what it is holding.
Verification converts claims into evidence
The spec's verification strategy names one criterion per requirement with
its exact command. The verification phase runs each row fresh against the
current commit and reports four honest statuses: PROVEN , FAILED ,
NOT-RUN , UNAVAILABLE . A single NOT-RUN blocks the word "complete" —
and nothing upgrades it by inference. Raw observed output lands in
per-criterion evidence files the outcomes table cites.
After a context loss or model switch, the controller reads the ledger —
not its recollection. Completed tasks are not re-dispatched; the ledger's
identity block refuses a ledger that belongs to another plan; live subagent
identities are recorded so a fix round can resume rather than replace.
Nothing in either store is ever deleted by a skill — pruning is a human
decision.
flowchart LR
s

[truncated]
