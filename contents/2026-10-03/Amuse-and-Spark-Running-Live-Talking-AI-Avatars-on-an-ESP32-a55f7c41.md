---
source: "https://geastack.com/blog-introducing-amuse-and-spark"
hn_url: "https://news.ycombinator.com/item?id=49944877"
title: "Amuse and Spark – Running Live Talking AI Avatars on an ESP32"
article_title: "Introducing Amuse and Spark — GEA Blog"
image: "https://geastack.com/og-cover.jpg"
author: "arbayi"
captured_at: "2026-10-03T15:03:38Z"
capture_tool: "hn-digest"
hn_id: 49944877
score: 1
comments: 0
posted_at: "2026-10-03T14:58:02Z"
tags:
  - hacker-news
---

# Amuse and Spark – Running Live Talking AI Avatars on an ESP32

- HN: [49944877](https://news.ycombinator.com/item?id=49944877)
- Source: [geastack.com](https://geastack.com/blog-introducing-amuse-and-spark)
- Score: 1
- Comments: 0
- Posted: 2026-10-03T14:58:02Z

## Translation

Title: Amuse and Spark – Running Live Talking AI Avatars on an ESP32
Article title: Introducing Amuse and Spark — GEA Blog
Description: Two open-source apps you can talk to on an ESP32-S3 watch board: Spark, a voice agent that connects directly to the OpenAI Realtime API, and Amuse, five LemonSlice video characters. Both are written in TypeScript, JSX and CSS and compiled with GeaStack.

Article text:
Skip to content
Executive summary
Examples
Docs
Blog
Services
Contact
Executive summary
Examples
Docs
Blog
Services
Contact
← Blog
Release
Introducing Amuse and Spark
Two open-source apps you can talk to, running on a few-dollar ESP32-S3 watch board. Spark is a voice
agent that connects straight to OpenAI's Realtime API. Amuse puts five animated video characters on
the screen. Both are written in TypeScript, JSX and CSS and compiled to native firmware with GeaStack.
Both apps run on the
Waveshare ESP32-S3 Touch AMOLED 2.06 ,
the watch-shaped board from our first post , and Amuse
also runs on its 1.8″ sibling. You swipe to pick a character, tap Start , and talk.
The source for both is on GitHub:
Spark is a voice agent. The board opens a WebSocket to OpenAI's Realtime API
( gpt-realtime-2.1 ) itself, streams microphone audio up and plays the reply as it
arrives. There is no server or phone in between. Audio is full duplex, so you can interrupt a reply by
speaking, the way you would with a person.
It has four characters, each with its own voice, persona and animation drawn on a Canvas:
KITT , a protective, dry-witted car, with a red scanner while it listens and three columns of voice lamps while it speaks.
Nova , an astronomer on an orbital observatory, with twinkling stars and planetary orbits.
Echo , a deep-sea navigator, with a sonar sweep and a speech waveform.
Flora , a mindfulness companion, with a growing leaf and a flower.
A character is one folder: a character.ts with its voice and instructions, a
View.tsx component, a stylesheet and an optional Canvas animation. Add a folder and
the build picks it up; there is nothing else to register. The conversation runs in a Worker, which
GeaStack compiles to native code alongside the UI.
Amuse puts a talking face on the screen. Its cast is Zuck, an awkward tech nerd; Alfred, a thoughtful
scholar; Felipe, an imaginative artist; Jojo, a lively optimist; and Todd, an easygoing frog. Each
is a LemonSlice video avatar with an
ElevenLabs voice. When you tap Start, the board creates a hosted
session, and the character's face, lip-synced to its voice, plays on the AMOLED as you talk.
Amuse needs a small bridge running on a computer on the same network. The bridge joins the hosted
session's call and relays the microphone, audio and video between it and the board. Several boards
can connect to one bridge at once, each with its own session. On the board, echo cancellation runs
natively, so the microphone stays open while the character speaks without hearing itself.
Swiping to a different character ends the current conversation. Each character keeps its
configuration and portraits in its own folder, prepared for both supported screen sizes.
Each repository's README covers setup. In short: npm ci , npx gea setup to
register the board and install ESP-IDF, fill in .env with your Wi-Fi and API
credentials, then build and flash with npm.
npm ci
npx gea setup
npm run build
npm run flash
Some things to know before you start:
Your API keys go into the firmware. Keep .env and any binaries you
build private.
Amuse needs the bridge (Python 3.11+) and LemonSlice credentials to hold a
conversation.
Both are meant to be forked. Change a persona, add a character, or point them at a different board
with npx gea setup . If you build something with them, we'd like to see it.
Building a device or a service on GeaStack? We'd like to hear about it — write to
contact@geastack.com .
TypeScript apps for microcontrollers, desktop, mobile, and consoles.

## Original Extract

Two open-source apps you can talk to on an ESP32-S3 watch board: Spark, a voice agent that connects directly to the OpenAI Realtime API, and Amuse, five LemonSlice video characters. Both are written in TypeScript, JSX and CSS and compiled with GeaStack.

Skip to content
Executive summary
Examples
Docs
Blog
Services
Contact
Executive summary
Examples
Docs
Blog
Services
Contact
← Blog
Release
Introducing Amuse and Spark
Two open-source apps you can talk to, running on a few-dollar ESP32-S3 watch board. Spark is a voice
agent that connects straight to OpenAI's Realtime API. Amuse puts five animated video characters on
the screen. Both are written in TypeScript, JSX and CSS and compiled to native firmware with GeaStack.
Both apps run on the
Waveshare ESP32-S3 Touch AMOLED 2.06 ,
the watch-shaped board from our first post , and Amuse
also runs on its 1.8″ sibling. You swipe to pick a character, tap Start , and talk.
The source for both is on GitHub:
Spark is a voice agent. The board opens a WebSocket to OpenAI's Realtime API
( gpt-realtime-2.1 ) itself, streams microphone audio up and plays the reply as it
arrives. There is no server or phone in between. Audio is full duplex, so you can interrupt a reply by
speaking, the way you would with a person.
It has four characters, each with its own voice, persona and animation drawn on a Canvas:
KITT , a protective, dry-witted car, with a red scanner while it listens and three columns of voice lamps while it speaks.
Nova , an astronomer on an orbital observatory, with twinkling stars and planetary orbits.
Echo , a deep-sea navigator, with a sonar sweep and a speech waveform.
Flora , a mindfulness companion, with a growing leaf and a flower.
A character is one folder: a character.ts with its voice and instructions, a
View.tsx component, a stylesheet and an optional Canvas animation. Add a folder and
the build picks it up; there is nothing else to register. The conversation runs in a Worker, which
GeaStack compiles to native code alongside the UI.
Amuse puts a talking face on the screen. Its cast is Zuck, an awkward tech nerd; Alfred, a thoughtful
scholar; Felipe, an imaginative artist; Jojo, a lively optimist; and Todd, an easygoing frog. Each
is a LemonSlice video avatar with an
ElevenLabs voice. When you tap Start, the board creates a hosted
session, and the character's face, lip-synced to its voice, plays on the AMOLED as you talk.
Amuse needs a small bridge running on a computer on the same network. The bridge joins the hosted
session's call and relays the microphone, audio and video between it and the board. Several boards
can connect to one bridge at once, each with its own session. On the board, echo cancellation runs
natively, so the microphone stays open while the character speaks without hearing itself.
Swiping to a different character ends the current conversation. Each character keeps its
configuration and portraits in its own folder, prepared for both supported screen sizes.
Each repository's README covers setup. In short: npm ci , npx gea setup to
register the board and install ESP-IDF, fill in .env with your Wi-Fi and API
credentials, then build and flash with npm.
npm ci
npx gea setup
npm run build
npm run flash
Some things to know before you start:
Your API keys go into the firmware. Keep .env and any binaries you
build private.
Amuse needs the bridge (Python 3.11+) and LemonSlice credentials to hold a
conversation.
Both are meant to be forked. Change a persona, add a character, or point them at a different board
with npx gea setup . If you build something with them, we'd like to see it.
Building a device or a service on GeaStack? We'd like to hear about it — write to
contact@geastack.com .
TypeScript apps for microcontrollers, desktop, mobile, and consoles.
