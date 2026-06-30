---
source: "https://loadr.io"
hn_url: "https://news.ycombinator.com/item?id=48730559"
title: "Loadr – a Rust load-testing tool that began as a single Claude prompt"
article_title: "loadr — the load testing platform. k6, JMeter, Gatling & Locust ideas in one Rust binary."
author: "reaandrew"
captured_at: "2026-06-30T11:01:28Z"
capture_tool: "hn-digest"
hn_id: 48730559
score: 1
comments: 0
posted_at: "2026-06-30T10:05:57Z"
tags:
  - hacker-news
  - translated
---

# Loadr – a Rust load-testing tool that began as a single Claude prompt

- HN: [48730559](https://news.ycombinator.com/item?id=48730559)
- Source: [loadr.io](https://loadr.io)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T10:05:57Z

## Translation

タイトル: Loadr – 単一のクロード プロンプトとして開始された Rust 負荷テスト ツール
記事のタイトル:loadr — 負荷テスト プラットフォーム。 k6、JMeter、Gatling、Locust のアイデアを 1 つの Rust バイナリにまとめました。
説明:loadr は、最新の負荷テスト プラットフォームです。宣言型 YAML テスト、埋め込み JavaScript、HTTP/2、WebSocket、gRPC、GraphQL、TCP および UDP、WASM プラグイン、正確なパーセンタイルを持つ分散エージェント、および組み込みのライブ Web UI を備えています。バイナリが 1 つ。

記事本文:
ランドマーク (各ページのフッターの前で閉じられます)。
-->
メインコンテンツにスキップ
ローダー.io
ベータ版
製品
特長
プロトコル
ウェブUI
分散型
k6 / JMeter との比較
ロードマップ
デモ
プラグイン
ドキュメント
GitHub
ダウンロード
デモ
プラグイン
ドキュメント
製品
loadr は、k6、JMeter、
Gatling と Locust を 1 か所に: 宣言型 YAML テスト、埋め込み JavaScript、6 つのプロトコル、プラグイン、
組み込みのライブ Web UI、および分散実行
数学的に正確なパーセンタイル。オールインワンのバイナリ。
ライブデモ — 1 シナリオ、8.1 秒
小切手............................: 100.00% — ✓ 752 ✗ 0
✓ ステータスは 200 (376 / 376)
✓ 200ms未満 (376 / 376)
http_req_duration................................: avg=4.96ms med=3.02ms p(95)=10.73ms p(99)=49.38ms
http_req_tls_handshaking....: avg=0µs (VU ごとに再利用される接続)
http_req_waiting....................: avg=4.85ms med=2.90ms p(95)=10.57ms
http_reqs................................: 376 (46.66/秒)
反復回数....................................: 376 (46.66/秒)
vus....................: 値=5 最小=4 最大=5
しきい値:
✓ http_req_duration: p(95)<500 (実測値: 10.73)
✓ チェック: レート > 0.95 (実測値: 1.00)
✓ http_req_failed: レート<0.01 (実測値: 0.00)
7
ロードエグゼキュータ
6
組み込まれたプロトコル
1
バイナリ、ゼロデプス
p99.9
HDR に正確、フリート全体
なぜローダーなのか
4つのツールの最高のアイデア。 Rust バイナリ 1 つ。
感謝の意を込めて、loadr は、負荷テストを形成した 4 つのプロジェクト、つまり k6 の実行モデル、JMeter のリクエスト アーセナル、Gatling のフロー制御 DSL とフィーダー、および Locust の動作モデルを基にして、Rust にまとめて再実装しました。私たちが構築したものをご覧ください →
オープン モデルの到着率エグゼキュータは、
システムの速度が低下します - 飽和は Dropped_iterations として表示されます。
静かにスループット (1 秒あたりのリクエスト、RPS) を低下させるわけではありません。すべての遅延は HDR の歴史です

グラム: p(99.9) は正確であり、決して推定されず、平均化もされません。
生成された JSON スキーマを含む宣言型 YAML — エディターがオートコンプリートします。
loadr validate は行番号を付けて lint を実行し、
修正を意味しており、PR の差分には実際に何か意味があります。
ロジックが必要とする箇所に正確に JavaScript を導入します。
組み込みの管理 Web UI (負荷テスト用の RabbitMQ を考えてください)、分散型
mTLS gRPC 経由のコントローラー/エージェント モード、6 つのメトリック エクスポーター、WASM およびネイティブ プラグイン、
既存の .jmx プランと k6 スクリプトを使用するインポーター。
本物の録音。本物が走る。モックアップはありません。
すべてのクリップは、ライブサーバーに対して実行される実際のローダーバイナリです。
ライブ チャート、しきい値ピル、実行コントロール、エディターおよびフリート ビュー - 実際のブラウザー セッション。
分散型フリート、ライブ 0:34
リストされている場合は、バイナリで出荷され、すべての項目がそのドキュメントにリンクされています。
constant-vus 、 Ramping-vus 、 constant-delivery-rate 、 Ramping-delivery-rate 、 per-vu-iterations 、shared-iterations 、 externally-controlled — 同一のセマンティクス、オープン モデルとクローズド モデル。
独立したエグゼキューター、ステージ、開始時間、正常な停止とランプダウンを備えたテストごとに任意の数の名前付きシナリオ (ブラウザー、API クライアント、バッチ ジョブを 1 回の実行で実行)。
p(95)<400 、 rate>0.99 、任意のパーセンタイル、タグでフィルター処理されたセレクター、ウォームアップ遅延のある abort_on_fail サーキット ブレーカー。失敗時の終了コード 99 — k6 互換。
CI ネイティブ: GitHub アクション + JUnit
ファーストパーティの setup-loadr / run アクションは CLI をインストールしてプランを実行し、 --junit は CI テスト パネルがレンダリングするすべての JUnit レポート (GitHub、GitLab、Jenkins、CircleCI) を出力します。
リクエストを決して失敗しない k6 スタイルのチェックと、ステータス、本文に含まれる/正規表現、JSONPath、XPath、期間、サイズ、ヘッダー、JS 式を実行する JMeter スタイルのアサーション (中止アクション付き)。
JSONPath、正規表現キャプチャ グループ、XPath 1.0

、CSS セレクター、境界抽出機能、およびヘッダー - 抽出された値は、後のリクエストに ${name} として流れ込み、JS に流れ込みます。
シーケンシャル、ランダムまたはシャッフル戦略 (共有または仮想ユーザー (VU) ごと、リサイクルまたは EOF で停止)、インライン行、およびログに到達しないシークレットを備えた CSV および JSON フィーダー。
repeat 、 while 、 if / else および加重 / 均一 / ラウンドロビン ランダム分岐 — 宣言型 YAML の Gatling のループとスイッチ、および Locust の加重タスク モデル。
ターゲットがどれほど高速であっても、既知のレート制限を下回るようにするための、任意のエグゼキューター (Gatling のリーチRps) の上にあるグローバルなリクエスト レートの上限 ( throttle: {requests_per_second } )。
k6 スタイルのインポート ( k6/http 、 check 、 sleep 、 metrics)、 setup() / teadown() 、シナリオ関数、 beforeRequest / afterRequest フック、インライン ${js: …} を備​​えた VU ごとの QuickJS — 時間とメモリ制限のあるサンドボックス化。
一定、一様ランダム、ガウスの思考時間。一定のスループット ペーシング (JMeter タイマー、適切に実行)。シナリオごとのデフォルトとグローバルなデフォルト。
DNS、接続、TLS ハンドシェイク、送信、TTFB、受信は、手作りのハイパー スタック上でリクエストごとに測定され、さらに正確なワイヤ レベルのバイト数も加算されます。平均的な推測はありません。
手動オーバーライドを備えた自動 VU ごとの RFC 6265 Cookie jar、リダイレクト ポリシー、JSON/form/multipart/file body、クエリ パラメータ - すべてが補間されます。
カスタム CA、クライアント証明書、SNI オーバーライド、ステージング用の安全でないモード、HTTP/HTTPS プロキシ (CONNECT)、バージョン強制による ALPN ネゴシエーション - すべて Rustl、OpenSSL なし。
1 つのファイル、多数のターゲット:loadr run -e staging 名前付きオーバーレイのディープマージ — より穏やかな CI ロード、ステージング URL、緩和されたしきい値、コピーアンドペーストなし。
JSON ライン、CSV、Prometheus (スクレイピング + リモート書き込み)、InfluxDB ライン プロトコル、OpenTelemetry OTLP (gRPC および HTTP)、StatsD — 加えて、事前に構築された Grafana ダッシュボード

ポ。
WIT インターフェースを備えたサンドボックス WASM コンポーネントと abi_stable ネイティブ ライブラリという 2 つのメカニズムにわたる 5 つのプラグイン タイプ (プロトコル、出力、エクストラクター、アサーション、サービス)。再構築は一度もありません。
オプションの mTLS を使用した gRPC 経由のコントローラー + エージェント: ロード パーティショニング、同期開始、データ ファイル送信、ハートビート、再接続、エージェント損失ポリシー、および中央 HDR マージ。
SSE 上のライブ ダッシュボード、ワンクリック検証を備えたテスト エディター、実行履歴、一時停止/停止/スケール コントロール、エージェント フリート ビュー、ログ テール - バイナリに埋め込まれたダーク モード ネイティブ。
loadr Convert plan.jmx は、スレッド グループ、サンプラー、タイマー、アサーション、エクストラクター、CSV 構成を変換します。 k6 インポーターは、オプション、シナリオ、チェック、http 呼び出しをマップし、残りについては明確な警告を出します。
k6 スタイルのコンソール概要、JSON エクスポート、自己完結型 HTML レポート (loadr レポート)、シェル補完、JSON スキーマ出力、構造化ログ、--quiet / -v spectrum。
スモークテストからフリート規模までを同じファイルで実行
名前: ロード中のチェックアウト
デフォルト:
http: { ベース URL: https://shop.example.com }
データ:
ユーザー: { タイプ: csv、パス: users.csv、モード: 共有、on_eof: recycle }
シナリオ:
買い物客:
エグゼキュータ: ランピング対
ステージ: [ { 持続時間: 2 メートル、目標: 100 }、 { 持続時間: 5 メートル、目標: 100 } ]
think_time: { タイプ: 均一、最小: 1 秒、最大: 3 秒 }
流れ:
- リクエスト:
URL: /ログイン
メソッド: POST
本文: { フォーム: { ユーザー: "${data.users.username}"、パス: "${data.users.password}" } }
抜粋：
- { タイプ: css、名前: csrf、式: "input[name=csrf]"、属性: 値 }
主張します:
- { タイプ: ステータス、等しい: 200 }
- リクエスト:
メソッド: POST
URL: /カート
本文: { フォーム: { sku: W-1、csrf: "${csrf}" } }
チェック:
- { タイプ: ステータス、等しい: 201 }
- { タイプ: 期間、名前: 高速チェックアウト、最大: 300ms }
しきい値:
http_req_duration: [ "p(95)<400", "p(99.9)<1500" ]
http_req_f

失敗: [ { しきい値: "rate<0.01"、abort_on_fail: true } ]
チェック: [ "レート>0.99" ]
script.js k6互換感・QuickJSサンドボックス
「k6/http」から http をインポートします。
import { チェック、スリープ、グループ } from 'k6';
'k6/metrics' から { トレンド } をインポートします。
const apiLatency = new Trend('api_latency');
エクスポート関数 setup() {
const res = http.post('/auth/token', JSON.stringify({ id: __ENV.CLIENT_ID }));
return { トークン: res.json().token };
}
デフォルト関数 (データ) をエクスポート {
group('カタログ', () => {
const res = http.get('/items?limit=20', {
ヘッダー: { 認可: `Bearer ${data.token}` }、
});
apiLatency.add(res.duration_ms);
check(res, {
'ステータス 200': (r) => r.status === 200、
'アイテムあり': (r) => r.json().items.length > 0、
});
});
sleep(Math.random() * 2 + 0.5);
}
// すべての YAML リクエスト ステップを中心に起動します
エクスポート関数 beforeRequest(req) {
req.headers['X-Request-Id'] = crypto.uuidv4();
要求を返す;
}
すべてのエージェントにわたるフリートの正確なパーセンタイル
# コントロール プレーン (Web UI も提供)
$loadrコントローラー --bind 0.0.0.0:7625 --ui-bind 0.0.0.0:6464
# すべてのロード ジェネレーターで
$loadragent --join ctrl-host:7625 --nameagent-$(ホスト名)
# submit: 600 リクエスト/秒をフリート全体に分割
$ロア
[切り捨てられた]
6 つのプロトコル。それぞれの完全なメトリクス。
すべてのプロトコルは、DNS、接続、TLS、TTFB、期間、送受信バイト数を報告し、同じ抽出/アサート/チェック ブロックで動作します。
- リクエスト:
メソッド: POST
URL: /注文
本文: { json: { sku: W-1、数量: 2 } }
チェック: [ { タイプ: ステータス、等しい: 201 } ]
ALPN、事前知識 h2、キープアライブ チューニング、VU ごとのプール。
- リクエスト:
URL: wss://chat.example.com/ws
WS:
送信: [ '{"タイプ":"hello"}' ]
accept_until: '"ack"'
セッション期間: 10秒
サブプロトコル、バイナリ フレーム、メッセージ カウンタ、セッション メトリック。
- リクエスト:
URL: grpc://svc:50051
グループ:
リフレクション: true # または proto

_ファイル
サービス: helloworld.Greeter
メソッド: SayHello
メッセージ: { 名前: "vu-${vu}" }
単項 + すべてのストリーミング シェイプ、インプロセス プロト コンパイル — プロトックなし。
- リクエスト:
URL: /graphql
プロトコル: グラフQL
グラフキュール:
クエリ: "クエリ($t:String!){ 検索(t:$t){ id } }"
変数: { t: ウィジェット }
GraphQL エラー セマンティクス、部分エラー認識、独自のメトリック ファミリ。
- リクエスト:
URL: tcp://ゲートウェイ:7000
ソケット:
send_text: "PING\r\n"
読み取りバイト数: 64
チェック: [ { タイプ: body_contains、値: PONG } ]
正確なバイトアカウンティング。生のペイロードに対する正規表現/境界抽出。
- リクエスト:
URL: udp://stats:8125
ソケット:
send_hex: "デッドビーフ 0102"
読み取りタイムアウト: 500ms
16 進ペイロードと損失認識タイムアウトによるデータグラム ラウンド トリップ。
MQTT、Kafka、または社内プロトコルが必要ですか?プロトコル プラグインを作成します。フォークやリビルドは必要ありません。
ほとんどのツールはノード全体でパーセンタイルを平均します。その数字は間違っています。
エージェント A の p99 が 100 ミリ秒、エージェント B の p99 が 1000 ミリ秒の場合、フリートの実際の p99 は次のようになります。
550ミリ秒ではありません。ローダー エージェント ストリーム HDR ヒストグラム デルタ
毎秒。コントローラーはヒストグラムをマージします - ロスレス
操作 — マージ後にのみパーセンタイルを計算します。しきい値を一元的に評価
艦隊全体の真実に反する。
▸ VU 数と到着率はエージェント間で正確に分割され、グローバルランプは正確なままです
▸ 同期開始バリア、ハートビート、ジッター再接続、エージェント喪失ポリシー
▸ テスト定義と CSV/プロト/JS ファイルがエージェントに自動的に送信されます
▸ エージェントごとに 1 つの双方向 gRPC ストリーム、プレーンテキストまたは mTLS
▸ リポジトリ内の Docker Compose スタックと Helm チャート: Agents.replicas=10 に移動します。
RabbitMQ スタイルの管理、負荷テスト用
バイナリに埋め込まれています。単一の実行の場合は、loadr run --ui、またはコントローラー上の完全なフリート コンソール。ブラウザーでテストを編集および検証し、ライブ パーセンタイルを監視し、一時停止、停止、またはオンにします。

走行中に VU をダイヤルします。
ローダー · 概要
ライブ · 実行 4588f1bf
リクエスト/秒
612.4
アクティブな VU
300
p95 レイテンシ
187ミリ秒
エラー率
0.02%
✓ p(95)<400 — 187.2
✓ レート<0.01 — 0.0002
✓ チェックレート>0.99 — 0.998
⏸一時停止
■停止
テスト ライブラリとエディター ブラウザーで YAML を保存、編集、検証します。診断はその行にジャンプします。
すべての実行の完全なサマリー (傾向、チェック、しきい値、合格/不合格) が保持されます。
エージェントの健全性、アクティブな VU、コア、ラベル、最後のハートビートが一目でわかります。
HTTP Basic およびベアラー トークン。デフォルトではループバックのみ。あらゆるものに対応する JSON API。
k6、JMeter、Gatling、Locust の隣のローダー
4 つはすべて、このスペースを形成した、広く愛されている優れたツールです。これは単に、それらに対してローダーが位置する場所にすぎません。あなたのチームに合ったものを選んでください。セルに [Enterprise] または [cloud] と表示されている場合、その機能はそのプロジェクトの有料/ホスト層に存在します。ローダーが何に基づいて構築されているかも参照してください。
loadr は新しい Rust 実装であり、何かのフォークではありませんが、負荷テストを定義した 4 つのツールから最高のアイデアを意図的に、そして感謝の気持ちを込めて取り入れています。
実行モデルは k6 です。7 つのエグゼキューター、オープン/クローズド ロード、4 つのメトリック タイプ、abortOnFail と終了コード 99 のしきい値、チェックとグループ、および埋め込み JS エクスペリエンスがあるため、API はインポート互換です ( import http from 'k6/http' )。
JMeter がリクエストを整形

[切り捨てられた]

## Original Extract

loadr is a modern load testing platform: declarative YAML tests, embedded JavaScript, HTTP/2, WebSocket, gRPC, GraphQL, TCP & UDP, WASM plugins, distributed agents with exact percentiles, and a built-in live web UI. One binary.

landmark (closed by each page before its footer).
-->
Skip to main content
loadr .io
Beta
Product
Features
Protocols
Web UI
Distributed
Compare vs k6 / JMeter
Roadmap
Demos
Plugins
Docs
GitHub
Download
Demos
Plugins
Docs
Product
loadr brings together the best ideas from k6, JMeter,
Gatling and Locust in one place: declarative YAML tests, embedded JavaScript, six protocols, plugins,
a built-in live web UI, and distributed execution with
mathematically exact percentiles . All in one binary.
live-demo — 1 scenario(s), 8.1s
checks........................: 100.00% — ✓ 752 ✗ 0
✓ status is 200 (376 / 376)
✓ under 200ms (376 / 376)
http_req_duration.............: avg=4.96ms med=3.02ms p(95)=10.73ms p(99)=49.38ms
http_req_tls_handshaking......: avg=0µs (connections reused per VU)
http_req_waiting..............: avg=4.85ms med=2.90ms p(95)=10.57ms
http_reqs.....................: 376 (46.66/s)
iterations....................: 376 (46.66/s)
vus...........................: value=5 min=4 max=5
thresholds:
✓ http_req_duration: p(95)<500 (observed: 10.73)
✓ checks: rate>0.95 (observed: 1.00)
✓ http_req_failed: rate<0.01 (observed: 0.00)
7
load executors
6
protocols built in
1
binary, zero deps
p99.9
HDR-exact, fleet-wide
Why loadr
Four tools' best ideas. One Rust binary.
loadr draws, with thanks, on four projects that shaped load testing: k6's execution model, JMeter's request arsenal, Gatling's flow-control DSL and feeders, and Locust's behaviour model — brought together and reimplemented in Rust. See what we built on →
Open-model arrival-rate executors keep the offered load constant even when your
system slows down — saturation shows up as dropped_iterations ,
not silently lower throughput (requests per second, RPS). Every latency is an HDR histogram: p(99.9) is exact, never estimated, never averaged.
Declarative YAML with a generated JSON Schema — your editor autocompletes it,
loadr validate lints it with line numbers and
did-you-mean fixes, and the diff in your PR actually means something.
Drop into JavaScript exactly where logic demands it.
Built-in management web UI (think RabbitMQ for load tests), distributed
controller/agent mode over mTLS gRPC, six metric exporters, WASM & native plugins,
and importers that eat your existing .jmx plans and k6 scripts.
Real recordings. Real runs. No mockups.
Every clip is the actual loadr binary executing against a live server.
Live charts, threshold pills, run controls, editor and fleet view — a real browser session.
A distributed fleet, live 0:34
If it's listed, it ships in the binary — and every item links to its documentation.
constant-vus , ramping-vus , constant-arrival-rate , ramping-arrival-rate , per-vu-iterations , shared-iterations , externally-controlled — identical semantics, open and closed models.
Any number of named scenarios per test with independent executors, stages, start times, graceful stop and ramp-down — browsers, API clients and batch jobs in one run.
p(95)<400 , rate>0.99 , any percentile, tag-filtered selectors, abort_on_fail circuit breakers with warm-up delay. Exit code 99 on failure — k6-compatible.
CI-native: GitHub Action + JUnit
A first-party setup-loadr / run action installs the CLI and runs your plan, and --junit emits a JUnit report every CI test panel renders — GitHub, GitLab, Jenkins, CircleCI.
k6-style checks that never fail requests, and JMeter-style assertions that do: status, body contains/regex, JSONPath, XPath, duration, size, headers, JS expressions — with abort actions.
JSONPath, regex capture groups, XPath 1.0, CSS selectors, boundary extractors and headers — extracted values flow into later requests as ${name} and into JS.
CSV & JSON feeders with sequential, random or shuffle strategies (shared or per virtual user (VU), recycle or stop-at-EOF), inline rows, and secrets that never reach logs.
repeat , while , if / else and weighted / uniform / round-robin random branches — Gatling's loops and switches and Locust's weighted-task model, in declarative YAML.
A global request-rate ceiling ( throttle: { requests_per_second } ) on top of any executor — Gatling's reachRps , for staying under a known rate limit no matter how fast the target is.
QuickJS per VU with k6-style imports ( k6/http , check , sleep , metrics), setup() / teardown() , scenario functions, beforeRequest / afterRequest hooks, inline ${js: …} — sandboxed with time & memory limits.
Constant, uniform-random and gaussian think time; constant-throughput pacing (the JMeter timer, done right); per-scenario and global defaults.
DNS, connect, TLS handshake, send, TTFB and receive measured per request on a hand-built hyper stack — plus exact wire-level byte counts. No averaged guesses.
Automatic per-VU RFC 6265 cookie jars with manual override, redirect policies, JSON/form/multipart/file bodies, query params — everything interpolated.
Custom CAs, client certificates, SNI override, insecure mode for staging, HTTP/HTTPS proxies (CONNECT), ALPN negotiation with version forcing — all rustls, no OpenSSL.
One file, many targets: loadr run -e staging deep-merges named overlays — gentler CI load, staging URLs, relaxed thresholds, without copy-paste.
JSON lines, CSV, Prometheus (scrape + remote-write), InfluxDB line protocol, OpenTelemetry OTLP (gRPC & HTTP), StatsD — plus a pre-built Grafana dashboard in the repo.
Five plugin types (protocol, output, extractor, assertion, service) over two mechanisms: sandboxed WASM components with a WIT interface, and abi_stable native libraries. No rebuilds, ever.
Controller + agents over gRPC with optional mTLS: load partitioning, synchronized starts, data-file shipping, heartbeats, reconnection, agent-loss policies — and central HDR merging.
Live dashboards over SSE, a test editor with one-click validation, run history, pause/stop/scale controls, agent fleet view, log tail — embedded in the binary, dark mode native.
loadr convert plan.jmx translates thread groups, samplers, timers, assertions, extractors and CSV configs; the k6 importer maps options, scenarios, checks and http calls — with clear warnings for the rest.
k6-style console summaries, JSON export, self-contained HTML reports ( loadr report ), shell completions, JSON Schema output, structured logs, --quiet / -v spectrum.
From smoke test to fleet-scale in the same file
name: checkout-under-load
defaults:
http: { base_url: https://shop.example.com }
data:
users: { type: csv, path: users.csv, mode: shared, on_eof: recycle }
scenarios:
shoppers:
executor: ramping-vus
stages: [ { duration: 2m, target: 100 }, { duration: 5m, target: 100 } ]
think_time: { type: uniform, min: 1s, max: 3s }
flow:
- request:
url: /login
method: POST
body: { form: { user: "${data.users.username}", pass: "${data.users.password}" } }
extract:
- { type: css, name: csrf, expression: "input[name=csrf]", attribute: value }
assert:
- { type: status, equals: 200 }
- request:
method: POST
url: /cart
body: { form: { sku: W-1, csrf: "${csrf}" } }
checks:
- { type: status, equals: 201 }
- { type: duration, name: fast checkout, max: 300ms }
thresholds:
http_req_duration: [ "p(95)<400", "p(99.9)<1500" ]
http_req_failed: [ { threshold: "rate<0.01", abort_on_fail: true } ]
checks: [ "rate>0.99" ]
script.js k6-compatible feel · QuickJS sandbox
import http from 'k6/http';
import { check, sleep, group } from 'k6';
import { Trend } from 'k6/metrics';
const apiLatency = new Trend('api_latency');
export function setup() {
const res = http.post('/auth/token', JSON.stringify({ id: __ENV.CLIENT_ID }));
return { token: res.json().token };
}
export default function (data) {
group('catalogue', () => {
const res = http.get('/items?limit=20', {
headers: { Authorization: `Bearer ${data.token}` },
});
apiLatency.add(res.duration_ms);
check(res, {
'status 200': (r) => r.status === 200,
'has items': (r) => r.json().items.length > 0,
});
});
sleep(Math.random() * 2 + 0.5);
}
// fires around every YAML request step
export function beforeRequest(req) {
req.headers['X-Request-Id'] = crypto.uuidv4();
return req;
}
fleet exact percentiles across every agent
# control plane (serves the web UI too)
$ loadr controller --bind 0.0.0.0:7625 --ui-bind 0.0.0.0:6464
# on every load generator
$ loadr agent --join ctrl-host:7625 --name agent-$(hostname)
# submit: 600 req/s split across the whole fleet
$ loa
[truncated]
Six protocols. Full metrics on every one.
Every protocol reports DNS, connect, TLS, TTFB, duration, and bytes sent/received — and works with the same extract/assert/check blocks.
- request:
method: POST
url: /orders
body: { json: { sku: W-1, qty: 2 } }
checks: [ { type: status, equals: 201 } ]
ALPN, prior-knowledge h2, keep-alive tuning, per-VU pools.
- request:
url: wss://chat.example.com/ws
ws:
send: [ '{"type":"hello"}' ]
receive_until: '"ack"'
session_duration: 10s
Subprotocols, binary frames, message counters, session metrics.
- request:
url: grpc://svc:50051
grpc:
reflection: true # or proto_files
service: helloworld.Greeter
method: SayHello
message: { name: "vu-${vu}" }
Unary + all streaming shapes, in-process proto compile — no protoc.
- request:
url: /graphql
protocol: graphql
graphql:
query: "query($t:String!){ search(t:$t){ id } }"
variables: { t: widget }
GraphQL error semantics, partial-error awareness, own metric family.
- request:
url: tcp://gateway:7000
socket:
send_text: "PING\r\n"
read_bytes: 64
checks: [ { type: body_contains, value: PONG } ]
Exact byte accounting; regex/boundary extraction over raw payloads.
- request:
url: udp://stats:8125
socket:
send_hex: "deadbeef 0102"
read_timeout: 500ms
Datagram round trips with hex payloads and loss-aware timeouts.
Need MQTT, Kafka, or your in-house protocol? Write a protocol plugin — no fork, no rebuild.
Most tools average percentiles across nodes. That number is wrong.
If agent A's p99 is 100 ms and agent B's is 1000 ms, the fleet's true p99 is
not 550 ms. loadr agents stream HDR histogram deltas
every second; the controller merges the histograms — a lossless
operation — and computes percentiles only after the merge. Thresholds evaluate centrally
against fleet-wide truth.
▸ VU counts and arrival rates partitioned exactly across agents — global ramps stay precise
▸ Synchronized start barrier, heartbeats, jittered reconnection, agent-loss policies
▸ Test definitions and CSV/proto/JS files shipped to agents automatically
▸ One bidirectional gRPC stream per agent, plaintext or mTLS
▸ Docker Compose stack and Helm chart in the repo: agents.replicas=10 and go
RabbitMQ-style management, for load tests
Embedded in the binary. loadr run --ui for a single run, or the full fleet console on the controller. Edit and validate tests in the browser, watch live percentiles, pause, stop, or turn the VU dial mid-run.
loadr · overview
live · run 4588f1bf
Requests / s
612.4
Active VUs
300
p95 latency
187 ms
Error rate
0.02%
✓ p(95)<400 — 187.2
✓ rate<0.01 — 0.0002
✓ checks rate>0.99 — 0.998
⏸ Pause
■ Stop
Test library & editor Save, edit and validate YAML in the browser — diagnostics jump to the line.
Every run's full summary persisted: trends, checks, thresholds, pass/fail.
Agent health, active VUs, cores, labels, last heartbeat — at a glance.
HTTP Basic and bearer tokens; loopback-only by default. JSON API for everything.
loadr next to k6, JMeter, Gatling & Locust
All four are excellent, widely-loved tools that shaped this space — this is simply where loadr sits relative to them; pick whatever fits your team. Where a cell says Enterprise or cloud , the capability exists in that project's paid/hosted tier. See also what loadr is built on .
loadr is a fresh Rust implementation — not a fork of anything — but it takes the best ideas from four tools that defined load testing, deliberately and with thanks.
The execution model is k6's: the seven executors, open/closed load, the four metric types, thresholds with abortOnFail and exit code 99, checks and groups — and an embedded-JS experience so close the API is import-compatible ( import http from 'k6/http' ).
JMeter shaped the reques

[truncated]
