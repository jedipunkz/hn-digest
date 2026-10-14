---
source: "https://techstackups.com/comparisons/gemini-vs-elevenlabs-voice-cloning-tts/"
hn_url: "https://news.ycombinator.com/item?id=50086559"
title: "Gemini vs ElevenLabs Voice Cloning"
article_title: "Gemini vs ElevenLabs Voice Cloning and TTS (2026): Gemini 3.8 Flash vs Eleven v4 | Tech Stackups"
image: "https://techstackups.com/img/comparisons/gemini-vs-elevenlabs/cover.jpg"
author: "ritzaco"
captured_at: "2026-10-14T13:21:49Z"
capture_tool: "hn-digest"
hn_id: 50086559
score: 1
comments: 0
posted_at: "2026-10-14T12:21:26Z"
tags:
  - hacker-news
---

# Gemini vs ElevenLabs Voice Cloning

- HN: [50086559](https://news.ycombinator.com/item?id=50086559)
- Source: [techstackups.com](https://techstackups.com/comparisons/gemini-vs-elevenlabs-voice-cloning-tts/)
- Score: 1
- Comments: 0
- Posted: 2026-10-14T12:21:26Z

## Translation

Title: Gemini vs ElevenLabs Voice Cloning
Article title: Gemini vs ElevenLabs Voice Cloning and TTS (2026): Gemini 3.8 Flash vs Eleven v4 | Tech Stackups
Description: Gemini vs ElevenLabs on voice cloning and text to speech: I ran Gemini 3.8 Flash TTS and Eleven v4 through emotion, hard text, long-form, dialogue and cloning tests, with the audio so you can judge.

Article text:
Skip to main content Tech Stackups Home Topics Comparisons Guides Articles Explainers News AX Search Gemini vs ElevenLabs Voice Cloning and TTS (2026): Gemini 3.8 Flash vs Eleven v4
Google released Gemini 3.8 Flash TTS on 23 September 2026. ElevenLabs released Eleven v4 five days later. I put both through the same five tests, covering emotional tags, hard-to-say text, long-form narration, a two-person argument, and cloning my own voice.
ElevenLabs won three of the five (hard text, long-form, dialogue). Gemini won two: emotion, and voice cloning, where it kept my South African accent and ElevenLabs gave me a British one. If you came here for the ElevenLabs voice cloning question specifically, that result is the interesting one, and it's further down.
Google ships a "Flash TTS" and a "Flash-Lite TTS" under the same 3.8 label, which is confusing. I tested Flash, and Eleven v4 on the ElevenLabs side.
Gemini vs ElevenLabs: the quick verdict ​
This is one person listening. I tested in English only, and I ran each scenario once with default settings.
Same text, same direction. Both models got identical text. Direction went inline, sentence by sentence ( [whispering] for ElevenLabs, <whispering> for Gemini). I left Gemini's separate Style instructions box untouched so the comparison is like for like. (My first Gemini attempt used one overall style line and came out hushed throughout, so I threw it out as a test-design flaw.)
Free-tier web UIs for everything except cloning: ElevenLabs with v4 selected, and Google AI Studio.
One caveat in ElevenLabs' favour. It returns two outputs per prompt and I picked the better one. Gemini returns one.
Cloning ran through the APIs , because AI Studio wouldn't run cloning without billing and instant voice cloning on ElevenLabs needs a paid plan in the UI. I cloned my own voice from the same 29.5 second sample on both. A friend listened to the cloning results too, so that verdict isn't mine alone. The other four are mine only.
Voices. Scenarios 1 to 3 used Gemini's Cleo (warm and engaging) against ElevenLabs' Lauren (friendly, comforting and soft). The dialogue test added Mako and Jack John.
Tested in early October 2026. Scroll to the end to try a blind A/B on the clips and disagree with me.
Gemini voice cloning vs ElevenLabs voice cloning ​
What this tests: can the model sound like a specific person from about 30 seconds of audio, and keep that identity when the line changes from neutral to emotional to narration. I'm South African, which makes accent drift easy to hear.
Gemini asks for a 10 to 30 second sample plus a spoken consent recording, and the consent has to come from the same voice as the sample. ElevenLabs takes just the sample. Creating the clean clone took 5.54 seconds on Gemini and 3.23 on ElevenLabs.
Here are the three clean-sample lines, same text on both:
Listen vs ElevenLabs eleven_v4 Listen
Cloned voice: emotional line Does the clone keep its identity when the delivery changes? Gemini gemini-3.8-flash-tts Listen vs ElevenLabs eleven_v4 Listen
Cloned voice: narration Longer read from the same clones Gemini gemini-3.8-flash-tts Listen vs ElevenLabs eleven_v4 Listen
Gemini wins this one clearly. Its clones kept my accent and tone far better. The ElevenLabs clones kept a slight resemblance to my voice, but it sounded like my voice and tone had been overlaid on someone else's accent, a heavily British one. That held across all three lines, so it looks like a general problem with ElevenLabs cloning on my sample.
Neither model captured expression well from a 30 second sample, though Gemini was less bad, and its emotional line didn't convey the emotion accurately. The Gemini narration clip had the best voice and accent match of anything I tested, but it sounded damp and fuzzy next to Gemini's other clips, like a recording of me.
What happens with a noisy sample ​
I also recorded the same text in a noisier setting. Gemini rejected it outright at clone creation:
Voice mismatch detected. The speaker in the consent audio does not match the speaker in the voice sample.
My consent clip was recorded in a quiet room, so Gemini was comparing a quiet recording against a noisy one and decided they weren't the same person. That's a useful safeguard and an annoying one if your only sample is noisy.
ElevenLabs accepted the same noisy sample with no consent step. The clone was much worse than its clean one: a very slight resemblance, and bad accent and expression. Here's the noisy narration, ElevenLabs only, since Gemini produced nothing:
Only clone a voice you own or have permission to use. Gemini enforces consent at this step and ElevenLabs doesn't, which matters if you're building a product that lets users clone voices.
What this tests: does the model perform the whisper, the laugh and the sigh, or just read the word, and how smooth is the shift between them. The script runs from a whisper to excitement to a laugh to a voice cracking with emotion.
Listen vs ElevenLabs eleven_v4 Listen
Gemini sounded significantly more natural. It inferred pauses from context, and the pauses sound like someone talking to a person and waiting for their reaction. My favourite moment is at 0:16. The line is "We're giving you the Lisbon project," and Gemini went from a whisper to a slight exclamation with no emotion tag at all, purely because it understood the news was exciting. The worst moment was ElevenLabs at 0:18, where the shift from laugh to whispering is jarring.
My read on why: ElevenLabs seems to treat the emotion tag as a parameter and then read that sentence with it. Gemini seems to read the current sentence together with the next one, so the emotion spans both. That is my impression from listening. I didn't hear any tag spoken aloud in this test.
Listen vs ElevenLabs eleven_v4 Listen
ElevenLabs won the dialogue test by a distance. On Gemini the interruption feels like a pause, and the emotion lands as a pre-sentence emote. For the [laughs, then quietly] tag, Gemini produced a normal laugh and then kept speaking loudly, ignoring "then quietly". ElevenLabs produced a laugh that fit the conversation. At 0:12 it also inferred the severity of the emotion from context, which it did well in every scenario. Neither is perfect, and I flagged robotic glitches and some overacting. Generation took 14.55 s on Gemini and 6.26 s on ElevenLabs.
Numbers, names and hard text ​
What this tests: the stuff that breaks customer-facing voice. An order number, a date, a currency amount, a phone number, a URL, and a run of hard names. No tags on either side.
The rest were fine on both: $1,249.50, R22,800, acme-labs.io/returns, 14:45, "Thursday the 9th," HTTP 404, 2.4 GHz, Xiaomi, Worcestershire, Leicester and Edinburgh. I'm not giving a score.
Listen vs ElevenLabs eleven_v4 Listen
ElevenLabs wins, mostly on the dropped dash in the order number and the Irish name. But remember the two-takes caveat: ElevenLabs gave me two outputs and I chose the better one. Generation took 18.88 s on Gemini and 17.53 s on ElevenLabs.
What this tests: about 2.5 minutes of Dickens (the opening of Great Expectations ), plain text with no tags, to see whether pace creeps, the voice changes character, or the tone jumps.
ElevenLabs handled it without trouble. No pace creep, the voice kept its character, the paragraph breaks sounded intentional, and there were no sudden tone or loudness jumps.
Gemini failed on my first day of testing. It would speak only the first four to six words of the passage and then fail, repeatedly. I lost several retries to it. Here's the screen recording:
I retried the next day and the full passage generated. I don't know why it failed the first time. But even the successful run has a problem: between 2:06 and 2:08 the tone drops jarringly, in a way that doesn't match the emotion the surrounding text calls for. Generation took 45.9 s on Gemini and 48.28 s on ElevenLabs.
Listen vs ElevenLabs eleven_v4 Listen
ElevenLabs was the better read of context here. I was impressed at how well it inferred expression from the text alone.
Which is the best TTS model? Which is the best voice cloning model? ​
This covers two models. For the 16-provider view, see Best TTS in 2026: blind benchmark , and for the realtime and price angle see Cartesia vs ElevenLabs . Within these two:
Creators and audiobooks: ElevenLabs. It stayed consistent across 2.5 minutes of narration and Gemini didn't.
Customer-facing voice and IVR: ElevenLabs, with caveats. Gemini's phone-number read was better, but the dropped dash in an order number is the kind of error that costs real money.
Developers watching API cost: Gemini, by a wide margin, if its long-form reliability holds up for you.
Voice cloning: Gemini, at least for an accent like mine and a clean sample. It also enforces consent. ElevenLabs may do better with other voices, but I only cloned one.
Dialogue and character voices: ElevenLabs.
Expressive reads where you want the model to infer emotion: Gemini.
Listen without knowing which is which, vote, and then see my verdict next to yours.
I'd use ElevenLabs for anything that has to be right first time: narration, dialogue, and text full of numbers. I'd use Gemini when cost matters, when I want a model that reads emotion across sentences, or when I need a clone of a real voice that keeps its accent. If you take one thing from this, take the cloning result. On a clean 30 second sample, Gemini kept me sounding like me while ElevenLabs didn't. Test with your own voice before you commit.
I didn't test voice design from a text description, re-roll consistency, other languages, or the UI friction of each tool.
Hash Kader is a software engineer and technical writer at Ritza . He contributes hands-on testing of AI tools and developer products to TechStackups, with a focus on what they're actually like to use.
Gemini vs ElevenLabs: the quick verdict
Gemini voice cloning vs ElevenLabs voice cloning What happens with a noisy sample
Which is the best TTS model? Which is the best voice cloning model?

## Original Extract

Gemini vs ElevenLabs on voice cloning and text to speech: I ran Gemini 3.8 Flash TTS and Eleven v4 through emotion, hard text, long-form, dialogue and cloning tests, with the audio so you can judge.

Skip to main content Tech Stackups Home Topics Comparisons Guides Articles Explainers News AX Search Gemini vs ElevenLabs Voice Cloning and TTS (2026): Gemini 3.8 Flash vs Eleven v4
Google released Gemini 3.8 Flash TTS on 23 September 2026. ElevenLabs released Eleven v4 five days later. I put both through the same five tests, covering emotional tags, hard-to-say text, long-form narration, a two-person argument, and cloning my own voice.
ElevenLabs won three of the five (hard text, long-form, dialogue). Gemini won two: emotion, and voice cloning, where it kept my South African accent and ElevenLabs gave me a British one. If you came here for the ElevenLabs voice cloning question specifically, that result is the interesting one, and it's further down.
Google ships a "Flash TTS" and a "Flash-Lite TTS" under the same 3.8 label, which is confusing. I tested Flash, and Eleven v4 on the ElevenLabs side.
Gemini vs ElevenLabs: the quick verdict ​
This is one person listening. I tested in English only, and I ran each scenario once with default settings.
Same text, same direction. Both models got identical text. Direction went inline, sentence by sentence ( [whispering] for ElevenLabs, <whispering> for Gemini). I left Gemini's separate Style instructions box untouched so the comparison is like for like. (My first Gemini attempt used one overall style line and came out hushed throughout, so I threw it out as a test-design flaw.)
Free-tier web UIs for everything except cloning: ElevenLabs with v4 selected, and Google AI Studio.
One caveat in ElevenLabs' favour. It returns two outputs per prompt and I picked the better one. Gemini returns one.
Cloning ran through the APIs , because AI Studio wouldn't run cloning without billing and instant voice cloning on ElevenLabs needs a paid plan in the UI. I cloned my own voice from the same 29.5 second sample on both. A friend listened to the cloning results too, so that verdict isn't mine alone. The other four are mine only.
Voices. Scenarios 1 to 3 used Gemini's Cleo (warm and engaging) against ElevenLabs' Lauren (friendly, comforting and soft). The dialogue test added Mako and Jack John.
Tested in early October 2026. Scroll to the end to try a blind A/B on the clips and disagree with me.
Gemini voice cloning vs ElevenLabs voice cloning ​
What this tests: can the model sound like a specific person from about 30 seconds of audio, and keep that identity when the line changes from neutral to emotional to narration. I'm South African, which makes accent drift easy to hear.
Gemini asks for a 10 to 30 second sample plus a spoken consent recording, and the consent has to come from the same voice as the sample. ElevenLabs takes just the sample. Creating the clean clone took 5.54 seconds on Gemini and 3.23 on ElevenLabs.
Here are the three clean-sample lines, same text on both:
Listen vs ElevenLabs eleven_v4 Listen
Cloned voice: emotional line Does the clone keep its identity when the delivery changes? Gemini gemini-3.8-flash-tts Listen vs ElevenLabs eleven_v4 Listen
Cloned voice: narration Longer read from the same clones Gemini gemini-3.8-flash-tts Listen vs ElevenLabs eleven_v4 Listen
Gemini wins this one clearly. Its clones kept my accent and tone far better. The ElevenLabs clones kept a slight resemblance to my voice, but it sounded like my voice and tone had been overlaid on someone else's accent, a heavily British one. That held across all three lines, so it looks like a general problem with ElevenLabs cloning on my sample.
Neither model captured expression well from a 30 second sample, though Gemini was less bad, and its emotional line didn't convey the emotion accurately. The Gemini narration clip had the best voice and accent match of anything I tested, but it sounded damp and fuzzy next to Gemini's other clips, like a recording of me.
What happens with a noisy sample ​
I also recorded the same text in a noisier setting. Gemini rejected it outright at clone creation:
Voice mismatch detected. The speaker in the consent audio does not match the speaker in the voice sample.
My consent clip was recorded in a quiet room, so Gemini was comparing a quiet recording against a noisy one and decided they weren't the same person. That's a useful safeguard and an annoying one if your only sample is noisy.
ElevenLabs accepted the same noisy sample with no consent step. The clone was much worse than its clean one: a very slight resemblance, and bad accent and expression. Here's the noisy narration, ElevenLabs only, since Gemini produced nothing:
Only clone a voice you own or have permission to use. Gemini enforces consent at this step and ElevenLabs doesn't, which matters if you're building a product that lets users clone voices.
What this tests: does the model perform the whisper, the laugh and the sigh, or just read the word, and how smooth is the shift between them. The script runs from a whisper to excitement to a laugh to a voice cracking with emotion.
Listen vs ElevenLabs eleven_v4 Listen
Gemini sounded significantly more natural. It inferred pauses from context, and the pauses sound like someone talking to a person and waiting for their reaction. My favourite moment is at 0:16. The line is "We're giving you the Lisbon project," and Gemini went from a whisper to a slight exclamation with no emotion tag at all, purely because it understood the news was exciting. The worst moment was ElevenLabs at 0:18, where the shift from laugh to whispering is jarring.
My read on why: ElevenLabs seems to treat the emotion tag as a parameter and then read that sentence with it. Gemini seems to read the current sentence together with the next one, so the emotion spans both. That is my impression from listening. I didn't hear any tag spoken aloud in this test.
Listen vs ElevenLabs eleven_v4 Listen
ElevenLabs won the dialogue test by a distance. On Gemini the interruption feels like a pause, and the emotion lands as a pre-sentence emote. For the [laughs, then quietly] tag, Gemini produced a normal laugh and then kept speaking loudly, ignoring "then quietly". ElevenLabs produced a laugh that fit the conversation. At 0:12 it also inferred the severity of the emotion from context, which it did well in every scenario. Neither is perfect, and I flagged robotic glitches and some overacting. Generation took 14.55 s on Gemini and 6.26 s on ElevenLabs.
Numbers, names and hard text ​
What this tests: the stuff that breaks customer-facing voice. An order number, a date, a currency amount, a phone number, a URL, and a run of hard names. No tags on either side.
The rest were fine on both: $1,249.50, R22,800, acme-labs.io/returns, 14:45, "Thursday the 9th," HTTP 404, 2.4 GHz, Xiaomi, Worcestershire, Leicester and Edinburgh. I'm not giving a score.
Listen vs ElevenLabs eleven_v4 Listen
ElevenLabs wins, mostly on the dropped dash in the order number and the Irish name. But remember the two-takes caveat: ElevenLabs gave me two outputs and I chose the better one. Generation took 18.88 s on Gemini and 17.53 s on ElevenLabs.
What this tests: about 2.5 minutes of Dickens (the opening of Great Expectations ), plain text with no tags, to see whether pace creeps, the voice changes character, or the tone jumps.
ElevenLabs handled it without trouble. No pace creep, the voice kept its character, the paragraph breaks sounded intentional, and there were no sudden tone or loudness jumps.
Gemini failed on my first day of testing. It would speak only the first four to six words of the passage and then fail, repeatedly. I lost several retries to it. Here's the screen recording:
I retried the next day and the full passage generated. I don't know why it failed the first time. But even the successful run has a problem: between 2:06 and 2:08 the tone drops jarringly, in a way that doesn't match the emotion the surrounding text calls for. Generation took 45.9 s on Gemini and 48.28 s on ElevenLabs.
Listen vs ElevenLabs eleven_v4 Listen
ElevenLabs was the better read of context here. I was impressed at how well it inferred expression from the text alone.
Which is the best TTS model? Which is the best voice cloning model? ​
This covers two models. For the 16-provider view, see Best TTS in 2026: blind benchmark , and for the realtime and price angle see Cartesia vs ElevenLabs . Within these two:
Creators and audiobooks: ElevenLabs. It stayed consistent across 2.5 minutes of narration and Gemini didn't.
Customer-facing voice and IVR: ElevenLabs, with caveats. Gemini's phone-number read was better, but the dropped dash in an order number is the kind of error that costs real money.
Developers watching API cost: Gemini, by a wide margin, if its long-form reliability holds up for you.
Voice cloning: Gemini, at least for an accent like mine and a clean sample. It also enforces consent. ElevenLabs may do better with other voices, but I only cloned one.
Dialogue and character voices: ElevenLabs.
Expressive reads where you want the model to infer emotion: Gemini.
Listen without knowing which is which, vote, and then see my verdict next to yours.
I'd use ElevenLabs for anything that has to be right first time: narration, dialogue, and text full of numbers. I'd use Gemini when cost matters, when I want a model that reads emotion across sentences, or when I need a clone of a real voice that keeps its accent. If you take one thing from this, take the cloning result. On a clean 30 second sample, Gemini kept me sounding like me while ElevenLabs didn't. Test with your own voice before you commit.
I didn't test voice design from a text description, re-roll consistency, other languages, or the UI friction of each tool.
Hash Kader is a software engineer and technical writer at Ritza . He contributes hands-on testing of AI tools and developer products to TechStackups, with a focus on what they're actually like to use.
Gemini vs ElevenLabs: the quick verdict
Gemini voice cloning vs ElevenLabs voice cloning What happens with a noisy sample
Which is the best TTS model? Which is the best voice cloning model?
