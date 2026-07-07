---
source: "https://github.com/s2xon/aeovim/"
hn_url: "https://news.ycombinator.com/item?id=48822567"
title: "AI Neovim"
article_title: "GitHub - s2xon/aeovim: aeovim instead if neovim · GitHub"
author: "s2xon"
captured_at: "2026-07-07T19:42:18Z"
capture_tool: "hn-digest"
hn_id: 48822567
score: 1
comments: 1
posted_at: "2026-07-07T19:32:05Z"
tags:
  - hacker-news
  - translated
---

# AI Neovim

- HN: [48822567](https://news.ycombinator.com/item?id=48822567)
- Source: [github.com](https://github.com/s2xon/aeovim/)
- Score: 1
- Comments: 1
- Posted: 2026-07-07T19:32:05Z

## Translation

タイトル：AIネオビム
記事のタイトル: GitHub - s2xon/aeovim: neovim の場合は代わりに aeovim · GitHub
説明: neovim の代わりに aeovim。 GitHub でアカウントを作成して、s2xon/aeovim の開発に貢献してください。

記事本文:
GitHub - s2xon/aeovim: neovim の場合は代わりに aeovim · GitHub
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
s2xon
/
アエオビム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
5 コミット 5 コミット src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md IMPLEMENTATION_PLAN.md IMPLEMENTATION_PLAN.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
vim ですが、バッファはライブ コーディング エージェントであり、オペレータがそれを駆動します。
aeovim は、LLM コーディング エージェントを多重化および調整するためのスタンドアロンのキーボード ネイティブ Rust TUI です。 Neovim メンタル モデル (モード、モーション、オペレーター、バッファー、タブ、分割) をコーディング エージェントとの会話に適用するため、多くのエージェントを一度に生成、操作、監視、レビューすることは、ウィンドウをジャグリングするのではなく筋肉の記憶になります。
プロジェクトは aeovim です。実行するコマンドは avim です ( Neovim -> nvim など)。
v1 は、ヘッドレス stream-json 上の長期存続子プロセスとしてクロード CLI (クロード コード) をラップします。
ビューアではなくオーケストレーター: 一流の求人掲示板での並列タスク、ファンアウト、スキル、avim 所有のループ。
]c / [c ハンク モーションとターンごとの git 承認/拒否を使用した vim ネイティブの差分レビュー。
すべてのバックエンドの詳細はアダプター シームの背後にあるため、他のモデル/CLI (または直接 API) は後でドロップインされます。
シングルユーザーのローカル macOS 日次ドライバー。配布されていません。
IMPLEMENTATION_PLAN.md — マイルストーンのラダーと構築計画。
実装前。設計が完了しました。次は歩くスケルトンのスパイクです (計画を参照)。
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

aeovim instead if neovim. Contribute to s2xon/aeovim development by creating an account on GitHub.

GitHub - s2xon/aeovim: aeovim instead if neovim · GitHub
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
s2xon
/
aeovim
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md IMPLEMENTATION_PLAN.md IMPLEMENTATION_PLAN.md README.md README.md View all files Repository files navigation
vim, but the buffers are live coding agents and the operators drive them.
aeovim is a standalone, keyboard-native Rust TUI for multiplexing and orchestrating LLM coding agents. It applies the Neovim mental model — modes, motions, operators, buffers, tabs, splits — to conversations with coding agents, so spawning, steering, watching, and reviewing many agents at once is muscle memory rather than window juggling.
The project is aeovim ; the command you run is avim (like Neovim -> nvim ).
v1 wraps the claude CLI (Claude Code) as long-lived child processes over headless stream-json .
An orchestrator, not a viewer: parallel tasks, fan-out, skills, and avim-owned loops on a first-class job board.
vim-native diff review with ]c / [c hunk motions and per-turn git approve/reject.
All backend detail sits behind an adapter seam so other models/CLIs (or a direct API) drop in later.
Single-user, local macOS daily driver. Not distributed.
IMPLEMENTATION_PLAN.md — milestone ladder and build plan.
Pre-implementation. Design complete; the walking-skeleton spike is next (see the plan).
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
