---
source: "https://github.com/innovatorved/sayitdev"
hn_url: "https://news.ycombinator.com/item?id=48882834"
title: "Show HN: SayItDev LLM and Speech capabilities with 0 dependencies and no models"
article_title: "GitHub - innovatorved/sayitdev: On-device AI and voice for Mac. dev CLI: UNIX tool, OpenAI-compatible server, speak/listen/transcribe — no cloud, no API keys. · GitHub"
author: "innovatorved"
captured_at: "2026-07-12T17:51:48Z"
capture_tool: "hn-digest"
hn_id: 48882834
score: 3
comments: 0
posted_at: "2026-07-12T17:29:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: SayItDev LLM and Speech capabilities with 0 dependencies and no models

- HN: [48882834](https://news.ycombinator.com/item?id=48882834)
- Source: [github.com](https://github.com/innovatorved/sayitdev)
- Score: 3
- Comments: 0
- Posted: 2026-07-12T17:29:07Z

## Translation

タイトル: HN を表示: 依存関係がなく、モデルがない SayItDev LLM および音声機能
記事のタイトル: GitHub - innovatorved/sayitdev: Mac 用のオンデバイス AI と音声。開発 CLI: UNIX ツール、OpenAI 互換サーバー、話す/聞く/文字起こし - クラウドも API キーもなし。 · GitHub
説明: Mac 用のオンデバイス AI と音声。開発 CLI: UNIX ツール、OpenAI 互換サーバー、話す/聞く/文字起こし - クラウドも API キーもなし。 - イノベーターブド/sayitdev

記事本文:
GitHub - innovatorved/sayitdev: Mac 用のオンデバイス AI と音声。開発 CLI: UNIX ツール、OpenAI 互換サーバー、話す/聞く/文字起こし - クラウドも API キーもなし。 · GitHub
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
革新的な
/
Sayitdev
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
593 コミット 593 コミット .claude/ ルーチン .claude/ ルーチン .github .github 例 例 ソース ソース テスト テストの完了 完了 デモ デモ ドキュメント ドキュメント man man mcp mcp スクリプト スクリプト .gitattributes .gitattributes .gitignore .gitignore .spi.yml .spi.yml .version .version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Info.plist Info.plist LICENSE LICENSE Makefile Makefile Package.resolved Package.resolved Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md STABILITY.md STABILITY.md pytest.ini pytest.ini すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Mac 用のオンデバイス AI と音声。開発 CLI は Apple Intelligence をローカルで実行し、話したり、聞いたり、文字起こししたり、チャットしたり、OpenAI 互換の API を提供したりします。クラウドも API キーもありません。
Apple Intelligence が有効 (LLM / dev "prompt" および --serve のみ)
醸造タップ革新/タップ
brew install innovatorved/tap/dev
dev --version
ソースからビルド: このリポジトリのクローンを作成し、 make build && sudo make install を実行します。トラブルシューティングについては、docs/brew-install.md を参照してください。
--listen の場合、システム設定 → プライバシーで Terminal.app にマイクと音声認識を許可します。 --transcribe および --speak の場合、音声認識/出力アクセスのみ (マイクなし)。
コマンド
何をするのか
dev --「こんにちは」と話します
テキスト読み上げ
dev --listen
マイク → トランスクリプト
dev --transcribe file.wav
音声ファイル → トランスクリプト
開発者「あなたのプロンプト」
オンデバイス LLM
dev --serve
ローカル HTTP サーバー (ポート 11434)
dev --speak "SayItDev からこんにちは"
dev --listen
dev --録音.wavの転写
開発者「2+2って何ですか？」
dev --serve
--transcribe は完全にデバイス上で実行されます。オプションのフラグ: --locale en-US 、 --timestamps (各セグメントの前に i を付けます)

ts 時間範囲)。
サーバー エンドポイントには、 /v1/chat/completions 、 /v1/audio/speech 、および /v1/audio/transcriptions が含まれます。サーバーをネットワーク上に公開する前に、docs/server-security.md を参照してください。
SayItDev / innovatorved による音声拡張機能を備えた apfel のフォーク。
Mac 用のオンデバイス AI と音声。開発 CLI: UNIX ツール、OpenAI 互換サーバー、話す/聞く/文字起こし - クラウドも API キーもなし。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
6
v1.0.5
最新の
2026 年 7 月 11 日
+ 5 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

On-device AI and voice for Mac. dev CLI: UNIX tool, OpenAI-compatible server, speak/listen/transcribe — no cloud, no API keys. - innovatorved/sayitdev

GitHub - innovatorved/sayitdev: On-device AI and voice for Mac. dev CLI: UNIX tool, OpenAI-compatible server, speak/listen/transcribe — no cloud, no API keys. · GitHub
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
innovatorved
/
sayitdev
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
593 Commits 593 Commits .claude/ routines .claude/ routines .github .github Examples Examples Sources Sources Tests Tests completions completions demo demo docs docs man man mcp mcp scripts scripts .gitattributes .gitattributes .gitignore .gitignore .spi.yml .spi.yml .version .version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Info.plist Info.plist LICENSE LICENSE Makefile Makefile Package.resolved Package.resolved Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md STABILITY.md STABILITY.md pytest.ini pytest.ini View all files Repository files navigation
On-device AI and voice for Mac. The dev CLI runs Apple Intelligence locally — speak, listen, transcribe, chat, and serve an OpenAI-compatible API. No cloud, no API keys.
Apple Intelligence enabled (for LLM / dev "prompt" and --serve only)
brew tap innovatorved/tap
brew install innovatorved/tap/dev
dev --version
Build from source: clone this repo, run make build && sudo make install . See docs/brew-install.md for troubleshooting.
For --listen , grant Microphone and Speech Recognition to Terminal.app in System Settings → Privacy. For --transcribe and --speak , Speech Recognition / output access only (no mic).
Command
What it does
dev --speak "hello"
Text to speech
dev --listen
Mic → transcript
dev --transcribe file.wav
Audio file → transcript
dev "your prompt"
On-device LLM
dev --serve
Local HTTP server (port 11434)
dev --speak " Hello from SayItDev "
dev --listen
dev --transcribe recording.wav
dev " What is 2+2? "
dev --serve
--transcribe runs fully on-device. Optional flags: --locale en-US , --timestamps (prefix each segment with its time range).
Server endpoints include /v1/chat/completions , /v1/audio/speech , and /v1/audio/transcriptions . See docs/server-security.md before exposing the server on your network.
Fork of apfel with voice extensions by SayItDev / innovatorved.
On-device AI and voice for Mac. dev CLI: UNIX tool, OpenAI-compatible server, speak/listen/transcribe — no cloud, no API keys.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
6
v1.0.5
Latest
Jul 11, 2026
+ 5 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
