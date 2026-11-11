---
source: "http://babieslearninglanguage.blogspot.com/2026/11/an-ai-audit-of-my-works-computational.html"
hn_url: "https://news.ycombinator.com/item?id=50455136"
title: "A developmental psychologist audits his own research with AI"
article_title: "Babies Learning Language: An AI audit of my work's computational reproducibility"
image: ""
author: "bluepeter"
captured_at: "2026-11-11T16:56:09Z"
capture_tool: "hn-digest"
hn_id: 50455136
score: 1
comments: 0
posted_at: "2026-11-11T16:18:19Z"
tags:
  - hacker-news
---

# A developmental psychologist audits his own research with AI

- HN: [50455136](https://news.ycombinator.com/item?id=50455136)
- Source: [babieslearninglanguage.blogspot.com](http://babieslearninglanguage.blogspot.com/2026/11/an-ai-audit-of-my-works-computational.html)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T16:18:19Z

## Translation

Title: A developmental psychologist audits his own research with AI
Article title: Babies Learning Language: An AI audit of my work's computational reproducibility
Description: Blog about fatherhood, langauge, developmental psychology, and cognitive science.

Article text:
Babies Learning Language: An AI audit of my work's computational reproducibility
Babies Learning Language
Thoughts on language learning, child development, and fatherhood; experimental methods, reproducibility, and open science; theoretical musings on cognitive science more broadly.
An AI audit of my work's computational reproducibility
(This post follows my previous one about new AI policies for my lab .)
I'm trying to come to terms with the challenges and opportunities of AI for science. Coding agents in particular offer the possibility of "checking" – reproducing and critiquing – scientific data analysis. This possibility is very exciting: AI can help us identify and correct errors in our own work. It can also help us find issues in the published literature, provided that code and data are publicly available. This all felt a little unrealistic even a year ago, but now if you ask a frontier model about the scientific literature on "high effort mode" it may dig into data as a matter of course! With cheaper tokens and computation, I suspect that computational reproducibility and error checking will soon be ubiquitous.
AI aside, open code and data for scientific papers is an ethical imperative . Transparency can help us deal with the fact that irreproducibility and error are unfortunately extremely common in the published literature. For example, around half of all psychology papers have at least one statistical test that – because of typos, rounding error, or other issues – is prima facie impossible (see statcheck ). In addition, computational reproducibility – getting the same numbers from the same raw data – is not achievable for much of the published literature ( one of our papers , new bigger study ). That's why I have made it a lab policy to post code and data.
So did that policy work? And – did open code and data help us avoid errors?
I have been using AI agents (Claude Code, especially with Opus 5.5, which is quite good at this) to do pre-submission checks on manuscripts. I've also had some occasions to use it to reproduce published work and provide critique (outside of the peer review process where it is not always allowed by review agreements). Also, increasingly, AI-driven audits are being used adversarially – including a recent high-profile audit in economics . So, I decided, with some trepidation to turn the lens on my own work.
This is a challenging test, in part because I collaborate broadly and many of the papers I've been involved in are collaborations with folks outside my own lab, often outside of Stanford. Some collaborations are with large consortia.
I turned Claude loose on all the journal papers I have an authorship role on from the last five years. That covers 79 papers. 59 of these were empirical papers – the rest were commentaries, reviews, methods papers, or stage 1 registered reports without data. Here's the Claude skill I used . I have examined some of these reports in detail, but have not checked most of them. Thus, this is a preliminary report at best, since it relies on complex and stochastic AI analyses. Here is what I found out about transparency, reproducibility, error, and privacy.
This whole pattern is still disappointing for me. The causes of the problems were heterogeneous. Many were apparently typos (60/216), but also a lack of random seeds; some others come from cases where the agents were forced to reconstruct analysis because code was not complete. Full and complete pre-publication reproducibility audits would likely have reduced this number significantly. Verdict: not as good as it should be, disappointing.
One coefficient in a secondary analysis was reported backwards so boys were higher on the main measure than girls instead of vice versa (not a claim interpreted anywhere in the paper).
A data coding error led to missing socio-educational status being marked as high-SES for many participants, a set of SES by item effects are thus underestimated by about 25%.
One task of more than a dozen in a multi-study replication has a stimulus coding issue, so a null effect becomes positive.
A bug in a computational simulation for one condition leads to that condition being wrongly interpreted as a hard learning context.
One other upsetting thing that our audits turned up was that my collaborators and I had not been as diligent in privacy protection as I had hoped and expected. In several collaborative studies, a team shared identifiers with the data analysts and they made their way into the repository. In several other cases, prolific IDs (which are anonymous but persistent) were pushed. I believe that all of these have been remediated, and in the case of the identifiers, the appropriate data protection office has been notified. Openness is great, and I love the idea of "born open data" – but humans are fallible! We need automated tools checking our repositories to prevent these errors, much the way various bots police repositories to keep passwords and credentials from being committed. Verdict: Not acceptable; policy changes needed.
Jermdemo November 10, 2026 at 11:15 AM Very cool. Added to https://github.com/leipzig/awesome-reproducible-research#ad-hoc-reproductions
Mixed effects models: Is it time to go Bayesian by default?
Where does logical language come from? The social bootstrapping hypothesis
Preventing statistical reporting errors by integrating writing and coding
▼
2026
(6)
▼
November
(1)
An AI audit of my work's computational reproducibi...

## Original Extract

Blog about fatherhood, langauge, developmental psychology, and cognitive science.

Babies Learning Language: An AI audit of my work's computational reproducibility
Babies Learning Language
Thoughts on language learning, child development, and fatherhood; experimental methods, reproducibility, and open science; theoretical musings on cognitive science more broadly.
An AI audit of my work's computational reproducibility
(This post follows my previous one about new AI policies for my lab .)
I'm trying to come to terms with the challenges and opportunities of AI for science. Coding agents in particular offer the possibility of "checking" – reproducing and critiquing – scientific data analysis. This possibility is very exciting: AI can help us identify and correct errors in our own work. It can also help us find issues in the published literature, provided that code and data are publicly available. This all felt a little unrealistic even a year ago, but now if you ask a frontier model about the scientific literature on "high effort mode" it may dig into data as a matter of course! With cheaper tokens and computation, I suspect that computational reproducibility and error checking will soon be ubiquitous.
AI aside, open code and data for scientific papers is an ethical imperative . Transparency can help us deal with the fact that irreproducibility and error are unfortunately extremely common in the published literature. For example, around half of all psychology papers have at least one statistical test that – because of typos, rounding error, or other issues – is prima facie impossible (see statcheck ). In addition, computational reproducibility – getting the same numbers from the same raw data – is not achievable for much of the published literature ( one of our papers , new bigger study ). That's why I have made it a lab policy to post code and data.
So did that policy work? And – did open code and data help us avoid errors?
I have been using AI agents (Claude Code, especially with Opus 5.5, which is quite good at this) to do pre-submission checks on manuscripts. I've also had some occasions to use it to reproduce published work and provide critique (outside of the peer review process where it is not always allowed by review agreements). Also, increasingly, AI-driven audits are being used adversarially – including a recent high-profile audit in economics . So, I decided, with some trepidation to turn the lens on my own work.
This is a challenging test, in part because I collaborate broadly and many of the papers I've been involved in are collaborations with folks outside my own lab, often outside of Stanford. Some collaborations are with large consortia.
I turned Claude loose on all the journal papers I have an authorship role on from the last five years. That covers 79 papers. 59 of these were empirical papers – the rest were commentaries, reviews, methods papers, or stage 1 registered reports without data. Here's the Claude skill I used . I have examined some of these reports in detail, but have not checked most of them. Thus, this is a preliminary report at best, since it relies on complex and stochastic AI analyses. Here is what I found out about transparency, reproducibility, error, and privacy.
This whole pattern is still disappointing for me. The causes of the problems were heterogeneous. Many were apparently typos (60/216), but also a lack of random seeds; some others come from cases where the agents were forced to reconstruct analysis because code was not complete. Full and complete pre-publication reproducibility audits would likely have reduced this number significantly. Verdict: not as good as it should be, disappointing.
One coefficient in a secondary analysis was reported backwards so boys were higher on the main measure than girls instead of vice versa (not a claim interpreted anywhere in the paper).
A data coding error led to missing socio-educational status being marked as high-SES for many participants, a set of SES by item effects are thus underestimated by about 25%.
One task of more than a dozen in a multi-study replication has a stimulus coding issue, so a null effect becomes positive.
A bug in a computational simulation for one condition leads to that condition being wrongly interpreted as a hard learning context.
One other upsetting thing that our audits turned up was that my collaborators and I had not been as diligent in privacy protection as I had hoped and expected. In several collaborative studies, a team shared identifiers with the data analysts and they made their way into the repository. In several other cases, prolific IDs (which are anonymous but persistent) were pushed. I believe that all of these have been remediated, and in the case of the identifiers, the appropriate data protection office has been notified. Openness is great, and I love the idea of "born open data" – but humans are fallible! We need automated tools checking our repositories to prevent these errors, much the way various bots police repositories to keep passwords and credentials from being committed. Verdict: Not acceptable; policy changes needed.
Jermdemo November 10, 2026 at 11:15 AM Very cool. Added to https://github.com/leipzig/awesome-reproducible-research#ad-hoc-reproductions
Mixed effects models: Is it time to go Bayesian by default?
Where does logical language come from? The social bootstrapping hypothesis
Preventing statistical reporting errors by integrating writing and coding
▼
2026
(6)
▼
November
(1)
An AI audit of my work's computational reproducibi...
