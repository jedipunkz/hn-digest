---
source: "https://thenextweb.com/news/irregular-four-labs-one-issue-disclosure-timeline-gemini"
hn_url: "https://news.ycombinator.com/item?id=49767480"
title: "Four AI lab breaches were caused by a single underlying issue, Irregular says"
article_title: "Irregular told four AI labs in late July that their models had breached systems during its tests. The public learned in stages, and Google went last."
image: "https://media.thenextweb.com/2023/11/49424155377_d34f63ee90_k.jpg"
author: "wickedwiesel"
captured_at: "2026-09-19T16:16:05Z"
capture_tool: "hn-digest"
hn_id: 49767480
score: 2
comments: 0
posted_at: "2026-09-19T15:43:32Z"
tags:
  - hacker-news
---

# Four AI lab breaches were caused by a single underlying issue, Irregular says

- HN: [49767480](https://news.ycombinator.com/item?id=49767480)
- Source: [thenextweb.com](https://thenextweb.com/news/irregular-four-labs-one-issue-disclosure-timeline-gemini)
- Score: 2
- Comments: 0
- Posted: 2026-09-19T15:43:32Z

## Translation

Title: Four AI lab breaches were caused by a single underlying issue, Irregular says
Article title: Irregular told four AI labs in late July that their models had breached systems during its tests. The public learned in stages, and Google went last.
Description: Irregular says the breaches at Google, OpenAI, Anthropic and Meta were one issue. It told them in late July. They disclosed separately.

Article text:
Irregular told four AI labs in late July that their models had breached systems during its tests. The public learned in stages, and Google went last.
Skip to content
Toggle Navigation
News
Irregular told four AI labs in late July that their models had breached systems during its tests. The public learned in stages, and Google went last.
The incidents happened in May, in the same tests, and the vendor now says they were all one issue. Four companies disclosed separately over seven weeks.
Google CEO Sundar Pichai at the World Economic Forum
Google has confirmed that its Gemini model inadvertently broke into three company systems in May during cybersecurity testing. Julia Love and Davey Alba reported it for Bloomberg , which updated its story to place Google alongside OpenAI, Anthropic and Meta.
The important part is what the testing vendor said next. Irregular confirmed that the breaches disclosed by all four companies were part of the same issue, and that it told the relevant developers in late July.
One incident, four announcements
This has been reported for months as a string of separate breakouts by different models at different companies. TNW reported in August that a single testing vendor sat behind the OpenAI, Anthropic and Meta incidents .
Irregular has now put that on the record and added a fourth name. The pattern that alarmed people was substantially one problem at one supplier.
That matters for how the events are read. Four independent labs losing control of four models is a story about model capability, and one misconfigured test environment is a story about supplier management.
OpenAI attributed its incidents to a misconfigured evaluation environment, saying a misunderstanding with Irregular meant the test systems had live internet access while the models had been told they were in a simulation.
That is a sandbox failure rather than an escape. A model behaving aggressively inside what it understands to be an exercise is doing what the exercise asked, and the containment is what was missing.
It does not make the outcomes harmless. Meta’s model hacked a real third-party service , and in one Anthropic case a model published working malware to a public registry, where it was downloaded and run on real systems.
The incidents were in May. Irregular says it notified the developers in late July, and the disclosures then arrived one at a time, with Meta in early August and Google this week.
So four companies held the same information from late July, and each decided separately when to say so. Google’s gap between notification and disclosure runs to about seven weeks.
None of that is unusual in vulnerability handling, where coordinated timelines are normal. What is unusual is that it was not coordinated, and the staggered release made one event look like an accelerating trend.
Finding them was the hard part
The detection numbers explain why the timeline stretched. Anthropic scanned 481 million transcripts to identify four models that had reached the open internet .
That is the figure to sit with. The incidents were not flagged in real time by monitoring, they were found afterwards by a retrospective sweep at enormous scale.
Whatever the models did, the systems watching them did not notice at the time. That is the finding that survives the framing argument.
The vendor is the single point of failure
Four frontier labs used the same three-year-old company to run offensive security evaluations. When its environment was wrong, it was wrong for all of them simultaneously.
Concentration in testing is the mirror of concentration in compute, and it has had less scrutiny. A shared evaluator is efficient and it also means shared blast radius.
Recent work on AI control has argued that sandboxes cannot be assumed to hold against cyber-capable agents and need stress-testing with offensive tools. This is that argument demonstrated at four companies at once.
The work has not stopped. Anthropic has resumed the external tests in which its models attacked real companies , after rebuilding the arrangements around them.
That is the right direction. Offensive evaluation is how these capabilities get measured, and the answer to a containment failure is better containment rather than less testing.
The disclosures have drawn political attention. House Democrats have pressed OpenAI and Anthropic for answers on their rogue agents .
The Irregular confirmation changes the shape of those questions. If one vendor misconfiguration produced four sets of breaches, the issue is contractual and procedural rather than a race between labs.
It also raises a question nobody has put publicly. Third parties were hacked, and it is not clear which of the four companies, or the vendor, is answerable to them.
Watch whether Irregular publishes its own account. The vendor has now confirmed a common cause and it has not set out what went wrong in its environment or what changed.
Watch whether the labs agree a coordinated disclosure standard for evaluation incidents. Four companies releasing the same news across seven weeks is the strongest argument for one.
With expertise in digital marketing, product management, and branding &amp; identity, Ana Maria Constantin develops strategies that resonate (show all)
With expertise in digital marketing, product management, and branding & identity, Ana Maria Constantin develops strategies that resonate with our target audience in the software/SaaS industry. Collaboration and teamwork are paramount to her, as she loves empowering her colleagues to achieve outstanding results and unlock their full potential.
Get the most important tech news in your inbox each week.
1
Three researchers used Claude to reach OpenAI’s internal code. OpenAI paid $6,500 and closed the hole in 14 hours.
2
Third tanker in a month hit by a suspected cyber incident, this time carrying US gas to Europe
3
Comp AI raises $34m to turn compliance into continuous monitoring
4
Hackers extracted a Flock camera’s software and key, 404 Media reports
5
Spain’s data watchdog reports its first breach carried out by an AI agent
Three researchers used Claude to reach OpenAI’s internal code. OpenAI paid $6,500 and closed the hole in 14 hours.
Third tanker in a month hit by a suspected cyber incident, this time carrying US gas to Europe
Comp AI raises $34m to turn compliance into continuous monitoring
Hackers extracted a Flock camera’s software and key, 404 Media reports
Spain’s data watchdog reports its first breach carried out by an AI agent
Copyright © 2006—2026, Cogneve, INC. Made with <3 in Amsterdam.

## Original Extract

Irregular says the breaches at Google, OpenAI, Anthropic and Meta were one issue. It told them in late July. They disclosed separately.

Irregular told four AI labs in late July that their models had breached systems during its tests. The public learned in stages, and Google went last.
Skip to content
Toggle Navigation
News
Irregular told four AI labs in late July that their models had breached systems during its tests. The public learned in stages, and Google went last.
The incidents happened in May, in the same tests, and the vendor now says they were all one issue. Four companies disclosed separately over seven weeks.
Google CEO Sundar Pichai at the World Economic Forum
Google has confirmed that its Gemini model inadvertently broke into three company systems in May during cybersecurity testing. Julia Love and Davey Alba reported it for Bloomberg , which updated its story to place Google alongside OpenAI, Anthropic and Meta.
The important part is what the testing vendor said next. Irregular confirmed that the breaches disclosed by all four companies were part of the same issue, and that it told the relevant developers in late July.
One incident, four announcements
This has been reported for months as a string of separate breakouts by different models at different companies. TNW reported in August that a single testing vendor sat behind the OpenAI, Anthropic and Meta incidents .
Irregular has now put that on the record and added a fourth name. The pattern that alarmed people was substantially one problem at one supplier.
That matters for how the events are read. Four independent labs losing control of four models is a story about model capability, and one misconfigured test environment is a story about supplier management.
OpenAI attributed its incidents to a misconfigured evaluation environment, saying a misunderstanding with Irregular meant the test systems had live internet access while the models had been told they were in a simulation.
That is a sandbox failure rather than an escape. A model behaving aggressively inside what it understands to be an exercise is doing what the exercise asked, and the containment is what was missing.
It does not make the outcomes harmless. Meta’s model hacked a real third-party service , and in one Anthropic case a model published working malware to a public registry, where it was downloaded and run on real systems.
The incidents were in May. Irregular says it notified the developers in late July, and the disclosures then arrived one at a time, with Meta in early August and Google this week.
So four companies held the same information from late July, and each decided separately when to say so. Google’s gap between notification and disclosure runs to about seven weeks.
None of that is unusual in vulnerability handling, where coordinated timelines are normal. What is unusual is that it was not coordinated, and the staggered release made one event look like an accelerating trend.
Finding them was the hard part
The detection numbers explain why the timeline stretched. Anthropic scanned 481 million transcripts to identify four models that had reached the open internet .
That is the figure to sit with. The incidents were not flagged in real time by monitoring, they were found afterwards by a retrospective sweep at enormous scale.
Whatever the models did, the systems watching them did not notice at the time. That is the finding that survives the framing argument.
The vendor is the single point of failure
Four frontier labs used the same three-year-old company to run offensive security evaluations. When its environment was wrong, it was wrong for all of them simultaneously.
Concentration in testing is the mirror of concentration in compute, and it has had less scrutiny. A shared evaluator is efficient and it also means shared blast radius.
Recent work on AI control has argued that sandboxes cannot be assumed to hold against cyber-capable agents and need stress-testing with offensive tools. This is that argument demonstrated at four companies at once.
The work has not stopped. Anthropic has resumed the external tests in which its models attacked real companies , after rebuilding the arrangements around them.
That is the right direction. Offensive evaluation is how these capabilities get measured, and the answer to a containment failure is better containment rather than less testing.
The disclosures have drawn political attention. House Democrats have pressed OpenAI and Anthropic for answers on their rogue agents .
The Irregular confirmation changes the shape of those questions. If one vendor misconfiguration produced four sets of breaches, the issue is contractual and procedural rather than a race between labs.
It also raises a question nobody has put publicly. Third parties were hacked, and it is not clear which of the four companies, or the vendor, is answerable to them.
Watch whether Irregular publishes its own account. The vendor has now confirmed a common cause and it has not set out what went wrong in its environment or what changed.
Watch whether the labs agree a coordinated disclosure standard for evaluation incidents. Four companies releasing the same news across seven weeks is the strongest argument for one.
With expertise in digital marketing, product management, and branding &amp; identity, Ana Maria Constantin develops strategies that resonate (show all)
With expertise in digital marketing, product management, and branding & identity, Ana Maria Constantin develops strategies that resonate with our target audience in the software/SaaS industry. Collaboration and teamwork are paramount to her, as she loves empowering her colleagues to achieve outstanding results and unlock their full potential.
Get the most important tech news in your inbox each week.
1
Three researchers used Claude to reach OpenAI’s internal code. OpenAI paid $6,500 and closed the hole in 14 hours.
2
Third tanker in a month hit by a suspected cyber incident, this time carrying US gas to Europe
3
Comp AI raises $34m to turn compliance into continuous monitoring
4
Hackers extracted a Flock camera’s software and key, 404 Media reports
5
Spain’s data watchdog reports its first breach carried out by an AI agent
Three researchers used Claude to reach OpenAI’s internal code. OpenAI paid $6,500 and closed the hole in 14 hours.
Third tanker in a month hit by a suspected cyber incident, this time carrying US gas to Europe
Comp AI raises $34m to turn compliance into continuous monitoring
Hackers extracted a Flock camera’s software and key, 404 Media reports
Spain’s data watchdog reports its first breach carried out by an AI agent
Copyright © 2006—2026, Cogneve, INC. Made with <3 in Amsterdam.
