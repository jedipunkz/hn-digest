---
source: "https://www.revenuecat.com/blog/engineering/lessons-paywalls-ai-editor"
hn_url: "https://news.ycombinator.com/item?id=50467731"
title: "Don't build your own harness: lessons from building Paywalls AI Editor"
article_title: "Don’t Build Your Own Agent Harness: Lessons From Paywalls AI Editor | RevenueCat"
image: "https://cdn.sanity.io/images/c3qnx9b0/production/149685e32fca9e8765e98aef1411b0fc35e95d3a-1600x800.png?w=1200&fit=max"
author: "azvaliev"
captured_at: "2026-11-12T15:01:34Z"
capture_tool: "hn-digest"
hn_id: 50467731
score: 3
comments: 1
posted_at: "2026-11-12T14:23:54Z"
tags:
  - hacker-news
---

# Don't build your own harness: lessons from building Paywalls AI Editor

- HN: [50467731](https://news.ycombinator.com/item?id=50467731)
- Source: [www.revenuecat.com](https://www.revenuecat.com/blog/engineering/lessons-paywalls-ai-editor)
- Score: 3
- Comments: 1
- Posted: 2026-11-12T14:23:54Z

## Translation

Title: Don't build your own harness: lessons from building Paywalls AI Editor
Article title: Don’t Build Your Own Agent Harness: Lessons From Paywalls AI Editor | RevenueCat
Description: Why RevenueCat swapped a custom JSON harness for a coding-agent loop with JSX, and how it lifted positive feedback 83% on its AI paywall editor.

Article text:
Don’t Build Your Own Agent Harness: Lessons From Paywalls AI Editor | RevenueCat Back to the RevenueCat homepage Why RevenueCat?
Don’t build your own harness: lessons from building Paywalls AI Editor
How one Slack message led us to rebuild our paywall agent
RevenueCat rebuilt its Paywalls AI Editor after a custom JSON-patching harness produced low-quality designs. Converting paywalls to JSX and giving the agent a standard coding loop with a compiler cut the system prompt by over half, raised positive feedback 83%, and reduced negative feedback 40% across about 46,000 requests.
Earlier this year, we shipped an AI editor for our remotely configurable paywalls product. Users can work with our agent to come up with a new paywall or edit an existing one.
On paper we fulfilled user requests, and adoption was growing rapidly. But there was a deeper issue that none of our eval or metrics could capture.
An inflection point: “it’s not even fair”
A pattern of issues started trickling in via user reports and internal feedback. On paper, the agent would accomplish the task the user asked, like creating the paywall, adding or overhauling a section of it. But the result it produced lacked the level of taste and quality that you’d expect a frontier agent to be capable of.
We worked to address this issue in several ways:
We tried different models from various providers, with higher effort and reasoning levels
We tweaked the system prompt and skills
We injected messages with screenshots of the paywall, pushing the agent to check its work and maintain a high bar for quality
There were notable improvements, but the gap between what users wanted, and what the AI delivered, remained massive.
We then got some feedback from Ryan, one of our Product Managers:
(left) our old paywall editor, (right) using Fable in Claude Code(FYI: Astra is our internal term for the paywall agent)
This was a brutal inflection point. The feedback was originally about our agent getting the app’s name wrong. But the bigger issue was blatantly obvious.
How does our agent not understand that the paywall it produced is atrocious? Spacing, alignment, card within a card… there were too many issues to name. How did it think this was acceptable?
And how can Claude Code just one shot something that looks great while our agent flounders? This thought haunted me. No one would use our paywall editor if you could get that much better results with a simple Claude Code prompt.
The old agent technically worked
To understand why our old agent was struggling, you first have to see how it worked.
Beginning with the data shape. Some remote paywall products use HTML, CSS and JS to render on all platforms. This provides the most flexibility, at the cost of feeling native, appearance, and performance.
We opted for native (well, a bit of both now ). Our paywalls are stored as JSON trees: each node is a component with its own attributes. When your paywall is rendered, the component tree is mapped into the native components.
Here’s what a purchase button component can look like in our structure:
{
"type" : "purchase_button" ,
"stack" : {
"type" : "stack" ,
"background" : { "type" : "color" , "value" : { "light" : { "type" : "hex" , "value" : "#ff8a00ff" } } },
"shape" : { "type" : "pill" , "corners" : null },
"padding" : { "top" : 12 , "bottom" : 12 , "leading" : 8 , "trailing" : 8 },
"components" : [
{
"type" : "text" ,
"text_lid" : "PBcCWPIKbT" ,
"color" : { "light" : { "type" : "hex" , "value" : "#ffffffff" } },
"font_weight" : "semibold" ,
"margin" : { "top" : 0 , "bottom" : 0 , "leading" : 0 , "trailing" : 0 }
}
]
}
}
// ... one of 20+ components on this screen Our agent had a very custom harness (set of tools, prompts and loop around the model). It understood the shape through an annotated JSON schema, and would operate on the paywall JSON directly, manipulating it with a set of tools.
This was a reasonable design for the first iteration.
The targeted JSON operations prevented it from making simple formatting mistakes we’d had in earlier prototypes
The schema showed it what was valid
Skills gave it reusable patterns and guidance
But we were asking a lot of the model: it needed to learn this new schema on the fly for representing UI, along with the new tools. It needed to translate the user request into visual intent, then translate that visual intent into a series of tool operations on the JSON. And lastly, it needed to self-evaluate using a screenshot.
It was a bespoke workflow and syntax — a far cry from the Jetpack Compose, JSX and SwiftUI use cases it had been extensively trained on.
Still, I couldn’t shake the feeling that we were not just “one more prompt” away from closing this gap. Something was fundamentally wrong.
I read through Claude Code’s system prompt; there’s very little in there telling Claude how to make tasteful design, and certainly nothing specific to paywalls. Yet it achieved far better results even with models the harness wasn’t built for, like GPT.
Anthropic has a frontend design skill but, in my experience, the newer models still create great UI without it.
This led me to suspect the taste was being lost in translation. In a minimal harness, the models could express taste into code. But not into JSON patching tool calls…
What if our paywall could just be code? The model knows how to build a great UI in JSX. And JSX components are a tree very similar to ours: each node has JSON attributes and an optional set of children.
So we tried it. Here’s that same purchase button from earlier, as JSX:
< PurchaseButton >
< Stack background = "#ff8a00ff" padding = {{ vertical : 12 , horizontal : 8 }} >
< Text text = "Continue" fontWeight = "semibold" color = "#ffffff" />
</ Stack >
</ PurchaseButton > It turned out to be fairly simple. A transpile step turns the JSON into this intermediate JSX, which the agent can edit as a file. Once it’s done editing, a strict compiler runs, flagging any issues that could prevent the paywall from displaying or functioning correctly. When there are no blocking issues, the compiler turns the JSX back into the same JSON structure our backend & SDKs expect.
Two days after rock bottom, the new prototype set a record on our blank paywalls eval dataset. That’s when we knew we were onto something.
If the theory is that leaning into coding agent RL will outperform learning a new harness on the fly, why not lean in fully. Let’s give the agent a full workspace, simple mechanisms to read, and additional context as needed.
So we did just that. We gave it a folder of snippets to see examples of how to do certain component structures, which it can copy-paste as needed. It doesn’t need the full documentation for every component on every request , so I split it into multiple d.ts files that can be lazily loaded, instead of cramming 175KB of JSON schema at the start of every session. This meant paywall translations just became entries in a strings/ folder. The agent can keep context about the app, like what the app does and who it is for, in a context.md .
And most importantly, this gave us room to expand. New capabilities can be represented with files/folders, and since things are lazily loaded in, there is minimal additional cost/overhead to do so.
/ workspace
├── context.md
├── paywall.tsx
├── snippets / (read - only)
│ ├── ctas / purchase_button.tsx
│ ├── feature_lists / checklist.tsx
│ └── ... ( 28 snippets)
├── strings /
│ ├── es_ES. json ( 14 strings)
│ └── fr_FR. json ( 14 strings, 2 stale)
└── types / (read - only)
├── _readme.d.ts
├── base.d.ts
├── components / layout.d.ts
├── components /... d.ts
└── overrides.d.ts We didn’t need to reinvent our own tools to work with this format. We used what the model already knows: a free loop with Read , Write , Edit and a compiler. The compiler even feels like tsc : giving line numbers and diagnostics with guidance when there are issues in the paywall code, and handing back a screenshot of the paywall when it succeeds.
With a familiar format and harness, We were also able to remove over half of our system prompt and almost all of our skills. Most of those skills and prompts were compensating for a subpar format. As we continue to iterate, we focus on improving format first, then only prompt our way out if we must.
The prototype broke eval records in two days. Getting that to production was a full team effort that took several weeks.
First, we had to build our confidence that this was production ready.
Our existing offline evaluations were a good start. Your users always use the agent in unexpected ways though, so we sampled hundreds of production traces and ran them against the new agent and created a dataset of new vs. old results. Our LLM judge mechanism was able to catch functional regressions: comparing the behavior, responses and results of the new vs. old agent, and ensuring the new agent was equal or superior in all cases. We also recruited many internal humans-as-judge to blind vote which version had better taste.
Only then did we started rolling out to external users — first, those who had given us feedback about taste and capability gaps.
Once we felt confident that this agent would be a definite step up in our production quality across the board, we had some infrastructure gaps to close.
The new agent was working hard. Between the large amount of turns it was using and the compilation work, we had to allocate additional CPU and memory. We also switched to Claude models: both Claude and GPT models perform well in the new harness, but Claude models currently have an edge in taste. Our new inference provider had different failure modes: stalled streams, empty responses, and increased downtime. So we hardened our retry and fallback mechanisms around the LLM calls.
After pushing through, we finished rolling out the new editor to 100% of users on September 2nd.
Survey says: compared to the two weeks prior to rollout, across approximately 46K requests, our rate of positive feedback grew by 83% and negative feedback fell 40% .
We saw little difference with the new agent in our simpler requests like changing copy, adding a simple button, checking app store compliance, and automatic translations. But the gains show up in large paywall edits, new paywall creation, and anything that requires taste .
Going back to Ryan’s paywall tells the story.
If you’re building your own agent
Lean into the model’s strengths: models have been trained extensively with certain tasks, environments and tools. Lean into that. A great example of where this is starting to show up is code mode with MCP tools; letting the model write code to string together data from several tool calls, then compose and manipulate it.
Needing lots of brittle system prompts and skill changes is a smell: when we got the format right, we cut the system prompt in half. We stopped fighting with tuning verbiage changes in a prompt to get the desired result. And as we continue to optimize and add features, we’re looking at making the interface intuitive first, so we don't stack more debt in this category.
If applicable, measure taste, not just completion: the problem flew under the radar as we didn’t have any automated, clearly-correlated measures for this. We still don’t have a great one. What’s helped is blind voting by humans between two different options, in addition to listening to user feedback.
Take big swings : if your agent has decent evaluation mechanisms, the cost of prototyping and measuring big agent changes is relatively cheap. Try it, test it and see if there’s a better way. The payoff can be massive.
Latency may be accepted for quality: the new agent takes more turns, reading in context, editing, compiling and iterating. This has significantly increased our latency. At first we anticipated users would be unhappy with this, but from the data it appears the quality of results is a far more important factor. (Of course, this can vary based on your type of application and users.)
Test the Paywal

[truncated]

## Original Extract

Why RevenueCat swapped a custom JSON harness for a coding-agent loop with JSX, and how it lifted positive feedback 83% on its AI paywall editor.

Don’t Build Your Own Agent Harness: Lessons From Paywalls AI Editor | RevenueCat Back to the RevenueCat homepage Why RevenueCat?
Don’t build your own harness: lessons from building Paywalls AI Editor
How one Slack message led us to rebuild our paywall agent
RevenueCat rebuilt its Paywalls AI Editor after a custom JSON-patching harness produced low-quality designs. Converting paywalls to JSX and giving the agent a standard coding loop with a compiler cut the system prompt by over half, raised positive feedback 83%, and reduced negative feedback 40% across about 46,000 requests.
Earlier this year, we shipped an AI editor for our remotely configurable paywalls product. Users can work with our agent to come up with a new paywall or edit an existing one.
On paper we fulfilled user requests, and adoption was growing rapidly. But there was a deeper issue that none of our eval or metrics could capture.
An inflection point: “it’s not even fair”
A pattern of issues started trickling in via user reports and internal feedback. On paper, the agent would accomplish the task the user asked, like creating the paywall, adding or overhauling a section of it. But the result it produced lacked the level of taste and quality that you’d expect a frontier agent to be capable of.
We worked to address this issue in several ways:
We tried different models from various providers, with higher effort and reasoning levels
We tweaked the system prompt and skills
We injected messages with screenshots of the paywall, pushing the agent to check its work and maintain a high bar for quality
There were notable improvements, but the gap between what users wanted, and what the AI delivered, remained massive.
We then got some feedback from Ryan, one of our Product Managers:
(left) our old paywall editor, (right) using Fable in Claude Code(FYI: Astra is our internal term for the paywall agent)
This was a brutal inflection point. The feedback was originally about our agent getting the app’s name wrong. But the bigger issue was blatantly obvious.
How does our agent not understand that the paywall it produced is atrocious? Spacing, alignment, card within a card… there were too many issues to name. How did it think this was acceptable?
And how can Claude Code just one shot something that looks great while our agent flounders? This thought haunted me. No one would use our paywall editor if you could get that much better results with a simple Claude Code prompt.
The old agent technically worked
To understand why our old agent was struggling, you first have to see how it worked.
Beginning with the data shape. Some remote paywall products use HTML, CSS and JS to render on all platforms. This provides the most flexibility, at the cost of feeling native, appearance, and performance.
We opted for native (well, a bit of both now ). Our paywalls are stored as JSON trees: each node is a component with its own attributes. When your paywall is rendered, the component tree is mapped into the native components.
Here’s what a purchase button component can look like in our structure:
{
"type" : "purchase_button" ,
"stack" : {
"type" : "stack" ,
"background" : { "type" : "color" , "value" : { "light" : { "type" : "hex" , "value" : "#ff8a00ff" } } },
"shape" : { "type" : "pill" , "corners" : null },
"padding" : { "top" : 12 , "bottom" : 12 , "leading" : 8 , "trailing" : 8 },
"components" : [
{
"type" : "text" ,
"text_lid" : "PBcCWPIKbT" ,
"color" : { "light" : { "type" : "hex" , "value" : "#ffffffff" } },
"font_weight" : "semibold" ,
"margin" : { "top" : 0 , "bottom" : 0 , "leading" : 0 , "trailing" : 0 }
}
]
}
}
// ... one of 20+ components on this screen Our agent had a very custom harness (set of tools, prompts and loop around the model). It understood the shape through an annotated JSON schema, and would operate on the paywall JSON directly, manipulating it with a set of tools.
This was a reasonable design for the first iteration.
The targeted JSON operations prevented it from making simple formatting mistakes we’d had in earlier prototypes
The schema showed it what was valid
Skills gave it reusable patterns and guidance
But we were asking a lot of the model: it needed to learn this new schema on the fly for representing UI, along with the new tools. It needed to translate the user request into visual intent, then translate that visual intent into a series of tool operations on the JSON. And lastly, it needed to self-evaluate using a screenshot.
It was a bespoke workflow and syntax — a far cry from the Jetpack Compose, JSX and SwiftUI use cases it had been extensively trained on.
Still, I couldn’t shake the feeling that we were not just “one more prompt” away from closing this gap. Something was fundamentally wrong.
I read through Claude Code’s system prompt; there’s very little in there telling Claude how to make tasteful design, and certainly nothing specific to paywalls. Yet it achieved far better results even with models the harness wasn’t built for, like GPT.
Anthropic has a frontend design skill but, in my experience, the newer models still create great UI without it.
This led me to suspect the taste was being lost in translation. In a minimal harness, the models could express taste into code. But not into JSON patching tool calls…
What if our paywall could just be code? The model knows how to build a great UI in JSX. And JSX components are a tree very similar to ours: each node has JSON attributes and an optional set of children.
So we tried it. Here’s that same purchase button from earlier, as JSX:
< PurchaseButton >
< Stack background = "#ff8a00ff" padding = {{ vertical : 12 , horizontal : 8 }} >
< Text text = "Continue" fontWeight = "semibold" color = "#ffffff" />
</ Stack >
</ PurchaseButton > It turned out to be fairly simple. A transpile step turns the JSON into this intermediate JSX, which the agent can edit as a file. Once it’s done editing, a strict compiler runs, flagging any issues that could prevent the paywall from displaying or functioning correctly. When there are no blocking issues, the compiler turns the JSX back into the same JSON structure our backend & SDKs expect.
Two days after rock bottom, the new prototype set a record on our blank paywalls eval dataset. That’s when we knew we were onto something.
If the theory is that leaning into coding agent RL will outperform learning a new harness on the fly, why not lean in fully. Let’s give the agent a full workspace, simple mechanisms to read, and additional context as needed.
So we did just that. We gave it a folder of snippets to see examples of how to do certain component structures, which it can copy-paste as needed. It doesn’t need the full documentation for every component on every request , so I split it into multiple d.ts files that can be lazily loaded, instead of cramming 175KB of JSON schema at the start of every session. This meant paywall translations just became entries in a strings/ folder. The agent can keep context about the app, like what the app does and who it is for, in a context.md .
And most importantly, this gave us room to expand. New capabilities can be represented with files/folders, and since things are lazily loaded in, there is minimal additional cost/overhead to do so.
/ workspace
├── context.md
├── paywall.tsx
├── snippets / (read - only)
│ ├── ctas / purchase_button.tsx
│ ├── feature_lists / checklist.tsx
│ └── ... ( 28 snippets)
├── strings /
│ ├── es_ES. json ( 14 strings)
│ └── fr_FR. json ( 14 strings, 2 stale)
└── types / (read - only)
├── _readme.d.ts
├── base.d.ts
├── components / layout.d.ts
├── components /... d.ts
└── overrides.d.ts We didn’t need to reinvent our own tools to work with this format. We used what the model already knows: a free loop with Read , Write , Edit and a compiler. The compiler even feels like tsc : giving line numbers and diagnostics with guidance when there are issues in the paywall code, and handing back a screenshot of the paywall when it succeeds.
With a familiar format and harness, We were also able to remove over half of our system prompt and almost all of our skills. Most of those skills and prompts were compensating for a subpar format. As we continue to iterate, we focus on improving format first, then only prompt our way out if we must.
The prototype broke eval records in two days. Getting that to production was a full team effort that took several weeks.
First, we had to build our confidence that this was production ready.
Our existing offline evaluations were a good start. Your users always use the agent in unexpected ways though, so we sampled hundreds of production traces and ran them against the new agent and created a dataset of new vs. old results. Our LLM judge mechanism was able to catch functional regressions: comparing the behavior, responses and results of the new vs. old agent, and ensuring the new agent was equal or superior in all cases. We also recruited many internal humans-as-judge to blind vote which version had better taste.
Only then did we started rolling out to external users — first, those who had given us feedback about taste and capability gaps.
Once we felt confident that this agent would be a definite step up in our production quality across the board, we had some infrastructure gaps to close.
The new agent was working hard. Between the large amount of turns it was using and the compilation work, we had to allocate additional CPU and memory. We also switched to Claude models: both Claude and GPT models perform well in the new harness, but Claude models currently have an edge in taste. Our new inference provider had different failure modes: stalled streams, empty responses, and increased downtime. So we hardened our retry and fallback mechanisms around the LLM calls.
After pushing through, we finished rolling out the new editor to 100% of users on September 2nd.
Survey says: compared to the two weeks prior to rollout, across approximately 46K requests, our rate of positive feedback grew by 83% and negative feedback fell 40% .
We saw little difference with the new agent in our simpler requests like changing copy, adding a simple button, checking app store compliance, and automatic translations. But the gains show up in large paywall edits, new paywall creation, and anything that requires taste .
Going back to Ryan’s paywall tells the story.
If you’re building your own agent
Lean into the model’s strengths: models have been trained extensively with certain tasks, environments and tools. Lean into that. A great example of where this is starting to show up is code mode with MCP tools; letting the model write code to string together data from several tool calls, then compose and manipulate it.
Needing lots of brittle system prompts and skill changes is a smell: when we got the format right, we cut the system prompt in half. We stopped fighting with tuning verbiage changes in a prompt to get the desired result. And as we continue to optimize and add features, we’re looking at making the interface intuitive first, so we don't stack more debt in this category.
If applicable, measure taste, not just completion: the problem flew under the radar as we didn’t have any automated, clearly-correlated measures for this. We still don’t have a great one. What’s helped is blind voting by humans between two different options, in addition to listening to user feedback.
Take big swings : if your agent has decent evaluation mechanisms, the cost of prototyping and measuring big agent changes is relatively cheap. Try it, test it and see if there’s a better way. The payoff can be massive.
Latency may be accepted for quality: the new agent takes more turns, reading in context, editing, compiling and iterating. This has significantly increased our latency. At first we anticipated users would be unhappy with this, but from the data it appears the quality of results is a far more important factor. (Of course, this can vary based on your type of application and users.)
Test the Paywal

[truncated]
