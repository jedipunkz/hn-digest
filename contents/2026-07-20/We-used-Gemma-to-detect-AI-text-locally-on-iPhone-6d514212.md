---
source: "https://imbue.com/blog/bouncer-leveraging-local-compute-to-detect-ai-slop"
hn_url: "https://news.ycombinator.com/item?id=48981692"
title: "We used Gemma to detect AI text locally on iPhone"
article_title: "Bouncer: Leveraging Local Compute to Detect AI Slop - Imbue"
author: "jjleads"
captured_at: "2026-07-20T17:23:21Z"
capture_tool: "hn-digest"
hn_id: 48981692
score: 4
comments: 0
posted_at: "2026-07-20T17:12:37Z"
tags:
  - hacker-news
  - translated
---

# We used Gemma to detect AI text locally on iPhone

- HN: [48981692](https://news.ycombinator.com/item?id=48981692)
- Source: [imbue.com](https://imbue.com/blog/bouncer-leveraging-local-compute-to-detect-ai-slop)
- Score: 4
- Comments: 0
- Posted: 2026-07-20T17:12:37Z

## Translation

タイトル: Gemma を使用して iPhone 上でローカルに AI テキストを検出しました
記事のタイトル: Bouncer: ローカル コンピューティングを活用して AI スロップを検出する - Imbue

記事本文:
Bouncer: ローカル コンピューティングを活用して AI スロップを検出 - Imbue
/ ブログ 製品 バウンサー
記事 / Product Bouncer: ローカル コンピューティングを活用して AI のスロップを検出する
Bouncer で AI テキストと画像検出を発表できることを嬉しく思います。多くの Bouncer ユーザーが、フィルター用語として「AI スロップ」、「AI ビデオ」、または「AI 書き込み」を追加して、フィードから AI コンテンツを除外しようとしているのを目にしてきました。これにより、バックエンドの LLM が投稿に関連付けられたテキストと画像を調べ、AI によって生成されたものであるかどうかを分類するように促されます。
ただし、このアプローチには問題があります。LLM にこの判断を促すだけでは不正確すぎるのです。テキストの場合、以前のベンチマークでは、特に小規模なモデルの場合、このようなゼロショット プロンプト手法がランダムよりかろうじて優れていることが示されています。 「AI が書いたもの」をフィルタリングすると、AI が書いたものであることを示すマーカー (エムダッシュ、「x ではありません、y です」など) を含む一部の投稿が検出される可能性がありますが、ほとんどの場合、この方法は失敗し、誤検知が発生する傾向があります。同様に、LLM は、画像を見ただけでは、その画像が AI によって生成されたものであるかどうかを確実に判断できません。
信頼性の高い検出には、タスク専用にトレーニングされたモデルが必要です。これに対処するために、テキストと画像の両方に専用の AI 検出モデルを構築しました。 Bouncer に移動し、「AI スロップを削除」をクリックするだけでフィルタリングを開始できます。
この機能を使用すると、次のコンテンツを検出して削除できます。
検出器のトレーニングはデータセットから始まります。 Amazon レビュー、Yelp レビュー、Reddit、FineWeb-EDU、ニュース記事、Twitter から人間が書いた 28,000 件のテキストを収集しました。次に、Claude Sonnet 4.6、Gemini 3 Flash、および GPT 5.3 にそれらのテキストを編集するように指示し (「これをもっと簡潔に」、「この音をより洗練させて」など)、AI 編集された一連のテキストを生成しました。最後に、同じモデルにゼロからテキストを生成するよう指示しました。

各ドメインとトピック (「炊飯器の Amazon レビューを生成する」) について。その結果、データセットは人間によるテキスト、AI によって編集されたテキスト、および AI によって生成されたテキストに均等に分割されます。世の中の投稿は必ずしも純粋に人間や純粋に AI であるとは限らないため、AI で編集されたテキストも含めました。投稿には両方が混在していたり​​、LLM を通じて元々人間が書いたテキストを言い換えたりする場合があり、私たちのモデルはこれらのケースも捉える必要があります。
アーキテクチャは比較的標準的です。投稿を分類するには、まずそのテキストを微調整された LLM に渡し、最後の入力トークンの非表示状態を取得します。この隠れた状態は、モデルが入力テキストについて学習したすべてを表す高次元ベクトルです。通常、これは次のトークンを生成するために使用されますが、私たちの場合は、テキストに AI 支援のレベルを割り当てる分類ヘッドを介してトークンを供給します。
LLM バックボーンとして Qwen 3 4B を使用することで、人間によるテキストと AI によって生成されたテキストを区別する際にテスト セットで 99.8% の精度を達成し、テキストを人間によるテキスト、AI によって編集されたテキスト、または AI によって生成されたテキストに分類する際に 91.7% の精度を達成しました。予想通り、混乱の大部分は AI で編集されたケースで発生します。
また、サードパーティのベンチマークでも検出器を評価しました。 RAID データセットに含まれるモデルでトレーニングが行われていないにもかかわらず、誤検知率 0.10% で 96.4% の精度を達成しています。
このモデルは、サーバー上の AI テキスト検出を強化します。しかし、私たちはさらに先に進みたかったのです。
Imue では、パーソナル コンピューティングに力を入れています。可能であれば、ユーザーが自分のデバイスでタスクを実行できるようにしたいと考えています。これが、当社の iOS アプリがデバイス上で AI テキスト検出を完全に提供する理由です。
強力なテキスト検出には 4B パラメータ バックボーンで十分であることが判明しました。 4 ビット量子化では、メモリの制約を考慮した最新の iOS デバイスが実行できる範囲のちょうど頂点に位置します。しかし、複雑な問題があります。

iOS バージョンの Bouncer はすでに Gemma 4 E2B を使用して投稿をフィルター カテゴリ別に分類しています。ユーザーは通常のフィルタリングと AI テキスト検出をローカルで利用するために Gemma 4 E2B と Qwen 3 4B の両方を合理的にダウンロードすることはできず、いずれにしても 2 つのモデルは同時にメモリに収まりません。解決策は、両方のタスクで 1 つの LLM バックボーンを共有することです。
両方のタスクをサポートするために、条件付きで LoRA アダプターを Gemma 4 E2B に適用します。通常のテキスト生成フローでは、アダプターなしで基本モデルが実行されます。 AI 検出フローはアダプターを適用し、最後のトークンの非表示状態を取得して、それを小さな分類ヘッドに渡します。 LoRA アダプターはわずか 10MB なので、ダウンロードが速く、メモリを節約できます。
これを実装するには、LiteRT-LM ランタイムと、ランタイムに必要な形式でモデルをエクスポートする LiteRT-Torch にいくつかの変更を加える必要がありました。私たちが遭遇した問題の 1 つは、LiteRT-Torch がモバイルに最適化された wNa8o8 スキーマへのエクスポートをサポートしていないことです。これは、Google の iOS 用事前構築モデルがメモリ使用量を削減するために使用するスキーマです。非表示の状態を出力として公開するには、モデルを新たにエクスポートする必要があります。この新しいスキーマをフォークに追加することで、Google が事前に構築した重みと少し同じ重みを持つモデルを生成し、通常のテキスト生成パスのモデル出力品質を維持することができます。
これを実装した後、LoRA アダプター、制約付きデコード、およびプレフィックス キャッシュのアプリケーションをサポートする機能を LiteRT-LM ランタイムに追加しました。結果として得られたオンデバイス モデルは、AI と人間のテキストを区別する場合にテスト セットで 99.6% の精度を達成し、人間のテキスト、AI 編集されたテキスト、および AI が生成したテキストを区別する場合に 88.2% の精度を達成しました。
このモデルの精度は、LLM バックボーンが小さいため、ホストされているバージョンよりわずかに悪くなります。

これは、デバイス上で完全に実行するための許容可能なトレードオフです。
ローカル AI テキスト検出も高速です。分類に必要なのは、自己回帰デコードなしで単一の前方パスのみであるためです。 iPhone 18 では、ツイートの平均テキスト検出リクエストにかかる時間はわずか 120 ミリ秒です。
サーバー モデルとオンデバイス モデルの両方のトレーニング コードと評価は、ここでオープン ソースです。 iOS デバイス上で独自のモデルを実行したい人向けに、モデルのエクスポート フローも提供します。
画像検出は非常に似た方法で構築されています。私たちのモデルは、NTIRE の 2026 年の AI 画像検出チャレンジで最高のパフォーマンスを発揮したソリューションから派生した、分類ヘッドを備えた微調整された DINOv3 バックボーンです。トレーニング データはさまざまな AI 画像モデルから生成され、トリミング、ぼかし、圧縮、および画像がオンラインで一般的に行われるその他の変換に対する堅牢性を向上させるために、さまざまな拡張を適用しています。 DINOv3 の 70 億パラメータ バージョンを使用していますが、これは一般消費者向けデバイスでは大きすぎて適切に処理できないため、現時点では画像検出はサーバー側でのみ実行されます。将来的には、この機能をローカルのユーザーに提供するために、より小規模な ViT バックボーンを使用したトレーニングを追求する可能性があります。
両方の機能は現在、次のプラットフォームで利用できる Bouncer で利用できます。
2026 年 4 月 8 日 -- Product Bouncer: フィードを修復する

## Original Extract

Bouncer: Leveraging Local Compute to Detect AI Slop - Imbue
/ blog Products Bouncer
Article / Product Bouncer: Leveraging Local Compute to Detect AI Slop
We're excited to announce AI text and image detection in Bouncer ! We've been seeing a lot of Bouncer users try to filter out AI content from their feeds by adding "AI slop", "AI video", or "AI written" as filter terms. This prompts an LLM in our backend to look at the text and images associated with a post and classify it as AI-generated or not.
However, there’s an issue with this approach: simply prompting an LLM to make this judgment is too inaccurate. For text, previous benchmarks have shown that zero-shot prompting methods like this are barely better than random, especially for smaller models. While filtering for "AI written" may catch some posts with telltale markers of AI writing (em-dashes, "it's not x, it's y", etc.), for the most part this method fails and is prone to false positives. Similarly, an LLM cannot reliably tell whether an image is AI-generated just by looking at it.
Reliable detection requires models trained specifically for the task. To address this, we've built dedicated AI detection models for both text and images. Simply go into Bouncer and click “Remove AI Slop” to start filtering.
With this feature, the following content can be detected and removed:
Training a detector starts with a dataset. We collected 28,000 human-written texts from Amazon reviews, Yelp reviews, Reddit, FineWeb-EDU, news articles, and Twitter. We then prompted Claude Sonnet 4.6, Gemini 3 Flash, and GPT 5.3 to edit those texts ("Make this more concise", "Make this sound more sophisticated", etc.), producing a set of AI-edited text. Finally, we prompted the same models to generate text from scratch based on each domain and topic ("Generate an Amazon review for a rice cooker"). The result is a dataset evenly split between human, AI-edited, and AI-generated text. We included AI-edited text as well since posts in the wild are not necessarily purely human or purely AI. A post might mix both, or paraphrase originally human text through an LLM, and our model needs to catch these cases too.
The architecture is relatively standard. To classify a post, we first pass its text through a fine-tuned LLM to get the hidden state of the last input token. This hidden state is a high dimensional vector that represents everything the model has learned about our input text. Typically it’s used to generate the next token, but in our case we feed it through a classification head that assigns the text a level of AI assistance.
Using Qwen 3 4B as our LLM backbone we achieve 99.8% accuracy on our test set when discriminating between human and AI-generated text, and 91.7% accuracy when classifying text as human, AI-edited, or AI-generated. As expected, the majority of confusion occurs in the AI-edited case.
We also evaluated the detector on third party benchmarks. Despite training on none of the models represented in the RAID dataset, it achieves 96.4% accuracy with a false positive rate of 0.10%.
This model powers AI text detection on our servers. But we wanted to go further.
At Imbue, we are committed to personal computing. Where possible, we want users to be able to run tasks on their own devices. This is why our iOS app offers AI text detection fully on device.
A 4B parameter backbone turned out to be enough for strong text detection. With 4-bit quantization, that sits right on the cusp of what the newest iOS devices can run given their memory constraints. But there's a complication: the iOS version of Bouncer already uses Gemma 4 E2B to classify posts by filter categories. Users can't reasonably download both Gemma 4 E2B and Qwen 3 4B to get normal filtering and AI text detection locally, and the two models wouldn't fit in memory simultaneously anyway. The solution is to share one LLM backbone across both tasks.
To support both tasks, we conditionally apply a LoRA adapter to Gemma 4 E2B. The normal text generation flow runs the base model without the adapter. The AI detection flow applies the adapter, takes the last token's hidden state, and passes it through a small classification head. The LoRA adapter is just 10MB which makes it quick to download and light on memory.
To implement this, we had to make some changes to the LiteRT-LM runtime and LiteRT-Torch , which exports models to the format required by the runtime. One of the issues we ran into is that LiteRT-Torch did not support exporting to the mobile optimized wNa8o8 schema. This is the schema that Google’s prebuilt model for iOS uses in order to reduce its memory footprint. In order to expose hidden states as output, a new export of the model is required. By adding this new schema in our fork, we are able to produce a model with bit identical weights to Google’s prebuilt, preserving model output quality for the normal text generation path.
After implementing this, we added functionality to the LiteRT-LM runtime to support the application of LoRA adapters, constrained decoding, and prefix caching. The resulting on-device model achieves 99.6% accuracy on our test set when discriminating between AI and human text, and 88.2% when discriminating between human, AI-edited, and AI-generated text.
The accuracy of this model is slightly worse than the hosted version due to its smaller LLM backbone, but this is an acceptable tradeoff for running fully on device.
Local AI text detection is also fast, since classification requires only a single forward pass with no autoregressive decoding. The average text detection request for a tweet takes just 120 milliseconds on an iPhone 18.
The training code and evals for both the server and on-device models are open source here . We also provide the model export flow for those wanting to run their own models on iOS devices.
Image detection is built in a very similar way. Our model is a fine-tuned DINOv3 backbone with a classification head, derived from top-performing solutions in NTIRE's 2026 AI image detection challenge . The training data is generated from a variety of AI image models, and we apply many different augmentations to improve robustness to cropping, blurring, compression, and other transformations that images commonly undergo online. We use the 7 billion parameter version of DINOv3, which is too large for consumer devices to handle well so image detection runs server-side only for now. In the future we may pursue training with a smaller ViT backbone in order to bring this feature to users locally.
Both features are available in Bouncer today which is available on the following platforms:
08 Apr 2026 -- Product Bouncer: Heal your feed
