---
source: "https://github.com/santosh79/current"
hn_url: "https://news.ycombinator.com/item?id=50393255"
title: "Show HN: Current – Change control for AI agents (Elixir)"
article_title: "GitHub - santosh79/current: Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log. · GitHub"
image: "https://opengraph.githubassets.com/7fa1240ef3d5ed4d74b4f2d435c85da098eb09a1fa482ea6c581a39a39f39d73/santosh79/current"
author: "santosh_ycomb"
captured_at: "2026-11-06T15:42:29Z"
capture_tool: "hn-digest"
hn_id: 50393255
score: 2
comments: 0
posted_at: "2026-11-06T15:09:07Z"
tags:
  - hacker-news
---

# Show HN: Current – Change control for AI agents (Elixir)

- HN: [50393255](https://news.ycombinator.com/item?id=50393255)
- Source: [github.com](https://github.com/santosh79/current)
- Score: 2
- Comments: 0
- Posted: 2026-11-06T15:09:07Z

## Translation

Title: Show HN: Current – Change control for AI agents (Elixir)
Article title: GitHub - santosh79/current: Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log. · GitHub
Description: Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log. - santosh79/current

Article text:
GitHub - santosh79/current: Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log. · GitHub
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
santosh79
current Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log.
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
191 Commits 191 Commits Folders and files
.github/ workflows .github/ workflows bin bin config config docs docs examples/ nextjs-auth examples/ nextjs-auth lib lib priv/ schema priv/ schema test test .formatter.exs .formatter.exs .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md mise.toml mise.toml mix.exs mix.exs mix.lock mix.lock View all files Repository files navigation
Agents propose changes. Current validates them, waits for your approval, applies them, and records every step in a tamper-evident log.
The Elevator Pitch · See it in action ·
Quickstart · How it works ·
Provenance ·
Reproducibility ·
Verifiability ·
Architecture · Intents · CLI ·
HTTP API · MCP · Configuration ·
Security · Status · Roadmap
Why not just use Claude Code or Codex? You can. But there, the agent acts and you approve step by step as it goes. In Current, the agent can’t act at all. It hands over a complete proposal, and nothing happens until that proposal is approved and applied.
Safety : The agent is limited to doing just one thing -- producing proposals . It sees only the files it asks for that Current's read policy permits, and can only propose changes from nine typed intents. Nothing touches your project until you approve it, and only allowlisted commands can run.
Verification : You don’t take the agent’s word for it. You review the exact proposal, or have other agents review it, before anything happens, and Current checks the agent’s claims against what actually ran.
Provenance : As agents write more of the code, someone will ask why was this changed, and who approved it? Current answers from a commit log: the oldest, best-tested idea in systems.
That log is hash-chained, so changes to past history are detectable, and the full state can be rebuilt from it on another machine.
Status: early. Single node, local projects. Current is not a sandbox
or an authentication system: approved commands run with the runtime's
privileges. See Security and the
threat model .
Ask Current to add CSRF protection to the included Next.js example:
current run " add CSRF protection to this app " \
--dir examples/nextjs-auth \
--model openai:gpt-5.6-sol
The model never gets the repository. It sees the layout and asks for what
it needs, and Current fetches each file:
… round 1: the model asked to see 8 things
✓ read_file app/api/login/route.ts 769 B
✓ read_file lib/auth.ts 626 B
✓ read_file lib/rate-limit.ts 2.4 KB
✓ read_file tests/login.test.ts 847 B
✓ read_file tests/rate-limit.test.ts 2.3 KB
✓ read_file package.json 343 B
✓ read_file tsconfig.json 221 B
✓ read_file .gitignore 83 B
Reads are confined to the project. Dependencies, build output, .git and
.env* files are never shown to the model, and every read is recorded.
Agent produced proposal
proposal sha256:672c4973787b…
The login route already invokes a missing CSRF validator before parsing
credentials. Implement that validator with fail-closed same-origin Origin
checking and add focused coverage for accepted and rejected origins.
PROPOSED
1 WriteFile lib/csrf.ts
2 WriteFile tests/csrf.test.ts
3 RunCommand npm test
4 Report Add fail-closed same-origin CSRF validation for the login
route and unit tests covering valid, absent, malformed,
and cross-origin headers.
✓ Proposal decoded
✓ Proposal validated
Waiting for approval. Nothing has been applied.
The agent has proposed a change. Your project is untouched.
Review the exact files and commands, or pass the machine-readable version
( --json ) to another agent for a second opinion:
current review sha256:672c4973787b
Then decide. Send it back with feedback, and the model reads more and
returns a new proposal:
current revise sha256:672c4973787b -m " add a test for a missing Origin header "
Or approve it. Approval is recorded, but nothing runs yet:
current approve sha256:672c4973787b -m " LGTM "
✓ Approved sha256:672c4973787b
message LGTM
Nothing has been applied. Apply it with:
current apply sha256:672c4973787b
Applying is its own step. Current, not the agent, writes the files and
runs the command:
current apply sha256:672c4973787b
✓ Proposal approved
✓ Workflow completed
MODE: APPLY
4 effects executed.
Inspect
current events sha256:1429719b7546
current proposal sha256:672c4973787b
current artifact sha256:fb30687f1722 (raw model output)
Every step above, from the request and each file the model read to its raw
reply, the approval and each effect's output, is now part of the workflow's
record. current events prints it.
Elixir 1.19 on Erlang/OTP ( mise.toml pins OTP 28.5 and Elixir 1.19.5;
run mise install if you use mise )
An OpenAI or Anthropic API key
git clone https://github.com/santosh79/current.git
cd current
mix deps.get
mix compile
make client
2. Start the runtime
In one terminal, export a key and start the server. It listens on
127.0.0.1:4000 and prints every request and commit as it happens.
export ANTHROPIC_API_KEY= " ... " # or OPENAI_API_KEY
export CURRENT_ALLOWED_COMMANDS=npm # commands are denied unless allowlisted
bin/current-server
3. Make a request
In another terminal, from the repository root:
export PATH= " $PWD : $PATH " # puts the local `current` CLI on PATH
current models # the models your keys give you
MODEL= " openai:gpt-5.6 " # replace with a model listed above
current run " add CSRF protection to this app " \
--dir examples/nextjs-auth \
--model $MODEL
run stops when the proposal is waiting for review. Output varies between
models and runs. Use plan instead of run for a proposal that can be
reviewed but never applied.
current review < proposal-id >
current revise < proposal-id > -m " keep the route check, add edge cases "
current approve < proposal-id > -m " reviewed "
current apply < proposal-id >
Or current reject <proposal-id> -m "..." to end the workflow. A revision
produces a new proposal ID, which must be reviewed and approved in turn;
the earlier one is stale. IDs can be shortened to a unique prefix of at
least four hex characters.
current events < workflow-id > # the full chain, in log order
current proposal < proposal-id > # intents, claims, and what the model read
current artifact < artifact-id > # exact bytes: a file, command output, a raw reply
How it works
In a typical agent, a tool call is the action: the file is written the
moment the model decides to write it. In Current, the model's output is a
proposal , a list of typed intents:
"in lib/signup.ex , replace this text with that"; "run mix test ";
"claim the tests pass, because that command exits 0"
A proposal does nothing on its own. Between it and your project, the
runtime:
decodes the model's reply, which is untrusted bytes, into a closed
set of intent types. Anything outside it is rejected, not interpreted;
validates it: paths confined to the project, files exist, every
edit matches exactly, claims cite earlier intents they can be checked
against;
waits for a decision. A proposal is reviewed as a unit and approved,
rejected, or sent back with feedback. The reviewer can be a person or
another agent: each decision is an API call, recorded as its own event
with the reviewer's metadata. A decision must name the workflow's latest
proposal, and revisions are bounded (10 by default);
applies it, in a separate step, with a small trusted interpreter
that re-checks every path immediately before touching it and records
every result;
checks the outcome itself. Claims are evaluated against recorded
results, and a claim that doesn't hold fails the workflow.
Every step is an event. Rounded boxes are where a workflow ends.
A task is submitted with a prompt, a project directory and a model.
The directory is resolved once to its real path and recorded.
The model is given the project's layout : paths and sizes, no
contents. Dependencies, build output, VCS directories and .env* files
are left out.
The model may reply with an observe request: up to 20 read_file ,
glob or grep intents, each with a reason. The runtime runs them
read-only, under path confinement and the read policy, and asks again
with the results. Up to 8 rounds, 400 KB of results in total, 200 KB
per file.
The model replies with a proposal . Its exact reply is stored as an
artifact, whatever it contains.
The proposal is decoded into typed intents
( wire contract ), every path is made absolute
against the project, and it is validated .
Validation answers "is this well-formed and in bounds right now?", not
"will it apply cleanly?". The project can change before apply, and those
failures are caught and recorded then. Producing a proposal is bounded by
llm_proposal_timeout_ms per call and the task's --timeout-minutes
(default 10) overall.
apply starts the interpreter. It re-checks every path against the
filesystem as it is now (following symlinks), then carries out the
intents in order. The first failure, or the first claim that doesn't
hold, stops it and fails the workflow. Effects are not rolled back. The
re-check narrows the window for a path to escape the project but doesn't
close it (T2 in the threat model ).
created → planning → validating → waiting_for_approval
plan mode: → planned
apply mode: → approved → applying → assembling_result → completed
→ rejected
→ revision_requested → planning → validating → waiting_for_approval
(bounded; each pass adds a proposal to the workflow)
any stage: → failed (production, validation or an effect failed)
→ abandoned (no further automatic progress is possible)
planning: the model sees the layout, may ask to read files, then returns a proposal
Three primitives
Primitive
When
Role
More
Intent
Before execution
A proposal, as data
intent.md
Event
After execution or decision
A fact
event.md
Artifact
Throughout
Immutable bytes, named by their SHA-256
artifact.md
Intents propose change. Artifacts hold bytes. Events record fact.
State is derived from the event log.
Provenance: from a change back to the request
Because the runtime is the only thing that acts, and records each step
before taking the next, every change sits at the end of an explicit chain:
TaskCreated the request: prompt, model, project, submitted_at
└─ WorkflowCreated
└─ ProposalProduced intents, what the model read first, its raw reply
└─ ProposalValidated
├─ Pr

[truncated]

## Original Extract

Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log. - santosh79/current

GitHub - santosh79/current: Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log. · GitHub
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
santosh79
current Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Change control for AI agents: agents propose, the runtime acts, every step lands in a tamper-evident log.
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
191 Commits 191 Commits Folders and files
.github/ workflows .github/ workflows bin bin config config docs docs examples/ nextjs-auth examples/ nextjs-auth lib lib priv/ schema priv/ schema test test .formatter.exs .formatter.exs .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md mise.toml mise.toml mix.exs mix.exs mix.lock mix.lock View all files Repository files navigation
Agents propose changes. Current validates them, waits for your approval, applies them, and records every step in a tamper-evident log.
The Elevator Pitch · See it in action ·
Quickstart · How it works ·
Provenance ·
Reproducibility ·
Verifiability ·
Architecture · Intents · CLI ·
HTTP API · MCP · Configuration ·
Security · Status · Roadmap
Why not just use Claude Code or Codex? You can. But there, the agent acts and you approve step by step as it goes. In Current, the agent can’t act at all. It hands over a complete proposal, and nothing happens until that proposal is approved and applied.
Safety : The agent is limited to doing just one thing -- producing proposals . It sees only the files it asks for that Current's read policy permits, and can only propose changes from nine typed intents. Nothing touches your project until you approve it, and only allowlisted commands can run.
Verification : You don’t take the agent’s word for it. You review the exact proposal, or have other agents review it, before anything happens, and Current checks the agent’s claims against what actually ran.
Provenance : As agents write more of the code, someone will ask why was this changed, and who approved it? Current answers from a commit log: the oldest, best-tested idea in systems.
That log is hash-chained, so changes to past history are detectable, and the full state can be rebuilt from it on another machine.
Status: early. Single node, local projects. Current is not a sandbox
or an authentication system: approved commands run with the runtime's
privileges. See Security and the
threat model .
Ask Current to add CSRF protection to the included Next.js example:
current run " add CSRF protection to this app " \
--dir examples/nextjs-auth \
--model openai:gpt-5.6-sol
The model never gets the repository. It sees the layout and asks for what
it needs, and Current fetches each file:
… round 1: the model asked to see 8 things
✓ read_file app/api/login/route.ts 769 B
✓ read_file lib/auth.ts 626 B
✓ read_file lib/rate-limit.ts 2.4 KB
✓ read_file tests/login.test.ts 847 B
✓ read_file tests/rate-limit.test.ts 2.3 KB
✓ read_file package.json 343 B
✓ read_file tsconfig.json 221 B
✓ read_file .gitignore 83 B
Reads are confined to the project. Dependencies, build output, .git and
.env* files are never shown to the model, and every read is recorded.
Agent produced proposal
proposal sha256:672c4973787b…
The login route already invokes a missing CSRF validator before parsing
credentials. Implement that validator with fail-closed same-origin Origin
checking and add focused coverage for accepted and rejected origins.
PROPOSED
1 WriteFile lib/csrf.ts
2 WriteFile tests/csrf.test.ts
3 RunCommand npm test
4 Report Add fail-closed same-origin CSRF validation for the login
route and unit tests covering valid, absent, malformed,
and cross-origin headers.
✓ Proposal decoded
✓ Proposal validated
Waiting for approval. Nothing has been applied.
The agent has proposed a change. Your project is untouched.
Review the exact files and commands, or pass the machine-readable version
( --json ) to another agent for a second opinion:
current review sha256:672c4973787b
Then decide. Send it back with feedback, and the model reads more and
returns a new proposal:
current revise sha256:672c4973787b -m " add a test for a missing Origin header "
Or approve it. Approval is recorded, but nothing runs yet:
current approve sha256:672c4973787b -m " LGTM "
✓ Approved sha256:672c4973787b
message LGTM
Nothing has been applied. Apply it with:
current apply sha256:672c4973787b
Applying is its own step. Current, not the agent, writes the files and
runs the command:
current apply sha256:672c4973787b
✓ Proposal approved
✓ Workflow completed
MODE: APPLY
4 effects executed.
Inspect
current events sha256:1429719b7546
current proposal sha256:672c4973787b
current artifact sha256:fb30687f1722 (raw model output)
Every step above, from the request and each file the model read to its raw
reply, the approval and each effect's output, is now part of the workflow's
record. current events prints it.
Elixir 1.19 on Erlang/OTP ( mise.toml pins OTP 28.5 and Elixir 1.19.5;
run mise install if you use mise )
An OpenAI or Anthropic API key
git clone https://github.com/santosh79/current.git
cd current
mix deps.get
mix compile
make client
2. Start the runtime
In one terminal, export a key and start the server. It listens on
127.0.0.1:4000 and prints every request and commit as it happens.
export ANTHROPIC_API_KEY= " ... " # or OPENAI_API_KEY
export CURRENT_ALLOWED_COMMANDS=npm # commands are denied unless allowlisted
bin/current-server
3. Make a request
In another terminal, from the repository root:
export PATH= " $PWD : $PATH " # puts the local `current` CLI on PATH
current models # the models your keys give you
MODEL= " openai:gpt-5.6 " # replace with a model listed above
current run " add CSRF protection to this app " \
--dir examples/nextjs-auth \
--model $MODEL
run stops when the proposal is waiting for review. Output varies between
models and runs. Use plan instead of run for a proposal that can be
reviewed but never applied.
current review < proposal-id >
current revise < proposal-id > -m " keep the route check, add edge cases "
current approve < proposal-id > -m " reviewed "
current apply < proposal-id >
Or current reject <proposal-id> -m "..." to end the workflow. A revision
produces a new proposal ID, which must be reviewed and approved in turn;
the earlier one is stale. IDs can be shortened to a unique prefix of at
least four hex characters.
current events < workflow-id > # the full chain, in log order
current proposal < proposal-id > # intents, claims, and what the model read
current artifact < artifact-id > # exact bytes: a file, command output, a raw reply
How it works
In a typical agent, a tool call is the action: the file is written the
moment the model decides to write it. In Current, the model's output is a
proposal , a list of typed intents:
"in lib/signup.ex , replace this text with that"; "run mix test ";
"claim the tests pass, because that command exits 0"
A proposal does nothing on its own. Between it and your project, the
runtime:
decodes the model's reply, which is untrusted bytes, into a closed
set of intent types. Anything outside it is rejected, not interpreted;
validates it: paths confined to the project, files exist, every
edit matches exactly, claims cite earlier intents they can be checked
against;
waits for a decision. A proposal is reviewed as a unit and approved,
rejected, or sent back with feedback. The reviewer can be a person or
another agent: each decision is an API call, recorded as its own event
with the reviewer's metadata. A decision must name the workflow's latest
proposal, and revisions are bounded (10 by default);
applies it, in a separate step, with a small trusted interpreter
that re-checks every path immediately before touching it and records
every result;
checks the outcome itself. Claims are evaluated against recorded
results, and a claim that doesn't hold fails the workflow.
Every step is an event. Rounded boxes are where a workflow ends.
A task is submitted with a prompt, a project directory and a model.
The directory is resolved once to its real path and recorded.
The model is given the project's layout : paths and sizes, no
contents. Dependencies, build output, VCS directories and .env* files
are left out.
The model may reply with an observe request: up to 20 read_file ,
glob or grep intents, each with a reason. The runtime runs them
read-only, under path confinement and the read policy, and asks again
with the results. Up to 8 rounds, 400 KB of results in total, 200 KB
per file.
The model replies with a proposal . Its exact reply is stored as an
artifact, whatever it contains.
The proposal is decoded into typed intents
( wire contract ), every path is made absolute
against the project, and it is validated .
Validation answers "is this well-formed and in bounds right now?", not
"will it apply cleanly?". The project can change before apply, and those
failures are caught and recorded then. Producing a proposal is bounded by
llm_proposal_timeout_ms per call and the task's --timeout-minutes
(default 10) overall.
apply starts the interpreter. It re-checks every path against the
filesystem as it is now (following symlinks), then carries out the
intents in order. The first failure, or the first claim that doesn't
hold, stops it and fails the workflow. Effects are not rolled back. The
re-check narrows the window for a path to escape the project but doesn't
close it (T2 in the threat model ).
created → planning → validating → waiting_for_approval
plan mode: → planned
apply mode: → approved → applying → assembling_result → completed
→ rejected
→ revision_requested → planning → validating → waiting_for_approval
(bounded; each pass adds a proposal to the workflow)
any stage: → failed (production, validation or an effect failed)
→ abandoned (no further automatic progress is possible)
planning: the model sees the layout, may ask to read files, then returns a proposal
Three primitives
Primitive
When
Role
More
Intent
Before execution
A proposal, as data
intent.md
Event
After execution or decision
A fact
event.md
Artifact
Throughout
Immutable bytes, named by their SHA-256
artifact.md
Intents propose change. Artifacts hold bytes. Events record fact.
State is derived from the event log.
Provenance: from a change back to the request
Because the runtime is the only thing that acts, and records each step
before taking the next, every change sits at the end of an explicit chain:
TaskCreated the request: prompt, model, project, submitted_at
└─ WorkflowCreated
└─ ProposalProduced intents, what the model read first, its raw reply
└─ ProposalValidated
├─ Pr

[truncated]
