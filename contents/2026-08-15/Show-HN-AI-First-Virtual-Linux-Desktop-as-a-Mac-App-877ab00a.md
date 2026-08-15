---
source: "https://aifcc.franzai.com/"
hn_url: "https://news.ycombinator.com/item?id=49308944"
title: "Show HN: AI First Virtual Linux Desktop as a Mac App"
article_title: "AIFCC - AI First Computer in a Computer"
author: "franze"
captured_at: "2026-08-15T09:17:00Z"
capture_tool: "hn-digest"
hn_id: 49308944
score: 1
comments: 0
posted_at: "2026-08-15T08:54:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI First Virtual Linux Desktop as a Mac App

- HN: [49308944](https://news.ycombinator.com/item?id=49308944)
- Source: [aifcc.franzai.com](https://aifcc.franzai.com/)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T08:54:44Z

## Translation

タイトル: Show HN: Mac アプリとしての AI First 仮想 Linux デスクトップ
記事のタイトル: AIFCC - コンピューターの中の AI ファーストコンピューター
説明: Claude Code、Codex、Gemini、および MiMo は、Mac 上のハードウェア サンドボックス化された Linux デスクトップ内で完全な sudo を使用して実行されます。ゲストは逃げることができないため、実際の Mac はそのまま残ります。

記事本文:
AIFCC
概要
特長
パワーユーザー
AIエージェント向け
よくある質問
フィードバック
ダウンロード
AIFCC: あなたのコンピュータの中の AI ファーストコンピュータ
ルート管理者としての AI が、ユーザーのコンピュータからサンドボックス化された独自のコンピュータ上に存在します。 Claude Code、Codex、Gemini、および MiMo は、そのマシンを完全に制御して実行されます。そこで何をしても、あなたの Mac は影響を受けません。
AI は root として実行されます。それは尋ねません - それは尋ねます。
この仮想コンピューターは 1 つの目的のために存在します。それは、AI エージェントを独自のコンピューターの root sudo 管理者としてフル権限の「YOLO」モードで実行することです。許可のプロンプトやベビーシッターはありません - Claude Code、Codex、Gemini、MiMo はフルスピードでインストール、構成、構築、破壊を行います。コンピューター全体がハードウェア サンドボックス化されているため、これを許可しても問題ありません。エージェントが失う可能性がある最悪の事態はエージェント自身であり、Mac では決してありません。
権限チェックをオフにしてエージェントを実行するように作られています - --dangerously-skip-permissions 、 --yolo およびその友人。ここでは、AI が管理者です。
apt パッケージ、ツールチェーン、デーモン、システム サービス - 「いいですか?」ダイアログでは、各ステップの承認を待つ必要はありません。
/etc、cron、systemd、デスクトップ自体。ルートとはルートを意味し、エージェントはシステム全体をジョブに合わせて形成します。
目標を与えて立ち去ってください。あなたが何か他のことをしている間に、フルスピードでビルド、テスト、中断、修正が行われます。
Apple の仮想化フレームワークによってハードウェアが分離されています。エージェントがそこで行うことはすべてそこに留まり、あなたの Mac が触られることはありません。
エージェントが自分のコンピュータをゴミ箱に捨てたとしても、ワンクリックで数分で新しいコンピュータが得られます。あなたの Mac は気付かなかったのです。
Mac から隔離された完全な AI 対応 Linux コンピューターへの 3 つのステップ。
ドラッグ アンド ドロップ (両方向)。セットアップは必要ありません。
ファイルを Finder から Linux ウィンドウにドラッグします。 Linux デスクトップに表示されます。フォルダーも機能します。
Linux 内でファイルをドラッグすると、右下に「Send to Mac」ゾーンが表示されます。そこにドロップするか、右クリック→送信

マックに。このファイルは [Files from Linux] ウィンドウ (タイトル バーのトレイ ボタン) に表示され、Finder に直接ドラッグできます。
テキスト、画像、ファイル - 両方の方法。 Mac では ⌘C を Linux に貼り付けます。
Linux オーディオは Mac のスピーカーから再生されます。
⌘いつものようにTabキーを押します。 Linux は、停止するまで実行し続けます。
ホームフォルダーは、終了、更新、さらにはシステムの再構築でも存続します。
起動画面からいずれかを起動します。やめなさい、そしてそれは決して存在しませんでした。
デスクトップメニュー→工場出荷時設定にリセット。新品の Linux コンピューター。
タイトルバー ボタン、または aifcc スクリーンショット / aifcc レコード。
この Mac のみであり、LAN では使用できません。コマンドとして aifcc ssh を実行します。
他のアプリと同様に、App Store を通じて。ホームフォルダーは常に残ります。
4 4 つの AI エージェント、プリインストール
インストール手順はありません。自分のアカウント/キーでサインインして、次に進みます。
準備ができて。 chrome-devtools-mcp はブラウザの自動化のために事前に接続されています。
同じchrome-devtools-mcp配線。
同じchrome-devtools-mcp配線。
プレインストールされていますが、chrome-devtools-mcp はありません - その構成スキーマがそのキーを拒否します。 mimo mcp 経由で自分で配線します。
5 パワーユーザー: 設定リファレンス
アプリ独自のタブ オーダーの ⌘ の下にあります。
GUI で実行できることはすべて、CLI でも実行できます。ほとんどのコマンドは、スクリプトまたは AI の解析用に --json をサポートしています。
7 パワーユーザー: AI (MCP) で運転する
デスクトップをツールとして公開するモデル コンテキスト プロトコル サーバー (JSON-RPC 2.0、stdio)。どの MCP クライアントでもそれを駆動できますが、Mac 上では何も動作しません。
claude mcp add aifcc -- aifcc mcpserve # MCP クライアントに登録します
aifcc mcp register # ゲスト自身のクロード コードに接続します
aifcc mcpserve # サーバーを直接実行します
Factory-Reset には MCP ツールがなく、公開されることはありません。新規インストールでは両方のスイッチ (ホスト + ゲスト ミラー) がデフォルトでオンになります。明示的に OFF のままです。
機械可読インデックス: /llms.txt
完全な CLI テーブル: #cli-table 上記 (15 行)
完全な MCP ツール テーブル: 上の #mcp-tools (24 行)

サーバー: aifcc mcpserve - JSON-RPC 2.0 over stdio、プロトコル 2024-11-05
ユーザーの同意が必要 ([設定] → [MCP]) - 新規インストールではデフォルトでオン
出荷時設定へのリセットが MCP 経由で公開されることはありません
「aifcc」MCP サーバーを介して、私の Mac 上で実行されている完全な Linux デスクトップを制御できます。
利用可能なツール: aifcc_status、aifcc_diagnostics、aifcc_logs、aifcc_screenshot
(フルデスクトップまたはアクティブウィンドウ、トークンを保存するためのmaxWidth)、およびデスクトップコントロール -
aifcc_desktop_launch、aifcc_desktop_exec、aifcc_desktop_windows、aifcc_desktop_click /
移動/タイプ/キー/スクロール、aifcc_desktop_activate。まずスクリーンショットを撮ってから、行動してください。
工場出荷時設定へのリセットは意図的に公開されていません。
アプリの設定画面で「AI のプロンプトをコピー」で提供されるのと同じプロンプト - このページは、耐久性があり、リンク可能なバージョンです。
1つの価格です。階層、アドオン、使用量測定はありません。
Mac App Storeでダウンロード→
同意があれば ([設定] → [MCP]、デフォルトでオン)、AI は Linux 内でスクリーンショット、クリック、入力、スクロール、ウィンドウの管理、シェル コマンドの実行を行うことができます。Mac を工場出荷時設定にリセットすることは決してありません。
デフォルトでは Mac のデスクトップです ([設定] → [Mac に送信] で変更します)。 Linux ダイアログでは、正確なフォルダーに名前が付けられます。
セットアップには約 34 GB の空き容量が必要ですが、使用すると約 50 GB まで増加します。アプリを削除した後も、そのデータは Mac に残ります。削除したい場合は、最初に工場出荷時設定にリセットしてください。
設定（⌘、 ）→表示サイズ。次回の起動に適用されます。
はい - Apple のハードウェア仮想化サンドボックス。自分でフォルダーをマウントしない限り、ファイルにアクセスできません。
[ヘルプ] → [電子メール サポート] 、デスクトップ メニュー → [問題を報告]、または以下の [フィードバックとバグ レポート] セクションを選択します。バッジからビルドについて言及します。
バグを見つけた場合、紛らわしい問題に遭遇した場合、または AIFCC についてのアイデアをお持ちですか? 1 つの受信トレイですべてを読み取ります。
バグレポートの場合は、ビルド番号 ([設定] → [バージョン情報]) と期待した内容を含めてください。

起こること。
クリックすると連絡先メールアドレスが表示されます

## Original Extract

Claude Code, Codex, Gemini and MiMo run with full sudo inside a hardware-sandboxed Linux desktop on your Mac. The guest cannot escape, so your real Mac stays untouched.

AIFCC
Overview
Features
Power User
For AI Agents
FAQ
Feedback
Download
AIFCC: AI First Computer in your Computer
An AI as root admin, on its own computer, sandboxed from yours. Claude Code, Codex, Gemini and MiMo run with full control of that machine. Whatever they do in there, your Mac stays untouched.
The AI runs as root. It doesn't ask - it does.
This virtual computer exists for one thing: running AI agents in full-permission "YOLO" mode, as the root sudo admin of their own machine. No permission prompts, no babysitting - Claude Code, Codex, Gemini and MiMo install, configure, build and break things at full speed. That's safe to allow, because the whole computer is hardware-sandboxed: the worst an agent can lose is itself, never your Mac.
Made to run agents with permission checks off - --dangerously-skip-permissions , --yolo and friends. Here, the AI is the admin.
apt packages, toolchains, daemons, system services - no "may I?" dialog, no waiting for you to approve each step.
/etc, cron, systemd, the desktop itself. Root means root - the agent shapes the whole system to the job.
Give it a goal and walk away. It builds, tests, breaks and fixes at full speed while you do something else.
Hardware-isolated by Apple's Virtualization framework. Whatever the agent does in there stays in there - your Mac is never touched.
If an agent trashes its computer, one click gives it a brand-new one in minutes. Your Mac never even noticed.
Three steps to a complete, AI-ready Linux computer - walled off from your Mac.
Drag and drop, both directions. No setup needed.
Drag a file from the Finder onto the Linux window. It lands on the Linux Desktop . Folders work too.
Drag a file inside Linux - a "Send to Mac" zone appears bottom-right. Drop it there, or right-click → Send to Mac . It appears in the Files from Linux window (tray button in the title bar), where you can drag it straight into the Finder.
Text, images, files - both ways. ⌘C on the Mac, paste in Linux.
Linux audio plays through your Mac's speakers.
⌘Tab as always. Linux keeps running until you stop it.
Home folder survives quitting, updates, even a system rebuild.
Start one from the launch screen. Quit, and it never existed.
Desktop menu → Factory Reset. Brand-new Linux computer.
Titlebar button, or aifcc screenshot / aifcc record .
THIS Mac only, never your LAN. Run aifcc ssh for the command.
Through the App Store, like any app. Home folder always stays.
4 Four AI agents, preinstalled
No install step. Sign in with your own account/key and go.
Ready. chrome-devtools-mcp pre-wired for browser automation.
Same chrome-devtools-mcp wiring.
Same chrome-devtools-mcp wiring.
Preinstalled, but no chrome-devtools-mcp - its config schema rejects that key. Wire it yourself via mimo mcp .
5 Power user: Settings reference
Under ⌘, , in the app's own tab order.
Everything the GUI can do, the CLI can do too - most commands support --json for scripting or an AI to parse.
7 Power user: drive it with an AI (MCP)
A Model Context Protocol server (JSON-RPC 2.0, stdio) exposing the desktop as tools. Any MCP client can drive it - never anything on your Mac.
claude mcp add aifcc -- aifcc mcp serve # register with an MCP client
aifcc mcp register # wire into the guest's own Claude Code
aifcc mcp serve # run the server directly
factory-reset has no MCP tool - never exposed. Both switches (host + guest mirror) are on by default on a fresh install; an explicit OFF sticks.
Machine-readable index: /llms.txt
Full CLI table: #cli-table above (15 rows)
Full MCP tool table: #mcp-tools above (24 rows)
Server: aifcc mcp serve - JSON-RPC 2.0 over stdio, protocol 2024-11-05
Requires user consent (Settings → MCP) - on by default on a fresh install
factory-reset is never exposed over MCP
You can control a full Linux desktop running on my Mac via the "aifcc" MCP server.
Available tools: aifcc_status, aifcc_diagnostics, aifcc_logs, aifcc_screenshot
(full desktop or active window; maxWidth to save tokens), and desktop control -
aifcc_desktop_launch, aifcc_desktop_exec, aifcc_desktop_windows, aifcc_desktop_click /
move / type / key / scroll, aifcc_desktop_activate. Screenshot first, then act.
Factory reset is intentionally not exposed.
Same prompt the app's Settings screen offers via "Copy prompt for an AI" - this page is its durable, linkable version.
One price. No tiers, no add-ons, no usage metering.
Download on the Mac App Store →
With consent (Settings → MCP, on by default), an AI can screenshot, click, type, scroll, manage windows and run shell commands inside Linux - never your Mac, never Factory Reset.
Your Mac's Desktop by default (change it in Settings → Send to Mac). The Linux dialog names the exact folder.
It needs about 34 GB free to set up, and grows to roughly 50 GB as you use it. That data stays on your Mac after you delete the app - run Factory Reset first if you want it gone.
Settings ( ⌘, ) → display size. Applies next launch.
Yes - Apple's hardware virtualization sandbox. No access to your files unless you mount a folder yourself.
Help → Email Support , Desktop menu → "Report a Problem", or the Feedback & bug reports section below. Mention the build from the badge.
Found a bug, hit something confusing, or have an idea for AIFCC? One inbox reads all of it.
For bug reports, include your build number (Settings → About) and what you expected to happen.
Click to reveal the contact email
