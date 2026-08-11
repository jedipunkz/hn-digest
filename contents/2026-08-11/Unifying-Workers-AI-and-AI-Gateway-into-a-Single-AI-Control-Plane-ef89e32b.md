---
source: "https://blog.cloudflare.com/workers-ai-gateway-unification/"
hn_url: "https://news.ycombinator.com/item?id=49251541"
title: "Unifying Workers AI and AI Gateway into a Single AI Control Plane"
article_title: "Unifying Workers AI and AI Gateway into a single AI control plane | Cloudflare Blog"
author: "c_f_"
captured_at: "2026-08-11T00:57:32Z"
capture_tool: "hn-digest"
hn_id: 49251541
score: 1
comments: 0
posted_at: "2026-08-11T00:01:46Z"
tags:
  - hacker-news
  - translated
---

# Unifying Workers AI and AI Gateway into a Single AI Control Plane

- HN: [49251541](https://news.ycombinator.com/item?id=49251541)
- Source: [blog.cloudflare.com](https://blog.cloudflare.com/workers-ai-gateway-unification/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T00:01:46Z

## Translation

タイトル: Workers AI と AI ゲートウェイを単一の AI コントロール プレーンに統合
記事のタイトル: Workers AI と AI ゲートウェイを単一の AI コントロール プレーンに統合 |クラウドフレアのブログ
説明: Cloudflare は、AI Gateway と Workers AI を単一のコントロール プレーンに統合し、開発者にマネージド GPU と外部プロバイダーの両方にわたる可観測性、課金、および動的ルーティングを提供します。統合バインディングとモデルファースト ルーティングにより、回復力のある AI アプリケーションの構築がどのように簡素化されるかを学びます。

記事本文:
Workers AI と AI Gateway を単一の AI コントロール プレーンに統合 |クラウドフレアのブログ
コンテンツへスキップ すべてのカテゴリ AI
ログイン 営業担当者へのお問い合わせ ブログ エージェント エージェントウィーク AI +5 さらに 5 個のタグを表示 8 個のタグ 8 個のタグを表示 選択したタグ
エージェント エージェントウィーク AI AI ゲートウェイ開発者 プラットフォーム開発者 製品ニュース 労働者 AI
プラットフォームの自動最適化
Cloudflare 1 ユーザーのリスクスコア
Workers AI と AI ゲートウェイを単一の AI コントロール プレーンに統合
AI Gateway と Workers AI は当初、別個の製品としてスタートしましたが、時間が経つにつれて、ユーザーが集中していることに気づきました。 AI ゲートウェイを使用すると、リクエストを任意のモデル プロバイダーにプロキシし、組み込みの可観測性、ロギング、アクセス、セキュリティを取得できます。 Workers AI では、管理する GPU インフラストラクチャ上でモデルをホストし、サービスとしての推論にアクセスするために利用できる API エンドポイントを公開します。
これらの製品のアーキテクチャは異なって見えますが、エンド ユーザーにとっては、洗練されたコントロール プレーンを備えたモデルに接続するという同じ目標を達成します。本日、これらの製品を 1 つの統合パスに統合することで、単一のコントロール プレーンから可観測性、課金、セキュリティ、ロギングなどを管理しながら、任意のモデル プロバイダー (Workers AI を含む) に接続できるようにする方法に関する計画を共有できることを嬉しく思います。
これは、私たちが計画しているいくつかの大きな計画に向けた次のステップです。モデル ルーティングの将来にとって、統合コントロール プレーンが何を意味するのかを学び続けてください。
私たちは、これらの製品がエントリーポイントである Workers バインディングと REST API を通じてより統合されつつあるという事実をほのめかしてきました。 AI Gateway と Workers AI を呼び出すために使用できる AI バインディングがあります。 AI ゲートウェイとワーカー AI バインディングを個別に分けるという概念はなく、すべて同じパスを通過します。私たちは「デフォルト」ゲートウェイのアイデアを数か月前に発表しました。

以前に AI ゲートウェイをセットアップしたことがあっても、AI ゲートウェイの可観測性とログ記録を自動的に継承できます。もちろん、アプリケーションを複数のプロジェクトに分割したい場合は、独自のゲートウェイを指定することもできます。
AI ゲートウェイ経由で Workers AI を呼び出している場合のバインディング呼び出しは次のようになります。
デフォルトのエクスポート {
非同期フェッチ ( request , env ) {
const 応答 = 環境を待機します。 AI 。走って（
'@cf/zai-org/glm-5.2' ,
{
メッセージ: [
{ 役割: 'ユーザー' 、内容: 'フランスの首都はどこですか?' }、
】
}、
{
ゲートウェイ: {
id: 'default' , // 組み込みゲートウェイには 'default' を使用します
}、
}
);
新しい応答を返します ( JSON . stringify (応答), {
ヘッダー: { 'Content-Type' : 'application/json' },
});
}、
};また、単一の統合 REST API、/ai/ エンドポイントも発表しました。これにより、AI Gateway 経由で Workers AI に対して同様の呼び出しを行うことができます。
カール "https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/run/@cf/zai-org/glm-5.2" \
- H "認可: ベアラー {api_token}" \
- H "コンテンツ タイプ: application/json" \
- H "cf-aig-gateway-id: デフォルト" \
- d'{
"messages": [{"role": "user", "content": "フランスの首都はどこですか?"}],
これにより、AI Gateway と Workers AI へのエントリ ポイントを統合できるため、最初にどの製品を使用するかを選択する必要はありません。すべての製品にはバッテリーが付属しています。
すべての Workers AI ユーザーに対する自動監視と制御
この統合の最も直接的なメリットの 1 つは、推論トラフィックの可視化を開始する前に AI ゲートウェイを明示的に作成する必要がなくなったことです。これまでにゲートウェイを設定したことがない場合は、バインディングまたは REST API 呼び出しでゲートウェイ ID としてデフォルトを渡すだけで、最初の認証されたリクエストで AI Gateway によってゲートウェイが自動的に作成されます。
これにより、すべてのリクエストが完全な状態でログに記録されます。

リクエストとレスポンスのペイロード、トークン数がモデルごとに追跡され、ダッシュボードを設定しなくてもコストの帰属が得られます。後でデフォルト ゲートウェイを超えた場合、つまりカスタム キャッシュ ルールが必要な場合、またはアプリケーションごとにトラフィックを分割する場合は、名前付きゲートウェイを作成し、パラメータを 1 回変更するだけでリクエストをそのゲートウェイに向けることができます。
バインディングに装着するとこんな感じになります。以前は、Workers AI に直接呼び出しました。
const 応答 = 環境を待機します。 AI 。 run ( '@cf/zai-org/glm-5.2' , {
メッセージ: [{ 役割: 'ユーザー' 、コンテンツ: 'こんにちは!' }]、
});次に、AI ゲートウェイを介してルーティングし、完全な可観測性を得るために 3 番目の引数を追加します。
const 応答 = 環境を待機します。 AI 。走って（
'@cf/zai-org/glm-5.2' ,
{ メッセージ: [{ 役割: 'ユーザー' 、コンテンツ: 'こんにちは!' }] }、
{gateway: { id: 'default' } } // 最初の使用時にゲートウェイを自動作成します
); Cloudflare AI Gatewayダッシュボードにアクセスすると、レイテンシの内訳、トークンの使用状況、エラー率、正確なプロンプトと応答など、すべてのリクエストが表示されます。モデルの動作をデバッグしたり、AI 出力を監査したりするチームにとって、これは目隠し飛行からの大きなアップグレードです。
新機能: Workers AI に AI Gateway クレジットを使用する
本日開始する新しい機能は、Workers AI に AI Gateway クレジットを使用できる機能です。以前は、外部モデル プロバイダー (OpenAI、Anthropic など) でのみ AI Gateway クレジットを使用できましたが、AI Gateway クレジットを Workers AI の使用に適用することはまだできませんでした。ついに、システムで Workers AI の統合請求を許可できるようになりました。これは、クレジットでいっぱいのウォレットをロードし、それを OpenAI、Anthropic、Workers AI、または当社がサポートする任意のプロバイダーで使用することを選択できることを意味します。
現在、Workers AI の前払い請求を提供しており、ユーザーにこの新しいパスの使用を奨励したいため、AI Gateway uni を使用している場合は Workers AI モデルのレート制限の引き上げも提供しています。

請求を行った。レート制限に関する最新情報と、より高いレート制限をリクエストする方法については、開発者ドキュメントを参照してください。
近日公開予定: モデルファースト ルーティング
すべての推論トラフィックが単一のコントロール プレーンを通過するため、管理する必要があるプロバイダーではなく、必要なモデルから各リクエストに対応する方法について、より賢明な決定を下すことができます。プロバイダーファーストのルーティングでは、インフラストラクチャについて考える必要があります。「どのプロバイダーに電話すればよいですか? プロバイダーがダウンしたらどうすればよいですか?」モデルファーストのルーティングはそれを逆転させます。必要なもの (有能な推論モデル、高速サマライザー、安価な埋め込みモデル) を考えると、コントロール プレーンがプロバイダーの選択、フェイルオーバー、負荷分散を処理します。
現在、モデルを呼び出したい場合は、どのプロバイダーがそれをホストしているかを知る必要があります。そのプロバイダーがダウンしていたり​​、レート制限がかかっている場合、アプリケーションは壊れます。私たちは、モデルを指定すれば、残りは AI Gateway が処理する世界に向かって進んでいます。
こうすることで、Workers AI、Moonshot 独自の API、または同じ重みをホストする別のプロバイダーからのものであるかどうかを気にせずに、Kimi K2.7 コードをリクエストできます。 Workers AI に能力がある場合は、マネージド インフラストラクチャの恩恵を受けることができます。 Workers AI が限界に達した場合、ゲートウェイは同じモデルを提供できる別のプロバイダーに透過的に負荷分散します。必要に応じて 1 つのプロバイダーに固執することもできますが、復元力を重視する場合は、モデルファースト ルーティングを使用すると、より柔軟な対応が可能になります。当社は精査されたプロバイダーと協力しているため、モデル出力の品質は引き続き最優先であり、ゼロ データ保持 (ZDR) などの要件も尊重できます。
カール - X POST "https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1/chat/completions" \
- H "認可: ベアラー {api_token}" \
-H"c

f-aig-ゲートウェイ ID: 私のゲートウェイ" \
- H "コンテンツ タイプ: application/json" \
- d'{
"モデル": "kimi-k2.7-コード",
"messages": [{"role": "user", "content": "この機能を確認する"}]
これは、デフォルトで復元力が向上することも意味します。あるプロバイダーのモデルのバージョンに問題がある場合、アプリケーション レベルの再試行やワーカーの複雑なフォールバック ロジックを必要とせずに、トラフィックは別のプロバイダーに移行します。ゲートウェイは、モデルの可用性をルーティングの問題として扱います。今後数か月以内に、すべての AI Gateway と Workers AI ユーザーを対象にこれを試験的に実施したいと考えています。
ルーティングの次の進化は、単純なフェイルオーバーを超えています。私たちは、ユーザーが何を求めているかを理解し、設定を必要とせずにジョブに適したモデルを選択するインテリジェントなルーティングを構築しています。
モデルを指定する代わりに、ゲートウェイに決定させることができます。内部では、Workers AI で実行されている分類器がプロンプトを読み取り、それがどのような種類のタスク (コーディング、調査、要約、一般的な Q&A) であるか、タスクがどの程度複雑か、コンテキストがどの程度重要かを予測します。次に、ヒューリスティック スコアラーが、厳選されたプールから最適なモデルにそれをマッピングします。コントロールしたいチームの場合は、正確なモデルを指定することもできます。他のすべての人にとって、構成ゼロのパスは、独自のルーティング ロジックを維持することなく、より優れた経済性とパフォーマンスが得られることを意味します。現在、社内でこれを試験運用しており、リリースまでの数週間で積極的にテストと反復を行う予定です。
すでに Workers AI を使用している場合、これを試す最も簡単な方法は、デフォルト ゲートウェイを介して既存の通話のルーティングを開始することです。モデルの呼び出し方法については何も変更せずに、リクエストのログ記録、トークンの追跡、コストの帰属をすぐに取得できます。
すでに AI Gateway を使用している場合、Workers AI をミックスに追加するのは、Workers AI モデルを呼び出すのと同じくらい簡単です。 AI ゲートウェイをロードします。

allet を使用すると、サポートされているすべてのプロバイダーで統一された請求が得られるほか、Workers AI モデルのレート制限も引き上げられます。
最初のゲートウェイをセットアップし、Workers AI モデル カタログを参照し、今すぐ構築を開始してください。
新しい投稿の通知を受け取るために購読する
あなたのメールアドレスを共有することはありません。
ご購読いただきありがとうございます!受信箱を確認して確認してください。
マルチテナントプラットフォームの開発
プライバシーの選択 セキュリティの問題を報告する |プライバシーポリシー |利用規約 | GDPR |商標検索は一時的に利用できません。製品
新しいタブで開きます 新しいタブで開きます 新しいタブで開きます すべてのカテゴリ AI

## Original Extract

Cloudflare is unifying AI Gateway and Workers AI into a single control plane, giving developers observability, billing, and dynamic routing across both managed GPUs and external providers. Learn how unified bindings and model-first routing simplify building resilient AI applications.

Unifying Workers AI and AI Gateway into a single AI control plane | Cloudflare Blog
Skip to content All Categories AI
Login Contact Sales Blog Agents Agents Week AI +5 Show 5 more tags 8 Tags Show 8 tags Selected Tags
Agents Agents Week AI AI Gateway Developer Platform Developers Product News Workers AI
Automatic Platform Optimization
Cloudflare One User Risk Score
Unifying Workers AI and AI Gateway into a single AI control plane
AI Gateway and Workers AI first started as distinct products, but over time, we noticed our users were converging. With AI Gateway, you can proxy requests to any model provider and get built-in observability, logging, access, and security. On Workers AI, we host models on the GPU infrastructure that we manage, exposing an API endpoint you can leverage to access inference-as-a-service.
The architecture of these products looks different, but to an end user, it achieves the same goal: connecting you to models with a sophisticated control plane. Today, we're excited to share our plans on how these products converge into one unified path, so you can connect to any model provider (including Workers AI), while managing things like observability, billing, security, and logging from a single control plane.
It’s the next step toward some big plans we have — read on to learn what a unified control plane means for the future of model routing.
We've been hinting at the fact that these products are becoming more unified through our entrypoints: the Workers binding and the REST API. We have an AI binding that you can use to call AI Gateway and Workers AI. There's no concept of a separate AI Gateway and Workers AI binding: it all goes through the same path. We shipped the idea of a “default” gateway a few months ago, so that if you have never set up an AI Gateway before, you could still automatically inherit the AI Gateway observability and logging. Of course, you can still specify your own gateway if you'd like to split up applications into multiple projects.
Here's what the binding call looks like, if you are calling Workers AI via AI Gateway:
export default {
async fetch ( request , env ) {
const response = await env. AI . run (
'@cf/zai-org/glm-5.2' ,
{
messages: [
{ role: 'user' , content: 'What is the capital of France?' },
]
},
{
gateway: {
id: 'default' , // Use 'default' for the built-in gateway
},
}
);
return new Response ( JSON . stringify (response), {
headers: { 'Content-Type' : 'application/json' },
});
},
}; We’ve also announced a single unified REST API — the /ai/ endpoint that allows you to make similar calls to Workers AI via AI Gateway.
curl "https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/run/@cf/zai-org/glm-5.2" \
- H "Authorization: Bearer {api_token}" \
- H "Content-Type: application/json" \
- H "cf-aig-gateway-id: default" \
- d '{
"messages": [{"role": "user", "content": "What is the capital of France?"}],
}' Doing this allows us to unify the entrypoints to AI Gateway and Workers AI, so you don't need to make a choice between which product to use first: it all comes with batteries included.
Automatic observability and control for all Workers AI users
One of the most immediate benefits of this convergence is that you no longer need to explicitly create an AI Gateway before you start getting visibility into your inference traffic. If you've never set up a gateway before, just pass default as the gateway ID in your binding or REST API calls, and AI Gateway will create it automatically on the first authenticated request.
With this, every request is logged with full request and response payloads, token counts are tracked per model, and you get cost attribution without any dashboard setup. If you later outgrow the default gateway — if you want custom caching rules or to split traffic by application — you can create a named gateway and point your requests at it with a single parameter change.
Here's how it looks in the binding. Before, you called Workers AI directly:
const response = await env. AI . run ( '@cf/zai-org/glm-5.2' , {
messages: [{ role: 'user' , content: 'Hello!' }],
}); Now, add a third argument to route through AI Gateway and get full observability:
const response = await env. AI . run (
'@cf/zai-org/glm-5.2' ,
{ messages: [{ role: 'user' , content: 'Hello!' }] },
{ gateway: { id: 'default' } } // Auto-creates the gateway on first use
); Head to the Cloudflare AI Gateway dashboard and you'll see every request: latency breakdowns, token usage, error rates, and the exact prompts and responses. For teams debugging model behavior or auditing AI output, this is a huge upgrade from flying blind.
New: using AI Gateway credits for Workers AI
A new thing we're launching today is the ability to use AI Gateway credits for Workers AI. Before, you could only use AI Gateway credits on external model providers (e.g., OpenAI, Anthropic) but you couldn't apply your AI Gateway credits to Workers AI usage just yet. We've finally enabled our systems to allow unified billing for Workers AI. This means that you can load a wallet full of credits, and then choose to spend that across OpenAI, Anthropic, Workers AI, or any provider that we support.
Since we're now offering pre-paid billing for Workers AI and want to encourage users to use this new path, we're also offering elevated rate limits on Workers AI models if you use AI Gateway unified billing. Please refer to the developer docs for up-to-date information regarding rate limits, as well as how to request a higher rate limit.
Coming soon: model-first routing
With all your inference traffic flowing through a single control plane, we can start making smarter decisions about how to serve each request starting with the model you want, not the provider you have to manage. Provider-first routing forces you to think about infrastructure: "Which provider do I call? What if they're down?" Model-first routing flips that. You think about what you need — a capable reasoning model, a fast summarizer, a cheap embedding model — and the control plane handles provider selection, failover, and load balancing.
Today, if you want to call a model, you have to know which provider hosts it. If that provider is down or rate-limiting you, your application breaks. We're moving toward a world where you specify the model, and AI Gateway handles the rest.
This way, you can request Kimi K2.7 Code and not care whether it comes from Workers AI, Moonshot's own API, or another provider that hosts the same weights. If Workers AI has capacity, you get the benefit of our managed infrastructure. If Workers AI is at capacity, the gateway transparently load balances you to another provider that can serve the same model. You can still choose to stick to a single provider if you’d like, but model-first routing enables you to get more flexibility if you care about resiliency. We work with vetted providers, so the quality of model outputs remains top priority, and will also be able to respect requirements such as Zero Data Retention (ZDR).
curl - X POST "https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1/chat/completions" \
- H "Authorization: Bearer {api_token}" \
- H "cf-aig-gateway-id: my-gateway" \
- H "Content-Type: application/json" \
- d '{
"model": "kimi-k2.7-code",
"messages": [{"role": "user", "content": "Review this function"}]
}' This also means better resiliency by default. If one provider's version of a model is having issues, traffic shifts to another without application-level retries or complex fallback logic in your Workers. The gateway treats model availability as a routing problem. We hope to pilot this in the coming months for all AI Gateway and Workers AI users.
The next evolution of routing goes beyond simple failover. We're building intelligent routing that understands what you're asking for and picks the right model for the job without any configuration required.
Instead of specifying a model, you can let the gateway decide. Under the hood, a classifier running on Workers AI reads your prompt and predicts what kind of task it is (coding, research, summarization, general Q&A), how complex it is, and how much context matters. A heuristic scorer then maps that to the best model from a curated pool. For teams that want control, you can still specify exact models. For everyone else, the zero-config path means you get better economics and performance without maintaining your own routing logic. We are currently piloting this internally, and we’ll be actively testing and iterating in the next few weeks before release.
If you're already using Workers AI, the easiest way to try this out is to start routing your existing calls through a default gateway. You'll immediately get request logging, token tracking, and cost attribution without changing anything else about how you call the models.
If you're already using AI Gateway, adding Workers AI to the mix is as simple as calling a Workers AI model. Load up your AI Gateway wallet and you'll get unified billing across every provider we support, plus elevated rate limits on Workers AI models.
Set up your first gateway , browse the Workers AI model catalog , and get started building today.
Subscribe to receive notifications of new posts
We’ll never share your email address.
Thanks for subscribing! Check your inbox to confirm.
Multi-Tenant Platform Development
Your privacy choices Report security issues | Privacy Policy | Terms of use | GDPR | Trademark Search is temporarily unavailable. Products
opens in a new tab opens in a new tab opens in a new tab All Categories AI
