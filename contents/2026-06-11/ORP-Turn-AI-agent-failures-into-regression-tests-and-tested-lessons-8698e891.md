---
source: "https://github.com/Fujo930/ORP"
hn_url: "https://news.ycombinator.com/item?id=48496737"
title: "ORP – Turn AI agent failures into regression tests and tested lessons"
article_title: "GitHub - Fujo930/ORP: Open Reflection Protocol — Turn agent failures into regression tests, reusable lessons, and measurable improvements. Built on OpenTelemetry. · GitHub"
author: "Fujo930"
captured_at: "2026-06-11T21:44:51Z"
capture_tool: "hn-digest"
hn_id: 48496737
score: 1
comments: 0
posted_at: "2026-06-11T21:33:38Z"
tags:
  - hacker-news
  - translated
---

# ORP – Turn AI agent failures into regression tests and tested lessons

- HN: [48496737](https://news.ycombinator.com/item?id=48496737)
- Source: [github.com](https://github.com/Fujo930/ORP)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T21:33:38Z

## Translation

タイトル: ORP – AI エージェントの失敗を回帰テストとテスト済みのレッスンに変える
記事のタイトル: GitHub - Fujo930/ORP: Open Reflection Protocol — エージェントの失敗を回帰テスト、再利用可能なレッスン、測定可能な改善に変えます。 OpenTelemetry に基づいて構築されています。 · GitHub
説明: オープン リフレクション プロトコル — エージェントの失敗を回帰テスト、再利用可能なレッスン、測定可能な改善に変えます。 OpenTelemetry に基づいて構築されています。 - ふじょ930/ORP

記事本文:
GitHub - Fujo930/ORP: Open Reflection Protocol — エージェントの失敗を回帰テスト、再利用可能なレッスン、測定可能な改善に変えます。 OpenTelemetry に基づいて構築されています。 · GitHub
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
ふじょ930
/
ORP
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット ブログ ブログ デモ デモの例 例 exps exps 統合 統合 orp orp テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md BLOG_OUTLINE.md BLOG_OUTLINE.md EXPERIMENTS.md EXPERIMENTS.md QUICK_START.md QUICK_START.md README.md README.md ROADMAP.md ROADMAP.md SPEC.md SPEC.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープン リフレクション プロトコル (ORP)
エージェントの失敗を回帰テスト、再利用可能なレッスン、測定可能な改善に変えます。
トレースすると、エージェントが何をしたかが分かります。 ORP は、起こったことをテストされた教訓に変えます。
コーディング エージェントは認証バグを修正しましたが、匿名ユーザー パスを見逃しました。テストは 34/35 点で不合格になります。
# 1. エージェントを ORP でラップする
orp ラップ -- python my_agent.py
# 2. ORP は失敗を捉え、証明されていない主張に異議を唱え、
# レッスン + 回帰 Eval をコンパイルします
orp 最新の情報を学ぶ
# 3. 同じエージェントが MCP 経由でレッスンを取得し、それを適用します
# -> 今回は 35 個のテストすべてに合格しました
orp mcp-server
#4. 前後の比較
orp diff exp_before exp_after
前:
タスクの成功: 失敗 (34/35 テスト)
主張: 1 証明されていない
後:
タスクの成功: 合格 (35/35 テスト)
主張: 0 証明されていない
それがループです。 1 つの間違い、1 つの教訓、1 つの目に見える改善。
ORP は、OpenTelemetry 上に構築された AI エージェント用のオープン エクスペリエンス レイヤーです。エージェント トレースを 3 つの実行可能なアーティファクトに変換します。
各レッスンはライフサイクルを経ます。
候補者 -> アクティブ -> 審査中 -> 非推奨 -> 拒否
|
（アクティブレッスンのみ）
回収可能です）
主要な概念
証拠優先 : ORP は、観察された事実 (ツールの出力、テスト結果) とエージェントの主張 (診断、信頼性ステータス) を区別します。

発言）。クレームが自動的に真実として扱われることはありません。
実行可能エクスペリエンス: レッスンはテキストだけでなく、実行可能な eval とガードレールにコンパイルされます。
成果価値：レッスンの質は、効果評価によって実際に成果が上がったかどうかで決まります。
OpenTelemetry に基づいて構築: ORP は、既存のトレース インフラストラクチャを置き換えるのではなく、拡張します。
デフォルトのプライベート : すべてのデータはローカルに残り、デフォルトでは匿名化され、プロンプト/ツール出力はアップロードされません。
pip インストール オープンリフレクション プロトコル
Python 3.10以降が必要です。
orp ラップ -- python my_agent.py --run-task
ORP は、標準出力、終了コード、テスト結果、git diff、および OpenTelemetry スパンを自動的にキャプチャします。
orp 最新の情報を学ぶ
これにより、以下が生成されます。
何が問題だったかの診断
異議申し立て（サポートされていないエージェントの声明）
orp 最新の検査
orp レポート --open # HTML レポート
orp diff exp_before exp_after
4. 教訓を今後の実行に活かす
# MCP レッスンサーバーを起動します
orp mcp-server --transport stdio
# 互換性のあるエージェントが次の MCP ツールを使用できるようになりました。
# orp_retrieve_lessons(タスク、制限=3)
# orp_acknowledge_lesson(lesson_id)
# orp_report_outcome(lesson_id、outcome、evidence_refs)
デモを実行する
git clone https://github.com/Fujo930/ORP
cd ORP
uv 実行 python デモ/orp_demo.py
出力:
実行 1: エージェントが匿名ユーザー パスを見逃す -> 失敗
ORP が失敗を分析 -> 1 つの証明されていない主張に異議を唱える
ORP はレッスン + 評価をコンパイルします
MCP がエージェントにレッスンを提供します
実行 2: エージェントがレッスンを適用 -> 合格
変更前: 34/35 回のテスト、1 件の未証明の主張
後: 35/35 テスト、未証明の主張は 0
推定効果: 0.5
実験結果
10 個の失敗タスク、それぞれ 5 回の試行、合計 100 回の実行。
ゴー/ノーゴー: >>> ゴー — 4/4 チェックに合格
自分自身を実行します: uv run python exps/runner.py
orp Wrap -- python Agent.py ORP を使用してエージェント プロセスをラップします。
orp Inspection [id] エクスペリエンスを検査します (デフォルト: l

テスト)
orp learn [id] 経験から教訓を生成する
orp リプレイ <id> 反事実的なリプレイ
orp レッスン一覧 レッスン一覧
orp レッスン validate <id> レッスンの整合性を検証します
orp レッスンの競合 競合するレッスンの自動検出
orp Less rollback <id> レッスンをロールバックします
orp レッスンが <id> を配信します レッスンを配信します
orp 効果を評価します <id> レッスン効果を評価します
orp 訓練候補者 訓練候補者一覧
orp トレーニングのエクスポート 承認されたトレーニング データをエクスポートする
orp mcp-server MCP レッスン サーバーを開始します
orp report --open HTML レポートの生成
orp diff <id1> <id2> 2 つのエクスペリエンスを比較する
orp import [id] JSON としてエクスポート
建築
エージェント/既存のトレース
|
v
トレース アダプター (OTel / OpenAI / LangGraph / 汎用 JSON)
|
v
エクスペリエンスビルダー -> 証拠検証者
-> リフレクションアナライザー (診断 + チャレンジャー)
-> 反事実リプレイヤー
|
v
エクスペリエンスコンパイラ
+----+----+------+
| | |
レッスン評価ガードレール
| | |
+---- 配信ルーター (MCP サーバー / プロンプト / ポリシー / ランタイムフック)
|
v
エフェクトエバリュエーター + ロールバック
寄稿者向け
uv 実行 pytest -q
#58 0.68秒で合格
このリポジトリの主要な設計ドキュメント:
オープン リフレクション プロトコル — エージェントの失敗を回帰テスト、再利用可能なレッスン、測定可能な改善に変えます。 OpenTelemetry に基づいて構築されています。
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

Open Reflection Protocol — Turn agent failures into regression tests, reusable lessons, and measurable improvements. Built on OpenTelemetry. - Fujo930/ORP

GitHub - Fujo930/ORP: Open Reflection Protocol — Turn agent failures into regression tests, reusable lessons, and measurable improvements. Built on OpenTelemetry. · GitHub
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
Fujo930
/
ORP
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits blog blog demo demo examples examples exps exps integrations integrations orp orp tests tests .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md BLOG_OUTLINE.md BLOG_OUTLINE.md EXPERIMENTS.md EXPERIMENTS.md QUICK_START.md QUICK_START.md README.md README.md ROADMAP.md ROADMAP.md SPEC.md SPEC.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Open Reflection Protocol (ORP)
Turn agent failures into regression tests, reusable lessons, and measurable improvements.
Tracing tells you what your agent did. ORP turns what happened into a tested lesson.
A coding agent fixes an auth bug but misses the anonymous user path. Tests fail at 34/35.
# 1. Wrap your agent with ORP
orp wrap -- python my_agent.py
# 2. ORP captures the failure, challenges unproven claims,
# and compiles a Lesson + regression Eval
orp learn latest
# 3. Same agent retrieves the Lesson via MCP, applies it
# -> All 35 tests pass this time
orp mcp-server
# 4. Before/after comparison
orp diff exp_before exp_after
Before:
Task success: FAILED (34/35 tests)
Claims: 1 unproven
After:
Task success: PASSED (35/35 tests)
Claims: 0 unproven
That's the loop. One mistake, one lesson, one measurable improvement.
ORP is an open experience layer for AI agents , built on OpenTelemetry. It converts agent traces into three executable artifacts:
Each Lesson goes through a lifecycle:
candidate -> active -> under_review -> deprecated -> rejected
|
(only active lessons
are retrievable)
Key Concepts
Evidence-first : ORP distinguishes observed facts (tool output, test results) from agent claims (diagnoses, confidence statements). Claims are never automatically treated as ground truth.
Executable experience : Lessons compile to runnable evals and guardrails, not just text.
Outcome-based value : Lesson quality is determined by whether it actually improves results, measured through effect evaluation.
Built on OpenTelemetry : ORP extends existing trace infrastructure instead of replacing it.
Default private : All data stays local, de-identified by default, no prompt/tool output uploaded.
pip install open-reflection-protocol
Requires Python 3.10+.
orp wrap -- python my_agent.py --run-task
ORP automatically captures stdout, exit codes, test results, git diff, and OpenTelemetry spans.
orp learn latest
This generates:
A diagnosis of what went wrong
Challenged claims (unsupported agent statements)
orp inspect latest
orp report --open # HTML report
orp diff exp_before exp_after
4. Deliver lessons to future runs
# Start the MCP Lesson server
orp mcp-server --transport stdio
# Compatible agents can now use these MCP tools:
# orp_retrieve_lessons(task, limit=3)
# orp_acknowledge_lesson(lesson_id)
# orp_report_outcome(lesson_id, outcome, evidence_refs)
Run the Demo
git clone https://github.com/Fujo930/ORP
cd ORP
uv run python demo/orp_demo.py
Output:
Run 1: Agent misses anonymous user path -> FAILED
ORP analyzes the failure -> challenges 1 unproven claim
ORP compiles Lesson + Eval
MCP delivers Lesson to Agent
Run 2: Agent applies Lesson -> PASSED
Before: 34/35 tests, 1 unproven claim
After: 35/35 tests, 0 unproven claims
Estimated effect: 0.5
Experimental Results
10 failure tasks, 5 trials each, 100 total runs.
Go/No-Go: >>> GO — 4/4 checks passed
Run yourself: uv run python exps/runner.py
orp wrap -- python agent.py Wrap an agent process with ORP
orp inspect [id] Inspect an experience (default: latest)
orp learn [id] Generate lessons from an experience
orp replay <id> Counterfactual replay
orp lessons list List lessons
orp lessons validate <id> Validate lesson integrity
orp lessons conflicts Auto-detect conflicting lessons
orp lessons rollback <id> Rollback a lesson
orp lessons deliver <id> Deliver a lesson
orp effects evaluate <id> Evaluate lesson effect
orp training candidates List training candidates
orp training export Export approved training data
orp mcp-server Start MCP lesson server
orp report --open Generate HTML report
orp diff <id1> <id2> Compare two experiences
orp export [id] Export as JSON
Architecture
Agent / Existing Trace
|
v
Trace Adapters (OTel / OpenAI / LangGraph / Generic JSON)
|
v
Experience Builder -> Evidence Verifier
-> Reflection Analyzer (diagnosis + challenger)
-> Counterfactual Replayer
|
v
Experience Compiler
+----+----+------+
| | |
Lesson Eval Guardrail
| | |
+---- Delivery Router (MCP Server / Prompt / Policy / Runtime Hook)
|
v
Effect Evaluator + Rollback
For Contributors
uv run pytest -q
# 58 passed in 0.68s
Key design documents in this repo:
Open Reflection Protocol — Turn agent failures into regression tests, reusable lessons, and measurable improvements. Built on OpenTelemetry.
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
