---
source: "https://www.frisson-labs.com/reyn"
hn_url: "https://news.ycombinator.com/item?id=50524802"
title: "We got AI to Wombo Combo – collaborative agents in Melee"
article_title: "Reyn — A language-steerable Melee player"
image: "https://www.frisson-labs.com/static/frisson/assets/social/articles/v1/reyn.jpg"
author: "coatol5"
captured_at: "2026-11-16T23:41:55Z"
capture_tool: "hn-digest"
hn_id: 50524802
score: 3
comments: 2
posted_at: "2026-11-16T23:03:19Z"
tags:
  - hacker-news
---

# We got AI to Wombo Combo – collaborative agents in Melee

- HN: [50524802](https://news.ycombinator.com/item?id=50524802)
- Source: [www.frisson-labs.com](https://www.frisson-labs.com/reyn)
- Score: 3
- Comments: 2
- Posted: 2026-11-16T23:03:19Z

## Translation

Title: We got AI to Wombo Combo – collaborative agents in Melee
Article title: Reyn — A language-steerable Melee player
Description: From Faynt to Reyn: language-steerable Melee, cooperative play, and the experiments behind a teammate you can talk to.

Article text:
Reyn — A language-steerable Melee player
Skip to article
frisson labs Home Research Careers Models Get in touch Home Research Careers Models Get in touch
Reyn
A language-steerable Melee player you can talk to.
A language-steerable Melee player you can talk to.
· Ali Janati · Tongke Xue November 11, 2026 15 minutes read Advisory contributions: Rohit Swamy · Charles Niu
Reyn1 and Reyn2 · two level-5 CPUs
Last month, we introduced Faynt , a family of open-weight models that can play Super Smash Bros. Melee. We trained it on approximately 840,000 human replays and refined it through reinforcement learning (RL).
Faynt could compete in one-on-one (1v1) matches, but we could not steer it with commands. What if it could also play two-versus-two (2v2)?
And to take it even further, what if it could become a language-steerable player you could coordinate with in real time, in both 1v1 and 2v2?
Would it still be able to fight back while following your commands?
Would it be able to coordinate with a human player?
What about friendly fire? How would it avoid hitting its teammate?
Enter Reyn 26C , our language-steerable evolution of Faynt. We started with Fox, expanded to four characters, and now steer one shared model across all 26 . Reyn 4C remains a development baseline in the comparisons below. Our Fox team and combo lab are separate experiments.
In the opening video, we pit two Reyn agents against two in-game bots. Each agent receives its own steering commands, displayed above its character, and adapts its gameplay in real time. The second part is a separate combo-lab relay. Both are selected examples.
In this article, we’ll go in depth about how we created Reyn, including:
Language steering: pre-training and RL, starting with 1v1 matches.
Voice: a conversational pipeline turns speech into commands for Reyn and spoken replies about the game.
2v2: bringing steering to team games and dealing with friendly fire.
Bonus: training a Wombo Combo 2v1 model, just for fun.
Steering: Reyn 26C on the original command suite against mirror level-9 CPUs. The team-game and teammate-damage results come from separate Fox experiments; the latter uses a human proxy.
Build a Melee model you can steer with natural language.
Push what a small model can do, and extend steering to all 26 characters.
Keep the model able to fight back while following your commands.
We are not trying to build the most competitive win-at-all-costs Melee bot or a standalone voice companion. The goal is a teammate that can still fight while responding to a person's instructions.
We started with Fox to keep the first steering experiments simple, then expanded to four characters. To reach the full roster, we tried reusing that gameplay stack and teaching its text router more requests. That became Reyn 26C . It handles special-move requests across the roster, although quality still varies by character. The experimental 2v2 branch currently uses Fox.
The plan was to make Faynt steerable, bring that steering to 2v2, then build toward a partner for a person.
Starting with language steering
We started by adding language steering to Faynt, with as little loss of playing strength as possible.
Following commands while still fighting
A request to grab, move somewhere or stop attacking changes which actions are useful. We wanted those requests to change Reyn's behaviour without making it ineffective at the game. To teach those behaviours, we began with Faynt's weights and added language labels to human replays. Rules called predicates identify actions or situations and pair them with matching instructions. Given an instruction and the game history, Reyn learns to predict the controller action from the replay. This imitation stage trains the policy network, cross-attention adapter and original text encoder together.
Asked to Grab him. , it tried more grabs, but few connected. We then used RL to reward grabs that landed.
RL also lets us reward a hit on the opponent or arrival at the requested platform. It gives the model practice with behaviours underrepresented in the labelled replays, such as moving to platforms. Imitation connects instructions to actions; RL improves what those actions achieve.
We developed the gameplay stack in Reyn 4C . After imitation learning, we froze the 10 million -parameter gameplay policy and trained its cross-attention adapter, command embeddings, state adapter and taunt head with RL. The state adapter reads 16 game features; the taunt head adds an input the frozen policy cannot produce. About 2.2 million parameters trained in this stage. With no command, the steering components are bypassed.
We train these parts with proximal policy optimization (PPO). Kullback–Leibler (KL) divergence measures how differently two policies assign probabilities to actions in the same game situation. We use a KL penalty to discourage large changes from the no-command policy while learning the requested behaviour. Specialists learn groups of commands, then distillation combines what they learned into one model. The Reyn 4C model then received a joint RL round across all six stages.
For Reyn 26C , we kept those gameplay weights fixed and fine-tuned only the text router on requests across the roster, including different ways to ask for the same move. It learns which supported command the words mean. This supervised training is separate from the RL that taught the gameplay response.
Scroll sideways to see the whole figure.
imitation
steering RL
router fitting
data / fixed rules
fixed during steering RL
All parameters stay fixed while playing.
ONCE PER REQUEST
EVERY GAME FRAME
Your request
“Go to the left platform.”
Text router
Supervised language training
Paraphrases → command meanings
Resolve the command
meaning + character + stage
command ID
Look up command tokens
Learned command tokens
…
or fixed fallback encodings
Game state
plus the previous action
Frame encoder
game state → features
Causal Transformer
the recent game history
becomes a hidden state
Hidden state h
Game features
positions, gaps, damage, …
h
STEERING PATH
Route the steering inputs
command tokens · command ID · hidden state h · game features
Each branch below names the signals that its module reads.
tokens + h
ID + h + game features
ID + h + game features
Cross-attention
296 thousand
parameters
adds a residual
State adapter
learned ID embedding
adds a residual
Taunt head
learned ID embedding
taunt command only
residual
residual
h
conditioned hidden state
Action head
buttons, then stick
taunt gate
Controller action
taunt action
when gated
No command or cancel: both residuals are zero; h reaches the action head unchanged.
Base policy: 10 million parameters in the encoder, Transformer and action head.
A pretrained text encoder builds initial command tokens offline.
How Reyn plays after RL Runtime after steering RL. All parameters are fixed during play. A separately fitted text router resolves requests by character and stage to a command ID; stored tokens are looked up once per request. Each frame, the encoder and causal Transformer produce hidden state h. Cross-attention reads h and command tokens. The state adapter and taunt head each read h, game features, and their own command-ID embedding. Cross-attention and state residuals are added to h before the original action head. A separate taunt gate sends D_UP to the controller. No command or cancel bypasses steering. The initial text encoder is offline and distinct from
[truncated]
trained
frozen weights
data / fixed rules
1 · Imitation from labelled replays
Predicates pair observed replay situations with matching instructions.
SOURCE
OBJECTIVE
WHAT TRAINS
Labelled human play
Instruction + state
recorded actions
Imitate the actions
Controller negative
log-likelihood
of recorded actions
All imitation modules
Game encoder + Transformer
Original action head
Cross-attention adapter
Large text encoder
The large text encoder is distinct from the later request router.
Replay predicates label training examples; the learned policy chooses actions.
initialize steering from the trained policy
2 · Specialist reinforcement learning (RL)
Continue cross-attention; add steering modules and learned command embeddings.
SOURCE
OBJECTIVE
WHAT TRAINS
Command practice
Requested behaviors
gameplay situations
Policy optimization
Proximal policy
optimization (PPO)
game rewards +
retention / teacher loss
Steering modules
Inherited cross-attention
State adapter
Command embeddings
Taunt head, when enabled
Specialists → distillation → joint refinement
Specialists
Command groups
gameplay rewards
Fixed teachers
trained specialists
Shared student
reward + teacher loss
Joint RL
Across stages
retain earlier skills
Fixed during this stage
Imitation-trained game encoder, Transformer and original action head.
Training-only critics learn values; they are not used to play.
Benchmark thresholds grade outcomes; behavior rewards guide learning.
3 · Router fine-tuning on paraphrases
A separate text model learns which supported meaning the wording requests.
SOURCE
OBJECTIVE
WHAT TRAINS
Labelled wordings
Request + meaning
or cancel / decline
Supervised fitting
Cross-entropy
on command labels
no gameplay reward
Text router
Own encoder + classifier
separate from policy RL
Gameplay policy fixed during router fitting
Router fitted separately; later policy RL reuses it.
How Reyn is trained
Teal marks imitation training, orange gameplay RL, and purple the router’s separate supervised fit. Firs
[truncated]
Fire Fox sends Fox flying, so a badly aimed launch can throw away a stock. Our first specialist landed hits but also self-destructed. We increased the self-destruct penalty, and it started avoiding using the Fire Fox move instead. This was reward hacking: avoiding the requested move was an easier way to avoid the penalty than learning to use it safely.
So we changed the reward to catch bad launches sooner. If a launch looked likely to send Fox too far past the stage edge, it got a penalty straight away and no reward for that launch, even if it landed a hit. We then made that check more cautious and rewarded aiming at the opponent. In the later test, the specialist still landed 5.2 Fire Fox hits per minute, with 0.08 self-destructs per minute.
Fire Fox is one specialist example. We distilled the specialists into Reyn 4C and refined it with joint RL. The later router update produced Reyn 26C , which meets at least the looser target in 386 of 424 original-suite cases: 267 meet the stricter target and 119 meet only the looser one. The comparison below shows both grades.
To tell Reyn what to do, we first wanted a small, fast path from words to a command. During Reyn 4C development, predicate-labelled pre-training used a 0.6 billion -parameter text encoder beside a 10 million -parameter gameplay policy. Running that encoder for each request would add latency and make the language machinery much larger than the model playing the game.
A smaller router recognizes a request and selects a stored command embedding. Reyn 26C extends that approach with a router trained on more character moves and varied phrasing. It reuses the gameplay weights from Reyn 4C : the extra character coverage comes from the language path, rather than another round of gameplay RL. In the additional text-only wording test, it routed 100 % of 208 generated test rows correctly.
We then tried more compact command representations and lighter routers. Every setup went through the same 424 test cases with held-out wordings. The question was practical: how much size and request time could we cut without losing command following?
In the Reyn 4C development experiment, the smaller-router setup cut the text-side files from 164.1 MB to 28.8 MB. On this held-out-wording benchmark, it met at least the looser command-following target in 88.2 % of cases, compared with 88.4 % for the original setup. Median text-request time fell from 13.84 ms to 4.09 ms. Compressing the command representation to a single token hurt comman

[truncated]

## Original Extract

From Faynt to Reyn: language-steerable Melee, cooperative play, and the experiments behind a teammate you can talk to.

Reyn — A language-steerable Melee player
Skip to article
frisson labs Home Research Careers Models Get in touch Home Research Careers Models Get in touch
Reyn
A language-steerable Melee player you can talk to.
A language-steerable Melee player you can talk to.
· Ali Janati · Tongke Xue November 11, 2026 15 minutes read Advisory contributions: Rohit Swamy · Charles Niu
Reyn1 and Reyn2 · two level-5 CPUs
Last month, we introduced Faynt , a family of open-weight models that can play Super Smash Bros. Melee. We trained it on approximately 840,000 human replays and refined it through reinforcement learning (RL).
Faynt could compete in one-on-one (1v1) matches, but we could not steer it with commands. What if it could also play two-versus-two (2v2)?
And to take it even further, what if it could become a language-steerable player you could coordinate with in real time, in both 1v1 and 2v2?
Would it still be able to fight back while following your commands?
Would it be able to coordinate with a human player?
What about friendly fire? How would it avoid hitting its teammate?
Enter Reyn 26C , our language-steerable evolution of Faynt. We started with Fox, expanded to four characters, and now steer one shared model across all 26 . Reyn 4C remains a development baseline in the comparisons below. Our Fox team and combo lab are separate experiments.
In the opening video, we pit two Reyn agents against two in-game bots. Each agent receives its own steering commands, displayed above its character, and adapts its gameplay in real time. The second part is a separate combo-lab relay. Both are selected examples.
In this article, we’ll go in depth about how we created Reyn, including:
Language steering: pre-training and RL, starting with 1v1 matches.
Voice: a conversational pipeline turns speech into commands for Reyn and spoken replies about the game.
2v2: bringing steering to team games and dealing with friendly fire.
Bonus: training a Wombo Combo 2v1 model, just for fun.
Steering: Reyn 26C on the original command suite against mirror level-9 CPUs. The team-game and teammate-damage results come from separate Fox experiments; the latter uses a human proxy.
Build a Melee model you can steer with natural language.
Push what a small model can do, and extend steering to all 26 characters.
Keep the model able to fight back while following your commands.
We are not trying to build the most competitive win-at-all-costs Melee bot or a standalone voice companion. The goal is a teammate that can still fight while responding to a person's instructions.
We started with Fox to keep the first steering experiments simple, then expanded to four characters. To reach the full roster, we tried reusing that gameplay stack and teaching its text router more requests. That became Reyn 26C . It handles special-move requests across the roster, although quality still varies by character. The experimental 2v2 branch currently uses Fox.
The plan was to make Faynt steerable, bring that steering to 2v2, then build toward a partner for a person.
Starting with language steering
We started by adding language steering to Faynt, with as little loss of playing strength as possible.
Following commands while still fighting
A request to grab, move somewhere or stop attacking changes which actions are useful. We wanted those requests to change Reyn's behaviour without making it ineffective at the game. To teach those behaviours, we began with Faynt's weights and added language labels to human replays. Rules called predicates identify actions or situations and pair them with matching instructions. Given an instruction and the game history, Reyn learns to predict the controller action from the replay. This imitation stage trains the policy network, cross-attention adapter and original text encoder together.
Asked to Grab him. , it tried more grabs, but few connected. We then used RL to reward grabs that landed.
RL also lets us reward a hit on the opponent or arrival at the requested platform. It gives the model practice with behaviours underrepresented in the labelled replays, such as moving to platforms. Imitation connects instructions to actions; RL improves what those actions achieve.
We developed the gameplay stack in Reyn 4C . After imitation learning, we froze the 10 million -parameter gameplay policy and trained its cross-attention adapter, command embeddings, state adapter and taunt head with RL. The state adapter reads 16 game features; the taunt head adds an input the frozen policy cannot produce. About 2.2 million parameters trained in this stage. With no command, the steering components are bypassed.
We train these parts with proximal policy optimization (PPO). Kullback–Leibler (KL) divergence measures how differently two policies assign probabilities to actions in the same game situation. We use a KL penalty to discourage large changes from the no-command policy while learning the requested behaviour. Specialists learn groups of commands, then distillation combines what they learned into one model. The Reyn 4C model then received a joint RL round across all six stages.
For Reyn 26C , we kept those gameplay weights fixed and fine-tuned only the text router on requests across the roster, including different ways to ask for the same move. It learns which supported command the words mean. This supervised training is separate from the RL that taught the gameplay response.
Scroll sideways to see the whole figure.
imitation
steering RL
router fitting
data / fixed rules
fixed during steering RL
All parameters stay fixed while playing.
ONCE PER REQUEST
EVERY GAME FRAME
Your request
“Go to the left platform.”
Text router
Supervised language training
Paraphrases → command meanings
Resolve the command
meaning + character + stage
command ID
Look up command tokens
Learned command tokens
…
or fixed fallback encodings
Game state
plus the previous action
Frame encoder
game state → features
Causal Transformer
the recent game history
becomes a hidden state
Hidden state h
Game features
positions, gaps, damage, …
h
STEERING PATH
Route the steering inputs
command tokens · command ID · hidden state h · game features
Each branch below names the signals that its module reads.
tokens + h
ID + h + game features
ID + h + game features
Cross-attention
296 thousand
parameters
adds a residual
State adapter
learned ID embedding
adds a residual
Taunt head
learned ID embedding
taunt command only
residual
residual
h
conditioned hidden state
Action head
buttons, then stick
taunt gate
Controller action
taunt action
when gated
No command or cancel: both residuals are zero; h reaches the action head unchanged.
Base policy: 10 million parameters in the encoder, Transformer and action head.
A pretrained text encoder builds initial command tokens offline.
How Reyn plays after RL Runtime after steering RL. All parameters are fixed during play. A separately fitted text router resolves requests by character and stage to a command ID; stored tokens are looked up once per request. Each frame, the encoder and causal Transformer produce hidden state h. Cross-attention reads h and command tokens. The state adapter and taunt head each read h, game features, and their own command-ID embedding. Cross-attention and state residuals are added to h before the original action head. A separate taunt gate sends D_UP to the controller. No command or cancel bypasses steering. The initial text encoder is offline and distinct from
[truncated]
trained
frozen weights
data / fixed rules
1 · Imitation from labelled replays
Predicates pair observed replay situations with matching instructions.
SOURCE
OBJECTIVE
WHAT TRAINS
Labelled human play
Instruction + state
recorded actions
Imitate the actions
Controller negative
log-likelihood
of recorded actions
All imitation modules
Game encoder + Transformer
Original action head
Cross-attention adapter
Large text encoder
The large text encoder is distinct from the later request router.
Replay predicates label training examples; the learned policy chooses actions.
initialize steering from the trained policy
2 · Specialist reinforcement learning (RL)
Continue cross-attention; add steering modules and learned command embeddings.
SOURCE
OBJECTIVE
WHAT TRAINS
Command practice
Requested behaviors
gameplay situations
Policy optimization
Proximal policy
optimization (PPO)
game rewards +
retention / teacher loss
Steering modules
Inherited cross-attention
State adapter
Command embeddings
Taunt head, when enabled
Specialists → distillation → joint refinement
Specialists
Command groups
gameplay rewards
Fixed teachers
trained specialists
Shared student
reward + teacher loss
Joint RL
Across stages
retain earlier skills
Fixed during this stage
Imitation-trained game encoder, Transformer and original action head.
Training-only critics learn values; they are not used to play.
Benchmark thresholds grade outcomes; behavior rewards guide learning.
3 · Router fine-tuning on paraphrases
A separate text model learns which supported meaning the wording requests.
SOURCE
OBJECTIVE
WHAT TRAINS
Labelled wordings
Request + meaning
or cancel / decline
Supervised fitting
Cross-entropy
on command labels
no gameplay reward
Text router
Own encoder + classifier
separate from policy RL
Gameplay policy fixed during router fitting
Router fitted separately; later policy RL reuses it.
How Reyn is trained
Teal marks imitation training, orange gameplay RL, and purple the router’s separate supervised fit. Firs
[truncated]
Fire Fox sends Fox flying, so a badly aimed launch can throw away a stock. Our first specialist landed hits but also self-destructed. We increased the self-destruct penalty, and it started avoiding using the Fire Fox move instead. This was reward hacking: avoiding the requested move was an easier way to avoid the penalty than learning to use it safely.
So we changed the reward to catch bad launches sooner. If a launch looked likely to send Fox too far past the stage edge, it got a penalty straight away and no reward for that launch, even if it landed a hit. We then made that check more cautious and rewarded aiming at the opponent. In the later test, the specialist still landed 5.2 Fire Fox hits per minute, with 0.08 self-destructs per minute.
Fire Fox is one specialist example. We distilled the specialists into Reyn 4C and refined it with joint RL. The later router update produced Reyn 26C , which meets at least the looser target in 386 of 424 original-suite cases: 267 meet the stricter target and 119 meet only the looser one. The comparison below shows both grades.
To tell Reyn what to do, we first wanted a small, fast path from words to a command. During Reyn 4C development, predicate-labelled pre-training used a 0.6 billion -parameter text encoder beside a 10 million -parameter gameplay policy. Running that encoder for each request would add latency and make the language machinery much larger than the model playing the game.
A smaller router recognizes a request and selects a stored command embedding. Reyn 26C extends that approach with a router trained on more character moves and varied phrasing. It reuses the gameplay weights from Reyn 4C : the extra character coverage comes from the language path, rather than another round of gameplay RL. In the additional text-only wording test, it routed 100 % of 208 generated test rows correctly.
We then tried more compact command representations and lighter routers. Every setup went through the same 424 test cases with held-out wordings. The question was practical: how much size and request time could we cut without losing command following?
In the Reyn 4C development experiment, the smaller-router setup cut the text-side files from 164.1 MB to 28.8 MB. On this held-out-wording benchmark, it met at least the looser command-following target in 88.2 % of cases, compared with 88.4 % for the original setup. Median text-request time fell from 13.84 ms to 4.09 ms. Compressing the command representation to a single token hurt comman

[truncated]
