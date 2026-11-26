---
source: "https://thesis.truffle.net/introduction"
hn_url: "https://news.ycombinator.com/item?id=50660962"
title: "The Local AI Opportunity"
article_title: "The Local AI Opportunity"
image: "https://thesis.truffle.net/share/thesis-cover-v1.jpg"
author: "arrowsmith"
captured_at: "2026-11-26T23:30:36Z"
capture_tool: "hn-digest"
hn_id: 50660962
score: 1
comments: 0
posted_at: "2026-11-26T23:21:14Z"
tags:
  - hacker-news
---

# The Local AI Opportunity

- HN: [50660962](https://news.ycombinator.com/item?id=50660962)
- Source: [thesis.truffle.net](https://thesis.truffle.net/introduction)
- Score: 1
- Comments: 0
- Posted: 2026-11-26T23:21:14Z

## Translation

Title: The Local AI Opportunity
Description: Why AI will diffuse through society as a GPU per person.

Article text:
November 1, 2026
The Local AI Opportunity
Why AI will diffuse through society as a GPU per person
Thus far, every era of computing has given birth to a new unit to measure its marginal cost in.
During the 1960s, with the mainframe, we got the timeshare credit—a unit to measure how much time one was using on an expensive shared resource. Then, with the internet, we got the megabit—a measure of digital information that had to be sent from point A to point B.
In the intelligence era, we got the token: a tiny piece of language that, when stitched together with others, conveys some new information.
As tokens got cheaper, we applied AI to entirely new use cases. Chatbots became tool-calling agents, which then became strong agents, which have now turned into 24/7 agents.
Something similar happened in 1865 when Britain learned how to get more work out of a ton of coal. A more efficient steam engine not only saved fuel but also made steam power affordable to the next factory owner. William Stanley Jevons described this result in 1865:
“It is wholly a confusion of ideas to suppose that the economical use of fuel is equivalent to a diminished consumption. The very contrary is the truth.” W. Stanley Jevons, The Coal Question, 1865
2023 2024 2025 2026 2027 2028
Chatbots
Early agents
Strong agents
24/7 agents
Chatbots answered simple questions and consumed 10–100 tokens per session.
By contrast, early agents could do small jobs on their own, like answering email or running a piece of research. On average, sessions would cost upwards of 5,000 tokens.
Strong agents then came along and could carry a long task from start to finish, such as building software. Here, we first saw the term “tokenmaxxing”: spending millions of tokens at a time, without worrying about the token content and evaluating only the output work.
Now, 24/7 background agents like Muse, Grok Bot, Instinct, and Dots are emerging and will consume tokens nonstop, even when not being prompted.
The demand for intelligence might not really have an upper bound. After all, why would you put a cap on how much effective work you could get out of a capable digital knowledge worker?
The supply of intelligence, however, is an entirely different story. Tokens are produced in datacenters, and datacenters hit the real-world constraints of energy, land, and regulations. Exponential growth in the world of atoms isn’t really a thing.
Supply without physical limits
Unmet demand — the opportunity
Unmet demand
The opportunity
Three caps — supply held down by real-world limits
Counterfactual supply without energy, land, or regulation limits
Token demand — rising exponentially
Token demand
Cloud supply — flattening under real-world constraints
Cloud supply
2026
2027
2028
2029
2030
Supply, if energy, land, and regulation
were not limiting factors
Supply, if energy, land, and regulation
were not limiting factors
2026
2027
2028
2029
2030
This creates a large imbalance between the demand and supply curves, represented by the area between them marked “Unmet demand.” 1
And even if energy, land, and regulation weren’t obstacles, datacenter supply would, at best, grow linearly.
In May of this year, I described the situation to some of my investors as the greatest opportunity everyone is ignoring. I called it the agentic paradox.
You see, the more capable an agent is, the longer it is worth keeping it running and doing work. As you get better memory, tool use, and attention over long horizons, you’d opt to keep such an agent running all the time. And if it could run even faster, even better.
Therefore, even if one considers a fixed model size, token demand can quickly grow without bound.
We think that a lot of the unmet demand is going to start spilling over into a personal inference device. A computer with an operating system designed around intelligence can keep running 24/7, producing tokens with the only cost being input electricity.
It’s an incredibly unsexy idea at a time when sexy ideas command enormous premiums. Would you really want to build an AI computer that plugs into the wall when you can tweet about putting mass drivers on the Moon or covering the Sahara with intelligence factories?
In the 1970s, a very similar thing occurred. Computing revolved around large machines owned by institutions, shared by many people who had to ration out time between them.
Early personal computers looked unimpressive beside these powerful mainframes. They had much less memory and very limited software, and they were harder to use.
But as the machines improved, someone who wanted to write a document or keep the accounts for a small business could increasingly do it on a computer of their own. Each useful new application gave another person a reason to buy one. Personal computers spread through these individual purchases, finding their way into offices that already relied on mainframes and gradually taking over more of the work those machines had done.
So what might a conservative estimate of the world’s demand for 24/7 agents look like?
Well, to do the math, we’re going to have to make some assumptions about the variables.
Let’s say 1 billion people each have an agent that decodes at 50 tokens a second continuously, and the model is in the ballpark of 100 billion parameters. Let’s also assume that it all happens on H200s. Here’s what that serving workload would look like:
Not many understand how outrageous a number 100 gigawatts is. That’s approximately the entire capacity of the American nuclear fleet . What we’re saying is that 24/7 agents in the hands of 1 billion people would keep the equivalent of all the nuclear reactors in America occupied around the clock, without even accounting for other use cases of AI, like enterprise or training.
The natural response from AI CEOs to this reality is to build more. I think we’ll build a lot of datacenters, but doing this in anything under a decade, at a time when America’s grid is buckling under the strain of electric vehicles and other new demand, is extremely unreasonable.
That’s not even accounting for things like community backlash or politicians who are riding the anti-AI sentiment of the general population. In many ways, narratives that are clearly false are still being parroted to avoid larger market drawdowns, at a time when almost all growth in the US economy is attributed to AI.
Not all things are built at the speed of bits
✕
The Grid
Years of waiting for grid connections and power
The Neighbors
Neighbors, noise, water, and what the polls really report
The State
State freezes, grid audits, and the price of saying yes
The Grid
Before anything else, a datacenter is a connection to the electric grid. Google calls transmission the biggest challenge in getting power to its datacenters. And in some parts of the country, the wait has stretched beyond a decade. 2
In Northern Virginia, JLL puts the average wait for a new 50-megawatt grid connection at about 7 years. 3 This is the reality of high-voltage construction. Things like routing, permitting, transformer backlogs, and securing land rights can slow the buildout to a crawl.
By comparison, the IEA puts datacenter construction at 1–3 years. New grid infrastructure can take 5–15 years to plan, permit, and build. 4
In October, Business Insider reported that Amazon was storing Nvidia GB300 racks in warehouses because some of its datacenters had room for the machines but not enough power to run them.
An internal email projected $5B of GB300 deliveries in August, followed by another $3B a month. At least one AWS team had no clear schedule for getting the incoming racks into datacenters. 5
Let’s imagine the grid weren’t an issue. Datacenters are still noisy eyesores, the size of several football fields, that feed on local power and water. Americans don’t want to live next to a datacenter, and they’re making their point clear during elections.
Data Center Watch initially counted $64B in American projects blocked or delayed by local opposition. 6 By Q1 2026, it recorded roughly $130B in a single quarter, with 833 opposition groups organized across 49 states. In Q2, another 45 projects worth $68B were blocked or delayed. 7
This opposition isn’t fringe. Anti-datacenter sentiment seems to be the only thing Democrats, independents, and Republicans agree on. In the poll above, 69% of respondents opposed an AI datacenter in their own community. 8
In a separate NBC poll, AI had a slightly lower net favorability rating than ICE. 9
While it’s easy to dismiss this sentiment as irrelevant, people are becoming more vocal and politically charged in their hatred of AI. This has the potential to become a headwind that only compounds in severity over time.
Where the neighbors organize, the state eventually follows. In August 2026 Governor Greg Abbott halted datacenter grid approvals across Texas pending an audit, 10 an order that froze 474 gigawatts of pending connection requests, more than 5 times the record peak demand of the entire ERCOT system. 11 In September he extended the halt to every environmental permit a datacenter needs, and the audit is due in December. New York went further. Governor Hochul froze new hyperscale datacenters by executive order in July, the legislature passed a 1-year ban behind it, and nearly 2/3 of New Yorkers polled want the ban kept when it expires. 12
Nor is this a passing mood in 2 statehouses. Wood Mackenzie counts more than 400 datacenter moratoriums enacted since January across 42 states, touching nearly 20% of the sites it tracks and more than 100 gigawatts of planned capacity, and the Texas pause alone puts 60 gigawatts of projects under audit. 13
Then came the midterms. Going in, NBC found 64% of voters less likely to support a candidate who backed a local datacenter, and 11% more likely. On November 3 datacenters were on the ballot 41 times in 10 states, and the industry won once, in Harwood, North Dakota. Wells, Maine, passed a retroactive 6-month moratorium with more than 80% of the vote. Of the 74 candidates for governor that Data Center Watch tracked, 43% had backed a moratorium or a ban, and that side won the statehouses that matter most. Texas, which holds about 25% of the datacenter capacity in the works, went to Gina Hinojosa on a promise to freeze approvals until the legislature writes rules; Ohio went to Amy Acton and Georgia to Keisha Lance Bottoms, who ran on the same issue. In the NBC exit poll, 56% of voters called AI and datacenters a major factor in their vote for Congress. 13 The political price of saying yes to a datacenter is no longer rising. It has been paid, in public, by candidates who said yes, and nothing about the demand curve in the introduction will bring it back down.
Against the requirements we just estimated for 24/7 agents, what’s being built is extremely small. Goldman Sachs’ May outlook projected American datacenter additions of 13.6 GW in 2026 and 36.3 GW in 2027. 14
Jefferies put the likely pace at 15–20 GW a year. By June, construction had begun on only 12 of the 24 GW planned for 2026, and up to 80% of the 2027–28 pipeline had yet to begin substantial construction. 15
1 billion always-on agents
100B model, 50 tok/s, 24/7
Planned additions, 2027
Planned additions, 2026
Realistic pace, per year
Jefferies
100 GW, continuous
36.3 GW
13.6 GW
15–20 GW
Fig. 4. What 1 billion always-on agents require, against what is planned and what is plausible. Sources: Goldman Sachs Research; Jefferies.
Financing the buildout
Chairman Kevin Warsh at the FOMC press conference, June 17, 2026. Federal Reserve photo.
The buildout is financed by equity and debt. The debt makes every delay more expensive as interest keeps accumulating. With floating-rate debt, the interest bill can even rise during the delay.
In September, the Fed under Warsh raised rates into this buildout, bringing an uncomfortable historical parallel to mind. 16
The internet boom was also a construction boom, fueled by telecom companies borrowing heavily to lay fi

[truncated]

## Original Extract

Why AI will diffuse through society as a GPU per person.

November 1, 2026
The Local AI Opportunity
Why AI will diffuse through society as a GPU per person
Thus far, every era of computing has given birth to a new unit to measure its marginal cost in.
During the 1960s, with the mainframe, we got the timeshare credit—a unit to measure how much time one was using on an expensive shared resource. Then, with the internet, we got the megabit—a measure of digital information that had to be sent from point A to point B.
In the intelligence era, we got the token: a tiny piece of language that, when stitched together with others, conveys some new information.
As tokens got cheaper, we applied AI to entirely new use cases. Chatbots became tool-calling agents, which then became strong agents, which have now turned into 24/7 agents.
Something similar happened in 1865 when Britain learned how to get more work out of a ton of coal. A more efficient steam engine not only saved fuel but also made steam power affordable to the next factory owner. William Stanley Jevons described this result in 1865:
“It is wholly a confusion of ideas to suppose that the economical use of fuel is equivalent to a diminished consumption. The very contrary is the truth.” W. Stanley Jevons, The Coal Question, 1865
2023 2024 2025 2026 2027 2028
Chatbots
Early agents
Strong agents
24/7 agents
Chatbots answered simple questions and consumed 10–100 tokens per session.
By contrast, early agents could do small jobs on their own, like answering email or running a piece of research. On average, sessions would cost upwards of 5,000 tokens.
Strong agents then came along and could carry a long task from start to finish, such as building software. Here, we first saw the term “tokenmaxxing”: spending millions of tokens at a time, without worrying about the token content and evaluating only the output work.
Now, 24/7 background agents like Muse, Grok Bot, Instinct, and Dots are emerging and will consume tokens nonstop, even when not being prompted.
The demand for intelligence might not really have an upper bound. After all, why would you put a cap on how much effective work you could get out of a capable digital knowledge worker?
The supply of intelligence, however, is an entirely different story. Tokens are produced in datacenters, and datacenters hit the real-world constraints of energy, land, and regulations. Exponential growth in the world of atoms isn’t really a thing.
Supply without physical limits
Unmet demand — the opportunity
Unmet demand
The opportunity
Three caps — supply held down by real-world limits
Counterfactual supply without energy, land, or regulation limits
Token demand — rising exponentially
Token demand
Cloud supply — flattening under real-world constraints
Cloud supply
2026
2027
2028
2029
2030
Supply, if energy, land, and regulation
were not limiting factors
Supply, if energy, land, and regulation
were not limiting factors
2026
2027
2028
2029
2030
This creates a large imbalance between the demand and supply curves, represented by the area between them marked “Unmet demand.” 1
And even if energy, land, and regulation weren’t obstacles, datacenter supply would, at best, grow linearly.
In May of this year, I described the situation to some of my investors as the greatest opportunity everyone is ignoring. I called it the agentic paradox.
You see, the more capable an agent is, the longer it is worth keeping it running and doing work. As you get better memory, tool use, and attention over long horizons, you’d opt to keep such an agent running all the time. And if it could run even faster, even better.
Therefore, even if one considers a fixed model size, token demand can quickly grow without bound.
We think that a lot of the unmet demand is going to start spilling over into a personal inference device. A computer with an operating system designed around intelligence can keep running 24/7, producing tokens with the only cost being input electricity.
It’s an incredibly unsexy idea at a time when sexy ideas command enormous premiums. Would you really want to build an AI computer that plugs into the wall when you can tweet about putting mass drivers on the Moon or covering the Sahara with intelligence factories?
In the 1970s, a very similar thing occurred. Computing revolved around large machines owned by institutions, shared by many people who had to ration out time between them.
Early personal computers looked unimpressive beside these powerful mainframes. They had much less memory and very limited software, and they were harder to use.
But as the machines improved, someone who wanted to write a document or keep the accounts for a small business could increasingly do it on a computer of their own. Each useful new application gave another person a reason to buy one. Personal computers spread through these individual purchases, finding their way into offices that already relied on mainframes and gradually taking over more of the work those machines had done.
So what might a conservative estimate of the world’s demand for 24/7 agents look like?
Well, to do the math, we’re going to have to make some assumptions about the variables.
Let’s say 1 billion people each have an agent that decodes at 50 tokens a second continuously, and the model is in the ballpark of 100 billion parameters. Let’s also assume that it all happens on H200s. Here’s what that serving workload would look like:
Not many understand how outrageous a number 100 gigawatts is. That’s approximately the entire capacity of the American nuclear fleet . What we’re saying is that 24/7 agents in the hands of 1 billion people would keep the equivalent of all the nuclear reactors in America occupied around the clock, without even accounting for other use cases of AI, like enterprise or training.
The natural response from AI CEOs to this reality is to build more. I think we’ll build a lot of datacenters, but doing this in anything under a decade, at a time when America’s grid is buckling under the strain of electric vehicles and other new demand, is extremely unreasonable.
That’s not even accounting for things like community backlash or politicians who are riding the anti-AI sentiment of the general population. In many ways, narratives that are clearly false are still being parroted to avoid larger market drawdowns, at a time when almost all growth in the US economy is attributed to AI.
Not all things are built at the speed of bits
✕
The Grid
Years of waiting for grid connections and power
The Neighbors
Neighbors, noise, water, and what the polls really report
The State
State freezes, grid audits, and the price of saying yes
The Grid
Before anything else, a datacenter is a connection to the electric grid. Google calls transmission the biggest challenge in getting power to its datacenters. And in some parts of the country, the wait has stretched beyond a decade. 2
In Northern Virginia, JLL puts the average wait for a new 50-megawatt grid connection at about 7 years. 3 This is the reality of high-voltage construction. Things like routing, permitting, transformer backlogs, and securing land rights can slow the buildout to a crawl.
By comparison, the IEA puts datacenter construction at 1–3 years. New grid infrastructure can take 5–15 years to plan, permit, and build. 4
In October, Business Insider reported that Amazon was storing Nvidia GB300 racks in warehouses because some of its datacenters had room for the machines but not enough power to run them.
An internal email projected $5B of GB300 deliveries in August, followed by another $3B a month. At least one AWS team had no clear schedule for getting the incoming racks into datacenters. 5
Let’s imagine the grid weren’t an issue. Datacenters are still noisy eyesores, the size of several football fields, that feed on local power and water. Americans don’t want to live next to a datacenter, and they’re making their point clear during elections.
Data Center Watch initially counted $64B in American projects blocked or delayed by local opposition. 6 By Q1 2026, it recorded roughly $130B in a single quarter, with 833 opposition groups organized across 49 states. In Q2, another 45 projects worth $68B were blocked or delayed. 7
This opposition isn’t fringe. Anti-datacenter sentiment seems to be the only thing Democrats, independents, and Republicans agree on. In the poll above, 69% of respondents opposed an AI datacenter in their own community. 8
In a separate NBC poll, AI had a slightly lower net favorability rating than ICE. 9
While it’s easy to dismiss this sentiment as irrelevant, people are becoming more vocal and politically charged in their hatred of AI. This has the potential to become a headwind that only compounds in severity over time.
Where the neighbors organize, the state eventually follows. In August 2026 Governor Greg Abbott halted datacenter grid approvals across Texas pending an audit, 10 an order that froze 474 gigawatts of pending connection requests, more than 5 times the record peak demand of the entire ERCOT system. 11 In September he extended the halt to every environmental permit a datacenter needs, and the audit is due in December. New York went further. Governor Hochul froze new hyperscale datacenters by executive order in July, the legislature passed a 1-year ban behind it, and nearly 2/3 of New Yorkers polled want the ban kept when it expires. 12
Nor is this a passing mood in 2 statehouses. Wood Mackenzie counts more than 400 datacenter moratoriums enacted since January across 42 states, touching nearly 20% of the sites it tracks and more than 100 gigawatts of planned capacity, and the Texas pause alone puts 60 gigawatts of projects under audit. 13
Then came the midterms. Going in, NBC found 64% of voters less likely to support a candidate who backed a local datacenter, and 11% more likely. On November 3 datacenters were on the ballot 41 times in 10 states, and the industry won once, in Harwood, North Dakota. Wells, Maine, passed a retroactive 6-month moratorium with more than 80% of the vote. Of the 74 candidates for governor that Data Center Watch tracked, 43% had backed a moratorium or a ban, and that side won the statehouses that matter most. Texas, which holds about 25% of the datacenter capacity in the works, went to Gina Hinojosa on a promise to freeze approvals until the legislature writes rules; Ohio went to Amy Acton and Georgia to Keisha Lance Bottoms, who ran on the same issue. In the NBC exit poll, 56% of voters called AI and datacenters a major factor in their vote for Congress. 13 The political price of saying yes to a datacenter is no longer rising. It has been paid, in public, by candidates who said yes, and nothing about the demand curve in the introduction will bring it back down.
Against the requirements we just estimated for 24/7 agents, what’s being built is extremely small. Goldman Sachs’ May outlook projected American datacenter additions of 13.6 GW in 2026 and 36.3 GW in 2027. 14
Jefferies put the likely pace at 15–20 GW a year. By June, construction had begun on only 12 of the 24 GW planned for 2026, and up to 80% of the 2027–28 pipeline had yet to begin substantial construction. 15
1 billion always-on agents
100B model, 50 tok/s, 24/7
Planned additions, 2027
Planned additions, 2026
Realistic pace, per year
Jefferies
100 GW, continuous
36.3 GW
13.6 GW
15–20 GW
Fig. 4. What 1 billion always-on agents require, against what is planned and what is plausible. Sources: Goldman Sachs Research; Jefferies.
Financing the buildout
Chairman Kevin Warsh at the FOMC press conference, June 17, 2026. Federal Reserve photo.
The buildout is financed by equity and debt. The debt makes every delay more expensive as interest keeps accumulating. With floating-rate debt, the interest bill can even rise during the delay.
In September, the Fed under Warsh raised rates into this buildout, bringing an uncomfortable historical parallel to mind. 16
The internet boom was also a construction boom, fueled by telecom companies borrowing heavily to lay fi

[truncated]
