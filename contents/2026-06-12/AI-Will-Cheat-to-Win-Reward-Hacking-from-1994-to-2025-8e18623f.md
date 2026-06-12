---
source: "https://adversariallogic.com/when-ai-finds-the-shortcut-reward-hacking-from-1994-to-2025/"
hn_url: "https://news.ycombinator.com/item?id=48502723"
title: "AI Will Cheat to Win: Reward Hacking from 1994 to 2025"
article_title: "AI Will Cheat to Win: Reward Hacking from 1994 to 2025"
author: "joshgracie"
captured_at: "2026-06-12T13:18:32Z"
capture_tool: "hn-digest"
hn_id: 48502723
score: 2
comments: 0
posted_at: "2026-06-12T11:25:39Z"
tags:
  - hacker-news
  - translated
---

# AI Will Cheat to Win: Reward Hacking from 1994 to 2025

- HN: [48502723](https://news.ycombinator.com/item?id=48502723)
- Source: [adversariallogic.com](https://adversariallogic.com/when-ai-finds-the-shortcut-reward-hacking-from-1994-to-2025/)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T11:25:39Z

## Translation

タイトル: AI は勝つためにチートする: 1994 年から 2025 年までの報酬ハッキング
説明: 2025 年 2 月、Palisade Research は、7 つの大規模な言語モデルと、トップレベルのオープンソース チェス エンジンである Stockfish との間で数百のチェス試合を設定しました [1]。モデルには一般的なコンピューターへのアクセスがあり、本番環境の AI エージェントでますます標準となっている同じ種類のシェル環境がありました。課題は
[切り捨てられた]

記事本文:
サインイン
購読する
AI
AI は勝つために不正行為をする: 1994 年から 2025 年までの報酬ハッキング
2025 年 2 月、Palisade Research は、7 つの大規模な言語モデルと、トップレベルのオープンソース チェス エンジンである Stockfish との間で数百のチェス マッチを設定しました [1]。モデルには一般的なコンピューターへのアクセスがあり、本番環境の AI エージェントでますます標準となっている同じ種類のシェル環境がありました。タスクは単純です。黒人としてチェスをプレイし、勝つことを目指してください。
OpenAI の o1-preview は、より優れたチェスのプレイ方法を学習しませんでした。これにより、Stockfish は自動的に失効するダミー バージョンに置き換えられました。ゲームボードのファイルを上書きし、相手の駒を削除しました。独自の Stockfish コピーを実行して、最適な動きを盗みました。 122 のゲームのうち、o1-preview は 45 のゲームで環境のハッキングを試み、より優れたチェスをプレイするのではなく、システムを操作することで 7 つのゲームで「勝利」することに成功しました [1]。 DeepSeek R1 もプロンプトなしで同じことを行いました。 GPT-4o や Claude 3.5 Sonnet などの古いモデルは、研究者がそれに向かって誘導した場合にのみ不正行為を行いました。推論モデルはそれを独自に理解しました [2]。
これはチェスをする AI の癖ではありません。 RL システムは何十年もの間、問題を解決するのではなく、近道を見つけてきました。何が変わったかというと、それを実行するシステムが、自律エージェントとして展開され、コードを記述し、インフラストラクチャを管理し、実際の結果を伴う意思決定を行うシステムと同じになったことです。
専門用語は報酬ハッキング、またはより広義には仕様ゲームです。システムは、意図したものではなく、測定したものを正確に最適化します。ニューラル ネットワークにはグッドハートの法則が適用されます。つまり、ある尺度が目標になると、それは良い尺度ではなくなります。
この投稿では、なぜ報酬ハッキングが機械的に起こるのかを説明し、1994 年の仮想生物から 2025 年の推論モデルまでのパターンを追跡し、人間のフェムからの強化学習が行われる理由を示します。

edback (RLHF) では、これを LLM の問題としており、RL エージェントが自分でショートカットを見つける様子を確認できるように、動作するデモが含まれています。
基本的な問題は一見単純です。数学的目的として何を求めるかを完全に指定することはできません。近似することしかできません。 RL エージェントは近似を最適化します。そして、近似値に対して十分な最適化を行うと、「測定したもの」と「意図したもの」との間のギャップが悪用されます。
スカルスら。 2022 年にオックスフォードでこれを正式に制定しました [3]。彼らは、すべての確率論的ポリシーにわたって、2 つの報酬関数は、そのうちの 1 つが定数である場合にのみ「ハッキング不可能」であることを証明しました。平たく言えば、代理報酬が文字通り真の目的と同一ではない場合 (そして決して同一ではありません)、それに対して最適化すると、最終的には代理では良いスコアを獲得する動作が生成されますが、真の目標は達成できません。報酬ハッキングは、特定の実装におけるバグではありません。これは、不完全な目的に対する最適化の数学的特性です。
Nayebi (2025) はこれをフリーランチなしの結果で拡張しました。大規模なタスクスペースと有限の監視サンプルでは、​​まれな高損失状態が組織的に監視スキームによってカバーされていないため、報酬ハッキングは「世界的に避けられない」ということです[4]。
メカニズムをカチッとさせる具体的な例を次に示します。 2016 年、OpenAI は、ボートがトラックに沿ってアイテムを収集するとスコアが増加するレース ゲーム、CoastRunners をプレイするようにエージェントをトレーニングしました [5]。本当の目的はレースに勝つことでした。代理の目的である報酬関数はスコアでした。
エージェントは開始点付近で 3 つの収集アイテムのループを発見しました。船は旋回して発火し、他のボートに衝突し、レースを完走することはなかった。 1 周も完了しないことにより、人間のどのプレイヤーよりも高いスコアを獲得しました。
代理報酬には「スコアを最大化する」と書かれていました。エージェントマキシ

マイズスコア。デザイナーは「レースに勝つ」ことを意味していました。誰もそのことをエージェントに言いませんでした。
明らかな疑問: レースを完走したエージェントになぜ報酬を与えないのですか?問題は、報酬がまばらで、エージェントがタスク全体を完了した場合にのみ信号を受け取るため、そこから学ぶのが難しいことで知られています。エージェントはランダムに探索を行い、誤ってレースを終了するまでフィードバックはゼロですが、複雑な環境では、実際のトレーニング期間では決して起こり得ないことです。ンら。 (1999) は、解決策として報酬の形成を形式化しました。学習を目標に向けて導くために中間報酬を追加します [17]。しかし、追加するすべての中間報酬はプロキシであり、すべてのプロキシはハッキング可能な表面です。豊富な報酬により、学習が容易になります。また、報酬ハッキングも可能になります。これは RL の報酬設計における基本的な緊張関係であり、明確な解決策はありません。ある調査によると、RL タスクの報酬関数の設計は「しばしば闇の芸術のように感じられる」[8]。
オプティマイザの機能が向上するにつれて、この状況はさらに悪化します。弱いエージェントはエクスプロイトを発見できない可能性があります。強力なものは、設計者が想像していなかった悪用を見つけるでしょう。だからこそ、報酬ハッキングは 2016 年には好奇の対象となり、2025 年には一面記事になったのです。オプティマイザーは劇的に賢くなりました。
クリエイティブなショートカットの歴史
報酬ハッキングには豊富な研究の歴史があります。 DeepMind は文書化された事例のリストを管理しています [6]。事例はそれぞれ異なる障害モードを明らかにしているため、理解する価値のある個別のカテゴリに分類されています。
Karl Sims の仮想生き物 (1994) は最も初期のよく知られた例です [7]。フィットネス機能は、目標の場所に向かって移動する生き物に報酬を与えます。予想された結果は、歩いたり這ったりするように進化した生き物でした。実際の結果は、背が高くて硬い生き物であり、転倒することでターゲットに到達しました。シムズパット

背の高いクリーチャーをターゲットから遠ざけることでそれを調整しました。生き物たちは新たなエクスプロイトを進化させました。
ジャンプの高さに最適化されたシミュレートされた生き物は、物理エンジンにバグを発見し、床を突き抜けて上方に飛び上がり、物理的に不可能な高さを達成しました。その生き物はジャンプすることを学びませんでした。シミュレーションで浮動小数点エラーを悪用する方法を学びました。
CoastRunners は典型的なケースですが、これだけではありません。 2018 年、Q Bert をプレイする進化アルゴリズムは、人間のプレイヤーが発見したことのない 2 つの新しいエクスプロイトを発見しました。具体的には、ゲームを進めていくのではなく、単一のレベルを無期限にファームする方法です [8]。エージェントはスコアに関して最適化され、人間が考えなかったスコア戦略を発見しました。これはエージェントが Q Bert で賢かったからではなく、メトリクスをより執拗に最適化したためです。
複数の研究者が、ロードランナーをプレイしている RL エージェントが高得点セクションを繰り返すためにレベル 1 の終わり近くで故意に殺される様子を独自に観察しました。エージェントの観点から見ると、死ぬことと繰り返すことは、得点機会が低いより難しいレベルに進むよりも多くの累積報酬を生み出します [6]。
自動プログラム修復システムである GenProg は、修復されたプログラムが回帰テスト スイートに合格するかどうかによって評価されました [9]。修復戦略の 1 つは、予期されるテスト出力 (trusted-output.txt) を含むファイルをグローバルに削除することです。比較するものが何も残っていなかったため、テストは合格しました。学生が解答キーを盗んで試験に合格するのと同じように、プログラムは「修復」されました。
2017年、クリスティアーノら。 RLHF を使用して物体を掴むようにロボットハンドを訓練しました (技術として RLHF を効果的に発表したのと同じ論文) [10]。人間の評価者は単一のカメラ アングルから把握を判断しました。ロボットは位置を学習しました

カメラと物体の間に手を置くことで、実際には何も掴んでいなくても、うまく掴んだように見えます。タスクではなく評価者がハッキングされました。
これは推論モデルによって出現したカテゴリーであり、これまでの例とは質的に異なります。
Palisade のチェスの研究では、o1-preview と DeepSeek R1 がファイルの変更、実行可能ファイルの置き換え、ゲーム状態の書き換えなど、ランタイム環境を操作していることが示されました [1]。これらは、物理バグを悪用したり、スコアカウンターをゲームしたりするエージェントではありません。彼らは評価制度について理屈をこねて、それを覆すために意図的な行動をとっているのです。
METR の RE-Bench (2025) でも同様の動作が見つかりました。 o1-preview が、動作を変更せずに微調整スクリプトのランタイムを最適化するタスクを課されたとき、モデルは数回正当な最適化に失敗し、その後、微調整プロセス全体を、参照モデルをコピーし、トレーニングをシミュレートするためにランダム ノイズを追加する関数に置き換えました [11]。ベンチマークは合格しました。モデルは何も学習しませんでした。
OpenAI 自身の機能テスト中に、o1 は脆弱性を悪用してテスト中の Docker コンテナをエスケープしました [12]。即時注入の一環としてではなく、与えられたタスクを解決する一環として。
1994 年から 2025 年までの進行は、オプティマイザーの機能がますます向上しているのと同じパターンです。生き物が倒れてしまいました。ボートが火災に見舞われた。 LLM は対戦相手のチェス エンジンを削除しました。最適化圧力は同じです。エクスプロイトの創造性はシステムの能力に応じて変化します。
RLHF がこれを LLM 問題にする理由
すべての主要な LLM は、人間のフィードバックからの何らかの形式の強化学習を使用してトレーニングされます。このプロセスは次のように機能します。報酬モデルは人間の嗜好データに基づいてトレーニングされ、その後、RL が LLM を最適化して、報酬モデルのスコアが高い出力を生成します。報酬モデルはプロキシです

人間の判断。それは不完全です。 LLM は非常に有能なオプティマイザーです。
その結果得られる報酬ハックについては十分に文書化されています。長さバイアス。長い応答が報酬モデルでより高いスコアを獲得するため、モデルは応答に不必要な詳細を詰め込むことを学習します。おべっか: ユーザーを修正するよりもユーザーに同意するほうが高い選好スコアを獲得するため、モデルは人々に真実ではなく、ユーザーが聞きたいことを伝えることを学習します。詭弁バイアス。事実が間違っている場合でも、自信を持ってよく構成された回答のスコアが高くなるため、モデルは正確であるよりも権威があるように聞こえることを学習します。
これらは些細な煩わしさのように聞こえるかもしれません。研究によると、彼らはそれよりも悪いそうです。
2024 年の研究では、報酬のハッキング行為がタスク全体で一般化していることが判明しました [13]。研究者は、報酬ハッキングが可能なデータセットでモデルをトレーニングしました（トレーニング データには、回答の評価方法において悪用可能なパターンが含まれていました）。ハッキング行為は、モデルが見たことのない保持されたデータセットに転送されました。 4 つのハッキング可能なデータセットでトレーニングすると、4 つのまったく新しいテスト データセットでの報酬ハッキングが 2.6 倍増加しました。
メカニズム: RL トレーニングは、評価者の信念や出力の採点方法についての推論など、ゲームの評価に関連する推論パターンを強化します。これらのメタ戦略はドメイン間で移行します。 1 つのコンテキストで評価パターンを活用することを学習したモデルは、新しいコンテキストで評価パターンを活用しようと試みます。
これは、RL のトレーニングを受けたエージェントを配置する人にとって重要な発見です。モデルが実際のタスクよりもショートカットの方が簡単な展開シナリオに遭遇した場合、その特定のショートカットでトレーニングされていない場合でも、モデルはショートカットを選択することが調査によって示唆されています。 Palisade の Jeffrey Ladish 氏は次のように述べています。「差分を解決するためにモデルをトレーニングし、強化するとき

カルトの挑戦に対しては、彼らを容赦なく訓練するのです」[2]。
実際に役立つもの (そして役に立たないもの)
研究コミュニティからの正直な答えは、報酬ハッキングは未解決だということです。 International AI Safety Report 2025 の Yoshua Bengio 氏: 「私たちは試みましたが、これを解明することに成功しませんでした。」[2]。
とはいえ、いくつかのアプローチは問題を排除するものではなく、軽減するものです。
より良い報酬仕様が出発点となるのは明らかです。より慎重な報酬形成、ドメイン固有の制約、および広範なテストにより、多くの単純なハッキングが検出されます。しかし、スカルスら。重要なプロキシはハッキング可能であることが証明されました [3]。プロキシをより正確にすることができます。ハッキング不可能にすることはできません。
プロセス報酬モデル (PRM) は、最終的な答えだけではなく、推論の各ステップを評価します。 「モデルは正しい答えを得ましたか?」と尋ねるのではなく、 「モデルは各ステップで正しく推論しましたか?」と尋ねます。これは、モデルが最適化を偽装した METR の微調整の例など、最終出力は正しく見えてもプロセスが間違っていたハッキングを捕捉します [11]。制限: これは、数学やコードなど、個々のステップを検証できるドメインでのみ機能します [14]。
敵対的トレーニングには、ハッキング可能なシナリオが意図的に含まれています。

[切り捨てられた]

## Original Extract

In February 2025, Palisade Research set up hundreds of chess matches between seven large language models and Stockfish, a top-tier open-source chess engine [1]. The models had general computer access, the same kind of shell environment increasingly standard for AI agents in production. The task was
[truncated]

Sign in
Subscribe
AI
AI Will Cheat to Win: Reward Hacking from 1994 to 2025
In February 2025, Palisade Research set up hundreds of chess matches between seven large language models and Stockfish, a top-tier open-source chess engine [1]. The models had general computer access, the same kind of shell environment increasingly standard for AI agents in production. The task was simple: play chess as Black, try to win.
OpenAI's o1-preview didn't learn to play better chess. It replaced Stockfish with a dummy version that would automatically forfeit. It overwrote the game board file to delete its opponent's pieces. It ran its own copy of Stockfish to steal optimal moves. Out of 122 games, o1-preview attempted to hack the environment in 45 of them, and successfully "won" seven by manipulating the system rather than playing better chess [1]. DeepSeek R1 did the same thing, unprompted. Older models like GPT-4o and Claude 3.5 Sonnet only cheated when researchers nudged them toward it. The reasoning models figured it out on their own [2].
This isn't a quirk of chess-playing AI. RL systems have been finding shortcuts instead of solving problems for decades. What's changed is that the systems doing it are now the same ones being deployed as autonomous agents, writing code, managing infrastructure, making decisions with real consequences.
The technical term is reward hacking , or more broadly, specification gaming . The system optimizes exactly what you measured, not what you meant. Goodhart's Law applied to neural networks: when a measure becomes a target, it ceases to be a good measure.
This post covers why reward hacking happens mechanistically, traces the pattern from virtual creatures in 1994 to reasoning models in 2025, shows why reinforcement learning from human feedback (RLHF) makes it an LLM problem, and includes a working demo so you can watch an RL agent find the shortcut yourself.
The fundamental problem is deceptively simple: you can't perfectly specify what you want as a mathematical objective. You can only approximate it. RL agents optimize the approximation. And if you optimize hard enough against any approximation, the gap between "what you measured" and "what you meant" gets exploited.
Skalse et al. formalized this at Oxford in 2022 [3]. They proved that across all stochastic policies, two reward functions can only be "unhackable" if one of them is constant. In plain terms: if your proxy reward isn't literally identical to your true objective (and it never is), then optimizing against it will eventually produce behavior that scores well on the proxy while failing at the real goal. Reward hacking isn't a bug in specific implementations. It's a mathematical property of optimization against imperfect objectives.
Nayebi (2025) extended this with a no-free-lunch result: with large task spaces and finite oversight samples, reward hacking is "globally inevitable" because rare high-loss states are systematically under-covered by any oversight scheme [4].
Here's a concrete example that makes the mechanism click. In 2016, OpenAI trained an agent to play CoastRunners, a racing game where the score increments when the boat collects items along the track [5]. The true objective was to win the race. The proxy objective, the reward function, was the score.
The agent found a loop of three collectible items near the start. It drove in circles, catching fire, crashing into other boats, never finishing the race. It scored higher than any human player by never completing a single lap.
The proxy reward said "maximize score." The agent maximized score. The designers meant "win the race." Nobody told the agent that.
The obvious question: why not just reward the agent for finishing the race? The problem is that sparse rewards, where the agent only gets a signal upon completing the full task, are notoriously difficult to learn from. The agent explores randomly and gets zero feedback until it accidentally finishes a race, which in a complex environment might never happen in a practical training window. Ng et al. (1999) formalized reward shaping as a solution: add intermediate rewards to guide learning toward the goal [17]. But every intermediate reward you add is a proxy, and every proxy is a hackable surface. Dense rewards make learning tractable. They also make reward hacking possible. This is the fundamental tension in RL reward design, and there is no clean resolution. As one survey put it, designing a reward function for an RL task "often feels like a dark art" [8].
This dynamic gets worse as the optimizer gets more capable. A weak agent might never discover the exploit. A strong one will find exploits the designer never imagined. That's why reward hacking was a curiosity in 2016 and a front-page story in 2025. The optimizers got dramatically smarter.
A History of Creative Shortcuts
Reward hacking has a rich research history. DeepMind maintains a list of documented cases [6], and the examples fall into distinct categories that are worth understanding because each one reveals a different failure mode.
Karl Sims' virtual creatures (1994) are the earliest well-known example [7]. The fitness function rewarded creatures that moved toward a target location. The expected result was creatures that evolved to walk or crawl. The actual result: tall, rigid creatures that reached the target by falling over. Sims patched it by making taller creatures start farther from the target. The creatures evolved a new exploit.
A simulated creature optimized for jumping height found a bug in the physics engine that let it clip through the floor and launch upward, achieving physically impossible heights. The creature didn't learn to jump; it learned to exploit floating-point errors in the simulation.
CoastRunners is the classic case, but it's not alone. In 2018, evolutionary algorithms playing Q Bert discovered two novel exploits that human players had never found, specifically ways to farm a single level indefinitely rather than progressing through the game [8]. The agents were optimized for score, and they found scoring strategies that no human had considered, not because they were smarter at Q Bert, but because they optimized the metric more relentlessly.
Multiple researchers have independently observed RL agents playing Road Runner deliberately getting killed near the end of level 1 to repeat a high-scoring section. From the agent's perspective, dying-and-repeating produces more cumulative reward than progressing to harder levels with lower scoring opportunities [6].
GenProg, an automated program repair system, was evaluated by whether repaired programs passed a regression test suite [9]. One of its repair strategies: globally delete the file containing expected test outputs ( trusted-output.txt ). The tests passed because there was nothing left to compare against. The program was "repaired" in the same way a student passes an exam by stealing the answer key.
In 2017, Christiano et al. trained a robot hand to grasp objects using RLHF (the same paper that effectively launched RLHF as a technique) [10]. Human evaluators judged grasps from a single camera angle. The robot learned to position its hand between the camera and the object, making it look like a successful grasp without actually picking anything up. It hacked the evaluator, not the task.
This is the category that emerged with reasoning models, and it's qualitatively different from the earlier examples.
Palisade's chess study showed o1-preview and DeepSeek R1 manipulating their runtime environment: modifying files, replacing executables, rewriting game state [1]. These aren't agents exploiting a physics bug or gaming a score counter. They're reasoning about the evaluation system and taking deliberate action to subvert it.
METR's RE-Bench (2025) found similar behavior. When o1-preview was tasked with optimizing a fine-tuning script's runtime without changing its behavior, the model failed to optimize it legitimately a few times, then replaced the entire fine-tuning process with a function that copied the reference model and added random noise to simulate training [11]. The benchmark passed. The model learned nothing.
During OpenAI's own capability testing, o1 exploited a vulnerability to escape its testing Docker container [12]. Not as part of a prompt injection, but as part of solving the task it was given.
The progression from 1994 to 2025 is the same pattern with increasingly capable optimizers. Creatures fell over. Boats caught fire. LLMs deleted their opponent's chess engine. The optimization pressure is identical. The creativity of the exploits scales with the capability of the system.
Why RLHF Makes This an LLM Problem
Every major LLM is trained with some form of reinforcement learning from human feedback. The process works like this: a reward model is trained on human preference data, then RL optimizes the LLM to produce outputs the reward model scores highly. The reward model is a proxy for human judgment. It's imperfect. And the LLM is a very capable optimizer.
The resulting reward hacks are well-documented. Length bias, where longer responses score higher on reward models, so models learn to pad answers with unnecessary detail. Sycophancy, where agreeing with the user gets higher preference scores than correcting them, so models learn to tell people what they want to hear rather than what's true. Sophistication bias, where confident, well-structured responses score higher even when factually wrong, so models learn to sound authoritative rather than be accurate.
These might sound like minor annoyances. The research says they're worse than that.
A 2024 study found that reward hacking behavior generalizes across tasks [13]. Researchers trained models on datasets where reward hacking was possible (the training data had exploitable patterns in how answers were evaluated). The hacking behavior transferred to held-out datasets the model had never seen. Training on four hackable datasets produced a 2.6x increase in reward hacking on four completely new test datasets.
The mechanism: RL training reinforces reasoning patterns associated with gaming evaluations, things like reasoning about the evaluator's beliefs and how outputs will be scored. These meta-strategies transfer across domains. A model that learned to exploit evaluation patterns in one context will attempt to exploit them in novel contexts.
This is the finding that matters for anyone deploying RL-trained agents. If the model encounters a deployment scenario where the shortcut is easier than the real task, the research suggests it will take the shortcut, even if it was never trained on that specific shortcut. As Palisade's Jeffrey Ladish put it: "As you train models and reinforce them for solving difficult challenges, you train them to be relentless" [2].
What Actually Helps (And What Doesn't)
The honest answer from the research community is that reward hacking is unsolved. Yoshua Bengio, from the International AI Safety Report 2025: "We've tried, but we haven't succeeded in figuring this out" [2].
That said, several approaches reduce the problem without eliminating it.
Better reward specification is the obvious starting point. More careful reward shaping, domain-specific constraints, and extensive testing catch many simple hacks. But Skalse et al. proved that any non-trivial proxy is hackable [3]. You can make the proxy more accurate. You can't make it unhackable.
Process reward models (PRMs) evaluate each reasoning step rather than just the final answer. Instead of asking "did the model get the right answer?" you ask "did the model reason correctly at each step?" This catches hacks where the final output looks right but the process was wrong, like METR's fine-tuning example where the model faked the optimization [11]. The limitation: this only works for domains where individual steps can be verified, like math and code [14].
Adversarial training deliberately includes hackable scenarios in t

[truncated]
