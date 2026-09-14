---
source: "https://stemjson.com/"
hn_url: "https://news.ycombinator.com/item?id=49695057"
title: "Show HN: StemJSON – a language for LLMs to extend native mobile apps on the fly"
article_title: "StemJSON - Mobile apps that build themselves"
image: "https://stemjson.com/og.png"
author: "vkrychun"
captured_at: "2026-09-14T11:58:49Z"
capture_tool: "hn-digest"
hn_id: 49695057
score: 2
comments: 0
posted_at: "2026-09-14T11:26:45Z"
tags:
  - hacker-news
---

# Show HN: StemJSON – a language for LLMs to extend native mobile apps on the fly

- HN: [49695057](https://news.ycombinator.com/item?id=49695057)
- Source: [stemjson.com](https://stemjson.com/)
- Score: 2
- Comments: 0
- Posted: 2026-09-14T11:26:45Z

## Translation

Title: Show HN: StemJSON – a language for LLMs to extend native mobile apps on the fly
Article title: StemJSON - Mobile apps that build themselves
Description: Open declarative JSON specification for AI-generated, server-driven, end-user-customizable native mobile apps.
HN text: Hey HN! I'd like to share something I've been working on for a while and hear what you think. It's a new domain-specific language built on JSON and designed for AI, along with an ecosystem for building native mobile features, screens, and even entire apps quickly, right on the devices where they live. On top of that, end users can use AI to customize what their apps do on the fly, which might take the user experience to a whole new level.

Article text:
StemJSON - Mobile apps that build themselves
StemJSON Home Specification Runtime
Products
StemStudio AI The mobile IDE for StemJSON MCP Server Mobile features from your AI chat Pluri Apps that build themselves About FAQ Home Specification Runtime Products StemStudio AI MCP Server Pluri About FAQ Mobile apps that build themselves
StemJSON is a new declarative language that turns descriptions into native apps - written by AI, pushed by your backend, or shaped by your users.
What it is The language AI uses to build native apps
StemJSON is a declarative language based on JSON, well-suited for AI to generate comprehensive, performant native app UI and logic - whether building new apps or extending existing ones.
The format is open. Drop it into AI tools, prototyping environments, IDE extensions, or alongside hand-written code to simplify and accelerate native app design, development, and delivery.
→ renders 9:41 ●●● ⌁ NATIVE Vasyl Krychun Creator of StemJSON Follow Use cases Where it fits
A sandboxed area inside your app where end users describe a feature in their own words - and the runtime renders a fully native module on the spot. Apps that adapt to each user, on demand.
An LLM emits StemJSON from the user's prompt; the runtime validates it, renders it natively, and keeps it inside the sandbox you defined. No new binary, no backend round-trip for the UI.
Push complete functional modules from your backend - state, actions, expressions, navigation, validation - straight to the device.
your server → StemJSON → device
From a prompt or a Figma export to a runnable native prototype. Tools emit StemJSON; the runtime renders it on-device - no glue code in between.
Figma / AI tool → StemJSON → runnable app
Drop StemJSON modules into a native app for the parts that change often - settings, onboarding, content feeds - and keep the rest hand-written.
StemJSON modules + native code → one app
The complete language, in one place.
Components, actions, expressions, state, navigation and the conformance rules both runtimes enforce - so you know exactly what a module can express before you write one.
Native rendering on iOS and Android.
Add the Swift or Kotlin SDK and render StemJSON inside an app you already ship: new screens from your backend without an app update, or a Stem AI Area where your users describe what they need and the runtime renders it natively, inside the sandbox you define. Free tier, no time limit.
Complete example apps you can run today.
Open-source apps with the runtime already integrated. Clone one, drop in your own .stem file, and reuse the patterns that fit what you're building.
Spotted something off in the spec? Open an issue on GitHub ↗ - feedback is the fastest way to make the next version better.

## Original Extract

Open declarative JSON specification for AI-generated, server-driven, end-user-customizable native mobile apps.

Hey HN! I'd like to share something I've been working on for a while and hear what you think. It's a new domain-specific language built on JSON and designed for AI, along with an ecosystem for building native mobile features, screens, and even entire apps quickly, right on the devices where they live. On top of that, end users can use AI to customize what their apps do on the fly, which might take the user experience to a whole new level.

StemJSON - Mobile apps that build themselves
StemJSON Home Specification Runtime
Products
StemStudio AI The mobile IDE for StemJSON MCP Server Mobile features from your AI chat Pluri Apps that build themselves About FAQ Home Specification Runtime Products StemStudio AI MCP Server Pluri About FAQ Mobile apps that build themselves
StemJSON is a new declarative language that turns descriptions into native apps - written by AI, pushed by your backend, or shaped by your users.
What it is The language AI uses to build native apps
StemJSON is a declarative language based on JSON, well-suited for AI to generate comprehensive, performant native app UI and logic - whether building new apps or extending existing ones.
The format is open. Drop it into AI tools, prototyping environments, IDE extensions, or alongside hand-written code to simplify and accelerate native app design, development, and delivery.
→ renders 9:41 ●●● ⌁ NATIVE Vasyl Krychun Creator of StemJSON Follow Use cases Where it fits
A sandboxed area inside your app where end users describe a feature in their own words - and the runtime renders a fully native module on the spot. Apps that adapt to each user, on demand.
An LLM emits StemJSON from the user's prompt; the runtime validates it, renders it natively, and keeps it inside the sandbox you defined. No new binary, no backend round-trip for the UI.
Push complete functional modules from your backend - state, actions, expressions, navigation, validation - straight to the device.
your server → StemJSON → device
From a prompt or a Figma export to a runnable native prototype. Tools emit StemJSON; the runtime renders it on-device - no glue code in between.
Figma / AI tool → StemJSON → runnable app
Drop StemJSON modules into a native app for the parts that change often - settings, onboarding, content feeds - and keep the rest hand-written.
StemJSON modules + native code → one app
The complete language, in one place.
Components, actions, expressions, state, navigation and the conformance rules both runtimes enforce - so you know exactly what a module can express before you write one.
Native rendering on iOS and Android.
Add the Swift or Kotlin SDK and render StemJSON inside an app you already ship: new screens from your backend without an app update, or a Stem AI Area where your users describe what they need and the runtime renders it natively, inside the sandbox you define. Free tier, no time limit.
Complete example apps you can run today.
Open-source apps with the runtime already integrated. Clone one, drop in your own .stem file, and reuse the patterns that fit what you're building.
Spotted something off in the spec? Open an issue on GitHub ↗ - feedback is the fastest way to make the next version better.
