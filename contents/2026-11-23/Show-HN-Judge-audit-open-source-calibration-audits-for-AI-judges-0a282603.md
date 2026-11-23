---
source: "https://github.com/kunko-ai-labs/judge-audit"
hn_url: "https://news.ycombinator.com/item?id=50620311"
title: "Show HN: Judge-audit – open-source calibration audits for AI judges"
article_title: "GitHub - kunko-ai-labs/judge-audit: Independent calibration audits for AI judges. The Moody's for AI judgment. · GitHub"
image: "https://opengraph.githubassets.com/b059cbcd0c4cbf418297cd160b4ae5548a8c56bac2fd200910962d1e4d729b53/kunko-ai-labs/judge-audit"
author: "kunkoailabs"
captured_at: "2026-11-23T23:28:05Z"
capture_tool: "hn-digest"
hn_id: 50620311
score: 1
comments: 0
posted_at: "2026-11-23T22:28:32Z"
tags:
  - hacker-news
---

# Show HN: Judge-audit – open-source calibration audits for AI judges

- HN: [50620311](https://news.ycombinator.com/item?id=50620311)
- Source: [github.com](https://github.com/kunko-ai-labs/judge-audit)
- Score: 1
- Comments: 0
- Posted: 2026-11-23T22:28:32Z

## Translation

Title: Show HN: Judge-audit – open-source calibration audits for AI judges
Article title: GitHub - kunko-ai-labs/judge-audit: Independent calibration audits for AI judges. The Moody's for AI judgment. · GitHub
Description: Independent calibration audits for AI judges. The Moody's for AI judgment. - kunko-ai-labs/judge-audit

Article text:
GitHub - kunko-ai-labs/judge-audit: Independent calibration audits for AI judges. The Moody's for AI judgment. · GitHub
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
kunko-ai-labs
judge-audit Public Notifications You must be signed in to change notification settings
Star 12 ( 12 ) You must be signed in to star a repository
Independent calibration audits for AI judges. The Moody's for AI judgment.
Security policy Activity Custom properties Stars
1 fork Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
149 Commits 149 Commits Folders and files
.github .github docs docs examples examples scripts scripts src/ judge_audit src/ judge_audit tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Turn AI confidence into a defensible decision boundary
Before an AI judge acts alone, show with evidence where autonomy is earned. judge-audit replays decisions your people already made and tells you which share an AI judge can take over at the error you accept, with the confidence threshold to deploy, a held-out check and the bound attached (approximate, for traffic like the decisions you labelled). Open source, runs in CI, no LLM inside the measurement, and every published number recomputes from committed evidence.
An independent audit, not a model. judge-audit measures the confidence an AI judge declares against decisions people already made. It judges nothing itself.
Evidence anyone can recompute. Every published figure regenerates from committed checkpoints, and CI fails when a report or this README drifts from its evidence ( scripts/verify_published.py , scripts/verify_readme.py ).
Pre-registered studies, separate reviews. The v0.5 and v0.6 studies were pre-registered before their first model call (v0.6: the plan's sha256 committed at the freeze; amendments A1–A4 hashed before the calls or the analysis they govern, A5, an erratum on wording, at publication). Each pull request goes through a separate review and QA pass before it merges, by AI review agents that did not write the change, with the verdicts posted on each pull request. No external human reviewer has read the study plans.
Our own security advisories. We publish the weaknesses we find in our own tools: GHSA-52xj-j738-jpm3 (judge-audit: an API key in error output, or sent on a redirect) and GHSA-f7jx-f5h6-pmx7 (agent-assurance).
Run the quickstart · Audit your own decisions · What you get · v0.6 findings · v0.5 findings · ▶ launch video
v0.6 in short. 12 AI judges, the same 2,974 items of MASSIVE 1.1 in English, Spanish and Catalan, pre-registered before the first model call ( plan , sha256 committed at the freeze; amendments A1–A5). Six pre-registered transfer tests asked whether a threshold validated (approximate) on English rows at a 5 % error target lets more than 5 % errors through on the same items in Spanish or Catalan: none was resolved at its Holm-corrected level, so no breach was shown; that is not a pass: the approximate one-sided 95 % upper bounds on the pooled target-language error run from 3.30 % (T6) to 8.65 % (T2), above 5 % in 5 of 6 tests, and 3 of the 6 came out as predicted. Demonstrated on this sample, and only on it (MASSIVE is translated public data, probably seen in pretraining): declared confidence does not always match accuracy; mean confidence minus accuracy is +0.080 to +0.085 for jev (Jev 1.13.0), −0.052 to −0.067 for gpt-6-luna and +0.063 to +0.099 for claude-sonnet-4.5 across the three languages. The differences between languages are 95 % intervals with no correction for multiplicity, so they are not called demonstrated ( §2 ). The number to act on is the same audit on your own decisions.
Jev version. Every Jev figure of v0.6 was measured with Jev 1.13.0, the model the replies say they served. The alias jev-latest has pointed to jev-1.14.0 since 2026-10-22, so a new run under the alias measures a different model and its figures are not comparable with these. Pin the version with JEV_MODEL (for example JEV_MODEL=jev-1.13.0 ; how ).
At a glance: v0.6, MASSIVE 1.1, 2,974 translated public items in English, Spanish and Catalan (not your traffic), 12 judges: the held-out share automated by thresholds chosen at a 5 % error target (the threshold chosen on one half of the items, applied to the other), with its approximate bound; met when that bound is within 5 %, at the study's split (seed 2026).
Met is per split. gpt-6-luna en-US is met at the study's split (seed 2026), and its share moves over split seeds 2027–2036 (21.8–39.4 %); on another split a met cell can miss the target. The bounds of the other splits are not computed here.
Bound: the one-sided 95 % upper bound on the held-out errors of both halves pooled, rounded up; approximate, since each half's threshold depends on the other half; computed after the study from its published counts, not pre-registered (#213). Met: that bound is within 5 %; only a share that meets it is called validated. Each met cell is a separate statement, each at 95 %, approximate; read together they do not all hold at 95 %, so no claim is made about all of them at once, and none should be picked as the best. None: no threshold was validated on either half at seed 2026. Splits: the same share over split seeds 2027–2036; that range need not contain the seed-2026 share, and a share far outside it rests on one split. Labels are carried over by localisation, and no share says anything about future traffic; judges marked exploratory are in no pre-registered test. No ranking: rows follow the study's roster. The shares, their ranges, every other figure with its intervals and the cost of each run are in the v0.6 findings in full ; the bounds are computed here from their counts.
In short. Asking Qwen3-8B and gemini-3.6-flash several times and counting the votes ranked their errors worse than their own verbalized number in 3 of 4 tests under the pre-registered rule (4 of 4 under the re-reading), and one verdict (H2 for gemini-3.6-flash) changes with how an answer that copies an option's description is scored. On BANKING77 at a 5 % error target, validated on held-out texts (approximate), Jev native probability can decide 41.0 % (27.0 %–45.1 % over split seeds 2027–2036) of the texts alone, gemini-3.6-flash verbalized 14.3 % (7.6 %–30.0 %), Qwen3-8B token log-probability none (0.0 %–4.4 %). Whether the bound is within the target depends on the split: at split seed 2026 Jev's 41.0 % has a bound of 3.9 %, within 5 %; on another split of the same checkpoints (split seeds 2033 and 2036) the same judge at the same 5 % target automates fewer texts with a bound above 5 %, not met; the range over split seeds shows how far the share moves. Caveats: public datasets, probably seen in pretraining; one run each, except Jev and Qwen3-8B verbalized on BANKING77 (3 runs); label noise not measured; the held-out slice not run (#106). The number to act on is the same audit on your own decisions.
At a glance: BANKING77, 3,080 human-labelled banking queries: the share each judge decides alone at a 5 % error target, validated on held-out texts (approximate) at split seed 2026; another split can change the result.
Validated: the threshold is chosen on one half of the texts and applied to the other, both ways, and the 95 % upper bound on the held-out error rate is within the target. The bound is the one-sided Clopper–Pearson bound on the held-out errors of both halves pooled, rounded up; it treats the two held-out halves as independent draws, an approximation since each half's threshold depends on the other half. A share observed on all rows without a bound is not a basis for deployment.
Jev version (v0.5). The Jev rows of v0.5 were also measured with Jev 1.13.0: every one of the 3,080 BANKING77 replies in each of the three Jev runs says it was served by jev-1.13.0 . Figures of another Jev version are not comparable with them.
Teams are shipping judgment models — TypeSafe's Jev, LLM-as-judge, guardrails, routers — that return a confidence with each decision. The literature studies calibration (Guo et al. 2017; Shao 2026; Huang et al. 2026); what a team needs before automating is the same measurement on its own decisions. judge-audit runs any judge in shadow mode against decisions your humans already made and answers the four questions that matter before you automate:
The study's design, the eight confirmatory tests and every caveat: v0.5 findings in full .
pip install kunko-judge-audit # every release ships Sigstore-signed build provenance
git clone --branch v0.6.3 https://github.com/kunko-ai-labs/judge-audit && cd judge-audit # the example data lives in the repo, not in the pip package
judge-audit run examples/email-routing/labels.jsonl --judge simulated --target 0.10 # no API key needed
judge-audit check examples/email-routing/labels.jsonl --judge simulated --baseline audit-result.json --target 0.10 --min-safe-rate 0.10:0.60
The start of the run line, SIMULATED :
SIMULATED — not a real vendor audit · judge=simulated n=200 accuracy=85.5% [75.1, 93.7] safe_automation@10%=not_met(75.0% bound_95_heldout_approx=16.6%) all_rows=65.5% …
simulated is a seeded simulator so you can see the whole pipeline in ten seconds; every report it touches is stamped SIMULATED , and its figures say nothing about any real judge. safe_automation@10% leads with the held-out check at a 10 % error target: not_met , because the one-sided 95 % upper bound on the held-out error rate (16.6 %, approximate: it pools the two held-out halves) is above the target, so the 75.0 % held out is not validated. A share is shown as validated only when that bound is within the target. all_rows (65.5 %) is the share the threshold to deploy covers on every labelled decision, the share check --min-safe-rate gates. At the default target, 5 %, the simulator reaches no threshold and the line reads safe_automation@5%=none . check --min-safe-rate 0.10:0.60 exits 0 here and exits 1 if that share falls below 60 %. To audit a real vendor, see docs/real-audits.md .
In the GitHub Action the same controls are inputs: target (the error rate the safe automation rate is headlined at, default 0.05 ), segment-by (where to look for the worst segment: label , meta.FIELD or none ) and min-safe-rate ( RISK:SHARE minimums that fail the job). It returns safe-automation-rate (at the target) and safe-automation-rates (every target) as outputs.
One self-contained report per run (Markdown or HTML, prints to A4): the answer first: th

[truncated]

## Original Extract

Independent calibration audits for AI judges. The Moody's for AI judgment. - kunko-ai-labs/judge-audit

GitHub - kunko-ai-labs/judge-audit: Independent calibration audits for AI judges. The Moody's for AI judgment. · GitHub
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
kunko-ai-labs
judge-audit Public Notifications You must be signed in to change notification settings
Star 12 ( 12 ) You must be signed in to star a repository
Independent calibration audits for AI judges. The Moody's for AI judgment.
Security policy Activity Custom properties Stars
1 fork Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
149 Commits 149 Commits Folders and files
.github .github docs docs examples examples scripts scripts src/ judge_audit src/ judge_audit tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Turn AI confidence into a defensible decision boundary
Before an AI judge acts alone, show with evidence where autonomy is earned. judge-audit replays decisions your people already made and tells you which share an AI judge can take over at the error you accept, with the confidence threshold to deploy, a held-out check and the bound attached (approximate, for traffic like the decisions you labelled). Open source, runs in CI, no LLM inside the measurement, and every published number recomputes from committed evidence.
An independent audit, not a model. judge-audit measures the confidence an AI judge declares against decisions people already made. It judges nothing itself.
Evidence anyone can recompute. Every published figure regenerates from committed checkpoints, and CI fails when a report or this README drifts from its evidence ( scripts/verify_published.py , scripts/verify_readme.py ).
Pre-registered studies, separate reviews. The v0.5 and v0.6 studies were pre-registered before their first model call (v0.6: the plan's sha256 committed at the freeze; amendments A1–A4 hashed before the calls or the analysis they govern, A5, an erratum on wording, at publication). Each pull request goes through a separate review and QA pass before it merges, by AI review agents that did not write the change, with the verdicts posted on each pull request. No external human reviewer has read the study plans.
Our own security advisories. We publish the weaknesses we find in our own tools: GHSA-52xj-j738-jpm3 (judge-audit: an API key in error output, or sent on a redirect) and GHSA-f7jx-f5h6-pmx7 (agent-assurance).
Run the quickstart · Audit your own decisions · What you get · v0.6 findings · v0.5 findings · ▶ launch video
v0.6 in short. 12 AI judges, the same 2,974 items of MASSIVE 1.1 in English, Spanish and Catalan, pre-registered before the first model call ( plan , sha256 committed at the freeze; amendments A1–A5). Six pre-registered transfer tests asked whether a threshold validated (approximate) on English rows at a 5 % error target lets more than 5 % errors through on the same items in Spanish or Catalan: none was resolved at its Holm-corrected level, so no breach was shown; that is not a pass: the approximate one-sided 95 % upper bounds on the pooled target-language error run from 3.30 % (T6) to 8.65 % (T2), above 5 % in 5 of 6 tests, and 3 of the 6 came out as predicted. Demonstrated on this sample, and only on it (MASSIVE is translated public data, probably seen in pretraining): declared confidence does not always match accuracy; mean confidence minus accuracy is +0.080 to +0.085 for jev (Jev 1.13.0), −0.052 to −0.067 for gpt-6-luna and +0.063 to +0.099 for claude-sonnet-4.5 across the three languages. The differences between languages are 95 % intervals with no correction for multiplicity, so they are not called demonstrated ( §2 ). The number to act on is the same audit on your own decisions.
Jev version. Every Jev figure of v0.6 was measured with Jev 1.13.0, the model the replies say they served. The alias jev-latest has pointed to jev-1.14.0 since 2026-10-22, so a new run under the alias measures a different model and its figures are not comparable with these. Pin the version with JEV_MODEL (for example JEV_MODEL=jev-1.13.0 ; how ).
At a glance: v0.6, MASSIVE 1.1, 2,974 translated public items in English, Spanish and Catalan (not your traffic), 12 judges: the held-out share automated by thresholds chosen at a 5 % error target (the threshold chosen on one half of the items, applied to the other), with its approximate bound; met when that bound is within 5 %, at the study's split (seed 2026).
Met is per split. gpt-6-luna en-US is met at the study's split (seed 2026), and its share moves over split seeds 2027–2036 (21.8–39.4 %); on another split a met cell can miss the target. The bounds of the other splits are not computed here.
Bound: the one-sided 95 % upper bound on the held-out errors of both halves pooled, rounded up; approximate, since each half's threshold depends on the other half; computed after the study from its published counts, not pre-registered (#213). Met: that bound is within 5 %; only a share that meets it is called validated. Each met cell is a separate statement, each at 95 %, approximate; read together they do not all hold at 95 %, so no claim is made about all of them at once, and none should be picked as the best. None: no threshold was validated on either half at seed 2026. Splits: the same share over split seeds 2027–2036; that range need not contain the seed-2026 share, and a share far outside it rests on one split. Labels are carried over by localisation, and no share says anything about future traffic; judges marked exploratory are in no pre-registered test. No ranking: rows follow the study's roster. The shares, their ranges, every other figure with its intervals and the cost of each run are in the v0.6 findings in full ; the bounds are computed here from their counts.
In short. Asking Qwen3-8B and gemini-3.6-flash several times and counting the votes ranked their errors worse than their own verbalized number in 3 of 4 tests under the pre-registered rule (4 of 4 under the re-reading), and one verdict (H2 for gemini-3.6-flash) changes with how an answer that copies an option's description is scored. On BANKING77 at a 5 % error target, validated on held-out texts (approximate), Jev native probability can decide 41.0 % (27.0 %–45.1 % over split seeds 2027–2036) of the texts alone, gemini-3.6-flash verbalized 14.3 % (7.6 %–30.0 %), Qwen3-8B token log-probability none (0.0 %–4.4 %). Whether the bound is within the target depends on the split: at split seed 2026 Jev's 41.0 % has a bound of 3.9 %, within 5 %; on another split of the same checkpoints (split seeds 2033 and 2036) the same judge at the same 5 % target automates fewer texts with a bound above 5 %, not met; the range over split seeds shows how far the share moves. Caveats: public datasets, probably seen in pretraining; one run each, except Jev and Qwen3-8B verbalized on BANKING77 (3 runs); label noise not measured; the held-out slice not run (#106). The number to act on is the same audit on your own decisions.
At a glance: BANKING77, 3,080 human-labelled banking queries: the share each judge decides alone at a 5 % error target, validated on held-out texts (approximate) at split seed 2026; another split can change the result.
Validated: the threshold is chosen on one half of the texts and applied to the other, both ways, and the 95 % upper bound on the held-out error rate is within the target. The bound is the one-sided Clopper–Pearson bound on the held-out errors of both halves pooled, rounded up; it treats the two held-out halves as independent draws, an approximation since each half's threshold depends on the other half. A share observed on all rows without a bound is not a basis for deployment.
Jev version (v0.5). The Jev rows of v0.5 were also measured with Jev 1.13.0: every one of the 3,080 BANKING77 replies in each of the three Jev runs says it was served by jev-1.13.0 . Figures of another Jev version are not comparable with them.
Teams are shipping judgment models — TypeSafe's Jev, LLM-as-judge, guardrails, routers — that return a confidence with each decision. The literature studies calibration (Guo et al. 2017; Shao 2026; Huang et al. 2026); what a team needs before automating is the same measurement on its own decisions. judge-audit runs any judge in shadow mode against decisions your humans already made and answers the four questions that matter before you automate:
The study's design, the eight confirmatory tests and every caveat: v0.5 findings in full .
pip install kunko-judge-audit # every release ships Sigstore-signed build provenance
git clone --branch v0.6.3 https://github.com/kunko-ai-labs/judge-audit && cd judge-audit # the example data lives in the repo, not in the pip package
judge-audit run examples/email-routing/labels.jsonl --judge simulated --target 0.10 # no API key needed
judge-audit check examples/email-routing/labels.jsonl --judge simulated --baseline audit-result.json --target 0.10 --min-safe-rate 0.10:0.60
The start of the run line, SIMULATED :
SIMULATED — not a real vendor audit · judge=simulated n=200 accuracy=85.5% [75.1, 93.7] safe_automation@10%=not_met(75.0% bound_95_heldout_approx=16.6%) all_rows=65.5% …
simulated is a seeded simulator so you can see the whole pipeline in ten seconds; every report it touches is stamped SIMULATED , and its figures say nothing about any real judge. safe_automation@10% leads with the held-out check at a 10 % error target: not_met , because the one-sided 95 % upper bound on the held-out error rate (16.6 %, approximate: it pools the two held-out halves) is above the target, so the 75.0 % held out is not validated. A share is shown as validated only when that bound is within the target. all_rows (65.5 %) is the share the threshold to deploy covers on every labelled decision, the share check --min-safe-rate gates. At the default target, 5 %, the simulator reaches no threshold and the line reads safe_automation@5%=none . check --min-safe-rate 0.10:0.60 exits 0 here and exits 1 if that share falls below 60 %. To audit a real vendor, see docs/real-audits.md .
In the GitHub Action the same controls are inputs: target (the error rate the safe automation rate is headlined at, default 0.05 ), segment-by (where to look for the worst segment: label , meta.FIELD or none ) and min-safe-rate ( RISK:SHARE minimums that fail the job). It returns safe-automation-rate (at the target) and safe-automation-rates (every target) as outputs.
One self-contained report per run (Markdown or HTML, prints to A4): the answer first: th

[truncated]
