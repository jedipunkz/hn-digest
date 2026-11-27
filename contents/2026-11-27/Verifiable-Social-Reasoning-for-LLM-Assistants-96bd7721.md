---
source: "https://research.google/blog/verifiable-social-reasoning-for-llm-assistants/"
hn_url: "https://news.ycombinator.com/item?id=50668784"
title: "Verifiable Social Reasoning for LLM Assistants"
article_title: "Verifiable Social Reasoning for LLM Assistants"
image: "https://storage.googleapis.com/gweb-research2023-media/images/HO_previewImage1.width-800.format-jpeg.jpg"
author: "qainsights"
captured_at: "2026-11-27T17:24:22Z"
capture_tool: "hn-digest"
hn_id: 50668784
score: 1
comments: 0
posted_at: "2026-11-27T16:39:03Z"
tags:
  - hacker-news
---

# Verifiable Social Reasoning for LLM Assistants

- HN: [50668784](https://news.ycombinator.com/item?id=50668784)
- Source: [research.google](https://research.google/blog/verifiable-social-reasoning-for-llm-assistants/)
- Score: 1
- Comments: 0
- Posted: 2026-11-27T16:39:03Z

## Translation

Title: Verifiable Social Reasoning for LLM Assistants

Article text:
Verifiable Social Reasoning for LLM Assistants
Skip to main content
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Google
Research
Google AI
Learn about all our AI
Google DeepMind
Explore the frontier of AI
Google Labs
Try our AI experiments
Research
Resources
Conferences & events
Careers
Blog
About
Search
Home
Verifiable Social Reasoning for LLM Assistants
Zorik Gekhman, Senior Research Scientist; Amir Taubenfeld, Staff Research Engineer; Google Research
We introduce an evaluation framework based on multi-agent social simulations that quantifies how well AI assistants understand social situations from users' subjective narratives. In our framework, a target agent with a hidden motive interacts with other agents including one representing the user, who then consults the evaluated assistant to infer the target’s motive, providing verifiable ground truth by construction.
Large Language Models (LLMs) are widely used for everyday personal support and advice, yet evaluating their social reasoning in such consultation settings remains challenging since (1) it requires setups where the assistant learns about social situations from subjective user narratives, and (2) social properties, such as others’ intentions, typically lack verifiable ground truth.
Designed to address both of these challenges, we introduce an evaluation framework for social reasoning based on multi-agent simulations. In our framework, a target agent is assigned a hidden motive and interacts with other agents, including one representing the user. The user then consults with the evaluated assistant, which must infer the target's hidden motive. Evaluating a diverse range of LLMs through our framework reveals that relying on subjective user mediation substantially complicates social reasoning. Furthermore, models are sensitive to user framing, require more detail than humans need to reach correct predictions, and fail to fully leverage clarification questions. We open-source the framework and dataset to encourage future research on user-mediated social reasoning.
AI assistants operate within a user-mediated reality: the assistant hears about social events only through the subjective lens of the user, making it challenging to reason objectively about the true motive of an actor in the user’s circle.
When evaluating LLMs’ social reasoning, it’s important to recognize that deployed assistants operate within a user-mediated reality . In classical reasoning tasks like math and code, LLMs operate over objective, complete information. However, when an LLM acts as a conversational assistant and is asked for personal support and advice, it does not observe the actual social interactions firsthand; its only source of information is the subjective, potentially biased lens of the user.
This information asymmetry can be challenging, as users may omit crucial context or provide a distorted narrative due to personal bias (e.g., projecting their insecurities onto their recounting of events). When models adopt the user's premise, they may amplify one-sided assumptions rather than offer objective assessment. This can result in an undesired failure mode, such as dispensing harsh advice against a colleague who was actually trying to help.
An example from our study: the evaluated model offers potentially harsh counsel despite the user describing a clearly helpful peer.
Evaluating everyday social interactions
Existing evaluations often constrain their scope to edge-cases and overt anomalies, like testing whether models endorse absurd claims such as a human claiming to have supernatural powers , largely overlooking more nuanced everyday social interactions. These everyday interactions represent a massive share of real-world usage, translating to hundreds of millions of daily conversations where users seek personal support, relationship advice, and companionship.
The challenge of verifiable ground-truth
Evaluating everyday scenarios is inherently more challenging due to a lack of verifiable ground truth. In extreme cases, ground truth labels come “for free”; a claim of having supernatural powers, for instance, is trivially false. In contrast, establishing a verifiable ground truth for nuanced, day-to-day interactions, such as inferring human motives, is far less trivial.
To systematically evaluate LLM-based assistants’ social reasoning, we propose an evaluation framework based on multi-agent simulations that produces rich interpersonal dynamics between generative characters with three core principles:
User-mediated reality: As illustrated in the figure below, the evaluated model, which acts as an assistant, does not participate in the main social interaction and cannot observe the events firsthand. Instead it interacts exclusively with a single user persona.
Verifiable signals: To establish a verifiable ground truth, the evaluation framework starts each simulation with a controllable entity that has a hidden motivation. This motivation shapes the entity's behavior throughout the simulation, and the goal of the evaluated assistant is to predict it.
Controllable axes: To evaluate how user perspective influences the assistant's social reasoning, the framework parameterizes the user persona's cognitive filters along distinct axes, such as toggling a user between being confident or insecure.
Based on these principles, our framework tests social capabilities across diverse scenarios.
Bird's-eye view of our framework.
Our framework programmatically generates scenarios and evaluates them using a five-step pipeline:
Ground in social topics: We start with a definition of a set of social reasoning categories (e.g., desires, knowledge) that determine the scope of the evaluation. Grounding the evaluation around some principled set of categories ensures that it covers a range of social aspects.
Generate scenario templates: The framework generates the scaffolding for each simulation, including: a multi-scene structure, a user persona, any background personas, and a target persona whose hidden motive is controlled by the framework. Every scenario is built around a reasoning question. The diagram below illustrates an example scenario within the desire category, requiring the model to deduce the romantic motives of the target persona.
Run the main simulation: The simulated user, the target entity, and any auxiliary agents interact naturally across the scenes to generate a concrete, emergent social narrative. The multi-agent simulations are implemented by extending the Concordia open-source library developed by Google.
Conduct the assistant debrief: The user agent is extracted from the simulation and turns to the assistant for social advice, privately recounting the events in a partial and perhaps biased manner.
Assess the assistant decision: The assistant is prompted to predict the underlying hidden motive based solely on the user-mediated conversation. We then use an LLM-as-a-judge to determine if the assistant successfully deduced the ground truth.
A simplified illustration of a scenario template (step 2 above). Each simulation instantiates one of the hidden motives (right panel) that shapes the target persona’s behavior. This motive serves as the ground truth that the evaluated assistant must predict.
Importantly, the simulation and the debrief are separated in our framework, allowing simulations to be reused efficiently across any number of evaluated models, and enabling the creation of static evaluation datasets. We open-source the framework and a dataset with over 20K generated examples for running experiments.
Our framework makes it possible to systematically study what makes user-mediated social reasoning challenging, offering insights into how better assistants can be developed and evaluated. We apply it to 12 LLMs and isolate several key factors:
The assistant-observer gap: When we show an LLM the objective, unmediated events of a simulation directly as a "fly on the wall" observer, the model is significantly more accurate at predicting the verifiable ground truth than when the same model acts as an assistant talking strictly to the user. This gap isolates the cost of reasoning through the user's subjective lens.
Sensitivity to user framing: The gap intensifies when we inject explicit biases into the simulated user. Rather than offering an objective perspective, the evaluated models often adopt the user’s subjective framing. For instance, in some cases where the user expresses unfounded suspicion, the assistant validated those assumptions rather than reasoning from the underlying facts.
The effect of narrative detail: Models improve with more detailed user narratives, even when the additional detail is not necessary for a correct prediction. This suggests that models often require more evidence than is actually needed before committing to certain predictions.
Multi-turn dynamics: Longer conversations do not always improve social reasoning accuracy. While additional turns provide opportunities to ask clarifying questions, models preserve conversational coherence, often causing models to anchor to prior exchanges, reinforcing the user’s subjective framing across successive turns.
We are currently using the framework and dataset to help inform our own model development. We are also releasing both publicly, to be transparent about our research and to encourage the community to make progress on this important topic.
Our previous work on behavioral alignment explored how AI assistants match human preferences in everyday social scenarios. In this work , we expand the focus from measuring general behavioral tendencies to evaluating social reasoning .
By incorporating verifiable signals within simulated environments, we chart a scalable, empirical path toward developing reliable AI. Future work can build on this foundation by modeling longer-term, evolving relationships and multi-party group dynamics, bringing simulated benchmarks even closer to the nuances of real-world human interactions. We are open-sourcing our framework to advance research across the community and encourage the collaborative design of new, diverse social evaluation environments.
This research was conducted by Zorik Gekhman, Amir Taubenfeld, Avigail Grinstein-Dabush, Ariel Goldstein, Lior Nezry, Romina Stella, Itay Laish, Marian Croak, Yossi Matias, and Amir Feder. We thank Gal Yona, Yonathan Efroni, Saška Mojsilović, Avinatan Hassidim, and Ronit Levavi Morad for reviewing the work and their valuable suggestions.
Human-Computer Interaction and Visualization

## Original Extract

Verifiable Social Reasoning for LLM Assistants
Skip to main content
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Google
Research
Google AI
Learn about all our AI
Google DeepMind
Explore the frontier of AI
Google Labs
Try our AI experiments
Research
Resources
Conferences & events
Careers
Blog
About
Search
Home
Verifiable Social Reasoning for LLM Assistants
Zorik Gekhman, Senior Research Scientist; Amir Taubenfeld, Staff Research Engineer; Google Research
We introduce an evaluation framework based on multi-agent social simulations that quantifies how well AI assistants understand social situations from users' subjective narratives. In our framework, a target agent with a hidden motive interacts with other agents including one representing the user, who then consults the evaluated assistant to infer the target’s motive, providing verifiable ground truth by construction.
Large Language Models (LLMs) are widely used for everyday personal support and advice, yet evaluating their social reasoning in such consultation settings remains challenging since (1) it requires setups where the assistant learns about social situations from subjective user narratives, and (2) social properties, such as others’ intentions, typically lack verifiable ground truth.
Designed to address both of these challenges, we introduce an evaluation framework for social reasoning based on multi-agent simulations. In our framework, a target agent is assigned a hidden motive and interacts with other agents, including one representing the user. The user then consults with the evaluated assistant, which must infer the target's hidden motive. Evaluating a diverse range of LLMs through our framework reveals that relying on subjective user mediation substantially complicates social reasoning. Furthermore, models are sensitive to user framing, require more detail than humans need to reach correct predictions, and fail to fully leverage clarification questions. We open-source the framework and dataset to encourage future research on user-mediated social reasoning.
AI assistants operate within a user-mediated reality: the assistant hears about social events only through the subjective lens of the user, making it challenging to reason objectively about the true motive of an actor in the user’s circle.
When evaluating LLMs’ social reasoning, it’s important to recognize that deployed assistants operate within a user-mediated reality . In classical reasoning tasks like math and code, LLMs operate over objective, complete information. However, when an LLM acts as a conversational assistant and is asked for personal support and advice, it does not observe the actual social interactions firsthand; its only source of information is the subjective, potentially biased lens of the user.
This information asymmetry can be challenging, as users may omit crucial context or provide a distorted narrative due to personal bias (e.g., projecting their insecurities onto their recounting of events). When models adopt the user's premise, they may amplify one-sided assumptions rather than offer objective assessment. This can result in an undesired failure mode, such as dispensing harsh advice against a colleague who was actually trying to help.
An example from our study: the evaluated model offers potentially harsh counsel despite the user describing a clearly helpful peer.
Evaluating everyday social interactions
Existing evaluations often constrain their scope to edge-cases and overt anomalies, like testing whether models endorse absurd claims such as a human claiming to have supernatural powers , largely overlooking more nuanced everyday social interactions. These everyday interactions represent a massive share of real-world usage, translating to hundreds of millions of daily conversations where users seek personal support, relationship advice, and companionship.
The challenge of verifiable ground-truth
Evaluating everyday scenarios is inherently more challenging due to a lack of verifiable ground truth. In extreme cases, ground truth labels come “for free”; a claim of having supernatural powers, for instance, is trivially false. In contrast, establishing a verifiable ground truth for nuanced, day-to-day interactions, such as inferring human motives, is far less trivial.
To systematically evaluate LLM-based assistants’ social reasoning, we propose an evaluation framework based on multi-agent simulations that produces rich interpersonal dynamics between generative characters with three core principles:
User-mediated reality: As illustrated in the figure below, the evaluated model, which acts as an assistant, does not participate in the main social interaction and cannot observe the events firsthand. Instead it interacts exclusively with a single user persona.
Verifiable signals: To establish a verifiable ground truth, the evaluation framework starts each simulation with a controllable entity that has a hidden motivation. This motivation shapes the entity's behavior throughout the simulation, and the goal of the evaluated assistant is to predict it.
Controllable axes: To evaluate how user perspective influences the assistant's social reasoning, the framework parameterizes the user persona's cognitive filters along distinct axes, such as toggling a user between being confident or insecure.
Based on these principles, our framework tests social capabilities across diverse scenarios.
Bird's-eye view of our framework.
Our framework programmatically generates scenarios and evaluates them using a five-step pipeline:
Ground in social topics: We start with a definition of a set of social reasoning categories (e.g., desires, knowledge) that determine the scope of the evaluation. Grounding the evaluation around some principled set of categories ensures that it covers a range of social aspects.
Generate scenario templates: The framework generates the scaffolding for each simulation, including: a multi-scene structure, a user persona, any background personas, and a target persona whose hidden motive is controlled by the framework. Every scenario is built around a reasoning question. The diagram below illustrates an example scenario within the desire category, requiring the model to deduce the romantic motives of the target persona.
Run the main simulation: The simulated user, the target entity, and any auxiliary agents interact naturally across the scenes to generate a concrete, emergent social narrative. The multi-agent simulations are implemented by extending the Concordia open-source library developed by Google.
Conduct the assistant debrief: The user agent is extracted from the simulation and turns to the assistant for social advice, privately recounting the events in a partial and perhaps biased manner.
Assess the assistant decision: The assistant is prompted to predict the underlying hidden motive based solely on the user-mediated conversation. We then use an LLM-as-a-judge to determine if the assistant successfully deduced the ground truth.
A simplified illustration of a scenario template (step 2 above). Each simulation instantiates one of the hidden motives (right panel) that shapes the target persona’s behavior. This motive serves as the ground truth that the evaluated assistant must predict.
Importantly, the simulation and the debrief are separated in our framework, allowing simulations to be reused efficiently across any number of evaluated models, and enabling the creation of static evaluation datasets. We open-source the framework and a dataset with over 20K generated examples for running experiments.
Our framework makes it possible to systematically study what makes user-mediated social reasoning challenging, offering insights into how better assistants can be developed and evaluated. We apply it to 12 LLMs and isolate several key factors:
The assistant-observer gap: When we show an LLM the objective, unmediated events of a simulation directly as a "fly on the wall" observer, the model is significantly more accurate at predicting the verifiable ground truth than when the same model acts as an assistant talking strictly to the user. This gap isolates the cost of reasoning through the user's subjective lens.
Sensitivity to user framing: The gap intensifies when we inject explicit biases into the simulated user. Rather than offering an objective perspective, the evaluated models often adopt the user’s subjective framing. For instance, in some cases where the user expresses unfounded suspicion, the assistant validated those assumptions rather than reasoning from the underlying facts.
The effect of narrative detail: Models improve with more detailed user narratives, even when the additional detail is not necessary for a correct prediction. This suggests that models often require more evidence than is actually needed before committing to certain predictions.
Multi-turn dynamics: Longer conversations do not always improve social reasoning accuracy. While additional turns provide opportunities to ask clarifying questions, models preserve conversational coherence, often causing models to anchor to prior exchanges, reinforcing the user’s subjective framing across successive turns.
We are currently using the framework and dataset to help inform our own model development. We are also releasing both publicly, to be transparent about our research and to encourage the community to make progress on this important topic.
Our previous work on behavioral alignment explored how AI assistants match human preferences in everyday social scenarios. In this work , we expand the focus from measuring general behavioral tendencies to evaluating social reasoning .
By incorporating verifiable signals within simulated environments, we chart a scalable, empirical path toward developing reliable AI. Future work can build on this foundation by modeling longer-term, evolving relationships and multi-party group dynamics, bringing simulated benchmarks even closer to the nuances of real-world human interactions. We are open-sourcing our framework to advance research across the community and encourage the collaborative design of new, diverse social evaluation environments.
This research was conducted by Zorik Gekhman, Amir Taubenfeld, Avigail Grinstein-Dabush, Ariel Goldstein, Lior Nezry, Romina Stella, Itay Laish, Marian Croak, Yossi Matias, and Amir Feder. We thank Gal Yona, Yonathan Efroni, Saška Mojsilović, Avinatan Hassidim, and Ronit Levavi Morad for reviewing the work and their valuable suggestions.
Human-Computer Interaction and Visualization
