---
source: "https://optimaflo.io/blog/connect-llm-to-your-data"
hn_url: "https://news.ycombinator.com/item?id=49161381"
title: "Connecting an LLM to Your Data Is the 21% Solution"
article_title: "Connecting an LLM to Your Data Is the 21% Solution. | OptimaFlo Blog"
author: "evan_rosa"
captured_at: "2026-08-03T21:59:32Z"
capture_tool: "hn-digest"
hn_id: 49161381
score: 1
comments: 0
posted_at: "2026-08-03T21:00:24Z"
tags:
  - hacker-news
  - translated
---

# Connecting an LLM to Your Data Is the 21% Solution

- HN: [49161381](https://news.ycombinator.com/item?id=49161381)
- Source: [optimaflo.io](https://optimaflo.io/blog/connect-llm-to-your-data)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T21:00:24Z

## Translation

タイトル: LLM をデータに接続することが 21% の解決策です
記事のタイトル: LLM をデータに接続することが 21% の解決策です。 |オプティマフロのブログ
説明: コネクタのピッチはどこにでもあります。公表されている数字は、Anthropic で 21%、Spider 2.0 で 10 ～ 17% です。ギャップを埋めるもの、そしてコネクターが決して構築しないもの。

記事本文:
LLM をデータに接続することが 21% の解決策です。 | OptimaFlo ブログ メイン コンテンツにスキップ OptimaFlo 製品
テーマの切り替え OptimaFlo 製品の構築を開始する
テーマの切り替え OptimaFlo データ エンジニアリングの構築を開始する LLM をデータに接続することが 21% の解決策です。
最初の会話のほぼすべてで同じ質問を耳にします。クロードに Postgres を教えて質問できるのに、なぜプラットフォームが必要なのでしょうか?
それは公平な質問です。コネクタは本物です。 Claude は Snowflake と Postgres に接続し、ChatGPT は MCP ブリッジを介してウェアハウスに到達し、Gemini は BigQuery 内に配置されます。デモは本当に印象的です。
また、デモの後に何が起こるかについての公開数字も入手しました。
Anthropic は、独自の分析ワークロードでそれを測定しました。「スキルがなければ、分析の質問に正確に答えるクロードの能力は、私たちの評価で 21% を超えませんでした。」データベースに接続するのと同じモデル。データ、質問、21%。
ベンチマークも一致しています。 Spider 2.0 は、数千列のスキーマ、100 行のクエリ、複数の SQL ダイアレクトなど、実際のエンタープライズ ウェアハウス ワークフローから構築されています。 GPT-4o のスコアは 10.1% です。 o1-preview は 17.1% に達します。同じモデルは、古いアカデミック ベンチマークで 85% をクリアしています。これが、デモが魔法のように感じられるのに、4 週間目は魔法のように感じられない理由です。
BIRDは崩壊について説明します。すべての BIRD の質問には、フィールドの意味を説明する手書きのヒントが付属しています。ヒントを削除すると、GPT-4 は 54.89% から 34.88% に低下します。ユーザーはヒントを書きません。彼らは「アクティブな顧客が何人いるか」を尋ね、モデルがアクティブの意味を理解していると仮定します。そうではありません。 40 個の候補列から 1 つを選択し、自信を持って合計します。
それが危険な部分です。間違った番号は間違っているようには見えません。 Anthropic はそれをサイレント失敗と名付け、強化されたスタックがむしろそれを減少させることを認めています

それを排除するよりも。
Anthropic はその差を 21% から 95% 以上に縮めました。より大規模なモデルやより多くのデータ アクセスを使用したわけではありません。エージェントに、何千もの自分自身の正しい過去のクエリへの grep アクセスを提供し、精度の向上は 1 ポイント未満でした。
機能したのは構造でした。管理された正規データセット。人間がすべてのメトリクス定義を所有するセマンティック レイヤー。系統。注意深く分析する人の仕事の仕方をコード化するスキル。そのすべてを監視する評価ハーネス。世界最高のデータ チームの 1 つによって構築および維持されている 4 つのレイヤー。この投稿が最初に着地したときに、誰がその足場を組み立てることができるかについて書きました。
他のベンダーも同じ結論を自社の製品に組み込みました。 Snowflake は、生のスキームにはアナリストが必要とする意味が欠けているため、独自の MCP サーバーをセマンティック ビューに制限します。 Google は Gemini をナレッジ カタログと Looker セマンティック レイヤーでラップし、このバンドルを製品と呼んでいます。コネクタはインターフェースです。その下にある管理された層が彼らの売り物なのです。
正確さはさておき、ピッチにはカテゴリー誤差があります。コネクタはテーブルを読み取ります。データの問題は、誰もテーブルを構築していないことです。
コネクタはソースを倉庫に配置しません。 Raw から Clean、Ready までをモデル化することはありません。チャットを開くと実行され、閉じると停止します。ChatGPT セッションは切断時に状態を落とし、スケジュールに従って何も実行しません。一晩中何も監視しません。チャット メモリはセマンティック レイヤーではないため、セッションごとにビジネスを再学習します。
そして、配管自体はあなたの負担です。 Anthropic は、クエリ タイムアウト、行キャップ、列ブロックリストのない参照 Postgres MCP サーバーを、セキュリティ保証なしという明白な警告とともにアーカイブしました。今でも月に約 312,000 回のインストールが行われていますが、そのほとんどはピッチ内でセットアップを正確に配線している人々によるものです。
OptimaFlo はこれらの 上で実行されます

アメモデル。自分のクロード、GPT、またはジェミニ キーを持参します。エンジンについては異論はありません。
私たちが売りにしているのは 74 点です。つまり、管理されたデータセット、セマンティック レイヤー、パイプライン、品質チェック、そして Anthropic が構築するためにチームにスタッフを配置した評価規律であり、1 人のデータ所有者が独自のクラウド内のオープンな Iceberg テーブルで実行できる製品として出荷されます。 SQL がテーブルにアクセスする前に、人間が SQL を承認します。チャットインターフェイスもそこにあります。これは最後のステップですが、簡単なステップです。
倉庫がすでにモデル化され管理されている場合、その上に 20 ドルのチャット シートを追加するのは非常にお買い得です。私は比較でそう言います。誰もその倉庫を建てなければ、椅子には立つものが何もありません。
21% はこのモデルがもたらすものです。残りは仕事です。私たちがどのようにそれを行うかを見に来てください。または、7 日間のパイロットで実際に試してみてください。
AI データ チームを独自のクラウドに。生データからダッシュボードまで、操作するプラットフォームはありません。
© 2026 オプティマフロ。無断転載を禁じます。
ROI を計算する 私たちはあなたのプライバシーを尊重します
当社は、ブラウジングエクスペリエンスを向上させ、パーソナライズされたコンテンツを提供し、トラフィックを分析するために Cookie を使用します。 「すべてを受け入れる」をクリックすると、Cookie の使用に同意したことになります。設定をカスタマイズしたり、Cookie ポリシーとプライバシー ポリシーで詳細を確認したりできます。

## Original Extract

The connector pitch is everywhere. The public numbers on it: 21% at Anthropic, 10 to 17% on Spider 2.0. What closes the gap, and what a connector never builds.

Connecting an LLM to Your Data Is the 21% Solution. | OptimaFlo Blog Skip to main content OptimaFlo Product
Toggle theme Start building OptimaFlo Product
Toggle theme Start building OptimaFlo Data Engineering Connecting an LLM to Your Data Is the 21% Solution.
I hear the same question in almost every early conversation: why do I need a platform when I can point Claude at Postgres and start asking questions?
It is a fair question. The connectors are real. Claude wires into Snowflake and Postgres, ChatGPT reaches warehouses over MCP bridges, Gemini sits inside BigQuery. The demo is genuinely impressive.
We also now have public numbers on what happens after the demo.
Anthropic measured it on their own analytics workload: "Without skills, Claude's ability to answer analytics questions accurately didn't exceed 21% on our evals." Same model you would connect to your database. Their data, their questions, 21%.
The benchmarks agree. Spider 2.0 is built from real enterprise warehouse workflows: thousand-column schemas, hundred-line queries, multiple SQL dialects. GPT-4o scores 10.1% on it. o1-preview reaches 17.1%. The same models clear 85% on the older academic benchmark, which is why the demo feels like magic and the fourth week does not.
BIRD explains the collapse. Every BIRD question ships with a hand-written hint explaining what the fields mean. Remove the hints and GPT-4 falls from 54.89% to 34.88% . Your users do not write hints. They ask "how many active customers do we have" and assume the model knows what active means. It does not. It picks one of your four dozen candidate columns and sums with confidence.
That is the dangerous part. A wrong number does not look wrong. Anthropic names it the silent failure and admits their hardened stack reduces it rather than eliminating it.
Anthropic closed the gap from 21% to above 95%. Not with a bigger model, and not with more data access: they gave the agent grep access to thousands of their own correct past queries and accuracy moved less than one point.
What worked was structure. Governed canonical datasets. A semantic layer where a human owns every metric definition. Lineage. Skills that encode how a careful analyst works. An eval harness watching all of it. Four layers, built and maintained by one of the best data teams in the world. I wrote about who gets to build that scaffolding when the post first landed.
The other vendors built the same conclusion into their products. Snowflake restricts its own MCP server to semantic views because raw schemas lack the meaning an analyst needs. Google wraps Gemini in a Knowledge Catalog and a Looker semantic layer and calls the bundle the product. The connector is the interface; the governed layer underneath is what they sell.
Accuracy aside, there is a category error in the pitch. A connector reads tables. Your data problem is that nobody is building the tables.
The connector will not land your sources in a warehouse. It will not model Raw into Clean into Ready. It runs when you open a chat and stops when you close it: ChatGPT sessions drop their state on disconnect and run nothing on a schedule . It watches nothing overnight. It re-learns your business every session, because chat memory is not a semantic layer.
And the plumbing itself is on you. Anthropic archived its reference Postgres MCP server, the one without query timeouts, row caps, or column blocklists, behind a plain warning: no security guarantees. It still gets about 312k installs a month , mostly from people wiring up exactly the setup in the pitch.
OptimaFlo runs on these same models. You bring your own Claude, GPT, or Gemini key. I have no argument with the engine.
What we sell is the 74 points: the governed datasets, the semantic layer, the pipelines, the quality checks, and the eval discipline that Anthropic staffed a team to build, shipped as a product one data owner can run, on open Iceberg tables in your own cloud. A human approves the SQL before it touches your tables. The chat interface is in there too. It is the last step, and it is the easy one.
If your warehouse is already modeled and governed, a $20 chat seat on top of it is a great deal, and I say so in our side-by-side comparison . If nobody is building that warehouse, the seat has nothing to stand on.
The 21% is what the model brings. The rest is the job. Come see how we do it , or kick the tires in a 7-day pilot .
Your AI data team, in your own cloud. From raw data to dashboards, no platform to operate.
© 2026 OptimaFlo. All rights reserved.
Calculate ROI We value your privacy
We use cookies to enhance your browsing experience, serve personalized content, and analyze our traffic. By clicking "Accept All", you consent to our use of cookies. You can customize your preferences or learn more in our Cookie Policy and Privacy Policy .
