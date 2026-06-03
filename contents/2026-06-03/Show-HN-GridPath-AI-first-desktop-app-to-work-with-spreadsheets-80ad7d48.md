---
source: "https://gridpath.dev/"
hn_url: "https://news.ycombinator.com/item?id=48369760"
title: "Show HN: GridPath – AI first desktop app to work with spreadsheets"
article_title: "GridPath — Best way to build spreadsheets with AI"
author: "pixelmash13"
captured_at: "2026-06-03T00:47:00Z"
capture_tool: "hn-digest"
hn_id: 48369760
score: 2
comments: 0
posted_at: "2026-06-02T13:06:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GridPath – AI first desktop app to work with spreadsheets

- HN: [48369760](https://news.ycombinator.com/item?id=48369760)
- Source: [gridpath.dev](https://gridpath.dev/)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:06:22Z

## Translation

タイトル: Show HN: GridPath – スプレッドシートを操作する AI 初のデスクトップ アプリ
記事のタイトル: GridPath — AI を使用してスプレッドシートを構築する最良の方法
説明: GridPath は、Excel 用のデスクトップ AI エージェントです。並列ワークフロー。可逆的な変化。電光石火のスピード。 Claude および OpenAI サブスクリプションで動作します。
HN テキスト: こんにちは HN — スプレッドシートの Cursor のように実行できるものが欲しかったので、GridPath を構築しました。スプレッドシートの作業と複数のモデル プロバイダーへの接続用にカスタマイズされたツール セットを使用して、並列ワークロード、可逆的な変更を高速に実行できます。これは Tauri/Rust デスクトップ アプリで、無料で使用できます。Claude トークンを接続するか、OpenAI サブスクリプションでサインインするだけです。理解するのに興味深い点がいくつかありました。Univer の無料ライセンスだけで Excel に非常に近づくことができます。最も重要な欠落要素はグラフ (既存のグラフは保持されますが、ビューには表示されません) とピボット テーブルです。 Rust + モデルで利用可能なツール (基本的には Web フェッチ + ユーザーとの対話 + Excel 操作のみ) の削減により、速度とトークンの書き込みが大幅に向上します。 SpaceX モデルは 2 分 30 秒で構築されましたが、Excel プラグインの場合は約 5 分でした。別の実行では、Claude CLI は同じモデルを構築しましたが、10 倍以上のトークンを焼きました (主にさらに多くの Web フェッチとモデルの反復)。次に、複数シートの作業におけるエージェントの効率の向上と、より長いセッションにおけるコンテキストの圧縮に取り組みます。リポジトリ: https://github.com/pixelsmasher13/gridpath
サイト: https://gridpath.dev
フィードバックは大歓迎です!

記事本文:
グリッドパス
能力
モデル
よくある質問
ダウンロード
構築するための最良の方法
AI を使用したスプレッドシート。
並列ワークフロー。可逆的な変化。電光石火のスピード。
Claude または OpenAI サブスクリプションでは無料です。
カリブリ
10
B
私は
U
S
あ
$
%
.00
あ
B
C
D
E
1
TESLA — 損益計算書モデル
2
2024年度
2025年度
2026年度E
2027年度E
3
自動車販売
72,480
65,821
=D3*(1+E10)
=E3*(1+F10)
4
その他の収益
25,210
29,006
=D4*(1+E11)
=E4*(1+F11)
5
総収益
97,690
94,827
=D3+D4
=E3+E4
6
収益原価
80,240
77,733
85,124
93,123
7
売上総利益
17,450
17,094
=D5-D6
=E5-E6
8
粗利率 %
17.9%
18.0%
18.5%
19.0%
9
営業利益
7,076
4,355
=D7-D11-D12
=E7-E11-E12
10
当期純利益
7,153
4,193
=D9+D13-D14
=E9+E13-E14
執筆中 — 361 件の編集がレビュー保留中
イン 30,133 / アウト 8,215
会話
25 年度および 26 ～ 30 年度の予測までの SEC 実績を使用して、5 年間のテスラ損益計算書を作成します。
SEC からテスラの最新の財務情報を入手して、完全なウォーターフォールを構築します。
⌖ fetch_web · sec.gov からの 2 つの URL
▦ set_range · 58 × 9 → Sheet1!A1
∑ set_format · 64 ops (1 回の一括呼び出し)
提案された変更
361 件の編集
✓ 受け入れる
× 拒否する
選択時：D3:D8（6マス）×
税率の 3 年間の感応度テーブルを追加します…
↑
Cursor のように構築されています。エクセルのような形をしています。
GridPath の使用感が他のものと異なる 6 つの点
あなたが試した他の「スプレッドシート用 AI」ツールすべて。
複数のワークブックを並べて実行します。各タブには独自のエージェントが割り当てられます。
差分優先編集。承認、拒否、⌘Z で元に戻します。あなたがそう言うまでは何も救いません。
Electron ではなく、ネイティブの Rust コア。即時コールド スタート、ターンあたり最大 75% のプロンプト キャッシュ レート。
グラフ、形式、名前付き範囲、結合 - すべてが保存時に保持されます。
貼り付けられた数値ではなく、 =SUM 、 =VLOOKUP を書き込みます。モデルはライブのままです。
.xlsx はディスク上に残ります。プロンプトのみが LLM に送信されます。
Claude または ChatGPT サブユーザーでサインインします

あなたがすでに持っている資格、
または API キーをプラグインします。私たちからは何も余分なものはありません。
クロード
OAuth または Anthropic API キー経由の Pro & Max
サポートされています
OpenAI
OAuth 経由の ChatGPT Plus および Pro (コーデックス)
サポートされています
[設定] からいつでもプロバイダーを切り替えます。さらに登場します。
.xlsx ファイル自体はディスク上に残ります - GridPath はありません
クラウド、アップロードなし、中間にサードパーティサーバーなし。どういうことですか
サインインした LLM プロバイダーにフローします (Anthropic for
Claude、OpenAI for ChatGPT) はプロンプト、シート構造、
エージェントが実行中に読み書きするセル。同じ
そのプロバイダーの製品を使用する場合と同じ形状と同じプライバシー条件
直接的に。長いバージョンについては、「プライバシー」を参照してください。
GridPath は Claude または ChatGPT サブスクリプションに接続します —
LLM プロバイダーには、既に支払った金額をそのまま支払います。
私たちにとっては余計な事。
はい、ただし、正直な注意点が 1 つあります。オリジナルのワークブックを保管します
ExcelJSを使用してメモリ内で
保存時にエージェントが実際に変更するセルにのみパッチを適用します。
それは他のすべてを意味します - チャート、条件付き
書式設定ルール、名前付き範囲、データ検証、結合されたセル、
図面、フリーズペイン、ピボットテーブル、テーマ - フロー
往復を無傷で通過します。後で Excel でファイルを開きます
見た目は以前とまったく同じで、さらに新しい編集も加えられています。
注意: GridPath のアプリ内グリッド (上に構築されています)
ユニバー）
グラフ、条件付き書式設定を視覚的にレンダリングしません
編集中に色付け、描画、ピボット テーブルを行うことができます。彼らは
ファイル内では完全にそのままです - 内部の画面上にペイントされていないだけです
アプリ。エージェントは引き続き結合されたセルを尊重し、ペインを固定し、
ネイティブに名前付き範囲を指定します。ワークフローでグラフを表示する必要がある場合
更新がライブの場合は、Excel でそのパスを実行します。構造的/公式の場合
GridPath の方が高速です。
3 つの方法。 (1) 提案されたバッチをその前に拒否する
コミット — ファイルには何も触れません。 (2) ⌘Zで元に戻します

セッション内で変更を受け入れました。 (3) 保存後、
ワンステップ用に前の状態の .bak を保持します。
再開後も回復。
Excel Copilot は、Excel 内のシングルターン アシスタントです。 GridPath の実行
マルチターンエージェントループ - Web データをフェッチし、58 行を書き込み、
64 の書式設定操作を適用し、独自の数式エラーを修正 (10 ～ 15 回)
ツールは、モデルが完全に完了するまで、プロンプトごとに呼び出します。さらにあなたも
基礎となる LLM を選択し、編集が反映される前にすべての編集を確認します。
今日は.xlsx。 .xls 、 .csv 、および
API 経由の Google スプレッドシートはロードマップにあります。
はい。クロスシート参照 ( =Sheet2!A1 ) はネイティブに機能します。
シートの名前を変更すると、他のシートの参照も自動更新されます。
外部ワークブック リンク ( =[Other.xlsx]Sheet1!A1 ) は存続します
キャッシュされた値として扱われますが、ライブで再計算されません。Google スプレッドシートの制限と同じです。
macOS (Apple Silicon、macOS 12+) および Windows (10/11)。 Linux がオンになっています
ロードマップ。リクエストに応じてセキュリティレビューに利用できるソース。
Claude または ChatGPT サブスクリプションでサインインし、60 秒以内に最初のワークブックを編集します。
Apple Silicon · macOS 12+ · Windows 10/11
グリッドパス
AI を使用してスプレッドシートを構築する最良の方法。

## Original Extract

GridPath is a desktop AI agent for Excel. Parallel workflows. Reversible changes. Lightning speed. Works with your Claude and OpenAI subscription.

Hi HN — I built GridPath as I wanted something that runs like Cursor for spreadsheets - parallel workloads, reversible changes, fast, with a tool set customized for spreadsheet work and connection to multiple model providers. It's a Tauri/Rust desktop app, free to use, just connect your Claude token or sign in with OpenAI subscription. A few things that were interesting to figure out: you can get very close to Excel with just Univer free license, the most material missing elements are charts (existing charts kept but not displaying in the view) and pivot tables. Rust + cutting down on tools available to the model (basically just web fetch + talk to user + excel operations) improves the speed and token burn materially. SpaceX model built in 2 min 30s, compared to ~5 min for Excel plug in, for another run Claude CLI built the same model but burned 10x+ tokens (primarily many more web fetches and model iterations). Working next on improving the agent efficiency for multi-sheet work + context compacting for longer sessions. Repo: https://github.com/pixelsmasher13/gridpath
Site: https://gridpath.dev
Feedback welcome!

GridPath
Capabilities
Models
FAQ
Download
Best way to build
spreadsheets with AI.
Parallel workflows. Reversible changes. Lightning speed.
Free with your Claude or OpenAI subscription.
Calibri
10
B
I
U
S
A
$
%
.00
A
B
C
D
E
1
TESLA — INCOME STATEMENT MODEL
2
FY 2024
FY 2025
FY 2026E
FY 2027E
3
Automotive sales
72,480
65,821
=D3*(1+E10)
=E3*(1+F10)
4
Other revenue
25,210
29,006
=D4*(1+E11)
=E4*(1+F11)
5
Total revenue
97,690
94,827
=D3+D4
=E3+E4
6
Cost of revenue
80,240
77,733
85,124
93,123
7
Gross profit
17,450
17,094
=D5-D6
=E5-E6
8
Gross margin %
17.9%
18.0%
18.5%
19.0%
9
Operating income
7,076
4,355
=D7-D11-D12
=E7-E11-E12
10
Net income
7,153
4,193
=D9+D13-D14
=E9+E13-E14
Writing — 361 edits pending review
in 30,133 / out 8,215
Conversation
Build a 5-year Tesla income statement with SEC actuals through FY25 and FY26–30 projections.
I'll fetch Tesla's latest financials from SEC, then build the full waterfall.
⌖ fetch_web · 2 URLs from sec.gov
▦ set_range · 58 × 9 → Sheet1!A1
∑ set_format · 64 ops (one bulk call)
Proposed changes
361 edits
✓ Accept
× Reject
Selected: D3:D8 (6 cells) ×
Add a 3-year sensitivity table for tax rate…
↑
Built like Cursor. Shaped like Excel.
Six things that make working with GridPath feel different from
every other "AI for spreadsheets" tool you've tried.
Run multiple workbooks side by side. Each tab gets its own agent.
Diff-first edits. Accept, reject, ⌘Z to undo. Nothing saves until you say so.
Native Rust core, not Electron. Instant cold start, ~75% prompt cache rate per turn.
Charts, formats, named ranges, merges — all preserved on save.
Writes =SUM , =VLOOKUP , not pasted numbers. Models stay live.
Your .xlsx stays on disk. Only the prompt goes to the LLM.
Sign in with the Claude or ChatGPT subscription you already have,
or plug in an API key. Nothing extra from us.
Claude
Pro & Max via OAuth, or Anthropic API key
Supported
OpenAI
ChatGPT Plus & Pro via OAuth (Codex)
Supported
Switch providers any time from Settings. More coming.
The .xlsx file itself stays on disk — no GridPath
cloud, no upload, no third-party server in the middle. What does
flow to the LLM provider you've signed in with (Anthropic for
Claude, OpenAI for ChatGPT) is your prompt, sheet structure, and
whatever cells the agent reads or writes during the run. Same
shape and same privacy terms as using that provider's product
directly. See Privacy for the long version.
GridPath connects to your Claude or ChatGPT subscription —
you pay the LLM provider whatever you already pay them, and nothing
extra to us.
Yes — but with one honest caveat. We keep the original workbook
in memory using ExcelJS
and only patch the cells the agent actually changes on save.
That means everything else — charts, conditional
formatting rules, named ranges, data validation, merged cells,
drawings, freeze panes, pivot tables, themes — flows
through the round-trip untouched. Open the file in Excel after
and it looks exactly like it did before, plus your new edits.
The caveat: GridPath's in-app grid (built on
Univer )
doesn't visually render charts, conditional-formatting
coloring, drawings, or pivot tables while you're editing. They're
fully intact in the file — just not painted on screen inside our
app. The agent still respects merged cells, freeze panes, and
named ranges natively. If your workflow needs to see chart
updates live, do that pass in Excel; if it's structural / formula
work, GridPath is faster.
Three ways. (1) Reject the proposed batch before it
commits — nothing touches the file. (2) ⌘Z to undo
accepted changes within the session. (3) After save,
we keep a .bak with the previous state for one-step
recovery, even after reopening.
Excel Copilot is a single-turn assistant inside Excel. GridPath runs
a multi-turn agent loop — fetch web data, write 58 rows,
apply 64 formatting ops, fix its own formula errors — across 10–15
tool calls per prompt, until the model is genuinely done. Plus you
pick the underlying LLM and review every edit before it lands.
.xlsx today. .xls , .csv , and
Google Sheets via API are on the roadmap.
Yes. Cross-sheet references ( =Sheet2!A1 ) work natively,
and renaming a sheet auto-updates references in other sheets.
External workbook links ( =[Other.xlsx]Sheet1!A1 ) survive
as cached values but don't recompute live — same limit Google Sheets has.
macOS (Apple Silicon, macOS 12+) and Windows (10/11). Linux is on
the roadmap. Source available for security review on request.
Sign in with your Claude or ChatGPT subscription and edit your first workbook in under 60 seconds.
Apple Silicon · macOS 12+ · Windows 10/11
GridPath
Best way to build spreadsheets with AI.
