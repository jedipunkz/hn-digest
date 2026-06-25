---
source: "https://www.cartai.ai/"
hn_url: "https://news.ycombinator.com/item?id=48677647"
title: "Show HN: CartAI – Checkout API for AI agents and apps"
article_title: "CartAI — The AI Agent That Handles Checkout"
author: "maniluppal"
captured_at: "2026-06-25T19:34:31Z"
capture_tool: "hn-digest"
hn_id: 48677647
score: 1
comments: 0
posted_at: "2026-06-25T18:49:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CartAI – Checkout API for AI agents and apps

- HN: [48677647](https://news.ycombinator.com/item?id=48677647)
- Source: [www.cartai.ai](https://www.cartai.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T18:49:12Z

## Translation

タイトル: HN を表示: CartAI – AI エージェントおよびアプリ用のチェックアウト API
記事のタイトル: CartAI — チェックアウトを処理する AI エージェント
説明: 任意の Web プロパティに移動してトランザクションを完了する特殊なエージェントをデプロイするための 1 つの API。チェックアウト、サブスクリプション、請求書、注文。構成可能。信頼性のある。開発者第一。

記事本文:
CartAI — チェックアウトを処理する AI エージェント
製品ソリューション ドキュメント ブログ サインイン API キーを取得 ライブ · Product Hunt で間もなく開始 チェックアウトを処理する AI エージェント。
1 つの API で、あらゆる Web プロパティに移動してトランザクション (チェックアウト、サブスクリプションのサインアップ、請求書の支払い、注文の送信など、トランザクションが発生するあらゆる場所) を完了する専用のエージェントをデプロイします。それを自動化しましょう。それを埋め込むか、これまでそれがなかった表面でそれを有効にします。構成可能。信頼性のある。開発者第一。
ドキュメントを読んでください checkout.ts webhook.ts // POST https://api.cartai.ai/checkout
const response = await fetch ( "https://api.cartai.ai/checkout" , {
メソッド: "POST" 、
ヘッダー: {
"Content-Type" : "application/json" ,
"認可" : ` Bearer ${ process.env.CARTAI_KEY } ` 、
}、
本文: JSON.stringify ({
顧客: {
電子メール: "buyer@example.com" 、
名前：「ジェーン・ドゥ」、
}、
タスク: [{
ターゲット: "bestbuy.com/site/sony-wh-1000xm5/6505727.p" ,
バリエーション: { カラー: "黒" } 、
送料: { プロフィール: "./buyer.json" } 、
支払い: { プロファイル: "./card.json" } 、
}]、
オプション: { サンドボックス: true } 、
})、
});
const { taskId, status } = 応答を待ちます。 json();
// → taskId: "task_abc123"、ステータス: "queued"
// 実行は非同期で実行されます。Webhook 経由で完了をリッスンします。 // Webhook リスナー — サブスクリプションのライフサイクル イベント
"express" からエクスプレスをインポートします。
const app = Express();
アプリ。 use (express.json());
アプリ。 post ( "/webhooks/cartai" , (req, res) => {
const {イベント、データ} = req.body;
スイッチ (イベント) {
ケース「開始済み」:
通知_顧客 (data.customerId);
休憩;
ケース「完了」:
send_post_purchase_flow (data.customerId);
休憩;
「失敗」の場合:
do_retry (data.customerId、data.executionId);
休憩;
}
解像度。 sendStatus (200);
} );
アプリ。聞いてください ( 3000 );サポートされるサーフェス 任意の Web プロパティ。ワークフローがナビゲートする必要がある場合は、Car

tAI はそれを介して取引できます。すべての e コマース チェックアウト カスタム カート イントラネット ポータル ベンダー ポータル ログインゲート ワークフロー カテゴリ 特殊な交差点にある特殊なレイヤー。
ブラウザの自動化とエージェントによる支払いはどちらも広大で重要なカテゴリです。どちらも、トランザクションを実際にクリアする必要があるワークフロー専用に構築されたものではありません。 CartAIです。
プログラム可能なトランザクションにより、信頼性の高い AI 主導のブラウザ実行と支払いインフラストラクチャが単一のプリミティブに融合されます。トランザクションをクリアする必要があるワークフロー専用に構築されており、トランザクションをクリアしないワークフロー用には構築されていません。範囲: 設計により特化されています。制約となるのは堀です。エージェント支払い エージェント向けの認証および支払いレール エージェントがユーザーに代わって安全に取引できるようにする ID、信頼、およびレール層 (KYA、署名済み資格情報、エージェント発行の支払い手段)。範囲: エージェントによる支払いを信頼できるものにするレール。支払いに移動できる実行レイヤーとペアになります。隣接するカテゴリは両方とも実際の問題を解決します。 CartAI は、トランザクションをクリアする必要があるという両者の間の問題を解決します。単一の、交渉の余地のない結果のために構築された専門的なプリミティブ。
誰でもナビゲートできます。トランザクションをクリアできる人はほとんどいません。
難しいのはクリックではありません。クリックしてから確認するまでのすべてです。私たちは 3 つの部分をプリミティブに組み込みました。正しくなければならない部分、そうでないとトランザクションが静かに失敗する部分です。
構成可能なワークフロー トランザクションが取り得るあらゆる形状に合わせて構築されています。すべてのフローには独自の形状があります。配送が必要なものもあれば、必要のないものもあります。 3DS が必要なものもあれば、必要ないものもあります。ログインが必要なものもあれば、必要のないものもあります。 CartAI ワークフローは、特定の状況で必要なステップを組み立て、常にトランザクションのどこかに到達します。
PCI 準拠の支払いスタック カードの取り込み、保管、およびエージェント

プロトコル — 正しく構築されています。完全に PCI に準拠したクレジット カードの取り込みと保管を、次世代のトランザクションで実行されるエージェント支払いプロトコルと統合します。
協調的なボット軽減 フラグは立てられず、クリアされています。私たちは、Cloudflare、Human、Fingerprint を回避しません。Web Bot Auth、Skyfire KYA、および署名されたエージェント ID を介してそれらと協力します。取引は、世界最大の販売者を保護するのと同じ信頼インフラストラクチャを通じてクリーンにクリアされます。
クラウドフレア / 人間 / 指紋
以下: 注文が確認されて終了したワークフロー。同じ API、同じ SDK、同じプリミティブ。実稼働マーチャントの表面に対して記録されます。演出やシミュレーションは行われません。
BestBuy Electronics をクリア Newegg Electronics をクリア Jomashop Luxury をクリア Ulta Beauty をクリア プラットフォーム 1 つのプリミティブ。製品は4つ。
チェックアウトは私たちが解決するのが難しい部分であり、他のすべてが配置されるくさびです。 4 つの製品は、製品の検索、注文の決済、資金の移動、利益の共有といったトランザクション ライフサイクル全体をカバーします。
01 カタログ 販売者全体で製品を検索し、バリエーションとライブ価格を取得し、カートが存在する前にチェックアウトの見積もりを返します。検索 · 詳細 · 見積もり カタログコアチェックアウトを探索する ライブマーチャントサーフェスで注文をクリアし、正規化された注文と Webhook を使用して確認まで追跡します。チェックアウト API · 注文 · Webhook チェックアウトを探索する 03 支払い Visa Intelligent Commerce および Mastercard Agent Pay を介してホストされた支払いセッション。 PCI スコープはスタックから離れたままになります。支払いセッション · 範囲外の PCI 支払いの探索 04 収益化 エージェントがあなたのサーフェスを通じて購入した場合でも、あなたは収益を得ることができます。帰属が組み込まれたアフィリエイト コミッション。 コミッション · 帰属 収益化の探索 署名、検証、許可されています。 CartAI は aro ではなく信頼インフラ内で動作します

それと。
さらに多くのものがオンラインで提供される 2 つの方法で構築できる API。乗り捨て用のカート。
完全な制御を必要とする開発者は、CartAI をワークフローに直接接続します。ターンキー チェックアウトが必要な Surface は、ホストされているカートに入れられ、午後に発送されます。どちらの方法でもその下にある同じプリミティブです。
完全なプログラム可能なトランザクション プリミティブ。支払い、注文、ワークフロー。トランザクションで終了するフローを作成します。
すべての状態遷移に対する Webhook
冪等、再試行安全な設計
これまで存在しなかったサーフェスのドロップイン チェックアウト。バリアントの選択、住所、支払い、確認 — を 1 行に埋め込みます。その下に同じプリミティブがあります。トランザクションは引き続き同じ方法でクリアされます。
バリアント、アドレス、支払い機能が組み込まれています
同じ API を基盤として利用されています
CartAI のすべてのユースケースは、手動のトランザクション プロセスを置き換える、トランザクションを新しいアプリケーションに埋め込む、トランザクションが存在しなかったサーフェスにトランザクションを追加する、という 3 つのモードのいずれかに対応します。
プログラム可能なトランザクションは現在運用中です。
API キーを取得してプラットフォームにアクセスし、最初のワークフローがクリアされるのを今すぐ確認してください。

## Original Extract

One API to deploy a specialized agent that navigates any web property and completes the transaction. Checkout, subscriptions, invoices, orders. Composable. Reliable. Developer-first.

CartAI — The AI Agent That Handles Checkout
Product Solutions Docs Blog Sign In Get API Key Live · Launching on Product Hunt soon The AI agent that handles checkout.
One API to deploy a specialized agent that navigates any web property and completes the transaction — checkout, subscription sign-up, invoice payment, order submission, anywhere a transaction has to happen. Automate it. Embed it, or enable it on surfaces that never had it. Composable. Reliable. Developer-first.
Read the docs checkout.ts webhook.ts // POST https://api.cartai.ai/checkout
const response = await fetch ( "https://api.cartai.ai/checkout" , {
method: "POST" ,
headers: {
"Content-Type" : "application/json" ,
"Authorization" : ` Bearer ${ process.env.CARTAI_KEY } ` ,
},
body: JSON.stringify ({
customer: {
email: "buyer@example.com" ,
name: "Jane Doe" ,
},
tasks: [{
target: "bestbuy.com/site/sony-wh-1000xm5/6505727.p" ,
variant: { color: "black" } ,
shipping: { profile: "./buyer.json" } ,
payment: { profile: "./card.json" } ,
}],
options: { sandbox: true } ,
}),
});
const { taskId, status } = await response. json ();
// → taskId: "task_abc123", status: "queued"
// Execution runs async — listen via webhook for completion. // Webhook listener — subscription lifecycle events
import express from "express" ;
const app = express ();
app. use (express. json ());
app. post ( "/webhooks/cartai" , (req, res) => {
const { event, data } = req.body;
switch (event) {
case "STARTED" :
notify_customer (data.customerId);
break ;
case "COMPLETED" :
send_post_purchase_flow (data.customerId);
break ;
case "FAILED" :
do_retry (data.customerId, data.executionId);
break ;
}
res. sendStatus ( 200 );
} );
app. listen ( 3000 ); Surfaces supported Any web property. If a workflow has to navigate it, CartAI can transact through it. All e-commerce checkouts Custom carts Intranet portals Vendor portals Login-gated workflows The category A specialized layer at a specialized intersection.
Browser automation and agentic payments are both vast, important categories. Neither is built specifically for the workflows where a transaction has to actually clear. CartAI is.
Programmable transactions, made reliable AI-driven browser execution and payment infrastructure fused into a single primitive. Built specifically for workflows where a transaction has to clear — and not built for workflows where one doesn't. Scope: specialized by design. The constraint is the moat. Agentic payments Authentication and payment rails for agents The identity, trust, and rails layer that lets agents transact safely on behalf of users — KYA, signed credentials, agent-issued payment instruments. Scope: the rails that make agentic payments trustable. Pairs with an execution layer that can navigate to the payment. Both adjacent categories solve real problems. CartAI solves the one between them: the transaction has to clear. A specialist primitive built for that single, non-negotiable outcome.
Anyone can navigate. Few can clear a transaction.
The hard part isn't the click. It's everything between the click and the confirmation. We engineered three pieces into the primitive — the ones that have to be right or the transaction silently fails.
Composable workflows Built for every shape a transaction can take. Every flow has its own shape — some need shipping, some don't; some need 3DS, some don't; some need login, some don't. CartAI workflows assemble the steps your specific situation requires, while always landing somewhere transactional.
PCI-compliant payment stack Card intake, vaulting, and agentic protocols — built right. Full PCI-compliant credit card intake and vaulting, integrated with the agentic payment protocols the next generation of transactions will run on.
Cooperative bot-mitigation Cleared, not flagged. We don't evade Cloudflare, Human, or Fingerprint — we cooperate with them via Web Bot Auth, Skyfire KYA, and signed agent identity. Transactions clear cleanly through the same trust infrastructure that protects the world's largest merchants.
Cloudflare / Human / Fingerprint
Below: workflows that ended in confirmed orders. Same API, same SDK, same primitive. Recorded against production merchant surfaces — not staged, not simulated.
Cleared BestBuy Electronics Cleared Newegg Electronics Cleared Jomashop Luxury Cleared Ulta Beauty The platform One primitive. Four products.
Checkout is the hard part we solve, and the wedge everything else sits on. Four products cover the full transaction lifecycle around it: find the product, clear the order, move the money, share the upside.
01 Catalog Search products across merchants, pull variants and live pricing, and return checkout estimates before any cart exists. Search · Details · Estimates Explore Catalog Core Checkouts Clear the order on the live merchant surface, then track it to confirmation with normalized orders and webhooks. Checkout API · Orders · Webhooks Explore Checkouts 03 Payments Hosted payment sessions via Visa Intelligent Commerce and Mastercard Agent Pay. PCI scope stays off your stack. Payment Sessions · PCI out of scope Explore Payments 04 Monetization When an agent buys through your surface, you still earn. Affiliate commissions with attribution built in. Commissions · Attribution Explore Monetization Signed, verified, allowed. CartAI operates inside the trust infrastructure, not around it.
+ more coming online Two ways to ship An API to build on. A cart to drop in.
Devs who want full control wire CartAI into their workflow directly. Surfaces that need a turnkey checkout drop in our hosted cart and ship in an afternoon. Same primitive underneath either way.
The full programmable-transaction primitive. Payments, Orders, Workflows. Compose any flow that ends in a transaction.
Webhooks for every state transition
Idempotent, retry-safe by design
Drop-in checkout for surfaces that never had one. Variant selection, addresses, payment, confirmation — embed it in one line. Same primitive underneath; the transaction still clears the same way.
Variants, addresses, payment built in
Powered by the same API underneath
Every CartAI use case maps to one of three modes — replace a manual transactional process, embed transactions into a new application, or add transactions to a surface that never had them.
Programmable transactions, in production today.
Get an API key, hit the platform, watch your first workflow clear today.
