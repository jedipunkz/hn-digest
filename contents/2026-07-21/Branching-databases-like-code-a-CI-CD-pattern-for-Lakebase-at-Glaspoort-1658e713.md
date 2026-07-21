---
source: "https://www.databricks.com/blog/branching-databases-code-cicd-pattern-lakebase-production-glaspoort"
hn_url: "https://news.ycombinator.com/item?id=48992846"
title: "Branching databases like code: a CI/CD pattern for Lakebase at Glaspoort"
article_title: "Branching databases like code: a CI/CD pattern for Lakebase, in production at Glaspoort | Databricks Blog"
author: "ryantsuji"
captured_at: "2026-07-21T14:56:49Z"
capture_tool: "hn-digest"
hn_id: 48992846
score: 1
comments: 0
posted_at: "2026-07-21T14:29:55Z"
tags:
  - hacker-news
  - translated
---

# Branching databases like code: a CI/CD pattern for Lakebase at Glaspoort

- HN: [48992846](https://news.ycombinator.com/item?id=48992846)
- Source: [www.databricks.com](https://www.databricks.com/blog/branching-databases-code-cicd-pattern-lakebase-production-glaspoort)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T14:29:55Z

## Translation

タイトル: コードのようなデータベースの分岐: グラスポートの Lakebase の CI/CD パターン
記事のタイトル: コードのようなデータベースの分岐: Glaspoort で実稼働中の Lakebase の CI/CD パターン |データブリックのブログ
説明: Glaspoort が Databricks Lakebase で OLTP CI/CD を実行する方法: 本番環境からのブランチ、一時的な PR ごとのデータベース、唯一の信頼できるソースとしての移行。

記事本文:
コードのようなデータベースの分岐: Glaspoort で実稼働中の Lakebase の CI/CD パターン | Databricks ブログ メイン コンテンツにスキップ ログイン Databricks がアプリ開発者向けに発見する理由
パートナー パートナーの概要 Databricks パートナー エコシステムを探索する
パートナー スポットライト 注目のパートナーの発表
パートナー プログラム メリット、レベル、パートナーになる方法を確認する
AWS、Azure、GCP 上のクラウド プロバイダー Databricks
パートナーを探す ニーズに合った Databricks パートナーを見つけてください
パートナー ソリューション カスタム業界および移行ソリューションを見つける
製品 Databricks プラットフォーム プラットフォームの概要 データ、分析、AI のための統合プラットフォーム
データ エンジニアリング ETL とバッチおよびストリーミング データのオーケストレーション
ビジネス ユーザー向け AI Assistant Agentic コワーカー
データ ウェアハウス SQL 分析用のサーバーレス データ ウェアハウス
アプリケーション開発 安全なデータと AI アプリを迅速に構築
データ アプリと AI エージェント用のデータベース Postgres
人工知能 ML および GenAI アプリケーションを構築してデプロイする
ガバナンス すべてのデータ、分析、AI 資産に対する統合ガバナンス
ビジネス インテリジェンス 実世界のデータに対するインテリジェントな分析
AI 時代に向けて構築されたセキュリティ オープン エージェント SIEM
Databricks に組み込まれた顧客データ プラットフォーム Agentic CDP
共有 データ、分析、AI のためのオープンデータ共有
統合とデータ マーケットプレイス データ、分析、AI のオープン マーケットプレイス
IDE 統合 お気に入りの IDE で Lakehouse を構築
Partner Connect Databricks エコシステムを発見して統合する
価格設定 Databricks の価格設定 製品の価格設定、DBU などを調べる
コスト計算ツール あらゆるクラウド上のコンピューティング コストを見積もります
オープンソース オープンソース テクノロジー プラットフォームの背後にあるイノベーションについて詳しく知る
産業向けソリューション Databricks 電気通信
異業種ソリューション AI エージェント
移住

n および展開データの移行
ソリューション アクセラレータ アクセラレータを探索する 重要な結果に向けてより迅速に移行する
リソース 学習トレーニング ニーズに合わせたカリキュラムを見つける
Databricks アカデミー Databricks 学習プラットフォームにサインインする
認証取得による認知と差別化
無料版 プロフェッショナルなデータと AI ツールを無料で学習
University Alliance Databricks を教えたいですか?その方法をご覧ください。
ブログとポッドキャスト Databricks ブログ ニュース、製品発表などをご覧ください
AI ブログ AI の研究とエンジニアリングの取り組みを詳しく見る
Data Brew ポッドキャスト データについて話しましょう!
Champions of Data + AI ポッドキャスト イノベーションを推進するデータ リーダーからの洞察
安心と信頼 安心と信頼
無視できなかった問題
湖底の分岐、60秒以内
グラスポートの設計: 常に生産から分岐
フォーク: 2 つのプロモーション モデル
サイドバー: 途中で思いついたもの
これによりグラスポールトが得られたもの
無視できなかった問題
湖底の分岐、60秒以内
グラスポートの設計: 常に生産から分岐
フォーク: 2 つのプロモーション モデル
サイドバー: 途中で思いついたもの
これによりグラスポールトが得られたもの
コードのようなデータベースの分岐: Glaspoort で実稼働中の Lakebase の CI/CD パターン
Glaspoort は、すべての環境を本番環境から分岐し、一時的な PR ごとのデータベースをスピンアップし、移行を唯一の信頼できる情報源として扱うことにより、アプリケーション コードと同じ厳密さでデータベースの変更をどのように出荷しているか。
ハディ・ファーハット、ギデオン・スパイリングス、リカルド・デ・フリース、レイモン・フェルドマン著
Glaspoort は、古い環境を更新すると削除と削除が強制される「親からのリセットの罠」を回避するために、開発と受け入れの両方が互いに積み重ねられるのではなく本番環境から直接分岐する、Lakebase 分岐セットアップをどのように設計したか。

その下にあるものをすべて再構築します。
PR ごとの CI/CD フロー: すべてのプル リクエストが本番環境からコピーされた独自の新しい使い捨て Lakebase ブランチを取得する方法、実際の環境に何かが触れる前に移行がどのように再生され、ライブ アプリ イメージに対してテストされるか、および移行自体 (データベースではなく) が信頼できる情報源として扱われる理由。
2 つのプロモーション モデル間のトレードオフ: CI が通過したらすぐにマージするか、PR が受け入れられて完全にプロモーションされた後にのみマージする。この投稿では、Glaspoort が速度優先のアプローチを選択した理由と、それを安全に保つために追加した保護手段 (スタックの再検証と危機パイプライン) について説明しています。
無視できなかった問題
Glaspoort は、オランダで光ファイバー インフラストラクチャを構築および運用しています。すべてはファイバー接続の数を増やすことを中心に展開しており、データ チームは長い間、その目標をサポートする BI レポートの作成に日々を費やしていましたが、各レポートの背後にある疑問は、レポートが完成した時点ではすでに時代遅れになっていました。その結果、単発的なレポートが大量に発生し、ユーザーはフォローアップの質問を受け付ける場所がなくなりました。
そこで依存関係を解消しました。次のレポートを出荷する代わりに、プロジェクト マネージャーがプロジェクトのチャンスがどこにあるかを直接確認できるカスタム フロントエンド アプリケーションを構築しました。新機能は内部にあります。Databricks 製品をアプリの構成要素として直接使用しています。 Genie は、データとチャットして簡単な分析を開始します。 AI/BI ダッシュボード: インサイトとセルフサービス分析用。 Agent Bricks による自動化されたワークフローは、プロジェクト上で何かが目立った瞬間にプロジェクト マネージャーに警告します。また、アプリケーションのトランザクション データについては、Databricks OLTP データベースである Lakebase を使用します。
この組み合わせにより、最近まで存在していた 2 つの世界が 1 つになります。

rt: 分析データがトランザクション データと出会う分析環境と運用フロントエンド。データ チームは現在、1 回限りのレポートではなく、Genie スペースとメタデータに時間を費やしています。しかし、これらはいずれも、CI/CD、データ品質テスト、コードとしてのインフラストラクチャ、データ ガバナンスなどの本格的な基盤がなければ、高速性を維持することはできません。その基盤の 1 つの部分には最も慎重な設計が行われ、それがこのストーリーの残りの部分、つまりアプリの背後にあるデータベースに変更を送信する方法についての部分です。
このアプリケーションの背後には、Databricks Lakebase データベースが配置されています。これはサーバーレス Postgres OLTP であり、レイクハウスにボルトで固定されるのではなく、レイクハウスの隣で実行されます。データ フローは説明が簡単で、見た目よりも操作が興味深いことがわかりました。
レイクハウスから厳選されたデータは Lakebase 本番ブランチに同期され、読み取り専用のアプリケーション スキーマに配置されます。
アプリケーションは独自の状態を同じブランチ上の別のスキーマに書き戻すため、レイクハウスから読み取られたデータとアプリによって書き込まれたデータは衝突することなく並列して存在します。
その上で、開発、受け入れ、本番という 3 つの論理環境を実行します。プル リクエスト、CI、ゲート プロモーションを通じて、アプリケーション コードに変更を送信するのと同じ方法で、このデータベースに変更を送信します。
最後のポイントは興味深い質問です。 OLTP データベースがアプリケーション コードと同じ厳密さに値すると判断した瞬間、次の 1 つの難しい質問に答える必要があります。
人々がお互いに踏みにじることなく、そもそもテストを意味のあるものにする新鮮なデータを失わずに、本番環境に似たデータベースに対してすべての PR テストを実行するにはどうすればよいでしょうか?
ここでの障害モードは、チーム間でデータベースを共有したことのある人なら誰でもよく知っているものです。孤立して通過する PR

彼らが一緒に着地したときのk。開発環境と受け入れ環境は、実際の運用環境から静かに遠ざかっています。そして、誰も望んでいないリフレッシュ日。クリーンなデータを取り戻すには、環境を破壊し、すべての接続文字列を再配線し、すべての許可を手動で再適用する必要があります。
この投稿では、そのほとんどをどのように回避したか、そして現在も積極的に議論している 1 つの決定について説明します。
湖底の分岐、60秒以内
Lakebase ブランチをまだ使用したことがない場合は、この投稿に必要なメンタル モデルはこれだけです。
Lakebase ブランチは、親から離れたコピーオンライト Postgres ブランチです。作成してもデータはコピーされません。安価かつ即座に状態をフォークし、ブランチは書き込み時にのみ親から分岐します。各ブランチは完全に分離されており、独自のエンドポイントと独自のデータを持ちます。数秒で作成でき、すぐに捨てることができます。
それが git のように聞こえる場合は、それがポイントです。以下のすべてを整理するアナロジーは単純です。git の機能ブランチは、Lakebase のデータベース ブランチにマップされます。 PR は独自のコード ブランチと独自のデータベース ブランチを取得し、両方に対してテストを実行し、結果が良好であれば本番環境に昇格します。
git ブランチとは決定的に異なるため、次のセクションに引き継ぐ制約が 1 つあります。ブランチをその親からリセットするには、まずそのブランチ自体の子を削除する必要があります。親を、それに依存するブランチの下からリセットすることはできません。
これは、ホワイトボード上に環境を描画する方法を反映しているため、ほとんどのチームが最初に到達するデザインです。
単一の長期にわたる開発が本番環境から分岐します。
受け入れ分岐開発。
機能は開発から分岐します。
これは明確な階層構造です。ルートに本番、次に開発、次に受け入れ、その下に機能がぶら下がっています。

打つ。また、2 つの予測可能な障害モードも導入されています。1 つはブランチのドリフトで、もう 1 つはドリフトの修正にはコストがかかるため、チームは適用を中止します。
まず、開発と受け入れが本番環境からずれています。本番環境では、新しい同期データと実際のアプリケーションの書き込みを受け取り続けます。長く存続する開発ブランチと受け入れブランチはそうではありません。 1 ～ 2 スプリント以内に、出荷先のデータベースとは似ていないデータベースに対してテストを行うことになります。
明らかな解決策は、開発環境を運用環境から更新することです。ここで、前のセクションの制約が負担となります。開発を親からリセットするには、最初に開発の子を削除する必要があります。このトポロジでは、承認と開発からぶら下がっているすべての機能ブランチを意味します。したがって、「新しいデータをください」というルーチンはカスケードに変わります。承認とすべての機能ブランチを削除し、開発をリセットし、環境を再作成し、古いブランチを指すすべての接続文字列を再接続し、すべての Postgres グラントを再適用します。グラントは削除したばかりのブランチに存在するためです。
これらの手順はどれもそれ自体は難しいものではありません。これらを定期的に組み合わせることで、チーム全体が更新を密かに回避できる確実な方法になります。つまり、全員が古くなったデータに対するテストに戻ることになります。それが元々の問題でした。単純なトポロジは、ただひどい午後を過ごすだけではありません。環境を健全に保つ衛生状態が損なわれます。
グラスポートの設計: 常に生産から分岐
この修正は、トポロジについての考え方を 1 行変更するだけで、非常に大きな影響を及ぼします。存続期間の長い環境ブランチはすべて、別の環境の子ではなく、運用環境の子になります。
開発と受け入れはどちらも本番環境から直接分岐したものです。これらは、製造中は互いに並べて配置され、互いに積み重ねられることはありません。 Neither is the

もう一方の親なので、一方を更新してももう一方の削除が強制されることはありません。
そのたった 1 つの変化が罠を解除します。開発または承認がドリフトした場合は、スプリントごとに 1 回程度、または新しいデータが必要なときはいつでも、本番環境 (今日の UI 操作) からリセットします。 dev または accept の下には何もぶら下がっていないため、最初に削除する子も、チーム全体で再スレッドする接続文字列も、許可の再適用のマラソンもありません。リセットは低コストなので実際に実行するので、環境は正直なままになります。リセットで変更が欠落した場合は、次の移行リプレイでその変更が再度適用されます。
アニメーション: PR ごとのライフサイクル。公開されたブログで再生されます。このドキュメント内の静止フレームとして表示されます。
日々の開発者ループは、安定したトポロジの上に一時的なブランチを重ねます。
開発者が一時的な PR (TTL を 1 時間に設定) を開きます。
CI は、新たな一時的な pr-xxxx ブランチを本番環境から切り離します。すべてを本番環境から分岐させ、データベースをマージして戻すことはありません。これらの PR ごとのブランチは使い捨てです。PR が終了した瞬間にブランチがアーカイブされます。これにより、Lakebase のデフォルト制限であるプロジェクトあたり 10 個の未アーカイブ ブランチを維持できます。
CI は、その新しいブランチに対して移行を再実行します。 git diff チェックにより、移行が再実行されるかどうかが判断されます。

[切り捨てられた]

## Original Extract

How Glaspoort runs OLTP CI/CD on Databricks Lakebase: branch from production, ephemeral per-PR databases, migrations as the single source of truth.

Branching databases like code: a CI/CD pattern for Lakebase, in production at Glaspoort | Databricks Blog Skip to main content Login Why Databricks Discover For App Developers
Partners Partner Overview Explore the Databricks partner ecosystem
Partner Spotlight Featured partner announcements
Partner Program Explore benefits, tiers and how to become a partner
Cloud Providers Databricks on AWS, Azure and GCP
Find a Partner Discover Databricks partners for your needs
Partner Solutions Find custom industry and migration solutions
Product Databricks Platform Platform Overview A unified platform for data, analytics and AI
Data Engineering ETL and orchestration for batch and streaming data
AI Assistant Agentic coworker for business users
Data Warehousing Serverless data warehouse for SQL analytics
Application Development Quickly build secure data and AI apps
Database Postgres for data apps and AI agents
Artificial Intelligence Build and deploy ML and GenAI applications
Governance Unified governance for all data, analytics and AI assets
Business Intelligence Intelligent analytics for real-world data
Security Open agentic SIEM built for the AI era
Customer Data Platform Agentic CDP embedded in Databricks
Sharing Open data sharing for data, analytics and AI
Integrations and Data Marketplace Open marketplace for data, analytics and AI
IDE Integrations Build on the Lakehouse in your favorite IDE
Partner Connect Discover and integrate with the Databricks ecosystem
Pricing Databricks Pricing Explore product pricing, DBUs and more
Cost Calculator Estimate your compute costs on any cloud
Open Source Open Source Technologies Learn more about the innovations behind the platform
Solutions Databricks for Industries Telecommunications
Cross Industry Solutions AI Agents
Migration & Deployment Data Migration
Solution Accelerators Explore Accelerators Move faster toward outcomes that matter
Resources Learning Training Discover curriculum tailored to your needs
Databricks Academy Sign in to the Databricks learning platform
Certification Gain recognition and differentiation
Free Edition Learn professional Data and AI tools for free
University Alliance Want to teach Databricks? See how.
Blog and Podcasts Databricks Blog Explore news, product announcements, and more
AI Blog Explore our AI research and engineering work
Data Brew Podcast Let’s talk data!
Champions of Data + AI Podcast Insights from data leaders powering innovation
Security and Trust Security and Trust
The problem we couldn't ignore
Lakebase branching, in 60 seconds
The Glaspoort design: always branch from production
The fork: two promotion models
Sidebars: the things that came up along the way
What this unlocked for Glaspoort
The problem we couldn't ignore
Lakebase branching, in 60 seconds
The Glaspoort design: always branch from production
The fork: two promotion models
Sidebars: the things that came up along the way
What this unlocked for Glaspoort
Branching databases like code: a CI/CD pattern for Lakebase, in production at Glaspoort
How Glaspoort ships database changes with the same rigor as application code by branching every environment from production, spinning up ephemeral per-PR databases, and treating migrations as the single source of truth.
by Hadi Farhat , Gideon Spierings , Ricardo de Vries and Raymon Veldman
How Glaspoort designed its Lakebase branching setup, with development and acceptance both branched directly from production instead of stacked on top of each other, to avoid the "reset-from-parent trap" where refreshing a stale environment forces you to delete and rebuild everything underneath it.
The per-PR CI/CD flow: how every pull request gets its own fresh, disposable Lakebase branch copied from production, how migrations get replayed and tested against a live app image before anything touches a real environment, and why the migrations themselves (not the databases) are treated as the source of truth.
The tradeoff between two promotion models: merging as soon as CI passes versus merging only after a PR is fully promoted through acceptance. The post explains why Glaspoort picked the velocity first approach, and the safeguards they added (stack revalidation and a crisis pipeline) to keep it safe.
The problem we couldn't ignore
Glaspoort builds and operates fiber infrastructure in the Netherlands. Everything revolves around growing the number of fiber connections, and for a long time the data team spent its days building BI reports in support of that goal, while the question behind each report was already outdated by the time the report was finished. The result was a sprawl of one-off reports, and users who had nowhere to take their follow-up questions.
So we broke the dependency. Instead of shipping the next report, we built a custom front-end application in which project managers see directly where the opportunities for their projects lie. What is new sits under the hood: we use Databricks products directly as building blocks in the app. Genie, to chat with the data and spin up quick analyses. AI/BI Dashboards, for insights and self-service analytics. Automated workflows with Agent Bricks, which alert project managers the moment something stands out on their projects and Lakebase, the Databricks OLTP database , for the application's transactional data.
That combination brings together two worlds that until recently lived apart: the analytical environment and an operational front-end, where analytical data meets transactional data. The data team now spends its time on Genie spaces and metadata instead of one-off reports. But none of this stays fast without a serious foundation underneath: CI/CD, data-quality testing, Infrastructure as Code, and data governance. One piece of that foundation took the most careful design, and it is the piece the rest of this story is about: how we ship changes to the database behind the app.
Behind that application sits a Databricks Lakebase database . It is serverless Postgres OLTP, running next to the lakehouse rather than bolted onto it. The data flow is simple to describe and, as we found out, more interesting to operate than it looks:
Curated data from the lakehouse is synced into a Lakebase production branch, where it lands in a read-only application schema.
The application writes its own state back into a separate schema on that same branch, so the data read from the lakehouse and the data written by the app live side by side without colliding.
On top of that we run three logical environments: development, acceptance, and production. We ship changes to this database the same way we ship changes to application code, through pull requests, CI, and gated promotion.
That last point is where the interesting question lives. The moment you decide an OLTP database deserves the same rigor as application code, you have to answer one hard question:
How do we let every PR test against a database that looks like production, without people stepping on each other, and without losing the fresh data that makes the test meaningful in the first place?
The failure modes here are familiar to anyone who has shared a database across a team. PRs that pass in isolation and break when they land together. Dev and acceptance environments that have quietly drifted away from what production actually looks like. And the refresh day nobody wants, where getting clean data back means tearing environments down, re-wiring every connection string, and re-applying every grant by hand.
This post is about how we avoided most of that, and the one decision we are still actively debating.
Lakebase branching, in 60 seconds
If you have not used Lakebase branches yet, here is the only mental model you need for this post.
A Lakebase branch is a copy-on-write Postgres branch off a parent . Creating one does not copy the data; it forks the state cheaply and instantly, and the branch only diverges from its parent as you write to it. Each branch is fully isolated, with its own endpoint and its own data. You can create one in seconds and throw it away just as fast.
If that sounds like git, that is the point. The analogy that organizes everything below is simple: a feature branch in git maps to a database branch in Lakebase. A PR gets its own code branch and its own database branch, tests run against both, and when the work is good it gets promoted toward production.
There is one constraint to carry into the next section, because it is a critical distinction from git branching. To reset a branch from its parent, you first have to delete that branch's own children. A parent cannot be reset out from under the branches that depend on it.
Here is the design most teams reach for first, because it mirrors the way we draw environments on a whiteboard:
A single long-lived development branch off production.
An acceptance branch off development.
Feature branches off development.
It is a clean hierarchy: production at the root, then dev, then acceptance and features hanging beneath it. It also introduces two predictable failure modes: the branches drift, and the fix for drift is expensive enough that teams stop applying it.
First, dev and acceptance drift from production. Production keeps receiving fresh synced data and real application writes; the long-lived dev and acceptance branches do not. Within a sprint or two you are testing against a database that no longer resembles the one you are shipping to.
The obvious fix is to refresh dev from production, and this is where the constraint from the last section turns into a tax. To reset development from its parent, you must first delete development's children, which in this topology means acceptance and every feature branch hanging off dev. So a routine "give me fresh data" turns into a cascade: delete acceptance and all feature branches, reset development, re-create the environments, re-wire every connection string that pointed at the old branches, and re-apply every Postgres grant, because grants live on the branch you just deleted.
None of those steps is hard on its own. Together, on a recurring basis, they are a reliable way to make the whole team quietly avoid refreshing, which means everyone goes back to testing against stale, drifted data. That was the original problem. The naive topology does not just cost you a bad afternoon; it discourages the hygiene that keeps the environment honest.
The Glaspoort design: always branch from production
The fix is a one-line change in how you think about the topology, and it has outsized effects: every long-lived environment branch is a child of production, not of another environment.
Development and acceptance are both branches taken directly from production. They sit beside each other under production, not stacked on top of one another. Neither is the other's parent, so refreshing one never forces you to delete the other.
That single change defuses the trap. When dev or acceptance drifts, we reset it from production (a UI operation today), roughly once a sprint or whenever we want fresh data. Because nothing hangs below dev or acceptance, there are no children to delete first, no connection strings to re-thread across the team, and no grant re-apply marathon. The reset is cheap, so we actually do it, so the environments stay honest. If a reset happens to miss a change, the next migration replay simply applies it again.
Animated: the per-PR lifecycle. Plays on the published blog; shown as a still frame inside this doc.
The day-to-day developer loop layers ephemeral branches on top of that stable topology:
A developer opens an ephemeral PR (TTL set to 1 hour).
CI cuts a fresh, ephemeral pr-xxxx branch from production. We branch everything from production and never merge databases back. These per-PR branches are disposable: we archive them the moment the PR closes, which keeps us under Lakebase’s default limit of 10 unarchived branches per project.
CI replays the migrations against that fresh branch. A git diff check decides whether a migration re

[truncated]
