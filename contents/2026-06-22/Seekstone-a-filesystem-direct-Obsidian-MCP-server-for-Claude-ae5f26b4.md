---
source: "https://seekstone.dev/"
hn_url: "https://news.ycombinator.com/item?id=48630401"
title: "Seekstone – a filesystem-direct Obsidian MCP server for Claude"
article_title: "Seekstone — the fastest Obsidian MCP server for Claude"
author: "shaqmughal"
captured_at: "2026-06-22T14:53:18Z"
capture_tool: "hn-digest"
hn_id: 48630401
score: 1
comments: 0
posted_at: "2026-06-22T14:12:34Z"
tags:
  - hacker-news
  - translated
---

# Seekstone – a filesystem-direct Obsidian MCP server for Claude

- HN: [48630401](https://news.ycombinator.com/item?id=48630401)
- Source: [seekstone.dev](https://seekstone.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T14:12:34Z

## Translation

タイトル: Seekstone – クロード用のファイルシステム直接の Obsidian MCP サーバー
記事のタイトル: Seekstone — クロード用の最速の Obsidian MCP サーバー
説明: Seekstone は、Claude 用のファイルシステム直接の Obsidian MCP サーバーです。 REST プラグインよりも最大 575 倍小さいペイロードで、ミリ秒単位でボールトを検索して編集できます。プラグインも Obsidian アプリも必要ありません。

記事本文:
Seekstone — クロード用の最速の Obsidian MCP サーバー
Seekstone ベンチマーク ツールのインストール FAQ 1.2k
シークストーンを入手する
オープンソース · MIT · Obsidian MCP
保管庫全体を検索する
ミリ秒単位で。
Seekstone は、クロード用の最速の Obsidian MCP サーバーです — ファイルシステムダイレクト、
プラグインもアプリも必要ありません。検索は 1.4 ～ 3.2 ミリ秒で返されます。
ペイロードが最大 575 倍小さいため、クロードはメモ ライブラリ全体を読み取ります
コンテキストを燃やさずに。
GitHub でスターを付ける
$ npx -y seenstone init コピー速度 25 ～ 160 倍 プラグイン不要 ネットワーク呼び出しなし macOS · Linux · Windows ベンチマーク済み、未評価 唯一の Obsidian MCP サーバー
公開されている数字付き。
実際のボールト上の 5 つの人気のあるサーバーに対して測定 — 1,955 ノート、それぞれ 20 回実行。ハーネスはオープンソースです。自分のボールトで実行し、すべての数値を確認します。
16 のツール · 1 つの温かいインデックス クロードに必要なものすべて
金庫を動作させるために。
読み取りと書き込み、リンクグラフ、定期的なメモなど、Obsidian MCP サーバーの中で最も広範なツールセットです。 4 つの機能は、他のテスト済みサーバーには搭載されていません。
読み取り · 8 書き込み · 8 検索 全文検索では、ランク付けされた ~200 文字の抜粋が返されます (完全なメモではありません)。
read_note パスごとにメモを読み取ります。単一のセクション、ブロック、または行範囲を返します。
list_notes メモをリストします。オプションでフォルダーのプレフィックスまたはタグでフィルター処理されます。
list_tags ボールト内のすべてのタグ。使用回数またはアルファベット順に並べ替えられます。
outline_note 見出しとブロックの構造 — 目的の内容を読む前の簡単なナビゲーション。
get_backlinks 特定のノートにリンクしているすべてのノートを検索します。
get_links ノートから発信されるすべてのウィキリンクとマークダウン リンクをリストします。
get_periodic_note 任意の日付の毎日/毎週/毎月のメモを読む — Obsidian は閉じられています。
create_note オプションのフロントマターを使用してノートを作成します。親ディレクトリがあなたのために作成されました。
append_note フロントマターに触れずにテキストをノート本文に追加します。

patch_frontmatter YAML をその場で編集します。キーの順序、引用スタイル、コメントは保持されます。
patch_note 前見出しはそのままで、見出しの直後にテキストを挿入します。
replace_in_note 最初に出現したフレーズを予行演習プレビューで置き換えます。
move_note メモを移動または名前変更します。宛先ディレクトリが自動的に作成されます。
append_periodic_note テンプレートから作成して、今日の定期メモに追加します。
delete_note メモを完全に削除します。取り消しは不可能です - あなたが尋ねた場合のみ。
ほとんどのサーバーは、ヒットごとに完全なメモの内容 (モデルが読み取る必要があるメガバイト) を返します。 Seekstone は短いランクの抜粋を返すため、459,000 トークンのコストがかかるクエリのコストは約 800 です。
Obsidian アプリ、ローカル REST API プラグイン、クラウドはありません。 Seekstone はディスクからファイルを直接読み取り、ネットワーク呼び出しを行わず、テレメトリも送信しません。書き込みは、ユーザーが要求した場合にのみ行われます。
マシン上のプレーンな Markdown — 何もロックされません。Frontmatter の編集は設計上バイト同一です。キーの順序、引用スタイル、コメントが保持されており、テスト ハーネスによって証明されています。
最初の検索まで 30 秒。自分の道を選んでください。
2 つの npm 名、1 つのサーバー。 Claude Desktop、Claude Code、Cursor、Windsurf など、あらゆる MCP-over-stdio クライアントで動作します。
1 GitHub Releases から Seekstone.mcpb をダウンロードします。
2 Claude Desktop で開きます — Finder でダブルクリックします。
3 プロンプトが表示されたら、黒曜石の保管庫を選択します。終わり。
.mcpbをダウンロード
ターミナルや Node.js は必要ありません。ボールトを自動検出して検証し、Claude Desktop にパッチを適用します。
コピー --write をパッチに追加するか、--vault を選択して追加します。 1 つのコマンドで Claude Code をエンドツーエンドで構成します。
コピーして claude_desktop_config.json に貼り付けます ([設定] → [開発者])。
コピー シークストーンとしてインストール obsidian-mcp-seekstone ノード ≥ 22 知っておきたい質問と回答。
いいえ。Seekstone は Vault フォルダーをディスクから直接読み取ります。 ○

bsidian は開いていても閉じていても構いません。実行している必要はありません。ローカル REST API プラグインは必要ですか?いいえ、シークストーンはそれを完全にバイパスします。これが最大 575 倍のペイロード削減の原因です。プラグインはまったく必要ありません。どの AI クライアントがサポートされていますか?標準入出力経由でモデル コンテキスト プロトコルを話すクライアント (Claude Desktop、Claude Code、Cursor、Windsurf、Continue など)。私の保管庫で使用しても安全ですか? Seekstone は、書き込みツールを明示的に呼び出した場合にのみファイルを変更します。ネットワーク リクエストは行われず、ボールト パスはサンドボックス化されており、どのツールもそのパスの外で読み書きすることはできません。 Windowsでも動作しますか?はい。 Seekstone は、コミットごとに CI の macOS、Linux、および Windows でテストされます。どれくらいの大きさの金庫まで対応できるのでしょうか？何千もの紙幣が保管されている保管庫に対してプロファイリングが行われています。メモリ内のインデックスは数 MB で、数秒で開始されます。探す時が来ました。
MIT のもとでの無料のオープンソース。クロードにボールトへのミリ秒アクセスを許可し、コンテキスト ウィンドウを元に戻します。
ソースを見る
最新情報を入手 新しいリリースがいつ出荷されるかを誰よりも早く知りましょう。スパムは一切ありません。
会社 購読すると、リリースの更新を受け取ることに同意したことになります。当社のプライバシー ポリシーをご覧ください。いつでも購読を解除してください。
Seekstone クロード用の最速の Obsidian MCP サーバー。

## Original Extract

Seekstone is a filesystem-direct Obsidian MCP server for Claude. Search and edit your vault in milliseconds, with ~575× smaller payloads than the REST plugin. No plugins, no Obsidian app required.

Seekstone — the fastest Obsidian MCP server for Claude
Seekstone Benchmarks Tools Install FAQ 1.2k
Get Seekstone
Open source · MIT · Obsidian MCP
Search your entire vault
in milliseconds.
Seekstone is the fastest Obsidian MCP server for Claude — filesystem-direct,
no plugins, no app required. Searches return in 1.4–3.2 ms with
~575× smaller payloads, so Claude reads your whole note library
without burning its context.
Star on GitHub
$ npx -y seekstone init Copy 25–160× faster No plugins required No network calls macOS · Linux · Windows Benchmarked, not claimed The only Obsidian MCP server
with published numbers.
Measured against 5 popular servers on a real vault — 1,955 notes, 20 runs each. The harness is open source; run it on your own vault and verify every figure.
16 tools · one warm index Everything Claude needs
to work your vault.
Read and write, link-graph and periodic notes — the broadest toolset of any Obsidian MCP server. Four capabilities no other tested server even ships.
Read · 8 Write · 8 search Full-text search returning ranked ~200-char excerpts — not full notes.
read_note Read a note by path. Return a single section, block, or line range.
list_notes List notes, optionally filtered by folder prefix or tag.
list_tags Every tag in the vault, sorted by usage count or alphabetically.
outline_note Heading & block structure — cheap navigation before a targeted read.
get_backlinks Find all notes that link to a given note.
get_links List all outgoing wikilinks and markdown links from a note.
get_periodic_note Read any date's daily / weekly / monthly note — Obsidian closed.
create_note Create a note with optional frontmatter; parent dirs made for you.
append_note Append text to a note body without ever touching frontmatter.
patch_frontmatter Edit YAML in place — key order, quote style and comments preserved.
patch_note Insert text immediately after a heading, frontmatter untouched.
replace_in_note Replace the first occurrence of a phrase, with a dry-run preview.
move_note Move or rename a note; destination directories created automatically.
append_periodic_note Append to today's periodic note, creating it from a template.
delete_note Permanently delete a note. Irreversible — only when you ask.
Most servers return full note content for every hit — megabytes your model has to read. Seekstone returns short ranked excerpts, so a query that cost 459,000 tokens costs about 800.
No Obsidian app, no Local REST API plugin, no cloud. Seekstone reads files straight from disk, makes zero network calls, and sends no telemetry. Writes only ever happen when you ask.
Plain Markdown on your machine — nothing to lock you in. Frontmatter edits are byte-identical by design: key order, quote style and comments preserved, proven by the test harness.
30 seconds to first search Pick your way in.
Two npm names, one server. Works with Claude Desktop, Claude Code, Cursor, Windsurf — any MCP-over-stdio client.
1 Download seekstone.mcpb from GitHub Releases.
2 Open it with Claude Desktop — double-click in Finder.
3 Pick your Obsidian vault when prompted. Done.
Download .mcpb
No terminal · no Node.js required Auto-detects your vault, validates it, and patches Claude Desktop for you.
Copy Add --write to patch in place, or --vault to choose. One command configures Claude Code end-to-end.
Copy Paste into claude_desktop_config.json (Settings → Developer).
Copy Installs as seekstone obsidian-mcp-seekstone Node ≥ 22 Good to know Questions, answered.
No. Seekstone reads the vault folder directly from disk. Obsidian can be open or closed — it never has to be running. Do I need the Local REST API plugin? No. Seekstone bypasses it entirely — that's the source of the ~575× payload reduction. No plugins are required at all. Which AI clients does it support? Any client that speaks the Model Context Protocol over stdio — Claude Desktop, Claude Code, Cursor, Windsurf, Continue, and more. Is it safe to use on my vault? Seekstone only modifies files when you explicitly call a write tool. It makes no network requests, and the vault path is sandboxed — no tool can read or write outside it. Does it work on Windows? Yes. Seekstone is tested on macOS, Linux, and Windows in CI on every commit. How big a vault can it handle? It has been profiled against vaults with thousands of notes. The in-memory index is a few MB and starts in a few seconds. It's your time to seek.
Free and open source under MIT. Give Claude millisecond access to your vault — and your context window back.
View source
Stay in the loop Be the first to know when a new release ships. No spam, ever.
Company By subscribing you agree to receive release updates. See our privacy policy . Unsubscribe anytime.
Seekstone The fastest Obsidian MCP server for Claude.
