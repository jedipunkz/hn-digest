---
source: "https://github.com/saineshnakra/automated-data-analyst"
hn_url: "https://news.ycombinator.com/item?id=48962405"
title: "Ada: An AI business intelligence software from CSV and Excel(yes LLMs but more)"
article_title: "GitHub - saineshnakra/automated-data-analyst: I tried to automate a data analyst · GitHub"
author: "saineshnakra"
captured_at: "2026-07-18T21:40:59Z"
capture_tool: "hn-digest"
hn_id: 48962405
score: 2
comments: 0
posted_at: "2026-07-18T21:06:49Z"
tags:
  - hacker-news
  - translated
---

# Ada: An AI business intelligence software from CSV and Excel(yes LLMs but more)

- HN: [48962405](https://news.ycombinator.com/item?id=48962405)
- Source: [github.com](https://github.com/saineshnakra/automated-data-analyst)
- Score: 2
- Comments: 0
- Posted: 2026-07-18T21:06:49Z

## Translation

タイトル: Ada: CSV と Excel による AI ビジネス インテリジェンス ソフトウェア (LLM ですがそれ以上)
記事タイトル: GitHub - saineshnakra/automated-data-analyst: データアナリストを自動化してみた · GitHub
説明: データ アナリストを自動化しようとしました。 GitHub でアカウントを作成して、saineshnakra/automated-data-analyst の開発に貢献してください。

記事本文:
GitHub - saineshnakra/automated-data-analyst: データ アナリストを自動化してみた · GitHub
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
サイネシュナクラ
/
自動データアナリスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
さ

ineshnakra/自動データアナリスト
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .github .github .streamlit .streamlit アセット アセット テスト テスト .gitignore .gitignore CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md ai_insights.py ai_insights.py Analysis.py Analysis.py anomalies.py anomalies.py app.py app.py business_insights.py business_insights.py Demon_data.py Demon_data.py file_io.py file_io.py Forecasting.py Forecasting.py nlq.py nlq.py Pipeline.py Pipeline.py pyproject.toml pyproject.toml 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt ui.py ui.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ADA — CSV と Excel による AI を活用したビジネス インテリジェンス
ビジネスファイルをドロップします。ダッシュボード、その背後にある証拠、そして次のアクションを取得します。
ライブ ダッシュボードを試す · エンジニアリング ストーリーを読む · ロードマップを見る · 貢献する · バグを報告する
ADA は、BI ツールを構成せずに回答を必要とするオペレーター向けのオープンソースの自動データ アナリストです。 CSV、XLSX、または XLSM ファイルをアップロードすると、ADA がファイルをクリーンアップし、そのビジネス スキーマを検出し、インタラクティブな Plotly ダッシュボードを作成し、異常な期間にフラグを立て、保護されたベースライン予測を予測し、重大な変更について説明し、次に何を調査するかを推奨します。また、データに関する平易な英語の質問に対して、すべての返信の背後にある計算を使用して回答します。
これは、分析ソフトウェアでは困難になることが多い単純なユースケース向けに設計されており、初めてのユーザーでもスプレッドシートをアップロードしてビジネスで何が起こっているかを理解できる必要があります。
ビジネス上の質問をします。数値とその計算を取得します。
1 つのセグメントに焦点を当てます。全体を見る

分析自体が再編成されます。
異常、予測、要因、証拠は、生成された散文の背後に隠れるのではなく、検査可能なままになります。
製品の背後にあるアーキテクチャ、トレードオフ、および障害モードについては、「幻覚を知らせる AI データ アナリストを構築しました」を参照してください。
ほとんどの CSV アナライザーはチャートで停止します。 ADA は 4 つのレイヤーを明示的に保持します。
すべての証拠カードとチャットの回答は、その計算を明らかにします。決定論的製品は、モデルが構成されているかどうかに関係なく、引き続き権威を持ちます。
CSV AI ツールによるチャット
従来のBI
ADA
セットアップ
アップロードしてプロンプトを表示する
データモデリング、数週間
アップロードのみ
答え
もっともらしい散文。隠された推論
そのとおりですが、すべてのグラフを作成するのはあなたです
示された数学による決定論的な計算
モデルに送信される行
通常
ベンダーによって異なります
決してしない — オプションの AI はスキーマと証拠のみを認識します
異常と予測
リクエストに応じて、検証不能
有料アドオン
バックテストされたエラーが組み込まれており、読み取ることができます
コスト
定期購入
ライセンス
無料でMITライセンスを取得
スプレッドシートから意思決定まで
フローチャート TD
A["CSV または Excel ワークシートをアップロードする"] --> B["クリーンアップして型を推測する"]
B --> C["メトリクス、日付、セグメント、ID を検出"]
C --> D[「証拠、異常、予測を計算する」]
D --> E[「ダッシュボード、エグゼクティブブリーフ、およびドリルダウン」]
C --> G["ADA に質問する: 質問 → QueryPlan → ローカル実行"]
G-。 "スキーマ + 質問のみ、ルールが解析できない場合" .-> H["オプションの AI クエリ プランナー"]
D-。 "計算された証拠のみ" .-> F["オプションの AI 戦略読み取り"]
読み込み中
ADA は以下を自動的に検索します。
収益、売上、利益、コスト、金額、単位などの主要な結果
期間の動き、異常検出、ベースライン予測のための時間フィールド
製品、カテゴリ、チャネル、地域、顧客、ステータスなどの有用なセグメント
識別子、欠損、外れ値、濃度、および数値関係
s

観察された事実とは切り離された、証拠に裏付けられた最も強力な次の調査
ソース スキーマが異常な場合、ユーザーはダッシュボードを再構築せずに、検出されたメトリクス、日付、セグメントをオーバーライドし、分析全体を 1 つのセグメント スライスにドリルできます。
統合デモが含まれたゼロ構成の CSV および Excel 分析
Ask ADA : わかりやすい英語の質問 (合計、ランキング、内訳、傾向、増加、カウント、時間、セグメント フィルター) は、示されている計算を使用してローカルで回答されます。
異常レーダー: 堅牢な傾向線バンドの外側の期間は、チャート、証拠台帳、および推奨アクションでフラグが付けられます。
年間月ごとの季節性、不確実性バンド、およびバックテストされた誤差をチャートの横に印刷した、保護されたベースライン予測
ドリルダウンの焦点: 1 つのセグメント値を分析し、次に有用なディメンションによって自動的に再グループ化します。
セグメントごとの最新の変化を調整する動きのウォーターフォールと、セグメントごとの期間ごとの強度ヒートマップ
マルチシート Excel ワークブック用のワークシート ピッカー
保守的なクリーンアップ、型推論、重複の削除、および目に見えるクリーニング監査
エグゼクティブの見出し、4 つのビジネス KPI、および平易な英語のブリーフィング
表示されるすべての信号の背後にある計算を含む証拠台帳
決定論的な証拠に関連付けられた優先順位付けされた推奨事項
オプションの AI クエリ プランナーと OpenAI Responses API を使用した構造化戦略合成
ダウンロード可能な Markdown エグゼクティブ ブリーフとクリーンな CSV
技術者以外のユーザー向けに構築されたレスポンシブな Streamlit インターフェイス
予測可能なホストパフォーマンスのためのファイル制限と行キャップ
ADA は API キーがなくても完全に機能します。決定論的モードでは、ルールベースのパーサーが自ら計画できるすべての Ask ADA 回答を含め、モデル呼び出しは行われません。
キーが存在する場合、入力と実行の両方で 2 つのナロー モデル呼び出しが利用可能になります。

同じローカル エンジンに対して:
クエリ プランナー — 決定論的パーサーが質問を読み取れない場合にのみ使用されます。列スキーマ (名前、タイプ、ロール) と質問自体を受け取り、型指定された QueryPlan を出力し、ADA はそのプランをローカルで実行します。解決できない計画は推測ではなく拒否され、AI が計画した回答には目に見えるバッジが付けられます。
戦略的読み取り — 計算されたスキーマ、概要、証拠カード、推奨事項、およびユーザーが提供したコンテキストのみを受け取ります。
どちらの呼び出しでも、アップロードされた行やセルの値はモデル プロンプトに入力されません。応答は型指定された Pydantic スキーマと一致する必要があり、ストレージは要求に対して無効になっており、安全性制御にはハッシュ化された匿名セッション ID が使用されます。
デフォルトのモデルは gpt-5.6-luna で、効率的な戦略読み取りのための推論が低くなります。中程度の推論を備えた gpt-5.6-terra は、より高いコストを正当化するほど決定が曖昧な場合に使用できます。モデル呼び出しはボタンでトリガーされ、偶発的な支出を避けるために証拠ペイロードごとにキャッシュされます。
完全なデータ処理および秘密管理ポリシーについては、SECURITY.md を参照してください。
パス
責任
app.py
Thin Streamlit オーケストレーションとセッション状態
パイプライン.py
限定された準備、クリーニング、スキーマの選択、ドリルダウン、および監査フレーム
分析.py
保守的なクリーニングとデータプロファイリング
business_insights.py
スキーマの検出、計算、証拠、および決定論的な推奨事項
nlq.py
自然言語の質問 → 監査可能なクエリ プラン → ローカル実行
anomalies.py
期間集計にわたる堅牢な傾向線の異常検出
予測.py
季節性と目に見えるバックテストを備えた、保護されたベースライン予測
ai_insights.py
オプションの型付き応答 API クエリの計画と証拠の統合
ui.py
再利用可能なプレゼンテーション コンポーネントと Plotly スタイル設定
file_io.py
wo による検証済みの CSV および Excel 解析

rkシートの選択
テスト/
ユニット、プライバシー コントラクト、パイプライン、ビジネス ロジック、レンダリング テスト
コードベースは、純粋な分析関数とモデル境界での依存関係の注入を優先します。これにより、Streamlit、ネットワーク アクセス、API クレジットなしでビジネス エンジンをテストできるようになります。
git clone https://github.com/saineshnakra/automated-data-analyst.git
cd 自動データアナリスト
Python -m venv .venv
ソース .venv/bin/activate # Windows: .venv\Scripts\activate
python -m pip install -rrequirements.txt
streamlit で app.py を実行
秘密は必要ありません。訪問者はセッション専用のサイドバーフィールドに独自の API キーを入力できます。信頼できるプライベート デプロイメントでは、代わりに環境または .streamlit/secrets.toml で OPENAI_API_KEY を設定できます。
OPENAI_API_KEY = " あなたのキー "
そのファイルは決してコミットしないでください。すでに無視されています。認証と支出の制御も追加しない限り、所有者が資金提供するキーをパブリック展開に配置することは避けてください。
python -m pip install -r 要件-dev.txt
ラフチェック。
python -m 単体テスト 検出 -s テスト -v
GitHub Actions は、プッシュおよびプル リクエストごとに lint、完全なテスト スイート、バイトコード コンパイルを実行します。
私のデータは私のマシンから離れますか?
いいえ。クリーニング、スキーマ検出、すべてのチャート、すべての証拠カード、およびすべての Ask ADA の回答は、パンダを使用してローカルに計算されます。 AI レイヤーを選択すると、列スキーマと計算された証拠のみが送信され、行やセル値は送信されません。
OpenAI API キーは必要ですか?
いいえ、ADA はアナリストなしで完全なアナリストです。キーは、珍しい質問と戦略的なナラティブに対するクエリ プランナーのフォールバックを追加するだけです。
どのような形式を分析できますか?
CSV (カンマ、セミコロン、またはタブ区切り)、XLSX、および XLSM — マルチシート ワークブックからの特定のワークシートの選択を含みます。
これはチャットボットに CSV を貼り付けることとどう違うのですか?
チャットボットで流暢に話せます

散文的には監査できず、行はプロンプトの一部になります。 ADA は質問を明示的なクエリ プランに変換し、マシン上のパンダで実行し、すべての回答の下に計算を出力します。
セルフホストできますか?
はい、これは標準の Streamlit アプリです。 pip install -rrequirements.txt && streamlit app.py を実行するか、Python を実行する任意のホストにデプロイします。
特に、新しい決定論的メトリクス、Ask ADA の質問形式、スキーマ検出フィクスチャ、チャートのアクセシビリティ、ファイル形式、および敵対的テスト データセットに関する貢献を歓迎します。最初の良好な問題から始めて、 CONTRIBUTING.md を読み、ロードマップから項目を選択するか、焦点を絞った提案を開きます。
優れた貢献により、洞察がより正確になり、より説明しやすくなり、技術者以外のユーザーでも簡単に行動できるようになります。すべての新しい推奨事項には、テストとそれをサポートする計算が含まれている必要があります。
ADA が役立つ場合、星印は他のオペレーターが ADA を見つけるのに役立ちます。
app.pyをStreamlitにデプロイします。リポジトリには、アプリのテーマ、依存関係マニフェスト、サーバーのアップロード制限、ヘッドレス構成が含まれます。オプションの戦略レイヤーを使用できる必要がある場合にのみ、Streamlit のシークレット マネージャーを介して OPENAI_API_KEY を追加します。
データアナリストを自動化してみた
読み込み中にエラーが発生しました

[切り捨てられた]

## Original Extract

I tried to automate a data analyst. Contribute to saineshnakra/automated-data-analyst development by creating an account on GitHub.

GitHub - saineshnakra/automated-data-analyst: I tried to automate a data analyst · GitHub
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
saineshnakra
/
automated-data-analyst
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
saineshnakra/automated-data-analyst
main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .github .github .streamlit .streamlit assets assets tests tests .gitignore .gitignore CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md ai_insights.py ai_insights.py analysis.py analysis.py anomalies.py anomalies.py app.py app.py business_insights.py business_insights.py demo_data.py demo_data.py file_io.py file_io.py forecasting.py forecasting.py nlq.py nlq.py pipeline.py pipeline.py pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt ui.py ui.py View all files Repository files navigation
ADA — AI-powered business intelligence from CSV and Excel
Drop in a business file. Get a dashboard, the evidence behind it, and the next action.
Try the live dashboard · Read the engineering story · See the roadmap · Contribute · Report a bug
ADA is an open-source automated data analyst for operators who need answers without configuring a BI tool. Upload a CSV, XLSX, or XLSM file and ADA cleans it, detects its business schema, creates an interactive Plotly dashboard, flags anomalous periods, projects a guarded baseline forecast, explains material changes, recommends what to investigate next — and answers plain-English questions about the data with the calculation behind every reply.
It is designed for the simple use case analytics software often makes difficult: even a first-time user should be able to upload a spreadsheet and understand what is happening in the business.
Ask a business question. Get the number and its calculation.
Focus on one segment. Watch the entire analysis regroup itself.
Anomalies, forecasts, drivers, and evidence stay inspectable instead of disappearing behind generated prose.
For the architecture, tradeoffs, and failure modes behind the product, read I Built an AI Data Analyst That Tells You When It Hallucinates .
Most CSV analyzers stop at charts. ADA keeps four layers explicit:
Every evidence card and chat answer exposes its calculation. The deterministic product remains authoritative whether or not a model is configured.
Chat-with-CSV AI tools
Traditional BI
ADA
Setup
Upload and prompt
Data modeling, weeks
Upload only
Answers
Plausible prose; reasoning hidden
Exact, but you build every chart
Deterministic calculations with the math shown
Rows sent to a model
Usually
Depends on vendor
Never — optional AI sees schema and evidence only
Anomalies and forecasts
On request, unverifiable
Paid add-ons
Built in, with a backtested error you can read
Cost
Subscription
License
Free and MIT-licensed
From spreadsheet to decision
flowchart TD
A["Upload CSV or Excel worksheet"] --> B["Clean and infer types"]
B --> C["Detect metric, date, segment, ID"]
C --> D["Calculate evidence, anomalies, and forecast"]
D --> E["Dashboard, executive brief, and drill-down"]
C --> G["Ask ADA: question → QueryPlan → local execution"]
G -. "schema + question only, when rules cannot parse" .-> H["Optional AI query planner"]
D -. "computed evidence only" .-> F["Optional AI strategic read"]
Loading
ADA automatically looks for:
A primary outcome such as revenue, sales, profit, cost, amount, or units
A time field for period movement, anomaly detection, and the baseline forecast
A useful segment such as product, category, channel, region, customer, or status
Identifiers, missingness, outliers, concentration, and numeric relationships
The strongest evidence-backed next investigation, separated from observed fact
If the source schema is unusual, users can override the detected metric, date, and segment without rebuilding the dashboard — and drill the whole analysis into a single segment slice.
Zero-configuration CSV and Excel analytics with an included synthetic demo
Ask ADA : plain-English questions (totals, rankings, breakdowns, trends, growth, counts, time and segment filters) answered locally with the calculation shown
Anomaly radar: periods outside a robust trendline band are flagged on the chart, in the evidence ledger, and in the recommended actions
Guarded baseline forecast with month-of-year seasonality, an uncertainty band, and its backtested error printed next to the chart
Drill-down focus: analyze one segment value and automatically regroup by the next useful dimension
Movement waterfall reconciling the latest change by segment, plus a segment-by-period intensity heatmap
Worksheet picker for multi-sheet Excel workbooks
Conservative cleanup, type inference, duplicate removal, and a visible cleaning audit
Executive headline, four business KPIs, and plain-English briefing
Evidence ledger with the calculation behind every displayed signal
Prioritized recommendations linked to deterministic evidence
Optional AI query planner and structured strategy synthesis using the OpenAI Responses API
Downloadable Markdown executive brief and cleaned CSV
Responsive Streamlit interface built for non-technical users
File limit and row cap for predictable hosted performance
ADA works fully without an API key. In deterministic mode, no model call is made — including every Ask ADA answer the rule-based parser can plan itself.
When a key is present, two narrow model calls become available, both typed and executed against the same local engine:
Query planner — used only when the deterministic parser cannot read a question. It receives the column schema (names, types, roles) and the question itself, emits a typed QueryPlan , and ADA executes that plan locally. Unresolvable plans are refused rather than guessed, and AI-planned answers are visibly badged.
Strategic read — receives only the calculated schema, summaries, evidence cards, recommendations, and user-supplied context.
Neither call puts uploaded rows or cell values into the model prompt. Responses must match a typed Pydantic schema, storage is disabled for the request, and a hashed anonymous session identifier is used for safety controls.
The default model is gpt-5.6-luna with low reasoning for an efficient strategic read. gpt-5.6-terra with medium reasoning is available when the decision is ambiguous enough to justify higher cost. Model calls are button-triggered and cached per evidence payload to avoid accidental spend.
See SECURITY.md for the complete data-handling and secret-management policy.
Path
Responsibility
app.py
Thin Streamlit orchestration and session state
pipeline.py
Bounded preparation, cleaning, schema selection, drill-down, and audit frames
analysis.py
Conservative cleaning and data profiling
business_insights.py
Schema detection, calculations, evidence, and deterministic recommendations
nlq.py
Natural-language questions → auditable query plans → local execution
anomalies.py
Robust trendline anomaly detection over period aggregates
forecasting.py
Guarded baseline forecast with seasonality and a visible backtest
ai_insights.py
Optional typed Responses API query planning and evidence synthesis
ui.py
Reusable presentation components and Plotly styling
file_io.py
Validated CSV and Excel parsing with worksheet selection
tests/
Unit, privacy-contract, pipeline, business-logic, and rendering tests
The codebase favors pure analysis functions and dependency injection at the model boundary. That keeps the business engine testable without Streamlit, network access, or API credits.
git clone https://github.com/saineshnakra/automated-data-analyst.git
cd automated-data-analyst
python -m venv .venv
source .venv/bin/activate # Windows: .venv\Scripts\activate
python -m pip install -r requirements.txt
streamlit run app.py
No secret is required. A visitor can enter their own API key in the session-only sidebar field. A trusted private deployment can instead set OPENAI_API_KEY in the environment or in .streamlit/secrets.toml :
OPENAI_API_KEY = " your-key "
Never commit that file; it is already ignored. Avoid putting an owner-funded key on a public deployment unless you also add authentication and spending controls.
python -m pip install -r requirements-dev.txt
ruff check .
python -m unittest discover -s tests -v
GitHub Actions runs linting, the complete test suite, and bytecode compilation on every push and pull request.
Does my data leave my machine?
No. Cleaning, schema detection, every chart, every evidence card, and every Ask ADA answer are computed locally with pandas. If you opt into the AI layer, only column schema and computed evidence are sent — never rows or cell values.
Do I need an OpenAI API key?
No. ADA is a complete analyst without one. A key only adds the query-planner fallback for unusual questions and the strategic narrative.
What formats can I analyze?
CSV (comma, semicolon, or tab delimited), XLSX, and XLSM — including picking a specific worksheet from a multi-sheet workbook.
How is this different from pasting a CSV into a chatbot?
A chatbot gives you fluent prose you cannot audit and your rows become part of a prompt. ADA turns questions into explicit query plans, executes them with pandas on your machine, and prints the calculation under every answer.
Can I self-host it?
Yes — it is a standard Streamlit app. pip install -r requirements.txt && streamlit run app.py , or deploy it to any host that runs Python.
Contributions are welcome, especially around new deterministic metrics, question shapes for Ask ADA, schema-detection fixtures, chart accessibility, file formats, and adversarial test datasets. Start with the good first issues , read CONTRIBUTING.md , choose an item from the roadmap , or open a focused proposal.
Good contributions make an insight more accurate, more explainable, or easier for a non-technical user to act on. Every new recommendation should include a test and the calculation that supports it.
If ADA is useful to you, a star helps other operators find it.
Deploy app.py on Streamlit. The repository includes its app theme, dependency manifest, server upload limit, and headless configuration. Add OPENAI_API_KEY through Streamlit's secret manager only if the optional strategy layer should be available.
I tried to automate a data analyst
There was an error while loa

[truncated]
