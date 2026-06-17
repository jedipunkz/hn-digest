---
source: "https://maxim-saplin.github.io/llm_chess/"
hn_url: "https://news.ycombinator.com/item?id=48573937"
title: "LLM Chess – Leaderboard"
article_title: "LLM Chess Leaderboard"
author: "elwell"
captured_at: "2026-06-17T18:57:25Z"
capture_tool: "hn-digest"
hn_id: 48573937
score: 3
comments: 0
posted_at: "2026-06-17T17:45:40Z"
tags:
  - hacker-news
  - translated
---

# LLM Chess – Leaderboard

- HN: [48573937](https://news.ycombinator.com/item?id=48573937)
- Source: [maxim-saplin.github.io](https://maxim-saplin.github.io/llm_chess/)
- Score: 3
- Comments: 0
- Posted: 2026-06-17T17:45:40Z

## Translation

タイトル: LLM チェス – リーダーボード
記事のタイトル: LLM チェス リーダーボード
説明: 模擬チェスとランダムおよびコモドドラゴンの LLM を比較する LLM チェス リーダーボード。 Elo、勝敗、ゲーム時間、トークンによるランキング。

記事本文:
LLM チェス リーダーボード
LLM チェス リーダーボード
ランダム プレイヤーと LLM 間のチェス ゲームをシミュレートします。チャット モデルの (1) チェスの熟練度と (2) の評価
指示に従う能力
__ __ ____ __
/\ \ /\ \ /'\_/`\ /\ _``。 /\ \
\ \ \ \ \ \ /\ \ \ \ \/\_\\ \ \___ __ ____ ____
\ \ \ __\ \ \ __\ \ \__\ \ \ \ \/_/_\ \ _ `\ /'__`\ /',__\ /',__\
\ \ \L\ \\ \ \L\ \\ \ \_/\ \ \ \ \L\ \\ \ \ \ \/\ __//\__, `\/\__, `\
\ \____/ \ \____/ \ \_\\ \_\ \ \____/ \ \_\ \_\ \____\/\____/\/\____/
\/___/ \/___/ \/_/ \/_/ \/___/ \/_/\/_/\/____/\/___/ \/___/
__ __ __ __
/\ \ /\ \ /\ "-./ \
\ \ \____ \ \ \____ \ \ \-./\ \
\ \_____\ \ \_____\ \ \_\ \ \_\
\/_____/ \/_____/ \/_/ \/_/
______ __ __ ______ ______ ______
/\ ___\ /\ \_\ \ /\ ___\ /\ ___\ /\ ___\
\ \ \____ \ \ __ \ \ \ __\ \ \___ \ \ \___ \
\ \_____\ \ \_\ \_\ \ \_____\ \/\_____\ \/\_____\
\/_____/ \/_/\/_/ \/_____/ \/_____/ \/_____/
ランダムプレイヤー（白）
♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
· · · · · · · · ·
· · · · · · · · ·
· · · · · · · · ·
· · · · · · · · ·
♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
ゲームオーバー
大規模言語モデルはチェスをすることができますか?調べてみましょうツ
このリーダーボードはチェスのスキルと指導を評価します。
エージェント設定の場合:
LLM は複数ターンのダイアログに参加し、そこでアクションの選択肢が提示されます (例: 「ボードを取得する」または「ボードを取得する」など)。
"make move") と対戦するとき
対戦相手 (ランダム プレイヤーまたはチェス エンジン)。
2024 年、私たちはカオス モンキー ベースライン、つまり合法的な手を選択するランダム プレイヤーから始まりました。
ランダム。
当時、ほとんどのモデルはほとんど競争できず、ゲームについていけずに負けていました。
指示
(つまり、違法な動きの幻覚や間違った行動をとる)、またはゲームを 200 メートルまでドラッグすることによって

過ぎた
彼らは制限するので、
勝てなかった。
2025 年には、より有能な推論モデルが、指示への従うこととチェスのスキルの両方を確立しました。
より有能な対戦相手としてコモドドラゴン チェス エンジンを追加しました。
エロレーティングオン
チェス.com。これにより、結果を現実世界の評価スケールに固定し、Elo 評価を計算できるようになりました。
各モデル。
メトリクス:
- プレイヤー: モデル名 (ブラックとしてプレイ)。 vsドラゴンでもプレイしたモデルにはマークが付いています
上付き文字にアスタリスクが付いています (例: 3*)。
- Elo: ドラゴンのスキルレベルと調整されたランダムによって固定された推定 Elo。を解決します
集計されたブロック (対戦相手の Elo、勝ち、引き分け、負け) に対する 1D MLE とレポート ±95% CI。どちらもランダムの場合
とドラゴンのデータが存在するので、それらを組み合わせます。極端な 100% の勝敗またはアンカーなしの場合、空の Elo が表示されます。
ゲーム。
- ゲーム時間: 完了した最大ゲーム時間の割合 (0 ～ 100%)。対策
多くの動きにわたって指示に従った安定性。 100% は、モデルが原因でゲームが中断されなかったことを意味します
幻覚的な動きや行動。 50% は、平均してモデルがゲームの途中でゲーム ループを中断したことを意味します
(最大 200 回の移動のうち平均 100 回の移動が許可されます)
- トークン: 移動ごとの完了トークン。冗長性/効率性シグナル。
- コスト/Elo (メイン): 1000 Elo ポイントあたりの推定コスト (コスト/ゲームを Elo で割って、1000 でスケールします)。低いほどコスト効率が高くなります。
- コスト/ゲーム (拡張): トークンの使用量とモデルの価格設定に基づいたゲームあたりの推定コスト。
アレンジメントとソース:
- 一次ソート: Elo (DESC)、次にゲーム期間 (DESC)、トークン (ASC)。
- データ ソースには、ランダム対 LLM ゲームとドラゴン対 LLM ゲームが混在しています。ドラゴン レベルは Elo にマッピングされ、
アンカー;ランダムはまずドラゴンに対して調整され、その後多くのモデルの対戦相手として使用されます。
- Elo レーティングはプレイヤープール間で比較できません。つまり、chess.com Elo と FIDE を比較することはできません。
エロ
- Chess.com

コンテキストに使用される参照 (2025 年 9 月現在): Rapid Leaderboard (Elo
プール）、マグナス・カールセン
統計、エロの説明、
プレーヤークラス。
マトリックスの視覚化:
このプロットは、2 つの主要な指標に基づいて LLM チェスプレイヤーを示しています。
- X 軸: ゲーム期間 (0 ～ 100%) - モデルが正しい通信をどの程度維持しているかを示します。
ゲーム全体のプロトコル。値が高いほど、指示に従う能力が優れていることを示します。
- Y 軸: 勝率 (0 ～ 100%) - この指標は勝敗よりも厳密ではありません (中断なし)
指示に従っていないことによる技術的損失を無視するため、リーダーボードで使用されます。より高い値
より良いチェスの戦略と意思決定を示します。
解釈:
- 右上: 優れたチェスのスキルと指示のフォローを兼ね備えたモデル。
- 左上: チェスのスキルは高いが、コミュニケーションを維持するのに苦労しているモデル
プロトコル。
- 右下: 指示にはよく従うものの、チェスの動きが下手なモデル。
- 左下: チェスの戦略とフォローの両方に苦戦するモデル
指示。
使用されるライブラリと依存関係:
- chess: チェスのゲームのルールと基本的な操作を処理するための Python ライブラリ。
ボードの表現、正当な手の評価、ゲームの状態の評価。これはチェスエンジンではありません
最善手の実際の計算を実行します。
- AG2 (別名 Autogen) は、LLM 通信のバックボーンとして使用されます。また、
チェスボードと GameAgent、RandomPlayerAgent、AutoReplyAgent などのカスタム エージェント間の対話
その他、さまざまなプレイヤータイプをシミュレートするためのもの。
- コモドドラゴン - 実際の最善手を計算するチェスエンジン。
実際のチェスプレイヤーのパフォーマンスがどのようなものかを示すための参考資料。
ワークフロー:
1. ゲームはチェス ボードと 2 人のプレーヤーで初期化されます: ランダム プレーヤー (白として) と LLM (として)
黒）。
2. ゲーム l

oop は、チェックメイト、ステイルメイト、到達などの終了条件が満たされるまで実行されます。
最大移動数。
3. 各プレイヤーは順番に動きます。ランダムプレイヤーは、リストからランダムに手を選択します。
法的な動き。
4. LLM は、構造化されたダイアログを使用して移動するように求められます。これには、
現在の取締役会の状態を確認し、合法的な手を取得し、手を動かします。
5. ゲームの状態は各移動後に更新され、有効になっている場合はボードが視覚化されます。
6. ゲーム統計が生成され、ゲーム終了時に保存されます。
ダイアログ:
以下は、モデルに移動を促すダイアログの例です。 UCI フォーマットはユニバーサル チェスを指します
ソースと宛先の四角形を使用するインターフェイス表記 (例: e2e4 は e2 から e4 への移動を意味します):
プロキシ (Player_Black へ):
あなたはプロのチェスプレイヤーで、黒人としてプレイします。今度はあなたが行動を起こす番です。移動する前に、次のいずれかのアクションを選択できます。
- 'get_current_board' はボードのスキーマと現在のステータスを取得します。
- 「get_legal_moves」: 利用可能な手の UCI 形式のリストを取得します。
- ターンを完了する準備ができたら、「make_move <UCI formatted move>」 (例: 「make_move e2e4」)
行動で応答します。
Player_Black (プロキシへ):
get_current_board
プロキシ (Player_Black へ):
♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
· · · · · · · · ·
· · · · · · · · ·
· · · · · · · · ·
· · · · · · · · ·
♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
Player_Black (プロキシへ):
get_legal_moves
プロキシ (Player_Black へ):
a7a6、b7b6、c7c6、d7d6、e7e6、f7f6、g7g6、h7h6、a7a5、b7b5、c7c5、d7d5、e7e5、f7f5、g7g5、h7h5
Player_Black (プロキシへ):
メイクムーブ e7e5
プロキシ (Player_Black へ):
移動しました、プレイヤーを切り替えます
1 移動 = 1 ダイアログ。ダイアログには限界がある

d ～ 10 ターン (10 ペアのプロキシ/エージェント リクエスト/レスポンス)
メッセージ)。間違いは3回まで（不適合）
コミュニケーション表記、間違ったアクションや間違った動きの選択など）はダイアログごとに許可されます。
最大ターン数または LLM のミス数に達すると、ゲームは終了し、LLM に損失が与えられます。
指定されたダイアログ。チェスの戦略と指示に正確に従うことの両方に優れたモデルは、
最高のパフォーマンスを発揮します。

## Original Extract

LLM Chess Leaderboard comparing LLMs in simulated chess vs Random and Komodo Dragon. Rankings by Elo, win/loss, game duration, and tokens.

LLM Chess Leaderboard
LLM Chess Leaderboard
Simulating chess games between a Random Player and an LLM. Evaluating Chat Models' (1) chess proficiency and (2)
instruction following abilities
__ __ ____ __
/\ \ /\ \ /'\_/`\ /\ _``. /\ \
\ \ \ \ \ \ /\ \ \ \ \/\_\\ \ \___ __ ____ ____
\ \ \ __\ \ \ __\ \ \__\ \ \ \ \/_/_\ \ _ `\ /'__`\ /',__\ /',__\
\ \ \L\ \\ \ \L\ \\ \ \_/\ \ \ \ \L\ \\ \ \ \ \/\ __//\__, `\/\__, `\
\ \____/ \ \____/ \ \_\\ \_\ \ \____/ \ \_\ \_\ \____\/\____/\/\____/
\/___/ \/___/ \/_/ \/_/ \/___/ \/_/\/_/\/____/\/___/ \/___/
__ __ __ __
/\ \ /\ \ /\ "-./ \
\ \ \____ \ \ \____ \ \ \-./\ \
\ \_____\ \ \_____\ \ \_\ \ \_\
\/_____/ \/_____/ \/_/ \/_/
______ __ __ ______ ______ ______
/\ ___\ /\ \_\ \ /\ ___\ /\ ___\ /\ ___\
\ \ \____ \ \ __ \ \ \ __\ \ \___ \ \ \___ \
\ \_____\ \ \_\ \_\ \ \_____\ \/\_____\ \/\_____\
\/_____/ \/_/\/_/ \/_____/ \/_____/ \/_____/
Random Player (White)
♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
· · · · · · · ·
· · · · · · · ·
· · · · · · · ·
· · · · · · · ·
♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
GAME OVER
Can Large Language Models play chess? Let's find out ツ
This leaderboard evaluates chess skill and instruction following
in an agentic setting:
LLMs engage in multi-turn dialogs where they are presented with a choice of actions (e.g., "get board" or
"make move") when playing against
an opponent (Random Player or Chess Engine).
In 2024, we began with a chaos monkey baseline — a Random Player that chooses legal moves at
random.
At the time, most models could barely compete and lost either due to an inability to follow game
instructions
(i.e., hallucinating illegal moves or taking incorrect actions) or by dragging the game to the 200-move
limit because they
couldn't win.
In 2025, more capable reasoning models nailed both instruction following and chess skill.
We've added the Komodo Dragon Chess Engine as a more capable opponent, which is also
Elo-rated on
chess.com. This allowed us to anchor the results to a real-world rating scale and compute an Elo rating for
each model.
METRICS:
- Player: Model name (playing as Black). Models that also played vs Dragon are marked
with an asterisk in superscript (e.g., 3*).
- Elo: Estimated Elo anchored by Dragon skill levels and calibrated Random. We solve a
1D MLE over aggregated blocks (opponent Elo, wins, draws, losses) and report ±95% CI. When both Random
and Dragon data exist, they are combined. Empty Elo appears for extreme 100% win/loss or no anchored
games.
- Game Duration: Share of maximum game length completed (0-100%); measures
instruction-following stability across many moves. 100% means no games were interrupted due to model
haluscinating moves or actions. 50% means that on average the model boroke the game loop mid-game
(making an average 100 moves out of max 200 allowed)
- Tokens: Completion tokens per move; verbosity/efficiency signal.
- Cost/Elo (main): Estimated cost per 1000 Elo points (Cost/Game divided by Elo, then scaled by 1000). Lower is more cost-efficient.
- Cost/Game (extended): Estimated cost per game based on token usage and model pricing.
ARRANGEMENT & SOURCES:
- Primary sorting: Elo (DESC), then Game Duration (DESC), Tokens (ASC).
- Data sources mix Random-vs-LLM and Dragon-vs-LLM games. Dragon levels map to Elo and provide the
anchor; Random is first calibrated vs Dragon and then used as an opponent for many models.
- Elo ratings are not comparable across player pools, i.e. you can not compare chess.com Elo to FIDE
Elo
- Chess.com references used for context (as of Sep 2025): Rapid Leaderboard (Elo
pool) , Magnus Carlsen
stats , and Elo explanation &
player classes .
MATRIX VISUALIZATION:
This plots shows LLM chess players based on two key metrics:
- X-Axis: Game Duration (0-100%) - Shows how well models maintain correct communication
protocols throughout the game. Higher values indicate better instruction following ability.
- Y-Axis: Win Rate (0-100%) - The metric is less strict than Win/Loss (Non-Interrupted)
used in the leaderboard as it ignores technichal losses due to poor instruction following. Higher values
indicate better chess strategy and decision making.
INTERPRETATION:
- Top-Right: Models with both excellent chess skill and instruction following.
- Top-Left: Models with good chess skill but struggle to maintain communication
protocol.
- Bottom-Right: Models that follow instructions well but make poor chess moves.
- Bottom-Left: Models that struggle with both chess strategy and following
instructions.
Libraries and Dependencies Used:
- chess: A Python library for handling chess game rules and basic operations, including
board representation, legal move evaluation, and game state evaluation. This is not a chess engine
running the actual calculation of the best move.
- AG2 (aka Autogen) is used as a backbone for LLM communication. It also implements the
interaction between a Chess Board and custom agents like GameAgent, RandomPlayerAgent, AutoReplyAgent,
and others for simulating different player types.
- Komodo Dragon - the chess engine doing the actual best move calculation, used as a
reference to demonstrate what a real chess player's performance is.
Workflow:
1. The game is initialized with a chess board and two players: a Random Player (as white) and an LLM (as
black).
2. The game loop runs until a termination condition is met, such as checkmate, stalemate, or reaching
the maximum number of moves.
3. Each player takes turns making a move. The Random Player selects a move randomly from the list of
legal moves.
4. The LLM is prompted to make a move using a structured dialog, which includes actions like getting the
current board state, retrieving legal moves, and making a move.
5. The game state is updated after each move, and the board is visualized if enabled.
6. Game statistics are generated and stored at the end of the game.
Dialog:
Here is an example of a dialog that prompts a model to make a move. UCI format refers to Universal Chess
Interface notation, which uses source and destination squares (e.g., e2e4 means move from e2 to e4):
Proxy (to Player_Black):
You are a professional chess player and you play as black. Now is your turn to make a move. Before making a move you can pick one of the following actions:
- 'get_current_board' to get the schema and current status of the board
- 'get_legal_moves' to get a UCI formatted list of available moves
- 'make_move <UCI formatted move>' when you are ready to complete your turn (e.g., 'make_move e2e4')
Respond with the action.
Player_Black (to Proxy):
get_current_board
Proxy (to Player_Black):
♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
· · · · · · · ·
· · · · · · · ·
· · · · · · · ·
· · · · · · · ·
♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
Player_Black (to Proxy):
get_legal_moves
Proxy (to Player_Black):
a7a6, b7b6, c7c6, d7d6, e7e6, f7f6, g7g6, h7h6, a7a5, b7b5, c7c5, d7d5, e7e5, f7f5, g7g5, h7h5
Player_Black (to Proxy):
make_move e7e5
Proxy (to Player_Black):
Move made, switching player
1 move = 1 dialog. The dialog is limited to 10 turns (10 pairs of Proxy/Agent Request/Response
messages). A maximum of 3 mistakes (not conforming to
communication notation, picking wrong action or wrong move) are allowed per dialog.
The game is terminated and LLM is given a loss if the max turns OR number of LLM mistakes is reached in
the given dialog. Models that excel at both chess strategy and following instructions precisely will
perform best.
