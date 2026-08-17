---
source: "https://mkornreich.me/projects/sokoban/"
hn_url: "https://news.ycombinator.com/item?id=49330215"
title: "Show HN: Sokoban AI Solver"
article_title: "Sokoban. Menachem Kornreich"
image: "https://mkornreich.me/og-image.png"
author: "enjoyyourlife"
captured_at: "2026-08-17T13:32:58Z"
capture_tool: "hn-digest"
hn_id: 49330215
score: 7
comments: 1
posted_at: "2026-08-17T13:07:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sokoban AI Solver

- HN: [49330215](https://news.ycombinator.com/item?id=49330215)
- Source: [mkornreich.me](https://mkornreich.me/projects/sokoban/)
- Score: 7
- Comments: 1
- Posted: 2026-08-17T13:07:00Z

## Translation

タイトル: 表示 HN: 倉庫番 AI ソルバー
記事タイトル：倉庫番。メナヘム・コーンライヒ
説明: ブラウザで倉庫番をプレイし、AI ソルバーがそれを解読する様子を見てください。ネイティブ C++ エンジンのプレーン JS ポートであり、完全にクライアント側で、最適と思われる (最小移動量の) ソリューションを返します。

記事本文:
倉庫番 (「倉庫番」) は 1980 年代のパズルで、すべての箱をゴールに押し込みます。このバリエーションでは
キーパーもゴールを決めなければなりません。
倉庫は格子状になっています。ステップごとに、キーパーは上下左右に 1 マス移動します。
キーパーは壁やボックスに入ることができません。単一のボックスをプッシュできるのは、
ボックスのすぐ向こう側（プッシュ方向）の四角は空のフロアまたはゴールです。 1箱だけ
ステップごとに移動し、ボックスを再びゴールから押し出してスペースを空けることができます。
コントロール: 矢印キー、WA S D 、またはオンスクリーン パッド。
元に戻すと元に戻ります。リセットするとボードが復元されます。
目標: すべての可動エンティティが完成すると、パズルに勝利します。すべてのボックス
そして飼育員さん。ゴールの上に座っています。だからこそ、各ボードにはもう 1 つの目標があるのです
ボックスがあるよりも、最後の目標はキーパーです。
目的: できるだけ少ない手数でその状態に到達すること。いくつかのボードでは、
最適な移動数は既知であり、上に示しています。 AI (許容可能なヒューリスティックを使用)
徹底的に検索できるボード上で最適なソリューションを返します。
倉庫番は A* 検索問題ですが、1 つのキーパー ステップを 1 つずつ探索する単純なバージョンです。
混雑したボードでは時間が爆発的に増えます。ここで実行されるのは、
私が作成したネイティブ C++ 最適ソルバーのプレーン JavaScript ポート。返されるのは、
おそらく最も少ない手数
単なる解決策ではなく、解決策:
移動最適マクロプッシュ A*。各検索エッジはボックス全体のプッシュコストがかかります
(キーパーのプッシュスポットまでの最短距離) + 1 なので、合計が真の最小値になります
キーパーの移動回数、検索は個々の歩行をスキップします
ステップ。
コンパクトなビットマスク状態。ボックスは 32 ビット整数にパックされます。
ボードの到達可能な「ライブ」セルとキーパーがもう 1 つの番号になるため、状態全体が
~1 KB オブジェクトの代わりに、単一の ~8 バイトのキー。数百万の州が数十の州に収まる
MB。
ダイヤルバケットキュー + オープンアドレス

dハッシュ。 A* フロンティアはキー付きのバケット キューです
コストによって、訪問セット (ソリューションの親リンクを含む) はアパートに住んでいます。
型付き配列のハッシュ。割り当てフリーでキャッシュに優しい。
デッドロック剪定。静的なデッドスクエア テーブル (からの逆到達可能性)
目標）に加えて、解決不可能と思われるポジションを凍結チェックで破棄します。
A* を許容できる (したがって最適な) 状態に保つ、壁を意識したプッシュ距離の下限。
ボード 1 ～ 14 はミリ秒単位でライブで証明された最適解に解決されます (移動
上で「最適」として示されているカウントは、まさにこのソルバーが返すものです)。ボード15. 8ボックス
迷路。例外は、最適な検索を探索することです。
約 4,900 万の州で 1 GB を超える必要があり、これには時間がかかりすぎます。
ブラウザタブ内で実行します。したがって、その最適値（184手）が計算されました
この正確なアルゴリズムのネイティブ C++ ビルドによるオフライン (並列 A* 検索、
24 コア全体で約 5 秒）、リプレイによって検証され、ページは単純に
事前に計算されたソリューションを再生します。そのため、ボード 15 の回答はハードコードされています。
ここで検索するのではなく。
から構築
私の倉庫番ソルバー。
倉庫番について →

## Original Extract

Play Sokoban in your browser and watch an AI solver crack it. A plain-JS port of a native C++ engine, it returns the provably optimal (fewest-move) solution, entirely client-side.

Sokoban ("warehouse keeper") is a 1980s puzzle: push every box onto a goal. In this variant
the keeper must also finish on a goal.
The warehouse is a grid. On each step the keeper moves one square up, down, left or right.
The keeper cannot walk into a wall or a box. It can push a single box if the
square just beyond the box (in the push direction) is empty floor or a goal. Only one box
moves per step, and a box can be pushed out of a goal again to make room.
Controls: arrow keys or W A S D , or the on-screen pad.
Undo steps back. Reset restores the board.
Goal: the puzzle is won when every movable entity. Every box
and the keeper. Is sitting on a goal. That is why each board has one more goal
than it has boxes: the last goal is for the keeper.
Objective: reach that state in as few moves as possible. For several boards the
optimal move count is known and shown above. The AI (with an admissible heuristic)
returns an optimal solution on the boards it can search exhaustively.
Sokoban is an A* search problem, but a naive version that explores one keeper step at a
time explodes on crowded boards. What runs here is a
plain-JavaScript port of a native C++ optimal solver I wrote. It returns the
provably fewest-moves
solution, not just some solution:
Move-optimal macro-push A*. Each search edge is a whole box push costed
as (the keeper's shortest walk to the push spot) + 1, so the total is the true minimum
number of keeper moves , while the search skips over the individual walking
steps.
Compact bitmask states. The boxes are packed into a 32-bit integer over the
board's reachable "live" cells and the keeper into one more number, so a whole state is
a single ~8-byte key instead of a ~1 KB object. Millions of states fit in tens of
MB.
Dial bucket queue + open-addressed hash. The A* frontier is a bucket queue keyed
by cost, and the visited set (with the solution's parent links) lives in a flat
typed-array hash. Allocation-free and cache-friendly.
Deadlock pruning. A static dead-square table (reverse-reachability from
the goals) plus a freeze check discard provably-unsolvable positions, guided by
a wall-aware push-distance lower bound that keeps A* admissible (hence optimal).
Boards 1–14 are solved live to the proven optimum in milliseconds (the move
counts shown as "Optimal" above are exactly what this solver returns). Board 15. The 8-box
maze. Is the exception: its optimal search explores
~49 million states and needs >1 GB , which would take far too long to
run inside a browser tab. So its optimum ( 184 moves ) was computed
offline by the native C++ build of this exact algorithm (a parallel A* search,
~5 s across 24 cores) and verified by replay, and the page simply
plays that precomputed solution back . That is why board 15's answer is hardcoded
rather than searched here.
Built from
my Sokoban solver .
About Sokoban →
