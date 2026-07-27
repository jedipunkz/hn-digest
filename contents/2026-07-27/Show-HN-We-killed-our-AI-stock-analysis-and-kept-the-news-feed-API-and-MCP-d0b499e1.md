---
source: "https://alphai.io/developers"
hn_url: "https://news.ycombinator.com/item?id=49068635"
title: "Show HN: We killed our AI stock analysis and kept the news feed (API and MCP)"
article_title: "Stock News API — Financial News API for Bots & Agents | alphai"
author: "mmakeev"
captured_at: "2026-07-27T12:51:18Z"
capture_tool: "hn-digest"
hn_id: 49068635
score: 1
comments: 0
posted_at: "2026-07-27T12:23:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: We killed our AI stock analysis and kept the news feed (API and MCP)

- HN: [49068635](https://news.ycombinator.com/item?id=49068635)
- Source: [alphai.io](https://alphai.io/developers)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T12:23:12Z

## Translation

タイトル: Show HN: AI 株分析を停止し、ニュース フィード (API と MCP) を保持しました
記事のタイトル: Stock News API — ボットとエージェント向けの金融ニュース API |アルファイ
説明: 関連性スコア付きティッカーリンク金融ニュース用の REST API: 取引ボット、AI エージェント、ダッシュボード用のリアルタイム株式ニュース API。無料枠、クレジットカードなし。
HN テキスト: 私は AlphaAI を実行しています。これは、REST と MCP を組み合わせた取引エージェント向けのニュース フィードです。以前は「AI銘柄分析」サイトでした。
そのバージョンのデモは素晴らしかったが、収益は得られなかった。私は市場データにお金を払っていましたが、ユーザーは私にお金を払っていませんでした。プロンプトの出力を購入する人は誰もいません。誰もがすでに独自の ChatGPT または Claude と独自のプロンプトを持っています。人々が見ていたのは評価付きのニュース フィードだけだったので、この機能を中心に新しい製品を構築したところ、非常に便利であることがわかりました。内部の内容: GDELT 解析、SEC EDGAR、Google パブリック RSS、およびその他のいくつかのソース。すべては重複排除と強化のパイプラインを通過します。本当に役立つ情報のほんの一部だけがフィードに含まれます。ユーザーにとって価値があるのはフィルタリングです。エージェントにニュースを送信する人への質問です。通常、不足しているメタデータは何ですか?私にとって、これがこの投稿から学びたい主なことです。

記事本文:
株式ニュース API — ボットおよびエージェント用の金融ニュース API | alphai AlphAI alphai — AI エージェント向けの金融ニュース フィード インサイダー リサーチ API MCP SDK TUI 例 価格設定 検索ティッカー… ⌘K 読み込み中… パブリック API 取引ボット用の株式ニュース API。
リクエストにより。
alphai を強化するのと同じフィードが、プレーンな JSON として提供されました。すべての記事は関連性がスコアリングされ、ティッカーにリンクされて到着するため、取引ボットまたはエージェントのバックエンドはリクエスト時にワイヤーをフィルタリングできます。 Auth は 1 つの Bearer ヘッダーです。
カール https://api.alphai.io/api/news/ \
-H "認可: ベアラー $ALPHAI_API_KEY" \
-G -d "symbol=NVDA" -d "min_relevance=7" ライブ フィードで現在測定されているフィード、15 分ごとに更新 過去 24 時間の記事 1,085 件 検証済みティッカーで関連性フロアをクリアした行。これは、/api/news/ ページへのリクエストを設定します。
スコアが付けられ、ティッカータグが付けられ、すでに API 上に公開されています。
それらの記事全体での言及数。 ?symbol= を使用してそれらのいずれかをクエリします。
新しい記事は公開から 1 時間以内に到着します。 SEC への提出は数分でフィード行になります (中央値 ~6)。自分のスケジュールでフィードをポーリングします。
記事には機械で抽出されたティッカーが掲載されており、すべてのティッカーが独自の分析を受けています。
すべての記事には、LLM によって割り当てられた関連性スコア (1 ～ 10) が付けられます。リクエスト時にフィルタリングします。
HTTPS 経由のプレーン JSON。任意の HTTP クライアントを使用するか、公式の型指定された Python / TypeScript SDK をインストールしてください。どちらにしてもロックインはありません。
ほとんどの金融ニュース API は、あなたに消防ホースを渡し、あなたの幸運を祈っています。ストーリーにリクエストが発生する前にフィルタリングが行われるため、API に到達するのはモデルのコンテキスト ウィンドウに相当する部分になります。
GDELT 主導のモニターが世界の金融報道を 24 時間体制で監視し、2 人目の職員が Form 4 および 8-K の提出のために SEC EDGAR を尾行します。パブリッシャーが厳選した RSS フィードは、次のようなギャップを埋めます。

えー、キャッチします。すべてが 1 つのフィードに表示されます。
LLM パスでは、各記事の関連性について 1 ～ 10 のスコアが付けられ、トピック カテゴリが割り当てられ、ティッカーごとのセンチメントが信頼レベルと推論とともに書き込まれます。
抽出されたティッカーは、アクティブシンボルリストおよび記事テキスト自体と照合されます。幻覚を見せたシンボルは削除され、PR ワイヤーのスパム ドメインは完全にブロックされ、ほぼ重複したものは 1 つの行にまとめられます。
関連性の下限をクリアし、少なくとも 1 つの検証済みティッカーを含む記事のみが API に到達します。 6,900 以上の監視対象ソースから合格できるのは 5 人に 1 人未満であるため、ノイズがパーサーに到達することはありません。
フィードからの実際の行です。
/api/news/ が返すとおり、スコア 8 以上の最新記事。このページで説明されていないフィールドは長さの関係で省略されています。形は本物の契約書です。
{
「オリジナル」: {
"uid": "2f9c5868891f5fcb",
"title": "デッカーズ ブランド、HOKA と UGG が牽引し、歴史的な第 1 四半期売上高 10 億ドルを達成",
"time_published": "2026-07-27T12:15:00Z",
"source_domain": "yahoo.com",
"url": "https://finance.yahoo.com/markets/stocks/articles/deckers-brands-hits-history-1bn-112734988.html"
}、
「強化」: {
「関連性スコア」: 8、
"カテゴリ": "収益",
"ai_trading_insights": {
"ティッカー分析": [
{
"ティッカー": "デッキ",
"影響分析": {
「感情」: 「ポジティブ」、
"自信": "中",
"reasoning": "この記事は、ブランドのパフォーマンスに関連した複数の将来予想データポイント (粗利予測の増加、21.5% をわずかに上回る営業利益率、EPS ガイダンスの引き上げ) を提供しています。通常、これが純粋な再評価と比較して再評価を促進します…"
}
}、
「…」
]
}
}
公式 SDK タイプ付きクライアント、2 つの言語。
ハンドロールされたリクエストよりも型指定されたクライアントを好みますか?公式の Python および TypeScript SDK は 9 つのエンドポイントすべてをラップし、カーソルのページネーション、429/5xx での再試行、型指定エラー、およびレート制限検査を処理します。

あなたのためのアクション。これらは同じキーを受け取り、同じデータを返すため、後で採用することも、まったく採用しないこともできます。
どちらも環境から ALPHAI_API_KEY を読み取り、デフォルトで api.alphai.io をターゲットにします。 Python クライアントは同期と非同期を提供します。 TypeScript クライアントにはランタイム依存関係はありません。
ターミナルダッシュボード: alpha-tui →
$ pip install alphai-sdk から alphai import クライアント
Client() を c として使用:
ページ = c.news.list(symbol="NVDA")
page.results の場合:
print(a.relevance_score, a.title) PyPI → ソース → TypeScript ノード · エッジ · デノ · ブン $ npm install alphai-sdk import { AlphaAI } from "alphai-sdk";
const c = 新しい AlphaAI();
const page = await c.news.list({
シンボル: "NVDA"、
});
for (const a of page.results)
console.log(a.enrichment.relevance_score); npm → ソース → SDK で構築: alphai-news-to-email — Python でデプロイ可能な news-to-email ダイジェスト · alphai-sdk-ts-examples — API サーフェス全体をカバーする実行可能な TypeScript スクリプト。すべての例 →
認可: Bearer ak_live_… で認証します。クォータには 1 分あたりのバースト キャップと 1 日あたりのボリューム キャップの 2 つの層があり、どちらもキーごとではなくアカウントに対してカウントされます。無料は評価用および非営利使用です。制作には有料プランが必要です。同じキーはヘッドレス エージェントの MCP サーバーも認証します (そこにもベアラー ヘッダーとして送信します)。インタラクティブ AI クライアントは、代わりに OAuth 経由で MCP に接続します。
記事は公開時間よりも遅れてフィードに到達します。一般ニュースは公開後中央値で約 33 分、SEC への提出は 5 ～ 9 分で配信されます (2026 年 7 月 25 日に測定)。したがって、time_published を追跡するポーラーは遅れた到着を見逃します。 sort=ingested は、代わりにフィードに入力された順序で行を返します。カーソルを使用せずに 1 回呼び出し、 next_cursor を保存してから、それを使用して呼び出しを繰り返します。このモードでは、rごとに

esponse にはカーソルが表示され、空の結果リストは追い付いていることを意味します。すべてのフィルターは変更されずに機能します。応答は 60 秒間キャッシュされるため、1 分に 1 回以上の頻度でポーリングしても何も得られません。すべての呼び出しで sort=ingested を渡し続けます。各モードは独自のカーソルを発行し、一方のモードをもう一方のモードで再生すると、フィードの先頭から静かに再開するのではなく、400 が返されます。同じモードは /api/news/insider/ でも機能しますが、ここではギャップが最も大きくなります。Form 4 はレポートされる取引の数日後に提出されるため、新しいイベントがすでにフィードの最新ページの下に入力されることがよくあります。
# pip alpha-sdk をインストールします
インポート時間
alpha インポートクライアントから
カーソル = なし # 実行間でこれを保持します
Client() をクライアントとして使用: # $ALPHAI_API_KEY を読み取ります
True の場合:
ページ = client.news.list(sort="取り込み", カーソル=カーソル)
page.result の記事の場合:
print(f"[{article.relevance_score}/10] {article.title}")
カーソル = page.next_cursor # このモードでは常に存在します
time.sleep(900) # 96 ポーリング/日は、無料の 1 日の上限に適合します。これは、カーソルの記録を行う公式の Python SDK です。同じ並べ替えオプションを備えた TypeScript もありますが、依存関係を持たない場合は、エンドポイントはプレーンな GET のままです。
すべてのニュース項目には、記事に対する LLM パスによって生成されたエンリッチメント ブロックが含まれています。これらは最もよく使用する 4 つのフィールドです。
スコアは、言及されている企業ではなく、記事を評価します。超大型株の 1 週間前の収益の要約はスコアが低く、小型株の新たな FDA 承認のスコアは高くなります。スコアリングは決定的です。同じ記事は常に同じスコアを獲得し、SEC Form 4 申請はモデルではなく取引自体 (規模、買い対売り、10b5-1 計画の有無) からスコアリングされます。
正式な契約が必要ですか?インタラクティブ API リファレンスは /api/schema/swagger-ui/ にあります (Swagger UI、生の OpenAPI 仕様リンク付き)

上部の d)。
記事自体がどれほどの取引価値を持っているか。
1–2 · 取引との関連性なし — ライフスタイル、勧誘、短期金利のスクレイピング。
3–4 · 既知の出来事に関する派生コンテンツ — 意見記事、総括、13F ステークの記事、マーケットラップ。
5–6 · リードアクロスのあるマクロ/セクター データポイント。マイナーだけど本当の企業ニュース。
7–8 · 実際の企業ニュース — 論文、決算、臨床データ、新鮮な触媒による同日の動きによるアナリストの行動。
9–10 · 一次資料、資料、新規開示 — M&A、番号付き印刷物、承認、主要契約。
/api/news/ フィードのデフォルトは ≥ 4 です。 ?min_relevance=N を渡してフィルターを緩めるか締めます。
ティッカーごとの価格の影響の方向。利益がプラスかマイナスかではなく、これは「このストーリーでチャートがどこに動くか」ということです。
各ティッカー ブロックには、impact_analysis.confidence (高 / 中 / 低) とフリーテキストの推論フィールドも含まれており、自動化パイプラインをフィルタリングして自信のある強気シグナルまたは弱気シグナルに絞り込むときに役立ちます。
記事ごとに 1 つのトピック バケット。 /api/news/ で ?category=earnings (CSV または OR マッチング用に繰り返し) を使用してフィードをフィルタリングするか、?exclude_categories= を使用して取引しないバケットを削除します。
数値は、ライブ フィードの過去 7 日間の記事数であり、15 分ごとに更新されます。
記事から抜粋したティッカー。 /api/symbols/ のアクティブなシンボルのリスト (米国株式、仮想通貨、外国上場銘柄) に存在するシンボルのみを保持するため、アクティブなティッカーに言及していない記事は API に到達する前に除外されます。米国株はベアシンボルを使用し、暗号通貨は BTC-USD のような -USD サフィックスを使用し、外国上場では VOD.L のような Yahoo サフィックスを使用します。各ティッカーには感情と推論を含む専用のticker_analysisエントリも取得されるため、マルチティッカーストア

ie は 1 回の読み取りに折りたたまれません。
N 個のアウトレットによって報告される 1 つの実際のイベントは、デフォルトでは N 行になります。代わりに、 /api/news/ に ?collapse=story を追加して、ストーリーごとに 1 行を取得します。代表的な記事には、story_id (独自の UID、 /api/news/ {uid} / で解決可能)、sources_count (記事を実行した個別のアウトレットの数。ほとんどの記事は 1 つのアウトレットによって伝えられるため、通常は 1)、およびソース (個別のドメイン) が含まれます。これらのフィールドは、デフォルトのフィードでは null です。シンボルやカテゴリなどのフィルタは代表行に適用されるため、代表が一致しないストーリーは省略されます。
専任の職員が SEC EDGAR を追跡し、すべての Form 4 提出 (役員と取締役が売買) を提出直後にフィード列に変換します。 1 行は 1 つの経済イベントをカバーします。つまり、同じ種類および保有形態の申告対象のすべての非デリバティブ取引が 1 つの記事に集約されます。 JSON 形状は他のすべての記事と一致するため、パーサーは特別なケースを必要としません。
フォーム 4 の概要は、申告自体から構築された決定論的なテンプレートです。インサイダー名、役割、取引コード、総株式数、価格 (取引が複数のトランシェにまたがる場合の出来高加重平均と約定範囲)、および合計金額です。あなたと開示の間には、LLM の言い換えはありません。関連性スコアはイベント自体 (イベントの合計サイズ、買いと売り、10b5-1 計画の有無など) から計算されるため、再現可能です。
GET /api/news/insider/ — カーソルでページ分割され、スコープに ?symbol= が設定された Form 4 インサイダー フィード。
GET /api/news/?category=insider — メイン フィードのカテゴリ フィルターを介した同じ行。
GET /api/symbols/ {ticker} /insider-summary/ — 30 日間のロールアップ: 買いと売り、ドルの出来高、10b5-1 プランに基づく割合、最もアクティブなインサイダー。
GET /api/symbols/ {ticker} /sentiment-summary/ — その兄弟: 7 日間の強気 / 中立 / 弱気カウント

ニュースの流れより。
インサイダー取引 API の詳細 →
{
"ティッカー": "NVDA",
「日」: 30、
"total_transactions": 14、
「購入数」: 2、
「販売数」: 12、
"buy_value_usd": "1240000.00",
"sell_value_usd": "224580213.05",
"pct_10b5_1": 64、
"top_insider": [
{
"名前": "スティーブンス マーク A",
"タイトル": "ディレクター",
「トランザクション数」: 3、
"net_value": "-221102600.00"
}、
{
"名前": "ダビリ ジョン",
"タイトル": "ディレクター",
「トランザクション数」: 1、
"net_value": "-133750.00"
}
]
値フィールドは USD の 10 進数文字列です。ウィンドウに買いまたは売りがない場合、buy_value_usd / sell_value_usd は null になります。 net_value はインサイダーごとの買いから売りを差し引いた値です。申告者が役員の役割を持たない取締役として報告する場合、タイトルは空にすることができます。
API 接続されたボットでできること。
ウォッチリストのアラートから Telegram ボットまで、どのスタックを実行しても同じ JSON です。 API は邪魔になりません。
「クロード、今日 NVDA を動かしたものは何ですか?」
02 REST API 「ウォッチリストに関連性​​の高い収益ニュースのみを送信します。」
/api/news/?category=収益 · min_relevance=7
03 REST API 「市場を動かすニュースの Telegram ボットを 10 分で構築」
/api/news/trending/ + ボット フレームワーク
04 MCP 「今週ウォッチリストにインサイダー買いはありますか?」
alphai_insider_news → MCP · /api/news/insider/
05 MCP 「MCP 経由で、エージェントに関連性スコアを付けた財務ニュースをフィードします。」
OAuth 2.1

[切り捨てられた]

## Original Extract

REST API for relevance-scored, ticker-linked financial news: a real-time stock news API for trading bots, AI agents, and dashboards. Free tier, no credit card.

I run AlphaAI, it is a news feed for trading agents, REST plus MCP. Before it was "AI stock analysis" site.
That version demoed great and made no money. i was paying for market data, users was not paying me. Nobody buys prompt output, everyone already has own ChatGPT or Claude and own prompts. The only thing people were watching was the news feed with ratings, so I built a new product around this feature and it turned out to be really useful. What's under the hood: GDELT parsing, SEC EDGAR, google public RSS and several other sources. Everything goes through a deduplication and enrichment pipeline. Only a small part of the really useful stuff gets into the feed, it's the filtering that's valuable to the user. Question to people who wire news into agents: what metadata you missing usually? For me this is main thing i want to learn from this post.

Stock News API — Financial News API for Bots & Agents | alphai AlphAI alphai — Financial news for AI agents Feed Insider Research API MCP SDK TUI Examples Pricing Search ticker… ⌘K Loading... Public API A stock news API for trading bots.
By the request.
The same feed that powers alphai, served as plain JSON. Every article arrives scored for relevance and linked to tickers, so a trading bot or an agent backend can filter the wire at request time. Auth is one Bearer header.
curl https://api.alphai.io/api/news/ \
-H "Authorization: Bearer $ALPHAI_API_KEY" \
-G -d "symbol=NVDA" -d "min_relevance=7" The feed right now measured on the live feed, updated every 15 minutes 1,085 Articles in the last 24 hours Rows that cleared the relevance floor with a verified ticker. This is the set a request to /api/news/ pages through.
Scored, ticker-tagged, and already live on the API.
Mention counts across those articles. Query any of them with ?symbol= .
New articles land within the hour of publication; SEC filings become feed rows in minutes (median ~6). Poll the feed on your own schedule.
Articles carry machine-extracted tickers, and every ticker gets its own analysis.
Every article carries an LLM-assigned relevance score (1–10). Filter at request time.
Plain JSON over HTTPS. Bring any HTTP client, or install the official typed Python / TypeScript SDK if you prefer one. No lock-in either way.
Most financial-news APIs hand you a firehose and wish you luck. We filter before a story ever costs you a request, so what reaches the API is the part worth a model's context window.
A GDELT-driven monitor watches the global financial press around the clock, while a second worker tails SEC EDGAR for Form 4 and 8-K filings. Hand-curated publisher RSS feeds fill the gaps neither catches. Everything lands in one feed.
An LLM pass scores each article 1–10 for relevance, assigns it a topic category, and writes per-ticker sentiment with a confidence level and reasoning.
Extracted tickers are checked against the active-symbol list and the article text itself. Hallucinated symbols get dropped, PR-wire spam domains are blocked outright, and near-duplicates collapse into one row.
Only articles that clear the relevance floor and carry at least one verified ticker reach the API. Fewer than 1 in 5 makes the cut, from 6,900+ monitored sources, so the noise never reaches your parser.
A real row, from the feed, right now.
The newest article scoring 8 or higher, exactly as /api/news/ returns it. Fields not covered on this page are cut for length; the shape is the real contract.
{
"original": {
"uid": "2f9c5868891f5fcb",
"title": "Deckers Brands hits historic $1bn Q1 sales driven by HOKA, UGG",
"time_published": "2026-07-27T12:15:00Z",
"source_domain": "yahoo.com",
"url": "https://finance.yahoo.com/markets/stocks/articles/deckers-brands-hits-historic-1bn-112734988.html"
},
"enrichment": {
"relevance_score": 8,
"category": "earnings",
"ai_trading_insights": {
"ticker_analysis": [
{
"ticker": "DECK",
"impact_analysis": {
"sentiment": "positive",
"confidence": "medium",
"reasoning": "The article provides multiple forward-looking datapoints (gross margin forecast up, operating margin slightly above 21.5%, EPS guidance raised) tied to brand performance, which typically drives re-rating versus pure reca…"
}
},
"…"
]
}
}
} Official SDKs Typed clients, two languages.
Prefer a typed client to hand-rolled requests? The official Python and TypeScript SDKs wrap all nine endpoints and handle cursor pagination, retries on 429/5xx, typed errors, and rate-limit inspection for you. They take the same key and return the same data, so you can adopt one later or never.
Both read ALPHAI_API_KEY from the environment and target api.alphai.io by default. The Python client offers sync and async; the TypeScript client ships zero runtime dependencies.
Terminal dashboard: alphai-tui →
$ pip install alphai-sdk from alphai import Client
with Client() as c:
page = c.news.list(symbol="NVDA")
for a in page.results:
print(a.relevance_score, a.title) PyPI → Source → TypeScript Node · Edge · Deno · Bun $ npm install alphai-sdk import { AlphaAI } from "alphai-sdk";
const c = new AlphaAI();
const page = await c.news.list({
symbol: "NVDA",
});
for (const a of page.results)
console.log(a.enrichment.relevance_score); npm → Source → Built with the SDK: alphai-news-to-email — a deployable news-to-email digest in Python · alphai-sdk-ts-examples — runnable TypeScript scripts covering the whole API surface. All examples →
Authenticate with Authorization: Bearer ak_live_… . Quotas have two layers, a per-minute burst cap and a per-day volume cap, both counted against your account rather than per key. Free is for evaluation and non-commercial use; production needs a paid plan. The same key also authenticates the MCP server for headless agents (send it as a Bearer header there too); interactive AI clients connect to MCP over OAuth instead.
Articles reach the feed later than their publish time: general news lands a median of ~33 minutes after publication, SEC filings in 5 to 9 minutes (measured 2026-07-25). A poller that tracks time_published therefore misses late arrivals. sort=ingested returns rows in the order they entered the feed instead: call once without a cursor, store the next_cursor , then repeat the call with it. In this mode every response carries a cursor, and an empty results list means you are caught up. All filters work unchanged. Responses are cached for 60 seconds, so polling more often than once a minute buys nothing. Keep passing sort=ingested on every call: each mode issues its own cursors, and replaying one into the other mode returns a 400 rather than silently restarting at the top of the feed. The same mode works on /api/news/insider/ , where the gap is widest: a Form 4 is filed days after the trade it reports, so a new event often enters the feed already below the newest page.
# pip install alphai-sdk
import time
from alphai import Client
cursor = None # persist this between runs
with Client() as client: # reads $ALPHAI_API_KEY
while True:
page = client.news.list(sort="ingested", cursor=cursor)
for article in page.results:
print(f"[{article.relevance_score}/10] {article.title}")
cursor = page.next_cursor # always present in this mode
time.sleep(900) # 96 polls/day fits the Free daily cap That is the official Python SDK doing the cursor bookkeeping. There is a TypeScript one with the same sort option, and the endpoint stays a plain GET if you would rather carry no dependency.
Every news item ships an enrichment block produced by an LLM pass over the article. These are the four fields you'll use most.
Scores rate the article , not the company it mentions: a recap of a mega-cap's week-old earnings scores low, while a small-cap's fresh FDA approval scores high. Scoring is deterministic: the same article always gets the same score, and SEC Form 4 filings are scored from the transaction itself (size, buy vs. sell, 10b5-1 plan or not) rather than by the model.
Need the formal contract? The interactive API reference lives at /api/schema/swagger-ui/ (Swagger UI, with the raw OpenAPI spec linked at the top).
How much trading value the article itself carries.
1–2 · no trading relevance — lifestyle, solicitations, short-interest scrapes.
3–4 · derivative content about already-known events — opinion pieces, recaps, 13F stake writeups, market wraps.
5–6 · macro/sector datapoints with read-across; minor-but-real company news.
7–8 · real company news — analyst actions with a thesis, settlements, clinical data, same-day moves with a fresh catalyst.
9–10 · primary, material, newly disclosed — M&A, prints with numbers, approvals, major contracts.
The /api/news/ feed defaults to ≥ 4 ; pass ?min_relevance=N to relax or tighten the filter.
Per-ticker price impact direction. Not earnings-positive vs earnings-negative — this is "where will the chart move on this story."
Each ticker block also carries impact_analysis.confidence ( high / medium / low ) and a free-text reasoning field, which helps when you filter an automation pipeline down to confident bullish or bearish signals.
One topic bucket per article. Use ?category=earnings (CSV or repeated for OR-matching) on /api/news/ to filter the feed, or ?exclude_categories= to drop buckets you never trade on.
The numbers are article counts from the last 7 days of the live feed, refreshed every 15 minutes.
Tickers extracted from the article. We only retain symbols that exist in our /api/symbols/ list of active symbols (US equities, cryptocurrencies, and foreign listings), so articles that don't mention any active ticker are filtered out before they reach the API. US equities use the bare symbol, cryptocurrencies use a -USD suffix like BTC-USD , and foreign listings use a Yahoo suffix like VOD.L . Each ticker also gets a dedicated ticker_analysis entry with sentiment and reasoning, so multi-ticker stories don't collapse to a single read.
One real event reported by N outlets is N rows by default. Add ?collapse=story on /api/news/ to get one row per story instead — the representative article carries story_id (its own UID, resolvable via /api/news/ {uid} / ), sources_count (how many distinct outlets ran it, usually 1 because most stories are carried by a single outlet), and sources (the distinct domains). These fields are null in the default feed. Filters like symbol and category apply to the representative row, so a story whose representative doesn't match is omitted.
A dedicated worker tails SEC EDGAR and turns every Form 4 filing (officer and director buys and sells) into a feed row shortly after it's filed. One row covers one economic event: all of a filing's non-derivative trades of the same type and holding form, aggregated into a single article. The JSON shape matches every other article, so your parser doesn't need a special case.
Form 4 summaries are deterministic templates built from the filing itself: insider name, role, transaction code, total shares, price (volume-weighted average and executed range when a sale runs across tranches), and total value. There is no LLM paraphrase between you and the disclosure. The relevance score is computed from the event itself — its total size, buy vs. sell, 10b5-1 plan or not — so it's reproducible.
GET /api/news/insider/ — the Form 4 insider feed, cursor-paginated, ?symbol= to scope.
GET /api/news/?category=insider — the same rows through the main feed's category filter.
GET /api/symbols/ {ticker} /insider-summary/ — 30-day rollup: buys vs. sells, dollar volumes, % under 10b5-1 plans, most active insiders.
GET /api/symbols/ {ticker} /sentiment-summary/ — its sibling: 7-day bullish / neutral / bearish counts from the news flow.
The insider trading API, in detail →
{
"ticker": "NVDA",
"days": 30,
"total_transactions": 14,
"buy_count": 2,
"sell_count": 12,
"buy_value_usd": "1240000.00",
"sell_value_usd": "224580213.05",
"pct_10b5_1": 64,
"top_insiders": [
{
"name": "STEVENS MARK A",
"title": "Director",
"transaction_count": 3,
"net_value": "-221102600.00"
},
{
"name": "Dabiri John",
"title": "Director",
"transaction_count": 1,
"net_value": "-133750.00"
}
]
} Value fields are decimal strings in USD; buy_value_usd / sell_value_usd are null when the window has no buys or sells. net_value is buys minus sells per insider; title can be empty when the filer reports as a director without an officer role.
What an API-connected bot can do.
From watchlist alerts to Telegram bots, it's the same JSON in whatever stack you run. The API stays out of your way.
"Claude, what moved NVDA today?"
02 REST API "Send only high-relevance earnings news for my watchlist."
/api/news/?category=earnings · min_relevance=7
03 REST API "Build a Telegram bot for market-moving news in 10 minutes."
/api/news/trending/ + your bot framework
04 MCP "Any insider buying in my watchlist this week?"
alphai_insider_news → MCP · /api/news/insider/
05 MCP "Feed my agent relevance-scored financial news via MCP."
OAuth 2.1

[truncated]
