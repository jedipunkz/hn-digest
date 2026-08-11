---
source: "https://www.docuseal.com/blog/gpu-ai-workloads-with-a-ruby-on-rails-monolith"
hn_url: "https://news.ycombinator.com/item?id=49256977"
title: "GPU AI workloads with a Ruby on Rails monolith"
article_title: "Running GPU AI workloads with a Ruby on Rails monolith | DocuSeal"
author: "babanooey21"
captured_at: "2026-08-11T12:41:48Z"
capture_tool: "hn-digest"
hn_id: 49256977
score: 2
comments: 1
posted_at: "2026-08-11T12:08:58Z"
tags:
  - hacker-news
  - translated
---

# GPU AI workloads with a Ruby on Rails monolith

- HN: [49256977](https://news.ycombinator.com/item?id=49256977)
- Source: [www.docuseal.com](https://www.docuseal.com/blog/gpu-ai-workloads-with-a-ruby-on-rails-monolith)
- Score: 2
- Comments: 1
- Posted: 2026-08-11T12:08:58Z

## Translation

タイトル: Ruby on Rails モノリスを使用した GPU AI ワークロード
記事のタイトル: Ruby on Rails モノリスを使用した GPU AI ワークロードの実行 |ドキュシール
説明: DocuSeal が、リクエスト パスに Python を使用せずに、TensorRT と Ruby を使用した Rails モノリス内で PDF フォーム フィールド検出用の GPU オブジェクト検出モデルを実行する方法。

記事本文:
ドキュシール
18k 18k
ソリューション
ソリューション
ドキュメント署名 API
自動化とワークフローの構築
埋め込み署名
署名エクスペリエンスをシームレスに統合
React App にドキュメント署名を埋め込む
エンタープライズ
当社のエンタープライズ ソリューションを探索する
オンプレミス
サーバー上の自己ホスト型 DocuSeal
支払いを受け入れる
署名中に代金を回収する
SaaS 向けのドキュメント署名
ドキュメント署名を製品に追加する
すべてのソリューション
他のソリューションを検討する
ソフトウェアとテクノロジー
ヘルスケア
不動産
教育機関
金融サービス
リソース
リソース
コンプライアンス
当社が電子サイン標準にどのように準拠しているかをご覧ください
APIリファレンス
API エンドポイントと Webhook を探索する
ブログ
最新の記事を入手してください
ドキュメント
答えを見つけて機能について学ぶ
開発者ガイド
開発ガイドと例
ハウツー
ユーザーのチュートリアルとリソース
GDPR
GDPR 準拠の EU クラウドを探索する
適格な電子署名
最高レベルの信頼で署名する
サインイン
始めましょう
ドキュシール
ソリューション
ドキュメント署名 API
自動化とワークフローの構築
埋め込み署名
ドキュメント署名をシームレスに統合
エンタープライズ
当社のエンタープライズ ソリューションを探索する
オンプレミス
サーバー上の自己ホスト型 DocuSeal
React ドキュメント署名
React App にドキュメント署名を埋め込む
SaaS 向けのドキュメント署名
ドキュメント署名を製品に追加する
支払いを受け入れる
署名中に代金を回収する
すべてのソリューション
他のソリューションを検討する
リソース
コンプライアンス
当社が電子サイン標準にどのように準拠しているかをご覧ください
ブログ
最新の記事を入手してください
APIリファレンス
API エンドポイントと Webhook を探索する
ドキュメント
答えを見つけて機能について学ぶ
開発者ガイド
開発ガイドと例
ハウツー
ユーザーのチュートリアルとリソース
GDPR
GDPR 準拠の EU クラウドを探索する
適格な電子署名
最高レベルの署名

信頼
知識ベースの認証
知識ベースの質問で本人確認を行う
サインイン
AIに聞く
すべての投稿
Ruby on Rails モノリスを使用した GPU AI ワークロードの実行
DocuSeal はオープンソースの電子署名プラットフォームであり、使用を開始するには、ユーザーは文書をアップロードし、入力または署名するフィールドをマップする必要があります。複雑なフォームに数十、さらには数百のフィールドを手動で追加するのは面倒な場合があるため、公開されているさまざまな PDF フォームのフィールドを検出するようにトレーニングされたコンピューター ビジョン モデルに基づく AI フィールド検出機能を実装しました。 AI ワークロードは、Python やマイクロサービスを使用せず、Rails モノリスの Ruby プロセス内の NVIDIA GPU インスタンス上で実行されます。
私たちは、Ruby を使用して Rails モノリス アプリを構築する利便性を気に入っています。また、AI 機能の開発と保守を容易にするために、同じ Ruby on Rails モノリス アプリ内に AI フィールド検出を搭載したいと考えていました。これを実現するために、Ruby と NVIDIA T4 GPU インスタンス上で実行される Sidekiq 非同期ジョブ プロセッサを使用して AI フィールド検出推論パイプラインを構築しました。以下のスクリーンショットは、フィールド検出中のプロダクション GPU ワーカー上の Ruby Sidekiq プロセスによる nvtop NVIDIA GPU 使用率を示しています。
コンピューター ビジョン モデルを GPU で効率的に実行するために、NVIDIA は TensorRT 推論ランタイムを提供しています。 TensorRT は C++ API を公開しますが、それに対応する Ruby バインディングは存在しませんでした。そのため、フォワード パスに必要なメソッド (エンジンのロード、テンソルの検査、デバイス メモリのバインド、実行、CUDA ストリームの同期) のみをリンクする非常に小さな単一ファイルの Rice C++ バインディングを構築しました。これらの TensorRT Ruby バインディングは、Apache 2.0 ライセンスに基づいてオープン ソースとして作成され、GitHub で入手できます。
コピー
コピーされました
gem インストール tensorrt
gem をビルドしてインストールするには、システム上に TensorRT と NVIDIA CUDA が必要です。全体

TensorRT バインディングは、11 のメソッドを持つ単一の TensorRT::Engine クラスです。
コピー
コピーされました
「tensorrt」が必要です
エンジン = TensorRT :: エンジン 。 new (model_path 、冗長: false )
エンジン。 num_io_tensors # 入力/出力テンソルの数
エンジン。 get_tensor_name (index) # インデックスによるテンソル名
エンジン。 _入力ですか? ( name ) # テンソルが入力されているかどうかを確認する
エンジン。 get_tensor_shape ( name ) # 配列として整形 [1, 3, 640, 640]
エンジン。 get_tensor_bytes ( name ) # バイト単位のサイズ
エンジン。 get_tensor_dtype ( name ) # データ型、例: float32
エンジン。 set_tensor_address ( name , device_ptr ) # GPU メモリをバインドする
エンジン。実行 # 同期実行
エンジン。 enqueue # 非同期実行
エンジン。 get_stream # CUDAストリームハンドル
エンジン。 stream_synchronize # ストリームの完了を待ちます
推論パイプライン
パイプラインは 3 つのステージで構成されます。
前処理中。 PDF ページをレンダリングし、入力テンソルを準備します。
フォワードパス。 TensorRT を使用して GPU 上でモデルを実行します。
後処理。出力テンソルをページ上のフィールド座標に変換します。
1. 前処理: PDF ページから入力テンソルまで
まず、PDF ページは PDFium を使用してレンダリングされます。レンダリングされたイメージは、アスペクト比を維持したままモデル入力解像度にスケーリングされ、正方形にパディングされ、標準の ImageNet 平均と標準偏差で正規化され、HWC レイアウトから CHW レイアウトに転置されます。画像操作は、libvips バインディングである Ruby-vips を使用して実行され、テンソル操作には、NumPy の代替となる Ruby である Numo を使用します。
コピー
コピーされました
平均 = [ 0.485 , 0.456 , 0.406 ]。フリーズする
STD = [ 0.229 、 0.224 、 0.225 ]。フリーズする
スケール = [解像度。 to_f / 画像 .幅、解像度。 to_f / 画像 .身長 ]。分
リサイズ＝画像。サイズ変更 (スケール、vscale:スケール、カーネル::lanczos3)
Pad_x = (( 解像度 - (画像 . 幅 * スケール ). 丸め ) / 2.0 )。丸い
パッドy = ((

解像度には（画像の高さ・スケール）。丸め）/2.0）。丸い
画像 = リサイズされました。埋め込み (pad_x 、pad_y 、解像度、解像度、背景: [ 255 , 255 , 255 ])
# ImageNet の正規化
画像 /= 255.0
画像 = (画像 - 平均) / STD
img_array = Numo :: SFloat 。 from_binary (画像 .write_to_memory , [解像度 , 解像度 , 3 ])
input_tensor = img_array 。転置 ( 2 , 0 , 1 )。 reshape ( 1 , 3 , 解像度 , 解像度 )
2. フォワードパス: GPU 上でモデルを実行する
フォワード パスを実行するには、入力テンソルがエンジンが宣言したデータ型にキャストされ、バイナリ文字列にシリアル化され、ホスト バッファに書き込まれ、GPU VRAM にコピーされます。実行は enqueue によって開始され、作業を送信してすぐに戻ります。次に、retrieve を使用して、フォワード パスが完了するのを待ち、出力を読み取ります。
コピー
コピーされました
def エンキュー ( input_tensor )
host_ptr = FFI::MemoryPointer 。新しい ( :uint8 、 @input_size )
host_ptr 。 write_bytes ( Numo :: SFloat . Cast ( input_tensor ) . to_binary )
TensorRT :: CUDA 。 memcpy_htod_async ( @input_ptr 、 host_ptr 、 @input_size 、 @cuda_stream )
@エンジン 。エンキュー # ノンブロッキング
終わり
デフ取得
@エンジン 。ストリーム同期
{ dets: read_output ( @dets_ptr 、 @dets_size )、
ラベル: read_output ( @labels_ptr 、 @labels_size ) }
終わり
def read_output ( device_ptr , size )
host_ptr = FFI::MemoryPointer 。新しい ( :uint8 、サイズ )
TensorRT :: CUDA 。 memcpy_dtoh ( host_ptr 、 device_ptr 、 サイズ )
Numo :: SFloat 。 from_binary ( host_ptr . read_bytes ( size ))
終わり
運用環境では、これらのホスト バッファはスレッドごとに 1 回割り当てられ、ジョブ全体で再利用されるため、定常状態推論では Ruby 側で割り当ては実行されません。各 Sidekiq スレッドは独自のエンジン インスタンスを保持します。
3. 後処理: 出力テンソルからフォームフィールドまで
エンジンはボックス予測と p を返します。

er クラスのロジッツ。後処理では、シグモイドを適用してロジットをスコアに変換し、各ボックスの最高スコア クラスを取得し、ボックスを中心から隅の形式に変換し、前処理中に適用されたスケールとパディングを反転し、信頼性のしきい値を下回る検出を破棄します。次に、非最大抑制により、同じ領域上の重複ボックスが削除され、座標はフォーム ビルダーが使用する 0..1 の範囲に正規化されます。
生き残った各検出は、相対ページ座標とフィールド タイプを持つ Field オブジェクトになります。
コピー
コピーされました
フィールド。新しい (タイプ: 'text' 、 x: 0.6123 、 y: 0.8402 、 w: 0.2510 、 h: 0.0338 、信頼度: 0.94 )
非同期パイプライン処理
ステージ 1 と 3 (前処理と後処理) は CPU で実行され、ステージ 2 (フォワード パス) は GPU で実行されるため、CPU は GPU フォワード パスを待機してアイドル状態になり、GPU は CPU の前処理と後処理を待機してアイドル状態になります。スループットを向上させるために、GPU が現在の PDF ページで順方向パスを実行している間に、CPU が次の PDF ページを前処理する非同期パイプラインを構築しました。
これを実現するために、TensorRT 非同期実行を利用します。この場合、enqueue は CUDA ストリームに作業を送信してブロックせずに戻り、stream_synchronize はストリームが完了するまでブロックします。 2 つの呼び出しを分離し、ラムダとして取得を返すことで、延期が明示的に行われます。
コピー
コピーされました
def エンキュー (入力:, **)
推論。エンキュー ( * 入力 )
-> { 推論 .取得する }
終わり
ページ ループはこれを使用してステージをオーバーラップします。ページ N の順方向パスがキューに入れられ、GPU の実行中にページ N+1 が CPU 上でレンダリングおよび前処理され、ページ N の結果は次のページの入力が準備された後にのみ読み取られます。
コピー
コピーされました
image = prepare_page_image ( doc . get_page ( page_indexes . f

最初に))
current_args = 推論 。 prepare_input ( image , ** prep_opts )
current_task = 推論 。エンキュー ( ** current_args 、 ** infer_opts )
page_indexes 。 each_with_index を実行します。現在のページ番号 , i |
next_n = ページインデックス [ i + 1 ]
if next_n # GPU が現在のページを実行している間に、CPU が次のページを前処理します
next_image = prepare_page_image ( doc . get_page ( next_n ))
next_args = 推論 。 prepare_input ( next_image , ** prep_opts )
終わり
出力 = current_task 。 call # GPU がフォワード パスを完了して出力を読み取るまで待機します
next_task = 推論 。 enqueue ( ** next_args , ** infer_opts ) if next_args # 次のページの前方パスを開始します (ノンブロッキング)
フィールド = 推論 。 process_outputs (outputs , ** current_args , ** infer_opts ) # 出力テンソルをフィールド座標に変換します
yield [attachment_uuid , current_page_number , field ] # このページのフィールドを公開する
current_args = next_args
現在のタスク = 次のタスク
終わり
ドキュメントの実行中、CPU と GPU の作業は重複します。同期的に実行される同じパイプラインに対して測定すると、約 80% 高いスループットが得られます。
結果をブラウザにストリーミングする
フィールド検出は GPU インスタンス上の Sidekiq ワーカーで実行され、結果を Server-Sent Events 経由でユーザーのブラウザにストリーミングして返すために、Redis pub/sub を使用してワーカー プロセスから Web プロセスにメッセージを運びます。 WebSocket で Rails Action Cable を使用することも実行可能なオプションですが、私たちはプレーンな Rails Live SSE コントローラーを使用することにしました。
ワーカーは、ページが完了するとすぐに各ページのフィールドを公開し、処理されたページ数のライブ進行状況インジケーターをユーザーに表示できるようにします。 DetectFields.call に渡されるブロックは、上記のパイプライン ループからページごとに 1 回呼び出されます。
コピー
コピーされました
クラステンプレートDe

tectFieldsジョブ
Sidekiq :: ジョブを含める
def 実行 (params = {})
# ... テンプレートとそのドキュメントをロードします
テンプレート:: DetectFields 。 call ( io 、添付ファイル: ドキュメント 、推論 :) do | (attachment_uuid、ページ、フィールド) |
Redisプール 。 call ( 'PUBLISH' , params [ 'channel_key' ], {attachment_uuid :, page :,fields: }.to_json )
終わり
Redisプール 。 call ( 'PUBLISH' , params [ 'channel_key' ], { completed: true }.to_json )
終わり
終わり
Rails 側では、SSE エンドポイントはプレーンな ActionController::Live アクションであり、チャネル キーを生成し、それをサブスクライブし、ジョブをキューに入れ、パブリッシュされた各メッセージを検出されたフィールドの JSON 配列としてブラウザーに中継します。
コピー
コピーされました
クラス TemplatesDetectFieldsCloudController < ApplicationController
ActionController :: Live を含める
デフォルト作成
sse = SSE 。新しい (応答のストリーム)
channel_key = SecureRandom 。 uuid
pubsub = RedisPool :: POOL 。パブサブ
パブサブ 。呼び出し ( 'SUBSCRIBE' 、channel_key )
TemplateDetectFieldsJob 。 Perform_async ( 'template_id' => @template . id , 'channel_key' => channel_key )
ループドゥ
タイプ、キー、ペイロード = pubsub 。 next_event (タイムアウト)
データ = JSON 。解析 (ペイロード)
っせ。書き込み（データ）
|| データ [ 'completed' ] の場合は中断します ||データ [ 'エラー' ]
終わり
確実にする

[切り捨てられた]

## Original Extract

How DocuSeal runs a GPU object detection model for PDF form field detection inside a Rails monolith with TensorRT and Ruby, with no Python in the request path.

DocuSeal
18k 18k
Solutions
Solutions
Document Signing API
Build automations and workflows
Embedded Signing
Seamlessly integrate signing experience
Embed document signing in React App
Enterprise
Explore our enterprise solutions
On-premises
Self-hosted DocuSeal on your servers
Accept Payments
Collect payments during signing
Document Signing for SaaS
Add document signing to your product
All Solutions
Explore other solutions
Software and Technology
Healthcare
Real Estate
Education Institution
Financial Services
Resources
Resources
Compliance
Learn about how we follow eSign standards
API Reference
Explore our API endpoints and webhooks
Blog
Stay updated with our latest articles
Documentation
Find answers and learn about the features
Developer Guides
Development guides and examples
How-To's
User tutorials and resources
GDPR
Explore our GDPR-compliant EU Cloud
Qualified Electronic Signature
Sign with the highest level of trust
Sign in
Get Started
DocuSeal
Solutions
Document Signing API
Build automations and workflows
Embedded Signing
Seamlessly integrate document signing
Enterprise
Explore our enterprise solutions
On-premises
Self-hosted DocuSeal on your servers
React Document Signing
Embed document signing in React App
Document Signing for SaaS
Add document signing to your product
Accept Payments
Collect payments during signing
All Solutions
Explore other solutions
Resources
Compliance
Learn about how we follow eSign standards
Blog
Stay updated with our latest articles
API Reference
Explore our API endpoints and webhooks
Documentation
Find answers and learn about the features
Developer Guides
Development guides and examples
How-To's
User tutorials and resources
GDPR
Explore our GDPR-compliant EU Cloud
Qualified Electronic Signature
Sign with the highest level of trust
Knowledge Based Authentication
Verify identity with knowledge-based questions
Sign in
Ask AI
All posts
Running GPU AI workloads with a Ruby on Rails monolith
DocuSeal is an open-source e-signature platform, and to get started, users need to upload their documents and map the fields to be filled or signed. Manually adding dozens or even hundreds of fields to complex forms can be tedious, so we implemented an AI field detection feature based on a computer vision model trained to detect fields on a wide range of publicly available PDF forms. We run our AI workloads on an NVIDIA GPU instance within the Ruby process of our Rails monolith, with no Python and no microservices.
We like the convenience of building our Rails monolith app with Ruby, and to make it easy to develop and maintain AI features, we also wanted to have AI field detection within the same Ruby on Rails monolith app. To achieve this, we’ve built an AI field detection inference pipeline with Ruby and the Sidekiq async jobs processor running on an NVIDIA T4 GPU instance. The screenshot below displays nvtop NVIDIA GPU utilization by the Ruby Sidekiq process on a production GPU worker during fields detection.
To run computer vision models efficiently on GPUs, NVIDIA provides the TensorRT inference runtime. TensorRT exposes a C++ API, and no Ruby binding for it existed, so we built a very small single-file Rice C++ binding that links only the methods a forward pass needs: loading an engine, inspecting its tensors, binding device memory, executing, and synchronizing the CUDA stream. We made these TensorRT Ruby bindings open source under the Apache 2.0 license, available on GitHub .
Copy
Copied
gem install tensorrt
Building and installing the gem requires TensorRT and NVIDIA CUDA on the system. The entire TensorRT binding is a single TensorRT::Engine class with 11 methods:
Copy
Copied
require 'tensorrt'
engine = TensorRT :: Engine . new ( model_path , verbose: false )
engine . num_io_tensors # Number of input/output tensors
engine . get_tensor_name ( index ) # Tensor name by index
engine . is_input? ( name ) # Check if tensor is input
engine . get_tensor_shape ( name ) # Shape as array [1, 3, 640, 640]
engine . get_tensor_bytes ( name ) # Size in bytes
engine . get_tensor_dtype ( name ) # Data type, e.g. float32
engine . set_tensor_address ( name , device_ptr ) # Bind GPU memory
engine . execute # Synchronous execution
engine . enqueue # Asynchronous execution
engine . get_stream # CUDA stream handle
engine . stream_synchronize # Wait for stream completion
The inference pipeline
The pipeline consists of three stages:
Preprocessing. Render the PDF page and prepare the input tensor.
Forward pass. Run the model on the GPU with TensorRT.
Postprocessing. Convert output tensors into field coordinates on the page.
1. Preprocessing: from PDF page to input tensor
First, PDF pages are rendered with PDFium. The rendered image is scaled to the model input resolution with aspect ratio preserved, padded to a square, normalized with the standard ImageNet mean and standard deviation, and transposed from HWC to CHW layout. Image operations are performed with ruby-vips, the libvips binding, and for tensor operations we use Numo, a Ruby alternative to NumPy:
Copy
Copied
MEAN = [ 0.485 , 0.456 , 0.406 ]. freeze
STD = [ 0.229 , 0.224 , 0.225 ]. freeze
scale = [ resolution . to_f / image . width , resolution . to_f / image . height ]. min
resized = image . resize ( scale , vscale: scale , kernel: :lanczos3 )
pad_x = (( resolution - ( image . width * scale ). round ) / 2.0 ). round
pad_y = (( resolution - ( image . height * scale ). round ) / 2.0 ). round
image = resized . embed ( pad_x , pad_y , resolution , resolution , background: [ 255 , 255 , 255 ])
# ImageNet normalization
image /= 255.0
image = ( image - MEAN ) / STD
img_array = Numo :: SFloat . from_binary ( image . write_to_memory , [ resolution , resolution , 3 ])
input_tensor = img_array . transpose ( 2 , 0 , 1 ). reshape ( 1 , 3 , resolution , resolution )
2. Forward pass: running the model on the GPU
To run the forward pass, the input tensor is cast to the data type the engine declares, serialized to a binary string, written into a host buffer, and copied to GPU VRAM. Execution is started with enqueue , which submits the work and returns immediately. retrieve is then used to wait for the forward pass to complete and read the output:
Copy
Copied
def enqueue ( input_tensor )
host_ptr = FFI :: MemoryPointer . new ( :uint8 , @input_size )
host_ptr . write_bytes ( Numo :: SFloat . cast ( input_tensor ). to_binary )
TensorRT :: CUDA . memcpy_htod_async ( @input_ptr , host_ptr , @input_size , @cuda_stream )
@engine . enqueue # non-blocking
end
def retrieve
@engine . stream_synchronize
{ dets: read_output ( @dets_ptr , @dets_size ),
labels: read_output ( @labels_ptr , @labels_size ) }
end
def read_output ( device_ptr , size )
host_ptr = FFI :: MemoryPointer . new ( :uint8 , size )
TensorRT :: CUDA . memcpy_dtoh ( host_ptr , device_ptr , size )
Numo :: SFloat . from_binary ( host_ptr . read_bytes ( size ))
end
In production these host buffers are allocated once per thread and reused across jobs, so steady state inference performs no allocation on the Ruby side. Each Sidekiq thread holds its own engine instance.
3. Postprocessing: from output tensors to form fields
The engine returns box predictions and per-class logits. Postprocessing applies a sigmoid to turn the logits into scores, takes the highest scoring class for each box, converts boxes from center to corner format, reverses the scale and padding applied during preprocessing, and discards detections below the confidence threshold. Non-maximum suppression then removes duplicate boxes over the same region, and the coordinates are normalized to the 0..1 range the form builder uses.
Each surviving detection becomes a Field object with relative page coordinates and a field type:
Copy
Copied
Field . new ( type: 'text' , x: 0.6123 , y: 0.8402 , w: 0.2510 , h: 0.0338 , confidence: 0.94 )
Asynchronous pipelining
Since stages 1 and 3 (preprocessing and postprocessing) execute on the CPU and stage 2 (the forward pass) executes on the GPU, the CPU can sit idle waiting for the GPU forward pass, and the GPU can sit idle waiting for CPU preprocessing and postprocessing. To increase throughput we built an async pipeline where the CPU preprocesses the next PDF page while the GPU is still running the forward pass on the current one.
To achieve this we utilize TensorRT asynchronous execution, where enqueue submits work to the CUDA stream and returns without blocking, and stream_synchronize blocks until the stream completes. Separating the two calls and returning retrieve as a lambda makes the deferral explicit:
Copy
Copied
def enqueue ( input :, ** )
inference . enqueue ( * input )
-> { inference . retrieve }
end
The page loop uses this to overlap the stages. The forward pass for page N is enqueued, page N+1 is rendered and preprocessed on the CPU while the GPU is executing, and the results for page N are read only after the next page’s input is prepared:
Copy
Copied
image = prepare_page_image ( doc . get_page ( page_indexes . first ))
current_args = inference . prepare_input ( image , ** prep_opts )
current_task = inference . enqueue ( ** current_args , ** infer_opts )
page_indexes . each_with_index do | current_page_number , i |
next_n = page_indexes [ i + 1 ]
if next_n # CPU preprocesses the next page while the GPU executes the current one
next_image = prepare_page_image ( doc . get_page ( next_n ))
next_args = inference . prepare_input ( next_image , ** prep_opts )
end
outputs = current_task . call # wait until the GPU completes the forward pass and read outputs
next_task = inference . enqueue ( ** next_args , ** infer_opts ) if next_args # start the next page's forward pass, non-blocking
fields = inference . process_outputs ( outputs , ** current_args , ** infer_opts ) # convert output tensors into field coordinates
yield [ attachment_uuid , current_page_number , fields ] # publish this page's fields
current_args = next_args
current_task = next_task
end
CPU and GPU work overlap for the duration of the document. Measured against the same pipeline executed synchronously, this yields approximately 80% higher throughput.
Streaming results to the browser
Fields detection runs in a Sidekiq worker on the GPU instance, and to stream results back to the user’s browser over Server-Sent Events, we use Redis pub/sub to carry messages from the worker process to the web process. Using Rails Action Cable with WebSockets would be a viable option as well, but we chose to stick with a plain Rails Live SSE controller.
The worker publishes each page’s fields as soon as that page completes, allowing us to show the user a live progress indicator of the number of pages processed. The block passed to DetectFields.call is invoked once per page, from the pipelined loop shown above:
Copy
Copied
class TemplateDetectFieldsJob
include Sidekiq :: Job
def perform ( params = {})
# ... load the template and its documents
Templates :: DetectFields . call ( io , attachment: document , inference :) do | ( attachment_uuid , page , fields ) |
RedisPool . call ( 'PUBLISH' , params [ 'channel_key' ], { attachment_uuid :, page :, fields: }. to_json )
end
RedisPool . call ( 'PUBLISH' , params [ 'channel_key' ], { completed: true }. to_json )
end
end
On the Rails side, the SSE endpoint is a plain ActionController::Live action where we generate a channel key, subscribe to it, enqueue the job, and relay each published message back to the browser as a JSON array of detected fields:
Copy
Copied
class TemplatesDetectFieldsCloudController < ApplicationController
include ActionController :: Live
def create
sse = SSE . new ( response . stream )
channel_key = SecureRandom . uuid
pubsub = RedisPool :: POOL . pubsub
pubsub . call ( 'SUBSCRIBE' , channel_key )
TemplateDetectFieldsJob . perform_async ( 'template_id' => @template . id , 'channel_key' => channel_key )
loop do
type , key , payload = pubsub . next_event ( TIMEOUT )
data = JSON . parse ( payload )
sse . write ( data )
break if data [ 'completed' ] || data [ 'error' ]
end
ensur

[truncated]
