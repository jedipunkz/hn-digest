---
source: "https://antigravity.google/blog/modernizing-40-year-old-scientific-code-with-gemini-4-argon/"
hn_url: "https://news.ycombinator.com/item?id=50397967"
title: "Modernising 40-year-old scientific code with Gemini 4 Argon"
article_title: "Modernizing 40-year-old scientific code with Gemini 4 Argon | Google Antigravity Blog"
image: "https://antigravity.google/assets/image/blog/translating-fortran-to-rust-wide.png"
author: "andrewstetsenko"
captured_at: "2026-11-06T21:19:00Z"
capture_tool: "hn-digest"
hn_id: 50397967
score: 1
comments: 0
posted_at: "2026-11-06T20:25:54Z"
tags:
  - hacker-news
---

# Modernising 40-year-old scientific code with Gemini 4 Argon

- HN: [50397967](https://news.ycombinator.com/item?id=50397967)
- Source: [antigravity.google](https://antigravity.google/blog/modernizing-40-year-old-scientific-code-with-gemini-4-argon/)
- Score: 1
- Comments: 0
- Posted: 2026-11-06T20:25:54Z

## Translation

Title: Modernising 40-year-old scientific code with Gemini 4 Argon
Article title: Modernizing 40-year-old scientific code with Gemini 4 Argon | Google Antigravity Blog
Description: Using Gemini 4 Argon in Google Antigravity, we migrated the core of a 174,000-line Fortran coastal ocean model into 100% memory-safe Rust in 59 minutes…

Article text:
Modernizing 40-year-old scientific code with Gemini 4 Argon | Google Antigravity Blog Copy Logo as SVG Copied! Press Guidelines Products Use Cases Pricing Enterprise Resources 🚀 Launch Remote Control Download 🚀 Launch Remote Control EN English (United States) Bahasa Indonesia (Indonesia) Deutsch (Deutschland) English (Canada) English (United Kingdom) Español (España) Español (México) Français (France) Italiano (Italia) Polski (Polska) Português (Brasil) हिन्दी (भारत) ไทย (ประเทศไทย) 日本語 (日本) 简体中文 (中国) 繁體中文 (台灣) 한국어 (대한민국) Explore our next generation products
Antigravity App Antigravity Extensions Antigravity CLI Antigravity IDE Antigravity SDK Built for developers and teams in the agent-first era
Everything you need to stay up-to-date and get help
Changelog Newsletter Support Press Releases Products Explore our next generation products
Antigravity App Antigravity Extensions Antigravity CLI Antigravity IDE Antigravity SDK Use Cases Built for developers and teams in the agent-first era
Everything you need to stay up-to-date and get help
Changelog Newsletter Support Press Releases Download New Introducing the Plugin Marketplace in Google Antigravity Read post 0:00 / 0:00 Modernizing 40-year-old scientific code with Gemini 4 Argon - Blog blog Blog Blog Modernizing 40-year-old scientific code with Gemini 4 Argon - Blog All posts / Research Modernizing 40-year-old scientific code with Gemini 4 Argon
Share Copy link Share on X Share on Bluesky Share on LinkedIn Using Gemini 4 Argon in Google Antigravity, we migrated the core of a 174,000-line Fortran coastal ocean model into 100% memory-safe Rust in 59 minutes. Along the way, the agent caught six decades-old bugs, cut memory usage in half, and sped up simulations by 1.88x. Today, we’re sharing how Antigravity ran the migration end to end and releasing the workflow as an open-source skill.
We built Google Antigravity for long-horizon engineering work where an agent needs to do much more than suggest the next few lines of code. Earlier this week, as part of the Gemini 4 Argon launch in Google Antigravity , we shared a short video from one of our favorite internal stress tests: pointing Antigravity at a 40-year-old Fortran codebase and letting it autonomously plan, rewrite, test, and optimize the simulation in Rust.
Gemini 4 Argon in Google Antigravity modernizes a 40-year-old Fortran ocean model into memory-safe Rust.
After we posted that clip on X and YouTube, we got a lot of questions from researchers and systems engineers asking how we kept the agent from breaking 40 years of validated math. Here’s a look under the hood at why this required an agent-first platform like Antigravity, what the agent actually did during the 59-minute run, and how you can run the same workflow on your own projects.
Why so much science still runs on 1980s Fortran
A surprising amount of the software that models the physical world still runs on Fortran written in the 1970s and 1980s. That includes the models we rely on for hurricane storm surges, weather forecasting, and computational physics.
Scientists still use these codebases for a simple reason: they trust them. Generations of researchers have spent decades calibrating the equations inside those files against real-world buoys, satellites, and lab data. Even when a single file grows to 14,000 lines of fixed-column Fortran with hundreds of global variables, the math works. Nobody wants to touch it and risk breaking decades of validated science.
The catch is that modern developer tools can barely read these files. IDEs, test runners, memory sanitizers, and CI pipelines choke on legacy Fortran syntax. When new grad students join a lab, they often spend their first six months fighting segmentation faults and brittle build scripts before they can do any real research.
Porting these systems by hand is brutal work. In 1980s Fortran, the physics equations are tangled up with hardware workarounds from another era: shared COMMON blocks, 1-based indexing tricks, and implicit variable types. A manual rewrite takes years. Worse, human translators inevitably introduce tiny off-by-one errors or change the order of floating-point operations. Two years later, the new code finally compiles, the output numbers don’t match the original Fortran, and the team spends another year hunting down where the math drifted.
Why this needed Antigravity, not just a chat prompt
Pasting snippets of a legacy simulation into a chat window doesn’t work. In old Fortran, a variable modified on line 4,000 of one file can silently alter a global array used by twenty other files. Chop the repo into small chunks, and a model quickly loses the big picture.
Even with a frontier model, a full codebase migration is a systems engineering job that can’t happen in a text box alone. An agent needs an environment where it can explore a repository, compile legacy binaries in the terminal, author a multi-file workspace, run test suites, read its own stack traces, compare binary outputs side by side, and run compiler profilers without a human copying and pasting between tools.
That’s why we paired Gemini 4 Argon with Google Antigravity. While Argon brings an expansive context window and deep mathematical reasoning, Antigravity complements it with the autonomous execution loop, terminal and compiler integration, structured artifacts, and reusable skills needed to carry a complex migration from initial audit to verified binary.
To put both to the test, we picked FVCOM (the Finite-Volume Community Ocean Model), a widely used coastal ocean simulation with 174,774 lines of Fortran 90 and C spread across 264 files. We gave Antigravity a single prompt with four hard rules:
Enforce #![forbid(unsafe_code)] across the entire Rust workspace, with zero unsafe blocks.
Eliminate heap allocations inside hot simulation loops and use cache-friendly Struct-of-Arrays (SoA) layouts.
Keep every Rust file under 400 lines of code, with runnable doc-tests for every module.
Pass at least 95% branch test coverage and match the original Fortran with zero difference in the numerical output across 100 randomized simulation seeds.
From that single prompt, Antigravity orchestrated a 59-minute autonomous run across three stages:
Stage 0: Baseline capture in the terminal (11 min). Working directly in the Antigravity terminal, the agent cloned FVCOM, inspected the legacy build flags, compiled the original Fortran binary with gfortran -O3 , and built a benchmark harness to record golden outputs, runtime, and memory usage across 100 random seeds.
Stage 1: Full-repo audit and blueprint artifact (3 min). Antigravity scanned all 264 files across the repository at once, tracing cross-file dependencies and following how every global variable moved through the 2D and 3D solver loops. Before writing any Rust, it generated a six-part architectural blueprint artifact in the workspace, logging latent bugs in the legacy code and organizing global arrays into clean Rust modules.
Stage 2: Rust port, closed-loop self-healing, and PGO (45 min). Executing against its blueprint artifact, Antigravity scaffolded the Cargo workspace, rewrote the core solver into 13 Rust modules (each under 240 lines), and drove an automated verification loop in the terminal: running cargo test and cargo clippy , verifying 100-seed parity against the Fortran binary, and profiling a live run with LLVM Profile-Guided Optimization ( llvm-profdata ).
What happened inside the 59-minute run
Catching six silent bugs from the 1990s
During the 3-minute blueprint stage, Argon read the codebase with the eye of a systems engineer and a physicist. At one point, it flagged a comment left in the Fortran momentum solver decades ago:
H1(0) = H1(1) ! STRANGE WAY TO PREVENT DIVISION BY ZERO
In the original code, missing neighbor triangles along the coastline were indexed at cell 0 . To keep boundary loops from dividing by zero, the original authors copied cell 1 ’s water depth into index 0 , quietly tying boundary math to whichever triangle happened to be numbered 1 . Argon figured out why the workaround was there and designed clean boundary types in the Antigravity blueprint artifact.
Fortran also lets you get away with memory hazards that Rust refuses to compile. As Argon mapped FVCOM’s arrays into strict Rust types, six real bugs in the original Fortran and C code surfaced immediately:
Buffer and integer overflows: A mesh builder blindly wrote past a 50-column limit on complex coastal grid points, and a triangle edge check multiplied three node IDs in 32-bit precision, overflowing on grids with more than 1,300 nodes.
Out-of-bounds and uninitialized memory reads: Vertical turbulence and momentum routines read out-of-bounds index 0 elements on shallow grids and pulled uninitialized values from dry land cells.
Unguarded math: A seawater density routine took the square root of salinity without checking if numerical undershoot had nudged it slightly negative during advection.
Self-healing in the Antigravity terminal
Once the Rust modules compiled, Antigravity generated 42 unit, property, and documentation tests, including edge-case geometries designed to stress the physics solver, and executed them in the workspace terminal.
On the first test run, 36 out of 37 unit tests passed. One failed: when simulating a very shallow ocean basin with only a single vertical water layer, the vertical flux solver tried to interpolate with the water layer below it and hit a Rust bounds check.
Because Antigravity runs in a closed agentic loop over the terminal and editor, it didn’t stop to ask a human what to do. The agent captured the panic trace from the test runner, opened the vertical solver module, added a guard for single-layer grids, and re-ran cargo test . Seconds later, all 42 tests passed with 96.8% branch coverage. Antigravity then launched its A/B parity harness against the Fortran baseline across 100 randomized simulation seeds, matching all 100 runs to 12 decimal places (3.91 × 10 -13 ) on the first try.
Where the 1.88x speedup came from
As Antigravity completed each optimization pass, it logged runtime and memory telemetry and generated a standalone interactive HTML scorecard artifact right in the workspace so we could see what actually moved the needle:
The biggest jump came from a bottleneck Argon spotted in the 2D Runge-Kutta loop: the original Fortran code was allocating and freeing eight temporary heap arrays 24 times per time step. Moving those into pre-allocated Struct-of-Arrays workspaces, and having Antigravity run an automated LLVM PGO pass on a live simulation trace, cut runtime nearly in half and dropped memory usage by 54%.
Using Antigravity’s custom skills framework, we packaged this entire multi-stage workflow into an open-source skill in the Google Skills repository on GitHub . It bundles the architectural prompt guardrails, baseline discovery, bug audit, #![forbid(unsafe_code)] Rust scaffolding, self-healing test loop, multi-seed parity harness, PGO pass, and the interactive HTML telemetry artifact generator.
Once installed in Google Antigravity, you can run the full pipeline on any repository with one command:
/migrate-to-rust <github-repo-url>
What comes next
Across science and engineering, millions of lines of hard-won mathematical models are at risk of becoming unmaintainable as the people who wrote them in the 1970s and 1980s retire. Until now, labs had to choose between nursing fragile 40-year-old code or pausing research for years to rewrite it by hand.
With Gemini 4 Argon in Google Antigravity, teams can bring that work forward in an afternoon, keeping the exact same math while running nearly twice as fast in safe, modern Rust.
Share Copy link Share on X Share on Bluesky Share on LinkedIn Stay in sync with Google Antigravity
Get the latest product updates, releases, and feature drops.
All Updates Antigravity App Antigravity CLI IDE Extensions Antigravity SDK Email address Subscribe I accept Google's Terms and Privacy Polic

[truncated]

## Original Extract

Using Gemini 4 Argon in Google Antigravity, we migrated the core of a 174,000-line Fortran coastal ocean model into 100% memory-safe Rust in 59 minutes…

Modernizing 40-year-old scientific code with Gemini 4 Argon | Google Antigravity Blog Copy Logo as SVG Copied! Press Guidelines Products Use Cases Pricing Enterprise Resources 🚀 Launch Remote Control Download 🚀 Launch Remote Control EN English (United States) Bahasa Indonesia (Indonesia) Deutsch (Deutschland) English (Canada) English (United Kingdom) Español (España) Español (México) Français (France) Italiano (Italia) Polski (Polska) Português (Brasil) हिन्दी (भारत) ไทย (ประเทศไทย) 日本語 (日本) 简体中文 (中国) 繁體中文 (台灣) 한국어 (대한민국) Explore our next generation products
Antigravity App Antigravity Extensions Antigravity CLI Antigravity IDE Antigravity SDK Built for developers and teams in the agent-first era
Everything you need to stay up-to-date and get help
Changelog Newsletter Support Press Releases Products Explore our next generation products
Antigravity App Antigravity Extensions Antigravity CLI Antigravity IDE Antigravity SDK Use Cases Built for developers and teams in the agent-first era
Everything you need to stay up-to-date and get help
Changelog Newsletter Support Press Releases Download New Introducing the Plugin Marketplace in Google Antigravity Read post 0:00 / 0:00 Modernizing 40-year-old scientific code with Gemini 4 Argon - Blog blog Blog Blog Modernizing 40-year-old scientific code with Gemini 4 Argon - Blog All posts / Research Modernizing 40-year-old scientific code with Gemini 4 Argon
Share Copy link Share on X Share on Bluesky Share on LinkedIn Using Gemini 4 Argon in Google Antigravity, we migrated the core of a 174,000-line Fortran coastal ocean model into 100% memory-safe Rust in 59 minutes. Along the way, the agent caught six decades-old bugs, cut memory usage in half, and sped up simulations by 1.88x. Today, we’re sharing how Antigravity ran the migration end to end and releasing the workflow as an open-source skill.
We built Google Antigravity for long-horizon engineering work where an agent needs to do much more than suggest the next few lines of code. Earlier this week, as part of the Gemini 4 Argon launch in Google Antigravity , we shared a short video from one of our favorite internal stress tests: pointing Antigravity at a 40-year-old Fortran codebase and letting it autonomously plan, rewrite, test, and optimize the simulation in Rust.
Gemini 4 Argon in Google Antigravity modernizes a 40-year-old Fortran ocean model into memory-safe Rust.
After we posted that clip on X and YouTube, we got a lot of questions from researchers and systems engineers asking how we kept the agent from breaking 40 years of validated math. Here’s a look under the hood at why this required an agent-first platform like Antigravity, what the agent actually did during the 59-minute run, and how you can run the same workflow on your own projects.
Why so much science still runs on 1980s Fortran
A surprising amount of the software that models the physical world still runs on Fortran written in the 1970s and 1980s. That includes the models we rely on for hurricane storm surges, weather forecasting, and computational physics.
Scientists still use these codebases for a simple reason: they trust them. Generations of researchers have spent decades calibrating the equations inside those files against real-world buoys, satellites, and lab data. Even when a single file grows to 14,000 lines of fixed-column Fortran with hundreds of global variables, the math works. Nobody wants to touch it and risk breaking decades of validated science.
The catch is that modern developer tools can barely read these files. IDEs, test runners, memory sanitizers, and CI pipelines choke on legacy Fortran syntax. When new grad students join a lab, they often spend their first six months fighting segmentation faults and brittle build scripts before they can do any real research.
Porting these systems by hand is brutal work. In 1980s Fortran, the physics equations are tangled up with hardware workarounds from another era: shared COMMON blocks, 1-based indexing tricks, and implicit variable types. A manual rewrite takes years. Worse, human translators inevitably introduce tiny off-by-one errors or change the order of floating-point operations. Two years later, the new code finally compiles, the output numbers don’t match the original Fortran, and the team spends another year hunting down where the math drifted.
Why this needed Antigravity, not just a chat prompt
Pasting snippets of a legacy simulation into a chat window doesn’t work. In old Fortran, a variable modified on line 4,000 of one file can silently alter a global array used by twenty other files. Chop the repo into small chunks, and a model quickly loses the big picture.
Even with a frontier model, a full codebase migration is a systems engineering job that can’t happen in a text box alone. An agent needs an environment where it can explore a repository, compile legacy binaries in the terminal, author a multi-file workspace, run test suites, read its own stack traces, compare binary outputs side by side, and run compiler profilers without a human copying and pasting between tools.
That’s why we paired Gemini 4 Argon with Google Antigravity. While Argon brings an expansive context window and deep mathematical reasoning, Antigravity complements it with the autonomous execution loop, terminal and compiler integration, structured artifacts, and reusable skills needed to carry a complex migration from initial audit to verified binary.
To put both to the test, we picked FVCOM (the Finite-Volume Community Ocean Model), a widely used coastal ocean simulation with 174,774 lines of Fortran 90 and C spread across 264 files. We gave Antigravity a single prompt with four hard rules:
Enforce #![forbid(unsafe_code)] across the entire Rust workspace, with zero unsafe blocks.
Eliminate heap allocations inside hot simulation loops and use cache-friendly Struct-of-Arrays (SoA) layouts.
Keep every Rust file under 400 lines of code, with runnable doc-tests for every module.
Pass at least 95% branch test coverage and match the original Fortran with zero difference in the numerical output across 100 randomized simulation seeds.
From that single prompt, Antigravity orchestrated a 59-minute autonomous run across three stages:
Stage 0: Baseline capture in the terminal (11 min). Working directly in the Antigravity terminal, the agent cloned FVCOM, inspected the legacy build flags, compiled the original Fortran binary with gfortran -O3 , and built a benchmark harness to record golden outputs, runtime, and memory usage across 100 random seeds.
Stage 1: Full-repo audit and blueprint artifact (3 min). Antigravity scanned all 264 files across the repository at once, tracing cross-file dependencies and following how every global variable moved through the 2D and 3D solver loops. Before writing any Rust, it generated a six-part architectural blueprint artifact in the workspace, logging latent bugs in the legacy code and organizing global arrays into clean Rust modules.
Stage 2: Rust port, closed-loop self-healing, and PGO (45 min). Executing against its blueprint artifact, Antigravity scaffolded the Cargo workspace, rewrote the core solver into 13 Rust modules (each under 240 lines), and drove an automated verification loop in the terminal: running cargo test and cargo clippy , verifying 100-seed parity against the Fortran binary, and profiling a live run with LLVM Profile-Guided Optimization ( llvm-profdata ).
What happened inside the 59-minute run
Catching six silent bugs from the 1990s
During the 3-minute blueprint stage, Argon read the codebase with the eye of a systems engineer and a physicist. At one point, it flagged a comment left in the Fortran momentum solver decades ago:
H1(0) = H1(1) ! STRANGE WAY TO PREVENT DIVISION BY ZERO
In the original code, missing neighbor triangles along the coastline were indexed at cell 0 . To keep boundary loops from dividing by zero, the original authors copied cell 1 ’s water depth into index 0 , quietly tying boundary math to whichever triangle happened to be numbered 1 . Argon figured out why the workaround was there and designed clean boundary types in the Antigravity blueprint artifact.
Fortran also lets you get away with memory hazards that Rust refuses to compile. As Argon mapped FVCOM’s arrays into strict Rust types, six real bugs in the original Fortran and C code surfaced immediately:
Buffer and integer overflows: A mesh builder blindly wrote past a 50-column limit on complex coastal grid points, and a triangle edge check multiplied three node IDs in 32-bit precision, overflowing on grids with more than 1,300 nodes.
Out-of-bounds and uninitialized memory reads: Vertical turbulence and momentum routines read out-of-bounds index 0 elements on shallow grids and pulled uninitialized values from dry land cells.
Unguarded math: A seawater density routine took the square root of salinity without checking if numerical undershoot had nudged it slightly negative during advection.
Self-healing in the Antigravity terminal
Once the Rust modules compiled, Antigravity generated 42 unit, property, and documentation tests, including edge-case geometries designed to stress the physics solver, and executed them in the workspace terminal.
On the first test run, 36 out of 37 unit tests passed. One failed: when simulating a very shallow ocean basin with only a single vertical water layer, the vertical flux solver tried to interpolate with the water layer below it and hit a Rust bounds check.
Because Antigravity runs in a closed agentic loop over the terminal and editor, it didn’t stop to ask a human what to do. The agent captured the panic trace from the test runner, opened the vertical solver module, added a guard for single-layer grids, and re-ran cargo test . Seconds later, all 42 tests passed with 96.8% branch coverage. Antigravity then launched its A/B parity harness against the Fortran baseline across 100 randomized simulation seeds, matching all 100 runs to 12 decimal places (3.91 × 10 -13 ) on the first try.
Where the 1.88x speedup came from
As Antigravity completed each optimization pass, it logged runtime and memory telemetry and generated a standalone interactive HTML scorecard artifact right in the workspace so we could see what actually moved the needle:
The biggest jump came from a bottleneck Argon spotted in the 2D Runge-Kutta loop: the original Fortran code was allocating and freeing eight temporary heap arrays 24 times per time step. Moving those into pre-allocated Struct-of-Arrays workspaces, and having Antigravity run an automated LLVM PGO pass on a live simulation trace, cut runtime nearly in half and dropped memory usage by 54%.
Using Antigravity’s custom skills framework, we packaged this entire multi-stage workflow into an open-source skill in the Google Skills repository on GitHub . It bundles the architectural prompt guardrails, baseline discovery, bug audit, #![forbid(unsafe_code)] Rust scaffolding, self-healing test loop, multi-seed parity harness, PGO pass, and the interactive HTML telemetry artifact generator.
Once installed in Google Antigravity, you can run the full pipeline on any repository with one command:
/migrate-to-rust <github-repo-url>
What comes next
Across science and engineering, millions of lines of hard-won mathematical models are at risk of becoming unmaintainable as the people who wrote them in the 1970s and 1980s retire. Until now, labs had to choose between nursing fragile 40-year-old code or pausing research for years to rewrite it by hand.
With Gemini 4 Argon in Google Antigravity, teams can bring that work forward in an afternoon, keeping the exact same math while running nearly twice as fast in safe, modern Rust.
Share Copy link Share on X Share on Bluesky Share on LinkedIn Stay in sync with Google Antigravity
Get the latest product updates, releases, and feature drops.
All Updates Antigravity App Antigravity CLI IDE Extensions Antigravity SDK Email address Subscribe I accept Google's Terms and Privacy Polic

[truncated]
