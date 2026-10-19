---
source: "https://arxiv.org/abs/2610.19264"
hn_url: "https://news.ycombinator.com/item?id=50158954"
title: "Profiling the Attack Surface of AI Agent Skills"
article_title: "[2610.19264] Beyond Maliciousness: Static and Runtime Attack-Surface Profiling of Agent Skills"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "ratneshiiita"
captured_at: "2026-10-19T23:21:38Z"
capture_tool: "hn-digest"
hn_id: 50158954
score: 1
comments: 0
posted_at: "2026-10-19T22:33:35Z"
tags:
  - hacker-news
---

# Profiling the Attack Surface of AI Agent Skills

- HN: [50158954](https://news.ycombinator.com/item?id=50158954)
- Source: [arxiv.org](https://arxiv.org/abs/2610.19264)
- Score: 1
- Comments: 0
- Posted: 2026-10-19T22:33:35Z

## Translation

Title: Profiling the Attack Surface of AI Agent Skills
Article title: [2610.19264] Beyond Maliciousness: Static and Runtime Attack-Surface Profiling of Agent Skills
Description: Abstract page for arXiv paper 2610.19264: Beyond Maliciousness: Static and Runtime Attack-Surface Profiling of Agent Skills

Article text:
Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 13 Oct 2026]
Title: Beyond Maliciousness: Static and Runtime Attack-Surface Profiling of Agent Skills
Abstract: Agent skills are Markdown documents that a tool-using agent reads as instructions and then acts on with the privileges of the session that loaded them. Installing one grants authority (the standing to direct the agent, not new privilege), and the size of that grant is what we measure. We present a two-phase attack-surface profiling methodology that combines whole-bundle static analysis and dependency provenance checking with instrumented sandbox execution and separate model-based assessment of the resulting traces. We apply the methodology to 6,309 listings from two marketplaces, selecting listings for execution without conditioning that decision on a prior threat or risk signal. Half the corpus carries evidence of shell-execution capability; the corresponding shares are 35% for filesystem read, 31% each for filesystem write and network egress, and 13% for credential access; seven listings in ten carry at least one elevated capability. That is the surface a user accepts at install time, and almost none of it is visible then. A permission model exists but is close to inert. Authors may list the tools a skill pre-approves in its own frontmatter, which removes confirmation prompts rather than restricting what the skill can reach; of the 6,061 listings whose frontmatter parses, 770 declare anything at all, the vocabulary cannot express credential or environment access, and nothing checks what it does express against the bundle. Declaring is also not declaring accurately: 59% of declaring listings understate what their bundle can do. ClawHub declares roughly twice as often as skillsmp, yet under-declares on 88% of its declarations against 31%. Triage at the marketplace boundary therefore cannot rest on what a listing declares. It must rest on what the bundle can do. These results position capability-oriented attack-surface profiling as a necessary complement to maliciousness detection.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

## Original Extract

Abstract page for arXiv paper 2610.19264: Beyond Maliciousness: Static and Runtime Attack-Surface Profiling of Agent Skills

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 13 Oct 2026]
Title: Beyond Maliciousness: Static and Runtime Attack-Surface Profiling of Agent Skills
Abstract: Agent skills are Markdown documents that a tool-using agent reads as instructions and then acts on with the privileges of the session that loaded them. Installing one grants authority (the standing to direct the agent, not new privilege), and the size of that grant is what we measure. We present a two-phase attack-surface profiling methodology that combines whole-bundle static analysis and dependency provenance checking with instrumented sandbox execution and separate model-based assessment of the resulting traces. We apply the methodology to 6,309 listings from two marketplaces, selecting listings for execution without conditioning that decision on a prior threat or risk signal. Half the corpus carries evidence of shell-execution capability; the corresponding shares are 35% for filesystem read, 31% each for filesystem write and network egress, and 13% for credential access; seven listings in ten carry at least one elevated capability. That is the surface a user accepts at install time, and almost none of it is visible then. A permission model exists but is close to inert. Authors may list the tools a skill pre-approves in its own frontmatter, which removes confirmation prompts rather than restricting what the skill can reach; of the 6,061 listings whose frontmatter parses, 770 declare anything at all, the vocabulary cannot express credential or environment access, and nothing checks what it does express against the bundle. Declaring is also not declaring accurately: 59% of declaring listings understate what their bundle can do. ClawHub declares roughly twice as often as skillsmp, yet under-declares on 88% of its declarations against 31%. Triage at the marketplace boundary therefore cannot rest on what a listing declares. It must rest on what the bundle can do. These results position capability-oriented attack-surface profiling as a necessary complement to maliciousness detection.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
