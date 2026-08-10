---
source: "https://github.com/yaminbakoh4-dot/NexusMem"
hn_url: "https://news.ycombinator.com/item?id=49240020"
title: "Show HN: NexusMem – Local context memory engine for AI coding agents"
article_title: "GitHub - yaminbakoh4-dot/NexusMem · GitHub"
author: "yamin_bakoh"
captured_at: "2026-08-10T07:10:15Z"
capture_tool: "hn-digest"
hn_id: 49240020
score: 2
comments: 0
posted_at: "2026-08-10T06:34:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: NexusMem – Local context memory engine for AI coding agents

- HN: [49240020](https://news.ycombinator.com/item?id=49240020)
- Source: [github.com](https://github.com/yaminbakoh4-dot/NexusMem)
- Score: 2
- Comments: 0
- Posted: 2026-08-10T06:34:23Z

## Translation

タイトル: Show HN: NexusMem – AI コーディング エージェント用のローカル コンテキスト メモリ エンジン
記事タイトル: GitHub - yaminbakoh4-dot/NexusMem · GitHub
説明: GitHub でアカウントを作成して、yaminbakoh4-dot/NexusMem の開発に貢献します。

記事本文:
GitHub - yaminbakoh4-dot/NexusMem · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
やみんばこ4-dot
/
ネクサスメモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コミット 16 コミット docs docs src src testing testing .gitignore .gitignore ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts すべてのファイルを表示 リポジト

ry ファイルのナビゲーション
AI コーディング エージェント (クロード コード、カーソル、MCP ベースのエージェント) 用のローカル ファースト永続メモリ エンジン。
AI コーディング アシスタントはセッションが終了するとコンテキストを忘れ、リポジトリ全体を次のように再アップロードします。
すべてのリクエストのコンテキストは遅くて高価です。 NexusMem はローカル マシン イベントを記録します — git 履歴、
シェル コマンド、ドキュメント、会話トランスクリプトをディスク上の SQLite データベースに保存し、のみを返します。
厳密なトークンバジェット内で関連するコンテキストスライス。
すべてのデータはマシン上のローカルに残ります。クラウドへの依存関係、アカウント、テレメトリはありません。
100% ローカルファースト: SQLite データベースはリポジトリ内の .nexusmem/ に保存されます。
sqlite-vec と FTS5 。完全にオフラインで動作します。
Kind-Agnostic Core : すべてのソースが単一の MemoryNode スキーマに正規化され、git が可能になります
コミット、シェルコマンド、ドキュメントがスコアリングされ、平等にランク付けされます。
ハイブリッド検索 (BM25 + Vector) : SQLite FTS5 (BM25) による正確なキーワード マッチングと、
Reciprocal Rank Fusion (RRF) を使用したセマンティック ベクトル検索 (ローカル Ollama モデルを介した sqlite-vec)。
RRF はランク位置のみに基づいて融合し、生のスコアには決して融合しません。これにより、BM25 を安全に組み合わせることができます。
無関係なスケールのベクトル距離によるコスト。 Ollama が有効な場合、BM25 のみに正常に機能が低下します。
オフライン。
ランク付けされた予算付き検索 : を使用して候補者をスコアリングします。
スコア = 関連性 × 信号^a × 最新性 ^b に基づいて、呼び出し元が指定したトークン バジェットにノードをパックします。
各因子は [0, 1] ではなく [floor, 1] に切り詰められるため、単一の低い因子をゼロにすることはできません
強い試合。指数 a と b は調整されずに導出されます。関連性のみが重要です。
クエリから導出された係数なので、クエリに依存しない各事前確率は、その全体を上限とする累乗になります。
2倍の関連性ギャップを覆すときの範囲（span^exponent = 2、 a ≈ 0.431 、 b ≈ 0.576 を与える）。

MCP サーバー ネイティブ : search_memory 、 sync_project 、および get_status をモデル コンテキストとして公開します
Claude Desktop、Cursor、Windsurf 用の stdio 経由のプロトコル (MCP) ツール。
git / シェル / ドキュメント / トランスクリプト
│
▼ コレクター/ 1 つの MemoryNode 形状に正規化する
│
▼ストア/ SQLite (FTS5 + sqlite-vec)
│
▼取得/RRFヒューズ -> ランク -> トークンバジェットへのパック
主要なサブシステム
Git Collector : コミット、差分統計、名前変更、および従来のコミット信号を取り込みます
ストリーム反復子を介して段階的に。同期カーソルは、HEAD の祖先として検証されてから、
信頼できるため、黙ってコミットをスキップするのではなく、リベースまたは修正によってウォーク範囲が広がります。
Shell Collector : デフォルトの履歴ファイル ( PSReadLine 、 .bash_history 、
.zsh_history )。オプションの PowerShell プロファイル フックは、正確なタイムスタンプを含むようにキャプチャをアップグレードします。
作業ディレクトリ、および終了コード (失敗したコマンドがより高度な構造シグナルを受け取る場所)。
Docs Collector : git によって追跡される Markdown ドキュメント ( .md ファイル) のインデックスを作成します。行末は、
Windows での CRLF 分割エラーを防ぐために、チャンク化する前に LF に正規化されます。範囲指定された枝刈り
見出しの名前が変更または削除されたときに、プロジェクトおよび正確なソースごとに範囲を指定して、孤立したセクションを削除します。
したがって、git、シェル、または会話ノードには影響しません。
会話コレクター (オプトイン): AI アシスタントのトランスクリプトにインデックスを付け、事前に秘密を編集します。
ディスクに書き込んでいます。返信は、次のように保存されるのではなく、見出しと太字の境界で分割されて保存されます。
全体の交換。
ノード ID はコンテンツアドレス指定 ( sha256(projectId + kind + NaturalKey) ) なので、同期を 2 回実行します。
重複を生成できず、カーソルが失われた場合でも取り込みは正しいままになります。プロジェクトのアイデンティティは、
正規化されたオリジン URL が存在する場合はそこから派生し、絶対パスにフォールバックするため、2 つ
同じリポジトリのクローンは 1 つのメモリを共有します

y 名前空間。
nodes_fts はトリガーによって設定され、自動的に一貫性を保ちます。 nodes_vec はコンピューティングではありません
埋め込みには Ollama への非同期呼び出しが必要ですが、同期 SQL トリガーでは実行できません。
同期書き込みノードの後に明示的なパスによって埋められ、コンテンツが変更されたノードは古くなります
再埋め込みのために埋め込みが削除されました。
埋め込みモデルを備えたローカル Ollama インスタンス (オプション、ベクトル検索用)
git クローン https://github.com/yaminbakoh4-dot/NexusMem.git
cd NexusMem
npmインストール
npm ビルドを実行する
npmリンク
npm link は nexusmem を PATH に配置するため、マシン上の任意のリポジトリに対して実行されます。
nexusmem 初期化
ネクサスメム同期
nexusmem クエリ「なぜ再試行ロジックが存在するのですか」
オプション: 高精度シェルフック
シェル履歴の正確な作業ディレクトリと終了ステータスを取得するには、次の手順を実行します。
nexusmem フックのインストール
これは既存の PowerShell プロンプトを置き換えるのではなくラップし、べき等であり、元に戻されます。
nexusmem フックを使ってきれいに取り外します。
MCP クライアント構成ファイルに以下を追加します。
{
"mcpサーバー": {
"ネクサスメム" : {
"コマンド" : " nexusmem " ,
"args" : [ "mcp " ]
}
}
}
利用可能なツール:
MCP ツール呼び出しには暗黙的なシェル動作が含まれないため、各ツールは明示的な projectRoot を受け取ります。
ディレクトリ。リポジトリがまだセットアップされていない場合、sync_project は最初に自動的に init を実行します。
NexusMem はパッカーの効率 (候補に対する内部パッキング パフォーマンス) を区別します。
セット）とエンドツーエンドのトークンの節約（コンテキスト請求における実際の節約）。二人はそうではない
交換可能であり、最初の引用をあたかも 2 番目の引用であるかのように引用することは、これを具体的に主張するものです。
を防ぐためにセクションが存在します。
ランキングパッカーが生のノードと比較してスコアの低い候補ノードをどれだけ効果的に削除するかを測定します。
厳密なトークン予算内での候補本体の合計:
効率性の由来は、

無関係な低得点候補をテキストからではなく完全に除外する
要約。コーパス サイズが大きくなるにつれて増加し、ノードごとに固定される小さいコーパス サイズでは負になります。
フォーマットのオーバーヘッドは、トリミングするわずかな作業を上回ります。
それを分割するベースラインは仮説です。NexusMem がなければ、これらの候補団体は決して得られなかったでしょう。
コンテキストウィンドウに入りました。この数値はランカーの調整に役立ちます。についての主張としてではありません。
セッションのトークン請求書。
同等の完全なソース ファイルをコンテキストに読み込むことに対するパックされたコンテキスト サイズを測定します。
測定結果: このコードベースに対して評価されたデザイン クエリで ~40% (README.md を読む +
docs/phase-2-spec.md の完全版、約 32,000 文字 ≈ 8 ～ 9,000 トークン、関連するパックされたコンテキストの取得と比較します)。
計測器を使用せず、実際の 1 つのセッションから手作業で集計したものです。桁違いの数値として扱ってください。
この規模では 70% を超える長期目標は達成されず、このリポジトリではそれを実証できません。
ターゲットは、省略することで利益が得られる大規模なリポジトリ (数千のコミット) を記述します。
少数の項目を削除するのではなく、何百もの無関係な履歴項目を削除する必要があります。リポジトリに対するベンチマーク
そのサイズはまだ目立っています。
NexusMem に有利な点の 1 つは、パーセンテージがまったくないことです。会話が変わり、シェル コマンドが発生します。
メモリ内には安価な grep に相当するものはありません。コレクターが録音しなければ、それらは消えてしまうのではなく、
単に取得コストが高くなるだけです。
このリポジトリのコーパス (約 530 ノード)、ウォーム、10 回の実行にわたる p50 で測定:
SQLite 側のすべての作業には合計約 5 ミリ秒かかります。エンドツーエンドの図はローカル埋め込みによって支配されます
これは、このパス上で唯一意味のある遅延ターゲットです。
コマンド
説明
nexusmem 初期化
.nexusmem/ ディレクトリと SQLite スキーマを初期化します。
ネクサスメム同期
新しいイベント (git、shell、docs; --conversation for trans) を取り込みます。

クリプト）。
ネクサスメンステータス
ソースおよびデータベースのステータスごとのメモリ数を出力します。
nexusmem クエリ <テキスト>
ハイブリッド検索を実行し、ランク付けし、コンテキストを標準出力にパックします。
nexusmem スキャン-git
DB に書き込むことなく、git ノードとシグナル スコアのドライラン プレビューを実行します。
nexusmem スキャンシェル
DB に書き込むことなく、シェル履歴ノードのドライラン プレビューを実行します。
nexusmem スキャンドキュメント
DB に書き込むことなく、ドキュメント セクション ノードのドライラン プレビューを実行します。
nexusmem スキャン会話
DB に書き込むことなく会話ノードのプレビューをドライランします。
nexusmem フックのインストール
高精度のシェル ログ用の PowerShell プロファイル ラッパーをインストールします。
nexusmem フックの削除
PowerShell プロファイル ラッパーを削除します。
nexusmem フックのステータス
フックがインストールされているかどうかを報告します。
ネクサスメンMCP
MCP stdio サーバーを起動します。
すべてのコマンドは、現在のディレクトリ以外のリポジトリをターゲットとする -C、--cwd <path> を受け入れます。
便利な同期フラグ: --conversation は、永続化せずに 1 回の実行で会話ソースをオプトインします。
それを設定します。 --no-embed はベクトル埋め込みパスをスキップします。 --rebuild はプロジェクトのノードを削除し、
最初から再摂取します。
<リポジトリ>/.nexusmem/
.gitignore '*' — ワークスペースはそれ自体を無視するため、init は所有していないファイルを編集しません
config.json は読み取り時に検証されます。破損した構成は大声で失敗しますが、静かに失敗することはありません
Memory.db SQLite (WAL): ノード、node_files、nodes_fts、nodes_vec、sync_state
.nexusmem/ を削除しても、同期で再構築できないものは何も失われません。
技術的な制限と特殊なケース
Windows の行末: マークダウン ファイルは、チャンク化の前に CRLF から LF に正規化されます。
正規化されていない CRLF により、段落スプリッター ( \n{2,} ) が起動されなくなります。 \r\n\r\n には何も含まれていません
2 つの連続した \n — ファイル全体を、見出しのないいくつかの粗いブロックに折りたたみます。
Git Rebase / Amend : git 履歴を書き換えると、到達不能なコミットのために孤立したノードが残ります。これら
はrです

eal イベントなので間違いではありませんが、対象となるプルーンはまだ存在しません。
sync --rebuild は、必要に応じてクリーンな再スキャンを強制します。
非セグメント言語: FTS5 unicode61 のトークン化は空白で分割されます。ない言語
空間境界 (タイ語、日本語、中国語) は、再現のためにベクトル検索パスに依存します。
スコープ外のシェル履歴: PowerShell フックが不足しているディレクトリを含まないスクレイピングされたシェル履歴ファイル
コンテキストに関連付けられており、同期が実行されたリポジトリに属します。尻尾に縛られている
保証ではなく近似値です。
PSReadLine 複数行エントリ : プロンプトで複数行にわたって入力された関数は次のように読み取られます。
別々の単一行コマンドであり、再構築されません。
スクレイプ フォールバック ID のドリフト: スクレイプ フォールバックの位置ベースの ID は、次の場合にドリフトする可能性があります。
基礎となる履歴ファイルは、同期の間に先頭からトリミングされます。フックを取り付けるとこれが修正されます。
会話取得の精度: 見出し境界で返信をチャンク化することで、会話の取得精度が向上しました。
長い返信がありますが、体系的に評価されていません。
埋め込みバッチ サイズ: 埋め込みパスは、同期ごとに制限されたバッチを処理します。大きなコーパス
完全に埋め込むには数回の実行が必要です。
フェーズ 1 と 2 が出荷されます。フェーズ 3 が進行中です。
フェーズ 1 — コアの取り込みと取得

[切り捨てられた]

## Original Extract

Contribute to yaminbakoh4-dot/NexusMem development by creating an account on GitHub.

GitHub - yaminbakoh4-dot/NexusMem · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
yaminbakoh4-dot
/
NexusMem
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits docs docs src src tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts View all files Repository files navigation
A local-first persistent memory engine for AI coding agents (Claude Code, Cursor, MCP-based agents).
AI coding assistants forget context once a session ends, and re-uploading the entire repository as
context on every request is slow and expensive. NexusMem records local machine events — git history,
shell commands, docs, and conversation transcripts — into an on-disk SQLite database, returning only
the relevant context slice within a strict token budget.
All data remains local on your machine. No cloud dependencies, accounts, or telemetry.
100% Local-First : SQLite database stored in .nexusmem/ inside your repository using
sqlite-vec and FTS5 . Works fully offline.
Kind-Agnostic Core : Every source normalizes to a single MemoryNode schema, allowing git
commits, shell commands, and documentation to be scored and ranked on an equal basis.
Hybrid Search (BM25 + Vector) : Combines exact keyword matching via SQLite FTS5 (BM25) with
semantic vector search ( sqlite-vec via a local Ollama model) using Reciprocal Rank Fusion (RRF).
RRF fuses on rank position only, never on raw scores, which is what makes it safe to combine a BM25
cost with a vector distance on an unrelated scale. Degrades gracefully to BM25-only if Ollama is
offline.
Ranked, Budgeted Retrieval : Scores candidates using
score = relevance × signal^a × recency^b , then packs nodes into a caller-specified token budget.
Each factor is floored into [floor, 1] rather than [0, 1] , so no single low factor can zero out
a strong match. The exponents a and b are derived, not tuned: relevance is the only
query-derived factor, so each query-independent prior is raised to the power that caps its entire
range at overturning a 2× relevance gap ( span^exponent = 2 , giving a ≈ 0.431 , b ≈ 0.576 ).
MCP Server Native : Exposes search_memory , sync_project , and get_status as Model Context
Protocol (MCP) tools over stdio for Claude Desktop, Cursor, and Windsurf.
git / shell / docs / transcripts
│
▼ collectors/ normalize to one MemoryNode shape
│
▼ store/ SQLite (FTS5 + sqlite-vec)
│
▼ retrieval/ RRF fuse -> rank -> pack to token budget
Key Subsystems
Git Collector : Ingests commits, diff statistics, renames, and conventional commit signals
incrementally via stream iterators. Sync cursors are validated as ancestors of HEAD before being
trusted, so a rebase or amend widens the walk instead of silently skipping commits.
Shell Collector : Scrapes default history files ( PSReadLine , .bash_history ,
.zsh_history ). An optional PowerShell profile hook upgrades capture to include exact timestamps,
working directories, and exit codes (where failed commands receive a higher structural signal).
Docs Collector : Indexes Markdown documentation ( .md files) tracked by git. Line endings are
normalized to LF before chunking to prevent CRLF splitting failures on Windows. Scoped pruning
removes orphaned sections when headings are renamed or deleted, scoped by project and exact source
so it cannot affect git, shell, or conversation nodes.
Conversation Collector (opt-in): Indexes AI assistant transcripts, redacting secrets before
writing to disk. Replies are chunked at heading and bold-lead boundaries rather than stored as
whole exchanges.
Node ids are content-addressed ( sha256(projectId + kind + naturalKey) ), so running sync twice
cannot produce duplicates and ingestion stays correct even if a cursor is lost. Project identity is
derived from the normalized origin URL when one exists, falling back to the absolute path, so two
clones of the same repository share one memory namespace.
nodes_fts is trigger-populated and stays consistent automatically. nodes_vec is not — computing
an embedding requires an async call to Ollama, which a synchronous SQL trigger cannot make — so it is
filled by an explicit pass after sync writes nodes, and a node whose content changes has its stale
embedding dropped for re-embedding.
Local Ollama instance with an embedding model (optional, for vector search)
git clone https://github.com/yaminbakoh4-dot/NexusMem.git
cd NexusMem
npm install
npm run build
npm link
npm link puts nexusmem on your PATH , so it runs against any repository on your machine.
nexusmem init
nexusmem sync
nexusmem query " why does the retry logic exist "
Optional: High-Precision Shell Hook
To capture exact working directory and exit status for shell history:
nexusmem hook install
This wraps your existing PowerShell prompt rather than replacing it, is idempotent, and is undone
cleanly by nexusmem hook remove .
Add the following to your MCP client configuration file:
{
"mcpServers" : {
"nexusmem" : {
"command" : " nexusmem " ,
"args" : [ " mcp " ]
}
}
}
Available tools:
Each tool takes an explicit projectRoot , because MCP tool calls carry no implicit shell working
directory. sync_project runs init first automatically if the repository has not been set up yet.
NexusMem distinguishes between packer efficiency (internal packing performance against candidate
sets) and end-to-end token savings (real-world savings on the context bill). The two are not
interchangeable, and quoting the first as if it were the second is the specific overclaim this
section exists to prevent.
Measures how effectively the ranking packer drops low-scoring candidate nodes relative to the raw
candidate body sum within a strict token budget:
Efficiency is derived from excluding irrelevant low-scoring candidates entirely, not from text
summarization. It increases with corpus size and goes negative on a tiny one, where fixed per-node
formatting overhead outweighs the little there is to trim.
The baseline it divides by is hypothetical: without NexusMem those candidate bodies would never have
entered the context window at all. This figure is useful for tuning the ranker, not as a claim about
a session's token bill.
Measures packed context size against reading the equivalent full source files into context.
Measured result: ~40% on design queries evaluated against this codebase (reading README.md +
docs/phase-2-spec.md in full, ~32k chars ≈ 8–9k tokens, versus retrieving relevant packed context).
Hand-tallied from one real session, not instrumented — treat it as an order-of-magnitude figure.
The long-term >70% target is not met at this scale, and this repository cannot demonstrate it.
The target describes large repositories (thousands of commits) where the win comes from omitting
hundreds of unrelated history items rather than shaving a handful. A benchmark against a repository
of that size is still outstanding.
One caveat in NexusMem's favour is not a percentage at all: the conversation turns and shell commands
in memory have no cheap grep equivalent. Without a collector recording them they are gone, not
merely more expensive to retrieve.
Measured on this repository's corpus (~530 nodes), warm, p50 over 10 runs:
All SQLite-side work totals roughly 5 ms. The end-to-end figure is dominated by the local embedding
call, which is the only meaningful latency target on this path.
Command
Description
nexusmem init
Initializes .nexusmem/ directory and SQLite schema.
nexusmem sync
Ingests new events (git, shell, docs; --conversation for transcripts).
nexusmem status
Prints memory counts per source and database status.
nexusmem query <text>
Executes hybrid search, ranks, and packs context to stdout.
nexusmem scan-git
Dry-run preview of git nodes and signal scores without writing to DB.
nexusmem scan-shell
Dry-run preview of shell history nodes without writing to DB.
nexusmem scan-docs
Dry-run preview of doc section nodes without writing to DB.
nexusmem scan-conversation
Dry-run preview of conversation nodes without writing to DB.
nexusmem hook install
Installs PowerShell profile wrapper for high-precision shell logs.
nexusmem hook remove
Removes the PowerShell profile wrapper.
nexusmem hook status
Reports whether the hook is installed.
nexusmem mcp
Starts the MCP stdio server.
Every command accepts -C, --cwd <path> to target a repository other than the current directory.
Useful sync flags: --conversation opts the conversation source in for one run without persisting
it to config; --no-embed skips the vector-embedding pass; --rebuild drops the project's nodes and
re-ingests from scratch.
<repo>/.nexusmem/
.gitignore '*' — the workspace ignores itself, so init never edits a file it does not own
config.json validated on read; a corrupt config fails loudly, never silently
memory.db SQLite (WAL): nodes, node_files, nodes_fts, nodes_vec, sync_state
Deleting .nexusmem/ loses nothing that sync cannot rebuild.
Technical Limitations & Edge Cases
Windows Line Endings : Markdown files are normalized from CRLF to LF prior to chunking.
Un-normalized CRLF causes the paragraph splitter ( \n{2,} ) to never fire — \r\n\r\n contains no
two consecutive \n — collapsing an entire file into a few coarse, heading-less blocks.
Git Rebase / Amend : Rewriting git history leaves orphaned nodes for unreachable commits. These
are real events, so they are not wrong, but a targeted prune does not exist yet;
sync --rebuild forces a clean re-scan if required.
Non-Segmented Languages : FTS5 unicode61 tokenization splits on whitespace. Languages without
space boundaries (Thai, Japanese, Chinese) rely on the vector search pass for recall.
Unscoped Shell History : Scraped shell history files without the PowerShell hook lack directory
context and are attributed to whichever repository sync was executed from. Bounded to the tail
window, and an approximation rather than a guarantee.
PSReadLine Multi-Line Entries : A function typed across several lines at the prompt is read as
separate single-line commands, not reconstructed.
Scrape-Fallback Id Drift : Position-based ids for the scrape fallbacks can drift if the
underlying history file is trimmed from the front between syncs. Installing the hook fixes this.
Conversation Retrieval Precision : Chunking replies at heading boundaries improved precision on
long replies but has not been evaluated systematically.
Embedding Batch Size : The embedding pass processes a bounded batch per sync ; a large corpus
needs several runs to embed fully.
Phases 1 and 2 are shipped. Phase 3 is in progress.
Phase 1 — Core ingestion and retrieval

[truncated]
