---
source: "https://github.com/archestra-ai/archestra/releases/tag/platform-v1.4.0"
hn_url: "https://news.ycombinator.com/item?id=50185318"
title: "Archestra 1.4: OpenAPPA and Self Hosted Runtime for Claude Code and Codex"
article_title: "Release Archestra: 1.4.0 \"Arpa\" · archestra-ai/archestra · GitHub"
image: "https://opengraph.githubassets.com/b2e7f9c5aa7aa2b1cf956bf724d98d99f3bc8af11e477bdc18e5ed52ac043569/archestra-ai/archestra/releases/tag/platform-v1.4.0"
author: "ildari"
captured_at: "2026-10-21T18:25:53Z"
capture_tool: "hn-digest"
hn_id: 50185318
score: 8
comments: 0
posted_at: "2026-10-21T18:21:42Z"
tags:
  - hacker-news
---

# Archestra 1.4: OpenAPPA and Self Hosted Runtime for Claude Code and Codex

- HN: [50185318](https://news.ycombinator.com/item?id=50185318)
- Source: [github.com](https://github.com/archestra-ai/archestra/releases/tag/platform-v1.4.0)
- Score: 8
- Comments: 0
- Posted: 2026-10-21T18:21:42Z

## Translation

Title: Archestra 1.4: OpenAPPA and Self Hosted Runtime for Claude Code and Codex
Article title: Release Archestra: 1.4.0 "Arpa" · archestra-ai/archestra · GitHub
Description: Enterprise AI Platform with guardrails, MCP registry, gateway & orchestrator - Release Archestra: 1.4.0 "Arpa" · archestra-ai/archestra

Article text:
Release Archestra: 1.4.0 "Arpa" · archestra-ai/archestra · GitHub
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
archestra-ai
Filter
Loading
Sorry, something went wrong.
There was an error while loading. Please reload this page .
Immutable
release. Only release title and notes can be modified.
platform-v1.4.0
4113e0a
This commit was created on GitHub.com and signed with GitHub’s verified signature .
GPG key ID: B5690EEEBB952194
Verified
Learn about vigilant mode .
1.4.0 "Arpa" (2026-10-21)
OpenAPPA guardrails replace the old guardrails. 86 PRs. Policy as text, batteries per provider, human review answered from chat, Slack, Teams, Telegram, Claude Code, Codex, OpenCode. Yells (report a bad block; break-the-glass turns guardrails off for the session). Policy validations, GitHub/GitLab sync, consult log. Old tool-invocation/trusted-data policies and Dual LLM are gone. ( #8620 , #8028 , #8711 , #8733 , #8956 , #8543 , #9282 )
RBAC. Per-object grants (View/Use/Edit/Full) replace visibility scopes. Multiple roles, team inheritance, SSO default roles. ( #7807 , #7724 , #8395 )
New Remote Agent Runtime. Hand work off laptop to runtime and back. Files and native conversation survive between turns. Slack gets progress updates and "needs input" pings. ( #7777 , #7951 , #9343 , #9361 )
Connect Page redesign and new clients. Step-by-step flows, script-first installers, OpenCode, Kilo Code, Claude Desktop with plugins and history import, managed instructions that auto-refresh at launch, admin-decided model routing. ( #8467 , #8622 , #8074 , #8825 , #8793 , #8649 , #8768 )
Slack: one bot per agent, streamed replies, native Stop, router bot. Teams keeps context in 1:1 and group chats. ( #8663 , #8959 , #8608 )
Service accounts replace team and org gateway tokens. Auto-migrated, same token values. ( #8928 )
LLM proxy: bill keys to a team with spend caps; cost explorer; several OpenAI-compatible endpoints on one key. ( #8580 , #8809 , #8384 )
Logs: session trace page with MCP tool calls in a waterfall, guardrail checks inline; metadata-only log mode. ( #9146 , #9181 , #8272 )
Git: GitHub and GitLab, cloud and self-hosted, for skills, plugins, policy sync. GitHub Apps only. ( #8739 )
A2A outbound delegation: external agents as subagents. ( #7757 )
Knowledge bases: personal/team sharing, rendered multi-site crawling. ( #7743 , #7771 )
Performance: ~60 perf PRs on page loads, chat cold start, agent lists, cursor pagination.
UI: search palette finds everything, pinning, time range picker, new agent creation flow, new logo.
Legacy guardrails removed; ARCHESTRA_OPENAPPA_ENABLED and ARCHESTRA_AGENT_RUNTIME_ENABLED removed; runtime auto-detects the Agent Sandbox controller. ( #8620 , #8240 )
fix(frontend): stop dialog footer lift from glowing in dark mode by @joeyorlando in #7715
feat(connection): let admins disable skills and LLM proxy on Connect by @piercypixel in #7717
fix(frontend): align identity icons with name inputs by @joeyorlando in #7718
feat(connection): raise Connect skill cap and prepare Desktop marketplace by @piercypixel in #7719
fix(chatops): reply consistently to bare bot mentions by @joeyorlando in #7720
fix(chat): show useful explanations in simplified error cards by @joeyorlando in #7721
feat(connection): let admins disable plugins on Connect by @piercypixel in #7722
feat(releases): isolate qualified stable releases from development by @iskhakov in #7709
chore(main): release platform 1.4.0-beta.1 by @archestra-ci[bot] in #7716
feat(costs): add proxy cost attribution by credential by @joeyorlando in #7723
docs(releases): make stable-line cutovers end-to-end by @piercypixel in #7728
ci(contrib): add anthonyhuynguyen by @archestra-contributor-pr-bot[bot] in #7730
test(releases): add isolated cutover harness by @piercypixel in #7729
feat(rbac): support multiple roles, team inheritance, and SSO defaults by @joeyorlando in #7724
fix(chat): defer Bedrock IAM authentication to the proxy by @Konstantinov-Innokentii in #7732
ci(contrib): add Thebestfishever2l by @archestra-contributor-pr-bot[bot] in #7734
refactor(apps): align settings with team dialogs and move creator into header by @joeyorlando in #7739
fix(agent-runtime): reject incompatible models before accepting runs by @piercypixel in #7740
docs(apps): clarify who can edit shared MCP Apps by @arsenyinfo in #7736
fix(deps): patch platform image security vulnerabilities by @joeyorlando in #7744
fix(agent-runtime): support Claude Code with Bedrock and Vertex AI by @joeyorlando in #7741
feat(knowledge): add personal and team sharing for knowledge bases by @joeyorlando in #7743
chore(deps): bump hono from 4.13.0 to 4.13.7 in /platform/mcp_server_docker_image by @dependabot [bot] in #7746
fix(releases): resume rolling beta release PRs by @joeyorlando in #7749
fix(agent-runtime): preflight incompatible model selections by @piercypixel in #7748
fix(subscriptions): add Astra and simplify credential dialogs by @joeyorlando in #7752
fix(agent-runtime): prevent Hermes startup dependency installs by @joeyorlando in #7755
fix(llm): surface subscription reconnect and model refresh failures by @joeyorlando in #7758
feat(frontend): show my teams and streamline personal settings by @joeyorlando in #7760
chore(main): release platform 1.4.0-beta.2 by @archestra-ci[bot] in #7751
feat(frontend): condense model tables and persist filters in the URL by @joeyorlando in #7762
feat(frontend): add responsive table filters and widen model IDs by @joeyorlando in #7766
fix(mcp): align registry view toggle order by @xdrdak in #7768
feat(frontend): float the save bar on agent, gateway, skill, and plugin edit pages by @joeyorlando in #7769
test(chatops): make binding recency tests deterministic by @joeyorlando in #7773
feat(knowledge): add rendered multi-site crawling and repeatable ingestion by @joeyorlando in #7771
chore(main): release platform 1.4.0-beta.3 by @archestra-ci[bot] in #7765
fix(frontend): update browser titles for chat and connect pages by @joeyorlando in #7776
feat(connection): connect coding agents from a prompt with browser approval by @joeyorlando in #7731
fix(frontend): simplify plugins table and widen visibility by @joeyorlando in #7779
fix(frontend): keep environment warnings visible below page tabs by @joeyorlando in #7784
fix(dev): bound Next.js memory and reduce shared-module reloads by @joeyorlando in #7782
feat(agent-runtime): replace Jobs with durable Agent Sandbox workspaces by @joeyorlando in #7777
fix(frontend): unify collection card styling by @xdrdak in #7786
fix(agents): prompt for all missing runtime credentials from setup links by @joeyorlando in #7788
fix(agent-runtime): preserve session identity and history across continuations by @joeyorlando in #7791
docs: add branch layout diagram to release checklist by @Matvey-Kuk in #7794
chore: consolidate shared agent instructions and skills by @joeyorlando in #7792
chore(main): release platform 1.4.0-beta.4 by @archestra-ci[bot] in #7781
ci(contrib): add parveshsaini by @archestra-contributor-pr-bot[bot] in #7805
perf(frontend): reuse navigation data and prefetch Apps on intent by @iskhakov in #7800
fix(frontend): label app and project visibility badges by @xdrdak in #7803
perf(frontend): defer optional syntax and emoji libraries by @iskhakov in #7801
fix(rbac): one rule set for role writes, so duplicating a role works by @iskhakov in #7808
feat(a2a): add outbound agent delegation by @xdrdak in #7757
feat(connect): install Claude Desktop with subscription reuse by @joeyorlando in #7796
deps: bump the rust-dependencies group in /ai-labs with 4 updates by @dependabot [bot] in #7798
deps: bump taiki-e/install-action from 2.87.0 to 2.87.4 in the github-actions-dependencies group by @dependabot [bot] in #7797
fix(proxy): keep SSE streams alive while tool calls are withheld by @iskhakov in #7810
fix(agent-runtime): survive transient completion monitoring failures by @joeyorlando in #7813
fix(connect): discover Desktop models and remove deprecated MCP configuration by @joeyorlando in #7814
feat(apps): add card quick actions by @xdrdak in #7795
fix(agents): preserve final task reports in Slack by @joeyorlando in #7815
fix(llm-proxy): preserve Model Router virtual key attribution by @xdrdak in #7816
ci: exempt xdrdak from automatic reviewer assignment by @joeyorlando in #7817
fix(database): repair columns skipped during stable-to-beta upgrades by @joeyorlando in #7819
fix(frontend): avoid reporting cancelled Apps requests by @xdrdak in #7824
ci(release): detect incompatible migration histories across release tracks by @joeyorlando in #7821
ci(release): open reviewed backport PRs from labels by @joeyorlando in #7825
fix(observability): preserve route categories on context compaction spans by @joeyorlando in #7827
fix(github-copilot): honor account API endpoints by @joeyorlando in #7828
ci(backport): process requests on main pushes and labels by @joeyorlando in #7831
chore(main): release platform 1.4.0-beta.5 by @archestra-ci[bot] in #7806
test(release): replay released-tag stable upgrades in PR validation by @piercypixel in #7835
fix(frontend): clarify inaccessible log sessions by @joeyorlando in #7839
fix(a2a): authenticate and validate agent connections by @xdrdak in #7836
fix(github-copilot): repair follow-up and nonstreaming inference by @joeyorlando in #7842
chore(main): release platform 1.4.0-beta.6 by @archestra-ci[bot] in #7840
fix(agent-runtime): use a dedicated platform API pod selector by @joeyorlando in #7845
fix(agent-runtime): allow platform Service IPs in runtime egress by @joeyorlando in #7847
feat(agents): expose scoped skills by @xdrdak in #7787
chore(main): release platform 1.4.0-beta.7 by @archestra-ci[bot] in #7846
fix(agents): clarify default agent action by @xdrdak in #7851
fix: recover agent runs from the current messaging thread by @joeyorlando in #7852
fix: include run identifiers in model-visible tool output by @joeyorlando in #7854
fix: use native Claude Code sign-in for personal subscriptions by @joeyorlando in #7850
fix: submit runtime steering and explain terminal cancellation by @joeyorlando in #7857
fix: simplify client connection setup and remove Desktop CLI prerequisites by @joeyorlando in #7856
fix: persist Claude runtime credentials and streamline runtime docs by @joeyorlando in #7858
fix: allow unhealthy staging pods to drain by @joeyorlando in #7859
fix(runtime): preserve native reasoning and recover detached runs by @joeyorlando in #7861
fix(agent-runtime): bound response previews to avoid completion delays by @joeyorlando in #7862
fix(ci): fetch canonical release histories for private PRs by @joeyorlando in #7863
fix(agent-runtime): retry recovery before waiting for completion by @joeyorlando in #7865
fix(agents): reuse canonical selector when adding sub

[truncated]

## Original Extract

Enterprise AI Platform with guardrails, MCP registry, gateway & orchestrator - Release Archestra: 1.4.0 "Arpa" · archestra-ai/archestra

Release Archestra: 1.4.0 "Arpa" · archestra-ai/archestra · GitHub
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
archestra-ai
Filter
Loading
Sorry, something went wrong.
There was an error while loading. Please reload this page .
Immutable
release. Only release title and notes can be modified.
platform-v1.4.0
4113e0a
This commit was created on GitHub.com and signed with GitHub’s verified signature .
GPG key ID: B5690EEEBB952194
Verified
Learn about vigilant mode .
1.4.0 "Arpa" (2026-10-21)
OpenAPPA guardrails replace the old guardrails. 86 PRs. Policy as text, batteries per provider, human review answered from chat, Slack, Teams, Telegram, Claude Code, Codex, OpenCode. Yells (report a bad block; break-the-glass turns guardrails off for the session). Policy validations, GitHub/GitLab sync, consult log. Old tool-invocation/trusted-data policies and Dual LLM are gone. ( #8620 , #8028 , #8711 , #8733 , #8956 , #8543 , #9282 )
RBAC. Per-object grants (View/Use/Edit/Full) replace visibility scopes. Multiple roles, team inheritance, SSO default roles. ( #7807 , #7724 , #8395 )
New Remote Agent Runtime. Hand work off laptop to runtime and back. Files and native conversation survive between turns. Slack gets progress updates and "needs input" pings. ( #7777 , #7951 , #9343 , #9361 )
Connect Page redesign and new clients. Step-by-step flows, script-first installers, OpenCode, Kilo Code, Claude Desktop with plugins and history import, managed instructions that auto-refresh at launch, admin-decided model routing. ( #8467 , #8622 , #8074 , #8825 , #8793 , #8649 , #8768 )
Slack: one bot per agent, streamed replies, native Stop, router bot. Teams keeps context in 1:1 and group chats. ( #8663 , #8959 , #8608 )
Service accounts replace team and org gateway tokens. Auto-migrated, same token values. ( #8928 )
LLM proxy: bill keys to a team with spend caps; cost explorer; several OpenAI-compatible endpoints on one key. ( #8580 , #8809 , #8384 )
Logs: session trace page with MCP tool calls in a waterfall, guardrail checks inline; metadata-only log mode. ( #9146 , #9181 , #8272 )
Git: GitHub and GitLab, cloud and self-hosted, for skills, plugins, policy sync. GitHub Apps only. ( #8739 )
A2A outbound delegation: external agents as subagents. ( #7757 )
Knowledge bases: personal/team sharing, rendered multi-site crawling. ( #7743 , #7771 )
Performance: ~60 perf PRs on page loads, chat cold start, agent lists, cursor pagination.
UI: search palette finds everything, pinning, time range picker, new agent creation flow, new logo.
Legacy guardrails removed; ARCHESTRA_OPENAPPA_ENABLED and ARCHESTRA_AGENT_RUNTIME_ENABLED removed; runtime auto-detects the Agent Sandbox controller. ( #8620 , #8240 )
fix(frontend): stop dialog footer lift from glowing in dark mode by @joeyorlando in #7715
feat(connection): let admins disable skills and LLM proxy on Connect by @piercypixel in #7717
fix(frontend): align identity icons with name inputs by @joeyorlando in #7718
feat(connection): raise Connect skill cap and prepare Desktop marketplace by @piercypixel in #7719
fix(chatops): reply consistently to bare bot mentions by @joeyorlando in #7720
fix(chat): show useful explanations in simplified error cards by @joeyorlando in #7721
feat(connection): let admins disable plugins on Connect by @piercypixel in #7722
feat(releases): isolate qualified stable releases from development by @iskhakov in #7709
chore(main): release platform 1.4.0-beta.1 by @archestra-ci[bot] in #7716
feat(costs): add proxy cost attribution by credential by @joeyorlando in #7723
docs(releases): make stable-line cutovers end-to-end by @piercypixel in #7728
ci(contrib): add anthonyhuynguyen by @archestra-contributor-pr-bot[bot] in #7730
test(releases): add isolated cutover harness by @piercypixel in #7729
feat(rbac): support multiple roles, team inheritance, and SSO defaults by @joeyorlando in #7724
fix(chat): defer Bedrock IAM authentication to the proxy by @Konstantinov-Innokentii in #7732
ci(contrib): add Thebestfishever2l by @archestra-contributor-pr-bot[bot] in #7734
refactor(apps): align settings with team dialogs and move creator into header by @joeyorlando in #7739
fix(agent-runtime): reject incompatible models before accepting runs by @piercypixel in #7740
docs(apps): clarify who can edit shared MCP Apps by @arsenyinfo in #7736
fix(deps): patch platform image security vulnerabilities by @joeyorlando in #7744
fix(agent-runtime): support Claude Code with Bedrock and Vertex AI by @joeyorlando in #7741
feat(knowledge): add personal and team sharing for knowledge bases by @joeyorlando in #7743
chore(deps): bump hono from 4.13.0 to 4.13.7 in /platform/mcp_server_docker_image by @dependabot [bot] in #7746
fix(releases): resume rolling beta release PRs by @joeyorlando in #7749
fix(agent-runtime): preflight incompatible model selections by @piercypixel in #7748
fix(subscriptions): add Astra and simplify credential dialogs by @joeyorlando in #7752
fix(agent-runtime): prevent Hermes startup dependency installs by @joeyorlando in #7755
fix(llm): surface subscription reconnect and model refresh failures by @joeyorlando in #7758
feat(frontend): show my teams and streamline personal settings by @joeyorlando in #7760
chore(main): release platform 1.4.0-beta.2 by @archestra-ci[bot] in #7751
feat(frontend): condense model tables and persist filters in the URL by @joeyorlando in #7762
feat(frontend): add responsive table filters and widen model IDs by @joeyorlando in #7766
fix(mcp): align registry view toggle order by @xdrdak in #7768
feat(frontend): float the save bar on agent, gateway, skill, and plugin edit pages by @joeyorlando in #7769
test(chatops): make binding recency tests deterministic by @joeyorlando in #7773
feat(knowledge): add rendered multi-site crawling and repeatable ingestion by @joeyorlando in #7771
chore(main): release platform 1.4.0-beta.3 by @archestra-ci[bot] in #7765
fix(frontend): update browser titles for chat and connect pages by @joeyorlando in #7776
feat(connection): connect coding agents from a prompt with browser approval by @joeyorlando in #7731
fix(frontend): simplify plugins table and widen visibility by @joeyorlando in #7779
fix(frontend): keep environment warnings visible below page tabs by @joeyorlando in #7784
fix(dev): bound Next.js memory and reduce shared-module reloads by @joeyorlando in #7782
feat(agent-runtime): replace Jobs with durable Agent Sandbox workspaces by @joeyorlando in #7777
fix(frontend): unify collection card styling by @xdrdak in #7786
fix(agents): prompt for all missing runtime credentials from setup links by @joeyorlando in #7788
fix(agent-runtime): preserve session identity and history across continuations by @joeyorlando in #7791
docs: add branch layout diagram to release checklist by @Matvey-Kuk in #7794
chore: consolidate shared agent instructions and skills by @joeyorlando in #7792
chore(main): release platform 1.4.0-beta.4 by @archestra-ci[bot] in #7781
ci(contrib): add parveshsaini by @archestra-contributor-pr-bot[bot] in #7805
perf(frontend): reuse navigation data and prefetch Apps on intent by @iskhakov in #7800
fix(frontend): label app and project visibility badges by @xdrdak in #7803
perf(frontend): defer optional syntax and emoji libraries by @iskhakov in #7801
fix(rbac): one rule set for role writes, so duplicating a role works by @iskhakov in #7808
feat(a2a): add outbound agent delegation by @xdrdak in #7757
feat(connect): install Claude Desktop with subscription reuse by @joeyorlando in #7796
deps: bump the rust-dependencies group in /ai-labs with 4 updates by @dependabot [bot] in #7798
deps: bump taiki-e/install-action from 2.87.0 to 2.87.4 in the github-actions-dependencies group by @dependabot [bot] in #7797
fix(proxy): keep SSE streams alive while tool calls are withheld by @iskhakov in #7810
fix(agent-runtime): survive transient completion monitoring failures by @joeyorlando in #7813
fix(connect): discover Desktop models and remove deprecated MCP configuration by @joeyorlando in #7814
feat(apps): add card quick actions by @xdrdak in #7795
fix(agents): preserve final task reports in Slack by @joeyorlando in #7815
fix(llm-proxy): preserve Model Router virtual key attribution by @xdrdak in #7816
ci: exempt xdrdak from automatic reviewer assignment by @joeyorlando in #7817
fix(database): repair columns skipped during stable-to-beta upgrades by @joeyorlando in #7819
fix(frontend): avoid reporting cancelled Apps requests by @xdrdak in #7824
ci(release): detect incompatible migration histories across release tracks by @joeyorlando in #7821
ci(release): open reviewed backport PRs from labels by @joeyorlando in #7825
fix(observability): preserve route categories on context compaction spans by @joeyorlando in #7827
fix(github-copilot): honor account API endpoints by @joeyorlando in #7828
ci(backport): process requests on main pushes and labels by @joeyorlando in #7831
chore(main): release platform 1.4.0-beta.5 by @archestra-ci[bot] in #7806
test(release): replay released-tag stable upgrades in PR validation by @piercypixel in #7835
fix(frontend): clarify inaccessible log sessions by @joeyorlando in #7839
fix(a2a): authenticate and validate agent connections by @xdrdak in #7836
fix(github-copilot): repair follow-up and nonstreaming inference by @joeyorlando in #7842
chore(main): release platform 1.4.0-beta.6 by @archestra-ci[bot] in #7840
fix(agent-runtime): use a dedicated platform API pod selector by @joeyorlando in #7845
fix(agent-runtime): allow platform Service IPs in runtime egress by @joeyorlando in #7847
feat(agents): expose scoped skills by @xdrdak in #7787
chore(main): release platform 1.4.0-beta.7 by @archestra-ci[bot] in #7846
fix(agents): clarify default agent action by @xdrdak in #7851
fix: recover agent runs from the current messaging thread by @joeyorlando in #7852
fix: include run identifiers in model-visible tool output by @joeyorlando in #7854
fix: use native Claude Code sign-in for personal subscriptions by @joeyorlando in #7850
fix: submit runtime steering and explain terminal cancellation by @joeyorlando in #7857
fix: simplify client connection setup and remove Desktop CLI prerequisites by @joeyorlando in #7856
fix: persist Claude runtime credentials and streamline runtime docs by @joeyorlando in #7858
fix: allow unhealthy staging pods to drain by @joeyorlando in #7859
fix(runtime): preserve native reasoning and recover detached runs by @joeyorlando in #7861
fix(agent-runtime): bound response previews to avoid completion delays by @joeyorlando in #7862
fix(ci): fetch canonical release histories for private PRs by @joeyorlando in #7863
fix(agent-runtime): retry recovery before waiting for completion by @joeyorlando in #7865
fix(agents): reuse canonical selector when adding sub

[truncated]
