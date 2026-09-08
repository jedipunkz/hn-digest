---
source: "https://github.com/doldecomp/melee"
hn_url: "https://news.ycombinator.com/item?id=49612081"
title: "Super Smash Brothers Melee has been 100% decompilated with the help of LLMs"
article_title: "GitHub - doldecomp/melee: A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks. · GitHub"
image: "https://opengraph.githubassets.com/f8f64ecb61d72239b5a781655e6d7e69dc2232400c82d0565b7e156344ca6759/doldecomp/melee"
author: "oveja"
captured_at: "2026-09-08T16:06:55Z"
capture_tool: "hn-digest"
hn_id: 49612081
score: 1
comments: 0
posted_at: "2026-09-08T15:49:30Z"
tags:
  - hacker-news
---

# Super Smash Brothers Melee has been 100% decompilated with the help of LLMs

- HN: [49612081](https://news.ycombinator.com/item?id=49612081)
- Source: [github.com](https://github.com/doldecomp/melee)
- Score: 1
- Comments: 0
- Posted: 2026-09-08T15:49:30Z

## Translation

Title: Super Smash Brothers Melee has been 100% decompilated with the help of LLMs
Article title: GitHub - doldecomp/melee: A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks. · GitHub
Description: A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks. - doldecomp/melee

Article text:
GitHub - doldecomp/melee: A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
doldecomp
/
melee
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
3,759 Commits 3,759 Commits Folders and files
.cargo .cargo .github .github .idea .idea .nix .nix .vscode .vscode config/ GALE01 config/ GALE01 docs docs extern extern orig/ GALE01/ sys orig/ GALE01/ sys reqs reqs src src tools tools .clang-format .clang-format .clang-tidy .clang-tidy .clangd .clangd .dockerignore .dockerignore .editorconfig .editorconfig .flake8 .flake8 .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .rustfmt.toml .rustfmt.toml Cargo.lock Cargo.lock Cargo.toml Cargo.toml Doxyfile Doxyfile configure.py configure.py default.nix default.nix flake.lock flake.lock flake.nix flake.nix View all files Repository files navigation
This repo contains a WIP decompilation of Super Smash Bros Melee (US).
The DOL this repository builds can be shifted! Meaning you are able to now add and remove code as you see fit, for modding or research purposes.
On Windows, it's highly recommended to use native tooling. WSL or msys2 are not required.
When running under WSL, objdiff is unable to get filesystem notifications for automatic rebuilds.
Install Python and add it to %PATH% .
Also available from the Windows Store .
Download ninja and add it to %PATH% .
Quick install via pip: pip install ninja
Install ninja :
brew install ninja
Install wine-crossover :
brew install --cask --no-quarantine gcenx/wine/wine-crossover
After OS upgrades, if macOS complains about Wine Crossover.app being unverified, you can unquarantine it using:
sudo xattr -rd com.apple.quarantine ' /Applications/Wine Crossover.app '
Linux:
For non-x86(_64) platforms: Install wine from your package manager.
For x86(_64), WiBo , a minimal 32-bit Windows binary wrapper, will be automatically downloaded and used.
Clone the repository:
git clone https://github.com/doldecomp/melee.git --depth=1
Using Dolphin Emulator , find your ISO and click Properties . Go to the Filesystem tab, right-click Disc - GALE01 and select Extract System Data . Choose orig/GALE01 of this repository.
To save space, only main.dol (and .gitkeep ) are necessary. Other files can be deleted.
Configure:
python configure.py
We use Python for our command line tooling. It is recommended that you use a virtual environment .
Create a virtual environment.
python -m venv --upgrade-deps ' .venv '
You'll need to activate it whenever you open a new shell.
Windows:
.venv / Scripts / Activate.ps1
Linux/macOS:
. .venv / bin / activate
After that, you can install or update our packages with:
pip install -r reqs/decomp.txt
Now you can run decomp.py to decomp a function using m2c . Pass it -h to see all the options.
python tools/decomp.py my_function_name
Once the initial build succeeds, an objdiff.json should exist in the project root.
Download the latest release from encounter/objdiff . Under project settings, set Project directory . The configuration should be loaded automatically.
Select an object from the left sidebar to begin diffing. Changes to the project will rebuild automatically: changes to source files, headers, configure.py , splits.txt or symbols.txt .
It's recommended that you enable the Relax relocation diffs option under Diff Options .
Contributions are welcome! If you're new to decomp, check out our Getting Started guide . Before opening a pull request , please read our contributing guidelines . If you're new to Git and don't know how to create a pull request, we encourage you to create an issue with your decomp.me link and a maintainer will add your code to the repository.
We're also happy to answer any questions in the #smash-bros-melee channel on Discord.
How is the codebase structured?
The code in src is divided into several modules, the main one being melee , which is the game code.
The main game code is divided into several two-letter folders, which were left behind by HAL in assert messages and game data on the original disc.
HAL also used two-letter abbreviations for each fighter.
1 Zako (雑魚) is Japanese for "trash mob" in video games, literally "small fish."
The Metrowerks Target Resident Kernel.
The Metrowerks Standard Library.
What can be done after decompiling Melee?
Note that this project's purpose is to only match the ASM with C code. This is entirely for research and archival purposes. After this is created, you essentially have a C project that can be compiled into Melee, but it won't be portable (aka you can't compile it to run on a normal computer).
So creating mods would be a lot easier as C code is much easier to consume than ASM. However, there are additional projects that could be undertaken once this is complete, but those technical endeavours are out-of-scope for this repo.
Do we know how the compiler works?
Kind of. We don’t have its source though.
How do we get the compiler to pick a certain register allocation?
Considering we don't have the source for the compiler, this is kind of "anything goes" territory. Unfortunately register allocation is an NP-hard problem which means there are all types of heuristics you can use to select registers, some of which can be confused by things as silly as variable names.
One option is to attempt to automatically permute the source code to get the correct register allocation.
A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks.
Contributing Activity Custom properties Stars
220 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks. - doldecomp/melee

GitHub - doldecomp/melee: A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
doldecomp
/
melee
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
3,759 Commits 3,759 Commits Folders and files
.cargo .cargo .github .github .idea .idea .nix .nix .vscode .vscode config/ GALE01 config/ GALE01 docs docs extern extern orig/ GALE01/ sys orig/ GALE01/ sys reqs reqs src src tools tools .clang-format .clang-format .clang-tidy .clang-tidy .clangd .clangd .dockerignore .dockerignore .editorconfig .editorconfig .flake8 .flake8 .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .rustfmt.toml .rustfmt.toml Cargo.lock Cargo.lock Cargo.toml Cargo.toml Doxyfile Doxyfile configure.py configure.py default.nix default.nix flake.lock flake.lock flake.nix flake.nix View all files Repository files navigation
This repo contains a WIP decompilation of Super Smash Bros Melee (US).
The DOL this repository builds can be shifted! Meaning you are able to now add and remove code as you see fit, for modding or research purposes.
On Windows, it's highly recommended to use native tooling. WSL or msys2 are not required.
When running under WSL, objdiff is unable to get filesystem notifications for automatic rebuilds.
Install Python and add it to %PATH% .
Also available from the Windows Store .
Download ninja and add it to %PATH% .
Quick install via pip: pip install ninja
Install ninja :
brew install ninja
Install wine-crossover :
brew install --cask --no-quarantine gcenx/wine/wine-crossover
After OS upgrades, if macOS complains about Wine Crossover.app being unverified, you can unquarantine it using:
sudo xattr -rd com.apple.quarantine ' /Applications/Wine Crossover.app '
Linux:
For non-x86(_64) platforms: Install wine from your package manager.
For x86(_64), WiBo , a minimal 32-bit Windows binary wrapper, will be automatically downloaded and used.
Clone the repository:
git clone https://github.com/doldecomp/melee.git --depth=1
Using Dolphin Emulator , find your ISO and click Properties . Go to the Filesystem tab, right-click Disc - GALE01 and select Extract System Data . Choose orig/GALE01 of this repository.
To save space, only main.dol (and .gitkeep ) are necessary. Other files can be deleted.
Configure:
python configure.py
We use Python for our command line tooling. It is recommended that you use a virtual environment .
Create a virtual environment.
python -m venv --upgrade-deps ' .venv '
You'll need to activate it whenever you open a new shell.
Windows:
.venv / Scripts / Activate.ps1
Linux/macOS:
. .venv / bin / activate
After that, you can install or update our packages with:
pip install -r reqs/decomp.txt
Now you can run decomp.py to decomp a function using m2c . Pass it -h to see all the options.
python tools/decomp.py my_function_name
Once the initial build succeeds, an objdiff.json should exist in the project root.
Download the latest release from encounter/objdiff . Under project settings, set Project directory . The configuration should be loaded automatically.
Select an object from the left sidebar to begin diffing. Changes to the project will rebuild automatically: changes to source files, headers, configure.py , splits.txt or symbols.txt .
It's recommended that you enable the Relax relocation diffs option under Diff Options .
Contributions are welcome! If you're new to decomp, check out our Getting Started guide . Before opening a pull request , please read our contributing guidelines . If you're new to Git and don't know how to create a pull request, we encourage you to create an issue with your decomp.me link and a maintainer will add your code to the repository.
We're also happy to answer any questions in the #smash-bros-melee channel on Discord.
How is the codebase structured?
The code in src is divided into several modules, the main one being melee , which is the game code.
The main game code is divided into several two-letter folders, which were left behind by HAL in assert messages and game data on the original disc.
HAL also used two-letter abbreviations for each fighter.
1 Zako (雑魚) is Japanese for "trash mob" in video games, literally "small fish."
The Metrowerks Target Resident Kernel.
The Metrowerks Standard Library.
What can be done after decompiling Melee?
Note that this project's purpose is to only match the ASM with C code. This is entirely for research and archival purposes. After this is created, you essentially have a C project that can be compiled into Melee, but it won't be portable (aka you can't compile it to run on a normal computer).
So creating mods would be a lot easier as C code is much easier to consume than ASM. However, there are additional projects that could be undertaken once this is complete, but those technical endeavours are out-of-scope for this repo.
Do we know how the compiler works?
Kind of. We don’t have its source though.
How do we get the compiler to pick a certain register allocation?
Considering we don't have the source for the compiler, this is kind of "anything goes" territory. Unfortunately register allocation is an NP-hard problem which means there are all types of heuristics you can use to select registers, some of which can be confused by things as silly as variable names.
One option is to attempt to automatically permute the source code to get the correct register allocation.
A decompilation of Super Smash Bros Melee brought to you by a bunch of clever folks.
Contributing Activity Custom properties Stars
220 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
