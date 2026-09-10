---
source: "https://automatedteach.com/p/designing-tests-for-agentic-ai-tools"
hn_url: "https://news.ycombinator.com/item?id=49647091"
title: "Designing Tests for Agentic AI Tools"
article_title: "Designing Tests for Agentic AI Tools"
image: "https://beehiiv-images-production.s3.amazonaws.com/uploads/asset/file/0690d258-e2a5-4208-a07f-2def3a4a723e/b7493f92-9635-4165-b4e7-077e1cfc8bc8.png?t=1789059732"
author: "professorgc"
captured_at: "2026-09-10T18:09:42Z"
capture_tool: "hn-digest"
hn_id: 49647091
score: 2
comments: 0
posted_at: "2026-09-10T17:11:12Z"
tags:
  - hacker-news
---

# Designing Tests for Agentic AI Tools

- HN: [49647091](https://news.ycombinator.com/item?id=49647091)
- Source: [automatedteach.com](https://automatedteach.com/p/designing-tests-for-agentic-ai-tools)
- Score: 2
- Comments: 0
- Posted: 2026-09-10T17:11:12Z

## Translation

Title: Designing Tests for Agentic AI Tools
Description: A passing test suite doesn't tell you how the agent got there

Article text:
Designing Tests for Agentic AI Tools AutomatED: Teaching Better with Tech Login Subscribe Department Course Scheduling Tool About Me Site Map 0 Designing Tests for Agentic AI Tools
A passing test suite doesn't tell you how the agent got there
Graham Clay
September 10, 2026 • Estimated Reading Time: 8 minutes
[image created with ChatGPT Images 2.0]
Today, a more technical note on the evaluation of agentic harnesses, drawing from the process of development for the Work Ledger , my harness-neutral agentic context management file system.
What did the Person Have to Do?
Follow One Run All the Way Through
Three Evidence Sources, Three Different Jobs
You ask a coding agent to update an API endpoint. The repository already contains a decision about backward compatibility: keep the existing response fields. You want the implementation, a contract test and a short migration note. Leave the commit to you.
The final diff looks good. The test passes. The migration note is in the right place.
But you also pointed the agent to the current decision after it started from an obsolete one, rejected a breaking field rename, approved the test command and reminded it to save the migration note.
This is an invented example, but it contains a real evaluation problem. If you score only the final diff, several different things disappear into “success”: repairs to delegated work, a permission decision, and a reminder to save the result.
You need to distinguish them before you can decide what to fix—or what to delegate next.
The method I use is to keep the evidence about the agent's actions separate from the evidence about the work it left behind , with my own contribution visible alongside both.
Agent-evaluation guidance already distinguishes a transcript from an outcome in the environment; Anthropic's guide explains that distinction well. Here I want to work through a particular complication: a person keeps changing the conditions and repairing the work (or, more generally, intervening) while the agent runs. What should the resulting record let us conclude? This is an especially important question if we are to avoid agent hand-holding that silently diminishes the productivity gains they are supposed to bring.
(Relevant context on how the resultant system, the Work Ledger, achieves the productivity gains in prior posts here , here , and here .)
What did the Person Have to Do?
An experienced engineer can make a shaky workflow look quite good. You catch the old decision, recognize the breaking change and ask the question that gets the agent unstuck. Because those interventions come naturally to you, you may barely count them as work.
Counting interruptions is a start, but it doesn't explain the engineering problem. An approval you deliberately reserved for a person is different from a correction to something the agent was supposed to handle.
I find these categories useful when reviewing a run. They are labels for diagnosis, not a universal scoring system:
A decision deliberately left to a person
Choose whether to accept a breaking API change.
Keep that decision with the intended owner.
Review the diff before merging.
Whether the review catches what it needs to catch.
Stop an unauthorized response-field rename.
The instructions, retrieved context and actual edit.
Help with the execution environment
Approve a test command or restore missing access.
Permissions and tooling; distinguish expected approvals from unexpected friction.
Ask for the migration note that appeared only in chat.
Where the workflow writes its result and how it verifies the save.
A single intervention can serve more than one role. Record what happened before choosing the labels. The point is to avoid treating every human action as an autonomy failure, or treating every rescue as ordinary review.
Follow One Run All the Way Through
Here's a more explicit version of the endpoint example. The files, commands and events below are synthetic. I haven't run this example or measured a model with it.
The task is to add an optional display_name field to a profile response while retaining name . Before editing, the agent must consult docs/adr/014-profile-response.md , which records that compatibility decision. It must update the contract test, save docs/profile-migration.md , and leave Git commits to the engineer.
Before starting, the engineer identifies three kinds of evidence: the activity record for the required sequence, the permission history for the test approval, and the final files and test output for the deliverable. The engineer also expects to approve the first test invocation in this environment.
Suppose the visible record contains these events:
01 agent: read docs/archive/profile-response-v0.md
02 agent: edit src/profile.py; replace name with display_name
03 human: point to ADR 014; instruct agent to keep name
04 agent: read docs/adr/014-profile-response.md
05 agent: edit src/profile.py; retain name and add display_name
06 agent: update tests/test_profile_contract.py
07 app: request permission for pytest tests/test_profile_contract.py
08 human: approve that command once, as planned
09 tool: test command exits 0
10 agent: describe migration note in chat; report task complete
11 human: request the missing docs/profile-migration.md file
12 agent: write docs/profile-migration.md
For this illustration, the engineer also preserves the diff before intervening at step 3, checks that the note is absent at step 10, and inspects the final files after step 12. The first attempt ends at the first corrective intervention; the rest is a recorded continuation. The final result can be perfectly usable. The record still supports several different judgments:
Consult the current decision before editing
The recorded read of ADR 014 comes after the first edit; the engineer redirects the agent.
Check how the workflow identifies the governing decision and makes it available before editing.
Preserve the existing response field
The first edit breaks the requirement; the engineer corrects it. The final diff preserves it.
Check the compatibility instruction and ensure the contract test actually protects the existing field.
The command runs after the planned approval and exits successfully.
Keep the intended approval; inspect what the test asserts before treating the pass as evidence of compatibility.
The file is absent at the agent's first completion report and present after a reminder.
Add a check of required deliverables before reporting completion.
The excerpt doesn't establish whether any Git commit occurred.
Inspect the available Git activity and saved history; keep any remaining event-coverage gap explicit.
The first two rows may share a cause. This trace doesn't establish whether the agent missed an instruction, retrieved the wrong context, or understood the decision and still edited incorrectly. It gives you a smaller place to investigate.
Nor do failed requirements, interventions and causes have to produce the same count. Here one correction addresses both the source and the field rename, and another addresses the missing note. Keep the actual record so those distinctions remain recoverable.
The useful conclusion is that the task was completed with corrective help and one planned permission approval. This run does not show that the workflow can satisfy those delegated requirements without that help. It also doesn't show that the permission approval was a defect.
Three Evidence Sources, Three Different Jobs
The starting materials tell you which task you tested. Keep the request, relevant files and starting revision. Git may omit local settings, so save the configuration that matters to the question as well. If you change the prompt, configuration or supplied context during recovery, record the change. The repaired run answers a different question from the unassisted first attempt.
The activity record tells you what you can observe about execution. Tool calls, visible approvals, errors and human messages can establish particular events and their order. They don't give you an all-seeing account of the runtime. In the synthetic trace, the late recorded read is visible. To claim that the agent had never received ADR 014 through any other route, you'd need to account for the starting context and other ways the app supplies material.
An absent read event is especially easy to overinterpret. The file might already have been in context, the event record might be incomplete, or the read might really have been skipped. If you can't distinguish those possibilities, retain the uncertainty.
The final files tell you what survived. Inspect the diff, deliverables and any saved revision. Equal before-and-after bytes don't prove that nothing happened between the snapshots. A file could have been changed and restored. Conversely, an agent can describe a correct result without writing it anywhere the next task will find it.
This matters for review, too. If a reviewer approves a particular diff and the agent then changes it, the earlier approval applies to the earlier candidate. Record a revision or a hash of the reviewed files so you can tell whether you're still looking at the same work. The hash identifies bytes; it doesn't certify their quality or explain how they were produced.
For a consequential check, I would start with something this small (n.b. the fields are an example template; tweak as needed):
question : " Can this workflow add display_name while preserving name?"
starting_revision : " <revision>"
request_and_context : " <saved request, instructions and supplied decisions>"
relevant_local_settings : " <saved files or hashes; unknowns listed>"
environment : " <app, engine/model, mode; unknown where unavailable>"
requirements :
- " consult the current ADR before editing"
- " preserve name; test compatibility; save the migration note; no commit"
planned_human_steps : [ " approve first test invocation" , " review final diff" ]
delegated_requirements : [ " choose current ADR" , " preserve name" , " save note" ]
first_attempt_boundary : " first completion, stop, or corrective intervention"
event_record : " <tool events, human messages, approvals; coverage gaps>"
interventions : " <event, purpose, expected or corrective, time if useful>"
final_candidate : " <diff/files and revision or content hashes>"
verdicts : " <procedure; permissions; saved result; remaining unknowns>"
Choose both the boundary and the division of responsibility before the run. Planned approvals can belong to that attempt; help repairing a predeclared delegated requirement starts a continuation. Classify the intervention against that earlier allocation. Otherwise you can quietly redefine a rescue as expected assistance. Preserve the first state when practical, then record the continuation separately. If you must intervene immediately to prevent harm, do so and record the stop and any missing snapshot.
For a small task, file copies and a short event note may be enough. If you are comparing repeated runs, preserve a fuller input inventory and use the same boundary each time. A time measure is useful only if you say what it includes: waiting for the tool, reading a permission dialog, reviewing the diff or repairing the result.
Before treating a failed requirement as an agent failure, check the test itself. Was the ADR actually supplied or retrievable? Does the contract test exercise the old response field, or does it merely pass because the new implementation and new test share the same mistaken assumption? A precise record of an invalid test is still an invalid test.
One run can reveal a failure path worth repairing. It doesn't estimate a reliability rate. And if you revise the workflow repeatedly, the eventual success shows what the revised arrangement achieved; it doesn't isolate which change caused the improvement.
I encountered a related problem while investigating Claude's Windows desktop app. In a September 4 report , I recorded a plain PowerShell command running silently under an allow rul

[truncated]

## Original Extract

A passing test suite doesn't tell you how the agent got there

Designing Tests for Agentic AI Tools AutomatED: Teaching Better with Tech Login Subscribe Department Course Scheduling Tool About Me Site Map 0 Designing Tests for Agentic AI Tools
A passing test suite doesn't tell you how the agent got there
Graham Clay
September 10, 2026 • Estimated Reading Time: 8 minutes
[image created with ChatGPT Images 2.0]
Today, a more technical note on the evaluation of agentic harnesses, drawing from the process of development for the Work Ledger , my harness-neutral agentic context management file system.
What did the Person Have to Do?
Follow One Run All the Way Through
Three Evidence Sources, Three Different Jobs
You ask a coding agent to update an API endpoint. The repository already contains a decision about backward compatibility: keep the existing response fields. You want the implementation, a contract test and a short migration note. Leave the commit to you.
The final diff looks good. The test passes. The migration note is in the right place.
But you also pointed the agent to the current decision after it started from an obsolete one, rejected a breaking field rename, approved the test command and reminded it to save the migration note.
This is an invented example, but it contains a real evaluation problem. If you score only the final diff, several different things disappear into “success”: repairs to delegated work, a permission decision, and a reminder to save the result.
You need to distinguish them before you can decide what to fix—or what to delegate next.
The method I use is to keep the evidence about the agent's actions separate from the evidence about the work it left behind , with my own contribution visible alongside both.
Agent-evaluation guidance already distinguishes a transcript from an outcome in the environment; Anthropic's guide explains that distinction well. Here I want to work through a particular complication: a person keeps changing the conditions and repairing the work (or, more generally, intervening) while the agent runs. What should the resulting record let us conclude? This is an especially important question if we are to avoid agent hand-holding that silently diminishes the productivity gains they are supposed to bring.
(Relevant context on how the resultant system, the Work Ledger, achieves the productivity gains in prior posts here , here , and here .)
What did the Person Have to Do?
An experienced engineer can make a shaky workflow look quite good. You catch the old decision, recognize the breaking change and ask the question that gets the agent unstuck. Because those interventions come naturally to you, you may barely count them as work.
Counting interruptions is a start, but it doesn't explain the engineering problem. An approval you deliberately reserved for a person is different from a correction to something the agent was supposed to handle.
I find these categories useful when reviewing a run. They are labels for diagnosis, not a universal scoring system:
A decision deliberately left to a person
Choose whether to accept a breaking API change.
Keep that decision with the intended owner.
Review the diff before merging.
Whether the review catches what it needs to catch.
Stop an unauthorized response-field rename.
The instructions, retrieved context and actual edit.
Help with the execution environment
Approve a test command or restore missing access.
Permissions and tooling; distinguish expected approvals from unexpected friction.
Ask for the migration note that appeared only in chat.
Where the workflow writes its result and how it verifies the save.
A single intervention can serve more than one role. Record what happened before choosing the labels. The point is to avoid treating every human action as an autonomy failure, or treating every rescue as ordinary review.
Follow One Run All the Way Through
Here's a more explicit version of the endpoint example. The files, commands and events below are synthetic. I haven't run this example or measured a model with it.
The task is to add an optional display_name field to a profile response while retaining name . Before editing, the agent must consult docs/adr/014-profile-response.md , which records that compatibility decision. It must update the contract test, save docs/profile-migration.md , and leave Git commits to the engineer.
Before starting, the engineer identifies three kinds of evidence: the activity record for the required sequence, the permission history for the test approval, and the final files and test output for the deliverable. The engineer also expects to approve the first test invocation in this environment.
Suppose the visible record contains these events:
01 agent: read docs/archive/profile-response-v0.md
02 agent: edit src/profile.py; replace name with display_name
03 human: point to ADR 014; instruct agent to keep name
04 agent: read docs/adr/014-profile-response.md
05 agent: edit src/profile.py; retain name and add display_name
06 agent: update tests/test_profile_contract.py
07 app: request permission for pytest tests/test_profile_contract.py
08 human: approve that command once, as planned
09 tool: test command exits 0
10 agent: describe migration note in chat; report task complete
11 human: request the missing docs/profile-migration.md file
12 agent: write docs/profile-migration.md
For this illustration, the engineer also preserves the diff before intervening at step 3, checks that the note is absent at step 10, and inspects the final files after step 12. The first attempt ends at the first corrective intervention; the rest is a recorded continuation. The final result can be perfectly usable. The record still supports several different judgments:
Consult the current decision before editing
The recorded read of ADR 014 comes after the first edit; the engineer redirects the agent.
Check how the workflow identifies the governing decision and makes it available before editing.
Preserve the existing response field
The first edit breaks the requirement; the engineer corrects it. The final diff preserves it.
Check the compatibility instruction and ensure the contract test actually protects the existing field.
The command runs after the planned approval and exits successfully.
Keep the intended approval; inspect what the test asserts before treating the pass as evidence of compatibility.
The file is absent at the agent's first completion report and present after a reminder.
Add a check of required deliverables before reporting completion.
The excerpt doesn't establish whether any Git commit occurred.
Inspect the available Git activity and saved history; keep any remaining event-coverage gap explicit.
The first two rows may share a cause. This trace doesn't establish whether the agent missed an instruction, retrieved the wrong context, or understood the decision and still edited incorrectly. It gives you a smaller place to investigate.
Nor do failed requirements, interventions and causes have to produce the same count. Here one correction addresses both the source and the field rename, and another addresses the missing note. Keep the actual record so those distinctions remain recoverable.
The useful conclusion is that the task was completed with corrective help and one planned permission approval. This run does not show that the workflow can satisfy those delegated requirements without that help. It also doesn't show that the permission approval was a defect.
Three Evidence Sources, Three Different Jobs
The starting materials tell you which task you tested. Keep the request, relevant files and starting revision. Git may omit local settings, so save the configuration that matters to the question as well. If you change the prompt, configuration or supplied context during recovery, record the change. The repaired run answers a different question from the unassisted first attempt.
The activity record tells you what you can observe about execution. Tool calls, visible approvals, errors and human messages can establish particular events and their order. They don't give you an all-seeing account of the runtime. In the synthetic trace, the late recorded read is visible. To claim that the agent had never received ADR 014 through any other route, you'd need to account for the starting context and other ways the app supplies material.
An absent read event is especially easy to overinterpret. The file might already have been in context, the event record might be incomplete, or the read might really have been skipped. If you can't distinguish those possibilities, retain the uncertainty.
The final files tell you what survived. Inspect the diff, deliverables and any saved revision. Equal before-and-after bytes don't prove that nothing happened between the snapshots. A file could have been changed and restored. Conversely, an agent can describe a correct result without writing it anywhere the next task will find it.
This matters for review, too. If a reviewer approves a particular diff and the agent then changes it, the earlier approval applies to the earlier candidate. Record a revision or a hash of the reviewed files so you can tell whether you're still looking at the same work. The hash identifies bytes; it doesn't certify their quality or explain how they were produced.
For a consequential check, I would start with something this small (n.b. the fields are an example template; tweak as needed):
question : " Can this workflow add display_name while preserving name?"
starting_revision : " <revision>"
request_and_context : " <saved request, instructions and supplied decisions>"
relevant_local_settings : " <saved files or hashes; unknowns listed>"
environment : " <app, engine/model, mode; unknown where unavailable>"
requirements :
- " consult the current ADR before editing"
- " preserve name; test compatibility; save the migration note; no commit"
planned_human_steps : [ " approve first test invocation" , " review final diff" ]
delegated_requirements : [ " choose current ADR" , " preserve name" , " save note" ]
first_attempt_boundary : " first completion, stop, or corrective intervention"
event_record : " <tool events, human messages, approvals; coverage gaps>"
interventions : " <event, purpose, expected or corrective, time if useful>"
final_candidate : " <diff/files and revision or content hashes>"
verdicts : " <procedure; permissions; saved result; remaining unknowns>"
Choose both the boundary and the division of responsibility before the run. Planned approvals can belong to that attempt; help repairing a predeclared delegated requirement starts a continuation. Classify the intervention against that earlier allocation. Otherwise you can quietly redefine a rescue as expected assistance. Preserve the first state when practical, then record the continuation separately. If you must intervene immediately to prevent harm, do so and record the stop and any missing snapshot.
For a small task, file copies and a short event note may be enough. If you are comparing repeated runs, preserve a fuller input inventory and use the same boundary each time. A time measure is useful only if you say what it includes: waiting for the tool, reading a permission dialog, reviewing the diff or repairing the result.
Before treating a failed requirement as an agent failure, check the test itself. Was the ADR actually supplied or retrievable? Does the contract test exercise the old response field, or does it merely pass because the new implementation and new test share the same mistaken assumption? A precise record of an invalid test is still an invalid test.
One run can reveal a failure path worth repairing. It doesn't estimate a reliability rate. And if you revise the workflow repeatedly, the eventual success shows what the revised arrangement achieved; it doesn't isolate which change caused the improvement.
I encountered a related problem while investigating Claude's Windows desktop app. In a September 4 report , I recorded a plain PowerShell command running silently under an allow rul

[truncated]
