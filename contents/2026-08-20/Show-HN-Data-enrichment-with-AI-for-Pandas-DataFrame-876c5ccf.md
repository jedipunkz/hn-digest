---
source: "https://github.com/mljar/enrichment"
hn_url: "https://news.ycombinator.com/item?id=49372451"
title: "Show HN: Data enrichment with AI for Pandas DataFrame"
article_title: "GitHub - mljar/enrichment: Data enrichment with AI for pandas DataFrame · GitHub"
image: "https://repository-images.githubusercontent.com/979881374/0788f01d-12a4-4807-a024-2a585ecd8dd4"
author: "pplonski86"
captured_at: "2026-08-20T10:20:24Z"
capture_tool: "hn-digest"
hn_id: 49372451
score: 1
comments: 0
posted_at: "2026-08-20T09:48:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Data enrichment with AI for Pandas DataFrame

- HN: [49372451](https://news.ycombinator.com/item?id=49372451)
- Source: [github.com](https://github.com/mljar/enrichment)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T09:48:41Z

## Translation

タイトル: Show HN: Pandas DataFrame の AI によるデータ エンリッチメント
記事タイトル: GitHub - mljar/enrichment: pandas DataFrame のための AI によるデータ エンリッチメント · GitHub
説明: pandas DataFrame の AI によるデータ強化。 GitHub でアカウントを作成して、mljar/エンリッチメントの開発に貢献します。

記事本文:
GitHub - mljar/enrichment: pandas DataFrame の AI によるデータ エンリッチメント · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
mljar
/
充実
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
20 コミット 20 コミット フォルダーとファイル
エンリッチメント エンリッチメント メディア メディア テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md pyproj

ect.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 列をパンダの DataFrame に追加します。
テーブルがあります。 「このレビューは満足していますか、それとも怒っていますか?」など、人間のみが記入できる新しい列が必要です。または「この会社は何の業界ですか?」
普通の英語で言いたいことを書きます。エンリッチメントでは、AI モデルにすべての行を要求し、新しいテーブルを返します。
あなたの指示: 「感情をポジティブ、ネガティブ、またはニュートラルに分類してください」
それが全体的な考え方です。 1 つの機能、1 つの命令文、1 つの新しい列。
パンダをPDとしてインポートする
エンリッチメントインポートからエンリッチメント
df = pd 。データフレーム (
{
「レビュー」: [
「商品が気に入りました！」 、
「壊れた状態で届きました。」 、
「大丈夫でした。」 、
]
}
)
結果 = 充実 (
DF、
input_col = "review" , # AI が読み取る列
Output_col = "sentiment" , # 新しい列の名前
プロンプト = "センチメントをポジティブ、ネガティブ、またはニュートラルに分類します" ,
)
印刷 (結果)
それだけです。入力する必要がある 4 つの項目: テーブル、読み取る列、作成する列、および必要なもの。
元の df は変更されません。常に新しい DataFrame が返されます。
同じテキストが 2 回出現した場合は、AI に 1 回送信されます。同じ行に対して 2 回支払うことはありません。
プロンプトは平易な英語なので、感情にとらわれることはありません。いくつかのアイデア:
# 乱雑なテキストから値を抽出する
エンリッチ ( df , input_col = "住所" , Output_col = "都市" ,
プロンプト = "都市名を抽出" )
# 物事をグループに分類する
エンリッチ ( df 、input_col = "チケット" 、output_col = "チーム" 、
プロンプト = "1 つのチームにルーティング: 請求、技術、または販売" )
# はい/いいえの質問
エンリッチ ( df , input_col = "email" , Output_col = "is_spam" ,
プロンプト = "はいまたはいいえで答えてください: この電子メールはスパムですか?" )
# 翻訳
エンリッチ ( df , input_col = "コメント" , Output_col = "英語" ,
プロンプト = "英語に翻訳" )
# 乱雑なデータをクリーンアップする

エンリッチ ( df , input_col = "job_title" , Output_col = "clean_title" ,
プロンプト = "標準の役職名に書き換えます。例: 「ソフトウェア エンジニア」" )
# 要約する
エンリッチ ( df , input_col = "記事" , Output_col = "概要" ,
プロンプト = "短い文で要約してください" )
複数の列を読む
場合によっては、1 つの列だけでは十分なコンテキストが得られないことがあります。 input_cols ( s に注意) を使用してリストを渡します。
結果 = 充実 (
DF、
input_cols = [ "会社名" , "ウェブサイト" ],
Output_col = "業界" ,
プロンプト = "会社の業界を決定する" ,
)
AI は両方の値を列名とともに確認するため、どちらがどちらであるかを認識します。
input_col または input_cols のいずれかを使用します。両方を同時に使用することはできません。
プロンプトは最も重要な部分です。小さな変化が大きな違いを生みます。
許可された回答をリストします。 「ポジティブ、ネガティブ、またはニュートラル」により、整然とした列が得られます。 「この人はどんな気持ちなんだろう？」段落を与えます。
答えはどれくらいの長さであるべきかを言ってください。 「一言で」「短い文で」。
最初に数行でテストします。 10,000 行を実行する前に df.head(10) を実行します。
# 最初は小さく試してみる
サンプル = エンリッチ ( df . head ( 10 )、input_col = "レビュー" 、output_col = "センチメント" 、
プロンプト = "センチメントをポジティブ、ネガティブ、またはニュートラルに分類します" )
プリント（サンプル）
大きなテーブル
大きなテーブルの場合は特別なことをする必要はありません。通常どおり、enrich() を呼び出すだけです。
小さなジョブは複数のリクエストとして同時に送信されるため、より速く完了します。
プロバイダーがサポートしている場合、大きなジョブ (50 以上の一意の値) は 1 つのバッチとして自動的に送信されます。 OpenAI では、バッチのコストが 50% 削減されます。最大 24 時間かかる場合もありますが、通常はそれよりもずっと短くなります。
プロバイダーが行を混ぜて返した場合でも、行は常に元の順序で返されます。
result = rich ( df , ..., use_batch = True ) # 常にバッチ処理
result = エンリッチ ( df , ..., use_batch = False ) # 決してしない

バッチ
result = rich ( df , ..., use_batch = None ) # デフォルト: 私が決める
OpenAI のバッチ制限: バッチごとに最大 50,000 の一意の行と 200 MB。
ネットワークの問題が発生します。エンリッチメントは一般的なものを処理します。レート制限 (HTTP 429)、タイムアウト、および一時的なサーバー エラーは自動的に再試行され、待機時間が長くなります。
空の行はスキップされ、pd.NA が取得されます。API 呼び出しやコストはかかりません。
失敗した行: デフォルトでは、行が失敗し続けるとジョブ全体が停止します。ジョブを終了して不良行をマークしたい場合は、 on_error="keep" を使用します。
結果 = 充実 (
DF、
input_col = "テキスト" ,
Output_col = "トピック" ,
プロンプト = "本題に戻ります" ,
on_error = "keep" 、失敗した行数はすべてを停止するのではなく pd.NA になります
)
リクエストの速度を遅くしたり速くしたりすることもできます。
結果 = 充実 (
DF、
input_col = "テキスト" ,
Output_col = "トピック" ,
プロンプト = "本題に戻ります" ,
max_concurrency = 10 、同時リクエスト数 (デフォルトは 5)
max_retries = 3 、諦めるまでの行ごとの試行回数
)
何が起こったか見てみましょう
return_report=True を追加すると、テーブルとともに小さなレポートが取得されます。コストやエラーの確認に役立ちます。
結果、レポート = 充実 (
DF、
input_col = "テキスト" ,
Output_col = "トピック" ,
プロンプト = "本題に戻ります" ,
return_report = True 、
)
print (レポート.完了) # 埋められた行数
print ( report . unique_requests ) # 実際に送信された呼び出しの数
print ( report . retries ) # 再試行が必要な回数
print (report .input_tokens , report .output_tokens ) # コストのための使用法
print (report .errors) # 何が失敗したか、もしあれば
現在、enrich() は 2 つのものを返すため、左側に 2 つの変数が必要であることに注意してください。
デフォルトの OpenAI モデルは gpt-5-nano です。高速かつ安価で、データの並べ替えと抽出に適しています。ほとんどの人がこれを使用します。

r.
難しいタスクでより高い品質が必要ですか?別のモデルを選択してください:
結果 = エンリッチ ( df , ..., モデル = "gpt-5.4" )
他の AI プロバイダーの使用
OpenAI に閉じ込められることはありません。自分のマシンで実行されているモデルも含め、OpenAI Chat Completions 形式を使用できるものはすべて機能します。
ローカル モデル (コンピューターからは何も残りません):
エンリッチメントからインポート OpenAICompatibilityProvider 、エンリッチメント
プロバイダー = OpenAI互換プロバイダー (
Base_url = "http://127.0.0.1:1234/v1" ,
モデル = "ローカルモデル" 、
)
結果 = 充実 (
DF、
input_col = "レビュー" ,
Output_col = "感情" ,
プロンプト = "感情を分類" ,
プロバイダー = プロバイダー、
)
キーを持つホスト型プロバイダー:
プロバイダー = OpenAI互換プロバイダー (
Base_url = "https://provider.example/v1" ,
api_key = "あなたの API キー" ,
model = "プロバイダー/モデル名" ,
headers = { "X-アプリ" : "私のアプリケーション" },
)
すべての設定
豊かにする（
DF、
input_col = なし、読み取る列の数
Output_col = None 、 # 作成する新しい列
プロンプト = なし 、 # 簡単な英語で必要なもの
model = None 、 # モデル名、空の場合はプロバイダーのデフォルト
api_key = None 、 # key、環境変数を使用しない場合
show_progress = True , # 進行状況バーを表示します
input_cols = None 、 # 複数の列を読み取る (input_col の代わりに)
プロバイダー = なし、 # カスタム プロバイダー オブジェクト
max_concurrency = 5 、並列リクエスト数
max_retries = 3 、行ごとの試行回数
retry_base_lay = 0.5 、最初の再試行の # 秒前
use_batch = None 、# True / False / None (自動)
on_error = "raise" 、 # 停止するには "raise"、 pd.NA を満たすには "keep"
return_report = False 、 # 実行レポートも返します
)
人々が尋ねる質問
AI について何か知る必要がありますか?
いいえ、文章を書くことができ、パンダを使用することができれば、準備は完了です。
私のデータはインターネットに送信されますか?
はい、どのプロバイダーを選択しても可能です。データがマシンから離れることができない場合は、次のコマンドを実行します。

ローカル モデルを作成し、それを OpenAICompatibilityProvider 経由で渡します。
いくらかかりますか？
それはこのパッケージではなく、プロバイダーとモデルによって異なります。 2 つの理由により、コストが低く抑えられます。繰り返される値は 1 回だけ送信されること、大規模なジョブでは安価なバッチが使用されることです。 return_report=True を実行して、正確なトークンの使用状況を確認します。
元のデータフレームは変更されますか?
いいえ、常に新しいものが戻ってきます。
空のセルはどうなりますか?
これらはスキップされ、 pd.NA で埋められます。これらに対して API 呼び出しは行われません。
毎回同じ答えを得ることができますか?
ほとんどの場合、AI モデルは若干異なる場合があります。重要な点については、出力のサンプルを自分で確認してください。
pip インストールのエンリッチメント
インストールに付属するパンダも必要です。
エンリッチメントには独自の AI はありません。 AI プロバイダーと対話します。最も簡単に始めるのは OpenAI です。
platform.openai.com にアクセスしてアカウントを作成します。
支払い方法を追加します (使用した分に対して支払います。通常、小さなテーブルの場合はセントです)。
API キーを作成してコピーします。 sk-... のようです。
エクスポート OPENAI_API_KEY= " あなたの API キー "
Windows (PowerShell):
$ env: OPENAI_API_KEY = " API キー "
または、必要に応じて、関数に直接渡します。
エンリッチ ( df , ..., api_key = "あなたの API キー" )
MLJAR スタジオをお使いですか?このセクション全体をスキップしても構いません。 Studio はサインインし、プロバイダーを選択します。 API キーは必要ありません。
独自のプロバイダーを作成する
プロバイダー インターフェイスは小さく、同期的です。
エンリッチメントインポートから CompletionResult 、プロバイダー
クラス MyProvider (プロバイダー):
名前 = "私のプロバイダー"
default_model = "私のモデル"
def complete ( self 、 request ):
値 = call_my_service (
指示 = 要求。説明書、
input_data = リクエスト 。入力データ 、
モデル＝リクエスト。モデルまたは自分自身。デフォルトモデル 、
)
CompletionResult を返す (コンテンツ = 値)
Bも実装して自動バッチサポートを追加します

マッチプロバイダー 。
エンリッチメントを組み込むアプリケーションはプロバイダーを登録できるため、ユーザーは何も設定する必要がありません。
エンリッチメントインポート register_provider から
register_provider ( "アプリケーション ランタイム" 、プロバイダー、優先度 = 100 )
プロバイダーは次の順序で選択されます。
api_key= を使用して明示的に構成された OpenAI
最も優先度の高い登録済みランタイム プロバイダー
MLJAR_RUNTIME_TOKEN_FILE からの MLJAR アカウント トークン
OPENAI_API_KEY を通じて構成された OpenAI
python -m pip install -e " .[dev] "
Python -m pytest
ライブ OpenAI テストは有料 API リクエストを行うため、デフォルトではスキップされます。
RUN_LIVE_API_TESTS=1 OPENAI_API_KEY= " API キー " python -m pytest -m live
ライブ Batch API テストは数分かかる場合があるため、個別にオプトインします。
RUN_LIVE_BATCH_TESTS=1 OPENAI_API_KEY= " API キー " \
python -m pytest -m live -k バッチ
ライセンス
MLJAR製。バグを見つけましたか、それともアイデアがありますか?問題を開きます。
pandas DataFrame の AI によるデータ強化
Readme Apache-2.0 ライセンス アクティビティ カスタム プロパティ スター
2 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Data enrichment with AI for pandas DataFrame. Contribute to mljar/enrichment development by creating an account on GitHub.

GitHub - mljar/enrichment: Data enrichment with AI for pandas DataFrame · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
mljar
/
enrichment
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
20 Commits 20 Commits Folders and files
enrichment enrichment media media tests tests .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Add AI columns to your pandas DataFrame.
You have a table. You want a new column that only a human could fill in — like "is this review happy or angry?" or "what industry is this company in?".
Write what you want in normal English. enrichment asks an AI model for every row and gives you back a new table.
Your instruction: "Classify sentiment as positive, negative, or neutral"
That is the whole idea. One function, one sentence of instructions, one new column.
import pandas as pd
from enrichment import enrich
df = pd . DataFrame (
{
"review" : [
"I loved the product!" ,
"It arrived broken." ,
"It was okay." ,
]
}
)
result = enrich (
df ,
input_col = "review" , # which column the AI reads
output_col = "sentiment" , # name of the new column
prompt = "Classify sentiment as positive, negative, or neutral" ,
)
print ( result )
That's it. Four things to fill in: your table, the column to read, the column to create, and what you want.
Your original df is not changed . You always get a new DataFrame back.
If the same text appears twice, it is sent to the AI once . You don't pay twice for the same row.
The prompt is just plain English, so you are not limited to sentiment. A few ideas:
# Pull a value out of messy text
enrich ( df , input_col = "address" , output_col = "city" ,
prompt = "Extract the city name" )
# Sort things into groups
enrich ( df , input_col = "ticket" , output_col = "team" ,
prompt = "Route to one team: billing, technical, or sales" )
# Yes / no questions
enrich ( df , input_col = "email" , output_col = "is_spam" ,
prompt = "Answer yes or no: is this email spam?" )
# Translate
enrich ( df , input_col = "comment" , output_col = "english" ,
prompt = "Translate to English" )
# Clean up untidy data
enrich ( df , input_col = "job_title" , output_col = "clean_title" ,
prompt = "Rewrite as a standard job title, e.g. 'Software Engineer'" )
# Summarize
enrich ( df , input_col = "article" , output_col = "summary" ,
prompt = "Summarize in one short sentence" )
Reading more than one column
Sometimes one column is not enough context. Use input_cols (note the s ) and pass a list:
result = enrich (
df ,
input_cols = [ "company_name" , "website" ],
output_col = "industry" ,
prompt = "Determine the company's industry" ,
)
The AI sees both values together, with their column names, so it knows which is which.
Use either input_col or input_cols — not both at the same time.
The prompt is the most important part. Small changes make a big difference.
List the allowed answers. "positive, negative, or neutral" gives you a tidy column. "How does this person feel?" gives you paragraphs.
Say how long the answer should be. "in one word", "in one short sentence".
Test on a few rows first. Run df.head(10) before running 10,000 rows.
# Try it small first
sample = enrich ( df . head ( 10 ), input_col = "review" , output_col = "sentiment" ,
prompt = "Classify sentiment as positive, negative, or neutral" )
print ( sample )
Big tables
You don't have to do anything special for large tables — just call enrich() as usual.
Small jobs are sent as several requests at the same time, so they finish faster.
Big jobs (50 or more unique values) are automatically sent as one batch, if your provider supports it. On OpenAI, batches cost 50% less . They can take up to 24 hours, but usually much less.
Rows always come back in the original order, even when the provider returns them mixed up.
result = enrich ( df , ..., use_batch = True ) # always batch
result = enrich ( df , ..., use_batch = False ) # never batch
result = enrich ( df , ..., use_batch = None ) # default: decide for me
OpenAI batch limits: up to 50,000 unique rows and 200 MB per batch.
Network problems happen. enrichment handles the common ones for you: rate limits (HTTP 429), timeouts, and temporary server errors are retried automatically with a growing wait time.
Empty rows are skipped and get pd.NA — no API call, no cost.
Failed rows: by default the whole job stops if a row keeps failing. If you'd rather finish the job and mark the bad rows, use on_error="keep" :
result = enrich (
df ,
input_col = "text" ,
output_col = "topic" ,
prompt = "Return the main topic" ,
on_error = "keep" , # failed rows become pd.NA instead of stopping everything
)
You can also slow down or speed up the requests:
result = enrich (
df ,
input_col = "text" ,
output_col = "topic" ,
prompt = "Return the main topic" ,
max_concurrency = 10 , # requests at the same time (default 5)
max_retries = 3 , # tries per row before giving up
)
See what happened
Add return_report=True to get a small report along with your table. Useful for checking cost and errors.
result , report = enrich (
df ,
input_col = "text" ,
output_col = "topic" ,
prompt = "Return the main topic" ,
return_report = True ,
)
print ( report . completed ) # how many rows were filled
print ( report . unique_requests ) # how many calls were actually sent
print ( report . retries ) # how many retries were needed
print ( report . input_tokens , report . output_tokens ) # usage, for cost
print ( report . errors ) # what failed, if anything
Note that enrich() now returns two things, so you need two variables on the left.
The default OpenAI model is gpt-5-nano . It is fast and cheap, and it's a good fit for sorting and extracting data — which is most of what people use this for.
Need better quality on a hard task? Pick another model:
result = enrich ( df , ..., model = "gpt-5.4" )
Using other AI providers
You are not locked into OpenAI. Anything that speaks the OpenAI Chat Completions format works — including models running on your own machine.
A local model (nothing leaves your computer):
from enrichment import OpenAICompatibleProvider , enrich
provider = OpenAICompatibleProvider (
base_url = "http://127.0.0.1:1234/v1" ,
model = "local-model" ,
)
result = enrich (
df ,
input_col = "review" ,
output_col = "sentiment" ,
prompt = "Classify sentiment" ,
provider = provider ,
)
A hosted provider with a key:
provider = OpenAICompatibleProvider (
base_url = "https://provider.example/v1" ,
api_key = "your-api-key" ,
model = "provider/model-name" ,
headers = { "X-App" : "My application" },
)
All the settings
enrich (
df ,
input_col = None , # column to read
output_col = None , # new column to create
prompt = None , # what you want, in plain English
model = None , # model name, provider default if empty
api_key = None , # key, if you don't use an env variable
show_progress = True , # show a progress bar
input_cols = None , # several columns to read (instead of input_col)
provider = None , # custom provider object
max_concurrency = 5 , # parallel requests
max_retries = 3 , # tries per row
retry_base_delay = 0.5 , # seconds before the first retry
use_batch = None , # True / False / None (automatic)
on_error = "raise" , # "raise" to stop, "keep" to fill pd.NA
return_report = False , # also return an execution report
)
Questions people ask
Do I need to know anything about AI?
No. If you can write a sentence and use pandas, you're ready.
Is my data sent to the internet?
Yes, to whichever provider you choose. If your data cannot leave your machine, run a local model and pass it through OpenAICompatibleProvider .
How much does it cost?
It depends on your provider and model, not on this package. Two things keep it low: repeated values are only sent once, and big jobs use cheaper batches. Run return_report=True to see your exact token usage.
Will it change my original DataFrame?
No. You always get a new one back.
What happens to empty cells?
They are skipped and filled with pd.NA . No API call is made for them.
Can I get the same answer every time?
Mostly, but AI models can vary a little. For anything important, check a sample of the output yourself.
pip install enrichment
You also need pandas, which comes along with the install.
enrichment does not have its own AI. It talks to an AI provider for you. The easiest one to start with is OpenAI.
Go to platform.openai.com and create an account.
Add a payment method (you pay for what you use — usually cents for small tables).
Create an API key and copy it. It looks like sk-... .
export OPENAI_API_KEY= " your-api-key "
Windows (PowerShell):
$ env: OPENAI_API_KEY = " your-api-key "
Or, if you prefer, pass it straight to the function:
enrich ( df , ..., api_key = "your-api-key" )
Using MLJAR Studio? You can skip this whole section. Studio signs you in and picks the provider for you. No API key needed.
Writing your own provider
The provider interface is small and synchronous:
from enrichment import CompletionResult , Provider
class MyProvider ( Provider ):
name = "my-provider"
default_model = "my-model"
def complete ( self , request ):
value = call_my_service (
instructions = request . instructions ,
input_data = request . input_data ,
model = request . model or self . default_model ,
)
return CompletionResult ( content = value )
Add automatic batch support by also implementing BatchProvider .
Applications that embed enrichment can register a provider so users never configure anything:
from enrichment import register_provider
register_provider ( "application-runtime" , provider , priority = 100 )
Providers are chosen in this order:
OpenAI configured explicitly with api_key=
Highest-priority registered runtime provider
MLJAR account token from MLJAR_RUNTIME_TOKEN_FILE
OpenAI configured through OPENAI_API_KEY
python -m pip install -e " .[dev] "
python -m pytest
Live OpenAI tests are skipped by default because they make paid API requests:
RUN_LIVE_API_TESTS=1 OPENAI_API_KEY= " your-api-key " python -m pytest -m live
The live Batch API test opts in separately, because it can take several minutes:
RUN_LIVE_BATCH_TESTS=1 OPENAI_API_KEY= " your-api-key " \
python -m pytest -m live -k batch
License
Made by MLJAR . Found a bug or have an idea? Open an issue .
Data enrichment with AI for pandas DataFrame
Readme Apache-2.0 license Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
