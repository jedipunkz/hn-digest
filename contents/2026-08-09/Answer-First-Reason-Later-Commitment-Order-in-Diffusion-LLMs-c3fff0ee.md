---
source: "https://arxiv.org/abs/2608.05687"
hn_url: "https://news.ycombinator.com/item?id=49227083"
title: "Answer First, Reason Later: Commitment Order in Diffusion LLMs"
article_title: "[2608.05687] Answer First, Reason Later: Commitment Order in Diffusion LLMs"
author: "sbulaev"
captured_at: "2026-08-09T00:56:24Z"
capture_tool: "hn-digest"
hn_id: 49227083
score: 2
comments: 0
posted_at: "2026-08-09T00:07:07Z"
tags:
  - hacker-news
  - translated
---

# Answer First, Reason Later: Commitment Order in Diffusion LLMs

- HN: [49227083](https://news.ycombinator.com/item?id=49227083)
- Source: [arxiv.org](https://arxiv.org/abs/2608.05687)
- Score: 2
- Comments: 0
- Posted: 2026-08-09T00:07:07Z

## Translation

タイトル: 答えが先、理由は後: 拡散 LLM のコミットメント注文
記事のタイトル: [2608.05687] 答えが先、理由は後で: 拡散 LLM のコミットメント命令
説明: arXiv 論文 2608.05687 の要約ページ: まず答え、後で理由: 拡散 LLM のコミットメント順序

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2026 年 8 月 6 日に提出]
タイトル: 答えが先、理由は後: 拡散 LLM のコミットメント注文
要約: マスクされた拡散言語モデル (dLLM) は、任意の順序でトークンをコミットできます。この自由度は、自己回帰デコードに対する主要な利点として宣伝されています。推論タスクでは、この自由がむしろ失敗の軸になることを示します。 GSM8K での LLaDA-8B のデコード中のすべてのコミットメントをログに記録すると、制約のない (純粋な) デコードでは、推論領域の半分がマスクされたままで、軌道の 15 ～ 24% で最終的な答えがコミットされ、キャンバスが大きくなるにつれて問題の最大 90% で答えのみの出力に崩壊することがわかります。原因はモデルの終了信念 (EOS の「圧力」はデコーダー間でほぼ同一) ではなく、到達可能性、つまりサンプラーが離れた位置でそれらの信念に基づいて動作するかどうかです。 2x2 プロンプトデコーダの設計は、順序付けられたコミットメント (インタラクション +34.8 パーセント ポイント、95% CI [26.8, 42.8]、テキストの推論がなければデコーダは区別できない) の下でのみ思考連鎖が役立つことを示しており、このインタラクションを崩壊チャネルと順序チャネルに分解し、Dream-7B と MATH-500 で複製します。単一ノブの介入 (フロンティア ゲート コミットメント) は、最大 4 倍の並列デコードを維持しながら、因果的に完全なギャップ (0.528 から 0.852) を回復します。測定されたフロンティアに沿って、その最適ウィンドウは完全リファインメント時の w=1 から 8 トークン/ステップの制約なしまで反転します。私たちの結果は、以前は効率性を動機としていた既存のウィンドウ スタイル サンプラーを、対処するように設計されていなかった推論的病理に対する最小限の修正として再構成しました。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (保留中)

登録）
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.05687: Answer First, Reason Later: Commitment Order in Diffusion LLMs

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 6 Aug 2026]
Title: Answer First, Reason Later: Commitment Order in Diffusion LLMs
Abstract: Masked diffusion language models (dLLMs) can commit tokens in any order -- a freedom marketed as their core advantage over autoregressive decoding. We show that on reasoning tasks this freedom is instead the axis of failure. Logging every commitment during decoding of LLaDA-8B on GSM8K, we find that unconstrained (pure) decoding commits the final answer at 15-24% of the trajectory while half the reasoning region is still masked, and collapses to answer-only outputs on up to 90% of problems as the canvas grows. The cause is not the model's termination beliefs -- EOS "pressure" is nearly identical across decoders -- but reachability: whether the sampler may act on those beliefs at distant positions. A 2x2 prompt-decoder design shows that chain-of-thought helps only under ordered commitment (interaction +34.8 percentage points, 95% CI [26.8, 42.8]; without reasoning text the decoders are indistinguishable), an interaction we decompose into a collapse channel and an order channel and replicate on Dream-7B and MATH-500. A single-knob intervention -- frontier-gated commitment -- causally recovers the full gap (0.528 to 0.852) while preserving up to 4x parallel decoding, along a measured frontier whose optimal window flips from w=1 at full refinement to unconstrained at 8 tokens/step. Our results reframe existing window-style samplers, previously motivated by efficiency, as the minimal fix for a reasoning pathology they were never designed to address.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
