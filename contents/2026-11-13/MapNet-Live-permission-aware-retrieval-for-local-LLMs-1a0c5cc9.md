---
source: "https://github.com/devinvok/MapNet"
hn_url: "https://news.ycombinator.com/item?id=50488267"
title: "MapNet: Live, permission-aware retrieval for local LLMs"
article_title: "GitHub - devinvok/MapNet · GitHub"
image: "https://opengraph.githubassets.com/c58f6658769f30fcb128398e7373ea91bbf41f5c4bd276faf491861e0ee74f37/devinvok/MapNet"
author: "invok"
captured_at: "2026-11-13T19:16:23Z"
capture_tool: "hn-digest"
hn_id: 50488267
score: 1
comments: 0
posted_at: "2026-11-13T19:08:35Z"
tags:
  - hacker-news
---

# MapNet: Live, permission-aware retrieval for local LLMs

- HN: [50488267](https://news.ycombinator.com/item?id=50488267)
- Source: [github.com](https://github.com/devinvok/MapNet)
- Score: 1
- Comments: 0
- Posted: 2026-11-13T19:08:35Z

## Translation

Title: MapNet: Live, permission-aware retrieval for local LLMs
Article title: GitHub - devinvok/MapNet · GitHub
Description: Contribute to devinvok/MapNet development by creating an account on GitHub.

Article text:
GitHub - devinvok/MapNet · GitHub
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
devinvok
MapNet Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
24 Commits 24 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples mapnet mapnet realtest realtest research research tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md agents_demo.py agents_demo.py ask.py ask.py compare.py compare.py pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
A small private model that knows where knowledge lives, not everything.
Instead of a giant model that memorizes the world, or one flat search index over
everything, MapNet keeps a live map of where each piece of knowledge is, the
way DNS knows which server holds a domain. A small local model only has to read
the few passages the map points to and write the answer.
Private. Runs entirely on your machine with a small model through
Ollama . Nothing leaves your computer.
Live. Add, edit or delete a file and the next question already sees it:
the map updates in milliseconds, with no re-indexing and no retraining.
Permission-aware. Put a .access file in a folder and only those roles can
reach it. Restricted branches are never even opened for other callers.
Agents leave trail markers. When an agent finds something, it leaves a
verified pointer so the next agent goes straight there. Pointers are checked on
every use; false ones are deleted.
Status: research prototype. It works end to end on the included sample company
and a synthetic benchmark of up to 92,000 documents. It has not been tested on
large real-world corpora yet.
git clone https://github.com/devinvok/MapNet.git
cd MapNet
pip install -r requirements.txt
# a small local model (optional: without it you get the best passage instead of an answer)
ollama pull qwen2.5:3b
python ask.py examples/company
Example session (the answer's wording depends on the model):
MapNet mapped 18 files (18 passages) in 0.1s. Model: qwen2.5:3b (local). Watching for changes.
> How many weeks of paid parental leave do birth parents get?
Birth parents get 20 weeks of fully paid parental leave. (source: hr/parental-leave-policy.md)
via map | hr, chat | opened 3 of 18 files | lookup 13 ms, answer 1.4s | note #1 left for next time
Then, while it's running, drop a new file into the folder (try
cp examples/drop-in/2026-10-23-checkout-outage.md examples/company/engineering/ )
and ask about it right away.
Point it at your own folder: python ask.py ~/Documents/work . Supported: .md ,
.txt , .csv , .pdf , .docx and other plain-text formats. Other models:
--model llama3.2:3b , --model phi4-mini , --model gemma3:4b . Roles:
--role hr,finance . Commands inside: /map , /notes , /role <names> , /quit .
question ──► notes: has an agent already found this? ──yes──► open that one file
│ no │
▼ │
walk the map (root → department → topic → file) │
│ keep the best few branches at each level │
▼ ▼
a handful of passages ─────────────► small local model ─► answer + source
│
└──► solid answer? leave a note for the next agent
Every node in the map advertises, about everything beneath it:
The map doesn't stop at the file. Like DNS, resolution goes from general to
specific: the folder map gets you to the file, and the file's own map gets you to
the passage. Each file is mapped the first time a question arrives at it : the
local model writes a short card per section, what it covers and the questions
people would ask in their own words, so "why is my namespace stuck deleting?"
lands on a page about finalizers . The map is cached like a DNS answer: revisits
are instant, a changed file is remapped on its next visit, and the cards flow up
into the folder map so the file is easier to find from then on. Files the lookup
touched but didn't open are mapped in the background, and every file gets a
one-line description at the folder level while the computer is idle, so the
router can find files no one has visited yet. With --meaning , the map is also
walked by meaning (a small local embedding model): it embeds the map, one vector
per file and, on arrival, per section, not every passage, and it only ever adds
candidates to the word router's, so a wrong guess can't push out a right one.
The work splits into three roles with very different models: a tiny
navigator that walks the map on every question and never writes text, a
cartographer that writes the map once per file, and a reader that only
ever sees the few passages found, so it can be a much bigger model
( --reader-model , --map-model ). Cards guide routing only;
they are never shown to the answering model. (To map a whole folder in advance:
python -m mapnet.cards FOLDER .)
Agent notes, and why they're safe
The idea came from a real failure mode: agents that discover a shared space and
start leaving each other instructions. MapNet's notes are deliberately limited:
Pointers, never instructions. A note has only "what it's about" and "which
file". Its text is never put in a model's prompt. Following a note only changes
which file gets opened.
Verified on every use. If the file doesn't answer, the note weakens; at zero
it's deleted. A planted false note dies the first time someone follows it.
Permission-bound and scoped. A note can only point to a file that exists,
is only shown to callers allowed to open that file, and is only used by callers
with the same roles as its writer.
Traceable and self-cleaning. Each note records its author, task and last
confirmation. Notes whose file is deleted disappear.
Command
What it shows
python ask.py examples/company
Interactive, private Q&A over a folder, live updates
python agents_demo.py
Later agents reach the right file in 1 step instead of ~10; five safety checks on notes
python compare.py
Accuracy, tokens and time on 27 questions, including 2 about files added seconds earlier
ANTHROPIC_API_KEY=... python compare.py --cloud --price-in 3 --price-out 15
Same, side by side with a large cloud model (reading the same passages, and reading every file)
python research/benchmark.py
Map vs flat search on a synthetic company of up to 2,304 systems
bash realtest/setup.sh && python -m realtest.run
The real test: 1,700 real Kubernetes docs, 40 hand-written questions, a local 3B model, against keyword and embedding search over everything. See realtest/README.md
See docs/DEMO_SCRIPT.md for a 2-minute recorded demo plan.
Sample company ( agents_demo.py ) , 18 files, 8 tasks:
Synthetic benchmark ( research/ ) , ordinary questions:
Real documents ( realtest/ ) , 475 Kubernetes docs (concepts, tasks, tutorials, setup),
7,703 passages, 4 passages handed to the reader per question, nomic-embed-text for
embeddings, everything on a CPU laptop. "Right file" means the file the question was written
from is among the 4 passages returned.
Answers, with qwen2.5:7b reading the 4 passages (35 hand-written questions, judged right
or wrong):
The honest reading: token-net+rerank finds the right file as often as embedding search,
but hands the reader slightly worse passages, and that costs 3 of 35 answers. On the test
half of the everyday questions (108, scored once, after the method was chosen on the other
half) it found the right file 51% of the time against embed-all's 56%. So embedding every
passage is still the more accurate search here; what token-net+rerank offers is search
that never embeds the corpus up front (7 s per question when nothing is cached yet), sees a
new file on the next question, and scores only what the caller may open. The samples are
small: one hand-written question is about 3 points, and a 7B judge kept only 114 of the 220
generated questions as sensible, so treat differences of a few points as noise.
We also tried, and measured, and closed: model-written file descriptions and reference
maps (51–63% right file), word associations, recursive spreading over the word network,
expanding questions with word embeddings, and letting a 7B model pick folders first (it
picked a folder holding the right page 39% of the time). Why each failed is in
docs/ROADMAP.md . Full tables are written to realtest/results/ by a run.
Honest caveats for the synthetic benchmark, all detailed in research/README.md :
the data is synthetic; when routing takes a wrong turn the answer is never seen
(flat search at least ranks it lower), so low-confidence queries should fall back
to a wider search; and on one machine a flat search is still faster in clock time,
since it's one matrix multiplication. The map's savings in work become savings in
time once sources live on different machines, or when the "searcher" is an agent
paying tokens for every file it opens.
mapnet/
core.py the map: nodes, routing, permissions, live add/remove, feedback
folder.py keeps the map in sync with a folder (polling watcher, .access files)
notes.py agent notes: pointer-only, verified, scoped, self-cleaning
llm.py readers: Ollama (local), Anthropic (cloud, for comparison), no-model fallback
cards.py section cards and one-line file descriptions written by a local model
embed.py Ollama embedding client with an on-disk cache
meaning.py walks the map by meaning as well as words
embedtree.py the map walked by embeddings, embedded on arrival
refmap.py a namespace of references (DNS style), with question-to-term translation
refgraph.py the same references as a navigable graph (HNSW, built from scratch)
tokennet.py a network of words pointing at the passages that hold them (BM25), per-caller permissions
shortlist.py the token network's top 50, embedded on arrival and ordered by meaning (the search core)
assistant.py the full loop: notes -> map -> reader -> leave a note
synthetic.py synthetic company generator for the benchmark
ask.py interactive CLI
agents_demo.py agents leaving trail markers + safety checks
compare.py local vs cloud comparison, including live updates
examples/ a fictional company's documents, files to drop in, question set
realtest/ the real-data test: 1,700 Kubernetes docs, hand-written questions, baselines
research/ scale benchmark, results and the synthetic walkthrough
tests/ pytest suite (runs without any model or network)
Swapping parts
Router encoder: the map uses a dependency-free hashed bag-of-words encoder
with live IDF, so new files never require re-fitting. Pass any object with
n_features and encode(texts) -> scipy.sparse rows as MapNet(embedder=...)
to try a learned embedding model.
Reader: any class with answer(question, hits, stream=None) -> Answer .
Sources: today a source i

[truncated]

## Original Extract

Contribute to devinvok/MapNet development by creating an account on GitHub.

GitHub - devinvok/MapNet · GitHub
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
devinvok
MapNet Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
24 Commits 24 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples mapnet mapnet realtest realtest research research tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md agents_demo.py agents_demo.py ask.py ask.py compare.py compare.py pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
A small private model that knows where knowledge lives, not everything.
Instead of a giant model that memorizes the world, or one flat search index over
everything, MapNet keeps a live map of where each piece of knowledge is, the
way DNS knows which server holds a domain. A small local model only has to read
the few passages the map points to and write the answer.
Private. Runs entirely on your machine with a small model through
Ollama . Nothing leaves your computer.
Live. Add, edit or delete a file and the next question already sees it:
the map updates in milliseconds, with no re-indexing and no retraining.
Permission-aware. Put a .access file in a folder and only those roles can
reach it. Restricted branches are never even opened for other callers.
Agents leave trail markers. When an agent finds something, it leaves a
verified pointer so the next agent goes straight there. Pointers are checked on
every use; false ones are deleted.
Status: research prototype. It works end to end on the included sample company
and a synthetic benchmark of up to 92,000 documents. It has not been tested on
large real-world corpora yet.
git clone https://github.com/devinvok/MapNet.git
cd MapNet
pip install -r requirements.txt
# a small local model (optional: without it you get the best passage instead of an answer)
ollama pull qwen2.5:3b
python ask.py examples/company
Example session (the answer's wording depends on the model):
MapNet mapped 18 files (18 passages) in 0.1s. Model: qwen2.5:3b (local). Watching for changes.
> How many weeks of paid parental leave do birth parents get?
Birth parents get 20 weeks of fully paid parental leave. (source: hr/parental-leave-policy.md)
via map | hr, chat | opened 3 of 18 files | lookup 13 ms, answer 1.4s | note #1 left for next time
Then, while it's running, drop a new file into the folder (try
cp examples/drop-in/2026-10-23-checkout-outage.md examples/company/engineering/ )
and ask about it right away.
Point it at your own folder: python ask.py ~/Documents/work . Supported: .md ,
.txt , .csv , .pdf , .docx and other plain-text formats. Other models:
--model llama3.2:3b , --model phi4-mini , --model gemma3:4b . Roles:
--role hr,finance . Commands inside: /map , /notes , /role <names> , /quit .
question ──► notes: has an agent already found this? ──yes──► open that one file
│ no │
▼ │
walk the map (root → department → topic → file) │
│ keep the best few branches at each level │
▼ ▼
a handful of passages ─────────────► small local model ─► answer + source
│
└──► solid answer? leave a note for the next agent
Every node in the map advertises, about everything beneath it:
The map doesn't stop at the file. Like DNS, resolution goes from general to
specific: the folder map gets you to the file, and the file's own map gets you to
the passage. Each file is mapped the first time a question arrives at it : the
local model writes a short card per section, what it covers and the questions
people would ask in their own words, so "why is my namespace stuck deleting?"
lands on a page about finalizers . The map is cached like a DNS answer: revisits
are instant, a changed file is remapped on its next visit, and the cards flow up
into the folder map so the file is easier to find from then on. Files the lookup
touched but didn't open are mapped in the background, and every file gets a
one-line description at the folder level while the computer is idle, so the
router can find files no one has visited yet. With --meaning , the map is also
walked by meaning (a small local embedding model): it embeds the map, one vector
per file and, on arrival, per section, not every passage, and it only ever adds
candidates to the word router's, so a wrong guess can't push out a right one.
The work splits into three roles with very different models: a tiny
navigator that walks the map on every question and never writes text, a
cartographer that writes the map once per file, and a reader that only
ever sees the few passages found, so it can be a much bigger model
( --reader-model , --map-model ). Cards guide routing only;
they are never shown to the answering model. (To map a whole folder in advance:
python -m mapnet.cards FOLDER .)
Agent notes, and why they're safe
The idea came from a real failure mode: agents that discover a shared space and
start leaving each other instructions. MapNet's notes are deliberately limited:
Pointers, never instructions. A note has only "what it's about" and "which
file". Its text is never put in a model's prompt. Following a note only changes
which file gets opened.
Verified on every use. If the file doesn't answer, the note weakens; at zero
it's deleted. A planted false note dies the first time someone follows it.
Permission-bound and scoped. A note can only point to a file that exists,
is only shown to callers allowed to open that file, and is only used by callers
with the same roles as its writer.
Traceable and self-cleaning. Each note records its author, task and last
confirmation. Notes whose file is deleted disappear.
Command
What it shows
python ask.py examples/company
Interactive, private Q&A over a folder, live updates
python agents_demo.py
Later agents reach the right file in 1 step instead of ~10; five safety checks on notes
python compare.py
Accuracy, tokens and time on 27 questions, including 2 about files added seconds earlier
ANTHROPIC_API_KEY=... python compare.py --cloud --price-in 3 --price-out 15
Same, side by side with a large cloud model (reading the same passages, and reading every file)
python research/benchmark.py
Map vs flat search on a synthetic company of up to 2,304 systems
bash realtest/setup.sh && python -m realtest.run
The real test: 1,700 real Kubernetes docs, 40 hand-written questions, a local 3B model, against keyword and embedding search over everything. See realtest/README.md
See docs/DEMO_SCRIPT.md for a 2-minute recorded demo plan.
Sample company ( agents_demo.py ) , 18 files, 8 tasks:
Synthetic benchmark ( research/ ) , ordinary questions:
Real documents ( realtest/ ) , 475 Kubernetes docs (concepts, tasks, tutorials, setup),
7,703 passages, 4 passages handed to the reader per question, nomic-embed-text for
embeddings, everything on a CPU laptop. "Right file" means the file the question was written
from is among the 4 passages returned.
Answers, with qwen2.5:7b reading the 4 passages (35 hand-written questions, judged right
or wrong):
The honest reading: token-net+rerank finds the right file as often as embedding search,
but hands the reader slightly worse passages, and that costs 3 of 35 answers. On the test
half of the everyday questions (108, scored once, after the method was chosen on the other
half) it found the right file 51% of the time against embed-all's 56%. So embedding every
passage is still the more accurate search here; what token-net+rerank offers is search
that never embeds the corpus up front (7 s per question when nothing is cached yet), sees a
new file on the next question, and scores only what the caller may open. The samples are
small: one hand-written question is about 3 points, and a 7B judge kept only 114 of the 220
generated questions as sensible, so treat differences of a few points as noise.
We also tried, and measured, and closed: model-written file descriptions and reference
maps (51–63% right file), word associations, recursive spreading over the word network,
expanding questions with word embeddings, and letting a 7B model pick folders first (it
picked a folder holding the right page 39% of the time). Why each failed is in
docs/ROADMAP.md . Full tables are written to realtest/results/ by a run.
Honest caveats for the synthetic benchmark, all detailed in research/README.md :
the data is synthetic; when routing takes a wrong turn the answer is never seen
(flat search at least ranks it lower), so low-confidence queries should fall back
to a wider search; and on one machine a flat search is still faster in clock time,
since it's one matrix multiplication. The map's savings in work become savings in
time once sources live on different machines, or when the "searcher" is an agent
paying tokens for every file it opens.
mapnet/
core.py the map: nodes, routing, permissions, live add/remove, feedback
folder.py keeps the map in sync with a folder (polling watcher, .access files)
notes.py agent notes: pointer-only, verified, scoped, self-cleaning
llm.py readers: Ollama (local), Anthropic (cloud, for comparison), no-model fallback
cards.py section cards and one-line file descriptions written by a local model
embed.py Ollama embedding client with an on-disk cache
meaning.py walks the map by meaning as well as words
embedtree.py the map walked by embeddings, embedded on arrival
refmap.py a namespace of references (DNS style), with question-to-term translation
refgraph.py the same references as a navigable graph (HNSW, built from scratch)
tokennet.py a network of words pointing at the passages that hold them (BM25), per-caller permissions
shortlist.py the token network's top 50, embedded on arrival and ordered by meaning (the search core)
assistant.py the full loop: notes -> map -> reader -> leave a note
synthetic.py synthetic company generator for the benchmark
ask.py interactive CLI
agents_demo.py agents leaving trail markers + safety checks
compare.py local vs cloud comparison, including live updates
examples/ a fictional company's documents, files to drop in, question set
realtest/ the real-data test: 1,700 Kubernetes docs, hand-written questions, baselines
research/ scale benchmark, results and the synthetic walkthrough
tests/ pytest suite (runs without any model or network)
Swapping parts
Router encoder: the map uses a dependency-free hashed bag-of-words encoder
with live IDF, so new files never require re-fitting. Pass any object with
n_features and encode(texts) -> scipy.sparse rows as MapNet(embedder=...)
to try a learned embedding model.
Reader: any class with answer(question, hits, stream=None) -> Answer .
Sources: today a source i

[truncated]
