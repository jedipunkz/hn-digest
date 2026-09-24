---
source: "https://performance.dev/chatgpt"
hn_url: "https://news.ycombinator.com/item?id=49837278"
title: "Reverse Engineering ChatGPT Web: How OpenAI Built for a Billion Users"
article_title: "Reverse Engineering ChatGPT Web:\nHow OpenAI Built for a Billion Users"
image: "https://media.performance.dev/cdn-cgi/image/width=1200,height=630,quality=100,fit=cover,format=auto/posts/p_IE23cQZdfRUJ/ZVpQoFJq3Ooo.jpg"
author: "fagnerbrack"
captured_at: "2026-09-24T22:01:20Z"
capture_tool: "hn-digest"
hn_id: 49837278
score: 1
comments: 0
posted_at: "2026-09-24T22:00:31Z"
tags:
  - hacker-news
---

# Reverse Engineering ChatGPT Web: How OpenAI Built for a Billion Users

- HN: [49837278](https://news.ycombinator.com/item?id=49837278)
- Source: [performance.dev](https://performance.dev/chatgpt)
- Score: 1
- Comments: 0
- Posted: 2026-09-24T22:00:31Z

## Translation

Title: Reverse Engineering ChatGPT Web: How OpenAI Built for a Billion Users
Article title: Reverse Engineering ChatGPT Web:
How OpenAI Built for a Billion Users
Description: A technical deep dive into how OpenAI built ChatGPT's web app: React Router 7, streaming SSR, code splitting, and feature flags at billion-user scale.

Article text:
Reverse Engineering ChatGPT Web:
How OpenAI Built for a Billion Users
Open a new tab, type chatgpt.com , and ask it something. No account, no sign in, no spinner. The page is interactive immediately, and the answer starts streaming a moment after you hit enter. That experience is served to roughly 1 billion people, making chatgpt.com a top website on the entire internet and one of the most used web apps in the world. On the surface it seems like a simple interface, but after digging in it's anything but. I spent days reverse engineering their web app by digging through the page source, bundled code, and network requests to understand how it was built.
The migration from Next.js to React Router
Don't reinvent the wheel with components
Every answer is a render problem
The fastest path to the first token
Letting a billion strangers in
The difference between ChatGPT and Claude
Before we start: I don't work at OpenAI and I've never seen their source code. What makes this piece different from my Linear and Conductor breakdowns is that OpenAI hasn't really published anything about ChatGPT web. This piece is a mixture of history, investigation, and technical insights. The whole story has been assembled from chatgpt.com itself (the HTML, JavaScript, CSS, and network requests) plus random talks, tweets, and their career page.
The two most important things I consider when architecting a web app are the goal and the constraints. These two factors shape every downstream decision. Who is the customer? How important is the initial load? Do you need an account to use it? Does it need to show up in search engines? What's the budget? How important is the developer experience? Is there existing architecture? And so on...
From the outside, the goal of ChatGPT feels clear to me: be the site the entire world uses to interact with AI. It's incredibly important that the site is accessible to anyone around the world on any device. It should cater to everyone from an Android user on a spotty connection who just discovered ChatGPT, to a power user with a Pro subscription on a MacBook Pro.
Having those constraints in place eliminates entire architectures for you. A standard client side rendered app is out, because you'd be staring at a blank page while a bunch of JavaScript downloads before anything meaningful shows up. Apps like Linear get away with CSR because their users log in once and live inside the app all day, but that approach doesn't fit OpenAI's goal.
Server-side rendering is a better approach in this case, but it doesn't come free. Every request now requires your server to fetch data, render HTML, and send it back, making it more expensive to run than serving static assets. As a developer, it's also a much harder model to reason about. Your app now exists in two environments (the server and the browser) and you're constantly thinking about which APIs are available where. You'll eventually hit hydration mismatches, accidentally reference window on the server, or spend time debugging caching and rendering behavior. It's a more complex architecture with more failure modes. But if your goal is the fastest possible first load, real HTML for search engines, and a great experience for first-time visitors, those are often tradeoffs worth making.
Let's see what ChatGPT does. The first step is to take a look under the hood to better understand the architecture and technical decisions they've made. From what I can gather, this is roughly the stack as of today:
Frontend
React 19 + react-dom (UI runtime, streamed server render)
React Router 7 (framework mode) (routing, loaders, streaming SSR)
TypeScript (language)
TanStack Query (server state on the client)
Tailwind CSS (utility styling + design tokens)
Radix UI primitives (menus, popovers, selects, toasts)
ProseMirror (the composer you type into)
CodeMirror 6 (answer code blocks + canvas editor)
Motion (animation)
silk-hq (native-feeling sheets on mobile)
KaTeX (math; fonts load on demand)
Mapbox GL (maps inside answers)
System font stack (no webfont for UI text)
Build and delivery
Vite (bundling; per-route JS + CSS chunks)
Cloudflare (CDN, cache, bot defense)
First-party assets (everything from chatgpt.com/cdn/assets)
Experimentation and observability
Statsig (flags + experiments, server evaluated)
Datadog RUM (real user monitoring)
Realtime
SSE over fetch (token streaming)
LiveKit + WebRTC (voice mode)
WebAssembly (audio processing, syntax grammars) What stands out to me the most is that they're mostly using off-the-shelf libraries instead of custom frameworks or components. If you've ever looked at Gemini's source code it's a very different beast full of proprietary Google stuff. ChatGPT feels simple and standard, which are incredibly important details when you want to build an app that scales to a billion users.
The migration from Next.js to React Router
The history of ChatGPT is very interesting to me. Before diving into the technical aspects of how it's built I wanted to understand the origin, key events, and path they took to get to where they are today.
ChatGPT launched on November 30, 2022 as a Next.js 12 app using the Pages Router (my personal favorite version of Next.js to this day). Inside OpenAI, ChatGPT was a “research preview” shipped with low expectations , a quick wrapper around a fine-tuned GPT-3.5 meant to collect feedback from the public. We all know how that went.
One of my favorite artifacts from the Next.js era is the route manifest, still sitting in the Wayback Machine. You can see how simple the initial web app was, plus hints of internal codespace and workspace tools that appear to have shipped inside the same build, likely a byproduct of racing the MVP out the door.
// _buildManifest.js from the December 2022 launch build
sortedPages : [
"/" ,
"/_app" ,
"/_error" ,
"/auth/error" ,
"/auth/login" ,
"/chat" ,
"/codespace" ,
"/error" ,
"/workspace/[[...tid]]"
] Through the entire Next.js App Router era that followed, chatgpt.com never adopted it. They rode the Pages Router for ~21 months, then migrated off Next.js entirely. In my past projects I've also followed the same path.
The migration was caught in the wild, not announced by ChatGPT. On August 27, 2024, Tibor Blaho noticed chatgpt.com serving a Remix build to a slice of users as an experiment. By September 4 it had rolled out widely, and Ryan Florence, Remix's co-creator, tweeted “New Remix app just dropped: chatgpt.com.” Wes Bos dug through the bundle on YouTube the next day.
What Wes found is the interesting part, because it explains the philosophy. The Remix-era app ran real server infrastructure (an Express server executing route loaders and serializing around 7,000 lines of JSON into window.__remixContext ) but rendered almost no actual UI on the server. A shell, some preload links, a theme script. No Remix actions anywhere. Mutations went through their own API instead of the library's built-in actions. Remix was a router, a data pipe, and a hydration shell around what was essentially a client side app. Ryan Florence later confirmed ChatGPT as a Remix app where “most of their data loading is done with TanStack Query instead of React Router's loaders.” Evan You, who created Vite, summarized the community's reading of the whole move in one line: “Many of you probably just need an SPA too.”
So far, the story goes as an initial MVP with Next.js Pages Router, then migrated to Remix in a SPA-like configuration. But there's more!
When Remix v2 merged back into React Router as v7 in November 2024, ChatGPT followed. Today you can inspect the page source of every page and see:
window.__reactRouterContext = {
"basename" : "/" ,
"ssr" : true ,
"isSpaMode" : false ,
"routeDiscovery" : {
"mode" : "lazy" ,
"manifestPath" : "/__manifest"
} ,
// ...
} ; ssr: true . This is React Router 7 framework mode, the full Vite-based, streaming, server rendered setup. And here's the evolution I find fascinating: the Remix era served an empty shell, but the app today server renders the entire logged-out experience as real HTML. They went even heavier on the SSR.
One more thing I found unique is hiding in the router. The client manifest lists 354 routes, and only about 13 of them are the actual chat app. The rest are marketing and landing pages, all living in the same codebase, same router, same deploy. The route names read straight out of React Router's file conventions:
routes/_conversation._index the chat
routes/_conversation.c.$conversationId a conversation
routes/($lang).codex.pricing marketing
routes/($lang).business._index marketing
routes/($lang).atlas.get-started marketing Most companies tend to split marketing into a separate site (Linear runs theirs on Next.js, away from the app). As far as I can tell, OpenAI ships theirs inside the product app, which means landing pages share the design system, the flag system, and the router. Clicking from a campaign page into the chat is a client side navigation, not a fresh page load, which is pretty cool.
The logged-out document I measured is served as 84 KB compressed, and it arrives with a time to first byte of around 50 to 65ms to my laptop in Vancouver, served by a nearby Cloudflare edge. Inside that single response is everything needed to paint a working app shell: roughly 30 KB of real markup (the sidebar, the “What's on the agenda today?” greeting, the composer) plus the styles to render it. After hydration the entire page is 548 DOM nodes (very light). The experience is entirely focused on simplicity and one goal: the chat input.
But the markup/DOM is the least interesting part of the document. The head is where the performance work lives. Before any bundle is requested, inline scripts run. The first one is the standard theme selection:
// inlined in <head>, runs before first paint
!function (){ try {
var d = document.documentElement, c = d.classList;
c. remove ( 'light' , 'dark' ) ;
var e = localStorage. getItem ( 'theme' ) ;
if ( 'system' === e || ( ! e && true )){
var t = '(prefers-color-scheme: dark)' , m = window. matchMedia ( t ) ;
if ( m.media !== t || m.matches ){ d.style.colorScheme = 'dark' ; c. add ( 'dark' ) }
else { d.style.colorScheme = 'light' ; c. add ( 'light' ) }
} else if ( e ){ c. add ( e || '' ) }
if ( e === 'light' || e === 'dark' ) d.style.colorScheme = e
} catch ( e ){}}() ; Sound familiar? It's the exact pattern I covered in the Linear article : read the user's saved theme out of localStorage and apply it to the html element before paint, so there's never a flash of the wrong theme.
Right after it comes one of my favorite details that hints at ChatGPT's goal:
window.__oai_logHTML ? window. __oai_logHTML ()
: window.__oai_SSR_HTML = window.__oai_SSR_HTML || Date. now () ;
requestAnimationFrame ( function (){
window.__oai_logTTI ? window. __oai_logTTI ()
: window.__oai_SSR_TTI = window.__oai_SSR_TTI || Date. now ()
}) ; They timestamp the moment the HTML starts executing and the first frame after it, for every single user, and feed it into their real user monitoring. Performance is clearly a priority: they're measuring the entire load path, from first byte to interactive. You can't improve what you don't measure!
The document response itself streams. This is React 19 streaming SSR through React Router: the server flushes the shell immediately and Suspense boundaries fill in as data resolves, patched into place by tiny inline scripts. Route loader data arrives over its own ReadableStream in parallel with the HTML. And tucked inside that loader data is a set of server decisions about what the client should warm up: shouldPrefetchModels , shouldPrefetchHistory , shouldPrefetchStarredConversations . In chatgpt.com's case, the server dictates a lot of the client's behavior. It hands the client a per-user prefetch plan. It really does feel like the client is a thin wrapper that is driven by the server.
When you think about it, it's a great approach for this usecase. If you believe in generative UI this is the first baby

[truncated]

## Original Extract

A technical deep dive into how OpenAI built ChatGPT's web app: React Router 7, streaming SSR, code splitting, and feature flags at billion-user scale.

Reverse Engineering ChatGPT Web:
How OpenAI Built for a Billion Users
Open a new tab, type chatgpt.com , and ask it something. No account, no sign in, no spinner. The page is interactive immediately, and the answer starts streaming a moment after you hit enter. That experience is served to roughly 1 billion people, making chatgpt.com a top website on the entire internet and one of the most used web apps in the world. On the surface it seems like a simple interface, but after digging in it's anything but. I spent days reverse engineering their web app by digging through the page source, bundled code, and network requests to understand how it was built.
The migration from Next.js to React Router
Don't reinvent the wheel with components
Every answer is a render problem
The fastest path to the first token
Letting a billion strangers in
The difference between ChatGPT and Claude
Before we start: I don't work at OpenAI and I've never seen their source code. What makes this piece different from my Linear and Conductor breakdowns is that OpenAI hasn't really published anything about ChatGPT web. This piece is a mixture of history, investigation, and technical insights. The whole story has been assembled from chatgpt.com itself (the HTML, JavaScript, CSS, and network requests) plus random talks, tweets, and their career page.
The two most important things I consider when architecting a web app are the goal and the constraints. These two factors shape every downstream decision. Who is the customer? How important is the initial load? Do you need an account to use it? Does it need to show up in search engines? What's the budget? How important is the developer experience? Is there existing architecture? And so on...
From the outside, the goal of ChatGPT feels clear to me: be the site the entire world uses to interact with AI. It's incredibly important that the site is accessible to anyone around the world on any device. It should cater to everyone from an Android user on a spotty connection who just discovered ChatGPT, to a power user with a Pro subscription on a MacBook Pro.
Having those constraints in place eliminates entire architectures for you. A standard client side rendered app is out, because you'd be staring at a blank page while a bunch of JavaScript downloads before anything meaningful shows up. Apps like Linear get away with CSR because their users log in once and live inside the app all day, but that approach doesn't fit OpenAI's goal.
Server-side rendering is a better approach in this case, but it doesn't come free. Every request now requires your server to fetch data, render HTML, and send it back, making it more expensive to run than serving static assets. As a developer, it's also a much harder model to reason about. Your app now exists in two environments (the server and the browser) and you're constantly thinking about which APIs are available where. You'll eventually hit hydration mismatches, accidentally reference window on the server, or spend time debugging caching and rendering behavior. It's a more complex architecture with more failure modes. But if your goal is the fastest possible first load, real HTML for search engines, and a great experience for first-time visitors, those are often tradeoffs worth making.
Let's see what ChatGPT does. The first step is to take a look under the hood to better understand the architecture and technical decisions they've made. From what I can gather, this is roughly the stack as of today:
Frontend
React 19 + react-dom (UI runtime, streamed server render)
React Router 7 (framework mode) (routing, loaders, streaming SSR)
TypeScript (language)
TanStack Query (server state on the client)
Tailwind CSS (utility styling + design tokens)
Radix UI primitives (menus, popovers, selects, toasts)
ProseMirror (the composer you type into)
CodeMirror 6 (answer code blocks + canvas editor)
Motion (animation)
silk-hq (native-feeling sheets on mobile)
KaTeX (math; fonts load on demand)
Mapbox GL (maps inside answers)
System font stack (no webfont for UI text)
Build and delivery
Vite (bundling; per-route JS + CSS chunks)
Cloudflare (CDN, cache, bot defense)
First-party assets (everything from chatgpt.com/cdn/assets)
Experimentation and observability
Statsig (flags + experiments, server evaluated)
Datadog RUM (real user monitoring)
Realtime
SSE over fetch (token streaming)
LiveKit + WebRTC (voice mode)
WebAssembly (audio processing, syntax grammars) What stands out to me the most is that they're mostly using off-the-shelf libraries instead of custom frameworks or components. If you've ever looked at Gemini's source code it's a very different beast full of proprietary Google stuff. ChatGPT feels simple and standard, which are incredibly important details when you want to build an app that scales to a billion users.
The migration from Next.js to React Router
The history of ChatGPT is very interesting to me. Before diving into the technical aspects of how it's built I wanted to understand the origin, key events, and path they took to get to where they are today.
ChatGPT launched on November 30, 2022 as a Next.js 12 app using the Pages Router (my personal favorite version of Next.js to this day). Inside OpenAI, ChatGPT was a “research preview” shipped with low expectations , a quick wrapper around a fine-tuned GPT-3.5 meant to collect feedback from the public. We all know how that went.
One of my favorite artifacts from the Next.js era is the route manifest, still sitting in the Wayback Machine. You can see how simple the initial web app was, plus hints of internal codespace and workspace tools that appear to have shipped inside the same build, likely a byproduct of racing the MVP out the door.
// _buildManifest.js from the December 2022 launch build
sortedPages : [
"/" ,
"/_app" ,
"/_error" ,
"/auth/error" ,
"/auth/login" ,
"/chat" ,
"/codespace" ,
"/error" ,
"/workspace/[[...tid]]"
] Through the entire Next.js App Router era that followed, chatgpt.com never adopted it. They rode the Pages Router for ~21 months, then migrated off Next.js entirely. In my past projects I've also followed the same path.
The migration was caught in the wild, not announced by ChatGPT. On August 27, 2024, Tibor Blaho noticed chatgpt.com serving a Remix build to a slice of users as an experiment. By September 4 it had rolled out widely, and Ryan Florence, Remix's co-creator, tweeted “New Remix app just dropped: chatgpt.com.” Wes Bos dug through the bundle on YouTube the next day.
What Wes found is the interesting part, because it explains the philosophy. The Remix-era app ran real server infrastructure (an Express server executing route loaders and serializing around 7,000 lines of JSON into window.__remixContext ) but rendered almost no actual UI on the server. A shell, some preload links, a theme script. No Remix actions anywhere. Mutations went through their own API instead of the library's built-in actions. Remix was a router, a data pipe, and a hydration shell around what was essentially a client side app. Ryan Florence later confirmed ChatGPT as a Remix app where “most of their data loading is done with TanStack Query instead of React Router's loaders.” Evan You, who created Vite, summarized the community's reading of the whole move in one line: “Many of you probably just need an SPA too.”
So far, the story goes as an initial MVP with Next.js Pages Router, then migrated to Remix in a SPA-like configuration. But there's more!
When Remix v2 merged back into React Router as v7 in November 2024, ChatGPT followed. Today you can inspect the page source of every page and see:
window.__reactRouterContext = {
"basename" : "/" ,
"ssr" : true ,
"isSpaMode" : false ,
"routeDiscovery" : {
"mode" : "lazy" ,
"manifestPath" : "/__manifest"
} ,
// ...
} ; ssr: true . This is React Router 7 framework mode, the full Vite-based, streaming, server rendered setup. And here's the evolution I find fascinating: the Remix era served an empty shell, but the app today server renders the entire logged-out experience as real HTML. They went even heavier on the SSR.
One more thing I found unique is hiding in the router. The client manifest lists 354 routes, and only about 13 of them are the actual chat app. The rest are marketing and landing pages, all living in the same codebase, same router, same deploy. The route names read straight out of React Router's file conventions:
routes/_conversation._index the chat
routes/_conversation.c.$conversationId a conversation
routes/($lang).codex.pricing marketing
routes/($lang).business._index marketing
routes/($lang).atlas.get-started marketing Most companies tend to split marketing into a separate site (Linear runs theirs on Next.js, away from the app). As far as I can tell, OpenAI ships theirs inside the product app, which means landing pages share the design system, the flag system, and the router. Clicking from a campaign page into the chat is a client side navigation, not a fresh page load, which is pretty cool.
The logged-out document I measured is served as 84 KB compressed, and it arrives with a time to first byte of around 50 to 65ms to my laptop in Vancouver, served by a nearby Cloudflare edge. Inside that single response is everything needed to paint a working app shell: roughly 30 KB of real markup (the sidebar, the “What's on the agenda today?” greeting, the composer) plus the styles to render it. After hydration the entire page is 548 DOM nodes (very light). The experience is entirely focused on simplicity and one goal: the chat input.
But the markup/DOM is the least interesting part of the document. The head is where the performance work lives. Before any bundle is requested, inline scripts run. The first one is the standard theme selection:
// inlined in <head>, runs before first paint
!function (){ try {
var d = document.documentElement, c = d.classList;
c. remove ( 'light' , 'dark' ) ;
var e = localStorage. getItem ( 'theme' ) ;
if ( 'system' === e || ( ! e && true )){
var t = '(prefers-color-scheme: dark)' , m = window. matchMedia ( t ) ;
if ( m.media !== t || m.matches ){ d.style.colorScheme = 'dark' ; c. add ( 'dark' ) }
else { d.style.colorScheme = 'light' ; c. add ( 'light' ) }
} else if ( e ){ c. add ( e || '' ) }
if ( e === 'light' || e === 'dark' ) d.style.colorScheme = e
} catch ( e ){}}() ; Sound familiar? It's the exact pattern I covered in the Linear article : read the user's saved theme out of localStorage and apply it to the html element before paint, so there's never a flash of the wrong theme.
Right after it comes one of my favorite details that hints at ChatGPT's goal:
window.__oai_logHTML ? window. __oai_logHTML ()
: window.__oai_SSR_HTML = window.__oai_SSR_HTML || Date. now () ;
requestAnimationFrame ( function (){
window.__oai_logTTI ? window. __oai_logTTI ()
: window.__oai_SSR_TTI = window.__oai_SSR_TTI || Date. now ()
}) ; They timestamp the moment the HTML starts executing and the first frame after it, for every single user, and feed it into their real user monitoring. Performance is clearly a priority: they're measuring the entire load path, from first byte to interactive. You can't improve what you don't measure!
The document response itself streams. This is React 19 streaming SSR through React Router: the server flushes the shell immediately and Suspense boundaries fill in as data resolves, patched into place by tiny inline scripts. Route loader data arrives over its own ReadableStream in parallel with the HTML. And tucked inside that loader data is a set of server decisions about what the client should warm up: shouldPrefetchModels , shouldPrefetchHistory , shouldPrefetchStarredConversations . In chatgpt.com's case, the server dictates a lot of the client's behavior. It hands the client a per-user prefetch plan. It really does feel like the client is a thin wrapper that is driven by the server.
When you think about it, it's a great approach for this usecase. If you believe in generative UI this is the first baby

[truncated]
