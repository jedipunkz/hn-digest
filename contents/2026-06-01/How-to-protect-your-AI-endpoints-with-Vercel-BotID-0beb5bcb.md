---
source: "https://vercel.com/kb/guide/protect-ai-endpoints-with-vercel-botid"
hn_url: "https://news.ycombinator.com/item?id=48342146"
title: "How to protect your AI endpoints with Vercel BotID"
article_title: "How to protect your AI endpoints with Vercel BotID | Vercel Knowledge Base"
author: "flashbrew"
captured_at: "2026-06-01T00:35:41Z"
capture_tool: "hn-digest"
hn_id: 48342146
score: 2
comments: 0
posted_at: "2026-05-31T01:06:00Z"
tags:
  - hacker-news
  - translated
---

# How to protect your AI endpoints with Vercel BotID

- HN: [48342146](https://news.ycombinator.com/item?id=48342146)
- Source: [vercel.com](https://vercel.com/kb/guide/protect-ai-endpoints-with-vercel-botid)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T01:06:00Z

## Translation

タイトル: Vercel BotID を使用して AI エンドポイントを保護する方法
記事のタイトル: Vercel BotID を使用して AI エンドポイントを保護する方法 |ヴァーセルナレッジベース
説明: Vercel BotID と checkBotId() を使用して AI エンドポイントへのすべてのリクエストをゲートし、検証された呼び出し元に対してのみ推論が実行されるようにします。セットアップ、詳細な分析、検証済みのボットについて説明します。

記事本文:
Vercel BotID を使用して AI エンドポイントを保護する方法 | Vercel ナレッジベース コンテンツにスキップ Vercel BotID を使用して AI エンドポイントを保護する方法
Vercel BotID と checkBotId() を使用して AI エンドポイントへのすべてのリクエストをゲートし、検証された呼び出し元に対してのみ推論が実行されるようにします。セットアップ、詳細な分析、検証済みのボットについて説明します。
ナレッジ ベース / AI URL をコピー ページをコピー このページについて AI に質問する 最終更新日 2026 年 5 月 31 日 Vercel BotID を使用すると、推論を実行する前に、AI エンドポイントへの各リクエストが実際のブラウザーからのものであることを検証できます。目に見えない CAPTCHA として機能し、保護するルート上のリクエストにクライアント側のチャレンジを添付し、サーバー側の checkBotId() 呼び出しでそれぞれを分類するため、自動化されたクライアントはモデルに到達する前に拒否されます。セッションごとに 1 回ではなく、リクエストごとに実行するということは、攻撃者が一度だけバイパスして、何千もの盗まれた呼び出しにわたってそのアクセスを再利用することができないことを意味します。この違いこそが、推論が横ばいのままで組織的な攻撃によってトラフィックが 3,000% 急増したときでも、Nous Research のチャット アプリをオンラインに保った理由です。
このガイドでは、BotID をインストールし、クライアントで AI ルートを宣言し、サーバー上で checkBotId() を使用してそのルートをゲートし、検証されたリクエストに対してのみ推論が実行されるようにする手順を説明します。また、ルートごとに検出レベルを設定し、最も価値の高いエンドポイントでの詳細な分析と他の場所での基本的なチェックを有効にし、Vercel WAF バイパス ルールを使用して正当な自動化を実現する方法を学びます。
Vercel にデプロイされた JavaScript プロジェクト
AI SDK で構築されたルートなどのフロントエンド リクエストを受け入れる AI エンドポイント
Deep Analysis を使用するための Pro または Enterprise プラン (Basic はすべてのプランで利用可能)
ターミナル npm i botid 2. プロキシの書き換えを構成する
Next.js 構成を withBotId でラップします。これによりプロキシ リライトが設定され、広告が

-ブロッカーやサードパーティのスクリプトは BotID の保護を弱めることはできません。
next.config.ts import { withBotId } から 'botid/next/config' ;
const nextConfig = { // 既存の Next.js 構成 } ;
デフォルトの withBotId ( nextConfig ) をエクスポートします。 Nuxt、SvelteKit、およびその他のフレームワークの場合、セットアップは同様のパターンに従います。フレームワークごとのバージョンについては、BotID スタート ガイドを参照してください。
3. クライアント上で AI ルートを宣言します
クライアントの初期化中に initBotId() を呼び出し、保護する AI ルートをリストします。 BotID はこのリストを使用して、一致するリクエストにチャレンジ ヘッダーを添付します。ここでルートが宣言されていない場合、そのリクエストはこれらのヘッダーなしで到着するため、checkBotId() は検証するものが何もなく、それらをボットとして扱います。
Next.js 15.3 以降の場合は、instrumentation-client.ts を使用します。
;instrumentation-client.ts import { initBotId } from 'botid/client/core' ;
initBotId ( { 保護 : [ { パス : '/api/chat' 、メソッド : 'POST' 、 } 、 ] 、 } ) ; Next.js の以前のバージョンでは、代わりにルート レイアウト ヘッドに <BotIdClient /> コンポーネントをマウントし、同じ保護配列を渡します。
4. サーバー上のすべてのリクエストを確認する
AI 呼び出しが実行される前に、ルート ハンドラー内で checkBotId() を呼び出します。これは負荷のかかるステップです。現在処理されているリクエストの分類を返すため、ブロックされたリクエストがモデルに到達することはありません。
app/api/chat/route.ts import { checkBotId } from 'botid/server' ; import { NextRequest , NextResponse } から 'next/server' ;
エクスポート非同期関数 POST (リクエスト: NextRequest) { const verify = await checkBotId() ;
if (検証 .isBot ) { NextResponse を返します。 json ( { エラー : 'アクセスが拒否されました' } , { ステータス : 403 } ) ; }
// 推論は検証が成功した後にのみ実行されます const body = await request 。 json() ; const result = await runInference (body) ;
NextRe を返す

スポンジ。 json ( { 結果 } ) ; }
async function runInference ( body :unknown ) { // ここで AI SDK またはモデルを呼び出します return {output : 'response' } ; runInference の前にチェックを入れると、検証されたリクエストに対してのみ推論コストが発生することになります。
基本的な検証は、多くのそれほど洗練されていないボットを捕捉し、すべてのプランで無料で実行されます。高価値の AI ルートの場合は、Kasada を利用した機械学習モデルを使用して数千のクライアント側信号を分析する深層分析を有効にします。
深層分析はリアルタイムで学習して適応するため、最初は正当なトラフィックのように見える組織的な攻撃を検出できます。あるインシデントでは、プロキシ ノード間を循環する同一のブラウザ フィンガープリントを相関させることで、新しいボット ネットワークへの 500% のトラフィック スパイクを追跡しました。その後、手動介入なしで、約 10 分以内にこれらのセッションが再分類されてブロックされました。完全な内訳については、BotID Deep Analysis が高度なボット ネットワークをリアルタイムでどのように捕捉したかをご覧ください。
プロジェクト設定の [ボット管理] ページにアクセスし、[構成] ボタンをクリックして構成設定を開き、詳細分析を有効にします。
この機能は、Pro および Enterprise プランのすべての顧客が利用できます。 checkBotId() を呼び出すリクエストのみが課金され、パッシブなページビューは課金されません。
推論の前にチェックを実行する : ハンドラーでモデル呼び出しの前に checkBotId() を保持しておくと、ブロックされたリクエストによってトークンが消費されることがなくなります。
ルートごとに検出レベルを設定する : AdvancedOptions.checkLevel を使用して、最も機密性の高いルートとその他の基本的なルートに deepAnalysis を適用します。 checkLevel は、各ルートのクライアントとサーバーの構成で同一である必要があります。同一でない場合、検証は失敗します。これは botid@1.4.5 以降で利用できます。
信頼できるエージェントに検証済みのボットを許可する
isBot のみに基づいてブロックすると、正規の自動ブロックもブロックされます

クローラー (Googlebot など) や AI アシスタント (ChatGPT など) などのエージェントと連携します。特定のエージェントを通過させるには、 checkBotId() が isBot とともに返す検証済みボット フィールドを使用します。
Vercel は、検証済みボット ディレクトリからこれらのエージェントを識別し、 isVerifiedBot 、verifiedBotName 、および verifyBotCategory を返すため、他のすべてをブロックしながら、ChatGPT Operator のようなエージェントを許可できます。
app/api/chat/route.ts import { checkBotId } from 'botid/server' ; 'next/server' から { NextResponse } をインポートします。
エクスポート非同期関数 POST(リクエスト:リクエスト) {const{isBot, isVerifiedBot,verifiedBotName} = await checkBotId();
// ChatGPT オペレーターの通過を許可します。他のすべてのボットをブロックする const isOperator = isVerifiedBot &&verifiedBotName === 'chatgpt-operator' ;
if (isBot && ! isOperator ) { NextResponse を返します。 json ( { エラー : 'アクセスが拒否されました' } , { ステータス : 403 } ) ; }
// 検証された人間と許可されたエージェントに対して推論が実行されます const body = await request 。 json() ; const result = await runInference (body) ;
NextResponse を返します。 json ( { 結果 } ) ; }
async function runInference ( body :unknown ) { // ここで AI SDK またはモデルを呼び出します return {output : 'response' } ;検証済みのボット ディレクトリにない信頼されたサービスの場合は、ルートから保護を削除するのではなく、Vercel WAF にバイパス ルールを追加します。エージェントとカテゴリの完全なリストについては、「検証済みボットの処理」を参照してください。
checkBotId() は常にボットを報告します
一致するパスとメソッドを使用してルートがクライアント保護配列で宣言されていることを確認します。 BotID は宣言されたルートにのみチャレンジ ヘッダーを添付するため、宣言されていないルートにはサーバーが検証するものはありません。
カールでテストすると 403 が返される
BotID はブラウザ セッションで JavaScript を実行し、ヘッダーをサーバーに送信するため、C からの直接リクエストが

URL またはブラウザのアドレス バーは、運用環境ではボットとして扱われます。保護されたルートをテストするには、独自のアプリケーションのページからフェッチ リクエストを作成します。
地域の発展は常に過ぎ去る
checkBotId() でdevelopmentOptions オプションを設定しない限り、ローカル開発は isBot: false を返します。ボット トラフィックをシミュレートする手順については、BotID ドキュメントの「ローカル開発動作」を参照してください。
Nous Research が BotID を使用して自動化された不正行為を大規模にブロックした方法
BotID Deep Analysis は、高度なボット ネットワークをリアルタイムで捕捉します。
Vercel BotID は、推論が実行される前に、実際のブラウザからのリクエストを確認する非表示の CAPTCHA です。保護するルートにクライアント側のチャレンジを付加し、サーバー側の checkBotId() 呼び出しで各リクエストを分類し、自動化されたクライアントがモデルに到達する前に拒否します。このチェックはセッションごとに 1 回ではなくリクエストごとに実行されるため、攻撃者が一度チェックをバイパスしてそのアクセスを再利用することはできません。
BotID を AI エンドポイントに追加するにはどうすればよいですか?
ステップは 4 つあります。 botid パッケージをインストールし、 withBotId でフレームワーク構成をラップし、 initBotId() でクライアント上でルートを宣言し、モデルを実行する前にルート ハンドラーで checkBotId() を呼び出します。推論呼び出しの前にチェックを続けるということは、ブロックされたリクエストによってトークンが消費されることがないことを意味します。
BotID により Web サイトの速度が遅くなりますか?
いいえ。検出はクライアント セッション内で非同期に実行されるため、ページの読み込みがブロックされたり、実際のユーザーに顕著な遅延が追加されることはありません。ブラウザーシグナルを収集するスクリプトは軽量であり、サーバー上では checkBotId() はリクエストに既に添付されている判定のみを読み取るため、ハンドラーは別の分析ステップを待機しません。チェックは推論の前に実行されるため、高価なモデル呼び出しがトリガーされる前にボットリクエストを停止することで全体のコストを削減できます。

BotID を使用するには有料プランが必要ですか?
いいえ。基本的な検証はすべてのプランで無料で実行され、多くのあまり洗練されていないボットを捕捉します。 Kasada を利用した機械学習モデルを使用して数千のクライアント側シグナルを読み取る Deep Analysis は、Pro プランと Enterprise プランで利用できます。 checkBotId() を呼び出すリクエストに対してのみ課金され、パッシブなページビューに対しては課金されません。
BotID を介して ChatGPT のような信頼できるボットを許可するにはどうすればよいですか?
checkBotId() が isBot と一緒に返す検証済みボット フィールドを使用します。 isVerifiedBot と VerifiedBotName をチェックして、ChatGPT Operator などの既知のエージェントを許可し、他のすべてをブロックします。 Vercel の検証済みボット ディレクトリにない信頼できるサービスの場合は、ルートの保護を削除するのではなく、Vercel WAF にバイパス ルールを追加します。
Curl でテストすると BotID が 403 を返すのはなぜですか?
これは予想通りです。 BotID はブラウザ セッションで JavaScript を実行して、チャレンジ ヘッダーをサーバーに送信します。そのため、curl またはブラウザのアドレス バーからの直接リクエストにはヘッダーがなく、運用環境ではボットとして扱われます。保護されたルートをテストするには、独自のアプリケーション内のページからフェッチ リクエストを作成します。
サポートされました。送信 関連ドキュメントを読む
Vercel のボット管理機能を活用する方法
AI アプリをボットから保護する方法
ライト ダーク 製品 AI クラウド
AI ゲートウェイ 1 つのエンドポイント、すべてのモデル
サンドボックス分離された安全なコード実行
Vercel Agent スタックを理解しているエージェント
AI SDK TypeScript 用 AI ツールキット
コア プラットフォーム CI/CD チームの出荷速度を 6 倍に支援
コンテンツ配信 高速、スケーラブル、信頼性の高いコンテンツ配信
サーバーレス形式の Fluid Compute Server
ワークフロー 大規模な長時間実行ワークフロー
可観測性 あらゆるステップを追跡
セキュリティボット管理 スケーラブルなボット保護
プラットフォームのセキュリティ DDoS 保護、ファイアウォール
Web アプリケーション ファイアウォール きめ細かなカスタム保護
レソ

urces 会社の顧客 最高のチームから信頼される
ブログ 最新の投稿と変更点
Learn Docs Vercel ドキュメント
レベルアップのためのアカデミー リニア コース
ナレッジベース ヘルプをすぐに見つける
コミュニティ 会話に参加する
オープンソース
Next.js ネイティブ Next.js プラットフォーム
Nuxt 進歩的な Web フレームワーク
Svelte Web の効率的な UI フレームワーク
エンタープライズ規模の Turborepo スピード
ソリューションの使用例
AI アプリ AI のスピードで導入
コンポーザブルコマース コンバージョンを実現する強力なストアフロント
マーケティング サイト キャンペーンを迅速に開始
マルチテナント プラットフォーム 1 つのコードベースでアプリを拡張
Web アプリ インフラストラクチャではなく機能を提供する
ツール マーケットプレイス ワークフローを拡張および自動化する
テンプレート アプリ開発のジャンプスタート
パートナー ファインダー ソリューション パートナーからサポートを受ける
ユーザー プラットフォーム エンジニア 繰り返しを自動化する
あらゆるアイデアに対応する設計エンジニアの配置

## Original Extract

Gate every request to your AI endpoints with Vercel BotID and checkBotId() so inference runs only for verified callers. Covers setup, Deep Analysis, and verified bots.

How to protect your AI endpoints with Vercel BotID | Vercel Knowledge Base Skip to content How to protect your AI endpoints with Vercel BotID
Gate every request to your AI endpoints with Vercel BotID and checkBotId() so inference runs only for verified callers. Covers setup, Deep Analysis, and verified bots.
Knowledge Base / AI Copy URL Copy page Ask AI about this page 5 min read Last updated May 31, 2026 Vercel BotID lets you verify that each request to your AI endpoints comes from a real browser before any inference runs. Working as an invisible CAPTCHA, it attaches a client-side challenge to requests on the routes you protect, and a server-side checkBotId() call classifies each one, so automated clients are turned away before they reach your model. Running it on every request, rather than once per session, means an attacker can't bypass it once and reuse that access across thousands of stolen calls. That difference is what kept Nous Research’s chat app online when a coordinated attack spiked its traffic by 3,000% while inference stayed flat.
This guide walks you through installing BotID, declaring an AI route on the client, and gating that route with checkBotId() on the server so inference runs only for verified requests. You'll also set detection levels per route, enabling Deep Analysis on your highest-value endpoints and basic checks elsewhere, and learn how to let legitimate automation through with a Vercel WAF bypass rule.
A JavaScript project deployed on Vercel
An AI endpoint that accepts frontend requests, such as a route built with AI SDK
A Pro or Enterprise plan to use Deep Analysis (Basic is available on all plans)
Terminal npm i botid 2. Configure proxy rewrites
Wrap your Next.js config with withBotId . This sets up proxy rewrites so that ad-blockers and third-party scripts can't weaken BotID's protection:
next.config.ts import { withBotId } from 'botid/next/config' ;
const nextConfig = { // Your existing Next.js config } ;
export default withBotId ( nextConfig ) ; For Nuxt, SvelteKit, and other frameworks, the setup follows a similar pattern. See the BotID getting started guide for the per-framework versions.
3. Declare your AI route on the client
Call initBotId() during client initialization and list the AI routes you want to protect. BotID uses this list to attach challenge headers to matching requests. If a route isn't declared here, its requests arrive without those headers, so checkBotId() has nothing to verify and treats them as bots.
For Next.js 15.3 and later, use instrumentation-client.ts :
instrumentation-client.ts import { initBotId } from 'botid/client/core' ;
initBotId ( { protect : [ { path : '/api/chat' , method : 'POST' , } , ] , } ) ; On earlier versions of Next.js, mount the <BotIdClient /> component in your root layout head instead, passing the same protect array.
4. Verify every request on the server
Call checkBotId() inside the route handler, before the AI call runs. This is the load-bearing step: it returns a classification for the request currently being served, so a blocked request never reaches your model.
app/api/chat/route.ts import { checkBotId } from 'botid/server' ; import { NextRequest , NextResponse } from 'next/server' ;
export async function POST ( request : NextRequest ) { const verification = await checkBotId ( ) ;
if ( verification . isBot ) { return NextResponse . json ( { error : 'Access denied' } , { status : 403 } ) ; }
// Inference runs only after verification passes const body = await request . json ( ) ; const result = await runInference ( body ) ;
return NextResponse . json ( { result } ) ; }
async function runInference ( body : unknown ) { // Your AI SDK or model call here return { output : 'response' } ; } Placing the check before runInference means you incur the inference cost only for verified requests.
Basic validation catches many less sophisticated bots and runs free on all plans. For high-value AI routes, enable Deep Analysis, which uses a Kasada-powered machine learning model to analyze thousands of client-side signals.
Because Deep Analysis learns and adapts in real time, it can detect coordinated attacks that initially appear as legitimate traffic. In one incident, it traced a 500% traffic spike to a new bot network by correlating identical browser fingerprints cycling across proxy nodes. It then reclassified and blocked those sessions within roughly 10 minutes, without any manual intervention. For the full breakdown, see how BotID Deep Analysis caught a sophisticated bot network in real time .
Visit the Bot Management page in your project settings, then click the Configure button to open the configuration settings and enable Deep Analysis.
This feature is available for all customers on Pro and Enterprise plans . Only requests that invoke checkBotId() are charged, passive page views are not.
Run the check before inference : Keep checkBotId() ahead of the model call in your handler, so a blocked request never costs you a token.
Set detection levels per route : Use advancedOptions.checkLevel to apply deepAnalysis to your most sensitive routes and basic elsewhere. The checkLevel must be identical in your client and server configurations for each route, or verification will fail. This is available in botid@1.4.5 and later.
Allow trusted agents with Verified Bots
Blocking based on isBot alone also blocks legitimate automated agents, such as crawlers (e.g., Googlebot) and AI assistants (e.g., ChatGPT). To let specific agents through, use the verified-bot fields that checkBotId() returns along with isBot .
Vercel identifies these agents from its verified bot directory and returns isVerifiedBot , verifiedBotName , and verifiedBotCategory , so you can allow an agent like ChatGPT Operator while still blocking everything else.
app/api/chat/route.ts import { checkBotId } from 'botid/server' ; import { NextResponse } from 'next/server' ;
export async function POST ( request : Request ) { const { isBot , isVerifiedBot , verifiedBotName } = await checkBotId ( ) ;
// Allow ChatGPT Operator through; block all other bots const isOperator = isVerifiedBot && verifiedBotName === 'chatgpt-operator' ;
if ( isBot && ! isOperator ) { return NextResponse . json ( { error : 'Access denied' } , { status : 403 } ) ; }
// Inference runs for verified humans and allowed agents const body = await request . json ( ) ; const result = await runInference ( body ) ;
return NextResponse . json ( { result } ) ; }
async function runInference ( body : unknown ) { // Your AI SDK or model call here return { output : 'response' } ; } For a trusted service that isn't in the verified bot directory, add a bypass rule in the Vercel WAF rather than removing protection from the route. See Handling Verified Bots for the full list of agents and categories.
checkBotId() always reports a bot
Confirm the route is declared in your client protect array with a matching path and method . BotID only attaches challenge headers to declared routes, so an undeclared route has nothing for the server to verify.
Testing with curl returns a 403
BotID runs JavaScript in the browser session and sends headers to the server, so a direct request from curl or a browser address bar is treated as a bot in production. To test a protected route, make a fetch request from a page in your own application.
Local development always passes
Local development returns isBot: false unless you set the developmentOptions option on checkBotId() . See Local Development Behavior in the BotID docs for instructions on simulating bot traffic.
How Nous Research used BotID to block automated abuse at scale
BotID Deep Analysis catches a sophisticated bot network in real-time
Vercel BotID is an invisible CAPTCHA that confirms a request comes from a real browser before any inference runs. It attaches a client-side challenge to the routes you protect, then a server-side checkBotId() call classifies each request and turns away automated clients before they reach your model. Because the check runs on every request rather than once per session, an attacker can't bypass it once and reuse that access.
How do I add BotID to an AI endpoint?
There are four steps. Install the botid package, wrap your framework config with withBotId , declare the route on the client with initBotId() , then call checkBotId() in your route handler before the model runs. Keeping the check ahead of the inference call means a blocked request never costs you a token.
Will BotID slow down my website?
No. The detection runs asynchronously inside the client session, so it doesn't block page loads or add noticeable latency for real users. The script that gathers browser signals is lightweight, and on the server checkBotId() only reads the verdict that's already attached to the request, so your handler isn't waiting on a separate analysis step. Since the check runs before inference, it can lower your overall costs by stopping bot requests before they trigger an expensive model call.
Do I need a paid plan to use BotID?
No. Basic validation runs free on all plans and catches many less sophisticated bots. Deep Analysis, which uses a Kasada-powered machine learning model to read thousands of client-side signals, is available on Pro and Enterprise plans. You're only charged for requests that invoke checkBotId() , not for passive page views.
How do I let trusted bots like ChatGPT through BotID?
Use the verified-bot fields that checkBotId() returns alongside isBot . Check isVerifiedBot and verifiedBotName to allow a known agent, such as ChatGPT Operator, while still blocking everything else. For a trusted service that isn't in Vercel's verified bot directory, add a bypass rule in the Vercel WAF rather than removing protection for the route.
Why does BotID return a 403 when I test with curl?
This is expected. BotID runs JavaScript in the browser session to send challenge headers to the server, so a direct request from curl or a browser address bar has no headers and gets treated as a bot in production. To test a protected route, make a fetch request from a page inside your own application.
supported. Send Read related documentation
How to Utilize Vercel’s Bot Management Features
How to protect your AI app from bots
light dark Products AI Cloud
AI Gateway One endpoint, all your models
Sandbox Isolated, safe code execution
Vercel Agent An agent that knows your stack
AI SDK The AI Toolkit for TypeScript
Core Platform CI/CD Helping teams ship 6× faster
Content Delivery Fast, scalable, and reliable
Fluid Compute Servers, in serverless form
Workflow Long-running workflows at scale
Observability Trace every step
Security Bot Management Scalable bot protection
Platform Security DDoS Protection, Firewall
Web Application Firewall Granular, custom protection
Resources Company Customers Trusted by the best teams
Blog The latest posts and changes
Learn Docs Vercel documentation
Academy Linear courses to level up
Knowledge Base Find help quickly
Community Join the conversation
Open Source
Next.js The native Next.js platform
Nuxt The progressive web framework
Svelte The web’s efficient UI framework
Turborepo Speed with Enterprise scale
Solutions Use Cases
AI Apps Deploy at the speed of AI
Composable Commerce Power storefronts that convert
Marketing Sites Launch campaigns fast
Multi-tenant Platforms Scale apps with one codebase
Web Apps Ship features, not infrastructure
Tools Marketplace Extend and automate workflows
Templates Jumpstart app development
Partner Finder Get help from solution partners
Users Platform Engineers Automate away repetition
Design Engineers Deploy for every idea
