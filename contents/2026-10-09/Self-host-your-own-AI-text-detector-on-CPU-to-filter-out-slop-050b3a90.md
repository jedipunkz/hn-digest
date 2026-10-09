---
source: "https://github.com/pablocaeg/sloptotal"
hn_url: "https://news.ycombinator.com/item?id=50018677"
title: "Self-host your own AI text detector on CPU to filter out slop"
article_title: "GitHub - pablocaeg/sloptotal: Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU. · GitHub"
image: "https://opengraph.githubassets.com/595587bf3b03def4c98ea27d034cf356d99633c74cf461a427a7a54f858a04b2/pablocaeg/sloptotal"
author: "sloptotal"
captured_at: "2026-10-09T11:41:12Z"
capture_tool: "hn-digest"
hn_id: 50018677
score: 1
comments: 0
posted_at: "2026-10-09T10:47:28Z"
tags:
  - hacker-news
---

# Self-host your own AI text detector on CPU to filter out slop

- HN: [50018677](https://news.ycombinator.com/item?id=50018677)
- Source: [github.com](https://github.com/pablocaeg/sloptotal)
- Score: 1
- Comments: 0
- Posted: 2026-10-09T10:47:28Z

## Translation

Title: Self-host your own AI text detector on CPU to filter out slop
Article title: GitHub - pablocaeg/sloptotal: Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU. · GitHub
Description: Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU. - pablocaeg/sloptotal

Article text:
GitHub - pablocaeg/sloptotal: Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU. · GitHub
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
pablocaeg
sloptotal Public Notifications You must be signed in to change notification settings
Star 55 ( 55 ) You must be signed in to star a repository
Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU.
Readme MIT license Code of conduct
Security policy Cite this repository Activity Stars
14 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
74 Commits 74 Commits Folders and files
.claude .claude .cursor/ rules .cursor/ rules .github .github app app benchmarks benchmarks docker docker docs docs examples examples scripts scripts tests tests web web .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
VirusTotal for AI-generated text. Paste text, drop in a PDF or Word file, or
give it a URL. Twenty-three independent AI detectors (neural classifiers,
statistical tests and linguistic heuristics) score it in parallel, and a
calibrated ensemble turns their votes into one verdict you can inspect engine by
engine. It runs on your own CPU, so nothing you scan leaves your machine.
It is a free, self-hosted, open-source alternative to hosted AI content
detectors such as GPTZero, Originality.ai, Copyleaks, ZeroGPT and Humalingo.
Instead of one number from one model, it shows you every model's opinion, and
it publishes how accurate that is, failures included.
Try it: sloptotal.com · Run it: docker run -p 8000:8000 ghcr.io/pablocaeg/sloptotal
23 detection engines, one calibrated score. DeBERTa and RoBERTa
classifiers, Binoculars, Fast-DetectGPT, GLTR, perplexity and burstiness
tests, and stock-phrase heuristics. Results stream in as each engine finishes.
Text, URLs and documents. Paste text, scan a web page (main content is
extracted automatically), or upload .pdf , .docx , .txt or .md .
Site check: was this website vibe-coded? Finds the fingerprints that
Lovable, v0, Bolt, Base44, Replit and Same leave in the sites they deploy,
and shows the evidence for each one. How it works
Per-paragraph heat map through the API, to see which parts read as AI.
Measured, not claimed. Every accuracy number below comes with the corpus,
the harness and the raw per-sample scores.
Private by default. Self-hosted, no third-party AI APIs, no tracking,
reports deleted after 30 days.
CPU-only is fine. Auto-detects your hardware; 4 GB RAM is enough for the
lite profile, a GPU is optional.
JSON API and a Chrome extension
that marks AI-looking results in Google Search and LinkedIn.
Most detectors publish an accuracy figure without saying what it was measured
on. SlopTotal is measured on SlopBench : 1,626 human
texts, every one written before ChatGPT, and 1,626 AI texts on the same topics
and at the same lengths from 14 current models, across 15 kinds of writing
(news, Wikipedia, arXiv, Stack Exchange, Reddit, reviews, student essays,
non-native English, fiction and literature published 1532-1915). Every number
below is measured on kinds of writing the model was not tuned on.
The score bands are anchored on that human text: 45 is where the top 5% of
human writing begins, 55 the top 2%, 80 the top 0.5%. So "Likely AI" means
fewer than 2 in 100 human texts score this high.
Other languages. Spanish, French, German, Italian, Portuguese, Dutch,
Polish, Russian and Japanese are supported (AUC 0.91 to 0.997); Arabic and
Korean are experimental; Hindi, Turkish and Chinese are not reliable yet, and
the report says so.
What does not work. Essays by non-native English writers are still flagged
more than native ones (15% of TOEFL essays called Likely AI, against none of 88
US school essays). Under about 80 words a score is a weak signal. AI text run
through a "humanizer" is caught about half the time. Source code is outside
what these engines do. All of it, per source, per model and per engine, is in
the findings , including that the
classifiers which top the RAID benchmark drop to AUC 0.75 on current models.
docker run -p 8000:8000 -v sloptotal-models:/app/models ghcr.io/pablocaeg/sloptotal
Open http://localhost:8000 . The first scan downloads about 2 GB of models into
the sloptotal-models volume, so later starts are quick. To build from source
instead, run docker compose -f docker/docker-compose.yml up .
Requires Python 3.10+ (macOS ships 3.9, which is too old).
git clone https://github.com/pablocaeg/sloptotal.git
cd sloptotal
python3.11 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
./scripts/start.sh # or: uvicorn app.main:app --port 8000
Check that every engine loads and scores, end to end:
python scripts/smoke_test.py # against http://localhost:8000
Site check: detect sites built with AI app builders
"Is this website vibe-coded?" checkers mostly score style (Tailwind class
counts, missing security headers, buzzwords) and turn it into a percentage.
Hand-written sites share all of those traits. SlopTotal looks only for markers
the builders themselves leave in what they deploy, each one confirmed on live
sites or in the builders' own templates:
A site with no marker may still have been written with AI: code exported from
these tools and hosted elsewhere, or written in an AI editor, carries no
fingerprint. So the result is evidence, not a probability. The page's copy is
scored separately by the text engines.
curl -X POST http://localhost:8000/api/scan/site \
-H " Content-Type: application/json " -d ' {"url": "example.com"} '
API
Endpoint
Method
What it does
Typical latency (CPU)
/api/analyze
POST
Full 23-engine report for text or url
2-8 s
/api/quick-score
POST
4 classifiers plus heuristics
0.1-0.5 s
/api/paragraph-score
POST
Score per paragraph (heat map)
1-3 s
/api/scan/site
POST
AI app builder fingerprints plus a copy score
1-3 s
/api/extract
POST
Text from an uploaded .pdf / .docx / .txt (multipart file )
< 1 s
/api/scan/snippets
POST
Batch of 1-30 short snippets
~0.5 s
/api/scan/urls
POST
Batch of 1-10 URLs, page-type aware
1-5 s
/api/engines
GET
Engine metadata
instant
/api/report/{id}
GET
A stored report
instant
/api/report/{id}/feedback
POST
Record who actually wrote the text: {"label": "human" | "ai" | "mixed" | "unsure"}
instant
/api/queue/status
GET
Queue capacity
instant
curl -X POST http://localhost:8000/api/analyze \
-H " Content-Type: application/json " \
-d ' {"text": "Your text to analyze here..."} '
From Python, examples/python_client.py analyses a
text, prints the five engines scoring highest and runs a site check, waiting
in the queue when the server is busy:
python examples/python_client.py " Paste at least 50 characters of text here... " example.com
The response lists every engine with its score, verdict and a plain-language
detail line, plus overall_score (0-100) and overall_verdict .
Every engine links to its page on sloptotal.com, which carries its measured scores against both corpora. AUC below is the probability the engine ranks a random AI passage above a random human one: 1.0 is perfect, 0.5 is a coin flip.
Engine
Model
AUC
Notes
Desklib DeBERTa
DeBERTa-v3-large (435M)
1.000
Strongest separation in our own tests
SuperAnnotate
RoBERTa-large (355M)
0.989
No measurable bias against archaic prose
E5-Small
E5 + LoRA (33M)
0.999
Matches far larger models at 33M params
TMR Detector
RoBERTa-base (125M)
1.000
RAID-trained, so RAID scores flatter it
BERT-tiny RAID
BERT-tiny (4.4M)
1.000
Answers in milliseconds
ReMoDetect
DeBERTa (184M)
0.941
Targets RLHF-aligned LLMs
ChatGPT Detector
RoBERTa-base (125M)
0.829
ChatGPT-specific
Fakespot
RoBERTa-base (125M)
0.999
Accurate on modern text, but +0.533 bias on pre-1920 prose
OpenAI Detector
RoBERTa-base (125M)
0.771
The 2019 GPT-2 detector; weaker on modern LLMs
Statistical Methods
Engine
Method
AUC
Log-Rank
Average log-rank under GPT-2
0.909
GLTR
Token rank distribution
0.904
Perplexity
GPT-2 perplexity scoring
0.901
Cross-Perplexity
Two-model perplexity comparison
0.891
Fast-DetectGPT
Conditional probability curvature
0.890
Binoculars
Cross-entropy ratio between two LMs
0.836
DivEye
Surprisal diversity
0.730
Linguistic Heuristics
Engine
Signal
AUC
Structural Analysis
Em-dash usage, sentence uniformity
0.836
Linguistic Markers
AI-preferred phrases ("delve", "tapestry"...)
0.713
Formulaic Patterns
Cliche openings and closings
0.698
Vocabulary Richness
Type-token ratio, hapax legomena
0.583
Readability Uniformity
Cross-paragraph consistency
0.581
Burstiness
Per-sentence perplexity variance
0.582
Sentiment & Hedging
Hedging and forced balance
0.522
The linguistic heuristics are weak on their own. They are kept because they fail independently of the neural classifiers, which is what makes them useful as tiebreakers rather than as evidence.
The final score is calibrated , not a simple average, and every weight is
derived from measurement rather than intuition. See
tests/eval/FINDINGS.md and
sloptotal.com/detect/ai-detector-ensemble/ .
Anchored on the unbiased classifiers -- Desklib, SuperAnnotate, E5 and
ReMoDetect all score high AUC with no measurable bias against older prose.
Their consensus is blended 60/40 with the full weighted set.
Weights from measurement -- each engine's share is proportional to
Somers' D (2*AUC - 1), scaled down by any bias it shows against archaic
writing. RAID-trained engines are damped because our corpus is RAID.
Confidence from agreement -- a tight cluster across independent engine
families is trustworthy; one confident engine is not.
Skepticism, but only when earned -- unanimous high classifier scores are
damped only when the text itself carries human markers (contractions,
first-person, slang). Applied unconditionally it fired on 69 of 70 AI samples
and 0 of 66 human ones, suppressing correct detections.
Fakespot was previously the anchor, weighted 0.13. It is accurate on modern text
(AUC 0.999) but scored pre-1920 human prose at 0.645 against 0.112 for modern
human writing -- the largest bias of any engine -- and anchoring amplified it.
Machiavelli scored 62.5. After demotion to 0.033, literary passages average 10.2
and none is flagged.
SlopTotal detects CPU, RAM and GPU at startup and picks a profile. Everything
can be overridden with environment variables; see .env.example .
High-RAM CPU servers (e.g. 64 GB, no GPU): you automatically get the performa

[truncated]

## Original Extract

Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU. - pablocaeg/sloptotal

GitHub - pablocaeg/sloptotal: Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU. · GitHub
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
pablocaeg
sloptotal Public Notifications You must be signed in to change notification settings
Star 55 ( 55 ) You must be signed in to star a repository
Open-source AI text detector and ChatGPT detector: 23 engines, one calibrated score, every number measured on a public benchmark. Self-hosted, runs on CPU.
Readme MIT license Code of conduct
Security policy Cite this repository Activity Stars
14 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
74 Commits 74 Commits Folders and files
.claude .claude .cursor/ rules .cursor/ rules .github .github app app benchmarks benchmarks docker docker docs docs examples examples scripts scripts tests tests web web .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
VirusTotal for AI-generated text. Paste text, drop in a PDF or Word file, or
give it a URL. Twenty-three independent AI detectors (neural classifiers,
statistical tests and linguistic heuristics) score it in parallel, and a
calibrated ensemble turns their votes into one verdict you can inspect engine by
engine. It runs on your own CPU, so nothing you scan leaves your machine.
It is a free, self-hosted, open-source alternative to hosted AI content
detectors such as GPTZero, Originality.ai, Copyleaks, ZeroGPT and Humalingo.
Instead of one number from one model, it shows you every model's opinion, and
it publishes how accurate that is, failures included.
Try it: sloptotal.com · Run it: docker run -p 8000:8000 ghcr.io/pablocaeg/sloptotal
23 detection engines, one calibrated score. DeBERTa and RoBERTa
classifiers, Binoculars, Fast-DetectGPT, GLTR, perplexity and burstiness
tests, and stock-phrase heuristics. Results stream in as each engine finishes.
Text, URLs and documents. Paste text, scan a web page (main content is
extracted automatically), or upload .pdf , .docx , .txt or .md .
Site check: was this website vibe-coded? Finds the fingerprints that
Lovable, v0, Bolt, Base44, Replit and Same leave in the sites they deploy,
and shows the evidence for each one. How it works
Per-paragraph heat map through the API, to see which parts read as AI.
Measured, not claimed. Every accuracy number below comes with the corpus,
the harness and the raw per-sample scores.
Private by default. Self-hosted, no third-party AI APIs, no tracking,
reports deleted after 30 days.
CPU-only is fine. Auto-detects your hardware; 4 GB RAM is enough for the
lite profile, a GPU is optional.
JSON API and a Chrome extension
that marks AI-looking results in Google Search and LinkedIn.
Most detectors publish an accuracy figure without saying what it was measured
on. SlopTotal is measured on SlopBench : 1,626 human
texts, every one written before ChatGPT, and 1,626 AI texts on the same topics
and at the same lengths from 14 current models, across 15 kinds of writing
(news, Wikipedia, arXiv, Stack Exchange, Reddit, reviews, student essays,
non-native English, fiction and literature published 1532-1915). Every number
below is measured on kinds of writing the model was not tuned on.
The score bands are anchored on that human text: 45 is where the top 5% of
human writing begins, 55 the top 2%, 80 the top 0.5%. So "Likely AI" means
fewer than 2 in 100 human texts score this high.
Other languages. Spanish, French, German, Italian, Portuguese, Dutch,
Polish, Russian and Japanese are supported (AUC 0.91 to 0.997); Arabic and
Korean are experimental; Hindi, Turkish and Chinese are not reliable yet, and
the report says so.
What does not work. Essays by non-native English writers are still flagged
more than native ones (15% of TOEFL essays called Likely AI, against none of 88
US school essays). Under about 80 words a score is a weak signal. AI text run
through a "humanizer" is caught about half the time. Source code is outside
what these engines do. All of it, per source, per model and per engine, is in
the findings , including that the
classifiers which top the RAID benchmark drop to AUC 0.75 on current models.
docker run -p 8000:8000 -v sloptotal-models:/app/models ghcr.io/pablocaeg/sloptotal
Open http://localhost:8000 . The first scan downloads about 2 GB of models into
the sloptotal-models volume, so later starts are quick. To build from source
instead, run docker compose -f docker/docker-compose.yml up .
Requires Python 3.10+ (macOS ships 3.9, which is too old).
git clone https://github.com/pablocaeg/sloptotal.git
cd sloptotal
python3.11 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
./scripts/start.sh # or: uvicorn app.main:app --port 8000
Check that every engine loads and scores, end to end:
python scripts/smoke_test.py # against http://localhost:8000
Site check: detect sites built with AI app builders
"Is this website vibe-coded?" checkers mostly score style (Tailwind class
counts, missing security headers, buzzwords) and turn it into a percentage.
Hand-written sites share all of those traits. SlopTotal looks only for markers
the builders themselves leave in what they deploy, each one confirmed on live
sites or in the builders' own templates:
A site with no marker may still have been written with AI: code exported from
these tools and hosted elsewhere, or written in an AI editor, carries no
fingerprint. So the result is evidence, not a probability. The page's copy is
scored separately by the text engines.
curl -X POST http://localhost:8000/api/scan/site \
-H " Content-Type: application/json " -d ' {"url": "example.com"} '
API
Endpoint
Method
What it does
Typical latency (CPU)
/api/analyze
POST
Full 23-engine report for text or url
2-8 s
/api/quick-score
POST
4 classifiers plus heuristics
0.1-0.5 s
/api/paragraph-score
POST
Score per paragraph (heat map)
1-3 s
/api/scan/site
POST
AI app builder fingerprints plus a copy score
1-3 s
/api/extract
POST
Text from an uploaded .pdf / .docx / .txt (multipart file )
< 1 s
/api/scan/snippets
POST
Batch of 1-30 short snippets
~0.5 s
/api/scan/urls
POST
Batch of 1-10 URLs, page-type aware
1-5 s
/api/engines
GET
Engine metadata
instant
/api/report/{id}
GET
A stored report
instant
/api/report/{id}/feedback
POST
Record who actually wrote the text: {"label": "human" | "ai" | "mixed" | "unsure"}
instant
/api/queue/status
GET
Queue capacity
instant
curl -X POST http://localhost:8000/api/analyze \
-H " Content-Type: application/json " \
-d ' {"text": "Your text to analyze here..."} '
From Python, examples/python_client.py analyses a
text, prints the five engines scoring highest and runs a site check, waiting
in the queue when the server is busy:
python examples/python_client.py " Paste at least 50 characters of text here... " example.com
The response lists every engine with its score, verdict and a plain-language
detail line, plus overall_score (0-100) and overall_verdict .
Every engine links to its page on sloptotal.com, which carries its measured scores against both corpora. AUC below is the probability the engine ranks a random AI passage above a random human one: 1.0 is perfect, 0.5 is a coin flip.
Engine
Model
AUC
Notes
Desklib DeBERTa
DeBERTa-v3-large (435M)
1.000
Strongest separation in our own tests
SuperAnnotate
RoBERTa-large (355M)
0.989
No measurable bias against archaic prose
E5-Small
E5 + LoRA (33M)
0.999
Matches far larger models at 33M params
TMR Detector
RoBERTa-base (125M)
1.000
RAID-trained, so RAID scores flatter it
BERT-tiny RAID
BERT-tiny (4.4M)
1.000
Answers in milliseconds
ReMoDetect
DeBERTa (184M)
0.941
Targets RLHF-aligned LLMs
ChatGPT Detector
RoBERTa-base (125M)
0.829
ChatGPT-specific
Fakespot
RoBERTa-base (125M)
0.999
Accurate on modern text, but +0.533 bias on pre-1920 prose
OpenAI Detector
RoBERTa-base (125M)
0.771
The 2019 GPT-2 detector; weaker on modern LLMs
Statistical Methods
Engine
Method
AUC
Log-Rank
Average log-rank under GPT-2
0.909
GLTR
Token rank distribution
0.904
Perplexity
GPT-2 perplexity scoring
0.901
Cross-Perplexity
Two-model perplexity comparison
0.891
Fast-DetectGPT
Conditional probability curvature
0.890
Binoculars
Cross-entropy ratio between two LMs
0.836
DivEye
Surprisal diversity
0.730
Linguistic Heuristics
Engine
Signal
AUC
Structural Analysis
Em-dash usage, sentence uniformity
0.836
Linguistic Markers
AI-preferred phrases ("delve", "tapestry"...)
0.713
Formulaic Patterns
Cliche openings and closings
0.698
Vocabulary Richness
Type-token ratio, hapax legomena
0.583
Readability Uniformity
Cross-paragraph consistency
0.581
Burstiness
Per-sentence perplexity variance
0.582
Sentiment & Hedging
Hedging and forced balance
0.522
The linguistic heuristics are weak on their own. They are kept because they fail independently of the neural classifiers, which is what makes them useful as tiebreakers rather than as evidence.
The final score is calibrated , not a simple average, and every weight is
derived from measurement rather than intuition. See
tests/eval/FINDINGS.md and
sloptotal.com/detect/ai-detector-ensemble/ .
Anchored on the unbiased classifiers -- Desklib, SuperAnnotate, E5 and
ReMoDetect all score high AUC with no measurable bias against older prose.
Their consensus is blended 60/40 with the full weighted set.
Weights from measurement -- each engine's share is proportional to
Somers' D (2*AUC - 1), scaled down by any bias it shows against archaic
writing. RAID-trained engines are damped because our corpus is RAID.
Confidence from agreement -- a tight cluster across independent engine
families is trustworthy; one confident engine is not.
Skepticism, but only when earned -- unanimous high classifier scores are
damped only when the text itself carries human markers (contractions,
first-person, slang). Applied unconditionally it fired on 69 of 70 AI samples
and 0 of 66 human ones, suppressing correct detections.
Fakespot was previously the anchor, weighted 0.13. It is accurate on modern text
(AUC 0.999) but scored pre-1920 human prose at 0.645 against 0.112 for modern
human writing -- the largest bias of any engine -- and anchoring amplified it.
Machiavelli scored 62.5. After demotion to 0.033, literary passages average 10.2
and none is flagged.
SlopTotal detects CPU, RAM and GPU at startup and picks a profile. Everything
can be overridden with environment variables; see .env.example .
High-RAM CPU servers (e.g. 64 GB, no GPU): you automatically get the performa

[truncated]
