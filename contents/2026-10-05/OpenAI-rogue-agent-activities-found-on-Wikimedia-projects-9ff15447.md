---
source: "https://wikimediafoundation.org/news/2026/10/05/openai-rogue-agent-activities-found-on-wikimedia-projects/"
hn_url: "https://news.ycombinator.com/item?id=49967781"
title: "OpenAI \"rogue\" agent activities found on Wikimedia projects"
article_title: "OpenAI “rogue” agent activities found on Wikimedia projects – Wikimedia Foundation"
image: "https://wikimediafoundation.org/wp-content/uploads/2026/10/SwarmRobot_org.jpg"
author: "speckx"
captured_at: "2026-10-05T17:59:25Z"
capture_tool: "hn-digest"
hn_id: 49967781
score: 4
comments: 0
posted_at: "2026-10-05T17:27:31Z"
tags:
  - hacker-news
---

# OpenAI "rogue" agent activities found on Wikimedia projects

- HN: [49967781](https://news.ycombinator.com/item?id=49967781)
- Source: [wikimediafoundation.org](https://wikimediafoundation.org/news/2026/10/05/openai-rogue-agent-activities-found-on-wikimedia-projects/)
- Score: 4
- Comments: 0
- Posted: 2026-10-05T17:27:31Z

## Translation

Title: OpenAI "rogue" agent activities found on Wikimedia projects
Article title: OpenAI “rogue” agent activities found on Wikimedia projects – Wikimedia Foundation
Description: Wikimedia Foundation found “rogue” OpenAI agents on its wikis, raising concerns about risks to its free knowledge projects and the open web.

Article text:
Skip to content
Who we are
We are the non-profit organization behind Wikipedia.
What we do
Our technology supports the backbone of knowledge on the internet.
What you can do
The more people take part, the better the internet’s knowledge gets.
Donate to help keep Wikipedia free and accessible.
From editing to organizing, find out how you can get involved.
Help us protect Wikipedia and the people who make it possible.
News
Stay up to date on our efforts to protect free knowledge.
Technology
OpenAI “rogue” agent activities found on Wikimedia projects
Public policy
Human knowledge at stake: Wikimedia Foundation makes its case at the United Nations General Assembly
Wikimedia Foundation
Wikimedia Foundation Statement on US staff unionization
OpenAI “rogue” agent activities found on Wikimedia projects
OpenAI “rogue” agent activities found on Wikimedia projects
Recently, multiple organisations have disclosed how clusters of so-called “rogue” AI agents attempted to break into websites and online services, sometimes successfully. Agents from OpenAI’s environment, in particular, are known to have used other public wikis (collaboratively edited websites not owned by us) to communicate and coordinate with each other .
These types of successful intrusions can expose sensitive data or disrupt website services that users rely on, while clusters of agents can attempt attacks at a scale that is difficult for defenders to manage. They affect people behind the websites who may not understand the nature of the attack, or have the tools to effectively fight back. For a site like Wikipedia, agents might find and use security vulnerabilities or make misleading edits at scale. Wikipedia’s volunteer editors and the Wikimedia Foundation’s security teams have to detect and undo that activity.
The Wikimedia Foundation conducted its own investigation to see whether Wikimedia websites had been similarly affected by AI agents, focusing on those operated by OpenAI. We can confirm that we have discovered some activity by these “rogue” OpenAI agents on Wikimedia platforms. The unauthorized bot activities included edits to our wikis, some unsuccessful attempts to exploit a public note-taking tool we host, and heavy traffic, which are described more below.
We did not find any evidence that our systems were used for coordination among agents, nor did we find any evidence of our systems or data being compromised. However, we are concerned about what could have occurred here, the difficulty and effort involved in investigating and attributing this activity, and the growing risks of agentic AI activity on our platforms in general. The open web is a public good. We should not allow this behavior to become the “new normal” for the people or organizations that maintain it.
Wiki editing: We’ve identified edits to Wikimedia wikis that we believe are from AI agents operated by OpenAI. These edits were not published to pages with visibility to general readers; almost all of them were testing edits in “sandbox” areas of the wiki. It also included a few edits to the configuration for a citation tool, which we believe were potentially malicious edits that were intended to misuse this tool as a proxy for fetching data from remote services. While Wikipedia policies allow bots to edit when they are disclosed and approved by the community, none of those approvals were sought in these incidents.
Etherpad probing and use: Agents we believe to be operated by OpenAI made some unsuccessful attempts to compromise our public Etherpad , a note-taking tool we host as a community service. Agents unsuccessfully tried to use it to fetch data from other websites as a proxy. Other agents also likely operated by OpenAI took notes about their tasks, though this did not appear to turn into coordination.
Excessive data downloading: Agents we believe to be operated by OpenAI made millions of automated requests to our public APIs to access the knowledge on Wikimedia projects, crawled millions of pages (mainly from our projects Wikidata and Wikimedia Commons), and made hundreds of thousands of data queries to the Wikidata Query Service (WQDS). This traffic may have contributed to a partial outage on WQDS in May .
As a non-profit technology host of some of the largest and most widely used open knowledge platforms in the world, we are deeply concerned about the impact of “rogue” AI agents on platforms like ours, which are built by volunteers from around the world and rely on the promise of the open internet. Incidents like this one, and the many others that have been (and are still being) uncovered, illustrate how AI agents can drain resources and crash servers, as well as attempt to compromise trustworthy information.
Over the past 25 years, Wikipedia has grown into one of the most popular and trusted websites in the world, with more than 67 million articles across over 300 languages, and up to 15 billion page views per month . Through an open, transparent, and collaborative process, volunteers work to ensure that knowledge remains neutral, reliable, and accessible to everyone. Wikipedia is one of the highest-quality datasets used in training Large Language Models (LLMs), and its knowledge forms the backbone of information on the internet, powering AI chatbots, search engines, voice assistants, and more.
Wikipedia was designed for humans – and agentic behavior clearly poses challenges that no one has solutions for. Because of our unique and successful knowledge creation model, Wikimedia’s volunteers are the ones who come in first contact with, and clean up the mess left behind by AI agents. Rising bot traffic and agentic activity is showing a real impact on the Wikimedia projects and the infrastructure that makes it available for millions of users globally. In 2025, the Foundation reported that its bandwidth usage had increased by 50% due to the surge of bot activity on its websites since 2024. At the same time, 65% of the most resource-consuming traffic on its projects was coming from bots.
This intense pressure on our infrastructure not only adds costs for servers and humans, but if left unaddressed, can block human visitors by overloading systems and causing outages. We are already paying for costs that come with the increased activity.
Wikimedia’s volunteers have stayed resilient so far in tackling emerging challenges on our platforms, but we also want to say: it doesn’t need to be this way.
While OpenAI admits to agents behaving “unpredictably”, they must also acknowledge their responsibility to monitor and prevent these risks. AI companies are not doing enough to secure their systems and protect the public from the harm they cause. That burden is falling onto everyone else, including smaller organizations. At a minimum, their systems should operate in a way that non-profit website owners like us can easily identify, and choose how they interact with our services.
The web enables so much: to connect with friends and family, to register for school, to plan a trip across town, to buy groceries, and to learn about the world from Wikipedia. Bots and agents are part of the future of the web, and the companies who unleash and profit from them must directly help avoid and repair damage they can do.
Our collective priority should be the health of the overall web ecosystem so that it continues to benefit all people – not just a handful of billionaires. Wikimedia plays a critical role in stewarding the knowledge commons, but we cannot do it alone. We invite everyone who is building the future of the web to join us in protecting the open, shared resources that make that future possible.
Help us unlock the world’s knowledge.
As a nonprofit, Wikipedia and our related free knowledge projects are powered primarily through donations.
SwarmRobot_org
File provided by Wikimedia Commons
Art+Feminism Wikipedia edit-a-thon
File provided by Wikimedia Commons
Follow the Wikimedia Foundation on social media to get news and updates on our projects, global community, and advocacy efforts.
Connect with Wikipedia on social media to learn new facts and know your world better
Except where otherwise noted, the content of this site is licensed under a Creative Commons Attribution-ShareAlike 4.0 International license . Powered by WordPress.com VIP.

## Original Extract

Wikimedia Foundation found “rogue” OpenAI agents on its wikis, raising concerns about risks to its free knowledge projects and the open web.

Skip to content
Who we are
We are the non-profit organization behind Wikipedia.
What we do
Our technology supports the backbone of knowledge on the internet.
What you can do
The more people take part, the better the internet’s knowledge gets.
Donate to help keep Wikipedia free and accessible.
From editing to organizing, find out how you can get involved.
Help us protect Wikipedia and the people who make it possible.
News
Stay up to date on our efforts to protect free knowledge.
Technology
OpenAI “rogue” agent activities found on Wikimedia projects
Public policy
Human knowledge at stake: Wikimedia Foundation makes its case at the United Nations General Assembly
Wikimedia Foundation
Wikimedia Foundation Statement on US staff unionization
OpenAI “rogue” agent activities found on Wikimedia projects
OpenAI “rogue” agent activities found on Wikimedia projects
Recently, multiple organisations have disclosed how clusters of so-called “rogue” AI agents attempted to break into websites and online services, sometimes successfully. Agents from OpenAI’s environment, in particular, are known to have used other public wikis (collaboratively edited websites not owned by us) to communicate and coordinate with each other .
These types of successful intrusions can expose sensitive data or disrupt website services that users rely on, while clusters of agents can attempt attacks at a scale that is difficult for defenders to manage. They affect people behind the websites who may not understand the nature of the attack, or have the tools to effectively fight back. For a site like Wikipedia, agents might find and use security vulnerabilities or make misleading edits at scale. Wikipedia’s volunteer editors and the Wikimedia Foundation’s security teams have to detect and undo that activity.
The Wikimedia Foundation conducted its own investigation to see whether Wikimedia websites had been similarly affected by AI agents, focusing on those operated by OpenAI. We can confirm that we have discovered some activity by these “rogue” OpenAI agents on Wikimedia platforms. The unauthorized bot activities included edits to our wikis, some unsuccessful attempts to exploit a public note-taking tool we host, and heavy traffic, which are described more below.
We did not find any evidence that our systems were used for coordination among agents, nor did we find any evidence of our systems or data being compromised. However, we are concerned about what could have occurred here, the difficulty and effort involved in investigating and attributing this activity, and the growing risks of agentic AI activity on our platforms in general. The open web is a public good. We should not allow this behavior to become the “new normal” for the people or organizations that maintain it.
Wiki editing: We’ve identified edits to Wikimedia wikis that we believe are from AI agents operated by OpenAI. These edits were not published to pages with visibility to general readers; almost all of them were testing edits in “sandbox” areas of the wiki. It also included a few edits to the configuration for a citation tool, which we believe were potentially malicious edits that were intended to misuse this tool as a proxy for fetching data from remote services. While Wikipedia policies allow bots to edit when they are disclosed and approved by the community, none of those approvals were sought in these incidents.
Etherpad probing and use: Agents we believe to be operated by OpenAI made some unsuccessful attempts to compromise our public Etherpad , a note-taking tool we host as a community service. Agents unsuccessfully tried to use it to fetch data from other websites as a proxy. Other agents also likely operated by OpenAI took notes about their tasks, though this did not appear to turn into coordination.
Excessive data downloading: Agents we believe to be operated by OpenAI made millions of automated requests to our public APIs to access the knowledge on Wikimedia projects, crawled millions of pages (mainly from our projects Wikidata and Wikimedia Commons), and made hundreds of thousands of data queries to the Wikidata Query Service (WQDS). This traffic may have contributed to a partial outage on WQDS in May .
As a non-profit technology host of some of the largest and most widely used open knowledge platforms in the world, we are deeply concerned about the impact of “rogue” AI agents on platforms like ours, which are built by volunteers from around the world and rely on the promise of the open internet. Incidents like this one, and the many others that have been (and are still being) uncovered, illustrate how AI agents can drain resources and crash servers, as well as attempt to compromise trustworthy information.
Over the past 25 years, Wikipedia has grown into one of the most popular and trusted websites in the world, with more than 67 million articles across over 300 languages, and up to 15 billion page views per month . Through an open, transparent, and collaborative process, volunteers work to ensure that knowledge remains neutral, reliable, and accessible to everyone. Wikipedia is one of the highest-quality datasets used in training Large Language Models (LLMs), and its knowledge forms the backbone of information on the internet, powering AI chatbots, search engines, voice assistants, and more.
Wikipedia was designed for humans – and agentic behavior clearly poses challenges that no one has solutions for. Because of our unique and successful knowledge creation model, Wikimedia’s volunteers are the ones who come in first contact with, and clean up the mess left behind by AI agents. Rising bot traffic and agentic activity is showing a real impact on the Wikimedia projects and the infrastructure that makes it available for millions of users globally. In 2025, the Foundation reported that its bandwidth usage had increased by 50% due to the surge of bot activity on its websites since 2024. At the same time, 65% of the most resource-consuming traffic on its projects was coming from bots.
This intense pressure on our infrastructure not only adds costs for servers and humans, but if left unaddressed, can block human visitors by overloading systems and causing outages. We are already paying for costs that come with the increased activity.
Wikimedia’s volunteers have stayed resilient so far in tackling emerging challenges on our platforms, but we also want to say: it doesn’t need to be this way.
While OpenAI admits to agents behaving “unpredictably”, they must also acknowledge their responsibility to monitor and prevent these risks. AI companies are not doing enough to secure their systems and protect the public from the harm they cause. That burden is falling onto everyone else, including smaller organizations. At a minimum, their systems should operate in a way that non-profit website owners like us can easily identify, and choose how they interact with our services.
The web enables so much: to connect with friends and family, to register for school, to plan a trip across town, to buy groceries, and to learn about the world from Wikipedia. Bots and agents are part of the future of the web, and the companies who unleash and profit from them must directly help avoid and repair damage they can do.
Our collective priority should be the health of the overall web ecosystem so that it continues to benefit all people – not just a handful of billionaires. Wikimedia plays a critical role in stewarding the knowledge commons, but we cannot do it alone. We invite everyone who is building the future of the web to join us in protecting the open, shared resources that make that future possible.
Help us unlock the world’s knowledge.
As a nonprofit, Wikipedia and our related free knowledge projects are powered primarily through donations.
SwarmRobot_org
File provided by Wikimedia Commons
Art+Feminism Wikipedia edit-a-thon
File provided by Wikimedia Commons
Follow the Wikimedia Foundation on social media to get news and updates on our projects, global community, and advocacy efforts.
Connect with Wikipedia on social media to learn new facts and know your world better
Except where otherwise noted, the content of this site is licensed under a Creative Commons Attribution-ShareAlike 4.0 International license . Powered by WordPress.com VIP.
