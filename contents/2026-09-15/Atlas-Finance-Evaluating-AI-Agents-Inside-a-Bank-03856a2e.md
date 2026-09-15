---
source: "https://joinhandshake.com/research/benchmarks/articles/atlas-finance-evaluating-ai-agents-inside-a-bank/"
hn_url: "https://news.ycombinator.com/item?id=49720319"
title: "Atlas-Finance: Evaluating AI Agents Inside a Bank"
article_title: "ATLAS-Finance: Evaluating AI Agents Inside a Bank | Handshake"
image: "https://cdn.sanity.io/images/mz2hls6g/production/496434fa3caf33de1e90da4ecac09ee9ab5ac5ed-1920x1080.png?w=1200&fm=webp"
author: "cjbarber"
captured_at: "2026-09-15T23:59:51Z"
capture_tool: "hn-digest"
hn_id: 49720319
score: 1
comments: 1
posted_at: "2026-09-15T23:37:11Z"
tags:
  - hacker-news
---

# Atlas-Finance: Evaluating AI Agents Inside a Bank

- HN: [49720319](https://news.ycombinator.com/item?id=49720319)
- Source: [joinhandshake.com](https://joinhandshake.com/research/benchmarks/articles/atlas-finance-evaluating-ai-agents-inside-a-bank/)
- Score: 1
- Comments: 1
- Posted: 2026-09-15T23:37:11Z

## Translation

Title: Atlas-Finance: Evaluating AI Agents Inside a Bank
Article title: ATLAS-Finance: Evaluating AI Agents Inside a Bank | Handshake
Description: Most finance benchmarks test AI on clean tasks. ATLAS Finance grades agents on 100 expert tasks in real banking environments. See how frontier models performed.

Article text:
ATLAS-Finance: Evaluating AI Agents Inside a Bank | Handshake Skip to content Employers
Sign up ATLAS-Finance: Evaluating AI Agents Inside a Bank
Copied! Share the Blog on Facebook (opens in new window) Share the Blog on X (opens in new window) Share the Blog on Linkedin (opens in new window) TL;DR
The gap: Existing finance benchmarks test agentic financial reasoning, data retrieval, and tool use through static, fully specified tasks. These one-off requests in clean contexts are not reflective of actual deployment. In reality, analysts must operate under ambiguity, gather key context from multiple (and sometimes conflicting) sources, and coordinate with leadership, colleagues across teams, and clients that may send relevant updates mid-task.
Our benchmark: Introducing ATLAS-Finance–100 expert-level tasks inside 13 realistic financial firm environments with multi-party working group teams, full inboxes, drives, notebooks, and calendars. Agent-generated work outputs are graded criterion-by-criterion against expert-authored rubrics with 42-486 binary criteria per task ( code ; data ).
Pushing the frontier: We designed ATLAS-Finance to closely reflect realistic working conditions in a financial firm. Agents that reliably navigate ATLAS-Finance are more prepared to help professionals deliver the real-world performance their clients expect.
What we found: Across 11 frontier models, the highest observed pass rate is 12.3% achieved by Claude Opus 5. Claude Fable 5.1 and GPT-6 Astra achieve 12.0% and 11.3%, respectively. The other eight models remain below 10%. Models consistently failed in three areas:
Applying the wrong financial logic, methodology, timing, or perimeter
Omitting required scope in the final output
Failing to propagate correctly calculated values downstream in the analyses.
The specific mix varies by model and workflow, but recurring, plausible-looking financial mistakes appear across all model families.
Why it matters: Identifying where agents fail under realistic working conditions like ATLAS-Finance environments is a key step in improving them to deliver larger productivity gains in the financial sector. In an actual bank, associates that make critical errors like the ones indicated above force senior members to audit all of their work. Similarly, if an agent makes a mistake like this in the workplace, it negates all of the time savings and expertise acceleration potentially provided by that model.
1. What’s different about ATLAS-Finance
Recent benchmarks have moved from short-form question answering to end-to-end financial work. Handshake’s BankerToolBench (BTB), in particular, evaluated the financial models produced by AI agents completing self-contained junior investment banker workflows. Its results show that frontier agents can complete parts of these workflows but remain unreliable at the standard required for professional delegation. BTB also identifies a remaining gap: simplified, static environments exclude the scattered, evolving, and sometimes contradictory nature of information found in the real world, in addition to organizational-level complexities that must be navigated where this work happens.
We address this gap with ATLAS-Finance by changing the unit of evaluation from a linear workflow to a workplace world. As a vastly richer RL environment, each world contains an advisory firm, specialized teams, updates from senior and junior professionals, clients and counterparties, active and background mandates, communication histories, shared files, calendars, permissions, and firm-specific conventions. The agent works through email, chat, document repositories, notes, calendars, and a local workstation. It must produce the requested financial artifacts and meet an explicit submission contract.
Table 1. Comparison of ATLAS Finance against existing finance and professional-work benchmarks. BTB (BankerToolBench, Handshake AI), APEX (APEX-Agents, Mercor; investment-banking subset), BFB (BigFinanceBench, Rogo), ALE (Agents' Last Exam, UC Berkeley), FAB (Finance Agent Benchmark, Vals AI), FinB (FinanceBench, Patronus AI), FF (FrontierFinance, Samaya AI), and GDPval (OpenAI).
Built by bankers up to VP-level from institutional banks and elite boutiques (such as Morgan Stanley, Bank of America, UBS, and Rothschild & Co), ATLAS-Finance reflects the real-world complexity that professionals navigate every day. The benchmark comprises 100 expert-level tasks in 13 environments, covering 9 industries and 20+ financial-analysis methodologies. These tasks typically take human finance experts 15-30 hours to complete manually.
Table 2. Expert-level workflows that undergird the ATLAS Finance benchmark.
To accurately complete tasks, agents must navigate organizational and relational complexity while collaborating with dozens of unique coworkers and clients. Before it populates a single spreadsheet cell, the agent must:
Assemble its workflow details from environment context
Understand its relationship to each counterparty and the role they play for the job at hand
Resolve controlled contradictions by finding governing sources or identifying the right decision-maker (e.g. “Who has the final say on the financing assumption?”).
A subset of tasks includes a real-time update with corrections to prior assumptions that must be incorporated before delivery.
Agent outputs are evaluated against 14,600 environment-grounded rubric criteria (146 per task on average and all expert-verified). ATLAS-Finance grades the work the way an industry professional would: every task carries 42-486 binary checks organized by schedule, with gating dependencies between them–if the debt schedule is not built with live formulas, everything downstream of it is recorded as unreachable, not merely wrong. A failing run produces a diagnosis in addition to a rubric score: which schedule broke, what the break cost downstream, and, read against the run's transcript, whether the model didn't know the finance, knew it and built it wrong, or built it right and delivered a workbook with static values instead of working formulas. Each rubric criterion comes with a weight indicating its relative importance according to experts. To grade any agent output against these rubrics, we use Gandalf the Grader : an agent-as-judge framework for verifying financial artifacts.
We evaluate 11 frontier models (see Figure 1), all run within the OpenCode agentic harness. Claude Opus 5 performs the best, yet still only manages to pass 12.3% of tasks . Claude Fable 5.1 and GPT-6 Astra are close behind, but the other eight models have significantly lower pass rates.
What counts as a pass: A run passes only if it meets every rubric criterion deemed critical (those weighted 3 or higher) and triggers no penalties (negative criteria). Penalties catch unacceptable issues such as fabricated data or a missing audit trail.
Figure 1. Pass@1 on ATLAS Finance. Error bars show 95% confidence intervals computed via bootstrap.
3. Example Task - LBO Rebuild Staffing at a FIG & media advisory
Below is an example task showcasing what agents must do in ATLAS-Finance environments. Here Project Osprey is a full LBO rebuild involving a specific tax treatment. It would take most human analysts 15+ hours to complete just the Excel modeling involved in this task, not counting the additional fact-finding work necessary to pull the relevant information into the analysis from the bank environment.
All ATLAS-Finance tasks start with the agent’s project “staffing”–it is assigned to a new project at the firm exactly how thousands of industry professionals experience it daily. The agent receives a heads-up message with high-level context that points to a subsequent project details email. This begins the workflow chain of understanding, discovery, analysis, and delivery. Understanding this kickoff notice is the first in dozens of steps an agent must handle to complete this ATLAS Finance task.
Step 1: Agent Staffing Kickoff
Look through your inbox and the Project Osprey attached source materials and rebuild the US sponsor LBO model your engagement lead assigned you (a sponsor’s 2019 LBO of a payments target, with the 338(h)(10) election). Your assignment email is titled “Project Osprey – rebuild the 2019 sponsor LBO (338(h)(10) election)”.
The brief goes on to fix the delivery contract that the agent must satisfy–an exact filename, exact recipients, an exact subject line– and the working standard: the finished workbook must carry live in-cell formulas so hardcoding a computed number where a formula belongs is an error even when the number is right.
3.1. Running Opus 5 powered Agent on this task (73 execution steps, outcome: Fail, rubric score: 0.336)
Here are the steps that a Claude agent took when assigned this task.
Steps 1–19 – Discovery. Searches the assignment subject (1) → reads the assignment thread and seven related threads (2–4, 14) → finds the desk’s OneNote: Modelling Approach, Output Standards, call notes, open items (5–6, 15–19) → lists and downloads every attachment (7–11) → dumps every source workbook programmatically to read real cell values (12–13).
Steps 20–25 – The shadow model. Before touching a single spreadsheet cell, the agent writes the complete LBO in pure Python: sources and uses, debt sweep, tax computation, returns, and iterates it to convergence (20–23), then extends it with a second validation pass (24–25). Every later self-check will be measured against this blueprint.
Steps 26–51 – Build. A single build script generates the whole workbook - nine tabs in the briefed order (26–27); first recalculation with the environment’s spreadsheet engine (28); targeted fixes to the NOL rows and a revenue formula (34–41). When convergence stalls, the agent reads the spreadsheet-recalculation utility’s own source code and reproduces its environment to debug the circularity (43–49). “Converged.” (52)
Steps 52–67 – Validate. Sweeps every output tab and ties the numbers back to the shadow model (52–67). All thirteen integrity checks on the Control sheet read OK; the workbook carries 2,641 live formulas and zero error cells.
Steps 68–73 – Delivery. Verifies the collection copy and the delivery-named copy are byte-identical by SHA-256 (68); sends to the exact recipients with the exact subject and attachment (69); reopens the sent message to confirm the attachment landed (70–72); closes: “Done. The model is built, recalculated, saved and delivered.” (73)
Figure 2. Anatomy of one agent rollout: 73 steps of work, two wrong ideas, and where the rubric deducted points (out of 1460).
During discovery, the agent surfaces a key email from the client’s tax lead. R. Calloway lays out each input for the tax-basis build and annotates each one with its operation: “ added to purchase price, ” “ subtracted in the basis calc, ” “ subtracted .” One line carries no operation at all: Existing Goodwill . The client gave Existing Goodwill no specified treatment, because under the record’s method Existing Goodwill plays no role in the amortizable step-up.
The agent’s main error is born at step 20, inside the agent’s own verification machinery. The Python shadow model–written to protect the build–contains one silent line that subtracts existing goodwill from the amortizable basis. Critically, the model failed to add a comment, consider alternative treatments given the task context, or flag the line for review. Six steps later, transcribing the inputs into the workbook, the agent rewrites the label itself: “ Existing Goodwill – carry-over tax basis ”–a treatment the governing email never assigned is recorded as if it were source data.
Figure 3. The LLM invents an incorrect treatment for the Existing Goodwill value shared in the client’s email. This treatment breaks the financial methodology for this scenario.
From that moment, the run’s guardrails point the wrong way. All of the agent’s self-checks verified against this shadow model, so all thirteen integrity checks read green,

[truncated]

## Original Extract

Most finance benchmarks test AI on clean tasks. ATLAS Finance grades agents on 100 expert tasks in real banking environments. See how frontier models performed.

ATLAS-Finance: Evaluating AI Agents Inside a Bank | Handshake Skip to content Employers
Sign up ATLAS-Finance: Evaluating AI Agents Inside a Bank
Copied! Share the Blog on Facebook (opens in new window) Share the Blog on X (opens in new window) Share the Blog on Linkedin (opens in new window) TL;DR
The gap: Existing finance benchmarks test agentic financial reasoning, data retrieval, and tool use through static, fully specified tasks. These one-off requests in clean contexts are not reflective of actual deployment. In reality, analysts must operate under ambiguity, gather key context from multiple (and sometimes conflicting) sources, and coordinate with leadership, colleagues across teams, and clients that may send relevant updates mid-task.
Our benchmark: Introducing ATLAS-Finance–100 expert-level tasks inside 13 realistic financial firm environments with multi-party working group teams, full inboxes, drives, notebooks, and calendars. Agent-generated work outputs are graded criterion-by-criterion against expert-authored rubrics with 42-486 binary criteria per task ( code ; data ).
Pushing the frontier: We designed ATLAS-Finance to closely reflect realistic working conditions in a financial firm. Agents that reliably navigate ATLAS-Finance are more prepared to help professionals deliver the real-world performance their clients expect.
What we found: Across 11 frontier models, the highest observed pass rate is 12.3% achieved by Claude Opus 5. Claude Fable 5.1 and GPT-6 Astra achieve 12.0% and 11.3%, respectively. The other eight models remain below 10%. Models consistently failed in three areas:
Applying the wrong financial logic, methodology, timing, or perimeter
Omitting required scope in the final output
Failing to propagate correctly calculated values downstream in the analyses.
The specific mix varies by model and workflow, but recurring, plausible-looking financial mistakes appear across all model families.
Why it matters: Identifying where agents fail under realistic working conditions like ATLAS-Finance environments is a key step in improving them to deliver larger productivity gains in the financial sector. In an actual bank, associates that make critical errors like the ones indicated above force senior members to audit all of their work. Similarly, if an agent makes a mistake like this in the workplace, it negates all of the time savings and expertise acceleration potentially provided by that model.
1. What’s different about ATLAS-Finance
Recent benchmarks have moved from short-form question answering to end-to-end financial work. Handshake’s BankerToolBench (BTB), in particular, evaluated the financial models produced by AI agents completing self-contained junior investment banker workflows. Its results show that frontier agents can complete parts of these workflows but remain unreliable at the standard required for professional delegation. BTB also identifies a remaining gap: simplified, static environments exclude the scattered, evolving, and sometimes contradictory nature of information found in the real world, in addition to organizational-level complexities that must be navigated where this work happens.
We address this gap with ATLAS-Finance by changing the unit of evaluation from a linear workflow to a workplace world. As a vastly richer RL environment, each world contains an advisory firm, specialized teams, updates from senior and junior professionals, clients and counterparties, active and background mandates, communication histories, shared files, calendars, permissions, and firm-specific conventions. The agent works through email, chat, document repositories, notes, calendars, and a local workstation. It must produce the requested financial artifacts and meet an explicit submission contract.
Table 1. Comparison of ATLAS Finance against existing finance and professional-work benchmarks. BTB (BankerToolBench, Handshake AI), APEX (APEX-Agents, Mercor; investment-banking subset), BFB (BigFinanceBench, Rogo), ALE (Agents' Last Exam, UC Berkeley), FAB (Finance Agent Benchmark, Vals AI), FinB (FinanceBench, Patronus AI), FF (FrontierFinance, Samaya AI), and GDPval (OpenAI).
Built by bankers up to VP-level from institutional banks and elite boutiques (such as Morgan Stanley, Bank of America, UBS, and Rothschild & Co), ATLAS-Finance reflects the real-world complexity that professionals navigate every day. The benchmark comprises 100 expert-level tasks in 13 environments, covering 9 industries and 20+ financial-analysis methodologies. These tasks typically take human finance experts 15-30 hours to complete manually.
Table 2. Expert-level workflows that undergird the ATLAS Finance benchmark.
To accurately complete tasks, agents must navigate organizational and relational complexity while collaborating with dozens of unique coworkers and clients. Before it populates a single spreadsheet cell, the agent must:
Assemble its workflow details from environment context
Understand its relationship to each counterparty and the role they play for the job at hand
Resolve controlled contradictions by finding governing sources or identifying the right decision-maker (e.g. “Who has the final say on the financing assumption?”).
A subset of tasks includes a real-time update with corrections to prior assumptions that must be incorporated before delivery.
Agent outputs are evaluated against 14,600 environment-grounded rubric criteria (146 per task on average and all expert-verified). ATLAS-Finance grades the work the way an industry professional would: every task carries 42-486 binary checks organized by schedule, with gating dependencies between them–if the debt schedule is not built with live formulas, everything downstream of it is recorded as unreachable, not merely wrong. A failing run produces a diagnosis in addition to a rubric score: which schedule broke, what the break cost downstream, and, read against the run's transcript, whether the model didn't know the finance, knew it and built it wrong, or built it right and delivered a workbook with static values instead of working formulas. Each rubric criterion comes with a weight indicating its relative importance according to experts. To grade any agent output against these rubrics, we use Gandalf the Grader : an agent-as-judge framework for verifying financial artifacts.
We evaluate 11 frontier models (see Figure 1), all run within the OpenCode agentic harness. Claude Opus 5 performs the best, yet still only manages to pass 12.3% of tasks . Claude Fable 5.1 and GPT-6 Astra are close behind, but the other eight models have significantly lower pass rates.
What counts as a pass: A run passes only if it meets every rubric criterion deemed critical (those weighted 3 or higher) and triggers no penalties (negative criteria). Penalties catch unacceptable issues such as fabricated data or a missing audit trail.
Figure 1. Pass@1 on ATLAS Finance. Error bars show 95% confidence intervals computed via bootstrap.
3. Example Task - LBO Rebuild Staffing at a FIG & media advisory
Below is an example task showcasing what agents must do in ATLAS-Finance environments. Here Project Osprey is a full LBO rebuild involving a specific tax treatment. It would take most human analysts 15+ hours to complete just the Excel modeling involved in this task, not counting the additional fact-finding work necessary to pull the relevant information into the analysis from the bank environment.
All ATLAS-Finance tasks start with the agent’s project “staffing”–it is assigned to a new project at the firm exactly how thousands of industry professionals experience it daily. The agent receives a heads-up message with high-level context that points to a subsequent project details email. This begins the workflow chain of understanding, discovery, analysis, and delivery. Understanding this kickoff notice is the first in dozens of steps an agent must handle to complete this ATLAS Finance task.
Step 1: Agent Staffing Kickoff
Look through your inbox and the Project Osprey attached source materials and rebuild the US sponsor LBO model your engagement lead assigned you (a sponsor’s 2019 LBO of a payments target, with the 338(h)(10) election). Your assignment email is titled “Project Osprey – rebuild the 2019 sponsor LBO (338(h)(10) election)”.
The brief goes on to fix the delivery contract that the agent must satisfy–an exact filename, exact recipients, an exact subject line– and the working standard: the finished workbook must carry live in-cell formulas so hardcoding a computed number where a formula belongs is an error even when the number is right.
3.1. Running Opus 5 powered Agent on this task (73 execution steps, outcome: Fail, rubric score: 0.336)
Here are the steps that a Claude agent took when assigned this task.
Steps 1–19 – Discovery. Searches the assignment subject (1) → reads the assignment thread and seven related threads (2–4, 14) → finds the desk’s OneNote: Modelling Approach, Output Standards, call notes, open items (5–6, 15–19) → lists and downloads every attachment (7–11) → dumps every source workbook programmatically to read real cell values (12–13).
Steps 20–25 – The shadow model. Before touching a single spreadsheet cell, the agent writes the complete LBO in pure Python: sources and uses, debt sweep, tax computation, returns, and iterates it to convergence (20–23), then extends it with a second validation pass (24–25). Every later self-check will be measured against this blueprint.
Steps 26–51 – Build. A single build script generates the whole workbook - nine tabs in the briefed order (26–27); first recalculation with the environment’s spreadsheet engine (28); targeted fixes to the NOL rows and a revenue formula (34–41). When convergence stalls, the agent reads the spreadsheet-recalculation utility’s own source code and reproduces its environment to debug the circularity (43–49). “Converged.” (52)
Steps 52–67 – Validate. Sweeps every output tab and ties the numbers back to the shadow model (52–67). All thirteen integrity checks on the Control sheet read OK; the workbook carries 2,641 live formulas and zero error cells.
Steps 68–73 – Delivery. Verifies the collection copy and the delivery-named copy are byte-identical by SHA-256 (68); sends to the exact recipients with the exact subject and attachment (69); reopens the sent message to confirm the attachment landed (70–72); closes: “Done. The model is built, recalculated, saved and delivered.” (73)
Figure 2. Anatomy of one agent rollout: 73 steps of work, two wrong ideas, and where the rubric deducted points (out of 1460).
During discovery, the agent surfaces a key email from the client’s tax lead. R. Calloway lays out each input for the tax-basis build and annotates each one with its operation: “ added to purchase price, ” “ subtracted in the basis calc, ” “ subtracted .” One line carries no operation at all: Existing Goodwill . The client gave Existing Goodwill no specified treatment, because under the record’s method Existing Goodwill plays no role in the amortizable step-up.
The agent’s main error is born at step 20, inside the agent’s own verification machinery. The Python shadow model–written to protect the build–contains one silent line that subtracts existing goodwill from the amortizable basis. Critically, the model failed to add a comment, consider alternative treatments given the task context, or flag the line for review. Six steps later, transcribing the inputs into the workbook, the agent rewrites the label itself: “ Existing Goodwill – carry-over tax basis ”–a treatment the governing email never assigned is recorded as if it were source data.
Figure 3. The LLM invents an incorrect treatment for the Existing Goodwill value shared in the client’s email. This treatment breaks the financial methodology for this scenario.
From that moment, the run’s guardrails point the wrong way. All of the agent’s self-checks verified against this shadow model, so all thirteen integrity checks read green,

[truncated]
