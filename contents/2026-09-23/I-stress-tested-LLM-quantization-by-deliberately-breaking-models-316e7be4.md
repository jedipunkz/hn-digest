---
source: "https://github.com/gracejackson-sudo/quant-delta-predictor"
hn_url: "https://news.ycombinator.com/item?id=49823211"
title: "I stress-tested LLM quantization by deliberately breaking models"
article_title: "GitHub - gracejackson-sudo/quant-delta-predictor: A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold · GitHub"
image: "https://opengraph.githubassets.com/dc8e23b1642ed7f96322cfbe6bd22387bb66c4217d98069392f07b1aaf013d85/gracejackson-sudo/quant-delta-predictor"
author: "GraceJackson07"
captured_at: "2026-09-23T22:44:06Z"
capture_tool: "hn-digest"
hn_id: 49823211
score: 1
comments: 0
posted_at: "2026-09-23T22:04:50Z"
tags:
  - hacker-news
---

# I stress-tested LLM quantization by deliberately breaking models

- HN: [49823211](https://news.ycombinator.com/item?id=49823211)
- Source: [github.com](https://github.com/gracejackson-sudo/quant-delta-predictor)
- Score: 1
- Comments: 0
- Posted: 2026-09-23T22:04:50Z

## Translation

Title: I stress-tested LLM quantization by deliberately breaking models
Article title: GitHub - gracejackson-sudo/quant-delta-predictor: A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold · GitHub
Description: A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold - gracejackson-sudo/quant-delta-predictor

Article text:
GitHub - gracejackson-sudo/quant-delta-predictor: A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold · GitHub
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
gracejackson-sudo
/
quant-delta-predictor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
data data gpu gpu out out paper paper src src tests tests verify verify .gitignore .gitignore ACCOUNTING.md ACCOUNTING.md ADVERSARIAL_AUDIT.md ADVERSARIAL_AUDIT.md AUDIT_DISCIPLINE.md AUDIT_DISCIPLINE.md BIAS_CORRECTION.md BIAS_CORRECTION.md FINDINGS.md FINDINGS.md LICENSE LICENSE NEGATIVE_RESULT.md NEGATIVE_RESULT.md PROVENANCE.md PROVENANCE.md RANKING.md RANKING.md README.md README.md RESEARCH.md RESEARCH.md SCOPE.md SCOPE.md TOOL_SUMMARY.md TOOL_SUMMARY.md feedback_config.json feedback_config.json feedback_form.html feedback_form.html index.html index.html View all files Repository files navigation
Feasibility spike: given (base model, quantization config) , predict the accuracy delta on
OpenLLM-style benchmarks with a calibrated prediction interval.
Result: the calibration works (90.1% empirical coverage on unseen checkpoints at a nominal 90%);
the point prediction carries almost no signal beyond the quantization scheme, and the fitted
artifact is a small number table — a calibrated historical baseline, not a predictor.
python3 -m venv .venv && ./.venv/bin/pip install numpy pandas scipy
# scikit-learn and pytest are only needed to re-run the research, not the tool
Use the tool
No network, no API key, no GPU. Three packages: numpy, pandas, scipy.
Run this one first. It is the command that shows you what the tool is:
./.venv/bin/python src/rank.py w4a16 --size 1.5B
It refuses to answer. For 4-bit weights on a sub-2B model, measured coverage is
68.8% against the 90% the interval claims, so instead of returning a number it
prints INSUFFICIENT CALIBRATION FOR THIS COMBINATION , shows you the coverage it
actually measured and the rows it measured it on, and stops.
That is the design. A quantization risk estimate is only worth having if it will
tell you when not to trust it, and the cells it refuses are exactly the ones where
a confident-sounding answer would do the most damage.
Once you have seen it decline, the rest:
./.venv/bin/python src/rank.py # rank every scheme
./.venv/bin/python src/rank.py fp8 # one scheme it will answer for
./.venv/bin/python src/rank.py --form-fields # what a contributed result needs
Each scheme comes back with a 90% interval, the worst loss ever observed for it, the
share of evaluations that lost more than 3pp, how many checkpoints and families back
it, and a tier. Two cells are refused outright and four more are blocked from the top
tier for resting on fewer than three distinct checkpoints.
The research scripts need two more packages than the tool does:
./.venv/bin/pip install scikit-learn pytest
./.venv/bin/python src/harvest.py # rebuild dataset.csv from cards (needs the cards; see data/README.md)
./.venv/bin/python src/run_final.py # all three split regimes and calibration variants
./.venv/bin/python src/demo_holdout.py # train on 2 families, predict unseen models
./.venv/bin/python src/real_use_case.py # prospective test on never-seen families (needs network)
./.venv/bin/python src/diagnose.py # which failure mode this is
./.venv/bin/python src/validate_strata.py # does size stratification help? mostly not - see RANKING.md
./.venv/bin/python src/build_envelope.py && ./.venv/bin/python src/cli.py --list
./.venv/bin/python -m pytest tests -q # the test suite
The audits, and the full exclusion census:
./.venv/bin/python src/audit.py && ./.venv/bin/python src/adversarial_audit.py && ./.venv/bin/python src/census.py
Independent re-verification — stdlib only, no project imports. Regenerates the headline coverage
from the raw cards and diffs it against the pipeline row by row:
python3 verify/independent_check.py
Data
850 rows of (model, quant_config, benchmark, accuracy_before, accuracy_after) scraped from 102
RedHatAI model cards, spanning 38 base checkpoints, 8 model
families, 6 quantization schemes and 16 benchmarks. Every three-column row is verified against
the card's own printed Recovery percentage; internally inconsistent rows are rejected rather than
guessed at.
It does not use the model family or size — those features made out-of-family accuracy worse .
Its intervals never exclude zero, so it cannot tell you a config will definitely hurt.
It is trained only on checkpoints Red Hat chose to publish, so it underpredicts damage from a
badly-tuned recipe.
Sub-2B, MoE and reasoning-distilled models fall outside the validated envelope.
Every quantitative statement in RANKING.md must trace to a value computed by the audit scripts.
This is enforced mechanically, not by discipline:
RANKING.md is generated by src/gen_ranking_doc.py from computed values. Do not hand-edit
figures in it.
Each figure carries a <!-- claim: key = value --> tag (invisible when rendered).
src/verify_claims.py recomputes every tagged value and fails on any mismatch or any tag not in
its registry.
tests/test_all.py::test_ranking_doc_claims_all_verify runs that check in CI.
tests/test_all.py::test_no_banned_prose_in_user_facing_output blocks phrases that were wrong
before ("safe to adopt", "ignores your model", "12 numbers", "26 of 29", "80%, not 90%") from
reappearing in printed output.
./.venv/bin/python src/cell_coverage.py # measure per-cell coverage
./.venv/bin/python src/gen_ranking_doc.py # regenerate RANKING.md
./.venv/bin/python src/verify_claims.py # re-verify every number
Where a (scheme, size band) cell has measured coverage below threshold, the tool prints
INSUFFICIENT CALIBRATION and no interval , rather than a number that would look as confident
as a well-calibrated one.
This repository is MIT licensed (see LICENSE ). The raw Hugging Face model
cards the dataset was extracted from are not redistributed here — they are
RedHatAI's content under several different licences. data/dataset.csv (the
extracted numbers) is included and is all the tool needs. See
data/README.md for what that costs you and how to fetch the
cards yourself.
Source data: evaluation results published by RedHatAI on Hugging Face —
https://huggingface.co/RedHatAI .
Frantar, Ashkboos, Hoefler & Alistarh. GPTQ: Accurate Post-Training
Quantization for Generative Pre-trained Transformers.
arXiv:2210.17323 — the method behind
every w4a16 checkpoint in this dataset.
Xiao et al. SmoothQuant.
arXiv:2211.10438 — the activation
smoothing used for the w8a8 checkpoints.
Zeng & Papailiopoulos. You Don't Need to Run Every Eval (BenchPress).
arXiv:2606.24020 — published the core
predict-then-conformal mechanism first; see NEGATIVE_RESULT.md §5.
Barber, Candès, Ramdas & Tibshirani. Conformal Prediction Beyond
Exchangeability. Ann. Statist. 51(2), 2023 — why the coverage guarantee is
conditional on the published population, not on your own recipe.
A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold - gracejackson-sudo/quant-delta-predictor

GitHub - gracejackson-sudo/quant-delta-predictor: A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold · GitHub
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
gracejackson-sudo
/
quant-delta-predictor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
data data gpu gpu out out paper paper src src tests tests verify verify .gitignore .gitignore ACCOUNTING.md ACCOUNTING.md ADVERSARIAL_AUDIT.md ADVERSARIAL_AUDIT.md AUDIT_DISCIPLINE.md AUDIT_DISCIPLINE.md BIAS_CORRECTION.md BIAS_CORRECTION.md FINDINGS.md FINDINGS.md LICENSE LICENSE NEGATIVE_RESULT.md NEGATIVE_RESULT.md PROVENANCE.md PROVENANCE.md RANKING.md RANKING.md README.md README.md RESEARCH.md RESEARCH.md SCOPE.md SCOPE.md TOOL_SUMMARY.md TOOL_SUMMARY.md feedback_config.json feedback_config.json feedback_form.html feedback_form.html index.html index.html View all files Repository files navigation
Feasibility spike: given (base model, quantization config) , predict the accuracy delta on
OpenLLM-style benchmarks with a calibrated prediction interval.
Result: the calibration works (90.1% empirical coverage on unseen checkpoints at a nominal 90%);
the point prediction carries almost no signal beyond the quantization scheme, and the fitted
artifact is a small number table — a calibrated historical baseline, not a predictor.
python3 -m venv .venv && ./.venv/bin/pip install numpy pandas scipy
# scikit-learn and pytest are only needed to re-run the research, not the tool
Use the tool
No network, no API key, no GPU. Three packages: numpy, pandas, scipy.
Run this one first. It is the command that shows you what the tool is:
./.venv/bin/python src/rank.py w4a16 --size 1.5B
It refuses to answer. For 4-bit weights on a sub-2B model, measured coverage is
68.8% against the 90% the interval claims, so instead of returning a number it
prints INSUFFICIENT CALIBRATION FOR THIS COMBINATION , shows you the coverage it
actually measured and the rows it measured it on, and stops.
That is the design. A quantization risk estimate is only worth having if it will
tell you when not to trust it, and the cells it refuses are exactly the ones where
a confident-sounding answer would do the most damage.
Once you have seen it decline, the rest:
./.venv/bin/python src/rank.py # rank every scheme
./.venv/bin/python src/rank.py fp8 # one scheme it will answer for
./.venv/bin/python src/rank.py --form-fields # what a contributed result needs
Each scheme comes back with a 90% interval, the worst loss ever observed for it, the
share of evaluations that lost more than 3pp, how many checkpoints and families back
it, and a tier. Two cells are refused outright and four more are blocked from the top
tier for resting on fewer than three distinct checkpoints.
The research scripts need two more packages than the tool does:
./.venv/bin/pip install scikit-learn pytest
./.venv/bin/python src/harvest.py # rebuild dataset.csv from cards (needs the cards; see data/README.md)
./.venv/bin/python src/run_final.py # all three split regimes and calibration variants
./.venv/bin/python src/demo_holdout.py # train on 2 families, predict unseen models
./.venv/bin/python src/real_use_case.py # prospective test on never-seen families (needs network)
./.venv/bin/python src/diagnose.py # which failure mode this is
./.venv/bin/python src/validate_strata.py # does size stratification help? mostly not - see RANKING.md
./.venv/bin/python src/build_envelope.py && ./.venv/bin/python src/cli.py --list
./.venv/bin/python -m pytest tests -q # the test suite
The audits, and the full exclusion census:
./.venv/bin/python src/audit.py && ./.venv/bin/python src/adversarial_audit.py && ./.venv/bin/python src/census.py
Independent re-verification — stdlib only, no project imports. Regenerates the headline coverage
from the raw cards and diffs it against the pipeline row by row:
python3 verify/independent_check.py
Data
850 rows of (model, quant_config, benchmark, accuracy_before, accuracy_after) scraped from 102
RedHatAI model cards, spanning 38 base checkpoints, 8 model
families, 6 quantization schemes and 16 benchmarks. Every three-column row is verified against
the card's own printed Recovery percentage; internally inconsistent rows are rejected rather than
guessed at.
It does not use the model family or size — those features made out-of-family accuracy worse .
Its intervals never exclude zero, so it cannot tell you a config will definitely hurt.
It is trained only on checkpoints Red Hat chose to publish, so it underpredicts damage from a
badly-tuned recipe.
Sub-2B, MoE and reasoning-distilled models fall outside the validated envelope.
Every quantitative statement in RANKING.md must trace to a value computed by the audit scripts.
This is enforced mechanically, not by discipline:
RANKING.md is generated by src/gen_ranking_doc.py from computed values. Do not hand-edit
figures in it.
Each figure carries a <!-- claim: key = value --> tag (invisible when rendered).
src/verify_claims.py recomputes every tagged value and fails on any mismatch or any tag not in
its registry.
tests/test_all.py::test_ranking_doc_claims_all_verify runs that check in CI.
tests/test_all.py::test_no_banned_prose_in_user_facing_output blocks phrases that were wrong
before ("safe to adopt", "ignores your model", "12 numbers", "26 of 29", "80%, not 90%") from
reappearing in printed output.
./.venv/bin/python src/cell_coverage.py # measure per-cell coverage
./.venv/bin/python src/gen_ranking_doc.py # regenerate RANKING.md
./.venv/bin/python src/verify_claims.py # re-verify every number
Where a (scheme, size band) cell has measured coverage below threshold, the tool prints
INSUFFICIENT CALIBRATION and no interval , rather than a number that would look as confident
as a well-calibrated one.
This repository is MIT licensed (see LICENSE ). The raw Hugging Face model
cards the dataset was extracted from are not redistributed here — they are
RedHatAI's content under several different licences. data/dataset.csv (the
extracted numbers) is included and is all the tool needs. See
data/README.md for what that costs you and how to fetch the
cards yourself.
Source data: evaluation results published by RedHatAI on Hugging Face —
https://huggingface.co/RedHatAI .
Frantar, Ashkboos, Hoefler & Alistarh. GPTQ: Accurate Post-Training
Quantization for Generative Pre-trained Transformers.
arXiv:2210.17323 — the method behind
every w4a16 checkpoint in this dataset.
Xiao et al. SmoothQuant.
arXiv:2211.10438 — the activation
smoothing used for the w8a8 checkpoints.
Zeng & Papailiopoulos. You Don't Need to Run Every Eval (BenchPress).
arXiv:2606.24020 — published the core
predict-then-conformal mechanism first; see NEGATIVE_RESULT.md §5.
Barber, Candès, Ramdas & Tibshirani. Conformal Prediction Beyond
Exchangeability. Ann. Statist. 51(2), 2023 — why the coverage guarantee is
conditional on the published population, not on your own recipe.
A calibrated risk envelope for LLM quantization that refuses to answer where its own coverage doesn't hold
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
