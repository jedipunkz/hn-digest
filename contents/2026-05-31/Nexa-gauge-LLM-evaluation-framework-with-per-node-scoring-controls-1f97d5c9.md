---
source: "https://harnexa.dev/nexa-gauge/docs/introduction"
hn_url: "https://news.ycombinator.com/item?id=48339968"
title: "Nexa-gauge – LLM evaluation framework with per-node scoring controls"
article_title: "Introduction — nexa-gauge Docs — harnexa.dev"
author: "Sardhendu"
captured_at: "2026-05-31T00:26:00Z"
capture_tool: "hn-digest"
hn_id: 48339968
score: 3
comments: 0
posted_at: "2026-05-30T19:50:37Z"
tags:
  - hacker-news
  - translated
---

# Nexa-gauge – LLM evaluation framework with per-node scoring controls

- HN: [48339968](https://news.ycombinator.com/item?id=48339968)
- Source: [harnexa.dev](https://harnexa.dev/nexa-gauge/docs/introduction)
- Score: 3
- Comments: 0
- Posted: 2026-05-30T19:50:37Z

## Translation

タイトル: Nexa-gauge – ノードごとのスコアリング制御を備えた LLM 評価フレームワーク
記事のタイトル: はじめに — nexa-gauge ドキュメント — harnexa.dev
説明: nexa-gauge ドキュメントの概要

記事本文:
はじめに — nexa-gauge ドキュメント — harnexa.dev ドキュメントを読み込み中… ドキュメント / はじめに
nexa-gauge は、LLM および LVLM アプリケーション出力のグラフベースの評価システムです。これにより、アドホックな手動チェックが、ローカル データセットまたはホストされたデータセットで実行できる反復可能なパイプラインに置き換えられます。
生のレコードを型付きの評価状態に正規化します。
選択したターゲットに必要なノードのみを実行します。
確定的キャッシュを通じて以前のノード出力を再利用します。
下流のツール向けに一貫したケースごとのレポートを作成します。
このアーキテクチャは、日々の即時反復、ベンチマーク実行、および測定可能な品質および安全性シグナルによるリリース ゲートをサポートします。
LLM-As-A-Judge が必要な理由
完全一致メトリクスは便利ですが、最新の生成システムでは制限されます。実際のタスクの多くでは、複数の回答が有効になる可能性があり、品質はコンテキストの使用に依存し、障害モードは語彙的ではなく意味論的になります。
LLM-as-a-judge は、明示的な基準に対して出力をスコアリングすることにより、スケーラブルな意味評価を提供します。 nexa-gauge では、この機能がターゲットを絞った指標と組み合わされているため、チームはさまざまな角度から品質を評価できます。
入出力調整との関連性。
提供されたコンテキストにおけるサポートの根拠。
安全性とリスク行動を担当するレッドチーム。
ルーブリックに基づいた判断のためのgeval。
既知の参考回答との重複を考慮して参考にしてください。
nexa-gauge は 2 つの動作モードを提供します。
run は選択したブランチを実行し、最終的な成果物を返します。
推定は、実行前にキャッシュされていない適格コストを計算します。
どちらのモードも同じブランチ計画ロジックに従い、完全な評価を実行する前にコスト見積もりを実行可能にします。
キャッシュはルートを認識し、決定的です。再利用は、入力コンテンツとルーティング セマンティクスが変更されていない場合にのみ発生します。入力、プロンプト、またはモデル ルーティングへの意図的な変更

影響を受けるステップを検証します。
チームは実行前に予算を見積もることができます。
反復実行により、安定したノードの再計算が回避されます。
結果は、固定入力とモデル ルートの下でも再現可能です。
コピー # データセット スライスの完全な評価コストを見積もる
nexagauge 推定値 eval --input sample.json --limit 100
# 完全な評価を実行し、ケースごとのレポート ファイルを作成します
nexagauge run eval --input sample.json --limit 100 --output-dir ./report
データセット フィールド、受け入れられるエイリアス、およびメトリック アクティベーション ルールについては、「データ スキーマ」を参照してください。
反復開発の場合、変更されていない入力とルーティングを繰り返し実行すると、キャッシュの再利用が高く、増分レイテンシが低くなります。

## Original Extract

Overview of nexa-gauge documentation

Introduction — nexa-gauge Docs — harnexa.dev Loading docs… Docs / Introduction Introduction
nexa-gauge is a graph-based evaluation system for LLM and LVLM application outputs. It replaces ad-hoc manual checks with a repeatable pipeline that can be run on local datasets or hosted datasets.
Normalizes raw records into a typed evaluation state.
Executes only the nodes required for the selected target.
Reuses prior node outputs through deterministic caching.
Produces a consistent per-case report for downstream tooling.
This architecture supports day-to-day prompt iteration, benchmark runs, and release gating with measurable quality and safety signals.
Why LLM-As-A-Judge Is Necessary
Exact-match metrics are useful but limited for modern generative systems. In many real tasks, multiple answers can be valid, quality depends on context use, and failure modes are semantic rather than lexical.
LLM-as-a-judge provides scalable semantic evaluation by scoring outputs against explicit criteria. In nexa-gauge, this capability is combined with targeted metrics so teams can evaluate quality from multiple angles:
relevance for input-output alignment.
grounding for support in provided context.
redteam for safety and risk behavior.
geval for rubric-based judgment.
reference for overlap with known reference answers.
nexa-gauge provides two operational modes:
run executes the selected branch and returns final artifacts.
estimate computes uncached eligible cost before execution.
Both modes follow the same branch-planning logic, which makes cost estimates actionable before you run full evaluations.
Caching is route-aware and deterministic. Reuse occurs only when input content and routing semantics are unchanged. Changes to inputs, prompts, or model routing intentionally invalidate affected steps.
Teams can estimate budget before execution.
Iterative runs avoid recomputing stable nodes.
Results remain reproducible under fixed inputs and model routes.
Copy # Estimate full evaluation cost for a dataset slice
nexagauge estimate eval --input sample.json --limit 100
# Run full evaluation and write per-case report files
nexagauge run eval --input sample.json --limit 100 --output-dir ./report
For dataset fields, accepted aliases, and metric activation rules, see the Data Schema .
For iterative development, repeated runs on unchanged inputs and routing should show high cache reuse and lower incremental latency.
