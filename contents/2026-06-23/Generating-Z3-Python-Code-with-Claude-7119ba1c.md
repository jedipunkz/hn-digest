---
source: "https://www.johndcook.com/blog/2026/06/20/z3-python-claude/"
hn_url: "https://news.ycombinator.com/item?id=48645223"
title: "Generating Z3/Python Code with Claude"
article_title: "Generating Z3/Python code with Claude"
author: "ibobev"
captured_at: "2026-06-23T15:01:48Z"
capture_tool: "hn-digest"
hn_id: 48645223
score: 2
comments: 0
posted_at: "2026-06-23T14:08:21Z"
tags:
  - hacker-news
  - translated
---

# Generating Z3/Python Code with Claude

- HN: [48645223](https://news.ycombinator.com/item?id=48645223)
- Source: [www.johndcook.com](https://www.johndcook.com/blog/2026/06/20/z3-python-claude/)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T14:08:21Z

## Translation

タイトル: Claude による Z3/Python コードの生成
記事のタイトル: Claude を使用した Z3/Python コードの生成
説明: 私は最近、チェスの問題を解決するコードを生成する LLM の取得についていくつかの投稿を書きました。 1 つ目は Claude を使用して Prolog を生成し、2 つ目は ChatGPT を使用して Prolog を生成しました。この投稿では、Claude を使用して Z3/Python コードを生成します。このパズルは私が以前に書いたものです: パイをすべて配置します
[切り捨てられた]

記事本文:
コンテンツにスキップ
数学
信号処理
私は最近、チェスの問題を解決するためのコードを生成する LLM の取得についていくつかの投稿を書きました。 1 つ目は Claude を使用して Prolog を生成し、2 つ目は ChatGPT を使用して Prolog を生成しました。この投稿では、Claude を使用して Z3/Python コードを生成します。
このパズルは私が以前に書いたものです:
すべての駒 (キング、クイーン、ビショップ 2 人、ナイト 2 人、ルーク 2 人) を 6 × 5 のチェス盤に配置します。ただし、2 人のビショップが反対色のマス目にいて、どの駒も他の駒を攻撃していないことが条件です。
ちなみに、上記のように「駒」にはポーンを含まないのが一般的です。では、チェス盤の上にあるものすべてを何と呼ぶのでしょうか?これらを「チェスの駒」と呼ぶこともできます。この場合、ポーンは「チェスの駒」ではありますが、「駒」ではありません。慣例の 1 つは、駒とポーンを含めるのに「チェスマン」または単に「メン」を使用することです。
Z3/Python コードを作成して、次のチェス パズルのすべての解決策を見つけます。
すべての駒 (キング、クイーン、ビショップ 2 人、ナイト 2 人、ルーク 2 人) を 6 × 5 のチェス盤に配置します。ただし、2 人のビショップが反対色のマス目にいて、どの駒も他の駒を攻撃していないことが条件です。
コードは 192 = 8 × 24 の解を見つけました。 8 という因数は、ビショップ、ナイト、ルークのペアを交換する 2 3 つの方法から得られます。スクリプトのレポート
合計生溶液: 192
独自のソリューション (ピースペアスワップの重複排除): 24
── 解決策1 ──
0 1 2 3 4
0 。 K. 。 N
1. 。 。 。 B
２． 。 R . 。
3Q. 。 。 。
４． 。 。 R .
５． B . 。 N
キング: (0,1)
女王: (3,0)
ビショップ1: (5,1) [光]
ビショップ2: (1,4) [ダーク]
騎士1: (5,4)
ナイト 2: (0,4)
ルーク1: (4,3)
ルーク2: (2,2)
── 解決策2 ──
0 1 2 3 4
0N。 。 K.
1B. 。 。 。
２． 。 R . 。
３． 。 。 。 Q
４． R . 。 。
5N。 。 B .
キング: (0,3)
女王: (3,4)
ビショップ1: (5,3) [光]
ビショップ2: (1,0) [ダーク]
K

夜1: (5,0)
騎士2: (0,0)
ルーク1: (4,1)
ルーク2: (2,2)
...
──解決策24──
0 1 2 3 4
0B. 。 B .
1N。 。 。 。
２． 。 。 。 Q
３． 。 R . 。
４． R . 。 。
5N。 。 K.
キング: (5,3)
女王: (2,4)
ビショップ1: (0,0) [光]
ビショップ2: (0,3) [ダーク]
ナイト 1: (5,0)
ナイト 2: (1,0)
ルーク1: (3,2)
ルーク2: (4,1)
コードは次のとおりです。
「」
Z3 を使用したチェス パズル ソルバー
キング、クイーン、ビショップ 2 名、ナイト 2 名、ルーク 2 名を 6×5 のボードに配置します。
次のような:
- 2 人の司教は反対色の正方形にいます
- 他の駒を攻撃する駒はありません
「」
z3 インポートから (
Int、ソルバー、OR、And、Not、Distinct、If、
sat、ForAll、暗黙的
)
itertoolsをインポートする
行 = 6
コル = 5
TOTAL_CELLS = 行 * 列数 # 30
Pieces = ["キング", "クイーン", "ビショップ1", "ビショップ2", "ナイト1", "ナイト2", "ルーク1", "ルーク2"]
N = len(PIECES) # 8 個
デフォルトセル(r, c):
r * COLS + c を返す
def pos_to_rc(pos):
return pos // COLS, pos % COLS
デフォルトカラー(r, c):
"""四角い色: 0=明るい、1=暗い"""
戻り値 (r + c) % 2
def Attacks_king(r1, c1, r2, c2):
戻り値 max(abs(r1 - r2), abs(c1 - c2)) == 1
def Attacks_queen(r1, c1, r2, c2):
return (r1 == r2 または c1 == c2 または
abs(r1 - r2) == abs(c1 - c2))
def Attacks_bishop(r1, c1, r2, c2):
戻り値 abs(r1 - r2) == abs(c1 - c2)
def 攻撃_ルーク(r1, c1, r2, c2):
r1 == r2 または c1 == c2 を返す
def Attacks_knight(r1, c1, r2, c2):
dr、dc = abs(r1 - r2)、abs(c1 - c2)
return (dr == 1 および dc == 2) または (dr == 2 および dc == 1)
# ピースインデックスを攻撃関数にマップします
ATTACK_FN = {
0: 攻撃_キング、
1: 攻撃クイーン、
2: 攻撃_ビショップ、
3: 攻撃_ビショップ、
4: アタックナイト、
5: アタックナイト、
6: 攻撃_飛車、
7: 攻撃_ルーク、
}
defsolve_all():
ソリューション = []
# 各ピースは [0, TOTAL_CELLS) の整数の位置を取得します。
pos = [Int(f"pos_{p}") for p (PIECES)]
s = ソルバー()
# 有効範囲内のすべての位置
正位置の p の場合:
s.add(p >= 0, p
関連記事：教訓の学び

Z3 SAT/SMT ソルバーを使用した場合
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
私と同僚は、データ プライバシー、応用数学、統計に関する複雑な問題を企業が解決できるよう支援してきた数十年のコンサルティング経験があります。
話しましょう。私たちも貴社を支援する機会を模索することを楽しみにしています。

## Original Extract

I've written a couple posts lately on getting an LLM to generate code to solve chess problems. The first used Claude to generate Prolog and the second used ChatGPT to generate Prolog. This post will use Claude to generate Z3/Python code. The puzzle is one I've written about before: Place all the pie
[truncated]

Skip to content
MATH
SIGNAL PROCESSING
I’ve written a couple posts lately on getting an LLM to generate code to solve chess problems. The first used Claude to generate Prolog and the second used ChatGPT to generate Prolog. This post will use Claude to generate Z3/Python code.
The puzzle is one I’ve written about before :
Place all the pieces—king, queen, two bishops, two knights, and two rooks—on a 6 × 5 chessboard, with the requirement that the two bishops be on opposite colored squares and no piece is attacking another.
Incidentally, it’s common for “piece” to exclude pawns, as above. But then what do you call all the things on a chessboard? You might call them “chess pieces,” in which case a pawn is a “chess piece” but not a “piece.” One convention is to use “chessmen” or simply “men” to include pieces and pawns.
Write Z3/Python code to find all solutions to the following chess puzzle.
Place all the pieces—king, queen, two bishops, two knights, and two rooks—on a 6 × 5 chessboard, with the requirement that the two bishops be on opposite colored squares and no piece is attacking another.
The code found 192 = 8 × 24 solutions. The factor of 8 comes from 2 3 ways of swapping the pairs of bishops, knights, and rooks. The script reports
Total raw solutions: 192
Unique solutions (deduplicating piece-pair swaps): 24
── Solution 1 ──
0 1 2 3 4
0 . K . . N
1 . . . . B
2 . . R . .
3 Q . . . .
4 . . . R .
5 . B . . N
King: (0,1)
Queen: (3,0)
Bishop1: (5,1) [light]
Bishop2: (1,4) [dark]
Knight1: (5,4)
Knight2: (0,4)
Rook1: (4,3)
Rook2: (2,2)
── Solution 2 ──
0 1 2 3 4
0 N . . K .
1 B . . . .
2 . . R . .
3 . . . . Q
4 . R . . .
5 N . . B .
King: (0,3)
Queen: (3,4)
Bishop1: (5,3) [light]
Bishop2: (1,0) [dark]
Knight1: (5,0)
Knight2: (0,0)
Rook1: (4,1)
Rook2: (2,2)
...
── Solution 24 ──
0 1 2 3 4
0 B . . B .
1 N . . . .
2 . . . . Q
3 . . R . .
4 . R . . .
5 N . . K .
King: (5,3)
Queen: (2,4)
Bishop1: (0,0) [light]
Bishop2: (0,3) [dark]
Knight1: (5,0)
Knight2: (1,0)
Rook1: (3,2)
Rook2: (4,1)
Here’s the code.
"""
Chess Puzzle Solver using Z3
Place king, queen, 2 bishops, 2 knights, 2 rooks on a 6×5 board
such that:
- The two bishops are on opposite-colored squares
- No piece attacks any other piece
"""
from z3 import (
Int, Solver, Or, And, Not, Distinct, If,
sat, ForAll, Implies
)
import itertools
ROWS = 6
COLS = 5
TOTAL_CELLS = ROWS * COLS # 30
PIECES = ["king", "queen", "bishop1", "bishop2", "knight1", "knight2", "rook1", "rook2"]
N = len(PIECES) # 8 pieces
def cell(r, c):
return r * COLS + c
def pos_to_rc(pos):
return pos // COLS, pos % COLS
def color(r, c):
"""Square color: 0=light, 1=dark"""
return (r + c) % 2
def attacks_king(r1, c1, r2, c2):
return max(abs(r1 - r2), abs(c1 - c2)) == 1
def attacks_queen(r1, c1, r2, c2):
return (r1 == r2 or c1 == c2 or
abs(r1 - r2) == abs(c1 - c2))
def attacks_bishop(r1, c1, r2, c2):
return abs(r1 - r2) == abs(c1 - c2)
def attacks_rook(r1, c1, r2, c2):
return r1 == r2 or c1 == c2
def attacks_knight(r1, c1, r2, c2):
dr, dc = abs(r1 - r2), abs(c1 - c2)
return (dr == 1 and dc == 2) or (dr == 2 and dc == 1)
# Map piece index to its attack function
ATTACK_FN = {
0: attacks_king,
1: attacks_queen,
2: attacks_bishop,
3: attacks_bishop,
4: attacks_knight,
5: attacks_knight,
6: attacks_rook,
7: attacks_rook,
}
def solve_all():
solutions = []
# Each piece gets an integer position in [0, TOTAL_CELLS)
pos = [Int(f"pos_{p}") for p in PIECES]
s = Solver()
# All positions in valid range
for p in pos:
s.add(p >= 0, p
Related post : Lessons Learned With the Z3 SAT/SMT Solver
Your email address will not be published. Required fields are marked *
My colleagues and I have decades of consulting experience helping companies solve complex problems involving data privacy , applied math , and statistics .
Let’s talk. We look forward to exploring the opportunity to help your company too.
