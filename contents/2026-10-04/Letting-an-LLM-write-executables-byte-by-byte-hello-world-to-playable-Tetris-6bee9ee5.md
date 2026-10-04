---
source: "https://github.com/lusob/nlc"
hn_url: "https://news.ycombinator.com/item?id=49954395"
title: "Letting an LLM write executables byte by byte (hello world to playable Tetris)"
article_title: "GitHub - lusob/nlc · GitHub"
image: "https://opengraph.githubassets.com/0cd998af7e48e8f9588d23e0a562835f6f7692ba257d48d7a7cbf9093c96b4e0/lusob/nlc"
author: "lusob"
captured_at: "2026-10-04T15:10:18Z"
capture_tool: "hn-digest"
hn_id: 49954395
score: 1
comments: 0
posted_at: "2026-10-04T14:48:16Z"
tags:
  - hacker-news
---

# Letting an LLM write executables byte by byte (hello world to playable Tetris)

- HN: [49954395](https://news.ycombinator.com/item?id=49954395)
- Source: [github.com](https://github.com/lusob/nlc)
- Score: 1
- Comments: 0
- Posted: 2026-10-04T14:48:16Z

## Translation

Title: Letting an LLM write executables byte by byte (hello world to playable Tetris)
Article title: GitHub - lusob/nlc · GitHub
Description: Contribute to lusob/nlc development by creating an account on GitHub.

Article text:
GitHub - lusob/nlc · GitHub
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
lusob
/
nlc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
docs/ img docs/ img examples examples harness harness results results .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
NLC: an LLM that writes executables directly
No source language. No compiler. No assembler. No linker.
Give an LLM a natural-language spec and it returns the bytes of the executable : ELF header, program header and machine code, by hand. The only thing between the model and the binary is xxd -r .
A 439-byte HTTP server, generated from a one-paragraph spec. That's the real output of the generated binary.
I've been running this experiment every time a new model comes out for a couple of years, and this is the first time the results are good enough to share.
You could say it isn't a compiler in the classical sense, but a model acting as the entire compiler (analysis, code generation, assembling and linking) in a single step, inside a loop that corrects it.
The loop is a bash script of about 100 lines. It hands the spec to the model ( claude -p ), turns whatever comes back into a binary, and tests it for real: it runs it, sends it a curl if it's a server, or compares a screenshot against a reference if it's a game. If it fails, the error goes back to the model exactly as it came out (exit code, output, what readelf sees in the file), and it starts over until the test passes or the attempts run out. Nobody is in the middle.
flowchart TD
S[Natural-language spec] --> M[LLM]
M -->|ELF bytes| X[xxd -r]
X --> T{Real test}
T -- pass --> OK([Binary])
T -- fail --> E[Real error<br/>+ readelf]
E --> M
Loading
To see how far down the stack a model can go with no human in the loop, the same loop runs at three levels:
Level C is the interesting one. A complete 165-byte hello world, byte by byte:
Compared with statically linked C, since these binaries use no library at all:
Caveats: a plain dynamically linked gcc hello.c is ~15 KB, but it depends on libc.so and the dynamic loader, which aren't counted. The 8-9 KB of the assembly versions is mostly the padding ld adds to align to 4 KB pages (the hello world has 47 bytes of code). With musl , static C would shrink noticeably; I haven't measured it.
With several attempts, the loop ends up producing a working program at all three levels, including raw bytes. It's not reliable at the bottom level, though. Five independent runs per case, x86-64, claude-sonnet-5-5 :
(Runs per case that passed within the iteration budget: 5-6 attempts for C, 6-8 for assembly, 8-15 for raw bytes. Median attempts for the passing runs: 1-2 for C and assembly, 1 for the raw-bytes hello world and 5 for the raw-bytes server.)
It also generated a playable Tetris in assembly (X11, no libc, talking to the X server over its Unix socket), about 8 KB, checked by comparing screenshots against a reference at four moments: piece spawn, move, rotate, and lock with a line clear.
Controls: a left, d right, s soft drop, w rotate (click the window to focus it). The keys in the GIF are sent by a script.
The first time I tried raw bytes on x86-64 everything failed: 8/8 attempts on the hello world and 15/15 on the server. The model wrote one extra hex digit in a long run of zeros in the header, so everything after it shifted by half a byte. To readelf it was still a valid ELF, but with absurd sizes, and running it gave a segfault.
A segfault tells the model nothing. I started sending back the real size of the file and what readelf actually decodes (entry point, FileSiz , MemSiz , etc.), and it began passing. I can't say how much of that is the change and how much is luck; I didn't measure it properly. The failed logs are kept as run.v1-failed.log .
Not deterministic. The same spec can come out in one attempt or in six.
Small programs. A 439-byte server isn't a browser; how far this scales, I don't know.
Model-dependent. Results above are Claude Sonnet 5.5 on x86-64 (AArch64 runs used a different model, so the two aren't comparable).
The tests do the work. What works is the model plus the automated verification, not the model alone.
It runs model-generated binaries. Don't run this on a machine you care about without reading what it does first.
Fewer resources. 439 bytes vs ~700 KB for static C. On microcontrollers or embedded systems that matters, but it's a toy server: real savings depend on how much of a program is your own logic and how much a library gives you.
No dependencies. No libc means nothing to go out of date and no inherited vulnerabilities, since the program talks straight to the kernel. But the generated code can have its own bugs and has no years of scrutiny behind it. You trade one risk for another.
Exotic architectures. I've only tried x86-64 and AArch64, which have plenty of compilers, so this is untested. But if a model can hand-encode an architecture's instructions and ELF header, you may not need a compiler for it. I'm leaving that as a question.
The spec as source code. What you maintain is the natural-language spec and the tests; the binary is a regenerable intermediate result. I'm not saying this replaces programming languages, but it changes what gets maintained.
What worries me. A hand-written binary is hard to audit (you can disassemble it, but it's not the same as reading source), and non-determinism makes reproducing it harder. That's why the automated tests aren't an afterthought.
There's also the question of malicious use. A model generating executables that differ every time and carry no compiler or linker fingerprints could slip past signature-based detection better. I don't see this as a radically new risk: modern detection mostly looks at behavior, which stays visible when the bytes change, and a model can already write harmful code in C or assembly. I haven't tried it and it isn't the goal; I'm mentioning it because it should be said.
git clone https://github.com/lusob/nlc && cd nlc
ARCH=x86_64 harness/run_loop_rawbytes.sh examples/hello-world-rawbytes 8 < model >
# <example-dir> [max_iters] [model] (default model: fable)
harness/run_loop.sh examples/hello-world # C, default 5 iters
harness/run_loop_asm.sh examples/hello-world-asm # asm, default 6 iters
harness/run_loop_rawbytes.sh examples/hello-world-rawbytes # raw, default 8 iters
Requirements: Linux (AArch64 or x86-64), Bash, xxd , readelf , curl , python3 , GNU binutils and gcc for levels A/B, and the Claude CLI logged in. ARCH=aarch64|x86_64 picks the target (default: uname -m ); pass your model as the third argument.
Each example directory has everything the loop needs:
examples/<name>/
├── spec.md # natural-language spec (the ONLY input the model gets)
├── spec.x86_64.md # arch-specific spec, when the ABI matters
├── smoke_test.sh # verification gate: takes the binary as $1, exit 0 = pass
├── main.c|.s|.hex # generated source (level-dependent)
└── run.log # full transcript of the loop (iterations + feedback)
x86-64 runs and logs are in results/x86_64/ ; the 5-run repetitions are in results/x86_64-runs/ .
spec.md uses @@XAUTH_COOKIE@@ instead of a cookie; the harness fills it at run time from xauth list $DISPLAY ( harness/xauth_cookie.sh ). The committed main.s files have the cookie bytes zeroed, so they won't authenticate as-is: re-run the loop on your own machine to regenerate them. Under Wayland you need XWayland.
Earlier AArch64 runs (model fable , native Linux AArch64), iteration that passed / budget:
First x86-64 run (single run, claude-sonnet-5-5 ):
Specs deliberately pin down every byte-level detail (struct layouts, endianness, syscall numbers, exact response strings), so the smoke test is unambiguous and the failure feedback is actionable.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to lusob/nlc development by creating an account on GitHub.

GitHub - lusob/nlc · GitHub
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
lusob
/
nlc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
docs/ img docs/ img examples examples harness harness results results .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
NLC: an LLM that writes executables directly
No source language. No compiler. No assembler. No linker.
Give an LLM a natural-language spec and it returns the bytes of the executable : ELF header, program header and machine code, by hand. The only thing between the model and the binary is xxd -r .
A 439-byte HTTP server, generated from a one-paragraph spec. That's the real output of the generated binary.
I've been running this experiment every time a new model comes out for a couple of years, and this is the first time the results are good enough to share.
You could say it isn't a compiler in the classical sense, but a model acting as the entire compiler (analysis, code generation, assembling and linking) in a single step, inside a loop that corrects it.
The loop is a bash script of about 100 lines. It hands the spec to the model ( claude -p ), turns whatever comes back into a binary, and tests it for real: it runs it, sends it a curl if it's a server, or compares a screenshot against a reference if it's a game. If it fails, the error goes back to the model exactly as it came out (exit code, output, what readelf sees in the file), and it starts over until the test passes or the attempts run out. Nobody is in the middle.
flowchart TD
S[Natural-language spec] --> M[LLM]
M -->|ELF bytes| X[xxd -r]
X --> T{Real test}
T -- pass --> OK([Binary])
T -- fail --> E[Real error<br/>+ readelf]
E --> M
Loading
To see how far down the stack a model can go with no human in the loop, the same loop runs at three levels:
Level C is the interesting one. A complete 165-byte hello world, byte by byte:
Compared with statically linked C, since these binaries use no library at all:
Caveats: a plain dynamically linked gcc hello.c is ~15 KB, but it depends on libc.so and the dynamic loader, which aren't counted. The 8-9 KB of the assembly versions is mostly the padding ld adds to align to 4 KB pages (the hello world has 47 bytes of code). With musl , static C would shrink noticeably; I haven't measured it.
With several attempts, the loop ends up producing a working program at all three levels, including raw bytes. It's not reliable at the bottom level, though. Five independent runs per case, x86-64, claude-sonnet-5-5 :
(Runs per case that passed within the iteration budget: 5-6 attempts for C, 6-8 for assembly, 8-15 for raw bytes. Median attempts for the passing runs: 1-2 for C and assembly, 1 for the raw-bytes hello world and 5 for the raw-bytes server.)
It also generated a playable Tetris in assembly (X11, no libc, talking to the X server over its Unix socket), about 8 KB, checked by comparing screenshots against a reference at four moments: piece spawn, move, rotate, and lock with a line clear.
Controls: a left, d right, s soft drop, w rotate (click the window to focus it). The keys in the GIF are sent by a script.
The first time I tried raw bytes on x86-64 everything failed: 8/8 attempts on the hello world and 15/15 on the server. The model wrote one extra hex digit in a long run of zeros in the header, so everything after it shifted by half a byte. To readelf it was still a valid ELF, but with absurd sizes, and running it gave a segfault.
A segfault tells the model nothing. I started sending back the real size of the file and what readelf actually decodes (entry point, FileSiz , MemSiz , etc.), and it began passing. I can't say how much of that is the change and how much is luck; I didn't measure it properly. The failed logs are kept as run.v1-failed.log .
Not deterministic. The same spec can come out in one attempt or in six.
Small programs. A 439-byte server isn't a browser; how far this scales, I don't know.
Model-dependent. Results above are Claude Sonnet 5.5 on x86-64 (AArch64 runs used a different model, so the two aren't comparable).
The tests do the work. What works is the model plus the automated verification, not the model alone.
It runs model-generated binaries. Don't run this on a machine you care about without reading what it does first.
Fewer resources. 439 bytes vs ~700 KB for static C. On microcontrollers or embedded systems that matters, but it's a toy server: real savings depend on how much of a program is your own logic and how much a library gives you.
No dependencies. No libc means nothing to go out of date and no inherited vulnerabilities, since the program talks straight to the kernel. But the generated code can have its own bugs and has no years of scrutiny behind it. You trade one risk for another.
Exotic architectures. I've only tried x86-64 and AArch64, which have plenty of compilers, so this is untested. But if a model can hand-encode an architecture's instructions and ELF header, you may not need a compiler for it. I'm leaving that as a question.
The spec as source code. What you maintain is the natural-language spec and the tests; the binary is a regenerable intermediate result. I'm not saying this replaces programming languages, but it changes what gets maintained.
What worries me. A hand-written binary is hard to audit (you can disassemble it, but it's not the same as reading source), and non-determinism makes reproducing it harder. That's why the automated tests aren't an afterthought.
There's also the question of malicious use. A model generating executables that differ every time and carry no compiler or linker fingerprints could slip past signature-based detection better. I don't see this as a radically new risk: modern detection mostly looks at behavior, which stays visible when the bytes change, and a model can already write harmful code in C or assembly. I haven't tried it and it isn't the goal; I'm mentioning it because it should be said.
git clone https://github.com/lusob/nlc && cd nlc
ARCH=x86_64 harness/run_loop_rawbytes.sh examples/hello-world-rawbytes 8 < model >
# <example-dir> [max_iters] [model] (default model: fable)
harness/run_loop.sh examples/hello-world # C, default 5 iters
harness/run_loop_asm.sh examples/hello-world-asm # asm, default 6 iters
harness/run_loop_rawbytes.sh examples/hello-world-rawbytes # raw, default 8 iters
Requirements: Linux (AArch64 or x86-64), Bash, xxd , readelf , curl , python3 , GNU binutils and gcc for levels A/B, and the Claude CLI logged in. ARCH=aarch64|x86_64 picks the target (default: uname -m ); pass your model as the third argument.
Each example directory has everything the loop needs:
examples/<name>/
├── spec.md # natural-language spec (the ONLY input the model gets)
├── spec.x86_64.md # arch-specific spec, when the ABI matters
├── smoke_test.sh # verification gate: takes the binary as $1, exit 0 = pass
├── main.c|.s|.hex # generated source (level-dependent)
└── run.log # full transcript of the loop (iterations + feedback)
x86-64 runs and logs are in results/x86_64/ ; the 5-run repetitions are in results/x86_64-runs/ .
spec.md uses @@XAUTH_COOKIE@@ instead of a cookie; the harness fills it at run time from xauth list $DISPLAY ( harness/xauth_cookie.sh ). The committed main.s files have the cookie bytes zeroed, so they won't authenticate as-is: re-run the loop on your own machine to regenerate them. Under Wayland you need XWayland.
Earlier AArch64 runs (model fable , native Linux AArch64), iteration that passed / budget:
First x86-64 run (single run, claude-sonnet-5-5 ):
Specs deliberately pin down every byte-level detail (struct layouts, endianness, syscall numbers, exact response strings), so the smoke test is unambiguous and the failure feedback is actionable.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
