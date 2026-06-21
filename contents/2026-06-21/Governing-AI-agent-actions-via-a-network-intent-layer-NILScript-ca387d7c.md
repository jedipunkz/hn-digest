---
source: "https://zenodo.org/records/20774491"
hn_url: "https://news.ycombinator.com/item?id=48616674"
title: "Governing AI-agent actions via a network intent layer (NILScript)"
article_title: "Unexpressible, Not Filtered: A Structural Framework for Governing AI-Agent Actions — the Network Intent Layer | Zenodo"
author: "bashierkh"
captured_at: "2026-06-21T10:18:40Z"
capture_tool: "hn-digest"
hn_id: 48616674
score: 1
comments: 0
posted_at: "2026-06-21T07:59:13Z"
tags:
  - hacker-news
  - translated
---

# Governing AI-agent actions via a network intent layer (NILScript)

- HN: [48616674](https://news.ycombinator.com/item?id=48616674)
- Source: [zenodo.org](https://zenodo.org/records/20774491)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T07:59:13Z

## Translation

タイトル: ネットワーク インテント レイヤーを介した AI エージェント アクションの管理 (NILScript)
記事のタイトル: 表現不可能、フィルタリングされていない: AI エージェントのアクションを管理するための構造フレームワーク — ネットワーク インテント レイヤー |ゼノド
説明: 大規模言語モデル (LLM) エージェントは、テキストの生成から、運用システム上でのアクション (払い戻しの発行、レコードの更新、メッセージの送信) の実行に移行しています。独立した企業データは、導入に対する主要な障壁として、モデルの機能ではなく、結果として生じる信頼ギャップを特定しています。スタンフォード大学の
[切り捨てられた]

記事本文:
表現不可能、フィルタリングされていない: AI エージェントのアクションを管理するための構造フレームワーク — ネットワーク インテント層 |ゼノド
メインにスキップ
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
表現不可能、フィルタリングされていない: AI エージェントのアクションを管理するための構造フレームワーク — ネットワーク インテント層
1.
独立研究者 · NILScript
大規模言語モデル (LLM) エージェントは、テキストの生成から、実稼働システム上でのアクション (払い戻しの発行、レコードの更新、メッセージの送信) の実行に移行しています。独立した企業データは、導入に対する主要な障壁として、モデルの機能ではなく、結果として生じる信頼ギャップを特定しています。スタンフォード大学の 2026 年の AI インデックスでは、組織の AI 導入が 88% に達し、実際のエージェント導入が 1 桁にとどまっているにもかかわらず、セキュリティとリスクがエージェント AI の拡張を妨げる最大の要因として 62% であり、次の要因を 24 ポイント上回っていると報告しています。優勢な防御は行動的なものです。エージェントがアクションを作成し、確率的フィルターが事後に安全でないものを捕らえようとします。これは、構築によりゼロ以外の失敗率を認める確率的ポリシーに対する確率的チェックです。構造フレームワークを提案します。ネットワーク インテント レイヤー (NIL) は、エージェントがアクションを発行しないニュートラル ワイヤ コントラクトです。バックエンドが明示的に宣言した操作に対するインテントのみを提案でき、すべての書き込みは決定的な提案、承認、コミット、ロールバックのライフサイクルを通過します。バックエンドが宣言していないアクションは、単にブロックされるだけでなく、表現不可能です。これにより、決定と実行が妨げられます。ポイズニングされた推論ループは依然として書き込みを作成できず、セキュリティ境界は、モデルとは無関係に、すべての推論ステップ (O(n)) から 1 つの意図と効果の境界 (O(1)) まで崩壊します。フレームワークを完全に示します: 4 つの str

構造的保証、静的に検証された複数ステップのプラン言語、監査可能なライフサイクルにわたる人間による承認ゲート、正直な複数ステップの可逆性、およびワイヤレベルの堅牢性 (型付き拒否、決定論的冪等性、サーキットブレーキング)、および InjecAgent 上でインスタンス化された制御された A/B 評価 (4,216 件の間接プロンプトインジェクションケース、2 つのモデル): NIL を介した不正な書き込みが行われました。モデルに依存せず、100% 良性タスク成功の場合は 0.00%。私たちは、メトリクスの定義、反トートロジーの規律、および妥当性に対する脅威を与えます。 NIL は、MCP などのツール統合標準で構成されています。
[切り捨てられた]
統計の収集方法の詳細....
10.5281/ゼノド.20774491
マークダウン
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg)](https://doi.org/10.5281/zenodo.20774491)
再構造化されたテキスト
.. 画像:: https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg
:target: https://doi.org/10.5281/zenodo.20774491
HTML
<a href="https://doi.org/10.5281/zenodo.20774491"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg" alt="DOI"></a>
画像URL
https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg
ターゲット URL
https://doi.org/10.5281/zenodo.20774491
リソースの種類
プレプリント
出版社
ゼノド
言語
英語
権利
提供元
CERN データセンターと InvenioRDM
このサイトでは Cookie を使用しています。 Cookieの使用方法について詳しくはこちらをご覧ください

## Original Extract

Large language model (LLM) agents are moving from generating text to taking actions on production systems: issuing refunds, updating records, sending messages. Independent enterprise data now identifies the resulting trust gap, not model capability, as the dominant barrier to deployment: Stanford's
[truncated]

Unexpressible, Not Filtered: A Structural Framework for Governing AI-Agent Actions — the Network Intent Layer | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Unexpressible, Not Filtered: A Structural Framework for Governing AI-Agent Actions — the Network Intent Layer
1.
Independent Researcher · NILScript
Large language model (LLM) agents are moving from generating text to taking actions on production systems: issuing refunds, updating records, sending messages. Independent enterprise data now identifies the resulting trust gap, not model capability, as the dominant barrier to deployment: Stanford's 2026 AI Index reports security and risk as the top blocker to scaling agentic AI at 62%, a 24-point margin over the next factor, even as organizational AI adoption reaches 88% and actual agent deployment remains in single digits. Prevailing defences are behavioural: the agent authors an action and a probabilistic filter attempts to catch unsafe ones after the fact, a probabilistic check over a probabilistic policy, which admits a nonzero failure rate by construction. We propose a structural framework. The Network Intent Layer (NIL) is a neutral wire contract under which an agent never issues an action; it can only propose intent against operations a backend has explicitly declared, and every write passes a deterministic propose-approve-commit-rollback lifecycle. An action a backend never declared is unexpressible, not merely blocked. This severs deciding from doing: a poisoned reasoning loop still cannot author a write, and the security perimeter collapses from every reasoning step (O(n)) to one intent-to-effect boundary (O(1)), independent of the model. We give the framework in full: four structural guarantees, a statically-validated multi-step plan language, a human-approval gate over an auditable lifecycle, honest multi-step reversibility, and wire-level robustness (typed refusals, deterministic idempotency, circuit-breaking), and a controlled A/B evaluation instantiated on InjecAgent (4,216 indirect prompt-injection cases, two models): unauthorized writes through NIL were 0.00% at 100% benign task-success, model-independently. We give metric definitions, an anti-tautology discipline, and threats to validity. NIL composes with tool-integration standards such as MCP as
[truncated]
More info on how stats are collected....
10.5281/zenodo.20774491
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg)](https://doi.org/10.5281/zenodo.20774491)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg
:target: https://doi.org/10.5281/zenodo.20774491
HTML
<a href="https://doi.org/10.5281/zenodo.20774491"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.20774491.svg
Target URL
https://doi.org/10.5281/zenodo.20774491
Resource type
Preprint
Publisher
Zenodo
Languages
English
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies
