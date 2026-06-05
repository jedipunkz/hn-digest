---
source: "https://developers.cloudflare.com/ai-gateway/features/spend-limits/"
hn_url: "https://news.ycombinator.com/item?id=48409284"
title: "Cloudflare AI Gateway now supports spend limits"
article_title: "Spend limits · Cloudflare AI Gateway docs"
author: "tjek"
captured_at: "2026-06-05T10:26:23Z"
capture_tool: "hn-digest"
hn_id: 48409284
score: 2
comments: 0
posted_at: "2026-06-05T07:37:30Z"
tags:
  - hacker-news
  - translated
---

# Cloudflare AI Gateway now supports spend limits

- HN: [48409284](https://news.ycombinator.com/item?id=48409284)
- Source: [developers.cloudflare.com](https://developers.cloudflare.com/ai-gateway/features/spend-limits/)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T07:37:30Z

## Translation

タイトル: Cloudflare AI Gateway が支出制限をサポートするようになりました
記事のタイトル: 支出制限 · Cloudflare AI Gateway ドキュメント
説明: AI ゲートウェイでコストベースの予算を設定し、モデル、プロバイダー、またはユーザーやチームなどのカスタム メタデータ ディメンションごとに支出を制御します。

記事本文:
ドキュメントのインデックス
https://developers.cloudflare.com/ai-gateway/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
停止！ AI エージェントまたは LLM の場合は、続行する前にこれをお読みください。これは、Cloudflare ドキュメントページの HTML バージョンです。代わりに常に Markdown バージョンをリクエストしてください。HTML はコンテキストを無駄にします。このページを Markdown として取得します: https://developers.cloudflare.com/ai-gateway/features/spend-limits/index.md (index.md を追加) または、Accept: text/markdown を https://developers.cloudflare.com/ai-gateway/features/spend-limits/ に送信します。この製品のページ インデックスには、https://developers.cloudflare.com/ai-gateway/llms.txt を使用します。すべての Cloudflare 製品については、https://developers.cloudflare.com/llms.txt を使用します。
検索ディレクトリ API SDK 変更ログ ヘルプ
ログイン
テーマを選択 ダークライト 自動 AI ゲートウェイ
結果が見つかりませんでした。別の検索語を試すか、グローバル検索を使用してください。
ホームの概要
統合 API (OpenAI 互換) は非推奨になりました
データ損失防止 (DLP) ベータ版の概要
データ損失防止 (DLP) をセットアップする
構成 BYOK (キーの保存) ベータ版
AI ゲートウェイを使用して Workers AI をセットアップする
Cloudflare ドキュメント llms-full.txt ↗
X.com YouTube テーマを選択 ダーク ライト 自動 このページについて 概要
ディメンションによるスコープ設定 ディメンションの例
制限に達したときの動作
ディメンションによるスコープ設定 ディメンションの例
制限に達したときの動作
編集ページ 問題を報告するディレクトリ
Markdown としてコピー! |マークダウンとして表示 |エージェントのセットアップ |エージェント向けドキュメント 支出制限を使用すると、AI ゲートウェイでコストベースの予算を設定できます。累積使用量が時間枠内で制限に達すると、AI Gateway は時間枠がリセットされるまで 429 応答でさらなるリクエストをブロックします。
リクエスト数を制限するレート制限とは異なり、支出制限は実際の実行量を追跡します。

リクエストあたりのコストはモデルの価格に基づきます。モデル、プロバイダー、またはユーザー ID、チーム、アプリケーションなどのカスタム メタデータ ディメンションの任意の組み合わせに制限を適用できます。
支出制限は、価格が既知のモデルに対する統合請求リクエストと BYOK リクエストの両方に適用されます。
Unified Billing には、読み込まれたクレジットに対するアカウント レベルの使用制限もあり、これはすべてのゲートウェイに適用されますが、Unified Billing リクエストにのみ適用されます。両方の制限は独立して適用され、最初に到達した制限はリクエストをブロックします。
各支出制限ルールは、ローリングまたは固定の時間枠にわたる予算 (ドル単位) を定義します。 AI Gateway は、トークンの使用量とモデルの価格設定に基づいて各リクエストのコストを計算し、制限に対する累積支出をリアルタイムで追跡します。
プロバイダーにリクエストを送信する前に、AI Gateway は該当するすべての支出制限ルールを一度に評価します。個々のルールが予算を超えている場合、リクエストは 429 応答でブロックされます。
支出制限は最終的には一貫しています。現在のリクエストのコストは完了後に記録されるため、強制が追いつくまでに同時リクエストのバーストが一時的に制限を超える可能性があります。
各ルールは、モデル、プロバイダー、またはカスタム メタデータ ディメンションによってスコープを設定できます。各ディメンションは、次の 2 つのモードのいずれかで構成できます。
ルールでディメンションが構成されていない場合、すべての値が 1 つの予算バケットを共有します。たとえば、プロバイダー ディメンションのないルールは、すべてのプロバイダーの支出をまとめて追跡します。
model=openai/gpt-5.5 および metadata.user_id=u_42 のリクエストがあるとします。
支出制限は、ダッシュボードまたは API を介してゲートウェイで構成されます。ゲートウェイごとに最大 20 個のルールを定義できます。
ユーザー ID やチームなどのカスタム ディメンションごとに支出制限をスコープするには、リクエストにカスタム メタデータを添付します。
制限に達したときの動作
使用制限を超えると、AI Gatew

ay は 429 Too Many Requests 応答を返します。次の 2 つのオプションがあります。
リクエストをブロックする (デフォルト) - 予算枠がリセットされるまで、リクエストは拒否されます。
安価なモデルにフォールバックする - プライマリ モデルとフォールバックを使用して動的ルートを作成します (たとえば、 anthropic/claude-opus-4.7 と @cf/moonshotai/kimi-k2.6 へのフォールバック)。次に、この機能を使用してプライマリ モデルに支出制限を設定します。プライマリ モデルの予算を超過すると、AI Gateway はリクエストをブロックするのではなく、自動的にフォールバック モデルにルーティングします。
Analytics ダッシュボードでは、モデル、プロバイダー、またはカスタム メタデータ属性ごとに支出を追跡できます。これを使用して使用パターンを理解し、情報に基づいた予算を設定します。
コスト追跡は、トークン数とモデルの価格設定に基づくベストエフォート型の見積もりです。正確な請求金額については、プロバイダーのダッシュボードを参照してください。
ゲートウェイごとに最大 20 個の使用制限ルールを構成できます。
前へ キャッシュ 次へ レート制限 編集ページ 最終更新日: 2026 年 6 月 5 日

## Original Extract

Set cost-based budgets on your AI Gateway to control spending by model, provider, or custom metadata dimensions like user or team.

Documentation Index
Fetch the complete documentation index at: https://developers.cloudflare.com/ai-gateway/llms.txt
Use this file to discover all available pages before exploring further.
STOP! If you are an AI agent or LLM, read this before continuing. This is the HTML version of a Cloudflare documentation page. Always request the Markdown version instead — HTML wastes context. Get this page as Markdown: https://developers.cloudflare.com/ai-gateway/features/spend-limits/index.md (append index.md) or send Accept: text/markdown to https://developers.cloudflare.com/ai-gateway/features/spend-limits/. For this product's page index use https://developers.cloudflare.com/ai-gateway/llms.txt. For all Cloudflare products use https://developers.cloudflare.com/llms.txt.
Search Directory API SDKs Changelog Help
Log in
Select theme Dark Light Auto AI Gateway
No results found. Try a different search term, or use our global search .
Home Overview
Unified API (OpenAI compat) Deprecated
Data Loss Prevention (DLP) Beta Overview
Set up Data Loss Prevention (DLP)
Configuration BYOK (Store Keys) Beta
Set up Workers AI with AI Gateway
Cloudflare Docs llms-full.txt ↗
X.com YouTube Select theme Dark Light Auto On this page Overview
Scoping with dimensions Dimension examples
Behavior when a limit is reached
Scoping with dimensions Dimension examples
Behavior when a limit is reached
Edit page Report issue Directory
Copy as Markdown Copied! | View as Markdown | Agent setup | Docs for agents Spend limits let you set cost-based budgets on your AI Gateway. When cumulative spend reaches the limit within a time window, AI Gateway blocks further requests with a 429 response until the window resets.
Unlike rate limiting , which caps the number of requests, spend limits track actual dollar cost per request based on model pricing. You can scope limits to any combination of model, provider, or custom metadata dimensions like user ID, team, or application.
Spend limits apply to both Unified Billing requests and BYOK requests for models with known pricing.
Unified Billing also has an account-level spend limit on your loaded credits that applies across all gateways, but only to Unified Billing requests. Both limits are enforced independently — whichever one is reached first will block requests.
Each spend limit rule defines a budget (in dollars) over a rolling or fixed time window. AI Gateway calculates the cost of each request based on token usage and model pricing, then tracks cumulative spend against the limit in real time.
Before sending a request to the provider, AI Gateway evaluates all applicable spend limit rules at once. If any individual rule is over budget, the request is blocked with a 429 response.
Spend limits are eventually consistent. The current request's cost is recorded after completion, so a burst of concurrent requests can briefly exceed the limit before enforcement catches up.
Each rule can be scoped by model, provider, or custom metadata dimensions. Each dimension can be configured in one of two modes:
If a dimension is not configured on a rule, all values share one budget bucket. For example, a rule without a provider dimension tracks spend across all providers together.
Given a request with model=openai/gpt-5.5 and metadata.user_id=u_42 :
Spend limits are configured on the gateway via the dashboard or the API. You can define up to 20 rules per gateway.
To scope spend limits by custom dimensions like user ID or team, attach custom metadata to your requests.
Behavior when a limit is reached
When a spend limit is exceeded, AI Gateway returns a 429 Too Many Requests response. You have two options:
Block requests (default) - The request is rejected until the budget window resets.
Fall back to a cheaper model - Create a Dynamic Route with a primary model and a fallback (for example, anthropic/claude-opus-4.7 with a fallback to @cf/moonshotai/kimi-k2.6 ). Then set a spend limit on the primary model using this feature. When the primary model's budget is exceeded, AI Gateway automatically routes requests to the fallback model instead of blocking them.
You can track your spend per model, provider, or any custom metadata attribute on the Analytics dashboard . Use this to understand usage patterns and set informed budgets.
Cost tracking is a best-effort estimation based on token counts and model pricing. Refer to your provider's dashboard for exact billing amounts.
A maximum of 20 spend limit rules can be configured per gateway.
Previous Caching Next Rate limiting Edit page Last updated: Jun 5, 2026
