---
source: "https://github.com/regent-protocol/regent-httpsig"
hn_url: "https://news.ycombinator.com/item?id=49361754"
title: "Show HN: Verify OpenAI's signed agent traffic in Python (RFC 9421/Web Bot Auth)"
article_title: "GitHub - regent-protocol/regent-httpsig: Verify and sign AI agent HTTP traffic in Python — RFC 9421 · Web Bot Auth · AAuth · GitHub"
image: "https://opengraph.githubassets.com/26eaecc1c64b4c1deaeb110bfe005fc5d0f9a9273abd460cb046fa787538187b/regent-protocol/regent-httpsig"
author: "abay_aubakirov"
captured_at: "2026-08-19T14:23:31Z"
capture_tool: "hn-digest"
hn_id: 49361754
score: 2
comments: 0
posted_at: "2026-08-19T14:01:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Verify OpenAI's signed agent traffic in Python (RFC 9421/Web Bot Auth)

- HN: [49361754](https://news.ycombinator.com/item?id=49361754)
- Source: [github.com](https://github.com/regent-protocol/regent-httpsig)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T14:01:41Z

## Translation

タイトル: HN を表示: Python で OpenAI の署名付きエージェント トラフィックを検証する (RFC 9421/Web Bot Auth)
記事のタイトル: GitHub - regent-protocol/regent-httpsig: Python での AI エージェント HTTP トラフィックの検証と署名 — RFC 9421 · Web Bot Auth · AAuth · GitHub
説明: Python での AI エージェント HTTP トラフィックの検証と署名 — RFC 9421 · Web Bot Auth · AAuth - regent-protocol/regent-httpsig

記事本文:
GitHub - regent-protocol/regent-httpsig: Python での AI エージェント HTTP トラフィックの検証と署名 — RFC 9421 · Web Bot Auth · AAuth · GitHub
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
リージェントプロトコル
/
リージェント-httpsig
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
7 コミット 7 コミット フォルダーとファイル
.github/ ワークフロー .github/ ワークフローの例 例 src/ re

gent_httpsig src/ regent_httpsig テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Python で AI エージェントの HTTP トラフィックを検証して署名します。これは、OpenAI が署名し、Cloudflare が検証する方法です。
RFC 9421 · Web ボット認証 · AAuth
OpenAI のエージェントは、行うすべての HTTP リクエストに暗号署名を行います。 Cloudflare、AWS WAF、および
Google はそれらの署名を検証します。このライブラリは、そのハンドシェイクの両側を Python にもたらします。
署名されたエージェントが API にアクセスしていることを確認し、ボット ウォールを防ぐために独自のエージェントのトラフィックに署名します。
それを認識してください。
pip インストール リージェント-httpsig
検証: どの AI エージェントが電話をかけているかを 5 行で確認します
fastapi から FastAPI をインポート
regent_httpsig からインポート HttpsigVerifier
regent_httpsig から。 fastapiインポートアタッチ、SignatureDep、VerifiedSignature
アプリ = FastAPI ()
Attach ( app , HttpsigVerifier ())
@アプリ。投稿 ( "/v1/orders" )
async def create_order ( sig : VerifiedSignature | None = SignatureDep ):
署名の場合:
print ( sig .agent ) # "https://chatgpt.com"
print ( sig . keyid ) # RFC 7638 キーの拇印
...
FastAPI はありませんか?コアにはフレームワークの依存関係はありません。
検証者 = HttpsigVerifier ()
sig = 検証者を待ちます。 verify (メソッド、URL、ヘッダー) #VerifiedSignature |なし
検証はデフォルトでエンリッチメントです: 署名ヘッダーがないためコストはかかりません。
署名は None を生成し、信頼できない入力では何も発生しません。使用する
署名が存在する必要がある場合の regent_httpsig.fastapi.RequiredSignatureDep — 401
エージェントに署名方法を正確に伝えます。
リバースプロキシの背後にあるのですか?エージェントはパブリック URL に署名しました
( https://api.example/… ) ただし、ASGI サーバーでは http://container/… が認識されます。 FastAPI
依存関係は、 X-Forwarded-Proto + Host から署名付き URL を再構築するため、

必ずあなたの
プロキシはスキームを転送します — nginx: proxy_set_header X-Forwarded-Proto $scheme; 。
本番環境で署名の検証が不思議なことに失敗した場合は、まずこれを確認してください。
サイン: エージェントをボットの壁を越えさせます
regent_httpsig から EgressSigner をインポート
署名者 = EgressSigner (シード = os .environ [ "AGENT_KEY_SEED" ],
Signature_agent = "https://myagent.example" )
ヘッダー = 署名者 。 sign ( "POST" , url , { "content-type" : "application/json" })
応答 = httpx 。 post ( url 、 json = body 、 headers = ヘッダー )
1 つのコマンドでキーとすぐに公開できる /.well-known/ ファイルを生成します。
regent-httpsig keygen --agent https://myagent.example --out ./well-known/
https://myagent.example/.well-known/http-message-signatures-directory でディレクトリを公開します。
これで、インターネット上のすべての Web Bot Auth Verifier がエージェントを識別できるようになります。
チェックする
ステータス
RFC 9421 付録 B.2.6 Ed25519 ベクトル (バイト正確)
✅CIで
Web ボット認証ドラフト -05 A.2.2 — ;key= でカバーされた sf-dictionary Signature-Agent
✅ CI¹
Web Bot Auth A.2.3 — 従来の sf-string 形式 (OpenAI が実稼働環境で出荷するもの)
✅CIで
署名 → ラウンドトリップ検証 (新しいキー、完全なパイプライン)
✅CIで
AAuth アイデンティティ モード ラウンドトリップ (aa-agent+jwt + cnf.jwk の所有証明)
✅CIで
aauth-signing によって署名 (jwt スキーム、keyid-less) → 検証済み
✅²
改ざんされたリクエスト / 期限切れの署名 / 間違ったディレクトリ キーが拒否されました
✅CIで
¹ 草案自身の A.2.2 例で印刷された署名バイトは、
ドラフト独自の署名ベース (従来の A.2.3 ベクターと RFC 9421 B.2.6 の両方に存在するため、欠陥は
は例にあるものであり、正規化ではありません）。 Ed25519 は決定論的であるため、テストでは
同じバイト正確なベース上で同じ RFC テスト鍵で再署名されたベクトル - アップストリームで報告されました。
² aauth-signing の j を使用したクロスライブラリ相互運用性

wt スキーム: トークン層、cnf.jwk の証明
所有権と正規化がすべて検証されます。署名者はオプションの keyid を正しく省略します
パラメータ — 基礎となる RFC 9421 ライブラリで読み取られる無条件の keyid を公開します
私たちが今扱っているもの。 1 つの逸脱が aauth-signing の上流で報告されます。
署名バイト シーケンスは Base64url ですが、RFC 8941 では標準の Base64 が必要です。の
keyid-less シェイプは CI に固定されます。
Web Bot Auth (draft-meunier-web-bot-auth-architecture): によるキー検出
{署名エージェント}/.well-known/http-message-signatures-directory 。両方のワイヤ形式
Signature-Agent は受け入れられます - 現在の sf-dictionary と従来のベア sf-string
OpenAIが実際に送信します。
AAuth (draft-hardt-oauth-aauth-protocol、アイデンティティベース モード): エージェントは
署名キーの JWT エージェントトークン ;発行者の JWKS はトークンを検証します。
cnf.jwk はリクエストの署名を検証します。 pip install 'regent-httpsig[aauth]' でインストールします。
-11 編集者のコピーを追跡します。完全に指定されたアルゴリズム (RFC 9864、Ed25519 —
-10 エコシステムの EdDSA の移行フラグ) と人物トークン ( aa-person+jwt 、
HttpsigConfig.resource_url 経由でオプトインします)。
フルプロトコルの AAuth 実装 (両方の役割、すべてのトークン タイプ) については、を参照してください。
christian-posta/aauth-python-library —
このライブラリは、両方の方言を処理するシン依存当事者検証ツールです。
セキュリティ モデル (単純な実装が間違っていること)
検証者は、攻撃者が指定できるオリジン (署名者が署名したもの) から重要なディレクトリを取得します。
リクエストはその Signature-Agent を選択します。 regent-httpsig には、以下のガード レールが付属しています。
デフォルトの SSRF 保護: https のみ、解決されたすべての IP はパブリックである必要があります (キャッチ
169.254.169.254 、ループバック、プライベート範囲、内部サービスへの DNS 名マッピング)、
リダイレクトは追跡されず、応答サイズは c

適用されました。
境界付きキャッシュ: エビクション付きのインスタンスごとの TTL キャッシュ - keyid-spam 攻撃は不可能
記憶を成長させる。失敗はネガティブキャッシュされるため、デッドオリジンを使用して速度を低下させることはできません。
有効な署名は、信頼性ではなく、キーの所有を証明します。 VerifiedSignature.trusted
設定された許可リストのみが反映されます。キーを信頼するかどうかを決定するのはポリシーです
レイヤーの仕事。
基礎となるエコシステムの既知の鋭いエッジ、すでに処理済み: 上流
http-message-signatures ライブラリは RFC 9421 ;key= 辞書メンバーを解決できません (
コンポーネント リゾルバーを提供します)、ASGI の間、大文字と小文字を区別してヘッダー名を検索します。
フレームワークはそれらを小文字にし（ラップします）、typing_extensions の宣言を忘れます（私たちは
宣言します）。
regent_httpsig からインポート HttpsigConfig 、 HttpsigVerifier
verifier = HttpsigVerifier ( HttpsigConfig (
Trusted_agents = frozenset ({ "https://chatgpt.com" , "https://operator.openai.com" }),
max_age_hours = 25 , # これより前に作成された署名を拒否します
ache_ttl = 600 , # キーディレクトリキャッシュ秒数
))
アプリの共有クライアントを渡して、そのプールを再利用します: HttpsigVerifier(http_client=my_async_client) 。
Web Bot Auth と AAuth は IETF ドラフトです (RFC 9421 自体が最終標準です)。私たちは追跡します
草案。ドラフトを破ると、変更は 0.x のままマイナー リリースとしてリリースされます。
Ed25519 は今のところのみ — これはエージェント エコシステムに同梱されているものです。
本文のカバレッジ ( content-digest ) は署名でカバーされている場合に検証されますが、これは
ライブラリはそれを必要としません。ルートごとに必要かどうかを決定します。
Cloudflare/web-bot-auth (TypeScript/Rust) ·
クリスチャン-ポスタ/aauth-python-ライブラリ
(完全な AAuth プロトコル) · pyauth/http-message-signatures
(これは RFC 9421 プリミティブに基づいて構築されています)
Regent Protocol によって実稼働環境で構築され、実戦テスト済み —
AI エージェントのランタイム制御とアイデンティティ。 Apache-2.0。
V

Python で AI エージェントの HTTP トラフィックを検証および署名する — RFC 9421 · Web Bot Auth · AAuth
Readme Apache-2.0 ライセンス セキュリティ ポリシー
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Verify and sign AI agent HTTP traffic in Python — RFC 9421 · Web Bot Auth · AAuth - regent-protocol/regent-httpsig

GitHub - regent-protocol/regent-httpsig: Verify and sign AI agent HTTP traffic in Python — RFC 9421 · Web Bot Auth · AAuth · GitHub
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
regent-protocol
/
regent-httpsig
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
7 Commits 7 Commits Folders and files
.github/ workflows .github/ workflows examples examples src/ regent_httpsig src/ regent_httpsig tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
Verify and sign AI agent HTTP traffic in Python — the way OpenAI signs and Cloudflare verifies.
RFC 9421 · Web Bot Auth · AAuth
OpenAI's agents cryptographically sign every HTTP request they make. Cloudflare, AWS WAF and
Google verify those signatures. This library brings both sides of that handshake to Python:
verify signed agents hitting your API, and sign your own agent's traffic so bot walls
recognize it.
pip install regent-httpsig
Verify: know which AI agent is calling — in 5 lines
from fastapi import FastAPI
from regent_httpsig import HttpsigVerifier
from regent_httpsig . fastapi import attach , SignatureDep , VerifiedSignature
app = FastAPI ()
attach ( app , HttpsigVerifier ())
@ app . post ( "/v1/orders" )
async def create_order ( sig : VerifiedSignature | None = SignatureDep ):
if sig :
print ( sig . agent ) # "https://chatgpt.com"
print ( sig . keyid ) # RFC 7638 key thumbprint
...
No FastAPI? The core has no framework dependencies:
verifier = HttpsigVerifier ()
sig = await verifier . verify ( method , url , headers ) # VerifiedSignature | None
Verification is enrichment by default : no Signature header costs nothing, a bad
signature yields None , and nothing ever raises on untrusted input. Use
regent_httpsig.fastapi.RequiredSignatureDep when a signature must be present — the 401
tells the agent exactly how to sign.
Behind a reverse proxy? The agent signed the public URL
( https://api.example/… ), but your ASGI server sees http://container/… . The FastAPI
dependency rebuilds the signed URL from X-Forwarded-Proto + Host , so make sure your
proxy forwards the scheme — nginx: proxy_set_header X-Forwarded-Proto $scheme; .
If signatures mysteriously fail to verify in production, check this first.
Sign: get your agent past bot walls
from regent_httpsig import EgressSigner
signer = EgressSigner ( seed = os . environ [ "AGENT_KEY_SEED" ],
signature_agent = "https://myagent.example" )
headers = signer . sign ( "POST" , url , { "content-type" : "application/json" })
resp = httpx . post ( url , json = body , headers = headers )
Generate a key and the ready-to-publish /.well-known/ files in one command:
regent-httpsig keygen --agent https://myagent.example --out ./well-known/
Publish the directory at https://myagent.example/.well-known/http-message-signatures-directory
and every Web Bot Auth verifier on the internet can now identify your agent.
Check
Status
RFC 9421 Appendix B.2.6 Ed25519 vector (byte-exact)
✅ in CI
Web Bot Auth draft -05 A.2.2 — sf-dictionary Signature-Agent covered with ;key=
✅ in CI¹
Web Bot Auth A.2.3 — legacy sf-string form ( what OpenAI ships in production )
✅ in CI
Sign → verify roundtrip (fresh keys, full pipeline)
✅ in CI
AAuth identity-mode roundtrip ( aa-agent+jwt + cnf.jwk proof of possession)
✅ in CI
Signed by aauth-signing (jwt scheme, keyid-less) → verified
✅²
Tampered request / expired signature / wrong directory key rejected
✅ in CI
¹ The signature bytes printed in the draft's own A.2.2 example do not verify over the
draft's own signature base (the legacy A.2.3 vector and RFC 9421 B.2.6 both do, so the defect
is in the example, not the canonicalization). Ed25519 is deterministic, so our test pins the
vector re-signed with the same RFC test key over the same byte-exact base — reported upstream.
² Cross-library interop with aauth-signing 's jwt scheme: token layer, cnf.jwk proof of
possession and canonicalization all verify. Its signers correctly omit the optional keyid
parameter — which exposed an unconditional keyid read in the underlying RFC 9421 library
that we now handle. One deviation reported upstream to aauth-signing : it emits the
Signature byte sequence as base64url, while RFC 8941 requires standard base64. The
keyid-less shape is pinned in CI.
Web Bot Auth ( draft-meunier-web-bot-auth-architecture ): key discovery via
{Signature-Agent}/.well-known/http-message-signatures-directory . Both wire forms of
Signature-Agent are accepted — the current sf-dictionary and the legacy bare sf-string
OpenAI actually sends.
AAuth ( draft-hardt-oauth-aauth-protocol , identity-based mode): the agent carries a
JWT agent_token in Signature-Key ; the issuer's JWKS verifies the token, the token's
cnf.jwk verifies the request signature. Install with pip install 'regent-httpsig[aauth]' .
Tracks the -11 editor's copy : fully-specified algorithms (RFC 9864, Ed25519 — with a
transition flag for the -10 ecosystem's EdDSA ) and person tokens ( aa-person+jwt ,
opt-in via HttpsigConfig.resource_url ).
For a full-protocol AAuth implementation (both roles, all token types) see
christian-posta/aauth-python-library —
this library is the thin relying-party verifier that handles both dialects.
Security model (what a naive implementation gets wrong)
The verifier fetches key directories from attacker-nameable origins — whoever signs a
request chooses its Signature-Agent . regent-httpsig ships with the guard rails on:
SSRF protection by default : https-only, every resolved IP must be public (catches
169.254.169.254 , loopback, private ranges, DNS names mapping to internal services),
redirects never followed, responses size-capped.
Bounded caching : per-instance TTL cache with eviction — a keyid-spam attack can't
grow memory; failures are negative-cached so a dead origin can't be used to slow you down.
A valid signature proves key possession — not trustworthiness. VerifiedSignature.trusted
reflects only your configured allow-list; deciding whether to trust a key is your policy
layer's job.
Known sharp edges of the underlying ecosystem, already handled: the upstream
http-message-signatures library cannot resolve RFC 9421 ;key= dictionary members (we
provide the component resolver), it looks up header names case-sensitively while ASGI
frameworks lowercase them (we wrap), and it forgets to declare typing_extensions (we
declare it).
from regent_httpsig import HttpsigConfig , HttpsigVerifier
verifier = HttpsigVerifier ( HttpsigConfig (
trusted_agents = frozenset ({ "https://chatgpt.com" , "https://operator.openai.com" }),
max_age_hours = 25 , # reject signatures created earlier than this
cache_ttl = 600 , # key-directory cache seconds
))
Pass your app's shared client to reuse its pool: HttpsigVerifier(http_client=my_async_client) .
Web Bot Auth and AAuth are IETF drafts (RFC 9421 itself is a final standard). We track
the drafts; breaking draft changes land as minor releases while we're 0.x.
Ed25519 only for now — it's what the agent ecosystem ships.
Body coverage ( content-digest ) is verified when covered by the signature, but this
library does not require it; decide per-route whether you need it.
cloudflare/web-bot-auth (TypeScript/Rust) ·
christian-posta/aauth-python-library
(full AAuth protocol) · pyauth/http-message-signatures
(the RFC 9421 primitive this builds on)
Built and battle-tested in production by Regent Protocol —
runtime control and identity for AI agents. Apache-2.0.
Verify and sign AI agent HTTP traffic in Python — RFC 9421 · Web Bot Auth · AAuth
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
