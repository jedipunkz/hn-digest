---
source: "https://github.com/Vectorlink-Labs/coalent"
hn_url: "https://news.ycombinator.com/item?id=49345960"
title: "Show HN: Coalent – an LLM answer cache that invalidates when source docs change"
article_title: "GitHub - Vectorlink-Labs/coalent: Real-time, provenance-invalidated cognitive cache for AI agents and RAG — build understanding once, reuse it everywhere, keep it fresh. · GitHub"
image: "https://opengraph.githubassets.com/0361fa9f9ddc93016dca978819fa56479f367472f4fc9f2a7e53a419094b3807/Vectorlink-Labs/coalent"
author: "nisarg-pujara"
captured_at: "2026-08-18T14:23:36Z"
capture_tool: "hn-digest"
hn_id: 49345960
score: 1
comments: 0
posted_at: "2026-08-18T14:16:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Coalent – an LLM answer cache that invalidates when source docs change

- HN: [49345960](https://news.ycombinator.com/item?id=49345960)
- Source: [github.com](https://github.com/Vectorlink-Labs/coalent)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T14:16:34Z

## Translation

タイトル: Show HN: Coalent – ソース ドキュメントが変更されたときに無効にする LLM 回答キャッシュ
記事のタイトル: GitHub - Vectorlink-Labs/coalent: AI エージェントと RAG のためのリアルタイムの出自無効化コグニティブ キャッシュ — 一度理解を構築すれば、どこでも再利用でき、新鮮さを保ちます。 · GitHub
説明: AI エージェントと RAG 向けの出自が無効化されたリアルタイムの認知キャッシュ — 一度理解を構築すれば、どこでも再利用でき、最新の状態に保たれます。 - Vectorlink-Labs/コーレント

記事本文:
GitHub - Vectorlink-Labs/coalent: AI エージェントおよび RAG 向けの出自が無効化されたリアルタイムの認知キャッシュ — 一度理解を構築すれば、どこでも再利用でき、新鮮さを保ちます。 · GitHub
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
Vectorlink-Labs
/
コアレント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
28 コミット 28 コミット フォルダーとファイル
.github/ ワークフロー .git

ハブ/ ワークフロー ベンチマーク/ news-n605 ベンチマーク/ news-n605 ブランド ブランドの例 例 統合/ langchain-coalent 統合/ langchain-coalent src/ coalent src/ coalent テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md REGISTRY-PUBLISH.md REGISTRY-PUBLISH.md UPGRADE-0.2-to-0.3.md UPGRADE-0.2-to-0.3.md UPGRADE-0.3-to-0.4.md UPGRADE-0.3-to-0.4.md UPGRADE-0.4-to-0.5.md UPGRADE-0.4-to-0.5.md UPGRADE-0.5-to-0.6.md UPGRADE-0.5-to-0.6.md pyproject.toml pyproject.toml server.json server.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントと RAG のための出自が無効化されたリアルタイムのコンテキスト。
一度理解を深めてください。あらゆる場所で再利用してください。自動的に新鮮な状態に保ちます。
📖 ドキュメント · coalent.ai · 💬 Discord
クイックスタート ·
v0.6 の新機能 ·
門梯子・
自分のスタックを持参してください ·
MCP・
ラングチェーン ·
ベンチマーク ·
CLI
エージェントは通話のたびに同じソースを再読み取りします。そして、ソースが変更された瞬間、キャッシュされたすべての回答は暗黙のうちに間違っています。
Coalent は理解を一度構築し、クエリの意味に基づいてキャッシュし、基礎となるソースが変更された瞬間にそれを外科的に無効にします。すべてを読み直すのと同じくらい正しく、わずかなコストで、しかも古くなることはありません。
すべてのコンテキスト層は 3 つのことをトレードオフすることを余儀なくされます。 Coalent は、次の 3 つすべてを一度に保持できるように構築されています。
🧠 断片的な理解ではなく、抽出的な理解。 LLM が抽出した、クエリに依存しないアトミックでソースに基づいたクレームのセットをキャッシュし、すべての数値と事実を保持します。そのため、キャッシュされた 1 つのユニットが、後のさまざまな質問に答えます。生の証拠は各ユニットで保持されるため、クエリを覆い隠すヒットは、薄い回答ではなく検索に戻ります。
♻️ クエリ、エージェント、ドキュメント全体で再利用します。クエリ平均をキーとしたセマンティック キャッシュ

ing : もう一度尋ねるか、別のエージェントに尋ねると、好評です。クロスユニット リコールは、余分な LLM コールをゼロにしながら、証拠が文書にまたがるマルチホップの質問に答えるために、ユニット間でクレームをプールします。
🌿 産地から見て新鮮です。すべてのユニットは、使用した正確な情報源を記憶します。変更すると、それを実際に使用したユニットだけが、正確に、自動的に、そしてゆっくりと古くなります。
Coalent は検索を超えて、あらゆる検索機能 (ベクター DB、ハイブリッド検索、GraphRAG、ツール、API) を提供します。これは、別の取得者ではなく、新鮮さと再利用のレイヤーです。GraphRAG のグラフ全体の事前構築の負荷とは意図的に反対です。軽量で独立したユニットであり、クエリが実際にユニットを必要とする場合にのみ遅延して構築され、単一のユニットをダーティにすることで更新されます (グラフの手術はありません)。
v0.6 の新機能 — プール読み取りパス ( read_path="pool" ): すべての読み取りは、トークン予算が設定され、グローバルにランク付けされたフレッシュクレーム プールにサービスを提供します。 605 の質問のニュース ベンチマークで測定 (厳格な採点): 981 コンテキスト トークンで 0.731 の精度 — トークンの最大 25% 少ない時点でナイーブ トップ 9 (0.711 @ 1,311) と一致し、トークンの最大 43% 少ない時点でナイーブの最高の測定ポイント (トップ 12: 0.731 @ 1,729) と一致します。さらに、デフォルトのオフの動作スタック (残留スパン → 拒否フォールバック → 追加のみの修復 → クエリ キー) があり、同じストアで拒否が -33%、+3.1 ポイントで測定されました。すべてオプトイン。デフォルトの読み取りパスは、v0.5 の動作から変更されていません。 「新機能」を参照してください。
v0.6.1 の新機能 — MCP サーバー: coalent-mcp は、クロード コード、カーソル、または任意の MCP クライアントから 1 行離れた場所にキャッシュを配置し (クロード コード / カーソルから使用します)、langchain-coalent は既存の LangChain スタックをキャッシュの基板にします。どちらも添加物のみ。
pip install coalent # コアには必要な依存関係がありません
クイックスタート
そのまま実行 — StubSynthesizer には API キーが必要ないため、10 秒でループを感じることができます。

coalent import SemanticCache 、 InMemoryRetriever 、 StubSynthesizer から
# 1. 任意のレトリーバー — ベクター DB、ツール、API。 (ここではデモ用にメモリ内にあります。)
レトリーバー = InMemoryRetriever ()
レトリーバー。 add ( "confluence:hr" , "休暇ポリシー: 年間 21 日の年次休暇。" )
# 2. キャッシュを構築します。以下の StubSynthesizer を実際の LLM と交換してください。
キャッシュ = SemanticCache (retriever , StubSynthesizer ())
#3. 尋ねる。最初の呼び出しで理解を深め、それをキャッシュします。次は温かいヒットです。
結果 = キャッシュ 。 get ( "当社の休暇ポリシーは何ですか?" )
print (結果 . コンテキスト [「理解」])
print ( result .cache_hit ) # False (コールド) -> 次の呼び出しで True
#4. ソースが変更されましたか?それを使用したユニットだけが外科的に古くなります。
キャッシュ。 source_changed ( "confluence:hr" , text = "ポリシーを離れる: 現在 25 日間。" )
# 次に一致する読み取りでは、その 1 つのユニットだけを遅延的に再構築します
実際のモデルに接続します。テキスト入力/テキスト出力 LLM はすべて機能します。 v0.4 では、シンセサイザーはデフォルトで抽出的理解を構築し (すべての事実を保持するクエリに依存しないアトミック クレーム)、キャッシュはユニット間の呼び出しを行います (両方とも自動的にオンになります)。
coalent インポートから SemanticCache 、 LLMSynthesizer 、 OpenAIProvider 、 OpenAIEmbedder
キャッシュ = SemanticCache (
レトリーバー、
LLMSynthesizer ( OpenAIProvider ()、model = "gpt-4o-mini" )、 # デフォルトで extract=True (v0.4)
embedder = OpenAIEmbedder (), # MEANING によるクエリの一致 (実際の使用を推奨)
)
# ドキュメント間でマルチホップしますか?リコールはすでにオンになっています。トリガーを上げてユニットをブリッジします。
# SemanticCache(レトリバー、シンセ、エンベッダー=...、リコール_しきい値=0.7)
v0.6 プール読み取りパス - オプトインすると、すべての読み取りが、1 つのルーティングされたユニットではなく、予算が詰め込まれたグローバルにランク付けされたフレッシュクレーム プールを提供します。アトリビューションは、各ユニットをマッピングする 3 行の pool_header 呼び出し可能コードです。

[タイトル |ソース |独自のコーパス メタデータからの日付]。これは測定されたゴールデン パスです。605 質問のニュース ベンチマーク (厳格な採点) では、裸の組み込みヘッダーの精度は 0.68 でしたが、この呼び出し可能な同じストア、同じクエリの精度は 0.73 でした。
DOC_META = { # アーティファクト ID をキーとしたコーパス メタデータ
"docs:azure-refresh" : { "title" : "Azure リージョンの更新" 、 "source" : "CloudWire" 、 "date" : "2026-05-02" },
}
def pool_header (unit) -> str : # [タイトル | ヘッダー]ソース |日付] ゴールデン パス — 3 行
メタ = DOC_META 。 get (ユニット . 証拠 [0]. artifact_id if ユニット . 証拠 else "" )
return f"[ { メタ [ 'タイトル' ] } | { メタ [ 'ソース' ] } | { メタ [ '日付' ] } ]" if メタ else f"[ソース: { ユニット . id } ]"
キャッシュ = SemanticCache ( レトリーバー 、 シンセサイザー 、 エンベッダー = OpenAIEmbedder ()、
read_path = "プール" 、 pool_header = プールヘッダー )
結果 = キャッシュ 。 get ( "どのリージョンが更新を取得しましたか?" )
結果。 context [ "pool" ] # パックされた属性付きクレーム ペイロード — 応答モデルに渡します
拒否ループを含む実行可能な API キーなしのデモ: example/pool_read_path.py 。
クロードコード/カーソル(MCP)から使用します
coalent-mcp は、Coalent キャッシュから任意の MCP クライアントに、新しい属性付きファクトを提供します。
そして事実は、その情報源が変更された瞬間に無効になります。配線するための 1 つの線
クロードコード:
pip インストール " coalent[mcp,openai] "
クロード mcp add coalent -- coalent-mcp --cache-factory my_cache:build
(カーソル / クロード デスクトップ / 任意の MCP クライアント: 同じ coalent-mcp ... コマンドを登録します
MCP 構成)
独自のキャッシュを使用する ( --cache-factory module:function ) — プライマリ モード。あなたの
ファクトリーは完全に構築された SemanticCache を返します: ベクター DB、エンベッダー、
LLM、すべてのノブ。サーバーはプロトコル グルーのみを追加します。グルーは追加するために測定されます。
品質損失ゼロ：事実

ory モードはライブラリ独自のベンチマーク結果を再現しました
バイト同一 (n=605 のニュースから抽出された 100 問の検証実行で 0.710
ベンチマーク — 同一の CI、100/100 サーブ、98/100 応答ペイロードがバイトと等しい
ライブラリを実行します）。
# my_cache.py — 起動したディレクトリからインポート可能
coalent import から ( SemanticCache 、 LLMSynthesizer 、 OpenAIProvider 、
OpenAIEmbedder 、 SQLiteCognitionStore )
def build() -> SemanticCache :
SemanticCache を返す (
my_vector_retriever , # あなたのベクター DB/レトリーバー
LLMSynthesizer ( OpenAIProvider ())、 # あなたの合成モデル
embedder = OpenAIEmbedder (), # あなたのエンベッダー
read_path = "プール" ,
Resident_spans = True 、query_keys = True 、 # 動作スタック、これまでと同様にオプトイン
pool_header = my_metadata_header , # [タイトル |ソース | date] — 測定されたゴールデン パス
store = SQLiteCognitionStore ( "kb.db" ), # 永続性もあなたのものです
)
ここでの鮮度はシグナル駆動です。取り込みパイプラインはsource_changedツールを呼び出します。
文書が変更され、影響を受けた事実が直ちに無効になる場合。 (追加
--watch DIR をファクトリと一緒に指定すると、ファイル編集時にも起動されます (無効化のみ)。それ
インデックスに取り込まれることはなく、アーティファクト ID が一致する場合にのみ一致します。
ウォッチ相対パス)。
ゼロ構成フォルダー モード ( --watch DIR ) — デモ ウェッジ。のフォルダーを指定します。
ドキュメントを参照すると、推奨される v0.6 デプロイメント (プール パス、残りのスパン、クエリ キー、
SQLite 永続化、自動 [パス |変更日] 帰属) コードがまったくありません:
クロード mcp add coalent --env OPENAI_API_KEY= $OPENAI_API_KEY -- coalent-mcp --watch ./docs
読み取りごとに、配信前に監視対象ファイル (mtime + コンテンツ ハッシュ) を再スキャンします。
ファイルを保存した後に古い応答が返され、手つかずのフォルダーが完全にウォーム状態で再起動されます。
正直な数字: オン

同じ 100 問の検証を実行し、フォルダー モードでスコアを取得
工場出荷時に構築されたキャッシュの場合は 0.46 対 0.71 (拒否数 40 対 18) — キャッシュの測定コスト
一般的な段落チャンカーとオンデマンドのキーホール ビルド。フレッシュネスループを感じるために使用してください
すぐに;本番品質を実現するために独自のスタックを導入してください。一つの政権メモ: 質問
追加したばかりのファイルについては、読み取りがトリガーされるまでウォーム キャッシュから正直に拒否できます。
そのファイルの最初のビルドは拒否であり、決して古い答えや間違った答えではありません。
多くのエージェントに対して 1 つの共有キャッシュ ( --transport http )。
COALENT_MCP_TOKEN= < シークレット > coalent-mcp --cache-factory my_cache:build --transport http --port 8765
1 つの長期存続プロセス、多数の同時 MCP クライアント、1 つの共有キャッシュ - 共有
複合化、ストア競合なし (検証: 2 つの同時クライアントがシーケンシャル クライアントと一致しました)
20 回の読み取りすべてで参照、重複ビルドはゼロ）。 COALENT_MCP_TOKEN が設定されている場合、
リクエストには Authorization が含まれている必要があります: Bearer <token> — ローカルホストまたは信頼できるネットワークをバインドします。
stdio の当然の帰結: 各 stdio の起動は独自のプロセスであるため、2 つのアプリを決してポイントしないでください。
同じ --store パス — HTTP モードが共有キャッシュの答えです。
7 つのツール: get_context(query, Budget?) → 属性付きの予算が詰め込まれたツール
ペイロード + read_id · report_refusal(read_id) / report_success(read_id) →
MCP 上の動作修復ループ · source_chang

[切り捨てられた]

## Original Extract

Real-time, provenance-invalidated cognitive cache for AI agents and RAG — build understanding once, reuse it everywhere, keep it fresh. - Vectorlink-Labs/coalent

GitHub - Vectorlink-Labs/coalent: Real-time, provenance-invalidated cognitive cache for AI agents and RAG — build understanding once, reuse it everywhere, keep it fresh. · GitHub
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
Vectorlink-Labs
/
coalent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
28 Commits 28 Commits Folders and files
.github/ workflows .github/ workflows benchmarks/ news-n605 benchmarks/ news-n605 brand brand examples examples integrations/ langchain-coalent integrations/ langchain-coalent src/ coalent src/ coalent tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md REGISTRY-PUBLISH.md REGISTRY-PUBLISH.md UPGRADE-0.2-to-0.3.md UPGRADE-0.2-to-0.3.md UPGRADE-0.3-to-0.4.md UPGRADE-0.3-to-0.4.md UPGRADE-0.4-to-0.5.md UPGRADE-0.4-to-0.5.md UPGRADE-0.5-to-0.6.md UPGRADE-0.5-to-0.6.md pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Real-time, provenance-invalidated context for AI agents & RAG.
Build understanding once. Reuse it everywhere. Keep it fresh — automatically.
📖 Documentation · coalent.ai · 💬 Discord
Quickstart ·
What's new in v0.6 ·
Gate ladder ·
Bring your own stack ·
MCP ·
LangChain ·
Benchmark ·
CLI
Your agent re-reads the same sources on every call — and the moment a source changes, every cached answer is silently wrong.
Coalent builds the understanding once, caches it by what the query means , and invalidates it surgically the instant an underlying source changes. As correct as re-reading everything, at a fraction of the cost — and never stale.
Every context layer is forced to trade off three things. Coalent is built to hold all three at once:
🧠 Extractive understanding, not chunks. It caches a query-independent set of atomic, source-grounded claims your LLM extracted — keeping every number and fact — so one cached unit answers many different later questions. The raw evidence is retained with each unit, so a hit that under-covers a query falls back to retrieval instead of answering thin.
♻️ Reuse across queries, agents — and documents. A semantic cache keyed by query meaning : ask again, or from another agent, and it's a warm hit. Cross-unit recall pools claims across units to answer multi-hop questions whose evidence spans documents — at zero extra LLM calls .
🌿 Fresh by provenance. Every unit remembers the exact sources it used. When one changes, only the units that actually used it go stale — precisely, automatically, and lazily.
Coalent sits above retrieval — bring any retriever (vector DB, hybrid search, GraphRAG, tools, APIs). It's the freshness-and-reuse layer, not another retriever — deliberately the opposite of GraphRAG's build-the-whole-graph-upfront tax: lightweight, independent units, built lazily only when a query actually needs one , and refreshed by dirtying a single unit (no graph surgery).
New in v0.6 — the pool read path ( read_path="pool" ): every read serves the token-budgeted, globally ranked fresh-claim pool. Measured on a 605-question news benchmark (strict grading): 0.731 accuracy @ 981 context tokens — matching naive top-9 (0.711 @ 1,311) at ~25% fewer tokens , and naive's best measured point (top-12: 0.731 @ 1,729) at ~43% fewer . Plus a default-OFF behavioral stack — residual spans → refusal fallback → append-only repair → query keys — measured at −33% refusals and +3.1 pts on the same store. All opt-in; the default read path is unchanged v0.5 behavior. See What's new .
New in v0.6.1 — the MCP server : coalent-mcp puts the cache one line away from Claude Code, Cursor, or any MCP client ( Use it from Claude Code / Cursor ), and langchain-coalent makes your existing LangChain stack the cache's substrate. Both additive-only.
pip install coalent # the core has zero required dependencies
Quickstart
Runs as-is — StubSynthesizer needs no API key, so you can feel the loop in ten seconds:
from coalent import SemanticCache , InMemoryRetriever , StubSynthesizer
# 1. Any retriever — a vector DB, a tool, an API. (In-memory here for the demo.)
retriever = InMemoryRetriever ()
retriever . add ( "confluence:hr" , "Leave policy: 21 days of annual leave per year." )
# 2. Build the cache. Swap StubSynthesizer for a real LLM below.
cache = SemanticCache ( retriever , StubSynthesizer ())
# 3. Ask. The first call builds understanding and caches it; the next is a warm hit.
result = cache . get ( "what is our leave policy?" )
print ( result . context [ "understanding" ])
print ( result . cache_hit ) # False (cold) -> True on the next call
# 4. A source changed? Only the units that used it go stale — surgically.
cache . source_changed ( "confluence:hr" , text = "Leave policy: now 25 days." )
# the next matching read rebuilds just that one unit, lazily
Wire in a real model — any text-in / text-out LLM works. In v0.4 the synthesizer builds extractive understanding by default (query-independent atomic claims that keep every fact), and the cache does cross-unit recall — both on automatically:
from coalent import SemanticCache , LLMSynthesizer , OpenAIProvider , OpenAIEmbedder
cache = SemanticCache (
retriever ,
LLMSynthesizer ( OpenAIProvider (), model = "gpt-4o-mini" ), # extract=True by default (v0.4)
embedder = OpenAIEmbedder (), # match queries by MEANING (recommended for real use)
)
# Multi-hop across documents? recall is already on; raise its trigger to bridge units:
# SemanticCache(retriever, synth, embedder=..., recall_threshold=0.7)
The v0.6 pool read path — opt in, and every read serves the budget-packed, globally ranked fresh-claim pool instead of one routed unit. Attribution is the one thing to wire: a 3-line pool_header callable mapping each unit to [title | source | date] from your own corpus metadata. This is the measured golden path — on a 605-question news benchmark (strict grading), 0.68 accuracy with the bare built-in header vs 0.73 with this callable, same store, same queries:
DOC_META = { # your corpus metadata, keyed by artifact id
"docs:azure-refresh" : { "title" : "Azure region refresh" , "source" : "CloudWire" , "date" : "2026-05-02" },
}
def pool_header ( unit ) -> str : # the [title | source | date] golden path — 3 lines
meta = DOC_META . get ( unit . evidence [ 0 ]. artifact_id if unit . evidence else "" )
return f"[ { meta [ 'title' ] } | { meta [ 'source' ] } | { meta [ 'date' ] } ]" if meta else f"[source: { unit . id } ]"
cache = SemanticCache ( retriever , synthesizer , embedder = OpenAIEmbedder (),
read_path = "pool" , pool_header = pool_header )
result = cache . get ( "which regions got the refresh?" )
result . context [ "pool" ] # the packed, attributed claim payload — hand it to your answer model
Runnable no-API-key demo, including the refusal loop: examples/pool_read_path.py .
Use it from Claude Code / Cursor (MCP)
coalent-mcp serves fresh, attributed facts from a Coalent cache to any MCP client —
and the facts are invalidated the instant their source changes. One line to wire it into
Claude Code:
pip install " coalent[mcp,openai] "
claude mcp add coalent -- coalent-mcp --cache-factory my_cache:build
(Cursor / Claude Desktop / any MCP client: register the same coalent-mcp ... command in
its MCP config.)
Bring your own cache ( --cache-factory module:function ) — the primary mode. Your
factory returns a fully constructed SemanticCache : your vector DB, your embedder, your
LLM, every knob. The server adds protocol glue only — and the glue is measured to add
zero quality loss : factory mode reproduced the library's own benchmark result
byte-identically (0.710 on a 100-question validation run drawn from our n=605 news
benchmark — identical CIs, 100/100 serves, 98/100 answer payloads byte-equal to the
library run).
# my_cache.py — importable from the directory you launch in
from coalent import ( SemanticCache , LLMSynthesizer , OpenAIProvider ,
OpenAIEmbedder , SQLiteCognitionStore )
def build () -> SemanticCache :
return SemanticCache (
my_vector_retriever , # YOUR vector DB / retriever
LLMSynthesizer ( OpenAIProvider ()), # YOUR synthesis model
embedder = OpenAIEmbedder (), # YOUR embedder
read_path = "pool" ,
residual_spans = True , query_keys = True , # the behavioral stack, opt-in as ever
pool_header = my_metadata_header , # [title | source | date] — the measured golden path
store = SQLiteCognitionStore ( "kb.db" ), # persistence is yours too
)
Freshness here is signal-driven: your ingestion pipeline calls the source_changed tool
when a document changes and the affected facts invalidate immediately. (Adding
--watch DIR alongside the factory also fires it on file edits — invalidation only; it
never ingests into your index, and it matches only when your artifact ids equal the
watch-relative paths.)
Zero-config folder mode ( --watch DIR ) — the demo wedge. Point it at a folder of
docs and you get the recommended v0.6 deployment (pool path, residual spans, query keys,
SQLite persistence, automatic [path | modified date] attribution) with no code at all:
claude mcp add coalent --env OPENAI_API_KEY= $OPENAI_API_KEY -- coalent-mcp --watch ./docs
Every read rescans the watched files (mtime + content hash) before serving — you cannot
get a stale answer after saving a file — and an untouched folder restarts fully warm.
The honest number: on the same 100-question validation run, folder mode scored
0.46 vs 0.71 for a factory-built cache (40 vs 18 refusals) — the measured cost of the
generic paragraph chunker and on-demand keyhole builds. Use it to feel the freshness loop
in a minute; bring your own stack for production quality. One regime note: a question
about a just-added file can honestly refuse from a warm cache until a read triggers
that file's first build — a refusal, never a stale or wrong answer.
One shared cache for many agents ( --transport http ).
COALENT_MCP_TOKEN= < secret > coalent-mcp --cache-factory my_cache:build --transport http --port 8765
One long-lived process, many concurrent MCP clients, ONE shared cache — shared
compounding, no store races (validated: two concurrent clients matched the sequential
reference on all 20 reads, zero duplicate builds). When COALENT_MCP_TOKEN is set, every
request must carry Authorization: Bearer <token> — bind localhost or trusted networks.
Corollary for stdio: each stdio launch is its own process, so never point two apps at the
same --store path — HTTP mode is the shared-cache answer.
The seven tools: get_context(query, budget?) → the attributed, budget-packed
payload + a read_id · report_refusal(read_id) / report_success(read_id) → the
behavioral repair loop over MCP · source_chang

[truncated]
