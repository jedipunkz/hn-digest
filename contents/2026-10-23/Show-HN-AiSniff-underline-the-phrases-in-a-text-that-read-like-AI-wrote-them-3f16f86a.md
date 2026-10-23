---
source: "https://github.com/Arthur031221/AiSniff"
hn_url: "https://news.ycombinator.com/item?id=50215712"
title: "Show HN: AiSniff, underline the phrases in a text that read like AI wrote them"
article_title: "GitHub - Arthur031221/AiSniff: Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint · GitHub"
image: "https://opengraph.githubassets.com/725b75b0abf3a81a3c3dacba3fb0e94f44f999d4ce3a4d4da1832caf6d63e518/Arthur031221/AiSniff"
author: "levi74108520963"
captured_at: "2026-10-23T21:52:29Z"
capture_tool: "hn-digest"
hn_id: 50215712
score: 1
comments: 0
posted_at: "2026-10-23T21:36:29Z"
tags:
  - hacker-news
---

# Show HN: AiSniff, underline the phrases in a text that read like AI wrote them

- HN: [50215712](https://news.ycombinator.com/item?id=50215712)
- Source: [github.com](https://github.com/Arthur031221/AiSniff)
- Score: 1
- Comments: 0
- Posted: 2026-10-23T21:36:29Z

## Translation

Title: Show HN: AiSniff, underline the phrases in a text that read like AI wrote them
Article title: GitHub - Arthur031221/AiSniff: Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint · GitHub
Description: Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint - Arthur031221/AiSniff

Article text:
GitHub - Arthur031221/AiSniff: Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint · GitHub
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
Arthur031221
AiSniff Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
.github/ workflows .github/ workflows assets assets docs docs public public rules rules scripts scripts skills/ aisniff skills/ aisniff src src tests tests .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.de.md README.de.md README.es.md README.es.md README.fr.md README.fr.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.pt-BR.md README.pt-BR.md README.zh-CN.md README.zh-CN.md README.zh-TW.md README.zh-TW.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md action.yml action.yml index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.cli.config.ts vite.cli.config.ts vite.config.ts vite.config.ts vitest.e2e.config.ts vitest.e2e.config.ts View all files Repository files navigation
Paste text. See which phrases read like AI wrote them.
English | Traditional Chinese | Simplified Chinese | Japanese | Korean | Spanish | French | German | Portuguese (BR)
AiSniff marks phrases in English and Chinese text that match documented AI-writing patterns. Each mark names its rule, cites a public source and suggests a plain rewrite. The web editor, the CLI and the agent skill all run from one set of JSON rules, with no server and no model call.
It reports patterns and says nothing about who wrote the text. A clean result doesn't prove a person wrote it, and a hit doesn't prove a model did.
Before
After
In today's fast-paced world, our update is a game-changer that unlocks your full potential. We are committed to excellence, and we are thrilled to share it.
Version 2.1 fixes the crash on startup when the config file is empty. Details are in the changelog.
3 hits : en.opener.todays-world , en.promo.commitment
0 hits
Hit counts come from aisniff lint run on these exact sentences. I rewrote the right column by hand, following the hints.
On answers held out from HC3 (a public set of questions with human and chatbot answers), the English rules marked at least one phrase in 29% of chatbot answers and 2% of human answers. For Chinese it was 24% and 6%.
The limits matter. The test half holds 23,752 human and 13,422 chatbot English answers, and 9,812 and 8,480 Chinese ones. The chatbot answers date from its first months, so the rules for newer-model vocabulary can't be checked against them. 71% of the English chatbot answers had no hit at all. The AUROC of the density is 0.63 for English and 0.59 for Chinese, so density alone is a weak detector. Several rules fire as often on human answers as on chatbot ones, and the evaluation page lists them. Method and per-rule counts are in docs/evaluation.md . Rerun it with npm run eval .
npx github:Arthur031221/AiSniff lint draft.md
Command
What it does
aisniff lint [files...]
Scans files, or stdin when no file is given. Prints each hit with position, confidence, rule, message and hint.
aisniff lint --format json
Same hits as JSON, with source URLs, counts per type and the density.
aisniff lint --format github
One GitHub Actions annotation per hit.
aisniff lint --min-confidence medium
Hides weak signals. Values: low , medium , high .
aisniff lint --max-density 20
Exits 1 when hits per 1,000 units are above the number.
aisniff lint --max-weighted-density 8.7
Exits 1 when the weighted density, the scale of the web meter, is above the number.
aisniff lint --disable en.promo
Skips one rule, or every rule under a prefix such as en.promo . Repeat the flag to skip more.
aisniff lint docs/
Scans a folder and its subfolders for .md , .markdown , .mdx and .txt files.
aisniff lint --markdown off
Scans Markdown files as plain text. By default, code blocks, inline code, front matter, comments and link addresses in .md , .markdown and .mdx files are skipped. Values: auto , on , off .
.aisniff.json
Sets the same options for a project: format , minConfidence , maxDensity , maxWeightedDensity , disable , markdown , lang , locale . Found from the current folder up to the repository root. Flags win. --config FILE picks a file, --no-config ignores it.
aisniff lint --lang en
Loads one pack. Values: auto , en , zh-TW .
aisniff lint --locale zh-TW
Messages and hints in Traditional Chinese.
aisniff rules
Lists all 61 rules with their confidence.
aisniff explain <rule-id>
Shows one rule with passing and failing examples and its source.
Exit codes: 0 scanned within the limit, 1 limit exceeded, 2 bad usage. The text output ends with a line saying it measures patterns, not authorship.
Open the live demo or run npm install && npm run dev .
Hover or tap a mark to see the rule, the rewrite hint and the source. Some rules offer a one-click rewrite, and Undo still works after it.
A side panel groups the hits by type. A meter shows hits per 1,000 words or characters, next to a line saying it counts patterns, not authorship.
A compare tab holds a before and an after text and shows how the hit count changed.
Save the annotated text as a PNG.
Open a .txt , .md , .markdown or .mdx file of up to 2 MB with the button or by dropping it on the editor. The file is read in the browser. The web editor scans Markdown as plain text, so use the CLI to skip code blocks.
Nine interface languages: English, Traditional Chinese, Simplified Chinese, Japanese, Korean, Spanish, French, German, Portuguese (BR). Rule messages exist in English and Traditional Chinese.
Simplified Chinese input goes through the Traditional Chinese rules after a character conversion that keeps positions intact.
Web
Use the live demo . Nothing you paste leaves the page. To host it yourself, run npm run build:web and serve dist/web .
Needs Node 22 or newer. CI runs 22 and 24.
npx github:Arthur031221/AiSniff lint draft.md
Or from a checkout: npm install && npm run build:cli && npm link . The npm package name is aisniff-cli , the command is aisniff .
git clone https://github.com/Arthur031221/AiSniff.git
cp -r AiSniff/skills/aisniff ~ / < agent-dir > /skills/aisniff
Then ask: "check this draft with AiSniff". Use <agent-dir>/skills/aisniff inside a project to scope it to that project.
Copy skills/aisniff/SKILL.md into your repository and point to it from AGENTS.md , or paste its body into the agent's instructions. The agent only needs a shell to run npx .
- uses : actions/checkout@v4
- uses : Arthur031221/AiSniff@v0.5.0
with :
paths : docs/
max-weighted-density : 8.7
Where the rules come from
Each rule cites one of six public sources: the Wikipedia guide Signs of AI writing , the pattern lists in blader/humanizer , Humanizer-zh-TW and speak-human-tw , the Chinese Wikipedia page on Europeanized Chinese , and the excess-vocabulary study by Kobak et al., Science Advances 2025 . Those lists are prose guides for rewriting. AiSniff turns their patterns into rules, tests them against examples and measures them on data. It also keeps some rules the data doesn't support, and the evaluation page lists each one. Every rule has passing and failing examples that run as tests, and npm run check:sources checks that each source link and section exists.
Some wording rules mark mainland usage in Traditional Chinese (for example the mainland words for software and video). They flag wording that differs from Taiwan usage. They aren't AI tells.
It finds phrases from published lists. A model that avoids them, or a person who uses them, defeats it.
Formal writing, second-language writing and old-fashioned prose will trigger rules.
English and Chinese only. Other languages are not checked.
The "many patterns" line on the meter is the 99th percentile of human HC3 answers (8.7 and 13.0 weighted hits per 1,000 units). It comes from short Q&A text and is a reference line, not a decision threshold.
Install | Usage | Configuration | Architecture | Writing a rule | Evaluation | Contributing
The Simplified to Traditional table comes from OpenCC (Apache-2.0), see THIRD_PARTY_NOTICES.md .
Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint - Arthur031221/AiSniff

GitHub - Arthur031221/AiSniff: Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint · GitHub
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
Arthur031221
AiSniff Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
.github/ workflows .github/ workflows assets assets docs docs public public rules rules scripts scripts skills/ aisniff skills/ aisniff src src tests tests .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.de.md README.de.md README.es.md README.es.md README.fr.md README.fr.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.pt-BR.md README.pt-BR.md README.zh-CN.md README.zh-CN.md README.zh-TW.md README.zh-TW.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md action.yml action.yml index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.cli.config.ts vite.cli.config.ts vite.config.ts vite.config.ts vitest.e2e.config.ts vitest.e2e.config.ts View all files Repository files navigation
Paste text. See which phrases read like AI wrote them.
English | Traditional Chinese | Simplified Chinese | Japanese | Korean | Spanish | French | German | Portuguese (BR)
AiSniff marks phrases in English and Chinese text that match documented AI-writing patterns. Each mark names its rule, cites a public source and suggests a plain rewrite. The web editor, the CLI and the agent skill all run from one set of JSON rules, with no server and no model call.
It reports patterns and says nothing about who wrote the text. A clean result doesn't prove a person wrote it, and a hit doesn't prove a model did.
Before
After
In today's fast-paced world, our update is a game-changer that unlocks your full potential. We are committed to excellence, and we are thrilled to share it.
Version 2.1 fixes the crash on startup when the config file is empty. Details are in the changelog.
3 hits : en.opener.todays-world , en.promo.commitment
0 hits
Hit counts come from aisniff lint run on these exact sentences. I rewrote the right column by hand, following the hints.
On answers held out from HC3 (a public set of questions with human and chatbot answers), the English rules marked at least one phrase in 29% of chatbot answers and 2% of human answers. For Chinese it was 24% and 6%.
The limits matter. The test half holds 23,752 human and 13,422 chatbot English answers, and 9,812 and 8,480 Chinese ones. The chatbot answers date from its first months, so the rules for newer-model vocabulary can't be checked against them. 71% of the English chatbot answers had no hit at all. The AUROC of the density is 0.63 for English and 0.59 for Chinese, so density alone is a weak detector. Several rules fire as often on human answers as on chatbot ones, and the evaluation page lists them. Method and per-rule counts are in docs/evaluation.md . Rerun it with npm run eval .
npx github:Arthur031221/AiSniff lint draft.md
Command
What it does
aisniff lint [files...]
Scans files, or stdin when no file is given. Prints each hit with position, confidence, rule, message and hint.
aisniff lint --format json
Same hits as JSON, with source URLs, counts per type and the density.
aisniff lint --format github
One GitHub Actions annotation per hit.
aisniff lint --min-confidence medium
Hides weak signals. Values: low , medium , high .
aisniff lint --max-density 20
Exits 1 when hits per 1,000 units are above the number.
aisniff lint --max-weighted-density 8.7
Exits 1 when the weighted density, the scale of the web meter, is above the number.
aisniff lint --disable en.promo
Skips one rule, or every rule under a prefix such as en.promo . Repeat the flag to skip more.
aisniff lint docs/
Scans a folder and its subfolders for .md , .markdown , .mdx and .txt files.
aisniff lint --markdown off
Scans Markdown files as plain text. By default, code blocks, inline code, front matter, comments and link addresses in .md , .markdown and .mdx files are skipped. Values: auto , on , off .
.aisniff.json
Sets the same options for a project: format , minConfidence , maxDensity , maxWeightedDensity , disable , markdown , lang , locale . Found from the current folder up to the repository root. Flags win. --config FILE picks a file, --no-config ignores it.
aisniff lint --lang en
Loads one pack. Values: auto , en , zh-TW .
aisniff lint --locale zh-TW
Messages and hints in Traditional Chinese.
aisniff rules
Lists all 61 rules with their confidence.
aisniff explain <rule-id>
Shows one rule with passing and failing examples and its source.
Exit codes: 0 scanned within the limit, 1 limit exceeded, 2 bad usage. The text output ends with a line saying it measures patterns, not authorship.
Open the live demo or run npm install && npm run dev .
Hover or tap a mark to see the rule, the rewrite hint and the source. Some rules offer a one-click rewrite, and Undo still works after it.
A side panel groups the hits by type. A meter shows hits per 1,000 words or characters, next to a line saying it counts patterns, not authorship.
A compare tab holds a before and an after text and shows how the hit count changed.
Save the annotated text as a PNG.
Open a .txt , .md , .markdown or .mdx file of up to 2 MB with the button or by dropping it on the editor. The file is read in the browser. The web editor scans Markdown as plain text, so use the CLI to skip code blocks.
Nine interface languages: English, Traditional Chinese, Simplified Chinese, Japanese, Korean, Spanish, French, German, Portuguese (BR). Rule messages exist in English and Traditional Chinese.
Simplified Chinese input goes through the Traditional Chinese rules after a character conversion that keeps positions intact.
Web
Use the live demo . Nothing you paste leaves the page. To host it yourself, run npm run build:web and serve dist/web .
Needs Node 22 or newer. CI runs 22 and 24.
npx github:Arthur031221/AiSniff lint draft.md
Or from a checkout: npm install && npm run build:cli && npm link . The npm package name is aisniff-cli , the command is aisniff .
git clone https://github.com/Arthur031221/AiSniff.git
cp -r AiSniff/skills/aisniff ~ / < agent-dir > /skills/aisniff
Then ask: "check this draft with AiSniff". Use <agent-dir>/skills/aisniff inside a project to scope it to that project.
Copy skills/aisniff/SKILL.md into your repository and point to it from AGENTS.md , or paste its body into the agent's instructions. The agent only needs a shell to run npx .
- uses : actions/checkout@v4
- uses : Arthur031221/AiSniff@v0.5.0
with :
paths : docs/
max-weighted-density : 8.7
Where the rules come from
Each rule cites one of six public sources: the Wikipedia guide Signs of AI writing , the pattern lists in blader/humanizer , Humanizer-zh-TW and speak-human-tw , the Chinese Wikipedia page on Europeanized Chinese , and the excess-vocabulary study by Kobak et al., Science Advances 2025 . Those lists are prose guides for rewriting. AiSniff turns their patterns into rules, tests them against examples and measures them on data. It also keeps some rules the data doesn't support, and the evaluation page lists each one. Every rule has passing and failing examples that run as tests, and npm run check:sources checks that each source link and section exists.
Some wording rules mark mainland usage in Traditional Chinese (for example the mainland words for software and video). They flag wording that differs from Taiwan usage. They aren't AI tells.
It finds phrases from published lists. A model that avoids them, or a person who uses them, defeats it.
Formal writing, second-language writing and old-fashioned prose will trigger rules.
English and Chinese only. Other languages are not checked.
The "many patterns" line on the meter is the 99th percentile of human HC3 answers (8.7 and 13.0 weighted hits per 1,000 units). It comes from short Q&A text and is a reference line, not a decision threshold.
Install | Usage | Configuration | Architecture | Writing a rule | Evaluation | Contributing
The Simplified to Traditional table comes from OpenCC (Apache-2.0), see THIRD_PARTY_NOTICES.md .
Paste English or Traditional Chinese text and see which phrases read like AI wrote them, with the rule, source and rewrite hint
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
