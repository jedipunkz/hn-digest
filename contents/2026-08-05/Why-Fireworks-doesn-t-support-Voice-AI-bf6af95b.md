---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49186754"
title: "Why Fireworks doesn't support Voice AI"
article_title: ""
author: "kushalpatil07"
captured_at: "2026-08-05T19:15:23Z"
capture_tool: "hn-digest"
hn_id: 49186754
score: 2
comments: 0
posted_at: "2026-08-05T18:18:27Z"
tags:
  - hacker-news
  - translated
---

# Why Fireworks doesn't support Voice AI

- HN: [49186754](https://news.ycombinator.com/item?id=49186754)
- Score: 2
- Comments: 0
- Posted: 2026-08-05T18:18:27Z

## Translation

タイトル: Fireworks が音声 AI をサポートしない理由
HN テキスト: 花火が音声モデルをサポートしていない理由を考え始めました。インコ、ココロ、Qwen ASR など、現在は非常に優れたオープンソース モデルが利用可能ですが、多数の GPU を自分で管理せずにそれを使用する方法はありません。音声エージェントで使用される Gemma 4 などの LLM もサポートされていません。そこで、使用しているユースケースの種類に応じて、推論プラットフォームを異なる方法で最適化する必要があると考えました。 STT や TTS ではなく、LLM の例を見てみましょう。
- コーディング エージェント -> キャッシュされた入力が大量にあるため、KV キャッシュ用に最適化する必要がある
- スライド/ブログの作成 -> 大量の出力、推測的なデコード用に最適化する必要がある
- 音声 LLM -> キャッシュされた入力が小さく、出力が小さいため、これを最適化する方法がまだわかっていません。したがって、TTS と STT はまったく異なる球技です。タイミングはわかりませんが、人々は今、kokoro、インコ、Qwen などのオープンソース モデルを使いたいと思っているのでしょうか?

## Original Extract

I started thinking over why doesn't fireworks support voice models. There are really good opensource models available now, like parakeet, kokoro, Qwen ASR etc but no way to use it without managing a bunch of GPUs yourself. Even LLMs like Gemma 4 used by voice agents are not supported. Then I figured that the inference platform needs to be optimized differently for the kind of usecase you are using. Lets take an example for LLMs, not even STT and TTS.
- Coding agents -> lot of cached input, needs to optimize for KV cache
- Creation slides/blogs -> lots of output, needs to optimize for speculative decoding
- Voice LLMs -> Cached input small output, not yet figured out on how to optimize this. So TTS and STT is a completely different ballgame. What I don't know is the timing, do people want to use open source models like kokoro, parakeet, Qwen etc RIGHT NOW?

