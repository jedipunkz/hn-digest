---
source: "https://github.com/mantis-llm-gateway"
hn_url: "https://news.ycombinator.com/item?id=48690749"
title: "Show HN: Mantis, A self-hosted LLM gateway"
article_title: "Mantis · GitHub"
author: "rizsyed1"
captured_at: "2026-06-26T19:32:04Z"
capture_tool: "hn-digest"
hn_id: 48690749
score: 2
comments: 0
posted_at: "2026-06-26T19:15:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Mantis, A self-hosted LLM gateway

- HN: [48690749](https://news.ycombinator.com/item?id=48690749)
- Source: [github.com](https://github.com/mantis-llm-gateway)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T19:15:04Z

## Translation

タイトル: 表示 HN: Mantis、セルフホスト型 LLM ゲートウェイ
記事タイトル: Mantis · GitHub
説明: 複数の LLM と対話するための単一インターフェイスを提供するオープンソースの自己ホスト型 LLM ゲートウェイ - Mantis
HN テキスト: こんにちは、HN の皆さん - リズです。私は数人で集まり、LLM ゲートウェイを構築しました。これは、初期段階の製品に取り組む小規模なチーム向けに設計されており、単一のコマンド (つまり「mantisdeploy」) を使用して AWS にデプロイできます。これは自己ホスト型であり、ユーザーに属するように設計されています。

記事本文:
カマキリ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
mantis-llm-ゲートウェイ
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
https://mantis-llm-gateway.github.io/
Mantis はオープンソースの自己ホスト型 LL です

複数のモデル ターゲットにわたるアプリケーションを構築するチーム向けの M ゲートウェイ。これにより、ルーティング ポリシー、フェイルオーバー動作、応答キャッシュ、ガードレール、可観測性、AWS デプロイメント構成を一元化しながら、クライアント アプリケーションに 1 つの安定したチャット完了 API が提供されます。
このプロジェクトは、インフラストラクチャやデータの制御を放棄することなく LLM ゲートウェイのメリットを享受したい小規模チーム向けに設計されています。
LLM 呼び出し用の 1 つの API: 各プロバイダーと直接統合するのではなく、単一のゲートウェイ エンドポイントを通じてチャット完了リクエストを送信します。
構成可能なルーティング: メタデータ、モデル エイリアス、加重ターゲット、フォールバック チェーン、再試行、タイムアウト、クールダウンによるルーティング。
応答キャッシュ: 正確なプロンプト キャッシュとオプションのセマンティック キャッシュにより、LLM 呼び出しの繰り返しを削減します。
ガードレール: AWS Bedrock ガードレールを使用して機密データをマスクし、ポリシーに違反するプロンプトや応答をブロックします。
可観測性: CloudWatch を通じてリクエスト ID、レイテンシー、トークン使用量、キャッシュ動作、エラー、リクエストの結果をキャプチャします。
AWS ネイティブのデプロイメント: Terraform、ECS Fargate、ALB、ElastiCache、Parameter Store、S3、IAM、CloudWatch を使用して Mantis をプロビジョニングして実行します。
llm-gateway : FastAPI ゲートウェイ サービス、React 構成ダッシュボード、Terraform インフラストラクチャ、およびデプロイメント スクリプト。
mantis-sdk : アプリケーション コードから Mantis /v1/chat/completions エンドポイントを呼び出すための Python SDK。
mantis-llm-gateway.github.io : 公開ドキュメント サイトとケーススタディ。
プロジェクトの概要、ガイド、API リファレンス、アーキテクチャのケーススタディに関するドキュメントをお読みください。
クイック スタートに従って、ゲートウェイを実行または展開します。
ルーティング構成ガイドを参照して、モデルの選択、フォールバック、キャッシュ、およびクールダウンの動作がどのように制御されるかを理解してください。
Mantis は、マルチ LLM アプリケーション開発の信頼性を高めるために存在します。

、観察可能であり、運用上管理可能です。プロバイダー固有のロジックをアプリケーション コード全体に分散させる代わりに、チームはモデルのルーティング、キャッシュ ポリシー、フェイルオーバー動作、ガードレール、展開上の懸念を 1 つのゲートウェイ層の背後に置くことができます。
その結果、アプリケーション コードはシンプルのまま、モデルの選択は構成可能なままとなり、チームが独自の AWS 環境内でリクエストがどのように移動するかを制御し続けるシステムが実現します。
llm-ゲートウェイ
llm-ゲートウェイ
公共
LLM ゲートウェイ
カマキリSDK
カマキリSDK
公共
SDK
mantis-llm-gateway.github.io
mantis-llm-gateway.github.io
公共
.github
.github
公共
組織の説明
読み込み中
種類
タイプを選択してください
すべて
公共
情報源
フォーク
アーカイブ済み
鏡
テンプレート
言語
言語を選択してください
すべて
HTML
パイソン
並べ替え
注文を選択してください
最終更新日
名前
スター
4 つのリポジトリ中 4 つを表示
mantis-llm-gateway.github.io
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
llm-ゲートウェイ
公共
LLM ゲートウェイ
読み込み中にエラーが発生しました。このページをリロードしてください。
12
アパッチ-2.0
0
0
0
2026 年 6 月 25 日更新
.github
公共
組織の説明
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
この組織には公的会員はいません。この組織のメンバーを確認するには、メンバーになる必要があります。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An open-source, self-hosted LLM gateway providing a single interface for interacting with multiple LLMs - Mantis

Hey HNers - Riz here. I got together with a few guys and we built an LLM gateway. It's designed for small teams working on early-stage products, and can be deployed to AWS using a single command (i.e. `mantis deploy`). It's self-hosted, and is designed to belong to you.

Mantis · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
mantis-llm-gateway
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
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
https://mantis-llm-gateway.github.io/
Mantis is an open-source, self-hosted LLM gateway for teams building applications across multiple model targets. It gives client applications one stable chat-completions API while centralizing routing policy, failover behavior, response caching, guardrails, observability, and AWS deployment configuration.
The project is designed for small teams that want the benefits of an LLM gateway without giving up control of their infrastructure or data.
One API for LLM calls: send chat-completion requests through a single gateway endpoint instead of integrating directly with each provider.
Configurable routing: route by metadata, model aliases, weighted targets, fallback chains, retries, timeouts, and cooldowns.
Response caching: reduce repeated LLM calls with exact prompt caching and optional semantic caching.
Guardrails: use AWS Bedrock guardrails to mask sensitive data and block policy-violating prompts or responses.
Observability: capture request IDs, latency, token usage, cache behavior, errors, and request outcomes through CloudWatch.
AWS-native deployment: provision and run Mantis with Terraform, ECS Fargate, ALB, ElastiCache, Parameter Store, S3, IAM, and CloudWatch.
llm-gateway : the FastAPI gateway service, React configuration dashboard, Terraform infrastructure, and deployment scripts.
mantis-sdk : a Python SDK for calling the Mantis /v1/chat/completions endpoint from application code.
mantis-llm-gateway.github.io : the public documentation site and case study.
Read the documentation for the project overview, guides, API reference, and architecture case study.
Follow the quick start to run or deploy the gateway.
Review the routing configuration guide to understand how model selection, fallback, caching, and cooldown behavior are controlled.
Mantis exists to make multi-LLM application development more reliable, observable, and operationally manageable. Instead of spreading provider-specific logic across application code, teams can put model routing, cache policy, failover behavior, guardrails, and deployment concerns behind one gateway layer.
The result is a system where application code stays simple, model choices remain configurable, and teams keep control over how requests move through their own AWS environment.
llm-gateway
llm-gateway
Public
An LLM Gateway
mantis-sdk
mantis-sdk
Public
SDK
mantis-llm-gateway.github.io
mantis-llm-gateway.github.io
Public
.github
.github
Public
Organization description
Loading
Type
Select type
All
Public
Sources
Forks
Archived
Mirrors
Templates
Language
Select language
All
HTML
Python
Sort
Select order
Last updated
Name
Stars
Showing 4 of 4 repositories
mantis-llm-gateway.github.io
Public
Uh oh!
There was an error while loading. Please reload this page .
llm-gateway
Public
An LLM Gateway
There was an error while loading. Please reload this page .
12
Apache-2.0
0
0
0
Updated Jun 25, 2026
.github
Public
Organization description
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
This organization has no public members. You must be a member to see who’s a part of this organization.
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
