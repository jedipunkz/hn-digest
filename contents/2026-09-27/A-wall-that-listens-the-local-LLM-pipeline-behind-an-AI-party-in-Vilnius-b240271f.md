---
source: "https://vania-novikau.me/ai-party-wall/"
hn_url: "https://news.ycombinator.com/item?id=49866154"
title: "A wall that listens: the local-LLM pipeline behind an AI party in Vilnius"
article_title: "A wall that listens: the local-LLM pipeline behind an AI party in Vilnius - Vania Novikau"
image: ""
author: "vanadiuz"
captured_at: "2026-09-27T12:48:56Z"
capture_tool: "hn-digest"
hn_id: 49866154
score: 1
comments: 0
posted_at: "2026-09-27T12:35:51Z"
tags:
  - hacker-news
---

# A wall that listens: the local-LLM pipeline behind an AI party in Vilnius

- HN: [49866154](https://news.ycombinator.com/item?id=49866154)
- Source: [vania-novikau.me](https://vania-novikau.me/ai-party-wall/)
- Score: 1
- Comments: 0
- Posted: 2026-09-27T12:35:51Z

## Translation

Title: A wall that listens: the local-LLM pipeline behind an AI party in Vilnius
Article title: A wall that listens: the local-LLM pipeline behind an AI party in Vilnius - Vania Novikau
Description: This is a 78-second first test of the 1st Dynamic AI Party, being built at the ZOLAK Micro Art Residence in Vilnius. The idea is simple to say and fiddly to build: the wall of the room is a projection, and what is on it is decided, live, by a language model that is watching

Article text:
A wall that listens: the local-LLM pipeline behind an AI party in Vilnius - Vania Novikau
Skip to content
Skip to content
Menu
Blog
A wall that listens: the local-LLM pipeline behind an AI party in Vilnius
This is a 78-second first test of the 1st Dynamic AI Party , being built at the ZOLAK Micro Art Residence in Vilnius. The idea is simple to say and fiddly to build: the wall of the room is a projection, and what is on it is decided, live, by a language model that is watching and listening to the people in front of it.
Watch on YouTube: https://www.youtube.com/shorts/7qU4OlB3CtU
An AI-driven light wall that reacts to the room. You move, it sees; you speak, it listens; then it answers on the wall. In the footage, words people say in the room (“basically”, “crazy”, “great”) come back as large projected type a moment later, and someone tests whether raising their hands makes the wall brighter. The on-screen label is honest about the state of it: real output, first test, unedited , version “AI-LOOP v0.1”.
Everything runs on one local rig, and each piece of hardware has one job:
Input — CCTV cameras and Shure SM58 microphones cover the whole room, live.
Eyes — a Coral Edge TPU (4 TOPS) runs person and motion detection on every camera feed: who is there and who is moving.
Ears — an RTX ADA 2000 (16 GB VRAM) does live speech-to-text on every mic.
Brain — an RTX 3090 (24 GB VRAM) runs a local LLM that takes what was seen and heard and decides what the wall says and shows next.
Hands — a 32-core CPU decodes all the streams, composites the next frame and feeds the projector.
Output — a laser projector puts it on the wall, where you see it instantly.
That makes it a closed loop rather than a visualiser: you react, the wall reacts, you react again. The interesting engineering is in the split. Detection lives on a tiny accelerator, transcription on a mid-range card, and the big card is left free for the model that makes the decision, so no single stage has to do everything at party latency.
Most “AI art” at events is pre-rendered or prompted by one operator at a laptop. Here the prompt is the room itself: the crowd’s movement and speech are the input, and the model runs locally on hardware in the building, which is what makes a live loop with every camera and every mic practical at all.
ZOLAK Micro Art Residence, Vytenio g. 50D, Vilnius, 3rd floor (enter via the old ELFA factory, Panerių g. 37). The date is still to be announced.
Physicist · Researcher · AI Explorer
Curiosity-driven research at the intersection of Physics and AI.
© 1995 – 2026 All Rights Reserved.
Made with
[vc_single_image image=”7100″ img_size=”full” css=””]

## Original Extract

This is a 78-second first test of the 1st Dynamic AI Party, being built at the ZOLAK Micro Art Residence in Vilnius. The idea is simple to say and fiddly to build: the wall of the room is a projection, and what is on it is decided, live, by a language model that is watching

A wall that listens: the local-LLM pipeline behind an AI party in Vilnius - Vania Novikau
Skip to content
Skip to content
Menu
Blog
A wall that listens: the local-LLM pipeline behind an AI party in Vilnius
This is a 78-second first test of the 1st Dynamic AI Party , being built at the ZOLAK Micro Art Residence in Vilnius. The idea is simple to say and fiddly to build: the wall of the room is a projection, and what is on it is decided, live, by a language model that is watching and listening to the people in front of it.
Watch on YouTube: https://www.youtube.com/shorts/7qU4OlB3CtU
An AI-driven light wall that reacts to the room. You move, it sees; you speak, it listens; then it answers on the wall. In the footage, words people say in the room (“basically”, “crazy”, “great”) come back as large projected type a moment later, and someone tests whether raising their hands makes the wall brighter. The on-screen label is honest about the state of it: real output, first test, unedited , version “AI-LOOP v0.1”.
Everything runs on one local rig, and each piece of hardware has one job:
Input — CCTV cameras and Shure SM58 microphones cover the whole room, live.
Eyes — a Coral Edge TPU (4 TOPS) runs person and motion detection on every camera feed: who is there and who is moving.
Ears — an RTX ADA 2000 (16 GB VRAM) does live speech-to-text on every mic.
Brain — an RTX 3090 (24 GB VRAM) runs a local LLM that takes what was seen and heard and decides what the wall says and shows next.
Hands — a 32-core CPU decodes all the streams, composites the next frame and feeds the projector.
Output — a laser projector puts it on the wall, where you see it instantly.
That makes it a closed loop rather than a visualiser: you react, the wall reacts, you react again. The interesting engineering is in the split. Detection lives on a tiny accelerator, transcription on a mid-range card, and the big card is left free for the model that makes the decision, so no single stage has to do everything at party latency.
Most “AI art” at events is pre-rendered or prompted by one operator at a laptop. Here the prompt is the room itself: the crowd’s movement and speech are the input, and the model runs locally on hardware in the building, which is what makes a live loop with every camera and every mic practical at all.
ZOLAK Micro Art Residence, Vytenio g. 50D, Vilnius, 3rd floor (enter via the old ELFA factory, Panerių g. 37). The date is still to be announced.
Physicist · Researcher · AI Explorer
Curiosity-driven research at the intersection of Physics and AI.
© 1995 – 2026 All Rights Reserved.
Made with
[vc_single_image image=”7100″ img_size=”full” css=””]
