---
source: "https://github.com/manavmishra/ZeroSlop"
hn_url: "https://news.ycombinator.com/item?id=49928432"
title: "Show HN: Zero Slop – Open-source skill that finds and fixes AI slop in writing"
article_title: "GitHub - manavmishra/ZeroSlop: Open-source Agent Skill that finds and removes AI slop from your writing · GitHub"
image: "https://repository-images.githubusercontent.com/1322140852/5f048a4c-df6f-40ce-a4a1-a08aba87c137"
author: "mmishra7"
captured_at: "2026-10-02T00:34:00Z"
capture_tool: "hn-digest"
hn_id: 49928432
score: 2
comments: 0
posted_at: "2026-10-02T00:04:16Z"
tags:
  - hacker-news
---

# Show HN: Zero Slop – Open-source skill that finds and fixes AI slop in writing

- HN: [49928432](https://news.ycombinator.com/item?id=49928432)
- Source: [github.com](https://github.com/manavmishra/ZeroSlop)
- Score: 2
- Comments: 0
- Posted: 2026-10-02T00:04:16Z

## Translation

Title: Show HN: Zero Slop – Open-source skill that finds and fixes AI slop in writing
Article title: GitHub - manavmishra/ZeroSlop: Open-source Agent Skill that finds and removes AI slop from your writing · GitHub
Description: Open-source Agent Skill that finds and removes AI slop from your writing - manavmishra/ZeroSlop

Article text:
GitHub - manavmishra/ZeroSlop: Open-source Agent Skill that finds and removes AI slop from your writing · GitHub
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
manavmishra
/
ZeroSlop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
330 Commits 330 Commits Folders and files
.claude-plugin .claude-plugin .codex-plugin .codex-plugin .codex-tmp .codex-tmp .github .github agents agents assets assets bench bench bin bin data data dist dist distribution distribution docs docs examples examples growth growth integrations integrations mcp mcp packaging/ zero_slop packaging/ zero_slop references references scripts scripts skills/ zero-slop skills/ zero-slop tests tests tooling tooling website website .codexignore .codexignore .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json .pre-commit-hooks.yaml .pre-commit-hooks.yaml AGENTS.md AGENTS.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DISTRIBUTION.md DISTRIBUTION.md EVALUATION.html EVALUATION.html LICENSE LICENSE ONE-PAGER.md ONE-PAGER.md README-nas-edit.md README-nas-edit.md README-pypi.md README-pypi.md README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md SUPPORT.md SUPPORT.md TERMS.md TERMS.md action.yml action.yml gemini-extension.json gemini-extension.json glama.json glama.json mcp.json mcp.json package-lock.json package-lock.json package.json package.json plugin.json plugin.json pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Find and remove AI slop in your writing. Get rid of workslop without losing your core intent and message.
Zero Slop is a free, open-source agent skill that finds and removes AI slop while checking that the core details of your message survive the edit. If your AI setup does not support agent skills, try the browser editor or use our MCP connector .
An AI draft can be grammatically sound and still read like workslop. In a compatible AI assistant, Zero Slop flags slop patterns, then guides the edit. The writing score finds patterns worth reviewing; it cannot tell who wrote the text.
Let's see Zero Slop at work. Imagine using AI to write a linkedin launch announcement and getting this:
We're thrilled to announce that our team has leveraged cutting-edge machine learning to deliver a seamless onboarding experience, reducing setup time by 40%.
The local Python scorer in Zero Slop scores the input Slop score 99.3/100. A high Slop score means the draft is more likely to contain sloppy patterns.
Zero Slop then strips the patterns and guides the AI agent to produce the deslopped output below:
We used machine learning to reduce onboarding setup time by 40%.
Writing score: 9.5/100 [clear]
Flagged phrases : 0 across 10 words
Quick start
You can try the browser editor without installing anything, install the skill in an assistant that supports skills, or use the hosted service through MCP and the API.
If you use Claude Code, Codex, or another assistant that supports skills, here's how to install Zero Slop there:
npx skills add manavmishra/ZeroSlop --global
/zero-slop (your writing)
To see the flagged passages without an edit, use /zero-slop inspect (your writing) .
The command above works with Claude Code and Codex. In Claude.ai, upload the skill ZIP . Other installation paths include Gemini CLI and remote MCP connections where your client allows them. You can also score a file locally without a model call.
Your AI assistant, whether Claude, GPT, or another compatible model, reads and edits the draft. The skill supplies the workflow and local tools: a 0 to 100 writing score, source-detail checks, and a final comparison with the original.
The scorer uses 294 weighted patterns and a 96-term lexicon. It checks for:
binary contrast formulas: “It's not X. It's Y.”
canned openers: “We're thrilled to…” and “Here's the thing…”
vague attribution: “experts agree” and “studies show”
significance inflation: “marks a pivotal moment” and “a testament to”
promotional wording: “robust,” “seamless,” and “leverage” when used as hype
repeated sentence shapes, crowded statistics, and overworked formatting
The Zero Slop agent uses an eight-stage workflow. Each stage is a job with a role, not a separate model; some run in the Python tools and others run in the user's AI app. We treat eight stages as an engineering convention, not eight separate models.
A saved, same-model editing test
We ran Zero Slop and three other open-source agent skills on our AI Slop test corpus, using GPT-5.4, high reasoning, and pinned instructions. Saved outputs are reproducible.
The RAID+ audit asks a different question: how much default writing from different models is flagged as AI slop by Zero Slop? The test corpus contains 7,627 anonymous, user-generated transcripts:
RAID+ records which model wrote each passage, not whether it reads well.
This is a feature comparison of Zero Slop against other popular slop tools.
The checks draw on research into predictable machine wording and overused vocabulary . Zero Slop cannot identify an author: detectors can misclassify non-native English .
You can teach Zero Slop a preference by giving it the original output, your edited version, and the reason for the change. Private data stays under $ZERO_SLOP_HOME and follows your privacy settings. Zero Slop is an AI slop detector and editor, not a plagiarism tool.
For developers: other ways to access Zero Slop
The endpoint for compatible clients is:
https://mcp.zero-slop.ai/mcp
Connection options .
For Gemini CLI, run gemini extensions install https://github.com/manavmishra/ZeroSlop --auto-update . For file-upload assistants, download the single-file bundle .
Score a file without sending it to a model:
npx zero-slop score draft.md
From a cloned checkout, check a folder against the review threshold of 25:
python3 scripts/slopscore.py --batch drafts/ --gate 25
Command line
The CLI sends a file to the hosted editor without changing the file on disk:
npx --yes zero-slop@2.12.12 deslop draft.md --genre professional
Use - for stdin and --json for structured output. --require-approved prints the result but exits nonzero when review is needed. Requires Node.js 22+; offline score also needs Python 3. CLI options and privacy .
The REST API accepts the same edit request:
curl --fail-with-body --max-time 75 https://mcp.zero-slop.ai/v1/deslop \
-H ' Content-Type: application/json ' \
--data ' {"text":"Maya owns the pricing review.","genre":"professional"} '
Check status before using an edit. Shared free capacity accepts up to 20,000 Unicode code points after trimming. API reference · OpenAPI contract
Path
Purpose
SKILL.md
The complete detect, rewrite, verify, and learn workflow
scripts/slopscore.py
Offline meter and source-detail gate
scripts/register.py
Performed-register and reading pass
references/
Genre guidance, tells, safeguards, and evaluation rules
examples/
Reproducible before-and-after edits
bench/
Frozen benchmarks, provenance, and limitations
mcp/
Optional hosted MCP server documentation
DISTRIBUTION.md
Direct installs, marketplace submissions, and release synchronization
Contribute or get help
Found a false positive, a broken check, or a better example? Use the issue forms or start a Discussion . If you want to change a pattern, read CONTRIBUTING.md and include tests with your pull request.
For setup help, see SUPPORT.md . Report security issues through SECURITY.md .
Zero Slop builds on ideas from First Reader , no-ai-slop , humanizer , de-slop , stop-slop , unslop-text , and avoid-ai-writing .
Open-source Agent Skill that finds and removes AI slop from your writing
Readme MIT license Code of conduct
Security policy Cite this repository Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source Agent Skill that finds and removes AI slop from your writing - manavmishra/ZeroSlop

GitHub - manavmishra/ZeroSlop: Open-source Agent Skill that finds and removes AI slop from your writing · GitHub
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
manavmishra
/
ZeroSlop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
330 Commits 330 Commits Folders and files
.claude-plugin .claude-plugin .codex-plugin .codex-plugin .codex-tmp .codex-tmp .github .github agents agents assets assets bench bench bin bin data data dist dist distribution distribution docs docs examples examples growth growth integrations integrations mcp mcp packaging/ zero_slop packaging/ zero_slop references references scripts scripts skills/ zero-slop skills/ zero-slop tests tests tooling tooling website website .codexignore .codexignore .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json .pre-commit-hooks.yaml .pre-commit-hooks.yaml AGENTS.md AGENTS.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DISTRIBUTION.md DISTRIBUTION.md EVALUATION.html EVALUATION.html LICENSE LICENSE ONE-PAGER.md ONE-PAGER.md README-nas-edit.md README-nas-edit.md README-pypi.md README-pypi.md README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md SUPPORT.md SUPPORT.md TERMS.md TERMS.md action.yml action.yml gemini-extension.json gemini-extension.json glama.json glama.json mcp.json mcp.json package-lock.json package-lock.json package.json package.json plugin.json plugin.json pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Find and remove AI slop in your writing. Get rid of workslop without losing your core intent and message.
Zero Slop is a free, open-source agent skill that finds and removes AI slop while checking that the core details of your message survive the edit. If your AI setup does not support agent skills, try the browser editor or use our MCP connector .
An AI draft can be grammatically sound and still read like workslop. In a compatible AI assistant, Zero Slop flags slop patterns, then guides the edit. The writing score finds patterns worth reviewing; it cannot tell who wrote the text.
Let's see Zero Slop at work. Imagine using AI to write a linkedin launch announcement and getting this:
We're thrilled to announce that our team has leveraged cutting-edge machine learning to deliver a seamless onboarding experience, reducing setup time by 40%.
The local Python scorer in Zero Slop scores the input Slop score 99.3/100. A high Slop score means the draft is more likely to contain sloppy patterns.
Zero Slop then strips the patterns and guides the AI agent to produce the deslopped output below:
We used machine learning to reduce onboarding setup time by 40%.
Writing score: 9.5/100 [clear]
Flagged phrases : 0 across 10 words
Quick start
You can try the browser editor without installing anything, install the skill in an assistant that supports skills, or use the hosted service through MCP and the API.
If you use Claude Code, Codex, or another assistant that supports skills, here's how to install Zero Slop there:
npx skills add manavmishra/ZeroSlop --global
/zero-slop (your writing)
To see the flagged passages without an edit, use /zero-slop inspect (your writing) .
The command above works with Claude Code and Codex. In Claude.ai, upload the skill ZIP . Other installation paths include Gemini CLI and remote MCP connections where your client allows them. You can also score a file locally without a model call.
Your AI assistant, whether Claude, GPT, or another compatible model, reads and edits the draft. The skill supplies the workflow and local tools: a 0 to 100 writing score, source-detail checks, and a final comparison with the original.
The scorer uses 294 weighted patterns and a 96-term lexicon. It checks for:
binary contrast formulas: “It's not X. It's Y.”
canned openers: “We're thrilled to…” and “Here's the thing…”
vague attribution: “experts agree” and “studies show”
significance inflation: “marks a pivotal moment” and “a testament to”
promotional wording: “robust,” “seamless,” and “leverage” when used as hype
repeated sentence shapes, crowded statistics, and overworked formatting
The Zero Slop agent uses an eight-stage workflow. Each stage is a job with a role, not a separate model; some run in the Python tools and others run in the user's AI app. We treat eight stages as an engineering convention, not eight separate models.
A saved, same-model editing test
We ran Zero Slop and three other open-source agent skills on our AI Slop test corpus, using GPT-5.4, high reasoning, and pinned instructions. Saved outputs are reproducible.
The RAID+ audit asks a different question: how much default writing from different models is flagged as AI slop by Zero Slop? The test corpus contains 7,627 anonymous, user-generated transcripts:
RAID+ records which model wrote each passage, not whether it reads well.
This is a feature comparison of Zero Slop against other popular slop tools.
The checks draw on research into predictable machine wording and overused vocabulary . Zero Slop cannot identify an author: detectors can misclassify non-native English .
You can teach Zero Slop a preference by giving it the original output, your edited version, and the reason for the change. Private data stays under $ZERO_SLOP_HOME and follows your privacy settings. Zero Slop is an AI slop detector and editor, not a plagiarism tool.
For developers: other ways to access Zero Slop
The endpoint for compatible clients is:
https://mcp.zero-slop.ai/mcp
Connection options .
For Gemini CLI, run gemini extensions install https://github.com/manavmishra/ZeroSlop --auto-update . For file-upload assistants, download the single-file bundle .
Score a file without sending it to a model:
npx zero-slop score draft.md
From a cloned checkout, check a folder against the review threshold of 25:
python3 scripts/slopscore.py --batch drafts/ --gate 25
Command line
The CLI sends a file to the hosted editor without changing the file on disk:
npx --yes zero-slop@2.12.12 deslop draft.md --genre professional
Use - for stdin and --json for structured output. --require-approved prints the result but exits nonzero when review is needed. Requires Node.js 22+; offline score also needs Python 3. CLI options and privacy .
The REST API accepts the same edit request:
curl --fail-with-body --max-time 75 https://mcp.zero-slop.ai/v1/deslop \
-H ' Content-Type: application/json ' \
--data ' {"text":"Maya owns the pricing review.","genre":"professional"} '
Check status before using an edit. Shared free capacity accepts up to 20,000 Unicode code points after trimming. API reference · OpenAPI contract
Path
Purpose
SKILL.md
The complete detect, rewrite, verify, and learn workflow
scripts/slopscore.py
Offline meter and source-detail gate
scripts/register.py
Performed-register and reading pass
references/
Genre guidance, tells, safeguards, and evaluation rules
examples/
Reproducible before-and-after edits
bench/
Frozen benchmarks, provenance, and limitations
mcp/
Optional hosted MCP server documentation
DISTRIBUTION.md
Direct installs, marketplace submissions, and release synchronization
Contribute or get help
Found a false positive, a broken check, or a better example? Use the issue forms or start a Discussion . If you want to change a pattern, read CONTRIBUTING.md and include tests with your pull request.
For setup help, see SUPPORT.md . Report security issues through SECURITY.md .
Zero Slop builds on ideas from First Reader , no-ai-slop , humanizer , de-slop , stop-slop , unslop-text , and avoid-ai-writing .
Open-source Agent Skill that finds and removes AI slop from your writing
Readme MIT license Code of conduct
Security policy Cite this repository Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
