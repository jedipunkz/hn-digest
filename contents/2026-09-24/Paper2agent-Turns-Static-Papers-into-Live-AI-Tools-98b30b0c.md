---
source: "https://spectrum.ieee.org/paper2agent-ai-agents-research-papers"
hn_url: "https://news.ycombinator.com/item?id=49833781"
title: "Paper2agent Turns Static Papers into Live AI Tools"
article_title: "Paper2agent Turns Static Papers Into Live AI Tools - IEEE Spectrum"
image: "https://spectrum.ieee.org/media-library/conceptual-illustration-of-an-ai-chatbot-observing-abstract-patterns-charts-and-shapes-for-its-deep-learning-algorithms.jpg?id=67806327&width=1200&height=600&coordinates=0%2C228%2C0%2C22"
author: "rbanffy"
captured_at: "2026-09-24T18:15:53Z"
capture_tool: "hn-digest"
hn_id: 49833781
score: 2
comments: 0
posted_at: "2026-09-24T17:16:06Z"
tags:
  - hacker-news
---

# Paper2agent Turns Static Papers into Live AI Tools

- HN: [49833781](https://news.ycombinator.com/item?id=49833781)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/paper2agent-ai-agents-research-papers)
- Score: 2
- Comments: 0
- Posted: 2026-09-24T17:16:06Z

## Translation

Title: Paper2agent Turns Static Papers into Live AI Tools
Article title: Paper2agent Turns Static Papers Into Live AI Tools - IEEE Spectrum
Description: Paper2Agent converts complex code and data into tested tools, empowering labs and classrooms to run cutting-edge methods on their own datasets.

Article text:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
Paper2agent Turns Static Papers Into Live AI Tools - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE Why Read a Research Paper When You Can Turn It Into an AI Agent? Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to Spectrum .
Why Read a Research Paper When You Can Turn It Into an AI Agent?
Paper2Agent lets scientists “agentify” their published research for others to query
Elie Dolgin is a science writer specializing in biomedical research.
iStock
Have you ever read a paper in Science or Nature and thought, “Man, that research was so cool. I wish I could try that method on my own data,” only to spend a week wrestling with someone else’s undocumented repo, broken dependencies, and half-finished readme.txt?
Well, now you can, more or less.
Say hello to Paper2Agent, a new open-source framework that transforms academic reports into interactive AI agents you can talk to. Give it a paper, along with the accompanying codebase, data, or other supplementary material, and the system automatically extracts the core workflows, then spins up a tested, runnable toolkit that you can use on your own datasets.
The concept may sound a little like Google’s NotebookLM (now called Gemini Notebook ), which lets you upload documents and chat with an AI about what’s in them. But Paper2Agent aims to go a step further: Rather than simply answering questions about a paper, its agents can actually run the methods described in it—and potentially combine those methods with tools from other papers.
The goal, explains Stanford computer scientist James Zou , is to change what a scientific paper fundamentally is. “Knowledge should not be static records,” Zou says. “It really should be dynamic and interactive—and this has many benefits, including making knowledge more reproducible but also enabling all sorts of new kinds of discovery.”
Zou and his colleagues described the tool 16 September in Nature. They tested Paper2Agent across diverse disciplines including statistics, econometrics, and astrophysics. However, the researchers focused their proof-of-concept demonstrations on computational biology, where turning published methods into usable tools can be particularly cumbersome.
The team started with AlphaGenome, a deep-learning model that predicts how mutations in DNA affect gene regulation. (A companion resource unveiled earlier this month , the AlphaGenome Atlas, cataloged the model’s predictions for all 9 billion possible single-letter changes in the human genome.)
The researchers fed Paper2Agent the corresponding documentation and code. About 45 minutes later, with no human intervention, the AI had produced 22 tools covering different aspects of AlphaGenome’s functionality, all on a personal laptop and for less than US $15 in computing costs. One tool, for example, could predict how a DNA change might affect gene activity, while others could compare those effects across tissues or analyze multiple variants at once.
“The idea of making papers more dynamic and executable through an agentic interface is quite compelling.” —Dongping Chen, University of Maryland
All 22 tools passed automated validation, thanks to a testing agent working behind the scenes to run the various AlphaGenome sub-tools against reference results. If a test failed, the agent would diagnose the problem and try to fix the tool, with up to six attempts per function. If that didn’t work, it could drop the tool altogether.
The validated tools were then packaged into a Model Context Protocol server and connected to Claude Code (though any compatible chat-based AI assistant could do). The result was a user-facing AlphaGenome agent that could take questions in plain English, run the appropriate analyses, and spit back results and visualizations.
The team then put the agent through its paces, hitting it with a battery of questions, ranging from simple requests to open-ended research problems. According to the researchers’ analysis, it outperformed both standard Claude given the AlphaGenome codebase and a specialist AI co-scientist tool called Biomni .
Multi-Agent Research Paper Collaboration
Going one step further, the researchers turned a couple of more papers into interactive agents and linked them with the AlphaGenome agent. (One of the additional papers was on how inherited DNA variants linked to autoimmune disease disrupt cell function, while the other was a more systematic exploration of how silencing every expressed gene alters immune cells.)
Prompted to investigate the genetic basis of psoriasis, an itchy skin disease, the three agents collectively zeroed in on a little-understood gene called GPR137 as a likely causal factor. What’s more, the AI proposed 10 ways to validate this inference. A human researcher selected one, and the resulting analysis found that silencing GPR137 produced changes in gene activity strikingly similar to those caused by the psoriasis-linked variant in immune cells.
“These agents, because they’re able to directly collaborate and communicate, can facilitate all these kinds of collaborations,” says Zou.
Other researchers see plenty of potential as well. “The idea of making papers more dynamic and executable through an agentic interface is quite compelling,” says Dongping Chen , a computer scientist at the University of Maryland in College Park.
“Agentification itself is a useful certificate that says, ‘This work is relatively complete and well documented.’” —James Zou, Stanford University
Olivier Elemento , a computational biologist who directs the Englander Institute for Precision Medicine at Weill Cornell Medicine in New York City, sees the approach as having broader implications for how researchers share their work.
“It’s a real advance in terms of how we think about the publication process,” he says, “with AI at the center and in a way that makes publications more interactive.” (Elemento peer-reviewed the study for Nature .)
The potential applications extend beyond the research side of academia, too. Artur Skowroński , head of application development at the Polish software company VirtusLab, noted in a blog post that Paper2Agent could help bring scientific papers to life in classrooms. For example, students could use the agent to play with methods described in the literature instead of merely reading about them.
The Future of Agentified Research Papers
With Paper2Agent now up and running, Zou and his colleagues have begun turning more of their own research papers into agents. Just one day after publishing their Nature paper on Paper2Agent, they unveiled the Virtual Biotech , a multi-agent platform modeled on a drug development company.
They described the system in Science and, at the same time, posted a Paper2Agent-generated incarnation of the paper.
However, not every study they threw at the tool could be converted into an agent. Of the 100 computational biology papers they tried, 26 failed to make the leap to agent form, often because of incomplete code, missing documentation, or other software packages that couldn’t be made to work.
But Zou sees that as a feature, not necessarily a bug. When the system gets stuck, it can expose missing information, errors in the code, or discrepancies between the paper and its implementation—problems that might otherwise go unnoticed. As Zou puts it: “Agentification itself is a useful certificate that says, ‘This work is relatively complete and well documented.’”
Human scientists, Zou says, will still have the final say. But he envisions agents becoming part of what it means to publish a paper. Today, papers come with data and code availability statements. Tomorrow, he suggests, they could come with an “agent availability” statement: a virtual corresponding author available around the clock, in any language, to answer the questions that real authors never have time to field.
Naturally, Zou and his colleagues decided to try the idea on their own study. They fed the Paper2Agent manuscript into Paper2Agent, creating an agent that now lives at paper2agent.ai . In other words, a paper about turning papers into agents has turned itself into an agent. The recursion , it seems, has already begun.
Will AI Agents Change the Internet Forever? ›
Are You Ready to Let an AI Agent Use Your Computer? ›
Reimagining research papers as interactive and reliable AI agents | Nature ›
Elie Dolgin is a science writer specializing in biomedical research and drug discovery. After a PhD spent studying the population genetics of nematodes, he swapped worms for words—entering journalism as an editor at The Scientist , Nature Medicine , and STAT . Now a freelancer, Elie is a frequent contributor to New Scientist , Nature , IEEE Spectrum , and more.
Lab on a Contact Lens Can Measure Stress Through Serotonin
The Future Is Fanless: 100% Heat Capture for Liquid Cooled AI Servers
The AI Inference Revolution Is Here
Why Andon Labs Puts AI Agents in Charge of Real Businesses
From AI Copilots to Agent Swarms
AI Safety Regulations in the U.S. Could Give Hackers an Edge

## Original Extract

Paper2Agent converts complex code and data into tested tools, empowering labs and classrooms to run cutting-edge methods on their own datasets.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
Paper2agent Turns Static Papers Into Live AI Tools - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE Why Read a Research Paper When You Can Turn It Into an AI Agent? Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to Spectrum .
Why Read a Research Paper When You Can Turn It Into an AI Agent?
Paper2Agent lets scientists “agentify” their published research for others to query
Elie Dolgin is a science writer specializing in biomedical research.
iStock
Have you ever read a paper in Science or Nature and thought, “Man, that research was so cool. I wish I could try that method on my own data,” only to spend a week wrestling with someone else’s undocumented repo, broken dependencies, and half-finished readme.txt?
Well, now you can, more or less.
Say hello to Paper2Agent, a new open-source framework that transforms academic reports into interactive AI agents you can talk to. Give it a paper, along with the accompanying codebase, data, or other supplementary material, and the system automatically extracts the core workflows, then spins up a tested, runnable toolkit that you can use on your own datasets.
The concept may sound a little like Google’s NotebookLM (now called Gemini Notebook ), which lets you upload documents and chat with an AI about what’s in them. But Paper2Agent aims to go a step further: Rather than simply answering questions about a paper, its agents can actually run the methods described in it—and potentially combine those methods with tools from other papers.
The goal, explains Stanford computer scientist James Zou , is to change what a scientific paper fundamentally is. “Knowledge should not be static records,” Zou says. “It really should be dynamic and interactive—and this has many benefits, including making knowledge more reproducible but also enabling all sorts of new kinds of discovery.”
Zou and his colleagues described the tool 16 September in Nature. They tested Paper2Agent across diverse disciplines including statistics, econometrics, and astrophysics. However, the researchers focused their proof-of-concept demonstrations on computational biology, where turning published methods into usable tools can be particularly cumbersome.
The team started with AlphaGenome, a deep-learning model that predicts how mutations in DNA affect gene regulation. (A companion resource unveiled earlier this month , the AlphaGenome Atlas, cataloged the model’s predictions for all 9 billion possible single-letter changes in the human genome.)
The researchers fed Paper2Agent the corresponding documentation and code. About 45 minutes later, with no human intervention, the AI had produced 22 tools covering different aspects of AlphaGenome’s functionality, all on a personal laptop and for less than US $15 in computing costs. One tool, for example, could predict how a DNA change might affect gene activity, while others could compare those effects across tissues or analyze multiple variants at once.
“The idea of making papers more dynamic and executable through an agentic interface is quite compelling.” —Dongping Chen, University of Maryland
All 22 tools passed automated validation, thanks to a testing agent working behind the scenes to run the various AlphaGenome sub-tools against reference results. If a test failed, the agent would diagnose the problem and try to fix the tool, with up to six attempts per function. If that didn’t work, it could drop the tool altogether.
The validated tools were then packaged into a Model Context Protocol server and connected to Claude Code (though any compatible chat-based AI assistant could do). The result was a user-facing AlphaGenome agent that could take questions in plain English, run the appropriate analyses, and spit back results and visualizations.
The team then put the agent through its paces, hitting it with a battery of questions, ranging from simple requests to open-ended research problems. According to the researchers’ analysis, it outperformed both standard Claude given the AlphaGenome codebase and a specialist AI co-scientist tool called Biomni .
Multi-Agent Research Paper Collaboration
Going one step further, the researchers turned a couple of more papers into interactive agents and linked them with the AlphaGenome agent. (One of the additional papers was on how inherited DNA variants linked to autoimmune disease disrupt cell function, while the other was a more systematic exploration of how silencing every expressed gene alters immune cells.)
Prompted to investigate the genetic basis of psoriasis, an itchy skin disease, the three agents collectively zeroed in on a little-understood gene called GPR137 as a likely causal factor. What’s more, the AI proposed 10 ways to validate this inference. A human researcher selected one, and the resulting analysis found that silencing GPR137 produced changes in gene activity strikingly similar to those caused by the psoriasis-linked variant in immune cells.
“These agents, because they’re able to directly collaborate and communicate, can facilitate all these kinds of collaborations,” says Zou.
Other researchers see plenty of potential as well. “The idea of making papers more dynamic and executable through an agentic interface is quite compelling,” says Dongping Chen , a computer scientist at the University of Maryland in College Park.
“Agentification itself is a useful certificate that says, ‘This work is relatively complete and well documented.’” —James Zou, Stanford University
Olivier Elemento , a computational biologist who directs the Englander Institute for Precision Medicine at Weill Cornell Medicine in New York City, sees the approach as having broader implications for how researchers share their work.
“It’s a real advance in terms of how we think about the publication process,” he says, “with AI at the center and in a way that makes publications more interactive.” (Elemento peer-reviewed the study for Nature .)
The potential applications extend beyond the research side of academia, too. Artur Skowroński , head of application development at the Polish software company VirtusLab, noted in a blog post that Paper2Agent could help bring scientific papers to life in classrooms. For example, students could use the agent to play with methods described in the literature instead of merely reading about them.
The Future of Agentified Research Papers
With Paper2Agent now up and running, Zou and his colleagues have begun turning more of their own research papers into agents. Just one day after publishing their Nature paper on Paper2Agent, they unveiled the Virtual Biotech , a multi-agent platform modeled on a drug development company.
They described the system in Science and, at the same time, posted a Paper2Agent-generated incarnation of the paper.
However, not every study they threw at the tool could be converted into an agent. Of the 100 computational biology papers they tried, 26 failed to make the leap to agent form, often because of incomplete code, missing documentation, or other software packages that couldn’t be made to work.
But Zou sees that as a feature, not necessarily a bug. When the system gets stuck, it can expose missing information, errors in the code, or discrepancies between the paper and its implementation—problems that might otherwise go unnoticed. As Zou puts it: “Agentification itself is a useful certificate that says, ‘This work is relatively complete and well documented.’”
Human scientists, Zou says, will still have the final say. But he envisions agents becoming part of what it means to publish a paper. Today, papers come with data and code availability statements. Tomorrow, he suggests, they could come with an “agent availability” statement: a virtual corresponding author available around the clock, in any language, to answer the questions that real authors never have time to field.
Naturally, Zou and his colleagues decided to try the idea on their own study. They fed the Paper2Agent manuscript into Paper2Agent, creating an agent that now lives at paper2agent.ai . In other words, a paper about turning papers into agents has turned itself into an agent. The recursion , it seems, has already begun.
Will AI Agents Change the Internet Forever? ›
Are You Ready to Let an AI Agent Use Your Computer? ›
Reimagining research papers as interactive and reliable AI agents | Nature ›
Elie Dolgin is a science writer specializing in biomedical research and drug discovery. After a PhD spent studying the population genetics of nematodes, he swapped worms for words—entering journalism as an editor at The Scientist , Nature Medicine , and STAT . Now a freelancer, Elie is a frequent contributor to New Scientist , Nature , IEEE Spectrum , and more.
Lab on a Contact Lens Can Measure Stress Through Serotonin
The Future Is Fanless: 100% Heat Capture for Liquid Cooled AI Servers
The AI Inference Revolution Is Here
Why Andon Labs Puts AI Agents in Charge of Real Businesses
From AI Copilots to Agent Swarms
AI Safety Regulations in the U.S. Could Give Hackers an Edge
