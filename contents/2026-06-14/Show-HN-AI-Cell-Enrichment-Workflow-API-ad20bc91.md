---
source: "https://ampledata.io"
hn_url: "https://news.ycombinator.com/item?id=48532372"
title: "Show HN: AI Cell Enrichment Workflow API"
article_title: "AmpleData - AI-Powered Data Enrichment for Any Dataset"
author: "blagoysimandoff"
captured_at: "2026-06-14T21:32:40Z"
capture_tool: "hn-digest"
hn_id: 48532372
score: 1
comments: 0
posted_at: "2026-06-14T20:36:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Cell Enrichment Workflow API

- HN: [48532372](https://news.ycombinator.com/item?id=48532372)
- Source: [ampledata.io](https://ampledata.io)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T20:36:23Z

## Translation

タイトル: HN を表示: AI 細胞濃縮ワークフロー API
記事のタイトル: AmpleData - あらゆるデータセットに対する AI を活用したデータ強化
説明: AI であらゆるデータセットを強化します。 CSV をアップロードし、列を選択すると、AmpleData が Web を検索して不足しているデータを自動的に埋めます。コードは必要ありません。リード、企業、製品などに使用できます。
HN テキスト: こんにちは、HackerNews!
最近、AI 強化プラットフォームであるampledata に API + ドキュメントを追加しました。アイデアはシンプルです。テーブル内の一連の行を指定すると、API 経由で ampledata にアクセスすると、Web を検索して引用と正確な情報源を提供することで、ワークフローが実行され、選択した列でテーブルが強化されます。これは、引用されたバルク エンリッチメント プラットフォームと考えることができます。 Ampledata も OSS なので、自由にセルフホストして独自の API キーを使用してください。あるいは、単にテクノロジーに興味がある場合でも。

記事本文:
AmpleData - あらゆるデータセットに対する AI を活用したデータ強化
API 価格ドキュメント 今すぐ試してみる → ✶ スプレッドシートの AI Research Worker 行の貼り付けを停止
ChatGPTに。
AmpleData は、スプレッドシートのすべての行を調査し、ソース間の引用、信頼スコア、競合解決を調査します。
エンリッチされたセルをクリックすると、出典の引用が表示されます。
AmpleData にリスト、会社、製品、論文、仕事、不動産など、名前のあるものを与えます。記入してほしい欄を平易な英語で説明してください:「価格設定に関する創業者の最新の公式声明」「この論文が撤回されたかどうか」「リモートワークに対する会社の表明されたスタンス」。 AmpleData は、Web 検索とクロールをディスパッチし、LLM で構造化された回答を抽出し、ソース間の競合を解決して、信頼できるクリーンアップされたシートを返します。すべてのセルは回答の出所にリンクされています。
ブラウザ上でセルが入力される様子をリアルタイムで観察します
各行が Web にディスパッチされ、ソースがクロールされ、LLM が回答を抽出し、競合が解決され、ソース URL と信頼スコアが添付された結果がセルに表示されます。
最長の機能リストではありません。最も特徴的なもの。
「AIが生成した」ものではありません。研究しました。任意のセルをクリックすると、ソース URL、抽出されたスニペット、および推論が表示されます。ソースが間違っている場合は、どれが間違っているかがわかります。
行動できる信頼スコア
AmpleData は、それが確実な場合とそうでない場合を教えてくれます。信頼性順に並べ替え、弱いセルをスポットチェックし、洗練されたプロンプトで列を再実行します。
シートごとではなく、セルごとの価格設定
あなたが豊かにしたものに対してのみ支払います。最小単位、シートごとのライセンス、年間契約はありません。同等のツールの料金の数分の一です。
独自のコードから直接強化する
API キーを生成し、HTTP 経由で同じエンリッチメント エンジンを駆動します。ループ内に UI はありません。 cronに接続します

仕事、パイプライン、または独自の製品。
# 1. API キーで認証する
エクスポート AMPLEDATA_KEY= "sk_live_..."
# 2. ソースのエンリッチメント実行を開始する
カール https://api.ampledata.ai/api/v1/sources/$SOURCE_ID/enrich \
-H "認可: ベアラー $AMPLEDATA_KEY" \
-H "コンテンツ タイプ: application/json" \
-d '{
"key_columns" : [ "会社" ],
"columns_metadata" : [
{ "名前" : "業界" 、 "タイプ" : "文字列" 、 "ジョブタイプ" : "エンリッチメント" }
】
}'
# 3. 引用された、信頼度スコア付けされたセルを撤回する
カール https://api.ampledata.ai/api/v1/jobs/$JOB_ID/results \
-H "認可: ベアラー $AMPLEDATA_KEY"
# => [{ "キー": "ストライプ.com",
# "抽出データ": { "業界": "フィンテック" },
# "信頼": { "業界": { "スコア": 0.9 } },
# "sources": ["https://..."] }] 比較 どのように比較しますか?
AmpleData は行全体で並行して実行され、すべての回答を引用し、独自の信頼度をスコアリングし、URL の幻覚を引き起こしません。それは、友人に尋ねることとアナリストを雇うことの違いです。
AmpleData は連絡先データベースを中心に構築されていません。これは、あらゆるリスト、あらゆる質問、あらゆるドメインを中心に構築されています。また、セルあたりのコストは約 10 分の 1 です。
データに適したプランを選択してください
すべてのプランには AI を活用したエンリッチメントが含まれています。いつでもアップグレードまたはダウングレードできます。
コミットメントなしで AmpleData をお試しください。
個人や小規模プロジェクトに最適です。
より高度なエンリッチメントのニーズを持つ成長中のチーム向け。
データドリブン組織向けの大容量エンリッチメント。
本当の質問。正直な答え。
お問い合わせください ご連絡をお待ちしております
AmpleData について質問がありますか?デモを見たり、エンタープライズ オプションについて話し合ったりしたいですか?ご連絡ください — 24 時間以内にご連絡いたします。
フォームにご記入ください。すぐにご連絡いたします。
私たちはあなたのプライバシーを尊重します。あなたの情報は決して共有されません。
研究者向けに構築
一度に 1 行ずつ。
リストをアップロードします。列を定義します。ゲットバー

ck が引用した、信頼度スコア付けされたデータを数分で取得できます。
ChatGPT に行を貼り付けることにうんざりした人が作成しました。

## Original Extract

Enrich any dataset with AI. Upload a CSV, pick your columns, and AmpleData searches the web to fill in missing data automatically. No code required. Works for leads, companies, products, and more.

Hi HackerNews!
I've recently added an API + Docs to ampledata - my AI enrichment platform. Idea is simple. Given a set of rows in a table you can hit ampledata via an api and a workflow will run to enrich the table with columns of your choosing by searching the web and giving you citations + exact sources. You can think of it as a cited bulk enrichment platform. Ampledata is also OSS so feel free to self host it and use your own api keys. Or if you are just interested in the tech.

AmpleData - AI-Powered Data Enrichment for Any Dataset
API Pricing Docs Try it out now → ✶ AI research worker for your spreadsheet Stop pasting rows
into ChatGPT.
AmpleData researches every row of your spreadsheet, with citations , confidence scores , and conflict resolution across sources.
Click any enriched cell to see its source citation
Give AmpleData a list, companies, products, papers, jobs, properties, anything with a name. Describe the columns you want filled in, in plain English: "the founder's most recent public statement on pricing," "whether this paper has been retracted," "the company's stated stance on remote work." AmpleData dispatches web search and crawl, extracts structured answers with an LLM, resolves conflicts across sources, and returns a cleaned-up sheet you can trust, every cell linked back to where the answer came from.
Watch cells fill in, live, in your browser
Each row gets dispatched to the web, sources get crawled, an LLM extracts the answer, conflicts are resolved, and the result lands in the cell, with a source URL and confidence score attached.
Not the longest feature list. The most distinctive one.
Not 'AI-generated.' Researched. Click any cell to see the source URL, the extracted snippet, and the reasoning. If a source is wrong, you'll know which one.
Confidence scores you can act on
AmpleData tells you when it's sure and when it isn't. Sort by confidence, spot-check the weak cells, and re-run a column with a refined prompt.
Per-cell pricing, not per-seat
Pay only for what you enrich. No minimum, no per-seat license, no annual contract. A fraction of what comparable tools charge.
Enrich straight from your own code
Generate an API key and drive the same enrichment engine over HTTP, no UI in the loop. Wire it into a cron job, a pipeline, or your own product.
# 1. Authenticate with your API key
export AMPLEDATA_KEY= "sk_live_..."
# 2. Kick off an enrichment run for a source
curl https://api.ampledata.ai/api/v1/sources/$SOURCE_ID/enrich \
-H "Authorization: Bearer $AMPLEDATA_KEY" \
-H "Content-Type: application/json" \
-d '{
"key_columns" : [ "company" ],
"columns_metadata" : [
{ "name" : "industry" , "type" : "string" , "job_type" : "enrichment" }
]
}'
# 3. Pull back cited, confidence-scored cells
curl https://api.ampledata.ai/api/v1/jobs/$JOB_ID/results \
-H "Authorization: Bearer $AMPLEDATA_KEY"
# => [{ "key": "stripe.com",
# "extracted_data": { "industry": "Fintech" },
# "confidence": { "industry": { "score": 0.9 } },
# "sources": ["https://..."] }] Comparisons How does it compare?
AmpleData runs in parallel across rows, cites every answer, scores its own confidence, and doesn't hallucinate URLs. It's the difference between asking a friend and hiring an analyst.
AmpleData isn't built around a contacts database. It's built around any list, any question, any domain. And it costs about 1/10th as much per cell.
Choose the plan that fits your data
All plans include AI-powered enrichment. Upgrade or downgrade at any time.
Try AmpleData with no commitment.
Perfect for individuals and small projects.
For growing teams with higher enrichment needs.
High-volume enrichment for data-driven orgs.
Real questions. Honest answers.
Get in touch We'd love to hear from you
Have a question about AmpleData? Want to see a demo or discuss enterprise options? Reach out — we'll get back to you within 24 hours.
Fill out the form and we'll be in touch soon.
We respect your privacy. Your information is never shared.
Built for people who do research
one row at a time.
Upload a list. Define your columns. Get back cited, confidence-scored data in minutes.
Made by a person who got tired of pasting rows into ChatGPT.
