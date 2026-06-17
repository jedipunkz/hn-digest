---
source: "https://github.com/davedx/cosmicgpt"
hn_url: "https://news.ycombinator.com/item?id=48571585"
title: "Cosmicgpt – A GPT-in-space simulator to research SpaceX AI satellite viability"
article_title: "GitHub - davedx/cosmicgpt: A GPT-in-space simulator to research SpaceX AI satellite viability · GitHub"
author: "davedx"
captured_at: "2026-06-17T16:22:19Z"
capture_tool: "hn-digest"
hn_id: 48571585
score: 2
comments: 1
posted_at: "2026-06-17T15:09:47Z"
tags:
  - hacker-news
  - translated
---

# Cosmicgpt – A GPT-in-space simulator to research SpaceX AI satellite viability

- HN: [48571585](https://news.ycombinator.com/item?id=48571585)
- Source: [github.com](https://github.com/davedx/cosmicgpt)
- Score: 2
- Comments: 1
- Posted: 2026-06-17T15:09:47Z

## Translation

タイトル: Cosmicgpt – SpaceX AI 衛星の実行可能性を研究するための GPT-in-space シミュレーター
記事のタイトル: GitHub - davedx/cosmicgpt: SpaceX AI 衛星の実行可能性を研究するための GPT-in-space シミュレーター · GitHub
説明: SpaceX AI 衛星の実行可能性を研究するための GPT-in-space シミュレーター - davedx/cosmicgpt

記事本文:
GitHub - davedx/cosmicgpt: SpaceX AI 衛星の実行可能性を研究するための GPT-in-space シミュレーター · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ダベックス
/
コズミックプト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット docs docs scenarios scenarios scripts scripts src/ cosmicgpt src/ cosmicgpt testing testing .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md DESIGN.md DESIGN.md README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
宇宙条件下で GPT 推論に何が起こるかをシミュレートする — cosmic-ray bit
モデルの重み、アクティベーション、
KV キャッシュと出力。
📊 ライブレポートを見る → davedx.github.io/cosmicgpt
放射線が AI モデルの出力に与える影響を確認する: シングル実行レポート
そして環境比較。
モデル化する目標と条件については、DESIGN.md を参照してください。
技術設計用の ARCHITECTURE.md。
ステータス: ビジュアライゼーション + HTML レポート (ステップ 5)
エンドツーエンドのループは、3 つの単一イベント効果分類法を完全にカバーします。
破損しやすい領域。手動で指定された障害または物理的な障害に起因する障害が発生します。
放射線環境: シードされた nanoGPT を構築する
(実際の KV キャッシュを使用して)、クリーンなベースラインを生成し、障害を取得します (手動または
フラックス スケジューラ )、それらを注入します (重みの突然変異、
アクティベーションフォワードフック、KVキャッシュミューテーション）、同じサンプリングシードで再生成、
そして差分。
障害の種類 ( --kind ): SEU (シングル ビット フリップ)、MBU (マルチ ビット アセット)、
STUCK_AT (セルは 0/1 に固定)、SEL (ラッチアップ - テンソル全体がゼロ化)、
SET (一時的なアクティベーショングリッチ)、SEFI (NaN/ガベージカスケード)。
リージョン ( --region ): 重み、アクティベーション (lm_head → logits を含む)、kv_cache 。
環境 ( --orbit ): LEO、SAA、POLAR、GEO、INTERPLANETARY、SOLAR_STORM、
推論の途中で λ(t) を引き上げるオプションの太陽フレア バースト ウィンドウ。
実行ごとに、失敗モード (silent_correct / 微妙な間違い / 繰り返し /
ゴミ/な

n_garbage / crash)、故障までの時間、および平均 KL 発散
出力分布を提供し、ステップごとの RunTrace を出力できます。
JSON (今後のビジュアライゼーションが消費するデータ)。
# 軌道から物理的に派生した断層 (短時間の実行で影響がわかるように磁束をスケール化)
cosmicgpt run --orbit SAA --flux-mult 1e4 --tokens 120
# 推論中期の太陽フレアを伴うミッション
cosmicgpt 実行シナリオ/mission_solar_storm.yaml
# 自己完結型の HTML レポートを作成します (トークンの差分 + 劣化タイムライン + ラスター)
cosmicgpt run --orbit SOLAR_STORM --flux-mult 1e4 --report report.html
# 保存されたトレースからレポートを再生成します - 再推論は行いません
cosmicgpt レポート run/storm/trace.json -o report.html
# 条件を並べて比較します (ビュー C)
cosmicgpt 比較 --orbits LEO、SAA、SOLAR_STORM -o 比較.html
レポートは完全に自己完結型です (インライン CSS + インライン SVG、外部アセットなし、
matplotlib) なので、電子メールで送信したり、アーカイブしたりできます。
python -m venv .venv && ソース .venv/bin/activate
pip install -e " .[dev] "
# 最小のシナリオ (SEU) を実行します
cosmicgpt 実行シナリオ/walking_skeleton.yaml
# 分類法を直接操作する
cosmicgpt run --kind SEFI --n-flips 1 --tokens 120 --fault-seed 3
cosmicgpt run --kind SEL --n-flips 8 --tokens 100
# ビットフリップ基盤 + インジェクションメカニズムを検証する
pytest
初期の発見
影響の少ないサイトでの単一障害 (バイアス、仮数部のビット数が少ない) が日常的に発生します。
マスクされた — 現実的: 宇宙線の衝突のほとんどは、目に見えるものは何もありません。
指数/符号の反転と SEL は、仮数の反転よりもはるかに破壊的です。
SET (一時的なアクティブ化グリッチ) は穏やかです。持続性がなければ、1 つの問題に影響を与えます。
ステップ、および放出された位置に着地した場合のみ。
モデルには実際の KV キャッシュ ( --region kv_cache ) が追加されました。そこにあるストライクが変更されました。
以降のすべてのトークンが破損したエントリを再読み取りするため、一度だけ実行されますが持続します。

を通して
注意。リージョンは障害の種類には依存しません — --regionweight|activation|kv_cache 。
LEO の単一の短い推論は、現実的なアップセット率では基本的に誤りがありません。
重大な破損には SAA、太陽嵐、または長時間露光が必要です。フレア付き
バースト ウィンドウでは、磁束が急増するとすぐに発散が始まります。
このモデルは、シードされた小さな、ランダムに初期化された文字レベル GPT であるため、ベースラインは
テキストは意味不明ですが、骨格としては問題ありません。重要なのは、
フォールト挿入ループとそれが (特に float exponer で) 目に見えて反転します
出力が破損します。後で scripts/train_tiny.py (ロードマップ) を介して一貫したモデルをトレーニングします。
src/cosmicgpt/
model/ nanogpt.py (+KV キャッシュ)、adapter.py、sites.py # モデル + 障害レジストリ
aults/ bitops.py、types.py、injector.py # 分類法 + インジェクション
環境/ flux.py、presets.py、scheduler.py # スケーリングされた物理フラックス
eval/ランナー、メトリクス、分類、トレース # ループ + メトリクス + RunTrace
viz/ svg、diffview、タイムライン、レポート # インライン SVG/HTML レポート
config.py、cli.py
シナリオ/ウォーキング_スケルトン.yaml、sefi_cascade.yaml、ミッション_ソーラー_ストーム.yaml
テスト/test_bitops、test_injection、test_kvcache、test_scheduler、test_eval、test_viz
ロードマップ
ARCHITECTURE.md §11 を参照してください。次 (ステップ 6): 緩和ラッパー
(ECC / TMR 投票 / スクラビング / NaN ガード) と費用対効果の実験を行った後、
プラグイン可能なより大きな GPT バックエンドを使用して、調査結果が一般化するかどうかをテストします。
SpaceX AI 衛星の実行可能性を研究するための GPT-in-space シミュレーター
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A GPT-in-space simulator to research SpaceX AI satellite viability - davedx/cosmicgpt

GitHub - davedx/cosmicgpt: A GPT-in-space simulator to research SpaceX AI satellite viability · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
davedx
/
cosmicgpt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits docs docs scenarios scenarios scripts scripts src/ cosmicgpt src/ cosmicgpt tests tests .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md DESIGN.md DESIGN.md README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Simulate what happens to GPT inference under space conditions — cosmic-ray bit
flips and other radiation-induced faults corrupting a model's weights, activations,
KV cache, and output.
📊 View the live reports → davedx.github.io/cosmicgpt
See what radiation does to an AI model's output: a single-run report
and an environment comparison .
See DESIGN.md for goals and the conditions we model, and
ARCHITECTURE.md for the technical design.
Status: visualizations + HTML reports (step 5)
The end-to-end loop covers the full Single-Event-Effect taxonomy across three
corruptible regions , with faults either hand-specified or derived from a physical
radiation environment : build a seeded nanoGPT
(with a real KV cache), generate a clean baseline, get faults (manual or from the
flux scheduler ), inject them (weight mutations,
activation forward-hooks, KV-cache mutations), regenerate with the same sampling seed,
and diff.
Fault kinds ( --kind ): SEU (single bit flip), MBU (multi-bit upset),
STUCK_AT (cell pinned 0/1), SEL (latch-up — a whole tensor zeroed),
SET (transient activation glitch), SEFI (NaN/garbage cascade).
Regions ( --region ): weight , activation (incl. lm_head → logits), kv_cache .
Environments ( --orbit ): LEO, SAA, POLAR, GEO, INTERPLANETARY, SOLAR_STORM , with an
optional solar-flare burst window raising λ(t) mid-inference.
Every run also reports a failure mode (silent_correct / subtle_wrong / repetition /
garbage / nan_garbage / crash), time-to-failure , and mean KL divergence of the
output distribution, and can emit a per-step RunTrace
JSON (the data the upcoming visualizations consume).
# physically-derived faults from an orbit (flux scaled so a short run shows effects)
cosmicgpt run --orbit SAA --flux-mult 1e4 --tokens 120
# a mission with a mid-inference solar flare
cosmicgpt run scenarios/mission_solar_storm.yaml
# write a self-contained HTML report (token diff + degradation timeline + raster)
cosmicgpt run --orbit SOLAR_STORM --flux-mult 1e4 --report report.html
# regenerate a report from a saved trace — no re-inference
cosmicgpt report runs/storm/trace.json -o report.html
# compare conditions side by side (View C)
cosmicgpt compare --orbits LEO,SAA,SOLAR_STORM -o comparison.html
Reports are fully self-contained (inline CSS + inline SVG, no external assets, no
matplotlib) so they're emailable and archivable.
python -m venv .venv && source .venv/bin/activate
pip install -e " .[dev] "
# run the smallest scenario (SEU)
cosmicgpt run scenarios/walking_skeleton.yaml
# drive the taxonomy directly
cosmicgpt run --kind SEFI --n-flips 1 --tokens 120 --fault-seed 3
cosmicgpt run --kind SEL --n-flips 8 --tokens 100
# verify the bit-flip foundation + injection mechanisms
pytest
Early findings
Single faults on low-impact sites (biases, low mantissa bits) are routinely
masked — realistic: most cosmic-ray hits do nothing visible.
Exponent/sign flips and SEL are far more destructive than mantissa flips.
SET (transient activation glitch) is gentle: without persistence it affects one
step, and only if it lands on the emitted position.
The model now has a real KV cache ( --region kv_cache ): a strike there is mutated
once but persists , because every later token re-reads the corrupted entry through
attention. Region is independent of fault kind — --region weight|activation|kv_cache .
A single short inference in LEO is essentially fault-free at realistic upset rates;
meaningful corruption needs the SAA, a solar storm, or long exposure. With a flare
burst window , divergence visibly begins right when the flux spikes.
The model is a small, seeded, randomly-initialized char-level GPT, so the baseline
text is gibberish — but that's fine for the skeleton: the point is to demonstrate the
fault-injection loop and that flips (especially in the float exponent ) measurably
corrupt the output. Train a coherent model later via scripts/train_tiny.py (roadmap).
src/cosmicgpt/
model/ nanogpt.py (+KV cache), adapter.py, sites.py # model + fault registry
faults/ bitops.py, types.py, injector.py # taxonomy + injection
environment/ flux.py, presets.py, scheduler.py # scaled-physical flux
eval/ runner, metrics, classify, trace # loop + metrics + RunTrace
viz/ svg, diffview, timeline, report # inline-SVG/HTML reports
config.py, cli.py
scenarios/ walking_skeleton.yaml, sefi_cascade.yaml, mission_solar_storm.yaml
tests/ test_bitops, test_injection, test_kvcache, test_scheduler, test_eval, test_viz
Roadmap
See ARCHITECTURE.md §11 . Next (step 6): mitigation wrappers
(ECC / TMR voting / scrubbing / NaN guards) with cost-benefit experiments, then a
pluggable larger-GPT backend to test whether findings generalize.
A GPT-in-space simulator to research SpaceX AI satellite viability
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
