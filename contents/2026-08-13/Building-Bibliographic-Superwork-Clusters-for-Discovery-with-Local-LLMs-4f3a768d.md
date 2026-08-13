---
source: "https://thisismattmiller.com/post/superwork/"
hn_url: "https://news.ycombinator.com/item?id=49291063"
title: "Building Bibliographic Superwork Clusters for Discovery with Local LLMs"
article_title: "Building Bibliographic Superwork Clusters for Discovery with Local LLMs"
author: "sebg"
captured_at: "2026-08-13T20:31:36Z"
capture_tool: "hn-digest"
hn_id: 49291063
score: 1
comments: 0
posted_at: "2026-08-13T19:55:20Z"
tags:
  - hacker-news
  - translated
---

# Building Bibliographic Superwork Clusters for Discovery with Local LLMs

- HN: [49291063](https://news.ycombinator.com/item?id=49291063)
- Source: [thisismattmiller.com](https://thisismattmiller.com/post/superwork/)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T19:55:20Z

## Translation

タイトル: ローカル LLM を使用した発見のための書誌スーパーワーク クラスターの構築
説明: 88K クラスターでの仕事関係の判断

記事本文:
ローカル LLM を使用した発見のための書誌スーパーワーク クラスターの構築
88,000 個のクラスターで仕事上の関係を判断する
書誌事項の説明では、さまざまな方法で外部の権威や記録との関係を構築できます。 MARC では、制御された語彙を使用でき、最終的には識別子と URI を使用して接続を確立する機能が進化しました。具体的には、名前タイトル権限またはシリアル カタログ作成に主に使用される上位 7XX フィールドの説明を介して、他のタイトルにリンクできます。 FRBR-LRM/WEMI/Linked Data に準拠した標準に移行するにつれ、Bibframe リレーションシップなどの標準がレコードの構築方法の中心となるようになります。作品とインスタンスは別個のリソースであり、Bibframe Hub は、カタログ化を目的として、Name Title Authorities の機能を置き換え、その機能を拡張します。ここでは、カタログ化ではなく発見のために構築できる関係の種類について考えているため、ここでは詳しく説明しません。次の小さな実験は、スーパーワークのアイデアに基づいて作品を集約する関係を構築する自動化された方法を検討しています。小説、翻訳、映画化、13 の異なる版はすべて、定義された関係に関連付けられています。
ハブはコロケーションを行いますが、特定のレベルの表現に範囲が限定されているという点で名前付きタイトルの権威に似ており、すべての権威の作業と同様に、カタログ作成時に追加の労力がかかります。つまり、作成時に非常に特殊なルールがあることを意味します。スーパーワークのこのアイデアはより拡張的であり、もちろん新しいアイデアではありません。ワークスとより高いものの役割については多くの議論があります。 Share-VDE プロジェクトは、この文書で説明する Opus と呼ばれるスーパーワーク レベルを実装しました。つまり、それは新しいアイデアではありません

良いアイデアかもしれませんが、問題は、さまざまな理由から実装が難しいことです。
書誌データのクラスタリングもよく知られた道です。 25 年前、OCLC のような組織は FRBR 化の取り組みのために作業をクラスタリングしていました。このプロセスには、さまざまなフィールドで MARC レコードの一部を照合して、リソースのグループを作成することが含まれます。したがって、クラスタリングは「簡単」ですが、クラスタリングした後は、それらを相互にどのように関連付けるのでしょうか?複雑なルールベースのプロセスを考え出す必要がありますが、さまざまな完全性レベルで記述された記録、矛盾する歴史的実践、無限に続く問題に対してすぐに破綻してしまいます。
これらすべての作業が互いにどのように関連しているかを判断するのに役立つ LLM を使用するというアイデアを取り入れてください。私は AI の実践派で、今年 AI による書籍の「カタログ化」に関するデモやピッチをいくつか受けてきましたが、図書館における AI に対するそのようなアプローチは本当に有害でひどいものであると感じています。 LLM は特定のタスクやワークフローの自動化に潜在的に役立つと思います。そして、レガシーデータを強化するというこの考え方には一理あると思います。また、高価なブラック ボックス フロンティア モデルの代わりに、自分の (頑丈な) コンピューター上で実行できるローカル モデルを使用することにも非常に興味がありました。
カタログに相当するレコードを作品にクラスタリングする
ローカル LLM を使用して、クラスターのスーパーワークが何か、およびすべての作業がそれにどのように関連しているかを判断します。
それがどの程度うまくいったか、結果がどの程度役立つかをテストします。
ソース データには米国議会図書館のカタログを使用します。私はそこで働いていますが、これはサイドプロジェクト (🫠) であり、公開されているダウンロード可能な一括データを使用しています。このカタログをクラスタリングするために、途方もなく広い網を張りたかったので、同じタイトルと一致する 1XX または 7XX の寄稿者を持つものを探します。これにより、次のようなものが見逃されます

翻訳されたタイトル (Der Idiot vs The Idiot) または音訳されたタイトルですが、私はこれらのギャップを埋めるために名前タイトル権限を使用しています。この非常識なクラスタリング手法を適用すると、約 2,000 万件のレコードから次の内訳が得られます。
完全な候補ユニバース (1,969,233 件のタイトル + 名前の一致 → Union-Find マージ)
したがって、クラスターに属するレコードは 20% だけである可能性があります。最大のクラスターは 2 つのもので、これらのクラスターは多くの場合、同じ作品の 2 つのフォーマットが別々に記録されている MARC カタログ作成の成果物です (たとえば、印刷物と電子ブック)。私のテストでは、非常に難しいクラスター サイズを実行したいと考えていましたが、すべてを実行したくはありません。私は 5 ～ 40 スライスを良いテスト グループとして選びました。グループ分けは次のように行われます。
カットオフは 40 レコードでしたが、名前、タイトル権限のマージにより、リソースに変化はなく、89,856 → 88,534 クラスター (-1,322) が崩壊しました。 146 個のマージされたクラスターが、元の 40 レコードの判定上限を超えています。
したがって、さまざまなサイズの 88,000 個のクラスターを判断する必要があります。これらの事柄が互いにどのように関連しているかを整理するための語彙が必要になります。限られた語彙を作成するために、クラスター内に存在するデータに基づいて関係プロパティを構築するのに役立つローカル LLM を使用することにしました。仕事上の人間関係に関する語彙は世の中にあり、それで十分かもしれませんが、ボトムアップのアプローチを試してみたかったのです。ただし、それが作成された後、既存のもの、具体的には Bibframe 作品と表現関係の語彙および RDA レジストリ ワーク関係と比較したいと思いました。以下は、有機的な関係の語彙のリストと、それが 2 つの確立された語彙にどのようにマッピングされるかを示しています。
BIBFRAME 関係へのマッピング
RDA レジストリへのマッピング 仕事関係
前に述べたように、ローカル LLM を使用してこれらの作品のクラスターを判断します。

これには、各クラスターのタイトルとメタデータ データが、可能な関係リストとともに表示されます。私は Qwen3.6-35B-A3B-bf16 モデルを使用しています。これは 350 億のパラメーター モデルですが、アクティブなパラメーターは 30 億しかありません。これは専門家混合 (MoE) モデルであり、多くのパラメーターを利用できますが、迅速に実行できます。私の M5 128GB マシンでは、4 ～ 6 個のプロンプトを並行して実行すると、それぞれ 1 秒あたり最大 17 トークンが発生します。これはまともなクリップですが、すべての 88K クラスターを通過するために数週間バックグラウンドで実行されました。 Open Router のようなプロバイダーで同じモデルを使用してこのようなジョブを実行した場合、トークンのコストは約 150 ～ 200 ドルになりますが、おそらく数週間ではなく 1 日かかります。 Gemini などで実行した場合、コストははるかに高くなりますが、おそらく品質は向上します。したがって、お金、時間、品質、地元など、すべてはトレードオフです。
すべてが判断されたので、まず LLM がどの程度うまく機能したかを評価し、次にクラスターがどのように機能したかを分析したいと思いました。それらを判断するために、私は 2 つのアプローチを実行しました。200 個のクラスターを手動でレビューしてその精度を判断し、またフロンティア モデル (Opus 4.8) に 2500 個のクラスターをレビューして判断させました。ここで正しさを考えるときは 2 つの要素があり、大きく 4 つのバケツがあります。
Belongs: 所属、コア表現の実現
関連: 属しており、ソース/派生/続編などの明確に関連する表現
フォールバック: 裁判官を困惑させ、彼らをどう扱ったらよいかわからない
それぞれの関係は、これらのバケットの 1 つにまとめられます。したがって、クラスター内のタイトルは、一般的には正しいものの、特定的には正しくない可能性があります。結果:
¹ はラベル全体で平均されるため、希少なラベルの重さは一般的なラベルと同じになります。
私はモデルよりも寛大にグレーディングしました。+40 の可能なプロパティのうち、モデルは非常に具体的な傾向があり、有益な統計です。

品質はおそらく 4 つの広範なカテゴリのスコアであり、私はモデルに 96%、モデルに 91% を与えました。以下にグレーディング結果の詳細を示します。
素材ごとの強いところ、弱いところ
どのラベルが過剰に使用されているのか、どのラベルが過小使用されているのでしょうか?
本当に重要な決断
01 どのような関係ですか?
各レコードには、44 の用語で管理された語彙から 1 つのラベルが付けられます。少数のレコードが優勢です。ほとんどのレコードは、単なる別の版、印刷、または同じ作品の一部です。
02 ズームアウト: それは属していますか?
44 個のラベルを 4 つのバケットに集約します。クラスタリングの中核となる仕事、つまりメンバーシップの決定は、圧倒的に「はい、これに属します」というものであり、小さな重要なスライスは偶然として除外されます。
表現（同一作品） 88% 関連表現 10% 該当しない 1.7% 不明瞭 0.0%
03 コーパスには何が入っているの？
マテリアルの種類ごとに記録し、クラスターがどのくらい大きくなるかを記録します。ほとんどのクラスターは小さいです。一般的なタイトルのロングテールは数十のレコードに及びます。
クラスターのサイズ分布 (クラスターごとのレコード)
2 人の独立した採点者が層別サンプル、つまり母集団を代表するセットに関する人間の専門家 (あなた) と、2,500 件のレコードに関するフロンティア モデル (Opus) を採点しました。正確なラベルの精度とバケットレベルの精度との間のギャップが問題です。ほとんどの「エラー」は、正しいバケット内でほぼ同数です。
0% 25% 50% 75% 100% 87% 76% 正確なラベル 97% 91% 粗いバケット
バケット レベルでは、両方のグレーダーがモデルが 90 ～ 97% の確率で正しいことに同意します。オーパスは、正確なラベルに関してより厳格なグレーダーです。
05 素材別の強いところ・弱いところ
マテリアルごとのフロンティアグレードの精度 (人口に重み付け; 暗い = 強い)。書籍、地図、ビデオ、連載は好調です。モデルが苦労するのは画像とソフトウェアです。
06 どのラベルが過剰に使用されているか、どのラベルが使用されていませんか?
すべての関係ラベルは、ドットのペアとして少なくとも 15 回グレーディングされます: 精度 (暗 — モデルがラベルを使用する場合、

それが正しい頻度）と思い出し（光 - 本当のケースが実際にそれを得る頻度）。両者のギャップによって分類されます。
0% 0% 25% 25% 50% 50% 75% 75% 100% 100% P / R UNDER-used — 適用すると正確ですが、ほとんどの真のケースを見逃します Illustration_edition — 精度 100%、リコール 10% (グレード 46x) Illustration_edition 100 / 10 Critical_or_annotated_edition — 精度 100%、リコール16% (57 倍の評価) Critical_or_annotated_edition 100 / 16 見かけの重複 — 精度 93%、再現率 25% (54 倍の評価) 見かけの重複 93 / 25 Different_sheet_area — 精度 98%、再現率 32% (52 倍の評価) Different_sheet_area 98 / 32 RELATED_WORK — 精度 93%、再現率31% (評価 276×) 関連作業 93 / 31 不確定 — 精度 68%、再現率 13% (評価 28×) 不確定 68 / 13 配置 — 精度 93%、再現率 48% (評価 65×) 配置 93 / 48 Different_release_issue — 精度 78%、再現率 39% (評価 41×) Different_release_issue 78 / 39 Different_state_or_impression — 適合率 63%、再現率 27% (等級付け 49×) Different_state_or_impression 63 / 27 abridged_edition — 適合率 98%、再現率 64% (等級付け 44×) abridged_edition 98 / 64Previous_title — 適合率 66%、再現率 34% (等級付け) 31×) 先行_タイトル 66 / 34 エピソードまたはインストール — 精度 100%、再現率 82% (グレード 108×) エピソードまたはインストール 100 / 82 Different_music_edition — 精度 49%、再現率 32% (グレード 26×) Different_music_edition 49 / 32 バランスの取れたコンポーネントパート — 精度 81%、再現率 69% (グレード 203×) コンポーネント_パート 81 / 69 楽器パーツ — 精度 78%、再現率 66% (グレード 33×) 楽器部品 78 / 66 改訂版マップ_エディション — 精度 92%、再現率 83% (グレード 68×) 改訂版_マップ_エディション 92 / 83 コンパイル_外観 — 精度 44%、再現率 35% (グレード21×) コンピレーション_appea

rance 44 / 35 Different_work — 精度 83%、再現率 75% (グレード 60×) Different_work 83 / 75 Different_platform — 精度 95%、再現率 92% (グレード 40×) Different_platform 95 / 92special_edition_release — 精度 97%、再現率 97% (グレード 37×)special_edition_release 97 / 97 excerpt_or_trailer — 精度 100%、再現率 100% (グレード 26×) excerpt_or_trailer 100 / 100 dubbed_
[切り捨てられた]
07 本当に重要な決断
クラスターからレコードを削除するラベルは 1 つだけです: Different_work 。ここでは、精度 (除外する場合は正しいか?) と再現率 (真の侵入者を捕捉するか?) を示します。人間とオーパスです。
0% 25% 50% 75% 100% 80% 83% 精度 100% 75% リコール
どちらも、モデルが 5 分の 1 を過剰に除外していることに同意しています (精度は 80% まで)。 Opus のリコールが低いのは、Opus 自体が厳しすぎることでした。Opus が記録の破棄を望んでいた 15/15 の境界ケースでは、人間はモデルを支持しました。
両方の記録を確認しました。モデルとの一致、および審査員間での一致 (Cohen の κ = 0.48、「中程度」)。
クラスタの 1.5% (1,374) では、ほとんどのレコードが除外されていました。一般的なタイトルが無関係な作品を一掃しました。しかし、ほとんどの「失敗」は無駄なものではありません。
すべてのレコードが除外される 26% 1 つのレコードを除くすべてのレコードが除外される 43% 大部分が除外される 31%
失敗したクラスターの 74% は、実際にはゴミではなく、1 つのタイトルの下にマージされた複数の本物の作品です。審査員がそれらを再分割することはできません。
例 -

[切り捨てられた]

## Original Extract

Judging work relationships in 88K clusters

Building Bibliographic Superwork Clusters for Discovery with Local LLMs
Judging work relationships in 88K clusters
In bibliographic description you are able to build relationships to external authorities and records in various ways. In MARC you can use controlled vocabularies which eventually evolved the ability to use identifiers and URIs to establish connections. You can link to other titles specifically via a Name Title authority or some description in the higher 7XX fields mostly used for Serial cataloging. As we move into more FRBR-LRM/WEMI/Linked Data aligned standards like Bibframe relationships become more central to how the record is constructed. The Work and the Instance are separate resources, the Bibframe Hub takes the place of and expands the functionality of Name Title Authorities for cataloging purposes. There are more details I’m glossing over here because I’m thinking about the type of relationships that can be built for discovery rather than cataloging. The following little experiment is looking at automated ways to build relationships that aggregate works under the idea of a Superwork: The novel, the translation, the movie adaptation, the 13 different editions, all linked with their relationships defined.
While Hubs do collocation they are similar to Named Title authorities in that they are scoped to specific levels of expression and like all authority work it takes additional effort while cataloging meaning there are very specific rules to when one is created. This idea of Superwork is much more expansive and is of course not a new idea, lots of debate about the role of Works vs something higher. The Share-VDE project implemented a Superwork level they called an Opus that is discussed in this paper . So it’s not a novel idea, it might be a good idea, but the problem is that it’s hard to implement for a number of reasons.
Clustering bibliographic data is also a well trodden path. Twenty-five years ago orgs like OCLC were clustering works for FRBR-ization efforts. The process involves matching parts of MARC records on various fields to create groupings of resources. So clustering is “easy” but once they are clustered how do you relate them to one another? You have to come up with elaborate rule based processes that quickly fall apart against records described at different levels of completeness, conflicting historical practices and an infinite long tail of problems.
Enter the idea of using a LLM to help you judge how all these works are related to one another. I’m an AI pragmatic, I’ve been subjected to a couple demos and pitches this year of AI “cataloging” books and I find that sort of approach to AI in libraries really harmful and gross. I think LLMs can potentially be useful for specific tasks or workflow automations. And I think it has a place in this idea of enriching legacy data. I’ve also been really interested in using local models that can run on your own (beefy) computer as alternatives to expensive black box frontier models.
Cluster a catalog’s worth of records into works
Use a local LLM to judge what is the Superwork of the cluster and how all the works relate to it
Test how well it did and how useful are the results.
For the source data I’m going to use the Library of Congress catalog. I do work there but this is a side project (🫠) and I’m using their public downloadable bulk data. To cluster this catalog I wanted to cast a ridiculously wide net, so it looks for things that have the same title and a matching 1XX or 7XX contributor. This will miss things like translated titles (Der Idiot vs The Idiot) or transliterated titles but I use Name Title authorities to try to fill in those gaps. Applying this insane clustering methodology we get the following breakdown out of the ~20M records:
Full candidate universe (1,969,233 title+name matches → union-find merged)
So only 20% of records possibly belong to a cluster. The biggest cluster is 2 things, these clusters are often artifacts of MARC cataloging where two formats of the same work are recorded separately (print + ebook for example). For my test I wanted to do a significantly challenging cluster size, but don’t want to do everything. I picked the 5-40 slice as a good test group. This is how that grouping breaks down:
While the cutoff was 40 records Name Title authority merging collapsed 89,856 → 88,534 clusters (−1,322) with no change in resources. 146 merged clusters now exceed the original 40-record judging cap.
So we have 88K clusters of varying sizes to judge. We now need a vocabulary to organize how these things relate to each other. I chose to use the local LLM to help build the relationship properties based on data present in the cluster to make a limited vocabulary. There are vocabularies out there for work relationships, which might have been sufficient but I wanted to try a bottom up approach. Though after it was created I did want to compare to what already exists, specifically the Bibframe work to expression relationship vocabulary and the RDA Registry Work relationships . Here is the list of the organic relationship vocabulary and how it maps to the two established ones:
Mapping to BIBFRAME relationships
Mapping to RDA Registry Work relationships
As I mentioned earlier we are going to use a local LLM to judge these clusters of works, basically prompting with the titles of each cluster and metadata data along with the possible relationship list. I’m using Qwen3.6-35B-A3B-bf16 model which is a 35 billion parameter model but only ever has 3 billion parameters active. This is a Mixture-of-Experts (MoE) model that allows it to draw on a lot of parameters but is able to execute quickly. Running 4-6 prompts in parallel results in ~17 tokens a second each on my M5 128GB machine. Even though that is a decent clip it still ran in the background for a couple weeks to get through all 88K clusters. If you ran a job like this using the same model on a provider like Open Router it would cost around $150-200 in tokens but probably take a day rather than weeks. If you ran it on something like Gemini it would cost MUCH more but probably better quality. So it’s all a trade off, money vs time vs quality vs local, etc.
Now that everything is judged I wanted to first evaluate how well the LLM worked and then break down how the clusters panned out. To judge them I did two approaches, I manually reviewed 200 clusters and judged their accuracy and I also had a frontier model (Opus 4.8) review 2500 and judge. There are two factors when thinking about correctness here, there are 4 broad buckets:
Belongs: belongs, a realization of the core expression
Related: belongs, a distinct related expression like source/derivative/sequel
Fallback: Stumps the judge, not sure what to do with them
Each of the relationships roll up into one of these buckets. So a title in the cluster could be generally correct but specifically incorrect. The results:
¹ averaged over labels, so rare labels weigh as much as common ones
I was more generous while grading than the model, out of +40 possible properties the model tends to be very specific, a useful stat that indicates quality is probably the 4 broad category score, I gave it a 96% the model a 91%. Below are more views into the grading results:
Strong and weak spots, by material
Which labels get over- and under-used?
The decision that actually matters
01 What kinds of relationships?
Each record gets one label from a 44-term controlled vocabulary. A handful dominate: most records are simply another edition, printing, or part of the same work.
02 Zoomed out: does it belong?
Collapsing the 44 labels into four buckets. The clustering's core job — deciding membership — is overwhelmingly "yes, this belongs," with a small, important slice excluded as coincidental.
manifestation (same work) 88% related expression 10% does not belong 1.7% unclear 0.0%
03 What's in the corpus?
Records by material type, and how big the clusters get. Most clusters are small; a long tail of generic titles runs to dozens of records.
Cluster size distribution (records per cluster)
Two independent graders scored a stratified sample: a human expert (you) on a population-representative set, and a frontier model (Opus) on 2,500 records. The gap between exact-label and bucket-level accuracy is the story — most "errors" are near-ties inside the right bucket.
0% 25% 50% 75% 100% 87% 76% Exact label 97% 91% Coarse bucket
At the bucket level both graders agree the model is right ~90–97% of the time. Opus is the stricter grader on exact labels.
05 Strong and weak spots, by material
Frontier-graded accuracy per material (population-weighted; darker = stronger). Books, maps, video and serials are strong; images and software are where the model struggles.
06 Which labels get over- and under-used?
Every relationship label graded at least 15 times, as a pair of dots: precision (dark — when the model uses the label, how often it's right) and recall (light — how often the true cases actually get it). Sorted by the gap between the two.
0% 0% 25% 25% 50% 50% 75% 75% 100% 100% P / R UNDER-USED — precise when applied, but misses most true cases illustrated_edition — precision 100%, recall 10% (graded 46×) illustrated_edition 100 / 10 critical_or_annotated_edition — precision 100%, recall 16% (graded 57×) critical_or_annotated_edition 100 / 16 apparent_duplicate — precision 93%, recall 25% (graded 54×) apparent_duplicate 93 / 25 different_sheet_area — precision 98%, recall 32% (graded 52×) different_sheet_area 98 / 32 related_work — precision 93%, recall 31% (graded 276×) related_work 93 / 31 indeterminate — precision 68%, recall 13% (graded 28×) indeterminate 68 / 13 arrangement — precision 93%, recall 48% (graded 65×) arrangement 93 / 48 different_release_issue — precision 78%, recall 39% (graded 41×) different_release_issue 78 / 39 different_state_or_impression — precision 63%, recall 27% (graded 49×) different_state_or_impression 63 / 27 abridged_edition — precision 98%, recall 64% (graded 44×) abridged_edition 98 / 64 preceding_title — precision 66%, recall 34% (graded 31×) preceding_title 66 / 34 episode_or_installment — precision 100%, recall 82% (graded 108×) episode_or_installment 100 / 82 different_music_edition — precision 49%, recall 32% (graded 26×) different_music_edition 49 / 32 BALANCED component_part — precision 81%, recall 69% (graded 203×) component_part 81 / 69 instrumental_parts — precision 78%, recall 66% (graded 33×) instrumental_parts 78 / 66 revised_map_edition — precision 92%, recall 83% (graded 68×) revised_map_edition 92 / 83 compilation_appearance — precision 44%, recall 35% (graded 21×) compilation_appearance 44 / 35 different_work — precision 83%, recall 75% (graded 60×) different_work 83 / 75 different_platform — precision 95%, recall 92% (graded 40×) different_platform 95 / 92 special_edition_release — precision 97%, recall 97% (graded 37×) special_edition_release 97 / 97 excerpt_or_trailer — precision 100%, recall 100% (graded 26×) excerpt_or_trailer 100 / 100 dubbed_
[truncated]
07 The decision that actually matters
Only one label removes a record from a cluster: different_work . Here's precision (when it excludes, is it right?) and recall (does it catch the true intruders?) — human vs Opus.
0% 25% 50% 75% 100% 80% 83% Precision 100% 75% Recall
Both agree the model over-excludes ~1 in 5 (precision ~80%). Opus's lower recall was Opus itself being over-strict: on 15/15 boundary cases where Opus wanted a record thrown out, the human sided with the model.
On the records both reviewed. Agreement with the model, and between the judges (Cohen's κ = 0.48, "moderate").
1.5% of clusters (1,374) had most records excluded — a generic title swept up unrelated works. But most "failures" aren't junk:
every record excluded 26% all but one record 43% majority excluded 31%
74% of failed clusters are actually several real works merged under one title , not garbage — the judge just can't re-partition them.
Example —

[truncated]
