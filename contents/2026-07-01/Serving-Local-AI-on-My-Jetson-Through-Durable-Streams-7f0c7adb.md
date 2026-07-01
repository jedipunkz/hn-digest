---
source: "https://s2.dev/blog/local-ai"
hn_url: "https://news.ycombinator.com/item?id=48741231"
title: "Serving Local AI on My Jetson Through Durable Streams"
article_title: "Serving Local AI on my Jetson through Durable Streams · S2"
author: "shikhar"
captured_at: "2026-07-01T01:44:34Z"
capture_tool: "hn-digest"
hn_id: 48741231
score: 1
comments: 0
posted_at: "2026-07-01T01:00:03Z"
tags:
  - hacker-news
  - translated
---

# Serving Local AI on My Jetson Through Durable Streams

- HN: [48741231](https://news.ycombinator.com/item?id=48741231)
- Source: [s2.dev](https://s2.dev/blog/local-ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T01:00:03Z

## Translation

タイトル: 耐久性のあるストリームを介して Jetson でローカル AI を提供する
記事のタイトル: Durable Streams · S2 を介して Jetson でローカル AI を提供する
説明: セルフホストの時代です!

記事本文:
Durable Streams を介して Jetson でローカル AI を提供する · S2 一般提供とシード ラウンド ドキュメント ブログ プレイグラウンドの価格設定 テーマの切り替え GitHub S2 チームとの接続 ダッシュボード テーマの切り替え メイン メニューを開く このページについて 2026 年 6 月 29 日 Durable Streams を介して Jetson でローカル AI を提供する
ローカル AI がますます実用的になるにつれて、私は独自のモデルを自己ホストし、サードパーティのプロバイダーを使用せずに独立してワークロードを実行したいと考えました。また、一部のユーザーにローカル モデルを確実に提供することも検討したいと考えました。 NVIDIA の Jetson シリーズは出発点として最適です。私は、「最も手頃な価格の生成 AI スーパーコンピューター」とも呼ばれる Jetson Orin Nano スーパー キットを選びました。 1024 個の CUDA コアと 32 個のテンソル コアがあり、67 TOPS (1 秒あたり 1 兆回の演算) と評価されています。これは、ニューラル テキスト読み上げモデルである Kokoro-82M を利用した小さなテキスト読み上げアプリである私の小さな実験には十分なはずです。
常にたくさんのテキストを読みたいわけではなく、むしろ聞きたいという必要性からインスピレーションを得たものです。したがって、テキストを選択し、音声を選択し、後で戻ったり、人々と共有したりできるリンクを取得できるものが必要です。のために
今ではそれはページにテキストを貼り付けることを意味しますが、最終的には同じコアアプリの上にあるより優れたフロントエンドとなる、さらに遅延防止のものが欲しいと思っています。アプリ自体を超えて、ローカル推論のための小さなリファレンス アーキテクチャにたどり着きたいと考えています。これは、クリーンな API を公開する自己完結型のサービング レイヤーであり、同じセットアップで Web アプリ、CLI、または別のサービスを再作業なしでサポートできるようにします。
streamtts.dev で試してみてください (私の Jetson で自己ホストされています! 😉):
通常のリクエスト/レスポンス API ではありません
これを設計する最も簡単な方法は次のようになります。
POST /生成
待ってください
audio.mp3 を返す this.classList.remove('rehype-pretty-copied'), 3000

);">
推論は通常の Web リクエストよりも遅くなります。この Jetson の Kokoro はリアルタイムよりも速く音声を生成できますが、それでも GPU の仕事です。 1 分の音声には何秒もの計算時間がかかる場合があります。コールドの最初の文は、モデル スタックがウォームアップするまで遅くなる可能性があります。複数のユーザーが同時に送信すると、ブロッキング リクエストは GPU 上で待機するソケットの列に変わります。
出力も当然増加します。 TTS では、リスナーが何かを聞く前に段落全体を終える必要はありません。モデルは 1 つの文を生成し、その文を MP3 にエンコードし、どこかに追加して、次に進むことができます。すべてを 1 つの応答本文に強制的に含めると、ワークロードの最適なプロパティが捨てられてしまいます。
そして、その結果を共有できるようにしたいと考えています。ユーザーは、モデルがすべてのバイトを生成するのを「待つ」ことができるリンクにすぐに誘導される必要があります。 Jetson がまだ動作している間にそれを開いた場合、プレフィックスが聞こえてから、ライブ エッジをたどるはずです。
リクエストとレスポンスから始めると、最終的には次のようなインフラストラクチャの山を追加することになります。
完成したファイルのオブジェクトストレージ
これらはすべて合理的です。しかし、次の 1 つの基本的な約束を組み合わせると、非常に多くのことが達成されます。
今すぐ仕事を受け入れる
後で出力を生成する
読者は this.classList.remove('rehype-pretty-copied'), 3000);"> に従ってください。
このリクエストは間違った生涯のように感じます。ネットワークの中断があっても推論ジョブがシームレスに機能するようにしたいと考えています。また、ドロップされたブラウザーのタブによって実行中の世代が強制終了されることも望ましくありません。したがって、出力は完了する前に ID を持っている必要があり、リーダーは最初から始めて最後に追いつくか、後で戻って同じバイトを再生できる必要があります。
作品を提出する
出力ストリームをすぐに取得する
ワーカーがモデル出力を追加します
クライアントはストリームを待機します this.classList.remove('rehype-pretty-copied'), 3000);">

これらすべては、耐久性のあるストリーム上できれいに抽象化できます。ストリームは順序付けられたレコードのシーケンスであり、レコードはほんの数バイト (ここでは、オーディオの塊と少量のメタデータ) です。永続的とは、すべてのレコードが永続化されるため、何も失われず、リーダーが後で戻ってきてまったく同じバイトを再生できることを意味します。 2 つを組み合わせると、シンプルだが強力な構成要素が得られます。
レコードを末尾に追加すると、リーダーは先頭から開始して既知のシーケンス番号を探すか、末尾に座って次のレコードが到着するのを待つことができます。ストリーム ストアでは、名前付きのタイムラインが提供されます。
レコードの追加
seq_num=N から読み取り
ライブ レコードの TAIL this.classList.remove('rehype-pretty-copied'), 3000);">
各レコードは進行の単位です。レコードには、シーケンス番号、タイムスタンプ、ヘッダー、および本文が含まれます。 StreamTTS にはそれ以上の構造は必要ありません。レコードを次のように表します。
ヘッダー:
e：オーディオ
私：3
d:4210
t:「文章のテキスト」
本体:
<生の mp3 バイト>
# e = イベントの種類
# i = インデックス
# d = 持続時間 (ミリ秒)
# t = 文のテキスト
# e = イベントの種類
# i = インデックス
# d = 持続時間 (ミリ秒)
# t = 文テキスト" class="rehype-pretty-copy" onclick="navigator.clipboard.writeText(this.attributes.data.value);this.classList.add('rehype-pretty-copied');window.setTimeout(() => this.classList.remove('rehype-pretty-copied'), 3000);">
出力は次のようになります。
パブ/キャスト/4LwnHZDl_vFC
シーケンス 0 メタ
シーケンス 1 開始
シーケンス 2 音声文 0
シーケンス 3 音声文 1
シーケンス 4 音声文 2
seq 5 eos # ストリームの終わり this.classList.remove('rehype-pretty-copied'), 3000);">
そのストリームは、オーディオ ファイル、ライブ フィード、再生ログ、および進行状況インジケーターです。これは、Web サーバー、GPU ワーカー、およびリンクを開くすべてのブラウザ間の契約でもあります。書き手は誰が聞いているかを知る必要はありません。読者はそうします

作家がまだ生きているかどうかを知る必要はありません。双方とも、名前付きの 1 つのレコード シーケンスについて合意するだけです。
接続専用の SSE または WebSocket はライブ配信に最適ですが、それ自体では耐久性のある再生を提供しません。現在接続されているクライアントにバイトを移動します。クライアント自身では、遅れて到着したり、切断したり、ページを更新したりするクライアントのバイトを覚えていません。したがって、誰も接続していない場合、WebSocket メッセージが永続的に送信される場所はありません。クライアントがドロップした場合、サーバーはクライアントが見逃したものを記憶するために他のストアを必要とします。生成の実行中に 2 番目のリスナーが同じリンクを開いた場合、WebSocket 接続はサーバーに先頭を再生してからライブ エッジに従う方法を指示しません。 SSE/WebSocket の隣にデータベースまたはオブジェクト ストアを配置することで、この問題を完全に解決できます。しかし現在、ライブ配信とリプレイは 2 つの別個の部分であり、同意する必要があります。
耐久性のあるストリームで、その分割を統一できます!ワーカーは出力を 1 回追加し、ライブ リスナーがストリームを追跡します。遅いリスナーは seq_num=0 から読み取り、同じストリームを追跡できます。リプレイとライブ再生は同じ読み取りパスであり、異なるオフセットから開始するだけです。
S2 Lite は、S2 耐久性ストリーム API のオープンソースのセルフホスト型シングルバイナリ実装です。この設定では、耐久性のあるストレージとしてローカル ディスクを備えたローカルホスト上で実行され、追加、読み取り、末尾、およびロングポーリングのセマンティクスを備えたストリームが提供されます。
s2 lite --local-root var/s2lite-data --port 4002 --no-cors
this.classList.remove('rehype-pretty-copied'), 3000);">
まず、名前空間として機能する盆地を作成し、サービス全体をいくつかの名前付きストリームとしてモデル化します。以下の矢印は、どのコンポーネントが各ストリームに追加され、どのコンポーネントがそこから読み取るかを示しています。
いくつかのストリームがすべてのキャストで共有されます

s:
jobs は取り込みログです: 推論リクエストごとに 1 つのレコード
jobs/_cursor は、ワーカーのジョブへのコミットされた読み取りオフセットを保持します。
jobs/dead は過去の再試行に失敗したジョブを収集します
progress/done は完了したキャストごとに 1 つのレシートを取得します
そして、各キャストは独自の 2 つのストリームを追加します。
category/<id> はプライベート レシピです: フルテキスト、音声、タイトル、作成時刻
pub/casts/<id> はパブリック出力ストリームです:meta、start、audio...、eos
s2 = S2(
os.environ.get( "S2_ACCESS_TOKEN" , "ローカルトークン" ),
エンドポイント = エンドポイント(
アカウント = lite_url、
盆地 = lite_url、
）、
)
config = BasinConfig(
default_stream_config = StreamConfig(
storage_class = ストレージクラス。エクスプレス、
保持ポリシー = RETENTION_SECS 、
)
)
await s2.ensure_basin(basin, config = config) this.classList.remove('rehype-pretty-copied'), 3000);">
各音声レコードのヘッダーには文のテキストとミリ秒単位の継続時間が含まれ、本文には生の MP3 バイトが含まれます。テキストはブラウザにキャプションとシーク ポイントを提供します。期間により、プレーヤーはチャンクをスケジュールできます。ブラウザは常に seq_num=0 で追跡を開始します。
ストリームが完了すると、ブラウザは eos を読み取って停止します。ワーカーがまだ追加している場合、ブラウザは既存のプレフィックスを読み取り、末尾に到達し、次のレコードを待ちます。ブラウザー プレーヤーもストリーム形状に基づいて構築されています。メディア ソース拡張機能を使用したり、成長する MP3 ファイルを構築したりすることはありません。各オーディオ レコードは、完全な文サイズの MP3 チャンクです。ブラウザは各文サイズの MP3 チャンクを受信し、Web Audio API でデコードし、
そしてそれを仮想タイムライン上に配置します。
単一の Jetson は、弾性推論クラスターのように動作することはできません 😅。たとえば 3 人がテキストを送信した場合、他の全員が待っている間に最初の長い段落が完全に終了することは望ましくありません。ワーカーは複数のキャストをアクティブにして追跡します

各ストリームが実時間の再生に対してどれだけ進んでいるか:
def リード (自分自身) -> float :
return self .total_ms / 1000.0 - (time.monotonic() - self .started) float:
return self.total_ms / 1000.0 - (time.monotonic() - self.started)" class="rehype-pretty-copy" onclick="navigator.clipboard.writeText(this.attributes.data.value);this.classList.add('rehype-pretty-copied');window.setTimeout(() => this.classList.remove('rehype-pretty-copied'), 3000);">
正のリードは、ストリームが再生前にバッファリングされたオーディオを生成したことを意味します。マイナスのリードは、リスナーがライブのテールに追い付いていることを意味します。
同時実行の上限までジョブを許可します
リードが最も低いアクティブなストリームを選択する
それに対して正確に 1 つの文を生成する
その文を追加します
リードを再計算する
これを繰り返します.classList.remove('rehype-pretty-copied'), 3000);">
すべてのアクティブなストリームが快適に先行している場合、ワーカーは 1 つのストリームを完了まで全力疾走してライブ出力スケジューリングを作成するのではなく、少しの間スリープします。目標は、複数のパブリック ストリームを再生可能な状態に保つことです。公平性の単位は要求ではなく、追加された 1 文です。
リクエストが到着しても、Web プロセスはモデルをロードしません。テキストと音声を検証し、決定的な ID を計算し、音声が表示される場所を作成します。
def content_id (テキスト: str 、音声: str ) -> str :
h = hashlib.sha256( f " { voice }\x00{ text.strip() } " .encode()).digest()
returnbase64.urlsafe_b64encode(h).decode().rstrip( "=" )[: 12 ] str:
h = hashlib.sha256(f"{音声}\x00{text.strip()}".encode()).digest()
returnbase64.urlsafe_b64encode(h).decode().rstrip("=")[:12]" class="rehype-pretty-copy" onclick="navigator.clipboard.writeText(this.attributes.data.value);this.classList.add('rehype-pretty-copied');window.setTimeout(() => this.classList.remove('rehype-pretty-copied'), 3000)

;">
同じ音声を持つ同じテキストは同じストリームにマッピングされます。これにより、繰り返しの送信がキャッシュ ヒットに変わります。
完全なレシピを含むカタログ/<id> を要求します
メタレコードを使用して pub/casts/<id> を要求します
ジョブ ストリームに 1 つのジョブを追加します
重要な操作はクレームです。 S2 は、 match_seq_num による条件付き追加をサポートします。 StreamTTS は match_seq_num=0 を使用します。これは、「このストリームが空の場合にのみ追加する」ことを意味します。
ペイロード = {
"レコード" : [{ "本体" : json.dumps(本体, 区切り文字 = ( "," , ":" ))}],
"match_seq_num" : 0 、
} this.classList.remove('rehype-pretty-copied'), 3000);">
2 人が同じテキストを同時に送信した場合、ちょうど 1 つのリクエストがクレームを獲得し、ジョブをキューに入れます。もう 1 つは同じリンクを取得し、同じ出力ストリームを追跡します。
その 1 つの追加により、ロック テーブル、一意性制約、および重複排除キャッシュが置き換えられます。
労働者は耐久消費者である
ワーカーは、モデルを所有し、GPU にアクセスする唯一のプロセスです。ジョブ ストリームから読み取り、Kokoro-82M を実行し、音声レコードをキャスト ストリームに追加します。
起動時に、ワーカーは jobs/_cursor から最後にコミットされたオフセットを読み取ります。
ジョブ/_カーソル
{"オフセット": 123} this.classList.remove('rehype-pretty-copied'), 3000);">
次に、そのオフセットから開始してジョブを読み取ります。

[切り捨てられた]

## Original Extract

Its time to self-host!

Serving Local AI on my Jetson through Durable Streams · S2 General availability & our seed round Docs Blog Playground Pricing Toggle theme GitHub Connect with S2 team Dashboard Toggle theme Open main menu On this page June 29, 2026 Serving Local AI on my Jetson through Durable Streams
With local AI feeling more and more practical, I wanted to self-host my own models and run my workloads independently without any third-party provider in the mix, and also look into serving my local model to some users reliably. The Jetson series by NVIDIA is a great starting point, and I went with the Jetson Orin Nano Super kit , aka “The most affordable generative AI supercomputer”! It has 1024 CUDA cores and 32 tensor cores and is rated at 67 TOPS (trillion operations per second), which should be good enough for my little experiment which is a small text-to-speech app powered by Kokoro-82M , a neural text-to-speech model.
It is mostly inspired out of need that I don't want to always read a lot of text, but would rather hear it. So I want something where I select some text, pick a voice, and get a link which I can come back to later or share with people. For
now that means pasting text into a page, but I'd want something even more lazy-proof eventually which would be a nicer frontend on top of the same core app. Beyond the app itself, I want to land on a small reference architecture for local inference: a self-contained serving layer that exposes a clean API, so the same setup can back a web app, a CLI, or another service without rework.
Try it out at streamtts.dev (It is self-hosted on my Jetson! 😉):
Not a normal Request/Response API
The simplest way to architect this would be:
POST /generate
wait
return audio.mp3 this.classList.remove('rehype-pretty-copied'), 3000);">
Inference is slower than a normal web request. Kokoro on this Jetson can produce speech faster than realtime, but it is still a GPU job. A minute of audio can take many seconds of compute. A cold first sentence can be slower while the model stack warms up. If multiple users submit at once, a blocking request turns into a line of sockets waiting on the GPU.
The output is also naturally incremental. TTS does not need to finish the entire paragraph before the listener hears anything. The model can generate one sentence, encode that sentence to MP3, append it somewhere, and move on. If I force the whole thing into a single response body, I throw away the best property of the workload.
And I want the result to be shareable. The user should be directed to a link immediately where they can "await" the model to produce all the bytes. If they open it while the Jetson is still working, they should hear the prefix and then follow the live edge.
If we start with request-response, we end up adding a pile of infrastructure like:
object storage for the finished file
All of this is reasonable. But together, it is a lot for one basic promise:
accept work now
produce output later
let readers follow along this.classList.remove('rehype-pretty-copied'), 3000);">
The request feels like the wrong lifetime for this. I want the inference job to work seamlessly across network disruptions. I also do not want a dropped browser tab to kill a running generation. Thus the output should have an identity before it is complete, and readers should be able to start at the beginning, catch up to the tail, or come back later and replay the same bytes!
submit work
get an output stream immediately
worker appends model output
client awaits the stream this.classList.remove('rehype-pretty-copied'), 3000);">
All of this can be cleanly abstracted over durable streams. A stream is an ordered sequence of records, where a record is just some bytes (here, a chunk of audio plus a little metadata). Durable means every record is persisted, so nothing is lost and a reader can come back later and replay the exact same bytes. Putting the two together, we get a simple but powerful building block.
Append records to the tail, and readers can start at the head, seek to a known sequence number, or sit at the tail and wait for the next record to arrive. A stream store gives you named timelines:
APPEND record
READ from seq_num=N
TAIL for live records this.classList.remove('rehype-pretty-copied'), 3000);">
Each record is the unit of progress. A record has a sequence number, timestamp, headers, and a body. StreamTTS does not need much more structure than that. We represent records like so:
headers:
e: audio
i: 3
d: 4210
t: "sentence text"
body:
<raw mp3 bytes>
# e = event type
# i = index
# d = duration (ms)
# t = sentence text
# e = event type
# i = index
# d = duration (ms)
# t = sentence text" class="rehype-pretty-copy" onclick="navigator.clipboard.writeText(this.attributes.data.value);this.classList.add('rehype-pretty-copied');window.setTimeout(() => this.classList.remove('rehype-pretty-copied'), 3000);">
And the output will be shaped like:
pub/casts/4LwnHZDl_vFC
seq 0 meta
seq 1 start
seq 2 audio sentence 0
seq 3 audio sentence 1
seq 4 audio sentence 2
seq 5 eos # end of stream this.classList.remove('rehype-pretty-copied'), 3000);">
That stream is the audio file, the live feed, the replay log, and the progress indicator. It is also the contract between the web server, the GPU worker, and every browser that opens the link. The writer does not need to know who is listening. The reader does not need to know whether the writer is still alive. Both sides just agree on one named sequence of records.
Connection-only SSE or WebSockets are great for live delivery, but they do not give you durable replay by themselves. They move bytes to clients that are currently connected. They do not, on their own, remember the bytes for clients that arrive late, disconnect, or refresh the page. So if nobody is connected, there is nowhere durable for a websocket message to go. If a client drops, the server needs some other store to remember what that client missed. If a second listener opens the same link while generation is still running, the websocket connection does not tell the server how to replay the beginning and then follow the live edge. You can absolutely solve this by putting a database or object store next to SSE/WebSockets. But now live delivery and replay are two separate pieces that have to agree.
With a durable stream, that split can be unified! The worker appends output once and a live listener tails the stream. A late listener can read from seq_num=0 and then tails the same stream. Replay and live playback are the same read path, just starting from different offsets.
S2 Lite is an open source self-hosted, single-binary implementation of the S2 durable streams API. In this setup, it runs on localhost with local disk for durable storage and gives me streams with append, read, tail, and long-polling semantics.
s2 lite --local-root var/s2lite-data --port 4002 --no-cors
this.classList.remove('rehype-pretty-copied'), 3000);">
We start by creating a basin, which acts as a namespace, and model the whole service as a handful of named streams. The arrows below show which component appends to each stream and which reads from it:
A few streams are shared across all casts:
jobs is the intake log: one record per inference request
jobs/_cursor holds the worker's committed read offset into jobs
jobs/dead collects jobs that failed past retries
progress/done gets one receipt per completed cast
And each cast adds two streams of its own:
catalog/<id> is the private recipe: full text, voice, title, created time
pub/casts/<id> is the public output stream: meta, start, audio..., eos
s2 = S2(
os.environ.get( "S2_ACCESS_TOKEN" , "local-token" ),
endpoints = Endpoints(
account = lite_url,
basin = lite_url,
),
)
config = BasinConfig(
default_stream_config = StreamConfig(
storage_class = StorageClass. EXPRESS ,
retention_policy = RETENTION_SECS ,
)
)
await s2.ensure_basin(basin, config = config) this.classList.remove('rehype-pretty-copied'), 3000);">
Each audio record carries the sentence text and duration in milliseconds in its headers, and the raw MP3 bytes in its body. The text gives the browser captions and seek points. The duration lets the player schedule chunks. The browser always starts tailing at seq_num=0 .
If the stream is complete, the browser reads through eos and stops. If the worker is still appending, the browser reads the existing prefix, reaches the tail, and waits for the next record. The browser player is also built around the stream shape. It does not use Media Source Extensions or build one growing MP3 file. Each audio record is a complete sentence-sized MP3 chunk. The browser receives each sentence-sized MP3 chunk, decodes it with the Web Audio API ,
and places it on a virtual timeline.
A single Jetson can’t behave like an elastic inference cluster 😅. If lets say three people submit text, I do not want the first long paragraph to finish completely while everyone else waits. The worker keeps several casts active and tracks how far ahead each stream is relative to wall-clock playback:
def lead (self) -> float :
return self .total_ms / 1000.0 - (time.monotonic() - self .started) float:
return self.total_ms / 1000.0 - (time.monotonic() - self.started)" class="rehype-pretty-copy" onclick="navigator.clipboard.writeText(this.attributes.data.value);this.classList.add('rehype-pretty-copied');window.setTimeout(() => this.classList.remove('rehype-pretty-copied'), 3000);">
Positive lead means the stream has generated audio buffered ahead of playback. Negative lead means the listener is catching up to the live tail.
admit jobs up to the concurrency cap
pick the active stream with the lowest lead
generate exactly one sentence for it
append that sentence
recompute lead
repeat this.classList.remove('rehype-pretty-copied'), 3000);">
When every active stream is comfortably ahead, the worker sleeps for a tiny bit instead of sprinting one stream to completion creating live-output scheduling. The goal is to keep multiple public streams playable. The unit of fairness is not a request, but one appended sentence.
When a request comes in, the web process does not load the model. It validates the text and voice, computes a deterministic id, and creates a place where audio will appear.
def content_id (text: str , voice: str ) -> str :
h = hashlib.sha256( f " { voice }\x00{ text.strip() } " .encode()).digest()
return base64.urlsafe_b64encode(h).decode().rstrip( "=" )[: 12 ] str:
h = hashlib.sha256(f"{voice}\x00{text.strip()}".encode()).digest()
return base64.urlsafe_b64encode(h).decode().rstrip("=")[:12]" class="rehype-pretty-copy" onclick="navigator.clipboard.writeText(this.attributes.data.value);this.classList.add('rehype-pretty-copied');window.setTimeout(() => this.classList.remove('rehype-pretty-copied'), 3000);">
Identical text with the same voice maps to the same stream. That turns repeated submissions into cache hits.
claim catalog/<id> with the full recipe
claim pub/casts/<id> with a meta record
append one job to the jobs stream
The important operation is the claim. S2 supports conditional append with match_seq_num . StreamTTS uses match_seq_num=0 , which means "only append if this stream is empty."
payload = {
"records" : [{ "body" : json.dumps(body, separators = ( "," , ":" ))}],
"match_seq_num" : 0 ,
} this.classList.remove('rehype-pretty-copied'), 3000);">
If two people submit the same text at the same time, exactly one request wins the claim and enqueues the job. The other gets the same link and tails the same output stream.
That one append replaces a lock table, a uniqueness constraint, and a dedupe cache.
The Worker is a Durable Consumer
The worker is the only process that owns the model and touches the GPU. It reads from the jobs stream, runs Kokoro-82M, and appends audio records to the cast stream.
On startup, the worker reads the last committed offset from jobs/_cursor :
jobs/_cursor
{"offset": 123} this.classList.remove('rehype-pretty-copied'), 3000);">
Then it reads jobs starting from that offset.

[truncated]
