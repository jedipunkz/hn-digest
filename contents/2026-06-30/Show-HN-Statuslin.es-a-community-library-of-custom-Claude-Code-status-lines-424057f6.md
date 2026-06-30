---
source: "https://statuslin.es"
hn_url: "https://news.ycombinator.com/item?id=48732336"
title: "Show HN: Statuslin.es – a community library of custom Claude Code status lines"
article_title: "statuslin.es"
author: "nastynate"
captured_at: "2026-06-30T13:34:10Z"
capture_tool: "hn-digest"
hn_id: 48732336
score: 1
comments: 0
posted_at: "2026-06-30T13:15:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Statuslin.es – a community library of custom Claude Code status lines

- HN: [48732336](https://news.ycombinator.com/item?id=48732336)
- Source: [statuslin.es](https://statuslin.es)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T13:15:39Z

## Translation

タイトル: HN を表示: Statuslin.es – カスタム クロード コード ステータス ラインのコミュニティ ライブラリ
記事のタイトル: statuslin.es
説明: クロード コード ステータス ライン カタログ
HN テキスト: 少し前、私はクロード コードで使用する素敵なステータス ラインを探していましたが、人々が自分の作ったものを披露するためのコミュニティ ライブラリ タイプの Web サイトが存在しないことに驚きました。そこで、ちょっと楽しんで作ってみることにしました。すべてのステータス行は手動で確認され、ライブ コンテナーでレンダリングされて出力がキャプチャされます。楽しむ！ソースコード: https://github.com/NathanAB/statuslin.es

記事本文:
statuslin.es ステータスリン 。 es GitHub statuslin でサインインします。エス▒
ステータス ラインは、Claude Code ターミナルの下部にあるバーです。独自のカスタム構成を他の人のために送信したり、気に入ったものを見つけて自分用に使用したりできます。
ステータス行を送信する Dreambase パネル
python ╭─ コンテキスト ─────────────────────╮
│ ████▊░░░░░░░░░░░░░░░░░ 200K の 22% · GOOD │
§─ 統計 ─────────────────────┤
│ ↓ 44.0K ↑ 1.4K │ $0.41 │ ⏱ 10 分 12 秒 │ +128 / -34 │
§─ REPO ──── ────── ──────────────┤
│ ◆ Opus 4.8 · main · 17:27 │
╰───────────────────────╯
自動サイズ調整の境界線を備えたボックス化された複数行パネル: テラコッタの Claude ブランド パレット内の、グラデーション バーとステータス ワードを含むコンテキスト セクション、統計セクション (トークン、コスト、期間、変更された行)、およびリポジトリ セクション (モデル、ファイル数を含む git ブランチ、クロック、vim モード)。ソース: github.com/kyleledbetter/claudecode-statusline
bash  アプリ   メイン   Opus 4.8  ━━ ━ ━ ━ ━ ━ ━ ━ ━ 22%   128  34    
オタク フォントの区切り文字と Dracula パレットを備えた Powerline スタイルのステータス ライン - 色付きのディレクトリ、git ブランチ、モデルごとのアイコン、および矢印の尾が付いた色褪せたコンテキスト プログレス バー。出典: github.com/LindseyB/claude-statuslin

e
bash Opus 4.8 |メイン | ██░░░░░░░░ 22.0% | 44k / 200k | 02:07 26% |水 7%
単一行のステータス: モデル、git ブランチ、トークン数を含むブロック コンテキスト バーン バー (緑→オレンジ→赤)、およびリセット時間付きの 5 時間および週ごとのレート制限パーセンテージ。出典: github.com/bitcoin21ideas/claude-statusline
bash Opus 4.8 [高] ~/app main
17:27:05 | ⛁ ██░░░░░░░░ 22% | 5 時間 ●◔○○○ 26% ↻2 時間 7 分 | 7日 ◔○○○○○○ 7% ↻2日1時間 │ $0.41 ⏱ 1000万
ブロック コンテキスト バーと 5 時間および週ごとの制限の 4 分の 1 ステップのドット メーター (◔◑◕●) を備えた 2 行のステータス — 塗りつぶしは使用状況を示し、色はペースが進んでいるか遅れているかを示します — 加えて、モデル、労力、Git、コストも表示します。 jqは必要ありません。出典: github.com/AndrewP-GH/cc-statusline
bash 📁 アプリ 🌿 メイン 🧠 Opus
🧊 22% [==----------] 💰 $0.41 ($2.41/時間)
⏳ 5 時間: 26% [==----------] 7 日: 7% [----------]
3 行のステータス: ディレクトリ、git ブランチ、モデルが一番上にあります。セッションコストと時間当たりの燃焼速度 ($/h) を示すコンテキストフィルバー。次に、コンテキスト警告アイコンが付いた 5 時間および週ごとのレート制限バー。出典: github.com/kangraemin/claude-status-bar
バッシュ5時間▓░░░░ 🟢
緑/黄/赤のステータス ドットでコンテキスト ウィンドウの使用状況を示す、最小の 5 スロットのブロック バー。出典: github.com/gwiener/claude-burn-bars
bash ネットワーク Opus 4.8 │ メイン │ アプリ │ ▓ ▓ ░ ░ ░ ░ ░ ░ ░ ░ 44k/200k │ $0.41
Catppuccin Frappé – モデル、git ブランチとダーティ フラグ、フォルダー、ブロック コンテキスト バー、およびセッション コストを含むステータス ライン。出典: github.com/danjdewhurst/cc-status-line
ノード ◆ Opus 4.8 │ ▪ app │ ⎇ main │ ◷ 16:02 │ █ █ ▄ ░ ░ ░ ░ ░ ░ ░ 26% (44K/167K) │ 26% ↻2h6m │ 7% ⟳2d0h ノードステータス行にautocompact 対応コンテキスト バー (compac を減算)

ション バッファー）、プロジェクト、ブランチ、最終アクティビティ時間、および 5 時間/7 日のレート制限セグメント。ソース: github.com/philipshurpik/claude-code-status-line
bash アプリ · メイン · Opus-4.8 · 努力高 · ctx 22% · 5h 26% ↻2h6m · 7d 7% ↻2d0h ディレクトリ、git ブランチ (汚れている場合は * 付き)、モデル、ライブ エフォート レベル、コンテキスト ウィンドウの使用状況、および 5 時間/7 日のレート制限の使用状況を示す 1 行のステータス バー。

## Original Extract

The Claude Code status line catalogue

A little while ago I was searching for a nice status line to use in Claude Code and was surprised that no community library type website existed for people to show off what they have made. So I decided to have a little fun and make one. All status lines are manually reviewed and rendered in a live container to capture the output. Enjoy! Source code: https://github.com/NathanAB/statuslin.es

statuslin.es statuslin . es Sign in with GitHub statuslin . es ▒
The status line is the bar at the bottom of your Claude Code terminal. Submit your own custom config for others, or find one you like and use it for yourself.
Submit a status line Dreambase Panel
python ╭─ CONTEXT ────────────────────────────────────────────╮
│ ████▊░░░░░░░░░░░░░░░░░ 22% of 200K · GOOD │
├─ STATS ──────────────────────────────────────────────┤
│ ↓ 44.0K ↑ 1.4K │ $0.41 │ ⏱ 10m 12s │ +128 / -34 │
├─ REPO ───────────────────────────────────────────────┤
│ ◆ Opus 4.8 · main · 17:27 │
╰──────────────────────────────────────────────────────╯
A boxed multi-line panel with auto-sizing borders: a context section with a gradient bar and status word, a stats section (tokens, cost, duration, lines changed), and a repo section (model, git branch with file counts, clock, vim mode), in a terracotta Claude-brand palette. Source: github.com/kyleledbetter/claudecode-statusline
bash  app   main   Opus 4.8  ━━ ━ ━ ━ ━ ━ ━ ━ ━ 22%   128  34    
A Powerline-style status line with nerd-font separators and a Dracula palette — colored directory, git branch, a per-model icon, and a faded context progress bar with an arrow tail. Source: github.com/LindseyB/claude-statusline
bash Opus 4.8 | main | ██░░░░░░░░ 22.0% | 44k / 200k | 02:07 26% | Wed 7%
Single-line status: model, git branch, a block context burn bar (green→orange→red) with token count, and 5-hour and weekly rate-limit percentages with reset times. Source: github.com/bitcoin21ideas/claude-statusline
bash Opus 4.8 [high] ~/app main
17:27:05 | ⛁ ██░░░░░░░░ 22% | 5h ●◔○○○ 26% ↻2h7m | 7d ◔○○○○○○ 7% ↻2d1h │ $0.41 ⏱ 10m
Two-line status with a block context bar and quarter-step dot meters (◔◑◕●) for the 5-hour and weekly limits — fill shows usage, color shows whether you are ahead of or behind pace — plus model, effort, git and cost. Needs no jq. Source: github.com/AndrewP-GH/cc-statusline
bash 📁 app 🌿 main 🧠 Opus
🧊 22% [==--------] 💰 $0.41 ($2.41/h)
⏳ 5h: 26% [==--------] 7d: 7% [----------]
Three-line status: directory, git branch and model up top; a context-fill bar with session cost and an hourly burn rate ($/h); then 5-hour and weekly rate-limit bars with context-warning icons. Source: github.com/kangraemin/claude-status-bar
bash 5h ▓░░░░ 🟢
A minimal 5-slot block bar showing context-window usage with a green/yellow/red status dot. Source: github.com/gwiener/claude-burn-bars
bash network Opus 4.8 │ main │ app │ ▓ ▓ ░ ░ ░ ░ ░ ░ ░ ░ 44k/200k │ $0.41
Catppuccin Frappé–themed status line with model, git branch and dirty flag, folder, a block context bar, and session cost. Source: github.com/danjdewhurst/cc-status-line
node ◆ Opus 4.8 │ ▪ app │ ⎇ main │ ◷ 16:02 │ █ █ ▄ ░ ░ ░ ░ ░ ░ ░ 26% (44K/167K) │ 26% ↻2h6m │ 7% ⟳2d0h Node status line with an autocompact-aware context bar (subtracts the compaction buffer), project, branch, last-activity time, and 5h/7d rate-limit segments. Source: github.com/philipshurpik/claude-code-status-line
bash app · main · Opus-4.8 · effort high · ctx 22% · 5h 26% ↻2h6m · 7d 7% ↻2d0h One-line status bar showing your directory, git branch (with a * when dirty), model, live effort level, context-window usage, and 5-hour/7-day rate-limit usage.
