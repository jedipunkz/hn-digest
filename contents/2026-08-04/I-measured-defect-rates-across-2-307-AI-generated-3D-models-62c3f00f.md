---
source: "https://github.com/alza123123/3dqa-benchmark"
hn_url: "https://news.ycombinator.com/item?id=49168919"
title: "I measured defect rates across 2,307 AI-generated 3D models"
article_title: "GitHub - alza123123/3dqa-benchmark: The State of AI-Generated 3D Assets — a public census over 2,307 real generative-AI exports. Reproducible with the published 3dqa package. · GitHub"
author: "soalza116"
captured_at: "2026-08-04T13:50:32Z"
capture_tool: "hn-digest"
hn_id: 49168919
score: 1
comments: 0
posted_at: "2026-08-04T13:45:20Z"
tags:
  - hacker-news
  - translated
---

# I measured defect rates across 2,307 AI-generated 3D models

- HN: [49168919](https://news.ycombinator.com/item?id=49168919)
- Source: [github.com](https://github.com/alza123123/3dqa-benchmark)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T13:45:20Z

## Translation

タイトル: AI で生成された 2,307 個の 3D モデルの欠陥率を測定しました
記事のタイトル: GitHub - alza123123/3dqa-benchmark: AI 生成 3D アセットの現状 — 2,307 件の実際の生成 AI 輸出に関する公開国勢調査。公開されている 3dqa パッケージで再現可能です。 · GitHub
説明: AI 生成 3D アセットの現状 — 2,307 件の実際の生成 AI 輸出に関する公的調査。公開されている 3dqa パッケージで再現可能です。 - alza123123/3dqa-ベンチマーク

記事本文:
GitHub - alza123123/3dqa-benchmark: AI 生成 3D アセットの現状 — 2,307 件の実際の生成 AI 輸出に関する公開国勢調査。公開されている 3dqa パッケージで再現可能です。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アルザ123123
/
3dqa-be

基準点
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット マニフェスト マニフェスト 結果 結果 スクリプト スクリプト .nojekyll .nojekyll ライセンス ライセンス PUBLIC_BENCHMARK_VISION_SAMPLE.md PUBLIC_BENCHMARK_VISION_SAMPLE.md README.md README.md REPRODUCE.md REPRODUCE.md Index.html Index.html すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI によって生成された 3D アセットの現状 — 公開国勢調査
scripts/generate_benchmark_report.py によって results/aggregate.json から生成されます。以下のすべての数値は、ディスク上の実際のファイルに対して、batch_lint.py /batch_lint_Parallel.py によって測定されたものであり、手作業で入力したり外挿したりするものはありません。ジオメトリのみ: 決定的、$0、アセットごとにミリ秒スケール。テクスチャ クラスの QA はありません (この層はテクスチャ マップをサンプリングしません)。
コーパス: 23 のジェネレーター、3D Arena (Hugging Face、MIT) にわたる 2,307 のアセット。
実行日: 2026 年 8 月 1 日。
lint エラー (不正な形式/読み取り不能なアセット、以下の料金から除外): 303
ここで分析されるすべてのアセットは 3D Arena からのものです
Hugging Face 上のデータセット、データセット レベルのライセンスを持つ MIT 。このリポジトリはそうではありません
アセット自体を再配布します (manifests/3d_arena_corpus.jsonl リスト)
リモート パス。 scripts/fetch_real_assets.sh は、
元のソース) — 測定値のみが引き継がれます。全額クレジット
基礎となる生成出力とデータセットのキュレーションが 3D に移行します
Arena プロジェクトと、その出力をホストする個々のジェネレーター チーム。
見出し: クラス別、生成者別の欠陥蔓延率
各セルは、その欠陥を含むジェネレーターのアセットのシェアです (1 つのアセットに複数のアセットが含まれる可能性があるため、行の合計は 100% にはなりません)。
測定できなかったものとその理由
欠陥発見とは区別されます: th

これらのファイルは、このエンジンのメッシュベースのチェックで評価できる範囲からまったく外れていたため (面トポロジのない点群など)、どちらにしても判定が下されず、上記のすべてのレートから除外され、クリーンとしてカウントされません。
顔数中央値: 77,402 · p90: 618,085 · Unity-mobile の公開予算: 30,000 トライアングル。
このコーパス内のすべてのアセットの 79.8% は、この lint run 独自の --max-polys しきい値によってフラグが付けられたかどうかに関係なく、エクスポートされたときにその予算を超えています。
修復速度 (ドライラン修復、ジオメトリのみ)
lint が失敗したアセット (n=1501): 45.0% が完全に修復、49.8% が部分的に修復、5.1% が修復せず、0.0% が修復中にエラーが発生しました。
治癒によってもたらされた回帰: 454 個のアセット。修復によってリグレッションが隠れてはいけないというこのプロジェクト独自のルールに従って、どんなに小さなものであってもここで報告されます。
障害分類 – ドライラン修復で修正できなかったもの
修正不可能な発見
出来事
修正不可能な発見の割合
UNCLOSED_HOLES
577
88.1%
NON_MANIFOLD_EDGES
55
8.4%
FLOATING_GEOMETRY
15
2.3%
NORMALS_INCONSISTENT
7
1.1%
POLY_COUNT_EXCEEDED
1
0.2%
タイミング (アセットごとの実時間、リント + 予行修復)
p50: 15,458 ms · p90: 45,358 ms · 最大: 450,462 ms 。このマシンで、シャードごとに単一プロセスで測定されます。このファイルを生成した並列実行については、scripts/batch_lint_Parallel.py を参照してください。
糸くずのみのタイミング (修復なし) — 厳選されたサンプルの主張に匹敵する数値
p50: 128.6 ミリ秒 · p90: 1294.3 ミリ秒 · 最大: 32061.5 ミリ秒 (n=2004)。個別に測定した 6 資産の厳選されたベースライン (22,000 ～ 65,000 面) も同様に「2.7 ～ 8 ミリ秒」を報告しました (糸くずのみ、ドライラン修復なし)。この行は、フル コーパス スケール (中央値 77,000 顔、最大 6.1M) での同じ測定値であり、中央値では 15 ～ 45 倍になります。一般的な 6 資産の数は引用しないでください。 scripts/measure_lint_timing.py を参照してください。
これは、独自の選択効果を持つ 1 つのデータセットです。 」

「3D Arena (Hugging Face、MIT) のアセット」は、「すべて AI によって生成された 3D」ではありません。
テクスチャクラスの QA クレームはありません。 raster_cpu / この lint 層はテクスチャ マップを決してサンプリングしないため、テクスチャの欠陥 (ストレッチ、継ぎ目、解像度) はここでは対象外であり、クリーンではありません。
このパスにはコントラスト グループがありません。専門家が作成した CC0 ベースライン (Poly Haven) が計画されていますが、このレポートでは実行されていません。上記の数値を AI と人間のギャップとして解釈しないでください。
ビジョンに裏付けられた合格率はこの表には含まれていません。これは小規模な層別サンプル (コスト管理 - 資産ごとに 1 つのモデル呼び出し) に対してのみ実行され、個別にレポートされ、上記の完全なコーパス ジオメトリの数値と混合されることはありません。
ビジョンに裏付けられたサンプル (別のスコープ)
さらに、層化されたシードされたサンプルが完全な閉ループ (修復 → 再レンダリング → ビジョン検証、アセットごとに 1 つのモデル呼び出し) を通じて実行されました。このレポートについては docs/PUBLIC_BENCHMARK_VISION_SAMPLE.md を参照してください。その数値は、上記の完全なコーパスではなく、サンプリングされたサブセットのみを説明しています。
完全なチュートリアルについては、REPRODUCE.md を参照してください。短いバージョン:
python3 -m venv .venv && ソース .venv/bin/activate
pip install "3dqa[repair,fast]" # このランナーが依存する公開パッケージ
./scripts/fetch_real_assets.sh --download # ~22 GB、1 回限り
python scripts/batch_lint_Parallel.py --from-dir サンプル/real_ai_corpus \
--manifest マニフェスト/real_ai_corpus.jsonl --out results/real_ai_corpus.jsonl --workers 5
Python スクリプト/measure_lint_timing.py --manifest マニフェスト/real_ai_corpus_lint.jsonl \
--from-dir サンプル/real_ai_corpus --out results/lint_timing.jsonl
python scripts/aggregate_results.py --results 結果/real_ai_corpus.jsonl \
--out-json 結果/aggregate.json --out-csv 結果/aggregate_by_generator.csv
python scripts/generate_benchmark_report.py --aggregate results/aggregate.json \
--out README.md \
--訪問

オンドキュメント PUBLIC_BENCHMARK_VISION_SAMPLE.md \
--lint-timing 結果/lint_timing.jsonl
このランナーは、geometry_linter /repair_engine / qa_profiles を
3dqa パッケージがインストールされています - このリポジトリにはプライベート ソースを必要としないものはありません。
この方法で公開することの重要な点は次のとおりです。
AI 生成 3D アセットの現状 — 2,307 件の実際の生成 AI 輸出に関する公的調査。公開されている 3dqa パッケージで再現可能です。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The State of AI-Generated 3D Assets — a public census over 2,307 real generative-AI exports. Reproducible with the published 3dqa package. - alza123123/3dqa-benchmark

GitHub - alza123123/3dqa-benchmark: The State of AI-Generated 3D Assets — a public census over 2,307 real generative-AI exports. Reproducible with the published 3dqa package. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
alza123123
/
3dqa-benchmark
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits manifests manifests results results scripts scripts .nojekyll .nojekyll LICENSE LICENSE PUBLIC_BENCHMARK_VISION_SAMPLE.md PUBLIC_BENCHMARK_VISION_SAMPLE.md README.md README.md REPRODUCE.md REPRODUCE.md index.html index.html View all files Repository files navigation
The State of AI-Generated 3D Assets — a public census
Produced by scripts/generate_benchmark_report.py from results/aggregate.json . Every number below was measured by batch_lint.py / batch_lint_parallel.py over real files on disk — none is hand-typed or extrapolated. Geometry-only: deterministic, $0, millisecond-scale per asset. No texture-class QA (this tier never samples a texture map).
Corpus: 2,307 assets across 23 generators, 3D Arena (Hugging Face, MIT).
Run: 2026-08-01.
Lint errors (malformed/unreadable assets, excluded from rates below): 303
All assets analyzed here come from the 3D Arena
dataset on Hugging Face, dataset-level licensed MIT . This repo does not
redistribute the assets themselves ( manifests/3d_arena_corpus.jsonl lists
their remote paths; scripts/fetch_real_assets.sh re-downloads them from the
original source) — only the measurements taken over them. Full credit for
the underlying generative outputs and the dataset's curation goes to the 3D
Arena project and the individual generator teams whose outputs it hosts.
Headline: defect prevalence by class, by generator
Each cell is the share of that generator's assets carrying that defect (a single asset can carry more than one, so rows don't sum to 100%).
What could not be measured, and why
Distinct from a defect finding: these files fell outside what this engine's mesh-based checks can assess at all (e.g. a point cloud with no face topology), so they carry no verdict either way and are excluded from every rate above, not counted as clean.
Median face count: 77,402 · p90: 618,085 · Unity-mobile's published budget: 30,000 triangles.
79.8% of all assets in this corpus exceed that budget as exported, regardless of whether this lint run's own --max-polys threshold flagged them.
Heal rate (dry-run repair, geometry-only)
Of assets that FAILED lint (n=1501): 45.0% heal fully, 49.8% heal partially, 5.1% do not heal, 0.0% error during repair.
Regressions introduced by healing: 454 asset(s). Reported here regardless of how small, per this project's own rule that a repair must never hide a regression.
Failure taxonomy — what dry-run repair could not fix
Unfixable finding
Occurrences
Share of unfixable findings
UNCLOSED_HOLES
577
88.1%
NON_MANIFOLD_EDGES
55
8.4%
FLOATING_GEOMETRY
15
2.3%
NORMALS_INCONSISTENT
7
1.1%
POLY_COUNT_EXCEEDED
1
0.2%
Timing (wall-clock per asset, lint + dry-run repair)
p50: 15,458 ms · p90: 45,358 ms · max: 450,462 ms . Measured on this machine, single-process-per-shard; see scripts/batch_lint_parallel.py for the parallel run that produced this file.
Lint-only timing (no repair) — the number comparable to the curated-sample claim
p50: 128.6 ms · p90: 1294.3 ms · max: 32061.5 ms (n=2004). A separately-measured 6-asset curated baseline (22k-65k faces) reported "2.7-8 ms" the same way (lint only, no dry-run repair); this row is the same measurement at full-corpus scale (median 77k faces, up to 6.1M) and is 15-45x higher at the median — do not quote the 6-asset number as general. See scripts/measure_lint_timing.py .
This is one dataset with its own selection effects. "Assets from 3D Arena (Hugging Face, MIT)" is not "all AI-generated 3D."
No texture-class QA claim. raster_cpu / this lint tier never samples a texture map, so texture defects (stretching, seams, resolution) are out of scope here, not clean.
No contrast group in this pass. A professionally-authored CC0 baseline (Poly Haven) is planned but not run in this report — do not read the numbers above as an AI-vs-human gap.
The vision-backed pass rate is NOT in this table. It only ever runs over a small stratified sample (cost control — one model call per asset) and is reported separately, never blended with the full-corpus geometry numbers above.
Vision-backed sample (separate scope)
A stratified, seeded sample was additionally run through the full closed loop (heal → re-render → vision verification, one model call per asset). See docs/PUBLIC_BENCHMARK_VISION_SAMPLE.md for that report — its numbers describe only the sampled subset, not the full corpus above.
See REPRODUCE.md for the full walkthrough. Short version:
python3 -m venv .venv && source .venv/bin/activate
pip install "3dqa[repair,fast]" # the published package this runner depends on
./scripts/fetch_real_assets.sh --download # ~22 GB, one-time
python scripts/batch_lint_parallel.py --from-dir samples/real_ai_corpus \
--manifest manifests/real_ai_corpus.jsonl --out results/real_ai_corpus.jsonl --workers 5
python scripts/measure_lint_timing.py --manifest manifests/real_ai_corpus_lint.jsonl \
--from-dir samples/real_ai_corpus --out results/lint_timing.jsonl
python scripts/aggregate_results.py --results results/real_ai_corpus.jsonl \
--out-json results/aggregate.json --out-csv results/aggregate_by_generator.csv
python scripts/generate_benchmark_report.py --aggregate results/aggregate.json \
--out README.md \
--vision-doc PUBLIC_BENCHMARK_VISION_SAMPLE.md \
--lint-timing results/lint_timing.jsonl
This runner imports geometry_linter / repair_engine / qa_profiles from the
installed 3dqa package — nothing in this repo needs private source, which
is the whole point of publishing it this way.
The State of AI-Generated 3D Assets — a public census over 2,307 real generative-AI exports. Reproducible with the published 3dqa package.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
