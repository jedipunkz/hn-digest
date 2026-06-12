---
source: "https://github.com/john-rocky/coreai-model-zoo"
hn_url: "https://news.ycombinator.com/item?id=48501723"
title: "Show HN: I ported 11 model families to Apple's new on-device AI framework"
article_title: "GitHub - john-rocky/coreai-model-zoo: Community model zoo + knowledge base for Apple Core AI (iOS/macOS 27): Qwen3.5 & Gemma 4 converted end-to-end, verified on-device (iPhone 17 Pro GPU/ANE), conversion gotchas, custom Metal kernels, Swift runner · GitHub"
author: "mlboy"
captured_at: "2026-06-12T10:35:21Z"
capture_tool: "hn-digest"
hn_id: 48501723
score: 2
comments: 1
posted_at: "2026-06-12T09:13:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I ported 11 model families to Apple's new on-device AI framework

- HN: [48501723](https://news.ycombinator.com/item?id=48501723)
- Source: [github.com](https://github.com/john-rocky/coreai-model-zoo)
- Score: 2
- Comments: 1
- Posted: 2026-06-12T09:13:27Z

## Translation

タイトル: Show HN: 11 のモデル ファミリを Apple の新しいオンデバイス AI フレームワークに移植しました
記事タイトル: GitHub - john-rocky/coreai-model-zoo: Apple Core AI (iOS/macOS 27) のコミュニティ モデル動物園 + ナレッジ ベース: Qwen3.5 と Gemma 4 のエンドツーエンド変換、オンデバイス (iPhone 17 Pro GPU/ANE) の検証、変換の問題、カスタム Metal カーネル、Swift ランナー · GitHub
説明: Apple Core AI (iOS/macOS 27) のコミュニティ モデル ズー + ナレッジ ベース: Qwen3.5 および Gemma 4 のエンドツーエンド変換、オンデバイス (iPhone 17 Pro GPU/ANE) の検証、変換の落とし穴、カスタム Metal カーネル、Swift ランナー - john-rocky/coreai-model-zoo

記事本文:
GitHub - john-rocky/coreai-model-zoo: Apple Core AI のコミュニティ モデル ズー + ナレッジ ベース (iOS/macOS 27): Qwen3.5 および Gemma 4 のエンドツーエンド変換、オンデバイス (iPhone 17 Pro GPU/ANE) の検証、変換の問題点、カスタム Metal カーネル、Swift ランナー · GitHub
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
却下エール

RT
{{ メッセージ }}
ジョン・ロッキー
/
コアイモデル動物園
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
71 コミット 71 コミット アプリ アプリ変換 知識 知識 公式 公式 Swift Swift Zoo Zoo .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Apple Core AI ( .aimodel 、iOS 27 / macOS 27) に変換された LLM — ダウンロード可能、検証済み
デバイス上で、変換コードとナレッジベースを使用して。の後継
CoreML モデル 。
モデル
ダウンロード ( .aimodel )
ライセンス
クウェン3.5-0.8B
🤗 qwen3.5-0.8B-CoreAI
アパッチ-2.0
クウェン3.5-2B
🤗 qwen3.5-2B-CoreAI
アパッチ-2.0
Qwen3.6-35B-A3B (MoE、Mac のみ)
🤗 Qwen3.6-35B-A3B-CoreAI
アパッチ-2.0
Qwen3.6-27B (高密度、Mac のみ)
🤗 Qwen3.6-27B-CoreAI
アパッチ-2.0
GLM-4.7-フラッシュ (MoE + MLA、Mac のみ)
🤗 GLM-4.7-フラッシュ-CoreAI
マサチューセッツ工科大学
Gemma 4 E2B (テキスト、公式 QAT int4 を含む)
🤗 gemma-4-E2B-CoreAI
ジェマ
Gemma 4 E4B (テキスト、公式 QAT int4)
🤗 gemma-4-E4B-CoreAI
ジェマ
LFM2.5-1.2B-命令
🤗 LFM2.5-1.2B-CoreAI
LFM オープン ライセンス v1.0
御影石 4.0-H 1B / 350M
🤗 granite-4.0-h-CoreAI
アパッチ-2.0
Qwen3-VL (ビジョン言語)
🤗 2B・4B・8B
アパッチ-2.0
Gemma 4 E2B ビジョン (VL) (画像 + テキスト)
vl/ で 🤗 gemma-4-E2B-CoreAI
ジェマ
RF-DETR ナノ/小/中/大 (物体検出、NMS なし)
🤗 RF-DETR-CoreAI
アパッチ-2.0
RF-DETR-Seg nano→2xlarge (インスタンス セグメンテーション、6 サイズ)
🤗 RF-DETR-CoreAI
アパッチ-2.0
デコード スループット (tok/s、貪欲; 出力トップ 1 と Hugging Face リファレンスの正確な比較)
iPhone 17 Pro・GPU
iPhone 17 Pro・ANE
M4 Max・GPU
クウェン3.5-0.8B
71.9
14.7
210
クウェン3.5-2B
29
—
161
LFM2.5-1.2B
45.4
—
276.5
御影石 4.0-H 1B
36.3
—
136.5
ジェマ 4 E2B
3

0.3 (QAT 30.7)
6
77.0 (QAT 78.9)
Gemma 4 E4B (公式 QAT)
15.1
—
55.8
Gemma 4 E2B VL (画像+テキスト、公式 QAT)
25.5
—
82.4
Qwen3.6-35B-A3B (MoE、35B/~3B アクティブ、Mac のみ)
—
—
30.9
Qwen3.6-27B (高密度、Mac のみ)
—
—
15.9
GLM-4.7-フラッシュ (MoE + MLA、30B/~3B アクティブ、Mac のみ)
—
—
20.3
iOS 27 / macOS 27 ベータ、Apple の coreai パイプライン GPU エンジン、ゼロカスタムで測定
カーネル (ANE 列を除く)。プレフィル、サイズ、モデルごとの注意事項:zoo/ 。
Qwen3.6-35B-A3B (MoE、35B/~3B アクティブ) — 30.9 tok/s は、
現在のベータ版。動物園/qwen3.6.md
Qwen3.6-27B (密) — 品質の選択: int8 出力 == fp16;密な場合は全体が読み取られます
トークンごとのモデルであるため、~3B-active MoE よりも遅くなります。動物園/qwen3.6-27b.md
GLM-4.7-Flash (MoE + MLA、30B/~3B アクティブ) — 動物園初のマルチヘッド潜在的注意力
モデル; 47 層すべてに完全な MLA を適用します (吸収された MLA は速度のフォローアップです)。
動物園/glm-4.7-flash.md
RF-DETR / RF-DETR-Seg — iPhone 17 Pro で 33 ～ 39 FPS のライブ検出。インスタンス
6 サイズのセグメンテーション、マスクゲート IoU 1.000、M4 Max で 10.7 ～ 59.1 ミリ秒/フレーム。
動物園/rf-detr.md
Gemma 4 E2B VL — 同じテキスト デコーダー + 3 行の画像スプライス。
動物園/gemma4-vl.md
CoreAIChat ( apps/ ) — iPhone 上のデバイス上で実行される動物園のモデル。
監督
何
動物園/
モデルカード - 構成、サイズ、パリティ、測定されたスループット。
知識/
フレームワークに関する確認済みのメモ: 変換、圧縮、ステートフル KV、カスタム Metal カーネル、AOT、コンピューティング ユニット ルール、Swift ランタイム。
変換/
再作成されたモデル + 変換/検証/圧縮スクリプト (PyTorch → .aimodel )。
素早い/
CoreAIRunner — 標準ランタイムを超えるアーキテクチャを含む .aimodel LLM バンドルを駆動する Swift パッケージ。
アプリ/
SwiftUI オンデバイス チャット アプリ (iOS 27): CoreAIChat (Gemma 4 E2B GPU/ANE/⚡ + Qwen3.5 / Qwen3.5)

-2B / LFM2.5 / Granite ⚡パイプライン化、1 ピッカー) + QwenChatFast (Qwen3.5 静的カーネル) (アプリ内モデル ダウンロード付き)。
ここから始めましょう
デバイス上でモデルを実行 →Knowledge/swift-runtime.md + モデルカード
モデルを変換する →Knowledge/conversion-guide.md
圧縮→knowledge/compression.md
高速化 →Knowledge/custom-metal-kernels.md ·Knowledge/performance-ceiling.md
既知のベータ版の問題 (グラフ内 KV 書き込みクラッシュ、回避策 + 入力マスク エスケープ) →Knowledge/coreai-beta-mpsgraph-kvwrite-bug.md — FB23024751 / apple/coreai-models#5
BSD-3 条項 (ライセンス)。再作成されたモデル コードは Apple の BSD-3-Clause から派生しています
coreai_models を管理し、その通知を保持します。モデルの重みは独自のライセンスに従います (それぞれのライセンスを参照)
ハグフェイスレポ）。
Apple Core AI (iOS/macOS 27) のコミュニティ モデル ズー + ナレッジ ベース: Qwen3.5 および Gemma 4 のエンドツーエンド変換、オンデバイス (iPhone 17 Pro GPU/ANE) の検証、変換の問題点、カスタム Metal カーネル、Swift ランナー
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Community model zoo + knowledge base for Apple Core AI (iOS/macOS 27): Qwen3.5 & Gemma 4 converted end-to-end, verified on-device (iPhone 17 Pro GPU/ANE), conversion gotchas, custom Metal kernels, Swift runner - john-rocky/coreai-model-zoo

GitHub - john-rocky/coreai-model-zoo: Community model zoo + knowledge base for Apple Core AI (iOS/macOS 27): Qwen3.5 & Gemma 4 converted end-to-end, verified on-device (iPhone 17 Pro GPU/ANE), conversion gotchas, custom Metal kernels, Swift runner · GitHub
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
john-rocky
/
coreai-model-zoo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
71 Commits 71 Commits apps apps conversion conversion knowledge knowledge official official swift swift zoo zoo .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
LLMs converted to Apple Core AI ( .aimodel , iOS 27 / macOS 27) — downloadable, verified
on-device, with the conversion code and a knowledge base. Successor to
CoreML-Models .
Model
Download ( .aimodel )
License
Qwen3.5-0.8B
🤗 qwen3.5-0.8B-CoreAI
Apache-2.0
Qwen3.5-2B
🤗 qwen3.5-2B-CoreAI
Apache-2.0
Qwen3.6-35B-A3B (MoE, Mac-only)
🤗 Qwen3.6-35B-A3B-CoreAI
Apache-2.0
Qwen3.6-27B (dense, Mac-only)
🤗 Qwen3.6-27B-CoreAI
Apache-2.0
GLM-4.7-Flash (MoE + MLA, Mac-only)
🤗 GLM-4.7-Flash-CoreAI
MIT
Gemma 4 E2B (text, incl. official-QAT int4)
🤗 gemma-4-E2B-CoreAI
Gemma
Gemma 4 E4B (text, official-QAT int4)
🤗 gemma-4-E4B-CoreAI
Gemma
LFM2.5-1.2B-Instruct
🤗 LFM2.5-1.2B-CoreAI
LFM Open License v1.0
Granite 4.0-H 1B / 350M
🤗 granite-4.0-h-CoreAI
Apache-2.0
Qwen3-VL (vision-language)
🤗 2B · 4B · 8B
Apache-2.0
Gemma 4 E2B vision (VL) (image+text)
vl/ in 🤗 gemma-4-E2B-CoreAI
Gemma
RF-DETR nano/small/medium/large (object detection, no NMS)
🤗 RF-DETR-CoreAI
Apache-2.0
RF-DETR-Seg nano→2xlarge (instance segmentation, 6 sizes)
🤗 RF-DETR-CoreAI
Apache-2.0
Decode throughput (tok/s, greedy; output top-1 exact vs the Hugging Face reference)
iPhone 17 Pro · GPU
iPhone 17 Pro · ANE
M4 Max · GPU
Qwen3.5-0.8B
71.9
14.7
210
Qwen3.5-2B
29
—
161
LFM2.5-1.2B
45.4
—
276.5
Granite 4.0-H 1B
36.3
—
136.5
Gemma 4 E2B
30.3 (QAT 30.7)
6
77.0 (QAT 78.9)
Gemma 4 E4B (official QAT)
15.1
—
55.8
Gemma 4 E2B VL (image+text, official QAT)
25.5
—
82.4
Qwen3.6-35B-A3B (MoE, 35B/~3B active, Mac-only)
—
—
30.9
Qwen3.6-27B (dense, Mac-only)
—
—
15.9
GLM-4.7-Flash (MoE + MLA, 30B/~3B active, Mac-only)
—
—
20.3
Measured on the iOS 27 / macOS 27 beta, Apple's coreai-pipelined GPU engine, zero custom
kernels (ANE column excepted). Prefill, sizes, per-model caveats: zoo/ .
Qwen3.6-35B-A3B (MoE, 35B/~3B active) — 30.9 tok/s is expert-gather-bound in the
current beta; zoo/qwen3.6.md
Qwen3.6-27B (dense) — the quality pick: int8 output == fp16; dense reads the whole
model per token, hence slower than the ~3B-active MoE; zoo/qwen3.6-27b.md
GLM-4.7-Flash (MoE + MLA, 30B/~3B active) — the zoo's first Multi-head Latent Attention
model; full-MLA attention on all 47 layers (absorbed-MLA is the speed follow-up);
zoo/glm-4.7-flash.md
RF-DETR / RF-DETR-Seg — detection 33–39 FPS live on iPhone 17 Pro; instance
segmentation in 6 sizes, masks gated IoU 1.000, 10.7–59.1 ms/frame on M4 Max;
zoo/rf-detr.md
Gemma 4 E2B VL — same text decoder + a 3-line image splice;
zoo/gemma4-vl.md
CoreAIChat ( apps/ ) — the zoo's models running on-device on iPhone.
Dir
What
zoo/
Model cards — configurations, sizes, parity, measured throughput.
knowledge/
Verified notes on the framework: conversion, compression, stateful KV, custom Metal kernels, AOT, compute-unit rules, the Swift runtime.
conversion/
Re-authored models + convert / verify / compress scripts (PyTorch → .aimodel ).
swift/
CoreAIRunner — a Swift package that drives .aimodel LLM bundles, including architectures beyond the standard runtime.
apps/
SwiftUI on-device chat apps (iOS 27): CoreAIChat (Gemma 4 E2B GPU/ANE/⚡ + Qwen3.5 / Qwen3.5-2B / LFM2.5 / Granite ⚡pipelined, one picker) + QwenChatFast (Qwen3.5 static kernels) with in-app model download.
Start here
Run a model on device → knowledge/swift-runtime.md + the model card
Convert a model → knowledge/conversion-guide.md
Compress → knowledge/compression.md
Make it fast → knowledge/custom-metal-kernels.md · knowledge/performance-ceiling.md
Known beta issue (in-graph KV-write crash; workarounds + the input-mask escape) → knowledge/coreai-beta-mpsgraph-kvwrite-bug.md — FB23024751 / apple/coreai-models#5
BSD-3-Clause ( LICENSE ). Re-authored model code derives from Apple's BSD-3-Clause
coreai_models and retains its notices. Model weights follow their own licenses (see each
Hugging Face repo).
Community model zoo + knowledge base for Apple Core AI (iOS/macOS 27): Qwen3.5 & Gemma 4 converted end-to-end, verified on-device (iPhone 17 Pro GPU/ANE), conversion gotchas, custom Metal kernels, Swift runner
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
