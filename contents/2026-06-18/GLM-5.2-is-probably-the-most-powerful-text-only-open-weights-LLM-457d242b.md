---
source: "https://simonwillison.net/2026/Jun/17/glm-52/"
hn_url: "https://news.ycombinator.com/item?id=48580759"
title: "GLM-5.2 is probably the most powerful text-only open weights LLM"
article_title: "GLM-5.2 is probably the most powerful text-only open weights LLM"
author: "lumpa"
captured_at: "2026-06-18T04:34:28Z"
capture_tool: "hn-digest"
hn_id: 48580759
score: 1
comments: 0
posted_at: "2026-06-18T04:18:39Z"
tags:
  - hacker-news
  - translated
---

# GLM-5.2 is probably the most powerful text-only open weights LLM

- HN: [48580759](https://news.ycombinator.com/item?id=48580759)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/17/glm-52/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T04:18:39Z

## Translation

タイトル: GLM-5.2 はおそらく最も強力なテキスト専用オープンウェイト LLM です
説明: 中国の AI ラボ Z.ai は、6 月 13 日にコーディング プランの加入者に GLM-5.2 をリリースし、その後昨日 (6 月 16 日) に MIT ライセンスに基づいて完全なオープン ウェイトをリリースしました。同様の点は…

記事本文:
GLM-5.2 はおそらく最も強力なテキスト専用オープンウェイト LLM です
サイモン・ウィリソンのウェブログ
GLM-5.2 はおそらく最も強力なテキスト専用オープンウェイト LLM です
中国の AI ラボ Z.ai は、6 月 13 日にコーディング プランの加入者向けに GLM-5.2 をリリースし、昨日 (6 月 16 日) には MIT ライセンスに基づいて完全なオープン ウェイトをリリースしました。以前の GLM-5 および GLM-5.1 リリースと同様のサイズで、これは 753B パラメータ、1.51TB のモンスターで、40 個のアクティブ パラメータを備えています (専門家の混合)。 GLM-5.2 はテキスト入力専用モデルです。Z.ai には、最近では GLM-5V-Turbo に代表される別のビジョン ファミリがありますが、これはオープン ウェイトではありません。 GLM-5.2 には、GLM-5.1 の 200,000 トークン コンテキスト ウィンドウから増加した 100 万トークン コンテキスト ウィンドウがあります。
このモデルに関する噂は根強い。
Artificial Analysis は、最も広く評価されている独立ベンチマークの 1 つを実行しています。GLM-5.2 は、Artificial Analysis Intelligence Index の新しい主要なオープン ウェイト モデルです。
GLM-5.2 は、Intelligence Index v4.1 の主要なオープン ウェイト モデルです。 51 では、MiniMax-M3 (44)、DeepSeek V4 Pro (最大 44)、および Kim K2.6 (43) を上回ります。
しかし、彼らはそれが非常にトークンを大量に消費することを発見しました。
GLM-5.2 は、他の主要なオープンウェイト モデルよりもタスクごとに多くの出力トークンを使用します。このモデルは、インテリジェンス インデックス タスクごとに 43,000 の出力トークンを使用します。これは、GLM-5.1 (26,000) 以上の MiniMax-M3 (24,000)、Kimi K2.6 (35,000)、および DeepSeek V4 Pro (最大 37,000) よりも増加しています。
このモデルは現在、Code Arena WebDev リーダーボードでも、Claude Fable 5 に次いで 2 位にランクされています。そのリーダーボードは、「エージェント コーディング ワークフローを含むフロントエンド Web 開発タスク」を測定します。画像入力がないにもかかわらず、これが非常に高いランクにランクされているのを見て感銘を受けました。画像入力は、本当に優れたフロントエンド コーディング モデルを構築するための重要な部分であると誤って思い込んでいたのです。
OpenRouter 経由で試してみました。

9 つの異なるプロバイダーから提供されており、ほとんどすべてのプロバイダーでは、入力に 100 万あたり 140 ドル、出力に 100 万あたり 440 ドルの料金がかかります。比較のために、GPT-5.5 は $5/$30、Claude Opus 4.5-4.8 は $5/$25 です。
GLM-5.1 は、私のお気に入りのペリカンと、ずっとお気に入りのオポッサムを提供してくれました (「電動スクーターで北バージニア州のオポッサムの SVG を生成する」というプロンプトに対して)。興味深いことに、どちらの場合でも、モデルは、CSS を使用して追加のアニメーションを追加した HTML ドキュメントにラップされた SVG を返すことを選択しました。
GLM-5.2を試してみましょう。 「自転車に乗っているペリカンのSVGを生成する」については、次のようになりました。
これは自己完結型の完全にアニメーション化された SVG であり、アニメーションは壊れていません。目が落ちたり、車輪が自転車とは関係なく回転したりするのをよく見かけますが、ここではすべてがうまく機能しています。ペリカンのとても素敵なベクターイラストです。とても印象的です。
残念なことに、電動スクーターに乗ったノースバージニアオポッサムは、それほどうまくいきませんでした。
これは GLM-5.1 から大幅にダウンしています。ちなみに、ポッサムはこんな感じでした。
5.2 ではアニメーション化さえしませんでした。
Pyodide で使用するために WASM ホイールを PyPI に公開 - 2026 年 6 月 13 日
クロード・ファブルは容赦なく積極的です - 2026 年 6 月 11 日
これは、2026 年 6 月 17 日に投稿された Simon Willison による GLM-5.2 はおそらく最も強力なテキスト専用オープン ウェイト LLM です。
Previous: Pyodide で使用するために WASM ホイールを PyPI に公開する
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

Chinese AI lab Z.ai released GLM-5.2 to their coding plan subscribers on June 13th, and then yesterday (June 16th) released the full open weights under an MIT license. Similar in …

GLM-5.2 is probably the most powerful text-only open weights LLM
Simon Willison’s Weblog
GLM-5.2 is probably the most powerful text-only open weights LLM
Chinese AI lab Z.ai released GLM-5.2 to their coding plan subscribers on June 13th, and then yesterday (June 16th) released the full open weights under an MIT license. Similar in size to their previous GLM-5 and GLM-5.1 releases, this is 753B parameter, 1.51TB monster—with 40 active parameters (Mixture of Experts). GLM-5.2 is a text input only model—Z.ai have a separate vision family most recently represented by GLM-5V-Turbo , but that one isn’t open weights. GLM-5.2 has a 1 million token context window, up from GLM-5.1’s 200,000.
The buzz around this model is strong.
Artificial Analysis, who run one of the most widely respected independent benchmarks: GLM-5.2 is the new leading open weights model on the Artificial Analysis Intelligence Index .
GLM-5.2 is the leading open weights model on the Intelligence Index v4.1. At 51, it leads MiniMax-M3 (44), DeepSeek V4 Pro (max, 44) and Kimi K2.6 (43)
They did however find it to be quite token-hungry:
GLM-5.2 uses more output tokens per task than other leading open weights models: the model uses 43k output tokens per Intelligence Index task, up from GLM-5.1 (26k) and above MiniMax-M3 (24k), Kimi K2.6 (35k) and DeepSeek V4 Pro (max, 37k)
The model is also now ranked 2nd on the Code Arena WebDev leaderboard , behind only Claude Fable 5. That leaderboard measures “front-end web development tasks, including agentic coding workflows”. I’m impressed to see it rank so highly given the lack of image input, which I had incorrectly assumed was a key part of building a truly great frontend coding model.
I’ve been trying it out via OpenRouter , which has it from 9 different providers, almost all of which are charging $1.40/million for input and $4.40/million for output. For comparison, GPT-5.5 is $5/$30 and Claude Opus 4.5-4.8 is $5/$25.
GLM-5.1 gave me one of my favorite pelicans and my all time favorite opossum (for the prompt “Generate an SVG of a NORTH VIRGINIA OPOSSUM ON AN E-SCOOTER”.) Interestingly, in both of those cases the model chose to return SVG wrapped in an HTML document that added additional animations using CSS.
Let’s try GLM-5.2. For “Generate an SVG of a pelican riding a bicycle” I got this :
It’s a self-contained fully animated SVG, and the animations aren’t broken! Often I’ll see eyes falling off or wheels rotating independently of the bicycle but here everything works great. It’s a very nice vector illustration of a pelican too. Very impressive.
Sadly, the NORTH VIRGINIA OPOSSUM ON AN E-SCOOTER did not come out nearly as well :
This is such a step down from GLM-5.1! As a reminder, that possum looked like this:
5.2 didn’t even try to animate it.
Publishing WASM wheels to PyPI for use with Pyodide - 13th June 2026
Claude Fable is relentlessly proactive - 11th June 2026
This is GLM-5.2 is probably the most powerful text-only open weights LLM by Simon Willison, posted on 17th June 2026 .
Previous: Publishing WASM wheels to PyPI for use with Pyodide
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.
