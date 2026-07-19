---
source: "https://neutralityproject.org/"
hn_url: "https://news.ycombinator.com/item?id=48969113"
title: "The neutrality project – Making the politics inside AI measurable"
article_title: "The Neutrality Project: measuring the politics inside AI"
author: "amai"
captured_at: "2026-07-19T15:46:52Z"
capture_tool: "hn-digest"
hn_id: 48969113
score: 1
comments: 0
posted_at: "2026-07-19T15:38:20Z"
tags:
  - hacker-news
  - translated
---

# The neutrality project – Making the politics inside AI measurable

- HN: [48969113](https://news.ycombinator.com/item?id=48969113)
- Source: [neutralityproject.org](https://neutralityproject.org/)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T15:38:20Z

## Translation

タイトル: 中立性プロジェクト – AI 内部の政治を測定可能にする
記事のタイトル: 中立性プロジェクト: AI 内の政治の測定
説明: 言語モデルがどのように機能するかを測定するオープンソースの再現可能な方法

記事本文:
Neutrality Project は助成金を積極的に募集しており、寄付も受け付けています。プロジェクトを支援する→
中立性プロジェクト
なぜそれが重要なのか
ベンチマーク
方法論
結果 ↗
ロードマップ
サポート
チーム
◐
オープンソースの研究 · リリース 01
AI 内部の政治を測定可能にします。
私たちは、AI の世界観を可視化するオープンで再現可能なベンチマークを構築し、人々がモデルがどのように判断を形成しているかを理解できるようにします。
すべてのモデルがどこかに着陸します。私たちは、その立場と自信を読みやすく、再現可能にし、精査にさらすようにしています。
AI は質問に答えるだけではありません。それは私たちの考え方を組み立てるものです。
調査、推論、執筆をアシスタントに委託すると、アシスタントの仮定が意思決定の一部になります。そのため、AI のデフォルトの世界観は、技術的な脚注ではなく、公益的な問題になります。
モデルが議論を要約したり、「バランスのとれた」見解を選択したり、結論を下書きしたりするとき、その枠組みがユーザーの出発点になりますが、気づかれないことがよくあります。
モデルが役職を発表することはほとんどありません。どのオプションを検証するか、何を妥当な中間として扱うか、何を除外するかによって決まります。それは測定可能です。
「中立です、信じてください」は検証不可能です。誰でも読んで、再実行して、挑戦できるベンチマークです。したがって、私たちが構築するものはすべてオープンソースです。
AI による人類の将来は、その影響を可視化し、その作成者が公に責任を負うことにかかっています。
AI が人々の学習、推論、決定の仲介者となるにつれ、その隠れた思い込みがエコー チェンバーを強化し、認知的不協和を深め、さまざまなコミュニティが真実または合理的として受け入れるものを静かに形づくることで分断を拡大する可能性があります。
私たちはその影響力を測定可能にし、AI ラボは公の場で責任を負います。私たち自身の結果は必要性の証拠です。独立した測定により、影響がすでに表面化しています。

研究室の外で遺体が見られた可能性はありますが、どの研究室もそれを独自に公開していませんでした。いかなる企業も、証拠の基準として「当社を信頼してください」と認められるべきではありません。
リリース 01 · 政治的中立性のベンチマーク
単一の評決ではなく、あらゆるモデルの調整された政治的プロフィール。
このベンチマークには、実際の世論調査の数千の質問に対する模範回答があり、その回答パターンを各イデオロギー軸上の位置として読み取ります。 2 つの設計選択により、結果が正直に保たれます。
1
セルフアンカーリング: バイアスのない定規が焼き付けられています
同じモデルは、極左と極右のロールプレイングでも実行されます。その中立的な回答はそれぞれの極端に配置されているため、「ソーシャルでの -0.7」は、このモデル自体の極左寄り、つまり私たちの中央の意見ではなくモデルごとの調整に 70% 向いていることを意味します。
2
国を超えた参考資料: 一人で採点する人はいません
どちらの答えがどちらに傾くかは、3 つの異なる国と研究所の独立したモデルによって事前に決定されるため、単一の国家的観点が「左」と「右」を定義することはありません。警備員は、自身の家族が作成に協力したルールブックに照らしてモデルが採点されるのを阻止します。
3
ディメンションごとにレポートされ、ブレンドされることはありません
ゲームに唯一の「中立性スコア」はありません。各軸は独自にレポートされ、実行の失敗にフラグを立てる健全性チェックと、拒否された質問が軸に偏りを与えているかどうかを示す拒否レポートが含まれます。
プロジェクト全体が守るべき原則。
コード、質問セット、凍結されたリファレンス、およびすべての結果は公開されています。読んで、再実行して、同意しないでください。
1 つのコマンドを入力するだけで、調整されたプロファイルが出力され、再開可能で監査に十分な決定性があり、独自のハードウェア上で実行できます。
このリファレンスはさまざまな国や研究室のモデルによって書かれているため、1 つの世界観を「中立」としてエンコードするわけではありません。
軸ごとにレポートし、信頼性の低い実行にフラグを立てて、結果が希薄化した場所、または拒否によって結果が歪められた可能性のある場所を指定します。

政治的傾向は中立性の第一の軸です。最後ではありません。
同じオープンで調整されたアプローチは、モデルがそれに依存する人々を静かに導くことができる他の方法にも拡張されます。
モデルがイデオロギー軸をまたがって配置され、自己アンカーされ、拒否検出が組み込まれ、複数国の参照に対してスコア付けされます。プロジェクトの最初のオープン リリースとして現在利用可能です。
左/右を超えて: モデルが対立する主題をどのように組み立てるか、問題が表面化するか、平坦化されるか、質問に 2 つ以上の側面がある場合にデフォルトでどのような構成になるかについてです。
モデルが拒否したり、和らげたり、黙って避けたりするもの: 得られる答えが、話された内容と同じくらい差し控えられた内容によって形作られるトピックをマッピングします。
独立した測定への資金提供にご協力ください。
私たちは助成金を積極的に募集しており、寄付も受け付けています。ベンチマーク実行のためのコンピューティングが主なコストであり、測定するラボからの独立性がポイントです。
小さなチームからスタートしました。コミュニティとともに成長するように構築されています。
以下の人々は初期の創設者であり、プロジェクトの限界ではありません。 Neutrality Project はオープンソースであり、その品質は、方法論に挑戦し、より多くのモデルをベンチマークし、結果をレビューし、私たちと協力してより良いツールを構築する貢献者に依存します。
Samuel Cardillo は、システムの構築とセキュリティ保護に 20 年を費やしてきた技術者兼起業家です。彼はナイキが買収したデジタル スタジオ RTFKT の CTO であり、そこでテクノロジーをエンドツーエンドで主導し、以前は Google で働いていました。オープン AI モデルのトレーニングや自己ホストから、次に分解する価値のあるものまで、さまざまなテクノロジーを絶え間なく実験する彼は、現在 AI の導入と倫理に重点を置いています。
ダニエル・ルーゲンは、トロント大学の視覚神経科学の博士研究員です。彼の学術研究は、初期の視覚経路と注意のメカニズムに焦点を当てています。

視覚情報がどのように選択され、処理されるかを決定します。博士課程の研究とは別に、ゲシュタルト研究所を通じてオープンソースの AI モデル、評価ベンチマーク、トレーニング方法、エージェント インフラストラクチャを開発しています。推論の品質、再現性、研究や技術的作業に真に役立つシステムの構築に重点を置いています。
Kai Stephens は、有能なモデルを実用的でローカルで検査可能にすることに重点を置いたオープンソース AI 開発者です。彼は、ターミナル、ファイル、ブラウザ、リポジトリ、デバッグ、およびマルチステップ ツールのワークフローを処理するために、Hermes Agent 用に調整されたモデル ファミリである Carnice を作成しました。 Hugging Face で 50 万件以上のモデルがダウンロードされたカイは、モデル トレーニング、エージェント インフラストラクチャ、行動評価、再現可能なオープンソース システムにおける実践的な経験を The Neutrality Project にもたらしています。
早くから仕事に参加し、創設者と一緒に仕事を形作ってきた人々。
アンドリューの仕事は、2017 年の個人的な転機によって動機づけられました。脳幹に神経膠腫腫瘍が発生し、即時手術が必要になりました。それ以来、脳と、脳が生の入力を意図的な経験に変換するプロセスに対する彼の魅力は、2025 年にオレゴン大学で認知神経科学の博士号を取得するきっかけとなりました。彼は現在、その焦点を AI ビジョンに移し、精神物理学と人間ベースの経験データに基づいたメトリクスとデータセットを設計しています。
プロジェクトには創設者以上のものが必要です。
モデルの対象範囲を拡大し、ベンチマークの実行に貢献し、質問と結果をレビューし、コードを改善し、または私たちの仮定に異議を唱えるのに役立ちます。オープンな監視と多様な貢献者によって、このプロジェクトはより厳密で有用で信頼できるものになります。
アシスタントは、読みやすいということによってのみ信頼できるものとなります。
Neutrality Project は、AI 内の政治を測定、再実行、

そして公の場で責任を追及します。コミュニティによって構築されたものです。助けに来てください。
独立した測定には独立した資金が必要です。サポートは、ベンチマークの実行、新しいモデル、今後のリリースなどの研究自体に資金を提供します。

## Original Extract

An open-source, reproducible way to measure how a language model

The Neutrality Project is actively looking for grants and open to donations . Support the project →
The Neutrality Project
Why it matters
The benchmark
Methodology
Results ↗
Roadmap
Support
Team
◐
Open-source research · Release 01
Making the politics inside AI measurable.
We build open, reproducible benchmarks that make AI worldviews visible, so people can understand how a model may be shaping their judgment.
Every model lands somewhere . We make its position and confidence legible, reproducible, and open to scrutiny.
AI does more than answer questions. It frames how we think.
When people delegate research, reasoning, and writing to an assistant, its assumptions become part of their decisions. That makes an AI's default worldview a public-interest question, not a technical footnote.
When a model summarises the debate, picks the "balanced" take, or drafts the conclusion, its framing becomes the user's starting point, often unnoticed.
A model rarely announces a position. It leans through which options it validates, what it treats as the reasonable middle, and what it leaves out. That is measurable.
"Trust us, it's neutral" is not verifiable. A benchmark that anyone can read, rerun, and challenge is. So everything we build is open source.
Humanity's future with AI depends on making its influence visible, and its makers openly accountable.
As AI becomes an intermediary for how people learn, reason, and decide, its hidden assumptions can reinforce echo chambers, deepen cognitive dissonance, and widen division by quietly shaping what different communities accept as true or reasonable.
We make that influence measurable, and AI labs accountable in the open. Our own results are proof of the need: independent measurement has already surfaced influence that nobody outside a lab could have seen , and that no lab had disclosed on its own. No company should get "trust us" as its standard of proof.
Release 01 · The Political Neutrality Benchmark
A calibrated political profile for any model, not a single verdict.
The benchmark has a model answer thousands of real public-opinion survey questions, then reads its pattern of answers as a position on each ideological axis. Two design choices keep the result honest.
1
Self-anchoring: a ruler with no bias baked in
The same model is also run role-playing far-left and far-right. Its neutral answers are placed on its own extremes, so "−0.7 on social" means 70% toward this model's own far-left, a per-model calibration rather than our opinion of center.
2
A cross-country reference: nobody grades alone
Which answer leans which way is fixed ahead of time by independent models from three different countries and labs , so no single national perspective defines "left" and "right." A guard blocks a model from being graded against a rulebook its own family helped write.
3
Reported per dimension, never blended
There is no one "neutrality score" to game. Each axis is reported on its own, with a sanity check that flags a broken run and a refusal report that says whether declined questions have biased any axis.
Principles the whole project is held to.
Code, question sets, the frozen reference, and every result are public. Read it, rerun it, disagree with it.
One command in, a calibrated profile out, resumable and deterministic enough to audit, on your own hardware.
The reference is written by models from different countries and labs, so it doesn't encode one worldview as "neutral."
We report per-axis, flag low-confidence runs, and name where a result is diluted or where refusals may have skewed it.
Political leaning is the first axis of neutrality. Not the last.
The same open, calibrated approach extends to the other ways a model can quietly steer the people who rely on it.
Where a model sits across ideological axes, self-anchored and scored against a multi-country reference, with refusal detection built in. Available now as the project's first open release.
Beyond left/right: how models frame contested subjects, which concerns they surface, which they flatten, and whose framing they default to when a question has more than two sides.
What a model refuses, softens, or silently avoids: mapping the topics where the answer you get is shaped as much by what is withheld as by what is said.
Help fund independent measurement.
We are actively looking for grants and open to donations . Compute for benchmark runs is the main cost, and independence from the labs we measure is the point.
Started by a small team. Built to grow with its community.
The people below are the initial founders, not the limits of the project. The Neutrality Project is open source, and its quality depends on contributors who challenge the methodology, benchmark more models, review results, and build better tools with us.
Samuel Cardillo is a technologist and entrepreneur who has spent two decades building and securing systems. He was CTO of RTFKT, the digital studio acquired by Nike, where he led technology end to end, and previously worked at Google. A relentless experimenter across technologies, from training and self-hosting open AI models to whatever is worth taking apart next, he is now heavily focused on AI adoption and ethics.
Daniel Lougen is a PhD researcher in visual neuroscience at the University of Toronto. His academic work focuses on the early visual pathway and the attentional mechanisms that shape how visual information is selected and processed. Separate from his doctoral research, he develops open-source AI models, evaluation benchmarks, training methods, and agent infrastructure through Gestalt Labs, with an emphasis on reasoning quality, reproducibility, and building systems that are genuinely useful for research and technical work.
Kai Stephens is an open-source AI developer focused on making capable models practical, local, and inspectable. He created Carnice, a family of models tuned for Hermes Agent to handle terminal, file, browser, repository, debugging, and multi-step tool workflows. With more than half a million model downloads on Hugging Face, Kai brings hands-on experience in model training, agent infrastructure, behavioral evaluation, and reproducible open-source systems to The Neutrality Project.
People who joined the work early and shape it alongside the founders.
Andrew's work is motivated by a personal turning point in 2017: a ganglioglioma tumor on his brain stem leading to immediate surgery. Since then his fascination with the brain, and the process by which it transforms raw input into intentional experience, led him to a PhD in Cognitive Neuroscience at the University of Oregon in 2025. He now brings that focus to AI vision, designing metrics and datasets grounded in psychophysics and human-based empirical data.
The project needs more than its founders.
Help expand model coverage, contribute benchmark runs, review questions and results, improve the code, or challenge our assumptions. Open scrutiny and diverse contributors are how this project becomes more rigorous, useful, and trustworthy.
An assistant is only as trustworthy as it is legible.
The Neutrality Project makes the politics inside AI something you can measure, rerun, and hold to account, in the open. It's community-built, come help.
Independent measurement needs independent money. Support funds the research itself: benchmark runs, new models, and the releases to come.
