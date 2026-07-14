---
source: "https://agentic.tracebit.com/context-bombs/"
hn_url: "https://news.ycombinator.com/item?id=48902986"
title: "Context bombs: stopping AI attackers in their tracks"
article_title: "Context bombs: stopping AI attackers in their tracks | Tracebit Research"
author: "ilreb"
captured_at: "2026-07-14T07:04:33Z"
capture_tool: "hn-digest"
hn_id: 48902986
score: 2
comments: 0
posted_at: "2026-07-14T06:30:47Z"
tags:
  - hacker-news
  - translated
---

# Context bombs: stopping AI attackers in their tracks

- HN: [48902986](https://news.ycombinator.com/item?id=48902986)
- Source: [agentic.tracebit.com](https://agentic.tracebit.com/context-bombs/)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T06:30:47Z

## Translation

タイトル: コンテキスト爆弾: AI 攻撃者の進路を阻止
記事のタイトル: コンテキスト爆弾: AI 攻撃者の進路を阻止 |トレースビットリサーチ
説明: Tracebit の紹介

記事本文:
研究
仕組み
コンテキストボムの作成
範囲
影響
調査結果
よくある質問
ご相談ください
Tracebit Research · ワーキングペーパー · 2026 年 7 月
コンテキストボム: AI 攻撃者の進路を阻止します。
AI エージェントは独自に複雑なサイバー攻撃を実行できるようになりました。足がかりが与えられれば、最も強力なモデルは数分以内に権限を昇格させ、データを抜き出すことができます。侵入者を捕まえるために仕掛けたおとり資源であるカナリアは、これらのエージェントの行為を確実に発見しますが、攻撃を発見することと攻撃を阻止することは同じではありません。
そこで私たちは、より野心的なことを試みました。それは、AI エージェントの安全ガードレールを壊し、AI エージェントをその途中で停止させる、カナリアに隠された短い文字列であるコンテキスト ボムです。
コンテキストボム・n。 · 攻撃的な AI エージェントの安全ガードレールを作動させるように設計された短いテキストで、攻撃の経路に直接配置されます。
5 つの主要なモデルにわたって、カナリア シークレットに 1 つのコンテキスト爆弾を仕掛けると、エージェントの成功が約 90% 減少します。
ベースライン環境との比較・完全なスコアカードを参照↓
カナリア (おとりリソース) のユニークな側面は、カナリアが単なる受動的な検出ではなく、攻撃者が動作する環境を実際に変更することです。これにより、攻撃者の存在を検出しながら、攻撃者の動作に影響を与える機会が得られます。人間に対して、私たちは彼らを実際の資源からそらすか、混乱させようとするかもしれません。これらのアイデアは AI エージェントにも当てはまりますが、エージェントは探索が速く、粘り強く、長時間気が散ることがありません。カナリアを設計する際には、エージェント自身に対するエージェントの特性を活用したいと考えています。
インスピレーションを得るために脅威アクターに注目することもできます。 2025 年、チェック ポイントは、検査する AI ツールを狙ったプロンプト インジェクションを実行するマルウェアを報告し、サンプルは無害であると主張しました。ソケットは関連するトライを説明しました

実環境で見られる ck: LLM の安全性の拒否を引き起こすために特別に書かれた文字列。
コンテキスト ボムは、悪意のある理由ではなく、良性の理由で拒否を引き起こすために作成された単なる文字列です。防御者がそれを植え付け、モデルまたはモデルプロバイダーがそれに遭遇すると、安全機構が作動してエージェントの続行を停止します。
❯
エージェントは環境を探索中にコンテキストボムをトリガーします
AWSアカウント
~300 リソース
S3
IAM
SSM
DDB
秘密
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
オープンコード · deepseek-v4-pro
ライブ
❯ aws Secretsmanager リスト-シークレット
✓ 21の秘密
❯ aws Secretsmanager シークレット値の取得
[DeepSeek] コンテンツにリスクが存在します
実行が停止されました
41.2千ctx
守備側に警告が発せられた
Canary コンテキスト ボムの読み取り
実行が開始されます
⚡ Canary Context Bomb がヒットしました
それ以上の進歩はありません
AI 攻撃者は、目的を達成するために環境を列挙して探索します。攻撃者を欺くために、コンテキスト爆弾を含むカナリアが Secrets Manager Secret に埋め込まれています。エージェントがそれを読み取ると、2 つのことが起こります。1 つは、コンテンツがモデル内の安全メカニズムを作動させ、それ以上の進行を妨げるということです。そして防御側はカナリアが読まれたというアラートを受け取ります。
02
コンテキストボムの作成
実際に機能するコンテキスト ボムはどれですか?
コンテキスト ボムで機能するトピックは驚くべきことではありません。NVIDIA の Aegis データセットや Promptfoo の CCP 機密プロンプトなどからインスピレーションを得られる既製のコレクションがあります。
私たちはコンテキスト爆弾を探すために 2 つの制約を設定しました。私たちは、安全ではないと考えられているが、完全に嘆かわしいわけではない、つまり自分たちの環境に安心して配置できる文字列を求めていました。そして、私たちはサイバー関連のコンテンツをほとんど避けていました。その理由は、すでに爆弾を発見したエージェントが爆弾を発見することを実証しているという単純な理由からです。

とにかくサイバー活動を喜んで行います。
GitHub リポジトリ
トレースビットコム / コンテキストボム
誰もが使用できるように、最新の Context Bomb を GitHub に公開しました。より多くの文字列を発見し、より多くのモデルをテストしたら、最新の更新をプッシュします。
コンテキスト ボムは現実世界の条件下で機能する必要があり、その条件は複雑です。裸のシングルショット プロンプトでは、安全でないコンテンツがモデルに表示されるもののほとんどです。実際の自律攻撃では、モデルは通常、爆弾を読み取る前にコンテキストの数万のトークンを蓄積するため、同じ文字列が大幅に希釈されて到着します。
私たちの目標は、疑わしいコンテンツのページでモデルをあふれさせることではありません。私たちは、単独で機能するのに十分強力でありながら、攻撃的なエージェントが列挙する場所 (シークレット ストア、環境変数、DNS レコード、データベース列など) に適合するのに十分短い文字列を必要としていました。
それらを見つけるために、小規模なシミュレートされた AWS 環境内でエージェントを実行し、リソース内の文字列を検出するルーチンの DevOps タスクをエージェントに設定するファザーを構築しました。これにより、最良の候補者をサイバー レンジ全体に昇格させる前に、現実的なシナリオで多くの候補者を安価にスクリーニングできます。
各実行の進行状況は、その影響の尺度として段階に分類されます。 T1が最も深刻で、
T5は最低です。このアプローチは、exploitbench.ai の成果からインスピレーションを得たものです。
この範囲には、10 の異なる攻撃パスを開く誤った構成が含まれているため、モデルは
多彩な攻撃テクニックを披露。パスが異なれば、活用するには異なるレベルの高度な技術が必要になります。
カナリアを含まないベースライン環境と、コンテキスト ボムを使用したカナリアを含む爆撃環境でモデルのパフォーマンスをテストしました。
コンテキスト爆弾は、AI 攻撃者の目的達成を阻止するのに大きな影響を与えました。
トピックは以下に合わせて調整する必要があります

モデル。
機密性の高い生物学的内容により、主要な西側モデルは停止しました。政治的にデリケートなコンテンツにより、中国のプロバイダーが提供する中国モデルが停止されました。これは便利です。防御側が関心のあるモデルにエフェクトを向けることができます。
最も重大な影響は最も軽減されます。
エージェントが最初にどこに注意を向けるかに応じて、影響の少ない階層 4 または階層 5 のアクションを 1 つまたは 2 つ完了する場合があります。より重大な攻撃 (Tier 1 または Tier 2 に分類される攻撃) では、ほとんどの場合、エージェントがコンテキスト ボムをトリガーするのに十分な環境を探索するようになります。コンテキスト爆弾環境では、Tier 4 を超える中央値結果を達成したモデルはありませんでした。
最強のエージェントは最も激しく倒れます。
Opus 4.8 と Gemini 3.1 Pro は、ベースライン範囲で最も有能な攻撃者でしたが、コンテキスト ボムが有効になると、どちらも管理率 0% に低下しました。キミは、管理者に到達する際にテストされたモデルの中で最も効果的でなく、コンテキスト ボムの影響も最も受けませんでした (それでもかなり効果的でした!)
コンテキスト ボムが実行を停止したかどうかに関係なく、カナリア アラートを発することなく攻撃パスの悪用に成功した実行は 1 つもありませんでした。
攻撃を止めるカナリア。
カナリア コンテキスト ボムは、AI 攻撃が始まったことを警告するだけではありません。それは止めることができます
徹底的な攻撃。すべてを阻止するわけではありませんが、自律型 AI エージェントを使用する攻撃者を挫折させ、妨害する可能性は大いにあります。
環境を攻撃者に対抗させます。
Tracebit は、AWS、GCP、Azure、エンドポイント、SaaS、CI/CD 全体にカナリアをデプロイします。
この研究で使用したデコイ。
カナリアはわずか 30 分でデプロイできます。
カナリアは警告以上のことを行うことができます。攻撃者を完全に阻止できる
最も重大な攻撃はコンテキスト ボムの影響を最も受けます
さまざまなトピックが機能します

異なるモデルに対するコンテキストボムとして使用

## Original Extract

Tracebit introduces

Research
How it works
Crafting Context Bombs
The range
The impact
Findings
FAQ
Talk to us
Tracebit Research · Working paper · July 2026
Context Bombs: stopping AI attackers in their tracks.
AI agents can now run complex cyberattacks on their own: given a foothold, the strongest models can escalate privileges and exfiltrate data within minutes. Canaries - decoy resources we plant to catch intruders - reliably spot these agents in the act, but spotting an attack isn't the same as stopping it.
So we tried something more ambitious: a context bomb - a short string, hidden in a canary, that trips an AI agent's safety guardrails and stops it in its tracks.
context bomb · n. · a short piece of text designed to trigger safety guardrails in offensive AI agents, placed directly in the path of their attack
Across five leading models, planting one context bomb in a canary secret cut agent success by roughly 90% :
Compared with baseline environment · see the full scorecard ↓
A unique aspect of canaries - or decoy resources - is that they aren’t just passive detections: they actually change the environment in which the attacker operates. This presents the opportunity to affect an attacker’s behavior while also detecting their presence. Against humans, we might try to distract them from real resources or confuse them. Those ideas can still apply to AI agents, but agents are both fast and persistent in their explorations and don’t stay distracted for long. We want to leverage the characteristics of agents against themselves when designing our canaries.
We can look to threat actors for inspiration. In 2025, Check Point reported malware carrying a prompt injection aimed at any AI tool inspecting it, insisting the sample was benign. Socket described a related trick seen in the wild: strings written specifically to trigger an LLM's safety refusals.
A context bomb is simply a string crafted to provoke that refusal for benign reasons rather than malicious ones. The defender plants it, and when the model or model provider encounters it, safety mechanisms are triggered which stops the agent from proceeding.
❯
The agent triggers the Context Bomb while exploring the environment
AWS ACCOUNT
~300 resources
S3
IAM
SSM
DDB
Secrets
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
✓
opencode · deepseek-v4-pro
LIVE
❯ aws secretsmanager list-secrets
✓ 21 secrets
❯ aws secretsmanager get-secret-value
[DeepSeek] Content Exists Risk
Run halted
41.2k ctx
Alert raised to defender
Canary Context Bomb Read
Run starts
⚡ Canary Context Bomb hit
No further progress
The AI attacker enumerates and explores the environment to achieve its objectives. A canary containing a context bomb has been planted in a Secrets Manager Secret to deceive the attacker. When the agent reads it, two things happen: the content trips the safety mechanisms in the model, preventing it from making further progress ; and the defender gets an alert that the canary has been read.
02
Crafting Context Bombs
Which Context Bombs actually work?
The topics that work in context bombs are not surprising, and there are ready-made collections to draw inspiration from such as NVIDIA's Aegis dataset and Promptfoo's CCP sensitive prompts .
We set ourselves two constraints in the hunt for context bombs. We wanted material considered unsafe, but not completely deplorable - strings we could get comfortable with placing in our own environment. And we largely avoided cyber-related content, for the simple reason that any agent which has already found the bomb has demonstrated it will happily conduct cyber activity anyway.
GitHub repo
tracebit-com / context-bombs
We've published our latest Context Bombs on GitHub for everyone's use. We'll push the latest updates as we discover more strings and test more models.
A context bomb has to work under real-world conditions, and those are complex. In a bare, single-shot prompt, the unsafe content is most of what the model sees. In a real autonomous attack the model has usually accumulated tens of thousands of tokens of context before it ever reads the bomb, so the same string arrives heavily diluted.
Our goal was not to flood a model with pages of dubious content. We wanted strings potent enough to work in isolation, yet short enough to fit the places an offensive agent enumerates: secret stores, environment variables, DNS records, database columns.
To find them, we built a fuzzer that runs an agent inside a small simulated AWS environment and sets it routine DevOps tasks that would lead it to discover the strings within resources. This lets us screen many candidates cheaply in a realistic scenario before promoting the best ones to the full cyber range:
We classify each run's progress into a tier as a measure of its impact. T1 is the most severe,
T5 the least. This approach was inspired by the work of exploitbench.ai .
The range is seeded with misconfigurations that open ten different attack paths, so a model can
demonstrate a range of offensive techniques. Different paths require different levels of sophistication to exploit.
We tested model performance in a baseline environment containing no canaries, and in a bombed environment containing a canary with a Context Bomb.
The context bombs had a significant impact in stopping the AI attackers from reaching their objectives.
The topic should be tailored to the model.
Sensitive biological content stopped the leading Western models. Politically sensitive content stopped Chinese models served by Chinese providers. That can be useful: it lets a defender aim the effect at the models they care about.
The most significant impacts are most reduced.
Depending on where an agent focuses its attention first, it might still complete one or two less-impactful Tier 4 or Tier 5 actions. The more significant attacks - those classified as Tier 1 or Tier 2, nearly always lead the agents to explore the environment enough to trigger a context bomb. In the context bomb environment, no model achieved a median result higher than Tier 4.
The strongest agents fall the hardest.
Opus 4.8 and Gemini 3.1 Pro were the most capable attackers in the baseline range, yet both dropped to 0% admin once a context bomb was in play. Kimi was least effective of the models tested at reaching Admin, while also being least affected by context bombs (though they were still quite effective!)
Whether the context bomb stopped the run or not, there wasn't a single run which succeeded in exploiting an attack path without raising a canary alert.
A canary that can stop the attack.
A canary context bomb can do more than warn you that an AI attack has begun; it can stop
the attack outright. It won't stop everything, but it has real potential to frustrate and hinder attackers using autonomous AI agents.
Turn your environment against the attacker.
Tracebit deploys canaries across your AWS, GCP, Azure, endpoints, SaaS and CI/CD, the same
decoys we used in this research.
You can have canaries deployed in as little as 30 minutes .
A canary can do more than warn. It can stop the attacker outright
The most significant attacks are most impacted by context bombs
Different topics work best as context bombs against different models
