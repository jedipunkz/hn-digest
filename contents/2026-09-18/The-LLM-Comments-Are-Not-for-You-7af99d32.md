---
source: "https://danilafe.com/blog/comments_not_for_you/"
hn_url: "https://news.ycombinator.com/item?id=49751659"
title: "The LLM Comments Are Not for You"
article_title: "The LLM Comments Are Not For You"
image: "https://danilafe.com/banner.png"
author: "kqr"
captured_at: "2026-09-18T08:49:16Z"
capture_tool: "hn-digest"
hn_id: 49751659
score: 1
comments: 0
posted_at: "2026-09-18T08:38:27Z"
tags:
  - hacker-news
---

# The LLM Comments Are Not for You

- HN: [49751659](https://news.ycombinator.com/item?id=49751659)
- Source: [danilafe.com](https://danilafe.com/blog/comments_not_for_you/)
- Score: 1
- Comments: 0
- Posted: 2026-09-18T08:38:27Z

## Translation

Title: The LLM Comments Are Not for You
Article title: The LLM Comments Are Not For You
Description: I’ve heard a lot of talk about LLMs recently, and among the most common topics of discussion have been the comments. LLM comments generally seem to suck: they are verbose, regardless of surrounding context, they seem to encode conversation decisions (like “do it this way, not that way”[note: In my h
[truncated]

Article text:
The LLM Comments Are Not For You
Daniel's Blog
The LLM Comments Are Not For You
I’ve heard a lot of talk about LLMs recently, and among the most common topics
of discussion have been the comments. LLM comments generally seem to suck: they are verbose,
regardless of surrounding context, they seem to encode conversation decisions
(like
“do it this way, not that way” [note:
In my head, I've been calling this "comment hysteresis", because the comments
are not a function of the final design, but of the path taken to get there.
]
), and they tend to use made-up vocabulary or terms. Engineers
have been working on ways to reduce the pain, by carefully crafting their
prompts or even separately running “desloppifier” agents to clean up PRs.
In my experience, these techniques are only mildly successful.
So why do LLMs keep writing comments like these, even as their software benchmark
scores climb ever higher with new model releases? I’d like to argue that they
remain bad precisely because the scores in benchmarks are getting
better. And very likely they will degrade, in the same way that
model tool calling has gotten worse over time
.
These comments are not for you, my (hopefully human) reader.
Fundamentally, as many in the coding agent space have pointed out (like,
say, Mario Zechner in his talk about pi
),
most modern models have been RL
’ed
to fit into agentic harnesses. These workflows are part of their training.
And the way that RL works is that it is outcome-based. A model takes steps,
edits files, does whatever it does in its agentic framework, and either
arrives at a solution or doesn’t. Behaviors that contributed to successful
outcomes are encouraged, and become more common.
The difficulty is that as benchmarks get harder and as models are asked
to take on larger and larger chunks of the software development workflow,
their ability to keep information “just” within their context is
pushed closer to its limit. Session compaction can accidentally destroy
design decisions or rationale, forcing the model to eventually re-discover
previous decisions or even switch directions. I suspect that in addition
to this — given the techniques used by OpenAI in
its formalization [note:
Whether this was really OpenAI's formalization or plagiarism remains open
to debate, and I do not know enough to claim one way or the other.
The possessive form here is just for convenience.
]
of the Navier-Stokes singularity
— models are also trained to operate in
swarms, which don’t share context but must find ways to coordinate with
each other. [note:
I don't know how agents talk to each other, but we've already seen that the
way they talk to themselves is
very different
from humans.
]
Comments are a hugely useful persistent store of contextual information.
If a model edits a file, chances are it will read it as well, discover
the comments, and re-load the given information into context. If
one agent makes a change to a file with some design rationale — “array,
not a linked list” — another agent that might want to change it back will
spot that and
tread carefully. [note:
I suspect, though with less certainty, that agents are very deferential
to pre-existing comments for this exact reason. Claude Code, for instance,
will be very insistent that when code says to do X, the new code written
should fit the "X model".
]
It should come as no surprise that agents that make use of verbose,
decision-making-included [note:
Interestingly, this turn of phrase is uncommon in pre-existing human comments,
which would make up the majority of the model's training data set. In
my opinion, this points towards this being something more than a reflection
of the "human style".
]
comments succeed more frequently, and get
rewarded. The result: +50/-1 comment blocks in your diff. It is irrelevant
whether these comments concisely describe the codebase; their intended
audience can read and parse them instantly. The comments are not for you.
It’s very hard to prompt this behavior out, and there’s a pretty good chance
that you don’t want to, by the simple evolutionary argument: these things
have helped the model do well in evaluations. By removing them, you are likely
undermining part of whatever mechanism makes it tick. On top of that, you
are fighting the wiring it has developed to do exactly this. It’s like trying
to get humans to stop liking hyperpalatable foods
.
So then, the natural conclusion is that we should be leaving these comments as
they are, right? As long as they’ve been known to improve agents’ performance,
the more the better? Some, who have
leaned heavily on models for self-regulating via persistent state
,
have found fascinating emergent behaviors, including whole organizational
structures with
agent-invented-names [note:
Did I mention that LLM comments tend to invent novel vocabulary?
]
. They believe that to be the future.
However, in my opinion, it may not be that simple. As we’ve seen with human
evolutionary adaptations, they don’t always do well outside of the environment
in which they arose. For instance — to hammer the point — humans love
hyperpalatable foods. They will eat them in excess, which can lead to
obesity and a variety of other conditions. These too are emergent behaviors,
and likely quite interesting from a medical perspective. That does not
make them desirable.
In the same way that scarcity of sugars and fats in nature balanced (and
motivated) human enjoyment of them, it’s possible that the ephemeral nature
of software engineering “tasks” is counteracting the accumulation of LLM-generated
commentary. In real-world contexts, human attention and editing may be doing
the same thing. We are yet to see what codebases maintained entirely
with RL’ed agentic behaviors over years look like, and whether there are
limitations.
Regardless of whether this new style holds up in the extreme,
it’s that way for a reason — and you are no longer its sole intended audience.
Liked this article? Have any questions or comments? Please don't hesitate to reach out to me at danila.fedorin@gmail.com ! I love receiving emails from readers, and I'm always happy to provide any additional clarification or assistance.

## Original Extract

I’ve heard a lot of talk about LLMs recently, and among the most common topics of discussion have been the comments. LLM comments generally seem to suck: they are verbose, regardless of surrounding context, they seem to encode conversation decisions (like “do it this way, not that way”[note: In my h
[truncated]

The LLM Comments Are Not For You
Daniel's Blog
The LLM Comments Are Not For You
I’ve heard a lot of talk about LLMs recently, and among the most common topics
of discussion have been the comments. LLM comments generally seem to suck: they are verbose,
regardless of surrounding context, they seem to encode conversation decisions
(like
“do it this way, not that way” [note:
In my head, I've been calling this "comment hysteresis", because the comments
are not a function of the final design, but of the path taken to get there.
]
), and they tend to use made-up vocabulary or terms. Engineers
have been working on ways to reduce the pain, by carefully crafting their
prompts or even separately running “desloppifier” agents to clean up PRs.
In my experience, these techniques are only mildly successful.
So why do LLMs keep writing comments like these, even as their software benchmark
scores climb ever higher with new model releases? I’d like to argue that they
remain bad precisely because the scores in benchmarks are getting
better. And very likely they will degrade, in the same way that
model tool calling has gotten worse over time
.
These comments are not for you, my (hopefully human) reader.
Fundamentally, as many in the coding agent space have pointed out (like,
say, Mario Zechner in his talk about pi
),
most modern models have been RL
’ed
to fit into agentic harnesses. These workflows are part of their training.
And the way that RL works is that it is outcome-based. A model takes steps,
edits files, does whatever it does in its agentic framework, and either
arrives at a solution or doesn’t. Behaviors that contributed to successful
outcomes are encouraged, and become more common.
The difficulty is that as benchmarks get harder and as models are asked
to take on larger and larger chunks of the software development workflow,
their ability to keep information “just” within their context is
pushed closer to its limit. Session compaction can accidentally destroy
design decisions or rationale, forcing the model to eventually re-discover
previous decisions or even switch directions. I suspect that in addition
to this — given the techniques used by OpenAI in
its formalization [note:
Whether this was really OpenAI's formalization or plagiarism remains open
to debate, and I do not know enough to claim one way or the other.
The possessive form here is just for convenience.
]
of the Navier-Stokes singularity
— models are also trained to operate in
swarms, which don’t share context but must find ways to coordinate with
each other. [note:
I don't know how agents talk to each other, but we've already seen that the
way they talk to themselves is
very different
from humans.
]
Comments are a hugely useful persistent store of contextual information.
If a model edits a file, chances are it will read it as well, discover
the comments, and re-load the given information into context. If
one agent makes a change to a file with some design rationale — “array,
not a linked list” — another agent that might want to change it back will
spot that and
tread carefully. [note:
I suspect, though with less certainty, that agents are very deferential
to pre-existing comments for this exact reason. Claude Code, for instance,
will be very insistent that when code says to do X, the new code written
should fit the "X model".
]
It should come as no surprise that agents that make use of verbose,
decision-making-included [note:
Interestingly, this turn of phrase is uncommon in pre-existing human comments,
which would make up the majority of the model's training data set. In
my opinion, this points towards this being something more than a reflection
of the "human style".
]
comments succeed more frequently, and get
rewarded. The result: +50/-1 comment blocks in your diff. It is irrelevant
whether these comments concisely describe the codebase; their intended
audience can read and parse them instantly. The comments are not for you.
It’s very hard to prompt this behavior out, and there’s a pretty good chance
that you don’t want to, by the simple evolutionary argument: these things
have helped the model do well in evaluations. By removing them, you are likely
undermining part of whatever mechanism makes it tick. On top of that, you
are fighting the wiring it has developed to do exactly this. It’s like trying
to get humans to stop liking hyperpalatable foods
.
So then, the natural conclusion is that we should be leaving these comments as
they are, right? As long as they’ve been known to improve agents’ performance,
the more the better? Some, who have
leaned heavily on models for self-regulating via persistent state
,
have found fascinating emergent behaviors, including whole organizational
structures with
agent-invented-names [note:
Did I mention that LLM comments tend to invent novel vocabulary?
]
. They believe that to be the future.
However, in my opinion, it may not be that simple. As we’ve seen with human
evolutionary adaptations, they don’t always do well outside of the environment
in which they arose. For instance — to hammer the point — humans love
hyperpalatable foods. They will eat them in excess, which can lead to
obesity and a variety of other conditions. These too are emergent behaviors,
and likely quite interesting from a medical perspective. That does not
make them desirable.
In the same way that scarcity of sugars and fats in nature balanced (and
motivated) human enjoyment of them, it’s possible that the ephemeral nature
of software engineering “tasks” is counteracting the accumulation of LLM-generated
commentary. In real-world contexts, human attention and editing may be doing
the same thing. We are yet to see what codebases maintained entirely
with RL’ed agentic behaviors over years look like, and whether there are
limitations.
Regardless of whether this new style holds up in the extreme,
it’s that way for a reason — and you are no longer its sole intended audience.
Liked this article? Have any questions or comments? Please don't hesitate to reach out to me at danila.fedorin@gmail.com ! I love receiving emails from readers, and I'm always happy to provide any additional clarification or assistance.
