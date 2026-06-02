---
source: "https://github.com/vercel/ai/issues/11280"
hn_url: "https://news.ycombinator.com/item?id=48364193"
title: "Vercel AI Gateway Appears to Block BYOK Requests When Account Balance Reaches $0"
article_title: "AI Gateway blocks all requests when credits are depleted, even with BYOK configured · Issue #11280 · vercel/ai · GitHub"
author: "kaicianflone"
captured_at: "2026-06-02T00:29:19Z"
capture_tool: "hn-digest"
hn_id: 48364193
score: 2
comments: 1
posted_at: "2026-06-02T00:07:24Z"
tags:
  - hacker-news
  - translated
---

# Vercel AI Gateway Appears to Block BYOK Requests When Account Balance Reaches $0

- HN: [48364193](https://news.ycombinator.com/item?id=48364193)
- Source: [github.com](https://github.com/vercel/ai/issues/11280)
- Score: 2
- Comments: 1
- Posted: 2026-06-02T00:07:24Z

## Translation

タイトル: Vercel AI ゲートウェイは、アカウント残高が 0 ドルに達すると BYOK リクエストをブロックするようです
記事のタイトル: AI Gateway は、BYOK が構成されている場合でも、クレジットが枯渇するとすべてのリクエストをブロックします · 問題 #11280 · vercel/ai · GitHub
説明: 説明 AI Gateway クレジットが枯渇すると (残高がゼロになると)、BYOK (Bring Your Own Key) 資格情報を使用する必要があるリクエストも含め、すべてのリクエストがブロックされます。これは、BYOK が次の方法で設定されているかどうかに関係なく発生します。 ダッシュボード統合...

記事本文:
AI ゲートウェイは、BYOK が構成されている場合でも、クレジットがなくなるとすべてのリクエストをブロックします · 問題 #11280 · vercel/ai · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ヴェルセル
/
あい
公共
通知
通知設定を変更するにはサインインする必要があります

ngs
追加のナビゲーション オプション
コード
AI ゲートウェイは、BYOK が構成されている場合でも、クレジットがなくなるとすべてのリクエストをブロックします #11280
リンクをコピー 新しい問題 リンクをコピー 開く Open AI Gateway は、BYOK が設定されている場合でも、クレジットがなくなるとすべてのリクエストをブロックします #11280 リンクをコピー 担当者 プロバイダー パッケージに関連する ai/プロバイダーにラベルを付けます。プロバイダー パッケージに関連する少なくとも 1 つの `provider/*` ラベルと一緒に割り当てる必要があります。少なくとも 1 つの `provider/*` ラベルと一緒に割り当てる必要があります バグ 文書化されたとおりに動作しない何か プロバイダー/ゲートウェイ 文書化されたとおりに動作しない何か @ai-sdk/gateway プロバイダーに関連する問題 @ai-sdk/gateway プロバイダーに関連する問題 説明
発行本体のアクション 説明
AI Gateway クレジットが枯渇すると (残高がゼロになると)、BYOK (Bring Your Own Key) 資格情報を使用する必要があるリクエストも含め、すべてのリクエストがブロックされます。これは、BYOK が次の方法で設定されているかどうかに関係なく発生します。
ダッシュボードの統合 (機能する「テストキー」検証を使用)
コード内のプログラムによる ProviderOptions.byok
BYOK 資格情報が (ダッシュボードまたは ProviderOptions.byok 経由で) 構成されている場合、リクエストは、AI Gateway クレジットを必要とせずに、それらの資格情報を使用してプロバイダーに直接ルーティングされる必要があります。
「外部 AI プロバイダーで独自の資格情報を使用すると、マークアップを追加せずに、AI Gateway がユーザーに代わってリクエストを認証できます。」
これは、BYOK がクレジット システムを完全にバイパスする必要があることを意味します。
資金が不足しています。 AI サービスを引き続き使用するには、アカウントにクレジットを追加してください。
次の場合でも:
BYOK 統合はダッシュボードと「テスト キー」パスで構成されます
ProviderOptions.byok はリクエスト内で明示的に設定されます
Vercel ダッシュボードで BYOK 統合を構成する (AI ゲートウェイ → 統合)
「テストキー」機能を使用してキーが機能することを確認します ✓
AI ゲートウェイを枯渇させる

クレジットは $0 まで
createGateway() を使用してリクエストを作成します。
import { createGateway ,generateText } from "ai" ;
const ゲートウェイ = createGateway() ;
// 方法 1: ダッシュボードの BYOK 構成に依存する
const response1 = await generatedText ( {
モデル: ゲートウェイ ( "anthropic/claude-3-5-haiku-latest" ) 、
プロンプト:「こんにちは」、
} ) ;
// 方法 2: ProviderOptions を使用した明示的な BYOK
const response2 = await generatedText ( {
モデル: ゲートウェイ ( "anthropic/claude-3-5-haiku-latest" ) 、
プロンプト:「こんにちは」、
プロバイダーオプション: {
ゲートウェイ: {
ビョク : {
anthropic:[{apiKey:プロセス .環境 。 ANTHROPIC_API_KEY } ] 、
} 、
} 、
} 、
} ) ;
どちらも「資金不足」エラーで失敗します。
API キーは有効です。プロバイダーの直接呼び出しは機能します。
"@ai-sdk/anthropic" から { anthropic } をインポートします。
"ai" からインポート {generateText} ;
// これは完璧に動作します
const 応答 = await generatedText ( {
モデル : anthropic ( "claude-3-5-haiku-latest" ) 、
プロンプト:「こんにちは」、
} ) ;
環境
aiパッケージバージョン：6.0.0-beta.128
Vercel クレジットコストを回避するために BYOK を設定する
利用可能なクォータを持つ有効なプロバイダー API キーを持っている
AI Gateway クレジットが不足する (偶発的であっても)
...AI ゲートウェイの使用が完全にブロックされ、BYOK の目的が無効になります。
次の場合には、信用調査をバイパスする必要があります。
有効な BYOK 資格情報が、要求されたプロバイダーのダッシュボードで構成されている
ProviderOptions.byok には、要求されたプロバイダーの有効な認証情報が含まれています
ゲートウェイは、Vercel のシステム資格情報を使用する場合にのみクレジットを必要とします。
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Description When AI Gateway credits are depleted (zero balance), all requests are blocked - including those that should use BYOK (Bring Your Own Key) credentials. This happens regardless of whether BYOK is configured via: Dashboard integ...

AI Gateway blocks all requests when credits are depleted, even with BYOK configured · Issue #11280 · vercel/ai · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
vercel
/
ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
AI Gateway blocks all requests when credits are depleted, even with BYOK configured #11280
Copy link New issue Copy link Open Open AI Gateway blocks all requests when credits are depleted, even with BYOK configured #11280 Copy link Assignees Labels ai/provider related to a provider package. Must be assigned together with at least one `provider/*` label related to a provider package. Must be assigned together with at least one `provider/*` label bug Something isn't working as documented Something isn't working as documented provider/gateway Issues related to the @ai-sdk/gateway provider Issues related to the @ai-sdk/gateway provider Description
Issue body actions Description
When AI Gateway credits are depleted (zero balance), all requests are blocked - including those that should use BYOK (Bring Your Own Key) credentials. This happens regardless of whether BYOK is configured via:
Dashboard integrations (with working "Test Key" verification)
Programmatic providerOptions.byok in code
When BYOK credentials are configured (either via dashboard or providerOptions.byok ), requests should be routed directly to the provider using those credentials, without requiring AI Gateway credits .
"Using your own credentials with an external AI provider allows AI Gateway to authenticate requests on your behalf with no added markup."
This implies BYOK should bypass the credit system entirely.
Insufficient funds. Please add credits to your account to continue using AI services.
Even when:
BYOK integrations are configured in dashboard AND "Test Key" passes
providerOptions.byok is explicitly set in the request
Configure BYOK integration in Vercel Dashboard (AI Gateway → Integrations)
Verify the key works using "Test Key" feature ✓
Deplete AI Gateway credits to $0
Make a request using createGateway() :
import { createGateway , generateText } from "ai" ;
const gateway = createGateway ( ) ;
// Method 1: Relying on dashboard BYOK config
const response1 = await generateText ( {
model : gateway ( "anthropic/claude-3-5-haiku-latest" ) ,
prompt : "Hello" ,
} ) ;
// Method 2: Explicit BYOK via providerOptions
const response2 = await generateText ( {
model : gateway ( "anthropic/claude-3-5-haiku-latest" ) ,
prompt : "Hello" ,
providerOptions : {
gateway : {
byok : {
anthropic : [ { apiKey : process . env . ANTHROPIC_API_KEY } ] ,
} ,
} ,
} ,
} ) ;
Both fail with "Insufficient funds" error.
The API keys are valid - direct provider calls work:
import { anthropic } from "@ai-sdk/anthropic" ;
import { generateText } from "ai" ;
// This works perfectly
const response = await generateText ( {
model : anthropic ( "claude-3-5-haiku-latest" ) ,
prompt : "Hello" ,
} ) ;
Environment
ai package version: 6.0.0-beta.128
Set up BYOK to avoid Vercel credit costs
Have valid provider API keys with available quota
Run out of AI Gateway credits (even accidentally)
...are completely blocked from using AI Gateway, defeating the purpose of BYOK.
The credit check should be bypassed when:
Valid BYOK credentials are configured in dashboard for the requested provider
providerOptions.byok contains valid credentials for the requested provider
The gateway should only require credits when using Vercel's system credentials.
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
