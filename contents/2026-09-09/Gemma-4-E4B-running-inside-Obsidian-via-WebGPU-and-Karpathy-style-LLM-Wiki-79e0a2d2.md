---
source: "https://github.com/itsyuimorii/obsidian-gemma4-litert-wiki"
hn_url: "https://news.ycombinator.com/item?id=49625192"
title: "Gemma 4 E4B running inside Obsidian via WebGPU and Karpathy-style LLM Wiki"
article_title: "GitHub - itsyuimorii/obsidian-gemma4-litert-wiki: Run Gemma 4 E4B locally inside Obsidian via LiteRT-LM and WebGPU — private, offline, no API key, no Ollama, no server. Builds Andrej Karpathy's LLM wiki over your vault: cards, index, concept pages, tags, links, provenance checks — then chat with a n\n[truncated]"
image: "https://opengraph.githubassets.com/c58e615a7d601fdd86e771abe54d22007368f70f5a44cbd9932e4d3c8594791d/itsyuimorii/obsidian-gemma4-litert-wiki"
author: "itsyuimorii"
captured_at: "2026-09-09T12:46:57Z"
capture_tool: "hn-digest"
hn_id: 49625192
score: 1
comments: 1
posted_at: "2026-09-09T12:09:52Z"
tags:
  - hacker-news
---

# Gemma 4 E4B running inside Obsidian via WebGPU and Karpathy-style LLM Wiki

- HN: [49625192](https://news.ycombinator.com/item?id=49625192)
- Source: [github.com](https://github.com/itsyuimorii/obsidian-gemma4-litert-wiki)
- Score: 1
- Comments: 1
- Posted: 2026-09-09T12:09:52Z

## Translation

Title: Gemma 4 E4B running inside Obsidian via WebGPU and Karpathy-style LLM Wiki
Article title: GitHub - itsyuimorii/obsidian-gemma4-litert-wiki: Run Gemma 4 E4B locally inside Obsidian via LiteRT-LM and WebGPU — private, offline, no API key, no Ollama, no server. Builds Andrej Karpathy's LLM wiki over your vault: cards, index, concept pages, tags, links, provenance checks — then chat with a n
[truncated]
Description: Run Gemma 4 E4B locally inside Obsidian via LiteRT-LM and WebGPU — private, offline, no API key, no Ollama, no server. Builds Andrej Karpathy's LLM wiki over your vault: cards, index, concept pages, tags, links, provenance checks — then chat with a note or the whole wiki, plus quizzes and flashcards
[truncated]

Article text:
GitHub - itsyuimorii/obsidian-gemma4-litert-wiki: Run Gemma 4 E4B locally inside Obsidian via LiteRT-LM and WebGPU — private, offline, no API key, no Ollama, no server. Builds Andrej Karpathy's LLM wiki over your vault: cards, index, concept pages, tags, links, provenance checks — then chat with a note or the whole wiki, plus quizzes and flashcards. · GitHub
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
itsyuimorii
/
obsidian-gemma4-litert-wiki
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
339 Commits 339 Commits Folders and files
.github .github assets assets docs docs scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.ja.md README.ja.md README.md README.md SECURITY.md SECURITY.md build.js build.js eslint.config.mjs eslint.config.mjs manifest.json manifest.json package-lock.json package-lock.json package.json package.json styles.css styles.css tsconfig.json tsconfig.json versions.json versions.json View all files Repository files navigation
Turn your notes into a living Karpathy LLM wiki.
Chat with your vault, connect ideas, find gaps, and generate quizzes & flashcards —
with Gemma 4 E4B running locally inside Obsidian. No API key, Ollama, or server.
⬇ Install from the community store
▶ See what it does · Step through the demo — nothing to install.
☔️ What do you do with everything you have already saved?
Clipped articles. Saved posts. Course notes. Papers you meant to read. A vault that looks organised and is mostly unread — not because there is too much in it, but because there is no easy way to see what you've actually collected without opening things one at a time.
Gemma 4 E4B builds Andrej Karpathy's LLM wiki out of the notes you already have.
🧠 Chat with one note or your whole vault — grounded in what you actually wrote, with sources listed by the plugin rather than generated by the model.
📇 A summary card for every note, and an index over all of them — so you can see what you have without opening anything.
🔗 Links between notes that turn out to be related — and tags drawn from one shared vocabulary rather than whatever each note happened to invent, which is what lets them find each other at all.
📚 Concept pages over the clusters that emerge — written above your notes, linking down into each one.
🔍 Gaps and contradictions surfaced — including claims in two notes that cannot both be true.
🎓 Quizzes and flashcards from any note — so the archive is something you revisit, not just something you kept.
💬 Or ask it anything at all — Direct mode drops the grounding and answers from the model itself, offline and free. It is a 4B model, so it is weaker at general knowledge than anything behind an API, and it says so by carrying no sources and never filing its answers into your wiki.
Everything it writes is plain Markdown, in your vault.
Your own notes are never modified by the model. The wiki is a separate layer built above them, and your notes stay exactly as you wrote them.
And the model is already inside Obsidian. Gemma 4 E4B runs in Obsidian's own process through LiteRT-LM and WebGPU — no API key, no Ollama, no LM Studio, no localhost server, no account, no subscription .
There is no provider to pick. Your notes are never uploaded anywhere, because there is no server to upload them to. Privacy is a property of the architecture, not a promise in a policy.
After the one-time downloads — the ~3 GB model and the WASM runtime, both under Privacy — inference runs locally and the network is no longer required.
💬 Chat with your notes — entirely offline
🔌 How this differs from other AI plugins
🔧 How it works
Manual test protocols
Sixteen commands, and this is what they add up to. Every row is covered in
detail below; the full command list is under
Current commands .
Two of the sixteen write into a note you wrote, both behind a preview. The
rest build and maintain a layer beside your notes and never touch them.
It implements Andrej Karpathy's LLM-wiki pattern : your raw notes are never modified by the model — the immutability constraint binds the LLM, not you, and editing your own notes is the normal fix when a chat surfaces a mistake (the wiki detects the change and offers to re-ingest). A separate, cross-linked wiki layer is built above them — one card per note, and every page is shown to you in full before a single byte is written .
From there: chat grounded in one note or in the whole wiki, quiz yourself on either, build concept pages across everything sharing a tag, and let it flag its own decay with lint, provenance and contradiction checks. Sources are listed by the plugin, never cited by the model — citation is the one thing a small local model would get wrong without anyone noticing.
Everything it writes is plain markdown in your vault — nothing is locked in a database, and nothing needs another plugin to read it back. Its own configuration is notes too: your tag vocabulary and naming rules live in schema.md , where you can edit them by hand and the plugin will obey; every operation is appended to log.md , so you can always see what it did and when; and dropping a markdown file into skills/ adds a command of your own to the ⚡ menu.
Status: 1.0.13, in the community plugin store . The full Karpathy loop is implemented and running; benchmarks below are from real use.
The Karpathy loop — raw notes stay read-only; the plugin maintains a separate gemma-wiki/ layer:
Ingest — one strict JSON extraction per note (summary, tags, key points, salient mentions, and a self-rated high / med / low confidence written to frontmatter), plus a validated multiple-choice pick of related pages from the index catalog. Everything previews in a review modal; nothing is written without approval. Ingested notes get a badge in the file explorer — pure UI decoration, the note file is untouched.
Scan (semi-automatic ingest) — sweep the folders you named for new or changed notes, draft a card for each, and review them all in one list sorted low-confidence first , so the pages most in need of a human eye are the ones you read first. Untick anything that should not have become a page. Opt-in by scope: leave the folder list blank and scan refuses to run rather than sweeping your vault.
index.md / log.md — a one-line-per-page catalog the query path reads first, and an append-only, grep-friendly activity log.
Query (Wiki mode) — index-first retrieval with stopword filtering; the catalog and recent log always ride along, so meta-questions ("what did I add today?") work; answers are grounded only in retrieved material with honest refusals otherwise.
Whole-wiki questions ground in every page , not in the ones that lexically match. "What connects my pages" is about the shape of the collection, so scoring it against page summaries matches nothing — the words in the question appear in no summary — and the model is handed zero pages and correctly reports it was given none. The two chips that ask this kind of question carry the flag; loadPages fills to the budget and stops.
Save an answer as a note — every reply has a save action (same review gate), and it writes the answer into your own notes , beside the note it came from, named gemma — Summarize — The perfect espresso shot so the file explorer shows at a glance both that a model wrote it and what it was about, and recording in frontmatter which model and from which sources — the two things a copy-paste destroys silently. It lands in your folders, under your organisation, and it survives deleting the wiki folder, which nothing the plugin owns is meant to. Being an ordinary note of yours, it is also the one thing a saved answer never used to be: ingestable . Scan it and it becomes a card like any other note — which is the only route by which anything becomes material, because everything the wiki retrieves derives from a note you wrote.
Everything retrievable derives from a note you wrote. That holds structurally rather than by convention: only cards/ (one card per ingested note) and concepts/ (built across those cards) are in the retrieval set.
Saved answers are never retrieved. Karpathy's gist says explorations should compound and says nothing about trust, which is the poisoning loop — a wrong or stale answer re-entering context with the same weight as the notes. So a saved answer is not in the retrieval set at all: it is an ordinary note of yours, kept to re-read, with frontmatter recording which model wrote it and from which sources. The one route back in is deliberate and human — judge it, and ingest it like anything else you wrote.
Concept pages — pick a tag or mention two or more pages share, and the plugin writes a page above them that links down into each one. Ingest ripples into them, and member lists self-heal in both directions.
Keeping it honest — a wiki nobody reviews is the failure mode this is built against:
Review board — three ways a page goes bad, in one queue: low self-rated confidence, source drift (the raw note changed since ingest, caught by source_hash ), and staleness. Concept overviews are checked too.
Provenance spot-check — takes a page's key points back to the raw note to find the sentence each came from, and flags the ones that cannot be traced.
Contradiction sweep — checks pairs of pages that share a tag for claims that disagree, recently-changed pairs first. It flags with the model's reason quoted and never edits — you decide which one was wrong.
Tidy the wiki — one command that checks first and then fixes only what you tick. The check is model-free: orphan pages, index entries pointing at deleted pages, pages missing from the index, and how far your pages' tags have drifted from the vocabulary. The four repairs — make links mutual, drop dead index entries, rebuild the vocabulary, apply it to existing pages — are always offered, ticked where the check found something; each runs in turn behind its own preview. Availability is never tied to detection, because a check good enough to pick a sensible default is not good enough to be the only way in.
Background count (off by default) — periodically counts new or changed notes into a status-bar chip. Counting never runs the model; drafting only happens when you click.
Config as notes — the rules live as plain markdown you can read, edit and version:
schema.md — your tag vocabulary, naming rules, a rejected list (a tag you deleted by hand does not come back), and an alias table : approving a retag records that the old tag means the new one, which is what lets a page still carrying llm-eval be recognised as being about evals . Before the plugin rewrites this file it saves the

[truncated]

## Original Extract

Run Gemma 4 E4B locally inside Obsidian via LiteRT-LM and WebGPU — private, offline, no API key, no Ollama, no server. Builds Andrej Karpathy's LLM wiki over your vault: cards, index, concept pages, tags, links, provenance checks — then chat with a note or the whole wiki, plus quizzes and flashcards
[truncated]

GitHub - itsyuimorii/obsidian-gemma4-litert-wiki: Run Gemma 4 E4B locally inside Obsidian via LiteRT-LM and WebGPU — private, offline, no API key, no Ollama, no server. Builds Andrej Karpathy's LLM wiki over your vault: cards, index, concept pages, tags, links, provenance checks — then chat with a note or the whole wiki, plus quizzes and flashcards. · GitHub
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
itsyuimorii
/
obsidian-gemma4-litert-wiki
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
339 Commits 339 Commits Folders and files
.github .github assets assets docs docs scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.ja.md README.ja.md README.md README.md SECURITY.md SECURITY.md build.js build.js eslint.config.mjs eslint.config.mjs manifest.json manifest.json package-lock.json package-lock.json package.json package.json styles.css styles.css tsconfig.json tsconfig.json versions.json versions.json View all files Repository files navigation
Turn your notes into a living Karpathy LLM wiki.
Chat with your vault, connect ideas, find gaps, and generate quizzes & flashcards —
with Gemma 4 E4B running locally inside Obsidian. No API key, Ollama, or server.
⬇ Install from the community store
▶ See what it does · Step through the demo — nothing to install.
☔️ What do you do with everything you have already saved?
Clipped articles. Saved posts. Course notes. Papers you meant to read. A vault that looks organised and is mostly unread — not because there is too much in it, but because there is no easy way to see what you've actually collected without opening things one at a time.
Gemma 4 E4B builds Andrej Karpathy's LLM wiki out of the notes you already have.
🧠 Chat with one note or your whole vault — grounded in what you actually wrote, with sources listed by the plugin rather than generated by the model.
📇 A summary card for every note, and an index over all of them — so you can see what you have without opening anything.
🔗 Links between notes that turn out to be related — and tags drawn from one shared vocabulary rather than whatever each note happened to invent, which is what lets them find each other at all.
📚 Concept pages over the clusters that emerge — written above your notes, linking down into each one.
🔍 Gaps and contradictions surfaced — including claims in two notes that cannot both be true.
🎓 Quizzes and flashcards from any note — so the archive is something you revisit, not just something you kept.
💬 Or ask it anything at all — Direct mode drops the grounding and answers from the model itself, offline and free. It is a 4B model, so it is weaker at general knowledge than anything behind an API, and it says so by carrying no sources and never filing its answers into your wiki.
Everything it writes is plain Markdown, in your vault.
Your own notes are never modified by the model. The wiki is a separate layer built above them, and your notes stay exactly as you wrote them.
And the model is already inside Obsidian. Gemma 4 E4B runs in Obsidian's own process through LiteRT-LM and WebGPU — no API key, no Ollama, no LM Studio, no localhost server, no account, no subscription .
There is no provider to pick. Your notes are never uploaded anywhere, because there is no server to upload them to. Privacy is a property of the architecture, not a promise in a policy.
After the one-time downloads — the ~3 GB model and the WASM runtime, both under Privacy — inference runs locally and the network is no longer required.
💬 Chat with your notes — entirely offline
🔌 How this differs from other AI plugins
🔧 How it works
Manual test protocols
Sixteen commands, and this is what they add up to. Every row is covered in
detail below; the full command list is under
Current commands .
Two of the sixteen write into a note you wrote, both behind a preview. The
rest build and maintain a layer beside your notes and never touch them.
It implements Andrej Karpathy's LLM-wiki pattern : your raw notes are never modified by the model — the immutability constraint binds the LLM, not you, and editing your own notes is the normal fix when a chat surfaces a mistake (the wiki detects the change and offers to re-ingest). A separate, cross-linked wiki layer is built above them — one card per note, and every page is shown to you in full before a single byte is written .
From there: chat grounded in one note or in the whole wiki, quiz yourself on either, build concept pages across everything sharing a tag, and let it flag its own decay with lint, provenance and contradiction checks. Sources are listed by the plugin, never cited by the model — citation is the one thing a small local model would get wrong without anyone noticing.
Everything it writes is plain markdown in your vault — nothing is locked in a database, and nothing needs another plugin to read it back. Its own configuration is notes too: your tag vocabulary and naming rules live in schema.md , where you can edit them by hand and the plugin will obey; every operation is appended to log.md , so you can always see what it did and when; and dropping a markdown file into skills/ adds a command of your own to the ⚡ menu.
Status: 1.0.13, in the community plugin store . The full Karpathy loop is implemented and running; benchmarks below are from real use.
The Karpathy loop — raw notes stay read-only; the plugin maintains a separate gemma-wiki/ layer:
Ingest — one strict JSON extraction per note (summary, tags, key points, salient mentions, and a self-rated high / med / low confidence written to frontmatter), plus a validated multiple-choice pick of related pages from the index catalog. Everything previews in a review modal; nothing is written without approval. Ingested notes get a badge in the file explorer — pure UI decoration, the note file is untouched.
Scan (semi-automatic ingest) — sweep the folders you named for new or changed notes, draft a card for each, and review them all in one list sorted low-confidence first , so the pages most in need of a human eye are the ones you read first. Untick anything that should not have become a page. Opt-in by scope: leave the folder list blank and scan refuses to run rather than sweeping your vault.
index.md / log.md — a one-line-per-page catalog the query path reads first, and an append-only, grep-friendly activity log.
Query (Wiki mode) — index-first retrieval with stopword filtering; the catalog and recent log always ride along, so meta-questions ("what did I add today?") work; answers are grounded only in retrieved material with honest refusals otherwise.
Whole-wiki questions ground in every page , not in the ones that lexically match. "What connects my pages" is about the shape of the collection, so scoring it against page summaries matches nothing — the words in the question appear in no summary — and the model is handed zero pages and correctly reports it was given none. The two chips that ask this kind of question carry the flag; loadPages fills to the budget and stops.
Save an answer as a note — every reply has a save action (same review gate), and it writes the answer into your own notes , beside the note it came from, named gemma — Summarize — The perfect espresso shot so the file explorer shows at a glance both that a model wrote it and what it was about, and recording in frontmatter which model and from which sources — the two things a copy-paste destroys silently. It lands in your folders, under your organisation, and it survives deleting the wiki folder, which nothing the plugin owns is meant to. Being an ordinary note of yours, it is also the one thing a saved answer never used to be: ingestable . Scan it and it becomes a card like any other note — which is the only route by which anything becomes material, because everything the wiki retrieves derives from a note you wrote.
Everything retrievable derives from a note you wrote. That holds structurally rather than by convention: only cards/ (one card per ingested note) and concepts/ (built across those cards) are in the retrieval set.
Saved answers are never retrieved. Karpathy's gist says explorations should compound and says nothing about trust, which is the poisoning loop — a wrong or stale answer re-entering context with the same weight as the notes. So a saved answer is not in the retrieval set at all: it is an ordinary note of yours, kept to re-read, with frontmatter recording which model wrote it and from which sources. The one route back in is deliberate and human — judge it, and ingest it like anything else you wrote.
Concept pages — pick a tag or mention two or more pages share, and the plugin writes a page above them that links down into each one. Ingest ripples into them, and member lists self-heal in both directions.
Keeping it honest — a wiki nobody reviews is the failure mode this is built against:
Review board — three ways a page goes bad, in one queue: low self-rated confidence, source drift (the raw note changed since ingest, caught by source_hash ), and staleness. Concept overviews are checked too.
Provenance spot-check — takes a page's key points back to the raw note to find the sentence each came from, and flags the ones that cannot be traced.
Contradiction sweep — checks pairs of pages that share a tag for claims that disagree, recently-changed pairs first. It flags with the model's reason quoted and never edits — you decide which one was wrong.
Tidy the wiki — one command that checks first and then fixes only what you tick. The check is model-free: orphan pages, index entries pointing at deleted pages, pages missing from the index, and how far your pages' tags have drifted from the vocabulary. The four repairs — make links mutual, drop dead index entries, rebuild the vocabulary, apply it to existing pages — are always offered, ticked where the check found something; each runs in turn behind its own preview. Availability is never tied to detection, because a check good enough to pick a sensible default is not good enough to be the only way in.
Background count (off by default) — periodically counts new or changed notes into a status-bar chip. Counting never runs the model; drafting only happens when you click.
Config as notes — the rules live as plain markdown you can read, edit and version:
schema.md — your tag vocabulary, naming rules, a rejected list (a tag you deleted by hand does not come back), and an alias table : approving a retag records that the old tag means the new one, which is what lets a page still carrying llm-eval be recognised as being about evals . Before the plugin rewrites this file it saves the

[truncated]
