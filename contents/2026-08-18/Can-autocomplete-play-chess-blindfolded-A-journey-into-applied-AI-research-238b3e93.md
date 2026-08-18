---
source: "https://alfredvc.no/blog/chess-as-autocomplete"
hn_url: "https://news.ycombinator.com/item?id=49344519"
title: "Can autocomplete play chess blindfolded? A journey into applied AI research"
article_title: "Chess as autocomplete | Alfredo"
image: "https://alfredvc.no/assets/img/chess-autocomplete-card.png"
author: "alfredvc"
captured_at: "2026-08-18T12:24:50Z"
capture_tool: "hn-digest"
hn_id: 49344519
score: 1
comments: 1
posted_at: "2026-08-18T12:09:44Z"
tags:
  - hacker-news
  - translated
---

# Can autocomplete play chess blindfolded? A journey into applied AI research

- HN: [49344519](https://news.ycombinator.com/item?id=49344519)
- Source: [alfredvc.no](https://alfredvc.no/blog/chess-as-autocomplete)
- Score: 1
- Comments: 1
- Posted: 2026-08-18T12:09:44Z

## Translation

タイトル: オートコンプリートは目隠しをしてもチェスをプレイできますか?応用 AI 研究への旅
記事のタイトル: オートコンプリートとしてのチェス |アルフレド
説明: 現代の LLM のトレーニングに使用されるのと同じレシピを使用して、目隠しチェスをプレイするためにトランスフォーマーをゼロからトレーニングします。

記事本文:
-->
avc。
ホーム
について
ブログ
プロジェクト
出版物
インタラクティブなグラフ
オートコンプリートとしてのチェス
オートコンプリートは目隠しをしてもチェスをプレイできますか?応用 AI 研究への旅
ブラウザで直接最終モデルと対戦してください。
ベースライン、どこかから始めないといけない
では、オートコンプリートは目隠しをしたままチェスをプレイできるのでしょうか?
最先端技術と比較してどうですか?
頂上のあの板？フルモデルです
91M パラメータ トランスフォーマーはチェスの動き (およびプレーヤーの評価と時間制御を指定する 3 つのトークン) のみを認識し、ボードやルールは決して認識しません。それは純粋に次のトークンを予測するのが上手になることによって、人間の動きの最先端に近い精度で目隠しをしてプレイすることを学びます。
前回の記事 では、LLM のメンタル モデルをオートコンプリートとして構築しました。
簡単に要約すると、LLM には次のような判決が下されます。
そして次のトークンを単純に予測します。
したがって、理論的には、チェスのゲームを言語のオートコンプリートの問題として組み立てれば、チェスをプレイするように LLM をトレーニングできるはずです。
チェスの言語を自動補完することで、目隠しをしてチェスをプレイするようにトランスフォーマーを訓練できるでしょうか?
目隠しチェスは、プレイヤーがボードを見ないチェスの変形です。誰かが手を読み上げると、彼らは自分の次の手を大声で宣言して答えます。
これは、私たちが望んでいるフレーミングとまったく同じように聞こえます。つまり、移動し、移動します。
チェスのゲームを表すために最も一般的に使用される言語は、Standard Algebraic Notation (SAN) と呼ばれます。たとえば、イタリアのゲームは動きから始まります。
そして、次のように理解できます。
SAN について詳しく知りたい場合は、chess.com の記事を参照してください。
…計算を利用する一般的な方法が、最終的には最も効果的であり、大差を付けています。
このプロジェクトはそのアイデアを極限まで突き詰めたものです。 Karvonen (2024) は、移動テキストだけで訓練された小さな GPT がボードを学習することを示しました。

1,600 万のゲームと 2,500 万のパラメータで、調査するには十分です。私の知る限り、そのアイデアを大規模に実行した人は誰もいません。人間のゲームの膨大なコーパスで訓練された一般的なモデル (チェス特有の要素が何も追加されていない、ボードも探索も余分なヘッドも何もない標準モデル) です。 1 つのトランスフォーマーは、そのコーパスの約 100 倍にあたる、約 18 億回の人間のゲームで訓練されました。
ゲームの次の動きを予測することは、LLM が次の単語を予測して解決する問題と同じであるため、私は装置全体を借用しました。つまり、最新の LLM が構築されるレシピでトレーニングされたデコーダー専用トランスフォーマー (ChatGPT の背後にあるモデル ファミリー) です。インターフェイスはトークンの入力、トークンの出力のみであり、それ以外は何もありません。モデルは一連の動きを見て、それを継続することを学習するだけです。
モデルは、すべての駒がどこに置かれているか、誰の順番であるか、どの手が正当であるかなど、その他すべてを計算する方法を学習する必要があります。移動ストリームのみからすべてを再構築します。それは文字通りの意味での目隠しチェスです。
他にも人間のようなチェスを機械に教えた人はたくさんいます。チェスがアーキテクチャにどの程度組み込まれているか、移動時に検索するかどうか、Stockfish のような強力なエンジンから学習するかどうか、およびトレーニングに使用するデータの量が異なります。このプロジェクトは、最初の 3 つ (ボードなし、検索なし、エンジン教師なし) のうちの予備に位置し、最後のものであるデータに頼っています。
ポジション = トレーニングされた 1 つのボードの状態。ムーブストリーム LM の場合、プライごとに 1 つのムーブ トークン (Allie、この投稿)。時間制御 = トレーニング データ内の Lichess ゲーム速度。 「すべて 4」 = ブレット、ブリッツ、ラピッド、クラシック。
マイア: モデルごとに ≈0.6B、評価ビンごとに 1 つ (×9)。 1,200 万ゲーム/モデル × ~50 キープムーブ/ゲーム;モデルでは 409.6M (400k ステップ × 1024 バッチ) が見られます。
Maia-3: コーパス サイズは未公開。トレーニングで処理された約 5 億 1,200 万の位置 (100 万ステップ × 512 の有効バッチ)、sca として表示

ルプロキシ。
グランドマスター: ≈0.53B の個別のポジションにわたる 1530 億個の Stockfish アクション値ラベル。
アリー: 66億トークンのトレーニング セット (9,100 万ゲーム) が約 40 エポック (200 万ステップ × 131,072 トークン) トレーニングされたため、約 262 億トークンが見られました。 6.6B はデータセット、262B はトレーニング中に確認された合計です。
ボードベースの人間模倣者: マイア ライン。 Maia (McIlroy-Young et al., 2020) は、チェス AI の目標を、最良の手をプレイすることから、特定のレーティングの人間が実際にプレイする手を予測することへと再構成しました。これは、ボードを 8 × 8 プレーンのスタックとして読み取る AlphaZero スタイルの畳み込みネットワークを使用し (そのため、ボード、そのジオメトリ、正当な動きが学習されるのではなく組み込まれます)、検索なしでプレイします。問題は、「特定の評価」とは、100-Elo バンドごとに 1 つずつ、合計 9 つの個別のモデルを意味するということです。 Maia-2 (Tang et al., 2024) は、両方のプレイヤーのレーティングを条件とする「スキルを意識した注意」モジュールを使用して、これらを 1 つのネットワークにまとめました。Maia-3 / Chessformer (Monroe et al., 2026) は、そのアイデアを最も推し進めました。変換器は、学習された幾何学的な注意バイアスを持つ 64 個の正方形のトークンとしてボードを読み取り、現在、ボードベースのモデル間で人間の動きの一致記録を保持しています。 (7,900 万パラメータで 57.1%)。
エンジン蒸留の外れ値。 DeepMind の検索なしのグランドマスターレベルのチェス (Ruoss et al., 2024) は、奇妙なものです。これはサーチレストランスフォーマーでもありますが、人間からはまったく学習しません。ボード (トークン化された FEN) を読み取り、人間のゲームをポジションのソースとしてのみ使用して、150 億回程度のエンジン評価ですべての正当な手に対する Stockfish 16 の勝利確率を再現するように訓練されています。結果は検索せずにグランドマスターの強さ (~2895 Lichess blitz Elo) で再生されます。それは人々を模倣するのではなく、エンジンを蒸留します。その目標は強さであり、人間らしさではありません。
ｗ

旧モデル プローブ: Chess-GPT。 Adam Karvonen (2024) は、PGN テキストのみで小さな GPT を 1 文字ずつトレーニングし、その内部を調べました。リニア プローブは、すべての正方形の状態を 99.6% の精度で回復します。ボードが与えられていないモデルでも、とにかくボードが構築されます。 2 番目のプローブでは、Elo 1550 未満のプレイヤーと Elo 2050 を超えるプレイヤーを 90.5% の精度で分離します。そのため、モデルは純粋に次のキャラクターをより正確に予測するための潜在変数としてスキルを推定します。これは、このプロジェクトが明示的にフィードするメタデータ トークンの教師なしバージョンです。これは、オセロ GPT の結果のチェス バージョンです。ここでの違いはスケールです。17 億 6,000 万のゲームに対して、2,500 万から 5,000 万のパラメータで 1,600 万のゲームです。カルボーネンはその効果が存在することを証明している。この投稿では、コーパスが 100 倍大きく、モデルが調査されるのではなく再生されるようにトレーニングされている場合に、その価値は何なのかを問うています。
ムーブストリームのいとこ：アリー。 Allie (Zhang et al., 2024) はこのプロジェクトに最も近いもので、一連の動き (UCI 表記法) に対するデコーダ専用のトランスフォーマーであり、私たちのものと同じ意味でのチェスの言語モデルです。これは「人間のみから」トレーニングし (エンジンの蒸留はありません)、人間が考える時間とゲームがどのように終わるかを予測するヘッドを追加します。その探索不要ポリシーは、このプロジェクトが到達する範囲と同じ範囲で 55.7% の手番一致を記録します。 見出しの違いは、その「ちょっとした探索」です。モンテカルロ ツリー探索の予算は、モデル自身が予測する人間の思考時間に合わせて調整されます。この探索により、アリーは非常に強い相手に対して調整され続けます。このプロジェクトでは検索をまったく行わず、それを放棄しています。
1つではなく、毎回コントロールします。ここの人間の模倣者たちはそれぞれ、リッチの 1 枚のスライスを訓練します。マイア 2 はラピッドで、マイア 3 とアリーはブリッツで、オリジナルのマイアでさえ弾丸を発射します。速いゲームはノイズとして無視される、c

ラッシックは希少すぎる。私たちはオールタイムコントロールでトレーニングします。
他のものではボード、検索、エンジン ティーチャー、または追加のヘッドが追加されますが、このものではデータのみが追加されます。つまり、狭いスライス上で多数のパスを通過するのではなく、毎回コントロールを通過する 1 つのパスです。これが、チェス固有のモデルではなく、上記の意味での一般的なモデルとなる理由です。
これは個人的なプロジェクトであるため、大規模な計算リソースにアクセスすることができません。 Intel i5-13500 CPU、NVIDIA RTX 4070 Ti Super (16 GB) GPU、16 GB の RAM、および 1 TB 未満の利用可能なストレージを搭載したパーソナル コンピューターを持っています。これは、プロジェクト全体で行われるいくつかのエンジニアリング上の選択に影響を与えます。プロジェクトの途中で、予期せず A100 および H100 GPU を搭載した 2 つの VM にアクセスできるようになりましたが、その時までに、それらの選択肢は上記のハードウェアによってすでに形成されていました。
チェスのオートコンプリートを事前トレーニングする目的は、チェスの言語、つまりゲームのルールを学習することです。次の 2 つの指標を使用して、事前トレーニングされたモデルのパフォーマンスを測定します。
有効な手確率質量 (VPM): モデルがポジション内のすべての正当な手に対して割り当てる確率の合計 (すべてのポジションの平均)。簡単に言うと、モデルの確率のどの部分が実際に合法な動きになるのかということです。この定義における正当な手には、駒の動きまたは有効な引き分け請求 (3 回の繰り返し) が含まれます。チェックメイトのような最終状態を誤って予測することも、無効な手とみなされます。ルールを完全に学習したモデルのスコアは 100% になります。
ゲーム終了予測精度 (GPA): ゲームの終了方法 (チェックメイト、ステイルメイト、反復による引き分け、50 手のルール) を予測するときのモデルの精度を測定します。モデルがゲームの終了時期と終了方法を正確に予測する頻度。
移動の有効性に関して重要な数値は、有効な移動確率の質量と

トップ1の法定金利ではありません。これは、次の手を選択するために出力分布からサンプリングする可能性があり、このメトリックは不正な手をサンプリングする確率を正確に示すためです。
ただし、ゲーム終了予測の場合、正しいメトリックは、最も可能性の高い出力が正しいゲーム終了予測である頻度です。モデルの出力がゲーム終了の予測である場合は、決してサンプリングしません。私たちは常に最も可能性の高いものを選択します。だからこそ、ここでは確率の質量ではなく、精度が適切な指標となるのです。
これら 2 つの指標は、モデルがチェスのルールをどの程度学習したかを評価する方法を提供します。私たちの目標は、100% の VPM と 100% の GPA に達することです。
また、さまざまな強さレベルで Stockfish 18 チェス エンジンと対戦することで、モデルのプレイの強さの感覚を得ることができます。これは明示的な最適化目標ではありません。このプロジェクトの目標は、最高評価のチェス モデルを構築することではなく、純粋なオートコンプリートとしてトレーニングされたモデルが最終的にどの程度うまくプレイできるかを理解することだけです。
オープンソースのオンライン プラットフォーム Lichess は、プラットフォーム上でプレイされるゲームをオープン データベースとして公開します。このデータベースには、完全な初心者からグランドマスターまで、13 年以上にわたる合計 7,863,012,346 の評価済みチェス ゲームが含まれています (2026 年 5 月現在)。このデータベースの圧縮後の合計サイズは 2.49 TB です。
Elo は約 1500 でピークに達します。これは、プレイヤーが Lichess で 1500 Elo から開始するためである可能性があります。最も人気のある時間コントロールはブリッツとバレットで、次にラピッド、ウルトラバレット、クラシックが続きます。
Lichess オープン データベースには約 78 億のゲームが保存されていますが、最終モデルではそのほとんどが表示されません。 2024 年 1 月から 2025 年 9 月までの 21 の月間標準レーティング ファイル、1,950,403,943 ゲームでトレーニングされます。検証セットは、最新のファイル (2025 年 9 月) からのゲームの最新 10% として保持されます。以前のすべて

それがトレーニングデータです。
ファイルは新しいものから順に読み取られ、プールが終了する前に 600,000 ステップのバジェットが使い果たされます。この実行は 2024 年 2 月のファイルの 13% で終了し、2024 年 1 月には開始されないため、プールの 90% に相当する 17 億 6,000 万のゲームがほぼすべて 1 回だけ表示されます。
Lichess は、他のプラットフォームとは異なる方法でいくつかのルールを実装します。 3 回の繰り返しは自動的に引き分けされるわけではありません。移動する側はそれを主張する必要があり、多くのプレーヤーは単にプレイを続けるだけです。そのため、データは定期的に誰も主張していない繰り返しに位置を保持しています (Lichess は 5 回目の繰り返しでのみ引き分けを強制します)。 50 手のルールはその逆です。Lichess は適用された瞬間に自動的にゲームを引き分けます。そのため、ゲームがそれを超えて続行されることはありません。
Lichess オープン データベース内のゲームは、zstd 圧縮 PGN (ポータブル ゲーム ノテーション) ファイルとして保存されます。 PGN ファイルは、1 つまたは複数のチェスのゲームのメタデータと手の両方を保存するために一般的に使用されるテキスト形式です。以下はチェス ゲームの PGN の例です。
【イベント「トロールマスターズ」】
[サイト「ガウスダルNOR」]
[日付「2001.01.05」]
[ ラウンド「1」 ]
[ ホワイト「エドヴァルドセン,R」 ]
[ ブラック「カールセン、マグナス」 ]
[結果「1/2-1/2」]
[ ホワイトエロ「2055」 ]
[ ブラックエロ "" ]
[ エコ「D12」 ]
1. d4 Nf6 2. Nf3 d5 3.

[切り捨てられた]

## Original Extract

Training a transformer from scratch to play blindfold chess, with the same recipe used to train modern LLMs.

-->
avc.
Home
About
Blog
Projects
Publications
Interactive Graph
Chess as autocomplete
Can autocomplete play chess blindfolded? A journey into applied AI research
Play against the final model directly in your browser.
A baseline, we have to start somewhere
So, can autocomplete play chess blindfolded?
How does it compare to the state of the art?
That board at the top? It's the full model
A 91M parameter transformer sees only chess moves (plus three tokens naming the players' ratings and the time control), never the board, never the rules. It learns to play blindfolded at near-state-of-the-art human move accuracy, purely by getting good at predicting the next token.
In my previous article , we built a mental model of LLMs as autocomplete.
As a quick recap: LLMs are given the sentence
and simply predict the next token.
So in theory, if we frame the game of chess as a language autocomplete problem, we should be able to train an LLM to play chess.
Can we train a transformer to play chess blindfolded by autocompleting the language of chess?
Blindfold chess is a variant of chess where the players do not see the board. Someone reads out the moves to them, and they reply by announcing their own next move out loud.
This sounds exactly like the framing we want: moves in, moves out.
The most commonly used language to represent chess games is called Standard Algebraic Notation (SAN). For example, the Italian Game starts with the moves
And can be understood as follows:
If you wish to explore SAN in detail you can go to this chess.com article .
…general methods that leverage computation are ultimately the most effective, and by a large margin.
This project takes that idea to an extreme. Karvonen (2024) showed that a small GPT trained on move text alone learns the board well enough to be probed for it, at 16 million games and 25M parameters. As far as I know nobody has run that idea at scale: a general model (a standard model with nothing chess-specific added: no board, no search, no extra heads) trained on a massive corpus of human games. One transformer, trained on nearly 1.8 billion human games, about a hundred times that corpus.
Predicting the next move in a game is the same problem an LLM solves predicting the next word, so I borrow the whole apparatus: a decoder-only transformer (the model family behind ChatGPT) trained with the recipe modern LLMs are built with. The interface is tokens in, tokens out, nothing else: the model only ever sees a stream of moves and learns to continue it.
The model must learn to calculate everything else: where every piece sits, whose turn it is, which moves are even legal. It reconstructs all of it from the move stream alone. That is blindfold chess in the most literal sense.
Plenty of others have taught machines human-like chess. They differ in how much chess is built into the architecture, whether they search at move time, whether they learn from a strong engine like Stockfish, and how much data they train on. This project sits at the spare end of the first three (no board, no search, no engine teacher) and leans on the last: data.
Position = one board state trained on; one move token per ply for the move-stream LMs (Allie, this post). Time controls = the Lichess game speeds in the training data; "all four" = bullet, blitz, rapid, classical.
Maia: ≈0.6B per model, one per rating bin (×9); 12M games/model × ~50 kept moves/game; model sees 409.6M (400k steps × 1024 batch).
Maia-3: corpus size unpublished; ≈512M positions processed in training (1M steps × 512 effective batch), shown as a scale proxy.
Grandmaster: 15.3B Stockfish action-value labels over ≈0.53B distinct positions.
Allie: 6.6B-token training set (91M games) trained ~40 epochs (2M steps × 131,072 tokens), so ≈262B tokens seen; the 6.6B is the dataset, the 262B the total seen during training.
Board-based human imitators: the Maia line. Maia (McIlroy-Young et al., 2020) reframed the goal of chess AI from playing the best move to predicting the move a human of a given rating would actually play. It uses an AlphaZero-style convolutional network that reads the board as a stack of 8×8 planes (so the board, its geometry, and the legal moves are built in rather than learned), and it plays with no search. The catch is that "a given rating" meant nine separate models, one per 100-Elo band. Maia-2 (Tang et al., 2024) folded those into one network with a "skill-aware attention" module that conditions on both players' ratings, and Maia-3 / Chessformer (Monroe et al., 2026) pushed the idea furthest: a transformer that still reads the board, now as 64 square-tokens with a learned geometric attention bias, and currently holds the human move match record among board-based models (57.1% at 79M parameters).
The engine-distillation outlier. DeepMind's Grandmaster-Level Chess Without Search (Ruoss et al., 2024) is the odd one out. It is also a searchless transformer, but it does not learn from humans at all: it reads a board (a tokenized FEN) and is trained to reproduce Stockfish 16's win probability for every legal move, on the order of 15 billion engine evaluations, with human games used only as a source of positions. The result plays at grandmaster strength (~2895 Lichess blitz Elo) without ever searching. It distills an engine rather than imitating people; its goal is strength, not human-likeness.
The world-model probes: Chess-GPT. Adam Karvonen (2024) trained small GPTs on nothing but PGN text, character by character, then went looking inside them. Linear probes recover the state of every square to 99.6% accuracy: a model given no board builds one anyway. A second probe separates players below 1550 Elo from those above 2050 with 90.5% accuracy, so the model estimates skill as a latent variable purely to predict the next character better, which is the unsupervised version of the metadata tokens this project feeds in explicitly. It is the chess version of the Othello-GPT result. What differs here is scale: 16 million games at 25M to 50M parameters, against 1.76 billion games. Karvonen establishes that the effect exists; this post asks what it is worth when the corpus is a hundred times larger and the model is trained to be played rather than probed.
The move-stream cousin: Allie. Allie (Zhang et al., 2024) is the closest relative to this project: a decoder-only transformer over a stream of moves (in UCI notation), a language model of chess in the same sense ours is. It trains "exclusively from humans" (no engine distillation) and adds heads that predict how long a human would think and how the game will end. Its search-free policy scores 55.7% move match, in the same range this project lands in. The headline difference is its "bit of search": a Monte-Carlo tree search whose budget scales with the model's own predicted human think-time. That search keeps Allie calibrated against very strong opponents; this project, with no search at all, gives that up.
Every time control, not one. The human imitators here each train on a single slice of Lichess: Maia-2 on rapid, Maia-3 and Allie on blitz, and even the original Maia throws out bullet. Fast games get dismissed as noise, classical as too scarce. We train on all time controls.
Where the others add a board, a search, an engine teacher, or extra heads, this one adds only data: a single pass across every time control rather than many passes over a narrow slice. That is what makes it a general model in the sense above, rather than a chess-specific one.
This being a personal project, I do not have access to large computational resources. I have a personal computer with an Intel i5-13500 CPU, an NVIDIA RTX 4070 Ti Super (16 GB) GPU, 16 GB of RAM, and under 1 TB of available storage. This will impact several engineering choices made throughout the project. Partway through the project I unexpectedly gained access to a couple of VMs with A100 and H100 GPUs, but by then those choices had already been shaped by the hardware above.
The goal of pre-training our chess autocomplete is to learn the language of chess , the rules of the game. We will measure the performance of a pre-trained model with two metrics:
Valid move probability mass (VPM): the sum of the probabilities the model assigns to all legal moves in a position, averaged over all positions. In plain terms: what fraction of the model's probability lands on moves that are actually legal. A legal move in this definition includes a piece move or a valid draw claim (threefold repetition). Predicting a terminal state like checkmate incorrectly is also considered an invalid move. A model that has fully learned the rules would score 100%.
Game-end prediction accuracy (GPA): measures the accuracy of the model when predicting how a game ended (checkmate, stalemate, draw by repetition, fifty move rule). How often does the model correctly predict when and how the game ends.
For move validity, the important number is the valid move probability mass and not the top-1-legal rate. This is because we may sample from the output distribution to choose the next move, and this metric correctly shows the probability of sampling an illegal move.
For game end prediction however, the correct metric is how often the most likely output is the correct game end prediction. Whenever the model output is a game end prediction, we never sample; we always pick the most likely one. That is why accuracy is the right metric here, not probability mass.
These two metrics will give us a way to evaluate how well the model has learned the rules of chess. Our goal is to reach 100% VPM and 100% GPA.
We will also get a sense of the model's playing strength by playing it against the Stockfish 18 chess engine at a range of strength levels. This is not an explicit optimization target: the goal of this project is not to build the highest-rated chess model, only to understand how well a model trained as pure autocomplete ends up playing.
The open-source online platform Lichess publishes the games played on their platform as an open database . This database contains (as of May 2026) a total of 7,863,012,346 rated chess games, spanning over 13 years, from complete beginners to grandmasters. This database has a total compressed size of 2.49 TB.
Elo peaks at around 1500, which may be because players start with 1500 Elo on Lichess. The most popular time controls are blitz and bullet, followed by rapid, ultrabullet, and classical.
The Lichess open database holds about 7.8 billion games, but the final model never sees most of them. It trains on the 21 monthly standard-rated files spanning January 2024 through September 2025, 1,950,403,943 games. The validation set is held out as the most recent 10% of games from the newest file (September 2025); everything before that is training data.
The files are read newest first, and the 600,000-step budget runs out before the pool does. The run ends 13% into the February 2024 file and never opens January 2024, so it sees 1.76 billion of those games, 90% of the pool, almost all of them exactly once.
Lichess implements a couple of rules differently than other platforms. A threefold repetition is not an automatic draw: the side to move has to claim it , and many players simply play on, so the data routinely holds positions sitting on a repetition nobody claimed (Lichess only forces the draw on the fifth repetition). The fifty move rule is the opposite: Lichess draws the game automatically the moment it applies, so no game ever continues past it.
The games in the Lichess open database are stored as zstd compressed PGN (portable game notation) files. A PGN file is a commonly used text format for storing both metadata and moves from one or more games of chess. Below is an example of the PGN of a chess game.
[ Event "Troll Masters" ]
[ Site "Gausdal NOR" ]
[ Date "2001.01.05" ]
[ Round "1" ]
[ White "Edvardsen,R" ]
[ Black "Carlsen,Magnus" ]
[ Result "1/2-1/2" ]
[ WhiteElo "2055" ]
[ BlackElo "" ]
[ ECO "D12" ]
1. d4 Nf6 2. Nf3 d5 3.

[truncated]
