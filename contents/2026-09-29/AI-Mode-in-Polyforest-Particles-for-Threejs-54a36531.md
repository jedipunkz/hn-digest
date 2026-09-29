---
source: "https://blog.polyforest.com/announcing-ai-mode/"
hn_url: "https://news.ycombinator.com/item?id=49896345"
title: "AI Mode in Polyforest Particles for Threejs"
article_title: "AI Mode in the Particle Editor — PolyForest Blog"
image: "https://blog.polyforest.com/social-card.png"
author: "nbilyk"
captured_at: "2026-09-29T17:46:02Z"
capture_tool: "hn-digest"
hn_id: 49896345
score: 1
comments: 0
posted_at: "2026-09-29T16:51:24Z"
tags:
  - hacker-news
---

# AI Mode in Polyforest Particles for Threejs

- HN: [49896345](https://news.ycombinator.com/item?id=49896345)
- Source: [blog.polyforest.com](https://blog.polyforest.com/announcing-ai-mode/)
- Score: 1
- Comments: 0
- Posted: 2026-09-29T16:51:24Z

## Translation

Title: AI Mode in Polyforest Particles for Threejs
Article title: AI Mode in the Particle Editor — PolyForest Blog
Description: The PolyForest Particle Editor now has an AI assistant that can build and edit effects with you.

Article text:
AI Mode in the Particle Editor — PolyForest Blog
Skip to content Posts About polyforest.com September 27, 2026
AI Mode in the Particle Editor
Nicholas Bilyk particles ai threejs
The PolyForest Particle Editor has a new AI mode! Open the chat, describe the
effect you want, and the assistant edits it while you watch the preview.
The PolyForest Particle editor is a web app I wrote many years ago that you’ve never heard of. Originally it was a tool
for a Kotlin multiplatform OpenGL/WebGL framework I wrote in 2016. While that’s been long abandoned, I recently ported
both the polyforest.com website and the editor to Three.js .
While I can honestly not promise that I’ll keep this editor maintained, as it’s just a hobby side project, the more
users it gets the more incentive I have to make it awesome.
In gen AI, creating effects by hand suddenly felt tedious and time-consuming. To make this easier, I created an AI
assistant that works with Anthropic or OpenAI keys. The editor uses
the three-particles runtime library. This library is
also a port from the Kotlin version of long past.
AI mode works with your own Anthropic or OpenAI API key. Add it under AI Assistant in the account menu. The key is
checked with the provider, then encrypted and stored on the server. It is never sent back to the browser, and you can
remove it at any time.
It knows what you’re looking at
The chat panel docks beside the preview, so the effect stays visible as it changes. The assistant sees the effect you
have open, so a request like “fade the flame to blue” applies to the flame in front of you.
The assistant works through the same operations as the editor:
create, duplicate, rename and delete effects
add and edit emitters and their timelines
draw new textures and add them to your texture library
Each step shows up in the chat as it happens.
Plug: the video playback engine in this blog is Amazon Vinyl
Screen recording background music by Suno
Blog written content by human hands, views expressed are my own
© 2026 PolyForest · RSS · polyforest.com · LinkedIn

## Original Extract

The PolyForest Particle Editor now has an AI assistant that can build and edit effects with you.

AI Mode in the Particle Editor — PolyForest Blog
Skip to content Posts About polyforest.com September 27, 2026
AI Mode in the Particle Editor
Nicholas Bilyk particles ai threejs
The PolyForest Particle Editor has a new AI mode! Open the chat, describe the
effect you want, and the assistant edits it while you watch the preview.
The PolyForest Particle editor is a web app I wrote many years ago that you’ve never heard of. Originally it was a tool
for a Kotlin multiplatform OpenGL/WebGL framework I wrote in 2016. While that’s been long abandoned, I recently ported
both the polyforest.com website and the editor to Three.js .
While I can honestly not promise that I’ll keep this editor maintained, as it’s just a hobby side project, the more
users it gets the more incentive I have to make it awesome.
In gen AI, creating effects by hand suddenly felt tedious and time-consuming. To make this easier, I created an AI
assistant that works with Anthropic or OpenAI keys. The editor uses
the three-particles runtime library. This library is
also a port from the Kotlin version of long past.
AI mode works with your own Anthropic or OpenAI API key. Add it under AI Assistant in the account menu. The key is
checked with the provider, then encrypted and stored on the server. It is never sent back to the browser, and you can
remove it at any time.
It knows what you’re looking at
The chat panel docks beside the preview, so the effect stays visible as it changes. The assistant sees the effect you
have open, so a request like “fade the flame to blue” applies to the flame in front of you.
The assistant works through the same operations as the editor:
create, duplicate, rename and delete effects
add and edit emitters and their timelines
draw new textures and add them to your texture library
Each step shows up in the chat as it happens.
Plug: the video playback engine in this blog is Amazon Vinyl
Screen recording background music by Suno
Blog written content by human hands, views expressed are my own
© 2026 PolyForest · RSS · polyforest.com · LinkedIn
