---
source: "https://fsdatalab.github.io/quail/docs"
hn_url: "https://news.ycombinator.com/item?id=50679888"
title: "Quail: A fast engine for AI functions in SQL"
article_title: "Introduction | Quail"
image: "https://fsdatalab.github.io/quail/og/docs/image.png"
author: "pretext"
captured_at: "2026-11-28T17:58:21Z"
capture_tool: "hn-digest"
hn_id: 50679888
score: 1
comments: 0
posted_at: "2026-11-28T17:28:30Z"
tags:
  - hacker-news
---

# Quail: A fast engine for AI functions in SQL

- HN: [50679888](https://news.ycombinator.com/item?id=50679888)
- Source: [fsdatalab.github.io](https://fsdatalab.github.io/quail/docs)
- Score: 1
- Comments: 0
- Posted: 2026-11-28T17:28:30Z

## Translation

Title: Quail: A fast engine for AI functions in SQL
Article title: Introduction | Quail

Article text:
Introduction | Quail Quail Docs Quail Docs Search ⌘ K Introduction When to use Quail User guide Quickstart Running queries SQL reference Querying multimodal data Query results and plans Metrics Quail Server Models and GPUs Cookbook Cookbook Reading history from old restaurant menus Developers Full Stack Data Lab at Carnegie Mellon University
A fast engine for AI functions in SQL.
GitHub query.sql SELECT r.id FROM reviews r
WHERE AI.IF ( PROMPT (
'Does the review discuss the ending? {0} ' ,
r.body)) id body AI.IF r1 The ending was perfect and the cast was great. true r2 Too long, and the ending made no sense. true r3 Great music, but I want my money back. false
Quickstart Run your first query. Running queries Choose a model, register tables, and run SQL. SQL reference Every AI function and option. Multimodal data Query images and PDFs. Models and GPUs Supported models and GPUs, and Modal. Quail Server Run queries on a remote GPU machine.
AI functions
Each AI function returns one answer per row. SQL uses the answer like any other value.
AI.IF Yes or no AI.SCORE A score from 0 to 1 AI.CLASSIFY One label from a list AI.EXTRACT Answers to questions AI.EXTRACT_SPAN Answers copied from the text AI.MAP Open-ended text AI.PARSE_DOCUMENT The text of a PDF AI.SAE_FEATURES Sparse autoencoder features
Install
Quail needs Python 3.12 and a supported CUDA GPU.
Terminal uv pip install quail-engine The package is quail-engine , and you import it as quail . With no
local GPU, run Quail on Modal
or connect to Quail Server .
Use Quail to run AI functions in SQL with open models on your own GPUs or on Modal, for much less than a hosted model costs.

## Original Extract

Introduction | Quail Quail Docs Quail Docs Search ⌘ K Introduction When to use Quail User guide Quickstart Running queries SQL reference Querying multimodal data Query results and plans Metrics Quail Server Models and GPUs Cookbook Cookbook Reading history from old restaurant menus Developers Full Stack Data Lab at Carnegie Mellon University
A fast engine for AI functions in SQL.
GitHub query.sql SELECT r.id FROM reviews r
WHERE AI.IF ( PROMPT (
'Does the review discuss the ending? {0} ' ,
r.body)) id body AI.IF r1 The ending was perfect and the cast was great. true r2 Too long, and the ending made no sense. true r3 Great music, but I want my money back. false
Quickstart Run your first query. Running queries Choose a model, register tables, and run SQL. SQL reference Every AI function and option. Multimodal data Query images and PDFs. Models and GPUs Supported models and GPUs, and Modal. Quail Server Run queries on a remote GPU machine.
AI functions
Each AI function returns one answer per row. SQL uses the answer like any other value.
AI.IF Yes or no AI.SCORE A score from 0 to 1 AI.CLASSIFY One label from a list AI.EXTRACT Answers to questions AI.EXTRACT_SPAN Answers copied from the text AI.MAP Open-ended text AI.PARSE_DOCUMENT The text of a PDF AI.SAE_FEATURES Sparse autoencoder features
Install
Quail needs Python 3.12 and a supported CUDA GPU.
Terminal uv pip install quail-engine The package is quail-engine , and you import it as quail . With no
local GPU, run Quail on Modal
or connect to Quail Server .
Use Quail to run AI functions in SQL with open models on your own GPUs or on Modal, for much less than a hosted model costs.
