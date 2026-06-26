---
source: "https://www.coinerella.com/energy-based-llm-billing-cut-my-bill-to-a-sixth/"
hn_url: "https://news.ycombinator.com/item?id=48684027"
title: "Paying for LLM inference by the kilowatt-hour instead of per token"
article_title: "Neural What? My LLM bill is down to a sixth - by no longer paying per token."
author: "willy__"
captured_at: "2026-06-26T09:03:04Z"
capture_tool: "hn-digest"
hn_id: 48684027
score: 1
comments: 0
posted_at: "2026-06-26T08:37:56Z"
tags:
  - hacker-news
  - translated
---

# Paying for LLM inference by the kilowatt-hour instead of per token

- HN: [48684027](https://news.ycombinator.com/item?id=48684027)
- Source: [www.coinerella.com](https://www.coinerella.com/energy-based-llm-billing-cut-my-bill-to-a-sixth/)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T08:37:56Z

## Translation

タイトル: LLM 推論の料金をトークンごとではなくキロワット時単位で支払う
記事タイトル: ニューラルとは？トークンごとに支払う必要がなくなったため、LLM の請求額は 6 分の 1 に減りました。
説明: 最近このブログで読んだかもしれませんが、hank.parts に対する私の調達の好みは基本的に次のとおりです。
* EU、
* (自己ホスト型) オープンソース、
* 英国/スイス、
* 世界のその他の地域、
この順序で。
この記事は「その他の世界」の事例であり、コンセプトと

記事本文:
神経って何？トークンごとに支払う必要がなくなったため、LLM の請求額は 6 分の 1 に減りました。
最近このブログで読んだ方もいるかもしれませんが、hank.parts に対する私の調達の好みは基本的に次のとおりです。
この記事は、コンセプトと製品の両方に対する私の第一印象に圧倒され、法的に問題のあることに参加しているのではないかと思った「その他の地域」の事例です。 「ちょっと待って、これはうますぎる」という稀な瞬間が、実際にあったのです...本当です。
私は日常のタスク、API 関係、コーディング用のオープン ウェイト モデルの大ファンです。これらは Nebius 経由で推論するか、歯を食いしばって OpenRouter で推論します。どちらもインプットとアウトプットの価格設定で動作し、各モデルは消費された 100 万トークンごとに課金されます。
それか、北半球では夏以外の月の満月か隔週金曜日の 12 時 34 分にリセットされる魔法の使用量割り当てを与える月額サブスクリプションを利用しているかのどちらかです。これらのプランは、API の使用に関して非常に制限的な場合があります。ビッグ 3 では、ツールの外でモデルを操作したい場合に、再びトークン ベースのメータリングを強制されます。これらのフロンティア モデルの使用プランを作成すると、確実に価格が上昇するか、ごく近い将来クォータ調整が行われることになるため、API トークン ベースのメータリングを介して非常に高価になります。例: Claude Opus 4.8 は、OpenRouter 上で入力トークン 100 万あたり 5 米ドル、出力トークン 100 万あたりなんと 25 米ドルです。
ということで、現状維持です。プランの料金を支払い、含まれる API アクセスを祈るか、実際に消費したトークンの料金を支払います。少なくとも私は一ヶ月ほど前まではそう思っていました。
次に、LinkedIn で Vito Botta の投稿を読みました。フォローすることを強くお勧めします。
ニューラルワット?聞いたこともありません。エネルギーベースの計量?わかった？！コンセプトについては Vito がすでに十分に説明しているので、ここで説明します。

コストの比較をご覧ください:
米国 A の NeuralWatt は、トークンとエネルギーベースの請求の両方を提供しているため、自分の節約額を直接比較できます。平均でなんと 82.9%、Qwen3.6-35b-fast では 95.2% であり、これは私にとって驚くべきことです。トークンベースのメータリングを使用して同じプラットフォームで同じ使用量を使用した場合、さらに 302.32 米ドルを支払うことになります。
また、私がコーディングに使用している Kim モデルのキャッシュ番号にも注目してください。 NeuralWatt では、繰り返しの入力に対して非常に効率的なキャッシュが行われているようで、コスト削減を顧客に問題なく提供しています。トークンベースの価格設定には個別の「キャッシュされた」レートもあり、入力/出力価格設定のデュオに一定の特徴を与えています。
これらすべてがとてもエキサイティングです。エネルギー消費を認識し、同時に安価な LLM 使用量を測定する新しい方法?!
NeuralWatt を使い始めて最初の 1 か月間で気づいたことをいくつか紹介します。
同時リクエストとレート制限は、他のプラットフォームからの制限であると感じる場合があります。指数バックオフを備えた再試行メカニズムにより、これまでのところ最初の問題は解決されています。
より長い応答に対していくつかの 524 エラーがポップアップし続けたため、API 応答にはストリーミング モードを使用する必要がありましたが、ストリーミングに切り替えてからはすべて問題ありませんでした。
Nebius に慣れていたものと比較すると、パフォーマンスが少し遅く感じることがあります。ここ数日は特に。成長痛だと思いますか？
DPA を取得する簡単な方法は見つかりませんでしたが、NeuralWatt の使用は GDPR 関連ではないため、あまり難しく考えませんでした。
データ保持ポリシーは厳密には ZDR ではありません。クロスリクエスト キャッシュのトレードオフだと思います。ユーザーデータに関するトレーニングは自信を生みません。
完全開示パート #1: 私は眠れない木曜日の夜にこの記事を携帯電話に入力し、LLM が私の支離滅裂なとりとめのない文章をこの最後の記事に構造化しました。 LLM はどれだと思いますか?
完全開示パート #2: 私は NeuralWatt の有料顧客であり、ニートです

この記事の執筆時点では、私も私の企業も NeuralWatt とさらなる関係を持っています。
車の部品を取引していますか、または取引している人を知っていますか?ハンクに一言欲しいです。
自分でメールをホストしないでください - 2026 年のリマインダー
「Made in EU」、思ったより大変でした。

## Original Extract

You might have read recently on this blog that my procurement preferences for hank.parts are basically
* EU,
* (self hosted) open source,
* UK/CH,
* Rest of the world,
in this order.
This article is a "rest of the world" case where my first impression of both the concept and the

Neural What? My LLM bill is down to a sixth - by no longer paying per token.
You might have read recently on this blog that my procurement preferences for hank.parts are basically
This article is a " rest of the world " case where my first impression of both the concept and the product overwhelmed me to a point where I thought I was taking part in something legally questionable. A rare "wait, this sounds too good to be true" moment that actually was... True.
I am a big fan of open weight models for daily tasks, API stuff and coding, which I either infer via Nebius or, gritting my teeth, OpenRouter. Both operate on input and output pricing where each model is metered per million tokens consumed.
It's either that or you are on some monthly subscription giving you some magic usage quotas that reset at the full moon or at 12:34 every other Friday in a non-summer month in the northern hemisphere. These plans can be quite restrictive with API usage, where the big three force you into token-based metering again if you want to interact with the models outside their tooling... Making these frontier model usage plans, which surely will rise in price or will see quota adjustments in the very near future, extremely expensive via API token-based metering. For example: Claude Opus 4.8 is 5 USD per million input tokens and a whopping 25 USD per million output tokens on OpenRouter.
So, that's the status quo. Pay for a plan and pray for included API access or pay for the tokens you actually consume . At least that's what I thought until more or less a month ago.
Then I read a post from Vito Botta , whom I highly recommend following, on LinkedIn:
NeuralWatt? Never heard of it. Energy-based metering ? Okay?! Vito explained the concept already well enough, so I will jump straight to the cost comparison:
NeuralWatt from the US of A offers both token and energy-based billing so I can directly compare my savings , a whopping 82.9% on average and 95.2% for Qwen3.6-35b-fast, which is mind-blowing to me. I would have paid 302.32 USD more for the same usage on the same platform with token-based metering.
Also, look at these cache numbers for the Kimi models I use for coding! NeuralWatt seem to have some very efficient caching going on for repeat input and have no problem with forwarding the cost savings on to the customer. Their token-based pricing even has separate "cached" rates, giving the input/output pricing duo some company.
All this is very exciting. A new way to meter LLM usage that is aware of energy consumption AND cheaper at the same time?!
Here are some things I noticed over my first month on NeuralWatt:
Concurrent-request and rate limits can feel restrictive coming from other platforms. Retry mechanisms with exponential backoff have solved my initial problems so far.
I had to use streaming mode for API responses, since some 524 errors kept popping up for longer answers - all fine since the streaming switch.
Performance sometimes can feel a bit slow compared to what I was used to from Nebius. Last couple of days especially; growing pains I guess?!
I didn't see a quick way to obtain a DPA , but I did not look too hard since my NeuralWatt usage is not GDPR-relevant.
Data retention policy is not strictly ZDR - I guess a cross-request caching trade-off? No training on user data inspires confidence.
Full disclosure part #1: I typed this article into my phone on a sleepless Thursday night and an LLM structured my incoherent ramblings into this final article. Guess which LLM?
Full disclosure part #2: I am a paying NeuralWatt customer and neither I nor my businesses have further associations with NeuralWatt at the time of writing.
Do you trade car parts or know someone who does? hank would like a word.
Don't host email yourself - your reminder in 2026
"Made in EU" - it was harder than I thought.
