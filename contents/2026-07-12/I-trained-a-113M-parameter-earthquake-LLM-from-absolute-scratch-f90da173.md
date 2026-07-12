---
source: "https://github.com/jiazhe868/nanogpt-seis"
hn_url: "https://news.ycombinator.com/item?id=48885236"
title: "I trained a 113M-parameter earthquake LLM from absolute scratch"
article_title: "GitHub - jiazhe868/nanogpt-seis · GitHub"
author: "jzsfg"
captured_at: "2026-07-12T22:38:37Z"
capture_tool: "hn-digest"
hn_id: 48885236
score: 2
comments: 0
posted_at: "2026-07-12T21:58:38Z"
tags:
  - hacker-news
  - translated
---

# I trained a 113M-parameter earthquake LLM from absolute scratch

- HN: [48885236](https://news.ycombinator.com/item?id=48885236)
- Source: [github.com](https://github.com/jiazhe868/nanogpt-seis)
- Score: 2
- Comments: 0
- Posted: 2026-07-12T21:58:38Z

## Translation

タイトル: 113M パラメータの地震 LLM をまったくのゼロから訓練しました
記事タイトル: GitHub - jiazhe868/nanogpt-seis · GitHub
説明: GitHub でアカウントを作成して、jiazhe868/nanogpt-seis の開発に貢献します。

記事本文:
GitHub - jiazhe868/nanogpt-seis · GitHub
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
jiazhe868
/
nanogpt-seis
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
14 コミット 14 コミット アセット アセット チェックポイント チェックポイント configs configs data/処理されたデータ/処理された src src .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.mdrequirements.txt required.txt すべてのファイルを表示 リポジトリ ファイルナビゲーション
地震科学用の小さな GPT をトレーニングします。空のフォルダーから会話モデルまで、LLM ライフサイクル全体がブロックごとに説明されます。
2× NVIDIA A30 (48 GB) 上で、クロール→クリーン→トークン化→モデル→トレーニング→推論。
6 つの無料データ ソース → クロール → クリーン/デアップ → 16k BPE →
113M GQA+RoPE デコーダー → 2-GPU DDP トレーニング → ストリーミング推論。各ステージは、
以下のセクション。
nanoGPT-seis は教育用リポジトリです。大地震を起こそうとしているわけではない
モデル — 言語モデルの事前トレーニングのすべての段階を読みやすくしようとしています。
データの出所、データのクリーンアップと重複排除の方法、トークナイザーの仕組み
構築されたもの、Transformer がそのように見える理由、2 つの GPU でどのようにトレーニングされるのか、
そしてそれがどのように提供されるか。すべての設計上の決定とすべての数値が説明されています
(パープレキシティ、VRAM、トークン) は、このハードウェアで実際に測定したものです。
このコーパスには、地震/地震学のテキスト (オープンアクセス論文が含まれています) が混合されています。
Crossref+Unpaywall、arXiv/EarthArXiv プレプリント、「Earthquake Insights」サブスタック)
平易な言語で流暢に話すための一般的なテキスト (Wikipedia + FineWeb-Edu) について
ドメイン 24% / 一般 76%。焦点を絞ったコーパスにより、~100M パラメータのモデルが可能になります。
単一ノード上で真に流暢に動作するため、ループ全体が 1 日で終了するのがわかります。
一ヶ月。 (なぜ一般的な組み合わせなのか? §1 を参照してください。)
ステータス: 事前トレーニングのライフサイクルが完了しました - クロールスルー推論。
ステージ 4 — モデル (RoPE、GQA など)
ステージ 5 — トレーニング (DDP、VRAM、LR)
値
コーパス
533,248日

ocs · 4 億 8,570 万ワード · 8 億 2,270 万トレーニング トークン (≈2.4:1 一般:ドメイン)
モデル
113M パラメータ - デコーダのみ、GQA + RoPE + RMSNorm + SwiGLU
ハードウェア
2× NVIDIA A30 (各 24 GB)、bf16、DDP — 単一の RTX 3090/4090、またはより小さいバッチの 12 ～ 16 GB でも動作します (§7.5 )
コンテキストの長さ
4096 トークン
トレーニング
8,000 イター (約 3.8 エポック)、約 6.5 時間、約 2.9 秒/イッター
流暢さ (ビット/バイト、一般的なテキスト)
0.997 — 対ドメインのみベースの場合は 1.527 ( -35% )
推論
KV キャッシュ ストリーミング、最初のトークンまで約 176 ミリ秒、アンチリピート サンプラー
立ち止まってみる価値のある 3 つの発見:
より長いコンテキストが役に立ちました。制御された A/B (データは固定されたまま) — 4096 での再トレーニング
vs 1024 では、混乱度が約 11% 低下し (9.74 → 10.93 ドメインのみ)、わずか約 26% 増加しました。
ステップごとに計算 — 論文は 1024 トークンのウィンドウでは認識できない長距離構造を持っています。
モデルはそのコンテキストを使用します。 4096 トークンのウィンドウでは、ポジションでのトークンの損失
2048 ～ 4096 は 0 ～ 64 より 25% 低いです。これは、以前の数千件を条件としています。
トークン (§8 を参照)。
一般的なテキストを混ぜることで流暢さを取り戻します。 Wikipedia + FineWeb-Edu の追加 (~2.4:1)
一般:ドメイン) 一般的な散文のビット/バイトを紙のみのベースと比較して 35% 削減します。
以下のデータミックスの比較。
データ混合: ドメインのみ (v1) → 一般 + ドメイン (v2)
紙のみのベースは紙の台帳では流暢ですが、平文では繰り返しや一貫性がありません
散文。 Wikipedia + FineWeb-Edu の ~5 億 4,000 万のトークンを追加 (→ ~8 億 2,300 万のトレイン トークン、
~2.4:1 一般:ドメイン、~3.8 エポック — 示されている ~4 エポックのリピート バジェット内
Muennighoff et al., 2023 によるほぼロスレス) は v2 を与えます。
バイトあたりのビット数で測定されます (トークナイザーに依存しないため、v1↔v2 は公平です。
src/compare_models.py ):
v2 は一般的な散文に対してはるかに流暢であり (-35% ビット/バイト)、一貫した内容を生成します。
地震以外のテキスト。v1 では意味不明の内容が出力されますが、ドメインの鮮明さは犠牲になります。
(+22%) —

古典的な流暢さと専門化のトレードオフです。この流暢なベースは、
SFT の正しい出発点。基本的な事前トレーニングだけではチャットが生成されることはありません。ドメインのみ
重みは、checkpoints/ckpt_v1_domain.pt として保持されます。
事前トレーニングされた 113M チェックポイントは、Hugging Face Hub でホストされています。
jiazhe868/nanogpt_seis 。
# 環境 (動作する CUDA-12.4 PyTorch; 以下の注を参照)
conda は nanogpt_seis をアクティブにします
pip install -r 要件.txt
# チェックポイント、トークナイザー、およびモデル構成を、によって予期されるパスにダウンロードします
# src/inference.py
ハグフェイス-cli ダウンロード jiazhe868/nanogpt_seis \
チェックポイント/ckpt.pt \
データ/トークン化/tokenizer.json \
データ/トークン化/meta.json \
configs/gpt120m_ctx4k.yaml \
--local-dir 。
python -m src.inference --prompt 「2011 年東北地方太平洋沖地震」
2.2 事前学習を最初から再現する
# 環境 (動作する CUDA-12.4 PyTorch; 以下の注を参照)
conda は nanogpt_seis をアクティブにします
pip install -r 要件.txt
# --- パイプライン全体を実行 ---
# 地震領域ソース
python -m src.crawl.wikipedia --max-pages 500 # 地震のタイトルのページ
python -m src.crawl.fulltext --ジャーナルごと 3000 --全体 30000 --workers 64
python -m src.crawl.preprints --arxiv 3000 --eartharxiv 2000
python -m src.crawl.substack --max 500
# 流暢さのための一般的なテキストの組み合わせ (~5 億 4,000 万のトークン: Wikipedia + FineWeb-Edu)
python -m src.crawl.general --wiki-tokens 300000000 --fineweb-tokens 240000000
python -m src.process.build_corpus --val-frac 0.005 # クリーン・重複排除・分割
python -m src.tokenizer.train_bpe --vocab-size 16384 # トークナイザーをトレーニングします
python -m src.tokenizer.encode # → uint16 シャード
torchrun --standalone --nproc_per_node=2 \
-m src.train --config configs/gpt120m_ctx4k.yaml # 2 つの GPU でトレーニングします
python -m src.inference --prompt 「2011 年東北地方太平洋沖地震」 # ライブ配信
⚠️環境問題

はー。ドライバーよりも新しい CUDA 用に構築された PyTorch
サイレントに cuda.is_available() == False を報告し、CPU にフォールバックします。で確認してください
python -c "import torch; print(torch.cuda.is_available())" — True を出力する必要があります。
このプロジェクトでは、CUDA-12.5 ドライバーと一致するように torch 2.6.0+cu124 を使用します。
目標: 無料のソースから大規模で合法的な全文地震コーパスを収集します。
コード: src/crawl/ 。
3.1 ソースとそれぞれの取得方法
ソース
モジュール
私たちが引っ張るもの
仕組み
研究論文
フルテキスト.py
地震論文全文PDFをOA
Crossref (DOI) → Unpaywall (OA PDF) → ダウンロード → 抽出
プレプリント
プレプリント.py
arXiv + EarthArXiv 全文
arXiv API + OSF/DOI → PDF
ウィキペディア
ウィキペディア.py
「地震」というタイトルのページ
MediaWiki API 平文抽出
サブスタック
サブスタック.py
「地震洞察」記事
アーカイブ API + HTML 本文の解析
一般的なテキスト
一般的な.py
Wikipedia + FineWeb-Edu (流暢性ミックス)
トークンバジェットにストリーミングする HF データセット
なぜ一般的なテキストを混合するのでしょうか?研究論文のみでトレーニングされた約 1 億 1,300 万のモデル
紙の台帳では流暢に話せるようになるが、平易な散文では反復的で支離滅裂になる。
そして 2 億 4,000 万のトークンは最適な計算量をはるかに下回っています (Hoffmann et al., 2022)。
したがって、最大 2 億 4,000 万のトークンを追加します
ウィキペディア (百科事典) +
FineWeb-Edu
(Penedo et al., 2024 ; 品質フィルタリングされた教育ウェブ)
~1:1 の一般:ドメインの混合の場合 — 計画された SFT ステージで使用できる流暢なベース
会話的なものにする。 (基本的な事前トレーニングだけではチャットが生成されることはありません。それが SFT です。)
すべてのドキュメントは 1 つのスキーマ ( src/crawl/common.py ) に正規化されます。
@データクラス
クラスドキュメント:
ソース: str # "フルテキスト" | "arxiv" | 「ウィキペディア」 | ...
id : str # ソースごとの安定した ID (重複除去に使用)
タイトル : ストラ
text : str # トークン化するクリーンなボディ
URL : str = ""
日付: str = ""
追加: dict = フィールド (default_factory = dict ) # 会場、引用者、フル

l_テキスト、...
3.2 Web クローリングの実際の仕組み
論文を見つける (Crossref)。 Crossref は、無料で最大 1 億 5,000 万の学術作品をインデックスします。
寛大な API。深いカーソルを使って地震雑誌の記事をページめくります
(オフセット制限なし)、ジャーナル ISSN でフィルタリング:
# src/crawl/fulltext.py — iter_crossref()
params = { "行" : 1000 、 "カーソル" : カーソル 、
"filter" : "type:journal-article,issn:0094-8276" , # 例: GRL
"query.bibliographic" : "earthquake" , "mailto" : EMAIL }
開いている PDF (Unpaywall) を見つけます。 DOI は PDF ではありません。 Unpaywall (無料、最大 100,000/日)
DOI をその合法的なオープンアクセス コピーにマッピングします。リポジトリ (緑色の OA) の場所を試します
まず、パブリッシャーのリンクはボットによってブロックされるスタブであることがよくあります。
# 個のリポジトリ コピーは、発行者リンクよりもはるかに確実にダウンロードされます
loc の場合は prio = 0。 get ( "host_type" ) == "リポジトリ" else 1
丁寧に並行してダウンロードします。 PDF は多くのホストから提供されるため、スレッドを使用します
プールしますが、ホストごとにスロットルします - 異なるサーバーが同時にダウンロードします。
単一サーバーはレート制限のままです:
クラス HostThrottle : # src/crawl/fulltext.py
def wait ( self 、 host ):
自分自身と。 _guard : # このホストのロックを取得/作成します
host_lock = self 。 _ロック 。 setdefault ( host , threading . Lock ())
自分自身。 _最後 。 setdefault (ホスト, 0.0)
with host_lock : # このホストのみをシリアル化します
デルタ = 自分自身。 min - ( time . 単調 () - self . _last [ host ])
デルタ > 0 の場合: 時間。スリープ (デルタ)
自分自身。 _last [ホスト] = 時間。単調（）
すべてのダウンロードを検証します。すべての「OA PDF」が本物であるわけではありません - 多くは 5 KB のアンチボットです
ランディングページ。十分なテキストを含む実際の PDF である場合にのみ、ダウンロードを受け入れます。
そうでない場合は pdf_bytes 。 beginswith ( b"%PDF" ): return None # HTML / スタブ
ドキュメントの場合。 page_count < 2 : なしを返します # 表紙
if len ( text ) < min_chars : return None # 抽出されたものが少なすぎます
無駄にしない

予算。一部のジャーナル (サイエンス、ネイチャー) は、ほぼすべてが
paywalled — 何千もの DOI をスキャンすると、Unpaywall の予算がゼロになります。
全文。低収量アボート ゲートは、ヒット率が一定以下になるとジャーナルをスキップします。
しきい値:
スキャンされた場合 >= abort_after および got_ft / max ( 1 , scanned ) < min_hit :
Break # この会場にはこれ以上 API を呼び出す価値はありません
再開可能。出力は行ごとに追加され、すでに取得された ID はスキップされます。
再起動すると、数時間のクロールが中断されても存続します。
done = _load_done_ids ( out ) # JSONL にすでに存在する ID
...
if iid inned : continue # すでに行っている作業をスキップします
戦争の話（なぜCrossref + Unpaywall）。このプロジェクトはもともと論文を列挙していました
OpenAlex 経由、ビルド途中で有料クレジット モデルに変更 - 毎日無料
約 100 件のリクエストで予算がなくなりました。 Crossref + Unpaywall は無料で堅牢です
交換品であり、コードに同梱されています。レッスン — データソースを固定する
仮定を作成し、クローラーを再開可能にする — が設計に組み込まれています。
全文の生成率はソースによって大きく異なります: arXiv ~99% (設計によりオープン)、広範な OA
プール ~15%、ペイウォールジャーナル ~0% (抽象フォールバック)。ネットコーパス: ~20k
論文全文 + 〜 26,000 件の要約。
3.3 一般的なパターン — スレッドセーフな BFS クローラー
上のクローラーは API 駆動型ですが、その下の形状は古典的なものです。
ブレア

[切り捨てられた]

## Original Extract

Contribute to jiazhe868/nanogpt-seis development by creating an account on GitHub.

GitHub - jiazhe868/nanogpt-seis · GitHub
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
jiazhe868
/
nanogpt-seis
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits assets assets checkpoints checkpoints configs configs data/ processed data/ processed src src .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md requirements.txt requirements.txt View all files Repository files navigation
Train a small GPT for earthquake science — the entire LLM lifecycle, from a blank folder to a talking model, explained block by block.
Crawl → Clean → Tokenize → Model → Train → Infer, on 2× NVIDIA A30 (48 GB).
Six free data sources → crawl → clean/dedup → 16k BPE →
113M GQA+RoPE decoder → 2-GPU DDP training → streaming inference. Each stage is a
section below.
nanoGPT-Seis is a teaching repository. It is not trying to be a great earthquake
model — it is trying to make every stage of pretraining a language model legible:
where the data comes from, how it is cleaned and deduplicated, how a tokenizer is
built, why the Transformer looks the way it does, how it is trained across two GPUs,
and how it is served. Every design decision is explained, and every number
(perplexity, VRAM, tokens) is one we actually measured on this hardware.
The corpus mixes earthquake / seismology text (open-access papers via
Crossref+Unpaywall, arXiv/EarthArXiv preprints, the "Earthquake Insights" Substack)
with general text (Wikipedia + FineWeb-Edu) for plain-language fluency — about
24% domain / 76% general. A focused corpus lets a ~100M-param model become
genuinely fluent on a single node, so you can see the whole loop close in a day, not
a month. (Why the general mix? See §1 .)
Status: the pretraining lifecycle is complete — crawl through inference.
Stage 4 — The model (RoPE, GQA, …)
Stage 5 — Training (DDP, VRAM, LR)
value
Corpus
533,248 docs · 485.7M words · 822.7M training tokens (≈2.4:1 general:domain)
Model
113M params — decoder-only, GQA + RoPE + RMSNorm + SwiGLU
Hardware
2× NVIDIA A30 (24 GB each), bf16, DDP — also runs on a single RTX 3090/4090, or 12–16 GB with a smaller batch ( §7.5 )
Context length
4096 tokens
Training
8,000 iters (~3.8 epochs), ~6.5 h, ~2.9 s/iter
Fluency (bits/byte, general text)
0.997 — vs 1.527 for a domain-only base ( −35% )
Inference
KV-cached streaming, ~176 ms to first token, anti-repeat sampler
Three findings worth pausing on:
Longer context helped. A controlled A/B (data held fixed) — retraining at 4096
vs 1024 dropped perplexity ~11% (9.74 → 10.93 domain-only) for only ~26% more
compute per step — papers have long-range structure a 1024-token window can't see across.
The model uses that context. In a 4096-token window, loss on tokens at positions
2048–4096 is 25% lower than at 0–64 — it conditions on thousands of preceding
tokens (see §8 ).
A general-text mix restores fluency. Adding Wikipedia + FineWeb-Edu (~2.4:1
general:domain) cut bits/byte on general prose by 35% vs a paper-only base — see the
data-mix comparison below.
Data mix: domain-only (v1) → general + domain (v2)
A paper-only base is fluent in paper-register but repetitive/incoherent in plain
prose. Adding ~540M tokens of Wikipedia + FineWeb-Edu (→ ~823M train tokens,
~2.4:1 general:domain, ~3.8 epochs — within the ~4-epoch repeat budget shown to be
near-lossless by Muennighoff et al., 2023 ) gives v2.
Measured with bits-per-byte (tokenizer-independent, so v1↔v2 is fair;
src/compare_models.py ):
v2 is far more fluent on general prose (−35% bits/byte) and generates coherent
non-earthquake text where v1 emits gibberish, at the cost of some domain sharpness
(+22%) — the classic fluency↔specialization trade-off. This fluent base is the
right starting point for SFT; base pretraining alone never yields chat. Domain-only
weights are kept as checkpoints/ckpt_v1_domain.pt .
The pretrained 113M checkpoint is hosted on the Hugging Face Hub:
jiazhe868/nanogpt_seis .
# environment (a working CUDA-12.4 PyTorch; see the note below)
conda activate nanogpt_seis
pip install -r requirements.txt
# download the checkpoint, tokenizer, and model config into the paths expected by
# src/inference.py
huggingface-cli download jiazhe868/nanogpt_seis \
checkpoints/ckpt.pt \
data/tokenized/tokenizer.json \
data/tokenized/meta.json \
configs/gpt120m_ctx4k.yaml \
--local-dir .
python -m src.inference --prompt " The 2011 Tohoku earthquake "
2.2 Reproduce pretraining from scratch
# environment (a working CUDA-12.4 PyTorch; see the note below)
conda activate nanogpt_seis
pip install -r requirements.txt
# --- run the whole pipeline ---
# earthquake-domain sources
python -m src.crawl.wikipedia --max-pages 500 # earthquake-titled pages
python -m src.crawl.fulltext --per-journal 3000 --broad 30000 --workers 64
python -m src.crawl.preprints --arxiv 3000 --eartharxiv 2000
python -m src.crawl.substack --max 500
# general-text mix for fluency (~540M tokens: Wikipedia + FineWeb-Edu)
python -m src.crawl.general --wiki-tokens 300000000 --fineweb-tokens 240000000
python -m src.process.build_corpus --val-frac 0.005 # clean · dedup · split
python -m src.tokenizer.train_bpe --vocab-size 16384 # train the tokenizer
python -m src.tokenizer.encode # → uint16 shards
torchrun --standalone --nproc_per_node=2 \
-m src.train --config configs/gpt120m_ctx4k.yaml # train on 2 GPUs
python -m src.inference --prompt " The 2011 Tohoku earthquake " # streams live
⚠️ Environment gotcha. A PyTorch built for a newer CUDA than your driver will
silently report cuda.is_available() == False and fall back to CPU. Verify with
python -c "import torch; print(torch.cuda.is_available())" — it must print True .
This project uses torch 2.6.0+cu124 to match a CUDA-12.5 driver.
Goal: assemble a large, legal, full-text earthquake corpus from free sources.
Code: src/crawl/ .
3.1 The sources and how each is fetched
source
module
what we pull
mechanism
Research papers
fulltext.py
OA full-text PDFs of earthquake papers
Crossref (DOIs) → Unpaywall (OA PDF) → download → extract
Preprints
preprints.py
arXiv + EarthArXiv full text
arXiv API + OSF/DOI → PDF
Wikipedia
wikipedia.py
pages titled "earthquake"
MediaWiki API plaintext extracts
Substack
substack.py
"Earthquake Insights" articles
archive API + HTML body parse
General text
general.py
Wikipedia + FineWeb-Edu (fluency mix)
HF datasets streaming to a token budget
Why a general-text mix? A ~113M model trained only on research papers
becomes fluent in paper-register but repetitive and incoherent in plain prose,
and 240M tokens is far below compute-optimal ( Hoffmann et al., 2022 ).
So we add ~240M tokens of
Wikipedia (encyclopedic) +
FineWeb-Edu
( Penedo et al., 2024 ; quality-filtered educational web)
for a ~1:1 general:domain mix — a fluent base that the planned SFT stage can then
make conversational. (Base pretraining alone never yields chat; that's SFT.)
Every document is normalized to one schema ( src/crawl/common.py ):
@ dataclass
class Doc :
source : str # "fulltext" | "arxiv" | "wikipedia" | ...
id : str # stable per-source id (used for dedup)
title : str
text : str # the cleaned body we will tokenize
url : str = ""
date : str = ""
extra : dict = field ( default_factory = dict ) # venue, cited_by, full_text, ...
3.2 How the web crawling actually works
Finding the papers (Crossref). Crossref indexes ~150M scholarly works with a free,
generous API. We page through earthquake journal-articles with a deep cursor
(no offset limit), filtered by journal ISSN:
# src/crawl/fulltext.py — iter_crossref()
params = { "rows" : 1000 , "cursor" : cursor ,
"filter" : "type:journal-article,issn:0094-8276" , # e.g. GRL
"query.bibliographic" : "earthquake" , "mailto" : EMAIL }
Finding the open PDF (Unpaywall). A DOI is not a PDF. Unpaywall (free, ~100k/day)
maps a DOI to its legal open-access copies. We try repository (green OA) locations
first — publisher links are frequently bot-blocked stubs:
# repository copies download far more reliably than publisher links
prio = 0 if loc . get ( "host_type" ) == "repository" else 1
Downloading in parallel, politely. PDFs come from many hosts, so we use a thread
pool but throttle per host — different servers download concurrently while any
single server stays rate-limited:
class HostThrottle : # src/crawl/fulltext.py
def wait ( self , host ):
with self . _guard : # get/create this host's lock
host_lock = self . _locks . setdefault ( host , threading . Lock ())
self . _last . setdefault ( host , 0.0 )
with host_lock : # serialize only this host
delta = self . min - ( time . monotonic () - self . _last [ host ])
if delta > 0 : time . sleep ( delta )
self . _last [ host ] = time . monotonic ()
Validating every download. Not every "OA PDF" is real — many are 5 KB anti-bot
landing pages. We accept a download only if it is a real PDF with enough text:
if not pdf_bytes . startswith ( b"%PDF" ): return None # HTML / stub
if doc . page_count < 2 : return None # cover page
if len ( text ) < min_chars : return None # too little extracted
Not wasting the budget. Some journals (Science, Nature) are almost entirely
paywalled — scanning thousands of their DOIs would burn the Unpaywall budget for zero
full text. A low-yield abort gate skips a journal once its hit-rate stays under a
threshold:
if scanned >= abort_after and got_ft / max ( 1 , scanned ) < min_hit :
break # this venue isn't worth more API calls
Resumable. Output is appended line-by-line and already-fetched ids are skipped on
restart, so a multi-hour crawl survives interruption:
done = _load_done_ids ( out ) # ids already in the JSONL
...
if iid in done : continue # skip work we already have
War story (why Crossref + Unpaywall). This project originally enumerated papers
via OpenAlex, which changed to a paid credit model mid-build — the free daily
budget ran out after ~100 requests. Crossref + Unpaywall is the free, robust
replacement, and it is what the code ships with. The lesson — pin your data source
assumptions and make the crawler resumable — is baked into the design.
Full-text yield varies a lot by source: arXiv ~99% (open by design), the broad OA
pool ~15%, paywalled journals ~0% (abstract fallback). Net corpus: ~20k
full-text papers + ~26k abstracts.
3.3 The general pattern — a thread-safe BFS crawler
The crawler above is API-driven, but the shape underneath is the classic one: a
brea

[truncated]
