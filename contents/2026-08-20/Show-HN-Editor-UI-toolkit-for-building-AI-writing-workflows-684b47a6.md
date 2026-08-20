---
source: "https://imperavi.com/redactor/ai-assistant/"
hn_url: "https://news.ycombinator.com/item?id=49372193"
title: "Show HN: Editor UI toolkit for building AI writing workflows"
article_title: "AI Assistant for Redactor — Build AI Writing Workflows"
image: "https://imperavi.com/assets/img/opengraph/redactor-1200x630.jpg"
author: "lessio"
captured_at: "2026-08-20T09:23:54Z"
capture_tool: "hn-digest"
hn_id: 49372193
score: 1
comments: 0
posted_at: "2026-08-20T09:09:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Editor UI toolkit for building AI writing workflows

- HN: [49372193](https://news.ycombinator.com/item?id=49372193)
- Source: [imperavi.com](https://imperavi.com/redactor/ai-assistant/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T09:09:31Z

## Translation

タイトル: Show HN: AI ライティング ワークフローを構築するためのエディター UI ツールキット
記事のタイトル: Redactor の AI アシスタント — AI ライティング ワークフローの構築
説明: AI の提案、レビュー、レポート、作成ツールを Redactor に追加するための、開発者が制御する UI および統合レイヤー。独自のバックエンド、プロンプト、モデルを使用します。

記事本文:
モデルに依存しない · サーバー側プロンプト · 構造化ドキュメントコンテキスト · 同期、ジョブ、またはカスタムトランスポート
製品内に AI 支援のライティング ワークフローを構築する
Redactor は、ドキュメント エディターとレビュー UI を提供します。独自のバックエンド、プロンプト、モデルを接続します。
モデルはデータを返します。 Redactor はそれを編集ワークフローに変えます。
Redactor は、選択されたブロックまたは完全なドキュメントを構造化コンテキストとして送信します。
エンドポイントは、プロンプト、モデル、制限、ログ、およびデータ ポリシーを選択します。
応答は、レビュー、レポート、代替案のリスト、または適用されたドキュメントの変更になります。
Redactor は特定の AI プロバイダーを必要とせず、トークンごとの課金も追加しません。
エディターはアクション ID とドキュメント コンテキストを送信します。バックエンドは完全なプロンプトを選択し、モデルを呼び出し、選択したモードで予期される応答形状を返します。
リダクター('#entry', {
プラグイン: ['アシスタント']、
あい: {
URL: '/api/ai/text',
モデル: 「あなたのモデル」
}
});
最低限のお願い
{
"アクション": "短く"、
"モード": "レビュー",
"コンテキスト": {
"スコープ": "ブロック",
「ブロック」: [
{
"uid": "ブロック-3",
"タイプ": "テキスト",
"タグ": "p",
"コンテンツ": "..."
}
]
}、
"lang": "en",
"モデル": "あなたのモデル"
}
応答モード
1 つの統合、複数の編集面
製品に必要なアクションを構築する
独自のコンテンツ、ユーザー、編集ルールに基づいたアクションを定義します。各アクションでは、選択したブロックまたはドキュメント全体を使用し、追加のコンテキストをバックエンドに送信し、結果をレビュー、適用された変更、またはレポートとして表示できます。
組み込みアクションは実用的な例を提供しますが、メニュー、プロンプト、モデル、応答処理、およびアプリケーション ロジックはすべてプロジェクトに適合させることができます。
エディターはアクション ID と構造化ドキュメント コンテキストをバックエンドに送信します。リクエストがどのように処理され、何が返されるかはアプリケーションによって決まります。
プロンプトを維持する

サーバー上のテンプレート
各アクションのモデルまたは処理パイプラインを選択します
ユーザー、テナント、ドキュメント、またはアプリケーション コンテキストを追加する
独自の認証と認可を適用する
レート制限と使用制限を設定する
ログ記録とデータ保持を制御する
レビューやレポート用に構造化された回答をフォーマットする
アシスタントは特定のモデル プロバイダーに関連付けられていません。バックエンドは、期待される応答形式を返す限り、OpenAI、Claude、ローカル モデル、またはカスタム サービスを呼び出すことができます。
イベントを使用すると、ホスト アプリケーションは重要な各段階でアシスタントのワークフローを検査および変更できます。
送信前にリクエストを変更またはキャンセルする
カスタムフィールドとHTTPヘッダーを追加する
成功したリクエストとエラーを観察する
提案が承認または拒否されたときに反応する
挿入前にコンテンツを検査する
イメージ生成応答を処理する
非同期ジョブの進行状況を観察する
アシスタントのアクティビティをアプリケーション分析または監査ログに接続する
コンテンツまたはリクエストを変更できるフックは、対応するアクションの前に実行されます。ライフサイクル イベントは、アクションの完了後に結果を提供します。
アシスタントのリクエストがアプリケーションに到達する方法を選択します。すべてのトランスポートは同じアクション コンテキストを使用し、同じレビューまたはレポート応答形式を返します。
HTTP リクエストを送信し、同じレスポンスで完了した結果を返します。
非同期ジョブを作成し、結果が準備できるまでアプリケーションをポーリングします。
組み込みの HTTP トランスポートを使用する代わりに、JavaScript リクエスト ハンドラーを提供します。
アシスタントは、小規模なリクエストとレスポンスのコントラクトを使用します。これらのガイドでは、バックエンド、プロンプト テンプレート、非同期処理、レポート出力について説明します。
構造化ドキュメント ブロックを受け取り、レビュー、変更、レポート用の JSON を返すサーバー側プロンプト テンプレートを定義します。
OpenAI のバックエンド例を含むリクエストとレスポンスのフォーマットを参照してください。

クロード、テキストアクション、画像生成。
長いドキュメント、キューに入れられた処理、および低速なモデル用に作成およびポーリング プロトコルを実装します。
付属の HTML フレームワークを使用して、概要、チェック項目の作成、キー ポイント、その他の構造化された結果をレポート タブに表示します。
構成オプション、アクション モード、応答形式、トランスポート、イベント、拡張ポイントを確認します。
アシスタントのリファレンスを開く →
Redactor Core 用アドオン · コアは含まれません
AI アシスタント、レビュー ワークフロー、レポート、構成可能なアクション、トランスポート オプション、完全なソース コードが含まれています。
独自のバックエンドとモデルプロバイダーを使用します。 AI リクエストとトークンの使用は含まれません。
購入したバージョンのライセンスは永久です。 Assistant Suite のアップデートは、現在の Redactor Core メジャー バージョンでは無料です。 Core の新しいメジャー バージョンに移行する場合は、有料のアシスタント アップグレードが必要になる場合があります。
統合とライセンスに関する質問
アシスタントがバックエンドに接続する方法、ライセンスに含まれるもの、アプリケーションの制御下に残るもの。
いいえ。独自のバックエンドとモデル プロバイダーに接続します。
いいえ。これはアドオンであり、別のコア ライセンスが必要です。
いいえ。プロバイダーの使用量は、選択したモデルとインフラストラクチャに応じて請求されます。
はい。アクションは ID によって識別され、サーバー側のプロンプト テンプレートまたはアプリケーション ロジックにマッピングできます。
はい。この統合は、特定のプロバイダー SDK ではなく、Redactor のリクエストとレスポンスの契約に基づいています。
ジョブ トランスポートは、長いドキュメントとキューに入れられた処理に対する非同期の作成とポーリングのワークフローをサポートします。
はい。レビュー アクションでは、元のコンテンツと提案されたコンテンツが表示され、作成者が各変更を承認または拒否できます。

## Original Extract

A developer-controlled UI and integration layer for adding AI suggestions, reviews, reports, and writing tools to Redactor. Use your own backend, prompts, and model.

Model-agnostic · Server-side prompts · Structured document context · Sync, jobs, or custom transport
Build AI-assisted writing workflows inside your product
Redactor provides the document editor and review UI. You connect your own backend, prompts, and model.
The model returns data. Redactor turns it into an editing workflow.
Redactor sends the selected blocks or the full document as structured context.
Your endpoint chooses the prompt, model, limits, logging, and data policy.
The response becomes a review, report, list of alternatives, or an applied document change.
Redactor does not require a specific AI provider and does not add per-token billing.
The editor sends an action ID and document context. Your backend selects the full prompt, calls the model, and returns the response shape expected by the selected mode.
Redactor('#entry', {
plugins: ['assistant'],
ai: {
url: '/api/ai/text',
model: 'your-model'
}
});
Minimal request
{
"action": "shorter",
"mode": "review",
"context": {
"scope": "blocks",
"blocks": [
{
"uid": "block-3",
"type": "text",
"tag": "p",
"content": "..."
}
]
},
"lang": "en",
"model": "your-model"
}
Response modes
One integration, several editing surfaces
Build the actions your product needs
Define actions around your own content, users, and editorial rules. Each action can use selected blocks or the full document, send additional context to your backend, and present the result as a review, an applied change, or a report.
Built-in actions provide working examples, but the menu, prompts, models, response handling, and application logic can all be adapted to your project.
The editor sends an action ID and structured document context to your backend. Your application decides how that request is processed and what is returned.
Keep prompt templates on the server
Choose the model or processing pipeline for each action
Add user, tenant, document, or application context
Apply your own authentication and authorization
Set rate limits and usage limits
Control logging and data retention
Format structured responses for reviews and reports
The Assistant is not tied to a particular model provider. Your backend can call OpenAI, Claude, a local model, or a custom service as long as it returns the expected response format.
Events let the host application inspect and modify the Assistant workflow at each important stage.
Modify or cancel a request before it is sent
Add custom fields and HTTP headers
Observe successful requests and errors
React when suggestions are accepted or rejected
Inspect content before insertion
Handle image-generation responses
Observe asynchronous job progress
Connect Assistant activity to application analytics or audit logs
Hooks that can change content or requests run before the corresponding action. Lifecycle events provide the result after the action has completed.
Choose how Assistant requests reach your application. All transports use the same action context and return the same review or report response formats.
Send an HTTP request and return the completed result in the same response.
Create an asynchronous job and poll your application until the result is ready.
Provide a JavaScript request handler instead of using the built-in HTTP transports.
The Assistant uses a small request and response contract. These guides cover the backend, prompt templates, asynchronous processing, and report output.
Define server-side prompt templates that receive structured document blocks and return JSON for reviews, changes, and reports.
See request and response formats with backend examples for OpenAI, Claude, text actions, and image generation.
Implement the create-and-poll protocol for long documents, queued processing, and slower models.
Use the included HTML framework to present summaries, writing checks, key points, and other structured results in report tabs.
Review configuration options, action modes, response formats, transports, events, and extension points.
Open the Assistant reference →
Add-on for Redactor Core · Core not included
Includes the AI Assistant, review workflows, reports, configurable actions, transport options, and full source code.
Uses your own backend and model provider. AI requests and token usage are not included.
Licenses are perpetual for the purchased versions. Assistant Suite updates are free for the current Redactor Core major version. A paid Assistant upgrade may be required when moving to a new major version of Core.
Questions about integration and licensing
How the Assistant connects to your backend, what the license includes, and what remains under your application’s control.
No. You connect it to your own backend and model provider.
No. It is an add-on and requires a separate Core license.
No. Provider usage is billed according to the model and infrastructure you choose.
Yes. Actions are identified by IDs and can be mapped to server-side prompt templates or application logic.
Yes. The integration is based on Redactor’s request and response contract, not a specific provider SDK.
The jobs transport supports asynchronous create-and-poll workflows for long documents and queued processing.
Yes. Review actions present the original and suggested content and let the author accept or reject each change.
