---
source: "https://pixelgust.com/blog/weather-mcp-server"
hn_url: "https://news.ycombinator.com/item?id=49009991"
title: "Show HN: An MCP server that gives AI assistants live weather and hazard data"
article_title: "Weather MCP Server: Connect Claude to Live Weather, Climate & Hazard Data | PixelGust"
author: "pixelgust"
captured_at: "2026-07-22T18:04:49Z"
capture_tool: "hn-digest"
hn_id: 49009991
score: 1
comments: 0
posted_at: "2026-07-22T17:09:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: An MCP server that gives AI assistants live weather and hazard data

- HN: [49009991](https://news.ycombinator.com/item?id=49009991)
- Source: [pixelgust.com](https://pixelgust.com/blog/weather-mcp-server)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T17:09:14Z

## Translation

タイトル: Show HN: AI アシスタントにライブ天気と危険データを提供する MCP サーバー
記事のタイトル: Weather MCP サーバー: Claude をライブ天気、気候、危険データに接続 |ピクセルガスト
説明: PixelGust は、クロード、カーソル、その他の AI アシスタントに地球上のあらゆる場所のライブ天気、予報、10 年間の気候履歴、地形、火災リスク、洪水データを提供する無料の MCP サーバーを実行しています。

記事本文:
ブログ起動アプリ
ホーム › ブログ › 気象 MCP サーバー
Weather MCP サーバー: Claude をライブ天気、気候、危険データに接続
AI アシスタントは推論には優れていますが、現在に対して盲目です。クロードに物件付近の現在の火災の危険性について尋ねると、数か月前に終了したトレーニング データからしか答えられません。モデル コンテキスト プロトコル (MCP) はこの問題を解決します。MCP は、AI アシスタントがライブ ツールを呼び出すことができるオープン スタンダードです。 PixelGust はパブリック MCP サーバーを実行しているため、互換性のあるアシスタントは会話中に、地球上の任意の地点の実際の天気、気候、地形、危険データを取得できます。
サーバー URL: https://pixelgust.com/mcp (ストリーミング可能な HTTP)。無料利用枠の制限内でサインアップなしで動作します。これを Claude、Cursor、または任意の MCP クライアントに追加して、位置に関する質問を開始します。
接続すると、アシスタントは推測ではなくライブ データを使用して次のような質問に答えます。
「アテネ対マルセイユの火災の危険性は現在どのくらいですか?」
「このブドウ畑の場所は、ここ10年で乾燥してきましたか？」
「これら 2 つの座標の地形、洪水の起こりやすさ、土壌を比較してください。」
「9月上旬の気温46.5度、11.3度ではどんな天気になると予想すべきでしょうか?」
「この多角形内の平均の傾きと人口はいくらですか?」
今日の実際のやり取りでは、アシスタントがサーバーを 2 回呼び出して比較しています。
同助手は、熱と乾燥がマルセイユの強い風を上回るため、両都市は危険度「高」に属し、アテネは「非常に危険」に近いと結論付けた。これが、MCP が可能にする、地に足の着いた現在の答えです。
get_current_weather : ある地点の温度、湿度、風、降水量、気圧、雲量、日射量 (NOAA GFS、6 時間ごとに更新)
get_weather_forecast : 7 日先までの毎日の天気予報
get_historyal_climate : 10 年間の気候平均 (2015 年から 2025 年、

ERA5)、年間または月ごと
get_climate_timeseries : 傾向分析用の 1 つの変数の 2015 年から 2025 年の月次値
get_terrain : 30m での標高、傾斜角、傾斜角 (Copernicus DEM)、および湿潤指数
get_hazards : 火災気象指数、土壌侵食リスク (RUSLE)、洪水感受性 (TWI)
get_environment : NDVI 植生指数、蒸発散量、土地被覆クラス
get_proximity : 最も近い都市、空港、港、病院、発電所までの距離
get_polygon_stats : カスタム ポリゴン領域にわたる統計の集計
[設定] → [コネクタ] → [カスタム コネクタを追加] → https://pixelgust.com/mcp を貼り付けます。
クライアントの MCP 構成に同じ URL を持つストリーミング可能 HTTP サーバーを追加します。
キーレス アクセスでは、無料利用枠の制限 (1 日あたり 50 件の通話) が使用されます。詳細については、Pixelgust.com/app で無料アカウントを作成し、[API キー] メニューから API キーを生成し、それをヘッダーとして渡します。
キーのプランは自動的に適用されます。有料レベルでは 1 日の割り当てが増加し、より大きなポリゴン エリアのロックが解除されます。詳細については、API ドキュメントを参照してください。
アシスタントに実際の環境データを提供する
URLは1つ。九つの道具。地球上のあらゆる場所。無料で始められます。
エージェントは、専門家がデータを操作する手段になりつつあります。農学者がアシスタントに分野間の干ばつストレスの比較を依頼したり、アナリストが気候リスクの特性をスクリーニングしたり、10 年にわたる気候シリーズを抽出した研究者は、最初に GIS スイートを開いたり、API 統合コードを作成したりする必要はありません。 MCP では、アシスタントが呼び出しを行い、人間が質問を行います。
内部では、サーバーは PixelGust REST API と同じデータ (NOAA GFS ナウキャストと予報、ERA5 再解析、コペルニクス DEM 地形、MODIS 蒸発散量、ESA WorldCover、SoilGrids、および毎日計算されるカナダ火災気象指数) を公開します。同じソース、新しいインターフェイス。
© 2026 ピクス

lガスト · ホーム · ブログ · ダッシュボード
あらゆる土地区画の天気、地形、危険プロファイルを即座に取得します。

## Original Extract

PixelGust runs a free MCP server that gives Claude, Cursor, and other AI assistants live weather, forecasts, 10-year climate history, terrain, fire risk, and flood data for any location on Earth.

Blog Launch App
Home › Blog › Weather MCP Server
Weather MCP Server: Connect Claude to Live Weather, Climate & Hazard Data
AI assistants are good at reasoning but blind to the present. Ask Claude about the current fire risk near a property and it can only answer from training data that ended months ago. The Model Context Protocol (MCP) fixes this: it is an open standard that lets AI assistants call live tools. PixelGust now runs a public MCP server, so any compatible assistant can pull real weather, climate, terrain, and hazard data for any point on Earth, mid-conversation.
Server URL: https://pixelgust.com/mcp (Streamable HTTP). Works with no signup at free-tier limits. Add it to Claude, Cursor, or any MCP client and start asking location questions.
Once connected, your assistant answers questions like these with live data instead of guesses:
"What is the fire risk in Athens versus Marseille right now?"
"Has this vineyard's location gotten drier over the last decade?"
"Compare the terrain, flood susceptibility, and soil of these two coordinates."
"What weather should I expect at 46.5, 11.3 in early September?"
"What is the average slope and population inside this polygon?"
A real exchange from today, with the assistant calling the server twice and comparing:
The assistant concluded that both cities sit in the High danger class, with Athens closer to Very High because heat and dryness outweigh Marseille's stronger wind. That is the kind of grounded, current answer MCP makes possible.
get_current_weather : temperature, humidity, wind, precipitation, pressure, cloud cover, and solar radiation for a point (NOAA GFS, updated every 6 hours)
get_weather_forecast : daily forecast up to 7 days ahead
get_historical_climate : 10-year climate averages (2015-2025, ERA5), annual or by month
get_climate_timeseries : monthly values 2015-2025 for one variable, for trend analysis
get_terrain : elevation, slope, and aspect at 30m (Copernicus DEM), plus wetness index
get_hazards : Fire Weather Index, soil erosion risk (RUSLE), flood susceptibility (TWI)
get_environment : NDVI vegetation index, evapotranspiration, land cover class
get_proximity : distance to the nearest city, airport, port, hospital, and power plant
get_polygon_stats : aggregate statistics over a custom polygon area
Settings → Connectors → Add custom connector → paste https://pixelgust.com/mcp .
Add a Streamable HTTP server with the same URL in your client's MCP configuration.
Keyless access uses free-tier limits (50 calls per day). For more, create a free account at pixelgust.com/app , generate an API key from the API Keys menu, and pass it as a header:
The key's plan applies automatically: paid tiers raise the daily quota and unlock larger polygon areas. See the API documentation for details.
Give Your Assistant Real Environmental Data
One URL. Nine tools. Any location on Earth. Free to start.
Agents are becoming the way professionals interact with data. An agronomist asking an assistant to compare drought stress across fields, an analyst screening properties for climate risk , or a researcher pulling decade-long climate series should not have to open a GIS suite or write API integration code first. With MCP, the assistant does the calling, and the human does the asking.
Under the hood the server exposes the same data as the PixelGust REST API : NOAA GFS nowcasts and forecasts, ERA5 reanalysis, Copernicus DEM terrain, MODIS evapotranspiration, ESA WorldCover, SoilGrids, and the Canadian Fire Weather Index computed daily. Same sources, new interface.
© 2026 PixelGust · Home · Blog · Dashboard
Get instant weather, terrain, and hazard profiles for any land parcel.
