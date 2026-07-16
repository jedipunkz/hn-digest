---
source: "https://github.com/d4n-larsson/aegisdb"
hn_url: "https://news.ycombinator.com/item?id=48941533"
title: "AegisDB – self-hosted memory for AI agents, in one C binary"
article_title: "GitHub - d4n-larsson/aegisdb: A fast, single-binary C database that gives AI agents durable, searchable long-term memory — with a Claude Code integration. · GitHub"
author: "d4n-larsson"
captured_at: "2026-07-16T23:50:30Z"
capture_tool: "hn-digest"
hn_id: 48941533
score: 2
comments: 0
posted_at: "2026-07-16T23:15:56Z"
tags:
  - hacker-news
  - translated
---

# AegisDB – self-hosted memory for AI agents, in one C binary

- HN: [48941533](https://news.ycombinator.com/item?id=48941533)
- Source: [github.com](https://github.com/d4n-larsson/aegisdb)
- Score: 2
- Comments: 0
- Posted: 2026-07-16T23:15:56Z

## Translation

タイトル: AegisDB – AI エージェント用のセルフホスト型メモリ、1 つの C バイナリ
記事のタイトル: GitHub - d4n-larsson/aegisdb: クロード コードの統合により、AI エージェントに耐久性と検索可能な長期メモリを提供する、高速なシングルバイナリ C データベース。 · GitHub
説明: クロード コードの統合により、AI エージェントに耐久性と検索可能な長期記憶を提供する、高速なシングルバイナリ C データベース。 - GitHub - d4n-larsson/aegisdb: クロード コードの統合により、AI エージェントに耐久性と検索可能な長期メモリを提供する、高速なシングルバイナリ C データベース。

記事本文:
GitHub - d4n-larsson/aegisdb: クロード コードの統合により、AI エージェントに耐久性と検索可能な長期メモリを提供する、高速なシングルバイナリ C データベース。 · GitHub
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
d4n-ラーソン
/
イージスデータベース
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
260 コミット 260 コミット .claude .claude .github/ workflows .github/ workflows bench bench docs docs include/ aegisdb include/ aegisdb 統合 統合スクリプト スクリプト サイト サイト src src テスト テスト サードパーティ サードパーティ .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt Dockerfile Dockerfile LAUNCH.md LAUNCH.md LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用のセルフホスト メモリ。 1 つの小さな C バイナリ — マルチテナント、
暗号化、バックアップ、リードレプリカ、および 1 つのコマンド Prometheus + Grafana を使用
積み重ねる。エージェントの記憶はボックスに残ります。 SaaS には何も送信されません。
AI エージェントはセッション間ですべてを忘れます。 AegisDB は耐久性を提供します。
検索可能な長期記憶 — エピソードの歴史、ベクトルを含む意味論的事実
検索と揮発性の作業メモリ - 非常に単純な JSON-over-TCP の背後にある
第一級クロードコードを使用したプロトコル
統合。これは、依存関係のない単一のバイナリを自分で実行するものです。つまり、データ、
あなたのボックスに第三者が関与することはありません。
サーバーを実行します。クローンやツールチェーンは使用しません (GHCR 上に事前構築されたマルチ アーキテクチャ イメージ)。
docker run -d --name aegisdb -p 9470:9470 -v aegis-data:/data \
ghcr.io/d4n-larsson/aegisdb:最新
それに話しかけてください — 同じバイナリがクライアントでもあります。
docker exec aegisdb aegisdb クライアント ping
docker exec aegisdb aegisdb client put --type semantic --tags user " ダーク モードを好む "
docker exec aegisdb aegisdb クライアント検索 --tags user --top-k 5
可観測性スタック全体 (サーバー + Prometheus + 事前に構築された Grafana) が必要です
ダッシュボード) を 1 つのカンマで囲みます

で？このリポジトリのクローンを作成し、次の操作を行います。
docker compose --profile モニタリングアップ # ダッシュボード (http://127.0.0.1:3000)
クロード コードに永続メモリを与えるのは、ワンライナーです (サーバーを使用すると
running): uvx --from aegisdb-mcp aegisdb-init — を参照してください。
クロードコードのメモリとして使用します。
自己ホスト型かつプライベート。エージェントの記憶がインフラストラクチャから離れることはありません
— SaaS、トークンごとの課金、データ共有はありません。保存時に 1 つで暗号化する
旗。
1 つのバイナリで、依存関係はありません。 C で書かれています。唯一のベンダーコードは cJSON です
そして暗号通貨。 JVM も Python ランタイムも、子守りする外部データベースもありません。
チーム向けに構築されています。マルチテナント認証 (名前空間ごと、スコープ指定されたトークン)、テナントごと
クォータ + レート制限、オンライン バックアップ、リード レプリカ、およびターンキー
プロメテウス/グラファナの可観測性。
クロードコードネイティブ。クロードが記憶できるように MCP サーバー + フックを出荷します
セッション — 1 つのコマンドでインストール可能。
生産志向。破損に強い追加専用ログ、クラッシュ回復、
文書化されたセキュリティレビュー、および ASan/UBSan/TSan を実行する CI と継続的な
毛羽立ちます。
耐久性のあるエピソード メモリ - マジック + CRC32 フレーミングによる追加専用ログ、破損に強いリカバリ、レガシー ログ移行
セマンティック ファクト — 更新可能なレコード (最新バージョンが優先)
作業メモリ — TTL とプロモーションを備えた揮発性のセッションごとのリング バッファー
取得 — IDによる検索、時間範囲検索、タグ検索（ all / any ）、
セマンティック (埋め込み) 検索は、コサイン類似度によって重み付けされてランク付けされます。
重要性 × 自信。同じフィルター上でカウントと統合 (重複除去)
セマンティック スケール — 小さい間は正確なコサイン。 --ann-threshold を過ぎた HNSW
サブリニア近似トップ K のグラフ。書き込みパスから構築され、シャーディングされたもの
ビルドは並列化 ( --ann-shard-target )、オプションで int8 量子化されます
関係 - レコード間の有向エッジ、グラフ走査、

そして
エージェントの名前空間の分離
マルチテナント認証 — オプションのベアラー トークン (常時チェック、ping 免除)、それぞれが名前空間 + スコープ (ro / rw /admin) にバインドされているため、1 つのサーバーで多くのテナントを安全に分離できます。
テナントごとの制限 — オプションのストレージ クォータ (レコード/バイト) と名前空間ごとのリクエスト レート制限により、1 人のチーム メンバーの暴走エージェントがディスクをいっぱいにしたり、共有サーバーを独占したりすることがなくなります。
保存時の暗号化 - ログ + チェックポイント上のオプションの XChaCha20-Poly1305 (ベンダー提供、暗号依存なし)。 --encryption-key-file 経由でオプトインし、オフライン マイグレーターと暗号化されたバックアップ/レプリカを使用します
可観測性 — 統計操作とドロップイン Prometheus エクスポータ + Grafana ダッシュボード ( docker compose --profile Monitoring up )
操作 - オンライン スナップショット/バックアップとリード レプリカの復元
同時実行性 — シャード化されたpoll()イベントループスレッド( --io-threads )。選択可能な fsync 耐久性 (同期 / バッチ / 間隔)
GCC 11 以降または Clang 14 以降を搭載した Linux (プライマリ ターゲット)
次のいずれか: CMake 3.20+ または GNU Make
Python 3.8+ (オプション、以下のサンプルクライアントの場合)
cmake -B build -DCMAKE_BUILD_TYPE=リリース
cmake --build ビルド
ctest --test-dir build --output-on-failure # 単体テスト スイートを実行します
Make あり (CMake は必要ありません)
make # ビルド build/aegisdb
make test # C 単体テストをビルドして実行します
統合を行う # ワイヤープロトコルコントラクトテスト (サーバーを起動)
チェック # ユニット + 統合を行う
きれいにする
サーバー バイナリは build/aegisdb で生成されます。
事前に構築されたマルチアーキテクチャ イメージ ( linux/amd64 、 linux/arm64 ) が GitHub に公開されます
メインおよびすべてのリリース タグへのプッシュごとのコンテナー レジストリ — クローンまたは
必要なツールチェーン:
docker run -p 9470:9470 -v aegis-data:/data ghcr.io/d4n-larsson/aegisdb:latest
# またはリリースを固定します: ghcr.io/d4n-larsson/aegisdb:0.1.0
代わりに自分で構築するには、多段階の

Dockerfile (Debian-slim) のコンパイル
サーバーにインストールされ、最小限のランタイム イメージが同梱されます。データは次の名前付きボリュームに保持されます。
/データ .
# Docker Compose でビルドして実行する
docker compose up --build # localhost:9470 で動作します
# またはイメージを直接ビルドして実行します
docker build -t aegisdb 。
docker run -p 127.0.0.1:9470:9470 -v aegis-data:/data aegisdb
Compose はオプションの .env ファイルによって構成されます。テンプレートをコピーして編集します。
cp .env.example .env # 次に、ポート、耐久性、テナント制限などを調整します。
docker 構成 --build
すべての設定にはデフォルトがあるため、.env はオプションです。共通フラグを公開します
名前付き変数 ( AEGIS_PORT 、 AEGIS_EMBEDDING_DIM 、 AEGIS_DURABILITY 、
AEGIS_TENANT_MAX_RECORDS , …) その他の場合は AEGIS_EXTRA_ARGS を加えます
( --auth-token-file 、 --io-threads 、ANN チューニングなど)。参照
完全なリストについては、.env.example を参照してください。
ビルドをスキップするには、docker-compose.yml で公開されたイメージを指定します。
ビルド: 。画像付き: ghcr.io/d4n-larsson/aegisdb:latest 。
イメージには、バイナリの組み込み --health-check を使用する HEALTHCHECK が同梱されています。
プローブ (イメージに追加のツールはありません) なので、 docker ps と Compose
depend_on: 条件: service_healthy は実際のサーバーの稼働状況を反映します。
コンテナは特権のないユーザーとして実行されます。サーバーは 0.0.0.0:9470 で待機します
コンテナ内にありますが、Compose はホストのループバックでそのポートを公開します
( 127.0.0.1 ) はデフォルトのみ - ワイヤプロトコルが認証されていないため、
平文なので、保護するまではオフボックスでアクセスできないようにする必要があります。
意図的に公開するには、AEGIS_BIND=0.0.0.0 (または特定のホスト IP) を設定します。
.env を指定し、最初に認証を有効にします。トークン ファイルを /data にマウントして追加します。
--auth-token-file /data/tokens.txt (Compose の AEGIS_EXTRA_ARGS へ; を参照)
認証)。認証があっても、トークンは平文で送信されます。
TLS を終了します

非ループバックのエクスポージャに対する信頼されたプロキシ。他をオーバーライドする
フラグを実行コマンドに追加してフラグを設定します。 docker run aegisdb --embedding-dim 1024 、または (Compose を使用して) .env 経由で実行します。
./build/aegisdb --data-dir ./data --port 9470
予想される起動出力:
2026-06-28 12:00:00.000 INFO [aegisdb] AegisDB 0.1.0 開始中 (ログレベル: info)
2026-06-28 12:00:00.000 警告 [aegisdb] 認証トークンが構成されていません。 ...
2026-06-28 12:00:00.000 情報 [aegisdb] リカバリが完了しました: N レコードがロードされました
2026-06-28 12:00:00.000 情報 [aegisdb] 0.0.0.0:9470 でリッスンしています
2026-06-28 12:00:00.000 情報 [aegisdb] データ ディレクトリ: ./data
ログは <timestamp> <LEVEL> [aegisdb] <message> として stderr に保存されます。を制御する
--log-level error|warn|info|debug (デフォルトの info ) または
AEGISDB_LOG_LEVEL 環境変数 - フラグが優先されます。で
debug を実行すると、サーバーは受け入れられたすべての接続とディスパッチされた操作をログに記録します。
WARN 行は、サーバーが何もせずに起動された場合にのみ表示されます。
--auth-token / --auth-token-file (「認証」を参照)。
同じバイナリがクライアントでもあります。 nc や手書きの JSON はありません。
aegisdb クライアントの ping
aegisdb クライアント put --type semantic --tags ユーザー「ダーク モードを好む」
aegisdb クライアントが 1 を取得
aegisdb クライアント検索 --tags ユーザー --top-k 5
aegisdb クライアントの統計
ホスト、ポート、トークンのデフォルトは $AEGIS_HOST / $AEGIS_PORT / $AEGIS_TOKEN
( 127.0.0.1 / 9470 / none) または --host / --port / --token 。終了コードは
OK 応答の場合は 0 なので、スクリプトはきれいに実行されます。 Docker 内:
docker exec aegisdb aegisdb クライアント統計 。
テナント トークンを作成するために、gen-token は準備完了のトークン ファイル行 (ハッシュ化された) を出力します。
そしてワンタイム平文トークン:
$ aegisdb gen-token --namespace acme --scope rw
sha256$… acme rw # --auth-token-file に貼り付けます
token: 9f3c… # クライアントに与える (AEGIS_TOKEN);回復不可能
設定

フラグ
旗
デフォルト
説明
--data-dir <パス>
./データ
永続ディレクトリ
--ポート <n>
9470
TCPリッスンポート
--フェーズ<1-4>
4
最も有効な機能フェーズ (ゲート操作)
--io-スレッド<n>
2× CPU (8 ～ 64)
ディスパッチ並列処理のための、poll() イベント ループ スレッド (同時接続を制限しません)。エイリアス: --workers
--max-payload <バイト>
1048576
最大データサイズ (1 MiB)
--embedding-dim <n>
384
予想される埋め込みベクトルの長さ
--ann-しきい値 <n>
10000
セマンティック検索が正確なスキャンから HNSW グラフに切り替わる前のライブ ベクトル
--ann-ef-search <n>
ニューサウスウェールズ州のデフォルト
HNSW クエリ ビーム幅 (リコール/レイテンシ ノブ)
--ann-shard-target <n>
25000
HNSW シャードごとのターゲット ベクター。グラフは ~ count/n 個のシャード (CPU によって制限される) に分割されるため、ビルドは並列化されます。
--ann-quantize
オフ
HNSW ベクトルを int8 として保存 (メモリが約 4 分の 1 で、リコール コストが小さい)
--耐久性 <モード>
間隔
同期 (書き込みごとの fsync)、バッチ (--fsync-batch レコードごと)、または間隔 (--fsync-interval-ms ごと)
--fsync-batch <n>
1000
バッチモードでの fsync 呼び出し間の記録
--fsync-interval-ms <n>
1000
インターバル モードでのフラッシュ ケイデンス (~1 秒のメンテナンス ティックで下限)
--チェックポイント秒 <n>
60
インデックス チェックポイントの頻度により、リカバリでは末尾のみが再生されます。 0 は無効になります
--compact-sec <n>
300
ログ圧縮チェックの頻度。コンパック

[切り捨てられた]

## Original Extract

A fast, single-binary C database that gives AI agents durable, searchable long-term memory — with a Claude Code integration. - GitHub - d4n-larsson/aegisdb: A fast, single-binary C database that gives AI agents durable, searchable long-term memory — with a Claude Code integration.

GitHub - d4n-larsson/aegisdb: A fast, single-binary C database that gives AI agents durable, searchable long-term memory — with a Claude Code integration. · GitHub
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
d4n-larsson
/
aegisdb
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
260 Commits 260 Commits .claude .claude .github/ workflows .github/ workflows bench bench docs docs include/ aegisdb include/ aegisdb integrations integrations scripts scripts site site src src tests tests third_party third_party .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt Dockerfile Dockerfile LAUNCH.md LAUNCH.md LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml View all files Repository files navigation
Self-hosted memory for your AI agents. One small C binary — multi-tenant,
encrypted, with backups, read replicas, and a one-command Prometheus + Grafana
stack. Your agents' memory stays on your box; nothing ships to a SaaS.
AI agents forget everything between sessions. AegisDB gives them durable,
searchable long-term memory — episodic history, semantic facts with vector
search, and volatile working memory — behind a dead-simple JSON-over-TCP
protocol, with a first-class Claude Code
integration. It's a single dependency-free binary you run yourself: your data,
your box, no third party in the loop.
Run the server — no clone, no toolchain (prebuilt multi-arch image on GHCR):
docker run -d --name aegisdb -p 9470:9470 -v aegis-data:/data \
ghcr.io/d4n-larsson/aegisdb:latest
Talk to it — the same binary is also the client:
docker exec aegisdb aegisdb client ping
docker exec aegisdb aegisdb client put --type semantic --tags user " prefers dark mode "
docker exec aegisdb aegisdb client search --tags user --top-k 5
Want the whole observability stack (server + Prometheus + a pre-built Grafana
dashboard) in one command? Clone this repo and:
docker compose --profile monitoring up # dashboard on http://127.0.0.1:3000
Giving Claude Code a persistent memory is a one-liner (with a server
running): uvx --from aegisdb-mcp aegisdb-init — see
Use as Claude Code memory .
Self-hosted & private. Your agents' memory never leaves your infrastructure
— no SaaS, no per-token billing, no data-sharing. Encrypt it at rest with one
flag.
One binary, no dependencies. Written in C; the only vendored code is cJSON
and the crypto. No JVM, no Python runtime, no external database to babysit.
Built for teams. Multi-tenant auth (per-namespace, scoped tokens), per-tenant
quotas + rate limits, online backups, read replicas, and turnkey
Prometheus/Grafana observability.
Claude Code native. Ships an MCP server + hooks so Claude remembers across
sessions — installable with a single command.
Production-minded. Corruption-resilient append-only log, crash recovery,
a documented security review, and CI that runs ASan/UBSan/TSan plus continuous
fuzzing.
Durable episodic memory — append-only log with magic + CRC32 framing, corruption-resilient recovery, and legacy-log migration
Semantic facts — updateable records (latest version wins)
Working memory — volatile per-session ring buffer with TTL and promotion
Retrieval — lookup by ID, time-range search, tag search ( all / any ),
semantic (embedding) search ranked by cosine similarity weighted by
importance × confidence; count and consolidate (dedup) over the same filters
Semantic scale — exact cosine while small; past --ann-threshold an HNSW
graph for sublinear approximate top-K, built off the write path and sharded so
the build parallelizes ( --ann-shard-target ), optionally int8-quantized
Relationships — directed edges between records, graph traversal, and
agent-namespace isolation
Multi-tenant auth — optional bearer tokens (constant-time check; ping exempt), each bound to a namespace + scope ( ro / rw /admin) so one server safely isolates many tenants
Per-tenant limits — optional storage quotas (records/bytes) and a request rate limit per namespace, so one team member's runaway agent can't fill the disk or monopolize the shared server
Encryption at rest — optional XChaCha20-Poly1305 (vendored, no crypto dependency) over the log + checkpoints; opt-in via --encryption-key-file , with an offline migrator and encrypted backups/replicas
Observability — stats op plus a drop-in Prometheus exporter + Grafana dashboard ( docker compose --profile monitoring up )
Operations — online snapshot /restore backups and read replicas
Concurrency — sharded poll() event-loop threads ( --io-threads ); selectable fsync durability ( sync / batch / interval )
Linux (primary target) with GCC 11+ or Clang 14+
One of: CMake 3.20+ or GNU Make
Python 3.8+ (optional, for the example client below)
cmake -B build -DCMAKE_BUILD_TYPE=Release
cmake --build build
ctest --test-dir build --output-on-failure # runs the unit test suite
With Make (no CMake required)
make # builds build/aegisdb
make test # builds and runs the C unit tests
make integration # wire-protocol contract tests (launches the server)
make check # unit + integration
make clean
The server binary is produced at build/aegisdb .
Prebuilt multi-arch images ( linux/amd64 , linux/arm64 ) are published to GitHub
Container Registry on every push to main and every release tag — no clone or
toolchain needed:
docker run -p 9470:9470 -v aegis-data:/data ghcr.io/d4n-larsson/aegisdb:latest
# or pin a release: ghcr.io/d4n-larsson/aegisdb:0.1.0
To build it yourself instead, a multi-stage Dockerfile (Debian-slim) compiles
the server and ships a minimal runtime image. Data persists in a named volume at
/data .
# Build and run with Docker Compose
docker compose up --build # serves on localhost:9470
# Or build and run the image directly
docker build -t aegisdb .
docker run -p 127.0.0.1:9470:9470 -v aegis-data:/data aegisdb
Compose is configured by an optional .env file — copy the template and edit:
cp .env.example .env # then tweak port, durability, tenant limits, …
docker compose up --build
Every setting has a default, so .env is optional. It exposes the common flags
as named vars ( AEGIS_PORT , AEGIS_EMBEDDING_DIM , AEGIS_DURABILITY ,
AEGIS_TENANT_MAX_RECORDS , …) plus AEGIS_EXTRA_ARGS for anything else
( --auth-token-file , --io-threads , ANN tuning, …). See
.env.example for the full list.
To skip building, point docker-compose.yml at the published image: replace
build: . with image: ghcr.io/d4n-larsson/aegisdb:latest .
The image ships a HEALTHCHECK that uses the binary's built-in --health-check
probe (no extra tooling in the image), so docker ps and Compose
depends_on: condition: service_healthy reflect real server liveness.
The container runs as an unprivileged user. The server listens on 0.0.0.0:9470
inside the container, but Compose publishes that port on the host's loopback
( 127.0.0.1 ) only by default — because the wire protocol is unauthenticated and
plaintext out of the box, it must not be reachable off-box until you secure it.
To expose it deliberately, set AEGIS_BIND=0.0.0.0 (or a specific host IP) in
.env , and first enable authentication: mount a token file into /data and add
--auth-token-file /data/tokens.txt (to AEGIS_EXTRA_ARGS under Compose; see
Authentication ). Even with auth, tokens travel in plaintext,
so terminate TLS at a trusted proxy for any non-loopback exposure. Override other
flags by appending them to the run command, e.g. docker run aegisdb --embedding-dim 1024 , or (with Compose) via .env .
./build/aegisdb --data-dir ./data --port 9470
Expected startup output:
2026-06-28 12:00:00.000 INFO [aegisdb] AegisDB 0.1.0 starting (log level: info)
2026-06-28 12:00:00.000 WARN [aegisdb] no auth tokens configured; ...
2026-06-28 12:00:00.000 INFO [aegisdb] recovery complete: N records loaded
2026-06-28 12:00:00.000 INFO [aegisdb] listening on 0.0.0.0:9470
2026-06-28 12:00:00.000 INFO [aegisdb] data directory: ./data
Logs go to stderr as <timestamp> <LEVEL> [aegisdb] <message> . Control the
verbosity with --log-level error|warn|info|debug (default info ) or the
AEGISDB_LOG_LEVEL environment variable — the flag takes precedence. At
debug , the server logs every accepted connection and dispatched operation.
The WARN line appears only when the server is started without
--auth-token / --auth-token-file (see Authentication ).
The same binary is also a client — no nc , no hand-written JSON:
aegisdb client ping
aegisdb client put --type semantic --tags user " prefers dark mode "
aegisdb client get 1
aegisdb client search --tags user --top-k 5
aegisdb client stats
Host, port, and token default to $AEGIS_HOST / $AEGIS_PORT / $AEGIS_TOKEN
( 127.0.0.1 / 9470 / none) or --host / --port / --token . The exit code is
0 on an ok response, so it scripts cleanly. Inside Docker:
docker exec aegisdb aegisdb client stats .
To create a tenant token, gen-token prints a ready token-file line (hashed)
and the one-time plaintext token:
$ aegisdb gen-token --namespace acme --scope rw
sha256$… acme rw # paste into your --auth-token-file
token: 9f3c… # give to the client (AEGIS_TOKEN); not recoverable
Configuration flags
Flag
Default
Description
--data-dir <path>
./data
Persistence directory
--port <n>
9470
TCP listen port
--phase <1-4>
4
Highest enabled feature phase (gates operations)
--io-threads <n>
2× CPUs (8–64)
poll() event-loop threads for dispatch parallelism (does not cap concurrent connections). Alias: --workers
--max-payload <bytes>
1048576
Max data size (1 MiB)
--embedding-dim <n>
384
Expected embedding vector length
--ann-threshold <n>
10000
Live vectors before semantic search switches from exact scan to the HNSW graph
--ann-ef-search <n>
HNSW default
HNSW query beam width (recall/latency knob)
--ann-shard-target <n>
25000
Target vectors per HNSW shard; the graph splits into ~ count/n shards (capped by CPUs) so the build parallelizes
--ann-quantize
off
Store HNSW vectors as int8 (~4× less memory, small recall cost)
--durability <mode>
interval
sync (fsync per write), batch (per --fsync-batch records), or interval (per --fsync-interval-ms )
--fsync-batch <n>
1000
Records between fsync calls in batch mode
--fsync-interval-ms <n>
1000
Flush cadence in interval mode (floored at the ~1s maintenance tick)
--checkpoint-sec <n>
60
Index checkpoint cadence so recovery replays only the tail; 0 disables
--compact-sec <n>
300
Log-compaction check cadence; compac

[truncated]
