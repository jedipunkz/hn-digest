---
source: "https://github.com/formgong/dead-forms-study"
hn_url: "https://news.ycombinator.com/item?id=50524591"
title: "Show HN: We submitted 105 AI-built contact forms – 27 said \"sent\", sent nothing"
article_title: "GitHub - formgong/dead-forms-study: Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results. · GitHub"
image: "https://opengraph.githubassets.com/c2203daf8e1fea72e0e6d9039ddee225405f1eb7b7a7d883359156b937b49083/formgong/dead-forms-study"
author: "Formgong"
captured_at: "2026-11-16T23:42:05Z"
capture_tool: "hn-digest"
hn_id: 50524591
score: 3
comments: 0
posted_at: "2026-11-16T22:47:31Z"
tags:
  - hacker-news
---

# Show HN: We submitted 105 AI-built contact forms – 27 said "sent", sent nothing

- HN: [50524591](https://news.ycombinator.com/item?id=50524591)
- Source: [github.com](https://github.com/formgong/dead-forms-study)
- Score: 3
- Comments: 0
- Posted: 2026-11-16T22:47:31Z

## Translation

Title: Show HN: We submitted 105 AI-built contact forms – 27 said "sent", sent nothing
Article title: GitHub - formgong/dead-forms-study: Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results. · GitHub
Description: Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results. - formgong/dead-forms-study

Article text:
GitHub - formgong/dead-forms-study: Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results. · GitHub
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
formgong
dead-forms-study Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results.
formgong.com/en/blog/contact-forms-that-send-nothing/ Topics
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
fixtures fixtures .gitignore .gitignore LICENSE LICENSE PREREG.md PREREG.md README.md README.md RESULTS.md RESULTS.md fixtures-run.mjs fixtures-run.mjs harness.mjs harness.mjs package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Contact forms that send nothing — study and test harness
We submitted the contact form on public Lovable and Bolt sites while recording every network request and blocking all of them , so no message ever reached a site owner. On the final run, 27 of 105 forms showed a success message ("Message sent!") and sent nothing; 54 of 105 sent the typed message. On sites without demo signs, 16 of 73 were fake.
Write-up: https://formgong.com/en/blog/contact-forms-that-send-nothing/
npm install
npm run fixtures # 23 known-answer pages; all must match
node harness.mjs https://your-site.example --out result.jsonl
Requires Google Chrome (the harness launches it with channel: "chrome" ).
After the contact form is found, every request is intercepted. Same-origin static files and cross-origin scripts, styles and fonts load; everything else is recorded and aborted . The harness fills visible fields with obvious test values (never honeypots), presses send and watches the page for five seconds.
SENDS — a request carried what was typed.
FAKE_SUCCESS — no such request, no other outgoing request, and a success message (or alert() ) appeared.
NO_EFFECT , MAILTO , RELOAD , REDIRECT_NO_SEND , OTHER_REQUEST , INVALID , VOID — see PREREG.md.
Open DevTools → Network and press send: a working form adds a request that carries the text you typed. The write-up's fix section lists three ways to fix it. Coding agents (Claude Code, Cursor, Codex) can install a skill that adds a working form and replaces fake handlers:
npx skills add formgong/skills
Ethics
Nothing is submitted to anyone. The site list and per-site results are not published, and the write-up names no site. Please use the harness the same way.
MIT licensed. Made by Formgong .
Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results.
formgong.com/en/blog/contact-forms-that-send-nothing/ Topics
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results. - formgong/dead-forms-study

GitHub - formgong/dead-forms-study: Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results. · GitHub
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
formgong
dead-forms-study Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results.
formgong.com/en/blog/contact-forms-that-send-nothing/ Topics
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
fixtures fixtures .gitignore .gitignore LICENSE LICENSE PREREG.md PREREG.md README.md README.md RESULTS.md RESULTS.md fixtures-run.mjs fixtures-run.mjs harness.mjs harness.mjs package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Contact forms that send nothing — study and test harness
We submitted the contact form on public Lovable and Bolt sites while recording every network request and blocking all of them , so no message ever reached a site owner. On the final run, 27 of 105 forms showed a success message ("Message sent!") and sent nothing; 54 of 105 sent the typed message. On sites without demo signs, 16 of 73 were fake.
Write-up: https://formgong.com/en/blog/contact-forms-that-send-nothing/
npm install
npm run fixtures # 23 known-answer pages; all must match
node harness.mjs https://your-site.example --out result.jsonl
Requires Google Chrome (the harness launches it with channel: "chrome" ).
After the contact form is found, every request is intercepted. Same-origin static files and cross-origin scripts, styles and fonts load; everything else is recorded and aborted . The harness fills visible fields with obvious test values (never honeypots), presses send and watches the page for five seconds.
SENDS — a request carried what was typed.
FAKE_SUCCESS — no such request, no other outgoing request, and a success message (or alert() ) appeared.
NO_EFFECT , MAILTO , RELOAD , REDIRECT_NO_SEND , OTHER_REQUEST , INVALID , VOID — see PREREG.md.
Open DevTools → Network and press send: a working form adds a request that carries the text you typed. The write-up's fix section lists three ways to fix it. Coding agents (Claude Code, Cursor, Codex) can install a skill that adds a working form and replaces fake handlers:
npx skills add formgong/skills
Ethics
Nothing is submitted to anyone. The site list and per-site results are not published, and the write-up names no site. Please use the harness the same way.
MIT licensed. Made by Formgong .
Do contact forms on AI-built sites send anything? A Playwright harness that checks without sending, with known-answer fixtures and the study's method and results.
formgong.com/en/blog/contact-forms-that-send-nothing/ Topics
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
