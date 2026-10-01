---
source: "https://www.fig.inc/blog/open-ended-data-is-a-bottleneck/"
hn_url: "https://news.ycombinator.com/item?id=49921816"
title: "Open-ended data is a bottleneck for AI systems tackling real knowledge work"
article_title: "Open-Ended Data is a Bottleneck for AI Tackling Real Work"
image: "https://storage.ghost.io/c/a0/84/a0848f29-ea39-4b8c-a26d-c283a5a64f0c/content/images/size/w1200/2026/09/Part_1_Splash_16x9-1.png"
author: "joehabeebs"
captured_at: "2026-10-01T14:15:45Z"
capture_tool: "hn-digest"
hn_id: 49921816
score: 7
comments: 1
posted_at: "2026-10-01T14:00:03Z"
tags:
  - hacker-news
---

# Open-ended data is a bottleneck for AI systems tackling real knowledge work

- HN: [49921816](https://news.ycombinator.com/item?id=49921816)
- Source: [www.fig.inc](https://www.fig.inc/blog/open-ended-data-is-a-bottleneck/)
- Score: 7
- Comments: 1
- Posted: 2026-10-01T14:00:03Z

## Translation

Title: Open-ended data is a bottleneck for AI systems tackling real knowledge work
Article title: Open-Ended Data is a Bottleneck for AI Tackling Real Work
Description: We believe there’s a requirement for fundamentally different data from the scripted workflows that agents are trained on today.

Article text:
Open-Ended Data is a Bottleneck for AI systems Tackling Real Knowledge Work
Real knowledge work is not a single task
Look at the screen in front of you. You’ll likely see a combination of notifications, message drafts, documents, and countless browser tabs across multiple windows and monitors. That desktop is uniquely yours and is the direct result of every action you’ve taken before this one.
We use our computers to progress towards goals, seamlessly navigating through incredibly complex environments. We shift between contexts, answer messages, ignore notifications, pause and resume tasks, and juggle multiple threads, all in service of completing tasks and making progress towards those goals. Naturally, as the capabilities of AI increase, we expect models to reliably work alongside us in the same spaces and towards the same goals.
Despite falling short of that today, models are quickly progressing in their ability to tackle self-contained long-horizon tasks (Opus 5.5 and GPT-6 Astra claim 72.6% and 81.8% on OSWorld 2.0 respectively). The structure of these tasks and the environments they’re completed in, however, remain far simpler than the reality that agents are expected to be deployed into. They begin with a clearly defined goal, a clear path to follow, a discrete outcome to work towards, and clear criteria that define when the task is ‘done’.
0:00
/ 0:27
1×
A real desktop session with interwoven tasks paused & resumed
Human work, on the other hand, is rarely defined that neatly. When returning to a task, the environment is inherently different than when we left it: messages have arrived, decisions have been made, and the task itself may depend on new information. We constantly carry that context forward and use it to triage what to resume, revise, or change course.
In order for agents to materially progress towards autonomous action across complex, open-ended environments, the data they learn from needs to be built on that foundation.
Computer-use agents are trained on trajectories: detailed recordings of a human or agent operating a computer broken out into individual steps or actions. Trajectories typically begin with the instructions for an objective or task and capture a combination of what was visible on screen, the inputs, and the resulting changes, with a focus on the app(s) used to complete the task.
During training, models learn from these signals. Each step imparts information about what action to take given the current state and the objective.
Beyond volume and length, the signals a trajectory carries scale across two dimensions: what can be observed and the realism of the actions and environments themselves:
Observability: How much of the ground truth, the inputs, and the resulting state change is accurately labeled in the recording. Images or video paired with keyboard and mouse input, accessibility trees, DOM snapshots, and app-switch provenance help narrow this gap in modern data sets.
Distribution: The range of events or scenarios that collection covers. Different applications, GUIs, solution paths, start states, and failures. A wide distribution of realism determines whether a model has been shown anything resembling the state it gets deployed into.
Open-ended work carries important signals in the surrounding context: which goals actions serve, what other work is in progress, and what is on hold or interrupted. For those cues to become useful learning signals in a trajectory, the data must preserve them.
Advancements in observability have led to synchronized accessibility and input telemetry becoming standard. It’s the easier axis to advance because it is an engineering problem: better tooling that is capable of measuring more at improved levels accuracy.
Distribution, on the other hand, has progressed primarily within the boundaries of a task. It’s taken the form of more apps, different solution paths, and varied tool use alongside longer and more complex objectives. The exercises themselves, however, remain structured, single-objective tasks performed on sterile desktops.
The complexity in what we need to capture doesn’t emerge from adding more steps to a predefined objective. It’s created from defining the threaded dependencies between tasks, new information, and the decisions made we move between them.
Have custom data requirements?
We design and generate data for frontier teams building agents, models, benchmarks, evals, and more.
To reach those capabilities, computer-use data must evolve. The scripted tasks that make up the majority of training data today are missing learning signals on how to navigate through the noise of real environments and progress towards open-ended goals.
Doing so means capturing real work as it unfolds: unscripted multi-thousand-step workflows with the nuance of the interwoven tasks within them and the decisions that happen before, after and between actions.
Open-ended work doesn’t have a set of instructions or tasks, nor does it include a rubric that define if it’s sufficiently complete. We draw from our learned experience to structure targets into the deliverables that get us there, the sequence they need to be completed in, and the point they’re considered ‘done’.
Reaching that frontier also requires reliable operation in a shared environment. Multiple windows, message notifications, meetings, and urgent requests will continue to take place. Data needs to demonstrate how we seamlessly sort through what’s relevant, carry relevant context forward, and pick up where we left off.
The hill to climb is in the decisions and complex environments that make open-ended work difficult: identifying what needs to be done, deciding what matters now, adapting when another goal intervenes, and determining when the outcome is sufficient.
To progress against these capabilities, we believe there’s a requirement for fundamentally different data from the scripted workflows that are trained on today: real work, observed continuously in evolving environments and structured into training signals for frontier models.
Fig is building the infrastructure to solve the data requirements for the next generation of AI systems that perceive environments, act in them reliably, and improve from the experience. Getting there requires data grounded in the patterns in which people actually work and in the complexity of the spaces in which that work takes place: real open-ended work across real desktops.
This is a gap we’re actively working to close.
If you’re building models, agents, evals, or benchmarks, we’d love to collaborate. Please reach out , tell us what you’re working on, and we’ll be in touch.
@article{fig_open_ended_data_bottleneck_2026,
title = {Open-Ended Data is a Bottleneck for AI systems Tackling Real Knowledge Work},
author = {Joseph Hakim},
year = {2026},
institution = {Fig},
url = {https://www.fig.inc/open-ended-data-is-a-bottleneck/}
}
Articles
Astra, Opus 5.5, and other Frontier Models Demonstrate Jagged Performance Across SoTA Agentic Tasks from Web Browsing to Robotics
Yangyue Wang1,
Harshvardhan Sikka1, 2,
Pranav Guruprasad1,
Sudipta Chowdhury1
1Fig;
2Georgia Institute of Technology.
Five task spaces by description, effort and error
Task failures (%)Fewer stepsSimilar task content
error threshold 0%:
0 have observed error at or above it
VisualWebArena
Bench2Drive-VL
IndEgo
Assembly101
VLABench
Error threshold
drag to turn
Goal Progress Decays with Task Horizon for Frontier VLMs in Interactive 2D Environments
Pranav Guruprasad1, 2,
Sean Rivera2,
Helen Lu2, 4,
Arushi Jain2,
Hangliang Ren2,
Harshvardhan Sikka1, 2, 3
1Fig;
2Manifold Research Group;
3Georgia Institute of Technology;
4Tufts University.
Relevant links
·
Code
·
Website
·
Evaluate your model
·
Cite this
TL;DR
* We evaluated three highly performant vision-language models on 2D mazes: 50 Minigrid
Fixing Failures in Browser-Use Models: Why More Data Isn't Enough
Investor Interest |
Fig Labs |
Twitter |
LinkedIn |
Contact |
Privacy
Based in beautiful San Francisco, California 🌁

## Original Extract

We believe there’s a requirement for fundamentally different data from the scripted workflows that agents are trained on today.

Open-Ended Data is a Bottleneck for AI systems Tackling Real Knowledge Work
Real knowledge work is not a single task
Look at the screen in front of you. You’ll likely see a combination of notifications, message drafts, documents, and countless browser tabs across multiple windows and monitors. That desktop is uniquely yours and is the direct result of every action you’ve taken before this one.
We use our computers to progress towards goals, seamlessly navigating through incredibly complex environments. We shift between contexts, answer messages, ignore notifications, pause and resume tasks, and juggle multiple threads, all in service of completing tasks and making progress towards those goals. Naturally, as the capabilities of AI increase, we expect models to reliably work alongside us in the same spaces and towards the same goals.
Despite falling short of that today, models are quickly progressing in their ability to tackle self-contained long-horizon tasks (Opus 5.5 and GPT-6 Astra claim 72.6% and 81.8% on OSWorld 2.0 respectively). The structure of these tasks and the environments they’re completed in, however, remain far simpler than the reality that agents are expected to be deployed into. They begin with a clearly defined goal, a clear path to follow, a discrete outcome to work towards, and clear criteria that define when the task is ‘done’.
0:00
/ 0:27
1×
A real desktop session with interwoven tasks paused & resumed
Human work, on the other hand, is rarely defined that neatly. When returning to a task, the environment is inherently different than when we left it: messages have arrived, decisions have been made, and the task itself may depend on new information. We constantly carry that context forward and use it to triage what to resume, revise, or change course.
In order for agents to materially progress towards autonomous action across complex, open-ended environments, the data they learn from needs to be built on that foundation.
Computer-use agents are trained on trajectories: detailed recordings of a human or agent operating a computer broken out into individual steps or actions. Trajectories typically begin with the instructions for an objective or task and capture a combination of what was visible on screen, the inputs, and the resulting changes, with a focus on the app(s) used to complete the task.
During training, models learn from these signals. Each step imparts information about what action to take given the current state and the objective.
Beyond volume and length, the signals a trajectory carries scale across two dimensions: what can be observed and the realism of the actions and environments themselves:
Observability: How much of the ground truth, the inputs, and the resulting state change is accurately labeled in the recording. Images or video paired with keyboard and mouse input, accessibility trees, DOM snapshots, and app-switch provenance help narrow this gap in modern data sets.
Distribution: The range of events or scenarios that collection covers. Different applications, GUIs, solution paths, start states, and failures. A wide distribution of realism determines whether a model has been shown anything resembling the state it gets deployed into.
Open-ended work carries important signals in the surrounding context: which goals actions serve, what other work is in progress, and what is on hold or interrupted. For those cues to become useful learning signals in a trajectory, the data must preserve them.
Advancements in observability have led to synchronized accessibility and input telemetry becoming standard. It’s the easier axis to advance because it is an engineering problem: better tooling that is capable of measuring more at improved levels accuracy.
Distribution, on the other hand, has progressed primarily within the boundaries of a task. It’s taken the form of more apps, different solution paths, and varied tool use alongside longer and more complex objectives. The exercises themselves, however, remain structured, single-objective tasks performed on sterile desktops.
The complexity in what we need to capture doesn’t emerge from adding more steps to a predefined objective. It’s created from defining the threaded dependencies between tasks, new information, and the decisions made we move between them.
Have custom data requirements?
We design and generate data for frontier teams building agents, models, benchmarks, evals, and more.
To reach those capabilities, computer-use data must evolve. The scripted tasks that make up the majority of training data today are missing learning signals on how to navigate through the noise of real environments and progress towards open-ended goals.
Doing so means capturing real work as it unfolds: unscripted multi-thousand-step workflows with the nuance of the interwoven tasks within them and the decisions that happen before, after and between actions.
Open-ended work doesn’t have a set of instructions or tasks, nor does it include a rubric that define if it’s sufficiently complete. We draw from our learned experience to structure targets into the deliverables that get us there, the sequence they need to be completed in, and the point they’re considered ‘done’.
Reaching that frontier also requires reliable operation in a shared environment. Multiple windows, message notifications, meetings, and urgent requests will continue to take place. Data needs to demonstrate how we seamlessly sort through what’s relevant, carry relevant context forward, and pick up where we left off.
The hill to climb is in the decisions and complex environments that make open-ended work difficult: identifying what needs to be done, deciding what matters now, adapting when another goal intervenes, and determining when the outcome is sufficient.
To progress against these capabilities, we believe there’s a requirement for fundamentally different data from the scripted workflows that are trained on today: real work, observed continuously in evolving environments and structured into training signals for frontier models.
Fig is building the infrastructure to solve the data requirements for the next generation of AI systems that perceive environments, act in them reliably, and improve from the experience. Getting there requires data grounded in the patterns in which people actually work and in the complexity of the spaces in which that work takes place: real open-ended work across real desktops.
This is a gap we’re actively working to close.
If you’re building models, agents, evals, or benchmarks, we’d love to collaborate. Please reach out , tell us what you’re working on, and we’ll be in touch.
@article{fig_open_ended_data_bottleneck_2026,
title = {Open-Ended Data is a Bottleneck for AI systems Tackling Real Knowledge Work},
author = {Joseph Hakim},
year = {2026},
institution = {Fig},
url = {https://www.fig.inc/open-ended-data-is-a-bottleneck/}
}
Articles
Astra, Opus 5.5, and other Frontier Models Demonstrate Jagged Performance Across SoTA Agentic Tasks from Web Browsing to Robotics
Yangyue Wang1,
Harshvardhan Sikka1, 2,
Pranav Guruprasad1,
Sudipta Chowdhury1
1Fig;
2Georgia Institute of Technology.
Five task spaces by description, effort and error
Task failures (%)Fewer stepsSimilar task content
error threshold 0%:
0 have observed error at or above it
VisualWebArena
Bench2Drive-VL
IndEgo
Assembly101
VLABench
Error threshold
drag to turn
Goal Progress Decays with Task Horizon for Frontier VLMs in Interactive 2D Environments
Pranav Guruprasad1, 2,
Sean Rivera2,
Helen Lu2, 4,
Arushi Jain2,
Hangliang Ren2,
Harshvardhan Sikka1, 2, 3
1Fig;
2Manifold Research Group;
3Georgia Institute of Technology;
4Tufts University.
Relevant links
·
Code
·
Website
·
Evaluate your model
·
Cite this
TL;DR
* We evaluated three highly performant vision-language models on 2D mazes: 50 Minigrid
Fixing Failures in Browser-Use Models: Why More Data Isn't Enough
Investor Interest |
Fig Labs |
Twitter |
LinkedIn |
Contact |
Privacy
Based in beautiful San Francisco, California 🌁
