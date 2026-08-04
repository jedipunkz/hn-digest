---
source: "https://mantis-llm-gateway.github.io/"
hn_url: "https://news.ycombinator.com/item?id=49173664"
title: "Mantis – an open-source, self-hosted LLM gateway"
article_title: "Mantis | Mantis"
author: "rizsyed1"
captured_at: "2026-08-04T20:20:15Z"
capture_tool: "hn-digest"
hn_id: 49173664
score: 1
comments: 0
posted_at: "2026-08-04T19:25:28Z"
tags:
  - hacker-news
  - translated
---

# Mantis – an open-source, self-hosted LLM gateway

- HN: [49173664](https://news.ycombinator.com/item?id=49173664)
- Source: [mantis-llm-gateway.github.io](https://mantis-llm-gateway.github.io/)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T19:25:28Z

## Translation

タイトル: Mantis – オープンソースの自己ホスト型 LLM ゲートウェイ
記事タイトル: カマキリ |カマキリ
説明: オープンソースのセルフホスト型 LLM ゲートウェイのドキュメント。

記事本文:
コンテンツへスキップ Mantis 検索 Ctrl K キャンセル チーム / Github Dark SELF_HOSTED: TRUE CLOUD_PROVIDER: AWS API_STYLE: OPENAI Mantis
ルーティング、キャッシュ、ガードレール、モデル プロバイダー間の可観測性を実現するセルフホスト型 LLM ゲートウェイ。
複数のモデル ターゲットの前に配置される安定したチャット完了エンドポイント。クライアント コードを 1 行も変更せずにプロバイダーを切り替えます。
AWS エコシステム向けに構築されています。高速キャッシュには ElastiCache を、基礎モデルには Bedrock を使用して、Terraform 経由で ECS にデプロイします。
ルーティング、再試行、およびフォールバック ロジック。タイムアウト、クールダウン期間、キャッシュ動作を構成内で直接定義します。
メタデータ、モデルのエイリアス、重み付けされたターゲット、およびフォールバック チェーンによってリクエストをルーティングします。
検証、キャッシュ チェック、クールダウン、プロバイダー呼び出し、再試行、端末応答を調整します。
オプションのルーティング メタデータを使用して、単一のゲートウェイ エンドポイントを介してチャット完了リクエストを送信します。
各 HTTP リクエストを手動で構築せずに、アプリケーション コードから Mantis を呼び出します。
レイ・ヴァン・デン・バーグ アイルランド、ダブリン リズ・サイード、イギリス ロンドン サミュエル・メカ 台北、台湾
// 次のステップ

## Original Extract

Open-source, self-hosted LLM gateway documentation.

Skip to content Mantis Search Ctrl K Cancel Team / Github Dark SELF_HOSTED: TRUE CLOUD_PROVIDER: AWS API_STYLE: OPENAI Man tis
A self-hosted LLM gateway for routing, caching, guardrails, and observability across model providers.
A stable chat completions endpoint that sits in front of multiple model targets. Switch providers without changing a single line of client code.
Built for the AWS ecosystem. Deploys via Terraform to ECS, using ElastiCache for fast caching and Bedrock for foundational models.
Routing, retry, and fallback logic. Define timeouts, cooldown periods, and cache behavior directly in your config.
Route requests by metadata, model aliases, weighted targets, and fallback chains.
Coordinate validation, cache checks, cooldowns, provider calls, retries, and terminal responses.
Send chat completion requests through a single gateway endpoint with optional routing metadata.
Call Mantis from application code without manually constructing each HTTP request.
Rey van den Berg Dublin, Ireland Riz Syed London, UK Samuel Meka Taipei, Taiwan
// Next Steps
