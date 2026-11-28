---
source: "https://platform.claude.com/docs/en/build-with-claude/refusals-and-fallback"
hn_url: "https://news.ycombinator.com/item?id=50675972"
title: "Refusals and Fallback of Anthropic Models"
article_title: "Refusals and fallback - Claude Platform Docs"
image: "https://platform.claude.com/web-api/og/docs/en/build-with-claude/refusals-and-fallback?design-rev=5"
author: "pretext"
captured_at: "2026-11-28T09:51:53Z"
capture_tool: "hn-digest"
hn_id: 50675972
score: 1
comments: 0
posted_at: "2026-11-28T09:07:04Z"
tags:
  - hacker-news
---

# Refusals and Fallback of Anthropic Models

- HN: [50675972](https://news.ycombinator.com/item?id=50675972)
- Source: [platform.claude.com](https://platform.claude.com/docs/en/build-with-claude/refusals-and-fallback)
- Score: 1
- Comments: 0
- Posted: 2026-11-28T09:07:04Z

## Translation

Title: Refusals and Fallback of Anthropic Models
Article title: Refusals and fallback - Claude Platform Docs
Description: How Claude Fable models, Claude Opus models, Claude Sonnet 5.5, and Claude Haiku 5.5 return classifier refusals and how to retry refused requests on a fallback model.

Article text:
Refusals and fallback - Claude Platform Docs Claude Platform Docs Messages
How Claude Fable models, Claude Opus models, Claude Sonnet 5.5, and Claude Haiku 5.5 return classifier refusals and how to retry refused requests on a fallback model.
 Limited support for other models; see details
Claude Fable 5.2, Claude Fable 5.1, Claude Fable 5, Claude Opus 5.5, Claude Opus 5, Claude Sonnet 5.5, and Claude Haiku 5.5 include safety classifiers that can decline a request. When that happens, you receive a normal response, not an error, with stop_reason: "refusal" . Its stop_details.category names the policy area (see What a refusal looks like ). You can usually still get an answer by sending the same request to another Claude model. This page shows you how to recognize a refusal and how to set up that retry.
Read this page when you build on any of these models and want declined requests to fall through to another model automatically. It also applies when you have seen "refusal" in a response and want to know what to do next.
Stop reasons and fallback : the full list of stop_reason values
Fallback credit : how to avoid paying the prompt-cache cost twice when you build the retry yourself
SDK middleware : the SDK helper that wraps all of this
Fallback and billing recipe : a worked end-to-end example
The simplest setup, in beta on the Claude API and on Claude Models on Google Cloud (Anthropic managed) : set fallbacks to "default" , and the API retries a declined request on the fallback model Anthropic recommends for its refusal category. For categories with no recommended fallback, the refusal stands. Claude Haiku 5.5 has no server-side fallback, so for it use client-side fallback with the SDK middleware or write the retry yourself .
client = Anthropic()
response = client.beta.messages.create(
model = "claude-fable-5" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Hello, Claude" }],
fallbacks = "default" ,
betas = [ "server-side-fallback-2026-07-01" ],
)
print (response.model)
The following sections cover what a refusal response contains, when to use server-side or client-side fallback, and how each is billed.
A refusal is a successful HTTP 200 response with stop_reason: "refusal" :
{
"id" : "msg_01XFUDYJgAACzvnptvVoYEL" ,
"type" : "message" ,
"role" : "assistant" ,
"model" : "claude-fable-5" ,
"content" : [],
"stop_reason" : "refusal" ,
"stop_details" : {
"type" : "refusal" ,
"category" : "cyber" ,
"explanation" : "This request was declined because it could enable cyber harm."
},
"usage" : {
"input_tokens" : 412 ,
"output_tokens" : 0
}
} 
The stop_details object explains the decline:
category : names the policy area that triggered the classifier.
explanation : a human-readable description. The text is not stable, so display it rather than parse it.
recommended_model : present only on requests that set fallbacks ( server-side fallback , beta). It names a model to retry directly when the API skipped the fallback attempt (for example, the fallback model was rate limited), and is null otherwise. It’s a hint, not a guarantee.
category and explanation are both null when the refusal does not map to a named category. That null is a normal, permanent value, not a placeholder.
stop_details itself is null for every stop reason other than refusal .
A refusal can arrive before any output, or mid-stream after partial output. In either case, treat any partial output as incomplete and discard it.
Keep reasoning in thinking blocks
A reasoning_extraction refusal usually comes from a prompt that asks the model to put its thinking or reasoning in the output, either verbatim or in a fixed format. Common examples:
A <thinking> , <reasoning> , or scratchpad section that the model fills in before it answers
A reasoning , thinking , or trace field in JSON output or in a tool input
The model’s private notes, or a running log of its reasoning
Reasoning asked for verbatim or in full
The wording can sit in a system prompt, a skill, or a tool description, so check those too. Earlier assistant turns and examples in the prompt that show reasoning in tags, such as <thinking> , might also trigger this refusal.
These models might return this category:
Claude Fable 5.2, Claude Fable 5.1, and Claude Fable 5
Claude Mythos 5.2, Claude Mythos 5.1, and Claude Mythos 5
Claude Opus 5.5 and Claude Opus 5
You can still ask Claude to explain its answer. Ask for a short explanation, the evidence behind a result, or a summary of the actions it took. To see a summary of the model’s reasoning, set thinking: {"type": "adaptive", "display": "summarized"} and read the thinking blocks . This category has no recommended fallback model, so change the prompt rather than retry the request. If a reworded request is still declined, contact support and include the request ID .
These billing rules apply on every platform.
Refusals before any output: To disrupt attempts to circumvent Anthropic’s safeguards at scale, a refusal that arrives before any output is billed when its stop_details.category is "bio" , "frontier_llm" , or "reasoning_extraction" . These are the categories where Anthropic measures low volumes of false positives, as of September 2026. These refusals are billed like any other request, at the rates of the model that ran it. A refusal before any output in any other category, or with a null category, is not billed. Either way, content is empty and token counts appear in usage . The request still counts against your rate limits.
Mid-stream refusals: A mid-stream refusal bills the input tokens and the output already streamed at normal rates.
Fallback: When you use fallback, the refusal that triggered it is billed, in addition to the fallback request, when it arrived mid-stream or is in one of the billed categories. Fallback credit compensates for the fallback request’s prompt-cache miss, so you don’t pay to cache the conversation twice. A Claude Haiku 5.5 refusal carries no fallback credit, so a fallback after one pays the full cost of writing the fallback model’s prompt cache. For how server-side fallback reports each attempt, see Billing and rate limits .
The billed categories may change as Anthropic keeps measuring and refining its safeguards’ false positive rates. The Billed before any output column in the refusal category table lists the billed categories.
There are three ways to retry a refused request on another model. The right one depends on where you are running and how much control you need.
Server-side fallback and the SDK middleware apply fallback credit for you. You only need the Fallback credit page when you build the retry yourself. Claude Haiku 5.5 has neither server-side fallback nor fallback credit, so use the SDK middleware or a manual retry.
Server-side fallback retries a refused request inside a single API call. In the default mode, when the primary model declines and the refusal category has a recommended fallback, the API runs the same request on the model Anthropic recommends for that category. You can instead name up to three fallback models of your own . Either way, you get back one response that names the model that answered, so your user gets an answer in one round trip.
Server-side fallback is in beta on the Claude API and on Claude Models on Google Cloud (Anthropic managed) . On platforms without server-side fallback, use client-side fallback with the SDK middleware instead. The fallbacks parameter is not supported on the Message Batches API (a batch item that includes it comes back as an errored result). Claude Haiku 5.5 has no server-side fallback: with fallbacks: "default" , a declined request stays declined, and a list of fallback models returns a 400 error.
Set the fallbacks parameter to the string "default" and send the server-side-fallback-2026-07-01 beta header. The API then applies the requested model’s server-defined default routing, which selects a recommended fallback model based on the refusal category the classifier reports, so refused requests are served without you maintaining a model list as recommendations change.
Default routing never draws the up-front oversized-image rejection for models you did not choose: a routed model that would resize an image marked "oversized_image": "error" is dropped from the routing instead, so a marked image is never served resized.
client = Anthropic()
response = client.beta.messages.create(
model = "claude-fable-5" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Hello, Claude" }],
fallbacks = "default" ,
betas = [ "server-side-fallback-2026-07-01" ],
)
# A fallback_message entry in usage.iterations means a fallback model ran;
# pair it with stop_reason to confirm the fallback served the response.
fallback_ran = any (
iteration.type == "fallback_message"
for iteration in response.usage.iterations or []
)
served_by_fallback = fallback_ran and response.stop_reason != "refusal"
print (
json.dumps(
{
"stop_reason" : response.stop_reason,
"model" : response.model,
"served_by_fallback" : served_by_fallback,
}
)
)
Anthropic sets safeguards for each model individually and for each policy category, in line with the model’s capability: depending on the category, a flagged request might fall back to a less capable model or be declined. The "default" mode encodes these per-model, per-category recommendations for you, so a refused request is retried on the model Anthropic recommends for that category. Fallbacks are visible either way: the response names the model that served it, and the fallback content block marks the handoff.
The routing is applied server-side and is not published for each model on the Models API Preview with links to the API reference available — press Tab until focus enters the card, Escape to return. . To see which model served a refused request, check the response’s top-level model field and look for a fallback_message entry in usage.iterations , as this page’s samples do.
Only a safety classifier decline triggers the fallback. A rate limit, overload, or server error on the requested model is returned to you as-is.
The beta header must carry exactly the date 2026-07-01 , which supports both "default" and the explicit-list form, or 2026-06-01 , which accepts only the explicit-list form. Under any other server-side-fallback-* value, the fallbacks parameter is rejected with a 400 error. If you built against an earlier preview of this feature, update the beta header and the request and response shapes together to the ones on this page.
Naming your own fallback models
Instead of default routing, you can set fallbacks to a list of up to three models. When the requested model declines, the API runs the next model in the chain on the same request. Use this form when you want to control exactly which models serve refused requests, such as pinning a model your application has qualified.
Named fallback models count toward the oversized-image check . A request whose image block sets "oversized_image": "error" is checked up front against the requested model and every named fallback. The request is rejected if any of them would resize that image, and the rejection’s reported rescale target fits them all.
The highlighted lines are the only difference from the default-routing request.
client = Anthropic()
response = client.beta.messages.create(
model = "claude-fable-5" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Hello, Claude" }],
fallbacks = [{ "model" : "claude-opus-4-8" }],
betas = [ "server-side-fallback-2026-07-01" ],
)
print (response.model)
A few rules apply to the fallbacks list:
Entries are tried in order. Each must be distinct from the other entries and from the requested model.
Each entry must be one of the requested model’s permitted targets. With the beta header set, that list is published as allowed_fallback_models on the model’s entry in the Models API Preview with links to the API reference available — press Tab until focus enters the card, Escape to return. .
Each entry names a

[truncated]

## Original Extract

How Claude Fable models, Claude Opus models, Claude Sonnet 5.5, and Claude Haiku 5.5 return classifier refusals and how to retry refused requests on a fallback model.

Refusals and fallback - Claude Platform Docs Claude Platform Docs Messages
How Claude Fable models, Claude Opus models, Claude Sonnet 5.5, and Claude Haiku 5.5 return classifier refusals and how to retry refused requests on a fallback model.
 Limited support for other models; see details
Claude Fable 5.2, Claude Fable 5.1, Claude Fable 5, Claude Opus 5.5, Claude Opus 5, Claude Sonnet 5.5, and Claude Haiku 5.5 include safety classifiers that can decline a request. When that happens, you receive a normal response, not an error, with stop_reason: "refusal" . Its stop_details.category names the policy area (see What a refusal looks like ). You can usually still get an answer by sending the same request to another Claude model. This page shows you how to recognize a refusal and how to set up that retry.
Read this page when you build on any of these models and want declined requests to fall through to another model automatically. It also applies when you have seen "refusal" in a response and want to know what to do next.
Stop reasons and fallback : the full list of stop_reason values
Fallback credit : how to avoid paying the prompt-cache cost twice when you build the retry yourself
SDK middleware : the SDK helper that wraps all of this
Fallback and billing recipe : a worked end-to-end example
The simplest setup, in beta on the Claude API and on Claude Models on Google Cloud (Anthropic managed) : set fallbacks to "default" , and the API retries a declined request on the fallback model Anthropic recommends for its refusal category. For categories with no recommended fallback, the refusal stands. Claude Haiku 5.5 has no server-side fallback, so for it use client-side fallback with the SDK middleware or write the retry yourself .
client = Anthropic()
response = client.beta.messages.create(
model = "claude-fable-5" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Hello, Claude" }],
fallbacks = "default" ,
betas = [ "server-side-fallback-2026-07-01" ],
)
print (response.model)
The following sections cover what a refusal response contains, when to use server-side or client-side fallback, and how each is billed.
A refusal is a successful HTTP 200 response with stop_reason: "refusal" :
{
"id" : "msg_01XFUDYJgAACzvnptvVoYEL" ,
"type" : "message" ,
"role" : "assistant" ,
"model" : "claude-fable-5" ,
"content" : [],
"stop_reason" : "refusal" ,
"stop_details" : {
"type" : "refusal" ,
"category" : "cyber" ,
"explanation" : "This request was declined because it could enable cyber harm."
},
"usage" : {
"input_tokens" : 412 ,
"output_tokens" : 0
}
} 
The stop_details object explains the decline:
category : names the policy area that triggered the classifier.
explanation : a human-readable description. The text is not stable, so display it rather than parse it.
recommended_model : present only on requests that set fallbacks ( server-side fallback , beta). It names a model to retry directly when the API skipped the fallback attempt (for example, the fallback model was rate limited), and is null otherwise. It’s a hint, not a guarantee.
category and explanation are both null when the refusal does not map to a named category. That null is a normal, permanent value, not a placeholder.
stop_details itself is null for every stop reason other than refusal .
A refusal can arrive before any output, or mid-stream after partial output. In either case, treat any partial output as incomplete and discard it.
Keep reasoning in thinking blocks
A reasoning_extraction refusal usually comes from a prompt that asks the model to put its thinking or reasoning in the output, either verbatim or in a fixed format. Common examples:
A <thinking> , <reasoning> , or scratchpad section that the model fills in before it answers
A reasoning , thinking , or trace field in JSON output or in a tool input
The model’s private notes, or a running log of its reasoning
Reasoning asked for verbatim or in full
The wording can sit in a system prompt, a skill, or a tool description, so check those too. Earlier assistant turns and examples in the prompt that show reasoning in tags, such as <thinking> , might also trigger this refusal.
These models might return this category:
Claude Fable 5.2, Claude Fable 5.1, and Claude Fable 5
Claude Mythos 5.2, Claude Mythos 5.1, and Claude Mythos 5
Claude Opus 5.5 and Claude Opus 5
You can still ask Claude to explain its answer. Ask for a short explanation, the evidence behind a result, or a summary of the actions it took. To see a summary of the model’s reasoning, set thinking: {"type": "adaptive", "display": "summarized"} and read the thinking blocks . This category has no recommended fallback model, so change the prompt rather than retry the request. If a reworded request is still declined, contact support and include the request ID .
These billing rules apply on every platform.
Refusals before any output: To disrupt attempts to circumvent Anthropic’s safeguards at scale, a refusal that arrives before any output is billed when its stop_details.category is "bio" , "frontier_llm" , or "reasoning_extraction" . These are the categories where Anthropic measures low volumes of false positives, as of September 2026. These refusals are billed like any other request, at the rates of the model that ran it. A refusal before any output in any other category, or with a null category, is not billed. Either way, content is empty and token counts appear in usage . The request still counts against your rate limits.
Mid-stream refusals: A mid-stream refusal bills the input tokens and the output already streamed at normal rates.
Fallback: When you use fallback, the refusal that triggered it is billed, in addition to the fallback request, when it arrived mid-stream or is in one of the billed categories. Fallback credit compensates for the fallback request’s prompt-cache miss, so you don’t pay to cache the conversation twice. A Claude Haiku 5.5 refusal carries no fallback credit, so a fallback after one pays the full cost of writing the fallback model’s prompt cache. For how server-side fallback reports each attempt, see Billing and rate limits .
The billed categories may change as Anthropic keeps measuring and refining its safeguards’ false positive rates. The Billed before any output column in the refusal category table lists the billed categories.
There are three ways to retry a refused request on another model. The right one depends on where you are running and how much control you need.
Server-side fallback and the SDK middleware apply fallback credit for you. You only need the Fallback credit page when you build the retry yourself. Claude Haiku 5.5 has neither server-side fallback nor fallback credit, so use the SDK middleware or a manual retry.
Server-side fallback retries a refused request inside a single API call. In the default mode, when the primary model declines and the refusal category has a recommended fallback, the API runs the same request on the model Anthropic recommends for that category. You can instead name up to three fallback models of your own . Either way, you get back one response that names the model that answered, so your user gets an answer in one round trip.
Server-side fallback is in beta on the Claude API and on Claude Models on Google Cloud (Anthropic managed) . On platforms without server-side fallback, use client-side fallback with the SDK middleware instead. The fallbacks parameter is not supported on the Message Batches API (a batch item that includes it comes back as an errored result). Claude Haiku 5.5 has no server-side fallback: with fallbacks: "default" , a declined request stays declined, and a list of fallback models returns a 400 error.
Set the fallbacks parameter to the string "default" and send the server-side-fallback-2026-07-01 beta header. The API then applies the requested model’s server-defined default routing, which selects a recommended fallback model based on the refusal category the classifier reports, so refused requests are served without you maintaining a model list as recommendations change.
Default routing never draws the up-front oversized-image rejection for models you did not choose: a routed model that would resize an image marked "oversized_image": "error" is dropped from the routing instead, so a marked image is never served resized.
client = Anthropic()
response = client.beta.messages.create(
model = "claude-fable-5" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Hello, Claude" }],
fallbacks = "default" ,
betas = [ "server-side-fallback-2026-07-01" ],
)
# A fallback_message entry in usage.iterations means a fallback model ran;
# pair it with stop_reason to confirm the fallback served the response.
fallback_ran = any (
iteration.type == "fallback_message"
for iteration in response.usage.iterations or []
)
served_by_fallback = fallback_ran and response.stop_reason != "refusal"
print (
json.dumps(
{
"stop_reason" : response.stop_reason,
"model" : response.model,
"served_by_fallback" : served_by_fallback,
}
)
)
Anthropic sets safeguards for each model individually and for each policy category, in line with the model’s capability: depending on the category, a flagged request might fall back to a less capable model or be declined. The "default" mode encodes these per-model, per-category recommendations for you, so a refused request is retried on the model Anthropic recommends for that category. Fallbacks are visible either way: the response names the model that served it, and the fallback content block marks the handoff.
The routing is applied server-side and is not published for each model on the Models API Preview with links to the API reference available — press Tab until focus enters the card, Escape to return. . To see which model served a refused request, check the response’s top-level model field and look for a fallback_message entry in usage.iterations , as this page’s samples do.
Only a safety classifier decline triggers the fallback. A rate limit, overload, or server error on the requested model is returned to you as-is.
The beta header must carry exactly the date 2026-07-01 , which supports both "default" and the explicit-list form, or 2026-06-01 , which accepts only the explicit-list form. Under any other server-side-fallback-* value, the fallbacks parameter is rejected with a 400 error. If you built against an earlier preview of this feature, update the beta header and the request and response shapes together to the ones on this page.
Naming your own fallback models
Instead of default routing, you can set fallbacks to a list of up to three models. When the requested model declines, the API runs the next model in the chain on the same request. Use this form when you want to control exactly which models serve refused requests, such as pinning a model your application has qualified.
Named fallback models count toward the oversized-image check . A request whose image block sets "oversized_image": "error" is checked up front against the requested model and every named fallback. The request is rejected if any of them would resize that image, and the rejection’s reported rescale target fits them all.
The highlighted lines are the only difference from the default-routing request.
client = Anthropic()
response = client.beta.messages.create(
model = "claude-fable-5" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Hello, Claude" }],
fallbacks = [{ "model" : "claude-opus-4-8" }],
betas = [ "server-side-fallback-2026-07-01" ],
)
print (response.model)
A few rules apply to the fallbacks list:
Entries are tried in order. Each must be distinct from the other entries and from the requested model.
Each entry must be one of the requested model’s permitted targets. With the beta header set, that list is published as allowed_fallback_models on the model’s entry in the Models API Preview with links to the API reference available — press Tab until focus enters the card, Escape to return. .
Each entry names a

[truncated]
