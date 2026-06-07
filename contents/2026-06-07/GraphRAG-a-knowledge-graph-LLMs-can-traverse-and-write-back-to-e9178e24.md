---
source: "https://github.com/mmkumar5401/GraphRag"
hn_url: "https://news.ycombinator.com/item?id=48437054"
title: "GraphRAG – a knowledge graph LLMs can traverse and write back to"
article_title: "GitHub - mmkumar5401/GraphRag · GitHub"
author: "mmkumar"
captured_at: "2026-06-07T18:33:20Z"
capture_tool: "hn-digest"
hn_id: 48437054
score: 2
comments: 0
posted_at: "2026-06-07T17:44:16Z"
tags:
  - hacker-news
  - translated
---

# GraphRAG – a knowledge graph LLMs can traverse and write back to

- HN: [48437054](https://news.ycombinator.com/item?id=48437054)
- Source: [github.com](https://github.com/mmkumar5401/GraphRag)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T17:44:16Z

## Translation

タイトル: GraphRAG – LLM が走査して書き戻すことができるナレッジ グラフ
記事タイトル: GitHub - mmkumar5401/GraphRag · GitHub
説明: GitHub でアカウントを作成して、mmkumar5401/GraphRag の開発に貢献します。

記事本文:
GitHub - mmkumar5401/GraphRag · GitHub
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
mmkumar5401
/
グラフラグ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダー

dファイル
1 コミット 1 コミット グラフ グラフ インジェスト インジェスト モデル モデル クエリ クエリ テスト テスト utils .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md PROGRESS.md PROGRESS.md README.md README.md config.py config.py docker-compose.yml docker-compose.yml main.py main.py 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
GraphRAG — LLM 用の永続メモリ
LLM に永続的な構造化メモリを提供するナレッジ グラフ システム。ドキュメントは、エンティティおよび関係として Neo4j グラフに取り込まれます。エージェント クエリ レイヤーにより、モデルはそのメモリを走査し、発見した新しいファクトを書き込み、新しいソースを取り込むことができます。すべてが 1 つのセッション内で行われます。
ほとんどの LLM は、コンテキスト ウィンドウがリセットされるとすべてを忘れます。このシステムは、セッション間で読み書きできるグラフをモデルに提供します。すべてのファクトは、名前付きエンティティ間の型付けされた関係として Neo4j 内に存在します。モデルはツールを使用してこのグラフをナビゲートし、検出した内容について理由を説明し、新しい知識をコミットしてセッションごとにグラフをよりスマートにします。
Python 3.10+ — python --version で確認する
Neo4j — すべてのメモリを保存するグラフ データベース (以下のオプションを 1 つ選択してください)
Anthropic API キー — console.anthropic.com → API Keys で取得します。
オプション A: Neo4j デスクトップ (最も簡単、ローカル)
neo4j.com/download からダウンロード
新しいプロジェクトを作成 → 追加 → ローカル DBMS
パスワードを設定し、「開始」をクリックします
デフォルトの URI:bolt://localhost:7687
ドッカー実行 \
--name neo4j \
-p 7474:7474 -p 7687:7687 \
-e NEO4J_AUTH=neo4j/あなたの_パスワード \
neo4j:最新
オプション C: AuraDB (クラウド、無料利用枠あり)
neo4j.com/cloud/aura にサインアップしてください
無料のインスタンスを作成します — URI、ユーザー名、パスワードを取得します
git clone https://github.com/your-username/graphrag
cd グラフフラグ
pip インストール anthropic neo4j python-dot

env PyPDF2 リッチ tqdm
4. 環境の設定
プロジェクト ルートに .env ファイルを作成します。
ANTHROPIC_API_KEY=your_key_here
NEO4J_URI=bolt://localhost:7687
NEO4J_USER=ネオ4j
NEO4J_PASSWORD=あなたのパスワード
AuraDB の場合は、インスタンスダッシュボードの URI ( neo4j+s:// で始まる) を使用します。
Python グラフ/query_cli.py --list-nodes
Neo4j が実行中で、認証情報が正しい場合、グラフには 0 個のノードが表示されます (空のグラフは問題ありません。まだ何も取り込んでいません)。接続エラーが発生した場合は、Neo4j が起動しており、.env パスワードが一致していることを確認してください。
ドキュメントを取り込んで質問します。
# PDF またはテキスト ファイルを取り込む
python main.py 取り込みパス/to/your_document.pdf
# 質問する (API クレジットは必要ありません)
cd /path/to/graphrag && クロード
# 次に、「この文書には X について何が書かれていますか?」と尋ねます。
# またはエージェントモードを直接実行します
python main.py query --agentic "このドキュメントの主な概念は何ですか?"
2つの使い方
モード
コスト
読む
事実を書く
ソースの取り込み
クロードコード CLI
読み取りは $0
✅
✅
✅
エージェント ( --agentic )
APIクレジット
✅
✅
✅
どちらのモードでも完全な永続メモリが提供されます。 Claude Code は明示的な CLI スクリプトを使用します。エージェント モードは、ツールを介してすべてを自律的に実行します。
Claude Code CLI (推奨、読み取りコストゼロ)
プロジェクト ディレクトリで Claude Code を開いて、自然に対話します。 Claude Code は以下のグラフ ツールを自動的に使用します。
cd /path/to/graphrag
クロード
グラフの見方
python chart/query_cli.py 「リップル伝播とは何ですか?」
python chart/query_cli.py " question " --hops 3 # より深いトラバーサル
python chart/query_cli.py --seeds " Ripple Propagation,BKT " # 特定のシード
python chart/query_cli.py --list-nodes # グラフ内のすべてを表示
Claude Code は、あらゆる質問に対してこれを自動的に実行し、返されたエッジをネイティブに推論します。API クレジットは消費されません。
摂取する

新しいソースを作成中 (グラフを維持)
# URL から — arxiv 抄録ページは PDF に自動リダイレクトされます
Python グラフ/ingest_cli.py --url " https://arxiv.org/abs/2105.00188 "
# ローカルファイルから
Python グラフ/ingest_cli.py --ファイル パス/to/paper.pdf
Python グラフ/ingest_cli.py --ファイル パス/to/notes.txt
# 生のテキストから
python chart/ingest_cli.py --text " リップル行列が全体への伝播を制御します... "
コンテンツをチャンク化し、Claude Haiku を使用してエンティティと関係を抽出し、すべてを Neo4j に永続的に書き込みます。その後、Claude Code はすぐに新しいノードを探索します。
ファクトをメモリに書き込む (グラフに保持される)
python chart/remember_cli.py "エンティティ A " " RELATION_TYPE " " エンティティ B "
# 例:
python chart/remember_cli.py " リップル行列 " " 制御 " " リップル伝播 "
python chart/remember_cli.py " BKT " " IS_A " " ベイジアン知識トレース "
python chart/remember_cli.py " スロットモデル " " DEPENDS_ON " " 忘却曲線 "
クロード コードがグラフ データを推論しているときに新しい関係を発見すると、これを呼び出して永続的にコミットします。取り込まれたファクトと区別できるようにタグ付けされます。
自然に質問してください。Claude Code はツール呼び出しを処理します。
「波紋伝播って何？」
→ query_cli.py を実行、理由、答え
「この論文を取り込む: https://arxiv.org/abs/2105.00188」
→ ingest_cli.py を実行し、グラフに書き込み、新しいノードを探索します
「この論文は、すでにグラフに含まれているものとどのように関連していますか?」
→ 両方のソースからのシードを使用して query_cli.py を実行し、接続を検索します
「その接続を保存します」
→ remember_cli.py を実行して関係を永続的にコミットします
エージェント モード (自律、API クレジット)
モデルはツールを使用してすべてを自動的に実行します。トラバーサル、ライトバック、セッション中の取り込みは、コマンドを指定しなくても自動的に行われます。
python main.py クエリ --agen

チック「リップル伝播とは何ですか?」
仕組み
1. シーズ探し
ノード名の完全なリストが Neo4j からフェッチされ、クエリとともに Haiku に表示されます。 Haiku は、実際のグラフ語彙から 2 ～ 5 個の開始ノードを選択します (幻覚的なエンティティ名は含まれません)。
2. 俳句横断ループ
Haiku は、次の 3 つのツールを使用してグラフ探索を推進します。
get_neighbours(node_names)
1 つの Neo4j クエリで 8 ～ 12 ノードのバッチのすべての直接隣接ノードを取得します。訪問済みセットにより、同じノードを 2 回再探索することができなくなります。 Haiku が停止することを決定すると、確認ゲートが起動します。トラバース中に出現したがまだ探索されていないすべてのノードが表示され、十分な数があることを確認するよう求められます。これは俳句が本当に完成するまで続きます。
remember(from_entity, relationship_type, to_entity)
新しいファクトをグラフにコミットします。書き込む前に、検証ステップでその事実が現在の走査ですでに見られたエッジによって直接サポートされていることを確認します。つまり、推測的なクロスドメイン推論は拒否されます。受け入れられたファクトにはセッション ID がタグ付けされ、将来のセッションにわたって保持されます。
ingest_source(ソース, ソースタイプ)
セッション中に新しい知識を取り入れます。 URL (source_type: "url") または生のテキスト (source_type: "text") を渡します。システムはトリプルをフェッチし、チャンク化し、抽出し、グラフに書き込みます。 Haiku は新しいノード名を受け取り、 get_neighbours を使用してすぐにそれらを探索できます。各 URL はセッションごとに 1 回だけフェッチされます。重複した取り込み呼び出しはブロックされます。
3. 通過確認ゲート
Haiku が完了を通知すると、システムはトラバース中に出現したがまだ探索されていないすべてのノードを表示します。 Haiku には再考のチャンスが 1 回与えられます。さらに多くのノードを選択した場合、ループは継続します。再び停止すると、合成が始まります。
4. ソネット合成
トラバース中に収集されたすべてのエッジは、単一の ca で Claude Sonnet に送信されます。

拡張思考が有効になっています。ソネットは根拠のある散文的な答えを導き出します。すべてのクレームは、収集されたサブグラフ内の特定のエッジまでトレースする必要があります。思考トレースと回答は両方ともリアルタイムで端末にストリーミングされます。
トラバーサル中のライブ ステータス: [9] 探索: リップル マトリックス、GRU アップデート、リップル伝播アップデート
[memory] は次のように書きました: (A) --[REL]--> (B) ファクトが検証に合格し、コミットされたとき
[メモリ] 拒否: ... 提案された事実が現在のエッジに基づいていない場合
[ingest] arxiv が検出されました — PDF を取得しています: ... 紙の URL が PDF に解決されるとき
[取り込み] 12 チャンク — 抽出の進行に応じてトリプルを抽出します...
[確認中 — N 個の未探索ノードが利用可能] 確認ゲートが起動したとき
思考ブロックがソネットの理由としてストリーミングされる
最終回答はトークンごとにストリーミングされます
探索されたすべてのノードを示すトラバーサル パス テーブル
概要: 探索されたノード、書き込まれた事実、取り込まれたソース、かかった時間
python main.py 取り込みパス/to/document.pdf
python main.py 取り込みパス/to/document.txt
ドキュメントを 800 文字のウィンドウに分割し、Claude Haiku を並行して使用してエンティティと関係を抽出し、Neo4j に書き込みます。べき等 — 安全に再実行できます。
BFS クエリ (高速、合成なし)
python main.py クエリ「リップル伝播とは何ですか?」
コンテキスト内の完全なグラフ
python main.py query --full-graph " リップル伝播とは何ですか? "
重複エンティティを解決する
python main.pysolve # 重複をマージする
python main.pysolve --dry-run # グラフに触れずにプレビューする
同じ概念を参照するエンティティ名 (「BKT」や「ベイジアン ナレッジ トレーシング」など) を検索し、それらを 1 つのノードにマージして、すべての関係を再ポイントします。
ドキュメント/URL/テキストの取り込み
↓
Neo4j グラフ ←─────────────

────┐
↓ │
シーズ探し（エントリーポイントの選定） │
↓ │
トラバーサルループ │
§── get_neighbours / query_cli → メモリの読み取り │
§── 覚えておく / remember_cli → 新しい事実を書く ────┤
└── ingest_source / ingest_cli → フェッチ + 書き込み ──────┘
↓
総合（根拠のある答え）
↓
ストリーミングされた回答
remember / remember_cli を介して書き込まれたファクトと、ingest_source / ingest_cli を介して取り込まれたトリプルにはセッション ID がタグ付けされるため、一括取り込まれたコンテンツと区別でき、後で個別にスコープ設定、エージング、または削除できます。
インジェスト — クリーンなソースの取得
arxiv URL の場合、システムは抽象 HTML ページから PDF にリダイレクトし、PyPDF2 でテキストを抽出し、ナビゲーション メニュー、引用ウィジェット、および HTML ノイズを回避します。汎用 URL は HTML タグの除去に戻ります。
取り込み - エンティティ名の正規化
エンティティ名は抽出中に正規化されます。内部の空白は折りたたまれ、末尾の句読点は削除され、ID は小文字形式から計算されます。これにより、小さなフォーマットの違いによって同じ概念が複数のノードとして表示されるのを防ぎます (例: 「確率行列」と「確率行列」は同じノードに解決されます)。
覚えておいてください — 接地検証 (エージェント モード)
ファクトがエージェント モードで永続メモリに書き込まれる前に、Haiku の検証呼び出しによって、ファクトが現在の tra に見られるグラフ エッジによって直接的かつ明示的にサポートされているかどうかがチェックされます。

[切り捨てられた]

## Original Extract

Contribute to mmkumar5401/GraphRag development by creating an account on GitHub.

GitHub - mmkumar5401/GraphRag · GitHub
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
mmkumar5401
/
GraphRag
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit graph graph ingestion ingestion models models query query tests tests utils utils .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md PROGRESS.md PROGRESS.md README.md README.md config.py config.py docker-compose.yml docker-compose.yml main.py main.py requirements.txt requirements.txt View all files Repository files navigation
GraphRAG — Persistent Memory for LLMs
A knowledge graph system that gives LLMs persistent, structured memory. Documents are ingested into a Neo4j graph as entities and relationships. An agentic query layer lets the model traverse that memory, write new facts it discovers, and pull in new sources — all within a single session.
Most LLMs forget everything when the context window resets. This system gives the model a graph it can read from and write to across sessions. Every fact lives in Neo4j as a typed relationship between named entities. The model navigates this graph using tools, reasons over what it finds, and commits new knowledge back — making the graph smarter with every session.
Python 3.10+ — check with python --version
Neo4j — the graph database that stores all memory (pick one option below)
Anthropic API key — get one at console.anthropic.com → API Keys
Option A: Neo4j Desktop (easiest, local)
Download from neo4j.com/download
Create a new project → Add → Local DBMS
Set a password, then click Start
Default URI: bolt://localhost:7687
docker run \
--name neo4j \
-p 7474:7474 -p 7687:7687 \
-e NEO4J_AUTH=neo4j/your_password \
neo4j:latest
Option C: AuraDB (cloud, free tier available)
Sign up at neo4j.com/cloud/aura
Create a free instance — you'll get a URI, username, and password
git clone https://github.com/your-username/graphrag
cd graphrag
pip install anthropic neo4j python-dotenv PyPDF2 rich tqdm
4. Configure environment
Create a .env file in the project root:
ANTHROPIC_API_KEY=your_key_here
NEO4J_URI=bolt://localhost:7687
NEO4J_USER=neo4j
NEO4J_PASSWORD=your_password
For AuraDB, use the URI from your instance dashboard (starts with neo4j+s:// ).
python graph/query_cli.py --list-nodes
If Neo4j is running and credentials are correct, you'll see 0 nodes in graph: (empty graph is fine — you haven't ingested anything yet). If you see a connection error, check that Neo4j is started and your .env password matches.
Ingest a document and ask a question:
# Ingest a PDF or text file
python main.py ingest path/to/your_document.pdf
# Ask a question (no API credits needed)
cd /path/to/graphrag && claude
# Then ask: "What does the document say about X?"
# Or run the agentic mode directly
python main.py query --agentic " What are the main concepts in the document? "
Two ways to use it
Mode
Cost
Read
Write facts
Ingest sources
Claude Code CLI
$0 for reads
✅
✅
✅
Agentic ( --agentic )
API credits
✅
✅
✅
Both modes give you full persistent memory. Claude Code uses explicit CLI scripts. Agentic mode drives everything autonomously via tools.
Claude Code CLI (recommended, zero read cost)
Open Claude Code in the project directory and interact naturally. Claude Code automatically uses the graph tools below.
cd /path/to/graphrag
claude
Reading the graph
python graph/query_cli.py " What is ripple propagation? "
python graph/query_cli.py " question " --hops 3 # deeper traversal
python graph/query_cli.py --seeds " Ripple Propagation,BKT " # specific seeds
python graph/query_cli.py --list-nodes # see everything in the graph
Claude Code runs this automatically for any question, then reasons over the returned edges natively — no API credits consumed.
Ingesting new sources (persists to graph)
# From a URL — arxiv abstract pages are auto-redirected to PDF
python graph/ingest_cli.py --url " https://arxiv.org/abs/2105.00188 "
# From a local file
python graph/ingest_cli.py --file path/to/paper.pdf
python graph/ingest_cli.py --file path/to/notes.txt
# From raw text
python graph/ingest_cli.py --text " The ripple matrix governs propagation across... "
Chunks the content, extracts entities and relationships using Claude Haiku, writes everything to Neo4j permanently. Claude Code then immediately explores the new nodes.
Writing facts to memory (persists to graph)
python graph/remember_cli.py " Entity A " " RELATION_TYPE " " Entity B "
# Examples:
python graph/remember_cli.py " Ripple Matrix " " GOVERNS " " Ripple Propagation "
python graph/remember_cli.py " BKT " " IS_A " " Bayesian Knowledge Tracing "
python graph/remember_cli.py " Slot Model " " DEPENDS_ON " " Forgetting Curve "
When Claude Code discovers a new relationship while reasoning over graph data, it calls this to commit it permanently — tagged so it's distinguishable from ingested facts.
Just ask naturally — Claude Code handles the tool calls:
"What is ripple propagation?"
→ runs query_cli.py, reasons over edges, answers
"Ingest this paper: https://arxiv.org/abs/2105.00188"
→ runs ingest_cli.py, writes to graph, explores new nodes
"How does this paper relate to what's already in the graph?"
→ runs query_cli.py with seeds from both sources, finds connections
"Save that connection"
→ runs remember_cli.py to commit the relationship permanently
Agentic mode (autonomous, API credits)
The model drives everything itself using tools — traversal, write-back, and mid-session ingestion happen automatically without you specifying commands.
python main.py query --agentic " What is ripple propagation? "
How it works
1. Seed finding
The full list of node names is fetched from Neo4j and shown to Haiku alongside the query. Haiku picks 2–5 starting nodes from the real graph vocabulary — no hallucinated entity names.
2. Haiku traversal loop
Haiku drives graph exploration using three tools:
get_neighbours(node_names)
Fetches all direct neighbours for a batch of 8–12 nodes in one Neo4j query. A visited set prevents re-exploring the same node twice. When Haiku decides to stop, a confirmation gate fires: it's shown all nodes that appeared during traversal but weren't explored yet, and asked to confirm it has enough. This continues until Haiku is genuinely done.
remember(from_entity, relation_type, to_entity)
Commits a new fact to the graph. Before writing, a validation step checks that the fact is directly supported by edges already seen in the current traversal — speculative cross-domain inferences are rejected. Accepted facts are tagged with the session ID and persist across future sessions.
ingest_source(source, source_type)
Pulls in new knowledge mid-session. Pass a URL ( source_type: "url" ) or raw text ( source_type: "text" ). The system fetches, chunks, extracts triples, and writes them to the graph. Haiku receives the new node names and can immediately explore them with get_neighbours . Each URL is only fetched once per session — duplicate ingest calls are blocked.
3. Traversal confirmation gate
When Haiku signals it's done, the system shows it all nodes that appeared during traversal but weren't explored yet. Haiku gets one chance to reconsider — if it picks more nodes, the loop continues; if it stops again, synthesis begins.
4. Sonnet synthesis
All edges collected during traversal are sent to Claude Sonnet in a single call with extended thinking enabled. Sonnet produces a grounded, prose answer. Every claim must trace to a specific edge in the collected subgraph. Both the thinking trace and the answer stream to the terminal in real time.
Live status during traversal: [9] exploring: Ripple Matrix, GRU update, Ripple-Propagated Updates
[memory] wrote: (A) --[REL]--> (B) when a fact passes validation and is committed
[memory] rejected: ... when a proposed fact isn't grounded in the current edges
[ingest] arxiv detected — fetching PDF: ... when a paper URL is resolved to its PDF
[ingest] 12 chunks — extracting triples... as extraction progresses
[confirming — N unexplored nodes available] when the confirmation gate fires
Thinking block streamed as Sonnet reasons
Final answer streamed token by token
Traversal path table showing every node explored
Summary: nodes explored, facts written, sources ingested, time taken
python main.py ingest path/to/document.pdf
python main.py ingest path/to/document.txt
Chunks the document into 800-character windows, extracts entities and relationships using Claude Haiku in parallel, and writes them to Neo4j. Idempotent — safe to re-run.
BFS query (fast, no synthesis)
python main.py query " What is ripple propagation? "
Full graph in context
python main.py query --full-graph " What is ripple propagation? "
Resolve duplicate entities
python main.py resolve # merge duplicates
python main.py resolve --dry-run # preview without touching the graph
Finds entity names that refer to the same concept (e.g. "BKT" and "Bayesian Knowledge Tracing") and merges them into a single node, re-pointing all relationships.
Ingest documents / URLs / text
↓
Neo4j graph ←──────────────────────────────────────┐
↓ │
Seed finding (picks entry points) │
↓ │
Traversal loop │
├── get_neighbours / query_cli → read memory │
├── remember / remember_cli → write new facts ────┤
└── ingest_source / ingest_cli → fetch + write ──────┘
↓
Synthesis (grounded answer)
↓
Streamed answer
Facts written via remember / remember_cli and triples ingested via ingest_source / ingest_cli are tagged with session ID, so they're distinguishable from bulk-ingested content and can later be scoped, aged, or removed independently.
Ingestion — clean source fetching
For arxiv URLs, the system redirects from the abstract HTML page to the PDF and extracts text with PyPDF2, avoiding navigation menus, citation widgets, and HTML noise. Generic URLs fall back to HTML tag stripping.
Ingestion — entity name normalisation
Entity names are normalised during extraction: internal whitespace is collapsed, trailing punctuation is stripped, and IDs are computed from the lowercase form. This prevents the same concept appearing as multiple nodes due to minor formatting differences (e.g. "Stochastic Matrices" and "Stochastic matrices" resolve to the same node).
Remember — grounding validation (agentic mode)
Before any fact is written to permanent memory in agentic mode, a Haiku validation call checks that the fact is directly and explicitly supported by the graph edges seen in the current tra

[truncated]
