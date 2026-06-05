---
source: "https://zilliz.com/blog/why-we-built-loon-a-storage-engine-for-ai-data-that-never-stops-changing"
hn_url: "https://news.ycombinator.com/item?id=48413538"
title: "Introduce Loon: lake-native storage engine behind Milvus 3.0 for AI data"
article_title: "AI Datasets Are Never Done. So We Built Loon. - Zilliz blog"
author: "redskyluan"
captured_at: "2026-06-05T16:06:30Z"
capture_tool: "hn-digest"
hn_id: 48413538
score: 3
comments: 0
posted_at: "2026-06-05T15:06:45Z"
tags:
  - hacker-news
  - translated
---

# Introduce Loon: lake-native storage engine behind Milvus 3.0 for AI data

- HN: [48413538](https://news.ycombinator.com/item?id=48413538)
- Source: [zilliz.com](https://zilliz.com/blog/why-we-built-loon-a-storage-engine-for-ai-data-that-never-stops-changing)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T15:06:45Z

## Translation

タイトル: Loon の紹介: AI データ用の Milvus 3.0 の背後にあるレイクネイティブ ストレージ エンジン
記事のタイトル: AI データセットは決して完成しない。それで私たちはLoonを作りました。 - ジリズブログ
説明: Loon は、Milvus 3.0 および Zilliz Vector Lakebase 用の新しいストレージ エンジンであり、ColumnGroups、行 ID の配置、およびマニフェストを使用して進化するベクター データセットを管理するために構築されています。

記事本文:
AI データセットは決して完成しません。それで私たちはLoonを作りました。 - Zilliz ブログ 製品 Zilliz Cloud
大規模な高い信頼性、パフォーマンス、コスト効率を実現するように設計された、フルマネージドの Vector Lakebase サービス。
10 億規模のベクトル類似性検索用に構築されたオープンソースのベクトル データベース。
Zilliz Vector Lakebase BYOC を発表
価格設定 料金プラン あらゆる予算であらゆるチームに柔軟な価格設定オプションを提供
定価 すべての請求項目の詳細な定価を表示します
無料利用枠 マネージド Milvus の力で想像力を解き放ちます
Zilliz Cloud 開発者ハブでは、Zilliz Cloud を使用するためのすべての情報が見つかります。
リソース ブログ ガイド リサーチ アナリスト レポート ウェビナー トレーニング
顧客検索 拡張生成 すべてのユースケースを見る 業界別で見る すべての顧客事例を見る OpenEvidence が Zilliz Cloud で医療 AI を強化
English 日本語 한국어 Español Français Deutsch Italiano Português Русский 44.6k ログイン デモを予約 無料で始める
Loon の紹介: 変化を止めることのないベクター データ用の新しいストレージ エンジン
ページをコピー Loon のご紹介: 変化を止めることのないベクター データ用の新しいストレージ エンジン
ベクター データセットが実際に完成することはありません。
最初の問題: 長いカラムにより書き込み増幅が高価になる
2 番目の問題: 同じデータがスキャンとポイント読み取りをサポートする必要がある
3 番目の問題: データセットが 1 つのエンジン内に存在しない
Loon: 進化するベクトル データセットのための Milvus および Zilliz Cloud の背後にあるストレージ エンジン
設計 1: 適切な列グループに適切なファイル形式を使用する
設計 2: 行 ID を使用して物理ファイルを整列する
デザイン 3: マニフェストを真実の源にする
ストレージがバージョン管理されるとユーザーに何が起こるか
Loon は Milvus 3.0 ベータ版と Zilliz Vector Lakebase で利用可能です
ジリズ ベクター湖ベースでルーンを体験してみませんか
完全に試してみてください-

GenAI アプリケーション用に構築されたマネージド ベクター データベース。
これは長くて詳細なエンジニアリングの説明になるため、詳細に入る前に重要なポイントを以下に示します。
AI データセットは静的なテーブルではありません。チームが埋め込みモデルを置き換え、スパース ベクトルを追加し、キャプションを修正し、ラベルを埋め戻し、インデックスを再構築し、オフライン分析を実行すると、同じ行が変化し続けます。
従来のストレージ レイアウトは 3 つの点に分類されます。長いベクトル列によりバックフィルが高価になること、単一のファイル形式ではスキャンとポイント読み取りの両方に適切に対応できないこと、プライベート データベース ストレージにより外部パイプラインに真実の追加コピーの作成が強制されることです。
Loon は、Milvus と Zilliz Vector Lakebase の新しいストレージ エンジンです。これは、ハイブリッド ファイル形式、行 ID の配置、データセットのバージョン管理された状態を定義するマニフェストを中心に構築されています。
目標は、データを常にコピー、書き換え、再インポートすることなく、単一のベクトル データセットでオンライン検索、オフライン分析、バックフィル、圧縮、外部コンピューティングをサポートできるようにすることです。
しばらくの間、ベクトル データベースに対して合理的に聞こえる議論が 1 つありました。
従来のデータベースには、整数、文字列、JSON、BLOB、インデックスがすでに格納されています。 _vector_ タイプを追加し、その横に ANN インデックスを構築して、それで終わりにしてみてはいかがでしょうか?
初期のセマンティック検索では、これで十分に機能します。ベクトル列とインデックスは、デモ、小規模な RAG アプリケーション、または内部検索機能をサポートできます。問題は後で、データセットがテーブルとしてではなく AI データ システムのように動作し始めるときに現れます。
運用ベクトル データセットには、行、主キー、スカラー フィールド、およびクエリ可能な列が含まれます。その意味では、データベースのテーブルのように見えます。しかし、データ レイクの規模とワークフローの形状も備えています。数億件のレコードが含まれる場合があります。スパによって繰り返し読み取られ、書き換えられる

rk、Ray、DuckDB、トレーニング パイプライン、評価ジョブ、データ品質システム。
オブジェクトストレージにも依存します。ソース オブジェクトは多くの場合、S3、GCS、OSS、または別のオブジェクト ストアに残るビデオ、画像、PDF、オーディオ ファイル、または Web ドキュメントです。データベースには、参照、メタデータ、派生機能、およびインデックスが保存されます。次に、従来のストレージ モデルがファーストクラスのオブジェクトとして管理するように構築されていなかったもの、つまり、密な埋め込み、スパース ベクトル、キャプション、ベクトル インデックス、テキスト インデックス、削除ログ、統計、モデル バージョン、パーサー バージョン、外部 BLOB 参照、およびそれらすべての間のバージョン関係が追加されます。
ここで、「ベクトル列を追加するだけ」では問題が生じ始めます。問題は、データベースがベクトル バイトを格納できるかどうかではありません。多くのシステムではそれが可能です。さらに難しい問題は、ベクトル データがどのように変化するか、データがどのようにクエリされるか、AI データ スタック全体でどのように共有されるかをストレージ モデルが処理できるかどうかです。
これが、Milvus と Zilliz Vector Lakebase (Zilliz Cloud の次の進化版) 用の新しいストレージ エンジンである Loon を構築した理由です。
Loon は 3 つのアイデアで設計されています。
列の種類ごとに異なる物理フォーマットを使用します。
共有行 ID スペースを介してこれらの列を位置合わせします。
マニフェストを使用して、データセットのバージョン管理された状態を定義します。
これらの部分がなぜ重要なのかを理解するために、一般的なマルチモーダル ワークフローから始めましょう。
ベクター データセットが実際に完成することはありません。
AI チームがマルチモーダル トレーニング用のビデオ データセットを構築しているところを想像してください。
長いビデオがオブジェクト ストレージにアップロードされます。パイプラインは、シーンの変更、ショットの境界、またはタイム ウィンドウに基づいて、クリップをクリップに分割します。長すぎるクリップ、短すぎるクリップ、ぼやけているクリップ、重複しているクリップ、または低品質のクリップは除外されます。残りのクリップは美的モデルによってスコア付けされ、別のモデルによってキャプションが付けられ、視覚言語モデルによって埋め込まれ、保存されます。

n 検索、重複排除、トレーニング データ フィルタリング用のベクトル データベース。
大まかに見ると、ワークフローは単純に見えます。
ビデオ
→ クリップ
→ メタデータ
→ 美的スコア
→ キャプション
→埋め込み
→ 検索/重複排除/トレーニングデータのフィルタリング
ただし、データセットは完全な形で到着するわけではありません。
最初の週は、テーブルに含まれるのは、 Clip_id 、 video_id 、 start_offset 、およびduration のみです。
2 週目では、チームは beautiful_score を追加します。
3 週目では、キャプション モデルが実行され、各クリップにキャプションが付けられます。
4 週目では、最初のエンベディング モデルがオンラインになり、各クリップは 768 次元の CLIP エンベディングを取得します。
1 か月後、チームはモデルを切り替えて embedding_v2 をバックフィルし、現在 1024 ディメンションになりました。
2 か月後、ハイブリッド検索が要件となるため、チームはスパース ベクトル列を追加します。
3 か月後、キャプションは人間によるレビューを受け、その場で修正する必要があります。
データセットは完成しませんでした。同じ基礎となる行の新しい解釈を蓄積し続けました。
これが、ベクター データと従来のビジネス データの主要な違いの 1 つです。同じ行が何度も再処理されます。そして、スケールによって、これは不便さから​​ストレージの問題に変わります。マルチモーダル データセットは、多くの場合、数百万のレコードではなく、数億または数十億のレコードになります。 LAION-5B は、この形状に関する有用なリファレンスです。それぞれにメタデータ、キャプション、埋め込みが含まれる何十億もの画像とテキストのペアが含まれています。したがって、難しいのは最初の挿入ではありません。難しいのは、データセットが進化し始めた後に起こることすべてです。その進化により 3 つの問題が明らかになります。
最初の問題: 長いカラムにより書き込み増幅が高価になる
Parquet などの列形式は、多くの分析ワークロードに最適です。スキーマがかなり安定しており、データが書き換えよりも頻繁に読み取られる場合、これらはうまく機能します。

10 番目、スキャンは列のサブセットのみにアクセスするため、圧縮が重要です。それは、多くの分析フォーマットが最適化された世界です。
ベクトル行は分析行よりもはるかに幅が広い
TPC-H ラインアイテムは適切なベースラインです。これには、整数キー、10 進数値、日付、短い文字列、および小さなコメント フィールドの 16 列があります。非圧縮の 1 行は約 150 バイトです。圧縮後は、かなり小さくなる可能性があります。 64 MB の行グループを使用すると、ストレージ システムは数十万の行を 1 つのグループにパックできます。
ベクター データセットはそのようには見えません。
LAION スタイルの画像テキスト データセットは、現在多くの AI パイプラインが生成するものに非常に近いものです。各行には、URL、キャプション、幅、高さ、品質スコア、ラベルなどの通常のメタデータが含まれています。ただし、埋め込みが追加されると、行の物理的な形状が変化します。
768 次元の CLIP ベクトルは、fp16 では約 1.5 KB、fp32 では約 3 KB を必要とします。その 1 つの列は、TPC-H ラインアイテム行全体よりもはるかに大きくなる可能性があります。
そして、768 次元は今日の標準からすると珍しいことでもなく、大きいものでもありません。マルチモーダル パイプラインでは、1024 次元または 2048 次元の埋め込みが一般的です。 OpenAI の text-embedding-3-large は最大 3072 次元に達し、fp32 ではベクトルあたり約 12 KB になります。
多くの AI データセットでは、ベクトル列は単なるフィールドではありません。物理的には列の大部分です。それにより、スキーマ進化のコストが変わります。
ベクトル列を 1 つ追加すると、数百ギガバイトが必要になる場合があります
データセットに 1 億個のビデオ クリップがあるとします。新しい 1024 次元の fp32 埋め込み列を追加すると、約 400 GB の生のベクトル データを書き込むことになります。これには、統計、インデックス、メタデータの更新、オブジェクト ストレージのオーバーヘッド、検証、またはサービス提供パスの統合は含まれません。
チームが毎月 1 つまたは 2 つのベクトルのような列 (embedding_v2 、 sparse_vector 、または再ランク付け機能など) を追加する場合、スキーマの進化

数百ギガバイトまたはテラバイト単位で測定される定期的な daAta エンジニアリング ジョブになります。
小規模な論理更新が大規模な物理的書き換えを引き起こす可能性がある
アップデートも同様に重要です。
カラム型システムでは、通常、古いデータはその場で更新されません。削除ログには変更内容が記録され、後で圧縮によってライブ行が新しいファイルに書き換えられます。このモデルは、行が小さい場合には管理可能です。
ベクター データの場合、小規模な論理更新が大規模な物理的な書き換えを引き起こす可能性があります。
人間によるレビュー作業では、キャプション内の数百バイトしか修正できません。ただし、キャプション、デンス ベクター、スパース ベクター、およびその他の派生フィーチャが同じ物理ファイル ライフサイクルを共有している場合、システムはベクターも書き換えることになる可能性があります。論理的な変更は小さいです。物理 I/O は巨大になる可能性があります。
これは、ベクトル ストレージにおける書き込み増幅の問題です。コストがかかるのは、ベクトルが大きいことだけではありません。それは、大きな派生フィールドと小さな可変フィールドが、それらを 1 つのユニットとして扱うストレージ レイアウトによって結び付けられることが多いということです。
AI データセットの場合、バックフィルは日常的なワークロードです
従来の分析テーブルの場合、スキーマの進化はまれにのみ発生する可能性があります。 AI データセットの場合、これは日常的なことです。キャプションモデルがアップグレードされました。埋め込みモデルが置き換えられます。スパース ベクトルは後で追加されます。リランク機能が表示されます。人間のラベルが修正されました。ガバナンスタグが埋め戻されます。インデックスが再構築されます。
これらの操作は単純な追加ではありません。既存の行を頻繁に変更または拡張します。
そのため、ベクトル ストレージはスキャン スループットのみを最適化することはできません。また、バックフィルや部分的な更新をより安価にする必要もあります。
2 番目の問題: 同じデータがスキャンとポイント読み取りをサポートする必要がある
データが書き込まれた後、読み取りパスが分割されます。通常、同じベクトル データセットには、分析スキャンとポイント読み取りという 2 つの異なるアクセス パターンがあります。
分析

重要なワークロードは広範囲の圧縮スキャンを必要とします
パイプラインは次のようなフィルターを実行できます。
WHERE エステティックスコア > 0.8 かつ 期間 > 5
あるいは、オフライン分析、完全な埋め込み評価、BM25 統計、ビットマップ構築、データ品質チェック、カウント、グループバイを実行することもあります。
このパターンでは、多くの行が読み取られますが、列は数列しか読み取られません。シーケンシャル I/O、大規模な行グループ、圧縮、列プルーニング、バッチ デコード、ベクトル化された実行を好みます。
ここでは大きな行グループが役に立ちます。これにより、単一の I/O リクエストで大量の有用なデータを取得できるようになり、圧縮効率が向上し、オーバーヘッドを償却するのに十分な連続データが実行エンジンに提供されます。複数の列を一緒に読み取る場合、スキャン スループットを考慮して列を整理しておくと、ベクトル化された実行中のキャッシュ ミスを減らすことにも役立ちます。
この道では寄木細工が強い。
ANN の結果には、狭い行レベルの検索が必要です
ANN 検索で候補行 ID が返された後、システムは多くの場合、次のようなフィールドをフェッチする必要があります。
キャプション
埋め込み
リランク機能
video_uri
メタデータ
このパターンでは、読み取り行数が少なくなり、多くの場合、数百または数千行になりますが、行 ID による正確なアクセスが必要です。特定の行と列を検索し、必要なバイト範囲のみをフェッチし、いくつかのレコードを取得するためだけに行グループ全体を取得することを避けたいと考えています。

[切り捨てられた]

## Original Extract

Loon is a new storage engine for Milvus 3.0 and Zilliz Vector Lakebase, built to manage evolving vector datasets with ColumnGroups, row ID alignment, and Manifests.

AI Datasets Are Never Done. So We Built Loon. - Zilliz blog Products Zilliz Cloud
Fully managed Vector Lakebase services designed for high reliability, performance, and cost efficiency at scale.
Open-source vector database built for billion-scale vector similarity search.
Announcing Zilliz Vector Lakebase BYOC
Pricing Pricing Plan Flexible pricing options for every team on any budget
List Price View detailed list prices for every billing item
Free Tier Unleash Your Imagination with the Power of Managed Milvus
The Zilliz Cloud Developer Hub where you can find all the information to work with Zilliz Cloud
Resources Blog Guides Research Analyst Reports Webinars Trainings
Customers Retrieval Augmented Generation View all use cases View by industry View all customer stories OpenEvidence Powers Medical AI with Zilliz Cloud
English 日本語 한국어 Español Français Deutsch Italiano Português Русский 44.6k Log In Book a Demo Get Started Free
Introducing Loon: A New Storage Engine for Vector Data That Never Stops Changing
Copy page Introducing Loon: A New Storage Engine for Vector Data That Never Stops Changing
A vector dataset is never really finished.
The first problem: long columns make write amplification expensive
The second problem: the same data must support scans and point reads
The third problem: the dataset does not live inside one engine
Loon: a storage engine behind Milvus and Zilliz Cloud for evolving vector datasets
Design 1: use the right file format for the right column group
Design 2: align physical files through row IDs
Design 3: make the Manifest the source of truth
What changes for users when storage becomes versioned
Loon is available in Milvus 3.0 beta and Zilliz Vector Lakebase
Try Loon under Zilliz Vector Lakebase
Try the fully-managed vector database built for your GenAI applications.
This is a long, in-depth engineering dive, so here are the key points before we get into the details.
AI datasets are not static tables. The same rows keep changing as teams replace embedding models, add sparse vectors, revise captions, backfill labels, rebuild indexes, and run offline analysis.
Traditional storage layouts break down in three ways: long vector columns make backfills expensive, a single file format cannot serve both scans and point reads well, and private database storage forces external pipelines to create extra copies of the truth.
Loon is the new storage engine for Milvus and Zilliz Vector Lakebase. It is built around hybrid file formats, row ID alignment, and a Manifest that defines the dataset’s versioned state.
The goal is to enable a single vector dataset to support online search, offline analysis, backfills, compaction, and external compute without constantly copying, rewriting, or reimporting data.
For a while, there was one argument against vector databases that sounded reasonable.
Traditional databases already store integers, strings, JSON, blobs, and indexes. Why not add a _vector_ type, build an ANN index beside it, and call it a day?
For early semantic search, that works well enough. A vector column plus an index can support a demo, a small RAG application, or an internal search feature. The problem shows up later, when the dataset starts behaving less like a table and more like an AI data system.
A production vector dataset has rows, primary keys, scalar fields, and queryable columns. In that sense, it looks like a database table. But it also has the scale and workflow shape of a data lake. It may contain hundreds of millions of records. It is repeatedly read and rewritten by Spark, Ray, DuckDB, training pipelines, evaluation jobs, and data quality systems.
It also depends on object storage. The source objects are often videos, images, PDFs, audio files, or web documents that remain in S3, GCS, OSS, or another object store. The database stores references, metadata, derived features, and indexes. Then it adds things traditional storage models were not built to manage as first-class objects: dense embeddings, sparse vectors, captions, vector indexes, text indexes, delete logs, statistics, model versions, parser versions, external blob references, and the version relationships between all of them.
That is where “just add a vector column” starts to break down. The issue is not whether a database can store vector bytes. Many systems can. The harder question is whether the storage model can handle how vector data changes, how it is queried, and how it is shared across the AI data stack.
This is why we built Loon, the new storage engine for Milvus and Zilliz Vector Lakebase (the next evolution of Zilliz Cloud).
Loon is designed with three ideas:
Use different physical formats for different kinds of columns.
Align those columns through a shared row ID space.
Use a Manifest to define the dataset's versioned state.
To see why those pieces matter, let's start with a common multimodal workflow.
A vector dataset is never really finished.
Imagine an AI team building a video dataset for multimodal training.
A long video is uploaded to object storage. A pipeline cuts it into clips based on scene changes, shot boundaries, or time windows. Clips that are too long or too short, blurry, duplicated, or low-quality are filtered out. The remaining clips are scored by an aesthetic model, captioned by another model, embedded by a vision-language model, and stored in a vector database for search, deduplication, and training-data filtering.
At a high level, the workflow looks simple:
video
→ clips
→ metadata
→ aesthetic_score
→ caption
→ embedding
→ search / dedup / training data filtering
But the dataset does not arrive fully formed.
In the first week, the table may only contain clip_id , video_id , start_offset , and duration .
In the second week, the team adds aesthetic_score .
In the third week, a captioning model runs, and each clip gets a caption .
In the fourth week, the first embedding model goes online, and each clip gets a 768-dimensional CLIP embedding.
A month later, the team switches models and backfills embedding_v2 , now with 1024 dimensions.
Two months later, hybrid search becomes a requirement, so the team adds a sparse vector column.
Three months later, captions undergo human review and must be corrected in place.
The dataset was never completed. It kept accumulating new interpretations of the same underlying rows.
That is one of the core differences between vector data and traditional business data. The same row gets reprocessed again and again. And scale turns this from an inconvenience into a storage problem: multimodal datasets are often not millions of records but hundreds of millions or billions. LAION-5B is a useful reference for the shape — billions of image-text pairs, each with metadata, captions, and embeddings. So the hard part is not the first insert. The hard part is everything that happens after the dataset starts evolving. That evolution exposes three problems.
The first problem: long columns make write amplification expensive
Columnar formats such as Parquet are excellent for many analytical workloads. They work well when schemas are fairly stable, data is read more often than rewritten, scans only touch a subset of columns, and compression matters. That is the world for which many analytical formats were optimized.
Vector rows are much wider than analytical rows
TPC-H lineitem is a good baseline. It has 16 columns: integer keys, decimal values, dates, short strings, and a small comment field. One uncompressed row is roughly 150 bytes. After compression, it may be much smaller. With a 64 MB row group, a storage system can pack hundreds of thousands of rows into one group.
Vector datasets do not look like that.
A LAION-style image-text dataset is much closer to what many AI pipelines produce today. Each row still has ordinary metadata: a URL, a caption, width, height, quality scores, labels, and so on. But once the embedding is added, the row's physical shape changes.
A 768-dimensional CLIP vector takes about 1.5 KB in fp16 or 3 KB in fp32. That one column can be much larger than an entire TPC-H lineitem row.
And 768 dimensions are not unusual or large by today’s standards. A 1024- or 2048-dimensional embedding is common in multimodal pipelines. OpenAI’s text-embedding-3-large goes up to 3072 dimensions, which is about 12 KB per vector in fp32.
In many AI datasets, the vector column is not just another field. Physically, it is most of the row. That changes the cost of schema evolution.
Adding one vector column can mean hundreds of gigabytes
Suppose a dataset has 100 million video clips. Adding a new 1024-dimensional fp32 embedding column means writing roughly 400 GB of raw vector data. That does not include statistics, indexes, metadata updates, object storage overhead, validation, or serving-path integration.
If the team adds one or two vector-like columns every month, such as embedding_v2 , sparse_vector , or rerank features, schema evolution becomes a recurring daAta engineering job measured in hundreds of gigabytes or terabytes.
Small logical updates can trigger large physical rewrites
Updates are just as important.
In columnar systems, old data is usually not updated in place. A delete log records what changed, and compaction later rewrites live rows into new files. That model is manageable when rows are small.
With vector data, a small logical update can trigger a large physical rewrite.
A human review job may only correct a few hundred bytes in a caption. But if the caption, dense vector, sparse vector, and other derived features share the same physical file lifecycle, the system may end up rewriting the vectors too. The logical change is small. The physical I/O can be huge.
This is the write amplification problem in vector storage. The expensive part is not only that vectors are large. It is that large derived fields and small mutable fields often get tied together by a storage layout that treats them as one unit.
For AI datasets, backfill is a routine workload
For traditional analytical tables, schema evolution may occur only occasionally. For AI datasets, it is routine. Caption models are upgraded. Embedding models are replaced. Sparse vectors are added later. Rerank features appear. Human labels are corrected. Governance tags are backfilled. Indexes are rebuilt.
These operations are not simple appends. They frequently modify or extend existing rows.
That is why vector storage cannot only optimize for scan throughput. It also has to make backfills and partial updates cheaper.
The second problem: the same data must support scans and point reads
After the data is written, the read path splits. The same vector dataset typically has two distinct access patterns: analytical scanning and point reads.
Analytical workloads want wide, compressed scans
A pipeline may run filters such as:
WHERE aesthetic_score > 0.8 AND duration > 5
Or it may run offline analysis, full embedding evaluation, BM25 statistics, bitmap construction, data quality checks, counts, and group-bys.
This pattern reads many rows but only a few columns. It likes sequential I/O, larger row groups, compression, column pruning, batch decoding, and vectorized execution.
Large row groups help here. They let a single I/O request pull a large amount of useful data, improve compression efficiency, and provide the execution engine with enough contiguous data to amortize overhead. When multiple columns are read together, keeping them organized for scan throughput also helps reduce cache misses during vectorized execution.
Parquet is strong on this path.
ANN results need narrow, row-level lookups
After the ANN search returns candidate row IDs, the system often needs to fetch fields such as:
caption
embedding
rerank feature
video_uri
metadata
This pattern reads fewer rows, often hundreds or thousands, but it needs precise access by row ID. It wants to locate a specific row and column, fetch only the required byte range, and avoid pulling an entire row group just to retrieve a few records.

[truncated]
