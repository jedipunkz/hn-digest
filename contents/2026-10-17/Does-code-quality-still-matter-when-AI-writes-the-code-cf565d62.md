---
source: "https://docspring.com/blog/posts/does-code-quality-still-matter-when-ai-writes-the-code/"
hn_url: "https://news.ycombinator.com/item?id=50136214"
title: "Does code quality still matter when AI writes the code?"
article_title: "Does code quality still matter when AI writes the code?"
image: "https://docspring.com/blog/content/images/size/w1200/2026/10/01-does-code-quality-still-matter-when-ai-writes-the-code-3x2.webp"
author: "nathan_f77"
captured_at: "2026-10-17T23:47:36Z"
capture_tool: "hn-digest"
hn_id: 50136214
score: 2
comments: 2
posted_at: "2026-10-17T23:33:53Z"
tags:
  - hacker-news
---

# Does code quality still matter when AI writes the code?

- HN: [50136214](https://news.ycombinator.com/item?id=50136214)
- Source: [docspring.com](https://docspring.com/blog/posts/does-code-quality-still-matter-when-ai-writes-the-code/)
- Score: 2
- Comments: 2
- Posted: 2026-10-17T23:33:53Z

## Translation

Title: Does code quality still matter when AI writes the code?
Description: We replayed agent-assisted mistakes in DocSpring. Shared logic and explicit contracts looked more useful than shorter files or blanket cleanup.

Article text:
Start Free Trial
PDF Filling API
Contact
EMAIL:
[email protected]
Does code quality still matter when AI writes the code?
I can ask Claude Code or Codex to investigate parts of DocSpring that I would struggle to keep in my head at once. They often find problems faster than I can. That has changed how I think about a codebase we've spent more than nine years building.
I've been seeing versions of this argument on X: if an agent can understand and change complicated code, perhaps technical debt matters much less now. Why spend time refactoring something a model can already work through?
I wanted to test that against our own experience. As a small team, we've tried to write clear code and useful tests while keeping DocSpring moving. AI has expanded how much of the application we can investigate and improve. It also gives us a long list of possible improvements, each of which still needs review and testing.
So I asked an agent to investigate which kinds of code quality work might actually help us ship correct changes. We screened available coding-agent transcripts and pull-request discussions from the previous three months, inspected selected changes, and ran small experiments against the code. This was a retrospective case study, with selected examples rather than a controlled measurement of overall productivity.
The clearest failures involved behavior spread across several places, and checks that did not express what we intended. Giving that behavior one implementation, or making a missing requirement fail a check, could have caught the mistakes we reproduced. We found much less evidence that making every file shorter would help.
A correct fix can still miss another path
DocSpring has conditional rules for form fields: show this field when another field has a particular value, or require it when a condition holds. The behavior has to agree across server validation, browser forms, and the JSON Schema we generate for other consumers.
That gives a coding agent several places to keep in sync. It also gives a human developer several places to forget.
One change involved field names containing a slash. In JSON Pointer, a slash within a path segment is escaped as ~1 . For a synthetic field named source/field , that segment becomes source~1field . Code looking up the field needs to resolve the encoded path correctly.
A correction reached the visibility handling, but missed a separate branch for rules that only make a field required. That branch still used the encoded name directly. As a result, it could fail to attach the conditional validator.
We replayed the historical implementation with a synthetic input. Before the follow-up fix, the validator was absent. Afterward, it was installed and rejected an empty target when the condition required a value. This particular omission was caught and fixed before the feature merged.
The first fix made sense where it was applied. The reviewer still needed to know that a second path existed and needed the same correction.
It would be good to give path resolution one implementation and test the combinations that matter across its callers. For conditional fields, that includes the action, the field path, the input value, and the form or validator consuming the result. A test for showing a field does not establish that making it required also works.
We've already taken a related approach with formulas. Our formula evaluator is now shared TypeScript code, used in the browser and on the backend. We still have to test serialization, error handling, and the behavior of each runtime, but there is one implementation of the calculation itself.
I'm planning to apply that approach to field dependencies too.
Strict typing needs to express the contract
Another example came from our SDK test tooling. A subclass defined lint_command , while the base class expected linter_command . The base implementation returned nil because the hook was optional. Calling the expected method therefore did nothing.
It is an easy typo to imagine making, and the faulty file already had Sorbet checking enabled. Would requiring strict typing have caught it?
We tested a small reproduction:
# typed: strict
class SuiteExample
extend T::Sig
sig { returns(T.nilable(String)) }
def linter_command
nil
end
end
class GoSuiteExample < SuiteExample
sig { returns(T.nilable(String)) }
def lint_command
'golangci-lint run'
end
end
Sorbet accepted this. Both methods have valid signatures. Defining a new method named lint_command is perfectly legal.
Then we changed the subclass signature:
sig { override.returns(T.nilable(String)) }
Now Sorbet rejected the misspelled method because it did not override anything. Correcting the name to linter_command satisfies that contract. Sorbet's override checking documentation describes this behavior.
That changes what I'd ask an agent to do when adding Sorbet coverage. The first version was already strict and fully annotated. The missing information was that this method was meant to override the base class.
The same reasoning applies to the tooling around the method. If a quality check is required, a test should establish that the command runs and that its failure fails the job. An optional hook returning nothing is a poor place to hide that requirement.
Long files were harder to blame
I expected to find stronger evidence against very long files. Earlier coding tools could struggle with reading or rewriting them, and splitting a file is an obvious task to give an agent.
Our transcript search did not find explicit code-file size or truncation errors in the Claude file-read results it inspected. That is a narrow finding. It does not tell us whether relevant context was omitted, whether shell output was truncated, or whether the agent read the right section and misunderstood it.
It does mean this investigation failed to establish a useful cutoff such as "an agent cannot reliably work on a file over 1,000 lines."
Some large files contain substantial inline tests or long runs of declarations. Others contain several responsibilities that are difficult to change independently. I care much more about the second case, especially in code we change frequently.
One CI refactor illustrated the distinction well. The main generator file became much shorter, but the generator and its new modules contained more lines in total. Named responsibilities and executable checks made the structure more useful. Total line count would have described that change badly.
I still want smaller modules where they let an agent work on a piece of behavior without rediscovering the rest of the application. Clear inputs, outputs, and tests are a better reason to draw that boundary than a line limit.
Duplication and lint rules need the same scrutiny
We also ran jscpd to look for copied code. It found useful candidates, including shared transformations implemented separately for the browser and server. Those are worth investigating when both copies are supposed to obey the same rules.
But matching text is only one kind of duplication. Two implementations can use different languages and look nothing alike while encoding the same business rule. The field-dependency example is a more convincing reason to share behavior than a repository-wide duplication percentage.
There is also a cost to a shared helper that needs a growing collection of flags to preserve unrelated behavior. Before extracting one, I'd write down what its callers are actually meant to share.
We applied similar scrutiny to disabled RuboCop rules. Size and complexity warnings are useful places to start looking. They do not establish that every flagged method needs a refactor. A rule that exposes a silently swallowed tooling failure has a more direct connection to our correction cycles.
For a new rule, I would want a concrete explanation of which mistake it would prevent. Then we can fix the real findings in the area where we enable it and keep the check useful. Producing a large list of warnings is the easy part.
Cleaner code can save tokens without improving correctness
A controlled study by Priyansh Trivedi and Olivier Schmitt , both at SonarSource, compared cleaner and messier versions of six Java and Python repositories. Across 660 Claude Code trials using Sonnet 4.6, the cleaner versions used about 7% fewer input tokens and 9% fewer output tokens. The mean fraction of hidden tests passed did not improve.
The authors constructed the repository pairs and tasks, and the result concerns that model and setup. Still, it is useful counterevidence to the assumption that cleaner-looking code automatically makes an agent more likely to finish correctly.
Our own cases give us plausible improvements to try. They don't tell us how much time a general cleanup would save.
For DocSpring, my attention is often the constraint. I need to select work, clarify requirements, review changes, and perform QA. CI and deployment queues also affect when a change reaches users. An agent reading fewer tokens can be useful, but I can't translate that directly into hours saved or software shipped.
The proposed next step is a small, time-boxed pilot. Pick a frequently changed area with a known source of rework, improve its shared behavior or executable contracts, and track what happens to the next comparable changes.
The measures I care about are:
Active human time spent prompting, reviewing, and doing QA.
Corrections that require another round of work, and what caused them.
CI reruns and release waiting time, recorded separately from active work.
We'll need to record the task type and the model and harness too. Comparing an easy bug fix with a feature spanning several systems would tell us very little.
If you're considering the same question, start with a change that required several rounds of correction. What did the agent keep missing? Reproduce the failure, then ask whether shared code, an explicit type contract, or a test could have made it visible earlier. You can evaluate that improvement on the next changes to the same area.
AI has made me more willing to revisit old parts of DocSpring. I can investigate improvements that would previously have stayed on the backlog because I didn't have enough time or confidence to tackle them. For the next round, I want to spend that opportunity on changes that reduce how much we have to remember each time we ship.
Fill out PDFs in minutes, not weeks
Sign up for a free trial and try out DocSpring with your own PDFs.

## Original Extract

We replayed agent-assisted mistakes in DocSpring. Shared logic and explicit contracts looked more useful than shorter files or blanket cleanup.

Start Free Trial
PDF Filling API
Contact
EMAIL:
[email protected]
Does code quality still matter when AI writes the code?
I can ask Claude Code or Codex to investigate parts of DocSpring that I would struggle to keep in my head at once. They often find problems faster than I can. That has changed how I think about a codebase we've spent more than nine years building.
I've been seeing versions of this argument on X: if an agent can understand and change complicated code, perhaps technical debt matters much less now. Why spend time refactoring something a model can already work through?
I wanted to test that against our own experience. As a small team, we've tried to write clear code and useful tests while keeping DocSpring moving. AI has expanded how much of the application we can investigate and improve. It also gives us a long list of possible improvements, each of which still needs review and testing.
So I asked an agent to investigate which kinds of code quality work might actually help us ship correct changes. We screened available coding-agent transcripts and pull-request discussions from the previous three months, inspected selected changes, and ran small experiments against the code. This was a retrospective case study, with selected examples rather than a controlled measurement of overall productivity.
The clearest failures involved behavior spread across several places, and checks that did not express what we intended. Giving that behavior one implementation, or making a missing requirement fail a check, could have caught the mistakes we reproduced. We found much less evidence that making every file shorter would help.
A correct fix can still miss another path
DocSpring has conditional rules for form fields: show this field when another field has a particular value, or require it when a condition holds. The behavior has to agree across server validation, browser forms, and the JSON Schema we generate for other consumers.
That gives a coding agent several places to keep in sync. It also gives a human developer several places to forget.
One change involved field names containing a slash. In JSON Pointer, a slash within a path segment is escaped as ~1 . For a synthetic field named source/field , that segment becomes source~1field . Code looking up the field needs to resolve the encoded path correctly.
A correction reached the visibility handling, but missed a separate branch for rules that only make a field required. That branch still used the encoded name directly. As a result, it could fail to attach the conditional validator.
We replayed the historical implementation with a synthetic input. Before the follow-up fix, the validator was absent. Afterward, it was installed and rejected an empty target when the condition required a value. This particular omission was caught and fixed before the feature merged.
The first fix made sense where it was applied. The reviewer still needed to know that a second path existed and needed the same correction.
It would be good to give path resolution one implementation and test the combinations that matter across its callers. For conditional fields, that includes the action, the field path, the input value, and the form or validator consuming the result. A test for showing a field does not establish that making it required also works.
We've already taken a related approach with formulas. Our formula evaluator is now shared TypeScript code, used in the browser and on the backend. We still have to test serialization, error handling, and the behavior of each runtime, but there is one implementation of the calculation itself.
I'm planning to apply that approach to field dependencies too.
Strict typing needs to express the contract
Another example came from our SDK test tooling. A subclass defined lint_command , while the base class expected linter_command . The base implementation returned nil because the hook was optional. Calling the expected method therefore did nothing.
It is an easy typo to imagine making, and the faulty file already had Sorbet checking enabled. Would requiring strict typing have caught it?
We tested a small reproduction:
# typed: strict
class SuiteExample
extend T::Sig
sig { returns(T.nilable(String)) }
def linter_command
nil
end
end
class GoSuiteExample < SuiteExample
sig { returns(T.nilable(String)) }
def lint_command
'golangci-lint run'
end
end
Sorbet accepted this. Both methods have valid signatures. Defining a new method named lint_command is perfectly legal.
Then we changed the subclass signature:
sig { override.returns(T.nilable(String)) }
Now Sorbet rejected the misspelled method because it did not override anything. Correcting the name to linter_command satisfies that contract. Sorbet's override checking documentation describes this behavior.
That changes what I'd ask an agent to do when adding Sorbet coverage. The first version was already strict and fully annotated. The missing information was that this method was meant to override the base class.
The same reasoning applies to the tooling around the method. If a quality check is required, a test should establish that the command runs and that its failure fails the job. An optional hook returning nothing is a poor place to hide that requirement.
Long files were harder to blame
I expected to find stronger evidence against very long files. Earlier coding tools could struggle with reading or rewriting them, and splitting a file is an obvious task to give an agent.
Our transcript search did not find explicit code-file size or truncation errors in the Claude file-read results it inspected. That is a narrow finding. It does not tell us whether relevant context was omitted, whether shell output was truncated, or whether the agent read the right section and misunderstood it.
It does mean this investigation failed to establish a useful cutoff such as "an agent cannot reliably work on a file over 1,000 lines."
Some large files contain substantial inline tests or long runs of declarations. Others contain several responsibilities that are difficult to change independently. I care much more about the second case, especially in code we change frequently.
One CI refactor illustrated the distinction well. The main generator file became much shorter, but the generator and its new modules contained more lines in total. Named responsibilities and executable checks made the structure more useful. Total line count would have described that change badly.
I still want smaller modules where they let an agent work on a piece of behavior without rediscovering the rest of the application. Clear inputs, outputs, and tests are a better reason to draw that boundary than a line limit.
Duplication and lint rules need the same scrutiny
We also ran jscpd to look for copied code. It found useful candidates, including shared transformations implemented separately for the browser and server. Those are worth investigating when both copies are supposed to obey the same rules.
But matching text is only one kind of duplication. Two implementations can use different languages and look nothing alike while encoding the same business rule. The field-dependency example is a more convincing reason to share behavior than a repository-wide duplication percentage.
There is also a cost to a shared helper that needs a growing collection of flags to preserve unrelated behavior. Before extracting one, I'd write down what its callers are actually meant to share.
We applied similar scrutiny to disabled RuboCop rules. Size and complexity warnings are useful places to start looking. They do not establish that every flagged method needs a refactor. A rule that exposes a silently swallowed tooling failure has a more direct connection to our correction cycles.
For a new rule, I would want a concrete explanation of which mistake it would prevent. Then we can fix the real findings in the area where we enable it and keep the check useful. Producing a large list of warnings is the easy part.
Cleaner code can save tokens without improving correctness
A controlled study by Priyansh Trivedi and Olivier Schmitt , both at SonarSource, compared cleaner and messier versions of six Java and Python repositories. Across 660 Claude Code trials using Sonnet 4.6, the cleaner versions used about 7% fewer input tokens and 9% fewer output tokens. The mean fraction of hidden tests passed did not improve.
The authors constructed the repository pairs and tasks, and the result concerns that model and setup. Still, it is useful counterevidence to the assumption that cleaner-looking code automatically makes an agent more likely to finish correctly.
Our own cases give us plausible improvements to try. They don't tell us how much time a general cleanup would save.
For DocSpring, my attention is often the constraint. I need to select work, clarify requirements, review changes, and perform QA. CI and deployment queues also affect when a change reaches users. An agent reading fewer tokens can be useful, but I can't translate that directly into hours saved or software shipped.
The proposed next step is a small, time-boxed pilot. Pick a frequently changed area with a known source of rework, improve its shared behavior or executable contracts, and track what happens to the next comparable changes.
The measures I care about are:
Active human time spent prompting, reviewing, and doing QA.
Corrections that require another round of work, and what caused them.
CI reruns and release waiting time, recorded separately from active work.
We'll need to record the task type and the model and harness too. Comparing an easy bug fix with a feature spanning several systems would tell us very little.
If you're considering the same question, start with a change that required several rounds of correction. What did the agent keep missing? Reproduce the failure, then ask whether shared code, an explicit type contract, or a test could have made it visible earlier. You can evaluate that improvement on the next changes to the same area.
AI has made me more willing to revisit old parts of DocSpring. I can investigate improvements that would previously have stayed on the backlog because I didn't have enough time or confidence to tackle them. For the next round, I want to spend that opportunity on changes that reduce how much we have to remember each time we ship.
Fill out PDFs in minutes, not weeks
Sign up for a free trial and try out DocSpring with your own PDFs.
