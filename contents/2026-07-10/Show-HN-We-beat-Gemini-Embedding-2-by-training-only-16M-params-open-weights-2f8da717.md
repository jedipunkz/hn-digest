---
source: "https://huggingface.co/EximiusLabs/fusion-embedding-1-2b-preview"
hn_url: "https://news.ycombinator.com/item?id=48854893"
title: "Show HN: We beat Gemini Embedding 2 by training only 16M params (open weights)"
article_title: "EximiusLabs/fusion-embedding-1-2b-preview · Hugging Face"
author: "abtonmoy"
captured_at: "2026-07-10T02:49:15Z"
capture_tool: "hn-digest"
hn_id: 48854893
score: 6
comments: 0
posted_at: "2026-07-10T02:05:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: We beat Gemini Embedding 2 by training only 16M params (open weights)

- HN: [48854893](https://news.ycombinator.com/item?id=48854893)
- Source: [huggingface.co](https://huggingface.co/EximiusLabs/fusion-embedding-1-2b-preview)
- Score: 6
- Comments: 0
- Posted: 2026-07-10T02:05:07Z

## Translation

タイトル: Show HN: 1,600 万パラメータのみをトレーニングすることで Gemini Embedding 2 に勝利しました (オープンウェイト)
記事タイトル: EximiusLabs/fusion-embedding-1-2b-preview · 顔を抱きしめる
説明: 私たちは、オープンソースとオープン サイエンスを通じて人工知能を進歩させ、民主化する旅の途中にあります。

記事本文:
EximiusLabs/fusion-embedding-1-2b-preview · 顔を抱きしめる
ハグ顔モデル
EximiusLabs / fusion-embedding-1-2b-preview いいね 4 Eximius Labs をフォロー 4
特徴抽出セーフテンサー 英語 fusion-embedding-connector embeddings マルチモーダル オーディオ検索 matryoshka qwen3-vl ライセンス: cc-by-nc-4.0 モデル カード ファイル ファイルとバージョン xet コミュニティ バケットにコピー 新しいアーキテクチャ
評価 クロスモーダル検索 - 統合埋め込みモデルとの比較
オーディオテキスト検索 — 専門の CLAP モデルとの比較
モデルは 1 つです。 1 つのベクトル空間。テキスト、画像、ビデオ、オーディオ、そして PDF。
最先端の機能を拡張するオープンウェイト マルチモーダル埋め込みモデル
単一のベースウェイトを変更することなく、オーディオを含むビジョン言語埋め込みベース。
Fusion Embedding 1 は Qwen3-VL-Embedding-2B を拡張します
オーディオモダリティ付き。トレーニングされたコネクタ (約 1,600 万のパラメーター) のマップが凍結されます
Qwen2.5-Omni オーディオタワー機能を搭載
ベースモデルの埋め込み空間。基本モデル自体は変更されていません。結果はシングルです
テキスト、画像、ビデオ、オーディオをカバーする埋め込みスペース。検索はサポートされています。
モダリティ間の任意の方向。
オーディオ↔テキストに関して測定したすべての統合埋め込みモデルをリードします。シングルで
クロスモーダル プロトコル、このモデルは ImageBind、LanguageBind、Gemini を超えています
オーディオ↔テキストに 2 を両方向で埋め込み、両方の言語に束縛されたベースラインをオンにします。
緊急音声↔画像（以下の完全な表）。
無改造ベース。コネクタのみがトレーニングされます。基本モデルのパラメータは次のとおりです。
オリジナルのリリースとバイトが同一であるため、テキスト/画像/ビデオの検索パフォーマンスが向上します。
(MMEB-V2) は変更なく引き継がれます。
緊急のクロスモーダル調整。コネクタは音声とテキストのみでトレーニングされています
ペア。それでも、音声→画像の検索は 696 の VGGSound 候補に対して R@10 0.407 に達します。
(確率: 0.014)

トレーニングではオーディオとビジュアルのペアがありません - テキストに合わせてオーディオを配置します
ベースがモダリティ間ですでに共有しているスペース。
マトリョーシカの表現。埋め込みは {2048、1536、1024、512、256、128、
64} 次元（繰り込みあり）。
コンパクトな配布。このリポジトリにはコネクタと正規化が同梱されています
統計 (~60 MB);フローズン タワーは元のリポジトリからダウンロードされます。
このリポジトリに表示されるパラメータ数 (16.4M) は、トレーニングされたコネクタです。
model.safetensors と .pt チェックポイントには同じ重みが含まれています。 inference.py
.pt をロードします。
これはリサーチ プレビューであり、現在 v0.3 : v0.2 の対照段階 (484K)
ペア) に続いて、AudioCaps トレイン スプリットでのコネクタのみのドメイン内微調整が行われます。
以前のバージョンも、v0.1-preview および v0.2-preview タグを介してダウンロードできます。
v0.3-preview は現在のバージョンを固定します。以下にすべてを比較します。ビルドする場合はタグを固定します
このモデルでは。
知覚リサンプラー (幅 384、64 の潜在クエリ) は、フリーズされたオーディオタワー フレームを変換します。
基本モデルの入力埋め込み空間に。その出力は、次のプレースホルダーの位置を占めます。
基本モデルのイメージ トークン メカニズムを反映した入力ストリーム。トレーニングというのは、
対照的（マトリョーシカの梯子上の情報、対称、完全なコーパス付き）
凍結テキストのネガティブ バンク — v0.2 で 484K キャプション）ベース モデルのテキストに対して
ネイティブ入力形式での埋め込み。 v0.3 では、コネクタのみの 2 番目の微調整が追加されています
AudioCaps トレイン分割のステージ (学習率を下げて 400 ステップ)、ウォームスタート
v0.2 チェックポイントから。
入力の書式設定。すべての入力は、基本モデルのチャット テンプレート形式を使用します (説明は
システムターン、ユーザーターンのコンテンツ、最後のトークンプーリング）。埋め込み品質は
この書式設定に敏感です。 inference.py のテンプレートを使用します。

クロスモーダル向け
ランキング、ギャラリーのモダリティごとの平均中心化が推奨されます ( FusionEmbedder.center )。
クロスモーダル検索 - 統合埋め込みモデルとの比較
VGGSound-AV、696 オーディオ/ビデオ フレーム ペア (チャンス R@10 = 0.014)。 R@10 は次のように表示されます
オーディオ側 → その他 / その他 → オーディオ側:
ImageBind は音声と画像のペアを直接トレーニングするため、そのペアが監視された方向になります。
オーディオとテキストの配置が明らかになります。 LanguageBind は言語に対して音声をトレーニングします (その
音声↔テキストは監視されています。表示されている値は、オーディオ ブランチの値を使用した場合の最良の読み出し値です。
独自のテキストタワー);その音声↔画像が出現します。このモデルは音声とテキストのみでトレーニングします。その
オーディオとイメージの整合性が現れています。すべてのモデルは同一のクリップ、フレーム、および
リリースされた imagebind_huge チェックポイントとリビジョン固定の LanguageBind を使用したスコアリング
チェックポイント (LanguageBind_Audio_FT + LanguageBind_Image)。 LanguageBind に関する注意:
ブランチは、分岐するテキスト タワーの別々のコピーを微調整します (キャプションのコサインを意味します)
音声ブランチと画像ブランチのテキスト埋め込みの間の 0.55) — ブランチ間のバインディング
これは、その出現した音声↔画像スコアと一致しています。このモデルは共有されています
構造によってスペースがドリフトすることはありません (ベースは凍結されており、トレーニングの実行ごとにアサートされます)
パラメータレベルのアイデンティティ）。 Gemini Embedding 2 は、Google のネイティブなマルチモーダル埋め込みです
API (1 つのスペース内のテキスト/画像/ビデオ/オーディオ)。文書化されたデフォルトの呼び出しで評価されます。
(モデル ID gemini-embedding-2 、3072-d ネイティブ出力、インライン オーディオ + 画像 + テキスト、
google-genai 2.10.0) 表示された評価日。その日以降、API モデルは変更される可能性があります。
共通の注意点: 評価キャプションはモデルによって生成されるため、モデルに有利になる可能性があります。
そのテキスト タワーはそのキャプション スタイルを共有しており、すべてのモデルが同一の入力を受け取りました。
フルオーディオ→i

メイジ メトリクス (モダリティごとの平均中心ギャラリー - によって実装される読み出し)
FusionEmbedder.center ;確率 R@10 = 0.014):
v0.3 のドメイン内微調整では、緊急音声→画像調整に最大 1 ポイントのコストがかかります。
オーディオ↔テキストの改善（クロスモーダル表を参照）。音声→画像の場合は v0.2 も引き続き利用可能
が主な使用例です。
音声→画像の検索はどのようなものですか。これらの数字は単なる集計ではありません。
検索は音によって整理されます。 VGGSound-696 の実例 (v0.2 チェックポイント)
(左にクリップのフレームをクエリ、右に取得した上位 5 つの画像、緑 = クリップの正確なフレーム):
VGGSound データセット (CC-BY-4.0) のフレーム例。評価の説明のために示しています。
直接ヒット — クリップ自体のフレームが、同じ種類のシーンの中でトップ 5 で返されます。
右近傍 - 正確なフレームのランクは下位 (多くの場合、まだ悪い) ですが、上位にあります。
結果は正しいサウンド カテゴリです。
MTEB チームの大規模オーディオ埋め込みベンチマークの 10 のタスクについて
(mteb 2.18.0、v0.2 チェックポイント; 2026 年 7 月 9 日時点のライブ リーダーボードに対するランク、21 ～ 65
タスクごとのモデル): UrbanSound8K T2A 取得 #3、Ravdes zero-shot #4、FSD2019Kaggle #6
(公開のみ — テスト クリップの 13.6% は、で使用されている FSD50K 開発分割に表示されます)
トレーニング、Freesound ID によって検証されるため、公式提出からは差し控えられます)、
北京オペラ #6、スピーチ/音楽タスクのミッドフィールド配置では、このモデルは決してありませんでした
のために訓練されました。公式リーダーボードの提出が進行中です。
オーディオテキスト検索 — 専門の CLAP モデルとの比較
AudioCaps テスト — 883 クリップ、クリップごとに 5 つの参照キャプション、リコールは次のように計算されます。
参照に対する最小ランク:
CLAPファミリーモデルは両方のエンコーダーをエンドツーエンドで微調整し、AudioCapsとClothoを含みます
トレーニングデータ。このモデルでは、両方のタワーをフリーズしたままにして、コネクタのみをトレーニングします。
クロト v2.1 評価

ション — 1,045 クリップ × 5 リファレンス、ゼロショット (クロソを除く)
トレーニング データから):
v0.3 のドメイン内 AudioCaps ステージは、ゼロショット Clotho A→T の 1.5 ポイントを交換します。
AudioCaps は上記を上回ります。 T→A はどちらの設定でも改善されます。
テキスト、画像、ビデオのベンチマークは、基本モデルの公開されている MMEB-V2 の結果です。
この拡張機能の影響を受けません。
# pip install git+https://github.com/Eximius-Labs/fusion-embedding-1 (+ トランスフォーマー、トーチビジョン、ピロー)
推論インポート FusionEmbedder から
fe = FusionEmbedder.from_pretrained( "EximiusLabs/fusion-embedding-1-2b-preview" ,
デバイス= "cuda" )
# またはバージョンを固定します:reviation="v0.3-preview" (current) / "v0.2-preview" / "v0.1-preview"
a = fe.embed_audio( "dog_barking.wav" ) # [2048]
t = fe.embed_text( "雨が降ると犬が吠える" ) # [2048]
i = fe.embed_image( "犬の写真.jpg" ) # [2048]
print ((a @ t), (a @ i), (t @ i)) # コサイン類似度
a256 = fe.embed_audio( "dog_barking.wav" , dim= 256 ) # マトリョーシカの切り詰め
トレーニングデータとライセンス
v0.2 は、約 484K のオーディオとキャプションのペアでトレーニングされました。フル AudioCaps トレイン スプリット (45K)、
FSD50K、WavCaps/AudioSet_SL、および LAION-FreeSound の 318K クリップ サブセット、10 秒を使用
トレーニング ウィンドウ (長いクリップのランダム クロップ)。 v0.3 は v0.2 チェックポイントを継続します。
AudioCaps トレイン スプリットのみで 400 ステップの微調整が可能です。 v0.1 では、131K ペアのサブセットが使用されていました。
同じソース。このミックスには YouTube から提供され、研究ライセンスを取得したコーパスが含まれているため、プレビューは
CC-BY-NC-4.0 に基づいてリリースされています。評価セット (AudioCaps テスト、Clotho、VGGSound、
ESC-50) は、クリップ ID によるトレーニングから除外されます。
サウンドイベントデータに基づいてトレーニングされています。スピーチの内容、話者の属性、音楽の説明
命令分類法によってサポートされていますが、同等の品質にまでトレーニングされていません。
英語のキャプション。 16 kHz モノラル入力。 1回あたり30秒

ウィンドウ (長いオーディオはチャンク化されます)。
このチェックポイントでは、音声テキストの取得は完全に微調整された CLAP ファミリー モデルを下回っています。
(「評価」を参照)。
さらなるコーパスのスケーリング、音声および音楽のカバレッジ、商用ライセンスのリリース層、
そして8Bモデル。
@ソフトウェア{fusion_embedding_2026,
title = {Fusion Embedding 1: テキスト用の統合された埋め込みスペース,
画像、ビデオ、オーディオ},
著者 = {トンモイ、アブドゥル・バシット}、
年 = {2026}、
URL = {https://github.com/Eximius-Labs/fusion-embedding-1}
}
AudioCaps、WavCaps、
そしてFSD50K。
セーフテンサー モデル サイズ 16.4M パラメータ テンソル タイプ F32 · ファイル情報
推論プロバイダーの新しい機能抽出 このモデルは、どの推論プロバイダーによっても展開されていません。 🙋 1 プロバイダーのサポートを依頼する EximiusLabs/fusion-embedding-1-2b-preview のモデル ツリー

## Original Extract

We’re on a journey to advance and democratize artificial intelligence through open source and open science.

EximiusLabs/fusion-embedding-1-2b-preview · Hugging Face
Hugging Face Models
EximiusLabs / fusion-embedding-1-2b-preview like 4 Follow Eximius Labs 4
Feature Extraction Safetensors English fusion-embedding-connector embeddings multimodal audio retrieval matryoshka qwen3-vl License: cc-by-nc-4.0 Model card Files Files and versions xet Community Copy to bucket new Architecture
Evaluation Cross-modal retrieval — versus unified embedding models
Audio–text retrieval — versus specialist CLAP models
One model. One vector space. Text, image, video, audio — and PDF.
An open-weight multimodal embedding model that extends a state-of-the-art
vision-language embedding base with audio — without modifying a single base weight.
Fusion Embedding 1 extends Qwen3-VL-Embedding-2B
with an audio modality. A trained connector (~16M parameters) maps frozen
Qwen2.5-Omni audio-tower features into the
base model's embedding space; the base model itself is unmodified. The result is a single
embedding space covering text, images, video, and audio , with retrieval supported in
any direction between modalities.
Leads every unified embedding model we measured on audio↔text. On a single
cross-modal protocol, this model exceeds ImageBind, LanguageBind, and Gemini
Embedding 2 on audio↔text in both directions, and both language-bound baselines on
emergent audio↔image (full tables below).
Unmodified base. Only the connector is trained; the base model's parameters are
byte-identical to the original release, so its text/image/video retrieval performance
(MMEB-V2) carries over unchanged.
Emergent cross-modal alignment. The connector is trained exclusively on audio–text
pairs. Audio→image retrieval nonetheless reaches R@10 0.407 over 696 VGGSound candidates
(chance: 0.014) with no audio-visual pairs in training — alignment to text places audio
in the space the base already shares across modalities.
Matryoshka representation. Embeddings truncate to {2048, 1536, 1024, 512, 256, 128,
64} dimensions with renormalization.
Compact distribution. This repository ships the connector and normalization
statistics (~60 MB); the frozen towers are downloaded from their original repositories.
The parameter count shown for this repository (16.4M) is the trained connector —
model.safetensors and the .pt checkpoint contain the same weights; inference.py
loads the .pt .
This is a research preview , currently at v0.3 : the v0.2 contrastive stage (484K
pairs) followed by a connector-only in-domain fine-tune on the AudioCaps train split.
Earlier versions remain downloadable via the v0.1-preview and v0.2-preview tags;
v0.3-preview pins the current version. All are compared below; pin a tag if you build
on this model.
A perceiver-resampler (width 384, 64 latent queries) translates frozen audio-tower frames
into the base model's input embedding space; its outputs occupy placeholder positions in
the input stream, mirroring the base model's image-token mechanism. Training is
contrastive (InfoNCE over the Matryoshka ladder, symmetric, with a full-corpus
frozen-text negative bank — 484K captions at v0.2) against the base model's text
embeddings in its native input format. v0.3 adds a second, connector-only fine-tuning
stage on the AudioCaps train split (400 steps at a reduced learning rate), warm-started
from the v0.2 checkpoint.
Input formatting. All inputs use the base model's chat-template format (instruction in
the system turn, content in the user turn, last-token pooling). Embedding quality is
sensitive to this formatting; use the templates in inference.py . For cross-modal
ranking, per-modality mean-centering of the gallery is recommended ( FusionEmbedder.center ).
Cross-modal retrieval — versus unified embedding models
VGGSound-AV, 696 audio/video-frame pairs (chance R@10 = 0.014). R@10 shown as
audio-side → other / other → audio-side:
ImageBind trains directly on audio–image pairs, so that pair is its supervised direction;
its audio–text alignment is emergent. LanguageBind trains audio against language (its
audio↔text is supervised; the value shown is its best readout, using the audio branch's
own text tower); its audio↔image is emergent. This model trains on audio–text only; its
audio–image alignment is emergent. All models evaluated with identical clips, frames, and
scoring, using the released imagebind_huge checkpoint and revision-pinned LanguageBind
checkpoints (LanguageBind_Audio_FT + LanguageBind_Image). Note on LanguageBind: its
branches fine-tune separate copies of the text tower, which diverge (mean caption cosine
0.55 between the audio and image branches' text embeddings) — the cross-branch binding
weakens, which is consistent with its emergent audio↔image score. This model's shared
space cannot drift by construction (the base is frozen; every training run asserts
parameter-level identity). Gemini Embedding 2 is Google's natively multimodal embedding
API (text/image/video/audio in one space), evaluated at its documented default invocation
(model id gemini-embedding-2 , 3072-d native output, inline audio+image+text,
google-genai 2.10.0) on the evaluation date shown; API models may change after that date.
One shared caveat: the evaluation captions are model-generated, which could favor models
whose text tower shares that caption style — all models received identical inputs.
Full audio→image metrics (per-modality mean-centered gallery — the readout implemented by
FusionEmbedder.center ; chance R@10 = 0.014):
The v0.3 in-domain fine-tune costs ~1 point of emergent audio→image alignment while
improving audio↔text (see the cross-modal table); v0.2 remains available if audio→image
is the primary use case.
What audio→image retrieval looks like. These numbers are not only aggregates — the
retrievals are organized by sound. Real examples (v0.2 checkpoint) on VGGSound-696
(query clip's frame left, top-5 retrieved images right; green = the clip's exact frame):
Example frames from the VGGSound dataset (CC-BY-4.0), shown for evaluation illustration.
Direct hits — the clip's own frame is returned in the top 5, among the same kind of scene:
Right neighbourhood — the exact frame ranks lower (often a poor still), but the top
results are the correct sound category:
On 10 tasks of the MTEB team's Massive Audio Embedding Benchmark
(mteb 2.18.0, v0.2 checkpoint; ranks vs the live leaderboard as of 2026-07-09, 21–65
models per task): UrbanSound8K T2A retrieval #3, Ravdess zero-shot #4, FSD2019Kaggle #6
(disclosed only — 13.6% of its test clips appear in the FSD50K dev split used in
training, verified by Freesound id, so it is withheld from the official submission),
BeijingOpera #6, with mid-field placements on speech/music tasks the model was never
trained for. Official leaderboard submission in progress.
Audio–text retrieval — versus specialist CLAP models
AudioCaps test — 883 clips, five reference captions per clip, recall computed as
min-rank over references:
CLAP-family models fine-tune both encoders end-to-end and include AudioCaps and Clotho
training data; this model keeps both towers frozen and trains only the connector.
Clotho v2.1 evaluation — 1,045 clips × 5 references, zero-shot (Clotho is excluded
from training data):
v0.3's in-domain AudioCaps stage trades 1.5 points of zero-shot Clotho A→T for the
AudioCaps gains above; T→A improves in both settings.
Text, image, and video benchmarks are the base model's published MMEB-V2 results, which
are unaffected by this extension.
# pip install git+https://github.com/Eximius-Labs/fusion-embedding-1 (+ transformers, torchvision, pillow)
from inference import FusionEmbedder
fe = FusionEmbedder.from_pretrained( "EximiusLabs/fusion-embedding-1-2b-preview" ,
device= "cuda" )
# or pin a version: revision="v0.3-preview" (current) / "v0.2-preview" / "v0.1-preview"
a = fe.embed_audio( "dog_barking.wav" ) # [2048]
t = fe.embed_text( "a dog barks while rain falls" ) # [2048]
i = fe.embed_image( "dog_photo.jpg" ) # [2048]
print ((a @ t), (a @ i), (t @ i)) # cosine similarities
a256 = fe.embed_audio( "dog_barking.wav" , dim= 256 ) # Matryoshka truncation
Training data and license
v0.2 was trained on ~484K audio–caption pairs: the full AudioCaps train split (45K),
FSD50K, WavCaps/AudioSet_SL, and a 318K-clip subset of LAION-FreeSound, using 10-second
training windows (random crop for longer clips). v0.3 continues the v0.2 checkpoint with
a 400-step fine-tune on the AudioCaps train split only. v0.1 used a 131K-pair subset of
the same sources. As this mix includes YouTube-sourced and research-licensed corpora, the preview
is released under CC-BY-NC-4.0 . Evaluation sets (AudioCaps test, Clotho, VGGSound,
ESC-50) are excluded from training by clip id.
Trained on sound-event data; speech content, speaker attributes, and music description
are supported by the instruction taxonomy but not yet trained to comparable quality.
English captions; 16 kHz mono input; 30 s per window (longer audio is chunked).
Audio–text retrieval is below fully fine-tuned CLAP-family models at this checkpoint
(see Evaluation).
Further corpus scaling, speech and music coverage, a commercially licensed release tier,
and the 8B model.
@software{fusion_embedding_2026,
title = {Fusion Embedding 1: A Unified Embedding Space for Text,
Image, Video, and Audio},
author = {Tonmoy, Abdul Basit},
year = {2026},
url = {https://github.com/Eximius-Labs/fusion-embedding-1}
}
Built on Qwen3-VL-Embedding and Qwen2.5-Omni, with training data from AudioCaps, WavCaps,
and FSD50K.
Safetensors Model size 16.4M params Tensor type F32 · Files info
Inference Providers NEW Feature Extraction This model isn't deployed by any Inference Provider. 🙋 1 Ask for provider support Model tree for EximiusLabs/fusion-embedding-1-2b-preview
