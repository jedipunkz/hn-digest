---
source: "https://github.com/HiteSit/real-scientist-skills"
hn_url: "https://news.ycombinator.com/item?id=50506529"
title: "Skills for Real Scientists – Steer Away AI Slop"
article_title: "GitHub - HiteSit/real-scientist-skills: Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked. · GitHub"
image: "https://opengraph.githubassets.com/ea6f6c4170ae222d8db5cc4aa4b1c03c125d1716e622e6300cf9f4a71d4f769c/HiteSit/real-scientist-skills"
author: "hitesit"
captured_at: "2026-11-15T13:31:12Z"
capture_tool: "hn-digest"
hn_id: 50506529
score: 1
comments: 0
posted_at: "2026-11-15T12:42:08Z"
tags:
  - hacker-news
---

# Skills for Real Scientists – Steer Away AI Slop

- HN: [50506529](https://news.ycombinator.com/item?id=50506529)
- Source: [github.com](https://github.com/HiteSit/real-scientist-skills)
- Score: 1
- Comments: 0
- Posted: 2026-11-15T12:42:08Z

## Translation

Title: Skills for Real Scientists – Steer Away AI Slop
Article title: GitHub - HiteSit/real-scientist-skills: Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked. · GitHub
Description: Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked. - HiteSit/real-scientist-skills

Article text:
GitHub - HiteSit/real-scientist-skills: Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
HiteSit
real-scientist-skills Public Notifications You must be signed in to change notification settings
Star 27 ( 27 ) You must be signed in to star a repository
Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked.
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
583 Commits 583 Commits Folders and files
.agents .agents .githooks .githooks .github/ workflows .github/ workflows .out-of-scope .out-of-scope assets assets customizations customizations customized customized docs docs scripts scripts skills skills .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md GLOSSARY.md GLOSSARY.md LICENSE LICENSE MAINTAINING.md MAINTAINING.md README.md README.md SCOPE.md SCOPE.md fork-removed.txt fork-removed.txt View all files Repository files navigation
Built on Matt Pocock's skills for real engineers .
Agent skills for reproducible scientific experiments with any coding agent. Agree on one experiment at a time, record the expected result before running it, and use the evidence to choose what to run next. Every run leaves a Git-tracked lab notebook and rerunnable files.
Matt's engineering flow plans the whole route first: /to-spec writes down the goal, /to-tickets splits it into tickets, and /implement-spec builds them, many in parallel. That works because software is predictable once you agree on what to build.
Research and scientific exploration aren't. A plan is a map, and in research nobody has mapped the territory yet. You learn what's relevant by running something and looking at the result: each result shows you more of the territory, and it often shows that the route you planned goes nowhere. So the science flow plans one experiment, called a probe, at a time, and you choose each next probe after seeing the last result. The study's notebook is the map of your scientific question, redrawn after every probe.
Ask a language model a question without steering it, and it gives you the most common answer, the middle of everything it has read. New findings sit away from that middle, and only a clear intention moves the work away from it, away from what you could call AI slop. So before each run the agent writes down the standard answer, and the result shows how far from it reality lands. That distance is what makes the work worth doing, and a rough measure of how unique it is.
When a protocol stops changing, it's predictable again, and Matt's engineering flow takes over.
Everything the science skills write goes in a .science/ folder at the root of your repo:
.science/
SETUP.md your setup
library/ what you read: papers, software docs; one page per question, the raw search results
data/ your input data (or data/ at the repo root, your choice)
<study>/ one folder per study
STUDY.md the starting direction and first probe, frozen; promotions and the outcome are added at the end
NOTEBOOK.md the map: threads of the question, each probe's asked / expected / found / means, a log of every operation
dashboard.html the study's dashboard
probes/ one folder per probe, named like 03-<slug>
Your code stays where it always was. Probes' work reaches it only through /promote , after you agree where each file goes. All studies share one branch, and every probe gets a Git tag, so you can go back to any version.
You type /setup-science , /experiment , /promote and /ask-science . The agent loads lab-notebook and desk-research on its own when the task fits.
setup-science : set up a repo for the science skills. Run once.
experiment : load or open a study and work on it, one probe at a time.
promote : move the work of one or more probes into the protocol, after a round of questions and a tree you approve.
lab-notebook : the study's layout and formats: notebook, probe folders, result briefs, plots and dashboard.
desk-research : answer questions by reading papers and documentation in parallel, into a shared library, with citations.
ask-science : step back, bring an older study to the current format, and check it is on track.
The flow step by step, the principles behind it and the vocabulary: skills/science/README.md .
Matt's engineering and productivity skills are included as he wrote them, and the science skills call several of them ( grilling , domain-modeling , tdd , code-review , pr ). For what each one does, see his README .
Paste this prompt into your coding agent. It installs the six science skills, and Matt's skills too if you don't have them yet:
Install the science skills from https://github.com/HiteSit/real-scientist-skills.
1. Clone the repo, or pull it if I already have a clone.
2. Find the folder you load skills from.
3. Check that folder for Matt Pocock's skills, such as grilling, tdd and domain-modeling.
- If they're there, keep them and install only the six skills in skills/science/.
- If they're missing, install the six science skills plus every skill in skills/engineering/ and skills/productivity/. When a skill has a folder in customized/, use that folder instead.
4. Install each skill as a symlink into the clone, so a git pull updates it. If a skill with the same name is already in the folder, ask me before replacing it.
5. Tell me what you installed and what you skipped. Then remind me to run /setup-science in my repo, and /setup-matt-pocock-skills if you installed Matt's skills.
To update, ask your agent to pull the clone.
Every week this repo merges Matt's latest skills, and it keeps a few small edits to two of them. MAINTAINING.md explains how.
The engineering and productivity skills, the repo's structure, and the ideas the science bucket grows from are the work of Matt Pocock . If you like these skills, read his README and join his newsletter .
Everything under skills/ except skills/science/ is Matt's, from mattpocock/skills , merged in unchanged every week.
customized/ holds edited copies of two of his skills, grilling and chief-of-staff , rebuilt from his originals plus the edits in customizations/ .
skills/science/ and the docs around it are new here.
Matt's skills are MIT licensed, and so is this repo. His LICENSE is kept as he wrote it.
Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked.
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked. - HiteSit/real-scientist-skills

GitHub - HiteSit/real-scientist-skills: Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
HiteSit
real-scientist-skills Public Notifications You must be signed in to change notification settings
Star 27 ( 27 ) You must be signed in to star a repository
Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked.
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
583 Commits 583 Commits Folders and files
.agents .agents .githooks .githooks .github/ workflows .github/ workflows .out-of-scope .out-of-scope assets assets customizations customizations customized customized docs docs scripts scripts skills skills .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md GLOSSARY.md GLOSSARY.md LICENSE LICENSE MAINTAINING.md MAINTAINING.md README.md README.md SCOPE.md SCOPE.md fork-removed.txt fork-removed.txt View all files Repository files navigation
Built on Matt Pocock's skills for real engineers .
Agent skills for reproducible scientific experiments with any coding agent. Agree on one experiment at a time, record the expected result before running it, and use the evidence to choose what to run next. Every run leaves a Git-tracked lab notebook and rerunnable files.
Matt's engineering flow plans the whole route first: /to-spec writes down the goal, /to-tickets splits it into tickets, and /implement-spec builds them, many in parallel. That works because software is predictable once you agree on what to build.
Research and scientific exploration aren't. A plan is a map, and in research nobody has mapped the territory yet. You learn what's relevant by running something and looking at the result: each result shows you more of the territory, and it often shows that the route you planned goes nowhere. So the science flow plans one experiment, called a probe, at a time, and you choose each next probe after seeing the last result. The study's notebook is the map of your scientific question, redrawn after every probe.
Ask a language model a question without steering it, and it gives you the most common answer, the middle of everything it has read. New findings sit away from that middle, and only a clear intention moves the work away from it, away from what you could call AI slop. So before each run the agent writes down the standard answer, and the result shows how far from it reality lands. That distance is what makes the work worth doing, and a rough measure of how unique it is.
When a protocol stops changing, it's predictable again, and Matt's engineering flow takes over.
Everything the science skills write goes in a .science/ folder at the root of your repo:
.science/
SETUP.md your setup
library/ what you read: papers, software docs; one page per question, the raw search results
data/ your input data (or data/ at the repo root, your choice)
<study>/ one folder per study
STUDY.md the starting direction and first probe, frozen; promotions and the outcome are added at the end
NOTEBOOK.md the map: threads of the question, each probe's asked / expected / found / means, a log of every operation
dashboard.html the study's dashboard
probes/ one folder per probe, named like 03-<slug>
Your code stays where it always was. Probes' work reaches it only through /promote , after you agree where each file goes. All studies share one branch, and every probe gets a Git tag, so you can go back to any version.
You type /setup-science , /experiment , /promote and /ask-science . The agent loads lab-notebook and desk-research on its own when the task fits.
setup-science : set up a repo for the science skills. Run once.
experiment : load or open a study and work on it, one probe at a time.
promote : move the work of one or more probes into the protocol, after a round of questions and a tree you approve.
lab-notebook : the study's layout and formats: notebook, probe folders, result briefs, plots and dashboard.
desk-research : answer questions by reading papers and documentation in parallel, into a shared library, with citations.
ask-science : step back, bring an older study to the current format, and check it is on track.
The flow step by step, the principles behind it and the vocabulary: skills/science/README.md .
Matt's engineering and productivity skills are included as he wrote them, and the science skills call several of them ( grilling , domain-modeling , tdd , code-review , pr ). For what each one does, see his README .
Paste this prompt into your coding agent. It installs the six science skills, and Matt's skills too if you don't have them yet:
Install the science skills from https://github.com/HiteSit/real-scientist-skills.
1. Clone the repo, or pull it if I already have a clone.
2. Find the folder you load skills from.
3. Check that folder for Matt Pocock's skills, such as grilling, tdd and domain-modeling.
- If they're there, keep them and install only the six skills in skills/science/.
- If they're missing, install the six science skills plus every skill in skills/engineering/ and skills/productivity/. When a skill has a folder in customized/, use that folder instead.
4. Install each skill as a symlink into the clone, so a git pull updates it. If a skill with the same name is already in the folder, ask me before replacing it.
5. Tell me what you installed and what you skipped. Then remind me to run /setup-science in my repo, and /setup-matt-pocock-skills if you installed Matt's skills.
To update, ask your agent to pull the clone.
Every week this repo merges Matt's latest skills, and it keeps a few small edits to two of them. MAINTAINING.md explains how.
The engineering and productivity skills, the repo's structure, and the ideas the science bucket grows from are the work of Matt Pocock . If you like these skills, read his README and join his newsletter .
Everything under skills/ except skills/science/ is Matt's, from mattpocock/skills , merged in unchanged every week.
customized/ holds edited copies of two of his skills, grilling and chief-of-staff , rebuilt from his originals plus the edits in customizations/ .
skills/science/ and the docs around it are new here.
Matt's skills are MIT licensed, and so is this repo. His LICENSE is kept as he wrote it.
Agent skills for scientists: one experiment at a time, each result deciding the next, every step tracked.
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
