---
source: "https://github.com/superlinked/synty"
hn_url: "https://news.ycombinator.com/item?id=48737949"
title: "Show HN: Distributed LLM tracing and GH PR/issue linking [Apache 2.0]"
article_title: "GitHub - superlinked/synty: Local-first memory for you and your coding agents. Every session and PR, searchable; the raw data is yours to analyze. · GitHub"
author: "supo"
captured_at: "2026-06-30T19:33:14Z"
capture_tool: "hn-digest"
hn_id: 48737949
score: 1
comments: 0
posted_at: "2026-06-30T19:22:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Distributed LLM tracing and GH PR/issue linking [Apache 2.0]

- HN: [48737949](https://news.ycombinator.com/item?id=48737949)
- Source: [github.com](https://github.com/superlinked/synty)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T19:22:36Z

## Translation

タイトル: HN を表示: 分散 LLM トレースと GH PR/問題リンク [Apache 2.0]
記事のタイトル: GitHub - スーパーリンク/synty: あなたとコーディング エージェントのためのローカルファーストのメモリ。すべてのセッションと PR が検索可能。生データはあなたが分析できます。 · GitHub
説明: あなたとあなたのコーディングエージェントのためのローカルファーストメモリ。すべてのセッションと PR が検索可能。生データはあなたが分析できます。 - スーパーリンク/シンティ
HN テキスト: コーディング エージェントがどこで苦労しているのか、どのツールがトークンを消費するのか、そしてそれが GitHub の PR や問題にどのように関係しているのかを確認してください。すべての開発マシン、VM、サンドボックスから S3 にデータをストリーミングするツールをオープンソース化しました。そこから独自の分析を実行することも、バンドルされている TUI を使用することもできます。これを構築した理由: 私は 10 人程度の AGI 専門エンジニアのチームを運営しており、自分たちが行っている作業が、計画の際に検討しているマイルストーン/プロジェクトとどのように一致しているかを追跡するのに苦労しています。 synty は、エージェントや人間の注目を集めているものの、計画には考慮されていない作業領域を特定するのに役立ちます。私たちはこのデータを使用して、エージェントの摩擦を取り除くためのリファクタリングや文書化の取り組みを正当化します。

記事本文:
GitHub - スーパーリンク/synty: あなたとあなたのコーディングエージェントのためのローカルファーストメモリ。すべてのセッションと PR が検索可能。生データはあなたが分析できます。 · GitHub
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
スーパーリンクされた
/
シン

タイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
210 コミット 210 コミット .github/ workflows .github/ workflows docs docs evals evals scripts scripts src src .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントはすべてのセッションをゼロから開始します。あなたもそうです。
synty は、すべてのコーディング エージェント セッション (Claude Code、Codex、Cowork) を受動的に記録します。
と GitHub アクティビティを 1 つのローカルの検索可能なメモリに保存します。
一緒に働くエージェント。
vhs を使用して docs/*.tape からレンダリングされたデモ。
再発見ではなく思い出す。 「認証フローに触れた人はいますか?」と尋ねてください。 、または単に
synty 関連を実行し、重要なセッションと PR を取得します。で閲覧します
TUI。エージェントは CLI および MCP 経由で同じ内容を読み取ります。
あなたのデータはあなたのもののままです。すべてがディスク上でオープンです。生のイベントは JSONL として、
SQLite インデックス、すべてのコマンドの --json。チャート エージェントの摩擦、微調整
モデルを作成し、独自のダッシュボードを構築します。視聴者には何も閉じ込められません。
地元第一主義。バイナリは 1 つで、API キーはなく、マシンからは何も残りません。さえも
1 行のサマリーは小規模なローカル モデルで実行されます。
カール -fsSL https://raw.githubusercontent.com/superlinked/synty/main/install.sh |しー
1 回貼り付けるとバイナリがインストールされ、ログイン時間トラッカーがオンになり、
視聴者。その後、2 つのコマンドで毎日のループをカバーします。
synty tui # ワークメモリを参照: トピック、セッション、検索、統計
synty 関連 # 今やっていることの以前の作業を表示します (クエリなし)
トラッカーはログイン時に実行されるため、メモリは自動的に構築され続けます。アップデート

どれでも
synty のアップグレードに伴う時間。代わりにソースからビルドしますか? 「ビルド」を参照してください。
すべての読み取りコマンドは、Markdown を標準出力に出力するか、バージョン管理されたエンベロープの場合は --json を出力します。
結果は、ドリル可能な ID を含むランク付けされた Markdown カードです ([a1b2c3d4] セッション、
repo#123 PRs/issues、[72a778f8] トピック) synty show をフィードします:
## レート制限ミドルウェア
1. [24.3] **pull_request api#1487** — ゲートウェイにトークンバケット リミッターを追加します
統合されました · https://github.com/acme/api/pull/1487
2. [21.8] _user_prompt · API · a3f1c2d9_
テナントごとのリミッターの状態をポッド間で共有するにはどうすればよいですか? Redisに落ち着きました…
3. [19.0] **api#1502 を発行** — 検索エンドポイントのバースト負荷で 429 秒
開く · https://github.com/acme/api/issues/1502
synty はキャプチャ レイヤーであり、壁に囲まれた庭園ではありません。記録されるものはすべてそこに収まります
~/.synty の下にプレーン ファイルがあるため、それを直接ビルドできます。
生のイベント (真実のソース): corpus/local/ の下に追加専用の JSONL
(および共有バケット内の events/<stream>/)。
ドキュメント : corpus/docs.jsonl 、メタを含む 1 行に 1 つのオブジェクト
(source 、 kind 、 repo 、 author 、 session_id 、 ts )。
メタデータ : Index/ の下にある SQLite データベース。任意の SQLite ツールでクエリ可能。
synty 統計 --json | jq ' .data.weeks[] | {week: .start, tok_in, tok_out} ' # 週次トークン傾向
synty ツール Bash --json | jq ' {calls: .data.calls、errors: .data.errors} ' # エージェントがスタックする場所
jq -r ' .meta.kind ' ~ /.synty/corpus/docs.jsonl |並べ替え | uniq -c # 生ファイルから直接
synty はすでに作品をクラスタリングし、ソース間でリンクしているため (PR、問題、
セッションなど）、分析はログの山ではなく構造化データから始まります。
個人的には、「バケット」はローカル ディレクトリです。チーム用、または自分だけのラップトップ用
デスクトップで synty を 1 つの共有 S3/GCS バケットにポイントすると、1 つのメモリを構築できます。
synty init gs://my-team
そのバケツだけが

共有インフラストラクチャ: ビルドサーバーなし、調整なし
サービス。すべてのマシンのトラッカーはイベントをプッシュします。ビューアを開いた人が
インデックスを付けて残りの部分を公開します。 1 台のトークン化されたマシンが GitHub をスクレイピング
みんな。クラウド バケットには --features s3 / gcs が必要です。
注意: バケットを持つすべてのメンバーは、バケット内のすべてのセッションを読み取ることができます。シンティ
信頼性の高いグループ向けに構築されています。読者ごとの編集はありません。それなら
適合しない、ローカルに留まる、またはバケットの範囲をすでにそれぞれを参照している人に限定する
他人の作品。
取得は遅いインタラクションです。 pylate-rs は小さな ColBERT モデルを実行します
(ModernBERT、32 M パラメータ) 各ドキュメントをトークンごとに 1 つのベクトルとしてエンコードします。
next-plaid は、SQLite メタデータ フィルター上で MaxSim を使用してクエリをスコアリングします。それ
トークンごとのスコアリングにより、まれな用語または特定の用語 (関数名、ファイル パス) が可能になります
1 つのドキュメント ベクトルに平均化されるのではなく、独自の信号を伝送します。
単一ベクトルの埋め込みの場合と同様です。
概要とトピック名は、小規模なローカル モデル (Qwen3-0.6B) から取得されます。
リモート API ではなく、抽出フォールバックを備えた CPU です。
出来事は真実の源です。インデックスとメタデータが導出されます
いつでも再構築可能で、バケットを通じて共有可能です。
アーキテクチャと理論的根拠は docs/design.md にあります。の
実データ上の検証は evals/ にあります。
カーゴビルド --release # プレーン CPU、ポータブル、依存関係が軽い (出荷されたビルド)
Apple Silicon では、GPU エンコード用に --features metal を追加します (約 5.7 倍高速)。
Accelerate (macOS) と mkl (Linux) は、CPU-BLAS の代替品です。埋め込み
初回使用時にモデル (~127 MB) をダウンロードします。サマライザー (~1.2 GB) 初回
何でも要約します。エアギャップ設定とステージごとのパイプラインについては、を参照してください。
docs/design.md および CONTRIBUTING 。
リリースのカットはメンテナのタスクです。「 CONTRIBUTING 」を参照してください。
あなたとコーディングエージェントのためのローカルファーストメモリ

。すべてのセッションと PR が検索可能。生データはあなたが分析できます。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0
最新の
2026 年 6 月 27 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-first memory for you and your coding agents. Every session and PR, searchable; the raw data is yours to analyze. - superlinked/synty

See where coding agents struggle, what tools eat tokens and how it all links to GitHub PRs & issues. I open-sourced a tool that streams data from all your dev machines, VMs & sandboxes into S3. You can run your own analysis from there or use the bundled TUI. Why I built this: I run a team of ~10 AGI-pilled engineers and struggle to keep track of how the work we do aligns to the milestones/projects we look in planning. synty helps me identify areas of work that are getting a lot of agent/human attention but they are not accounted for in planning and we use this data to justify re-factoring/documentation efforts to remove friction for agents

GitHub - superlinked/synty: Local-first memory for you and your coding agents. Every session and PR, searchable; the raw data is yours to analyze. · GitHub
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
superlinked
/
synty
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
210 Commits 210 Commits .github/ workflows .github/ workflows docs docs evals evals scripts scripts src src .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Your coding agents start every session from zero. So do you.
synty passively records every coding-agent session (Claude Code, Codex, Cowork)
and your GitHub activity into one local, searchable memory, for you and the
agents you work with.
Demos rendered from docs/*.tape with vhs .
Recall, not re-discovery. Ask "has anyone touched the auth flow?" , or just
run synty related , and get the sessions and PRs that matter. You browse in
the TUI; your agents read the same over the CLI and MCP.
Your data stays yours. Everything is open on disk: raw events as JSONL, a
SQLite index, --json on every command. Chart agent friction, fine-tune a
model, build your own dashboards. Nothing is trapped in a viewer.
Local-first. One binary, no API keys, nothing leaves your machine. Even the
one-line summaries run on a small local model.
curl -fsSL https://raw.githubusercontent.com/superlinked/synty/main/install.sh | sh
One paste installs the binary, turns on the login-time tracker, and opens the
viewer. After that, two commands cover the daily loop:
synty tui # browse your work memory: topics, sessions, search, stats
synty related # surface prior work for whatever you're doing now (no query)
The tracker runs at login, so the memory keeps building on its own. Update any
time with synty upgrade . Building from source instead? See Build .
Every read command prints Markdown to stdout, or --json for a versioned envelope.
A result is a ranked Markdown card with ids you can drill ( [a1b2c3d4] sessions,
repo#123 PRs/issues, [72a778f8] topics) that feed synty show :
## rate limiting middleware
1. [24.3] **pull_request api#1487** — Add a token-bucket limiter to the gateway
merged · https://github.com/acme/api/pull/1487
2. [21.8] _user_prompt · api · a3f1c2d9_
how do we share the per-tenant limiter's state across pods? settled on Redis…
3. [19.0] **issue api#1502** — 429s under burst load on the search endpoint
open · https://github.com/acme/api/issues/1502
synty is a capture layer, not a walled garden. Everything it records sits in
plain files under ~/.synty , so you can build on it directly:
Raw events (the source of truth): append-only JSONL under corpus/local/
(and events/<stream>/ in a shared bucket).
Documents : corpus/docs.jsonl , one object per line with meta
( source , kind , repo , author , session_id , ts ).
Metadata : a SQLite database under index/ , queryable with any SQLite tool.
synty stats --json | jq ' .data.weeks[] | {week: .start, tok_in, tok_out} ' # weekly token trend
synty tool Bash --json | jq ' {calls: .data.calls, errors: .data.errors} ' # where agents get stuck
jq -r ' .meta.kind ' ~ /.synty/corpus/docs.jsonl | sort | uniq -c # straight from the raw file
Because synty already clusters the work and links across sources (PRs, issues,
sessions), your analysis starts from structured data, not a heap of logs.
Solo, the "bucket" is a local directory. For a team, or just your own laptop and
desktop, point synty at one shared S3/GCS bucket and you build one memory:
synty init gs://my-team
That bucket is the only shared infrastructure: no build server, no coordination
service. Every machine's tracker pushes events; whoever opens a viewer builds the
index and publishes it for the rest; one tokened machine scrapes GitHub for
everyone. Cloud buckets need --features s3 / gcs .
Heads up: every member with the bucket can read every session in it. synty
is built for high-trust groups; there's no per-reader redaction. If that
doesn't fit, stay local, or scope the bucket to people who already see each
other's work.
Retrieval is late interaction. pylate-rs runs a small ColBERT model
(ModernBERT, 32 M params) that encodes each document as one vector per token;
next-plaid scores queries with MaxSim over a SQLite metadata filter. That
per-token scoring lets a rare or specific term (a function name, a file path)
carry its own signal, instead of being averaged into one document vector the
way single-vector embeddings do.
Summaries and topic names come from a small local model (Qwen3-0.6B) on
your CPU, never a remote API, with an extractive fallback.
Events are the source of truth. The index and metadata are derived
projections, rebuildable any time and shareable through a bucket.
Architecture and rationale live in docs/design.md ; the
on-real-data validation lives in evals/ .
cargo build --release # plain CPU, portable, dependency-light (the shipped build)
On Apple Silicon, add --features metal for GPU encode (~5.7× faster);
accelerate (macOS) and mkl (Linux) are CPU-BLAS alternatives. The embedding
model (~127 MB) downloads on first use; the summarizer (~1.2 GB) the first time
anything summarizes. For an air-gapped setup and the per-stage pipeline, see
docs/design.md and CONTRIBUTING .
Cutting a release is a maintainer task: see CONTRIBUTING .
Local-first memory for you and your coding agents. Every session and PR, searchable; the raw data is yours to analyze.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.1.0
Latest
Jun 27, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
