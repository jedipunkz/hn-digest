---
source: "https://github.com/itsmeduncan/commonplace"
hn_url: "https://news.ycombinator.com/item?id=48740235"
title: "Commonplace: Self-hosted, privacy-tiered memory for your AI agents"
article_title: "GitHub - itsmeduncan/commonplace: Self-hosted, privacy-tiered memory for your AI agents. · GitHub"
author: "itsmeduncan"
captured_at: "2026-06-30T23:26:57Z"
capture_tool: "hn-digest"
hn_id: 48740235
score: 1
comments: 0
posted_at: "2026-06-30T22:46:21Z"
tags:
  - hacker-news
  - translated
---

# Commonplace: Self-hosted, privacy-tiered memory for your AI agents

- HN: [48740235](https://news.ycombinator.com/item?id=48740235)
- Source: [github.com](https://github.com/itsmeduncan/commonplace)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T22:46:21Z

## Translation

タイトル: Commonplace: AI エージェント用の自己ホスト型のプライバシー階層型メモリ
記事のタイトル: GitHub - itsmeduncan/commonplace: AI エージェント用の自己ホスト型のプライバシー階層型メモリ。 · GitHub
説明: AI エージェント用の自己ホスト型のプライバシー階層型メモリ。 - そのメダンカン/ありふれたもの

記事本文:
GitHub - itsmeduncan/commonplace: AI エージェント用の自己ホスト型のプライバシー階層型メモリ。 · GitHub
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
そのメドゥンカン
/
ありふれた
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインB

牧場 タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット .github .github config config docs docs eval eval ゲートウェイ ゲートウェイ スクリプト スクリプト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml patch_transport_security.py patch_transport_security.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
MCP が提供する自己ホスト型の 2 層 Graphiti ナレッジ グラフ
クライアント (Claude Code や Pi など) はプライベート経由で読み取りおよび書き込みを行います。
テールスケールネットワーク。オフラインファーストです。デフォルトでは、すべての部分が含まれます。
グラフを抽出する LLM — 独自のハードウェアで実行されるため、箱から出てくるものは何もありません。
Docker とコンシューマ NVIDIA GPU を備えた単一の常時稼働 Linux ホスト上で実行されます。あなたのラップトップと
他のデバイスは純粋なクライアントであり、何もホストしません。
ナレッジグラフの取り込みでは、LLM を使用してテキストからエンティティと関係を抽出します。それ
抽出はデータがモデルに公開される場所であるため、デフォルトで一般的に行われています。
ローカルで、GPU 上で、両方の層に対して。 2 つの層は、機密性とセキュリティによってメモリを分割します。
品質のために地域性を犠牲にすることが許されるかどうか:
個人層はデフォルトではローカルですが、ホストされたモデル (Claude Haiku など) を指す場合もあります。
非機密データの高品質なグラフについては、.env 経由でオプトインします (「ホスト型アップグレード?」を参照)。
[セットアップ] の下)。クライアント層は常にローカルです。それが要点です。
どちらの層でも取得は安価でプライベートです。検索はエンベディング + BM25 + グラフ走査です
クエリ パスに LLM がありません。 GPU は低速の非同期バックグラウンドのみを実行します
抽出

— クエリのレイテンシは影響を受けません。したがって、ローカル抽出は遅くても大丈夫です。
両方の層は 1 つのエンベッダー (Ollama nomic-embed-text 、 768-dim) と 1 つの FalkorDB を共有します。
2 つの別々のグラフを保持しているため、2 つのメモリは分離されたままですが、インフラストラクチャはシンプルなままです。
フローチャートTB
CC["クロード・コード<br/>(クライアント)"]
PI["ピ<br/>(クライアント)"]
TS{{"テールスケール<br/>MagicDNS · テールネットのみ"}}
ANT["Anthropic API<br/>Claude Haiku 4.5 · ホスト"]
CC --> TS
PI --> TS
サブグラフ HOST["あなたのサーバー — Docker"]
TB方向
GW["<b>ゲートウェイ</b> :8000 / :8001<br/>階層ごとの認証 · ロギング · メトリクス"]
MP["<b>mcp-personal</b><br/>個人層 · 内部"]
MC["<b>mcp-client</b><br/>クライアント機密 · 内部"]
GW --> MP
GW → MC
OL["<b>オラマ</b> :11434<br/>nomic-embed-text · mistral:7b<br/>ローカル GPU"]
サブグラフ FALKOR["FalkorDB :6379 · ブラウザ UI :3000"]
LR方向
GP[("commonplace_personal")]
GC[("commonplace_client")]
終わり
MP -->|ストア| GP
MC -->|ストア| GC
議員 - 。埋め込む .-> OL
MC-。埋め込む .-> OL
MC -->|抽出・ローカル| OL
終わり
TS -->|ベアラートークン| GW
MP -->|抽出 · ホスト|アリ
classDef ext fill:#fff3e0、ストローク:#e67e22、color:#111;
classDef 層の塗りつぶし:#e8f0fe、ストローク:#4285f4、色:#111;
クラス ANT 内線
クラスMP、MC層
読み込み中
1 つの FalkorDB 、FALKORDB_DATABASE によってインスタンスごとに選択された 2 つのグラフ
( commonplace_personal と commonplace_client )。
2 つの Graphiti MCP インスタンス ( commonplace-mcp:local 、から構築)
zepai/knowledge-graph-mcp:standalone — Dockerfile を参照)、HTTP トランスポート、パスで提供
/mcp/ (末尾のスラッシュ)。
両方のインスタンスで使用される 1 つの共有 Ollama エンベッダー ( nomic-embed-text 、 768-dim)。しないでください
エンベッダーを混合する — 異なるエンベッダーからのベクトルは比較できません。
ゲートウェイ (キャディ) は両方の層の前面にあります。ゲートウェイ (キャディ) はホスト ポートを所有し、層ごとのベアラーを必要とします。
トークン (つまり、クライアントのみを持つクライアント)

トークンは個人層に到達できません)、アクセスを発行します
ログ (監査) + Prometheus メトリクス。 MCP コンテナ自体は内部専用です。
your-server.your-tailnet.ts.net をホストの Tailscale MagicDNS 名に全体的に置き換えます
(ホスト上で tailscale status を実行して見つけます)。
パーソナル/クライアント エンドポイントには承認が必要です: Bearer <tier-token> (set PERSONAL_TOKEN /
.env の CLIENT_TOKEN)。正しいトークンのないリクエストは 401 を受け取ります。
Ollama はホスト上で実行され、共有エンベッダーとローカル サーバーにサービスを提供します。
抽出モデル。 MCP コンテナは HTTP 経由でアクセスします。GPU は、GPU ではなく Ollama によって使用されます。
コンテナなので、Docker への GPU パススルーは必要ありません。約 8 GB の VRAM を備えたコンシューマ NVIDIA GPU
misstral:7b-instruct-q4_0 を快適に実行します。 CPU のみで動作しますが、ローカル抽出は遅いです。
テールスケール — MCP エンドポイントは、テールネット経由ではなく、テールネット経由で提供されます。
公共のインターネット。
API キーは必要ありません。両方の層はデフォルトでローカルに抽出します。 Anthropic API キーが必要です
個人層をホスト型モデルにオプトインした場合のみ (下記のホスト型アップグレードを参照)。
各クライアント (ラップトップなど): Tailscale、および MCP 対応クライアント (Claude Code、Pi など)。
このリポジトリのクローンからホスト上で実行します (例: ~/commonplace ):
#1. オラマが仕えるモデルを引っ張る
オラマプル nomic-embed-text
オラマ プル ミストラル:7b-instruct-q4_0
# 2. シークレットを設定する
cp .env.example .env
# .env を編集して設定します:
# FALKORDB_PASSWORD (openssl rand -hex 24)
# PERSONAL_TOKEN / CLIENT_TOKEN ゲートウェイベアラートークン (openssl rand -hex それぞれ 32)
# (ANTHROPIC_API_KEY は必要ありません。デフォルトでは抽出はローカルに行われます)
# 3. ローカルイメージを構築してスタックを開始する
ドッカー構成 -d
docker compose ps # すべてのサービスが正常であることを報告する必要があります
次に、クライアントを 2 つのエンドポイントに向けます。「 クライアントの構成 」を参照してください。
ホスト型アップグレード?私がすべて

■ デフォルトではローカルです。個人層をホストされたサーバーに向けるには
高品質のグラフ (非機密データのみ) のモデル、 .env で設定:
PERSONAL_LLM_PROVIDER=人間性 、PERSONAL_LLM_MODEL=クロード-俳句-4-5 、
PERSONAL_SEMAPHORE_LIMIT=5 、および ANTHROPIC_API_KEY=… 。それにもかかわらず、クライアント層はローカルのままです。
ゲートウェイ導入前の環境からアップグレードしますか? PERSONAL_TOKEN / CLIENT_TOKEN を .env に追加してから、
docker compose up -d --build --force-recreate (MCP 層はゲートウェイの背後に移動し、
オントロジーの変更には再作成が必要です)。各クライアントを Authorization: Bearer ヘッダーを使用して再追加します。
既存のトークンレスクライアントは 401 を取得し始めます。
注意事項 (苦労して学んだので、これをコピーする前に読んでください)
これらは、現在 (2026 年) Graphiti MCP サーバーに特有の地雷です。いくつかの矛盾がある
古いドキュメント。
openai_generic プロバイダー文字列はありません。 Ollama を使用するには、プロバイダーを「openai」に設定します。
api_url を非 OpenAI URL に指定します。その後、サーバーは OpenAIGenericClient を自動選択します
内部的には。この汎用クライアントは、OpenAI のベータ版のresponse.parse() (Ollama が実行する) を回避するものです。
実装しません）。プロバイダーの設定:「openai_generic」が無効です。
small_model 設定はありません。 MCP サーバーには 1 つの llm.model があります。オープンナイで
パスでは、「小さい」スロットにも同じモデルが使用されます。悪名高い gpt-4.1-mini は単なる
モデルが None の場合に使用されるフォールバック — ヒットしないようにするには、llm.model を固定するだけで十分です。
json_schema 構造化出力はローカル パスに対して常にオンであり、無効にすることはできません。
そこではインストラクターは使用されません。再試行が組み込まれています (粘り強さ、4 回の試行)。ありません
いずれかの設定ノブを使用します。小規模なローカル モデルが無効な JSON を生成する場合、唯一の手段は、より多くの方法です。
有能なモデル。
Ollama にはコンテナ内からアクセスできる必要があります。 Ollama はホスト上で実行されるため、各 MCP
サービスには追加料金が必要です_

hosts: ["host.docker.internal:host-gateway"] および api_url
http://host.docker.internal:11434/v1 。 Ollama は 0.0.0.0:11434 でリッスンする必要があります (デフォルトでリッスンします)。
FALKORDB_DATABASE はグラフを選択します。 group_id はそうではありません。 1 つの FalkorDB に 2 つのグラフ =
同じ FALKORDB_URI と異なる FALKORDB_DATABASE を持つ 2 つのインスタンス。 group_id のみ
グラフ内の名前空間ノード。
FalkorDB ホスト/ポートは FALKORDB_URI から解析されます — FALKORDB_HOST / FALKORDB_PORT は
無視されました。読み取られる env オーバーライドは FALKORDB_URI と FALKORDB_PASSWORD のみです。
FalkorDB パスワードは、オーバーライドではなく、環境変数 REDIS_ARGS=--requirepass … を介して設定されます
コンテナー コマンド (FalkorDB モジュールのロードを停止します)。
:latest ではなく、 :standalone イメージを使用してください。 zepai/knowledge-graph-mcp:latest バンドル
独自の FalkorDB。 :standalone は外部のものを想定しています — 単一の FalkorDB を共有するために必要です
2つのインスタンス。
MCP パスの末尾にはスラッシュが付いています: /mcp/ (FastMCP のデフォルト。構成不可)。
人間モデル ID: claude-haiku-4-5-latest ではなく、裸のエイリアス claude-haiku-4-5 を使用します。
-latest サフィックスは OpenAI 主義です。 Anthropic API 404 ( not_found_error: model )。
裸のエイリアスは、現在の日付のスナップショット ( claude-haiku-4-5-20251001 ) に解決されます。
Anthropic プロバイダーには、明示的な数値 llm.temperature が必要です。グラフィティパス
温度=config.温度 ;何も設定されていない場合は、null と API 400 を送信します。
(温度: 入力は有効な数値である必要があります)、個人レベルのエピソードはすべてキューに入れられますが、
決して処理しません。 OpenAI/Ollama 汎用クライアントは null を許容するため、これは
人間層。例を設定します。温度: 0.0 。
:standalone イメージには、anthropic SDK は含まれていません。プロバイダー: anthropic は失敗します
起動時 — 「Anthropic クライアントは現在のグラフィティコア バージョンでは利用できません」 (f

俳優の
インポート anthropic がレイズするため、HAS_ANTHROPIC は False です)。バンドルされている Dockerfile により追加されます
( uv pip install anthropic )。
graphiti-core は、初期化時に OPENAI_API_KEY を要求するデフォルトの OpenAI リランカーを構築します
ただし、検索パスは NODE_HYBRID_SEARCH_RRF (クロスエンコーダーなし) を使用します。各層にダミーを与える
OPENAI_API_KEY を構築できるようにします。 OPENAI_BASE_URL が Ollama を指すようにすることで、偶発的であっても
通話はオンボックスのままになります。実際には呼び出されることはありません。
FastMCP は、ローカルホスト以外のホスト ヘッダーを HTTP 421「無効なホスト ヘッダー」で拒否します。それ
構築時にローカルホストのみの許可リストを使用して DNS 再バインド保護を自動的に有効にし、パスします
そのオブジェクトはそのpydantic設定に明示的に組み込まれるため、FASTMCP_…環境変数はそれをオーバーライドできません
(init kwargs ビート env)。バンドルされている patch_transport_security.py (Dockerfile で実行) は、
保護 - ネットワークが信頼境界でクライアントがエージェントであるテールネット上では安全です。
ブラウザではありません。厳格にするには、代わりに明示的に allowed_hosts を設定します。
OpenAI 互換のベース URL のコンテナ環境変数は OPENAI_API_URL (graphiti の
config 拡張)、 OPENAI_BASE_URL ではありません。リランカー (#13) はその逆であることに注意してください。
OpenAI SDK の OPENAI_BASE_URL 。 2 つの異なるクライアントに対する 2 つの異なる名前。
ホスト上で実行します。

[切り捨てられた]

## Original Extract

Self-hosted, privacy-tiered memory for your AI agents. - itsmeduncan/commonplace

GitHub - itsmeduncan/commonplace: Self-hosted, privacy-tiered memory for your AI agents. · GitHub
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
itsmeduncan
/
commonplace
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits .github .github config config docs docs eval eval gateway gateway scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml patch_transport_security.py patch_transport_security.py View all files Repository files navigation
A self-hosted, two-tier Graphiti knowledge graph that MCP
clients (for example Claude Code and Pi ) read from and write to over a private
Tailscale network. It's offline-first: by default every part — including
the LLM that extracts your graph — runs on your own hardware, so nothing leaves the box.
It runs on a single always-on Linux host with Docker and a consumer NVIDIA GPU. Your laptops and
other devices are pure clients — they host nothing.
Knowledge-graph ingestion uses an LLM to extract entities and relationships from text. That
extraction is where your data would be exposed to a model — so by default commonplace does it
locally , on your GPU, for both tiers. The two tiers split memory by confidentiality and by
whether you're allowed to trade locality for quality:
The personal tier is local by default but may be pointed at a hosted model (e.g. Claude Haiku)
for higher-quality graphs on non-confidential data — opt in via .env (see Hosted upgrade?
under Setup ). The client tier is always local; that's the whole point of it.
Retrieval is cheap and private on both tiers. Search is embeddings + BM25 + graph traversal
with no LLM in the query path . The GPU only ever does slow, asynchronous background
extraction — query latency is never affected. Slow local extraction is therefore fine.
Both tiers share one embedder (Ollama nomic-embed-text , 768-dim) and one FalkorDB
holding two separate graphs, so the two memories stay isolated but the infrastructure stays simple.
flowchart TB
CC["Claude Code<br/>(client)"]
PI["Pi<br/>(client)"]
TS{{"Tailscale<br/>MagicDNS · tailnet-only"}}
ANT["Anthropic API<br/>Claude Haiku 4.5 · hosted"]
CC --> TS
PI --> TS
subgraph HOST["your server — Docker"]
direction TB
GW["<b>gateway</b> :8000 / :8001<br/>per-tier auth · logging · metrics"]
MP["<b>mcp-personal</b><br/>personal tier · internal"]
MC["<b>mcp-client</b><br/>client-confidential · internal"]
GW --> MP
GW --> MC
OL["<b>Ollama</b> :11434<br/>nomic-embed-text · mistral:7b<br/>local GPU"]
subgraph FALKOR["FalkorDB :6379 · browser UI :3000"]
direction LR
GP[("commonplace_personal")]
GC[("commonplace_client")]
end
MP -->|store| GP
MC -->|store| GC
MP -. embed .-> OL
MC -. embed .-> OL
MC -->|extract · local| OL
end
TS -->|Bearer token| GW
MP -->|extract · hosted| ANT
classDef ext fill:#fff3e0,stroke:#e67e22,color:#111;
classDef tier fill:#e8f0fe,stroke:#4285f4,color:#111;
class ANT ext
class MP,MC tier
Loading
One FalkorDB , two graphs selected per-instance by FALKORDB_DATABASE
( commonplace_personal vs commonplace_client ).
Two Graphiti MCP instances ( commonplace-mcp:local , built from
zepai/knowledge-graph-mcp:standalone — see Dockerfile ), HTTP transport, served at path
/mcp/ (trailing slash).
One shared Ollama embedder ( nomic-embed-text , 768-dim) used by both instances. Do not
mix embedders — vectors from different embedders are not comparable.
A gateway (Caddy) fronts both tiers: it owns the host ports, requires a per-tier bearer
token (so a client with only the client token can't reach the personal tier), and emits access
logs (audit) + Prometheus metrics. The MCP containers themselves are internal-only.
Replace your-server.your-tailnet.ts.net with your host's Tailscale MagicDNS name throughout
(run tailscale status on the host to find it).
The personal/client endpoints require Authorization: Bearer <tier-token> (set PERSONAL_TOKEN /
CLIENT_TOKEN in .env ). A request without the right token gets 401 .
Ollama running on the host, serving the shared embedder and the local
extraction model. The MCP containers reach it over HTTP — the GPU is used by Ollama, not by the
containers, so no GPU passthrough into Docker is required. A consumer NVIDIA GPU with ~8 GB VRAM
runs mistral:7b-instruct-q4_0 comfortably; CPU-only works but local extraction is slow.
Tailscale — the MCP endpoints are served over the tailnet, not the
public internet.
No API keys required. Both tiers extract locally by default. An Anthropic API key is needed
only if you opt the personal tier into a hosted model (see Hosted upgrade? below).
On each client (laptop, etc.): Tailscale, plus an MCP-capable client (Claude Code, Pi, …).
Run on the host, from a clone of this repo (e.g. ~/commonplace ):
# 1. Pull the models Ollama will serve
ollama pull nomic-embed-text
ollama pull mistral:7b-instruct-q4_0
# 2. Configure secrets
cp .env.example .env
# edit .env and set:
# FALKORDB_PASSWORD (openssl rand -hex 24)
# PERSONAL_TOKEN / CLIENT_TOKEN gateway bearer tokens (openssl rand -hex 32 each)
# (no ANTHROPIC_API_KEY needed — extraction is local by default)
# 3. Build the local image and start the stack
docker compose up -d
docker compose ps # all services should report healthy
Then point a client at the two endpoints — see Client configuration .
Hosted upgrade? Everything is local by default. To point the personal tier at a hosted
model for higher-quality graphs (non-confidential data only), set in .env :
PERSONAL_LLM_PROVIDER=anthropic , PERSONAL_LLM_MODEL=claude-haiku-4-5 ,
PERSONAL_SEMAPHORE_LIMIT=5 , and ANTHROPIC_API_KEY=… . The client tier stays local regardless.
Upgrading from a pre-gateway deploy? Add PERSONAL_TOKEN / CLIENT_TOKEN to .env , then
docker compose up -d --build --force-recreate (the MCP tiers move behind the gateway and the
ontology change needs a recreate). Re-add each client with its Authorization: Bearer header —
existing token-less clients will start getting 401 .
Gotchas (learned the hard way — read before you copy this)
These are the landmines specific to the current (2026) Graphiti MCP server. Several contradict
older docs.
There is no openai_generic provider string. To use Ollama you set provider: "openai"
and point api_url at a non-OpenAI URL; the server then auto-selects its OpenAIGenericClient
internally. That generic client is what avoids OpenAI's beta responses.parse() (which Ollama
does not implement). Setting provider: "openai_generic" is invalid.
There is no small_model setting. The MCP server has a single llm.model . On the openai
path it uses that same model for the "small" slot too. The infamous gpt-4.1-mini is only a
fallback used when model is None — pinning llm.model is enough to never hit it.
json_schema structured output is always on for the local path and cannot be disabled , and
instructor is not used there — retries are built-in (tenacity, 4 attempts). There is no
config knob for either. If a small local model produces invalid JSON, the only lever is a more
capable model.
Ollama must be reachable from inside the containers. Ollama runs on the host , so each MCP
service needs extra_hosts: ["host.docker.internal:host-gateway"] and an api_url of
http://host.docker.internal:11434/v1 . Ollama must listen on 0.0.0.0:11434 (it does by default).
FALKORDB_DATABASE selects the graph; group_id does not. Two graphs in one FalkorDB =
two instances with the same FALKORDB_URI and different FALKORDB_DATABASE . group_id only
namespaces nodes within a graph.
FalkorDB host/port are parsed from FALKORDB_URI — FALKORDB_HOST / FALKORDB_PORT are
ignored. The only env overrides read are FALKORDB_URI and FALKORDB_PASSWORD .
FalkorDB password is set via REDIS_ARGS=--requirepass … , an env var — not by overriding
the container command (that would stop the FalkorDB module from loading).
Use the :standalone image, not :latest . zepai/knowledge-graph-mcp:latest bundles its
own FalkorDB; :standalone expects an external one — required to share a single FalkorDB across
two instances.
The MCP path has a trailing slash: /mcp/ (FastMCP default; not configurable).
Anthropic model id: use the bare alias claude-haiku-4-5 , not claude-haiku-4-5-latest .
The -latest suffix is an OpenAI-ism; the Anthropic API 404s on it ( not_found_error: model ).
The bare alias resolves to the current dated snapshot ( claude-haiku-4-5-20251001 ).
The Anthropic provider needs an explicit numeric llm.temperature . graphiti passes
temperature=config.temperature ; with none set it sends null and the API 400s
( temperature: Input should be a valid number ), so every personal-tier episode queues but
never processes. The OpenAI/Ollama generic client tolerates null , so this bites only the
Anthropic tier. Set e.g. temperature: 0.0 .
The :standalone image ships WITHOUT the anthropic SDK. provider: anthropic then fails
at startup — "Anthropic client not available in current graphiti-core version" (the factory's
HAS_ANTHROPIC is False because import anthropic raises). The bundled Dockerfile adds it
( uv pip install anthropic ).
graphiti-core builds a default OpenAI reranker at init that demands OPENAI_API_KEY even
though the search path uses NODE_HYBRID_SEARCH_RRF (no cross-encoder). Give each tier a dummy
OPENAI_API_KEY so it can construct; point OPENAI_BASE_URL at Ollama so even an accidental
call stays on-box. In practice it is never called.
FastMCP rejects non-localhost Host headers with HTTP 421 "Invalid Host header". It
auto-enables DNS-rebinding protection with a localhost-only allow-list at construction and passes
that object explicitly into its pydantic Settings, so the FASTMCP_… env vars cannot override it
(init kwargs beat env). The bundled patch_transport_security.py (run in the Dockerfile) disables
the protection — safe on a tailnet, where the network is the trust boundary and clients are agents,
not browsers. To tighten, set explicit allowed_hosts instead.
The container env var for the OpenAI-compatible base URL is OPENAI_API_URL (graphiti's
config expansion), not OPENAI_BASE_URL . Note the reranker (#13) is the opposite — it reads the
OpenAI SDK's OPENAI_BASE_URL . Two different names for two different clients.
Run on the host, from the

[truncated]
