---
source: "https://github.com/jmuncor/tokentap"
hn_url: "https://news.ycombinator.com/item?id=48810046"
title: "Tokentap – Print what your LLM is thinking"
article_title: "GitHub - jmuncor/tokentap: Intercept LLM API traffic and visualize token usage in a real-time terminal dashboard. Track costs, debug prompts, and monitor context window usage across your AI development sessions. · GitHub"
author: "jmuncor"
captured_at: "2026-07-06T21:22:44Z"
capture_tool: "hn-digest"
hn_id: 48810046
score: 1
comments: 1
posted_at: "2026-07-06T20:27:02Z"
tags:
  - hacker-news
  - translated
---

# Tokentap – Print what your LLM is thinking

- HN: [48810046](https://news.ycombinator.com/item?id=48810046)
- Source: [github.com](https://github.com/jmuncor/tokentap)
- Score: 1
- Comments: 1
- Posted: 2026-07-06T20:27:02Z

## Translation

タイトル: Tokentap – LLM が考えていることを出力します
記事のタイトル: GitHub - jmuncor/tokentap: LLM API トラフィックをインターセプトし、リアルタイムのターミナル ダッシュボードでトークンの使用状況を視覚化します。 AI 開発セッション全体でコストを追跡し、プロンプトをデバッグし、コンテキスト ウィンドウの使用状況を監視します。 · GitHub
説明: LLM API トラフィックを傍受し、リアルタイムのターミナル ダッシュボードでトークンの使用状況を視覚化します。 AI 開発セッション全体でコストを追跡し、プロンプトをデバッグし、コンテキスト ウィンドウの使用状況を監視します。 - GitHub - jmuncor/tokentap: LLM API トラフィックをインターセプトし、リアルタイムのターミナル ダッシュボードでトークンの使用状況を視覚化します。
[切り捨てられた]

記事本文:
GitHub - jmuncor/tokentap: LLM API トラフィックをインターセプトし、リアルタイムのターミナル ダッシュボードでトークンの使用状況を視覚化します。 AI 開発セッション全体でコストを追跡し、プロンプトをデバッグし、コンテキスト ウィンドウの使用状況を監視します。 · GitHub
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
ジュムンコル
/
トークンタップ

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
42 コミット 42 コミット .github/ workflows .github/ workflows testing testing tokentap tokentap .gitignore .gitignore LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM CLI ツールのトークン トラッカー
インストール •
クイックスタート •
特徴 •
コマンド •
貢献する
tokentap は、ライブ ターミナル ダッシュボードを使用して、LLM CLI ツールのトークンの使用状況を追跡します。使用しているトークンの数をリアルタイムで正確に確認できます。
トークンの使用状況を追跡 : 各リクエストが消費するトークンの数を正確に確認します。
コンテキスト ウィンドウを監視: 視覚的な残量ゲージで制限に対する累積使用量を表示
デバッグ プロンプト : すべてのプロンプトをレビュー用にマークダウンおよび JSON として自動的に保存します。
ゼロ構成: 証明書もセットアップも不要 - インストールするだけで使用可能
pip インストールトークンタップ
またはソースからインストールします。
git clone https://github.com/jmuncor/tokentap.git
CDトークンタップ
pip install -e 。
要件
ターミナル 1: ダッシュボードを開始する
トークンタップスタート
キャプチャしたプロンプトを保存する場所を選択するように求められ、ダッシュボードが表示されます。
┌───────────────────────────┐
│ TOKENTAP - LLM トラフィックインスペクター │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ コンテキストの使用率 ████████████░░░░░░░░░░░░░░░░ 42% │
│ (84,231 / 200,000

トークン) │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ タイムプロバイダーモデルトークン │
│ 14:23:01 人間のクロード・ソネット-4-20250514 12,847 │
│ 14:23:45 人間のクロード・ソネット-4-20250514 8,234 │
│ 14:24:12 人間のクロード・ソネット-4-20250514 15,102 │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ 最後のプロンプト: 「この関数のリファクタリングを手伝ってくれませんか...」 │
━━━━━━━━━━━━━━━━━━━━━━━━┘
ターミナル 2: LLM ツールを実行する
# クロードコードの場合
トークンタップ クロード
# Gemini CLI の場合 (既知の問題を参照)
トークンタップジェミニ
# OpenAI コーデックスの場合
トークンタップコーデックス
# MiniMax を利用したツールの場合
tokentap run --provider minimax python my_app.py
それです！作業中にリアルタイムでダッシュボードが更新されるのを確認します。
色分けされた残量計によるリアルタイムのトークン追跡:
インターセプトされたすべてのリクエストは、選択したディレクトリに保存されます。
Markdown - メタデータを含む人間が読める形式
JSON - デバッグ用の生の API リクエスト本文
終了すると、合計使用量が表示されます。
セッションが完了しました。合計: 12 リクエスト全体で 84,231 トークン。
コマンド
コマンド
説明
トークンタップスタート
プロキシとダッシュボードを開始する
トークンタップ クロード
プロキシを構成してクロード コードを実行する
トークンタップジェミニ
プロキシを構成して Gemini CLI を実行する
トークンタップコーデックス
プロキシを構成して OpenAI Codex CLI を実行する
tokentap run --provider <名前> <cmd>
プロキシを構成して任意のコマンドを実行する
--provider でサポートされているプロバイダー:

人間 、 オープンアイ 、 ジェミニ 、 ミニマックス
トークンタップスタート [オプション]
オプション:
-p, --port NUM プロキシ ポート (デフォルト: 8080)
-l, --limit 燃料ゲージの NUM トークン制限 (デフォルト: 200000)
tokentap クロード [オプション] [ARGS]...
オプション:
-p, --port NUM プロキシ ポート (デフォルト: 8080)
仕組み
┌───────────────────────────┐
│ ターミナル 1: tokentap スタート │
│ ┌───────────────────────────┐│
│ │ HTTP プロキシ (localhost:8080) ││
│ │ + ダッシュボード ││
│ │ + プロンプトアーカイブ ││
│ ━━━━━━━━━━━━━━━━━━━━┘│
━━━━━━━━━━━━━━━━━━━━━━┬───────────────┘
│ HTTP
│
┌───────────────────┴─────────────────┐
│ ターミナル 2: トークンタップ クロード │
│ ┌───────────────────────────┐│
│ │ ANTHROPIC_BASE_URL=http://localhost:8080 を設定します ││
│ │ 実行: クロード ││
│ ━─────────

───────────────────────┘│
━━━━━━━━━━━━━━━━━━━━━━━━━┘
│
│ HTTPS
▼
┌─────────────┐
│ api.anthropic.com │
━━━━━━━━━━┘
MiniMax のような OpenAI 互換プロバイダーの場合、 tokentap は
リクエストが正しいアップストリーム API に転送されるように、パス プレフィックス ルーティングを使用します。
tokentap run --provider minimax python my_app.py
→ OPENAI_BASE_URL=http://localhost:8080/minimax/v1 を設定します
→ リクエストは /minimax/v1/chat/completions に到着します
→ プロキシはプレフィックスを取り除き、https://api.minimax.io/v1/chat/completions に転送します。
サポートされているプロバイダー
プロバイダー
コマンド
ステータス
人間性 (クロード・コード)
トークンタップ クロード
サポートされています
Google (Gemini CLI)
トークンタップジェミニ
上流の問題によりブロックされました
OpenAI (コーデックス)
トークンタップコーデックス
サポートされています
ミニマックス
tokentap run --provider minimax <cmd>
サポートされています
既知の問題
現在、Gemini CLI には、OAuth 認証の使用時にカスタム ベース URL が無視されるという既知の問題があります。 Gemini CLI チームがこの問題を修正すると、tokentap の Gemini サポートは自動的に機能するようになります。
貢献は大歓迎です！お手伝いできる方法は次のとおりです。
機能ブランチを作成します ( git checkout -b feature/amazing-feature )
変更をコミットします ( git commit -m 'Addすばらしい機能' )
ブランチにプッシュする ( git Push Origin feature/amazing-feature )
git clone https://github.com/jmuncor/tokentap.git
CDトークンタップ
Python -m venv venv
ソース venv/bin/activate
pip install -e 。
ライセンス
このプロジェクトは MIT ライセンスに基づいてライセンスされています - ライセンス fi を参照してください。

詳細については、「」を参照してください。
実際に LLM に送信されるものを確認してください。追跡。学ぶ。最適化する。
LLM API トラフィックを傍受し、リアルタイムのターミナル ダッシュボードでトークンの使用状況を視覚化します。 AI 開発セッション全体でコストを追跡し、プロンプトをデバッグし、コンテキスト ウィンドウの使用状況を監視します。
github.com/jmuncor/tokentap
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
37
フォーク
レポートリポジトリ
リリース
2
v0.1.1
最新の
2026 年 4 月 3 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Intercept LLM API traffic and visualize token usage in a real-time terminal dashboard. Track costs, debug prompts, and monitor context window usage across your AI development sessions. - GitHub - jmuncor/tokentap: Intercept LLM API traffic and visualize token usage in a real-time terminal dashboard.
[truncated]

GitHub - jmuncor/tokentap: Intercept LLM API traffic and visualize token usage in a real-time terminal dashboard. Track costs, debug prompts, and monitor context window usage across your AI development sessions. · GitHub
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
jmuncor
/
tokentap
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
42 Commits 42 Commits .github/ workflows .github/ workflows tests tests tokentap tokentap .gitignore .gitignore LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Token Tracker for LLM CLI Tools
Installation •
Quick Start •
Features •
Commands •
Contributing
tokentap tracks token usage for LLM CLI tools with a live terminal dashboard. See exactly how many tokens you're using in real-time.
Track Token Usage : See exactly how many tokens each request consumes
Monitor Context Windows : Visual fuel gauge shows cumulative usage against your limit
Debug Prompts : Automatically saves every prompt as markdown and JSON for review
Zero Configuration : No certificates, no setup - just install and go
pip install tokentap
Or install from source:
git clone https://github.com/jmuncor/tokentap.git
cd tokentap
pip install -e .
Requirements
Terminal 1: Start the Dashboard
tokentap start
You'll be prompted to choose where to save captured prompts, then the dashboard appears:
┌─────────────────────────────────────────────────────────────┐
│ TOKENTAP - LLM Traffic Inspector │
├─────────────────────────────────────────────────────────────┤
│ Context Usage ████████████░░░░░░░░░░░░░░░░ 42% │
│ (84,231 / 200,000 tokens) │
├─────────────────────────────────────────────────────────────┤
│ Time Provider Model Tokens │
│ 14:23:01 Anthropic claude-sonnet-4-20250514 12,847 │
│ 14:23:45 Anthropic claude-sonnet-4-20250514 8,234 │
│ 14:24:12 Anthropic claude-sonnet-4-20250514 15,102 │
├─────────────────────────────────────────────────────────────┤
│ Last Prompt: "Can you help me refactor this function..." │
└─────────────────────────────────────────────────────────────┘
Terminal 2: Run Your LLM Tool
# For Claude Code
tokentap claude
# For Gemini CLI (see known issues)
tokentap gemini
# For OpenAI Codex
tokentap codex
# For MiniMax-powered tools
tokentap run --provider minimax python my_app.py
That's it! Watch the dashboard update in real-time as you work.
Real-time token tracking with color-coded fuel gauge:
Every intercepted request is saved to your chosen directory:
Markdown - Human-readable format with metadata
JSON - Raw API request body for debugging
When you exit, see your total usage:
Session complete. Total: 84,231 tokens across 12 requests.
Commands
Command
Description
tokentap start
Start the proxy and dashboard
tokentap claude
Run Claude Code with proxy configured
tokentap gemini
Run Gemini CLI with proxy configured
tokentap codex
Run OpenAI Codex CLI with proxy configured
tokentap run --provider <name> <cmd>
Run any command with proxy configured
Supported providers for --provider : anthropic , openai , gemini , minimax
tokentap start [OPTIONS]
Options:
-p, --port NUM Proxy port (default: 8080)
-l, --limit NUM Token limit for fuel gauge (default: 200000)
tokentap claude [OPTIONS] [ARGS]...
Options:
-p, --port NUM Proxy port (default: 8080)
How It Works
┌─────────────────────────────────────────────────────────────────┐
│ Terminal 1: tokentap start │
│ ┌─────────────────────────────────────────────────────────────┐│
│ │ HTTP Proxy (localhost:8080) ││
│ │ + Dashboard ││
│ │ + Prompt Archive ││
│ └─────────────────────────────────────────────────────────────┘│
└───────────────────────────────┬─────────────────────────────────┘
│ HTTP
│
┌───────────────────────────────┴─────────────────────────────────┐
│ Terminal 2: tokentap claude │
│ ┌─────────────────────────────────────────────────────────────┐│
│ │ Sets ANTHROPIC_BASE_URL=http://localhost:8080 ││
│ │ Runs: claude ││
│ └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
│
│ HTTPS
▼
┌───────────────────┐
│ api.anthropic.com │
└───────────────────┘
For OpenAI-compatible providers like MiniMax , tokentap uses
path-prefix routing so requests are forwarded to the correct upstream API:
tokentap run --provider minimax python my_app.py
→ sets OPENAI_BASE_URL=http://localhost:8080/minimax/v1
→ requests arrive at /minimax/v1/chat/completions
→ proxy strips prefix, forwards to https://api.minimax.io/v1/chat/completions
Supported Providers
Provider
Command
Status
Anthropic (Claude Code)
tokentap claude
Supported
Google (Gemini CLI)
tokentap gemini
Blocked by upstream issue
OpenAI (Codex)
tokentap codex
Supported
MiniMax
tokentap run --provider minimax <cmd>
Supported
Known Issues
Gemini CLI currently has a known issue where it ignores custom base URLs when using OAuth authentication. tokentap's Gemini support will work automatically once the Gemini CLI team fixes this issue.
Contributions are welcome! Here's how you can help:
Create a feature branch ( git checkout -b feature/amazing-feature )
Commit your changes ( git commit -m 'Add amazing feature' )
Push to the branch ( git push origin feature/amazing-feature )
git clone https://github.com/jmuncor/tokentap.git
cd tokentap
python -m venv venv
source venv/bin/activate
pip install -e .
License
This project is licensed under the MIT License - see the LICENSE file for details.
See what's really being sent to the LLM. Track. Learn. Optimize.
Intercept LLM API traffic and visualize token usage in a real-time terminal dashboard. Track costs, debug prompts, and monitor context window usage across your AI development sessions.
github.com/jmuncor/tokentap
Topics
There was an error while loading. Please reload this page .
37
forks
Report repository
Releases
2
v0.1.1
Latest
Apr 3, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
