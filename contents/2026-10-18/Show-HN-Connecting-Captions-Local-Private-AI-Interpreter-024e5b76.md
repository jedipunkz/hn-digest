---
source: "https://local-host.ai/connectingCaptions/"
hn_url: "https://news.ycombinator.com/item?id=50138839"
title: "Show HN: Connecting Captions Local Private AI Interpreter"
article_title: "Connecting Captions · local-host.ai"
image: "https://local-host.ai/images/fluidsubtitles.jpg"
author: "chrisswimlee"
captured_at: "2026-10-18T08:36:23Z"
capture_tool: "hn-digest"
hn_id: 50138839
score: 1
comments: 0
posted_at: "2026-10-18T07:41:31Z"
tags:
  - hacker-news
---

# Show HN: Connecting Captions Local Private AI Interpreter

- HN: [50138839](https://news.ycombinator.com/item?id=50138839)
- Source: [local-host.ai](https://local-host.ai/connectingCaptions/)
- Score: 1
- Comments: 0
- Posted: 2026-10-18T07:41:31Z

## Translation

Title: Show HN: Connecting Captions Local Private AI Interpreter
Article title: Connecting Captions · local-host.ai
Description: On-device live captions on a Mac. Speech and translation stay on the machine. For presentations, bilingual rooms, and any conversation that needs a translation.
HN text: Hi HN, I built Connecting Captions, a live captioning and translation overlay for macOS that runs entirely on-device. Nothing goes over the network: no API calls, no audio leaving your machine, and it keeps working offline. I wanted live translation I could use in meetings and conversations without worrying about connectivity, latency spikes, or sending sensitive audio to a third party. The technically interesting part is a low-latency, self-correcting pipeline for speech-to-text and translation. It has two modes: Sentence mode: buffers until a sentence boundary, then translates with full context. It’s more accurate, especially between languages with different word order.
Streaming mode: translates word by word as speech arrives, trading grammatical accuracy for minimal delay. Current state: Platform: macOS
Models: Whisper
Latency: ~200 ms
Languages: 25 Code: https://github.com/chrisswimlee/connectingCaptions I’d love feedback, especially on the self-correction approach and which use cases matter most to people.

Article text:
local-host.ai
Connecting Captions
For work
Chris Swim Lee
Connecting Captions
Live captions on a Mac. Each sentence appears when it is ready. Use it
for a presentation, a bilingual room, or any conversation that needs a
translation.
Completely on-device by default. Speech recognition and Apple Translation
run on this Mac. Voice never leaves the machine unless someone turns on a
cloud speech model. Theater doesn’t call home. No account. No telemetry.
Arabic · Chinese · Danish · Dutch · English · Finnish · French · German · Hebrew · Hindi · Indonesian · Italian · Japanese · Korean · Malay · Norwegian · Polish · Portuguese · Russian · Spanish · Swedish · Thai · Turkish · Ukrainian · Vietnamese
Presentations.
Theater sits over the slides. A whole-screen share includes it. Share the slides window when the room should not see the captions.
Bilingual rooms.
Speak any of the 25 languages. The other side sees a caption or a
translation.
Any conversation.
Put the Mac in the room. Each sentence appears when it is ready.
Voice stays on the machine.
macOS 15 or later. Theater streaming needs Apple Silicon. The zip is notarized.
Unzip Connecting-Captions.zip and drag Connecting Captions to Applications.
Open Theater, allow the microphone, then press Listen.
A Theater window over slides. A whole-screen Zoom or Meet share includes it. Share the slides window when viewers should not see the captions.
On-device Apple Translation. Speech recognition comes from FluidVoice.
Voice stays on this Mac. A cloud speech model is something you turn on.
No account, no telemetry, no captions uploaded. Works offline once models are on disk.
macOS 15 or later. Theater streaming needs Apple Silicon.
Microphone for Listen. Accessibility only if you type into another app.
Personal use stays free. A named license is $60 per Mac per year, five
Macs minimum. IT can see the prices without emailing for a quote.
How does this monetize?
Named license $60 per Mac per year, five Macs minimum. Written SLA
+$90 per Mac per year. MDM pkg $1,500 once. Transcripts, voiceprints,
and training data are not part of the sale.
Does voice leave this Mac?
No, by default. Speech and translation run on the device. Voice stays
on the Mac unless someone turns on a cloud speech model. Theater
doesn’t send anything home. No account, no telemetry.
Who fixes a break?
Security reports get a reply within 7 days. A paid license can add a
written SLA.
A project from
local-host.ai
by
Chris Swim Lee

## Original Extract

On-device live captions on a Mac. Speech and translation stay on the machine. For presentations, bilingual rooms, and any conversation that needs a translation.

Hi HN, I built Connecting Captions, a live captioning and translation overlay for macOS that runs entirely on-device. Nothing goes over the network: no API calls, no audio leaving your machine, and it keeps working offline. I wanted live translation I could use in meetings and conversations without worrying about connectivity, latency spikes, or sending sensitive audio to a third party. The technically interesting part is a low-latency, self-correcting pipeline for speech-to-text and translation. It has two modes: Sentence mode: buffers until a sentence boundary, then translates with full context. It’s more accurate, especially between languages with different word order.
Streaming mode: translates word by word as speech arrives, trading grammatical accuracy for minimal delay. Current state: Platform: macOS
Models: Whisper
Latency: ~200 ms
Languages: 25 Code: https://github.com/chrisswimlee/connectingCaptions I’d love feedback, especially on the self-correction approach and which use cases matter most to people.

local-host.ai
Connecting Captions
For work
Chris Swim Lee
Connecting Captions
Live captions on a Mac. Each sentence appears when it is ready. Use it
for a presentation, a bilingual room, or any conversation that needs a
translation.
Completely on-device by default. Speech recognition and Apple Translation
run on this Mac. Voice never leaves the machine unless someone turns on a
cloud speech model. Theater doesn’t call home. No account. No telemetry.
Arabic · Chinese · Danish · Dutch · English · Finnish · French · German · Hebrew · Hindi · Indonesian · Italian · Japanese · Korean · Malay · Norwegian · Polish · Portuguese · Russian · Spanish · Swedish · Thai · Turkish · Ukrainian · Vietnamese
Presentations.
Theater sits over the slides. A whole-screen share includes it. Share the slides window when the room should not see the captions.
Bilingual rooms.
Speak any of the 25 languages. The other side sees a caption or a
translation.
Any conversation.
Put the Mac in the room. Each sentence appears when it is ready.
Voice stays on the machine.
macOS 15 or later. Theater streaming needs Apple Silicon. The zip is notarized.
Unzip Connecting-Captions.zip and drag Connecting Captions to Applications.
Open Theater, allow the microphone, then press Listen.
A Theater window over slides. A whole-screen Zoom or Meet share includes it. Share the slides window when viewers should not see the captions.
On-device Apple Translation. Speech recognition comes from FluidVoice.
Voice stays on this Mac. A cloud speech model is something you turn on.
No account, no telemetry, no captions uploaded. Works offline once models are on disk.
macOS 15 or later. Theater streaming needs Apple Silicon.
Microphone for Listen. Accessibility only if you type into another app.
Personal use stays free. A named license is $60 per Mac per year, five
Macs minimum. IT can see the prices without emailing for a quote.
How does this monetize?
Named license $60 per Mac per year, five Macs minimum. Written SLA
+$90 per Mac per year. MDM pkg $1,500 once. Transcripts, voiceprints,
and training data are not part of the sale.
Does voice leave this Mac?
No, by default. Speech and translation run on the device. Voice stays
on the Mac unless someone turns on a cloud speech model. Theater
doesn’t send anything home. No account, no telemetry.
Who fixes a break?
Security reports get a reply within 7 days. A paid license can add a
written SLA.
A project from
local-host.ai
by
Chris Swim Lee
