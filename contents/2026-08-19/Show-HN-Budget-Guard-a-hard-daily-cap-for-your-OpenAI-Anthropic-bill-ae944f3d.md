---
source: "https://github.com/kimbeomgyu/budget-guard"
hn_url: "https://news.ycombinator.com/item?id=49361479"
title: "Show HN: Budget Guard – a hard daily cap for your OpenAI/Anthropic bill"
article_title: "GitHub - kimbeomgyu/budget-guard: A circuit breaker for your LLM API bill — hard budget caps + per-feature cost attribution for OpenAI & Anthropic. · GitHub"
image: "https://opengraph.githubassets.com/8691ad3276e71abcf4bfd10a1069b9489d7c256a1b4260d198171e0b028425a8/kimbeomgyu/budget-guard"
author: "kimbeomgyu"
captured_at: "2026-08-19T14:24:29Z"
capture_tool: "hn-digest"
hn_id: 49361479
score: 1
comments: 1
posted_at: "2026-08-19T13:41:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Budget Guard – a hard daily cap for your OpenAI/Anthropic bill

- HN: [49361479](https://news.ycombinator.com/item?id=49361479)
- Source: [github.com](https://github.com/kimbeomgyu/budget-guard)
- Score: 1
- Comments: 1
- Posted: 2026-08-19T13:41:15Z

## Translation

タイトル: Show HN: Budget Guard – OpenAI/Anthropic 請求の 1 日あたりの厳しい上限
記事のタイトル: GitHub - kimbeomgyu/budget-guard: LLM API 請求のサーキット ブレーカー — OpenAI と Anthropic のハード予算上限 + 機能ごとのコスト帰属。 · GitHub
説明: LLM API 請求のサーキット ブレーカー — OpenAI と Anthropic のハードバジェット キャップ + 機能ごとのコスト帰属。 - キムボムギュ/バジェットガード

記事本文:
GitHub - kimbeomgyu/budget-guard: LLM API 請求のサーキット ブレーカー — OpenAI と Anthropic のハード予算上限 + 機能ごとのコスト帰属。 · GitHub
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
キムボムギュ
/
予算を守る
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
106 コミット 106 コミット フォルダーとファイル
.github .github 例 例 src src テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_

CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM API 請求のサーキット ブレーカー。 1 回のラップで、1 日の厳しい上限を設定します。暴走した再試行ループは、料金が請求される前にブロックされます。さらに、機能ごとのコストの帰属も表示されるため、実際に何にどれくらいのコストがかかるのかがわかります。
「5 ドルのタスクで 40 ドルの請求」を経験したインディーズ開発者向けに構築されています。 OpenAI / Anthropic / Gemini SDK と直接連携するか、Vercel AI SDK (v5 + v7)、LangChain.js、LlamaIndex.TS、Mastra アダプタを通じて動作します。ドロップイン: 通話は引き続きプロバイダーに直接送信されます。予算はカウントされ、上限が設定されるだけです。
無料およびオープンソース (MIT)。プロジェクト間の支出とアラートを備えたホスト型ダッシュボードが計画されていますが、SDK は無料であり、今後も無料です。
npm i 予算ガード
使ってみよう（OpenAI）
'openai' から OpenAI をインポートします。
'budget-guard' から {ガード} をインポートします。
const openai = new OpenAI() ;
const ai = Guard (openai . chat . completed , { project : 'my-app' , dailyCapUSD : 50 } ) ;
// 以前とまったく同じように使用します。オプションの機能タグを追加するだけです
const res = ai を待ちます。作成(
{ モデル : 'gpt-5.6' 、メッセージ : [ { 役割 : 'user' 、コンテンツ : 'hi' } ] } 、
{ 機能 : 'チャット' } ,
) ;
my-app の今日の支出がすでに $50 を超えている場合、次の呼び出しでは請求前に BudgetExceededError がスローされます。午前 3 時に突然請求書が届くことはもうありません。
'@anthropic-ai/sdk' から Anthropic をインポートします。
'budget-guard' から {ガード} をインポートします。
const anthropic = new Anthropic () ;
const ai = Guard ( anthropic .messages , { project : 'my-app' , dailyCapUSD : 50 } ) ;
待っててね。作成(
{ モデル : 'claude-opus-5' 、 max_tokens : 1024 、メッセージ : [ { 役割

: 'ユーザー' 、コンテンツ: 'こんにちは' } ] } 、
{ 機能 : '要約' } ,
) ;
Budget-Guard は、OpenAI (Azure、Mistral、DeepSeek、xAI を含む)、Anthropic、Google Gemini ( useMetadata )、AWS Bedrock Converse、Cohere ( billed_units ) の使用状況を自動検出します。キャッシュされたトークンと推論トークンは、プロバイダーの癖を含むクラスごとの実際の料金で請求されます (xAI と Gemini は出力カウントの範囲外で推論を報告します。予算ガードは暗黙的に過小カウントされないようにそれを追加します)。それ以外の場合は、独自のエクストラクターを渡します。
Guard ( client , opts , { useOf : ( res ) => ( { input : res . in , Output : res . out } ) } ;
何に何がかかるかを知る
import { payReport } から 'budget-guard' ;
putReport ( 'my-app' ) ; を待ちます。 // 非同期
// → { チャット: 2.41、要約: 0.88 } (今日、米ドル)
インスタンス間での共有上限 (Redis)
デフォルトでは、レジャーはメモリ内に (プロセスごとに) 存在します。これは、単一のスクリプト、ワーカー、またはエージェントに最適です。複数のインスタンスを実行していますか?共有ストアを渡すと、一緒に 1 つの上限が適用され、再起動後も存続できるようになります。
'redis' から { createClient } をインポートします。
import { ガード , redisStore } から 'budget-guard' ;
const redis = createClient() ;
redis を待ちます。接続する （ ） ;
const ai = ガード (openai . チャット . 完了 , {
プロジェクト: 'my-app' 、
デイリーキャップUSD : 50 、
ストア : redisStore ( redis ) , // ノード-redis v4;キーは自動的に期限切れになります (~2 日)
} ) ;
store は、小さな SpendStore インターフェイスを実装するもの ( add / get /entrys 、さらにアトミック予約用のオプションの addIfUnder — redisStore はサーバー側の Lua スクリプトとして実装します) を受け入れるため、既に実行しているものを使用してそれをバックアップできます。
スクリプトの実行全体で持続する (ファイル)
Cron ジョブと有効期間の短いスクリプトでは、メモリ内上限が静かに失敗します。毎回の実行は $0 から始まります。 fileStore は台帳を単一の JSON ファイルに保持するため、100 回実行されます

1日1キャップを共有します。
'budget-guard' から {ガード} をインポートします。
'budget-guard/file' から { fileStore } をインポートします。
const ai = ガード (openai . チャット . 完了 , {
プロジェクト : '夜の仕事' 、
dailyCapUSD : 5 、
ストア: fileStore ( '/var/tmp/my-app-spend.json' ) 、
} ) ;
書き込みはアトミック (一時ファイル + 名前変更) で、親ディレクトリが作成され、予算をサイレントにリセットする代わりに破損したファイルがスローされます。ストレージ層: メモリ = 1 つのプロセス、ファイル = 1 つのマシン、redis = フリート。 (同時プロセスは redisStore を使用する必要があります。ファイル ストアは順次実行をターゲットとしています。)
呼び出し前にブロック (オーバーシュートなし)
デフォルトでは、上限を超えると次のコールに上限が適用されるため、1 つのコールがオーバーシュートする可能性があります。推定子を与えると、問題のある呼び出し自体がブロックされます。組み込みのものはワンライナーです。
import { ガード , 推定子 } から 'budget-guard' ;
const ai = ガード (openai . チャット . 完了 , {
プロジェクト: 'my-app' 、
デイリーキャップUSD : 50 、
estimateUsage : estimator () , // chars/4 ヒューリスティック — サーキット ブレーカーとしては問題ありません
} ) ;
estimator() は、入力推定のプロンプト / システム / メッセージと、出力の宣言された max_tokens (または maxOutputTokens ) を読み取ります。新しい Claude トークナイザーの生成では、トークンの数が最大 30% 増加することが認識され (Opus 4.7 以降、Sonnet 5 以降、Fable、Mythos は自動的に修正されます)、 tools を渡すときにツール スキーマのオーバーヘッドが追加されます。正確な数を知りたいですか?トークナイザーを挿入します。
'gpt-tokenizer' から { countTokens } をインポートします。
estimateUsage: 推定子 ( { countTokens } ) ;
同時実行安全: ストアがサポートしている場合 (組み込みメモリ、ファイル、Redis ストアはすべてサポートしています)、推定コストは呼び出し前にアトミックに予約され、呼び出し後に実際のコストに精算されます。そのため、100 個の並列ワーカーが一緒に上限を超えて競争することはできません。失敗した通話は予約をロールバックします。
LlamaInde をラップする

x LLM — 上限は各通話前および非ストリーミング前に適用されます。
chat() は応答から測定されます。
'budget-guard' から {guardLlamaIndex } をインポートします。
const llm = GuardLlamaIndex (openai ( {モデル : 'gpt-5.6' } ) , { プロジェクト : 'my-app' , dailyCapUSD : 50 } ) ;
設定 。 llm = llm ; // または llm.chat(...) を直接呼び出します
使用法は、response.raw から読み取られます (プロバイダー間で機能します)。ストリーミングチャット()は
計測も可能 — 各チャンクの未処理の OpenAI / Anthropic / Gemini の使用状況が監視されます
ストリームが終了すると支出が精算されます。プロバイダーが指定しない場合
ストリーム内で使用すると、サイレント ゼロの代わりに警告が表示されます。新規ゼロ
依存関係。
コールバック ハンドラーを任意の LangChain モデルまたはチェーンにアタッチします。上限が適用されます。
各呼び出しの前に、コストは応答から測定されます。
'budget-guard/langchain' から { BudgetGuardHandler } をインポートします。
const handler = new BudgetGuardHandler ( { プロジェクト : 'my-app' , dailyCapUSD : 50 , モデル : 'gpt-5.6' } ) ;
モデルを待っています。 invoke(input, {callbacks:[handler]}) ; // オーバーキャップ → 呼び出し前にスロー
use_metadata から使用状況を読み取ります ( llmOutput.tokenUsage にフォールバックします)。パス
信頼性の高い価格設定のモデル (また、応答から自動検出されます)
現在）。 @langchain/core (オプションのピア依存関係) が必要です。
GuardOpenAI / GuardAnthropic / GuardGemini は、以下を設定する薄いラッパーです。
プロバイダーに代わってストリーミングが行われるため、ストリーミングは、
オプション:
'budget-guard' から {guardAnthropic } をインポートします。
const ai = GuardAnthropic ( anthropic .messages , { project : 'my-app' , dailyCapUSD : 50 } ) ;
// ストリーミングはすでにそれが人類的であることを知っています — 忘れるべき「プロバイダー」はありません
ストリーミング
ストリーミング呼び出しも従量制で行われます。通常どおり stream: true を渡すだけです。予算を守る
すべてのチャンクを直接渡して使用状況を読み取ります

決勝から
チャンクなので、ストリームが終了した後にコストが一度着地します。
const ai = Guard (openai . chat . completed , { project : 'my-app' , dailyCapUSD : 50 } ) ;
const stream = ai を待ちます。 create ( { モデル : 'gpt-5.6' , ストリーム : true } , { 機能 : 'チャット' } ) ;
for await (ストリームのconstチャンク) {
プロセスを標準出力 。 write (チャンク . 選択肢 [ 0 ] ?. デルタ ?. コンテンツ ?? '' ) ;
}
// ループが終了するとコストが記録されます
OpenAI の場合、stream_options: { include_usage: true } を挿入します (OpenAI
そのフラグが設定されている場合、最後のチャンクでのみ使用量が送信されます)。キャップはそのままです
呼び出しの前に強制されます。
Anthropic ストリーミングの場合、プロバイダーを「anthropic」に設定します。これは、
message_start / message_delta イベントを実行し、OpenAI のみのインジェクションをスキップします。
const ai = Guard (人間の . メッセージ , {
プロジェクト: 'my-app' 、
デイリーキャップUSD : 50 、
プロバイダー: '人間' 、
} ) ;
const stream = ai を待ちます。 create ( { モデル : 'claude-sonnet-4-6' 、ストリーム : true 、max_tokens : 1024 } ) ;
for await (ストリームの定数イベント) { /* ... */ }
Gemini ストリーミングの場合、プロバイダーを「gemini」に設定します。使用量は各チャンクから取得されます。
useMetadata (最後のものには合計が含まれます)。
AI SDK を使用しますか?任意のモデルをミドルウェアでラップします —
クライアントを保護する必要がないため、上限と機能ごとの測定が自動的に適用されます。作品
AI SDK v5 と v7 の両方で (使用状況は呼び出しごとに自動検出されます - 1 つ)
エントリ ポイント、何も設定する必要はありません):
import {wrapLanguageModel,generateText} from 'ai' ;
'@ai-sdk/openai' から {openai } をインポートします。
'budget-guard' から { BudgetGuardMiddleware } をインポートします。
const モデル = WrapLanguageModel ( {
モデル: openai ('gpt-5.6') 、
ミドルウェア : BudgetGuardMiddleware ( { プロジェクト : 'my-app' 、 dailyCapUSD : 50 、機能 : 'チャット' } ) 、
} ) ;
await generatedText ( { モデル , プロンプト : 'hi' } ) ; // ov

er cap → モデル呼び出しの前にスロー
generateText と streamText の両方を測定します (ストリームの終了から読み取られる使用量)
部分）。モデル呼び出しの前に、streamText ブロックをオーバーキャップします。エラーは
SDK の標準エラー チャネル ( onError 、または await result.text ) — textStream
デフォルトでストリームエラーを飲み込みます。
Mastra エージェントは Vercel AI SDK モデルで実行されるため、上記のミドルウェアはすでにカバーしています
それら — エージェントに渡す前にモデルをラップします。Mastra 固有のコードは使用しません。
'ai' から {wrapLanguageModel} をインポートします。
'@mastra/core/agent' から { エージェント } をインポートします。
'budget-guard' から { BudgetGuardMiddleware } をインポートします。
const モデル = WrapLanguageModel ( {
モデル: openai ('gpt-5.6') 、
ミドルウェア : BudgetGuardMiddleware ( { プロジェクト : 'my-app' , dailyCapUSD : 50 } ) 、
} ) ;
const エージェント = 新しいエージェント ({ id : 'サポート' , モデル, 指示 : '…' } ) ;
// (または @mastra/ai-sdk を直接使用する場合は withMastra(model, { … })
キャップとメータリングは、エージェントが行うすべてのモデル コールに適用されます。
すべての通話のコストを確認する (可観測性)
硬いキャップで出血を止めます。 onSpend で視聴できます。毎回発火します
SpendEvent を使用して呼び出しが成功すると、呼び出しごとのコストを直接パイプすることができます。
ログ、トレース、またはダッシュボード:
const ai = ガード (openai . チャット . 完了 , {
プロジェクト:

[切り捨てられた]

## Original Extract

A circuit breaker for your LLM API bill — hard budget caps + per-feature cost attribution for OpenAI & Anthropic. - kimbeomgyu/budget-guard

GitHub - kimbeomgyu/budget-guard: A circuit breaker for your LLM API bill — hard budget caps + per-feature cost attribution for OpenAI & Anthropic. · GitHub
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
kimbeomgyu
/
budget-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
106 Commits 106 Commits Folders and files
.github .github examples examples src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A circuit breaker for your LLM API bill. One wrap, set a hard daily cap — runaway retry loops get blocked before they bill you. Plus per-feature cost attribution so you know what actually costs what.
Built for indie devs who've seen "a $40 bill from a $5 task." Works with the OpenAI / Anthropic / Gemini SDKs directly, or through the Vercel AI SDK (v5 + v7), LangChain.js, LlamaIndex.TS and Mastra adapters. Drop-in: your calls still go straight to the provider — budget-guard just counts and caps.
Free & open source (MIT). A hosted dashboard with cross-project spend + alerts is planned — but the SDK is, and stays, free.
npm i budget-guard
Use it (OpenAI)
import OpenAI from 'openai' ;
import { guard } from 'budget-guard' ;
const openai = new OpenAI ( ) ;
const ai = guard ( openai . chat . completions , { project : 'my-app' , dailyCapUSD : 50 } ) ;
// use it exactly like before — just add an optional feature tag
const res = await ai . create (
{ model : 'gpt-5.6' , messages : [ { role : 'user' , content : 'hi' } ] } ,
{ feature : 'chat' } ,
) ;
If today's spend for my-app is already past $50 , the next call throws BudgetExceededError before it bills . No more 3am surprise invoices.
import Anthropic from '@anthropic-ai/sdk' ;
import { guard } from 'budget-guard' ;
const anthropic = new Anthropic ( ) ;
const ai = guard ( anthropic . messages , { project : 'my-app' , dailyCapUSD : 50 } ) ;
await ai . create (
{ model : 'claude-opus-5' , max_tokens : 1024 , messages : [ { role : 'user' , content : 'hi' } ] } ,
{ feature : 'summarize' } ,
) ;
budget-guard auto-detects the usage shapes of OpenAI (incl. Azure, Mistral, DeepSeek, xAI), Anthropic, Google Gemini ( usageMetadata ), AWS Bedrock Converse and Cohere ( billed_units ). Cached and reasoning tokens are billed at their real per-class rates — including the provider quirks (xAI and Gemini report reasoning outside the output count; budget-guard adds it back so you're not silently under-counting). For anything else, pass your own extractor:
guard ( client , opts , { usageOf : ( res ) => ( { input : res . in , output : res . out } ) } ) ;
Know what costs what
import { spendReport } from 'budget-guard' ;
await spendReport ( 'my-app' ) ; // async
// → { chat: 2.41, summarize: 0.88 } (today, in USD)
Shared caps across instances (Redis)
By default the ledger lives in memory (per process) — great for a single script, worker, or agent. Running multiple instances? Pass a shared store so they enforce one cap together and survive restarts:
import { createClient } from 'redis' ;
import { guard , redisStore } from 'budget-guard' ;
const redis = createClient ( ) ;
await redis . connect ( ) ;
const ai = guard ( openai . chat . completions , {
project : 'my-app' ,
dailyCapUSD : 50 ,
store : redisStore ( redis ) , // node-redis v4; keys auto-expire (~2 days)
} ) ;
store accepts anything implementing the tiny SpendStore interface ( add / get / entries , plus optional addIfUnder for atomic reservations — redisStore implements that as a server-side Lua script), so you can back it with whatever you already run.
Persist across script runs (file)
Cron jobs and short-lived scripts are where in-memory caps quietly fail — every run starts from $0. fileStore keeps the ledger in a single JSON file, so 100 runs a day share one cap:
import { guard } from 'budget-guard' ;
import { fileStore } from 'budget-guard/file' ;
const ai = guard ( openai . chat . completions , {
project : 'nightly-job' ,
dailyCapUSD : 5 ,
store : fileStore ( '/var/tmp/my-app-spend.json' ) ,
} ) ;
Writes are atomic (temp file + rename), parent directories are created, and a corrupted file throws instead of silently resetting your budget. Storage tiers: memory = one process, file = one machine, redis = a fleet. (Concurrent processes should use redisStore — the file store targets sequential runs.)
Block before the call (no overshoot)
By default the cap is enforced on the next call after you cross it, so one call can overshoot. Give it an estimator and it blocks the offending call itself — the built-in one is a one-liner:
import { guard , estimator } from 'budget-guard' ;
const ai = guard ( openai . chat . completions , {
project : 'my-app' ,
dailyCapUSD : 50 ,
estimateUsage : estimator ( ) , // chars/4 heuristic — fine for a circuit breaker
} ) ;
estimator() reads prompt / system / messages for the input estimate and the declared max_tokens (or maxOutputTokens ) for the output. It knows the new Claude tokenizer generation counts ~30% more tokens (Opus 4.7+, Sonnet 5+, Fable, Mythos — corrected automatically), and it adds tool-schema overhead when you pass tools . Want exact counts? Inject any tokenizer:
import { countTokens } from 'gpt-tokenizer' ;
estimateUsage: estimator ( { countTokens } ) ;
Concurrency-safe: when the store supports it (built-in memory, file and redis stores all do), the estimated cost is reserved atomically before the call and settled to the actual cost after — so 100 parallel workers can't race past the cap together. Failed calls roll their reservation back.
Wrap any LlamaIndex LLM — the cap applies before each call and non-streaming
chat() is metered from the response:
import { guardLlamaIndex } from 'budget-guard' ;
const llm = guardLlamaIndex ( openai ( { model : 'gpt-5.6' } ) , { project : 'my-app' , dailyCapUSD : 50 } ) ;
Settings . llm = llm ; // or call llm.chat(...) directly
Usage is read from response.raw (works across providers). Streaming chat() is
metered too — each chunk's raw is watched for OpenAI / Anthropic / Gemini usage
shapes and the spend is settled when the stream ends. If the provider doesn't put
usage in the stream, you get a warning instead of a silent zero. Zero new
dependencies.
Attach the callback handler to any LangChain model or chain — the cap is enforced
before each call, and cost is metered from the response:
import { BudgetGuardHandler } from 'budget-guard/langchain' ;
const handler = new BudgetGuardHandler ( { project : 'my-app' , dailyCapUSD : 50 , model : 'gpt-5.6' } ) ;
await model . invoke ( input , { callbacks : [ handler ] } ) ; // over cap → throws before the call
Reads usage from usage_metadata (falling back to llmOutput.tokenUsage ). Pass
model for reliable pricing (it's also auto-detected from the response when
present). Needs @langchain/core (an optional peer dependency).
guardOpenAI / guardAnthropic / guardGemini are thin wrappers that set
provider for you — so streaming is metered correctly without remembering the
option:
import { guardAnthropic } from 'budget-guard' ;
const ai = guardAnthropic ( anthropic . messages , { project : 'my-app' , dailyCapUSD : 50 } ) ;
// streaming already knows it's Anthropic — no `provider` to forget
Streaming
Streaming calls are metered too — just pass stream: true as usual. budget-guard
passes every chunk straight through to you and reads the usage from the final
chunk, so the cost lands once, after the stream finishes:
const ai = guard ( openai . chat . completions , { project : 'my-app' , dailyCapUSD : 50 } ) ;
const stream = await ai . create ( { model : 'gpt-5.6' , stream : true } , { feature : 'chat' } ) ;
for await ( const chunk of stream ) {
process . stdout . write ( chunk . choices [ 0 ] ?. delta ?. content ?? '' ) ;
}
// cost is recorded once the loop finishes
For OpenAI it injects stream_options: { include_usage: true } for you (OpenAI
only sends usage on the final chunk when that flag is set). The cap is still
enforced before the call.
For Anthropic streaming, set provider: 'anthropic' — it reads usage from the
message_start / message_delta events and skips the OpenAI-only injection:
const ai = guard ( anthropic . messages , {
project : 'my-app' ,
dailyCapUSD : 50 ,
provider : 'anthropic' ,
} ) ;
const stream = await ai . create ( { model : 'claude-sonnet-4-6' , stream : true , max_tokens : 1024 } ) ;
for await ( const event of stream ) { /* ... */ }
For Gemini streaming, set provider: 'gemini' — usage comes from each chunk's
usageMetadata (the last one carries the totals).
Using the AI SDK ? Wrap any model with the middleware —
no client to guard, the cap and per-feature metering apply automatically. Works
with both AI SDK v5 and v7 (the usage shape is auto-detected per call — one
entry point, nothing to configure):
import { wrapLanguageModel , generateText } from 'ai' ;
import { openai } from '@ai-sdk/openai' ;
import { budgetGuardMiddleware } from 'budget-guard' ;
const model = wrapLanguageModel ( {
model : openai ( 'gpt-5.6' ) ,
middleware : budgetGuardMiddleware ( { project : 'my-app' , dailyCapUSD : 50 , feature : 'chat' } ) ,
} ) ;
await generateText ( { model , prompt : 'hi' } ) ; // over cap → throws before the model call
Meters both generateText and streamText (usage read from the stream's finish
part). Over-cap streamText blocks before the model call; the error arrives on the
SDK's standard error channel ( onError , or await result.text ) — textStream
swallows stream errors by default.
Mastra agents run on Vercel AI SDK models, so the middleware above already covers
them — wrap the model before handing it to your agent, no Mastra-specific code:
import { wrapLanguageModel } from 'ai' ;
import { Agent } from '@mastra/core/agent' ;
import { budgetGuardMiddleware } from 'budget-guard' ;
const model = wrapLanguageModel ( {
model : openai ( 'gpt-5.6' ) ,
middleware : budgetGuardMiddleware ( { project : 'my-app' , dailyCapUSD : 50 } ) ,
} ) ;
const agent = new Agent ( { id : 'support' , model , instructions : '…' } ) ;
// (or withMastra(model, { … }) if you use @mastra/ai-sdk directly)
The cap and metering apply to every model call the agent makes.
See every call's cost (observability)
A hard cap stops the bleeding; onSpend lets you watch it. It fires on every
successful call with a SpendEvent , so you can pipe per-call cost straight into
your logs, traces, or a dashboard:
const ai = guard ( openai . chat . completions , {
project :

[truncated]
