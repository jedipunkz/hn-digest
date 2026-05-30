---
source: "https://www.modash.io/engineering/how-we-run-gemini-at-scale-across-billions-of-posts"
hn_url: "https://news.ycombinator.com/item?id=48334011"
title: "How we run Gemini at scale across billions of posts"
article_title: "How we run Gemini at scale across billions of posts"
author: "igarnedo"
captured_at: "2026-05-30T11:34:31Z"
capture_tool: "hn-digest"
hn_id: 48334011
score: 1
comments: 0
posted_at: "2026-05-30T08:31:41Z"
tags:
  - hacker-news
  - translated
---

# How we run Gemini at scale across billions of posts

- HN: [48334011](https://news.ycombinator.com/item?id=48334011)
- Source: [www.modash.io](https://www.modash.io/engineering/how-we-run-gemini-at-scale-across-billions-of-posts)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T08:31:41Z

## Translation

タイトル: 何十億もの投稿にわたって Gemini を大規模に実行する方法
説明: Modash データ エンジニアリング チームの Iván から、マルチクラウド セットアップで数十億のポストにわたって Gemini を実行する方法について聞きます。コストとスループットの最適化により、請求額が制御不能になることなく、LLM 推論を毎日何百万もの新しい入力に拡張できるようになります。

記事本文:
何十億もの投稿にわたって Gemini を大規模に実行する方法
プラットフォーム管理 すべてのコラボレーションを最初から最後まで組織的に保ちます。 Discover あなたのブランドに適したクリエイターを大規模に見つけ、精査し、つながります。追跡 インフルエンサーのキャンペーン指標とコンテンツを自動追跡します。リーチからクリック、そして売上まで。支払い 世界中のクリエイターに迅速かつ安全な支払いを送信します。統合 Shopify プロモーション コード、リンク、無料製品を数秒で送信します。電子メールでより多くのクリエイターに連絡し、コミュニケーションを 1 か所にまとめます。詳細 製品ツアー Modash に切り替える お客様 API API の概要 コンセプトからリリースまで、最も豊富なクリエイター データ API を使用してこれまでよりも迅速に実行できます。 Discovery API プログラムによるクリエイターの検索、視聴者の分析などを行います。 Raw API コメントやタグなどを含む、フィルターされていないライブ データをソーシャル アカウントから取得します。 AI 検索 AI を活用したインフルエンサー ツールとワークフローを構築する さらに多くのブランド コラボレーション ブランドとクリエイターのコラボレーションをマップする 類似クリエイターをプラットフォーム間で類似のクリエイターを発見する ドキュメント 包括的でインタラクティブなドキュメントを始めましょう。 API 価格設定 ユースケース マーケティング代理店向け AI ツールおよびエージェント向け 調査および分析向け ユースケース ギフト ギフト プログラムの管理と成長 アフィリエイト アフィリエイト プログラムの管理と成長 有料パートナーシップ 有料パートナーシップ プログラムの管理と成長 より多くの顧客 製品ツアー リソース アフィリエイト ガイド ギフト ガイド インフルエンサーの価格設定 リソース ブログ 独自のリサーチ、ハウツーなどを含む、B2C インフルエンサー プログラムに関する長文コンテンツ。ニュースレター インフルエンサー マーケティングのトピックに関する、示唆に富む短く定期的なメールを購読してください。 Eleni Zoe 著 ウェビナー インフルエンサー マーケティングの専門家によるライブおよびオンデマンドのウェビナー。ディスカッションに参加して、プログラムの新しい戦略を学びましょう。 YouTube チャンネルのヒント、コツ

、インフルエンサーとのパートナーシップの開始と拡大に役立つ分解情報。詳細 ヘルプデスク 当社のデータ コラボレーション ライブラリ 価格 デモをリクエストする ログイン デモを取得する ログイン 無料で試す エンジニアリング ブログ 2026 年 5 月 29 日 方法
[切り捨てられた]
マルチクラウド設定で数十億の入力を含む LLM を使用する
Modash では、毎日何百万もの投稿が増加するクリエイター発掘データセットを管理しています。そのパイプラインの拡大するスライスが LLM を介して実行されるようになりました。
この膨大な量の推論により、クラウドの料金と運用の複雑さが増大します。この記事では、実際に破綻することなく、数十億の入力に対して LLM を実行する方法を学びます。
AIの誇大広告にはそれだけの価値があるのか​​？ LLM には 24 時間年中無休のチャットボットとしての役割以外に実際の用途はありますか?私たちはそう考えており、昨年、LLM が顧客に提供するデータを目に見えて改善しているいくつかの本番パイプラインを出荷してきました。
これらのパイプラインのいくつかは、乱雑で多言語、マルチモーダルなソーシャル コンテンツから構造化された意味を抽出するために存在します。歴史的に、これらは正規表現ルール、キーワード リスト、および手動でコーディングされたエクストラクターのパッチワークでした。カバレッジではなく、コード行単位でスケールされました。
「これはスポンサー付きの投稿ではありません」または「このプロモーションに対して報酬は支払われません」のようなキャプションを使用します。これには、ルールが探していたすべてのキーワードが含まれていますが、まったく逆の意味になります。これを正しく処理する唯一の方法は、内容を実際に理解することです。
このような誤検知は、顧客がお金を払っている製品の知覚品質を損なうため、容認できません。
LLM は、これらをパターン マッチングのタスクではなく、言語と視覚のタスクとして再構成します。トレードオフはコスト、スループット、検証です。これがこの記事の残りの部分の内容です。
アップストリーム データは S3 上の Iceberg テーブルに存在します。各 LLM ユースケースには、対応する Airflow DAG があり、

rs PySpark ETL は、厳選されたテーブルを読み取り、推論が必要な行を抽出します。
AWS Batch ジョブは JSONL ファイルを (Gemini プロンプトを使用して) 生成し、異なる GCP バケットに保存します (コンピューティング能力を最大限に活用するためにリージョンごとに 1 つ、詳細は後述します)。パブリッシュ/サブスクライブがイベントを検出し、JSONL ファイルを Gemini Enterprise Agent プラットフォームに送信します (50% 安い Batch API を使用します)。
Gemini Enterprise (以前の Vertex ) は、ファイルのパスから使用されるモデルを読み取り、同じパーティション化戦略を使用して出力を保存します。そこから、出力 JSONL を S3 にプルし、Parquet として配置する定期的な同期ジョブを実行します。
各入力行は、LLM 出力にも存在する一意の ID によって識別されます。
最後に、データを使用する準備が整い、スケジュールされた EMR ジョブによって、顧客向けに生成されるデータが生成されます。 EMR ジョブを最適化する方法について詳しく知りたい場合は、このリンクを確認してください。
そこから、各 Parquet ファイルに対して、Airflow が 1 つの AWS Batch ジョブをトリガーし、大まかに言うと、Gemini がそれを消化できるように生のプラットフォーム データを準備します。
必要なポストデータを読み取り、負荷の高い I/O タスク (メディアのダウンロードやエンコードなど) を処理するため、並列処理する必要があり、そうしないとジョブのリソースがインフラ使用されます。
データを Gemini リクエストにカプセル化します。各投稿 (または投稿のバッチ) は、プロンプト指示および構造化された出力スキーマとともに、単一の自己完結型リクエスト ペイロードにパッケージ化されます。
これらのリクエスト ペイロードを集約し、最大 900 MB に達するとローテーションされる大きな JSONL ファイルに書き込みます。
[
{
"キー" : "XYZ" ,
「リクエスト」: {
「内容」：[
{
「パーツ」: [
{
"テキスト" : "<バッチ_リクエスト>
<post_entry> 投稿データ </post_entry>
</batch_request>"
}
]、
「役割」：「ユーザー」
}
]、
"世代構成" : {
「気性

特性" : 0 、
"maxOutputTokens" : 8192 、
"responseMimeType" : "アプリケーション/json" ,
"応答スキーマ" : {
"タイプ" : "オブジェクト" ,
"title" : "SponsoredPostBatchResponse" ,
"プロパティ" : {
"結果" : {
"タイプ" : "配列" ,
"title" : "結果" ,
"description" : "解析結果の一覧です。入力投稿ごとに 1 つの結果のみを含める必要があります。" ,
「アイテム」: {
"タイプ" : "オブジェクト" ,
"title" : "スポンサー ID" ,
"プロパティ" : {
「ピダンティックオブジェクトスキーマ」
}、
}
}
}、
}
}、
" ThinkingConfig" : { " ThinkingBudget " : 0 },
"システム命令" : {
「パーツ」: [
{
"text" : "当社のスポンサー投稿検出プロンプト"
}
】
}
}
}, ...
】
JSONL が大きい理由
900 MB という数字は意図的なもので、Gemini のハード入力上限は 1 GB です。しかし、リクエストを大きな JSONL に詰め込んでも、トークンのコストは節約できません。
同時バッチジョブ割り当て。 Gemini Enterprise では、リージョンごとに同時に実行できるバッチ ジョブの数に上限が設けられています。 1 KB ファイルと 900 MB ファイルはそれぞれ 1 つのクォータ スロットを消費します。
GCS オブジェクトの操作。 Gemini Enterprise は、入力形状をミラーリングするファイルとして結果を書き戻します。数千のマイクロ JSONL は数千のマイクロ出力ファイルになり、GCS ではそれらすべてのコストが 10 億規模ではるかに高くなります。
ダウンストリームのシンプルさ。取り込みジョブで列挙、グロブ、マージするファイルが少なくなります。
GCP リージョン全体に負荷を分散する
Gemini Enterprise AI Batch クォータは、プロジェクトごとではなく、リージョンごとです。すべての JSONL を 1 つの GCS バケットに単純にアップロードし、そのバケットと同じリージョンで実行すると、リージョンの同時実行数の上限に達し、キューに入れられ、他のリージョンがアイドル状態になる間に飢えてしまいます。
このボトルネックを回避するために、GCP トラフィック ルーティング層を中心にパイプラインを設計します。
地域エンドポイント ( europe-west4 、 us-central1 、…)。リクエストはまさにその地域に着陸します。データはそこに残り、クォータはそこでカウントされます。また、リージョンが

満席でお待ちしております。
マルチリージョンのエンドポイント (例: us 、 eu )。 Gemini Enterprise は、データを地理上の居住境界内に保ちながら、単一の地理内のリージョン間でトラフィックを透過的にシフトします。
グローバルエンドポイント。 Google が現在容量のある世界のどの地域にリクエストをルーティングする単一のエンドポイント。可用性を確保するためにレイテンシーと常駐時間を費やすことをいとわない場合、これは正しい決定です。
問題点: すべてのモデルがすべてのエンドポイント タイプをサポートしているわけではありません。
そのため、マルチリージョン サポートのないモデルの場合は、ヨーロッパのリージョン全体に GCS バケットをプロビジョニングしました (サポートされているリージョンごとに 1 つ)。バッチ ジョブは、負荷に応じて各 JSONL をこれらのバケットの 1 つにアップロードします。
まず、各リージョンの Gemini Enterprise Batch API をポーリングします。
次に、そのリージョン内で現在 QUEUED 状態にあるバッチ ジョブの数をカウントします。
ジョブが新しい​​ JSONL をアップロードする準備ができると、トラッカーは盲目的なランダムな選択ではなく、最も負荷の少ない領域を選択します。
私たちが持っている最も安価なパフォーマンスレバーはプロンプト自体です。
より良いプロンプトは、より多くの思考予算、より有能なモデル、またはより多くの推論パスよりも安価に実行されます。
いくつかの苦痛な反復を経て、以下のプラクティスが本番環境との接触を生き残ったものです。
プロンプトごとに複数の行をバッチ処理する
Gemini Enterprise AI Batch API を利用する場合、コストはトークンの量 (入力トークン + 出力 + 思考トークン) によって厳密に決定されます。モデルを通過する単一のトークンごとに料金が請求されるため、タスクを 1 つずつ実行すると、静的システム ガイドラインの全額を何度も支払う必要があります。これを解消するには、複数の行を同じプロンプトにパックして、固定命令ブロックを「キャッシュ」できるようにします (標準的なリアルタイム呼び出しの場合は、キャッシュのオプションもありますが、残念なことに、

はバッチ API 呼び出しには適用されません)。
これは、プロンプトがリクエストの半分を占めている場合にのみ有効です。マルチモーダル ジョブの場合、入力は圧倒的に主要なトークン コンシューマーです (単純な命令、膨大な入力)。そこでは、複数の行を 1 つのリクエストにバッチ処理すると、わずかながらメリットが得られます。ただし、命令ブロックが行ごとの入力よりもはるかに大きい場合、バッチ処理によりプロンプト コストが多くの行にわたって償却されます。
経験則は次のとおりです。プロンプト トークンが行ごとの入力トークンよりも多い場合はバッチです。それ以外の場合は、簡素化とデバッグ容易性のために、行ごとに 1 つのリクエストを推奨します。
バッチ入力レート: 1,000,000 トークンあたり 0.15 ドル
バッチ出力レート: 1,000,000 トークンあたり 1.25 ドル
システム指示プロンプト: 1,189 トークン
投稿あたりのデータ サイズ: 300 入力トークン
投稿ごとの出力トークン: 100 出力トークン
これらの正確な数字をさまざまなバッチ サイズにマッピングすると、プロンプトあたり 10 ～ 20 件の投稿が最終的なスイート スポットである理由がわかり、収穫逓減の法則が明らかになります。 1 件の投稿から 20 件の投稿に移行することで、数百万のレコードに対して数千ドルを節約できます。ただし、20 を超えると、平坦な停滞期に突入します。ラインに 50 または 100 のポストを詰め込むと、金銭的利益は 2% 未満ですが、重大なエンジニアリング リスクが生じます。
リクエストごとの出力トークン制限。 Gemini モデルの出力トークンの上限は 65,000 です。入力トークンにも制限がありますが、その上限は 1M です。
「すべての入力投稿 ID に対して結果オブジェクトを返す必要がある」や「リスト内の各アカウントを個別に処理する。接続を推測しない」などの指示を追加して、バッチの各項目が個別に処理されるようにします。
結果カウント == 入力カウント を常に検証します。非常に大きなバッチでは、順序のシャッフルが多くなり、場合によっては結果オブジェクトが削除されることがあります。
LLM はデフォルトで遅延します。モデルに単純なカテゴリを尋ねると、

l ラベルまたは単純な 2 値判定の場合、実際の内部推論を実行することなく (しばしば幻覚や怠惰で間違った推測につながる)、その答えを即座に出力します。非推論モデルの場合でも、ラベル付けする前に証拠を要約することを強制すると (構造化された思考連鎖)、推測ができなくなります。これにより、モデルは判定にコミットする前に入力全体を処理できるようになります。
構造化された思考連鎖パターンを適用するには、モデルが最終分類 (スポンサー付き投稿か?) に達する前に、モデルに Analysis_trace フィールドへの入力を強制し、構造的に尻からの攻撃を防ぎます。これらのフィールドは、モデルが分類を発行する前に推論を言語化することを純粋に強制するために存在します。ただし、一時的な修正ほど永続的なものはないため、デバッグ中や LLM の推論を理解しようとするときに、これらを追加することは非常に役立ちました。
しっかりと蓋をしてください (20 単語)。長い形式の思考連鎖はトークンを無駄にします。
これらをスキーマ定義の最初に置きます。モデルは宣言された順序でフィールドに入力するため、推論は分類の前に着地する必要があります。
デバッグに使用してください。結果が間違っているように見える場合、analytics_trace は、それがルール解釈の問題 (プロンプトを修正する) なのか、それとも問題なのかを示します。

[切り捨てられた]

## Original Extract

Hear from Iván from the Modash Data Engineering team about how they run Gemini across billions of posts in a multi-cloud setup. Learn the cost and throughput optimizations that let them scale LLM inference to millions of new inputs daily — without the bill spiraling out of control.

How we run Gemini at scale across billions of posts
Platform Manage Keep every collaboration organized from start to finish. Discover Find, vet, and connect with the right creators for your brand, at scale. Track Auto-track influencer campaign metrics & content. From reach, to clicks, to sales. Pay Send fast, secure payments to creators worldwide. Integrations Shopify Send promo codes, links, and free products in seconds. Email Reach more creators and keep comms in one place. More Product tour Switch to Modash Customers API API Overview From concept to launch, faster than ever using the richest creator data APIs. Discovery API Programmatically find creators, analyze audience, and more. Raw API Get live, unfiltered data from social accounts, including comments, tags, and more! AI search Build AI powered influencer tools and workflows More Brand collaborations Map brand and creator collaborations Creator lookalikes Discover similar creators across platforms Documentation Get started with our comprehensive and interactive documentation. API Pricing Use cases For marketing agencies For AI tools and agents For research and analysis Use cases Gifting Manage & grow your gifting programs Affiliates Manage & grow your affiliate programs Paid partnerships Manage & grow paid partnerships programs More Customers Product tour Resources Affiliate guide Gifting guide Influencer pricing Resources Blog Long-form written content on B2C influencer programs including original research, how-tos and more. Newsletter Subscribe for short, regular, thought provoking emails on influencer marketing topics. Written by Eleni Zoe Webinars Live & on-demand webinars with influencer marketing pros. Join the discussion and learn new strategies for your program. YouTube channel Tips, tricks, and teardowns that will help you start and scale your influencer partnerships. More Helpdesk Our data Collabs library Pricing Request a demo Login Get a demo Log in Try for free Engineering blog May 29, 2026 How
[truncated]
Using LLMs with billions of inputs in a multi-cloud setup
At Modash we sit on top of a creator-discovery dataset that grows by millions of posts every day. A growing slice of that pipeline now runs through LLMs.
This massive volume of inference adds up on our cloud bills and our operational complexity. In this article you will learn how we actually run an LLM against billions of inputs without going broke .
Is the AI hype worth it? Do LLMs have any real use beyond being a 24/7 chatbot? We think so, and over the last year we’ve shipped several production pipelines where LLMs are visibly improving the data we deliver to our customers.
Several of those pipelines exist to extract structured meaning from messy, multilingual, multimodal social content. Historically these were patchworks of regex rules, keyword lists, and hand-coded extractors. They scaled in lines of code, not in coverage.
Take a caption like “This is not a sponsored post” or “I’m not being paid for this promotion” : it contains every keyword the rules were looking for, but means the exact opposite. The only way to handle it correctly is to actually understand the content.
Those false positives are unacceptable as they erode the perceived quality of a product customers are paying for .
LLMs reframe these as language and vision tasks instead of pattern-matching ones. The tradeoff is cost, throughput, and validation — which is what the rest of this article is about.
Our upstream data lives in Iceberg tables on S3 . Each LLM use case has a corresponding Airflow DAG that triggers PySpark ETL’s that read our curated tables and extracts the rows that need inference.
The AWS Batch jobs generate the JSONL files (with the Gemini prompts) and stores them in different GCP buckets (one per region to leverage as much as possible compute capacity, more on that below), a pub/sub detects the event and send the JSONL file to a Gemini Enterprise Agent Platform (using the Batch API as it’s 50% cheaper ).
Gemini Enterprise (formerly Vertex ) will read the model to be used from the path of the file and will store the output using the same partitioning strategy. From there, we run a periodic sync job that pulls those output JSONLs into S3 and lands them as Parquet.
Each input row is identified by a unique ID that is also present in the LLM output.
Finally, the data is ready to be used and our scheduled EMR jobs generate the data that we produce for our customers. Check this link if you want to lear more about how we optimize our EMR jobs.
From there, for each Parquet file, Airflow triggers out one AWS Batch job that, at a high level, prepares our raw platform data so Gemini can digest it:
Reads the necessary post data and handles heavy I/O tasks: (like downloading and encoding media) so it has to be parallel or the job’s resources will be infra-utilized.
Encapsulates the data into Gemini requests: Each post (or batch of posts) is packaged into a single, self-contained request payload along with its prompt instructions and structured output schemas.
Aggregates these request payloads and writes them into a large JSONL file that rotates when it hits ~900 MB.
[
{
"key" : "XYZ" ,
"request" : {
"contents" : [
{
"parts" : [
{
"text" : "<batch_request>
<post_entry> POST DATA </post_entry>
</batch_request>"
}
],
"role" : "user"
}
],
"generationConfig" : {
"temperature" : 0 ,
"maxOutputTokens" : 8192 ,
"responseMimeType" : "application/json" ,
"responseSchema" : {
"type" : "object" ,
"title" : "SponsoredPostBatchResponse" ,
"properties" : {
"results" : {
"type" : "array" ,
"title" : "Results" ,
"description" : "List of analysis results. Must contain exactly one result per input post." ,
"items" : {
"type" : "object" ,
"title" : "SponsorIdentification" ,
"properties" : {
"PYDANTIC OBJECT SCHEMA"
},
}
}
},
}
},
"thinkingConfig" : { "thinkingBudget" : 0 },
"systemInstruction" : {
"parts" : [
{
"text" : "OUR SPONSORED POST DETECTION PROMPT"
}
]
}
}
}, ...
]
Why JSONLs are big
The 900 MB number is deliberate — Gemini’s hard input cap is 1 GB . But packing requests into large JSONLs does not save us token money.
Concurrent-batch-job quota . Gemini Enterprise enforces a cap on how many batch jobs we can have running per region simultaneously. A 1 KB file and a 900 MB file each consume one quota slot.
GCS object operations . Gemini Enterprise writes results back as files mirroring the input shape. Thousands of micro-JSONLs become thousands of micro-output-files, all of which cost way more on GCS at billion-scale.
Downstream simplicity . Fewer files for our ingestion job to enumerate, glob, and merge.
Spreading load across GCP regions
Gemini Enterprise AI Batch quotas are per-region , not per-project. If we naively upload all our JSONLs into one GCS bucket and run them in the same region as the bucket, we’ll hit the regional concurrency cap, queue up, and starve while other regions sit idle.
To bypass this bottleneck, we architect our pipeline around GCP traffic routing tiers:
Regional endpoints ( europe-west4 , us-central1 , …). Request lands in exactly that region. Data stays there, quota is counted there, and if the region is full we wait.
Multi-region endpoints (e.g. us , eu ). Gemini Enterprise transparently shifts traffic between regions inside a single geography while keeping data within the geography’s residency boundary.
The global endpoint . A single endpoint where Google routes requests to whichever region in the world currently has capacity. It’s the right decision when we are willing to spend latency and residency to get availability.
The catch: not every model supports every endpoint type.
So, for the models without multi-region support , we provisioned GCS buckets across European regions (one per supported region ). The batch job’s upload each JSONL to one of those buckets depending on their load:
First polls each region’s Gemini Enterprise Batch API.
Then counts how many of our batch jobs in that region are currently in QUEUED state.
When the Job is ready to upload a new JSONL, the tracker picks the least-loaded region instead of a blind random choice.
The cheapest performance lever we have is the prompt itself.
Better prompts run cheaper than more thinking budget, more capable models, or more inference passes .
After several painful iterations, the practices below are the ones that survived contact with production.
Batch multiple rows per prompt
When utilizing the Gemini Enterprise AI Batch API, costs are determined strictly by token volume: input tokens + output + thinking ones. Because you are billed for every single token that passes through the model, executing tasks one-by-one forces you to pay the full price of your static system guidelines over and over again. To eliminate this, multiple rows can be packed into the same prompt, allowing us to “cache” the fixed instruction block (for standard real-time calls there is also the option of Caching but, sadly, that’s not applicable to Batch API calls).
This only pays off when the prompt is the heavy half of the request . For our multimodal jobs the input is by far the dominant token consumer (simple instructions, huge input). There, batching multiple rows into one request would have marginally benefits. But, for cases where the instruction block is much larger than the per-row input, batching amortizes the prompt cost across many rows.
The rule of thumb is: batch when prompt tokens dominate per-row input tokens; otherwise prefer one request per row for simplicity and debuggability.
Batch Input Rate: $0.15 per 1,000,000 tokens
Batch Output Rate: $1.25 per 1,000,000 tokens
System instructions prompt: 1,189 tokens
Data Size per Post: 300 input tokens
Output tokens per Post: 100 output tokens
When we map these exact numbers out across different batch sizes, the law of diminishing returns becomes obvious as it illustrates why 10 to 20 posts per prompt is our definitive sweet spot . Moving from 1 to 20 posts saves us thousands of dollars over millions of records. However, scaling past 20 pushes us straight into a flat plateau. Jamming 50 or 100 posts into a line introduces substantial engineering risk for less than a 2% financial gain.
Output token limit per request . Gemini models caps at ~65K output tokens . Input tokens also have a limit but its ~1M.
Add instructions like “you must return a result object for EVERY input post ID” and “Process each account in the list independently; do not infer connections” to ensure that each item of the batch is processed independently.
Always validate result count == input count . Very large batches see more order-shuffling and occasional dropped result objects.
LLMs are lazy by default. If we ask a model for a straightforward categorical label or a simple binary verdict, it will happily emit that answer instantly without performing any real internal reasoning (frequently leading to hallucinations or lazy, incorrect guesses). Even for non-reasoning models, forcing them to summarize evidence before labeling (structured Chain-of-Thought) prevents them from guessing. It allows the models to process the entire input before committing to a verdict.
To apply the structured Chain-of-Thought pattern, we force the model to populate an analysis_trace field before it reaches the final classification ( is a sponsored post? ), we structurally prevent it from shooting from the hip. Those fields exist purely to force the model to verbalize the reasoning before it emits the classification. But, as nothing is more permanent as a temporary fix, adding those was really useful while debugging and while trying to understand the LLM reasoning.
Cap them tightly (20 words). Long-form chain-of-thought wastes tokens.
Put them first in the schema definition . The model fills fields in the order they’re declared, so the reasoning has to land before the classification.
Use them for debugging . When a result looks wrong, analysis_trace tells whether it was a rule-interpretation issue (fix the prompt) or a

[truncated]
