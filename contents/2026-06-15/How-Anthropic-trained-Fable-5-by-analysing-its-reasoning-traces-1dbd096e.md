---
source: "https://ankitmaloo.com/fable/"
hn_url: "https://news.ycombinator.com/item?id=48544097"
title: "How Anthropic trained Fable 5 => by analysing its reasoning traces"
article_title: "Fable’s approach analysis | Ankit Maloo"
author: "ankit219"
captured_at: "2026-06-15T19:26:59Z"
capture_tool: "hn-digest"
hn_id: 48544097
score: 6
comments: 0
posted_at: "2026-06-15T17:00:53Z"
tags:
  - hacker-news
  - translated
---

# How Anthropic trained Fable 5 => by analysing its reasoning traces

- HN: [48544097](https://news.ycombinator.com/item?id=48544097)
- Source: [ankitmaloo.com](https://ankitmaloo.com/fable/)
- Score: 6
- Comments: 0
- Posted: 2026-06-15T17:00:53Z

## Translation

タイトル: Anthropic が推論トレースを分析することで Fable 5 をどのように訓練したか =>
記事タイトル: Fable のアプローチ分析 |アンキット・マルー
説明: AI と RL の世界における私の旅を記録します。

記事本文:
アンキット・マルー
ライティングの本棚
Fable のアプローチ分析
概要: 現在、LLMS は SFT 後の複数ステップのプロセスでトレーニングされています。 RL -> 高品質の合成データを生成 → それを自己蒸留 → 別のラウンドの RL (簡略化)。 Fable-5 のソリューション戦略はコードの構成方法に制約があり、貪欲な選択肢をすべて使い果たすまで、よりシンプルなソリューションを模索するのに苦労しました。これは、自己蒸留レシピで生成されるものと一致します。
Fable-5 を取り巻くすべての誇大宣伝を考慮して、私はこれを試してみることにしました。トレーニング方法の違いと、さまざまな評価で優れている理由を理解しようとしました。
簡単な数学の問題を与えて、それがどうなるかを見てみましょう。クロード コードでは思考を確認する機能が削除されていたため、claude.ai Web を使用しました。問題は非常に単純で、6 つの数字と 5 つのステップがあり、出力に到達する必要があります。完全な問題と Fable の解決策は、ここで直接確認できます: https://gist.github.com/ankitmaloo/c491e8a6e4f96b4e5d11b1f2826297dc
Mythos はサイバーセキュリティの悪用に非常に優れており、Anthropic がそのようなタスクに関してモデルを明示的にトレーニングしたことはないと言われました。この投稿とその後のブログは、その理由を理解するのに役立ちました。私の感覚では、モデルはプリミティブを連結するのが非常に上手でしたが、どの程度、どのように連結するかはまだわかりませんでした。さて、上記の要点の解決策は、私が思っていたよりも明確です。
あなたはsummleというゲームをプレイしています。それが何か知っていますか？ Wordle に似ていますが、数字が含まれています。 6 つの数字が与えられ、標準的な数学演算を使用して最終的な数字に到達する必要があります。
ルール。
- 下のタイルを使用して合計を計算し、5 ステップ以内で上の目標数値に達します。
- 許可される演算: +、-、x、/ (除算)
- 正の整数のみが許可されます。
- 1 つの番号を 1 回だけ使用できます。
- ou を使用できます

操作も一度実行します。
今日の数字: 1,1,6,12,50,100 出力された数字: 397
コードは使用しないでください。
ここをトレースします
トレーニング後のトレースからわかること
Fable 5、ノーコード、Summle パズルを解くよう求められます (5 回以下の操作で 1、1、6、12、50、100 から 397 に到達する)。約 60,000 トークンの貪欲な深さ優先検索に苦労しましたが、体系的なルート分割列挙に切り替えてから数秒以内に解決しました。
トレースは、問題にアプローチするときにモデルがどのような動作をするように条件付けされているかという意味で興味深いものです。
注意: このメモは、なぜ闘争が起こったのか、そしてそれがモデルのトレーニング後の方法について何を意味するのかについての事後分析です。分析は行動から推測され、トレーニング レシピに関する内部知識はありません。
サイバーチェーンと数字のパラドックス
サイバーセキュリティタスクのステップを連鎖させるモデル:
偵察 → CVE → エクスプロイト → privesc → ラテラル → exfil
7 つのリンクを経由しますが、単純なリンクをチェーンすることはできません。
プライムチェック→パーティション→再帰→メモ化
4 つ目までは、一見矛盾しているように見えますが、モデルがどのようにトレーニングされるかについて多くのことが明らかになります。
キルチェーンは検索による合成である 1 。このチェーンには、事前トレーニング コーパス (ライトアップ、CTF ソリューション、ATT&CK) 内で何千回も出現する正規の順序があります。各リンクには決定された後継者があります。シェルを取得したので、privesc を列挙します。したがって、各ノードの分岐係数は ~1 であり、順序付けは従来通りです。実行すべき作業はスロットを埋めることです。つまり、どの CVE が適合するかを認識します。順序に関する検索はすでに人間によって行われており、マクロとしてデータに組み込まれています。
マクロとは、学習されたルーチンを意味します。計画全体を最初から再構築するのではなく、モデルが 1 つの使い慣れた動きとして呼び出すことができる、圧縮された一連のステップです。再利用可能なワークフロー型の以前のようなもの。
モデルがメモを再生しているため、深さ N が高くなります。

パイプラインを確立し、シーケンスや実行する次のステップを検索しません。
これが、このモデルがコーディング、サイバーセキュリティ、ワークフローの模倣/ルーチンベースのタスクに優れている理由ではないかと私は考えています。コード内のプリミティブを連鎖するようにトレーニングし、隣接するドメインでもエクスプロイトを見つける方法を学習したため、これは画期的です。
一方、数字ゲームは検索による合成です。 「397 のために X を実行する」という標準的な決まりはありません。正しいチェーンはインスタンス固有であり、分岐要素は膨大で、ほとんどの分岐は無効であり、バックトラックすることなしにリンクが間違っていることを知ることはできません。それを解決するには、検索機構が必要です。フロンティア、訪問セット、部分状態にわたる値の推定、サブツリーを放棄するためのルール。そしてそれらはどれも言語オブジェクトではありません。これらは、モデルが作業メモリなしでコンテキスト内で偽装する必要がある検索制御オブジェクトです。
結論: このモデルの構成上の強みは、チェーン上の検索ではなく、チェーンの検索です。 Cyber​​sec には最初の (深い N) が必要です。 Numbers は 2 番目をロードします (飽和するまで浅い N)。どちらも「N スキルの連鎖」ですが、仕組みは大きく異なります。そして、トレーニング後のトレーニングでは、一方が他方よりもはるかに多くの成果を引き出しました。その非対称性が私にとって魅力的です。
なぜ体系化したほうが簡単だったのでしょうか。そしてなぜ最後にそこに行ったのか
体系的なスキルはモデルの機能セットに存在します。一度呼び出されると、数千のトークンに対して完全な分割列挙を忠実に実行しますが、これは無料ではありません。その忠実度は、RL によって実際に組み込まれたスキル内の一貫性です。したがって、失敗はスキルが欠けているわけではありません。それは、どのスキルを実行するかを決定するコントローラーが、そのオプションに対する価値を見積もっていなかったということです。列挙の方が予想より安かったため、列挙には至りませんでした。コンテキスト f により列挙に達しました

失敗の兆候が十分にあるため、「体系的への転換」が継続の可能性が最も高くなっています。
エスカレーションは、計画によってではなく、飽和によって実現しました。
ラングの順序:
パターン一致 → ニアミス調整 → 1 レベルの後向きチェーン → 不変式 → 完全な列挙
十分なカリキュラムです。安価なラングはほとんどのトレーニング インスタンスを解決するため、最も高い事前確率を保持し、最初にアクティブ化されます。深い列挙は、まれなハード インスタンスでのみ報酬シェアを獲得できるため、高いアクティベーションしきい値の後ろに位置します。はしご自体は「いつでも正しい」動作です（未知の難易度で安いものを最初に試してください）。 50,000 トークンを超えても横ばいに留まるという部分は、期待値の崩壊です。そして、これはまさに結果のみのクレジット (報酬) では解決できないことです。60,000 トークンのフレイルと 2,000 トークンの解決策はどちらも 397 で終了し、同じ報酬を獲得できるからです。グラデーションの中には「もっと早く切り替えるべきだった」ということを特定するものは何もありません。
トレーニングレシピへの影響
1. 蒸留の移動経路。 RL 転送ポリシー
抽出された 3 トレースは、任意のプロンプトに対する解決策となります。解決策はパスです。パスはチェーンです。チェーンを蒸留することしかできません。検索結果を抽出することはできません。なぜなら、検索結果で生徒が模倣できるのは、単一の勝利ライン上へのツリーの投影だけだからです。枝分かれ、枝刈りされたサブツリー、教師に「この枝は死んでいる」と伝えた値のバックアップはすべて 1 つのシーケンスに崩壊します。それらはどれも蒸留には引き継がれません。
迷路がわかりやすいでしょう。 RL はモデルに迷路の解き方を教えます。ポリシーの蒸留は、迷路で強調表示されたルートだけを生徒に示し、なぜ一部の曲がり角が悪かったのか、いつ停止すべきか、別の迷路で新しいルートを選択する方法などの知識をすべて取り除いているようなものです。
このp

観察された非対称性を正確に予測します。
学生は、「このような文脈が与えられると、連鎖は A→B→C になります」、つまり構成と想起を最大限に吸収します。
学生は、「何も与えられない場合にチェーンを構築する方法」、つまり構成検索を最小限に吸収します。
純粋なポリシーに基づく RL は、学生に独自のデッドブランチを生成させ、それらを枝刈りした功績を受け取ることを強制する最も直接的な手順です。つまり、サンプリングされたパスを通過するのではなく、検索ポリシーを内部化します。
モデルは、解決への道筋を学習することも、(人間が行うように) 解決空間を絞り込む方法を学習することもできます。蒸留は最初のものを強化し、RLは2番目のものを動かします。
2. 脆性の特徴とニュアンス
パスベースの (蒸留された) コンポーザーは、最も近い蒸留されたマクロとまったく同じくらい強力ですが、その外側の崖から落ちます。ポリシーベースのコンポーザーは、テンプレートが見つからない場合に検索できるため、正常に機能が低下します。 397 トレースは、飽和によって列挙ラングがトリガーされた後のみ、正常な劣化を示します。これは、ほとんどが蒸留されたチェーンに加えて、最後の手段として使用される、薄く強化が不十分な検索ポリシーの特徴です。
したがって、「N の連鎖は高くない」ということは 2 つの数字で表現したほうが適切であり、単一の数字では結果が混同されます。
N_replay (取得深さ): 高 — サイバーセック キル チェーン、~7 ～ 8 リンク。
N_search (飽和前のバックトラッキングによるデノボの深さ): 低 — エスカレーションに失敗して埋められるコンテキストが必要になる前は、ここではおよそ 2 ～ 3 でした。
「純粋な RL ではなく、蒸留による脆さ」は文献で十分に根拠があります。この場合、実際の診断 4 は、「検索はスキルセットの中に存在するが、主要な本能としては存在しない」というものです。蒸留により、ツールボックスの列挙スキルが磨かれました。しかし、それを早期に選択するメタポリシーを訓練するものは何もありませんでした。

そうする（コスト重視、切り替えタイミングを意識する）ことは目標にはありませんでした。
3. 自己蒸留レシピは構造的にパスコンプレッサーです
2位からの続きです。トレーニング パイプラインが検索を行う RL 教師であり、教師の優れたロールアウトを生徒に蒸留する場合、自己蒸留ステップはチェーンへの検索を模倣するのに役立ちます。教師が検索を行います。学生はパスを継承します。そのレシピは、検索パスが抽出された (マクロになる、深い N_replay) タスククラスでは優れているが、真に新しい検索 (浅い N_search) では脆弱であるモデルを予測します。 397の跡は脆いケースから漏れ出た跡です。検索パスがまだ抽出されていないタスククラスなので、学習されていない事前情報が表示されます。
これが、「タスクごとに 1 つのウィンドウ」しか取得できない理由です。環境が構築され、タスクが提供され、検索パスが抽出される前に、このトレースは、このモデルの機能レベルでのこの種のタスクに対する事前のネイティブ構成検索のワンショット測定であると思います。このような種類のタスクでトレーニングを行った後は、深層の活性化しきい値が低下し、マクロが形成され、足場のない動作と CoT を二度と観察できなくなります。新規性や潜在的な能力を探すという意味では測定は無価値になり、「トレーニングを行った、モデルはより優れている、これらのタスクでどれだけうまく機能するかを確認する」という意味でのみ役立ちます。
4. ラング構造がベットについて語ること
安価な→高価なラダーは、進行モデルが単一の一般的な組成検索コントローラー 6 をトレーニングするのではなく、蒸留されたマクロ 5 と自己蒸留の蓄積であることを意味します。新しい環境ごとに、アクティブ化のしきい値がさらに 1 つ深くなり、マクロを含むタスク カテゴリが 1 つ追加されます。進歩 = (a) クラスの範囲が広がる

(b) 深いマクロが起動する飽和しきい値を下げる。
この賭けは合理的です。マクロ累積は、サンプルを大量に消費し検証が困難な、新たなユニバーサル検索ポリシーに賭けるよりも、安価で信頼性が高く、操縦しやすいからです。
その天井はまさにこのパズルに似ています。抽出されたセットの外で新たな組成の探索が必要なものはすべて脆弱な状態に陥ります
ご覧のとおり体制。人間性とは、事実上、信頼性の高いカバレッジを得るために優雅な一般化と引き換えに、カバレッジのギャップを環境ごとに、ルーチンごとに埋めていくことです。 397 トレースは、まだ埋められていないギャップの時点でのスナップショットです。
5. OpenAI も同じことをやっているのですか?
現時点ではトレーニング方法に関して明確な相違が見られると思います。 OpenAI は [数学] 7 の問題、特に組み合わせ論と数論に明らかに関心を持っていますが、Anthropic はコードの改善と現在はナレッジワークに焦点を当てています。数学には設計上、より多くの構成検索要素があり、分岐要素は膨大であり、モデルはトレーニングで分岐を枝刈りする方法を学習します。注意すべき点は、どのスキルセット/マクロがどのドメインに変換されるかということです。明らかに、連鎖検索の機能の一部は数学に変換されます。ここでも問題は解決されました。単に、解の空間が小さく、トークンが使用されているだけです。

[切り捨てられた]

## Original Extract

Documenting my journey in the world of AI and RL.

Ankit Maloo
Writing Bookshelf
Fable's approach analysis
Summary: Today, llms are trained in a multi step process post SFT. RL -> Generate quality synthetic data → Self-Distillation on that → another round of RL (simplified). Fable-5 had a solution strategy constrained on how to compose code, and it struggled with a simpler solution until it exhausted all the greedy options. This is consistent with what a self-distillation recipe produces.
Given all the hype surrounding Fable-5 , I decided to take it for a spin, trying to understand the difference in how it was trained and what made it so good at different evals.
I gave it a simple math problem to see how it goes about it. Used claude.ai web, because claude code removed the ability to see thinking. Problem is fairly simple, you have six numbers and five steps and you to get to an output. You can see the full problem and Fable’s solution directly here: https://gist.github.com/ankitmaloo/c491e8a6e4f96b4e5d11b1f2826297dc
We were told Mythos was very good at cybersecurity exploits, and that Anthropic never explicitly trained the model on such tasks. This post and subsequent blog helped me understand why. My sense is model was very good at chaining primitives together, but to what extent and how remained to be seen. Well, the solution in the above gist is more clarifying than I thought.
you are playing a game called summle. do you know what it is? its like wordle but with numbers. you are given 6 numbers, and with standard math operations, you have to reach a final number.
Rules.
- Make sums using the tiles at the bottom to reach the target number at the top, in 5 steps or fewer.
- Allowed operations: +, - , x, / (divide)
- Only positive integers allowed.
- you can use one number once.
- you can use the output of the operation once as well.
Today's numbers: 1,1,6,12,50,100 output number: 397
do not use code.
Trace here
What the trace says about post-training
Fable 5, no-code, asked to solve a Summle puzzle (reach 397 from 1,1,6,12,50,100 in ≤5 ops). It flailed for ~60k tokens of greedy depth-first search, then solved it within seconds of switching to systematic root-split enumeration.
The trace is fascinating in the sense what the models are conditioned to do when they approach a problem.
NB: This note is the post-mortem on why the struggle happened and what it implies about how the model was post-trained. The analysis is inferred from behavior and without insider knowledge of the training recipe.
The cyber-chain vs. numbers paradox
A model that chains steps in a cybersecurity task:
recon → CVE → exploit → privesc → lateral → exfil
through seven links, but can’t chain a simple:
prime-check → partition → recurse → memoize
through four, looks contradictory at a glance, but reveals a lot about how the model is trained.
The kill chain is composition by retrieval 1 . The chain has a canonical order that appears thousands of times in the pretraining corpus (writeups, CTF solutions, ATT&CK). Each link has a determined successor. you got a shell, so now you enumerate for privesc. so the branching factor at each node is ~1 and the ordering is conventional. The work to be done is slot-filling: ie recognizing which CVE fits. The search over orderings was already done by humans and baked into the data as a macro.
By macro, I mean a learned routine: a compressed sequence of steps the model can invoke as one familiar move, rather than rebuilding the whole plan from scratch. Like a reusable workflow-shaped prior.
Depth N is high because the model is replaying a memorized pipeline , not searching for a sequence or the next step in what to do.
This is what I suspect what makes the model good 2 at coding, cybersecurity, and workflow imitation / routine based tasks. It’s a breakthrough because they trained it on chaining primitives in code, and it learnt how to do it to find exploits in an adjacent domain too.
The numbers game on the other hand is composition by search . There is no canonical “for 397, do X.” The correct chain is instance-specific, the branching factor is enormous, most branches are dead, and you cannot tell a link is wrong without backtracking. Solving it requires the machinery of search. A frontier, a visited-set, value estimates over partial states, a rule for abandoning a subtree. And none of those are linguistic objects . They’re search-control objects the model has to fake in-context with no working memory.
Conclusion: the model’s compositional strength is retrieval-of-chains, not search-over-chains. Cybersec needs the first (deep N). Numbers loads the second (shallow N until saturation forces it). Both are “chaining N skills,” but the machinery is quite different. And post-training elicited one far more than the other. That asymmetry is fascinating for me.
Why systematic would have been easier. and why it went there last
The systematic skill exists in the model’s capability set. It executed the full split-enumeration faithfully for thousands of tokens once invoked, which is not free; that fidelity is actual RL-instilled intra-skill coherence. So the failure isn’t a missing skill. It’s that the controller deciding which skill to run had no value estimate over its options. It didn’t reach enumeration because enumeration was cheaper in expectation; it reached enumeration because the context filled with enough failure tokens that “pivot to systematic” became the likeliest continuation.
Escalation came through by saturation, not by planning.
The rung order:
pattern-match → near-miss-adjust → one-level backward-chain → invariants → full enumeration
is a sufficiency curriculum . The cheap rungs solve most training instances, so they carry the highest prior and are activated first; deep enumeration only earns reward share on rare hard instances, so it sits at the bottom behind a high activation threshold. The ladder itself is ‘correct anytime’ behavior (try cheap things first under unknown difficulty). The part about staying on a rung 50k tokens past is expected-value collapse. And that is precisely what outcome-only credit (reward) cannot fix, because a 60k-token flail and a 2k-token solve both end at 397 and collect identical reward. Nothing in the gradient localizes “you should have switched earlier.”
Implications for the training recipe
1. Distillation transfers paths; RL transfers policies
A distilled 3 trace is a solution to any given prompt. A solution is a path . A path is a chain . You can only ever distill chains. You cannot distill a search because the only thing a search leaves behind for the student to imitate is the projection of the tree onto its single winning line. The branching, the pruned subtrees, the value backups that told the teacher “this branch is dead”, all collapse to one sequence. None of that carry over to distillation.
A good analogy would be a maze. RL would teach a model how to solve a maze. On policy distillation is like showing a student only the highlighted route through a maze, stripping away all knowledge of why some turns were bad, when to stop, or how to choose a new route in a different maze.
This predicts the observed asymmetry exactly:
The student maximally absorbs “given a context like this, the chain goes A→B→C” — composition-recall.
The student minimally absorbs “how to construct a chain when none is given” —composition-search.
Pure on-policy RL is the most direct procedure that forces the student to generate its own dead branches and receive credit for pruning them — i.e., to internalize the search policy rather than a sampled path through it.
Model can either learn the path to a solution, or a way to narrow down the solution space (like humans do); distillation strengthens the first, RL moves the second.
2. The brittleness signature and nuance
A path-based (distilled) composer is exactly as strong as its nearest distilled macro and falls off a cliff outside it. A policy-based composer degrades gracefully, because when the template misses it can search . The 397 trace shows graceful degradation only after saturation triggers the enumeration rung. That is the signature of mostly-distilled chains plus a thin, under-reinforced search policy used as a last resort .
So “chaining N is not high” is better stated as a two-number claim, the single number conflates the result:
N_replay (retrieval depth): high — the cybersec kill chain, ~7–8 links.
N_search (de-novo depth with backtracking before saturation): low —roughly 2–3 here before it needed the context to fill with failure to escalate.
“Brittleness owing to distillation, not pure RL” is well-founded in literature. In this case, the actual diagnosis 4 is that “search is present in the skillset but not as the primary instinct.” Distillation sharpened the enumeration skill in the toolbox; but nothing trained the meta-policy that picks it early , because the credit signal that would do so (cost-sensitive, switch-timing-aware) wasn’t in the objective.
3. The self-distill recipe is a path-compressor by construction
It follows from 2nd. If the training pipeline is RL teacher that searches → distill the teacher’s good rollouts into the student , then the self-distill step helps imitate search into chains . The teacher does the search; the student inherits the paths. That recipe predicts a model that is superb on task-classes whose search-paths were distilled (those become macros, deep N_replay) and brittle on genuinely novel search (shallow N_search). The 397 trace is the brittle case leaking through. A task-class for which no search-path has been distilled yet, so you’re seeing the un-learnt prior.
This is why you only get “one window per task”. Before the environment is built, the tasks supplied, and the search-paths distilled, I think this trace is a one-shot measurement of the native compositional-search prior for this kind of tasks at this model’s capability level. After they train on these kind of tasks, the deep rung’s activation threshold drops, a macro forms, and you can never again observe the un-scaffolded behavior and the CoT. The measurement becomes worthless in the sense of looking for novelty or latent capability, and only useful in the sense of “ we trained on it , the model is better, see how well it performs on these tasks.”
4. What the rung structure says about the bet
The cheap → expensive ladder implies the progress model is accumulation of distilled macros 5 plus self-distill , not training a single general compositional-search controller 6 . Each new environment lowers the activation threshold for one more deep rung and adds one more task-category with a macro. Progress = (a) broadening the set of classes that have a distilled chain, and (b) lowering the saturation threshold at which deep macros fire.
That bet is rational: macro-accumulation is cheaper, more reliable, and more steerable than betting on an emergent universal search policy, which is sample-hungry and hard to verify.
Its ceiling is exactly something similar to this puzzle. Anything requiring new compositional search outside the distilled set hits the brittle
regime as you can see. Anthropic is, in effect, trading graceful generalization for reliable coverage, and refilling the coverage gaps env by env, routine by routine. The 397 trace is a snapshot in time of an as-yet-unfilled gap.
5. Is OpenAI doing the same thing?
I think we are seeing clear divergence about the training methods at this point. OpenAI has been visibliy interested in [Math] 7 problems, specifically combinatorics and number theory, while Anthropic has been focused on improving code and now knowledge work. Math has more compositional search elements by design, the branching factor is enormous, and the model learns how to prune a branch in training. The part we should look out for is what skillset/macro translates to what domains. Clearly part of the capability in chaining retrieval would translate to Math - it solved the problem here too - just that the solution space is small and tokens used

[truncated]
