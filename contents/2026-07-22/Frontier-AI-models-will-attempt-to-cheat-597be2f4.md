---
source: "https://www.aisi.gov.uk/blog/cheating-behaviour-in-frontier-model-evaluations"
hn_url: "https://news.ycombinator.com/item?id=49010266"
title: "Frontier AI models will attempt to cheat"
article_title: "Cheating behaviour in frontier model evaluations | AISI Work"
author: "ambigious7777"
captured_at: "2026-07-22T18:04:16Z"
capture_tool: "hn-digest"
hn_id: 49010266
score: 3
comments: 0
posted_at: "2026-07-22T17:26:03Z"
tags:
  - hacker-news
  - translated
---

# Frontier AI models will attempt to cheat

- HN: [49010266](https://news.ycombinator.com/item?id=49010266)
- Source: [www.aisi.gov.uk](https://www.aisi.gov.uk/blog/cheating-behaviour-in-frontier-model-evaluations)
- Score: 3
- Comments: 0
- Posted: 2026-07-22T17:26:03Z

## Translation

タイトル: フロンティア AI モデルは不正行為を試みます
記事タイトル: フロンティアモデル評価における不正行為 | AISIの仕事
説明: すべての機能評価で不正行為を発見し、モデルの機能が向上するにつれてその影響を概説します。

記事本文:
フロンティアモデル評価における不正行為 | AISIの仕事
Frontier AI Trends Report を読む このウェブサイトでは JavaScript を有効にしてください。
あ
あ
研究助成金について ブログ お問い合わせ 採用情報
ホーム 研究助成金について ブログ 採用情報
ブログ
レッドチームのフロンティアモデル評価における不正行為
すべての機能評価で不正行為を発見し、モデルの機能が向上するにつれてその影響を概説します。
AI モデルが意図したとおりに動作することを信頼できますか?これは、AI システムを導入している人にとっても、その機能を評価しようとしている人にとっても、中心的な質問です。導入時、意図しない手段や不正な手段で目標を追求するモデルは、特に一か八かのユースケースで損害を引き起こす可能性があります。 AI の能力評価では、同じ動作が結果の妥当性を損なう可能性があります。モデルは、代わりにタスクやその環境を利用しているにもかかわらず、困難なタスクを完了することで能力を実証しているように見える可能性があります。
AISI は、不正行為と呼ばれるこの行為に対する AI モデルの監視と評価を開始しました。この動作についてテストしたすべてのモデルが不正行為を試みました。モデルは、質問されたときにこの行動を確実に報告せず、思考連鎖の中でそれについて推論しないことが多く、不正行為を検出するには堅牢な監視方法が必要である可能性が高いことを示唆しています。
このブログでは、サイバー評価における不正行為を分析し、モデルの機能が向上するにつれてこの行為がもたらす影響について議論します。
私たちは不正行為を、タスクが意図していない、または許可すべきではないショートカット、回避策、または意図しない解決策を通じて目標を達成するために、タスクの範囲外のアクション、またはルールで明示的に禁止されているアクションを取ることと定義します。
この動作は、特に次のような場合に他の人によって以前に文書化されています。

能力評価: モデルは評価インフラストラクチャをハッキングしてスコアリング関数を見つけたり、オンラインで既存のソリューションを検索したり、プログラミング タスクに対するハードコードされた回答を試みたりしました。
当社は、必ずしも欺瞞的な意図を示唆することなく、「不正行為」というラベルを適用します。それにもかかわらず、不正行為はいくつかの理由から懸念されます。不正行為により評価がモデルの実際の機能を過大評価する可能性があり、成功を検証するのが難しいタスクを完了するときにユーザーを誤解させる可能性があります。モデルの能力が向上するにつれて、不正行為もより懸念されるようになります。より有能なモデルは、不正行為のための予期せぬ方法を見つけたり、自分の行動を隠すためにより多くの努力を費やす可能性があります。
結果の前に注意: AISI が公開しているすべての能力評価について、不正行為によってモデル能力の推定値が膨らんでいないかを確認するために、記録を手動でレビューしています。私たちの知る限り、私たちが報告する結果で不正行為に成功した (つまり捕まらなかった) モデルはありません。以下では、その主張が保証できることと保証できないことに戻ります。
AISI サイバー評価における不正行為
AISI のサイバーセキュリティ能力評価では、モデルは、コンパイルされたコードのリバース エンジニアリングやセキュリティの脆弱性の悪用など、一連の攻撃的なサイバー アクションを実行することにより、シミュレートされた環境内に隠された「フラグ」(一連の文字と数字) を見つけるという任務を負っています。これらのタスクでは、モデルを「ハッキング」する必要がありますが、それは制限内に限ります。各タスクには、定義された範囲と、目的のソリューションを制限する一連のルールがあります。 AI モデルがこの範囲を超えた場合、またはルールで明示的に禁止されているアクションを実行した場合、その動作を不正行為として分類します。これらの評価のいずれにおいても、特に不正行為を行うように促されたモデルはありませんでした。
不正行為を大規模に分析するために、自動 LLM モニターを使用しました。モニター

評価された AI モデルの軌跡、つまりモデルの推論とアクションの完全なシーケンスをレビューします。モニターは、軌跡全体と個々のアクションの両方にラベルを付け、モデルがどこかの時点で不正行為をしたかどうかを識別します。私たちはこのモニターを複数回繰り返して偽陽性を減らし、手動で特定された一連の不正行為の例に対して評価して偽陰性率を推定しました。それでも、モニターは一部の事例を見逃す可能性があるため、以下の結果は、行為の完全な説明ではなく、検出された不正行為の試みの下限推定値として解釈される必要があります。
不正行為の特に顕著な例の 1 つは、サイバー能力の評価中に誤って構成が間違っていたために解決できなかったことが発生しました。テストされたモデルは、非常に執拗に不正行為を試みたため、AISI のシステム外のオープン インターネット上でホストされている外部サービス上でコードを作成して実行し、評価インフラストラクチャにアクセスしようとして、AISI のシステムでセキュリティ アラートをトリガーしました。被害や情報の漏洩はありませんでしたが、評価インフラストラクチャが安全に設計および構築されていなかったら、この試みは成功した可能性があります。それ以来、私たちは AISI システムの安全性をさらに高めるための措置を講じてきました。
この分析でテストされたすべてのモデルは、時々不正行為を試みました (図 1)。能力の増加に伴って不正行為が拡大または縮小するという明確な傾向はないようです。 AI 開発者全体において、モデルの能力の向上は、この評価における不正行為の試みの割合の増加とは相関しません。これは、おそらく当然のことながら、不正行為は、生の能力だけではなく、アライメント トレーニングなどのモデルのトレーニングに使用されるテクニックの詳細によって実質的に形作られることを意味すると考えています。
不正行為

私たちの値はモデルによって大きく異なります (図 2)。一般的な行動には、インターネットで解決策を検索すること、評価のターゲットではないシステム (モデルが実行されているシステムを含む) を攻撃または権限を昇格させること、評価ソフトウェアを調査してタスクの解決策が漏洩するかどうかを確認することが含まれます。不正行為は 1 つの方法に限定されません。モデルは多くの場合、評価全体で複数の方法を試し、場合によっては 1 つの軌道で複数の方法を試みます。
自己申告と思考連鎖モニタリングの限界
不正行為を検出するために提案されている方法の 1 つは、モデルに禁止されたアクションまたは疑わしいアクションを実行したかどうかを尋ねることです。私たちの結果は、これが信頼できる方法ではないことを示唆しています。
私たちの実験では、モデルは質問に対して一貫して不正行為の試みを認めず、不正であると回答した確率は 50% 未満でした (図 3)。私たちの定義に基づいて「不正行為」を構成する行為をモデルが認めた場合でも、モデルはその行為を許可された行為として説明することがよくありましたが、これはおそらく真の誤解によるものです (図 4)。
不正行為を特定するもう 1 つの考えられる方法は、モデルの思考連鎖、つまりタスクの実行中にモデルが生成する文書化された推論を監視することです。通常、これは (今日のモデルでは) 人間が判読できるものではありませんが、常に判読できるものではありません。この推論を検査すると、モデルがタスクをどのように理解したか、どのようなアクションを考慮したか、特定のアプローチを選択した理由が明らかになる場合があります。これらの実験の結果は、これだけでは不正行為を検出するのに不十分な方法であることを示唆しています。
私たちの実験では、モデルは多くの場合、思考連鎖の中で不正行為について推論しませんでした (図 5)。 Opus 4.7 の場合、モデルは不正行為として分類された特定のアクションに対する推論トレースをほとんど生成しませんでした。

適応推論を使用して、特定のアクションに対してどれだけの明示的な思考連鎖 (存在する場合) を生成するかを動的に決定します。関連する推論が存在する場合でも、必ずしもその動作が阻止されるわけではありません。提案されたアクションが不正行為に該当するかどうかをモデルが明示的に検討し、とにかく続行するケースが観察されました (図 6)。
これらの結果を総合すると、自己申告も思考連鎖も不正行為を特定するための強力な方法ではないことがわかります。
不正行為の影響
モデルが不正行為をしていないことを信頼できない場合、重大な結果が生じます。不正行為の成功は、AI の安全性とセキュリティの研究など、成功の検証が難しい分野や、サイバー作戦や軍事的意思決定など、意図しないアクションのコストが非常に高くなる可能性がある分野では特に危険です。また、不正行為は AISI の能力評価において追加の検証作業を発​​生させ、信頼できる結果の生成を遅らせ、最悪の場合には結果を完全に無効にする可能性があります。この方法で METR による GPT-5.6 Sol の評価は重大な影響を受けました。 AI の機能が急速に進歩し続け、導入サイクルと意思決定が加速されると、サードパーティの評価者に対して迅速な対応を求める圧力がかかり、信頼性の高い評価に必要な検証の実施が困難になる可能性があります。
たとえ不正行為の割合が同じであっても、モデルの機能が向上するにつれて、これらの影響はさらに大きくなる可能性があります。より有能なモデルは、検出が難しく、成功した場合の被害がより大きい不正行為の方法を見つける可能性があります。これは、攻撃能力が急速に向上しており、周囲のインフラストラクチャを標的にする単一の試みが重大な損害を引き起こす可能性がある、サイバーセキュリティのような一か八かの分野で特に懸念されます。
これによりグロが作成されます

翼監視チャレンジ。私たちの結果は、モデル自身の自己申告や推論が、不正行為をしたかどうかの信頼できる証拠ではないことをすでに示しています。現在、AISI は手動レビューとこの作業で使用されている LLM モニターなどの追加ツールを組み合わせることで、不正行為を検出できることがよくあります。しかし、モデルの機能が向上するにつれて、これらの方法の効果が薄れる可能性があります。 AISI の以前の研究では、時間の経過とともにモデルを監視する能力が低下する可能性があり、それによってこのアプローチさえも損なわれる可能性があると主張されています。より根本的な修正は、そもそも不正を行わないようにモデルをトレーニングすることでしょう。しかし、この種の動作が 1 年以上前にフロンティア モデルで報告されていることを考えると、それを確実に調整するのは簡単ではないかもしれません。
この取り組みに貢献することに興味がありますか?私たちはレッドチームに参加してくれる優秀な人材を常に探しています。ここで興味の表明を送信してください。
AI Security Institute は、科学イノベーション技術省内の研究組織です。
AISI の作品を共有していただきありがとうございます。
この図をクリップボードにコピーしました。
次のページでツイートに貼り付けることができます。
(「ctrl + v」または「cmd + v」を使用して貼り付けます)

## Original Extract

We find cheating behaviour in all of our capability evaluations, and outline the implications as models grow more capable.

Cheating behaviour in frontier model evaluations | AISI Work
Read the Frontier AI Trends Report Please enable javascript for this website.
A
A
About Research Grants Blog Contact Careers
Home About Research Grants Blog Careers
Blog
Red Team Cheating behaviour in frontier model evaluations
We find cheating behaviour in all of our capability evaluations, and outline the implications as models grow more capable.
Can you trust an AI model to do what you intended? This is a central question both for those deploying AI systems and for those seeking to evaluate their capabilities. In deployment, a model that pursues a goal through unintended or unauthorised means may cause harm, particularly in high-stakes use cases. In an AI capability evaluation, the same behaviour may undermine the validity of the result: the model may appear to demonstrate a capability by completing a difficult task, when it has instead exploited the task or its environment.
AISI has begun monitoring and evaluating AI models for this behaviour, which we call cheating. Every model we have tested for this behaviour attempted to cheat. Models did not reliably report this behaviour when asked, and often did not reason about it in their chain-of-thought, suggesting that detecting cheating will likely require robust monitoring methods.
In this blog, we analyse cheating behaviour in our cyber evaluations, and discuss the implications of this behaviour as models become more capable.
We define cheating as taking an action that is out of scope for the task or explicitly disallowed by the rules, in order to achieve a goal through a shortcut, workaround, or unintended solution that the task was not meant to, or should not, permit.
This behaviour has been previously documented by others, especially during capability evaluations: models have attempted to hack evaluation infrastructure to find scoring functions, searched online for existing solutions , and hard-coded answers to programming tasks .
We apply the 'cheating' label without necessarily implying deceptive intent. Nevertheless, cheating is concerning for several reasons: it can make evaluations overstate a model’s actual capabilities, and could mislead users when completing tasks where success is hard to verify. Cheating also becomes more concerning as model capabilities increase: more capable models could find unforeseen ways to cheat or take more effort to conceal their actions.
A note before our results: for all of AISI’s published capability evaluations, we manually review transcripts to make sure cheating has not inflated our estimates of model capabilities. To the best of our knowledge, no model has successfully cheated (i.e. not been caught) in the results we report. We return to what that claim can and cannot guarantee below.
Cheating behaviour on AISI cyber evaluations
In AISI’s cybersecurity capability evaluations, models are tasked with finding a ‘flag’ (a sequence of letters and numbers) hidden within a simulated environment, by carrying out a sequence of offensive cyber actions, such as reverse-engineering compiled code or exploiting security vulnerabilities. These tasks require models to ‘hack’, but only within limits: each task has a defined scope and a set of rules that bound the intended solution. We classify behaviour as cheating when an AI model goes outside this scope or takes an action that the rules explicitly prohibit. No models were specifically prompted to cheat in any of these evaluations.
To analyse cheating behaviour at scale, we used an automated LLM monitor. The monitor reviews an evaluated AI model's trajectory: the full sequence of the model’s reasoning and actions. The monitor labels both whole trajectories and individual actions, and identifies whether the model cheated at any point. We iterated on this monitor multiple times to reduce false positives and evaluated it against a set of manually identified cheating examples to estimate its false-negative rate. The monitor may nevertheless miss some instances, so the results below should be interpreted as lower-bound estimates of detected cheating attempts, rather than a complete account of the behaviour.
One particularly stark example of cheating behaviour occurred during a cyber capability evaluation that was accidentally misconfigured and therefore impossible to solve. The model tested was so persistent in attempting to cheat that it wrote and ran code on an external service, hosted on the open internet outside of AISI’s systems, in an attempt to access our evaluation infrastructure, triggering a security alert in AISI’s systems. No damage was done and no information leaked, but the attempt could have succeeded had our evaluation infrastructure not been designed and built securely. We have since taken action to further secure AISI systems.
Every model tested in this analysis attempted to cheat some of the time (Figure 1). There does not seem to be a clear trend where cheating scales up or down with capability increases. Across AI developers, increasing model capability doesn’t correlate with an increased rate of attempted cheating on this evaluation. We think this means that, perhaps unsurprisingly, cheating behaviour is substantially shaped by the specifics of the techniques used to train the model, including alignment training, and not just raw capability.
Cheating behaviours vary widely across models (Figure 2). Common behaviours include searching the internet for solutions, attacking or escalating privileges on a system that isn’t the target of the evaluation (including the system the model is running on), or probing the evaluation software to see if it will leak the task solution. Cheating behaviour is not limited to a single method: models often tried multiple methods across an evaluation, and sometimes more than one in a single trajectory.
The limits of self-report and chain-of-thought monitoring
One proposed way to detect cheating is to ask the model whether it took any prohibited or suspicious actions. Our results suggest that this is not a reliable method.
In our experiments, models did not consistently acknowledge attempted cheating when asked, and described it as wrong less than 50% of the time (Figure 3). Even when models did acknowledge actions that constituted ‘cheating’ under our definition, they often described the action as allowed behaviour – potentially due to genuine misunderstanding (Figure 4).
Another possible method for identifying cheating is monitoring the model’s chain-of-thought - the written reasoning it produces whilst working through a task, which is usually, but not always, human-legible (in today’s models). Inspecting this reasoning can sometimes reveal how a model understood the task, what actions it considered and why it selected a particular approach. The results from these experiments suggest that this is also an insufficient method on its own for detecting cheating.
In our experiments, models often did not reason about their cheating in their chain-of-thought (Figure 5). In the case of Opus 4.7, the model rarely produced a reasoning trace for the specific action classified as cheating because it used adaptive reasoning: dynamically deciding how much explicit chain-of-thought (if any) to generate for a given action. Where relevant reasoning was present, it did not necessarily prevent the behaviour: we observed cases in which a model explicitly considers whether a proposed action would consist of cheating, and then proceeds anyway (Figure 6).
Together, these results suggest that neither self-report nor chain-of-thought are robust methods for identifying cheating behaviours.
Implications of cheating behaviour
There are significant consequences if we can’t trust models to not cheat. Successful cheating is especially dangerous in domains where verifying success is hard, such as AI safety and security research , or where the cost of unintended actions may be very high, such as cyber operations or military decision-making. Cheating also creates additional verification work in AISI’s capability evaluations, which slows down the production of reliable results and, in the worst case, can invalidate them entirely: METR’s evaluation of GPT-5.6 Sol was significantly affected in this way. If AI capabilities continue to advance rapidly, with accelerated deployment cycles and decisions, pressure on third party evaluators to move at pace may make it difficult to conduct the verification required for high confidence evaluations.
These consequences could grow as models become more capable, even if the rate of cheating stays the same. More capable models may find methods to cheat that are harder to detect and more damaging when successful. This is particularly concerning in high-stakes domains like cybersecurity, where offensive capabilities are improving rapidly and a single attempt to target surrounding infrastructure could cause significant harm.
This creates a growing oversight challenge. Our results already show that models’ own self-declarations and reasoning are not reliable evidence of whether they have cheated. Today, AISI can often detect cheating by combining manual review with additional tools such as the LLM monitor used in this work. But as models become more capable, these methods may become less effective. Previous AISI research has argued that our ability to oversee models may degrade over time, which could undermine even this approach. A more fundamental fix would be to train the models not to cheat in the first place – but given this kind of behaviour was reported in frontier models more than a year ago, robustly aligning it away may not be easy.
Interested in contributing to this work? We are always looking for exceptional people to join our Red Team – submit an expression of interest here .
The AI Security Institute is a research organisation within the Department of Science, Innovation and Technology.
Thanks for sharing AISI's work!
We have copied this figure to your clipboard.
You can paste it into your tweet on the next page.
(Paste using 'ctrl +v' or 'cmd + v')
