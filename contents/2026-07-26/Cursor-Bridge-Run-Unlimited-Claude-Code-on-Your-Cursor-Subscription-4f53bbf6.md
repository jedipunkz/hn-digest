---
source: "https://github.com/hkc5/cursor-bridge"
hn_url: "https://news.ycombinator.com/item?id=49063186"
title: "Cursor Bridge – Run Unlimited Claude Code on Your Cursor Subscription"
article_title: "GitHub - hkc5/cursor-bridge · GitHub"
author: "hakancan"
captured_at: "2026-07-26T22:50:53Z"
capture_tool: "hn-digest"
hn_id: 49063186
score: 1
comments: 1
posted_at: "2026-07-26T22:48:09Z"
tags:
  - hacker-news
  - translated
---

# Cursor Bridge – Run Unlimited Claude Code on Your Cursor Subscription

- HN: [49063186](https://news.ycombinator.com/item?id=49063186)
- Source: [github.com](https://github.com/hkc5/cursor-bridge)
- Score: 1
- Comments: 1
- Posted: 2026-07-26T22:48:09Z

## Translation

タイトル: Cursor Bridge – カーソル サブスクリプションで無制限のクロード コードを実行
記事タイトル: GitHub - hkc5/cursor-bridge · GitHub
説明: GitHub でアカウントを作成して、hkc5/cursor-bridge の開発に貢献します。

記事本文:
GitHub - hkc5/cursor-bridge · GitHub
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
hkc5
/
カーソルブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード

「その他のアクション」メニューを開く フォルダーとファイル
16 コミット 16 コミット .github/ workflows .github/ workflows src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バイナリが 1 つ。 Cursor のバックエンドのクロード コード。設定ゼロ。
Cursor サブスクリプションを持っています。クロード コード (CLI) を使用したいと考えています。
Cursor の Auto モデルはサブスクリプションに含まれており、無料、無制限、トークンごとの追加費用はありません。
カーソルブリッジを使用しない場合、Anthropic API クレジットまたは Claude Pro プランを別途支払うことになります。
カーソルブリッジを使用すると、カーソルブリッジを実行するだけで機能します。クロードコードはカーソルバックエンドで実行されます。
すでに Cursor の料金を支払っている → その上で Claude Code を無料で入手
Anthropic 課金なしで Claude Code のエージェント機能 (ファイル編集、シェル コマンド、ツールの使用) が必要な場合
Cursor の Auto モデルは無料で、サブスクリプションにより無制限になります — Claude コードは実質的に無料で実行できるようになります
カーソルブリッジ # インタラクティブセッション
カーソルブリッジ " このファイルをリファクタリング " # ワンショット プロンプト
カーソルブリッジ -p " list files " # パイプモード
それだけです。プロキシ管理はありません。環境変数はありません。すべてが自動です。
カーソルブリッジ (Rust バイナリ)
§── ランダムなポートでローカル HTTP プロキシを開始します
§── macOS キーチェーン (または Linux の CURSOR_TOKEN 環境変数) からカーソル認証トークンを読み取ります。
§── プロキシを指す環境変数を使用して `claude` を生成します
§── プロキシが Anthropic API 呼び出しを変換 → カーソルエージェント CLI
└── 終了時にクリーンアップ
プロキシが表示されません。あなたはそれを管理できません。それはそこにあり、そして消えていきます。
# 前提条件
# - `エージェント` CLI 認証済みでインストールされたカーソル (`エージェント ログイン`)
# - クロード コードがインストールされています (`curl -O https://claude-code.anthropic.com/claude && chmod +x claude`)
カーゴインストールカーソルブリッジ
# あとはそれを使うだけ

カーソルブリッジ
または、リリースからバイナリをダウンロードします。
カーソルサブスクリプション (PATH 内のエージェント CLI を使用)
クロード コード CLI (PATH 内のクロード)
macOS : キーチェーンからのトークンの自動読み取り
Linux : CURSOR_TOKEN 環境変数を設定します (キーチェーン フォールバックなし)
他のプロキシとの違い
他のすべてのソリューションは、お客様が管理するバックグラウンド サーバーです。カーソルブリッジは実行するコマンドです。
既存のプロキシ (cursor-api-proxy 、cursor-composer-in-claude 、cursor-proxy ) は、プロセス リスト内に存在し、ポートを占有し、手動の環境変数の配線が必要な Node.js サーバーです。クロード コードは同梱されていません。ユーザーとクロード コードの間に挟まれて儀式が加えられます。
カーソルブリッジはその逆です。開始、停止、または設定する必要はありません。それはセッションです:
デーモンはありません。 npm install はありません。環境変数はありません。ポートハンティングはありません。クリーンアップはありません。動作するバイナリは 1 つだけです。
カーソルブリッジはクロードを完全に置き換えます。プロキシのライフサイクルを内部で管理し、CLI を生成し、完了時にそれ自体をクリーンアップします。
Linux : CURSOR_TOKEN 環境変数が必要です (キーチェーンのサポートなし)。
ワークスペースのサンドボックス化はありません。エージェントは現在のディレクトリで実行されます。
単一アカウント — 複数アカウントのローテーションは（まだ）ありません。
このプロジェクトは Anthropic または Cursor/Anysphere とは関係ありません。ご自身の責任でご使用ください。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to hkc5/cursor-bridge development by creating an account on GitHub.

GitHub - hkc5/cursor-bridge · GitHub
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
hkc5
/
cursor-bridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits .github/ workflows .github/ workflows src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
One binary. Claude Code on Cursor's backend. Zero config.
You have a Cursor subscription . You want to use Claude Code (the CLI).
Cursor's Auto model is included with your subscription — free, unlimited, no extra per-token cost.
Without cursor-bridge, you'd pay separately for Anthropic API credits or a Claude Pro plan.
With cursor-bridge, you just run cursor-bridge and it works — Claude Code runs on your Cursor backend.
You're already paying for Cursor → get Claude Code for free on top
You want Claude Code's agent capabilities (file editing, shell commands, tool use) without Anthropic billing
Cursor's Auto model is free and unlimited with subscription — Claude Code becomes effectively free to run
cursor-bridge # interactive session
cursor-bridge " refactor this file " # one-shot prompt
cursor-bridge -p " list files " # pipe mode
That's it. No proxy management. No env vars. Everything automatic.
cursor-bridge (Rust binary)
├── Starts a local HTTP proxy on a random port
├── Reads your Cursor auth token from macOS keychain (or CURSOR_TOKEN env var on Linux)
├── Spawns `claude` with env vars pointing at the proxy
├── Proxy translates Anthropic API calls → Cursor agent CLI
└── Cleans up on exit
You don't see the proxy. You don't manage it. It's there and gone.
# Prerequisites
# - Cursor installed with `agent` CLI authenticated (`agent login`)
# - Claude Code installed (`curl -O https://claude-code.anthropic.com/claude && chmod +x claude`)
cargo install cursor-bridge
# Then just use it
cursor-bridge
Or download a binary from Releases.
Cursor subscription (with agent CLI in PATH)
Claude Code CLI ( claude in PATH)
macOS : token auto-read from keychain
Linux : set CURSOR_TOKEN env var (no keychain fallback)
How it differs from other proxies
All other solutions are background servers you manage. cursor-bridge is a command you run.
Existing proxies ( cursor-api-proxy , cursor-composer-in-claude , cursor-proxy ) are Node.js servers that live in your process list, occupy a port, and need manual env var wiring. They don't ship with Claude Code — they sit between you and it, adding ceremony.
cursor-bridge is the opposite. There is nothing to start, stop, or configure. It is the session:
No daemon. No npm install . No env vars. No port hunting. No cleanup. Just a single binary that works.
cursor-bridge replaces claude entirely — it manages the proxy lifecycle internally, spawns the CLI, and cleans up after itself when you're done.
Linux : requires CURSOR_TOKEN env var (no keychain support).
No workspace sandboxing — the agent runs in your current directory.
Single account — no multi-account rotation (yet).
This project is not affiliated with Anthropic or Cursor/Anysphere. Use at your own risk.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
