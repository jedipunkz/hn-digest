---
source: "https://github.com/team-telnyx/telnyx-code-examples/tree/main/build-conversational-workflow-nodejs"
hn_url: "https://news.ycombinator.com/item?id=48690730"
title: "Building Voice AI Workflows with Branches Instead of One Giant Prompt"
article_title: "telnyx-code-examples/build-conversational-workflow-nodejs at main · team-telnyx/telnyx-code-examples · GitHub"
author: "anushathukral"
captured_at: "2026-06-26T19:32:12Z"
capture_tool: "hn-digest"
hn_id: 48690730
score: 2
comments: 0
posted_at: "2026-06-26T19:13:40Z"
tags:
  - hacker-news
  - translated
---

# Building Voice AI Workflows with Branches Instead of One Giant Prompt

- HN: [48690730](https://news.ycombinator.com/item?id=48690730)
- Source: [github.com](https://github.com/team-telnyx/telnyx-code-examples/tree/main/build-conversational-workflow-nodejs)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T19:13:40Z

## Translation

タイトル: 1 つの巨大なプロンプトではなく分岐を使用した音声 AI ワークフローの構築
記事のタイトル: メインの telnyx-code-examples/build-conversational-workflow-nodejs · チーム-telnyx/telnyx-code-examples · GitHub
説明: Telnyx AI 通信インフラストラクチャの本番対応コード例 — 音声 AI、SMS、SIP、IoT API - メインの telnyx-code-examples/build-conversational-workflow-nodejs · Team-telnyx/telnyx-code-examples

記事本文:
メインの telnyx-code-examples/build-conversational-workflow-nodejs · チーム Telnyx/telnyx-code-examples · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
チームテルニクス
/
telnyx コードの例
公共
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ビルド-会話型ワークフロー-nodejs
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
ビルド-会話型ワークフロー-nodejs
コピーパスのトップフォルダーとファイル
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md package.json package.json server.js server.js validate-workflow.js validate-workflow.js workflow.json workflow.json すべてのファイルを表示 README.md
概要
名前
会話型ワークフローの構築
タイトル
会話型ワークフローを構築する
説明
構造化されたブランチ、バックエンド ツール、および優先フォローアップを使用して、インバウンドの自動車保険請求受付のための Telnyx 会話ワークフローを構築します。
言語
ノードjs
フレームワーク
特急
テルニクス製品
音声AI
チャンネル
声
会話型ワークフローを構築する
自動車保険の損失受付の最初の通知を受信するための Telnyx 会話ワークフローを構築します。ワークフローは構造化された質問をし、緊急のケースに分岐し、バックエンド ツールを呼び出し、呼び出し元にクレームまたはフォールバック参照を返します。
この例は、Telnyx ポータルで会話ワークフローとして構成されています。ローカル Express アプリは、Telnyx がワークフロー ツール ノードから呼び出す Webhook ツール エンドポイントを公開します。
クレーム取り込みツールの作成 : POST /tools/create-claim-intake
ログ フォールバック ツール: POST /tools/log-claim-intake-fallback
フラグ優先度フォローアップツール: POST /tools/flag-priority-follow-up
PSTN 着信通話
│
▼
┌───────────────┐
│ Telnyx 会話型 │
│ ワークフロー │
━───────┬─────────┘
│ ノード + ブランチ
▼
┌───────────────┐
│オートクレー

m 吸気流量 │
│ 安全、発信者、損失フィールド │
━───────┬─────────┘
│ ワークフローツール呼び出し
▼
┌───────────────┐
│ エクスプレスツールサーバー │
│ /tools/create-claim-intake │
━───────┬─────────┘
│ JSON結果
▼
クレーム参照またはフォールバック参照
ワークフローのライフサイクル
発信者は、ワークフローに割り当てられた Telnyx 電話番号に到達します。
ワークフローにより、これが新しい自動請求であることが確認されます。
このワークフローは、怪我、危険、または危険な車両の状態をチェックします。
ワークフローは、発信者、保険契約、車両、およびインシデントの詳細を収集します。
ワークフローは、必須フィールドが存在する場合に create_claim_intake を呼び出します。
取り込みを完了できない場合、ワークフローは log_claim_intake_fallback を呼び出します。
priority_flag が true の場合、ワークフローは flag_priority_follow_up を呼び出します。
Telnyx は、音声、AI、SIP、メッセージング、およびプログラム可能な通信を 1 つのプライベートなグローバル ネットワーク上に配置する AI 通信インフラストラクチャ プラットフォームです。会話型ワークフローを使用すると、目に見える分岐やツール呼び出しを使用して構造化された音声エクスペリエンスを構築できるため、保険金請求の受付などの運用プロセスは、単一の大きなプロンプトよりも監査が容易になります。
この例で保険を使用する理由
保険の損失の最初の通知は、現実的なビジネス上の制約があるため、有用なワークフロー パターンです。
ワークフローは特定のフィールドを収集する必要があります
緊急の場合は早めの分岐が必要
バックエンド レコードは、必要な情報が存在した後にのみ作成する必要があります。
音声エクスペリエンスは、補償、支払い、修理、賠償責任、法的アドバイス、または医学的アドバイスを約束するものであってはなりません
同じ構造をスケジューリング、サポートトリアージ、ファイル管理に再利用できます。

広告資格、患者の受け入れ、サービスの派遣。
.env.example を .env にコピーし、次のように入力します。
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/build-conversational-workflow-nodejs
cp .env.example .env # CLAIM_TOOL_SECRET を入力します
npmインストール
npm test
ノードserver.js # http://localhost:8787 で開始します
Workflow Configuration
Telnyx ポータルで、「自動請求取り込み会話ワークフロー」という名前の会話ワークフローを作成します。
GUIDE.md からノードとブランチを構築します。
ngrok ベース URL を使用してツール ノードを構成します。
POST https://<id>.ngrok.io/tools/create-claim-intake
POST https://<id>.ngrok.io/tools/log-claim-intake-fallback
POST https://<id>.ngrok.io/tools/flag-priority-follow-up
各ツールに共有ヘッダーを追加します。
受信 Telnyx 電話番号をワークフローに割り当てて電話をかけます。
POST /tools/create-claim-intake
ワークフローが必須フィールドを収集したときに、疑似請求取り込みを作成します。
curl -X POST http://localhost:8787/tools/create-claim-intake \
-H " コンテンツタイプ: application/json " \
-H " x-tool-secret: dev-secret " \
-d ' {
"caller_name": "ジェーン サンプル",
"発信者電話": "+15551234567",
"loss_type": "auto",
"損失日": "2026-06-15",
"loss_location": "ミッション ストリートと 5 番ストリート、サンフランシスコ",
"loss_description": "停止中に追突されました",
"priority_flag": false,
「続行への同意」: true
} '
応答:
{
"success" : true ,
"claim_intake_id" : " aci_123 " ,
"優先フラグ" : false 、
"next_step" : " クレームチームのフォローアップ "
}
POST /tools/log-claim-intake-fallback
発信者が間違った方向に誘導された場合、情報が不完全である場合、または主要なツールが使用できない場合に、フォールバックを記録します。
POST /tools/フラグ-優先順位-フォローアップ
緊急の場合に模擬優先度フォローアップ タスクを作成します。
監視用のヘルスチェックエンドポイント。
curl http://local

ホスト:8787/健康
応答:
{
"ステータス" : " OK " 、
"サービス" : "claim_intake_tools"
}
トラブルシューティング
問題
原因
修正
401 {"成功":false、"エラー":"不正"}
x-tool-secret が見つからないか、 CLAIM_TOOL_SECRET と一致しません。
同じ x-tool-secret ヘッダーを各 Telnyx ワークフロー ツール ノードに追加し、 .env を編集した後にサーバーを再起動します。
422 {"エラー":"必須フィールドがありません"}
ワークフローは、必須フィールドが収集される前にツールを呼び出しました。
workflow.json を確認し、最小限のフィールド チェックを行った後でのみ create_claim_intake を呼び出します。
ツール呼び出しがサーバーに到達しない
ローカルサーバーはパブリックにアクセスできません。
ngrok http 8787 を実行し、各ワークフロー ツール ノードで HTTPS URL を使用します。
発信者はサポートされていない Promise を受け取ります
ワークフローの指示またはノードのコピーが広すぎます。
ワークフローを摂取に重点を置き、補償、支払い、賠償責任、修理、法的、または医学的決定を回避します。
編集中にワークフロー グラフが中断される
ノードは、欠落しているノードまたはツールを指します。
npm test を実行して workflow.json を検証します。
関連する例
AI Insurance Claims Intake Voice (Python) - 保険金請求受付の音声エージェント
音声 AI エージェント (Node.js) を構築する - Telnyx 推論を使用して通話に応答し、応答を生成します
AI エージェントへの電話のルーティング (Node.js) - 着信通話を AI エージェントにルーティングします
AI アシスタント マルチ ツール (Python) - AI アシスタントのツール呼び出しパターン
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Production-ready code examples for Telnyx AI Communications Infrastructure — Voice AI, SMS, SIP, and IoT APIs - telnyx-code-examples/build-conversational-workflow-nodejs at main · team-telnyx/telnyx-code-examples

telnyx-code-examples/build-conversational-workflow-nodejs at main · team-telnyx/telnyx-code-examples · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
team-telnyx
/
telnyx-code-examples
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
build-conversational-workflow-nodejs
More options Directory actions
History History main Breadcrumbs
build-conversational-workflow-nodejs
Copy path Top Folders and files
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md package.json package.json server.js server.js validate-workflow.js validate-workflow.js workflow.json workflow.json View all files README.md
Outline
name
build-conversational-workflow
title
Build a Conversational Workflow
description
Build a Telnyx Conversational Workflow for inbound auto insurance claim intake with structured branches, backend tools, and priority follow-up.
language
nodejs
framework
express
telnyx_products
Voice AI
channel
voice
Build a Conversational Workflow
Build a Telnyx Conversational Workflow for inbound auto insurance first notice of loss intake. The workflow asks structured questions, branches for urgent cases, calls backend tools, and returns a claim or fallback reference to the caller.
This example is configured in the Telnyx Portal as a Conversational Workflow. The local Express app exposes webhook tool endpoints that Telnyx calls from workflow tool nodes:
Create Claim Intake Tool : POST /tools/create-claim-intake
Log Fallback Tool : POST /tools/log-claim-intake-fallback
Flag Priority Follow-Up Tool : POST /tools/flag-priority-follow-up
Inbound PSTN call
│
▼
┌────────────────────────────┐
│ Telnyx Conversational │
│ Workflow │
└─────────────┬──────────────┘
│ nodes + branches
▼
┌────────────────────────────┐
│ Auto claim intake flow │
│ safety, caller, loss fields │
└─────────────┬──────────────┘
│ workflow tool calls
▼
┌────────────────────────────┐
│ Express tool server │
│ /tools/create-claim-intake │
└─────────────┬──────────────┘
│ JSON result
▼
claim reference or fallback reference
Workflow lifecycle
Caller reaches the Telnyx phone number assigned to the workflow.
The workflow confirms this is a new auto claim.
The workflow checks for injury, danger, or unsafe vehicle status.
The workflow collects caller, policy, vehicle, and incident details.
The workflow calls create_claim_intake when required fields are present.
The workflow calls log_claim_intake_fallback when the intake cannot be completed.
The workflow calls flag_priority_follow_up when priority_flag is true.
Telnyx is an AI Communications Infrastructure platform that puts voice, AI, SIP, messaging, and programmable communications on one private, global network. Conversational Workflows let you build structured voice experiences with visible branches and tool calls, so operational processes like insurance claim intake are easier to audit than a single large prompt.
Why This Example Uses Insurance
Insurance first notice of loss is a useful workflow pattern because it has realistic business constraints:
the workflow must collect specific fields
urgent cases need early branching
backend records should only be created after required information is present
the voice experience must not promise coverage, payment, repairs, liability, legal advice, or medical advice
The same structure can be reused for scheduling, support triage, lead qualification, patient intake, and service dispatch.
Copy .env.example to .env and fill in:
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/build-conversational-workflow-nodejs
cp .env.example .env # fill in CLAIM_TOOL_SECRET
npm install
npm test
node server.js # starts on http://localhost:8787
Workflow Configuration
In the Telnyx Portal , create a Conversational Workflow named auto claim intake conversational workflow .
Build the nodes and branches from GUIDE.md .
Configure tool nodes with your ngrok base URL:
POST https://<id>.ngrok.io/tools/create-claim-intake
POST https://<id>.ngrok.io/tools/log-claim-intake-fallback
POST https://<id>.ngrok.io/tools/flag-priority-follow-up
Add the shared header to each tool:
Assign an inbound Telnyx phone number to the workflow and call it.
POST /tools/create-claim-intake
Creates a mock claim intake when the workflow has collected required fields.
curl -X POST http://localhost:8787/tools/create-claim-intake \
-H " content-type: application/json " \
-H " x-tool-secret: dev-secret " \
-d ' {
"caller_name": "jane sample",
"caller_phone": "+15551234567",
"loss_type": "auto",
"loss_date": "2026-06-15",
"loss_location": "mission street and 5th street, san francisco",
"loss_description": "rear-ended while stopped",
"priority_flag": false,
"consent_to_continue": true
} '
Response:
{
"success" : true ,
"claim_intake_id" : " aci_123 " ,
"priority_flag" : false ,
"next_step" : " claims team follow-up "
}
POST /tools/log-claim-intake-fallback
Records a fallback when the caller is misdirected, information is incomplete, or the primary tool cannot be used.
POST /tools/flag-priority-follow-up
Creates a mock priority follow-up task for urgent cases.
Health check endpoint for monitoring.
curl http://localhost:8787/health
Response:
{
"status" : " ok " ,
"service" : " claim_intake_tools "
}
Troubleshooting
Issue
Cause
Fix
401 {"success":false,"error":"unauthorized"}
x-tool-secret is missing or does not match CLAIM_TOOL_SECRET .
Add the same x-tool-secret header to each Telnyx workflow tool node and restart the server after editing .env .
422 {"error":"missing_required_fields"}
The workflow called a tool before required fields were collected.
Check workflow.json and only call create_claim_intake after the minimum field check.
Tool call never reaches the server
Local server is not publicly reachable.
Run ngrok http 8787 and use the HTTPS URL in each workflow tool node.
Caller receives an unsupported promise
Workflow instructions or node copy are too broad.
Keep the workflow focused on intake and avoid coverage, payment, liability, repair, legal, or medical decisions.
Workflow graph breaks during editing
A node points to a missing node or tool.
Run npm test to validate workflow.json .
Related Examples
AI Insurance Claims Intake Voice (Python) - voice agent for insurance claim intake
Build a Voice AI Agent (Node.js) - answer calls and generate replies with Telnyx Inference
Route Phone Calls to AI Agent (Node.js) - route inbound calls to an AI agent
AI Assistant Multi Tool (Python) - tool-calling pattern for AI assistants
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
