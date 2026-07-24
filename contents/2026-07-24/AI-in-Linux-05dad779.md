---
source: "https://drewdevault.com/blog/AI-in-Linux/"
hn_url: "https://news.ycombinator.com/item?id=49041851"
title: "AI in Linux"
article_title: "AI in Linux"
author: "zdw"
captured_at: "2026-07-24T21:56:39Z"
capture_tool: "hn-digest"
hn_id: 49041851
score: 3
comments: 0
posted_at: "2026-07-24T21:33:34Z"
tags:
  - hacker-news
  - translated
---

# AI in Linux

- HN: [49041851](https://news.ycombinator.com/item?id=49041851)
- Source: [drewdevault.com](https://drewdevault.com/blog/AI-in-Linux/)
- Score: 3
- Comments: 0
- Posted: 2026-07-24T21:33:34Z

## Translation

タイトル: Linux における AI

記事本文:
Linux における AI ツール (主に LLM) の役割については議論中です。あるいは、Linus Torvalds が Linux カーネル開発における AI の使用を支持するために「足を踏み入れるまで」議論されていました。
Linux カーネル開発で AI が使用される主な方法は 2 つあります。それは、コードの作成とコードのレビューです。この記事の執筆時点では、2025 年 9 月から現在まで、「Assisted-by」タグが付いたカーネル コミットが 1,200 件をわずかに超えています。そのほとんどは、LLM ツールによって作成または支援されたパッチを示しています。
Linux 向け AI の 2 番目の重要な用途には、Sasako と呼ばれる新しいコード レビュー ツールが付属しています。これは、さまざまなサブシステムに対して考慮されたパッチのコード レビューを生成します。 Sashiko は、Linux における AI の限界を押し広げるため、Linux における AI に関する現在の議論に火をつけました。AI に反対する人、または AI の使用を望まない人は、以前はパッチを作成するために AI を使用することを控えるだけで済みましたが、現在では、Linux に貢献したい人は誰でも、自分の作業について AI によって生成されたフィードバックを反復するために Sashiko やその他の同様の AI ツールと対話する必要があるという期待が高まっています。
Linux カーネル コミュニティにおけるこの議論の主な内容の 1 つは、LLM の使用の倫理的考慮に関するものでした。ライナスはこの論拠を完全に封じ、議論を技術的なメリットにしっかりと根付かせ、この問題に関するいかなる政治的議論も拒否した。
カーネル プロジェクトはこれまでも、そしてこれからもテクノロジーに関するものです。
確かに、オープンソースに取り組むという社会的な側面は重要であり、多くの場合、プロジェクトの非常に動機付けとなる部分ですが、結局のところ、それは副次的な利点であり、プロジェクトの目的ではありません。
これはある種の「社会戦士」プロジェクトではありませんし、これまでも、そしてこれからも決してありません。
カーネルコミュニティでは、オープンソースを採用しています。

宗教的な理由ではなく、テクノロジーを活用してください。
この議論は不誠実で偽善的です。 Linux は政治プロジェクトであり、Linus は政治活動家です。 Linux のライセンス付与に GPLv2 の使用を検討してください。技術的な利点から議論することもできます。たとえば、GPLv2 のコピーレフトの性質により、人々、特に営利団体がドライバーやその他の貢献物を Linux カーネルのアップストリームに押し上げています。これは、結果としてカーネルの技術的な卓越性に貢献します。
しかし、これは政治的選択ではなく、政治的行為ではないでしょうか？この選択の目的は、他の人の行動に影響を与え、カーネルの利益を自分の利益よりも優先させることです。そして、Linus は GPLv3 が導入されたとき、政治的理由からこの決定に固執し、ライセンスとその展開方法に反対するときは道徳と倫理から推論し、FSF の暗黙のボイコットを呼びかけました。
Linus と Linux は世界中に多大な権力と影響力を行使しており、それは責任を持って行使されるべきです。ライナスが次のように言ったとき：
Linux はそれらの反 AI プロジェクトの 1 つではありません。もし誰かがそれに問題を抱えているなら、オープンソースにしてフォークすることができます。
それはまったく不誠実だと思います。 Linus は、実際的な目的から見て Linux をフォークできないことを確実に認識しています。これは世界最大のソフトウェア プロジェクトであり、最も資金が豊富なプロジェクトの 1 つでもあります。貢献者間の組織的な知識や、猛烈な変化のペースについていく見込み、さらには上流から独立してカーネルのコードの一部を理解して維持するための時間と資金を持った人々のグループをまとめる見込みさえ、まったく手に負えないのです。ライナスもこれを知っています。これは、

GPL、GPL 専用シンボル、および不安定な内部カーネル ABI: カーネルのフォークを独立して維持するプロセスを可能な限り困難にするため。
人々は、不可能な手段に訴える前に、Linux アップストリームにその動作とポリシーを修正するよう請願するのが当然です。 Linux に取り組むには、社内で政治を実践する必要があります。たとえば、高い技術的卓越性と低い社会的/政治的能力を備えた bcachef をめぐる議論に対する Linus の反応を参照してください。また、Linux とその他の世界の交差点でも政治を実践する必要があります。したがって、仕事の進め方について議論するときは、政治的、道徳的、倫理的な議論を展開する必要があります。自分の主張に合わない場合には政治的および倫理的な考慮事項から議論を逸らし、必要な場合には議論を提起するという、安っぽく弱い議論です。
私は、LLM を利用した Sashiko のコード レビューが多くの優れた洞察を提供すると信じたいと思っています。ただし、このツールの外部性を 1 つだけ取り上げるには、AI 企業が消費者向けハードウェアの価格を吊り上げていることを考慮してください。 AI を活用したコード レビューによってパッチが改善される可能性がありますが、そのパッチは、より良いパッチを享受するために Linux を実行できるハードウェアの価格が高くなる、ますます多くの人々にとってはあまり役に立ちません。
別の角度から見てみると、大気中に何トンの CO₂ が追加されるか、あるいは何リットルの真水供給が中断されるかは、より良いコードレビューのために許容できる代償となるでしょうか?インドでは熱波中の気温が50度を超え、数万人が死亡している。構築された AI は、世界で最も急速に成長しているエネルギー消費者であり、化石燃料を使用して構築されているか、他の産業に依存している化石燃料の代替からグリーン エネルギー需要を引き出しています。同じプロセスで通常の価格設定が行われます

熱波の最中に人々はエアコンに電力を供給するために必要なエネルギーを使い果たし、それを可能にするテクノロジーによって人々は労働市場から追い出され、貧困に陥っています。
これらの外部性は非常に現実的であり、これらの問題を Linus の注意を引こうとしている Linux カーネルの貢献者やメンテナ、そしてその友人や家族を含む多くの人々に影響を与えています。
Linux でこれらのツールを使用すると、目に見えない影響はどのようなものになるでしょうか? Linux は、これらのツールの作成者を正当化するためにプロジェクトの計り知れない影響力の一部を与えており、これらのツールの技術的応用に焦点を絞るという Linus の主張は、彼らへの寛大な贈り物です。彼らは役員室、ロビイストと政府との会合、その他自分たちの利益を促進するあらゆる場面で、必ず彼の支持について言及するだろう。
「それらの利益は何ですか?そしてこの影響力を使って何をしているのでしょうか?」と尋ねるのは当然です。
ドナルド・トランプ大統領の就任式に出席したGoogle CEOのサンダー・ピチャイ氏と他のAI分野で影響力のある商業リーダーらの写真。 Sashiko は主に、Linux カーネル コード レビューのために Google Gemini に依存しています。
この種の質問に関心を持つ、賢くて情熱的な人々がた​​くさんいます。この道徳的計算を行うことを拒否し、他の誰にもそれを許可することを拒否し、これらの問題に関心を持つ才能のある Linux 貢献者を追い払うことにおいて、技術的な卓越性はどこにあるのでしょうか? Linus Torvalds、Linux コミュニティ、そして私たちのすべてのコミュニティは、AI の問題のこれらの側面に正直かつ誠実に対処する勇気と洞察力を備えている必要があります。

## Original Extract

The role of AI tools (LLMs, mainly) in Linux is under discussion, or it was, until Linus Torvalds “ put his foot down ” in support of the use of AI in Linux kernel development.
I can identify two major ways in which AI is used for Linux kernel development: authoring code and reviewing code. There are, at the time of writing, just over 1,200 kernel commits with an “Assisted-by” tag, from September 2025 to the present, most of which indicate patches which were written or assisted by LLM tools.
The second important use of AI for Linux comes with a new code review tool called Sashiko, which generates code reviews for patches considered for various subsystems. Sashiko ignited the current debate on AI in Linux because it pushes the envelope on AI in Linux: people who oppose or do not want to use AI could previously just refrain from using it to write their patches, but now there is a growing expectation that anyone who wants to contribute to Linux will have to interact with Sashiko or other AI tools like it to iterate on AI-generated feedback on their work.
One of the major lines of this discussion in the Linux kernel community has been with respect to the ethical considerations of the use of LLMs. Linus shuts this line of reasoning down entirely, firmly grounding the discussion in technical merits and rejecting any political discourse on the matter:
The kernel project has been and will continue to be about the technology.
Sure, the social angle of working on open source is important and often a very motivating part of the project, but in the end that’s a side benefit, not the point of the project.
This is NOT some kind of “social warrior” project, never has been, and never will be.
In the kernel community we do open source because it results in better technology, not because of religious reasons.
This argumentation is disingenuous and hypocritical. Linux is a political project and Linus is a political actor. Consider the use of the GPLv2 for licensing Linux. One can argue from technical merits – for instance, the copyleft nature of the GPLv2 pushes people, and in particular commercial entities, to upstream their drivers and other contributions into the Linux kernel. This contributes to the technical excellence for the kernel as a result.
But is this not a political choice, and a political act? The purpose of this choice is to influence the behavior of others and to advance the interests of the kernel ahead of their own. And Linus stuck to this decision for political reasons when GPLv3 was introduced, reasoning from morality and ethics when objecting to the license and the manner in which it was deployed, and called for a tacit boycott of the FSF.
Linus, and Linux, wields a tremendous degree of power and influence over the world, and it should be wielded responsibly. When Linus says the following:
Linux is not one of those anti-AI projects, and if somebody has issues with that, they can do the open-source thing and fork it.
I find it completely disingenuous. Linus is surely aware that, for all practical purposes, Linux cannot be forked. It is the world’s largest software project, and one of the most well-funded, too. The institutional knowledge among its contributors, the prospect of keeping up with the blistering pace of change, or even putting together a group of people with the time and funding to understand and maintain even a fraction of the kernel’s code independently of upstream, is, quite simply, intractable. Linus knows, this, too – it’s an explicitly cited reason for decisions like the use of the GPL, GPL-only symbols, and the unstable internal kernel ABI: to make the process of independently maintaining a fork of the kernel as difficult as possible.
People are right to petition Linux upstream to amend its behavior and policies before resorting to the impossible. Working on Linux requires practicing politics, both internally – see for example Linus’ response to the discussions around bcachefs, which had high technical excellence and low social/political competence – and at the intersection of Linux and the rest of the world. Therefore, we must table political, moral, and ethical arguments when we discuss how we go about the work. It’s a cheap, weak argument to direct the discussion away from political and ethical considerations when it wouldn’t serve your point and to table it when it would.
I’m willing to believe that the LLM-powered Sashiko code reviews provide a lot of good insights. However, to address just one externality of this tool, consider that AI companies are driving up the price of consumer hardware. An AI powered code review may improve a patch, but that patch won’t be of much use to the increasingly large cohort of people who are being priced out of the hardware they could run Linux on to enjoy the better patch.
Looking at it from another angle: how many tons of CO₂ added to the atmosphere or liters of fresh water supplies disrupted is a tolerable price for a better code review? Temperatures during heat waves are exceeding 50°C in India, causing tens thousands of deaths. The AI built-out is by far the fastest growing energy consumer in the world , and they’re being built with fossil fuels, or drawing green energy demand away from replacing the fossil fuels depended on by other industries. The same process is pricing regular people out of the energy they need to power their air conditioner during those heat waves and the technology it enables is pushing them out of the labor market and into poverty.
These externalities are very real, and are affecting a lot of people, including Linux kernel contributors and maintainers, and their friends and families, who are trying to bring these problems to Linus’ attention.
And what of the intangible effects of the use of these tools in Linux? Linux is lending some of the project’s immense influence towards legitimizing the makers of these tools, and Linus’ insistence on focusing narrowly on the technical applications of these tools is a generous gift to them. They will be sure to mention his support in boardrooms, meetings between lobbyists and governments, and anywhere else it will advance their interests.
It’s fair to ask: what are those interests, and what are they doing with this influence?
Google CEO Sundar Pichai pictured at Donald Trump’s inauguration, together with other influential commercial leaders in AI. Sashiko primarily depends on Google Gemini for Linux kernel code reviews.
There are a lot of smart, passionate people who care about these kinds of questions. Where is the technical excellence in refusing to do this moral calculus, refusing to allow anyone else to do so, and driving away the talented Linux contributors who care about these problems? Linus Torvalds, the Linux community, and all of our communities should have the courage and insight to address these dimensions of the AI question honestly and in good faith.
