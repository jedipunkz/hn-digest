---
source: "https://ridgetext.com/blog/mapbox-llm-composition"
hn_url: "https://news.ycombinator.com/item?id=48789986"
title: "Mapping with In-Memory Layers to Reduce LLM Overload"
article_title: "Mapping with In-Memory Layers to Reduce LLM Overload | RidgeText Blog | RidgeText SMS AI"
author: "Buckwheat469"
captured_at: "2026-07-04T23:53:41Z"
capture_tool: "hn-digest"
hn_id: 48789986
score: 1
comments: 0
posted_at: "2026-07-04T23:25:53Z"
tags:
  - hacker-news
  - translated
---

# Mapping with In-Memory Layers to Reduce LLM Overload

- HN: [48789986](https://news.ycombinator.com/item?id=48789986)
- Source: [ridgetext.com](https://ridgetext.com/blog/mapbox-llm-composition)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T23:25:53Z

## Translation

タイトル: LLM オーバーロードを軽減するためのインメモリ レイヤーを使用したマッピング
記事のタイトル: LLM オーバーロードを軽減するためのインメモリ レイヤーを使用したマッピング |リッジテキストブログ |リッジテキスト SMS AI
説明: コンテキスト ウィンドウを小さく保ち、レンダリングを決定論的に保ちながら、LLM が GeoJSON に一切触れずにレイヤーを調整する、Mapbox 互換のマップ コンポジターを構築した方法。

記事本文:
LLM のオーバーロードを軽減するためのインメモリ レイヤーとのマッピング |リッジテキストブログ | RidgeText SMS AI 機能 ハウツー 価格 ブログ サインイン マップボックス LLM アーキテクチャ マッピング LLM オーバーロードを軽減するためのインメモリ レイヤーを使用したマッピング
コンテキスト ウィンドウを小さく保ち、レンダリングを決定論的に保ちながら、LLM が GeoJSON に触れることなくレイヤーを調整する、Mapbox 互換のマップ コンポジターを構築した方法。
RidgeText にトレイルを重ねた火災境界マップの生成を依頼すると、応答には火災周囲とその上に重ねられたトレイル ルートを含む地図が含まれます。すべてが 1 つの画像でレンダリングされ、SMS 経由で送信されます。
トレイル ルートがオーバーレイされた火災の境界線 (影付きのポリゴン)。これらはそれぞれ、マップがレンダリングされる前に個別にキューに入れられる個別のレイヤーです。
この階層化機能の開発中に、私たちは 1 つのことに気づきました。LLM ツール呼び出しを介して GeoJSON を渡すことはできません。単にデータが多すぎて役に立たないということです。
RidgeText は、LLM 上のオーケストレーション層として構築されます。ユーザーは、アプリや UI を使用せずに SMS を通じて対話し、LLM が自然言語の理解を処理し、呼び出すツールを決定し、返送する明確な応答を作成します。これはルールエンジンではありません。 LLM はあらゆるステップで判断を下します。
この非決定性は会話のための機能ですが、機能に対する制約を生み出します。LLM が触れるものはすべて、変動に対する耐性が必要です。ツールが返すデータが多すぎると、LLM が切り詰められたり、要約が幻覚を示したり、警告なしに失敗したりする可能性があります。ツールのインターフェイスがあいまいな場合、LLM がそれを誤って呼び出す可能性があります。優れたツール設計とは、一連の合理的な応答がすべて正しい結果につながるように、LLM が認識するものを形作ることを意味します。
マップの生成は、これが重要であることを示す明らかな例です。
素朴なアプローチ (そしてそれが失敗する理由)
明らかな実装は次のとおりです

1 つは火災データを取得して LLM に返すツール、次に 2 番目のツールはそのデータを受け入れてマップをレンダリングします。次のようなもの:
LLM は get_wildfire_data() を呼び出します → 2,000 個の火災ポリゴンを GeoJSON として受け取ります
LLM は render_map(geojson: ...) を呼び出します → その GeoJSON を渡します
実際には、適度な山火事データセットは 50 ～ 500 KB の生の GeoJSON です。トークンは約 4 バイトであるため、500 KB は約 125,000 トークンに相当します。これは多くのコンテキスト ウィンドウよりも大きく、収まる場合でも高価です。
LLM は、推論できないデータのパイプになります。 GeoJSON を単純化することも、検証することもできず、毎回コンテキスト コスト全体を支払うことになります。
私たちのソリューションは Mapbox 自体の仕組みを反映しており、レイヤーは個別に追加され、レンダリング時に合成されます。
データを LLM に返す代わりに、各データ取得ツールはその結果をサーバー側に保存し、軽量の確認応答のみを返します。
LLM は、retrieve_wildfire_layer(場所: "カスケード") を呼び出します。
→ { ステータス: "キューに入っています"、レイヤ ID: "wildfires-0"、機能カウント: 847 }
LLM は、retrieve_trail_layer(trailName: "PCT セクション J") を呼び出します。
→ { ステータス: "キュー中"、layerId: "trail-1"、featureCount: 1 }
LLM は、generate_map() を呼び出します
→ {mapUrl: "https://storage.../map-abc123.jpg" }
LLM のコンテキストは、確認応答 (小さな JSON オブジェクト) のみを認識します。 GeoJSON は、generate_map がキューを排出して画像を合成するまで、サーバー メモリ内に存在します。
各retrieve_*呼び出しは、リクエストコンテキストに保持されている順序付けされたレイヤー配列に追加されます。
インターフェイス MapLayer {
タイプ: '野火の境界' | '火災ホットスポット' | 'トレイル' | 'ヒートマップ' | ...;
データ: GeoJSON 。機能コレクション ;
スタイル: LayerStyle ;
}
// インプロセスキュー — 1 LLM ターンの存続期間中存続します
constlayerQueue : MapLayer [] = [];
generate_map は、Mapbox のレイヤー スタックとまったく同じように、挿入順序でレンダリングします。

後のレイヤーは後のレイヤーの下に位置します。 LLM は、ツールを呼び出す順序によって暗黙的に順序を制御します。
私たちの実装では、30 分の TTL でセッション ID をキーとするインプロセス マップを使用するため、キューに入れられたもののレンダリングされなかったレイヤーは、蓄積されるのではなく自動的に削除されます。特定のメカニズムは重要ではありません。Redis、データベース、またはリクエストスコープのコンテキストオブジェクトはすべて機能します。重要なのは、データが LLM のコンテキスト ウィンドウ以外の場所に存在することです。
Mapbox のコア モデルは次のとおりです。ソースはデータを提供し、レイヤーはデータのレンダリング方法を定義し、レイヤーは宣言順に構成されます。サーバー側のキューも同じ抽象化であり、ブラウザなしで実行されるだけです。
レンダラーにおける「Mapbox」が何を意味するのかを正確に把握しておくことが重要です。 Mapbox Static API 画像 (地形、道路、ラベル) をベース タイルとして取得し、キャンバスを使用してその上にデータ レイヤーを合成します。 Mapbox 自体が GeoJSON を認識することはありません。背景を提供するだけです。私たちが保存するレイヤー記述子は、レンダラーが必要とするからではなく、意図的に将来を見据えた選択として Mapbox の形式に従っています。
3D 地形、複雑なブレンディング、アニメーション レイヤーなどで静的タイルが大きくなりすぎた場合は、レンダラーを Playwright ブラウザで実行されているヘッドレス Mapbox GL JS インスタンスに交換できます。そのレンダラーは、ツールや LLM のインターフェイスを変更することなく、まったく同じレイヤー キューを消費します。コストと速度のプロファイルも変化します。JavaScript ベースのレンダラーは、リクエスト全体でタイルと GL アセットをキャッシュできるため、マップごとに新しい静的 API 呼び出しを行う場合と比較して、マップごとのコストが削減され、コールド スタートのレイテンシーが改善される可能性があります。
新しいデータ ソースを追加すると、新しいデータ ソースが追加されます。

retrieve_* ツール — レンダー パイプラインは変更されません
レイヤーの順序付けは、ツール呼び出しシーケンスを通じて表現するのが自然です
スタイリングはレイヤー タイプと同じ場所に配置され、LLM を通過しません。
レンダラーは、LLM レイヤーに触れることなく交換可能です (今日は静的タイル、明日はヘッドレス GL)
generate_map は決定的です。 LLM からは GeoJSON を受け取りません。ズーム レベルやマップ スタイルなどのオプションのパラメーターのみを受け取ります。レイヤー キューを読み取り、各フィーチャ セットを Mapbox Static API ベース イメージに投影し、sharp を使用してそれらを合成します。
async function generatedMap (options : MapOptions ) : Promise < string > {
const レイヤー =ドレインLayerQueue(); // 消費してクリア
const Base = await fetchMapboxBase (オプション);
const generated = await complexLayers (ベース、レイヤー);
return await UploadToStorage (構成済み);
}
その結果、LLM が SMS 応答に添付できる単一の画像 URL が得られます。
「PCT 付近の山火事を見せて」リクエストの場合、LLM のツール呼び出しシーケンスは次のようになります。
[
{ "name" : "retrieve_wildfire_layer" , "result" : { "status" : "queued" , "featureCount" : 847 } },
{ "name" : "retrieve_trail_layer" , "result" : { "status" : "queued" , "featureCount" : 1 } },
{ "名前" : "generate_map" , "結果" : { "mapUrl" : "https://..." } }
】
ツール結果の合計トークン: ~150。このパターンがない場合: ~125,000+。
データセットのサイズに関係なく、コンテキスト ウィンドウは小さいままです
レンダリング パイプラインは決定論的であり、単独でテスト可能です
新しいレイヤー タイプを追加する場合、LLM がシステムと対話する方法を変更する必要はありません。
LLM は、キューに入れられたレイヤーの基礎となるジオメトリについて推論することができません。トレイルが追加されたことと、それに含まれるフィーチャの数は知っていますが、座標自体については何も知りません。 「私のトレイルに最も近い都市はどこですか?」のような質問。キューに入れられた l からは正確に答えることができません

ayer データ — LLM には、それをテキストとして返す別のツールが必要になります。ただし、後からストレージ リンクを介してレンダリングされたマップ イメージを表示し、合成結果を視覚的に観察することはできます。
レイヤーキューは一時的です - ターンをまたいで存続しないため、複数ターンのマップを改良するには再フェッチが必要です
2 番目のトレードオフは、会話応答と同じ履歴保持ポリシーに従うデータベース テーブルにレイヤーを永続化することで解決できます。レンダリングされた各マップ応答は、関連付けられたレイヤーを参照によって保存するため、ユーザーがズームイン、スタイルの変更、または別のデータ ソースの追加を要求した場合、元の API から再フェッチすることなく、レイヤーをデータベースから直接リハイドレートできます。
ここでのパターンは地図に関するものではありません。それは、LLM が推論エンジンではなくデータ パイプとして機能していることを認識し、その役割から LLM を削除することです。
探すシグナル: ツール A がデータをフェッチ → LLM がそれを受信 → LLM がデータをツール B に直接渡す。LLM がコンテンツに基づいて決定を行っていない場合、LLM はコンテンツをまったく保持すべきではありません。
同じアプローチが適用されるいくつかのシナリオ:
マルチソースデータの強化。 NREL は EV 充電ステーションのデータセットを提供していますが、他の API にあるアメニティ データ、リアルタイムの空き状況、ユーザー レビューが欠けています。各取得ツールは、各ソースをフェッチし、結合されたペイロードを LLM に戻して結合するのではなく、そのデータセットをサーバー側のキューに入れます。コンポジターはステーション ID に基づいてこれらを結合し、強化された結果を生成します。 LLM は、どのソースをプルするかを調整し、概要を受け取ります。中間のソースごとのペイロードやマージされたデータセットを保持することはありません。
複数のパスを使用したログ分析。ユーザーが「過去 1 時間で何件のエラーがありましたか?」と尋ねます。

これは、同じ基盤となるデータに対する 2 つの分析上の質問です。このパターンがなければ、ログを 2 回フェッチするか (ツール呼び出しごとに 1 回)、または呼び出しと呼び出しの間で生のログ ペイロードを LLM に戻すことになります。代わりに、ログ フェッチは 1 回実行され、結果が保存されます。エラーカウント ツールと根本原因ツールは両方とも同じ保存されたデータセットから独立して読み取ります。LLM は、業務上保持していない生のログ データではなく、2 つの小さな構造化された回答を取得します。
出力が重要な ETL パイプライン。複数の上流ソースから取得し、それらを 1 つの結果 (レポート、ダッシュボード データセット、ランク付けリスト) に合成する場合、LLM の役割は、何を取得して結果を説明するかを決定することであり、マージに参加することではありません。ソース間の中間状態は、コンテキスト ウィンドウではなくパイプラインに属します。
RidgeText は、まさにこのパターンを使用して、SMS 経由で火災の境界、トレイルの状況、気象警報を配信します。アプリやデータ プランは必要ありません。
© 2026 リッジテキスト。すべての著作権は留保されています。

## Original Extract

How we built a Mapbox-compatible map compositor where the LLM orchestrates layers without ever touching GeoJSON — keeping context windows small and rendering deterministic.

Mapping with In-Memory Layers to Reduce LLM Overload | RidgeText Blog | RidgeText SMS AI Features How-To Pricing Blog Sign In mapbox llm architecture mapping Mapping with In-Memory Layers to Reduce LLM Overload
How we built a Mapbox-compatible map compositor where the LLM orchestrates layers without ever touching GeoJSON — keeping context windows small and rendering deterministic.
If you ask RidgeText to generate a fire perimeter map with trails overlaid, the response now includes a map with fire perimeters and your trail route layered on top — all rendered in a single image and sent via SMS.
Fire perimeter (shaded polygon) with a trail route overlaid. Each of these is a separate layer queued independently before the map is rendered.
While developing this layering feature, we realized one thing: you cannot pass GeoJSON through an LLM tool call , it's simply too much data to be useful.
RidgeText is built as an orchestration layer on top of an LLM. Users interact through SMS — no app, no UI — and the LLM handles natural language understanding, decides which tools to call, and composes a clear response to send back. It's not a rules engine; the LLM is making judgment calls at every step.
That non-determinism is a feature for conversation, but it creates a constraint for features: anything the LLM touches needs to be resilient to variation. If a tool returns too much data, the LLM may truncate, hallucinate a summary, or fail silently. If a tool's interface is ambiguous, the LLM may call it incorrectly. Good tool design means shaping what the LLM sees so that the range of reasonable responses all lead to correct outcomes.
Map generation is a clear example of where this matters.
The Naive Approach (and Why It Fails)
The obvious implementation is to have a tool that fetches fire data and returns it to the LLM, then a second tool that accepts that data and renders a map. Something like:
LLM calls get_wildfire_data() → receives 2,000 fire polygons as GeoJSON
LLM calls render_map(geojson: ...) → passes that GeoJSON along
In practice, a modest wildfire dataset is 50–500KB of raw GeoJSON. A token is roughly 4 bytes, so 500KB is ~125,000 tokens — larger than many context windows, and expensive even when it fits.
The LLM becomes a pipe for data it cannot reason about. It can't simplify the GeoJSON, it can't validate it, and it pays the full context cost every time.
Our solution mirrors how Mapbox itself works: layers are added independently and composited at render time.
Instead of returning data to the LLM, each data-fetching tool stores its result server-side and returns only a lightweight acknowledgment:
LLM calls retrieve_wildfire_layer(location: "Cascades")
→ { status: "queued", layerId: "wildfires-0", featureCount: 847 }
LLM calls retrieve_trail_layer(trailName: "PCT Section J")
→ { status: "queued", layerId: "trail-1", featureCount: 1 }
LLM calls generate_map()
→ { mapUrl: "https://storage.../map-abc123.jpg" }
The LLM's context only ever sees the acknowledgments — tiny JSON objects. The GeoJSON lives in server memory until generate_map drains the queue and composites the image.
Each retrieve_* call appends to an ordered layer array held in the request context:
interface MapLayer {
type : 'wildfire-perimeters' | 'fire-hotspots' | 'trail' | 'heatmap' | ...;
data : GeoJSON . FeatureCollection ;
style : LayerStyle ;
}
// In-process queue — lives for the lifetime of one LLM turn
const layerQueue : MapLayer [] = [];
generate_map renders them in insertion order — exactly like Mapbox's layer stack, where earlier layers sit below later ones. The LLM controls ordering implicitly by the sequence it calls the tools: if it calls retrieve_satellite_base before retrieve_trail , the trail draws on top of the satellite imagery.
Our implementation uses an in-process Map keyed by session ID with a 30-minute TTL, so layers that were queued but never rendered are evicted automatically rather than accumulating. The specific mechanism isn't the point — Redis, a database, or a request-scoped context object would all work. What matters is that the data lives somewhere other than the LLM's context window.
Mapbox's core model is: sources provide data, layers define how to render it, and layers compose in declaration order. Our server-side queue is the same abstraction, just running without a browser.
It's worth being precise about what "Mapbox" means in our renderer. We fetch a Mapbox Static API image as the base tile — terrain, roads, labels — and then composite the data layers on top of it using canvas. Mapbox itself never sees the GeoJSON; it only provides the background. The layer descriptors we store follow Mapbox's format not because the renderer requires it, but as a deliberate forward-looking choice.
If we ever outgrow static tiles — for 3D terrain, complex blending, or animated layers — we can swap the renderer for a headless Mapbox GL JS instance running in a Playwright browser. That renderer would consume the exact same layer queue without any changes to the tools or the LLM's interface. The cost and speed profile would also change: a JavaScript-based renderer can cache tiles and GL assets across requests, potentially reducing per-map costs and improving cold-start latency compared to making a fresh Static API call for every map.
Adding a new data source is adding a new retrieve_* tool — the render pipeline doesn't change
Layer ordering is natural to express through tool call sequence
Styling is co-located with the layer type, not passed through the LLM
The renderer is swappable — static tiles today, headless GL tomorrow — without touching the LLM layer
generate_map is deterministic. It receives no GeoJSON from the LLM — only optional parameters like zoom level or map style. It reads the layer queue, projects each feature set onto a Mapbox Static API base image, and composites them using sharp.
async function generateMap ( options : MapOptions ) : Promise < string > {
const layers = drainLayerQueue (); // consume and clear
const base = await fetchMapboxBase (options);
const composed = await compositeLayers (base, layers);
return await uploadToStorage (composed);
}
The result is a single image URL the LLM can attach to its SMS response.
For a "show me wildfires near the PCT" request, the LLM's tool call sequence looks like:
[
{ "name" : "retrieve_wildfire_layer" , "result" : { "status" : "queued" , "featureCount" : 847 } },
{ "name" : "retrieve_trail_layer" , "result" : { "status" : "queued" , "featureCount" : 1 } },
{ "name" : "generate_map" , "result" : { "mapUrl" : "https://..." } }
]
Total tokens in tool results: ~150. Without this pattern: ~125,000+.
Context window stays small regardless of dataset size
Render pipeline is deterministic and testable in isolation
Adding new layer types doesn't require changing how the LLM interacts with the system
The LLM can't reason about the underlying geometry of any layer it queued. It knows a trail was added and how many features it contains, but nothing about the coordinates themselves. A question like "what's the nearest city to my trail?" can't be answered precisely from the queued layer data — the LLM would need a separate tool that returns that as text. It can, however, view the rendered map image later via the Storage link and make visual observations about the composite result.
Layer queue is ephemeral — it doesn't survive across turns, so multi-turn map refinement requires re-fetching
The second tradeoff could be addressed by persisting layers to a database table that follows the same history retention policies as conversation responses. Each rendered map response would store its associated layers by reference, so if the user follows up — asking to zoom in, change the style, or add another data source — the layers can be rehydrated directly from the database without re-fetching from the original APIs.
The pattern here isn't about maps. It's about recognizing when your LLM is acting as a data pipe instead of a reasoning engine — and removing it from that role.
The signal to look for: Tool A fetches data → LLM receives it → LLM passes it directly to Tool B. If the LLM isn't making a decision based on the content, it shouldn't be holding the content at all.
A few scenarios where the same approach applies:
Multi-source data enrichment. NREL provides a dataset of EV charging stations, but it's missing amenity data, real-time availability, and user reviews that other APIs carry. Rather than fetching each source and passing the merged payload back to the LLM for combination, each retrieval tool queues its dataset server-side. A compositor joins them on station ID and produces the enriched result. The LLM orchestrates which sources to pull and receives a summary — it never holds the intermediate per-source payloads or the merged dataset.
Log analysis with multiple passes. A user asks "how many errors did we have in the last hour, and what are the primary issues?" That's two analytical questions over the same underlying data. Without this pattern, you'd fetch the logs twice — once per tool call — or pass the raw log payload back to the LLM between calls. Instead, the log fetch runs once and stores the result. An error-count tool and a root-cause tool both read from the same stored dataset independently. The LLM gets two small structured answers rather than raw log data that it has no business holding.
Any ETL pipeline where the output is what matters. If you're pulling from multiple upstream sources and compositing them into a single result — a report, a dashboard dataset, a ranked list — the LLM's role is to decide what to pull and describe the result, not to participate in the merge. The intermediate state between sources belongs in your pipeline, not in the context window.
RidgeText uses this exact pattern to deliver fire perimeters, trail conditions, and weather alerts over SMS — no app or data plan required.
© 2026 RidgeText. All rights reserved.
