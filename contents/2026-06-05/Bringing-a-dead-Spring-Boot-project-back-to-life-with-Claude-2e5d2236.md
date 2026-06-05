---
source: "https://tomaytotomato.com/spring-data-solr-lazarus/"
hn_url: "https://news.ycombinator.com/item?id=48416959"
title: "Bringing a dead Spring Boot project back to life with Claude"
article_title: "Bringing a dead Spring Boot project back to life with Claude"
author: "tomaytotomato"
captured_at: "2026-06-05T21:36:28Z"
capture_tool: "hn-digest"
hn_id: 48416959
score: 1
comments: 0
posted_at: "2026-06-05T19:18:36Z"
tags:
  - hacker-news
  - translated
---

# Bringing a dead Spring Boot project back to life with Claude

- HN: [48416959](https://news.ycombinator.com/item?id=48416959)
- Source: [tomaytotomato.com](https://tomaytotomato.com/spring-data-solr-lazarus/)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T19:18:36Z

## Translation

タイトル: クロードとともに死んだ Spring Boot プロジェクトを生き返らせる
説明: Spring Data Solr は 2020 年に活発な開発を停止し、2023 年にアーカイブされました。3 年後、クロードとペアになり、屋根裏部屋から戻ってきました。

記事本文:
春
注目の
クロードとともに死んだ Spring Boot プロジェクトを生き返らせる
Spring Data Solr は 2020 年に活発な開発を停止し、2023 年にアーカイブされました。3 年後、クロードとのペアリングが行われ、屋根裏部屋から戻ってきました。
Github リンク: https://github.com/tomaytoTomato/spring-data-solr
Apache Solr を使用している、または使用したことがある場合は、これが世の中の検索エンジンの中で「最もクールな」検索エンジンではないことがわかるでしょう。ただし、これは現在も積極的に開発されており、非常に高速で耐久性があり、ドキュメントのクエリに簡単に使用できます。
新興のライバルである Elasticsearch と同じ Lucene インデックス エンジンを使用していることを知って驚く人もいます。
これまでのキャリアの中でいくつかのプロジェクトで Solr を使用してきたので、その強力な検索機能とファセット機能を楽しんでいます。本当の威力は、データのインデックス付け方法を形成する明確に定義されたスキーマと、その下の Lucene クエリ構文の 2 つの連携によって実現されます。これにより、さまざまなフィルターを使用してほぼあらゆる方向にデータをスライスできます。
ディーゼルエンジンを搭載した青い車が販売されています。
エンジンサイズは2000cc以下。
場所はコーンウォール。
価格でランク付けし、3 つのグループに分類します
- 低走行距離 < 80,000 マイル
- 中走行距離 120,000 マイル
- 走行距離が多い > 160,000 マイル 車の Solr 検索クエリの疑似コードの例
以前のプロジェクトでは、Spring Framework を使用した Java バックエンド サービスを通じてこの検索機能を提供していました。
Spring フレームワークと Spring Boot がどのように機能するかという核心には触れません。これは本質的に、さまざまなインフラストラクチャとの統合を可能にするベルトとブレースの依存関係管理フレームワークです。その後、REST、GraphQL、RPC、SOAP (まだ使用している人がいるでしょうか?)、CLI などを使用して、そのデータをクライアントに表現できます。
Sprin を使用すると、統合が大幅に容易になります

これらの要件に対応する g Boot のスターター ライブラリ。依存関係を追加するだけで、最小限の構成でそれらが接続されます。
ここで遊んでみてください - https://start.spring.io/
長い間、Solr のスターター ライブラリがありましたが、2020 年に積極的に開発されなくなり、最終的には 2023 年に Spring Attic に置かれる運命にありました。
ちょっとしたハックなプロジェクトと Spring Boot スターター ライブラリ開発への寄り道を経て、Spring Data Solr ライブラリを復活させることができるでしょうか。
開発者として、私たちは何かをハッキングするのが大好きです。ほとんどの場合、仕事中または休日に、ランダムなアイデアや作りたいものがあります。残念なことに、30代になると時間はより制限され貴重なものになりつつあります。
しかし、LLM は、こうした知的なかゆみを解消し、通常なら何時間もかかり、中途半端な Github プロジェクトと「もしもどうなるか?」という答えのない質問に終わるアイデアをテストする機会を提供するために現れました。
私の最近の知的かゆみは、地域の計画文書 (パブリック ドメイン) をわかりにくくするためにインデックスを付けて保存するための検索エンジンを作成することでした。メタデータとその他の興味深いデータを抽出してインデックスを作成し、簡単に検索できるようにします。
この背後にある動機は、英国の地方議会の計画 Web サイトの多くがひどい検索機能を備えており、住所や日付などに限定されていることが多いためです。マンチェスター市議会のウェブサイトの例を参照してください。
Solr を使用してから何年も経っていたので、最新バージョンを試してみる良い機会だと思いました。しかし、Spring Boot と統合するための最良の方法に関するドキュメントを探すことになったとき、この素晴らしいリポジトリが提示されました。すべてアーカイブされ、埃をかぶっていました。
まあ、大きなブロッカーではありませんが、SolrJ クライアントを使用して、独自の検索 API 呼び出しを処理するだけで済みます。

Spring Boot を使用した検索 API サービスもあれば、追加するだけで優れたリポジトリ抽象化とミドルウェア層を備えた Spring Data Solr ライブラリがないため、非常に不足していると感じました。
これにより、次に最適なポートが残りました。それは、SolrJ Java クライアントを使用して Solr へのクエリを実行することでした。これは、私自身のリポジトリとインターフェイスを手作業でローリングすることを意味します。
これは非常に簡単に実行でき、概念実証は機能しました。ただし、コードベースを見ると、多くの定型文があることに気づきます。
RESTコントローラー
│
▼
サービス層
│
▼
ハンドロールリポジトリ
│
▼
SolrJ クライアント ──────► Apache Solr
(手動クエリ構築、応答マッピング、ページネーションなし、並べ替えなし) Spring Data ミドルウェアを使用しない Direct Solr
ただし、Spring Data Solr ライブラリがまだ生きていて動作していれば、次のようになっているでしょう。
RESTコントローラー
│
▼
サービス層
│
▼
SolrRepository<PlanningDoc> (findByKeywords、findByAuthor)
│
▼
SolrTemplate ──────► Apache Solr
(自動構成、ページネーション、ソート、基準クエリ組み込み) Spring Data ミドルウェアと統合された Solr
そこで疑問が生じます。Claude に新しい Spring Data Solr スターター ライブラリを構築してもらうことはできないでしょうか?
最後に Solr を使用したのは、2019/20 年の Spring Boot 2.4.x でした。 Spring エコシステムではずっと昔のことです。
6 年が経過し、Spring Boot はバージョン 4.x.x になり、基礎となる Spring フレームワークはバージョン 7 になりました。リリース速度がすぐに遅くならないため、最新バージョン以外で Solr 用の新しい Spring Data スターター ライブラリを作成するのは無意味であるように思えました。
同時に、Solr はバージョン 10 のリリースで検索エンジンと同様に進化しました。
まあまあ、そこに固定しておきます

e バージョン。
Spring Data ライブラリにはすべて、Spring 開発者にとって非常に便利な共通の機能があります。
使用を開始するための便利なメソッドを提供する CrudRepository インターフェイス (findById、delete、update、create、findAll)
派生クエリメソッド: Spring Data は、メソッド名に基づいてどのクエリを構築するかを推測します (例: findByTitleContaining(String keyword))
@Query : 派生クエリ メソッドでは処理できない、より複雑なクエリ用。 @Query("タイトル:?0 AND 著者:?1")
Criteria API: プログラムでクエリを構築するための流暢な Java API
Criteria.where("title").contains("spring").and("price").between(10, 50) ページネーションと並べ替え: Pageable のサポートにより、100 を超えるデータセットには必須です。
自動構成: 依存関係を追加し、構成値を設定すると、リポジトリ経由で Solr にアクセスできるようになります。
したがって、ライブラリを復活させるこの実験では、これらすべての機能が存在し、動作する必要があります。
クロードやエージェントを PLAN.md ファイルで逃がす前に、その地域を偵察する価値はありました。
既存の Spring Data Solr ライブラリはどのような状態でしたか?
GitHub にはどのような未マージおよび未解決の問題がありましたか?
Spring フレームワークと Spring Boot ではどのような重要な API 変更が行われましたか?
Apache Solr の最新バージョンではどのような機能が利用できますか?
クロードはこの点で優れており、特に群エージェント パターンを使用して複数の異なる情報源を調査し、それを結び付けました。 swarm を試してみることをお勧めします。これは非常に便利でした。
そこから得られた洞察の例は次のとおりです。
元の Spring Data Solr プロジェクトの 8 つの未解決のバグ。修正するか、新しいアプローチで回避する必要があります。
コミュニティ フォークは存在しましたが、サポートされているのは Solr バージョン 9 未満のみです
SolrJ ライブラリはその中でいくつかのアーティファクトに分割されていました

最新バージョンと古いバージョンの比較 solr-solrj 、 solr-solrj-jetty 、 sollr-solrj-zookeeper
Springs のヘルスおよびメトリクス ツールは、v2.x.x から v4.x.x までに大幅に変更されました。
このデジタル考古学への自動化された進出は、v.001 の実装を軌道に乗せるために何が必要になるのかについての入門書を提供しました。また、古い Spring Data Solr ライブラリに影響を与えていたいくつかの問題が、Spring Data の最新バージョンでは回避できることがわかって安心しました。
最後のステップは、この新しいライブラリを組み立てる際のガイドとなるテンプレートまたはアーキタイプを見つけることでした。確かに、Claude は何千もの Java ライブラリ リポジトリについて訓練を受けており、Java ライブラリ リポジトリを構築する方法について意見やアイデアを持っていますが、私は Spring のやり方から逸脱したくありませんでした。そこで、ElasticSearch の成熟した実際のライブラリと、ericus20 が提供するこの比較的最新の Spring Boot スターター テンプレートを組み合わせて使用​​することにしました。
それをビルドする時が来て、次のような狡猾な PLAN.md が生成されました。
構築する機能 (必須):
1. マルチモジュール Maven プロジェクト — 自動構成、スターター、サンプル
2. SolrProperties バインディング spring.solr.* 構成
3. SolrAutoConfiguration による HttpJdkSolrClient (スタンドアロン) または CloudSolrClient (SolrCloud) の作成
4. SolrTemplate — コレクション対応の CRUD とクエリ実行のラッピング SolrJ
5. Criteria API — ネイティブ SolrQuery に変換する流暢なクエリ ビルダー
6. SolrRepository<T> — @EnableSolrRepositories を使用した Spring Data リポジトリの抽象化
7. Spring Data PartTree → SolrQueryCreator による派生クエリメソッド
8.生のSolrクエリ文字列に対する@Queryアノテーションのサポート
9. SolrHealthIndicator — 管理エンドポイントフォールバックを使用した収集 ping
10. クエリレイヤー上の強調表示、ファセット、およびカーソルベースのディープページング
11. テストコンテナの統合

実際の Solr 9 および Solr 10 に対するイオン テスト
12. Spring Boot 自動構成 — XML ゼロ、アノテーションのみのセットアップ 次に、実装時にいくつかのエージェントをリリースし、いくつかの段階でチェックインする時期が来ました。
私は、好みのスタイルで単体テストと統合テストを作成する、TDD エージェントと開発者のアプローチを選択しました。その後、エージェントにフォローアップして実装を行います。
ほとんどの機能はローカルまたは Github でレビューできるほど小規模でした。
短期間のうちに、最初のバージョンを試す準備が整いました。自動構成、 SolrTemplate 、 Criteria API、部分更新、ページネーション、Spring Data リポジトリのサポート、および Testcontainers を備えています。
0:00
/ 0:48
1×
ソースの可視化
ローカルでアプリを試してみたところ、それは良いように思えましたが、簡単に把握できる外部のレビュアーがいませんでした。
もし私が何人かの査読者をシミュレートして、偏見やニッチな興味分野を持ったライブラリをチェックできたらどうなるでしょうか。クロードはそれをうまく表現し、シミュレートできるでしょうか。
Spring コミュニティの 2 人の非常に重要な人物をシミュレートし、Claude にコード レビュー ファイルを出力させることにしました。
ロッド・ジョンソン ;春の創造者
「これは、Spring Data Commons の拡張ポイントと Boot の自動構成規則を正しく使用する、よく構造化され、明確に実装された Spring Boot スターターです。0.1.0 アーティファクトとして期待していたよりも大幅に完成しています。」
「評決: このままでは Maven Central の準備ができていません。インジェクションの脆弱性、変異する `count()`、壊れた `CrudRepository` コントラクトを修正し、インターフェースから `getSolrClient()` を削除します。そうすれば、これは正当な 0.1.0 になります。」
ジョシュ・ロングSpring 開発者およびエバンジェリスト
「何ということでしょう。どこから始めればよいでしょうか。このリポジトリへのリンクを渡されました。その名前を読んだ後の私の最初の反応は、思わずニヤリとしてしまいました。」

g Data Solr Lazarus。大胆さ。ドラマ。大好きだよ。」
「これについてツイートしてもいいですか？はい - インジェクションの問題が解決されれば。デモしてみませんか？絶対に。」
これにより、機能の構築とライブラリの修正を開始するための非常に役立つフィードバックが得られたため、クロードとの 2 時間のハッキング セッションよりも本格的な作業になる可能性があります。
さらに数日間調整を行った後、Spring Data Solr ライブラリのオリジナルの作成者に連絡しました。彼らは構造についていくつかのヒントを与えてくれました
Elastic または Solr モジュールをガイドとして使用するのではなく、テンプレートとリポジトリ レイヤーを MongoDB または Cassandra に合わせて調整し、他の Spring Data 実装から自由に描画して、エコシステムにネイティブに感じられるようにします。
オリジナルがアーカイブされた理由について彼らから聞いたのも興味深いものでした
Solr はメジャー アップグレードごとに互換性を壊すという習慣があるため、Spring Data ポートフォリオの残りの部分とともにモジュールに必要な注意を払うことができなくなりました。
これにより、私たちが正しい軌道に乗っているという素晴らしい確認が得られました。あと少し調整するだけです。
次の数週間は、Claude とコードをレビューし、Github の問題を作成し、それらの修正と機能をレビューするための PR を作成するという反復的なアプローチで構成されていました。
図書館に備え付けられている、役立つもの、または煩わしさを防ぐもの:
リ

[切り捨てられた]

## Original Extract

Spring Data Solr stopped active development in 2020 and was archived in 2023. Three years later and some pairing with Claude, its now back from the attic.

spring
Featured
Bringing a dead Spring Boot project back to life with Claude
Spring Data Solr stopped active development in 2020 and was archived in 2023. Three years later and some pairing with Claude, its now back from the attic.
Github Link: https://github.com/tomaytotomato/spring-data-solr
If you are using or have used Apache Solr , you will know its not the "coolest" of search engines out there. However it is still being actively developed, is very fast, durable and is easy to use for querying documents.
Some people are even surprised to find out that it uses the same Lucene indexing engine as its newer kid-on-the-block rival, Elasticsearch.
Having used Solr for several projects in my career, I enjoy its powerful search and faceting capabilities. The real power comes from two things working together: a well-defined schema that shapes how your data is indexed, and the Lucene query syntax underneath that lets you slice it in almost any direction with a variety of filters.
Find blue cars for sale with diesel engines.
Engine size no larger than 2000cc.
Location Cornwall.
Rank them in price and also facet them into three groups
- low mileage < 80,000 miles
- medium mileage 120,000 milee
- high mileage > 160,000 miles Example pseudo code for a Solr search query for cars
For those previous projects I was providing this search capability through a Java backend service using the Spring Framework.
Without going into the nitty gritty of how Spring framework and Spring Boot works; it essentially is a belt and braces dependency management framework that lets you integrate with a whole variety of infrastructure. You can then express that data to a client with REST, GraphQL, RPC, SOAP (who still uses that?), CLI etc.
A lot of the ease of integration comes with Spring Boot's starter libraries for these requirements. Just add your dependencies and it wires them up with minimal config.
You can have a play around here - https://start.spring.io/
For a long time there was a starter library for Solr but it stopped being actively developed in 2020 and eventually was doomed to the Spring Attic in 2023.
After a little hacky project and detour into Spring Boot starter library development, could the Spring Data Solr library be brought back into life.
As developers we love to hack away at stuff. Most days I have random ideas or things I want to build at work or in my days off. Unfortunately time is becoming a more restricted and precious commodity in my 30s.
LLMs have turned up though to offer you those chances to scratch those intellectual itches and test out some ideas that would normally take hours and hours and result in a half done Github project and the unanswered question of, "what if?".
My most recent intellectual itch was to make a search engine to index and store to obscure local planning documents (public domain). Extracting the metadata and other interesting data from them into an index to be easily searched.
The motivation behind this was that a lot of local council planning websites in the UK have terrible search features and are often limited to things like address or date. See this example of Manchester City council's website .
It has been years since I used Solr, so it seemed like a good opportunity to try out its latest version. However when it came to search for documentation on the best way to integrate it with Spring Boot, I was presented with this nice repository , all archived and gathering dust.
Fair enough, not a major blocker, I can just use the SolrJ client and handle my own search API calls it came time to throw together a search API service using Spring Boot, it felt very lacking in that there was no Spring Data Solr library I could just add and have a nice repository abstraction and middleware layer.
This left me with my next best port of call which was to use the SolrJ Java client to execute my queries to Solr. This would mean hand rolling my own repositories and interfaces.
This was easy enough to do and the proof of concept worked. When looking at the codebase though you notice that there is a lot of boilerplate.
REST Controller
│
▼
Service Layer
│
▼
Hand-rolled Repository
│
▼
SolrJ Client ──────► Apache Solr
(manual query building,response mapping, no pagination, no sorting) Direct Solr with no Spring Data middleware
However if the Spring Data Solr library was still alive and kicking it would have looked more like this:
REST Controller
│
▼
Service Layer
│
▼
SolrRepository<PlanningDoc> (findByKeywords, findByAuthor)
│
▼
SolrTemplate ──────► Apache Solr
(auto-configured, pagination, sorting, criteria queries built-in) Integrated Solr with Spring Data middleware
So it begs the question, couldn't we get Claude to build us a new Spring Data Solr starter library?
The last time I used Solr was with Spring Boot 2.4.x which was back in 2019/20; which is a long long time ago in the Spring ecosystem.
With six years passing, Spring boot has now reached version 4.x.x and the underlying Spring framework is onto version 7 . With its release velocity not slowing anytime soon, it seemed pointless to entertain making a new Spring Data starter library for Solr in anything other than the latest versions.
At the same time Solr had moved on as well as a search engine with its release of version 10 .
Fair enough lets just keep it anchored at those versions.
Spring Data libraries all have features in common that are very handy for a Spring developer to use:
CrudRepository interfaces that provide useful methods for getting started with (findById, delete, update, create, findAll)
Derived query methods: Spring Data infers what query to construct based on the method name e.g findByTitleContaining(String keyword)
@Query : for more involved queries that can't be handled by derived query methods. @Query("title:?0 AND author:?1")
Criteria API: fluent Java API for building queries programmatically
Criteria.where("title").contains("spring").and("price").between(10, 50) Pagination and sorting: with Pageable support, a must for any dataset greater than 100.
Auto-configuration: add the dependency, set the config values and you have access to Solr via repositories.
So for this experiment to revive a library, we would need all these pieces of functionality present and working.
Before we let Claude or any agents loose with a PLAN.md file it was worth doing some reconnaissance of the area.
What was the state of the existing Spring Data Solr library?
What un-merged and unresolved issues were there on GitHub?
What significant API changes have happened in Spring framework and Spring boot?
What features are available in the newest version of Apache Solr?
Claude was excellent at this, especially using a swarm agentic pattern to research multiple disparate sources of information and tie it together. I would recommend playing around with swarms, this one was quite useful.
Some examples of the insights it gave were:
8 unresolved bugs in the original Spring Data Solr project; they should be fixed or avoided by new approaches
A community fork existed but only supported Solr versions < 9
SolrJ library was split into several artifacts in its latest version compared to older ones solr-solrj , solr-solrj-jetty , sollr-solrj-zookeeper
Springs health and metrics tooling had changed considerably since v2.x.x to v4.x.x
This automated foraying into digital archaelogy gave a primer in what would be involved in getting a v.001 implementation off the ground. It was also a relief to see that some issues that affected the older Spring Data Solr library could be bypassed in the newest versions of Spring Data.
The final step was to find a template or archetype to be our guide when assembling this new library. For certain Claude has been trained on thousands of Java library repos and has opinions and ideas on how to structure one, but I didn't want to stray away from the Spring way of doing things. So I opted for a combination of using a mature real library from ElasticSearch and this relatively up to date Spring Boot starter template provided by ericus20.
It was time to then build it and a cunning PLAN.md was hatched, which went roughly like this.
Features to build (essential):
1. Multi-module Maven project — autoconfigure, starter, sample
2. SolrProperties binding spring.solr.* config
3. SolrAutoConfiguration creating HttpJdkSolrClient (standalone) or CloudSolrClient (SolrCloud)
4. SolrTemplate — collection-aware CRUD and query execution wrapping SolrJ
5. Criteria API — fluent query builder converting to native SolrQuery
6. SolrRepository<T> — Spring Data repository abstraction with @EnableSolrRepositories
7. Derived query methods via Spring Data PartTree → SolrQueryCreator
8. @Query annotation support for raw Solr query strings
9. SolrHealthIndicator — collection ping with admin endpoint fallback
10. Highlighting, faceting, and cursor-based deep paging on top of the query layer
11. Testcontainers integration tests against real Solr 9 and Solr 10
12. Spring Boot auto-configuration — zero XML, annotation-only setup Then it was time to release several agents at the implementation and check in at several stages.
I opted for a TDD agent and developer approach with writing unit and integration tests in a style I like. Then following up with an agent to do the implementation.
Most features were small enough that I could then review them locally or on Github.
Within a short period of time, a first version was ready to try out; with auto-configuration, SolrTemplate , Criteria API, partial update, pagination, Spring Data repository support, and Testcontainers.
0:00
/ 0:48
1×
gource visualisation
After playing around with the app locally it looked good to me, but I didn't have any external reviewers I could grab a hold of easily.
What if I could simulate some reviewers to check through the library with a bias or niche area of interest, could Claude express and simulate that well.
I decided to simulate two very important people in the Spring community and get Claude to output code review files.
Rod Johnson ; the creator of Spring
"This is a well-structured, cleanly implemented Spring Boot starter that correctly uses Spring Data Commons' extension points and Boot's auto-configuration conventions. It is substantially more complete than I expected for a 0.1.0 artefact."
"Verdict: Not ready for Maven Central as-is. Fix the injection vulnerability, the mutating `count()`, the broken `CrudRepository` contract, and remove `getSolrClient()` from the interface. Then this is a legitimate 0.1.0."
Josh Long ; Spring dev and evangelist
"Oh my goodness, where do I begin. I was handed a link to this repository and my first reaction — after reading the name — was an involuntary grin. 'Spring Data Solr Lazarus.' The audacity. The drama. I love it."
"Would I tweet about this? Yes — once the injection issue is fixed. Would I demo it? Absolutely."
This gave some really useful feedback to start building out features and fixing the library so it could become more serious than a two hour hack session with Claude.
After a few more days of tweaking I then reached out to the original creators of the Spring Data Solr library. They gave some pointers on the structure
Align the template and repository layer with MongoDB or Cassandra rather than using the Elastic or Solr modules as a guide, and draw freely from the other Spring Data implementations to make it feel native to the ecosystem.
It was also interesting to hear from them on why the original was archived
Solr's habit of breaking compatibility with each major upgrade made it impossible to give the module the attention it needed alongside the rest of the Spring Data portfolio.
This gave a nice confirmation that we were on the right tracks, just a few more tweaks.
The following weeks consisted of an iterative approach, reviewing code with Claude, creating Github issues and then creating PRs to review those fixes and features.
Some thing furnished into the library which would be of use or prevent annoyance:
Re

[truncated]
