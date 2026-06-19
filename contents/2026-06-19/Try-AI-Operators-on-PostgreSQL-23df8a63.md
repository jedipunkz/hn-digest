---
source: "https://samtsql.com/"
hn_url: "https://news.ycombinator.com/item?id=48601666"
title: "Try AI Operators on PostgreSQL"
article_title: "samtSQL | SQL + AI on Your Database"
author: "itrummer"
captured_at: "2026-06-19T18:42:15Z"
capture_tool: "hn-digest"
hn_id: 48601666
score: 1
comments: 0
posted_at: "2026-06-19T18:34:32Z"
tags:
  - hacker-news
  - translated
---

# Try AI Operators on PostgreSQL

- HN: [48601666](https://news.ycombinator.com/item?id=48601666)
- Source: [samtsql.com](https://samtsql.com/)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T18:34:32Z

## Translation

タイトル: PostgreSQL で AI オペレーターを試す
記事のタイトル: samtSQL |データベース上の SQL + AI
説明: PostgreSQL に接続し、samtSQL を介して組み込み AI 演算子を使用して SQL を実行します。

記事本文:
samtSQL |データベース上の SQL + AI
サムトSQL
探検する
データ型
マルチモーダルテーブルのセマンティック分析
クラウド データベース上で AI を使用して SQL を実行します。
samtSQL を使用すると、既存の PostgreSQL データベースに保存されているテキスト、画像、オーディオ、テーブルに対して AI 演算子を使用して SQL を作成できます。
平均価格を選択
ハウス広告から
WHERE year_built > 1980
かつ寝室数が 3 以上
そしてAIFILTER（写真、「家にはプールがあります」）。
結果
997,500ドル
平均2列目と4列目の上
仕組み
接続してクエリする
既存の PostgreSQL データベースに接続します。
画像、音声、またはテキスト データを含む CSV または ZIP ファイルをアップロードします。
REST API またはビジュアル インターフェイスを使用して AI オペレーターで SQL を実行します。
混合データのワークロードを 1 つの環境で取り込んで分析します。
テキスト、整数、浮動小数点値、ブール値などの標準 SQL データ型。
JPEG や PNG などの一般的な画像形式に対する AI 分析。
ドキュメントおよびフリーテキスト分析はリレーショナル テーブルに直接保存されます。
MP3やWAVなどの音声ファイルをAIオペレーターで解析します。
SQL を使用しながら、分析の意図を自然言語で表現します。
自然言語条件を使用してマルチモーダル データをフィルター処理します。
レコードをラベル名のみで分類します。
自然言語基準に基づいてスコアを割り当てます。
自然言語命令を使用して情報を変換または抽出します。
自然言語で定義された意味一致ルールを使用してデータセットを結合します。
アップロードされた各形式は、SQL + AI ワークフローのファーストクラスの入力として利用できます。
リレーショナル分析のための構造化された表形式のアップロード。
混合マルチモーダル アセットのアーカイブ コンテナー。
視覚的な分析タスク用の圧縮画像形式。
視覚的忠実度の高いロスレス画像フォーマット。
アップロードされたメディア データセットでサポートされる画像形式。
コンパクトに保存できる最新の Web 画像形式。
音声データの圧縮オーディオ形式。
非圧縮オーディオ形式の詳細

信号品質。
AI クレジット + 毎日のコンピューティング時間
すべてのプランで、PostgreSQL 上の完全な SQL + AI 演算子セットのロックが解除されます。各プランには、LLM 使用量に対する AI クレジットの毎月の割り当てと、クエリ コンピューティング時間の 1 日あたりの許容量がバンドルされています。
たまに探索するのに最適
通常の分析ワークフローに最適
*毎日の計算時間は 00:00 UTC にリセットされます。
代表的なユースケースを切り替えて、構文の強調表示を使用して各 SQL パターンを検査します。
アンケートのコメントから苦情を見つける
自由記述形式の回答で価格に関する懸念について言及した顧客をカウントします。
選択数(*)
FROM調査結果
WHERE AIFILTER(コメント、「顧客が価格について苦情を言っている」);
画像コレクションの分類
アップロードされた休日画像にセマンティック クラスを 1 つのクエリでタグ付けします。
SELECT AICLASSIFY(写真, 'ビーチの写真', '市内訪問', 'その他')
ホリデーイメージより;
アップロードされた音声の文字起こし
1 つのクエリで、テーブルに保存されている MP3 クリップをプレーンテキストのトランスクリプトに変換します。
SELECT AIMAP(content, 'この音声をプレーン テキストに書き起こします。') AS トランスクリプト
オーディオクリップから
WHERE ファイル名は '%.mp3' のようになります。
レビューでのセンチメントのランキング
コメントをポジティブスコアで並べ替えて、アウトリーチに優先順位を付けます。
選択 *
FROM調査結果
ORDER BY AISCORE(コメント、1、10、'ポジティブ');
カスタマーサポートチケットの要約
トリアージのために長いチケットを短い要約に凝縮します。
SELECT AIMAP(ticket_body, '顧客の問題を 2 つの文で要約してください')
サポートチケットから;
製品カタログ間でのエンティティの一致
セマンティック結合を使用して、独立したデータセットからの同等の製品を照合します。
選択 *
FROM ProductDB1 P1、ProductDB2 P2
WHERE AIJOIN(P1.製品名, P2.製品名, 'これは同じ製品です');
ドキュメントとガイド付きチュートリアル
エンドポイントごとの API ドキュメント、オペレーター仕様、ビジュアル インターフェイスと Python REST クライアントの両方の実践的なチュートリアルを読む

。

## Original Extract

Connect PostgreSQL and run SQL with built-in AI operators through samtSQL.

samtSQL | SQL + AI on Your Database
samt SQL
Explore
Data types
Semantic Analytics on Multimodal Tables
Run SQL with AI on your cloud database.
samtSQL lets you write SQL with AI operators over text, images, audio, and tables stored in your existing PostgreSQL database.
SELECT AVG(price)
FROM HouseAds
WHERE year_built > 1980
AND bedrooms >= 3
AND AIFILTER(photo, 'the house has a pool');
Result
$997,500
avg. over rows 2 and 4
How it works
Connect and query
Connect your existing PostgreSQL database.
Upload CSV or ZIP files containing images, audio, or text data.
Run SQL with AI operators using REST API or visual interface.
Ingest and analyze mixed data workloads in one environment.
Standard SQL data types including text, integers, floating-point values, and Booleans.
AI analysis over common image formats such as JPEG and PNG.
Document and free-text analysis stored directly in relational tables.
Analyze audio files such as MP3 and WAV with AI operators.
Express analytical intent in natural language while staying in SQL.
Filter multimodal data with natural-language conditions.
Classify records with label names only.
Assign scores based on natural-language criteria.
Transform or extract information with natural-language instructions.
Join datasets using semantic matching rules defined in natural language.
Each uploaded format is available as a first-class input for SQL + AI workflows.
Structured tabular uploads for relational analysis.
Archive container for mixed multimodal assets.
Compressed image format for visual analysis tasks.
Lossless image format with high visual fidelity.
Image format supported for uploaded media datasets.
Modern web image format for compact storage.
Compressed audio format for speech and sound data.
Uncompressed audio format for detailed signal quality.
AI Credits + daily compute time
Every plan unlocks the full SQL + AI operator set on PostgreSQL. Each plan bundles a monthly allotment of AI Credits for LLM usage and a per-day allowance of query compute time .
Best for occasional exploration
Best for regular analytics workflows
*Daily compute time resets at 00:00 UTC.
Switch between representative use cases and inspect each SQL pattern with syntax highlighting.
Finding complaints in survey comments
Count customers who mention pricing concerns in free-text responses.
SELECT COUNT(*)
FROM SurveyResults
WHERE AIFILTER(Comments, 'the customer complains about the price');
Classifying image collections
Tag uploaded holiday images with semantic classes in one query.
SELECT AICLASSIFY(pic, 'beach picture', 'city visit', 'other')
FROM HolidayImages;
Transcribing uploaded audio
Turn MP3 clips stored in your tables into plain-text transcripts with one query.
SELECT AIMAP(content, 'Transcribe this audio into plain text.') AS transcript
FROM AudioClips
WHERE filename LIKE '%.mp3';
Ranking sentiment in reviews
Sort comments by positivity score to prioritize outreach.
SELECT *
FROM SurveyResults
ORDER BY AISCORE(Comments, 1, 10, 'Positivity');
Summarizing customer-support tickets
Condense long tickets into short summaries for triage.
SELECT AIMAP(ticket_body, 'Summarize the customer issue in two sentences')
FROM SupportTickets;
Entity matching across product catalogs
Match equivalent products from independent datasets with semantic joins.
SELECT *
FROM ProductDB1 P1, ProductDB2 P2
WHERE AIJOIN(P1.productName, P2.productName, 'this is the same product');
Documentation and guided tutorials
Read endpoint-by-endpoint API docs, operator specs, and practical tutorials for both the visual interface and Python REST clients.
