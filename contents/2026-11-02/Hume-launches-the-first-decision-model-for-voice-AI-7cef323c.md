---
source: "https://www.hume.ai/products/decision-model-api"
hn_url: "https://news.ycombinator.com/item?id=50339949"
title: "Hume launches the first decision model for voice AI"
article_title: "Decision Model API | Hume AI"
image: "https://www.hume.ai/opengraph-image.jpg"
author: "dndx"
captured_at: "2026-11-02T17:53:32Z"
capture_tool: "hn-digest"
hn_id: 50339949
score: 1
comments: 0
posted_at: "2026-11-02T17:07:10Z"
tags:
  - hacker-news
---

# Hume launches the first decision model for voice AI

- HN: [50339949](https://news.ycombinator.com/item?id=50339949)
- Source: [www.hume.ai](https://www.hume.ai/products/decision-model-api)
- Score: 1
- Comments: 0
- Posted: 2026-11-02T17:07:10Z

## Translation

Title: Hume launches the first decision model for voice AI
Article title: Decision Model API | Hume AI
Description: Hume's audio decision model: send audio, a question and possible answers, and get a probability for each.

Article text:
Decision Model API | Hume AI New Introducing the Hume Audio Decisions API One model, any question, straight from audio hume.ai logo Emotion Science
Decision models for voice AI D e c i s i o n m o d e l s
f o r v o i c e A I
The bar is human. Hume diagnoses system breaks and feeds failures into improvement. T h e b a r i s h u m a n . H u m e d i a g n o s e s s y s t e m b r e a k s a n d f e e d s f a i l u r e s i n t o i m p r o v e m e n t .
A sample telehealth call plays while three questions about it are answered for every five seconds of audio.
Sample call Microphone Upload files Listening.. Stop Telehealth consultation 0:00 / 0:56
0:00 0:14 0:28 0:42 0:56 Context A telehealth consultation between a clinician and a patient.
Does the patient sound anxious? Pick One Which best describes the clinician’s tone? Scale How reassured does the patient sound by the end? Add question You can add or change a question any time. Yes No Not sure Answered 0 of 12 chunks · every 5 seconds is answered on its own
Why audio needs a decision model
Voice AI considers tone, pauses, and sounds. Developers set Q&A; the audio decision model gives probabilities for quick audio judgments.
Watch Video Different questions. The same audio.
Ask objective, subjective or perceptual questions about speech and sound. Select a tab to explore examples.
Find the audio worth acting on
Apply several questions to the same recordings, from naturalness and clarity to empathy and timing.
Repeat evaluations across your collection as your model or criteria change.
Pinpoint cases for your next step
Find candidate examples for training, isolate evaluation failures, and surface cases for escalation review.
Customer sounds frustrated is Yes Issue resolved is No Add Question Call 56 1:12 Yes 71 % No 63 % Call 908 1:55 Yes 66 % No 48 % Call 122 1:24 Yes 68 % No 50 % Call 87 0:58 Yes 74 % No 38 % Call 341 2:07 Yes 61 % No 57 % More reliable than a single listener
The Audio Decision Model agrees with the human crowd 70.7% of the time, while a single human rater agrees 56.5% of the time.
Agreement with a panel of human raters : Audio Decision Model 70.7%, Individual Rater 56.5% .
Agreement with a panel of human raters
Accuracy against speed : Hume’s Decision Model averages 0.82 over 10 leading benchmarks at 0.23 seconds median latency per screen, scoring higher and answering faster than Gemini 3.8 Flash, Qwen3.8-Omni-Flash, Inkling-Small, gpt-audio-1.5, Gemini 3.5 Flash Lite and Nemotron 3 Nano Omni.
0.1 s 0.2 s 1 s 2 s 10 s 0.50 0.60 0.70 0.80 Median latency per screen, seconds (log scale) Average accuracy across leading benchmarks Hume’s Decision Model Gemini 3.8 Flash Qwen3.8-Omni-Flash Inkling-Small gpt-audio-1.5 Gemini 3.5 Flash Lite Nemotron 3 Nano Omni The Audio Decision Model answers audio questions faster and more accurately than any other model on the market.
Compare the judgments that matter
Accuracy across the criteria you care about.
Hume’s Decision Model Hume 82%
Qwen3.8-Omni-Flash Alibaba 76%
Inkling-Small Thinking Machines 67%
What is Hume’s audio decision model?
It’s Hume’s audio-native decision model. Developers submit audio, a question and possible answers; it returns a probability for each answer. It helps teams evaluate speech and sound, assess how a voice comes across to listeners, and inform decisions in voice applications.
What is a decision model for audio or voice?
How is the audio decision model different from speech-to-text or a general-purpose audio LLM?
Can I define my own questions and ask several in one API call?
How can the audio decision model support voice-model evaluation and training?
How do I request Audio Decision Model API access?
Audio decisions that fit your stack.
Send audio, your question and possible answers. The audio decision model returns a probability for each
curl https://api.hume.ai/v0/decisions/file \
-H "X-Hume-Api-Key: $HUME_API_KEY" \
-F file=@support_call.wav \
-F 'request={
"questions": {
"background": {
"type": "noul",
"instructions": "Is there background speech?"
},
"speakers": {
"type": "choice",
"instructions": "How many speakers can you hear?",
"criteria": [
{ "name": "One" },
{ "name": "Two" },
{ "name": "Three or more" }
]
}
}
};type=application/json'
curl https://api.hume.ai/v0/decisions/file \
-H "X-Hume-Api-Key: $HUME_API_KEY" \
-F file=@support_call.wav \
-F 'request={
"questions": {
"background": {
"type": "noul",
"instructions": "Is there background speech?"
},
"speakers": {
"type": "choice",
"instructions": "How many speakers can you hear?",
"criteria": [
{ "name": "One" },
{ "name": "Two" },
{ "name": "Three or more" }
]
}
}
};type=application/json'
Klingt der Sprecher gelangweilt? Is there background speech? どんな言語が話されているのですか？ Apa kabar anda saat ini? Có bao nhiêu loa phát ra âm thanh? ¿El cliente suena frustrado? Use the Decision Model API as a decision layer in a voice-agent system with speech recognition and generation.
Make emotionally
intelligent voice AI measurable
Get new benchmark research and leaderboard updates.
Methodology (opens in a new tab)

## Original Extract

Hume's audio decision model: send audio, a question and possible answers, and get a probability for each.

Decision Model API | Hume AI New Introducing the Hume Audio Decisions API One model, any question, straight from audio hume.ai logo Emotion Science
Decision models for voice AI D e c i s i o n m o d e l s
f o r v o i c e A I
The bar is human. Hume diagnoses system breaks and feeds failures into improvement. T h e b a r i s h u m a n . H u m e d i a g n o s e s s y s t e m b r e a k s a n d f e e d s f a i l u r e s i n t o i m p r o v e m e n t .
A sample telehealth call plays while three questions about it are answered for every five seconds of audio.
Sample call Microphone Upload files Listening.. Stop Telehealth consultation 0:00 / 0:56
0:00 0:14 0:28 0:42 0:56 Context A telehealth consultation between a clinician and a patient.
Does the patient sound anxious? Pick One Which best describes the clinician’s tone? Scale How reassured does the patient sound by the end? Add question You can add or change a question any time. Yes No Not sure Answered 0 of 12 chunks · every 5 seconds is answered on its own
Why audio needs a decision model
Voice AI considers tone, pauses, and sounds. Developers set Q&A; the audio decision model gives probabilities for quick audio judgments.
Watch Video Different questions. The same audio.
Ask objective, subjective or perceptual questions about speech and sound. Select a tab to explore examples.
Find the audio worth acting on
Apply several questions to the same recordings, from naturalness and clarity to empathy and timing.
Repeat evaluations across your collection as your model or criteria change.
Pinpoint cases for your next step
Find candidate examples for training, isolate evaluation failures, and surface cases for escalation review.
Customer sounds frustrated is Yes Issue resolved is No Add Question Call 56 1:12 Yes 71 % No 63 % Call 908 1:55 Yes 66 % No 48 % Call 122 1:24 Yes 68 % No 50 % Call 87 0:58 Yes 74 % No 38 % Call 341 2:07 Yes 61 % No 57 % More reliable than a single listener
The Audio Decision Model agrees with the human crowd 70.7% of the time, while a single human rater agrees 56.5% of the time.
Agreement with a panel of human raters : Audio Decision Model 70.7%, Individual Rater 56.5% .
Agreement with a panel of human raters
Accuracy against speed : Hume’s Decision Model averages 0.82 over 10 leading benchmarks at 0.23 seconds median latency per screen, scoring higher and answering faster than Gemini 3.8 Flash, Qwen3.8-Omni-Flash, Inkling-Small, gpt-audio-1.5, Gemini 3.5 Flash Lite and Nemotron 3 Nano Omni.
0.1 s 0.2 s 1 s 2 s 10 s 0.50 0.60 0.70 0.80 Median latency per screen, seconds (log scale) Average accuracy across leading benchmarks Hume’s Decision Model Gemini 3.8 Flash Qwen3.8-Omni-Flash Inkling-Small gpt-audio-1.5 Gemini 3.5 Flash Lite Nemotron 3 Nano Omni The Audio Decision Model answers audio questions faster and more accurately than any other model on the market.
Compare the judgments that matter
Accuracy across the criteria you care about.
Hume’s Decision Model Hume 82%
Qwen3.8-Omni-Flash Alibaba 76%
Inkling-Small Thinking Machines 67%
What is Hume’s audio decision model?
It’s Hume’s audio-native decision model. Developers submit audio, a question and possible answers; it returns a probability for each answer. It helps teams evaluate speech and sound, assess how a voice comes across to listeners, and inform decisions in voice applications.
What is a decision model for audio or voice?
How is the audio decision model different from speech-to-text or a general-purpose audio LLM?
Can I define my own questions and ask several in one API call?
How can the audio decision model support voice-model evaluation and training?
How do I request Audio Decision Model API access?
Audio decisions that fit your stack.
Send audio, your question and possible answers. The audio decision model returns a probability for each
curl https://api.hume.ai/v0/decisions/file \
-H "X-Hume-Api-Key: $HUME_API_KEY" \
-F file=@support_call.wav \
-F 'request={
"questions": {
"background": {
"type": "noul",
"instructions": "Is there background speech?"
},
"speakers": {
"type": "choice",
"instructions": "How many speakers can you hear?",
"criteria": [
{ "name": "One" },
{ "name": "Two" },
{ "name": "Three or more" }
]
}
}
};type=application/json'
curl https://api.hume.ai/v0/decisions/file \
-H "X-Hume-Api-Key: $HUME_API_KEY" \
-F file=@support_call.wav \
-F 'request={
"questions": {
"background": {
"type": "noul",
"instructions": "Is there background speech?"
},
"speakers": {
"type": "choice",
"instructions": "How many speakers can you hear?",
"criteria": [
{ "name": "One" },
{ "name": "Two" },
{ "name": "Three or more" }
]
}
}
};type=application/json'
Klingt der Sprecher gelangweilt? Is there background speech? どんな言語が話されているのですか？ Apa kabar anda saat ini? Có bao nhiêu loa phát ra âm thanh? ¿El cliente suena frustrado? Use the Decision Model API as a decision layer in a voice-agent system with speech recognition and generation.
Make emotionally
intelligent voice AI measurable
Get new benchmark research and leaderboard updates.
Methodology (opens in a new tab)
