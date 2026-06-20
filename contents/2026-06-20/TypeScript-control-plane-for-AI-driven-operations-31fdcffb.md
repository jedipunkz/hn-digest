---
source: "https://github.com/AI-gateway-systems/attestor"
hn_url: "https://news.ycombinator.com/item?id=48609772"
title: "TypeScript control plane for AI driven operations"
article_title: "GitHub - AI-gateway-systems/attestor: AI Consequence Control Plane · GitHub"
author: "Oxlamarr"
captured_at: "2026-06-20T15:37:23Z"
capture_tool: "hn-digest"
hn_id: 48609772
score: 3
comments: 0
posted_at: "2026-06-20T15:10:46Z"
tags:
  - hacker-news
  - translated
---

# TypeScript control plane for AI driven operations

- HN: [48609772](https://news.ycombinator.com/item?id=48609772)
- Source: [github.com](https://github.com/AI-gateway-systems/attestor)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T15:10:46Z

## Translation

タイトル: AI 主導の操作のための TypeScript コントロール プレーン
記事のタイトル: GitHub - AI ゲートウェイ システム/認証者: AI Consequence Control Plane · GitHub
説明: AI 結果コントロール プレーン。 GitHub でアカウントを作成して、AI ゲートウェイ システム/認証者の開発に貢献します。

記事本文:
GitHub - AI ゲートウェイ システム/認証者: AI Consequence コントロール プレーン · GitHub
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
AIゲートウェイシステム
/
認証者
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスターブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
1,869 コミット 1,869 コミット .github .github .well-known .well-known docs docs 例 例 フィクスチャ フィクスチャ ops ops スクリプト スクリプト 仕様 仕様 src src テスト テスト ベンダー/ schematron/ 2026-CMS-QRDA-III ベンダー/ schematron/ 2026-CMS-QRDA-III .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.dr.yml docker-compose.dr.yml docker-compose.ha.yml docker-compose.ha.yml docker-compose.observability.yml docker-compose.observability.yml docker-compose.yml docker-compose.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バッジはリポジトリの証拠を指します。
アテスターが既存のシステムに接続する方法
リスクの高い AI 主導の運用のための制御インフラストラクチャ。
認証者は、AI が準備した操作とそれを実行するシステムの間に位置します。
プロンプトは行動を導くことはできますが、それを強制したり、危険な行動を止めたりすることはできません。
不正なサービス呼び出し、または範囲外のサービス呼び出し。
安全でないリクエストは、幻覚、古いコンテキスト、汚染されたツールの出力などから発生する可能性があります。
リプレイ、承認の欠落、または敵対的なコンテンツ。何かが実行される前に、認証者はチェックします
ポリシー、権限、スコープ、鮮度、再生、および証拠の場合は、 accept を返します。
絞り込み、レビュー、またはブロック。
顧客所有のゲートは実行前に決定します。トレイルには提案された内容が記録されています。
何がチェックされたのか、なぜそれが保留または許可されたのか。
AI システムはチャットから、支払い、データ、
アクセス、顧客メッセージ、インフラストラクチャ、およびプログラム可能なお金。
それはもはやプロンプトの品質の問題ではありません。チームはその前に停止ポイントが必要です
実行とレビュー後の記録: wh

o 何がチェックされたのか、なぜそれが保持されたのかを尋ねました
またはブロックされ、次に何が実行される可能性があるか。
コンテキストアンカー: EU AI 法 、NIST AI リスク管理フレームワーク 、および DORA 。これらはコンプライアンスの主張ではありません。
認証者は AI の意図を構造化された結果に変換し、それを次のように還元します。
決定、ゲートステータス、証明参照。
ポリシー、承認、証拠、許可された範囲、鮮度、再生、テナント、
とトークンの場合は、理由付きの 1 つの制限された決定を返します: accept 、Narrow 、
レビューするか、ブロックしてください。
要求可能な承認については、承認されたタスクがまだ一致しているかどうかを確認します。
実行前の現在のポリシーとマテリアルスコープのコンテキスト。
実際のサービスは、顧客所有のゲートを介してのみ実行されるべきです。
システム メタデータは、危険なアクションが形成されている場所を示すことができます。既存の API、ツール、
ジョブ、テレメトリ、イベント、ゲートウェイ ログは、
アクションの発見、ルール草案、入場決定、顧客ゲート、証明。
完全な結果パスマップを表示する
AIエージェント
-> 作戦を提案する
認証者
-> 認める / 絞り込む / レビュー / ブロック + 理由と証明の参照
顧客所有のゲート
-> 許可されている場合にのみ実際のサービスを呼び出します
顧客側のゲートがなければ、決定は強制ではなく証拠となります。と
その下流点、
終点となります。
アテスターをシャドウ パイロット モードで実行する
観察モードでは、エージェントがどのようなアクションを試行するか、それが危険である理由、および
どのポリシー、承認、証拠が存在するか。事前にリスクがわかる
実際のサービスが実行されます。
アテスターをシャドウ パイロット モードで実行する
オペレーション間で同じパターン
同じゲートを次の操作クラスの前に置くことができます。
パッケージバージョン: 0.3.0-評価版
リリースタグ: v0.3.0-評価版
リリース段階: 評価ベースライン
リリース タイプ: リポジトリ ベースライン / マルチパス ローカル レビュー
このベースラインは、ローカルでのレビューと統合計画用です。

アニング。ライブ顧客
導入と外部セキュリティ監査は別個の証明手順です。
アテスターはコントロール ポイントであり、データ レイクではありません。構造化されたリクエストが必要です
生の顧客データではなく、コンテキストと証拠の参照。お客様のシステムは
モデル、エージェント、ワークフロー、ウォレット、データベース、サービスコール、および記録システム。
ライトを開始します。詳細が必要な場合にのみ、さらに深く進みます。
初めての場合は、ローカル実行、シャドウ パイロット、カスタマー ゲートの順序に従ってください。
まず Attestor を試してください。最小のローカル返金パスを実行して、決定の証跡を確認してください。
Attestor をシャドウ パイロット モードで実行します。何かを強制する前に、1 つの実際のアクション パスを観察します。
Attestor を統合する方法 - 実際の副作用を見つけて、顧客所有のゲートを配置します。
リポジトリ ナビゲーター - ホスト、価格設定、サポート、証明、またはメンテナの作業に関する詳細なドキュメントを検索します。
使用境界: ライセンスと使用、およびセキュリティ ポリシー。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI Consequence Control Plane. Contribute to AI-gateway-systems/attestor development by creating an account on GitHub.

GitHub - AI-gateway-systems/attestor: AI Consequence Control Plane · GitHub
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
AI-gateway-systems
/
attestor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1,869 Commits 1,869 Commits .github .github .well-known .well-known docs docs examples examples fixtures fixtures ops ops scripts scripts specs specs src src tests tests vendor/ schematron/ 2026-CMS-QRDA-III vendor/ schematron/ 2026-CMS-QRDA-III .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.dr.yml docker-compose.dr.yml docker-compose.ha.yml docker-compose.ha.yml docker-compose.observability.yml docker-compose.observability.yml docker-compose.yml docker-compose.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Badges point to repository evidence.
How Attestor connects to existing systems
Control infrastructure for high-risk AI-driven operations.
Attestor sits between an AI-prepared operation and the system that would execute it.
Prompts can guide behavior, but they cannot enforce it or stop an unsafe,
unauthorized, or out-of-scope service call.
Unsafe requests can come from hallucination, stale context, poisoned tool output,
replay, missing approval, or hostile content. Before anything runs, Attestor checks
policy, authority, scope, freshness, replay, and evidence, then returns admit ,
narrow , review , or block .
The customer-owned gate decides before execution. The trail records what was proposed,
what was checked, and why it was held or allowed.
AI systems are moving from chat into tools that can touch payments, data,
access, customer messages, infrastructure, and programmable money.
That is no longer a prompt-quality problem. Teams need a stop point before
execution, and a record after review: who asked, what was checked, why it held
or blocked, and what may run next.
Context anchors: EU AI Act , NIST AI Risk Management Framework , and DORA . These are not compliance claims.
Attestor translates AI intent into a structured consequence, then reduces it to
a decision, gate status, and proof references.
It checks policy, approval, evidence, allowed scope, freshness, replay, tenant,
and token, then returns one bounded decision with reasons: admit , narrow ,
review , or block .
For requestable approvals, it checks that the approved task still matches the
current policy and material scope context before execution.
The real service should run only through the customer-owned gate.
System metadata can show where risky actions are forming. Existing APIs, tools,
jobs, telemetry, events, and gateway logs can become review material for
action discovery, rule drafts, admission decisions, customer gates, and proof.
View the full consequence path map
AI agent
-> proposes an operation
Attestor
-> admit / narrow / review / block + reasons and proof references
Customer-owned gate
-> calls the real service only when allowed
Without a customer-side gate, the decision is evidence, not enforcement. With
that downstream point ,
it becomes the stop point.
Run Attestor in shadow pilot mode
Observe mode shows what actions agents would try, why they may be risky, and
which policy, approval, and evidence are present. You see the risk before a
real service runs.
Run Attestor in shadow pilot mode
The Same Pattern Across Operations
The same gate can sit before these operation classes:
Package version: 0.3.0-evaluation
Release tag: v0.3.0-evaluation
Release stage: evaluation baseline
Release type: repository baseline / multi-path local review
This baseline is for local review and integration planning. Live customer
deployment and external security audit are separate proof steps.
Attestor is a control point, not a data lake. It needs structured request
context and proof references, not raw customer data. Customer systems keep the
model, agent, workflow, wallet, database, service call, and system of record.
Start light. Go deeper only when you need the detail.
If you are new, follow this order: local run , shadow pilot , then customer gate .
Try Attestor first - run the smallest local refund path and see the decision trail.
Run Attestor in shadow pilot mode - observe one real action path before enforcing anything.
How to integrate Attestor - find the real side effect and place the customer-owned gate.
Repository navigator - find deeper docs for hosted, pricing, support, proof, or maintainer work.
Use boundaries: License and use and Security Policy .
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
