---
source: "https://github.com/altic-dev/FluidVoice"
hn_url: "https://news.ycombinator.com/item?id=48739409"
title: "FluidVoice - Open source voice-to-text dictation app for macOS with local AI"
article_title: "GitHub - altic-dev/FluidVoice: Fastest and only macOS Dictation app with on-device STT and custom trained AI enhancement model - Local Wispr Flow alternative. One ⭐ takes us a long way :)) Windows, iOS and Linux coming soon. · GitHub"
author: "danboarder"
captured_at: "2026-06-30T21:34:04Z"
capture_tool: "hn-digest"
hn_id: 48739409
score: 1
comments: 0
posted_at: "2026-06-30T21:24:33Z"
tags:
  - hacker-news
  - translated
---

# FluidVoice - Open source voice-to-text dictation app for macOS with local AI

- HN: [48739409](https://news.ycombinator.com/item?id=48739409)
- Source: [github.com](https://github.com/altic-dev/FluidVoice)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T21:24:33Z

## Translation

タイトル: FluidVoice - ローカル AI を備えた macOS 用のオープンソースの音声からテキストへのディクテーション アプリ
記事のタイトル: GitHub - altic-dev/FluidVoice: オンデバイス STT とカスタム トレーニングされた AI 強化モデルを備えた最速かつ唯一の macOS ディクテーション アプリ - ローカル Wispr Flow の代替品。 1 つの ⭐ は長い道のりです :)) Windows、iOS、Linux は近日公開予定です。 · GitHub
説明: オンデバイス STT とカスタムトレーニングされた AI 強化モデルを備えた最速かつ唯一の macOS ディクテーション アプリ - ローカル Wispr Flow の代替品。 1 つの ⭐ は長い道のりです :)) Windows、iOS、Linux は近日公開予定です。 - アルティック開発/FluidVoice

記事本文:
GitHub - altic-dev/FluidVoice: オンデバイス STT とカスタム トレーニングされた AI 強化モデルを備えた最速かつ唯一の macOS ディクテーション アプリ - ローカル Wispr Flow の代替品。 1 つの ⭐ は長い道のりです :)) Windows、iOS、Linux は近日公開予定です。 · GitHub
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
アルティ

c-dev
/
流体音声
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
751 コミット 751 コミット .github .github Fluid.xcodeproj Fluid.xcodeproj ソース/ Fluid ソース/ Fluid テスト/ FluidDictationIntegrationTests テスト/ FluidDictationIntegrationTests アセット アセット ドキュメント ドキュメント スクリプト スクリプト .gitattributes .gitattributes .gitignore .gitignore .swiftformat .swiftformat .swiftlint.yml .swiftlint.yml Fluid.entitlements Fluid.entitlements Info.plist Info.plist LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md build.sh build.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オンデバイス AI 強化を備えた macOS 用のオープンソースの音声からテキストへのディクテーション アプリ。
Homebrew でインストール: brew install --cask Fluidvoice
マニュアルダウンロード: 最新リリース
このプロジェクトは、GPLv3 に基づく無料のオープンソースです。 FluidVoice が役立つ場合は、リポジトリにスターを付けてください。可視化に役立ち、開発を継続できます。
FluidVoice がお役に立てば、GitHub スポンサー で継続的な開発と iOS および Windows の将来のプラットフォーム作業をサポートできます。
めちゃくちゃ速い Parakeet — 画面上で言葉を話してから見るまでの遅延がほぼゼロで再構築された Parakeet 実装
Fluid Intelligence — オンデバイスのディクテーション強化のための完全にローカルな AI モデル。クラウドも API キーも、Mac からデータが流出することもありません
より良いテーマ — コンパクトなツールバースイッチャーを備えた適応型のライト/ダークテーマ
更新されたオンボーディング — 言語優先の音声エンジンのセットアップ、実際のディクテーションのトライアウト、および AI 強化のセットアップを 1 つのクリーンなパスで実行
初期のフィードバックに基づいて、Fluid Intelligence により次のようなことが起こる可能性があります。

他のディクテーション アプリを購読してお金を節約しましょう。警告されましたね。
FluidVoice は GPLv3 に基づく完全なオープンソースです。 Fluid Intelligence は、非公開で管理される別個のローカル AI ランタイムであり、高度なオンデバイス ディクテーション機能強化 (スマート フォーマット、コンテキストを認識した大文字化、後処理) をすべて Mac 上でローカルに実行します。
このアプリは、サポートされている音声モデルやオプションのクラウド AI プロバイダーと組み合わせて単独でうまく機能します。 Fluid Intelligence は、データをどこにも送信せずにデバイス上の機能強化を必要とするユーザー向けに、完全にローカルなプライベート AI レイヤーを追加します。
核となるディクテーション エクスペリエンスを無料で継続的に提供できるよう、現時点では Fluid Intelligence を非公開にしています。これは将来変更される可能性があります。
メールテンプレート
花
変更時間名.mp4
電子メール_テンプレート.mp4
時間と名前の変更
絵文字
絵文字.mp4
花.mp4
ハイフンと数字
ハイフン123.mp4
デモ
コマンド モード — FluidVoice を使用して Mac 上であらゆるアクションを実行します
FluidVoice_1.5_cmd_mode.mp4
書き込みモード — 任意のアプリの任意のテキスト ボックスにテキストを書き込むか書き換えます
writemode_FluidVoice_1.5.mp4
スクリーンショット
Fluid Intelligence — スマート フォーマッティング、コンテキストを認識した大文字化、後処理のためのオンデバイス AI の強化。すべて、マシンからデータを出さずに Mac 上でローカルに実行されます。
コマンド モード — 音声で Mac を制御します。キーボードに触れることなく、アプリの起動、ショートカットの実行、システム アクションのトリガー、ワークフローの自動化を行います。
書き込みモード — あらゆるアプリのあらゆるテキスト フィールドにテキストを直接書き込みまたは書き換えます。テキストを選択して書き直すか、新しいコンテンツをインラインで口述筆記します
ライブ プレビュー — ノッチ サポートを備えたリアルタイムの文字起こしオーバーレイにより、話しているときに単語が表示されるのがわかります
複数の音声モデル — Nemotron Speech 3.5、Parakeet Flash、Parakeet TDT v3 および v2、Cohere Transcribe、Apple Speech、Whisper。というモデルを選んでください

言語と遅延のニーズに適合します
AI の強化 — OpenAI、Groq、カスタム プロバイダー、またはローカルの Fluid Intelligence によるオプションの後処理により、よりクリーンで正確なトランスクリプトを実現します。
オーディオ履歴 — 予算管理と ZIP エクスポートを備えたオプションのローカル録音履歴により、クラウド ストレージなしで過去のディクテーションを確認できます
今日の使用状況の統計 — 統計ヘッダー カードとツールバー ピルを使用して、毎日の使用状況を一目で追跡できます。
アダプティブ テーマ — コンパクトなツールバー スイッチャーを備えた、システムに追従する明るい/暗いテーマ
グローバル ホットキー — どこからでも瞬時に音声をキャプチャでき、アプリの切り替えは不要です
スマート タイピング — アクセシビリティ API を介して任意のアプリに直接挿入し、アプリに依存しない信頼性の高いテキスト入力を実現します。
メニュー バーの統合 — メニュー バーからのクイック アクセス、ステータス、設定
自動更新 — 初期プレビュー用のオプションのベータ チャネルを使用したシームレスな更新
アプリごとの構成 — 異なるアプリに異なるプロンプト セットを割り当てると、作業内容に合わせてディクテーションが適応されます。完全にオプションです。
ノッチ対応オーバーレイ — MacBook のノッチの周囲にきれいにフィットする転写オーバーレイ。Mac に標準オーバーレイがない場合は標準オーバーレイを使用します。
ローカルファースト — クラウド AI プロバイダーにオプトインしない限り、音声とテキストがマシンから離れることはありません
Mac 上で最速の Parakeet — macOS 上での Parakeet の最速ネイティブ実装の 1 つで、ほぼ瞬時の転写と最小限の遅延を実現します。
構成可能なオーバーレイ — 錠剤型から大きなオーバーレイ サイズまで選択してライブ プレビューを表示するか、最小限に抑えます。すべてはオプションです
すべてがオプションです - AI 強化、Fluid Intelligence、オーディオ履歴、分析、ベータ版ビルドはすべてオプトインです。コアのディクテーション エクスペリエンスは、アクセス許可とホットキー以外の設定を必要とせず、すぐに使用できます。
モデル
こんな方に最適
言語補助

ポート
ダウンロードサイズ
ハードウェア
Nemotron Speech 3.5 — 超高速、低遅延
ストリーミング対応の多言語ディクテーション
約40の言語
~670MB
アップルシリコン
Nemotron 3.5 多言語対応
より高精度な多言語ディクテーション
約40の言語
～530MB
アップルシリコン
インコ フラッシュ (ベータ版)
遅延が最も少ないライブ英語ディクテーション
英語
～250MB
アップルシリコン
インコ TDT v3
デフォルトの高速多言語ディクテーション
25の言語
~500MB
アップルシリコン
インコ TDT v2
最速の英語のみのディクテーション
英語
~500MB
アップルシリコン
コヒア転写
高精度の多言語ディクテーション
14言語
～1.4GB
アップルシリコン
アップルスピーチ
ダウンロード不要のネイティブ macOS 音声
システム言語
内蔵
Appleシリコン+インテル
ウィスパーティニー/ベース/スモール/ミディアム/ラージ
Intel Mac を含む幅広い互換性
99の言語
～75MB ～ ～2.9GB
Appleシリコン+インテル
インコ TDT v3 言語
ブルガリア語、クロアチア語、チェコ語、デンマーク語、オランダ語、英語、エストニア語、フィンランド語、フランス語、ドイツ語、ギリシャ語、ハンガリー語、イタリア語、ラトビア語、リトアニア語、マルタ語、ポーランド語、ポルトガル語、ルーマニア語、ロシア語、スロバキア語、スロベニア語、スペイン語、スウェーデン語、ウクライナ語。
英語、フランス語、ドイツ語、イタリア語、スペイン語、ポルトガル語、ギリシャ語、オランダ語、ポーランド語、北京語、日本語、韓国語、ベトナム語、アラビア語。
システム言語のサポートは、マシンで利用可能な macOS 音声認識言語によって異なります。
Whisper は、選択したモデルのサイズに応じて、最大 99 の言語をサポートします。
brew install --cask Fluidvoice
または、最新リリースをダウンロードしてください。
許可を与える — FluidVoice はマイクとアクセシビリティへのアクセスを要求します。どちらもディクテーションや他のアプリへの入力に必要です。
ホットキーを設定します — どこからでも音声キャプチャをトリガーするグローバル ホットキーを設定で選択します。
オンボーディングを実行します。言語と遅延のニーズに基づいて音声モデルを選択します。モデル

ダウンロード不要の Apple Speech から高精度の Nemotron や Whisper まで多岐にわたります。
(オプション) Fluid Intelligence を有効にする — オンデバイスのディクテーション強化のため、オンボーディング中にローカル AI モデルをダウンロードします。すべてがローカルで実行され、Mac からデータが流出することはありません。
(オプション) 独自の AI プロバイダーを導入する — クラウドベースの拡張のために OpenAI、Groq、またはカスタム プロバイダー API キーを追加します。キーは macOS キーチェーンに安全に保存されます。キーアクセスを「常に許可」を選択します。
(オプション) ベータ ビルドにオプトインします。新機能に早期にアクセスするには、[設定] → [自動更新] → [ベータ リリース] を選択します。
Apple Silicon Mac 全モデル対応
Whisper モデルでサポートされる Intel Mac (1.5.1 以降)
音声モデルの場合は最大 1 GB のディスク容量
Fluid Intelligence モデル用の最大 3.5 GB のディスク容量 (オプション)
入力のためのアクセシビリティ権限
git クローン https://github.com/altic-dev/FluidVoice.git
cd 流体音声
Fluid.xcodeproj を開く
Xcode でビルドして実行します。すべての依存関係は Swift Package Manager を通じて管理されます。
xcodebuild -project Fluid.xcodeproj -scheme Fluid -destination ' platform=macOS ' build CODE_SIGNING_ALLOWED=NO
貢献する
貢献は大歓迎です!プル リクエストを送信する前に、まず問題を作成して大きな変更について話し合ってください。
上記のようにクローンを作成して Xcode で開きます。
署名: FluidVoice → 署名と機能 → 署名を自動的に管理 → チームを選択します (個人チームでも問題ありません)。 xcuserdata/ (gitignored) に保存されます。
ビルドして実行 — SPM が依存関係を処理します。
(オプション) 誤ってチーム ID がコミットされるのを防ぐための事前コミットフック:
cp scripts/check-team-id.sh .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
PR ごとに 1 つの機能または修正 — 変更を集中的かつアトミックに保つ
最初に問題を作成して、レビュー前に作業を追跡できるようにします
PR を開始する前に、重要な変更について話し合う
マシン上で徹底的にテストする
決してコミットしないでください

個人のチーム ID または API キー
コミットする前に git diff を確認する
xcodebuild test -project Fluid.xcodeproj -scheme Fluid -destination ' platform=macOS '
CI は署名されていないビルドを使用します。
xcodebuild test -project Fluid.xcodeproj -scheme Fluid -destination ' platform=macOS ' CODE_SIGNING_REQUIRED=NO CODE_SIGNING_ALLOWED=NO
プライバシーと分析
FluidVoice はローカルファーストです。クラウド AI プロバイダーを明示的にオプトインしない限り、音声、オーディオ、文字起こしされたテキストがマシンから離れることはありません。
アプリの健全性と機能の使用状況を追跡するために、匿名分析がデフォルトで有効になっています。 [設定] → [匿名分析の共有] からいつでも無効にできます。
アプリのバージョン、ビルド、macOS バージョン
低カーディナリティの機能/構成フラグ (アプリ モード、主要設定など)
おおよその使用範囲（正確な値ではありません）
成功/エラーの高レベルの結果
音声、生の音声、または文字に起こしたテキスト
選択したテキスト、プロンプト、または AI 応答
ターミナル コマンド、ウィンドウ タイトル、ファイル パス、クリップボード、または入力されたコンテンツ
あらゆる個人情報または個人情報
Discord に参加してください: https://discord.gg/VUPHaKSvYV
X の開発をフォローしてください: @ALTIC_DEV
2026 年 2 月 23 日以降、このプロジェクトは GNU General Public License、バージョン 3.0 (GPLv3) に基づいてライセンスされます。
この日付より前に公開されたバージョンは、Apache License 2.0 に基づいてライセンスされていました。
最速かつ唯一の macOS Dictati

[切り捨てられた]

## Original Extract

Fastest and only macOS Dictation app with on-device STT and custom trained AI enhancement model - Local Wispr Flow alternative. One ⭐ takes us a long way :)) Windows, iOS and Linux coming soon. - altic-dev/FluidVoice

GitHub - altic-dev/FluidVoice: Fastest and only macOS Dictation app with on-device STT and custom trained AI enhancement model - Local Wispr Flow alternative. One ⭐ takes us a long way :)) Windows, iOS and Linux coming soon. · GitHub
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
altic-dev
/
FluidVoice
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
751 Commits 751 Commits .github .github Fluid.xcodeproj Fluid.xcodeproj Sources/ Fluid Sources/ Fluid Tests/ FluidDictationIntegrationTests Tests/ FluidDictationIntegrationTests assets assets docs docs scripts scripts .gitattributes .gitattributes .gitignore .gitignore .swiftformat .swiftformat .swiftlint.yml .swiftlint.yml Fluid.entitlements Fluid.entitlements Info.plist Info.plist LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md build.sh build.sh View all files Repository files navigation
Open source voice-to-text dictation app for macOS with on-device AI enhancement.
Install with Homebrew: brew install --cask fluidvoice
Manual download: latest release
This project is free and open source under GPLv3. If FluidVoice is useful to you, please star the repository — it helps visibility and keeps development going.
If FluidVoice helps you, you can support continued development and future platform work for iOS and Windows on GitHub Sponsors .
Insanely fast Parakeet — rebuilt Parakeet implementation with pretty much zero delay between speaking and seeing words on screen
Fluid Intelligence — fully local AI model for on-device dictation enhancement. No cloud, no API keys, no data leaving your Mac
Better Theming — adaptive light/dark theme with a compact toolbar switcher
Refreshed Onboarding — language-first voice engine setup, real dictation tryout, and AI enhancement setup in one clean pass
Based on early feedback, Fluid Intelligence may cause you to unsubscribe from other dictation apps and save money. You've been warned.
FluidVoice is fully open source under GPLv3. Fluid Intelligence is a separate, privately maintained local AI runtime that powers advanced on-device dictation enhancement — smart formatting, context-aware capitalization, and post-processing — all running locally on your Mac.
The app works great on its own with any supported speech model and optional cloud AI providers. Fluid Intelligence adds a fully local, private AI layer for users who want on-device enhancement without sending data anywhere.
We're keeping Fluid Intelligence private for now so we can sustainably offer the core dictation experience for free. This may change in the future.
Email Template
Flowers
change_time_name.mp4
Email_template.mp4
Change Time & Name
Emoji
emoji.mp4
flowers.mp4
Hyphens & Numbers
hyphen123.mp4
Demo
Command Mode — Take any action on your Mac using FluidVoice
FluidVoice_1.5_cmd_mode.mp4
Write Mode — Write or rewrite text in any text box in any app
writemode_FluidVoice_1.5.mp4
Screenshots
Fluid Intelligence — on-device AI enhancement for smart formatting, context-aware capitalization, and post-processing, all running locally on your Mac with zero data leaving your machine
Command Mode — control your Mac by voice: launch apps, run shortcuts, trigger system actions, and automate workflows without touching the keyboard
Write Mode — write or rewrite text directly in any text field across any app. Select text and rewrite it, or dictate new content inline
Live Preview — real-time transcription overlay with notch support, so you see words appear as you speak
Multiple Speech Models — Nemotron Speech 3.5, Parakeet Flash, Parakeet TDT v3 & v2, Cohere Transcribe, Apple Speech, and Whisper. Pick the model that fits your language and latency needs
AI Enhancement — optional post-processing via OpenAI, Groq, custom providers, or local Fluid Intelligence for cleaner, more accurate transcripts
Audio History — optional local recording history with budget controls and ZIP export, so you can review past dictations without cloud storage
Today-Usage Stats — daily usage tracking at a glance with a stats header card and toolbar pill
Adaptive Theming — light/dark theme that follows your system, with a compact toolbar switcher
Global Hotkey — instant voice capture from anywhere, no app switching needed
Smart Typing — direct insertion into any app via accessibility APIs for reliable, app-independent text entry
Menu Bar Integration — quick access, status, and settings from the menu bar
Auto-Updates — seamless updates with an optional beta channel for early previews
Per-App Configuration — assign different prompt sets to different apps, so your dictation adapts to whatever you're working in. Fully optional
Notch-Aware Overlay — transcription overlay that fits cleanly around the MacBook notch, or use a standard overlay if your Mac doesn't have one
Local-First — your voice and text never leave your machine unless you opt in to a cloud AI provider
Fastest Parakeet on Mac — one of the fastest native implementations of Parakeet on macOS, with near-instant transcription and minimal latency
Configurable Overlay — choose from pill-shaped to large overlay sizes to show live preview, or keep it minimal. Everything is optional
Everything is Optional — AI enhancement, Fluid Intelligence, audio history, analytics, and beta builds are all opt-in. The core dictation experience works out of the box with zero configuration beyond permissions and a hotkey
Model
Best for
Language support
Download size
Hardware
Nemotron Speech 3.5 — Ultra Fast Low Latency
Streaming-capable multilingual dictation
~40 languages
~670 MB
Apple Silicon
Nemotron 3.5 Multilingual
Higher-accuracy multilingual dictation
~40 languages
~530 MB
Apple Silicon
Parakeet Flash (Beta)
Lowest-latency live English dictation
English
~250 MB
Apple Silicon
Parakeet TDT v3
Fast default multilingual dictation
25 languages
~500 MB
Apple Silicon
Parakeet TDT v2
Fastest English-only dictation
English
~500 MB
Apple Silicon
Cohere Transcribe
High-accuracy multilingual dictation
14 languages
~1.4 GB
Apple Silicon
Apple Speech
Zero-download native macOS speech
System languages
Built-in
Apple Silicon + Intel
Whisper Tiny / Base / Small / Medium / Large
Broad compatibility, including Intel Macs
99 languages
~75 MB to ~2.9 GB
Apple Silicon + Intel
Parakeet TDT v3 Languages
Bulgarian, Croatian, Czech, Danish, Dutch, English, Estonian, Finnish, French, German, Greek, Hungarian, Italian, Latvian, Lithuanian, Maltese, Polish, Portuguese, Romanian, Russian, Slovak, Slovenian, Spanish, Swedish, and Ukrainian.
English, French, German, Italian, Spanish, Portuguese, Greek, Dutch, Polish, Mandarin, Japanese, Korean, Vietnamese, and Arabic.
System language support depends on the macOS speech recognition languages available on your machine.
Whisper supports up to 99 languages, depending on the model size you choose.
brew install --cask fluidvoice
Or download the latest release .
Grant permissions — FluidVoice will ask for microphone and accessibility access. Both are required for dictation and typing into other apps.
Set your hotkey — pick a global hotkey in settings that triggers voice capture from anywhere.
Go through onboarding — choose your voice model based on your language and latency needs. Models range from zero-download Apple Speech to high-accuracy Nemotron and Whisper.
(Optional) Enable Fluid Intelligence — download the local AI model during onboarding for on-device dictation enhancement. Everything runs locally, no data leaves your Mac.
(Optional) Bring your own AI provider — add an OpenAI, Groq, or custom provider API key for cloud-based enhancement. Keys are stored securely in macOS Keychain. Select "Always allow" for key access.
(Optional) Opt in to beta builds — Settings → Automatic Updates → Beta Releases for early access to new features.
Apple Silicon Mac for all models
Intel Macs supported via Whisper models (from 1.5.1+)
~1 GB disk space for a voice model
~3.5 GB disk space for the Fluid Intelligence model (optional)
Accessibility permissions for typing
git clone https://github.com/altic-dev/FluidVoice.git
cd FluidVoice
open Fluid.xcodeproj
Build and run in Xcode. All dependencies are managed via Swift Package Manager.
xcodebuild -project Fluid.xcodeproj -scheme Fluid -destination ' platform=macOS ' build CODE_SIGNING_ALLOWED=NO
Contributing
Contributions are welcome! Please create an issue first to discuss major changes before submitting a pull request.
Clone and open in Xcode as above.
Signing: FluidVoice → Signing & Capabilities → Automatically manage signing → pick your Team (Personal Team is fine). Stored in xcuserdata/ (gitignored).
Build and run — SPM handles dependencies.
(Optional) Pre-commit hook to prevent accidental team ID commits:
cp scripts/check-team-id.sh .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
One feature or fix per PR — keep changes focused and atomic
Create an issue first so work is trackable before review
Discuss non-trivial changes before opening a PR
Test thoroughly on your machine
Never commit personal team IDs or API keys
Check git diff before committing
xcodebuild test -project Fluid.xcodeproj -scheme Fluid -destination ' platform=macOS '
CI uses unsigned builds:
xcodebuild test -project Fluid.xcodeproj -scheme Fluid -destination ' platform=macOS ' CODE_SIGNING_REQUIRED=NO CODE_SIGNING_ALLOWED=NO
Privacy & Analytics
FluidVoice is local-first . Your voice, audio, and transcribed text never leave your machine unless you explicitly opt in to a cloud AI provider.
Anonymous analytics are enabled by default to track app health and feature usage. You can disable at any time from Settings → Share Anonymous Analytics .
App version, build, macOS version
Low-cardinality feature/config flags (e.g. app mode, major settings)
Approximate usage ranges (not exact values)
High-level success/error outcomes
Voice, raw audio, or transcribed text
Selected text, prompts, or AI responses
Terminal commands, window titles, file paths, clipboard, or typed content
Any personal or private information
Join our Discord: https://discord.gg/VUPHaKSvYV
Follow development on X: @ALTIC_DEV
From 2026-02-23 onward, this project is licensed under the GNU General Public License, Version 3.0 (GPLv3) .
Versions published before this date were licensed under Apache License 2.0.
Fastest and only macOS Dictati

[truncated]
