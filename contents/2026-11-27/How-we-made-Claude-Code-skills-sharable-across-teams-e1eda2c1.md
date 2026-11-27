---
source: "https://cleancommit.io/blog/claude-code-skills/"
hn_url: "https://news.ycombinator.com/item?id=50668619"
title: "How we made Claude Code skills sharable across teams?"
article_title: "How we made Claude Code skills sharable across teams"
image: "https://cdn.cleancommit.io/blog/2026/11/claude-code-skills-cover.png"
author: "mrkaluzny"
captured_at: "2026-11-27T17:24:28Z"
capture_tool: "hn-digest"
hn_id: 50668619
score: 1
comments: 0
posted_at: "2026-11-27T16:24:39Z"
tags:
  - hacker-news
---

# How we made Claude Code skills sharable across teams?

- HN: [50668619](https://news.ycombinator.com/item?id=50668619)
- Source: [cleancommit.io](https://cleancommit.io/blog/claude-code-skills/)
- Score: 1
- Comments: 0
- Posted: 2026-11-27T16:24:39Z

## Translation

Title: How we made Claude Code skills sharable across teams?
Article title: How we made Claude Code skills sharable across teams
Description: What Claude Code skills are, how they differ from plugins, and two ways to share them with your team: a private plugin marketplace or the CLI we built.

Article text:
How we made Claude Code skills sharable across teams Browse tests that generated $7.3M in revenue. Live now Clean Commit Case Studies Resources Blog Contact Book your free strategy call Open main menu Conversion Rate Optimization Case Studies Resources A/B Experiments Library Tools Blog Contact Talk to a human Book your free strategy call Clean Commit's dribbble Clean Commit's github Clean Commit's linkedin Clean Commit's facebook Clean Commit's instagram Back to blog Blog Development How we made Claude Code skills sharable across teams How we made Claude Code skills sharable across teams
Find out where your store is leaking revenue.
Get the audit that uncovered millions in missed revenue for $25M+ eComm stores.
Claude Code skills are code your whole team runs, so they should be shared, reviewed and versioned like code, not copied from one laptop to the next. We run 45 Claude Code skills across a team of strategists, developers, writers and sales. A shared folder of skills stopped working long before we got to 45. So we built cc-skills , a small CLI that installs the team's skills from one GitHub repo into every coding agent on a machine, keeps them current with one command, and lets each person pick only the skills their job needs. This post covers what skills are, how they differ from plugins, where a shared folder broke for us, and what the CLI does that a plugin doesn't.
A skill is a folder with a SKILL.md file in it. The file starts with a short frontmatter block, a name and a description , followed by instructions in plain markdown. The folder can also hold scripts, reference documents and templates the instructions point to.
Claude Code reads only the name and description of every installed skill when a session starts. When your request matches a description, it loads the full SKILL.md and follows it. You can also call a skill directly by typing a slash and its name. That's why the description matters more than anything else in the file: it decides whether the skill ever gets used.
Skills live in ~/.claude/skills for you personally, or in .claude/skills inside a project. As of November 2026 the same format is an open standard, Agent Skills , so Codex , Cursor, GitHub Copilot, OpenCode and others read the same folders. A skill written once works in most coding agents, which turned out to matter a lot for how we distribute them.
A skill and a plugin aren't alternatives. A skill is one piece of know-how. A plugin is a package that installs several pieces at once, and skills are usually what's inside it.
A Claude Code plugin is a folder with a .claude-plugin/plugin.json manifest. It can hold skills, subagents, hooks, MCP servers, slash commands and executables for the agent's shell. You install it from a marketplace , which is a git repo that lists plugins, with /plugin install name@marketplace . Figma's plugin is a typical one: one install brings the Figma MCP server and around 15 skills.
Plugin skills are namespaced. A skill called review inside my-plugin runs as /my-plugin:review , while a skill in ~/.claude/skills keeps its plain name.
Plugins can be private. The marketplace repo can be a private GitHub repo , and Claude Code clones it with the git credentials already on your machine. A team can add the marketplace to a project's .claude/settings.json so teammates are offered it when they trust the folder, and on Team and Enterprise plans an admin can sync a private repo through the organization's GitHub connection. Marketplaces outside Anthropic's own don't auto-update by default, so each person turns that on or runs /plugin marketplace update .
Where each one works is what decided it for us. Other agents have their own plugin formats, but they all read plain skill folders. A Claude Code plugin only reaches the people on the team who work in Claude Code.
Where a shared folder of skills broke
Our skills started in March 2026 as a git repo with a bash install script. Clone it, run the script, and it symlinked everything into ~/.claude/skills . That worked for one person and about a dozen skills.
Then the team started writing their own. Strategists, developers and writers each added skills for their own work. Within a few months we had the problems every team with shared skills runs into:
Skills drifted. People edited a copy in their own ~/.claude/skills to fix something on a client job, and the fix never made it back. Two people would run "the same" skill and get different results.
Nobody knew what version they had. A change to how one skill records its work had to reach every developer before the next client job, and there was no way to check who had it.
Everyone got everything. A copywriter doesn't need a developer's skills in every session, and every installed description is something the agent has to read and choose between.
It only covered Claude Code. Not everyone on the team works in Claude Code, and Codex, Cursor and the rest each keep skills in their own folder.
One command installs the team's skills
cc-skills is a TypeScript CLI with no dependencies. Node 24 runs it directly, so there's no build step and nothing to install beyond cloning the repo. Running it opens one screen with four tabs: the agents it detected on your machine, packs, individual skills and tools.
It links skills rather than copying them, so the clone on your machine is the install. cc-skills update pulls the repo and relinks. New skills in your packs appear, moved skills get repointed, and retired skills are unlinked with the reason printed. That last part matters more than it sounds: when we retire a skill, nobody keeps running the old one by accident.
It links into every agent it finds: Claude Code, Codex, the universal ~/.agents/skills folder, OpenCode, GitHub Copilot and Cursor. A developer who switches from Claude Code to Codex for a task has the same skills there.
It also handles the drift problem directly. If you already have your own copy of a skill where the link should go, the CLI compares it with the repo version. An identical copy gets replaced by the link. A different one is shown as a conflict with the files that differ and is left alone unless you say yes, in which case it keeps a timestamped backup. And cc-skills import goes the other way: it finds skills you wrote locally, asks which pack each belongs to, copies them into the repo and swaps your local copy for a link.
A GitHub repo for versions, ideas and issues
Because every skill lives in one repo, a skill is reviewed like code. Changes go through a pull request with a short template, CI runs cc-skills check (directory names match skill names, links between skills resolve, the catalogue is current), and every release gets a changelog entry that says what changed for the person who installs it. Version 1.0.0 shipped in October 2026; we're on 1.6 now.
This is what fixed "which version do you have". When we changed one development skill so that two people working on the same client project stopped creating merge conflicts, the change reached every developer at their next cc-skills update . Nobody had to be told to copy anything.
It also gives everyone a place to put an idea or a complaint. Anyone can open an issue when a skill gets something wrong, or a PR when they've fixed it. So far 18 pull requests have come from four people, and not all of them are developers.
Retiring a skill is a git mv into an archive folder with a line explaining why. We archived 11 skills after finding none of them had been used in 31 days: a few were replaced by better ones, a few were built for one client, and two depended on a site that blocks automated access.
Packs: install only what your job needs
Not everyone uses every skill. Skills are grouped into packs by role, not by topic: one for strategists, one for developers, one for writers, and smaller ones for marketing, sales and our internal knowledge base. The strategist pack is the biggest, at 21 skills.
You tick the packs for your job and get all their skills. On top of that you can tick a single skill from another pack, or untick one inside yours, and the CLI remembers those as overrides. When we add a new pack, everyone is offered it once at their next update, and never asked again.
This keeps each person's agent focused. A strategist's session sees research skills, not front-end coding conventions. Fewer descriptions to choose from means the agent picks the right skill more often.
Tools for what MCP servers don't cover
The last piece came later. Official MCP servers are fine, but sometimes the one we need doesn't exist, or doesn't do what we need, so we built tools we can trust. A tool is code, not instructions: a command-line script with a short skill that tells the agent when to run it. Two examples:
sentry-issues lists errors across every client's Sentry organization with one token, worst project first. Sentry has its own MCP server , but we work across many client organizations and wanted one command that answers "what broke this week, everywhere".
rybbit pulls visitor analytics per client site from Rybbit . Adding up every site tells us how many real visitors a new browser target would cut off before we raise it.
Tools also save tokens. When an agent works through an MCP server, every step is a round trip through the model: call the server, read the JSON, decide the next call, reshape the output. For a task we repeat every week, like a Sentry triage across every client project, that's a lot of tokens spent on fetching and formatting. A script does the same work the same way every time and hands the agent a short table to reason about. The agent's job becomes deciding what matters, not paging through API responses.
cc-skills install --tool sentry,rybbit links each command onto your PATH and its skill into every agent. The CLI never asks for, reads or writes API keys. It only tells you which ones are missing.
What changed for the team and our clients
Setting up a new machine is cc-skills install --pack dev -y . A fix to a skill reaches everyone in one update, so a lesson from one client job carries over to the next. Skills people wrote on their own now land in the repo through cc-skills import instead of staying on one laptop.
For the stores we run CRO programs for, the difference is consistency. Every developer builds and rolls out an A/B test the same way, and every strategist runs a CRO audit the same way, because they run the same skills. When we change how we do something, we change the skill once.
When a plain folder or a plugin is enough
You don't need any of this if you're one person, or a small team that all use Claude Code and the same handful of skills. A folder in ~/.claude/skills , or .claude/skills committed to a project, is simpler and works. If your skills only matter inside one codebase, keep them in that codebase.
If your whole team works in Claude Code, start with a private plugin marketplace, one plugin per role. That gives you versioned installs, updates and bundles per role without writing anything. We'd have done the same if Claude Code were the only agent we used. What it doesn't give you is the other agents, single skills on top of a bundle, or a way to bring someone's local copy back into the repo.
A CLI earned its keep for us once three things were true: several people writing skills, people in different roles needing different sets, and more than one coding agent in use. The last one is what rules out a plugin. If your team uses more than one agent, put your skills in one repo before they drift, and write the install script before the folder gets past a dozen skills.
Where a shared folder of skills broke
One command installs the team's skills
A GitHub repo for versions, ideas and issues
Packs: install only what your job needs
Tools for what MCP servers don't cover
What changed for the team and our clients
When a plain folder or a plugin is enough
Where a shared folder of skills broke
One command installs the team's skills
A GitHub repo for versions, id

[truncated]

## Original Extract

What Claude Code skills are, how they differ from plugins, and two ways to share them with your team: a private plugin marketplace or the CLI we built.

How we made Claude Code skills sharable across teams Browse tests that generated $7.3M in revenue. Live now Clean Commit Case Studies Resources Blog Contact Book your free strategy call Open main menu Conversion Rate Optimization Case Studies Resources A/B Experiments Library Tools Blog Contact Talk to a human Book your free strategy call Clean Commit's dribbble Clean Commit's github Clean Commit's linkedin Clean Commit's facebook Clean Commit's instagram Back to blog Blog Development How we made Claude Code skills sharable across teams How we made Claude Code skills sharable across teams
Find out where your store is leaking revenue.
Get the audit that uncovered millions in missed revenue for $25M+ eComm stores.
Claude Code skills are code your whole team runs, so they should be shared, reviewed and versioned like code, not copied from one laptop to the next. We run 45 Claude Code skills across a team of strategists, developers, writers and sales. A shared folder of skills stopped working long before we got to 45. So we built cc-skills , a small CLI that installs the team's skills from one GitHub repo into every coding agent on a machine, keeps them current with one command, and lets each person pick only the skills their job needs. This post covers what skills are, how they differ from plugins, where a shared folder broke for us, and what the CLI does that a plugin doesn't.
A skill is a folder with a SKILL.md file in it. The file starts with a short frontmatter block, a name and a description , followed by instructions in plain markdown. The folder can also hold scripts, reference documents and templates the instructions point to.
Claude Code reads only the name and description of every installed skill when a session starts. When your request matches a description, it loads the full SKILL.md and follows it. You can also call a skill directly by typing a slash and its name. That's why the description matters more than anything else in the file: it decides whether the skill ever gets used.
Skills live in ~/.claude/skills for you personally, or in .claude/skills inside a project. As of November 2026 the same format is an open standard, Agent Skills , so Codex , Cursor, GitHub Copilot, OpenCode and others read the same folders. A skill written once works in most coding agents, which turned out to matter a lot for how we distribute them.
A skill and a plugin aren't alternatives. A skill is one piece of know-how. A plugin is a package that installs several pieces at once, and skills are usually what's inside it.
A Claude Code plugin is a folder with a .claude-plugin/plugin.json manifest. It can hold skills, subagents, hooks, MCP servers, slash commands and executables for the agent's shell. You install it from a marketplace , which is a git repo that lists plugins, with /plugin install name@marketplace . Figma's plugin is a typical one: one install brings the Figma MCP server and around 15 skills.
Plugin skills are namespaced. A skill called review inside my-plugin runs as /my-plugin:review , while a skill in ~/.claude/skills keeps its plain name.
Plugins can be private. The marketplace repo can be a private GitHub repo , and Claude Code clones it with the git credentials already on your machine. A team can add the marketplace to a project's .claude/settings.json so teammates are offered it when they trust the folder, and on Team and Enterprise plans an admin can sync a private repo through the organization's GitHub connection. Marketplaces outside Anthropic's own don't auto-update by default, so each person turns that on or runs /plugin marketplace update .
Where each one works is what decided it for us. Other agents have their own plugin formats, but they all read plain skill folders. A Claude Code plugin only reaches the people on the team who work in Claude Code.
Where a shared folder of skills broke
Our skills started in March 2026 as a git repo with a bash install script. Clone it, run the script, and it symlinked everything into ~/.claude/skills . That worked for one person and about a dozen skills.
Then the team started writing their own. Strategists, developers and writers each added skills for their own work. Within a few months we had the problems every team with shared skills runs into:
Skills drifted. People edited a copy in their own ~/.claude/skills to fix something on a client job, and the fix never made it back. Two people would run "the same" skill and get different results.
Nobody knew what version they had. A change to how one skill records its work had to reach every developer before the next client job, and there was no way to check who had it.
Everyone got everything. A copywriter doesn't need a developer's skills in every session, and every installed description is something the agent has to read and choose between.
It only covered Claude Code. Not everyone on the team works in Claude Code, and Codex, Cursor and the rest each keep skills in their own folder.
One command installs the team's skills
cc-skills is a TypeScript CLI with no dependencies. Node 24 runs it directly, so there's no build step and nothing to install beyond cloning the repo. Running it opens one screen with four tabs: the agents it detected on your machine, packs, individual skills and tools.
It links skills rather than copying them, so the clone on your machine is the install. cc-skills update pulls the repo and relinks. New skills in your packs appear, moved skills get repointed, and retired skills are unlinked with the reason printed. That last part matters more than it sounds: when we retire a skill, nobody keeps running the old one by accident.
It links into every agent it finds: Claude Code, Codex, the universal ~/.agents/skills folder, OpenCode, GitHub Copilot and Cursor. A developer who switches from Claude Code to Codex for a task has the same skills there.
It also handles the drift problem directly. If you already have your own copy of a skill where the link should go, the CLI compares it with the repo version. An identical copy gets replaced by the link. A different one is shown as a conflict with the files that differ and is left alone unless you say yes, in which case it keeps a timestamped backup. And cc-skills import goes the other way: it finds skills you wrote locally, asks which pack each belongs to, copies them into the repo and swaps your local copy for a link.
A GitHub repo for versions, ideas and issues
Because every skill lives in one repo, a skill is reviewed like code. Changes go through a pull request with a short template, CI runs cc-skills check (directory names match skill names, links between skills resolve, the catalogue is current), and every release gets a changelog entry that says what changed for the person who installs it. Version 1.0.0 shipped in October 2026; we're on 1.6 now.
This is what fixed "which version do you have". When we changed one development skill so that two people working on the same client project stopped creating merge conflicts, the change reached every developer at their next cc-skills update . Nobody had to be told to copy anything.
It also gives everyone a place to put an idea or a complaint. Anyone can open an issue when a skill gets something wrong, or a PR when they've fixed it. So far 18 pull requests have come from four people, and not all of them are developers.
Retiring a skill is a git mv into an archive folder with a line explaining why. We archived 11 skills after finding none of them had been used in 31 days: a few were replaced by better ones, a few were built for one client, and two depended on a site that blocks automated access.
Packs: install only what your job needs
Not everyone uses every skill. Skills are grouped into packs by role, not by topic: one for strategists, one for developers, one for writers, and smaller ones for marketing, sales and our internal knowledge base. The strategist pack is the biggest, at 21 skills.
You tick the packs for your job and get all their skills. On top of that you can tick a single skill from another pack, or untick one inside yours, and the CLI remembers those as overrides. When we add a new pack, everyone is offered it once at their next update, and never asked again.
This keeps each person's agent focused. A strategist's session sees research skills, not front-end coding conventions. Fewer descriptions to choose from means the agent picks the right skill more often.
Tools for what MCP servers don't cover
The last piece came later. Official MCP servers are fine, but sometimes the one we need doesn't exist, or doesn't do what we need, so we built tools we can trust. A tool is code, not instructions: a command-line script with a short skill that tells the agent when to run it. Two examples:
sentry-issues lists errors across every client's Sentry organization with one token, worst project first. Sentry has its own MCP server , but we work across many client organizations and wanted one command that answers "what broke this week, everywhere".
rybbit pulls visitor analytics per client site from Rybbit . Adding up every site tells us how many real visitors a new browser target would cut off before we raise it.
Tools also save tokens. When an agent works through an MCP server, every step is a round trip through the model: call the server, read the JSON, decide the next call, reshape the output. For a task we repeat every week, like a Sentry triage across every client project, that's a lot of tokens spent on fetching and formatting. A script does the same work the same way every time and hands the agent a short table to reason about. The agent's job becomes deciding what matters, not paging through API responses.
cc-skills install --tool sentry,rybbit links each command onto your PATH and its skill into every agent. The CLI never asks for, reads or writes API keys. It only tells you which ones are missing.
What changed for the team and our clients
Setting up a new machine is cc-skills install --pack dev -y . A fix to a skill reaches everyone in one update, so a lesson from one client job carries over to the next. Skills people wrote on their own now land in the repo through cc-skills import instead of staying on one laptop.
For the stores we run CRO programs for, the difference is consistency. Every developer builds and rolls out an A/B test the same way, and every strategist runs a CRO audit the same way, because they run the same skills. When we change how we do something, we change the skill once.
When a plain folder or a plugin is enough
You don't need any of this if you're one person, or a small team that all use Claude Code and the same handful of skills. A folder in ~/.claude/skills , or .claude/skills committed to a project, is simpler and works. If your skills only matter inside one codebase, keep them in that codebase.
If your whole team works in Claude Code, start with a private plugin marketplace, one plugin per role. That gives you versioned installs, updates and bundles per role without writing anything. We'd have done the same if Claude Code were the only agent we used. What it doesn't give you is the other agents, single skills on top of a bundle, or a way to bring someone's local copy back into the repo.
A CLI earned its keep for us once three things were true: several people writing skills, people in different roles needing different sets, and more than one coding agent in use. The last one is what rules out a plugin. If your team uses more than one agent, put your skills in one repo before they drift, and write the install script before the folder gets past a dozen skills.
Where a shared folder of skills broke
One command installs the team's skills
A GitHub repo for versions, ideas and issues
Packs: install only what your job needs
Tools for what MCP servers don't cover
What changed for the team and our clients
When a plain folder or a plugin is enough
Where a shared folder of skills broke
One command installs the team's skills
A GitHub repo for versions, id

[truncated]
