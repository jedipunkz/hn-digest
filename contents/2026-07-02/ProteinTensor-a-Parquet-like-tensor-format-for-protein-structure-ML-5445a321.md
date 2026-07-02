---
source: "https://github.com/mooreneural/HelixDB"
hn_url: "https://news.ycombinator.com/item?id=48766566"
title: "ProteinTensor – a Parquet-like tensor format for protein-structure ML"
article_title: "GitHub - mooreneural/HelixDB: AI-native tensor format (.ptt) for protein structure ML. Convert mmCIF/PDB structures or raw sequences once, then load backbone, MSA, embeddings, and bond graphs in milliseconds. Feeds AlphaFold/Boltz pipelines. Zarr-backed, lossless, PyTorch/JAX-ready. · GitHub"
author: "ramen0w0"
captured_at: "2026-07-02T20:55:13Z"
capture_tool: "hn-digest"
hn_id: 48766566
score: 2
comments: 0
posted_at: "2026-07-02T19:58:27Z"
tags:
  - hacker-news
  - translated
---

# ProteinTensor – a Parquet-like tensor format for protein-structure ML

- HN: [48766566](https://news.ycombinator.com/item?id=48766566)
- Source: [github.com](https://github.com/mooreneural/HelixDB)
- Score: 2
- Comments: 0
- Posted: 2026-07-02T19:58:27Z

## Translation

タイトル: ProteinTensor – タンパク質構造 ML 用の Parquet のようなテンソル形式
記事のタイトル: GitHub - mooreneural/HelixDB: タンパク質構造 ML 用の AI ネイティブ テンソル形式 (.ptt)。 mmCIF/PDB 構造または生シーケンスを一度変換し、バックボーン、MSA、エンベディング、および結合グラフをミリ秒単位でロードします。 AlphaFold/Boltz パイプラインにフィードします。 Zarr サポート、ロスレス、PyTorch/JAX 対応。 · GitHub
説明: タンパク質構造 ML の AI ネイティブ テンソル形式 (.ptt)。 mmCIF/PDB 構造または生シーケンスを一度変換し、バックボーン、MSA、エンベディング、および結合グラフをミリ秒単位でロードします。 AlphaFold/Boltz パイプラインにフィードします。 Zarr サポート、ロスレス、PyTorch/JAX 対応。 - モアニューラル/HelixDB

記事本文:
GitHub - mooreneural/HelixDB: タンパク質構造 ML 用の AI ネイティブ テンソル形式 (.ptt)。 mmCIF/PDB 構造または生シーケンスを一度変換し、バックボーン、MSA、エンベディング、および結合グラフをミリ秒単位でロードします。 AlphaFold/Boltz パイプラインにフィードします。 Zarr サポート、ロスレス、PyTorch/JAX 対応。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロード

セッションを更新します。
アラートを閉じる
{{ メッセージ }}
ムーアニューラル
/
HelixDB
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
32 コミット 32 コミット .github/ workflows .github/ workflows ベンチマーク ベンチマーク プロテインテンソル プロテインテンソル テスト テスト .gitignore .gitignore CITATION.cff CITATION.cff ライセンス ライセンス README.md README.md ボルトz_benchmark.py ボルトツ_ベンチマーク.py 紙.bib 紙.bib 紙.md Paper.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ProteinTensor は、AI ネイティブの生体分子ストレージ形式であり、
現代の構造生物学の機械学習パイプラインにおける前処理のボトルネック。
研究者が AlphaFold、Boltz、RoseTTAFold、OpenFold などをトレーニングするたびに
構造予測モデルでは、単一の GPU 操作が実行される前に同じ作業が行われます。
データセット内の各タンパク質について:
mmCIF / PDB ファイルを解析 (構造ごとに 30 ～ 350 ミリ秒)
シーケンストークンを抽出する
原子座標配列を構築する
バックボーン ジオメトリを構築する
共有結合グラフを計算する
MSA のロードまたは再生成 (JackHMMER で 2 ～ 480 分)
ESM2 / ESM3 推論の実行 (GPU 上のタンパク質あたりの秒数)
距離行列を計算する
...
-> 最後に:model.forward(features)
100,000 構造のトレーニングを実行する場合、この前処理には数千の CPU 時間のコストがかかります
エポックごと - ほとんどの場合、毎回同じ結果が生成されます。 mmCIF ファイルは
変わりました。順序は変わっていません。物理学は変わっていません。それでも、走るたびに
すべてを最初から再計算します。
ProteinTensor は、PDB エントリを .ptt ファイルに一度変換することでこの問題を解決します。
Zarr ベースの、LZ4 圧縮された、モデルのすべてのテンソルを保持するメモリ マップ可能なストア
ニーズ - そしてそれらの数十を読み込みます

または、解析なしでトレーニング時に直接実行されます。
1 回: mmCIF -> ProteinTensor (.ptt)
常に: .ptt -> model.forward()
これは誰のためのものですか
AlphaFold 3、Boltz、または Chai-1 を実行している構造生物学の研究者
新しい実験が行われるたびに、MSA の生成を数時間待機します。
大規模な構造を反復処理する製薬会社やバイオテクノロジー企業の ML エンジニア
I/O スループットがトレーニングとなるデータベース (PDB、AlphaFold Database、ESMAtlas)
壁時計の時代で測定されたボトルネック。
GPU 予算が限られており、コンピューティング サイクルを無駄にする余裕がない学術研究機関
GPU 時間をモデルのトレーニングに費やす必要がある場合に、テキスト ファイルを再解析するときに使用します。
構造生物学パイプラインを構築しているソフトウェア エンジニアは、単一の
明確に定義された中間形式で、PyTorch、JAX、および NumPy を使用せずに動作します。
すべてのモデルのカスタム ローダーを作成します。
ProteinTensor は構造生物学にとって、Parquet は分析にとって、セーフテンソルとは何なのか
重みをモデル化するものであり、交換をモデル化するものは ONNX です - 共通、オープン、高
定期的な計算税を 1 回限りの費用に変えるパフォーマンス形式。
ベンチマーク: 従来のパイプラインと ProteinTensor
すべてのタイミングは、NVIDIA RTX 5080、CUDA 12.8、Python 3.11 での 30 ラウンドの中央値です。
タンパク質は、76 残基のドメインから 3,525 残基の CRISPR 酵素までの全範囲に及びます。
再現するには、pythonboltz_benchmark.py を実行します。
構造
方法
解像度
MSA 配列
mmCIF 解析
ポイント: いっぱい
ptt: バックボーン
ptt: 絆
ptt:MSA
ポイント: 距離MX
1UBQ - ユビキチン
X線
76
512
7.2ミリ秒
2.8ミリ秒
1.2ミリ秒
0.7ミリ秒
1.6ミリ秒
0.8ミリ秒
6LU7 - SARS-CoV-2 Mpro
X線
312
1,024
29.6ミリ秒
2.9ミリ秒
1.2ミリ秒
0.7ミリ秒
5.1ミリ秒
2.0ミリ秒
4HHB - ヘモグロビン
X線
574
2,048
55.3ミリ秒
2.9ミリ秒
1.2ミリ秒
0.7ミリ秒
11.3ミリ秒
3.5ミリ秒
6M0J - ACE2 + RBD
クライオEM
791
2,048
74.7ミリ秒
2.9ミリ秒
1.2ミリ秒
0.7ミリ秒
14.7ミリ秒
6.4ミリ秒
6VXX - スパイクトリマー
クライオEM
2,916
8,192
283.4m

s
3.3ミリ秒
1.3ミリ秒
0.9ミリ秒
208.3ミリ秒
71.1ミリ秒
6OHW - Cas12a
クライオEM
3,525
8,192
352.4ミリ秒
3.3ミリ秒
1.2ミリ秒
1.0ミリ秒
240.7ミリ秒
104.5ミリ秒
列の定義
ptt: フル - read() - すべての原子、バックボーン、結合、メタデータ
ptt: バックボーン - read_backbone() - N/CA/C/O 座標 + シーケンスのみ
ptt: 結合 - read_bonds() - 共有結合グラフのみ
ptt: MSA - read_msa() - MSA トークン + プロファイル (.ptt キャッシュからロード)
ptt: dist mx - read_pair_feature(" distance_matrix ") - Ca-Ca 距離行列
構造
解像度
いっぱい
バックボーン
債券
MSA
ディストMX
1UBQ - ユビキチン
76
3倍
6倍
11倍
4倍
9倍
6LU7 - SARS-CoV-2 Mpro
312
10倍
24倍
43倍
6倍
15倍
4HHB - ヘモグロビン
574
19倍
45倍
78倍
5倍
16倍
6M0J - ACE2 + RBD
791
26倍
61倍
102倍
5倍
12倍
6VXX - スパイクトリマー
2,916
87倍
223x
308倍
1x*
4倍
6OHW - Cas12a
3,525
108倍
284x
370倍
1x*
3倍
*MSA の高速化は、mmCIF 解析と比較して 1x として示されています。これは、両方が同じ時間範囲にあるためです。
大きなタンパク質 - 実際の MSA の比較は、JackHMMER 世代との比較です (以下を参照)。
機能アセンブリ: model.forward() 用にすべてのテンソルを準備する時間
従来の = mmCIF 解析 + A3M ファイルから MSA を読み取ります。 ProteinTensor = 単一の .ptt 読み取り
すべての特徴が事前にキャッシュされています (シーケンス、バックボーン、結合、MSA、距離行列、
ESM2 埋め込み)。
6 つの構造すべての平均速度向上: フル機能アセンブリで 34 倍。
KRAS 腫瘍学にまたがる 6 つの高価値の薬剤ターゲットにわたって同じ方法論、
HIV 抗ウイルス薬、PD-L1 免疫療法、p53、心血管系 (PCSK9)、および完全な
IgG1抗体。数値は上記の構造生物学ベンチマークと一致しています。
ProteinDataset + ProteinDataset.collate() を使用して測定され、構造を
パディングされたバッチは、model.forward() の準備ができています。単一プロセス、プリフェッチ ワーカーなし。
スケール予測: 100,000 構造、1 トレーニング エポック
これらは、測定された構造ごとのタイミングから外挿された予測です。
上 - 終わりではありません

100kスケールでのOエンド測定。
MSA 生成は、32 コア サーバー (PDB90 データベース、標準) でタンパク質あたり 2.4 分を想定しています。
AlphaFold 設定)。 ProteinTensor は MSA を 1 回生成し、.ptt キャッシュからロードします。
その後の実行ごとに。 4,000 時間という数字は、AlphaFold2 と Voltz の実質コストです。
ユーザーは、トレーニング データセットをゼロから構築するために料金を支払います。
測定値と予測値 - これをお読みください。上記の 1,794x は MSA 世代です
(JackHMMER を使用してアライメントを一度構築) であり、文献ベースです。
投影、ここでベンチマークされたものではありません。ハードウェア上で測定されるのは、
エポックごとに繰り返される MSA ロード - キャッシュされた MSA を .ptt から読み取る場合と、
A3M テキストを各エポックで再解析します (ベクトル化された A3M パーサー ベースラインに対して):
3.4x ～ 5.9x 、MSA の深さに応じて増加します。参照
ベンチマーク/MSA_RESULTS.md 。これらは異なります
量; 1,794x を負荷の高速化の測定値として読み取らないでください。
フル機能の .ptt (8,192 シーケンス MSA + 距離行列 + ESM2-650M 埋め込み
float16) は、6 つのベンチマーク構造全体でソース mmCIF よりも平均して 23 倍大きくなります。
このトレードオフは意図的です。ディスク容量を一度支払うことで、GPU 時間と CPU 時間の支払いを回避できます。
トレーニングの実行ごとに。キャッシュされたフィーチャを持たない構造のみの .ptt は、
ソース mmCIF。
pip install -e " .[dev] " # コア + 開発ツール
pip install -e " .[cloud] " # リモート読み取り用に fsspec、s3fs、gcsfs を追加します
pip install -e " .[dev,cloud] " # すべて
Python >= 3.9、 gemmi 、 zarr 、 numpy 、 click 、 rich が必要です。
プロテインテンソル変換 1abc.cif 1abc.ptt
プロテインテンソル情報 1abc.ptt
シーケンスを変換します (構造は必要ありません)
AlphaFold や Voltz などのシーケンス駆動型の予測変数の場合、主な入力は
構造ではなく順序です。 ProteinTensor はシーケンスのみの .ptt (いいえ
座標) を生の文字列または FASTA ファイルから直接取得します。
プロテインテンソル変換シーケンス MQIFVKTLTG

KTITLEVEPSDTIENVKAKIQDKEGIPPDQQRLIFAGKQLEDG ubq.ptt
タンパク質テンソル変換-seq complex.fasta complex.ptt # マルチレコード FASTA -> マルチチェーン
プロテインテンソルを pt としてインポート
データ = pt 。 from_sequence ( "MQIFVKTLTGK..." 、 pdb_id = "UBQ" 、chain_id = "A" )
データ。 has_struct # False - シーケンスのみのエントリ
データ。 sequence_tokens # (N_res,) int32
点。 write ( data , "ubq.ptt" )
# FASTA: 単一レコード -> 1 チェーン。複数のレコード -> マルチチェーン複合体
データ = pt 。 from_fasta ( "complex.fasta" )
ディレクトリを一括変換する
構造のディレクトリ全体を並行して変換し、進行状況をレポートします。
解析に失敗したファイルはスキップされ、概要にリストされます。すでに変換済み
出力はデフォルトではスキップされます。
プロテインテンソル変換ディレクトリ ./pdb_files/ ./ptt_files/ # 自動ワーカー数
プロテインテンソル変換ディレクトリ ./pdb_files/ ./ptt_files/ --workers 16 --recursive
プロテインテンソル変換ディレクトリ ./pdb_files/ ./ptt_files/ --overwrite # 既存のファイルを再構築
mmCIF に対するベンチマーク
プロテインテンソル ベンチマーク 1abc.cif --ラウンド 20
MSA をキャッシュする (JackHMMER / ColabFold の実行後)
プロテインテンソルを pt としてインポート
msa = pt 。 from_a3m ( "1abc_uniref90.a3m" ,
ツール = "ジャックハンマー" 、ツールバージョン = "3.3.2" 、
データベース = "uniref90" 、データベース日付 = "2024-01" )
点。 add_msa ( "1abc.ptt" , msa , ソース = "uniref90" )
ESM2 埋め込みのキャッシュ (GPU 推論後)
点。 add_embedding ( "1abc.ptt" , esm_representations ,
モデル = "esm2_t33_650M_UR50D" 、レイヤー = - 1 、dtype = "float16" 、
シーケンスハッシュ = pt 。 embedding_sequence_hash ( data . sequence_tokens ))
.ptt ファイルから直接 Voltz2 を実行する
プロテインテンソルインポートからのBoltzAdapter
アダプター = BoltzAdapter ( "1abc.ptt" )
予測 = アダプター 。予測する(
"boltz_output/" ,
モデル = "boltz2" 、
拡散サンプル = 5 、
リサイクルステップ = 3 、
アクセラレータ = "GPU" 、
)
# -> ボルト出力

ut/predictions/1abc/1abc_model_0.cif (予測された構造)
# ->boltz_output/predictions/1abc/pae_*.npz (PAE マトリックス)
# ->boltz_output/predictions/1abc/plddt_*.npz (残基ごとの信頼度)
Python API
プロテインテンソルを pt としてインポート
# ------ 構造 ------
データ = pt 。読み取り ( "1abc.ptt" )
データ。 atom_positions 。形状番号 (N_atoms, 3) float32
データ。シーケンストークン 。形状番号 (N_res,) int32
データ。 backbone_positions 。形状番号 (N_res, 4, 3) float32 N/CA/C/O
データ。ボンドエッジインデックス 。形状 # (2, N_edges) int32 双方向
# バックボーンのみ (最速の構造負荷)
bb = pt 。 read_backbone ( "1abc.ptt" )
bb 。ポジション。形状 # (N_res, 4, 3)
# ボンドグラフのみ
債券 = pt 。 read_bonds ( "1abc.ptt" )
# ------ MSA ------
msa = pt 。 read_msa ( "1abc.ptt" 、ソース = "uniref90" )
うさ。トークン 。形状番号 (N_seq, N_res) int32
うさ。プロフィール。形状番号 (N_res, 23) float32
# ------ ペア機能 ------
点。 compute_and_store_ distances ( "1abc.ptt" ) # Ca-Ca 距離行列
点。 compute_and_store_contacts ( "1abc.ptt" 、しきい値 = 8.0 )
距離 = pt 。 read_pair_feature ( "1abc.ptt" , " distance_matrix " )
距離 。データ 。形状番号 (N_res, N_res, 1) float32
# 任意のペア テンソルを保存 (テンプレート特徴、MSA 共分散など)
点。追加

[切り捨てられた]

## Original Extract

AI-native tensor format (.ptt) for protein structure ML. Convert mmCIF/PDB structures or raw sequences once, then load backbone, MSA, embeddings, and bond graphs in milliseconds. Feeds AlphaFold/Boltz pipelines. Zarr-backed, lossless, PyTorch/JAX-ready. - mooreneural/HelixDB

GitHub - mooreneural/HelixDB: AI-native tensor format (.ptt) for protein structure ML. Convert mmCIF/PDB structures or raw sequences once, then load backbone, MSA, embeddings, and bond graphs in milliseconds. Feeds AlphaFold/Boltz pipelines. Zarr-backed, lossless, PyTorch/JAX-ready. · GitHub
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
mooreneural
/
HelixDB
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
32 Commits 32 Commits .github/ workflows .github/ workflows benchmarks benchmarks proteintensor proteintensor tests tests .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md boltz_benchmark.py boltz_benchmark.py paper.bib paper.bib paper.md paper.md pyproject.toml pyproject.toml View all files Repository files navigation
ProteinTensor is an AI-native biomolecular storage format designed to eliminate
the preprocessing bottleneck in modern structural biology machine learning pipelines.
Every time a researcher trains AlphaFold, Boltz, RoseTTAFold, OpenFold, or any
structure-prediction model, the same work happens before a single GPU operation runs:
for each protein in dataset:
parse mmCIF / PDB file (30 – 350 ms per structure)
extract sequence tokens
build atom coordinate arrays
construct backbone geometry
compute covalent bond graph
load or regenerate MSA (2 – 480 min with JackHMMER)
run ESM2 / ESM3 inference (seconds per protein on GPU)
compute distance matrices
...
-> finally: model.forward(features)
For a 100,000-structure training run this preprocessing costs thousands of CPU-hours
per epoch - most of it producing identical results every time. The mmCIF file has not
changed. The sequence has not changed. The physics has not changed. Yet every run
recomputes everything from scratch.
ProteinTensor solves this by converting the PDB entry once into a .ptt file - a
Zarr-backed, LZ4-compressed, memory-mappable store that holds every tensor a model
needs - and then loading those tensors directly at training time with zero parsing.
once: mmCIF -> ProteinTensor (.ptt)
always: .ptt -> model.forward()
Who This Is For
Structural biology researchers running AlphaFold 3, Boltz, or Chai-1 who spend
hours waiting for MSA generation before every new experiment.
ML engineers at pharma and biotech companies iterating over large structure
databases (PDB, AlphaFold Database, ESMAtlas) where I/O throughput is a training
bottleneck measured in wall-clock days.
Academic labs with limited GPU budgets who cannot afford to waste compute cycles
on re-parsing text files when those GPU-hours should go toward model training.
Software engineers building structural biology pipelines who want a single,
well-defined intermediate format that works with PyTorch, JAX, and NumPy without
writing custom loaders for every model.
ProteinTensor is to structural biology what Parquet is to analytics, what safetensors
is to model weights, and what ONNX is to model exchange - a common, open, high-
performance format that turns a recurring computational tax into a one-time cost.
Benchmark: Traditional Pipeline vs ProteinTensor
All timings are median over 30 rounds on an NVIDIA RTX 5080, CUDA 12.8, Python 3.11.
Proteins span the full range from a 76-residue domain to a 3,525-residue CRISPR enzyme.
Run python boltz_benchmark.py to reproduce.
Structure
Method
Res
MSA seqs
mmCIF parse
ptt: full
ptt: backbone
ptt: bonds
ptt: MSA
ptt: dist mx
1UBQ - Ubiquitin
X-ray
76
512
7.2 ms
2.8 ms
1.2 ms
0.7 ms
1.6 ms
0.8 ms
6LU7 - SARS-CoV-2 Mpro
X-ray
312
1,024
29.6 ms
2.9 ms
1.2 ms
0.7 ms
5.1 ms
2.0 ms
4HHB - Hemoglobin
X-ray
574
2,048
55.3 ms
2.9 ms
1.2 ms
0.7 ms
11.3 ms
3.5 ms
6M0J - ACE2 + RBD
Cryo-EM
791
2,048
74.7 ms
2.9 ms
1.2 ms
0.7 ms
14.7 ms
6.4 ms
6VXX - Spike trimer
Cryo-EM
2,916
8,192
283.4 ms
3.3 ms
1.3 ms
0.9 ms
208.3 ms
71.1 ms
6OHW - Cas12a
Cryo-EM
3,525
8,192
352.4 ms
3.3 ms
1.2 ms
1.0 ms
240.7 ms
104.5 ms
Column definitions
ptt: full - read() - all atoms, backbone, bonds, metadata
ptt: backbone - read_backbone() - N/CA/C/O coordinates + sequence only
ptt: bonds - read_bonds() - covalent graph only
ptt: MSA - read_msa() - MSA tokens + profile (loaded from .ptt cache)
ptt: dist mx - read_pair_feature("distance_matrix") - Ca-Ca distance matrix
Structure
Res
full
backbone
bonds
MSA
dist mx
1UBQ - Ubiquitin
76
3x
6x
11x
4x
9x
6LU7 - SARS-CoV-2 Mpro
312
10x
24x
43x
6x
15x
4HHB - Hemoglobin
574
19x
45x
78x
5x
16x
6M0J - ACE2 + RBD
791
26x
61x
102x
5x
12x
6VXX - Spike trimer
2,916
87x
223x
308x
1x*
4x
6OHW - Cas12a
3,525
108x
284x
370x
1x*
3x
*MSA speedup shown as 1x vs mmCIF parse because both are in the same time range for
large proteins - the real MSA comparison is vs JackHMMER generation (see below).
Feature assembly: time to prepare all tensors for model.forward()
Traditional = mmCIF parse + read MSA from A3M file. ProteinTensor = single .ptt read
with all features pre-cached (sequence, backbone, bonds, MSA, distance matrix,
ESM2 embedding).
Average speedup across all six structures: 34x for full feature assembly.
Same methodology across six high-value drug targets spanning KRAS oncology,
HIV antivirals, PD-L1 immunotherapy, p53, cardiovascular (PCSK9), and a full
IgG1 antibody. Numbers are consistent with the structural biology benchmark above.
Measured using ProteinDataset + ProteinDataset.collate() , loading structures into
padded batches ready for model.forward() . Single process, no prefetch workers.
Scale projection: 100,000 structures, one training epoch
These are projections , extrapolated from the measured per-structure timings
above - not end-to-end measurements at 100k scale.
MSA generation assumes 2.4 min/protein on a 32-core server (PDB90 database, standard
AlphaFold settings). ProteinTensor generates MSAs once and loads from the .ptt cache
on every subsequent run. The 4,000-hour figure is the real cost AlphaFold2 and Boltz
users pay to build training datasets from scratch.
Measured vs projected - read this. The 1,794x above is MSA generation
(building the alignment once with JackHMMER) and is a literature-based
projection , not something benchmarked here. What is measured on hardware is
the recurring per-epoch MSA load - reading a cached MSA from .ptt vs
re-parsing A3M text each epoch (against a vectorized A3M parser baseline):
3.4x-5.9x , growing with MSA depth. See
benchmarks/MSA_RESULTS.md . These are different
quantities; do not read the 1,794x as a measured load speedup.
A full-featured .ptt (8,192-sequence MSA + distance matrix + ESM2-650M embedding at
float16) averages 23x larger than the source mmCIF across the six benchmark structures.
The tradeoff is deliberate: pay disk space once to avoid paying GPU-hours and CPU-hours
on every training run. A structure-only .ptt with no cached features is smaller than
the source mmCIF.
pip install -e " .[dev] " # core + dev tools
pip install -e " .[cloud] " # adds fsspec, s3fs, gcsfs for remote reads
pip install -e " .[dev,cloud] " # everything
Requires Python >= 3.9, gemmi , zarr , numpy , click , rich .
proteintensor convert 1abc.cif 1abc.ptt
proteintensor info 1abc.ptt
Convert a sequence (no structure required)
For sequence-driven predictors like AlphaFold and Boltz, the primary input is a
sequence, not a structure. ProteinTensor can build a sequence-only .ptt (no
coordinates) directly from a raw string or a FASTA file:
proteintensor convert-seq MQIFVKTLTGKTITLEVEPSDTIENVKAKIQDKEGIPPDQQRLIFAGKQLEDG ubq.ptt
proteintensor convert-seq complex.fasta complex.ptt # multi-record FASTA -> multi-chain
import proteintensor as pt
data = pt . from_sequence ( "MQIFVKTLTGK..." , pdb_id = "UBQ" , chain_id = "A" )
data . has_structure # False - sequence-only entry
data . sequence_tokens # (N_res,) int32
pt . write ( data , "ubq.ptt" )
# FASTA: a single record -> one chain; multiple records -> multi-chain complex
data = pt . from_fasta ( "complex.fasta" )
Batch-convert a directory
Convert an entire directory of structures in parallel, with progress reporting.
Files that fail to parse are skipped and listed in the summary; already-converted
outputs are skipped by default.
proteintensor convert-dir ./pdb_files/ ./ptt_files/ # auto worker count
proteintensor convert-dir ./pdb_files/ ./ptt_files/ --workers 16 --recursive
proteintensor convert-dir ./pdb_files/ ./ptt_files/ --overwrite # rebuild existing
Benchmark against mmCIF
proteintensor benchmark 1abc.cif --rounds 20
Cache an MSA (after running JackHMMER / ColabFold)
import proteintensor as pt
msa = pt . from_a3m ( "1abc_uniref90.a3m" ,
tool = "jackhammer" , tool_version = "3.3.2" ,
database = "uniref90" , database_date = "2024-01" )
pt . add_msa ( "1abc.ptt" , msa , source = "uniref90" )
Cache ESM2 embeddings (after GPU inference)
pt . add_embedding ( "1abc.ptt" , esm_representations ,
model = "esm2_t33_650M_UR50D" , layer = - 1 , dtype = "float16" ,
sequence_hash = pt . embedding_sequence_hash ( data . sequence_tokens ))
Run Boltz2 directly from a .ptt file
from proteintensor import BoltzAdapter
adapter = BoltzAdapter ( "1abc.ptt" )
predictions = adapter . predict (
"boltz_output/" ,
model = "boltz2" ,
diffusion_samples = 5 ,
recycling_steps = 3 ,
accelerator = "gpu" ,
)
# -> boltz_output/predictions/1abc/1abc_model_0.cif (predicted structure)
# -> boltz_output/predictions/1abc/pae_*.npz (PAE matrix)
# -> boltz_output/predictions/1abc/plddt_*.npz (per-residue confidence)
Python API
import proteintensor as pt
# ------ Structure ------
data = pt . read ( "1abc.ptt" )
data . atom_positions . shape # (N_atoms, 3) float32
data . sequence_tokens . shape # (N_res,) int32
data . backbone_positions . shape # (N_res, 4, 3) float32 N/CA/C/O
data . bond_edge_index . shape # (2, N_edges) int32 bidirectional
# Backbone only (fastest structural load)
bb = pt . read_backbone ( "1abc.ptt" )
bb . positions . shape # (N_res, 4, 3)
# Bond graph only
bonds = pt . read_bonds ( "1abc.ptt" )
# ------ MSA ------
msa = pt . read_msa ( "1abc.ptt" , source = "uniref90" )
msa . tokens . shape # (N_seq, N_res) int32
msa . profile . shape # (N_res, 23) float32
# ------ Pair features ------
pt . compute_and_store_distances ( "1abc.ptt" ) # Ca-Ca distance matrix
pt . compute_and_store_contacts ( "1abc.ptt" , threshold = 8.0 )
dist = pt . read_pair_feature ( "1abc.ptt" , "distance_matrix" )
dist . data . shape # (N_res, N_res, 1) float32
# Store arbitrary pair tensors (template features, MSA covariance, …)
pt . add

[truncated]
