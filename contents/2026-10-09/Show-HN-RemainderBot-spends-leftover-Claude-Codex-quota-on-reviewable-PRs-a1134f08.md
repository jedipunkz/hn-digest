---
source: "https://github.com/arimwilson/openremainderbot"
hn_url: "https://news.ycombinator.com/item?id=50023930"
title: "Show HN: RemainderBot - spends leftover Claude/Codex quota on reviewable PRs"
article_title: "GitHub - arimwilson/openremainderbot: Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo. · GitHub"
image: "https://opengraph.githubassets.com/b9f9b698a61b51adcce75cabf85979b97d07e19c0345d2151d95e65e670e539c/arimwilson/openremainderbot"
author: "ariwilson"
captured_at: "2026-10-09T18:02:07Z"
capture_tool: "hn-digest"
hn_id: 50023930
score: 1
comments: 0
posted_at: "2026-10-09T17:31:50Z"
tags:
  - hacker-news
---

# Show HN: RemainderBot - spends leftover Claude/Codex quota on reviewable PRs

- HN: [50023930](https://news.ycombinator.com/item?id=50023930)
- Source: [github.com](https://github.com/arimwilson/openremainderbot)
- Score: 1
- Comments: 0
- Posted: 2026-10-09T17:31:50Z

## Translation

Title: Show HN: RemainderBot - spends leftover Claude/Codex quota on reviewable PRs
Article title: GitHub - arimwilson/openremainderbot: Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo. · GitHub
Description: Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo. - arimwilson/openremainderbot

Article text:
GitHub - arimwilson/openremainderbot: Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo. · GitHub
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
arimwilson
openremainderbot Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo.
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
docs docs examples examples remainderbot remainderbot tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md GOALS.example.md GOALS.example.md LICENSE LICENSE README.md README.md sources.example.toml sources.example.toml View all files Repository files navigation
Your Claude or Codex subscription resets every week, and whatever quota you didn't use is
gone. RemainderBot notices when a reset is a few hours away with quota left, and spends
it on finished, easily reviewable pieces of work: pull requests in your own private repo.
One run, illustrated, for a made-up founder's dog-grooming scheduler. The real run it's
drawn from is in examples/groomer-time-off , with its
pull request , plan, and patch. Three more runs,
including a rejected one and the run that redid it from the rejection note, are in
examples/ .
Trigger. An hourly cron job reads each CLI's own quota numbers. When the weekly
reset is less than 4 hours away and at least 10% is left, it runs.
( DESIGN.md §2 )
Select. The agent reads three things. A fresh snapshot of your sources (the task
lists, issues, docs, and notes you already keep) supplies the candidate tasks. Your
GOALS.md says which of those matter and in what order, since your sources mix
errands with real work and don't rank anything. The verdicts on past runs keep it from
repeating itself. It picks one task, checks that the task isn't already done, and
commits a plan before building anything.
( §3 )
Build. It works against a hard deadline set from the reset time, sized to the quota
left: a document, a feature with tests, or a prototype.
( §4 )
Finish. If the run ends without a done marker, a short finalize pass may only cut
scope and write up what works, never add to it.
PR and feedback. Every run opens a PR whose body says what was built, how it was
verified, and the one step that ships it. Merge means approved; close with a comment
means rejected, and the next run reads your comment. ( §5 )
A dedicated machine , such as a small VPS. The agent runs there with every
permission check off, so that machine is the sandbox (see Security ).
Python 3.11 or newer, git , and gh .
Claude Code ( claude ) and/or the Codex CLI ( codex ), logged in with a subscription.
gws , only if you use Google
Tasks or Docs as sources.
RemainderBot is a small Python package with no dependencies outside the standard library.
Every instance is a private repo. A run branch carries a copy of everything the agent
read and its whole log, so the code refuses to publish to a public repo
( Privacy ).
On your own computer (anywhere gh can create a repo):
git clone https://github.com/arimwilson/openremainderbot.git remainderbot
cd remainderbot
git remote rename origin upstream
gh repo create remainderbot --private --source . --remote origin --push
cp -n GOALS.example.md GOALS.md
cp -n sources.example.toml sources.toml
Don't fork. A GitHub fork of a public repo can't be private. Don't use "Use this
template" either: a template copy has unrelated history, so your first update would
conflict on every file. A clone with upstream as a second remote gets updates with a
plain git merge ( Updating ).
Edit GOALS.md and sources.toml (the next two sections), then commit and push them:
git add GOALS.md sources.toml
git commit -m " My goals and sources "
git push
The bot creates INBOX.md and state.json itself.
Your sources say what could be done; GOALS.md says what should be. Of everything a
run reads, it's the only thing you write for the bot ( sources.toml just points at
things you already keep), and every run reads it first. GOALS.example.md is a complete
example for a made-up founder. It has four sections:
Priorities : what the work is for, in tiers. A modest artifact in a high tier beats
an impressive one in a low tier.
What to take off my plate (optional): the work you would hand to an assistant,
and the work you'd rather keep.
Rules : your additions to the defaults built into the prompt. The defaults are: no
errands, drafted never sent, one action to ship, no repeats, and nothing that needs
your credentials to verify. For example:
## Rules
- Prefer, in tier order: (a) issues labeled ` customer ` ; (b) newsletter drafts from
my outline notes; (c) anything in the repos' TODO files.
- Never change billing code; I write every change there myself.
Manual goals : one-off requests. An open one outranks everything else.
Edit it whenever you like and push. Each tick pulls main first.
sources.toml says what the snapshot reads. Each [[sources]] table becomes
goals/snapshot/<name>.md , written by the adapter its type names:
version = 1
[[ sources ]]
name = " issues "
type = " gh-issues "
owner = " your-github-user "
[[ sources ]]
name = " todo "
type = " file "
paths = [ " ~/notes/TODO.md " ]
[[ sources ]]
name = " notes "
type = " gh-markdown "
repo = " your-github-user/notes "
private = true # nothing drawn from it goes into anything meant for publication
[[ sources . parts ]] # belongs to the [[sources]] table above it, however it's indented
include = " projects/*.md "
sections = [ " Open questions " , " Next " ]
Type
Reads
gws-tasks
open Google Tasks
gws-doc
a Google Doc, as Markdown
gh-issues
open GitHub issues across an owner or a list of repos
gh-roadmaps
each repo's README head plus its TODO/ROADMAP/PLAN files
gh-markdown
chosen sections, frontmatter, or whole files from a repo of Markdown
file
local files, by glob
command
the stdout of any command: an export from Linear, Notion, or your own script
Every field, each adapter's output, and the access each one needs are in
docs/sources.md . A source whose fetch fails keeps its last content
under a stale: header; the others still refresh.
Do this on the machine that will run the bot. A run there has full access, and that is
by design. Install the CLIs, then log in to each one as the user the bot will run as:
gh auth login # see Security for a least-privilege token
gh auth setup-git # lets git push over HTTPS with that login
claude # log in with your subscription, then exit; and/or: codex login
gws auth login --readonly -s tasks,docs,drive # only for gws-* sources
gh repo clone your-github-user/remainderbot ~ /remainderbot
cd ~ /remainderbot
Then:
python3 -m remainderbot doctor
python3 -m remainderbot snapshot # writes goals/snapshot/*.md; read them
python3 -m remainderbot decide --dry-run # what the trigger rule would do now; claims nothing
python3 -m remainderbot run --provider claude --deadline-minutes 20 --size small
doctor prints one line per check and exits nonzero if any check fails. It checks:
the CLIs, their logins, and whether it can read your quota;
that sources.toml is valid and GOALS.md is no longer the example;
that your edits are committed;
that origin is private and accepts a push (a git push --dry-run , which creates
nothing).
run launches a real agent session on a branch run/<id> . That spends real quota: a
small run typically uses a few percent of the week and much of the 5-hour window. When it
ends, run goes back to main . Nothing is pushed unless you add --publish . To read
what it made:
git show run/ < id > :runs/ < id > /README.md
git checkout run/ < id > # look around, then: git checkout main
Install on a server
Run crontab -e and add two lines:
PATH=/home/you/.local/bin:/usr/local/bin:/usr/bin:/bin
0 * * * * cd $HOME/remainderbot && RUN_DISABLED=1 python3 -m remainderbot tick >> $HOME/remainderbot.log 2>&1
Cron's own PATH is only /usr/bin:/bin . For the PATH= line, paste the output of
echo $PATH from your login shell, since cron does not expand variables there.
RUN_DISABLED=1 makes each tick log its decision without running anything. Each hour
the log should show something like this:
2026-10-08 02:00:01Z tick start
2026-10-08 02:00:07Z snapshot issues.md: ok
2026-10-08 02:00:08Z usage: claude 5h 3% / weekly 41% used; codex 5h 0% / weekly 12% used
2026-10-08 02:00:08Z claude: weekly reset in 51.2h (> 4.0h window), 59% left
2026-10-08 02:00:08Z codex: weekly reset in 101.7h (> 4.0h window), 88% left
2026-10-08 02:00:08Z skip: no provider eligible
2026-10-08 02:00:08Z tick end
Inside the window, the tick logs RUN_DISABLED=1: would run now instead. After a day of
clean ticks, take RUN_DISABLED=1 off the crontab line.
The PR is the notification. The bot opens PRs with your gh login, and GitHub
doesn't email you about your own actions by default. Turn on Include your own updates
in GitHub's email notification settings, or give the bot its own account.
Each run's README (the PR body) says what exists now, how it was verified, and your
one next step: merge, git am a patch, paste a draft, or run one command. The work
itself is in runs/<id>/artifact/ , or in artifact.patch when it belongs in another
repo.
Merge means approved. Close means rejected. The last PR comment or review is
the note the next run reads, so before you close a PR, say what should change.
INBOX.md is the ledger on main : one row per run, with its status and your note.
Each tick updates it from the PRs, and the agent reads it so it never repeats a task.
Steer the bot through GOALS.md and PR comments, not by editing INBOX.md.
Data
Where it goes
Who can read it
Your sources
read by gws , gh , local files, or your command
unchanged
goals/snapshot/
the server's disk; gitignored on main
the server
runs/<id>/snapshot/ , PLAN.md , log/
committed on the run branch
anyone with read access to your instance repo
The PR and INBOX.md
your instance repo
same
The agent's session
Anthropic or OpenAI, as with any Claude Code or Codex use
the provider
The public-repo guard. Before the first push in a process, gh repo view must say
that every URL origin pushes to is private (or internal). If it can't tell, the push
is refused and tried again next tick. tick checks this right after its pull, before
it syncs or runs anything, so a public instance never spends quota.
ALLOW_PUBLIC_REPO=1 turns the guard off; use it only for a repo meant to be public,
with sources that are all public.
private sources. The worker prompt names every source marked private . Nothing
drawn from them may go into anyt

[truncated]

## Original Extract

Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo. - arimwilson/openremainderbot

GitHub - arimwilson/openremainderbot: Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo. · GitHub
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
arimwilson
openremainderbot Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Spend the subscription quota you didn't use and achieve your goals: finished, reviewable PRs in your own private repo.
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
docs docs examples examples remainderbot remainderbot tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md GOALS.example.md GOALS.example.md LICENSE LICENSE README.md README.md sources.example.toml sources.example.toml View all files Repository files navigation
Your Claude or Codex subscription resets every week, and whatever quota you didn't use is
gone. RemainderBot notices when a reset is a few hours away with quota left, and spends
it on finished, easily reviewable pieces of work: pull requests in your own private repo.
One run, illustrated, for a made-up founder's dog-grooming scheduler. The real run it's
drawn from is in examples/groomer-time-off , with its
pull request , plan, and patch. Three more runs,
including a rejected one and the run that redid it from the rejection note, are in
examples/ .
Trigger. An hourly cron job reads each CLI's own quota numbers. When the weekly
reset is less than 4 hours away and at least 10% is left, it runs.
( DESIGN.md §2 )
Select. The agent reads three things. A fresh snapshot of your sources (the task
lists, issues, docs, and notes you already keep) supplies the candidate tasks. Your
GOALS.md says which of those matter and in what order, since your sources mix
errands with real work and don't rank anything. The verdicts on past runs keep it from
repeating itself. It picks one task, checks that the task isn't already done, and
commits a plan before building anything.
( §3 )
Build. It works against a hard deadline set from the reset time, sized to the quota
left: a document, a feature with tests, or a prototype.
( §4 )
Finish. If the run ends without a done marker, a short finalize pass may only cut
scope and write up what works, never add to it.
PR and feedback. Every run opens a PR whose body says what was built, how it was
verified, and the one step that ships it. Merge means approved; close with a comment
means rejected, and the next run reads your comment. ( §5 )
A dedicated machine , such as a small VPS. The agent runs there with every
permission check off, so that machine is the sandbox (see Security ).
Python 3.11 or newer, git , and gh .
Claude Code ( claude ) and/or the Codex CLI ( codex ), logged in with a subscription.
gws , only if you use Google
Tasks or Docs as sources.
RemainderBot is a small Python package with no dependencies outside the standard library.
Every instance is a private repo. A run branch carries a copy of everything the agent
read and its whole log, so the code refuses to publish to a public repo
( Privacy ).
On your own computer (anywhere gh can create a repo):
git clone https://github.com/arimwilson/openremainderbot.git remainderbot
cd remainderbot
git remote rename origin upstream
gh repo create remainderbot --private --source . --remote origin --push
cp -n GOALS.example.md GOALS.md
cp -n sources.example.toml sources.toml
Don't fork. A GitHub fork of a public repo can't be private. Don't use "Use this
template" either: a template copy has unrelated history, so your first update would
conflict on every file. A clone with upstream as a second remote gets updates with a
plain git merge ( Updating ).
Edit GOALS.md and sources.toml (the next two sections), then commit and push them:
git add GOALS.md sources.toml
git commit -m " My goals and sources "
git push
The bot creates INBOX.md and state.json itself.
Your sources say what could be done; GOALS.md says what should be. Of everything a
run reads, it's the only thing you write for the bot ( sources.toml just points at
things you already keep), and every run reads it first. GOALS.example.md is a complete
example for a made-up founder. It has four sections:
Priorities : what the work is for, in tiers. A modest artifact in a high tier beats
an impressive one in a low tier.
What to take off my plate (optional): the work you would hand to an assistant,
and the work you'd rather keep.
Rules : your additions to the defaults built into the prompt. The defaults are: no
errands, drafted never sent, one action to ship, no repeats, and nothing that needs
your credentials to verify. For example:
## Rules
- Prefer, in tier order: (a) issues labeled ` customer ` ; (b) newsletter drafts from
my outline notes; (c) anything in the repos' TODO files.
- Never change billing code; I write every change there myself.
Manual goals : one-off requests. An open one outranks everything else.
Edit it whenever you like and push. Each tick pulls main first.
sources.toml says what the snapshot reads. Each [[sources]] table becomes
goals/snapshot/<name>.md , written by the adapter its type names:
version = 1
[[ sources ]]
name = " issues "
type = " gh-issues "
owner = " your-github-user "
[[ sources ]]
name = " todo "
type = " file "
paths = [ " ~/notes/TODO.md " ]
[[ sources ]]
name = " notes "
type = " gh-markdown "
repo = " your-github-user/notes "
private = true # nothing drawn from it goes into anything meant for publication
[[ sources . parts ]] # belongs to the [[sources]] table above it, however it's indented
include = " projects/*.md "
sections = [ " Open questions " , " Next " ]
Type
Reads
gws-tasks
open Google Tasks
gws-doc
a Google Doc, as Markdown
gh-issues
open GitHub issues across an owner or a list of repos
gh-roadmaps
each repo's README head plus its TODO/ROADMAP/PLAN files
gh-markdown
chosen sections, frontmatter, or whole files from a repo of Markdown
file
local files, by glob
command
the stdout of any command: an export from Linear, Notion, or your own script
Every field, each adapter's output, and the access each one needs are in
docs/sources.md . A source whose fetch fails keeps its last content
under a stale: header; the others still refresh.
Do this on the machine that will run the bot. A run there has full access, and that is
by design. Install the CLIs, then log in to each one as the user the bot will run as:
gh auth login # see Security for a least-privilege token
gh auth setup-git # lets git push over HTTPS with that login
claude # log in with your subscription, then exit; and/or: codex login
gws auth login --readonly -s tasks,docs,drive # only for gws-* sources
gh repo clone your-github-user/remainderbot ~ /remainderbot
cd ~ /remainderbot
Then:
python3 -m remainderbot doctor
python3 -m remainderbot snapshot # writes goals/snapshot/*.md; read them
python3 -m remainderbot decide --dry-run # what the trigger rule would do now; claims nothing
python3 -m remainderbot run --provider claude --deadline-minutes 20 --size small
doctor prints one line per check and exits nonzero if any check fails. It checks:
the CLIs, their logins, and whether it can read your quota;
that sources.toml is valid and GOALS.md is no longer the example;
that your edits are committed;
that origin is private and accepts a push (a git push --dry-run , which creates
nothing).
run launches a real agent session on a branch run/<id> . That spends real quota: a
small run typically uses a few percent of the week and much of the 5-hour window. When it
ends, run goes back to main . Nothing is pushed unless you add --publish . To read
what it made:
git show run/ < id > :runs/ < id > /README.md
git checkout run/ < id > # look around, then: git checkout main
Install on a server
Run crontab -e and add two lines:
PATH=/home/you/.local/bin:/usr/local/bin:/usr/bin:/bin
0 * * * * cd $HOME/remainderbot && RUN_DISABLED=1 python3 -m remainderbot tick >> $HOME/remainderbot.log 2>&1
Cron's own PATH is only /usr/bin:/bin . For the PATH= line, paste the output of
echo $PATH from your login shell, since cron does not expand variables there.
RUN_DISABLED=1 makes each tick log its decision without running anything. Each hour
the log should show something like this:
2026-10-08 02:00:01Z tick start
2026-10-08 02:00:07Z snapshot issues.md: ok
2026-10-08 02:00:08Z usage: claude 5h 3% / weekly 41% used; codex 5h 0% / weekly 12% used
2026-10-08 02:00:08Z claude: weekly reset in 51.2h (> 4.0h window), 59% left
2026-10-08 02:00:08Z codex: weekly reset in 101.7h (> 4.0h window), 88% left
2026-10-08 02:00:08Z skip: no provider eligible
2026-10-08 02:00:08Z tick end
Inside the window, the tick logs RUN_DISABLED=1: would run now instead. After a day of
clean ticks, take RUN_DISABLED=1 off the crontab line.
The PR is the notification. The bot opens PRs with your gh login, and GitHub
doesn't email you about your own actions by default. Turn on Include your own updates
in GitHub's email notification settings, or give the bot its own account.
Each run's README (the PR body) says what exists now, how it was verified, and your
one next step: merge, git am a patch, paste a draft, or run one command. The work
itself is in runs/<id>/artifact/ , or in artifact.patch when it belongs in another
repo.
Merge means approved. Close means rejected. The last PR comment or review is
the note the next run reads, so before you close a PR, say what should change.
INBOX.md is the ledger on main : one row per run, with its status and your note.
Each tick updates it from the PRs, and the agent reads it so it never repeats a task.
Steer the bot through GOALS.md and PR comments, not by editing INBOX.md.
Data
Where it goes
Who can read it
Your sources
read by gws , gh , local files, or your command
unchanged
goals/snapshot/
the server's disk; gitignored on main
the server
runs/<id>/snapshot/ , PLAN.md , log/
committed on the run branch
anyone with read access to your instance repo
The PR and INBOX.md
your instance repo
same
The agent's session
Anthropic or OpenAI, as with any Claude Code or Codex use
the provider
The public-repo guard. Before the first push in a process, gh repo view must say
that every URL origin pushes to is private (or internal). If it can't tell, the push
is refused and tried again next tick. tick checks this right after its pull, before
it syncs or runs anything, so a public instance never spends quota.
ALLOW_PUBLIC_REPO=1 turns the guard off; use it only for a repo meant to be public,
with sources that are all public.
private sources. The worker prompt names every source marked private . Nothing
drawn from them may go into anyt

[truncated]
