---
source: "https://code.claude.com/docs/en/plugin-evals"
hn_url: "https://news.ycombinator.com/item?id=49666190"
title: "Anthropic released a CLI to evaluate skills and plugins"
article_title: "Test plugins with evals - Claude Code Docs"
image: "https://claude-code.mintlify.app/_next/image?url=%2F_mintlify%2Fapi%2Fog%3Fdivision%3DPlugins%26title%3DTest%2Bplugins%2Bwith%2Bevals%26description%3DWrite%2Beval%2Bcases%2Bfor%2Byour%2BClaude%2BCode%2Bplugin%252C%2Brun%2Bthem%2Bwith%2Bclaude%2Bplugin%2Beval%252C%2Bgrade%2Bthe%2Bresults%252C%2Bcompare%2Bagainst%2Ba%2Bno-plugin%2Bbaseline%252C%2Band%2Bgate%2BCI%2Bon%2Bthe%2Bscor%26theme%3D03628e99c753a03aec319053&w=1200&q=100"
author: "edonadei"
captured_at: "2026-09-11T23:00:59Z"
capture_tool: "hn-digest"
hn_id: 49666190
score: 3
comments: 1
posted_at: "2026-09-11T22:20:10Z"
tags:
  - hacker-news
---

# Anthropic released a CLI to evaluate skills and plugins

- HN: [49666190](https://news.ycombinator.com/item?id=49666190)
- Source: [code.claude.com](https://code.claude.com/docs/en/plugin-evals)
- Score: 3
- Comments: 1
- Posted: 2026-09-11T22:20:10Z

## Translation

Title: Anthropic released a CLI to evaluate skills and plugins
Article title: Test plugins with evals - Claude Code Docs
Description: Write eval cases for your Claude Code plugin, run them with claude plugin eval, grade the results, compare against a no-plugin baseline, and gate CI on the score.

Article text:
Test plugins with evals - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant ⌘ I Claude Developer Platform
Search... Navigation Plugins Test plugins with evals Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Agents and parallel work
Isolate sessions with worktrees
Discover and install prebuilt plugins
Share session output as artifacts
Push external events to Claude
Troubleshoot installation and login
Troubleshoot performance and stability
How an eval run works What happens in a run
Write and refine cases Write a case manually
Set run limits and tools in prompt.md
Choose graders that give a stable signal
Score against the no-plugin baseline
Use a different eval directory
Set up fixtures and mocks Seed the workspace or conversation
Run evals Choose what to evaluate
What a run can access Trust the plugin directory
Eval suite reference prompt.md frontmatter
Troubleshooting ”plugin eval is currently in early access”
”plugin eval is currently unavailable”
”is not a trusted plugin directory, and this run cannot stop to ask you about it”
The baseline arm shows no plugin, or delta is zero
Everything scores zero although the right files were produced
A regex over the trace doesn’t match text I can see
Tools are denied, MCP tools are missing, or Bash won’t run
The run exits 1 but the results look fine
”—json output path must end in .json”
A grader shows passed: false under a run that scored 1.0
Runs fail with a usage-limit or rate-limit error partway through
Runs time out or hit the turn cap
Copy page Copy page Write eval cases for your Claude Code plugin, run them with claude plugin eval, grade the results, compare against a no-plugin baseline, and gate CI on the score.
Copy page Copy page claude plugin eval runs your plugin against a suite of test cases and scores the results. Each case is a realistic prompt plus one or more graders. A grader is a pass/fail check on what Claude produced, such as a regex over the reply, whether a particular tool was called, or a rubric that a second model judges the reply against.
You don’t have to write the suite manually. claude plugin eval init asks you about your plugin, proposes the cases and graders, tries them, and writes the files. You can also ask Claude to do the same from a session you already have open.
Use evals to measure how reliably your plugin steers Claude to the right outcome, to catch regressions when you change the plugin or a new model ships, and to see what the plugin contributes compared with no plugin at all.
This page is for plugin and skill authors who have a working plugin and want to test its behavior, and for teams that gate plugin changes in CI. Its case format is separate from the evals/evals.json file the skill-creator plugin uses. To create a plugin, see Create plugins ; to check a plugin’s files for syntax and schema errors rather than its behavior, use claude plugin validate .
Every eval run and every judge grader is a real model call on your account, counted against your plan’s usage or your API bill, so check the requirements first. Then create your first eval suite , or go to Run evals in CI if you already have one.
​ Requirements
Claude Code v2.1.269 or later. Run claude --version to check and claude update to upgrade.
A plugin directory with a plugin.json or .claude-plugin/plugin.json manifest, or a skills-directory plugin .
The same authentication and model provider your normal Claude Code sessions use. Eval runs, judge-scored graders, and claude plugin eval init call the model with your credentials, so they count against your plan’s usage limits or your API bill. When the command reports a cost, the figure is a list-price estimate of those calls.
​ Create your first eval suite
Claude Code v2.1.269 or later and the other requirements
A terminal open at your plugin’s root directory, the one containing plugin.json or .claude-plugin/plugin.json
One skill in the plugin you want to test, and a request a user would type that should trigger it
claude plugin eval init
If Claude Code doesn’t already trust this directory it first asks Trust this plugin directory? ; answer y . An interactive Claude Code session then opens. Claude reads your plugin and asks you what a good result looks like, proposes prompts that should and shouldn’t trigger the plugin, designs graders for each, pilots them once to check they behave, and writes one case directory per prompt under evals/ , each named after its prompt. When Claude tells you the suite is ready, exit that session with /exit or Ctrl+D to return to your shell. If you already have a Claude Code session open at the plugin root, you can instead ask Claude there to run claude plugin eval init . Claude runs the command and then asks you the same questions in that conversation. If you’d rather write a case yourself to see exactly what the files contain, follow Write a case manually and come back here to run it. 2 Run the suite
claude plugin eval .
You already trusted this directory during step 1, so the run starts immediately. If you wrote the case manually instead, the run first asks Trust this plugin directory? [y/N] ; answer y . What a run can access explains what you’re agreeing to. Each case runs three times with your plugin and three times without it, so one case is six runs. A progress line prints as each run finishes, with that run’s score and each grader’s verdict. 3 Read the summary
CASE WITH W/OUT Δ RUNS COST NOTES
first-case 1.00 0.33 +0.67 6 $0.41
1 case(s) · mean Δ +0.67 · 74s · $0.41
Report: /Users/you/my-plugin/evals/results/2026-09-10T17-02-11-482Z/report.html
Published: https://claude.ai/... · keep local next time with --no-publish
WITH is the case’s score with your plugin loaded, W/OUT is the score without it, and a positive Δ means the plugin raised the score. COST is a list-price estimate of the model calls, and NOTES shows the highest-weight failing grader’s explanation, or the run’s error, from the with-arm. 4 Open the report and iterate
claude plugin eval . --case < case-nam e > --runs 1 --ablation none
Replace <case-name> with one of the directory names under evals/ .
​ Write and refine cases
my-plugin/
├── .claude-plugin/plugin.json
├── skills/...
└── evals/
├── first-case/
│ ├── prompt.md # frontmatter: case fields; body: the prompt
│ ├── graders/
│ │ ├── criteria.md # frontmatter: type + options; body: rubric or pattern
│ │ └── skill-fired.md
│ └── case.yaml # optional: only for context.* fields
├── ignores-unrelated-request/
│ └── ...
└── results/ # written by each run; add to .gitignore
​ Write a case manually
claude plugin eval init --bare first-case
evals/first-case/
├── prompt.md # the prompt sent to Claude, plus run limits
└── graders/
└── criteria.md # one grader: how to score the result
In prompt.md you write the message Claude receives in each run, and set the run’s limits and the tools the case may use in its frontmatter. Open evals/first-case/prompt.md and replace the placeholder body with a request one of your skills should handle, phrased the way a user would type it rather than naming the skill. This example is for a skill that drafts commit messages; use your own request:
---
max_turns : 10
allowed_tools : [ Read , Glob , Grep , Skill ]
---
Write me a commit message for this change: I renamed getUser to fetchUser and updated the three call sites.
Each run starts in an empty working directory, so put whatever the task needs in the prompt itself, or set up the workspace first. The full list of frontmatter fields covers the model, timeout, tags, and environment variables.
Each file under graders/ is one check applied after the run. Open evals/first-case/graders/criteria.md and replace the placeholder with a rubric for the judge model, written as concrete PASS and FAIL conditions:
---
type : llm
---
PASS if < what a correct response contains > .
FAIL if < what a wrong or missing response looks like > .
Then add a second grader that checks whether your skill is what produced the answer. Create evals/first-case/graders/skill-fired.md , replacing your-skill-name with the name from your skill’s SKILL.md :
---
type : tool_used
tool : Skill
input_match : '"skill"\s*:\s*"(?:[\w-]+:)?your-skill-name"'
---
This passes when Claude invoked that skill at least once during the run, including by its namespaced plugin-name:skill-name form. Grader types lists the other checks available, such as matching a regex or confirming a file was created.
With both files saved, run the case the way the quickstart does, with claude plugin eval . from the plugin root.
​ Set run limits and to
[truncated]
Choose graders that give a stable signal
An llm grader asks a model for a verdict, so its answer can differ between runs, and it differs more the longer the text it has to read. These habits keep a suite’s scores steady enough to trust:
For long output such as a generated file, grade it with a regex grader over the file’s contents, which checks the whole file the same way every time. Keep llm graders for short outputs, with rubrics written as concrete PASS and FAIL conditions.
Give each case one grader on the result, such as the final message or a produced file, and one on how Claude got there, such as tool_used or tool_order . Together they tell you both whether the answer was right and whether your plugin produced it.
If a case’s tool_used: Skill grader passes but Δ is negative, suspect the judge before the plugin. A small judge model can mark a correct answer wrong because it’s formatted differently from what the rubric describes. Re-run with --judge-model sonnet , and tighten the rubric so formatting doesn’t decide the verdict.
To check that a build or test passed inside the run, have the prompt ask Claude to run it and write the outcome to a file, grade that file, and assert the command ran with a tool_used grader whose input_match names the command.
​ Score against the no-plugin baseline
Every tool_used grader whose tool is Skill
Any grader you mark arm: with-only
​ Use a different eval directory
In plugin.json : add "experimental": { "evals": "quality/evals" } .
On the command line : pass --eval-dir quality/evals to both claude plugin eval and claude plugin eval init .
​ Seed the workspace or conversation
schema_version : "1.1"
name : changelog-from-diff
tags : [ smoke ]
context :
scaffold_script : fixture.sh
add_dirs : [ resources ]
​ Mock MCP servers
---
expect :
title : string
priority : [ low , medium , high ]
---
Created issue #4821: {{input.title}}
Insert fields from the call’s input with {{input.<field>}} , and the contents of a fixture file beside the mock with {{file:fixtures/{input.<field>}.json}} . The expect: block guards the input. If a call violates it, the run aborts with score 0 and records why, so a case can assert what your plugin asked the server to do. Set error: true to return the body as a tool error instead, or type: agent to have a small model answer as the server from instructions in the body. The mock file reference lists every key and the _server.md and _tools.json files.
To grade the calls themselves, point a grader at target: mock_calls .
To run against the plugin’s real MCP servers instead, pass one of these flags. Either way those processes run as you, outside the run’s sandbox, and their tools need an --allow-tools grant :
--allow-real-servers : start the real process for each server you haven’t mocked, and keep answering mocked tools from their files
--mocks off : ignore mocks/ entirely and start every server the plugin declares
Replay agent mock answers
A type: agent mock answers with a call to the --judge-model , so its output varies between runs and changes if you change the judge. When a run completes without an error or abort, Claude Code saves each answer an agent mock

[truncated]

## Original Extract

Write eval cases for your Claude Code plugin, run them with claude plugin eval, grade the results, compare against a no-plugin baseline, and gate CI on the score.

Test plugins with evals - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant ⌘ I Claude Developer Platform
Search... Navigation Plugins Test plugins with evals Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Agents and parallel work
Isolate sessions with worktrees
Discover and install prebuilt plugins
Share session output as artifacts
Push external events to Claude
Troubleshoot installation and login
Troubleshoot performance and stability
How an eval run works What happens in a run
Write and refine cases Write a case manually
Set run limits and tools in prompt.md
Choose graders that give a stable signal
Score against the no-plugin baseline
Use a different eval directory
Set up fixtures and mocks Seed the workspace or conversation
Run evals Choose what to evaluate
What a run can access Trust the plugin directory
Eval suite reference prompt.md frontmatter
Troubleshooting ”plugin eval is currently in early access”
”plugin eval is currently unavailable”
”is not a trusted plugin directory, and this run cannot stop to ask you about it”
The baseline arm shows no plugin, or delta is zero
Everything scores zero although the right files were produced
A regex over the trace doesn’t match text I can see
Tools are denied, MCP tools are missing, or Bash won’t run
The run exits 1 but the results look fine
”—json output path must end in .json”
A grader shows passed: false under a run that scored 1.0
Runs fail with a usage-limit or rate-limit error partway through
Runs time out or hit the turn cap
Copy page Copy page Write eval cases for your Claude Code plugin, run them with claude plugin eval, grade the results, compare against a no-plugin baseline, and gate CI on the score.
Copy page Copy page claude plugin eval runs your plugin against a suite of test cases and scores the results. Each case is a realistic prompt plus one or more graders. A grader is a pass/fail check on what Claude produced, such as a regex over the reply, whether a particular tool was called, or a rubric that a second model judges the reply against.
You don’t have to write the suite manually. claude plugin eval init asks you about your plugin, proposes the cases and graders, tries them, and writes the files. You can also ask Claude to do the same from a session you already have open.
Use evals to measure how reliably your plugin steers Claude to the right outcome, to catch regressions when you change the plugin or a new model ships, and to see what the plugin contributes compared with no plugin at all.
This page is for plugin and skill authors who have a working plugin and want to test its behavior, and for teams that gate plugin changes in CI. Its case format is separate from the evals/evals.json file the skill-creator plugin uses. To create a plugin, see Create plugins ; to check a plugin’s files for syntax and schema errors rather than its behavior, use claude plugin validate .
Every eval run and every judge grader is a real model call on your account, counted against your plan’s usage or your API bill, so check the requirements first. Then create your first eval suite , or go to Run evals in CI if you already have one.
​ Requirements
Claude Code v2.1.269 or later. Run claude --version to check and claude update to upgrade.
A plugin directory with a plugin.json or .claude-plugin/plugin.json manifest, or a skills-directory plugin .
The same authentication and model provider your normal Claude Code sessions use. Eval runs, judge-scored graders, and claude plugin eval init call the model with your credentials, so they count against your plan’s usage limits or your API bill. When the command reports a cost, the figure is a list-price estimate of those calls.
​ Create your first eval suite
Claude Code v2.1.269 or later and the other requirements
A terminal open at your plugin’s root directory, the one containing plugin.json or .claude-plugin/plugin.json
One skill in the plugin you want to test, and a request a user would type that should trigger it
claude plugin eval init
If Claude Code doesn’t already trust this directory it first asks Trust this plugin directory? ; answer y . An interactive Claude Code session then opens. Claude reads your plugin and asks you what a good result looks like, proposes prompts that should and shouldn’t trigger the plugin, designs graders for each, pilots them once to check they behave, and writes one case directory per prompt under evals/ , each named after its prompt. When Claude tells you the suite is ready, exit that session with /exit or Ctrl+D to return to your shell. If you already have a Claude Code session open at the plugin root, you can instead ask Claude there to run claude plugin eval init . Claude runs the command and then asks you the same questions in that conversation. If you’d rather write a case yourself to see exactly what the files contain, follow Write a case manually and come back here to run it. 2 Run the suite
claude plugin eval .
You already trusted this directory during step 1, so the run starts immediately. If you wrote the case manually instead, the run first asks Trust this plugin directory? [y/N] ; answer y . What a run can access explains what you’re agreeing to. Each case runs three times with your plugin and three times without it, so one case is six runs. A progress line prints as each run finishes, with that run’s score and each grader’s verdict. 3 Read the summary
CASE WITH W/OUT Δ RUNS COST NOTES
first-case 1.00 0.33 +0.67 6 $0.41
1 case(s) · mean Δ +0.67 · 74s · $0.41
Report: /Users/you/my-plugin/evals/results/2026-09-10T17-02-11-482Z/report.html
Published: https://claude.ai/... · keep local next time with --no-publish
WITH is the case’s score with your plugin loaded, W/OUT is the score without it, and a positive Δ means the plugin raised the score. COST is a list-price estimate of the model calls, and NOTES shows the highest-weight failing grader’s explanation, or the run’s error, from the with-arm. 4 Open the report and iterate
claude plugin eval . --case < case-nam e > --runs 1 --ablation none
Replace <case-name> with one of the directory names under evals/ .
​ Write and refine cases
my-plugin/
├── .claude-plugin/plugin.json
├── skills/...
└── evals/
├── first-case/
│ ├── prompt.md # frontmatter: case fields; body: the prompt
│ ├── graders/
│ │ ├── criteria.md # frontmatter: type + options; body: rubric or pattern
│ │ └── skill-fired.md
│ └── case.yaml # optional: only for context.* fields
├── ignores-unrelated-request/
│ └── ...
└── results/ # written by each run; add to .gitignore
​ Write a case manually
claude plugin eval init --bare first-case
evals/first-case/
├── prompt.md # the prompt sent to Claude, plus run limits
└── graders/
└── criteria.md # one grader: how to score the result
In prompt.md you write the message Claude receives in each run, and set the run’s limits and the tools the case may use in its frontmatter. Open evals/first-case/prompt.md and replace the placeholder body with a request one of your skills should handle, phrased the way a user would type it rather than naming the skill. This example is for a skill that drafts commit messages; use your own request:
---
max_turns : 10
allowed_tools : [ Read , Glob , Grep , Skill ]
---
Write me a commit message for this change: I renamed getUser to fetchUser and updated the three call sites.
Each run starts in an empty working directory, so put whatever the task needs in the prompt itself, or set up the workspace first. The full list of frontmatter fields covers the model, timeout, tags, and environment variables.
Each file under graders/ is one check applied after the run. Open evals/first-case/graders/criteria.md and replace the placeholder with a rubric for the judge model, written as concrete PASS and FAIL conditions:
---
type : llm
---
PASS if < what a correct response contains > .
FAIL if < what a wrong or missing response looks like > .
Then add a second grader that checks whether your skill is what produced the answer. Create evals/first-case/graders/skill-fired.md , replacing your-skill-name with the name from your skill’s SKILL.md :
---
type : tool_used
tool : Skill
input_match : '"skill"\s*:\s*"(?:[\w-]+:)?your-skill-name"'
---
This passes when Claude invoked that skill at least once during the run, including by its namespaced plugin-name:skill-name form. Grader types lists the other checks available, such as matching a regex or confirming a file was created.
With both files saved, run the case the way the quickstart does, with claude plugin eval . from the plugin root.
​ Set run limits and to
[truncated]
Choose graders that give a stable signal
An llm grader asks a model for a verdict, so its answer can differ between runs, and it differs more the longer the text it has to read. These habits keep a suite’s scores steady enough to trust:
For long output such as a generated file, grade it with a regex grader over the file’s contents, which checks the whole file the same way every time. Keep llm graders for short outputs, with rubrics written as concrete PASS and FAIL conditions.
Give each case one grader on the result, such as the final message or a produced file, and one on how Claude got there, such as tool_used or tool_order . Together they tell you both whether the answer was right and whether your plugin produced it.
If a case’s tool_used: Skill grader passes but Δ is negative, suspect the judge before the plugin. A small judge model can mark a correct answer wrong because it’s formatted differently from what the rubric describes. Re-run with --judge-model sonnet , and tighten the rubric so formatting doesn’t decide the verdict.
To check that a build or test passed inside the run, have the prompt ask Claude to run it and write the outcome to a file, grade that file, and assert the command ran with a tool_used grader whose input_match names the command.
​ Score against the no-plugin baseline
Every tool_used grader whose tool is Skill
Any grader you mark arm: with-only
​ Use a different eval directory
In plugin.json : add "experimental": { "evals": "quality/evals" } .
On the command line : pass --eval-dir quality/evals to both claude plugin eval and claude plugin eval init .
​ Seed the workspace or conversation
schema_version : "1.1"
name : changelog-from-diff
tags : [ smoke ]
context :
scaffold_script : fixture.sh
add_dirs : [ resources ]
​ Mock MCP servers
---
expect :
title : string
priority : [ low , medium , high ]
---
Created issue #4821: {{input.title}}
Insert fields from the call’s input with {{input.<field>}} , and the contents of a fixture file beside the mock with {{file:fixtures/{input.<field>}.json}} . The expect: block guards the input. If a call violates it, the run aborts with score 0 and records why, so a case can assert what your plugin asked the server to do. Set error: true to return the body as a tool error instead, or type: agent to have a small model answer as the server from instructions in the body. The mock file reference lists every key and the _server.md and _tools.json files.
To grade the calls themselves, point a grader at target: mock_calls .
To run against the plugin’s real MCP servers instead, pass one of these flags. Either way those processes run as you, outside the run’s sandbox, and their tools need an --allow-tools grant :
--allow-real-servers : start the real process for each server you haven’t mocked, and keep answering mocked tools from their files
--mocks off : ignore mocks/ entirely and start every server the plugin declares
Replay agent mock answers
A type: agent mock answers with a call to the --judge-model , so its output varies between runs and changes if you change the judge. When a run completes without an error or abort, Claude Code saves each answer an agent mock

[truncated]
