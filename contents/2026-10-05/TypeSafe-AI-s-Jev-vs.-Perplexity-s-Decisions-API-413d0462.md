---
source: "https://vercel.com/i/jev-vs-perplexity"
hn_url: "https://news.ycombinator.com/item?id=49961749"
title: "TypeSafe AI's Jev vs. Perplexity's Decisions API"
article_title: "TypeSafe AI's Jev vs. Perplexity's Decisions API - Vercel"
image: "https://images.ctfassets.net/e5382hct74si/1cH9kPmvxLbJt2XXVAPBNv/5f5110484a448e9c0f582dc6c5aad686/image.png"
author: "flashbrew"
captured_at: "2026-10-05T08:28:25Z"
capture_tool: "hn-digest"
hn_id: 49961749
score: 1
comments: 0
posted_at: "2026-10-05T07:31:16Z"
tags:
  - hacker-news
---

# TypeSafe AI's Jev vs. Perplexity's Decisions API

- HN: [49961749](https://news.ycombinator.com/item?id=49961749)
- Source: [vercel.com](https://vercel.com/i/jev-vs-perplexity)
- Score: 1
- Comments: 0
- Posted: 2026-10-05T07:31:16Z

## Translation

Title: TypeSafe AI's Jev vs. Perplexity's Decisions API
Article title: TypeSafe AI's Jev vs. Perplexity's Decisions API - Vercel
Description: Compare Jev and Perplexity's Decisions API by input support, answer types, deployment options, and the work required to switch providers.

Article text:
TypeSafe AI's Jev vs. Perplexity's Decisions API - Vercel Skip to content Copy Wordmark Copy Logo Download Brand Assets Brand Guidelines Products Agent Stack AI SDK
TypeSafe AI's Jev vs. Perplexity's Decisions API
Articles Build with AI 5 Oct 2026
TypeSafe AI's Jev and Perplexity's Decisions API both evaluate supplied state using typed questions. Jev works with text, including structured records. Perplexity also accepts images and publishes decision-model weights. For text workflows, compare their predictions on your own examples before deciding whether to switch.
The choice starts with the evidence your application needs and where you intend to run evaluation. Similar request fields can reduce integration work, but they don't establish that two models will make the same decisions.
Copy link to heading How do their interfaces compare?
Both follow the decision-model pattern : supply evidence, define the possible answers, and let application code determine the next action. Their direct APIs use the same names for the three question types.
Text, JSON objects, and arrays containing text
Hosted model identifier in this example
https://api.typesafe.ai/v1/systemone
https://api.perplexity.ai/v1/decisions
TypeSafe API key as a bearer token
Perplexity API key as a bearer token
These are the providers' direct interfaces. Jev also has an AI Gateway evaluation integration , which uses a different endpoint and model identifier. Keep that distinction explicit when adapting code from an existing Gateway application.
Copy link to heading When does image support change the decision?
Jev accepts textual state . If a workflow starts with a screenshot, you would need to extract the relevant information into text before passing it to Jev. That extra step becomes part of the system you must evaluate.
Perplexity's image input gives you a way to include the screenshot itself. For a support ticket that says “This keeps happening,” the visible error message and the screen around it may supply the information needed to select a queue.
Compare complete workflows when assessing this difference. Sending a screenshot to one model and only the customer's vague message to another measures access to evidence as well as model behavior. For a text-only comparison, give both providers the same record. For an image workflow, evaluate the extraction step alongside the final classification if you use one.
Copy link to heading What changes in a text-routing request?
The following example sends the same invented ticket and routing question to both providers. It prints their answers without assuming which queue either model selects. Use Node.js 22 or later, set both API keys in your server environment, and save the file as compare-decisions.mjs .
const providers = [ { name : "Jev" , endpoint : "https://api.typesafe.ai/v1/systemone" , model : "jev-latest" , apiKey : process . env . TYPESAFE_API_KEY , } , { name : "Perplexity" , endpoint : "https://api.perplexity.ai/v1/decisions" , model : "pplx-decider-v1-27b" , apiKey : process . env . PERPLEXITY_API_KEY , } , ] ;
const task = { state : { subject : "Can't open my invoices" , message : "My password reset link has expired. I need to sign in to download an invoice." , } , questions : { queue : { type : "choice" , instructions : "Choose the queue that can resolve the immediate blocker." , criteria : { billing : "Incorrect charges or invoice contents, with no sign-in blocker." , account_access : "Sign-in, password reset, or account recovery problems." , review : "The blocker is unclear or falls outside these queues." , } , } , } , } ;
for ( const provider of providers ) { if ( ! provider . apiKey ) throw new Error ( ` Set the API key for ${ provider . name } . ` ) ; }
for ( const provider of providers ) { const response = await fetch ( provider . endpoint , { method : "POST" , headers : { Authorization : ` Bearer ${ provider . apiKey } ` , "Content-Type" : "application/json" , } , body : JSON . stringify ( { model : provider . model , ... task } ) , signal : AbortSignal . timeout ( 30_000 ) , } ) ;
if ( ! response . ok ) { throw new Error ( ` ${ provider . name } returned HTTP ${ response . status } ` ) ; }
const result = await response . json ( ) ; console . log ( provider . name , result . model , result . answers . queue ) ; }
Run node compare-decisions.mjs to inspect the two responses. The question gives account access priority when signing in blocks the customer from reaching an invoice. Without that rule, a billing prediction could reflect an ambiguous task definition rather than a failure to understand the message.
The shared body uses fields defined by both TypeSafe's API and Perplexity's API . That overlap supports this specific example. Before adapting a larger integration, check its question structures and response handling rather than treating the two APIs as universally interchangeable.
Copy link to heading Can you reuse the same thresholds?
Reuse the business policy, then measure how each model implements it. “Automatically route only when the error rate is acceptable” can remain the goal. The numeric cutoff needed to reach that goal may differ.
For Choice questions, distinguish the selected option's probability from the separate confidence field. TypeSafe derives confidence from the probability distribution . Even when two APIs expose similarly named fields, matching names don't establish identical behavior on your dataset.
Keep the field used by your routing rule explicit. If the current rule tests probabilities.account_access , switching it to confidence changes the rule as well as the model. You would then need to determine which change caused any difference in routing.
Use a labeled validation set to select thresholds for each provider, and reserve other examples for the final comparison. Report both the errors among accepted routes and the proportion of tickets held for review. Sending nearly everything to a person can produce few automatic errors while providing little automation.
Copy link to heading How do the deployment options affect your application?
Jev's documented AI Gateway evaluation path lets a Vercel application use the gateway's evaluation interface. If you already use that path, include the response shape and authentication setup in your migration review. The Jev form router guide provides an existing implementation to build from.
Perplexity provides a hosted Decisions endpoint and downloadable model weights with inference code . The code above calls the hosted service. Running the model yourself adds responsibility for serving it, managing capacity, and maintaining the deployment.
Open weights can matter when operating the model on your own infrastructure is a requirement. They don't tell you whether your existing application can adopt the supplied inference code unchanged. Review the serving interface and infrastructure separately from the question definitions.
Copy link to heading How should you evaluate a switch?
Start with a set of real cases that your team has labeled, including examples that are difficult to assign. For support routing, include messages that mention billing but are blocked by login, tickets with two unrelated requests, and cases that should go to review.
Use the same category definitions for the initial text comparison. Keep a record of each provider's answer, the expected queue, and the eventual action your policy would take. Review disagreements individually. Sometimes the model chose the wrong category; sometimes the rubric left two plausible routes.
Then test the candidate against a held-out set using its own selected thresholds. Include failure cases at the application boundary, such as a missing answer or an unsuccessful HTTP response. Those cases need a review or retry path, not a default queue that conceals the failed evaluation.
Before replacing an existing router, run the candidate without applying its proposed routes. Compare its suggestions with the workflow's actual outcomes. This lets you inspect differences without sending customer work to new queues during the evaluation.
Copy link to heading What doesn't this comparison establish?
Interface compatibility doesn't establish equal predictions, and a successful request doesn't establish accuracy. The examples above demonstrate how to submit a comparable task. Choosing a provider for production still requires labeled cases that reflect your application.
For work that needs a written reply, either decision API would be one part of a larger workflow. Keep the routing decision separate from the model that drafts the response so you can evaluate each task against its own requirements.
Copy link to heading Frequently asked questions
Copy link to heading Is Perplexity's Decisions API a drop-in replacement for Jev?
No. The direct APIs share several request concepts, but changing providers also changes the endpoint, model, and credentials. Check the fields your integration uses and reevaluate the resulting decisions before switching application traffic.
Copy link to heading Can both Jev and Perplexity evaluate screenshots directly?
Perplexity accepts image inputs in its Decisions API, while Jev's documented state input is text. Using Jev for a screenshot workflow therefore requires a step that turns the relevant image content into text.
Copy link to heading Can I keep my Jev routing threshold when switching models?
Treat the existing cutoff as a candidate to test. Select a threshold using labeled validation examples for the new model, then assess accepted-route errors and review volume on separate examples.
Copy link to heading How can I compare these APIs from a Vercel application?
Call the providers from server-side code and keep their credentials out of the browser. Jev also has a documented AI Gateway evaluation integration; the shared HTTP example in this article uses each provider's direct endpoint.
Copy link to heading Related resources
Classify images with Perplexity's Decisions API
Route form submissions with Jev and AI SDK
Perplexity's decision-model weights and inference example
Perplexity Decisions API reference
Build with AI Jev vs. Laya vs. Liquid d1: Which decision model should you use?
Compare Jev, Laya, and Liquid d1 by API compatibility, deployment options, and probability handling, with a product-catalog example through AI Gateway.
Build with AI What is Perplexity's Decisions API?
Perplexity's Decisions API evaluates text and images with typed answers. Learn how its questions, probabilities, and open model fit application workflows.
Build with AI How to classify images with Perplexity's Decisions API
Build a support screenshot classifier with Perplexity's Decisions API, image resizing, typed categories, and a review path for uncertain results.

## Original Extract

Compare Jev and Perplexity's Decisions API by input support, answer types, deployment options, and the work required to switch providers.

TypeSafe AI's Jev vs. Perplexity's Decisions API - Vercel Skip to content Copy Wordmark Copy Logo Download Brand Assets Brand Guidelines Products Agent Stack AI SDK
TypeSafe AI's Jev vs. Perplexity's Decisions API
Articles Build with AI 5 Oct 2026
TypeSafe AI's Jev and Perplexity's Decisions API both evaluate supplied state using typed questions. Jev works with text, including structured records. Perplexity also accepts images and publishes decision-model weights. For text workflows, compare their predictions on your own examples before deciding whether to switch.
The choice starts with the evidence your application needs and where you intend to run evaluation. Similar request fields can reduce integration work, but they don't establish that two models will make the same decisions.
Copy link to heading How do their interfaces compare?
Both follow the decision-model pattern : supply evidence, define the possible answers, and let application code determine the next action. Their direct APIs use the same names for the three question types.
Text, JSON objects, and arrays containing text
Hosted model identifier in this example
https://api.typesafe.ai/v1/systemone
https://api.perplexity.ai/v1/decisions
TypeSafe API key as a bearer token
Perplexity API key as a bearer token
These are the providers' direct interfaces. Jev also has an AI Gateway evaluation integration , which uses a different endpoint and model identifier. Keep that distinction explicit when adapting code from an existing Gateway application.
Copy link to heading When does image support change the decision?
Jev accepts textual state . If a workflow starts with a screenshot, you would need to extract the relevant information into text before passing it to Jev. That extra step becomes part of the system you must evaluate.
Perplexity's image input gives you a way to include the screenshot itself. For a support ticket that says “This keeps happening,” the visible error message and the screen around it may supply the information needed to select a queue.
Compare complete workflows when assessing this difference. Sending a screenshot to one model and only the customer's vague message to another measures access to evidence as well as model behavior. For a text-only comparison, give both providers the same record. For an image workflow, evaluate the extraction step alongside the final classification if you use one.
Copy link to heading What changes in a text-routing request?
The following example sends the same invented ticket and routing question to both providers. It prints their answers without assuming which queue either model selects. Use Node.js 22 or later, set both API keys in your server environment, and save the file as compare-decisions.mjs .
const providers = [ { name : "Jev" , endpoint : "https://api.typesafe.ai/v1/systemone" , model : "jev-latest" , apiKey : process . env . TYPESAFE_API_KEY , } , { name : "Perplexity" , endpoint : "https://api.perplexity.ai/v1/decisions" , model : "pplx-decider-v1-27b" , apiKey : process . env . PERPLEXITY_API_KEY , } , ] ;
const task = { state : { subject : "Can't open my invoices" , message : "My password reset link has expired. I need to sign in to download an invoice." , } , questions : { queue : { type : "choice" , instructions : "Choose the queue that can resolve the immediate blocker." , criteria : { billing : "Incorrect charges or invoice contents, with no sign-in blocker." , account_access : "Sign-in, password reset, or account recovery problems." , review : "The blocker is unclear or falls outside these queues." , } , } , } , } ;
for ( const provider of providers ) { if ( ! provider . apiKey ) throw new Error ( ` Set the API key for ${ provider . name } . ` ) ; }
for ( const provider of providers ) { const response = await fetch ( provider . endpoint , { method : "POST" , headers : { Authorization : ` Bearer ${ provider . apiKey } ` , "Content-Type" : "application/json" , } , body : JSON . stringify ( { model : provider . model , ... task } ) , signal : AbortSignal . timeout ( 30_000 ) , } ) ;
if ( ! response . ok ) { throw new Error ( ` ${ provider . name } returned HTTP ${ response . status } ` ) ; }
const result = await response . json ( ) ; console . log ( provider . name , result . model , result . answers . queue ) ; }
Run node compare-decisions.mjs to inspect the two responses. The question gives account access priority when signing in blocks the customer from reaching an invoice. Without that rule, a billing prediction could reflect an ambiguous task definition rather than a failure to understand the message.
The shared body uses fields defined by both TypeSafe's API and Perplexity's API . That overlap supports this specific example. Before adapting a larger integration, check its question structures and response handling rather than treating the two APIs as universally interchangeable.
Copy link to heading Can you reuse the same thresholds?
Reuse the business policy, then measure how each model implements it. “Automatically route only when the error rate is acceptable” can remain the goal. The numeric cutoff needed to reach that goal may differ.
For Choice questions, distinguish the selected option's probability from the separate confidence field. TypeSafe derives confidence from the probability distribution . Even when two APIs expose similarly named fields, matching names don't establish identical behavior on your dataset.
Keep the field used by your routing rule explicit. If the current rule tests probabilities.account_access , switching it to confidence changes the rule as well as the model. You would then need to determine which change caused any difference in routing.
Use a labeled validation set to select thresholds for each provider, and reserve other examples for the final comparison. Report both the errors among accepted routes and the proportion of tickets held for review. Sending nearly everything to a person can produce few automatic errors while providing little automation.
Copy link to heading How do the deployment options affect your application?
Jev's documented AI Gateway evaluation path lets a Vercel application use the gateway's evaluation interface. If you already use that path, include the response shape and authentication setup in your migration review. The Jev form router guide provides an existing implementation to build from.
Perplexity provides a hosted Decisions endpoint and downloadable model weights with inference code . The code above calls the hosted service. Running the model yourself adds responsibility for serving it, managing capacity, and maintaining the deployment.
Open weights can matter when operating the model on your own infrastructure is a requirement. They don't tell you whether your existing application can adopt the supplied inference code unchanged. Review the serving interface and infrastructure separately from the question definitions.
Copy link to heading How should you evaluate a switch?
Start with a set of real cases that your team has labeled, including examples that are difficult to assign. For support routing, include messages that mention billing but are blocked by login, tickets with two unrelated requests, and cases that should go to review.
Use the same category definitions for the initial text comparison. Keep a record of each provider's answer, the expected queue, and the eventual action your policy would take. Review disagreements individually. Sometimes the model chose the wrong category; sometimes the rubric left two plausible routes.
Then test the candidate against a held-out set using its own selected thresholds. Include failure cases at the application boundary, such as a missing answer or an unsuccessful HTTP response. Those cases need a review or retry path, not a default queue that conceals the failed evaluation.
Before replacing an existing router, run the candidate without applying its proposed routes. Compare its suggestions with the workflow's actual outcomes. This lets you inspect differences without sending customer work to new queues during the evaluation.
Copy link to heading What doesn't this comparison establish?
Interface compatibility doesn't establish equal predictions, and a successful request doesn't establish accuracy. The examples above demonstrate how to submit a comparable task. Choosing a provider for production still requires labeled cases that reflect your application.
For work that needs a written reply, either decision API would be one part of a larger workflow. Keep the routing decision separate from the model that drafts the response so you can evaluate each task against its own requirements.
Copy link to heading Frequently asked questions
Copy link to heading Is Perplexity's Decisions API a drop-in replacement for Jev?
No. The direct APIs share several request concepts, but changing providers also changes the endpoint, model, and credentials. Check the fields your integration uses and reevaluate the resulting decisions before switching application traffic.
Copy link to heading Can both Jev and Perplexity evaluate screenshots directly?
Perplexity accepts image inputs in its Decisions API, while Jev's documented state input is text. Using Jev for a screenshot workflow therefore requires a step that turns the relevant image content into text.
Copy link to heading Can I keep my Jev routing threshold when switching models?
Treat the existing cutoff as a candidate to test. Select a threshold using labeled validation examples for the new model, then assess accepted-route errors and review volume on separate examples.
Copy link to heading How can I compare these APIs from a Vercel application?
Call the providers from server-side code and keep their credentials out of the browser. Jev also has a documented AI Gateway evaluation integration; the shared HTTP example in this article uses each provider's direct endpoint.
Copy link to heading Related resources
Classify images with Perplexity's Decisions API
Route form submissions with Jev and AI SDK
Perplexity's decision-model weights and inference example
Perplexity Decisions API reference
Build with AI Jev vs. Laya vs. Liquid d1: Which decision model should you use?
Compare Jev, Laya, and Liquid d1 by API compatibility, deployment options, and probability handling, with a product-catalog example through AI Gateway.
Build with AI What is Perplexity's Decisions API?
Perplexity's Decisions API evaluates text and images with typed answers. Learn how its questions, probabilities, and open model fit application workflows.
Build with AI How to classify images with Perplexity's Decisions API
Build a support screenshot classifier with Perplexity's Decisions API, image resizing, typed categories, and a review path for uncertain results.
