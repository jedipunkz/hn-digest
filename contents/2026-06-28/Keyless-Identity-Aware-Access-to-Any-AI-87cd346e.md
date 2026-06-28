---
source: "https://netbird.ai/"
hn_url: "https://news.ycombinator.com/item?id=48710420"
title: "Keyless, Identity-Aware Access to Any AI"
article_title: "NetBird Agent Network - Keyless, Identity-Aware Access for AI Gateways"
author: "braginini"
captured_at: "2026-06-28T19:29:41Z"
capture_tool: "hn-digest"
hn_id: 48710420
score: 1
comments: 0
posted_at: "2026-06-28T19:06:29Z"
tags:
  - hacker-news
  - translated
---

# Keyless, Identity-Aware Access to Any AI

- HN: [48710420](https://news.ycombinator.com/item?id=48710420)
- Source: [netbird.ai](https://netbird.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T19:06:29Z

## Translation

タイトル: あらゆる AI へのキーレスで ID を認識したアクセス
記事のタイトル: NetBird エージェント ネットワーク - AI ゲートウェイ向けのキーレス、アイデンティティ認識型アクセス
説明: 有効期間の長い AI API キーを ID プロバイダーのグループに置き換えます。検証されたアイデンティティは、監査、コストの帰属、およびポリシーの適用のために、LiteLLM Cloudflare、およびその他のゲートウェイまたはプロバイダーに流れます。

記事本文:
NetBird エージェント ネットワーク - AI ゲートウェイ向けのキーレス、アイデンティティ認識アクセス 価格設定ドキュメントについて オープン ソース エージェント ネットワーク キーレス、
アイデンティティを意識したアクセス
あらゆるAIに。
NetBird は、有効期間の長い AI API キーを、アイデンティティ プロバイダーのグループに関連付けられたネットワーク層のアクセスに置き換えます。検証された ID は、監査、コストの帰属、およびポリシーの適用のために、LiteLLM、Cloudflare、およびその他のゲートウェイに流れます。
ポリシー LiteLLM AI ゲートウェイ スクロール到達可能性 トンネルのみのアクセス。
NetBird は、パブリック イングレスのないプライベート WireGuard ネットワークで AI ゲートウェイをラップします。このネットワークには、OIDC IdP (Okta、Entra、Google) に関連付けられたポリシーゲートの暗号化トンネル経由でのみアクセスできます。ユーザーをグループから削除するか、ユーザーのポリシーを無効にすると、数秒以内にアクセスが切断されます。
NetBird プロキシ https://ai.netbird LiteLLM AI ゲートウェイ ID 共有 API キーはありません。
すべてのリクエストには、実際の発信者の ID (電子メールまたはエージェント名と IdP グループのメンバーシップ) が含まれており、LiteLLM、Cloudflare、または任意のゲートウェイのヘッダーとして NetBird によってスタンプされます。監査ログには実際の人物の名前が記録され、コストは適切なチームに帰属し、グループごとの制限が自動的に適用されます。これらはすべて、静的な API キーではなく IdP によって駆動されます。
設定に API キーがありません。 ID は NetBird プロキシによってスタンプされ、ヘッダーまたはメタデータとしてゲートウェイに転送されます。ガバナンス 支出上限、レート制限、完全な監査。
ゲートウェイがありませんか、または NetBird 自体の内部で支出を制御したいですか?グループまたは個人ごとに、任意のポリシーにトークンとドルの上限を付加します。すべてのリクエストは、ID、モデル、トークン、コスト、レイテンシー、ステータスとともにアクセス ログに記録され、支出を属性化し、暴走エージェントを捕捉し、すべてを SIEM にストリーミングします。
トークン制限グループ: 100k · 個人: 10k · 1 日ごとにリセット 予算制限グループ: $10000 · 個人: $500 · 30 日ごとにリセット アクセス ログ ステータス 時間 ユーザー/エージェント モデル トークン コスト スタ

tus 14:32:08 S sarah.chen@acme.io ユーザー gpt-5.5 1,240 $0.0124 200 S sarah.chen@acme.io ユーザー · 14:32:08 200 gpt-5.5 1,240 · $0.0124 14:32:01 D データ抽出エージェントclaude-opus-4.7 8,512 $0.0851 200 D データ抽出エージェント · 14:32:01 200 claude-opus-4.7 8,512 · $0.0851 14:31:54 M marcus.lee@acme.io ユーザー gpt-5.5 2,104 $0.0421 200 M marcus.lee@acme.io ユーザー · 14:31:54 200 gpt-5.5 2,104 · $0.0421 14:31:47 C crm-sync エージェント gpt-5.5 — — 429 ↳ 予算超過 C crm-sync エージェント · 14:31:47 429 gpt-5.5 — · — ↳ 予算超過 Beyond AI ユニバーサル アクセス飛行機。
AI ゲートウェイの前面にある同じオーバーレイが、データベース、内部サーバー、ステージング、プライベート リソースなど、他のすべての前面にも配置されます。エージェントとユーザーは、暗号化されたピアツーピア WireGuard トンネルを介して直接接続します。WireGuard トンネルは、ポリシーによって管理される、クラウド、オンプレミス、ハイブリッドにわたる 1 つの ID 認識ネットワークです。
データベース Postgres CRM 内部エンジニアリング 384 ユーザー エージェント 192 エージェント マーケティング 66 ユーザー データベース Postgres CRM 内部 エージェント ネットワークを 10 分以内にセットアップします。

## Original Extract

Replace long-lived AI API keys with groups from your identity provider. Verified identity flows into LiteLLM Cloudflare, and other gateways or providers for audit, cost attribution, and policy enforcement.

NetBird Agent Network - Keyless, Identity-Aware Access for AI Gateways About Pricing Docs Open Source Agent Network Keyless,
Identity-Aware Access
to Any AI .
NetBird replaces long-lived AI API keys with network-layer access tied to groups in your identity provider. Verified identity flows into LiteLLM, Cloudflare, and other gateways for audit, cost attribution, and policy enforcement.
Policy LiteLLM AI Gateway Scroll Reachability Tunnel-only access.
NetBird wraps your AI gateway in a private WireGuard network with no public ingress — reachable only through policy-gated encrypted tunnels tied to your OIDC IdP (Okta, Entra, Google). Drop a user from the group or disable their policy, and access drops within seconds.
NetBird Proxy https://ai.netbird LiteLLM AI Gateway Identity No shared API keys.
Every request carries the real caller's identity — email or agent name plus IdP group memberships — stamped by NetBird as headers for LiteLLM, Cloudflare, or any gateway. Audit logs name real people, costs attribute to the right team, and per-group limits enforce themselves, all driven by your IdP instead of a static API key.
No API key in the config. Identity is stamped by the NetBird proxy and forwarded to the gateway as headers or metadata. Governance Spend caps, rate limits, full audit.
No gateway, or want spend controls inside NetBird itself? Attach token and dollar caps to any policy, per group or individual. Every request hits the access log with identity, model, tokens, cost, latency, and status — attribute spend, catch runaway agents, and stream it all to your SIEM.
Token Limit Group: 100k · Individual: 10k · resets every 1d Budget Limit Group: $10000 · Individual: $500 · resets every 30d Access Log Status Time User / Agent Model Tokens Cost Status 14:32:08 S sarah.chen@acme.io User gpt-5.5 1,240 $0.0124 200 S sarah.chen@acme.io User · 14:32:08 200 gpt-5.5 1,240 · $0.0124 14:32:01 D data-extractor Agent claude-opus-4.7 8,512 $0.0851 200 D data-extractor Agent · 14:32:01 200 claude-opus-4.7 8,512 · $0.0851 14:31:54 M marcus.lee@acme.io User gpt-5.5 2,104 $0.0421 200 M marcus.lee@acme.io User · 14:31:54 200 gpt-5.5 2,104 · $0.0421 14:31:47 C crm-sync Agent gpt-5.5 — — 429 ↳ Budget exceeded C crm-sync Agent · 14:31:47 429 gpt-5.5 — · — ↳ Budget exceeded Beyond AI Universal access plane.
The same overlay that fronts your AI gateway fronts everything else too — databases, internal servers, staging, any private resource. Agents and users connect directly over encrypted peer-to-peer WireGuard tunnels: one identity-aware network across cloud, on-prem, and hybrid, governed by your policies.
Database Postgres CRM Internal Engineering 384 Users Agents 192 Agents Marketing 66 Users Database Postgres CRM Internal Set up your Agent Network in under 10 minutes.
