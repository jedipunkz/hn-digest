---
source: "https://realvuln.com"
hn_url: "https://news.ycombinator.com/item?id=48503425"
title: "Show HN: We're inviting Anthropic to put the real Mythos 5 on our open benchmark"
article_title: "Dashboard — RealVuln"
author: "jfaganel99"
captured_at: "2026-06-12T13:16:45Z"
capture_tool: "hn-digest"
hn_id: 48503425
score: 3
comments: 1
posted_at: "2026-06-12T12:44:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: We're inviting Anthropic to put the real Mythos 5 on our open benchmark

- HN: [48503425](https://news.ycombinator.com/item?id=48503425)
- Source: [realvuln.com](https://realvuln.com)
- Score: 3
- Comments: 1
- Posted: 2026-06-12T12:44:50Z

## Translation

タイトル: Show HN: 本物の Mythos 5 をオープン ベンチマークに入れるよう Anthropic を招待します
記事のタイトル: ダッシュボード — RealVuln
説明: インタラクティブな RealVuln v1.0 ダッシュボード: リーダーボード、精度と再現率、コスト効率、再現率と精度のランキング、カテゴリ分析、および 15 台のスキャナーにわたるデータセット構成。

記事本文:
▚ 本物の脆弱性 v1.0
ダッシュボード
方法論
データセット
調査結果
ロードマップ
GitHub
≡
ダッシュボード
方法論
データセット
調査結果
ロードマップ
GitHub ↗
ベンチマークダッシュボード
24 のスキャナー、26 のリポジトリ、F3 (厳密) によるランク付け
メトリック
F2
F3
24 スキャナー 3 カテゴリ
26 リポジトリ Python · タイプ 1
92.4 ベスト F3 (厳密) コレガ エンタープライズ
95.3 最高再現率 % Kolega Enterprise
93.2 最高精度 % Grok 4.20
リーダーボード
脆弱性クラスによる検出
▸ LLM ベースのスキャナーは、SQL インジェクション、コマンド インジェクション、安全でない逆シリアル化など、セマンティックなデータ フローの理解を必要とするクラスを支配します。 ▸ ルールベースのツールは構文パターンでのみ競争力を維持しており、それでも全体的な再現率は低いままです。
すべての数値は、24 のスキャナーと 26 のリポジトリにわたる RealVuln のライブ結果です。 F3 重みは 9 倍の精度を再現します。厳密モードでは、未完了のリポジトリはミスとしてカウントされます。コストは、スコアリングされた実行の合計支出です (ルールベースのツールは無料または変動価格です)。メトリクスの定義 →

## Original Extract

Interactive RealVuln v1.0 dashboard: leaderboard, precision–recall, cost-efficiency, recall and precision rankings, category analysis, and dataset composition across 15 scanners.

▚ real vuln v1.0
Dashboard
Methodology
Dataset
Findings
Roadmap
GitHub
≡
Dashboard
Methodology
Dataset
Findings
Roadmap
GitHub ↗
Benchmark dashboard
24 scanners · 26 repositories · ranked by F3 (strict)
Metric
F2
F3
24 Scanners 3 categories
26 Repositories Python · Type 1
92.4 Best F3 (strict) Kolega Enterprise
95.3 Highest recall % Kolega Enterprise
93.2 Highest precision % Grok 4.20
Leaderboard
Detection by vulnerability class
▸ LLM-based scanners dominate classes that need semantic data-flow understanding — SQL injection, command injection, insecure deserialization. ▸ Rule-based tools stay competitive only on syntactic patterns, and even there overall recall remains low.
All figures are live RealVuln results across 24 scanners and 26 repositories. F3 weights recall nine times over precision; strict mode counts unfinished repositories as misses. Cost is the total spend for the scored run (rule-based tools are free or variably priced). Metric definitions →
