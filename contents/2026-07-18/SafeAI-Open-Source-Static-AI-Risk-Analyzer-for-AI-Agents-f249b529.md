---
source: "https://github.com/ikaruscareer/SafeAI"
hn_url: "https://news.ycombinator.com/item?id=48963061"
title: "SafeAI – Open-Source Static AI Risk Analyzer for AI Agents"
article_title: "GitHub - ikaruscareer/SafeAI · GitHub"
author: "ikaruscareer"
captured_at: "2026-07-18T22:39:13Z"
capture_tool: "hn-digest"
hn_id: 48963061
score: 1
comments: 0
posted_at: "2026-07-18T22:24:40Z"
tags:
  - hacker-news
  - translated
---

# SafeAI – Open-Source Static AI Risk Analyzer for AI Agents

- HN: [48963061](https://news.ycombinator.com/item?id=48963061)
- Source: [github.com](https://github.com/ikaruscareer/SafeAI)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T22:24:40Z

## Translation

タイトル: SafeAI – AI エージェント向けのオープンソース静的 AI リスク アナライザー
記事タイトル: GitHub - ikaruscareer/SafeAI · GitHub
説明: GitHub でアカウントを作成して、ikaruscareer/SafeAI の開発に貢献します。

記事本文:
GitHub - ikaruscareer/SafeAI · GitHub
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
イカルスキャリア
/
セーフAI
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーと

ファイル
3 コミット 3 コミット PLUGIN_TEMPLATE PLUGIN_TEMPLATE 例 例 safetyai safetyai スクリプト スクリプト テスト テスト .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md ARCHITECTURE_FOR_CONTRIBUTORS.md ARCHITECTURE_FOR_CONTRIBUTORS.md CAPABILITIES.md CAPABILITIES.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMMUNITY_PROJECTS.md COMMUNITY_PROJECTS.md CONTRIBUTING.md CONTRIBUTOR_ROADMAP.md CONTRIBUTOR_ROADMAP.md FRAMEWORK_SUPPORT.md FRAMEWORK_SUPPORT.md GITHUB_RELEASE.md GITHUB_RELEASE.md GOOD_FIRST_ISSUES.md GOOD_FIRST_ISSUES.md HOW_TO_ADD_ANALYZER.md HOW_TO_ADD_ANALYZER.md HOW_TO_ADD_CAPABILITY.md HOW_TO_ADD_CAPABILITY.md HOW_TO_ADD_FRAMEWORK.md HOW_TO_ADD_FRAMEWORK.md HOW_TO_ADD_RULE.md HOW_TO_ADD_RULE.md LABELS_GUIDE.md LABELS_GUIDE.md ライセンス ライセンス MCP_SECURITY.md MCP_SECURITY.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md RISK_MODEL.md RISK_MODEL.md ROADMAP.md ROADMAP.md RULES_REFERENCE.md RULES_REFERENCE.md SECURITY.md SECURITY.md SECURITY_MODEL.md SECURITY_MODEL.md TESTING_CORPUS_GUIDE.md TESTING_CORPUS_GUIDE.md USER_GUIDE.md USER_GUIDE.md pyproject.toml pyproject.toml 要件-dev.lock 要件-dev.lock要件.lock要件.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SafeAI — 静的 AI 機能およびリスク アナライザー
SafeAI は、AI アプリケーションのソース コードをスキャンして、セキュリティ リスク、機能の露出、ガバナンスのギャップを検出する静的分析ツールです。完全にオフラインで実行され、エージェントの実行や LLM の呼び出しは一切行われず、CI/CD パイプラインに統合されます。
🌐safeai-analyzer.ikaruscareer.com — プロジェクトのランディング ページ
従来のアプリケーション セキュリティ ツール (SAST、SCA、IaC スキャン) は、AI エージェント システム用に設計されていません。 AI アプリケーションは、次のような新たなリスク面をもたらします。
プロンプト注入 — 信頼できない入力 f

モデルのプロンプトを入力します
エージェントツールの誤用 - ファイルシステム、シェル、またはデータベースにアクセスできるエージェント
機能のスプロール — フレームワークは可視化されずに機能を公開します
MCP エクスポージャ — モデル コンテキスト プロトコルのエンドポイントとツール
ガバナンスのギャップ - 認証、権限、監査証跡の欠落
SafeAI は、展開前にフレームワーク、エージェント、ツール、機能、MCP 統合を保存時に分析することで、このギャップを埋めます。
SafeAI は、AI アプリケーションを実行せずに分析し、開発者が機能を発見し、潜在的なリスクを特定し、ソフトウェア ライフサイクルの早い段階でガバナンスを改善できるように支援します。
SafeAI は、軽量で説明しやすく、コミュニティ主導型になるように設計されており、AI 機能とリスク分析のためのオープン基盤になることを目指しています。
SafeAI は、セキュリティ ライフサイクルにおいて、ランタイム ガードレールやレッド チーム ツールの前に位置します。エージェントをステージングにデプロイする前に、コミット時にエージェントのソース コードをスキャンし、フレームワーク固有の機能、MCP の構成ミス、プロンプト インジェクション パターンを検出します。ランタイム ツール (Microsoft AGT)、評価フレームワーク (LangSmith、DeepEval)、またはレッドチーム スキャナー (Promptfoo、Garak) を置き換えるものではありません。それはそれらを補完します。最初にコード内のリスクを見つけてから、実行時に検証します。
特徴
説明
フレームワークの検出
8 つの AI エージェント フレームワークを検出および解析します
能力の発見
ファイルシステム、シェル、ネットワーク、データベース、その他の機能を識別します
AIリスク分析
加重信頼スコアリングを使用して調査結果を 7 つのリスク カテゴリに分類します
迅速なリスク分析
インジェクションパターン、デリミタの問題、システムリーク、ロールオーバーライドを検出します。
ツール分析
エージェントバインドツールとそのリスクプロファイルを特定します
メモリ分析
エージェントのワークフローでのメモリ/チェックポインタの使用状況を検出します。
MCP分析
MCP サーバー、クライアント、ツール、リソース、および VA を検出します

蓋付き構成
データ漏洩の検出
ハードコードされたシークレット、トークン、API キーにフラグを立てます
CI/CDの統合
SARIF 出力、終了コード、GitHub Actions ワークフローが含まれる
マルチフォーマットレポート
ターミナルサマリー、JSON、SARIF 2.1.0、HTML
ファイル間分析
グラフのインポート、シンボル解決、およびプロジェクト グラフ
信頼度に基づいて調停された解析
ファイルごとに複数のパーサー、来歴とマージ
仕組み
ソースコード
│
▼
フレームワークの検出 — import、config、deps を通じて AI フレームワークを識別します
│
▼
静的分析 — AST 解析、機能パターン、依存関係スキャン
│
▼
機能マッピング — フレームワーク オブジェクトを正規化されたリスク カテゴリにマッピングします。
│
▼
リスク ルール — 構成可能な重大度と重みを備えたルール エンジンを適用します
│
▼
トラストスコア — 0 ～ 100 の決定論的なカテゴリ加重スコアリング
│
▼
レポート — ターミナル、JSON、SARIF、HTML
サポートされているフレームワーク
フレームワーク
検出
発見
能力分析
リスク分析
ステータス
ランググラフ
✔
部分的
部分的
部分的
早期プレビュー
CrewAI
✔
部分的
部分的
部分的
早期プレビュー
ラングチェーン
✔
部分的
部分的
部分的
早期プレビュー
セマンティックカーネル
✔
部分的
部分的
部分的
早期プレビュー
OpenAI エージェント SDK
✔
部分的
部分的
部分的
早期プレビュー
Microsoft エージェント フレームワーク
✔
部分的
部分的
部分的
早期プレビュー
Azure AI ファウンドリ
✔
最小限
最小限
最小限
早期プレビュー
岩盤エージェント
✔
最小限
最小限
最小限
早期プレビュー
フレームワークのサポートの詳細
LangGraph — StateGraph 、 add_edge 、 binding_tools 、ノード、モデルを検出します
CrewAI — エージェント、タスク、ツール、モデルを検出します
LangChain — AgentExecutor 、 Chain 、 Tool 、 PromptTemplate 、モデルを検出します
セマンティック カーネル — Kernel.invoke 、プラグイン、関数、スキル、メモリを検出します
OpenAI Agents SDK — エージェント、ツール、ハンドオフ、MCP 参照を検出します
マイクロソフト エージェント F

ramework — AgentClient 、ツール、ワークフロー、Azure モデルを検出します
Azure AI Foundry — Azure リソースを使用した YAML 構成を検出します
Bedrock Agent — Bedrock リソースを使用して JSON 構成を検出します
フレームワーク オブジェクト レベルおよびフォールバック正規表現パターンを介した SafeAI フィンガープリント機能。各機能には、証拠、信頼スコア、解決された定義、および来歴が含まれます。
注: 一部の機能 (ブラウザ、GitHub、Slack、電子メール、RAG、人間による承認) は、主に MCP 構成分析を通じて検出されます。これらの機能のフレームワーク アダプターの検出が計画されています。
PyYAML (YAML 構成解析用)
git clone https://github.com/ikaruscareer/SafeAI.git
cd セーフAI
pip install -e 。
開発依存関係をインストールする
pip install -e " .[dev] "
CLIの使用法
python -msafeai scan <ディレクトリ> [オプション]
オプション
オプション
デフォルト
説明
ディレクトリ
必須
スキャンするパス
--サリフ
レポート.サリフ
SARIF 出力パス (スキップする空の文字列)
--json
—
JSON出力パス
--html
—
HTMLレポート出力パス
--ルール
内蔵
カスタムルールディレクトリ
--フェイルオン
クリティカル
終了コードしきい値: クリティカル、高、中
--冗長
—
詳細出力を有効にする
終了コード
コード
状態
0
閾値以上の所見なし
1
しきい値以上の検出結果が検出されました
出力例
SafeAI スキャンの概要
ファイル: 12
フレームワーク: langgraph、creai
MCP アセット: 2
全体的な AI リスク スコア: 73
クリティカル: 1
高: 3
中:5
調査結果:
[クリティカル] app.py:10 - 信頼できない入力がプロンプトに挿入されました
[高] app.py:22 - 検出された機能:shell_execution
[高] mcp.json:1 - MCP 構成で認証が定義されていません
例: MCP を使用した LangGraph エージェント
{
"フレームワーク" : " LangGraph " 、
"機能" : [ " プランナー " 、 " メモリ " 、 " ファイルシステム " 、 " MCP " ]、
「リスクスコア」：73、
「調査結果」 : 9 、
「クリティカル」：1、
」

高" : 3
}
CI/CDの統合
ワークフローは .github/workflows/ci.yml に含まれています。プロジェクトで使用するには:
ジョブ:
セーフアイスキャン:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 使用:actions/setup-python@v5
付き:
Python バージョン: ' 3.12 '
- 名前 : SafeAI をインストールする
実行: |
pip install -e 。
- 名前 : スキャンの実行
実行: |
Python -m セーフアイ スキャン 。 --sarif 結果.sarif --html レポート.html
- 名前 : SARIF をアップロード
使用: github/codeql-action/upload-sarif@v3
付き:
sarif_file : results.sarif
GitLab CI
セーフアイスキャン:
画像：Python：3.12
スクリプト:
-pip install -e 。
- セーフアイスキャン。 --sarif 結果.sarif --html レポート.html
アーティファクト:
パス:
- results.sarif
- レポート.html
Azure DevOps
- タスク : PythonScript@0
入力:
スクリプトソース: ' インライン '
スクリプト: |
サブプロセスのインポート
subprocess.run(["pip", "install", "-e", "."])
subprocess.run(["safeai", "scan", ".", "--sarif", "$(Build.ArtifactStagingDirectory)/results.sarif"])
SARIFの統合
SafeAI は、GitHub Advanced Security、Azure DevOps、およびその他の SARIF 準拠ツールと互換性のある SARIF 2.1.0 形式を出力します。
5 つのフェーズすべてをカバーする詳細なロードマップについては、ROADMAP.md を参照してください。
フェーズ 1 — 静的 AI リスク スキャナー (OSS) — 鋭意開発中
フェーズ 1.5 — AI コンポーネントのセキュリティ
フェーズ 2 — AI セキュリティ テスト (将来のオプション)
フェーズ 4 — エンタープライズ (商用)
フェーズ 5 — コミュニティ インテリジェンス
SafeAI は、Apache 2.0 ライセンスに基づいてリリースされています。
Apache-2.0ライセンス
行動規範
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

Contribute to ikaruscareer/SafeAI development by creating an account on GitHub.

GitHub - ikaruscareer/SafeAI · GitHub
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
ikaruscareer
/
SafeAI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits PLUGIN_TEMPLATE PLUGIN_TEMPLATE examples examples safeai safeai scripts scripts tests tests .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md ARCHITECTURE_FOR_CONTRIBUTORS.md ARCHITECTURE_FOR_CONTRIBUTORS.md CAPABILITIES.md CAPABILITIES.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMMUNITY_PROJECTS.md COMMUNITY_PROJECTS.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTOR_ROADMAP.md CONTRIBUTOR_ROADMAP.md FRAMEWORK_SUPPORT.md FRAMEWORK_SUPPORT.md GITHUB_RELEASE.md GITHUB_RELEASE.md GOOD_FIRST_ISSUES.md GOOD_FIRST_ISSUES.md HOW_TO_ADD_ANALYZER.md HOW_TO_ADD_ANALYZER.md HOW_TO_ADD_CAPABILITY.md HOW_TO_ADD_CAPABILITY.md HOW_TO_ADD_FRAMEWORK.md HOW_TO_ADD_FRAMEWORK.md HOW_TO_ADD_RULE.md HOW_TO_ADD_RULE.md LABELS_GUIDE.md LABELS_GUIDE.md LICENSE LICENSE MCP_SECURITY.md MCP_SECURITY.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md RISK_MODEL.md RISK_MODEL.md ROADMAP.md ROADMAP.md RULES_REFERENCE.md RULES_REFERENCE.md SECURITY.md SECURITY.md SECURITY_MODEL.md SECURITY_MODEL.md TESTING_CORPUS_GUIDE.md TESTING_CORPUS_GUIDE.md USER_GUIDE.md USER_GUIDE.md pyproject.toml pyproject.toml requirements-dev.lock requirements-dev.lock requirements.lock requirements.lock View all files Repository files navigation
SafeAI — Static AI Capability & Risk Analyzer
SafeAI is a static analysis tool that scans AI application source code for security risks, capability exposure, and governance gaps. It runs entirely offline, never executes agents or calls LLMs, and integrates into CI/CD pipelines.
🌐 safeai-analyzer.ikaruscareer.com — project landing page
Traditional application security tools (SAST, SCA, IaC scanning) are not designed for AI agent systems. AI applications introduce new risk surfaces:
Prompt injection — untrusted input flows into model prompts
Agent tool misuse — agents with filesystem, shell, or database access
Capability sprawl — frameworks expose capabilities without visibility
MCP exposure — Model Context Protocol endpoints and tools
Governance gaps — missing authentication, permissions, audit trails
SafeAI fills this gap by analyzing frameworks, agents, tools, capabilities, and MCP integrations at rest—before deployment.
SafeAI analyzes AI applications without executing them, helping developers discover capabilities, identify potential risks, and improve governance early in the software lifecycle.
Designed to be lightweight, explainable, and community-driven, SafeAI aims to become an open foundation for AI capability and risk analysis.
SafeAI sits before runtime guardrails and red-teaming tools in the security lifecycle. It scans agent source code at commit time — detecting framework-specific capabilities, MCP misconfigurations, and prompt injection patterns — before you ever deploy an agent to staging. It does not replace runtime tools (Microsoft AGT), evaluation frameworks (LangSmith, DeepEval), or red-teaming scanners (Promptfoo, Garak). It complements them: find the risk in code first, then validate at runtime.
Feature
Description
Framework Detection
Detects and parses 8 AI agent frameworks
Capability Discovery
Identifies filesystem, shell, network, database, and other capabilities
AI Risk Analysis
Categorizes findings into 7 risk categories with weighted trust scoring
Prompt Risk Analysis
Detects injection patterns, delimiter issues, system leak, role override
Tool Analysis
Identifies agent-bound tools and their risk profiles
Memory Analysis
Detects memory/checkpointer usage in agent workflows
MCP Analysis
Discovers MCP servers, clients, tools, resources, and validates configuration
Data Leakage Detection
Flags hardcoded secrets, tokens, and API keys
CI/CD Integration
SARIF output, exit codes, GitHub Actions workflow included
Multi-Format Reports
Terminal summary, JSON, SARIF 2.1.0, HTML
Cross-File Analysis
Import graph, symbol resolution, and project graph
Confidence-Arbitrated Parsing
Multiple parsers per file, merged with provenance
How It Works
Source Code
│
▼
Framework Detection — identifies AI frameworks via imports, configs, deps
│
▼
Static Analysis — AST parsing, capability patterns, dependency scanning
│
▼
Capability Mapping — maps framework objects to normalized risk categories
│
▼
Risk Rules — applies rule engine with configurable severity and weights
│
▼
Trust Score — deterministic category-weighted scoring from 0–100
│
▼
Reports — terminal, JSON, SARIF, HTML
Supported Frameworks
Framework
Detection
Discovery
Capability Analysis
Risk Analysis
Status
LangGraph
✔
Partial
Partial
Partial
Early Preview
CrewAI
✔
Partial
Partial
Partial
Early Preview
LangChain
✔
Partial
Partial
Partial
Early Preview
Semantic Kernel
✔
Partial
Partial
Partial
Early Preview
OpenAI Agents SDK
✔
Partial
Partial
Partial
Early Preview
Microsoft Agent Framework
✔
Partial
Partial
Partial
Early Preview
Azure AI Foundry
✔
Minimal
Minimal
Minimal
Early Preview
Bedrock Agent
✔
Minimal
Minimal
Minimal
Early Preview
Framework Support Details
LangGraph — detects StateGraph , add_edge , bind_tools , nodes, models
CrewAI — detects Agent , Task , tools, models
LangChain — detects AgentExecutor , Chain , Tool , PromptTemplate , models
Semantic Kernel — detects Kernel.invoke , plugins, functions, skills, memory
OpenAI Agents SDK — detects Agent , tools, handoffs, MCP references
Microsoft Agent Framework — detects AgentClient , tools, workflows, Azure models
Azure AI Foundry — detects YAML configurations with Azure resources
Bedrock Agent — detects JSON configurations with Bedrock resources
SafeAI fingerprints capabilities at the framework object level and via fallback regex patterns. Each capability includes evidence, confidence score, resolved definition, and provenance.
Note: Some capabilities (Browser, GitHub, Slack, Email, RAG, Human Approval) are detected primarily through MCP configuration analysis. Framework adapter detection for these capabilities is planned.
PyYAML (for YAML configuration parsing)
git clone https://github.com/ikaruscareer/SafeAI.git
cd SafeAI
pip install -e .
Install development dependencies
pip install -e " .[dev] "
CLI Usage
python -m safeai scan < directory > [options]
Options
Option
Default
Description
directory
required
Path to scan
--sarif
report.sarif
SARIF output path (empty string to skip)
--json
—
JSON output path
--html
—
HTML report output path
--rules
built-in
Custom rules directory
--fail-on
critical
Exit code threshold: critical , high , medium
--verbose
—
Enable verbose output
Exit Codes
Code
Condition
0
No findings at or above threshold
1
Finding at or above threshold detected
Example Output
SafeAI Scan Summary
Files: 12
Frameworks: langgraph, crewai
MCP assets: 2
Overall AI Risk Score: 73
critical: 1
high: 3
medium: 5
Findings:
[critical] app.py:10 - Untrusted input interpolated into prompt
[high] app.py:22 - Capability detected: shell_execution
[high] mcp.json:1 - MCP configuration does not define authentication
Example: LangGraph agent with MCP
{
"Framework" : " LangGraph " ,
"Capabilities" : [ " Planner " , " Memory " , " Filesystem " , " MCP " ],
"Risk Score" : 73 ,
"Findings" : 9 ,
"Critical" : 1 ,
"High" : 3
}
CI/CD Integration
A workflow is included at .github/workflows/ci.yml . To use in your project:
jobs :
safeai-scan :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : actions/setup-python@v5
with :
python-version : ' 3.12 '
- name : Install SafeAI
run : |
pip install -e .
- name : Run scan
run : |
python -m safeai scan . --sarif results.sarif --html report.html
- name : Upload SARIF
uses : github/codeql-action/upload-sarif@v3
with :
sarif_file : results.sarif
GitLab CI
safeai-scan :
image : python:3.12
script :
- pip install -e .
- safeai scan . --sarif results.sarif --html report.html
artifacts :
paths :
- results.sarif
- report.html
Azure DevOps
- task : PythonScript@0
inputs :
scriptSource : ' inline '
script : |
import subprocess
subprocess.run(["pip", "install", "-e", "."])
subprocess.run(["safeai", "scan", ".", "--sarif", "$(Build.ArtifactStagingDirectory)/results.sarif"])
SARIF Integration
SafeAI outputs SARIF 2.1.0 format, compatible with GitHub Advanced Security, Azure DevOps, and other SARIF-compliant tools.
See ROADMAP.md for the detailed roadmap covering all 5 phases:
Phase 1 — Static AI Risk Scanner (OSS) — in active development
Phase 1.5 — AI Component Security
Phase 2 — AI Security Testing (optional future)
Phase 4 — Enterprise (Commercial)
Phase 5 — Community Intelligence
SafeAI is released under the Apache 2.0 License.
Apache-2.0 license
Code of conduct
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
