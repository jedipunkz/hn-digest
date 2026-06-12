---
source: "https://www.johndcook.com/blog/2026/06/11/prolog-claude/"
hn_url: "https://news.ycombinator.com/item?id=48504176"
title: "Solving a chess puzzle with Claude and Prolog"
article_title: "Solving a chess puzzle with Claude and Prolog"
author: "ibobev"
captured_at: "2026-06-12T16:08:10Z"
capture_tool: "hn-digest"
hn_id: 48504176
score: 2
comments: 0
posted_at: "2026-06-12T14:00:17Z"
tags:
  - hacker-news
  - translated
---

# Solving a chess puzzle with Claude and Prolog

- HN: [48504176](https://news.ycombinator.com/item?id=48504176)
- Source: [www.johndcook.com](https://www.johndcook.com/blog/2026/06/11/prolog-claude/)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T14:00:17Z

## Translation

タイトル: クロードとプロローグと一緒にチェスのパズルを解く
説明: Martin Gardner がクロードを使用して Prolog コードを生成するチェス問題を解決します。

記事本文:
コンテンツにスキップ
数学
信号処理
クロードとプロローグと一緒にチェスのパズルを解く
Prolog はオリジナルの論理プログラミング言語です。この名前は、ロジックのプログラミングに由来しています。具体的には、この言語の発明者であるフィリップ・ルーセルがフランス人であるため、この名前は「programmation en log ique」に由来しています。
プロローグにはメリットとデメリットがあります。利点の 1 つは、この言語が論理的な問題を直接表現していることです。欠点の 1 つは、構文が奇妙になる可能性があることです。しかし、LLM がコードを作成している場合、または少なくともコードの作成を支援している場合、構文はそれほど重要ではありません。
私が 1 年ちょっと前に書いた Martin Gardner のチェス パズルをクロード (ソネット 4.6、中程度の努力) がどれだけ解けるかを試してみたかったのです。 LLM はパズルに関する無数の記事の 1 つを引用するだけで済むため、エイト クイーン パズルのようなものではなく、比較的あいまいな問題を選択しました。
昨年の投稿で述べたように、タスクは 4 × 4 のチェス盤にルーク 2 枚、ビショップ 2 枚、ナイト 2 枚を配置して、どの駒も他の駒を攻撃しないようにすることです。
基本的なソリューションは 2 つあり、回転と反射を別のソリューションとしてカウントすると 12 になります。
クロードは SWI-Prolog プログラムを作成し、私はそれを実行しました
swipl -g "実行、停止" chess_placement.pl
そして、次の出力が得られました。
合計ソリューション (生): 12
D4 対称までのソリューション: 2
D4 対称までのソリューション: 2
解決策 1:
B . 。 N
N. 。 B
。 R . 。
。 。 R .
解決策 2:
。 B N 。
。 。 。 R
R . 。 。
。 N B 。
クロードは解決策の素晴らしいイメージを作りました。
回転と反転を組み合わせて、最初のソリューションの 8 つのバリエーションを作成できることに注意してください。 2 回転すると開始位置に戻るため、2 番目の解決策には 4 つのバリエーションしかありません。
クロードが生成した Prolog コードは次のとおりです。
% ===

====================================================
ルーク 2 枚、ビショップ 2 枚、ナイト 2 枚の非攻撃配置の割合 (%)
4x4 チェス盤上の %
%
% SWI-Prolog で実行:
% swipl -g "実行、停止" chess_placement.pl
%
% または対話的に:
% swipl chess_placement.pl
% ?- 実行します。
% ?- count_solutions(N)。
% ?- count_unique(N)。
% ===========================================================
:- use_module(ライブラリ(リスト))。
% ---- ボードの正方形: (行、列) 0..3 の行、列 ----------------
square(R, C) :- メンバー(R, [0,1,2,3])、メンバー(C, [0,1,2,3])。
% ---- 攻撃述語 --------------------------------------
rook_攻撃(R,_,R,_)。 % 同じ行
rook_攻撃(_,C,_,C)。 % 同じ列
ビショップ攻撃(R1,C1,R2,C2) :-
DRはabs(R1-R2)、DCはabs(C1-C2)、
DR =:= DC、DR > 0。
騎士の攻撃(R1,C1,R2,C2) :-
DRはabs(R1-R2)、DCはabs(C1-C2)、
( DR =:= 1、 DC =:= 2 ; DR =:= 2、 DC =:= 1 )。
% 攻撃(+Type1、+R1、+C1、+Type2、+R2、+C2)
% どちらかの駒がもう一方の正方形を脅かす場合は True。
攻撃(ルーク、R1、C1、_、R2、C2) :- ルーク_攻撃(R1、C1、R2、C2)。
攻撃(_, R1,C1, ルーク, R2,C2) :- ルーク_攻撃(R1,C1,R2,C2)。
攻撃(ビショップ、R1、C1、_、R2、C2) :- ビショップ_攻撃(R1、C1、R2、C2)。
攻撃(_, R1,C1, ビショップ,R2,C2) :- ビショップ_攻撃(R1,C1,R2,C2)。
攻撃(ナイト、R1、C1、_、R2、C2) :- ナイト_攻撃(R1、C1、R2、C2)。
攻撃(_, R1,C1, ナイト,R2,C2) :- ナイト_攻撃(R1,C1,R2,C2)。
% ---- すべてのペアは安全です --------------------------------------------------
% リスト内の順序付けされていないペアをすべてチェックします。
no_アタック_ペア(_, [])。
no_ Attack_pair(T1-(R1,C1), [T2-(R2,C2)|レスト]) :-
\+ 攻撃(T1,R1,C1,T2,R2,C2)、
no_ Attack_pair(T1-(R1,C1), レスト)。
all_pairs_safe([])。
all_pairs_safe([P|休息]) :-
no_ Attack_pair(P, レスト),
all_pairs_safe(残り)。
% ---- プレースメントを生成する --------------------------------------
%
% 配置 = [rook-(R1,C

1)、ルーク-(R2,C2)、
% ビショップ-(R3,C3)、ビショップ-(R4,C4)、
% ナイト-(R5,C5)、ナイト-(R6,C6)]
%
% カウントを避けるために、各ペア内で順序付けを強制します。
% 同じタイプの重複 (rook1 と rook2 が交換)
[切り捨てられた]
「クロードとプロローグと一緒にチェスパズルを解く」についての 3 つの考え
カルロ・ドンゼラ
2026年6月11日 14:37
この問題はクロードにどのように説明されましたか?
ジョン
2026年6月11日 15:06
「4 × 4 のチェス盤を考えてみましょう。ルーク 2 人、ビショップ 2 人、ナイト 2 人を盤上に配置して、どの駒も他の駒を攻撃しないようにすることは可能でしょうか? 対称にするための解はいくつありますか? Prolog コードを生成して答えを検証してください。」
ロス
2026年6月11日 20:44
とても印象的です！
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
私と同僚は、データ プライバシー、応用数学、統計に関する複雑な問題を企業が解決できるよう支援してきた数十年のコンサルティング経験があります。
話しましょう。私たちも貴社を支援する機会を模索することを楽しみにしています。

## Original Extract

Solving a chess problem by Martin Gardner using Claude to generate Prolog code.

Skip to content
MATH
SIGNAL PROCESSING
Solving a chess puzzle with Claude and Prolog
Prolog is the original logic programming language. The name comes from pro gramming in log ic. More specifically, the name comes from pro grammation en log ique because the inventor of the language, Philippe Roussel, is French.
Prolog has its advantages and disadvantages. One of the advantages is that the language represents logical problems directly. One of the disadvantages is that the syntax can be quirky. But if an LLM is writing the code, or at least helping to write the code, the syntax doesn’t matter so much.
I wanted to see how well Claude (Sonnet 4.6, medium effort) could solve a chess puzzle by Martin Gardner that I wrote about a little over a year ago. I chose a relatively obscure problem rather than something like the Eight Queens puzzle because an LLM could simply quote one of countless articles on the puzzle.
As I stated in the post last year, the task is to place two rooks, two bishops, and two knights on a 4 by 4 chessboard so that no piece attacks any other.
There are two basic solutions, twelve if you count rotations and reflections as different solutions.
Claude wrote an SWI-Prolog program that I ran with
swipl -g "run, halt" chess_placement.pl
and it gave the following output.
Total solutions (raw): 12
Solutions up to D4 symmetry: 2
Solutions up to D4 symmetry: 2
Solution 1:
B . . N
N . . B
. R . .
. . R .
Solution 2:
. B N .
. . . R
R . . .
. N B .
Claude made a nice image of the solutions.
Note that you can make eight variations of the first solution by a combination of rotations and flips. There are only four variations of the second solution because two rotations put you back where you started.
Here’s the Prolog code that Claude generated.
% ============================================================
% Non-attacking placement of 2 Rooks, 2 Bishops, 2 Knights
% on a 4x4 chessboard
%
% Run with SWI-Prolog:
% swipl -g "run, halt" chess_placement.pl
%
% Or interactively:
% swipl chess_placement.pl
% ?- run.
% ?- count_solutions(N).
% ?- count_unique(N).
% ============================================================
:- use_module(library(lists)).
% ---- Board squares: (Row, Col) with Row, Col in 0..3 ----------------
square(R, C) :- member(R, [0,1,2,3]), member(C, [0,1,2,3]).
% ---- Attack predicates ----------------------------------------------
rook_attacks(R,_,R,_). % same row
rook_attacks(_,C,_,C). % same column
bishop_attacks(R1,C1,R2,C2) :-
DR is abs(R1-R2), DC is abs(C1-C2),
DR =:= DC, DR > 0.
knight_attacks(R1,C1,R2,C2) :-
DR is abs(R1-R2), DC is abs(C1-C2),
( DR =:= 1, DC =:= 2 ; DR =:= 2, DC =:= 1 ).
% attacks(+Type1, +R1, +C1, +Type2, +R2, +C2)
% True if either piece threatens the other's square.
attacks(rook, R1,C1, _, R2,C2) :- rook_attacks(R1,C1,R2,C2).
attacks(_, R1,C1, rook, R2,C2) :- rook_attacks(R1,C1,R2,C2).
attacks(bishop, R1,C1, _, R2,C2) :- bishop_attacks(R1,C1,R2,C2).
attacks(_, R1,C1, bishop,R2,C2) :- bishop_attacks(R1,C1,R2,C2).
attacks(knight, R1,C1, _, R2,C2) :- knight_attacks(R1,C1,R2,C2).
attacks(_, R1,C1, knight,R2,C2) :- knight_attacks(R1,C1,R2,C2).
% ---- All pairs safe -------------------------------------------------
% Check all unordered pairs in a list.
no_attack_pair(_, []).
no_attack_pair(T1-(R1,C1), [T2-(R2,C2)|Rest]) :-
\+ attacks(T1,R1,C1,T2,R2,C2),
no_attack_pair(T1-(R1,C1), Rest).
all_pairs_safe([]).
all_pairs_safe([P|Rest]) :-
no_attack_pair(P, Rest),
all_pairs_safe(Rest).
% ---- Generate a placement -------------------------------------------
%
% Placement = [rook-(R1,C1), rook-(R2,C2),
% bishop-(R3,C3), bishop-(R4,C4),
% knight-(R5,C5), knight-(R6,C6)]
%
% We enforce an ordering within each pair to avoid counting
% same-type duplicates (rook1 vs rook2 swapped)
[truncated]
3 thoughts on “ Solving a chess puzzle with Claude and Prolog ”
Carlo Donzella
11 June 2026 at 14:37
how was the problem described to Claude?
John
11 June 2026 at 15:06
“Consider a 4 by 4 chess board. Is it possible to place two rooks, two bishops, and two knights on the board so that no piece is attacking any other? How many solutions are there up to symmetry? Generate Prolog code to verify the answer.”
ross
11 June 2026 at 20:44
So impressive!
Your email address will not be published. Required fields are marked *
My colleagues and I have decades of consulting experience helping companies solve complex problems involving data privacy , applied math , and statistics .
Let’s talk. We look forward to exploring the opportunity to help your company too.
