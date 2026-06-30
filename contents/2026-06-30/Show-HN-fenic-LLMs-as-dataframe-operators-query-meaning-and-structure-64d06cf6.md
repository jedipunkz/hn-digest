---
source: "https://github.com/typedef-ai/fenic"
hn_url: "https://news.ycombinator.com/item?id=48735301"
title: "Show HN: fenic – LLMs as dataframe operators, query meaning and structure"
article_title: "GitHub - typedef-ai/fenic: Semantic DataFrames for humans and agents · GitHub"
author: "cpard"
captured_at: "2026-06-30T16:42:18Z"
capture_tool: "hn-digest"
hn_id: 48735301
score: 2
comments: 0
posted_at: "2026-06-30T16:39:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: fenic – LLMs as dataframe operators, query meaning and structure

- HN: [48735301](https://news.ycombinator.com/item?id=48735301)
- Source: [github.com](https://github.com/typedef-ai/fenic)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T16:39:40Z

## Translation

タイトル: 表示 HN: fenic – データフレーム演算子としての LLM、クエリの意味と構造
記事のタイトル: GitHub - typedef-ai/fenic: 人間とエージェントのためのセマンティック データフレーム · GitHub
説明: 人間とエージェントのためのセマンティック データフレーム。 GitHub でアカウントを作成して、typedef-ai/fenic の開発に貢献してください。
HN テキスト: やあ、友達。私にとって大切なプロジェクトを共有したいと思います。 fenic は、LLM が第一級市民として追加されたデータフレーム API であり、LLM をサポートする新しい演算子で拡張された古典的な遅延データフレーム API です。これにより、同じコンテキストで構造化データと非構造化データを操作できるようになります。最も重要なのは、LLM が不透明な UDF ブラック ボックスとして統合されていないことです。これらは、プランナーが古典的な演算子と並んで推論できる「セマンティック」演算子として公開されます。 (すべてがどのように連携するかを確認するために、リポジトリに例とコード スニペットがあります) これを構築する理由は何ですか?私はデータインフラ/システム担当者です。 LLM が登場したとき、私が目にしたのは、システムの特性を変える新しいタイプのコンピューティングでした。
私たちが扱うワークロード。私は、現在のシステムがこれらの新しいワークロードとコンピューティング タイプをどのように吸収できるか、DX を可能な限りシームレスにするために何が必要かを実験したかったのですが、UDF + 任意のプロンプトがあまりにも問題だと感じていたのはそこです。これを適切にサポートするには、新しいプラン演算子という非常に優れた機能をいくつか導入する必要がありました。 LLM ではプロンプトを送信するだけではありません。セマンティック結合、セマンティック マップとリデュース、セマンティック フィルターなどの演算子を使用します。これらは古典的な演算子と混合されており、プランナーはそれらをブラック ボックスではなく実際の演算子と見なすため、それらを回避する作業の順序を変更できます。型付き出力。人間工学に基づいて、セマンティック演算子の出力を型指定されたデータフレーム列に直接変換します。ピダンティック

LLM 出力のスキーマは、ネストを解除したり、展開したりできる型付き構造体列になります。マークダウン データ型のような新しいデータ型。マークダウンは、プレゼンテーション用にテキストをフォーマットする方法として誕生したにもかかわらず、LLM と情報を共有するための重要な方法になりました。これは構造を保持しており、構造体または JSON 型と同じ方法でその構造にアクセスできるため、前述した開発者のエクスペリエンスが向上します。非同期 UDF。 LLM の爆発的な増加によるワークロードの最も興味深い変化の 1 つは、API からの応答の取得、Web サイトのクロールなど、I/O バウンドの高いステップをパイプラインに組み込む必要があることです。非同期 UDF はそのギャップを埋め、同時実行性、再試行などの微妙な違いは実装によって処理されます。 LLM 推論対応のプランナーおよびランタイム。これは私が最も楽しみにしている部分の 1 つであり、やるべきことはまだたくさんあります。現在: バッチ内の同一のプロンプトは単一のモデル呼び出しに折りたたまれるため、複製のコストはゼロトークンになります。リクエストは、プロバイダーごとの rpm/tpm 制限の下で、再試行とバックオフを使用して同時にディスパッチされます。 null セルと空のセルはモデルを完全にスキップします。そして、オペレーターごとのトークンとコストのメトリクスを取得します。オプションの永続的な応答キャッシュもあるため、再実行ではモデルがスキップされます。新しいカタログ プリミティブとしての MCP。登録されたビューと同様に、データフレーム パイプラインを MCP ツールとしてカタログに登録できます。 fenic は、データ上で実行されるツールのロジックとしてそのパイプラインを備えた MCP サーバーを提供します。これらは、LLM をコンピューティング インフラストラクチャの一部にする方法を実験する際に fenic に導入されたものの一部にすぎません。すでにあるものをさらに磨くべきことはまだたくさんあります。私はあらゆる種類のことにfenicを使用してきました。小規模/個人的な目的では、ポッドキャストの音声録音を取り込み、素敵な音声ファイルに変換するために使用しています。

調査できるメタデータの構造化されたテーブル。より重い側面では、エージェントが Pydantic Logfire からエクスポートされたエージェント トレースを分析し、eval を検出してデータフレーム パイプラインの形式で再現可能なアーティファクトに変換するためのツールとして使用します。 pip インストール fenic
リポジトリ: https://github.com/typedef-ai/fenic
ドキュメント: https://docs.fenic.ai
クロード コードやコーデックスなどで使用して、お気に入りのエージェント コーディング環境で fenic をすぐに使い始めることができるスキルもあります。ご意見、ご批判、その他思いついたことがあればぜひお聞かせください。質問に答えるためにここにいます。

記事本文:
GitHub - typedef-ai/fenic: 人間とエージェントのためのセマンティック データフレーム · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
typedef-ai
/
フェニック
公共
通知
通知設定を変更するにはサインインする必要があります
あ

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
245 コミット 245 コミット .claude-plugin .claude-plugin .claude .claude .cursor/ rules .cursor/ rules .github .github .trunk .trunk ベンチマーク ベンチマーク docs docs 例 例 protos/logical_plan/ v1 protos/logical_plan/ v1 Rust Rust スクリプト スクリプト仕様 仕様 src/ fenic src/ fenic テスト テスト ツール ツール .gitignore .gitignore .python-version .python-version .release-please-config.json .release-please-config.json .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス PYPI.rst PYPI.rst README.md README.md ROADMAP.md ROADMAP.md buf.gen.py.yaml buf.gen.py.yaml justfile justfile mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml ruff.toml ruff.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
fenic: 人間とエージェントのためのセマンティック データフレーム
fenic は、構造化データと非構造化データの AI 支援探索を、再利用可能で検査可能な DataFrame パイプラインに変換します。
これはセマンティック データ処理用の DataFrame クエリ エンジンであり、AI 演算子 (抽出、分類、要約、埋め込み、セマンティック結合など) がクエリ モデルに組み込まれています。これを使用して、ドキュメント、トランスクリプト、ログ、評価トレース、チケット、テーブル、API を型指定された行と反復可能なワークフローに変換します。
重要なのは、データ作業が生み出すものの変化です。人間とエージェントは同じパイプラインで作業し、どちらもパイプラインを作成、検査、再利用できます。結果は、1 回限りのプロンプトや、後でリバース エンジニアリングが必要な脆弱な正規表現スクリプトではなく、入力され、検査可能、再実行可能、呼び出し可能な耐久性のある成果物です。
pip インストール fenic
AI コーディング エージェントを使用して fenic を作成する

? fenic スキルのインストールを実行して、クロード コード/カーソル/コーデックスが正しく書き込むようにし、fenic チェックを行って lint を実行します。詳細は以下をご覧ください。
fenic はセマンティック データフレーム エンジンです。すでに知っている PySpark/SQL スタイルの操作 ( select 、 filter 、 join 、 group_by 、 agg ) を、クエリのファーストクラスの部分として言語モデルを呼び出すセマンティック演算子と並行して作成します。 Session 上でモデルを一度設定し、パイプラインを遅延構築し、自動バッチ処理、レート制限、再試行、トークン/コスト アカウンティング、応答キャッシュなどの推論用に構築されたクエリ エンジン上で fenic がコンパイルして実行します。
2 つのアイデアにより、パンダに LLM を貼り付けることとは異なります。
推論はクエリ モデル内に存在します。抽出、分類、要約、埋め込みは、手動で調整するサイドコールではなく、スキーマと型を持つ演算子です。
パイプラインは成果物です。作業は型指定された演算子として表現されるため、すでに検査可能 (行レベルの系統、 Explain 、クエリごとのメトリクス)、再実行可能 (遅延プラン + キャッシュ)、およびエージェントが呼び出すことができる名前付きテーブル、ビュー、または MCP ツールにプロモート可能です。
60秒: 乱雑なテキスト → 入力された行
脆弱な解析と 1 回限りのプロンプトを、型指定されたスキーマ バインドされた演算子に置き換えます。 Pydantic モデルとして必要な形状を定義します。 fenic はクエリ可能な構造化列を返します。
fenicをfcとしてインポート
pydantic import BaseModel 、フィールドから
クラスチケット (BaseModel):
product : str = Field ( description = "ユーザーが問い合わせている製品" )
センチメント : str = フィールド (説明 = "ポジティブ、ニュートラル、またはネガティブ" )
issue : str = フィールド ( description = "ユーザーの問題の 1 行の概要" )
セッション = fc 。セッション 。 get_or_create (
fc 。 SessionConfig (
app_name = "クイックスタート" ,
セマンティック = fc 。 SemanticConfig (
言語モデル = {
「ミニ」：fc. OpenAILanguageModel ( m

odel_name = "gpt-4o-mini" 、rpm = 500 、tpm = 200_000 )
}、
）、
)
)
df = セッション 。 create_dataframe ([
{ "id" : 1 , "text" : "レポートの CSV エクスポートは、最後の更新以降、タイムアウトになり続けます。 }、
{ "id" : 2 , "text" : "新しいダッシュボードは気に入っていますが、モバイルでは SSO ログインが機能しません。" }、
])
# フリーテキスト -> 入力されたクエリ可能な行
チケット = (
DF 。 select ( "id" , fc . semantic . extract ( "text" , Ticket ). alias ( "t" ))
。ネスト解除 ( "t" )
)
チケット。表示 ()
＃ID |製品 |感情 |問題
＃1 |レポート |負の | CSV エクスポートがタイムアウトする
＃2 |認証 |負の |モバイルで SSO ログインが壊れる
使用するプロバイダーの API キーを設定します。
import OPENAI_API_KEY=... # または ANTHROPIC_API_KEY / GOOGLE_API_KEY / COHERE_API_KEY / OPENROUTER_API_KEY
なぜフェニックなのか？
非構造化データはどこにでもありますが、それを扱うのは脆弱です。チームは正規表現、一回限りのスクリプト、ノートブック、プロンプト チェーンを利用して、ドキュメント、ログ、チケット、トランスクリプト、トレースから意味を引き出します。結果を再現するのは難しく、検査するのも困難です。
エージェントは探索を容易にし、新たな問題を引き起こしました。エージェントは乱雑なデータを調べて有用なものを見つけることができますが、その発見がコード、データ、またはパイプラインにならない限り、チャットの記録として消えてしまいます。次の人は何が起こったのかをリバースエンジニアリングする必要があります。
fenic は、セマンティック データ作業に DataFrame 抽象化を提供します。探索を fenic オペレーターとして表現すると、それはすでにアーティファクトになります。
モデルはまだ確率的ですが、その周りのパイプラインは型指定され、制限され、キャッシュされ、再生可能で検査可能です。
注目のワークフロー: 評価の探索から永続的な評価インテリジェンスまで
評価分析はその完璧なケースです。データは乱雑で半構造化されており (トレース、ツール呼び出し、出力、判定メモ)、有用なパターンは通常、一度発見されるだけで失われます。エージェントのfenicと

(またはあなたが) トレースを探索すると、有用な操作が新しいバッチごとに再実行されるパイプラインになります。
fenicをfcとしてインポート
import リテラルの入力から
pydantic import BaseModel 、フィールドから
クラス FailureMode (BaseModel):
failed : bool = フィールド ( description = "エージェントがタスクに失敗したかどうか" )
category : リテラル [ "tool_error" 、 "instruction_following" 、 "retrieval" 、 "reasoning" 、 "none" ] = フィールド (
description = "主な失敗カテゴリ、または実行が成功した場合は「なし」
)
証拠 : str = フィールド ( description = "分類を正当化する短い引用または要約" )
セッション = fc 。セッション 。 get_or_create (
fc 。 SessionConfig (
app_name = "eval_triage" ,
セマンティック = fc 。 SemanticConfig (
言語モデル = {
「ミニ」：fc. OpenAILanguageModel (モデル名 = "gpt-4o-mini" 、rpm = 500、tpm = 200_000)
}、
）、
)
)
# Eval トレースはほとんどの場合 JSON です。ここでは実行ごとに 1 つのファイル (JSONL ファイルまたは
#倉庫テーブル作業も)。 「content」はJSONとしてロードされます。
トレース = セッション 。読む 。 docs ( "eval_runs/**/*.json" 、 content_type = "json" 、 recursive = True )
# オプション: トークンコストを制御するためにモデルが参照するフィールドのみを投影します。
# トレース =traces.with_column("convo", fc.json.jq(fc.col("content"), ".messages"))
失敗 = (
痕跡
。 with_column ( "分析" , fc . semantic . extract ( fc .col ( "content" ). Cast ( fc . StringType ), FailureMode ))
。 unnest (「分析」)
。フィルター ( fc .col ( "失敗" ))
)
# 障害カテゴリごとに 1 つの検査可能な根本原因の概要
Failure_modes = 失敗。 group_by (「カテゴリー」)。集約 (
fc 。カウント ( "*" )。エイリアス ( "n" )、
fc 。意味論的。減らす(
"これらの障害に共通する根本原因を要約します" ,
列 = fc 。列 (「証拠」)、
）。エイリアス (「パターン」)、
)
失敗モード 。書く 。 save_as_table ( "failure_modes" , mode = "over

書いてください」）
探索は累積的になりました。次のモデル バージョンで再実行して回帰を検出し、failures.lineage() を使用して任意の行をソース トレースに戻して検査し、(以下を参照) 結果をツールとしてエージェントに渡します。新しい分析はそれぞれ、拡大するライブラリ内の別の再利用可能な変換になります。
[切り捨てられた]
意味とメタデータを一緒にクエリする
興味深い質問には通常、顧客のメタデータとサポート チケット、評価スコアと軌跡、CRM フィールドと通話記録など、構造化データと非構造化データの両方が必要です。 fenic は、正確なキーではなく意味に基づく結合など、同じパイプラインでリレーショナルおよびセマンティックな作業を実行します。
# 正確な値ではなく、意味に基づいて一致します
一致 = 候補 。意味論的。参加（
役割、
述語 = (
"候補者の背景: {{ left_on }} \n "
"役割の要件: {{ right_on }} \n "
「候補者はその役割にぴったりです。」
）、
left_on = fc 。列 (「再開」)、
right_on = fc 。列 ( "ジョブの説明" )、
)
# ...その後、通常の DataFrame 操作を使用してグループ化、集計、ランク付けを行います
ランク付け = (
一致します。 group_by ( "role_id" )
。 agg ( fc . count ( "*" ). alias ( "n_candidates" ))
。 order_by ( fc . desc ( "n_candidates" ))
)
エージェントが再利用するアーティファクトにする
fenic のテーブルまたはビューは、MCP (モデル コンテキスト プロトコル) を介して管理されたツール サーフェスになります。そのため、人間が検査するのと同じパイプラインが、型指定されたパラメーターと制限された結果サイズを持つエージェントによって呼び出されます。
テーブルに対して一連のシステム ツール (スキーマ、プロファイル、読み取り、検索、分析) を自動生成します。
fenicインポートからSystemToolConfig
セッション。カタログもset_table_description (
"failure_modes" , "エージェントの定期的な障害モードとその数と根本原因の概要"
)
サーバー = fc 。 create_mcp_server (
セッション、
「エヴァル・インテリジェンス」、
system_tools = SystemToolConfig (
table_names = [ "ファイ

ルアー_モード" ],
tool_namespace = "evals" ,
max_result_rows = 100 、
）、
)
fc 。 run_mcp_server_sync (サーバー、トランスポート = "http"、ポート = 8000)
または、型指定された SQL マクロのような、正確なパラメーター化されたツールを定義します。
最近 = セッション。テーブル (「failure_modes」)。フィルター (
fc 。 Col ( "カテゴリ" ) == fc . tool_param ( "カテゴリ" , fc . StringType )
)
セッション。カタログもcreate_tool (
"カテゴリ別の失敗" ,
"特定のカテゴリで繰り返される障害を検索する" ,
最近、
tool_params = [ fc . ToolParam ( name = "カテゴリ" 、説明 = "フィルタリングする失敗カテゴリ" )],
)
CLI を使用して、登録されているすべてのツールをカタログから直接提供します。
fenic-serve --アプリ名 eval_triage --ポート 8000
AI コーディング エージェントを使用して fenic を作成する
fenic は新しいため、コーディング エージェントはすぐにその正確な API を常に認識できるわけではありません。 2 つのツールにより、エージェントが確実に fenic を作成できるようになります。
糸くずが出る。 fenic チェックは、インストールされている fenic に対してパイプラインの fc.* シンボルを静的に解決し、名前空間/インポートの間違い ( fenic.functions 、 fc.array と fc.arr 、 fc.explode など) にフラグを立てます。スクリプトは実行されません。 fenic を作成した後にエージェントに実行してもらい、報告された内容を修正してもらいます。
fenic チェックのパイプライン.py
# → {"ok": true, "findings": []} または修正の提案を含む名前空間/シンボルの検出

[切り捨てられた]

## Original Extract

Semantic DataFrames for humans and agents. Contribute to typedef-ai/fenic development by creating an account on GitHub.

Hey friends. I'd like to share a project that's dear to me. fenic is a dataframe API with LLMs added as first-class citizens, a classic lazy dataframe API extended with new operators that are backed by LLMs. What this gets you is the ability to work with structured and unstructured data in the same context. Most importantly, the LLMs aren't integrates as opaque UDF black boxes. They're exposed as "semantic" operators that the planner can reason about alongside the classic ones. (There are examples and code snippets on the repo to see how everything works together) Why build this? I'm a data infra / systems person. When LLMs showed up, what I saw was a new type of compute that changes the characteristics of the
workloads we deal with. I wanted to experiment with how our current systems can absorb these new workloads and compute types, and what it would take to make the DX as seamless as possible, that's where the UDF + arbitrary prompt was feeling too problematic. To support this properly, we had to introduce a few really cool things: New plan operators. You don't just send prompts at an LLM. You use operators like semantic join, semantic map and reduce, and semantic filter, among others. They mix with the classic operators, and because the planner sees them as real operators rather than black boxes, it can reorder work around them. Typed outputs. There's ergonomics to turn the output of a semantic operator straight into a typed dataframe column. A Pydantic schema for the LLM output becomes a typed struct column you can unnest, explode, and so on. New data types like a markdown data type. Markdown became an important way to share information with LLMs, even though it started life as a way to format text for presentation. It carries structure, and being able to access that structure the way you would a struct or JSON type adds to the developer experience I mentioned. Async UDFs. One of the more interesting shifts in workloads from the LLM explosion is the need to put heavily I/O-bound steps in your pipeline: fetching a response from an API, crawling a website, and so on. Async UDFs fill that gap, and the implementation handles the nuances for you: concurrency, retries, and the rest. An LLM-inference-aware planner and runtime. This is one of the parts I'm most excited about, and there's a lot still to do. Today: identical prompts within a batch collapse to a single model call, so duplicates cost zero tokens; requests are dispatched concurrently under per-provider rpm/tpm limits with retries and backoff; null and empty cells skip the model entirely; and you get token and cost metrics per operator. There's also an optional persistent response cache so re-runs skip the model. MCP as a new catalog primitive. Much like a registered view, you can register a dataframe pipeline as an MCP tool in the catalog. fenic then serves an MCP server with that pipeline as the tool's logic, executed over your data. These are just some of what's gone into fenic while experimenting with how LLMs can become part of our compute infrastructure. There's more, and plenty more to polish on what's already there. I've been using fenic for all sorts of things. On the small/personal end, I use it to take my podcast audio recordings and turn them into nicely structured tables of metadata I can research. On the heavier end, I use it as tooling for agents to analyze agent traces exported from Pydantic Logfire, to discover evals and turn them into reproducible artifacts in the form of dataframe pipelines. pip install fenic
Repo: https://github.com/typedef-ai/fenic
Docs: https://docs.fenic.ai
There's also a skill you can use with claude code, codex etc. to quickly get started with fenic in your favourite agentic coding environment. I'd love to hear your thoughts, criticism, and anything else that comes to mind. I'm here to answer questions.

GitHub - typedef-ai/fenic: Semantic DataFrames for humans and agents · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
typedef-ai
/
fenic
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
245 Commits 245 Commits .claude-plugin .claude-plugin .claude .claude .cursor/ rules .cursor/ rules .github .github .trunk .trunk benchmarks benchmarks docs docs examples examples protos/ logical_plan/ v1 protos/ logical_plan/ v1 rust rust scripts scripts specs specs src/ fenic src/ fenic tests tests tools tools .gitignore .gitignore .python-version .python-version .release-please-config.json .release-please-config.json .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PYPI.rst PYPI.rst README.md README.md ROADMAP.md ROADMAP.md buf.gen.py.yaml buf.gen.py.yaml justfile justfile mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml ruff.toml ruff.toml uv.lock uv.lock View all files Repository files navigation
fenic: semantic DataFrames for humans and agents
fenic turns AI-assisted exploration of structured and unstructured data into reusable, inspectable DataFrame pipelines.
It's a DataFrame query engine for semantic data processing, with AI operators — extract , classify , summarize , embed , semantic join , and more — built into the query model. Use it to turn documents, transcripts, logs, eval traces, tickets, tables, and APIs into typed rows and repeatable workflows.
The point is a shift in what your data work produces . Humans and agents work on the same pipelines — both can author, inspect, and reuse them. The result isn't a one-off prompt or a brittle regex script that has to be reverse-engineered later — it's a durable artifact: typed, inspectable, rerunnable, and callable.
pip install fenic
Writing fenic with an AI coding agent? Run fenic skill install so Claude Code / Cursor / Codex write it correctly, and fenic check to lint it — details below .
fenic is a semantic DataFrame engine . You write the PySpark/SQL-style operations you already know — select , filter , join , group_by , agg — alongside semantic operators that call language models as a first-class part of the query. You configure models once on a Session , build a pipeline lazily, and fenic compiles and runs it on a query engine built for inference: automatic batching, rate limiting, retries, token/cost accounting, and response caching.
Two ideas make it different from gluing an LLM onto pandas:
Inference lives inside the query model. Extraction, classification, summarization, and embeddings are operators with schemas and types — not side calls you orchestrate by hand.
The pipeline is the artifact. Because the work is expressed as typed operators, it's already inspectable (row-level lineage, explain , per-query metrics), rerunnable (lazy plans + caching), and promotable into a named table, view, or MCP tool an agent can call.
60 seconds: messy text → typed rows
Replace brittle parsing and one-off prompts with a typed, schema-bound operator. Define the shape you want as a Pydantic model; fenic returns structured columns you can query.
import fenic as fc
from pydantic import BaseModel , Field
class Ticket ( BaseModel ):
product : str = Field ( description = "The product the user is asking about" )
sentiment : str = Field ( description = "positive, neutral, or negative" )
issue : str = Field ( description = "One-line summary of the user's problem" )
session = fc . Session . get_or_create (
fc . SessionConfig (
app_name = "quickstart" ,
semantic = fc . SemanticConfig (
language_models = {
"mini" : fc . OpenAILanguageModel ( model_name = "gpt-4o-mini" , rpm = 500 , tpm = 200_000 )
},
),
)
)
df = session . create_dataframe ([
{ "id" : 1 , "text" : "The CSV export in Reports keeps timing out since the last update." },
{ "id" : 2 , "text" : "Love the new dashboard, but SSO login is broken on mobile." },
])
# Free text -> typed, queryable rows
tickets = (
df . select ( "id" , fc . semantic . extract ( "text" , Ticket ). alias ( "t" ))
. unnest ( "t" )
)
tickets . show ()
# id | product | sentiment | issue
# 1 | Reports | negative | CSV export times out
# 2 | Auth | negative | SSO login broken on mobile
Set the API key for whichever provider you use:
export OPENAI_API_KEY=... # or ANTHROPIC_API_KEY / GOOGLE_API_KEY / COHERE_API_KEY / OPENROUTER_API_KEY
Why fenic?
Unstructured data is everywhere, and working with it is brittle. Teams reach for regex, one-off scripts, notebooks, and prompt chains to pull meaning out of documents, logs, tickets, transcripts, and traces. The results are hard to reproduce and hard to inspect.
Agents made exploration easy and introduced a new problem. An agent can dig through messy data and find something useful — but unless that discovery becomes code, data, or a pipeline, it dies as a chat transcript. The next person has to reverse-engineer what happened.
fenic gives semantic data work a DataFrame abstraction. Express the exploration as fenic operators and it's already the artifact:
The model is still probabilistic — but the pipeline around it is typed, bounded, cached, replayable, and inspectable.
Featured workflow: from eval exploration to durable eval intelligence
Eval analysis is the perfect case: the data is messy and semi-structured (traces, tool calls, outputs, judge notes), and the useful patterns usually get discovered once and lost. With fenic, an agent (or you) explores the traces, and the useful operations become a pipeline that reruns on every new batch.
import fenic as fc
from typing import Literal
from pydantic import BaseModel , Field
class FailureMode ( BaseModel ):
failed : bool = Field ( description = "Whether the agent failed the task" )
category : Literal [ "tool_error" , "instruction_following" , "retrieval" , "reasoning" , "none" ] = Field (
description = "Primary failure category, or 'none' if the run succeeded"
)
evidence : str = Field ( description = "Short quote or summary justifying the classification" )
session = fc . Session . get_or_create (
fc . SessionConfig (
app_name = "eval_triage" ,
semantic = fc . SemanticConfig (
language_models = {
"mini" : fc . OpenAILanguageModel ( model_name = "gpt-4o-mini" , rpm = 500 , tpm = 200_000 )
},
),
)
)
# Eval traces are almost always JSON — one file per run here (a JSONL file or a
# warehouse table work too). `content` is loaded as JSON.
traces = session . read . docs ( "eval_runs/**/*.json" , content_type = "json" , recursive = True )
# Optional: project just the fields the model should see to control token cost, e.g.
# traces = traces.with_column("convo", fc.json.jq(fc.col("content"), ".messages"))
failures = (
traces
. with_column ( "analysis" , fc . semantic . extract ( fc . col ( "content" ). cast ( fc . StringType ), FailureMode ))
. unnest ( "analysis" )
. filter ( fc . col ( "failed" ))
)
# One inspectable, root-caused summary per failure category
failure_modes = failures . group_by ( "category" ). agg (
fc . count ( "*" ). alias ( "n" ),
fc . semantic . reduce (
"Summarize the common root cause across these failures" ,
column = fc . col ( "evidence" ),
). alias ( "pattern" ),
)
failure_modes . write . save_as_table ( "failure_modes" , mode = "overwrite" )
The exploration is now cumulative : rerun it on the next model version to detect regressions, inspect any row back to its source trace with failures.lineage() , and — see below — hand the result to the agent as a tool. Each new analysis becomes another reusable transform in a growing library for und
[truncated]
Query meaning and metadata together
The interesting questions usually need both structured and unstructured data — customer metadata and support tickets, eval scores and trajectories, CRM fields and call transcripts. fenic does relational and semantic work in the same pipeline, including joins on meaning rather than exact keys.
# Match on meaning, not exact values
matches = candidates . semantic . join (
roles ,
predicate = (
"Candidate background: {{ left_on }} \n "
"Role requirements: {{ right_on }} \n "
"The candidate is a strong fit for the role."
),
left_on = fc . col ( "resume" ),
right_on = fc . col ( "job_description" ),
)
# ...then group, aggregate, and rank with ordinary DataFrame ops
ranked = (
matches . group_by ( "role_id" )
. agg ( fc . count ( "*" ). alias ( "n_candidates" ))
. order_by ( fc . desc ( "n_candidates" ))
)
Make it an artifact your agents reuse
A fenic table or view becomes a governed tool surface over MCP (Model Context Protocol) — so the same pipeline a human inspects is what an agent calls, with typed parameters and bounded result sizes.
Auto-generate a suite of system tools (schema, profile, read, search, analyze) over your tables:
from fenic import SystemToolConfig
session . catalog . set_table_description (
"failure_modes" , "Recurring agent failure modes with counts and root-cause summaries"
)
server = fc . create_mcp_server (
session ,
"Eval Intelligence" ,
system_tools = SystemToolConfig (
table_names = [ "failure_modes" ],
tool_namespace = "evals" ,
max_result_rows = 100 ,
),
)
fc . run_mcp_server_sync ( server , transport = "http" , port = 8000 )
Or define a precise, parameterized tool — like a typed SQL macro:
recent = session . table ( "failure_modes" ). filter (
fc . col ( "category" ) == fc . tool_param ( "category" , fc . StringType )
)
session . catalog . create_tool (
"failures_by_category" ,
"Look up recurring failures for a given category" ,
recent ,
tool_params = [ fc . ToolParam ( name = "category" , description = "Failure category to filter by" )],
)
Serve every registered tool straight from the catalog with the CLI:
fenic-serve --app-name eval_triage --port 8000
Writing fenic with an AI coding agent
fenic is new, so a coding agent won't always know its exact API out of the box. Two tools make any agent reliable at writing fenic:
Lint it. fenic check statically resolves a pipeline's fc.* symbols against the installed fenic and flags namespace/import mistakes ( fenic.functions , fc.array vs fc.arr , fc.explode , …) — it does not execute your script. Have your agent run it after writing fenic and fix what it reports:
fenic check pipeline.py
# → {"ok": true, "findings": []} or a namespace/symbol finding with a fix suggest

[truncated]
