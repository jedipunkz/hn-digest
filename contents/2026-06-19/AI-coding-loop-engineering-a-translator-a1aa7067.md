---
source: "https://jeena.net/loop-engineering"
hn_url: "https://news.ycombinator.com/item?id=48594371"
title: "AI coding: loop engineering a translator"
article_title: "AI coding: loop engineering a translator; I accidentally did that years ago - jeena.net"
author: "jeena"
captured_at: "2026-06-19T04:36:45Z"
capture_tool: "hn-digest"
hn_id: 48594371
score: 2
comments: 0
posted_at: "2026-06-19T03:00:50Z"
tags:
  - hacker-news
  - translated
---

# AI coding: loop engineering a translator

- HN: [48594371](https://news.ycombinator.com/item?id=48594371)
- Source: [jeena.net](https://jeena.net/loop-engineering)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T03:00:50Z

## Translation

タイトル: AI コーディング: トランスレータのループ エンジニアリング
記事のタイトル: AI コーディング: トランスレータのループ エンジニアリング。何年も前にうっかりやってしまいました - jeena.net
説明: 役に立ちましたか?

記事本文:
ホーム /
ブログ
/
ポッドキャスト /
メモ
/
写真
/
について/
もっと見る
AI コーディング: トランスレータのループ エンジニアリング
何年も前に偶然やってしまった
投稿しました
約3時間前
2026-06-19 03:00:30
によって
ジーナ
ここ 1 ～ 2 週間で、「プロンプト エンジニアリング」や「エージェント ワークフロー」ではもう通用しないようになった後、「ループ エンジニアリング」という用語が登場するようになり、さらなる抽象化が必要です。
私が初めて LLM を使って役に立つ何かを書こうとしたのは、約 2 年前の 2024 年 12 月でした。12 GB VRAM を搭載したローカルの RTX 3060 で、大きなドキュメントを韓国語から英語に翻訳したいと思いました。 ChatGPT チャットをそのまま使用することはできませんでした。当時は、翻訳が必要な大量のドキュメントに対して ChatGPT のコンテキスト ウィンドウが小さすぎて、10% を超えると翻訳が停止してしまいました。そこで、ドキュメントの分割を自動化できないかと考えました。
しかし、私はエージェントについて何も知らなかったので、いくつかのアーキテクチャを考え出す必要がありました。そこで、次のような非常に複雑なエージェント パイプラインを作成しました。
このパイプラインは、計画→実行→批評→修復のループを通じて、生の韓国語トランスクリプトを英語翻訳に変換します。リテラル NLLB 翻訳は公平な参照として機能し、翻訳メモリはチャンク全体で用語の一貫性を保ちます。
何時間にもわたる教育用 YouTube ビデオのトランスクリプトを翻訳する必要があったため、数週間この作業に取り組みました。最初は単純なループでしたが、ローカル モデルの品質が良くなかったので、批評家を導入し、そこからの情報をプロンプトとしてフィードバックすることで翻訳の品質を向上させようとしました。その後、物事が漂っていることに気付き、同じ単語が文書の途中で異なって翻訳されていることに気づきました。そこで、漂流しないようにメモリを導入しました。重複するウィンドウなど、さらに多くの小さな最適化を行いました。

チャンクなど
しかし最終的に、批評家からのあらゆる情報と記憶がより良い翻訳につながることはなく、何が起こったかというと、批評家は翻訳が十分ではないというフラグを立て続けてループバックし、翻訳者は批評家が満足するほど良い翻訳をすることができませんでした。翻訳を大幅に改善するような微調整ができなかったので、数週間経っても諦めて、より良いローカル翻訳モデルが到着するのを待っていますが、そうなればこのループ全体はおそらくもう必要なくなるでしょう。
返事を書きましたか？ URLを教えてください:
インディーズ コメント (ウェブメンション) のサポートもあります。

## Original Extract

Did it help?

home /
blog
/
podcasts /
notes
/
photos
/
about /
more
AI coding: loop engineering a translator
I accidentally did that years ago
Posted
about 3 hours ago
2026-06-19 03:00:30
by
Jeena
Since a week or two the term 'loop engineering' is coming up after 'prompt engineering' and 'agentic workflow' seem to not cut it anymore and we need more abstractions.
The first time I tried to write something which would do something useful with a LLM was about two years ago in December 2024, I wanted to do translation of big documents from Korean to English on my local RTX 3060 with 12 GB VRAM. I couldn't just use the ChatGPT chat for that because back then it's context window was way too small for the massive documents I needed translated and it would stop translating after 10%. So I thought I can automate the splitting up of the document.
But because I did not know anything about agents I had to come up with some architecture. So I made this very complicated agent pipeline:
The pipeline turns a raw Korean transcript into an English translation through a plan → execute → critique → repair loop. A literal NLLB translation acts as an impartial reference, and a translation memory keeps terminology consistent across chunks.
I worked on this for some weeks because I had a real need to translate hours of educational YouTube video transcripts. It started as a simple loop but the quality of the local models was not good so I tried to improve the quality of the translation by introducing a critic and feed back the information from it as a prompt, then I realized things are drifting and the same words are being translated differently over the course of the document so I introduced a memory so it would not drift. I did many more small optimizations like a overlapping window of the chunks, etc.
But in the end all the information from the critic and the memory did not lead to a better translation, what happened was that the critic kept flagging that the translation is not good enough and looping back and the translator was not able to translate good enough for the critic to be satisfied. I was not able to fine tune it so that it would improve the translation in any significant manner and after a couple of weeks I kind of gave up and am waiting for better local translation models to arrive, but then this whole loop thing will probably not be necessary anymore.
Have you written a response? Let me know the URL:
There's also indie comments (webmentions) support.
