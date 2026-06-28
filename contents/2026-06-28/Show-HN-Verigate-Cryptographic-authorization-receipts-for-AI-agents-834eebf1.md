---
source: "https://verigate.cloud"
hn_url: "https://news.ycombinator.com/item?id=48709538"
title: "Show HN: Verigate – Cryptographic authorization receipts for AI agents"
article_title: "Verigate — Cryptographic Trust Infrastructure for AI Agents"
author: "heartlinmachado"
captured_at: "2026-06-28T18:23:08Z"
capture_tool: "hn-digest"
hn_id: 48709538
score: 1
comments: 0
posted_at: "2026-06-28T17:33:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Verigate – Cryptographic authorization receipts for AI agents

- HN: [48709538](https://news.ycombinator.com/item?id=48709538)
- Source: [verigate.cloud](https://verigate.cloud)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T17:33:49Z

## Translation

タイトル: HN を表示: Verigate – AI エージェントの暗号化認証レシート
記事のタイトル: Verigate — AI エージェントのための暗号化信頼インフラストラクチャ
説明: 認証受領書、コンプライアンス レポート、および検証可能なエージェント ID。トラストパス内の LLM がゼロ。

記事本文:
検証する
特長
価格設定
コンプライアンス
ドキュメント
無料で始める
EU AI 法の施行は -- 日以内
エージェントのあらゆるアクション。
署名入り。連鎖した。検証済み。
暗号化された認証レシート、自動化されたコンプライアンス レポート、および検証可能なエージェント ID。トラストパス内の LLM がゼロ。
無料で始める
受信チェーン.ログ
$ 検証検証 --chain
--- 領収書 #1047 ---
シーケンス: 1047
アクション: 読み取り:顧客:プロファイル
エージェント: エージェント:crm アシスタント
署名: ed25519:9f3a...c8e1
前: sha256:7d2b...4a90
--- 領収書 #1048 ---
シーケンス: 1048
アクション: 書き込み:注文:キャンセル
エージェント: エージェント:ops-bot
署名: ed25519:1b7e...f2d3
チェーン: 有効 (受信 2 回、休憩 0 回)
_
Ed25519 署名入り ·
ミリ秒未満のレイテンシ ·
6 フレームワーク ·
ゼロLLM ·
ハッシュチェーン監査 ·
マークルの証明 ·
Ed25519 署名入り ·
ミリ秒未満のレイテンシ ·
6 フレームワーク ·
ゼロLLM ·
ハッシュチェーン監査 ·
マークルの証明 ·
Ed25519 署名入り ·
ミリ秒未満のレイテンシ ·
6 フレームワーク ·
ゼロLLM ·
ハッシュチェーン監査 ·
マークルの証明 ·
Ed25519 署名入り ·
ミリ秒未満のレイテンシ ·
6 フレームワーク ·
ゼロLLM ·
ハッシュチェーン監査 ·
マークルの証明 ·
インフラストラクチャー
ゼロトラスト エージェント環境向けに構築
すべてのエージェントのアクションは、Ed25519 署名付きの承認受領書を発行する決定論的なゲートウェイを通過します。トラストパスに LLM がありません。各レシートはその前のレシートとハッシュチェーンされ、改ざんが明らかな監査証跡が作成されます。
EU AI 法、DORA、HIPAA、SEC 17a-4、NIST、OWASP にマッピングされた監査対応レポートを生成します。監査エージェントは領収書チェーンをスキャンし、証拠の引用を含むコンプライアンス判定を生成します。 LLM でサニタイズされた出力により、監査アーティファクトへの即時挿入が防止されます。
Ed25519 キー ペアを使用した DPoP バインド エージェント登録。エージェントは、生存チャレンジを通じて自分の身元を証明します。エージェント カードは、検出可能なメタデータを提供します。 SPIFFE ワークロード ID 標準と互換性があります。
領収書はマークル ツリーと

ベース L2 に固定されています。ハッシュを再計算し、Ed25519 署名を検証し、チェーン上のルートに対してマークル証明をチェックするなど、誰でも独立してチェーンを検証できます。完全なオフライン検証がサポートされています。
エージェントのアクティビティを 6 つの規制およびセキュリティのフレームワークに自動的にマッピングします。
透明性と予測可能な価格設定
席ごとの料金はかかりません。エージェントが使用した分を支払います。
コンプライアンスレポート（6つのフレームワークすべて）
第 14 条では、リスクの高い AI システムに対して人間による監視と監査証跡を義務付けています。 Verigate は、すべてのエージェントのアクションが承認され、連鎖され、検証可能であることを暗号化して証明します。

## Original Extract

Authorization receipts, compliance reports, and verifiable agent identity. Zero LLM in the trust path.

Verigate
Features
Pricing
Compliance
Docs
Start Free
EU AI Act enforcement in -- days
Every agent action.
Signed. Chained. Verified.
Cryptographic authorization receipts, automated compliance reports, and verifiable agent identity. Zero LLM in the trust path.
Start Free
receipt-chain.log
$ verigate verify --chain
--- receipt #1047 ---
seq: 1047
action: read:customer:profile
agent: agent:crm-assistant
sig: ed25519:9f3a...c8e1
prev: sha256:7d2b...4a90
--- receipt #1048 ---
seq: 1048
action: write:order:cancel
agent: agent:ops-bot
sig: ed25519:1b7e...f2d3
chain: VALID (2 receipts, 0 breaks)
_
Ed25519 Signed ·
Sub-ms Latency ·
6 Frameworks ·
Zero LLM ·
Hash Chain Audit ·
Merkle Proofs ·
Ed25519 Signed ·
Sub-ms Latency ·
6 Frameworks ·
Zero LLM ·
Hash Chain Audit ·
Merkle Proofs ·
Ed25519 Signed ·
Sub-ms Latency ·
6 Frameworks ·
Zero LLM ·
Hash Chain Audit ·
Merkle Proofs ·
Ed25519 Signed ·
Sub-ms Latency ·
6 Frameworks ·
Zero LLM ·
Hash Chain Audit ·
Merkle Proofs ·
Infrastructure
Built for zero-trust agent environments
Every agent action passes through a deterministic gateway that issues Ed25519-signed authorization receipts. No LLM in the trust path. Each receipt is hash-chained to its predecessor, creating a tamper-evident audit trail.
Generate audit-ready reports mapped to EU AI Act, DORA, HIPAA, SEC 17a-4, NIST, and OWASP. The auditor agent scans your receipt chain and produces compliance verdicts with evidence citations. LLM-sanitized output prevents prompt injection in audit artifacts.
DPoP-bound agent registration with Ed25519 key pairs. Agents prove their identity via liveness challenges. Agent cards provide discoverable metadata. Compatible with SPIFFE workload identity standards.
Receipts are organized into Merkle trees and anchored to Base L2. Anyone can independently verify the chain: recompute hashes, validate Ed25519 signatures, check Merkle proofs against on-chain roots. Full offline verification supported.
Automated mapping of your agent activity to six regulatory and security frameworks.
Transparent, predictable pricing
No per-seat fees. Pay for what your agents use.
Compliance reports (all 6 frameworks)
Article 14 requires human oversight and audit trails for high-risk AI systems. Verigate gives you cryptographic proof that every agent action was authorized, chained, and verifiable.
