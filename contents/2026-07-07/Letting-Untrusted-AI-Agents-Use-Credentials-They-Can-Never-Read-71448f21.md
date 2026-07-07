---
source: "https://declaw.ai/blog/credentials-agents-can-never-read"
hn_url: "https://news.ycombinator.com/item?id=48819806"
title: "Letting Untrusted AI Agents Use Credentials They Can Never Read"
article_title: "Letting Untrusted AI Agents Use Credentials They Can Never Read | Declaw"
author: "ShivamNayak11"
captured_at: "2026-07-07T17:07:58Z"
capture_tool: "hn-digest"
hn_id: 48819806
score: 1
comments: 0
posted_at: "2026-07-07T16:10:40Z"
tags:
  - hacker-news
  - translated
---

# Letting Untrusted AI Agents Use Credentials They Can Never Read

- HN: [48819806](https://news.ycombinator.com/item?id=48819806)
- Source: [declaw.ai](https://declaw.ai/blog/credentials-agents-can-never-read)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T16:10:40Z

## Translation

タイトル: 信頼できない AI エージェントに決して読み取れない認証情報を使用させる
記事のタイトル: 信頼できない AI エージェントに決して読み取れない認証情報を使用させる |デクロー
説明: Declaw の資格情報ボールトの完全な設計: シークレットがどのようにしてサンドボックスの外に留まり、途中で出口プロキシに接続されるのか、そしてなぜ完全に侵害されたプロンプト注入されたエージェントが依然としてプレースホルダーのみをリークするのか。

記事本文:
信頼できない AI エージェントに決して読み取ることができない認証情報を使用させる | Declaw declaw Arena Startups デモ ベンチマーク ブログ 価格 ドキュメントについて デモを予約 サインイン ブログに戻る 2026-07-07 セキュリティ 信頼できない AI エージェントに決して読み取ることのできない認証情報を使用させる
Declaw の資格情報ボールトの完全な設計: シークレットがどのようにしてサンドボックスの外に留まり、途中で出力プロキシに接続されるのか、そして、完全に侵害されたプロンプト注入されたエージェントが依然としてプレースホルダーのみをリークする理由。
AI エージェントは、LLM API キー、データベース パスワード、内部サービスのトークンなどのシークレットを与えて初めて役に立ちます。引き渡す通常の方法は、エージェントが実行されるサンドボックス内の環境変数です。それが問題でもあります。サンドボックスはまさに信頼できないモデル駆動型コードが実行される場所です。プロンプトによって挿入されたエージェント (または侵害された依存関係) は、独自の環境を読み取り、 /proc をダンプし、プロセス テーブルを移動し、キーも一緒に移動する可能性があります。すでに信頼しないと決めた 1 つのコンポーネントにキーを渡しましたが、サンドボックス化してもそれは変わりません。優れたサンドボックスはコードの爆発範囲を制限しますが、秘密はその範囲内にあります。
「資格情報を渡さずに AI エージェントにコードを実行させる方法」では、シークレットの保存、スコープ設定、キーを決して保持しないエージェントからの API の呼び出しなど、実践的な側面について説明しました。この投稿は残りの半分です。エンドツーエンドの設計全体 (ストア、インジェクション、分離) と、完全に侵害されたエージェントがプレースホルダーのみをリークする理由について説明します。
そこで私たちは、箱の中にもっと良い隠れ場所を見つけるのをやめ、秘密が箱の中にまったく入らないようにしました。
シークレット値はサーバー側のボールトに保存されます (OpenBao を使用します)。これは書き込み専用です。一度保存すると、API はそれを再度返すことはありません。

エージェントではなくあなたに。
サンドボックスは参照を使用して作成されます (シークレットの名前のみ)。 VM 内では、環境変数にプレースホルダー文字列 declaw:vault-managed が保持されます。
すべてのサンドボックスの送信トラフィックは、すでに出力プロキシを経由してルーティングされています。エージェントがシークレットのスコープを設定したホストにリクエストを行うと、リクエストが VM を離れた後、プロキシは実際の資格情報を添付します。
エージェントは、意図した宛先に対して資格情報を使用し、資格情報を保持することはありません。環境全体をダンプするプロンプト挿入エージェントは、プレースホルダーをリークします。
これがデモではなく実際の境界となるのは、プロキシが実行される場所です。エージェントは Firecracker microVM で実行されます。プロキシは、サンドボックスのホスト側ネットワーク名前空間 (ゲストの外側、仮想化境界を越えた場所) で実行され、VM のネットワークへの唯一のルートです。シークレットは 1 回のリクエストに対してのみプロキシのメモリ内に存在し、ファイル、環境変数、ログ行などには決して書き込まれません。
VM がプロキシの証明書を信頼する方法
エンジニアが最初に抱く疑問は、プロキシが TLS を終了して HTTPS リクエストに挿入する場合、エージェントのクライアントはなぜ置換された証明書を拒否しないのかということです。
各サンドボックスは独自の一時 CA を取得します。これはサンドボックスごとの ECDSA P-256 キーであり、その秘密キーはプロキシのメモリから離れることはなく、ディスクに書き込まれることもありません。エージェントが接続する上流ホストごとに、プロキシはその CA によって署名された短期間のリーフを作成します。
ゲストがこれらのリーフを受け入れるために、CA の公開証明書が起動時にゲストの信頼ストアにインストールされ、システム バンドルに追加され、すべての主要なランタイムが受け入れる変数を通じてエクスポートされます。
SSL_CERT_FILE → /etc/ssl/certs/ca-certificates.crt
REQUESTS_CA_BUNDLE → /etc/ssl/certs/ca-certificates.crt
CURL_CA_BUNDLE → /etc/ssl/certs

/ca-certificates.crt
NODE_EXTRA_CA_CERTS → /usr/local/share/ca-certificates/declaw-sandbox.crt
したがって、OpenSSL、Python リクエスト、curl、および Node はすべて、アプリごとの構成を行わずにプロキシのリーフをチェーン検証します。 CA はサンドボックスごとに存在するため、あるサンドボックスの CA は別のサンドボックスのトラフィックに対しては無価値です。
同じメカニズムから 1 つの制限が生じます。システム ストアを信頼するのではなく、特定のアップストリーム証明書を固定するクライアントは、プロキシのミントされたリーフを拒否するため、これらのエンドポイントを仲介できません。サーバー間の API 呼び出しではまれですが、これは実際に起こります。
インジェクションは、環境変数ではなく、ドメインの一致で発生します。プレースホルダーは、エージェントの SDK に読み取る値があるための礼儀です。プロキシは、エージェントが変数に触れたかどうかに関係なく、スコープ指定されたホストへのリクエストを挿入します。また、プロキシは Set を使用するため、エージェントが自身に提供しようとしたヘッダーもオーバーライドします。
スコープとそれが固定されている理由
シークレットは、そのスコープのいずれかに一致するホストにのみアタッチされます。スコープのドメイン パターンは、正確なホスト、* です。 -wildcard、または ~ -接頭辞が付いた正規表現。照合は RE2 上で実行され、致命的なバックトラッキングはなく、線形時間で実行され、すべてのスコープのパターンは ^(?:…)$ に固定されているため、 api.example.com をスコープとする資格情報を api.example.com.evil.com に誘導することはできません。アンカーリングは勧告ではなく強制されます。パターンはラップされており (そのため、トップレベルの代替 a|b でも ^(?:a|b)$ として正しくアンカーされます)、組み込みのプリセットだけでなく、手書きのカスタム スコープにも適用されます。
値は KV-v2 マウントの下で OpenBao に入力され、チームごとに 1 つの名前空間になります。つまり、クエリ フィルターだけでなくストアでのテナントの分離です。私たち自身の Postgres は、メタデータとドメインごとのインジェクション ルールのみを保持します。値の列はどこにもありません。ボールトへの書き込み後にメタデータの書き込みが失敗した場合、ボールトはロールされます。

書き戻すことはできないので、中途半端に作成された秘密が孤児を残すことはありません。
秘密を私たちに保管する必要さえありません。シークレットの代わりに、コネクタ記述子、つまりすでに実行しているストアへの小さな JSON ポインタを使用することもできます。
{"プロバイダー": "aws_sm", "secret_id": "prod/openai", "リージョン": "us-east-1"}
キャッシュ ミスが発生した場合、ワーカーはストアからライブ値を読み取り、その 1 つのリクエストに使用し、永続化することはありません。 AWS Secrets Manager と SSM Parameter Store、GCP Secret Manager、Azure Key Vault、HashiCorp Vault / OpenBao、Infisical、Doppler、Kubernetes Secrets、1Password Connect、Cyber​​Ark Conjur、Akeyless の 11 のストアが接続されています。コネクタは、以下のブローカー インジェクション タイプと同じ資格情報生成機構を再利用するため、コネクタは、保存されたキーの代わりにワークロード ID を使用してストアに対して認証できます。
一般的なケースでは、プリセット カタログ (最大 39 プロバイダー) があります。プロバイダーを選択してキーを貼り付けると、宛先ホスト、インジェクション形式、およびプロバイダーの癖が自動的に入力されます。ある API が必要とするバージョン ヘッダー、別の API が必要とするトークンとキーのスキーム ワード、3 番目の固定の基本認証ユーザー名です。プリセットはサーバー側を通常のスコープに拡張するため、別個のコード パスではなく、同じ機械上の砂糖となります。
静的タイプは退屈だが重要なものです: 認可: Bearer <v> 、任意のヘッダー、HTTP 基本、またはクエリ パラメーター — 加えて 4 つのアフォーダンス (値プレフィックス、固定の基本認証ユーザー名、追加の非シークレット ヘッダー、静的クエリ パラメーター) により、新しいメカニズムを使用せずに単一のスコープでプロバイダーの実際のコントラクトを表現できます。これらはすべて、セキュリティ スキャンと PII 編集の実行後にプロキシのリクエスト ディレクタに適用されるため、値がスキャナや監査証跡に到達することはありません。
興味深いタイプは短命な関係を仲介します

保存されている資格情報を転送する代わりに、出力時に資格情報を使用します。
AWS SigV4。エージェントの SDK は使い捨て資格情報を使用して署名します。プロキシはその署名を削除し、リクエストに再署名します。保存されるマテリアルは、静的キー、AssumeRole (有効期限が切れる前に更新される短期間の認証情報)、または — キーレス パス — AssumeRoleWithWebIdentity のいずれかです。OIDC Web ID トークンが認証であるため、STS への呼び出しは署名なしで送信されます。有効期間の長い AWS キーはどこにも保存されません。
OIDC。プロキシは、ワーカー側で有効期間の短いベアラーを作成し、それを (チーム、シークレット) ごとにキャッシュし、有効期限が切れる前に再作成します。クライアント認証情報、RFC 8693 トークン交換 (サブジェクトとして投影されたサービス アカウントまたは CI OIDC トークン - GCP Workload Identity へのフェデレーションなど)、private_key_jwt、および jwt-bearer をサポートし、発行者の .well-known/openid-configuration からトークン エンドポイントを検出します。クライアント シークレットと作成されたトークンは両方ともプロキシ側に残ります。結果として得られる Authorization: Bearer のみが上流に到達します。
HMAC。構成可能なダイジェスト (sha1/256/512)、エンコーディング (hex/base64)、 {body}{method}{path}{timestamp} の署名文字列テンプレート、プレフィックス、およびヘッダー レイアウトを使用した署名をリクエストします。GitHub、Stripe、Slack の署名形状はすべて同じ構成に含まれます。
private_key_jwt (RS256/PS256/ES256) の背後にある JOSE 署名は、Go 標準ライブラリに実装されています。それで、私が最も好きな部分に行き着きます。
エージェントは、パスワードをまったく使用せずに Postgres、MySQL、MongoDB、Redis、または SMTP リレーと通信できます。通常の接続を開きます。プロキシは、プロキシに代わってアップストリームで認証ハンドシェイクを実行し、その後脇に置いてバイトをパイプします。エージェントは動作する認証されたセッションを取得し、パスワードを保持することはありません。
これは、各ワイヤ プロトコルの認証側を標準ライブラリに実装することを意味します。
ポストグル —

クライアントの起動メッセージを読み取り、SSLRequest を N で拒否してニア レッグをプレーンに保ち、クリアテキスト / MD5 / SCRAM-SHA-256 (完全な SCRAM、サーバー署名が検証済み) でアップストリームを認証し、クライアントに対して AuthenticationOk を合成し、ReadyForQuery を通じてサーバーの認証後メッセージを再生します。
MySQL — サーバー グリーティングの CLIENT_SSL 機能ビットをクリアし (MySQL には「拒否 TLS」応答がありません)、エージェントの認証応答を破棄し、独自の mysql_native_password、caching_sha2_password fast-auth、およびキャッシュ ミス時の完全認証パスを送信します。サーバーの RSA 公開キーを要求し、パスワードとスクランブルを XOR し、RSA-OAEP で暗号化します。
MongoDB — OP_MSG 上の SCRAM-SHA-256。SASL 会話を伝送するのに十分な手動でロールされた BSON を使用します。ドライバーはいません。
Redis — エージェントのトラフィックよりも先にある RESP AUTH (存在する場合は ACL ユーザー名付き)。
SMTP — EHLO、STARTTLS、AUTH PLAIN を使用してアップストリーム レッグをアップグレードし、認証済みのセッションで続行できるように、AUTH と STARTTLS を削除して機能リストをクライアントに返します。
ネットワーク上で TLS を義務付けるデータベースの場合、各ブローカーはアップストリーム TLS (STARTTLS 経由の SMTP、接続時の Redis と MongoDB、およびプロトコル内ネゴシエーションを通じて Postgres と MySQL) を選択できます。エージェントとプロキシのレッグは、プロキシが唯一のゲートウェイであるサンドボックスのネットワーク名前空間内で平文のままです。
SCRAM エンジンは Postgres と Mongo の間で共有され、crypto/hmac + crypto/sha256 に基づいて構築されています。SCRAM の Hi() は単一ブロックの PBKDF2 であるため、ライブラリは必要ありませんでした。仲介層全体は golang.org/x/crypto import を実行しません。 JWS 署名者も同様です。
スコープは蓄積されます。ドメインがリクエストに一致するすべてのスコープがインジェクトするため、ゲートウェイは独自のキーと上流プロバイダーのキーの両方を必要とします (ボット)

同じホストをスコープとする - 単一のリクエストで両方を取得します。スコープは配信順に適用されます。 2 人が同じヘッダーを書き込んだ場合、最後の書き込み者が勝ちます。これは意図的なものです。これにより、プロバイダー互換プリセットは、特別なケースを使用せずに、資格情報の上に必須の非機密ヘッダーを重ねることができます。
障害モードは明示的であり、フェールセーフです。
HTTP インジェクションが 401 で失敗します。フェッチ、ミント、または署名エラーが発生した場合、プロキシはそのスコープをスキップし、認証情報なしでリクエストを転送します。アップストリームは 401 でリクエストを拒否します。フォールバックを置き換えたり、部分的な値を漏洩したりすることはありません。
ソケット ブローカリングはハード クローズに失敗します。データベース ハンドシェイクでエラーが発生した場合、接続は半分認証された状態でエージェントに渡されるのではなく、切断されます。
失効は TTL です。フェッチされた値は、構成可能なウィンドウ (デフォルトは 60 秒) に対してワーカー側でキャッシュされます。保存された値をローテーションするか、コネクタの場合は独自のストア内でローテーションすると、実行中のサンドボックスが次のリクエストでその値を取得します。ウィンドウをゼロに設定すると、すべてのリクエストが再フェッチされます。つまり、インスタント キル スイッチです。 (シークレットの再スコープ (新しい宛先または挿入形式) もインプレース更新であり、削除して再作成する必要はありません。)
すべての注入は、マスクされた監査イベント (どの環境変数、どの注入タイプ、ターゲット ヘッダー、宛先ホスト、カウント) を発行しますが、値は発行されません。

[切り捨てられた]

## Original Extract

The full design of Declaw's credential vault: how a secret stays out of the sandbox, gets attached at the egress proxy on the way out, and why a fully compromised, prompt-injected agent still leaks nothing but a placeholder.

Letting Untrusted AI Agents Use Credentials They Can Never Read | Declaw declaw Arena Startups Demos Benchmarks Blog Pricing About Docs Book a Demo Sign In Back to blog 2026-07-07 Security Letting Untrusted AI Agents Use Credentials They Can Never Read
The full design of Declaw's credential vault: how a secret stays out of the sandbox, gets attached at the egress proxy on the way out, and why a fully compromised, prompt-injected agent still leaks nothing but a placeholder.
An AI agent is only useful once you give it secrets — an LLM API key, a database password, a token for an internal service. The usual way to hand one over is an environment variable in the sandbox where the agent runs. That's also the problem: the sandbox is exactly where untrusted, model-driven code executes. A prompt-injected agent — or a compromised dependency — can read its own environment, dump /proc , walk the process table, and your key walks out with it. You've handed the keys to the one component you've already decided not to trust, and sandboxing doesn't change that: a good sandbox bounds the blast radius of the code, but the secret is sitting inside that radius with it.
We covered the practical side — store a secret, scope it, call an API from an agent that never holds the key — in How to Give an AI Agent Code Execution Without Handing Over Your Credentials . This post is the other half: the whole design, end to end — store, injection, isolation — and why a fully compromised agent leaks nothing but a placeholder.
So we stopped trying to find a better hiding spot inside the box, and made the secret never enter the box at all.
The secret value lives server-side in a vault (we use OpenBao). It's write-only: you store it once, and no API ever returns it again — not to you, not to the agent.
The sandbox is created with a reference — just the secret's name. Inside the VM, the environment variable holds a placeholder string, declaw:vault-managed .
Every sandbox's outbound traffic already routes through an egress proxy. When the agent makes a request to a host you've scoped the secret to, the proxy attaches the real credential — after the request has left the VM.
The agent uses the credential against the destination you intend, and never holds it. A prompt-injected agent that dumps its whole environment leaks a placeholder.
What makes this a real boundary and not a demo is where the proxy runs. The agent runs in a Firecracker microVM. The proxy runs in that sandbox's host-side network namespace — outside the guest, across the virtualization boundary — and it is the VM's only route to the network. The secret lives in the proxy's memory for exactly one request and is never written into the guest: not a file, not an env var, not a log line.
How the VM trusts the proxy's certs
The first question any engineer asks: if the proxy terminates TLS to inject into an HTTPS request, why doesn't the agent's client reject the substituted certificate?
Each sandbox gets its own ephemeral CA — a per-sandbox ECDSA P-256 key whose private key never leaves the proxy's memory and is never written to disk. For each upstream host the agent contacts, the proxy mints a short-lived leaf signed by that CA.
For the guest to accept those leaves, the CA's public cert is installed into the guest's trust stores at boot — appended to the system bundle and exported through the variables every major runtime honors:
SSL_CERT_FILE → /etc/ssl/certs/ca-certificates.crt
REQUESTS_CA_BUNDLE → /etc/ssl/certs/ca-certificates.crt
CURL_CA_BUNDLE → /etc/ssl/certs/ca-certificates.crt
NODE_EXTRA_CA_CERTS → /usr/local/share/ca-certificates/declaw-sandbox.crt
So OpenSSL, Python requests, curl, and Node all chain-validate the proxy's leaf with no per-app configuration. The CA is per-sandbox, so one sandbox's CA is worthless against another's traffic.
One limitation falls out of the same mechanism: a client that pins a specific upstream certificate — rather than trusting the system store — will reject the proxy's minted leaf, so those endpoints can't be brokered. It's rare in server-to-server API calls, but it's real.
Injection fires on domain match, not on the env var. The placeholder is a courtesy so the agent's SDK has a value to read; the proxy injects into any request to a scoped host, whether or not the agent ever touched the variable — and it uses Set, so it also overrides any header the agent tried to supply itself.
Scoping, and why it's anchored
A secret is attached only to hosts that match one of its scopes. A scope's domain pattern is an exact host, a *. -wildcard, or a ~ -prefixed regex. Matching runs on RE2 — linear-time, no catastrophic backtracking — and every scope's pattern is anchored to ^(?:…)$ , so a credential scoped to api.example.com can't be coaxed onto api.example.com.evil.com . Anchoring is enforced rather than advisory: the pattern is wrapped (so even a top-level alternation a|b anchors correctly as ^(?:a|b)$ ), and it applies to hand-written custom scopes, not just the built-in presets.
Values go into OpenBao under a KV-v2 mount, one namespace per team — tenant isolation at the store, not just in a query filter. Our own Postgres holds only metadata and the per-domain injection rules; there is no value column anywhere in it. If the metadata write fails after the vault write, we roll the vault write back, so a half-created secret never leaves an orphan.
You don't even have to store the secret with us. A secret can instead be a connector descriptor — a small JSON pointer into a store you already run:
{"provider": "aws_sm", "secret_id": "prod/openai", "region": "us-east-1"}
On a cache miss the worker reads the live value from your store, uses it for that one request, and never persists it. Eleven stores are wired: AWS Secrets Manager and SSM Parameter Store, GCP Secret Manager, Azure Key Vault, HashiCorp Vault / OpenBao, Infisical, Doppler, Kubernetes Secrets, 1Password Connect, CyberArk Conjur, and Akeyless. The connectors reuse the same credential-minting machinery as the brokered injection types below, so a connector can authenticate to your store with workload identity instead of a stored key.
For the common case there's a preset catalog (~39 providers): you pick the provider and paste the key, and the destination host, injection format, and provider quirks are filled in for you — the version header one API requires, the Token vs Key scheme word another wants, the fixed basic-auth username a third pins. Presets expand server-side into ordinary scopes, so they're sugar over the same machinery, not a separate code path.
The static types are the boring-but-essential ones: Authorization: Bearer <v> , an arbitrary header, HTTP basic, or a query parameter — plus four affordances (a value prefix, a fixed basic-auth username, extra non-secret headers, static query params) so a single scope can express a provider's real contract without a new mechanism. All of it is applied in the proxy's request director, after the security scan and PII redaction have run, so the value never reaches the scanner or the audit trail.
The interesting types broker a short-lived credential at egress instead of forwarding a stored one:
AWS SigV4. The agent's SDK signs with throwaway credentials; the proxy strips that signature and re-signs the request. The stored material can be static keys, an AssumeRole (short-lived creds, refreshed before expiry), or — the keyless path — AssumeRoleWithWebIdentity, where the call to STS is sent unsigned because the OIDC web-identity token is the authentication. No long-lived AWS keys stored anywhere.
OIDC. The proxy mints a short-lived bearer worker-side, caches it per (team, secret), and re-mints before expiry. It supports client-credentials, RFC 8693 token-exchange (a projected service-account or CI OIDC token as the subject — e.g. federating into GCP Workload Identity), private_key_jwt, and jwt-bearer, discovering the token endpoint from the issuer's .well-known/openid-configuration . The client secret and the minted token both stay proxy-side; only the resulting Authorization: Bearer reaches upstream.
HMAC. Request signing with a configurable digest (sha1/256/512), encoding (hex/base64), signing-string template over {body}{method}{path}{timestamp} , prefix, and header layout — the GitHub, Stripe, and Slack signature shapes all fall out of the same config.
The JOSE signing behind private_key_jwt (RS256/PS256/ES256) is implemented on the Go standard library. Which brings me to the part I like most.
The agent can talk to Postgres, MySQL, MongoDB, Redis, or an SMTP relay with no password at all. It opens a normal connection; the proxy performs the authentication handshake upstream on its behalf, then steps aside and pipes bytes. The agent gets a working, authenticated session and never holds the password.
That meant implementing the auth side of each wire protocol, on the standard library:
Postgres — read the client's startup message, decline its SSLRequest with N to keep the near leg plain, authenticate upstream with cleartext / MD5 / SCRAM-SHA-256 (full SCRAM, server signature verified), then synthesize AuthenticationOk to the client and replay the server's post-auth messages through ReadyForQuery.
MySQL — clear the CLIENT_SSL capability bit in the server greeting (MySQL has no "decline TLS" reply), discard the agent's auth response, and send our own: mysql_native_password, caching_sha2_password fast-auth, and the full-auth path on a cache miss — request the server's RSA public key, XOR the password with the scramble, and RSA-OAEP encrypt it.
MongoDB — SCRAM-SHA-256 over OP_MSG, with just enough hand-rolled BSON to carry the SASL conversation. No driver.
Redis — a RESP AUTH (with ACL username if present) ahead of the agent's traffic.
SMTP — EHLO, upgrade the upstream leg with STARTTLS, AUTH PLAIN, then return the capability list to the client with AUTH and STARTTLS stripped so it proceeds on the already-authenticated session.
For a database that mandates TLS on the wire, each broker can opt into upstream TLS — SMTP via STARTTLS, Redis and MongoDB at connect time, and Postgres and MySQL through their in-protocol negotiation. The agent↔proxy leg stays cleartext inside the sandbox's network namespace, where the proxy is the only gateway.
The SCRAM engine is shared between Postgres and Mongo and built on crypto/hmac + crypto/sha256 — SCRAM's Hi() is just a single-block PBKDF2, so we didn't need a library for it. The whole brokering layer carries no golang.org/x/crypto import; same for the JWS signer.
Scopes accumulate. Every scope whose domain matches the request injects, so a gateway that needs both its own key and the upstream provider's key — both scoped to the same host — gets both on a single request. Scopes apply in delivery order; if two write the same header, last writer wins. That's deliberate: it's what lets a provider-quirk preset layer a required non-secret header over the credential with no special case.
The failure modes are explicit, and they fail safe:
HTTP injection fails open-to-401. If a fetch, mint, or signature errors, the proxy skips that scope and forwards the request without the credential — the upstream rejects it with a 401. It never substitutes a fallback or leaks a partial value.
Socket brokering fails hard-closed. If a database handshake errors, the connection is dropped rather than handed to the agent half-authenticated.
Revocation is a TTL. Fetched values are cached worker-side for a configurable window (default 60s). Rotate the stored value — or rotate it in your own store, for a connector — and running sandboxes pick it up on their next request. Set the window to zero and every request re-fetches: an instant kill switch. (Re-scoping a secret — a new destination or injection format — is an in-place update too, no delete-and-recreate.)
Every injection emits a masked audit event — which env var, which injection type, the target header, the destination host, and a count — and never the value.

[truncated]
