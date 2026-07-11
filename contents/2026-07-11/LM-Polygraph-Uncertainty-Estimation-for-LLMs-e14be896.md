---
source: "https://github.com/IINemo/lm-polygraph"
hn_url: "https://news.ycombinator.com/item?id=48876823"
title: "LM-Polygraph: Uncertainty Estimation for LLMs"
article_title: "GitHub - IINemo/lm-polygraph · GitHub"
author: "modinfo"
captured_at: "2026-07-11T23:43:47Z"
capture_tool: "hn-digest"
hn_id: 48876823
score: 1
comments: 0
posted_at: "2026-07-11T23:28:37Z"
tags:
  - hacker-news
  - translated
---

# LM-Polygraph: Uncertainty Estimation for LLMs

- HN: [48876823](https://news.ycombinator.com/item?id=48876823)
- Source: [github.com](https://github.com/IINemo/lm-polygraph)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T23:28:37Z

## Translation

タイトル: LM-Polygraph: LLM の不確実性推定
記事タイトル: GitHub - INemo/lm-polygraph · GitHub
説明: GitHub でアカウントを作成して、IInemo/lm-polygraph の開発に貢献します。

記事本文:
GitHub - IINemo/lm-polygraph · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
イイネモ
/
lm-ポリグラフ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーと

ファイル
1,325 コミット 1,325 コミット .github/ workflows .github/ workflows dataset_builders dataset_builders docs docs 例 例 ノートブック ノートブック スクリプト スクリプト src/ lm_polygraph src/ lm_polygraph test test .dockerignore .dockerignore .gitignore .gitignore .readthedocs.yaml .readthedocs.yaml CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE.md LICENSE.md README.md README.md pyproject.toml pyproject.toml要件.txt要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LM-Polygraph: LLM の不確実性推定
インストール |基本的な使い方 |概要 |ベンチマーク |デモアプリケーション |ドキュメント
LM-Polygraph は、テキスト生成タスクにおける LLM に一連の最先端の不確実性推定 (UE) 手法を提供します。不確実性が高いということは幻覚の存在を示している可能性があり、不確実性を推定するスコアを知ることは、LLM のアプリケーションをより安全にするのに役立ちます。
LM-Polygraph は、不確実性の推定および幻覚検出方法の一貫した評価に最も広く使用されているベンチマークの 1 つでもあります。何百もの研究者やテクノロジー企業に採用されています。
最新の安定バージョンはメイン ブランチで入手できます。仮想環境を使用することをお勧めします。
python -m venv env # これを仮想環境作成コマンドに置き換えます
ソース環境/bin/activate
pip install git+https://github.com/IINemo/lm-polygraph.git
タグも使用できます。
pip install git+https://github.com/IINemo/lm-polygraph.git@v0.5.0
PyPIから
最新のタグ付きバージョンは、PyPI 経由でも入手できます。
pip インストール lm-ポリグラフ
オプションの依存関係
一部の機能には、デフォルトではインストールされない追加のパッケージが必要です。
COMET メトリック (変換評価): unbabel-comet ピン numpy<2.0。vLLM などのパッケージと競合する可能性があります。インスタ

エクストラ経由で:
pip install lm-polygraph[comet]
numpy 2.x (vLLM など) が必要な場合は、余分なものをインストールせずにインストールし、comet を手動で追加します。
pip インストール lm-ポリグラフ
pip install unbabel-comet --no-deps
基本モデル (エンコーダー-デコーダーまたはデコーダーのみ) とトークナイザーを HuggingFace またはローカル ファイルから初期化し、それらを使用して評価用に WhiteboxModel を初期化します。
トランスフォーマーからのインポート AutoModelForCausalLM 、 AutoTokenizer
lm_polygraph から。ユーティリティ。モデルのインポート WhiteboxModel
model_path = "Qwen/Qwen2.5-0.5B-Instruct"
Base_model = AutoModelForCausalLM 。 from_pretrained (モデルパス、デバイスマップ = "cuda:0")
トークナイザー = AutoTokenizer 。 from_pretrained (モデルパス)
モデル = WhiteboxModel (ベースモデル、トークナイザー、モデルパス = モデルパス)
UE メソッドを指定します。
lm_polygraph から。エスティメーターのインポート *
ue_method = MeanTokenEntropy ()
予測とその不確実性スコアを取得します。
lm_polygraph から。 utils importestimate_uncertainty
input_text = "ジョージ・ブッシュとは誰ですか?"
ue = 推定不確実性 (モデル、ue_method、input_text = input_text)
プリント (うえ)
# UncertaintyOutput(uncertainty=-6.504108926902215, input_text='ジョージ・ブッシュとは誰ですか?',generation_text='アメリカ合衆国大統領', model_path='Qwen/Qwen2.5-0.5B-Instruct')
その他の例:basic_example.ipynb
コードに効率的に統合するための低レベルの例: low_level_example.ipynb も参照してください。
サービスとしてデプロイされた LLM での使用
LM-Polygraph は、OpenAI 互換の API サービスと連携できます。
lm_polygraph から BlackboxModel をインポート
lm_polygraph から。推定器は Perplexity 、 MaximumSequenceProbability をインポートします
モデル = BlackboxModel 。 from_openai (
openai_api_key = 'YOUR_API_KEY' ,
モデルパス = 'gpt-4o' 、
Supports_logprobs = True # デプロイメントを有効にする
)
ue_method = Perplexity () # または MeanTokenEntropy()、EigValL

アプラシアン()など
estimate_uncertainty (model , ue_method , input_text = '頭と尾はあるが胴体がないものは何ですか?' )
EigValLaplacian() などの UE メソッドは、ロジットを提供しない完全なブラックボックス LLM をサポートします。
Basic_example.ipynb : 個々のクエリをスコアリングする簡単な例
low_level_example.ipynb : 推論およびクレームレベルの UE への低レベルの統合
low_level_vllm_example.ipynb : 推論を高速化するために vLLM を使用する低レベルの例
Basic_visual_llm_example.ipynb : ビジュアル LLM の例
不確かさの推定方法
種類
カテゴリ
コンピューティング
記憶
トレーニング データが必要ですか?
レベル
最大系列確率
ホワイトボックス
情報ベース
低い
低い
いいえ
シーケンス/クレーム
困惑 (Fomicheva et al., 2020a)
ホワイトボックス
情報ベース
低い
低い
いいえ
シーケンス/クレーム
平均/最大トークン エントロピー (Fomicheva et al.、2020a)
ホワイトボックス
情報ベース
低い
低い
いいえ
シーケンス/クレーム
モンテカルロ数列エントロピー (Kuhn et al., 2023)
ホワイトボックス
情報ベース
高
低い
いいえ
シーケンス
点別相互情報量 (PMI) (高山、荒瀬、2019)
ホワイトボックス
情報ベース
中
低い
いいえ
シーケンス/クレーム
条件付き PMI (van der Poel 他、2022)
ホワイトボックス
情報ベース
中
中
いいえ
シーケンス
レンイの発散 (Darrin et al., 2023)
ホワイトボックス
情報ベース
低い
低い
いいえ
シーケンス
フィッシャー・ラオ距離 (Darrin et al., 2023)
ホワイトボックス
情報ベース
低い
低い
いいえ
シーケンス
注意スコア (Sriramanan et al.、2024)
ホワイトボックス
情報ベース
中
低い
いいえ
シーケンス/クレーム
文脈化された配列尤度 (CSL) (Lin et al., 2024)
ホワイトボックス
情報ベース
中
低い
いいえ
シーケンス
再発注意ベースの不確実性定量化 (RAUQ) (Vazhentsev et al., 2025)
ホワイトボックス
情報ベース
低い
低い
いいえ
シーケンス
フォーカス (Zhang et al., 2023)
ホワイトボックス
情報ベース
中
低い
いいえ
シーケンス/クレーム
BoostedProb (ディン語)

他、2025)
ホワイトボックス
情報ベース
低い
低い
いいえ
シーケンス/クレーム
意味論的エントロピー (Kuhn et al., 2023)
ホワイトボックス
多様性という意味
高
低い
いいえ
シーケンス
主張条件付き確率 (Fadeeva et al., 2024)
ホワイトボックス
多様性という意味
低い
低い
いいえ
シーケンス/クレーム
FrequencyScoring (Mohri et al., 2024)
ホワイトボックス
多様性という意味
高
低い
いいえ
主張する
TokenSAR (Duan 他、2023)
ホワイトボックス
多様性という意味
高
低い
いいえ
シーケンス/クレーム
SentenceSAR (Duan et al., 2023)
ホワイトボックス
多様性という意味
高
低い
いいえ
シーケンス
SAR (Duan 他、2023)
ホワイトボックス
多様性という意味
高
低い
いいえ
シーケンス
SemanticDensity (Qiu et al., 2024)
ホワイトボックス
多様性という意味
高
低い
いいえ
シーケンス
CoCoA (Vashurin et al.、2025)
ウィ
[切り捨てられた]
不確実性推定手法のパフォーマンスを評価するには、簡単な例を考えてみましょう。
CUDA_VISIBLE_DEVICES=0 ポリグラフ_eval \
--config-dir=./examples/configs/ \
--config-name=polygraph_eval_coqa.yaml \
モデル.パス=メタ-ラマ/ラマ-3.1-8B \
サブサンプル_eval_dataset=100
vLLM を生成に使用して不確実性推定手法のパフォーマンスを評価するには、次の例を検討してください。
CUDA_VISIBLE_DEVICES=0 ポリグラフ_eval \
--config-dir=./examples/configs/ \
--config-name=polygraph_eval_coqa.yaml \
モデル=vllm \
モデル.パス=メタ-ラマ/ラマ-3.1-8B \
estimators=default_estimators_vllm \
stat_calculators=default_calculators_vllm \
サブサンプル_eval_dataset=100
事前に構築された Docker コンテナをベンチマークに使用することもできます。例:
docker run --gpus '"device=0"' --rm \
-w /app \
inemo/lm_polygraph \
bash -c "ポリグラフ_eval \
--config-dir=./examples/configs/ \
--config-name=polygraph_eval_coqa.yaml \
モデル.パス=メタ-ラマ/ラマ-3.1-8B \
subsample_eval_dataset=100"
正しい形式のベンチマーク データセットは HF リポジトリにあります。データセット準備のためのスクリプト

dataset_builders ディレクトリにあるはずです。
実験用の要約テーブルを生成するには、visualization_tables.ipynb または result_tables.ipynb を使用します。
ベンチマークの詳細な説明はドキュメントにあります。
(廃止) デモ Web アプリケーション
@article{shelmanovvashurin2025,
著者 = {ヴァシュリン、ローマンとファデーワ、エカテリーナとヴァジェンツェフ、アルチョムとルヴァノワ、リュドミラとワシレフ、ダニイルとツヴィグン、アキムとペトラコフ、セルゲイとシン、ルイとサダラ、アブデルラフマンとグリシチェンコフ、キリルとパンチェンコ、アレクサンダーとボールドウィン、ティモシーとナコフ、プレスラフとパノフ、マキシムとシェルマノフ、アルチョム}、
title = {LM-Polygraph を使用した大規模言語モデルの不確実性定量化手法のベンチマーク},
ジャーナル = {計算言語学協会のトランザクション},
ボリューム = {13}、
ページ = {220-248}、
年 = {2025}、
月 = {03}、
issn = {2307-387X}、
ドイ = {10.1162/tacl_a_00737}、
URL = {https://doi.org/10.1162/tacl\_a\_00737},
eprint = {https://direct.mit.edu/tacl/article-pdf/doi/10.1162/tacl\_a\_00737/2511955/tacl\_a\_00737.pdf},
}
ACL-2025 チュートリアル:
@inproceedings{シェルマノフ-etal-2025-不確実性、
title = "大規模言語モデルの不確実性の定量化",
著者 = 「シェルマノフ、アルチョム、
パノフ、マキシム、
ヴァシュリン、ローマン、
ヴァジェンツェフ、アルチョム、
ファデーワ、エカテリーナ、
ボールドウィン、ティモシー」、
編集者＝『荒瀬と有紀と
ユルゲンス、デイヴィッド、
シア、フェイ」、
booktitle = "計算言語学協会第 63 回年次総会議事録 (第 5 巻: チュートリアル要約)",
月 = 7 月、
年 = "2025"、
住所 = "オーストリア、ウィーン"、
出版社 = 「計算言語学協会」、
URL = "https://aclanthology.org/2025.acl-tutorials.3/",
doi = "10.18653/v1/2025.acl-tutorials.3",
ページ = "3--4"、
ISBN = "979-8-89176-255-8"
}
EMNLP-2023 論文:
@進行中{ファ

deeva-etal-2023-lm、
title = "{LM}-ポリグラフ: 言語モデルの不確実性推定",
著者 = 「ファデーワ、エカテリーナ、そして
ヴァシュリン、ローマン、
ツヴィグン、アキム、
ヴァジェンツェフ、アルチョム、
ペトラコフ、セルゲイ、
フェディアニン、キリル、
ヴァシレフ、ダニール、
ゴンチャロワ、エリザベタ、
パンチェンコ、アレクサンダー、
パノフ、マキシム、
ボールドウィン、ティモシー、
シェルマノフ、アルチョム」、
編集者 = 「フェン、ヤンソン、そして
レフィーバー、エルス」、
booktitle = "自然言語処理における経験的手法に関する 2023 年会議議事録: システムのデモンストレーション",
月 = 12 月、
年 = "2023"、
住所 = "シンガポール"、
出版社 = 「計算言語学協会」、
URL = "https://aclanthology.org/2023.emnlp-demo.41",
doi = "10.18653/v1/2023.emnlp-demo.41",
ページ = "446--461",
abstract = "大規模言語モデル (LLM) の機能における最近の進歩により、さまざまな分野で無数の画期的なアプリケーションへの道が開かれました。しかし、これらのモデルはしばしば{``}幻覚を起こす{''}、つまり、発言の真実性を識別する明確な手段をユーザーに提供せずに事実を捏造するため、重大な課題が生じます。不確実性推定 (UE) 手法は、LLM をより安全で、より責任があり、より効果的に使用するための 1 つの方法です。ただし、

[切り捨てられた]

## Original Extract

Contribute to IINemo/lm-polygraph development by creating an account on GitHub.

GitHub - IINemo/lm-polygraph · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
IINemo
/
lm-polygraph
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,325 Commits 1,325 Commits .github/ workflows .github/ workflows dataset_builders dataset_builders docs docs examples examples notebooks notebooks scripts scripts src/ lm_polygraph src/ lm_polygraph test test .dockerignore .dockerignore .gitignore .gitignore .readthedocs.yaml .readthedocs.yaml CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE.md LICENSE.md README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
LM-Polygraph: Uncertainty Estimation for LLMs
Installation | Basic usage | Overview | Benchmark | Demo application | Documentation
LM-Polygraph provides a battery of state-of-the-art uncertainty estimation (UE) methods for LLMs in text generation tasks. High uncertainty can indicate the presence of hallucinations and knowing a score that estimates uncertainty can help to make applications of LLMs safer.
LM-Polygraph is also one of the most widely used benchmarks for the consistent evaluation of uncertainty estimation and hallucination detection methods. It is adopted by hundreds of researchers and technology companies.
The latest stable version is available in the main branch, it is recommended to use a virtual environment:
python -m venv env # Substitute this with your virtual environment creation command
source env/bin/activate
pip install git+https://github.com/IINemo/lm-polygraph.git
You can also use tags:
pip install git+https://github.com/IINemo/lm-polygraph.git@v0.5.0
From PyPI
The latest tagged version is also available via PyPI:
pip install lm-polygraph
Optional dependencies
Some features require additional packages that are not installed by default:
COMET metric (translation evaluation): unbabel-comet pins numpy<2.0 which may conflict with packages like vLLM. Install via extras:
pip install lm-polygraph[comet]
If you need numpy 2.x (e.g., for vLLM), install without the extra and add comet manually:
pip install lm-polygraph
pip install unbabel-comet --no-deps
Initialize the base model (encoder-decoder or decoder-only) and tokenizer from HuggingFace or a local file, and use them to initialize the WhiteboxModel for evaluation:
from transformers import AutoModelForCausalLM , AutoTokenizer
from lm_polygraph . utils . model import WhiteboxModel
model_path = "Qwen/Qwen2.5-0.5B-Instruct"
base_model = AutoModelForCausalLM . from_pretrained ( model_path , device_map = "cuda:0" )
tokenizer = AutoTokenizer . from_pretrained ( model_path )
model = WhiteboxModel ( base_model , tokenizer , model_path = model_path )
Specify the UE method:
from lm_polygraph . estimators import *
ue_method = MeanTokenEntropy ()
Get predictions and their uncertainty scores:
from lm_polygraph . utils import estimate_uncertainty
input_text = "Who is George Bush?"
ue = estimate_uncertainty ( model , ue_method , input_text = input_text )
print ( ue )
# UncertaintyOutput(uncertainty=-6.504108926902215, input_text='Who is George Bush?', generation_text=' President of the United States', model_path='Qwen/Qwen2.5-0.5B-Instruct')
More examples: basic_example.ipynb
See also a low-level example for efficient integration into your code: low_level_example.ipynb
Using with LLMs deployed as a service
LM-Polygraph can work with any OpenAI-compatible API services:
from lm_polygraph import BlackboxModel
from lm_polygraph . estimators import Perplexity , MaximumSequenceProbability
model = BlackboxModel . from_openai (
openai_api_key = 'YOUR_API_KEY' ,
model_path = 'gpt-4o' ,
supports_logprobs = True # Enable for deployments
)
ue_method = Perplexity () # or MeanTokenEntropy(), EigValLaplacian(), etc.
estimate_uncertainty ( model , ue_method , input_text = 'What has a head and a tail but no body?' )
UE methods such as EigValLaplacian() support fully blackbox LLMs that do not provide logits.
basic_example.ipynb : simple examples of scoring individual queries
low_level_example.ipynb : low-level integration into inference and claim-level UE
low_level_vllm_example.ipynb : low-level example using vLLM for faster inference
basic_visual_llm_example.ipynb : examples for visual LLMs
Uncertainty Estimation Method
Type
Category
Compute
Memory
Need Training Data?
Level
Maximum sequence probability
White-box
Information-based
Low
Low
No
sequence/claim
Perplexity (Fomicheva et al., 2020a)
White-box
Information-based
Low
Low
No
sequence/claim
Mean/max token entropy (Fomicheva et al., 2020a)
White-box
Information-based
Low
Low
No
sequence/claim
Monte Carlo sequence entropy (Kuhn et al., 2023)
White-box
Information-based
High
Low
No
sequence
Pointwise mutual information (PMI) (Takayama and Arase, 2019)
White-box
Information-based
Medium
Low
No
sequence/claim
Conditional PMI (van der Poel et al., 2022)
White-box
Information-based
Medium
Medium
No
sequence
Rényi divergence (Darrin et al., 2023)
White-box
Information-based
Low
Low
No
sequence
Fisher-Rao distance (Darrin et al., 2023)
White-box
Information-based
Low
Low
No
sequence
Attention Score (Sriramanan et al., 2024)
White-box
Information-based
Medium
Low
No
sequence/claim
Contextualized Sequence Likelihood (CSL) (Lin et al., 2024)
White-box
Information-based
Medium
Low
No
sequence
Recurrent Attention-based Uncertainty Quantification (RAUQ) (Vazhentsev et al., 2025)
White-box
Information-based
Low
Low
No
sequence
Focus (Zhang et al., 2023)
White-box
Information-based
Medium
Low
No
sequence/claim
BoostedProb (Dinh et al., 2025)
White-box
Information-based
Low
Low
No
sequence/claim
Semantic entropy (Kuhn et al., 2023)
White-box
Meaning diversity
High
Low
No
sequence
Claim-Conditioned Probability (Fadeeva et al., 2024)
White-box
Meaning diversity
Low
Low
No
sequence/claim
FrequencyScoring (Mohri et al., 2024)
White-box
Meaning diversity
High
Low
No
claim
TokenSAR (Duan et al., 2023)
White-box
Meaning diversity
High
Low
No
sequence/claim
SentenceSAR (Duan et al., 2023)
White-box
Meaning diversity
High
Low
No
sequence
SAR (Duan et al., 2023)
White-box
Meaning diversity
High
Low
No
sequence
SemanticDensity (Qiu et al., 2024)
White-box
Meaning diversity
High
Low
No
sequence
CoCoA (Vashurin et al., 2025)
Whi
[truncated]
To evaluate the performance of uncertainty estimation methods consider a quick example:
CUDA_VISIBLE_DEVICES=0 polygraph_eval \
--config-dir=./examples/configs/ \
--config-name=polygraph_eval_coqa.yaml \
model.path=meta-llama/Llama-3.1-8B \
subsample_eval_dataset=100
To evaluate the performance of uncertainty estimation methods using vLLM for generation, consider the following example:
CUDA_VISIBLE_DEVICES=0 polygraph_eval \
--config-dir=./examples/configs/ \
--config-name=polygraph_eval_coqa.yaml \
model=vllm \
model.path=meta-llama/Llama-3.1-8B \
estimators=default_estimators_vllm \
stat_calculators=default_calculators_vllm \
subsample_eval_dataset=100
You can also use a pre-built docker container for benchmarking, example:
docker run --gpus '"device=0"' --rm \
-w /app \
inemo/lm_polygraph \
bash -c "polygraph_eval \
--config-dir=./examples/configs/ \
--config-name=polygraph_eval_coqa.yaml \
model.path=meta-llama/Llama-3.1-8B \
subsample_eval_dataset=100"
The benchmark datasets in the correct format could be found in the HF repo . The scripts for dataset preparation could be found in the dataset_builders directory.
Use visualization_tables.ipynb or result_tables.ipynb to generate the summarizing tables for an experiment.
A detailed description of the benchmark is in the documentation .
(Obsolete) Demo web application
@article{shelmanovvashurin2025,
author = {Vashurin, Roman and Fadeeva, Ekaterina and Vazhentsev, Artem and Rvanova, Lyudmila and Vasilev, Daniil and Tsvigun, Akim and Petrakov, Sergey and Xing, Rui and Sadallah, Abdelrahman and Grishchenkov, Kirill and Panchenko, Alexander and Baldwin, Timothy and Nakov, Preslav and Panov, Maxim and Shelmanov, Artem},
title = {Benchmarking Uncertainty Quantification Methods for Large Language Models with LM-Polygraph},
journal = {Transactions of the Association for Computational Linguistics},
volume = {13},
pages = {220-248},
year = {2025},
month = {03},
issn = {2307-387X},
doi = {10.1162/tacl_a_00737},
url = {https://doi.org/10.1162/tacl\_a\_00737},
eprint = {https://direct.mit.edu/tacl/article-pdf/doi/10.1162/tacl\_a\_00737/2511955/tacl\_a\_00737.pdf},
}
ACL-2025 Tutorial:
@inproceedings{shelmanov-etal-2025-uncertainty,
title = "Uncertainty Quantification for Large Language Models",
author = "Shelmanov, Artem and
Panov, Maxim and
Vashurin, Roman and
Vazhentsev, Artem and
Fadeeva, Ekaterina and
Baldwin, Timothy",
editor = "Arase, Yuki and
Jurgens, David and
Xia, Fei",
booktitle = "Proceedings of the 63rd Annual Meeting of the Association for Computational Linguistics (Volume 5: Tutorial Abstracts)",
month = jul,
year = "2025",
address = "Vienna, Austria",
publisher = "Association for Computational Linguistics",
url = "https://aclanthology.org/2025.acl-tutorials.3/",
doi = "10.18653/v1/2025.acl-tutorials.3",
pages = "3--4",
ISBN = "979-8-89176-255-8"
}
EMNLP-2023 paper:
@inproceedings{fadeeva-etal-2023-lm,
title = "{LM}-Polygraph: Uncertainty Estimation for Language Models",
author = "Fadeeva, Ekaterina and
Vashurin, Roman and
Tsvigun, Akim and
Vazhentsev, Artem and
Petrakov, Sergey and
Fedyanin, Kirill and
Vasilev, Daniil and
Goncharova, Elizaveta and
Panchenko, Alexander and
Panov, Maxim and
Baldwin, Timothy and
Shelmanov, Artem",
editor = "Feng, Yansong and
Lefever, Els",
booktitle = "Proceedings of the 2023 Conference on Empirical Methods in Natural Language Processing: System Demonstrations",
month = dec,
year = "2023",
address = "Singapore",
publisher = "Association for Computational Linguistics",
url = "https://aclanthology.org/2023.emnlp-demo.41",
doi = "10.18653/v1/2023.emnlp-demo.41",
pages = "446--461",
abstract = "Recent advancements in the capabilities of large language models (LLMs) have paved the way for a myriad of groundbreaking applications in various fields. However, a significant challenge arises as these models often {``}hallucinate{''}, i.e., fabricate facts without providing users an apparent means to discern the veracity of their statements. Uncertainty estimation (UE) methods are one path to safer, more responsible, and more effective use of LLMs. However

[truncated]
