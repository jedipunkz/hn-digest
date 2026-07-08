---
source: "https://www.resilientcyber.io/p/how-america-talked-itself-into-chinese"
hn_url: "https://news.ycombinator.com/item?id=48832505"
title: "America Talked Itself into Chinese Open Source AI"
article_title: "How America Talked Itself Into Chinese Open Source AI"
author: "smurda"
captured_at: "2026-07-08T15:11:34Z"
capture_tool: "hn-digest"
hn_id: 48832505
score: 2
comments: 0
posted_at: "2026-07-08T14:32:46Z"
tags:
  - hacker-news
  - translated
---

# America Talked Itself into Chinese Open Source AI

- HN: [48832505](https://news.ycombinator.com/item?id=48832505)
- Source: [www.resilientcyber.io](https://www.resilientcyber.io/p/how-america-talked-itself-into-chinese)
- Score: 2
- Comments: 0
- Posted: 2026-07-08T14:32:46Z

## Translation

タイトル: アメリカは中国のオープンソース AI について自ら語る
記事のタイトル: アメリカはどのようにして中国のオープンソース AI に乗り出したのか
説明: ミトスのエピソード、無差別級の急増、そしてなぜ制限が攻撃者よりも防御者を傷つけるのか

記事本文:
アメリカはいかにして中国のオープンソースAIに乗り出したか
回復力のあるサイバー
アメリカはどのようにして中国のオープンソース AI に参入したか
ミトスのエピソード、無差別級の急増、そしてなぜ制限が攻撃者よりも防御者にダメージを与えるのか
Chris Hughes 2026 年 7 月 8 日 シェア 過去 2 年間のほとんどにおいて、オープン AI とクローズ AI に関する議論は主に経済的およびイデオロギー的なものでした。
非公開の研究所は、フロンティアの能力は危険すぎて自由に配布できないと主張した。オープン支持者は、透明性、コスト、管理は、わずかな能力のリードよりも価値があると主張しました。日々のセキュリティ業務においては、閉鎖的フロンティアモデルの方が単純に優れており、オープンモデルは一歩か二歩遅れていたため、実務家は主にその議論を傍観することができた。
その差は縮まり、議論は学術的なものではなくなった。これは現在、CISO、セキュリティ エンジニア、プラットフォーム チームが現在行っているアーキテクチャ上の決定であり、多くの場合、どちらの側にもセキュリティへの影響を推論するための明確な枠組みがありません。
私たちがどのようにして 2026 年半ばにここにたどり着いたのかを説明し、その後、私たちが全体として過小評価していると思われる部分にほとんどの時間を費やしたいと思います。これが、オープンかクローズドの選択が、防御側と私たち全員が依存しているソフトウェア サプライ チェーンにとって実際に意味するものです。
Resilient Cyber​​ ニュースレターをお読みいただきありがとうございます。無料で購読し、20,000 人以上の読者に加わって、AppSec、リーダーシップ、AI、サプライ チェーンなどのサイバーセキュリティに関する最新ニュースを毎週更新してください。
米国の政策がいかにして自らを追い詰めたか
目標と行動が重要な点で乖離しているため、戦略的目標から始めます。 2025 年 7 月、政府はアメリカ産の輸出を促進する大統領令 14320 を発行しました。

AI テクノロジー スタック。
明確な野心は、米国の AI ハードウェア、モデル、標準、ガバナンスが、世界中の同盟国やパートナーにとってのデフォルトのスタックになるということでした。米国が中国との AI 競争に勝ってほしければ、米国のスタックを輸出し、世界がそれを基にして構築することが首尾一貫した方法だった。
2026年6月、アンスロピックは『フェイブル5』と『ミトス5』を発売したが、その約3日後、米国政府は国家安全保障と輸出管理当局を理由に、同社に対し、両方への外国からのアクセスを遮断するよう命令した。さまざまな大手メディアの報道は、この命令を信頼できるパートナーからのジェイルブレイク報告と結びつけており、報道ではアマゾンを指摘しており、モデルが無制限のサイバーツールに転用される可能性があることを示唆していた。
アントロピック氏は脱獄が実際にどれだけ過酷なものであったかに異議を唱え、不透明なプロセスを批判した。約 18 日後、制限は解除され、モデルはオンラインに戻りました。
禁止自体についてはすでに「サイバーセキュリティのフレンドリーファイア問題」で取り上げたので、ここではすべてを再議論するつもりはありません。簡単に言えば、米国のフロンティアモデルをゲートして禁止しても、脅威の状況から根本的な機能が削除されるわけではないということです。
これにより、米国はその能力が使用されているのを監視することができなくなり、欧州内外のすべての同盟国に、米国の AI スタック上に構築することが政治的リスクを伴うかどうかを疑問視する具体的な理由が与えられた。それは大統領令 14320 が達成しようとしていたこととは逆です。
輸出行動から数日以内にフランスとオランダが AI の主権を求める声を高め始めれば、輸出戦略自体が不利になります。
人々をオープンウェイトへと駆り立てる静かな力
この禁止は見出しで大きく取り上げられましたが、オープンソース AI への関心が再び高まっているのはそれだけではありません。

その役割を誇張しすぎないように注意したい。いくつかの構造的な力が一度に集中しており、そのほとんどは地政学とは何の関係もありません。
1 つ目はコストであり、実践者が最も直接感じるのはコストです。 2026 年初頭の短期間、大企業内での AI 導入の最も大きな兆候は、トークン消費量の増加でした。
それが大きく逆転した。 TechCrunch などの報道では、財務チームが現在、同じ数字を削減しようとしていることが文書化されています。 Amazon は、AI を使用するためだけに AI を使用すべきではないという内部規定により、トークン消費量によって開発者をランク付けする内部リーダーボードを 2026 年 5 月下旬に閉鎖したと伝えられています。ウーバーは、2026年のAIコーディングツール予算を4カ月で使い果たし、ツール当たりの支出の上限を従業員1人当たり月額1,500ドルに設定したと述べた。
エージェント ワークフローは、標準的なチャットボット インタラクションの 5 倍から 30 倍のトークンを消費するため、従量制のフロンティア API を通じてすべてを実行する経済性は、大規模に運用している組織にとって意味がなくなりました。
Adrian Sanabria 氏はこの点を以前から主張しており、大量の自動分析を実行することが多く、トークンごとのコストが急速に増大するセキュリティの分野では特に困難を伴います。ゲイリー・マーカスが、トークン最大化がトークン最小化に取って代わられると言うとき、それは雰囲気ではなく、予算の限界です。また、AI が引き起こすあらゆるセキュリティ問題への答えは、AI を使用してそれと戦うことであるというテーマに冷や水を浴びせています。
2番目の力は、フロンティアラボ自身が安全の名の下に自社の製品に摩擦を加えていることです。 Anthropic が 2026 年 7 月に Fable 5 を再デプロイしたとき、新しい分類子レイヤーを使用して再デプロイしました。
Anthropic 自身の「Redeploying Claude Fable 5」の投稿によると、新しいサイバー分類子をトリップさせる Fable 5 リクエストは、Claude Opus 4.8 に再ルーティングされます。

ユーザーに通知されたため、作業はより保守的なモデルで完了します。 Anthropic 氏はトレードオフについて正直であり、私もそれを尊重しており、日常的なコーディングやデバッグ中に分類子が無害なリクエストにフラグを立てることが多くなったと指摘しました。
あなたが正規の脆弱性調査やエクスプロイトに隣接するコードのデバッグを行っているセキュリティ エンジニアである場合、安全マージンの分類子によって無害な作業が危険に見えると判断されたため、タスクの途中で別のモデルにダウンシフトされる可能性が現実にあります。これは Anthropic 側の合理的な安全上の決定であり、実践者にとっては真の能力課税です。これはまさに、あなたが完全に制御する自己ホスト型モデルを魅力的に見せる摩擦のようなものです。
私はセキュリティ研究者やリーダーとさまざまなプライベート グループ チャットに参加していますが、正当なセキュリティ作業を格下げする分類子に不満を抱いている人たちでいっぱいです。
3番目の力はすべてを変えるものであり、それはオープンウェイトが追いついたということです。 2026 年 6 月、Z.ai は GLM-5.2 をリリースしました。これは、約 400 億のアクティブ パラメーター、100 万のトークン コンテキスト ウィンドウ、および MIT ライセンスを備えた、約 7,440 億のパラメーターの専門家混合モデルです。
これは、人工分析インテリジェンス指数のオープンウェイト モデルの中で 1 位、全体では 4 位にランクされています。コーディングベンチマークでは、SWE-bench Pro の GPT-5.5 を上回り、FrontierSWE の Claude Opus 4.8 のポイント内に収まり、コストはおよそ 6 分の 1 です。主流のコーディングおよびエンジニアリング作業では、GLM-5.2 は事実上、クローズド フロンティアと同等です。
このギャップが有意義に広がるのは、最も困難な超長期的なタスクにおいてのみです。オープンオプションが同等で、6倍安く、自己ホスト可能で、政府命令によって無効にできない場合、その魅力は明らかです。
Joshua Saxe などの人々は、例として GLM-5.2 を使用しています。

彼の記事「GLM-5.2は神話ではなく、本当の安全保障上の緊急事態である」では、米国によるクローズドソースのフロンティアモデルの禁止がいかに逆効果であるか、そしてそれが攻撃者よりも防御者にどれほどのダメージを与えるかについて述べています。
業界のシグナルも積み重なっています。パランティアのアレックス・カープ氏は2026年7月1日にCNBCに出演し、フロンティア・ラボのトークンベースの価格設定モデルは「完全に間違っている」と述べ、コンピューティング、モデル、データ、アルファを制御したい顧客への答えとしてオープンウェイト・モデルを組み立てた。
同氏はまた、中国が国家安全保障の態勢をシリコンバレーのコンセンサス的な見解に委ねることを本当に望んでいるかどうかを尋ね、中国の動きがどれほど速いかを過小評価しないよう警告した。 Palantir の真の動機については好きに言ってください。しかし、彼の指摘の多くは、AI に関しては多くの人が共有しているものです。以下のビデオが急速に広まりました。
これとは別に、Axios は、Microsoft が Copilot 内の低コストのオプションとして、DeepSeek V4 の微調整された Azure ホスト型バージョン、または別のオープン モデルを検討していると報告しました。
これらの両方に慎重にフラグを立てたいと思います。カープのコメントは記録に残されており、彼の構成は彼自身のものです。 Microsoft のレポートでは、Microsoft が最終段階ではないとしている調査について説明されており、そのモデルはオプションであり、追加の安全策を備えて Azure 内でホストされる予定です。どちらも中国のオープンウェイトを標準化するという決定が定まったものではなく、私はそれらをそのように扱うつもりはありません。ただし、他の多くの西側企業も同様に中国の AI モデルを採用していると報告されているため、Microsoft も同様ではありません。
彼らが示しているのは、「オープンウェイトは愛好家向け」というフレーミングは時代遅れであり、真剣な企業と真剣なバイヤーが積極的にスイッチの価格を設定しているということです。また、MSF が米国政府への最大のソフトウェア供給業者の 1 つであることも興味深い。

T 社は DeepSeek の使用を検討していますが、政府が承認した製品には使用しない可能性が高いです。
政治的なシグナルはおそらくより顕著だ。米国の元AI皇帝デビッド・サックスは、私が週末に聞いたAll-Inポッドキャストの最近のエピソードを利用して、オープンソースAIを直接主張し、オープンソースは中国との競争における米国にとって不利な要素ではなく、最も強力なカードの1つであると主張した。
同氏は、企業が実際に望んでいるのはコンピューティング、モデル、データの制御であるというカープ氏の指摘に同調し、フロンティア研究所が自社の顧客と競合する垂直アプリケーションを立ち上げようとしているときに、独自の知識をフロンティア研究所に提供することは危険であるとの懸念を提起した。
それがどこまで研究所の公正な特徴付けであるかについて議論することはできるが、Anthropic と OpenAI は強く反発するだろう。この議論にとって重要なことは、米国の AI 政策の中心にいた人物が、オープンソースを制限されるべきものではなく、戦略的資産として公に組み立てているということです。これは政治の重心がどこに位置するかにおける意味のある変化だ。
Michael Spencer のブログ「The Token Apocalypse」や This Week in AI HSM @ This Week in AI ブログ「The State of AI: The US is Losing its AI Monopoly」など、他のいくつかのブログもこの変化に関する素晴らしい記事を発表しています。
通常の直感を覆す安全保障論
ここで、セキュリティ担当者は直観をアップデートする必要があります。再帰的に考えると、クローズドでゲートされたモデルの方が、アクセスが制御されているため安全であるということになります。 Joshua Saxe は、Mythos ではなく GLM-5.2 が本当の安全保障上の緊急事態であるという、より鋭い主張を自身の記事で行っています。そして、私は彼のメカニズムについては正しいと思います。
攻撃者がクローズド モデルを使用する場合、攻撃者はプロバイダーのインフラストラクチャ上で動作します。

、信頼と安全のチームが不正行為を監視している監視下にあります。これは仮説上の利益ではなく、Anthropic が初めて文書化された AI を利用して組織された大規模サイバースパイ活動をどのように捕捉したかそのものです。同社はこの活動を 2025 年 11 月に公開し、中国国家支援グループによるものであると高い信頼性を持っています。
攻撃者はクロード コードをジェイルブレイクし、およそ 30 のターゲットを実行し、キャンペーンの推定 80 ～ 90% をモデルに自律的に処理させ、人間が介入するのはほんの一握りの決定点のみでした。この操作は監視された API ゲート型のインフラストラクチャ上で実行されていたため、Anthropic はそれを捕捉し、マッピングしてアカウントを禁止し、被害者に通知しました。ラボの枠組みの一部をマーケティングとして割り引いたとしても、可視性の点は重要なので、見えないアカウントを禁止することはできません。これは、企業のセキュリティ リーダーが何年もシャドウ IT/SaaS、そして現在は AI と格闘してきた経験から、すでに直感的に知っていることです。
次に、自己ホスト型オープンウェイト モデルで同じキャンペーンを実行します。
使用ログ、信頼性と安全性のチーム、禁止するアカウント、トリガーする検出はありません。閉じたモデルを制限しても、この機能は失われません。
それは暗闇の中で、誰も見ていない環境の H200 のラックに移動します。それでカレー

[切り捨てられた]

## Original Extract

The Mythos episode, the open-weights surge, and why restriction hurts defenders more than attackers

How America Talked Itself Into Chinese Open Source AI
Resilient Cyber
Subscribe Sign in How America Talked Itself Into Chinese Open Source AI
The Mythos episode, the open-weights surge, and why restriction hurts defenders more than attackers
Chris Hughes Jul 08, 2026 Share For most of the last two years, the debate about open versus closed AI was mostly an economic and ideological one.
Closed labs argued that frontier capabilities were too dangerous to hand out freely. Open advocates argued that transparency, cost, and control were worth more than a marginal capability lead. Practitioners could mostly watch that argument from the sidelines, because in day-to-day security work the closed frontier models were simply better and the open models were a step or two behind.
That gap has closed, and the debate has stopped being academic. It is now a live architectural decision that CISOs, security engineers, and platform teams are making right now, often without a clear framework for reasoning about the security implications on either side.
I want to walk through how we got here in the middle of 2026, and then spend most of my time on the part I think we are collectively underweighting, which is what the open versus closed choice actually means for defenders and for the software supply chain we all depend on.
Thanks for reading the Resilient Cyber Newsletter! Subscribe for FREE and join 20,000+ readers to receive weekly updates with the latest news across AppSec, Leadership, AI, Supply Chain, and more for Cybersecurity.
How U.S. policy talked itself into a corner
Start with the strategic goal, because the goal and the actions have drifted apart in a way that matters. In July 2025 the administration issued Executive Order 14320, Promoting the Export of the American AI Technology Stack .
The explicit ambition was that American AI hardware, models, standards, and governance would become the default stack for allies and partners around the world. If you wanted the U.S. to win the AI competition with China, exporting the American stack and getting the world to build on it was a coherent way to do it.
In June 2026 Anthropic launched Fable 5 and Mythos 5, and roughly three days later the U.S. government ordered the company to cut off foreign access to both, citing national security and export-control authority. Reporting from various leading media outlets tied the order to a jailbreak report from a trusted partner, with reporting pointing to Amazon, that suggested the models could be turned into unrestricted cyber tools.
Anthropic disputed how severe the jailbreak actually was and criticized the opaque process. About eighteen days later the restrictions were lifted and the models came back online.
I already worked through the ban itself in Cybersecurity’s Friendly Fire Problem , so I will not relitigate the whole thing here. The short version is that gating and banning a U.S. frontier model did not remove the underlying capability from the threat landscape.
It removed the U.S. ability to watch that capability being used, and it handed every ally in Europe and beyond a concrete reason to question whether building on an American AI stack carries political risk. That is the opposite of what Executive Order 14320 set out to accomplish.
When France and the Netherlands start amplifying calls for AI sovereignty within days of your export action, the export strategy is working against itself.
The quieter forces pushing people toward open weights
The ban gets the headlines, but it is not the only thing reviving interest in open source AI, and I want to be careful not to overstate its role. Several structural forces are converging at once, and most of them have nothing to do with geopolitics.
The first is cost, and it is the one practitioners feel most directly. For a brief window in early 2026 the loudest signal of AI adoption inside big companies was token consumption going up.
That has reversed hard. Coverage from TechCrunch and others documented finance teams now trying to drive that same number down. Amazon reportedly shut down an internal leaderboard that ranked developers by token consumption in late May 2026, with the internal line being that you should not use AI just to use AI. Uber said it burned through its entire 2026 AI coding-tools budget in four months and capped spend at $1,500 per employee per month per tool.
Agentic workflows consume something like 5x-30x the tokens of a standard chatbot interaction, so the economics of running everything through a metered frontier API stopped making sense for organizations operating at scale.
Adrian Sanabria has been making this point for a while, and it lands especially hard in security, where you are often running high-volume automated analysis and per-token costs compound quickly. When Gary Marcus says tokenmaxxing is giving way to tokenminimizing, that is not a vibe, it is a budget line. It also pours some cold water on the theme that the answer to every security problem AI introduces is to use AI to fight it.
The second force is the frontier labs themselves adding friction to their own products in the name of safety. When Anthropic redeployed Fable 5 in July 2026, it did so with a new classifier layer.
Per Anthropic’s own Redeploying Claude Fable 5 post, a Fable 5 request that trips the new cyber classifier is rerouted to Claude Opus 4.8, with the user notified, so the work still completes on a more conservative model. Anthropic was honest about the tradeoff, which I respect, noting that the classifier flags benign requests more often during routine coding and debugging.
If you are a security engineer doing legitimate vulnerability research or debugging exploit-adjacent code, you now have a real chance of being downshifted to a different model mid-task because a safety-margin classifier decided your benign work looked risky. That is a rational safety decision on Anthropic’s part and a genuine capability tax on the practitioner. It is exactly the kind of friction that makes a self-hosted model you fully control look attractive.
I’m in various private group chats with security researchers and leaders and it is full of folks frustrated with the classifiers downgrading their legitimate security work .
The third force is the one that changes everything, which is that open weights have caught up. In June 2026, Z.ai released GLM-5.2 , a roughly 744-billion-parameter mixture-of-experts model with around 40 billion active parameters, a one-million-token context window, and an MIT license.
It ranks first among open-weights models on the Artificial Analysis Intelligence Index and fourth overall. On coding benchmarks it beats GPT-5.5 on SWE-bench Pro and lands within a point of Claude Opus 4.8 on FrontierSWE, and it does it at roughly one-sixth the cost. For mainstream coding and engineering work, GLM-5.2 is effectively at parity with the closed frontier.
The gap only opens meaningfully on the hardest ultra-long-horizon tasks. When the open option is at parity, six times cheaper, self-hostable, and cannot be turned off by a government order, the pull is obvious.
Folks such as Joshua Saxe have used GLM-5.2 as an example of how counterproductive the U.S. ban on closed source frontier models are, and how it hurts defenders more than attackers in his piece “ GLM-5.2 not Mythos, is the real security emergency ”.
The industry signals are stacking up as well. Alex Karp at Palantir went on CNBC on July 1, 2026 and called the token-based pricing model of the frontier labs “completely wrong, ” framing open-weight models as the answer for customers who want control over their compute, their models, their data, and their alpha.
He also asked whether the country really wants to outsource its national security posture to the consensus view of Silicon Valley, and he warned against underestimating how fast China is moving. Say what you want about Palantir’s true motives, but many of his points are ones a lot of folks do share when it comes to AI. The below video has now gone viral:
Separately, Axios reported that Microsoft is exploring a fine-tuned, Azure-hosted version of DeepSeek V4, or another open model, as a lower-cost option inside Copilot.
I want to flag both of these carefully. Karp’s comments are on the record and his framing is his own. The Microsoft reporting describes an exploration that Microsoft says is not final, with the model that would be optional and hosted inside Azure with added safeguards. Neither is a settled decision to standardize on Chinese open weights, and I would not treat them as such. Microsoft isn’t along either though, as it is reported many other Western companies are adopting Chinese AI models as well:
What they do show is that the “open weights are for hobbyists” framing is dead, and serious enterprises and serious buyers are actively pricing the switch. It is also interesting that one of the largest software suppliers to the U.S. Government in MSFT is looking at potentially using DeepSeek, albeit they likely wouldn’t use it for their Government approved offerings.
The political signal is arguably more striking. Former U.S. AI Czar David Sacks used a recent episode of the All-In podcast that I listened to over the weekend to make the case for open source AI directly, arguing that open source is one of America’s strongest cards in the competition with China rather than a liability.
He echoed Karp’s point that what enterprises actually want is control over their compute, models, and data, and he raised the concern that feeding proprietary knowledge to frontier labs is risky when those same labs are launching vertical applications that compete with their own customers.
You can debate how much of that is a fair characterization of the labs, and Anthropic and OpenAI would push back hard. What matters for this discussion is that a figure who sat at the center of U.S. AI policy is now publicly framing open source as a strategic asset, not the thing to be restricted. That is a meaningful shift in where the political center of gravity sits.
Several others have had great pieces on this shift too, such as Michael Spencer ’s blog titled “ The Token Apocalypse ”, or This Week in AI HSM @ This Week in AI blog “ The State of AI: The US is Losing its AI Monopoly ”.
The security argument that flips the usual intuition
Here is where security practitioners have to update their intuition. The reflexive take is that closed, gated models are safer because access is controlled. Joshua Saxe made the sharper argument in his piece that GLM-5.2, not Mythos, is the real security emergency , and I think he is right about the mechanism.
When an attacker uses a closed model, they are operating on the provider’s infrastructure, under monitoring, with trust and safety teams watching for abuse. That is not a hypothetical benefit, it is exactly how Anthropic caught the first documented large-scale AI-orchestrated cyber espionage campaign , which it disclosed in November 2025 and attributed with high confidence to a Chinese state-sponsored group.
The attackers jailbroke Claude Code, ran roughly thirty targets, and let the model handle an estimated 80 to 90 percent of the campaign autonomously, with humans stepping in at only a handful of decision points. Anthropic caught it, mapped it, banned the accounts, and notified victims, because the operation ran on monitored, API-gated infrastructure. Even if you discount some of the lab’s framing as marketing, the visibility point holds, you cannot ban an account you cannot see. This is something enterprise security leaders already know intuitively after years of wrestling with shadow IT/SaaS and now AI.
Now run the same campaign on a self-hosted open-weights model.
There is no usage log, no trust and safety team, no account to ban, no detection to trigger. The capability does not go away when you restrict the closed model.
It relocates into the dark, onto a rack of H200s in an environment nobody is watching. So the curre

[truncated]
