---
source: "https://blog.trine.dev/posts/2026-02-28-df-ai-exp/"
hn_url: "https://news.ycombinator.com/item?id=49208128"
title: "Building an LLM Agent to Play Dwarf Fortress"
article_title: "Building an LLM Agent to Play Dwarf Fortress | trine"
author: "totetsu"
captured_at: "2026-08-07T10:42:43Z"
capture_tool: "hn-digest"
hn_id: 49208128
score: 1
comments: 0
posted_at: "2026-08-07T10:01:06Z"
tags:
  - hacker-news
  - translated
---

# Building an LLM Agent to Play Dwarf Fortress

- HN: [49208128](https://news.ycombinator.com/item?id=49208128)
- Source: [blog.trine.dev](https://blog.trine.dev/posts/2026-02-28-df-ai-exp/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T10:01:06Z

## Translation

タイトル: ドワーフ要塞をプレイするための LLM エージェントの構築
記事のタイトル: ドワーフ要塞をプレイするための LLM エージェントの構築 |トライン
説明: それが原因ではありません

記事本文:
ドワーフ要塞をプレイするための LLM エージェントの構築 |トライン
コンテンツへスキップ
投稿
戻る ドワーフ要塞をプレイするための LLM エージェントの構築
公開日: 2026 年 2 月 28 日 |午後 6 時 目次
Dwarf Fortress はこれまでに作られた中で最も複雑なゲームだと言われています。それが私がそれを選んだ理由ではありません。
抽象的なので選びました。
NetHack、DF — これらは Go のように構造化されています。ほぼ無限の新たな複雑性を生み出す小さなルールセット。囲碁は 19 × 19 の盤で、簡単な攻略ルールを備えています。 DF はターミナル上の ASCII 文字、数十の DFHack コマンドですが、地下の川を壁で囲うのを忘れると、数分で要塞が崩壊する可能性があります。
その構造 — シンプルなルール、巨大な状態空間 — は、実際には AI にとって理想的です。
状態はピクセルではなく構造化されます。テキスト出力とデータ ファイル。 LLM はビジョン モデルなしで直接読み取ることができます
アクション空間は離散的です。定義された一連のコマンド (prospect、quickfort run、dig-now、showmood …)
フィードバックは明確です。ドワーフは死ぬか生きているかです。食事の数は数字です。部屋が掘られたかどうかはファイルシステムチェックです
記憶可能な最適な戦略はありません。すべての世界は手続き的に生成されます。エージェントは状況を実際に理解する必要がある
最近の 3D ゲームは、ピクセル、フレーム レート、オクルージョンを扱うため、逆説的にエージェントを構築するのが難しくなります。 DF の ASCII インターフェイスは偶然の利点です。
基本的なアイデア: ゲーム UI には決して触れないでください。 DFHack のみに話しかけてください。
ナレッジ レイヤー — DF 戦略のナレッジ + dreamfort ブループリント ライブラリ + エクスペリエンス ログ。フェーズ固有のコンテキストに抽出され、LLM プロンプトに挿入されます。
デシジョン層 (LLM) — 現在の状態 + フェーズの知識 + アクション履歴を受け取り、構造化されたアクション JSON を出力します。アクションの種類: dfhack (quickfort/dig-now/build-now)、キーストローク (UI ナビゲーション)、クエリ (Lua 状態プローブ)、wait 、done 。

実行層 — dfhack-run / xdotool / Lua を実行し、各ステップの後に JSONL ログを書き込み、状態を更新します。
フィードバック レイヤー — プライマリ: DFHack Lua クエリ (ユニット、建物、アイテム、マップ、見通し)。フォールバック: コマンドが状態を判断できない場合の scrot スクリーンショット + Vision API。イベント: gamelog.txt リアルタイム監視 - 戦闘/死亡/気分/季節の変化により即時に対応します。
VPS の TEXT モード。 DF には PRINT_MODE:TEXT があり、Xvfb ではなく ncurses TUI としてレンダリングされ、PTY でヘッドレスで実行されます。
Lua RPC ではなく、dfhack-run コマンド パイプ。 Lua RPC チャネルはヘッドレスで不安定です。 dfhack-run コマンド パイプは確実に動作します。
--cursor を使用したクイックフォート。部屋を掘るのにメニューナビゲーションは必要ありません。クイックフォートは、 blueprint.csv --cursor x,y,z を実行します。 AI に必要なのは UI スキルではなく、座標です。
状態オラクルとしてのファイルシステム。要塞は動いているのか？ data/save/region*/region_snapshot-*.dat をカウントします。何ティック?ファイル名はそれをエンコードします。基本的な状態には内部クエリは必要ありません。
フェーズ 1 の目標: 生存の基本 — 食べ物、飲み物、住居。
この計画では、ドリームフォートのブループリントを順番に実行します: セットアップ → 地表の整地 → 農業レベルの掘削 → ワークショップの配置 → 農場区画 + ダイニング ルーム → 一時停止解除 → 検証。
エージェントは実行され、状態を読み取り、それを介して推論します。興味深いデバッグの瞬間 - DF の時間システムの追跡:
cur_year_tick は、年間のティック カウンタです。 cur_season_tick はシーズン内です。 1 シーズンは ~33600 ティック (1200 ティック/日 × 28 日) です。エージェントは、季節はまだ春の半ばで、作物が熟すには夏に進む必要があると判断しました。ヘッドレス設定ではゲームの実行が遅くなりました。DFHack のタイムストリームを使用して高速化しました。
建物は 1 → 45 (農業 2) → 68 (農業 3) になりました。 23 の吊り下げられた建物は正常です。建築計画はまだ製造されていない材料を待っています。計画が完了しました。
中期：段階的な目標

ここから続きます。
フェーズ 2: 生産チェーン (農業 → 醸造 → 木工 → 製錬)
フェーズ 3: 防御 (廊下、罠、軍事)
各フェーズには、独自の目標関数と検証可能な成功条件があります。
長期: セッション間の記憶。 Fortress が失敗しました → 理由を記録 → 次のロード時にコンテキストを挿入します。 DF の gamelog.txt には、要塞の開始からのすべてのイベントが記録されます。それは自然なエピソードの記憶です。 LLM は現在の状況を計画するだけでなく、実行全体で経験を蓄積します。
ゲームは負けることを中心に設計されています。それぞれの失敗は信号です。

## Original Extract

Not because it

Building an LLM Agent to Play Dwarf Fortress | trine
Skip to content trine
Posts
Go back Building an LLM Agent to Play Dwarf Fortress
Published: Feb 28, 2026 | at 06:00 PM Table of contents
People say Dwarf Fortress is the most complex game ever made. That’s not why I picked it.
I picked it because it’s abstract .
NetHack, DF — they’re structured like Go. A small ruleset that produces near-infinite emergent complexity. Go is a 19×19 board with simple capture rules. DF is ASCII characters on a terminal, a few dozen DFHack commands, but forgetting to wall off an underground river can end your fortress in minutes.
That structure — simple rules, enormous state space — is actually ideal for AI:
State is structured, not pixels. Text output and data files. An LLM can read it directly without a vision model
Action space is discrete. A defined set of commands ( prospect , quickfort run , dig-now , showmood …)
Feedback is unambiguous. Dwarves are dead or alive. Food count is a number. Whether a room was dug is a filesystem check
No memorizable optimal strategy. Every world is procedurally generated. The agent has to actually understand the situation
Modern 3D games are paradoxically harder to build agents for — you deal with pixels, frame rates, occlusion. DF’s ASCII interface is an accidental advantage.
Core idea: never touch the game UI. Talk to DFHack only.
Knowledge Layer — DF strategy knowledge + dreamfort blueprint library + experience logs, distilled into phase-specific context injected into the LLM prompt.
Decision Layer (LLM) — receives current state + phase knowledge + action history, outputs a structured Action JSON. Action types: dfhack (quickfort/dig-now/build-now), keystroke (UI navigation), query (Lua state probe), wait , done .
Execution Layer — runs dfhack-run / xdotool / Lua, writes a JSONL log after each step, refreshes state.
Feedback Layer — primary: DFHack Lua queries (units, buildings, items, map, prospect). Fallback: scrot screenshot + Vision API when commands can’t determine state. Event: gamelog.txt real-time monitoring — combat/death/mood/season changes trigger immediate response.
TEXT mode on VPS. DF has PRINT_MODE:TEXT — renders as ncurses TUI, no Xvfb, runs headless in a PTY.
dfhack-run command pipe, not Lua RPC. The Lua RPC channel is unstable headless. dfhack-run command pipe works reliably.
quickfort with --cursor . To dig a room, no menu navigation needed. quickfort run blueprint.csv --cursor x,y,z . The AI needs coordinates, not UI skills.
Filesystem as state oracle. Is the fortress running? Count data/save/region*/region_snapshot-*.dat . How many ticks? The filenames encode it. No internal queries needed for basic state.
Phase 1 goal: survival basics — food, drink, shelter.
The plan runs dreamfort blueprints in sequence: setup → surface clearing → dig farming level → place workshops → farm plots + dining room → unpause → verify.
The agent runs, reads state, reasons through it. One interesting debugging moment — tracking DF’s time system:
cur_year_tick is the year-wide tick counter. cur_season_tick is within-season. A season is ~33600 ticks (1200 ticks/day × 28 days). The agent figures out it’s still spring mid-season and needs to advance to summer for crops to ripen. Game was running slow on the headless setup — used DFHack’s timestream to accelerate.
Buildings went from 1 → 45 (farming2) → 68 (farming3). 23 suspended buildings are normal — buildingplan waiting on materials not yet produced. Plan complete.
Mid-term: phase-based goals continuing from here.
Phase 2: production chain (farming → brewing → woodworking → smelting)
Phase 3: defense (corridors, traps, military)
Each phase has its own goal function and verifiable success condition.
Long-term: cross-session memory. Fortress failed → record why → inject context on next load. DF’s gamelog.txt logs every event from the beginning of the fortress. That’s a natural episode memory. The LLM becomes not just a planner for the current situation, but an experience accumulator across runs.
The game is designed around losing. Each failure is signal.
