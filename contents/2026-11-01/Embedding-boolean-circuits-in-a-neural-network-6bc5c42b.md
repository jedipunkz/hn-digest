---
source: "https://edwardbenson.com/2026/10/embedding-boolean-circuits-in-a-neural-network"
hn_url: "https://news.ycombinator.com/item?id=50326545"
title: "Embedding boolean circuits in a neural network"
article_title: "Embedding boolean circuits in a neural network"
image: ""
author: "eob"
captured_at: "2026-11-01T15:49:32Z"
capture_tool: "hn-digest"
hn_id: 50326545
score: 2
comments: 0
posted_at: "2026-11-01T15:27:44Z"
tags:
  - hacker-news
---

# Embedding boolean circuits in a neural network

- HN: [50326545](https://news.ycombinator.com/item?id=50326545)
- Source: [edwardbenson.com](https://edwardbenson.com/2026/10/embedding-boolean-circuits-in-a-neural-network)
- Score: 2
- Comments: 0
- Posted: 2026-11-01T15:27:44Z

## Translation

Title: Embedding boolean circuits in a neural network
Description: Construct AND, OR, and NOT gates with fixed neural weights, then combine them into circuits that read model activations.

Article text:
Embedding boolean circuits in a neural network > Ted Benson / Writing
Embedding boolean circuits in a neural network
Embedded boolean circuit inside a neural network
In which an LLM diagrams this idea more eloquently than I wrote it
Real-valued activity flows through the model.
This post is mostly me thinking out loud: how is it possible to install a boolean circuit sidecar to an otherwise trained neural network?
It strikes me that there are essentially two distinct problems here:
Representing the boolean circuit as a set of neurons
Connecting the boolean circuit to a trained network
The representation part is relatively straightforward. A neuron is a weighted sum of inputs, plus a bias, passed through an activation function.
If we assert that the inputs are all either 0 or 1, and the activation function is simply “is it greater than zero”, then we can turn boolean operators into neurons like so:
AND(A, B) = A + B - 1.5 > 0
OR(A, B) = A + B - 0.5 > 0
NOT(A) = 0.5 - A > 0
There’s no extra story needed here with respect to boolean expressions. These gates are all you need to create any expression, and the neurons can compose exactly as the symbolic operators would compose on paper.
The second problem is patching this circuit to a trained network.
Signals in a trained neural network aren’t necessarily a clean 0 or 1 depending on where you draw them from.
So you need an ingress and egress mapping: from the neural net into the boolean circuit, and then back into the neural net.
Kind of like Neuralink but for LLMs.
For the ingress side, let’s call a Probe a mapping from a trained neuron with real-valued output into a boolean-valued signal.
We’ll notate a probe like this:
probe Trigger: [L3 N234 >1]
This probe would be 0 if neuron L3 N234 was equal to or less than 1, and 1 otherwise.
It can be implemented with a single neuron at the ingress of the boolean circuit, with a threshold set based on knowledge of the targeted neuron and observed meaning of it firing under different values.
For the egress side, let’s call a Stim an additive push that a boolean value switches on or off.
A stim identifies a neuron and a signed steering magnitude.
stim Watermelon: [L11 N15292 -12.5]
Here -12.5 sets the direction and size of the added push.
This size value has to be chosen based on the observed and desired effect in the trained LLM.
For example, if you found a neuron that made the model act like a watermelon when pushed at -12.5, then you might connect this stim as the output of some boolean circuit intended to switch it on at the right time.
An individual MLP neuron has a scalar activation. In a transformer, its output weights define a direction in the vector-valued residual stream. Here, a stim adds the stated signed magnitude along that direction, normalized to length 1.
Within one forward pass, the probe must run before the location it stimulates, as in the L3-to-L11 example above. Reading a later layer and stimulating an earlier one requires carrying the signal into a subsequent processing step.
I’m fairly certain this describes:
A method of translating a boolean expression into a neural circuit
A method of adapting real-valued neurons into boolean-valued inputs for this circuit
A method of adapting boolean-valued outputs of this circuit back into real-valued outputs in the network
No code example attached this time. I’m using this blog post to structure my thinking for something slightly more involved I’m hacking on.
LinkedIn Robot Book Club 🤖 Kaya 🥥 · © 2026 Edward Benson. All rights reserved.

## Original Extract

Construct AND, OR, and NOT gates with fixed neural weights, then combine them into circuits that read model activations.

Embedding boolean circuits in a neural network > Ted Benson / Writing
Embedding boolean circuits in a neural network
Embedded boolean circuit inside a neural network
In which an LLM diagrams this idea more eloquently than I wrote it
Real-valued activity flows through the model.
This post is mostly me thinking out loud: how is it possible to install a boolean circuit sidecar to an otherwise trained neural network?
It strikes me that there are essentially two distinct problems here:
Representing the boolean circuit as a set of neurons
Connecting the boolean circuit to a trained network
The representation part is relatively straightforward. A neuron is a weighted sum of inputs, plus a bias, passed through an activation function.
If we assert that the inputs are all either 0 or 1, and the activation function is simply “is it greater than zero”, then we can turn boolean operators into neurons like so:
AND(A, B) = A + B - 1.5 > 0
OR(A, B) = A + B - 0.5 > 0
NOT(A) = 0.5 - A > 0
There’s no extra story needed here with respect to boolean expressions. These gates are all you need to create any expression, and the neurons can compose exactly as the symbolic operators would compose on paper.
The second problem is patching this circuit to a trained network.
Signals in a trained neural network aren’t necessarily a clean 0 or 1 depending on where you draw them from.
So you need an ingress and egress mapping: from the neural net into the boolean circuit, and then back into the neural net.
Kind of like Neuralink but for LLMs.
For the ingress side, let’s call a Probe a mapping from a trained neuron with real-valued output into a boolean-valued signal.
We’ll notate a probe like this:
probe Trigger: [L3 N234 >1]
This probe would be 0 if neuron L3 N234 was equal to or less than 1, and 1 otherwise.
It can be implemented with a single neuron at the ingress of the boolean circuit, with a threshold set based on knowledge of the targeted neuron and observed meaning of it firing under different values.
For the egress side, let’s call a Stim an additive push that a boolean value switches on or off.
A stim identifies a neuron and a signed steering magnitude.
stim Watermelon: [L11 N15292 -12.5]
Here -12.5 sets the direction and size of the added push.
This size value has to be chosen based on the observed and desired effect in the trained LLM.
For example, if you found a neuron that made the model act like a watermelon when pushed at -12.5, then you might connect this stim as the output of some boolean circuit intended to switch it on at the right time.
An individual MLP neuron has a scalar activation. In a transformer, its output weights define a direction in the vector-valued residual stream. Here, a stim adds the stated signed magnitude along that direction, normalized to length 1.
Within one forward pass, the probe must run before the location it stimulates, as in the L3-to-L11 example above. Reading a later layer and stimulating an earlier one requires carrying the signal into a subsequent processing step.
I’m fairly certain this describes:
A method of translating a boolean expression into a neural circuit
A method of adapting real-valued neurons into boolean-valued inputs for this circuit
A method of adapting boolean-valued outputs of this circuit back into real-valued outputs in the network
No code example attached this time. I’m using this blog post to structure my thinking for something slightly more involved I’m hacking on.
LinkedIn Robot Book Club 🤖 Kaya 🥥 · © 2026 Edward Benson. All rights reserved.
