---
source: "https://punnerud.github.io/pyspell/"
hn_url: "https://news.ycombinator.com/item?id=48576793"
title: "Show HN: ESP32 512kB – Tailscale, English to Python LLM and 8 containers local"
article_title: "PySpell — sandboxed Rust/Python expressions, live on ESP32"
author: "punnerud"
captured_at: "2026-06-17T21:42:25Z"
capture_tool: "hn-digest"
hn_id: 48576793
score: 1
comments: 0
posted_at: "2026-06-17T20:58:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ESP32 512kB – Tailscale, English to Python LLM and 8 containers local

- HN: [48576793](https://news.ycombinator.com/item?id=48576793)
- Source: [punnerud.github.io](https://punnerud.github.io/pyspell/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T20:58:08Z

## Translation

タイトル: 表示 HN: ESP32 512kB – Tailscale、英語から Python LLM およびローカルの 8 つのコンテナー
記事のタイトル: PySpell — サンドボックス化された Rust/Python 式、ESP32 で公開

記事本文:
PySpell を使えば Python の綴りができるのに、なぜ Python の綴りを学ぶ必要があるでしょうか?
言いたいことを平易な英語で言うと、デバイス上の小さなモデルがそれを実行可能な Python に変換します。
512 KB RAM を搭載した 10 ドルの ESP32 で 100% ローカル (PSRAM、クラウド、API キーなし)、
Tailscale 上のどこにでも到達可能。
GitHubで見る→ ・仕組み↓
ライブで試してみる — モデルはブラウザで実行されます
インストールもデバイスもクラウドも必要ありません。 ESP32 で実行されるものと同じ ~0.5 MB のモデル + サンドボックスが、ここでは WebAssembly としてロードされます。わかりやすい英語を入力してください。 Python を記述して実行します。
画面 + LED 最後の結果を保持 · リセット
ブラウザ上で 100% 動作します。携帯電話からアクセスできるようにしたいですか?このタブを独自のテールネット上のプライベート ノードにするには、以下の Connect Tailscale を使用します。仕組み →
オプトイン — クリックするまで何も接続しません。 Rust Tailscale クライアント (~321 kB WASM) は、WebSocket 経由でテールネットに参加します。ノードキーはこのブラウザに残ります。次に、テールネット上の任意のデバイスから http://<node-ip>/ を開いて、トンネル経由で同じデモを取得します。
このチップは、Web エージェント IDE、ネイティブ MCP サーバー、および REST API を提供します。
サンドボックスでコードを実行し（ライブ Web リクエスト / フェッチ サポート付き）、ドライブ
独自の画面 + LED — 同じ 0.5 メガバイトの RAM 上で最大 8 つの PySpell プロセスを並列処理します。
MicroPython と似ていますが、構文が 2 つあり、パーサーがデバイスに同梱されることはなく、英語が 3 番目の構文です。
PySpell プログラムは、単一の式 (Python)、またはいくつかの let バインディングとそれに続くものです。
末尾の式 (Rust)。数値、ブール値、文字列、またはリストの値として評価されます。無料
識別子は、ホスト提供の環境に対して評価時に解決されます: CLI 変数
ラップトップ上では、またはマイクロコントローラー上のライブデバイスの測定値。唯一の I/O はホストによって許可され、ホワイトリストに登録されているものです
フェッチjson ;ループ、関数、インポートはありません。

それが重要です: 小さく、速く、そして
他の場所から受け入れても安全です。
オフチップで提供されるオフライン AI コーディング エージェント
トンネル経由で http://<dongle>/ を開くと、カーソルのようなエージェントが表示されます。種類
「ライトを点滅させます」、「「こんにちは」というテキストを表示します」、「7 プラス 5 は何ですか」、
または「単語を逆にするとロボット」 — ~0.45 M パラメーターの言語モデル (< 500 kB、int8)
それを PySpell コードに変換し、チップ上でライブで実行し、結果を表示します。
物理的な動作 (画面が点灯し、RGB LED が点滅します)。ランタイム、モデル、トークナイザー、辞書
すべてはドングルからオフラインで提供されます。クラウドもキーもありません (OpenAI はオプションで、背後にあります)
⚙）。
これほど小さなモデルが役立つのは、一連のトリックがある場合に限られます。完全な記事は次のとおりです。
技術.MD 。見出し:
モデルがポイントし、ブラウザがコピーする
0.45 M モデルは任意のトークン (数値、文字列、リスト) を確実にコピーできないため、要求されません
に。それは小さなセマンティックディレクティブを発行します。ブラウザはリテラルの内容をそのままコピーします。
3 + 2 を計算する → print( 3 + 2 ) ;変更して追加する
減算 → @@ + ==> - 。引用符で囲まれたテキストはリテラルの内容であり、バイトごとにコピーされます。
語彙チェックの対象外となります。
デバイスは機能します。ブラウザが計算する
推論はクライアント側の WebAssembly で実行されます。 0.5 MB のモデル画像がフラッシュから送信されます。
TCP セグメントは一度に (HTTP 範囲)、チップの ~60 KB ヒープに常駐することはありません。反転した
エッジ推論: 制約付きデバイスがサービスとグレードを提供し、ブラウザーがモデルを実行します。
512 トークンの語彙は、all-MiniLM (22 M パラメータ) で埋め込まれ、128 ディメンションに PCA され、
品詞ベクトル、および凍結 — 小さなモデルは意味のある単語から始まります
わずかな予算を費やして幾何学を学ぶのではなく、幾何学を学べます。
語彙は辞書です
これらの同じ 512 トークン + 埋め込みが、入力検証のためにブラウザーに戻されます。
(「外では

モデルの語彙…") と、モデル自身の語彙に対する関連語の RAG を組み合わせます。
free_heap > 100000 および uptime_s < 60
距離 > 1000 の場合は 250、それ以外の場合は 0
0 < 温度 < 60 # チェーン
20人はピアに属さない
sum([1, 2, 3])
読み取り値[-1] # 負のインデックス
最大(a, b)
さび
free_heap > 100000 && uptime_s < 60
距離 > 1000 { 250 } の場合 { 0 }
let used = 合計 - 無料;使用済み * 100 / 合計
!peers.contains(20)
sum([1, 2, 3])
読み取り値[読み取り値.len() - 1]
最大(a, b)
言語リファレンス
古典的なワンライナー — 値をフェッチし、ドングルの画面に表示します。
show("オスロ: " + fetch_json(
"https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=59.91&lon=10.75",
"properties.timeseries.0.data.instant.details.air_temperature") + " C")
# 画面には次が表示されます: オスロ: 14.9 C (そして呼び出しはその文字列を返します)
ネットワークとJSON
fetch(url) + json_get(text, path) を使用すると、プログラムがライブ データを取得してデータを読み取ることができます。
それからフィールド。フェッチは仲介機能です。ホスト/デバイスがどのホストをフェッチするかを決定します。
到達可能 (許可リスト) であるため、プログラムは任意の URL に到達できません。
# ホスト CLI (ホストを明示的に許可):
pyspell 実行 oslo_temp.py --allow-host api.met.no
# oslo_temp.py は次のとおりです。
json_get(
fetch("https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=59.91&lon=10.75"),
"properties.timeseries.0.data.instant.details.air_temperature")
# → 14.9
メモ (デバイス): json_get はパス指定されているため、
RAM 内のドキュメント全体 — 一致した値のみを具体化します。 ESP32 上 (約 60 KB 空き、PSRAM なし)
fetch_json がストリームをストリーミングするため、大きな応答からフィールドを読み取ることが可能です。
HTTP(S) 本文を作成し、フィールドが見つかった瞬間に停止します (TLS バッファを早期に解放します)。つまり、年間最大 50 KB です。
応答が一度に RAM に収まる必要はありません。
# ESP32 上、Tailscale 上 (単一プロセス; ≈60 KB 空き; ライブで検証済み):

fetch_json(
"https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=59.91&lon=10.75",
"properties.timeseries.0.data.instant.details.air_temperature")
# → 14.9 (ドングルは年番号自体を取得しました)
ホスト上で実行中
# 自由変数をバインドして評価します。
カーゴ run -p pyspell-cli --examples/health.py を実行 --set free_heap=120000 --set uptime_ms=45000
#→本当
# ポータブル IR BLOB にコンパイルします。
カーゴ実行 -p pyspell-cli --examples/health.py をコンパイルします # → example/health.py.psb
# USB シリアルまたは対話型 REPL 経由でデバイスにライブ プッシュします。
カーゴ ラン -p pyspell-cli -- repl --port /dev/cu.usbmodem2101 --lang python
ESP32で実行中
ポータブル エバリュエーター ( pyspell-core 、 no_std + alloc ) は変更されずに実行されます。
ESP32-S3。プログラムは環境からライブデバイス変数を読み取ります。
demo/esp32-tailscale-pyspell ファームウェアは、Web テキスト ウィンドウと /run を追加します。
Tailscale トンネル内の API — ブラウザでデバイスの Tailscale IP を開き、
式を作成し、タイムアウトを設定してチップ上で実行します。 PySpell はネットワーク上に最大 62 KB しか追加しません
ファームウェア。
# Web ウィンドウ:
http://100.x.y.z/ を開きます
# POST (推奨): 本体にプログラム、クエリに lang/timeout。
# URL よりコード用のスペースが多く、パーセント エンコーディングはありません。
curl -X POST 'http://100.x.y.z/run?lang=py&timeout=10' --data 'free_heap > 100000' # → true
curl -X POST 'http://100.x.y.z/run?lang=rs&timeout=10' --data 'uptime_ms / 1000' # → 22
# GET (こちらもサポートされています): コードはクエリ内で URL エンコードされます。
カール 'http://100.x.y.z/run?lang=py&timeout=10&code=free_heap%20%3E%20100000' # → true
タイムアウトは秒単位で、1 ～ 60 に固定され、実際の時計の期限として適用されます。
デバイス。 1 つのリクエストは 1 つの TCP セグメント (約 1.2 KB) に収まる必要があります。POST では、そのセグメントの多くがコード用に残されます。
応答は text/plain (JSON ラッパーなし) です。
の

ESP32-S3 には 512 kB の SRAM があり、PSRAM はありませんが、完全な Tailscale ノードを実行します
(コントロール プレーンと DERP)、PySpell エバリュエーター、チップから提供されるブラウザ エージェント IDE、
ネイティブ MCP サーバー、および TLS から api.met.no。これが当てはまるのは、長い記憶トリックの連鎖があるからです。
CA チェーン検証の代わりに SPKI リーフキー ピニング — RSA-PSS 検証が 1 つ、6 kB なし
チェーン バッファー (TLS フェッチでは約 45→30 KB が減少します)。ヒープ許可ゲートが同時実行性を制限する
したがって、ピークヒープは K × フェッチごとにあり、決して N × フェッチごとにありません。
ネットマップは serde_json::from_reader を使用して HTTP/2 フレーム経由で読み取られるため、serde
巨大な DERPMap フィールドをバッファリングする代わりにスキップします (約 60 KB → 1 つの 4 KB チャンク)。
fetch_json は値が見つかった瞬間に停止し、生のバイトスキャンを実行します。
JSON DOM ツリーを置き換えます。
静的コンテンツは &'static str (ゼロ ヒープ) としてフラッシュ内に存在し、次のようにストリーミングされます。
512 バイトの TCP セグメント - 現在のセグメントのみが RAM に存在するため、4.3 KB のエージェント
IDE はフルページ バッファなしで機能します。
ヒープとスタックは 1 つの DRAM プール ( +16 kB ヒープ = −16 kB スタック) を共有し、手動で調整されます。
SO_LINGER=0 は、lwIP ソケットを直ちに解放します (TIME_WAIT パイルアップは発生しません)。
無駄のないビルドでの協調的な共有スタックにより、スレッドごとの並列処理が安価になります
スタックではできません。
完全なカタログ (正確なファイルとシンボルを含むすべてのトリック) は、
docs/memory-512kb.md 。
デフォルトで拒否する文法。ホワイトリストに登録された式ノードと
上記の組み込みは存在します。ループ、関数、再帰、属性アクセス、インポート、文字列、または I/O はありません。
指導予算。各評価にはステップ制限（暴走ガード）があります。
壁時計のタイムアウト。呼び出し元は期限 (例: 10 秒) を指定できます。デバイス上で
ESP タイマーがそれを強制します。
パーサーは小さいままです。オンデバイスパーサーは安全なサブセットのみを受け入れるため、
デバイスの攻撃対象領域は、単なる制限されたデコーダーとエバリュエーターです。

## Original Extract

Why learn to spell Python when PySpell can conjure it for you?
Say what you want in plain English and a tiny on-device model turns it into runnable Python —
100% locally on a $10 ESP32 with 512 kB RAM (no PSRAM, no cloud, no API key),
reachable anywhere over Tailscale .
View on GitHub → · How it works ↓
Try it live — the model runs in your browser
No install, no device, no cloud. The same ~0.5 MB model + sandbox that runs on the ESP32, loaded here as WebAssembly. Type plain English; it writes Python and runs it.
screen + LED keep the last result · Reset
Runs 100% in your browser. Want it reachable from your phone ? Use Connect Tailscale below to turn this tab into a private node on your own tailnet. How it works →
Opt-in — nothing connects until you click. Our Rust Tailscale client (~321 kB WASM) joins your tailnet over WebSocket; node keys stay in this browser. Then open http://<node-ip>/ from any device on your tailnet to get this same demo, over the tunnel.
The chip serves a web agent IDE, a native MCP server and a REST API ,
runs the code in a sandbox (with live web-request / fetch support), and drives
its own screen + LED — up to 8 parallel PySpell processes on that same half-megabyte of RAM.
Like MicroPython, but two syntaxes, the parser never ships to the device, and English is the third syntax.
A PySpell program is a single expression (Python) or some let bindings followed by a
trailing expression (Rust). It evaluates to a value — a number, a boolean, a string, or a list. Free
identifiers are resolved at evaluation time against a host-supplied environment : CLI variables
on a laptop, or live device readings on a microcontroller. The only I/O is a host-granted, allowlisted
fetch_json ; there are no loops, functions, or imports — that is the point: small, fast, and
safe to accept from elsewhere.
An offline AI coding agent, served off the chip
Open http://<dongle>/ over the tunnel and you get a Cursor-like agent. Type
"flash the light" , "show the text "hello"" , "what is 7 plus 5" ,
or "reverse the word robot" — a ~0.45 M-parameter language model (< 500 kB, int8)
turns it into PySpell code, runs it live on the chip , and shows the result, or the
physical action (the screen lights up, the RGB LED blinks). Runtime, model, tokenizer and dictionary
are all served from the dongle, offline — no cloud, no key (OpenAI is optional, behind
the ⚙).
A model that small is only useful because of a chain of tricks — the full write-up is in
tech.md . The headlines:
The model points, the browser copies
A 0.45 M model can't reliably copy arbitrary tokens (numbers, strings, lists), so it isn't asked
to. It emits tiny semantic directives; the browser copies the literal content verbatim.
calculate 3 + 2 → print( 3 + 2 ) ; change add to
subtract → @@ + ==> - . Quoted text is literal content — copied byte-for-byte,
excluded from vocab checks.
The device serves; the browser computes
Inference runs in WebAssembly, client-side. The 0.5 MB model image streams off flash a
TCP segment at a time (HTTP Range) and is never resident in the chip's ~60 kB heap. Inverted
edge inference: the constrained device serves and grades, the browser runs the model.
The 512-token vocab is embedded with all-MiniLM (22 M params), PCA'd to 128 dims, folded with a
part-of-speech vector, and frozen — the tiny model starts with meaningful word
geometry instead of spending its tiny budget learning it.
The vocabulary is the dictionary
Those same 512 tokens + embeddings are served back to the browser for input validation
("outside the model's vocabulary…") and related-word RAG over the model's own vocabulary.
free_heap > 100000 and uptime_s < 60
250 if distance > 1000 else 0
0 < temp < 60 # chained
20 not in peers
sum([1, 2, 3])
readings[-1] # negative index
max(a, b)
Rust
free_heap > 100000 && uptime_s < 60
if distance > 1000 { 250 } else { 0 }
let used = total - free; used * 100 / total
!peers.contains(20)
sum([1, 2, 3])
readings[readings.len() - 1]
max(a, b)
Language reference
Classic one-liner — fetch a value and show it on the dongle's screen:
show("Oslo: " + fetch_json(
"https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=59.91&lon=10.75",
"properties.timeseries.0.data.instant.details.air_temperature") + " C")
# screen shows: Oslo: 14.9 C (and the call returns that string)
Network & JSON
fetch(url) + json_get(text, path) let a program pull live data and read one
field out of it. fetch is a mediated capability — the host/device decides which hosts are
reachable (an allowlist), so a program can't reach arbitrary URLs.
# Host CLI (allow the host explicitly):
pyspell run oslo_temp.py --allow-host api.met.no
# where oslo_temp.py is:
json_get(
fetch("https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=59.91&lon=10.75"),
"properties.timeseries.0.data.instant.details.air_temperature")
# → 14.9
Memory note (device): json_get is path-directed so it never builds the
whole document in RAM — it materializes only the matched value. On the ESP32 (≈60 kB free, no PSRAM)
reading a field out of a large response is feasible because fetch_json streams the
HTTP(S) body and stops the moment the field is found (freeing the TLS buffers early) — so a ~50 kB yr.no
response never has to fit in RAM at once.
# On the ESP32, over Tailscale (single process; ≈60 kB free; verified live):
fetch_json(
"https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=59.91&lon=10.75",
"properties.timeseries.0.data.instant.details.air_temperature")
# → 14.9 (the dongle fetched yr.no itself)
Running on the host
# Evaluate, binding free variables:
cargo run -p pyspell-cli -- run examples/health.py --set free_heap=120000 --set uptime_ms=45000
# → true
# Compile to a portable IR blob:
cargo run -p pyspell-cli -- compile examples/health.py # → examples/health.py.psb
# Push live to a device over USB-serial, or an interactive REPL:
cargo run -p pyspell-cli -- repl --port /dev/cu.usbmodem2101 --lang python
Running on the ESP32
The portable evaluator ( pyspell-core , no_std + alloc ) runs unchanged on the
ESP32-S3. Programs read live device variables from the environment:
The demo/esp32-tailscale-pyspell firmware adds a web text window and a /run
API inside a Tailscale tunnel — open the device's Tailscale IP in a browser, type an
expression, set a timeout, and run it on the chip. PySpell adds only ~62 kB on top of the networking
firmware.
# Web window:
open http://100.x.y.z/
# POST (preferred): program in the body, lang/timeout in the query.
# More room for code than a URL, and no percent-encoding.
curl -X POST 'http://100.x.y.z/run?lang=py&timeout=10' --data 'free_heap > 100000' # → true
curl -X POST 'http://100.x.y.z/run?lang=rs&timeout=10' --data 'uptime_ms / 1000' # → 22
# GET (also supported): code is URL-encoded in the query.
curl 'http://100.x.y.z/run?lang=py&timeout=10&code=free_heap%20%3E%20100000' # → true
timeout is in seconds, clamped to 1–60, and enforced as a real wall-clock deadline on
the device. The single request must fit one TCP segment (≈1.2 kB) — POST leaves more of that for code.
The reply is text/plain (no JSON wrapper):
The ESP32-S3 has 512 kB of SRAM and no PSRAM , yet it runs a full Tailscale node
(control plane and DERP), the PySpell evaluator, a browser agent IDE served off the chip, a
native MCP server, and TLS to api.met.no. That only fits because of a long chain of memory tricks.
SPKI leaf-key pinning instead of CA-chain validation — one RSA-PSS verify, no 6 kB
chain buffer (a TLS fetch drops ~45→30 kB). A heap admission gate bounds concurrency
so peak heap is K × per-fetch , never N × per-fetch .
The netmap is read with serde_json::from_reader over the HTTP/2 frames, so serde
skips the huge DERPMap field instead of buffering it (~60 kB → one 4 kB chunk).
fetch_json stops the moment the value is found, and raw byte-scans
replace JSON DOM trees.
Static content lives in flash as &'static str (zero heap) and is streamed out as
512-byte TCP segments — only the current segment is ever in RAM, so the 4.3 kB agent
IDE serves without a full-page buffer.
Heap and stack share one DRAM pool ( +16 kB heap = −16 kB stack ), tuned by hand.
SO_LINGER=0 frees lwIP sockets immediately (no TIME_WAIT pile-up), and a
cooperative shared stack on the lean build makes parallelism cheap where per-thread
stacks can't.
The full catalog — every trick with the exact file and symbol — is in
docs/memory-512kb.md .
Deny-by-default grammar. Only the whitelisted expression nodes and the
built-ins above exist — no loops, functions, recursion, attribute access, imports, strings, or I/O.
Instruction budget. Every evaluation has a step limit (runaway guard).
Wall-clock timeout. A caller can supply a deadline (e.g. 10 s); on the device the
ESP timer enforces it.
Parser stays small. The on-device parser accepts only the safe subset, so the
device's attack surface is just a bounded decoder + evaluator.
