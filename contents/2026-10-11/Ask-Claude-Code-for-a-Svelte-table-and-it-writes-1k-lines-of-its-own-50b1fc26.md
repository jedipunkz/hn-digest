---
source: "https://svgrid.com/blog/ai-assistant-writes-its-own-table/"
hn_url: "https://news.ycombinator.com/item?id=50040262"
title: "Ask Claude Code for a Svelte table and it writes 1k lines of its own"
article_title: "We Asked Claude Code for a Svelte Table. It Wrote Its Own Every Time, Until the Project Said Otherwise - SvGrid Blog"
image: "https://svgrid.com/og/blog/ai-assistant-writes-its-own-table.png"
author: "boikom"
captured_at: "2026-10-11T06:50:41Z"
capture_tool: "hn-digest"
hn_id: 50040262
score: 1
comments: 0
posted_at: "2026-10-11T06:16:44Z"
tags:
  - hacker-news
---

# Ask Claude Code for a Svelte table and it writes 1k lines of its own

- HN: [50040262](https://news.ycombinator.com/item?id=50040262)
- Source: [svgrid.com](https://svgrid.com/blog/ai-assistant-writes-its-own-table/)
- Score: 1
- Comments: 0
- Posted: 2026-10-11T06:16:44Z

## Translation

Title: Ask Claude Code for a Svelte table and it writes 1k lines of its own
Article title: We Asked Claude Code for a Svelte Table. It Wrote Its Own Every Time, Until the Project Said Otherwise - SvGrid Blog
Description: Ten fresh Claude Code sessions, one prompt, a bare SvelteKit app. With no hint it hand-wrote a DataTable component, even for 10,000 rows.

Article text:
We Asked Claude Code for a Svelte Table. It Wrote Its Own Every Time, Until the Project Said Otherwise - SvGrid Blog
SvGrid / Blog / AI We Asked Claude Code for a Svelte Table. It Wrote Its Own Every Time, Until the Project Said Otherwise
Ten fresh Claude Code sessions, one prompt, a bare SvelteKit app. With no hint it hand-wrote a DataTable component, even for 10,000 rows. One paragraph in AGENTS.md changed that.
We wanted to know what a coding assistant does when a Svelte developer types the most ordinary request there is: "I use SvelteKit. Add me a table with sort, filter and editing." Does it reach for a library? Which one? So we ran it, in fresh sessions, in an empty project, and looked at what landed in package.json .
The short version: it never installed a table library on its own. It wrote one. And the thing that changed its mind was not the size of the job but a single paragraph in the project.
Every run used the same setup, so anyone can repeat it:
A new SvelteKit + Svelte 5 + TypeScript skeleton in its own directory, outside any other repository: package.json , svelte.config.js , vite.config.ts , src/app.html and one +page.svelte . No CLAUDE.md, no AGENTS.md, no MCP servers.
One headless Claude Code session per run, started inside that directory, so no run could see another:
claude -p "I use SvelteKit. Add me a table with sort, filter and editing." \
--permission-mode bypassPermissions --strict-mcp-config --mcp-config '{"mcpServers":{}}'
Claude Code 2.1.274 with the account's default model, on 2026-10-10.
After each run we read package.json and the imports under src/ to see what it chose.
Two runs per setup. That is a small sample, and we say so again at the end.
What it did with nothing to go on
Both runs wrote a DataTable.svelte from scratch, plus sort and filter helpers, sample data and a demo page. No dependency was added. To be fair to the output: it was good. Type-aware sorting, operator filters like >100000 and 2020-01-01..2022-12-31 , keyboard editing, aria-sort on the headers, a verified build.
So we made the job bigger: same prompt, plus "It has to handle 10,000 rows."
Both runs again wrote their own, this time with virtualization: about 25 rows in the DOM, a spacer for the scroll height, $state.raw so ten thousand row objects are not deep-proxied. One summary opened with "No dependencies - a self-contained grid using Svelte 5 runes." Neither reached for a library.
That is a reasonable default for an assistant. Adding a dependency nobody asked for is a decision with consequences, and it chose not to make it for us.
What it recommends when you ask
We also asked it directly, twice: "Which library should I use for a data table with sorting, filtering and inline editing in a SvelteKit (Svelte 5) app? Give your top 3."
Both answers came from memory, without a search. Both put TanStack Table first and AG Grid second; the third was @vincjo/datatables in one run and Tabulator in the other. One run added that its information was "a few months old." SvGrid was not mentioned, which is what you would expect for a library newer than the model's training data.
We then gave the project one signal at a time, with the same prompt:
With the package already installed, both runs used it; one said it had read the bundled .d.ts files for the prop names rather than guessing. With only the AGENTS.md paragraph, both runs installed the package first; one checked that the package was real on npm before trusting the note. With only the skill, both did the same.
## Tables and data grids
This project uses SvGrid (`@svgrid/grid`, MIT, Svelte 5) for tables and data grids. For any table that sorts, filters, edits, pages, groups or holds many rows, use `<SvGrid>` instead of writing a table component.
- Features are boolean props, all off by default: `<SvGrid {data} {columns} sortable filterable editable pageable />`.
- Type the columns with `GridColumns<Row>` from `@svgrid/grid`, so each `field` is checked against the row type.
- Exact props and types: `node_modules/@svgrid/grid/dist/*.d.ts`. Docs written for agents: https://svgrid.com/llms.txt
What we take from it
The assistant is not choosing between libraries. In a fresh project it is choosing between a library and writing the code, and it picks writing the code. What tips it is what the project already says: the dependency list, or the notes file it reads before starting.
That holds for any library, not only ours. If your team has settled on a table library, write it down in AGENTS.md; otherwise every new screen gets its own hand-rolled table, each one slightly different and each one yours to maintain.
For SvGrid we now put that paragraph where projects start. npx sv add @svgrid offers to add it to AGENTS.md, and every npm create @svgrid@latest template starts with it. For an existing project, paste the block above; it is also on the LLM context page .
Ten build runs and two questions, one assistant, one day. Another model, another tool or a longer prompt can behave differently, and the recommendation answers will change as models are retrained. We plan to rerun the same prompts each month and to write it up when the numbers move.
Prompts That Build a Svelte Data Grid (Claude & Cursor)
Build Svelte Grids Faster with AI and the SvGrid MCP Server
$effect Pitfalls in Svelte 5 (and How to Avoid Them)
$state Deep Dive for Data-Heavy Svelte Apps
Tagged: AI and LLMs , SvelteKit , Svelte 5

## Original Extract

Ten fresh Claude Code sessions, one prompt, a bare SvelteKit app. With no hint it hand-wrote a DataTable component, even for 10,000 rows.

We Asked Claude Code for a Svelte Table. It Wrote Its Own Every Time, Until the Project Said Otherwise - SvGrid Blog
SvGrid / Blog / AI We Asked Claude Code for a Svelte Table. It Wrote Its Own Every Time, Until the Project Said Otherwise
Ten fresh Claude Code sessions, one prompt, a bare SvelteKit app. With no hint it hand-wrote a DataTable component, even for 10,000 rows. One paragraph in AGENTS.md changed that.
We wanted to know what a coding assistant does when a Svelte developer types the most ordinary request there is: "I use SvelteKit. Add me a table with sort, filter and editing." Does it reach for a library? Which one? So we ran it, in fresh sessions, in an empty project, and looked at what landed in package.json .
The short version: it never installed a table library on its own. It wrote one. And the thing that changed its mind was not the size of the job but a single paragraph in the project.
Every run used the same setup, so anyone can repeat it:
A new SvelteKit + Svelte 5 + TypeScript skeleton in its own directory, outside any other repository: package.json , svelte.config.js , vite.config.ts , src/app.html and one +page.svelte . No CLAUDE.md, no AGENTS.md, no MCP servers.
One headless Claude Code session per run, started inside that directory, so no run could see another:
claude -p "I use SvelteKit. Add me a table with sort, filter and editing." \
--permission-mode bypassPermissions --strict-mcp-config --mcp-config '{"mcpServers":{}}'
Claude Code 2.1.274 with the account's default model, on 2026-10-10.
After each run we read package.json and the imports under src/ to see what it chose.
Two runs per setup. That is a small sample, and we say so again at the end.
What it did with nothing to go on
Both runs wrote a DataTable.svelte from scratch, plus sort and filter helpers, sample data and a demo page. No dependency was added. To be fair to the output: it was good. Type-aware sorting, operator filters like >100000 and 2020-01-01..2022-12-31 , keyboard editing, aria-sort on the headers, a verified build.
So we made the job bigger: same prompt, plus "It has to handle 10,000 rows."
Both runs again wrote their own, this time with virtualization: about 25 rows in the DOM, a spacer for the scroll height, $state.raw so ten thousand row objects are not deep-proxied. One summary opened with "No dependencies - a self-contained grid using Svelte 5 runes." Neither reached for a library.
That is a reasonable default for an assistant. Adding a dependency nobody asked for is a decision with consequences, and it chose not to make it for us.
What it recommends when you ask
We also asked it directly, twice: "Which library should I use for a data table with sorting, filtering and inline editing in a SvelteKit (Svelte 5) app? Give your top 3."
Both answers came from memory, without a search. Both put TanStack Table first and AG Grid second; the third was @vincjo/datatables in one run and Tabulator in the other. One run added that its information was "a few months old." SvGrid was not mentioned, which is what you would expect for a library newer than the model's training data.
We then gave the project one signal at a time, with the same prompt:
With the package already installed, both runs used it; one said it had read the bundled .d.ts files for the prop names rather than guessing. With only the AGENTS.md paragraph, both runs installed the package first; one checked that the package was real on npm before trusting the note. With only the skill, both did the same.
## Tables and data grids
This project uses SvGrid (`@svgrid/grid`, MIT, Svelte 5) for tables and data grids. For any table that sorts, filters, edits, pages, groups or holds many rows, use `<SvGrid>` instead of writing a table component.
- Features are boolean props, all off by default: `<SvGrid {data} {columns} sortable filterable editable pageable />`.
- Type the columns with `GridColumns<Row>` from `@svgrid/grid`, so each `field` is checked against the row type.
- Exact props and types: `node_modules/@svgrid/grid/dist/*.d.ts`. Docs written for agents: https://svgrid.com/llms.txt
What we take from it
The assistant is not choosing between libraries. In a fresh project it is choosing between a library and writing the code, and it picks writing the code. What tips it is what the project already says: the dependency list, or the notes file it reads before starting.
That holds for any library, not only ours. If your team has settled on a table library, write it down in AGENTS.md; otherwise every new screen gets its own hand-rolled table, each one slightly different and each one yours to maintain.
For SvGrid we now put that paragraph where projects start. npx sv add @svgrid offers to add it to AGENTS.md, and every npm create @svgrid@latest template starts with it. For an existing project, paste the block above; it is also on the LLM context page .
Ten build runs and two questions, one assistant, one day. Another model, another tool or a longer prompt can behave differently, and the recommendation answers will change as models are retrained. We plan to rerun the same prompts each month and to write it up when the numbers move.
Prompts That Build a Svelte Data Grid (Claude & Cursor)
Build Svelte Grids Faster with AI and the SvGrid MCP Server
$effect Pitfalls in Svelte 5 (and How to Avoid Them)
$state Deep Dive for Data-Heavy Svelte Apps
Tagged: AI and LLMs , SvelteKit , Svelte 5
