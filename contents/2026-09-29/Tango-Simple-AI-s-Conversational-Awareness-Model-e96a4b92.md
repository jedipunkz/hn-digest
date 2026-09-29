---
source: "https://www.usesimple.ai/blog/tango-voice-ai-end-of-turn-detection"
hn_url: "https://news.ycombinator.com/item?id=49897047"
title: "Tango: Simple AI's Conversational Awareness Model"
article_title: "Introducing Tango: Simple AI’s Conversational Awareness Model | Blog"
image: "https://framerusercontent.com/images/EQznguRUkupnDWZhYcyNCtVdwk.png?width=3840&height=2560"
author: "catherynl"
captured_at: "2026-09-29T17:45:47Z"
capture_tool: "hn-digest"
hn_id: 49897047
score: 2
comments: 0
posted_at: "2026-09-29T17:29:58Z"
tags:
  - hacker-news
---

# Tango: Simple AI's Conversational Awareness Model

- HN: [49897047](https://news.ycombinator.com/item?id=49897047)
- Source: [www.usesimple.ai](https://www.usesimple.ai/blog/tango-voice-ai-end-of-turn-detection)
- Score: 2
- Comments: 0
- Posted: 2026-09-29T17:29:58Z

## Translation

Title: Tango: Simple AI's Conversational Awareness Model
Article title: Introducing Tango: Simple AI’s Conversational Awareness Model | Blog
Description: Simple AI launches Tango, an audio-native model that combines end-of-turn detection and sentiment analysis for faster, more natural voice AI conversations.

Article text:
Introducing Tango: Simple AI’s Conversational Awareness Model
Introducing Tango: Simple AI’s Conversational Awareness Model
Simple AI launches Tango, an audio-native model that combines end-of-turn detection and sentiment analysis for faster, more natural voice AI conversations.
Voice AI’s biggest challenge is timing. For a conversation to feel natural, an agent has to recognize the moment a customer is done speaking and craft their response quickly and accurately. Most models still struggle with this, either interrupting mid-sentence, or pausing too long to confirm their turn to speak. Customers are left with conversations that feel robotic rather than human. It’s the biggest obstacle standing between voice AI and mainstream adoption.
We tried all of the end-of-turn models on the market and couldn’t find one that actually worked well in production. In Simple AI fashion, we chose to train our own.
Today, we’re introducing Tango , the system behind Simple’s most natural voice agents yet. In production, its end-of-turn decision lands a median 300ms before a conventional voice-activity detector would even detect a pause — measured on real telephony audio, not just clean VoIP.
Most voice AI on the market uses a cascaded approach: transcribe what the customer says, generate a response from that transcript, convert it back to audio. Everything the agent needs has to survive being flattened into text first, and a lot gets lost in translation. The result is both slower and less natural: no matter how fast your agent is, it’s waiting on that transcript before it can say anything.
Tango skips that step for the decisions that matter most, reading end-of-turn and interruption directly from audio. Transcripts still get generated for every call, but they no longer sit in the critical path.
Benefits of Tango’s Audio Processing
Tango isn’t just faster than competitor end-of-turn models, it’s also more accurate. This is crucial, since false positives mean more interruptions and false negatives mean uncomfortable pauses. With improved end-of-turn, conversations flow more freely. On the public LiveKit benchmark, it’s the fastest of 11 systems once you require at least 90% detection, missing just 15 of 400 true turn endings versus 54 for Soniox and 197 for Deepgram Flux. Against Smart Turn specifically, it fires 3.7 to 4.7 times less often mid-speech at equal turn-end coverage.
Genuine interruptions and quick backchannel cues like “mm-hmm” can look identical in a transcript, but they call for opposite behavior: stop talking for one, keep going for the other. Tango tells them apart in real time, so it doesn’t cut off a customer who’s simply agreeing. That’s cut interruptions by 64% versus our previous agent in production, recognizing a real interruption about 100ms after the customer starts talking.
Type “thank you” and there’s no way to tell if it was said with a smile or through gritted teeth. Text strips out tone, prosody, and pacing: the exact signals that make an interaction feel human. Audio-native processing keeps that signal intact.
The old pipeline adds roughly one second of latency before your agent starts thinking. Tango skips it: audio comes in, decisions come out, with both agent and customer channels processed at once. By going straight to streaming audio models, we can make assessments live every 80 ms, and run inference in less than half that time. There’s no transcript to wait for, and no second pipeline to run.
What This Looks (and Sounds) Like
A transcript-based model sees silence and guesses. Tango hears the customer still thinking, and responds at just the right time.
Tango can tell a real interruption from a conversational nod, so your agents don’t stop talking every time a customer says “okay”.
What This Means For Your Call Center
In practice, the difference between a natural-feeling agent and a robotic one comes down to about one second of latency. This gap shows up directly in the metrics you’re already tracking.
Customer Satisfaction Score (CSAT): A customer’s satisfaction is shaped by their overall experience rather than resolution alone. Fewer awkward pauses and false interruptions means fewer calls that feel broken, even when the agent is correct.
Average Handle Time (AHT): Every pause, interruption, or repeated “sorry, go ahead” adds seconds to your call. At call center volume, these seconds compound fast.
Abandonment: The customers you lose aren’t the ones who got a wrong answer. They’re the ones who hung up in frustration before you got the chance to help them at all.
A system near the top of this chart catches more end-of-turns, while a system to the left catches them sooner. Tango is the most accurate model for its speed, and the fastest model for its accuracy. Because the best, most natural conversations have both.
Built for the Calls You Actually Get
Most end-of-turn models are trained and benchmarked on clean, high-bandwidth audio. Real customer calls are not. They’re compressed, 8kHz telephony audio, running through phone systems built way before there were AI agents on the other end of the line. That’s why models that perform well in a demo often degrade the moment they meet a call center’s actual call volume.
Tango is built dual-channel and telephony-native from the ground up. It plugs directly into the inbound and outbound calling systems you already run, without asking you to upgrade your infrastructure first. That’s the difference between a benchmark number and a number you’ll actually see in production.
Tango is already live across Simple’s production traffic, delivering lower latency, higher accuracy, and more natural conversations for the calls our customers handle every day. It’s already shaping how we think about every agent we build from here.
They say it takes two to Tango. Schedule a demo today to see how Tango could sound for you.
Put your best rep on every call.
How is Tango different from a speech-to-text pipeline?
')"> Tango analyzes the audio directly instead of waiting for speech to become text and then turning a response back into audio. That preserves cues such as tone, pacing, and prosody while removing roughly one second of transcription latency from the response path. Transcripts are still generated after each call for review and reporting.
Why does Tango handle turn-taking and sentiment in one model?
')"> A caller’s tone and rhythm help reveal both how they feel and whether they have finished speaking. Tango processes those signals together in one pass. During development, Simple found that joint training improved both tasks and made the system simpler and cheaper to train and deploy.
Do I need to upgrade my phone system to use Tango?
')"> No. Tango handles compressed 8 kHz telephony audio as well as cleaner 24 kHz VoIP audio, and it connects to existing inbound and outbound calling systems.
Revenue Agent vs. Deflection Bot: How to Tell Them Apart
Build It on Vapi or ElevenLabs, or Buy a Platform?
How Omaha Steaks Cut Call Abandonment From 16% to 3%

## Original Extract

Simple AI launches Tango, an audio-native model that combines end-of-turn detection and sentiment analysis for faster, more natural voice AI conversations.

Introducing Tango: Simple AI’s Conversational Awareness Model
Introducing Tango: Simple AI’s Conversational Awareness Model
Simple AI launches Tango, an audio-native model that combines end-of-turn detection and sentiment analysis for faster, more natural voice AI conversations.
Voice AI’s biggest challenge is timing. For a conversation to feel natural, an agent has to recognize the moment a customer is done speaking and craft their response quickly and accurately. Most models still struggle with this, either interrupting mid-sentence, or pausing too long to confirm their turn to speak. Customers are left with conversations that feel robotic rather than human. It’s the biggest obstacle standing between voice AI and mainstream adoption.
We tried all of the end-of-turn models on the market and couldn’t find one that actually worked well in production. In Simple AI fashion, we chose to train our own.
Today, we’re introducing Tango , the system behind Simple’s most natural voice agents yet. In production, its end-of-turn decision lands a median 300ms before a conventional voice-activity detector would even detect a pause — measured on real telephony audio, not just clean VoIP.
Most voice AI on the market uses a cascaded approach: transcribe what the customer says, generate a response from that transcript, convert it back to audio. Everything the agent needs has to survive being flattened into text first, and a lot gets lost in translation. The result is both slower and less natural: no matter how fast your agent is, it’s waiting on that transcript before it can say anything.
Tango skips that step for the decisions that matter most, reading end-of-turn and interruption directly from audio. Transcripts still get generated for every call, but they no longer sit in the critical path.
Benefits of Tango’s Audio Processing
Tango isn’t just faster than competitor end-of-turn models, it’s also more accurate. This is crucial, since false positives mean more interruptions and false negatives mean uncomfortable pauses. With improved end-of-turn, conversations flow more freely. On the public LiveKit benchmark, it’s the fastest of 11 systems once you require at least 90% detection, missing just 15 of 400 true turn endings versus 54 for Soniox and 197 for Deepgram Flux. Against Smart Turn specifically, it fires 3.7 to 4.7 times less often mid-speech at equal turn-end coverage.
Genuine interruptions and quick backchannel cues like “mm-hmm” can look identical in a transcript, but they call for opposite behavior: stop talking for one, keep going for the other. Tango tells them apart in real time, so it doesn’t cut off a customer who’s simply agreeing. That’s cut interruptions by 64% versus our previous agent in production, recognizing a real interruption about 100ms after the customer starts talking.
Type “thank you” and there’s no way to tell if it was said with a smile or through gritted teeth. Text strips out tone, prosody, and pacing: the exact signals that make an interaction feel human. Audio-native processing keeps that signal intact.
The old pipeline adds roughly one second of latency before your agent starts thinking. Tango skips it: audio comes in, decisions come out, with both agent and customer channels processed at once. By going straight to streaming audio models, we can make assessments live every 80 ms, and run inference in less than half that time. There’s no transcript to wait for, and no second pipeline to run.
What This Looks (and Sounds) Like
A transcript-based model sees silence and guesses. Tango hears the customer still thinking, and responds at just the right time.
Tango can tell a real interruption from a conversational nod, so your agents don’t stop talking every time a customer says “okay”.
What This Means For Your Call Center
In practice, the difference between a natural-feeling agent and a robotic one comes down to about one second of latency. This gap shows up directly in the metrics you’re already tracking.
Customer Satisfaction Score (CSAT): A customer’s satisfaction is shaped by their overall experience rather than resolution alone. Fewer awkward pauses and false interruptions means fewer calls that feel broken, even when the agent is correct.
Average Handle Time (AHT): Every pause, interruption, or repeated “sorry, go ahead” adds seconds to your call. At call center volume, these seconds compound fast.
Abandonment: The customers you lose aren’t the ones who got a wrong answer. They’re the ones who hung up in frustration before you got the chance to help them at all.
A system near the top of this chart catches more end-of-turns, while a system to the left catches them sooner. Tango is the most accurate model for its speed, and the fastest model for its accuracy. Because the best, most natural conversations have both.
Built for the Calls You Actually Get
Most end-of-turn models are trained and benchmarked on clean, high-bandwidth audio. Real customer calls are not. They’re compressed, 8kHz telephony audio, running through phone systems built way before there were AI agents on the other end of the line. That’s why models that perform well in a demo often degrade the moment they meet a call center’s actual call volume.
Tango is built dual-channel and telephony-native from the ground up. It plugs directly into the inbound and outbound calling systems you already run, without asking you to upgrade your infrastructure first. That’s the difference between a benchmark number and a number you’ll actually see in production.
Tango is already live across Simple’s production traffic, delivering lower latency, higher accuracy, and more natural conversations for the calls our customers handle every day. It’s already shaping how we think about every agent we build from here.
They say it takes two to Tango. Schedule a demo today to see how Tango could sound for you.
Put your best rep on every call.
How is Tango different from a speech-to-text pipeline?
')"> Tango analyzes the audio directly instead of waiting for speech to become text and then turning a response back into audio. That preserves cues such as tone, pacing, and prosody while removing roughly one second of transcription latency from the response path. Transcripts are still generated after each call for review and reporting.
Why does Tango handle turn-taking and sentiment in one model?
')"> A caller’s tone and rhythm help reveal both how they feel and whether they have finished speaking. Tango processes those signals together in one pass. During development, Simple found that joint training improved both tasks and made the system simpler and cheaper to train and deploy.
Do I need to upgrade my phone system to use Tango?
')"> No. Tango handles compressed 8 kHz telephony audio as well as cleaner 24 kHz VoIP audio, and it connects to existing inbound and outbound calling systems.
Revenue Agent vs. Deflection Bot: How to Tell Them Apart
Build It on Vapi or ElevenLabs, or Buy a Platform?
How Omaha Steaks Cut Call Abandonment From 16% to 3%
