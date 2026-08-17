---
source: "https://muvon.io/blog/introducing-octoweb-keyboard-first-ai-browser"
hn_url: "https://news.ycombinator.com/item?id=49328342"
title: "Show HN: Octoweb – A Keyboard-First Browser with AI as a Citizen"
article_title: "Introducing Octoweb: A Keyboard-First Browser Where the AI Is a Citizen, Not an Extension | Muvon"
author: "donk8r"
captured_at: "2026-08-17T09:30:33Z"
capture_tool: "hn-digest"
hn_id: 49328342
score: 1
comments: 0
posted_at: "2026-08-17T09:28:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Octoweb – A Keyboard-First Browser with AI as a Citizen

- HN: [49328342](https://news.ycombinator.com/item?id=49328342)
- Source: [muvon.io](https://muvon.io/blog/introducing-octoweb-keyboard-first-ai-browser)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T09:28:04Z

## Translation

タイトル: Show HN: Octoweb – AI を市民として備えたキーボードファーストのブラウザ
記事のタイトル: Octoweb の紹介: AI が拡張機能ではなく市民であるキーボードファーストのブラウザ |ムヴォン
説明: Octoweb は、コマンド パレット、高速アクセス スロット、どこでも readline バインディングなど、キーボードを使用して生活している人のために Rust で書かれた WebKit ブラウザです。そして、AI を両方向で第一級市民として扱います。サイドバーにはコーディングもできる Octomind 搭載のエージェントと、トークン効率の高い MCP が表示されます。
[切り捨てられた]

記事本文:
Octoweb のご紹介: AI が拡張機能ではなく市民となる、キーボードファーストのブラウザー |ムヴォン
サービス 製品 オープンソース ブログ チーム JP について 私たちと協力する ブログに戻る Octoweb のご紹介: AI が拡張機能ではなく市民であるキーボードファーストのブラウザ
Octoweb は、コマンド パレット、高速アクセス スロット、どこでも readline バインディングなど、キーボードを使用して生活している人のために Rust で書かれた WebKit ブラウザです。また、AI を両方向で第一級市民として扱います。サイドバーには Octomind を利用してコードを作成できるエージェントがあり、ブラウザー内にはトークン効率の高い MCP サーバーがあり、エージェントがそれを駆動できます。私たちは自分たちのためにそれを作りました。 5 か月後、それは私たちの毎日のブラウザになりました。 macOS、オープンソース、Apache-2.0。
Octoweb のご紹介: AI が拡張機能ではなく市民となる、キーボードファーストのブラウザー
すべてのブラウザに突然 AI が組み込まれるようになります。そして、彼らは皆、同じ方法でそれを行いました。Chrome を使用し、側面にチャット パネルをボルトで固定し、それを AI ブラウザと呼びます。ブラウザはエージェントの存在を認識しません。エージェントはブラウザにアクセスできません。見知らぬ二人が窓を共有している。
私たちは今年、エージェント (ランタイム、メモリ、コード検索、ファイルシステム層) の構築に費やしましたが、ある時点で、ブラウザが最初のコミットから AI 向けに設計されていればどのようなものになるのかという疑問が避けられなくなりました。 AI が追加されたわけではありません。市民としての AI — 双方向。エージェントはブラウザ内に存在し、ユーザーが見ているものを見ることができます。そして、ブラウザはエージェントが選択して使用できるツールです。
Octoweb は、Rust で書かれた WebKit ブラウザです。Electron も Chromium も使用せず、約 25,000 行あります。重なり合う 2 つのユーザー向けに作られています。キーボードを使って生活しているユーザーと、エージェントと一日中仕事をしているユーザーです。私たちは二人とも5ヶ月です

私たちは毎朝このブラウザーで仕事を始めます。ドキュメントを読んだり、独自の Web アプリをデバッグしたり、読書を続けている間、退屈なブラウザーの仕事を Octomind に任せたりすることができます。今日、それは 0.10.0 に到達し、それについてお知らせする時が来ました。
まず、マニアックな部分: マウスは必要ありません
AI が何かをする前に、Octoweb はキーストロークで考える人々のためのブラウザとしての地位を獲得する必要があります。設計ルールはシンプルで、すべてのアクションにショートカットがあり、クリックを必要とするものはありません。
その中心はコマンド パレット ( ⌘⇧P ) です。URL、検索クエリ、または訪問したページの断片を入力すると、開いているタブと履歴の間であいまい一致が行われ、一致の品質と訪問頻度によってランク付けされます。もう 1 回キーを押すと、開く、切り替える、または検索します。そして、筋肉の記憶は神聖なものであるため、パレットはリードラインを読み上げます。⌃A / ⌃E でジャンプ、⌃K / ⌃U でキル、⌃N / ⌃P で移動します。 Emacs やシェルを知っている人は、すでに Octoweb を知っています。
高速アクセス スロット — ⌘⇧1 – ⌘⇧0 は、現在のページを番号付きスロットに固定します。 ⌘1 – ⌘0 はどこからでもジャンプして戻ります。 10 ページ、それぞれ 1 回のキーストロークが再起動後も保持されます。これは、Vim マークに相当するブラウザです。 0.10.0 の新機能: ⌘⇧N は、現在のページを最初の空きスロットに固定し、すでにそこにある場合は固定を解除します。
MRU タブの循環 — ⌃N / ⌃P は、タブ バーにたまたま配置されている順序ではなく、最近使用した順序でタブを移動します。これは、実際に 2 つのページを切り替える方法です。
Vim スタイルのスクロール — ⌃D / ⌃U で半分のページ、⌃T / ⌃B で上下に移動します。
クリップボードへのスクリーンショット — ⌘S ビューポート、⌘⇧S フルページ。保存ダイアログや、後でクリーンアップするファイルはありません。
AI プロンプト ボックスの ⌃R — プロンプト履歴の逆増分検索。はい: readline の履歴検索をブラウザの AI 入力に組み込みました。これは、一度シェルに入れてしまえば、

どこにでもあります。
グローバル ショートカットは、macOS ネイティブで構成可能な CGEventTap を通じて実行されます。これが、Octoweb が現時点では macOS 専用である理由です。2 番目のオペレーティング システムを zip ファイルで出荷するのではなく、プラットフォーム (WKWebView、AppKit、CGEventTap) に依存しています。ブラウザ全体が 300 MB の Electron アプリではなく、単一の無駄のないバイナリである理由もここにあります。
サイドバーのエージェント - Octomind で実行中
⌘⇧A を押すと、エージェントがページ上にスライドして表示されます。チャット Web サイトを含む iframe ではなく、エージェント クライアント プロトコルを介して接続され、ローカルで実行される実際のエージェント プロセスです。
エージェントは、プラグアンドプレイ エージェント ランタイムである Octomind です。これは、ターミナル セッションと CI レビューを強化するものと同じです。そして、ここであなたがやらないことは、それを始めることです。 Octoweb は octomind acp octoweb:assistant 自体を起動します。セッションごとにサンドボックス化されたサブプロセスで、終了した場合は再起動され、次回起動時に会話がそのまま再開されます。 Octomind を一度インストールすると、そこからサイドバーにインストールされます。
その octoweb:assistant タグは、コミュニティ タップ レジストリ (モデル、システム プロンプト、ツール) からのマニフェストであり、自動的にフェッチされてアセンブルされます。手動で作成された構成ファイルはゼロです。それ以降は、読んでいるページについて質問したり、そのページに関するコードの説明を受けたり、時間がないスレッドを要約したりすることができます。応答がストリーミングされ、セッションが持続して再開され、豊富な出力カードを使用してコマンドをスラッシュすることができます。
そして、これは本物のエージェントであるため、サイドバーは意見が並ぶチャット ボックスではなく、機能します。タグをdeveloper:rustに交換すると、ブラウザからコードを作成できます。エージェントはサンドボックス化されたワークスペースでファイルを編集し、ツール呼び出しはライブステータスとともにクリック可能な行として表示され、散文よりも充実したもの（フォーム、テーブル、小さなインタラクティブUI）を表示したい場合は、A2UIサーフェスを画面内でレンダリングします。

彼のサイドバー。 Octoweb のバグを Octoweb から修正しましたが、問題が発生したページは次のタブに表示されていました。
エージェントタグは編集可能です。サイドバーのヘッダーにはデフォルトで octoweb:assistant が表示されますが、Octomind が知っている任意のタグ (developer:rust 、独自のカスタム エージェント) を入力すると、サイドバーがその場でタグに再接続します。ブラウザは、どの脳に話しかけているかを気にしません。
ブラウザーが AI プロバイダーを呼び出すことはありません。すべてのモデル呼び出しは、ブラウザ プロセスの外部で Octomind によって行われ、エージェントのファイル システム アクセスはそのワークスペースにサンドボックス化されます。エージェントがどこかにデータを送信しない限り、データがマシンから離れることはありません。どのエージェントがどのプロバイダーを使用するかは完全にあなたの判断です。
セッションは双方向です。 0.10.0 以降、Octomind ランタイムによって挿入されたメッセージ (スペシャリストのタップからの返信、スケジュールされたジョブの終了、Webhook の起動) は、会話中に独自のラベル付きバブルとしてサイドバーに到着します。ブラウザは、エージェントが報告を返す場所になります。
小さなドアもあります。任意のページでテキストを選択し、 ⌘⇧E を押すと、インライン編集モーダルが選択内容を変換します (書き換え、要約、翻訳)。同じタップから独自の専用エージェント (octoweb:editor) を実行し、独自のプロンプト履歴、⌃R などを備えています。
ツールとしてのブラウザ: 内部から見た MCP
これは他の誰も構築していない方向性であり、私たちが最も気に入っている部分です。 Octoweb は、ブラウザ内で MCP サーバー ( l​​ocalhost:3434/mcp ) を実行します。任意の MCP クライアント (Claude Desktop、Octomind セッション、独自のスクリプト) を接続して実行できます。ナビゲーション、タブ、クリック、入力、スクロール、スクリーンショット、コンテンツ抽出、コンソール出力、ネットワーク アクティビティをカバーする 26 個のツールが含まれています。
ブラウザの自動化については、一度は見たことがあると思います。ここでの違いは、エージェントがヘッドレス Chromium を運転していないことです。

atacenter には Cookie がなく、背面にボット検出ターゲットが付いています。視聴中にブラウザ、つまりセッション、ログイン、タブが動作します。そして、ツールの設計では「視聴中」を重視しています。
ナビゲーションがフォーカスを奪うことはありません。 browser_navigate は常にバックグラウンドで開きます。表示している内容を変更できるのは 1 つのツール (browser_switch_tab) だけです。エージェントは、カーソルの下からページを引き出すことなく、背後にある 10 個のタブで調査できます。あなたの代わりではなく、あなたのそばで閲覧します。
嘘をつかないクリック。 browser_click は、要素が存在し、安定し、障害物がなくなるまで再試行します。要素が何かで覆われている場合は、エラーが表示されます。 browser_snapshot は、iframe、Shadow DOM、およびリスナーのみのクリッカブルを貫通する安定した @N ref を持つインタラクティブな要素のマップを返します。 browser_wait は、ページの読み込みだけでなく、SPA の準備状況についても認識しています。
目も付属。 browser_console_messages とbrowser_network_requests は、エージェントにページのコンソール エラーと、ステータスとタイミングを含むフェッチ/XHR アクティビティを提供します。この 2 つは、DevTools をツールとして開く場合に使用します。
境界にはガードレール。スナップショットやページ コンテンツから流出するテキストはサニタイザーを通過します。カード番号はプロンプトに到達する前に編集されます。
トークンは予算です。すべてのツールの回答は、デバッガーではなくモデルに合わせて作成されます。browser_snapshot は、生の HTML の代わりにコンパクトな要素マップを返します (get_html ツールは意図的に削除しました)。コンソールとネットワークのログはキャップ付きでフィルター可能であり、@N refs は、エージェントが CSS セレクターを推測するコンテキスト ウィンドウの半分を書き込むことがないことを意味します。エージェントのターンにはトークンが課金されます。ブラウザはそれを尊重します。
エージェント ランタイムが Octomind の場合、セッション内の 1 行で接続できます。
/mcp octoweb http://localhost:3434/mcp を追加
エージェントy

すでに話している相手は、会話の途中でも再起動することなく、ドキュメントを開いて読んだり、フォームに記入したり、ページのコンソールが何を叫んでいるかを確認したりできるようになります。これが、ブラウザの雑用が私たちの To Do リストから外される方法です。エージェントはバックグラウンド タブでバグを再現し、コンソールを読み取り、失敗したリクエストを返します。クリックする気のなかった管理パネルが表示されます。フォアグラウンド タブで読み取りを続けながら、ログインしたブラウザーでセッションを使用して、デプロイが実際にレンダリングされたかどうかを確認します。
あなたと一緒に学ぶブラウザ
もう 1 つの実験: プロアクティブな学習です。 Octoweb は、一定の間隔 (デフォルトでは 30 分) で、完全に独立した 2 番目のエージェントである octoweb:learning タグを起動します。このエージェントは、開いているタブと最近の履歴を調べ、作業内容を Octomind のメモリに抽出します。次のセッションでは、サイドバー アシスタントは、あなたが 3 日間閲覧してきたドキュメントをすでに認識しています。
これはデフォルトでオンになっており、[設定] のトグルを 1 つ押すとオフになります。これはローカルであり、間隔は自由に設定でき、2 つのエージェントは完全に独立しており、どちらか一方でも実行されます。私たちは、ブラウザ AI が実際に興味深いのはアンビエント メモリであると考えており、デモではなく、その正直な v1 を出荷したいと考えています。
はい、それは単なる優れたブラウザでもあります
ブラウジングがあなたが残したものよりも悪い場合、これは何も問題ではありません。基本的なものはそこにありますが、それより優れているものもいくつかあります。
スマート タブの休止状態 — バックグラウンド タブは、RAM を認識するしきい値によりメモリ負荷がかかるとフリーズします。 50錠はライフスタイルであり、漏れではありません。そして 0.10.0 以降、開いているタブ (デフォルトでは 500 個、設定可能) に上限が設けられ、それを超えると最も最近使用されていないタブが静かに閉じられます。これは、タブをため込む人のためのエアバッグです。
組み込みのコンテンツ ブロック — WebKit のネイティブ コンテンツ ルール リストを介したトラッカーと広告。拡張機能は必要ありません。
コールドスタートステーション

y cold — ファビコンはデータ URI としてキャッシュされるため、ブラウザを起動しても、起動するまでネットワーク リクエストは発生しません。
セッションの復元、WebAuthn パスキー、PDF および DOCX の表示、CSS Custom Highlight API でのページ内検索、ネイティブのフルスクリーン、適切なポップアップ ウィンドウの処理。
拡張機能についてはどうですか?正直なステータス: Safari Web Extensions の統合 (1Password、Bitwarden、およびその友人向け) はコードに完全に実装されており、Apple のみが付与できる承認である Apple の com.apple.developer.web-browser 資格でブロックされています。着陸するまで、コンテンツ スクリプトは挿入されません。拡張機能がロードマップ項目であるかのように振る舞うのではなく、そのことをお伝えしたいと思います。
Octoweb は 3 月下旬に実験として開始されました。ACP サイドバーを備えたキーボード ブラウザを 1 週間で使用できるようになるでしょうか? 0.6.0 では、タブ休止状態と最初の MCP 自動化ツールが搭載されました。 0.7.0 では、永続的なチャット セッション、パスキー、およびスラッシュ コマンドが導入されました。 0.8.0 は、アクション性ハーネス、独自の MCP サーバーを通じてブラウザをテストする e2e スイート、エージェント サンドボックス、構成可能なキーバインディング、A2UI サーフェスなど、大きなものでした。 7 月下旬にカットされた 0.9.0 は、アカウント認証と安定化作業のロングテールを備えた ACP v1.2.0 に移行しました。そして今日の 0.10.0 は、固定されたタブ、開いているタブの上限、自動的に更新されるアカウント クォータ ステータス、サイドバーに表示される専門家のメッセージ、さらにメモリ リークとオーフの修正など、毎日の使用でロードマップに影響を与えます。

[切り捨てられた]

## Original Extract

Octoweb is a WebKit browser written in Rust for people who live on the keyboard — command palette, fast-access slots, readline bindings everywhere. And it treats AI as a first-class citizen in both directions: an Octomind-powered agent in the sidebar you can even code with, and a token-efficient MCP
[truncated]

Introducing Octoweb: A Keyboard-First Browser Where the AI Is a Citizen, Not an Extension | Muvon
Services Products Open Source Blog Team About EN Work with us Back to Blog Introducing Octoweb: A Keyboard-First Browser Where the AI Is a Citizen, Not an Extension
Octoweb is a WebKit browser written in Rust for people who live on the keyboard — command palette, fast-access slots, readline bindings everywhere. And it treats AI as a first-class citizen in both directions: an Octomind-powered agent in the sidebar you can even code with, and a token-efficient MCP server inside the browser so your agents can drive it. We built it for ourselves; five months in, it is our daily browser. macOS, open source, Apache-2.0.
Introducing Octoweb: A Keyboard-First Browser Where the AI Is a Citizen, Not an Extension
Every browser suddenly has AI in it. And every one of them did it the same way: take Chrome, bolt a chat panel onto the side, call it an AI browser. The browser doesn't know the agent exists. The agent can't touch the browser. Two strangers sharing a window.
We spent this year building agents — a runtime , memory , code search , a filesystem layer — and at some point the question became unavoidable: what would a browser look like if it were designed for AI from the first commit? Not AI added . AI as a citizen — in both directions. The agent lives in the browser and can see what you see. And the browser is a tool the agent can pick up and use.
So we built it — for ourselves, the way you build a tool you intend to live in. Octoweb is a WebKit browser written in Rust — no Electron, no Chromium, about 25,000 lines — made for two overlapping audiences: people who live on the keyboard, and people who work with agents all day. We're both, and five months later it's the browser we start every morning in: reading docs, debugging our own web apps, handing the boring browser chores to Octomind while we keep reading. Today it hits 0.10.0, and it's time to tell you about it.
First, the geek part: no mouse required
Before the AI does anything, Octoweb has to earn its place as a browser for people who think in keystrokes. The design rule was simple: every action has a shortcut, nothing requires a click.
The center of it is the command palette ( ⌘⇧P ) — type a URL, a search query, or any fragment of a page you've visited, and it fuzzy-matches across open tabs and history, ranked by match quality and visit frequency. One more keystroke opens, switches, or searches. And because muscle memory is sacred, the palette speaks readline: ⌃A / ⌃E to jump, ⌃K / ⌃U to kill, ⌃N / ⌃P to move. If your fingers know Emacs or a shell, they already know Octoweb.
Fast-access slots — ⌘⇧1 – ⌘⇧0 pins the current page to a numbered slot; ⌘1 – ⌘0 jumps back from anywhere. Ten pages, one keystroke each, persisted across restarts. It's the browser equivalent of Vim marks. New in 0.10.0: ⌘⇧N pins the current page to the first free slot — and unpins it if it's already there.
MRU tab cycling — ⌃N / ⌃P walks tabs in most-recently-used order, the way you actually switch between two pages, not the order they happen to sit in the tab bar.
Vim-style scrolling — ⌃D / ⌃U for half pages, ⌃T / ⌃B for top and bottom.
Screenshots to clipboard — ⌘S viewport, ⌘⇧S full page. No save dialog, no file to clean up later.
⌃R in the AI prompt box — reverse incremental search through your prompt history. Yes: we put readline's history search in the browser's AI input, because once you've had it in a shell you want it everywhere.
Global shortcuts run through CGEventTap, macOS-native and configurable. And this is why Octoweb is macOS-only for now — it leans on the platform (WKWebView, AppKit, CGEventTap) instead of shipping a second operating system in a zip file. That's also why the whole browser is a single lean binary instead of a 300 MB Electron app.
The agent in the sidebar — running on Octomind
Press ⌘⇧A and an agent slides in over the page. Not an iframe with a chat website in it — a real agent process, running locally, connected over the Agent Client Protocol .
The agent is Octomind , our plug-and-play agent runtime — the same one that powers our terminal sessions and CI reviews. And here's the part you don't do: start it. Octoweb launches octomind acp octoweb:assistant itself — a sandboxed subprocess per session, restarted if it dies, resumed with your conversation intact on the next launch. Install Octomind once, and the sidebar takes it from there.
That octoweb:assistant tag is a manifest from the community tap registry : model, system prompt, and tools, fetched and assembled automatically. Zero config files written by hand. From then on you can ask about the page you're reading, get code on it explained, or summarize a thread you don't have time for — with responses streaming in, sessions that persist and resume, and slash commands with rich output cards.
And because it's a real agent, the sidebar isn't a chat box with opinions — it works. Swap the tag to developer:rust and you can code from the browser: the agent edits files in its sandboxed workspace, tool calls show up as clickable rows with live status, and when it wants to show you something richer than prose — a form, a table, a small interactive UI — it renders an A2UI surface right in the sidebar. We've fixed bugs in Octoweb, from Octoweb, while the failing page sat in the next tab.
The agent tag is editable. The sidebar header shows octoweb:assistant by default, but type any tag your Octomind knows — developer:rust , your own custom agent — and the sidebar reconnects to it on the spot. The browser doesn't care which brain it's talking to.
The browser never calls an AI provider. Every model call is made by Octomind, outside the browser process, with the agent's filesystem access sandboxed to its workspace. No data leaves your machine unless your agent sends it somewhere — and which agent that is, with which provider, is entirely your call.
The session is a two-way street. Since 0.10.0, messages injected by the Octomind runtime — a specialist replying from the tap, a scheduled job finishing, a webhook firing — arrive in the sidebar as their own labeled bubbles, mid-conversation. Your browser becomes the place where your agents report back.
There's also a smaller door: select text on any page, press ⌘⇧E , and an inline edit modal transforms the selection — rewrite, summarize, translate. It runs its own dedicated agent from the same tap — octoweb:editor — with its own prompt history, ⌃R and all.
The browser as a tool: MCP from the inside
This is the direction nobody else builds, and it's our favorite part. Octoweb runs an MCP server inside the browser — localhost:3434/mcp . Any MCP client — Claude Desktop, an Octomind session, your own script — can connect and drive it: twenty-six tools covering navigation, tabs, clicking, typing, scrolling, screenshots, content extraction, console output, and network activity.
You've seen browser automation before. The difference here is that the agent isn't driving a headless Chromium in a datacenter with no cookies and a bot-detection target on its back. It's driving your browser — your sessions, your logins, your tabs — while you watch. And the tool design takes "while you watch" seriously:
Navigation never steals focus. browser_navigate always opens in the background. Exactly one tool — browser_switch_tab — is allowed to change what you're looking at. An agent can research in ten tabs behind your back without ever yanking the page out from under your cursor. It browses beside you, not instead of you.
Clicks that don't lie. browser_click retries until the element is present, stable, and unobstructed — and if something is covering it, the error says what . browser_snapshot returns a map of interactive elements with stable @N refs that pierces iframes, shadow DOM, and listener-only clickables. browser_wait knows about SPA readiness, not just page load.
Eyes included. browser_console_messages and browser_network_requests give the agent the page's console errors and fetch/XHR activity with statuses and timings — the two things you'd open DevTools for, as tools.
Guard rails at the boundary. Text that flows out through snapshots and page content passes a sanitizer — card numbers get redacted before they ever reach a prompt.
Tokens are the budget. Every tool answer is shaped for a model, not for a debugger: browser_snapshot returns a compact element map instead of raw HTML (we deleted the get_html tool on purpose), console and network logs come capped and filterable, and @N refs mean the agent never burns half a context window guessing CSS selectors. An agent's turn is priced in tokens; the browser respects that.
And if your agent runtime is Octomind, wiring it up is one line inside a session:
/mcp add octoweb http://localhost:3434/mcp
The agent you were already talking to can now open docs, read them, fill forms, and check what a page's console is screaming about — mid-conversation, no restart. This is how browser chores leave our to-do lists: the agent reproduces a bug in a background tab, reads the console, and comes back with the failing request; it walks an admin panel we didn't feel like clicking through; it checks that a deploy actually rendered — in our logged-in browser, with our sessions, while we keep reading in the foreground tab.
The browser that learns with you
One more experiment: proactive learning . On an interval — 30 minutes by default — Octoweb wakes a second, fully independent agent, the octoweb:learning tag, which looks at your open tabs and recent history and distills what you're working on into Octomind's memory . Next session, the sidebar assistant already knows the docs you've been circling for three days.
It's on by default — one toggle in Settings turns it off — it's local, the interval is yours to set, and the two agents are fully independent: run either without the other. We think ambient memory is where browser AI actually gets interesting, and we'd rather ship the honest v1 of it than a demo.
Yes, it's also just a good browser
None of this matters if the browsing is worse than what you left. The fundamentals are there, and a few are better than there:
Smart tab hibernation — background tabs freeze under memory pressure, with RAM-aware thresholds. Fifty tabs is a lifestyle, not a leak. And since 0.10.0 there's a cap on open tabs (500 by default, configurable) that quietly closes the least-recently-used ones past it — an airbag for tab hoarders.
Content blocking built in — trackers and ads via WebKit's native content rule lists, no extension needed.
Cold starts stay cold — favicons are cached as data-URIs, so launching the browser makes zero network requests until you do.
Session restore, WebAuthn passkeys, PDF and DOCX viewing, find-in-page on the CSS Custom Highlight API, native fullscreen, proper popup-window handling.
What about extensions? Honest status: the Safari Web Extensions integration (for 1Password, Bitwarden, and friends) is fully implemented in the code — and blocked on Apple's com.apple.developer.web-browser entitlement, an approval only Apple can grant. Until it lands, content scripts don't inject. We'd rather tell you that than pretend extensions are a roadmap item.
Octoweb started in late March as an experiment: could a keyboard browser with an ACP sidebar be usable in a week? By 0.6.0 it had tab hibernation and the first MCP automation tools. 0.7.0 brought persistent chat sessions, passkeys, and slash commands. 0.8.0 was the big one — the actionability harness, the e2e suite that tests the browser through its own MCP server, agent sandboxing, configurable keybindings, A2UI surfaces. 0.9.0, cut in late July, moved to ACP v1.2.0 with account authentication and a long tail of stability work. And today's 0.10.0 is what daily use does to a roadmap: pinned tabs, the open-tab cap, account quota status that refreshes itself, specialist messages landing in the sidebar — plus fixes for a memory leak and orph

[truncated]
