---
source: "https://ctxgov.github.io/ctxgov/try-in-5-minutes.html"
hn_url: "https://news.ycombinator.com/item?id=48482262"
title: "Show HN: CtxGov – a local claim firewall for AI memory claims"
article_title: "Try CtxGov in 5 minutes"
author: "LuxBennu"
captured_at: "2026-06-10T21:44:47Z"
capture_tool: "hn-digest"
hn_id: 48482262
score: 1
comments: 0
posted_at: "2026-06-10T20:33:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CtxGov – a local claim firewall for AI memory claims

- HN: [48482262](https://news.ycombinator.com/item?id=48482262)
- Source: [ctxgov.github.io](https://ctxgov.github.io/ctxgov/try-in-5-minutes.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T20:33:14Z

## Translation

タイトル: 表示 HN: CtxGov – AI メモリ クレーム用のローカル クレーム ファイアウォール
記事のタイトル: 5 分で CtxGov を試す
説明: パブリックセーフな 1 コマンドの Memory X-Ray プレビューを実行し、短いフィードバックを残します。

記事本文:
公共の安全な地元製品パス
CtxGov のクローンを作成し、コマンド 1 つで Memory X-Ray プレビューを実行し、レポートを読みます。
オプションでコンパニオン評価レシートを検査します。このパスは、
ベンチマークの解釈ではなく、製品の形状と再現性です。
git クローン https://github.com/ctxgov/ctxgov.git
cd ctxgov
02
走る
python3 スクリプト/run_memory_xray_demo.py
03
レポートを読む
メモリ-xray-demo-report.htmlを開く
または docs/memory-xray-demo-report.md を調べてください。
同行者を調べる
v0.11.0 フレッシュクローンの受領書
ローカル証拠コンテキスト用。
採用に関する主張や下流の本番使用に関する主張はありません。
任意のターゲット リポジトリ スキャンはありません。
Memory X-Ray へのファイル入力を表示する修正されたパブリック セーフ フィクスチャを参照してください。
レポートの形状。これは任意のリポジトリ スキャナーではありません。
python3 スクリプト/run_tiny_fixture_memory_xray_demo.py
このコマンドは、固定された公衆安全フィクスチャのみを読み取り、書き込みます。
docs/tiny-fixture-memory-xray-demo.html プラス
docs/tiny-fixture-memory-xray-demo.json 。
これを貼り付けます
GitHub の問題 #22 。
短いコメントが 1 つあれば十分です。
走行経路がクリアになっていますか?はい/いいえ:
レポートは役に立ちますか?はい/いいえ:
欠落しているフィールド:
統合形状:
紛らわしい表現:
境界: 公的ベンチマークの主張はありません。養子縁組の主張はありません。セキュリティなし
主張する。プロバイダー/モデルの互換性については主張しません。

## Original Extract

Run the public-safe one-command Memory X-Ray preview and leave short feedback.

Public-safe local product path
Clone CtxGov, run the one-command Memory X-Ray preview, read the report,
and optionally inspect the companion eval receipt. This path is for
product shape and reproducibility, not benchmark interpretation.
git clone https://github.com/ctxgov/ctxgov.git
cd ctxgov
02
Run
python3 scripts/run_memory_xray_demo.py
03
Read report
Open memory-xray-demo-report.html
or inspect docs/memory-xray-demo-report.md .
Inspect the companion
v0.11.0 fresh-clone receipt
for local evidence context.
No adoption claim or downstream production-use claim.
No arbitrary target repository scan.
See a fixed public-safe fixture that shows file input to Memory X-Ray
report shape. It is not an arbitrary repo scanner.
python3 scripts/run_tiny_fixture_memory_xray_demo.py
The command reads the fixed public-safe fixture only and writes
docs/tiny-fixture-memory-xray-demo.html plus
docs/tiny-fixture-memory-xray-demo.json .
Paste this into
GitHub issue #22 .
One short comment is enough.
Run path clear? yes/no:
Report useful? yes/no:
Missing field:
Integration shape:
Confusing wording:
Boundary: No public benchmark claim. No adoption claim. No security
claim. No provider/model compatibility claim.
