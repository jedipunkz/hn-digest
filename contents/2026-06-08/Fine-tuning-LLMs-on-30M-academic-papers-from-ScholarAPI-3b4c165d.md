---
source: "https://scholarapi.net/case_study/ai_training"
hn_url: "https://news.ycombinator.com/item?id=48446928"
title: "Fine-tuning LLMs on 30M academic papers from ScholarAPI"
article_title: "AI Training Case Study - ScholarAPI"
author: "mwojnars"
captured_at: "2026-06-08T16:28:03Z"
capture_tool: "hn-digest"
hn_id: 48446928
score: 1
comments: 0
posted_at: "2026-06-08T15:49:25Z"
tags:
  - hacker-news
  - translated
---

# Fine-tuning LLMs on 30M academic papers from ScholarAPI

- HN: [48446928](https://news.ycombinator.com/item?id=48446928)
- Source: [scholarapi.net](https://scholarapi.net/case_study/ai_training)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T15:49:25Z

## Translation

タイトル: ScholarAPI の 3,000 万学術論文の LLM の微調整
記事のタイトル: AI トレーニングのケーススタディ - ScholarAPI
説明: 世界中の学術文献をすぐに入手できます。全文検索と一括ダウンロードを備えた、学術 PDF やメタデータへの大規模なアクセスのための信頼できる Google Scholar の代替手段。壊れやすいスクレイパーを 1 つの効率的な API に置き換えます。

記事本文:
AI トレーニングのケーススタディ - ScholarAPI
ScholarAPI の機能 クイック スタート API ドキュメント 料金 FAQ ケース サインイン API キーの取得 ケース スタディ
高品質の学術データをモデルに活用します。
データセットを構築し、モデルを微調整し、包括的な知識ベースを作成します。
ChatGPT や Google Gemini などの汎用 LLM は強力ですが、ニッチなテーマについて質問されると幻覚を起こすことがよくあります。
「腫瘍随伴性天疱瘡」についてジェネリックモデルに聞く、
まれな自己免疫疾患であり、何もないところから治療プロトコルを発明する可能性があります。
まれな免疫疾患に特化した AI アシスタントを構築したいとします。
ウィキペディアだけでなく、医療症例報告の「ロングテール」でトレーニングする必要があります。
重要なステップは、オープン Web を超えて、学術出版物のような信頼できる知識源に移行することです。
ScholarAPI は、プログラム可能なアクセスを即座に提供することでこれを容易にします。
シンプルな REST インターフェイスを通じて数百万の論文を閲覧できます。
このケーススタディは医学に焦点を当てていますが、同じワークフローが材料科学、法的技術、化学工学、
深い科学的精度を必要とするあらゆる分野。
腫瘍随伴性天疱瘡とは・・・
日光への曝露によって引き起こされる一般的な皮膚疾患。通常、顔や腕に軽度の発疹として現れます。治療には局所保湿剤を使用し、直射日光を避けます。考えられる単語の連想に基づいた幻覚コンテンツ (!)。
免疫学向けに微調整された専門家向け AI モデル
腫瘍随伴性天疱瘡とは・・・
まれな自己免疫性皮膚粘膜水疱疾患で、リンパ増殖性腫瘍と関連することが多い。重度の痛みを伴う口内炎と多形性の皮膚発疹が特徴です。この症状は、プラキンファミリータンパク質、特にエンボプラキンとペリプラキンを標的とする自己抗体によって媒介されます。正確な定義は次のとおりです。

学術的なテキストを微調整します。
インポートリクエスト
パラメータ = {
' q ': [' "自己抗体" '、' "プラキンタンパク質" '、' "エンボプラキン" ']
}
論文 = []
True の場合:
resp = リクエスト.get(
"https://scholarapi.net/api/v1 /list ",
params=パラメータ、
headers={"X-API-Key": "YOUR_KEY"}
)
結果 = resp.json().get('結果')
結果がでない場合: 中断
結果のヒットの場合:
テキスト = リクエスト.get(
f"https://scholarapi.net/api/v1 /text /{hit['id']}",
headers={"X-API-Key": "YOUR_KEY"}
).テキスト
papers.append({**hit, 'full_text': text})
params['indexed_after'] = results[-1]['indexed_at'] 知識の収集
効果的なトレーニングを行うには、重要な量のドメイン固有のデータが必要です。
免疫学、バイオマーカー、病理学に関する何千もの論文を集約するために、
ScholarAPI の /list エンドポイントを使用し、次のような特定の用語を渡します。
q パラメータによる「自己抗体」、「プラキンタンパク質」、または「エンボプラキン」。
1 つ以上のクエリ フレーズを含む出版物はインデックス順に返されます。
最大 1,000 レコード (デフォルトでは 100) のバッチで。
次に、/text または /texts エンドポイントを使用して、各記事のプレーン テキストをダウンロードします。
このプロセスにより、この分野の真の複雑さを反映した生の学術散文の高密度で高品質なコーパスが作成されます。
アカウントを作成すると、API キーがダッシュボードで利用可能になります。
生のテキストだけでは命令のチューニングには十分ではありません。モデルに実行させたいタスクの構造化された例が必要です。
ヘルパー LLM を使用して、生の学術コーパスをスキャンし、さまざまなカテゴリにわたる何千ものトレーニング ペアを自動的に生成します。
要約: 複雑な要約を要約します。
Q&A: 調査結果に基づいて質問を作成します。
抽出: バイオマーカーと投与量を抽出します。
臨床推論: 診断ロジックをシミュレートします。
これにより、受動的な読み物が能動的なトレーニング ドリルに変わります。
から生成されたサンプル

学術書がモデルを教えてくれる
コンテキストとして指示 (タスクの説明) と入力 (記事のチャンク) が提供されたときに出力を生成します。
{
"instruction" : "眼症状には何が考えられますか
腫瘍随伴性天疱瘡のことでしょうか?",
"input" : "眼障害が頻繁に発生し、重篤です...
結膜炎は傷跡を引き起こす可能性があります...",
"output" : "重度の結膜炎が原因で
瘢痕化とシンブルファロンが特徴です...」
R 臨床推論
{
"instruction" : "臨床的意義を分析する
患者の IgE レベルの上昇」、
"input" : "患者は...検査結果を提示します
IgE > 2000 IU/mL を表示...",
"output" : "IgE の顕著な上昇は、
高IgE症候群または重度のアトピー
皮膚炎…」
E 抽出
{
"instruction" : "言及されたすべてのバイオマーカーを抽出します
入力テキスト内",
"input" : "この研究では、
50 人の患者における IL-6、TNF-α、および CRP。",
"出力" : ["IL-6"、"TNF-α"、"CRP"]
4 Python Node.js インポート リクエスト
def fetch_multimodal_data(paper_id):
# 完全な PDF バイナリを取得する
resp = リクエスト.get(
f"https://scholarapi.net/api/v1 /pdf /{paper_id}",
headers={"X-API-Key": "YOUR_KEY"}
)
open(f"{paper_id}.pdf", "wb") を f として使用:
f.write(それぞれのコンテンツ)
# 抽出パイプラインに渡す (例: PyMuPDF、非構造化)
# extract_images_and_captions(f.name) マルチモーダル エンリッチメント
医療診断は視覚的に行われます。証拠が組織学スライド、X 線、フローサイトメトリー チャートにある場合、テキストだけでは十分ではありません。
ScholarAPI の /pdf エンドポイントを使用すると、完全なバイナリ ドキュメントにアクセスできます。
完全な PDF をダウンロードすると、画像とそれに関連するキャプションを抽出してデータセットを充実させることができます
医療画像を「見る」ことができるだけでなく臨床メモを読むこともできるマルチモーダル モデルをトレーニングします。
生成された命令データセットと、場合によっては抽出された

マルチモーダル機能、
これで、教師あり微調整 (SFT) の準備が整いました。
準備された命令と出力のペアを、Llama や Mistral などの事前トレーニングされた基本モデルにフィードします。
このモデルは、数千のトレーニング ステップにわたって、その重みを免疫学の特殊な語彙と推論パターンに適応させます。
その結果、推測ではなく領域を真に理解するエキスパート モデルが誕生しました。
トランスフォーマーから AutoModelForCausalLM、TrainingArguments をインポート
trl インポート SFTTrainer から
# ドメイン固有のデータセットをロードします
dataset =load_dataset("json", data_files="免疫学_タスク.jsonl")
# 効率を高めるために低ランク適応 (LoRA) を構成する
peft_config = LoraConfig(
r=16、
lora_alpha=32、
target_modules=["q_proj", "v_proj"]
)
# トレーナーを初期化する
トレーナー = SFTTrainer(
モデル="ミストラライ/ミストラル-7B-v0.3",
train_dataset=データセット、
peft_config=peft_config,
args=トレーニング引数(
Output_dir="./immunogen-v1",
num_train_epochs=3、
per_device_train_batch_size=4
)
)
Trainer.train() ユーザー クエリ 「CAR-T の安全性」に関する 2026 年の最新の調査結果は何ですか?
SCHOLARAPI RETRIEVAL 3 つの情報源が見つかりました 📄 Journal of Immunology (2026): 安全性プロファイル... 📄 臨床試験最新情報: CAR-T 結果... ✨ AI 応答 「2026 年の研究によると、CAR-T の安全性は新しいサイトカイン管理プロトコルで改善されました...」
トレーニングは特定の日付で終了しますが、医学は日々進化しています。
モデルが最新のトライアルとプロトコルについて知る必要がある場合は、
検索拡張生成 (RAG) を使用してモデル推論を拡張します。
臨床医が質問すると、システムは重要なキーワードを特定し、リアルタイムで ScholarAPI にクエリを実行します。
最も関連性の高い最新の論文を見つけるために。
これらは生のテキストとしてモデルのコンテキスト ウィンドウに挿入され、最新の精度で応答できるようになります。
このハイブリッド アプローチ - 深い領域の知識

SFT の棚に RAG の最新の事実を加えて、賢明かつ最新の AI を作成します。
Domain Expert AI をトレーニングする準備はできていますか?
ScholarAPI 発見の未来を加速します。
科学のスピードを加速します。

## Original Extract

Global academic literature at your fingertips. Reliable Google Scholar alternative for large-scale access to academic PDFs and metadata, with full-text search and bulk download. Replace fragile scrapers with one efficient API.

AI Training Case Study - ScholarAPI
ScholarAPI Features Quick Start API Docs Pricing FAQs Cases Sign In Get API Key Case Study
Fuel your models with high-quality academic data.
Build datasets, fine-tune models, and create comprehensive knowledge bases.
General-purpose LLMs like ChatGPT or Google Gemini are powerful, but they often hallucinate when asked about niche subjects.
Ask a generic model about "paraneoplastic pemphigus" ,
a rare autoimmune disease, and it might invent a treatment protocol from thin air.
Suppose you want to build a specialized AI assistant for rare immunological disorders.
You need to train it on the "long tail" of medical case reports, not just Wikipedia.
The crucial step is moving beyond the open web to a trusted source of knowledge like academic publications.
ScholarAPI makes this easy by giving you instant, programmable access to
millions of papers through a simple REST interface.
While this case study focuses on medicine, the same workflow applies to materials science, legal tech, chemical engineering,
and any domain requiring deep scientific precision.
Paraneoplastic pemphigus is ...
a common skin condition caused by sun exposure. It typically presents as a mild rash on the face and arms. Treatment involves topical moisturizers and avoiding direct sunlight. Hallucinated content (!) based on probable word associations.
Specialist AI Model fine-tuned for Immunology
Paraneoplastic pemphigus is ...
a rare autoimmune mucocutaneous blistering disease often associated with lymphoproliferative neoplasms. It is characterized by severe painful stomatitis and polymorphous cutaneous eruptions. The condition is mediated by autoantibodies targeting plakin family proteins, specifically envoplakin and periplakin. Precise definition derived from fine-tuning on academic texts.
import requests
params = {
' q ': [' "autoantibodies" ', ' "plakin proteins" ', ' "envoplakin" ']
}
papers = []
while True:
resp = requests.get(
"https://scholarapi.net/api/v1 /list ",
params=params,
headers={"X-API-Key": "YOUR_KEY"}
)
results = resp.json().get('results')
if not results: break
for hit in results:
text = requests.get(
f"https://scholarapi.net/api/v1 /text /{hit['id']}",
headers={"X-API-Key": "YOUR_KEY"}
).text
papers.append({**hit, 'full_text': text})
params['indexed_after'] = results[-1]['indexed_at'] Gathering Knowledge
For effective training, you need a critical mass of domain-specific data.
To aggregate thousands of papers on immunology, biomarkers, and pathology,
use ScholarAPI's /list endpoint and pass specific terms like
"autoantibodies", "plakin proteins", or "envoplakin" via the q parameter.
Publications that contain one or more of the query phrases will be returned in indexing order,
in batches of up to 1,000 records (100 by default).
Then, use the /text or /texts endpoint to download the plain text of each article.
This process creates a dense, high-quality corpus of raw academic prose that reflects the true complexity of the field.
Your API key is available in the Dashboard after creating an account.
Raw text isn't enough for instruction tuning; you need structured examples of the tasks you want the model to perform.
Use a helper LLM to scan your raw academic corpus and automatically generate thousands of training pairs across diverse categories:
Summarization: Condense complex abstracts.
Q&A: Create questions based on findings.
Extraction: Pull out biomarkers and dosages.
Clinical Reasoning: Simulate diagnostic logic.
This transforms passive reading material into active training drills.
Samples generated from academic texts will teach the model
to generate the output when provided with an instruction (task description) and input (article chunk) as context.
{
"instruction" : "What can be the ocular manifestations
of paraneoplastic pemphigus?",
"input" : "Ocular involvement is frequent and severe...
Conjunctivitis can lead to scarring...",
"output" : "Severe conjunctivitis leading to
scarring and symblepharon is a hallmark..."
} R Clinical Reasoning
{
"instruction" : "Analyze the clinical significance
of the patient's elevated IgE levels",
"input" : "Patient presents with... Lab results
show IgE > 2000 IU/mL...",
"output" : "The markedly elevated IgE suggests
a hyper-IgE syndrome or severe atopic
dermatitis..."
} E Extraction
{
"instruction" : "Extract all biomarkers mentioned
in the input text",
"input" : "The study analyzed serum levels of
IL-6, TNF-alpha, and CRP in 50 patients.",
"output" : ["IL-6", "TNF-alpha", "CRP"]
} 4 Python Node.js import requests
def fetch_multimodal_data(paper_id):
# Get the full PDF binary
resp = requests.get(
f"https://scholarapi.net/api/v1 /pdf /{paper_id}",
headers={"X-API-Key": "YOUR_KEY"}
)
with open(f"{paper_id}.pdf", "wb") as f:
f.write(resp.content)
# Pass to extraction pipeline (e.g. PyMuPDF, Unstructured)
# extract_images_and_captions(f.name) Multi-Modal Enrichment
Medical diagnosis is visual. Text alone isn't enough when the evidence lies in histology slides, X-rays, and flow cytometry charts.
ScholarAPI's /pdf endpoint gives you access to the full binary documents.
By downloading the complete PDFs, you can extract images and their associated captions to enrich your dataset
and train multi-modal models that can "see" medical imagery as well as read the clinical notes.
With your generated instruction dataset and, possibly, the extracted multi-modal features,
you are ready for Supervised Fine-Tuning (SFT) .
Feed the prepared instruction / output pairs into a pretrained base model like Llama or Mistral.
Over thousands of training steps, the model adapts its weights to the specialized vocabulary and reasoning patterns of immunology.
The result is an expert model that no longer guesses but genuinely knows the domain.
from transformers import AutoModelForCausalLM, TrainingArguments
from trl import SFTTrainer
# Load your domain-specific dataset
dataset = load_dataset("json", data_files="immunology_tasks.jsonl")
# Configure Low-Rank Adaptation (LoRA) for efficiency
peft_config = LoraConfig(
r=16,
lora_alpha=32,
target_modules=["q_proj", "v_proj"]
)
# Initialize the trainer
trainer = SFTTrainer(
model="mistralai/Mistral-7B-v0.3",
train_dataset=dataset,
peft_config=peft_config,
args=TrainingArguments(
output_dir="./immunogen-v1",
num_train_epochs=3,
per_device_train_batch_size=4
)
)
trainer.train() User Query What are the latest 2026 findings on "CAR-T safety"?
SCHOLARAPI RETRIEVAL 3 sources found 📄 Journal of Immunology (2026): Safety profiles... 📄 Clinical Trials Update: CAR-T outcomes... ✨ AI Response "According to 2026 studies, CAR-T safety has improved with new cytokine management protocols..."
Training cuts off at a certain date, but medicine evolves daily.
If your model needs to know about the latest trials and protocols,
extend model inference with Retrieval-Augmented Generation (RAG) .
When a clinician asks a question, your system identifies the important keywords and queries ScholarAPI in real-time
to find the most relevant fresh papers.
These are injected as raw text into the model's context window, allowing it to answer with up-to-the-minute accuracy.
This hybrid approach—deep domain knowledge from SFT plus fresh facts from RAG—creates an AI that is both wise and current.
Ready to train your Domain Expert AI?
ScholarAPI Fueling the future of discovery.
Accelerating the speed of science.
