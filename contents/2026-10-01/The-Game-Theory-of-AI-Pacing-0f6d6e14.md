---
source: "https://www.paradigm.xyz/writing/the-game-theory-of-ai-pacing"
hn_url: "https://news.ycombinator.com/item?id=49926409"
title: "The Game Theory of AI Pacing"
article_title: "Pace, a multiplayer web game about AI race dynamics, by Paradigm."
image: "https://images.ctfassets.net/vb2n37v5ldjn/7ipgwUoO8wdP7gPqyS1CpM/ccb390518aa7b1e3ebbb307837632c45/opengraph-pace.jpg?w=1200&h=630&fit=fill&f=center&fm=jpg&q=85"
author: "ronfriedhaber"
captured_at: "2026-10-01T20:10:00Z"
capture_tool: "hn-digest"
hn_id: 49926409
score: 1
comments: 0
posted_at: "2026-10-01T20:02:44Z"
tags:
  - hacker-news
---

# The Game Theory of AI Pacing

- HN: [49926409](https://news.ycombinator.com/item?id=49926409)
- Source: [www.paradigm.xyz](https://www.paradigm.xyz/writing/the-game-theory-of-ai-pacing)
- Score: 1
- Comments: 0
- Posted: 2026-10-01T20:02:44Z

## Translation

Title: The Game Theory of AI Pacing
Article title: Pace, a multiplayer web game about AI race dynamics, by Paradigm.
Description: Introducing Pace, a multiplayer web game about AI race dynamics, by Paradigm.

Article text:
Pace, a multiplayer web game about AI race dynamics, by Paradigm.
/ Writing [ M ] [M] [X] Menu About Team Investments Research Research Index Paradigm Puzzles Build Incubations Open Source Writing Terms , Disclosures , Privacy , CA Privacy
By Justin Wang , Dan Robinson , Andrew Koh
Some major AI labs, including Anthropic and OpenAI , have been calling for a coordinated slowdown in frontier AI research. In response, some have been reasonably asking why, if a given company is so concerned about the risks of what it is building, it can’t just voluntarily slow down.
One answer comes from game theory. And it also explains why the U.S. has been hesitant to slow down AI progress without international coordination.
We created a multiplayer web game inspired by a recent paper from Drew Fudenberg and Andrew Koh, which explains the game theory of competition in AI development. The paper models AI labs as players in a two-player game, who can either accelerate or decelerate. Companies make money when they are close to or ahead of their competitor, and lose money if they fall too far behind.
There is a shared safety frontier, which advances over time. The further the leading model is past that safety frontier, the higher the risk of catastrophe.
Zoom The paper shows that whether firms race (develop at max speed) or pace (develop along the safety frontier) depends on parameters that include the probability of catastrophic risk, discount rate, and transparency around each other’s progress. The strategic situation at hand exhibits elements of both competition : labs want to be ahead and make money, and cooperation : nobody wants to be too reckless and cause unintended harm.
Our game demonstrates this problem, and includes some in-game mechanics that contribute to the risk:
Labs do not have immediate transparency into the other’s research progress, and can only react to decisions their competitor made weeks ago. In the game, acceleration is private for a little while, and so can only be reciprocated with a lag.
Acceleration and deceleration are gradual, and development is irreversible.
The path of the safety frontier is unpredictable: sometimes it flattens out, and sometimes it rises steeply.
As the game progresses, the acceleration and maximum speed become faster, reflecting the possibility of recursive self-improvement .
In practice, we think it is even more difficult to control the pace of progress in the real world, due to factors that are not modeled by the game.
While the game only exposes one control for acceleration and deceleration, the AI development loop takes many factors as input — such as compute, data, and labor — which often require long-term investments and commitments.
In the real world, labs can keep internal progress secret for longer.
The real safety frontier is unknown and arguably unknowable, and labs may disagree on the degree of risk from being at a particular capability level. This disagreement can make the effective safety frontier driven by the player who is least concerned or most skeptical of AI risks.
There are more than two competitors in the real world.
The model may have implications beyond competition between labs. Similar dynamics are at play in competition between countries like the U.S. and China. An international slowdown is difficult because of the race between multiple nations. In the absence of verification and transparency measures, or even simple agreements around shared safety standards, the natural emergent behavior is likely to resemble a race as well.
We hope this game provides helpful intuition and helps show how race dynamics can contribute to risk, especially in the absence of cooperation.
Thank you to Will Robinson and Kevin Liu for helpful feedback.
Disclaimer: This post is for general information purposes only. It does not constitute
investment advice or a recommendation or solicitation to buy or sell any investment and
should not be used in the evaluation of the merits of making any investment decision. It
should not be relied upon for accounting, legal or tax advice or investment recommendations.
This post reflects the current opinions of the authors and is not made on behalf of Paradigm
or its affiliates and does not necessarily reflect the opinions of Paradigm, its affiliates
or individuals associated with Paradigm. The opinions reflected herein are subject to change
without being updated.
Terms , Disclosures , Privacy , CA Privacy

## Original Extract

Introducing Pace, a multiplayer web game about AI race dynamics, by Paradigm.

Pace, a multiplayer web game about AI race dynamics, by Paradigm.
/ Writing [ M ] [M] [X] Menu About Team Investments Research Research Index Paradigm Puzzles Build Incubations Open Source Writing Terms , Disclosures , Privacy , CA Privacy
By Justin Wang , Dan Robinson , Andrew Koh
Some major AI labs, including Anthropic and OpenAI , have been calling for a coordinated slowdown in frontier AI research. In response, some have been reasonably asking why, if a given company is so concerned about the risks of what it is building, it can’t just voluntarily slow down.
One answer comes from game theory. And it also explains why the U.S. has been hesitant to slow down AI progress without international coordination.
We created a multiplayer web game inspired by a recent paper from Drew Fudenberg and Andrew Koh, which explains the game theory of competition in AI development. The paper models AI labs as players in a two-player game, who can either accelerate or decelerate. Companies make money when they are close to or ahead of their competitor, and lose money if they fall too far behind.
There is a shared safety frontier, which advances over time. The further the leading model is past that safety frontier, the higher the risk of catastrophe.
Zoom The paper shows that whether firms race (develop at max speed) or pace (develop along the safety frontier) depends on parameters that include the probability of catastrophic risk, discount rate, and transparency around each other’s progress. The strategic situation at hand exhibits elements of both competition : labs want to be ahead and make money, and cooperation : nobody wants to be too reckless and cause unintended harm.
Our game demonstrates this problem, and includes some in-game mechanics that contribute to the risk:
Labs do not have immediate transparency into the other’s research progress, and can only react to decisions their competitor made weeks ago. In the game, acceleration is private for a little while, and so can only be reciprocated with a lag.
Acceleration and deceleration are gradual, and development is irreversible.
The path of the safety frontier is unpredictable: sometimes it flattens out, and sometimes it rises steeply.
As the game progresses, the acceleration and maximum speed become faster, reflecting the possibility of recursive self-improvement .
In practice, we think it is even more difficult to control the pace of progress in the real world, due to factors that are not modeled by the game.
While the game only exposes one control for acceleration and deceleration, the AI development loop takes many factors as input — such as compute, data, and labor — which often require long-term investments and commitments.
In the real world, labs can keep internal progress secret for longer.
The real safety frontier is unknown and arguably unknowable, and labs may disagree on the degree of risk from being at a particular capability level. This disagreement can make the effective safety frontier driven by the player who is least concerned or most skeptical of AI risks.
There are more than two competitors in the real world.
The model may have implications beyond competition between labs. Similar dynamics are at play in competition between countries like the U.S. and China. An international slowdown is difficult because of the race between multiple nations. In the absence of verification and transparency measures, or even simple agreements around shared safety standards, the natural emergent behavior is likely to resemble a race as well.
We hope this game provides helpful intuition and helps show how race dynamics can contribute to risk, especially in the absence of cooperation.
Thank you to Will Robinson and Kevin Liu for helpful feedback.
Disclaimer: This post is for general information purposes only. It does not constitute
investment advice or a recommendation or solicitation to buy or sell any investment and
should not be used in the evaluation of the merits of making any investment decision. It
should not be relied upon for accounting, legal or tax advice or investment recommendations.
This post reflects the current opinions of the authors and is not made on behalf of Paradigm
or its affiliates and does not necessarily reflect the opinions of Paradigm, its affiliates
or individuals associated with Paradigm. The opinions reflected herein are subject to change
without being updated.
Terms , Disclosures , Privacy , CA Privacy
