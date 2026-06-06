---
source: "https://github.com/waldekmastykarz/atifact"
hn_url: "https://news.ycombinator.com/item?id=48422341"
title: "Turn HAR Files, Claude Code, Copilot CLI, and Codex CLI Logs into ATIF"
article_title: "GitHub - waldekmastykarz/atifact: Turn HAR files, Claude Code, Copilot CLI, and Codex CLI logs into standardized ATIF v1.7 trajectory JSON — ready for debugging, visualization, fine-tuning, and RL pipelines. · GitHub"
author: "waldekm"
captured_at: "2026-06-06T09:45:49Z"
capture_tool: "hn-digest"
hn_id: 48422341
score: 2
comments: 0
posted_at: "2026-06-06T07:24:54Z"
tags:
  - hacker-news
  - translated
---

# Turn HAR Files, Claude Code, Copilot CLI, and Codex CLI Logs into ATIF

- HN: [48422341](https://news.ycombinator.com/item?id=48422341)
- Source: [github.com](https://github.com/waldekmastykarz/atifact)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T07:24:54Z

## Translation

タイトル: HAR ファイル、Claude Code、Copilot CLI、および Codex CLI ログを ATIF に変換する
記事のタイトル: GitHub - waldekmastykarz/atifact: HAR ファイル、Claude Code、Copilot CLI、および Codex CLI ログを標準化された ATIF v1.7 軌跡 JSON に変換します。これにより、デバッグ、視覚化、微調整、および RL パイプラインの準備が整います。 · GitHub
説明: HAR ファイル、Claude Code、Copilot CLI、および Codex CLI ログを標準化された ATIF v1.7 軌跡 JSON に変換します。これにより、デバッグ、視覚化、微調整、および RL パイプラインの準備が整います。 - ワルデクマスティカルツ/アティファクト

記事本文:
GitHub - waldekmastykarz/atifact: HAR ファイル、Claude コード、Copilot CLI、および Codex CLI ログを標準化された ATIF v1.7 軌跡 JSON に変換します。これにより、デバッグ、視覚化、微調整、および RL パイプラインの準備が整います。 · GitHub
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
ワルデクマスティカルツ
/
アティ

事実
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
53 コミット 53 コミット .github .github 資産 アセット スキル/ atifact スキル/ atifact src src test test .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント ログを ATIF トラジェクトリに変換します。コマンドは 1 つです。依存関係ゼロ。
HAR ファイル、Claude Code ログ、Copilot CLI ログ、および Codex CLI ログを標準化された ATIF v1.7 軌跡 JSON に変換します。これにより、デバッグ、視覚化、微調整、および RL パイプラインの準備が整います。
AI コーディング エージェントに atifact スキルを与えて、ユーザーに代わって軌跡を抽出できるようにします。
npx スキル追加 waldekmastykarz/atifact
インストールしたら、エージェントに「この HAR ファイルから軌跡を抽出する」、「クロード コード ログを ATIF に変換する」、「Copilot CLI ログを変換する」、または「Codex CLI ログを変換する」ように依頼すると、残りの作業はエージェントが処理します。
npm install -g atifact
クイックスタート
# HAR ファイルを変換します (自動検出)
アティファクトセッション.har
# クロードコードのログを変換する
atifact クロードログ.jsonl
# Copilot CLI ログを変換する
atifact copilot-session.jsonl
# Codex CLI ログを変換する
atifact codex-session.jsonl
# stdout にパイプします (軌跡の JSON 配列を返します)
atifact session.har --json | jq ' .ステップ |長さ '
出力: ATIF v1.7 形式の <input>.trajectory.json。サブエージェントを含む Copilot CLI および Codex CLI ログでは、追加の <input>.trajectory.<name>.json ファイルが生成されます。
--json モードは、subagent_trajectories 配列にサブエージェントが埋め込まれた単一の軌跡を、ファイルを書き込まずに標準出力に出力します。
フォーマット
ソース
旗
ハール
OpenAI

チャット完了 API、OpenAI 応答 API、Anthropic メッセージ API
ハール
JSONL
Claude Code CLI セッション ログ
クロードコード-jsonl
JSONL
Copilot CLI セッション ログ
コパイロット-cli-jsonl
JSONL
Codex CLI exec --json ログ
コーデックス-cli-jsonl
形式はファイルの内容 (拡張子ではなく) から自動検出されます。 -f で強制的に実行します。
atifact myfile.log -f クロードコード-jsonl
atifact myfile.log -f copilot-cli-jsonl
atifact myfile.log -f codex-cli-jsonl
使用法
atifact <入力ファイル> [オプション]
オプション
説明
-o, --output <プレフィックス>
出力パスの接頭辞 (デフォルト: 入力ファイルのパス)。メイン: <prefix>.trajectory.json 、サブエージェント: <prefix>.trajectory.<name>.json
-f、--format <fmt>
強制入力形式: har 、 claude-code-jsonl 、 copilot-cli-jsonl 、 codex-cli-jsonl
--json
サブエージェントが埋め込まれた軌跡を標準出力に書き込みます (ファイルは書き込まれません)
-q、--静か
進行状況メッセージを抑制する
-h、--ヘルプ
ヘルプを表示する
--バージョン
印刷版
終了コード
コード
意味
0
成功
1
実行時エラー (解析失敗、I/O エラー)
2
無効な使用法 (引数が間違っている、ファイルが見つからない)
出力形式
atifact は、以下を含む ATIF v1.7 JSON を生成します。
ステップ - ユーザーメッセージ、エージェントの応答、ツール呼び出し、および観察
メトリクス - ステップごとのトークン数、コスト、キャッシュされたトークン
ツール呼び出し — 構造化された関数名 + 観測結果を含む引数
サブエージェントの軌跡 — Copilot CLI および Codex CLI のサブエージェントは、trajectory_id 解決を持つ subagent_trajectory_ref を介してリンクされた個別の軌跡ファイルを生成します。 --json モードでは、それらを subagent_trajectories に埋め込みます。
最終的なメトリクス — 軌跡全体の集計合計
ソース データから ISO 8601 として保存されたすべてのタイムスタンプ
コンパクトな出力では Null/未定義フィールドが除外される
atifact Recording.har -o my-trajectory
# 書き込み: my-trajectory.trajectory.json
Claude Code → 軌跡、パイプ
atifact ~ /.claude/projects/ * /sessions/ * .jsonl --

json --quiet > trajectory.json
Copilot CLI → サブエージェントを使用した軌跡
atifact copilot-session.jsonl
# 書き込み: copilot-session.jsonl.trajectory.json
# 書き込み: copilot-session.jsonl.trajectory.<サブエージェント名>.json (サブエージェントごと)
Codex CLI → 軌跡
atifact codex-session.jsonl
# 書き込み: codex-session.jsonl.trajectory.json
エージェントのステップをカウントする
atifact session.har --json | jq ' [.steps[] | select(.source == "エージェント")] |長さ '
要件
HAR ファイル、Claude Code、Copilot CLI、および Codex CLI ログを標準化された ATIF v1.7 軌跡 JSON に変換します。これにより、デバッグ、視覚化、微調整、および RL パイプラインの準備が整います。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Turn HAR files, Claude Code, Copilot CLI, and Codex CLI logs into standardized ATIF v1.7 trajectory JSON — ready for debugging, visualization, fine-tuning, and RL pipelines. - waldekmastykarz/atifact

GitHub - waldekmastykarz/atifact: Turn HAR files, Claude Code, Copilot CLI, and Codex CLI logs into standardized ATIF v1.7 trajectory JSON — ready for debugging, visualization, fine-tuning, and RL pipelines. · GitHub
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
waldekmastykarz
/
atifact
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
53 Commits 53 Commits .github .github assets assets skills/ atifact skills/ atifact src src test test .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Convert agent logs to ATIF trajectories. One command. Zero dependencies.
Turn HAR files, Claude Code logs, Copilot CLI logs, and Codex CLI logs into standardized ATIF v1.7 trajectory JSON — ready for debugging, visualization, fine-tuning, and RL pipelines.
Give your AI coding agent the atifact skill so it can extract trajectories on your behalf:
npx skills add waldekmastykarz/atifact
Once installed, ask your agent to "extract the trajectory from this HAR file" , "convert Claude Code logs to ATIF" , "convert Copilot CLI logs" , or "convert Codex CLI logs" and it will handle the rest.
npm install -g atifact
Quick start
# Convert a HAR file (auto-detected)
atifact session.har
# Convert Claude Code logs
atifact claude-log.jsonl
# Convert Copilot CLI logs
atifact copilot-session.jsonl
# Convert Codex CLI logs
atifact codex-session.jsonl
# Pipe to stdout (returns JSON array of trajectories)
atifact session.har --json | jq ' .steps | length '
Output: <input>.trajectory.json in ATIF v1.7 format. Copilot CLI and Codex CLI logs with subagents produce additional <input>.trajectory.<name>.json files.
--json mode outputs a single trajectory with subagents embedded in the subagent_trajectories array to stdout with no files written.
Format
Source
Flag
HAR
OpenAI Chat Completions API, OpenAI Responses API, Anthropic Messages API
har
JSONL
Claude Code CLI session logs
claude-code-jsonl
JSONL
Copilot CLI session logs
copilot-cli-jsonl
JSONL
Codex CLI exec --json logs
codex-cli-jsonl
Format is auto-detected from file contents (not extension). Force it with -f :
atifact myfile.log -f claude-code-jsonl
atifact myfile.log -f copilot-cli-jsonl
atifact myfile.log -f codex-cli-jsonl
Usage
atifact <input-file> [options]
Option
Description
-o, --output <prefix>
Output path prefix (default: input file path). Main: <prefix>.trajectory.json , subagents: <prefix>.trajectory.<name>.json
-f, --format <fmt>
Force input format: har , claude-code-jsonl , copilot-cli-jsonl , codex-cli-jsonl
--json
Write trajectory to stdout with subagents embedded (no files written)
-q, --quiet
Suppress progress messages
-h, --help
Show help
--version
Print version
Exit codes
Code
Meaning
0
Success
1
Runtime error (parse failure, I/O error)
2
Invalid usage (bad arguments, missing file)
Output format
atifact produces ATIF v1.7 JSON with:
Steps — user messages, agent responses, tool calls, and observations
Metrics — token counts, costs, cached tokens per step
Tool calls — structured function name + arguments with observation results
Subagent trajectories — Copilot CLI and Codex CLI subagents produce separate trajectory files linked via subagent_trajectory_ref with trajectory_id resolution; --json mode embeds them in subagent_trajectories
Final metrics — aggregated totals across the trajectory
All timestamps preserved as ISO 8601 from source data
Null/undefined fields excluded for compact output
atifact recording.har -o my-trajectory
# Writes: my-trajectory.trajectory.json
Claude Code → trajectory, piped
atifact ~ /.claude/projects/ * /sessions/ * .jsonl --json --quiet > trajectory.json
Copilot CLI → trajectory with subagents
atifact copilot-session.jsonl
# Writes: copilot-session.jsonl.trajectory.json
# Writes: copilot-session.jsonl.trajectory.<subagent-name>.json (per subagent)
Codex CLI → trajectory
atifact codex-session.jsonl
# Writes: codex-session.jsonl.trajectory.json
Count agent steps
atifact session.har --json | jq ' [.steps[] | select(.source == "agent")] | length '
Requirements
Turn HAR files, Claude Code, Copilot CLI, and Codex CLI logs into standardized ATIF v1.7 trajectory JSON — ready for debugging, visualization, fine-tuning, and RL pipelines.
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
