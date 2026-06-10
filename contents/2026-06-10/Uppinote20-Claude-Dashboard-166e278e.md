---
source: "https://github.com/uppinote20/claude-dashboard"
hn_url: "https://news.ycombinator.com/item?id=48473677"
title: "Uppinote20 Claude-Dashboard"
article_title: "GitHub - uppinote20/claude-dashboard: Comprehensive status line plugin for Claude Code with context usage, API rate limits, and cost tracking · GitHub"
author: "geoffbp"
captured_at: "2026-06-10T10:29:39Z"
capture_tool: "hn-digest"
hn_id: 48473677
score: 1
comments: 0
posted_at: "2026-06-10T09:21:45Z"
tags:
  - hacker-news
  - translated
---

# Uppinote20 Claude-Dashboard

- HN: [48473677](https://news.ycombinator.com/item?id=48473677)
- Source: [github.com](https://github.com/uppinote20/claude-dashboard)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T09:21:45Z

## Translation

タイトル: Uppinote20 クロード ダッシュボード
記事のタイトル: GitHub - uppinote20/claude-dashboard: コンテキストの使用法、API レート制限、コスト追跡を備えたクロード コード用の包括的なステータス ライン プラグイン · GitHub
説明: コンテキストの使用法、API レート制限、コスト追跡を備えたクロード コード用の包括的なステータス ライン プラグイン - uppinote20/claude-dashboard

記事本文:
GitHub - uppinote20/claude-dashboard: コンテキストの使用法、API レート制限、コスト追跡を備えたクロード コード用の包括的なステータス ライン プラグイン · GitHub
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
アップピノート20
/
クロードダッシュボード
公共
通知
通知を変更するにはサインインする必要があります

カチオン設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
363 コミット 363 コミット .claude-plugin .claude-plugin .github .github コマンド コマンド dist dist イメージ 画像 ロケール ロケール スクリプト スクリプト Web サイト Web サイト .gitignore .gitignore CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.mdデモ.gif デモ.gif パッケージ-lock.json パッケージ-lock.json パッケージ.json パッケージ.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code 用の包括的なステータス ライン プラグイン — コンテキスト、レート制限、コスト追跡、モジュラー ウィジェット システムを備えた、Claude、Codex、Gemini、および z.ai CLI にわたる統合使用状況監視。
要件: クロード コード v1.0.80+ / Node.js 18+
/プラグイン マーケットプレイス add uppinote20/claude-dashboard
/プラグインインストール クロードダッシュボード
/claude-ダッシュボード:セットアップ
手動インストール
git clone https://github.com/uppinote20/claude-dashboard.git ~ /.claude/plugins/claude-dashboard
/claude-ダッシュボード:セットアップ
表示モード
モデル、コンテキスト進行状況バー、コスト、レート制限 (5 時間/7 日/7 日ソネット) または z.ai の使用法 (プロバイダーによって相互に排他的)
プロジェクト情報、セッションID、セッション期間、バーンレート、ToDoの進行状況を追加します
枯渇時間、構成数、ツール/エージェントのステータス、キャッシュ ヒット、パフォーマンス バッジ、トークンの内訳、予測、予算、Codex/Gemini の使用状況を追加します
マルチプロバイダーのサポート: z.ai/ZHIPU、Codex、Gemini はインストール時に自動検出されます。
カテゴリ
ウィジェット
説明
コア
モデル
絵文字付きモデル名、オーパス/ソネットのエフォートレベル（X/H/M/L）、オーパスの高速モード（↯）
コンテキスト
プログレスバー、パーセンテージ、トークン (🟢 0-50% / 🟡 51-80% / 🔴 81-100%)
コンテキストバー
プログレスバーのみ (コンテキストのサブウィジェット)
コンテキストパーセンテージ
パーセンテージのみ (サブ幅

コンテキストの取得)
コンテキストの使用法
トークン数のみ、例: 42K/200K (コンテキストのサブウィジェット)
コスト
セッション費用 (USD)
プロジェクト情報
ディレクトリ + git ブランチ (クリック可能な OSC8 リンク) + 前後 (↑↓)、project_dir からのサブパス、ワークツリー インジケーター (🌳)
レート制限
rateLimit5h
リセットカウントダウン付きの 5 時間のレート制限
rateLimit7d
7 日間のレート制限 (Pro および Max)
rateLimit7dソネット
7 日間のソネット制限 (最大のみ)
セッション
セッションID
セッションID (短い8文字)
セッションID完全
セッションID (完全なUUID)
セッション名
/rename コマンドからのセッション名
セッション期間
セッション期間
最後のプロンプト
タイムスタンプ付きの最後のユーザー プロンプト
構成カウント
CLAUDE.md、AGENTS.md、ルール、MCP、フック、+Dirs
アクティビティ
ツールアクティビティ
ターゲットを含む実行中/完了したツール (例: Read(app.ts) )
エージェントステータス
サブエージェントの進捗状況
todo進捗状況
Todo完了率
分析
燃焼率
1分あたりのトークン消費量
キャッシュヒット
キャッシュヒット率のパーセンテージ
枯渇時間
レート制限までの推定時間 (およそ)¹
マルチCLI
コーデックスの使用法
OpenAI Codex CLI の使用法 (インストールされていない場合は自動非表示)²
ジェミニの使い方
Google Gemini CLI - 現在のモデル (インストールされていない場合は自動的に非表示になります)¶
ジェミニの使い方すべて
Google Gemini CLI - すべてのモデル (インストールされていない場合は自動非表示)¶
zai使い方
z.ai/ZHIPU の使用法 (z.ai を使用していない場合は自動非表示)⁴
洞察
トークンの内訳
入力/出力/キャッシュ書き込み/読み取りトークンの内訳
パフォーマンス
複合効率バッジ (キャッシュ ヒット + 出力率)
予測
セッション料金に基づく推定時間当たりコスト
予算
1 日の支出と設定された予算制限⁵
トークンスピード
出力トークン生成速度 (tok/s)
今日のコスト
今日のすべてのセッションにわたる合計支出額
ステータス
ピーク時間
カウントダウン付きのピーク時間インジケーター (PeakClaude に基づく)⁶
タグステータス
一致した git タグより先にコミットします (デフォルトのパターン v* 、 c
[切り捨てられた]
i18n: 英語と韓国語をサポート (自動検出またはセットアップ経由で設定)

。
# プリセットモード
/claude-dashboard:コンパクト # 1 行のセットアップ (デフォルト)
/claude-dashboard:通常のプロ # 2 行のセットアップ、英語、プロ プラン
/claude-dashboard:setup 詳細な ko max # 6 行、韓国語、最大プラン
# カスタムモード: ウィジェットの順序と行構成を制御します
# 形式: "widget1,widget2,...|widget3,widget4,..." (| は行を区切ります)
/claude-dashboard:カスタム自動最大値の設定 "モデル、コンテキスト、コスト|プロジェクト情報、todoProgress"
プランの違い:
対話型モード: 引数なしで /claude-dashboard:setup を実行します。プリセットの選択に最適です。ウィジェットを完全に制御するには、ダイレクト モードを使用するか、JSON ファイルを直接編集します。
表示モードのプリセット (zaiUsage と rateLimit* はプロバイダーに基づいて相互に排他的です):
設定ファイル ( ~/.claude/claude-dashboard.local.json ):
{
"言語" : "自動" ,
"計画" : " 最大 " 、
"displayMode" : "カスタム" ,
「行」: [
[ " モデル " 、 " コンテキスト " 、 " コスト " 、 " rateLimit5h " ]、
[ " projectInfo " 、 " todoProgress " ]
]、
"テーマ" : "デフォルト" ,
"セパレータ" : "パイプ" ,
"毎日の予算" : 15 、
"無効なウィジェット" : [],
"キャッシュ" : { "ttlSeconds" : 60 }
}
または、プリセットの省略表現を使用して素早い設定を行うこともできます。
{
"プリセット" : " MC$R|BDO " ,
"テーマ" : "東京ナイト" ,
"区切り文字" : "ドット"
}
テーマ：default（パステル）/minimal（モノクロ）/catppuccin/catppuccinLatte（ライトモード）/dracula/gruvbox/nord/tokyoNight/solarized
区切り文字: パイプ (│、デフォルト) / スペース / ドット (・) / 矢印 (›)
プリセット ショートカット: 単一文字によるクイック レイアウト — "プリセット": "MC$R|BDO" (M=model、C=context、$=cost、R=rateLimit5h など)
予算の追跡: 毎日の支出を追跡するには、「dailyBudget」: 15 を設定します。 80% で ⚠️ 、95% で 🚨 を表示します。
タグステータス: "tagPatterns": ["v*", "release-*"] を介して tagStatus パターンをカスタマイズします。デフォルトは ["v*"] です。到達可能なタグに一致するパターンがない場合、ウィジェットは自動的に非表示になります。

ウィジェットの切り替え: ウィジェット ID を無効なウィジェットに追加して、表示モードから非表示にします。
色の凡例: 🟢 0-50% 安全 / 🟡 51-80% 警告 / 🔴 81-100% クリティカル
ステータスラインの表示モード、言語、プランを設定します。 「構成」を参照してください。
すべての AI CLI (Claude、Codex、Gemini、z.ai) の使用制限を一度に確認し、どの CLI が最も利用可能な容量を持つかを推奨します。
/claude-dashboard:check-usage # 色付きのインタラクティブな出力
/claude-dashboard:check-usage --json # スクリプト用の JSON 出力
/claude-dashboard:check-usage --lang ko # 言語を指定します
check-ai シェル エイリアスを追加して、端末からすべての AI CLI の使用状況をすばやく確認します。 macOS/Linux (zsh/bash) および Windows (PowerShell) をサポートします。
/claude-dashboard:setup-alias
セットアップ後:
check-ai # きれいな出力
check-ai --json # スクリプト用の JSON 出力
/クロードダッシュボード:更新
プラグインを更新し、設定の statusLine パスを更新します。 git pull またはマーケットプレイス経由で更新した後に実行します。
/クロードダッシュボード:更新
トラブルシューティング
ステータス行が表示されない
プラグインがインストールされているかどうかを確認します: /plugin list
settings.json に statusLine 構成があることを確認する
API トークンの有効期限が切れている可能性があります - Claude Code に再ログインしてください
ネットワークの問題 - インターネット接続を確認してください
API レート制限 - キャッシュの更新まで 60 秒待ちます
明示的な言語を使用してセットアップを実行します。
/claude-dashboard:通常のセットアップ ko # 韓国語
/claude-dashboard:通常のセットアップ en # 英語
キャッシュの問題
API レスポンス キャッシュは ~/.cache/claude-dashboard/ に保存されます。クリアするには:
rm -rf ~ /.cache/claude-dashboard/
キャッシュ ファイルは 1 時間後に自動的にクリーンアップされます。
npm install && npm run build
詳細については、CONTRIBUTING.md を参照してください。
このプラグインが役立つと思われる場合は、開発をサポートしてください。
コンテキストの使用法、API レート制限、コスト追跡を備えた Claude Code 用の包括的なステータス ライン プラグイン
クロードダッシュボード.uppinote.de

v
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
50
フォーク
レポートリポジトリ
リリース
46
v1.29.0
最新の
2026 年 5 月 30 日
+ 45 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Comprehensive status line plugin for Claude Code with context usage, API rate limits, and cost tracking - uppinote20/claude-dashboard

GitHub - uppinote20/claude-dashboard: Comprehensive status line plugin for Claude Code with context usage, API rate limits, and cost tracking · GitHub
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
uppinote20
/
claude-dashboard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
363 Commits 363 Commits .claude-plugin .claude-plugin .github .github commands commands dist dist images images locales locales scripts scripts website website .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md demo.gif demo.gif package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Comprehensive status line plugin for Claude Code — unified usage monitoring across Claude, Codex, Gemini, and z.ai CLIs, with context, rate limits, cost tracking, and a modular widget system.
Requirements: Claude Code v1.0.80+ / Node.js 18+
/plugin marketplace add uppinote20/claude-dashboard
/plugin install claude-dashboard
/claude-dashboard:setup
Manual Installation
git clone https://github.com/uppinote20/claude-dashboard.git ~ /.claude/plugins/claude-dashboard
/claude-dashboard:setup
Display Modes
Model, context progress bar, cost, rate limits (5h/7d/7d-Sonnet) or z.ai usage (mutually exclusive by provider)
Adds project info, session ID, session duration, burn rate, todo progress
Adds depletion time, config counts, tool/agent status, cache hit, performance badge, token breakdown, forecast, budget, Codex/Gemini usage
Multi-provider support: z.ai/ZHIPU, Codex, Gemini auto-detected when installed.
Category
Widget
Description
Core
model
Model name with emoji, effort level for Opus/Sonnet (X/H/M/L), fast mode for Opus (↯)
context
Progress bar, percentage, tokens (🟢 0-50% / 🟡 51-80% / 🔴 81-100%)
contextBar
Progress bar only (sub-widget of context )
contextPercentage
Percentage only (sub-widget of context )
contextUsage
Token count only, e.g. 42K/200K (sub-widget of context )
cost
Session cost in USD
projectInfo
Directory + git branch (clickable OSC8 link) + ahead/behind (↑↓), subpath from project_dir, worktree indicator (🌳)
Rate Limits
rateLimit5h
5-hour rate limit with reset countdown
rateLimit7d
7-day rate limit (Pro and Max)
rateLimit7dSonnet
7-day Sonnet limit (Max only)
Session
sessionId
Session ID (short 8-char)
sessionIdFull
Session ID (full UUID)
sessionName
Session name from /rename command
sessionDuration
Session duration
lastPrompt
Last user prompt with timestamp
configCounts
CLAUDE.md, AGENTS.md, rules, MCPs, hooks, +Dirs
Activity
toolActivity
Running/completed tools with targets (e.g., Read(app.ts) )
agentStatus
Subagent progress
todoProgress
Todo completion rate
Analytics
burnRate
Token consumption per minute
cacheHit
Cache hit rate percentage
depletionTime
Estimated time to rate limit (approx)¹
Multi-CLI
codexUsage
OpenAI Codex CLI usage (auto-hide if not installed)²
geminiUsage
Google Gemini CLI - current model (auto-hide if not installed)³
geminiUsageAll
Google Gemini CLI - all models (auto-hide if not installed)³
zaiUsage
z.ai/ZHIPU usage (auto-hide if not using z.ai)⁴
Insights
tokenBreakdown
Input/output/cache write/read token breakdown
performance
Composite efficiency badge (cache hit + output ratio)
forecast
Estimated hourly cost based on session rate
budget
Daily spending vs configured budget limit⁵
tokenSpeed
Output token generation speed (tok/s)
todayCost
Total spending across all sessions today
Status
peakHours
Peak hours indicator with countdown ( based on PeakClaude )⁶
tagStatus
Commits ahead of matched git tags (default pattern v* , c
[truncated]
i18n: English and Korean supported (auto-detect or set via setup).
# Preset modes
/claude-dashboard:setup compact # 1 line (default)
/claude-dashboard:setup normal en pro # 2 lines, English, Pro plan
/claude-dashboard:setup detailed ko max # 6 lines, Korean, Max plan
# Custom mode: control widget order and line composition
# Format: "widget1,widget2,...|widget3,widget4,..." (| separates lines)
/claude-dashboard:setup custom auto max " model,context,cost|projectInfo,todoProgress "
Plan differences:
Interactive Mode: Run /claude-dashboard:setup without arguments. Best for preset selection; for full widget control, use Direct Mode or edit the JSON file directly.
Display Mode Presets ( zaiUsage and rateLimit* are mutually exclusive based on provider):
Configuration file ( ~/.claude/claude-dashboard.local.json ):
{
"language" : " auto " ,
"plan" : " max " ,
"displayMode" : " custom " ,
"lines" : [
[ " model " , " context " , " cost " , " rateLimit5h " ],
[ " projectInfo " , " todoProgress " ]
],
"theme" : " default " ,
"separator" : " pipe " ,
"dailyBudget" : 15 ,
"disabledWidgets" : [],
"cache" : { "ttlSeconds" : 60 }
}
Or use preset shorthand for quick configuration:
{
"preset" : " MC$R|BDO " ,
"theme" : " tokyoNight " ,
"separator" : " dot "
}
Themes: default (pastel) / minimal (monochrome) / catppuccin / catppuccinLatte (light-mode) / dracula / gruvbox / nord / tokyoNight / solarized
Separators: pipe (│, default) / space / dot (·) / arrow (›)
Preset Shortcuts: Quick layout with single characters — "preset": "MC$R|BDO" (M=model, C=context, $=cost, R=rateLimit5h, etc.)
Budget Tracking: Set "dailyBudget": 15 to track daily spending. Shows ⚠️ at 80% and 🚨 at 95%.
Tag Status: Customize tagStatus patterns via "tagPatterns": ["v*", "release-*"] . Default is ["v*"] . The widget auto-hides when no pattern matches a reachable tag.
Widget Toggle: Add widget IDs to disabledWidgets to hide them from any display mode.
Color Legend: 🟢 0-50% Safe / 🟡 51-80% Warning / 🔴 81-100% Critical
Configure the status line display mode, language, and plan. See Configuration .
Check usage limits for all AI CLIs (Claude, Codex, Gemini, z.ai) at once and get a recommendation for which CLI has the most available capacity.
/claude-dashboard:check-usage # Interactive output with colors
/claude-dashboard:check-usage --json # JSON output for scripting
/claude-dashboard:check-usage --lang ko # Specify language
Add a check-ai shell alias to quickly check all AI CLI usage from your terminal. Supports macOS/Linux (zsh/bash) and Windows (PowerShell).
/claude-dashboard:setup-alias
After setup:
check-ai # Pretty output
check-ai --json # JSON output for scripting
/claude-dashboard:update
Update the plugin and refresh the statusLine path in settings. Run after updating via git pull or marketplace.
/claude-dashboard:update
Troubleshooting
Status line not showing
Check if plugin is installed: /plugin list
Verify settings.json has statusLine config
API token may be expired - re-login to Claude Code
Network issue - check internet connection
API rate limited - wait 60 seconds for cache refresh
Run setup with explicit language:
/claude-dashboard:setup normal ko # Korean
/claude-dashboard:setup normal en # English
Cache issues
API response cache is stored in ~/.cache/claude-dashboard/ . To clear:
rm -rf ~ /.cache/claude-dashboard/
Cache files are automatically cleaned up after 1 hour.
npm install && npm run build
See CONTRIBUTING.md for details.
If you find this plugin useful, support development:
Comprehensive status line plugin for Claude Code with context usage, API rate limits, and cost tracking
claude-dashboard.uppinote.dev
Resources
There was an error while loading. Please reload this page .
50
forks
Report repository
Releases
46
v1.29.0
Latest
May 30, 2026
+ 45 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
