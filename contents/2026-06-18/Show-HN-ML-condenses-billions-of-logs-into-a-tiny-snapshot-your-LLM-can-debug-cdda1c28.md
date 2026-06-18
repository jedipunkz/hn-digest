---
source: "https://github.com/Rocketgraph/rocketgraph"
hn_url: "https://news.ycombinator.com/item?id=48578324"
title: "Show HN: ML condenses billions of logs into a tiny snapshot your LLM can debug"
article_title: "GitHub - Rocketgraph/rocketgraph: Agent layer for observability · GitHub"
author: "kvaranasi_"
captured_at: "2026-06-18T01:02:27Z"
capture_tool: "hn-digest"
hn_id: 48578324
score: 8
comments: 2
posted_at: "2026-06-17T23:14:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ML condenses billions of logs into a tiny snapshot your LLM can debug

- HN: [48578324](https://news.ycombinator.com/item?id=48578324)
- Source: [github.com](https://github.com/Rocketgraph/rocketgraph)
- Score: 8
- Comments: 2
- Posted: 2026-06-17T23:14:37Z

## Translation

タイトル: HN を表示: ML は、LLM がデバッグできる小さなスナップショットに数十億のログを凝縮します
記事タイトル: GitHub - Rocketgraph/rocketgraph: 可観測性のためのエージェント層 · GitHub
説明: 可観測性のためのエージェント層。 GitHub でアカウントを作成して、Rocketgraph/ロケットグラフの開発に貢献します。
HN テキスト: こんにちは HN、私は Kaushik です。Rocketgraph を構築しました。他の空間が AI の波に追いついている一方で、人間が書いたコードのログを分析するために使用しているのと同じツールやダッシュボードを使用する可観測性の空間はまだ遅れていると私は考えています。しかし現在では、コードの作成とデバッグは AI によって行われるため、オブザーバー自体が AI である場合に可観測性を実現する方法を再考する必要があります。私が遭遇する問題は、アラートが発生したときに手動で Grafana ダッシュボードを確認し、LogQL クエリを作成する必要があることです。これは grep によく似ています。しかし、通常、スキーマの不一致、DB 接続の問題、または数百万のログ行に埋もれている見たことのないログ行が原因で、本番環境が中断します。さらに悪いことに、アラートは決して発火せず、いつ grep すればよいのかわかりません。Rocketgraph はそれを修正します。ログをフィンガープリントしてパターンに変換し、ML を使用して、頻度、テキストの類似性、その他のベクトルなどの特徴によって異常スコアを付けます。したがって、通常、これにより、100 万件のログが、異常スコアと特徴ベクトルを含む 200 ～ 300 のパターンに凝縮され、消防ホース全体を送信することなく、LLM が簡単に分析できるようになります。これは特定の時点で実行されるため、ログに基づくオンラインの異常検出のようなものです。一部の企業はメトリクスに対して異常検出を行っていますが、これはログに対して行われます。この分野の他のアプローチでは、既存の Grafana ダッシュボードの上に AI をボルトで固定しますが、これは追加の手順を手動で grep するのと同じことです。ローカルでホストするためのセットアップ例を確認してください。

そしてログファイルに対して実行します。皆さんの意見を聞かせてください!

記事本文:
GitHub - Rocketgraph/rocketgraph: 可観測性のためのエージェント層 · GitHub
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
ロケットグラフ
/
ロケットグラフ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル C に移動

ode 「その他のアクション」メニューを開く フォルダーとファイル
17 コミット 17 コミット example-setups example-setups イメージ イメージ ml ml パッケージ/ otel-node パッケージ/ otel-node .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE.txt LICENSE.txt README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すでに実行している可観測性スタックの隣に組み込まれる自己ホスト型のログ クラスタリングとストリーミング異常検出。
ここには何が入っているのか
•
クイックスタート
•
例
•
ウェブサイト
•
コミュニティ
監視ツールでは、何を検索したかがわかります。現在何が異常なのかを伝えることはめったにありません。
Rocketgraph は、Datadog、New Relic、Loki、CloudWatch、Sentry、ClickHouse など、すでにお金を払っているものの隣にあり、ログのウィンドウを取得し、構造テンプレートをマイニングし、異常なものにフラグを立てます。完全にネットワーク内で実行されます。ログが VPC から離れることはありません。料金を支払う SaaS 層はありません。
コンポーネント
何をするのか
🧠 MLエンジン
クラスターは構造テンプレートにログインし、異常を検出します。既存のログ ソースから直接取得します。並列取り込みパイプラインはありません。
⚡ @rgraph/otel-node
OpenTelemetry を使用して任意の Node.js サービスを約 90 秒で自動計測する AI エージェント。
90秒以内にお試しください
git clone https://github.com/Rocketgraph/rocketgraph
cdロケットグラフ/ml
cp .env.example .env # 手持ちのソースを入力します
docker compose up --build # → http://localhost:9020
すでに使用しているソースを指定します。
カール「 http://localhost:9020/clusters?source=loki&window=1h 」
または、認証情報を完全にスキップして、ログ ファイルをダウンロードして実行します。 Datadog (CSV/JSON)、kubectl logs > app.log 、または任意の生のログからエクスポートし、ドロップインしてローカルで分析します。
カール -XPOST ' http://localhost:9020/clusters/train?source=file ' # ファイル

_PATH=/data/app.log
1 つのコマンドでログ ファイルを作成するクイックスタートを参照してください。
これでインストールは完了です。プロビジョニングするスキーマ、作成するアカウント、ホスト上のエージェントは必要ありません。
👉 詳細: ML エンジンの場合は ml/README.md、OTel エージェントの場合は package/otel-node
仕組み (30 秒バージョン)
3 つの決定論的アルゴリズムを順に実行 - LLM なし、幻覚なし、完全に再現可能:
Drain3 は生の丸太ラインから構造テンプレートをマイニングします。
Isolation Forest は、サービスごとにテンプレートにスコアを付けて、異常なものを明らかにします。
Half-Space-Trees は、トレーニングされたモデルに対して真新しいログをリアルタイムでスコア付けします。
実際の運用バーストでは、200 万ログ → 58 テンプレート → 9 異常、実測 90 秒、単一コンテナに対してテストしました。詳細については、ml/README.md を参照してください。
ログ ファイルをローカルで分析する —analyze.py
Rocketgraph の動作を確認する最速の方法: ログ ファイルを ./logs/ にドロップし、ログ ファイルを実行します。
コマンドを実行し、異常のフラグが立てられたクラスター テーブルを取得します。アカウントも API もありません
キーを押しても、マシンからは何も離れません。オプションのクロード優先順位付けのために --ai を追加します。
上 — エンジン自体は意図的に LLM フリーのままであり、再現可能です。モデル
決定論的クラスターのみを説明します。
cd 例-セットアップ/ログファイル-クイックスタート
docker compose up --build -d # ML エンジン (http://localhost:9020)
python gen_sample_log.py # または: cp ~/Downloads/whatever.log ./logs/file.log
pip install request # --ai を使用する場合は anthropic も要求します
python Analyzer.py # すべてのクラスターのテーブル
python Analyzer.py --anomalies-only # フラグが立てられたもののみ
Python Analyzer.py --ai # テーブル + AI トリアージ
python Analyzer.py mylogs.log --ai # 特定のファイル
analyze.py はファイルを自動検出し、エンジンにそのファイルを指定し、クラスターをプルします。
そしてそれらを印刷します。約 15,000 の生の線が約 11 の構造テンプレートに分解されます。の
新しい「データベース フェイルオーバー」テンプレート — 8 行、これまで見たことのない、エラー

r
レベル — 異常としてフラグが立てられて戻ってきます。ルールもラベルも書かれていない:
15188 ログ → 11 クラスター (3 異常)
Anom サービス ログの深さのテンプレート
--------------------------------------
*payment-svc 8 3 データベース フェイルオーバー: レプリカ <*> は ... 後にプライマリに昇格しました
* auth-svc 1573 2 セッション <NUM> のトークンが更新されました
payment-svc 1686 $<FLOAT> に対する請求 <NUM> が承認されました
...
表の見方: ANOM はクラスターに Isolation Forest フラグを付けます。ログ
そのテンプレートに折りたたまれた生の行の数です。 DEPTH は分離です
異常なクラスターの深さ (低い = より異常)。テンプレートは
構造パターン Drain3 が採掘されました。フラグが設定されたフェールオーバー クラスターはまれであり、新しいものです。
それがまさにそれを表面化させたものです。
--ai を使用すると、同じクラスターが SRE スタイルのトリアージのためにクロードに渡されます。
可能性の高いインシデント、ランク付けされた根本原因仮説、具体的な次のステップ - 根拠のあるもの
上記のクラスター内のみ。完全なウォークスルー
ログファイルのクイックスタート。
example-setups/ には、ポイントできるリファレンス アプリも含まれています
otel-node でパイプライン全体の動作を確認します — サービスを計測し、出荷します
OTLP をシンクに接続し、Rocketgraph クラスターを監視してログにフラグを立てます。
ロードマップにはさらに多くの例 (Fastify、NestJS、Next.js) が含まれています。広報担当者は歓迎します。
ステータス
プラットフォーム
✅ サポートされています
ログ ファイル ( .log / .json / .csv ) · OpenTelemetry · Loki · New Relic · Datadog · CloudWatch · Sentry · ClickHouse
🛣️ロードマップ
Splunk · Elastic / OpenSearch · Azure Monitor · GCP Cloud Logging
コミュニティ
💬 Discord — サポートとデザインのディスカッション
🐛 GitHub の問題 — バグと機能リクエスト
PR の方は大歓迎です。現在最も影響力のある貢献:
新しい ML コネクタ (Splunk、OpenSearch、Azure Monitor、GCP Cloud Logging)
@rgraph/otel-node での追加のフレームワーク サポート (Fastify、NestJS、Remix、Bun-native servi)

セス)
example-setups/ にあるその他のエンドツーエンドのリファレンス アプリ
詳細なドキュメントについては、ml/README.md および package/otel-node を参照してください。
自己ホスト型。オープンソース。すでに実行しているものの隣にドロップします。
ロケットグラフ.app
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
9
フォーク
レポートリポジトリ
リリース
1
🚀 Rocketgraph v0.1.0 - 数十億のログを小さなスナップショットに圧縮して異常を検出します
最新の
2026 年 5 月 27 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent layer for observability. Contribute to Rocketgraph/rocketgraph development by creating an account on GitHub.

Hi HN, I'm Kaushik, and I built Rocketgraph. I believe that while other spaces have caught up to the AI wave, the observability space is still lagging behind, using the same tools and dashboards that we use to analyse logs from human-written code. But now the code is written and debugged by AI, so we need to rethink how we do observability where the observer itself is an AI. The problem that I run into is when an alert fires, I have to manually check the Grafana dashboards and write LogQL queries, which is pretty much like greping. But production usually breaks due to a schema mismatch, or a DB connection issue or a log line that I haven't seen before that's buried under millions of log lines. Much worse, the alert never fires, and I don't know when to grep Rocketgraph fixes that. It turns your logs into patterns by fingerprinting them, then uses ML to anomaly score them by features like frequency, text similarity and other vectors. So, usually this condenses a million logs into 200-300 patterns with anomaly scores and feature vectors that your LLM can easily analyse without sending the entire firehose. This runs at specific points in time, so it's like an online anomaly detection based on logs. Some companies do anomaly detection on metrics, but this is done for logs. Other approaches in this space bolt an AI on top of existing Grafana dashboards, but it's the same thing as manually greping with extra steps. Please check out the example setups to host it locally and run it on your log files. Let me know what you guys think!

GitHub - Rocketgraph/rocketgraph: Agent layer for observability · GitHub
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
Rocketgraph
/
rocketgraph
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits example-setups example-setups images images ml ml packages/ otel-node packages/ otel-node .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE.txt LICENSE.txt README.md README.md View all files Repository files navigation
Self-hosted log clustering and streaming anomaly detection that drops in next to the observability stack you already run.
What's in here
•
Quick start
•
Examples
•
Website
•
Community
Your monitoring tool tells you what you searched for . It rarely tells you what's unusual right now .
Rocketgraph sits next to whatever you already pay for — Datadog, New Relic, Loki, CloudWatch, Sentry, ClickHouse — pulls a window of logs, mines structural templates, and flags the anomalous ones. It runs entirely inside your network. Your logs never leave your VPC. There's no SaaS tier to pay for.
Component
What it does
🧠 ML engine
Clusters logs into structural templates and detects anomalies. Pulls directly from your existing log source — no parallel ingest pipeline.
⚡ @rgraph/otel-node
AI agent that auto-instruments any Node.js service with OpenTelemetry in ~90 seconds.
Try it in 90 seconds
git clone https://github.com/Rocketgraph/rocketgraph
cd rocketgraph/ml
cp .env.example .env # fill in whichever sources you have
docker compose up --build # → http://localhost:9020
Point it at any source you already use:
curl ' http://localhost:9020/clusters?source=loki&window=1h '
Or skip the credentials entirely — download a log file and run it. Export from Datadog (CSV/JSON), kubectl logs > app.log , or any raw log, drop it in, and analyse it locally:
curl -XPOST ' http://localhost:9020/clusters/train?source=file ' # FILE_PATH=/data/app.log
See the one-command log-file quickstart .
That's the whole install. No schemas to provision, no accounts to create, no agents on hosts.
👉 Deep dive: ml/README.md for the ML engine · packages/otel-node for the OTel agent
How it works (30-second version)
Three deterministic algorithms in sequence — no LLM, no hallucination, fully reproducible:
Drain3 mines structural templates from raw log lines.
Isolation Forest scores templates per service to surface the unusual ones.
Half-Space-Trees scores brand-new logs against the trained model in real time.
On a real production burst we test against: 2M logs → 58 templates → 9 anomalies, 90 seconds wall-clock, single container. Full details in ml/README.md .
Analyse a log file locally — analyze.py
The fastest way to see Rocketgraph work: drop a log file in ./logs/ , run one
command, and get a cluster table with the anomalies flagged. No accounts, no API
keys, nothing leaves your machine. Add --ai for an optional Claude triage on
top — the engine itself stays deliberately LLM-free and reproducible; the model
only explains the deterministic clusters.
cd example-setups/logfile-quickstart
docker compose up --build -d # ML engine on http://localhost:9020
python gen_sample_log.py # or: cp ~/Downloads/whatever.log ./logs/file.log
pip install requests # anthropic too, if you'll use --ai
python analyze.py # table of all clusters
python analyze.py --anomalies-only # just the flagged ones
python analyze.py --ai # table + AI triage
python analyze.py mylogs.log --ai # a specific file
analyze.py auto-detects the file, points the engine at it, pulls the clusters,
and prints them. ~15,000 raw lines collapse to ~11 structural templates; the
brand-new "database failover" template — 8 lines, never seen before, error
level — comes back flagged as an anomaly. No rules written, no labels:
15188 logs → 11 clusters (3 anomalous)
ANOM SERVICE LOGS DEPTH TEMPLATE
----------------------------------------
* payment-svc 8 3 Database failover: replica <*> promoted to primary after ...
* auth-svc 1573 2 Token refreshed for session <NUM>
payment-svc 1686 Charge <NUM> authorized for $<FLOAT>
...
Reading the table: ANOM marks the clusters Isolation Forest flagged; LOGS
is how many raw lines collapsed into that template; DEPTH is the isolation
depth on anomalous clusters (lower = more anomalous); TEMPLATE is the
structural pattern Drain3 mined. The flagged failover cluster is rare and new,
which is exactly what surfaces it.
With --ai , the same clusters are handed to Claude for an SRE-style triage —
likely incident, ranked root-cause hypotheses, and concrete next steps — grounded
only in the clusters above. Full walkthrough in the
log-file quickstart .
example-setups/ also contains reference apps you can point
otel-node at to see the whole pipeline working — instrument the service, ship
OTLP into your sink, then watch Rocketgraph cluster and flag the logs.
More examples (Fastify, NestJS, Next.js) are on the roadmap — PRs welcome.
Status
Platforms
✅ Supported
Log file ( .log / .json / .csv ) · OpenTelemetry · Loki · New Relic · Datadog · CloudWatch · Sentry · ClickHouse
🛣️ Roadmap
Splunk · Elastic / OpenSearch · Azure Monitor · GCP Cloud Logging
Community
💬 Discord — support and design discussions
🐛 GitHub Issues — bugs and feature requests
PRs welcome. The most impactful contributions right now:
New ML connectors (Splunk, OpenSearch, Azure Monitor, GCP Cloud Logging)
Additional framework support in @rgraph/otel-node (Fastify, NestJS, Remix, Bun-native services)
More end-to-end reference apps under example-setups/
See ml/README.md and packages/otel-node for the deep-dive docs.
Self-hosted. Open source. Drops in next to what you already run.
rocketgraph.app
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
9
forks
Report repository
Releases
1
🚀 Rocketgraph v0.1.0 - Compress billions of logs into tiny snapshots to detect anomalies
Latest
May 27, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
