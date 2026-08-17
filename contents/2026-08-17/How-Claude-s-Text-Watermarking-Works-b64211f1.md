---
source: "https://sebastianraschka.com/blog/2026/claude-text-watermarking.html"
hn_url: "https://news.ycombinator.com/item?id=49331004"
title: "How Claude's Text Watermarking Works"
article_title: "How Claude's Text Watermarking Works | Sebastian Raschka, PhD"
image: "https://sebastianraschka.com/images/blog/2026/claude-watermarking-explained/claude-watermarking.png"
author: "ModelForge"
captured_at: "2026-08-17T14:18:48Z"
capture_tool: "hn-digest"
hn_id: 49331004
score: 3
comments: 0
posted_at: "2026-08-17T13:58:08Z"
tags:
  - hacker-news
  - translated
---

# How Claude's Text Watermarking Works

- HN: [49331004](https://news.ycombinator.com/item?id=49331004)
- Source: [sebastianraschka.com](https://sebastianraschka.com/blog/2026/claude-text-watermarking.html)
- Score: 3
- Comments: 0
- Posted: 2026-08-17T13:58:08Z

## Translation

タイトル: Claude のテキスト透かしの仕組み
記事のタイトル: Claude のテキスト透かしの仕組み |セバスチャン・ラシュカ博士
説明: Anthropic がリリースした資料に基づいて、クロードのテキストの透かしがどのように機能するかを示す短い図。

記事本文:
メインコンテンツにスキップ
セバスチャン
ラシュカ
-->
🌙
このウェブサイトを検索
検索を送信
ホーム
ブログ
ブログ -->
本
AI ニュースレター -->
コース
LLM ギャラリー
ゼロからの LLM
推論モデル
会談
もっと見る
ブログアーカイブ
クイックペーパーとモデルノート
研究
について
お問い合わせ -->
詳細 -->
クロードのテキスト透かしの仕組み
2026 年 8 月 15 日
セバスチャン・ラシュカ著
クロードの透かしがどのように機能するかを示す短い図 (リリースされた資料を読んだことに基づく)。
一般に、トークンを生成するとき、特定の次の単語の位置に複数の高スコアのトークンが存在する可能性があります。通常、top-k または top-p サンプリングを使用してサンプリングするため、最もスコアの高いトークンが選択されます (サンプリングを何度も繰り返す場合) が、他のトークンも選択される可能性があります。
ウォーターマークには、最もスコアの高いトークンのどれを選択するかに影響を与えるキーがあります。 (透かしキー + 最近のトークン コンテキスト → 擬似乱数シード → 候補トークンの擬似乱数スコア → 有利なスコアを持つ候補に偏ったサンプリング手順。)
または、より具体的には、秘密キーと以前のトークン (のウィンドウ) がここでのランダム性に影響します。これを多くのトークン位置で繰り返すと、統計的に他の方法では (組み合わせ論により) 得られる可能性が低いパターン (統計的相関) となるため、ウォーターマークが作成されます。
もちろん、この透かしは、中程度から重度の編集や、別の非透かし LLM による言い換えによって削除できます。ただし、単語を入れ替えるテキストにいくつかの変更を加える必要があるため、これによりテキストが悪化する可能性があります (透かしの入った位置がわからないため、優れた透かし解除では多くの位置を編集する必要があります)。
私が混乱していることの 1 つは、彼らは基本的に、全員に対してこれを行わなければならないと言っていることです。

EUの規制のため。なぜ？もちろんですが、これは別のモデルの再トレーニングやトレーニングを必要としない推論時の手法なので、その気になれば EU ユーザーに対してのみ実行できるのでしょうか?
これが EU ユーザーだけに適用されない理由についての質問に答えるには、次のようにします。
Anthropic はプロバイダーであり、EU 市場で Claude を提供し (第 2 条)、その出力は EU に拠点を置く可能性があるため、EU AI 法の対象となります。つまり、透明性規制（第 50 条）を遵守する義務がプロバイダーにあります。生成時点では、Anthropic は出力が最終的にどこに行くのかを知ることができません。法律を遵守するための信頼できる唯一の方法は、生成中にテキストにマークを付けること、そしてそれをユーザーや場所ごとにセグメント化するのではなくグローバルに実行することです。
しかし、一般的にこの規制は非常に厳しいと思います。たとえば、誰かがまだ承認されていないEUに医薬品を輸出する可能性があるため、医薬品メーカーがその医薬品を製造して米国の患者に提供できないというシナリオを考えています。つまり、この場合、メーカーではなく輸出業者が責任を負うべきではないでしょうか？
出典: 私の Substack ノートの Web サイト版。

## Original Extract

Short illustration of how Claude’s text watermarking is supposed to work based on Anthropic’s released materials.

Skip to main content
Sebastian
Raschka
-->
🌙
Search this website
Submit search
Home
Blog
Blog -->
Books
AI Newsletter -->
Courses
LLM Gallery
LLMs From Scratch
Reasoning Models
Talks
More
Blog Archive
Quick Paper and Model Notes
Research
About
Contact -->
More -->
How Claude's Text Watermarking Works
Aug 15, 2026
by Sebastian Raschka
A short illustration of how Claude’s watermarking is supposed to work (based on my reading of their released materials ).
In general, when we are generating tokens, there can be multiple high-scoring tokens at certain next-word positions. Usually, we sample with top-k or top-p sampling so the highest-scoring token is most often selected (if we repeat the sampling many times), but other tokens may be selected as well.
With watermarking, there is a key that influences which of the highest-scoring tokens to select. (Watermark key + recent token context → pseudorandom seed → pseudorandom scores for candidate tokens → sampling procedure biased toward candidates with favorable scores.)
Or, more concretely, the secret key and (a window of) previous token(s) influence the randomness here. Now, if we repeat this at many token positions, this creates the watermark, as it will be a pattern (statistical correlation) that is statistically unlikely to get otherwise (due to combinatorics).
Of course, this watermark can be removed by moderate to severe editing and rephrasing via a different, non-watermarking LLM. But this could potentially make the text worse, as it would require making several changes to the text where it swaps out words (since we don’t know the watermarked positions, a good de-watermarker would have to edit many positions).
One thing I am confused about: They basically say that they HAVE to do this for everyone due to EU regulation. Why? Sure, but this is an inference-time technique that doesn’t require retraining or training a separate model, so if they wanted, they could only do that for EU users?
To answer your question about why this does not just apply to EU users:
Anthropic is a provider and falls under EU AI Act because it offers Claude on the EU Market (Article 2) and its outputs may be based in the EU. That means duty to comply with transparency regulation (art 50) is on the provider. At the time of generation, Anthropic can’t know where the outputs will end up. The only reliable way to comply with the law is to mark text during generation, and to do it globally rather than segmenting by user or location.
In general though, I find this regulation very strict. E.g., I am thinking of a scenario where a pharmaceutical drug manufacturer can’t make and offer said pharmaceutical drug to patients in the US because someone could export it to the EU where it’s not approved yet. I.e., it shouldn’t the exporter be held liable rather than the manufacturer in this case?
Source: website version of my Substack note .
