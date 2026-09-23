---
source: "https://blog.master.dev/introducing-solo/"
hn_url: "https://news.ycombinator.com/item?id=49820421"
title: "Discover Solo: The Ultimate AI-Integrated Control Panel"
article_title: "Discover Solo: The Ultimate AI-Integrated Control Panel"
image: "https://blog.master.dev/wp-json/social-image-generator/v1/image/11114"
author: "ibobev"
captured_at: "2026-09-23T19:29:32Z"
capture_tool: "hn-digest"
hn_id: 49820421
score: 1
comments: 0
posted_at: "2026-09-23T18:30:33Z"
tags:
  - hacker-news
---

# Discover Solo: The Ultimate AI-Integrated Control Panel

- HN: [49820421](https://news.ycombinator.com/item?id=49820421)
- Source: [blog.master.dev](https://blog.master.dev/introducing-solo/)
- Score: 1
- Comments: 0
- Posted: 2026-09-23T18:30:33Z

## Translation

Title: Discover Solo: The Ultimate AI-Integrated Control Panel
Description: Discover Solo: the ultimate tool for streamlining coding projects with AI integration, terminal management, and powerful debugging features.

Article text:
$100 off Fall Sale
Master the Fundamentals. Build Better With AI.
Get Discount Now →
← Back to Master.dev
Courses
Learn
Become a Member
Guest Writing
RSS
Blog
AI CLI Solo TanStack
Introducing Solo
Solo is one of my favorite new tools. I heard about it recently from its creator, Aaron Francis, while at a conference he was emceeing.
Solo’s website describes it as a Meta-harness for coding agents, but I don’t think that does the project justice. To me, Solo is a control panel for whatever project I’m working on, with deep AI integration (naturally). It’s home to any terminals I might need, with common commands preloaded into dedicated slots in Solo, with the option to auto-start. And again, AI is deeply integrated: you can launch agents, and what’s especially neat is that Solo provides its own MCP server that gives your agents access to your tasks and terminals to help debug problems you’re having.
I won’t walk you through installation or adding a project. It’ll make some best guesses on which commands you’ll likely want. Tweak as desired (you can always adjust later) and create it.
Here’s what mine looks like for a fitness tracking application I’ve been messing with.
This is a TanStack Start web application , and as you can see, I’ve got two commands (basically a terminal embedded in Solo) running: my dev web server and Postgres via Docker.
TanStack Start & TanStack Query
As you hover over those commands, you’d see buttons to stop, start, or restart any of these commands. Naturally, you can click any of these commands and see that terminal’s output in the main Solo window. In the image above, you can see my dev server.
And obviously you can edit these commands anytime.
So far, all I’ve shown is an app that manages multiple terminals in one convenient place, with common commands pre-set.
But it’s 2026, so obviously you want to see the AI integration. Naturally you can start agents in Solo; there’s even a dedicated section for it, which you can see in the image above. There’s no shortage of shortcuts and UI commands for this, but just hit Command-T and type “agent”
Don’t worry, Solo supports virtually any agent you’ve ever heard of; only Claude and Codex show up here because that’s all I bothered to set up.
Once you start an agent, it’s living as normal right inside Solo, just like you’re used to.
Nothing changes for you as the user.
I moved kind of fast above because I wanted to get to the more interesting AI pieces. I mentioned earlier that Solo has a built-in MCP server that lets your agents inspect (among other things) your other commands and their outputs to help debug problems.
Let’s try it out. I’ll stop my database process.
Obviously nothing will work now.
As a control, before I touch Solo’s MCP, let’s make sure this isn’t something a vanilla Claude agent could easily debug. When I ask it why I have errors in my dev server, it starts taking steps to start the dev server and reproduce.
I’d prefer it to just look at my existing output, along with neighboring commands.
The MCP section in Solo has a dirt-simple way to enable it for whatever agent you use. It gives you a bash command, or if that’s too much effort, a nice fat “Run” button that executes that command for you.
Find your harness of choice and smash that Run button, and it should show as installed.
When we try to debug the same problem with the same prompt, unfortunately nothing really changes.
If you want to engage Solo’s MCP, you need to be a bit more specific in your prompt
Like any software engineer, I tend to be lazy and would prefer to not have to manually tell it “hey, use Solo’s MCP to blah blah” every time I want it to. A skill is a nice way to wrap that bit of functionality up.
At the time of writing, Skills have sort of gotten a bad name, with devs dumping way too many of them in their repo, flooding context windows, potentially affecting skill selection quality, etc. But a skill that can only be manually invoked can avoid those issues and essentially serve as a subroutine for wrapping common functionality (like any function we’re used to writing).
Here’s the skill I whipped up:
---
name: solo-debug
description: Debug this application using Solo MCP
disable-model-invocation: true
---
# Solo Debug
Use the Solo MCP tools to inspect the processes already running
for this project.
1. Inspect recent output and errors.
2. Use those process outputs to help diagnose what's being asked.
3. Do not start new processes unless necessary. Code language: Markdown ( markdown )
Note this line:
disable-model-invocation: true Code language: YAML ( yaml )
That prevents models from invoking it on their own.
I put that in .claude/skills/solo-debug/SKILL.md and with that, I can now just do /solo-debug and type my original prompt
I’m about to wrap this post up, but if you’re feeling underwhelmed with Solo, I promise I’m barely scratching the surface. Solo also supports scratchpads and to-dos, which, of course, can integrate with your agents. And there are entire AI orchestration workflows . They even have guides on building better daily workflows .
Solo is a superb tool for managing multiple processes and agents and connecting everything seamlessly. The process management alone is nice, but the built-in MCP support for improved debugging really makes this tool a favorite of mine. And that is before we’ve even scratched the surface of the deeper AI workflows it supports.
Learn to build production-ready apps with TanStack from Adam Rackis. Explore the full ecosystem, from Router to Start to Query, while leveraging server-side rendering, type-safe routing, caching strategies, and React Server Components. Walk away with the skills to ship high-performance full-stack React apps. Access 300+ courses with a Master.dev subscription and get 20% off today!
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
Notify me of follow-up comments by email.
Notify me of new posts by email.
Solo
TanStack Start & TanStack Query
Using Solo’s MCP Adding a Skill
Our courses go beyond frontend into fullstack, devops, and AI.
Master.dev donates to open source projects through thanks.dev and Open Collective , as well as donates to non-profits like The Last Mile , Annie Canons , and Vets Who Code .

## Original Extract

Discover Solo: the ultimate tool for streamlining coding projects with AI integration, terminal management, and powerful debugging features.

$100 off Fall Sale
Master the Fundamentals. Build Better With AI.
Get Discount Now →
← Back to Master.dev
Courses
Learn
Become a Member
Guest Writing
RSS
Blog
AI CLI Solo TanStack
Introducing Solo
Solo is one of my favorite new tools. I heard about it recently from its creator, Aaron Francis, while at a conference he was emceeing.
Solo’s website describes it as a Meta-harness for coding agents, but I don’t think that does the project justice. To me, Solo is a control panel for whatever project I’m working on, with deep AI integration (naturally). It’s home to any terminals I might need, with common commands preloaded into dedicated slots in Solo, with the option to auto-start. And again, AI is deeply integrated: you can launch agents, and what’s especially neat is that Solo provides its own MCP server that gives your agents access to your tasks and terminals to help debug problems you’re having.
I won’t walk you through installation or adding a project. It’ll make some best guesses on which commands you’ll likely want. Tweak as desired (you can always adjust later) and create it.
Here’s what mine looks like for a fitness tracking application I’ve been messing with.
This is a TanStack Start web application , and as you can see, I’ve got two commands (basically a terminal embedded in Solo) running: my dev web server and Postgres via Docker.
TanStack Start & TanStack Query
As you hover over those commands, you’d see buttons to stop, start, or restart any of these commands. Naturally, you can click any of these commands and see that terminal’s output in the main Solo window. In the image above, you can see my dev server.
And obviously you can edit these commands anytime.
So far, all I’ve shown is an app that manages multiple terminals in one convenient place, with common commands pre-set.
But it’s 2026, so obviously you want to see the AI integration. Naturally you can start agents in Solo; there’s even a dedicated section for it, which you can see in the image above. There’s no shortage of shortcuts and UI commands for this, but just hit Command-T and type “agent”
Don’t worry, Solo supports virtually any agent you’ve ever heard of; only Claude and Codex show up here because that’s all I bothered to set up.
Once you start an agent, it’s living as normal right inside Solo, just like you’re used to.
Nothing changes for you as the user.
I moved kind of fast above because I wanted to get to the more interesting AI pieces. I mentioned earlier that Solo has a built-in MCP server that lets your agents inspect (among other things) your other commands and their outputs to help debug problems.
Let’s try it out. I’ll stop my database process.
Obviously nothing will work now.
As a control, before I touch Solo’s MCP, let’s make sure this isn’t something a vanilla Claude agent could easily debug. When I ask it why I have errors in my dev server, it starts taking steps to start the dev server and reproduce.
I’d prefer it to just look at my existing output, along with neighboring commands.
The MCP section in Solo has a dirt-simple way to enable it for whatever agent you use. It gives you a bash command, or if that’s too much effort, a nice fat “Run” button that executes that command for you.
Find your harness of choice and smash that Run button, and it should show as installed.
When we try to debug the same problem with the same prompt, unfortunately nothing really changes.
If you want to engage Solo’s MCP, you need to be a bit more specific in your prompt
Like any software engineer, I tend to be lazy and would prefer to not have to manually tell it “hey, use Solo’s MCP to blah blah” every time I want it to. A skill is a nice way to wrap that bit of functionality up.
At the time of writing, Skills have sort of gotten a bad name, with devs dumping way too many of them in their repo, flooding context windows, potentially affecting skill selection quality, etc. But a skill that can only be manually invoked can avoid those issues and essentially serve as a subroutine for wrapping common functionality (like any function we’re used to writing).
Here’s the skill I whipped up:
---
name: solo-debug
description: Debug this application using Solo MCP
disable-model-invocation: true
---
# Solo Debug
Use the Solo MCP tools to inspect the processes already running
for this project.
1. Inspect recent output and errors.
2. Use those process outputs to help diagnose what's being asked.
3. Do not start new processes unless necessary. Code language: Markdown ( markdown )
Note this line:
disable-model-invocation: true Code language: YAML ( yaml )
That prevents models from invoking it on their own.
I put that in .claude/skills/solo-debug/SKILL.md and with that, I can now just do /solo-debug and type my original prompt
I’m about to wrap this post up, but if you’re feeling underwhelmed with Solo, I promise I’m barely scratching the surface. Solo also supports scratchpads and to-dos, which, of course, can integrate with your agents. And there are entire AI orchestration workflows . They even have guides on building better daily workflows .
Solo is a superb tool for managing multiple processes and agents and connecting everything seamlessly. The process management alone is nice, but the built-in MCP support for improved debugging really makes this tool a favorite of mine. And that is before we’ve even scratched the surface of the deeper AI workflows it supports.
Learn to build production-ready apps with TanStack from Adam Rackis. Explore the full ecosystem, from Router to Start to Query, while leveraging server-side rendering, type-safe routing, caching strategies, and React Server Components. Walk away with the skills to ship high-performance full-stack React apps. Access 300+ courses with a Master.dev subscription and get 20% off today!
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
Notify me of follow-up comments by email.
Notify me of new posts by email.
Solo
TanStack Start & TanStack Query
Using Solo’s MCP Adding a Skill
Our courses go beyond frontend into fullstack, devops, and AI.
Master.dev donates to open source projects through thanks.dev and Open Collective , as well as donates to non-profits like The Last Mile , Annie Canons , and Vets Who Code .
