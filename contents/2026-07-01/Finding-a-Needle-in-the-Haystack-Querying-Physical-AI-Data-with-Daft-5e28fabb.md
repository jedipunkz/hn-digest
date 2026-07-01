---
source: "https://www.eventual.ai/blog/egodex-scenario-search"
hn_url: "https://news.ycombinator.com/item?id=48754121"
title: "Finding a Needle in the Haystack: Querying Physical AI Data with Daft"
article_title: "Finding a Needle in the Haystack: Querying Physical AI Data with Daft"
author: "sammysidhu"
captured_at: "2026-07-01T23:28:39Z"
capture_tool: "hn-digest"
hn_id: 48754121
score: 2
comments: 0
posted_at: "2026-07-01T22:45:06Z"
tags:
  - hacker-news
  - translated
---

# Finding a Needle in the Haystack: Querying Physical AI Data with Daft

- HN: [48754121](https://news.ycombinator.com/item?id=48754121)
- Source: [www.eventual.ai](https://www.eventual.ai/blog/egodex-scenario-search)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T22:45:06Z

## Translation

タイトル: 干し草の山から針を見つける: Daft を使用して物理 AI データをクエリする
説明: Daft を使用した Apple の EgoDex ハンド操作データセットに対するポーズ + セマンティック検索: SigLIP 埋め込みとハンドポーズ ジオメトリの出会い。物理 AI データの場合は Ctrl+F。

記事本文:
干し草の山から針を見つける: Daft を使用して物理 AI データをクエリする [メニュー] ブログ エンジニアリングに戻る 2026 年 6 月 26 日 干し草の山から針を見つける: Daft を使用して物理 AI データをクエリする
Daft を使用した Apple の EgoDex 手動操作データセットに対するポーズ + セマンティック検索: SigLIP 埋め込みと手のポーズ ジオメトリの出会い。物理 AI データの場合は Ctrl+F。
私たちは Apple の EgoDex 上で Daft を実行し、Daft によって何が可能になるかを示しました。フレームレベルのセマンティック埋め込みと幾何学的特徴を組み合わせた自然言語クエリを使用してビデオを検索することです。つまり、「筆記用具を持った手が箸を持ち上げるデータセット内のすべてのクリップを見つける」ということです。 Daft を使用すると、ロボット データセット上で CTRL+F を押すことができるようになりました。
歴史的に、ロボットは事前定義されたメニューにリストされたタスクを実行してきました。 Amazon Kiva はかつて、棚ポッドを保管場所からピッカーステーションに移動する方法を知るだけで済みました。ルンバは床の掃除方法を知るだけで済みます。これにより、データセットのキュレーションがある程度手動でキュレーションできるようになりました。研究者やエンジニアは、注意深く設定された実験を使用して、特定のタスク用のデータを簡単に生成できます。データセットの内容は構築によって簡単にわかりました。
管理されていない環境では、その想定は崩れます。衣類をたたむロボットをトレーニングする場合、事前に衣服の構成をすべて列挙することはできません。したがって、実験を計画する代わりに、人間にヘッドマウント カメラを装着し、折り畳んだり、手を伸ばしたり、掴んだりさせて、データの分布を完全に捉えるために必要なバリエーションが世界に提供されることを期待します。そのデータにはラベルが付けられていません。後でトレーニング ミックスを監査したり、失敗したケース (袖のもつれ、シルクのブラウス、裏返しの靴下など) について再トレーニングする必要がある場合、その瞬間がどこにあるのかわかりません。
それをさらにスケールアップします。500 台の自動運転車がペタバイトのビデオをアップロードします。

センサーデータを毎日クラウドに送信します。研究者として、再トレーニング用の 10 万時間以上のビデオの中から、事故に遭いそうになった箇所や交通違反をどのように特定しますか?
データはメニューから提供されなくなりました。データは到着したばかりで、誰も理解できないほどの速さで増え続けていますが、その中に何が含まれているかはあまり知られていないため、発見する必要があります。下流のタスクを微調整するために、ラベルのないマルチモーダル機能に基づいてデータの特定のサブセットを取得するにはどうすればよいでしょうか?データ内で困難なエッジケースや障害をどのように見つけられるでしょうか?
これは、最先端のロボット研究室のほとんどが直面しているデータ理解の問題です。
干し草の山から針を見つける: EgoDex
EgoDex は、さまざまな卓上タスクにわたる手のポーズの注釈とヘッドビュー ビデオのペアで構成される Apple の自己中心的なデータセットです。ユニークなシナリオと豊富なセンサーおよびビデオ データが豊富に含まれているため、マルチモーダルな理解を実現するのに最適です。
しかし、メニューの問題に遭遇します。たとえば、「座ったまま木製テーブルの上で小さな T シャツを折りたたむ」というタスクの説明では、いくつかの重要な幾何学的および視覚的プリミティブがわかりにくくなっています。
トレーニング ミックスにツイスト アクションが不足しているかどうかは、どうすればわかりますか?
人物がハンマー グリップを強く握って何かを持っているエピソードのサブセットを選択するにはどうすればよいですか?
Daft がどのようにソリューションを提供できるかを見てみましょう。
HDF5 ファイル形式では、ファイルごとに小さなファイルシステムのように、単一のファイルに多くの名前付き n 次元配列とメタデータが保持されます。 EgoDex および他の多くのロボット データセットは HDF5 形式でパッケージ化されているため、VideoFile、AudioFile などの既存のタイプと並行して、Daft でネイティブ サポートを備えた新しい Hdf5File タイプを作成することにしました。
ここから Apple から生の EgoDex データセットをダウンロードします。
mkdir -p .data
カール " https://ml-site.cdn-apple.com/datasets/egodex/test.zip " -o .data/te

st.zip
.data/test.zipを解凍します
すべての EgoDex エピソードは、対応する {i}.mp4 (ヘッドカメラ ビデオ) の隣にある 1 つの {i}.hdf5 (ハンドポーズ変換) です。 read_egodex というメソッドを作成して、そのデータセット全体を Daft の単一のフレームごとの DataFrame に変換します。
egodex_lib.egodex から read_egodex をインポート
# Raw EgoDex HDF5 → フレームごとに 1 行、Daft に直接入力。
df = read_egodex( " .data/**/*.hdf5 " , with_video = True )
HDF5 から Daft への実装の詳細
daft.from_files では、HDF5 ファイルごとに 1 行が生成されます。単一の @daft.func UDF は、ネイティブ Hdf5File API を通じて各ファイルを開き、必要なすべての変換を 1 回の呼び出しでバッチ読み取りし、NumPy 特徴量計算 (build_state、build_skeleton、build_extrinsics) を実行して、フレームごとの構造体のリストとしてエピソードを返します。次に、爆発を実行し、ブロードキャストされたタスクを使用してフレームごとに 1 行にリストをファンします。
@daft.func ( return_dtype = EPISODE_DTYPE )
def process_egodex_episode ( file_ : daft.File) -> dict :
h = file_.as_hdf5()
変換 = h.read( list ( dict .fromkeys( STATE_TRANSFORMS + SKELETON_TRANSFORMS + [ CAMERA ])))
state = build_state(変換)
スケルトン = build_skeleton(変換)
extrinsics = build_extrinsics(transforms)
アクション = next_frame_action(状態)
task =solve_task(h.attrs()) # ネイティブ attrs() は辞書を返します
フレーム = [
{
" フレームインデックス " : i、
" 観察.状態 " : 状態[i],
" 観察.スケルトン " : スケルトン[i],
" 観察.extrinsics " : extrinsics[i],
" アクション " : アクション[i]、
}
範囲内の i の場合 ( len (状態))
]
return { " タスク " : タスク、 " フレーム " : フレーム}
def read_egodex ( hdf5_glob 、 with_video : bool = False ):
per_file = Window().order_by(col( " file " ).file_path())
エピソード = (
daft.from_files(hdf5_glob)
.sort(col( " ファイル " ).file_path())
.with_column( " エピソードインデックス " , row_number().over(ファイルごと) - 1 )
# を運ぶ

e HDF5 パス。ビデオ デコーダーが各エピソードの兄弟 .mp4 を見つけられるようにします。
.with_column( " _src " ,col( " file " ).file_path())
.into_batches( 8 )
.with_column( " _ep " , process_egodex_episode(col( " file " )))
.with_column( " タスク " ,col( " _ep " )[ " タスク " ])
.with_column( " フレーム " ,col( " _ep " )[ " フレーム " ])
)
フレーム = (
エピソード.explode( " フレーム " )
.select( " エピソードインデックス " 、 " タスク " 、 " _src " 、col( " フレーム " ).unnest())
.with_column( " タイムスタンプ " , (col( " フレームインデックス " ) / FPS ).cast(DataType.float32()))
)
ビデオ付きの場合:
フレーム = Frames.with_column( " 観察.画像 " , _decode_sibling_mp4(col( " _src " ),col( " タイムスタンプ " )))
フレームを返す.exclude( " _src " )
最初の 3 行を読み取った後のデータフレーム出力は次のようになります。
以下にいくつかのビデオの例を示します (スキーマの要約)。
各フレームの内容: SigLIP 埋め込み
ここで、Daft を使用して、Google の SigLIP-2 画像エンコーダをサブサンプリングされた一連のフレーム (1 fps) に対してエピソード全体で実行し、結果を Daft DataFrame の Vector_column として保存します。
daft import DataType、Series から
トランスフォーマーから AutoModel、AutoProcessor をインポート
輸入トーチ
def _auto_device () -> str :
_HAS_CUDA の場合:
「cuda」を返します
torch.backends.mps.is_available() の場合:
「mps」を返します
「CPU」を返します
_HAS_CUDA = torch.cuda.is_available()
_HAS_CUDA の場合は GPUS = 1、それ以外の場合は 0
DEVICE = os.environ.get( " CLIP_DEVICE " , _auto_device())
DTYPE = torch.float16 if DEVICE == " cuda " else torch.float32
SUBSAMPLE = 30 # 30 フレームごとに 1 フレームを保持します (~1 fps)。セマンティックコンテンツは隣接するフレーム間でほとんど変化しません
MODEL_ID = " google/siglip2-base-patch16-224 "
EMB_DIM = 768 # SigLIP2 ベースの共有画像/テキスト埋め込み寸法 (モデルと一致する必要があります)
エンコーダーは、ステートフル UDF である @daft.cls でラップされます。プラン UDF とは異なり、クラス UDF インスタ

ワーカーごとに 1 回初期化され、GPU メモリに留まり、すべてのバッチで再利用されます。
def _normalized_embedding ( model_output ) -> torch.Tensor:
"""トランスフォーマーの出力から埋め込みテンソルを取り出し、それを L2 正規化します。
トランスフォーマー 5.x は get_image_features / からモデル出力オブジェクトを返します。
get_text_features;古いバージョンは裸のテンソルを返しました。両方を扱います。
「」
torch.is_tensor(model_output) の場合:
特技 = モデル出力
それ以外の場合:
特技 = モデル出力.プーラー出力
偉業 = 偉業.float()
return feats / feats.norm( dim =- 1 , keepdim = True ) # 単位ノルム so cosine == ドット積
@daft.cls ( gpus = GPUS 、 max_concurrency = 1 、 use_process = False )
クラス SiglipEmbedder :
def __init__ ( self ) -> なし :
self .model = AutoModel.from_pretrained( MODEL_ID , torch_dtype = DTYPE ).to( DEVICE ).eval()
self .processor = AutoProcessor.from_pretrained( MODEL_ID )
@daft.method.batch ( return_dtype = DataType.embedding(DataType.float32(), EMB_DIM )、batch_size = 16 )
def embed_image ( self 、画像 : シリーズ):
# image.to_pylist() は uint8 H×W×C numpy 配列を生成します。 SigLIPプロセッサがそれらを受け取ります
# 直接 (PIL パスと同一であることが確認された) なので、フレームごとの Image.fromarray は必要ありません。
inputs = self .processor(images =images.to_pylist(),return_tensors = " pt " ).to( DEVICE )
torch.no_grad() を使用:
model_output = self .model.get_image_features( ** 入力)
embeddings = _normalized_embedding(model_output)
戻りリスト (embeddings.cpu().numpy())
egodex_lib から egodex をインポート
ダフインポートコルから
emb = (
egodex.embed_frames(
df.where(col( " エピソード_インデックス " ).is_in( EPISODES ))
)
.select( " エピソードインデックス " 、 " フレームインデックス " 、 " クリップ_emb " )
.collect()
)
emb.show( 3 )
次に、Daft はエンコーダーを介してフレームをバッチでストリーミングし、各 768D ベクトルを DataFrame の列として書き戻します。
埋め込みます

一度実行すれば、クエリ時に再利用できます。クエリの後で、テキスト クエリ (「箸」、「折りたたんだシャツ」) が同じ SigLip モデルによってエンコードされ、clip_emb に対するコサイン類似度が DataFrame の類似度列になります。
手は何をしているのか: 幾何学的特徴
SigLIP では、フレームに箸が入っていることがわかります。それらを持っている手が筆記用グリップを握っているとは言えません。これは、センサー データ内の 48D 手首のポーズと 204D 関節スケルトンから直接計算できる幾何学的事実です。
私たちは、EgoDex における幾何学的シナリオの抽象化を「状態とアクション」として提案します。
状態 = 1 フレームのみで計算できる手/手首/腕のプロパティ。
私たちは手の外科的研究を利用して手のポーズを研究しました。その結果、この分野は指先と親指で物体を繊細につまむ「精密グリップ」と、手全体で物体を包み込む「パワーグリップ」の 2 つのグループに分かれることがわかりました。また、手の屈曲モデルを使用して、手の開き具合を状態として分類します。
アクション = 複数のフレームにわたって計算する必要がある手/手首/腕のプロパティ。
たとえば、物を持ち上げる動作は、時間の経過とともに手首の Y 位置が通常よりも早く増加するのが特徴です。これはほとんどのアクションに一般化され、アクションの開始時と停止時だけでなく、アクションの発生時を検出するために変化率などの何らかのメトリクスを計算します。
これは、どのポーズを検出し、どのように検出するかをまとめたものです。
各フレームのセンサー データは、ビデオだけでは得られない幾何学的量をエンコードしています。つまり、observation.state (48 個の数値、手首 + 両手の指先 5 個) および Observation.skeleton (204 個の数値、68 関節) です。
静的なポーズと進行中のアクションの検出をサポートするために、これらの幾何学的量を使用してフレームごとの状態と時間の経過に伴うアクション率を計算し、それらを

SigLIP 埋め込みの横にある列。
前腕ロールの計算 - daft udf
@daft.func (return_dtype = DataType.float64())
def forearm_roll ( rot6d 、 rot6d_next 、 forearm_axis ):
"""あるフレームから次のフレームまで、前腕の軸を中心としたリスト ロール (rad) (エピソードの最後のフレームでは 0)。""
rot6d が None または rot6d_next が None の場合:
0.0を返す
デルタ = _rotation_matrix(rot6d_next) @ _rotation_matrix(rot6d).T
角度 = np.arccos(np.clip((np.trace(delta) - 1 ) / 2 , - 1 , 1 ))
axis = np.array([delta[ 2 , 1 ] - delta[ 1 , 2 ], delta[ 0 , 2 ] - delta[ 2 , 0 ], delta[ 1 , 0 ] - delta[ 0 , 1 ]])
大きさ = np.linalg.norm(axis)
マグニチュード < 1e-9 の場合:
0.0を返す
return float ( abs (角度 * np.dot(axis / 大きさ, np.asarray(forearm_axis))))
6 桁の回転からの方向 — 手の完全な方向 (および手のひらに向かう方向) を rot6d から Gram-Schmidt 経由で再構築します。
デフォルトの回転から_rot6d ( rot6d ):
"""rot6d (N, 6) -> (N, 3, 3) 回転行列 (列 = 手の x、y 軸 + 手のひらの法線)。"""
first_column = rot6d[: , 0 : 3 ]
first_column = first_column / (np.linalg.norm(first_column, axis = 1 , keepdims = True ) + 1e-9 )
Second_column = rot6d[ : , 3 : 6 ]
2 番目の列 = 2 番目の列 - (最初の列 * 2 番目の列).sum( 1 , keepdims = True )

[切り捨てられた]

## Original Extract

Pose + semantic search over Apple's EgoDex hand-manipulation dataset with Daft: SigLIP embeddings meet hand-pose geometry. Ctrl+F for physical AI data.

Finding a Needle in the Haystack: Querying Physical AI Data with Daft [menu] Back to Blog Engineering June 26, 2026 Finding a Needle in the Haystack: Querying Physical AI Data with Daft
Pose + semantic search over Apple's EgoDex hand-manipulation dataset with Daft: SigLIP embeddings meet hand-pose geometry. Ctrl+F for physical AI data.
We ran Daft on Apple's EgoDex and showed what Daft makes possible: searching on video using natural language queries combining frame-level semantic embeddings and geometric features, i.e., “find every clip in my dataset where a writing-gripped hand lifts chopsticks.” Using Daft, you can now CTRL+F over your robotics dataset.
Historically, robots have performed tasks listed on a pre-defined menu. An Amazon Kiva used to only need to know how to move shelving pods from storage to picker stations. Your Roomba only needs to know how to clean your floor. This made dataset curation somewhat hand-curatable. Using carefully set up experiments, a researcher or engineer could easily generate data for a specific task. Contents of the dataset were easily known by construction.
In uncontrolled environments, that assumption falls apart. When training a clothes-folding robot, you can’t enumerate every garment configuration in advance. So instead of designing experiments, you outfit humans with head-mounted cameras and let them fold, reach, and grasp, hoping that the world provides the variation needed to capture the full distribution of data. That data doesn’t come labeled. When you need to later audit your training mix or retrain on failure cases (like a tangled sleeve, silk blouse, inside-out sock), you have no idea where those moments are.
Scale that up even further: A fleet of 500 autonomous vehicles uploads petabytes of video and sensor data to the cloud every day. How do you, as a researcher, identify your near crashes and traffic violations amongst hundreds of 100k+ hours of video for re-training?
Your data is no longer served off a menu. The data just arrives, continuously growing faster than anyone could understand it, and what’s in it isn’t well-known: it needs to be discovered. How do you retrieve a specific subset of data based on unlabeled multimodal features for downstream task fine-tuning? How can you find difficult edge cases and failures within your data?
This is the data understanding problem most frontier robotics labs are facing.
Finding a Needle in the Haystack: EgoDex
EgoDex is Apple’s egocentric dataset consisting of paired hand pose annotations and head-view video across varying tabletop tasks. It’s full of unique scenarios and rich sensor and video data, making it a perfect candidate for multimodal understanding.
But it runs into the menu problem. For example, the task description “fold a small t-shirt on a wooden table while sitting” obscures some important geometric and visual primitives.
How can I tell if my training mix is short on twisting actions?
How can I select a subset of episodes where the person holds something with a tight hammer grip?
Let's see how Daft can provide the solution.
In the HDF5 file format, a single file holds many named n-dimensional arrays plus metadata, like a tiny filesystem per file. Since EgoDex as well as many other robotics datasets come packaged in the HDF5 format, we decided to write a new Hdf5File type with native support in Daft, alongside existing types like VideoFile, AudioFile, etc.
Download the raw EgoDex dataset from Apple here:
mkdir -p .data
curl " https://ml-site.cdn-apple.com/datasets/egodex/test.zip " -o .data/test.zip
unzip .data/test.zip
Every EgoDex episode is one {i}.hdf5 (the hand-pose transforms) sitting next to a corresponding {i}.mp4 (the head-cam video). We write a method, read_egodex turns that whole dataset into a single per-frame DataFrame in Daft:
from egodex_lib.egodex import read_egodex
# Raw EgoDex HDF5 → one row per frame, straight into Daft.
df = read_egodex( " .data/**/*.hdf5 " , with_video = True )
HDF5 to Daft Implementation Details
daft.from_files yields one row per HDF5 file. A single @daft.func UDF opens each file through the native Hdf5File API, batch-reads every transform it needs in one call, runs NumPy feature math (build_state, build_skeleton, build_extrinsics), and returns the episode as a list of per-frame structs. We then run an explode then fans that list into one row per frame with the task broadcasted.
@daft.func ( return_dtype = EPISODE_DTYPE )
def process_egodex_episode ( file_ : daft.File) -> dict :
h = file_.as_hdf5()
transforms = h.read( list ( dict .fromkeys( STATE_TRANSFORMS + SKELETON_TRANSFORMS + [ CAMERA ])))
state = build_state(transforms)
skeleton = build_skeleton(transforms)
extrinsics = build_extrinsics(transforms)
action = next_frame_action(state)
task = resolve_task(h.attrs()) # native attrs() returns a dict
frames = [
{
" frame_index " : i,
" observation.state " : state[i],
" observation.skeleton " : skeleton[i],
" observation.extrinsics " : extrinsics[i],
" action " : action[i],
}
for i in range ( len (state))
]
return { " task " : task, " frames " : frames}
def read_egodex ( hdf5_glob , with_video : bool = False ):
per_file = Window().order_by(col( " file " ).file_path())
episodes = (
daft.from_files(hdf5_glob)
.sort(col( " file " ).file_path())
.with_column( " episode_index " , row_number().over(per_file) - 1 )
# carry the HDF5 path so the video decoder can find each episode's sibling .mp4
.with_column( " _src " , col( " file " ).file_path())
.into_batches( 8 )
.with_column( " _ep " , process_egodex_episode(col( " file " )))
.with_column( " task " , col( " _ep " )[ " task " ])
.with_column( " frames " , col( " _ep " )[ " frames " ])
)
frames = (
episodes.explode( " frames " )
.select( " episode_index " , " task " , " _src " , col( " frames " ).unnest())
.with_column( " timestamp " , (col( " frame_index " ) / FPS ).cast(DataType.float32()))
)
if with_video:
frames = frames.with_column( " observation.image " , _decode_sibling_mp4(col( " _src " ), col( " timestamp " )))
return frames.exclude( " _src " )
Here's what your dataframe output should look like after reading the first 3 rows:
Here are some example videos (abridged schemas):
What's in Each Frame: SigLIP Embeddings
Now, using Daft, we'll run Google’s SigLIP-2 image encoder over a subsampled set of frames (1 fps), across episodes, and store the result as a vector_column in the Daft DataFrame.
from daft import DataType, Series
from transformers import AutoModel, AutoProcessor
import torch
def _auto_device () -> str :
if _HAS_CUDA :
return " cuda "
if torch.backends.mps.is_available():
return " mps "
return " cpu "
_HAS_CUDA = torch.cuda.is_available()
GPUS = 1 if _HAS_CUDA else 0
DEVICE = os.environ.get( " CLIP_DEVICE " , _auto_device())
DTYPE = torch.float16 if DEVICE == " cuda " else torch.float32
SUBSAMPLE = 30 # keep 1 of every 30 frames (~1 fps); semantic content barely changes between adjacent frames
MODEL_ID = " google/siglip2-base-patch16-224 "
EMB_DIM = 768 # SigLIP2-base shared image/text embedding dim (must match the model)
The encoder is wrapped in a @daft.cls, which is a stateful UDF. Unlike plan UDFs, a class UDF instantiates once per worker and stays in GPU memory to be reused across all batches:
def _normalized_embedding ( model_output ) -> torch.Tensor:
"""Pull the embedding tensor out of a transformers output and L2-normalize it.
transformers 5.x returns a model-output object from get_image_features /
get_text_features; older versions returned a bare tensor. Handle both.
"""
if torch.is_tensor(model_output):
feats = model_output
else :
feats = model_output.pooler_output
feats = feats.float()
return feats / feats.norm( dim =- 1 , keepdim = True ) # unit-norm so cosine == dot product
@daft.cls ( gpus = GPUS , max_concurrency = 1 , use_process = False )
class SiglipEmbedder :
def __init__ ( self ) -> None :
self .model = AutoModel.from_pretrained( MODEL_ID , torch_dtype = DTYPE ).to( DEVICE ).eval()
self .processor = AutoProcessor.from_pretrained( MODEL_ID )
@daft.method.batch ( return_dtype = DataType.embedding(DataType.float32(), EMB_DIM ), batch_size = 16 )
def embed_image ( self , images : Series):
# images.to_pylist() yields uint8 H×W×C numpy arrays; the SigLIP processor takes them
# directly (verified identical to the PIL path), so no per-frame Image.fromarray needed.
inputs = self .processor( images = images.to_pylist(), return_tensors = " pt " ).to( DEVICE )
with torch.no_grad():
model_output = self .model.get_image_features( ** inputs)
embeddings = _normalized_embedding(model_output)
return list (embeddings.cpu().numpy())
from egodex_lib import egodex
from daft import col
emb = (
egodex.embed_frames(
df.where(col( " episode_index " ).is_in( EPISODES ))
)
.select( " episode_index " , " frame_index " , " clip_emb " )
.collect()
)
emb.show( 3 )
Then, Daft streams the frames through the encoder in batches and writes each 768D vector back as a column in the DataFrame.
We embed once and reuse at query time. Later on query, a text query (”chopsticks”, “folded shirt”) will be encoded by the same SigLip model, and cosine similarity against clip_emb becomes a similarity column in the DataFrame.
What the Hands Are Doing: Geometric Features
SigLIP can tell you a frame contains chopsticks. It can’t tell you the hand holding them is in a writing grip. That’s a geometric fact that can be computed directly from the 48D wrist pose and the 204D joint skeleton in the sensor data.
We propose an abstraction for geometric scenarios in EgoDex as "states and actions".
States = a property of the hands/wrist/arms that can be computed over just one frame.
We researched hand poses using hand-surgery research, and what we discovered is that the field splits into two families: precision grips, where the fingertips and thumb delicately pinch an object, and power grips, where the whole hand wraps around it. We also classify hand openness as a state using the hand flexion model.
Actions = a property of the hands/wrist/arms that must be computed over several frames.
Lifting, for example, is characterized by the Y-position of the wrist increasing quicker than usual over time. This generalizes to most actions, where we compute some kind of metric like a rate of change to detect when an action occurs, as well as when it starts and stops.
This is a summary of which poses we detect and how we do it:
Each frame's sensor data encodes geometric quantities the video alone doesn't: observation.state (48 numbers, wrist + 5 fingertips per hand) and observation.skeleton (204 numbers, 68 joints).
To support detecting static poses as well as in-progress actions, we use these geometric quantities to compute states per frame as well as action rates over time and write them as columns alongside the SigLIP embeddings.
Computing forearm roll - daft udf
@daft.func ( return_dtype = DataType.float64())
def forearm_roll ( rot6d , rot6d_next , forearm_axis ):
"""Wrist roll (rad) about the forearm axis from one frame to the next (0 at an episode's last frame)."""
if rot6d is None or rot6d_next is None :
return 0.0
delta = _rotation_matrix(rot6d_next) @ _rotation_matrix(rot6d).T
angle = np.arccos(np.clip((np.trace(delta) - 1 ) / 2 , - 1 , 1 ))
axis = np.array([delta[ 2 , 1 ] - delta[ 1 , 2 ], delta[ 0 , 2 ] - delta[ 2 , 0 ], delta[ 1 , 0 ] - delta[ 0 , 1 ]])
magnitude = np.linalg.norm(axis)
if magnitude < 1e-9 :
return 0.0
return float ( abs (angle * np.dot(axis / magnitude, np.asarray(forearm_axis))))
Orientation from the 6-number rotation — rebuild a hand's full orientation (and the palm-facing direction) from rot6d via Gram–Schmidt:
def rotation_from_rot6d ( rot6d ):
"""rot6d (N, 6) -> (N, 3, 3) rotation matrices (columns = hand x, y axes + palm normal)."""
first_column = rot6d[ : , 0 : 3 ]
first_column = first_column / (np.linalg.norm(first_column, axis = 1 , keepdims = True ) + 1e-9 )
second_column = rot6d[ : , 3 : 6 ]
second_column = second_column - (first_column * second_column).sum( 1 , keepdims = True )

[truncated]
