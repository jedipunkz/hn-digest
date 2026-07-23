---
source: "https://github.com/GiovanniGatti/cve-bench"
hn_url: "https://news.ycombinator.com/item?id=49021912"
title: "Ask HN: Which CVEs should I add to my Python security benchmark for AI agents?"
article_title: "GitHub - GiovanniGatti/cve-bench: A benchmark for evaluating AI agents on fixing real-world security vulnerabilities. · GitHub"
author: "ggattip"
captured_at: "2026-07-23T15:03:44Z"
capture_tool: "hn-digest"
hn_id: 49021912
score: 1
comments: 1
posted_at: "2026-07-23T14:10:28Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Which CVEs should I add to my Python security benchmark for AI agents?

- HN: [49021912](https://news.ycombinator.com/item?id=49021912)
- Source: [github.com](https://github.com/GiovanniGatti/cve-bench)
- Score: 1
- Comments: 1
- Posted: 2026-07-23T14:10:28Z

## Translation

タイトル: HN に聞く: AI エージェントの Python セキュリティ ベンチマークにどの CVE を追加する必要がありますか?
記事のタイトル: GitHub - GiovanniGatti/cve-bench: 現実世界のセキュリティ脆弱性の修正に関して AI エージェントを評価するためのベンチマーク。 · GitHub
説明: 現実世界のセキュリティ脆弱性の修正に関して AI エージェントを評価するためのベンチマーク。 - GiovanniGatti/cve-bench

記事本文:
GitHub - GiovanniGatti/cve-bench: 現実世界のセキュリティ脆弱性の修正に関して AI エージェントを評価するためのベンチマーク。 · GitHub
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
ジョバンニ・ガッティ
/
CVEベンチ
公共
通知
Cha にサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット docker docker docs docs results results src/ harness src/ harness Tasks タスク template/ CVE-2026-XXXX template/ CVE-2026-XXXX .coverage .coverage .dockerignore .dockerignore .gitignore .gitignore CITATION.cff CITATION.cff ライセンス ライセンスREADME.md README.md Benchmark.py Benchmark.py build.py build.py Generate_charts.py Generate_charts.py Poetry.lock Poetry.lock power_analysis.py power_analysis.py pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json validate.py validate.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
現実世界のセキュリティ脆弱性の修正に関して LLM エージェントを評価するためのベンチマーク。エージェントはサンドボックス化された Docker コンテナ内で実行され、メンテナのセキュリティ テスト スイートに対してスコア付けされます。
環境内の OPENAI_API_KEY 、 ANTHROPIC_API_KEY 、および/または POOLSIDE_API_KEY (または .env ファイル)
pip 詩をインストールする
詩のインストール
タスク構造
各タスクは、tasks/{CVE-ID}/ の下に存在し、以下が含まれます。
タスク/CVE-2026-33175/
§── meta.json # GHSA ID、CWE、CVSS、リポジトリ URL、脆弱な SHA と修正された SHA
§── setup.sh # リポジトリをクローンし、脆弱な SHA をチェックアウトし、依存関係をインストールします
§── run_tests.sh # test_security.py をリポジトリに挿入し、pytest を実行します
§── test_security.py # セキュリティ テスト (脆弱なコードで xfail、修正をパス)
§── Advisory.md # 完全な GHSA アドバイザリー (最もリッチなプロンプト)
§── Diagnostic.md # 動作の説明のみ - ファイル名や関数名なし
§──locate.md # ファイルと関数のみ - 欠陥の説明なし
└── Dockerfile # オプション。タスクに追加のシステム依存関係が必要な場合にのみ存在します
meta.json の例:
{
"ghsa_id" : " GHSA-xxxx-xxxx-xxxx "

、
"cwe" : [ "CWE-287 " ]、
"cvss" : 9.1 、
「リポジトリ」: {
"url" : " https://github.com/org/project " ,
"vulnerable_sha" : " abc123^ " ,
"fixed_sha" : " abc123 "
}
}
setup.sh はべき等であり、安全に再実行できます。 test_security.py は、実行中はエージェントに対して隠蔽され、エージェントの終了後にのみ挿入されます。
共有ベースイメージ ( cve-bench/base ) — Python 3.12、git、詩、およびハーネス。
タスクごとに 1 つのタスク イメージ ( cve-bench/{task-id} ) — ベースを拡張し、タスク ディレクトリをコピーし、 setup.sh を実行します。
# 特定のタスクのみをビルドする
python build.py --task CVE-2026-33175 CVE-2026-42561
# 基本イメージの再構築をスキップする
python build.py --skip-base
タスク イメージは並行して構築されます (最大 5 つのワーカー)。タスク ディレクトリに Dockerfile が含まれている場合、それが汎用の docker/task.Dockerfile の代わりに使用されます。
ベンチマークを実行する前に、各タスクのセキュリティ テストで脆弱なコードと修正されたコードが正しく区別されていることを確認します。
Python検証.py
タスクごとに、タスク コンテナー内で 3 つのフェーズが実行されます。
結果はライブテーブルとして表示されます。タスクがいずれかのフェーズに失敗した場合、終了コードは 1 になります。
# 特定のタスクのみを検証する
python validate.py --task CVE-2026-33175 GHSA-r758-8hxw-4845
# 検証前のイメージの再構築をスキップする
python validate.py --skip-build
ベンチマークの実行
python benchmark.py --model openai:gpt-5.5 poolside:laguna-m.1 --prompt-type 勧告
オプション:
実行ごとに、 results/ に JSON 結果ファイルが生成されます。
results/{タスクID}__{プロバイダー}:{モデル}__{プロンプトタイプ}.json
既存の結果ファイルは自動的にスキップされます。実行はタスク (最大 20 ワーカー) 間で同時に実行され、429 を回避するためにプロバイダーごとのレート制限 (プロバイダーごとに一度に 1 つのアクティブなリクエスト) が設定されます。
各結果ファイルは、次の構造を持つ JSON オブジェクトです。
{
"cve_id" : " CVE-2026-33175 " 、
"モデルID": "

openai:gpt-5.5 " ,
"prompt_type" : "アドバイザリー" ,
"タイムスタンプ" : " 2026-05-01T12:00:00 " ,
"モデル期間_s" : 142.3 、
"test_duration_s" : 8.1 ,
「ターン」 : [
{
"tool_calls_and_results" : [ ... ],
"input_tokens" : 12400 、
"output_tokens" : 310
}
]、
「テスト」: [
{
"種類" : "セキュリティ" ,
"name" : " test_email_verified " ,
"結果" : " 合格 "
}
】
}
testing[].kind は、「security」( test_security.py から) または「regression」(プロジェクト独自のテスト スイートから) のいずれかです。すべてのセキュリティ テストに合格し、回帰テストに失敗しない場合にのみ、実行は解決されたとみなされます。
python generate_charts.py
すべての結果ファイルを results/ から読み取り、SVG チャートを docs/images/charts/ に書き込みます。 Bokeh のヘッドレス エクスポートには Chrome/Chromium が必要です ( chromedriver-binary 経由)。
ハーネスは、各 Docker コンテナ内で python -m harness.run として実行されます。プロンプトのロード、エージェント ループの実行、結果ファイルの書き込みを行います。
ソース/ハーネス/
├── run.py # Entry point;引数を解析し、コンポーネントを接続し、BenchmarkRunner を呼び出します
├── client/
│ §── Factory.py # Provider:model-id を解析し、正しい LLMClient を返す
│ §── _client.py # 抽象 LLMClient、ToolCall、LLMTurn データクラス
│ §── anthropic.py # Anthropic SDK の統合
│ └─ oai.py # OpenAI SDK 統合 (Poolside にも使用)
├── agent/
│ §── core.py # エージェントループ: クライアントの呼び出し、ツール呼び出しのディスパッチ、メッセージのスレッド化
│ └──runner.py # エージェントをラップし、タイミングとターンリストを追跡します
├── bench/
│ §──runner.py # セットアップ → エージェント → セキュリティ テスト → 回帰テストをオーケストレーションします
│ §── result.py # BenchmarkResult および TestResult データクラス、JSON シリアル化
│ └─ repository.py # 結果ファイルをディスクに書き込みます
└── タスク/
├── tools.py # Tool implem

エントリ: ListFiles、ReadFile、SearchInFiles、
│ # EditFile、CreateFile、DeleteFile、RunPytest
━──prompt_loader.py # Advisory.md / Diagnostic.md / Locate.md を読み込みます
エージェントが利用できるツール:
すべてのツールは、ディレクトリのトラバーサルを防ぐために、リポジトリ ルートに対してパスを検証します。エージェントは test_security.py または git 履歴にアクセスできません。
エージェント ループは最大 20 ターン実行されます。ターン上限に達した場合、実行はそのまま記録され、エージェントがリポジトリから離れた状態にかかわらず、セキュリティ テストは引き続き実行されます。
Tasks/{CVE-ID}/ を作成し、 meta.json 、 setup.sh 、 run_tests.sh 、 test_security.py 、 Advisory.md 、 Diagnostic.md 、 Locate.md を追加します。
setup.sh と run_tests.sh を実行可能にします ( chmod +x )。
検証: python validate.py --task {CVE-ID} 。
ビルド: python build.py --task {CVE-ID} 。
この研究は自主研究として実施されました。調査を実施し、このリポジトリを準備した時点では、私は所属機関に所属していませんでした。
@misc { gattipinheiro2026cvebench 、
著者 = { ガッティ・ピニェイロ、ジョバンニ } 、
title = { {CVE-Bench}: 現実世界のセキュリティ脆弱性修正に関する {LLM} エージェントのベンチマーク } ,
年 = { 2026 } 、
howpublished = { \url{https://giovannigatti.github.io/cve-bench} } ,
note = { コードは \url{https://github.com/GiovanniGatti/cve-bench} で入手できます }
}
ライセンス
現実世界のセキュリティ脆弱性の修正に関して AI エージェントを評価するためのベンチマーク。
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

A benchmark for evaluating AI agents on fixing real-world security vulnerabilities. - GiovanniGatti/cve-bench

GitHub - GiovanniGatti/cve-bench: A benchmark for evaluating AI agents on fixing real-world security vulnerabilities. · GitHub
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
GiovanniGatti
/
cve-bench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits docker docker docs docs results results src/ harness src/ harness tasks tasks template/ CVE-2026-XXXX template/ CVE-2026-XXXX .coverage .coverage .dockerignore .dockerignore .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md benchmark.py benchmark.py build.py build.py generate_charts.py generate_charts.py poetry.lock poetry.lock power_analysis.py power_analysis.py pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json validate.py validate.py View all files Repository files navigation
A benchmark for evaluating LLM agents on fixing real-world security vulnerabilities. Agents run inside sandboxed Docker containers and are scored against the maintainer's security test suite.
OPENAI_API_KEY , ANTHROPIC_API_KEY , and/or POOLSIDE_API_KEY in your environment (or a .env file)
pip install poetry
poetry install
Task structure
Each task lives under tasks/{CVE-ID}/ and contains:
tasks/CVE-2026-33175/
├── meta.json # GHSA ID, CWE, CVSS, repo URL, vulnerable and fixed SHAs
├── setup.sh # Clones repo, checks out the vulnerable SHA, installs dependencies
├── run_tests.sh # Injects test_security.py into the repo and runs pytest
├── test_security.py # Security tests (xfail on vulnerable code, pass on the fix)
├── advisory.md # Full GHSA advisory (richest prompt)
├── diagnose.md # Behavioural description only — no file or function names
├── locate.md # File and function only — no description of the flaw
└── Dockerfile # Optional; only present when the task needs extra system deps
meta.json example:
{
"ghsa_id" : " GHSA-xxxx-xxxx-xxxx " ,
"cwe" : [ " CWE-287 " ],
"cvss" : 9.1 ,
"repo" : {
"url" : " https://github.com/org/project " ,
"vulnerable_sha" : " abc123^ " ,
"fixed_sha" : " abc123 "
}
}
setup.sh is idempotent and safe to re-run. test_security.py is kept hidden from the agent during the run and injected only after the agent finishes.
A shared base image ( cve-bench/base ) — Python 3.12, git, poetry, and the harness.
One task image per task ( cve-bench/{task-id} ) — extends the base, copies the task directory, and runs setup.sh .
# Build specific tasks only
python build.py --task CVE-2026-33175 CVE-2026-42561
# Skip rebuilding the base image
python build.py --skip-base
Task images are built in parallel (up to 5 workers). If a task directory contains a Dockerfile , it is used instead of the generic docker/task.Dockerfile .
Before running the benchmark, verify that each task's security tests correctly distinguish vulnerable from fixed code:
python validate.py
For each task, this runs three phases inside the task container:
Results are displayed as a live table. Exit code is 1 if any task fails any phase.
# Validate specific tasks only
python validate.py --task CVE-2026-33175 GHSA-r758-8hxw-4845
# Skip rebuilding images before validation
python validate.py --skip-build
Running the benchmark
python benchmark.py --model openai:gpt-5.5 poolside:laguna-m.1 --prompt-type advisory
Options:
Each run produces a JSON result file in results/ :
results/{task-id}__{provider}:{model}__{prompt-type}.json
Existing result files are skipped automatically. Runs execute concurrently across tasks (up to 20 workers), with per-provider rate limiting (one active request per provider at a time) to avoid 429s.
Each result file is a JSON object with the following structure:
{
"cve_id" : " CVE-2026-33175 " ,
"model_id" : " openai:gpt-5.5 " ,
"prompt_type" : " advisory " ,
"timestamp" : " 2026-05-01T12:00:00 " ,
"model_duration_s" : 142.3 ,
"test_duration_s" : 8.1 ,
"turns" : [
{
"tool_calls_and_results" : [ ... ],
"input_tokens" : 12400 ,
"output_tokens" : 310
}
],
"tests" : [
{
"kind" : " security " ,
"name" : " test_email_verified " ,
"outcome" : " passed "
}
]
}
tests[].kind is either "security" (from test_security.py ) or "regression" (from the project's own test suite). A run is considered solved only if all security tests pass and no regression tests fail.
python generate_charts.py
Reads all result files from results/ and writes SVG charts to docs/images/charts/ . Requires Chrome/Chromium for Bokeh's headless export (via chromedriver-binary ).
The harness runs inside each Docker container as python -m harness.run . It is responsible for loading the prompt, running the agentic loop, and writing the result file.
src/harness/
├── run.py # Entry point; parses args, wires components, calls BenchmarkRunner
├── client/
│ ├── factory.py # Parses provider:model-id, returns the correct LLMClient
│ ├── _client.py # Abstract LLMClient, ToolCall and LLMTurn dataclasses
│ ├── anthropic.py # Anthropic SDK integration
│ └── oai.py # OpenAI SDK integration (also used for Poolside)
├── agent/
│ ├── core.py # Agentic loop: calls client, dispatches tool calls, threads messages
│ └── runner.py # Wraps Agent, tracks timing and turn list
├── bench/
│ ├── runner.py # Orchestrates setup → agent → security tests → regression tests
│ ├── result.py # BenchmarkResult and TestResult dataclasses, JSON serialisation
│ └── repository.py # Writes result files to disk
└── task/
├── tools.py # Tool implementations: ListFiles, ReadFile, SearchInFiles,
│ # EditFile, CreateFile, DeleteFile, RunPytest
└── prompt_loader.py # Reads advisory.md / diagnose.md / locate.md
Tools available to the agent:
All tools validate paths against the repository root to prevent directory traversal. The agent does not have access to test_security.py or to the git history.
The agent loop runs for at most 20 turns. If the turn ceiling is reached, the run is recorded as-is and the security tests are still executed against whatever state the agent left the repository in.
Create tasks/{CVE-ID}/ and add meta.json , setup.sh , run_tests.sh , test_security.py , advisory.md , diagnose.md , locate.md .
Make setup.sh and run_tests.sh executable ( chmod +x ).
Validate: python validate.py --task {CVE-ID} .
Build: python build.py --task {CVE-ID} .
This work was conducted as independent research. At the time of conducting the research and preparing this repository, I had no institutional affiliation.
@misc { gattipinheiro2026cvebench ,
author = { Gatti Pinheiro, Giovanni } ,
title = { {CVE-Bench}: Benchmarking {LLM} Agents on Real-World Security Vulnerability Fixes } ,
year = { 2026 } ,
howpublished = { \url{https://giovannigatti.github.io/cve-bench} } ,
note = { Code available at \url{https://github.com/GiovanniGatti/cve-bench} }
}
License
A benchmark for evaluating AI agents on fixing real-world security vulnerabilities.
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
