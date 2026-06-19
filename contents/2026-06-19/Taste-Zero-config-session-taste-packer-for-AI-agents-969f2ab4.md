---
source: "https://github.com/dvcoolarun/taste-ai"
hn_url: "https://news.ycombinator.com/item?id=48597587"
title: "Taste – Zero-config session-taste packer for AI agents"
article_title: "GitHub - dvcoolarun/taste-ai: taste - Zero-config session-taste packer for AI agents · GitHub"
author: "dvcoolarun"
captured_at: "2026-06-19T13:18:53Z"
capture_tool: "hn-digest"
hn_id: 48597587
score: 2
comments: 0
posted_at: "2026-06-19T11:58:57Z"
tags:
  - hacker-news
  - translated
---

# Taste – Zero-config session-taste packer for AI agents

- HN: [48597587](https://news.ycombinator.com/item?id=48597587)
- Source: [github.com](https://github.com/dvcoolarun/taste-ai)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T11:58:57Z

## Translation

タイトル: Taste – AI エージェント向けの設定不要のセッション テイスト パッカー
記事のタイトル: GitHub - dvcoolarun/taste-ai:taste - AI エージェント用のゼロ構成セッション テイスト パッカー · GitHub
説明:taste - AI エージェント用のゼロ構成セッション テイスト パッカー - dvcoolarun/taste-ai

記事本文:
GitHub - dvcoolarun/taste-ai:taste - AI エージェント用のゼロ構成セッション テイスト パッカー · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
dvcoolarun
/
味あい
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット アセット アセット ドキュメント ドキュメント サンプル サンプル .agent-taste.json .agent-taste.json .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh テイスト テイスト すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ゼロ構成。自動学習します。ただ動作します。
97% 小さいコンテキスト · スタイルを自動学習 · あらゆるエージェントと連携
AI エージェント コンテキストを 56K トークンから 1.9K トークンに圧縮します。 git 履歴とセッション ログからコーディング パターンを学習します。結果はプロジェクトの規模とセッション履歴によって異なります。
あなたは問題を知っています。 AI エージェント セッションを開始します。セッション ログ、git diff、README、設定ファイルなど、プロジェクト全体を読み取ります。あなたのスタイルに合わないコードが書かれます。トークンを無駄に消費し、汎用の肥大化したコードを生成します。
味はそれを止めます。それはあなたのパターンを学習します。コンテキストが圧縮されます。エージェントもあなたと同じようにコードを書くことができます。
レートリミッターを要求します。エージェントはコンテキストの 56K トークンを読み取り、ライブラリをインストールし、汎用実装を作成し、Redis セットアップについて質問します。
味
# 1.9K トークンを使用して .session-doc.md を作成します
# エージェントがあなたのパターンを読み取り、あなたの方法でコードを書きます
その他の例は、examples/ にあります。
5 つの指標、1 つの目標: エージェントが少ないコンテキストでより良いコードを作成できるようにします。
コンテキストが 97% 小さくなり、スタイルが自動学習され、あらゆるエージェントと連携します。テイストが学習するすべてのパターンは、コード内で信頼スコアでマークされます。自分で再現してみてください。どのプロジェクトでも実行して味を学びましょう。方法と生の数値: ベンチマーク/ 。実際の例: example/ 。
それは副産物であり、ピッチではありません。これらは平均的な数値であり、プロジェクトによって異なります。より多くのセッション履歴を持つ大規模なプロジェクトでは、圧縮率が向上します。歴史が浅い小規模なプロジェクトでは、節約できる額は小さくなります。そして、これらすべてが反復されます

ve:taste learn を実行するたびに、より多くのパターンを学習し、次の圧縮をより良くします。ルールは決して「トークンの最小数」ではありませんでした。それは、プロジェクトに必要なものだけを学び、検証、エラー処理、セキュリティ、アクセシビリティを決してスキップしないことです。コンテキストはトリミングされるのではなく必要であるため、最終的に小さくなり、それが有用な部分になります。コードの品質が向上するのは自分のスタイルを学ぶことの副作用であり、それが重要な部分です。
コンテキストを圧縮する前に、味覚はパターンを学習します。
1. セッションデータを収集 → git diff、セッションログ、テイスト設定
2. 概要を作成 → 必要な情報のみをまとめたコンパクトなフォーマット
3. エージェントで分析 → opencode または claude を呼び出す
4. パターンの抽出 → 命名、アーキテクチャ、インポート、エラー処理、スタイル
5. テイスト構成を更新 → TASTE.md に追加し、.agent-taste.json を更新します
怠惰ではなく、怠慢ではありません。検証、エラー処理、セキュリティ、およびアクセシビリティのパターンがスキップされることはありません。
これまでで最も努力を要する味は次のとおりです。
1 行インストール (推奨)
カール -fsSL https://raw.githubusercontent.com/dvcoolarun/taste-ai/main/install.sh |バッシュ
手動インストール
git clone https://github.com/dvcoolarun/taste-ai.git
CD テイストアイ
chmod +x テイスト
cp 味 ~ /.local/bin/
インストールの検証
味の助け
それはそれでした。それは誇りに思うでしょう。それは言わないでしょう。
# 任意のプロジェクトディレクトリ内
味
# 圧縮されたコンテキストを使用して .session-doc.md を生成します
パターンを学ぶ
# 過去 3 つのセッションを分析してパターンを学習する
味わう
# 過去 5 つのセッションを分析する
味を学ぶ -- 深さ 5
# ファイルを更新せずにパターンを表示する
味を知る -- 予行練習
テイストコンフィグの初期化
# デフォルトの .agent-taste.json を作成する
味の初期化
現在の構成を表示
# 現在のテイスト設定を表示します
味覚ショー
コマンド
コマンド
何をするのか
味
セッションコンテキストを.session-doc.mdにパックする
テイストパック [

ファイル]
特定の出力ファイルにパックする
味の初期化
デフォルトの .agent-taste.json を作成する
味覚ショー
現在のテイスト構成を表示
味わう
最近のセッションからパターンを学習する (エージェント支援)
味の助け
ヘルプを表示する
パターン学習
Taste Learn はコーディング セッションを分析し、パターンを抽出します。
┌─────────────────────────┐
│ データ収集 │
│ - 最新 3 ～ 5 セッションのログ │
│ - 最新の 3 ～ 5 つのプロンプト ログ │
│ - Git の差分 (最後の 3 ～ 5 コミット) │
│ - 現在のテイスト構成 │
━━━━━━━━━━━━━━━━━━━┘
│
▼
┌─────────────────────────┐
│概要作成│
│ - コンパクトフォーマット (通常 18KB) │
│ - トークン効率の高い構造 │
│ - 必須情報のみ │
━━━━━━━━━━━━━━━━━━━┘
│
▼
┌─────────────────────────┐
│ エージェント分析 │
│ - opencode または claude を呼び出します │
│ - パターン抽出プロンプトを使用する │
│ - 構造化パターンを返します │
━━━━━━━━━━━━━━━━━━━┘

│
▼
┌─────────────────────────┐
│ パターン抽出 │
│ - 命名規則 → TASTE.md │
│ - アーキテクチャパターン → TASTE.md │
│ - インポートスタイル → TASTE.md │
│ - ERROR_HANDLING パターン → TASTE.md │
│ - スタイル設定 → TASTE.md │
│ - 禁止パターン → .agent-taste.json │
━━━━━━━━━━━━━━━━━━━┘
│
▼
┌─────────────────────────┐
│ 自動アップデート │
│ - ポジティブパターン → TASTE.md │
│ - 禁止パターン → .agent-taste.json │
│ - 既存のパターンを保存する │
│ - 重複を避ける │
━━━━━━━━━━━━━━━━━━━┘
パターンのカテゴリ
カテゴリ
学習内容
出力場所
ネーミング
関数の命名規則 (snake_case、camelCase など)
テイスト.md
アーキテクチャ
プロジェクト構造パターン、依存関係管理
テイスト.md
輸入
インポートのスタイル、順序、遅延インポートと即時インポート
テイスト.md
エラー処理
トライ/キャッチパターン、エラー伝播
テイスト.md
スタイル
コードのフォーマット、関数の長さ、コメント
テイスト.md
禁止されたパターン
してはいけないこととその理由
.agent-taste.json
禁止パターン
味覚はポジティブなパターン（何をすべきか）とネガティブなパターン（何をすべきではないか）の両方を学習します。禁止パターンはユーザーの修正から抽出されます。

間違いやフィードバック。
{
"禁止パターン" : [
" --single-process_Chromium_flag_on_macOS (理由: クラッシュの原因、文書化された失敗) " ,
"hardcoding_connection_URLs_or_env_specific_values (理由: 'キュー サービスが利用できない' エラーが発生しました) " ,
" Jump_to_implementation_before_design_alignment (理由: 価格設定モデルが確認されていないときの無駄な作業) " ,
"removing_comments_during_code_rewrites (理由: ユーザーが明示的に呼び出し、保存を期待しています) " ,
" using_browser_only_Node_APIs_in_subprocess (理由: ErrorEvent により ReferenceError が発生しました) "
】
}
禁止されたパターンが重要な理由:
特定 - 汎用 (「クラスを使用しない」) ではなく、具体的 (「--single_process_Chrome_flag を使用しない」)
実用的 - 禁止される理由を説明する明確な理由
間違いから学んだ - 「page.setContent に置き換えられた」は歴史的背景を示しています
プラットフォーム対応 - 「macOS でのクラッシュ」は環境固有の知識を示します
エージェントがセッション データから BANNED_PATTERNS を抽出します
パターンは JSON 配列として .agent-taste.json に書き込まれます。
テイストを実行すると、禁止されたパターンが .session-doc.md に含まれます
エージェントは禁止されたパターンを読み取り、それらのパターンを回避します
各パターンには信頼度スコア (0 ～ 1) が含まれています。
0.9-1.0 : 非常に高い信頼性 (複数回確認)
0.8-0.9 : 高い信頼性 (一貫して見られる)
0.7-0.8 : 中程度の信頼度 (時折見られる)
0.6-0.7 : 信頼度が低い (1 回または 2 回確認された)
<0.6 : 含まれていない (証拠不十分)
味の学習を実行するたびに、次のことが行われます。
最新のセッション ログと git diff を読み取ります
AIエージェントでパターンを分析
信頼スコアを使用して構造化パターンを抽出します
テイスト構成を更新します (TASTE.md、.agent-taste.json)
既存のパターンを保持し、重複を回避します
テイストを使えば使うほど、あなたのスタイルを学習していきます。パターンの信頼性は、

複数のセッションにわたって同じパターン。
# オープンコードを開始する前に
味
# オープンコードにフィードする
オープンコード 。
クロードコード付き
# クロードを始める前に
味
# クロードにフィードする
クロード 。
どのエージェントでも
# コンテキストを生成する
味
# エージェントは .session-doc.md を自動的に読み取ります
出力例
# 味覚の境界線
ソース: ` .agent-taste.json `
「」json
{
"flavor" : " Functional TypeScript、厳密な型、依存関係なし " ,
"banned_patterns" : [ " クラス " 、 " 任意の " 、 " console.log " ]、
"style" : " 暗黙的な戻り値、関数あたり最大 20 行 "
}
最近の仕事
ブランチ: メイン
最後の 5 つのコミット:
abc1234 リファクタリング: 認証を /core に抽出します
def5678 特技: トークン検証の追加
変更されたファイル (最後のコミット):
src/auth.ts | 12 +++---
src/utils.ts | 5 +++-
### 「味を知る」出力
味: 過去 3 回のセッションを分析中...
味: セッションデータを収集中...
テイスト: 作成された概要: 18818 バイト
味: 分析のためにオープンコードを呼び出しています...
学習したパターン (過去 3 セッション):
関数説明アクション動詞 (信頼度: 0.9)
classes_use_PascalCase_Prefixed (信頼度: 0.9)
変数_アンダースコア_分離_スネーク_ケース (信頼度: 0.8)
Python_FastAPI_fronts_with_Node_subprocess_backend_via_stdin_stdout_bridge
async_job_queue_with_redis_backend_and_RQ_worker
Dual_storage_PDF_disk_and_Redis_cache
Lazy_import_inside_endpoint_to_avoid_side_Effects
from_stdlib_then_third_party_then_local_grouped
明示的インポート_非スターインポート_使用
log_then_raise_precise_HTTPException_with_detail
check_rate_limit_before_database_operation
セッション終了前に、値を保存することでクレジットを払い戻す
short_direct_corrections_fix_agent_behavior_正確に
comment_preservation_expected_across_rewrites
--single-process_Chromium_flag_on_macOS (理由: クラッシュの原因、文書化された失敗)
Hardcoding_connection_URLs_or_env_specific_values (理由: 'Queue servi' が発生しました)

ce が利用できませんの失敗)
Jumping_to_implementation_before_design_alignment (理由: 価格モデルが確認されていないときの無駄な作業)
Remove_comments_during_code_rewrites (理由: ユーザーが明示的に呼び出し、保存を期待しています)
using_browser_only_Node_APIs_in_subprocess (理由: ErrorEvent により ReferenceError が発生しました)
更新: TASTE.md、.agent-taste.json
## 自動キャプチャ
最近のセッション ファイルが存在しない場合、taste learn は現在のセッションを自動的にキャプチャします。
「」バッシュ
# docs/session-*.md が存在しないか、2 時間以上古い場合
味わう
# 次のようになります:
#1. git 履歴 + 差分をキャプチャする
# 2. 現在のテイスト構成をキャプチャする
# 3. docs/session-[YYYY-MM-DD-HHMM].md を作成する
# 4. プロンプト/prompt-[YYYY-MM-DD-HHMM].md を作成します。
#5. パターンを分析する

[切り捨てられた]

## Original Extract

taste - Zero-config session-taste packer for AI agents - dvcoolarun/taste-ai

GitHub - dvcoolarun/taste-ai: taste - Zero-config session-taste packer for AI agents · GitHub
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
dvcoolarun
/
taste-ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits assets assets docs docs examples examples .agent-taste.json .agent-taste.json .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh taste taste View all files Repository files navigation
Zero-config. Auto-learns. Just works.
97% smaller context · Auto-learns your style · Works with any agent
Compresses AI agent context from 56K tokens to 1.9K tokens. Learns your coding patterns from git history and session logs. Results vary by project size and session history.
You know the problem. You start an AI agent session. It reads your entire project: session logs, git diffs, READMEs, config files. It writes code that doesn't match your style. It wastes tokens and produces generic, bloated code.
taste puts a stop to that. It learns your patterns. It compresses your context. It makes your agents write code like you do.
You ask for a rate limiter. Your agent reads 56K tokens of context, installs a library, writes a generic implementation, and asks about your Redis setup.
taste
# Creates .session-doc.md with 1.9K tokens
# Agent reads your patterns, writes code your way
More examples in examples/ .
Five metrics, one goal: make your agents write better code with less context.
97% smaller context, auto-learns your style, and works with any agent. Every pattern taste learns is marked in the code with confidence scores. Reproduce it yourself: run taste learn in any project. Method and raw numbers: benchmarks/ . Real-world examples: examples/ .
That is the byproduct, not the pitch. These are average numbers, and they vary by project. Larger projects with more session history see better compression. Smaller projects with less history see smaller savings. And all of this is iterative: each time you run taste learn , it learns more patterns, which makes the next compression better. The rule was never "fewest tokens." It is: learn only what the project needs, and never skip validation, error handling, security, or accessibility. The context ends up small because it is necessary, not trimmed, and that is the part that stays useful. Better code quality is a side effect of learning your style, and that is the part that matters.
Before compressing context, taste learns your patterns:
1. Collect session data → git diffs, session logs, taste config
2. Create summary → compact format, essential information only
3. Analyze with agent → calls opencode or claude
4. Extract patterns → naming, architecture, imports, error handling, style
5. Update taste config → append to TASTE.md, update .agent-taste.json
Lazy, not negligent: validation, error handling, security, and accessibility patterns are never skipped.
The most effort taste will ever ask of you:
One-Line Install (Recommended)
curl -fsSL https://raw.githubusercontent.com/dvcoolarun/taste-ai/main/install.sh | bash
Manual Install
git clone https://github.com/dvcoolarun/taste-ai.git
cd taste-ai
chmod +x taste
cp taste ~ /.local/bin/
Verify Installation
taste help
That was it. It would be proud. It won't say it.
# In any project directory
taste
# Generates .session-doc.md with compressed context
Learn Patterns
# Analyze last 3 sessions and learn patterns
taste learn
# Analyze last 5 sessions
taste learn --depth 5
# Show patterns without updating files
taste learn --dry-run
Initialize Taste Config
# Create default .agent-taste.json
taste init
Show Current Config
# Display current taste configuration
taste show
Commands
Command
What it does
taste
Pack session context into .session-doc.md
taste pack [file]
Pack to specific output file
taste init
Create default .agent-taste.json
taste show
Show current taste config
taste learn
Learn patterns from recent sessions (agent-assisted)
taste help
Show help
Pattern Learning
taste learn analyzes your coding sessions and extracts patterns:
┌─────────────────────────────────────────────────────────┐
│ Data Collection │
│ - Last 3-5 session logs │
│ - Last 3-5 prompt logs │
│ - Git diffs (last 3-5 commits) │
│ - Current taste config │
└─────────────────────────────────────────────────────────┘
│
▼
┌─────────────────────────────────────────────────────────┐
│ Summary Creation │
│ - Compact format (18KB typical) │
│ - Token-efficient structure │
│ - Essential information only │
└─────────────────────────────────────────────────────────┘
│
▼
┌─────────────────────────────────────────────────────────┐
│ Agent Analysis │
│ - Calls opencode or claude │
│ - Uses pattern extraction prompt │
│ - Returns structured patterns │
└─────────────────────────────────────────────────────────┘
│
▼
┌─────────────────────────────────────────────────────────┐
│ Pattern Extraction │
│ - NAMING conventions → TASTE.md │
│ - ARCHITECTURE patterns → TASTE.md │
│ - IMPORTS style → TASTE.md │
│ - ERROR_HANDLING patterns → TASTE.md │
│ - STYLE preferences → TASTE.md │
│ - BANNED_PATTERNS → .agent-taste.json │
└─────────────────────────────────────────────────────────┘
│
▼
┌─────────────────────────────────────────────────────────┐
│ Auto-Update │
│ - Positive patterns → TASTE.md │
│ - Banned patterns → .agent-taste.json │
│ - Preserve existing patterns │
│ - Avoid duplicates │
└─────────────────────────────────────────────────────────┘
Pattern Categories
Category
What It Learns
Output Location
NAMING
Function naming conventions (snake_case, camelCase, etc.)
TASTE.md
ARCHITECTURE
Project structure patterns, dependency management
TASTE.md
IMPORTS
Import style, ordering, lazy vs eager imports
TASTE.md
ERROR_HANDLING
Try/catch patterns, error propagation
TASTE.md
STYLE
Code formatting, function length, comments
TASTE.md
BANNED_PATTERNS
What NOT to do, with reasons
.agent-taste.json
Banned Patterns
taste learns both positive patterns (what to do) and negative patterns (what NOT to do). Banned patterns are extracted from user corrections, past mistakes, and feedback.
{
"banned_patterns" : [
" --single-process_Chromium_flag_on_macOS (reason: causes crashes, documented failure) " ,
" hardcoding_connection_URLs_or_env_specific_values (reason: caused 'Queue service unavailable' failure) " ,
" jumping_to_implementation_before_design_alignment (reason: wasted work when pricing model wasn't confirmed) " ,
" removing_comments_during_code_rewrites (reason: user explicitly called out and expects preservation) " ,
" using_browser_only_Node_APIs_in_subprocess (reason: ErrorEvent caused ReferenceError) "
]
}
Why banned patterns matter:
Specific - Not generic ("don't use classes") but concrete ("don't use --single_process_Chrome_flag")
Actionable - Clear reasons that explain WHY it's banned
Learned from mistakes - "was replaced with page.setContent" shows historical context
Platform-aware - "crashes on macOS" shows environment-specific knowledge
Agent extracts BANNED_PATTERNS from session data
Patterns are written to .agent-taste.json as a JSON array
When you run taste , banned patterns are included in .session-doc.md
Agents read the banned patterns and avoid those patterns
Each pattern includes a confidence score (0-1):
0.9-1.0 : Very high confidence (seen multiple times)
0.8-0.9 : High confidence (seen consistently)
0.7-0.8 : Medium confidence (seen occasionally)
0.6-0.7 : Low confidence (seen once or twice)
<0.6 : Not included (insufficient evidence)
Each time you run taste learn , it:
Reads your latest session logs and git diffs
Analyzes patterns with an AI agent
Extracts structured patterns with confidence scores
Updates your taste config (TASTE.md, .agent-taste.json)
Preserves existing patterns and avoids duplicates
The more you use taste, the better it learns your style. Pattern confidence increases as it sees the same patterns across multiple sessions.
# Before starting opencode
taste
# Feed to opencode
opencode .
With Claude Code
# Before starting claude
taste
# Feed to claude
claude .
With Any Agent
# Generate context
taste
# Agent reads .session-doc.md automatically
Example Output
# TASTE BOUNDARIES
Source: ` .agent-taste.json `
``` json
{
"flavor" : " Functional TypeScript, strict types, zero dependencies " ,
"banned_patterns" : [ " classes " , " any " , " console.log " ],
"style" : " Implicit returns, max 20 lines per function "
}
RECENT WORK
Branch: main
Last 5 commits:
abc1234 refactor: extract auth to /core
def5678 feat: add token validation
Changed files (last commit):
src/auth.ts | 12 +++---
src/utils.ts | 5 +++-
### `taste learn` Output
taste: Analyzing last 3 sessions...
taste: Collecting session data...
taste: Summary created: 18818 bytes
taste: Calling opencode for analysis...
LEARNED PATTERNS (last 3 sessions):
functions_describe_action_verbs (confidence: 0.9)
classes_use_PascalCase_Prefixed (confidence: 0.9)
variables_underscore_separated_snake_case (confidence: 0.8)
Python_FastAPI_fronts_with_Node_subprocess_backend_via_stdin_stdout_bridge
async_job_queue_with_redis_backend_and_RQ_worker
dual_storage_PDF_disk_and_Redis_cache
lazy_import_inside_endpoint_to_avoid_side_effects
from_stdlib_then_third_party_then_local_grouped
explicit_imports_not_star_imports_used
log_then_raise_precise_HTTPException_with_detail
check_rate_limit_before_database_operation
refund_credits_by_saving_values_before_session_closes
short_direct_corrections_fix_agent_behavior_precisely
comment_preservation_expected_across_rewrites
--single-process_Chromium_flag_on_macOS (reason: causes crashes, documented failure)
hardcoding_connection_URLs_or_env_specific_values (reason: caused 'Queue service unavailable' failure)
jumping_to_implementation_before_design_alignment (reason: wasted work when pricing model wasn't confirmed)
removing_comments_during_code_rewrites (reason: user explicitly called out and expects preservation)
using_browser_only_Node_APIs_in_subprocess (reason: ErrorEvent caused ReferenceError)
Updated: TASTE.md, .agent-taste.json
## Auto-Capture
taste learn automatically captures your current session if no recent session files exist:
```bash
# If docs/session-*.md doesn't exist or is older than 2 hours
taste learn
# It will:
# 1. Capture git history + diffs
# 2. Capture current taste config
# 3. Create docs/session-[YYYY-MM-DD-HHMM].md
# 4. Create prompts/prompt-[YYYY-MM-DD-HHMM].md
# 5. Analyze patterns

[truncated]
