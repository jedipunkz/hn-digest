---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49318304"
title: "Show HN: RNet – AI token service provider"
article_title: ""
author: "nextma"
captured_at: "2026-08-16T09:17:41Z"
capture_tool: "hn-digest"
hn_id: 49318304
score: 1
comments: 0
posted_at: "2026-08-16T09:11:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RNet – AI token service provider

- HN: [49318304](https://news.ycombinator.com/item?id=49318304)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T09:11:32Z

## Translation

タイトル: HN を表示: RNet – AI トークン サービス プロバイダー
HN テキスト: rNet を構築しました。なぜ私が建てたのか？ : 私はエージェント IDE と Hostinger デプロイメント エージェントの両方を使用していました。ある日、展開エージェントのクレジットが不足してしまいました。使い続けるには、クレジットがリセットされるまで待つか、より高いサブスクリプションにアップグレードするか、トークンを購入する必要がありました。同時に、私はすでに IDE のサブスクリプションを持っていましたが、それらのクレジットを Hostinger で使用することはできませんでした。クレジットを持っているにもかかわらず、それを使用できないだけです。基本的な考え方はシンプルです。ユーザーはプラットフォームでクレジットを一度購入すると、それらのクレジットをさまざまな AI アプリケーションで使用します。 (OpenRouter のように聞こえますが、2 つのプラットフォームはまったく異なる問題を解決します。) 仕組み: 1. まず開発者が製品をプラットフォームに登録し、AI ゲートウェイを統合して AI モデルを呼び出します。
2. ユーザーは当社のプラットフォームからクレジットを購入し、自分のアカウントをサポートされている AI 製品に接続します。その後、サポートされているすべての AI 製品で同じクレジットを使用できます。詳細情報「仕組み」:- 開発者向け:
rNet は単なる AI ゲートウェイです。開発者はゲートウェイを使用して AI モデルを呼び出します。これは開発者向けの支払いレイヤーではありません。信用の流れ:
ユーザー → rNet がクレジットを保持 → AI プロバイダー 例:
ユーザーは rNet から 10 ドルのクレジットを購入します。これらのクレジットはユーザーのアカウントに保持されます。ユーザーがサポートされている AI アプリケーションに接続し、AI モデルを使用する場合:
1. アプリケーションは AI ゲートウェイ経由で AI リクエストを送信します。
2. rNet は、適切なユーザーの残高から必要なクレジットを差し引きます。
3. その後、その使用量に対して AI プロバイダーに支払います。これはインターネット データ プランと同様に機能します (これが、「AI トークン サービス プロバイダー」を「インターネット サービス プロバイダー」のように書いている理由です): 2 GB のインターネット データを購入すると、そのデータをハッカー ニュース、X、またはその他のサービスで使用できます。必要ありません

アプリごとに個別のデータプラン。データがゼロになるまでデータを使用します。
同様に、rNet を使用すると、ユーザーは AI クレジットを一度購入すると、ユーザーの残高がゼロになるまで、サポートされている AI アプリケーション全体でそれらのクレジットを使用できます。その後、引き続き使用するには再充電する必要があります。デモビデオ: https://youtu.be/W7U3HdI37N0
製品の基本バージョンは公開中です。 https://www.rnetai.org/ 結果 :- 1. ユーザーは複数の AI 製品で同じクレジットを使用することでコストを節約できます。
2. 開発者は AI トークンのコストを前払いする必要はありません。
3. シンプルなフロー 私が検討している将来の機能 (まだ構築されていません): 1. 開発者はオープンソース モデルを微調整し、自社の製品で使用できます。
2. AIモデル統一体
3. 企業向けのエンタープライズ版
4. ユーザーはAIクレジットをお金のように送受信できる
5. 開発者とユーザーの両方のための機能の追加 これが実際の問題を解決するかどうかについてのフィードバックを特にお待ちしています。また、可能であれば、自分の言葉で製品について説明してください。私たちは現在、卵が先か鶏が先かの問題に直面しているため、これを作り続けるべきか、それともゴミ箱に捨てて次に進むべきかを判断するために待機リストを作成することにしました。このコンセプトが気に入ったら、順番待ちリストに参加してください: https://www.rnetai.org/reserve-spot

## Original Extract

I built rNet. Why I built ? : I was using both an agentic IDE and a Hostinger deployment agent. One day, I ran out of credits on the deployment agent. To keep using it, I either had to wait for credits to reset or upgrade to a higher subscription or buy tokens. At the same time, I already had a subscription for the IDE, but I could not use those credits on Hostinger. simply despite having credits, we cannot use them. The basic idea is simple: users buy credits once on our platform and then use those credits across different AI applications. (It sound like OpenRouter, but the two platforms are solve totally different problems.) How it works: 1. First developers register their product on the our platform and integrate our AI gatway to call AI models
2. Users purchase credits from our platform and connect their accounts to supported AI products. They can then use the same credits across all supported AI products. In details information "how it works" :- For developers:
rNet is simply an AI Gateway. Developers use our gateway to call AI models. It is not a payment layer for developers. Credit flow:
User → rNet holds credits → AI Provider Example:
A user buys $10 of credits from rNet. We hold those credits in the user's account. When the user connects to a supported AI application and uses an AI model:
1. The application sends the AI request through our AI Gateway.
2. rNet deduct the required credits from the right user's balance.
3. then we pay the AI provider for that usage. It works similar to an internet data plan (that why i write "AI token service provider" like "internet service provider"): You buy 2 GB of internet data, and then you can use that data on hacker news, X or any other service. You don't need a separate data plan for each app. you use data until their data reaches zero.
Similarly, with rNet, users buy AI credits once and can use those credits across any supported AI applications until user's balance reaches zero. Then they need to recharge to continue using them. demo video : https://youtu.be/W7U3HdI37N0
basic version of product is live. https://www.rnetai.org/ Result :- 1. users can save money by using the same credits across multiple AI products.
2. Developers don't need to pay AI token costs upfront.
3. simple flow Future features I'm considering (not built yet): 1. Developers can fine-tune open-source models and use them in their products
2. AI model unified body
3. An enterprise version for companies
4. Users can send and receive AI credits like money
5. More features for both developers and users We’d especially love feedback on whether this solves a real problem or not ? and also if possible then you would describe the product in your own words. We’re currently facing a chicken-and-egg problem, so we decided to create a waitlist to help us figure out whether we should keep building this or put it in the trash and move on. If you like the concept, then join our waitlist: https://www.rnetai.org/reserve-spot Thank You

