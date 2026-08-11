---
source: "https://github.com/theirish81/frags"
hn_url: "https://news.ycombinator.com/item?id=49257009"
title: "Frags framework for precision in AI agents"
article_title: "GitHub - theirish81/frags: Frags is an advanced AI agent for complex data workflows—retrieval, transformation, extraction, and aggregation. Highly customizable and extensible, it prioritizes precision. · GitHub"
author: "mansilladev"
captured_at: "2026-08-11T12:41:36Z"
capture_tool: "hn-digest"
hn_id: 49257009
score: 1
comments: 0
posted_at: "2026-08-11T12:11:51Z"
tags:
  - hacker-news
  - translated
---

# Frags framework for precision in AI agents

- HN: [49257009](https://news.ycombinator.com/item?id=49257009)
- Source: [github.com](https://github.com/theirish81/frags)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T12:11:51Z

## Translation

タイトル: AI エージェントの精度を高めるための Frags フレームワーク
記事のタイトル: GitHub - theirish81/frags: Frags は、複雑なデータ ワークフロー (取得、変換、抽出、集計) のための高度な AI エージェントです。カスタマイズ性と拡張性が高く、精度を優先します。 · GitHub
説明: Frags は、複雑なデータ ワークフロー (取得、変換、抽出、集計) のための高度な AI エージェントです。カスタマイズ性と拡張性が高く、精度を優先します。 - theirish81/フラグス

記事本文:
GitHub - theirish81/frags: Frags は、複雑なデータ ワークフロー (取得、変換、抽出、集計) のための高度な AI エージェントです。カスタマイズ性と拡張性が高く、精度を優先します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
彼らっぽい81
/
破片
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
218 コミット 218 コミット .github/ workflows .github/ workflows anthropic anthropic chatgpt cha

tgpt cli cli エバリュエーター エバリュエーター 例 例 gemini gemini httpfactory httpfactory log log mcpauth mcpauth ollama ollama リソース リソース スキーマ スキーマscoper スコープr scriptengines scriptengines spec spec test_data test_data util util .gitignore .gitignore ライセンス ライセンス README.md README.md ai.go ai.go config.go config.go config_test.go config_test.go context.go context.go context_config.go context_config.go context_test.go context_test.go dependency.go dependency.go dependency_test.go dependency_test.go external_functions.go external_functions.go external_functions_test.go external_functions_test.go function_callers.go function_callers.go go.mod go.mod go.sum go.sum mcp.go mcp.gorunner.gorunner.gorunner_test.gorunner_test.goscripting.goscripting.gosession_manager.gosession_manager.gosession_manager_test.gosession_manager_test.gosessions.gosessions.gosessions_test.gosessions_test.go thread_safe.go thread_safe.go thread_safe_test.go thread_safe_test.go tools.go tools.goトランスフォーマー.go トランスフォーマー.go トランスフォーマー_テスト.ゴー トランスフォーマー_テスト.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
注: プロジェクトはまだ開発中ですが、すでに試すことができます。
Frags は、データの取得、変換、抽出の複雑なワークフローの実行に特化した高度な AI/LLM エージェントです
そして集約。高度にカスタマイズ可能で拡張可能であるように設計されているため、独自のものと統合できます。
ツールとプロセス。その主な目標は精度と集中力であり、エンジニアと専門家に特化したシステムです。
コード不要のクイックフィックスではなく、専門家に相談してください。
Frags は、CLI ツールとして、また Golang プロジェクトに統合されるライブラリとして提供されます。
マルチ LLM: Frags は複数の LLM をサポートしているため、ニーズに最も適した LLM を選択できます。
ほぼ専用

構造化コンテンツの作成専用: Frags の目的は、高度な機能に統合されることです。
したがって、その出力は完全に予測可能であり、マシンで利用できる必要があります。
オーケストレーション システム: Frags は、質問すると単純に答えが返ってくるエージェントではありません。全体
Frags の目的は、ユーザーが複雑なデータの取得、変換、抽出、集計を記述できるようにすることです。
複雑なデータ構造を生成します。
ツールの高度なサポート: Frags には、内部システムと統合するための標準化されたシステム全体があります (次のように: によって提供されます)。
インテグレータ）と外部ツール（MCP サーバーなど）。
コンテキストの肥大化防止: Frags マルチセッション システムにより、LLM に存在するものを定義および整理できます。
セッションのタスクに基づいてコンテキストを分析し、集中力を高め、幻覚のリスクを軽減します。
出力のセグメント化: Frags を使用すると、出力を複数の部分にセグメント化して、問題を克服できるようになります。
出力トークンの制限と回答の品質の向上。
高度な前処理/後処理: Frags を使用すると、カスタムの前処理/後処理ステップ、スクリプト、ツール、および
変圧器、不必要な LLM 作業量の削減、コストの削減、パフォーマンスと応答の向上
品質。
モジュール性: Frags は簡単に拡張できるように設計されており、新しい機能を追加して統合することができます。
自分だけのツール。
研究/論文: ニーズが単純な答えを得るだけに留まらず、全体的な構造化された研究が必要な場合
洗練されたトピックについて。コンテキストの削減、フォーカスの強化、インターネット検索と組み合わせたドキュメントの取り込み
論文の各セクションにどのような内容を含めるかを設計できるため、LLM は各目的に焦点を当てることができ、
必要に応じて処理できるデータを生成します。
データ抽出: Frags を使用すると、複雑なデータ抽出を定義できます。

n パイプライン、からデータを抽出できます。
文書を作成し、出力が構造化され、予測可能であることを確認します。これにより、他のシステムに簡単に接続できるようになります。
事前定義されたフィールドと値が期待されます。
データ変換/分析: データ取得 (インターネット、データベース、または利用可能な MCP ツール経由) から
分析または変換では、Frags はプロセスをガイドし、コンテキストをざっと読み取って堅牢なデータ構造を提供できます。
結果の信頼性が高まります。
レポート: フラグをデータ ソースに接続し、ステータス全体を説明する複雑なレポート テンプレートを定義します。
システム、部門、または会社の。出力をレポート ツールに接続して、高品質のレポートを作成します。
メモの拡張: Frags にメモを与え、LLM がメモを完全な機能に拡張する方法を設計します。
文書。
チャットボットの拡張: チャットボットを「応答マシン」から「説明エンジン」に改善します。
クリエイティブなライティング: フラグを使用してクリエイティブなコンテンツを生成できるため、ライティングの方法をデザインできます。
どのように見えるか、そしてどのように構造化されるべきか。
Frags Wiki で完全なドキュメントを見つけてください。
Frags は、複雑なデータ ワークフロー (取得、変換、抽出、集計) のための高度な AI エージェントです。カスタマイズ性と拡張性が高く、精度を優先します。
Readme AGPL-3.0 ライセンス アクティビティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Frags is an advanced AI agent for complex data workflows—retrieval, transformation, extraction, and aggregation. Highly customizable and extensible, it prioritizes precision. - theirish81/frags

GitHub - theirish81/frags: Frags is an advanced AI agent for complex data workflows—retrieval, transformation, extraction, and aggregation. Highly customizable and extensible, it prioritizes precision. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
theirish81
/
frags
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
218 Commits 218 Commits .github/ workflows .github/ workflows anthropic anthropic chatgpt chatgpt cli cli evaluators evaluators examples examples gemini gemini httpfactory httpfactory log log mcpauth mcpauth ollama ollama resources resources schema schema scoper scoper scriptengines scriptengines spec spec test_data test_data util util .gitignore .gitignore LICENSE LICENSE README.md README.md ai.go ai.go config.go config.go config_test.go config_test.go context.go context.go context_config.go context_config.go context_test.go context_test.go dependencies.go dependencies.go dependencies_test.go dependencies_test.go external_functions.go external_functions.go external_functions_test.go external_functions_test.go function_callers.go function_callers.go go.mod go.mod go.sum go.sum mcp.go mcp.go runner.go runner.go runner_test.go runner_test.go scripting.go scripting.go session_manager.go session_manager.go session_manager_test.go session_manager_test.go sessions.go sessions.go sessions_test.go sessions_test.go thread_safe.go thread_safe.go thread_safe_test.go thread_safe_test.go tools.go tools.go transformers.go transformers.go transformers_test.go transformers_test.go View all files Repository files navigation
Note: The project is still in development, but you can already try it out.
Frags is an advanced AI/LLM Agent dedicated to executing complex workflows of data retrieval, transformation, extraction
and aggregation. It is designed to be highly customizable and extensible, allowing you to integrate it with your own
tools and processes. Its main goal is precision and focus , and it's a system dedicated to engineers and
specialists rather than a code-free quick fix .
Frags comes as a CLI tool and as a library to be integrated into Golang projects .
Multi LLM: Frags supports multiple LLMs, allowing you to choose the one that best suits your needs.
Dedicated almost exclusively to producing structured content: the purpose of Frags is to be integrated in advanced
workflows, therefore its output needs to be perfectly predictable and consumable by a machine.
Orchestration system: Frags is not an agent to which you ask a question a simple answer back. The whole
purpose of Frags is to allow the user to describe complex data retrieval, transformation, extraction and aggregation
to produce complex data structures.
Advanced support for tools: Frags has a whole standardized system to integrate with internal (as in: provided by
the integrator) and external tools (as in: MCP servers).
Anti-context-bloating: Frags multi-session system allows you to define and organize what is present in the LLM
context, based on the session task, improving focus and reducing the risk of hallucinations
Output segmentation: Frags allows you to segment your output into multiple parts, allowing you to overcome
output token limitations, and improving answer quality.
Advanced pre/post-processing: Frags allows you to define custom pre/post-processing steps, scripts, tools, and
transformers, reducing the amount of LLM work where not necessary, reducing cost, improving performance and answer
quality.
Modularity: Frags is designed to be easily extensible, allowing you to add new features and integrate them with
your own tools.
Research/Paper: when your needs go beyond getting a straight answer, but need a whole structured research
on sophisticated topics. The context reduction, focus enhancement, document ingestion combined with internet search
allows you to design how your paper should contain in each section, allowing the LLM to focus on each objective and
produce data you can process however you want.
Data extraction: Frags allows you to define complex data extraction pipelines, allowing you to extract data from
documents, making sure the output is structured and predictable. This makes it easy to plug into other systems that
expect predefined fields and values.
Data transformation/analysis: From data retrieval (via the Internet, databases or any MCP tool available) to
analysis or transformation, Frags can guide the process and provide solid data structures, skimming the context and
increasing the credibility of the results.
Reporting: Connect Frags to data sources and define complex reporting templates that describe the entire status
of a system, division or company. Connect the output to a reporting tool to produce quality reports.
Notes augmentation: give Frags your notes and design how you want the LLM to expand them into a fully featured
document.
Chatbot augmentation: improve your chatbot from an "answers machine" to an "explanation engine."
Creative writing: Frags can be used to generate creative content, allowing you to design how your writing should
look like, and how it should be structured.
Find the full documentation in the Frags Wiki
Frags is an advanced AI agent for complex data workflows—retrieval, transformation, extraction, and aggregation. Highly customizable and extensible, it prioritizes precision.
Readme AGPL-3.0 license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
