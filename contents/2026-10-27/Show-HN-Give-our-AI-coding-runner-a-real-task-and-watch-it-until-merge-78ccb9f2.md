---
source: "https://mcptask.online/live?language=en"
hn_url: "https://news.ycombinator.com/item?id=50259111"
title: "Show HN: Give our AI coding runner a real task and watch it until merge"
article_title: "Live – watch the runners build mcp4mail | mcptask.online"
image: "https://mcptask.online/assets/og_live_en@2x-baa9c82c94aa9a5a0a299485555d347b36017b9206993027c59bebd32e202183.jpg"
author: "josef_chmel"
captured_at: "2026-10-27T14:02:28Z"
capture_tool: "hn-digest"
hn_id: 50259111
score: 1
comments: 0
posted_at: "2026-10-27T13:37:38Z"
tags:
  - hacker-news
---

# Show HN: Give our AI coding runner a real task and watch it until merge

- HN: [50259111](https://news.ycombinator.com/item?id=50259111)
- Source: [mcptask.online](https://mcptask.online/live?language=en)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T13:37:38Z

## Translation

Title: Show HN: Give our AI coding runner a real task and watch it until merge
Article title: Live – watch the runners build mcp4mail | mcptask.online
Description: Watch mcptask.online runners build the open-source mcp4mail project right now: every change goes from task to branch, pull request, tests and merge.

Article text:
Features Pricing AI developer How it works From task to merge in six steps Introduction What an AI developer is and how it works on its own Onboard your AI developer Same as any new hire: a computer, accounts, a first task Installation Install mcptask_runner in your project Runner mcptask_runner — autonomous developer Runner docs Features, changelog and CLI reference Integrations Connect with services you already use About us Contact Media Blog EN CS EN SK Login Register Features Pricing AI developer How it works Introduction Onboard your AI developer Installation Runner Runner docs Integrations About us Contact Media Blog CS EN SK Login Register Live · the mcp4mail.online project Watch software being built. Right now.
mcp4mail is a real open-source product. A person writes the tasks, our runners program them — and you can check every step: the task, the pull request, the tests and the resulting code on GitHub.
schedule The site is deployed once a day — a merged change is on it on work days at around 18:00 Prague time (CET/CEST), not the moment you see it merged here.
In 40 days they delivered 179 finished tasks and are now waiting for the next assignment.
Runner 1 · Claude Code waiting for the next task — nothing ready in the queue —
Runner 2 · Claude Code waiting for the next task — nothing ready in the queue —
Runners run on our own machines. None of them has direct access to the main branch.
work_history The runners work on work days from 8:00 until 18:00 Prague time (CET/CEST) and often longer — outside those hours the panel is quiet.
Easy GitHub issue templates: bug report, feature request, "my mail provider does not work" Context: Launch-day task for Show HN (2026-10-21). New visitors will open issues; today .github/ has only dependabot.yml, pull_request_te... Pick this task →
Moderate Self-hosting dry run from a clean clone: fix every gap in docs/self-hosting.md and .env.example Context: Launch-day task for Show HN (2026-10-21). "Run it yourself with docker compose" is a main promise of the README; HN readers will... Pick this task →
Easy README: "Tested mail providers" table built from the existing guides and tests Context: Launch-day task for Show HN (2026-10-21). "Does it work with my provider?" will be a frequent question. The backlog refill track... Pick this task →
Easy README screenshots: the AI-permission switch and the send-approval page Context: Launch-day task for Show HN (2026-10-21). The README explains the safety model only in text; two screenshots make it understanda... Pick this task →
Easy mcp4mail.online: visible links to the GitHub repo and to the live build (homepage + footer, cs/en) Context: Launch-day task for Show HN (2026-10-21). Visitors coming from HN will move between the repo, the hosted site and https://mcptas... Pick this task →
Moderate Mail index: read/unread state goes stale, so "unread since Monday" lists mail already read Stable key: dogfood: unread since Monday (source: Dogfooding, backlog refill 2026-W44). Pick this task →
GitHub issue templates: bug report, feature request, "my mail provider does not work"
Context: Launch-day task for Show HN (2026-10-21). New visitors will open issues; today .github/ has only dependabot.yml , pull_request_template.md and workflows – no issue templates, so reports arrive without the details we need.
Do:
1. Add .github/ISSUE_TEMPLATE/ with issue forms (YAML):
- bug_report.yml : hosted or self-hosted, version/commit, MCP client used, steps, expected vs. actual, relevant log lines (with a reminder never to paste passwords or tokens).
- feature_request.yml : problem, proposed solution, which tool it concerns.
- mail_provider.yml : provider name, IMAP host/port, whether an app password is used, error message shown in the app; link to the existing provider guides.
2. Add config.yml with blank_issues_enabled: false and a contact link to SECURITY.md for vulnerabilities (if SECURITY.md is not merged yet, link to the GitHub security advisories page of the repo).
3. Mention the templates in CONTRIBUTING.md in one sentence.
Acceptance criteria: "New issue" on GitHub shows the three forms; forms validate (GitHub renders them without errors in the PR preview / after merge); bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
Self-hosting dry run from a clean clone: fix every gap in docs/self-hosting.md and .env.example
Context: Launch-day task for Show HN (2026-10-21). "Run it yourself with docker compose" is a main promise of the README; HN readers will try it within minutes and comment if it breaks.
Do:
1. In a fresh clone in a temporary directory, follow README.md → Self-hosting and docs/self-hosting.md literally, step by step, as a newcomer would: cp .env.example .env , fill only what the docs tell you, docker compose up -d , open the app, sign up, reach the Mail accounts page, check that /.well-known/oauth-protected-resource (or whatever metadata endpoint the docs mention) answers with the configured resource URI.
2. Write down every place where the docs were missing a step, were ambiguous, or where .env.example disagrees with the ENV reads in app/ , config/ , lib/ (the "docs drift" check from docs/backlog-refill.md ).
3. Fix the docs and .env.example . If something in the code/compose file is actually broken, fix it only if small; otherwise open a separate bug task in this project and mention it in the PR.
4. Put the list of findings (fixed / filed) into the PR description.
Acceptance criteria: a newcomer following only the docs gets a running instance; .env.example and docs match the code's ENV reads; findings listed in the PR; bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
README: "Tested mail providers" table built from the existing guides and tests
Context: Launch-day task for Show HN (2026-10-21). "Does it work with my provider?" will be a frequent question. The backlog refill tracks IMAP compatibility for Dovecot, Cyrus, Courier, Gmail, iCloud, Seznam, Zoho, Fastmail (guides in app/views/guides/content/ , tests in test/ ; recent tasks #13679 Dovecot, #13680 Seznam, #13681 iCloud), but the README does not show it.
Do:
1. Add a short "Mail providers" table to the README (in or right after the "Mail accounts" section): provider, status ( tested = there is a test in test/ , guide = setup guide exists, untested ), app password needed (yes/no, per the guide), link to the guide on mcp4mail.online if public.
2. Derive every row from the code/tests/guides – do not mark anything tested that has no test. Note Outlook/Exchange status honestly (only if covered by a guide or test; otherwise "untested").
3. Add one line inviting people to report a provider via the issue template (if merged) or a GitHub issue.
Acceptance criteria: table present, each status traceable to a file in the repo (list the files in the PR description), bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
README screenshots: the AI-permission switch and the send-approval page
Context: Launch-day task for Show HN (2026-10-21). The README explains the safety model only in text; two screenshots make it understandable in seconds.
Do:
1. Produce two screenshots from the system tests (or a dedicated system test that saves them), with seeded fake data only – no real addresses:
- Mail accounts page with "Allow the AI to make changes to this mailbox" switched on and the permission checkboxes unfolded.
- The outgoing-message approval page (the page behind the link a user gets before an e-mail from send_message leaves).
2. Save them as optimized PNG/WebP under docs/images/ (keep each under ~200 KB), English UI.
3. Embed them in the README section "MCP endpoint: read-only by design" with descriptive alt text.
Acceptance criteria: two images in docs/images/ , shown in the README on GitHub, no real personal data in them, bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
mcp4mail.online: visible links to the GitHub repo and to the live build (homepage + footer, cs/en)
Context: Launch-day task for Show HN (2026-10-21). Visitors coming from HN will move between the repo, the hosted site and https://mcptask.online/live. The hosted site must make the open-source and built-in-public story easy to find.
Do:
1. Check the homepage and the footer of mcp4mail.online (cs and en). If missing, add: a "Source code on GitHub" link to https://github.com/jchsoft/mcp4mail and a "Watch it being built live" link to https://mcptask.online/live (with ?language=en / ?language=cs according to the current locale). If links already exist, just verify they work in both locales and close the task with a note.
2. Add the i18n strings in cs and en (i18n parity check must pass).
3. Use rel="noopener" for external links; keep the existing design.
Acceptance criteria: both links visible on the homepage and in the footer in cs and en; system test covering their presence; bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
Mail index: read/unread state goes stale, so "unread since Monday" lists mail already read
Stable key: dogfood: unread since Monday (source: Dogfooding, backlog refill 2026-W44).
Imap::MessageSync#import_new_messages only fetches UIDs above MailFolder#last_synced_uid . A message's FLAGS are stored once, when it is first indexed, and never refreshed. When the owner reads (or flags) mail in their own mail client, the index keeps unread: true , so search_messages unread: true returns mail that is no longer unread. The same pass never notices messages expunged on the server, so they stay searchable until get_message fails with "message gone".
Reproduction (live, 2026-10-22)
list_folders for a mailbox: INBOX unseen_count: 0 (live IMAP STATUS).
search_messages account_id: <same>, folder: "INBOX", unread: true : 5 rows, all unread: true .
The two disagree; STATUS is the truth.
In sync_folder , after importing new messages, refresh the flags of already-indexed messages: UID FETCH 1:<last_synced_uid> (FLAGS) (batched like the import), update_all the changed flags , and delete index rows whose UID the server no longer returns. Use CONDSTORE/ CHANGEDSINCE only if it stays simple; a plain FLAGS fetch is cheap enough for the 15-minute refresh. Mirror Gmail's per- gm_msgid copies the way FlagSetter does.
Extend test/services/imap/message_sync_test.rb with the fake IMAP server: a message's flags change between two syncs → the index follows; a message is expunged between two syncs → its row is gone. No setting, no new tool.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
Every change goes through the same five steps. No exceptions.
Afraid AI will break your code? So are we. That is why a runner gets nothing you would not give a new developer on their first day: it works on its own branch and reaches main only through a pull request that passes CI – plus a person's review where the project requires it.
A person writes the brief. The runner takes what is ready — and when a task turns out to be bigger than one run, it sets it aside for a person instead of half-implementing it.
Work happens in isolation. The main branch is locked for people and runners alike.
A small, readable change linked to its task. Anyone can open it on GitHub.
Lint, security scan, unit and system tests. Red means it goes no further.
After green CI the

[truncated]

## Original Extract

Watch mcptask.online runners build the open-source mcp4mail project right now: every change goes from task to branch, pull request, tests and merge.

Features Pricing AI developer How it works From task to merge in six steps Introduction What an AI developer is and how it works on its own Onboard your AI developer Same as any new hire: a computer, accounts, a first task Installation Install mcptask_runner in your project Runner mcptask_runner — autonomous developer Runner docs Features, changelog and CLI reference Integrations Connect with services you already use About us Contact Media Blog EN CS EN SK Login Register Features Pricing AI developer How it works Introduction Onboard your AI developer Installation Runner Runner docs Integrations About us Contact Media Blog CS EN SK Login Register Live · the mcp4mail.online project Watch software being built. Right now.
mcp4mail is a real open-source product. A person writes the tasks, our runners program them — and you can check every step: the task, the pull request, the tests and the resulting code on GitHub.
schedule The site is deployed once a day — a merged change is on it on work days at around 18:00 Prague time (CET/CEST), not the moment you see it merged here.
In 40 days they delivered 179 finished tasks and are now waiting for the next assignment.
Runner 1 · Claude Code waiting for the next task — nothing ready in the queue —
Runner 2 · Claude Code waiting for the next task — nothing ready in the queue —
Runners run on our own machines. None of them has direct access to the main branch.
work_history The runners work on work days from 8:00 until 18:00 Prague time (CET/CEST) and often longer — outside those hours the panel is quiet.
Easy GitHub issue templates: bug report, feature request, "my mail provider does not work" Context: Launch-day task for Show HN (2026-10-21). New visitors will open issues; today .github/ has only dependabot.yml, pull_request_te... Pick this task →
Moderate Self-hosting dry run from a clean clone: fix every gap in docs/self-hosting.md and .env.example Context: Launch-day task for Show HN (2026-10-21). "Run it yourself with docker compose" is a main promise of the README; HN readers will... Pick this task →
Easy README: "Tested mail providers" table built from the existing guides and tests Context: Launch-day task for Show HN (2026-10-21). "Does it work with my provider?" will be a frequent question. The backlog refill track... Pick this task →
Easy README screenshots: the AI-permission switch and the send-approval page Context: Launch-day task for Show HN (2026-10-21). The README explains the safety model only in text; two screenshots make it understanda... Pick this task →
Easy mcp4mail.online: visible links to the GitHub repo and to the live build (homepage + footer, cs/en) Context: Launch-day task for Show HN (2026-10-21). Visitors coming from HN will move between the repo, the hosted site and https://mcptas... Pick this task →
Moderate Mail index: read/unread state goes stale, so "unread since Monday" lists mail already read Stable key: dogfood: unread since Monday (source: Dogfooding, backlog refill 2026-W44). Pick this task →
GitHub issue templates: bug report, feature request, "my mail provider does not work"
Context: Launch-day task for Show HN (2026-10-21). New visitors will open issues; today .github/ has only dependabot.yml , pull_request_template.md and workflows – no issue templates, so reports arrive without the details we need.
Do:
1. Add .github/ISSUE_TEMPLATE/ with issue forms (YAML):
- bug_report.yml : hosted or self-hosted, version/commit, MCP client used, steps, expected vs. actual, relevant log lines (with a reminder never to paste passwords or tokens).
- feature_request.yml : problem, proposed solution, which tool it concerns.
- mail_provider.yml : provider name, IMAP host/port, whether an app password is used, error message shown in the app; link to the existing provider guides.
2. Add config.yml with blank_issues_enabled: false and a contact link to SECURITY.md for vulnerabilities (if SECURITY.md is not merged yet, link to the GitHub security advisories page of the repo).
3. Mention the templates in CONTRIBUTING.md in one sentence.
Acceptance criteria: "New issue" on GitHub shows the three forms; forms validate (GitHub renders them without errors in the PR preview / after merge); bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
Self-hosting dry run from a clean clone: fix every gap in docs/self-hosting.md and .env.example
Context: Launch-day task for Show HN (2026-10-21). "Run it yourself with docker compose" is a main promise of the README; HN readers will try it within minutes and comment if it breaks.
Do:
1. In a fresh clone in a temporary directory, follow README.md → Self-hosting and docs/self-hosting.md literally, step by step, as a newcomer would: cp .env.example .env , fill only what the docs tell you, docker compose up -d , open the app, sign up, reach the Mail accounts page, check that /.well-known/oauth-protected-resource (or whatever metadata endpoint the docs mention) answers with the configured resource URI.
2. Write down every place where the docs were missing a step, were ambiguous, or where .env.example disagrees with the ENV reads in app/ , config/ , lib/ (the "docs drift" check from docs/backlog-refill.md ).
3. Fix the docs and .env.example . If something in the code/compose file is actually broken, fix it only if small; otherwise open a separate bug task in this project and mention it in the PR.
4. Put the list of findings (fixed / filed) into the PR description.
Acceptance criteria: a newcomer following only the docs gets a running instance; .env.example and docs match the code's ENV reads; findings listed in the PR; bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
README: "Tested mail providers" table built from the existing guides and tests
Context: Launch-day task for Show HN (2026-10-21). "Does it work with my provider?" will be a frequent question. The backlog refill tracks IMAP compatibility for Dovecot, Cyrus, Courier, Gmail, iCloud, Seznam, Zoho, Fastmail (guides in app/views/guides/content/ , tests in test/ ; recent tasks #13679 Dovecot, #13680 Seznam, #13681 iCloud), but the README does not show it.
Do:
1. Add a short "Mail providers" table to the README (in or right after the "Mail accounts" section): provider, status ( tested = there is a test in test/ , guide = setup guide exists, untested ), app password needed (yes/no, per the guide), link to the guide on mcp4mail.online if public.
2. Derive every row from the code/tests/guides – do not mark anything tested that has no test. Note Outlook/Exchange status honestly (only if covered by a guide or test; otherwise "untested").
3. Add one line inviting people to report a provider via the issue template (if merged) or a GitHub issue.
Acceptance criteria: table present, each status traceable to a file in the repo (list the files in the PR description), bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
README screenshots: the AI-permission switch and the send-approval page
Context: Launch-day task for Show HN (2026-10-21). The README explains the safety model only in text; two screenshots make it understandable in seconds.
Do:
1. Produce two screenshots from the system tests (or a dedicated system test that saves them), with seeded fake data only – no real addresses:
- Mail accounts page with "Allow the AI to make changes to this mailbox" switched on and the permission checkboxes unfolded.
- The outgoing-message approval page (the page behind the link a user gets before an e-mail from send_message leaves).
2. Save them as optimized PNG/WebP under docs/images/ (keep each under ~200 KB), English UI.
3. Embed them in the README section "MCP endpoint: read-only by design" with descriptive alt text.
Acceptance criteria: two images in docs/images/ , shown in the README on GitHub, no real personal data in them, bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
mcp4mail.online: visible links to the GitHub repo and to the live build (homepage + footer, cs/en)
Context: Launch-day task for Show HN (2026-10-21). Visitors coming from HN will move between the repo, the hosted site and https://mcptask.online/live. The hosted site must make the open-source and built-in-public story easy to find.
Do:
1. Check the homepage and the footer of mcp4mail.online (cs and en). If missing, add: a "Source code on GitHub" link to https://github.com/jchsoft/mcp4mail and a "Watch it being built live" link to https://mcptask.online/live (with ?language=en / ?language=cs according to the current locale). If links already exist, just verify they work in both locales and close the task with a note.
2. Add the i18n strings in cs and en (i18n parity check must pass).
3. Use rel="noopener" for external links; keep the existing design.
Acceptance criteria: both links visible on the homepage and in the footer in cs and en; system test covering their presence; bin/ci green.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
Mail index: read/unread state goes stale, so "unread since Monday" lists mail already read
Stable key: dogfood: unread since Monday (source: Dogfooding, backlog refill 2026-W44).
Imap::MessageSync#import_new_messages only fetches UIDs above MailFolder#last_synced_uid . A message's FLAGS are stored once, when it is first indexed, and never refreshed. When the owner reads (or flags) mail in their own mail client, the index keeps unread: true , so search_messages unread: true returns mail that is no longer unread. The same pass never notices messages expunged on the server, so they stay searchable until get_message fails with "message gone".
Reproduction (live, 2026-10-22)
list_folders for a mailbox: INBOX unseen_count: 0 (live IMAP STATUS).
search_messages account_id: <same>, folder: "INBOX", unread: true : 5 rows, all unread: true .
The two disagree; STATUS is the truth.
In sync_folder , after importing new messages, refresh the flags of already-indexed messages: UID FETCH 1:<last_synced_uid> (FLAGS) (batched like the import), update_all the changed flags , and delete index rows whose UID the server no longer returns. Use CONDSTORE/ CHANGEDSINCE only if it stays simple; a plain FLAGS fetch is cheap enough for the 15-minute refresh. Mirror Gmail's per- gm_msgid copies the way FlagSetter does.
Extend test/services/imap/message_sync_test.rb with the fake IMAP server: a message's flags change between two syncs → the index follows; a message is expunged between two syncs → its row is gone. No setting, no new tool.
Why we ask for your e-mail: we'll e-mail you when the work starts, when the pull request is ready and when it's done. Nothing else, and it is never shown on this page.
Every change goes through the same five steps. No exceptions.
Afraid AI will break your code? So are we. That is why a runner gets nothing you would not give a new developer on their first day: it works on its own branch and reaches main only through a pull request that passes CI – plus a person's review where the project requires it.
A person writes the brief. The runner takes what is ready — and when a task turns out to be bigger than one run, it sets it aside for a person instead of half-implementing it.
Work happens in isolation. The main branch is locked for people and runners alike.
A small, readable change linked to its task. Anyone can open it on GitHub.
Lint, security scan, unit and system tests. Red means it goes no further.
After green CI the

[truncated]
