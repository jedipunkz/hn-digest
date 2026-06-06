---
source: "https://byacommonthread.com/blog/what-should-i-eat"
hn_url: "https://news.ycombinator.com/item?id=48428548"
title: "Made my first plugin – How I orchestrated 3 LLMs to ship a plugin in 2 hours"
article_title: "What Should I Eat | ByACommonThread"
author: "kaydub"
captured_at: "2026-06-06T21:26:51Z"
capture_tool: "hn-digest"
hn_id: 48428548
score: 2
comments: 1
posted_at: "2026-06-06T20:12:06Z"
tags:
  - hacker-news
  - translated
---

# Made my first plugin – How I orchestrated 3 LLMs to ship a plugin in 2 hours

- HN: [48428548](https://news.ycombinator.com/item?id=48428548)
- Source: [byacommonthread.com](https://byacommonthread.com/blog/what-should-i-eat)
- Score: 2
- Comments: 1
- Posted: 2026-06-06T20:12:06Z

## Translation

タイトル: 初めてのプラグインの作成 – 3 つの LLM を調整して 2 時間でプラグインをリリースする方法
記事タイトル: 何を食べよう | ByACommonThread

記事本文:
ByACommonThread ホーム ブログ 何を食べるべきか
サイドプロジェクト · AIエージェント · aws
レストランで一番おいしい料理は何か、あるいは一番おいしい料理のランキングをクラウドソーシングするアプリがあることを私は知りませんでした。また、私はブラウザーのプラグインを公開したことがなかったので、それがどのように機能するかに興味がありました。そこでこのプラグインを作りました。
この v1 バージョンは Chrome および Firefox 用です (このブログ投稿の草稿を書いている時点では、Chrome バージョンのみが公開されています)。
これはおそらく Android/iOS アプリとしての方が便利なので、それについては説明しますが、その仕組みと、これを構築する際のワークフローについて、ここで詳細を記入できると思いました。
非常にシンプルです。v1 は Google マップでのみ機能します。地図上でレストランを選択し、プラグインをクリックします。これまでクリックしたことのないレストランの場合は、自動的に追加されるので、一番おいしいと思う料理を入力してください。以前にそのレストランをクリックしたことがある場合は、他の人の投票が表示され、新しいものを入力するか、表示された料理の 1 つをクリックすることができます。 UI は、料理の投票数に基づいてテキスト サイズが大きくなる「ピル クラウド」を使用しているため、強打者が視覚的に泡立ちます。
認証や入力された料理の検証はありません。これは物事をシンプルにするためでもありましたが、摩擦がゼロであるという特徴でもあります。ラザニアが美味しいと言うために、さらに別のサービスに登録する必要はありません。はい、重複して入手できます。はい、ユーザーはローカル データをクリアして、料理に複数回投票できます。
ただし、データのプライバシーやプラグインの承認などは扱いたくないので、非常にシンプルで基本的なものにすることにしました。
レストランのエントリーを正確に管理する方法を理解するのに苦労しました。将来的にはこれを Google マップ以外でも機能させたいので、Google Place を使用しないことにしました。

d. Google の ID に依存すると、ちょっとした人質にもなります。Google が内部 ID システムを変更したり、後で Apple Maps や Yelp をサポートしたくなったりした場合、ゼロからのスタートになります。
最終的には、GPS 座標を使用することにたどり着きました。 GPS 座標は、後でモバイル アプリの機能にうまく機能するため、プラスになります。そして、それらの GPS 座標に基づいてハッシュ/uuid を作成します。これらは座標に基づいて決定的であるため、単一テーブル dynamodb 設計での検索が簡単になります。また、Geohashing を使用して、レストランを約 5km 四方 (5 文字の接頭辞) に分割しています。これにより、API はテーブル全体をスキャンするのではなく、近くの「バケット」だけをデータベースにクエリできるため、DynamoDB のコストは基本的にゼロに保たれます。近くの結果が得られたら、ハバーサインの公式を使用して正確な距離をメートル単位で計算し、それらを並べ替えます。 (ここで注意しなければならないのは、私はハバーシン公式が何なのか全く知りませんでした。それが LLM が考え出したものです)
「クラスタリング」問題と GPS ドリフトに対処するために、ハッシュ化する前に座標を小数点以下 3 桁に四捨五入します。これにより、約 110 メートルの「あいまいな」スナップが得られます。つまり、レストランの反対側に立っている場合でも、同じ ID に解決されるはずです。狭い敷地内に多くのレストランが密集している場所では状況は少し悪くなりますが、プラットフォームに依存しないことを考慮して、これは喜んでトレードオフとします。将来的には、レストランが密集している場合に特定のレストランを選択する方法を追加する予定です。ある時点で、スポットが閉店し、新しいものとして再オープンするときにレストランの名前を変更する方法も考えなければなりません。
バックエンドも非常にシンプルです。すべては AWS 上でサーバーレス API として実行されます。すべてのコンピューティングは AWS Lambda にあり、Apigw がリクエストを処理し、Dynamod

b データストアの場合。認証やユーザー アカウントはなく、非常にシンプルな API です。ある程度までは自動的にスケールアップされ、ほとんど無料で実行できます。
そうですね、現在これを使用している人はほとんどいないので、現時点では無料です (ホストゾーンを除けば、私はそこから複数のことを実行していますが、月額わずか 0.50 ドルです)。これにお金がかかるまでになったら、かなり成功したと思います。スタック上の無料利用枠からロールオーバーすると、ユーザー数に応じて線形に拡張されます。
収益化はありません。コストが私の欲求を超えた場合にのみ、これを収益化することを検討します。月額 100 ドル以上の費用がかかる場合は、コストをカバーするために、シンプルで煩わしくない収益化方法を見つけようとします。現時点では、そのようなことは考えられていないので、あまり心配していません。
私が個人的にコードを書くことはほとんどなく、LLM にほとんどのコードを作成してもらいました。 Gemini、Claude、ChatGPT はすべて参加していました。これは主に Gemini のものだったと思います (これを書いている 2026 年 6 月の時点では、これは数あるツールの中で最悪のツールだと思います)。私は完全に手を離さず、コードに雰囲気を与え、ガイダンスを与え、特定の時点で LLM をゲートし、特定の検証手順/ツールを手動で実行します。むしろ、喜んでもらいたいと思っているジュニアエンジニアと会話しているような気分になります。
まず、何を構築したいのかを説明し、LLM に AGENTS/CLAUDE/GEMINI マークダウン ファイルの構築を手伝ってもらうように指示します。通常、この部分には 1 ～ 5 つのプロンプトが必要です。また、マークダウン ファイルを自分で何度か読み込み、ここで手動で微調整を加える可能性があります。
マークダウン ファイルが完成したら、確認したいアーキテクチャとコード パターンについて説明します。 LLM がレビューするために、いくつかのリファレンス アーキテクチャ/プロジェクトと Terraform モジュールを取り込みました。 3以降-

5 つのプロンプトを表示すると、通常はある程度の調整が行われます。 「これで十分だ」と感じたら、LLM に綿密な計画を立てるように指示します。私はこれを個人的に選択しますが、LLM にもこれを選択するよう依頼します。そして通常、ここで別のコンテキストまたは llm モデルに切り替えることになります。計画と全体的なアイデアをさらに洗練させるために、いくつかのラウンドロビンを行うこともできます。
プランが良さそうだと感じたら、エージェントにリッピングしてもらいます。可能であれば、仕事をサブエージェントに引き継ぐように伝えます。通常、この時点ではすべての編集を許可しますが、状況には常に注意を払っています。 LLM から完了の通知が来たら、ローカルにデプロイまたは実行します。試してみてください。
この時点ではすべてがうまくいっているように見えても、さらに多くのエージェントを立ち上げ始め、彼らをより敵対的にし始めます。彼らの目標は、コードの匂い、不適切な設計、重複/無効なコードなどを見つけることです。通常、エージェントの 1 つはテストを作成することのみを目的としており、テストをレビューするために別のエージェントを開始します。テストのレビューでは、価値の低いテスト、ゲッター/セッター/モックのテストに注意するように伝えることが常に重要です。私が望んでいるのは、問題を隠すことでカバレッジを増やしたり、障害を修正したりするだけのテストではなく、価値の高いテストです。私はいつも彼らに、幸せなパスだけでなく、考えられるすべてのエラーパスもテストしていることを確認するように言います。
問題を見つけ始める頃です。この時点で LLM を最も活用することになると思います。ここにはパターンの修正がたくさんあります。 LLM は、針に糸を通すために悪いコードを書くことを好みます。少なくとも、私の一般的なワークフローはそのようになります。しかし、何か機能するものができたら、レビュー、修正、テストを何回も繰り返します。時々、何らかの回帰が発生することがありますが、LLM はそれらも喜んで修正します。
コードベースが表示されているのを見ると、

形になり始め、パターンが形成され始めているので、物事を「完成」状態にしてリリースの準備を始めます。 CI/CD を改良するかもしれません。何らかのタイプの e2e スモークテストを作成するかもしれません。そして、さらにいくつかの手動レビューを行います。アプリをデプロイし、開発/デバッグ モードでローカルにインストールし、ほとんどの機能を実行します。
次に、LLM にこれを Google Chrome プラグインとして配布する手順を説明してもらいました。
このアイデアを考え、構築し、デプロイし、配布するのに 1 ～ 2 時間しかかからなかったと思います。
これをいじり続けます。まだFirefoxプラグインとしてリリースする必要があります。いつか、いくつかの追加機能 (近くで一番おいしい料理は何ですか?) を備えた Flutter アプリを作成すると思います。
人々がこれを利用して、私が訪問するレストランでどの料理を食べるべきかについてより良い情報を得ることができることを願っています。

## Original Extract

ByACommonThread Home Blog What Should I Eat
side-projects · ai-agents · aws
I didn’t know of any apps out there that crowdsource what the best dish at a restaurant is or a ranking of the best dishes. I’ve also never published browser plugins, I was interested in how that works. So I made this plugin.
This v1 version is for Chrome and Firefox (as of writing the draft of this blog post only the Chrome version is published).
This will probably be more useful as an android/ios app, so I’ll get around to that, but I figured I could fill in some details here about how it works and my workflow in building this.
Very simple, v1 only works in Google Maps, you select a restaurant on the map, click the plugin. If the restaurant has never been clicked before, we add it automatically, then you can input what you think the best dish is. If the restaurant has been clicked before, you will see other people’s votes and you can either enter something new or click one of the dishes displayed. The UI uses a “pill cloud” where the text size grows based on how many votes a dish has, so the heavy hitters bubble up visually.
There’s no authentication, no validation on dishes entered. This was partly to keep things simple, but it’s also a feature: there is zero friction to contribute. You don’t have to sign up for yet another service just to say the lasagna is good. Yeah, you can get duplicates. Yeah, users can clear their local data and vote multiple times on a dish.
But, I don’t want to deal with data privacy, plugin approvals, etc. So I chose to make it super simple and basic.
I had a tough time figuring out how exactly to manage restaurant entries. I want this to work in more than just google maps in the future, so I decided against using a google place id. Relying on Google’s IDs also makes you a bit of a hostage—if they change their internal ID system or you want to support Apple Maps or Yelp later, you’re starting from scratch.
Ultimately, I landed on using GPS coordinates. GPS coordinates will work well for mobile app functionality later, so that’s a plus. And I create a hash/uuid based on those gps coordinates, and these are deterministic based on the coordinates, so it’s easier to lookup in our single table dynamodb design. I’m also using Geohashing to bucket restaurants into ~5km squares (5-character prefixes). This lets the API query the database for just the nearby “bucket” instead of scanning the whole table, which keeps the DynamoDB costs at basically zero. Once I have the nearby results, I use the Haversine formula to calculate the exact distance in meters and sort them. (Here’s where I should note, I had/have no idea what the Haversine formula is, that’s what the LLM came up with)
To handle the “clustering” problem and GPS drift, I round the coordinates to 3 decimal places before hashing. This gives us about 110 meters of “fuzzy” snapping. It means even if you’re standing on the other side of the restaurant, it should still resolve to the same ID. It’s a little worse for locations where a lot of restaurants cluster close to each other in a tiny footprint, but it’s a trade-off I’m willing to make for platform agnosticism. In the future, I’ll add a way to select a specific restaurant if they’re clustered. At some point I’ll also have to think of a way to rename restaurants for when spots close and re-open as something new too.
The backend is also pretty simple. It’s all in AWS running as a serverless api. All compute is in AWS Lambda, Apigw handles requests, Dynamodb for the datastore. There’s no authentication, no user accounts, it’s a really simple API. It will scale up automatically to a point and it’s pretty much free to run.
Well, there are very few people using this right now, so currently it’s free (besides the hosted zone, but I run multiple things off of that and it’s only $.50 a month). If this thing even gets to the point of costing me money I’d consider it pretty successful. Once it rolls over from the free tier on my stack, it will scale linearly with the number of users.
There’s no monetization and I’d only consider monetizing this if my costs exceed my appetite. If it costs me $100+ a month, I’d try to find a simple, non-intrusive way to monetize to cover my costs. Right now, I don’t see that being in the cards so it’s not really a concern.
I wrote almost zero code for this personally, I had an LLM do most of it. Gemini, Claude, ChatGPT all had a hand. I think this one was mostly Gemini (which at the time of writing this, June 2026, I’d say it’s the worst tool of the bunch). I don’t stay completely hands off and vibe code, I give guidance, gate the llm at certain points, manually run certain verification steps/tools. It feels more like having a conversation with a Jr Engineer who is eager to please.
I start by explaining what I want to build and telling the LLM to help me build an AGENTS/CLAUDE/GEMINI markdown file. This part usually takes anywhere from 1 to 5 prompts. I’ll also read the markdown file myself a couple times and this is where I might manually make some tweaks.
Once the markdown file is completed, I’ll discuss architecture and code patterns that I want to see. I pulled in some reference architecture/projects as well as my terraform modules for the LLM to review. After 3-5 prompts, we’ll usually reach some alignment. When I feel like it’s “good enough” I tell the LLM to write up a thorough plan. I’ll pick this apart personally, but I also ask the LLM to pick this apart. And generally, this is where I’d switch to another context or llm model. Maybe even round robin a couple of them to further refine the plan and overall idea.
Once I feel like the plan looks good, I just let an agent rip. I tell them to pass of work to sub-agents if possible. I usually allow all edits at this point, but I keep an eye on things. Once the LLM tells me it’s done, I’ll deploy or run things locally. Take it for a spin.
Even if everything looks okay at this point, I start spinning up more agents and I start to make them more adversarial. Their goals are to find code smells, bad design, duplicate/dead code, etc. Typically one of the agents only purpose is to write tests and again, I’ll start another agent to review the tests. For the reviews on tests, it’s always important to tell them to be on the lookout for low value tests, testing of getters/setters/mocks. I want high value tests, not just tests that increase coverage and fix failures by hiding problems. I always tell them to also make sure we’re not only testing the happy path but all possible error paths as well.
This is around the time that I’ll start finding problems. I feel like at this point is where I’m going to use the LLM the most. There’s a lot of fixes to patterns here. The LLMs like to write bad code to thread the needle, at least, that’s how my general workflow goes. But then once there’s something functional I do many rounds of these reviews, fixes, testing. Sometimes there’s some regression, but the LLM is happy to fix those too.
When I see the codebase is really starting to take shape, patterns are forming, I’ll start getting things into a “finished” state and ready for release. Maybe spruce up ci/cd. Maybe create some type of e2e smoketest. And I’ll do some more manual reviews: deploy the app, install things locally in a dev/debug mode, and run through most of the functionality.
I then had the LLM walk me through distributing this as a Google Chrome plugin.
I think it only took me an hour or two to think of this idea, build it, deploy it, distribute it.
I’ll continue to fiddle with this. I still need to release it as a firefox plugin. I think I’ll make a flutter app with some extra functionality (What are the best dishes nearby?) at some point.
Hopefully people will use this so I can be better informed about what dish to get at any restaurant I visit.
