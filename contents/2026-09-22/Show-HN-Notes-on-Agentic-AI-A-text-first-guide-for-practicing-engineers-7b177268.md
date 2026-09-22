---
source: "https://github.com/mrsachindixit/agentixit"
hn_url: "https://news.ycombinator.com/item?id=49805807"
title: "Show HN: Notes on Agentic AI – A text-first guide for practicing engineers"
article_title: "GitHub - mrsachindixit/agentixit: Notes on Agentic AI - a text-first teaching module for practising engineers, with runnable local-first code samples. · GitHub"
image: "https://opengraph.githubassets.com/ecd474b3aa422e5960381d5c5ff0977700dfd0cc7e0bfbe61c103e459bcc9f65/mrsachindixit/agentixit"
author: "tbaadm"
captured_at: "2026-09-22T18:32:51Z"
capture_tool: "hn-digest"
hn_id: 49805807
score: 1
comments: 0
posted_at: "2026-09-22T18:19:28Z"
tags:
  - hacker-news
---

# Show HN: Notes on Agentic AI – A text-first guide for practicing engineers

- HN: [49805807](https://news.ycombinator.com/item?id=49805807)
- Source: [github.com](https://github.com/mrsachindixit/agentixit)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T18:19:28Z

## Translation

Title: Show HN: Notes on Agentic AI – A text-first guide for practicing engineers
Article title: GitHub - mrsachindixit/agentixit: Notes on Agentic AI - a text-first teaching module for practising engineers, with runnable local-first code samples. · GitHub
Description: Notes on Agentic AI - a text-first teaching module for practising engineers, with runnable local-first code samples. - mrsachindixit/agentixit

Article text:
GitHub - mrsachindixit/agentixit: Notes on Agentic AI - a text-first teaching module for practising engineers, with runnable local-first code samples. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mrsachindixit
/
agentixit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
21 Commits 21 Commits Folders and files
.github .github code code docs docs images images mermaid mermaid .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md docker-compose.yaml docker-compose.yaml requirements.short.txt requirements.short.txt requirements.txt requirements.txt scratchpad_run_summary.txt scratchpad_run_summary.txt View all files Repository files navigation
This is a practical textbook-style guide for building production-ready AI agents.
It pairs concept-first chapters with runnable, module-based code examples.
🤖 AgenticAI — Hands-On Code Companion
Build real AI agents from scratch: tool-calling, RAG, memory, multi-agent systems, and production hardening — all running locally with Ollama.
👉📘 Start here: Coding Tutorial
📚 Notes on Agentic AI — Textbook Companion
The content below is a structured textbook-style guide: concepts, diagrams, design patterns, and production trade-offs — with direct mapping to runnable modules.
The text is CC BY-NC-SA 4.0 ; the code is MIT . See License at the end.
About
🤖 AgenticAI — Hands-On Code Companion
📚 Notes on Agentic AI — Textbook Companion
Basics of LLMs, Prompts and Tool Calls
What is an Agent?
Language Model (LLM)
Prompt Engineering
Tool Calling (Function Calling)
Foundation of Agentic AI
Retrieval-Augmented Generation (RAG)
Agentic Execution and Patterns
An Agent
DSPy, Embabel, and LlamaIndex
Why these three matter
How students should compare them
Production Aspects
Post-training (out of scope)
Hallucinations and Factfulness
Agent Benchmarks and Evaluation
Cost and Resources Utilization
Enterprise Suites and Protocols
MCP (Model Context Protocol)
A2A and Other Enterprise Concerns
I started teaching engineers about agentic AI. Over the period I realized that teaching practicing programmers about any new Framework has to be mindful of their prior learning. At times the wisdom gained so far feeds into the new framework and technology paradigm. At times the prior learnings hamper their ability to see nuances and retain curiosity about the new kid.
So my teaching style evolved into using lots of code first to drive in the new point of view. And then supplement that with details on the new tech and also commentary on where it has continuation, breakup, evolution and new beginnings. Over period of time my notes and call transcripts were long enough to motivate me into writing this. This text is very opinionated hence, it also assumes the student has prior understanding of basic building blocks and will dig out more when clues are provided. All the text is written by me, citations have been given where due.
When it came to coding examples I generated them using Copilot. However I realized they were too cryptic. So I took help from Sumit Toshniwal. He is one of the AI Engineers working at Actimize on Agentic projects. First he helped me get a feel of younger generation of developers :). Then he helped me simplify the code examples according to the intended learning outcome. We have kept the code samples minimal and self contained because I am expecting experienced developers will fill in the necessary design blanks. We thank NICE Actimize for being the org where we could do all this as one of the part of the work.
This repo is dedicated to Prof. Andrew Ng who is sharing AI related knowledge freely nurturing AI minds across the Globe.
Running the code alongside the text
Each module is a numbered sequence of small, self-contained scripts; the full index is in
code/README.md .
Basics of LLMs, Prompts and Tool Calls
This section introduces the concepts related to Agentic AI. We have grouped the concepts in accordance with the code samples in modules, so that learners can quickly test out the concepts.
Classically, an agent can be defined as a component that has some perception of its environment and has the ability to perform an action/task. A lot of your home automation devices fit into this definition. This also implies that AI is not a prerequisite for an agent. However, the wave of Agentic AI defines an Agent as a component that has the ability to perform some goal-oriented action with some sort of reasoning and planning capability (with or without direct perception of its environment). It goes without saying that LLMs have significantly enabled agents with the capability to reason and plan. This adds dynamism and adaptability to agents given that LLMs have a generic ability to reason about practically everything.
The hard part of achieving the action, however, is still accomplished by tools and retrieval systems that assist the agent. This underlines that agents are not first-class citizens of LLMs; LLMs can be made aware that agentic processing is happening to some degree.
A lot of agentic AI's potential comes from the application of traditional design patterns facilitated by tools like LangChain or LangGraph. These frameworks allow tighter control over processing flow by introducing concepts like task graphs, memory, and tools.
The final effect is a system that can plan and execute complex tasks with a good degree of certainty while offering adaptability for newer scenarios. We will revisit most of these concepts in the sections that follow.
Figure: each paradigm shift traded control for capability — agents buy adaptability at the price of determinism, which is why the engineering disciplines later in this book exist
graph TB
subgraph AGENT["🤖 AGENT"]
LLM["🧠 Language Model\n(Reasoning Engine)"]
Tools["🔧 Tools\n(Actions & APIs)"]
Memory["💾 Memory\n(State & History)"]
Prompt["📝 Prompts\n(Instructions)"]
end
ENV["🌍 Environment\n(Data, Users, Systems)"] -->|"Perception"| AGENT
AGENT -->|"Actions"| ENV
Prompt -->|"Instructs"| LLM
LLM -->|"Selects"| Tools
LLM -->|"Reads/Writes"| Memory
Tools -->|"Results"| LLM
Memory -->|"Context"| LLM
style LLM fill:#EF4444,stroke:#B91C1C,color:#fff,stroke-width:3px
style Tools fill:#10B981,stroke:#047857,color:#fff,stroke-width:3px
style Memory fill:#8B5CF6,stroke:#6D28D9,color:#fff,stroke-width:3px
style Prompt fill:#F59E0B,stroke:#B45309,color:#fff,stroke-width:3px
style ENV fill:#6366F1,stroke:#4338CA,color:#fff,stroke-width:3px
style AGENT fill:#F1F5F9,stroke:#3B82F6,stroke-width:3px
linkStyle default stroke:#475569,stroke-width:2px
Loading
Figure: anatomy of an agent — LLM, tools, memory, and prompts interacting with the environment
Figure: agent capability ladder — scope to the lowest rung that solves the problem
The LLM is the cognitive core of an agent. It processes natural language input and generates inferences. The instruction that gives data and direction to the LLM is called a prompt. At the same time, LLMs are trained on world data. This inherent information plus the prompt and the internal mathematical techniques that go into creating the LLM produce the effect of interpretation, inference, and reasoning. This effect of interpretation, inference, and reasoning is what makes LLMs popular and enables the whole Agentic phenomenon! A Survey of Reasoning with Foundation Models .
In this course, we use local LLMs via Ollama to avoid cloud dependencies. Most of the code demos show only the name of the LLM/URL when calls are made to it. However, LLMs also give us a few more parameters to tweak that can affect its outcome. Settings like top_p, top_k, and temperature help us control the randomness and sampling of the processing. There are more flags like reasoning effort, streaming, and verbosity that a given LLM might support; developers should read the documentation before moving to production.
For most use cases, these are not things we tweak daily in our code. In some cases, one can adjust the context length and the depth of reasoning. It is recommended to read the documentation of your LLM model to understand these settings. However, given the way LLM models and agentic use cases are evolving, we are good with the defaults for all practical purposes.
graph LR
Input["📝 Input\nPrompt"] --> Tokenize["🔤 Tokenize"]
Tokenize --> Encode["🔢 Encode\nEmbeddings"]
Encode --> Attend["🧠 Multi-Head\nAttention"]
Attend --> Decode["📊 Decode\nProbabilities"]
Decode --> Sample["🎲 Sample\n(temp, top_p)"]
Sample --> Output["✅ Output\nTokens"]
style Input fill:#6366F1,stroke:#4338CA,color:#fff,stroke-width:3px
style Tokenize fill:#3B82F6,stroke:#1E40AF,color:#fff,stroke-width:3px
style Encode fill:#8B5CF6,stroke:#6D28D9,color:#fff,stroke-width:3px
style Attend fill:#EF4444,stroke:#B91C1C,color:#fff,stroke-width:3px
style Decode fill:#F59E0B,stroke:#B45309,color:#fff,stroke-width:3px
style Sample fill:#EC4899,stroke:#BE185D,color:#fff,stroke-width:3px
style Output fill:#10B981,stroke:#047857,color:#fff,stroke-width:3px
linkStyle default stroke:#475569,stroke-width:2px
Loading
Figure: LLM request-response lifecycle — tokenize, encode, attend, decode, sample
Prompts are instructions to the LLM that shape its behavior. Prompt also significantly shapes the outcome in terms of the result as well as the format. In fact it is the prompt that actually causes the effect to take place. So, it is very important to master prompt engineering Paper ) to improve agentic outcomes.
First and foremost, ChatGPT has made many to believe that the chat interactions we have with it are how prompts are/should be. Statements like "English is your new programming language" has also fuelled this nonsense grasp of what prompts need to be.
When it comes to getting the LLM to produce the effect we want, it is the sum total of one or many sentences with examples and keywords that have to come together as a coherent instruction, agentic or not. Here is a good paper you must read Can Large Language Models Reason and Plan?
Suggestions to craft good prompts:
Role: This can include system, user, or assistant. Specifying the system role has an overarching effect. Usually user or assistant roles help us get the effect we want.
Task: Framing the prompt as a task directs the LLM into narrowing down the range of outcomes. However, task is a generic word; it can also be replaced with aim, job, or similar words. The task itself can be to think and reflect, a chain of actions, a review, or mid-conversation instructions for longer tasks.
Constraints: Framing prompts with constraints also narrows down the range of alternatives the LLM has for interpreting and processing the prompt. It g

[truncated]

## Original Extract

Notes on Agentic AI - a text-first teaching module for practising engineers, with runnable local-first code samples. - mrsachindixit/agentixit

GitHub - mrsachindixit/agentixit: Notes on Agentic AI - a text-first teaching module for practising engineers, with runnable local-first code samples. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mrsachindixit
/
agentixit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
21 Commits 21 Commits Folders and files
.github .github code code docs docs images images mermaid mermaid .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md docker-compose.yaml docker-compose.yaml requirements.short.txt requirements.short.txt requirements.txt requirements.txt scratchpad_run_summary.txt scratchpad_run_summary.txt View all files Repository files navigation
This is a practical textbook-style guide for building production-ready AI agents.
It pairs concept-first chapters with runnable, module-based code examples.
🤖 AgenticAI — Hands-On Code Companion
Build real AI agents from scratch: tool-calling, RAG, memory, multi-agent systems, and production hardening — all running locally with Ollama.
👉📘 Start here: Coding Tutorial
📚 Notes on Agentic AI — Textbook Companion
The content below is a structured textbook-style guide: concepts, diagrams, design patterns, and production trade-offs — with direct mapping to runnable modules.
The text is CC BY-NC-SA 4.0 ; the code is MIT . See License at the end.
About
🤖 AgenticAI — Hands-On Code Companion
📚 Notes on Agentic AI — Textbook Companion
Basics of LLMs, Prompts and Tool Calls
What is an Agent?
Language Model (LLM)
Prompt Engineering
Tool Calling (Function Calling)
Foundation of Agentic AI
Retrieval-Augmented Generation (RAG)
Agentic Execution and Patterns
An Agent
DSPy, Embabel, and LlamaIndex
Why these three matter
How students should compare them
Production Aspects
Post-training (out of scope)
Hallucinations and Factfulness
Agent Benchmarks and Evaluation
Cost and Resources Utilization
Enterprise Suites and Protocols
MCP (Model Context Protocol)
A2A and Other Enterprise Concerns
I started teaching engineers about agentic AI. Over the period I realized that teaching practicing programmers about any new Framework has to be mindful of their prior learning. At times the wisdom gained so far feeds into the new framework and technology paradigm. At times the prior learnings hamper their ability to see nuances and retain curiosity about the new kid.
So my teaching style evolved into using lots of code first to drive in the new point of view. And then supplement that with details on the new tech and also commentary on where it has continuation, breakup, evolution and new beginnings. Over period of time my notes and call transcripts were long enough to motivate me into writing this. This text is very opinionated hence, it also assumes the student has prior understanding of basic building blocks and will dig out more when clues are provided. All the text is written by me, citations have been given where due.
When it came to coding examples I generated them using Copilot. However I realized they were too cryptic. So I took help from Sumit Toshniwal. He is one of the AI Engineers working at Actimize on Agentic projects. First he helped me get a feel of younger generation of developers :). Then he helped me simplify the code examples according to the intended learning outcome. We have kept the code samples minimal and self contained because I am expecting experienced developers will fill in the necessary design blanks. We thank NICE Actimize for being the org where we could do all this as one of the part of the work.
This repo is dedicated to Prof. Andrew Ng who is sharing AI related knowledge freely nurturing AI minds across the Globe.
Running the code alongside the text
Each module is a numbered sequence of small, self-contained scripts; the full index is in
code/README.md .
Basics of LLMs, Prompts and Tool Calls
This section introduces the concepts related to Agentic AI. We have grouped the concepts in accordance with the code samples in modules, so that learners can quickly test out the concepts.
Classically, an agent can be defined as a component that has some perception of its environment and has the ability to perform an action/task. A lot of your home automation devices fit into this definition. This also implies that AI is not a prerequisite for an agent. However, the wave of Agentic AI defines an Agent as a component that has the ability to perform some goal-oriented action with some sort of reasoning and planning capability (with or without direct perception of its environment). It goes without saying that LLMs have significantly enabled agents with the capability to reason and plan. This adds dynamism and adaptability to agents given that LLMs have a generic ability to reason about practically everything.
The hard part of achieving the action, however, is still accomplished by tools and retrieval systems that assist the agent. This underlines that agents are not first-class citizens of LLMs; LLMs can be made aware that agentic processing is happening to some degree.
A lot of agentic AI's potential comes from the application of traditional design patterns facilitated by tools like LangChain or LangGraph. These frameworks allow tighter control over processing flow by introducing concepts like task graphs, memory, and tools.
The final effect is a system that can plan and execute complex tasks with a good degree of certainty while offering adaptability for newer scenarios. We will revisit most of these concepts in the sections that follow.
Figure: each paradigm shift traded control for capability — agents buy adaptability at the price of determinism, which is why the engineering disciplines later in this book exist
graph TB
subgraph AGENT["🤖 AGENT"]
LLM["🧠 Language Model\n(Reasoning Engine)"]
Tools["🔧 Tools\n(Actions & APIs)"]
Memory["💾 Memory\n(State & History)"]
Prompt["📝 Prompts\n(Instructions)"]
end
ENV["🌍 Environment\n(Data, Users, Systems)"] -->|"Perception"| AGENT
AGENT -->|"Actions"| ENV
Prompt -->|"Instructs"| LLM
LLM -->|"Selects"| Tools
LLM -->|"Reads/Writes"| Memory
Tools -->|"Results"| LLM
Memory -->|"Context"| LLM
style LLM fill:#EF4444,stroke:#B91C1C,color:#fff,stroke-width:3px
style Tools fill:#10B981,stroke:#047857,color:#fff,stroke-width:3px
style Memory fill:#8B5CF6,stroke:#6D28D9,color:#fff,stroke-width:3px
style Prompt fill:#F59E0B,stroke:#B45309,color:#fff,stroke-width:3px
style ENV fill:#6366F1,stroke:#4338CA,color:#fff,stroke-width:3px
style AGENT fill:#F1F5F9,stroke:#3B82F6,stroke-width:3px
linkStyle default stroke:#475569,stroke-width:2px
Loading
Figure: anatomy of an agent — LLM, tools, memory, and prompts interacting with the environment
Figure: agent capability ladder — scope to the lowest rung that solves the problem
The LLM is the cognitive core of an agent. It processes natural language input and generates inferences. The instruction that gives data and direction to the LLM is called a prompt. At the same time, LLMs are trained on world data. This inherent information plus the prompt and the internal mathematical techniques that go into creating the LLM produce the effect of interpretation, inference, and reasoning. This effect of interpretation, inference, and reasoning is what makes LLMs popular and enables the whole Agentic phenomenon! A Survey of Reasoning with Foundation Models .
In this course, we use local LLMs via Ollama to avoid cloud dependencies. Most of the code demos show only the name of the LLM/URL when calls are made to it. However, LLMs also give us a few more parameters to tweak that can affect its outcome. Settings like top_p, top_k, and temperature help us control the randomness and sampling of the processing. There are more flags like reasoning effort, streaming, and verbosity that a given LLM might support; developers should read the documentation before moving to production.
For most use cases, these are not things we tweak daily in our code. In some cases, one can adjust the context length and the depth of reasoning. It is recommended to read the documentation of your LLM model to understand these settings. However, given the way LLM models and agentic use cases are evolving, we are good with the defaults for all practical purposes.
graph LR
Input["📝 Input\nPrompt"] --> Tokenize["🔤 Tokenize"]
Tokenize --> Encode["🔢 Encode\nEmbeddings"]
Encode --> Attend["🧠 Multi-Head\nAttention"]
Attend --> Decode["📊 Decode\nProbabilities"]
Decode --> Sample["🎲 Sample\n(temp, top_p)"]
Sample --> Output["✅ Output\nTokens"]
style Input fill:#6366F1,stroke:#4338CA,color:#fff,stroke-width:3px
style Tokenize fill:#3B82F6,stroke:#1E40AF,color:#fff,stroke-width:3px
style Encode fill:#8B5CF6,stroke:#6D28D9,color:#fff,stroke-width:3px
style Attend fill:#EF4444,stroke:#B91C1C,color:#fff,stroke-width:3px
style Decode fill:#F59E0B,stroke:#B45309,color:#fff,stroke-width:3px
style Sample fill:#EC4899,stroke:#BE185D,color:#fff,stroke-width:3px
style Output fill:#10B981,stroke:#047857,color:#fff,stroke-width:3px
linkStyle default stroke:#475569,stroke-width:2px
Loading
Figure: LLM request-response lifecycle — tokenize, encode, attend, decode, sample
Prompts are instructions to the LLM that shape its behavior. Prompt also significantly shapes the outcome in terms of the result as well as the format. In fact it is the prompt that actually causes the effect to take place. So, it is very important to master prompt engineering Paper ) to improve agentic outcomes.
First and foremost, ChatGPT has made many to believe that the chat interactions we have with it are how prompts are/should be. Statements like "English is your new programming language" has also fuelled this nonsense grasp of what prompts need to be.
When it comes to getting the LLM to produce the effect we want, it is the sum total of one or many sentences with examples and keywords that have to come together as a coherent instruction, agentic or not. Here is a good paper you must read Can Large Language Models Reason and Plan?
Suggestions to craft good prompts:
Role: This can include system, user, or assistant. Specifying the system role has an overarching effect. Usually user or assistant roles help us get the effect we want.
Task: Framing the prompt as a task directs the LLM into narrowing down the range of outcomes. However, task is a generic word; it can also be replaced with aim, job, or similar words. The task itself can be to think and reflect, a chain of actions, a review, or mid-conversation instructions for longer tasks.
Constraints: Framing prompts with constraints also narrows down the range of alternatives the LLM has for interpreting and processing the prompt. It g

[truncated]
