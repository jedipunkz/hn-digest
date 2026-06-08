---
source: "https://olshansky.info/posts/2026-06-08-reading-of-openais-self-improving-tax-agents"
hn_url: "https://news.ycombinator.com/item?id=48446431"
title: "Reading of OpenAI's Self-Improving Tax Agents"
article_title: "Reading of OpenAI's Self-Improving Tax Agents | 🦉 olshansky 🦁"
author: "Olshansky"
captured_at: "2026-06-08T16:29:14Z"
capture_tool: "hn-digest"
hn_id: 48446431
score: 2
comments: 0
posted_at: "2026-06-08T15:13:58Z"
tags:
  - hacker-news
  - translated
---

# Reading of OpenAI's Self-Improving Tax Agents

- HN: [48446431](https://news.ycombinator.com/item?id=48446431)
- Source: [olshansky.info](https://olshansky.info/posts/2026-06-08-reading-of-openais-self-improving-tax-agents)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T15:13:58Z

## Translation

タイトル: OpenAI の自己改善型税務代理人の読書
記事のタイトル: OpenAI の自己改善型税務代理人に関する読書 | 🦉オルシャンスキー🦁
説明: OpenAI のブログのタイトルを最初に見たとき、私はあまり興奮しませんでした。それは次のとおりです。Codex を使用して自己改善型の税務担当者を構築します。
税務職員というとあまり魅力的ではありません。すべては自己改善です。わかった。しかし、この短い投稿は、最新のツール、ワークフロー、
[切り捨てられた]

記事本文:
OpenAI の自己改善型税務代理人の読書
登るのに最適な丘を見つける
本番トレースから評価まで
前方展開したエンジニアがどこにも行かない理由
数字と抱えている問題の解決
OpenAI のブログのタイトルを最初に見たとき、私はそれほど興奮しませんでした。それは次のとおりです: Codex を使用して自己改善型の税務担当者を構築する 。
税務職員というとあまり魅力的ではありません。すべては自己改善です。わかった。しかし、この短い投稿は、最新のツール、ワークフロー、実際のユースケースで評価を行う方法についての実践的で実践的な学習のマスタークラスです。それは技術的なことだけではありません。実世界の経験に基づいて直感を構築します。
これは、私が最初に Waymo で働き始めたときに欲しかった投稿です。
登るのに最適な丘を見つける #
ML/AI 実践者なら誰でも、この強力なオープン性を高く評価するでしょう。
「現実世界のシステムは、本番環境では実験室とは異なる動作をし、展開前に予測するのが困難な方法で壊れます。」
まず始めに、問題は「AI を税金に使用できるか?」ということではありませんでした。 、それは次のとおりでした。
「彼らは、納税シーズンの最も忙しい時期に納税準備が大きなボトルネックになっていると指摘しました。」
本題に入りませんが、中心的な問題には、大量のドキュメントをアップロードし、PDF からデータを正確に解析して抽出し、データの完了率と手動レビューのフラグが立てられた返品を確認することで精度を測定することが含まれます。
現実世界の評価問題と同様に、フィードバック ループには 3 つの重要な柱があります。
フィードバック : 専門家ではない分野で、適切な問題を適切な順序で確実に解決するには、エンド顧客とできるだけ緊密に連携する必要があります。これにより、表示されるデータが決まります。
Eval : 本番トレース、ホールドアウト セット、ゴールデン テスト セット コーパス

エスなどなど。これには経験と試行錯誤が必要です。フィードバックによって、何を評価すべきかが決まります。
自己改善 : 実際に改善するにはモデルが必要です。これは、モデルのルーティング、プロンプト調整、ポストトレーニング (微調整など)、評価に基づくしきい値の設定などです。これが実際に結果を左右するものです。
私は問題文の最後にあるこの引用がとても気に入っています。
「登るべき丘を特定するための信号がありませんでした。」
そのギャップこそが、あらゆる分野に適用できる彼らの方法論を推進したのです。
これは、AI 時代の前方展開エンジニアのためのハンドブックです。
本番環境で証拠が作成されるように製品を構築する
Codex 主導の改善ループを作成する
プロダクショントレースから評価まで #
この投稿では、財務詳細が複数の文書 (スプレッドシート、PDF、手書きのメモなど) にまたがる可能性がある賃貸物件の管理に関する具体的な例を取り上げています。
精度と再現率を取得する際の興味深い点は、ネガティブ サンプルが必ずしも客観的であるとは限らず、実践者の好みである可能性があることです。
とてもシンプルですが、とてもリアルなフレームです。だからこそ、組織ごとに微調整する必要があるのです。もしかしたら、いつかチームごとに？ 1人あたり？ 「パーソナルコンピュータ」というアイデアも人々に笑われました。私の考えについては、別の投稿でさらに共有します。
製品トレースを eval に変換することに関する OpenAI のメモは 3 つに分かれています。
(人間によって) 提出された申告書と Tax AI の出力の違いを把握します。これは素晴らしいことです。最終的な戻り値は、グラウンド トゥルース コーパスとして取得できる最良のサンプルです。
関連するエラーをグループ化して、繰り返し発生するエラーをワークフローのノイズから分離します。これはデータを理解するのに役立ちますが、特定の分野の障害に過度に集中したり、過度にトレーニングしたりすることも避けられます。
繰り返されるパターンを eval t に変換する

レビューと測定を経て目標を達成します。丘を見つけ、道具を揃え、基本的なスキルを身につければ、登山はとても楽しいものになります。
上のメモを書き終えて次の文に移ったとき、「修正が丘の候補になる」という文字が目に入りました。 :) OpenAI は非常にクールなグラフィックを提供してくれたので、ここにリンクしておきます。
次のステップは、エンジニアリング作業と時間のほとんどがかかる部分です。
パイプラインを調査します: パッケージ、スキーマ、コード パス、バグ、トレース、ログ、デプロイメント環境、オフライン/オンライン スキューなど。
修正を実装します: ソース選択、解析、税金エンジンなどを更新します。
検証と提案 : ターゲットを絞った評価を再実行し、測定し、回帰スイートを実行し、A/B テストによるデプロイを検討し、トラフィックをシャドウし、再評価します。
ループを閉じる : 新しいメトリクス、テスト セット、ダッシュボード、デプロイされた変更の可視性を得るために、できるだけ多く (すべてではありません) を自動化します。
次の図は、インラインに落とし込むほど気に入ったものです。
前方展開したエンジニアがどこにも行かない理由 #
先ほども述べたように、これは Codex に特化したブログではありません。これは、評価フィードバックが現実の世界でどのようなものであるかを示すためのブログです。チームは次のように述べています。
「賃貸物件の例は、エージェントの能力を向上させるために本番環境の成果物とトレースを使用するという、より広範な再利用可能なパターンを象徴しています。」
そして、なぜ前方展開エンジニアがすぐにどこにも行かないのかと尋ねる人がいたら、これが答えです。
その証拠は自動的にコーデックスのタスクになるわけではありません。担当者の修正には、抽出ミス、マッピングの問題、サポートされていない製品の動作、税務上の判断、または予想されるワークフローのノイズが反映される場合があります。 […] エンジニアは引き続きアーキテクチャ、製品の決定、出荷に責任を負います。
Codex が動作する候補ワークスペースは、おおよそ次のようになります。
/ca

ndidates/FIND-RENTAL-0042/
§── repo/ # ブランチ: codex/fix-rental-0042
│ §── AGENTS.md
│ §── タスク/FIND-RENTAL-0042/
│ §── app/tax-ai/家賃収入/
│ §── evals/
│ §── スキル/
│ └── ドキュメント/
└──scoped-tools/ # 読み取り専用の実稼働コンテキスト
§── 生産トレース
§── ソースアーティファクト
└── 税金エンジンのドキュメント
数字と抱えている問題の解決 #
最後になりますが、いくつかの数字は常に素晴らしいものです。
プロジェクトには 6 週間かかり、精度と再現率は 90% に達しました。
ある上級会計士は、昨年、税務準備に 180 時間から 15 時間まで費やしました。
また、BitTorrent の作成者であり、私がこれまで数回お会いする機会に恵まれた人物である Bram Cohen から最近聞いた引用を 1 つ追加します。彼はとても率直で、現実的で、楽しくて実践的です。現実世界での eval に関して言えば、次のようになります。
「秘訣は、あなたが望む問題ではなく、あなたが抱えている問題を解決することです。」

## Original Extract

When I first saw the title of OpenAI’s blog, I wasn’t too excited. Here it is: Building self-improving tax agents with Codex.
Tax agents don’t sound exciting. Everything is self-improving. I got it. But this short post is a masterclass in practical, hands-on learnings on the latest tools, workflows,
[truncated]

Reading of OpenAI's Self-Improving Tax Agents
Finding the Right Hill to Climb
From Production Traces to Evals
Why Forward Deployed Engineers Aren’t Going Anywhere
The Numbers, and Fixing the Problem You Have
When I first saw the title of OpenAI’s blog, I wasn’t too excited. Here it is: Building self-improving tax agents with Codex .
Tax agents don’t sound exciting. Everything is self-improving. I got it. But this short post is a masterclass in practical, hands-on learnings on the latest tools, workflows, and approach to how to do evals for real-world use cases. It’s not just technical; it builds intuition by relying on real-world experience.
This is the post I wish I had when I first started my job at Waymo.
Finding the Right Hill to Climb #
Any ML/AI practitioner would appreciate this strong open:
“Real-world systems behave differently in production than they do in a lab, breaking in ways that are hard to anticipate before deployment.”"
To get things started: the problem wasn’t “Can we use AI for taxes?” , it was:
“They pointed us to tax preparation as a significant bottleneck during the busiest stretch of tax season.”
Without getting into the weeds, the core problem involves uploading a bunch of documents -> accurately parsing and extracting data from PDFs, and measuring accuracy by looking at data completion % along with returns flagged for manual review.
Like any real-world eval problem, there are three key pillars to the feedback loop:
Feedback : To make sure you’re solving the right problem(s), in the right order, in a domain you might not be an expert in, you need to work as closely as possible with the end customer. This drives the data you’ll be looking at.
Eval : Production traces, hold-out sets, golden test set corpuses, and much more. This requires trial-and-error along with experience. Feedback will drive what you should evaluate.
Self-improvement : You need the model to actually improve. This might be model routing, prompt tuning, post-training (e.g. fine-tuning), setting thresholds based on the evals, etc. This is what actually drives the results.
I really like this quote at the end of the problem statement:
“We did not have the signal to identify the right hill to climb.”
That gap is what drove their methodology, one applicable to any field.
This is the playbook for Forward Deployed Engineers in the era of AI:
Build the product so production creates evidence
Create a Codex-driven improvement loop
From Production Traces to Evals #
The post goes into a concrete example on managing rental properties where the financial details may be spread across multiple documents: spreadsheets, PDFs, hand-written notes, etc…
The interesting part about capturing your precision and recall is that negative samples are not always objective, they may be practitioner preferences.
This is a very simple but very real framing. It’s why we’ll need fine-tuning per organization. Maybe, one day, per team? Per person? People laughed at the idea of “personal computers” too. I’ll share more of my thoughts here in another post.
OpenAI’s note on turning product traces into evals is split into three:
Capture the difference between the filed return (by the human) and Tax AI’s output. This is great, the final return is the best sample we can get for our ground truth corpus.
Group related failures to separate recurring errors from workflow noise. This helps with understanding the data, but also avoids overtraining/overfocusing on a specific vertical of failures.
Turn repeated patterns into eval targets after review and measurement. Once you’ve found the hill, and you have the tools, and you have the baseline skills, climbing is quite fun.
As I finished writing the note above, and moved on to the next sentence, I saw: “Corrections become hill candidates.” :) OpenAI provided a really cool graphic that I’ll simply link to here .
The next step is where most of the engineering work and time lies:
Investigate the pipeline : packages, schemas, code paths, bugs, traces, logs, deployment environments, offline/online skew, etc.
Implement the fixes : Update source selection, parsing, tax-engine, etc.
Validate & propose : Re-run targeted evals, measure, run regression suites, consider deploying with an A/B test, shadow traffic, re-evaluate.
Close the loop : Automate as much (not everything) as possible to have new metrics, test sets, dashboards, and visibility into the deployed changes.
This next diagram I liked enough to drop inline:
Why Forward Deployed Engineers Aren’t Going Anywhere #
Like I said earlier on, this isn’t a Codex-specific blog. It is a blog to provide a view into what eval feedback looks like in the real world. The team says:
“The rental property example is emblematic of a broader reusable pattern: using production artifacts and traces to improve an agent’s capabilities.”
And if anyone ever asks why Forward Deployed Engineers aren’t going anywhere anytime soon, this is the answer:
That evidence does not become a Codex task automatically. A practitioner correction may reflect an extraction miss, a mapping issue, unsupported product behavior, tax judgment, or expected workflow noise. […] Engineers remain responsible for architecture, product decisions, and shipping.
The candidate workspace Codex operates in looks roughly like this:
/candidates/FIND-RENTAL-0042/
├── repo/ # branch: codex/fix-rental-0042
│ ├── AGENTS.md
│ ├── tasks/FIND-RENTAL-0042/
│ ├── app/tax-ai/rental-income/
│ ├── evals/
│ ├── skills/
│ └── docs/
└── scoped-tools/ # read-only production context
├── production-trace
├── source-artifacts
└── tax-engine-docs
The Numbers, and Fixing the Problem You Have #
To close it out, some numbers are always great.
The project took 6 weeks and reached 90% precision and recall .
One senior accountant went from 180 hours to 15 hours on tax preparation last year.
I’ll also add one quote I came by recently from Bram Cohen, the creator of BitTorrent and someone I’ve had the pleasure of meeting a handful of times now. He’s very direct, pragmatic, fun, and practical. When it comes to evals in the real world, this hits:
“The trick is to fix the problem you have, rather than the problem you want.”
