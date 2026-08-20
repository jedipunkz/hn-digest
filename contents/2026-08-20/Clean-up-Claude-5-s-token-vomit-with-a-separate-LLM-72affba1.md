---
source: "https://github.com/zachahn/vomit"
hn_url: "https://news.ycombinator.com/item?id=49375996"
title: "Clean up Claude 5's token vomit with a separate LLM"
article_title: "GitHub - zachahn/vomit: Clean up Claude 5's token vomit with a separate LLM. Save your tokens, Claude 5 is hopeless · GitHub"
image: "https://opengraph.githubassets.com/d3f8cbfffbf6cb01fd61c9aff33266947fc23e3ed74610995c43c2f63ed961bd/zachahn/vomit"
author: "Bluestein"
captured_at: "2026-08-20T16:23:22Z"
capture_tool: "hn-digest"
hn_id: 49375996
score: 34
comments: 22
posted_at: "2026-08-20T15:26:02Z"
tags:
  - hacker-news
  - translated
---

# Clean up Claude 5's token vomit with a separate LLM

- HN: [49375996](https://news.ycombinator.com/item?id=49375996)
- Source: [github.com](https://github.com/zachahn/vomit)
- Score: 34
- Comments: 22
- Posted: 2026-08-20T15:26:02Z

## Translation

タイトル: 別の LLM を使用してクロード 5 のトークンの吐き出しをクリーンアップ
記事のタイトル: GitHub - zachahn/vomit: Claude 5 のトークンの嘔吐を別の LLM でクリーンアップします。トークンを保存してください。クロード 5 は絶望的です · GitHub
説明: クロード 5 のトークン嘔吐を別の LLM でクリーンアップします。トークンを保存してください、クロード 5 は絶望的です - ザカーン/嘔吐

記事本文:
GitHub - zachahn/vomit: Claude 5 のトークン嘔吐を別の LLM でクリーンアップします。トークンを保存してください。クロード 5 は絶望的です · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ザカーン
/
嘔吐する
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
28 コミット 28 コミット フォルダーとファイル
cmd/ decomment cmd/ decomment 内部 内部 .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md Rakefile Rakefile go.mod go.mo

d init.go init.go list.go list.go main.go main.go main_test.go main_test.go printer.go printer.go scrub.go scrub.go scrub_test.go scrub_test.go tail.go tail.go version.go version.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Vomit は、ローカル LLM にパイプして、Claude のトークン vomit を英語に変換します。これは完全にローカル (テレメトリなし) であり、外部依存関係はありません。
ローカル LLM は、クロードが通信しようとしている内容しか見ることができない (アクションやファイルにはアクセスできない) ため、少し幻覚が見えます。
これは完全にバイブコーディングされており、Mac でのみテストされています
クロードのメッセージを完全に見逃してしまう可能性があります。 Vomit は実行時に何も操作しないため (技術的には TMPDIR にファイルを書き込むことを除いて)、AgentsView などを使用して元のメッセージを取得できます。
# バイナリを GOPATH にインストールします
github.com/zachahn/vomit@latest をインストールしてください
# LLM への接続の詳細を設定します
嘔吐初期
# フックを介してクロードの出力を置き換える方法についての説明
嘔吐スクラブ - クロード
使用法
さらに、Vomit を並行して実行したい場合は、非侵襲モードがあります。
嘔吐リスト。クロードセッション識別子のリスト
尾を吐きます [<セッション識別子>] 。指定されたセッションのクロードのトークンを変換するか、最新のトークンを追跡します
嘔吐の助け。その他のコマンドについては、「 」を参照してください。
OpenAI APIを使用するものはありますか?
ローカル LLM がまだセットアップされていない場合は、 Llama.app を使用し、それから GPT-OSS 20B をダウンロードすることをお勧めします。これを設定した後、init サブコマンドを実行します。
クロード 5 のトークンの嘔吐を別の LLM でクリーンアップします。トークンを節約してください、クロード 5 は絶望的です
Readme GPL-3.0 ライセンス アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Clean up Claude 5's token vomit with a separate LLM. Save your tokens, Claude 5 is hopeless - zachahn/vomit

GitHub - zachahn/vomit: Clean up Claude 5's token vomit with a separate LLM. Save your tokens, Claude 5 is hopeless · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
zachahn
/
vomit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
28 Commits 28 Commits Folders and files
cmd/ decomment cmd/ decomment internal internal .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md Rakefile Rakefile go.mod go.mod init.go init.go list.go list.go main.go main.go main_test.go main_test.go printer.go printer.go scrub.go scrub.go scrub_test.go scrub_test.go tail.go tail.go version.go version.go View all files Repository files navigation
Vomit converts Claude's token vomit into English by piping it through a local LLM. It's fully local (no telemetry) and has no external dependencies.
The local LLM can only see what Claude tries to communicate (no access to any actions or files), so it hallucinates a bit
This is totally vibe-coded, only tested on Mac
There is a possibility you'll completely miss Claude's message. You can use something like AgentsView to get the original messages, as Vomit does not touch anything during runtime (except technically it writes files to your TMPDIR)
# Install the binary to your GOPATH
go install github.com/zachahn/vomit@latest
# Setup connection details to your LLM
vomit init
# Instructions on how to replace Claude's output via hooks
vomit scrub -claude
Usage
In addition, there's a non-invasive mode if you want to run Vomit on the side.
vomit list . List Claude session identifiers
vomit tail [<session_identifier>] . Translate Claude's tokens for the specified session, or follow the latest one
vomit help . See for more commands
Anything that uses the OpenAI API?
If you don't have a local LLM set up already, I think I recommend using Llama.app , then downloading GPT-OSS 20B through it. Run the init subcommand after you set this up.
Clean up Claude 5's token vomit with a separate LLM. Save your tokens, Claude 5 is hopeless
Readme GPL-3.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
