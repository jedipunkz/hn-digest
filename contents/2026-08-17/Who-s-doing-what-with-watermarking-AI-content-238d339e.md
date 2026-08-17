---
source: "https://dilpreet.co/blog/2026/8/17/whos-doing-what-with-watermarking-ai-content"
hn_url: "https://news.ycombinator.com/item?id=49329962"
title: "Who's doing what with watermarking AI content?"
article_title: "Who’s doing what with watermarking AI content? — Dilpreet Singh"
image: "http://static1.squarespace.com/static/532e6d5ce4b04d947b0a3249/t/6a770e9f01f73d19939516be/1786187423848/dilpreet-social-logo.jpg?format=1500w"
author: "argilium"
captured_at: "2026-08-17T13:33:25Z"
capture_tool: "hn-digest"
hn_id: 49329962
score: 1
comments: 0
posted_at: "2026-08-17T12:46:03Z"
tags:
  - hacker-news
  - translated
---

# Who's doing what with watermarking AI content?

- HN: [49329962](https://news.ycombinator.com/item?id=49329962)
- Source: [dilpreet.co](https://dilpreet.co/blog/2026/8/17/whos-doing-what-with-watermarking-ai-content)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T12:46:03Z

## Translation

タイトル: AI コンテンツに透かしを入れて誰が何をしているのか?
記事のタイトル: AI コンテンツに透かしを入れて誰が何をしているのか? — ディルプリート・シン
説明: Anthropic は SynthID スタイルのテキスト透かしを Claude に導入しています。John Gruber はいくつかの反対意見を持っていますが、James Padolsey はその仕組みについて私が見た中で最も明確な説明をしています。この議論全体を通して、他の人は何をしているのか疑問に思いました。主要な AI 企業の現在の立場は次のとおりです

記事本文:
AI コンテンツに透かしを入れて誰が何をしているのでしょうか?
Anthropic は SynthID スタイルのテキスト透かしを Claude に導入していますが、John Gruber はいくつかの反対意見を持っており、James Padolsey はその仕組みについて私が見た中で最も明確な説明をしています。この議論全体を通して、他の人は何をしているのか疑問に思いました。
主要な AI 企業の現在の立場は次のとおりです。
Google DeepMind は、最初は生成された画像用に SynthID を開発しました。 2024 年に、SynthID を Gemini アプリと Web エクスペリエンスのテキストに拡張し、その後、 SynthID Text をオープンソース化しました。それが、Anthropic がクロードのバージョンを採用できる理由であり、他の誰でも同様です。
Google の公開 Gemini および SynthID 検証ツールは現在、画像、音声、ビデオを受け入れますが、貼り付けられたテキストは受け入れません。 Google には (もちろん) AI Content Detection API という同様の名前の別の製品があります。これは、ピクセルレベルのアーティファクトやその他の統計的手がかりを使用して、Google およびサードパーティのモデルからの画像を検査します。幅広い名前にもかかわらず、現在のドキュメントには JPEG、PNG、および WebP 入力のみがリストされています。
将来の Claude モデルは、Google の SynthID Text のバージョンを使用する予定です。既存のモデルは今後数か月以内に移行されます。 C2PA メタデータは、サポートされている一部のファイルにも含まれています。
Anthropic は、テキスト検出 API が間もなく登場すると述べています。
OpenAIはテキスト透かしをリリースしていないが、来歴サポートをテキストに拡張するつもりだと述べている。サポートされている画像は C2PA と SynthID を使用し、サポートされているオーディオは SynthID を使用します。
OpenAI Verify は現在、サポートされている画像と音声をチェックしています。
Meta AI または Llama の発表されたテキスト ウォーターマークは見つかりませんでした。 Meta の新しいコンテンツ シールは、生成された画像にマークを付けます。ビデオのサポートも予定されています。
Microsoft はテキストの実装を発表していません。 Microsoft 365 は、生成された画像、ビデオ、オーディオにウォーターマークと来歴メタデータを追加できます

, ただし、一般的なテキストチェッカーはありません。
どちらも EU コードに準拠していますが、私が見つけたテキスト透かしの実装についてはどちらも公に説明していません。
xAI はテキスト透かしについて発表していません。 Grok は生成された画像やビデオに透かしを追加しますが、テキストをチェックするための公開機能は何も提供しません。
この突然の活動は、EU AI法によって部分的に推進されています。その透明性ルールは 2026 年 8 月 2 日に適用され始め、生成されたテキスト、画像、音声、ビデオを技術的に可能な場合には機械可読形式でマークすることを義務付けています。 Anthropic、Google、OpenAI、Meta、Microsoft、Mistral、Cohere はすべて、EU 実施規範の関連セクションに署名しています。
オープンウェイトモデルはどうですか？
Google は SynthID Text をオープンソース化しており、Hugging Face Transformers で利用できます。つまり、 Qwen3-8B などのオープンウェイト モデルをダウンロードして自宅で実行し、独自のキーを選択して透かし入りのテキストを生成できるということです。ただし、SynthID がなくても同じモデルを簡単に実行できます。このテクニックによるオープンウェイトとウォーターマークは実際には役に立ちません。

## Original Extract

Anthropic is bringing SynthID-style text watermarking to Claude , John Gruber has some objections , and James Padolsey has the clearest explanation I’ve seen of how it works . This whole discussion got me wondering what everyone else is doing. Here is where the major AI companies currently stand

Who’s doing what with watermarking AI content?
Anthropic is bringing SynthID-style text watermarking to Claude , John Gruber has some objections , and James Padolsey has the clearest explanation I’ve seen of how it works . This whole discussion got me wondering what everyone else is doing.
Here is where the major AI companies currently stand:
Google DeepMind developed SynthID, initially for generated images. In 2024 it extended SynthID to text in the Gemini app and web experience, then open-sourced SynthID Text . That is why Anthropic can adopt a version for Claude, and so can anyone else.
Google’s public Gemini and SynthID verification tools currently accept images, audio and video, but not pasted text. Google (of course) has another similarly named product called the AI Content Detection API . That one examines images from Google and third-party models using pixel-level artifacts and other statistical clues. Despite the broad name, its current documentation only lists JPEG, PNG and WebP inputs.
Future Claude models will use a version of Google’s SynthID Text. Existing models will move over during the coming months. C2PA metadata is also included with some supported files.
Anthropic says a text-detection API is coming soon .
OpenAI has not released a text watermark, although it says it intends to expand provenance support to text. Supported images use C2PA and SynthID, while supported audio uses SynthID.
OpenAI Verify currently checks supported images and audio.
I couldn’t find an announced text watermark for Meta AI or Llama. Meta’s new Content Seal marks generated images, with video support planned.
Microsoft has not announced a text implementation. Microsoft 365 can add watermarks and provenance metadata to generated images, video and audio , but there is no general text checker.
Both have committed to the EU Code , but neither has publicly described a text-watermarking implementation that I could find.
xAI has not announced a text watermark. Grok adds watermarks to generated images and video, but offers nothing public for checking text.
The sudden activity is being driven in part by the EU AI Act. Its transparency rules began applying on 2 August 2026 and require generated text, images, audio and video to be marked in a machine-readable form where technically feasible. Anthropic, Google, OpenAI, Meta, Microsoft, Mistral and Cohere have all signed the relevant section of the EU code of practice .
What about open-weight models?
Google has open-sourced SynthID Text , and it is available in Hugging Face Transformers. That means I could download an open-weight model such as Qwen3-8B , run it at home, choose my own keys and produce watermarked text. But you can just as easily run the same model without SynthID; open-weights and watermarking via this technique don't really jibe.
