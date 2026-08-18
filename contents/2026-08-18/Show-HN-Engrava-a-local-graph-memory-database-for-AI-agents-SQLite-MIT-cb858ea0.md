---
source: "https://github.com/sovantica/engrava"
hn_url: "https://news.ycombinator.com/item?id=49345004"
title: "Show HN: Engrava – a local graph memory database for AI agents (SQLite, MIT)"
article_title: "GitHub - sovantica/engrava: The memory database for AI agents - graph memory, hybrid search, tamper-evident thought/edge journal. · GitHub"
image: "https://repository-images.githubusercontent.com/1256473670/1645fae1-d36a-4208-b33d-c265af0268a9"
author: "przemarzec"
captured_at: "2026-08-18T13:36:30Z"
capture_tool: "hn-digest"
hn_id: 49345004
score: 3
comments: 0
posted_at: "2026-08-18T13:01:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Engrava – a local graph memory database for AI agents (SQLite, MIT)

- HN: [49345004](https://news.ycombinator.com/item?id=49345004)
- Source: [github.com](https://github.com/sovantica/engrava)
- Score: 3
- Comments: 0
- Posted: 2026-08-18T13:01:08Z

## Translation

タイトル: Show HN: Engrava – AI エージェント用のローカル グラフ メモリ データベース (SQLite、MIT)
記事のタイトル: GitHub - sovantica/engrava: AI エージェント用のメモリ データベース - グラフ メモリ、ハイブリッド検索、改ざん防止思考/エッジ ジャーナル。 · GitHub
説明: AI エージェント用のメモリ データベース - グラフ メモリ、ハイブリッド検索、改ざん防止思考/エッジ ジャーナル。 - ソバンティカ/エングラヴァ

記事本文:
GitHub - sovantica/engrava: AI エージェント用のメモリ データベース - グラフ メモリ、ハイブリッド検索、改ざん防止思考/エッジ ジャーナル。 · GitHub
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
ソバンティカ
/
刻む
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
dev ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
275 コミット 275 コミット フォルダーとファイル
.github .github docs docs 例 例 スクリプト スクリプト src/ engra

va src/ engrava テスト テスト .commitlintrc.js .commitlintrc.js .cursorrules .cursorrules .gitattributes .gitattributes .gitignore .gitignore .releaserc.json .releaserc.json AGENTS.md AGENTS.md BRANCHING.md BRANCHING.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md SECURITY.md SECURITY.md commit-scopes.json commit-scopes.json pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのメモリ データベース。
グラフ メモリ、ハイブリッド検索、および改ざん防止思考/エッジ ジャーナル — 1 回の pip インストールで、サーバーや LLM は不要です。
Engrava は、AI エージェント メモリ用のスタンドアロンの組み込みデータベースです。上に構築
SQLite、思考 CRUD、エッジベースのナレッジ グラフ、埋め込みベースを提供します。
類似検索、全文検索（FTS5/BM25）、宣言型拡張機能
システム — すべてが 1 つのパッケージにまとめられており、外部サービスへの依存関係はありません。
ベンチマーク結果は公開されており、実行は別のリポジトリから再現できます。毎
パブリッシュされた実行はグループ A —memory_pipeline_llms: [] 、メモリのどこにも言語モデルがありません
層。これは測定値ではなくアーキテクチャの特性であるため、リリースをまたいで維持されます。
ベンチマーク結果 — ライブ テーブル: すべてのパブリッシュされた行と、それが属する比較セグメント。
engrava-benchmark — ランナー。それを複製し、PyPI からのパッケージに対して結果を再現します。マサチューセッツ工科大学
ディスカッション — 質問。再現に関する質問は Q&A カテゴリに属します。
セマンティック検索を備えた会話ストレージ
連想リンク付きの研究ノート
埋め込みを含む思考グラフを必要とするアプリケーション
pip インストール engrava
オプションの追加物:
pip インストール ' engrava[vec] '

# sqlite-vec ベクトル検索バックエンド
pip install ' engrava[embeddings-local] ' # 文変換の埋め込み (ローカル モデル)
pip install ' engrava[embeddings-openai] ' # OpenAI 互換の埋め込み API
pip install ' engrava[embeddings-ollama] ' # Ollama ローカル埋め込みサーバー
pip install ' engrava[embeddings-hf] ' # HuggingFace Inference API の埋め込み
夢見ること/統合することと知識グラフには余分なものは必要ありません - それらは一部です
of the base install.
メモリを保存し、2 回の呼び出しで検索します。ID の生成や記録は必要ありません。
to assemble:
import asyncio
import aiosqlite
engrava インポートから SqliteEngravaCore
async def main () -> なし :
# SqliteEngravaCore は、オープンな aiosqlite 接続をラップします。
async with aiosqlite . connect ( ":memory:" ) as conn :
conn . row_factory = aiosqlite 。行
ストア = SqliteEngravaCore ( conn )
await store . ensure_schema ()
await store .覚えておいてください (「Python は AI エージェントに最適です」)
await store .覚えておいてください (「SQLite にはサーバーは必要ありません」)
result = await store .思い出してください (「エージェントにとってどの言語が適していますか?」)
thought_id の場合、結果のスコア。 results :
thought = await store . get_thought (thought_id)
if thought is not None :
print ( f" { 思考 . エッセンス } (スコア: { スコア :.3f } )" )
asyncio .実行 (メイン ())
remember() はテキストを思考として保存し (その ID を生成します)、
保存されている ThoughtRecord を返します。 remember() は、次と同じハイブリッド検索を実行します。
search_hybrid() を実行し、ランク付けされた結果を返します。フルコントロールの場合 — 設定
優先順位、思考タイプ、メタデータ、または書き込み時の認知サイクル — を構築する
Thought自分自身を記録し、 create_thought() を呼び出します。
ここから、思考をタイプされたエッジと結び付け、
MindQL でクエリするか、完全なコマンドを実行します。
クイックスタートガイドの「取り込み」→「夢」→「検索」ツアー。
engrava インポートから SqliteEngravaCore
#from_co

nfig は接続を開いて所有します。これを非同期コンテキスト マネージャーとして使用します。
await SqliteEngravaCore を使用した非同期。 from_config ( "engrava.yaml" ) をストアとして:
# スキーマは from_config によってすでに適用されています。
考え＝ストアを待ちます。 get_thought ( "some-id" )
完全な YAML スキーマについては、docs/configuration.md を参照してください。
自動スキーマ移行は最初の接続時に実行されます。を参照してください。
互換性に関する注意事項に関するアップグレード ガイド、バックアップ ガイダンス、
トラブルシューティングの手順。
完全なライフサイクル管理により、考えを作成、読み取り、更新、アーカイブします。
すべてのモデルは凍結された Pydantic オブジェクトです。突然変異は、EVOLUTION() を介して発生します。
思考を型付けされた重み付けされたエッジにリンクします。エッジ タイプには ASSOCIATED 、
DEPENDS_ON 、 DERIVED_FROM 、 CONSOLIDATED_FROM (夢によって作成)、および
CONTESTED_BY 。
埋め込みを思考と一緒に保存し、組み込みの NumPy で検索します
cosine バックエンドまたはオプションの sqlite-vec バックエンド ( pip install 'engrava[vec]' )。プラグイン可能な埋め込みプロバイダー:
BM25 ランキングの SQLite FTS5 仮想テーブル。ハイブリッド検索はベクトルを結合します
類似性、テキストの関連性、最新性、優先度、グラフの接続性。信号
クエリに対して実行できないものはスキップされ、残りの重みは
再配布されました。
思考グラフの宣言型クエリ言語:
思考を検索 WHERE thought_type = 'OBSERVATION' AND priority = 'P1' LIMIT 10
COUNT 個の思考 WHERE lifecycle_status = 'ACTIVE'
SELECT thought_id、エッセンス FROM thought WHERE thought_type = 'BELIEF'
MindQLExecutor または
拡張マニフェスト ;ライフサイクル フック レジストリは予約されており、予約されていません。
中核執行者から相談を受けます。
EngravaHooksProtocol を介して思考のライフサイクルに接続します。サブクラス
選択したアクティブなメソッドのみが必要な場合の DefaultEngravaHooks:
from engrava import DefaultEngravaHooks 、 ThoughtRecord
クラスMyHooks

( DefaultEngravaHooks ):
async def on_store (self , thought : ThoughtRecord ) -> ThoughtRecord :
# 永続化後に返されるオブジェクトを観察または強化します。
思考を返す
async def Decay_function (
self 、思考 : ThoughtRecord 、elapsed_cycles : int
) -> 浮動小数点 :
# 有効な Memory Hygiene パスに減衰乗数を指定します。
1.0を返す
コアは現在、 on_store 、 on_retrieve 、およびdecay_function を呼び出します。
on_store は、ソースが永続的であると考えられた後に実行されます。戻り値を変更する
永続化された行は書き換えられません。スコア関数と
minql_extension_registry() は予約されたプロトコル メソッドのままであり、
コアによって呼び出されます。
夢を見る/記憶の定着
定期的な記憶の統合のための組み込みの DreamingExtension — スコア
構成可能なシグナルを介して思考を行い、価値の高いエントリを促進し、
意味的に関連したクラスタリングによって REFLECTION 思考を作成する
思考と重心埋め込みの計算 (LLM は必要ありません)。利用可能
0.3.0以降。
→ 再現性については docs/benchmarks.md を参照してください。
証拠 (合成ベンチマーク スイートは約 5 分で実行可能)。
記憶維持の減算的半分とドリーミングの組み合わせ: オプトイン、
信号の少ない冷たい思考をアーカイブする、可逆的な非 LLM ループ。
個別にオプトインされたステップでは、サイクルと
壁掛け時計の復元ウィンドウ。デフォルトではオフです。有効にすると、アーカイブされた思考が削除されます
デフォルトの取得の復元が可能です (restore_thought / include_archived)。
→ 詳細については、docs/memory-hygiene.md を参照してください。
ループ、保護、ウィンドウの復元、および正直な削除姿勢。
タンパーエビデント思考/エッジジャーナル
思考とエッジの突然変異を記録するオプトイン ハッシュ チェーン ジャーナル (さらに
アクションステータス / 検証ステータス遷移) SHA-256 リンクとして、前/後
エントリ — 改ざんが明らかな思考/エッジジャーナル

データベース全体の監査ではありません
(埋め込みとアクションの作成は対象外です)。デフォルトではオフ、1 つの設定フラグが
有効にする。 store.journal.get_entries(...) を使用して履歴をクエリし、
store.verify_journal() を使用したチェーン。ディスク上にあるチェーンを監査します。
現在のjournal.enabled状態とは独立した、store.journalライター
ジャーナリングがオフの間はハンドルは存在しませんが、以前のセッションはチェーンされます
write はまだjournal_entryにあるため、まだ検証する必要があります。
→ 有効化、クエリ、
検証とセキュリティ モデル (「改ざん防止」で何が行われるか、何が行われないか)
保証）。
1 つの EngravaManager の下で複数の独立したデータベースを実行します。
pathlibインポートパスから
engrava インポートから EngravaManager
EngravaManager ( data_dir = Path ( "./data" )) を mgr として非同期:
エージェント_a = マネージャを待ちます。 get_store ( "エージェント-a" )
Agent_b = マネージャを待ちます。 get_store ( "エージェント-b" )
# 完全に分離されたデータベース
MCPサーバー
エージェントのメモリ サーバーとして Engrava を使用したいですか? MCP サーバーは独自のサーバーとして出荷されます
パッケージ、engrava-mcp — 読み取り機能を備えたネイティブ stdio サーバー (HTTP シムなし)
ツール、オプションの書き込みツール、アタッチ可能な engrava:// リソース、およびガイド付き
任意の MCP クライアント (Claude Desktop、Claude Code、Cursor、Windsurf、
VS コード):
uvx engrava-mcp # または: pip install engrava-mcp
engrava-mcp は engrava を推移的に取り込むので、それをインストールすると、
engrava ライブラリをインポートします。を参照してください。
インストール用の engrava-mcp パッケージ、
クライアント構成、完全なツール/リソース/プロンプトのリファレンス、および読み取り専用
モード。
engrava --db mydata.db info # データベース統計
engrava --db mydata.db query " FIND thought WHERE thought_type = 'OBSERVATION' LIMIT 5 "
engrava --db mydata.db スナップショット -o バックアップ.jsonl
engrava --db mydata.db復元 -i バックアップ.jsonl
engrava --db mydata.db gc # ARCHI を収集する

VED の思考 + そのエッジ/埋め込み/アクション
engrava --db mydata.db merge # スキーマが最新であることを確認します
engrava --db mydata.db export -oportable.json
gc は、アーカイブされた思考を、エッジに触れているすべての思考とともに物理的に削除します
どちらかの端 (もう一方の端がまだ生きているエッジを含む) の埋め込み
と、それらをソースとするアクションを削除し、ベクトル インデックスを調整します。
すべての vec0 行は、埋め込み行を所有しません。 vec0 インデックス付きストアでは、
sqlite-vec をロードできません - 最も一般的な理由は、engrava[vec] がロードされていないためです
インストール済み — 削除しようとしているパスは、何かを削除する前に停止し、
それらのベクトルをインデックスに取り残さず、1 を終了します。その後、何も到達できなくなります。
それらを通過します。
engrava info は、によって公開されているのと同じメトリクス スナップショット コントラクトをレンダリングするようになりました。
store.metrics() を待ちます。
すべてのコマンドとオプションについては、CLI リファレンスを参照してください。
同時読み取り用の WAL モードを備えた SQLite
凍結された Pydantic モデル — 不変ドメイン オブジェクト
非同期ファースト — すべての I/O は aiosqlite 経由
フックベースの拡張 — モンキーパッチなし
テンプレート メソッド パターン — 拡張スキーマのサブクラス SqliteEngravaCore
外部サービスなし - すべてがローカルでインプロセスで実行されます
ドキュメントインデックス — 学習、構築、運用、拡張、移行などのタスクごとにパスを選択します
中心となる概念

[切り捨てられた]

## Original Extract

The memory database for AI agents - graph memory, hybrid search, tamper-evident thought/edge journal. - sovantica/engrava

GitHub - sovantica/engrava: The memory database for AI agents - graph memory, hybrid search, tamper-evident thought/edge journal. · GitHub
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
sovantica
/
engrava
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
dev Branches Tags Go to file Code Open more actions menu Latest commit
275 Commits 275 Commits Folders and files
.github .github docs docs examples examples scripts scripts src/ engrava src/ engrava tests tests .commitlintrc.js .commitlintrc.js .cursorrules .cursorrules .gitattributes .gitattributes .gitignore .gitignore .releaserc.json .releaserc.json AGENTS.md AGENTS.md BRANCHING.md BRANCHING.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md SECURITY.md SECURITY.md commit-scopes.json commit-scopes.json pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
The memory database for AI agents.
Graph memory, hybrid search, and a tamper-evident thought/edge journal — one pip install , no server, no LLM.
Engrava is a standalone embedded database for AI agent memory. Built on
SQLite, it provides thought CRUD, edge-based knowledge graphs, embedding-based
similarity search, full-text search (FTS5/BM25), and a declarative extension
system — all in a single package with zero external service dependencies.
Benchmark results are published, and the runs are reproducible from a separate repository. Every
published run is Group A — memory_pipeline_llms: [] , no language model anywhere in the memory
layer. That is a property of the architecture rather than a measurement, so it holds across releases.
Benchmark results — the live table: every published row with the comparability segment it belongs to.
engrava-benchmark — the runner. Clone it and reproduce a result against the package from PyPI. MIT.
Discussions — questions; reproduction questions belong in the Q&A category.
Conversation storage with semantic search
Research notes with associative linking
Any application that needs a thought-graph with embeddings
pip install engrava
Optional extras:
pip install ' engrava[vec] ' # sqlite-vec vector search backend
pip install ' engrava[embeddings-local] ' # sentence-transformers embeddings (local model)
pip install ' engrava[embeddings-openai] ' # OpenAI-compatible embeddings API
pip install ' engrava[embeddings-ollama] ' # Ollama local embeddings server
pip install ' engrava[embeddings-hf] ' # HuggingFace Inference API embeddings
Dreaming/consolidation and the knowledge graph need no extra — they are part
of the base install.
Store a memory and search for it in two calls — no IDs to generate, no record
to assemble:
import asyncio
import aiosqlite
from engrava import SqliteEngravaCore
async def main () -> None :
# SqliteEngravaCore wraps an open aiosqlite connection.
async with aiosqlite . connect ( ":memory:" ) as conn :
conn . row_factory = aiosqlite . Row
store = SqliteEngravaCore ( conn )
await store . ensure_schema ()
await store . remember ( "Python is great for AI agents" )
await store . remember ( "SQLite needs no server" )
result = await store . recall ( "what language is good for agents?" )
for thought_id , score in result . results :
thought = await store . get_thought ( thought_id )
if thought is not None :
print ( f" { thought . essence } (score: { score :.3f } )" )
asyncio . run ( main ())
remember() stores the text as a thought (generating its ID for you) and
returns the stored ThoughtRecord ; recall() runs the same hybrid search as
search_hybrid() and returns the ranked results. For full control — setting
priority, thought type, metadata, or the cognitive cycle on a write — build a
ThoughtRecord yourself and call create_thought() .
From here, link thoughts with typed edges ,
query them with MindQL , or run the full
ingest → dream → search tour in the Quick Start guide .
from engrava import SqliteEngravaCore
# from_config opens and OWNS the connection — use it as an async context manager.
async with await SqliteEngravaCore . from_config ( "engrava.yaml" ) as store :
# The schema is already applied by from_config.
thought = await store . get_thought ( "some-id" )
See docs/configuration.md for the full YAML schema.
Automatic schema migration runs on first connection. See the
upgrade guide for compatibility notes, backup guidance, and
troubleshooting steps.
Create, read, update, and archive thoughts with full lifecycle management.
All models are frozen Pydantic objects — mutations happen via evolve() .
Link thoughts with typed, weighted edges. Edge types include ASSOCIATED ,
DEPENDS_ON , DERIVED_FROM , CONSOLIDATED_FROM (created by dreaming), and
CONTESTED_BY .
Store embeddings alongside thoughts and search them with the built-in NumPy
cosine backend or the optional sqlite-vec backend ( pip install 'engrava[vec]' ). Pluggable embedding providers:
SQLite FTS5 virtual table with BM25 ranking. Hybrid search combines vector
similarity, text relevance, recency, priority, and graph connectivity. Signals
that cannot run for a query are skipped and the remaining weights are
redistributed.
Declarative query language for the thought-graph:
FIND thoughts WHERE thought_type = 'OBSERVATION' AND priority = 'P1' LIMIT 10
COUNT thoughts WHERE lifecycle_status = 'ACTIVE'
SELECT thought_id, essence FROM thought WHERE thought_type = 'BELIEF'
Extensible with custom commands through MindQLExecutor or an
ExtensionManifest ; the lifecycle hook registry is reserved and is not
consulted by the core executor.
Plug into the thought lifecycle via EngravaHooksProtocol . Subclass
DefaultEngravaHooks when you only need selected active methods:
from engrava import DefaultEngravaHooks , ThoughtRecord
class MyHooks ( DefaultEngravaHooks ):
async def on_store ( self , thought : ThoughtRecord ) -> ThoughtRecord :
# Observe or enrich the object returned after persistence.
return thought
async def decay_function (
self , thought : ThoughtRecord , elapsed_cycles : int
) -> float :
# Supply a decay multiplier to an enabled Memory Hygiene pass.
return 1.0
Core currently invokes on_store , on_retrieve , and decay_function .
on_store runs after the source thought is durable; changing its return value
does not rewrite the persisted row. score_function and
mindql_extension_registry() remain reserved protocol methods and are not
called by core.
Dreaming / Memory Consolidation
Built-in DreamingExtension for periodic memory consolidation — scores
thoughts via configurable signals, promotes high-value entries, and
creates REFLECTION thoughts by clustering semantically related
thoughts and computing centroid embeddings (no LLM required). Available
since 0.3.0.
→ See docs/benchmarks.md for reproducible
evidence (synthetic benchmark suite runnable in ~5 minutes).
The subtractive half of memory maintenance, paired with Dreaming: an opt-in,
reversible , no-LLM loop that archives cold, low-signal thoughts — and, as a
separately opted-in step, garbage-collects them only after both a cycle and a
wall-clock restore window. OFF by default; once enabled, archived thoughts drop out
of default retrieval and can be restored ( restore_thought / include_archived ).
→ See docs/memory-hygiene.md for the
loop, protection, restore windows, and the honest deletion posture.
Tamper-Evident Thought/Edge Journal
Opt-in hash-chain journal that records thought and edge mutations (plus
action status / verification_status transitions) as SHA-256-linked, before/after
entries — a tamper-evident thought/edge journal, not a whole-database audit
(embeddings and action creation are not covered). Off by default, one config flag to
enable. Query history with store.journal.get_entries(...) and validate the
chain with store.verify_journal() , which audits whatever chain is on disk
independent of the current journal.enabled state — the store.journal writer
handle does not exist while journaling is off, but the chain earlier sessions
wrote is still in journal_entry and still needs verifying.
→ See docs/audit-trail.md for enabling, querying,
verification, and the security model (what "tamper-evident" does and does not
guarantee).
Run multiple independent databases under one EngravaManager :
from pathlib import Path
from engrava import EngravaManager
async with EngravaManager ( data_dir = Path ( "./data" )) as mgr :
agent_a = await mgr . get_store ( "agent-a" )
agent_b = await mgr . get_store ( "agent-b" )
# Completely isolated databases
MCP Server
Want Engrava as a memory server for your agent? The MCP server ships as its own
package, engrava-mcp — a native stdio server (no HTTP shim) with read
tools, optional write tools, attachable engrava:// resources, and guided
prompts, for any MCP client (Claude Desktop, Claude Code, Cursor, Windsurf,
VS Code):
uvx engrava-mcp # or: pip install engrava-mcp
engrava-mcp pulls engrava in transitively, so installing it also gives you
the import engrava library. See the
engrava-mcp package for install,
client configuration, the full tool/resource/prompt reference, and read-only
mode.
engrava --db mydata.db info # Database stats
engrava --db mydata.db query " FIND thoughts WHERE thought_type = 'OBSERVATION' LIMIT 5 "
engrava --db mydata.db snapshot -o backup.jsonl
engrava --db mydata.db restore -i backup.jsonl
engrava --db mydata.db gc # Collect ARCHIVED thoughts + their edges/embeddings/actions
engrava --db mydata.db migrate # Ensure schema is up-to-date
engrava --db mydata.db export -o portable.json
gc physically deletes ARCHIVED thoughts together with every edge touching one
on either end — including edges whose other end is still live — their embeddings
and the actions sourced from them, then reconciles the vector index by removing
every vec0 row no embedding row owns; on a vec0 -indexed store where
sqlite-vec cannot be loaded — most commonly because engrava[vec] is not
installed — a pass that is about to delete stops before deleting anything and
exits 1 rather than stranding those vectors in an index nothing can then reach
them through.
engrava info now renders the same metrics snapshot contract exposed by
await store.metrics() .
See the CLI reference for every command and option.
SQLite with WAL mode for concurrent reads
Frozen Pydantic models — immutable domain objects
Async-first — all I/O via aiosqlite
Hook-based extension — zero monkey-patching
Template method pattern — subclass SqliteEngravaCore for extended schemas
Zero external services — everything runs locally in-process
Documentation Index — choose a path by task: learn, build, operate, extend, or migrate
Core Concepts

[truncated]
