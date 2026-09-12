---
source: "https://withspecific.com/benchmarks/real-swe"
hn_url: "https://news.ycombinator.com/item?id=49676820"
title: "Real-SWE: Benchmarking AI models on private, real-world, enterprise codebases"
article_title: "Real-SWE Benchmark — Specific Labs"
image: ""
author: "theanonymousone"
captured_at: "2026-09-12T21:22:08Z"
capture_tool: "hn-digest"
hn_id: 49676820
score: 8
comments: 1
posted_at: "2026-09-12T20:25:48Z"
tags:
  - hacker-news
---

# Real-SWE: Benchmarking AI models on private, real-world, enterprise codebases

- HN: [49676820](https://news.ycombinator.com/item?id=49676820)
- Source: [withspecific.com](https://withspecific.com/benchmarks/real-swe)
- Score: 8
- Comments: 1
- Posted: 2026-09-12T20:25:48Z

## Translation

Title: Real-SWE: Benchmarking AI models on private, real-world, enterprise codebases
Article title: Real-SWE Benchmark — Specific Labs
Description: Real-SWE benchmarks frontier AI models on private production codebases licensed from real companies. Eight model and harness configurations, ten tasks, 640 scored rollouts.

Article text:
Real-SWE Benchmark — Specific Labs ← Benchmarks September 2026 Introducing Real-SWE
Benchmarking frontier AI models on private, real-world, enterprise codebases.
.
;+;+;:;+;+;.#
; ;
; . . . . ; .
+ ;
; . . : @;;+;+# +;+;+;+;:;+:#
: : ; . ; ; ; @;+;+.+;:;+;+*
; . . : ; . . : ; @++:;:;+.;;@ ;. . . . .: . ; . . . .; @++;;;;+;:;@
; : ; ; @;+++:@;++;:+ ; : : . :+.+;;;+:;@ ; : + .
; . . . ; ; . ; : . . . . ; . . . . .; ; . . . . .+ + . . . .; ; . . . . .; ; . . :
; : ; ; ; ; : ; . ; . ; ; ; ; :
; . . % . . ; : . . . . ; ; ... . . .; +. . . . .: ;. . :.. .: : . . ::. . . ; . . . .;
: *. @ ; ; +.@ . ; : +: *+ : ; :. %+; +. ; *. *# @: :.*. ;
; . . .;::@;; . . ; :#@;. ..@:. . ;.@%: .:@%:: :. . #+. :@%;. .:@*. .: @+ . #*. . ;@#.;.* . . .:
.::: ..:%@@+@+#... :+:%@@+..:@@@::*.:;@@:..+@@:; *.::#@@#:..:@@#...+@#: :*:%@@+ .:.+::..*@@:@@@#;.. ::.#
+@#.@ .*@: .*@%: *: :@%. ;@@: .@%. @ .;@@: .: :@#..#@*.
+@#.@ .@ .*@%:. *: %+ ;@@: #; +# .;@@: .. @: .#@*.
@. .. ;@ .*@%: *: @@ ;@@: %# * # .;@@: @+ .@
:*#.. @:: @. . ::;. @@ .. . @% .*@ ;@.
%.@ .. .@; . @@ .. .@@. .. @:#
... . ;%# . @@ +;#: .. .@ @
. ... ... .. ..
... .. ... .
01 Introduction
Today we are releasing Real-SWE, a benchmark that evaluates frontier AI models on private, real-world, enterprise codebases. Each task comes from a private production codebase that we licensed from a real-world company. These are problems their engineers work on, with all the context and complexity that comes with an existing product.
Private codebases. Agents must navigate proprietary systems whose code and solutions aren’t available on the public internet.
Work with business consequences. Getting billing right, calculating taxes, migrating customers. Changes that affect how a business runs, often across multiple services.
Company-specific complexity. Every company has its own rules and ways of writing code. Agents have to understand those conventions and make changes that work with what’s already there.
Can a coding agent actually do the work of a software engineer in the real world?
1 Fable 5.1 Claude Code Resolution rate: 38.8%
2 GPT-6 Astra Codex CLI Resolution rate: 33.8%
3 Gemini 3.8 Flash Gemini CLI Resolution rate: 31.2%
4 GLM 5.3 Claude Code Resolution rate: 28.8%
=5 Grok 4.6 Grok Build Resolution rate: 23.8%
=5 Muse Spark 1.3 Muse Code Resolution rate: 23.8%
7 Kimi K3 Kimi Code Resolution rate: 18.8%
8 GPT-5.6 Sol Codex CLI Resolution rate: 16.2%
Expert-generated or synthetic tasks can be well designed, but they aren’t the verbatim, actual tasks that engineers in real companies need to do. Our tasks differ on two axes: the underlying coding artifact and specificity of the instruction. Both add complexities that challenge today’s frontier models.
We use native harnesses to reflect how enterprise engineers work in practice, evaluating model-and-harness combinations rather than models in isolation.
Real company tasks require company-specific context
Correct billing depends on business rules and external services
Fix invoice billing so each business charges the right tax and exempt customers aren't taxed.
Billing reopens on Monday and every invoice this service issues is coming out untaxed. Each business on the platform settles its tax a different way: some maintain a rate themselves, some want each invoice priced against the buyer's destination by our tax authority provider, and some collect nothing at all, while a customer we hold an exemption for is charged nothing whichever way its business is configured. Pricing a destination means going to the authority with both addresses, the priced lines and the product category that business sells under, on the sandbox or the production authority according to the account the business is on; an address the authority refuses must be reported without stopping the invoice. The rate, the tax and the gross belong on the issued invoice, and once an invoice is settled the sale is filed back to the authority under that invoice's number so the returns reconcile. Invoices between European parties show both sides' VAT registrations. The authority and ledger are available at TAX_JAR_URL , PROD_TAX_JAR_URL and INFLUX_URL .
Agents work across code, infrastructure, and business tools
Tools and services across Real-SWE task environments. Each task exposes only the services its workflow needs.
We selected codebases through a rigorous screening process, focusing on real companies with substantial usage, strong engineering teams, and demanding production workloads. The sample tasks analyzed below come from these codebases, including:
A Luma/Partiful competitor with 200K+ users and a top 100 App Store ranking
A consumer fintech platform processing 100K+ bank statements
Enterprise AI sales platforms supporting complex business workflows
We prioritize code written to meet an actual user or business need over code written solely to create a benchmark task. Production engineering requires understanding existing architecture, preserving behavior that users rely on, and making changes within real operational constraints.
Brief instructions can require changes across many files
Our tasks describe the change needed, leaving agents to discover implementation details in the codebase and surrounding tools. Any behavior required by the verifier must be stated or reasonably discoverable. This leads to our prompts being slightly underspecified, about par with DeepSWE and Terminal Bench, but specific enough to not omit instructions.
The work is cross-functional and complex: a single change can span multiple parts of the application. Agents must understand existing business logic and company coding patterns while keeping the surrounding system working.
A typical Real-SWE instruction is 1,742 characters.
11 files in Real-SWE, compared with 6 in FrontierCode and DeepSWE.
Models fail even in short rollouts.
71.4 % of rollouts under 10 minutes failed, compared with 73.4 % of longer rollouts.
Triaging multiple systems and understanding requirements in codebases riddled with existing business logic and coding patterns is difficult.
73.4 % 26.6 % 398 / 542 failed
Every task is inspired or lifted verbatim from a private, real-world codebase. We find these types of tasks super interesting for three reasons:
Tasks on private codebases are natively out of distribution. These types of coding tasks are not available anywhere on the internet and are unlikely to have ever been trained on by any other ai model. 99% of tokens in real-world enterprises are hidden away from the frontier models.
These tasks are economically viable work. Each task here has a direct relationship to spend and was assigned to an engineer earning a salary. Most benchmarks test interesting, experimental capabilities that are often unlikely to be widespread in the real-world.
Company-specific engineering patterns matter. Does AI code match the bar of a real-world enterprise? Our results show us that we're far from that reality. Many enterprises care about code standards and patterns. We've found that today's models are weaker at understanding company coding patterns and frequently miss requirements or don't verify their assumptions.
Here's an analysis of a small sample of tasks from our benchmark. If you're interested in the sample, request access here .
6 of 10 tasks have resolution rates below 15%
Select a task to view model results. Percentages show the overall resolution rate.
Missed requirements are the most common failure
Failures are grouped by observed submission behavior using the same taxonomy across models, following DeepSWE .
One square per rollout: each row is a task, each column a trial, eight trials per task for every model.
Percentages are out of each model's failed runs, not all runs.
Builds on a guess about the system instead of checking it in the workspace.
GPT-5.6 Sol 43.3% : 29 of 67 failed runs
GPT-6 Astra 34.0% : 18 of 53 failed runs
GLM 5.3 28.1% : 16 of 57 failed runs
Grok 4.6 24.6% : 15 of 61 failed runs
Fable 5.1 24.5% : 12 of 49 failed runs
Muse Spark 1.3 19.7% : 12 of 61 failed runs
Kimi K3 15.4% : 10 of 65 failed runs
Gemini 3.8 Flash 10.9% : 6 of 55 failed runs
Leaves out behavior the instruction requires.
Grok 4.6 67.2% : 41 of 61 failed runs
Kimi K3 53.8% : 35 of 65 failed runs
GLM 5.3 38.6% : 22 of 57 failed runs
Fable 5.1 36.7% : 18 of 49 failed runs
Muse Spark 1.3 36.1% : 22 of 61 failed runs
GPT-5.6 Sol 31.3% : 21 of 67 failed runs
Gemini 3.8 Flash 29.1% : 16 of 55 failed runs
GPT-6 Astra 28.3% : 15 of 53 failed runs
Right idea, wired into the surrounding system incorrectly.
Gemini 3.8 Flash 49.1% : 27 of 55 failed runs
Muse Spark 1.3 41.0% : 25 of 61 failed runs
Fable 5.1 34.7% : 17 of 49 failed runs
GPT-6 Astra 34.0% : 18 of 53 failed runs
Kimi K3 27.7% : 18 of 65 failed runs
GLM 5.3 26.3% : 15 of 57 failed runs
GPT-5.6 Sol 16.4% : 11 of 67 failed runs
Grok 4.6 8.2% : 5 of 61 failed runs
Breaks existing behavior while making the change.
Gemini 3.8 Flash 10.9% : 6 of 55 failed runs
GPT-5.6 Sol 9.0% : 6 of 67 failed runs
Fable 5.1 4.1% : 2 of 49 failed runs
GPT-6 Astra 3.8% : 2 of 53 failed runs
Muse Spark 1.3 3.3% : 2 of 61 failed runs
GLM 5.3 0% : 0 of 57 failed runs
Grok 4.6 0% : 0 of 61 failed runs
Kimi K3 0% : 0 of 65 failed runs
Delivers the change somewhere the running application never calls, such as a one-off script.
GLM 5.3 7.0% : 4 of 57 failed runs
Kimi K3 3.1% : 2 of 65 failed runs
Fable 5.1 0% : 0 of 49 failed runs
GPT-6 Astra 0% : 0 of 53 failed runs
Gemini 3.8 Flash 0% : 0 of 55 failed runs
Grok 4.6 0% : 0 of 61 failed runs
Muse Spark 1.3 0% : 0 of 61 failed runs
GPT-5.6 Sol 0% : 0 of 67 failed runs
Higher cost does not guarantee a higher resolution rate
10 15 20 25 30 35 40 45 $2 $3 $5 $10 Cost per rollout (USD, log scale) Gemini 3.8 Flash: 31.2% · $2.50; Gemini CLI Gemini 3.8 Flash 31.2% · $2.50 GPT-5.6 Sol: 16.2% · $2.65; Codex CLI GPT-5.6 Sol 16.2% · $2.65 Muse Spark 1.3: 23.8% · $2.74; Muse Code Muse Spark 1.3 23.8% · $2.74 Grok 4.6: 23.8% · $3.44; Grok Build; incomplete usage, actual cost may be higher Grok 4.6 23.8% · $3.44 Kimi K3: 18.8% · $3.90; Kimi Code; incomplete usage, actual cost may be higher Kimi K3 18.8% · $3.90 GPT-6 Astra: 33.8% · $4.67; Codex CLI GPT-6 Astra 33.8% · $4.67 GLM 5.3: 28.8% · $5.12; Claude Code GLM 5.3 28.8% · $5.12 Fable 5.1: 38.8% · $6.96; Claude Code Fable 5.1 38.8% · $6.96 Resolution rate (%) 10 15 20 25 30 35 40 45 $2 $3 $5 $10 Cost / rollout ($, log scale) Gemini 3.8 Flash: 31.2% · $2.50; Gemini CLI 3 GPT-5.6 Sol: 16.2% · $2.65; Codex CLI 8 Muse Spark 1.3: 23.8% · $2.74; Muse Code 6 Grok 4.6: 23.8% · $3.44; Grok Build; incomplete usage, actual cost may be higher 5 Kimi K3: 18.8% · $3.90; Kimi Code; incomplete usage, actual cost may be higher 7 GPT-6 Astra: 33.8% · $4.67; Codex CLI 2 GLM 5.3: 28.8% · $5.12; Claude Code 4 Fable 5.1: 38.8% · $6.96; Claude Code 1 1 Fable 5.1 38.8% · $6.96
3 Gemini 3.8 Flash 31.2% · $2.50
6 Muse Spark 1.3 23.8% · $2.74
Estimated rollout costs range from $2.50 to $6.96
Swipe the chart to see all tasks.
0 100k 200k 300k 400k 01 02 03 04 05 06 07 08 09 10 task Entitlement overage lines · Fable 5.1: 34k Multi-region sweep · Fable 5.1: 30k Tax jurisdiction · Fable 5.1: 78k API token metering · Fable 5.1: 95k API keys & environments · Fable 5.1: 71k S3 datastore measurement · Fable 5.1: 62k Customer identity migration · Fable 5.1: 67k Billing schedule migration · Fable 5.1: 26k Linearizable scan · Fable 5.1: 86k Analytics stream reducer · Fable 5.1: 88k Entitlement overage lines · GPT-6 Astra: 13k Multi-region sweep · GPT-6 Astra: 13k Tax jurisdiction · GPT-6 Astra: 24k API token metering · GPT-6 Astra: 31k API keys & environments · GPT-6 Astra: 32k S3 datastore measurement · GPT-6 Astra: 22k Customer identity migration · GPT-6 Astra: 25k Billing schedule migration · GPT-6 Astra: 15k Linearizable scan · GPT-6 Astra: 33k Analytics stream reducer · GPT-6 Astra: 29k Entitlement overage lines · Gemini 3.8 Flash: 78k Multi-region sweep · Gemini 3.8 Flash: 67k Tax jurisdictio

[truncated]

## Original Extract

Real-SWE benchmarks frontier AI models on private production codebases licensed from real companies. Eight model and harness configurations, ten tasks, 640 scored rollouts.

Real-SWE Benchmark — Specific Labs ← Benchmarks September 2026 Introducing Real-SWE
Benchmarking frontier AI models on private, real-world, enterprise codebases.
.
;+;+;:;+;+;.#
; ;
; . . . . ; .
+ ;
; . . : @;;+;+# +;+;+;+;:;+:#
: : ; . ; ; ; @;+;+.+;:;+;+*
; . . : ; . . : ; @++:;:;+.;;@ ;. . . . .: . ; . . . .; @++;;;;+;:;@
; : ; ; @;+++:@;++;:+ ; : : . :+.+;;;+:;@ ; : + .
; . . . ; ; . ; : . . . . ; . . . . .; ; . . . . .+ + . . . .; ; . . . . .; ; . . :
; : ; ; ; ; : ; . ; . ; ; ; ; :
; . . % . . ; : . . . . ; ; ... . . .; +. . . . .: ;. . :.. .: : . . ::. . . ; . . . .;
: *. @ ; ; +.@ . ; : +: *+ : ; :. %+; +. ; *. *# @: :.*. ;
; . . .;::@;; . . ; :#@;. ..@:. . ;.@%: .:@%:: :. . #+. :@%;. .:@*. .: @+ . #*. . ;@#.;.* . . .:
.::: ..:%@@+@+#... :+:%@@+..:@@@::*.:;@@:..+@@:; *.::#@@#:..:@@#...+@#: :*:%@@+ .:.+::..*@@:@@@#;.. ::.#
+@#.@ .*@: .*@%: *: :@%. ;@@: .@%. @ .;@@: .: :@#..#@*.
+@#.@ .@ .*@%:. *: %+ ;@@: #; +# .;@@: .. @: .#@*.
@. .. ;@ .*@%: *: @@ ;@@: %# * # .;@@: @+ .@
:*#.. @:: @. . ::;. @@ .. . @% .*@ ;@.
%.@ .. .@; . @@ .. .@@. .. @:#
... . ;%# . @@ +;#: .. .@ @
. ... ... .. ..
... .. ... .
01 Introduction
Today we are releasing Real-SWE, a benchmark that evaluates frontier AI models on private, real-world, enterprise codebases. Each task comes from a private production codebase that we licensed from a real-world company. These are problems their engineers work on, with all the context and complexity that comes with an existing product.
Private codebases. Agents must navigate proprietary systems whose code and solutions aren’t available on the public internet.
Work with business consequences. Getting billing right, calculating taxes, migrating customers. Changes that affect how a business runs, often across multiple services.
Company-specific complexity. Every company has its own rules and ways of writing code. Agents have to understand those conventions and make changes that work with what’s already there.
Can a coding agent actually do the work of a software engineer in the real world?
1 Fable 5.1 Claude Code Resolution rate: 38.8%
2 GPT-6 Astra Codex CLI Resolution rate: 33.8%
3 Gemini 3.8 Flash Gemini CLI Resolution rate: 31.2%
4 GLM 5.3 Claude Code Resolution rate: 28.8%
=5 Grok 4.6 Grok Build Resolution rate: 23.8%
=5 Muse Spark 1.3 Muse Code Resolution rate: 23.8%
7 Kimi K3 Kimi Code Resolution rate: 18.8%
8 GPT-5.6 Sol Codex CLI Resolution rate: 16.2%
Expert-generated or synthetic tasks can be well designed, but they aren’t the verbatim, actual tasks that engineers in real companies need to do. Our tasks differ on two axes: the underlying coding artifact and specificity of the instruction. Both add complexities that challenge today’s frontier models.
We use native harnesses to reflect how enterprise engineers work in practice, evaluating model-and-harness combinations rather than models in isolation.
Real company tasks require company-specific context
Correct billing depends on business rules and external services
Fix invoice billing so each business charges the right tax and exempt customers aren't taxed.
Billing reopens on Monday and every invoice this service issues is coming out untaxed. Each business on the platform settles its tax a different way: some maintain a rate themselves, some want each invoice priced against the buyer's destination by our tax authority provider, and some collect nothing at all, while a customer we hold an exemption for is charged nothing whichever way its business is configured. Pricing a destination means going to the authority with both addresses, the priced lines and the product category that business sells under, on the sandbox or the production authority according to the account the business is on; an address the authority refuses must be reported without stopping the invoice. The rate, the tax and the gross belong on the issued invoice, and once an invoice is settled the sale is filed back to the authority under that invoice's number so the returns reconcile. Invoices between European parties show both sides' VAT registrations. The authority and ledger are available at TAX_JAR_URL , PROD_TAX_JAR_URL and INFLUX_URL .
Agents work across code, infrastructure, and business tools
Tools and services across Real-SWE task environments. Each task exposes only the services its workflow needs.
We selected codebases through a rigorous screening process, focusing on real companies with substantial usage, strong engineering teams, and demanding production workloads. The sample tasks analyzed below come from these codebases, including:
A Luma/Partiful competitor with 200K+ users and a top 100 App Store ranking
A consumer fintech platform processing 100K+ bank statements
Enterprise AI sales platforms supporting complex business workflows
We prioritize code written to meet an actual user or business need over code written solely to create a benchmark task. Production engineering requires understanding existing architecture, preserving behavior that users rely on, and making changes within real operational constraints.
Brief instructions can require changes across many files
Our tasks describe the change needed, leaving agents to discover implementation details in the codebase and surrounding tools. Any behavior required by the verifier must be stated or reasonably discoverable. This leads to our prompts being slightly underspecified, about par with DeepSWE and Terminal Bench, but specific enough to not omit instructions.
The work is cross-functional and complex: a single change can span multiple parts of the application. Agents must understand existing business logic and company coding patterns while keeping the surrounding system working.
A typical Real-SWE instruction is 1,742 characters.
11 files in Real-SWE, compared with 6 in FrontierCode and DeepSWE.
Models fail even in short rollouts.
71.4 % of rollouts under 10 minutes failed, compared with 73.4 % of longer rollouts.
Triaging multiple systems and understanding requirements in codebases riddled with existing business logic and coding patterns is difficult.
73.4 % 26.6 % 398 / 542 failed
Every task is inspired or lifted verbatim from a private, real-world codebase. We find these types of tasks super interesting for three reasons:
Tasks on private codebases are natively out of distribution. These types of coding tasks are not available anywhere on the internet and are unlikely to have ever been trained on by any other ai model. 99% of tokens in real-world enterprises are hidden away from the frontier models.
These tasks are economically viable work. Each task here has a direct relationship to spend and was assigned to an engineer earning a salary. Most benchmarks test interesting, experimental capabilities that are often unlikely to be widespread in the real-world.
Company-specific engineering patterns matter. Does AI code match the bar of a real-world enterprise? Our results show us that we're far from that reality. Many enterprises care about code standards and patterns. We've found that today's models are weaker at understanding company coding patterns and frequently miss requirements or don't verify their assumptions.
Here's an analysis of a small sample of tasks from our benchmark. If you're interested in the sample, request access here .
6 of 10 tasks have resolution rates below 15%
Select a task to view model results. Percentages show the overall resolution rate.
Missed requirements are the most common failure
Failures are grouped by observed submission behavior using the same taxonomy across models, following DeepSWE .
One square per rollout: each row is a task, each column a trial, eight trials per task for every model.
Percentages are out of each model's failed runs, not all runs.
Builds on a guess about the system instead of checking it in the workspace.
GPT-5.6 Sol 43.3% : 29 of 67 failed runs
GPT-6 Astra 34.0% : 18 of 53 failed runs
GLM 5.3 28.1% : 16 of 57 failed runs
Grok 4.6 24.6% : 15 of 61 failed runs
Fable 5.1 24.5% : 12 of 49 failed runs
Muse Spark 1.3 19.7% : 12 of 61 failed runs
Kimi K3 15.4% : 10 of 65 failed runs
Gemini 3.8 Flash 10.9% : 6 of 55 failed runs
Leaves out behavior the instruction requires.
Grok 4.6 67.2% : 41 of 61 failed runs
Kimi K3 53.8% : 35 of 65 failed runs
GLM 5.3 38.6% : 22 of 57 failed runs
Fable 5.1 36.7% : 18 of 49 failed runs
Muse Spark 1.3 36.1% : 22 of 61 failed runs
GPT-5.6 Sol 31.3% : 21 of 67 failed runs
Gemini 3.8 Flash 29.1% : 16 of 55 failed runs
GPT-6 Astra 28.3% : 15 of 53 failed runs
Right idea, wired into the surrounding system incorrectly.
Gemini 3.8 Flash 49.1% : 27 of 55 failed runs
Muse Spark 1.3 41.0% : 25 of 61 failed runs
Fable 5.1 34.7% : 17 of 49 failed runs
GPT-6 Astra 34.0% : 18 of 53 failed runs
Kimi K3 27.7% : 18 of 65 failed runs
GLM 5.3 26.3% : 15 of 57 failed runs
GPT-5.6 Sol 16.4% : 11 of 67 failed runs
Grok 4.6 8.2% : 5 of 61 failed runs
Breaks existing behavior while making the change.
Gemini 3.8 Flash 10.9% : 6 of 55 failed runs
GPT-5.6 Sol 9.0% : 6 of 67 failed runs
Fable 5.1 4.1% : 2 of 49 failed runs
GPT-6 Astra 3.8% : 2 of 53 failed runs
Muse Spark 1.3 3.3% : 2 of 61 failed runs
GLM 5.3 0% : 0 of 57 failed runs
Grok 4.6 0% : 0 of 61 failed runs
Kimi K3 0% : 0 of 65 failed runs
Delivers the change somewhere the running application never calls, such as a one-off script.
GLM 5.3 7.0% : 4 of 57 failed runs
Kimi K3 3.1% : 2 of 65 failed runs
Fable 5.1 0% : 0 of 49 failed runs
GPT-6 Astra 0% : 0 of 53 failed runs
Gemini 3.8 Flash 0% : 0 of 55 failed runs
Grok 4.6 0% : 0 of 61 failed runs
Muse Spark 1.3 0% : 0 of 61 failed runs
GPT-5.6 Sol 0% : 0 of 67 failed runs
Higher cost does not guarantee a higher resolution rate
10 15 20 25 30 35 40 45 $2 $3 $5 $10 Cost per rollout (USD, log scale) Gemini 3.8 Flash: 31.2% · $2.50; Gemini CLI Gemini 3.8 Flash 31.2% · $2.50 GPT-5.6 Sol: 16.2% · $2.65; Codex CLI GPT-5.6 Sol 16.2% · $2.65 Muse Spark 1.3: 23.8% · $2.74; Muse Code Muse Spark 1.3 23.8% · $2.74 Grok 4.6: 23.8% · $3.44; Grok Build; incomplete usage, actual cost may be higher Grok 4.6 23.8% · $3.44 Kimi K3: 18.8% · $3.90; Kimi Code; incomplete usage, actual cost may be higher Kimi K3 18.8% · $3.90 GPT-6 Astra: 33.8% · $4.67; Codex CLI GPT-6 Astra 33.8% · $4.67 GLM 5.3: 28.8% · $5.12; Claude Code GLM 5.3 28.8% · $5.12 Fable 5.1: 38.8% · $6.96; Claude Code Fable 5.1 38.8% · $6.96 Resolution rate (%) 10 15 20 25 30 35 40 45 $2 $3 $5 $10 Cost / rollout ($, log scale) Gemini 3.8 Flash: 31.2% · $2.50; Gemini CLI 3 GPT-5.6 Sol: 16.2% · $2.65; Codex CLI 8 Muse Spark 1.3: 23.8% · $2.74; Muse Code 6 Grok 4.6: 23.8% · $3.44; Grok Build; incomplete usage, actual cost may be higher 5 Kimi K3: 18.8% · $3.90; Kimi Code; incomplete usage, actual cost may be higher 7 GPT-6 Astra: 33.8% · $4.67; Codex CLI 2 GLM 5.3: 28.8% · $5.12; Claude Code 4 Fable 5.1: 38.8% · $6.96; Claude Code 1 1 Fable 5.1 38.8% · $6.96
3 Gemini 3.8 Flash 31.2% · $2.50
6 Muse Spark 1.3 23.8% · $2.74
Estimated rollout costs range from $2.50 to $6.96
Swipe the chart to see all tasks.
0 100k 200k 300k 400k 01 02 03 04 05 06 07 08 09 10 task Entitlement overage lines · Fable 5.1: 34k Multi-region sweep · Fable 5.1: 30k Tax jurisdiction · Fable 5.1: 78k API token metering · Fable 5.1: 95k API keys & environments · Fable 5.1: 71k S3 datastore measurement · Fable 5.1: 62k Customer identity migration · Fable 5.1: 67k Billing schedule migration · Fable 5.1: 26k Linearizable scan · Fable 5.1: 86k Analytics stream reducer · Fable 5.1: 88k Entitlement overage lines · GPT-6 Astra: 13k Multi-region sweep · GPT-6 Astra: 13k Tax jurisdiction · GPT-6 Astra: 24k API token metering · GPT-6 Astra: 31k API keys & environments · GPT-6 Astra: 32k S3 datastore measurement · GPT-6 Astra: 22k Customer identity migration · GPT-6 Astra: 25k Billing schedule migration · GPT-6 Astra: 15k Linearizable scan · GPT-6 Astra: 33k Analytics stream reducer · GPT-6 Astra: 29k Entitlement overage lines · Gemini 3.8 Flash: 78k Multi-region sweep · Gemini 3.8 Flash: 67k Tax jurisdictio

[truncated]
