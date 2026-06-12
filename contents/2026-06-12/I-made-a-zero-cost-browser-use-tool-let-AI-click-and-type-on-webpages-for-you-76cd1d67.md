---
source: "https://github.com/pdufour/browser-use-wasm"
hn_url: "https://news.ycombinator.com/item?id=48502402"
title: "I made a zero cost browser-use tool – let AI click and type on webpages for you"
article_title: "GitHub - pdufour/browser-use-wasm · GitHub"
author: "pdufour"
captured_at: "2026-06-12T13:19:03Z"
capture_tool: "hn-digest"
hn_id: 48502402
score: 1
comments: 1
posted_at: "2026-06-12T10:50:25Z"
tags:
  - hacker-news
  - translated
---

# I made a zero cost browser-use tool – let AI click and type on webpages for you

- HN: [48502402](https://news.ycombinator.com/item?id=48502402)
- Source: [github.com](https://github.com/pdufour/browser-use-wasm)
- Score: 1
- Comments: 1
- Posted: 2026-06-12T10:50:25Z

## Translation

タイトル: ゼロコストのブラウザ利用ツールを作りました – AI に Web ページをクリックして入力させます
記事のタイトル: GitHub - pdufour/browser-use-wasm · GitHub
説明: GitHub でアカウントを作成して、pdufour/browser-use-wasm の開発に貢献します。

記事本文:
GitHub - pdufour/browser-use-wasm · GitHub
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
プデュフォー
/
ブラウザ使用-wasm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く Fo

フォルダとファイル
11 コミット 11 コミット .cursor .cursor .github/ workflows .github/ workflows 資産/ Visual-diff アセット/ Visual-diff ドキュメント ドキュメント サンプル サンプル スクリプト スクリプト src src テスト テスト .gitignore .gitignore .npmignore .npmignore GEMINI.md GEMINI.md README.md README.md デモ-vid.gif デモ-vid.gif デモ-vid.mp4 デモ-vid.mp4 パッケージ-ロック.json パッケージ-lock.json パッケージ.json パッケージ.json playwright.config.js playwright.config.js 試した方法.txt 試した方法.txt tsconfig.json tsconfig.json tsconfig.src.json tsconfig.src.json vite.browse-proxy.js vite.browse-proxy.js vite.coi-pages.js vite.coi-pages.js vite.config.js vite.config.js vite.config.src.js vite.config.src.js vite.coop-corp.js vite.coop-corp.js vite.demo-pages.js vite.demo-pages.js vite.fixtures.js vite.fixtures.js vite.model-cache.js vite.model-cache.js vite.wllama-plugins.js vite.wllama-plugins.js vite.wllama-wasm.js vite.wllama-wasm.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ライブデモ → https://pdufour.github.io/browser-use-wasm/
WASM でのブラウザーの使用 — バックエンドを使用しないビジョン言語の自動化。 browser-use-wasm は、クライアント側で完全なブラウザー使用ループを実行します: SnapDOM スクリーンショット → ShowUI グラウンディング → ライブページのクリック/入力/選択。デフォルトのモデルは ShowUI-2B です。組み込みのモデル セレクターから他のものを交換します。 Chrome または Edge のみ (WebGPU + WASM ワーカー)。
デモビデオ.mp4
ブラウザの使用方法
ここでのブラウザの使用とは、人間のようにページを見て (スクリーンショット)、何をするかを決定し (VLA)、接地点で実際の DOM に基づいて動作することを意味します。
┌─────────┐ ┌─────────┐ ┌─────────┐ ┌──────────┐
│ ライブページ │ ──► │ SnapDOM │ ──► │ ShowUI VLA │ ──► │ DOM アクション

│
│ (iframe) │ │ スクリーンショット │ │ WASM ワーカー │ │ [x, y] │
━━━━━━━━━━━━━━━━━━━━━━━━━━┘ ━━━━━━━━━━━━┘
キャプチャ — SnapDOM は、ターゲット要素 (または iframe コンテンツ) をキャンバスに複製します。これはブラウザで「見える」ものです。
サイズ変更 — 画像はモデルのビジョン トークン バジェットに合わせて縮小されます。
Infer — ユーザーの目標ごとに 1 つの ShowUI ナビゲーション呼び出し。モデルは、スクリーンショットのみで正規化された [x, y] を使用して構造化アクション (CLICK、INPUT、SELECT、ENTER など) を出力します。
実行 — ブラウザーの使用は、その座標の elementFromPoint を介してライブ ページ上で動作します。グラウンディングのためのラベル テキスト DOM ルックアップはありません。
1 つの目標 → 1 つの推論 → 解析されたすべてのアクションを実行します。バッチの途中で UI を変更すると (例: CLICK してから INPUT )、次のステップの前に自動再キャプチャがトリガーされます。
座標はライブ DOM レイアウト、セレクター、OCR から取得されることはなく、SnapDOM バッファー上のビジョン推論からのみ取得されます。
browser-use-wasm をローカルで実行します。ブラウザー使用スタックは完全にタブ内に存在します。
WebGPU + JSPI を備えた Chrome または Edge — ブラウザ使用の推論はサーバーではなくここで実行されます
開発サーバーは COOP/COEP ヘッダー ( vite.config.js に含まれる) を送信する必要があります。
パブリック レジストリ モデルは、最初の読み込み時にブラウザにダウンロードされます (ShowUI-2B Q4_K_M + mmproj の場合は約 1.8 GB)。内蔵セレクターからモデルを交換します
npm install github:pdufour/browser-use-wasm#main
# または: 糸追加 pdufour/browser-use-wasm#main
Examples/ は公開された tarball から除外されます ( .npmignore を参照)。
ライブラリはルート パッケージ (browser-use-wasm) です。例/アプリは、 "browser-use-wasm": "file:.." を介してローカルにリンクします。
npm install # root — テスト、キャッシュスクリプト、ライブラリ開発
cd の例 && npm install && cd .. # リンク

リポジトリルートからのブラウザ使用-wasm
npm run dev # ホームページは /、オペレーターは /home/
npm run build # dist/lib/ + dist/examples/
npm run typecheck # src + 例
ブラウザでの使用をアプリに埋め込む
browser-use-wasm は、ブラウザ使用の自動化のための小さな埋め込み API を公開します。高レベルのエントリ ポイント: createWebOperator() → モデルのロード → SnapDOM キャプチャ → instruct(goal) 。
import { createWebOperator } から 'browser-use-wasm' ;
const 演算子 = createWebOperator ( {
// スクリーンショットの内容 (デフォルト: document.body)
CaptureRoot : ( ) => ドキュメント 。 getElementById ( 'app' ) 、
// クリック/入力するドキュメント (デフォルト: ホストドキュメント。埋め込みには iframe contentDocument を使用)
targetDocument : ( ) => ドキュメント 。 getElementById ( 'frame' ) ?。コンテンツドキュメント ??書類、
} ) ;
演算子を待ちます。ロード ( {
onStatus: (msg) => コンソール。ログ (メッセージ) 、
} ) ;
const cap = 演算子を待ちます。捕獲 （ ） ;
キャップを待ちます。 whenEncoded ; // ワーカーで JPEG エンコードを待つ
const result = 演算子を待ちます。 instruct ( '送信をクリック' , {
onStatus: (msg) => コンソール。ログ (メッセージ) 、
onBeforeStep : async (ステップ) => {
// 各 CLICK / INPUT / …の前にカーソルをアニメーション化します
} 、
} ) ;
コンソール。ログ (結果の概要) ; // 例: 「クリック→入力→決定」
コンソール。ログ (結果.ステップ) ; // アクションごとの OK/詳細 + キャプチャーノルムポイント
createWebOperator(オプション?)
オプション
デフォルト
説明
モデルID
ShowUI-2B
レジストリモデルID
wasmURL
/wllama/wllama.wasm
同一起源のラマ WASM
キャプチャルート
文書.本文
要素の SnapDOM キャプチャ
ターゲットドキュメント
ホストドキュメント
ライブページ DOM ツールが動作する
推論タイムアウトMs
12000
推論ごとの上限
onPerfMark
—
キャプチャ パイプライン タイミング フック
WebOperator メソッド
方法
説明
ロード(オプション?)
GGUF をダウンロードしてブラウザ WASM ワーカーにロードします
キャプチャ()
ブラウザ使用のキャプチャ — SnapDOM → ビジョンのサイズ変更

→ 非同期JPEGエンコード
instruct(タスク、オプト?)
1 つの ShowUI 推論 + 解析されたブラウザ使用アクションをすべて実行
検索(ラベル)
現在のキャプチャを実行せずに <label> をグランドクリックします
hasCapture()
スクリーンショットバッファの準備ができているかどうか
クリアキャプチャ()
スクリーンショットをドロップします (次の命令には再度 Capture() が必要です)
setModel(id)
レジストリ モデルの切り替え (load() が再度必要)
プローブ()
ワーカーで WebGPU を確認する
破棄()
労働者を解雇する
instruct() コールバック
コールバック
いつ
onBeforeExecute
DOM を実行する前に、解析済みのアクションが準備完了
onBeforeStep
各ステップの前 — キャプチャノルム空間の { アクション、値、ポイント }
オンステップ
各ステップの後 — { アクション、値、ポイント、OK、詳細 }
オンステータス
人間が読める進行状況 (CLICK … 、 INPUT … 、 Done — … )
onRecapture
UIを変更する中間バッチステップ後の新しいOperatorCapture
オペレーターキャプチャー
フィールド
説明
キャンバス
フル SnapDOM キャンバス (マーカー オーバーレイ スペース)
幅 / 高さ
ビジョン ビットマップが薄暗くなります (モデルが見ているもの)
キャプチャ幅 / キャプチャ高さ
キャンバス全体のピクセルが薄暗くなります
css幅 / css高さ
画面上の整数の CSS サイズ
エンコード時
Promise<ArrayBuffer | null> — 推論 JPEG
世代
キャプチャごとにバンプが発生します。古い命令は拒否される
環境チェック
import { getWllamaEnvIssues , canLoadVlModelInBrowser } から 'browser-use-wasm' ;
const issues = getWllamaEnvIssues() ; // COOP/COEP、ブラウザ、WebGPU のヒント
テスト
browser-use-wasm は、実際の Chrome のブラックボックス E2E で検証されます (ライブ ページ → SnapDOM → ShowUI-2B → スクリーンショット マーカー チェック)。 Playwright は、事前に .model-cache/ に重みを格納する必要があります。このノードのダウンロード パスはテストと評価のみに使用され、アプリでの通常のブラウザーでは使用されません。
npx playwright インストールクロム
npm run queue:showui # E2E 用の ShowUI-2B を事前キャッシュします
npm run test # ShowUI-2B Gate — 実際の Chrome での 34 のブラウザ使用例
npm 実行テスト:ベンチ

ark # オプションのクロスモデル比較 (npm run queue:public first)
E2E コントラクトについては、.cursor/rules/blackbox-e2e.mdc を参照してください。
npm run eval:mind2web # オフライン Mind2Web グラウンディング — docs/mind2web-eval.md
npm run キャッシュ:miniwob && npm run eval:miniwob
について
pdufour.github.io/browser-use-wasm/
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to pdufour/browser-use-wasm development by creating an account on GitHub.

GitHub - pdufour/browser-use-wasm · GitHub
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
pdufour
/
browser-use-wasm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .cursor .cursor .github/ workflows .github/ workflows assets/ visual-diff assets/ visual-diff docs docs examples examples scripts scripts src src tests tests .gitignore .gitignore .npmignore .npmignore GEMINI.md GEMINI.md README.md README.md demo-vid.gif demo-vid.gif demo-vid.mp4 demo-vid.mp4 package-lock.json package-lock.json package.json package.json playwright.config.js playwright.config.js tried-methods.txt tried-methods.txt tsconfig.json tsconfig.json tsconfig.src.json tsconfig.src.json vite.browse-proxy.js vite.browse-proxy.js vite.coi-pages.js vite.coi-pages.js vite.config.js vite.config.js vite.config.src.js vite.config.src.js vite.coop-corp.js vite.coop-corp.js vite.demo-pages.js vite.demo-pages.js vite.fixtures.js vite.fixtures.js vite.model-cache.js vite.model-cache.js vite.wllama-plugins.js vite.wllama-plugins.js vite.wllama-wasm.js vite.wllama-wasm.js View all files Repository files navigation
Live demo → https://pdufour.github.io/browser-use-wasm/
browser-use in WASM — vision-language automation with no backend. browser-use-wasm runs the full browser-use loop client-side: SnapDOM screenshot → ShowUI grounding → live-page click/type/select. Default model is ShowUI-2B ; swap others from the built-in model selector. Chrome or Edge only (WebGPU + WASM workers).
demo-vid.mp4
How browser-use works
browser-use here means: look at the page like a human (screenshot), decide what to do (VLA), act on the real DOM at the grounded point.
┌─────────────┐ ┌──────────────┐ ┌─────────────┐ ┌──────────────┐
│ Live page │ ──► │ SnapDOM │ ──► │ ShowUI VLA │ ──► │ DOM action │
│ (iframe) │ │ screenshot │ │ WASM worker │ │ at [x, y] │
└─────────────┘ └──────────────┘ └─────────────┘ └──────────────┘
Capture — SnapDOM clones the target element (or iframe content) into a canvas. This is what browser-use “sees”.
Resize — the image is downscaled to the model’s vision token budget.
Infer — one ShowUI navigation call per user goal. The model outputs structured actions ( CLICK , INPUT , SELECT , ENTER , …) with normalized [x, y] on the screenshot only .
Execute — browser-use acts on the live page via elementFromPoint at that coordinate. No label-text DOM lookup for grounding.
One goal → one inference → execute all parsed actions. Mid-batch UI changes (e.g. CLICK then INPUT ) trigger an automatic re-capture before the next step.
Coordinates never come from live DOM layout, selectors, or OCR — only from vision inference on the SnapDOM buffer.
Run browser-use-wasm locally — the browser-use stack lives entirely in the tab.
Chrome or Edge with WebGPU + JSPI — browser-use inference runs here, not on a server
Dev server must send COOP/COEP headers (included in vite.config.js )
Public registry models download in the browser on first Load (~1.8 GB for ShowUI-2B Q4_K_M + mmproj); swap models from the built-in selector
npm install github:pdufour/browser-use-wasm#main
# or: yarn add pdufour/browser-use-wasm#main
examples/ is excluded from the published tarball (see .npmignore ).
Library is the root package ( browser-use-wasm ). The examples/ app links it locally via "browser-use-wasm": "file:.." .
npm install # root — tests, cache scripts, library dev
cd examples && npm install && cd .. # links browser-use-wasm from repo root
npm run dev # homepage at /, operator at /home/
npm run build # dist/lib/ + dist/examples/
npm run typecheck # src + examples
Embed browser-use in your app
browser-use-wasm exposes a small embed API for browser-use automation. High-level entry point: createWebOperator() → load model → SnapDOM capture → instruct(goal) .
import { createWebOperator } from 'browser-use-wasm' ;
const operator = createWebOperator ( {
// What to screenshot (default: document.body)
captureRoot : ( ) => document . getElementById ( 'app' ) ,
// What document to click/type in (default: host document; use iframe contentDocument for embeds)
targetDocument : ( ) => document . getElementById ( 'frame' ) ?. contentDocument ?? document ,
} ) ;
await operator . load ( {
onStatus : ( msg ) => console . log ( msg ) ,
} ) ;
const cap = await operator . capture ( ) ;
await cap . whenEncoded ; // wait for JPEG encode in worker
const result = await operator . instruct ( 'click Submit' , {
onStatus : ( msg ) => console . log ( msg ) ,
onBeforeStep : async ( step ) => {
// animate cursor before each CLICK / INPUT / …
} ,
} ) ;
console . log ( result . summary ) ; // e.g. "CLICK → INPUT → ENTER"
console . log ( result . steps ) ; // per-action ok/detail + capture-norm points
createWebOperator(options?)
Option
Default
Description
modelId
ShowUI-2B
Registry model id
wasmUrl
/wllama/wllama.wasm
Same-origin wllama WASM
captureRoot
document.body
Element SnapDOM captures
targetDocument
host document
Live page DOM tools act on
inferenceTimeoutMs
12000
Per-inference ceiling
onPerfMark
—
Capture pipeline timing hooks
WebOperator methods
Method
Description
load(opts?)
Download and load GGUF in the browser WASM worker
capture()
browser-use capture — SnapDOM → vision resize → async JPEG encode
instruct(task, opts?)
One ShowUI inference + execute all parsed browser-use actions
locate(label)
Ground click <label> on current capture without executing
hasCapture()
Whether a screenshot buffer is ready
clearCapture()
Drop screenshot (next instruct needs capture() again)
setModel(id)
Switch registry model (requires load() again)
probe()
Check WebGPU in the worker
dispose()
Terminate workers
instruct() callbacks
Callback
When
onBeforeExecute
Parsed actions ready, before any DOM execution
onBeforeStep
Before each step — { action, value, point } in capture-norm space
onStep
After each step — { action, value, point, ok, detail }
onStatus
Human-readable progress ( CLICK … , INPUT … , Done — … )
onRecapture
Fresh OperatorCapture after a UI-changing mid-batch step
OperatorCapture
Field
Description
canvas
Full SnapDOM canvas (marker overlay space)
width / height
Vision bitmap dims (what the model sees)
captureWidth / captureHeight
Full canvas pixel dims
cssWidth / cssHeight
Integer CSS size on screen
whenEncoded
Promise<ArrayBuffer | null> — inference JPEG
generation
Bumps on each capture; stale instructs are rejected
Environment checks
import { getWllamaEnvIssues , canLoadVlModelInBrowser } from 'browser-use-wasm' ;
const issues = getWllamaEnvIssues ( ) ; // COOP/COEP, browser, WebGPU hints
Tests
browser-use-wasm is validated with blackbox E2E in real Chrome — live page → SnapDOM → ShowUI-2B → screenshot marker checks. Playwright needs weights in .model-cache/ ahead of time; that Node download path is for tests and evals only , not normal browser-use in the app.
npx playwright install chrome
npm run cache:showui # pre-cache ShowUI-2B for E2E
npm run test # ShowUI-2B gate — 34 browser-use cases in real Chrome
npm run test:benchmark # optional cross-model comparison (npm run cache:public first)
See .cursor/rules/blackbox-e2e.mdc for the E2E contract.
npm run eval:mind2web # offline Mind2Web grounding — docs/mind2web-eval.md
npm run cache:miniwob && npm run eval:miniwob
About
pdufour.github.io/browser-use-wasm/
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
