---
source: "https://github.com/kiritbasu/causal-analyst"
hn_url: "https://news.ycombinator.com/item?id=49855115"
title: "Show HN: Causal analyst agent skill for Claude"
article_title: "GitHub - kiritbasu/causal-analyst: Causal analysis for people who aren't data scientists: an Agent Skill for Claude · GitHub"
image: "https://opengraph.githubassets.com/8d47e99f03f7e76be653963d7e24be460bd31cb4dea7b50a98aee153f5e217e0/kiritbasu/causal-analyst"
author: "2au_observer"
captured_at: "2026-09-26T10:54:57Z"
capture_tool: "hn-digest"
hn_id: 49855115
score: 2
comments: 0
posted_at: "2026-09-26T10:24:33Z"
tags:
  - hacker-news
---

# Show HN: Causal analyst agent skill for Claude

- HN: [49855115](https://news.ycombinator.com/item?id=49855115)
- Source: [github.com](https://github.com/kiritbasu/causal-analyst)
- Score: 2
- Comments: 0
- Posted: 2026-09-26T10:24:33Z

## Translation

Title: Show HN: Causal analyst agent skill for Claude
Article title: GitHub - kiritbasu/causal-analyst: Causal analysis for people who aren't data scientists: an Agent Skill for Claude · GitHub
Description: Causal analysis for people who aren't data scientists: an Agent Skill for Claude - kiritbasu/causal-analyst
HN text: I'm not a Data Scientist but have been super intrigued with causal inference off late, so I made a Claude skill to simplify the modelling and experimentation process. You upload your dataset to Claude (Opus) and this skill picks your brain for your domain knowledge about the dataset and comes up with a causal dag that you can iterate through. It then runs the data through a number of classical statistics, ML and newer transformer models to come up with outcomes. And it finally creates a html report to help you understand the results.

Article text:
GitHub - kiritbasu/causal-analyst: Causal analysis for people who aren't data scientists: an Agent Skill for Claude · GitHub
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
kiritbasu
/
causal-analyst
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.github/ workflows .github/ workflows docs/ images docs/ images evals evals examples examples skills/ causal-analyst skills/ causal-analyst tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md requirements-optional.txt requirements-optional.txt requirements.txt requirements.txt View all files Repository files navigation
Causal analysis for people who aren't data scientists. An Agent Skill that lets Claude answer "did X actually cause Y?" from your data. You bring the question and what you know about your business. Claude does the modelling, checks how far to trust the answer, and hands back a one-page report. When the data can't answer the question, it says so.
Live example reports: Loyalty program (grade C) · Sales calls (grade D, "can't tell") · AI training (grade B, three traps) · Quick start
Contents: Why this exists · What it's like to use · The report · Quick start · How it works · Does it work? · Which model to use · Data and privacy · Limitations and roadmap
Good causal analysis needs two kinds of knowledge that rarely sit in one person:
Data science: which methods to use, how to check them, and what the numbers can and can't support.
Domain knowledge: what happened before what, how people ended up getting the treatment, and what important factor isn't in the data.
The marketing lead knows that "points redeemed" only exists after someone joins, and that the program was pushed to big spenders. The data scientist knows that controlling for points will wreck the estimate, and that targeting big spenders creates a gap that isn't the program's doing. Usually it takes both people and a lot of back and forth. Often the domain expert simply ends up with a correlation.
This skill lets the domain expert work directly with Claude. Claude handles the modelling. It asks, in plain language, only the questions that need domain knowledge, turns the answers into a causal diagram the expert confirms, then runs and checks the analysis and explains the result. The expert supplies what only they know; the skill supplies the rest.
It also guards against the ways capable models go wrong. Frontier models already get the arithmetic right: in our tests plain Claude matched the skill's estimate on a clean case to within a few percent. The failures are judgment calls:
Controlling for the wrong thing. Adjusting for "points redeemed" turns a +$10 effect into −$22.
Answering a question the data can't support. If reps pick whom to call using a gut feel that isn't recorded, no amount of adjustment recovers the effect of a call.
Choosing the method after seeing the results , and quietly drifting toward the hoped-for answer.
Overstating certainty. Plain Claude twice gave confident ranges that missed the true answer.
So the workflow puts those judgments in the open. Controls are confirmed as recorded before the treatment. The expert signs off the diagram. The method is fixed in advance. There's an explicit "we can't tell" grade, and the risk from hidden factors is sized.
You attach a file and ask in your own words. The skill takes it from there, stopping three times for your input.
You: [attaches customers.csv] Did our loyalty program actually raise monthly spend? We think newer customers benefit more.
1. Claude checks the data and points out anything odd before going further:
4,000 customers, 9 columns. points_redeemed almost perfectly predicts who joined, which usually means it's a result of joining. And 44% of non-members have some points, which shouldn't happen. I'll flag that.
2. It confirms what each column means. Column names mislead (is usage_idx measured before or after the training?), so Claude writes down its reading of every column and when it was recorded, and you correct it.
3. It asks what only you know , one question at a time, as clickable options:
Were income, tenure, age, city and last quarter's spend all recorded before customers joined? ○ Yes ○ Some were later ○ Not sure
Is there anything that affects both who joins and how much they spend that isn't in this file? ○ No ○ Yes: ___ ○ Not sure
Could past spending have driven who joined, rather than only the other way round? ○ Yes, big spenders were targeted ○ No ○ Not sure
4. You confirm the diagram and the plan. Claude draws how it thinks things work, you correct it, and the main method is fixed before anything runs. Where you weren't sure, Claude writes down the alternative diagrams too.
5. You get the report about a minute later. It gives the answer, how much to trust it and why, who benefits most, and the randomized test that would settle it.
A self-contained HTML page: it works offline, on a phone, and prints cleanly. It reads top to bottom as a story. Here is the answer; here is what it stood on; here is how sure we are; here is what to do.
When the data can't answer the question , the report leads with that. The diagram shows why, and the page gives what can be said: a range the true effect lies in, and how to find out.
Open the full reports: loyalty program (grade C) · sales calls (grade D) · AI training (grade B: a misleadingly named mediator, a collider and reverse causation). The HTML files, data and every intermediate file are in examples/ .
Claude apps (claude.ai / desktop): download causal-analyst.skill from Releases and upload it in Settings, under Skills. Then attach a data file and ask your question. The skill triggers on questions like "did our loyalty program raise spend?" or "is this cause or just coincidence?"
Claude Code: copy skills/causal-analyst/ into ~/.claude/skills/ (personal) or .claude/skills/ (project).
Python dependencies are installed automatically where the environment allows. Otherwise:
pip install -r requirements.txt # core: numpy, pandas, scikit-learn, statsmodels, econml, dowhy, matplotlib
pip install -r requirements-optional.txt # optional: CausalPFN (local), tabpfn-client (hosted)
Without an agent: the toolkit is a plain CLI, so you can reproduce any report by hand:
cd examples/loyalty-program
S=../../skills/causal-analyst/scripts
python $S /ca.py profile --data data.csv --treatment joined_loyalty --outcome monthly_spend --out profile.json
python $S /ca.py dag spec.json --outdir .
python $S /ca.py identify spec.json --out identification.json
python $S /ca.py run spec.json --out results.json # ~1 minute on 4,000 rows
python $S /ca.py power --sd 31 --lift 5 --results results.json
python $S /ca.py report results.json --narrative narrative.json --out report.html
How it works
flowchart LR
A[Your data + question] --> B[1. Profile + codebook<br/>flags odd values; you confirm<br/>what each column means]
B --> C[2. The question<br/>action, outcome, target,<br/>groups to compare]
C --> D[3. Assumptions interview<br/>timing, assignment, hidden drivers,<br/>reverse causation, colliders]
D --> E[4. Diagram + plan<br/>expert confirms the DAG;<br/>alternatives named; method fixed]
E --> F[5. Run<br/>estimators, diagnostics, design checks,<br/>sensitivity, trust grade]
F --> G[6. Report<br/>designed HTML page]
E -. not identifiable .-> H[Grade D:<br/>bounds, complier effect,<br/>test sizing]
H --> G
Loading
Two rules hold throughout:
Numbers come from the script, never from the model's head. scripts/ca.py produces every estimate. Claude writes only the words ( narrative.json ), and the report renderer draws every chart from results.json .
The plan comes before the results. The main method, controls and target are written into spec.json and approved before anything runs. Other methods are shown as cross-checks, never averaged in.
For amount treatments (e.g. discount size), the skill uses g-computation with the outcome model chosen by cross-validation. For "can't tell" cases it reports Manski / Manski–Pepper bounds, or instrument-based bounds and the complier effect when a random nudge exists.
propensity overlap, with a trimmed estimate;
covariate balance before and after weighting;
a fake-treatment placebo and a random-common-cause test;
stability across 80% subsamples;
the Cinelli–Hazlett robustness value against the strongest measured confounder;
power calculations for a confirming experiment.
Checks on the diagram itself. A correct method on the wrong diagram gives a confident wrong answer, so the run also tests the design:
Column meanings: a codebook of what each column is and when it was recorded, drafted by Claude and confirmed by you. Unconfirmed meanings lower the grade.
Reverse causation: the plan names a before-the-action measure of the outcome (last quarter's score, prior spend) and controls for it, or the report says why not.
Alternative diagrams: the answer is re-estimated under the alternatives you named, with each control dropped in turn, and with each left-out column added. If a plausible alternative moves the answer outside the range, the grade says so.
Structure second opinion: a light PC-algorithm search on the data flags controls that look like consequences (collider patterns) and unused columns linked to both action and outcome. Its findings become questions for you, never silent edits.
Planted-effect test: the main method is rerun on your real columns and real assignment with a simulated outcome carrying a known effect. If it can't find that effect, the grade drops.
Trust grades: A randomized and checks pass · B observational, good overlap, robust to moderate hidden bias · C a weakness (weak overlap, fragile to hidden bias, methods disagree) · D the data can't answer this.
We ran the skill-creator eval loop on three scenarios, comparing Claude with the skill against Claude with the same prompt and no skill ( details ):
Assertion pass rate: 100% with the skill vs 50–57% without , at about 2–3 minutes and ~20% more tokens per run. Accuracy on clean cases is similar either way; on the AI-training case both runs avoided every trap, which tells us current frontier models handle well-described traps. The skill's added value there is the checks and the audit trail: the report shows what the answer would have been under each wrong diagram. The difference is honest ranges, abstention, pre-registration, and a report someone can act on.
The foundation-model cross-checks were benchmarked on 14 semi-synthetic datasets ( results ):
CausalPFN was more accurate, especially under poo

[truncated]

## Original Extract

Causal analysis for people who aren't data scientists: an Agent Skill for Claude - kiritbasu/causal-analyst

I'm not a Data Scientist but have been super intrigued with causal inference off late, so I made a Claude skill to simplify the modelling and experimentation process. You upload your dataset to Claude (Opus) and this skill picks your brain for your domain knowledge about the dataset and comes up with a causal dag that you can iterate through. It then runs the data through a number of classical statistics, ML and newer transformer models to come up with outcomes. And it finally creates a html report to help you understand the results.

GitHub - kiritbasu/causal-analyst: Causal analysis for people who aren't data scientists: an Agent Skill for Claude · GitHub
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
kiritbasu
/
causal-analyst
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.github/ workflows .github/ workflows docs/ images docs/ images evals evals examples examples skills/ causal-analyst skills/ causal-analyst tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md requirements-optional.txt requirements-optional.txt requirements.txt requirements.txt View all files Repository files navigation
Causal analysis for people who aren't data scientists. An Agent Skill that lets Claude answer "did X actually cause Y?" from your data. You bring the question and what you know about your business. Claude does the modelling, checks how far to trust the answer, and hands back a one-page report. When the data can't answer the question, it says so.
Live example reports: Loyalty program (grade C) · Sales calls (grade D, "can't tell") · AI training (grade B, three traps) · Quick start
Contents: Why this exists · What it's like to use · The report · Quick start · How it works · Does it work? · Which model to use · Data and privacy · Limitations and roadmap
Good causal analysis needs two kinds of knowledge that rarely sit in one person:
Data science: which methods to use, how to check them, and what the numbers can and can't support.
Domain knowledge: what happened before what, how people ended up getting the treatment, and what important factor isn't in the data.
The marketing lead knows that "points redeemed" only exists after someone joins, and that the program was pushed to big spenders. The data scientist knows that controlling for points will wreck the estimate, and that targeting big spenders creates a gap that isn't the program's doing. Usually it takes both people and a lot of back and forth. Often the domain expert simply ends up with a correlation.
This skill lets the domain expert work directly with Claude. Claude handles the modelling. It asks, in plain language, only the questions that need domain knowledge, turns the answers into a causal diagram the expert confirms, then runs and checks the analysis and explains the result. The expert supplies what only they know; the skill supplies the rest.
It also guards against the ways capable models go wrong. Frontier models already get the arithmetic right: in our tests plain Claude matched the skill's estimate on a clean case to within a few percent. The failures are judgment calls:
Controlling for the wrong thing. Adjusting for "points redeemed" turns a +$10 effect into −$22.
Answering a question the data can't support. If reps pick whom to call using a gut feel that isn't recorded, no amount of adjustment recovers the effect of a call.
Choosing the method after seeing the results , and quietly drifting toward the hoped-for answer.
Overstating certainty. Plain Claude twice gave confident ranges that missed the true answer.
So the workflow puts those judgments in the open. Controls are confirmed as recorded before the treatment. The expert signs off the diagram. The method is fixed in advance. There's an explicit "we can't tell" grade, and the risk from hidden factors is sized.
You attach a file and ask in your own words. The skill takes it from there, stopping three times for your input.
You: [attaches customers.csv] Did our loyalty program actually raise monthly spend? We think newer customers benefit more.
1. Claude checks the data and points out anything odd before going further:
4,000 customers, 9 columns. points_redeemed almost perfectly predicts who joined, which usually means it's a result of joining. And 44% of non-members have some points, which shouldn't happen. I'll flag that.
2. It confirms what each column means. Column names mislead (is usage_idx measured before or after the training?), so Claude writes down its reading of every column and when it was recorded, and you correct it.
3. It asks what only you know , one question at a time, as clickable options:
Were income, tenure, age, city and last quarter's spend all recorded before customers joined? ○ Yes ○ Some were later ○ Not sure
Is there anything that affects both who joins and how much they spend that isn't in this file? ○ No ○ Yes: ___ ○ Not sure
Could past spending have driven who joined, rather than only the other way round? ○ Yes, big spenders were targeted ○ No ○ Not sure
4. You confirm the diagram and the plan. Claude draws how it thinks things work, you correct it, and the main method is fixed before anything runs. Where you weren't sure, Claude writes down the alternative diagrams too.
5. You get the report about a minute later. It gives the answer, how much to trust it and why, who benefits most, and the randomized test that would settle it.
A self-contained HTML page: it works offline, on a phone, and prints cleanly. It reads top to bottom as a story. Here is the answer; here is what it stood on; here is how sure we are; here is what to do.
When the data can't answer the question , the report leads with that. The diagram shows why, and the page gives what can be said: a range the true effect lies in, and how to find out.
Open the full reports: loyalty program (grade C) · sales calls (grade D) · AI training (grade B: a misleadingly named mediator, a collider and reverse causation). The HTML files, data and every intermediate file are in examples/ .
Claude apps (claude.ai / desktop): download causal-analyst.skill from Releases and upload it in Settings, under Skills. Then attach a data file and ask your question. The skill triggers on questions like "did our loyalty program raise spend?" or "is this cause or just coincidence?"
Claude Code: copy skills/causal-analyst/ into ~/.claude/skills/ (personal) or .claude/skills/ (project).
Python dependencies are installed automatically where the environment allows. Otherwise:
pip install -r requirements.txt # core: numpy, pandas, scikit-learn, statsmodels, econml, dowhy, matplotlib
pip install -r requirements-optional.txt # optional: CausalPFN (local), tabpfn-client (hosted)
Without an agent: the toolkit is a plain CLI, so you can reproduce any report by hand:
cd examples/loyalty-program
S=../../skills/causal-analyst/scripts
python $S /ca.py profile --data data.csv --treatment joined_loyalty --outcome monthly_spend --out profile.json
python $S /ca.py dag spec.json --outdir .
python $S /ca.py identify spec.json --out identification.json
python $S /ca.py run spec.json --out results.json # ~1 minute on 4,000 rows
python $S /ca.py power --sd 31 --lift 5 --results results.json
python $S /ca.py report results.json --narrative narrative.json --out report.html
How it works
flowchart LR
A[Your data + question] --> B[1. Profile + codebook<br/>flags odd values; you confirm<br/>what each column means]
B --> C[2. The question<br/>action, outcome, target,<br/>groups to compare]
C --> D[3. Assumptions interview<br/>timing, assignment, hidden drivers,<br/>reverse causation, colliders]
D --> E[4. Diagram + plan<br/>expert confirms the DAG;<br/>alternatives named; method fixed]
E --> F[5. Run<br/>estimators, diagnostics, design checks,<br/>sensitivity, trust grade]
F --> G[6. Report<br/>designed HTML page]
E -. not identifiable .-> H[Grade D:<br/>bounds, complier effect,<br/>test sizing]
H --> G
Loading
Two rules hold throughout:
Numbers come from the script, never from the model's head. scripts/ca.py produces every estimate. Claude writes only the words ( narrative.json ), and the report renderer draws every chart from results.json .
The plan comes before the results. The main method, controls and target are written into spec.json and approved before anything runs. Other methods are shown as cross-checks, never averaged in.
For amount treatments (e.g. discount size), the skill uses g-computation with the outcome model chosen by cross-validation. For "can't tell" cases it reports Manski / Manski–Pepper bounds, or instrument-based bounds and the complier effect when a random nudge exists.
propensity overlap, with a trimmed estimate;
covariate balance before and after weighting;
a fake-treatment placebo and a random-common-cause test;
stability across 80% subsamples;
the Cinelli–Hazlett robustness value against the strongest measured confounder;
power calculations for a confirming experiment.
Checks on the diagram itself. A correct method on the wrong diagram gives a confident wrong answer, so the run also tests the design:
Column meanings: a codebook of what each column is and when it was recorded, drafted by Claude and confirmed by you. Unconfirmed meanings lower the grade.
Reverse causation: the plan names a before-the-action measure of the outcome (last quarter's score, prior spend) and controls for it, or the report says why not.
Alternative diagrams: the answer is re-estimated under the alternatives you named, with each control dropped in turn, and with each left-out column added. If a plausible alternative moves the answer outside the range, the grade says so.
Structure second opinion: a light PC-algorithm search on the data flags controls that look like consequences (collider patterns) and unused columns linked to both action and outcome. Its findings become questions for you, never silent edits.
Planted-effect test: the main method is rerun on your real columns and real assignment with a simulated outcome carrying a known effect. If it can't find that effect, the grade drops.
Trust grades: A randomized and checks pass · B observational, good overlap, robust to moderate hidden bias · C a weakness (weak overlap, fragile to hidden bias, methods disagree) · D the data can't answer this.
We ran the skill-creator eval loop on three scenarios, comparing Claude with the skill against Claude with the same prompt and no skill ( details ):
Assertion pass rate: 100% with the skill vs 50–57% without , at about 2–3 minutes and ~20% more tokens per run. Accuracy on clean cases is similar either way; on the AI-training case both runs avoided every trap, which tells us current frontier models handle well-described traps. The skill's added value there is the checks and the audit trail: the report shows what the answer would have been under each wrong diagram. The difference is honest ranges, abstention, pre-registration, and a report someone can act on.
The foundation-model cross-checks were benchmarked on 14 semi-synthetic datasets ( results ):
CausalPFN was more accurate, especially under poo

[truncated]
