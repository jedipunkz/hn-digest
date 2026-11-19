---
source: "https://www.edgee.ai/blog/posts/switching-llms-mid-conversation-is-a-ski-rental-problem"
hn_url: "https://news.ycombinator.com/item?id=50561397"
title: "Switching LLMs mid-conversation is a ski-rental problem"
article_title: "Switching LLMs mid-conversation is a ski-rental problem - Edgee Blog"
image: "https://cdn.edgee.app/blog/inline/6a68b5d5-df68-46e9-a8c7-b560b7b591bb.png"
author: "gmays"
captured_at: "2026-11-19T16:21:42Z"
capture_tool: "hn-digest"
hn_id: 50561397
score: 1
comments: 0
posted_at: "2026-11-19T15:30:02Z"
tags:
  - hacker-news
---

# Switching LLMs mid-conversation is a ski-rental problem

- HN: [50561397](https://news.ycombinator.com/item?id=50561397)
- Source: [www.edgee.ai](https://www.edgee.ai/blog/posts/switching-llms-mid-conversation-is-a-ski-rental-problem)
- Score: 1
- Comments: 0
- Posted: 2026-11-19T15:30:02Z

## Translation

Title: Switching LLMs mid-conversation is a ski-rental problem
Article title: Switching LLMs mid-conversation is a ski-rental problem - Edgee Blog
Description: Most LLM routers answer one question: which model should handle this request? For a coding agent that is only half the story. Every switch throws away the prompt cache, so deciding when a session should change model is a ski-rental problem. Self-Learning Routing learns both halves from an organizati
[truncated]

Article text:
Switching LLMs mid-conversation is a ski-rental problem - Edgee Blog Features Smart Routing & Budgets Token Compression Team Observability Integrations Plugins Claude Mod On-Premise Edgee vs OpenRouter Edgee vs LiteLLM Models Blog Pricing Documentation Login Open main menu Switching LLMs mid-conversation is a ski-rental problem
Most work on routing between language models answers one question: which model should handle this
request? A classifier scores the request, and if it's an easy one directs it to a cheap model while a hard one goes
to a frontier model. For a chatbot answering one-off questions independently, that is the whole characterization of the problem.
For a coding agent it is only half of it. A coding session is a long conversation in which the agent
resends everything said so far with every request, and LLM providers cache that history so they need
not reprocess it each time. A router that picks the right model for every turn, but ignores the
cache, throws it away at each change of model and pays to rebuild it. The second half of the problem
is therefore when a running session should change model. That is an online problem, decided turn
by turn without knowing the future, between a one-off cost and a stream of small savings, and it has
a classic form: the ski-rental problem.
Edgee's Agent Gateway treats routing as both halves. A tier map assigns a model to each level of
difficulty, and a switching policy based on ski rental decides when a session actually moves.
Self-Learning Routing learns both from an organization's own traffic. The gateway labels every turn
as it serves it. Once a week the control plane replays the past 30 days of sessions through a copy of
the gateway's algorithm, under every candidate map and several switching settings, and measures each
against the schedule that hindsight would have chosen. It recommends the cheapest combination that
keeps every turn on a capable model, and applies it when a series of guardrails allow.
Why difficulty is only half the story
Picking a model once per session is easy, and wasteful: a session that opens with a hard design
question often spends the next hour on edits a far cheaper model could make. The savings come from
moving mid-task. But every move has a price.
On Anthropic, caching a prefix for five minutes costs 1.25 times the normal input price, and reading
it back costs a tenth ( prompt caching ).
A session that changes model loses its warm cache. The new model reads the whole conversation cold
and writes it to its own cache. For a long session, that one rewrite can cost more than many turns of
the price difference between the two models.
So the two halves interact. A tier map that puts a cheap model on easy turns only saves money if
sessions actually move down to it, and stay long enough to recoup the rewrite. Switch on every turn
that looks easy and the cache rewrites eat the savings. Never switch and the cheap model is never
used. Which map is best therefore depends on the switching policy, and the switching policy that is
best depends on the map. Self-Learning Routing learns them together for that reason.
At the start of each user turn the gateway asks a classifier one question about the latest message:
how much capability does this task need? The answer comes on a four-point scale, from a trivial edit
to the hardest architectural problem, as a probability-weighted level index from 0 to 3. The gateway
normalizes it to a score in [0, 1] and maps the score to a tier:
The requests inside a turn, as the agent calls tools and reads their output, never decide afresh:
they stay on the model the turn began with.
Moving up a tier is a constraint, not a bet. A turn that needs a more capable model gets one, at once
and whatever the switching cost: quality is not traded for money. The only filter is the
classifier's confidence, so that a single noisy label cannot drag a session upmarket:
Upgrades sit outside the cost accounting that follows. Everything else, moving to a cheaper model or
to a different model for the same tier, is a cost decision, and that is the second half.
Half two: switching as a ski-rental problem
An algorithm is online when it must act on each part of its input as it arrives, without seeing
what comes next. Its offline counterpart is handed the whole input in advance. Caching is the
classic example: a full cache must evict an entry now, and only later requests reveal which entry
was not needed again.
Online algorithms are judged by competitive analysis (Sleator and Tarjan, Amortized efficiency of
list update and paging rules , Communications of the ACM, 1985). An online algorithm is
c-competitive if, on every input, it pays at most c times what the best offline algorithm pays on
that same input. The input may be chosen by an adversary who knows the algorithm, so the ratio is a
promise about the worst case, not the average one.
Routing a coding session is an online problem: the gateway must pick a model for this turn before the
developer has written the next one. Half one barely notices, since a turn's tier depends only on that
turn. Switching does. Whether a move pays depends on the past, which decides where the warm cache is,
and on the future, which decides how long the cheaper model will stay the right one. The gateway
knows only the past.
The simplest online problem with that shape is ski rental. A skier can rent skis for a fee each day
or buy a pair for B , without knowing how many days the season will last. Renting through a long
season costs far more than the skis; buying on the first day wastes the money if the season ends the
next morning. The problem was first analyzed in this form for caches too: multiprocessor caches
deciding whether to keep paying to update a shared block or to drop it (Karlin, Manasse, Rudolph and
Sleator, Competitive snoopy caching , Algorithmica, 1988).
Why break-even is 2-competitive
Let S be what renting would cost over the whole season. The offline optimum knows S , so it pays
min(S, B) : it rents through a short season and buys on the first day of a long one. The break-even
rule rents until the rent paid reaches B , then buys. (Rent is treated as continuous, a fair
approximation when a day's rent is small next to the price of skis, as a turn's saving is next to a
cache rewrite.)
If the season ends before the rent paid reaches B , the rule has rented throughout, as the optimum
did, and pays S .
Otherwise it pays B in rent, then B for the skis: 2B , against the optimum's B .
Either way it pays at most twice the optimum. Its worst case is the one an adversary would pick: the
season ends the moment the skis are bought.
The same adversary shows that no deterministic rule does better. Any such rule amounts to buying once
the rent paid reaches x·B , for some x . End the season just after the purchase: the rule pays
(1 + x)·B and the optimum min(x, 1)·B , a ratio of
1 + max(x, 1/x)
That is also the worst any season can do to the rule, so it is the rule's exact competitive ratio. It
is smallest at x = 1 , where it equals 2. Buying early loses to a short season, buying late loses to
a long one, and break-even balances the two regrets. Randomizing the moment of purchase lowers the
bound to e/(e−1) ≈ 1.58 (Karlin, Manasse, McGeoch and Owicki, Competitive randomized algorithms
for nonuniform problems , Algorithmica, 1994). The gateway stays deterministic: the same session
always gets the same decisions, which is what lets the replay re-run them.
When a cheaper model becomes the right one for a session's turns, the gateway faces the skier's
choice:
The fit is close but not exact, and each difference shaped the design:
The price of skis rises. A switch rewrites the whole conversation, which usually grows with
every turn, so the longer the gateway waits, the dearer buying gets. The rent moves too, with the
length of the answers. The gateway re-prices both on every turn. The proof above assumes fixed
prices, so in a real session it is a guide rather than a theorem; the replay measures what the
rule actually leaves on the table.
Seasons repeat. A session holds many stretches, not one. When the candidate changes, because
the turn's tier or the Strategy's map changed, a new instance begins, and the evidence gathered
against the old candidate stops counting.
The skis can be taken away. A hard turn forces an upgrade whatever it costs, which ends the
stretch; coming back down means buying again. A threshold above break-even guards against exactly
this.
Sometimes the skis are free. When the cache has expired or the conversation has just been
compacted, the current model must rewrite the conversation anyway, so switching costs nothing
extra.
For each turn the gateway estimates four token counts: P , the conversation prefix already warm on
the current model p ; N , the tokens new this turn; O and R , the expected output and reasoning
tokens, taken from running averages of the session's own past turns. With a model's rates for cache
writes w , cache reads h , output o and reasoning ρ , and c the cheaper candidate:
cold(m) = (P + N)·w(m) + O·o(m) + R·ρ(m) # nothing cached: the whole prefix is written
warm(m) = P·h(m) + N·w(m) + O·o(m) + R·ρ(m) # the prefix is a cache hit
buy B = max(0, cold(c) − warm(p)) # one-off premium of switching this turn
rent r = max(0, warm(p) − warm(c)) # saving forgone per turn by staying
The break-even rule in the gateway
The gateway keeps a running total A of the rent paid against the current candidate, and buys when
it reaches a multiple of the buy price:
A ← A + r on every turn the candidate stays the same
switch when A ≥ λ·B (λ = 1.5 by default)
With λ = 1 this is the rule proved above. If a cold start on the candidate already costs less than
staying warm ( B = 0 ), the switch happens at once.
The guarantee has an honest scope. It holds for one stretch, with a stable candidate, no upgrade and
fixed prices, against that pair of models. It is not a bound against the best schedule for the whole
session; that comparison is the replay's job, below.
By the argument above, a threshold λ has a worst-case ratio of 1 + max(λ, 1/λ) : 2.5 at
λ = 1.5 . The gateway accepts the weaker bound because its seasons can end in a way the textbook's
cannot. Upgrades fire regardless of cost, so a downgrade bought at exactly break-even can be undone
by a hard task on the very next turn, paying a second rewrite for nothing. Waiting until the rent
reaches one and a half times the buy price demands more evidence that the easy stretch is real. Which
threshold does best on average depends on how long an organization's easy stretches tend to last,
which no worst-case argument can tell; the replay measures it.
Forgetting without getting stuck
A router should forgive an isolated blip, so the running total decays. Decay it on every turn,
though, and a subtle bug appears. With a constant rent r and a decay factor d , the recurrence
A ← d·A + r converges to r / (1 − d) . If the buy price exceeds that ceiling the total can never
cross the threshold, and the session overpays forever with no way out. The gateway therefore decays
the total (by half per turn) only once the session already sits on its target model. While evidence
is building against a stable candidate, the total only grows, so a long enough run of savings always
triggers the switch. When the candidate itself changes, a new instance of the problem starts and the
total resets to zero.
Free switches, and the order of checks
A switch is free when the cache has expired since the session's last request, or when the
conversation has just shrunk below 60% of its previous length, as when Claude Code compacts it. The
gateway assumes a cache lives five minutes, the shorter of Anthropic's two lifetimes, which can only
make it treat a warm cache as cold early, never the reverse. After any switch, the gateway waits four
turns before it will switch again to save money, which bounds how often a session changes model
whatever the totals

[truncated]

## Original Extract

Most LLM routers answer one question: which model should handle this request? For a coding agent that is only half the story. Every switch throws away the prompt cache, so deciding when a session should change model is a ski-rental problem. Self-Learning Routing learns both halves from an organizati
[truncated]

Switching LLMs mid-conversation is a ski-rental problem - Edgee Blog Features Smart Routing & Budgets Token Compression Team Observability Integrations Plugins Claude Mod On-Premise Edgee vs OpenRouter Edgee vs LiteLLM Models Blog Pricing Documentation Login Open main menu Switching LLMs mid-conversation is a ski-rental problem
Most work on routing between language models answers one question: which model should handle this
request? A classifier scores the request, and if it's an easy one directs it to a cheap model while a hard one goes
to a frontier model. For a chatbot answering one-off questions independently, that is the whole characterization of the problem.
For a coding agent it is only half of it. A coding session is a long conversation in which the agent
resends everything said so far with every request, and LLM providers cache that history so they need
not reprocess it each time. A router that picks the right model for every turn, but ignores the
cache, throws it away at each change of model and pays to rebuild it. The second half of the problem
is therefore when a running session should change model. That is an online problem, decided turn
by turn without knowing the future, between a one-off cost and a stream of small savings, and it has
a classic form: the ski-rental problem.
Edgee's Agent Gateway treats routing as both halves. A tier map assigns a model to each level of
difficulty, and a switching policy based on ski rental decides when a session actually moves.
Self-Learning Routing learns both from an organization's own traffic. The gateway labels every turn
as it serves it. Once a week the control plane replays the past 30 days of sessions through a copy of
the gateway's algorithm, under every candidate map and several switching settings, and measures each
against the schedule that hindsight would have chosen. It recommends the cheapest combination that
keeps every turn on a capable model, and applies it when a series of guardrails allow.
Why difficulty is only half the story
Picking a model once per session is easy, and wasteful: a session that opens with a hard design
question often spends the next hour on edits a far cheaper model could make. The savings come from
moving mid-task. But every move has a price.
On Anthropic, caching a prefix for five minutes costs 1.25 times the normal input price, and reading
it back costs a tenth ( prompt caching ).
A session that changes model loses its warm cache. The new model reads the whole conversation cold
and writes it to its own cache. For a long session, that one rewrite can cost more than many turns of
the price difference between the two models.
So the two halves interact. A tier map that puts a cheap model on easy turns only saves money if
sessions actually move down to it, and stay long enough to recoup the rewrite. Switch on every turn
that looks easy and the cache rewrites eat the savings. Never switch and the cheap model is never
used. Which map is best therefore depends on the switching policy, and the switching policy that is
best depends on the map. Self-Learning Routing learns them together for that reason.
At the start of each user turn the gateway asks a classifier one question about the latest message:
how much capability does this task need? The answer comes on a four-point scale, from a trivial edit
to the hardest architectural problem, as a probability-weighted level index from 0 to 3. The gateway
normalizes it to a score in [0, 1] and maps the score to a tier:
The requests inside a turn, as the agent calls tools and reads their output, never decide afresh:
they stay on the model the turn began with.
Moving up a tier is a constraint, not a bet. A turn that needs a more capable model gets one, at once
and whatever the switching cost: quality is not traded for money. The only filter is the
classifier's confidence, so that a single noisy label cannot drag a session upmarket:
Upgrades sit outside the cost accounting that follows. Everything else, moving to a cheaper model or
to a different model for the same tier, is a cost decision, and that is the second half.
Half two: switching as a ski-rental problem
An algorithm is online when it must act on each part of its input as it arrives, without seeing
what comes next. Its offline counterpart is handed the whole input in advance. Caching is the
classic example: a full cache must evict an entry now, and only later requests reveal which entry
was not needed again.
Online algorithms are judged by competitive analysis (Sleator and Tarjan, Amortized efficiency of
list update and paging rules , Communications of the ACM, 1985). An online algorithm is
c-competitive if, on every input, it pays at most c times what the best offline algorithm pays on
that same input. The input may be chosen by an adversary who knows the algorithm, so the ratio is a
promise about the worst case, not the average one.
Routing a coding session is an online problem: the gateway must pick a model for this turn before the
developer has written the next one. Half one barely notices, since a turn's tier depends only on that
turn. Switching does. Whether a move pays depends on the past, which decides where the warm cache is,
and on the future, which decides how long the cheaper model will stay the right one. The gateway
knows only the past.
The simplest online problem with that shape is ski rental. A skier can rent skis for a fee each day
or buy a pair for B , without knowing how many days the season will last. Renting through a long
season costs far more than the skis; buying on the first day wastes the money if the season ends the
next morning. The problem was first analyzed in this form for caches too: multiprocessor caches
deciding whether to keep paying to update a shared block or to drop it (Karlin, Manasse, Rudolph and
Sleator, Competitive snoopy caching , Algorithmica, 1988).
Why break-even is 2-competitive
Let S be what renting would cost over the whole season. The offline optimum knows S , so it pays
min(S, B) : it rents through a short season and buys on the first day of a long one. The break-even
rule rents until the rent paid reaches B , then buys. (Rent is treated as continuous, a fair
approximation when a day's rent is small next to the price of skis, as a turn's saving is next to a
cache rewrite.)
If the season ends before the rent paid reaches B , the rule has rented throughout, as the optimum
did, and pays S .
Otherwise it pays B in rent, then B for the skis: 2B , against the optimum's B .
Either way it pays at most twice the optimum. Its worst case is the one an adversary would pick: the
season ends the moment the skis are bought.
The same adversary shows that no deterministic rule does better. Any such rule amounts to buying once
the rent paid reaches x·B , for some x . End the season just after the purchase: the rule pays
(1 + x)·B and the optimum min(x, 1)·B , a ratio of
1 + max(x, 1/x)
That is also the worst any season can do to the rule, so it is the rule's exact competitive ratio. It
is smallest at x = 1 , where it equals 2. Buying early loses to a short season, buying late loses to
a long one, and break-even balances the two regrets. Randomizing the moment of purchase lowers the
bound to e/(e−1) ≈ 1.58 (Karlin, Manasse, McGeoch and Owicki, Competitive randomized algorithms
for nonuniform problems , Algorithmica, 1994). The gateway stays deterministic: the same session
always gets the same decisions, which is what lets the replay re-run them.
When a cheaper model becomes the right one for a session's turns, the gateway faces the skier's
choice:
The fit is close but not exact, and each difference shaped the design:
The price of skis rises. A switch rewrites the whole conversation, which usually grows with
every turn, so the longer the gateway waits, the dearer buying gets. The rent moves too, with the
length of the answers. The gateway re-prices both on every turn. The proof above assumes fixed
prices, so in a real session it is a guide rather than a theorem; the replay measures what the
rule actually leaves on the table.
Seasons repeat. A session holds many stretches, not one. When the candidate changes, because
the turn's tier or the Strategy's map changed, a new instance begins, and the evidence gathered
against the old candidate stops counting.
The skis can be taken away. A hard turn forces an upgrade whatever it costs, which ends the
stretch; coming back down means buying again. A threshold above break-even guards against exactly
this.
Sometimes the skis are free. When the cache has expired or the conversation has just been
compacted, the current model must rewrite the conversation anyway, so switching costs nothing
extra.
For each turn the gateway estimates four token counts: P , the conversation prefix already warm on
the current model p ; N , the tokens new this turn; O and R , the expected output and reasoning
tokens, taken from running averages of the session's own past turns. With a model's rates for cache
writes w , cache reads h , output o and reasoning ρ , and c the cheaper candidate:
cold(m) = (P + N)·w(m) + O·o(m) + R·ρ(m) # nothing cached: the whole prefix is written
warm(m) = P·h(m) + N·w(m) + O·o(m) + R·ρ(m) # the prefix is a cache hit
buy B = max(0, cold(c) − warm(p)) # one-off premium of switching this turn
rent r = max(0, warm(p) − warm(c)) # saving forgone per turn by staying
The break-even rule in the gateway
The gateway keeps a running total A of the rent paid against the current candidate, and buys when
it reaches a multiple of the buy price:
A ← A + r on every turn the candidate stays the same
switch when A ≥ λ·B (λ = 1.5 by default)
With λ = 1 this is the rule proved above. If a cold start on the candidate already costs less than
staying warm ( B = 0 ), the switch happens at once.
The guarantee has an honest scope. It holds for one stretch, with a stable candidate, no upgrade and
fixed prices, against that pair of models. It is not a bound against the best schedule for the whole
session; that comparison is the replay's job, below.
By the argument above, a threshold λ has a worst-case ratio of 1 + max(λ, 1/λ) : 2.5 at
λ = 1.5 . The gateway accepts the weaker bound because its seasons can end in a way the textbook's
cannot. Upgrades fire regardless of cost, so a downgrade bought at exactly break-even can be undone
by a hard task on the very next turn, paying a second rewrite for nothing. Waiting until the rent
reaches one and a half times the buy price demands more evidence that the easy stretch is real. Which
threshold does best on average depends on how long an organization's easy stretches tend to last,
which no worst-case argument can tell; the replay measures it.
Forgetting without getting stuck
A router should forgive an isolated blip, so the running total decays. Decay it on every turn,
though, and a subtle bug appears. With a constant rent r and a decay factor d , the recurrence
A ← d·A + r converges to r / (1 − d) . If the buy price exceeds that ceiling the total can never
cross the threshold, and the session overpays forever with no way out. The gateway therefore decays
the total (by half per turn) only once the session already sits on its target model. While evidence
is building against a stable candidate, the total only grows, so a long enough run of savings always
triggers the switch. When the candidate itself changes, a new instance of the problem starts and the
total resets to zero.
Free switches, and the order of checks
A switch is free when the cache has expired since the session's last request, or when the
conversation has just shrunk below 60% of its previous length, as when Claude Code compacts it. The
gateway assumes a cache lives five minutes, the shorter of Anthropic's two lifetimes, which can only
make it treat a warm cache as cold early, never the reverse. After any switch, the gateway waits four
turns before it will switch again to save money, which bounds how often a session changes model
whatever the totals

[truncated]
