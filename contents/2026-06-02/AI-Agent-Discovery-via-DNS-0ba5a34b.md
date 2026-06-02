---
source: "https://dns-aid.org/"
hn_url: "https://news.ycombinator.com/item?id=48355160"
title: "AI Agent Discovery via DNS"
article_title: "DNS-AID — AI Agent Discovery via DNS"
author: "oogali"
captured_at: "2026-06-02T00:42:21Z"
capture_tool: "hn-digest"
hn_id: 48355160
score: 1
comments: 0
posted_at: "2026-06-01T10:56:34Z"
tags:
  - hacker-news
  - translated
---

# AI Agent Discovery via DNS

- HN: [48355160](https://news.ycombinator.com/item?id=48355160)
- Source: [dns-aid.org](https://dns-aid.org/)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T10:56:34Z

## Translation

タイトル: DNS 経由の AI エージェント検出
記事のタイトル: DNS-AID — DNS を介した AI エージェントの検出

記事本文:
D
DNS-AID
概要
DNS-AID 概要資料
ポリシー
さらに詳しく
ルッキンググラス
☀️
GitHub
IETF ドラフト · オープンソース
普遍的なもの。 AI エージェントの検出レイヤー。
エージェントを DNS に公開し、Web サイトと同様に検出し、DNSSEC との信頼性を検証します。一元化されたレジストリはなく、シグナルのみです。
完全な SDK をワンショットでインストールします。
DNS-AID が提供するものは、DNS-AID プロトコルに基づいて構築されています。
新しいインフラストラクチャはゼロです。
すでに実行している DNS 上に構築されます。
DNS-AID は、既存の SVCB、TXT、および TLSA レコードに基づく命名規則です。新しいレコード タイプ、新しいサーバー、新しいプロトコルはなく、RFC 9460 と RFC 4033 の標準だけが使用されます。
エージェントの記録が本物であり、改ざんされていないことを示す暗号による証明。
MCP、A2A、HTTPS、および alpn を介した将来のプロトコル。
名前による検索、機能による検索、またはドメイン インデックスのクロールを行います。
内部エージェントと外部エージェントではエージェントが異なります。組み込みのテナント分離。
CLI、Python SDK、MCP サーバー。 8 つのバックエンドが同梱されています。
DNS は自動的にキャッシュされます。中央 API はありません。分散ルックアップ。
エージェント レコードの決定的で人間が判読できる名前付けパターン。
_<エージェント名> 。 _<プロトコル> ._agents.<ドメイン>
例:
_チャットボット 。 _mcp ._agents.example.com MCP チャットボット
_検索 。 _a2a ._agents.example.com A2A 検索エージェント
_data-cleaner 。 _a2a ._agents.acme.com 機能ベース
_index._agents.example.com 完全なエージェント インデックス
マルチテナント:
_analytics 。 _mcp ._agents.customer1.saas.com
エージェントレコードの構造
各エージェントは、機械可読メタデータが詰め込まれた SVCB レコードです。
_my-agent._mcp._agents.example.com。 3600 IN SVCB 1 エージェント.example.com。 (
alpn = "mcp" ;プロトコル
ポート =443 ;サービスポート
キャップ = "https://example.com/cap.json" ;機能ドキュメント
cap-sha256 = "abc123..." ;整合性ハッシュ
bap = "mcp=1.0,a2a=0.2" ;プロトコルのバージョン
ポリシー = "https://example.com/policy" ;ガバナンス URL
レルム =

"生産" ;テナントの範囲
ipv4hint =192.0.2.1 ;アドレスのヒント
)
alpn 通信プロトコル（mcp、a2a、h2）
port サービスポート番号
cap 機能ドキュメント URI
cap-sha256 改ざん検出のための整合性ハッシュ
bap バルクプロトコルバージョン宣言
ポリシー ガバナンスと使用ポリシーの URL
レルム テナントまたは環境スコープ
ipv4hint 余分な検索を減らすためのアドレス ヒント
仕組み
公開から接続までの 4 つのステップ。
CLI または SDK を使用して、エンドポイント、プロトコル、機能を含む SVCB レコードをドメインの _agents ゾーンの下に作成します。
権威 DNS がレコードに署名し、ルートからエージェントまでの暗号化された信頼チェーンを作成します。
リモート エージェントは、名前、機能タイプ、または完全なドメイン インデックスを使用して、SVCB レコードの DNS をクエリします。
Discoverer は DNSSEC + DANE を検証し、SVCB レコード内のプロトコルを介して直接接続します。
dns-aid-core Python パッケージを起動して実行します。
pip install "dns-aid[all]" # すべて
pip install "dns-aid[cli]" # CLI のみ
pip install "dns-aid[route53]" # AWS バックエンド
pip install "dns-aid[cloudflare]" # Cloudflare バックエンド
pip install "dns-aid[nios]" # Infoblox NIOS バックエンド
pip install "dns-aid[mcp]" # MCP サーバー
発行する
DNS-AID 公開 \
--name 私のチャットボット \
--ドメイン example.com \
--プロトコル mcp \
--エンドポイント エージェント.example.com \
--能力チャット
発見する
dns-aid Discover example.com
dns-aid Discover example.com --json
dns-aid Discover example.com --use-http-index
検証と診断
dns-aid verify _my-chatbot._mcp._agents.example.com
DNS-AID ドクター --ドメイン example.com
エージェントを呼び出す
# MCP エージェント上のツールをリストします。
dns-aid リスト-ツール https://mcp.example.com/mcp
# 特定のツールを呼び出す
DNS-AID 呼び出し https://mcp.example.com/mcp Analyzer_security \
--arguments '{"ドメイン":"example.com"}'
# A2A エージェントにメッセージを送信します (検出優先)
DNS-Aid メッセージ

e「DNS-AIDとは何ですか?」 \
-d ai.infoblox.com -n セキュリティアナライザー
管理する
# DNS からエージェントを削除する
dns-aid delete -n my-chatbot -d example.com -p mcp
発行する
dns_aid からインポート 公開
result = パブリッシュを待ちます(
name= "私のチャットボット" ,
ドメイン= "example.com" 、
プロトコル= "mcp" 、
エンドポイント= "agent.example.com" ,
機能=[ "チャット" , "要約" ],
description= "汎用チャットエージェント" ,
)
print(f "公開済み: {result.agent.fqdn}" )
print(f "レコード: {result.records_created}" )
発見する
非同期をインポートする
dns_aid インポートから検出、検証
非同期デフォルトメイン():
result = await Discover( "example.com" )
result.agents のエージェントの場合:
print(f " {agent.name} — {agent.protocol} @ {agent.endpoint_url}" )
check = await verify( "_my-agent._mcp._agents.example.com" )
print(f "DNSSEC が有効です: {check.dnssec_valid}" )
asyncio.run(main())
発見してから呼び出す
dns_aid インポートから検出、呼び出し
async def find_and_call ():
result = await Discover( "partner.com" 、protocol= "mcp" )
エージェント = result.agents[0]
resp = await invoke(agent, method= "tools/list" )
print(f "レイテンシー: {resp.signal.invocation_latency
[切り捨てられた]
すべて標準の DNS クエリ経由です。特別なクライアントは必要ありません。
あなたはエージェントを知っています。エンドポイントの詳細については、SVCB レコードを直接クエリします。
エージェントの活動内容からエージェントを見つけます。エージェント ゾーンの機能タイプを照会します。
既知のインデックス エントリ ポイントからドメインの完全なエージェント インベントリを取得します。
組織間のエージェント検出フロー。
ORG 1 (検出) ORG 2 (公開)
+----------------+ +-------------------+
| AI エージェント |---- 1. DNS SVCB クエリ ----------->|権威ある |
| (org1) | _search._a2a._agents.org2.com | DNSサーバー |
| |
R53 Amazon Route 53 AWS ホストゾーン
CF Cloudflare グローバル エッジ DNS
IB Infoblox NIOS エンタープライズ DDI
UDインフォブロックス UDDI ユニバーサルDDIクラウド
AZ Azure DNS マイクロソフト クラウド
GC

Google Cloud DNS GCP マネージド
NS1 NS1 マネージド DNS およびトラフィック ステアリング
RFC RFC 2136 DDNS 標準に準拠した任意の DNS
B9 BIND9 セルフホストおよびローカル開発
ポリシーの施行
Discovery はエージェントを適切なエンドポイントに誘導します。ポリシーは、ホームページをプロトコルのメモに変えることなく、展開時に誰が呼び出し可能か、どの認証が必要か、どのランタイム チェックを適用するかを表現するのに役立ちます。
DNS-AID 学習資料では、ディスカバリ メタデータ、認証要件、およびpolicy.json などの展開固有のポリシー バンドルとともに、ランタイム ポリシー評価についてすでに説明しています。
呼び出し元、ターゲット、インフラストラクチャ
チームは、発信者側とターゲット側のチェックから開始し、DNS またはトラフィック インフラストラクチャがサポートしている場合にのみ、リゾルバーまたはプロキシの適用を追加できます。
一部の展開では、中央のポリシー サービスを介してトラフィックをルーティングせずに、PII、プロンプト インジェクション、またはデータ処理チェックに対するローカルのリクエストまたはレスポンスの検査を追加する場合があります。
このサイトに文書化されたポリシーの例には、呼び出し元ドメインの制限、必要な認証タイプ、可用性ウィンドウ、レート制限、DNSSEC に依存する決定、およびより厳密な実行時チェックのための CEL 式が含まれます。
最初から強力な集中型セキュリティ アーキテクチャにコミットするのではなく、現在のランタイム ポリシー モデルから開始して、後でリゾルバーまたはプロキシの適用に拡張することができます。
現実世界のエージェント発見シナリオ。
内部エージェントは DNS にクエリを実行してパートナーの承認されたエージェントを検出し、委任を検証して、安全なセッションを自動的に開始します。
大学は独自のドメインでエージェントを公開しています。協力者は、組織の信頼境界を尊重しながら、能力に基づいてサービスを発見します。
SaaS プロバイダーは、テナント固有のゾーンでエージェントをホストします。 DNS ゾーンの委任により、顧客ごとに自然な分離と範囲を絞った検出が可能になります。
const 上の軽量エージェント

雨が降ったデバイスは、低遅延ブートストラップのための SVCB ヒントを備えた DNS の分散キャッシュ可能なアーキテクチャの恩恵を受けます。
インターネットの実証済みのセキュリティ インフラストラクチャ上に構築されています。
パブリックゾーンでは必須です。暗号化された信頼チェーンにより、なりすましや改ざんが防止されます。
TLS 証明書を DNS レコードにバインドします。認証局の信頼の問題のないエンドポイント検証。
エージェントは、DCV TXT レコードを通じて承認を証明します。範囲が限定され、検証可能で、一時的なエージェントに最適です。
cap-sha256 ハッシュにより、機能文書が改ざんされていないことが保証されます。
内部エージェントは外部からは見えません。リゾルバーコンテキストごとに異なるビュー。
TXT レコードは、特定のサービスと操作に範囲を限定したエージェントごとのロールと権限を定義します。
SDK をインストールし、最初のエージェントを公開し、AI 用のオープン ユニバーサル ディスカバリ レイヤー上に構築します。

## Original Extract

D
DNS-AID
Overview
DNS-AID One Pager
Policy
Learn More
Looking Glass
☀️
GitHub
IETF Draft · Open Source
The universal . discovery layer for AI agents.
Publish agents to DNS, discover them like websites, and verify trust with DNSSEC. No centralized registry, just signal.
Install the full SDK in one shot:
What DNS-AID gives you, built on the DNS-AID protocol.
Zero new infrastructure.
Built on DNS you already run.
DNS-AID is a naming convention on top of existing SVCB, TXT, and TLSA records. No new record types, no new servers, no new protocols — just standards from RFC 9460 and RFC 4033.
Cryptographic proof that agent records are authentic and untampered.
MCP, A2A, HTTPS, and any future protocol via alpn .
Lookup by name, search by capability, or crawl a domain index.
Different agents to internal vs. external. Built-in tenant isolation.
CLI, Python SDK, MCP server. Eight backends ship in the box.
DNS caches automatically. No central API. Distributed lookups.
A deterministic, human-readable naming pattern for agent records.
_<agent-name> . _<protocol> ._agents.<your-domain>
Examples:
_chatbot . _mcp ._agents.example.com MCP chatbot
_search . _a2a ._agents.example.com A2A search agent
_data-cleaner . _a2a ._agents.acme.com capability-based
_index._agents.example.com full agent index
Multi-tenant:
_analytics . _mcp ._agents.customer1.saas.com
Anatomy of an agent record
Each agent is an SVCB record packed with machine-readable metadata.
_my-agent._mcp._agents.example.com. 3600 IN SVCB 1 agent.example.com. (
alpn = "mcp" ; protocol
port =443 ; service port
cap = "https://example.com/cap.json" ; capability doc
cap-sha256 = "abc123..." ; integrity hash
bap = "mcp=1.0,a2a=0.2" ; protocol versions
policy = "https://example.com/policy" ; governance URL
realm = "production" ; tenant scope
ipv4hint =192.0.2.1 ; address hint
)
alpn Communication protocol (mcp, a2a, h2)
port Service port number
cap Capability document URI
cap-sha256 Integrity hash for tamper detection
bap Bulk protocol version declarations
policy Governance and usage policy URL
realm Tenant or environment scope
ipv4hint Address hint to reduce extra lookups
How it works
Four steps from publish to connect.
Use the CLI or SDK to create an SVCB record under your domain's _agents zone with endpoint, protocol, and capabilities.
Your authoritative DNS signs the records, creating a cryptographic chain of trust from root to your agent.
Remote agents query DNS for your SVCB record by name, capability type, or full domain index.
The discoverer validates DNSSEC + DANE, then connects directly via the protocol in your SVCB record.
Get up and running with the dns-aid-core Python package.
pip install "dns-aid[all]" # everything
pip install "dns-aid[cli]" # CLI only
pip install "dns-aid[route53]" # AWS backend
pip install "dns-aid[cloudflare]" # Cloudflare backend
pip install "dns-aid[nios]" # Infoblox NIOS backend
pip install "dns-aid[mcp]" # MCP server
Publish
dns-aid publish \
--name my-chatbot \
--domain example.com \
--protocol mcp \
--endpoint agent.example.com \
--capability chat
Discover
dns-aid discover example.com
dns-aid discover example.com --json
dns-aid discover example.com --use-http-index
Verify & Diagnose
dns-aid verify _my-chatbot._mcp._agents.example.com
dns-aid doctor --domain example.com
Invoke agents
# List tools on an MCP agent
dns-aid list-tools https://mcp.example.com/mcp
# Call a specific tool
dns-aid call https://mcp.example.com/mcp analyze_security \
--arguments '{"domain":"example.com"}'
# Send a message to an A2A agent (discover-first)
dns-aid message "What is DNS-AID?" \
-d ai.infoblox.com -n security-analyzer
Manage
# Delete an agent from DNS
dns-aid delete -n my-chatbot -d example.com -p mcp
Publish
from dns_aid import publish
result = await publish(
name= "my-chatbot" ,
domain= "example.com" ,
protocol= "mcp" ,
endpoint= "agent.example.com" ,
capabilities=[ "chat" , "summarize" ],
description= "General-purpose chat agent" ,
)
print(f "Published: {result.agent.fqdn}" )
print(f "Records: {result.records_created}" )
Discover
import asyncio
from dns_aid import discover, verify
async def main ():
result = await discover( "example.com" )
for agent in result.agents:
print(f " {agent.name} — {agent.protocol} @ {agent.endpoint_url}" )
check = await verify( "_my-agent._mcp._agents.example.com" )
print(f "DNSSEC valid: {check.dnssec_valid}" )
asyncio.run(main())
Discover-then-Invoke
from dns_aid import discover, invoke
async def find_and_call ():
result = await discover( "partner.com" , protocol= "mcp" )
agent = result.agents[ 0 ]
resp = await invoke(agent, method= "tools/list" )
print(f "Latency: {resp.signal.invocation_latency
[truncated]
All via standard DNS queries. No special client needed.
You know the agent. Query its SVCB record directly for endpoint details.
Find agents by what they do. Query a capability type under the agent zone.
Fetch a domain's full agent inventory from a well-known index entry point.
Cross-organization agent discovery flow.
ORG 1 (Discovering) ORG 2 (Publishing)
+----------------+ +-------------------+
| AI Agent |---- 1. DNS SVCB Query ----------->| Authoritative |
| (org1) | _search._a2a._agents.org2.com | DNS Server |
| |
R53 Amazon Route 53 AWS hosted zones
CF Cloudflare Global edge DNS
IB Infoblox NIOS Enterprise DDI
UD Infoblox UDDI Universal DDI cloud
AZ Azure DNS Microsoft cloud
GC Google Cloud DNS GCP managed
NS1 NS1 Managed DNS & traffic steering
RFC RFC 2136 DDNS Any standards-compliant DNS
B9 BIND9 Self-hosted & local dev
Policy enforcement
Discovery gets agents to the right endpoint. Policy helps deployments express who may call, what auth is required, and which runtime checks to apply without turning the homepage into a protocol memo.
The DNS-AID learning materials already describe runtime policy evaluation alongside discovery metadata, auth requirements, and deployment-specific policy bundles such as policy.json .
Caller, target, and infrastructure
Teams can start with caller-side and target-side checks, then add resolver or proxy enforcement only if their DNS or traffic infrastructure supports it.
Some deployments may add local request or response inspection for PII, prompt injection, or data handling checks without routing traffic through a central policy service.
Documented policy examples on this site include caller-domain restrictions, required auth types, availability windows, rate limits, DNSSEC-sensitive decisions, and CEL expressions for tighter runtime checks.
You can start with the current runtime policy model and extend into resolver or proxy enforcement later, instead of committing to a heavyweight centralized security architecture on day one.
Real-world agent discovery scenarios.
An internal agent queries DNS to discover a partner's authorized agents, validates delegation, and initiates a secure session automatically.
Universities publish agents under their own domains. Collaborators discover services by capability while respecting institutional trust boundaries.
SaaS providers host agents under tenant-specific zones. DNS zone delegation provides natural isolation and scoped discovery per customer.
Lightweight agents on constrained devices benefit from DNS's distributed, cacheable architecture with SVCB hints for low-latency bootstrapping.
Built on the internet's battle-tested security infrastructure.
Mandatory for public zones. Cryptographic chain of trust prevents spoofing and tampering.
Binds TLS certificates to DNS records. Endpoint verification without certificate authority trust issues.
Agents prove authorization via DCV TXT records. Scoped, verifiable, and ideal for ephemeral agents.
cap-sha256 hash ensures capability documents haven't been tampered with.
Internal agents stay invisible externally. Different views for different resolver contexts.
TXT records define per-agent roles and permissions scoped to specific services and operations.
Install the SDK, publish your first agent, and build on the open universal discovery layer for AI.
