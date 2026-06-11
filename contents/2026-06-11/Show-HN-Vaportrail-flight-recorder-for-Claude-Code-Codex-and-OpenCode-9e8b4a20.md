---
source: "https://github.com/B33BMO/vaportrail"
hn_url: "https://news.ycombinator.com/item?id=48486482"
title: "Show HN: Vaportrail – flight recorder for Claude Code, Codex, and OpenCode"
article_title: "GitHub - B33BMO/vaportrail: Your agents leave trails. vaportrail reads them — a local-first flight recorder for AI coding agent sessions. · GitHub"
author: "b33bmo"
captured_at: "2026-06-11T07:50:30Z"
capture_tool: "hn-digest"
hn_id: 48486482
score: 2
comments: 1
posted_at: "2026-06-11T05:15:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Vaportrail – flight recorder for Claude Code, Codex, and OpenCode

- HN: [48486482](https://news.ycombinator.com/item?id=48486482)
- Source: [github.com](https://github.com/B33BMO/vaportrail)
- Score: 2
- Comments: 1
- Posted: 2026-06-11T05:15:42Z

## Translation

タイトル: Show HN: Vaportrail – Claude Code、Codex、OpenCode 用のフライト レコーダー
記事のタイトル: GitHub - B33BMO/vaportrail: エージェントが証跡を残します。 Vaportrail はそれらを読み取ります。これは、AI コーディング エージェント セッション用のローカル ファーストのフライト レコーダーです。 · GitHub
説明: エージェントは痕跡を残します。 Vaportrail はそれらを読み取ります。これは、AI コーディング エージェント セッション用のローカル ファーストのフライト レコーダーです。 - B33BMO/ヴァポートレール

記事本文:
GitHub - B33BMO/vaportrail: エージェントは証跡を残します。 Vaportrail はそれらを読み取ります。これは、AI コーディング エージェント セッション用のローカル ファーストのフライト レコーダーです。 · GitHub
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
B33BMO
/
ヴェイパートレイル
公共
通知
通知設定を変更するにはサインインする必要があります

ngs
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット デモ デモ src src .gitignore .gitignore ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
エージェントは痕跡を残します。ヴェイパートレイルはそれらを読みます。
AI コーディング エージェント セッション用のローカル ファーストのフライト レコーダー。すべての Claude Code 、 Codex CLI 、 opencode 、 Gemini CLI 、および Aider セッションは、ディスク上に詳細なトランスクリプトを残します。つまり、ユーザーが尋ねたこと、エージェントが行ったこと、エージェントが触れたすべてのファイル、実行されたすべてのコマンド、書き込まれたすべてのトークンです。 Vaportrail は 5 つのフォーマットをすべて読み取り、そのトランスクリプトの山を検索可能で再生可能な 1 つの履歴に変換します。
実行時の依存関係はゼロです。マシンからは何も残りません。試してみるコマンドは 1 つです。
AI エージェントは現在、リポジトリ内の作業の重要な役割を担っていますが、その履歴は書き込み専用であり、ギガバイト単位のトランスクリプトは誰も読み取ることができません。 Vaportrail は、データが常に答えることができた次の質問に答えます。
エージェントは先週の火曜日に実際に何をしましたか? → ヴェイパートレイルショー <id>
この正確な問題を以前どこで解決したのでしょうか? →vaportrail検索「jwtリフレッシュ」
エージェントが最も頻繁に使用するファイルはどれですか? 1ヶ月の使用感はどんな感じですか？ → ヴェイパートレイル統計
すべてのプロジェクトの最近のセッション — ID、年齢、プロジェクト、AI 生成のタイトル、プロンプト数、ツール呼び出し、出力トークン、モデル。
セッションをタイムラインとして再生します。プロンプト、エージェントのナレーション、および操作内容を含むすべてのツール呼び出しです。
ヴェイパートレイル · セッション cca29f96…
タイトル エージェント登録による WebSocket 経由の SSH の構築
プロジェクト ~/wssh (メイン)
いつ 2026-06-02 22:14 → 2026-06-03 00:04 (1 時間 49 分)
アクティビティ 16 プロンプト · 183 ターン · 168 ツール呼び出し · opus-4-8
22:14:30

❯ それで.. 素敵なアイデアがあります。つまり… SSH .. ただし WebSocket 経由…
22:14:47 ⚒ Bash ls -la /Users/b/wssh && command -v go node python3…
22:15:40 ✦ そうですね — これは素晴らしいアイデアで、非常に構築可能です…
22:17:03 ⚒ /Users/b/wssh/go.mod を書き込みます
22:17:15 ⚒ /Users/b/wssh/main.go を書き込みます
...
切り詰められていないテキストには --full を使用します。
エージェント履歴全体 (プロンプトと応答) の全文検索。パターンには --regex、コマンドとファイル パスにも一致する --tools を使用します。
全体像: セッション、プロンプト、ターン、トークン合計 (キャッシュ分割あり)、30 日間のアクティビティ スパークライン、ツール使用率のリーダーボード、モデルごとのトークン テーブル、エージェントが最も頻繁に編集するファイル。
旗
意味
-s、--source <s>
読み取るエージェント: claude 、 codex 、 opencode (カンマリスト; デフォルトはすべて)
-p、--プロジェクト<s>
プロジェクト パスに <s> ( . = 現在のディレクトリ) が含まれるセッションのみ
-n、--limit <n>
最大行数/一致数
-a、--すべて
制限なし
-j、--json
すべてのコマンドの機械可読出力
-d、--dir <パス>
トランスクリプト ルート オーバーライド (非クロード レイアウトの場合は -s と組み合わせます)
--エージェント
サブエージェントの記録を含める
仕組み
Vaportrail は、フック、ラッパー、テレメトリなしで、各エージェントのネイティブのディスク上の形式を読み取ります。
すべてが 1 つのセッション モデルに正規化されます。トークンの使用はフォーマットによってエントリ間で繰り返され、サブエージェントの作業は独自のプロンプトから分離され、ツール名はエージェント間で正規化され (shell_command と bash は両方とも Bash としてカウントされます)、簿記のノイズがフィルターで除去されます。約 100 MB の履歴が約 1 秒で解析されます。読み取り専用: Vaportrail はトランスクリプトを変更しません。
Codex CLI および opencode トランスクリプト形式
カーソル CLI (セッションは SQLite 内に存在します。node:sqlite が必要です)
Vaportrail エクスポート <id> — セッション → マークダウン / 共有可能な要点
セッション/モデルごとのコストの見積もり
ローカル Web UI (va

portrail ui ) クロスセッション分析付き
監視モード — 実行中のセッションのライブテール
エージェントは痕跡を残します。 Vaportrail はそれらを読み取ります。これは、AI コーディング エージェント セッション用のローカル ファーストのフライト レコーダーです。
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

Your agents leave trails. vaportrail reads them — a local-first flight recorder for AI coding agent sessions. - B33BMO/vaportrail

GitHub - B33BMO/vaportrail: Your agents leave trails. vaportrail reads them — a local-first flight recorder for AI coding agent sessions. · GitHub
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
B33BMO
/
vaportrail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits demo demo src src .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Your agents leave trails. vaportrail reads them.
A local-first flight recorder for AI coding agent sessions. Every Claude Code , Codex CLI , opencode , Gemini CLI , and Aider session leaves a detailed transcript on disk — what you asked, what the agent did, every file it touched, every command it ran, every token it burned. vaportrail reads all five formats and turns that pile of transcripts into one searchable, replayable history.
Zero runtime dependencies. Nothing leaves your machine. One command to try it:
AI agents now do a meaningful share of the work in your repos, but their history is write-only: gigabytes of transcripts nobody can read. vaportrail answers the questions that data was always able to answer:
What did the agent actually do last Tuesday? → vaportrail show <id>
Where did we solve this exact problem before? → vaportrail search "jwt refresh"
Which files do agents churn on the most? What does a month of usage look like? → vaportrail stats
Recent sessions across every project — id, age, project, AI-generated title, prompt count, tool calls, output tokens, model.
Replay a session as a timeline: your prompts, the agent's narration, and every tool call with what it operated on.
vaportrail · session cca29f96…
title Build SSH over WebSocket with agent enrollment
project ~/wssh (main)
when 2026-06-02 22:14 → 2026-06-03 00:04 (1h 49m)
activity 16 prompts · 183 turns · 168 tool calls · opus-4-8
22:14:30 ❯ So.. I have a cool ass idea. So... SSH.. but over websocket…
22:14:47 ⚒ Bash ls -la /Users/b/wssh && command -v go node python3…
22:15:40 ✦ Hell yes — this is a great idea, and it's very buildable…
22:17:03 ⚒ Write /Users/b/wssh/go.mod
22:17:15 ⚒ Write /Users/b/wssh/main.go
...
Use --full for untruncated text.
Full-text search across your entire agent history — prompts and responses, with --regex for patterns and --tools to also match commands and file paths.
The big picture: sessions, prompts, turns, token totals (with cache split), a 30-day activity sparkline, a tool-usage leaderboard, per-model token tables, and the files your agents edit the most.
flag
meaning
-s, --source <s>
agents to read: claude , codex , opencode (comma list; default all)
-p, --project <s>
only sessions whose project path contains <s> ( . = current dir)
-n, --limit <n>
max rows / matches
-a, --all
no limit
-j, --json
machine-readable output for every command
-d, --dir <path>
transcript root override (combine with -s for non-claude layouts)
--agents
include subagent transcripts
How it works
vaportrail reads each agent's native on-disk format — no hooks, no wrappers, no telemetry:
Everything is normalized into one session model: token usage is deduplicated where formats repeat it across entries, subagent work is separated from your own prompts, tool names are normalized across agents ( shell_command and bash both count as Bash ), and bookkeeping noise is filtered out. ~100 MB of history parses in about a second. Read-only: vaportrail never modifies a transcript.
Codex CLI and opencode transcript formats
Cursor CLI (sessions live in SQLite; needs node:sqlite )
vaportrail export <id> — session → markdown / shareable gist
Cost estimates per session/model
Local web UI ( vaportrail ui ) with cross-session analytics
Watch mode — live tail of a running session
Your agents leave trails. vaportrail reads them — a local-first flight recorder for AI coding agent sessions.
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
