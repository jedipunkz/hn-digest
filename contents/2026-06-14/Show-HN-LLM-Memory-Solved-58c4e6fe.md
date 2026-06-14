---
source: "https://github.com/gary23w/neuron-db"
hn_url: "https://news.ycombinator.com/item?id=48527346"
title: "Show HN: LLM Memory Solved?"
article_title: "GitHub - gary23w/neuron-db: An associative memory you can run anywhere. Write facts in plain language, recall them by meaning. No tables, no schema, no embeddings, no model required. · GitHub"
author: "gary23w"
captured_at: "2026-06-14T15:38:00Z"
capture_tool: "hn-digest"
hn_id: 48527346
score: 1
comments: 0
posted_at: "2026-06-14T14:07:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LLM Memory Solved?

- HN: [48527346](https://news.ycombinator.com/item?id=48527346)
- Source: [github.com](https://github.com/gary23w/neuron-db)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T14:07:28Z

## Translation

タイトル: HN を表示: LLM メモリは解決されましたか?
記事タイトル: GitHub - gary23w/neuron-db: どこでも実行できる連想メモリ。事実を平易な言葉で書き、意味を理解して思い出しましょう。テーブル、スキーマ、埋め込み、モデルは必要ありません。 · GitHub
説明: どこでも実行できる連想メモリ。事実を平易な言葉で書き、意味を理解して思い出しましょう。テーブル、スキーマ、埋め込み、モデルは必要ありません。 - gary23w/ニューロン-db

記事本文:
GitHub - gary23w/neuron-db: どこでも実行できる連想メモリ。事実を平易な言葉で書き、意味を理解して思い出しましょう。テーブル、スキーマ、埋め込み、モデルは必要ありません。 · GitHub
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
ゲイリー23w
/
ニューロンデータベース
公共
通知
君はね

サインインして通知設定を変更できない
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
56 コミット 56 コミット .claude .claude docs docs 例 例 Rust/ ニューロンコア Rust/ ニューロンコア ワーカー ワーカー .dockerignore .dockerignore .gitignore .gitignore .hf_model_card.md .hf_model_card.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md build.sh build.sh docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
どこでも実行できる連想メモリと、定額料金の長期記憶。
LLM 。事実を平易な言葉で書き、意味によって思い出し、ニューロンを結びつける
追加のモデルコストなしで任意の深さのチェーンを作成できます。テーブルもスキーマも埋め込みもなし
モデルが必要です。コアは依存関係のない純粋な Rust であり、WebAssembly にコンパイルされます。
耐久性のあるストレージ、暗号化、HTTP サーバー、および MCP サーバーはオプトイン機能です。
./build.sh
ニューロン --db app.db 変えてください 「私の計画はプロです」
ニューロン --db app.db で「どのような計画を立てていますか?」 ' # -> プロ
▶ ブラウザーで、本物の Rust コアであるシナプスの起動を 3D で観察してください。
LLM メモリ: 定額のコストで無限のニューロンをリンクします。
LLM のコンテキスト ウィンドウは小さいです。 neuron-db は、その外部に存在するメモリです。あ
関係質問 - 「Aurora の所有者のマネージャーのタイムゾーン」 - 通常
モデルに事実を思い出させ、待って、次の事実を思い出させ、待って… N ホップ = N+1 モデル呼び出し。
remember_chain は次のように折りたたまれます: モデルは 1 つのパスを送信し、シナプスはそのパスを歩きます。
チェーン全体のサーバー側、各ホップのマイクロ秒のリコール。深度はマイクロ秒単位で支払われます。
モデルターンではありません。
現在ほとんどの LLM が使用しているメモリに対してライブで測定 (ダンプされたすべてのファクトのマークダウン ファイル)
毎ターンコンテキストに)、 gpt-4o-mini 、 700-fac

ユーザー:
マークダウンダンプは毎ターンメモリ全体を再注入し、最終的にはウィンドウをオーバーランさせます。
neuron-db は、思い出したものだけを注入します — 定額コスト、上限なし、マッチングまたはビートなし
精度。完全な数値: docs/COMPARISON.md · リコールの起動速度:
docs/SYNAPSE.md 。
1列に並べて取り付けます。 neuron-mcp はネイティブの std 専用 stdio MCP サーバーです - 任意の MCP を指します
クライアント (Claude Desktop/Code、Cursor) がバイナリにあると、モデルは記憶/呼び出しを取得します。
/recall_chain をツールとして使用します。ノードもPythonもHTTPプロセスもありません。参照
docs/MEMORY_HARNESS.md および example/mcp_chat/ 。
カーゴ ビルド --release --features mcp --bin ニューロン-mcp
それは何ですか
ファクトは文です (「API キーは zeta-9931」)。ニューロンデータベースは驚くべき言葉を保存します
取得可能な値として残りをキューとしてインデックス付けします。スコープはファクトの名前付きバッグです
( user:42 )、データベースはスコープのファイルです。物事を述べることで挿入し、読む
質問する — 検索は連想的 (キューの重複) であるため、列や
SQLを書きます。完全なモデルとすべての操作: docs/API.md 。
ニューロン_コア :: db :: NeuronDB を使用します。
let db = NeuronDB :: open ( "app.db" , 500 ) ;
データベース。観察 ( "ユーザー:42" , "計画はプロです" ) ;
データベース。 get ( "ユーザー:42" , "どのプランですか?" ) ; // 一部("プロ")
データベース。 remember ( "user:42" , Some ( "plan" ) ) ; // 部分文字列ごとに削除
階層
Neuron — メモリ内連想ストア (デフォルト、標準のみ)。マイクロ秒単位で呼び出します。
PlasticNeuron — 想起は適応します: 使用時の強度、不使用時の減衰、ヘビアン リンク、および
神経伝達物質による拡散活性化の想起。
NeuronRouter — 多くの小さなニューロンにシャーディングし、クエリを扇状に広げます ( --features none)。
NeuronDB — 1 つの SQLite ファイル内のスコープの耐久性のあるデータベース ( --features sqlite )。
SecureNeuronDB — AES-256-GCM 値、スコープごとのシークレットは保存されません ( --features secure )。
HTT

P サーバー + バイナリの提供 — スコープごとに 1 つのエンドポイント ( --features サーバー )。
ニューロン-mcp — 標準出力 MCP サーバー。LLM はニューロン データベースをメモリとしてマウントします ( --features mcp )。
小さい。ファクトの検索状態はステムとスカラーであり、密なベクトルではありません - 約 48
シリアル化されたファクトあたりのバイト数、1536 次元の浮動小数点数ベクトルよりも GiB あたり約 130 倍のファクト数
店。 docs/STORAGE.md を参照してください。
高速で依存性がありません。マイクロ秒のリコール、GPU なし、モデルなし。デフォルトのビルドが実行される
1 MB WebAssembly ワーカー内。
適応型。プラスチック層は、O(1) スカラー更新での使用から学習します。再埋め込みはなく、
インデックスの再作成はありません。
トレードポイント: スカラー優先、語彙の想起であり、学習された意味上の類似性ではありません。それは橋渡しします
形態 (所有者 / 所有者 / 所有者) および厳選された同義語オントロジー (レポート ↔ マネージャー、
↔ city に住んでいます) は無料ですが、オープンボキャブラリーの言い換え (「オンラインにアクセスするために使用するもの」→「Wi-Fi パスワード」) には、やはり埋め込み層が必要です。引き換えにあなたは得ます
マイクロ秒のリコール、ベクトル ストアの密度の約 130 倍、ホット パス上にモデルなし。
./build.sh # sqlite + セキュア + サーバー
カーゴ ビルド --release --features " sqlite セキュア サーバー "
カーゴインストール --path Rust/neuron-core --features " sqlite セキュア サーバー "
デフォルトのビルドは依存関係がゼロで、 wasm32-unknown-unknown をターゲットとします。ネイティブ層は
オプトイン機能により、wasm ビルドには一切触れなくなります。サービス (および Docker) として実行します。
docs/DEPLOY.md 。
埋め込み SQLite にはログインがありません。ファイルシステムのアクセス許可、HTTP サーバーのアクセス許可によってアクセスを制御します。
NEURON_DB_KEY ベアラー トークン、または SecureNeuronDB によるスコープごとの暗号化。詳細:
セキュリティ.md 。
ストア層とサービス層は Rust の標準です (rust/neuron-core/)。ニシキヘビ
リファレンス実装 — ゲイリーニューロン皮質ブリッジとトレーニングツールを含む —
Legacy-Python に保存されます

支店。
実行可能なコードと統合ガイドは、examples/ — クイックスタート、
チャットボットのメモリ ループ、ユーザーごとのプロファイル、シャーディング、暗号化されたシークレット、HTTP クライアント
(curl/browser/Node/Python)、およびニューロン データベースをチャットボットに接続するためのガイド
または既存の API 。
docs/SYNAPSE.md — LLM のリコールの起動速度 (測定値)
docs/COMPARISON.md — マルチホップ + ニューロンデータベースとマークダウンダンプメモリの比較
docs/MEMORY_HARNESS.md — LLM メモリとしてマウントします。 MCP サーバーとツール
docs/API.md — データモデルとすべての操作 (ライブラリ/CLI/HTTP)
docs/DEPLOY.md — ビルド、インストール、Docker、環境、バックアップ
docs/STORAGE.md — ストレージ密度とベクトル データベース
SECURITY.md — 暗号化とアクセス モデル
MITライセンス取得済み。作者: gary23w
どこでも実行できる連想メモリ。事実を平易な言葉で書き、意味を理解して思い出しましょう。テーブル、スキーマ、埋め込み、モデルは必要ありません。
gary23w.github.io/neuron-db/
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
3
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An associative memory you can run anywhere. Write facts in plain language, recall them by meaning. No tables, no schema, no embeddings, no model required. - gary23w/neuron-db

GitHub - gary23w/neuron-db: An associative memory you can run anywhere. Write facts in plain language, recall them by meaning. No tables, no schema, no embeddings, no model required. · GitHub
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
gary23w
/
neuron-db
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
56 Commits 56 Commits .claude .claude docs docs examples examples rust/ neuron-core rust/ neuron-core worker worker .dockerignore .dockerignore .gitignore .gitignore .hf_model_card.md .hf_model_card.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md build.sh build.sh docker-compose.yml docker-compose.yml View all files Repository files navigation
An associative memory you can run anywhere — and the flat-cost long-term memory for an
LLM . Write facts in plain language, recall them by meaning, and link neurons across
arbitrarily deep chains at no extra model cost . No tables, no schema, no embeddings, no
model required. The core is pure Rust with zero dependencies and compiles to WebAssembly;
durable storage, encryption, an HTTP server, and an MCP server are opt-in features.
./build.sh
neuron --db app.db turn me ' my plan is pro '
neuron --db app.db get me ' what plan am i on? ' # -> pro
▶ Watch the synapse fire in 3D — the real Rust core, in your browser.
LLM memory: link infinite neurons, at flat cost
An LLM's context window is small; neuron-db is the memory that lives outside it. A
relational question — "the timezone of the manager of the owner of Aurora" — normally
makes the model recall a fact, wait, recall the next, wait… N hops = N+1 model calls.
recall_chain collapses that: the model sends one path, and the synapse walks the
whole chain server-side, each hop a microsecond recall. Depth is paid in microseconds,
not model turns.
Measured live against the memory most LLMs use today (a markdown file of all facts dumped
into context every turn), gpt-4o-mini , a 700-fact user:
The markdown-dump reinjects the whole memory every turn and eventually overruns the window;
neuron-db injects only what it recalled — flat cost, no ceiling, matching or beating
accuracy. Full numbers: docs/COMPARISON.md · how fast recall fires:
docs/SYNAPSE.md .
Mount it in one line. neuron-mcp is a native std-only stdio MCP server — point any MCP
client (Claude Desktop/Code, Cursor) at the binary and your model gets remember / recall
/ recall_chain as tools. No Node, no Python, no HTTP process. See
docs/MEMORY_HARNESS.md and examples/mcp_chat/ .
cargo build --release --features mcp --bin neuron-mcp
What it is
A fact is a sentence ( "the api key is zeta-9931" ); neuron-db keeps the surprising word
as the retrievable value and indexes the rest as cues. A scope is a named bag of facts
( user:42 ), and a database is a file of scopes. You insert by stating things and read by
asking questions — retrieval is associative (cue overlap), so you never declare a column or
write SQL. Full model and every operation: docs/API.md .
use neuron_core :: db :: NeuronDB ;
let db = NeuronDB :: open ( "app.db" , 500 ) ;
db . observe ( "user:42" , "the plan is pro" ) ;
db . get ( "user:42" , "what plan?" ) ; // Some("pro")
db . forget ( "user:42" , Some ( "plan" ) ) ; // delete by substring
Tiers
Neuron — in-memory associative store (default, std-only). Recall in microseconds.
PlasticNeuron — recall adapts: strength on use, decay on disuse, Hebbian links, and
a neurotransmitter-style spreading-activation recall.
NeuronRouter — shard across many small neurons and fan a query out ( --features none).
NeuronDB — durable database of scopes in one SQLite file ( --features sqlite ).
SecureNeuronDB — AES-256-GCM values, per-scope secret never stored ( --features secure ).
HTTP server + serve binary — one endpoint per scope ( --features server ).
neuron-mcp — stdio MCP server so any LLM mounts neuron-db as memory ( --features mcp ).
Tiny. A fact's retrieval state is stems and scalars, not a dense vector — about 48
bytes/fact serialized, roughly 130× more facts per GiB than a 1536-dim float vector
store. See docs/STORAGE.md .
Fast and dependency-free. Microsecond recall, no GPU, no model. The default build runs
in a 1 MB WebAssembly worker.
Adaptive. The plastic tier learns from use with O(1) scalar updates — no re-embedding,
no re-indexing.
The trade: it's scalar-first, lexical recall — not learned semantic similarity. It bridges
morphology ( owner / owned / owns ) and a curated synonym ontology ( reports to ↔ manager ,
lives in ↔ city ) for free, but open-vocabulary paraphrase ( "the thing I use to get online" → "wifi password" ) would still want an embedding tier. In exchange you get
microsecond recall, ~130× the density of a vector store, and no model on the hot path.
./build.sh # sqlite + secure + server
cargo build --release --features " sqlite secure server "
cargo install --path rust/neuron-core --features " sqlite secure server "
Default build is zero-dependency and targets wasm32-unknown-unknown ; the native tiers are
opt-in features so they never touch the wasm build. Running it as a service (and Docker):
docs/DEPLOY.md .
Embedded SQLite has no login — control access by filesystem permissions, the HTTP server's
NEURON_DB_KEY bearer token, or per-scope encryption with SecureNeuronDB . Details:
SECURITY.md .
The store and service tiers are canonical in Rust ( rust/neuron-core/ ). A Python
reference implementation — including the gary-neuron cortex bridge and training tooling —
is preserved on the legacy-python branch.
Runnable code and integration guides are in examples/ — quickstart, a
chatbot-memory loop, per-user profiles, sharding, encrypted secrets, HTTP clients
(curl/browser/Node/Python), and guides for wiring neuron-db into a chatbot
or an existing API .
docs/SYNAPSE.md — how fast recall fires for an LLM, measured
docs/COMPARISON.md — multi-hop + neuron-db vs the markdown-dump memory
docs/MEMORY_HARNESS.md — mount as LLM memory; the MCP server & tools
docs/API.md — data model and every operation (library / CLI / HTTP)
docs/DEPLOY.md — build, install, Docker, env, backups
docs/STORAGE.md — storage density vs vector databases
SECURITY.md — encryption and access model
MIT licensed. Author: gary23w.
An associative memory you can run anywhere. Write facts in plain language, recall them by meaning. No tables, no schema, no embeddings, no model required.
gary23w.github.io/neuron-db/
Resources
There was an error while loading. Please reload this page .
3
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
