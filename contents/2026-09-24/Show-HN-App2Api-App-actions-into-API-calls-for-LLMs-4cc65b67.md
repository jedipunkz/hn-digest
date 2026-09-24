---
source: "https://app2api.com/"
hn_url: "https://news.ycombinator.com/item?id=49829882"
title: "Show HN: App2Api – App actions into API calls for LLMs"
article_title: "App2Api — Give your AI the API behind your app"
image: "https://app2api.com/og.png"
author: "flkdnfsnsf"
captured_at: "2026-09-24T13:17:46Z"
capture_tool: "hn-digest"
hn_id: 49829882
score: 1
comments: 0
posted_at: "2026-09-24T12:50:37Z"
tags:
  - hacker-news
---

# Show HN: App2Api – App actions into API calls for LLMs

- HN: [49829882](https://news.ycombinator.com/item?id=49829882)
- Source: [app2api.com](https://app2api.com/)
- Score: 1
- Comments: 0
- Posted: 2026-09-24T12:50:37Z

## Translation

Title: Show HN: App2Api – App actions into API calls for LLMs
Article title: App2Api — Give your AI the API behind your app
Description: App2Api reverse-engineers the API behind a mobile app you own or are authorized to test — the endpoints, the auth chain, a replayable client — and keeps it working as the app updates. Free beta.
HN text: Please give me feedback

Article text:
app 2 api
How it works
What you get
Who it's for
FAQ
Get beta access
Free beta
Give your AI the API behind your app.
Describe what you want your app to do, in plain English. App2Api captures the real API call
behind it — login, tokens and all — and gives your AI agent an OpenAPI spec it can call.
Free during the beta. No card, no setup.
Goal-directed, not a blind crawl.
You don't get a dump of everything the app happened to touch. You describe one workflow, and App2Api walks it end to end — then hands you a call you can reproduce standalone.
Point it at the app + describe the goal
Send the app and a plain-English goal — "trigger the loyalty lookup," "add an item to my cart." Any credentials are typed locally, never sent to a model.
It drives the real app and captures the traffic
A vision planner drives a live emulator toward that goal, capturing every HTTPS call the app makes along the way.
It traces the call back to a replayable chain
When the action fires, we find the request behind it and walk backward for everything it needs — session bootstrap, tokens, attestation.
You get a runnable flow + OpenAPI spec
A runnable flow.sh , an OpenAPI 3.1 spec, and an owner report. Monitoring is coming in beta, so you'll know the moment an update breaks the chain.
Three files, not a screenshot.
Every run hands back the same set — a call to run, a spec to build on, a report to share.
The runnable call with its full auth chain. Paste it into a terminal and it works standalone.
An OpenAPI 3.1 spec for the endpoints the workflow touched — drop it straight into your tooling.
A shareable report of the captured requests, ready to hand to a teammate or a client.
Built for mobile, QA, and security teams.
Whether you build the app, test it, or ship on top of it, App2Api turns its API into something you can depend on.
Document your own app's real API surface, catch undeclared endpoints, and get a contract you can test against every release.
Turn a manual workflow into a scripted, replayable call — and know the instant an app update changes the contract underneath it.
Trace auth chains and attestation flows on your engagements, with an owner report you can hand straight to the client.
No. App2Api is for apps you own or are explicitly authorized to test. Every request includes an authorization confirmation, and we decline anything we can't reasonably believe is authorized.
A runnable flow.sh , an OpenAPI 3.1 spec, and an owner report — the endpoint you asked for plus the full auth chain needed to replay it standalone.
We email you the traced call once your run finishes. During the beta we run each request by hand, so depending on the queue it can take a little while — you don't need to resubmit.
Does it handle login and hardened apps?
Login flows are supported — credentials are typed locally at the keystroke and never sent to a model. Some hardened apps (for example ones shipping certain integrity protections) may not run in the emulator, and we'll tell you when that's the case.
Nothing during the beta. We run it on our own infrastructure and email you the result.
Tell us the workflow you want to reach.
Send us an app you're authorized to test and the goal you're after, and we'll email you the traced call once it's done. It's free during the beta — and depending on the queue, it can take a little while.
We'll run it and email you the traced call. Depending on the queue that can take a while, so no need to resubmit — just keep an eye on your inbox.
app 2 api
© 2026 App2Api · Free beta
App2Api is intended solely for use on applications you own or are explicitly authorized to
test. You are responsible for ensuring your use complies with the target service's terms,
applicable law, and any authorization you rely on. We decline requests we cannot reasonably
believe are authorized.

## Original Extract

App2Api reverse-engineers the API behind a mobile app you own or are authorized to test — the endpoints, the auth chain, a replayable client — and keeps it working as the app updates. Free beta.

Please give me feedback

app 2 api
How it works
What you get
Who it's for
FAQ
Get beta access
Free beta
Give your AI the API behind your app.
Describe what you want your app to do, in plain English. App2Api captures the real API call
behind it — login, tokens and all — and gives your AI agent an OpenAPI spec it can call.
Free during the beta. No card, no setup.
Goal-directed, not a blind crawl.
You don't get a dump of everything the app happened to touch. You describe one workflow, and App2Api walks it end to end — then hands you a call you can reproduce standalone.
Point it at the app + describe the goal
Send the app and a plain-English goal — "trigger the loyalty lookup," "add an item to my cart." Any credentials are typed locally, never sent to a model.
It drives the real app and captures the traffic
A vision planner drives a live emulator toward that goal, capturing every HTTPS call the app makes along the way.
It traces the call back to a replayable chain
When the action fires, we find the request behind it and walk backward for everything it needs — session bootstrap, tokens, attestation.
You get a runnable flow + OpenAPI spec
A runnable flow.sh , an OpenAPI 3.1 spec, and an owner report. Monitoring is coming in beta, so you'll know the moment an update breaks the chain.
Three files, not a screenshot.
Every run hands back the same set — a call to run, a spec to build on, a report to share.
The runnable call with its full auth chain. Paste it into a terminal and it works standalone.
An OpenAPI 3.1 spec for the endpoints the workflow touched — drop it straight into your tooling.
A shareable report of the captured requests, ready to hand to a teammate or a client.
Built for mobile, QA, and security teams.
Whether you build the app, test it, or ship on top of it, App2Api turns its API into something you can depend on.
Document your own app's real API surface, catch undeclared endpoints, and get a contract you can test against every release.
Turn a manual workflow into a scripted, replayable call — and know the instant an app update changes the contract underneath it.
Trace auth chains and attestation flows on your engagements, with an owner report you can hand straight to the client.
No. App2Api is for apps you own or are explicitly authorized to test. Every request includes an authorization confirmation, and we decline anything we can't reasonably believe is authorized.
A runnable flow.sh , an OpenAPI 3.1 spec, and an owner report — the endpoint you asked for plus the full auth chain needed to replay it standalone.
We email you the traced call once your run finishes. During the beta we run each request by hand, so depending on the queue it can take a little while — you don't need to resubmit.
Does it handle login and hardened apps?
Login flows are supported — credentials are typed locally at the keystroke and never sent to a model. Some hardened apps (for example ones shipping certain integrity protections) may not run in the emulator, and we'll tell you when that's the case.
Nothing during the beta. We run it on our own infrastructure and email you the result.
Tell us the workflow you want to reach.
Send us an app you're authorized to test and the goal you're after, and we'll email you the traced call once it's done. It's free during the beta — and depending on the queue, it can take a little while.
We'll run it and email you the traced call. Depending on the queue that can take a while, so no need to resubmit — just keep an eye on your inbox.
app 2 api
© 2026 App2Api · Free beta
App2Api is intended solely for use on applications you own or are explicitly authorized to
test. You are responsible for ensuring your use complies with the target service's terms,
applicable law, and any authorization you rely on. We decline requests we cannot reasonably
believe are authorized.
