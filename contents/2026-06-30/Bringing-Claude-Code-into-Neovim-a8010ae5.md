---
source: "https://inacioklassmann.com/posts/claude-chat-nvim/"
hn_url: "https://news.ycombinator.com/item?id=48739649"
title: "Bringing Claude Code into Neovim"
article_title: "Bringing Claude Code Into Neovim | Inácio Klassmann"
author: "samsgro"
captured_at: "2026-06-30T22:27:01Z"
capture_tool: "hn-digest"
hn_id: 48739649
score: 1
comments: 0
posted_at: "2026-06-30T21:46:02Z"
tags:
  - hacker-news
  - translated
---

# Bringing Claude Code into Neovim

- HN: [48739649](https://news.ycombinator.com/item?id=48739649)
- Source: [inacioklassmann.com](https://inacioklassmann.com/posts/claude-chat-nvim/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T21:46:02Z

## Translation

タイトル: Neovim へのクロード コードの導入
記事のタイトル: Neovim にクロード コードを導入する |イナシオ・クラスマン
説明: 私は 2 つの場所に住んでいます: Claude Code と Neovim を実行している端末です。何ヶ月もの間、私は彼らの間を行き来しました。 Neovim で編集し、Alt キーを押しながら Tab キーを押してターミナルに移動し、ファイル パスを貼り付け、クロードに何かを尋ね、Alt キーを押しながら Tab キーを押して戻り、手動で変更を適用します。摩擦は小さいですが一定であり、一定の摩擦は私たちが
[切り捨てられた]

記事本文:
Neovim への Claude コードの導入 | Inácio Klassmann Inácio Klassmann 個人ブログ ソース コード ホーム カテゴリー タグ アーカイブ ABOUT ホーム Neovim へのクロード コードの導入 投稿 キャンセル Neovim へのクロード コードの導入
私は 2 つの場所に住んでいます。Claude Code と Neovim を実行している端末です。何ヶ月もの間、私は彼らの間を行き来しました。 Neovim で編集し、Alt キーを押しながら Tab キーを押してターミナルに移動し、ファイル パスを貼り付け、クロードに何かを尋ね、Alt キーを押しながら Tab キーを押して戻り、手動で変更を適用します。摩擦は小さいですが一定であり、絶え間ない摩擦は一日の中で気分を損なうようなものです。
VS Code と JetBrains の群衆は、これの素晴らしいバージョンをすでに持っていました。クロードはエディター内に住み、開いているファイル、カーソル、選択内容、診断を確認します。ネオビムはそうではなかった。それで私はそれを建てました。これは claude-chat.nvim という名前で、99.8% Lua です。
実際には次のようになります。
Neovim サイドバーにある Claude Code TUI は、私が編集しているファイルを認識しています。
私の最初の本能は、すべてのことと同じように、重いものでした。 Anthropic API をラップします。自分の会話状態を管理します。ストリーミングを再実装します。許可プロンプトを再構築します。 Neovim バッファーに小さなチャット UI を構築し、そこにトークンをパイプします。
これは非常に大きな表面積であり、その隅々までクロード・コードがすでに行っていることであり、私よりも優れています。ストリーミングを再実装したり、許可ダイアログを再実装したりしていることに気づいた瞬間、私は間違った方向に進んでしまいます。ツールはそれらをすでに所有しています。
そこで私はそのアイデアを捨てて、もっと小さな質問をしました。「何もラップしなかったらどうなるでしょうか?」
Claude Code はターミナル UI です。 Neovim にはターミナル バッファがあります。プラグイン全体の核心は、Neovim ターミナル バッファ内のサイドバー内で実際の claude CLI を子プロセスとして起動するだけです。
その 1 つの決断はすぐに報われます。だってそれは

実際の TUI では、ストリーミング応答、対話型の許可プロンプト、セッションの再開など、すべてがターミナルで claude を実行しているのとまったく同じように動作します。これらの機能のいずれについてもコードを作成しませんでした。私がそれらを所有しようとするのをやめた瞬間に、それらは無料で提供されました。
これは、私が常に学び続けている教訓と同じです。ツールの最良のバージョンは、通常、既存の部分にその仕事をさせるため、最も機能しないバージョンです。私の仕事は再構築ではなく接続でした。
興味深い部分: Neovim を IDE のように見せること
サイドバーにターミナルがあると便利ですが、それは本当の価値ではありません。本当の利点は編集者の認識です。クロードは、私がどのファイルを見ているのか、カーソルがどこにあるのか、何を選択したのか、LSP が何について文句を言っているのかを知ってくれます。これが、VS Code エクスペリエンスを、単に同じ場所にあるのではなく、魔法のように感じさせる理由です。
クロード・コードはすでに編集者と話す方法を知っていることが判明しました。 VS Code および JetBrains 拡張機能は WebSocket 経由で MCP を話し、CLI はその接続にダイヤルバックする方法を認識します。プロトコルを発明する必要はありませんでした。すでに存在するプロトコルを話さなければなりませんでした。
したがって、プラグインは 2 つの MCP サーバーを起動し、どちらもセッションごとの認証トークンでローカルホストにバインドされます。
Claude の IDE 拡張機能が使用するのと同じプロトコルを話す WebSocket MCP サーバーは、エディターの状態 (ファイルを開く、カーソル、選択、診断など) を同期します。 open_file 、 current_file などのインプロセス ツールをホストする HTTP MCP サーバー、および Neovim の LSP へのブリッジ (定義、参照、ホバー、シンボル検索)。ディスカバリーハンドシェイクは、静かに満足感を与えるものです。プラグインは、ロック ファイルを ~/.claude/ide/<port>.lock に書き込み、CLAUDE_CODE_SSE_PORT を設定して CLI を起動します。 CLI はロック ファイルを見つけて再度接続し、それ以降、クロードは VS Code を見るのと同じようにエディターを見ることができるようになります。フーです

自動 — 手動で実行する必要はありません。
その LSP ブリッジが私が最も気に入っている部分です。クロードは生のテキストを取得するだけではありません。それは私と同じ意味論的な見解を取得します。関数について質問すると、文字列から推測するのではなく、プロジェクト全体の参照と定義をたどることができます。
<leader>ai でサイドバーを切り替え、 <leader>af で現在のファイルを会話にプッシュし、通常の <C-h/j/k/l> でエディタとチャットの間を移動します。クロードが編集を提案すると、ファイルに何かが触れる前に、編集がスクラッチ バッファーに並べて表示されます。他の IDE と同様に、編集を承認します。
インストールは退屈な作業ですが、それが私は気に入っています。
1
2
3
4
5
6
{
"codegik/claude-chat.nvim" ,
cmd = { "クロードチャット" , "クロードチャットリセット" , "クロードチャットファイル" },
キー = { { "<リーダー>ai" , "<cmd>ClaudeChat<cr>" } },
config = function () require ( "claude-chat" )。セットアップ()終了、
}
Neovim 0.10+ と PATH 上の claude CLI が必要です。プラグインの API キー、2 回目のログイン、別のアカウントは必要ありません。当然のことながら、クロード コードがすでに設定したものはすべて再利用されます。それが要点でした。
このプラグインには、API クライアント、状態ストア、ストリーミング レンダラー、権限システムなど、広大なアプリケーションであるバージョンがあります。そのバージョンは存在する必要がなかったため、存在しません。 Claude Code が会話と認証を所有します。 MCP プロトコルはエディターの同期を所有します。 Neovim の端末がレンダリングを所有し、その LSP がセマンティクスを所有します。私が書いたのは、すでに存在するもの間の配線であり、さらに Neovim を IDE として紹介するのに十分なプロトコルの話でした。
優れたエンジニアリングは引き算であることが判明し続けています。コードは GitHub: codegik/claude-chat.nvim にあります。
Neovim に Claude コードを取り込む 私の Cla を監視するための小さな Waybar モジュール

ude コードの制限 Scala + Springboot の相互運用性 ショート ポーリング、ロング ポーリング、および Websocket コードの曖昧さの回避 トレンドのタグ アジャイル リーダーシップ スプリングブート アーキテクチャの学習 クロード コードのフィードバック グッド プラクティス Linux スキル 目次 参考資料
私は自分のマシンで Omarchy を実行しており、最近の多くの人たちと同じように、Claude Code を一日中開いたままにしています。それは素晴らしいことです。何かの途中でレート制限に達し、その後初めてその制限を思い出すまでは...
私のクロードコードの制限を監視するための小さな Waybar モジュール
© 2026 イナシオ・クラスマン .一部の権利は留保されています。
新しいバージョンのコンテンツが利用可能です。

## Original Extract

I live in two places: a terminal running Claude Code and Neovim. For months I bounced between them. Edit in Neovim, alt-tab to the terminal, paste a file path, ask Claude something, alt-tab back, apply the change by hand. The friction was small but constant, and constant friction is the kind that we
[truncated]

Bringing Claude Code Into Neovim | Inácio Klassmann Inácio Klassmann Personal Blog Source Code HOME CATEGORIES TAGS ARCHIVES ABOUT Home Bringing Claude Code Into Neovim Post Cancel Bringing Claude Code Into Neovim
I live in two places: a terminal running Claude Code and Neovim. For months I bounced between them. Edit in Neovim, alt-tab to the terminal, paste a file path, ask Claude something, alt-tab back, apply the change by hand. The friction was small but constant, and constant friction is the kind that wears a groove in your day.
The VS Code and JetBrains crowd already had the nice version of this: Claude living inside the editor, seeing the open file, the cursor, the selection, the diagnostics. Neovim didn’t. So I built it. It’s called claude-chat.nvim , and it’s 99.8% Lua.
Here’s what it looks like in practice:
The Claude Code TUI living in a Neovim sidebar, aware of the file I’m editing.
My first instinct — same as with everything — was the heavy one. Wrap the Anthropic API. Manage my own conversation state. Re-implement streaming. Rebuild permission prompts. Build a little chat UI in a Neovim buffer and pipe tokens into it.
That’s a lot of surface area, and every inch of it is something Claude Code already does and does better than I would. The moment I find myself re-implementing streaming or re-implementing a permission dialog, I’ve taken a wrong turn. The tool already owns those.
So I threw the idea out and asked a smaller question: what if I didn’t wrap anything at all?
Claude Code is a terminal UI. Neovim has terminal buffers. The whole plugin, at its core, is just: launch the actual claude CLI as a child process inside a Neovim terminal buffer, in a sidebar.
That one decision pays for itself immediately. Because it is the real TUI, everything behaves exactly like running claude in a terminal — streaming replies, interactive permission prompts, session resume, all of it. I wrote zero code for any of those features. They came for free the instant I stopped trying to own them.
This is the same lesson I keep relearning: the best version of a tool is usually the one that does the least , because it lets the existing pieces do their jobs. My job was connection, not reconstruction.
The interesting part: making Neovim look like an IDE
A terminal in a sidebar is convenient, but it isn’t the real prize. The real prize is editor awareness — Claude knowing what file I’m looking at, where my cursor is, what I’ve selected, what the LSP is complaining about. That’s what makes the VS Code experience feel magical instead of just colocated.
It turns out Claude Code already knows how to talk to an editor. The VS Code and JetBrains extensions speak MCP over a WebSocket, and the CLI knows how to dial back into that connection. I didn’t have to invent a protocol — I had to speak the one that already exists .
So the plugin stands up two MCP servers, both bound to localhost with a per-session auth token:
A WebSocket MCP server speaking the same protocol Claude’s IDE extensions use, to sync editor state — open files, cursor, selection, diagnostics. An HTTP MCP server hosting in-process tools like open_file , current_file , and bridges into Neovim’s LSP — definitions, references, hover, symbol search. The discovery handshake is the quietly satisfying bit. The plugin writes a lock file to ~/.claude/ide/<port>.lock and launches the CLI with CLAUDE_CODE_SSE_PORT set. The CLI finds the lock file, connects back, and from then on Claude can see my editor the same way it sees VS Code. It’s fully automatic — nothing to run by hand.
That LSP bridge is the part I’m most fond of. Claude doesn’t just get raw text; it gets the same semantic view I have. Ask it about a function and it can follow references and definitions through the project instead of guessing from strings.
You toggle the sidebar with <leader>ai , push the current file into the conversation with <leader>af , and move between editor and chat with the usual <C-h/j/k/l> . When Claude proposes an edit, it shows up as a side-by-side diff in a scratch buffer before anything touches your file — you approve it like you would in any IDE.
Installation is the boring kind, which is how I like it:
1
2
3
4
5
6
{
"codegik/claude-chat.nvim" ,
cmd = { "ClaudeChat" , "ClaudeChatReset" , "ClaudeChatFile" },
keys = { { "<leader>ai" , "<cmd>ClaudeChat<cr>" } },
config = function () require ( "claude-chat" ). setup () end ,
}
You need Neovim 0.10+ and the claude CLI on your PATH . No API key for the plugin, no second login, no separate account — it reuses whatever Claude Code already set up, because of course it does. That was the whole point.
There’s a version of this plugin that’s a sprawling application: an API client, a state store, a streaming renderer, a permission system. That version doesn’t exist, because none of it needed to. Claude Code owns the conversation and the auth. The MCP protocol owns the editor sync. Neovim’s terminal owns the rendering and its LSP owns the semantics. What I wrote was the wiring between things that were already there — plus enough protocol-speaking to make Neovim introduce itself as an IDE.
Good engineering keeps turning out to be subtraction. The code is on GitHub: codegik/claude-chat.nvim .
Bringing Claude Code Into Neovim A Tiny Waybar Module to Watch My Claude Code Limits Scala + Springboot Interoperability Short polling, Long polling and Websockets Avoiding Code Obscurity Trending Tags agile leadership learning springboot architecture claude-code feedback good practice linux skills Contents Further Reading
I run Omarchy on my machine, and like a lot of people these days I keep Claude Code open all day. It’s great — until you hit a rate limit in the middle of something and only then remember those lim...
A Tiny Waybar Module to Watch My Claude Code Limits
© 2026 Inácio Klassmann . Some rights reserved.
A new version of content is available.
