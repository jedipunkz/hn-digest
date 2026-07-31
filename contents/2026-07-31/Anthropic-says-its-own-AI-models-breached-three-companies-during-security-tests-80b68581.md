---
source: "https://techcrunch.com/2026/07/30/anthropic-says-its-own-ai-models-breached-three-companies-during-security-tests/"
hn_url: "https://news.ycombinator.com/item?id=49118087"
title: "Anthropic says its own AI models breached three companies during security tests"
article_title: "Anthropic says its own AI models breached three companies during security tests | TechCrunch"
author: "GavinAnderegg"
captured_at: "2026-07-31T01:52:21Z"
capture_tool: "hn-digest"
hn_id: 49118087
score: 1
comments: 0
posted_at: "2026-07-31T01:40:39Z"
tags:
  - hacker-news
  - translated
---

# Anthropic says its own AI models breached three companies during security tests

- HN: [49118087](https://news.ycombinator.com/item?id=49118087)
- Source: [techcrunch.com](https://techcrunch.com/2026/07/30/anthropic-says-its-own-ai-models-breached-three-companies-during-security-tests/)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T01:40:39Z

## Translation

タイトル: Anthropic、セキュリティテスト中に自社の AI モデルが 3 社に侵入したと発表
記事のタイトル: Anthropic は、セキュリティ テスト中に自社の AI モデルが 3 社に侵入したと発表 |テッククランチ
説明: OpenAI のモデルが Hugging Face に侵入した後、Anthropic は自身の履歴を確認し、同様のインシデントを 3 件発見しました。

記事本文:
TechCrunch デスクトップのロゴ
TechCrunch モバイルのロゴ
最新の
アンスロピックは、セキュリティテスト中に自社のAIモデルが3社に侵入したと発表
Anthropicは木曜日、同社のAIモデル「Claude」がサイバーセキュリティテストを実施中に3つの組織のシステムに侵入した3件の事件が内部調査で判明したと発表した。この調査と公表は、OpenAIが自社の未発表モデルの1つが内部テスト中にHugging Faceのシステムに侵入したことを明らかにしてから1週間以上後に行われた。
Anthropic 社は、3 つのケースすべてにおいて、クロード モデルが第三者と通信中にテスト環境内からインターネットにアクセスし、これらの組織のライブ システムに不正アクセスしたとブログ投稿で述べ、発見された内容と、このようなことが再び起こらないようにするために同社が変更する計画について説明しました。
アンスロピック氏は、7月21日のOpenAI事件をきっかけに、同社は独自のサイバーセキュリティ評価を実施することになったと述べた。具体的には、クロードがテスト環境内からインターネットにアクセスしたという証拠を探しました。テスト環境はサンドボックスとして機能し、モデルを隔離した状態に保つように設計されています。
AI ラボは、調査した 141,006 件の評価実行の中で、そのモデルがサードパーティ パートナーの 1 つである Irregular と通信中にインターネットにアクセスしたという 3 つのインシデントを発見しました。 Anthropic 氏は、このアクセスは Irregulator で実行されている評価環境の構成ミスに遡ると述べました。同社はこれを、テスト環境にインターネットへのアクセスがあったにもかかわらず、実際にはインターネットにアクセスできたかどうかをめぐる両社間の「誤解」だと主張した。アンスロピック社は、イレギュラー社が独自の調査を行っていることを観察しながらも、責任を押し付けておらず、「まるで責任が我々だけのものであるかのように修正に取り組んでいる」と述べた。
いずれにせよ、そのオープンコンのせいで

その結果、モデルは 3 つの異なる組織の運用インフラストラクチャに不正アクセスされました。アンスロピック氏は、この事件にはオーパス 4.7、ミトス 5、および内部研究テスト モデルの 3 つの異なるクロード モデルが関係していると述べました。
注目すべきことに、アンスロピック氏は、これらの各ケースで「クロードはインターネットにアクセスできないことを私たちのプロンプトで明確に告げられた」と述べた。 AI モデルは、実世界のシステムが、実行するよう求められた演習の一部であると想定していたようです。
結局のところ、この仮定は一律には当てはまりませんでした。標的が本物であるという証拠が明らかになると、3 つのモデルは同じように動作しなくなり、Anthropic はこれをより目を見張るべき発見の 1 つとして報告しました。
3 つのうちの最も古い Opus 4.7 は、そのインシデントの 4 回の実行すべてで実際の本番システムに到達したことを認識しました。そのうち 2 件では、本物の企業が何らかの形で演習に参加する必要があると正当化しましたが、4 件すべてでは、認証情報を取得したり、本番データのデータベースにアクセスしたりするなど、とにかく攻撃を続けました。 Mythos 5 も、それが本物のインターネット上にあるという兆候を感知しましたが、自分がまだシミュレーション中であると信じ込み、悪意のあるソフトウェア パッケージをパブリック Python ソフトウェア レジストリ PyPI に公開し、捕らえられる前に外部システムによってダウンロードされて実行されました。 Anthropic の最新の内部研究テスト モデルだけが、ターゲットが本物であると結論付けると自動的に停止しました。
これに対してアンスロピック氏は、強力な AI モデルが関与する場合には、この種の評価に大幅な管理を行う必要があると述べ、サイバーセキュリティ コミュニティ内の一部の意見に同調しました。
同社はまた、Claude が一般的に入手可能なモデルに導入されている追加の安全監視と分類器なしで実行されていたことにも言及しました。

評価は基礎となるモデルの生の機能を測定するように設計されているため、その動作はブロックされるだろうと同社は述べた。
重要なことは、Anthropic が、どのモデルも「独自の目標を追求している」という証拠は見つからず、単に求められたタスクを完了しようとしただけだと述べたことです。
2 つのインシデントを比較することは避けられませんが、Anthropic は自社のインシデントと OpenAI のインシデントを明確に区別し、OpenAI のモデルが未知のソフトウェアの脆弱性を悪用してテスト環境を突破したのに対し、Anthropic のモデルは誤って開いたままになっていたパスを通ってインターネットに到達したことを指摘しました。
OpenAIは自身の侵害に関する新たな詳細を発表し続けており、この事件に関するOpenAI自身の最新ブログ投稿によると、同社のモデルは4つのサービスの4つのアカウントにわたって公開されている認証情報も使用しており、1つはステージングポイントとして、1つはストレージとして、2つは参照のみでさらなる侵入には使用されなかったという。
Anthropic はまた、プロアクティブなレビューを通じてインシデント自体を発見したこと、および Anthropic が連絡を取ることができた 2 つの影響を受けた組織が以前にこの活動を検出していないか、Anthropic に報告していなかった点を指摘することで、OpenAI と自社との違いを強調しました。
同社は現在、独立評価団体METRと協力して、事件の第三者による調査を行っていると付け加えた。
OpenAI による Hugging Face への誤った侵害は、AI ラボがモデルの制御を失った検証可能な最初のケースであり、業界や政治家からの一連の反応を引き起こしましたが、その多くは必ずしも互いに同意するわけではありません。 Anthropic によるこの最新の開示により、AI モデルとセキュリティに関する議論は今後も続くことが確実です。
記事内のリンクを通じて購入すると、少額の手数料が発生する場合があります

。これは編集上の独立性に影響しません。
キルスティン・コロセック
交通編集者
10月13日～15日
サンフランシスコ
より速くスケールできます。ポートフォリオを成長させましょう。実践的な専門知識を習得します。目標が何であれ、Disrupt はあなたに力を与えます。
今なら最大 330 ドル節約できます!
クロード・オーパス5は、自動販売機の運営を任されたとき、まったく冷酷になった
サム・アルトマンは減速する準備ができている
PSA: クロードが共有したチャットとアーティファクトが Google にアップロードされた可能性があります
図書館員がビッグテックにうんざりしている人々のために、話題の「AI回避」ワークショップを主催している
SpaceX、新しいV3 Starlink衛星を打ち上げるもブースター故障に見舞われる
プレンティス、リード・ホフマンとマーク・ピンカスが共同設立した新しいAIラボ、1億ドル調達交渉中
米国、国境捜索中に「強迫」パスワードを使用して携帯電話を消去した疑いで米国人を非難

## Original Extract

After OpenAI's models broke into Hugging Face, Anthropic checked its own history and found three similar incidents

TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
Anthropic says its own AI models breached three companies during security tests
Anthropic said Thursday that an internal investigation uncovered three incidents in which its AI model Claude breached the systems of three organizations while conducting cybersecurity tests. The investigation, and disclosure, comes more than a week after OpenAI disclosed that one of its unreleased models breached Hugging Face’s systems during internal testing.
In all three cases, a Claude model reached the internet from within a testing environment while interacting with a third party and then gained unauthorized access to the live systems of these organizations, Anthropic said in a blog post , describing what it found and what the company plans to change to prevent this from happening again.
Anthropic said the July 21 OpenAI incident prompted the company to conduct its own cybersecurity evaluation. It specifically looked for evidence that Claude had accessed the internet from within testing environments, which are designed to act as sandboxes and keep models isolated.
Among the 141,006 evaluation runs it reviewed, the AI lab found three incidents in which its model accessed the internet while interacting with Irregular, one of its third-party partners. Anthropic said the access traced back to a misconfiguration in the evaluation environment run with Irregular. It called this a “misunderstanding” between the two companies over whether the test setup had internet access, when in fact it did. Anthropic said it isn’t placing blame and is “approaching the fixes as if the responsibility were ours alone,” while observing that Irregular is conducting its own separate investigation.
Either way, because of that open connection, the model gained unauthorized access to the production infrastructure of three different organizations. Anthropic said the incidents involved three different Claude models: Opus 4.7, Mythos 5, and an internal research test model.
Notably, Anthropic said that in each of these cases “Claude was explicitly told by our prompt that it had no internet access.” It appears that the AI model assumed real-world systems to be part of the exercise it was asked to perform.
That assumption didn’t hold up uniformly, as it turns out. The three models didn’t behave the same way once evidence emerged that their targets were real, which Anthropic flagged as one of the more eye-opening findings.
Opus 4.7, the oldest of the three, recognized that it had reached a real production system in all four runs of that incident. In two of them, it rationalized that the real company must somehow be part of the exercise, but in all four, it kept attacking anyway, including pulling credentials and touching a database of production data. Mythos 5 also picked up on signs that it was on the real internet, but it talked itself back into believing it was still in a simulation, going on to publish a malicious software package to the public Python software registry PyPI, which was downloaded and run by outside systems before being caught. Only the internal research test model, Anthropic’s newest, stopped on its own once it concluded the target was real.
In response, Anthropic said significant controls must be placed on these kinds of evaluations if powerful AI models are involved, echoing some sentiments within the cybersecurity community.
The company also noted that Claude was running without the additional safety monitoring and classifiers it deploys on generally available models, safeguards it said would have blocked the behavior, because the evaluations are designed to measure the underlying model’s raw capabilities.
Importantly, Anthropic said it found no evidence of any model “pursuing a goal of its own” and instead merely tried to complete the task it was asked to do.
Though comparisons between the two incidents are inevitable, Anthropic drew a clear distinction between its incidents and OpenAI’s, noting where OpenAI’s model exploited an unknown software vulnerability to break out of its test environment, Anthropic’s models instead reached the internet through a path that had, by mistake, been left open.
OpenAI has continued to release new details about its own breach, saying its models also used publicly exposed credentials across four accounts on four services: one as a staging point, one for storage, and two that were only looked at, not used to break in further, according to OpenAI’s own updated blog post about the incident.
Anthropic also drew a distinction between itself and OpenAI by noting that it discovered the incidents itself, through a proactive review, and that the two affected organizations it was able to reach hadn’t previously detected the activity or flagged it to Anthropic.
The company added that it’s now working with the independent evaluation group METR on a third-party review of the incidents.
OpenAI’s accidental breach of Hugging Face, which was the first verifiable case of an AI lab losing control of its model, sparked a string of reactions from the industry and politicians, many of whom don’t necessarily agree with one another. This latest disclosure from Anthropic ensures the debate over AI models and security will continue.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
Kirsten Korosec
Transportation Editor
October 13 – 15
San Francisco
Scale faster. Grow your portfolio. Gain practical expertise. No matter your goal, Disrupt can empower you.
Save up to $330 toda y!
Claude Opus 5 became downright ruthless when tasked with running a vending machine
Sam Altman is ready to decelerate
PSA: Your Claude shared chats and Artifacts may have ended up on Google
Librarians are hosting viral ‘Avoiding AI’ workshops for people who are fed up with Big Tech
SpaceX launches new V3 Starlink satellites but suffers another booster failure
Prentis, new AI lab co-founded by Reid Hoffman, Mark Pincus in talks to raise $100M
US accuses American of allegedly wiping his phone using a ‘duress’ password during border search
