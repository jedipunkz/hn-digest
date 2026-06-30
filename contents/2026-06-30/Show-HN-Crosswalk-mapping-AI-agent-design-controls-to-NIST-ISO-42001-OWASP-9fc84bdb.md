---
source: "https://www.agent-kits.com/agentaz-crosswalk"
hn_url: "https://news.ycombinator.com/item?id=48730140"
title: "Show HN: Crosswalk mapping AI-agent design controls to NIST, ISO 42001, OWASP"
article_title: "AgentAz™ Regulatory Crosswalk — NIST AI RMF, ISO 42001, OWASP Agentic — AgentKits"
author: "stoicstoic"
captured_at: "2026-06-30T09:03:30Z"
capture_tool: "hn-digest"
hn_id: 48730140
score: 1
comments: 0
posted_at: "2026-06-30T08:57:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Crosswalk mapping AI-agent design controls to NIST, ISO 42001, OWASP

- HN: [48730140](https://news.ycombinator.com/item?id=48730140)
- Source: [www.agent-kits.com](https://www.agent-kits.com/agentaz-crosswalk)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T08:57:38Z

## Translation

タイトル: Show HN: AI エージェント設計制御を NIST、ISO 42001、OWASP にマッピングする Crosswalk
記事のタイトル: AgentAz™ 規制横断歩道 — NIST AI RMF、ISO 42001、OWASP Agentic — AgentKits
Description: How AgentAz™ design-time governance dimensions map to NIST AI RMF 1.0, ISO/IEC 42001:2023, and the OWASP Top 10 for Agentic Applications (ASI01–ASI10).コンプライアンスに関するアンケートの無料リファレンス。正直な範囲のギャップがマークされています。

記事本文:
AgentAz™ Regulatory Crosswalk — NIST AI RMF、ISO 42001、OWASP Agentic — AgentKits コンテンツにスキップ AgentKits 参照ツール ブログ 保存された検索について ⌘ K AgentAz™ 仕様 / Regulatory Crosswalk AgentAz™ Regulatory Crosswalk
AgentAz は、設計時のガバナンス ボキャブラリです。このクロスウォークは、AgentAz の各ディメンションを、企業がすでに監査されている可能性が高い 3 つのフレームワークを満たすのに役立つコントロールにマッピングします。そのため、機械可読な Agentaz.json がセキュリティ アンケートのガバナンス セクションのショートカットになります。
スペックの価値は、それがどのようにマッピングされるかです。各行を「この AgentAz ディメンションを宣言するエージェントは、これらのコントロールに対する設計時の証拠を生成しています。」と読み替えます。これは認証ではなく、 に対する証拠です。マッピングは、監査人が実行中のシステムを検証したということではなく、意図が文書化されていることを示しています。
この横断歩道が主張しないもの
正直な範囲が重要です。ガバナンス マッピングは、そのギャップが明示されている場合にのみ役立ちます。 AgentAz は、設計時、機械可読、ブループリント レベルの 1 つのレーンにとどまります。実行時の証明、エージェントの ID、または認証は対象外です。具体的には範囲外です:
OWASP ASI04 (サプライ チェーン)、ASI05 (コード実行のサンドボックス化)、ASI06 (メモリ/RAG ポイズニング)、および ASI07 (エージェント間通信) は、ランタイムおよびインフラストラクチャの防御です。 AgentAz は設計時の意図を文書化します。これらは実装されていません。これらはランタイム層とセキュリティ層に属します。
NIST MEASURE バイアス/公平性の深さと完全な TEVV 方法論。 AgentAz は境界に焦点を当てており、公平性テストの方法論ではありません。これらはせいぜい部分的なものとして扱ってください。
ISO/IEC 42001 A.7 (データ ガバナンス) および A.10 (サードパーティとの関係)。単一のブループリントの設計時の仕様を大幅に逸脱しています。
マッピングはアンケートの開始点であり、アンケートの開始点ではありません。

コンプライアンス判定。監査人は、環境内でコントロールが満たされているかどうかを判断します。
NIST AI RMF 1.0 (2023)、ISO/IEC 42001:2023、および OWASP Top 10 for Agentic Applications (ASI01–ASI10、2025 年 12 月) の公開された構造に対してマッピングされています。
Crosswalk バージョン 1.0 · 最終レビュー日 2026 年 6 月 30 日。これらのフレームワークは時間の経過とともに改訂されます (特に OWASP エージェント リストは新しく、進化しています)。監査で信頼する前に、現在公開されている制御テキストに対して行を検証します。
AgentAz™ 仕様を読む それに対してエージェントをスキャンする
本番環境に対応した AI エージェント ブループリント — 実際のエージェントを構築するためのアーキテクチャ、プロンプト、ツール、ワークフロー。ログインがありません。

## Original Extract

How AgentAz™ design-time governance dimensions map to NIST AI RMF 1.0, ISO/IEC 42001:2023, and the OWASP Top 10 for Agentic Applications (ASI01–ASI10). A free reference for compliance questionnaires — with honest scope gaps marked.

AgentAz™ Regulatory Crosswalk — NIST AI RMF, ISO 42001, OWASP Agentic — AgentKits Skip to content AgentKits Browse Tools Blog About Saved Search ⌘ K AgentAz™ Specifications / Regulatory Crosswalk AgentAz™ Regulatory Crosswalk
AgentAz is a design-time governance vocabulary. This crosswalk maps each AgentAz dimension to the controls it helps satisfy in three frameworks an enterprise is likely already audited against — so a machine-readable agentaz.json becomes a shortcut through the governance section of a security questionnaire.
A spec’s worth is what it maps up to. Read each row as: “an agent that declares this AgentAz dimension is producing design-time evidence toward these controls.” It is evidence toward , not a certification — the mapping shows intent is documented, not that an auditor has verified the running system.
What this crosswalk does not claim
Honest scope is the point — a governance mapping is only useful if its gaps are stated. AgentAz stays in one lane: design-time, machine-readable, blueprint-level. It does not cover runtime proof, agent identity, or certification. Specifically out of scope:
OWASP ASI04 (Supply Chain), ASI05 (sandboxing of code execution), ASI06 (memory/RAG poisoning), and ASI07 (inter-agent communication) are runtime and infrastructure defenses. AgentAz documents design-time intent; it does not implement these — they belong to your runtime and security layers.
NIST MEASURE bias/fairness depth and full TEVV methodology. AgentAz is boundary-focused, not a fairness-testing methodology; treat these as partial at best.
ISO/IEC 42001 A.7 (data governance) and A.10 (third-party relationships). Largely outside a single blueprint's design-time spec.
A mapping is a starting point for a questionnaire, not a compliance verdict. Your auditor still determines whether a control is satisfied in your environment.
Mapped against the published structure of NIST AI RMF 1.0 (2023), ISO/IEC 42001:2023 , and the OWASP Top 10 for Agentic Applications (ASI01–ASI10, December 2025).
Crosswalk version 1.0 · Last reviewed 2026-06-30. These frameworks are revised over time (the OWASP agentic list especially is new and evolving) — verify any row against the current published control text before relying on it in an audit.
Read the AgentAz™ spec Scan an agent against it
Production-ready AI Agent Blueprints — architecture, prompts, tools, and workflows to build real agents. No login.
