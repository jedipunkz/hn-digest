---
source: "https://huggingface.co/KRLabsOrg/verbatim-rag-modern-bert-v2"
hn_url: "https://news.ycombinator.com/item?id=48478775"
title: "Show HN: A 150M model that extracts verbatim evidence spans for RAG, no LLM call"
article_title: "KRLabsOrg/verbatim-rag-modern-bert-v2 · Hugging Face"
author: "justacoolname"
captured_at: "2026-06-10T19:02:53Z"
capture_tool: "hn-digest"
hn_id: 48478775
score: 5
comments: 0
posted_at: "2026-06-10T16:29:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A 150M model that extracts verbatim evidence spans for RAG, no LLM call

- HN: [48478775](https://news.ycombinator.com/item?id=48478775)
- Source: [huggingface.co](https://huggingface.co/KRLabsOrg/verbatim-rag-modern-bert-v2)
- Score: 5
- Comments: 0
- Posted: 2026-06-10T16:29:07Z

## Translation

タイトル: HN を表示: RAG の逐語的証拠スパンを抽出する 150M モデル、LLM 呼び出しなし
記事のタイトル: KRLabsOrg/verbatim-rag-modern-bert-v2 · 顔を抱きしめる
説明: 私たちは、オープンソースとオープン サイエンスを通じて人工知能を進歩させ、民主化する旅の途中にあります。

記事本文:
KRLabsOrg/verbatim-rag-modern-bert-v2 · 顔を抱きしめる
ハグ顔モデル
KRLabsOrg / verbatim-rag-modern-bert-v2 いいね 12 KR Labs をフォロー 47
トークン分類 Transformers Safetensors KRLabsOrg/verbatim-spans 英語 modernbert semantic-highlighting extractive-qaevidence-selection verbatim-rag custom_code arxiv: 2605.21102 ライセンス: apache-2.0 モデル カード ファイル ファイルとバージョン xet コミュニティ デプロイ バケットにコピー 新しい このモデルを使用する 使用手順KRLabsOrg/verbatim-rag-modern-bert-v2 にはライブラリ、推論プロバイダー、ノートブック、ローカル アプリが含まれます。これらのリンクに従って開始してください。
トランスフォーマー トランスフォーマーで KRLabsOrg/verbatim-rag-modern-bert-v2 を使用する方法:
# パイプラインを高レベルのヘルパーとして使用する
変圧器からのインポートパイプライン
Pipe = Pipeline("token-classification", model="KRLabsOrg/verbatim-rag-modern-bert-v2", trust_remote_code=True) # モデルを直接ロードする
トランスフォーマーから AutoTokenizer、AutoModelForTokenClassification をインポート
tokenizer = AutoTokenizer.from_pretrained("KRLabsOrg/verbatim-rag-modern-bert-v2", trust_remote_code=True)
model = AutoModelForTokenClassification.from_pretrained("KRLabsOrg/verbatim-rag-modern-bert-v2", trust_remote_code=True)
Verbatim-RAG エクストラクターの概要
モデルの詳細 トレーニングデータの構成
モデル名:verbatim-rag-modern-bert-v2
組織: KRLabsOrg
Github: https://github.com/KRLabsOrg/verbatim-rag
Verbatim-RAG Extractor は、クエリ条件付きのトークン分類子です。
質問に答える文章の範囲をそのまま強調表示します。それは
VerbatimRAG のエンコーダ コンパニオン
そしてその後継者
逐語的-rag-modern-bert-v1 。
上に構築
Alibaba-NLP/gte-reranker-modernbert-base 、
これは長い ModernBERT コンテキスト (最大 8192 トークン) と
スパン抽出に基づくクエリ条件付き再ランキング

イオンは微調整されています。
目標は、多くの LLM ベースの証拠を置き換えることができる軽量の抽出ツールです。
実稼働 RAG システムでの呼び出しの強調表示: ローカル、決定論的、低コスト
サービスを提供しており、スパンオーバーラップ品質では依然として競争力があります。 ACL-Verbatim ゴールドでは
ベンチマークでは、ACL に特化した兄弟モデルは強力な LLM と同等です
ワードレベルの F1 による抽出機能を備えていますが、この汎用マルチドメイン モデルはパブリック モデルを上回っています。
ACL gold、RAGBench、Squeez、QASPER 全体にわたる抽出ベースライン。
これを VerbatimRAG 内の抽出ステージとして使用したり、
取得/再ランキング後に独自の RAG パイプラインを作成し、取得したチャンクを
根拠のある証拠は、ユーザーに表示したり、管理者に渡したりする前に確認する必要があります。
発電機。
ほとんどの公開証拠抽出ツール (Provence、Zilliz Semantic-Highlight、
MultiSpanQA でトレーニングされたモデル）は、Wikipedia スタイルの散文 QA のみでトレーニングされます。
このモデルは以下でトレーニングされています
KRLabsOrg/verbatim-spans 、
財務表、法的契約、医学文献、製品を追加します
マニュアル、および公開されている抽出プログラムの中で独自のコーディング エージェント ツールの出力
(pytest の失敗、git diff ハンク、スタック トレース)。結果はシングルです
実際の RAG またはエージェントを形成するコンテンツ全体で使用可能な 150M パラメータのエンコーダ
パイプラインは記事の段落だけでなく、取得する傾向があります。
ACL アンソロジーに特化したバリアントについては、を参照してください。
KRLabsOrg/acl-verbatim-modernbert 。
アーキテクチャ: ModernBERT (gte-reranker-modernbert-base)、8192 トークン コンテキスト
タスク: トークン分類 — 文字スパンにマッピングされたバイナリ証拠ラベル
トレーニング データセット: KRLabsOrg/verbatim-spans (マルチドメイン)
(質問、コンテキスト) ペアは単一のシーケンスとしてエンコードされます。モデル
コンテキスト トークンに対するトークンごとのポジティブ クラスの確率を予測します。上
しきい値を設定すると、連続するポジティブなランが文字スパンにマージされます。
ポストプロ

削除する cessing ( min_span_chars 、 merge_gap_chars )
断片化アーティファクト。長いコンテキストは、次のスライディング ウィンドウで処理されます。
max_length トークンは doc_stride によってステップ化され、スパンは複数のトークン間でマージされます。
窓。
トランスフォーマーから AutoModel をインポート
モデル = AutoModel.from_pretrained(
"KRLabsOrg/verbatim-rag-modern-bert-v2" ,
trust_remote_code= True 、
）
結果 = モデル.プロセス(
question= 「ModernBERT とは何ですか?」 、
コンテキスト=(
「ModernBERT は NLP 用のロングコンテキスト エンコーダーです。」
"最大 8192 トークンのシーケンスをサポートします。"
「以前の BERT の亜種とは異なり、回転位置の埋め込みを使用します。」
）、
しきい値= 0.2 、
）
result[ "spans" ] のスパンの場合:
print ( f"[ {span[ 'score' ]: .2 f} ] {span[ 'text' ]} " )
VerbatimRAG 内で使用する
verbatim_rag.core から VerbatimRAG をインポート
verbatim_rag.index から VerbatimIndex をインポート
verbatim_rag.extractors から ModelSpanExtractor をインポート
verbatim_rag.vector_stores から LocalMilvusStore をインポート
verbatim_rag.embedding_providers から SpladeProvider をインポート
# v2 はデフォルトの ModelSpanExtractor モデルですが、これを明示的に渡すと
# 依存関係をクリアします。
エクストラクタ = ModelSpanExtractor(
model_path= "KRLabsOrg/verbatim-rag-modern-bert-v2" ,
しきい値= 0.2 、
min_span_chars= 30 、
merge_gap_chars= 20 、
device= None 、 # cuda、次に mps、そして cpu を自動検出します
）
sparse_provider = SpladeProvider(
model_name= "opensearch-project/opensearch-neural-sparse-encoding-doc-v2-distill" ,
device= "cuda" , # 利用可能な GPU がない場合は "cpu" を使用します
）
Vector_store = LocalMilvusStore(
db_path= "./index.db" ,
collection_name= "verbatim_rag" ,
Enable_dense= False 、
Enable_sparse= True 、
）
# インデックスにはすでにドキュメントが設定されていると仮定します。
インデックス = VerbatimIndex(
ベクトルストア=ベクトルストア、
sparse_provider=sparse_provider、
）
rag = VerbatimRAG(
インデックス=インデックス、
エクストラクター=抽出者、
k= 5 、

）
response = rag.query( "論文の主な発見?" )
印刷 (response.answer)
独自のリトリーバー/リランカーの直後にモデルを使用することもできます。
トランスフォーマーから AutoModel をインポート
エクストラクター = AutoModel.from_pretrained(
"KRLabsOrg/verbatim-rag-modern-bert-v2" ,
trust_remote_code= True 、
）
question = 「DINOv2 をビジュアル バックボーンとして使用することを裏付ける証拠は何ですか?」
コンテキスト = (
「特徴抽出のためにさまざまな視覚的バックボーンを調査します。」
「結果は、特徴抽出ツールとしての DINOv2 の有効性を示しています。」
手話通訳のため。
）
result = extractor.process(question=質問、context=コンテキスト)
result[ "spans" ] のスパンの場合:
印刷（
{
"開始" : スパン[ "開始" ]、
"終了" : スパン[ "終了" ],
"テキスト" : スパン[ "テキスト" ],
"スコア" : スパン[ "スコア" ],
}
）
.process() は次のものを受け入れます: question 、 context 、threshold (デフォルト 0.2 )、
max_length (デフォルト 8192 )、doc_stride (デフォルト 256 )、min_span_chars
(デフォルト 30 )、merge_gap_chars (デフォルト 20 )、return_sentence_metrics
(デフォルトは False )。記述式ベンチマークの場合 (ファイル パス、表のセル、
数値)、threshold=0.1、min_span_chars=10 はリコール調整された設定です
以下のパフォーマンスに記載されています。
戻り値の形式は {"spans": [{"start": int, "end": int, "text": str, "score": float}, ...]} で、次の場合に "sentences" が追加されます。
return_sentence_metrics=True 。スパンは入力への文字オフセットです
コンテキストとスライディング ウィンドウ間でマージされます。
が使用する共有スパン抽出ハーネスを使用して評価しました。
acl-verbatim 。現在の
メトリック プロトコルはスライス内のすべての行をスコアリングします。ゴールド スパンのない行は
陰性の例、および誤検知の抽出されたテキストにより精度が低下します。
以下の表は、この一般モデルを 2 つの公開抽出ベースラインと比較しています。
Zilliz セマンティック ハイライトとプロヴァンス。すべてのシステムは次の方法で評価されます。
すべての行のスコアが同じ

r.実行時間が依存するため、レイテンシーは意図的に省略されています
デバイス上、バッチ処理、およびサービス提供のセットアップ。
汎用モデルは、評価された 4 つのスライスすべてで最高の Word-F1 を達成します。
の一部ではない QASPER を含むパブリック抽出ベースラインに対して
トレーニングミックス。これが主な結果です。150M パラメータのローカル エンコーダは、
科学論文全体にわたる強力な汎用証拠ハイライトとして機能します。
RAGBench QA ドメイン、コーディング エージェント ツールの出力、および QASPER 科学 QA。の
RAGBench と Squeez での利点が最も強く、マルチドメインと一致します。
トレーニングミックス。 ACLゴールドではジェネリックモデルも一般モデルより強い
ベースラインの枝刈り/強調表示。ただし、ACL に特化したモデルが最適であることに変わりはありません
ホームドメイン上で。プロヴァンスは、次のようなリコール指向の指標において強いことがよくあります。
AnyOverlap と同様ですが、大幅に過剰予測する傾向があります。ジリズは一般的に
より保守的で再現率が低くなります。
評価コマンドとスライス構築については、次の文書に記載されています。
docs/GENERIC_EVAL.md 。
@misc{レクスキー:2026,
title={ACL-Verbatim: 研究のための幻覚のない質問応答},
著者={ガーボール・レクスキー、シルベスター・トート、ナディア・ヴェルダ、イシュトヴァーン・ボロス、アダム・コヴァチ}、
年={2026}、
eprint={2605.21102}、
archivePrefix={arXiv}、
プライマリクラス={cs.CL}、
URL={https://arxiv.org/abs/2605.21102}、
}
先月のダウンロード数 939
Safetensor モデル サイズ 0.1B params Tensor タイプ F32 · ファイル情報
推論プロバイダーの新しいトークン分類 このモデルは、どの推論プロバイダーによってもデプロイされていません。 🙋 プロバイダーのサポートを求める KRLabsOrg/verbatim-rag-modern-bert-v2 のモデル ツリー
Finetuned Alibaba-NLP/gte-reranker-modernbert-base Finetuned ( 17 ) このモデル KRLabsOrg/verbatim-rag-modern-bert-v2 のトレーニングに使用されるデータセット
KRLabsOrg/verbatim-spans ビューア • 8 日前に更新 • 391k • 173 • 3 スパ

KRLabsOrg/verbatim-rag-modern-bert-v2 を使用する 1
KRLabsOrg/verbatim-rag-modern-bert-v2 を含むコレクション
9 KRLabsOrg/verbatim-rag-modern-bert-v2 の論文

## Original Extract

We’re on a journey to advance and democratize artificial intelligence through open source and open science.

KRLabsOrg/verbatim-rag-modern-bert-v2 · Hugging Face
Hugging Face Models
KRLabsOrg / verbatim-rag-modern-bert-v2 like 12 Follow KR Labs 47
Token Classification Transformers Safetensors KRLabsOrg/verbatim-spans English modernbert semantic-highlighting extractive-qa evidence-selection verbatim-rag custom_code arxiv: 2605.21102 License: apache-2.0 Model card Files Files and versions xet Community Deploy Copy to bucket new Use this model Instructions to use KRLabsOrg/verbatim-rag-modern-bert-v2 with libraries, inference providers, notebooks, and local apps. Follow these links to get started.
Transformers How to use KRLabsOrg/verbatim-rag-modern-bert-v2 with Transformers:
# Use a pipeline as a high-level helper
from transformers import pipeline
pipe = pipeline("token-classification", model="KRLabsOrg/verbatim-rag-modern-bert-v2", trust_remote_code=True) # Load model directly
from transformers import AutoTokenizer, AutoModelForTokenClassification
tokenizer = AutoTokenizer.from_pretrained("KRLabsOrg/verbatim-rag-modern-bert-v2", trust_remote_code=True)
model = AutoModelForTokenClassification.from_pretrained("KRLabsOrg/verbatim-rag-modern-bert-v2", trust_remote_code=True)
Verbatim-RAG Extractor Overview
Model Details Training data composition
Model Name: verbatim-rag-modern-bert-v2
Organization: KRLabsOrg
Github: https://github.com/KRLabsOrg/verbatim-rag
The Verbatim-RAG Extractor is a query-conditioned token classifier that
highlights the verbatim spans of a passage that answer a question. It is the
encoder companion to VerbatimRAG
and the successor to
verbatim-rag-modern-bert-v1 .
Built on
Alibaba-NLP/gte-reranker-modernbert-base ,
which provides the long ModernBERT context (up to 8192 tokens) and a
query-conditioned reranking prior on top of which span extraction is fine-tuned.
The goal is a lightweight extractor that can replace many LLM-based evidence
highlighting calls in production RAG systems: local, deterministic, cheap to
serve, and still competitive on span-overlap quality. In our ACL-Verbatim gold
benchmark, the ACL-specialized sibling model is on par with strong LLM
extractors by word-level F1, while this generic multi-domain model beats public
extractive baselines across ACL gold, RAGBench, Squeez, and QASPER.
You can use it as the extraction stage inside VerbatimRAG, or drop it into your
own RAG pipeline after retrieval/reranking to turn retrieved chunks into
grounded evidence spans before displaying them to users or passing them to a
generator.
Most public evidence extractors (Provence, Zilliz Semantic-Highlight,
MultiSpanQA-trained models) are trained on Wikipedia-style prose QA only.
This model is trained on
KRLabsOrg/verbatim-spans ,
which adds financial tables, legal contracts, medical literature, product
manuals, and — uniquely among public extractors — coding-agent tool output
( pytest failures, git diff hunks, stack traces). The result is a single
150M-parameter encoder usable across the content shapes a real RAG or agent
pipeline tends to retrieve, not just article paragraphs.
For an ACL-Anthology-specialized variant, see
KRLabsOrg/acl-verbatim-modernbert .
Architecture: ModernBERT (gte-reranker-modernbert-base) with 8192-token context
Task: Token classification — binary evidence labels mapped to character spans
Training Dataset: KRLabsOrg/verbatim-spans (multi-domain)
A (question, context) pair is encoded as a single sequence; the model
predicts a per-token positive-class probability over the context tokens. Above
a threshold, contiguous positive runs are merged into character spans, with
post-processing ( min_span_chars , merge_gap_chars ) that removes
fragmentation artifacts. Long contexts are handled with sliding windows of
max_length tokens stepped by doc_stride , and spans are merged across
windows.
from transformers import AutoModel
model = AutoModel.from_pretrained(
"KRLabsOrg/verbatim-rag-modern-bert-v2" ,
trust_remote_code= True ,
)
result = model.process(
question= "What is ModernBERT?" ,
context=(
"ModernBERT is a long-context encoder for NLP. "
"It supports sequences up to 8192 tokens. "
"Unlike earlier BERT variants, it uses rotary position embeddings."
),
threshold= 0.2 ,
)
for span in result[ "spans" ]:
print ( f"[ {span[ 'score' ]: .2 f} ] {span[ 'text' ]} " )
Use inside VerbatimRAG
from verbatim_rag.core import VerbatimRAG
from verbatim_rag.index import VerbatimIndex
from verbatim_rag.extractors import ModelSpanExtractor
from verbatim_rag.vector_stores import LocalMilvusStore
from verbatim_rag.embedding_providers import SpladeProvider
# v2 is the default ModelSpanExtractor model, but passing it explicitly makes
# the dependency clear.
extractor = ModelSpanExtractor(
model_path= "KRLabsOrg/verbatim-rag-modern-bert-v2" ,
threshold= 0.2 ,
min_span_chars= 30 ,
merge_gap_chars= 20 ,
device= None , # auto-detects cuda, then mps, then cpu
)
sparse_provider = SpladeProvider(
model_name= "opensearch-project/opensearch-neural-sparse-encoding-doc-v2-distill" ,
device= "cuda" , # use "cpu" if no GPU is available
)
vector_store = LocalMilvusStore(
db_path= "./index.db" ,
collection_name= "verbatim_rag" ,
enable_dense= False ,
enable_sparse= True ,
)
# Assumes the index has already been populated with your documents.
index = VerbatimIndex(
vector_store=vector_store,
sparse_provider=sparse_provider,
)
rag = VerbatimRAG(
index=index,
extractor=extractor,
k= 5 ,
)
response = rag.query( "Main findings of the paper?" )
print (response.answer)
You can also use the model directly after your own retriever/reranker:
from transformers import AutoModel
extractor = AutoModel.from_pretrained(
"KRLabsOrg/verbatim-rag-modern-bert-v2" ,
trust_remote_code= True ,
)
question = "What evidence supports using DINOv2 as the visual backbone?"
context = (
"We investigate different visual backbones for feature extraction. "
"The results demonstrate DINOv2's effectiveness as a feature extractor "
"for sign language translation."
)
result = extractor.process(question=question, context=context)
for span in result[ "spans" ]:
print (
{
"start" : span[ "start" ],
"end" : span[ "end" ],
"text" : span[ "text" ],
"score" : span[ "score" ],
}
)
.process() accepts: question , context , threshold (default 0.2 ),
max_length (default 8192 ), doc_stride (default 256 ), min_span_chars
(default 30 ), merge_gap_chars (default 20 ), return_sentence_metrics
(default False ). For short-answer benchmarks (file paths, table cells,
numbers), threshold=0.1 and min_span_chars=10 is the recall-tuned config
documented in Performance below.
The return shape is {"spans": [{"start": int, "end": int, "text": str, "score": float}, ...]} , with "sentences" added when
return_sentence_metrics=True . Spans are character offsets into the input
context and are merged across sliding windows.
Evaluated with the shared span-extraction harness used by
acl-verbatim . The current
metric protocol scores every row in a slice: rows without gold spans are
negative examples, and false-positive extracted text lowers precision.
The table below compares this generic model to two public extractive baselines:
Zilliz Semantic Highlight and Provence. All systems are evaluated with the
same all-row scorer. Latency is intentionally omitted because runtime depends
on device, batching, and serving setup.
The generic model achieves the best Word-F1 on all four evaluated slices
against the public extractor baselines, including QASPER, which is not part of
the training mix. This is the main result: a 150M-parameter local encoder can
act as a strong general-purpose evidence highlighter across scientific papers,
RAGBench QA domains, coding-agent tool output, and QASPER scientific QA. The
advantage is strongest on RAGBench and Squeez, matching the multi-domain
training mix. On ACL gold, the generic model is also stronger than the public
pruning/highlighting baselines, though the ACL-specialized model remains best
on its home domain. Provence is often stronger on recall-oriented metrics such
as AnyOverlap, but tends to over-predict substantially more; Zilliz is generally
more conservative and lower recall.
Evaluation commands and slice construction are documented in
docs/GENERIC_EVAL.md .
@misc{Recski:2026,
title={ACL-Verbatim: hallucination-free question answering for research},
author={Gábor Recski and Szilveszter Tóth and Nadia Verdha and István Boros and Ádám Kovács},
year={2026},
eprint={2605.21102},
archivePrefix={arXiv},
primaryClass={cs.CL},
url={https://arxiv.org/abs/2605.21102},
}
Downloads last month 939
Safetensors Model size 0.1B params Tensor type F32 · Files info
Inference Providers NEW Token Classification This model isn't deployed by any Inference Provider. 🙋 Ask for provider support Model tree for KRLabsOrg/verbatim-rag-modern-bert-v2
Finetuned Alibaba-NLP/gte-reranker-modernbert-base Finetuned ( 17 ) this model Dataset used to train KRLabsOrg/verbatim-rag-modern-bert-v2
KRLabsOrg/verbatim-spans Viewer • Updated 8 days ago • 391k • 173 • 3 Space using KRLabsOrg/verbatim-rag-modern-bert-v2 1
Collection including KRLabsOrg/verbatim-rag-modern-bert-v2
9 Paper for KRLabsOrg/verbatim-rag-modern-bert-v2
