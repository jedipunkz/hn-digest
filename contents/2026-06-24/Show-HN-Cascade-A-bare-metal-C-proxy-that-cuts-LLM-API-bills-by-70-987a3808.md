---
source: "https://cascade-router.github.io/cascade-router/"
hn_url: "https://news.ycombinator.com/item?id=48662106"
title: "Show HN: Cascade – A bare-metal C++ proxy that cuts LLM API bills by 70%"
article_title: "Cascade Router | Intelligent AI Proxy"
author: "AmixxM"
captured_at: "2026-06-24T16:42:29Z"
capture_tool: "hn-digest"
hn_id: 48662106
score: 1
comments: 0
posted_at: "2026-06-24T16:17:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cascade – A bare-metal C++ proxy that cuts LLM API bills by 70%

- HN: [48662106](https://news.ycombinator.com/item?id=48662106)
- Source: [cascade-router.github.io](https://cascade-router.github.io/cascade-router/)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T16:17:33Z

## Translation

タイトル: Show HN: Cascade – LLM API の料金を 70% 削減するベアメタル C++ プロキシ
記事のタイトル: カスケードルーター |インテリジェント AI プロキシ
HNテキスト: https://github.com/Cascade-Router/cascade-router

記事本文:
">
カスケードルーター
ホワイトペーパー
GitHub
エンタープライズに連絡する
過剰な支払いをやめる
予測可能な推論。
ベアメタル C++ AI プロキシは、プロンプトの複雑さを 4.59 ミリ秒で予測し、最もコスト効率の高い LLM にトラフィックを動的にルーティングします。
トークン化、ONNX 埋め込み、ML 予測をカバーするエンドツーエンドの実行。
基本的な抽出および分類タスクをフロンティア層ではなく高速モデルに自動的にルーティングすることで、待ち時間を追加することなく OpenAI と Anthropic のコストを削減します。
小規模モデルが検証に失敗した場合、リクエストはシームレスにフロンティア モデルにエスカレーションされます。
遅延を追加することなく、予測可能なプロンプトをより安価なモデルにルーティングすることで、Cascade がどれだけ節約できるかを見積もります。 Cascade は、単純なワークロードの 75% をフロンティア モデル (gpt-4o や Claude 3.5 Sonnet など) ではなく高速モデル (gpt-4o-mini や Claude 3 Haiku など) に自動的にルーティングします。
ルーティング可能なトラフィックでは、高速モデルとフロンティア モデルのコストが最大 60% 低いと想定されます。実際の節約額は、ワークロードの組み合わせやプロバイダー (OpenAI、Anthropic など) によって異なります。
今すぐオープンソース プロキシをセルフホストしましょう。エンタープライズ ライセンスにアップグレードすると、独自のデータ、SOC2 監査ログ、および SSO 統合に基づいてトレーニングされたカスタム ルーティング ウェイトのロックが解除されます。

## Original Extract

https://github.com/Cascade-Router/cascade-router

">
Cascade Router
Whitepaper
GitHub
Contact Enterprise
Stop overpaying for
predictable inference.
A bare-metal C++ AI proxy that predicts prompt complexity in 4.59 milliseconds and dynamically routes traffic to the most cost-effective LLM.
End-to-end execution covering tokenization, ONNX embedding, and ML prediction.
Automatically routes basic extraction and classification tasks to fast models instead of frontier tiers—cutting OpenAI and Anthropic bills without added latency.
If a small model fails validation, requests seamlessly escalate to frontier models.
Estimate how much Cascade can save by routing predictable prompts to cheaper models without adding latency. Cascade automatically routes 75% of your simple workloads to fast models (like gpt-4o-mini or Claude 3 Haiku) instead of frontier models (like gpt-4o or Claude 3.5 Sonnet).
Assumes ~60% lower cost for fast models vs frontier models on routable traffic. Actual savings vary by workload mix and provider (OpenAI, Anthropic, etc.).
Self-host the open-source proxy today. Upgrade to an Enterprise License to unlock custom routing weights trained on your proprietary data, SOC2 audit logging, and SSO integration.
