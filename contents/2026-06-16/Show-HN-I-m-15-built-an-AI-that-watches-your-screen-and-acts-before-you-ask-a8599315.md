---
source: "https://github.com/Helmus101/weave"
hn_url: "https://news.ycombinator.com/item?id=48561017"
title: "Show HN: I'm 15, built an AI that watches your screen and acts before you ask"
article_title: "GitHub - Helmus101/weave · GitHub"
author: "anqer"
captured_at: "2026-06-16T21:53:56Z"
capture_tool: "hn-digest"
hn_id: 48561017
score: 5
comments: 1
posted_at: "2026-06-16T19:54:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I'm 15, built an AI that watches your screen and acts before you ask

- HN: [48561017](https://news.ycombinator.com/item?id=48561017)
- Source: [github.com](https://github.com/Helmus101/weave)
- Score: 5
- Comments: 1
- Posted: 2026-06-16T19:54:07Z

## Translation

タイトル: Show HN: 私は 15 歳です。画面を監視し、要求される前に行動する AI を構築しました。
記事タイトル: GitHub - Helmus101/weave · GitHub
説明: GitHub でアカウントを作成して、Helmus101/weave の開発に貢献します。
HN テキスト: フィードバックをいただけますか?

記事本文:
GitHub - Helmus101/weave · GitHub
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
ヘルムス101
/
織ります
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
72

コミット数 72 コミット アセット アセット 拡張子 拡張子 ネイティブ/ ocr ネイティブ/ ocr スクリプト スクリプト src src .gitignore .gitignore README.md README.md UI_ENGINEERING_REPORT.md UI_ENGINEERING_REPORT.md build-icon.py build-icon.py build-icon.sh build-icon.sh electric-builder.json electron-builder.jsonindex.htmlindex.htmlpackage-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.main.json tsconfig.main.json tsconfig.preload.json tsconfig.preload.json vite.config.mts vite.config.mts Weave_design_refinement.docx Weave_design_refinement.docx すべてのファイルを表示 リポジトリ ファイルのナビゲーション
macOS 用のローカルファーストのリレーションシップ & ライフ インテリジェンス エンジン。静かに織り上げるプライベートな空間
見たこと、行ったこと（画面 OCR + Gmail/カレンダー + 連絡先）を記憶し、それを実行します。
積極的なナッジであなたの背後を監視し、あなた自身の記憶からの質問に答え、行動することができます
ブラウザーとネイティブ macOS アプリで利用できます。
すべてはローカルに保存されます (sql.js 経由の SQLite)。クラウド呼び出しは LLM/検索に対してのみ発生します
ユーザーが設定する機能。
ホーム — 毎日の朝の概要 (やるべきこと、Weave が行ったこと、通知、世界のニュース、オプション)
スポーツ）、プロアクティブな「Do Now」カード、人間関係レーダー、および対応義務。
メモリ グラフ — スクリーン キャプチャは OCR 処理され、レイアウトによって再構築されます (サイドバー/メイン/パネルは
ごちゃ混ぜにせず、別々に保持し、安価なモデルで検索可能な 1 つの文にまとめて保存します
正確な事実の信頼できる情報源として完全な OCR が保持されています。
チャット — 自分の記憶とウェブに対して質問します。 Googleカレンダーを読み取ったり、
Gmail で「明日の会議は何ですか?」という質問に答えます。 / 「サラは私にメールをくれましたか?」と正確な日付/数字を引用します
完全なソースから逐語的に。
エクスプローラー — 生のメモリ ノードを参照します。
コミットメントトラッカー — 約束します

あなたが作った/借りているもの、OCR + Gmail から抽出。
応答負債 — 応答を待っている電子メール スレッド。
ルーティン / プレイブック — スケジュールされたブリーフィングまたはオンデマンドのブリーフィング (例: 15 分前までの会議ブリーフィング)
カレンダーイベント)。
ブラウザ エージェント (Chrome 拡張機能) — 目標を与えます。計画し、実際の Chrome タブで動作し、
比較された推奨成果物を過剰に納品します。 Cookie バナーを自動的に受け入れ、次の場合に回復します
スタックし（オプションのビジョンフォールバックを含む）、支払い/送信/確認をクリックすることはありません。
macOS コントロール エージェント — アクセシビリティ API を介して音声/目標によってネイティブ アプリを駆動します (オープン + プレイ)
Spotify/音楽、地図のルート案内、メモ/メールなどでの複数ステップのアクション)。
音声 — グローバル ショートカットにより、システム全体のマイク オーバーレイがポップされます。音声は文字に起こされ、次のいずれかです
アプリを制御したり、ブラウザー タスクを実行したり、[応答] タブで回答したりします。
プロアクティブ インテリジェンスは、グラフを認識した検索とメモリ減衰の基盤に基づいて実行されるため、提案が
グラフが大きくなるにつれて、答えはより鮮明になります。
npmインストール
npm run build:ocr # ネイティブ Apple Vision OCR ヘルパーをビルドします (macOS)
npm run dev:electron # フルアプリ: キャプチャ、同期、エージェント、音声、トレイ
npm run dev + http://127.0.0.1:5173 はレンダラのみのブラウザ プレビューを提供しますが、
ネイティブ ブリッジ — OCR、同期、エージェント、およびアカウントの動作には npm run dev:electron が必要です。
環境経由でキーを提供します (従来のアプリ内値は macOS キーチェーンに移行されます)。
画面録画 — OCR キャプチャに必要です。
アクセシビリティ — macOS コントロール エージェントに必要です (キーストローク / ネイティブ UI のクリック)。
Chrome拡張機能（ブラウザエージェント）
アプリを実行します (拡張機能がポーリングするローカルホスト取り込みサーバーが開始されます)。
chrome://extensions → 開発者モードを有効にする → 解凍したものをロード → extension/ フォルダーを選択します。
Do-Now 項目または音声によってブラウザー タスクをトリガーします。エージェントは専用のタブで作業します

nd
成果物を Weave にポストします。 WEAVE_INGEST_PORT を変更した場合は、それに合わせてください
拡張子/background.js 。
Gmail 作成、カレンダー イベント、ドライブ スコープが、元の読み取り専用権限を超えて追加されました。
以前に接続したアカウントは、設定で切断して再接続し、再同意する必要があります。
ローカル MCP サーバーは、Weave のメモリ グラフを Claude Desktop や ChatGPT などのクライアントに公開します。
npm run mcp # stdio
npm run mcp:http # ローカル HTTP
ツール: ステータス/ヘルス、メモリ検索、ask-weave/チャット、プロアクティブな提案、ルーチンと実行、
メモリノードの検査、コミットメント、応答負債、アクティビティのタイムライン。
オプションの環境: WEAVE_USER_DATA_DIR 、 WEAVE_MCP_PORT 、 WEAVE_MCP_HOST (デフォルトは 0.0.0.0 )。
Weave のクラウドとリモート アクセス設定が有効になっていない限り、リモート HTTP モードはオフになります。それはSupabaseを使用します
セッション認証を使用し、一致するサインインしたアカウントの名前空間のみを提供します。独自の HTTPS の背後に配置します
Web コネクタで使用する前に境界を確認してください。
クロードデスクトップ構成の例:
{
"mcpサーバー": {
「織り」 : {
"コマンド" : "npm" ,
"args" : [ " run " 、 " mcp " ]、
"cwd" : " /absolute/path/to/weave "
}
}
}
リモートリクエストには Authorization: Bearer <supabase_access_token> が含まれている必要があります。
src/main — Electron 特権サービス: SQLite 永続性 ( db/ )、キャプチャ ウォッチャー +
ocrBridge / ocrLayout (レイアウト認識 OCR)、検索 (ベクター + BM25 + グラフ)、インテリジェンス
(チャット合成)、プロアクティブ、ルーチン、MorningBrief、コミットメント/レスポンスDebt、
browserAgent + localIngest (拡張サーバー)、 macAgent 、 google 、 deepseek 、および IPC。
src/preload — 型付きのレンダラーセーフな API ブリッジ ( window.weave )。
src/renderer — React アプリ: ホーム、チャット、タスク、ルーチン、応答、エクスプローラー、設定、音声
オーバーレイ、およびデザイン システム スタイルシート (:root トークンによって駆動される軽量の編集テーマ)。
拡張子/ — MV3 Chro

me 拡張子: ブラウザー エージェントのバックグラウンド ワーカー + CDP 駆動の DOM アクション。
ネイティブ/ocr — 画面をキャプチャし、Apple Vision OCR (行ごとの境界あり) を実行する Swift CLI。
npm run typecheck 、 npm test 、および npm run build はクリーンなマシン上でパスします (または npm run beta:check )。
パッケージ化する前にネイティブ ヘルパーをビルドします: npm run build:ocr 、 npm run build:contacts 。
macOS での確認: 画面録画 + アクセシビリティ フロー、Google 接続/再接続、Apple 連絡先
同期、Chrome 拡張機能の読み込み + ブラウザー タスク、アカウントの切り替え、現在のアカウントの削除。
ベータ版ユーザーは次の点を理解する必要があります。連絡先調査は、明示的に有効になるまでローカルのみです。データの削除
現在のアカウントのローカル Weave データのみを削除します。ブラウザのプレビューはランタイムではありません。
署名/公証は手動で行われます。ベータ版のビルド前: リリース候補にタグを付け、Electron をパッケージ化します。
ネイティブ OCR + Apple Contacts バイナリを含むアプリ、Apple Developer ID で署名、
公証とステープル処理を行ってから、署名されたビルドを別の macOS マシンで検証します。
npm run dmg # build:ocr + build:contacts + signed dmg (CSC ID が必要)
注意事項
データベースは、sql.js を介したローカル SQLite ファイルです (ネイティブ アドオン コンパイルはありません)。ベクトル検索が座っています
決定論的な埋め込みとファイルに基づく永続性を備えたローカル アダプターの背後にあるため、実ベクトル DB
ウォッチャーや UI 呼び出しサイトに触れずにスワップインできます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
ダメージ
最新の
2026 年 4 月 28 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to Helmus101/weave development by creating an account on GitHub.

Can I have some feedback?

GitHub - Helmus101/weave · GitHub
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
Helmus101
/
weave
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
72 Commits 72 Commits assets assets extension extension native/ ocr native/ ocr scripts scripts src src .gitignore .gitignore README.md README.md UI_ENGINEERING_REPORT.md UI_ENGINEERING_REPORT.md build-icon.py build-icon.py build-icon.sh build-icon.sh electron-builder.json electron-builder.json index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.main.json tsconfig.main.json tsconfig.preload.json tsconfig.preload.json vite.config.mts vite.config.mts weave_design_refinement.docx weave_design_refinement.docx View all files Repository files navigation
Local-first relationship & life intelligence engine for macOS. Weave quietly builds a private
memory of what you see and do (screen OCR + Gmail/Calendar + contacts), then puts it to work:
it watches your back with proactive nudges, answers questions from your own memory, and can act
for you in the browser and in native macOS apps.
Everything is stored locally (SQLite via sql.js ); cloud calls happen only for the LLM/search
features you configure.
Home — a daily morning brief (to-do, what Weave did, notifications, world news, optional
sports), proactive "Do-Now" cards, a relationship radar, and response debt.
Memory graph — screen captures are OCR'd, reconstructed by layout (sidebar/main/panel are
kept separate, not jumbled), summarized into one searchable sentence by a cheap model, and stored
with the full OCR retained as the authoritative source for exact facts.
Chat — ask questions against your own memory + the web. It can read your Google Calendar and
Gmail to answer "what meetings tomorrow?" / "did Sarah email me?", and quotes exact dates/numbers
verbatim from the full source.
Explorer — browse the raw memory nodes.
Commitment tracker — promises you made/owe, extracted from OCR + Gmail.
Response debt — email threads awaiting your reply.
Routines / Playbooks — scheduled or on-demand briefings (e.g. a meeting brief ~15 min before
a calendar event).
Browser agent (Chrome extension) — give it a goal; it plans, works in a real Chrome tab, and
over-delivers a compared, recommended deliverable. Auto-accepts cookie banners, recovers when
stuck (incl. an optional vision fallback), and never clicks pay/submit/confirm.
macOS control agent — drive native apps by voice/goal via the Accessibility API (open + play
Spotify/Music, Maps directions, multi-step actions in Notes/Mail/etc.).
Voice — a global shortcut pops a system-wide mic overlay; speech is transcribed and either
controls an app, runs a browser task, or answers in the Responses tab.
Proactive intelligence runs on a graph-aware retrieval + memory-decay foundation, so suggestions
and answers get sharper as the graph grows.
npm install
npm run build:ocr # build the native Apple Vision OCR helper (macOS)
npm run dev:electron # full app: capture, sync, agents, voice, tray
npm run dev + http://127.0.0.1:5173 gives a renderer-only browser preview, but it has no
native bridge — OCR, sync, agents, and account behavior require npm run dev:electron .
Provide keys via the environment (legacy in-app values are migrated to the macOS Keychain):
Screen Recording — required for OCR capture.
Accessibility — required for the macOS control agent (keystrokes / clicking native UI).
Chrome extension (browser agent)
Run the app (it starts the localhost ingest server the extension polls).
chrome://extensions → enable Developer mode → Load unpacked → select the extension/ folder.
Trigger a browser task from a Do-Now item or by voice; the agent works in a dedicated tab and
posts its deliverable back to Weave. If you change WEAVE_INGEST_PORT , match it in
extension/background.js .
Gmail compose, Calendar events, and Drive scopes were added beyond the original read-only grant.
Accounts connected earlier must disconnect and reconnect in Settings to re-consent.
A local MCP server exposes Weave's memory graph to clients like Claude Desktop or ChatGPT.
npm run mcp # stdio
npm run mcp:http # local HTTP
Tools: status/health, memory search, ask-weave / chat, proactive suggestions, routines & runs,
memory-node inspection, commitments, response debt, activity timeline.
Optional env: WEAVE_USER_DATA_DIR , WEAVE_MCP_PORT , WEAVE_MCP_HOST (default 0.0.0.0 ).
Remote HTTP mode is off unless Weave's Cloud & Remote Access setting is enabled; it uses Supabase
session auth and serves only the matching signed-in account namespace. Put it behind your own HTTPS
boundary before using it with web connectors.
Example Claude Desktop config:
{
"mcpServers" : {
"weave" : {
"command" : " npm " ,
"args" : [ " run " , " mcp " ],
"cwd" : " /absolute/path/to/weave "
}
}
}
Remote requests must include Authorization: Bearer <supabase_access_token> .
src/main — Electron privileged services: SQLite persistence ( db/ ), the capture watcher +
ocrBridge / ocrLayout (layout-aware OCR), retrieval (vector + BM25 + graph), intelligence
(chat synthesis), proactive , routines , morningBrief , commitments / responseDebt ,
browserAgent + localIngest (extension server), macAgent , google , deepseek , and IPC.
src/preload — typed, renderer-safe API bridge ( window.weave ).
src/renderer — React app: Home, Chat, Tasks, Routines, Responses, Explorer, Settings, the voice
overlay, and the design-system stylesheet (light, editorial theme driven by :root tokens).
extension/ — MV3 Chrome extension: the browser agent's background worker + CDP-driven DOM actions.
native/ocr — Swift CLI that captures the screen and runs Apple Vision OCR (with per-line bounds).
npm run typecheck , npm test , and npm run build pass on a clean machine (or npm run beta:check ).
Build native helpers before packaging: npm run build:ocr , npm run build:contacts .
Verify on macOS: Screen Recording + Accessibility flows, Google connect/reconnect, Apple Contacts
sync, the Chrome extension load + a browser task, account switching and current-account deletion.
Beta users should understand: contact research is local-only until explicitly enabled; deleting data
removes only the current account's local Weave data; the browser preview is not the runtime.
Signing/notarization are manual. Before a beta build: tag a release candidate, package the Electron
app with the native OCR + Apple Contacts binaries included, sign with the Apple Developer identity,
notarize and staple, then validate the signed build on a separate macOS machine.
npm run dmg # build:ocr + build:contacts + signed dmg (CSC identity required)
Notes
The database is a local SQLite file via sql.js (no native addon compilation). Vector search sits
behind a local adapter with deterministic embeddings and file-backed persistence, so a real vector DB
can be swapped in without touching the watcher or UI call sites.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
DMG
Latest
Apr 28, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
