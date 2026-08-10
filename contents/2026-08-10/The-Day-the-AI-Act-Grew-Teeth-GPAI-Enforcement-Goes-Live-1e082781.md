---
source: "https://simonroses.com/2026/08/the-day-the-ai-act-grew-teeth-gpai-enforcement-goes-live/"
hn_url: "https://news.ycombinator.com/item?id=49246042"
title: "The Day the AI Act Grew Teeth: GPAI Enforcement Goes Live"
article_title: "The Day the AI Act Grew Teeth: GPAI Enforcement Goes Live | Simon Roses Femerling – Blog"
author: "speckx"
captured_at: "2026-08-10T16:43:27Z"
capture_tool: "hn-digest"
hn_id: 49246042
score: 1
comments: 0
posted_at: "2026-08-10T16:34:51Z"
tags:
  - hacker-news
  - translated
---

# The Day the AI Act Grew Teeth: GPAI Enforcement Goes Live

- HN: [49246042](https://news.ycombinator.com/item?id=49246042)
- Source: [simonroses.com](https://simonroses.com/2026/08/the-day-the-ai-act-grew-teeth-gpai-enforcement-goes-live/)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T16:34:51Z

## Translation

タイトル: AI 法が誕生した日: GPAI の施行が本格化
記事のタイトル: AI 法が誕生した日: GPAI の施行が本格化 |サイモン・ローゼズ・フェマーリング – ブログ

記事本文:
サイモン・ローゼズ・フェマーリング – ブログ
サイバースペースの不安 3.X
コンテンツにスキップ
ホーム
AI 法が本格化した日: GPAI の施行が本格化
2026 年 8 月 2 日、誰もが黙って無視していた EU AI 法の一部が施行可能になりました。AI 局は汎用 AI モデルのプロバイダーに対し、世界の年間売上高の 3% または 1,500 万ユーロのいずれか高い方の罰金を科すことができ、さらに技術文書の要求、モデルの独自の評価の実施、「措置を取る」よう命令することができ、最悪の場合は EU 市場からのモデルの制限、撤退、リコールを強制することができます。この義務自体は技術的には 2025 年 8 月 2 日から存在していました (技術文書、下流の透明性、robots.txt とオプトアウトを尊重する著作権ポリシー、トレーニング データの公開概要) が、これまでは審査員のいないルールでした。それは変わりました。システミック リスクを伴うと判断されたモデル (10^25 FLOP を超えてトレーニングされたモデル) には、規制当局によって作成されたセキュリティ チェックリストのような、より重い義務が課せられます。つまり、モデルの評価、敵対的テスト / レッドチーム、安全性とセキュリティのフレームワーク、AI オフィスへの重大インシデントの報告、およびモデルの重み自体のサイバーセキュリティ保護です。自主的な行動規範は、免責を与えるものではなく、軽い気持ちを与えるものです。 2025 年 8 月 2 日より前にすでに発売されているモデルは、2027 年 8 月 2 日までにラインナップされます。これは 2 部構成のストーリーの規制に関する部分です。民事責任に関する部分である新しい製造物責任指令については、この記事の直後に私が公開する投稿です。合わせて読んでください。EU はあなたに罰金を課す規制当局と、あなたに請求を行う法廷という挟み撃ちを構築しました。以下は、セキュリティ専門家である私の読んだ規制上の課題です。
いつもの免責事項（いつもと同じ）：私は弁護士ではないので、これは法的アドバイスではありません

— それは、私が攻撃対象領域を読むのと同じように規制を読み、実際に圧力がかかる場所と最終的に誰がそれを保持することになるかを探すセキュリティ専門家です。 AI モデルを構築したり EU に出荷したりする場合は、実際の弁護士に相談してください。私が言えることは、これがこれらのモデルを構築、保護、展開する人々にとって運用面でどのような変化をもたらすかということです。なぜなら、コンプライアンスの文言の下に埋もれているのは、私がこのブログで 2 年間要求してきたことのリストであり、現在は罰金に裏付けられているからです。
始める前にフレームに関するメモ。この投稿はペアのうちの 1 つです。 EU は同時に 2 つの異なる軌道で AI に歯を立てており、彼らは異なる方法で噛みつきます。これは、AI 法の汎用 AI に関する規則であり、公的機関である AI 局が調査し、罰金を課す権限を持っています。新しい製造物責任指令に関する関連記事は、民事に関するものです。民間の原告と裁判所、厳格責任、欠陥のあるソフトウェアが損害を与えた人に支払われる損害賠償です。 1 つだけを追跡すると、露出の判断を誤ってしまうため、意図的にこれらを連続して公開しています。あなたに罰金を科す規制当局とあなたを訴える請求者は 2 つの別のドアであり、この夏以降はどちらも開かれます。
私はしばらくの間、規制の扉をくぐり抜けてきました。なぜ「ChatGPT を使用する」が AI 戦略ではないのか、モデル自体が攻撃者になったときに何が起こるのか、オープンウェイト モデルとフロンティア モデルをまだ区別できるのかどうかなどです。 8月2日は、EUが執行予算を添付してこれらの質問の一部に回答する日である。
8月2日に実際に変わったこと
ここが人々を混乱させる部分なので、正確に言わせていただきますが、2026 年 8 月 2 日によって新たな義務が創設されたわけではありません。汎用 AI (GPAI) モデルの実質的な規則は、その 1 年前の 2025 年 8 月 2 日に施行されました。

今まで欠けていたのは執行機械だった。 12 か月間、AI 法の GPAI 章は、誰もそれについて何もできなくても、技術的に破ることができる法律でした。
その猶予期間は終わりました。 2026 年 8 月 2 日の時点で、委員会の専任 AI 執行機関である AI Office は、8 月 1 日にはなかった 4 つの具体的な権限を持ちます。
資料を要求してください。 GPAI プロバイダーに、モデルに関する技術文書と情報を引き渡すことが必要になる場合があります。
モデルを評価します。モデルの独自の評価を実行して、コンプライアンスをチェックし、アクセスの要求などのシステムリスクを調査できます。
注文コンプライアンス措置。モデルを一致させるために「適切な措置を講じる」ことが必要になる場合があります。
モデルを引っ張ります。最悪の場合、その入手可能性を制限したり、EU市場から撤回したり、リコールしたりする可能性があります。
これを読んでいる弁護士は、AI 庁が調査、評価、訴訟を提起する運営機関であるが、第 101 条に基づく罰金の正式な決定は欧州委員会が行うことになるため、正確に押さえておく価値がある点が 1 つあります。実際には AI オフィスとやり取りします。ペナルティへの署名は委員会のものである。
そして、4 つすべての背後には、注目を集める数字があります。第 101 条に基づく罰金は、世界の年間売上高の 3% または 1,500 万ユーロのいずれか高い方です。これは GPAI 固有の上限であることに注意してください。同法の見出しである売上高の 7% の罰金は、禁止されている AI 慣行の導入に対するものであり、これとは異なる制度です。しかし、フロンティアラボにとって、世界売上高の 3% は取締役会レベルの数字です。
図 1. 執行が実際にどのように行われるか: AI 局のエスカレーションのはしご (文書要求 → モデル評価 → 遵守命令) を乗り越えた不遵守ギャップは、最大 3% の罰金で終わる

総売上高または 1,500 万ユーロ、極端な場合には EU 市場からの制限または撤退が考えられます。
つまり、今週モデルの義務については何も変わりませんでした。変わったのは、それらを無視することには代償があり、審判がおり、停止ボタンがあることです。
そして、この日付が本当であることがわかります。 EUのデジタルオムニバス（業界が熱心に働きかけ、2025年11月に提出され、今春合意された簡素化パッケージ）は、AI法の高リスク期限を延期し、付属書IIIの義務を2027年12月に、組み込みシステムを2028年にスライドさせた。これにより、GPAIは放置された。ブリュッセルが移動を迫られたすべての期限のうち、移動しなかった期限は、この投稿が取り上げている期限である。あなたが書いていること以外のほぼすべてでレギュレータが点滅する場合、それが優先されます。
AI法は役割にうるさく、細かい部分が重要だ。 GPAI の義務はモデルのプロバイダー、つまり汎用モデルをトレーニングして市場に出すラボに課されます。明らかなフロンティアの名前だけでなく、成長を続けるオープンウェイト ラボの分野、そして重要なことに、モデルを微調整したり大幅に変更して、事実上新しいプロバイダーになる人も考えてください。委員会独自のガイダンスでは、これについて大まかな境界線が定められています。元のトレーニング コンピューティングの約 3 分の 1 以上を使用してモデルを変更すると、ユーザー自身がプロバイダーになったとみなされ、それに伴う義務が課せられます。この条項は、自分たちを単なる「ユーザー」だと思っている多くの企業を、気付かないうちにその対象に引き込んでしまう条項です。
あなたがデプロイ担当者である場合、つまり他人のモデルに基づいて製品を構築する場合、これらの GPAI の義務のほとんどはまだ直接あなたのものではありません (ハイリスク システムのルールが適用されれば、あなたの日は後で来ます)。タイムラインはデジタル オムニバスによって延期されたばかりで、

技術基準の条件なし）。しかし、その結果はあなたにも受け継がれます。上流のプロバイダーが提供する必要がある透明性情報は、まさにあなた自身のコンプライアンスとセキュリティレビューが依存する資料です。セキュリティ専門家が最初に気を引き締めるべきなのは、この法律により、モデル ベンダーに対し、これまで企業秘密として扱っていたことを伝えるよう強制されているということです。それを使ってください。
ベースライン: すべての GPAI プロバイダーが現在負っているもの
EU 市場のすべての汎用モデルには、サイズに関係なく、次の 4 つの義務があります。
技術文書。モデルに関する詳細な維持管理書類 (アーキテクチャ、トレーニング プロセス、使用目的と除外用途、エネルギー消費量) は 10 年間保存され、要求に応じて AI オフィスに提出されます。これは調査で読み上げられるファイルです。
下流の透明性。連絡先の詳細を公開し、モデルを責任を持って統合するために必要な情報を下流プロバイダーの要求に応答する必要があります。実務規範の透明性ガイダンスでは、応答期間が短く数日であるとされています (法律事務所の測定値では約 14 日とされています)。同時に、正当な知的財産と企業秘密を保護する必要があります。モデル統合の「すべて独自のもので、自分で考えてください」という時代は終わりつつあります。
実際の仕組みを備えた著作権ポリシー。意図の段落ではありません。技術的保護措置を尊重し、既知の著作権侵害ソースを除外し、robots.txt と機械読み取り可能なオプトアウト信号を尊重し、権利所有者に連絡先と苦情経路を提供する作業方針です。これがトレーニングデータ訴訟の対象となる条項だ。
公開トレーニングデータの概要。モデルがトレーニングされた内容を要約した、入力済みの AI Office テンプレート。データセットではありませんが、ブラック ボックスにラベルを付けるには十分な概要です。

成熟したエンジニアリング工場を経営したことのある人にとって、これは珍しいことではありません。新しいのは、それが必須であり、強制可能であり、発見可能であることです。
システミックリスク：規制が私の言葉を話し始めるとき
この記事を書こうと思ったきっかけとなった部分です。モデルのサブセット、つまり「影響の大きい機能」（トレーニングの計算が 10^25 FLOP を超えると推定される）を備えたモデルと、委員会が指定するその他のモデルは、システミック リスクを伴うものとして分類されます。そして、それらに付随する義務は、私自身のエンゲージメントチェックリストの 1 つから直接得られたものかもしれません。
安全性とセキュリティのフレームワーク (ここでのタイミングの詳細は、行動規範の安全性とセキュリティの章から引用されています) は、通知から数週間以内に策定され、モデルの出荷前に完成しました。
モデルの評価と敵対的テスト — 同法では、レッドチーム化を声高に謳っています。モデルの危険な機能を最新鋭で評価することは、現在では法的な義務となっており、予算を争う安全チームにとっては嬉しいことではありません。
ライフサイクル全体にわたるシステムリスクの評価と軽減: フィルタリング、モニタリング、入出力制御。
重大なインシデントについては、時差のあるタイムラインで AI オフィスおよび各国当局に報告します。これは、モデルの実際のインシデント対応義務です。
モデルとその物理インフラストラクチャのサイバーセキュリティ保護、つまり重りの保護。体重の流出は今や、単なる恥ずかしい見出しではなく、コンプライアンス違反となっています。
文書は 10 年間保存され、外部評価者の調査結果を含む安全性とセキュリティのレポートが提供されます。
調達に関する公平な注意: 法律自体がこれらの義務を原則のレベルで定めています (第 55 条)。上記の具体的な詳細のいくつか（フレームワークのタイミング、安全性報告書の正確な形式）は、行動規範の安全性から来ています。

nd-security 章は、法定基準を満たしていることを証明するための舗装された道路です。義務は法律です。詳細の一部はコードです。
そのリストを読んでから、Hugging Face / OpenAI モデル評価インシデントの後に私が書いたことを読み直してください。ほとんどの防御側はモデル層とエージェント層でテレメトリがゼロであり、マシンスピードで完全な侵入を実行するツールは、独自の安全性テスト中に偶然トリガーできるものになりました。 AI 法は、まさにその層に対するテレメトリ、評価、インシデント報告をシステミック リスク モデルの規制上の義務としたところです。私はこの法律のすべての行が気に入っているわけではありませんが、地球上で最も有能なモデルに敵対的テストと重量セキュリティを義務付けることが、文句を言う価値のある部分であると主張するつもりはありません。私が求めていた部分です。
セキュリティ経済上の微妙な点もあります。敵対的テストは、敵対的テストの効果を発揮します。厳格さを定義せずにレッドチームを要求する規制は、フレンドリーな社内チーム、週末、グリーンレポートといったチェックボックスバージョンを招きます。これを実際の攻撃的な研究（外部的、適応的、実際にモデルを破るように動機付けられた）として扱う研究所は、何か意味のある安全性の証拠を生み出すでしょう。それをコンプライアンス成果物として扱うものは、調査で誰かが読むまでは見栄えの良い文書を作成します。

[切り捨てられた]

## Original Extract

Simon Roses Femerling – Blog
CyberSpace Insecurity 3.X
Skip to content
Home
The Day the AI Act Grew Teeth: GPAI Enforcement Goes Live
On August 2, 2026 , the part of the EU AI Act everyone was quietly ignoring became enforceable: the AI Office can now fine providers of general-purpose AI models up to 3% of global annual turnover or €15 million, whichever is higher , and can demand your technical documentation, run its own evaluations of your model, order you to “take measures,” and in the worst case make you restrict, withdraw, or recall the model from the EU market . The obligations themselves have technically existed since August 2, 2025 — technical documentation, downstream transparency, a copyright policy that honors robots.txt and opt-outs, a public summary of training data — but until now they were rules without a referee. That changed. Models judged to carry systemic risk (trained above 10^25 FLOP) carry heavier duties that read like a security checklist written by a regulator: model evaluations, adversarial testing / red-teaming , a safety-and-security framework, serious-incident reporting to the AI Office, and cybersecurity protection of the model weights themselves. The voluntary Code of Practice buys you a lighter touch, not immunity. Models already on the market before August 2, 2025 get until August 2, 2027 to fall in line. This is the regulatory half of a two-part story — the civil-liability half, the new Product Liability Directive, is the post I’m publishing right after this one. Read together, the EU has built a pincer: a regulator that fines you, and a courtroom that bills you. Here’s my security-practitioner read of the regulatory jaw.
The usual disclaimer, same as always: I am not a lawyer, and this is not legal advice — it’s a security practitioner reading a regulation the way I read an attack surface, looking for where the pressure actually lands and who ends up holding it. If you build or ship AI models into the EU, talk to actual counsel. What I can tell you is what this changes operationally for the people who build, secure, and deploy these models — because buried under the compliance language is a list of things I have been demanding on this blog for two years, now backed by a fine.
A note on framing before we start. This post is one of a pair. The EU is putting teeth behind AI on two different tracks at once , and they bite in different ways. This one — the AI Act’s rules for general-purpose AI — is regulatory : a public authority, the AI Office, with the power to investigate you and fine you. The companion piece, on the new Product Liability Directive , is civil : private plaintiffs and courts, strict liability, damages paid to the person your defective software harmed. I’m publishing them back to back on purpose, because if you only track one you’ll misjudge your exposure. A regulator fining you and a claimant suing you are two separate doors, and after this summer both are open.
I have been circling the regulatory door for a while — why “we use ChatGPT” isn’t an AI strategy , what happens when the model itself becomes the attacker , whether you can still tell an open-weight model from a frontier one . August 2 is the EU answering a slice of those questions with an enforcement budget attached.
What Actually Changed on August 2
Here’s the part that confuses people, so let me be precise: August 2, 2026 did not create new obligations. The substantive rules for general-purpose AI (GPAI) models kicked in a full year earlier, on August 2, 2025. What was missing until now was the enforcement machinery. For twelve months the AI Act’s GPAI chapter has been law you could technically break without anyone able to do much about it.
That grace period is over. As of August 2, 2026, the AI Office , the Commission’s dedicated AI enforcement body, has four concrete powers it did not have on August 1:
Demand your documentation. It can require a GPAI provider to hand over technical documentation and information about the model.
Evaluate your model. It can run its own assessments of your model to check compliance and investigate systemic risk — including requesting access.
Order compliance measures. It can require you to “take appropriate measures” to bring the model into line.
Pull the model. In the worst case it can make you restrict its availability, withdraw it, or recall it from the EU market.
One precision point worth keeping straight, because the lawyers reading this will: the AI Office is the operational body that investigates, evaluates, and builds the case, but the formal decision to fine under Article 101 is the European Commission’s . In practice you deal with the AI Office; the signature on the penalty is the Commission’s.
And behind all four sits the number that focuses minds: fines of up to 3% of global annual turnover or €15 million, whichever is higher , under Article 101. Note that’s the GPAI-specific ceiling; the Act’s headline 7%-of-turnover fines are for deploying prohibited AI practices, a different regime. But for a frontier lab, 3% of global turnover is a board-level number.
Figure 1. How enforcement actually lands: a non-compliance gap that survives the AI Office’s escalation ladder — documentation request → model evaluation → compliance order — ends in a fine of up to 3% of global turnover or €15M, and at the extreme, restriction or withdrawal from the EU market.
So nothing about your model’s obligations changed this week. What changed is that ignoring them now has a price, a referee, and a stop button.
And here is the tell that this date is real. The EU’s Digital Omnibus — a simplification package the industry lobbied for hard, tabled in November 2025 and agreed this spring — postponed the AI Act’s high-risk deadlines, sliding the Annex III obligations to December 2027 and embedded systems into 2028. It left GPAI alone. Of all the deadlines Brussels was pressed to move, the one it would not move is the one this post is about. When a regulator blinks on nearly everything except the thing you are writing about, that thing is the priority.
The AI Act is fussy about roles, and the fine print matters. The GPAI obligations land on the provider of the model — the lab that trains and places the general-purpose model on the market. Think the obvious frontier names, but also the growing field of open-weight labs, and, importantly, anyone who fine-tunes or substantially modifies a model to the point of becoming, in effect, its new provider. The Commission’s own guidance puts a rough line on that: modify a model using more than about a third of its original training compute and you’re presumed to have become a provider yourself, with the obligations that follow. That clause is the one that pulls a lot of companies who think of themselves as mere “users” into scope without noticing.
If you’re a deployer — you build a product on top of someone else’s model — most of these GPAI duties are not directly yours yet (your day comes later, once the high-risk-system rules bite — a timeline the Digital Omnibus just pushed back and made conditional on technical standards). But you inherit the consequences: the transparency information your upstream provider must now give you is exactly the material your own compliance, and your own security review, depend on. Which is the first place a security practitioner should perk up: the Act is forcing your model vendor to tell you things they previously treated as trade secrets. Use that.
The Baseline: What Every GPAI Provider Now Owes
For every general-purpose model on the EU market, regardless of size, four duties:
Technical documentation. A detailed, maintained dossier on the model — architecture, training process, intended and excluded uses, energy consumption — retained for ten years and produced to the AI Office on request. This is the file that gets read aloud in an investigation.
Downstream transparency. You must publish contact details and respond to downstream providers’ requests with the information they need to integrate the model responsibly — the Code of Practice’s transparency guidance points at a short, days-long response window (law-firm readings cite roughly 14 days) — while still protecting legitimate IP and trade secrets. The “it’s all proprietary, figure it out yourself” era of model integration is ending.
A copyright policy with actual mechanics. Not a paragraph of intent — a working policy that respects technological protection measures, excludes known piracy sources, honors robots.txt and machine-readable opt-out signals, and gives rightsholders a contact and a complaint path. This is the provision the training-data lawsuits will hang on.
A public training-data summary. A filled-in AI Office template summarizing what the model was trained on. Not the dataset, but enough of a summary that the black box gets a label.
None of this is exotic to anyone who has run a mature engineering shop. What’s new is that it’s mandatory, enforceable, and discoverable.
Systemic Risk: When the Regulation Starts Speaking My Language
Here is the section that made me want to write this post. A subset of models, those with “high-impact capabilities” (presumed once training compute crosses 10^25 FLOP ) plus any others the Commission designates, are classified as carrying systemic risk . And the obligations that attach to them could have come straight off one of my own engagement checklists:
A safety and security framework (the timing specifics here come from the Code of Practice’s safety-and-security chapter), stood up within weeks of notification and finalized before the model ships.
Model evaluations and adversarial testing — the Act says red-teaming out loud. State-of-the-art evaluation of the model’s dangerous capabilities is now a legal duty, not a nice-to-have your safety team fights for budget on.
Systemic-risk assessment and mitigation across the lifecycle: filtering, monitoring, input/output controls.
Serious-incident reporting to the AI Office and national authorities on staggered timelines — an actual incident-response obligation for models.
Cybersecurity protection of the model and its physical infrastructure — i.e., protect the weights. Weight exfiltration is now a compliance failure, not just an embarrassing headline.
Ten-year documentation retention, and a safety-and-security report including external evaluators’ findings.
A fair note on sourcing: the Act itself sets these duties at the level of principle (Article 55); several of the concrete specifics above — the framework’s timing, the exact shape of the safety report — come from the Code of Practice’s safety-and-security chapter, which is the paved road for demonstrating you met the statutory bar. The duty is law; some of the detail is the Code.
Read that list and then reread what I wrote after the Hugging Face / OpenAI model-evaluation incident : most defenders have zero telemetry at the model and agent layer, and the tooling to run a full intrusion at machine speed is now something you can trigger by accident during your own safety testing. The AI Act just made the telemetry, the evaluation, and the incident reporting for that exact layer a regulated obligation for systemic-risk models. I don’t love every line of this Act, but I’m not going to pretend that mandating adversarial testing and weight security for the most capable models on Earth is the part worth complaining about. It’s the part I’ve been asking for.
There’s also a subtle security-economics point. Adversarial testing is only as good as the adversary. A regulation that requires red-teaming without defining rigor invites the checkbox version — a friendly internal team, a weekend, a green report. The labs that treat this as real offensive work (external, adaptive, incentivized to actually break the model) will produce safety evidence that means something; the ones that treat it as a compliance artifact will produce a document that looks great right up until someone reads it in an investiga

[truncated]
