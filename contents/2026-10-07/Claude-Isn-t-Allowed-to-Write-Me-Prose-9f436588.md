---
source: "https://blog.kvit.app/posts/agent-not-allowed-to-write-prose/"
hn_url: "https://news.ycombinator.com/item?id=49987880"
title: "Claude Isn't Allowed to Write Me Prose"
article_title: "Claude Isn't Allowed to Write Me Prose | Kvit Blog"
image: "https://blog.kvit.app/images/og-default.png"
author: "skolos"
captured_at: "2026-10-07T04:38:26Z"
capture_tool: "hn-digest"
hn_id: 49987880
score: 4
comments: 0
posted_at: "2026-10-07T03:45:41Z"
tags:
  - hacker-news
---

# Claude Isn't Allowed to Write Me Prose

- HN: [49987880](https://news.ycombinator.com/item?id=49987880)
- Source: [blog.kvit.app](https://blog.kvit.app/posts/agent-not-allowed-to-write-prose/)
- Score: 4
- Comments: 0
- Posted: 2026-10-07T03:45:41Z

## Translation

Title: Claude Isn't Allowed to Write Me Prose
Article title: Claude Isn't Allowed to Write Me Prose | Kvit Blog
Description: The agent reports through a fixed set of typed blocks instead of prose, and the UI draws them.

Article text:
">
Claude Isn't Allowed to Write Me Prose | Kvit Blog
Kvit Blog
Claude Isn't Allowed to Write Me Prose
The agent reports through a fixed set of typed blocks instead of prose, and the UI draws them.
My global CLAUDE.md has a section called Wording. It bans “honestly”, “genuine”, “landed”, and the phrase “earns its keep”. A later section bans false-contrast pairs such as “This is not theoretical. It is standard engineering.”, runs of short sentences written for punch, and one-line closers. Each rule is there because reading Claude’s prose day after day wore me down.
I run about a dozen agents at once, so whether I can read and process a report from an agent in seconds or need several minutes impacts my work speed. Until recently I treated that as a writing style problem: if the reports were still hard to read, style rules were not complete yet. The list kept growing, fastest after Opus 5 came out, when the prose became unbearable.
Opus 5.5 writes quite a bit better, but finding the parts of a report that matter is still a chore: the model tells me in detail which problems it ran into, how it tried to solve them and which paths it took, and most of that I do not need. Somewhere in the middle, often in passing, it mentions that it noticed something that might be a problem and might need a closer look, and that sentence is usually buried among details I could have skipped.
I am now convinced I was working on the wrong layer. Telling the model how to write, and what not to write, still lets it produce an essay of whatever length and shape it chooses; the rules only change the words inside it. And a rule in CLAUDE.md is only a request. What I need from a report is the two or three things I have to decide, and prose hides those even when it is well written.
What if, instead, the model could talk to me only through a small set of predefined blocks, each with fields it has to fill in?
To tell me what it changed, it uses a change block: a one-sentence summary, plus a longer description that stays folded until I open it.
To tell me something it learned, it uses a finding block, which also has to say how much the finding matters and whether it stops the work.
To tell me how it checked the work, it uses a check block: passed, failed or not run, and if not run, why not.
To ask me to decide something, it uses a question block with two to six options, each saying what picking it will do.
A harness checks every report before it reaches me and sends back anything that does not fit (as a tool call error), so the model never gets to hand me an essay. It still decides what to tell me, but everything it tells me has to go into one of these blocks.
Once the output is a list of typed blocks, the UI can present it the way I want to read it. Here are examples of what can be done:
UI doesn’t need to display block in order model sends them:
Click an option, or click the panel and press 1 or 2 .
The recommended option is marked
When model suggests options of what to do I often ask to tell me what it recommends and why. No need for round trip - the block could just require this info up front:
Long descriptions, logs and evidence stay hidden behind the one-line summaries until I open the ones I want:
Or here’s how I would structure it into single card:
Click a card, then press 1 to 4 to pick an option, or 0 to open or close the details.
For comparison, this is the text Claude Code actually ended that session with:
On the card I spend 10-15 seconds to check title, possible options and maybe glancing on “what’s happened section”. The prose from model I need to read and process every word - significantly more time and effort.
That moves the question of what an agent may show you from the model to the person running it. Today the model decides everything, in prose, and the interface prints whatever arrives. With a fixed set of blocks, the model still chooses which blocks to use and what goes in them, while I decide which blocks exist and what each has to contain.
It also makes the interaction more robust. The way a model writes its reports differs between models, and even between minor versions of the same model. With blocks, a new model changes what goes into the fields, while the layout I read and the buttons I press stay the same.
The same report from a working agent
To try the idea, I built it into kvit-coder , a small coding agent I work on. Here is the same result written as one of its reports and printed by its renderer:
What separates this from a style rule is that the report is a tool call and the only way the model can stop its work in kvit-coder .
Which blocks exist is a design choice, and this is the set kvit-coder uses today:
Changing the agent’s behaviour by changing the blocks
When a prose report goes wrong, the only remedy is another line in the prompt. When a block report goes wrong, the failure is usually a specific missing block or a field used for the wrong thing, and the fix goes to that place.
Nothing in the approach requires blocks to be text. A block type is a schema the model fills in plus code in the UI that draws it, so a UI that can draw pictures can offer the model blocks that are pictures:
A diagram block. The model supplies nodes and edges, or Mermaid source, and a caption, and the check can make sure the source parses before the report is accepted. After a refactor, this is where “the request now goes through these three services in this order” belongs.
A plot block. The model supplies the series, the axis labels and units, and the command that produced the numbers, and the check can make sure every series matches the axis. Benchmark results, or memory use before and after a change, would go here.
A step-through block. The model supplies an ordered list of states with a caption for each, and the UI plays them as an animation or steps through them one key at a time, which fits a race condition or a data migration.
Here is a step-through of the game bug from earlier in this post. The boxes, the arrows and what changes at each step all come from the data under “What the model sent”; the page only draws it.
In none of these does the model draw anything: it decides what should be shown and supplies the content, and UI code does the drawing the same way every time. Today an agent that wants to explain how three services fit together writes three paragraphs or an ASCII drawing, because text is all its interface can show. The figure above is a prototype: none of these blocks exist in kvit-coder yet.
The full version is in kvit-coder , under internal/report : report.go has the block types and the schema, validate.go the checks, and render.go the layout.
Part of it works in Claude Code through the What’s Next plugin , which uses a hook to require that a turn which changed something ends with Claude Code’s built-in question card.

## Original Extract

The agent reports through a fixed set of typed blocks instead of prose, and the UI draws them.

">
Claude Isn't Allowed to Write Me Prose | Kvit Blog
Kvit Blog
Claude Isn't Allowed to Write Me Prose
The agent reports through a fixed set of typed blocks instead of prose, and the UI draws them.
My global CLAUDE.md has a section called Wording. It bans “honestly”, “genuine”, “landed”, and the phrase “earns its keep”. A later section bans false-contrast pairs such as “This is not theoretical. It is standard engineering.”, runs of short sentences written for punch, and one-line closers. Each rule is there because reading Claude’s prose day after day wore me down.
I run about a dozen agents at once, so whether I can read and process a report from an agent in seconds or need several minutes impacts my work speed. Until recently I treated that as a writing style problem: if the reports were still hard to read, style rules were not complete yet. The list kept growing, fastest after Opus 5 came out, when the prose became unbearable.
Opus 5.5 writes quite a bit better, but finding the parts of a report that matter is still a chore: the model tells me in detail which problems it ran into, how it tried to solve them and which paths it took, and most of that I do not need. Somewhere in the middle, often in passing, it mentions that it noticed something that might be a problem and might need a closer look, and that sentence is usually buried among details I could have skipped.
I am now convinced I was working on the wrong layer. Telling the model how to write, and what not to write, still lets it produce an essay of whatever length and shape it chooses; the rules only change the words inside it. And a rule in CLAUDE.md is only a request. What I need from a report is the two or three things I have to decide, and prose hides those even when it is well written.
What if, instead, the model could talk to me only through a small set of predefined blocks, each with fields it has to fill in?
To tell me what it changed, it uses a change block: a one-sentence summary, plus a longer description that stays folded until I open it.
To tell me something it learned, it uses a finding block, which also has to say how much the finding matters and whether it stops the work.
To tell me how it checked the work, it uses a check block: passed, failed or not run, and if not run, why not.
To ask me to decide something, it uses a question block with two to six options, each saying what picking it will do.
A harness checks every report before it reaches me and sends back anything that does not fit (as a tool call error), so the model never gets to hand me an essay. It still decides what to tell me, but everything it tells me has to go into one of these blocks.
Once the output is a list of typed blocks, the UI can present it the way I want to read it. Here are examples of what can be done:
UI doesn’t need to display block in order model sends them:
Click an option, or click the panel and press 1 or 2 .
The recommended option is marked
When model suggests options of what to do I often ask to tell me what it recommends and why. No need for round trip - the block could just require this info up front:
Long descriptions, logs and evidence stay hidden behind the one-line summaries until I open the ones I want:
Or here’s how I would structure it into single card:
Click a card, then press 1 to 4 to pick an option, or 0 to open or close the details.
For comparison, this is the text Claude Code actually ended that session with:
On the card I spend 10-15 seconds to check title, possible options and maybe glancing on “what’s happened section”. The prose from model I need to read and process every word - significantly more time and effort.
That moves the question of what an agent may show you from the model to the person running it. Today the model decides everything, in prose, and the interface prints whatever arrives. With a fixed set of blocks, the model still chooses which blocks to use and what goes in them, while I decide which blocks exist and what each has to contain.
It also makes the interaction more robust. The way a model writes its reports differs between models, and even between minor versions of the same model. With blocks, a new model changes what goes into the fields, while the layout I read and the buttons I press stay the same.
The same report from a working agent
To try the idea, I built it into kvit-coder , a small coding agent I work on. Here is the same result written as one of its reports and printed by its renderer:
What separates this from a style rule is that the report is a tool call and the only way the model can stop its work in kvit-coder .
Which blocks exist is a design choice, and this is the set kvit-coder uses today:
Changing the agent’s behaviour by changing the blocks
When a prose report goes wrong, the only remedy is another line in the prompt. When a block report goes wrong, the failure is usually a specific missing block or a field used for the wrong thing, and the fix goes to that place.
Nothing in the approach requires blocks to be text. A block type is a schema the model fills in plus code in the UI that draws it, so a UI that can draw pictures can offer the model blocks that are pictures:
A diagram block. The model supplies nodes and edges, or Mermaid source, and a caption, and the check can make sure the source parses before the report is accepted. After a refactor, this is where “the request now goes through these three services in this order” belongs.
A plot block. The model supplies the series, the axis labels and units, and the command that produced the numbers, and the check can make sure every series matches the axis. Benchmark results, or memory use before and after a change, would go here.
A step-through block. The model supplies an ordered list of states with a caption for each, and the UI plays them as an animation or steps through them one key at a time, which fits a race condition or a data migration.
Here is a step-through of the game bug from earlier in this post. The boxes, the arrows and what changes at each step all come from the data under “What the model sent”; the page only draws it.
In none of these does the model draw anything: it decides what should be shown and supplies the content, and UI code does the drawing the same way every time. Today an agent that wants to explain how three services fit together writes three paragraphs or an ASCII drawing, because text is all its interface can show. The figure above is a prototype: none of these blocks exist in kvit-coder yet.
The full version is in kvit-coder , under internal/report : report.go has the block types and the schema, validate.go the checks, and render.go the layout.
Part of it works in Claude Code through the What’s Next plugin , which uses a hook to require that a turn which changed something ends with Claude Code’s built-in question card.
