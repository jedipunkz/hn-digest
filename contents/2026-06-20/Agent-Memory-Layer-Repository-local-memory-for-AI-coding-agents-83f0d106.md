---
source: "https://github.com/ragnarok268/agent-memory-layer"
hn_url: "https://news.ycombinator.com/item?id=48610121"
title: "Agent Memory Layer: Repository-local memory for AI coding agents"
article_title: "GitHub - ragnarok268/agent-memory-layer: A persistent memory layer for autonomous coding agents that preserves project intent, decisions, and handoff context. · GitHub"
author: "einherjarlabs"
captured_at: "2026-06-20T18:36:58Z"
capture_tool: "hn-digest"
hn_id: 48610121
score: 3
comments: 0
posted_at: "2026-06-20T15:48:58Z"
tags:
  - hacker-news
  - translated
---

# Agent Memory Layer: Repository-local memory for AI coding agents

- HN: [48610121](https://news.ycombinator.com/item?id=48610121)
- Source: [github.com](https://github.com/ragnarok268/agent-memory-layer)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T15:48:58Z

## Translation

タイトル: エージェント メモリ層: AI コーディング エージェントのリポジトリ ローカル メモリ
記事のタイトル: GitHub - ragnarok268/agent-memory-layer: プロジェクトの意図、決定、ハンドオフ コンテキストを保存する自律コーディング エージェント用の永続メモリ層。 · GitHub
説明: プロジェクトの意図、決定、ハンドオフ コンテキストを保存する自律コーディング エージェント用の永続メモリ層。 - ragnarok268/エージェントメモリ層

記事本文:
GitHub - ragnarok268/agent-memory-layer: プロジェクトの意図、決定、ハンドオフ コンテキストを保持する自律コーディング エージェント用の永続メモリ層。 · GitHub
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
ラグナロク268
/
エージェントメモリ層
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
ragnarok268/エージェントメモリ層
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .cursor/ rules .cursor/ rules .github/ workflows .github/ workflows artifacts/ ナレッジアーティファクト/ ナレッジ自動化自動化実験/ ab_adoption 実験/ ab_adoption テスト テスト .gitattributes .gitattributes .gitignore .gitignore ADOPTION_DRILL.md ADOPTION_DRILL.md AGENTS.md AGENTS.md AGENT_BOOTSTRAP.md AGENT_BOOTSTRAP.md AGENT_GUIDE.md AGENT_GUIDE.md ARTIFACT_MODEL.md ARTIFACT_MODEL.md AUTOMATION_ARCHITECTURE.md AUTOMATION_ARCHITECTURE.md CI_AND_HOOKS.md CI_AND_HOOKS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md EVENT_MODEL.md EVENT_MODEL.md EVIDENCE.md EVIDENCE.md EXAMPLES.md EXAMPLES.md GEMINI.md GEMINI.md ライセンス ライセンスマニフェスト.md MANIFESTO.md README.md README.md ROADMAP.md ROADMAP.md WORKFLOW.md WORKFLOW.md pytest.ini pytest.ini すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Agent Memory Layer は、AI 支援ソフトウェアの動作を簡単にレビュー、修復し、後で継続できるようにするための、実験的なドキュメント優先のワークフローです。
このリポジトリは、AI 支援エンジニアリング ワークフロー向けに設計されています。人間と、Codex、Cursor、Claude Code、Gemini、または同様のシステムなどの AI コーディング エージェントの両方が読み取ることができる、リポジトリ ローカルのメモリ、意図、決定、および証拠のアーティファクトを提供します。
従来の Python ライブラリ、パッケージ、SDK、フレームワーク、またはエンドユーザー アプリケーションを意図したものではありません。付属の Python スクリプトは、ルーティング チェックとアーティファクトの書き込みのためのシン リポジトリ ローカル自動化ヘルパーです。
これは、次の 3 つの補完的な機能を中心に構築されています。
意図検証 (IA が代表)
DS2 に代表される依存性と機能の認識
エンジニア

SCP で表される ng メモリ
このリポジトリは業界標準ではありません。これはオープンソースの方法論と研究の方向性であり、現在ローカル自動化と再現可能な A/B トライアルによって評価されています。
人間の場合: README.md 、 WORKFLOW.md 、および EVIDENCE.md から始めます。
AI エージェントの場合: AGENT_BOOTSTRAP.md 、 AGENTS.md 、および ARTIFACT_MODEL.md から始めます。
AI はコードを迅速に生成できますが、リポジトリは周囲のエンジニアリング メモリを失うことがよくあります。
どの機能表面が変化したか
変更の出荷を裏付ける証拠は何ですか
その記憶が失われた場合、将来のすべての人間または AI エージェントはそれを再発見する必要があります。
このプロジェクトは以下に最も関連します:
コーディング エージェントに大きく依存する AI 支援開発者
ドメインの専門家が内部ツールを構築
AI コーディング エージェントを実験するチーム
より明確な意図、レビュー、コンテキスト保持の習慣を求めるジュニアおよび中レベルの開発者
耐久性のあるエンジニアリングメモリと再現可能なハンドオフを必要とする経験豊富なエンジニア
通常、使い捨てのスクリプト、些細なプロトタイプ、AI をほとんど使用していないチーム、またはすでに強力で耐久性のあるエンジニアリング メモリの実践を行っているチームにはあまり役に立ちません。
アイデア
-> AIがコードを生成
-> IA が意図を確認する
-> DS2 マップの依存関係と機能面
-> SCPは重要な場合には論理的根拠を保持します
-> AIによる修復か人間によるレビュー
-> 証拠を添えて発送します
目標は、保存されたエンジニアリングのコンテキストを手動の儀式ではなく、静かなインフラストラクチャのように感じさせることです。
このリポジトリはドキュメント優先であり、薄いローカル オートメーション レイヤーを使用します。
リポジトリ自体にはパッケージのインストール手順はありません。
この README と WORKFLOW.md をお読みください。
Codex、Cursor、Claude Code、Gemini、または同様のエージェントを使用している場合は、 AGENT_BOOTSTRAP.md を読んでください。
Python -m pytest
ドキュメントやコードに小さな変更を加えます。
変更されたファイルに対してガードレール ランナーを実行します。

Python オートメーション/guardrail_runner.py --変更された README.md オートメーション/guardrail_runner.py
artifacts/knowledge/ で生成されたアーティファクトを確認します。
このリポジトリに必要なローカル検証は python -m pytest です。
意図されたオペレーティング モデルは、意図の検証、依存関係/機能の認識、およびエンジニアリング メモリを組み合わせたものです。実装はモジュール式であるため、軽量の使用では個々のツールを省略または置き換えることができますが、完全な方法論ではこれらの機能が連携して機能することが前提となっています。
リポジトリローカル検証では、外部ツールのインストールはオプションです。
ia は意図の検証を有効にします。
ds2 は、依存関係と機能の表面スキャンを有効にします。
ここではSCPはローカルドラフトアーティファクトによって表現されています。別個の SCP プロジェクトは、より広範な意思決定保存のリファレンス実装を提供します。
ia または ds2 がインストールされていない場合、ランナーはそれらが失敗するのではなく、スキップされたと報告します。つまり、完全なコンパニオン ツールチェーンをインストールしなくても、リポジトリ ローカルの概念実証をテストできるということです。
現在どのような証拠が存在するのか
このリポジトリの現在の証拠は暫定的なものであり、ローカルなものです。
これまでの Codex A/B トライアルで観察されたこと:
ワークフローが有効な状態でアーティファクトをより適切に使用できるようになりました。
現在の最も強力なサマリーは EVIDENCE.md です。再現可能な実験ハーネスは、 Experiments/ab_adoption にあります。
普遍的な品質の向上
有効性に対する既知の脅威には、プロジェクト履歴全体にわたる小規模なタスク セット、ローカルのみの実行、および混合タイミング手法が含まれます。
次の検証マイルストーンは、より広範な外部使用、つまりより多くのモデル、より多くの開発者、長期プロジェクト、そして実際のケーススタディです。
概要と最初の 30 分: README.md
オートメーション設計: AUTOMATION_ARCHITECTURE.md
エージェント操作説明書：AGENTS.md
アーティファクトモデル: ARTIFACT_MODEL.md
実験方法: Experiments/ab_adoption/README.md

貢献ガイダンス: COTRIBUTING.md
DS2: github.com/ragnarok268/DS2
SCP: github.com/ragnarok268/scp
このプロジェクトは MIT ライセンスに基づいてライセンスされています。
フィードバックは具体的な場合に最も役立ちます。
どのアーティファクトが役に立ったのか、またはノイズが多かったのか
どの実験ステップが再現できなかったのか
追加の検証によって自信が変わるのはどれですか
GitHub でリポジトリを公開する場合、最も明確なフィードバック チャネルは、具体的な再現、批判、または提案された実験の改善に関する問題です。
現時点で最も安全な枠組みは次のとおりです。このリポジトリは、実証済みの標準ではなく、予備的なローカル証拠を備えた、もっともらしいエージェント メモリ ワークフローを示しています。
プロジェクトの意図、決定、およびハンドオフのコンテキストを保存する自律コーディング エージェント用の永続メモリ層。
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

A persistent memory layer for autonomous coding agents that preserves project intent, decisions, and handoff context. - ragnarok268/agent-memory-layer

GitHub - ragnarok268/agent-memory-layer: A persistent memory layer for autonomous coding agents that preserves project intent, decisions, and handoff context. · GitHub
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
ragnarok268
/
agent-memory-layer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
ragnarok268/agent-memory-layer
master Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .cursor/ rules .cursor/ rules .github/ workflows .github/ workflows artifacts/ knowledge artifacts/ knowledge automation automation experiments/ ab_adoption experiments/ ab_adoption tests tests .gitattributes .gitattributes .gitignore .gitignore ADOPTION_DRILL.md ADOPTION_DRILL.md AGENTS.md AGENTS.md AGENT_BOOTSTRAP.md AGENT_BOOTSTRAP.md AGENT_GUIDE.md AGENT_GUIDE.md ARTIFACT_MODEL.md ARTIFACT_MODEL.md AUTOMATION_ARCHITECTURE.md AUTOMATION_ARCHITECTURE.md CI_AND_HOOKS.md CI_AND_HOOKS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md EVENT_MODEL.md EVENT_MODEL.md EVIDENCE.md EVIDENCE.md EXAMPLES.md EXAMPLES.md GEMINI.md GEMINI.md LICENSE LICENSE MANIFESTO.md MANIFESTO.md README.md README.md ROADMAP.md ROADMAP.md WORKFLOW.md WORKFLOW.md pytest.ini pytest.ini View all files Repository files navigation
Agent Memory Layer is an experimental, documentation-first workflow for making AI-assisted software work easier to review, repair, and continue later.
This repository is designed for AI-assisted engineering workflows. It provides repository-local memory, intent, decision, and evidence artifacts that can be read by both humans and AI coding agents such as Codex, Cursor, Claude Code, Gemini, or similar systems.
It is not intended to be a conventional Python library, package, SDK, framework, or end-user application. The included Python scripts are thin repo-local automation helpers for routing checks and writing artifacts.
It is built around three complementary capabilities:
intent verification, represented by IA
dependency and capability awareness, represented by DS2
engineering memory, represented by SCP
This repository is not an industry standard. It is an open-source methodology and research direction that is currently evaluated with local automation and reproducible A/B trials.
For humans: start with README.md , WORKFLOW.md , and EVIDENCE.md .
For AI agents: start with AGENT_BOOTSTRAP.md , AGENTS.md , and ARTIFACT_MODEL.md .
AI can generate code quickly, but repositories often lose the surrounding engineering memory:
what capability surface changed
what evidence supports shipping the change
When that memory is missing, every future human or AI agent has to rediscover it.
This project is most relevant to:
AI-assisted developers who rely heavily on coding agents
domain experts building internal tools
teams experimenting with AI coding agents
junior and mid-level developers who want clearer intent, review, and context-preservation habits
experienced engineers who want durable engineering memory and reproducible handoff
It is usually less useful for throwaway scripts, trivial prototypes, teams with little AI usage, or teams that already have strong durable engineering-memory practices.
Idea
-> AI generates code
-> IA verifies intent
-> DS2 maps dependency and capability surfaces
-> SCP preserves rationale when it matters
-> AI repairs or a human reviews
-> ship with evidence
The goal is to make preserved engineering context feel more like quiet infrastructure than manual ceremony.
This repo is documentation-first and uses a thin local automation layer.
There is no package install step for the repo itself.
Read this README and WORKFLOW.md .
If you are using Codex, Cursor, Claude Code, Gemini, or a similar agent, read AGENT_BOOTSTRAP.md .
python -m pytest
Make a small documentation or code change.
Run the guardrail runner on the changed files:
python automation/guardrail_runner.py --changed README.md automation/guardrail_runner.py
Review the generated artifacts under artifacts/knowledge/ .
Required local validation for this repository is python -m pytest .
The intended operating model combines intent verification, dependency/capability awareness, and engineering memory. The implementation is modular, so lightweight use can omit or replace individual tools, but the complete methodology assumes these capabilities work together.
External tool installation is optional for repo-local validation:
ia enables intent verification.
ds2 enables dependency and capability-surface scanning.
SCP is represented here by local draft artifacts; the separate SCP project provides the broader decision-preservation reference implementation.
If ia or ds2 are not installed, the runner reports them as skipped rather than failing. That means the repo-local proof of concept can still be tested without installing the full companion toolchain.
What evidence currently exists
Current evidence in this repo is preliminary and local.
Observed in the Codex A/B trials so far:
better artifact usage in the workflow-enabled condition
The strongest current summary is EVIDENCE.md . The reproducible experiment harness lives in experiments/ab_adoption .
universal quality improvements
Known threats to validity include small task sets, local-only runs, and mixed timing methodologies across the project history.
The next validation milestone is broader external use: more models, more developers, longer projects, and real-world case studies.
Overview and first 30 minutes: README.md
Automation design: AUTOMATION_ARCHITECTURE.md
Agent operating instructions: AGENTS.md
Artifact model: ARTIFACT_MODEL.md
Experiment methodology: experiments/ab_adoption/README.md
Contribution guidance: CONTRIBUTING.md
DS2: github.com/ragnarok268/DS2
SCP: github.com/ragnarok268/scp
This project is licensed under the MIT License .
Feedback is most useful when it is concrete:
which artifact was useful or noisy
which experiment step was not reproducible
which additional validation would change your confidence
If you publish the repository on GitHub, the clearest feedback channel is an issue with a concrete reproduction, criticism, or suggested experiment improvement.
For now, the safest framing is: this repository shows a plausible agent-memory workflow with preliminary local evidence, not a proven standard.
A persistent memory layer for autonomous coding agents that preserves project intent, decisions, and handoff context.
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
