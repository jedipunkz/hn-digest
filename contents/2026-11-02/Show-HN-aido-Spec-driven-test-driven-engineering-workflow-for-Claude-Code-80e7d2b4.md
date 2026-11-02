---
source: "https://github.com/joebew42/aido"
hn_url: "https://news.ycombinator.com/item?id=50340386"
title: "Show HN: aido – Spec-driven, test-driven engineering workflow for Claude Code"
article_title: "GitHub - joebew42/aido: AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack. · GitHub"
image: "https://opengraph.githubassets.com/793df47c72d89016bb0a6bc0f0cc6e67c6fee422c4b940ee352a05ebf848c2ec/joebew42/aido"
author: "joebew42"
captured_at: "2026-11-02T17:53:20Z"
capture_tool: "hn-digest"
hn_id: 50340386
score: 1
comments: 0
posted_at: "2026-11-02T17:34:10Z"
tags:
  - hacker-news
---

# Show HN: aido – Spec-driven, test-driven engineering workflow for Claude Code

- HN: [50340386](https://news.ycombinator.com/item?id=50340386)
- Source: [github.com](https://github.com/joebew42/aido)
- Score: 1
- Comments: 0
- Posted: 2026-11-02T17:34:10Z

## Translation

Title: Show HN: aido – Spec-driven, test-driven engineering workflow for Claude Code
Article title: GitHub - joebew42/aido: AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack. · GitHub
Description: AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack. - joebew42/aido

Article text:
GitHub - joebew42/aido: AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack. · GitHub
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
joebew42
aido Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
19 Commits 19 Commits Folders and files
.claude-plugin .claude-plugin docs docs plugins/ aido plugins/ aido tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Hand Claude Code the whole feature: the spec, the tests, the code, the commits and the pull request.
aido stops at every decision that matters and asks you, so what gets built stays your call. It works
with any stack.
/plugin marketplace add joebew42/aido
/plugin install aido@aido
/aido:1-plan
Every skill is available as /aido:<skill> .
Step
Skill
What it does
1
/aido:1-plan
Specifies and plans the feature in SPEC.md
2
/aido:2-implement
Builds the feature
3
/aido:3-organize-commits
Cleans up the commit history
4
/aido:4-create-pr
Creates the pull request
5
/aido:5-handle-review
Handles the review feedback
/aido:1-plan drafts the specification in SPEC.md.draft (from a GitHub issue, one you write, or a
conversation with you), checks it against your repository and refines it with you until you approve
it as SPEC.md , plans the work, creates the branch, and offers to move on to implementation. When
the change looks small and clear, it proposes a lighter route: a brief specification and plan that
you approve once, without the enabling refactoring analysis. /aido:3-organize-commits asks whether
to revert SPEC.md (the default) or keep it in the repository's docs. /aido:4-create-pr offers to
verify the implementation against the specification before pushing.
The guidelines ask for commits and PRs without AI attribution and for no auto-memory. A plugin
can't set these, so add them to your ~/.claude/settings.json :
{
"autoMemoryEnabled" : false ,
"attribution" : { "commit" : " " , "pr" : " " }
}
The plugin's hooks also block commit messages with Co-Authored-By trailers and writes to Claude's
memory directory.
Why aido: AI-Augmented Engineering
An AI-augmented engineer uses an AI agent to do more of the work, without handing over the
judgement. The agent explores the code, drafts the specification, writes the tests and the code,
and shapes the commits and the pull request. The engineer decides what gets built, approves every
step that matters, and stays accountable for the result.
aido stands for "AI do" and echoes sudo : the agent does the work, on your say-so. It is a
workflow for working that way, and each of its practices keeps that split in place:
The engineer decides, the agent proposes. Each stage ends with an explicit approval — the
specification, the plan, the commit structure, the pull request — and the agent asks instead of
guessing whenever a choice is yours.
Specification before code. Every feature starts as a draft that is refined with you into
SPEC.md , then planned. A small, clear change can take a lighter route, with the specification
and the plan approved together. Implementation only starts from an approved plan.
Small, reviewable steps. Code grows test by test with test-driven development, so each change
is easy to follow and to stop.
Good engineering practices by default. Clean architecture, deliberate refactoring, focused
tests and a clean commit history are part of every step, not an afterthought.
Your repository's rules come first. aido brings practices, not conventions. How to
test, lint, name or structure code comes from your repository — first from its CLAUDE.md and
AGENTS.md , then from the code itself.
To learn more about the wider field, see
AI-Augmented Software Engineering at
Carnegie Mellon's Software Engineering Institute.
The specification is at the center
Every stage works from one file, SPEC.md . Planning writes the specification, the plan and the
design notes into it. Implementation reads the plan and records its progress there. The pull
request description is written from it, and review feedback is checked against it. /aido:verify
checks the code against it.
Every specification has the same shape. Context says why the change is needed, Intent what
it must achieve, Out of scope what it leaves out, and Acceptance criteria — numbered AC1 ,
AC2 , … — how you will know it is done. The specification says what and why; how belongs to the
plan that planning adds below it, followed by the progress that implementation records.
Acceptance criteria are checked at every stage
The acceptance criteria are what verifies that the change does what the specification asks, not
just what the plan lists:
planning proposes them when the draft has none, and every plan item names the criteria it
serves — none is left without one
implementation writes each criterion's acceptance test first and keeps it red as the goal
post while the code behind it is built test-driven. Before it calls the feature done, it checks
every criterion: a test that covers it, or — for criteria only a person can judge, marked
(manual) — your confirmation. When one is not met, you choose to implement it, change it, or
defer it
verification checks them again from a fresh context, and any deferred criterion is named in
the pull request description
The specification is checked against your repository
A specification can be clear and still conflict with the repository it is meant for. /aido:1-plan
checks for that before you approve it, and /aido:2-implement checks again at the start of every
session, in case the default branch has moved. When nothing relevant changed, that costs one line.
A misalignment is significant when the specification contradicts your CLAUDE.md or AGENTS.md ,
crosses an architectural boundary, or changes behaviour other code relies on. The agent then explains
the conflict, cites the files that show it, and recommends revising or rejecting the specification.
You can still proceed, since intentional architectural changes are sometimes the point, and the
decision and your reason are recorded in the specification.
Verifying that you built what you agreed
/aido:verify asks one question of the current branch or any PR: does the implementation satisfy
what was agreed? /aido:4-create-pr offers to run it before pushing, and you can run it at any time.
It always runs from a fresh context, so the agent that built the change is not the one judging it,
and it changes nothing.
Its report sorts every requirement as verified , insufficient evidence or unmet , and lists
deviations from what was agreed separately. A requirement is verified only by a test that passed
during the verification, or by your recorded confirmation; claims in the specification's progress
notes do not count. On a PR with no specification, it uses the description and linked issues, and
says that this is weaker.
It complements Claude Code's /code-review rather than replacing it. /aido:verify asks whether you
built what you agreed to build, and /code-review asks whether the code has defects.
The agent does the work, but it stops before every decision that matters: it does not approve the
specification, choose a refactoring that reaches into another module, rewrite history or push
without you. The only way past these checkpoints is to deliberately run the workflow fully
unsupervised. The diagram at the top of this page shows where it stops at each stage.
For example, when you run /aido:1-plan , it stops to ask where the specification comes from: a
GitHub issue it turns into a draft, a draft you write yourself, or a conversation in which it
interviews you and writes the draft. The screenshot at the top of this page shows that question.
Skill
What it does
/aido:1-plan
draft, refine and plan a feature into SPEC.md
/aido:2-implement
implement the feature with TDD
/aido:3-organize-commits
organize the commit history
/aido:4-create-pr
create a pull request
/aido:5-handle-review
handle PR review feedback
Multi-phase feature
Skill
What it does
/aido:multi-phase-spec-pr
create the specification tracking PR
/aido:multi-phase-switch
switch to the main plan branch
/aido:multi-phase-sync
sync the main plan with progress
Utilities
Skill
What it does
/aido:resolve-migration-conflict
resolve database migration conflicts after a rebase, with any migration tool
/aido:verify
check the current branch or a PR against what was agreed, from a fresh context
/aido:stabilize
take over a PR, remote branch, or local changes made outside this workflow
/aido:map-architecture
build or update a slim index of the repo's domain areas and their building blocks
/aido:assess-architecture
assess the architecture area by area and open GitHub issues for the refactorings you pick
Other ways to use it
Write the overall feature specification in SPEC.md.draft
/aido:multi-phase-spec-pr — creates the main tracking PR
For each phase (or set of parallel phases):
/aido:1-plan — plans the phase, and offers parallel worktrees for independent phases
/aido:2-implement , /aido:3-organize-commits , /aido:4-create-pr
/aido:multi-phase-sync — updates the main plan and starts the next phase
Taking over someone else's work
/aido:stabilize , pointed at the PR, remote branch, or local changes
Choose whether to reimplement from scratch or carry the original changes over and adjust them
Approve the plan it writes into SPEC.md
Continue with /aido:2-implement , /aido:3-organize-commits and /aido:4-create-pr
Run /aido:verify <PR number> to check it against what was agreed, and /code-review to look for
defects.
A migration conflict after a rebase
Run /aido:resolve-migration-conflict when another PR's migration was merged while yours was open.
The skills load these from plugins/aido/guidelines/ :
AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack. - joebew42/aido

GitHub - joebew42/aido: AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack. · GitHub
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
joebew42
aido Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
19 Commits 19 Commits Folders and files
.claude-plugin .claude-plugin docs docs plugins/ aido plugins/ aido tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Hand Claude Code the whole feature: the spec, the tests, the code, the commits and the pull request.
aido stops at every decision that matters and asks you, so what gets built stays your call. It works
with any stack.
/plugin marketplace add joebew42/aido
/plugin install aido@aido
/aido:1-plan
Every skill is available as /aido:<skill> .
Step
Skill
What it does
1
/aido:1-plan
Specifies and plans the feature in SPEC.md
2
/aido:2-implement
Builds the feature
3
/aido:3-organize-commits
Cleans up the commit history
4
/aido:4-create-pr
Creates the pull request
5
/aido:5-handle-review
Handles the review feedback
/aido:1-plan drafts the specification in SPEC.md.draft (from a GitHub issue, one you write, or a
conversation with you), checks it against your repository and refines it with you until you approve
it as SPEC.md , plans the work, creates the branch, and offers to move on to implementation. When
the change looks small and clear, it proposes a lighter route: a brief specification and plan that
you approve once, without the enabling refactoring analysis. /aido:3-organize-commits asks whether
to revert SPEC.md (the default) or keep it in the repository's docs. /aido:4-create-pr offers to
verify the implementation against the specification before pushing.
The guidelines ask for commits and PRs without AI attribution and for no auto-memory. A plugin
can't set these, so add them to your ~/.claude/settings.json :
{
"autoMemoryEnabled" : false ,
"attribution" : { "commit" : " " , "pr" : " " }
}
The plugin's hooks also block commit messages with Co-Authored-By trailers and writes to Claude's
memory directory.
Why aido: AI-Augmented Engineering
An AI-augmented engineer uses an AI agent to do more of the work, without handing over the
judgement. The agent explores the code, drafts the specification, writes the tests and the code,
and shapes the commits and the pull request. The engineer decides what gets built, approves every
step that matters, and stays accountable for the result.
aido stands for "AI do" and echoes sudo : the agent does the work, on your say-so. It is a
workflow for working that way, and each of its practices keeps that split in place:
The engineer decides, the agent proposes. Each stage ends with an explicit approval — the
specification, the plan, the commit structure, the pull request — and the agent asks instead of
guessing whenever a choice is yours.
Specification before code. Every feature starts as a draft that is refined with you into
SPEC.md , then planned. A small, clear change can take a lighter route, with the specification
and the plan approved together. Implementation only starts from an approved plan.
Small, reviewable steps. Code grows test by test with test-driven development, so each change
is easy to follow and to stop.
Good engineering practices by default. Clean architecture, deliberate refactoring, focused
tests and a clean commit history are part of every step, not an afterthought.
Your repository's rules come first. aido brings practices, not conventions. How to
test, lint, name or structure code comes from your repository — first from its CLAUDE.md and
AGENTS.md , then from the code itself.
To learn more about the wider field, see
AI-Augmented Software Engineering at
Carnegie Mellon's Software Engineering Institute.
The specification is at the center
Every stage works from one file, SPEC.md . Planning writes the specification, the plan and the
design notes into it. Implementation reads the plan and records its progress there. The pull
request description is written from it, and review feedback is checked against it. /aido:verify
checks the code against it.
Every specification has the same shape. Context says why the change is needed, Intent what
it must achieve, Out of scope what it leaves out, and Acceptance criteria — numbered AC1 ,
AC2 , … — how you will know it is done. The specification says what and why; how belongs to the
plan that planning adds below it, followed by the progress that implementation records.
Acceptance criteria are checked at every stage
The acceptance criteria are what verifies that the change does what the specification asks, not
just what the plan lists:
planning proposes them when the draft has none, and every plan item names the criteria it
serves — none is left without one
implementation writes each criterion's acceptance test first and keeps it red as the goal
post while the code behind it is built test-driven. Before it calls the feature done, it checks
every criterion: a test that covers it, or — for criteria only a person can judge, marked
(manual) — your confirmation. When one is not met, you choose to implement it, change it, or
defer it
verification checks them again from a fresh context, and any deferred criterion is named in
the pull request description
The specification is checked against your repository
A specification can be clear and still conflict with the repository it is meant for. /aido:1-plan
checks for that before you approve it, and /aido:2-implement checks again at the start of every
session, in case the default branch has moved. When nothing relevant changed, that costs one line.
A misalignment is significant when the specification contradicts your CLAUDE.md or AGENTS.md ,
crosses an architectural boundary, or changes behaviour other code relies on. The agent then explains
the conflict, cites the files that show it, and recommends revising or rejecting the specification.
You can still proceed, since intentional architectural changes are sometimes the point, and the
decision and your reason are recorded in the specification.
Verifying that you built what you agreed
/aido:verify asks one question of the current branch or any PR: does the implementation satisfy
what was agreed? /aido:4-create-pr offers to run it before pushing, and you can run it at any time.
It always runs from a fresh context, so the agent that built the change is not the one judging it,
and it changes nothing.
Its report sorts every requirement as verified , insufficient evidence or unmet , and lists
deviations from what was agreed separately. A requirement is verified only by a test that passed
during the verification, or by your recorded confirmation; claims in the specification's progress
notes do not count. On a PR with no specification, it uses the description and linked issues, and
says that this is weaker.
It complements Claude Code's /code-review rather than replacing it. /aido:verify asks whether you
built what you agreed to build, and /code-review asks whether the code has defects.
The agent does the work, but it stops before every decision that matters: it does not approve the
specification, choose a refactoring that reaches into another module, rewrite history or push
without you. The only way past these checkpoints is to deliberately run the workflow fully
unsupervised. The diagram at the top of this page shows where it stops at each stage.
For example, when you run /aido:1-plan , it stops to ask where the specification comes from: a
GitHub issue it turns into a draft, a draft you write yourself, or a conversation in which it
interviews you and writes the draft. The screenshot at the top of this page shows that question.
Skill
What it does
/aido:1-plan
draft, refine and plan a feature into SPEC.md
/aido:2-implement
implement the feature with TDD
/aido:3-organize-commits
organize the commit history
/aido:4-create-pr
create a pull request
/aido:5-handle-review
handle PR review feedback
Multi-phase feature
Skill
What it does
/aido:multi-phase-spec-pr
create the specification tracking PR
/aido:multi-phase-switch
switch to the main plan branch
/aido:multi-phase-sync
sync the main plan with progress
Utilities
Skill
What it does
/aido:resolve-migration-conflict
resolve database migration conflicts after a rebase, with any migration tool
/aido:verify
check the current branch or a PR against what was agreed, from a fresh context
/aido:stabilize
take over a PR, remote branch, or local changes made outside this workflow
/aido:map-architecture
build or update a slim index of the repo's domain areas and their building blocks
/aido:assess-architecture
assess the architecture area by area and open GitHub issues for the refactorings you pick
Other ways to use it
Write the overall feature specification in SPEC.md.draft
/aido:multi-phase-spec-pr — creates the main tracking PR
For each phase (or set of parallel phases):
/aido:1-plan — plans the phase, and offers parallel worktrees for independent phases
/aido:2-implement , /aido:3-organize-commits , /aido:4-create-pr
/aido:multi-phase-sync — updates the main plan and starts the next phase
Taking over someone else's work
/aido:stabilize , pointed at the PR, remote branch, or local changes
Choose whether to reimplement from scratch or carry the original changes over and adjust them
Approve the plan it writes into SPEC.md
Continue with /aido:2-implement , /aido:3-organize-commits and /aido:4-create-pr
Run /aido:verify <PR number> to check it against what was agreed, and /code-review to look for
defects.
A migration conflict after a rebase
Run /aido:resolve-migration-conflict when another PR's migration was merged while yours was open.
The skills load these from plugins/aido/guidelines/ :
AI does, you decide. A spec-driven, test-driven engineering workflow for Claude Code. The agent writes the spec, the code and the PRs, and you approve each step. It works with any stack.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
