---
source: "https://viborc.com/what-recursive-self-improvement-actually-means/"
hn_url: "https://news.ycombinator.com/item?id=49333469"
title: "Has AI started improving itself? RSI depeneds on what \"itself\" means"
article_title: "What recursive self-improvement actually means"
image: "https://viborc.com/what-recursive-self-improvement-actually-means/recursive-system-boundary-og.jpg"
author: "viborcip"
captured_at: "2026-08-17T17:18:26Z"
capture_tool: "hn-digest"
hn_id: 49333469
score: 1
comments: 0
posted_at: "2026-08-17T16:17:40Z"
tags:
  - hacker-news
  - translated
---

# Has AI started improving itself? RSI depeneds on what "itself" means

- HN: [49333469](https://news.ycombinator.com/item?id=49333469)
- Source: [viborc.com](https://viborc.com/what-recursive-self-improvement-actually-means/)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T16:17:40Z

## Translation

タイトル: AI は自らを改善し始めたのか? RSI は「それ自体」が何を意味するかによって決まります
記事のタイトル: 再帰的自己改善が実際に意味するもの
説明: 再帰的自己改善とは、AI が改善能力を向上させることです。一部の現在のシステムはそのループの一部を閉じていますが、私がレビューした公開ケースでは完全な後継サイクルを実行するものはありません。

記事本文:
コンテンツにスキップ
viborc.com トピックス
研究
メニュートピックスについて
研究
について
人間の判断
発展途上のエッジ
システムとデータ
Power and synthesis Artificial Intelligence
Has AI started improving itself?
再帰的自己改善とは、潜在的により優れた後継者を構築することによって、AI が次の改善を行う能力が向上するという考えです。私は、そのループのどのくらいがすでに現実のもので、どのくらいがモデルの周りの人々やインフラストラクチャにまだ属しているのかを知りたかったのです。
人工知能 · インテリジェンス & トレードクラフト
On this page The short version
What recursive self-improvement means
First decide what counts as the system
Models can already improve an answer
ダーウィン ゲーデル マシン: 凍結モデルの周りの再帰
AlphaEvolve: 答えがチェック可能になると何が変わるか
自己報酬型言語モデル: 提案者と裁判官が同じ家族の出身である場合
The evaluator decides what counts as improvement
Recursion does not imply accelerating progress
AI は、限られた方法で、周囲のシステムの一部を改善し始めています。しかし、独自に構築、検証、展開し、より有能な後継システムから学習する一般的なシステムの公開例は見つかりませんでした。私たちはまだそこに到達していません。
したがって、企業が自社の AI 自体が改善されたと言う場合は、何が変わったのか、誰が結果を判断したのか、何が保存されたのか、そして変更されたシステムが次のラウンドを実行したかどうかを尋ねてください。多くの場合、評価者は主張が最も弱くなる場所です。
再帰的自己改善は、AI の議論では通常 RSI と短縮されます。
AI システムがそれ自体の能力を向上させ、さらなる改善を行うという考えです。
最も強力なバージョンでは、より有能な後継機が設計され、それが役立つようになります。
design the version after that. If those gains keep strengthening the
改善プロセスでは、ループが加速して次のような状態になる可能性があります。

I. J. グッドは
「知性の爆発」。グッド氏がこの議論を提案したのは 1965 年に遡ります。
当時は推測の域を出ませんでしたが、まだ既存のものを説明したものではありません。
エンドツーエンドのシステム。 1
60年後の2025年7月30日、マーク・ザッカーバーグはメタが「始まった」と書いた。
私たちの AI システムが自ら改善している様子を垣間見ることができました。」彼は、
改善は遅いが否定できない、当時は超知能が導入されたと言われていた
光景。 2
製品の発表で、同じ安全性を主張するバージョンを見たことがありました。
フレームワーク、論文、そして私が毎日行っているコーディング エージェントの仕事です。人類のレポート
2026 年 5 月の時点で、コードの 80% 以上をクロードが作成し、
コードベース。 OpenAIは、自動化されたGPT-Redモデルが生成した攻撃をトレーニングに使用したと述べている
GPT-5.6 は即時注射に対して有効です。それでも、Anthropic は完全再帰的とも言います
自己改善は到着していません。 OpenAI の準備フレームワークでは、
評価された GPT-5.6 モデルは AI の上限しきい値に達しませんでした
自己改善。 METR は RSI を専門用語として避けています。
相容れないもの。 3
それらの情報源は、自己改善をさまざまな意味で使用していました。私も
もっと普通の説明の可能性も排除できませんでした。言語の一部が行っていたのです。
マーケティングの仕事。それは当然のことですが、正確に何をするのか知りたくなりました
主張の後ろに座っていた。そこで私は調査を行いました。
再帰的自己改善の意味
事例を検討する前に、1 つのルールが必要でした。AI システムは改善できるということです。
それ自体を再帰的に改善しない何か。 3 つの一般的な例を示します。
なぜ:
回答を批判して書き直すモデルにより、出力が向上しました。
有用なテクニックを記憶したシステムは、次回は異なる動作をする可能性があります。
ツールまたはコードを編集するエージェントが機構の一部を変更した
模型の周り。
パーシステンスゲートはchです

アンジェは保存して、後の実行で再利用しますか?いいえ 改善のみ 変更はこの出力で終了します。はい - もう 1 回のテスト次のラウンドで改善プロセスを再度変更できますか?いいえ 永続的な変更 はい 再帰ループ より良い出力は依然として改善の余地があります。メモリとツールの変更は実行後も保持される可能性があります。このループは、変更されたシステムが、後の改善方法を変更できる別の改善ラウンドに参加する場合にのみ再帰的になります。方法: これは、エッセイ全体で使用される永続性と再利用のテストを適用します。この作品では、より強力なステップ、つまり受け入れられた変更のために再帰を使用しています。
存続すると、変更されたシステムは別の改善ラウンドに参加し、
ラウンドによってプロセスが再び変更される可能性があります。
古い AGI の議論は、そのスペクトルの遠端から始まります。
より有能な後継者と潜在的な知性の爆発を設計しています。
現在の工学論文では、固定境界内の小さなループを研究することがよくあります。
一方、AI ラボでは、エージェントや自動化をコーディングするときにこのフレーズをより広範囲に使用します。
実験は、人間が運営する組織が次のモデルを構築するのに役立ちます。これらすべて
フィードバックは含まれますが、AI に同じ役割を与えるわけではありません。
AI は、周囲のシステムの一部を改善し始めています。ダーウィン ゲーデル マシン
世代を超えてコーディング エージェントの足場を変更します。自分にご褒美を与える言葉
モデルは、外部で実行されるトレーニング ループでモデルによって生成された判断を使用します。
AlphaEvolve は、Gemini のトレーニングに使用されるコードを含むターゲット プログラムを改善します。 4
しかし、研究目標を選択し、より多くの研究を構築する公的システムは見つかりませんでした。
有能な一般的な後継者は、それを独自に検証し、デプロイしてから、
後継者は、重要な部分を閉じることなく次のサイクルを実行します。
ループ。
答えは、それ自体に何が含まれているかによって異なります。それはモデルの重みを意味するかもしれません、
ラップされたエージェント

モデル、または会社の声明では研究室全体を中心に:
人、モデル、ベンチマーク、トレーニング クラスター、展開の決定。いつ
AI は自らを改良したと誰かが言いますが、AI の内部には一体何があるのでしょうか?
まず、何をシステムとしてカウントするかを決定します
言語モデルのみを意味する場合、そのプロンプトを編集する人は言語モデルの外にいます。
境界線。モデルとそのエージェント ハーネス、ツール、メモリ、および
評価者が評価を行っている場合でも、ハーネスの変更は自己改善になります。
モデルの重みは固定されたままになります。トレーニング インフラストラクチャ、チェックポイントを追加する
選択、安全性レビュー、展開が行われ、自律的に見えたループは、
いくつかの重要な点で人に依存していることが判明しました。
すべての質問に対して単一の正しい境界はありません。トラブルが始まるのは、
発表では、成果に対して 1 つの境界を使用し、成果に対してはより大きな境界を使用します。
主張する。
私はどの事件でも同じ事実に従います。議論されているシステム、
提案された変更、その変更を誰が判断したか、何が保存され、誰が参加したか
次のラウンドで。この最後の事実により、改良された製品と既存の製品が区別されます。
改善された改良剤。
たとえば、モデルを変更することなく、より高速な並べ替えプログラムを設計できます。
モデル、検索アルゴリズム、またはそれを作成した評価者。プログラム
改善されました。それを見つけたシステムはまったく同じである可能性があります。
3 秒後に使用されるより良い答えは、1 つの会話内でのフィードバックです。
ただし、実行全体で改善するには、何かを保存する必要があります。それは可能性があります
メモリ、コードコミット、または後の実行で実際に使用されるチェックポイント。
水平にスクロールして図全体を読んでください
宣言されたシステム境界内では、現在のシステムがプロポーザーを駆動し、変更候補を作成します。評価者または検証者は、選択および安全ゲートに証拠を提供します。受け入れられた変更は永続ストアに入ります。

次のシステムをインスタンス化し、次のサイクルで再利用されます。人間の研究課題、タスクの配分、評価者とベンチマークの作成者、コンピューティング プロバイダー、トレーニング インフラストラクチャ、専門家の検証、展開の承認は、境界の外側に表示されたままになります。 DGM は、モデルの重みと外側の選択を固定したまま、有界のエージェント コード ループを閉じます。 AlphaEvolve は、限定されたターゲット プログラム検索ループを閉じますが、独自のオーケストレーターは変更しません。自己報酬型言語モデルは新しいチェックポイントを保持しますが、生成と判断は系統を共有し、トレーニング操作は外部に残ります。
ループを呼び出す前に境界を宣言する 自律 自動化された内部境界 外部または人的必須のコミット 宣言されたシステム境界 外部 研究課題 タスク配分 外部 評価者作成者 ベンチマーク管理者 外部 コンピューティング / API トレーニング インフラストラクチャ 外部 専門家による検証 導入の承認 1 · 現在のシステム S t 2 · 提案者を実行 3 · 修正候補変更 4 ・ TEST / PROVE 評価者/検証者 5 ・ DECIDE 選択 + セーフティゲート 6 ・ ACCEPT / ROLLBACK 永続ストア 7 ・ INSTANTIATE 次のシステム S t+1 機能の生成 証拠の評価 コミット 次のサイクルで再利用！提案者と評価者の系統が共有されると、独立した検証が弱まります。同じシーケンスにマッピングされた 3 つのケース ダーウィン ゲーデル マシン AGENT-SCAFFOLD ループの変異 コーディング エージェント リポジトリの検証 固定ベンチマーク + ゲート 永続化 アーカイブされた子リポジトリの外側 FM 重み、外側の選択 U3 / C2 · 部分 C3 AlphaEvolve ARTIFACT-SEARCH LOOP 変異 ユーザー指定のプログラム 検証 ユーザー実行可能評価者 PERSISTS 進化的データベース OUTSIDE タスク、評価者、デプロイメント X1 / C2 · RUNNER FIXED 自己報酬型 LM WEIGHT + JUDGE LOOP MUTATES モデルの重み VERIFIES 同じ系統

モデル ジャッジ PERSISTS チェックポイント M t+1 OUTSIDE DPO、計算、選択 U4 + U5 · C1–C2 BY BOUNDARY 自動化されたシーケンスでも、評価、統合、展開を人に依存することができます。宣言された境界によって、「自分」に何が含まれるかが決まります。システムは、人々がベンチマーク、トレーニング インフラストラクチャ、専門家のレビュー、または展開の決定を提供しながら、変更を提案、評価、保存、再利用できます。ソースと方法: 引用された DGM、AlphaEvolve、および Self-Rewarding Language Models レコードを同じシーケンスにマッピングしました。証拠の締め切り: 2026 年 8 月 15 日。アナウンスでは通常、変更の提案、
それを評価し、勝者を維持し、繰り返します。図の破線の接続
これらのアカウントがしばしば省略しているものを示します。ベンチマーク作成者、モデル提供者、
トレーニングインフラ
[切り捨てられた]
委員会は、生産量の改定、
足場の編集、重量の更新、評価機能の改善、自動調査、
そして後継システムの設計。 AlphaEvolve はそのリストの問題を明らかにしました。
事例を追跡した結果、変更できる可能性があることが 5 つ見つかりました。
永続メモリには独自のカテゴリが必要でした。 「自動化された研究」はそうではありませんでした
別のこと。それはいくつかの変更を調整できるプロセスでした。
「後継構築」は、コンポーネントではなく結果を表します。必要です
新しいシステムが次のシステムを引き継ぐ前に、いくつかの変更が連携して機能するようにする必要があります。
丸い。
ケース全体で、次の 5 か所に変更が加えられました。
模型の周囲の足場、
人が重要な手順を実行することも、システムが提案、テスト、保持することもできます。
人間が設計したプロセス内で自動的に変更されます。しかし、将軍
後継者は、それを製造するプロセスも改善する必要があります。
水平にスクロールしてマトリックス全体を読んでください
アップデート対象は5つ

3 つのプロセス タイプ間で比較しました。出力または作業状態のリビジョンは、ヒューマンエッセンシャルおよび境界のあるエピソード内ループで観察されますが、それ自体は後続ループではありません。永続メモリ、スキャフォールドまたはエージェント コード、およびモデル パラメーターが有界ループで観察されますが、システムの構築と後継システムの再利用の証拠は部分的です。評価者または目的の更新には限界のあるプロトタイプがありますが、一般的な独立して検証された後続ループを実証するレビューされたケースはありません。自動化された調査は、5 つのターゲットすべてを調整できます。後継構築では、複数のターゲットを検証、統合、展開、および次のシステムによる再利用と組み合わせます。
変更内容 / 誰がプロセスを実行するか 観察された 部分的 表示されない テストされなかった カテゴリ、スコアではない 更新ターゲット C1 人間必須 C2 制限された自動化 C3 再帰的後継またはメタ改善 U1 出力 / 作業状態 観察された 人間が選択または修正する 観察される 1 つのエピソード内で観察される 表示されない 出力だけでリセットされる 持続性のしきい値 · 後で再利用 U2 記憶 / スキル観察された キュレートされた永続性 観察された 境界のある環境 部分的 次のシステムを構築 U3 足場 / エージェントコード 観察された 人間の統合 観察された 外部プロセスの固定 部分的 境界のある自己参照 U4 重み / ポリシー 観察

[切り捨てられた]

## Original Extract

Recursive self-improvement is AI improving its ability to improve. Some current systems close parts of that loop, but no public case I reviewed runs the full successor cycle.

Skip to content
viborc .com Topics
Research
About Menu Topics
Research
About
Human judgment
The developing edge
Systems and data
Power and synthesis Artificial Intelligence
Has AI started improving itself?
Recursive self-improvement is the idea that an AI gets better at making the next improvement, potentially by building a better successor. I wanted to know how much of that loop is already real, and how much still belongs to the people and infrastructure around the model.
Artificial Intelligence · Intelligence & Tradecraft
On this page The short version
What recursive self-improvement means
First decide what counts as the system
Models can already improve an answer
Darwin Gödel Machine: recursion around a frozen model
AlphaEvolve: what changes when the answer is checkable
Self-Rewarding Language Models: when proposer and judge come from the same family
The evaluator decides what counts as improvement
Recursion does not imply accelerating progress
In bounded ways, AI has started improving parts of the systems around it. However, I found no public example of a general system that builds, validates, deploys, and then learns from a more capable successor on its own. We are not there yet.
So when a company says its AI improved itself, ask what changed, who judged the result, what was saved, and whether the changed system performed the next round. The evaluator is often where the claim becomes weakest.
Recursive self-improvement, usually shortened to RSI in AI discussions, is the
idea that an AI system improves its own ability to make further improvements.
In the strongest version, it designs a more capable successor, which then helps
design the version after that. If those gains keep strengthening the
improvement process, the loop could accelerate into what I. J. Good called an
“intelligence explosion.” Good proposed that argument all the way back in 1965.
It was speculative then, and it is still not a description of an existing
end-to-end system. 1
Sixty years later, on 30 July 2025, Mark Zuckerberg wrote that Meta had “begun
to see glimpses of our AI systems improving themselves.” He called the
improvement slow but undeniable, then said superintelligence was now in
sight. 2
I had seen versions of the same claim in product announcements, safety
frameworks, papers, and the coding-agent work I do every day. Anthropic reports
that, as of May 2026, Claude authored more than 80% of the code merged into its
codebase. OpenAI says its automated GPT-Red model generated attacks used to train
GPT-5.6 against prompt injections. Yet Anthropic also says full recursive
self-improvement has not arrived. Under OpenAI’s Preparedness Framework, the
evaluated GPT-5.6 models did not reach its High threshold for AI
self-improvement. METR avoids RSI as a technical term because people use it for
incompatible things. 3
Those sources were using self-improvement to mean different things. I also
could not rule out a more ordinary explanation: some of the language was doing
marketing work. That is legitimate, but it made me want to know exactly what
sat behind the claim. So I did the research.
What recursive self-improvement means
Before I looked at the cases, I needed one rule: an AI system can improve
something without recursively improving itself. Three common examples show
why:
A model that critiques and rewrites an answer has improved an output.
A system that remembers a useful technique can behave differently next time.
An agent that edits its tools or code has changed part of the machinery
around the model.
PERSISTENCE GATE Is the change saved and reused in a later run? NO Improvement only The change ends with this output. YES - ONE MORE TEST Can the next round change the improvement process again? NO Persistent change YES Recursive loop A better output is still an improvement. Memory and tool changes can persist across runs. The loop becomes recursive only when the changed system takes part in another improvement round that can alter how later improvements are made. Method: this applies the persistence-and-reuse test used throughout the essay. In this piece, I use recursive for the stronger step: an accepted change
survives, the changed system takes part in another improvement round, and that
round can change the process again.
The older AGI argument starts at the far end of that spectrum, with a system
designing a more capable successor and a possible intelligence explosion.
Current engineering papers often study smaller loops inside fixed boundaries,
while AI labs use the phrase more broadly when coding agents or automated
experiments help the human-run organization build the next model. All of these
involve feedback, but they do not give the AI the same role.
AI has started improving parts of the systems around it. Darwin Gödel Machine
changes a coding-agent scaffold across generations. Self-Rewarding Language
Models use model-generated judgments in an externally run training loop.
AlphaEvolve improves target programs, including code used to train Gemini. 4
But I did not find a public system that chooses the research goal, builds a more
capable general successor, independently validates it, deploys it, and then lets
that successor run the next cycle without people closing essential parts of
the loop.
The answer depends on what itself includes. It might mean the model weights,
an agent wrapped around a model, or, in a company statement, a whole lab:
people, models, benchmarks, training clusters, and deployment decisions. When
someone says an AI improved itself, what exactly is inside itself ?
First decide what counts as the system
If I mean only the language model, a person editing its prompt is outside the
boundary. If I mean the model plus its agent harness, tools, memory, and
evaluator, then a change to the harness can be self-improvement even while the
model weights stay frozen. Add the training infrastructure, checkpoint
selection, safety review, and deployment, and a loop that looked autonomous may
turn out to depend on people at several essential points.
There is no single correct boundary for every question. Trouble starts when an
announcement uses one boundary for the achievement and a larger one for the
claim.
I follow the same facts through every case: the system being discussed, the
change it proposed, who judged that change, what was saved, and what took part
in the next round. That final fact separates an improved product from an
improved improver.
For example, a model can design a faster sorting program without changing the
model, the search algorithm, or the evaluator that produced it. The program
improved. The system that found it may be exactly the same.
A better answer used three seconds later is feedback inside one conversation.
But for improvement across runs, something has to be saved. That could be a
memory, a code commit, or a checkpoint that a later run actually uses.
Scroll horizontally to read the full diagram
Inside a declared system boundary, the current system drives a proposer, which creates a candidate change. An evaluator or verifier supplies evidence to a selection and safety gate. Accepted changes enter a persistent store, instantiate the next system, and are reused in the following cycle. Human research agendas, task distributions, evaluator and benchmark authors, compute providers, training infrastructure, expert validation, and deployment approval remain visible outside the boundary. DGM closes a bounded agent-code loop while its model weights and outer selection stay fixed. AlphaEvolve closes a bounded target-program search loop but does not modify its own orchestrator. Self-Rewarding Language Models persist new checkpoints, but generation and judgment share a lineage while training operations remain external.
DECLARE THE BOUNDARY BEFORE CALLING THE LOOP AUTONOMOUS Automated inside boundary External or human-essential Commit DECLARED SYSTEM BOUNDARY EXTERNAL Research agenda Task distribution EXTERNAL Evaluator author Benchmark maintainer EXTERNAL Compute / API Training infrastructure EXTERNAL Expert validation Deployment approval 1 · CURRENT System S t 2 · EXECUTE Proposer 3 · MODIFY Candidate change 4 · TEST / PROVE Evaluator / verifier 5 · DECIDE Selection + safety gate 6 · ACCEPT / ROLLBACK Persistent store 7 · INSTANTIATE Next system S t+1 capability generate evaluate evidence commit reuse next cycle ! Shared proposer and evaluator lineage weakens independent validation. THREE CASES MAPPED TO THE SAME SEQUENCE Darwin Gödel Machine AGENT-SCAFFOLD LOOP MUTATES Coding-agent repository VERIFIES Fixed benchmarks + gates PERSISTS Archived child repository OUTSIDE FM weights, outer selection U3 / C2 · PARTIAL C3 AlphaEvolve ARTIFACT-SEARCH LOOP MUTATES User-designated program VERIFIES User executable evaluator PERSISTS Evolutionary database OUTSIDE Task, evaluator, deployment X1 / C2 · RUNNER FIXED Self-Rewarding LM WEIGHT + JUDGE LOOP MUTATES Model weights VERIFIES Same-lineage model judge PERSISTS Checkpoint M t+1 OUTSIDE DPO, compute, selection U4 + U5 · C1–C2 BY BOUNDARY An automated sequence can still rely on people for evaluation, integration, and deployment. The declared boundary determines what "self" includes. A system may propose, evaluate, save, and reuse a change while people still provide the benchmarks, training infrastructure, expert review, or deployment decision. Source and method: I mapped the cited DGM, AlphaEvolve, and Self-Rewarding Language Models records onto the same sequence. Evidence cutoff: 15 August 2026. Announcements usually describe the automated sequence: propose a change,
evaluate it, keep the winner, and repeat. The dashed connections in the figure
show what those accounts often omit. Benchmark authors, model providers,
training infrastructu
[truncated]
The commission began with six provisional categories: output revision,
scaffold editing, weight updates, evaluator improvement, automated research,
and successor-system design. AlphaEvolve exposed the problem with that list.
After tracing the cases, I ended up with five things that could change.
Persistent memory needed its own category. “Automated research” was not
another thing; it was a process that could coordinate several changes.
“Successor construction” described an outcome, not a component. It requires
several changes to work together before the new system can take over the next
round.
Across the cases, changes landed in five places:
the scaffold around the model,
People may perform essential steps, or a system may propose, test, and retain
changes automatically inside a process people designed. But a general
successor would also have to improve the process that produced it.
Scroll horizontally to read the full matrix
Five update targets are compared across three process types. Output or working-state revision is observed in human-essential and bounded within-episode loops but is not a successor loop by itself. Persistent memory, scaffold or agent code, and model parameters are observed in bounded loops, while evidence for a system building and reusing its successor is partial. Evaluator or objective updating has bounded prototypes, but no reviewed case demonstrates a general independently validated successor loop. Automated research can orchestrate all five targets. Successor construction combines several targets with validation, integration, deployment, and reuse by the next system.
WHAT CHANGES / WHO RUNS THE PROCESS Observed Partial Not shown Not tested CATEGORIES, NOT SCORES UPDATE TARGET C1 Human-essential C2 Bounded automated C3 Recursive successor or meta-improvement U1 Output / working state Observed Human selects or revises Observed Within one episode Not shown Output alone resets PERSISTENCE THRESHOLD · REUSE LATER U2 Memory / skills Observed Curated persistence Observed Bounded environments Partial Builds next system U3 Scaffold / agent code Observed Human integration Observed Fixed outer process Partial Bounded self-reference U4 Weights / policy Observe

[truncated]
