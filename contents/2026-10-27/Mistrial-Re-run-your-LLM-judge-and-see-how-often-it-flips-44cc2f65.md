---
source: "https://github.com/hemanth1999k/mistrial"
hn_url: "https://news.ycombinator.com/item?id=50253334"
title: "Mistrial – Re run your LLM judge and see how often it flips"
article_title: "GitHub - hemanth1999k/mistrial: Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card. · GitHub"
image: "https://repository-images.githubusercontent.com/1433955083/d40771bc-3826-4933-85e1-6ca470e783f3"
author: "hemanth2k"
captured_at: "2026-10-27T01:12:48Z"
capture_tool: "hn-digest"
hn_id: 50253334
score: 1
comments: 0
posted_at: "2026-10-27T00:46:14Z"
tags:
  - hacker-news
---

# Mistrial – Re run your LLM judge and see how often it flips

- HN: [50253334](https://news.ycombinator.com/item?id=50253334)
- Source: [github.com](https://github.com/hemanth1999k/mistrial)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T00:46:14Z

## Translation

Title: Mistrial – Re run your LLM judge and see how often it flips
Article title: GitHub - hemanth1999k/mistrial: Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card. · GitHub
Description: Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card. - hemanth1999k/mistrial

Article text:
GitHub - hemanth1999k/mistrial: Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card. · GitHub
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
hemanth1999k
mistrial Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card.
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
38 Commits 38 Commits Folders and files
.github .github docs docs examples examples src/ mistrial src/ mistrial tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Is your LLM judge flaky?
Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card.
When the verdict can't be trusted, declare a mistrial.
pip install mistrial
export OPENAI_API_KEY=... # or ANTHROPIC_API_KEY, GEMINI_API_KEY, GROQ_API_KEY ... see `mistrial providers`
mistrial run cases.jsonl -j openai:gpt-5.4-mini -j openai:gpt-5.5 --runs 5
No key? pipx run mistrial demo runs three fake judges that are built to be flaky on purpose.
You ship AI features based on a score from an LLM judge. The score gates CI, decides which prompt wins, and goes on the dashboard. But the judge is a model sampling from a distribution. Ask it the same question twice and it changes its verdict on some of your cases. Swap the model and it changes its mind on more. Reword the rubric without changing its meaning and it moves again. One eval run is one sample from a noisy process, and almost nobody measures the noise.
These are real replies. On the first case, gpt-5.4-nano wrote "the added mention of 1811 does not contradict the core requirement" on one run and "1811 is not part of the reference, so the answer is incorrect" on the next. On the second, it worked out that 239,000 miles is 384,000 km and then failed the answer for being in miles. Each of those judges reports a clean score on its own. The flakiness only shows up when you ask again and compare.
The research agrees. The same judge gives different scores to the same answer across runs ( Rating Roulette, EMNLP 2025 ). Raw percent agreement overstates reliability by 30 to 40 points compared with chance-corrected kappa, across 21 judges ( Reliability without Validity, 2026 ).
Eval platforms report one number and stop. LangSmith, Braintrust, Langfuse and friends run the judge once per case. Some can repeat an experiment and show a standard deviation, none report self-agreement, kappa between judges, or which cases flipped.
The statistics exist, but not as a tool. Test-retest reliability, Cohen's kappa and Krippendorff's alpha are a century old and standard in every field that uses human raters. They live in research notebooks, not in something you can run against your evaluator in a minute.
The projects that tried stop at the card. judgecheck, judgebench, Assay rejudge and judge-drift-sentinel each cover a piece of this well, and docs/COMPARISON.md credits what each does better. None of them take your existing evaluator function, run in CI, or tell you what to do when the judge is flaky.
Before "how good is my model?" you need "can I trust the number that told me?"
Point it at the evaluator you already use: a model name, or the function you run in LangSmith or openevals. It asks the judge several times, across one or more judge models, and under rewordings of the rubric, and reports the spread: self-agreement, flipped cases, kappa and alpha between judges, accuracy against your human labels, and the exact cases that move. Then it does something about it: a gate for CI, a stabilized judge with the same signature, and a list of the cases a human should look at.
The 90 second version. Every number in it is from the runs below.
Three current OpenAI judges on examples/tricky.jsonl , 32 cases built to be hard to judge (hedges, rounding, unit conversions, extra details, equivalent code), five runs each. 480 live calls on 25 October 2026. Every number below is from that run, and the GIF replays those replies.
mistrial · 32 cases · 3 judges · 5 runs each · 480 verdicts in 76.66s
Consistency (same judge, asked again)
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━┓
┃ judge ┃ self-agreement ┃ flipped cases ┃ vs human ┃ errors ┃ reading ┃
┡━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━━━━━━━━━┩
│ openai:gpt-5.4-nano │ 92% │ 5/32 (16%) │ 81% (κ 0.63) │ 0 │ noticeably flaky │
│ openai:gpt-5.4-mini │ 95% │ 4/32 (12%) │ 88% (κ 0.75) │ 0 │ mostly stable │
│ openai:gpt-5.5 │ 100% │ 0/32 (0%) │ 91% (κ 0.81) │ 0 │ stable │
└─────────────────────┴────────────────┴───────────────┴───────────────┴────────┴──────────────────┘
Flakiest cases (everyone's verdicts pooled)
┏━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━┓
┃ case ┃ disagreement ┃ openai:gpt-5.4-nano ┃ openai:gpt-5.4-mini ┃ openai:gpt-5.5 ┃ human ┃
┡━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━╇━━━━━━━┩
│ arigato │ 40% │ F F F F F │ F T T T T │ T T T T T │ T │
│ austen-wrong-year │ 40% │ T F F F T │ T T T T F │ F F F F F │ F │
│ largest-country-hedge │ 40% │ F F F F F │ T F F F F │ T T T T T │ F │
│ moon-miles │ 33% │ F F F F F │ F F F F F │ T T T T T │ T │
│ percent-unit │ 33% │ T T T T T │ F F F F F │ F F F F F │ F │
└───────────────────────┴──────────────┴─────────────────────┴─────────────────────┴────────────────┴───────┘
What the run says, in order:
gpt-5.4-nano flips on 5 of 32 cases with nothing changed, and agrees with the human label on 81%. gpt-5.5 never flips and agrees on 91%. The judges agree with each other at kappa 0.69 to 0.81, Krippendorff's alpha 0.75.
A CI gate at a 0.5 pass mark on gpt-5.4-mini goes red 90% of the time with nothing changed, because its score sits in a 0.41 to 0.50 band. On gpt-5.4-nano, 100%. On gpt-5.5, 0%.
Rewording the rubric moves 16% of gpt-5.4-nano's verdicts. gpt-5.5 moves none. Asked for "incorrect" instead of "correct", gpt-5.4-nano fails to invert on 4 cases.
Switching from gpt-5.4-nano to gpt-5.5 changes 6 verdicts , and a 0.5 pass mark on the old judge is about 0.62 on the new one.
The fix works and you can see by how much. gpt-5.4-nano wrapped in a 5-vote majority goes from 92% to 99% self-agreement, and adaptive re-asking got there with 177 fewer calls than asking five times everywhere.
Stronger judges are not immune, they are harder to catch
On the 32 tricky cases, gpt-5.5 and the three gpt-6 judges never flip. So we wrote examples/hard.jsonl : 64 cases that pit precision against reasonableness. Rounding that is close enough or not, unit conversions, hedges that commit to the right answer and hedges that don't, extra facts that are right and extra facts that are wrong, equivalent code, multi-step answers with a wrong step and a right total. Six judges, five runs each, 1,920 live verdicts on 26 October 2026.
Every gpt-6 judge flipped, and the rationales show what they flip on:
gpt-6.1-sol on "6.02 x 10^23" for Avogadro's number : fails it four runs out of five for "omits the required unit, per mole", then passes it because "per mole is implicit in the context."
gpt-6-astra on a shirt-price answer with the right total and wrong working : fails it four times, "25% of $20 is $5, not $4", then passes it once because "the final price matches the reference."
gpt-6-luna on "two hydrogen atoms bonded to one oxygen atom" for H2O : F, T, F, T, F across five runs. gpt-5.5 flipped on the same case.
The pattern: weak judges flip everywhere, strong judges flip exactly where a reasonable person could go either way, at 1 to 3% of cases. On an eval of a few hundred cases that is still several verdicts that depend on which run you got. And the bigger number is between tiers: Avogadro without "per mole" passes on every gpt-5 judge and fails on every gpt-6 judge; "32 degrees Fahrenheit" for the freezing point fails on both 5.4 judges and passes on everything above. Kappa between gpt-5.4-nano and gpt-6.1-sol is 0.60. Krippendorff's alpha across all six is 0.81, across the three gpt-6 judges 0.97.
One more thing the run caught: OpenAI's moderation rejected one of gpt-6-luna's 320 calls, on a question about sum(xs) / len(xs) , as a flagged prompt. The parser audit recorded it and marked that judge "fragile". A pipeline without error accounting would have counted that as a fail or skipped it silently.
Both sets ship with the package. examples/tricky.jsonl is where cheap judges fall apart; examples/hard.jsonl is where the expensive ones do.
You get the card in the terminal, a single-file report.html you can send to anyone, JSON for machines, and markdown for a pull request.
Self-agreement per judge: the chance two runs give the same verdict on a case. Independent of how many runs you did. The headline number.
Flipped cases : which cases had more than one verdict across runs.
Agreement between judges : percent agreement and Cohen's kappa for every pair, Krippendorff's alpha across all of them.
Accuracy vs human labels : accuracy and kappa against your label column, plus the misses.
Flakiest cases : every verdict from every judge and run, as chips, ranked by disagreement. This is where you go to fix your rubric.
Parser audit : how each reply was read (strict JSON, code fence, buried in prose, no verdict), so a parsing problem never gets blamed on the model.
Every run is saved as JSON under .mistrial/runs/ and every response is cached on disk, so the probes below can read a saved run, re-rendering is free, and --fresh is explicit. Beyond the card: the probes for the ways judges go wrong in production, the fixes for the flakiness they find, and the CI setup to keep it that way.
One case per line. JSONL, JSON or CSV.
{ "id" : " capital-france " , "inputs" : " What is the capital of France? " , "outputs" : " Paris. " , "reference_outputs" : " Paris " , "label" : true }
{ "id" : " seven-times-eight " , "inputs" : " What is 7 × 8? " , "outputs" : " 54 " , "reference_outputs" : " 56 " , "label" : false }
inputs and outputs are required. reference_outputs and the human label are optional but unlock the correctness rubric and the accuracy column. Common aliases work: input , question , answer , expected , ground_truth .
Three kinds, mixed freely with repeated --judge

[truncated]

## Original Extract

Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card. - hemanth1999k/mistrial

GitHub - hemanth1999k/mistrial: Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card. · GitHub
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
hemanth1999k
mistrial Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Is your LLM judge flaky? Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card.
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
38 Commits 38 Commits Folders and files
.github .github docs docs examples examples src/ mistrial src/ mistrial tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Is your LLM judge flaky?
Re-run any evaluator across runs, judge models and rubric wordings, and get a reliability card.
When the verdict can't be trusted, declare a mistrial.
pip install mistrial
export OPENAI_API_KEY=... # or ANTHROPIC_API_KEY, GEMINI_API_KEY, GROQ_API_KEY ... see `mistrial providers`
mistrial run cases.jsonl -j openai:gpt-5.4-mini -j openai:gpt-5.5 --runs 5
No key? pipx run mistrial demo runs three fake judges that are built to be flaky on purpose.
You ship AI features based on a score from an LLM judge. The score gates CI, decides which prompt wins, and goes on the dashboard. But the judge is a model sampling from a distribution. Ask it the same question twice and it changes its verdict on some of your cases. Swap the model and it changes its mind on more. Reword the rubric without changing its meaning and it moves again. One eval run is one sample from a noisy process, and almost nobody measures the noise.
These are real replies. On the first case, gpt-5.4-nano wrote "the added mention of 1811 does not contradict the core requirement" on one run and "1811 is not part of the reference, so the answer is incorrect" on the next. On the second, it worked out that 239,000 miles is 384,000 km and then failed the answer for being in miles. Each of those judges reports a clean score on its own. The flakiness only shows up when you ask again and compare.
The research agrees. The same judge gives different scores to the same answer across runs ( Rating Roulette, EMNLP 2025 ). Raw percent agreement overstates reliability by 30 to 40 points compared with chance-corrected kappa, across 21 judges ( Reliability without Validity, 2026 ).
Eval platforms report one number and stop. LangSmith, Braintrust, Langfuse and friends run the judge once per case. Some can repeat an experiment and show a standard deviation, none report self-agreement, kappa between judges, or which cases flipped.
The statistics exist, but not as a tool. Test-retest reliability, Cohen's kappa and Krippendorff's alpha are a century old and standard in every field that uses human raters. They live in research notebooks, not in something you can run against your evaluator in a minute.
The projects that tried stop at the card. judgecheck, judgebench, Assay rejudge and judge-drift-sentinel each cover a piece of this well, and docs/COMPARISON.md credits what each does better. None of them take your existing evaluator function, run in CI, or tell you what to do when the judge is flaky.
Before "how good is my model?" you need "can I trust the number that told me?"
Point it at the evaluator you already use: a model name, or the function you run in LangSmith or openevals. It asks the judge several times, across one or more judge models, and under rewordings of the rubric, and reports the spread: self-agreement, flipped cases, kappa and alpha between judges, accuracy against your human labels, and the exact cases that move. Then it does something about it: a gate for CI, a stabilized judge with the same signature, and a list of the cases a human should look at.
The 90 second version. Every number in it is from the runs below.
Three current OpenAI judges on examples/tricky.jsonl , 32 cases built to be hard to judge (hedges, rounding, unit conversions, extra details, equivalent code), five runs each. 480 live calls on 25 October 2026. Every number below is from that run, and the GIF replays those replies.
mistrial · 32 cases · 3 judges · 5 runs each · 480 verdicts in 76.66s
Consistency (same judge, asked again)
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━┓
┃ judge ┃ self-agreement ┃ flipped cases ┃ vs human ┃ errors ┃ reading ┃
┡━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━━━━━━━━━┩
│ openai:gpt-5.4-nano │ 92% │ 5/32 (16%) │ 81% (κ 0.63) │ 0 │ noticeably flaky │
│ openai:gpt-5.4-mini │ 95% │ 4/32 (12%) │ 88% (κ 0.75) │ 0 │ mostly stable │
│ openai:gpt-5.5 │ 100% │ 0/32 (0%) │ 91% (κ 0.81) │ 0 │ stable │
└─────────────────────┴────────────────┴───────────────┴───────────────┴────────┴──────────────────┘
Flakiest cases (everyone's verdicts pooled)
┏━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━┓
┃ case ┃ disagreement ┃ openai:gpt-5.4-nano ┃ openai:gpt-5.4-mini ┃ openai:gpt-5.5 ┃ human ┃
┡━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━╇━━━━━━━┩
│ arigato │ 40% │ F F F F F │ F T T T T │ T T T T T │ T │
│ austen-wrong-year │ 40% │ T F F F T │ T T T T F │ F F F F F │ F │
│ largest-country-hedge │ 40% │ F F F F F │ T F F F F │ T T T T T │ F │
│ moon-miles │ 33% │ F F F F F │ F F F F F │ T T T T T │ T │
│ percent-unit │ 33% │ T T T T T │ F F F F F │ F F F F F │ F │
└───────────────────────┴──────────────┴─────────────────────┴─────────────────────┴────────────────┴───────┘
What the run says, in order:
gpt-5.4-nano flips on 5 of 32 cases with nothing changed, and agrees with the human label on 81%. gpt-5.5 never flips and agrees on 91%. The judges agree with each other at kappa 0.69 to 0.81, Krippendorff's alpha 0.75.
A CI gate at a 0.5 pass mark on gpt-5.4-mini goes red 90% of the time with nothing changed, because its score sits in a 0.41 to 0.50 band. On gpt-5.4-nano, 100%. On gpt-5.5, 0%.
Rewording the rubric moves 16% of gpt-5.4-nano's verdicts. gpt-5.5 moves none. Asked for "incorrect" instead of "correct", gpt-5.4-nano fails to invert on 4 cases.
Switching from gpt-5.4-nano to gpt-5.5 changes 6 verdicts , and a 0.5 pass mark on the old judge is about 0.62 on the new one.
The fix works and you can see by how much. gpt-5.4-nano wrapped in a 5-vote majority goes from 92% to 99% self-agreement, and adaptive re-asking got there with 177 fewer calls than asking five times everywhere.
Stronger judges are not immune, they are harder to catch
On the 32 tricky cases, gpt-5.5 and the three gpt-6 judges never flip. So we wrote examples/hard.jsonl : 64 cases that pit precision against reasonableness. Rounding that is close enough or not, unit conversions, hedges that commit to the right answer and hedges that don't, extra facts that are right and extra facts that are wrong, equivalent code, multi-step answers with a wrong step and a right total. Six judges, five runs each, 1,920 live verdicts on 26 October 2026.
Every gpt-6 judge flipped, and the rationales show what they flip on:
gpt-6.1-sol on "6.02 x 10^23" for Avogadro's number : fails it four runs out of five for "omits the required unit, per mole", then passes it because "per mole is implicit in the context."
gpt-6-astra on a shirt-price answer with the right total and wrong working : fails it four times, "25% of $20 is $5, not $4", then passes it once because "the final price matches the reference."
gpt-6-luna on "two hydrogen atoms bonded to one oxygen atom" for H2O : F, T, F, T, F across five runs. gpt-5.5 flipped on the same case.
The pattern: weak judges flip everywhere, strong judges flip exactly where a reasonable person could go either way, at 1 to 3% of cases. On an eval of a few hundred cases that is still several verdicts that depend on which run you got. And the bigger number is between tiers: Avogadro without "per mole" passes on every gpt-5 judge and fails on every gpt-6 judge; "32 degrees Fahrenheit" for the freezing point fails on both 5.4 judges and passes on everything above. Kappa between gpt-5.4-nano and gpt-6.1-sol is 0.60. Krippendorff's alpha across all six is 0.81, across the three gpt-6 judges 0.97.
One more thing the run caught: OpenAI's moderation rejected one of gpt-6-luna's 320 calls, on a question about sum(xs) / len(xs) , as a flagged prompt. The parser audit recorded it and marked that judge "fragile". A pipeline without error accounting would have counted that as a fail or skipped it silently.
Both sets ship with the package. examples/tricky.jsonl is where cheap judges fall apart; examples/hard.jsonl is where the expensive ones do.
You get the card in the terminal, a single-file report.html you can send to anyone, JSON for machines, and markdown for a pull request.
Self-agreement per judge: the chance two runs give the same verdict on a case. Independent of how many runs you did. The headline number.
Flipped cases : which cases had more than one verdict across runs.
Agreement between judges : percent agreement and Cohen's kappa for every pair, Krippendorff's alpha across all of them.
Accuracy vs human labels : accuracy and kappa against your label column, plus the misses.
Flakiest cases : every verdict from every judge and run, as chips, ranked by disagreement. This is where you go to fix your rubric.
Parser audit : how each reply was read (strict JSON, code fence, buried in prose, no verdict), so a parsing problem never gets blamed on the model.
Every run is saved as JSON under .mistrial/runs/ and every response is cached on disk, so the probes below can read a saved run, re-rendering is free, and --fresh is explicit. Beyond the card: the probes for the ways judges go wrong in production, the fixes for the flakiness they find, and the CI setup to keep it that way.
One case per line. JSONL, JSON or CSV.
{ "id" : " capital-france " , "inputs" : " What is the capital of France? " , "outputs" : " Paris. " , "reference_outputs" : " Paris " , "label" : true }
{ "id" : " seven-times-eight " , "inputs" : " What is 7 × 8? " , "outputs" : " 54 " , "reference_outputs" : " 56 " , "label" : false }
inputs and outputs are required. reference_outputs and the human label are optional but unlock the correctness rubric and the accuracy column. Common aliases work: input , question , answer , expected , ground_truth .
Three kinds, mixed freely with repeated --judge

[truncated]
