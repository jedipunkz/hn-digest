---
source: "https://www.doubleai.com/research/can-llms-redesign-themselves-as-they-train"
hn_url: "https://news.ycombinator.com/item?id=50398045"
title: "Can LLMs redesign themselves as they train?"
article_title: "Can LLMs redesign themselves as they train?"
image: "https://www.doubleai.com/api/media/file/code-layers-doubleai.jpg"
author: "bilsbie"
captured_at: "2026-11-06T21:18:53Z"
capture_tool: "hn-digest"
hn_id: 50398045
score: 4
comments: 0
posted_at: "2026-11-06T20:30:38Z"
tags:
  - hacker-news
---

# Can LLMs redesign themselves as they train?

- HN: [50398045](https://news.ycombinator.com/item?id=50398045)
- Source: [www.doubleai.com](https://www.doubleai.com/research/can-llms-redesign-themselves-as-they-train)
- Score: 4
- Comments: 0
- Posted: 2026-11-06T20:30:38Z

## Translation

Title: Can LLMs redesign themselves as they train?
Description: Architecture is the last hand-designed piece. We make it part of training.

Article text:
Can LLMs redesign themselves as they train? New Meet Maxwell, the first Code Execution Model. Research About Careers News Visit Maxwell Research About Careers News Visit Maxwell Research / Can LLMs redesign themselves as they train? Can LLMs redesign themselves as they train?
Architecture is the last hand-designed piece. We make it part of training.
01 TL;DR 02 Nanochat: where does a language model need help? 03 Evolution is mandatory 04 Taking on jaggedness 05 Closing the loop: DeepSeek improving itself 06 Looking ahead Related research
Publications · Oct 25, 2026 Alignment via Generalization
Publications · Oct 19, 2026 A New Record in Extremal Chess Composition
Publications · Oct 6, 2026 Advancing a decades-old problem in extremal graph theory
We put an AI agent inside a neural network's training run, where it redesigns the network on the fly.
The agent does what we call mechanistic intervention . It reads the network's internal state, much as mechanistic interpretability research does, but instead of explaining a finished network, it changes the network while it is still learning . The agent diagnoses what the network fails to learn and fixes the problem by changing the data , the objective , the update rule or the architecture itself .
This requires an entirely new toolkit , which the agent builds itself . On nanochat , the agent discovered one of its main instruments, which we call the phantom gradient : a measure of unmet demands.
The agent writes small programs and grows them into the network as code layers . The code layers are often “alien”. For example, one such layer computes gradients of hidden representations within a single forward pass and uses those to self-correct.
On CLRS-Text , a dataset of classical algorithms, the agent takes on jaggedness . We find that a transformer trained on this dataset handles the input sizes it was trained on, yet fails on the sizes in between, even when trained for twice as long. Code layers grown during training raise accuracy at those in-between sizes from 33% to 79% .
A model can also do this to itself. We let DeepSeek-V4.1-Flash act as the agent on itself . Using a code layer that reads its own hidden states, it raised its accuracy on AIME 2026 and HMMT 2026 from 64% to 71% .
Today, a transformer still transforms everything but itself. We are on the cusp of models that train and re-architect themselves , with people setting only the goal and the rules. Models will be able to carry themselves beyond the tasks they already know: they'll do science on themselves, testing ideas, experimenting and expanding their capabilities.
However long you train a neural network, only its weights change. Every other component, from the optimiser to the data mix and the architecture, remains static. What if, instead, we made everything learnable, including the architecture?
In this post, we find early signs that coding agents can redesign neural networks while they train . The resulting architectures can look rather alien. Below, we report early results on four experiments, each aimed at a basic limitation of today's transformers.
Recent work has let coding agents run research projects end to end , tune training code overnight , rewrite the update rules of networks as they train [Bai et al., 2026] and even edit their weights directly [Takehi et al., 2026] . We simply ask: why not let them take over every part of training at once, including even the architecture itself?
Transformers have well-known limitations. They often fail to generalise to inputs longer than those seen in training, and scale alone does not fix this [Anil et al., 2022] [Delétang et al., 2023] . They struggle with compositional tasks that take many steps [Dziri et al., 2023] . Even simple logic (like inferring "B is A" from "A is B") is unreliable [Berglund et al., 2023] .
Some of these limitations trace back to gradient descent itself. Gradient descent fits statistics extremely well, but it settles for whatever lowers the loss fastest, often shortcuts that hold only on the training data [Geirhos et al., 2020] . The representations it finds can be fractured and entangled [Kumar et al., 2025] . Code behaves differently. Programs apply the same rule to every input, seen or unseen, so they generalise by construction. Programs are also short. The shorter the description of a hypothesis, the fewer examples it needs to generalise [Blumer et al., 1987] , an idea formalised in the minimum description length principle [Rissanen, 1978] [Shalev-Shwartz and Ben-David, 2014] .
Coding agents can cover for gradient descent. When placed inside the training loop, the agents can find what the network fails to learn, work out why and write exact programs from a handful of failing examples. These programs go into the network itself, as new layers, which we call code layers . Such a change is beyond gradient descent, which moves in small, continuous steps. Rules can be written as code, but the countless statistical patterns hidden in massive data cannot. The agents can still use gradient descent to train everything else, including how much to rely on each code layer.
To decide where code layers are needed and what they should compute, the agents need an entirely new toolkit, which we call mechanistic intervention . Like mechanistic interpretability , these tools read the network's internals, but they do so during training, while the network can still adapt to what the agents change. These tools are still at an early stage. We call on the community to help develop them.
Finding demand (nanochat). We introduce the phantom gradient, the first mechanistic intervention tool discovered by the agent, and show how the agent uses it to find where a language model needs help. We then look at the layers it wrote, including one that lets the network correct itself using its own recent errors.
Evolution is mandatory (SWA). Comparing layers grown during training with the same layers planted at the start shows that timing matters, a new kind of training dynamic.
Jaggedness (CLRS-Text). We train a transformer on CLRS-Text, a benchmark of classical algorithms. Trained normally, the model solves problems only at the sizes it saw in training, failing at the sizes in between. Code layers grown during training close most of these gaps.
Closing the loop (DeepSeek). DeepSeek-V4.1-Flash acts as the agent on itself. Using a code layer it designed, which reads its own hidden states, it raises its accuracy on AIME 2026 and HMMT 2026 , without thinking, from 64% to 71%.
To avoid overfitting, we keep the test set hidden from the agent in every experiment. Beyond that, the agent has free rein to come up with ideas, design its own experiments and mechanistic intervention tools, and insert the code layers. We describe each experiment in detail below.
Nanochat: where does a language model need help?
We start by training nanochat [Karpathy, 2025], a small ChatGPT-style language model. We train it as usual: next-token prediction, with plain gradient descent. The network is a twelve-layer transformer of about 286 million parameters. 1
Can a coding agent, brought into the training loop, help the network learn more from the same data? In other words, can we improve sample efficiency?
We train the same network again, this time with a coding agent inside the training loop. The setup is identical (same data, schedule, steps and gradient-based updates) except that, at checkpoints of its choosing, the agent is allowed to insert new layers written as code, editing the architecture itself . Once a layer is inserted, ordinary training resumes from that point until the next insertion.
Over the span of this experiment, the agent inserted three code layers, one at a time. After each insertion, the loss dropped further below ordinary training, and the gap persisted to the end of training.
How did the agent decide what to insert, and where? We describe its process below.
From mechanistic interpretability to mechanistic intervention
Mechanistic interpretability is a young field that tries to reverse-engineer a trained and frozen neural network. It asks which internal components carry which computations, and how they combine to produce behaviour. Its tools, such as probes [Alain and Bengio, 2016], ablations and attribution graphs [Ameisen et al., 2025], are often used to audit models for safety [Marks et al., 2025], steer behaviour and understand what their neurons learned.
In our case, the agent turns the question around: what is the network failing to learn, why, and what can be added to it while it still trains? To answer it, the same tools apply, and more . This task gives interpretability something it often lacks: an objective test. An intervention either improves the metric or it doesn't.
To answer this question, the agents have begun building a toolkit of their own, through trial and error on networks as they train. The tool the agent reaches for most often is a quantity computed on the live network, which we dub the phantom gradient . 2
While ordinary gradients are taken with respect to the weights, the phantom gradient is taken with respect to the transformer’s internal representations, at every token and every layer, so it gives a far more detailed map of the change the network wants , which ordinary gradients average away.
We call it the phantom gradient because it is the gradient of weights that are not there. Imagine giving every token, in every document of the batch, its own zero-initialised weight at each layer, added to its representation. The phantom gradient is exactly the gradient those weights would receive.
More specifically, the residual stream at layer $\ell$ is:
$h_\ell$ holds everything the network has computed up to layer $\ell$, a tensor of shape [B, T, D] , with B documents in the batch, T token positions in each, and an embedding of width D at every position.
The phantom gradient at layer $\ell$ is $\partial L/\partial h_\ell$, the gradient of the loss with respect to this representation. It has the same shape, [B, T, D] , and at every position it points in the direction that would lower the loss there most.
Watching this demand map answers the agents’ first two questions: when to intervene, and where in the network to do so. Phantom gradients that persist in a layer across checkpoints mark demand that SGD does not satisfy, and trigger a code-layer insertion there.
Interestingly, when used raw, the phantom gradient tensor misleads. The agents discovered that one normalisation is required first, centring: for each layer, subtract from every phantom vector the mean over all documents whose correct next token is the same. The raw vectors are dominated by which token the answer is, and the subtraction leaves what the context asked for.
Using phantom gradients to find and score code
There are many ways to use phantom gradients. Below is one recipe, in the form the agents converged on during the CLRS-Text runs. The agents begin by grading the network on its own training data, every task and input size separately, at successive checkpoints. Cells whose errors stay high and flat are the targets. Cells that are still improving are left to gradient descent.
From the target cells, the agents collect the tokens the network gets wrong. For each such token, they backpropagate only that token's loss and read the phantom gradient at its position, at every layer. Backpropagating one loss at a time keeps other tokens' errors out of the measurement.
They then cluster the wrong tokens with spherical k-means, which groups vectors by direction rather than size. Each token's fingerprint gives equal weight to two directions, its phantom gradient at a layer near the demand peak and the network's hidden state at that token. Tokens in the same cluster therefore ask for a similar change from a similar state. Repeating the clustering on half the tokens gives the same clusters.
Each cluster yields examples. A member is the prefix of its document

[truncated]

## Original Extract

Architecture is the last hand-designed piece. We make it part of training.

Can LLMs redesign themselves as they train? New Meet Maxwell, the first Code Execution Model. Research About Careers News Visit Maxwell Research About Careers News Visit Maxwell Research / Can LLMs redesign themselves as they train? Can LLMs redesign themselves as they train?
Architecture is the last hand-designed piece. We make it part of training.
01 TL;DR 02 Nanochat: where does a language model need help? 03 Evolution is mandatory 04 Taking on jaggedness 05 Closing the loop: DeepSeek improving itself 06 Looking ahead Related research
Publications · Oct 25, 2026 Alignment via Generalization
Publications · Oct 19, 2026 A New Record in Extremal Chess Composition
Publications · Oct 6, 2026 Advancing a decades-old problem in extremal graph theory
We put an AI agent inside a neural network's training run, where it redesigns the network on the fly.
The agent does what we call mechanistic intervention . It reads the network's internal state, much as mechanistic interpretability research does, but instead of explaining a finished network, it changes the network while it is still learning . The agent diagnoses what the network fails to learn and fixes the problem by changing the data , the objective , the update rule or the architecture itself .
This requires an entirely new toolkit , which the agent builds itself . On nanochat , the agent discovered one of its main instruments, which we call the phantom gradient : a measure of unmet demands.
The agent writes small programs and grows them into the network as code layers . The code layers are often “alien”. For example, one such layer computes gradients of hidden representations within a single forward pass and uses those to self-correct.
On CLRS-Text , a dataset of classical algorithms, the agent takes on jaggedness . We find that a transformer trained on this dataset handles the input sizes it was trained on, yet fails on the sizes in between, even when trained for twice as long. Code layers grown during training raise accuracy at those in-between sizes from 33% to 79% .
A model can also do this to itself. We let DeepSeek-V4.1-Flash act as the agent on itself . Using a code layer that reads its own hidden states, it raised its accuracy on AIME 2026 and HMMT 2026 from 64% to 71% .
Today, a transformer still transforms everything but itself. We are on the cusp of models that train and re-architect themselves , with people setting only the goal and the rules. Models will be able to carry themselves beyond the tasks they already know: they'll do science on themselves, testing ideas, experimenting and expanding their capabilities.
However long you train a neural network, only its weights change. Every other component, from the optimiser to the data mix and the architecture, remains static. What if, instead, we made everything learnable, including the architecture?
In this post, we find early signs that coding agents can redesign neural networks while they train . The resulting architectures can look rather alien. Below, we report early results on four experiments, each aimed at a basic limitation of today's transformers.
Recent work has let coding agents run research projects end to end , tune training code overnight , rewrite the update rules of networks as they train [Bai et al., 2026] and even edit their weights directly [Takehi et al., 2026] . We simply ask: why not let them take over every part of training at once, including even the architecture itself?
Transformers have well-known limitations. They often fail to generalise to inputs longer than those seen in training, and scale alone does not fix this [Anil et al., 2022] [Delétang et al., 2023] . They struggle with compositional tasks that take many steps [Dziri et al., 2023] . Even simple logic (like inferring "B is A" from "A is B") is unreliable [Berglund et al., 2023] .
Some of these limitations trace back to gradient descent itself. Gradient descent fits statistics extremely well, but it settles for whatever lowers the loss fastest, often shortcuts that hold only on the training data [Geirhos et al., 2020] . The representations it finds can be fractured and entangled [Kumar et al., 2025] . Code behaves differently. Programs apply the same rule to every input, seen or unseen, so they generalise by construction. Programs are also short. The shorter the description of a hypothesis, the fewer examples it needs to generalise [Blumer et al., 1987] , an idea formalised in the minimum description length principle [Rissanen, 1978] [Shalev-Shwartz and Ben-David, 2014] .
Coding agents can cover for gradient descent. When placed inside the training loop, the agents can find what the network fails to learn, work out why and write exact programs from a handful of failing examples. These programs go into the network itself, as new layers, which we call code layers . Such a change is beyond gradient descent, which moves in small, continuous steps. Rules can be written as code, but the countless statistical patterns hidden in massive data cannot. The agents can still use gradient descent to train everything else, including how much to rely on each code layer.
To decide where code layers are needed and what they should compute, the agents need an entirely new toolkit, which we call mechanistic intervention . Like mechanistic interpretability , these tools read the network's internals, but they do so during training, while the network can still adapt to what the agents change. These tools are still at an early stage. We call on the community to help develop them.
Finding demand (nanochat). We introduce the phantom gradient, the first mechanistic intervention tool discovered by the agent, and show how the agent uses it to find where a language model needs help. We then look at the layers it wrote, including one that lets the network correct itself using its own recent errors.
Evolution is mandatory (SWA). Comparing layers grown during training with the same layers planted at the start shows that timing matters, a new kind of training dynamic.
Jaggedness (CLRS-Text). We train a transformer on CLRS-Text, a benchmark of classical algorithms. Trained normally, the model solves problems only at the sizes it saw in training, failing at the sizes in between. Code layers grown during training close most of these gaps.
Closing the loop (DeepSeek). DeepSeek-V4.1-Flash acts as the agent on itself. Using a code layer it designed, which reads its own hidden states, it raises its accuracy on AIME 2026 and HMMT 2026 , without thinking, from 64% to 71%.
To avoid overfitting, we keep the test set hidden from the agent in every experiment. Beyond that, the agent has free rein to come up with ideas, design its own experiments and mechanistic intervention tools, and insert the code layers. We describe each experiment in detail below.
Nanochat: where does a language model need help?
We start by training nanochat [Karpathy, 2025], a small ChatGPT-style language model. We train it as usual: next-token prediction, with plain gradient descent. The network is a twelve-layer transformer of about 286 million parameters. 1
Can a coding agent, brought into the training loop, help the network learn more from the same data? In other words, can we improve sample efficiency?
We train the same network again, this time with a coding agent inside the training loop. The setup is identical (same data, schedule, steps and gradient-based updates) except that, at checkpoints of its choosing, the agent is allowed to insert new layers written as code, editing the architecture itself . Once a layer is inserted, ordinary training resumes from that point until the next insertion.
Over the span of this experiment, the agent inserted three code layers, one at a time. After each insertion, the loss dropped further below ordinary training, and the gap persisted to the end of training.
How did the agent decide what to insert, and where? We describe its process below.
From mechanistic interpretability to mechanistic intervention
Mechanistic interpretability is a young field that tries to reverse-engineer a trained and frozen neural network. It asks which internal components carry which computations, and how they combine to produce behaviour. Its tools, such as probes [Alain and Bengio, 2016], ablations and attribution graphs [Ameisen et al., 2025], are often used to audit models for safety [Marks et al., 2025], steer behaviour and understand what their neurons learned.
In our case, the agent turns the question around: what is the network failing to learn, why, and what can be added to it while it still trains? To answer it, the same tools apply, and more . This task gives interpretability something it often lacks: an objective test. An intervention either improves the metric or it doesn't.
To answer this question, the agents have begun building a toolkit of their own, through trial and error on networks as they train. The tool the agent reaches for most often is a quantity computed on the live network, which we dub the phantom gradient . 2
While ordinary gradients are taken with respect to the weights, the phantom gradient is taken with respect to the transformer’s internal representations, at every token and every layer, so it gives a far more detailed map of the change the network wants , which ordinary gradients average away.
We call it the phantom gradient because it is the gradient of weights that are not there. Imagine giving every token, in every document of the batch, its own zero-initialised weight at each layer, added to its representation. The phantom gradient is exactly the gradient those weights would receive.
More specifically, the residual stream at layer $\ell$ is:
$h_\ell$ holds everything the network has computed up to layer $\ell$, a tensor of shape [B, T, D] , with B documents in the batch, T token positions in each, and an embedding of width D at every position.
The phantom gradient at layer $\ell$ is $\partial L/\partial h_\ell$, the gradient of the loss with respect to this representation. It has the same shape, [B, T, D] , and at every position it points in the direction that would lower the loss there most.
Watching this demand map answers the agents’ first two questions: when to intervene, and where in the network to do so. Phantom gradients that persist in a layer across checkpoints mark demand that SGD does not satisfy, and trigger a code-layer insertion there.
Interestingly, when used raw, the phantom gradient tensor misleads. The agents discovered that one normalisation is required first, centring: for each layer, subtract from every phantom vector the mean over all documents whose correct next token is the same. The raw vectors are dominated by which token the answer is, and the subtraction leaves what the context asked for.
Using phantom gradients to find and score code
There are many ways to use phantom gradients. Below is one recipe, in the form the agents converged on during the CLRS-Text runs. The agents begin by grading the network on its own training data, every task and input size separately, at successive checkpoints. Cells whose errors stay high and flat are the targets. Cells that are still improving are left to gradient descent.
From the target cells, the agents collect the tokens the network gets wrong. For each such token, they backpropagate only that token's loss and read the phantom gradient at its position, at every layer. Backpropagating one loss at a time keeps other tokens' errors out of the measurement.
They then cluster the wrong tokens with spherical k-means, which groups vectors by direction rather than size. Each token's fingerprint gives equal weight to two directions, its phantom gradient at a layer near the demand peak and the network's hidden state at that token. Tokens in the same cluster therefore ask for a similar change from a similar state. Repeating the clustering on half the tokens gives the same clusters.
Each cluster yields examples. A member is the prefix of its document

[truncated]
