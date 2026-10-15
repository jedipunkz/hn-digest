---
source: "https://github.com/AssemblyAI/voice-standard-bench"
hn_url: "https://news.ycombinator.com/item?id=50107279"
title: "Show HN: Voice Standard Bench – Open benchmark of LLMs on speech transcripts"
article_title: "GitHub - AssemblyAI/voice-standard-bench: AssemblyAIs Open Source Benchmarking Framework · GitHub"
image: "https://opengraph.githubassets.com/33fcdebd17f9e8ff47327a02b8e6b9e6b72033bf8390dc25ce207389db56f607/AssemblyAI/voice-standard-bench"
author: "iharnoor"
captured_at: "2026-10-15T18:02:30Z"
capture_tool: "hn-digest"
hn_id: 50107279
score: 1
comments: 0
posted_at: "2026-10-15T17:41:52Z"
tags:
  - hacker-news
---

# Show HN: Voice Standard Bench – Open benchmark of LLMs on speech transcripts

- HN: [50107279](https://news.ycombinator.com/item?id=50107279)
- Source: [github.com](https://github.com/AssemblyAI/voice-standard-bench)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T17:41:52Z

## Translation

Title: Show HN: Voice Standard Bench – Open benchmark of LLMs on speech transcripts
Article title: GitHub - AssemblyAI/voice-standard-bench: AssemblyAIs Open Source Benchmarking Framework · GitHub
Description: AssemblyAIs Open Source Benchmarking Framework. Contribute to AssemblyAI/voice-standard-bench development by creating an account on GitHub.

Article text:
GitHub - AssemblyAI/voice-standard-bench: AssemblyAIs Open Source Benchmarking Framework · GitHub
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
AssemblyAI
voice-standard-bench Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
AssemblyAIs Open Source Benchmarking Framework
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
internal internal results/ canonical/ 0-1-3 results/ canonical/ 0-1-3 src src .gitignore .gitignore LICENSE LICENSE README.md README.md go.mod go.mod View all files Repository files navigation
A Go CLI that benchmarks every model behind llm_gateway on four tasks —
summarization, information extraction, text correction, and dictation —
against real AssemblyAI call transcripts (or, for dictation, a synthetic
dictation dataset — see "Dictation" below), tracking latency and quality
broken down by audio-length size. Summarization and extraction are graded
reference-free (there's no gold summary or annotated entity list to compare
against); text correction is graded the same way, by a single reference-free
LLM judge (its dataset's original transcript is not used for scoring — see
"What it measures").
Dictation also has a real answer key, and — unlike every other task — is
graded ENTIRELY mechanically, with no LLM judge call at all (see
"Dictation").
export ASSEMBLY_AI_TOKEN=... # set in your environment; there is deliberately no -api-key flag
go run .
Run these commands from the src/ directory ( cd src ): the normalized/ ,
results/ , results_ttft/ , errors/ , and errors_ttft/ directories live
one level up ( ../normalized etc., at the repo root). Tests run with
go test ./... from the same directory.
Run with -h for the full flag list — every flag has a one-line description
there rather than duplicated here. The API key is read only from the
ASSEMBLY_AI_TOKEN environment variable (there is no -api-key flag, so the
key never appears in the process list or shell history). GET /v1/models
needs no key at all; running completions does.
For every (model × task × transcript) combination:
Latency : total wall-clock latency and tokens/second (completion tokens ÷
total latency), each reported as min/max/mean/p50/p95 ( Stats in
src ). Time-to-first-token is also tracked but is always identical
to total latency: RunCompletion calls the gateway synchronously (a single
request, a single response), not via streaming, so there is no earlier
moment to measure a first token against — see RunCompletion 's doc comment
in gateway.go . The field is kept rather than dropped so the report JSON schema stays
stable; expect the TTFT and latency columns to
read identically for every row.
Quality : an LLM judge scores the output against the SOURCE TRANSCRIPT for
every task — reference-free, because summarization and extraction have no
gold summary or annotated entity list to compare against. See
src for the full rationale and the papers each rubric is
drawn from. text_correction is graded the same way : the candidate
model corrects missed_text , a simulated corrupted transcript, and one
reference-free judge call grades the output against missed_text
( faithfulness , improvement ). The dataset also carries the original
transcript in text , but it is not used for scoring : it turned out not to
be an accurate answer key, so any comparison with it — a word-level distance
( truth_wer / input_wer , accuracy , error_reduction ) or a second,
reference-based judge ( meaning_similarity ) — would score the dataset's
errors as much as the model's. Reports written with those scores keep them,
but summarizing ignores them and quality_score is built from the two judge
dimensions only; re-run text_correction to get a score that stands on its own
(see correctionSpec 's doc comment in prompts/correction.go ).
Every statistic is broken down by audio-length bucket ( 30s , 120s ,
900s , 3600s , ... — whatever's actually present in the loaded data; buckets
are read off each record, never hardcoded, so a new bucket the normalization
side adds later works with no change here). That's the whole point of
grouping by bucket rather than reporting one number per model: a model that's
fast and accurate on a 30-second call and falls over on an hour-long one is a
different tool than a report that just says "fast, accurate" would suggest.
Dictation is included in -tasks ( -tasks dictation , or alongside the other
three) and flows through the exact same staging and canonical-merge pipeline, but
is loaded and scored differently under the hood — see the package doc comment at the top of src for the full
picture.
Every record gets the exact same system prompt —
prompts.dictationSpec.Task ( src ), describing the
spoken-command grammar being tested ("comma" -> ",", "all caps <word> "
uppercases it, "new paragraph" starts a new paragraph, and so on) — exactly
like the other three tasks' fixed task prompt. Its records come from
normalized/dictation/*.jsonl (e.g. synthetic.jsonl ), never
aai_transcripts.jsonl — the schema and filename both differ. Each record also carries its own input.instruction field, but it
is deliberately never read (see rawDictationRecord in dictation.go ): the
one shared prompt above is the only one ever sent, regardless of what that
field happens to say.
What DOES still keep dictation off the shared runOne / runModel pipeline
is grading: every record already carries a real answer key
( expected.truth_text ), so grading is entirely mechanical — ScoreDictation
in src computes some or all of exact_match , word_error_rate ,
punctuation_accuracy , capitalization_accuracy , and
formatting_command_accuracy (whichever names that record's own metrics
field asks for) directly against truth_text , in code, with no LLM call
spent on it at all — unlike every other task, which always calls RunJudge .
Because it's scored mechanically, dictation's five score names aren't a 1-5
Likert judge rubric — they're percentages (0-100, except exact_match ,
which is a plain 0/1) already in the units you'd read out of a raw report.
They still fold into the same quality_score / task_quality breakdown every
other task's scores do (see "Per-task quality breakdown" below) —
src 's normalize just knows to invert word_error_rate (lower is
better) and scale the rest.
The standard benchmark's RunCompletion calls the gateway synchronously, so
its TimeToFirstToken is always identical to TotalLatency — see "What it
measures, and how" above. Every run therefore also includes a separate pass
that measures a real time-to-first-token via streaming ( RunCompletionStream
in src ), for the 30s and 120s audio-length buckets only — TTFT is a
real-time-adjacent responsiveness signal, not the full transcript-size sweep
the standard benchmark runs. It never judges (no LLM-judge call, no schema
check): a TTFT sample only measures how fast the first token came back,
never what the model said.
The two passes run sequentially, never concurrently: running them together
would send streaming AND synchronous requests to the same candidate models
at once, compounding against the gateway's per-account-per-model rate limit
(see NewGateway ) for BOTH passes' numbers, not just one. The TTFT pass has
its own everything besides the shared model-discovery/dataset-loading setup,
so it can never change a single number the standard pass produces: its own
sample/cell/report types ( src ) and its own staging directories
( results_ttft/ and errors_ttft/ , mirroring results/ and errors/ ,
never shared with them). A run's TTFT directory is named like its standard one
— results_ttft/<name> beside results/<name> — so one name finds everything
the run staged (see -replay below).
The TTFT reports are merged into the summary: a TTFT latency breakdown
( ttft_n through ttft_max ) appears in every summary.csv /table/JSON row by
(model, bucket) — new columns alongside the existing ones, blank for a bucket
the TTFT pass never measured (anything other than 30s/120s).
results/canonical/<version>/ (version folder name uses dashes, e.g. 0.1.2 →
0-1-2 ) is a single running result set that accumulates coverage across
however many separate runs it took to benchmark every model — unlike
results/ 's fresh, uniquely-named subfolder per run, it's the same
report-<judge>.json / ttft-report.json / summary.json / summary.csv files
updated in place. Appending is keyed by (model, reasoning effort, task, bucket) : new data for a combination entirely REPLACES old data for that same
combination (however many records the old run sampled for that bucket), and
every combination the new run didn't touch passes through byte-for-byte
unchanged — e.g. a run restricted to -buckets 3600s only replaces that
bucket's rows, leaving 30s / 120s / 900s 's already-merged data alone. See
src 's package doc comment for why a naive recompute-on-every-append is
wrong.
Every run does the whole thing in one command: it runs the standard benchmark
AND the TTFT pass back-to-back (sequentially, never concurrently, for the same
rate-limit reason as above) for -models x -tasks , stages both under
results/ and results_ttft/ , prints a combined review summary so a human
can sanity-check the numbers, then either merges into canonical immediately
( -approve-merge , for scripted/CI use) or asks for interactive confirmation
on stdin first — the same plan-then-confirm shape as terraform apply .
Declining leaves both staged directories on disk untouched. A successful merge
re-runs the summary code over the canonical version directory afterward, so its
summary.json / summary.csv reflect the merge immediately.
go run . \
-models new-model-1,new-model-2
Stashed task outputs ( task_runs/ )
Every model call costs money, and the judge step is the part that gets
reworked — a new rubric, a different judge model, a judge that was down — so
redoing the judging shouldn't mean paying for the model calls again. Every call
that comes back is therefore written to disk the moment it returns, before it is
judged, in the run's staged directory beside its report — and the TTFT pass's
streaming calls the same way, in the directory of the same name under
results_ttft/ :
results/<name>/ <name> is <date>-<uuid>
report-<judge>.json
task_runs/<task>/<model>__<reasoning effort>__<bucket>__<record id>.json
results_ttft/<name>/
ttft-report.json
task_runs/<task>/<model>__<reasoning effort>__<bucket>__<record id>.json
One file per (model, reasoning effort, task, bucket, record). Characters a file
name can't carry (the / in a model ID, the : in a record ID) become _ in
the name only. Each file holds who it was for ( model , reasoning_effort ,
task , bucket , r

[truncated]

## Original Extract

AssemblyAIs Open Source Benchmarking Framework. Contribute to AssemblyAI/voice-standard-bench development by creating an account on GitHub.

GitHub - AssemblyAI/voice-standard-bench: AssemblyAIs Open Source Benchmarking Framework · GitHub
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
AssemblyAI
voice-standard-bench Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
AssemblyAIs Open Source Benchmarking Framework
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
internal internal results/ canonical/ 0-1-3 results/ canonical/ 0-1-3 src src .gitignore .gitignore LICENSE LICENSE README.md README.md go.mod go.mod View all files Repository files navigation
A Go CLI that benchmarks every model behind llm_gateway on four tasks —
summarization, information extraction, text correction, and dictation —
against real AssemblyAI call transcripts (or, for dictation, a synthetic
dictation dataset — see "Dictation" below), tracking latency and quality
broken down by audio-length size. Summarization and extraction are graded
reference-free (there's no gold summary or annotated entity list to compare
against); text correction is graded the same way, by a single reference-free
LLM judge (its dataset's original transcript is not used for scoring — see
"What it measures").
Dictation also has a real answer key, and — unlike every other task — is
graded ENTIRELY mechanically, with no LLM judge call at all (see
"Dictation").
export ASSEMBLY_AI_TOKEN=... # set in your environment; there is deliberately no -api-key flag
go run .
Run these commands from the src/ directory ( cd src ): the normalized/ ,
results/ , results_ttft/ , errors/ , and errors_ttft/ directories live
one level up ( ../normalized etc., at the repo root). Tests run with
go test ./... from the same directory.
Run with -h for the full flag list — every flag has a one-line description
there rather than duplicated here. The API key is read only from the
ASSEMBLY_AI_TOKEN environment variable (there is no -api-key flag, so the
key never appears in the process list or shell history). GET /v1/models
needs no key at all; running completions does.
For every (model × task × transcript) combination:
Latency : total wall-clock latency and tokens/second (completion tokens ÷
total latency), each reported as min/max/mean/p50/p95 ( Stats in
src ). Time-to-first-token is also tracked but is always identical
to total latency: RunCompletion calls the gateway synchronously (a single
request, a single response), not via streaming, so there is no earlier
moment to measure a first token against — see RunCompletion 's doc comment
in gateway.go . The field is kept rather than dropped so the report JSON schema stays
stable; expect the TTFT and latency columns to
read identically for every row.
Quality : an LLM judge scores the output against the SOURCE TRANSCRIPT for
every task — reference-free, because summarization and extraction have no
gold summary or annotated entity list to compare against. See
src for the full rationale and the papers each rubric is
drawn from. text_correction is graded the same way : the candidate
model corrects missed_text , a simulated corrupted transcript, and one
reference-free judge call grades the output against missed_text
( faithfulness , improvement ). The dataset also carries the original
transcript in text , but it is not used for scoring : it turned out not to
be an accurate answer key, so any comparison with it — a word-level distance
( truth_wer / input_wer , accuracy , error_reduction ) or a second,
reference-based judge ( meaning_similarity ) — would score the dataset's
errors as much as the model's. Reports written with those scores keep them,
but summarizing ignores them and quality_score is built from the two judge
dimensions only; re-run text_correction to get a score that stands on its own
(see correctionSpec 's doc comment in prompts/correction.go ).
Every statistic is broken down by audio-length bucket ( 30s , 120s ,
900s , 3600s , ... — whatever's actually present in the loaded data; buckets
are read off each record, never hardcoded, so a new bucket the normalization
side adds later works with no change here). That's the whole point of
grouping by bucket rather than reporting one number per model: a model that's
fast and accurate on a 30-second call and falls over on an hour-long one is a
different tool than a report that just says "fast, accurate" would suggest.
Dictation is included in -tasks ( -tasks dictation , or alongside the other
three) and flows through the exact same staging and canonical-merge pipeline, but
is loaded and scored differently under the hood — see the package doc comment at the top of src for the full
picture.
Every record gets the exact same system prompt —
prompts.dictationSpec.Task ( src ), describing the
spoken-command grammar being tested ("comma" -> ",", "all caps <word> "
uppercases it, "new paragraph" starts a new paragraph, and so on) — exactly
like the other three tasks' fixed task prompt. Its records come from
normalized/dictation/*.jsonl (e.g. synthetic.jsonl ), never
aai_transcripts.jsonl — the schema and filename both differ. Each record also carries its own input.instruction field, but it
is deliberately never read (see rawDictationRecord in dictation.go ): the
one shared prompt above is the only one ever sent, regardless of what that
field happens to say.
What DOES still keep dictation off the shared runOne / runModel pipeline
is grading: every record already carries a real answer key
( expected.truth_text ), so grading is entirely mechanical — ScoreDictation
in src computes some or all of exact_match , word_error_rate ,
punctuation_accuracy , capitalization_accuracy , and
formatting_command_accuracy (whichever names that record's own metrics
field asks for) directly against truth_text , in code, with no LLM call
spent on it at all — unlike every other task, which always calls RunJudge .
Because it's scored mechanically, dictation's five score names aren't a 1-5
Likert judge rubric — they're percentages (0-100, except exact_match ,
which is a plain 0/1) already in the units you'd read out of a raw report.
They still fold into the same quality_score / task_quality breakdown every
other task's scores do (see "Per-task quality breakdown" below) —
src 's normalize just knows to invert word_error_rate (lower is
better) and scale the rest.
The standard benchmark's RunCompletion calls the gateway synchronously, so
its TimeToFirstToken is always identical to TotalLatency — see "What it
measures, and how" above. Every run therefore also includes a separate pass
that measures a real time-to-first-token via streaming ( RunCompletionStream
in src ), for the 30s and 120s audio-length buckets only — TTFT is a
real-time-adjacent responsiveness signal, not the full transcript-size sweep
the standard benchmark runs. It never judges (no LLM-judge call, no schema
check): a TTFT sample only measures how fast the first token came back,
never what the model said.
The two passes run sequentially, never concurrently: running them together
would send streaming AND synchronous requests to the same candidate models
at once, compounding against the gateway's per-account-per-model rate limit
(see NewGateway ) for BOTH passes' numbers, not just one. The TTFT pass has
its own everything besides the shared model-discovery/dataset-loading setup,
so it can never change a single number the standard pass produces: its own
sample/cell/report types ( src ) and its own staging directories
( results_ttft/ and errors_ttft/ , mirroring results/ and errors/ ,
never shared with them). A run's TTFT directory is named like its standard one
— results_ttft/<name> beside results/<name> — so one name finds everything
the run staged (see -replay below).
The TTFT reports are merged into the summary: a TTFT latency breakdown
( ttft_n through ttft_max ) appears in every summary.csv /table/JSON row by
(model, bucket) — new columns alongside the existing ones, blank for a bucket
the TTFT pass never measured (anything other than 30s/120s).
results/canonical/<version>/ (version folder name uses dashes, e.g. 0.1.2 →
0-1-2 ) is a single running result set that accumulates coverage across
however many separate runs it took to benchmark every model — unlike
results/ 's fresh, uniquely-named subfolder per run, it's the same
report-<judge>.json / ttft-report.json / summary.json / summary.csv files
updated in place. Appending is keyed by (model, reasoning effort, task, bucket) : new data for a combination entirely REPLACES old data for that same
combination (however many records the old run sampled for that bucket), and
every combination the new run didn't touch passes through byte-for-byte
unchanged — e.g. a run restricted to -buckets 3600s only replaces that
bucket's rows, leaving 30s / 120s / 900s 's already-merged data alone. See
src 's package doc comment for why a naive recompute-on-every-append is
wrong.
Every run does the whole thing in one command: it runs the standard benchmark
AND the TTFT pass back-to-back (sequentially, never concurrently, for the same
rate-limit reason as above) for -models x -tasks , stages both under
results/ and results_ttft/ , prints a combined review summary so a human
can sanity-check the numbers, then either merges into canonical immediately
( -approve-merge , for scripted/CI use) or asks for interactive confirmation
on stdin first — the same plan-then-confirm shape as terraform apply .
Declining leaves both staged directories on disk untouched. A successful merge
re-runs the summary code over the canonical version directory afterward, so its
summary.json / summary.csv reflect the merge immediately.
go run . \
-models new-model-1,new-model-2
Stashed task outputs ( task_runs/ )
Every model call costs money, and the judge step is the part that gets
reworked — a new rubric, a different judge model, a judge that was down — so
redoing the judging shouldn't mean paying for the model calls again. Every call
that comes back is therefore written to disk the moment it returns, before it is
judged, in the run's staged directory beside its report — and the TTFT pass's
streaming calls the same way, in the directory of the same name under
results_ttft/ :
results/<name>/ <name> is <date>-<uuid>
report-<judge>.json
task_runs/<task>/<model>__<reasoning effort>__<bucket>__<record id>.json
results_ttft/<name>/
ttft-report.json
task_runs/<task>/<model>__<reasoning effort>__<bucket>__<record id>.json
One file per (model, reasoning effort, task, bucket, record). Characters a file
name can't carry (the / in a model ID, the : in a record ID) become _ in
the name only. Each file holds who it was for ( model , reasoning_effort ,
task , bucket , r

[truncated]
