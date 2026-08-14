---
source: "https://www.jevy.org/articles/my-vp-eng-told-me-to-stop-working-on-observability/"
hn_url: "https://news.ycombinator.com/item?id=49297738"
title: "My VP Eng told me to stop working on observability"
article_title: "My VP Eng Told Me to Stop Working on Observability. Here's Why He Was Wrong. · Jevin Maltais"
author: "jevyjevjevs"
captured_at: "2026-08-14T12:41:47Z"
capture_tool: "hn-digest"
hn_id: 49297738
score: 1
comments: 0
posted_at: "2026-08-14T12:17:00Z"
tags:
  - hacker-news
  - translated
---

# My VP Eng told me to stop working on observability

- HN: [49297738](https://news.ycombinator.com/item?id=49297738)
- Source: [www.jevy.org](https://www.jevy.org/articles/my-vp-eng-told-me-to-stop-working-on-observability/)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T12:17:00Z

## Translation

タイトル: 担当副社長が可観測性への取り組みをやめるように言った
記事のタイトル: 私の担当副社長は、可観測性への取り組みをやめるように私に言いました。彼が間違っていた理由は次のとおりです。 · ジェビン・マルタイス
説明: 私は、急速に成長している小規模な新興企業のスタッフ エンジニアでした。私たちの新しい技術担当副社長が私を脇に引きました。
開発者のエクスペリエンスと可観測性を非常に重視しているようですが、機能の開発にもっと投資してもらいたいと思います。
私は彼がどこから来たのか理解しました、私たちは成長する重要な時期にありました
[切り捨てられた]

記事本文:
私の担当副社長は私に、可観測性への取り組みをやめるよう言いました。彼が間違っていた理由は次のとおりです。 · ジェビン・マルタイス
ジェビン・マルタイス
Fractional シリーズ A/B 企業のエンジニア副社長兼主任エンジニア。私はエンジニアリング チームを修正し、本番環境で動作する AI を実装し、新しい製品ラインを構築します。
私の担当副社長は私に、可観測性への取り組みをやめるよう言いました。彼が間違っていた理由は次のとおりです。
私は急成長を遂げている小規模な新興企業のスタッフ エンジニアでした。私たちの新しい技術担当副社長が私を脇に引きました。
開発者のエクスペリエンスと可観測性を非常に重視しているようですが、機能の開発にもっと投資してもらいたいと思います。
私は彼がどこから来たのかを理解しました。私たちはビジネスを成長させる重要な時期にあり、機能は収益拡大への明確な道筋でした。しかし、テクノロジー負債に直接関係する、彼には見えない部分がありました。
1、2週間ごとに停電がありました
ユーザーは SSE との接続で問題を抱えていました
AI エージェントは信頼性の低い失敗をしていました。
何らかの問題が発生するたびに、問題を見つけて修正を開始するのに少なくとも 1 時間 (または 2 時間!) かかる必要がありました。大変だったし、良くならなかった。バグやエラーはよくあるもので、解決までに長い時間がかかる場合があります。問題を発見し、問題を切り分け、影響を理解し、修正するためのループを大幅に削減するためにできることは何でも、最終的にはユーザー エクスペリエンスを向上させ、実際に製品を構築するための時間を増やすことができます。
私の考える可観測性とは、システムの継続的な健全性と問題の発見を理解するために必要な全体像です。可観測性の品質は、次の要素を組み合わせた尺度です。
開発者が問題の原因をどれだけ早く見つけられるか。
その問題がユーザーにどの程度の影響を与えているかを判断します。
プロアクティブな問題が重大度別にどのように発生するか、ユーザーによってもたらされるか。
最初の 1 つ

チームに参加するときに私がやっているのは、自分たちの責任の範囲と観察可能性がどのようなものかを把握することです。私が見たTierリストは次のとおりです。
1. すぐに使えるツール以外は何もありません - 完全に苦情主導型
デプロイに使用した Vercel ログまたは kubernetes ログを使用しているだけです。
アラートが何もないのに、どうしてできるでしょうか？例外を発行していません
バグレポートはすべてユーザーから届きます。
2. 統合機能を備えた Datadog を追加しました (私が参加したほとんどのチーム)
おめでとうございます！ Datadog に資金を提供するためにテクノロジー企業が存在するのであれば、それは成功です。
ボーナス
ホストされている Signoz をチェックしてください。同様の機能、オープンソース、低コスト、Open Telemetry ベース。
自動インスツルメンテーション テレメトリが有効になっているため、基本的なログと例外がプッシュされます。
多数の統合とそのデフォルトのダッシュボードを追加しました
どこかでエラーが発生したとき、確実な状況を把握できます
チームにアラートとページングを行うためのフレームワークがある
すべてのデータを 1 か所に保管できます。
MCP にはデータがあるため、基本的な調査を独自に行うことができます。ログの完全な帰属が存在しないため、多くの情報が失われる可能性があります。
すべてのサービスが切断されます。ブラウザからのリクエストはルートから切断され、データベース クエリから切断され、Redis キャッシュから切断され、ジョブ ランナーから切断されます。特定のユーザーのリクエストがどこで失敗したかを追跡するのは悪夢です
あなたの丸太はゴミ箱の火です。まったくキュレーションを行っていません。重複したものや本当に必要のないものがたくさんあります。
ログは構造化されていません。エラーを特定のユーザーに分離することは不可能なので、ユーザーを詳しく調査するときに、時間に基づいて関連付けようとしているとは考えられません。どれの

最低だ。検索ツールは優れていますが、フリーテキストでしか検索できません
どこにでもエラーが重複しているため、影響の範囲はありません。エラーをグループ化していないため、特定の問題がどのくらいの頻度で発生しているのか、またはユーザーに影響を与えているのかを知ることができません。
コスト - すべてのデータを Datadog などにダンプすると、特にすべてがまだ混乱しているために期待した値が得られない場合、すぐに費用がかかる可能性があります。
あなたのチームの誰かが、デバッグで十分なフラストレーションを抱えて、苦しむことを経験し、それらを結びつける作業を行ってきました。ユーザーが問題を抱えている場合は、パス全体を確認できます。
ブラウザリクエスト -> nginx/vercel -> ルーター -> ビジネスロジックとジョブ -> データベースリクエスト
何が遅いのかについてのすぐに使えるメトリクス - すべてのトレースには開始時間と終了時間が設定されているため、システムの何が遅いのかを詳細なレベルで客観的に把握できるようになりました。
完全なコンテキストでエラーを見つけます。冒頭で述べた SSE エラーを覚えていますか?トレースを使用すると、受信したすべてのリクエストを確認できましたが、バックエンド サービスがトリガーされていないことも確認できました。 2 つを接続しないと、他のユーザーと同じように、ユーザーの受信リクエストのみが表示されました。フルパスが確認できると、パターンが明らかになりました。影響を受けるユーザーはすべて、SSE 接続をブロックしている企業ファイアウォールの内側にいたということです。数週間にわたる混乱が、午後には解決した。
誰も報告していない問題を表面化 - すぐに使えるもう 1 つの成果: トレースが入ったら、1000 人のユーザー チャット セッションのサンプルの散布図を作成しました。ユーザーの 15% がサイレントにタイムアウトしていました。タイムアウトしたユーザーのログとタイムアウトしなかったユーザーのログを比較すると、プロバイダーが返されたプロンプトを途中で遮断し、ユーザーに何もストリーミングしていないことがわかりました。それらのユーザーの中でバグを報告した人は一人もいませんでした

報告。彼らはただ悪い経験をしただけで、次に進みました。
MCP は、問題解決の再現性を高めるためにトレースを上下に移動して原因となる問題を見つける方法を知っているため、問題を迅速に見つけることができます。
自分にとって重要なエンドポイント、サービス、SLA を監視する非常にクリーンなダッシュボードを作成できます。
これには、積極的に保守するための手動作業が必要です。トレースの形状とそれに追加するメタデータを決定する必要があります。
フロントエンドの可観測性のための新しいルート、おそらく 1 つまたは 2 つのコレクター、または生のログを取得し、トレースし、フィルター処理するなど、追加のインフラストラクチャを立ち上げる必要がある場合があります。小規模なチームでやりたいことができる場合を除き、セットアップに時間を割り当てる必要があります。
まだ大量のデータを処理する必要があり、データの 80% はおそらくノイズです。
4. ニルヴァーナ - 開発の一環としての観察可能性
チームの何人かが定期的にダッシュボードを使用しており、データがクリーンであると確信しており、ツールに精通しています。
機能を構築する前に、リリース後にその機能を実際に使用する人の数とそのパフォーマンスをどのように測定するかを考えます。機能をデプロイすると、その機能がどのように機能しているかをすぐに確認できるようになり、機能を改善する能力が高まります。
あなたは実戦テストを経た強固なプラットフォームを持っており、報告された問題がどこにあるのかを正確に把握している可能性が高いでしょう。
それが確立されたパターンとなっているため、チームは可観測性のための本当に強固なベスト プラクティスを構築できるようになります。
ここに到達するには多大な努力が必要であり、組織はそれを常に把握しておく必要があります。問題をデバッグするために徹底的に調査する必要がある場合、怠惰でベスト プラクティスを実行せずに盲点を作る人もいます。
新しいサービスにはアプローチが必要です。

追加のテクノロジー層をパスに追加する場合は常に、それを分散トレースの一部にする必要があり、これには労力がかかり、理解するのは簡単ではない場合があります。
副社長に戻ります。機能がビジネスを成長させるという点では彼は間違っていませんでしたが、可観測性と機能の取り組みという枠組みは誤った選択であることが判明しました。何週間も続いていた SSE 問題は、痕跡が見つかってから午後には解決しました。チャット セッションの 15% にサイレント タイムアウトが発生していますか?影響を受けるユーザーが 1 人もいなかったため、どれだけの機能を開発してもそのようなことは表面化することはなかったでしょう。彼らは静かに悪い経験をしただけです。
切断された丸太を掘るのに費やされる 1 時間は、構築に費やされない 1 時間になります。目に見えない問題はすべて、ユーザーを失う可能性があるということです。可観測性への投資によって機能の作業が遅れることはありませんでした。それが私たちにそれを行うための時間と自信を与えてくれました。
現在ティア 1 または 2 にいる場合、厳選された分散トレースへの移行が最大のロック解除です。これは大変な作業ですが、「奇妙な断続的な問題」が数週間にわたる謎ではなく、5 分間の調査に変わったときに初めて元が取れます。
これらの投稿を書いていて気に入っていることの 1 つは、常に何人かの人々が連絡をくれてつながりを持てることです。ここで何か楽しいことはありますか？
殴ってくれ！ jevin@quickjack.ca 。

## Original Extract

I was a Staff Engineer at a small, quickly growing start-up. Our new VP eng pulled me aside:
it seems like you care a lot about developer experience and observability, but I need you to invest more in the feature work.
I understood where he was coming from, we were in a critical time of growing the
[truncated]

My VP Eng Told Me to Stop Working on Observability. Here's Why He Was Wrong. · Jevin Maltais
Jevin Maltais
Fractional VP Eng & Principal Engineer for Series A/B companies. I fix engineering teams, implement AI that works in production, and build new product lines.
My VP Eng Told Me to Stop Working on Observability. Here's Why He Was Wrong.
I was a Staff Engineer at a small, quickly growing start-up. Our new VP eng pulled me aside:
it seems like you care a lot about developer experience and observability, but I need you to invest more in the feature work.
I understood where he was coming from, we were in a critical time of growing the business and the features were the clear path to growing revenue. However, the unseen part that he couldn’t see that were directly related to tech debt:
we had outages every week or two
users were getting issues connecting with SSE
AI agents were failing unreliably.
Everytime some issue happened, I had to take at least an hour (or two!) to find the issue then start in on the fix. It was a grind and it wasn’t getting better. Bugs and errors are common and can take a long time to resolve. Anything we can do to significantly cut down on the loop for discovering the issue, isolating the problem, understanding impact and fixing them will at the end of the day allow for a better user experience, and allow for more time to actually build the product .
Observability in my mind is: the holistic picture we have to understand the ongoing health of the system and finding issues. The quality of the observability is a combined measure of:
How quickly a dev can find the source of an issue.
Determine how widely that issue is impacting users.
How proactive issues bubble up by severity vs being brought to us by the users.
One of the first things I do when joining a team is getting the lay of the land of the scope of our responsibility and what our observability looks like. Here is the tier list I’ve seen:
1. Nothing apart from the out of the box tools - entirely complaint driven
you’re just using the Vercel logs or kubernetes logs you’re deployed with.
you don’t have any alerting because how could you? You aren’t emitting exceptions
all your bug reports come in from your users.
2. You added Datadog with integrations (Most teams I’ve been on)
Congratulations! If technology companies exist in part to give Datadog money, you’ve made it!
Bonus
Check out hosted Signoz. Similar functionality, open source, less expensive and Open Telemetry based.
you have auto-instrumenting telemetry enabled so you get basic logs and exceptions pushed
you added a bunch of your integrations and their default dashboards
you have a solid picture when an error is emitted somewhere
you have a framework for alerting and paging your team
you have a single place for all your data.
MCP can do a basic investigation on its own since the data is there. It may miss a lot since the full attribution of a log isn’t there.
all your services are disconnected: a request from the browser is disconnected from your routes which is disconnected from your db queries which is disconnected from your redis cache which is disconnected from your job runner. Tracking down where a request went wrong for a particular user is a nightmare
your logs are a dumpster fire. You haven’t done any curation at all. Lots of duplication and stuff you really don’t need.
logs aren’t structured: it’s impossible to able to isolate errors to particular users so when you’re digging into a user, you have no idea you try to correlate based on time. Which sucks. While the search tooling is good, you can only search by free text
no scope of impact since you have duplication of errors everywhere - you don’t have any grouping of errors so you can’t tell how often certain issues are happening or user impact.
cost - dumping all your data into something like Datadog can get expensive quick, especially if you don’t get the value you were hoping for because everything is still a mess.
Someone on your team has had enough frustration debugging that they bit the bullet and they have done the work to tie together. When a user has an issue you can see the entire path:
The browser request -> nginx/vercel -> router -> business logic and jobs -> database requests
out of the box metrics on what is slow - since all traces have start and end times, you can now have an objective window into what is slow in your system at a granular level
find that error with full context - remember that SSE error I mentioned at the beginning? With traces we could see all the requests coming in but we could also see no backend services being triggered. Without connecting the two, we only saw the incoming request of the user just like everyone else. Once we could see the full path, the pattern jumped out: the affected users were all behind corporate firewalls that were blocking the SSE connections. Weeks of confusion, solved in an afternoon.
surface issues nobody is reporting - another out of the box win: once traces were in, I made a scatterplot of a sample of 1000 user chat sessions. 15% of users were silently timing out. Comparing the logs of the users who timed out against the ones who didn’t, we found the provider was cutting off the returned prompt midway and never streaming anything back to the user. Not a single one of those users filed a bug report. They just had a bad experience and moved on.
MCPs can quickly find the issues since it knows how to go up and down the trace to find the source issue to help with reproducibility to solve the issue.
you can create really clean dashboards monitoring the endpoints, services and SLAs that matter to you!
this requires manual work to actively maintain. You have to decide the shape of your traces and what metadata you want to add to it.
You might have additional infrastructure that needs to be stood up: new routes for your front end observability, maybe a collector or two or take the raw logs, traces and filter them. Unless you’re a small team and you can do what you want, you’ll need to allocate time to set it up.
you still have a ton of data to wade through and 80% of the data is probably noise.
4. Nirvana - Observability as part of your Development
A few folks on your team are using your dashboards regularly and you’re confident your data is clean and you’re very familiar with the tool.
In advance of building a feature, you’re now thinking how you’re going to measure how many people actually use it and its performance once it’s released. Once you deploy your feature, you have a window into how it’s doing right away accelerating your ability to improve it.
You have a solid platform that has been battle tested and you’re likely very good at knowing exactly where you can see any reported issue
Your team is enabled to build really solid best practices for observability since that’s now the established pattern
it’s taken a lot of work to get here and your org needs to stay on top of it. Some folks will be lazy and not implement best practices creating blind spots when YOU have to dig in to debug an issue.
you’ll need an approach for new service. Whenever you add an additional technology layer into your path, that needs to be part of your distributed trace and that takes effort and may not be trivial to figure out.
Back to my VP. He wasn’t wrong that features grow the business, but the framing of observability vs feature work turned out to be a false choice. The SSE issue that had been dragging on for weeks was solved in an afternoon once we had traces. The silent timeouts hitting 15% of chat sessions? No amount of feature work would have surfaced that, because not a single affected user told us. They just quietly had a bad experience.
Every hour spent digging through disconnected logs is an hour not spent building. Every issue you can’t see is a user you might be losing. Investing in observability didn’t slow down our feature work. It’s what gave us the time and the confidence to do it.
If you’re at tier 1 or 2 today, the jump to curated distributed traces is the single biggest unlock. It’s real work, but it pays for itself the first time a “weird intermittent issue” turns into a five minute investigation instead of a multi-week mystery.
One of my favourite things with writing these posts is that there are always a few folks who reach out and I get to connect with. Anything here fun?
Hit me up! jevin@quickjack.ca .
