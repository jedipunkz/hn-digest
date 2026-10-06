---
source: "https://hieraticbench.vercel.app/"
hn_url: "https://news.ycombinator.com/item?id=49982822"
title: "Show HN: HieraticBench – Can AI read ancient Egyptian handwriting?"
article_title: "HieraticBench. Can AI read ancient Egyptian handwriting?"
image: "https://hieraticbench.vercel.app/og.png"
author: "amoursy"
captured_at: "2026-10-06T19:23:17Z"
capture_tool: "hn-digest"
hn_id: 49982822
score: 1
comments: 0
posted_at: "2026-10-06T19:22:15Z"
tags:
  - hacker-news
---

# Show HN: HieraticBench – Can AI read ancient Egyptian handwriting?

- HN: [49982822](https://news.ycombinator.com/item?id=49982822)
- Source: [hieraticbench.vercel.app](https://hieraticbench.vercel.app/)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T19:22:15Z

## Translation

Title: Show HN: HieraticBench – Can AI read ancient Egyptian handwriting?
Article title: HieraticBench. Can AI read ancient Egyptian handwriting?
Description: An Oxford Egyptologist wrote one sentence in hieratic, the handwriting of ancient Egypt. No AI model can read it. HieraticBench measures how close they are.
HN text: I'm obsessed with Egyptology. During covid, I wanted to get a hieratic tattoo. For those who don't know, hieratic is the cursive form of hieroglyphs that was used in day-to-day life in ancient Egypt. It's as old as hieroglyphs. I then went on a mission to have a sentence translated into hieratic. Luckily, in 2022, I was able to commission an Oxford Egyptologist to write it for me. After that, I took it to a hieratic specialist I found on Instagram to redraw it with "biblical" accuracy. When generative AI came out, I realized that models very rarely identify hieratic and almost never translate it correctly. Every single time a new model came out, I ended up testing my sentence and hieratic texts in general, and I realized this could actually become a benchmark. Since so many benchmarks are becoming saturated, we need to find weird ways to test new models. That's why I decided to build HieraticBench. The first version of the benchmark has 268 items, including real documents, signs, and other Egyptian scripts, which I used as controls, as well as my unpublished sentence in two different renditions. The current state-of-the-art foundation models are hilariously bad at identifying my sentence, especially the one the Egyptologist wrote with her pen. Some of them said Tibetan, Urdu, Korean, and, funnily enough, "Reformed Egyptian," which is from the Book of Mormon. On real documents, the best model in the benchmark identified hieratic correctly 95% of the time, and the best score for reading individual signs was around 13%. That's when we tell the model specifically that it's hieratic. None of them can really translate a sealed sentence that's never been published and has no answer key. I haven't run every model on every task just to keep costs down. My next steps were to run Astra and Gemini 3.1 Pro on real documents. Currently, only the Claude models have done the sign readings. And just FYI, I have no experience whatsoever running benchmarks or evaluations, so not 100% sure I scored things right. Things we need:
- Running the models I haven't covered
- More sealed sentences and labeled signs If you know hieratic or have an Egyptologist friend, reach out at hieraticbench@veeza.ai. Source code: https://github.com/alymoursy/hieraticbench I'll be around to answer questions.

Article text:
HieraticBench. Can AI read ancient Egyptian handwriting? HieraticBench v0.1 Leaderboard Dataset Method Join Open menu Claude Opus 5.5 said Not a known writing system . Wrong.
Claude Fable 5.1 said Not a real writing system . Wrong.
Claude Fable 5.1 said Hieratic, Demotic, Meroitic, Pahlavi or Elbasan. None confirmed . Wrong.
Claude Fable 5.1 said Amenemhat, “Amun is in front” . Wrong.
Claude Haiku 4.5 said Urdu, “One should learn every day” . Wrong.
Claude Fable 5.1 said Pahawh Hmong . Wrong.
Claude Fable 5.1 said Kharosthi . Wrong.
Claude Fable 5.1 said Old Uyghur . Wrong.
Claude Sonnet 5.5 said Nüshu . Wrong.
GPT-6 Astra said Tibetan . Wrong.
GPT-6 Astra said Geba . Wrong.
GPT-6 Astra said Marchung . Wrong.
GPT-6 Astra said Marchen . Wrong.
GPT-6.1 Sol said Tibetan . Wrong.
GPT-6.1 Sol said Enochian . Wrong.
Gemini 3.1 Pro said Paleo-Hebrew . Wrong.
Gemini 3.1 Pro said Reformed Egyptian . Wrong.
Gemini 3.8 Flash said Nabataean . Wrong.
Gemini 3.8 Flash said Linear Elamite . Wrong.
Gemini 3.8 Flash said Proto-Sinaitic . Wrong.
Gemini 3.8 Flash said Tibetan . Wrong.
Grok 4.7 said Georgian Mkhedruli . Wrong.
Grok 4.7 said Baybayin . Wrong.
Kimi K3 said Kharosthi . Wrong.
Llama 4 Maverick said Mongolian . Wrong.
Llama 4 Maverick said Hangul . Wrong.
Llama 4 Maverick said Mkhedruli Georgian . Wrong.
Mistral Medium 3.5 said Rongorongo . Wrong.
Mistral Medium 3.5 said SINHALA . Wrong.
Mistral Medium 3.5 said Thai . Wrong.
Mistral Medium 3.5 said Syriac . Wrong.
Qwen3.8 Max said Mongolian script . Wrong.
Claude Opus 5.5 said Not a known writing system
Claude Fable 5.1 said Not a real writing system
Claude Fable 5.1 said Hieratic, Demotic, Meroitic, Pahlavi or Elbasan. None confirmed
Claude Fable 5.1 said Amenemhat, “Amun is in front”
Claude Haiku 4.5 said Urdu, “One should learn every day”
Claude Fable 5.1 said Pahawh Hmong
Claude Fable 5.1 said Kharosthi
Claude Fable 5.1 said Old Uyghur
Gemini 3.1 Pro said Paleo-Hebrew
Gemini 3.1 Pro said Reformed Egyptian
Gemini 3.8 Flash said Nabataean
Gemini 3.8 Flash said Linear Elamite
Gemini 3.8 Flash said Proto-Sinaitic
Grok 4.7 said Georgian Mkhedruli
Llama 4 Maverick said Mongolian
Llama 4 Maverick said Mkhedruli Georgian
Mistral Medium 3.5 said Rongorongo
Mistral Medium 3.5 said SINHALA
Mistral Medium 3.5 said Syriac
Qwen3.8 Max said Mongolian script
An Oxford Egyptologist wrote it for me in 2022, in hieratic, the everyday handwriting of ancient Egypt. We tested 13 models. Not one could even name the script.
Aly Moursy , Founder of Veeza AI ( YC F26 )
The same image, the same plain question. Asked in the apps people actually use, and through the API.
Can you identify and translate this script?
I can't confidently identify this as any known writing system, so I can't give you a translation. I'd rather say so than invent one.
It doesn't match Chinese, Japanese, Korean, Tangut, Khitan, Jurchen, Nüshu, or the common shorthand systems like Gregg or Pitman.
Can you translate this script? First identify what it is and then translate it!
I looked at this carefully, and I don't think it's a real writing system, so I can't give you a translation without making one up.
Follow-up after the cold attempt: This may be a script that's no longer in use or a reconstruction of one.
Even against extinct and reconstructed scripts, I still can't place it, so I'd rather not guess at a translation.
It listed Egyptian hieratic and Demotic among the closest visual families, then rejected them.
Named the right family only after a nudge, then rejected it
Told the image was hieratic and that a speculative transcription was welcome.
What survives narrowing: a personal name built on Amun, most likely Amenemhat.
Confident sign-by-sign reading. Wrong.
5 October 2026 , API, HieraticBench harness
Can you identify and translate this script?
This appears to be Urdu script (a variant of Persian/Arabic script).
It then supplied a word-by-word Urdu reading and translated it as “One should learn every day.”
Wrong script, invented translation
Both images, several tries each, the same plain question
Opus 5.5 checked Tangut, Khitan, Jurchen and Nüshu. It checked Gregg shorthand. It never checked Egypt.
Astra also failed to name the script. We didn't keep the transcript.
They know what a papyrus looks like. They can't read the writing.
Show the best model a real hieratic document, a papyrus, a pottery shard or a plate from an Egyptology book, and it names the script 95% of the time. Show it one sentence in the same script, written fresh by an expert, and it scores 0% across 78 tries.
Ask it which hieroglyph a single hieratic sign stands for and the best model is right 13% of the time. Models seem to know what a page of hieratic looks like. They don't know the signs it is made of.
The goal isn't this sentence. It's hieratic.
The sentence is the exam. Its answer has never been published and isn't stored anywhere, so a model can't pass by remembering something it read online. The only way through is to actually read hieratic.
What we want is AI that can read the handwriting of ancient Egypt, sign by sign, on any papyrus or pottery shard, and help the few people who read it today with the many documents still waiting. A model that can do that will read this sentence too.
Hieroglyphs were for stone. Hieratic was for everything else.
Hieratic is the cursive form of Egyptian hieroglyphs, written fast with a reed brush on papyrus, pottery shards and wooden boards. For more than three thousand years it was how Egypt actually wrote. Letters, tax records, medical manuals, maths problems, love poems, the stories people told.
The Rhind Mathematical Papyrus is hieratic. So is the Edwin Smith surgical papyrus, the Tale of Sinuhe, and the Diary of Merer, a logbook kept during the building of the Great Pyramid and the oldest inscribed papyrus ever found.
Hieroglyphs get the museum walls. Hieratic is where the people are. And almost nobody alive can read it.
Four rungs, from seeing to reading.
Every rung after the first tells the model it is looking at hieratic, so a model that can't name the script still gets a fair shot at reading it.
Asked cold, the way you would ask a friend. Demotic and hieroglyphic controls sit in the set, so answering hieratic every time doesn't pay.
Best so far 95% on real documents, 0% on the sentence
Which hieroglyph is each sign?
Answered in Gardiner sign-list codes, the standard Egyptologists use. Scored by sign error rate.
Standard Egyptological transliteration. Not scored on the sentence, since no answer is stored.
English. Not scored on the sentence either. Scoring is built in for future public texts.
Every score comes from the open harness, through Anthropic's API for Claude and OpenRouter for everyone else. The first two columns ask a model to name the script, on the professor's sentence and on real ancient documents. Updated 5 October 2026 .
Claude Fable 5.1 (high) declined 6 questions, which count as wrong. Real documents counts hieratic documents only, not the demotic and hieroglyphic controls. Transliterating and translating the sentence aren't scored, because no answer is stored anywhere. Chat-app transcripts above are quoted for the record and never counted here.
268 items. 87 real hieratic documents, 29 controls in other Egyptian scripts, 150 single signs and 2 sealed sentences . Every public image is openly licensed and credited.
Sealed © Aly Moursy. Free to use for evaluating models.
Sealed © Aly Moursy. Free to use for evaluating models.
Papyrus Berlin 3022 (Story of Sinuhe), opening section
Ägyptisches Museum und Papyrussammlung, Berlin
Ostracon with Pharaoh Spearing a Lion and a Royal Hymn on its Back, Met 26.7.1453
ca. 1200–1080 BCE, The Metropolitan Museum of Art, New York
Hieratic CC0 (Met Open Access)
Rhind Mathematical Papyrus (BM EA 10057), detail
c. 1650 BCE, British Museum, London (EA 10057)
Ostracon: letter from the scribe Mose to the vizier Pesiur (Ashmolean HO 71)
13th century BCE, Ashmolean Museum, Oxford (HO 71)
Marriage Contract, Met 35.4.1a, b
380–343 BCE, The Metropolitan Museum of Art, New York
Single hieratic sign, London, British Museum, EA 9999 (AKU HT 10440)
New Kingdom, Dynasty 20, Ramesses III, regnal year 32, month 3, season šm.w, day 6
Help AI learn to read hieratic.
This only works with more people. It needs the few who can read hieratic, and the people building the models that one day might.
AI researchers Run the harness on your model with your own key, or build a reader from scratch. Every result goes on the public leaderboard. Run the benchmark Everyone else Share it. Introduce us to an Egyptologist. Sponsor a commissioned sentence. The benchmark grows one sentence at a time. Get in touch Or go see it for yourself.
Egypt's museums are full of hieratic, on papyrus and on pottery shards. Deir el-Medina, near Luxor, is the village where the workmen who built the royal tombs left thousands of notes in it.
Need a visa? Veeza AI handles the Egypt e-visa before you fly.
The Nile Abu Simbel HieraticBench was started by Aly Moursy , founder of Veeza AI ( YC F26 ), in 2026. Code is MIT licensed and results are CC BY 4.0. Every image is credited on the dataset page .

## Original Extract

An Oxford Egyptologist wrote one sentence in hieratic, the handwriting of ancient Egypt. No AI model can read it. HieraticBench measures how close they are.

I'm obsessed with Egyptology. During covid, I wanted to get a hieratic tattoo. For those who don't know, hieratic is the cursive form of hieroglyphs that was used in day-to-day life in ancient Egypt. It's as old as hieroglyphs. I then went on a mission to have a sentence translated into hieratic. Luckily, in 2022, I was able to commission an Oxford Egyptologist to write it for me. After that, I took it to a hieratic specialist I found on Instagram to redraw it with "biblical" accuracy. When generative AI came out, I realized that models very rarely identify hieratic and almost never translate it correctly. Every single time a new model came out, I ended up testing my sentence and hieratic texts in general, and I realized this could actually become a benchmark. Since so many benchmarks are becoming saturated, we need to find weird ways to test new models. That's why I decided to build HieraticBench. The first version of the benchmark has 268 items, including real documents, signs, and other Egyptian scripts, which I used as controls, as well as my unpublished sentence in two different renditions. The current state-of-the-art foundation models are hilariously bad at identifying my sentence, especially the one the Egyptologist wrote with her pen. Some of them said Tibetan, Urdu, Korean, and, funnily enough, "Reformed Egyptian," which is from the Book of Mormon. On real documents, the best model in the benchmark identified hieratic correctly 95% of the time, and the best score for reading individual signs was around 13%. That's when we tell the model specifically that it's hieratic. None of them can really translate a sealed sentence that's never been published and has no answer key. I haven't run every model on every task just to keep costs down. My next steps were to run Astra and Gemini 3.1 Pro on real documents. Currently, only the Claude models have done the sign readings. And just FYI, I have no experience whatsoever running benchmarks or evaluations, so not 100% sure I scored things right. Things we need:
- Running the models I haven't covered
- More sealed sentences and labeled signs If you know hieratic or have an Egyptologist friend, reach out at hieraticbench@veeza.ai. Source code: https://github.com/alymoursy/hieraticbench I'll be around to answer questions.

HieraticBench. Can AI read ancient Egyptian handwriting? HieraticBench v0.1 Leaderboard Dataset Method Join Open menu Claude Opus 5.5 said Not a known writing system . Wrong.
Claude Fable 5.1 said Not a real writing system . Wrong.
Claude Fable 5.1 said Hieratic, Demotic, Meroitic, Pahlavi or Elbasan. None confirmed . Wrong.
Claude Fable 5.1 said Amenemhat, “Amun is in front” . Wrong.
Claude Haiku 4.5 said Urdu, “One should learn every day” . Wrong.
Claude Fable 5.1 said Pahawh Hmong . Wrong.
Claude Fable 5.1 said Kharosthi . Wrong.
Claude Fable 5.1 said Old Uyghur . Wrong.
Claude Sonnet 5.5 said Nüshu . Wrong.
GPT-6 Astra said Tibetan . Wrong.
GPT-6 Astra said Geba . Wrong.
GPT-6 Astra said Marchung . Wrong.
GPT-6 Astra said Marchen . Wrong.
GPT-6.1 Sol said Tibetan . Wrong.
GPT-6.1 Sol said Enochian . Wrong.
Gemini 3.1 Pro said Paleo-Hebrew . Wrong.
Gemini 3.1 Pro said Reformed Egyptian . Wrong.
Gemini 3.8 Flash said Nabataean . Wrong.
Gemini 3.8 Flash said Linear Elamite . Wrong.
Gemini 3.8 Flash said Proto-Sinaitic . Wrong.
Gemini 3.8 Flash said Tibetan . Wrong.
Grok 4.7 said Georgian Mkhedruli . Wrong.
Grok 4.7 said Baybayin . Wrong.
Kimi K3 said Kharosthi . Wrong.
Llama 4 Maverick said Mongolian . Wrong.
Llama 4 Maverick said Hangul . Wrong.
Llama 4 Maverick said Mkhedruli Georgian . Wrong.
Mistral Medium 3.5 said Rongorongo . Wrong.
Mistral Medium 3.5 said SINHALA . Wrong.
Mistral Medium 3.5 said Thai . Wrong.
Mistral Medium 3.5 said Syriac . Wrong.
Qwen3.8 Max said Mongolian script . Wrong.
Claude Opus 5.5 said Not a known writing system
Claude Fable 5.1 said Not a real writing system
Claude Fable 5.1 said Hieratic, Demotic, Meroitic, Pahlavi or Elbasan. None confirmed
Claude Fable 5.1 said Amenemhat, “Amun is in front”
Claude Haiku 4.5 said Urdu, “One should learn every day”
Claude Fable 5.1 said Pahawh Hmong
Claude Fable 5.1 said Kharosthi
Claude Fable 5.1 said Old Uyghur
Gemini 3.1 Pro said Paleo-Hebrew
Gemini 3.1 Pro said Reformed Egyptian
Gemini 3.8 Flash said Nabataean
Gemini 3.8 Flash said Linear Elamite
Gemini 3.8 Flash said Proto-Sinaitic
Grok 4.7 said Georgian Mkhedruli
Llama 4 Maverick said Mongolian
Llama 4 Maverick said Mkhedruli Georgian
Mistral Medium 3.5 said Rongorongo
Mistral Medium 3.5 said SINHALA
Mistral Medium 3.5 said Syriac
Qwen3.8 Max said Mongolian script
An Oxford Egyptologist wrote it for me in 2022, in hieratic, the everyday handwriting of ancient Egypt. We tested 13 models. Not one could even name the script.
Aly Moursy , Founder of Veeza AI ( YC F26 )
The same image, the same plain question. Asked in the apps people actually use, and through the API.
Can you identify and translate this script?
I can't confidently identify this as any known writing system, so I can't give you a translation. I'd rather say so than invent one.
It doesn't match Chinese, Japanese, Korean, Tangut, Khitan, Jurchen, Nüshu, or the common shorthand systems like Gregg or Pitman.
Can you translate this script? First identify what it is and then translate it!
I looked at this carefully, and I don't think it's a real writing system, so I can't give you a translation without making one up.
Follow-up after the cold attempt: This may be a script that's no longer in use or a reconstruction of one.
Even against extinct and reconstructed scripts, I still can't place it, so I'd rather not guess at a translation.
It listed Egyptian hieratic and Demotic among the closest visual families, then rejected them.
Named the right family only after a nudge, then rejected it
Told the image was hieratic and that a speculative transcription was welcome.
What survives narrowing: a personal name built on Amun, most likely Amenemhat.
Confident sign-by-sign reading. Wrong.
5 October 2026 , API, HieraticBench harness
Can you identify and translate this script?
This appears to be Urdu script (a variant of Persian/Arabic script).
It then supplied a word-by-word Urdu reading and translated it as “One should learn every day.”
Wrong script, invented translation
Both images, several tries each, the same plain question
Opus 5.5 checked Tangut, Khitan, Jurchen and Nüshu. It checked Gregg shorthand. It never checked Egypt.
Astra also failed to name the script. We didn't keep the transcript.
They know what a papyrus looks like. They can't read the writing.
Show the best model a real hieratic document, a papyrus, a pottery shard or a plate from an Egyptology book, and it names the script 95% of the time. Show it one sentence in the same script, written fresh by an expert, and it scores 0% across 78 tries.
Ask it which hieroglyph a single hieratic sign stands for and the best model is right 13% of the time. Models seem to know what a page of hieratic looks like. They don't know the signs it is made of.
The goal isn't this sentence. It's hieratic.
The sentence is the exam. Its answer has never been published and isn't stored anywhere, so a model can't pass by remembering something it read online. The only way through is to actually read hieratic.
What we want is AI that can read the handwriting of ancient Egypt, sign by sign, on any papyrus or pottery shard, and help the few people who read it today with the many documents still waiting. A model that can do that will read this sentence too.
Hieroglyphs were for stone. Hieratic was for everything else.
Hieratic is the cursive form of Egyptian hieroglyphs, written fast with a reed brush on papyrus, pottery shards and wooden boards. For more than three thousand years it was how Egypt actually wrote. Letters, tax records, medical manuals, maths problems, love poems, the stories people told.
The Rhind Mathematical Papyrus is hieratic. So is the Edwin Smith surgical papyrus, the Tale of Sinuhe, and the Diary of Merer, a logbook kept during the building of the Great Pyramid and the oldest inscribed papyrus ever found.
Hieroglyphs get the museum walls. Hieratic is where the people are. And almost nobody alive can read it.
Four rungs, from seeing to reading.
Every rung after the first tells the model it is looking at hieratic, so a model that can't name the script still gets a fair shot at reading it.
Asked cold, the way you would ask a friend. Demotic and hieroglyphic controls sit in the set, so answering hieratic every time doesn't pay.
Best so far 95% on real documents, 0% on the sentence
Which hieroglyph is each sign?
Answered in Gardiner sign-list codes, the standard Egyptologists use. Scored by sign error rate.
Standard Egyptological transliteration. Not scored on the sentence, since no answer is stored.
English. Not scored on the sentence either. Scoring is built in for future public texts.
Every score comes from the open harness, through Anthropic's API for Claude and OpenRouter for everyone else. The first two columns ask a model to name the script, on the professor's sentence and on real ancient documents. Updated 5 October 2026 .
Claude Fable 5.1 (high) declined 6 questions, which count as wrong. Real documents counts hieratic documents only, not the demotic and hieroglyphic controls. Transliterating and translating the sentence aren't scored, because no answer is stored anywhere. Chat-app transcripts above are quoted for the record and never counted here.
268 items. 87 real hieratic documents, 29 controls in other Egyptian scripts, 150 single signs and 2 sealed sentences . Every public image is openly licensed and credited.
Sealed © Aly Moursy. Free to use for evaluating models.
Sealed © Aly Moursy. Free to use for evaluating models.
Papyrus Berlin 3022 (Story of Sinuhe), opening section
Ägyptisches Museum und Papyrussammlung, Berlin
Ostracon with Pharaoh Spearing a Lion and a Royal Hymn on its Back, Met 26.7.1453
ca. 1200–1080 BCE, The Metropolitan Museum of Art, New York
Hieratic CC0 (Met Open Access)
Rhind Mathematical Papyrus (BM EA 10057), detail
c. 1650 BCE, British Museum, London (EA 10057)
Ostracon: letter from the scribe Mose to the vizier Pesiur (Ashmolean HO 71)
13th century BCE, Ashmolean Museum, Oxford (HO 71)
Marriage Contract, Met 35.4.1a, b
380–343 BCE, The Metropolitan Museum of Art, New York
Single hieratic sign, London, British Museum, EA 9999 (AKU HT 10440)
New Kingdom, Dynasty 20, Ramesses III, regnal year 32, month 3, season šm.w, day 6
Help AI learn to read hieratic.
This only works with more people. It needs the few who can read hieratic, and the people building the models that one day might.
AI researchers Run the harness on your model with your own key, or build a reader from scratch. Every result goes on the public leaderboard. Run the benchmark Everyone else Share it. Introduce us to an Egyptologist. Sponsor a commissioned sentence. The benchmark grows one sentence at a time. Get in touch Or go see it for yourself.
Egypt's museums are full of hieratic, on papyrus and on pottery shards. Deir el-Medina, near Luxor, is the village where the workmen who built the royal tombs left thousands of notes in it.
Need a visa? Veeza AI handles the Egypt e-visa before you fly.
The Nile Abu Simbel HieraticBench was started by Aly Moursy , founder of Veeza AI ( YC F26 ), in 2026. Code is MIT licensed and results are CC BY 4.0. Every image is credited on the dataset page .
