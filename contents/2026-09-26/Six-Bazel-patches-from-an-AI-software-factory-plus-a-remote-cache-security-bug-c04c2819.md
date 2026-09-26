---
source: "https://www.incredibuild.com/blog/bazel-contributions-security-fix-autonomous-sdlc"
hn_url: "https://news.ycombinator.com/item?id=49858854"
title: "Six Bazel patches from an AI software factory, plus a remote-cache security bug"
article_title: "From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix - Incredibuild"
image: "https://www.incredibuild.com/wp-content/uploads/2026/09/bazel-hero.png"
author: "lol-lol-lol-2"
captured_at: "2026-09-26T18:22:59Z"
capture_tool: "hn-digest"
hn_id: 49858854
score: 1
comments: 0
posted_at: "2026-09-26T17:48:55Z"
tags:
  - hacker-news
---

# Six Bazel patches from an AI software factory, plus a remote-cache security bug

- HN: [49858854](https://news.ycombinator.com/item?id=49858854)
- Source: [www.incredibuild.com](https://www.incredibuild.com/blog/bazel-contributions-security-fix-autonomous-sdlc)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T17:48:55Z

## Translation

Title: Six Bazel patches from an AI software factory, plus a remote-cache security bug
Article title: From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix - Incredibuild
Description: Six upstream Bazel contributions and a remote-cache path-traversal fix: what Incredibuild's Software Factory found, and what Bazel users should check.

Article text:
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix - Incredibuild
Skip to content
Islo: The AI sandbox for always-on coding agents. Learn more
System Updates
AI SANDBOX
ISLO: The always-on execution layer for AI coding agents
·
PRODUCT · RELEASED MAY 2026
LEARN MORE
System Updates
GAME DEVELOPMENT
Shaders and the Unity Shader Compiler: From the GPU Pipeline to Faster Builds
·
BLOG · PUBLISHED JULY 2026
READ MORE
System Updates
CI ACCELERATION
Get 8× faster CI runners for free
·
PRODUCT · JOIN THE EARLY ACCESS
JOIN NOW
Product
Close Product
Open Product
Accelerate Everything
Solutions
Close Solutions
Open Solutions
Industries
Resources
Close Resources
Open Resources
Resources​
Company
Close Company
Open Company
Incredibuild empowers teams to build faster, create better products, and have greater control over their dev processes.
Get Started
English
日本語 ( Japanese )
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
Using an early internal version of Incredibuild’s Software Factory, we landed six upstream Bazel contributions, uncovered a remote-cache path-containment security flaw, and came away with a concrete set of requirements for automated software engineering. A companion to our Bazel series: Part 1 , Part 2 , and Part 3 .
Generating code is getting fast and cheap. Building, testing, validating, and deciding what reaches main is where agents, and people, now stand waiting.
For more than two decades, Incredibuild has focused on one problem: shortening the build-test loop without forcing engineering teams to reinvent how they build software. That began with distributed compilation and grew into caching, CI, testing, telemetry, and build integrity.
AI is moving the bottleneck. When an agent can draft a patch in seconds, the long pole is everything after it: the build, the test run, the evidence, the review. That’s the problem our Software Factory is built to solve. Before pointing it at anyone else’s code, we pointed it at our own products and at the ecosystem around us, and Bazel is a big part of that ecosystem.
Six upstream Bazel contributions from our Software Factory were imported to Bazel master through Google’s Copybara workflow.
Along the way, we found a path-containment flaw in --experimental_repo_contents_cache , reproduced it end to end on a shipping Bazel 9.0.2 binary, and reported it privately.
Google fixed it upstream in c37a6a1bf7f4 . The fix is in Bazel 9.3.0rc2 ; Bazel 9.2.0 still ships the pre-fix logic.
If you use that flag: upgrade to a build that contains the fix, or turn the flag off until you can.
Why Bazel was the right proving ground
Bazel models software as a graph of declared work, makes caching and remote execution first-class concepts, and treats reproducibility as part of correctness. Those are the same problems Incredibuild has spent two decades solving from a different layer of the stack. (If you want the full background, start with Bazel, Inside and Out .)
It’s also a demanding target. Bazel’s Remote Build Execution (RBE) model is powerful, but adopting it is a real engineering project: hermetic actions, rule coverage, toolchains, platforms, and REAPI infrastructure all have to line up. And real repositories rarely live entirely inside a clean action graph. They contain nested Make or Ninja builds, custom compilers, test harnesses, scripts, packaging, code generation, and proprietary tools.
Our approach is complementary. Bazel orchestrates the build; Incredibuild provides caching and distributed execution across the existing workload , including the parts that are hard to move to RBE. Teams get faster builds now, while the RBE migration continues on its own schedule.
So Bazel mattered to us twice: as a workload we accelerate, and as a build system we have to understand all the way down to its cache boundaries.
What the Software Factory actually does
The Software Factory isn’t an “agent writes code” demo. It runs the whole engineering loop as a factory line:
goal → plan → isolated task → implement → build → test → inspect → retry → reviewable PR
Orchestration manages the project as a DAG of tasks.
Islo sandboxes isolate every task.
Incredibuild accelerates the expensive build and test loops.
Agents explore aggressively, but tests, evidence, and human review decide what gets promoted.
One rule doesn’t bend: automated search and model output can propose patches, run experiments, or return cached results. Test verification and human review keep authority over what merges to main.
Six contributions landed upstream
Six changes from this effort were imported to Bazel master through Google’s Copybara workflow. None of them is flashy. All of them are about the same thing: making automation trustworthy.
Links to each change: PR #30365 · PR #30358 · PR #30364 · PR #30363 · PR #30360 · PR #30359 .
The pattern is the point. Deterministic outputs, complete evidence, clear authority boundaries, and safe failure modes aren’t nice-to-haves for automated software engineering. They’re the difference between an agent you can let run and one you have to babysit. Partial evidence must never be treated as complete evidence.
Then the cache boundary got interesting
As we went deeper into Bazel’s caching architecture, we examined --experimental_repo_contents_cache .
The feature restores the output of reproducible repository rules from a remote Action Cache (AC) and Content Addressable Storage (CAS). That’s good for performance, but it moves the trust boundary: cache-supplied metadata now participates in filesystem materialization.
Before the fix, Bazel trusted cache-supplied tree node names without sufficiently enforcing that the resulting path stayed inside the intended external repository directory. A node name containing parent-directory traversal or an absolute path could land outside it.
We didn’t stop at reading source. The Software Factory generated a reproducible, end-to-end test case against a real Bazel 9.0.2 binary. Under the required conditions, a poisoned remote-cache entry caused Bazel to overwrite a pre-existing file outside the repository root, while the build itself still completed successfully.
All three conditions have to hold:
This is not a drive-by internet vulnerability. But that’s exactly why it matters: once cache metadata can steer filesystem writes, anyone who can write to your remote cache is part of your trusted computing base.
We reported the finding privately to the Bazel security team, with the reproduction attached.
Google’s Bazel team fixed the issue in commit c37a6a1bf7f4 , “Improve path validation in remote repository contents cache.” The fix:
rejects remote tree node names that contain separators, parent-directory traversal, or absolute paths;
checks that injected file, directory, and symlink paths stay inside the repository directory;
resolves symlink targets, absolute and relative, and verifies they stay inside it too;
adds a regression test asserting that invalid paths are rejected.
That’s the right place for the check: the exact boundary where untrusted cache metadata becomes a filesystem path.
What to do if you use --experimental_repo_contents_cache
1. Find out whether the flag is on. Search every place Bazel options come from:
grep -rn "experimental_repo_contents_cache" \
.bazelrc ~/.bazelrc /etc/bazel.bazelrc ci/ tools/ 2>/dev/null
2. Check the exact binary you run , not the one you think you run. With Bazelisk, also check .bazelversion and USE_BAZEL_VERSION .
bazel --version
Bazel 9.2.0 still contains the pre-fix path logic. Bazel 9.3.0rc2 contains the path-containment validation.
3. Upgrade or disable. Move to a Bazel build that contains c37a6a1bf7f4 . If your organization can’t deploy a release candidate, disable --experimental_repo_contents_cache until a stable release with the fix is available.
4. Treat cache write access as a security boundary , not a performance setting. Know exactly who and what can write to your shared remote Action Cache and CAS.
What this taught us about automated SDLC
Build acceleration runs on structured graphs, deterministic inputs, reusable outputs, process isolation, scheduling, and defined trust boundaries. Those same primitives are the foundation of automated software engineering:
Caching , because agents repeat work at enormous scale.
Distribution , because once coding gets cheap, build, test, analysis, and validation become the long pole.
Isolation , because many agent loops run concurrently.
Observability , because humans need a receipt, not a promise.
Determinism , because exploration and promotion can’t have the same authority.
Bazel turns source code into an explicit execution graph. The Software Factory extends the same idea upward: it turns the software development lifecycle itself into a managed graph.
We ran the system against an active open-source project to see whether its architecture holds up against real maintainer feedback, build failures, iterative changes, and security disclosure:
Six upstream contributions landed.
A security boundary was found and reproduced on a shipping binary.
The issue was disclosed privately and responsibly.
After more than two decades of making builds faster, the next step is making the entire engineering loop faster, from intent to tested, reviewable, evidence-backed software, without giving up the controls that make software trustworthy.
Bazel PR #30365: Make bazel-distfile.tar reproducible
Bazel PR #30358: Skip SSL tracking-issue mutation on PR validation
Bazel PR #30364: Ignore COMMENTED reviews in community-review status
Bazel PR #30363: Locale-independent certificate expiration parsing
Bazel PR #30360: Fix BAZEL_DOC_TRIGGER_TOKEN documented scope
Bazel PR #30359: Paginate PR-review retrieval and fail closed
Commit c37a6a1bf7f4: Improve path validation in remote repository contents cache
Bazel, Inside and Out (Part 1)
Bazel Caching, Remote Execution, and the Build Supply Chain (Part 2)
Bazel in the Real World (Part 3)
Incredibuild empowers your teams to be productive and focus on innovating.
Shorten Your Builds
Incredibuild empowers your teams to be productive and focus on innovating.
X-twitter
Linkedin
Slack
Product
Build Acceleration
Incredibuild empowers teams to build faster, create better products, and have greater control over their dev processes.

## Original Extract

Six upstream Bazel contributions and a remote-cache path-traversal fix: what Incredibuild's Software Factory found, and what Bazel users should check.

From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix - Incredibuild
Skip to content
Islo: The AI sandbox for always-on coding agents. Learn more
System Updates
AI SANDBOX
ISLO: The always-on execution layer for AI coding agents
·
PRODUCT · RELEASED MAY 2026
LEARN MORE
System Updates
GAME DEVELOPMENT
Shaders and the Unity Shader Compiler: From the GPU Pipeline to Faster Builds
·
BLOG · PUBLISHED JULY 2026
READ MORE
System Updates
CI ACCELERATION
Get 8× faster CI runners for free
·
PRODUCT · JOIN THE EARLY ACCESS
JOIN NOW
Product
Close Product
Open Product
Accelerate Everything
Solutions
Close Solutions
Open Solutions
Industries
Resources
Close Resources
Open Resources
Resources​
Company
Close Company
Open Company
Incredibuild empowers teams to build faster, create better products, and have greater control over their dev processes.
Get Started
English
日本語 ( Japanese )
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
From Build Acceleration to Autonomous SDLC: Six Bazel Contributions and a Security Fix
Using an early internal version of Incredibuild’s Software Factory, we landed six upstream Bazel contributions, uncovered a remote-cache path-containment security flaw, and came away with a concrete set of requirements for automated software engineering. A companion to our Bazel series: Part 1 , Part 2 , and Part 3 .
Generating code is getting fast and cheap. Building, testing, validating, and deciding what reaches main is where agents, and people, now stand waiting.
For more than two decades, Incredibuild has focused on one problem: shortening the build-test loop without forcing engineering teams to reinvent how they build software. That began with distributed compilation and grew into caching, CI, testing, telemetry, and build integrity.
AI is moving the bottleneck. When an agent can draft a patch in seconds, the long pole is everything after it: the build, the test run, the evidence, the review. That’s the problem our Software Factory is built to solve. Before pointing it at anyone else’s code, we pointed it at our own products and at the ecosystem around us, and Bazel is a big part of that ecosystem.
Six upstream Bazel contributions from our Software Factory were imported to Bazel master through Google’s Copybara workflow.
Along the way, we found a path-containment flaw in --experimental_repo_contents_cache , reproduced it end to end on a shipping Bazel 9.0.2 binary, and reported it privately.
Google fixed it upstream in c37a6a1bf7f4 . The fix is in Bazel 9.3.0rc2 ; Bazel 9.2.0 still ships the pre-fix logic.
If you use that flag: upgrade to a build that contains the fix, or turn the flag off until you can.
Why Bazel was the right proving ground
Bazel models software as a graph of declared work, makes caching and remote execution first-class concepts, and treats reproducibility as part of correctness. Those are the same problems Incredibuild has spent two decades solving from a different layer of the stack. (If you want the full background, start with Bazel, Inside and Out .)
It’s also a demanding target. Bazel’s Remote Build Execution (RBE) model is powerful, but adopting it is a real engineering project: hermetic actions, rule coverage, toolchains, platforms, and REAPI infrastructure all have to line up. And real repositories rarely live entirely inside a clean action graph. They contain nested Make or Ninja builds, custom compilers, test harnesses, scripts, packaging, code generation, and proprietary tools.
Our approach is complementary. Bazel orchestrates the build; Incredibuild provides caching and distributed execution across the existing workload , including the parts that are hard to move to RBE. Teams get faster builds now, while the RBE migration continues on its own schedule.
So Bazel mattered to us twice: as a workload we accelerate, and as a build system we have to understand all the way down to its cache boundaries.
What the Software Factory actually does
The Software Factory isn’t an “agent writes code” demo. It runs the whole engineering loop as a factory line:
goal → plan → isolated task → implement → build → test → inspect → retry → reviewable PR
Orchestration manages the project as a DAG of tasks.
Islo sandboxes isolate every task.
Incredibuild accelerates the expensive build and test loops.
Agents explore aggressively, but tests, evidence, and human review decide what gets promoted.
One rule doesn’t bend: automated search and model output can propose patches, run experiments, or return cached results. Test verification and human review keep authority over what merges to main.
Six contributions landed upstream
Six changes from this effort were imported to Bazel master through Google’s Copybara workflow. None of them is flashy. All of them are about the same thing: making automation trustworthy.
Links to each change: PR #30365 · PR #30358 · PR #30364 · PR #30363 · PR #30360 · PR #30359 .
The pattern is the point. Deterministic outputs, complete evidence, clear authority boundaries, and safe failure modes aren’t nice-to-haves for automated software engineering. They’re the difference between an agent you can let run and one you have to babysit. Partial evidence must never be treated as complete evidence.
Then the cache boundary got interesting
As we went deeper into Bazel’s caching architecture, we examined --experimental_repo_contents_cache .
The feature restores the output of reproducible repository rules from a remote Action Cache (AC) and Content Addressable Storage (CAS). That’s good for performance, but it moves the trust boundary: cache-supplied metadata now participates in filesystem materialization.
Before the fix, Bazel trusted cache-supplied tree node names without sufficiently enforcing that the resulting path stayed inside the intended external repository directory. A node name containing parent-directory traversal or an absolute path could land outside it.
We didn’t stop at reading source. The Software Factory generated a reproducible, end-to-end test case against a real Bazel 9.0.2 binary. Under the required conditions, a poisoned remote-cache entry caused Bazel to overwrite a pre-existing file outside the repository root, while the build itself still completed successfully.
All three conditions have to hold:
This is not a drive-by internet vulnerability. But that’s exactly why it matters: once cache metadata can steer filesystem writes, anyone who can write to your remote cache is part of your trusted computing base.
We reported the finding privately to the Bazel security team, with the reproduction attached.
Google’s Bazel team fixed the issue in commit c37a6a1bf7f4 , “Improve path validation in remote repository contents cache.” The fix:
rejects remote tree node names that contain separators, parent-directory traversal, or absolute paths;
checks that injected file, directory, and symlink paths stay inside the repository directory;
resolves symlink targets, absolute and relative, and verifies they stay inside it too;
adds a regression test asserting that invalid paths are rejected.
That’s the right place for the check: the exact boundary where untrusted cache metadata becomes a filesystem path.
What to do if you use --experimental_repo_contents_cache
1. Find out whether the flag is on. Search every place Bazel options come from:
grep -rn "experimental_repo_contents_cache" \
.bazelrc ~/.bazelrc /etc/bazel.bazelrc ci/ tools/ 2>/dev/null
2. Check the exact binary you run , not the one you think you run. With Bazelisk, also check .bazelversion and USE_BAZEL_VERSION .
bazel --version
Bazel 9.2.0 still contains the pre-fix path logic. Bazel 9.3.0rc2 contains the path-containment validation.
3. Upgrade or disable. Move to a Bazel build that contains c37a6a1bf7f4 . If your organization can’t deploy a release candidate, disable --experimental_repo_contents_cache until a stable release with the fix is available.
4. Treat cache write access as a security boundary , not a performance setting. Know exactly who and what can write to your shared remote Action Cache and CAS.
What this taught us about automated SDLC
Build acceleration runs on structured graphs, deterministic inputs, reusable outputs, process isolation, scheduling, and defined trust boundaries. Those same primitives are the foundation of automated software engineering:
Caching , because agents repeat work at enormous scale.
Distribution , because once coding gets cheap, build, test, analysis, and validation become the long pole.
Isolation , because many agent loops run concurrently.
Observability , because humans need a receipt, not a promise.
Determinism , because exploration and promotion can’t have the same authority.
Bazel turns source code into an explicit execution graph. The Software Factory extends the same idea upward: it turns the software development lifecycle itself into a managed graph.
We ran the system against an active open-source project to see whether its architecture holds up against real maintainer feedback, build failures, iterative changes, and security disclosure:
Six upstream contributions landed.
A security boundary was found and reproduced on a shipping binary.
The issue was disclosed privately and responsibly.
After more than two decades of making builds faster, the next step is making the entire engineering loop faster, from intent to tested, reviewable, evidence-backed software, without giving up the controls that make software trustworthy.
Bazel PR #30365: Make bazel-distfile.tar reproducible
Bazel PR #30358: Skip SSL tracking-issue mutation on PR validation
Bazel PR #30364: Ignore COMMENTED reviews in community-review status
Bazel PR #30363: Locale-independent certificate expiration parsing
Bazel PR #30360: Fix BAZEL_DOC_TRIGGER_TOKEN documented scope
Bazel PR #30359: Paginate PR-review retrieval and fail closed
Commit c37a6a1bf7f4: Improve path validation in remote repository contents cache
Bazel, Inside and Out (Part 1)
Bazel Caching, Remote Execution, and the Build Supply Chain (Part 2)
Bazel in the Real World (Part 3)
Incredibuild empowers your teams to be productive and focus on innovating.
Shorten Your Builds
Incredibuild empowers your teams to be productive and focus on innovating.
X-twitter
Linkedin
Slack
Product
Build Acceleration
Incredibuild empowers teams to build faster, create better products, and have greater control over their dev processes.
