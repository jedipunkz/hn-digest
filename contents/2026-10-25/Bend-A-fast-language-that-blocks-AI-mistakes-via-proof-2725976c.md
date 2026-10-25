---
source: "https://github.com/bendlang/bend"
hn_url: "https://news.ycombinator.com/item?id=50237170"
title: "Bend: A fast language that blocks AI mistakes via proof"
article_title: "GitHub - bendlang/bend: Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh · GitHub"
image: "https://opengraph.githubassets.com/e2bc41f5d1c6de845dbd27cb3e5b317608be1f85e73d4f21b950b96d67b29ed8/bendlang/bend"
author: "airhangerf15"
captured_at: "2026-10-25T23:51:39Z"
capture_tool: "hn-digest"
hn_id: 50237170
score: 1
comments: 0
posted_at: "2026-10-25T23:41:04Z"
tags:
  - hacker-news
---

# Bend: A fast language that blocks AI mistakes via proof

- HN: [50237170](https://news.ycombinator.com/item?id=50237170)
- Source: [github.com](https://github.com/bendlang/bend)
- Score: 1
- Comments: 0
- Posted: 2026-10-25T23:41:04Z

## Translation

Title: Bend: A fast language that blocks AI mistakes via proof
Article title: GitHub - bendlang/bend: Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh · GitHub
Description: Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh - bendlang/bend

Article text:
GitHub - bendlang/bend: Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh · GitHub
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
bendlang
bend Public Notifications You must be signed in to change notification settings
Star 24.2k ( 24.2k ) You must be signed in to star a repository
Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh
Readme Apache-2.0 license Activity Custom properties Stars
770 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3,366 Commits 3,366 Commits Folders and files
.github .github bench bench bend2 bend2 demos demos evals evals gates gates guide guide media media paper paper tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md WONTFIX.txt WONTFIX.txt flake.nix flake.nix View all files Repository files navigation
In the post-AGI economy, humans will eventually stop writing and reading code,
but we still need an ambiguity-free language to communicate our intents to the
AIs building the world around us. Bend is that language.
With laws , intents can be more precise than natural language. With
proofs , we can mechanically verify the AI implemented our prompts correctly.
And with a fast compiler , we can run that code at peak compute.
That's Bend - and nothing else.
Target: be as fast as C on the CPU, as fast as CUDA on the GPU. Status:
Thanks to strong types, purity and linearity, Bend compiles to fast executables
as fast as hand-written C (single-core), and even faster (on 10000s cores). The
entire language runs on the GPU, with full memory unification.
Target: outperform every proof assistant by several OOMs. Status:
Bend's compiler is so powerful it can verify mathematical proofs. Usually, this
is slow. Bend is not. It checks, in under a second, files that other projects
would take minutes, making proofs way more practical.
No threads, no locks, no kernels to write. Split the work in two, and Bend
spreads the calls over every core it can find, then joins them back. Below,
pow2(20) divides until one task sits on each of 4,096 GPU cores:
Bend BLOCKS mistakes - with proof
PROBLEM: How can you trust AI code, without reading it?
SOLUTION: By forcing your AI to write a correctness proof .
Bend introduces LAWS.bend , a file where you declare rules that your app must
not break. Bend's compiler then guarantees that these laws always hold, by
demanding mathematical proof whenever your code is edited. For example,
consider a game with one law: winning is impossible . Here's how it plays out:
Law : winning is impossible
So far, it works!
New feature: "Claude, make the board wrap around"
Without LAWS.bend:
Laws broken. AI mistake: merged .
With LAWS.bend:
Laws intact. AI mistake: blocked !
Without LAWS.bend , a bug was merged. With it, the AI had to retry, until no
bugs were left! In this case, it added a wall, but it could have moved the flag,
made the room kill you, or whatever. The only thing it can't do is commit a bug,
because it is mathematically impossible to break laws in LAWS.bend . The
compiler enforces it.
Ask your AI to formalize your app's rules in LAWS.bend . Example:
LAW: "the sum of all balances must be zero"
LAW: "players can never pass through solid walls"
LAW: "list_sort() must always return ascending numbers"
LAW: "array_set() may never be called out-of-bounds"
LAW: "winning is impossible" (the demo above!)
And so on. Anything you can spell can become a law.
Ask your AI to run bend PROOF.bend after editing any code.
That's it. Rejoice as your app never again breaks or violates your rules.
You can also edit LAWS.bend yourself. Here's how it looks:
# LAWS.bend
law you_cant_win : # "winning is impossible"
for moves : List < Game . Move > # any sequence of moves
board = Game . replay ( Game . start (), moves ) # replayed from the start
{ Game . is_won ( board ) == False {} : Bool } # never leads to victory
# PROOF.bend
def Laws . you_cant_win ( moves ):
# ... written by the AI
In short, LAWS.bend is AGENTS.md backed by proof .
With LAWS.bend , "make no mistakes" becomes enforceable.
Skeptical? Edit the demo's code and break the "you can't win" law!
curl -fsSL https://bend-lang.com/install.sh | sh
This bend ignores a project's bunfig.toml and .env ; bun bend2/main.ts
from a checkout reads them, so check untrusted code with bend .
2. Tell your agent to use Bend:
When using Bend:
- run `bend guide` to learn it
- use `LAWS.bend` to keep important rules
- run `bend PROOF.bend` before committing
- parallelize the code whenever possible
Then, just say: "use Bend"!
3. Enjoy bug-free, fast vibe-coded apps!
Ask it to write laws for anything that can't break.
Ask it to parallelize anything you want to be fast.
Bend is young. If anything goes wrong, ask it to open an issue. <3
Bend works best on the back-end, on Linux or macOS.
Syntax == Python + dependent types
import Base
# Performs effects on the CPU.
def main () -> IO ( Unit ):
do IO < Unit > :
name : String < - IO . try ( String , IO . get_env ( "USER" ))
IO . print ( "Hello, " + + name )
Parallelism == divide-and-conquer
import Base
# Computes 2^d in parallel: a tree of d levels, one leaf per unit.
def pow2 ( + d : Nat ) -> U32 :
match d :
case 0 n :
1
case 1 n + p :
a b = pow2 ( p ) pow2 ( p )
( a + b : U32 )
# Runs pow2 on the GPU, via `!`.
def main () -> IO ( Unit ):
result = pow2 !( 20 n )
IO . print ( U32 . show ( result ))
Theorems == laws, Proofs == defs
import Base
# CLAIM: for every nat x, x + 0 equals x.
law add_zero :
for x : Nat
{ Nat . add ( x , 0 n ) == x : Nat }
# PROOF: induction on `x`, one rewrite (`%`) per step.
def add_zero ( x ):
match x :
case 0 n :
{ == }
case 1 n + xp :
% add_zero ( xp ) : { 1 n + Nat . add ( xp , 0 n ) == 1 n + _ : Nat }
{ == }
References
Guide: GUIDE.md , also printed by bend guide .
Demos: demos/ , apps, servers and proofs, each with its LAWS.bend .
Base: base.bend , the base library, also printed by bend base .
Paper: BendTT: An Affine Dependent Type Theory .
Paper: BendRT: A Parallel Runtime for CPUs and GPUs .
Formalization: bendtt.lean , the kernel of --verdict and its proofs, in Lean.
Benches: bench/ , every bench used to make the charts above.
Formatter: bend-fmt-lsp , a formatting-only Bend 2 language server.
Community language server: bend2-lsp , with formatting, diagnostics, and hover.
Community workshop: Bend 2 pocket workshop , the checker and JS compiler in one HTML page, with goals and proof tools, phone first.
Website: https://bend-lang.com
Discord: https://discord.bend-lang.com
Twitter/X: https://x.com/bendlang
Reddit: https://www.reddit.com/r/bendlang/
Issues: https://github.com/bendlang/bend/issues
- Bend 2 is a new language. Bend 1 programs and HVM do not carry over.
- Everything is annotated and nothing is inferred, so code is verbose.
- No type classes, no traits, and no macros beyond compile-time templates.
- Bend has no tactics or proof search; proving theorems takes extra effort.
- Values are affine: closures and arrays cannot be shared.
- Recursion must terminate; `@unsafe` or `def f?(..)` turns the check off.
- Computed matches (`match f(x)`) aren't supported. Must split it manually.
- There is no syntax for if-then-else: a branch is a match on True and False.
- Numbers are Nat, U32 and F32 only: no U64, I64 or F64 (Metal has no f64).
- F32 is axiomatic: nothing about floating point can be proven.
- Strings are linked lists of characters, so text processing is slow.
- Base is small: expect to write helpers other languages ship built in.
- Effects are few: print, env, time, sleep, spawn, channels, files, TCP, UDP.
UDP buffers fit IPv4's 65,507-byte payload limit.
- No TLS, HTTP library, JSON or regex for now (you can add them as foreigns).
- Targets are C, Metal, CUDA and JavaScript; Lua, Luau and Python are planned.
- The JavaScript target runs on one core and has no graphics or audio.
- Parallelism requires balanced calls. Flexible parallelism will be added later.
- Atomic arrays shared across threads are experimental and need `@unsafe`.
- One GPU per program, one event loop, and no multi-machine execution yet.
- One C file per program: no separate compilation, no incremental builds.
- Compiling to native is slow (clang/CUDA/Metal). For fast development, use JS.
- The compiler is young and has blind spots (unusually slow programs). Report.
- We don't have as many benchmarks as we'd like yet, especially for the checker.
- The compiler (not kernel) is 99% AI-written and not yet fully audited.
- The checker has no proof and may have bugs; `--verdict` uses a proven kernel.
- Laws are proven about Bend's semantics. The backends can still differ, for exampl
[truncated]
Bend is created by Victor Taelin and built
by the team:
If you were part of this and your name is missing, please get in touch so we
can add it here.
Thanks to Ayush Somani for reserving the
bend-lang name for us.
Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh
Readme Apache-2.0 license Activity Custom properties Stars
770 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh - bendlang/bend

GitHub - bendlang/bend: Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh · GitHub
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
bendlang
bend Public Notifications You must be signed in to change notification settings
Star 24.2k ( 24.2k ) You must be signed in to star a repository
Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh
Readme Apache-2.0 license Activity Custom properties Stars
770 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3,366 Commits 3,366 Commits Folders and files
.github .github bench bench bend2 bend2 demos demos evals evals gates gates guide guide media media paper paper tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md WONTFIX.txt WONTFIX.txt flake.nix flake.nix View all files Repository files navigation
In the post-AGI economy, humans will eventually stop writing and reading code,
but we still need an ambiguity-free language to communicate our intents to the
AIs building the world around us. Bend is that language.
With laws , intents can be more precise than natural language. With
proofs , we can mechanically verify the AI implemented our prompts correctly.
And with a fast compiler , we can run that code at peak compute.
That's Bend - and nothing else.
Target: be as fast as C on the CPU, as fast as CUDA on the GPU. Status:
Thanks to strong types, purity and linearity, Bend compiles to fast executables
as fast as hand-written C (single-core), and even faster (on 10000s cores). The
entire language runs on the GPU, with full memory unification.
Target: outperform every proof assistant by several OOMs. Status:
Bend's compiler is so powerful it can verify mathematical proofs. Usually, this
is slow. Bend is not. It checks, in under a second, files that other projects
would take minutes, making proofs way more practical.
No threads, no locks, no kernels to write. Split the work in two, and Bend
spreads the calls over every core it can find, then joins them back. Below,
pow2(20) divides until one task sits on each of 4,096 GPU cores:
Bend BLOCKS mistakes - with proof
PROBLEM: How can you trust AI code, without reading it?
SOLUTION: By forcing your AI to write a correctness proof .
Bend introduces LAWS.bend , a file where you declare rules that your app must
not break. Bend's compiler then guarantees that these laws always hold, by
demanding mathematical proof whenever your code is edited. For example,
consider a game with one law: winning is impossible . Here's how it plays out:
Law : winning is impossible
So far, it works!
New feature: "Claude, make the board wrap around"
Without LAWS.bend:
Laws broken. AI mistake: merged .
With LAWS.bend:
Laws intact. AI mistake: blocked !
Without LAWS.bend , a bug was merged. With it, the AI had to retry, until no
bugs were left! In this case, it added a wall, but it could have moved the flag,
made the room kill you, or whatever. The only thing it can't do is commit a bug,
because it is mathematically impossible to break laws in LAWS.bend . The
compiler enforces it.
Ask your AI to formalize your app's rules in LAWS.bend . Example:
LAW: "the sum of all balances must be zero"
LAW: "players can never pass through solid walls"
LAW: "list_sort() must always return ascending numbers"
LAW: "array_set() may never be called out-of-bounds"
LAW: "winning is impossible" (the demo above!)
And so on. Anything you can spell can become a law.
Ask your AI to run bend PROOF.bend after editing any code.
That's it. Rejoice as your app never again breaks or violates your rules.
You can also edit LAWS.bend yourself. Here's how it looks:
# LAWS.bend
law you_cant_win : # "winning is impossible"
for moves : List < Game . Move > # any sequence of moves
board = Game . replay ( Game . start (), moves ) # replayed from the start
{ Game . is_won ( board ) == False {} : Bool } # never leads to victory
# PROOF.bend
def Laws . you_cant_win ( moves ):
# ... written by the AI
In short, LAWS.bend is AGENTS.md backed by proof .
With LAWS.bend , "make no mistakes" becomes enforceable.
Skeptical? Edit the demo's code and break the "you can't win" law!
curl -fsSL https://bend-lang.com/install.sh | sh
This bend ignores a project's bunfig.toml and .env ; bun bend2/main.ts
from a checkout reads them, so check untrusted code with bend .
2. Tell your agent to use Bend:
When using Bend:
- run `bend guide` to learn it
- use `LAWS.bend` to keep important rules
- run `bend PROOF.bend` before committing
- parallelize the code whenever possible
Then, just say: "use Bend"!
3. Enjoy bug-free, fast vibe-coded apps!
Ask it to write laws for anything that can't break.
Ask it to parallelize anything you want to be fast.
Bend is young. If anything goes wrong, ask it to open an issue. <3
Bend works best on the back-end, on Linux or macOS.
Syntax == Python + dependent types
import Base
# Performs effects on the CPU.
def main () -> IO ( Unit ):
do IO < Unit > :
name : String < - IO . try ( String , IO . get_env ( "USER" ))
IO . print ( "Hello, " + + name )
Parallelism == divide-and-conquer
import Base
# Computes 2^d in parallel: a tree of d levels, one leaf per unit.
def pow2 ( + d : Nat ) -> U32 :
match d :
case 0 n :
1
case 1 n + p :
a b = pow2 ( p ) pow2 ( p )
( a + b : U32 )
# Runs pow2 on the GPU, via `!`.
def main () -> IO ( Unit ):
result = pow2 !( 20 n )
IO . print ( U32 . show ( result ))
Theorems == laws, Proofs == defs
import Base
# CLAIM: for every nat x, x + 0 equals x.
law add_zero :
for x : Nat
{ Nat . add ( x , 0 n ) == x : Nat }
# PROOF: induction on `x`, one rewrite (`%`) per step.
def add_zero ( x ):
match x :
case 0 n :
{ == }
case 1 n + xp :
% add_zero ( xp ) : { 1 n + Nat . add ( xp , 0 n ) == 1 n + _ : Nat }
{ == }
References
Guide: GUIDE.md , also printed by bend guide .
Demos: demos/ , apps, servers and proofs, each with its LAWS.bend .
Base: base.bend , the base library, also printed by bend base .
Paper: BendTT: An Affine Dependent Type Theory .
Paper: BendRT: A Parallel Runtime for CPUs and GPUs .
Formalization: bendtt.lean , the kernel of --verdict and its proofs, in Lean.
Benches: bench/ , every bench used to make the charts above.
Formatter: bend-fmt-lsp , a formatting-only Bend 2 language server.
Community language server: bend2-lsp , with formatting, diagnostics, and hover.
Community workshop: Bend 2 pocket workshop , the checker and JS compiler in one HTML page, with goals and proof tools, phone first.
Website: https://bend-lang.com
Discord: https://discord.bend-lang.com
Twitter/X: https://x.com/bendlang
Reddit: https://www.reddit.com/r/bendlang/
Issues: https://github.com/bendlang/bend/issues
- Bend 2 is a new language. Bend 1 programs and HVM do not carry over.
- Everything is annotated and nothing is inferred, so code is verbose.
- No type classes, no traits, and no macros beyond compile-time templates.
- Bend has no tactics or proof search; proving theorems takes extra effort.
- Values are affine: closures and arrays cannot be shared.
- Recursion must terminate; `@unsafe` or `def f?(..)` turns the check off.
- Computed matches (`match f(x)`) aren't supported. Must split it manually.
- There is no syntax for if-then-else: a branch is a match on True and False.
- Numbers are Nat, U32 and F32 only: no U64, I64 or F64 (Metal has no f64).
- F32 is axiomatic: nothing about floating point can be proven.
- Strings are linked lists of characters, so text processing is slow.
- Base is small: expect to write helpers other languages ship built in.
- Effects are few: print, env, time, sleep, spawn, channels, files, TCP, UDP.
UDP buffers fit IPv4's 65,507-byte payload limit.
- No TLS, HTTP library, JSON or regex for now (you can add them as foreigns).
- Targets are C, Metal, CUDA and JavaScript; Lua, Luau and Python are planned.
- The JavaScript target runs on one core and has no graphics or audio.
- Parallelism requires balanced calls. Flexible parallelism will be added later.
- Atomic arrays shared across threads are experimental and need `@unsafe`.
- One GPU per program, one event loop, and no multi-machine execution yet.
- One C file per program: no separate compilation, no incremental builds.
- Compiling to native is slow (clang/CUDA/Metal). For fast development, use JS.
- The compiler is young and has blind spots (unusually slow programs). Report.
- We don't have as many benchmarks as we'd like yet, especially for the checker.
- The compiler (not kernel) is 99% AI-written and not yet fully audited.
- The checker has no proof and may have bugs; `--verdict` uses a proven kernel.
- Laws are proven about Bend's semantics. The backends can still differ, for exampl
[truncated]
Bend is created by Victor Taelin and built
by the team:
If you were part of this and your name is missing, please get in touch so we
can add it here.
Thanks to Ayush Somani for reserving the
bend-lang name for us.
Bend 2: a fast language that blocks AI mistakes via proof. Install: curl -fsSL https://bend-lang.com/install.sh | sh
Readme Apache-2.0 license Activity Custom properties Stars
770 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
