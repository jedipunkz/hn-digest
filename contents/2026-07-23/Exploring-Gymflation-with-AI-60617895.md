---
source: "https://cbowdon.github.io/posts/gymflation/"
hn_url: "https://news.ycombinator.com/item?id=49028017"
title: "Exploring Gymflation with AI"
article_title: "Exploring Gymflation with AI – Chris Bowdon"
author: "sebg"
captured_at: "2026-07-23T21:58:14Z"
capture_tool: "hn-digest"
hn_id: 49028017
score: 1
comments: 0
posted_at: "2026-07-23T21:02:31Z"
tags:
  - hacker-news
  - translated
---

# Exploring Gymflation with AI

- HN: [49028017](https://news.ycombinator.com/item?id=49028017)
- Source: [cbowdon.github.io](https://cbowdon.github.io/posts/gymflation/)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T21:02:31Z

## Translation

タイトル: AI でジムフレーションを探る
記事のタイトル: AI によるジムフレーションの探索 – Chris Bowdon
説明: それは本物ですか? YouTube データと LLM (Ellmer、R) を使用して調べます。

記事本文:
その他
42 London 向け AI MVP ワークショップ (YouTube)
ブリストル データ インテリジェンス ネットワークの Gymflation
それは本当ですか？ YouTube データと LLM (Ellmer、R) を使用して調べます。
スーパーヒーローのインフレについてご存知ですか？時間が経つにつれて、スクリーン上のスーパーヒーローの体格はますます誇張されるようになりました。アダム・ウェストからベン・アフレックへのバットマンの成長は、その良い例です。
ジムフレーションもほぼ同じ考えです。ジム文化が急増するにつれて、人々の強さの妙技もソーシャルメディアに記録されているようです。 YouTube や Instagram を見ていると、デッドリフト 200kg は今ではむしろ平均的なものであるように感じられます。
しかし、それは本当に本当なのでしょうか？それとも、私たちはただアルゴリズムの影響を感じていて、青く照らされた顔に極端な例を押し付けているだけなのでしょうか？
以下は、それを調べるための小さなプロジェクトです。私たちは次のことを行います:
httr2 を使用して YouTube API をスクレイピングして 10 年分のデッドリフト個人記録 (PR) 動画を取得する
構造化された応答を使用して、AI で持ち上げた重量を抽出します
抽出の品質を評価し、ICL で改善します
マルチモーダル モデルを使用してリフターの性別を判断する
エラーによって後戻りしないようにキャッシュを最大限に活用する
古き良き ggplot2 を使用して結果をプロットします。
大きな満足感を持って帰ってください
私の隠れた動機は、これが R プロジェクトで API と LLM を使用するためのガイドとして機能し、私にとって役立ついくつかのテクニックと実践方法を示すことです。いつものように、完全なコードは折り目の下にあります。
ライブラリ (整理整頓)
ライブラリ (ニット)
ライブラリ (httr2)
図書館（エルマー）
#ライブラリ(メモワーズ)
#ライブラリ(キャッシュ)
ライブラリ（キャレット）
httr2 とディスク キャッシュを使用して YouTube データをスクレイピングする
YouTube API は httr2 で使用するのが比較的簡単であり、それ自体が物語っています。行う必要があるリクエストは GET リクエストであるため、便利な req_cache 関数を使用できます。

繰り返しのリクエストでクォータが消費されないようにするための機能です。
YOUTUBE_API_KEY <- Sys.getenv ( "YOUTUBE_API_KEY" )
search_videos <- 関数 (
クエリ、
published_after = "2026-01-01T00:00:00Z" ,
公開前 = "2026-01-02T00:00:00Z"
) {
req <- request ( "https://youtube.googleapis.com/youtube/v3/search" ) |>
req_cache ( ".cache/httr2" ) |> # 幸いにもこのエンドポイントは GET なので、キャッシュできます
req_url_query (
部分 = "スニペット" 、
q = クエリ、
type = "ビデオ" 、
maxResults = 50 、
期間 = "短い" 、
地域 = "英国" 、
公開後 = 公開後、
公開前 = 公開前、
キー = YOUTUBE_API_KEY
) |>
req_headers ( Accept = "application/json" ) |>
req_timeout ( 30 )
resp <- req_perform (要求)
if ( resp_is_error (resp)) {
stop ( "YouTube API リクエストが失敗しました: HTTP " 、 resp_status (resp))
}
resp_body_json (応答、simplifyVector = TRUE )
}
クエリ <- " \" デッドリフト PR \" -anatoly"
-anatoly という用語は、非常に人気があるが無関係な Anatoly ビデオを削除することを意味します。 （彼はエリートのパワーリフターで、清掃員の格好をしてボディービルダーたちにいたずらをしています。どうやらインターネットではこれが十分に理解できないようです。）
YouTube の無料 API は非常に制限が厳しいため、妥協せざるを得ません。毎月の初日のみ、および結果の最初のページ (最大 50) のみをサンプリングします。私の愚かな研究に資金を提供してくれる人がいれば、喜んでこれらの問題を解決します。
rds.file <- "data/youtube_df.rds"
df <- {
if ( file.exists (rds.file)) {
readRDS (rds.file)
} それ以外の場合は {
日付 <- seq ( as.Date ( "2015-01-01" )、 as.Date ( "2025-12-01" )、 by = "month" )
後 <- 日付[ - 長さ (日付)]
以前 <- 日付[ - 1 ]
query_results <- purrr ::map2 (
その後、
以前、
.f = 関数 (a, b) {
検索ビデオ (
クエリ、
published_after = sprintf ( "%sT00:00:00Z" , a),
published_before = sprintf ( "%sT00:00:00Z" , b)
)
}、
.progress = TRUE
)
DF <

- クエリ結果 |>
purrr ::map (\(r) unnest (r $ items,snippet)) |>
バインド行 () |>
mutate (publish.date = as.Date (publishedAt)) |>
unnest ( c (id, サムネイル), names_sep = "." ) |>
unnest (starts_with ( "thumb" ), names_sep = "." ) |>
ユニークな ()
saveRDS (df、ファイル = rds.file、圧縮 = FALSE )
DF
}
}
これを書いているときに、どういうわけかキャッシュを無効にすることができました。おっと。しかし、これは私にとって初めてのロデオではなく、フォールバックとして結果の RDS ファイルを持っていました。これまでのところ失敗していないことを確認するために、毎日の結果数をチェックしてみましょう。
ggplot (df) +
aes ( x = 公開.日付) +
geom_histogram () +
テーマミニマル () +
labs ( x = "公開日" 、 y = "ビデオの数" )
図 1: 日付別のビデオの分布。均一性は、通常、フェッチしていない他のページがあることを示していますが、これらが関連する結果であるかどうかは保証されません。
抽出精度の評価
精度をチェックする必要があります。そうでない場合は、タスクを AI に委任せず、願いと祈りに委任しています。私は、手動でラベルを付けるために小さなサンプルを CSV にダンプしています。自分のデータを知るために、最初のインスタンスでは常にこれをお勧めします。
より大きなサンプルが必要な場合は、使用している LLM よりも大きな LLM に委任することができます (このタスクに関して人間の精度と同等かどうかはわかりません…評価者を評価する必要があるかもしれません…) もう 1 つのアプローチは、人間が可能な限り効率的にラベル付けできるように、独自のラベル付けアプリをコーディングすることです。 Prodigy のようなラベル付けアプリを購入することもできますが、2026 年には、データに合わせて何かをバイブコーディングするためにトークンを費やす方が安くなります。このアプローチを取る場合は、人間の偏見にも注意してください。
セットシード ( 42 )
df |>
スライスサンプル ( n = 100 ) |>
select (id.videoId、タイトル、説明) |>
write_csv ( "data/unannotated-weight.csv" )
ラベル

このデータを調べてみると、PR を複数回繰り返して投稿し、体重を含めることがあることがわかりました。 180x2 @ 83kg、つまり 180kg を 2 回持ち上げ、体重は 83kg になります。重量が軽くなるため、複数回の繰り返しを考慮する必要があります。適切なアプローチは、エプリーの 1 Rep Max (1RM) 式を使用して、1 回の繰り返しにおけるリフターの最大値を推定することです。
ラベル付き <- read_csv ( "data/annotated-weight.csv" ) |>
# 行がレビューされたかどうかを示すフラグを CSV に入れると、すべてにラベルを付けることはできないことが多いため便利です
フィルター (レビュー済み) |>
選択 ( - レビュー済み)
head (ラベル付き) |> kable ()
id.videoId
タイトル
説明
重量
ユニット
担当者
特別な
Ms64fHyxQoI
デッドリフト PR 2020/11/29 - 495 @ 286 ポンド
NA
495
ポンド
1
NA
nKAUQ5c23vA
ケイラ C 190 デッドリフト PR 1/7/19
NA
NA
NA
NA
NA
dBVlfi-UNNY
190 kg デッドリフト pr #ジム #パワーリフティング #デッドリフト
NA
190
kg
1
NA
Sfxc3yTeKz8
デッドリフト PR
5 レップの PR が以前の 1 レップの最大値に達したら、その日は良い日であることがわかります。こちらは355ポンド×3×5です。スマートなプログラミングはすべてを意味します…
355
ポンド
5
NA
1yt6w5HbQuA
675LB デッドリフト担当 PR |心配する必要はありますか? |パワーリフティング準備 Ep. 10
ビデオをお楽しみください?いいね＆購読してください！私のウェブサイト: http://www.russwole.com Alphalete アパレルを入手: http://alphaleteathletics.com 入手 …
675
ポンド
1
NA
NvzE_4wITTw
405# デッドリフト PR
5キロ走った後、デッドリフトで自己PRを達成しました。
405
NA
1
NA
したがって、単位変換と 1RM 変換という 2 つの変換を適用します。
Convert_kg <- 関数 (重量、単位) {
case_when (
単位 == "kg" ~ 重量、
単位 == "ポンド" ~ 重量 * 0.45359237 、
.default = - 1
)
}
epley_1rm <- 関数 (weight_kg, reps) {
if_else (
担当者 == 1 、
体重_kg、
# 0.5 に四捨五入
ラウンド (体重 kg * ( 1 + 担当者 / 30 ) * 2 ) / 2
)
}
ラベル付き <- ラベル付き |>
突然変異(
reps = 合体 (reps, 1 ), # デフォルト

不特定のすべての繰り返しは 1 にカウントされます
Weight.kg = Convert_kg (重量、単位)、
Weight.kg.1rm = epley_1rm (weight.kg、回数)
)
head (ラベル付き) |> kable ()
表 1: ラベル付きデータセットの例
id.videoId
タイトル
説明
重量
ユニット
担当者
特別な
重量.kg
重量.kg.1rm
Ms64fHyxQoI
デッドリフト PR 2020/11/29 - 495 @ 286 ポンド
NA
495
ポンド
1
NA
224.5282
224.5282
nKAUQ5c23vA
ケイラ C 190 デッドリフト PR 1/7/19
NA
NA
NA
1
NA
-1.0000
-1.0000
dBVlfi-UNNY
190 kg デッドリフト pr #ジム #パワーリフティング #デッドリフト
NA
190
kg
1
NA
190.0000
190.0000
Sfxc3yTeKz8
デッドリフト PR
5 レップの PR が以前の 1 レップの最大値に達したら、その日は良い日であることがわかります。こちらは355ポンド×3×5です。スマートなプログラミングはすべてを意味します…
355
ポンド
5
NA
161.0253
188.0000
1yt6w5HbQuA
675LB デッドリフト担当 PR |心配する必要はありますか? |パワーリフティング準備 Ep. 10
ビデオをお楽しみください?いいね＆購読してください！私のウェブサイト: http://www.russwole.com Alphalete アパレルを入手: http://alphaleteathletics.com 入手 …
675
ポンド
1
NA
306.1748
306.1748
NvzE_4wITTw
405# デッドリフト PR
5キロ走った後、デッドリフトで自己PRを達成しました。
405
NA
1
NA
-1.0000
-1.0000
NA は実際に役に立ちます。LLM が NA を抽出不可能なものとして正しく識別したことを確認したいからです。
ラベリングの際に、ケースの数パーセントがトラップまたはヘックスバーのデッドリフトであることにも気づきました。これは特別なタイプのバーで、重いですが持ち上げやすいです。単純な古い正規表現を使用してそれらを除外します。
ラベル付き <- ラベル付き |> フィルタ (is.na (特殊))
df <- df |>
filter ( ! str_detect (タイトル、 "trap|hex" ) | str_detect (説明、 "trap|hex" ))
最後に、データをテスト セットと開発セットに分割します。評価用のテストセットを差し上げます
テスト <- ラベル付き |> スライスヘッド ( n = 50 )
dev <- ラベル付き |> スライステール ( n = nrow (ラベル付き) - 50 )
tibble ( nrow (テスト), nrow (開発)) |> kable ()
狭い(テスト)
行(開発)
50
48
あ

この時点で、抽出「モデル」を構築する準備が整いました。
なぜ引用符なのでしょうか？それは、ほとんどの人が慣れ親しんでいる意味でのモデリングとはまったく異なります。これから行うことは、非常に強力な基本モデルがタスクで適切に実行されるように調整するための高度なプロンプトです。その一部はそれ自体を物語っています。システム プロンプトは、一貫した入力を期待し、特定の方法で応答するようにモデルを設定します。
この場合のモデルは、Mistral AI のオープンソース ビジョン言語モデル Ministral 3B のインスタンスであり、LM Studio を介してローカルで実行されます。ラップトップに十分な機能があれば、オープンソース モデルをローカルで実行するのは非常に簡単です (これには少なくとも 8GB の VRAM が必要だと思います)。
一貫性があり解析しやすい出力が必要なため、特定のキーを含む JSON 応答を期待していることを指定していることに注意してください。
LLM_SERVER_URL <- "http://localhost:8080/v1"
extract_dl_weight_model <- 関数 ( Turns = list ()) {
system_prompt <- "ユーザーは YouTube ビデオのタイトルと説明を入力します。デッドリフトの重量について説明されている場合は、デッドリフトの重量、単位、および反復回数を抽出する必要があります。
現在のデッドリフトの世界記録は 510kg/1122lb であることに注意してください。
キーを持つ JSON オブジェクトのみを返します。
- `重量` (数値)
- 「単位」 (kg、lb、または NA)
- `reps` (数値、指定されていない場合は 1 と仮定します)。
単位変換は行わないでください。
説明にデッドリフト重量が記載されていない場合、または単位が明確でない場合は、単位として NA を返してください。」
チャット <- chat_openai_compatibility (
LLM_SERVER_URL、
モデル = "ミストラライ/ミニストラル-3-3b" ,
credentials = function () "" , # ローカル、認証情報を無視する
params = params ( max_tokens = 50 ), # 必要以上に生成しない
システムプロンプト = システムプロンプト
)
チャット $ set_turns (ターン)
チャット
}
セットターンのビジネスとは何でしょうか?だから私たちはc

会話履歴のプライム - 詳細については後ほど説明します。
構造化された応答の生成
ほとんどの場合、単純に JSON 応答を要求するだけで十分ですが、構造化された応答を使用して強制することもできます。 Ellmer を使用すると、受信したいものを記述する JSON スキーマを構築できます。また、LLM へのさらなるガイダンスとして各フィールドの説明を追加する機会でもあります。
extract_dl_weight_type <- type_object (
Weight = type_number ( "デッドリフトの重量を数値で示します。" ),
単位 = type_enum (値 = c ( "kg" 、 "lb" 、 "NA" ))、
reps = type_integer ( "繰り返しの数 (指定されていない場合は 1)" )
)
構造化された応答は非常に単純です。各ステップで、LLM は考えられるすべての次のトークン (数十万個) にわたる確率分布を生成します。通常、上位 k 個の最も可能性の高いトークンから 1 つをサンプリングし、それを入力の最後に追加して繰り返します。ただし、構造化された応答では、無​​効なトークンではなく、指定された文法に準拠するトークンのみが配布に含まれます。どのような文法でも使用できますが、JSON スキーマが最も一般的な実装です。
これは、ここで使用しているかなり弱いモデルであっても、うまく機能します。
example.response <- extract_dl_weight_model () $ chat_structed (
タイトル: 体重 188 ポンドで 665 ポンドのデッドリフト #deadlift

[切り捨てられた]

## Original Extract

Is it real? Using YouTube data, and LLMs (Ellmer, R) to find out.

Misc
AI MVP workshop for 42 London (YouTube)
Gymflation for Bristol Data Intelligence Network
Is it real? Using YouTube data, and LLMs (Ellmer, R) to find out.
You might be familiar with super hero inflation? Over time, super hero physiques on screen have become increasingly exaggerated. Batman’s progression from Adam West to Ben Affleck is a great example of this.
Gymflation is much the same idea. As gym culture has skyrocketed, it seems like so too have people’s feats of strength - as recorded on social media. Going on YouTube or Instagram one gets the feeling that a 200kg deadlift is really rather average nowadays.
But is it really true? Or are we just feeling the effects of the Algorithm, pushing extreme examples in our blue-lit faces?
What follows is a little project to find out. We will:
Scrape the YouTube API for 10 years of deadlift personal record (PR) videos with httr2
Use structured responses to extract the lifted weight with AI
Evaluate the quality of our extraction and improve it with ICL
Use a multimodal model to judge the lifter’s gender
Make extensive use of caching so errors don’t set us back
Plot our results with good old ggplot2
Come away with a sense of tremendous satisfaction
My secret ulterior motive is that this can serve as a guide to using APIs and LLMs in an R project, demonstrating some techniques and practices that have been helpful to me. As usual the full code is below the folds.
library (tidyverse)
library (knitr)
library (httr2)
library (ellmer)
#library(memoise)
#library(cachem)
library (caret)
Scraping YouTube data with httr2 and a disk cache
The YouTube API is relatively straightforward to use with httr2 and speaks for itself. Because the requests we need to make are GET requests, we can use the handy req_cache function to avoid burning quota on repeat requests.
YOUTUBE_API_KEY <- Sys.getenv ( "YOUTUBE_API_KEY" )
search_videos <- function (
query,
published_after = "2026-01-01T00:00:00Z" ,
published_before = "2026-01-02T00:00:00Z"
) {
req <- request ( "https://youtube.googleapis.com/youtube/v3/search" ) |>
req_cache ( ".cache/httr2" ) |> # happily this endpoint is a GET so we can cache
req_url_query (
part = "snippet" ,
q = query,
type = "video" ,
maxResults = 50 ,
duration = "short" ,
region = "UK" ,
published_after = published_after,
published_before = published_before,
key = YOUTUBE_API_KEY
) |>
req_headers ( Accept = "application/json" ) |>
req_timeout ( 30 )
resp <- req_perform (req)
if ( resp_is_error (resp)) {
stop ( "YouTube API request failed: HTTP " , resp_status (resp))
}
resp_body_json (resp, simplifyVector = TRUE )
}
query <- " \" deadlift PR \" -anatoly"
The -anatoly term is to get rid of the wildly popular but irrelevant Anatoly videos. (He’s an elite powerlifter who dresses as a cleaner to prank bodybuilders. Apparently the Internet can’t get enough of this.)
The YouTube free API is quite restrictive so we are forced to compromise. We sample only the first day of each month, and only the first page of results (up to 50). If anyone is willing to fund my stupid research I will gladly resolve these issues.
rds.file <- "data/youtube_df.rds"
df <- {
if ( file.exists (rds.file)) {
readRDS (rds.file)
} else {
dates <- seq ( as.Date ( "2015-01-01" ), as.Date ( "2025-12-01" ), by = "month" )
afters <- dates[ - length (dates)]
befores <- dates[ - 1 ]
query_results <- purrr :: map2 (
afters,
befores,
.f = function (a, b) {
search_videos (
query,
published_after = sprintf ( "%sT00:00:00Z" , a),
published_before = sprintf ( "%sT00:00:00Z" , b)
)
},
.progress = TRUE
)
df <- query_results |>
purrr :: map (\(r) unnest (r $ items, snippet)) |>
bind_rows () |>
mutate ( publish.date = as.Date (publishedAt)) |>
unnest ( c (id, thumbnails), names_sep = "." ) |>
unnest ( starts_with ( "thumb" ), names_sep = "." ) |>
unique ()
saveRDS (df, file = rds.file, compress = FALSE )
df
}
}
Somewhow when writing this up I managed to nuke the cache. Oof. But it’s not my first rodeo, I had an RDS file of the results as a fallback. Let’s check out the daily result counts to ensure we haven’t messed up so far.
ggplot (df) +
aes ( x = publish.date) +
geom_histogram () +
theme_minimal () +
labs ( x = "Publish date" , y = "Count of videos" )
Figure 1: Distribution of videos by date. The uniformity indicates that there are usually other pages that we haven’t fetched, though it’s not guaranteed that these are relevant results.
Evaluating extraction accuracy
We will need to check accuracy, otherwise we’ve not delegated a task to AI we’ve delegated it to a wish and a prayer. I’m dumping a small sample into a CSV to be labelled manually, which I’d always recommend in the first instance so that you get to know your own data.
If you need a larger sample then it could be delegated to a larger LLM than the one you are using (which may or may not equal human accuracy on this task… you may need to evaluate your evaluator…) Another approach is to (vibe) code your own labelling app so that a human can label as efficiently as possible. You can buy a labelling app like Prodigy, but in 2026 it’s cheaper to spend the tokens vibe-coding something bespoke to your data. If you take this approach beware human bias too.
set.seed ( 42 )
df |>
slice_sample ( n = 100 ) |>
select (id.videoId, title, description) |>
write_csv ( "data/unannotated-weight.csv" )
Labelling this data I discovered that sometimes people post a multi-repetition PR and include their body weight, e.g. 180x2 @ 83kg, meaning they lifted 180kg twice and their bodyweight is 83kg. We will need to account for multiple reps, because the weight will be lower. A decent approach is to use Epley’s 1 Rep Max (1RM) formula to estimate the lifter’s maximum for a single repetition.
labelled <- read_csv ( "data/annotated-weight.csv" ) |>
# putting a flag in your CSV for whether the row has been reviewed or not is helpful because you often won't label it all
filter (reviewed) |>
select ( - reviewed)
head (labelled) |> kable ()
id.videoId
title
description
weight
unit
reps
special
Ms64fHyxQoI
Deadlift PR 11/29/2020 - 495 @ 286 lb
NA
495
lb
1
NA
nKAUQ5c23vA
Kayla C 190 Deadlift PR 1/7/19
NA
NA
NA
NA
NA
dBVlfi-UNnY
190 kg deadlift pr #gym #powerlifting #deadlift
NA
190
kg
1
NA
Sfxc3yTeKz8
DEADLIFT PR
You know it’s a good day when you hit your former 1 rep max for a 5 rep PR. Here is 355lbs.x3x5. Smart programming means all …
355
lb
5
NA
1yt6w5HbQuA
675LB Rep PR On Deadlift | Should I Be Worried? | Powerlifting Prep Ep. 10
Enjoy the video? Like & Subscribe! My Website: http://www.russwole.com Get Alphalete Apparel: http://alphaleteathletics.com Get …
675
lb
1
NA
NvzE_4wITTw
405# Deadlift PR
After a 5k run I got my PR on the deadlift.
405
NA
1
NA
We’ll therefore apply two conversions: unit conversion and 1RM conversion.
convert_kg <- function (weight, unit) {
case_when (
unit == "kg" ~ weight,
unit == "lb" ~ weight * 0.45359237 ,
.default = - 1
)
}
epley_1rm <- function (weight_kg, reps) {
if_else (
reps == 1 ,
weight_kg,
# rounded to nearest 0.5
round (weight_kg * ( 1 + reps / 30 ) * 2 ) / 2
)
}
labelled <- labelled |>
mutate (
reps = coalesce (reps, 1 ), # default all unspecified rep counts to 1
weight.kg = convert_kg (weight, unit),
weight.kg.1rm = epley_1rm (weight.kg, reps)
)
head (labelled) |> kable ()
Table 1: Example of labelled dataset
id.videoId
title
description
weight
unit
reps
special
weight.kg
weight.kg.1rm
Ms64fHyxQoI
Deadlift PR 11/29/2020 - 495 @ 286 lb
NA
495
lb
1
NA
224.5282
224.5282
nKAUQ5c23vA
Kayla C 190 Deadlift PR 1/7/19
NA
NA
NA
1
NA
-1.0000
-1.0000
dBVlfi-UNnY
190 kg deadlift pr #gym #powerlifting #deadlift
NA
190
kg
1
NA
190.0000
190.0000
Sfxc3yTeKz8
DEADLIFT PR
You know it’s a good day when you hit your former 1 rep max for a 5 rep PR. Here is 355lbs.x3x5. Smart programming means all …
355
lb
5
NA
161.0253
188.0000
1yt6w5HbQuA
675LB Rep PR On Deadlift | Should I Be Worried? | Powerlifting Prep Ep. 10
Enjoy the video? Like & Subscribe! My Website: http://www.russwole.com Get Alphalete Apparel: http://alphaleteathletics.com Get …
675
lb
1
NA
306.1748
306.1748
NvzE_4wITTw
405# Deadlift PR
After a 5k run I got my PR on the deadlift.
405
NA
1
NA
-1.0000
-1.0000
The NAs are actually useful, because we want to confirm that the LLM has correctly identified these as non-extractable.
In labelling I also noticed that a small percent of cases are trap or hex bar deadlifts. This is a special type of bar that’s heavier but easier to lift. We’ll filter those out with a plain old regular expression.
labelled <- labelled |> filter ( is.na (special))
df <- df |>
filter ( ! str_detect (title, "trap|hex" ) | str_detect (description, "trap|hex" ))
Finally we’re split the data into a test and dev set. We hold out the test set for evaluation
test <- labelled |> slice_head ( n = 50 )
dev <- labelled |> slice_tail ( n = nrow (labelled) - 50 )
tibble ( nrow (test), nrow (dev)) |> kable ()
nrow(test)
nrow(dev)
50
48
At this point we’re ready to build an extraction “model”.
Why the quotes? Well it’s very unlike modelling in the sense most people are used to. What we’re going to do is advanced prompting to condition a very powerful base model into performing well on our task. Some of it speaks for itself: the system prompt sets the model up to expect a consistent input and to respond in a certain way.
The model in this case is an instance of the open source vision-language model Ministral 3B from Mistral AI, running locally via LM Studio. Running open source models locally is very easy provided your laptop has the grunt (I think you’d need at least 8GB of VRAM for this one).
Note that we specify that we’re expecting a JSON response with certain keys, because we want consistent and easy-to-parse output.
LLM_SERVER_URL <- "http://localhost:8080/v1"
extract_dl_weight_model <- function ( turns = list ()) {
system_prompt <- "The user will provide the title and description of a YouTube video. If they describe a weight being deadlifted, you must extract the weight that was deadlifted, with the unit, and the number of repetitions.
Note that the deadlift world record is currently 510kg/1122lb.
Return only a JSON object with keys:
- `weight` (numeric)
- `unit` (kg, lb or NA)
- `reps` (numeric, assume 1 if not provided).
Do not perform unit conversion.
Return NA as the unit if the description doesn't mention a weight deadlifted or the unit isn't clear."
chat <- chat_openai_compatible (
LLM_SERVER_URL,
model = "mistralai/ministral-3-3b" ,
credentials = function () "" , # local, just ignore creds
params = params ( max_tokens = 50 ), # don't generate more than we need
system_prompt = system_prompt
)
chat $ set_turns (turns)
chat
}
What’s all that set_turns business? That’s so we can prime the conversation history - more on that later.
Structured response generation
Simply asking for JSON response is good enough most of the time, but we can do better and use structured responses to enforce it. With Ellmer we can build a JSON schema describing what we want to receive. It’s also an opportunity to add a description of each field for further guidance to the LLM.
extract_dl_weight_type <- type_object (
weight = type_number ( "The weight deadlifted, as a number." ),
unit = type_enum ( values = c ( "kg" , "lb" , "NA" )),
reps = type_integer ( "Number of repetitions (1 if unspecified)" )
)
Structured responses are quite simple: at each step, the LLM generates a probability distribution over all possible next tokens (hundreds of thousands of them). Usually we sample one from the top k most likely tokens, tack it onto the end of our input and repeat. However with structured responses we only include tokens that adhere to a specified grammar in the distribution, not those that would be invalid. Any grammar is possible, but JSON schema is the most popular implementation.
This works nicely, even with the fairly weak model we’re using here.
example.response <- extract_dl_weight_model () $ chat_structured (
"Title: 665 lbs deadlift at 188 lbs bodyweight #deadlift

[truncated]
