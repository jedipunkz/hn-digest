---
source: "https://languageops.com/blog/the-software-exists-because-of-the-developers-agency/"
hn_url: "https://news.ycombinator.com/item?id=49608523"
title: "Developer Agency in the Age of AI"
article_title: "Developer Agency in the Age of AI"
image: "https://languageops.com/images/languageops-og.png"
author: "luxpir"
captured_at: "2026-09-08T11:28:17Z"
capture_tool: "hn-digest"
hn_id: 49608523
score: 1
comments: 0
posted_at: "2026-09-08T10:50:04Z"
tags:
  - hacker-news
---

# Developer Agency in the Age of AI

- HN: [49608523](https://news.ycombinator.com/item?id=49608523)
- Source: [languageops.com](https://languageops.com/blog/the-software-exists-because-of-the-developers-agency/)
- Score: 1
- Comments: 0
- Posted: 2026-09-08T10:50:04Z

## Translation

Title: Developer Agency in the Age of AI
Description: Why LanguageOps is built on decades of experience in infrastructure, open source software, translation, and responsible AI-assisted development.

Article text:
Tools
Live Interpreting Live multilingual audio, captions, terminology, and event reports.
Collaborative Translation Translate in parallel, then review, proofread, and approve together.
Transcreation Service Research-led creative adaptation with meaning groups and human polish.
AV Dubbing & Subtitling Video editing, subtitles, dubbing, batch transcription, and batch TTS.
Alignment Tool Auto-align source and target documents into reusable translation memory.
Term Extraction Turn approved language into a termbase for consistent future work.
OCR Tool Recover editable content from complex PDFs and images.
Large File Processing Translate million-word JSON, Excel, and other large structured files.
Content Studio SEO, competitor research, planning, and AI-assisted content creation.
25 August 2026 | 6 min Read
Developer Agency in the Age of AI
As a bit of background, I have written about performance and security in posts that hit the front page of Hacker News, long before LLMs were in use. I’ve been setting up and securing Linux machines and servers for decades. I was writing in QBasic, a menu-driven “OS”, then simple games in C++ from around 12 (i.e. in the 90s). I am a big proponent of open source (FOSS) 1 and software anyone can run and work on, enabling me to have enterprise capability at no cost or low cost because I learned how to set it all up securely and to run at scale. I’ve brought this knowledge, of course, into the LanguageOps project. This is one area often missing from AI developed projects. People tend to underestimate the effort to operationalize a project.
Then later years spent running a freelance and agency operation of my own in the translation industry, where I’ve worked with big brands on translation, transcreation and even sent out and organised interpreters for days across countries or remote events. I’ve handled file engineering for dozens of agencies, meaning the extraction of translatables, and ensuring the compatibility of various pieces of industry software. My own personal TM (translation memory) is millions of segments long. I’ve done a lot of typing. I’ve done a lot of rewriting other people’s words. Some better than others! I was granted membership of the Chartered Institute of Linguists, as MCIL, after a French degree from living and working in Paris for 7 years.
All of that knowledge is also in the project.
Then more recently I’ve taken to developing with AI assistance. First it helped to write boilerplate faster, autocomplete on steroids, then it took my plans and structured them and autocompleted those, for checking, finally it did all that but following my long instructions file that ensures the entire site context is visible, to not overlap or regress, atomic commits of work to easily roll back, all relevant notes written up after every feature and decision, docs written automatically from that knowledge, unit tests for every feature, now in the thousands, run constantly through development.
I will happily and voluntarily disclose the use of AI in development, but, and I’m sure I’m not alone in this situation, I’d very much like other people to understand what’s behind that decision and why I feel qualified to do so responsibly. The other developer I work with, who has even more decades of experience than me, is a heavy and responsible user of AI-driven development too. We frequently discuss best practices and the most reliable frameworks.
And let’s remember, nearly no developers write machine code. Source code is a layer of abstraction. AI-assisted development adds another layer of abstraction. Yet its output is still code that we can inspect, cross-check, test, profile, audit and reject as if we’d written it ourselves. Responsibility for the finished software remains with the developer.
The software exists because of the developer’s agency, whether they wrote the code directly or directed, reviewed and accepted it.
Experience then lends itself to the security and performance audits then run and built into the development process, on top of using a framework used by NASA, 2 Instagram 3 and many others due to its performance and security. 4 It provides secure defaults and built-in protection against several common vulnerabilities, reducing the number of security decisions that have to be implemented from scratch. Application design, deployment and ongoing auditing still remain our responsibility. Likewise I take the earlier knowledge of server security and optimisation to make sure we’re getting the latest and best features available at any level. What the big players don’t tell you is that most key features around security and performance and development are open source. Much of this essential work is maintained by highly skilled people who receive little or no direct compensation, even when their software becomes critical infrastructure for some of the world’s largest companies. 5 There are sometimes traces of credit given by the enforced accreditation of its use somewhere deep in a README file of an obscure directory, but rest assured open source is nearly universal in modern software development. 6
Larger companies often choose managed cloud platforms, paying a premium for infrastructure, monitoring, engineering support and the ability to scale without operating every layer themselves. They work on the derivative platform offerings, all based on these key pieces of software, with some proprietary bells and whistles for monitoring and menu-driven setup. This all comes out in the price to the end-user. And then major infrastructure providers document several significant outages in a year, with dependent services affected at scale. 7 8
We have chosen to accept more operational responsibility in exchange for portability, architectural simplicity and cost control. That approach will not suit every organisation, but it allows us to avoid unnecessary platform dependencies.
We manage that complexity ourselves by simplifying the architecture and reducing the performance overhead and attack surfaces. We can keep spinning up new instances to match demand. On any remote server in the world. Our load tests have included hundreds of concurrent users, translation memories containing 800,000 entries and files containing two million words. We continue to test the limits and document the conditions under which those results are achieved.
This feeds into our eco stance, with very few wasted “cycles” as they are called, referring to CPU cycles, where a poorly written function can cause 1000s of times more requests to the computer memory, CPU or database than necessary because the smarter filtering or algorithm wasn’t deployed. We strive to run lean to avoid waste at all levels.
A principle that stayed with me from time spent in Canada during my childhood was to take only what you need and leave the rest. It has influenced my preference for lean systems, restrained resource use and infrastructure that scales in response to genuine demand.
So to me this is much more than just a vibe-coded weekend project. It represents years of knowledge, decades of philosophical thought, and months of late nights and early mornings. Its technical scope is also deliberately constrained compared with some of the finance and statistics projects I’ve worked on. That gives us safer ground for constraining the LLM towards good output in development and translation: the system primarily manipulates text and media, with clearly bounded transformation, processing and analysis. The developer’s experience and guidance are as important in software as the linguist’s are in translation. Models can reproduce patterns of taste and judgement from their training data, but they do not possess our product context or bear responsibility for the result.
Free and open-source software , Wikipedia. ↩︎
NASA/TM-2020-220513 , NASA Technical Reports Server. ↩︎
Django overview and sites using Django , Django Software Foundation. ↩︎
Security in Django , Django documentation. ↩︎
2024 World of Open Source Global Spotlight , Linux Foundation Research. ↩︎
World of Open Source Europe Spotlight 2023 , Linux Foundation Research. ↩︎
AWS post-event summaries , Amazon Web Services. ↩︎
Code Orange: Fail Small , Cloudflare. ↩︎
There's plenty more to share, let's keep this going a bit longer!
Subscribe
Next: AGI v1.0 Will Be Harness + Skills
comments powered by Disqus
About Us

## Original Extract

Why LanguageOps is built on decades of experience in infrastructure, open source software, translation, and responsible AI-assisted development.

Tools
Live Interpreting Live multilingual audio, captions, terminology, and event reports.
Collaborative Translation Translate in parallel, then review, proofread, and approve together.
Transcreation Service Research-led creative adaptation with meaning groups and human polish.
AV Dubbing & Subtitling Video editing, subtitles, dubbing, batch transcription, and batch TTS.
Alignment Tool Auto-align source and target documents into reusable translation memory.
Term Extraction Turn approved language into a termbase for consistent future work.
OCR Tool Recover editable content from complex PDFs and images.
Large File Processing Translate million-word JSON, Excel, and other large structured files.
Content Studio SEO, competitor research, planning, and AI-assisted content creation.
25 August 2026 | 6 min Read
Developer Agency in the Age of AI
As a bit of background, I have written about performance and security in posts that hit the front page of Hacker News, long before LLMs were in use. I’ve been setting up and securing Linux machines and servers for decades. I was writing in QBasic, a menu-driven “OS”, then simple games in C++ from around 12 (i.e. in the 90s). I am a big proponent of open source (FOSS) 1 and software anyone can run and work on, enabling me to have enterprise capability at no cost or low cost because I learned how to set it all up securely and to run at scale. I’ve brought this knowledge, of course, into the LanguageOps project. This is one area often missing from AI developed projects. People tend to underestimate the effort to operationalize a project.
Then later years spent running a freelance and agency operation of my own in the translation industry, where I’ve worked with big brands on translation, transcreation and even sent out and organised interpreters for days across countries or remote events. I’ve handled file engineering for dozens of agencies, meaning the extraction of translatables, and ensuring the compatibility of various pieces of industry software. My own personal TM (translation memory) is millions of segments long. I’ve done a lot of typing. I’ve done a lot of rewriting other people’s words. Some better than others! I was granted membership of the Chartered Institute of Linguists, as MCIL, after a French degree from living and working in Paris for 7 years.
All of that knowledge is also in the project.
Then more recently I’ve taken to developing with AI assistance. First it helped to write boilerplate faster, autocomplete on steroids, then it took my plans and structured them and autocompleted those, for checking, finally it did all that but following my long instructions file that ensures the entire site context is visible, to not overlap or regress, atomic commits of work to easily roll back, all relevant notes written up after every feature and decision, docs written automatically from that knowledge, unit tests for every feature, now in the thousands, run constantly through development.
I will happily and voluntarily disclose the use of AI in development, but, and I’m sure I’m not alone in this situation, I’d very much like other people to understand what’s behind that decision and why I feel qualified to do so responsibly. The other developer I work with, who has even more decades of experience than me, is a heavy and responsible user of AI-driven development too. We frequently discuss best practices and the most reliable frameworks.
And let’s remember, nearly no developers write machine code. Source code is a layer of abstraction. AI-assisted development adds another layer of abstraction. Yet its output is still code that we can inspect, cross-check, test, profile, audit and reject as if we’d written it ourselves. Responsibility for the finished software remains with the developer.
The software exists because of the developer’s agency, whether they wrote the code directly or directed, reviewed and accepted it.
Experience then lends itself to the security and performance audits then run and built into the development process, on top of using a framework used by NASA, 2 Instagram 3 and many others due to its performance and security. 4 It provides secure defaults and built-in protection against several common vulnerabilities, reducing the number of security decisions that have to be implemented from scratch. Application design, deployment and ongoing auditing still remain our responsibility. Likewise I take the earlier knowledge of server security and optimisation to make sure we’re getting the latest and best features available at any level. What the big players don’t tell you is that most key features around security and performance and development are open source. Much of this essential work is maintained by highly skilled people who receive little or no direct compensation, even when their software becomes critical infrastructure for some of the world’s largest companies. 5 There are sometimes traces of credit given by the enforced accreditation of its use somewhere deep in a README file of an obscure directory, but rest assured open source is nearly universal in modern software development. 6
Larger companies often choose managed cloud platforms, paying a premium for infrastructure, monitoring, engineering support and the ability to scale without operating every layer themselves. They work on the derivative platform offerings, all based on these key pieces of software, with some proprietary bells and whistles for monitoring and menu-driven setup. This all comes out in the price to the end-user. And then major infrastructure providers document several significant outages in a year, with dependent services affected at scale. 7 8
We have chosen to accept more operational responsibility in exchange for portability, architectural simplicity and cost control. That approach will not suit every organisation, but it allows us to avoid unnecessary platform dependencies.
We manage that complexity ourselves by simplifying the architecture and reducing the performance overhead and attack surfaces. We can keep spinning up new instances to match demand. On any remote server in the world. Our load tests have included hundreds of concurrent users, translation memories containing 800,000 entries and files containing two million words. We continue to test the limits and document the conditions under which those results are achieved.
This feeds into our eco stance, with very few wasted “cycles” as they are called, referring to CPU cycles, where a poorly written function can cause 1000s of times more requests to the computer memory, CPU or database than necessary because the smarter filtering or algorithm wasn’t deployed. We strive to run lean to avoid waste at all levels.
A principle that stayed with me from time spent in Canada during my childhood was to take only what you need and leave the rest. It has influenced my preference for lean systems, restrained resource use and infrastructure that scales in response to genuine demand.
So to me this is much more than just a vibe-coded weekend project. It represents years of knowledge, decades of philosophical thought, and months of late nights and early mornings. Its technical scope is also deliberately constrained compared with some of the finance and statistics projects I’ve worked on. That gives us safer ground for constraining the LLM towards good output in development and translation: the system primarily manipulates text and media, with clearly bounded transformation, processing and analysis. The developer’s experience and guidance are as important in software as the linguist’s are in translation. Models can reproduce patterns of taste and judgement from their training data, but they do not possess our product context or bear responsibility for the result.
Free and open-source software , Wikipedia. ↩︎
NASA/TM-2020-220513 , NASA Technical Reports Server. ↩︎
Django overview and sites using Django , Django Software Foundation. ↩︎
Security in Django , Django documentation. ↩︎
2024 World of Open Source Global Spotlight , Linux Foundation Research. ↩︎
World of Open Source Europe Spotlight 2023 , Linux Foundation Research. ↩︎
AWS post-event summaries , Amazon Web Services. ↩︎
Code Orange: Fail Small , Cloudflare. ↩︎
There's plenty more to share, let's keep this going a bit longer!
Subscribe
Next: AGI v1.0 Will Be Harness + Skills
comments powered by Disqus
About Us
