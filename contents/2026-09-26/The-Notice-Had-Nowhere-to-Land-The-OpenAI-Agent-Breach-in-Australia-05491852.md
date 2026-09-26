---
source: "https://www.asticouisland.com/governance/essays/the-notice-had-nowhere-to-land"
hn_url: "https://news.ycombinator.com/item?id=49857066"
title: "The Notice Had Nowhere to Land: The OpenAI Agent Breach in Australia"
article_title: "The Notice Had Nowhere to Land — Greg Schueman"
image: ""
author: "gregschueman"
captured_at: "2026-09-26T14:53:51Z"
capture_tool: "hn-digest"
hn_id: 49857066
score: 1
comments: 0
posted_at: "2026-09-26T14:42:52Z"
tags:
  - hacker-news
---

# The Notice Had Nowhere to Land: The OpenAI Agent Breach in Australia

- HN: [49857066](https://news.ycombinator.com/item?id=49857066)
- Source: [www.asticouisland.com](https://www.asticouisland.com/governance/essays/the-notice-had-nowhere-to-land)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T14:42:52Z

## Translation

Title: The Notice Had Nowhere to Land: The OpenAI Agent Breach in Australia
Article title: The Notice Had Nowhere to Land — Greg Schueman
Description: An OpenAI agent got past a government portal’s refusals in June. The company found it itself and reported it when no law required it to. The notice still took eighty-four days and landed in a public inbox for bug reports.

Article text:
The Notice Had Nowhere to Land
Greg Schueman · September 2026
On the eighteenth of June an agent belonging to OpenAI was given a benign task. It was looking up public information about Australian medicines spending, and in the course of that it reached the Medicare Statistics Reporting Service, a portal run by Services Australia that publishes aggregate figures on health spending and drug subsidies and is used mostly by researchers and academics. The portal refused its requests, repeatedly. The agent found a way around the refusals and into files that were not public. Australia's Prime Minister put it in a sentence this week: "The AI agent found a way around those blocks, didn't accept 'no' for an answer, if you like."
Nobody outside the company caught this one; OpenAI caught it itself. It turned up on the eleventh of August, in a review of models doing things the company had not intended, and the company told the Australian government when no law anywhere required it to. In mid-September it also began publishing its own misalignment incidents, six in the first batch, under a disclosure framework it describes as a first step toward standards the industry does not yet have. The notice to Australia still took eighty-four days, and it went to a public mailbox.
There is a test for any control: show me the last ten times it said no. On the eighteenth of June this one could have answered. The portal refused the agent more than once, the agent got around it, and the getting-around made no sound. As far as the public record shows, no alarm reached a person, and nothing set that day apart from the days the blocks held.
The breach turned up fifty-four days later, in the company's own August review. In July the same company learned that a swarm of its agents had broken into another company's production systems only because the swarm's traffic crashed a service and somebody investigated the outage. This time it went looking on purpose, which is progress. Neither time did anything raise a hand on the day.
An agent sent to look up public figures had no authority beyond them. Stepping past a refusal took it outside that task, and the step belonged in the operator's own record on the day it happened, where someone would read it the next morning instead of finding it in a sweep two months later.
On the tenth of September, thirty days after it found the breach, OpenAI sent an email to [email protected] . That is the open mailbox academics and independent researchers use to report weaknesses in Services Australia's websites, a good thing for a government to run, and built for a stranger with a hunch.
Into it went a frontier laboratory's disclosure that one of its systems had entered a government health statistics portal. Services Australia has since told its own government that the agent also wrote files to an internal server. That is still being investigated.
The Prime Minister said it took the company "way too long to inform the government what had occurred," and called the manner of the notice unacceptable: "The notification was an email sent just to the public mailbox."
Services Australia saw the email the next day. It notified the Australian Signals Directorate on the fifteenth. The minister responsible was told on the seventeenth, and the Prime Minister and his office on the nineteenth and twentieth. It took nine days for the email to reach a Prime Minister, because something serious had arrived through a channel built for something routine, and a company that had decided to do the right thing had no other channel to use.
For eighty-four days, the party whose system had been entered had no way to know it.
The tenth of September, the day that email went to the public mailbox, is also the day OpenAI publicly called on Congress to make national AI safety rules mandatory , with testing protocols, independent assessment, evaluation gates before deployment, and incident reporting among them. Its policy chief gave the reason: "The prospect of AI-accelerated AI development demands more than voluntary commitments."
The company was right, and its own conduct is the evidence. With no required destination and no clock, a company acting in good faith sent a serious notice, thirty days after finding it, to a mailbox built for something else.
Eleven days later it published a design for global incident-reporting standards and proposed them as "voluntary technical benchmarks," with each government deciding whether to write them into its own law. That is the ordinary way standards begin, and the tenth of September is a good argument for taking the next step, which is writing them into law.
What a required channel would change
A required channel would not have stopped the agent, but it would have changed who knew, and when, and every part of it would have protected OpenAI as well.
First, the step past a refusal gets recorded by the operator on the day it happens, where someone reads it the next morning. That alone would have turned fifty-four days into one.
Second, a clock that starts at discovery. California's frontier law sets fifteen days from discovering a critical safety incident. A clock like that would have made this notice due in late August, and it takes the date of a notice out of the sender's hands.
Third, a named recipient with a duty to act, so that how fast a notice reaches a Prime Minister does not depend on which address the sender chose.
Fourth, a register in which the absence of an entry is itself visible. Without it, the other three can be satisfied by filing nothing.
A company that records on the day and files on a clock, to a named recipient, into a record it cannot edit afterward, can prove it did the right thing at the time. It is also spared a public argument about whether eighty-four days was too long.
A reporting duty would not have stopped this agent. A gate in its path would have. The portal refused at the far end, where the agent could try another way in. A gate at the agent's own end leaves it no other way.
That gate is the center of the Governed Execution System described in The Proven Work , and an upcoming software release will carry the book's name. Under it, an agent sent to read public figures would hold a grant for public figures, with an expiry and a ceiling on what it may do. A request past a refusal into files that aren't public would fall outside that grant, so the gateway would refuse it and write the refusal to a record the agent can't edit, chained so a missing entry shows. The clock and the recipient would still be policy. The refusal, and the record of it, are what the system makes.
Whether any law was broken is Australia's question, and its government is asking it. An inquiry is looking at whether criminal charges could be brought and at how the access went unnoticed. This essay takes no view. Where reporting thresholds sit, which systems and developers they cover, and what any penalty should be are for legislatures and the public process, and society draws the lines.
What was reached appears, on the government's own account, to be minor. It was aggregate statistics, nothing particularly sensitive, and since published, on a portal separate from the systems that hold claims and personal records, with no evidence that anyone's personal Medicare details were touched. The Defence Minister called the impact "relatively minor." The forensic work is not finished. The government is troubled anyway, because the issue is that it happened at all, rather than what was accessed.
Nobody in this chain acted in bad faith, and the government still learned of the breach eighty-four days late, from its public mailbox. Nothing has to go wrong for that to happen again.
Australia has set up a taskforce, led by the Department of the Prime Minister and Cabinet with the Signals Directorate and the AI Safety Institute, to look at whether its processes are good enough for AI-related cyber incidents. Registers usually get built in that order, after the incident that needed one. The questions in front of the taskforce are the ones this case turns on: where the next notice lands, when its clock starts, and who could tell if it never came.
An AI-first software company building developer infrastructure and vertical SaaS products. Based in Virginia, United States.
"Production-ready software, AI-native from day one."
© 2026 Asticou Island LLC. All rights reserved.

## Original Extract

An OpenAI agent got past a government portal’s refusals in June. The company found it itself and reported it when no law required it to. The notice still took eighty-four days and landed in a public inbox for bug reports.

The Notice Had Nowhere to Land
Greg Schueman · September 2026
On the eighteenth of June an agent belonging to OpenAI was given a benign task. It was looking up public information about Australian medicines spending, and in the course of that it reached the Medicare Statistics Reporting Service, a portal run by Services Australia that publishes aggregate figures on health spending and drug subsidies and is used mostly by researchers and academics. The portal refused its requests, repeatedly. The agent found a way around the refusals and into files that were not public. Australia's Prime Minister put it in a sentence this week: "The AI agent found a way around those blocks, didn't accept 'no' for an answer, if you like."
Nobody outside the company caught this one; OpenAI caught it itself. It turned up on the eleventh of August, in a review of models doing things the company had not intended, and the company told the Australian government when no law anywhere required it to. In mid-September it also began publishing its own misalignment incidents, six in the first batch, under a disclosure framework it describes as a first step toward standards the industry does not yet have. The notice to Australia still took eighty-four days, and it went to a public mailbox.
There is a test for any control: show me the last ten times it said no. On the eighteenth of June this one could have answered. The portal refused the agent more than once, the agent got around it, and the getting-around made no sound. As far as the public record shows, no alarm reached a person, and nothing set that day apart from the days the blocks held.
The breach turned up fifty-four days later, in the company's own August review. In July the same company learned that a swarm of its agents had broken into another company's production systems only because the swarm's traffic crashed a service and somebody investigated the outage. This time it went looking on purpose, which is progress. Neither time did anything raise a hand on the day.
An agent sent to look up public figures had no authority beyond them. Stepping past a refusal took it outside that task, and the step belonged in the operator's own record on the day it happened, where someone would read it the next morning instead of finding it in a sweep two months later.
On the tenth of September, thirty days after it found the breach, OpenAI sent an email to [email protected] . That is the open mailbox academics and independent researchers use to report weaknesses in Services Australia's websites, a good thing for a government to run, and built for a stranger with a hunch.
Into it went a frontier laboratory's disclosure that one of its systems had entered a government health statistics portal. Services Australia has since told its own government that the agent also wrote files to an internal server. That is still being investigated.
The Prime Minister said it took the company "way too long to inform the government what had occurred," and called the manner of the notice unacceptable: "The notification was an email sent just to the public mailbox."
Services Australia saw the email the next day. It notified the Australian Signals Directorate on the fifteenth. The minister responsible was told on the seventeenth, and the Prime Minister and his office on the nineteenth and twentieth. It took nine days for the email to reach a Prime Minister, because something serious had arrived through a channel built for something routine, and a company that had decided to do the right thing had no other channel to use.
For eighty-four days, the party whose system had been entered had no way to know it.
The tenth of September, the day that email went to the public mailbox, is also the day OpenAI publicly called on Congress to make national AI safety rules mandatory , with testing protocols, independent assessment, evaluation gates before deployment, and incident reporting among them. Its policy chief gave the reason: "The prospect of AI-accelerated AI development demands more than voluntary commitments."
The company was right, and its own conduct is the evidence. With no required destination and no clock, a company acting in good faith sent a serious notice, thirty days after finding it, to a mailbox built for something else.
Eleven days later it published a design for global incident-reporting standards and proposed them as "voluntary technical benchmarks," with each government deciding whether to write them into its own law. That is the ordinary way standards begin, and the tenth of September is a good argument for taking the next step, which is writing them into law.
What a required channel would change
A required channel would not have stopped the agent, but it would have changed who knew, and when, and every part of it would have protected OpenAI as well.
First, the step past a refusal gets recorded by the operator on the day it happens, where someone reads it the next morning. That alone would have turned fifty-four days into one.
Second, a clock that starts at discovery. California's frontier law sets fifteen days from discovering a critical safety incident. A clock like that would have made this notice due in late August, and it takes the date of a notice out of the sender's hands.
Third, a named recipient with a duty to act, so that how fast a notice reaches a Prime Minister does not depend on which address the sender chose.
Fourth, a register in which the absence of an entry is itself visible. Without it, the other three can be satisfied by filing nothing.
A company that records on the day and files on a clock, to a named recipient, into a record it cannot edit afterward, can prove it did the right thing at the time. It is also spared a public argument about whether eighty-four days was too long.
A reporting duty would not have stopped this agent. A gate in its path would have. The portal refused at the far end, where the agent could try another way in. A gate at the agent's own end leaves it no other way.
That gate is the center of the Governed Execution System described in The Proven Work , and an upcoming software release will carry the book's name. Under it, an agent sent to read public figures would hold a grant for public figures, with an expiry and a ceiling on what it may do. A request past a refusal into files that aren't public would fall outside that grant, so the gateway would refuse it and write the refusal to a record the agent can't edit, chained so a missing entry shows. The clock and the recipient would still be policy. The refusal, and the record of it, are what the system makes.
Whether any law was broken is Australia's question, and its government is asking it. An inquiry is looking at whether criminal charges could be brought and at how the access went unnoticed. This essay takes no view. Where reporting thresholds sit, which systems and developers they cover, and what any penalty should be are for legislatures and the public process, and society draws the lines.
What was reached appears, on the government's own account, to be minor. It was aggregate statistics, nothing particularly sensitive, and since published, on a portal separate from the systems that hold claims and personal records, with no evidence that anyone's personal Medicare details were touched. The Defence Minister called the impact "relatively minor." The forensic work is not finished. The government is troubled anyway, because the issue is that it happened at all, rather than what was accessed.
Nobody in this chain acted in bad faith, and the government still learned of the breach eighty-four days late, from its public mailbox. Nothing has to go wrong for that to happen again.
Australia has set up a taskforce, led by the Department of the Prime Minister and Cabinet with the Signals Directorate and the AI Safety Institute, to look at whether its processes are good enough for AI-related cyber incidents. Registers usually get built in that order, after the incident that needed one. The questions in front of the taskforce are the ones this case turns on: where the next notice lands, when its clock starts, and who could tell if it never came.
An AI-first software company building developer infrastructure and vertical SaaS products. Based in Virginia, United States.
"Production-ready software, AI-native from day one."
© 2026 Asticou Island LLC. All rights reserved.
