---
source: "https://cruxevals.com/crux/can-ai-agents-conduct-research/"
hn_url: "https://news.ycombinator.com/item?id=49195882"
title: "AI Agents cannot currently conduct open-ended research"
article_title: "Can AI agents conduct open-ended AI research? Early evidence from two case studies | CRUX"
author: "automatic6131"
captured_at: "2026-08-06T12:50:59Z"
capture_tool: "hn-digest"
hn_id: 49195882
score: 1
comments: 0
posted_at: "2026-08-06T12:43:37Z"
tags:
  - hacker-news
  - translated
---

# AI Agents cannot currently conduct open-ended research

- HN: [49195882](https://news.ycombinator.com/item?id=49195882)
- Source: [cruxevals.com](https://cruxevals.com/crux/can-ai-agents-conduct-research/)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T12:43:37Z

## Translation

タイトル: AI エージェントは現在、無制限の研究を行うことができません
記事のタイトル: AI エージェントはオープンエンドの AI 研究を行うことができますか? 2 つのケーススタディからの初期の証拠 |クラックス
説明: 私たちは、2 つの非公開の NeurIPS 提出物からフロンティア エージェントに研究質問を与え、元の著者に結果を採点してもらいました。代理人が作成した書類は明確に拒否されました。

記事本文:
AI期待のアップデートに向けたCRUX共同研究
AI エージェントは無制限の AI 研究を行うことができますか? 2 つのケーススタディからの初期の証拠
私たちはフロンティア エージェントに 2 つの非公開の NeurIPS 提出物からの研究質問を与え、元の著者に結果を採点してもらいました。代理人が作成した書類は明確に拒否されました。
AI の爆発的な進歩の予測は、AI 研究を自動化する AI エージェントにかかっています。しかし、エージェントが無制限の AI 研究を実行できるかどうかについての証拠は乏しい。現在の評価では、限定的で検証可能なタスクでエージェントをテストするか、無制限の研究を除外するか、AI で生成された論文をブラインド査読に提出しますが、これは過剰で確率的であり、レビューの質が低いという問題があります。 AI 研究開発の自動化に向けた進捗状況を測定する 3 番目の方法を紹介します。エージェントは、質の高い未発表論文の中心となる自由回答の研究課題に取り組み、論文の元の著者がその成果を採点します。これらをシャドウ評価と呼びます。私たちは 2 つの未公開の NeurIPS 2026 提出物に対してシャドウ評価を実行し、フロンティア エージェントに 6 日間と数千ドルのコンピューティングを与えました。エージェントは人間の助けなしですべてのエンジニアリングを完了しましたが、研究上の疑問の答えに向けて実質的な進歩を遂げることはできませんでした。その結果、両方の論文は著者によって明確に拒否されました。私たちは、繰り返される 5 つの失敗モードを特定します。それは、出版可能な研究の基準に関する誤った判断、研究設計の欠点に対する創造性のない対応、行き止まりからの非効果的な後戻り、不十分なリソース認識、指示の逸脱です。 2 番目のモデルと足場を使用した堅牢性チェックにより、これらの障害が再現されました。専門家のレビュー、アンケートの回答、エージェントのリポジトリ、およびログを公開します。私たちの結果は、今日の状況が次のようなものであるという初期の証拠を提供します。

男性は AI 研究のエンジニアリングを行うことができますが、研究ライフサイクルの重要な部分で苦労しています。
AI の機能に関する最も重大な未解決の疑問の 1 つは、AI エージェントが AI 研究を実施できるかどうかです。 AI の爆発的な進歩に関する予測の多くは、AI システムが間もなく AI 研究を自ら行うようになるだろうと推測しています。これは、主要な AI ラボの明確な前提でもあります。 6 月に Anthropic は「AI が自らを構築するとき」というタイトルの投稿を公開し、7 月には OpenAI が自社の新しいモデル GPT-5.6 Sol が小規模モデルの事後トレーニングに役立ち、研究者が数週間節約できると宣伝しました。 1 1. 特に、ルナのトレーニング後の GPT-5.6 ソルの貢献は、81 ページのシステム カードには記載されていません。しかし、この研究方向の重要性とこれらの主張に注目が集まっているにもかかわらず、エージェントが自由回答型の研究問題を解決できるかどうかに関する証拠は薄い。
自律型 AI 研究に関する最近の研究のほとんどを統一しているのは、エージェントが固定された狭い指標を改善する必要がある検証可能なタスクに焦点を当てていることです。ベンチマークはエージェントに既知の指標を改善するよう求め、自動検証機能が結果をスコア付けします。ベンチマーク以外にも、GPT-2 レベルのモデルの最適化、より弱いモデルを使用したより強いモデルのトレーニング、自動調査ハーネスの最適化などのタスクにおいて、AI エージェントが専門家の人間のパフォーマンスを上回っていることが、多くの評価で示されています。 2 2. 付録 4 では、これらおよび検証可能なタスクを構成する他の多くの AI 研究実験について説明します。
しかし、AI 研究の多くは検証可能なタスクを解決するだけにとどまりません。エージェントは、一連の候補仮説を選択したり、研究上の疑問を解決する証拠を決定したり、フィードバックを効果的に組み込んでアプローチが失敗し、最初からやり直すのが正しい行動であると認識したりするために、坂を登ることができません。
小さな数字 o

f プロジェクトは、AI 研究を評価するために別の方法を採用しています。つまり、AI で生成された論文を、AI カンファレンスやワークショップなどのブラインド査読プロセスに提出します。しかし、査読は研究能力を測る尺度としては弱い。会議の査読は過剰で確率論的であり、最終的に受理されるまでに AI によって生成された投稿がどれだけ拒否されたかは明らかにされていない。
CRUX は、現実世界の困難なタスクに関して、フロンティア AI システムのオープンエンドかつ長期的な評価を実施するための私たちのプロジェクトです。これは、少数の現実的な課題に焦点を当て、AI システムのパフォーマンスを深く分析することで、ベンチマークよりも最先端の AI システムを前進させます。 3 か月前、私たちは長期的なオープンワールド評価の基礎を説明する論文を書きました。この評価では、エージェントは AI 研究の自由回答形式の質問を解決できるか? と尋ねます。
この論文では、AI エージェントがシャドウ評価と呼ばれる新しい方法でオープンエンドの AI 研究を実行できるかどうかを評価します。これには、まだ公開されていない質の高い研究論文から研究の中心となる質問を取り出し、十分なリソースを備えたフロンティアエージェントにそれに答えるよう依頼し、論文の原著者にエージェントの成果を会議に投稿する場合と同様に採点するよう依頼することが含まれます。エージェントは元の研究を「シャドウ」します。つまり、元の著者の論文や調査結果にはアクセスせずに、元の著者と同じ研究課題に取り組みます。この設計により、自由形式のタスク、汚染のない質問、および評価対象の正確な質問に関する深い専門知識を持つレビュー担当者が得られます。私たちはシャドウ評価が、自動化された AI 研究の限定的で検証可能な評価とブラインドレビュー評価の両方を補完すると考えています。私たちは、将来の研究がこれら 3 つすべてを使用して、戦争への進歩を測定するさまざまな側面を探求することを願っています。

AI研究を自動化するds。
シャドウ評価を実行するために、NeurIPS 2026 に提出されたまだ公開されていない 2 つの論文の著者と提携しました。未発表の論文は、エージェントが自由回答型の研究問題を解決できるかどうかをテストする方法を提供します。元の著者は、数か月に及ぶ質問についての独自の努力を通じて、エージェントの仕事の品質を評価する独自の立場にありますが、エージェントは研究者の結果を検索することができません。なぜなら、研究者はまだウェブ上に存在していないからです。私たちはエージェントに、論文の調査質問、6 日間の実時間、エージェントが 1 週間にわたる無制限の探索と実験の実行を可能にする 3,000 ドルの Anthropic API クレジット、実験用の GPU クレジット、および VM とオープン Web へのフルアクセスを与えました。目標は、一流の AI カンファレンスで発表されるに値する論文を作成することでした。その後、両方の論文の原著者が会議の査読者として結果を採点しました。
私たちの重要な発見は、エージェントは研究を行うために必要なエンジニアリング上の問題を解決できたものの、トップの ML カンファレンスに匹敵するオリジナルの研究を生み出すことができなかったということでした (詳細については、以下の表を参照)。主なポイント:
エージェントには、問題が適切に解決された時期を特定する判断力が欠けていました。彼らは研究上の疑問を理解し、元の論文の著者の意見を忠実に反映した方向性を提案しました。しかし、彼らはその後、手動で厳選された、または合成された小規模なデータセットを使用して、それらの仮説を偽りました。最終論文では、文献との関わりが浅く、実質的な知見としては不十分な否定的な結果を提示しました。
エージェントは、利用可能なリソースとプロジェクトのスケジュールについての認識が不足していました。どちらの実行も、使用された API 予算の 50% 未満で終了しました。

ただし、エージェントは自分の使用状況をリアルタイムで監視でき、残りの予算を使用することが奨励されました。エージェントは、これらのリソース制限、特に時間制限の意味を直観的に理解していないようでした。彼らは最初の探索を数時間かけて急いで行い、論文が独自の成功基準を満たしていなかったにもかかわらず、数時間の時計時間を残して終了しました。
エージェントは、不適切な研究デザインに関するフィードバックに創造的に対応しませんでした。私たちはエージェントに対し、研究の品質と進捗状況を評価するために、AI レビューのために論文をサブエージェントと、refine.ink などの外部ツールの両方に送信するように指示しました。数十回の改訂を通じて、エージェントの自己レビューでは一度も承認が返されませんでした (「結果」セクションのレビュー概要の図を参照)。これらのレビューにより、人間のレビュー担当者が後に提起した問題の多くが明らかになりました (「結果」セクションのレビュー比較表を参照)。しかし、エージェントは AI レビューからのフィードバックに創造的に対処しませんでした。否定的なフィードバックに直面したとき、彼らは既存の発見に警告を追加することで対応し、見込みのない研究の方向性を追求し続けました。
エージェントは、見込みのないアプローチから効果的に後退しませんでした。エージェントは当初、複数の異なる研究方向を実験し、プロセス全体を通じて局所的に後戻りしましたが、どちらも最初の 10 時間以内に最も野心的な研究目標を撤回し、それ以降はどちらのエージェントもアプローチを根本的に変更することはありませんでした。
エージェントは具体的な指示に従わなかった。彼らは指示のずれに悩まされ、探索に費やす時間、AI レビュー ツールからレビューを取得する頻度、論文の長さの厳格な制限などに関する明確なルールを無視しました。その結果、両方のフィナ

l 論文は、AI カンファレンスに提出するための技術的要件を満たしていませんでした。
エージェントが提出した作品に対する原論文著者のレビューの概要
タスクと研究のセットアップに関するセクションには、セットアップに関する追加の詳細が含まれており、付録 1 では、両方の論文の研究上の疑問と関連する背景の概要が説明されています。
OpenClaw を使用してこれらの実験を実行したため、スキャフォールドはモデル プロバイダーに依存しませんでした。私たちは OpenAI と Anthropic のモデルを使用して予行実験を行った後、最もパフォーマンスの高いモデルとして Opus 4.8 に落ち着きました。私たちの結果は主にスキャフォールドの制限によって説明されるのではないかという懸念に応え、GPT-5.6 Sol とそのネイティブ スキャフォールドである Codex を使用し、同じ時間と API 予算で 1 つの論文で実験を繰り返しました。この実験の結果は、OpenClaw/Opus 4.8 の実験と同様でした。これにより、私たちの結果は単なる足場欠損によるアーチファクトではないと確信することができます。この実行では、特定された障害モードのほぼすべてが再現されました。
シャドウ評価には、エージェントによって生成された論文を元の著者がレビューすることが含まれることに注意してください。これは潜在的なバイアスにつながる可能性があります (彼らは論文が AI によって生成されたものであることを知っています。エージェントがとったアプローチではなく、質問に答えるためにとったアプローチを好む可能性があります)。このアプローチの制限については、以下の長所と制限の表および「制限」セクションで説明します。しかし、設計の非盲検性は理想的ではありませんが、私たちの実験でエージェントによって生成された論文は明らかに品質が低かったです。それでも、私たちはこの論文とともに成果物を公開し、それぞれのトピック分野の他の専門家に評価してもらうことを歓迎します。 3 3. シャドウ評価に含まれている論文の 1 つは、まだ評価されていません。

公開されているため、エージェントの詳細なログやエージェントが作成した文書は公開されません。もう 1 つの論文は、実験が終了した後に公開されました。
私たちは、GPT-5.6 Sol、Opus 5、および Fable 5 を使用して、より多くの研究論文の追跡実験を実行し、足場とモデルの改善に対する結果の感度をよりよく理解するために足場をさらに最適化する予定です。しかし、私たちの結果は、今日のフロンティア モデルでは数週間にわたる無制限の AI 研究課題を解決できないという初期の証拠を提供していると考えています。
AIの研究開発を研究する実験の厳選された評価とデモンストレーション
これまでの研究のほとんどは、次の 2 つのカテゴリのいずれかに分類されます。a) プログラムまたは LLM を介して指定された狭い指標に照らして評価される研究タスク、または b) 人間のレビューによって評価されるオープンエンドの研究。
シャドウ評価: AI 研究の自動化に向けた進捗状況を測定する新しい方法
自動化された AI 研究を評価するには、エージェントが作成する研究の品質を測定する方法が必要です。既存の評価のほとんどは、検証可能なタスクに対する評価またはブラインドレビューという 2 つのアプローチのいずれかを使用します。検証可能なタスクの評価では、エージェントが固定指標を改善し、自動検証器が結果をスコア付けします。元もたくさんいるよ

[切り捨てられた]

## Original Extract

We gave frontier agents research questions from two non-public NeurIPS submissions and had the original authors grade the results. The papers produced by the agents were unambiguously rejected.

CRUX Collaborative Research for Updating AI eXpectations Home
Can AI agents conduct open-ended AI research? Early evidence from two case studies
We gave frontier agents research questions from two non-public NeurIPS submissions and had the original authors grade the results. The papers produced by the agents were unambiguously rejected.
Forecasts of explosive AI progress hinge on AI agents automating AI research. But evidence on whether agents can carry out open-ended AI research is thin. Current evaluations either test agents on narrow, verifiable tasks, which excludes open-ended research, or submit AI-generated papers to blind peer review, which is overstretched, stochastic, and suffers from poor review quality. We introduce a third way to measure progress towards AI R&D automation. An agent takes on the central, open-ended research question of a high-quality unpublished paper, and the paper's original authors grade its output. We call these shadow evaluations. We ran shadow evaluations on two unpublished NeurIPS 2026 submissions, giving frontier agents six days and thousands of dollars of compute. The agents completed all of the engineering without human help, yet could not make substantial progress towards answering the research questions. As a result, both papers were unambiguously rejected by the authors. We identify five recurring failure modes: poor judgment about the bar for publishable research, uncreative responses to shortcomings in the research design, ineffective backtracking from dead ends, poor resource awareness, and instruction drift. A robustness check with a second model and scaffold reproduced these failures. We release the expert reviews, survey responses, agent repositories, and logs. Our results provide early evidence that today's agents can do the engineering of AI research, but struggle with critical parts of the research lifecycle.
One of the most consequential open questions about AI capabilities is whether AI agents can conduct AI research. Many forecasts of explosive AI progress speculate that AI systems will soon do AI research themselves . This is also the explicit premise of leading AI labs; in June, Anthropic published a post entitled “ When AI Builds Itself ,” and in July, OpenAI advertised that their new model, GPT-5.6 Sol, had helped post-train a smaller model , saving researchers multiple weeks. 1 1. Notably, GPT-5.6 Sol’s contribution to Luna’s post-training is not mentioned in the 81-page system card . But despite the significance of this research direction and the attention paid to these claims, the evidence base on whether agents can solve open-ended research questions is thin.
What unifies most of this recent work on autonomous AI research is its focus on verifiable tasks where agents need to improve a fixed, narrow metric. Benchmarks ask agents to improve a known metric, and an automatic verifier scores the result. Beyond benchmarks, a number of evaluations have shown AI agents beating expert human performance in tasks such as optimizing GPT-2 level models , using weaker models to train stronger ones , and optimizing an autoresearch harness . 2 2. Appendix 4 describes these and many other AI research experiments that comprise verifiable tasks.
But much AI research goes beyond solving verifiable tasks. Agents can’t hill-climb their way into choosing a set of candidate hypotheses, deciding what evidence would settle a research question, or incorporating feedback effectively and recognizing that an approach has failed and the right move is to start over.
A small number of projects have adopted a different method for evaluating AI research: they submit AI-generated papers to blind peer-review processes, such as to AI conferences and workshops. But peer review is a weak measure of research ability: conference reviewing is overstretched and highly stochastic , and it does not reveal how many AI-generated submissions were rejected before the eventual acceptance.
CRUX is our project to conduct open-ended, long-horizon evaluations of frontier AI systems on challenging real-world tasks. It pushes frontier AI systems farther than benchmarks can by focusing on a small set of realistic challenges, on which we analyze AI systems’ performance deeply. Three months ago, we wrote a paper laying out the foundations of long-horizon open-world evaluations. In this evaluation, we ask: can agents solve open-ended AI research questions?
In this paper, we evaluate whether AI agents can conduct open-ended AI research with a new method, which we call shadow evaluation . This involves taking the central research question from a high-quality research paper that is not yet public, tasking a well-resourced frontier agent with answering it, and asking the paper’s original authors to grade the agent’s output as they would a conference submission. The agent “shadows” the original study: it works on the same research question as the original authors, without access to their paper or findings. This design gives us open-ended tasks, uncontaminated questions, and reviewers with deep expertise in the exact question being evaluated. We consider shadow evaluations complementary to both narrow, verifiable evaluations and blind review evaluations of automated AI research. We hope future research uses all three to explore different facets of measuring progress towards automating AI research.
To carry out shadow evaluations, we partnered with the authors of two papers submitted to NeurIPS 2026 that were not yet public. Unpublished papers give us a way to test if agents can solve open-ended research problems. Through their own months-long effort in thinking about the question, the original authors are uniquely positioned to grade the quality of the agent’s work, and the agent cannot look up the researchers’ findings, because they are not yet on the web. We gave an agent the paper’s research question, six days of wall-clock time, $3,000 in Anthropic API credits to allow agents to conduct open-ended exploration and run experiments over the course of a week, GPU credits for experiments, and full access to a VM and the open web. The goal was to produce a paper worthy of publication at a top-tier AI conference. The original authors of both papers then graded the results as conference reviewers.
Our key finding was that while agents could solve the engineering problems necessary to do the research, they failed to produce original research at the caliber of a top ML conference (see the table below for details). Our main takeaways:
The agents lacked the judgment to identify when a problem was adequately solved. They understood the research questions and proposed directions that closely mirrored those of the authors of the original papers. But they then falsified those hypotheses using small, hand-curated or synthetic datasets. In their final papers, they engaged only shallowly with the literature, and they presented underpowered negative results as substantive findings.
The agents lacked awareness about the resources available to them and the timeline for the project. Both runs ended with less than 50% of the API budget spent, even though the agents could monitor their own usage in real time, and were encouraged to use their remaining budgets. The agents did not appear to intuitively grasp the meaning of these resource limits, particularly the time limit. They rushed through their initial exploration in a number of hours, and finished with hours of clock time remaining despite papers that did not meet their own bar for success.
The agents did not creatively respond to feedback about poor research design. We instructed the agents to send their papers for AI review — both to a subagent and to external tools such as refine.ink — to assess the quality and progress of their research. Across dozens of rounds of revision, the agent’s self-review never once returned an acceptance (see the review summary figure in the Results section). These reviews surfaced many of the issues that the human reviewers later raised (see the review comparison table in the Results section). But the agents did not creatively address the feedback from the AI reviews; when faced with negative feedback they responded by adding caveats to existing findings, and continued to pursue unpromising research directions.
The agents did not effectively backtrack from unpromising approaches. While the agents initially experimented with multiple distinct research directions and backtracked locally throughout the process, they both retired their most ambitious research targets within the first ten hours, and neither agent fundamentally shifted its approach after that point.
The agents did not follow concrete instructions. They suffered from instruction drift and ignored explicit rules about how much time should be spent on exploration, how often to get reviews from AI review tools, and strict limits on paper length. As a result, both final papers failed the technical requirements of submitting these papers to AI conferences.
Summary of the original paper authors’ reviews of the agents’ submitted work
The section on the task and the research setup contains additional details on our setup, and Appendix 1 outlines the research questions and relevant context for both papers.
We used OpenClaw to run these experiments so that our scaffold was agnostic to the model provider. We conducted dry-run experiments with models from OpenAI and Anthropic before settling on Opus 4.8 as the best-performing model. In response to concerns that our results might be principally explained by a limitation in our scaffold, we repeated our experiment on one paper using GPT-5.6 Sol and Codex, its native scaffold, with the same time and API budgets. The results of this experiment were similar to our OpenClaw/Opus 4.8 experiments. This makes us more confident that our results are not simply artifacts of a scaffold deficiency; this run reproduced nearly every single one of our identified failure modes.
Note that shadow evaluations involve the original authors reviewing the paper generated by the agent. This might lead to potential biases (they know the paper is AI-generated; they might prefer the approach they took to answer the question rather than the one the agent took). We discuss the limitations of this approach in the strengths-and-limitations table below and in the Limitations section. But while the non-blindedness of the design is not ideal, the papers generated by the agents in our experiments were unambiguously of poor quality. Still, we release the artifacts alongside this paper and we welcome other experts in the respective topic areas to judge them. 3 3. One of the papers included in our shadow evaluation is not yet public, so we do not disclose the detailed agent logs or the agent-generated paper. The other paper was made public after we concluded our experiments.
We plan to run follow-up experiments on a larger set of research papers using GPT-5.6 Sol, Opus 5, and Fable 5, and further optimize scaffolds to better understand the sensitivity of our results to scaffold and model improvements. But we think our results provide early evidence that today’s frontier models cannot solve weeks-long, open-ended AI research questions.
Selected evaluations and demonstrations of experiments studying AI research and development
Most previous work falls into one of two categories: a) research tasks evaluated against a narrow metric, specified programmatically or via LLM-as-a-judge, or b) open-ended research evaluated by human review.
Shadow evaluations: a new method for measuring progress towards automating AI research
Evaluating automated AI research requires a way to measure the quality of the research that agents produce. Most existing evaluations use one of two approaches: evaluations on verifiable tasks or blind review. In evaluations on verifiable tasks, the agent improves a fixed metric, and an automatic verifier scores the result. There are many ex

[truncated]
