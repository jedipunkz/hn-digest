---
source: "https://github.com/authai-io/authai"
hn_url: "https://news.ycombinator.com/item?id=48477326"
title: "Show HN: AuthAI, an open-source relay for user-authorized AI sessions"
article_title: "GitHub - authai-io/authai · GitHub"
author: "riccardoio"
captured_at: "2026-06-10T16:21:29Z"
capture_tool: "hn-digest"
hn_id: 48477326
score: 1
comments: 0
posted_at: "2026-06-10T14:56:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AuthAI, an open-source relay for user-authorized AI sessions

- HN: [48477326](https://news.ycombinator.com/item?id=48477326)
- Source: [github.com](https://github.com/authai-io/authai)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T14:56:54Z

## Translation

タイトル: Show HN: AuthAI、ユーザー承認 AI セッション用のオープンソース リレー
記事タイトル: GitHub - authai-io/authai · GitHub
説明: GitHub でアカウントを作成して、authai-io/authai の開発に貢献します。
HN テキスト: こんにちは、HN。私の名前は Riccardo です。インディーズ ハッカー向けに AuthAI を作成しました。アイデアは非常にシンプルです。エンド ユーザーが chatgpt/grok/copilot アカウントに接続し、AI サブスクリプションを通じて AI リクエストをルーティングできるようにします。これにより、ビジネス モデルやユニット エコノミクスが必ずしも意味をなさない多くの新しいクールなアイデアが可能になります。フローは簡単です。「AI でログイン」をクリックし、プロバイダーを選択し、プロバイダーの Web サイトでデバイスを認証します。トークンはユーザーごとの AES-256-GCM 暗号化キーを使用して暗号化されます。このキーはサーバー側には保存されず、ユーザーの JWT セッション内にのみ存在します。セキュリティ モデル全体は、Web サイト/github で見つけることができます。デモは次のとおりです: https://demo.authai.io 開発者の観点から見ると、目的は OpenAI SDK にできるだけ近い状態を維持することです: ```ts
const openai = new OpenAI({
APIキー: jwt、
ベースURL: " https://relay.authai.io/v1 "、
デフォルトヘッダー: {
"x-authai-secret": process.env.AUTH_AI_SECRET,
}、
});
``` また、接続フローを処理するための React SDK もあります。 * MIT ライセンスを取得しており、完全にオープンソースであり、ホストされたリレーが利用可能ですが、スタック全体は自己ホスト可能です。 GitHub リポジトリ: https://github.com/authai-io/authai 生成したアプリケーションやサイド プロジェクトにこのようなものを使用しますか?他に何を追加できますか?

記事本文:
GitHub - authai-io/authai · GitHub
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
オータイオ
/
オータイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

137 コミット 137 コミット .github .github apps apps docs docs package package scripts scripts .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
アプリビルダーの場合は、ChatGPT、Grok、または Copilot を使用してサインインします。
AuthAI を楽しんでいますか?さらに野心的な計画が進行中です。お問い合わせください。
エンドユーザーが既存の ChatGPT、Grok、または GitHub Copilot サブスクリプションを使用して AI 機能の料金を支払えるようにします。ユーザーは OAuth 経由でサインインし、アプリがユーザーに代わってモデルを呼び出し、コストはプラン内に残ります。 React コンポーネントをドロップし、公式の openai SDK をリレーに向ければ完了です。
使用方法は 2 つあります。authai.io でホストされているサービス (無料、セットアップなし) または自分でリレーをセルフホストする方法です。
実験的。 AuthAI は、各プロバイダーのパブリック デバイス コード OAuth フローを使用します。これは、公式 CLI が使用するものと同じです。これらのサーフェスは非公式であり、プロバイダーが変更する可能性があります。 OpenAI、GitHub、xAI とは無関係です。個人的なプロジェクトやデモに使用します。
最速の道。これを新しいプロジェクトで実行します。
npx authai-cloud init
# → ブラウザで authai.io を開き、GitHub にサインインします。
# アプリを作成し、AUTH_AI_SECRET=... を .env に書き込みます
次に、バックエンドで次のようにします。
「openai」から OpenAI をインポートします。
const openai = 新しい OpenAI ( {
apiKey : jwt , // フロントエンドの @authai -io/react から
BaseURL : "https://relay.authai.io/v1" ,
defaultHeaders : { "x-authai-secret" : プロセス。環境 。 AUTH_AI_SECRET ! } 、
} ) ;
AUTH_AI_SECRET はアプリごとの資格情報です。サーバー側に保持し、ブラウザーに送信したり、コミットしたりしないでください。リレーはハッシュのみを保存します。紛失した場合は、アプリを取り消して、新しいアプリを作成してください。
オート

hAI Cloud は無料でレート制限があり、セルフホスト型リレーと同じコードを実行します。暗号文のみを保存します。レコードごとの AES キーは各ユーザーの JWT 内にのみ存在し、復号可能な形式でリレーのサーバーに到達することはありません。自分で実行してみませんか?以下のリレーの自己ホストを参照してください。
ChatGPT に表示されます。 Grok と Copilot は同じパターンに従い、それぞれ独自のプロバイダー固有のデバイスコード OAuth フローを使用します。
エンドユーザーのブラウザ
│「ChatGPTでサインイン」でサインイン → JWTを受け取る
│ JWT をバックエンドに送信しますが、通常は認証を送信します
▼
既存のバックエンド (api.example.com)
│ 新しい OpenAI({ apiKey: jwt, BaseURL:relayUrl + "/v1" })
│ すでに行っているように、openai.chat.completions.create(...) を使用します
▼
AuthAI リレー (relay.authai.io または独自のホスト)
│ JWT を検証し、JWT 内のキーを使用してユーザーの OpenAI トークンを復号化します
│ チャット完了 → コーデックス応答を変換し、chatgpt.com/backend-api を呼び出します
│ 必要に応じてサーバー側でトークンを更新し、適切な場所で再暗号化します
▼
ChatGPT サブスクリプション (ユーザーが支払う)
セキュリティモデル
リレーは、各ユーザーの OAuth トークンをレコードごとの新しい AES-256 キーで暗号化します。キーはサーバー側で永続化されることはなく、クライアントに発行される JWT に埋め込まれます。資格情報を使用するには、暗号化された BLOB (ディスク上) とキー (JWT 内) の両方が必要です。
AuthAI Cloud を使用する代わりにデータ プレーンを自分で実行したい場合:
git clone https://github.com/authai-io/authai.git && cd authai
pnpmインストール
cat > apps/relay-server/.env << EOF
AUTH_AI_JWT_SECRET= $( openssl rand -hex 32 )
AUTH_AI_ORIGINATOR=私のアプリ
AUTH_AI_DB_URL=./relay.db
終了後
pnpm 開発:リレー
# http://localhost:3000 でリッスンする AuthAI リレー
AUTH_AI_ORIGINATOR は、サインイン中に ChatGPT 同意画面に表示される名前です。代わりに、フロントエンド SDK を http://localhost:3000 に指定します。

https://relay.authai.io の。 「統合」セクションのその他の項目はすべて同じように機能します。
2 つの統合パス。シングルトン パスは、クライアント SPA に推奨されるデフォルトです。プロバイダー パスは SSR (Next.js、Remix) およびマルチテナント用です。
import {configureAuthAI , SignIn , useAuthAI } from "@authai-io/react" ;
// モジュールスコープで一度呼び出します。プロバイダツリーはありません。
configureAuthAI ( {
リレーUrl : "https://your-relay.com" ,
appName : "私のアプリ" 、
} ) ;
関数アプリ ( ) {
const {jwt、isSignedIn、signOut} = useAuthAI();
if (!isSignedIn) return <SignIn> AI でサインイン </SignIn> ;
// 通常は auth を送信しますが、`jwt` をバックエンドに送信します
}
useAuthAI() と <SignIn> はツリー内のどこでも機能します。ラッパーは必要ありません。サインイン ダイアログは、最初の使用時にポータル経由で自動的にマウントされます。
// app/layout.tsx — Next.js アプリルーター
"next/headers" から {cookie} をインポートします。
import { AuthAIProvider } から "@authai-io/react" ;
デフォルトの非同期関数をエクスポートします。 Layout ( { Children } ) {
const jwt = (await cookies() ) 。 get ( "authai-jwt" ) ?。価値 ？？ null ;
戻る (
< AuthAIプロバイダー
リレーUrl = {プロセス。環境 。 NEXT_PUBLIC_AUTHAI_RELAY ! }
appName = "私のアプリ"
初期Jwt = { jwt }
ストレージ = "クッキー"
>
{子供たち}
</ AuthAIProvider >
) ;
}
InitialJwt は SSR ハンドオフです。どこからでも (Cookie、NextAuth セッション、カスタム ヘッダー) JWT を渡すと、最初のレンダリングが正しくサインインされます。 storage="cookie" は JWT を Cookie にミラーリングして、サーバー コンポーネントがそれを読み取れるようにします。完全なデモは apps/demo-nextjs にあります。
SDK は JWT のみを公開します。 client.chat() メソッドや openai のラッパーはありません。モデルの呼び出しは、既に使用しているパッケージを使用してバックエンドで行われます。
「openai」から OpenAI をインポートします。
// jwt は受信リクエスト (ヘッダー、Cookie など) から取得されます。
const openai = 新しい OpenAI ( {
APIキー: jwt 、
ベースURL:

プロセス 。環境 。 AUTH_AI_RELAY_URL + "/v1" 、
} ) ;
const stream = openai を待ちます。チャット 。完成品。作成 ( {
モデル：「gpt-5.4」、
メッセージ : [ { 役割 : "ユーザー" 、コンテンツ : "こんにちは" } ] 、
ストリーム : true 、
} ) ;
for await (ストリームのconstチャンク) {
プロセスを標準出力 。 write (チャンク . 選択肢 [ 0 ] ?. デルタ ?. コンテンツ ?? "" ) ;
}
既存の openai ベースのコードはそのまま動作し続けます。 2 つのコンストラクター フィールドを交換します。
Codex バックエンドは、OpenAI の API のサブセットのみを受け入れます。
モデル
使用する
gpt-5.4
デフォルト。スピードと品質のバランスが良い
gpt-5.4-mini
より安く/より速く
gpt-5.4-プロ
より高品質
gpt-5.4-コーデックス
コーディング調整済み
gpt-5.5
最新の
gpt-5.5-プロ
最新の最上位層
サポートされていないモデル名 (例: gpt-4 ) を送信すると、Codex バックエンドから 400 が返されます。
変数
デフォルト
注意事項
AUTH_AI_JWT_SECRET
必須
32 バイト以上の 16 進数 ( openssl rand -hex 32 を使用)
AUTH_AI_ORIGINATOR
必須
ChatGPTの同意画面に表示されます
AUTH_AI_DB_DRIVER
スクライト
sqlite (デフォルト) または postgres
AUTH_AI_DB_URL
./relay.db
SQLite ファイル パス、または Postgres ドライバーの postgres://...
AUTH_AI_PORT
3000
AuthAIクラウドアーキテクチャ
ホストされたサービスは 2 つのドメインで実行されます。
authai.io — Next.js Webアプリ。ランディング ページ、GitHub サインイン、ダッシュボード、ドキュメント ビューア、CLI ブリッジ。ここでアプリを管理します。
lay.authai.io — ほのリレー。純粋なデータプレーン。エンド ユーザーがサインインするエンドポイントと、バックエンドがモデル呼び出しにヒットするエンドポイント。
どちらもセルフホステッド リレーと同じコードを実行します。完全なアーキテクチャについては、docs/reference.md を参照してください。
パッケージ/
§── リレーコア: OAuth フロー、JWT、AES-GCM、OpenAI 互換プロキシ
§──relay-store-sqlite デフォルトの SQLite ストレージ ドライバー
§──relay-store-postgres Postgresドライバ（クラウド版はこれを使用）
§── クラウド クラウドエディション: テナント、管理 API、キルスイッチ、レート制限
§── cli npx authai-cloud — one-comm

そしてアプリの登録
§── 反応configureAuthAI()、<AuthAIProvider>、<SignIn>、useAuthAI()、cookieAdapter
└── サーバー authai.session()、decodeAuthAIToken() — バックエンド ヘルパー
アプリ/
§── コミュニティ (セルフホスト) リレーを起動するリレーサーバー実行可能ファイル
§── クラウド エディションのリレーを起動する、cloud-relay-server 実行可能ファイル (Hetzner+Dokku)
§── Cloud-web AuthAI Cloud 用 Next.js Web アプリ (Hetzner+Dokku) — ランディング、サインイン、ダッシュボード、ドキュメント ビューア、CLI ブリッジ
§── デモ-バックエンド リレーに対して openai SDK を使用した小さなノードのデモ
§── デモ-react Vite + React SPA デモ (シングルトン パス)
━──demo-nextjs Next.js App Router デモ (プロバイダー + SSR パス)
デモをエンドツーエンドで実行する
pnpm dev:リレー # → :3000
pnpm dev:demo:backend # → :4000 (リレーに対して openai npm パッケージを使用します)
pnpm dev:demo # → :5173 (Vite + React)
http://localhost:5173 を開き、 ChatGPT でサインイン をクリックし、 auth.openai.com/codex/device で認証して、メッセージを送信します。
pnpmテスト
29 の単体テストは、暗号化プリミティブ、JWT 発行/検証、チャット完了 ↔ コーデックス応答変換層をカバーします。上位レベルの配線 (認証ルート、SQLite ストア CRUD、リフレッシュ フロー) はデモによって実行されますが、自動テストではまだカバーされていません。
フレームワークに依存しない @authai-io/web SDK (バニラ / Web コンポーネント)
HttpOnly Cookie モード (リレー側セッション エンドポイント)
すでに出荷されています: Postgres ストレージ ドライバー ( @authai-io/relay-store-postgres )、マルチテナント ダッシュボードとブランドの同意オリジネーターを備えた AuthAI クラウド ( https://authai.io + https://relay.authai.io )。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
2
v0.1.1 — authai-cloud CLI での utf-8 文字セットの修正
最新の
2026 年 6 月 9 日
+ 1 リリース
パッケージ
0
そこw

ロード中のエラーとして。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to authai-io/authai development by creating an account on GitHub.

Hello HN, My name is Riccardo and I created AuthAI for indie hackers. The idea is quite simple: let the end users connect their chatgpt/grok/copilot account and route the AI requests through their AI subscriptions. This enable a lot of new cool ideas where the business model/unit economics don't always make sense. The flow is straightforward: They click on "login with AI", choose their provider, and authorise the device on the provider's website. Tokens get encrypted using a per-user AES-256-GCM encryption key, which isn't stored anywhere server-side and only exists within the user's JWT session. The whole security model can be found on the website/github. Here is a demo: https://demo.authai.io From a developer's perspective, the objective is to stay as close to the OpenAI SDK as possible: ```ts
const openai = new OpenAI({
apiKey: jwt,
baseURL: " https://relay.authai.io/v1 ",
defaultHeaders: {
"x-authai-secret": process.env.AUTH_AI_SECRET,
},
});
``` Also, there is a React SDK for handling the connection flow. * It's MIT licensed and completely open-source, there's a hosted relay available, however, the entire stack is self-hostable. GitHub repo: https://github.com/authai-io/authai Would you use something like this for your generated applications and side projects? What else could I add?

GitHub - authai-io/authai · GitHub
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
authai-io
/
authai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
137 Commits 137 Commits .github .github apps apps docs docs packages packages scripts scripts .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Sign in with ChatGPT, Grok, or Copilot — for app builders.
Enjoying AuthAI? Something more ambitious is in the works — get in touch .
Let your end users pay for AI features with their existing ChatGPT, Grok, or GitHub Copilot subscription. They sign in via OAuth, your app calls models on their behalf, the cost stays on their plan. Drop in a React component, point the official openai SDK at the relay, done.
Two ways to use it: the hosted service at authai.io (free, no setup) or self-host the relay yourself.
Experimental. AuthAI uses each provider's public device-code OAuth flow — the same one their official CLIs use. These surfaces are unofficial and providers can change them. Not affiliated with OpenAI, GitHub, or xAI. Use for personal projects and demos.
The fastest path. Run this in a fresh project:
npx authai-cloud init
# → opens authai.io in your browser to sign in with GitHub,
# create an app, and writes AUTH_AI_SECRET=... to .env
Then in your backend:
import OpenAI from "openai" ;
const openai = new OpenAI ( {
apiKey : jwt , // from @authai -io/react on the frontend
baseURL : "https://relay.authai.io/v1" ,
defaultHeaders : { "x-authai-secret" : process . env . AUTH_AI_SECRET ! } ,
} ) ;
AUTH_AI_SECRET is a per-app credential — keep it server-side, never ship it to the browser, never commit it. The relay stores only a hash; if you lose it, revoke the app and create a new one.
AuthAI Cloud is free, rate-limited, and runs the same code as the self-hosted relay. It stores only ciphertext — per-record AES keys live exclusively in each user's JWT and never reach the relay's servers in a decryptable form. Want to run it yourself? See Self-host the relay below.
Shown for ChatGPT. Grok and Copilot follow the same pattern — each uses its own provider-specific device-code OAuth flow.
end-user browser
│ signs in with "Sign in with ChatGPT" → receives a JWT
│ sends JWT to your backend however it normally sends auth
▼
your existing backend (api.example.com)
│ new OpenAI({ apiKey: jwt, baseURL: relayUrl + "/v1" })
│ uses openai.chat.completions.create(...) as you already do
▼
AuthAI relay (relay.authai.io or your own host)
│ verifies JWT, decrypts the user's OpenAI tokens using the key in the JWT
│ translates Chat Completions → Codex Responses, calls chatgpt.com/backend-api
│ refreshes tokens server-side if needed, re-encrypts in place
▼
ChatGPT subscription (user pays)
Security model
The relay encrypts each user's OAuth tokens with a fresh per-record AES-256 key. The key is never persisted server-side — it's embedded in the JWT issued to the client. Both the encrypted blob (on disk) and the key (in the JWT) are required to use the credential.
If you'd rather run the data plane yourself instead of using AuthAI Cloud:
git clone https://github.com/authai-io/authai.git && cd authai
pnpm install
cat > apps/relay-server/.env << EOF
AUTH_AI_JWT_SECRET= $( openssl rand -hex 32 )
AUTH_AI_ORIGINATOR=my-app
AUTH_AI_DB_URL=./relay.db
EOF
pnpm dev:relay
# AuthAI relay listening on http://localhost:3000
AUTH_AI_ORIGINATOR is the name shown on the ChatGPT consent screen during sign-in. Point your frontend SDK at http://localhost:3000 instead of https://relay.authai.io . Everything else in the Integrate section works identically.
Two integration paths. The singleton path is the recommended default for client SPAs; the provider path is for SSR (Next.js, Remix) and multi-tenant.
import { configureAuthAI , SignIn , useAuthAI } from "@authai-io/react" ;
// Call once, at module scope. No provider tree.
configureAuthAI ( {
relayUrl : "https://your-relay.com" ,
appName : "My App" ,
} ) ;
function App ( ) {
const { jwt , isSignedIn , signOut } = useAuthAI ( ) ;
if ( ! isSignedIn ) return < SignIn > Sign in with AI </ SignIn > ;
// send `jwt` to your backend however you normally send auth
}
useAuthAI() and <SignIn> work anywhere in the tree — no wrapper required. The sign-in dialog auto-mounts via portal on first use.
// app/layout.tsx — Next.js App Router
import { cookies } from "next/headers" ;
import { AuthAIProvider } from "@authai-io/react" ;
export default async function Layout ( { children } ) {
const jwt = ( await cookies ( ) ) . get ( "authai-jwt" ) ?. value ?? null ;
return (
< AuthAIProvider
relayUrl = { process . env . NEXT_PUBLIC_AUTHAI_RELAY ! }
appName = "My App"
initialJwt = { jwt }
storage = "cookie"
>
{ children }
</ AuthAIProvider >
) ;
}
initialJwt is the SSR hand-off: pass a JWT from anywhere (cookie, NextAuth session, custom header) and the first render is correctly signed-in. storage="cookie" mirrors the JWT to a cookie so server components can read it. Full demo in apps/demo-nextjs .
The SDK only exposes the JWT. There's no client.chat() method, no wrapper around openai — model calls happen in your backend, using the package you already use.
import OpenAI from "openai" ;
// jwt comes from your incoming request (header, cookie, etc.)
const openai = new OpenAI ( {
apiKey : jwt ,
baseURL : process . env . AUTH_AI_RELAY_URL + "/v1" ,
} ) ;
const stream = await openai . chat . completions . create ( {
model : "gpt-5.4" ,
messages : [ { role : "user" , content : "hi" } ] ,
stream : true ,
} ) ;
for await ( const chunk of stream ) {
process . stdout . write ( chunk . choices [ 0 ] ?. delta ?. content ?? "" ) ;
}
Any existing openai -based code keeps working as-is. You swap two constructor fields.
The Codex backend only accepts a subset of OpenAI's API:
Model
Use
gpt-5.4
Default. Good balance of speed and quality
gpt-5.4-mini
Cheaper / faster
gpt-5.4-pro
Higher quality
gpt-5.4-codex
Coding-tuned
gpt-5.5
Newest
gpt-5.5-pro
Newest, top tier
Sending an unsupported model name (e.g. gpt-4 ) returns a 400 from the Codex backend.
Variable
Default
Notes
AUTH_AI_JWT_SECRET
required
32+ bytes hex (use openssl rand -hex 32 )
AUTH_AI_ORIGINATOR
required
Shown on the ChatGPT consent screen
AUTH_AI_DB_DRIVER
sqlite
sqlite (default) or postgres
AUTH_AI_DB_URL
./relay.db
SQLite file path, or postgres://... for the Postgres driver
AUTH_AI_PORT
3000
AuthAI Cloud architecture
The hosted service runs on two domains:
authai.io — Next.js webapp. Landing page, GitHub sign-in, dashboard, docs viewer, CLI bridge. You manage apps here.
relay.authai.io — Hono relay. Pure data plane. The endpoint your end users sign in against and your backend hits for model calls.
Both run the same code as the self-hosted relay. See docs/reference.md for the full architecture.
packages/
├── relay core: OAuth flow, JWT, AES-GCM, OpenAI-compat proxy
├── relay-store-sqlite default SQLite storage driver
├── relay-store-postgres Postgres driver (cloud edition uses this)
├── cloud cloud edition: tenant, admin API, kill switch, rate limits
├── cli npx authai-cloud — one-command app registration
├── react configureAuthAI(), <AuthAIProvider>, <SignIn>, useAuthAI(), cookieAdapter
└── server authai.session(), decodeAuthAIToken() — backend helpers
apps/
├── relay-server executable that boots the community (self-hosted) relay
├── cloud-relay-server executable that boots the cloud edition's relay (Hetzner+Dokku)
├── cloud-web Next.js webapp for AuthAI Cloud (Hetzner+Dokku) — landing, sign-in, dashboard, docs viewer, CLI bridge
├── demo-backend tiny Node demo using the openai SDK against the relay
├── demo-react Vite + React SPA demo (singleton path)
└── demo-nextjs Next.js App Router demo (provider + SSR path)
Run the demo end-to-end
pnpm dev:relay # → :3000
pnpm dev:demo:backend # → :4000 (uses the openai npm package against the relay)
pnpm dev:demo # → :5173 (Vite + React)
Open http://localhost:5173 , click Sign in with ChatGPT , authorize at auth.openai.com/codex/device , send a message.
pnpm test
29 unit tests cover the encryption primitives, JWT issue/verify, and the Chat Completions ↔ Codex Responses translation layer. Higher-level wiring (auth routes, SQLite store CRUD, refresh flow) is exercised by the demo but not yet covered by automated tests.
Framework-agnostic @authai-io/web SDK (vanilla / web component)
HttpOnly cookie mode (relay-side session endpoint)
Already shipped: Postgres storage driver ( @authai-io/relay-store-postgres ), AuthAI Cloud ( https://authai.io + https://relay.authai.io ) with multi-tenant dashboard and branded consent originators.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
v0.1.1 — utf-8 charset fix in authai-cloud CLI
Latest
Jun 9, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
