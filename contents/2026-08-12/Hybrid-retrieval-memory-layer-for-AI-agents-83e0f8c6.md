---
source: "https://github.com/deepmemteam/deepmem"
hn_url: "https://news.ycombinator.com/item?id=49270630"
title: "Hybrid-retrieval memory layer for AI agents"
article_title: "GitHub - deepmemteam/deepmem: Drop-in AI memory layer with 2x faster response and 10x lower cost. Fully compatible with Mem0 API. Migrate in 5 minutes without any code changes. Self-host for free. · GitHub"
author: "qikouki"
captured_at: "2026-08-12T11:39:44Z"
capture_tool: "hn-digest"
hn_id: 49270630
score: 2
comments: 0
posted_at: "2026-08-12T11:14:59Z"
tags:
  - hacker-news
  - translated
---

# Hybrid-retrieval memory layer for AI agents

- HN: [49270630](https://news.ycombinator.com/item?id=49270630)
- Source: [github.com](https://github.com/deepmemteam/deepmem)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T11:14:59Z

## Translation

タイトル: AI エージェント用のハイブリッド検索メモリ層
記事のタイトル: GitHub - deepmemteam/deepmem: 2 倍の高速応答と 10 倍のコストを備えたドロップイン AI メモリ レイヤー。 Mem0 APIと完全互換。コードを変更することなく 5 分で移行できます。無料のセルフホスト。 · GitHub
説明: 2 倍の高速応答と 10 倍のコストを備えたドロップイン AI メモリ レイヤー。 Mem0 APIと完全互換。コードを変更することなく 5 分で移行できます。無料のセルフホスト。 - ディープメムチーム/ディープメム

記事本文:
GitHub - deepmemteam/deepmem: 2 倍の高速応答と 10 倍のコストを備えたドロップイン AI メモリ レイヤー。 Mem0 APIと完全互換。コードを変更することなく 5 分で移行できます。無料のセルフホスト。 · GitHub
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
ディープメンチーム
/
ディープメモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット ベンチマーク ベンチマーク deepmem deepmem スクリプト スクリプト サーバー サーバー テスト tes

ts .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore Dockerfile Dockerfile ライセンス ライセンス README.md README.md config.json.example config.json.exampleデモ.gif デモ.gif docker-compose.yml docker-compose.yml pytest.ini pytest.ini 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
2 倍の高速応答と 10 倍のコストを実現するドロップイン AI メモリ層。
Mem0 APIと完全互換。 1 つのインポート行で 5 分で移行します。
自己ホスト可能。認証も支払いもロックインもありません。または、 deepmem.dev にあるマネージド クラウドを使用します。
クラウド
·
セルフホスト
·
ベンチマーク
·
それらを再現する
Mem0 から 1 行で移行します - 同じ MemoryClient 、同じメソッド シグネチャ:
# 前 - Mem0
mem0からMemoryClientをインポート
client = MemoryClient ( api_key = "m0-..." )
# 後 - DeepMem (インポートのみが変更されます)
deepmemインポートMemoryClientから
client = MemoryClient ( api_key = "dm_live-..." ) # deepmem.dev でキーを取得します
pip インストール deepmem クライアント
会話を検索可能な長期記憶に変える: FastAPI HTTP API
Qdrant ベクトル ストアの前面、LLM ファクト抽出、ハイブリッド検索を使用
(ベクター + BM25 + エンティティ ブースト + 時間減衰)、セマンティック キャッシュ、非同期バッチ処理
蒸留、GDPR 管理、内蔵 MCP サーバー。オープンで走ります
モード - API キーやユーザー登録は不要 - そのため独自にデプロイできます
エージェントは数分で完了します。マルチテナントの分離は、
リクエスト本文。
セルフホストをしたくないですか? DeepMem Cloud は、のマネージド バージョンです。
deepmem.dev にあるこの正確なエンジン - 同じ API、いいえ
インフラ。サインアップし、キー ( dm_live_... ) を取得し、ベース URL を指定します。
https://deepmem.dev 、完了しました。クラウドとオープンソースサーバーが語る
同じ Mem0 互換 API なので、クライアント コードは同一です

アル。
すでに Mem0 を使用していますか? 1 行で DeepMem クラウドに切り替えます。 deepmemクライアント
パッケージは mem0.MemoryClient をミラーリングします - 同じクラス名、同じメソッド シグネチャ、
同じ filters={"user_id": ...} スタイル - インポート後のすべてがそのまま残ります
手付かずの。
pip インストール deepmem クライアント
# 前 (Mem0)
mem0からMemoryClientをインポート
client = MemoryClient ( api_key = "m0-..." )
クライアント。 add ( メッセージ , user_id = "alex" )
クライアント。 search ( "Alex は何を料理できますか?" , filters = { "user_id" : "alex" })
# 後 (DeepMem クラウド) - インポート行を 1 つ変更します
deepmemインポートMemoryClientから
client = MemoryClient ( api_key = "dm_live_..." ) # https://deepmem.dev のキー
クライアント。 add (messages , user_id = "alex" ) # 同一の呼び出し
クライアント。 search ( "Alex は何を料理できますか?" , filters = { "user_id" : "alex" })
行動メモ
add(infer=True) (デフォルト) は DeepMem クラウドでは非同期です -
pending=True を results=[] で返し、抽出されたファクトがいくつか見つかります
数秒後。 (Mem0 クラウドの追加も非同期です - PENDING が返されます。)
infer=False は、すぐに検索できる同期生テキスト ストレージの場合です。
グラフ関係なし - DeepMem はハイブリッド ベクトル取得 (ベクトル + BM25) を使用します
時間減衰)、したがって関係は常に [] になります。 Mem0 のグラフ機能はそうではありません。
複製されました。
リセットは異なります - Mem0 はアカウント全体です。 DeepMem は user_id ごとにあります
確認ガード。
DeepMem Cloud は、すべての有料レベルで Mem0 よりも 10 倍安い - 同じ形状
プランの料金は10分の1。
代わりに自己ホストを使用すると $0 - 独自の LLM/埋め込み料金のみを支払います
プロバイダー (プラン料金に加えて Mem0 が請求するのと同じ LLM コスト)、
メモリサービスのマークアップ。バッチ蒸留では LLM コールも最大 80% 削減されるため、
プロバイダーの請求額は、メッセージごとのエクストラクターよりも少額になります。
計画と制限: deepmem.dev · mem0.ai 。
厳選された見出しはありません。スクリプト

2番目のワークロードの出荷
/benchmarks - 自分で実行します。これが私たちのものです
測定値とそれを生成した正確な構成:
¹ GTX 1070 GPU (2016 年頃) 上の BGE-M3、ローカル ファイル Qdrant、infer=False 、
100 オペレーション、同時実行性 1. ² ローカル ファイル Qdrant I/O によって支配される - Qdrant
サーバーはこれを大幅にカットします。 ³ Mem0 には raw ストア モードがありません。常に実行される追加
LLM 抽出なので、この行は完全一致ではありません。
セルフホスト型では、本質的な遅延が発生します。インターネット RTT はありません。
埋め込み者、あなたのQdrant。古いコンシューマ GPU での p50 検索は 73 ミリ秒。
DeepMem クラウドは検索 p50 で Mem0 クラウドを上回ります (643 ミリ秒対 653 ミリ秒)。
検索ごとに最大 2.3 倍の候補を返します (40 件中 195 件対 84 件のヒット)
クエリ）。
クラウドの遅延は RTT に支配されます - 両方のクラウド列が測定されました
中国本土からの代理人を通じて。実行間のジッターは最大±10%です。走る
/benchmarks を低 RTT の場所から取得
数字。
エージェント フレームワークには、永続的で取得可能なものが必要であることが再認識され続けています。
記憶。ホスト型オプションは通話ごとに請求され、データを他の人のサーバーに送信します。
雲。 DeepMem は自己ホスト可能な代替手段であり、同じ Mem0 型の API です。
ドロップインすることもできますが、エンベッダー、LLM キー、および
あなたの Qdrant - それを確認するためのコードがここにあります。
DeepMem は他の Mem0 代替手段とどのように比較されますか?ほとんどはホスト専用、または他人のベクター DB の上に重ねられたメモリです。 DeepMem は 3 つの機能を同時に組み合わせたものです。自己ホスト可能 (データはボックスに残り、LLM キーを超えると $0)、MCP ネイティブ (クロード デスクトップ/カーソルがツールとしてメモリを直接読み書き)、完全にオープンソースです。マネージド クラウドはまったく同じエンジンを実行するため、クラウドとセルフホストは 2 つの製品ではなく 1 つの API です。
ナレッジグラフではなく、ハイブリッド検索。検索ヒューズ ベクトルの類似性、
BM25 キーワード マッチ、エンティティ ブースト、および時間減衰スコア

してる。ありません
時間グラフ層。それが必要な場合は、Zep を見てください。
コードダンプではなく設定を保存します。大きなフェンスで囲まれたコードブロックは、
LLM を抽出する前に除去されるため、ストアには耐久性のある材料が充填されます。
貼り付けられた実装ではなく、ユーザー/プロジェクトのファクト。
BYOK、マルチプロバイダー。独自の LLM を持ち込みます (OpenAI / Anthropic / 任意)
OpenAI 互換エンドポイント）とエンベディング（BGE-M3 / Google / OpenAI 互換）。
MCPネイティブ。 MCP サーバーを同梱しているため、Claude Desktop / Cursor は読み取りと
思い出を直接書き込む。
3つの走り方。すべて同じ Mem0 互換 API を話します。
export DEEPMEM_API_KEY=dm_live_... # https://deepmem.dev から
カール https://deepmem.dev/v1/memories \
-H " 認可: Bearer $DEEPMEM_API_KEY " -H " コンテンツタイプ: application/json " \
-d ' {"messages":[{"role":"user","content":"私はパットです。リスボンに住んでいます。"}],"user_id":"pat","infer":false} '
カール https://deepmem.dev/v1/memories/search \
-H " 認可: Bearer $DEEPMEM_API_KEY " -H " コンテンツタイプ: application/json " \
-d ' {"query":"パットはどこに住んでいますか?","user_id":"パット"} '
2. Docker (1 つのコマンド)
cp .env.example .env # LLM キーを追加します
docker compose up --build # DeepMem (HTTP :8000 + MCP :8001) + Qdrant サイドカー
カール http://localhost:8000/health
または、公開されたイメージをプルします。
docker pull langdeepmem/deepmem:latest
docker run -p 8000:8000 -p 8001:8001 -e DEEPSEEK_API_KEY=sk-... langdeepmem/deepmem:latest
このイメージでは、 :8000 (HTTP) および :8001 (MCP) が公開されています。 Dockerfile
および docker-compose.yml は GPU バリアント (CUDA トーチ + BGE_DEVICE=cuda ) をカバーします。
BGE-M3 モデル ダウンロード オプション (HF ミラー、プロキシ、またはローカル マウント)。
git clone https://github.com/deepmemteam/deepmem.git && cd deepmem
pip install -r 要件.txt
cp .env.example .env # LLM キー + 埋め込み設定を追加します
python サーバー/start.py # HTTP :8000 + MCP :8001
書く

3 行で検索します。
httpx をインポートする
httpx 。 post ( "http://localhost:8000/v1/memories" ,
json = { "messages" :[{ "role" : "user" , "content" : "私はパットです、リスボンに住んでいます。" }]、
"user_id" : "パット" })
print ( httpx .post ( "http://localhost:8000/v1/memories/search" ,
json = { "クエリ" : "パットはどこに住んでいますか?" , "user_id" : "pat" })。 json ()[ "結果" ])
user_id はオプションです (デフォルトは "default" です)。異なる user_id を送信する
エンドユーザーを隔離するため。 infer: false は生のテキストをすぐに保存します
(テストしやすい);デフォルトの infer: true は、LLM ファクト抽出のキューを作成します。
コードではなく構成によるマルチプロバイダー。両方の層が環境変数をオンにします。
BGE_DEVICE=auto|cpu|cuda が利用可能な場合は GPU を選択し、それ以外の場合は CPU (強制的に CPU)
マルチプロセスの競合を避けるため、小さな VRAM カード上で使用されます)。 BYOK は
リクエストごとの LLM。
ハイブリッド検索。ベクトル類似度 + BM25 キーワード + エンティティ ブースト +
時間の減衰が 1 つのスコアに融合されます。オーバーフェッチ、再ランク付け、リターン。
非同期バッチ蒸留。書き込みキューは沈黙ウィンドウの背後にあり、
バッチで抽出 - メッセージごとの抽出よりも LLM 呼び出しが最大 80% 少なくなります。
セマンティック キャッシュ。追加/検索を繰り返すと類似性ゲート キャッシュにヒットし、
Qdrant の再埋め込みや再クエリを行わずに、キャッシュされたファクトを返します。
コードではなく設定を保存します。 extract_filter ストリップ 大型フェンス付き
LLM 抽出前にコード ブロックを作成するため、ストアは永続的なファクトで満たされます。
実装が貼り付けられていません。
MCPサーバー。 deepmem_write / deepmem_search / deepmem_delete ツール
Claude Desktop、Cursor、および任意の MCP クライアント用。
GDPR。保持期間付きの論理的な削除、ハード削除のリセット、SHA-256
ログ内の ID マスキング、移植性のためのエクスポート/インポート。
方法
パス
説明
投稿
/v1/思い出
メッセージを書きます。 LLM - ファクトの抽出 ( infer=false は生のデータを保存します)
投稿
/v1/思い出/検索
セマンティック検索 (ベクトル + BM25 + エンティティ + 時間減衰)
ゲット
/v1/m

思い出
user_id のすべてのリスト (ページ分割)
ゲット
/v1/思い出/{id}
IDで取得
置く
/v1/思い出/{id}
1 つの思い出のテキストを更新する
削除
/v1/思い出/{id}
1 つを論理的に削除します
削除
/v1/思い出
user_id のすべての論理的な削除 (GDPR)
ゲット
/v1/思い出/{id}/履歴
監査ログの追加/更新/削除
投稿
/v1/リセット
すべて + 履歴を物理的に削除します (confirm_user_id が必要です)
ゲット
/v1/エクスポート · POST /v1/インポート
ポータブルな JSON エクスポート/インポート
ゲット
/健康 · /準備完了
Liveness / Readiness プローブ
Agent_id / run_id オプションでスコープの書き込み/読み取り (Mem0 の 3 レベルを反映)
分離: ユーザー -> エージェント -> 実行)。インタラクティブなドキュメントは /docs にあります。
{
"mcpサーバー": {
"ディープメム" : {
"コマンド" : " Python " ,
"args" : [ "server/mcp_server.py " ],
"env" : { "DEEPMEMORY_BASE_URL" : " http://localhost:8000 " }
}
}
}
オープン モードには API キーは必要ありません。DEEPMEMORY_API_KEY は空のままです。
クライアント(HTTP/MCP)
-> FastAPI (:8000)
-> レート制限ミドルウェア (追加/検索時の IP トークン バケットごと)
-> SemanticCache.check (類似性ヒット時にキャッシュされたファクトを返す)
-> AsyncBatchDistiller.enqueue (POST /v1/memories、infer=true)
↳ Silence-window または max_batch が on_batch_ready をトリガーする
↳ VectorStore.process_batch (LLM 抽出 -> Qdrant 更新/挿入)
-> VectorStore.search (ベクター + BM25 + エンティティ + 時間減衰)

[切り捨てられた]

## Original Extract

Drop-in AI memory layer with 2x faster response and 10x lower cost. Fully compatible with Mem0 API. Migrate in 5 minutes without any code changes. Self-host for free. - deepmemteam/deepmem

GitHub - deepmemteam/deepmem: Drop-in AI memory layer with 2x faster response and 10x lower cost. Fully compatible with Mem0 API. Migrate in 5 minutes without any code changes. Self-host for free. · GitHub
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
deepmemteam
/
deepmem
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits benchmarks benchmarks deepmem deepmem scripts scripts server server tests tests .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md config.json.example config.json.example demo.gif demo.gif docker-compose.yml docker-compose.yml pytest.ini pytest.ini requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
Drop-in AI memory layer with 2× faster response and 10× lower cost.
Fully compatible with Mem0 API. Migrate in 5 minutes - one import line.
Self-hostable. No auth, no payment, no lock-in. Or use the managed cloud at deepmem.dev .
Cloud
·
Self-host
·
Benchmarks
·
Reproduce them
Migrate from Mem0 in one line - same MemoryClient , same method signatures:
# Before - Mem0
from mem0 import MemoryClient
client = MemoryClient ( api_key = "m0-..." )
# After - DeepMem (only the import changes)
from deepmem import MemoryClient
client = MemoryClient ( api_key = "dm_live-..." ) # get a key at deepmem.dev
pip install deepmem-client
Turn conversations into searchable long-term memory: a FastAPI HTTP API in
front of a Qdrant vector store, with LLM fact extraction, hybrid retrieval
(vector + BM25 + entity boost + time-decay), semantic caching, async batched
distillation, GDPR controls, and a built-in MCP server. It runs in open
mode - no API key, no user registration - so you can deploy it for your own
agents in minutes. Multi-tenant isolation is driven by user_id in the
request body.
Prefer not to self-host? DeepMem Cloud is the managed version of
this exact engine at deepmem.dev - same API, no
infra. Sign up, grab a key ( dm_live_... ), point your base URL at
https://deepmem.dev , done. The cloud and the open-source server speak the
same Mem0-compatible API, so client code is identical.
Already using Mem0? Switch to DeepMem cloud in one line. The deepmem-client
package mirrors mem0.MemoryClient - same class name, same method signatures,
same filters={"user_id": ...} style - so everything after the import stays
untouched.
pip install deepmem-client
# before (Mem0)
from mem0 import MemoryClient
client = MemoryClient ( api_key = "m0-..." )
client . add ( messages , user_id = "alex" )
client . search ( "What can Alex cook?" , filters = { "user_id" : "alex" })
# after (DeepMem cloud) - change one import line
from deepmem import MemoryClient
client = MemoryClient ( api_key = "dm_live_..." ) # key at https://deepmem.dev
client . add ( messages , user_id = "alex" ) # identical calls
client . search ( "What can Alex cook?" , filters = { "user_id" : "alex" })
Behavioral notes
add(infer=True) (the default) is asynchronous on DeepMem cloud - it
returns pending=True with results=[] and extracted facts land a few
seconds later. (Mem0 cloud's add is async too - it returns PENDING .) Pass
infer=False for synchronous raw-text storage that's immediately searchable.
No graph relations - DeepMem uses hybrid vector retrieval (vector + BM25
time-decay), so relations is always [] . Mem0's graph features aren't
replicated.
reset differs - Mem0's is account-wide; DeepMem's is per- user_id with
a confirm guard.
DeepMem Cloud is 10x cheaper than Mem0 at every paid tier - the same shape
of plans, a tenth of the price.
Self-host instead and it's $0 - you pay only your own LLM/embedding
provider (the same LLM cost Mem0 charges on top of its plan price), with no
memory-service markup. Batched distillation also cuts LLM calls ~80%, so even
your provider bill is smaller than per-message extractors.
Plans and limits: deepmem.dev · mem0.ai .
No cherry-picked headline. The scripts and workload ship in
/benchmarks - run them yourself. Here's what we
measured and the exact config that produced it:
¹ BGE-M3 on a GTX 1070 GPU (2016-era), local file Qdrant, infer=False ,
100 ops, concurrency 1. ² Dominated by local-file Qdrant I/O - a Qdrant
server cuts this sharply. ³ Mem0 has no raw-store mode; add always runs
LLM extraction, so this row isn't apples-to-apples.
Self-hosted is where intrinsic latency lives - no internet RTT, your
embedder, your Qdrant. 73 ms p50 search on an old consumer GPU.
DeepMem cloud beats Mem0 cloud on search p50 (643 ms vs 653 ms) and
returns ~2.3x more candidates per search (195 vs 84 hits across 40
queries).
Cloud latency is RTT-dominated - both cloud columns were measured
through a proxy from mainland China; run-to-run jitter is ~±10%. Run
/benchmarks from a low-RTT location for your own
numbers.
Agent frameworks keep re-discovering that they need persistent, retrievable
memory. The hosted options bill per call and send your data to someone else's
cloud. DeepMem is the self-hostable alternative: the same Mem0-shaped API you
can drop in, but it runs on your box, with your embedder, your LLM key, and
your Qdrant - and the code is right here to verify it.
How does DeepMem compare to other Mem0 alternatives? Most are hosted-only or layer memory on top of someone else's vector DB. DeepMem combines three things at once: it's self-hostable (your data stays on your box - $0 beyond your own LLM key), MCP-native (Claude Desktop / Cursor read and write memories directly as tools), and fully open-source - and the managed cloud runs the exact same engine, so cloud and self-host are one API, not two products.
Hybrid retrieval, not a knowledge graph. Search fuses vector similarity,
BM25 keyword match, entity boost, and time-decay scoring. There is no
temporal graph layer; if that's what you need, look at Zep.
Stores preferences, not code dumps. Large fenced code blocks are
stripped before LLM extraction, so the store fills with durable
user/project facts instead of pasted implementations.
BYOK, multi-provider. Bring your own LLM (OpenAI / Anthropic / any
OpenAI-compatible endpoint) and embedding (BGE-M3 / Google / OpenAI-compatible).
MCP-native. Ships an MCP server so Claude Desktop / Cursor can read and
write memories directly.
Three ways to run. All speak the same Mem0-compatible API.
export DEEPMEM_API_KEY=dm_live_... # from https://deepmem.dev
curl https://deepmem.dev/v1/memories \
-H " Authorization: Bearer $DEEPMEM_API_KEY " -H " Content-Type: application/json " \
-d ' {"messages":[{"role":"user","content":"I am Pat, I live in Lisbon."}],"user_id":"pat","infer":false} '
curl https://deepmem.dev/v1/memories/search \
-H " Authorization: Bearer $DEEPMEM_API_KEY " -H " Content-Type: application/json " \
-d ' {"query":"Where does Pat live?","user_id":"pat"} '
2. Docker (one command)
cp .env.example .env # add an LLM key
docker compose up --build # DeepMem (HTTP :8000 + MCP :8001) + Qdrant sidecar
curl http://localhost:8000/health
Or pull the published image:
docker pull langdeepmem/deepmem:latest
docker run -p 8000:8000 -p 8001:8001 -e DEEPSEEK_API_KEY=sk-... langdeepmem/deepmem:latest
The image exposes :8000 (HTTP) and :8001 (MCP) . The Dockerfile
and docker-compose.yml cover the GPU variant (CUDA torch + BGE_DEVICE=cuda )
and BGE-M3 model-download options (HF mirror, proxy, or local mount).
git clone https://github.com/deepmemteam/deepmem.git && cd deepmem
pip install -r requirements.txt
cp .env.example .env # add an LLM key + embedder config
python server/start.py # HTTP :8000 + MCP :8001
Write and search in three lines:
import httpx
httpx . post ( "http://localhost:8000/v1/memories" ,
json = { "messages" :[{ "role" : "user" , "content" : "I'm Pat, I live in Lisbon." }],
"user_id" : "pat" })
print ( httpx . post ( "http://localhost:8000/v1/memories/search" ,
json = { "query" : "Where does Pat live?" , "user_id" : "pat" }). json ()[ "results" ])
user_id is optional (defaults to "default" ); send different user_id s
to isolate end-users. infer: false stores raw text immediately
(test-friendly); the default infer: true queues for LLM fact extraction.
Multi-provider by config, not code. Both layers switch on env vars:
BGE_DEVICE=auto|cpu|cuda picks GPU when available, else CPU (force cpu
on small-VRAM cards to avoid multi-process contention). BYOK overrides the
LLM per-request.
Hybrid retrieval. Vector similarity + BM25 keyword + entity boost +
time-decay, fused into one score. Over-fetch, re-rank, return.
Async batched distillation. Writes queue behind a silence window and are
extracted in batches - ~80% fewer LLM calls than per-message extraction.
Semantic cache. Repeat adds/searches hit a similarity-gated cache and
return cached facts without re-embedding or re-querying Qdrant.
Stores preferences, not code. extraction_filter strips large fenced
code blocks before LLM extraction, so the store fills with durable facts,
not pasted implementations.
MCP server. deepmem_write / deepmem_search / deepmem_delete tools
for Claude Desktop, Cursor, and any MCP client.
GDPR. Soft-delete with retention window, hard-delete reset , SHA-256
id masking in logs, export/import for portability.
Method
Path
Description
POST
/v1/memories
Write messages; LLM-extract facts ( infer=false stores raw)
POST
/v1/memories/search
Semantic search (vector + BM25 + entity + time-decay)
GET
/v1/memories
List all for a user_id (paginated)
GET
/v1/memories/{id}
Get one by ID
PUT
/v1/memories/{id}
Update one memory's text
DELETE
/v1/memories/{id}
Soft-delete one
DELETE
/v1/memories
Soft-delete all for a user_id (GDPR)
GET
/v1/memories/{id}/history
ADD/UPDATE/DELETE audit log
POST
/v1/reset
Hard-delete all + history (needs confirm_user_id )
GET
/v1/export · POST /v1/import
Portable JSON export / import
GET
/health · /ready
Liveness / readiness probes
agent_id / run_id optionally scope writes/reads (mirrors Mem0's three-level
isolation: user -> agent -> run). Interactive docs at /docs .
{
"mcpServers" : {
"deepmem" : {
"command" : " python " ,
"args" : [ " server/mcp_server.py " ],
"env" : { "DEEPMEMORY_BASE_URL" : " http://localhost:8000 " }
}
}
}
Open mode needs no API key - DEEPMEMORY_API_KEY stays empty.
client (HTTP / MCP)
-> FastAPI (:8000)
-> rate-limit middleware (per-IP token bucket on add/search)
-> SemanticCache.check (return cached facts on similarity hit)
-> AsyncBatchDistiller.enqueue (POST /v1/memories, infer=true)
↳ silence-window or max_batch triggers on_batch_ready
↳ VectorStore.process_batch (LLM extraction -> Qdrant upsert)
-> VectorStore.search (vector + BM25 + entity + time-decay)

[truncated]
