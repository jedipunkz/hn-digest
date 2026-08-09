---
source: "https://anona-labs.github.io/mirror/"
hn_url: "https://news.ycombinator.com/item?id=49228375"
title: "Show HN: Open-source, searchable HTML workspace for Claude Code"
article_title: "Mirror, a live HTML view of your Claude Code conversations"
author: "anoop_kumar"
captured_at: "2026-08-09T04:16:46Z"
capture_tool: "hn-digest"
hn_id: 49228375
score: 1
comments: 0
posted_at: "2026-08-09T04:10:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source, searchable HTML workspace for Claude Code

- HN: [49228375](https://news.ycombinator.com/item?id=49228375)
- Source: [anona-labs.github.io](https://anona-labs.github.io/mirror/)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T04:10:50Z

## Translation

タイトル: Show HN: クロード コード用のオープンソースの検索可能な HTML ワークスペース
記事のタイトル: Mirror、クロード コードの会話のライブ HTML ビュー
説明: Mirror は、会話をライブの検索可能な HTML ワークスペースとして表示する Claude Code プラグインです。 API コストなし、ローカルホストのみ、依存関係なし。

記事本文:
クロード コードの会話のための、ライブで検索可能な HTML ワークスペース。
ターミナルでの作業を続けます。 Mirror はリンクを印刷し、そのページにはすべてのセッションが、実行に応じて更新されるクリーンなドキュメントとして表示されます。
読み取り可能なワークスペースとしてのターミナル セッション。
Claude Code は端末内で実行され、長い会話はすぐにスクロールして消えます。 Mirror は、レンダリングされたマークダウン、構文が強調表示されたコード、折りたたみ可能なツール呼び出し、すべてのセッションのリスト、セッション全体の全文検索などの会話をブラウザー タブにミラーリングするプラグインです。 Claude Code が既に書き込んだトランスクリプト ファイルを読み取るため、モデルを呼び出すことはなく、トークンのコストもかかりません。
クロード コードを初めて使用しますか?これは、端末内で実行される Anthropic のコーディング エージェントです。最初にインストールします ( Node.js が必要です)。
# クロード コードをインストールし、起動時にサインインします
npm install -g @anthropic-ai/claude-code
クロード
完全な手順: docs.claude.com/en/docs/claude-code 。 claude がターミナルで動作したら、以下に Mirror を追加します。
プラグインのクローンを作成します。
マシン上の任意の場所に置きます。
git クローン https://github.com/anona-labs/mirror.git
プラグインをロードした状態で Claude Code を起動します。
クローンを作成したフォルダーに Claude Code を指定します。作業したい任意のプロジェクトからこれを実行します。
クロード --plugin-dir /path/to/mirror
永続的なインストールをご希望ですか?代わりにローカル マーケットプレイスを通じて追加します。
クロード プラグイン マーケットプレイスに /path/to/mirror を追加
クロードプラグインのインストールミラー
ミラープリントのリンクを開きます。
起動時にこのような行が表示されます。ブラウザで開きます。
🪞 ミラーライブビュー: http://localhost:7842
要件: Python 3.8 以降 (macOS およびほとんどの Linux にすでに搭載されています) および Claude Code。 pip インストールまたはビルドするものは何もありません。開発中にプラグイン ファイルを編集する場合は、クロード コード内で /reload-plugins を実行します。
現在参加している会話が右側に表示されます。

スクロール位置を維持しながら、ターンごとに更新されます。緑色の点はライブであることを示します。
左側のサイドバーには、すべての Claude Code セッションがプロジェクトごとにグループ化され、最新性によって並べ替えられてリストされます。どれかをクリックして読んでください。
検索ボックスに入力すると、すべてのセッションにわたるメッセージが検索され、関連性によってランク付けされ、結果を現在のプロジェクトに限定することができます。 ⌘F / Ctrl-F を押して開いているセッション内を検索し、崩壊した思考とツール呼び出しに到達します。結果をクリックしてその結果にジャンプします。
上部のバーから: 明暗を切り替え、図のオンとオフを切り替え、思考やツールの呼び出しを非表示にし、セッションを Markdown にエクスポートし、再開コマンドをコピーします (またはサイドバーの行にマウスを置きます)。あなたの選択を記憶します。 /mirror を実行してリンクを再度開きます。
5分でその感触を味わってください。
これらのプロンプトのいずれかをクロードに与えて、Mirror がレンダリングするものを正確に確認してください。矢印、点、またはスワイプを使用して、例内を移動します。
また、上部のバーから: ⌘F / Ctrl‑F を押してセッション内を検索し、 Insights を開き、セッションを Markdown にエクスポートし、 Diagrams を切り替え、フィルター メニューから思考またはツール呼び出しを非表示にし、再開コマンドをコピーし (開いているセッションの上部バー、またはサイドバーの任意の行にマウスを置きます)、いつでも /mirror を実行してリンクを再度開きます。
ミラーは決してモデルを呼びません。セッションがすでに書き込んだトランスクリプトをレンダリングするため、所有している Claude サブスクリプションに基づいて実行されます。
サーバーは 127.0.0.1 にバインドします。あなたのトランスクリプトは、そのシークレットとファイルの内容とともに、あなたのマシンから離れることはありません。
すべてのセッションが 1 か所にまとめられ、プロジェクトごとにグループ化され、ライブ インジケーターとメッセージ数が表示されます。
SQLite を利用した、履歴全体にわたる関連性ランクの全文検索。必要なときにプロジェクトを範囲指定できます。 ⌘F / Ctrl-F を押して、折りたたまれたブロック内のテキストを含むセッション内を検索します。
セッション、最も忙しいプロジェクト、および

ターンごとのセッションごとの内訳、名前によるツール呼び出し、および思考はすべてローカル インデックスから読み取られます。 APIコストはかかりません。
ワンクリックでセッションを Markdown にエクスポートします。トップバーまたはサイドバーの行から claude --resume コマンドをコピーして、会話を元に戻します。
マークダウン、強調表示されたコード、テーブル、折りたたみ可能なツールの呼び出しと考え方。幅広いコンテンツがスペースを使用します。
Mermaid ブロックは、図ごとにソースを切り替えて、図としてレンダリングされます。貼り付けられた画像とスクリーンショット ツールの結果がインラインで表示されます。
ツール呼び出しの実行はグループに折りたたまれ、コード ブロックにはコピー ボタンが表示され、長い出力は「さらに表示」の後ろに折りたたまれます。
サーバー上の純粋な Python 標準ライブラリ。インストールするものや構築するものは何もありません。
ミラーは設定なしで機能します。デフォルトを変更するには、 ~/.mirror/config.json を作成します。
{
"theme" : "dark", // "dark" または "light"
"port" : 7842, // 優先ポート (ビジーの場合はフォールバック)
"auto_open" : false // サーバーの起動時にブラウザを開きます
}
よくある質問
いいえ、ミラーはクロード コードがディスクに書き込んだトランスクリプト ファイルを読み取ります。モデルコールは決して行われないため、請求には何も追加されません。
すべてがローカルのままです。サーバーは 127.0.0.1 でのみリッスンするため、ページには自分のマシンからのみアクセスできます。ミラーは何もアップロードしません。
macOS と Linux。 Python 3.8 以降が必要です。Python 3.8 は、ほとんどのシステムにすでに存在しています。
Mirror は、Claude Code が起動すると、小さなバックグラウンド サーバーを自動的に起動します。印刷されたリンクを開くだけです。
kill $(cat ~/.mirror/server.pid) を実行します。次回クロード コードを起動すると、再び開始されます。

## Original Extract

Mirror is a Claude Code plugin that shows your conversations as a live, searchable HTML workspace. No API cost, localhost only, zero dependencies.

A live, searchable HTML workspace for your Claude Code conversations.
Keep working in the terminal. Mirror prints a link, and that page shows every session as a clean document that updates as you go.
Your terminal sessions, as a readable workspace.
Claude Code runs in your terminal, where long conversations scroll away fast. Mirror is a plugin that mirrors those conversations into a browser tab: rendered markdown, syntax-highlighted code, collapsible tool calls, a list of all your sessions, and full-text search across them. It reads the transcript files Claude Code already writes, so it never calls a model and never costs a token.
New to Claude Code? It is Anthropic's coding agent that runs in your terminal. Install it first (it needs Node.js ):
# install Claude Code, then sign in when it launches
npm install -g @anthropic-ai/claude-code
claude
Full instructions: docs.claude.com/en/docs/claude-code . Once claude works in your terminal, add Mirror below.
Clone the plugin.
Put it anywhere on your machine.
git clone https://github.com/anona-labs/mirror.git
Start Claude Code with the plugin loaded.
Point Claude Code at the folder you just cloned. Run this from any project you want to work in.
claude --plugin-dir /path/to/mirror
Prefer a permanent install? Add it through a local marketplace instead:
claude plugin marketplace add /path/to/mirror
claude plugin install mirror
Open the link Mirror prints.
On startup you will see a line like this. Open it in your browser.
🪞 Mirror live view: http://localhost:7842
Requirements: Python 3.8+ (already on macOS and most Linux) and Claude Code. There is nothing to pip install or build. If you edit plugin files during development, run /reload-plugins inside Claude Code.
The conversation you are in appears on the right and updates after every turn, keeping your scroll position. A green dot marks it live.
The left sidebar lists every Claude Code session, grouped by project and sorted by recency. Click any one to read it.
Type in the search box to find any message across all your sessions, ranked by relevance, with a toggle to limit results to the current project. Press ⌘F / Ctrl‑F to find within the open session, reaching into collapsed thinking and tool calls. Click a result to jump to it.
From the top bar: toggle dark and light, turn diagrams on or off, hide thinking or tool calls, export a session to Markdown, and copy a Resume command (or hover any sidebar row). It remembers your choices. Run /mirror to reopen the link.
Get a feel for it in five minutes.
Give Claude one of these prompts and see exactly what Mirror renders. Use the arrows, the dots, or swipe to move through the examples.
And from the top bar: press ⌘F / Ctrl‑F to find within a session, open Insights , Export a session to Markdown, toggle Diagrams , hide thinking or tool calls from the filter menu, copy a Resume command (top bar for the open session, or hover any row in the sidebar), and run /mirror any time to reopen the link.
Mirror never calls a model. It renders the transcript your session already writes, so it rides on the Claude subscription you have.
The server binds to 127.0.0.1. Your transcripts, with their secrets and file contents, never leave your machine.
Every session in one place, grouped by project, with live indicators and message counts.
Relevance-ranked full-text search across your whole history, powered by SQLite, scoped to a project when you want. Plus ⌘F / Ctrl‑F to find within a session, including text inside collapsed blocks.
Totals across your sessions, the busiest projects, and a per-session breakdown of turns, tool calls by name, and thinking, all read from the local index. No API cost.
Export any session to Markdown with one click. Copy a claude --resume command from the top bar or any sidebar row to pick a conversation back up.
Markdown, highlighted code, tables, and collapsible tool calls and thinking. Wide content uses the space.
Mermaid blocks render as diagrams, with a per-diagram source toggle. Pasted images and screenshot tool results show inline.
Runs of tool calls collapse into a group, code blocks get a copy button, and long output folds behind "Show more".
Pure Python standard library on the server. Nothing to install, nothing to build.
Mirror works with no config. To change defaults, create ~/.mirror/config.json :
{
"theme" : "dark", // "dark" or "light"
"port" : 7842, // preferred port (falls back if busy)
"auto_open" : false // open the browser when the server starts
}
FAQ
No. Mirror reads the transcript files Claude Code writes to disk. It never makes a model call, so it adds nothing to your bill.
Everything stays local. The server only listens on 127.0.0.1, so the page is reachable only from your own machine. Mirror does not upload anything.
macOS and Linux. You need Python 3.8 or newer, which is already present on most systems.
Mirror starts a small background server automatically when Claude Code launches. Just open the printed link.
Run kill $(cat ~/.mirror/server.pid) . It starts again next time you launch Claude Code.
