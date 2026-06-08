---
source: "https://github.com/salimassili62-afk/ai-costguard"
hn_url: "https://news.ycombinator.com/item?id=48446076"
title: "AI CostGuard – Local-first runtime safety layer for AI agents"
article_title: "GitHub - salimassili62-afk/ai-costguard: Local-first runtime safety layer for AI agents that blocks runaway costs, loops, retries, and budget overruns before provider API calls execute. · GitHub"
author: "salim2006"
captured_at: "2026-06-08T16:30:42Z"
capture_tool: "hn-digest"
hn_id: 48446076
score: 2
comments: 0
posted_at: "2026-06-08T14:43:23Z"
tags:
  - hacker-news
  - translated
---

# AI CostGuard – Local-first runtime safety layer for AI agents

- HN: [48446076](https://news.ycombinator.com/item?id=48446076)
- Source: [github.com](https://github.com/salimassili62-afk/ai-costguard)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T14:43:23Z

## Translation

タイトル: AI CostGuard – AI エージェント向けのローカルファーストのランタイム安全レイヤー
記事のタイトル: GitHub - salimassili62-afk/ai-costguard: プロバイダー API 呼び出しが実行される前に暴走コスト、ループ、再試行、予算超過をブロックする AI エージェント用のローカルファーストのランタイム安全レイヤー。 · GitHub
説明: プロバイダー API 呼び出しが実行される前に、暴走コスト、ループ、再試行、予算超過をブロックする、AI エージェント用のローカルファーストのランタイム安全レイヤー。 - salimassili62-afk/ai-costguard

記事本文:
GitHub - salimassili62-afk/ai-costguard: プロバイダー API 呼び出しが実行される前に暴走コスト、ループ、再試行、予算超過をブロックする AI エージェント用のローカルファーストのランタイム安全レイヤー。 · GitHub
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
salimassili62-afk
/
ai-コストガード
公共
通知

イケーション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
salimassili62-afk/ai-costguard
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 コミット 40 コミット .github/ workflows .github/ workflows ベンチマーク ベンチマーク カバレッジ カバレッジ ドキュメント ドキュメントの例 例 ランディング ランディング レガシー アーティファクト レガシー アーティファクト スクリプト スクリプト src src テンプレート テンプレート テスト テスト .gitignore .gitignore .prettierrc .prettierrc ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md aifw.example.js aifw.example.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべて表示ファイル リポジトリ ファイルのナビゲーション
AI CostGuard は、API 呼び出しが実行される前に暴走コスト、ループ、再試行、予算の爆発を防ぐ、AI エージェント向けのローカルファーストのランタイム安全レイヤーです。 OpenAI 互換クライアントと関数スタイルの SDK 呼び出しをラップし、リクエストのコストをローカルで見積もり、予算超過をブロックし、繰り返されるプロンプトを検出し、構造化されたイベントを発行し、CLI チェックとローカル ダッシュボードを公開します。
地元第一主義です。これには、SaaS コントロール プレーン、クラウド ダッシュボード、プロキシ ゲートウェイ、テレメトリ サービス、請求調整サービス、またはハード セキュリティ境界は含まれません。
npm インストール @salimassili/ai-costguard
クイックスタート
'openai' から OpenAI をインポートします。
import { ガード , GuardError } から '@salimassili/ai-costguard' ;
const openai =guard (new OpenAI({apiKey:process.env.OPENAI_API_KEY}),{
予算：5、
最大ステップ数 : 50 、
スコープ : { プロジェクト ID : 'my-app' } 、
} ) ;
{を試してください
const 応答 = openai を待ちます。チャット 。完成品。作成 ( {
モデル: 'gpt-4o-mini' 、
メッセージ : [ { 役割 : 'ユーザー' 、コンテンツ : '短い金額を書き込みます

メアリー。 } ]、
max_tokens : 200 、
} ) ;
コンソール。 log (応答 . 選択肢 [0] ?. メッセージ ?. 内容 ) ;
} キャッチ (エラー) {
if ( GuardError のエラーインスタンス) {
コンソール。 error (エラー.コード、エラー.メッセージ、エラー.コンテキスト);
} それ以外の場合は {
エラーをスローします。
}
}
守るもの
デフォルトでは、AI CostGuard は次の SDK メソッド パスを評価します。
他のクライアント メソッドはコスト チェックなしでパススルーされます。カスタム クライアント メソッドを保護するには:
const client = ガード (customClient , {
予算：2、
GuardedMethods : [ 'agent.run' ] 、
価格設定オーバーライド : [
{
モデル: '内部モデル' 、
inputPer1kTokens : 0.001 、
出力Per1kTokens : 0.002 、
最終更新日 : '2026-06-07' 、
出典 : 「社内価格設定シート」 、
} 、
]、
} ) ;
Vercel AI SDK アダプター、LangChain ラッパー、エージェント ランナーなどの関数スタイル SDK の場合:
'@salimassili/ai-costguard' から {guardFunction } をインポートします。
const GuardedGenerateText = GuardFunction (generateTextAdapter , {
予算：1、
スコープ : { プロジェクト ID : 'チャットボット' } 、
} ) ;
await GuardedGenerateText ( {
モデル: 'gpt-4o-mini' 、
プロンプト: 'ユーザーに 1 つの段落で答えてください。' 、
max_tokens : 200 、
} ) ;
決定と誤り
ブロックされたリクエストは、プロバイダー メソッドが呼び出される前に GuardError をスローします。
{を試してください
オープンアイを待ってください。チャット 。完成品。作成 (リクエスト) ;
} キャッチ (エラー) {
if ( GuardError のエラーインスタンス) {
コンソール。ログ (エラーコード) ;
コンソール。ログ (エラー.メタデータ) ;
}
}
現在のランタイム ブロック コード:
INVALID_LICENSE は、古い呼び出し元との互換性のためにエクスポートされたタイプに残りますが、現在の Pro ヘルパーはローカル ライセンス チェックを強制しません。
ガード (クライアント、{
予算：10、
maxSteps : 100 、
動作分析 : true 、
最大履歴 : 32 、
履歴TtlMs : 5 * 60 * 1000 、
ループ類似性しきい値 : 0.85 、
ループ最小繰り返し数 : 2 、
再試行しきい値 : 2 、
スコープ: {
プロジェクト ID : 'producti

オン API 、
userId : 'オプションのユーザー' 、
sessionId : 'オプションのエージェント実行' 、
} 、
GuardedMethods : [ 'chat.completions.create' , 'responses.create' ] ,
価格設定オーバーライド: [ ] 、
Webhook : {
スラック：プロセスの緩み環境 。 SLACK_WEBHOOK 、
discord : プロセス。環境 。 DISCORD_WEBHOOK 、
再試行: 2 、
タイムアウトMs : 1500 、
} 、
イベントログパス: '.ai-costguard/events.jsonl' 、
イベントログプロンプト: 'なし' 、
} ) ;
スコープは予算と行動履歴を分離します。スコープが指定されていない場合、ガードは 1 つのプロセスローカルのデフォルト スコープを使用します。
AI CostGuard は通話前の見積もりであり、請求元帳ではありません。
試行コスト: ブロックされた試行を含む、保護されたすべての試行の推定コスト。
totalCost : 許可された呼び出しの推定コスト。
BlockedCost : プロバイダーの実行前に停止した推定コスト。
actualCost : 応答に認識可能な使用量フィールドが含まれている場合に、プロバイダーが報告した使用量コスト。
予算の決定には、推定許容コストが使用されます。実際の使用状況は可観測性のために記録されますが、以前の決定が書き換えられることはありません。
既知のモデルの価格は、組み込みのレジストリ エントリ、ランタイム登録、またはガードごとのオーバーライドによって決まります。不明なモデルはデフォルトでブロックされます。
import { registerPricing } から '@salimassili/ai-costguard' ;
登録価格 ( [
{
モデル: '私の会社のモデル' 、
inputPer1kTokens : 0.001 、
出力Per1kTokens : 0.002 、
最終更新日 : '2026-06-07' 、
ソース : '内部' 、
} 、
]);
意図的に不明なモデルのフォールバック価格を設定したい場合:
ガード (クライアント、{
予算：5、
不明なモデルポリシー: 'フォールバック' 、
不明モデル価格 : {
モデル: 'フォールバック' 、
inputPer1kTokens : 0.001 、
出力Per1kTokens : 0.002 、
最終更新日 : '2026-06-07' 、
ソース: 'アプリケーション フォールバック' 、
} 、
} ) ;
価格は頻繁に変更されます。運用環境で使用する前にプロバイダーの価格を確認し、必要に応じてエントリを上書きします。
const unsubscribe = openai 。に ( 'ブロ

ck' , ( イベント ) => {
コンソール。 log (イベント . コード 、 イベント . 理由 、 イベント . コンテキスト . 推定コスト ) ;
} ) ;
購読解除 ( ) ;
サポートされているイベントは、cost 、allow 、および block です。ハンドラーエラーは無視されるため、可観測性コードはガードの決定を変更できません。
ローカル JSONL イベント ログにオプトインします。
const openai = ガード ( クライアント , {
予算：5、
イベントログパス: '.ai-costguard/events.jsonl' 、
} ) ;
ローカル専用ダッシュボードを開始します。
ai-costguard ダッシュボード --events .ai-costguard/events.jsonl --budget 5
1 回限りのパッケージ実行の場合:
npx @salimassili/ai-costguard ダッシュボード --events .ai-costguard/events.jsonl --budget 5
パッケージがローカルにインストールされている場合は、npx ai-costguard ダッシュボードも機能します。ダッシュボードはデフォルトで 127.0.0.1 にバインドされ、ローカル イベント ファイルのみを読み取ります。
ai-costguard ダッシュボード --events .ai-costguard/events.jsonl --budget 5 --once --json
docs/DASHBOARD.md を参照してください。
実行可能なモック化されたサンプルには、次のものが含まれています。
OpenAI SDK エージェントのループ保護
Anthropic SDK ワークフローの予算ガード
Vercel AI SDK チャットボットの予算上限
LangChain の再試行ストームの防止
Mastra スタイルのエージェント ランナー保護
docs/INTEGRATIONS.md および example/integrations を参照してください。
ミドルウェアは手動チェッカーを添付します。すべてのルートを自動的に解析または検査するわけではありません。
import { ミドルウェア , GuardError } から '@salimassili/ai-costguard' ;
アプリ 。 use(ミドルウェア({予算:2}));
アプリ 。 post ( '/chat' , async ( req , res , next ) => {
{を試してください
要求 .ローカルセーフティ 。チェック ( {
モデル: 'gpt-4o-mini' 、
トークン: 500、
inputTokens : 100 、
出力トークン: 400 、
推定コスト : 0.0003 、
タイムスタンプ: 日付。今 ( ) 、
プロンプト: 文字列 (req .body ?.プロンプト ?? '')、
} ) ;
レス。 json ({ok : true }) ;
} キャッチ (エラー) {
if ( GuardError のエラーインスタンス) {
レス。ステータス (403) 。 json ( { コード : エラー . コード 、理由 : エラー . messa

ゲ } ) ;
戻る ;
}
次へ (エラー) ;
}
} ) ;
オプションの Redis / Pro ヘルパー
Redis を利用した共有支出追跡は、サブパス インポートの背後に隔離されています。
'@salimassili/ai-costguard/pro' から { GuardPro } をインポートします。
const pro = 新しい GuardPro ( {
redisUrl : プロセス。環境 。 REDIS_URL ?? ” 、
予算：25、
ウィンドウ秒 : 86400 、
} ) ;
プロを待ってください。 checkAndCharge ( 'production' , 0.0042 ) ;
プロを待ってください。シャットダウン ( ) ;
ioredis はオプションの依存関係であり、ルート インポートによってロードされません。
LicenseKey は、非推奨の互換性フィールドとしてのみ受け入れられます。 AI CostGuard は、商用ライセンスをローカルで強制しません。また、validateLicense() は、セキュリティではなく形式の健全性ヘルパーです。
aifw check --budget 1 --model gpt-4o-mini --input-tokens 500 --tokens 1000 --max-steps 5
このパッケージは、ai-costguard bin エイリアスもインストールします。
ai-costguard check --budget 1 --model gpt-4o-mini --tokens 1000 --max-steps 5
ai-costguard ダッシュボード --events .ai-costguard/events.jsonl --budget 5
カスタムモデルの場合:
aifw check --budget 1 --model external-model --tokens 1000 --input-price-per-1k 0.001 --output-price-per-1k 0.002
終了コード:
0 : 予測コストは予算内です
1 : 予想コストが予算を超えています
npm ビルドを実行する
npm 実行ベンチマーク
このスクリプトは、実行時のオーバーヘッド、おおよそのヒープ デルタ、誤検知シナリオ、ループ検出動作、およびコスト見積もりの境界を報告します。結果は局所的な測定値であり、普遍的な保証ではありません。 docs/BENCHMARKS.md を参照してください。
Node v24.14.1 / Windows 上のこのリポジトリの最新のローカル ベンチマークでは、5000 回の反復を超えるモックされたガードされた呼び出しごとに 0.020691 ミリ秒の追加が測定されました。パフォーマンス重視の要求でこの数値を使用する前に、ターゲット ランタイムで再実行してください。
簡単な自家製予算チェックでは、1 つのカウンターが 1 つの番号を横切った後に 1 つのリクエストを停止できます。 AI CostGuard は、通常は乱雑になる部分をパッケージ化します。

エージェントが本番環境に入ると、次のようになります。
ランタイムオーバーライドと不明なモデルのブロックを備えたプロバイダー価格レジストリ。
API 応答の構造化された GuardError コードとメタデータ。
プロジェクト、ユーザー、またはセッションごとにスコープを設定した予算と動作の状態。
ループおよび再試行ストームの検出。
推定使用量、試行使用量、ブロック使用量、および実際の使用量のアカウンティング。
メソッド フィルタリングにより、非 AI SDK 呼び出しは課金されません。
イベント フック、ベストエフォート Webhook、JSONL イベント ログ、ローカル ダッシュボードの可視性。
CI 予算のチェックと実行可能な統合の例。
npmci
npm ビルドを実行する
npm タイプチェックを実行する
npmテスト
npm 煙を実行する
npm 実行ベンチマーク
npm Audit --omit=dev
npm パック --dry-run
制限事項
トークンのカウントは概算であり、依存関係はありません。
価格エントリは古くなる可能性があります。運用環境ではそれらをオーバーライドします。
フリー ガードはプロセスローカルです。
ループ検出では、埋め込みではなく、文字のトライグラムの類似性を使用します。
Webhook はベストエフォート型であり、強制には影響しません。
ダッシュボードはローカルの JSONL ログのみを読み取ります。これはホスト型分析製品ではありません。
プロバイダーの使用状況の調整は、応答が認識可能な使用法フィールドを公開している場合にのみ機能します。
プロバイダー API の前に暴走コスト、ループ、再試行、予算超過をブロックする AI エージェント向けのローカルファーストのランタイム安全レイヤー

[切り捨てられた]

## Original Extract

Local-first runtime safety layer for AI agents that blocks runaway costs, loops, retries, and budget overruns before provider API calls execute. - salimassili62-afk/ai-costguard

GitHub - salimassili62-afk/ai-costguard: Local-first runtime safety layer for AI agents that blocks runaway costs, loops, retries, and budget overruns before provider API calls execute. · GitHub
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
salimassili62-afk
/
ai-costguard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
salimassili62-afk/ai-costguard
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits .github/ workflows .github/ workflows benchmarks benchmarks coverage coverage docs docs examples examples landing landing legacy-artifacts legacy-artifacts scripts scripts src src templates templates test test .gitignore .gitignore .prettierrc .prettierrc ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md aifw.example.js aifw.example.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
AI CostGuard is a local-first runtime safety layer for AI agents that prevents runaway costs, loops, retries, and budget explosions before API calls execute. It wraps OpenAI-compatible clients and function-style SDK calls, estimates request cost locally, blocks budget overruns, detects repeated prompts, emits structured events, and exposes CLI checks plus a local dashboard.
It is local-first. It does not include a SaaS control plane, cloud dashboard, proxy gateway, telemetry service, billing reconciliation service, or hard security boundary.
npm install @salimassili/ai-costguard
Quick Start
import OpenAI from 'openai' ;
import { guard , GuardError } from '@salimassili/ai-costguard' ;
const openai = guard ( new OpenAI ( { apiKey : process . env . OPENAI_API_KEY } ) , {
budget : 5 ,
maxSteps : 50 ,
scope : { projectId : 'my-app' } ,
} ) ;
try {
const response = await openai . chat . completions . create ( {
model : 'gpt-4o-mini' ,
messages : [ { role : 'user' , content : 'Write a short summary.' } ] ,
max_tokens : 200 ,
} ) ;
console . log ( response . choices [ 0 ] ?. message ?. content ) ;
} catch ( error ) {
if ( error instanceof GuardError ) {
console . error ( error . code , error . message , error . context ) ;
} else {
throw error ;
}
}
What It Guards
By default AI CostGuard evaluates these SDK method paths:
Other client methods are passed through without cost checks. To protect a custom client method:
const client = guard ( customClient , {
budget : 2 ,
guardedMethods : [ 'agent.run' ] ,
pricingOverrides : [
{
model : 'internal-model' ,
inputPer1kTokens : 0.001 ,
outputPer1kTokens : 0.002 ,
lastUpdated : '2026-06-07' ,
source : 'internal pricing sheet' ,
} ,
] ,
} ) ;
For function-style SDKs such as Vercel AI SDK adapters, LangChain wrappers, or agent runners:
import { guardFunction } from '@salimassili/ai-costguard' ;
const guardedGenerateText = guardFunction ( generateTextAdapter , {
budget : 1 ,
scope : { projectId : 'chatbot' } ,
} ) ;
await guardedGenerateText ( {
model : 'gpt-4o-mini' ,
prompt : 'Answer the user in one paragraph.' ,
max_tokens : 200 ,
} ) ;
Decisions And Errors
Blocked requests throw GuardError before the provider method is called.
try {
await openai . chat . completions . create ( request ) ;
} catch ( error ) {
if ( error instanceof GuardError ) {
console . log ( error . code ) ;
console . log ( error . metadata ) ;
}
}
Current runtime block codes:
INVALID_LICENSE remains in the exported type for compatibility with older callers, but the current Pro helper does not enforce local license checks.
guard ( client , {
budget : 10 ,
maxSteps : 100 ,
behaviorAnalysis : true ,
maxHistory : 32 ,
historyTtlMs : 5 * 60 * 1000 ,
loopSimilarityThreshold : 0.85 ,
loopMinRepeats : 2 ,
retryThreshold : 2 ,
scope : {
projectId : 'production-api' ,
userId : 'optional-user' ,
sessionId : 'optional-agent-run' ,
} ,
guardedMethods : [ 'chat.completions.create' , 'responses.create' ] ,
pricingOverrides : [ ] ,
webhooks : {
slack : process . env . SLACK_WEBHOOK ,
discord : process . env . DISCORD_WEBHOOK ,
retries : 2 ,
timeoutMs : 1500 ,
} ,
eventLogPath : '.ai-costguard/events.jsonl' ,
eventLogPrompt : 'none' ,
} ) ;
scope isolates budgets and behavior history. If no scope is supplied, the guard uses one process-local default scope.
AI CostGuard is a pre-call estimator, not a billing ledger.
attemptedCost : estimated cost of every guarded attempt, including blocked attempts.
totalCost : estimated cost of allowed calls.
blockedCost : estimated cost stopped before provider execution.
actualCost : provider-reported usage cost when the response includes recognizable usage fields.
Budget decisions use estimated allowed cost. Actual usage is recorded for observability but does not rewrite earlier decisions.
Known model pricing comes from built-in registry entries, runtime registrations, or per-guard overrides. Unknown models are blocked by default.
import { registerPricing } from '@salimassili/ai-costguard' ;
registerPricing ( [
{
model : 'my-company-model' ,
inputPer1kTokens : 0.001 ,
outputPer1kTokens : 0.002 ,
lastUpdated : '2026-06-07' ,
source : 'internal' ,
} ,
] ) ;
If you intentionally want fallback pricing for unknown models:
guard ( client , {
budget : 5 ,
unknownModelPolicy : 'fallback' ,
unknownModelPricing : {
model : 'fallback' ,
inputPer1kTokens : 0.001 ,
outputPer1kTokens : 0.002 ,
lastUpdated : '2026-06-07' ,
source : 'application fallback' ,
} ,
} ) ;
Pricing changes frequently. Verify provider pricing before production use and override entries when needed.
const unsubscribe = openai . on ( 'block' , ( event ) => {
console . log ( event . code , event . reason , event . context . estimatedCost ) ;
} ) ;
unsubscribe ( ) ;
Supported events are cost , allow , and block . Handler errors are swallowed so observability code cannot change guard decisions.
Opt into a local JSONL event log:
const openai = guard ( client , {
budget : 5 ,
eventLogPath : '.ai-costguard/events.jsonl' ,
} ) ;
Start the local-only dashboard:
ai-costguard dashboard --events .ai-costguard/events.jsonl --budget 5
For one-off package execution:
npx @salimassili/ai-costguard dashboard --events .ai-costguard/events.jsonl --budget 5
If the package is installed locally, npx ai-costguard dashboard also works. The dashboard binds to 127.0.0.1 by default and reads only local event files.
ai-costguard dashboard --events .ai-costguard/events.jsonl --budget 5 --once --json
See docs/DASHBOARD.md .
Runnable mocked examples are included for:
OpenAI SDK agent loop protection
Anthropic SDK workflow budget guard
Vercel AI SDK chatbot budget cap
LangChain retry-storm prevention
Mastra-style agent runner protection
See docs/INTEGRATIONS.md and examples/integrations .
The middleware attaches a manual checker. It does not automatically parse or inspect every route.
import { middleware , GuardError } from '@salimassili/ai-costguard' ;
app . use ( middleware ( { budget : 2 } ) ) ;
app . post ( '/chat' , async ( req , res , next ) => {
try {
req . localSafety . check ( {
model : 'gpt-4o-mini' ,
tokens : 500 ,
inputTokens : 100 ,
outputTokens : 400 ,
estimatedCost : 0.0003 ,
timestamp : Date . now ( ) ,
prompt : String ( req . body ?. prompt ?? '' ) ,
} ) ;
res . json ( { ok : true } ) ;
} catch ( error ) {
if ( error instanceof GuardError ) {
res . status ( 403 ) . json ( { code : error . code , reason : error . message } ) ;
return ;
}
next ( error ) ;
}
} ) ;
Optional Redis / Pro Helper
Redis-backed shared spend tracking is isolated behind a subpath import:
import { GuardPro } from '@salimassili/ai-costguard/pro' ;
const pro = new GuardPro ( {
redisUrl : process . env . REDIS_URL ?? '' ,
budget : 25 ,
windowSeconds : 86400 ,
} ) ;
await pro . checkAndCharge ( 'production' , 0.0042 ) ;
await pro . shutdown ( ) ;
ioredis is an optional dependency and is not loaded by the root import.
licenseKey is accepted as a deprecated compatibility field only. AI CostGuard does not enforce commercial licenses locally, and validateLicense() is a format sanity helper, not security.
aifw check --budget 1 --model gpt-4o-mini --input-tokens 500 --tokens 1000 --max-steps 5
The package also installs an ai-costguard bin alias:
ai-costguard check --budget 1 --model gpt-4o-mini --tokens 1000 --max-steps 5
ai-costguard dashboard --events .ai-costguard/events.jsonl --budget 5
For custom models:
aifw check --budget 1 --model internal-model --tokens 1000 --input-price-per-1k 0.001 --output-price-per-1k 0.002
Exit codes:
0 : projected cost is within budget
1 : projected cost exceeds budget
npm run build
npm run benchmark
The script reports runtime overhead, approximate heap delta, false-positive scenarios, loop detection behavior, and cost-estimation boundaries. Results are local measurements, not universal guarantees. See docs/BENCHMARKS.md .
Latest local benchmark in this repo on Node v24.14.1 / Windows measured 0.020691 ms added per mocked guarded call over 5000 iterations. Re-run on your target runtime before using this number in performance-sensitive claims.
A simple homemade budget check can stop one request after one counter crosses one number. AI CostGuard packages the parts that usually become messy once agents enter production:
Provider pricing registry with runtime overrides and unknown-model blocking.
Structured GuardError codes and metadata for API responses.
Scoped budget and behavior state per project, user, or session.
Loop and retry-storm detection.
Estimated, attempted, blocked, and actual usage accounting.
Method filtering so non-AI SDK calls are not charged.
Event hooks, best-effort webhooks, JSONL event logs, and local dashboard visibility.
CI budget checks and runnable integration examples.
npm ci
npm run build
npm run typecheck
npm test
npm run smoke
npm run benchmark
npm audit --omit=dev
npm pack --dry-run
Limitations
Token counting is approximate and dependency-free.
Pricing entries can become stale; override them for production.
The free guard is process-local.
Loop detection uses character trigram similarity, not embeddings.
Webhooks are best-effort and never affect enforcement.
The dashboard reads local JSONL logs only; it is not a hosted analytics product.
Provider usage reconciliation only works when responses expose recognizable usage fields.
Local-first runtime safety layer for AI agents that blocks runaway costs, loops, retries, and budget overruns before provider API

[truncated]
