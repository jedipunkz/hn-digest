---
source: "https://paretoops.dev/"
hn_url: "https://news.ycombinator.com/item?id=50145643"
title: "ParetoOps – CI gate on cost per successful task for LLM evals"
article_title: "ParetoOps: A CI Check for Prompt and Model Changes"
image: "https://paretoops.dev/og.png"
author: "manasapalas"
captured_at: "2026-10-18T22:15:30Z"
capture_tool: "hn-digest"
hn_id: 50145643
score: 1
comments: 0
posted_at: "2026-10-18T21:25:47Z"
tags:
  - hacker-news
---

# ParetoOps – CI gate on cost per successful task for LLM evals

- HN: [50145643](https://news.ycombinator.com/item?id=50145643)
- Source: [paretoops.dev](https://paretoops.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-10-18T21:25:47Z

## Translation

Title: ParetoOps – CI gate on cost per successful task for LLM evals
Article title: ParetoOps: A CI Check for Prompt and Model Changes
Description: Get a verdict on every prompt and model change before it merges: did it get worse, or more expensive? A CI check that reads your existing eval results. Free to start.

Article text:
Skip to content ParetoOps Product Calculator Models Learn Pricing Docs Free key Get early access Early access Product Calculator Models Learn Pricing Docs Free License Key Research FAQ A CI check for AI changes Did that prompt or model change make your AI agent worse, or more expensive?
ParetoOps runs in your CI and reads the results your eval tool already produces. It checks every pull request against your limits, and with Pro against your last good run, so you get a clear verdict before it merges.
Want the regression check on your pull requests? Join the interest list for a 30-day Pro trial →
Try it on your own results: npx pareto-ops analyze results.json --min-acc 0.9
Works with Promptfoo, LangSmith, Braintrust and DeepEval · Runs offline, never calls a model · Free to use today
Pro A regression, caught on the pull request
The change still clears the 50% accuracy threshold, so a threshold check alone would have passed it.
Free keys get a pull request comment too: your options ranked by cost per successful task, and whether your accuracy and cost limits passed. Pro adds this comparison with your last good run.
Three steps, no new infrastructure
Keep running the evals you have
Promptfoo, LangSmith, Braintrust, DeepEval, or any CSV or JSON with one row per task. ParetoOps doesn't run evals or call a model.
uses: certiqa/paretoops-action@v1 in GitHub Actions, or one command in GitLab, Jenkins or a cron job. Set it up →
Get a verdict on every pull request
A comment says pass or fail against your limits, what you're paying for each task that passes, and with Pro, what changed since your last good run.
See it on sample data in two minutes
Download the sample eval results ParetoOps ships with and run the free analysis. No sign-up, no key, nothing leaves your machine.
Needs Node 18 or newer. Prefer Python? pip install pareto-ops installs the same command. Walk through the report →
$ curl -O https://paretoops.dev/eval_benchmark.json
$ npx pareto-ops analyze eval_benchmark.json --min-acc 0.9
Gemini 2.5 Flash 67.0% $0.0022 per success ✔ OPTIMAL
Claude Haiku 4.5 92.0% $0.0035 per success ✔ OPTIMAL
Claude Sonnet 5 94.0% $0.0233 per success ✔ OPTIMAL
Claude Opus 4.8 80.0% $0.0327 per success ✖ DOMINATED
🌟 Recommended Sweet Spot: Claude Haiku 4.5 at $0.0035/successful
task (92.0% accuracy) Condensed from the real output on the sample dataset. The sweet spot is the best pick: the cheapest option above your accuracy bar.
Code changes get tests. Prompt and model changes mostly get a glance.
A pass mark like "accuracy above 50%" only asks whether a number clears a line. It can't tell you the change made things worse than what you already had. The same sample change, two verdicts:
Accuracy is 66.5%, above the 50% bar you set. Nothing in a threshold check says the change got worse.
On the 200 tasks both runs share, accuracy fell 15.0 points (95% CI -22.5 to -7.5) and cost per success rose ×1.42 (×1.24–×1.64). The change is worse, not unlucky.
Rewording a prompt moved results 11 to 58 times more than rerunning it ( Chen, Qian et al., 2026 ). Comparing two averages can't catch that.
The cheaper price list isn't always cheaper
In 32% of model pairs studied, the model with the lower list price cost more overall ( Chen, Zhang et al., 2026 ).
Join the interest list to request a 30-day Pro trial →
The verdicts are real engine output on the sample data shipped with ParetoOps. See the full report →
Two questions every team with an eval ends up asking
1. Is this change safe to merge?
Prompts, retrieval settings and model versions change all the time, and each change can quietly cost accuracy or money. This is the question for any team, even with a single model: ParetoOps checks every pull request against your limits (Free) and against your last good run (Pro).
2. Which configuration should we run?
When you pick a model, switch to a new release, or route easy requests to a cheaper one, ParetoOps ranks every option by what a successful task actually costs, and recommends the cheapest one that clears your accuracy bar.
Find the cheapest option that's good enough
Cost per successful task is total spend divided by the tasks that passed. A cheap model that fails often can cost more per good answer than a pricier one that doesn't.
Flags options you can drop: beaten on both cost and accuracy.
Recommends a best pick: the cheapest option above your accuracy bar.
Prices tokens correctly: prompt caching and reasoning tokens included.
50% 60% 70% 80% 90% 100% $0.002 $0.005 $0.01 $0.02 $0.05 Cost per successful task (log scale) → Success rate ↑ Your accuracy bar: 90% Claude Sonnet 5 Claude Haiku 4.5 Claude Opus 4.8 (legacy) Gemini 2.5 Flash On the frontier Best pick: cheapest option above the bar Beaten: costs more for the same or worse results "We only evaluate one model." Why would anyone run more than one?
New releases and price changes: "can we switch?" is answered by re-running the same eval.
Routing: easy requests to a small model, hard ones to a large one, split by measured results.
Agents: each step may use a different model.
Prompts count too: two prompts on one model are two options to compare.
A tested fallback if your provider has an outage or raises prices.
Compare with your own numbers → · Can't I just always use the frontier model? →
Offline , never calls a model, no telemetry · Deterministic , arithmetic on your data · Any CI , one command that exits 0 or 1 ( GitHub Action , GitLab, Jenkins) · Closed source, how to verify →
Every statistic here traces to a cited study or the sample data shipped with the engine ( research ). See every command →
Start free. Add the regression gate with Pro.
Cost per successful task for every option
Which options to drop, and the best pick
CI check on accuracy and cost limits, with pull request comments
$99 per month, billed yearly · per GitHub organization
Regression gate against your last good run, task by task, with confidence intervals
Baselines stored on a git branch
Early stopping for failing evals
Request a 30-day Pro trial key of the regression gate. We'll get in touch.
Cost per successful task (also called cost-of-pass in academic literature) is the total inference spend on a set of benchmark tasks divided by the number of tasks that actually passed. Because failed model calls still cost money, a model with a low token price that fails often frequently costs significantly more per useful output than a more expensive, reliable model.
Price per token only reflects nominal provider rates per million tokens. It ignores how many tokens a model consumes to solve a prompt, how many reasoning or "thinking" tokens it generates, and how often it fails. In published research, in 32% of model pairs studied, the model with the lower list price cost more in total—in extreme cases up to 28x more—due to verbosity and reasoning token overhead.
Yes. The main use is checking every prompt, retrieval or model-version change against your last good run, which works with a single model. The model comparison is the second use: when a new model ships, a price changes, or you want to route easy requests to a cheaper option, you run the same eval on the alternatives and see which one costs least per successful task. A configuration is more than a model, so two prompts on the same model count as two options.
With ParetoOps CI gating (`pareto-ops gate` or `pareto-ops ci`), every pull request can be evaluated against your minimum accuracy and maximum cost-per-success thresholds. The gate exits with code 0 on pass or 1 on failure, blocking regressions in GitHub Actions, GitLab CI, Jenkins, or any CI runner. A threshold only checks a fixed number, so a change can clear it and still be worse than before. Pro adds a paired, task-by-task comparison against your last good run, with confidence intervals, which catches that.
No. ParetoOps is 100% offline and deterministic. It performs local arithmetic on the benchmark results your eval framework generates. It makes no API calls to LLM providers, sends zero telemetry, and verifies license keys completely offline.
Pro is for engineering teams shipping LLM changes through pull requests: a paired, task-by-task regression gate against your last good run (with confidence intervals), baselines stored on a git branch, early stopping for failing eval runs, step-level cost attribution for agentic runs, and a 3D frontier with P95 latency. Pull request comments are part of Free with a free key.
Put a verdict on every prompt and model change
Run the free check on your eval results in a minute. No sign-up.
Cost per successful task for AI. Find the configurations worth paying for, and keep them that way.

## Original Extract

Get a verdict on every prompt and model change before it merges: did it get worse, or more expensive? A CI check that reads your existing eval results. Free to start.

Skip to content ParetoOps Product Calculator Models Learn Pricing Docs Free key Get early access Early access Product Calculator Models Learn Pricing Docs Free License Key Research FAQ A CI check for AI changes Did that prompt or model change make your AI agent worse, or more expensive?
ParetoOps runs in your CI and reads the results your eval tool already produces. It checks every pull request against your limits, and with Pro against your last good run, so you get a clear verdict before it merges.
Want the regression check on your pull requests? Join the interest list for a 30-day Pro trial →
Try it on your own results: npx pareto-ops analyze results.json --min-acc 0.9
Works with Promptfoo, LangSmith, Braintrust and DeepEval · Runs offline, never calls a model · Free to use today
Pro A regression, caught on the pull request
The change still clears the 50% accuracy threshold, so a threshold check alone would have passed it.
Free keys get a pull request comment too: your options ranked by cost per successful task, and whether your accuracy and cost limits passed. Pro adds this comparison with your last good run.
Three steps, no new infrastructure
Keep running the evals you have
Promptfoo, LangSmith, Braintrust, DeepEval, or any CSV or JSON with one row per task. ParetoOps doesn't run evals or call a model.
uses: certiqa/paretoops-action@v1 in GitHub Actions, or one command in GitLab, Jenkins or a cron job. Set it up →
Get a verdict on every pull request
A comment says pass or fail against your limits, what you're paying for each task that passes, and with Pro, what changed since your last good run.
See it on sample data in two minutes
Download the sample eval results ParetoOps ships with and run the free analysis. No sign-up, no key, nothing leaves your machine.
Needs Node 18 or newer. Prefer Python? pip install pareto-ops installs the same command. Walk through the report →
$ curl -O https://paretoops.dev/eval_benchmark.json
$ npx pareto-ops analyze eval_benchmark.json --min-acc 0.9
Gemini 2.5 Flash 67.0% $0.0022 per success ✔ OPTIMAL
Claude Haiku 4.5 92.0% $0.0035 per success ✔ OPTIMAL
Claude Sonnet 5 94.0% $0.0233 per success ✔ OPTIMAL
Claude Opus 4.8 80.0% $0.0327 per success ✖ DOMINATED
🌟 Recommended Sweet Spot: Claude Haiku 4.5 at $0.0035/successful
task (92.0% accuracy) Condensed from the real output on the sample dataset. The sweet spot is the best pick: the cheapest option above your accuracy bar.
Code changes get tests. Prompt and model changes mostly get a glance.
A pass mark like "accuracy above 50%" only asks whether a number clears a line. It can't tell you the change made things worse than what you already had. The same sample change, two verdicts:
Accuracy is 66.5%, above the 50% bar you set. Nothing in a threshold check says the change got worse.
On the 200 tasks both runs share, accuracy fell 15.0 points (95% CI -22.5 to -7.5) and cost per success rose ×1.42 (×1.24–×1.64). The change is worse, not unlucky.
Rewording a prompt moved results 11 to 58 times more than rerunning it ( Chen, Qian et al., 2026 ). Comparing two averages can't catch that.
The cheaper price list isn't always cheaper
In 32% of model pairs studied, the model with the lower list price cost more overall ( Chen, Zhang et al., 2026 ).
Join the interest list to request a 30-day Pro trial →
The verdicts are real engine output on the sample data shipped with ParetoOps. See the full report →
Two questions every team with an eval ends up asking
1. Is this change safe to merge?
Prompts, retrieval settings and model versions change all the time, and each change can quietly cost accuracy or money. This is the question for any team, even with a single model: ParetoOps checks every pull request against your limits (Free) and against your last good run (Pro).
2. Which configuration should we run?
When you pick a model, switch to a new release, or route easy requests to a cheaper one, ParetoOps ranks every option by what a successful task actually costs, and recommends the cheapest one that clears your accuracy bar.
Find the cheapest option that's good enough
Cost per successful task is total spend divided by the tasks that passed. A cheap model that fails often can cost more per good answer than a pricier one that doesn't.
Flags options you can drop: beaten on both cost and accuracy.
Recommends a best pick: the cheapest option above your accuracy bar.
Prices tokens correctly: prompt caching and reasoning tokens included.
50% 60% 70% 80% 90% 100% $0.002 $0.005 $0.01 $0.02 $0.05 Cost per successful task (log scale) → Success rate ↑ Your accuracy bar: 90% Claude Sonnet 5 Claude Haiku 4.5 Claude Opus 4.8 (legacy) Gemini 2.5 Flash On the frontier Best pick: cheapest option above the bar Beaten: costs more for the same or worse results "We only evaluate one model." Why would anyone run more than one?
New releases and price changes: "can we switch?" is answered by re-running the same eval.
Routing: easy requests to a small model, hard ones to a large one, split by measured results.
Agents: each step may use a different model.
Prompts count too: two prompts on one model are two options to compare.
A tested fallback if your provider has an outage or raises prices.
Compare with your own numbers → · Can't I just always use the frontier model? →
Offline , never calls a model, no telemetry · Deterministic , arithmetic on your data · Any CI , one command that exits 0 or 1 ( GitHub Action , GitLab, Jenkins) · Closed source, how to verify →
Every statistic here traces to a cited study or the sample data shipped with the engine ( research ). See every command →
Start free. Add the regression gate with Pro.
Cost per successful task for every option
Which options to drop, and the best pick
CI check on accuracy and cost limits, with pull request comments
$99 per month, billed yearly · per GitHub organization
Regression gate against your last good run, task by task, with confidence intervals
Baselines stored on a git branch
Early stopping for failing evals
Request a 30-day Pro trial key of the regression gate. We'll get in touch.
Cost per successful task (also called cost-of-pass in academic literature) is the total inference spend on a set of benchmark tasks divided by the number of tasks that actually passed. Because failed model calls still cost money, a model with a low token price that fails often frequently costs significantly more per useful output than a more expensive, reliable model.
Price per token only reflects nominal provider rates per million tokens. It ignores how many tokens a model consumes to solve a prompt, how many reasoning or "thinking" tokens it generates, and how often it fails. In published research, in 32% of model pairs studied, the model with the lower list price cost more in total—in extreme cases up to 28x more—due to verbosity and reasoning token overhead.
Yes. The main use is checking every prompt, retrieval or model-version change against your last good run, which works with a single model. The model comparison is the second use: when a new model ships, a price changes, or you want to route easy requests to a cheaper option, you run the same eval on the alternatives and see which one costs least per successful task. A configuration is more than a model, so two prompts on the same model count as two options.
With ParetoOps CI gating (`pareto-ops gate` or `pareto-ops ci`), every pull request can be evaluated against your minimum accuracy and maximum cost-per-success thresholds. The gate exits with code 0 on pass or 1 on failure, blocking regressions in GitHub Actions, GitLab CI, Jenkins, or any CI runner. A threshold only checks a fixed number, so a change can clear it and still be worse than before. Pro adds a paired, task-by-task comparison against your last good run, with confidence intervals, which catches that.
No. ParetoOps is 100% offline and deterministic. It performs local arithmetic on the benchmark results your eval framework generates. It makes no API calls to LLM providers, sends zero telemetry, and verifies license keys completely offline.
Pro is for engineering teams shipping LLM changes through pull requests: a paired, task-by-task regression gate against your last good run (with confidence intervals), baselines stored on a git branch, early stopping for failing eval runs, step-level cost attribution for agentic runs, and a 3D frontier with P95 latency. Pull request comments are part of Free with a free key.
Put a verdict on every prompt and model change
Run the free check on your eval results in a minute. No sign-up.
Cost per successful task for AI. Find the configurations worth paying for, and keep them that way.
