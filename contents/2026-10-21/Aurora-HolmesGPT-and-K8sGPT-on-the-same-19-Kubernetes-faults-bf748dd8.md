---
source: "https://www.aurorasre.ai/benchmark"
hn_url: "https://news.ycombinator.com/item?id=50180570"
title: "Aurora, HolmesGPT and K8sGPT on the same 19 Kubernetes faults"
article_title: "AI SRE Benchmark: Aurora, HolmesGPT and K8sGPT on 19 Kubernetes Faults"
image: "https://www.aurorasre.ai/opengraph-image"
author: "beng360"
captured_at: "2026-10-21T13:29:45Z"
capture_tool: "hn-digest"
hn_id: 50180570
score: 1
comments: 0
posted_at: "2026-10-21T13:25:29Z"
tags:
  - hacker-news
---

# Aurora, HolmesGPT and K8sGPT on the same 19 Kubernetes faults

- HN: [50180570](https://news.ycombinator.com/item?id=50180570)
- Source: [www.aurorasre.ai](https://www.aurorasre.ai/benchmark)
- Score: 1
- Comments: 0
- Posted: 2026-10-21T13:25:29Z

## Translation

Title: Aurora, HolmesGPT and K8sGPT on the same 19 Kubernetes faults
Article title: AI SRE Benchmark: Aurora, HolmesGPT and K8sGPT on 19 Kubernetes Faults
Description: 19 injected Kubernetes faults, three AI incident agents, two runs each, every answer graded by a judge model against the injected fault. Every trace and every miss is public.

Article text:
AI SRE Benchmark: Aurora, HolmesGPT and K8sGPT on 19 Kubernetes Faults Aurora by Arvo AI Documentation Features About Us Blog Runbooks Benchmark Contact Invest Star on GitHub 359 Log in Book a demo AI SRE benchmark
19 Kubernetes faults. 3 agents. Every trace public.
We break a small app on a kind cluster in known ways, hand each agent only the alert, and grade the root cause against the fault we injected. A judge model scores every run. A person has checked Aurora’s verdicts; HolmesGPT’s and K8sGPT’s are the judge’s, read over by a second model, not yet checked by a person. Misses are published with the hits. Arvo builds Aurora and wrote the scenarios and the grading criteria; read the numbers with that in mind.
Aurora has 3 more correct runs of 38. Fisher exact test on correct against not, two-sided: p = 0.52 . That gap is not significant; with 38 runs per agent and a model judge, read the two as roughly even.
kind v0.33.0, Kubernetes v1.37.0, 3 nodes . The shop app: nginx, a Python API, a worker, redis, a load generator. Synthetic, no customer code.
One fault is injected with the scenario's break.sh ; check.sh proves it is visible.
Each agent gets the alert text. Aurora's prompt adds "Cluster: aurora-lab (kubeconfig). Investigate and find the root cause."; HolmesGPT's adds "Namespace shop, cluster kind-aurora-lab. Find the root cause and say what changed." K8sGPT has no notion of an alert and scans the namespace. Aurora runs under its default read-only observability_only policy; HolmesGPT and K8sGPT got the cluster admin kubeconfig, and no transcript shows a mutating command running.
HolmesGPT ran through holmes ask with default toolsets and a 40-step cap it never reached; the lab has no Alertmanager, so holmes investigate was not used. Its own command allow-list blocked kubectl rollout history in 12 runs, 2 of its 7 non-correct runs among them; we did not rerun those with the command allowed. Aurora's policy had the same block in the first lab report.
Correct means every must_identify item in the scenario is stated and no listed wrong answer is asserted; partial means some items. Judge: us.anthropic.claude-opus-5-5 , the same vendor as every model under test. The prompt does not name the agent, though the tool names in each command list give it away.
Every tool ran twice per scenario. Aurora's 38 verdicts were checked by a person, who changed 0 of them. HolmesGPT’s and K8sGPT’s were graded by the judge model, then re-graded from scratch by a second model. It agreed on all 38 HolmesGPT verdicts and disputed 6 K8sGPT labels, where K8sGPT only restated the symptom: the judge said partial or wrong, the reviewer said no answer. We kept the judge’s labels, and every grade file carries both opinions. No person has checked these 76 yet. As a bias check, every answer was graded again with the agent’s name, its tool names and the cluster name removed from the judge’s prompt (runner/grade_blind.py). The blind verdict matched the published one for 38 of 38 Aurora, 37 of 38 HolmesGPT, 36 of 38 K8sGPT runs; blind correct counts are Aurora 34, HolmesGPT 30, K8sGPT 0. Every disagreement is listed in runs/blind-regrade.md in the repo. Aurora's runs are from commit 87d22baf, the day before the scenarios were written, so it was not tuned on them. The fixes the first lab report prompted are open pull requests on the Aurora repo; the six reruns on the fixed build are in the lab repo and not in this table.
What is not comparable: K8sGPT's row. Its analyzers flag objects and the model explains the flag, so it is scored on a criterion built for alert-driven investigators. It ran on Sonnet 4.5 because it sends a temperature parameter that Bedrock rejects for Sonnet 5; we did not try a proxy that strips it. Times: Aurora's include contention from its own background jobs in the lab stack (REPORT.md in the repo); K8sGPT's cover two analyze passes, JSON then text. Costs are recorded for Aurora only; HolmesGPT's token counts were not captured.
Everything is in Arvo-AI/aurora-lab : scenarios, the runner, every transcript and grade, and this table as results.json .
git clone https://github.com/Arvo-AI/aurora-lab && cd aurora-lab
./lab.sh up && ./lab.sh break 03-bad-image-tag
./runner/h2h.sh 03-bad-image-tag runs/mine To put another agent on the board, open an issue titled "Benchmark: <agent>" with a script that takes a kubeconfig and the alert text and prints the root cause. We run it on all 19 faults and publish the result, whatever it is. Submit an agent .
Read the full write-up: lab report v1 . Fix the faults yourself: runbooks , each with Aurora's real trace.
Results updated 2026-10-21 . Scenario and tool versions are recorded with every run in the repo.
Want to see what the community is building?
The best new features, integrations, and incident stories from the Aurora community, plus what we're shipping next. Bi-weekly, never spam.
© 2026 Arvo A.I. Ltd. All rights reserved.

## Original Extract

19 injected Kubernetes faults, three AI incident agents, two runs each, every answer graded by a judge model against the injected fault. Every trace and every miss is public.

AI SRE Benchmark: Aurora, HolmesGPT and K8sGPT on 19 Kubernetes Faults Aurora by Arvo AI Documentation Features About Us Blog Runbooks Benchmark Contact Invest Star on GitHub 359 Log in Book a demo AI SRE benchmark
19 Kubernetes faults. 3 agents. Every trace public.
We break a small app on a kind cluster in known ways, hand each agent only the alert, and grade the root cause against the fault we injected. A judge model scores every run. A person has checked Aurora’s verdicts; HolmesGPT’s and K8sGPT’s are the judge’s, read over by a second model, not yet checked by a person. Misses are published with the hits. Arvo builds Aurora and wrote the scenarios and the grading criteria; read the numbers with that in mind.
Aurora has 3 more correct runs of 38. Fisher exact test on correct against not, two-sided: p = 0.52 . That gap is not significant; with 38 runs per agent and a model judge, read the two as roughly even.
kind v0.33.0, Kubernetes v1.37.0, 3 nodes . The shop app: nginx, a Python API, a worker, redis, a load generator. Synthetic, no customer code.
One fault is injected with the scenario's break.sh ; check.sh proves it is visible.
Each agent gets the alert text. Aurora's prompt adds "Cluster: aurora-lab (kubeconfig). Investigate and find the root cause."; HolmesGPT's adds "Namespace shop, cluster kind-aurora-lab. Find the root cause and say what changed." K8sGPT has no notion of an alert and scans the namespace. Aurora runs under its default read-only observability_only policy; HolmesGPT and K8sGPT got the cluster admin kubeconfig, and no transcript shows a mutating command running.
HolmesGPT ran through holmes ask with default toolsets and a 40-step cap it never reached; the lab has no Alertmanager, so holmes investigate was not used. Its own command allow-list blocked kubectl rollout history in 12 runs, 2 of its 7 non-correct runs among them; we did not rerun those with the command allowed. Aurora's policy had the same block in the first lab report.
Correct means every must_identify item in the scenario is stated and no listed wrong answer is asserted; partial means some items. Judge: us.anthropic.claude-opus-5-5 , the same vendor as every model under test. The prompt does not name the agent, though the tool names in each command list give it away.
Every tool ran twice per scenario. Aurora's 38 verdicts were checked by a person, who changed 0 of them. HolmesGPT’s and K8sGPT’s were graded by the judge model, then re-graded from scratch by a second model. It agreed on all 38 HolmesGPT verdicts and disputed 6 K8sGPT labels, where K8sGPT only restated the symptom: the judge said partial or wrong, the reviewer said no answer. We kept the judge’s labels, and every grade file carries both opinions. No person has checked these 76 yet. As a bias check, every answer was graded again with the agent’s name, its tool names and the cluster name removed from the judge’s prompt (runner/grade_blind.py). The blind verdict matched the published one for 38 of 38 Aurora, 37 of 38 HolmesGPT, 36 of 38 K8sGPT runs; blind correct counts are Aurora 34, HolmesGPT 30, K8sGPT 0. Every disagreement is listed in runs/blind-regrade.md in the repo. Aurora's runs are from commit 87d22baf, the day before the scenarios were written, so it was not tuned on them. The fixes the first lab report prompted are open pull requests on the Aurora repo; the six reruns on the fixed build are in the lab repo and not in this table.
What is not comparable: K8sGPT's row. Its analyzers flag objects and the model explains the flag, so it is scored on a criterion built for alert-driven investigators. It ran on Sonnet 4.5 because it sends a temperature parameter that Bedrock rejects for Sonnet 5; we did not try a proxy that strips it. Times: Aurora's include contention from its own background jobs in the lab stack (REPORT.md in the repo); K8sGPT's cover two analyze passes, JSON then text. Costs are recorded for Aurora only; HolmesGPT's token counts were not captured.
Everything is in Arvo-AI/aurora-lab : scenarios, the runner, every transcript and grade, and this table as results.json .
git clone https://github.com/Arvo-AI/aurora-lab && cd aurora-lab
./lab.sh up && ./lab.sh break 03-bad-image-tag
./runner/h2h.sh 03-bad-image-tag runs/mine To put another agent on the board, open an issue titled "Benchmark: <agent>" with a script that takes a kubeconfig and the alert text and prints the root cause. We run it on all 19 faults and publish the result, whatever it is. Submit an agent .
Read the full write-up: lab report v1 . Fix the faults yourself: runbooks , each with Aurora's real trace.
Results updated 2026-10-21 . Scenario and tool versions are recorded with every run in the repo.
Want to see what the community is building?
The best new features, integrations, and incident stories from the Aurora community, plus what we're shipping next. Bi-weekly, never spam.
© 2026 Arvo A.I. Ltd. All rights reserved.
