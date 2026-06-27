---
source: "https://github.com/Arkalogy/best-photo-picker"
hn_url: "https://news.ycombinator.com/item?id=48693638"
title: "AI Powered Photo Gallery Without the Cloud"
article_title: "GitHub - Arkalogy/best-photo-picker: Pick the best photos from large libraries. Quality scoring, perceptual dedup, and per-person boost — runs entirely on your machine, no cloud. · GitHub"
author: "anonu"
captured_at: "2026-06-27T00:30:38Z"
capture_tool: "hn-digest"
hn_id: 48693638
score: 1
comments: 0
posted_at: "2026-06-27T00:07:04Z"
tags:
  - hacker-news
  - translated
---

# AI Powered Photo Gallery Without the Cloud

- HN: [48693638](https://news.ycombinator.com/item?id=48693638)
- Source: [github.com](https://github.com/Arkalogy/best-photo-picker)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T00:07:04Z

## Translation

タイトル: クラウドを使用しない AI を活用したフォト ギャラリー
記事のタイトル: GitHub - Arkalogy/best-photo-picker: 大規模なライブラリから最高の写真を選択します。品質スコアリング、知覚的重複排除、および個人ごとのブーストは、クラウドを使用せず、完全にマシン上で実行されます。 · GitHub
説明: 大規模なライブラリから最高の写真を選択します。品質スコアリング、知覚的重複排除、および個人ごとのブーストは、クラウドを使用せず、完全にマシン上で実行されます。 - Arkalogy/ベストフォトピッカー

記事本文:
GitHub - Arkalogy/best-photo-picker: 大規模なライブラリから最高の写真を選択します。品質スコアリング、知覚的重複排除、および個人ごとのブーストは、クラウドを使用せず、完全にマシン上で実行されます。 · GitHub
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
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。もう一度お願いします

このページをロードします。
アーカロ学
/
ベストフォトピッカー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .githooks .githooks .github .github bpp bpp デスクトップ デスクトップ ドキュメント ドキュメント 例/ plugin_example 例/ plugin_example スクリプト スクリプト テスト-e2e テスト-e2e テスト-js テスト-js テスト テスト .dockerignore .dockerignore .eslint-globals.json .eslint-globals.json .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE MANIFEST.in MANIFEST.in MODEL_POLICY.md MODEL_POLICY.md NOTICE.txt NOTICE.txt README.md README.md SECURITY.md SECURITY.md bpp-server.spec bpp-server.spec eslint.config.js eslint.config.js launch.command.template launch.command.template package-lock.json package-lock.json package.json package.json playwright.config.mjs playwright.config.mjs pyproject.toml pyproject.toml sample_config.yaml sample_config.yaml tsconfig.json tsconfig.json vitest.config.mjs vitest.config.mjs すべてのファイルを表示 リポジトリ ファイルのナビゲーション
15,000 枚の家族写真に溺れていませんか? Apple フォトと Google フォトは、ベストショットをフィードに埋め込みます。 Best Photo Picker はマシン上で完全に実行され、すべての写真の鮮明さ、照明、顔、構図を採点し、数分で完璧な 50 枚を厳選できます。クラウドもサブスクリプションも妥協もありません。
これは、何もアップロードせずにキュレーションのためにライブラリ全体を自動スコアリングし、ほぼ重複したものをクラスタリングして、特定の人をブーストできる唯一の写真ツールです。

したがって、最終的な選択は、重要な顔、つまりあなたの子供、祖父母を優先します。
30 秒でお試しいただけます。写真は必要ありません。
pip install " bppicker[web] " && bpp デモ
サンプル ライブラリを生成し、完全な UI をローカルで実行し、タブを閉じると正常に終了します。マシンからは何も残りません。
Arkalogy による — PM 主導の AI 構築製品。アーキテクチャ上の決定は docs/adr/ に取り込まれます。なぜ、どのようにして建てられたのか→
完全なワークフロー (インポート → 調整 → 選択 → エクスポート) と側面 (面、カレンダー、マップ、複製) の注釈付きウォークスルーについては、 docs/quickstart-gallery.md を参照してください。
📣 バグを見つけましたか? それとも機能のアイデアがありますか?問題を開く · ディスカッションを開始する · ガイドに貢献する。セキュリティの問題: 公開の問題ではなく、非公開の脆弱性レポートを使用してください。
ベストフォトピッカー
アップルの写真
Googleフォト
ライトルーム
デジカム
ローカルファースト / オフラインで実行
✅
⚠️ iCloud連携
❌ 雲
⚠️クラウド同期
✅
写真を選択するための自動スコアリング
✅
❌
⚠️「思い出」
⚠️旗のみ
❌
指名された人物によるブーストピック
✅
⚠️
⚠️
❌
❌
ほぼ重複した重複排除
✅
⚠️
⚠️
❌
✅
無料 + オープンソース
✅ マサチューセッツ工科大学
❌
❌
❌ $120/年
✅ GPL
アカウントなし/サブスクリプションなし
✅
❌
❌
❌
✅
オプションのデバイスペアリングによる LAN 共有
✅
❌
❌
❌
❌
あなたが Apple/Google エコシステムに住んでいて、自分のコレクションがそのフィードに適合している場合、これらのツールはうまく機能します。 bpp は、彼らが苦手なケースに対応しています。子供、愛犬、昨年の夏の旅行など、特定のテーマの写真が何千枚もあり、どこにもアップロードせずに実際のベスト 50 枚を見つけたいと考えています。
品質スコアリング - シャープネス、露出、顔検出 (YuNet + SCRFD + BlazeFace + MediaPipe + dlib)、合成
スマートな重複排除 — 知覚的ハッシュ (dHash + aHash) および CL

IP セマンティックの類似性
時間的多様性 — バランスのとれた選択のための 1 日あたりの上限と月次の補償範囲
顔認識 — 自動クラスタリング、個人ごとのスマート アルバム、結合、削除、再割り当て、誤検出された bbox のドラッグによる修正
ライブラリ管理 - SHA-256 重複除去を使用して管理ライブラリにコピーしてインポート
アルバム システム — マニュアル + 19 種類のスマート アルバム (人物、時間、スコア、重複、ペットなど)
インタラクティブな Web UI — リアルタイムのスライダー調整、写真グリッド、ライトボックス、比較ビュー、バッチ操作
デスクトップ アプリ — Tauri v2 経由のネイティブ macOS アプリ (Web UI をネイティブ ウィンドウにラップ)
論理的な削除 - 30 日間の回復、最近削除したアルバム
デモ モード — 生成されたサンプル写真を使ってすぐに試してみましょう
プライバシー — ローカルファースト、テレメトリーや分析は不要です。 bpp が行ういくつかのネットワーク呼び出し (最初の分析時のモデルのダウンロード、マップ ビューを開いたときの OpenStreetMap タイル、GitHub リリースに対する更新チェック、オプションの LAN 共有、オプションの追加機能の pip インストール) はすべて以下に列挙されており、個別に開示可能です。ライブラリに関するものは何も送信されません。
Python 3.11 が必要です (3.12 以降はまだサポートされていません - 一致するように固定されています)
デスクトップサイドカー）。 macOS では、Homebrew にはデフォルトで新しい Python が同梱されています。
brew install python@3.11 で 3.11 をインストールするか、
pyenv 、または非 Python を使用します
以下のデスクトップアプリ。
推奨 — pipx は bpp を独自に保持します
環境を変更し、1 つのコマンドで更新を行います。
pipx インストール " bppicker[web] "
# 後で最新リリースに更新します:
pipx アップグレード bppicker
プレーンな pip も機能します。
pip インストール " bppicker[web] "
# オプション: 顔認識
pip install " bppicker[web,faces] "
# オプション: HEIC サポート
pip インストール " bppicker[heic] "
ML を利用したほとんどの機能 (顔認識、NudeNet、RAW インポート、
HEIC、AI 修復) 設定 → 詳細からオンデマンドでインストール
→ ML モデル

実行中のアプリ内で、またはプレインストールすることもできます
pip install "bppicker[heic,faces,raw,nudity,inpaint]" ですべてを一度に実行します。
デスクトップ アプリ (macOS、Apple Silicon)
ターミナルよりもドックアイコンの方が好みですか?各リリースにはスタンドアロンが同梱されます
マックのアプリ。 Python もターミナルもありません:
BestPhotoPicker-macOS-arm64.dmg をダウンロードします (常に最新リリース)。
.dmg をダブルクリックし、Best Photo Picker を [アプリケーション] にドラッグします。
アプリケーションまたはスポットライトから開きます。
アプリは Apple Developer ID で署名され、Apple によって公証されています。
ダブルクリックで開きます。セキュリティ警告はありません。更新するには、ダウンロードしてください
新しい .dmg も同様です。写真ライブラリと設定が公開されます
アプリの外ではそのまま引き継がれます。
Apple Silicon (M1/M2/M3 以降) のみ。 Intel Mac または
Windows/Linux の場合は、上記の pipx install "bppicker[web]" パスを使用します。
bpp デモ
サンプル写真を生成し、Web UI を起動します。構成は必要ありません。
# 写真管理サーバーを起動します (デフォルトのライブラリ: ~/Pictures/BestPhotoPicker)
BPPサーブ
# またはカスタムライブラリのパスを指定します
bpp サーブ --library ~ /写真/2024
ブラウザで http://127.0.0.1:5001 を開きます。フォルダーをインポートし、スライダーでスコアの重みを調整し、厳選した選択をエクスポートします。
# ワンショット: 分析 + ベスト 50 枚の写真を選択
bpp run --input ~ /Photos/MyKid --k 50 --out ~ /curated --gallery
# または段階的に:
bpp 分析 --input ~ /Photos/MyKid --out ~ /workdir
bpp select --workdir ~ /workdir --k 50 --out ~ /curated --gallery
仕組み
インポート → 分析 → スコア → 重複排除 → 選択 → エクスポート
インポート : SHA-256 重複除去を使用して管理ライブラリにコピーされた写真
分析: SQLite にキャッシュされた並列特徴抽出 (ブラー、露出、顔、合成)
スコア: サブスコアの重み付けされた組み合わせ、スライダーを使用してリアルタイムで調整可能
重複排除 : 知覚的ハッシュ クラスタリングの削除

エスバーストの重複。オプションの CLIP セマンティック重複排除
選択 : 時間的多様性を考慮した 1 日あたりの上限と月次の範囲を備えた貪欲な選択
エクスポート: 選択した写真をオプションの HTML ギャラリーとともにコピー/ハードリンク/シンボリックリンクします。
特徴
説明
フォトグリッド
スコアバッジ付きのサムネイルグリッド、ズームコントロール、並べ替えとフィルター
ライトボックス
キーボードナビゲーションを備えたフルサイズのビューア
スライダー
リアルタイムのウェイト調整（ブラー、露出、顔、構図）
BPP のおすすめ
サイドバーのサブ項目は、ライブラリ全体から最高の K 写真を選択します。ツールバーチップフィルターは現在のアルバム/ビューで選択します
アルバム
マニュアルアルバムとスマートアルバム（時間別、スコア別、人物別）
顔
結合、消去、人物ごとのアルバムを使用した自動クラスタ化された顔検出
バッチ操作
Cmd/Ctrl キーを押しながらクリックして複数選択、一括で含める/除外する/お気に入り
インポート
フォルダーまたはアーカイブをライブラリにドラッグします
エクスポート
選択した写真を出力ディレクトリにコピーまたはリンクします
チューニングノブ
パラメータ
デフォルト
説明
ぼかしの重み
0.30
集計スコアの鮮明さの重み付け
露出の重み
0.20
露出品質の重み付け
顔の重み
0.35
顔検出とフレーミングのためのウェイト
組成_重量
0.15
構図の重み付け（三分割法）
最大長辺
1024
分析前に画像をこれまでダウンスケールする
ハッシュ距離閾値
10
「同じ」画像の dHash ハミング距離
time_window_秒
15
このウィンドウ内でバースト写真をクラスター化します
1 日あたりの最大値
3
暦日ごとに選択できる写真の最大数
min_per_month
1
少なくとも月に 1 件は含めるようにしてください
ハードウェア
bpp は、デフォルトですべての ML 推論を CPU 上で実行します。 ONNX ベースのモデル (SCRFD)
顔検出、CLIP セマンティック検索、YOLOv11n ペット検出) を選択可能
BPP_ONNX_PROVIDERS 環境変数を介してハードウェア アクセラレーションに移行します。
# Apple Silicon (M1/M2/M3) — CoreML 経由で Apple Neural Engine を使用する
BPP_ONNX_PROVIDERS= " CoreMLExecutionProvider,CPUExecutionProvider " \
bpp サーブ --librar

y ~ /写真/BestPhotoPicker
# NVIDIA GPU (Linux) — 最初に onnxruntime-gpu をインストールします
pip uninstall -y onnxruntime && pip install onnxruntime-gpu
BPP_ONNX_PROVIDERS= " CUDAExecutionProvider,CPUExecutionProvider " \
bppserve --library ~ /Pictures/BestPhotoPicker
# Windows DirectML
BPP_ONNX_PROVIDERS= " DmlExecutionProvider,CPUExecutionProvider " \
bppserve --library ~ /Pictures/BestPhotoPicker
CPU は常に最終フォールバックとして追加されるため、プロバイダーが欠落しています
あなたのホイールでは、警告ではなくきれいに通過します。
クラッシュする。順序は重要です。プロバイダーはリストされている順序で試行されます。
表面
CIでテスト済み
サポートされています
注意事項
macOS arm64 (M1/M2/M3) — CPU
はい (開発ボックス)
はい
プライマリ開発ターゲット。 Tauri デスクトップ アプリはこれのみを同梱しています。
Linux x86_64 — CPU
はい (ubuntu-最新)
はい
PyPI インストールのデフォルト。
Linux arm64 — CPU
いいえ
はい
Docker イメージのビルド。ランタイムは CI では実行されません。
Windows — CPU
いいえ
はい
コードパスが存在します。決してエンドツーエンドで実行しないでください。 GitHub の問題経由でのレポートを歓迎します。
macOS arm64 — CoreMLExecutionProvider (ANE)
いいえ
はい
コードパスが存在します。顔推論の速度が 5 ～ 10 倍向上すると予想されます。正確さは、特定の ONNX グラフと ONNX ランタイムのバージョンによって異なります。
Linux x86_64 — CUDAExecutionProvider
いいえ
はい
pip install onnxruntime-gpu が必要です (ベース onnxruntime を置き換えます)。無料層 CI には GPU ランナーがありません。
Windows — DmlExecutionProvider
いいえ
はい
Windows CPUと同じ注意点 — タラ

[切り捨てられた]

## Original Extract

Pick the best photos from large libraries. Quality scoring, perceptual dedup, and per-person boost — runs entirely on your machine, no cloud. - Arkalogy/best-photo-picker

GitHub - Arkalogy/best-photo-picker: Pick the best photos from large libraries. Quality scoring, perceptual dedup, and per-person boost — runs entirely on your machine, no cloud. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Arkalogy
/
best-photo-picker
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .githooks .githooks .github .github bpp bpp desktop desktop docs docs examples/ plugin_example examples/ plugin_example scripts scripts tests-e2e tests-e2e tests-js tests-js tests tests .dockerignore .dockerignore .eslint-globals.json .eslint-globals.json .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE MANIFEST.in MANIFEST.in MODEL_POLICY.md MODEL_POLICY.md NOTICE.txt NOTICE.txt README.md README.md SECURITY.md SECURITY.md bpp-server.spec bpp-server.spec eslint.config.js eslint.config.js launch.command.template launch.command.template package-lock.json package-lock.json package.json package.json playwright.config.mjs playwright.config.mjs pyproject.toml pyproject.toml sample_config.yaml sample_config.yaml tsconfig.json tsconfig.json vitest.config.mjs vitest.config.mjs View all files Repository files navigation
Drowning in 15,000 family photos? Apple Photos and Google Photos bury your best shots in their feed. Best Photo Picker runs entirely on your machine, scores every photo on sharpness, lighting, faces, and composition, and lets you curate the perfect 50 in minutes — no cloud, no subscription, no compromise.
It's the only photo tool that auto-scores your whole library for curation without uploading anything , then clusters near-duplicates and lets you boost specific people so the final pick favors the faces that matter — your kid, for the grandparents.
Try it in 30 seconds — no photos needed:
pip install " bppicker[web] " && bpp demo
Generates a sample library, runs the full UI locally, and quits cleanly when you close the tab. Nothing leaves your machine.
By Arkalogy — a PM-directed, AI-built product; the architectural decisions are captured in docs/adr/ . Why & how it was built →
For an annotated walkthrough of the full workflow — import → adjust → pick → export — plus the side surfaces (faces, calendar, map, duplicates), see docs/quickstart-gallery.md .
📣 Found a bug or have a feature idea? Open an issue · Start a discussion · Contributing guide . Security issues: please use private vulnerability reporting , not a public issue.
Best Photo Picker
Apple Photos
Google Photos
Lightroom
digiKam
Local-first / runs offline
✅
⚠️ iCloud-coupled
❌ cloud
⚠️ cloud sync
✅
Auto-scores photos for picking
✅
❌
⚠️ "Memories"
⚠️ flags only
❌
Boost picks by named person
✅
⚠️
⚠️
❌
❌
Near-duplicate deduplication
✅
⚠️
⚠️
❌
✅
Free + open source
✅ MIT
❌
❌
❌ $120/yr
✅ GPL
No account / no subscription
✅
❌
❌
❌
✅
Optional LAN sharing w/ device pairing
✅
❌
❌
❌
❌
If you live in the Apple/Google ecosystem and your collection fits their feeds, those tools work great. bpp exists for the case they're bad at: you have thousands of photos of a specific subject — your kid, your dog, last summer's trip — and you want to find the actual best fifty without uploading anything anywhere.
Quality scoring — sharpness, exposure, face detection (YuNet + SCRFD + BlazeFace + MediaPipe + dlib), composition
Smart deduplication — perceptual hashing (dHash + aHash) and CLIP semantic similarity
Temporal diversity — per-day caps and monthly coverage for balanced selection
Face recognition — automatic clustering, per-person smart albums, merge, dismiss, reassign, drag-to-fix mis-detected bboxes
Library management — import by copying to managed library with SHA-256 dedup
Album system — manual + 19 smart album types (person, time, score, duplicates, pets, etc.)
Interactive web UI — real-time slider tuning, photo grid, lightbox, compare view, batch operations
Desktop app — native macOS app via Tauri v2 (wraps the web UI in a native window)
Soft delete — 30-day recovery, Recently Deleted album
Demo mode — try instantly with generated sample photos
Privacy — local-first, no telemetry, no analytics. The few network calls bpp makes (model downloads on first analyze, OpenStreetMap tiles when you open the Map view, update checks against GitHub Releases, optional LAN sharing, optional pip install of extra features) are all enumerated below and individually disclosable; nothing about your library is ever sent.
Requires Python 3.11 (3.12+ is not yet supported — it's pinned to match
the desktop sidecar). On macOS, Homebrew ships a newer Python by default;
install 3.11 with brew install python@3.11 or
pyenv , or use the no-Python
desktop app below.
Recommended — pipx keeps bpp in its own
environment and makes updating one command:
pipx install " bppicker[web] "
# Update to the latest release later:
pipx upgrade bppicker
Plain pip works too:
pip install " bppicker[web] "
# Optional: face recognition
pip install " bppicker[web,faces] "
# Optional: HEIC support
pip install " bppicker[heic] "
Most ML-powered features (face recognition, NudeNet, RAW import,
HEIC, AI inpainting) install on demand from Settings → Advanced
→ ML Models in the running app, or you can pre-install
everything at once with pip install "bppicker[heic,faces,raw,nudity,inpaint]" .
Desktop app (macOS, Apple Silicon)
Prefer a dock icon over a terminal? Each release ships a standalone
Mac app. No Python, no terminal:
Download BestPhotoPicker-macOS-arm64.dmg (always the latest release).
Double-click the .dmg and drag Best Photo Picker into Applications .
Open it from Applications or Spotlight.
The app is signed with an Apple Developer ID and notarized by Apple,
so it opens on double-click — no security warning. To update, download
the newer .dmg the same way; your photo library and settings live
outside the app and carry over untouched.
Apple Silicon (M1/M2/M3 and later) only. On an Intel Mac, or on
Windows/Linux, use the pipx install "bppicker[web]" path above.
bpp demo
Generates sample photos and launches the web UI — no configuration needed.
# Start the photo management server (default library: ~/Pictures/BestPhotoPicker)
bpp serve
# Or specify a custom library path
bpp serve --library ~ /Photos/2024
Open http://127.0.0.1:5001 in your browser. Import folders, adjust scoring weights with sliders, and export your curated selection.
# One-shot: analyze + select best 50 photos
bpp run --input ~ /Photos/MyKid --k 50 --out ~ /curated --gallery
# Or step-by-step:
bpp analyze --input ~ /Photos/MyKid --out ~ /workdir
bpp select --workdir ~ /workdir --k 50 --out ~ /curated --gallery
How It Works
Import → Analyze → Score → Deduplicate → Select → Export
Import : photos copied to managed library with SHA-256 dedup
Analyze : parallel feature extraction (blur, exposure, faces, composition) cached in SQLite
Score : weighted combination of sub-scores, tunable in real-time via sliders
Deduplicate : perceptual hash clustering removes burst duplicates; optional CLIP semantic dedup
Select : greedy selection with per-day caps and monthly coverage for temporal diversity
Export : copy/hardlink/symlink selected photos with optional HTML gallery
Feature
Description
Photo grid
Thumbnail grid with score badges, zoom control, sort & filter
Lightbox
Full-size viewer with keyboard navigation
Sliders
Real-time weight tuning (blur, exposure, face, composition)
BPP Picks
Sidebar sub-item picks best K photos across the full library; toolbar chip filters picks in the current album/view
Albums
Manual and smart albums (by time, score, person)
Faces
Auto-clustered face detection with merge, dismiss, and per-person albums
Batch ops
Multi-select with Cmd/Ctrl-click, bulk include/exclude/favorite
Import
Drag folders or archives into the library
Export
Copy or link selected photos to an output directory
Tuning Knobs
Parameter
Default
Description
blur_weight
0.30
Weight for sharpness in aggregate score
exposure_weight
0.20
Weight for exposure quality
face_weight
0.35
Weight for face detection & framing
composition_weight
0.15
Weight for composition (rule of thirds)
max_long_side
1024
Downscale images to this before analysis
hash_distance_threshold
10
dHash Hamming distance for "same" image
time_window_seconds
15
Cluster burst photos within this window
max_per_day
3
Max selected photos per calendar day
min_per_month
1
Try to include at least 1 per month
Hardware
bpp runs all ML inference on CPU by default. ONNX-based models (SCRFD
face detection, CLIP semantic search, YOLOv11n pet detection) can opt
into hardware acceleration through the BPP_ONNX_PROVIDERS env var:
# Apple Silicon (M1/M2/M3) — use the Apple Neural Engine via CoreML
BPP_ONNX_PROVIDERS= " CoreMLExecutionProvider,CPUExecutionProvider " \
bpp serve --library ~ /Pictures/BestPhotoPicker
# NVIDIA GPU (Linux) — install onnxruntime-gpu first
pip uninstall -y onnxruntime && pip install onnxruntime-gpu
BPP_ONNX_PROVIDERS= " CUDAExecutionProvider,CPUExecutionProvider " \
bpp serve --library ~ /Pictures/BestPhotoPicker
# Windows DirectML
BPP_ONNX_PROVIDERS= " DmlExecutionProvider,CPUExecutionProvider " \
bpp serve --library ~ /Pictures/BestPhotoPicker
CPU is always appended as the final fallback so a missing provider
in your wheel falls through cleanly with a warning rather than
crashing. Ordering matters — providers are tried in the order listed.
Surface
Tested in CI
Supported
Notes
macOS arm64 (M1/M2/M3) — CPU
yes (dev box)
yes
Primary dev target. Tauri desktop app ships only this.
Linux x86_64 — CPU
yes ( ubuntu-latest )
yes
Default for PyPI install.
Linux arm64 — CPU
no
yes
Docker image builds; runtime not exercised in CI.
Windows — CPU
no
yes
Code path exists; never run end-to-end. Reports welcome via GitHub issues.
macOS arm64 — CoreMLExecutionProvider (ANE)
no
yes
Code path exists. Expected 5-10× face inference speedup; correctness depends on the specific ONNX graph and ONNX Runtime version.
Linux x86_64 — CUDAExecutionProvider
no
yes
Requires pip install onnxruntime-gpu (replaces base onnxruntime ). Free-tier CI doesn't have a GPU runner.
Windows — DmlExecutionProvider
no
yes
Same caveat as Windows CPU — cod

[truncated]
