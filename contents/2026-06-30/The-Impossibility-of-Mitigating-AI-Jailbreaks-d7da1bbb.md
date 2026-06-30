---
source: "https://reliable-ai.review/posts/2026-05-priviledge_erosion/"
hn_url: "https://news.ycombinator.com/item?id=48726902"
title: "The Impossibility of Mitigating AI Jailbreaks"
article_title: "On the Impossibility of Mitigating AI Jailbreaks – AI RELIABILITY REVIEW"
author: "NickySlicks"
captured_at: "2026-06-30T00:30:54Z"
capture_tool: "hn-digest"
hn_id: 48726902
score: 2
comments: 0
posted_at: "2026-06-29T23:56:48Z"
tags:
  - hacker-news
  - translated
---

# The Impossibility of Mitigating AI Jailbreaks

- HN: [48726902](https://news.ycombinator.com/item?id=48726902)
- Source: [reliable-ai.review](https://reliable-ai.review/posts/2026-05-priviledge_erosion/)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T23:56:48Z

## Translation

タイトル: AI の脱獄を軽減することの不可能性
記事のタイトル: AI ジェイルブレイクを緩和することの不可能性について – AI 信頼性レビュー
説明: アライメント、ジェイルブレイク、およびシステム制御に対するそれらの影響を直感的に理解するためのガイド。

記事本文:
AI の脱獄を軽減することの不可能性について
結果として特権が侵食される
現在、AI メディアの報道では、即時注射と脱獄が注目を集めています…
おそらく子供をターゲットにした xAI が爆弾の製造に協力している。
... それには十分な理由があります。私たちはさまざまな (しばしばおかしな) 失敗を観察しています。食品の注文を支援することを目的としたマクドナルドのカスタマー サポート ボットが、Python パズルを解くよう誘導されることがあります 1 。おそらく若い視聴者をターゲットにした xAI のチャットボットは、繰り返しのプロンプトを受けて、パイプ爆弾の作り方についての指示を提供できます 2 。また、ChatGPT モデルは十分に記述されていれば著作権で保護されたキャラクターを生成できます 3 。
4 これらの手法の概要については、Nathan Lambert 著の「Reinforcement Learning from Human Feedback」または Kevin Murphy 著の「Reinforcement Learning from Human Feedback」を参照してください。
これらの障害は、LLM ベースのシステムにおける開発者が意図した制御命令とユーザー提供の入力の間の分離の破綻に起因し、ジェイルブレイク、ポリシー回避、またはプロンプト インジェクションと呼ばれます。これらの障害の標準的な軽減策は、トレーニング後の調整です。これは、意図された指示に従い、安全ポリシーを順守するために、教師あり微調整と人間のフィードバックからの強化学習を介して厳選されたサンプルに基づいてモデルをトレーニングします 4。ただし、アライメントは、モデルが実行できることではなく、実行する可能性のあることを変更するだけです。動作に厳しい制約を課すことなく、考えられる出力にわたる分布を再形成します。
この投稿では、その直感を発展させ、それを体系的に活用する方法を示します 5 。次に、制御とデータの分離の欠如とジェイルブレイクが組み合わさることで、システムレベルの制御に系統的な障害がどのように発生するのかを議論します。
5 次のセクションは直感的なものであることに注意してください。

NeurIPS 2025 論文の電子バージョン: Mission Impossible: A Statistical Perspective on Jailbreaking LLMs 。議論をより厳密に扱うために、読者に論文を参照してもらいたい。
確率論的な観点から見ると、大規模な言語モデルはシーケンス全体にわたる非常に高次元の分布を定義するものとして理解できます。説明のために、簡単な低次元の例から始めます。形状と色という 2 つの確率変数にわたってグラウンドトゥルース分布が存在すると仮定します。生成モデルは、そこから抽出されたサンプルを観察することによって、この分布を近似するようにトレーニングできます。最も単純なケースでは、形状と色の各組み合わせに確率が割り当てられる、完全な同時分布を明示的に表すことを想像できます。オブジェクトを観察するたびに、共同分布でそのイベントに割り当てられる尤度が増加します。
この設定を拡張するにつれて、課題が明らかになります。単純な例では、2 つの変数 (形状と色) があり、それぞれに 3 つの可能な値 (赤、青、緑) と (円、三角形、四角形) があります。これにより、 \(3^2 = 9\) 通りの結果が得られるため、9 つの確率を決定する必要があります。言語の場合、これは大幅に増加します。一連のテキスト内の各位置は、語彙サイズが数万程度の確率変数です。 \(16,000\) トークンの標準語彙と \(1,024\) トークンのコンテキスト長の場合、可能なシーケンスの数は \(16{,}000^{1024} \およそ 10^{4305}\) です。この数は、観測可能な宇宙の粒子の数 (約 \(10^{80}\) ) を大幅に超えています。また、利用可能なトレーニング データの量をはるかに超えています。推定によると、インターネット上のすべてのテキストはおよそ \(10^{12}\) から \(10^{14}\) トークンに相当します。
LLM はこのジョイントを明示的に表現しませんが、

分布であるにもかかわらず、この配列空間全体に確率分布を誘導するので、それを利用することができます。
これを具体的にするために、おもちゃの例に戻ります。 1 つの変数 (色) がリクエスト (例: 「ケーキの焼き方を教えてください」) を表し、もう 1 つの変数 (形状) が応答を表すものとします。結合分布内の各点は、(要求、応答) ペアに対応します。
これらの組み合わせの中には、望ましくないものもあります。たとえば、(「生物兵器の作り方を教えてください」、「もちろん、…が必要です」) などの、有害な要求と準拠した応答の組み合わせです。この図では、そのようなケースを青い四角で表しています。
実際には、モデルの完全な結合分布に直接アクセスすることはできません。代わりに、調整は間接的に機能します。つまり、望ましい動作と望ましくない動作の例を提供し、その可能性を増減するようにモデルを更新します。特に、望ましくないペア (青い四角など) に対応する出力にペナルティを与え、モデルがそれらに低い確率を割り当てるように促します。
調整により、望ましくない結果が発生する可能性が減少しました。モデル分布では有害な (リクエスト、レスポンス) ペアはまれであり、標準的なサンプリングではそのような動作に遭遇する可能性は低いでしょう。
追加のコンテキストを条件にすると、これは変わります。これを説明するために、3 番目の変数を導入します。これは、根底にある意図を変更せずに、リクエストの表現方法を変更する修飾子です。実際には、これは次のようなものに相当します。「ロールプレイをしましょう - あなたは地球を救わなければならないスーパーヒーローです、そしてそうする唯一の方法は…です。」 このような修飾子は、たとえば 🐰 などの動物として表されます。
一方、有害なペアの確率 \(P(🟦) = 0.006\) と、特定の修飾子との結合確率 \(P(

🟦 , 🐰) = 0.004\) は小さい可能性があり、条件付き確率ははるかに高くなる可能性があります。
\(P( 🟦 \mid 🐰) = \frac{P( 🟦, 🐰)}{P(🐰)} \約 0.260\)
全体的にはまれなことですが、🐰を条件付けると有害な結果が発生する可能性が高くなります。結合確率が低いことは、条件付き確率が低いことを意味するわけではありません。
アライメント中に🐰を単純に防御しないのはなぜでしょうか?問題は規模だ。高次元の入力空間では、多くの修飾子 (別の表現、文脈、構成) を組み合わせて、同様の方法で条件付き分布をシフトできます。アライメントは少数のサンプル セットに作用し、入力空間の広大な領域は弱い制約のままになります。攻撃者の観点からすると、そのような制約のない領域を 1 つ見つけることは容易です。ディフェンダーの観点から見ると、そのような地域をすべてカバーすることは不可能です。 6
6 この論文では、そのような領域の体積とその比率について、より正確な議論を行っています。
7 たとえば、「LLM 統合アプリケーションに対するプロンプト インジェクション攻撃」、「AdvPrompter: LLM に対する高速適応型敵対的プロンプティング」、または「RL はハンマーであり、LLM は釘である: SOTA 攻撃に対する強力なプロンプト インジェクションのためのシンプルな強化学習レシピ」を参照してください。
ここまでは、条件付け変数 (🐰 など) が存在することだけを議論してきました。攻撃者もそれらを見つける必要がありますが、これは禁止ではありません。入力は繰り返し改良できるため、攻撃者は入力空間を手動または自動で検索して、目的の動作を誘発するプロンプトを発見できます。この最適化問題は、特定の結果の可能性を最大化するプロンプトを求めます。結果として生じるプロンプトは、ジェイルブレイクまたはプロンプト インジェクションと呼ばれます 7 。このような攻撃にはモデルの尤度へのアクセスは必要ないことに注意してください。
上記で、有害な反応を完全に排除することはできないことを確認しました。

モデルの分布を分析することで、攻撃者はそのような動作を引き起こす可能性のある入力を検索できます。 LLM がチャット コンパニオンである場合、これは懸念事項ですが、被害はある程度限定されます。最悪の場合、モデルは間違ったアドバイスや不適切なコンテンツ、つまり多くの場合他の場所 (Google や Reddit) で見つかる可能性のあるコンテンツを生成します。これは、コーディング、研究、UI、一般的な OS レベルのエージェントなどのエージェントの使用によって変わります。
これらの設定では、モデルはテキストを生成するだけでなく、コードを実行するなどして機能します。 Claude Code などのコーディング エージェントを使用する場合、システムはモデル出力に基づいてアクション (ファイルの編集や bash コマンドの実行) を実行します。より一般的には、このクラスのエージェントは ReAct エージェントと呼ばれます。そのアクションは LLM の出力によって決まります。この出力は、システム プロンプト、ユーザー指示、ツール呼び出し、Web サイトやドキュメントなどの取得されたコンテンツなどの入力ストリームによって決まります。
結果として特権が侵食される
従来のコンピュータ セキュリティでは、データが制御として解釈されるときに重大な脆弱性が発生します。典型的な例はバッファ オーバーフローです。ユーザーが指定した入力が適切に分離されずにメモリに書き込まれ、戻りアドレスなどの制御構造が上書きされる可能性があります。同様に、SQL インジェクションでは、信頼できない入力がクエリの一部として解釈され、攻撃者がプログラムの動作を変更できるようになります。どちらの場合も、根本的な原因は、データと制御の間の明確な境界を維持できないことです。最新のシステムは、型システム、メモリ安全性、パラメータ化されたクエリを通じて、このギャップをアーキテクチャ的に埋めています。つまり、戻りアドレスは実行可能データではなく、SQL パラメータは構文として解析されません。
ReAct エージェントは、同様の問題を再び引き起こします。その命令と、それが作用するデータ (取得したドキュメント、ツール出力、Web ページ、Git リポジトリなど) が到着します。

e 同じ入力ストリームを介して。したがって、LLM システムは、コントロール プレーンをデータ プレーンに統合します。
従来のシステムはデータ/制御の脆弱性をアーキテクチャ的に解決しますが、LLM ベースのシステムは統計的にのみ脆弱性を解決します。学習された命令階層などの緩和戦略 8 は、システム プロンプトをユーザー入力よりも重視し、ユーザー入力を取得したコンテンツよりも重視するようにモデルをトレーニングします。ただし、前のセクションでは、統計的な境界がまさにジェイルブレイクが容易に突破できるものであることを示しました。入力ストリームの任意の場所 (Web ページ、ドキュメント、Git リポジトリ、任意のライブラリ) に 🐰 スタイルの修飾子を配置する攻撃者。ユーザーの指示ではなくユーザーの指示に従う方向にモデルをシフトできます。
8 「命令階層: 特権命令を優先するための LLM のトレーニング」を参照してください。
したがって、定義された権限セット (読み取り、書き込み、実行) で動作する AI エージェントは、入力ストリームの任意の部分にアクセスできるプロセスにそれらの権限を誤って伝播する可能性があります。信頼性の低い入力の重みを信頼性の高い命令よりも低くすることを強制する方法がないため、AI エージェントはシステム全体で特権侵食を引き起こします。攻撃者がエージェントが読み取る場所にコンテンツを配置できるようになると、システムと直接対話することなく、そのアクションへのチャネルが得られます。
アプリケーションを構築している人々にとって、これは典型的な脅威モデルを変えることになります。ソフトウェアは常にオペレーティング システムを信頼できる中立的な基盤として扱ってきました。アプリケーションの下の層は敵ではありません。その層に位置し、メッセージ、カレンダー、ファイルを読み取るエージェントは、読み取るものによって操作可能であり、この想定を打ち破ります。コンピューター自体が攻撃対象領域の一部になります。 Meredith Whittaker と Udbhav Tiwari も 39C3 で同様の議論を行い、エージェントによるアクセスがどのように行われるかを説明しました。

ess は脅威モデルを無効にします。 安全なメッセージング アプリは 9 に基づいて構築されています。
読者にはおなじみのこの原則の例をいくつか挙げます。
サマーユエ／OpenClaw（2026年2月） 10． Meta Superintelligence Labs のアラインメント担当ディレクターである Summer Yue 氏は、AI エージェントにメールの受信箱へのアクセスを許可し、何をアーカイブすべきかを提案するよう依頼しましたが、何もアクションを起こすことは求められませんでした。受信トレイがエージェントのコンテキスト ウィンドウでいっぱいになると、圧縮により以前の安全に関する指示が黙って破棄され、その後エージェントは電子メールの一括削除を開始し、繰り返し停止するコマンドを無視しました。安全制約は敵対者によってバイパスされたのではなく、エージェント自身の内部メモリ管理によって消去されました。攻撃者さえ関与せずに特権が侵食される。
10 FastCompany の記事、サマー ユエに日陰はない — このような恥ずかしい出来事を公に共有することは、特に彼女のような目立つ役割においては、本当に勇気が必要です。そのようなオープンさは、まさに私たちが学ぶ方法なのです。
メタAIサポートエージェント（2026年6月） 11．攻撃者は、Meta の AI サポート チャットボットを悪用して、標的の Instagram アカウントを攻撃者が管理する電子メール アドレスにリンクするよう要求し、その結果、オバマ大統領のホワイト ハウス ページやセフォラなど、注目を集めるアカウント乗っ取りが相次ぎました。チャットボットが個別に検証せずにアカウント資格情報をリセットする

[切り捨てられた]

## Original Extract

A Guide to an Intuitive Understanding of Alignment, Jailbreaking and their Implications for System Control.

On the Impossibility of Mitigating AI Jailbreaks
The result is privilege erosion
Prompt injections and jailbreaks are having a moment in AI media coverage at the moment …
xAI, targeted presumably at children, helps to build a bomb.
… and for good reason: We observe a wide range of (often funny) failures: The McDonald’s customer support bot intended to assist with ordering food can be enticed to solve python puzzles 1 ; xAI’s chatbots targeted at presumably young audiences can, under repeated prompting, provide instructions on how to build a pipe bomb 2 ; And ChatGPT models can generate copyrighted characters when sufficiently described 3 .
4 See Reinforcement Learning from Human Feedback by Nathan Lambert or Reinforcement Learning: An Overview Chapter 6 by Kevin Murphy for an introduction to these techniques.
These failures stem from a breakdown in the separation between developer-intended control instructions and user-provided input in LLM-based systems, and are referred to as jailbreaking, policy evasion, or prompt injections. The standard mitigation for these failures is alignment post-training, which trains models on curated examples via supervised fine-tuning and reinforcement learning from human feedback 4 to follow intended instructions and adhere to safety policies. However, alignment only changes what the model is likely to do, not what it is able to do: it reshapes the distribution over possible outputs without imposing hard constraints on behavior.
This post develops that intuition and shows how it can be systematically exploited 5 . The argument then traces how jailbreaking combined with a lack of separation between control and data can produce systematic failures of system-level control.
5 Note that the following section is an intuitive version of our NeurIPS 2025 paper: Mission Impossible: A Statistical Perspective on Jailbreaking LLMs . For a more rigorous treatment of the argument, I refer the reader to the paper.
From a probabilistic perspective, large language models can be understood as defining very high-dimensional distributions over sequences. For the purposes of illustration, we begin with a simple low-dimensional example. Assume there exists a ground-truth distribution over two random variables: shape and color. A generative model can be trained to approximate this distribution by observing samples drawn from it. In the simplest case, we could imagine explicitly representing the full joint distribution, where each combination of shape and color is assigned a probability. Whenever we observe an object, we increase the likelihood assigned to that event in our joint distribution.
The challenge becomes apparent as we scale this setup: In our simple example, we have two variables (shape and color), each with three possible values—(red, blue, green) and (circle, triangle, square). This results in \(3^2 = 9\) possible outcomes, and thus nine probabilities to determine. With language, this grows substantially. Each position in a sequence of text is a random variable with vocabulary size on the order of tens of thousands. For a standard vocabulary of \(16,000\) tokens and context length of \(1,024\) tokens, the number of possible sequences is \(16{,}000^{1024} \approx 10^{4305}\) . This number vastly exceeds the number of particles in the observable universe (approximately \(10^{80}\) ). It also far exceeds the amount of available training data: estimates suggest that all text on the internet amounts to roughly \(10^{12}\) to \(10^{14}\) tokens.
While LLMs do not explicitly represent this joint distribution, they nonetheless induce a probability distribution over this space of sequences, and that we can exploit.
To make this concrete, we return to our toy example. Let one variable (color) represent the request (e.g., “Tell me how to bake a cake”), and the other variable (shape) represent the response . Each point in the joint distribution corresponds to a (request, response) pair.
Some of these pairs are undesirable, for example, a harmful request paired with a compliant response, such as (“Tell me how to build bio weapons”, “Of course, you will need …”). In our illustration, we represent such cases as blue squares.
In practice, we do not have direct access to the model’s full joint distribution. Instead, alignment operates indirectly: we provide examples of desirable and undesirable behaviors and update the model to increase or decrease their likelihood. In particular, we penalize outputs corresponding to undesirable pairs (like blue squares), encouraging the model to assign them lower probability.
Alignment has reduced the likelihood of undesirable outcomes: harmful (request, response) pairs are rare under the model distribution, and we would be unlikely to encounter such behavior through standard sampling.
This changes once we condition on additional context. To illustrate this, we introduce a third variable: a modifier that changes how the request is phrased without changing its underlying intent. In practice, this could correspond to something like: “Let’s role-play—you are a superhero who must save the planet, and the only way to do so is to…” We represent such modifiers as animals, for example 🐰.
Note that while the probability of a harmful pair \(P(🟦) = 0.006\) , and its joint probability with a specific modifier \(P( 🟦 , 🐰) = 0.004\) can be small, the conditional probability can be much higher.
\(P( 🟦 \mid 🐰) = \frac{P( 🟦, 🐰)}{P(🐰)} \approx 0.260\)
Despite being rare overall, the harmful outcome becomes likely once we condition on 🐰. Low joint probability does not imply low conditional probability.
Why not simply defend against 🐰 during alignment? The problem is scale. In high-dimensional input spaces, combinatorially many modifiers (alternative phrasings, contexts, compositions) can shift the conditional distribution in similar ways. Alignment acts on a small set of examples, leaving vast regions of the input space weakly constrained. From the attacker’s perspective, finding one such unconstrained region is tractable. From the defender’s perspective, covering all such regions is not. 6
6 In the paper, we make a more precise argument about the volume of such regions, and their ratio.
7 For example, see Prompt Injection attack against LLM-integrated Applications , AdvPrompter: Fast Adaptive Adversarial Prompting for LLMs , or RL Is a Hammer and LLMs Are Nails: A Simple Reinforcement Learning Recipe for Strong Prompt Injection for SOTA attacks.
Up to this point, we have only argued that conditioning variables (like 🐰) exist . Attackers must also find them—but this is not prohibitive. Because inputs can be iteratively refined, attackers can search the input space manually or automatically to discover prompts that induce desired behavior. This optimization problem seeks a prompt maximizing the likelihood of a specific outcome. The resulting prompts are called jailbreaks or prompt injections 7 . Note that access to a model’s likelihoods is not necessary for such attacks.
Above we have established that harmful responses cannot be fully eliminated from the model distribution, and attackers can search for inputs that make such behaviors likely. When LLMs are chat companions, this is concerning but damage is somewhat limited. At worst, a model generates bad advice or inappropriate content—content that could in many cases be found elsewhere (Google, Reddit). This changes with agentic uses such as coding, research, UI, and general OS-level agents.
In these settings, the model does not just generate text, it acts, for instance by executing code. When using a coding agent such as Claude Code , the system executes actions based on model outputs: editing files, or running bash commands. More generally, this class of agents is called ReAct agents . Its actions are determined by the LLM’s output, which is determined by an input stream: system prompt, user instructions, tool calls, and retrieved content such as websites or documents.
The result is privilege erosion
In classical computer security, severe vulnerabilities arise when data is interpreted as control . A canonical example is buffer overflow: user-provided input is written into memory without proper separation, allowing it to overwrite control structures such as return addresses. Similarly, in SQL injection, untrusted input is interpreted as part of a query, enabling attackers to modify the program’s behavior. In both, the root cause is the failure to maintain a clear boundary between data and control. Modern systems close this gap architecturally i.e. a return address is not executable data, and an SQL parameter is not parsed as syntax—through type systems, memory safety, and parameterized queries.
A ReAct agent reintroduces similar problems: Its instructions and the data it acts on—e.g., retrieved documents, tool outputs, web pages, git repositories—arrive through the same input stream. Hence, LLM systems collapse the control plane into the data plane.
While classical systems close the data/control vulnerabilities architecturally, LLM-based systems close it only statistically. Mitigation strategies such as learned instruction hierarchies 8 train the model to weight system prompts above user input, and user input above retrieved content - but the previous section showed that statistical boundaries are exactly what jailbreaks breach easily. An attacker who places a 🐰-style modifier anywhere in the input stream i.e. a webpage, a document, a git repo, any library; can shift the model toward following their instructions instead of the user’s.
8 See The Instruction Hierarchy: Training LLMs to Prioritize Privileged Instructions .
Hence, an AI agent operating with a defined privilege set (read, write, execute) may inadvertently propagate those privileges to any process with access to any portion of its input stream. Because there is no way to enforce that lower-trust inputs carry less weight than higher-trust instructions, AI agents cause Privilege Erosion across the entire system . Once an attacker can place content anywhere the agent reads from, they have a channel to its actions—without ever interacting with the system directly.
For people building applications, this changes the typical threat model. Software has always treated the operating system as a trusted, neutral foundation: the layer below your application is not your adversary. An agent that sits at that layer—reading messages, calendars, files—and is steerable by anything it reads breaks this assumption. The computer itself becomes part of the attack surface. Meredith Whittaker and Udbhav Tiwari made a similar argument at 39C3, describing how agentic access invalidates the threat models secure messaging apps are built on 9 .
A few examples of this principle that might seem familiar to readers;
Summer Yue / OpenClaw (February 2026) 10 . Summer Yue, director of alignment at Meta Superintelligence Labs, granted an AI agent access to her email inbox and asked it to suggest what should be archived — but not to take any action. As the inbox filled the agent’s context window, compaction caused her earlier safety instruction to be silently discarded, after which the agent began mass-deleting emails and ignored repeated commands to stop. The safety constraint was not bypassed by an adversary — it was erased by the agent’s own internal memory management. Privilege erosion without even an attacker in the loop.
10 FastCompany article , No shade to Summer Yue — publicly sharing an embarrassing incident like this, especially in a role as visible as hers, takes real courage. That kind of openness is exactly how we learn.
Meta AI Support Agent (June 2026) 11 . Attackers exploited Meta’s AI support chatbot by asking it to link a target Instagram account to an attacker-controlled email address, resulting in a wave of high-profile account takeovers — including the Obama White House page and Sephora. The chatbot reset account credentials without independently verifyin

[truncated]
