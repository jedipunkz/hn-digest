---
source: "https://github.com/AllThingsSmitty/ai-flight-recorder"
hn_url: "https://news.ycombinator.com/item?id=49284938"
title: "Show HN: AI Flight Recorder – record, replay, and track AI session costs"
article_title: "GitHub - AllThingsSmitty/ai-flight-recorder: DevTools for AI apps. Record, replay, and inspect every prompt, token, tool call, and cost in a visual timeline. · GitHub"
author: "AllThingsSmitty"
captured_at: "2026-08-13T12:45:48Z"
capture_tool: "hn-digest"
hn_id: 49284938
score: 2
comments: 0
posted_at: "2026-08-13T12:30:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Flight Recorder – record, replay, and track AI session costs

- HN: [49284938](https://news.ycombinator.com/item?id=49284938)
- Source: [github.com](https://github.com/AllThingsSmitty/ai-flight-recorder)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T12:30:28Z

## Translation

タイトル: HN を表示: AI フライト レコーダー – AI セッション コストの記録、再生、追跡
記事のタイトル: GitHub - AllThingsSmitty/ai-flight-recorder: AI アプリ用の DevTools。すべてのプロンプト、トークン、ツール呼び出し、コストを視覚的なタイムラインで記録、再生、検査します。 · GitHub
説明: AI アプリ用の DevTools。すべてのプロンプト、トークン、ツール呼び出し、コストを視覚的なタイムラインで記録、再生、検査します。 - AllThingsSmitty/ai-フライトレコーダー

記事本文:
GitHub - AllThingsSmitty/ai-flight-recorder: AI アプリ用の DevTools。すべてのプロンプト、トークン、ツール呼び出し、コストを視覚的なタイムラインで記録、再生、検査します。 · GitHub
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
すべてのものSmitty
/
AIフライトレコーダー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
129 コミット 129 コミット .changeset .changeset .github .github .vscode .vscode アプリ アプリの例 例 パッケージ パッケージ

es サンプル サンプル スクリプト スクリプト .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc CODE-OF-CONDUCT.md CODE-OF-CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI Flight Recorder は、AI アプリケーション内のすべてのインタラクション (プロンプト、ストリーミング トークン、ツール呼び出し、レイテンシー、コスト) をすべて 1 か所で記録、再生、検査するためのオープンソースの開発者ツールです。
事後にコンソール ログをつなぎ合わせる代わりに、1 行の SDK ラッパーをドロップして、一時停止したり巻き戻したり、.flight ファイルとしてチームメイトに渡すことができる完全な DevTools スタイルのタイムラインを取得します。
セッション記録: すべてのプロンプト、トークン、ツール呼び出し、完了を構造化されたイベント ストリームとしてキャプチャします。
ストリーミング リプレイ: 速度制御 (0.25 倍 ～ 8 倍) を使用して、リアルタイムでセッションの再生を視聴します。
タイムラインとウォーターフォール: 並列ツール呼び出しやストリーミング遅延を含むリクエストのライフサイクル全体を視覚化します。
コスト分析: トークンの使用量とセッションあたりの推定費用を分析します。
検索とフィルター: タイムライン全体にわたってイベントをタイプまたはキーワードでフィルターします。
プロバイダー アダプター: OpenAI、Anthropic、Google Gemini 用の 1 行ラッパー (ストリーミングおよび非ストリーミング)
.flight エクスポート/インポート: セッションをポータブル ファイルとして共有し、別の開発者がローカルで再生できるようにします。
プラグイン システム: カスタム オブザーバーを使用してレコーダーのライフサイクルに接続します。
トランスポート システム: 任意のストレージ バックエンド (メモリ内、ファイル システム、独自の API) を接続します。
OpenTelemetry エクスポート: あらゆるセッションを OTLP トレース ペイロードに変換し、Jaeger、Grafana Tempo、Honeycomb、または任意の OTel-comp に取り込みます。

atible バックエンド ( @ai-flight-recorder/sdk からの toOtlp )
AIフライトレコーダー/
§── アプリ/
│ §── devtools/ Next.js DevTools アプリケーション
│ §── docs/Starlight ドキュメント サイト
│ └── vscode/ VS Code 拡張子 — .flight ファイル用のカスタム エディター
§── パッケージ/
│ §── コア/ドメイン モデル — イベント、セッション、レコーダー、リプレイ エンジン
│ §── sdk/ 開発者向け API — FlightRecorder、アダプター、プラグイン、トランスポート
│ §── ui/ 共有 React コンポーネント (将来)
│ └── タイプ/共有 TypeScript タイプ (将来)
§── スクリプト/
│ ━──smoke.ts SDK統合スモークテスト
└── 例/
§── nextjs-chat/ フルスタック チャット アプリ — OpenAI ストリーミング + .flight エクスポート
§── node-anthropic/ Node.js の例 — Anthropic + FileTransport
━──node-gemini/ Node.js の例 — Google Gemini + FileTransport
はじめに
pnpmインストール
DevTools アプリを実行する
pnpm開発
http://localhost:3000 を開きます。アプリは 2 つのデモ セッションで読み込まれるため、API キーは必要なく、すぐに UI を探索できます。
pnpm煙
記録、プラグイン、トランスポート、シリアル化、再生をエンドツーエンドで実行します。 40 個のアサーションすべてが合格する必要があります。
import { FlightRecorder } から "@ai-flight-recorder/sdk" ;
const fr = 新しいフライトレコーダー () ;
const セッション = fr 。 startSession ( { ラベル : "my-chat" } ) ;
フランス。レコード ( {
タイプ: "プロンプト" 、
モデル：「gpt-4o」、
プロンプト:「フランスの首都はどこですか?」 、
} ) ;
フランス。レコード ( {
タイプ: "完了" 、
応答：「パリ」 、
終了理由 : "停止" 、
合計トークン数 : 18 、
} ) ;
const 終了 = fr 。 endSession() ;
プロバイダーアダプター
プロバイダー クライアントをインターセプトし、すべての呼び出しを自動的に記録するドロップイン ラッパー。
「openai」から OpenAI をインポートします。
import { FlightRecorder , WrapOpenAI } from "@ai-flight-recorder/sdk" ;

const fr = 新しいフライトレコーダー () ;
const openai =wrapOpenAI(new OpenAI(), fr.recordor);
フランス。 startSession ( { ラベル : "チャット" } ) ;
const 応答 = openai を待ちます。チャット 。完成品。作成 ( {
モデル：「gpt-4o」、
メッセージ : [ { 役割 : "ユーザー" 、コンテンツ : "こんにちは" } ] 、
} ) ;
フランス。 endSession() ;
人間的
"@anthropic-ai/sdk" から Anthropic をインポートします。
import { FlightRecorder , WrapAnthropic } from "@ai-flight-recorder/sdk" ;
const fr = 新しいフライトレコーダー () ;
const client =wrapAnthropic(new Anthropic(), fr.recordor);
フランス。 startSession ( { ラベル : "クロードチャット" } ) ;
const message = クライアントを待ちます。メッセージ。作成 ( {
モデル：「クロード・ソネット-4-5」、
max_tokens : 1024 、
メッセージ : [ { 役割 : "ユーザー" 、コンテンツ : "こんにちは" } ] 、
} ) ;
フランス。 endSession() ;
Google ジェミニ
"@google/generative-ai" から { GoogleGenerativeAI } をインポートします。
import { FlightRecorder , WrapGeminiModel } from "@ai-flight-recorder/sdk" ;
const fr = 新しいフライトレコーダー () ;
const genAI = new GoogleGenerativeAI (プロセス . 環境 . GOOGLE_API_KEY ! ) ;
const モデル = WrapGeminiModel (
ゲンアイ 。 getGenerativeModel ( { モデル : "gemini-1.5-pro" } ) 、
フランス。レコーダー、
) ;
フランス。 startSession ( { ラベル : "ジェミニチャット" } ) ;
const result = モデルを待ちます。 generateContent ( "Hello" ) ;
フランス。 endSession() ;
3 つのアダプターはすべてストリーミングをサポートしています。既存のクライアントをラップすると、すべての通話が自動的に記録されます。
import { FlightRecorder , ConsoleLogPlugin } from "@ai-flight-recorder/sdk" ;
const fr = 新しいフライトレコーダー ( {
プラグイン: [
新しい ConsoleLogPlugin ( { logEvents : true , logsummary : true } ) 、
// インラインプラグイン
{
名前: "my-plugin" 、
onSessionStart: (セッション) => コンソール。 log ( "開始:" 、セッション . id ) 、
onEvent: (event) => myMetrics 。記録（イベント）、
onSessionEnd:(セッション) => アラート。 f

ラッシュ (セッション) 、
} 、
]、
} ) ;
use() はチェーン可能であり、登録時に重複する名前をチェックします。
フランス。 (プラグインA)を使用します。使用 (プラグインB) ;
輸送
import { FlightRecorder , InMemoryTransport } from "@ai-flight-recorder/sdk" ;
const Transport = new InMemoryTransport();
const fr = new FlightRecorder ({transport }) ;
フランス。 startSession() ;
// ... イベントを記録します ...
フランス。 endSession() ; // 自動的にトランスポートに保存されます
const セッション = トランスポート 。 getAll() ;
Node.js ファイルシステムのトランスポート:
import { FlightRecorder } から "@ai-flight-recorder/sdk" ;
import { FileTransport } から "@ai-flight-recorder/sdk/node" ;
const Transport = new FileTransport ( "./recordings" ) ;
const fr = new FlightRecorder ({transport }) ;
フランス。 startSession ( { label : "my-session" } ) ;
// ... イベントを記録します ...
フランス。 endSession() ;
// ./recordings/<sessionId>.flight に保存します
const セッション = トランスポート 。ロードオール() ;
独自のトランスポートを実装する
import type { Transport } from "@ai-flight-recorder/sdk" ;
class MyApiTransport は Transport { を実装します。
非同期保存 (セッション) {
await fetch ( "/api/sessions" , {
メソッド: "POST" 、
本文: JSON 。 stringify (セッション) 、
} ) ;
}
}
const fr = new FlightRecorder({transport:new MyApiTransport()});
.flight ファイル形式
セッションは、ポータブルな .flight ファイル (バージョン エンベロープを含む JSON) としてエクスポートできます。
{
"バージョン" : " 1 " 、
"exportedAt" : 1721484000000 、
「セッション」: {
"id" : " ... " 、
"label" : " バグレポート-123 " ,
"ステータス" : "終了" 、
"startedAt" : 1721484000000 、
"終了後" : 1721484060000 、
「イベント」: [ ... ]
}
}
DevTools UI からエクスポートする: セッションがアクティブなときにツールバーの [エクスポート] ボタンをクリックします。
DevTools UI にインポートします。[インポート] をクリックして .flight ファイルを選択します。セッションがセッション リストに追加され、アクトになります

すぐにセッションを開始します。
import {serializeSession , deserializeSession } from "@ai-flight-recorder/sdk" ;
import { writeFileSync , readFileSync } from "node:fs" ;
// エクスポート
writeFileSync (「bug-123.flight」、serializeSession (終了セッション) ) ;
// インポート
const session = deserializeSession ( readFileSync ( "bug-123.flight" , "utf-8" ) ) ;
開発ツールアプリケーション
DevTools アプリ ( apps/devtools ) は、記録されたセッションに視覚的なインターフェイスを提供する Next.js アプリケーションです。
タイムライン: タイプ バッジ、説明、タイミング オフセットを含む時系列のイベント リスト
ウォーターフォール: ストリーミング スパンとツール呼び出し時間を示す視覚的な遅延の内訳
コスト分析: トークンの使用量の内訳とリクエストあたりの推定費用
「セッションを再生」をクリックして再生モードに入ります
速度制御: 0.25×、0.5×、1×、2×、4×、8×
セッション内の任意のポイントにジャンプするためのシークバー
トークンが再生されると、トークン ストリームがリアルタイムで組み立てられます
チップ行を使用してイベント タイプでフィルタリングします (プロンプト、トークン、ツール、結果、完了、エラー)。
イベントコンテンツ全体のテキスト検索
Examples/nextjs-chat は、完全なエンドツーエンドの統合 (GPT-4o-mini によるストリーミング チャット、自動セッション記録、.flight エクスポート) を示す最小限の Next.js アプリです。
cd サンプル/nextjs-chat
cp .env.example .env.local
.env.local を編集し、OpenAI API キーを追加します。
pnpm開発
http://localhost:3000 を開きます。アシスタントとチャットし、ヘッダーの [Export .flight] をクリックしてセッションをダウンロードします。
DevTools アプリ (リポジトリ ルートからの pnpm dev) を開き、ツールバーの [インポート] をクリックして、.flight ファイルを選択します。セッションは、タイムライン、ウォーターフォール、コストの内訳、完全なストリーミング再生など、即座に読み込まれます。
この例では、SDK から 3 つのものを接続します。
FlightRecorder : リクエストごとにセッションを開始します
WrapOpenAI : OpenAI クライアントをインターセプトし、すべてのプロンプトを記録します。

ken、および自動的に完了
SerializeSession : 終了したセッションをダウンロード用に JSON にシリアル化します。
代わりに Anthropic または Gemini を使用するには、 src/app/api/chat/route.ts で WrapOpenAI を WrapAnthropic または WrapGeminiModel に置き換えます。
# すべてのパッケージをビルドする
pnpmビルド
# 開発モードで DevTools を実行する
pnpm開発
# すべてのパッケージをタイプチェックします
pnpm タイプチェック
# すべてのパッケージをリントします
pnpm リント
# SDK スモークテスト (ビルドは必要ありません)
pnpm煙
新しいイベントタイプの追加
型リテラルをpackages/core/src/events/EventType.tsに追加します。
Packages/core/src/events/YourEvent.ts に BaseEvent を拡張するインターフェイスを作成します。
これをpackages/core/src/events/AIEvent.tsのAIEvent共用体に追加します。
package/core/src/events/index.ts からエクスポートします。
メタデータを表示するために、DevTools アプリのeventMeta.tsにケースを追加します。
@ai-flight-recorder/core からプラグイン インターフェイスを実装します。
import type { Plugin , AIEvent , Session } from "@ai-flight-recorder/sdk" ;
エクスポート クラス MyPlugin はプラグインを実装します {
読み取り専用名 = "my-plugin" ;
onSessionStart (セッション : セッション) { ... }
onEvent (イベント : AIEvent ) { ... }
onSessionEnd (セッション : セッション) { ... }
}
ライセンス
このプロジェクトは MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
AI アプリ用の DevTools。すべてのプロンプト、トークン、ツール呼び出し、コストを視覚的なタイムラインで記録、再生、検査します。
ai-flight-recorder.vercel.app トピックス
お読みください

[切り捨てられた]

## Original Extract

DevTools for AI apps. Record, replay, and inspect every prompt, token, tool call, and cost in a visual timeline. - AllThingsSmitty/ai-flight-recorder

GitHub - AllThingsSmitty/ai-flight-recorder: DevTools for AI apps. Record, replay, and inspect every prompt, token, tool call, and cost in a visual timeline. · GitHub
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
AllThingsSmitty
/
ai-flight-recorder
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
129 Commits 129 Commits .changeset .changeset .github .github .vscode .vscode apps apps examples examples packages packages samples samples scripts scripts .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc CODE-OF-CONDUCT.md CODE-OF-CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json turbo.json turbo.json View all files Repository files navigation
AI Flight Recorder is an open-source developer tool for recording, replaying, and inspecting every interaction in an AI application — prompts, streamed tokens, tool calls, latency, and cost — all in one place.
Instead of piecing together console logs after the fact, you drop in a one-line SDK wrapper and get a full DevTools-style timeline you can pause, rewind, and hand off to a teammate as a .flight file.
Session recording: capture every prompt, token, tool call, and completion as a structured event stream
Streaming replay: watch a session play back in real time with speed controls (0.25×–8×)
Timeline & Waterfall: visualize the full request lifecycle including parallel tool calls and streaming latency
Cost Analysis: break down token usage and estimated spend per session
Search & Filter: filter events by type or keyword across the full timeline
Provider Adapters: one-line wrappers for OpenAI, Anthropic, and Google Gemini (streaming and non-streaming)
.flight Export/Import: share a session as a portable file another developer can replay locally
Plugin System: hook into the recorder lifecycle with custom observers
Transport System: plug in any storage backend (in-memory, filesystem, your own API)
OpenTelemetry Export: convert any session to an OTLP trace payload for ingestion into Jaeger, Grafana Tempo, Honeycomb, or any OTel-compatible backend ( toOtlp from @ai-flight-recorder/sdk )
ai-flight-recorder/
├── apps/
│ ├── devtools/ Next.js DevTools application
│ ├── docs/ Starlight documentation site
│ └── vscode/ VS Code extension — custom editor for .flight files
├── packages/
│ ├── core/ Domain model — events, session, recorder, replay engine
│ ├── sdk/ Developer-facing API — FlightRecorder, adapters, plugins, transports
│ ├── ui/ Shared React components (future)
│ └── types/ Shared TypeScript types (future)
├── scripts/
│ └── smoke.ts SDK integration smoke test
└── examples/
├── nextjs-chat/ Full-stack chat app — OpenAI streaming + .flight export
├── node-anthropic/ Node.js example — Anthropic + FileTransport
└── node-gemini/ Node.js example — Google Gemini + FileTransport
Getting Started
pnpm install
Run the DevTools app
pnpm dev
Open http://localhost:3000 . The app loads with two demo sessions so you can explore the UI immediately — no API keys required.
pnpm smoke
Exercises recording, plugins, transport, serialization, and replay end-to-end. All 40 assertions should pass.
import { FlightRecorder } from "@ai-flight-recorder/sdk" ;
const fr = new FlightRecorder ( ) ;
const session = fr . startSession ( { label : "my-chat" } ) ;
fr . record ( {
type : "prompt" ,
model : "gpt-4o" ,
prompt : "What is the capital of France?" ,
} ) ;
fr . record ( {
type : "completion" ,
response : "Paris." ,
finishReason : "stop" ,
totalTokens : 18 ,
} ) ;
const ended = fr . endSession ( ) ;
Provider adapters
Drop-in wrappers that intercept the provider client and record every call automatically.
import OpenAI from "openai" ;
import { FlightRecorder , wrapOpenAI } from "@ai-flight-recorder/sdk" ;
const fr = new FlightRecorder ( ) ;
const openai = wrapOpenAI ( new OpenAI ( ) , fr . recorder ) ;
fr . startSession ( { label : "chat" } ) ;
const response = await openai . chat . completions . create ( {
model : "gpt-4o" ,
messages : [ { role : "user" , content : "Hello" } ] ,
} ) ;
fr . endSession ( ) ;
Anthropic
import Anthropic from "@anthropic-ai/sdk" ;
import { FlightRecorder , wrapAnthropic } from "@ai-flight-recorder/sdk" ;
const fr = new FlightRecorder ( ) ;
const client = wrapAnthropic ( new Anthropic ( ) , fr . recorder ) ;
fr . startSession ( { label : "claude-chat" } ) ;
const message = await client . messages . create ( {
model : "claude-sonnet-4-5" ,
max_tokens : 1024 ,
messages : [ { role : "user" , content : "Hello" } ] ,
} ) ;
fr . endSession ( ) ;
Google Gemini
import { GoogleGenerativeAI } from "@google/generative-ai" ;
import { FlightRecorder , wrapGeminiModel } from "@ai-flight-recorder/sdk" ;
const fr = new FlightRecorder ( ) ;
const genAI = new GoogleGenerativeAI ( process . env . GOOGLE_API_KEY ! ) ;
const model = wrapGeminiModel (
genAI . getGenerativeModel ( { model : "gemini-1.5-pro" } ) ,
fr . recorder ,
) ;
fr . startSession ( { label : "gemini-chat" } ) ;
const result = await model . generateContent ( "Hello" ) ;
fr . endSession ( ) ;
All three adapters support streaming. Wrap your existing client and all calls are recorded automatically.
import { FlightRecorder , ConsoleLogPlugin } from "@ai-flight-recorder/sdk" ;
const fr = new FlightRecorder ( {
plugins : [
new ConsoleLogPlugin ( { logEvents : true , logSummary : true } ) ,
// Inline plugin
{
name : "my-plugin" ,
onSessionStart : ( session ) => console . log ( "Started:" , session . id ) ,
onEvent : ( event ) => myMetrics . record ( event ) ,
onSessionEnd : ( session ) => alerting . flush ( session ) ,
} ,
] ,
} ) ;
use() is chainable and checks for duplicate names at registration time:
fr . use ( pluginA ) . use ( pluginB ) ;
Transport
import { FlightRecorder , InMemoryTransport } from "@ai-flight-recorder/sdk" ;
const transport = new InMemoryTransport ( ) ;
const fr = new FlightRecorder ( { transport } ) ;
fr . startSession ( ) ;
// ... record events ...
fr . endSession ( ) ; // automatically saves to transport
const sessions = transport . getAll ( ) ;
Node.js filesystem transport:
import { FlightRecorder } from "@ai-flight-recorder/sdk" ;
import { FileTransport } from "@ai-flight-recorder/sdk/node" ;
const transport = new FileTransport ( "./recordings" ) ;
const fr = new FlightRecorder ( { transport } ) ;
fr . startSession ( { label : "my-session" } ) ;
// ... record events ...
fr . endSession ( ) ;
// saves to ./recordings/<sessionId>.flight
const sessions = transport . loadAll ( ) ;
Implement your own transport
import type { Transport } from "@ai-flight-recorder/sdk" ;
class MyApiTransport implements Transport {
async save ( session ) {
await fetch ( "/api/sessions" , {
method : "POST" ,
body : JSON . stringify ( session ) ,
} ) ;
}
}
const fr = new FlightRecorder ( { transport : new MyApiTransport ( ) } ) ;
.flight File Format
Sessions can be exported as portable .flight files (JSON with a version envelope):
{
"version" : " 1 " ,
"exportedAt" : 1721484000000 ,
"session" : {
"id" : " ... " ,
"label" : " bug-report-123 " ,
"status" : " ended " ,
"startedAt" : 1721484000000 ,
"endedAt" : 1721484060000 ,
"events" : [ ... ]
}
}
Export from the DevTools UI: click the Export button in the toolbar while a session is active.
Import into the DevTools UI: click Import and select a .flight file. The session is added to the session list and becomes the active session immediately.
import { serializeSession , deserializeSession } from "@ai-flight-recorder/sdk" ;
import { writeFileSync , readFileSync } from "node:fs" ;
// Export
writeFileSync ( "bug-123.flight" , serializeSession ( endedSession ) ) ;
// Import
const session = deserializeSession ( readFileSync ( "bug-123.flight" , "utf-8" ) ) ;
DevTools Application
The DevTools app ( apps/devtools ) is a Next.js application providing a visual interface for recorded sessions.
Timeline: chronological event list with type badges, descriptions, and timing offsets
Waterfall: visual latency breakdown showing streaming spans and tool call durations
Cost Analysis: token usage breakdown and estimated spend per request
Click "Replay Session" to enter replay mode
Speed controls: 0.25×, 0.5×, 1×, 2×, 4×, 8×
Seek bar for jumping to any point in the session
Token stream assembles in real time as tokens replay
Filter by event type using the chip row (Prompt, Token, Tool, Result, Completion, Error)
Text search across event content
examples/nextjs-chat is a minimal Next.js app showing a full end-to-end integration — streaming chat with GPT-4o-mini, automatic session recording, and .flight export.
cd examples/nextjs-chat
cp .env.example .env.local
Edit .env.local and add your OpenAI API key:
pnpm dev
Open http://localhost:3000 . Chat with the assistant, then click Export .flight in the header to download your session.
Open the DevTools app ( pnpm dev from the repo root), click Import in the toolbar, and select the .flight file. Your session loads instantly — timeline, waterfall, cost breakdown, and full streaming replay.
The example wires up three things from the SDK:
FlightRecorder : starts a session per request
wrapOpenAI : intercepts the OpenAI client and records every prompt, token, and completion automatically
serializeSession : serializes the ended session to JSON for download
To use Anthropic or Gemini instead, swap wrapOpenAI for wrapAnthropic or wrapGeminiModel in src/app/api/chat/route.ts .
# Build all packages
pnpm build
# Run DevTools in development mode
pnpm dev
# Typecheck all packages
pnpm typecheck
# Lint all packages
pnpm lint
# SDK smoke test (no build required)
pnpm smoke
Adding a new event type
Add the type literal to packages/core/src/events/EventType.ts
Create the interface in packages/core/src/events/YourEvent.ts extending BaseEvent
Add it to the AIEvent union in packages/core/src/events/AIEvent.ts
Export it from packages/core/src/events/index.ts
Add a case to eventMeta.ts in the DevTools app for display metadata
Implement the Plugin interface from @ai-flight-recorder/core :
import type { Plugin , AIEvent , Session } from "@ai-flight-recorder/sdk" ;
export class MyPlugin implements Plugin {
readonly name = "my-plugin" ;
onSessionStart ( session : Session ) { ... }
onEvent ( event : AIEvent ) { ... }
onSessionEnd ( session : Session ) { ... }
}
License
This project is licensed under the MIT License - see the LICENSE file for details.
DevTools for AI apps. Record, replay, and inspect every prompt, token, tool call, and cost in a visual timeline.
ai-flight-recorder.vercel.app Topics
Readme

[truncated]
