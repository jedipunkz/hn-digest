---
source: "https://barrasso.me/posts/2026-06-23-shipping-ai-generated-code-safely-with-static-sites/"
hn_url: "https://news.ycombinator.com/item?id=48672329"
title: "Shipping AI-Generated Code Safely with Static Sites"
article_title: "Shipping AI-Generated Code Safely with Static Sites - Tom Barrasso"
author: "podlp"
captured_at: "2026-06-25T12:45:06Z"
capture_tool: "hn-digest"
hn_id: 48672329
score: 1
comments: 0
posted_at: "2026-06-25T12:23:15Z"
tags:
  - hacker-news
  - translated
---

# Shipping AI-Generated Code Safely with Static Sites

- HN: [48672329](https://news.ycombinator.com/item?id=48672329)
- Source: [barrasso.me](https://barrasso.me/posts/2026-06-23-shipping-ai-generated-code-safely-with-static-sites/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T12:23:15Z

## Translation

タイトル: 静的サイトを使用した AI 生成コードの安全な配布
記事のタイトル: AI で生成されたコードを静的サイトで安全に配布する - Tom Barrasso
説明: コーディング エージェントがパイプラインを作成しました。静電気の発生により安全に輸送できました。 Claude、Python、DuckDB、Hugo を使用して毎日更新される不動産サイトを構築します。

記事本文:
AI 生成コードを静的サイトで安全に配布 - Tom Barrasso AI 生成コードを静的サイトで安全に配布
私は何年もの間、Hugo を使用して Web サイトを構築してきました。このブログ 、 Cloud Phone for Developers 、そして現在は私の個人的な不動産 Web サイト、 Unique Homes Massachusetts を支えています。 PodLP Podcast API の初期バージョンでさえ、静的 Web サイトとして設計されており、Podcatcher が RSS フィードをクロールし、フィードのコンテンツが変更されると、S3 に保存され CloudFront によってホストされる JSON フラグメントとして API ペイロードを事前に計算しました。
簡単に言えば、静的に生成できる Web サイトはすべて、 である必要があります。 Claude Code のようなコーディング エージェントや、Cloudflare Workers のようなサーバーレス サービスとしてのプラットフォーム (PaaS) ホスティング サービスを利用すれば、頻繁に更新され、情報密度が高く、カスタム設計された、手頃な価格の静的 Web サイトを構築することは簡単です。 Unique Homes MA とそのアーキテクチャ、および Cloudflare ページで再現可能な毎日のビルドの構築時間を構築するために AI をどのように検証して移行したかについて、さらに詳しい内容を共有します。
不動産分野でニッチ市場を開拓する
#
私の故郷であるマサチューセッツ州 (MA) は、最初の 13 の植民地の 1 つとして、数百年前に遡る長く豊かな歴史を持っています。マサチューセッツ州は、その規模と人口にもかかわらず、建築上の影響力に関してはその重みを超えています。マサチューセッツ州は、バイオテクノロジーやテクノロジー、教育によって経済が健全に成長しているため、不動産業を営むのに最適な州でもあります。私の専門分野として、私は MA のユニークな歴史的目録に焦点を当てることにしました。
マサチューセッツ州の歴史と建築
#
ベイ州は住宅と建築に多大な影響を与えており、マサチューセッツ州発祥のスタイルで今日まで多くの家が建てられ続けています。いくつかの注目すべきスタイルは次のとおりです。
マサチューセッツ州ブロックトンのブロックトンハイツにある最古の家、古典的なケープコッド（1910 年のポストカード）。プ

blic ドメイン、ウィキメディア コモンズ経由。 17 世紀後半に初めて建てられたケープ コッド スタイルは、マサチューセッツ州の沿岸地域に由来すると考えられています。ケープ コッド ハウスは、雪を落とすための急な屋根、中央の煙突、グレーのトーンで磨かれた象徴的な杉の屋根板を備えた、米国で最も有名なスタイルの 1 つです。
コネチカット州ギルフォードにあるコンフォート スター ハウスは、植民地時代の塩箱です。写真は米国歴史的建造物調査 (国立公園局) より、パブリック ドメインです。塩を保管するために使用される木箱にちなんで名付けられたソルトボックスは、非対称のケープです。後部の急な屋根のラインが後部の平屋まで続いています。
マサチューセッツ州ケンブリッジにある 3 階建ての建物で、各階にポーチが積み上げられています。写真提供：Bcorr、CC BY-SA 3.0。移民労働者を収容するためにウースター、ボストン、フォールリバーなどのコミュニティに建てられたトリプルデッカーは、前後のポーチが一致するもので、手頃な価格の高密度住宅への MA のもう 1 つの貢献です。
マサチューセッツ州レキシントンのファイブ フィールズ地区にある TAC 設計の住宅 (1954 年)。写真提供：Fothergilla、CC BY-SA 4.0。バウハウスの創設者であるウォルター グロピウスは、1937 年にハーバード大学大学院デザイン大学院に入学し、マサチューセッツにモダニズムをもたらしました。カール コッホや建築家協同組合 (TAC) のメンバーとともに、モダニズムはリンカーンのようなボストン郊外で形を整えました。
不動産データの入門書
#
Zillow や Redfin などの Web サイトは、Multiple Listing Services (MLS) を通じて不動産物件情報にアクセスし、Internet Data Exchange (IDX) 形式を使用してデータを共有し、仲介業者に提供します。技術的な観点から見ると、データは通常、パイプ文字 ( | ) で区切られたカンマ区切り値 (CSV) 形式で提供され、LOT_SIZE などの大文字の列名が付けられます。ブール値は Yes および No として保存され、IDX 消費者はオフィスレベルのオプトアウト要求を受け入れる必要があります。

DisplayOnInternet = No や ShowAddress = No などの可視性コントロール。
幸いなことに、これらのファイルは形式とサイズの両方で管理可能です。過去 1 年間にマサチューセッツ州で販売されたすべての一戸建て住宅の合計記録は、わずか約 65 MB です。課題は、さまざまなソースからのデータを正確かつ最小限の損失で確実につなぎ合わせることにあります。
IDX フィードの他に、不動産標準化機構 (RESO) は、公開仕様を含む RESO Web API の標準を維持しています。 RESO Web API は RESTful で、リアルタイムでデータを提供し、読み取りと書き込みをサポートし、JSON 形式で交換される OAuth 2.0 と OpenID Connect を使用します。 RESO データ ディクショナリは、ローカライズされた MLS のバリエーションを避けるために、BedroomsTotal などのフィールドを標準化します。
IDX フィードと比較して、RESO Web API にはより多くの (非公開の) データ フィールドがありますが、コストが高くつき、コンプライアンス要件も厳しくなります。リード生成に関しては、公開されている IDX データが、見込み客を惹きつけるのに十分な情報を提供します。
UniqueHomesMA.com のアーキテクチャ Unique Homes MA は、MLS PIN、マサチューセッツ地理情報システム (GIS)、マサチューセッツ文化資源情報システム (MACRIS)、国家歴史登録簿 (NRHP)、ウィキデータなどのソースから区画、住所、証書、場所、歴史的指定、その他のデータを取得します。次に、町の接尾辞 (例: 「Southboro」対「Southborough」)、単位の指定、および道路の略称 (例: 「St」対「Street」) を含む住所を正規化する必要があります。
緯度、経度、半径を提供するソースもあれば、土地区画の境界を識別するための特定の境界ボックスを提供するソースもあります。これらすべてをテストしてつなぎ合わせるのに必要な構成とロジックは非常に面倒ですが、これは間違いなくコーディングアシスタントに喜んで任せたいタイプの仕事です。
ほとんどの

e Unique Homes MA パイプラインは Python で書かれています。 Python には、データの操作と変換を処理するためのライブラリの広範なカタログがあります。 DuckDB は、リストとプロパティ データの保存とクエリに使用されます。これは埋め込み型 (実行するサーバーがないため安価)、オープンソースであり、CSV、Parquet、および時系列データを効率的に読み取ります。その列指向エンジンは、SQLite よりもはるかに優れた完全なテーブル スキャンを処理するため、基礎となるデータを再構築することなく、さまざまな分類を迅速にテストできます。最後に、最初に述べたように、私は静的サイト ジェネレーターとして Hugo を使用しています。これは、使い慣れており、高速で、柔軟性が高いためです。エンドツーエンドのパイプラインには最大 5 分かかり、そのうち最大 30 秒が Hugo --minify で 13,000 以上のページを生成します。
コストと鮮度のバランスを考慮して、Unique Homes MA は、EC2 インスタンスをトリガーして systemd サービスを実行する EventBridge スケジューラーを使用して毎日 cron ジョブを実行します。パイプラインと Hugo --minify のビルドが完了すると、サービスは wranglerdeploy を使用して HTML ドキュメントを Cloudflare Pages に公開し、サービス自体を停止します。データは Elastic Block Storage (EBS) に保存され、実行実行全体で状態が保持されます。
t4g.medium EC2 インスタンスの場合は 0.034 ドル/時間、EBS では 1 GB 月あたり 0.05 ドル、S3 (バックアップ用) では 1 GB 月あたり 0.023 ドルで、合計コストは 1 か月あたり ~2.35 ドル、または .com ドメイン登録料金を含めると年間約 39 ドルになります。これは、年間ライセンス更新料、MLSの会費、ブローカーデスク手数料などの一般的な不動産費用を大幅に下回ります。これを、月額 39 ドルから始まる AgentWebsite のようなサービスと比較してください。
私は正規表現 (regex) を書くのが好きではありませんでしたが、正規表現はパターン マッチングや生のテキストから情報を抽出するための強力なツールです。 Unique Homes MA は、正規表現を広範囲に使用して、物件仲介業者のコメント内の単語やフレーズを照合します。クロードは提案し、

過去のデータに対して各式を検証し、外れ値を特定し、カバレッジと精度を継続的に改善します。端末からの claude -p の呼び出しとは異なり、正規表現は決定的であり、ネットワーク アクセスを必要とせず、実行コストは (実質的に) かかりません。他のパターン認識技術と組み合わせることで、Unique Homes MA は 1 万軒以上の住宅をスタイル、建築者、特徴ごとに迅速に識別、カタログ化、並べ替えることができます。
_ARTS_CRAFTS_STYLE_RE: 最終 = re 。コンパイル(
r "\barts\s*(?:&|and)\s*crafts[-\s]+"
r "(?:(?!または\b|および\b|for\b|with\b|スタジオ|部屋?|スペース|エリア|オフィス|隅|コーナー|"
r "駅|クローゼット|ロフト|書斎|テーブル)[a-z][\w-]*\s+)?"
r "(?:家|家|バンガロー|コロニアル|コテージ|邸宅|岬|ビクトリア朝|チューダー|フォースクエア|"
r "スタイル?|運動|時代|時代|デザイン|建築\w+|美学|性格|魅力|"
r "影響を受けましたか?|インスピレーション|復活|インテリア|木工|構築[\s-]?ins?|詳細\w*|トリム|"
r "精巧な細工|宝石|美しさ\w+|傑作|血統|才能|壮大|要素?)\b" ,
再。無視してください、
)
たとえば、この厄介な正規表現は、プロパティを Arts and Crafts スタイルとして分類するための信号として使用されます。建築年や MLS スタイル コードなどの他のシグナルと組み合わせると、「美術品や工芸品やオフィス用の追加スペース」などの説明は無視され、「F.A. デイによって建てられた歴史的なアーツ & クラフツ コロニアル」として説明される物件が分類されます。
注: 上場ブローカーの発言は一貫性がないことで知られています。文字数制限により、エージェントはニックネーム (例: 「ナブ レイク」対「ナブナセット レイク」) や略語 (例: 「強制熱風」の代わりに「FHA」) を使用することが推奨されます。誇張 (例: 「湖畔」対「湖の眺め」) やタイプミス (例: 「シアーズ & ロバック」) が見られることは珍しくありません。正規表現や決定論的なソリューションは、複数の環境で完全に機能することはありません。

数万件のリスト。正確性と正確性を確保するために、ルールはライブ リストに対して定期的に検証される必要があります。
静的 Web サイトでは、侵害されるサーバーやインターネット トラフィックにさらされるデータベースがないため、コストが低くなり、セキュリティ上の懸念も少なくなります。ジャストインタイム処理を行わなくても本質的に高速です。
静的サイトは、AWS S3 や Cloudflare R2 などのバケット ストレージ、GitHub Pages や Cloudflare Pages などのホスティング サービス、さらには Apache や NGINX を実行する Raspberry Pi など、ほぼどこでもホストできます。それぞれに、コスト、コンテンツの鮮度、複雑さ、メンテナンスの点でトレードオフがあります。
Hugo を選んだのは、使い慣れていて、ポータブルで完成度が高く、軽量で高速であるためです。とはいえ、静的サイト ジェネレーターは多数あるので、使い慣れたものを選んでください。
静的サイト用に AI が生成したコードは、本質的に安全というわけではありません。静的なコンテキストで「安全に出荷」するには、「良好」とはどのようなものかをエージェントに伝え、その結果に対してテストを作成し、反復し、構築して、出荷する必要があります。トークンコストを低く抑え、構築を決定論的に行うため、LLM は構築時にリストを評価および分類しません。代わりに、テストによってゲートされたパイプライン内でルールを作成し、最終的に一連のアーティファクト (つまり、Markdown と HTML) を生成します。
Unique Homes MA に変更を加えるたびに、次のような多くのチェックが実行されます。
pytest 単体テスト (VCR のようなネットワーク再生テストを含む)
mypy と ruff のリンティングとコードのフォーマット
Hugo --minify ビルドの完了
動的 Web サイトはクラッシュしたり、バグがある場合に HTTP 500 を返したりする可能性がありますが、静的サイトは静かに「失敗」する可能性があります。訪問者がエラーページにアクセスして文句を言うことはありません。その代わり、物件が欠落していることに気づかず、物件がいつ契約されるのかも知りません。ユーザーが壊れたリンクをクリックすると、不一致に気づく可能性があります

ZillowかRedfinを選び、最終的には去ります。サイレント障害は障害ではありませんが、別の戦略が必要です。つまり、出荷前のビルド時の検出と検証です。 Unique Homes MA のビルドごとに AWS SES 経由でメールが送信されるので、壊れた情報や古い情報を静かに提供するのではなく、ビルドが成功するか失敗するか停止するかを常に監視できます。
Cloudflare Workers のような PaaS サービスを使用すると、静的 Web サイトが 100% 静的である必要はありません。 Workers を使用すると、問い合わせフォームの処理やメール リストの購読などの小さな関数を作成できます。サーバーレス プラットフォームは次のものと統合できます。
Cloudflare 電子メール – 従業員の給与 (5 ドル/月) + 0.35 ドル/1,000 メール
AWS Simple Email Service (SES) – $0.10/1000 メール (価格を参照)
サードパーティ API – Resend、SendGrid など
Webhook – Discord、Slack など
動的 Web サイトは主にセッション Cookie を使用して詳細なユーザー情報を保存しますが、静的 Web サイトでも localStorage 、 sessionStorage 、 IndexedDB 、URL クエリ パラメーター、およびその他のストレージ テクノロジを使用して訪問者データを「記憶」できます。
Unique Homes MA への訪問者が「詳細情報をリクエスト」をクリックすると、URLSearchParams から前のリストに関する情報を取得する静的ページが表示されます。ホームページには、 localStorage に保存されている「最近表示した」プロパティも表示されるため、訪問者はすぐに見つけることができます。

[切り捨てられた]

## Original Extract

Coding agents wrote the pipeline; static generation made it safe to ship. Building a daily-updated real estate site with Claude, Python, DuckDB, and Hugo.

Shipping AI-Generated Code Safely with Static Sites - Tom Barrasso Shipping AI-Generated Code Safely with Static Sites
For years, I’ve built websites using Hugo . It powers this blog , Cloud Phone for Developers , and now my personal real estate website, Unique Homes Massachusetts . Even early versions of the PodLP Podcast API were architected as static websites, where the podcatcher crawled RSS feeds and, when the feed content changed, pre-computed API payloads as JSON fragments stored in S3 and hosted by CloudFront.
Simply put, any website that can be statically-generated, should be . With coding agents like Claude Code , and serverless platform-as-a-service (PaaS) hosting services like Cloudflare Workers, it’s trivial to build frequently-updated, information-dense, custom-designed, affordable static websites. I’ll share more context on Unique Homes MA, its architecture, and how I’ve validated and shifted AI to build time for reproducible, daily builds on Cloudflare Pages.
Carving out a niche in real estate
#
As one of the original thirteen colonies, my home state of Massachusetts (MA) has a long and rich history dating back hundreds of years. Despite its size and population, MA punches above its weight when it comes to architectural influence. MA is also a great state to practice real estate, because it has a healthy and growing economy driven by biotech/tech and education. For my niche, I chose to focus on MA’s unique historical inventory.
Massachusetts History & Architecture
#
The Bay State has had an outsized impact on housing and architecture, where many homes continue to be built to this day in styles originating in MA. A few notable styles include:
The oldest house in Brockton Heights, Brockton, MA, a classic Cape Cod (1910 postcard). Public domain , via Wikimedia Commons. First built in the late 17th century, the Cape Cod style is attributed to Massachusetts’ coastal region. The Cape Cod House is one of the US’ most recognizable styles, with its steep roof to shed snow, central chimney, and iconic cedar shingles worn to a gray tone.
The Comfort Starr House in Guilford, CT, a colonial saltbox. Photo from the Historic American Buildings Survey (National Park Service), public domain. Named after wooden boxes used to store salt, the Saltbox is an asymmetric Cape: a steep rear roofline that sweeps down to a single story at the back.
A triple-decker in Cambridge, MA, with stacked porches on each floor. Photo by Bcorr , CC BY-SA 3.0 . Built in communities like Worcester, Boston, and Fall River to house immigrant workers, the Triple-Decker with matching front and rear porches is another MA contribution to affordable, dense housing.
A TAC-designed house (1954) in the Five Fields neighborhood of Lexington, MA. Photo by Fothergilla , CC BY-SA 4.0 . Walter Gropius, founder of the Bauhaus, brought Modernism to Massachusetts when he joined the Harvard Graduate School of Design in 1937. Along with Carl Koch and fellow members of The Architects Collaborative (TAC), Modernism took shape in Boston suburbs like Lincoln.
A primer on real estate data
#
Websites like Zillow and Redfin access property listings through Multiple Listing Services (MLS), which share and vend data to brokerages using the Internet Data Exchange (IDX) format. From a technical perspective, data is typically served in comma-separated value (CSV) format, delimited using the pipe character ( | ), with column names like LOT_SIZE in upper-snake case. Booleans are stored as Yes and No , and IDX consumers need to honor office-level opt out requests or visibility controls like DisplayOnInternet = No or ShowAddress = No .
Fortunately, these files are manageable in both format and size. The total records for all single-family homes sold in MA in the past year are only ~65MB. The challenge is in reliably stitching together data from a variety of sources accurately and with minimal loss.
Other than IDX feeds, the Real Estate Standards Organization (RESO) maintains a standard for the RESO Web API , including public specifications . The RESO Web API is RESTful, serves data in real time, supports reads and writes, and uses OAuth 2.0 & OpenID Connect exchanged in JSON format. The RESO Data Dictionary standardizes fields like BedroomsTotal to avoid localized MLS variations.
Compared to IDX feeds, the RESO Web API has more (often non-public) data fields, but comes at a greater cost and with stricter compliance requirements. For lead generation, public IDX data offers enough information to entice prospective buyers.
Architecture of UniqueHomesMA.com Unique Homes MA pulls parcels, address, deed, location, historical designation, and other data from sources including MLS PIN, Massachusetts Geographic Information Systems (GIS), Massachusetts Cultural Resource Information System (MACRIS), the National Registry of Historic Places (NRHP), and Wikidata. Next, addresses need to be normalized including town suffix (e.g. “Southboro” vs “Southborough”), unit designation, and road abbreviations (e.g. “St” vs “Street”).
Some sources provide latitude, longitude, and a radius, while others give specific bounding boxes to identify parcel boundaries. The configuration and logic needed to test and stitch all of this together gets very messy, and is certainly the type of job I’ll happily offload to a coding assistant.
Most of the Unique Homes MA pipelines are written in Python. Python has an extensive catalog of libraries for handling data manipulation and transformation. DuckDB is used to store and query listings and property data. It’s embedded (no server to run, so it’s cheap), open-source, and reads CSV, Parquet, and time-series data efficiently. Its columnar engine handles full table scans far better than SQLite, which lets me test different categorizations quickly without reshaping the underlying data. Lastly, as I mentioned at the start, I’m using Hugo as the static-site generator because it’s familiar, fast, and flexible. End-to-end pipeline takes ~5 minutes, ~30 seconds of which is hugo --minify generating 13,000+ pages.
Balancing cost and freshness, Unique Homes MA runs a daily cron job using an EventBridge Scheduler that triggers an EC2 instance to run a systemd service. When the pipeline and hugo --minify build finish, the service publishes HTML documents to Cloudflare Pages using wrangler deploy , then stops itself. Data is stored on Elastic Block Storage (EBS) to preserve state across execution runs.
At $0.034/hour for a t4g.medium EC2 instance, $0.05 per GB-month of EBS, and $0.023 per GB-month on S3 (for backups), the total cost is ~$2.35 per month , or about $39 per year including .com domain registration fees. That’s well below common real estate expenses like annual licensure renewal fees, MLS membership dues, or broker desk fees. Compare that to services like AgentWebsite that start at $39 per month .
I’ve never enjoyed writing regular expressions (regex), but they’re a powerful tool for pattern-matching and extracting information from raw text. Unique Homes MA uses regex extensively to match words and phrases in the listing broker remarks. Claude proposes and validates each expression against historical data, identifies outliers, and continuously improves coverage and accuracy. Unlike calls to claude -p from the terminal, regular expressions are deterministic, don’t require network access, and cost (effectively) nothing to execute. Combined with other pattern-recognition techniques, this allows Unique Homes MA to quickly identify, catalog, and sort over ten thousand homes by style, builder, and feature.
_ARTS_CRAFTS_STYLE_RE: Final = re . compile(
r "\barts\s*(?:&|and)\s*crafts[-\s]+"
r "(?:(?!or\b|and\b|for\b|with\b|studio|rooms?|space|area|office|nook|corner|"
r "station|closet|loft|den|table)[a-z][\w-]*\s+)?"
r "(?:home|house|bungalow|colonial|cottage|residence|cape|victorian|tudor|foursquare|"
r "styled?|movement|period|era|design|architectur\w+|aesthetic|character|charm|"
r "influenced?|inspired|revival|interior|woodwork|built[\s-]?ins?|detail\w*|trim|"
r "millwork|gem|beaut\w+|masterpiece|pedigree|flair|grandeur|elements?)\b" ,
re . IGNORECASE,
)
For example, this gnarly regular expression is used as a signal to classify a property as Arts and Crafts style . Combined with other signals like build year and MLS style codes, this categorizes properties described as “historic Arts & Crafts Colonial built by F.A. Day,” while ignoring descriptions like “extra space for arts & crafts or office.”
Note: listing broker remarks are notoriously inconsistent. Character count limits encourage agents to use nicknames (e.g. “Nab Lake” vs “Nabnasset Lake”) and abbreviations (e.g. “FHA” instead of “forced hot air”). It’s not uncommon to see exaggerations (e.g. “lakefront” vs “lake view”), or typos (e.g. “Sears & Robuck”). No regular expression or deterministic solution will work perfectly across tens of thousands of listings. Rules must be validated against live listings regularly for accuracy and precision.
Static websites have lower costs and fewer security concerns since there’s no server to compromise, and no database exposed to internet traffic. They’re inherently fast without just-in-time processing.
Static sites can be hosted just about anywhere including bucket storage like AWS S3 and Cloudflare R2, hosting services like GitHub Pages or Cloudflare Pages, or even a Raspberry Pi running Apache or NGINX. Each comes with trade-offs in terms of cost, content freshness, complexity, and maintenance.
I chose Hugo because I’m familiar with it and it’s portable, mature, lightweight, and fast . That said, there are dozens of static site generators, so pick one you’re familiar with.
AI-generated code for a static site isn’t inherently safer. “Safe to ship” in a static context still requires you to tell the agent what “good” looks like, write tests against that outcome, iterate, build, and ship. To keep token cost low and builds deterministic, an LLM doesn’t evaluate and classify listings at build time. Instead, it authors rules inside of a pipeline, gated by tests, that ultimately produce a set of artifacts (namely, Markdown and HTML).
Every change to Unique Homes MA runs a number of checks including:
pytest unit tests, including VCR-like network replay tests
mypy and ruff linting & code formatting
hugo --minify build completion
Although dynamic websites might crash or return an HTTP 500 if there’s a bug, static sites can silently “fail.” A visitor won’t hit an error page and complain. Instead, they won’t realize a listing is missing or know when a property goes under contract. Users might click a broken link, notice a discrepancy against Zillow or Redfin, and ultimately leave. Silent failures aren’t a blocker, but they do require a different strategy: detection and validation at build time, before anything ships . Each build of Unique Homes MA sends me an email through AWS SES so I can keep an eye on whether builds succeed, fail, or stall, instead of quietly serving broken or stale information.
With PaaS services like Cloudflare Workers, static websites don’t have to be 100% static. Workers lets you write small functions for things like handling a contact form or subscribing to an email list. Serverless platforms can then integrate with:
Cloudflare Email – Workers Paid ($5/m) + $0.35/1,000 emails
AWS Simple Email Service (SES) – $0.10/1000 emails (see pricing )
Third-party APIs – Resend, SendGrid, etc
Webhooks – Discord , Slack, etc
While dynamic websites primarily use session cookies to store detailed user information, static websites can still “remember” visitor data using localStorage , sessionStorage , indexedDB , URL query parameters, and other storage technologies.
When visitors to Unique Homes MA click “Request more info” they are taken to a static page that pulls information from URLSearchParams about the previous listing. The homepage also displays “Recently viewed” properties stored in localStorage , so visitors can quickly find

[truncated]
