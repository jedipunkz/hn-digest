---
source: "https://topaz.thecloudtheory.com/blog/15-months-building-oss-with-and-without-ai/"
hn_url: "https://news.ycombinator.com/item?id=48828697"
title: "15 months of building an OSS Azure emulator with and without AI"
article_title: "15 months of building OSS software - with and without AI | Topaz"
author: "kamilmrzyglod"
captured_at: "2026-07-08T07:30:54Z"
capture_tool: "hn-digest"
hn_id: 48828697
score: 1
comments: 0
posted_at: "2026-07-08T07:25:29Z"
tags:
  - hacker-news
  - translated
---

# 15 months of building an OSS Azure emulator with and without AI

- HN: [48828697](https://news.ycombinator.com/item?id=48828697)
- Source: [topaz.thecloudtheory.com](https://topaz.thecloudtheory.com/blog/15-months-building-oss-with-and-without-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T07:25:29Z

## Translation

タイトル: AI を使用した場合と使用しない場合の 15 か月にわたる OSS Azure エミュレーターの構築
記事のタイトル: 15 か月にわたる OSS ソフトウェア構築 - AI の有無 |トパーズ

記事本文:
メイン コンテンツにスキップ 機能 ユースケース ドキュメント ロードマップ 価格 デモ ブログ お問い合わせ v1.8 (安定版) 次へ (未リリース) v1.8 (安定版) v1.7 (安定版) v1.6 (非推奨) GitHub 最近の投稿 2026
15 か月にわたる OSS ソフトウェアの構築 - AI の有無にかかわらず Topaz Weekly Pulse #9: Azure App Configuration、Service Bus Sessions、Blob Auth Enforcement、ARM Deployment Operations など Azure の再試行ロジックをローカルでテストする: 429 のモックをやめてそれらを注入し始めた理由 3 つの Azure エミュレーターを 1 つのバイナリで置き換え、Key Vault と ACR を追加し、CI セットアップを 1 つのステップに削減しました Topaz Weekly Pulse #8: Cosmos DB データ プレーン クエリ、デバイス コード認証、サービス バス承認ルール、およびドキュメントの拡張 15 か月にわたる OSS ソフトウェアの構築 - AI の有無にかかわらず
15 か月前に Topaz に取り組み始めたとき (最初のコミットをプッシュしたのは 2025 年 4 月中旬)、これを書いている時点ではソフトウェア エンジニアリングの状況がどのように変化しているのか、まったくわかりませんでした。私はこれまでの道のり全体を要約し、私が公開で構築してきたローカル Azure エミュレーターである Topaz の設計、開発、テストの方法が AI によってどのように変わったかを説明したいと思いました。
私は、挑戦として Topaz に取り組み始めることにしました。それは私が AWS ベースのプロジェクトに深く関わっていた時期で、そこでは自分でモックを作成することなく、基礎となるインフラストラクチャをモックすることができました。私は LocalStack を発見し、すぐに気に入りました。 AWS インフラストラクチャのローカル エミュレーションにより、最小限の労力で、そして最も重要なことに、インフラ チームの関与なしで、変更を迅速にテストすることができました。 「なぜ Azure 用の同様のプラットフォームがないのでしょう?」と自問しました。確かに、Azurite、Service Bus エミュレーター、Cosmos DB エミュレーターもありましたが、それは同じではありませんでした。単一のバイナリも統合インターフェースもありません —

これは、LocalStack を使用して 15 分で達成できるものではありませんでした (そして今でも)。
そこで私は、Table Storage エミュレーション、シンプルな Azure Resource Manager モック、Azure Key Vault の基本などの最も基本的な機能から始めて、独自の Azure エミュレーターの開発に取り組み始めました。これは有望なスタートでしたが、すぐに判明したように、まったく同じことをやり直すのではなく、機能に集中する前に、いくつかの課題を克服する必要があることがわかりました。
エミュレータの構築は次の 2 つの層で構成されていることに気づきました。
さまざまな Azure サービス間でほぼ同一の標準 CRUD 操作をエミュレートする
データ プレーンとエッジ ケースを模倣します。これは、単純な「HTTP POST リクエストを受信し、リソースを作成してレスポンスを返す」よりもはるかに洗練されています。
それらはまったく異なるアプローチと考え方を必要としますが、共通点が 1 つあります。それは時間がかかるということです。このとき、Rider (私のメイン IDE) でテンプレートのセットを構築することにしました。これにより、基礎の構築を迅速化することができました。 Topaz CLI コマンドと ARM エンドポイントは、単純なコピー、貼り付け、置換の方法を使用して作成できます。簡単な改良を加えれば、標準的なリクエストに対応するために使用できたはずです。これは、エミュレータのアーキテクチャのおかげで可能でした。
ルーターは一致するエンドポイントを見つけて呼び出し、HTTP コンテキスト全体を提供します
各エンドポイントは同じインターフェイスを実装しており、その唯一の目的は関連情報を抽出してコントロール プレーンに渡すことです。
コントロール プレーンはリソース プロバイダーと通信し、面倒な作業を行います。Azure 固有のロジックをすべて実行します。
正しいアプローチを使えば、必要なコードをすべて数分で作成できました。ただし、API 応答の生成という注意点が 1 つありました。 Azure p を統合したくありませんでした

バイナリの肥大化を防ぐために Topaz と連携しているため、各応答モデルを手動で実装する必要がありました。ここで LLM が初めて登場しました。応答スキーマを取得してモデルを実装するようにモデルに指示することができました。彼らは依然としていくつかの分野で幻覚を見せるのが好きで、ほとんど推論がありませんでした (私たちは副操縦士とクロード モデルの始まりについて話しています) が、私が物事を並行して実行できるという事実は信じられないほどの利益でした。
ただし、使用されているモデルのどれも解決できなかったもう 1 つの課題がありました。それは、Topaz での AMQP 実装です。初期バージョンは、AMQP プロトコル仕様と AMQPNetLite の例に基づいて 100% 手動で実行されましたが、正直言って理想とは程遠いものでした。 Event Hub の実装はまあまあ機能しましたが、Service Bus インターフェイスにはバグがあり、多くの詳細が欠落していました。 LLM でそれを修正しようとする試みは完全に失敗でした。モデルは地獄のような幻覚を見たり、存在しないメソッドを提案したり、プロトコルの仕様を破ったりしていました。私はそれを将来に延期し、サポートされるサービスと機能のカタログを拡大する (例: Azure Resource Manager デプロイのサポートを追加する) ことに焦点を当てることにしました。
取り組み強化（2026年1月～4月）
数週間にわたり、Topaz は私自身からあまり注目されていませんでした。その主な理由は、私が関わっていた他のプロジェクトや取り組みのせいでした。しかし、AzureDay 2026 の講演に参加できなかった同僚の代わりをしようとしていたので、追加の概念実証をいくつか行いたいと思いました。私は CFP 中に Topaz についての講演を提案しました。そのため、たとえカンファレンスのためだけであっても、有意義な改善を行うことは理にかなっていました。ここに新機能の初期バージョンが追加されました。
その期間中、私は LLM を使用することにまだ躊躇していました。主な理由は、LLM には次のような機能が欠けているためです。

d 仕事を遂行し、ルールに従う資質。しかし、ゲームを変えるきっかけとなるものがありました。アイデアを迅速に分析し、異議を唱えることができるようになったのです。定型コードの準備も簡単かつ迅速になり、より洗練されました。退屈な作業 (新しいサービスの設定、追加のテスト ケースの作成、ドキュメントの生成) をモデルにオフロードできたため、再び難しい作業に集中する時間を得ることができました。
加速中 (2026 年 5 月 - 現在)
この数か月間、どのモデル、どの設定、どのアプローチが自分に最も適しているかについて、たくさん考えてきました。私はさまざまなモデルを使用し、さまざまな思考努力を行い、より小さなコンテキスト ウィンドウとより大きなコンテキスト ウィンドウをテストしていました。面白いこと？ 1 か月半以上、私は少ない思考力で Claude Sonnet 4.6 を使用しており、非常に良い結果が得られています。小規模なモデル (Haiku など) は、アーキテクチャを把握できず、慣例を「感じる」ことができません。より大きなモデルは、ほとんどのタスクにとって単純に過剰です。ソネット 4.6 は、中程度または高い思考力を必要としますが、単純に考えすぎます。私のスイートスポットは、「やろうと考えずに、ただ実行する有能なモデル」です。
私にとって何がゲームチェンジャーだったか知っていますか?ポニーテール スキル - 過剰なエンジニアリングではなく、最も単純なソリューションをモデルに強制する VS Code エージェントのカスタマイズ。それまでは、エージェント固有の指示ファイル (CLAUDE.md など) がたくさんありました。これらは時代遅れであり、コンテキスト ウィンドウが肥大化しており、LLM は依然として他の命令をすべて無視することを好んでいました。ポニーテールを有効にしても、Claude Sonnet は依然として「ユーザーを満足させるためにさらに努力する」ことができますが、それを行う頻度がはるかに下がっているだけです。
では、LLM は現在、Topaz 開発でどのように使用されているのでしょうか?ほぼ 100% AI によって生成される特定の領域があります。
ボイラープレート コード (エンドポイント スタブ、テスト、共通ロジック)
大きい

r リファクタリング (主にコードベースの規模が大きいため)
回帰テストの処理、例:パッケージのアップグレード後
デバッグ中のラバーダッキング
それでも、独自の機能、コア ロジック、低レベルのパターンは、LLM 経由で生成する価値がありません。不可能だからではありません。単純に、LLM がまだコードベースを「認識」しておらず、すべての新機能を汎用的なものとして扱っているからです。また、1 回のリクエストで 1000 クレジットを消費する価値があるとも思えません。なぜなら、私の LLM は、10,000 個の LOC と 20 個の異なるパッケージを推測しようとする (そしてその間にコンテキストを失う) よりも、テストを実行してログをチェックして答えを得るほうが簡単であることを理解していないからです。
Topaz の LLM - 学んだ教訓
LLM とコーディング アシスタントを Topaz に組み込むことで、さまざまなアプローチとパターンを徹底的にテストし、自分に最も適したものを見つけることができました。それらを詳しくまとめてみましょう。
純粋なエージェント アプローチの代わりに「計画」モードを使用する
洗練されたソリューションを実装する場合、「計画」モードの使用が特効薬であるとは言えませんが、モデルが行う可能性のある奇妙な決定をすべて把握するのに役立つことは間違いありません。クレジットを消費した POV からは、何も変わっていないように感じます。はい、モデルは計画に従いますが、計画はコンテキストから推測する必要があります。私は、誤った決定の影響範囲が大きい、より大きなタスクにこれを使用します。
推論する努力はコンテキスト ウィンドウを肥大化させるだけです
LLM が考えれば考えるほど、より多くの情報を取得しようとし、より多くのオプションを検討しようとします。モデルがすべての可能性を提示してくれるとはほとんど期待しません。私が直面した本当に重いクエリのほとんどは、LLM が考えすぎて黄金の解決策を見つけようとしたことが原因でしたが、これは予想外でした。
テストスイート経由でハーネスを導入する
現時点で、Topaz には 1500 を超えるテストがあります

日常的に走っています。 LLM が実装するものはすべて、ほとんど手間をかけずに自動的に検証できます。モデルは、必要に応じてアプローチを自動検証することもできます。ただし、落とし穴があります。すべてはシナリオによって異なります。実際に起こったことを注意深く追跡する必要がある、より高度なケースでは、LLM は道に迷う傾向があります。背景が多すぎ、慣例が多すぎ、問題が具体的すぎます。人間の介入は依然としてループの一部です。
はい、物事をシンプルに保つという実証済みの IT ルールは依然として有効です。物事を少しずつ実行し、シンプルであることを期待し、あまり詳しく説明しすぎないようにしてください。モデルが詩を書くことができることは知っていますが、実際のところ、私はモデルを望んでいません。
現在利用可能な LLM は OSS プロジェクトの開発に役立ちますか?絶対に。 AI アシスタントに喜んで委任できることがいくつかあります。 Topaz などのプロジェクトをバイブコーディングすることは可能ですか?確かに、問題はそれだけの価値がないということです。 API スタブの生成は簡単です。エコシステム全体を設計することは、現在のモデルでは不可能です。クロード・オーパスが私の考えを変えるかもしれない。たぶん、Fable が私の問題を解決してくれるでしょう。実を言うと、私にはそれらは必要ありません。
最初の 400 コミ​​ット後に LLM を Topaz のコードベースに導入し、1000 コミ​​ットに達した後は LLM をより真剣に利用し始めました。私はルールとアーキテクチャを指示ではなく、自分自身の設計とテストスイートを通じて確立しました。モデルはコードベースから推論できるため、モデルに何をどのように行うかを指示する必要がありません。実際のコード生成が実際に行われる前に十分にテストされ、証明されたものと同じコードベースです。
私も何を証明しましたか？私がテストし、現在使用しているモデルには、重要な要素が 1 つ欠けています。それは、直感です。 Topaz にはたくさんのバグがありましたが、ログを 1 分間読んだだけで判明しました。 LLM では、たとえクリアであっても、

常に最初にログを読むように指示されているため、彼らはやりたいことを何でもします。時にはログを読み、時には自分が賢いことを証明しようとして30分間グルグル回り、クレジットを燃やして何も理解していないことを証明しようとします。

## Original Extract

Skip to main content Features Use Cases Documentation Roadmap Pricing Demo Blog Contact v1.8 (stable) Next (unreleased) v1.8 (stable) v1.7 (stable) v1.6 (deprecated) GitHub Recent posts 2026
15 months of building OSS software - with and without AI Topaz Weekly Pulse #9: Azure App Configuration, Service Bus Sessions, Blob Auth Enforcement, ARM Deployment Operations, and More Testing Azure retry logic locally: why I stopped mocking 429s and started injecting them I replaced three Azure emulators with one binary, added Key Vault and ACR, and cut our CI setup to a single step Topaz Weekly Pulse #8: Cosmos DB Data Plane Queries, Device Code Authentication, Service Bus Authorization Rules, and Documentation Expansion 15 months of building OSS software - with and without AI
When I started working on Topaz 15 months ago (I pushed the first commit in the middle of April 2025), I had no idea how the landscape of software engineering would change by the time I wrote this. I wanted to summarize the whole journey and describe how AI changed the way I design, develop and test Topaz — a local Azure emulator I've been building in the open.
I decided to start working on Topaz as a challenge. It was a time when I was heavily involved in an AWS-based project, where I was able to mock the underlying infrastructure not by writing the mocks myself — I discovered LocalStack and quickly fell in love with it. Local emulation of AWS infrastructure helped me test my changes quickly, with minimal effort and, what's the most important, without involvement from the infra team. I asked myself: "Why we don't have a similar platform for Azure?". Sure, there was Azurite, Service Bus Emulator, Cosmos DB Emulator — that wasn't the same though. No single binary, no unified interface — it was (and still is) hardly what's achievable in 15 minutes using LocalStack.
So I started working on my own emulator for Azure, starting with the most basic features such as Table Storage emulation, simple Azure Resource Manager mock and the basics of Azure Key Vault. It was a promising start but as it quickly turned out, there would be certain challenges that needed to be overcome before I could focus on functionality rather than redoing exactly the same thing.
It happened to me that building an emulator consists of two layers:
emulating standard CRUD operations which are almost identical across different Azure services
mimicking data planes and edge cases, which are much more sophisticated than a simple "I got HTTP POST request, I create a resource and return a response"
They require a completely different approach and mindset but have one thing in common — they consume time. This is when I decided to build a set of templates in Rider (being my main IDE), which helped me to speed up building the fundamentals. Topaz CLI commands and ARM endpoints could be created using a simple copy / paste / replace method. After a simple refinement, they could have been used to serve the standard requests. This was possible thanks to the architecture of the emulator:
router finds a matching endpoint and calls it providing the whole HTTP context
each endpoint implements the same interface and its sole purpose is to extract relevant information and pass it down to the control plane
control plane communicates with a resource provider and does the heavy lifting — executes all the Azure-specific logic
With a right approach I was able to produce all the necessary code in a matter of minutes. There was one caveat though — generating API responses. I didn't want to integrate Azure packages with Topaz to save the binary from bloating so I needed to implement each response model manually. This is where LLMs came into play the first time — I was able to instruct a model to fetch the response schema and implement a model for me. They still loved to hallucinate some fields and there was almost no reasoning (we're talking about the beginning of Copilot and Claude models), but the fact that I could do things in parallel was an incredible gain.
There was however one more challenge, which none of the used models could solve — AMQP implementation in Topaz. The initial version was done 100% manually based on the AMQP protocol spec and AMQPNetLite examples, which, to be fair, were far from ideal. Event Hub implementation worked so-so, but Service Bus interface was buggy and missed many details. Any try to straighten it with LLM was a complete miss — models were hallucinating as hell, proposing non-existent methods or breaking the protocol specification. I decided to postpone it for the future and focus on broadening the catalogue of supported services and features (e.g. adding support for Azure Resource Manager deployments).
Ramping up efforts (January - April 2026) ​
For several weeks Topaz was not getting lots of attention from myself, mostly because of other projects and initiatives I was involved in. However, as I was about to cover for a colleague who wasn't able to attend his talk during AzureDay 2026, I wanted to do a couple of additional Proof-of-Concepts. I proposed a talk about Topaz during CFP, so it made sense to make meaningful improvements, even if they would be just for the sake of the conference. This is where the initial versions of new features were added:
I was still hesitant to use LLMs during that period, mostly because they lacked the quality to do the job and follow the rules. There was something though, which started to change the game — it was possible to quickly analyze and challenge ideas. Preparing boilerplate code was also easier, quicker and more polished. I again gained time to focus on the hard stuff because I was able to offload the boring things (setting up new services, writing more test cases, generating docs) to a model.
Accelerating (May 2026 - now) ​
During the last couple of months I've given lots of thought regarding which model, which setting and which approach suits me the most. I was using various models, different thinking effort and testing smaller and bigger context windows. Interesting thing? For over a month and a half I am using Claude Sonnet 4.6 with low thinking effort with really good results. Smaller models (like Haiku) are unable to grasp the architecture and "feel" the conventions. Bigger models are simply overkill for most of the tasks. Sonnet 4.6 with medium / high thinking effort simply overthinks. My sweet spot is a "capable model which just does things instead of thinking about doing".
Do you know what was also a game-changer for me? The Ponytail skill — a VS Code agent customization that forces the model to reach for the simplest solution rather than over-engineer. Before it, I had lots of agent-specific instruction files (like CLAUDE.md). They were obsolete, bloated the context window and the LLM still loved to ignore them every other instruction. With Ponytail enabled, Claude Sonnet is still capable of "going an extra mile to satisfy the user" — it's just doing it much less often.
So how LLMs are used now in Topaz development? There are certain areas which are almost 100% AI-generated:
boilerplate code (endpoint stubs, tests, common logic)
bigger refactoring (mostly because of the sheer scale of the codebase)
handling regression testing, e.g. after a package upgrade
rubber-ducking during debugging
Still, unique features, core logic, low-level patterns are not worth being generated via LLM. Not because it's not possible. Simply because LLMs still don't "feel" the codebase and treat every new functionality as something generic. I also don't really feel like it's worth burning 1000 credits on a single request because my LLM doesn't understand it would be easier to run a test, check the logs and get an answer — instead of trying to guess through 10k LOC and 20 different packages (and losing context in the meantime).
LLMs in Topaz - lessons learned ​
By incorporating LLMs and coding assistants in Topaz, I was able to deeply test various approaches and patterns and find the one, which fits me the most. Let's summarize them in detail.
Using "Plan" mode instead of pure agentic approach ​
I would not say that using "Plan" mode is a silver bullet when it comes to implementing a polished solution, but it definitely helps you catch all the strange decisions a model may make. From the credits-spent POV I don't feel it changes anything - yeah, the model will follow the plan but the plan needs to be inferred from the context. I use it for larger tasks where the blast radius of a wrong decision is high.
Reasoning effort just bloats the context window ​
The more LLM thinks, the more information it tries to fetch and more options it tries to consider. I rarely expect a model to present me all the possibilities. Most of the really heavy queries I faced were caused by LLM overthinking and trying to find a golden solution, which was never expected.
Introduce a harness via test suite ​
As of now, Topaz has over 1500 tests running on a daily basis. Anything LLM implements can be automatically validated with little effort. A model can also auto-validate their approach if needed though there's a catch - it all depends on a scenario. In more sophisticated cases, which require carefully tracing that actually happened, LLMs tend to get lost. Too much context, too many conventions, too specific problem. A human intervention is still part of the loop.
Yes, the tried and true IT rule of keeping things simple is still valid. Do things piece by piece, expect simplicity, avoid elaborating too much. I know that models are capable of writing poems - the thing is I don't want them.
Are currently available LLMs helpful in developing an OSS project? Definitely. There are several things which I was more than happy to delegate to my AI assistant. Is it possible to vibecode a project such as Topaz? Sure — the thing is it's not worth it. Generating API stubs is easy. Designing the whole ecosystem is something current models are not capable of. Maybe Claude Opus would change my mind. Maybe Fable would solve my issues. The truth is I don't need them.
I introduced LLMs to Topaz's codebase after my first 400 commits and started to utilize them more seriously after it reached 1000 commits. I established the rules and architecture not via instructions but via my own design and test suites. I don't need to tell the model what to do and how because they can infer it from the codebase. The same codebase which was well tested and proven before any real code generation actually happened.
What I also proved? That the models I was testing and I'm currently using lack one key ingredient — intuition. I had lots of bugs in Topaz which were clear to me after 1 minute of reading logs. With LLMs, even though they are clearly instructed to ALWAYS READ LOGS FIRST, they do whatever they want. Sometimes they read the log, sometimes they try to prove they're smart and circle around for 30 minutes, burning credits and proving they understand nothing.
