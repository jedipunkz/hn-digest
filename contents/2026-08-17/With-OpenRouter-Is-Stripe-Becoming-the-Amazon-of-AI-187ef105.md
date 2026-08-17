---
source: "https://github.com/getlago/lago/wiki/With-OpenRouter,-is-Stripe-becoming-the-Amazon-of-AI"
hn_url: "https://news.ycombinator.com/item?id=49324994"
title: "With OpenRouter, Is Stripe Becoming the Amazon of AI"
article_title: "With OpenRouter, is Stripe becoming the Amazon of AI · getlago/lago Wiki · GitHub"
author: "AnhTho_FR"
captured_at: "2026-08-17T00:39:16Z"
capture_tool: "hn-digest"
hn_id: 49324994
score: 2
comments: 2
posted_at: "2026-08-16T23:47:45Z"
tags:
  - hacker-news
  - translated
---

# With OpenRouter, Is Stripe Becoming the Amazon of AI

- HN: [49324994](https://news.ycombinator.com/item?id=49324994)
- Source: [github.com](https://github.com/getlago/lago/wiki/With-OpenRouter,-is-Stripe-becoming-the-Amazon-of-AI)
- Score: 2
- Comments: 2
- Posted: 2026-08-16T23:47:45Z

## Translation

タイトル: OpenRouter により、Stripe は AI のアマゾンになるのか
記事のタイトル: OpenRouter により、Stripe は AI のアマゾンになる · getlago/lago Wiki · GitHub
説明: オープンソースのメーターリングと使用量ベースの請求 API ⭐️ 消費量の追跡、サブスクリプション管理、価格設定の反復、支払いオーケストレーションと収益分析 - OpenRouter を使用すると、Stripe は AI の Amazon になります · getlago/lago Wiki

記事本文:
OpenRouter により、Stripe は AI のアマゾンになるのか · getlago/lago Wiki · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ゲトラゴ
/
ラゴ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
OpenRouter により、Stripe は AI のアマゾンになるのか
OpenRouter により、Stripe は AI のアマゾンになるのでしょうか?
Stripe は OpenRouter を 70 億ドル以上で買収することに合意したと伝えられている。ストライプは契約を確認していない。ブルームバーグはfinaとして報じた

とTechCrunchも同じ数字を報告した。
OpenRouter が API プロキシである場合、価格は法外に見えます。 Stripe が AI 推論のための Amazon マーケットプレイスの始まりを見ているとすれば、それはより理にかなっています。
OpenRouter は、開発者に 400 を超えるモデルに対して 1 つのアカウントと 1 つの API を提供します。その背後では、70 を超える推論プロバイダーがリクエストに対応するために競合しています。 OpenRouterによると、現在800万人のユーザーがいるという。
リクエストを転送するコードには 70 億ドルの価値はありません。大規模で増え続けるリクエストのプールをどこに送信するかを決定する権利がある可能性があります。
ゲートウェイはマーケットプレイスになる可能性がある
OpenRouter は通常、ゲートウェイとして説明されます。これは技術的でかなり中立的に聞こえます。リクエストが到着し、ゲートウェイがそれを転送し、応答が返されます。
ただし、OpenRouter は、どのサプライヤーが需要にアクセスできるかを決定します。
デフォルトのルーティング ロジックでは、まず、最近障害が発生したプロバイダーを削除します。次に、価格の逆二乗を使用して、残りの候補の中からより安価なプロバイダーを優先します。 OpenRouter 自身の例では、100 万トークンあたり 1 ドルを請求するエンドポイントは、3 ドルを請求するエンドポイントよりも最初のリクエストを受信する可能性が 9 倍高くなります。
それは開発者にとって良いことです。価格が下がり、フォールバック容量が得られ、維持する必要のあるプロバイダー統合が少なくなります。
推論プロバイダーにとって、これは分散システムです。価格とパフォーマンスによって、受信するトラフィックの量が決まります。プロバイダーは、価格を引き下げたり、信頼性を向上させたりすることで、より多くの需要を獲得できます。開発者が切り替えを明確に決定しなければ、需要が失われる可能性があります。
オンラインで商品を出品するのが困難だったため、Amazon マーケットプレイスは強力になりませんでした。購入者を 1 か所に集め、販売者が購入者に連絡する方法を制御しました。検索ランキングと購入ボックスは、販売者の商品と同じくらい重要である可能性があります。
OpenRouter はまだ存在しません。しかし、それはすでに認めています

er、パフォーマンスを測定し、需要をそれらの間で割り当てます。デフォルト ルートを受け入れる開発者が増えるにつれて、そのルーティング ルールがプロバイダーの収益を決定するようになります。
OpenRouterは顧客がクレジットを購入する際に5.5％の手数料を請求し、基礎となるプロバイダーのトークン価格を値上げするものではないとしている。目に見える料金はシンプルです。さらに興味深い資産は、その背後にある需要です。
これが Stripe にとって 70 億ドルの価値がある理由
Stripe はすでに OpenRouter をよく知っています。 OpenRouter は、支払い、請求、税金、不正行為の防止に Stripe を使用します。 1月には、両社はStripe Projectsを通じてOpenRouterも利用できるようにしたため、開発者やコーディングエージェントはアカウントをプロビジョニングし、StripeのコマンドラインからAPIキーを受け取ることができるようになった。 Stripe はここでパートナーシップについて説明し、OpenRouter はここでプロジェクトの統合について説明しました。
OpenRouter を所有すると、トランザクションの早い段階で Stripe が移動することになります。
現在、Stripe は顧客の支払いを確認できるようになりました。 OpenRouter は、どのモデルがリクエストされたか、どのプロバイダーがそれを提供したか、その推論コストを確認できます。この 2 つを組み合わせると、Stripe は AI 機能の作成コストとその販売から得た収益を結び付けることができます。
AI 製品の利益率は異常に変動するため、これは便利です。 2 つの同様の顧客アクションは、モデル、コンテキストの長さ、キャッシュ動作、再試行、選択したプロバイダーに応じて、コストが大きく異なる可能性があります。企業はそのコストを転嫁したり、マークアップを追加したり、クレジットを使い果たしたり、サブスクリプション内で吸収したりする可能性があります。
Stripe は、メーター、クレジット、請求、税金、不正行為の管理、プロバイダー決済など、そのアクティビティに関する支払い処理以上のものを販売できます。 OpenRouter は、コストが発生した瞬間にそれらの製品を配布する場所を提供します。
Stripe はすでにそのスタックを組み立てています。 Metronomeを約10億ドルで買収した

lion は、大容量使用量の測定と複雑な料金設定についてさらに詳しく説明します。 Stripe Billing を再建せずに Stripe を買収した理由について書きました。 Metronome は現在、Stripe 製品として販売されていますが、Stripe Billing と並んで独自のアプリケーションと製品面を持っています。不足しているレイヤーを購入する方が、再構築するよりも高速でした。 2 つの製品を 1 つの製品のように感じさせるのは遅くなります。
報道されたPayPal入札も背景にある。ロイター通信によると、StripeとAdventはPayPalに530億ドル以上を提示したという。 PayPal は販売者と消費者をもたらすでしょう。 OpenRouter は開発者と推論プロバイダーをもたらします。どちらも、複数の当事者が資金の移動、調整、収益化を必要とするネットワークです。
Stripe は 70 億ドルをはるかに下回るコストで AI ルーターを構築できる可能性があります。 OpenRouter のプロバイダー関係をすぐに再現したり、800 万人のユーザーに運用トラフィックを新しいプロバイダーに移動するよう説得したりすることはできませんでした。プレミアムは、ルーティング コードではなく、それらの関係とその需要に対してのものです。
OpenRouter を離れるのはまだ簡単です
アマゾンの例えには明らかに限界があります。多くの販売者はアマゾンを離れるわけにはいきません。 OpenRouter ユーザーにはまだ選択肢があります。
開発者は、プロバイダーの固定、最大価格の設定、レイテンシやスループットの並べ替え、独自のプロバイダー キーの使用、別のゲートウェイへの移動、または直接統合を行うことができます。大規模なチームは、オープンソースのルーティング ソフトウェアを自分たちで実行できます。プロバイダーは複数のゲートウェイを介して販売し、顧客との直接的な関係を維持できます。
スイッチングが安価である限り、Stripe はどちらの側にもそれほど強く圧迫することはできません。
OpenRouter は決して完全に中立ではありませんでした。そのデフォルトでは、適切なルートとは何かというビューがすでにエンコードされています。それが製品です。開発者は、ルールが文書化されており、通常は開発者の利益にかなうため、これを受け入れます。
Stripe は、より多くの製品を OpenRouter に組み込む予定です。

それが購入のポイントです。リスクは、OpenRouter が商業目的になることではありません。すでにそうです。リスクは、開発者が、ルーティング、価格設定、製品の決定が、Stripe 向けに行われる前に行われているのではないかと疑い始めることです。
そうなっても、彼らはそれを回避することができます。
Stripe は Amazon の地位を獲得しようとしているのかもしれない。同社はアマゾンの囲い込みを買っていない。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
[Ruby] BigDecimal の使用について
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
「QA フリー」の理念に反する
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
B2B 支払いの構造
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
開発ツールではコストがパフォーマンスに優先します
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
開発環境
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
Webflow の読み込み遅延の問題を修正する
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
ORM は依然として「アンチパターン」ですか?
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
オープンソース製品の使用に関する誤解
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
YC の後に米国に移住する場合: SF、NYC...、それともマイアミ?
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
あなたがエンジニアである場合の「非エンジニア」への面接について [危険信号]
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
オープンソースは安価なだけでは勝てない
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
最初の YC スタートアップの事後分析: リバース ETL
ああ、ああ！
エラーが発生しました

ルをロードしています。このページをリロードしてください。
読み込み中
AI 製品の価格設定は単位ベースの問題である
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
低迷市場におけるソフトウェア製品の価格設定
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
Product Hunt の開始: ハンドブック
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
払い戻し、クーポン、クレジットノート: それぞれが異なる理由
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
Replit は Shopify-2012 より前です
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
ストライプ データとオープンソースの代替案: MRR の例
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
Stripe の実際の価格設定: 入門書
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
味はスペックです
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
777 ポイントの Hacker News が 000 スターを獲得しました
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
エンタープライズ AI 競争に勝つのは最高のモデルではない
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
レートの罠: 1 つのアーキテクチャ決定がいかに柔軟性を失わせるか
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
タイムゾーン: エンジニアのための生き残るヒント
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
必要なアルファベータプログラムをスキップする企業が多すぎる
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
Clickhouse を使用してイベント エンジンをスケールする
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
プリペイド クレジットについて誰かに教えてほしいこと
ああ、ああ！
途中でエラーが発生しました

eロード中。このページをリロードしてください。
読み込み中
ユーザーが多額の AI 請求を立てて支払わなかった場合の対処方法
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
レプリットの90億ドルの評価額が安く見える理由
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
Stripe の「オープン」が実際にはオープンではない理由 — 開発者のジレンマ
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中
OpenRouter により、Stripe は AI のアマゾンになるのか
OpenRouter により、Stripe は AI のアマゾンになるのでしょうか?
ゲートウェイはマーケットプレイスになる可能性がある
これが Stripe にとって 70 億ドルの価値がある理由
OpenRouter を離れるのはまだ簡単です
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open Source Metering and Usage Based Billing API ⭐️ Consumption tracking, Subscription management, Pricing iterations, Payment orchestration & Revenue analytics - With OpenRouter, is Stripe becoming the Amazon of AI · getlago/lago Wiki

With OpenRouter, is Stripe becoming the Amazon of AI · getlago/lago Wiki · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
getlago
/
lago
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
With OpenRouter, is Stripe becoming the Amazon of AI
With OpenRouter, is Stripe becoming the Amazon of AI?
Stripe has reportedly agreed to buy OpenRouter for more than $7 billion. Stripe has not confirmed the deal. Bloomberg reported it as finalized , and TechCrunch reported the same figure .
The price looks absurd if OpenRouter is an API proxy. It makes more sense if Stripe sees the beginnings of Amazon Marketplace for AI inference.
OpenRouter gives developers one account and one API for more than 400 models. Behind it, more than 70 inference providers compete to serve requests. OpenRouter says it now has 8 million users.
The code that forwards a request is not worth $7 billion. The right to decide where a large and growing pool of requests goes might be.
A gateway can become a marketplace
OpenRouter is usually described as a gateway. That sounds technical and fairly neutral: a request arrives, the gateway forwards it, a response comes back.
But OpenRouter also decides which suppliers get access to demand.
Its default routing logic first removes providers that have recently failed. It then favors cheaper providers among the remaining candidates, using the inverse square of price. In OpenRouter's own example , an endpoint charging $1 per million tokens is nine times more likely to receive the first request than one charging $3.
That is good for developers. They get lower prices, fallback capacity and fewer provider integrations to maintain.
For inference providers, it is a distribution system. Price and performance determine how much traffic they receive. A provider can win more demand by cutting its price or improving reliability. It can lose demand without a developer ever making an explicit decision to switch.
Amazon Marketplace did not become powerful because listing products online was difficult. It gathered buyers in one place, then controlled how merchants reached them. Search ranking and the Buy Box could matter as much as the seller's product.
OpenRouter is not there yet. But it already admits providers, measures their performance and allocates demand between them. As more developers accept the default route, its routing rules start to determine provider revenue.
OpenRouter charges customers a 5.5% fee when they buy credits and says it does not mark up the underlying provider's token price. The visible fee is simple. The more interesting asset is the demand sitting behind it.
Why this could be worth $7 billion to Stripe
Stripe already knows OpenRouter well. OpenRouter uses Stripe for payments, invoicing, tax and fraud prevention. In January, the two companies also made OpenRouter available through Stripe Projects, so a developer or coding agent could provision an account and receive an API key from Stripe's command line. Stripe described the partnership here , and OpenRouter described the Projects integration here .
Owning OpenRouter would move Stripe earlier in the transaction.
Today Stripe can see a customer payment. OpenRouter can see which model was requested, which provider served it and what that inference cost. Put the two together and Stripe can connect the cost of producing an AI feature with the revenue earned from selling it.
That is useful because AI products have unusually variable margins. Two similar customer actions can have very different costs depending on the model, context length, cache behavior, retries and provider selected. A company may pass that cost through, add a markup, spend down credits or absorb it inside a subscription.
Stripe can sell more than payment processing around that activity: metering, credits, billing, tax, fraud controls and provider settlement. OpenRouter gives it a place to distribute those products at the moment the cost is created.
Stripe has already been assembling that stack. It acquired Metronome for roughly $1 billion to go deeper into high-volume usage metering and complex pricing. I wrote about why Stripe bought it instead of rebuilding Stripe Billing . Metronome is now marketed as a Stripe product, but it still has its own application and product surface alongside Stripe Billing. Buying the missing layer was faster than rebuilding it. Making two products feel like one is slower.
The reported PayPal bid belongs in the background too. Reuters reported that Stripe and Advent offered more than $53 billion for PayPal. PayPal would bring merchants and consumers. OpenRouter brings developers and inference providers. Both are networks in which several parties need money moved, reconciled and monetized.
Stripe could build an AI router for far less than $7 billion. It could not quickly reproduce OpenRouter's provider relationships or persuade 8 million users to move production traffic through a new one. The premium is for those relationships and that demand, not the routing code.
OpenRouter is still easy enough to leave
The Amazon analogy has an obvious limit. Many merchants cannot afford to leave Amazon. OpenRouter users still have options.
Developers can pin a provider, set a maximum price, sort for latency or throughput, bring their own provider keys, move to another gateway or integrate directly. Large teams can run open-source routing software themselves. Providers can sell through several gateways and keep direct customer relationships.
As long as switching remains cheap, Stripe cannot squeeze either side very hard.
OpenRouter has never been perfectly neutral. Its defaults already encode a view of what a good route is. That is the product. Developers accept it because the rules are documented and usually serve their interests.
Stripe will put more of its products into OpenRouter. That is the point of buying it. The risk is not that OpenRouter becomes commercially motivated. It already is. The risk is that developers begin to suspect routing, pricing or product decisions are being made for Stripe before they are being made for them.
If that happens, they can still route around it.
Stripe may be buying its way into the Amazon position. It has not bought Amazon's lock-in.
There was an error while loading. Please reload this page .
Loading
[Ruby] On using BigDecimal
Uh oh!
There was an error while loading. Please reload this page .
Loading
Against the ‘QA‐free’ philosophy
Uh oh!
There was an error while loading. Please reload this page .
Loading
Anatomy of a B2B payment
Uh oh!
There was an error while loading. Please reload this page .
Loading
Cost precedes performance for devtools
Uh oh!
There was an error while loading. Please reload this page .
Loading
Development Environment
Uh oh!
There was an error while loading. Please reload this page .
Loading
Fixing Webflow's load latency issues
Uh oh!
There was an error while loading. Please reload this page .
Loading
Is ORM still an 'anti pattern'?
Uh oh!
There was an error while loading. Please reload this page .
Loading
Misconceptions About Using Open‐Source Products
Uh oh!
There was an error while loading. Please reload this page .
Loading
Moving to the US after YC: SF, NYC... or Miami?
Uh oh!
There was an error while loading. Please reload this page .
Loading
On interviewing "non engineers" when you're an engineer [Red Flags]
Uh oh!
There was an error while loading. Please reload this page .
Loading
Open Source does not win by being cheaper
Uh oh!
There was an error while loading. Please reload this page .
Loading
Post mortem of our 1st YC startup: a Reverse ETL
Uh oh!
There was an error while loading. Please reload this page .
Loading
Pricing AI products is a unit‐based problem
Uh oh!
There was an error while loading. Please reload this page .
Loading
Pricing software products in a down market
Uh oh!
There was an error while loading. Please reload this page .
Loading
Product Hunt launch : our handbook
Uh oh!
There was an error while loading. Please reload this page .
Loading
Refunds, Coupons & Credit Notes: why they are different
Uh oh!
There was an error while loading. Please reload this page .
Loading
Replit Is Pre‐Shopify‐2012
Uh oh!
There was an error while loading. Please reload this page .
Loading
Stripe Data vs Open‐Source Alternatives: a MRR example
Uh oh!
There was an error while loading. Please reload this page .
Loading
Stripe's real pricing: a primer
Uh oh!
There was an error while loading. Please reload this page .
Loading
Taste is the spec
Uh oh!
There was an error while loading. Please reload this page .
Loading
The 777‐point Hacker News hit that got us 000 star
Uh oh!
There was an error while loading. Please reload this page .
Loading
The enterprise AI race won't be won by the best model
Uh oh!
There was an error while loading. Please reload this page .
Loading
The rate trap: how one architecture decision kills flexibility
Uh oh!
There was an error while loading. Please reload this page .
Loading
Timezones: survival tips for engineers
Uh oh!
There was an error while loading. Please reload this page .
Loading
Too many companies skip the needed alpha beta program
Uh oh!
There was an error while loading. Please reload this page .
Loading
Using Clickhouse to scale an events engine
Uh oh!
There was an error while loading. Please reload this page .
Loading
What I Wish Someone Told Me About Prepaid Credits
Uh oh!
There was an error while loading. Please reload this page .
Loading
What to do if a user racks up a massive AI bill and doesn't pay
Uh oh!
There was an error while loading. Please reload this page .
Loading
Why Replit's $9B Valuation Looks Cheap
Uh oh!
There was an error while loading. Please reload this page .
Loading
Why Stripe’s ‘Open’ Isn’t Really Open—The Developer’s Dilemma
Uh oh!
There was an error while loading. Please reload this page .
Loading
With OpenRouter, is Stripe becoming the Amazon of AI
With OpenRouter, is Stripe becoming the Amazon of AI?
A gateway can become a marketplace
Why this could be worth $7 billion to Stripe
OpenRouter is still easy enough to leave
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
