---
source: "https://ethereal-agents.space/blog/launching-arxiv-scholar.html"
hn_url: "https://news.ycombinator.com/item?id=48554932"
title: "Building a Research Search Engine for 5,600 AI Papers on a $1 Budget"
article_title: "Building ArXiv Scholar: A Production RAG Pipeline — Without a GPU, Without LangChain, and Without a Budget"
author: "dubeyaayush07"
captured_at: "2026-06-16T13:55:22Z"
capture_tool: "hn-digest"
hn_id: 48554932
score: 1
comments: 1
posted_at: "2026-06-16T13:24:20Z"
tags:
  - hacker-news
  - translated
---

# Building a Research Search Engine for 5,600 AI Papers on a $1 Budget

- HN: [48554932](https://news.ycombinator.com/item?id=48554932)
- Source: [ethereal-agents.space](https://ethereal-agents.space/blog/launching-arxiv-scholar.html)
- Score: 1
- Comments: 1
- Posted: 2026-06-16T13:24:20Z

## Translation

タイトル: 1 ドルの予算で 5,600 件の AI 論文のための研究検索エンジンを構築する
記事のタイトル: ArXiv Scholar の構築: プロダクション RAG パイプライン — GPU なし、LangChain なし、予算なし
説明: 6 つの無料 Google Colab アカウント、予算ゼロのインフラストラクチャ、抽象化フレームワークを使用せずに、5,600 件の学術論文を超える実稼働 RAG システムを構築した方法。 98.8% の True Recall@20 を達成。

記事本文:
あ
ArXiv 研究者
ホーム
検索してみる
GitHub
このページについて
これを構築した理由
建築
データの取得
レイアウトを考慮したチャンク化
デュアル埋め込みパイプライン
6-コラボ戦略
Qdrantストレージ
インテリジェントな検索
リランカーの評価
LLM とストリーミング
評価
ベンチマーク
私たちが学んだこと
自分で試してみる
ライブ — 5.6K の論文がインデックス付き
2026年6月
ArXiv Scholar の構築: プロダクション RAG パイプライン — GPU なし、LangChain なし、予算なし
6 つの無料 Colab アカウント、無料層 Qdrant、およびゼロ抽象化フレームワークを使用して、5,600 件の学術論文を超えるエンドツーエンドの検索システムを構築し、98.8% True Recall@20 を達成した方法を説明します。
私たちは、arXiv のコーパスの厳選された 5,600 枚の論文のサブセットに対してゼロから RAG パイプライン (抽象化フレームワークなし) を構築し、並行して実行される 6 つの無料 Google Colab アカウントを使用してすべてを処理し、結果を Qdrant の無料クラウド層にアップロードし、インテリジェントなクエリ ルーティング、HyDE、クエリ分解、ハイブリッド密 + スパース検索を備えた軽量ストリーミング フロントエンドを出荷しました。運用インフラストラクチャ全体の費用はちょうど 1 ドルでした (ドメインの場合)名前）。
私たちは、研究者が実際に考えている方法で学術論文を検索したいと考えました。タイトルとキーワードを照合するのではなく、「2023 年以降に発表されるロングコンテキストの注意メカニズムの最先端の技術は何ですか?」のような本当の質問をすることです。そして、実際の arXiv 出版物からの回答を引用して、背景に戻ります。
そこで私たちは ArXiv Scholar を構築しました。これは、arXiv から何千もの学術論文を取り込み、解析、チャンク、埋め込み、検索するエンドツーエンドの検索拡張生成 (RAG) システムです。ラングチェーンはありません。本番環境には GPU はありません。有料のインフラストラクチャはありません。
この投稿は、それを構築する正直なストーリーです。何がうまくいったのか、何がうまくいかなかったのか、そして予算ゼロのプロジェクトが 98.8% 真実であることを達成したエンジニアリングのトリックです。

5,600 件以上の論文を高精度に再ランキングした Recall@20。
毎週、何千もの新しい論文が arXiv に掲載されます。研究者はキーワード検索、Twitter スレッド、リストを手動でスクロールして関連する研究を見つけます。 arXiv 上の従来の検索 (arXiv 独自の検索を含む) は、基本的なテキスト検索を使用してタイトルと要約を照合します。それは概念を理解していません。
私たちは簡単な質問をしました。arXiv に平易な英語で質問して、実際の論文から引用された合成された回答が返って来るとしたらどうしますか?
問題は次のような制約でした。
コンピューティング予算はゼロ。 AWS、GCP、レンタル GPU はありません。カスタム ドメインの合計請求額はちょうど 1 ドルでした。
高レベルのフレームワークはありません。私たちは、LangChain や LlamaIndex ではなく、完全なアーキテクチャ制御を必要としていました。Python、生の API 呼び出し、および各バイトが何を行っているかを理解するだけでした。
すべてを無料枠で。処理用の無料の Colab、ベクター ストレージ用の無料の Qdrant Cloud、GCS からの無料の arXiv データ、Hugging Face Spaces でホストされている API、GitHub Pages のフロントエンド、およびルーティング用の Cloudflare 無料枠。
これらの制約は制限ではなく、設計パラメータでした。彼らは私たちに、あらゆる層でエンジニアリング上の慎重な決定を下すことを強いました。
システムは、オフライン (Colab 内) で実行される取り込みパイプラインと、ライブ クエリを提供する取得パイプラインの 2 つの分離された部分に分割されます。それぞれの決定について見ていきましょう。
1. データ取得: 1.4TB の科学への無料アクセス
ArXiv は、その出版物アーカイブ全体をパブリック Google Cloud Storage バケット ( arxiv-dataset ) としてミラーリングします。これまでにアップロードされたすべての論文 (300 万を超える PDF、約 1.4 TB) は、匿名の GCS 読み取り経由で自由にアクセスできます。
当社の ArxivUnifiedEngine は、ステートフルでクラッシュセーフなバッチ ダウンローダーです。 ev後にディスクに永続化されたJSONカーソルを使用して進行状況を追跡します。

単一ファイル:
プロセスがバッチの途中でクラッシュした場合、正確に次のファイルから再起動されます。重複やギャップはありません。エンジンはシームレスに月の境界を越えて ( 2604 → 2605 )、現在に追いつくと履歴バックフィルからライブモードに移行することもあります。
キュレーションの決定: パイプラインは 300 万件の論文すべてを取り込むことができますが、無料枠の Qdrant は最大 5,600 件の論文相当の埋め込みを快適に保持できます。そこで、4 段階のマニフェスト フィルターを構築しました。
論文は 2022 年 1 月以降に更新され、コア CS カテゴリ ( cs.AI 、 cs.CL 、 cs.IR 、 cs.LG 、 cs.SE ) に属している必要があります。
クロスリストされた医学、物理学、純粋な数学の論文を除外するための積極的なアンチノイズ フィルタリング
含めるには、VIP ツール (vLLM、LangChain など) への言及、または 3 つ以上の AI トピック グループにわたる高密度のキーワード マッチが必要です
予算上限はちょうど 5,600 論文で、関連性と最新性によってランク付けされます
このマニフェストはコスト削減策であり、技術的な制限ではありません。それを削除すると、同じパイプラインが何百万ものデータを取り込みます。
2. ドクリングによるレイアウトを意識したチャンク化
これは、ほとんどの RAG パイプラインがサイレントに失敗する場所です。デフォルトのアプローチ (500 文字ごとに分割) では、学術論文の意味構造が破壊されます。最終的には、方程式の途中から開始するチャンクができたり、テーブルが半分に分割されたり、セクション ヘッダーがコンテンツから分離されたりすることになります。
文書を視覚的に理解するために、IBM の Docling ライブラリを使用します。 PDF をフラットな文字列として扱う代わりに、Docling はレイアウトを理解します。
ヘッダーが何であるかを認識し、それを後続の段落にバインドします。
単一のチャンク内でテーブルをそのままの状態に保ちます
リスト構造とコードブロックを認識します
セマンティック要素は、下限チャンク凝集サイズ ( target_chunk_size=1000 ) に達するまでバッファーに蓄積され、その後チャンクが生成されます。すべてのチャンクには、グローバルな続きの先頭に論文のタイトルが付加されます

ext — 「提案された方法」に関するテキストがどの論文から来たのかを参照していないという古典的な「孤立したチャンク」問題を解決します。
チャンクの凝集性の影響: 当初はチャンク サイズの下限がなかったため、小さなチャンクが多すぎました。この target_chunk_size=1000 を強制することで、700 件の論文のデータセットで実験を実行し、大幅な改善が見られました。
OCR フォールバック: 当初はすべての PDF に OCR を使用していましたが、これにより処理時間が大幅に増加しました。私たちは、学術論文 (PDF を作成するために画像をスキャンしない) では、ほとんどの場合、テキストがメタデータにネイティブに存在することに気づきました。そこで、デフォルトで OCR を無効にし、厳密にフォールバックとして保持しました。
ベンチマークの違いは明らかでした。
デフォルトの OCR の場合: PDF あたりの平均時間は 31.10 秒でした
OCR なし (フォールバックのみ): PDF あたりの平均時間が 21.12 秒に短縮され、取り込み時間が約 32% 節約されました。
古い arXiv 論文 (または特定の LaTeX エンジンでコンパイルされた論文) では、内部フォント エンコーディングが壊れています。テキストは視覚的には適切に表示されますが、意味不明なものとして抽出されます。私たちのチャンカーはこれを自動的に検出し、OCR を有効にして再実行します。
レイアウト ブロックが max_chunk_size (1,500 文字) を超えると、システムは 200 文字のオーバーラップを持つスライディング ウィンドウ チャンカーに動的にフォールバックします。これにより、他のすべての Docling の品質を維持しながらデータが切り捨てられることがなくなります。
これが重要な理由: レイアウトを意識したチャンク化は計算コストが高くなります。Docling はすべての PDF で完全な文書理解モデルを実行します。これが、取り込みに GPU コンピューティングが必要な主な理由です。しかし、品質の違いは劇的です。セマンティック境界を尊重するチャンクは、スマートな結合制限とターゲットを絞った OCR と組み合わせることで、数分の 1 のコストで単純なテキスト分割よりも大幅に優れた埋め込みを生成します。
3.

埋め込み: BGE-M3 + BM25 デュアル パイプライン
私たちは、高密度埋め込みのために BAAI/bge-m3 を選択しました。これは、MTEB リーダーボードの上位近くに常にランク付けされている 1024 次元の多言語モデルです。スパース ベクトルの場合、Qdrant/bm25 を使用して、密なモデルでは見逃されるキーワードの正確な一致 (ライブラリ名、特定の頭字語、著者名) をキャプチャします。
Recall@100 による埋め込みボトルネックの診断: 最初はより小さな埋め込みモデルを実験しましたが、検索結果は不十分でした。ボトルネックを特定するために、Recall@100 (上位 100 件の結果のどこかに正しいチャンクを配置するモデルの能力) を測定しました。スコアはひどいものだった。これは重要な洞察でした。関連するチャンクがトップ 100 にすら入っていない場合、埋め込みモデルには、複雑な科学テキストを理解してマッピングするための意味論的な能力が欠けていることを意味します。ダウンストリームの再ランカーは最初から取得されなかったチャンクを再ソートできないため、この Recall@100 の低さは、BGE-M3 のようなより大規模で高次元のモデルにアップグレードする必要があることを決定的に証明しました。
重要なアーキテクチャ上の洞察は、デュアル バックエンド設計です。
どちらのバックエンドも、まったく同じ 1024 次元の BGE-M3 ベクトルを生成します。ここで決定を下したのはコストでした。ライブ クエリの処理に GPU をレンタルすると、予算ゼロの制約が破られるため、完全に無料の CPU 層で実行されるように取得バックエンドを設計しました。操作上の違いは次のようになります。
これは、取り込みパイプラインが Colab GPU を最大限に活用している間、実稼働 Docker イメージが無駄のない状態を維持する (複数 GB の PyTorch インストールがない) ことを意味します。
4. 6-Colab 並列処理戦略
ここで事態は厄介になります。 Docling のレイアウト分析と BGE-M3 埋め込みを通じて 5,600 の学術 PDF を処理するには、大量の計算が必要です。単一の Colab セッション (無料の T4 GPU を使用した場合でも) には数日かかり、Colab は強制終了します。

約12時間後のセッション。
私たちのソリューションは、並行して実行されている 6 つの無料の Google Colab アカウントに作業を分散することです。
各 Colab セッションでは、 --embedding-batch-size 128 と --colab-gpu フラグを指定して、batch_gcs_to_drive.py スクリプトを実行しました。これにより、Docling のコンバーターがオーバーライドされ、4 つのスレッドで CUDA アクセラレーションが使用されます。
耐衝撃性は非常に重要でした。 Colab セッションがランダムに切断されます。私たちのスクリプトは、JSONL 出力を Google ドライブにコピーすることにより、50 ドキュメントごとにチェックポイントを作成します。セッションが終了すると、最後のチェックポイントを指す --start-paper を使用して再起動します。アップロード スクリプトは UUID-v5 重複排除を使用するため、同じ用紙を再処理しても問題はありません。
JSONL 形式は、移植可能な中間表現として機能します。
これにより、取り込みがストレージから完全に切り離されます。 6 つの Colab アカウントは JSONL を生成します。別のアップロード スクリプト ( import_remote_qdrant_Parallel.py ) は、8 つのスレッドと自動再試行ロジック (バッチごとに指数バックオフで最大 10 回の試行) を使用して並行して Qdrant にプッシュします。
制約が私たちに教えてくれたこと: 迅速に反復することはできませんでした。完全な再処理を実行するたびに、Colab セッションの調整に数時間かかりました。このため、パイプラインを上流に配置する必要があり、「再実行だけ」に頼るのではなく、クラッシュ セーフティ、チェックポイント、冪等アップロードに多額の投資をする必要がありました。
5. ストレージ: 無料階層 Qdrant クラウド
Qdrant を選択したのは、基本的に Qdrant が当社のアーキテクチャに技術的に最適であったためです。コストだけが要因ではありませんでした。これは、Rust で書かれた高性能で堅牢なベクター データベースであり、開発/テスト (Docker 経由) のためにローカルでシームレスに実行し、運用のためにクラウドに移行できます。
具体的には、Qdrant は 4 つの重要な利点を提供しました。
ネイティブのマルチベクトルのサポート — 各ポイントには、密ベクトル (1024 次元コサイン) と名前付きスパース ベクトル ( bm25 ) sim の両方が格納されます。

同時に
サーバー側の融合 — Qdrant は、1 回のラウンドトリップで両方のベクトル空間にわたってプリフェッチ クエリを実行できます。
Rust に裏付けられた HNSW パフォーマンス — 高度に最適化された Hierarchical Navigable Small World (HNSW) グラフを利用し、限られたハードウェアでも非常に高い 1 秒あたりのリクエスト (RPS) スループットと 25 ミリ秒未満のレイテンシーを実現します。
無料のクラウド利用枠 — 5,600 の論文コーパスに十分な量
すべてのチャンクは、SHA-256 コンテンツ ハッシュから派生した決定的な UUID-v5 を取得します。これにより、upsert が冪等になります。同じチャンクを 2 回アップロードすると、Qdrant は重複ではなく上書きします。
また、コレクションの作成時に、metadata.year にペイロード インデックスを構築します。これにより、Qdrant がベクトル検索中 (後ではなく) に年ベースのフィルターを適用できるようになります。これは、クエリ分解パイプラインにとって重要です。
6. 取得: インテリジェント ルーティングとアダプティブ RAG
画一的な検索の代わりに、構造に基づいてさまざまな戦略を介してクエリをルーティングするカスタム パイプラインを構築しました。
データベース クエリが実行される前に、ルーターはユーザーの意図を分類します。
ルーターは、メタデータ パターンのハード正規表現オーバーライドを使用して、密なクエリの埋め込みに対して事前トレーニングされた分類子を使用します。オーバーライドは意図的です。私たちは時間的制約について ML 分類を信頼していません。

[切り捨てられた]

## Original Extract

How we built a production RAG system over 5,600 academic papers using 6 free Google Colab accounts, zero-budget infrastructure, and no abstraction frameworks. Achieving 98.8% True Recall@20.

A
ArXiv Scholar
Home
Try Search
GitHub
On This Page
Why We Built This
Architecture
Data Acquisition
Layout-Aware Chunking
Dual Embedding Pipeline
6-Colab Strategy
Qdrant Storage
Intelligent Retrieval
Evaluating the Re-Ranker
LLM & Streaming
Evaluation
Benchmarks
What We Learned
Try It Yourself
Live — 5.6K Papers Indexed
June 2026
Building ArXiv Scholar: A Production RAG Pipeline — Without a GPU, Without LangChain, and Without a Budget
How we built an end-to-end retrieval system over 5,600 academic papers using 6 free Colab accounts, free-tier Qdrant, and zero abstraction frameworks — achieving 98.8% True Recall@20.
We built a from-scratch RAG pipeline (no abstraction frameworks) over a 5,600-paper curated subset of arXiv's corpus, processed everything using 6 free Google Colab accounts running in parallel, uploaded the results to Qdrant's free cloud tier, and shipped a lightweight streaming frontend with intelligent query routing, HyDE, query decomposition, and hybrid dense+sparse search — our entire production infrastructure cost exactly $1 (for the domain name).
We wanted to search academic papers the way researchers actually think — not keyword-matching against titles, but asking real questions like "What is the state of the art for long-context attention mechanisms published after 2023?" and getting back grounded, cited answers from actual arXiv publications.
So we built ArXiv Scholar : an end-to-end Retrieval-Augmented Generation (RAG) system that ingests, parses, chunks, embeds, and searches thousands of academic papers from arXiv. No LangChain. No GPU in production. No paid infrastructure.
This post is the honest story of building it — what worked, what didn't, and the engineering tricks that made a zero-budget project achieve 98.8% True Recall@20 with high-precision reranking over 5,600 papers.
Every week, thousands of new papers appear on arXiv. Researchers rely on keyword searches, Twitter threads, or manually scrolling through listings to find relevant work. Traditional search over arXiv — including arXiv's own search — matches against titles and abstracts using basic text retrieval. It doesn't understand concepts.
We asked a simple question: What if you could ask arXiv a question in plain English and get back a synthesized, cited answer from the actual papers?
The catch was our constraints:
Zero compute budget. No AWS, no GCP, no rented GPUs. Our total bill was exactly $1 for the custom domain.
No high-level frameworks. We wanted full architectural control — no LangChain, no LlamaIndex — just Python, raw API calls, and an understanding of what every byte was doing.
Free-tier everything. Free Colab for processing, free Qdrant Cloud for vector storage, free arXiv data from GCS, API hosted on Hugging Face Spaces, frontend on GitHub Pages, and Cloudflare free-tier for routing.
These constraints weren't limitations — they were design parameters. They forced us to make thoughtful engineering decisions at every layer.
The system is split into two decoupled halves: an ingestion pipeline that runs offline (in Colab), and a retrieval pipeline that serves live queries. Let's walk through each decision.
1. Data Acquisition: Free Access to 1.4TB of Science
ArXiv mirrors its entire publication archive as a public Google Cloud Storage bucket ( arxiv-dataset ). Every paper ever uploaded — over 3 million PDFs, roughly 1.4TB — is freely accessible via anonymous GCS reads.
Our ArxivUnifiedEngine is a stateful, crash-safe batch downloader. It tracks progress with a JSON cursor persisted to disk after every single file:
If the process crashes mid-batch, restart picks up from the exact next file. No duplicates, no gaps. The engine seamlessly rolls over month boundaries ( 2604 → 2605 ) and even transitions from historical backfill to live-mode when it catches up to the present.
The curation decision: While the pipeline can ingest all 3 million papers, free-tier Qdrant comfortably holds ~5,600 papers worth of embeddings. So we built a 4-stage manifest filter:
Papers must be updated after January 2022 and belong to core CS categories ( cs.AI , cs.CL , cs.IR , cs.LG , cs.SE )
Aggressive anti-noise filtering to exclude cross-listed medical, physics, and pure math papers
Inclusion requires mentions of VIP tools (vLLM, LangChain, etc.) OR dense keyword matches across 3+ AI topic groups
Budget cap at exactly 5,600 papers, ranked by relevance tier and recency
This manifest is a cost-saving measure, not a technical limitation. Remove it, and the same pipeline ingests millions.
2. Layout-Aware Chunking with Docling
This is where most RAG pipelines fail silently. The default approach — split every 500 characters — destroys the semantic structure of academic papers. You end up with chunks that start mid-equation, split a table in half, or separate a section header from its content.
We use IBM's Docling library for visual document understanding. Instead of treating a PDF as a flat string, Docling understands the layout:
It knows what a header is and binds it to the paragraph that follows
It keeps tables intact within a single chunk
It recognizes list structures and code blocks
We accumulate semantic elements into a buffer until they reach a lowerbound chunk cohesion size ( target_chunk_size=1000 ), then yield a chunk. Every chunk gets the paper's title prepended for global context — solving the classic "orphaned chunk" problem where a piece of text about "the proposed method" has no reference to what paper it came from.
The impact of chunk cohesion: Initially, we didn't have a lowerbound on chunk size, which resulted in too many small chunks. By enforcing this target_chunk_size=1000 , we ran an experiment on a dataset of 700 papers and saw a massive improvement:
The OCR fallback: We initially used OCR for all PDFs, but this increased processing time significantly. We realized that for academic papers (where people don't scan images to create PDFs), the text is almost always natively present in the metadata. So we disabled OCR by default and kept it strictly as a fallback.
The benchmark difference was stark:
With default OCR: Avg Time per PDF was 31.10 s
Without OCR (Fallback only): Avg Time per PDF dropped to 21.12 s , saving roughly 32% of ingestion time.
Older arXiv papers (or those compiled with certain LaTeX engines) have broken internal font encodings. The text renders fine visually but extracts as gibberish. Our chunker detects this automatically and re-runs with OCR enabled:
When a layout block exceeds our max_chunk_size (1,500 chars), the system dynamically falls back to a sliding window chunker with 200-character overlap — ensuring we never truncate data while maintaining Docling's quality for everything else.
Why this matters: Layout-aware chunking is computationally expensive — Docling runs a full document understanding model on every PDF. This is the primary reason we needed GPU compute for ingestion. But the quality difference is dramatic: chunks that respect semantic boundaries, combined with smart cohesion limits and targeted OCR, produce significantly better embeddings than naive text splits at a fraction of the cost.
3. Embedding: The BGE-M3 + BM25 Dual Pipeline
We chose BAAI/bge-m3 for dense embeddings — a 1024-dimensional multilingual model that consistently ranks near the top of the MTEB leaderboard. For sparse vectors, we use Qdrant/bm25 to capture exact keyword matches that dense models miss (library names, specific acronyms, author names).
Diagnosing the embedding bottleneck with Recall@100: We initially experimented with smaller embedding models, but our retrieval results were poor. To pinpoint the bottleneck, we measured our Recall@100 — the ability of the model to place the correct chunk anywhere in the top 100 results. The score was abysmal. This was a critical insight: if relevant chunks aren't even in the top 100, it means the embedding model lacks the semantic capacity to understand and map the complex scientific text. Because no downstream re-ranker can re-sort chunks that were never retrieved in the first place, this low Recall@100 definitively proved we had to upgrade to a larger, higher-dimensional model like BGE-M3.
The key architectural insight is our dual-backend design:
Both backends produce the exact same 1024-dimensional BGE-M3 vectors. Cost was the driving decision here — renting a GPU for live query serving breaks our zero-budget constraint, so we engineered the retrieval backend to run entirely on free CPU tiers. The operational difference looks like this:
This means our production Docker image stays lean (no multi-GB PyTorch installation) while our ingestion pipeline maxes out Colab GPUs.
4. The 6-Colab Parallel Processing Strategy
Here's where things get scrappy. Processing 5,600 academic PDFs through Docling's layout analysis + BGE-M3 embedding is compute-intensive. A single Colab session (even with a free T4 GPU) would take days — and Colab kills sessions after ~12 hours.
Our solution: distribute the work across 6 free Google Colab accounts running in parallel.
Each Colab session ran our batch_gcs_to_drive.py script with an --embedding-batch-size 128 and --colab-gpu flag that overrides Docling's converter to use CUDA acceleration with 4 threads:
Crash resilience was critical. Colab sessions disconnect randomly. Our script checkpoints every 50 documents by copying the JSONL output to Google Drive. When a session dies, we restart with --start-paper pointing to the last checkpoint. The upload script uses UUID-v5 deduplication, so re-processing the same paper is harmless.
The JSONL format serves as our portable intermediate representation:
This decouples ingestion from storage completely. The 6 Colab accounts produce JSONL; a separate upload script ( import_remote_qdrant_parallel.py ) pushes to Qdrant in parallel with 8 threads and automatic retry logic (up to 10 attempts with exponential backoff per batch).
What the constraints taught us: We couldn't iterate quickly. Every full re-processing run took several hours of coordinating Colab sessions. This forced us to get the pipeline right upstream — investing heavily in crash-safety, checkpointing, and idempotent uploads rather than relying on "just re-run it."
5. Storage: Free-Tier Qdrant Cloud
We chose Qdrant because it was fundamentally the best technical fit for our architecture — cost was not the only factor. It's a highly performant, robust vector database written in Rust that we could seamlessly run locally for dev/testing (via Docker) and transition to the cloud for production.
Specifically, Qdrant offered four critical advantages:
Native multi-vector support — Each point stores both a dense vector (1024-dim cosine) and a named sparse vector ( bm25 ) simultaneously
Server-side fusion — Qdrant can execute prefetch queries across both vector spaces in a single round-trip
Rust-backed HNSW Performance — It utilizes highly optimized Hierarchical Navigable Small World (HNSW) graphs, delivering extremely high request-per-second (RPS) throughput and sub-25ms latency even on limited hardware
Free cloud tier — Generous enough for our 5,600-paper corpus
Every chunk gets a deterministic UUID-v5 derived from its SHA-256 content hash. This makes upserts idempotent — upload the same chunk twice, and Qdrant overwrites rather than duplicates.
We also build a payload index on metadata.year at collection creation, enabling Qdrant to apply year-based filters during the vector search (not after), which is critical for our query decomposition pipeline.
6. Retrieval: Intelligent Routing & Adaptive RAG
Instead of one-size-fits-all retrieval, we built a custom pipeline to route queries through different strategies based on their structure.
Before any database query fires, the router classifies the user's intent:
The router uses a pre-trained classifier on the dense query embedding, with hard regex overrides for metadata patterns. The override is deliberate — we don't trust ML classification for temporal constraints beca

[truncated]
