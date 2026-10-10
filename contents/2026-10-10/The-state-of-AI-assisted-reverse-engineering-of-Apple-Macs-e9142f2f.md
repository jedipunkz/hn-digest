---
source: "https://github.com/aurora-silicon/linux/discussions/70"
hn_url: "https://news.ycombinator.com/item?id=50036511"
title: "The state of AI-assisted reverse-engineering of Apple Macs"
article_title: "🧭 Apple Silicon hardware directory · products & boards · aurora-silicon/linux · Discussion #70 · GitHub"
image: "https://opengraph.githubassets.com/f14569aa041d9a9a4a91d8b66d8673aa284fe1ffb2c66245a2f76ab3015c9026/aurora-silicon/linux/discussions/70"
author: "cromka"
captured_at: "2026-10-10T20:29:42Z"
capture_tool: "hn-digest"
hn_id: 50036511
score: 2
comments: 0
posted_at: "2026-10-10T19:52:08Z"
tags:
  - hacker-news
---

# The state of AI-assisted reverse-engineering of Apple Macs

- HN: [50036511](https://news.ycombinator.com/item?id=50036511)
- Source: [github.com](https://github.com/aurora-silicon/linux/discussions/70)
- Score: 2
- Comments: 0
- Posted: 2026-10-10T19:52:08Z

## Translation

Title: The state of AI-assisted reverse-engineering of Apple Macs
Article title: 🧭 Apple Silicon hardware directory · products & boards · aurora-silicon/linux · Discussion #70 · GitHub
Description: 🧭 Apple Silicon hardware directory · products & boards

Article text:
🧭 Apple Silicon hardware directory · products & boards · aurora-silicon/linux · Discussion #70 · GitHub
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
aurora-silicon
🧭 Apple Silicon hardware directory · products & boards
#70
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
{{actor}} deleted this content
.
{{editor}}'s edit
There was an error while loading. Please reload this page .
Acelogic
Oct 1, 2026
Maintainer
🧭 Apple Silicon hardware directory
Pick a generation. Open your product. Find your board.
25 hardware threads · 59 board configurations · A18 and M1–M6
Each product has one discussion per silicon generation. Screen sizes, port configurations and Pro / Max / Ultra variants live together in its board table.
Quick links: MacBook Neo · A18 · Mac mini · M6 · Mac Studio · M5 Max / Ultra
Aurora Linux is the support baseline. All 25 product discussions / 59 board configurations were refreshed against public PRs, source and named-board reports. Source, feature-branch results and complete-system qualification remain separate.
Aurora Linux · A18–M6 at a glance
59 board configurations · A18 and M1–M6 · Pro / Max / Ultra grouped within their families
↔ Scroll horizontally to compare all families. Capabilities stay in rows; each family has its own column.
Refreshed October 8, 2026: linux/aurora-wip@35c7d87921c0 and m1n1/aurora-wip@049a304e87c4 , with open and feature-branch work identified per product. PR states, destination branches, source contents and follow-up hardware reports were checked; no fresh hardware test was run for this update.
🔵 means implementation is integrated in the pinned source; 🟢 means a cited hardware report for a named board and build. Neither marker automatically qualifies every feature on every sibling board. Open PR results remain development work; merged rows identify the destination and actual source integration. The product matrices above contain exact board exceptions, branch states and source links.
Current development changes · October 8, 2026
A18 / Neo: J700 has native internal DCP/brightness/DPMS reports and automatic Wi-Fi/Bluetooth boot evidence ( linux #203 , linux #204 ). G17 source and open-branch test reports are documented with public-stack and completion/memory limits. SEP session, thermal and suspend source has advanced; physical acceptance remains incomplete.
M1 / M2: USB4 is configured; Thunderbolt, Broadcom recovery, AWDL, haptics, SEP and ANE contributions have merged. ANE runtime PM and exact chip enablement are now documented. M1 Pro and M2 Pro include newer contributor-kernel display, ANE and Touch ID reports without applying them to every board.
M3: core platform and fitted peripheral source remains present; no M3 GPU driver match or full native DCP path was found. Ultra wireless gaps remain. New shared driver merges are not hardware passes.
M4: J616S MacBook Pro now has open m1n1 proxy CPU/input/audio/camera reports in m1n1 #20 ; Linux handoff remains outside their scope. Other M4 assessments retain their source-only or missing-target status.
M5: J813 MacBook Air retains its native internal-panel, ten-core, input and USB3 reports. The relevant m1n1/U-Boot tips are unchanged. J815 and other M5 configurations remain separately unqualified.
M6: J873G Mac mini retains its twelve-core, bounded-frequency and one USB-C console-route reports. Merged linux #194 preserves unmanaged P-state ceilings; its own hardware qualification is pending.
M5 identity mapping: The Apple Wiki currently lists T6050 for M5 Pro, M5 Max and M5 Ultra. This directory retains those marketing variants on each board. The shared catalogue ID is not evidence of identical dies, firmware, drivers or support.
Implementation, a successful build, a hardware test, and a usable complete system are separate milestones. Results on a sibling board do not automatically apply here. Open PRs are development work, not merged-release support.
Report in the product thread. Include the exact board ID, SoC / chip variant, screen or port configuration, software revisions and a repeatable public artifact.
Keep configuration differences in the same thread. Screen size alone does not need a new discussion. Put any differing result beside its tested board ID.
Admin editing: repository administrators and organization owners can edit this directory and every hardware thread through ⋯ → Edit . Community members can reply with evidence.
Shared silicon work: use the soc:Txxxx labels to find affected product threads and cross-link the evidence.
Snapshots are dated. Open and draft PRs are development work. A merged change is not automatic release or whole-system qualification.
Old links remain usable. The former per-board and SoC-index discussions are closed redirects to the consolidated threads; their original posts remain available in a folded archive.
Sources: The Apple Wiki models · M6 · M5 Ultra · Public hardware catalogue · Linux snapshot
Established and consolidated: October 1, 2026
1
You must be logged in to vote
All reactions
Replies:
0 comments
Heading
Bold
Italic
Quote
Code
Link
Numbered list
Unordered list
Task list
Attach files
Mention
Reference
Menu
Heading
There was an error while loading. Please reload this page .
Create a new saved reply
👍
1
reacted with thumbs up emoji
👎
1
reacted with thumbs down emoji
😄
1
reacted with laugh emoji
🎉
1
reacted with hooray emoji
😕
1
reacted with confused emoji
❤️
1
reacted with heart emoji
🚀
1
reacted with rocket emoji
👀
1
reacted with eyes emoji
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

🧭 Apple Silicon hardware directory · products & boards

🧭 Apple Silicon hardware directory · products & boards · aurora-silicon/linux · Discussion #70 · GitHub
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
aurora-silicon
🧭 Apple Silicon hardware directory · products & boards
#70
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
{{actor}} deleted this content
.
{{editor}}'s edit
There was an error while loading. Please reload this page .
Acelogic
Oct 1, 2026
Maintainer
🧭 Apple Silicon hardware directory
Pick a generation. Open your product. Find your board.
25 hardware threads · 59 board configurations · A18 and M1–M6
Each product has one discussion per silicon generation. Screen sizes, port configurations and Pro / Max / Ultra variants live together in its board table.
Quick links: MacBook Neo · A18 · Mac mini · M6 · Mac Studio · M5 Max / Ultra
Aurora Linux is the support baseline. All 25 product discussions / 59 board configurations were refreshed against public PRs, source and named-board reports. Source, feature-branch results and complete-system qualification remain separate.
Aurora Linux · A18–M6 at a glance
59 board configurations · A18 and M1–M6 · Pro / Max / Ultra grouped within their families
↔ Scroll horizontally to compare all families. Capabilities stay in rows; each family has its own column.
Refreshed October 8, 2026: linux/aurora-wip@35c7d87921c0 and m1n1/aurora-wip@049a304e87c4 , with open and feature-branch work identified per product. PR states, destination branches, source contents and follow-up hardware reports were checked; no fresh hardware test was run for this update.
🔵 means implementation is integrated in the pinned source; 🟢 means a cited hardware report for a named board and build. Neither marker automatically qualifies every feature on every sibling board. Open PR results remain development work; merged rows identify the destination and actual source integration. The product matrices above contain exact board exceptions, branch states and source links.
Current development changes · October 8, 2026
A18 / Neo: J700 has native internal DCP/brightness/DPMS reports and automatic Wi-Fi/Bluetooth boot evidence ( linux #203 , linux #204 ). G17 source and open-branch test reports are documented with public-stack and completion/memory limits. SEP session, thermal and suspend source has advanced; physical acceptance remains incomplete.
M1 / M2: USB4 is configured; Thunderbolt, Broadcom recovery, AWDL, haptics, SEP and ANE contributions have merged. ANE runtime PM and exact chip enablement are now documented. M1 Pro and M2 Pro include newer contributor-kernel display, ANE and Touch ID reports without applying them to every board.
M3: core platform and fitted peripheral source remains present; no M3 GPU driver match or full native DCP path was found. Ultra wireless gaps remain. New shared driver merges are not hardware passes.
M4: J616S MacBook Pro now has open m1n1 proxy CPU/input/audio/camera reports in m1n1 #20 ; Linux handoff remains outside their scope. Other M4 assessments retain their source-only or missing-target status.
M5: J813 MacBook Air retains its native internal-panel, ten-core, input and USB3 reports. The relevant m1n1/U-Boot tips are unchanged. J815 and other M5 configurations remain separately unqualified.
M6: J873G Mac mini retains its twelve-core, bounded-frequency and one USB-C console-route reports. Merged linux #194 preserves unmanaged P-state ceilings; its own hardware qualification is pending.
M5 identity mapping: The Apple Wiki currently lists T6050 for M5 Pro, M5 Max and M5 Ultra. This directory retains those marketing variants on each board. The shared catalogue ID is not evidence of identical dies, firmware, drivers or support.
Implementation, a successful build, a hardware test, and a usable complete system are separate milestones. Results on a sibling board do not automatically apply here. Open PRs are development work, not merged-release support.
Report in the product thread. Include the exact board ID, SoC / chip variant, screen or port configuration, software revisions and a repeatable public artifact.
Keep configuration differences in the same thread. Screen size alone does not need a new discussion. Put any differing result beside its tested board ID.
Admin editing: repository administrators and organization owners can edit this directory and every hardware thread through ⋯ → Edit . Community members can reply with evidence.
Shared silicon work: use the soc:Txxxx labels to find affected product threads and cross-link the evidence.
Snapshots are dated. Open and draft PRs are development work. A merged change is not automatic release or whole-system qualification.
Old links remain usable. The former per-board and SoC-index discussions are closed redirects to the consolidated threads; their original posts remain available in a folded archive.
Sources: The Apple Wiki models · M6 · M5 Ultra · Public hardware catalogue · Linux snapshot
Established and consolidated: October 1, 2026
1
You must be logged in to vote
All reactions
Replies:
0 comments
Heading
Bold
Italic
Quote
Code
Link
Numbered list
Unordered list
Task list
Attach files
Mention
Reference
Menu
Heading
There was an error while loading. Please reload this page .
Create a new saved reply
👍
1
reacted with thumbs up emoji
👎
1
reacted with thumbs down emoji
😄
1
reacted with laugh emoji
🎉
1
reacted with hooray emoji
😕
1
reacted with confused emoji
❤️
1
reacted with heart emoji
🚀
1
reacted with rocket emoji
👀
1
reacted with eyes emoji
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
