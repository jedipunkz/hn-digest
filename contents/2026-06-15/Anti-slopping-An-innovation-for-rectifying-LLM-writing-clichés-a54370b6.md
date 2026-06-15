---
source: "https://research.thoughtworks.com/library/anti-slopping-an-innovation-for-rectifying-llm-writing-cliches"
hn_url: "https://news.ycombinator.com/item?id=48544760"
title: "Anti-slopping: An innovation for rectifying LLM writing clichés"
article_title: "Anti-slopping — An innovation for rectifying LLM writing clichés | Thoughtworks Research"
author: "freeatnet"
captured_at: "2026-06-15T19:25:51Z"
capture_tool: "hn-digest"
hn_id: 48544760
score: 2
comments: 0
posted_at: "2026-06-15T17:50:37Z"
tags:
  - hacker-news
  - translated
---

# Anti-slopping: An innovation for rectifying LLM writing clichés

- HN: [48544760](https://news.ycombinator.com/item?id=48544760)
- Source: [research.thoughtworks.com](https://research.thoughtworks.com/library/anti-slopping-an-innovation-for-rectifying-llm-writing-cliches)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T17:50:37Z

## Translation

タイトル: アンチスロッピング: LLM の書き方の常套句を修正するためのイノベーション
記事のタイトル: アンチスロッピング — LLM の書き方の常套句を修正するためのイノベーション |思考回路の研究
説明: 滑り止め

記事本文:
チーム
お問い合わせ
戻る
アンチスロッピング: LLM の書き方の常套句を修正するためのイノベーション
今日の LLM は「スロップ」に悩まされています。これは、テキストの品質が劣るだけでなく、認識可能で当たり障りのない、反復的で予測可能なロボット的なテキスト出力と定義されます。このようなパターンを抑制することは明らかな解決策のように見えますが、そうではありません。他の有用な単語を無効にして出力をさらに損傷します。 「消化不良」という単語を禁止すると、「in」や「digest」などの単語も禁止され、テキストが無意味になります。抑制は「逆効果」を引き起こす可能性もあります。プロンプト中に用語や概念を禁止すると、モデルがそれについてさらに話すことによって誤って優先順位を付ける可能性があります。
この投稿 (独自の調査に基づく) では、出力から 90% の「スロップ削減」を達成することで、この問題の解決策を紹介します。私たちは、使いすぎたパターンを検出して置き換えるための 3 つのイノベーションを組み合わせたフレームワークを提供しました。このフレームワークは、出力を劣化させることなく 8000 のパターンを削除できます。フレームワークには以下が含まれます。
反復的な単語に対してバックトラッキングを使用し、モデルに新しい、より人間らしい単語を選択させるアンチスロップ サンプラー。
モデル出力内の単語とフレーズの頻度比率と人間のベースラインを計算することにより、「スロップ」を特定する自動パイプライン。 2 つのベースラインを使用しました。 1 つ目は wordfreq ライブラリで、2 つ目は Reddit と Project Gutenberg の人間が書いたテキスト コーパスです。
最終トークン優先最適化 (FTPO) は、モデルが「いい加減な」単語を選択した正確なトークンを調べるトレーニング アルゴリズムです。その後、途中で多くの「ソフトタッチ」メカニズムを実装することで、特定のロジットのみを調整します。
以下は、反復的なパターンや遅延を減らすことができる現在の戦略のリストです。

op.ただし、それぞれの戦略には独自の特徴がある一方で、それぞれに欠点もあります。以下に表を示します。
一貫した出力における繰り返しの傾向に対処しないでください。
高確率トークンのみを対象とします。
統計的に出現した反復パターンを識別できない。
ExLlamaの文字列禁止機能
推論時に、指定された文字列のセットをハード禁止します。
ビームプルーニングによって禁止された単語や語句を除外します。
直接優先最適化 (DPO)
優先応答の可能性が低下し、多様性の崩壊が引き起こされ、出力の構文および N グラムの多様性が減少します。
すべてのモデルは異なり、いくつかの固有の傾向があります。私たちはそれらが何であるかをさらに詳しく調べたいと思いました。これを行うために、頻度比を使用して傾斜パターンを生成する原因となる、過度に使用されている単語と N グラムを収集しました。 n グラムは、テキストからの n 個の連続した項目 (通常は単語、場合によっては文字) のシーケンスです。周波数比は次の式に基づいて計算されました。
\[ \rho(\mathbf{p}) = \frac{f_{\text{LLM}}(\mathbf{p})}{f_{\text{human}}(\mathbf{p})} \]
\[ f_{\text{LLM}}(\mathbf{p}) \text{ - パターンの頻度 } (\mathbf{p}) \text{ LLM 内} \]
\[ f_{\text{human}}(\mathbf{p}) \text{ - 人間の体におけるパターン } (\mathbf{p}) \text{} の頻度 \]
私たちは Reddit からのクリエイティブ ライティング プロンプトを使用して 2,000 を超える出力を生成し、さまざまな言語モデルでさまざまな過剰表現を発見しました。
LLM のサンプラーの仕事は、確率分布から次のトークンを選択することです。サンプラーは、モデルの出力の創造性、多様性、一貫性を制御します。
私たちの革新的なサンプラーは、単語のシーケンス全体が推論トレースに現れた後にのみトリガーされます。次のように動作します。
モデルが特定の入力に対する出力を生成するとき、アルゴリズム k

トークンとロジット分布の追跡を記録します。各トークンの後で不要なパターンがスキャンされます。このようなパターンが検出されると、それを抑制する代わりに、アルゴリズムはその原点に戻り、次のように定義される構成可能な禁止強度パラメータ「s」によって開始確率を下げます。
\[ \mathbf{p}_{\text{new}} = \mathbf{p}_{\text{old}} \times 10^{-10s} \]
\[ \text{ここで、 } 0 \leq s \leq 1.0 \]
次のステップでは、min-p フィルタリングを使用して、調整された分布を制限します。このステップでは、事前定義された確率しきい値を満たす一貫した候補が選択されます。私たちのアンチスロップバックトラッキングアルゴリズムは次のとおりです。
図:2 アンチスロップバックトラッキングアルゴリズム
最後のステップはソフト禁止です。ここでは、高い確率分布を持つフォワードパスを通過する禁止されたパターンのみを許可します。このメカニズムは、禁止強度パラメータ「s」を通じて増分制御を提供します。 s = 0 の場合、パターンは自由に許可されます。 0 から 1 までの値は禁止リストの増分抑制を提供しますが、s = 1 は完全な抑制を強制します。
図 3: ソフトバンニング手法の例
アンチスロップ バックトラッキング メカニズムは、推論フェーズで不要なパターンを検出し、その禁止されたシーケンスの最初のトークンにバックトラックして、その確率を下げます。その後、再サンプリングされます。
図 4: スロップ防止バックトラッキング プロセス
このサンプラー プロセスの制限は、頻繁なバックトラッキングによりパフォーマンスが 69 ～ 96% 低下することです。 1000 から 8000 までの禁止リストのサイズも、パフォーマンスの低下に重要な役割を果たします。このサイズの禁止リストは、ほとんどの現実世界のモデルにとって過剰になります。バックトラッキングはリソースを大量に消費します。 API 実装で使用されるサンプラー技術では、さらにパフォーマンスの低下が観察される場合があります。
最終トークン優先最適化 (FTPO)
FTPOは研修です

ng アルゴリズムは、アンチスロップ サンプラーで使用されるバンリストとバックトラッキングによるパフォーマンスの損失を回避するように設計されています。サンプラーは不良トークンを検出し、繰り返しバックトラックするため、スループットが大幅に低下します。しかし、根本的な問題は、モデル自体が依然としてこれらの不要なトークンに高い確率を割り当てていることです。
FTPO は、これらのトークンを直接回避するようにモデルをトレーニングすることでこの問題を修正し、バックトラッキングの必要性がほとんどなくなります。 FTPO はシーケンス全体を再トレーニングするのではなく、エラーが発生した正確な決定点でのみモデルを調整し、拒否されたトークンをプッシュダウンして、より良い代替を優先します。これにより、モデル全体の品質へのダメージを最小限に抑えながら、悪いパターンが永続的に内部的に抑制されます。私たちは基本的に、学習された動作によって遅いランタイムを高速なスループットに変えました。
FTPO はロジットで機能します。そのため、トークン配布全体を妨げることなく、正確でローカルな制御を「繊細に」実現します。 (ロジットは、確率に変換する前の、考えられる各トークンの生のスコアです。) FTPO メソッドの重要な部分は、3 つの損失項の計算です。その方法は次のとおりです。
これにより、選択されたトークンが拒否されたトークンよりも一定のマージン (m) だけ上にプッシュされます。
\[ L_{\text{pref}} = \frac{\sum_{c \in C} \omega_c \cdot \text{softplus}(m - \Delta_c)}{\sum_{c \in C} \omega_c} \]
\[ \Delta_c = y[c] - y[r] \text{ 選択されたトークンと拒否されたトークンの間のロジット ギャップです。} \]
\[ \text{} \Delta < 0 \text{ の場合 (負けを選択した場合)、ペナルティは大きくなります。 } \Delta \text{ が } \] に向かって増加するにつれて
\[ \text{margin } m \text{、ペナルティは滑らかに減少します。 } \Delta \geq m \text{ を実行すると、重みはゼロになります} \]
\[ \text{そして優先順位の損失はもはや寄与しません。} \]
\[ \omega_c = \text{クランプ} \left(\frac{m - \Delta_c}{m}, 0, 1\right) \text{ 非アクティブ化

マージンに達したときの損失です。} \]
これは、選択されたトークンと拒否されたトークンに作用します。これにより、事前定義された基準に近い値が維持されます。トークンがペナルティを受ける前に、小さな空き領域 ( \(\tau_{\text{target}}) \) が許可されます。
\[ L_{\text{target}} = \frac{1}{|T|} \sum_{j \in T} \max(|y[j] - y_{\text{ref}}[j]| - \tau_{\text{target}}, 0)^2 \]
\[ \text{where,} \] \[ T = C \cup \{r\} \text{ すべてのターゲット トークンが含まれます} \]
これにより、他のすべてのトークンが参照に固定されます。これにより、語彙の無関係な部分が意図せず変更されるのを防ぎます。
\[ L_{\text{nontarget}} = \frac{1}{|N|} \sum_{j \in N} (y[j] - y_{\text{ref}}[j])^2 \]
\[ \text{どこ,} \] \[ N \text{ はすべて非ターゲット トークンです} \]
\[ L_{\text{FTPO}} = L_{\text{pref}} + \lambda_{\text{target}}L_{\text{target}} + \lambda_{\text{nontarget}}L_{\text{nontarget}} \] \[ \text{Where,} \]
\[ \lambda_{\text{target}} \text{ および } \lambda_{\text{nontarget}} \text{ は重み係数です} \]
FTPO を効果的にするための 3 つの設計原則があります。
ロジット空間操作: FTPO は、MSE 損失を生の「ロジット」(スコア) に適用します。これにより、モデルは語彙の無関係な部分を妨げることなく、特定の「選択された」トークンと「拒否された」トークンのみをターゲットにして変更することができます。
マージンの無効化: FTPO はマージン m を使用します。良いトークンと悪いトークンの間のギャップが十分に大きくなると、重み変数 \(w_c\) は自動的にゼロになります。これにより、その特定のペアのトレーニングが停止され、オーバートレーニングが防止されます。
2 部構成の正則化: FTPO は 2 部構成の MSE 損失を使用します。これにより、残りの語彙を参照に制限しながら、ターゲット ロジットが比較的自由に移動できるようになります。これにより、破壊的なロジットの発散を回避しながら、高い嗜好精度のトレーニングが可能になります。
以下の図は、データをトレーニングするプロセス全体を示しています。

r FTPO:
図 5: アンチスロップサンプラーと FTPO トレーニング
グラフは、4 つの方法に対する FTPO 禁止リスト抑制のパフォーマンスを視覚化します。トークン禁止、アンチスロップ サンプラー、DPO、Gemma3 -12b ベースライン。各メソッドの禁止リスト抑制は、ライティング ルーブリックで測定された出力低下に対してプロットされています。 Writing-Quality Rubric は、生成されたテキストの構造的および文体の完全性を客観的に測定するように設計された自動評価フレームワークです。人間の採点者に依存する代わりに、大容量モデル (GPT-5) を使用して、次の 4 つの特定の次元に基づいて出力を採点します。
FTPO は、品質損失を最小限に抑えながら 90% のスロップ抑制を達成し、DPO やトークン禁止を上回ります。図 6 は、2k、4k、および 8k パターンのバンリスト サイズにわたって gemma-3-12b に対する 4 つの抑制方法を評価しています。 FTPO は、ベースラインの書き込み品質を維持しながら、不要なパターンの 85 ～ 90% を抑制します。対照的に、DPO は 80 ～ 82% の抑制しか達成していないにもかかわらず、品質を 6 ～ 15 ポイント低下させます。トークンの禁止は壊滅的な品質の崩壊を示しています。誤差バーは 95% 信頼区間 (\(CI_{95}\)) を示します。 n= 条件ごとに 1,000 個の出力。
以下の表は、FTPO と DPO の比較を示しています。
FTPO は DPO よりも 8.5% 強力な抑制を実現します。
FTPO は、GSM8k での数学的推論と MMLU での世界知識機能をベースラインの 1 ～ 3% 以内に維持しています。 DPO は両方の指標を 2 ～ 5% 低下させます。
FTPO でトレーニングされたモデルは、2k、4k、および 8k バンリスト サイズのベースライン gemma3 スコアの周囲にクラスタリングします。一方、DPO では品質が大幅に低下します。
FTPO は多様性を維持または強化します (ベースラインの 95 ～ 102%) が、DPO は進行性の崩壊を引き起こします (74 ～ 92%)。
FTPO は最小限の低下でほぼ 100% の優先精度までトレーニングできますが、DPO は 40% しか管理できず、その後は大幅な低下が発生します。
FTPO と DPO の比較

興味深い結果が得られました。
図7。 FTPO はトレーニングがより高い優先精度に進むにつれて書き込み品質を維持しますが、DPO は 40% の精度マークを超えると急激に低下します。この実験では、1,000 アイテムの禁止リストに基づいて gemma-3-12b をトレーニングします。
図 8: FTPO では、(1) MSE 損失条件と、(2) すでに勝利しているトークンと拒否されたトークンのトレーニング シグナルを無効にする早期スイッチオフ機能により、ロジットは基準値に近いままになります。
アンチスロップ サンプラーは、頻繁なバックトラッキングにより、1k ～ 8k のバンリスト サイズに対してスループットを 69 ～ 96% 低下させます。私たちは、モデルの使いすぎた書き込みパターンを自動的にプロファイリングし、トレーニング セットを生成して、これらのパターンを抑制するようにモデルをトレーニングするパイプラインを開発しました。当社の FTPO トレーナーは、分布への最小限の変更で、モデルの使いすぎた書き込み傾向を的を絞った調整を行うように設計されています。
私たちのテスト全体を通じて、FTPO とサンプラーは DPO やロジットベースのトークン禁止よりも高い抑制を達成し、ルーブリックで測定可能な品質の損失は無視できました。コードとデータセットは MIT ライセンスに基づいてリリースされます。
私たちの実験とその結果の詳細については、元の論文を参照してください: Antislop: AComprehensive Framework for identifying and Elirepetitive pattern i

[切り捨てられた]

## Original Extract

Anti slopping

Team
Contact
Back
Anti-slopping: An innovation for rectifying LLM writing clichés
Today’s LLMs suffer from “slop”, which we define as the repetitive, predictable and robotic text output that makes the text not only of inferior quality but also recognizable and bland. Suppressing such patterns appears an obvious solution but it is not. It further damages the output by killing other useful words, e.g. banning a word “indigestible” will also ban the words like “in”, and “digest” making the text meaningless. Suppression may also introduce a “backfire effect” — if we forbid a term or concept while prompting, the model can accidentally prioritize it by talking about it more.
This post (based on original research) presents a solution to this problem, by achieving 90% “slop reduction” from its outputs. We have provided a framework that combines three innovations to detect and replace the overused patterns. The framework can eliminate 8000 patterns without degrading the output. The framework includes,
The anti-slop sampler that uses back-tracking for the repetitive words and forces the model to select a new and more human-like word.
An automated pipeline that identifies "slop" by calculating the frequency ratio of words and phrases in the model output versus human baselines. We have used two baselines; the first is wordfreq library and the second is a human-written text corpus from Reddit and Project Gutenberg.
Final token-preference optimization (FTPO) is a training algorithm that looks at the exact token where the model has chosen a “sloppy” word. It then adjusts only those specific logits by implementing many “soft-touch” mechanisms along the way.
Below is a list of current strategies that can reduce repetitive patterns or slop.However, while each strategy possesses unique features, they all have their own shortcomings. Here is a table:
Don’t address repetitive tendencies in coherent outputs.
Targets only high probability tokens.
Can’t identify statistically emerging repetitive patterns.
String banning feature of ExLlama
Hard-bans a provided set of strings at inference time.
Exclude forbidden words or phrases by beam pruning.
Direct preference optimization (DPO)
Lowers the likelihood of preferred responses, inducing diversity collapse and reducing syntactic and n-gram variety in outputs.
Every model is different, having some inherent tendencies; we wanted to explore what those are in more depth. To do this, we collected overly used words and n-grams that are responsible for producing slop patterns using the frequency ratio. An n-gram is a sequence of n consecutive items (usually words, sometimes characters) from the text. The frequency ratio was calculated based on the following formula:
\[ \rho(\mathbf{p}) = \frac{f_{\text{LLM}}(\mathbf{p})}{f_{\text{human}}(\mathbf{p})} \]
\[ f_{\text{LLM}}(\mathbf{p}) \text{ - frequencies of pattern } (\mathbf{p}) \text{ in LLM} \]
\[ f_{\text{human}}(\mathbf{p}) \text{ - frequencies of pattern } (\mathbf{p}) \text{ in human corpora} \]
We generated more than 2000 outputs using creative writing prompts from Reddit and discovered various overrepresentations in different language models.
The job of a sampler in LLMs is to select the next token from the probability distribution. The sampler controls creativity, diversity and the coherence of the model’s output.
Our innovative sampler triggers only after the entire sequence of words appears in the inference trace. It works as follows:
As the model generates the output for certain input, our algorithm keeps the trace of tokens and logit distributions. It scans for unwanted patterns after each token. Once such a pattern is detected, instead of suppressing it the algorithm backtracks to its origin and lowers the initiating probability by the configurable ban-strength parameter “s” defined as:
\[ \mathbf{p}_{\text{new}} = \mathbf{p}_{\text{old}} \times 10^{-10s} \]
\[ \text{Where, } 0 \leq s \leq 1.0 \]
The next step is to use the min-p filtering to constrain the adjusted distribution. This step selects the coherent candidates who meet the predefined probability threshold. Our anti-slop backtracking algorithm is as follows:
Fig:2 Anti-slop backtracking algorithm
The final step is soft banning, where we allow only those banned patterns through the forward pass having high probability distribution. The mechanism provides incremental control through the ban-streangth parameter “s”. When s = 0, then patterns are allowed freely. Values between zero and one provide incremental suppression of the banlist, while s = one enforces complete suppression.
Fig 3: Example of soft banning technique
The anti-slop backtracking mechanism detects unwanted patterns in the inference phase, backtracks to the first token of that banned sequence and lowers its probability. It then resamples.
Fig 4: The anti-slop backtracking process
The limitation of this sampler process is that performance degrades by 69-96% due to frequent backtracking. The size of the banlist from 1000 to 8000 also plays a crucial role in performance degradation. A banlist of this size would be overkill for most real-world models; backtracking is resource intensive. Further performance losses may be observed with sampler techniques used in API implementation.
Final token preference optimization (FTPO)
FTPO is a training algorithm designed to avoid loss of performance due to banlist and backtracking used in anti-slop sampler. The sampler detects bad tokens and backtracks repeatedly, which drastically reduces throughput. But the root problem is that the model itself still assigns high probability to those unwanted tokens.
FTPO fixes this by training the model to avoid those tokens directly, so the need for backtracking largely disappears. Instead of retraining the entire sequence, FTPO adjusts the model only at the exact decision point where the error occurs, pushing down the rejected token and favoring better alternatives. This results in a permanent, internalized suppression of bad patterns while minimizing damage to overall model quality. We’ve essentially turned a slow runtime into a fast throughput by learned behavior.
FTPO works on logits. That’s why it gives precise, local control “delicately” without disturbing the entire token distribution. (A logit is the raw score for each possible token before converting it into probability.) An important part of the FTPO method is the calculation of three loss terms. Here’s how:
This pushes the chosen token above the rejected one by a certain margin (m):
\[ L_{\text{pref}} = \frac{\sum_{c \in C} \omega_c \cdot \text{softplus}(m - \Delta_c)}{\sum_{c \in C} \omega_c} \]
\[ \Delta_c = y[c] - y[r] \text{ It is the logit gap between chosen and rejected token.} \]
\[ \text{When } \Delta < 0 \text{ (chosen losing), the penalty is large. As } \Delta \text{ increases toward the} \]
\[ \text{margin } m \text{, the penalty smoothly tapers. Once } \Delta \geq m \text{, the weight goes to zero} \]
\[ \text{and the preference loss no longer contributes.} \]
\[ \omega_c = \text{clamp} \left(\frac{m - \Delta_c}{m}, 0, 1\right) \text{ deactivates the loss when the margin is achieved.} \]
This acts on chosen and rejected tokens. It keeps them close to the pre-defined reference. It allows a small free region ( \(\tau_{\text{target}}) \) before the token is penalized.
\[ L_{\text{target}} = \frac{1}{|T|} \sum_{j \in T} \max(|y[j] - y_{\text{ref}}[j]| - \tau_{\text{target}}, 0)^2 \]
\[ \text{where,} \] \[ T = C \cup \{r\} \text{ contains all target tokens} \]
This anchors all other tokens to the reference. It prevents unintended changes in unrelated parts of a vocabulary.
\[ L_{\text{nontarget}} = \frac{1}{|N|} \sum_{j \in N} (y[j] - y_{\text{ref}}[j])^2 \]
\[ \text{Where,} \] \[ N \text{ are all nontarget tokens} \]
\[ L_{\text{FTPO}} = L_{\text{pref}} + \lambda_{\text{target}}L_{\text{target}} + \lambda_{\text{nontarget}}L_{\text{nontarget}} \] \[ \text{Where,} \]
\[ \lambda_{\text{target}} \text{ and } \lambda_{\text{nontarget}} \text{ are weighting coefficients} \]
There are three principles of design that makes FTPO effective.
Logit-space operation: FTPO applies MSE loss to the raw "logits" (scores). This allows the model to target and change only the specific "chosen" and "rejected" tokens without disturbing unrelated parts of the vocabulary.
Margin deactivation: FTPO uses margin m. Once the gap between the good token and the bad token is wide enough, a weight variable \(w_c\) ​automatically drops to zero. This stops the training for that specific pair, thus preventing overtraining.
Two-part regularization: FTPO uses the two-part MSE loss that allows target logits to move relatively freely, while constraining the remaining vocabulary to the reference. This allows training to high preference accuracy while avoiding destructive logit divergences.
The diagram below shows the entire process of training data for FTPO:
Fig. 5: Anti-slop sampler and FTPO training
The graph visualizes the performance of FTPO banlist suppression against four methods; token banning, antislop sampler, DPO and Gemma3 -12b baseline. Banlist suppression for each method is plotted against output degradation as measured by our writing rubric. The Writing-Quality Rubric is an automated evaluation framework designed to objectively measure the structural and stylistic integrity of generated text. Instead of relying on human graders, it uses a high-capacity model (GPT-5) to score outputs based on four specific dimensions:
FTPO achieves 90% slop suppression with minimal quality loss, outperforming DPO and token banning. The figure 6 evaluates four suppression methods on gemma-3-12b across banlist sizes of 2k, 4k, and 8k patterns. FTPO maintains baseline writing quality while suppressing 85-90% of unwanted patterns. In contrast, DPO degrades quality by 6-15 points despite achieving only 80-82% suppression. Token banning shows catastrophic quality collapse. Error bars show 95% confidence intervals (\(CI_{95}\)). n= 1,000 outputs per condition.
The table below provides a comparison between FTPO and DPO:
FTPO achieves 8.5% stronger suppression than DPO.
FTPO maintains math reasoning on GSM8k and world-knowledge capabilities on MMLU within 1-3% of baseline. DPO degrades both metrics by 2-5%.
FTPO-trained models cluster around the baseline gemma3 score for 2k, 4k and 8k banlist sizes; while DPO experiences a large degradation in quality.
FTPO maintains or enhances diversity (95-102% of baseline), while DPO causes progressive collapse (74-92%).
FTPO can train to nearly 100% preference accuracy with minimal degradation, while DPO only manages 40%, after which substantial degradation.
Comparing FTPO and DPO provides us with some interesting results.
Fig.7. FTPO maintains writing quality as training progresses to higher preference accuracies, while DPO degrades sharply after the 40% accuracy mark. This experiment trains gemma-3-12b on a banlist of 1,000 items.
Fig 8: With FTPO, logits stay close to reference due to (1) the MSE loss terms and (2) the early switch-off feature which nulls the training signal for chosen tokens that are already winning vs rejected.
The anti-slop sampler reduces throughput by 69-96% for the banlist sizes of 1k to 8k due to frequent backtracking. We developed a pipeline that automatically profiles a model’s overused writing patterns, generates a training set and trains the model to suppress these patterns. Our FTPO trainer is designed to make targeted adjustments to the model’s over-used writing tendencies with minimal changes to its distribution.
Across our tests, FTPO and the sampler achieved higher suppression than DPO and logit-based token banning, with negligible measurable quality loss on our rubric. We release code and datasets under the MIT license.
For more detailed information on our experiment and its results please refer to the original paper: Antislop: A comprehensive framework for identifying and eliminating repetitive patterns i

[truncated]
