---
source: "https://github.com/AmrDab/clawdcursor"
hn_url: "https://news.ycombinator.com/item?id=48488026"
title: "AI can control your desktop through scripts"
article_title: "GitHub - AmrDab/clawdcursor: MCP-powered fallback layer that lets AI agents execute tasks through the GUI when APIs, tools, or direct integrations are unavailable. Cross-OS, accessibility-first, local-only. · GitHub"
author: "AmDab"
captured_at: "2026-06-11T10:36:25Z"
capture_tool: "hn-digest"
hn_id: 48488026
score: 1
comments: 0
posted_at: "2026-06-11T09:11:51Z"
tags:
  - hacker-news
  - translated
---

# AI can control your desktop through scripts

- HN: [48488026](https://news.ycombinator.com/item?id=48488026)
- Source: [github.com](https://github.com/AmrDab/clawdcursor)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T09:11:51Z

## Translation

タイトル: AI はスクリプトを通じてデスクトップを制御できる
記事のタイトル: GitHub - AmrDab/clawdcursor: API、ツール、または直接統合が利用できない場合に、AI エージェントが GUI を介してタスクを実行できるようにする MCP を利用したフォールバック レイヤー。クロス OS、アクセシビリティ最優先、ローカルのみ。 · GitHub
説明: API、ツール、または直接統合が利用できない場合に、AI エージェントが GUI を介してタスクを実行できるようにする MCP を利用したフォールバック レイヤー。クロス OS、アクセシビリティ最優先、ローカルのみ。 - AmrDab/clawdcursor

記事本文:
GitHub - AmrDab/clawdcursor: API、ツール、または直接統合が利用できない場合に、AI エージェントが GUI を介してタスクを実行できるようにする MCP を利用したフォールバック レイヤー。クロス OS、アクセシビリティ最優先、ローカルのみ。 · GitHub
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
アムルダブ
/
爪カーソル
パブ

リック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
796 コミット 796 コミット .github .github ドキュメント ドキュメント ガイド ガイド ネイティブ ネイティブ perf perf スクリプト スクリプト シード レジストリ シード レジストリ src src テスト テスト .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json schema.snapshot.json schema.snapshot.json tsconfig.json tsconfig.json tsconfig.tests.json tsconfig.tests.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すべてのエージェントに安全なデスクトップ制御を提供するローカル MCP サーバー。
どのモデルでも。どのアプリでも。 1 つの MCP エントリ。地元限定。
クイックスタート ·
なぜ・
ツールボックス ·
仕組み ·
プラットフォーム ·
変更履歴
Clawd Cursor はローカル MCP サーバーです。一度インストールしてください。マシン上のツール呼び出しエージェント (Claude Code、Cursor、Windsurf、OpenClaw、Claude Agent SDK、独自のループ) は MCP 経由で接続し、実際のデスクトップを安全に制御できます。エージェントは人間と同じように、クリック、入力、画面の読み取り、アプリの起動、GUI の操作を行います。
雲はありません。デフォルトではテレメトリはありません。サーバーは 127.0.0.1 にバインドします。クラウド モデルを指定しない限り、スクリーンショットは RAM に残ります。 Ollama やその他のローカル モデルでは、マシンからは何も出ません。
単一の safety.evaluate() チョークポイント。すべてのツール呼び出しは、それが標準入出力経由のエディター ホストからのものであっても、HTTP 経由の外部エージェントからのものであっても、組み込みの自律ループからのものであっても、デスクトップにアクセスする前に 1 つのセーフティ ゲートを通過します。年齢

nt はこのパスをバイパスできません。
HTTP でのベアラー トークン認証。デーモンは 127.0.0.1:3847 にバインドします。すべての HTTP リクエストには認証が必要です: Bearer $(cat ~/.clawdcursor/token) 。デフォルトではローカルのみ。バインドアドレスは設定可能です。
人間が画面上で実行できるのであれば、AI も実行できます。 API がありませんか?統合はありませんか?問題ない。
不可能な仕事はありません。 GUI、マウス、キーボードがあれば、必要なものがすべて揃います。 「このアプリではそれができない」ということはなく、読み取り、クリック、キー、待機の正しいシーケンスがあるだけです。 Clawd Cursor はそれらすべてを提供します。
モデルに依存せず (Claude、GPT、Gemini、Llama、Kimi、Ollama など)、アプリに依存せず (アクセシビリティ、OCR、ビジョン フォールバックを介して任意のウィンドウを駆動)、OS に依存しない (1 つの PlatformAdapter で Windows、macOS、Linux X11、Linux Wayland をカバー) です。
第一選択ではなく、フォールバックとして使用してください。ネイティブAPIは存在しますか?使ってください。 CLIは存在しますか?使ってください。ファイルを直接編集できますか?そうしなさい。 Playwright の脚本はすでに完成していますか?それを使ってください。 Clawd Cursor は、クリック、レガシー アプリ、パブリック サーフェスのない GUI など、ラスト マイル用です。
UI 状態コンパイラ。 compile_ui は、アクセシビリティ ツリーと OCR を、画面の信頼スコア付けされた 1 つのマップに融合し、すべての要素に安定した el_NN ID がタグ付けされます。ピクセルではなく、{element_id, snapshot_id} によって要素に作用します。トークンはほぼ無料で、DPI、サイズ変更、レイアウトの変更にも耐えます。 find_button / find_field 意味によってターゲットを特定し、ID を渡します。
事後検証。結果的なアクションでexpectを渡し、clawdcursorは結果を確認し(非同期UIの場合は短い確定ウィンドウを使用)、UIが従わなかった場合は空虚な成功ではなくDEVIATIONを報告します。
正直なクロスプラットフォーム同等性。コンパイラ、セキュアフィールドのリダクション、座標処理は macOS 上でも完全に同等になりました。現在の外部エージェント (MCP) サーフェス

セーフティ ゲートを介して el_NN 参照を解決し、それがいつ既存のブラウザにアタッチされたかを明らかにします。ツールの名前は変更されませんでした。既存の権限許可リストは引き続き機能します。
完全なリストについては、変更履歴を参照してください。
ツールボックス — 7 つの複合ツール (推奨)
2 つのカタログが並べて出荷されます。ツールボックス (このセクション) は 7 つの複合ツールであり、それぞれに約 10 ～ 15 の動詞をカバーするアクション列挙型が含まれています。ツール (次のセクション) は 98 個の基礎となる詳細なプリミティブであり、動詞ごとに 1 つのスキーマがあります。
コンパウンドはデフォルトのサーフェスです。カタログのフットプリントは約 1,500 トークン (粒状の約 12 倍小さい) なので、小規模なモデルはプリミティブに溺れるのではなく、アクションの選択に集中できます。 Anthropic が使用しているのと同じコンピューター_20250124 シェイプなので、編集ホストはすでにその操作方法を知っています。
Computer ( { action : "key" , combo : "mod+s" } ) // Cmd+S / Ctrl+S に解決されます
アクセシビリティ ( { アクション : "呼び出し" 、名前 : "送信" } )
window ( { アクション : "open_app" 、名前 : "Outlook" } )
system ( { action : "ocr" } ) // OS レベルの OCR、LLM ビジョンなし
task ( { 命令 : "メモ帳を開いて hello と入力してください" } ) // シン エージェント ループに委任します
バッチ ( { ステップ : [ // N 回のコールを 1 往復にまとめます
{ 名前 : "アクセシビリティ" 、引数 : { アクション : "set_value" 、名前 : "To" 、値 : "amy@x.com" } } 、
{ 名前 : "アクセシビリティ" 、引数 : { アクション : "set_value" 、名前 : "件名" 、値 : "Hi" } } 、
{ 名前 : "コンピュータ" 、引数 : { アクション : "タイプ" 、テキスト : "ここに本文があります。" } }
] } )
クイックスタート
ゼロからデスクトップ上のツールを呼び出すエージェントまで 60 秒。
最も単純 — 任意の OS (現在は npm):
npm i -g clawdcursor
Windows および Linux ではそのまま動作します。 macOS では、Swift ツールチェーン (Xcode コマンド ライン ツール) が存在する場合、ネイティブ ヘルパーはインストール中に自動的に構築されます。そうでなかった場合 — 例:ツールチャのないマシン上の npm i -g

in — cd "$(npm root -g)/clawdcursor" && bashnative/build.sh で一度ビルドし、clawdcursor give を実行してアクセシビリティと画面録画を承認します (grant はアクセス許可を承認します。ヘルパーはコンパイルしません)。以下の OS インストーラー スクリプトは両方の手順を実行します。
または、OS ごとに 1 行 (リポジトリをクローンし、ビルドし、macOS ネイティブ ビルドを自動的に処理します):
powershell - c " irm https://clawdcursor.com/install.ps1 | iex "
macOS / Linux:
カール -fsSL https://clawdcursor.com/install.sh |バッシュ
次に:
clawdcursor 同意 --accept # 1 回限りのデスクトップ制御同意 (必須)
clawdcursor Doctor # 権限を確認し、(オプションで) LLM プロバイダーを構成します
clawdcursor エージェント # または `clawdcursor mcp` — 上の表を参照
インストーラーは ~/clawdcursor にクローンを作成し、 npm install を実行してビルドし、npm link をグローバル shim として作成します。実行時の状態は ~/.clawdcursor/ (認証トークン、pidfiles、ログ) にあります。エージェント ホストの構成は編集されません。その手順は以下にあります。
それを Claude Code、Cursor、Windsurf、または Zed に配線します。
// ~/.claude/settings.json (またはエディターの MCP 構成)
{
"mcpサーバー": {
"爪カーソル" : {
"コマンド" : "clawdcursor" ,
"args" : [ "mcp " , " --compact " ]
}
}
}
それだけです。エージェントに「Outlook を開いて、サラからの最新のメールに返信してください」と依頼し、メールが実行されるのを確認してください。
clawdcursor mcp をターミナルで自分で実行しないでください。サーバーが必要な場合、エディタは標準入出力経由で自動的に起動します。手動で実行するコマンドは、上記の install、consent、および Doctor のステップのみです。
編集者権限の許可リスト (クロード コード、カーソルなど)。エディターがツールごとの権限許可リスト ( mcp__clawdcursor__window などのキー) を維持している場合は、代わりにサーバー レベルのワイルドカード「mcp__clawdcursor」を使用してください。 1 つのエントリですべてのツールをカバーし、バージョン間でのツール名の変更の影響を受けません (ツールごと)。

ツールが追加、削除、または名前変更されるたびに、静かに中断されます。
macOSの初回起動時。 clawdcursor Grant を実行して権限ダイアログを表示し、[システム設定] → [プライバシーとセキュリティ] を開き、[アクセシビリティ] と [画面録画] の両方で ClawdCursor という名前のエントリを有効にします。 clawdcursor は、すべてのデスクトップ コントロールをこの単一のネイティブ アプリ ID に統合します。両方のエントリが必要です。
Linux: tesseract-ocr 、 python3-gi 、 gir1.2-atspi-2.0 、および (Wayland のみ) ydotool または wtype をインストールします。
「エージェントにコンピュータを使用させる」ツールのほとんどは、ブラウザ専用、単一 OS、またはビジョン専用です。 Clawd Cursor は、クロス OS、アクセシビリティ第一、MCP ネイティブのカーソルであり、すべてのコールがルーティングされる単一のセーフティ ゲートを備えています。
他のものにはない 2 つのメカニズム:
最安階層優先の設計。アクセシビリティツリー（無料）→OCR（安価）→スクリーンショット（中）→ビジョン（高価）;エージェントは必要な場合にのみ登るので、トークンのコストはタスクの難易度を追跡します。バッチ ツールは、決定論的なストレッチを 1 回のラウンドトリップにまとめて効率を高めます。
1 つのプロトコル、2 つのトランスポート。エディター ホストの場合は MCP over stdio、デーモンの場合は MCP over HTTP — 同じカタログ、同じ JSON-RPC エンベロープ。
脳がどこに存在するかによって、clawdcursor がどのように使用されるかが決まります。どちらのモードも並行して実行できます。デーモンとエディターによって生成された stdio の子は独立したプロセスです。
直接ツール — エージェントが主導権を握る
LLM が通話を選択します。 clawdcursor は、実際のデスクトップからの安全な作動と最新の観察を提供します。これは、独自の推論ループを持つエージェントにとっての主要なモードです。
フローチャートTB
task["ユーザー タスク"] --> ループ["外部エージェント LLM ループ<br/>計画、ツールの選択、検証"]
ループ --> 観察{"質問に答える最も安価な観察<br/>"}
観察 -- "obs·a11y — 無料<br/>accessibility.read_tree/find/get_element<br/>window.li

st/active" --> a11y["A11y 観察<br/>(構造化テキスト + 要素ハンドル)"]
観察 -- "obs·ocr — 安価な<br/>system.ocr<br/>a11y ツリーが空または疎である" --> ocr["OCR 観察<br/>(OS レベルのテキスト、ビジョン LLM なし)"]
観察 -- "obs·dom — 中<br/>browser.read_text / page_context<br/>WebView / Electron / Chrome" --> dom["DOM 観察<br/>(CDP、構造化ブラウザ コンテンツ)"]
観察 -- "obs·vision — 高価な<br/>コンピュータのスクリーンショット<br/>キャンバスのみまたはピクセル推論" --> ビジョン["ビジョン観察<br/>(LLM コンテキストへの画像バイト)"]
a11y --> 行動する
ocr --> 動作
ダム --> 行動する
ビジョン --> 行動
ループ -- "デリゲート サブタスク" --> handoff["タスク({命令:...})<br/>シン ループにハンドオフ"]
ハンドオフ --> Thinloop["シン エージェント ループ<br/>(デーモン LLM)"]
シンループ --> 安全性
act["デスクトップ上で動作<br/>computer.click/type/key/drag<br/>accessibility.invoke/set_value<br/>window.open_app<br/>system.shortcuts_run<br/>browser.click/type<br/>バッチ — 1回の呼び出しでNステップ"] --> 安全性
safety["単一のセーフティ ゲート<br/>safety.evaluate()<br/>許可 / 確認 / ブロック"] -- 許可 --> tools["clawdcursor ツール レジストリ<br/>98 粒状 + 7 複合"]
安全 -- ユーザーが必要 --> 確認["人による確認"] --> ツール
安全性 -- 拒否 --> ブロックされました["ブロックされました"]
ツール --> デスクトップ["リアルデスクトップ<br/>ネイティブアプリ ·

[切り捨てられた]

## Original Extract

MCP-powered fallback layer that lets AI agents execute tasks through the GUI when APIs, tools, or direct integrations are unavailable. Cross-OS, accessibility-first, local-only. - AmrDab/clawdcursor

GitHub - AmrDab/clawdcursor: MCP-powered fallback layer that lets AI agents execute tasks through the GUI when APIs, tools, or direct integrations are unavailable. Cross-OS, accessibility-first, local-only. · GitHub
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
AmrDab
/
clawdcursor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
796 Commits 796 Commits .github .github docs docs guides guides native native perf perf scripts scripts seed-registry seed-registry src src tests tests .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json schema.snapshot.json schema.snapshot.json tsconfig.json tsconfig.json tsconfig.tests.json tsconfig.tests.json vitest.config.ts vitest.config.ts View all files Repository files navigation
The local MCP server that gives any agent safe desktop control.
Any model. Any app. One MCP entry. Local-only.
Quickstart ·
Why ·
Toolbox ·
How it works ·
Platforms ·
Changelog
Clawd Cursor is a local MCP server . Install it once. Any tool-calling agent on the machine — Claude Code, Cursor, Windsurf, OpenClaw, Claude Agent SDK, your own loop — connects via MCP and gets safe control of the real desktop. The agent clicks, types, reads the screen, opens apps, and drives any GUI the same way a human would.
No cloud. No telemetry by default. Server binds to 127.0.0.1 . Screenshots stay in RAM unless you point a cloud model at them. With Ollama or any local model, nothing leaves the machine.
Single safety.evaluate() chokepoint. Every tool call — whether it comes from an editor host over stdio, from an external agent over HTTP, or from the built-in autonomous loop — routes through one safety gate before it touches the desktop. The agent cannot bypass this path.
Bearer-token auth on HTTP. The daemon binds to 127.0.0.1:3847 . Every HTTP request needs Authorization: Bearer $(cat ~/.clawdcursor/token) . Local-only by default; the bind address is configurable.
If a human can do it on a screen, your AI can do it too. No API? No integration? No problem.
No task is impossible. GUI plus a mouse plus a keyboard equals everything you need. There is no "I can't do that in this app" — only the right sequence of reads, clicks, keys, and waits. Clawd Cursor gives you all of them.
It's model-agnostic (Claude, GPT, Gemini, Llama, Kimi, Ollama, …), app-agnostic (drives any window via accessibility, OCR, or vision fallback), and OS-agnostic (one PlatformAdapter covers Windows, macOS, Linux X11, and Linux Wayland).
Use as a fallback, not first choice. Native API exists? Use it. CLI exists? Use it. Direct file edit possible? Do that. A Playwright script already wired up? Use that. Clawd Cursor is for the last mile — the click, the legacy app, the GUI with no public surface.
UI State Compiler. compile_ui fuses the accessibility tree and OCR into one confidence-scored map of the screen, every element tagged with a stable el_NN id. Act on an element by {element_id, snapshot_id} instead of pixels — it's near-free in tokens and survives DPI, resize, and layout shifts. find_button / find_field locate a target by meaning and hand you the id.
Reactive verification. Pass expect on a consequential action and clawdcursor confirms the outcome (with a short settle window for async UIs), reporting a DEVIATION when the UI didn't obey instead of a hollow success.
Honest cross-platform parity. The compiler, secure-field redaction, and coordinate handling were brought to full parity on macOS; the external-agent (MCP) surface now resolves el_NN refs through the safety gate and discloses when it attached to your existing browser. No tools were renamed — existing permission allowlists keep working.
See the changelog for the full list.
Toolbox — 7 compound tools (recommended)
Two catalogs ship side-by-side. The toolbox (this section) is 7 compound tools, each with an action enum that covers ~10-15 verbs. Tools (next section) is the 98 underlying granular primitives, one schema per verb.
Compound is the default surface. Catalog footprint is ~1,500 tokens (about 12× smaller than granular), which keeps small models focused on the action choice instead of drowning in primitives. Same computer_20250124 shape Anthropic uses, so editor hosts already know how to drive it.
computer ( { action : "key" , combo : "mod+s" } ) // resolves to Cmd+S / Ctrl+S
accessibility ( { action : "invoke" , name : "Send" } )
window ( { action : "open_app" , name : "Outlook" } )
system ( { action : "ocr" } ) // OS-level OCR, no LLM vision
task ( { instruction : "open Notepad and type hello" } ) // delegates to the thin agent loop
batch ( { steps : [ // collapse N calls into 1 round-trip
{ name : "accessibility" , arguments : { action : "set_value" , name : "To" , value : "amy@x.com" } } ,
{ name : "accessibility" , arguments : { action : "set_value" , name : "Subject" , value : "Hi" } } ,
{ name : "computer" , arguments : { action : "type" , text : "Body here." } }
] } )
Quickstart
Sixty seconds from zero to a tool-calling agent on your desktop.
Simplest — any OS (now on npm):
npm i -g clawdcursor
Works as-is on Windows and Linux. On macOS , the native helper builds automatically during install when the Swift toolchain (Xcode Command Line Tools) is present. If it wasn't — e.g. npm i -g on a machine without the toolchain — build it once with cd "$(npm root -g)/clawdcursor" && bash native/build.sh , then run clawdcursor grant to approve Accessibility + Screen Recording ( grant approves permissions; it does not compile the helper). The OS installer scripts below do both steps for you.
Or one line per OS (clones the repo, builds, and handles the macOS native build automatically):
powershell - c " irm https://clawdcursor.com/install.ps1 | iex "
macOS / Linux:
curl -fsSL https://clawdcursor.com/install.sh | bash
Then:
clawdcursor consent --accept # one-time desktop-control consent (required)
clawdcursor doctor # verify permissions + (optionally) configure an LLM provider
clawdcursor agent # OR `clawdcursor mcp` — see the table above
The installer clones into ~/clawdcursor , runs npm install , builds, and npm link s a global shim. Runtime state lives at ~/.clawdcursor/ (auth token, pidfiles, logs). It does not edit any agent host config — that step is below.
Wire it into Claude Code, Cursor, Windsurf, or Zed:
// ~/.claude/settings.json (or your editor's MCP config)
{
"mcpServers" : {
"clawdcursor" : {
"command" : " clawdcursor " ,
"args" : [ " mcp " , " --compact " ]
}
}
}
That's it. Ask your agent to "open Outlook and reply to the latest email from Sarah" and watch it run.
Don't run clawdcursor mcp in a terminal yourself — your editor launches it automatically over stdio when it needs the server. The only commands you run by hand are the install, consent , and doctor steps above.
Editor permission allowlist (Claude Code, Cursor, …). If your editor maintains a per-tool permission allowlist (keys like mcp__clawdcursor__window ), use the server-level wildcard "mcp__clawdcursor" instead. It covers every tool in one entry and is immune to tool renames across versions — per-tool entries silently break whenever a tool is added, removed, or renamed.
macOS first run. Run clawdcursor grant to walk through the permission dialogs, then open System Settings → Privacy & Security and enable the entry named ClawdCursor under both Accessibility and Screen Recording. clawdcursor consolidates all desktop control under this single native-app identity — both entries are required.
Linux: install tesseract-ocr , python3-gi , gir1.2-atspi-2.0 , and (Wayland only) ydotool or wtype .
Most "let an agent use the computer" tools are browser-only, single-OS, or vision-only. Clawd Cursor is the cross-OS, accessibility-first, MCP-native one — with a single safety gate every call routes through.
Two mechanisms the others don't have:
Cheapest-tier-first by design. Accessibility tree (free) → OCR (cheap) → screenshot (medium) → vision (expensive); the agent climbs only when it must, so token cost tracks task difficulty. The batch tool collapses deterministic stretches into one round-trip for additional efficiency.
One protocol, two transports. MCP over stdio for editor hosts, MCP over HTTP for daemons — same catalog, same JSON-RPC envelope.
Where the brain lives decides how clawdcursor is used. Both modes can run side-by-side — the daemon and editor-spawned stdio child are independent processes.
Direct tools — your agent drives
Your LLM picks the calls; clawdcursor supplies safe actuation and fresh observations from the real desktop. This is the primary mode for any agent with its own reasoning loop.
flowchart TB
task["User task"] --> loop["External agent LLM loop<br/>plans, chooses tools, verifies"]
loop --> observe{"Cheapest observation<br/>that answers the question"}
observe -- "obs·a11y — free<br/>accessibility.read_tree/find/get_element<br/>window.list/active" --> a11y["A11y observation<br/>(structured text + element handles)"]
observe -- "obs·ocr — cheap<br/>system.ocr<br/>a11y tree empty or sparse" --> ocr["OCR observation<br/>(OS-level text, no vision LLM)"]
observe -- "obs·dom — medium<br/>browser.read_text / page_context<br/>WebView / Electron / Chrome" --> dom["DOM observation<br/>(CDP, structured browser content)"]
observe -- "obs·vision — expensive<br/>computer.screenshot<br/>canvas-only or pixel reasoning" --> vision["Vision observation<br/>(image bytes into LLM context)"]
a11y --> act
ocr --> act
dom --> act
vision --> act
loop -- "delegate subtask" --> handoff["task({instruction:...})<br/>hand off to thin loop"]
handoff --> thinloop["Thin agent loop<br/>(daemon LLM)"]
thinloop --> safety
act["Act on the desktop<br/>computer.click/type/key/drag<br/>accessibility.invoke/set_value<br/>window.open_app<br/>system.shortcuts_run<br/>browser.click/type<br/>batch — N steps in 1 call"] --> safety
safety["Single safety gate<br/>safety.evaluate()<br/>allow / confirm / block"] -- allowed --> tools["clawdcursor tool registry<br/>98 granular + 7 compound"]
safety -- needs user --> confirm["Human confirmation"] --> tools
safety -- denied --> blocked["blocked"]
tools --> desktop["Real desktop<br/>native app ·

[truncated]
