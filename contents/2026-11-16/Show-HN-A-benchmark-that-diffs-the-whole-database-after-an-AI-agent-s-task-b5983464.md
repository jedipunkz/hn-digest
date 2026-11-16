---
source: "https://github.com/shubhamjoshipromail-svg/software-world-model"
hn_url: "https://news.ycombinator.com/item?id=50514157"
title: "Show HN: A benchmark that diffs the whole database after an AI agent's task"
article_title: "GitHub - shubhamjoshipromail-svg/software-world-model: Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion. · GitHub"
image: "https://opengraph.githubassets.com/b39ff64823e1adfe0b796c4fcdc3e688d21f1c114c20b55853dc95c49c661224/shubhamjoshipromail-svg/software-world-model"
author: "shubhamjoshipro"
captured_at: "2026-11-16T05:03:14Z"
capture_tool: "hn-digest"
hn_id: 50514157
score: 1
comments: 0
posted_at: "2026-11-16T04:49:45Z"
tags:
  - hacker-news
---

# Show HN: A benchmark that diffs the whole database after an AI agent's task

- HN: [50514157](https://news.ycombinator.com/item?id=50514157)
- Source: [github.com](https://github.com/shubhamjoshipromail-svg/software-world-model)
- Score: 1
- Comments: 0
- Posted: 2026-11-16T04:49:45Z

## Translation

Title: Show HN: A benchmark that diffs the whole database after an AI agent's task
Article title: GitHub - shubhamjoshipromail-svg/software-world-model: Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion. · GitHub
Description: Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion. - shubhamjoshipromail-svg/software-world-model

Article text:
GitHub - shubhamjoshipromail-svg/software-world-model: Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion. · GitHub
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
shubhamjoshipromail-svg
software-world-model Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion.
Readme MIT license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
.github .github colab colab docs docs envs envs reports reports results results scripts scripts swm swm tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
A benchmark that checks what a computer-use agent actually did to the database, not what it said it did.
Agents drive real, self-hosted business apps (Invoice Ninja, Dolibarr, Odoo 17, Odoo 18) in a browser. Every
episode starts from the same database snapshot, and afterwards the whole database is diffed . That gives four
verdicts per episode instead of one:
What we found (Phase 0, 635 scored episodes)
A version upgrade broke things quietly. From Odoo 17 to Odoo 18, task success held (78% → 76%), but
collateral damage rose from 15% to 21% and irreversible changes from 6% to 10%. Odoo 18 locks fields on posted
invoices, so to edit a reference the agent clicks Reset to Draft , which silently un-reconciles the
customer's payment . It then re-posts the invoice and reports success.
Success rates hide damage. "Clean success" (task done and nothing else touched) is up to 12 points lower
than raw success on Odoo.
Composition collapsed. Single-step skills reach up to 100%. A 3-step task (find the oldest unpaid invoice,
pay it, report the new balance) succeeded 0 of 50 times across 4 apps and 2 agents (0 of 66 counting
perturbed runs).
A small world model survives a version upgrade but not a new app. A latent dynamics model trained on
label-free exploration predicts state transitions within an app and from Odoo 17 to 18, and does not transfer
across apps at this scale.
Agents: GPT-6 Luna (accessibility tree + screenshot) and UI-TARS-1.5-7B (screenshot only), same harness, 30-step
budget. Full write-up with every table, the failure taxonomy and the limitations:
reports/phase0_writeup.md .
These are small-sample results from one person's Phase 0: 36 to 72 episodes per cell, one run per task, Wilson
95% intervals. Read Limitations before quoting a number.
You need Docker, Python 3.10 to 3.12 (or uv , which fetches one for you), and about
8 GB of free disk.
git clone https://github.com/shubhamjoshipromail-svg/software-world-model.git
cd software-world-model
make install # virtualenv, package, Chromium
make test # 18 unit tests, no Docker or API key needed
Check the published numbers without running anything (rebuilds every table from the committed episode
summaries: 635 scored, plus 35 superseded ones kept for a sensitivity check):
make report
Run an agent yourself. Start one app, then run one task per skill:
cp .env.example .env # add OPENAI_API_KEY (or OPENAI_BASE_URL for a local model)
.venv/bin/swm doctor # checks Docker, browser, key, disk
.venv/bin/swm up odoo17 # creates local credentials, starts Odoo, seeds 60 clients and 250 invoices
.venv/bin/swm run --app odoo17 --per-skill 1 --tag smoke/odoo17/clean
.venv/bin/swm report # tables and figure from your own runs/
swm up takes about six minutes the first time (image download and seeding). In Phase 0, GPT-6 Luna cost one
to three cents per successful task. Each episode writes screenshots, the step log and a summary.json under runs/<tag>/<task>/ .
Apps: invoiceninja · dolibarr · odoo17 · odoo18 . Everything binds to 127.0.0.1 , agents cannot navigate
off localhost, and no real email can be sent.
Three ways, from least to most work:
# 1. Any chat model on the OpenAI API
swm run --app odoo17 --model < model-name > --tag mymodel/odoo17/clean
# 2. Any OpenAI-compatible server (vLLM, Ollama, OpenRouter, LiteLLM): set OPENAI_BASE_URL in .env
swm run --app odoo17 --model qwen3-vl --tag qwen/odoo17/clean
# 3. Your own class
swm run --app odoo17 --model custom:my_pkg.my_agent:MyAgent --tag mine/odoo17/clean
A custom agent is one class with one method. It returns a BrowserGym
action string:
class MyAgent :
tok_in = tok_out = 0 # token counters, reported per episode
def __init__ ( self , actions , app ): ... # actions.describe() lists the action space
def act ( self , goal , obs , axtree_text , jpeg_bytes ):
return "click('42')" , "raw model output" , {}
def cost ( self ):
return 0.0 # dollars so far
Details and a worked example: docs/ADDING_AN_AGENT.md .
9 skills × 8 instances = 72 tasks per app : find a record, count, cross-record total, create client, create
invoice, edit a field, bulk status change, record a payment, and a 3-step composite. Tasks are written in each
app's own vocabulary and generated deterministically from the seed data.
One synthetic company in every app , from one generator: 60 clients (about 15% near-duplicate names such as
"Dunder Health Ltd" and "Dunder Health Limited"), 250 invoices in mixed states, about 109 payments.
Perturbations ( --perturb ): a blocking "What's new" modal, a flaky network (1 to 4 s latency, 8% HTTP 500),
and session expiry mid-task.
Per episode : success, collateral damage, irreversible changes, false completion, gave up or timed out,
steps, tokens, dollars, wall time.
The diff lives in swm/dbdiff.py (about 90 lines, MySQL and Postgres). What counts as a
legitimate change for each task is in swm/tasks_*.py .
Clean condition. Full table with perturbations and per-skill breakdown:
reports/phase0_results.md .
Neither agent was trained on any of these apps, so this table measures app difficulty, not transfer.
swm/ harness, agents, database diff, per-app tasks and verifiers, report
apps/ one adapter per app (login, snapshot, restore, capture)
envs/ docker compose files and credential templates, images pinned by digest
results/ phase0_episodes.jsonl: one line per published episode
reports/ write-up, result tables, failure taxonomy, world-model probe
tests/ unit tests, including one that rebuilds the published table
docs/ adding an agent, adding an app
colab/ serving UI-TARS-1.5-7B with vLLM on a Colab GPU
Limitations
Small samples, one run per task. Intervals are wide; several differences between cells are inside them.
Collateral whitelists are hand-built per app. App-generated noise (sessions, counters, wizards) is
whitelisted by inspection and listed in code. A different whitelist would move collateral rates by a few
points. The reset-to-draft case does not depend on it.
The diff copies every table , which is fine at this data size and would not scale to a large database.
A harness bug was found and fixed mid-run. 12 Invoice Ninja episodes could not be rerun and are excluded;
counting them as failures is reported alongside.
Developed on macOS. The unit tests run on Linux in CI; the full harness has not been run on Linux.
Raw step logs and screenshots are not in the repo (about 800 MB). The per-episode summaries are.
The write-up's limitations section is longer and more
specific.
State-based grading of agents is an active area. Agent-Diff diffs full
environment state for API agents on Slack, Linear, Box and Calendar replicas.
EntWorld verifies enterprise GUI agents by querying Odoo and EspoCRM
databases. WebArena and WorkArena are the
standard web-agent benchmarks, and this harness is built on BrowserGym .
What this repo adds is the same tasks across two major versions of one app , with collateral and irreversible
changes reported separately from success.
New apps, agents and verifier fixes are welcome: see CONTRIBUTING.md . The most useful
contribution is a result on a model we have not run; open a pull request with your results JSONL.
@misc { joshi2026softwareworldmodel ,
author = { Joshi, Shubham } ,
title = { software-world-model: do computer-use agents transfer to business software they have never seen? } ,
year = { 2026 } ,
url = { https://github.com/shubhamjoshipromail-svg/software-world-model }
}
License
MIT for the code in this repository: see LICENSE . The apps are pulled as their publishers' Docker
images and keep their own licenses; none of their code or data is redistributed here.
Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion.
Readme MIT license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion. - shubhamjoshipromail-svg/software-world-model

GitHub - shubhamjoshipromail-svg/software-world-model: Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion. · GitHub
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
shubhamjoshipromail-svg
software-world-model Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion.
Readme MIT license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
.github .github colab colab docs docs envs envs reports reports results results scripts scripts swm swm tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
A benchmark that checks what a computer-use agent actually did to the database, not what it said it did.
Agents drive real, self-hosted business apps (Invoice Ninja, Dolibarr, Odoo 17, Odoo 18) in a browser. Every
episode starts from the same database snapshot, and afterwards the whole database is diffed . That gives four
verdicts per episode instead of one:
What we found (Phase 0, 635 scored episodes)
A version upgrade broke things quietly. From Odoo 17 to Odoo 18, task success held (78% → 76%), but
collateral damage rose from 15% to 21% and irreversible changes from 6% to 10%. Odoo 18 locks fields on posted
invoices, so to edit a reference the agent clicks Reset to Draft , which silently un-reconciles the
customer's payment . It then re-posts the invoice and reports success.
Success rates hide damage. "Clean success" (task done and nothing else touched) is up to 12 points lower
than raw success on Odoo.
Composition collapsed. Single-step skills reach up to 100%. A 3-step task (find the oldest unpaid invoice,
pay it, report the new balance) succeeded 0 of 50 times across 4 apps and 2 agents (0 of 66 counting
perturbed runs).
A small world model survives a version upgrade but not a new app. A latent dynamics model trained on
label-free exploration predicts state transitions within an app and from Odoo 17 to 18, and does not transfer
across apps at this scale.
Agents: GPT-6 Luna (accessibility tree + screenshot) and UI-TARS-1.5-7B (screenshot only), same harness, 30-step
budget. Full write-up with every table, the failure taxonomy and the limitations:
reports/phase0_writeup.md .
These are small-sample results from one person's Phase 0: 36 to 72 episodes per cell, one run per task, Wilson
95% intervals. Read Limitations before quoting a number.
You need Docker, Python 3.10 to 3.12 (or uv , which fetches one for you), and about
8 GB of free disk.
git clone https://github.com/shubhamjoshipromail-svg/software-world-model.git
cd software-world-model
make install # virtualenv, package, Chromium
make test # 18 unit tests, no Docker or API key needed
Check the published numbers without running anything (rebuilds every table from the committed episode
summaries: 635 scored, plus 35 superseded ones kept for a sensitivity check):
make report
Run an agent yourself. Start one app, then run one task per skill:
cp .env.example .env # add OPENAI_API_KEY (or OPENAI_BASE_URL for a local model)
.venv/bin/swm doctor # checks Docker, browser, key, disk
.venv/bin/swm up odoo17 # creates local credentials, starts Odoo, seeds 60 clients and 250 invoices
.venv/bin/swm run --app odoo17 --per-skill 1 --tag smoke/odoo17/clean
.venv/bin/swm report # tables and figure from your own runs/
swm up takes about six minutes the first time (image download and seeding). In Phase 0, GPT-6 Luna cost one
to three cents per successful task. Each episode writes screenshots, the step log and a summary.json under runs/<tag>/<task>/ .
Apps: invoiceninja · dolibarr · odoo17 · odoo18 . Everything binds to 127.0.0.1 , agents cannot navigate
off localhost, and no real email can be sent.
Three ways, from least to most work:
# 1. Any chat model on the OpenAI API
swm run --app odoo17 --model < model-name > --tag mymodel/odoo17/clean
# 2. Any OpenAI-compatible server (vLLM, Ollama, OpenRouter, LiteLLM): set OPENAI_BASE_URL in .env
swm run --app odoo17 --model qwen3-vl --tag qwen/odoo17/clean
# 3. Your own class
swm run --app odoo17 --model custom:my_pkg.my_agent:MyAgent --tag mine/odoo17/clean
A custom agent is one class with one method. It returns a BrowserGym
action string:
class MyAgent :
tok_in = tok_out = 0 # token counters, reported per episode
def __init__ ( self , actions , app ): ... # actions.describe() lists the action space
def act ( self , goal , obs , axtree_text , jpeg_bytes ):
return "click('42')" , "raw model output" , {}
def cost ( self ):
return 0.0 # dollars so far
Details and a worked example: docs/ADDING_AN_AGENT.md .
9 skills × 8 instances = 72 tasks per app : find a record, count, cross-record total, create client, create
invoice, edit a field, bulk status change, record a payment, and a 3-step composite. Tasks are written in each
app's own vocabulary and generated deterministically from the seed data.
One synthetic company in every app , from one generator: 60 clients (about 15% near-duplicate names such as
"Dunder Health Ltd" and "Dunder Health Limited"), 250 invoices in mixed states, about 109 payments.
Perturbations ( --perturb ): a blocking "What's new" modal, a flaky network (1 to 4 s latency, 8% HTTP 500),
and session expiry mid-task.
Per episode : success, collateral damage, irreversible changes, false completion, gave up or timed out,
steps, tokens, dollars, wall time.
The diff lives in swm/dbdiff.py (about 90 lines, MySQL and Postgres). What counts as a
legitimate change for each task is in swm/tasks_*.py .
Clean condition. Full table with perturbations and per-skill breakdown:
reports/phase0_results.md .
Neither agent was trained on any of these apps, so this table measures app difficulty, not transfer.
swm/ harness, agents, database diff, per-app tasks and verifiers, report
apps/ one adapter per app (login, snapshot, restore, capture)
envs/ docker compose files and credential templates, images pinned by digest
results/ phase0_episodes.jsonl: one line per published episode
reports/ write-up, result tables, failure taxonomy, world-model probe
tests/ unit tests, including one that rebuilds the published table
docs/ adding an agent, adding an app
colab/ serving UI-TARS-1.5-7B with vLLM on a Colab GPU
Limitations
Small samples, one run per task. Intervals are wide; several differences between cells are inside them.
Collateral whitelists are hand-built per app. App-generated noise (sessions, counters, wizards) is
whitelisted by inspection and listed in code. A different whitelist would move collateral rates by a few
points. The reset-to-draft case does not depend on it.
The diff copies every table , which is fine at this data size and would not scale to a large database.
A harness bug was found and fixed mid-run. 12 Invoice Ninja episodes could not be rerun and are excluded;
counting them as failures is reported alongside.
Developed on macOS. The unit tests run on Linux in CI; the full harness has not been run on Linux.
Raw step logs and screenshots are not in the repo (about 800 MB). The per-episode summaries are.
The write-up's limitations section is longer and more
specific.
State-based grading of agents is an active area. Agent-Diff diffs full
environment state for API agents on Slack, Linear, Box and Calendar replicas.
EntWorld verifies enterprise GUI agents by querying Odoo and EspoCRM
databases. WebArena and WorkArena are the
standard web-agent benchmarks, and this harness is built on BrowserGym .
What this repo adds is the same tasks across two major versions of one app , with collateral and irreversible
changes reported separately from success.
New apps, agents and verifier fixes are welcome: see CONTRIBUTING.md . The most useful
contribution is a result on a model we have not run; open a pull request with your results JSONL.
@misc { joshi2026softwareworldmodel ,
author = { Joshi, Shubham } ,
title = { software-world-model: do computer-use agents transfer to business software they have never seen? } ,
year = { 2026 } ,
url = { https://github.com/shubhamjoshipromail-svg/software-world-model }
}
License
MIT for the code in this repository: see LICENSE . The apps are pulled as their publishers' Docker
images and keep their own licenses; none of their code or data is redistributed here.
Benchmark for computer-use agents on self-hosted business apps, graded by whole-database diffs: success, collateral damage, irreversible changes, false completion.
Readme MIT license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
