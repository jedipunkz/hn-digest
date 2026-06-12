---
source: "https://github.com/jpatel3/claude-dashboard"
hn_url: "https://news.ycombinator.com/item?id=48505103"
title: "Claude dashboard – local web interface"
article_title: "GitHub - jpatel3/claude-dashboard: Local web dashboard for Claude Code sessions: live monitoring, history, usage charts, resume commands, and LLM usage analysis · GitHub"
author: "jpatel3"
captured_at: "2026-06-12T16:07:18Z"
capture_tool: "hn-digest"
hn_id: 48505103
score: 2
comments: 0
posted_at: "2026-06-12T15:04:09Z"
tags:
  - hacker-news
  - translated
---

# Claude dashboard – local web interface

- HN: [48505103](https://news.ycombinator.com/item?id=48505103)
- Source: [github.com](https://github.com/jpatel3/claude-dashboard)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T15:04:09Z

## Translation

タイトル: クロード ダッシュボード – ローカル Web インターフェイス
記事のタイトル: GitHub - jpatel3/claude-dashboard: クロード コード セッションのローカル Web ダッシュボード: ライブ モニタリング、履歴、使用状況グラフ、再開コマンド、および LLM 使用状況分析 · GitHub
説明: クロード コード セッション用のローカル Web ダッシュボード: ライブ モニタリング、履歴、使用状況グラフ、再開コマンド、LLM 使用状況分析 - jpatel3/claude-dashboard

記事本文:
GitHub - jpatel3/claude-dashboard: クロード コード セッション用のローカル Web ダッシュボード: ライブ モニタリング、履歴、使用状況グラフ、再開コマンド、LLM 使用状況分析 · GitHub
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
jpatel3
/
クロードダッシュボード
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
29 コミット 29 コミット client client docs docs server server .gitignore .gitignore ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Node 20+、macOS (ライブセッション検出に ps/lsof を使用)、および Claude Code CLI が必要です。
このマシン上のクロード コード セッションのローカル ダッシュボード: ライブ セッション モニタリング、履歴、使用状況グラフ、ワンクリック再開コマンド、オンデマンド LLM 使用状況分析。
スクリーンショットはサンプルデータを示しています。名前付きセッション ( /rename 経由) にはカスタム名が表示されます。それ以外はすべてセッションの最初のプロンプトに戻ります。
npmインストール
npm 実行開発
http://localhost:5173 を開きます。 APIサーバーは:3456で動作します。
~/.claude/projects/*/*.jsonl からトランスクリプトを読み取ります (増分、追加のみの解析)。
クロード プロセス ( ps + lsof cwd) をトランスクリプト ディレクトリと照合することで、ライブ セッションを検出します。プロセス検出が利用できない場合は、「過去 2 分間に変更されました」に戻ります。
「使用状況を分析する」は、過去 14 日間のダイジェストに対して claude -p headless を実行し、レポートを server/data/last-analysis.json にキャッシュします。
サーバーは 127.0.0.1 のみにバインドし (0.0.0.0 にはバインドしない)、ホスト許可リストを強制して DNS 再バインド攻撃を防ぎます。
POST /api/analyze には、x-claude-dashboard: 1 ヘッダー (UI が自動的に送信します) が必要です。ヘッダーのない裸のカールは 403 で拒否されます。
npmテスト
すべてのテストはserver/test/に存在し、vitestで実行されます。
仕様: docs/superpowers/specs/2026-06-11-claude-dashboard-design.md
同じディレクトリの同時セッション: 同じ作業ディレクトリ内の 2 つの同時クロード セッションが 1 つのアクティブなセッションとして表示されます — cwd (lsof から) がオンになっています

ly シグナルは、プロセスとトランスクリプトを照合するために使用されるため、そのディレクトリ内で最も最近アクティブだったトランスクリプトが優先されます。
サブエージェントのトランスクリプトはカウントされません:projects/*/<session>/subagents/ に保存されているトランスクリプトはスキャンされません。トークンの使用量とセッション数には、トップレベルのセッション ファイルのみが反映されます。
セッションと統計パネルはページ読み込みごとに 1 回読み込まれます。[最近のセッション] テーブルとグラフ パネルは起動時にデータを 1 回取得し、自動更新はしません。新しいデータを表示するにはページを再読み込みします。 [アクティブ セッション] パネルは 5 秒ごとに自動ポーリングします。
クロード コード セッション用のローカル Web ダッシュボード: ライブ モニタリング、履歴、使用状況グラフ、再開コマンド、LLM 使用状況分析
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

Local web dashboard for Claude Code sessions: live monitoring, history, usage charts, resume commands, and LLM usage analysis - jpatel3/claude-dashboard

GitHub - jpatel3/claude-dashboard: Local web dashboard for Claude Code sessions: live monitoring, history, usage charts, resume commands, and LLM usage analysis · GitHub
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
jpatel3
/
claude-dashboard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
29 Commits 29 Commits client client docs docs server server .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Requires Node 20+, macOS (uses ps/lsof for live-session detection), and the Claude Code CLI.
Local dashboard for Claude Code sessions on this machine: live session monitoring, history, usage charts, one-click resume commands, and on-demand LLM usage analysis.
Screenshot shows sample data. Named sessions (via /rename ) display their custom name; everything else falls back to the session's first prompt.
npm install
npm run dev
Open http://localhost:5173 . API server runs on :3456.
Reads transcripts from ~/.claude/projects/*/*.jsonl (incremental, append-only parsing).
Detects live sessions by matching claude processes ( ps + lsof cwd) to transcript directories; falls back to "modified in last 2 minutes" if process detection is unavailable.
"Analyze my usage" runs claude -p headless over a digest of the last 14 days and caches the report to server/data/last-analysis.json .
The server binds to 127.0.0.1 only (never 0.0.0.0) and enforces a Host allowlist to prevent DNS-rebinding attacks.
POST /api/analyze requires the x-claude-dashboard: 1 header (the UI sends it automatically); bare curl without the header is rejected with 403.
npm test
All tests live in server/test/ and are run with vitest.
Spec: docs/superpowers/specs/2026-06-11-claude-dashboard-design.md
Same-directory concurrent sessions: Two concurrent Claude sessions in the same working directory appear as a single active session — cwd (from lsof ) is the only signal used to match a process to a transcript, so the most-recently-active transcript in that directory wins.
Subagent transcripts not counted: Transcripts stored under projects/*/<session>/subagents/ are not scanned. Token usage and session counts reflect only top-level session files.
Sessions and stats panels load once per page load: The Recent Sessions table and Charts panel fetch data once at startup and do not auto-refresh — reload the page to see new data. The Active Sessions panel auto-polls every 5 seconds.
Local web dashboard for Claude Code sessions: live monitoring, history, usage charts, resume commands, and LLM usage analysis
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
