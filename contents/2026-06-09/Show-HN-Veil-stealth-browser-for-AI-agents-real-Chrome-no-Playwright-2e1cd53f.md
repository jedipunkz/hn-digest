---
source: "https://github.com/acunningham-ship-it/veilbrowser"
hn_url: "https://news.ycombinator.com/item?id=48460784"
title: "Show HN: Veil – stealth browser for AI agents (real Chrome no Playwright)"
article_title: "GitHub - acunningham-ship-it/veilbrowser: Stealth browser for AI agents — real Chrome over raw CDP, no Playwright/Puppeteer. TypeScript + MCP-native. Passes sannysoft 57/57, bypasses Cloudflare. · GitHub"
author: "TrustLayerDev"
captured_at: "2026-06-09T16:06:56Z"
capture_tool: "hn-digest"
hn_id: 48460784
score: 3
comments: 0
posted_at: "2026-06-09T13:20:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Veil – stealth browser for AI agents (real Chrome no Playwright)

- HN: [48460784](https://news.ycombinator.com/item?id=48460784)
- Source: [github.com](https://github.com/acunningham-ship-it/veilbrowser)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T13:20:15Z

## Translation

タイトル: Show HN: Veil – AI エージェント用のステルス ブラウザ (実際の Chrome には Playwright はありません)
記事のタイトル: GitHub - acuningham-ship-it/veilbrowser: AI エージェント用のステルス ブラウザ — 生の CDP 上の本物の Chrome、Playwright/Puppeteer なし。 TypeScript + MCP ネイティブ。 sannysoft 57/57 をパスし、Cloudflare をバイパスします。 · GitHub
説明: AI エージェント用のステルス ブラウザ — 生の CDP 上の本物の Chrome、Playwright/Puppeteer なし。 TypeScript + MCP ネイティブ。 sannysoft 57/57 をパスし、Cloudflare をバイパスします。 - acuningham-ship-it/veilbrowser

記事本文:
GitHub - acuningham-ship-it/veilbrowser: AI エージェント用のステルス ブラウザ — 生の CDP 上の本物の Chrome、Playwright/Puppeteer なし。 TypeScript + MCP ネイティブ。 sannysoft 57/57 をパスし、Cloudflare をバイパスします。 · GitHub
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
アカニンガム・シップ・イット
/
ベールブラウザ

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
アキュニンガム-シップイット/ベールブラウザ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github .github 資産 資産 dist dist 例 例 src src テスト テスト .gitignore .gitignore COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI エージェント用のステルス自動化ランタイム。生の CDP 上で本物の Chrome を駆動します。Playwright、Puppeteer、WebDriver は不要で、ランタイム依存関係はありません。
本物の Chrome を動かすベール: bot.sannysoft.com はすべて緑色で、その後 Cloudflare の JS チャレンジを真っ直ぐに通過します — パッチもプラグインもありません。
Instagram、Google、Reddit、Datadome、Akamai へ — Veil は Chrome です。同じバイナリ、同じ
TLS、同じ JS エンジン、人間のブラウザが持つ同じキャンバス/WebGL/フォント フィンガープリント。私たち
ブラウザを再実装しないでください (これは Chromium の 20 年にわたる 1,000 人のエンジニアの仕事であり、
手巻きエンジンの方が指紋採取は簡単ですが、難しくはありません)。という部品を交換していきます
引っかかるのは自動化層です。
なぜ劇作家/人形遣いではないのでしょうか?
これらは強力で、スクリプト化された QA に最適です。敵対的なサイトにいるエージェントには、次の 3 つの構造的な問題があります。
Veil は依存関係がありません — Node 24 / Bun はグローバル WebSocket を出荷するため、全体
CDP トランスポートは、当社が所有する最大 120 回線です。パッチするものは何もありませんし、漏れるものもありません。
「生の CDP よりも本物の Chrome」というアイデアは新しいものではありません — Python の nodriver
Camoufox (C++ パッチを適用した Firefox) はそれを先駆的に獲得し、同点のスコアを獲得しました。
純粋なステルス性の方が優れています。ベールは彼らを隠し通すとは主張していない。そのくさびがそこにあります
生活とエージェントがそれをどのように使用するか:
TypeScript ネイティブ。 JS/TS エージェント エコシステム (Vercel AI SDK、LangChain.js、MCP) には、
強力な raw-CDP がない

ステルス ドライバー — Playwright + ステルス プラグインでスタックしている、または
Python nodriver にシェルアウトします。ベールは欠けている部分です。
MCPネイティブ。 MCP サーバーを同梱しているため、どのエージェントもツールとしてステルス ブラウジングを利用できます。
接着剤ゼロ。
スクレイパーファーストではなく、エージェントファースト。アクセシビリティ ツリーの参照と人間による入力は、
スクレイピング スクリプト用ではなく、ブラウザを駆動する LLM です。
Python を使用していて、生のステルス機能だけが必要な場合は、nodriver または Camoufox を使用してください。これらは素晴らしいものです。
Veil は TypeScript エージェントと MCP ホスト用です。
前提条件: PATH 上の Chrome/Chromium (または VEIL_CHROME=/path/to/chrome )、および Bun。
バンインストール
bun run example/selftest.ts # 実際の Chrome を起動し、完全なチェーンを実行します
コード内:
"veilbrowser" から { ブラウザ } をインポートします。
const browser = ブラウザを待ちます。 launch({headless:false}) ; // 頭がいっぱい = 最もステルス
const page = ブラウザを待ちます。 newPage() ;
ページを待ちます。 goto ( "https://example.com" ) ;
// アクセシビリティ ツリーのスナップショット → 安定した整数参照 (セレクターなし)。
const snap = ページを待ちます。スナップショット ( ) ;
コンソール。ログ (スナップ.テキスト) ;
// [1] テキストボックス「検索」
// [2]ボタン「サインイン」
ページを待ちます。 fill ( 1 , "こんにちは" ) ; // ref によって動作 — 人間によるタイピング、ジッターのあるタイミング
ページを待ちます。 (2) をクリックします。 // 曲線ベジェ マウス パス、実際の CDP 入力
const png = ページを待ちます。スクリーンショット ( ) ; // ビジョンモデル用の PNG バッファー
ブラウザを待ちます。近い （ ） ;
ステルスの仕組み
Launch ( launcher.ts ) — 通常のプロファイルが使用するフラグを備えた本物の Chrome
Playwright が追加するオートメーション スイッチを除きます。 --disable-blink-features=AutomationControlled
navigator.webdriver をエンジン レベルで false に切り替えます。永続的な userDataDir
そのため、プロファイルは新しく作成されたものではなく、使用済み (履歴、Cookie) のように見えます。
トランスポート ( cdp.ts ) — 生の WebSocket、フラット セッション モード。私たちは決して電話しません
Runtime.enable — そのコマンドは

プライマリ CDP 検出ベクトル。ランタイム.評価
それがなくても動作します。
ページパッチ (steelth.ts) — addScriptToEvaluateOnNewDocument 経由で挿入
サイト コードの前、すべてのフレーム: webdriver 、 window.chrome 、を正規化します。
plugins 、 language 、 Permissions.query 、 WebGL ベンダー — パッチを適用したものになります
関数の toString() はネイティブに見えるため、パッチ自体は検出できません。保たれた
意図的に小さい。オーバーパッチはそれ自身の指紋です。
UA / クライアント ヒント ( page.ts ) — UA から HeadlessChrome トークンを削除します
および Sec-CH-UA ブランド ヘッダー。
人間による入力 ( human.ts ) — シード可能な PRNG は、湾曲したマウス パスを駆動し、
キーストロークのタイミングが不安定。
エージェント ツール (製品の残りの半分)
セールスポイントはステルス性だけではありません。エージェントがそれをうまく推進していることです。
snapshot() は、アクセシビリティからフラットな番号付きインデックスとしてページを返します。
ツリー (スクリーン リーダーが使用するセマンティック レイヤー)。エージェント破損の最大の原因 —
推測された CSS/XPath セレクター — はなくなりました。エージェントは安定した ref に基づいて動作します。
スクリーンショット() は PNG バッファを返し、ビジョンの準備が整います。
クリック / フィル / タイプで実際の CDP 入力を人間のダイナミクスで操作します。
waitFor(expr) は不安定な固定スリープを置き換えます。
検出スコアカード (測定、Chrome 148)
自分で実行してください: bun run example/detect.ts (ヘッドレス) または VEIL_HEADFUL=1 bun run example/detect.ts (ヘッドフル — Veil は独自の Xvfb を自動起動します。ラッパーは必要ありません)。
AMD Radeon (Renoir APU) ホスト、ANGLE/EGL 経由の実際のハードウェア GL で測定:
ライブターゲット (bun run example/hardtargets.ts 、住宅用 IP):
私たちがまだ主張していない正直なギャップ: インタラクティブなターンスタイル/reCAPTCHA、エンタープライズ
DataDome/Kasada、ログイン セッション、大容量の行動信頼。私たちは主張する前にテストします。
変化をもたらしたもの (スイートを再実行することでそれぞれを検証):
SwiftShaderではなく、実際のGPU。 --use-gl=角度 --use-angl

e=gl-egl は実際の
AMD GPU → 本物の自己一貫性のある WebGL フィンガープリント。ベンダーのなりすましがない = 嘘がない
CreepJS のピクセルハッシュをキャッチするため。
サーバー上で頭がいっぱいです。 Veil は独自の Xvfb ディスプレイを管理するため、「ヘッドフル」には何も必要ありません
デスクトップ。ヘッドレスレンダリングの癖と小さな画面の表示を排除します (33% → 0%)。
スリムなセルフゲート ステルス機能。最大の驚き: ステルス パッチ自体
それは「20%ステルス」信号でした。正しく起動された Chrome ではすでにレポートが表示されます
webdriver === false (人間の正しい値 — 未定義を強制するのは悪いです)、5
プラグイン、実際の Chrome オブジェクト。したがって、各パッチは値が次の場合にのみ起動されるようになりました。
本当に異常です。健全な Chrome では何も操作しません。表面が小さい = 何も検出されません。
Veil は MCP サーバー ( src/mcp.ts ) を出荷しています - すでに persoje に接続されています
( ~/.config/persoje/mcp.json )、8 つのツールを公開します: goto 、 snapshot 、 click 、 fill 、
type 、スクリーンショット、eval、close 。 persoje 独自の MCP を通じてエンドツーエンドで検証済み
クライアント (検出 → 移動 → スナップショット)。どの MCP ホストでも動作します。
{ "サーバー" : { "ベール" : {
"コマンド" : " /home/armani/.bun/bin/bun " ,
"args" : [ " run " , " /home/armani/projects/veil/src/mcp.ts " ],
"env" : {} // ヘッドフル + auto-Xvfb + 実 GPU (0% CreepJS)。
// より高速なサーバー モードには VEIL_HEADLESS=1 を設定します。
} } }
テスト
bun run example/selftest.ts # エンドツーエンド: 起動、ステルス、スナップショット、対話
bun run example/detect.ts # ボット検出スコアカード (bot.sannysoft.com など)
deno test testing/ * .test.ts # 単体テスト: PRNG、マウスパス、参照番号付け
単体テストの対象範囲は次のとおりです。
PRNG ( human.test.ts ) : xorshift32 決定論、範囲/int 境界、キーストロークのリズム、マウスのタイミング
スナップショット参照 ( snapshot.test.ts ) : 参照番号付け (1 ベース、連続、ギャップなし)、AX ツリー フィルタリング
CDP フレーミング ( cdp-messages.test.ts ) : JSON-RPC 構造、sessionId ルーティング、コマンド

d/応答相関
今日の作業 (Chrome 148 に対して検証済み):
ゼロデプ CDP ランタイム (生の WebSocket、フラット セッション モード)
ステルス起動 + ページスクリプトインジェクション ( --disable-blink-features=AutomationControlled )
UA/クライアント ヒント スクラブ (「HeadlessChrome」トークンなし)
WebGL バックエンドの選択 (ANGLE/EGL 経由のハードウェア GPU、または SwiftShader + ベンダー マスキング)
AX ツリー スナップショット → エージェントフレンドリーな対話のための安定した整数参照
人間のような入力: 湾曲したベジェ マウス パス、ジッターのあるキーストローク タイミング、実際の CDP 入力
PNG スクリーンショット (ビジョンモデル対応)
MCP サーバー ( src/mcp.ts ) — 標準出力 JSON-RPC; persoje、Claude、任意の MCP ホストが Veil をネイティブに駆動
敵対的フィンガープリント スイート — 継続的スコアリング (CreepJS、sannysoft、Datadome デモ)
Runtime.enable - すべての評価に対して分離された世界を介したリーク強化
マネージド Xvfb を介したヘッドフルオンサーバー。プロファイル + プロキシ プール
ビジョンベースの要素グラウンディング フォールバック (スパース AX ツリー、キャンバス アプリ)
ネットワーク傍受。応答キャプチャ。セッションの永続性とプロファイルのウォームアップ
タブごとの同時実行 (多数のタブ、1 つのソケット — トランスポートはすでにサポートしています)
AI エージェント用のステルス ブラウザ — 生の CDP 上の本物の Chrome、Playwright/Puppeteer はありません。 TypeScript + MCP ネイティブ。 sannysoft 57/57 をパスし、Cloudflare をバイパスします。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
Veil v0.1.0 — 最初のリリース
最新の
2026 年 6 月 9 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Stealth browser for AI agents — real Chrome over raw CDP, no Playwright/Puppeteer. TypeScript + MCP-native. Passes sannysoft 57/57, bypasses Cloudflare. - acunningham-ship-it/veilbrowser

GitHub - acunningham-ship-it/veilbrowser: Stealth browser for AI agents — real Chrome over raw CDP, no Playwright/Puppeteer. TypeScript + MCP-native. Passes sannysoft 57/57, bypasses Cloudflare. · GitHub
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
acunningham-ship-it
/
veilbrowser
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
acunningham-ship-it/veilbrowser
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github .github assets assets dist dist examples examples src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A stealth automation runtime for AI agents. Drives real Chrome over raw CDP — no Playwright, no Puppeteer, no WebDriver, zero runtime dependencies.
Veil driving real Chrome: bot.sannysoft.com all-green, then straight through Cloudflare's JS challenge — no patches, no plugins.
To Instagram, Google, Reddit, Datadome, Akamai — Veil is Chrome. Same binary, same
TLS, same JS engine, same canvas/WebGL/font fingerprint a human's browser has. We
don't reimplement the browser (that's Chromium's 20-year, 1000-engineer job, and a
hand-rolled engine is easier to fingerprint, not harder). We replace the part that
gets you caught: the automation layer.
Why not Playwright / Puppeteer?
They're powerful and great for scripted QA. For agents on hostile sites they have three structural problems:
Veil is dependency-free — Node 24 / Bun ship a global WebSocket , so the entire
CDP transport is ~120 lines we own. Nothing to patch, nothing to leak.
The "real Chrome over raw CDP" idea isn't new — Python's nodriver
pioneered it, and Camoufox (a C++-patched Firefox) scores even
better on pure stealth. Veil isn't claiming to out-stealth them. Its wedge is where it
lives and how agents use it :
TypeScript-native. The JS/TS agent ecosystem (Vercel AI SDK, LangChain.js, MCP) has
no strong raw-CDP stealth driver — it's stuck on Playwright + stealth plugins, or
shelling out to Python nodriver . Veil is that missing piece.
MCP-native. Ships an MCP server, so any agent gets stealth browsing as tools with
zero glue.
Agent-first, not scraper-first. Accessibility-tree refs and human input are built for
an LLM driving the browser, not for a scraping script.
If you're in Python and just want raw stealth, use nodriver or Camoufox — they're great.
Veil is for TypeScript agents and MCP hosts.
Prerequisites: Chrome/Chromium on PATH (or VEIL_CHROME=/path/to/chrome ), and Bun.
bun install
bun run examples/selftest.ts # launches real Chrome, runs the full chain
In your code:
import { Browser } from "veilbrowser" ;
const browser = await Browser . launch ( { headless : false } ) ; // headful = stealthiest
const page = await browser . newPage ( ) ;
await page . goto ( "https://example.com" ) ;
// Accessibility-tree snapshot → stable integer refs (no selectors).
const snap = await page . snapshot ( ) ;
console . log ( snap . text ) ;
// [1] textbox "Search"
// [2] button "Sign in"
await page . fill ( 1 , "hello" ) ; // act by ref — human typing, jittered timing
await page . click ( 2 ) ; // curved Bézier mouse path, real CDP input
const png = await page . screenshot ( ) ; // PNG buffer for a vision model
await browser . close ( ) ;
How the stealth works
Launch ( launcher.ts ) — a real Chrome with the flags a normal profile uses,
minus the automation switches Playwright adds. --disable-blink-features=AutomationControlled
flips navigator.webdriver to false at the engine level. Persistent userDataDir
so the profile looks used (history, cookies), not freshly minted.
Transport ( cdp.ts ) — raw WebSocket, flat session mode. We never call
Runtime.enable — that command is a primary CDP detection vector. Runtime.evaluate
works without it.
Page patch ( stealth.ts ) — injected via addScriptToEvaluateOnNewDocument
before any site code, on every frame: normalizes webdriver , window.chrome ,
plugins , languages , permissions.query , WebGL vendor — and makes the patched
functions' toString() look native so the patch itself can't be detected. Kept
deliberately small; over-patching is its own fingerprint.
UA / client hints ( page.ts ) — strips the HeadlessChrome token from the UA
and the Sec-CH-UA brand headers.
Human input ( human.ts ) — seedable PRNG drives curved mouse paths and
jittered keystroke timing.
Agent tooling (the other half of the product)
The selling point isn't only stealth — it's that agents drive it well :
snapshot() returns the page as a flat numbered index from the accessibility
tree (the semantic layer screen readers use). The #1 cause of agent breakage —
guessed CSS/XPath selectors — is gone. The agent acts on a stable ref .
screenshot() returns a PNG buffer, ready for vision grounding.
click / fill / type drive real CDP input with human dynamics.
waitFor(expr) replaces flaky fixed sleeps.
Detection scorecard (measured, Chrome 148)
Run it yourself: bun run examples/detect.ts (headless) or VEIL_HEADFUL=1 bun run examples/detect.ts (headful — Veil auto-starts its own Xvfb, no wrapper needed).
Measured on an AMD Radeon (Renoir APU) host, real hardware GL via ANGLE/EGL:
Live targets ( bun run examples/hardtargets.ts , residential IP):
Honest gaps we do not yet claim: interactive Turnstile/reCAPTCHA, enterprise
DataDome/Kasada, logged-in sessions, high-volume behavioural trust. We test before we claim.
What moved the needle (each verified by re-running the suite):
Real GPU, not SwiftShader. --use-gl=angle --use-angle=gl-egl drives the actual
AMD GPU → an authentic, self-consistent WebGL fingerprint. No vendor spoof = no lie
for CreepJS's pixel-hash to catch.
Headful on a server. Veil manages its own Xvfb display, so "headful" needs no
desktop. Eliminates the headless render quirks + tiny-screen tell (33% → 0%).
Slim, self-gating stealth. The biggest surprise: the stealth patches themselves
were the "20% stealth" signal. A correctly-launched Chrome already reports
webdriver === false (the right human value — forcing undefined is worse ), 5
plugins, a real chrome object. So each patch now fires only when the value is
genuinely anomalous; on healthy Chrome it's a no-op. Smaller surface = nothing to detect.
Veil ships an MCP server ( src/mcp.ts ) — already wired into persoje
( ~/.config/persoje/mcp.json ), exposing 8 tools: goto , snapshot , click , fill ,
type , screenshot , eval , close . Verified end-to-end through persoje's own MCP
client (discover → goto → snapshot). Any MCP host works:
{ "servers" : { "veil" : {
"command" : " /home/armani/.bun/bin/bun " ,
"args" : [ " run " , " /home/armani/projects/veil/src/mcp.ts " ],
"env" : {} // headful + auto-Xvfb + real GPU (0% CreepJS).
// Set VEIL_HEADLESS=1 for the faster server mode.
} } }
Testing
bun run examples/selftest.ts # end-to-end: launch, stealth, snapshot, interact
bun run examples/detect.ts # bot-detection scorecard (bot.sannysoft.com, etc.)
deno test tests/ * .test.ts # unit tests: PRNG, mouse paths, ref numbering
Unit tests cover:
PRNG ( human.test.ts ) : xorshift32 determinism, range/int bounds, keystroke cadence, mouse timing
Snapshot refs ( snapshot.test.ts ) : ref numbering (1-based, sequential, no gaps), AX-tree filtering
CDP framing ( cdp-messages.test.ts ) : JSON-RPC structure, sessionId routing, command/response correlation
Working today (verified against Chrome 148):
Zero-dep CDP runtime (raw WebSocket, flat session mode)
Stealth launch + page-script injection ( --disable-blink-features=AutomationControlled )
UA/client-hint scrub (no "HeadlessChrome" token)
WebGL backend selection (hardware GPU via ANGLE/EGL, or SwiftShader + vendor masking)
AX-tree snapshot → stable integer refs for agent-friendly interaction
Human-like input: curved Bézier mouse paths, jittered keystroke timing, real CDP input
PNG screenshots (vision-model ready)
MCP server ( src/mcp.ts ) — stdio JSON-RPC; persoje, Claude, any MCP host drives Veil natively
Adversarial fingerprint suite — continuous scoring (CreepJS, sannysoft, Datadome demo)
Runtime.enable -leak hardening via isolated worlds for all eval
Headful-on-server via managed Xvfb; profile + proxy pools
Vision-based element grounding fallback (sparse AX-trees, canvas apps)
Network interception; response capture; session persistence & profile warm-up
Per-tab concurrency (many tabs, one socket — transport already supports it)
Stealth browser for AI agents — real Chrome over raw CDP, no Playwright/Puppeteer. TypeScript + MCP-native. Passes sannysoft 57/57, bypasses Cloudflare.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Veil v0.1.0 — first release
Latest
Jun 9, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
