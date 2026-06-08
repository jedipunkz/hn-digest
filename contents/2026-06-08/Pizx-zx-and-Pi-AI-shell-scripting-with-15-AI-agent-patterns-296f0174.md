---
source: "https://github.com/topce/pizx"
hn_url: "https://news.ycombinator.com/item?id=48443198"
title: "Pizx – zx and Pi AI = shell scripting with 15 AI agent patterns"
article_title: "GitHub - topce/pizx · GitHub"
author: "topce"
captured_at: "2026-06-08T10:41:11Z"
capture_tool: "hn-digest"
hn_id: 48443198
score: 1
comments: 0
posted_at: "2026-06-08T09:43:58Z"
tags:
  - hacker-news
  - translated
---

# Pizx – zx and Pi AI = shell scripting with 15 AI agent patterns

- HN: [48443198](https://news.ycombinator.com/item?id=48443198)
- Source: [github.com](https://github.com/topce/pizx)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T09:43:58Z

## Translation

タイトル: Pizx – zx および Pi AI = 15 の AI エージェント パターンを使用したシェル スクリプト
記事タイトル: GitHub - topce/pizx · GitHub
説明: GitHub でアカウントを作成して、topce/pizx の開発に貢献します。

記事本文:
GitHub - topce/pizx · GitHub
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
トプス
/
ピズ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3

コミット ドキュメント ドキュメントの例 サンプル スクリプト スクリプト src src .gitignore .gitignore ライセンス ライセンス README.md README.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ネイティブ Pi AI 統合を備えた zx フォーク — シェル スクリプト、AI テキスト生成、コーディング エージェント、エージェント パターン、通信、およびオーケストレーション トポロジ用の 15 個のテンプレート タグ。
npm インストール @topce/pizx
pi auth login # one-time: Pi AI 資格情報を構成する
スクリプト ( hello.mjs ) を作成します。
#!/usr/bin/env pizx
// 単純な AI クエリ
const 答え = await π 「フランスの首都はどこですか?」
エコー（応答）
// エージェントのパターン
const files = await $ `ls src/`
const summary = await π `これらのファイルを一文に要約します: ${ files } `
コンソール。ログ（概要）
実行してください:
chmod +x hello.mjs
./hello.mjs
# または:
pizx hello.mjs
インストール
npm インストール @topce/pizx
前提条件:
Pi AI のインストールと構成 (pi 認証ログイン)
zx からのシェル コマンド ( $ 、 cd 、 echo 、 fetch など)
#!/usr/bin/env pizx
const name = 質問を待ちます (「あなたの名前は何ですか?」)
const intro = await π ` ${ name } ` にフレンドリーな挨拶を書きます
エコー（イントロ）
プログラムによるインポート
import { $ , π , Π , Ρ , Φ , Σ } から '@topce/pizx'
const 出力 = await $ `ls src/ | grep '.ts'`
コンソール。ログ (出力.stdout)
const review = await π `このコードに問題がないか確認してください:\n ${ 出力 。標準出力 } `
コンソール。ログ（レビューのテキスト）
// コーディング エージェントを使用して問題を修正する
await Π `src/ の TypeScript エラーを修正してください`
CLI クイック クエリ
pizx -p " JavaScript での async/await について説明します"
pizx -p --model deepseek/deepseek-chat " このコードを要約します: @file.ts "
pizx --version
タグリファレンス
各タグにはデータがあります

docs/ 内の iled ドキュメント:
タグ
名前
説明
ドキュメント
$
シェル
シェルコマンド（zxから変更なし）
—
π
円周率
pi-aiによるAIテキスト生成
docs/pi.md
Π
大文字ピ
ツールを備えた Pi コーディング エージェント (読み取り、bash、編集、書き込み)
docs/capital-pi.md
エージェント パターン (Ρ Φ Σ Δ Λ Ψ Ω Ν)
タグ
名前
流れ
ドキュメント
Ρ
ラルフ・ループ
分析 → 計画 → 実行 → レビュー ↺
docs/ralph.md
Φ
艦隊
A、B、C 並列 → 集合
docs/フリート.md
Σ
サブエージェント
分解 → サブエージェント → 合成
docs/subagent.md
Δ
ディベート
視点→収束する
docs/debate.md
Λ
パイプライン
ステージ₁ → ステージ₂ → ステージ₃
docs/pipeline.md
Ψ
批評
生成→批評→改善
docs/critique.md
Ω
オーケストレーター
企画→発送→合成
docs/orchestrator.md
Ν
ヌ
分析 → 役割の交渉 → 実行 → 合成
docs/nu.md
コミュニケーションパターン (Θ Μ Β)
タグ
名前
パターン
ドキュメント
Θ
スレッド
マルチエージェントの会話
docs/thread.md
Μ
記憶
共有黒板
docs/memory.md
Β
ブロードキャスト
1対多のメッセージング
docs/broadcast.md
オーケストレーション トポロジ (Α Γ Χ Τ)
タグ
名前
パターン
ドキュメント
Α
アダプティブ
自動調整ワークフロー
docs/adaptive.md
Γ
グラフ
DAGベースの実行
docs/graph.md
Χ
チー
トレースを解析→パターンを抽出
docs/chi.md
Τ
タウ
スキーマの定義 → 書き込み → 洗練 → 統合
docs/tau.md
高度な機能
すべてのパターンは、高レベルの推論と実行をさまざまなモデルにルーティングするための plannerModel と WorkerModel をサポートしています。
待ってくださいΩ ( {
plannerModel : 'deepseek/deepseek-v4-pro' , // 計画 + 合成
workerModel : 'deepseek/deepseek-v4-flash' , // ワーカーの実行
} ) `通知システムを設計する`
フェーズごとのモデルがない場合、パターンはモデル → Pi のデフォルトに戻ります。
すべてのタグは、出力を抑制するオプション チェーンと .quit モードをサポートしています。
await π ( { モデル : 'anthropic/claude-sonnet-4

-5' } ) `このアルゴリズムを説明してください`
お待ちくださいΠ。 Quiet `src/ の lint の問題を修正してください`
await Φ ( { concurrency : 5 } ) `すべての .ts ファイルを確認する`
Σを待ってください。静かに「コードベース全体のセキュリティを分析する」
await Θ ( { エージェント : 4 , ターン : 3 } ) `アーキテクチャについて議論する`
await Γ ( { グラフ : { ノード : [ ... ] , エッジ : [ ... ] } } ) `ワークフローを実行`
グローバル構成
import { configurePi , configureAgent } から '@topce/pizx'
configurePi ( { モデル : 'anthropic/claude-sonnet-4-5' , maxTokens : 8000 } )
configureAgent ( { maxTurns : 5 , excludeTools : [ 'write' ] } )
CLI リファレンス
pizx [オプション] < script > # pizx スクリプトを実行します
pizx -p <プロンプト> # クイック pi-ai クエリ
pizx --version # バージョンを印刷します
pizx --help # ヘルプを印刷する
オプション:
-p, --prompt <text> — 簡単な pi-ai クエリを実行します (スクリプトは必要ありません)。
-m, --model <id> — 使用する AI モデルを指定します
--quiet — エラー以外の出力を抑制します
--shell <path> — 使用するシェル (デフォルト: 自動検出)
npm run build # ビルド (JS + DTS)
npm run check # Lint と Biome でフォーマットする
npm テスト #95 単体テスト
npm run example:hello # hello サンプルを実行する
npm run example:π # pi-ai の例を実行する
npm run example:all # すべての例を実行します
例
すべてのパターンの実行可能な例については、examples/ を参照してください。
hello-pizx.mjs — シェル + AI を使用した基本スクリプト
Basic-pi.mjs — π テキストの生成
Basic-capital-pi.mjs — Π コーディング エージェント
pattern-ralph.mjs — ラルフ ループ
pattern-fleet.mjs — フリートの並列実行
pattern-debate.mjs — 複数の視点からの議論
...パターンごとにさらに詳しく
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

Contribute to topce/pizx development by creating an account on GitHub.

GitHub - topce/pizx · GitHub
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
topce
/
pizx
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits docs docs examples examples scripts scripts src src .gitignore .gitignore LICENSE LICENSE README.md README.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
zx fork with native Pi AI integration — 15 template tags for shell scripting, AI text generation, coding agents, agentic patterns, communication, and orchestration topologies.
npm install @topce/pizx
pi auth login # one-time: configure Pi AI credentials
Write a script ( hello.mjs ):
#!/usr/bin/env pizx
// Simple AI query
const answer = await π `what is the capital of France?`
echo ( answer )
// Agent patterns
const files = await $ `ls src/`
const summary = await π `summarize these files in one sentence: ${ files } `
console . log ( summary )
Run it:
chmod +x hello.mjs
./hello.mjs
# Or:
pizx hello.mjs
Install
npm install @topce/pizx
Prerequisites:
Pi AI installed and configured ( pi auth login )
Shell commands from zx ( $ , cd , echo , fetch , etc.)
#!/usr/bin/env pizx
const name = await question ( 'What is your name? ' )
const intro = await π `write a friendly greeting for ${ name } `
echo ( intro )
Programmatic Import
import { $ , π , Π , Ρ , Φ , Σ } from '@topce/pizx'
const output = await $ `ls src/ | grep '.ts'`
console . log ( output . stdout )
const review = await π `review this code for issues:\n ${ output . stdout } `
console . log ( review . text )
// Use the coding agent to fix issues
await Π `fix the TypeScript errors in src/`
CLI Quick Queries
pizx -p " explain async/await in JavaScript "
pizx -p --model deepseek/deepseek-chat " summarize this code: @file.ts "
pizx --version
Tags Reference
Each tag has detailed documentation in docs/ :
Tag
Name
Description
Docs
$
Shell
Shell commands (unchanged from zx)
—
π
Pi
AI text generation via pi-ai
docs/pi.md
Π
Capital Pi
Pi coding agent with tools (read, bash, edit, write)
docs/capital-pi.md
Agent Patterns (Ρ Φ Σ Δ Λ Ψ Ω Ν)
Tag
Name
Flow
Docs
Ρ
Ralph Loop
analyze → plan → execute → review ↺
docs/ralph.md
Φ
Fleet
A, B, C in parallel → aggregate
docs/fleet.md
Σ
Subagents
decompose → sub-agents → synthesize
docs/subagent.md
Δ
Debate
perspectives → converge
docs/debate.md
Λ
Pipeline
stage₁ → stage₂ → stage₃
docs/pipeline.md
Ψ
Critique
generate → critique → improve
docs/critique.md
Ω
Orchestrator
plan → dispatch → synthesize
docs/orchestrator.md
Ν
Nu
analyze → negotiate roles → execute → synthesize
docs/nu.md
Communication Patterns (Θ Μ Β)
Tag
Name
Pattern
Docs
Θ
Thread
Multi-agent conversation
docs/thread.md
Μ
Memory
Shared blackboard
docs/memory.md
Β
Broadcast
One-to-many messaging
docs/broadcast.md
Orchestration Topologies (Α Γ Χ Τ)
Tag
Name
Pattern
Docs
Α
Adaptive
Self-adjusting workflow
docs/adaptive.md
Γ
Graph
DAG-based execution
docs/graph.md
Χ
Chi
Analyze traces → extract patterns
docs/chi.md
Τ
Tau
Define schema → write → refine → consolidate
docs/tau.md
Advanced Features
All patterns support plannerModel and workerModel for routing high-level reasoning vs execution to different models:
await Ω ( {
plannerModel : 'deepseek/deepseek-v4-pro' , // planning + synthesis
workerModel : 'deepseek/deepseek-v4-flash' , // worker execution
} ) `design a notification system`
Without per-phase models, patterns fall back to model → Pi default.
All tags support option chaining and .quiet mode to suppress output:
await π ( { model : 'anthropic/claude-sonnet-4-5' } ) `explain this algorithm`
await Π . quiet `fix the lint issues in src/`
await Φ ( { concurrency : 5 } ) `review all .ts files`
await Σ . quiet `analyze security across the codebase`
await Θ ( { agents : 4 , turns : 3 } ) `debate the architecture`
await Γ ( { graph : { nodes : [ ... ] , edges : [ ... ] } } ) `execute workflow`
Global Configuration
import { configurePi , configureAgent } from '@topce/pizx'
configurePi ( { model : 'anthropic/claude-sonnet-4-5' , maxTokens : 8000 } )
configureAgent ( { maxTurns : 5 , excludeTools : [ 'write' ] } )
CLI Reference
pizx [options] < script > # Run a pizx script
pizx -p < prompt > # Quick pi-ai query
pizx --version # Print version
pizx --help # Print help
Options:
-p, --prompt <text> — Run a quick pi-ai query (no script needed)
-m, --model <id> — Specify AI model to use
--quiet — Suppress output except errors
--shell <path> — Shell to use (default: auto-detect)
npm run build # Build (JS + DTS)
npm run check # Lint and format with Biome
npm test # 95 unit tests
npm run example:hello # Run hello example
npm run example:π # Run pi-ai example
npm run example:all # Run all examples
Examples
See examples/ for runnable examples of every pattern:
hello-pizx.mjs — Basic script with shell + AI
basic-pi.mjs — π text generation
basic-capital-pi.mjs — Π coding agent
pattern-ralph.mjs — Ralph Loop
pattern-fleet.mjs — Fleet parallel execution
pattern-debate.mjs — Multi-perspective debate
... and more for every pattern
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
