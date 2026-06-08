---
source: "https://willhackett.com/opus-tax/"
hn_url: "https://news.ycombinator.com/item?id=48448797"
title: "Turn a $3M AI bill into $1.9M"
article_title: "Turn a $3m AI bill into $1.9m | Will Hackett"
author: "speckx"
captured_at: "2026-06-08T18:57:03Z"
capture_tool: "hn-digest"
hn_id: 48448797
score: 1
comments: 0
posted_at: "2026-06-08T18:01:46Z"
tags:
  - hacker-news
  - translated
---

# Turn a $3M AI bill into $1.9M

- HN: [48448797](https://news.ycombinator.com/item?id=48448797)
- Source: [willhackett.com](https://willhackett.com/opus-tax/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T18:01:46Z

## Translation

タイトル: 300 万ドルの AI 請求を 190 万ドルに変える
記事のタイトル: 300 万ドルの AI 請求を 190 万ドルに変える |ウィル・ハケット
説明: ほとんどのチームは、気付かないうちに AI に 2 回払いすぎています。すべてをデフォルトのフラッグシップ モデルで実行し、可能な従量制プランでトークンを購入しています。

記事本文:
300 万ドルの AI 請求を 190 万ドルに変える |ウィル・ハケット
あなたの購読が確認されました。ありがとう。
← 戻る
2026 年 6 月 8 日 300 万ドルの AI 請求額を 190 万ドルに変える
現在、チームの誰かが最も高価な AI モデルを使用してスライド デッキを編集しています。彼らはそれを選択しませんでした。それは単なるデフォルトです。その目に見えない選択を 1 日に数千回繰り返すと、AI の請求書はすぐに給与計算に似てきます。
2 つのことがこの数字を膨らませています。デフォルトのモデルはこの目的には適していません。フラッグシップ価格で仕事をするなら、もっと安いモデルが最適です。そして、そのタスクは請求書には表示されず、どのプロジェクトやどのモデルで行われたのかを知る方法のない一括金額です。
Flowstate はリクエスト パス内に存在し、両方のリークを閉じます。各プロンプトをタスクが実際に必要とするモデルにルーティングし、すべてのドルをその作業に支払った作業に結び付けます。出荷額がこれより少ない人はいません。300 万ドルかかった同じ作品の価格が 190 万ドルになり、そのお金が実際にどの作品を購入したかが初めてわかります。
ソネットの作品に対してオーパスの価格を支払っていることになります
モデルを選ぶ人はほとんどいません。ボックスのロード時に選択されたものがすべて使用され、デフォルトはフラッグシップ、つまり提供されている中で最も高価なモデルです。これは、本当に難しい問題と純粋な無駄を 1 行のメールで扱うための正しい呼びかけです。デフォルトのチャット ウィンドウに必要なコストの 5 倍のコストがかかることをマーケティング担当者が知ることは期待できません。価格は画面に表示されず、ベンダーには価格を表示する動機がありません。
だから、彼らにそれを学ばせないでください。タスクは入力者ではなくモデルを選択する必要があり、その決定は誰かの頭の中にあるのではなく、リクエスト層に属します。要約や再フォーマットは Haiku に、日常のコーディングと草稿は Sonnet に、真に厳しい推論は Opus に割り当てられます。単一のジョブを分割して、Opus で計画し、Sonnet で実行することもできます。

必要なステップについては高価な考えを持っています。その人は、一日中何をしていても、同じプロンプトを入力し、同じ答えを受け取ります。請求書はただ小さいです。これはクロード コードだけではありません。営業、運用、マーケティング担当者が開くすべてのチャットの前にも同じデフォルトが表示されます。
どのくらい小さいですか？ Ding らの Hybrid LLM のような査読済みの研究では、品質を目に見えるほど低下させることなく、高価なモデルへの通話を最大 40% 削減できることが示されています 1 。これはモデルの組み合わせに関する単なる算術演算であり、合法的に実行するあらゆるデプロイメントで機能します。
これは、使用するにつれて増大するレバーです。チームが AI に依存するほど、間違ったモデルのデフォルトによるコストが増加し、手を戻すルーティングも増えます。下の計算機では、太線と緑色の線の間のギャップです。
1日目です。エンジニアが会社に入社し、Enterprise Claude アカウントを渡され、最初の 5 つのプロンプトで 145 ドルを消費します。定額プランでは、その使用量は 1 週間にわたって延長されるでしょう。従量制のエンタープライズ プランでは、昼食前に終わってしまいます。人事担当者はすでに答えられない質問をしており、月給 5,000 ドルを「私の給料よりも高い」と計算しています。使用量ページで制限を表示する必要がある場合、「 Unlimited 」という 1 つの単語が表示されます。これは r/ClaudeCode からの実際の投稿であり、単一のスクリーンショットでの 2 番目のリークです。
最初のリークは誰も選ばなかったモデルだった。もう一つは、誰も見ていないメーターです。今年の時点で、Enterprise は、チームがチャット、Claude Code、および Cowork で費やすすべてのトークンに対して、シートに加えて標準 API レートで課金します 2 。 (チームは代わりに手当が含まれたフラットシートを維持します。) 従量制料金は、ライトチームにとっては安いですが、規模が大きくなると離れていきます。また、区別されていない 1 つの請求書として着地するため、財務担当者が報告するまで誰もその高騰に気づきません。

旗を立てます。目に見えないものをルーティングすることはできません。また、比較したことのない 2 つのデプロイメントのどちらかを選択することもできません。したがって、それらを比較してください。
ドアを選んでください。チームの規模を調整します。用途をドラッグします。
選択したデプロイメント (太字)、Flowstate ルーティングを使用した同じデプロイメント、および比較のための他のドアの年間 AI 支出。緑の線までのギャップが配線によってあなたを救います。最安の破線との差がドアの価格となります。
$0 $2.63m $5.25m $7.88m $10.50m 0M 250M 500M 750M 1000M チームメンバーあたりのトークン / 月 クロード for Enterprise — 現在 クロード for Enterprise + Flowstate クロード for Teams (プレミアム) 導入戦略 クロード for Enterprise クロード for Teams (プレミアム) チームメンバー (シート) トークン / チームメンバー / 月375 M ▶ これがどのように計算されるか、およびこれが単なる例である理由
ドラッグすると何が起こるか見てみましょう。使用率が低い場合、2 つのドアに違いはほとんどなく、実際には Enterprise の方が安価です。そのため、ライト チームにとってはこれは重要ではありません。使用量を押し上げると、従量線が暴走します。ルーティングは直線で3分の1から半分を引き戻し、トップエンドではフラットなチームシートに移動するだけでも勝ち始めます。しかし、どちらの行動も取れるのは、請求書を比較できるほど明確に理解できる場合のみであり、ほとんどのチームはそれができません。
実際に元が取れたのはどのプロジェクトか
ルーティングにより、タスクごとに支払う金額が決まります。さらに難しい質問は、それで何を買ったかということですが、それは請求書からはまったく読み取ることができません。人々が議論するのはコストの半分だけです。アトリビューションは、静かにコストがかかる半分です。
誰かが今月 Opus に 300 ドルを費やした場合、問題はどのモデルかではなく、どのプロジェクトかです。それに答えられない場合は、すべてのドルが同じ未差別の OpEx バケットに保管され、費やされた瞬間に費用処理されることになります。財務省は Anthropic 社からの告発を検討している

数字は、個人や作品に結び付けることができないため、それを扱うことはできず、その成長を見守ることしかできません。これはコストセンターのない第二の給与計算です。
文脈のない法案は単なる法案、数字が上がっただけです。コンテキストがあれば地図に変わります。新しい請求フローを構築しているチームは月に 40,000 ドルのモデル時間を消費しているのに対し、誰もサインオフしていない実験では 60,000 ドルを消費していることがわかります。どの機能の出荷コストが収益よりも高いのか、そしてどの安価な機能がひっそりとロードマップを実行しているのかがわかります。それはコスト削減ではありません。それは、自分のレバレッジがどこにあるのか、どの仕事に栄養を与え、どの仕事を枯渇させるのかを知ることです。支出の帰属は金融機関が恐れる数字ではなくなり、実際に価値が生み出されている場所を最も正確に読み取ることができるようになります。
そして、それは報告だけでなく会計も変えます。新しいソフトウェアの構築にかかる AI 支出は、IAS 38 または ASC 350-40 3 に基づく従来のソフトウェア開発と同様に、資産化して耐用年数にわたって償却できます。妨げとなるのは会計規則ではありませんでした。それは帰属の欠如でした。帰属できないものを大文字にすることはできず、プロバイダーの請求書には何も帰属しません。 Flowstate はすべての呼び出しを人、プロジェクト、モデル、コスト クラスに結び付けるため、実際の価値を構築する作業が OpEx に隠れることがなくなります。
そして、あなたの仕事がより多くの資格を得るほど、この問題は大きくなります。開発努力の 70% が真に新製品の構築に費やされている場合 (そして多くのチームがそうです)、アトリビューションにより AI 支出の大部分が今四半期の損益から貸借対照表に移され、ソフトウェアが稼いだ年数にわたって償却されます。ハウスキーピングではない 7 桁の AI 請求について。それは、今マージンに達するか、後で取り戻せる資産かの違いです。 (特定のプロジェクトが適格であるかどうかは審査員によって決まります)

財務および監査チーム向けであり、ブログ投稿ではありません。)
Flowstate はインテリジェントなプロキシです。Zscaler を思い浮かべてください。ただし、AI トラフィック用です。当社はアカウントをプールしたり、契約を保持したりしません。あなたは自分のキーを保持し、チームが使用するすべてのプロバイダーとの独自の契約を保持します。私たちはリクエスト パス内にいて、通過する各呼び出しに対して 3 つのことを行います。タスクが実際に必要とするモデルにルーティングし、残すべきではないもの (ソース コード、離れてはならない場所に向かう顧客 PII) を検査し、人、プロジェクト、コスト クラスに対して記録します。それが、Enterprise がプレミアムを請求する、またはプレミアムを請求しない、そして誰にも契約を渡すことなく、可視化することです。
当社は代理店であり、アカウントプールではないため、プロバイダーの条件に従うかどうかは、暗闇ではなく目の前で全体像を見て下されたお客様の決定となります。各導入に実際にかかるコストを確認し、支出を削減し、どの程度のリスクを負うかに応じてチームごとのタップを増減することができます。上記の 2 つのリークは、同じマシンが 2 つのジョブを実行しています。つまり、各リクエストを適切なモデルに送信し、選択したデプロイメントを管理できるほど読みやすくします。
はっきり言って、いくつかの注意点があります。 Flowstate により、デプロイメントが監視可能になり、制御可能になります。契約が書き換えられるわけではありません。 BAA、データ常駐、または契約上のトレーニングなし条項が必要な場合は、そこがエンタープライズへの入り口であり、そこで私たちの仕事はルーティングと台帳であり、従量制請求書がお客様から逃げないようにすることです。そして、すべては大量の使用量の話です。使用量を引き下げた瞬間に計算機が示すように、ライトチームの場合、従量制請求額は決して元が取れるところには近づきません。
何年もの間、この取引は二分法で行われてきました。つまり、人々は目の前にあるモデルに手を伸ばせるというものでした。

d 請求書を食べるか、すべてをロックして、すべてのプロンプトを手動で監視します。
これは、請求額を無駄にするか、使用制限でチームを停止させるかの二者択一であってはなりません。タスクをルーティングすると、Sonnet の作業に対して Opus 価格を支払う必要がなくなります。支出の原因を特定すれば、AI は差別化されていないマージンヒットではなくなります。コントロールを提供するために中間にプロキシが必要なだけです。 4
Ding et al.、『Hybrid LLM: Cost-Efficient and Quality-Aware Query Routing』、ICLR 2024 では、応答品質の低下がなく、大規模モデルへの呼び出しが最大 40% 減少すると報告しています。 Ong らの『RouteLLM: Learning to Route LLMs with Preference Data』では、品質を損なうことなく、ベンチマークの一部で 2 倍を超えるコスト削減が報告されています。ベンダー ルーターのアドバタイズ量は高くなります (40 ～ 70%)。査読済みの図をモデルにしました。 arxiv.org/abs/2404.14618 · arxiv.org/abs/2406.18665 ↩
2026 年にエンタープライズ シート料金をトークン使用量から人為的に切り離す: Claude ヘルプ センターによれば、「使用量はシート料金に含まれていません…チームがチャット、Claude Code、または Cowork で使用するすべてのトークンは、シート料金に加えて標準 API レートで請求されます。」シートごとのセッション制限はレベルによって異なり、Team Premium シートは Pro のセッションごとの制限の約 6.25 倍です。公開されている基本シートは 20 ドル (チーム スタンダードおよびエンタープライズ) と 100 ドル (チーム プレミアム) です。真のエンタープライズ価格は販売交渉によって決定されます。 support.claude.com/en/articles/9797531 · support.claude.com/en/articles/9266767 · claude.com/pricing ↩
IAS 38 無形固定資産 (開発段階) および ASC 350-40 (社内使用ソフトウェア) は、開発コストを費用計上するのではなく資産計上できる時期を規定します。適格性は判断と監査の問題であり、主張ではありません。ここにあるのは会計上のアドバイスではありません。 ↩
私は Flowstate を共同設立したので、明らかな利益相反の開示はありません

が適用されます。上記のルーティングの数値は、顧客アカウントからではなく、公開価格と引用された調査に基づいてモデル化されています。 ↩
AI に関するムーアの法則は正式に消滅した
エージェント コーディングはエンジニアリングの労働力を変えていますが、それはあなたの考え方ではありません
2,000 万ポンドのエンジニアリング支出をどう説明しますか?
メールアドレス 購読 このページで あなたはソネットの作品に対してオーパスの価格を払っています
実際に元が取れたのはどのプロジェクトか

## Original Extract

Most teams overpay for AI twice without noticing: they run everything through the default flagship model, and they buy the tokens on a metered plan they can

Turn a $3m AI bill into $1.9m | Will Hackett
Your subscription has been confirmed. Thank you.
← Back
8 June 2026 Turn a $3m AI bill into $1.9m
Right now, someone on your team is using the most expensive AI model to edit a slide deck. They didn’t choose it; it’s just the default. Repeat that invisible choice a few thousand times a day, and your AI bill quickly starts to resemble payroll.
Two things are inflating that number. The default model is wrong for the task: flagship prices for work a cheaper one would nail. And the task is invisible on the invoice, a single lump sum with no way to tell which project or which model it went on.
Flowstate sits in the request path to close both leaks. We route each prompt to the model the task actually needs, and tie every dollar to the work it paid for. Nobody ships less: the same output that ran you $3m now costs $1.9m, and for the first time you can see which work the money actually bought.
You’re paying Opus prices for Sonnet work
Almost nobody picks a model. They use whatever’s selected when the box loads, and the default is the flagship: the most expensive model on offer. That’s the right call for a genuinely hard problem and pure waste on a one-line email. You can’t expect a marketer to know their default chat window costs five times more than necessary. The price isn’t on the screen, and the vendor has no incentive to show it.
So don’t make them learn it. The task should pick the model, not the person typing, and that decision belongs at the request layer rather than in anyone’s head. A summary or a reformat goes to Haiku, everyday coding and drafting to Sonnet, the genuinely hard reasoning to Opus. You can even split a single job, planning it in Opus and running the execution on Sonnet, keeping the expensive thinking for the steps that need it. The person, whatever they do all day, types the same prompt and gets the same answer. The bill is just smaller. And this isn’t only Claude Code: the same default sits in front of every chat your sales, ops and marketing people open too.
How much smaller? Peer-reviewed research, like Ding et al.’s Hybrid LLM, shows you can cut calls to the expensive model by up to 40% with no measurable drop in quality 1 . It’s just arithmetic on your model mix, and it works on any deployment you legitimately run.
This is the lever that grows with usage: the harder your team leans on AI, the more a wrong-model default costs you, and the more routing hands back. In the calculator below it’s the gap between your bold line and the green one.
It’s day one. An engineer joins a company, gets handed an Enterprise Claude account, and burns $145 in his first five prompts. On a flat-rate plan that usage would have stretched all week; on a metered Enterprise plan, it’s gone before lunch. HR is already asking questions he can’t answer, and he’s doing the maths on a $5,000 month: “more than my salary.” Where the usage page should show a limit, it shows one word: Unlimited . That’s a real post from r/ClaudeCode , and it’s the second leak in a single screenshot.
The first leak was the model nobody chose. This is the other: the meter nobody’s watching. As of this year, Enterprise charges for every token your team spends in chat, Claude Code and Cowork, at standard API rates on top of the seat 2 . (Teams keeps a flat seat with an included allowance instead.) Metered pricing is cheap for a light team but runs away from you at scale, and because it lands as one undifferentiated invoice, nobody catches the spike until finance raises a flag. You can’t route what you can’t see, and you can’t choose between two deployments you’ve never compared. So compare them:
Pick your door. Size your team. Drag the usage.
Annual AI spend for your chosen deployment (bold), the same deployment with Flowstate routing, and the other doors for comparison. The gap to the green line is what routing saves you; the gap to the cheapest dashed line is what the door costs you.
$0 $2.63m $5.25m $7.88m $10.50m 0 M 250 M 500 M 750 M 1000 M tokens per team member / month Claude for Enterprise — today Claude for Enterprise + Flowstate Claude for Teams (Premium) Deployment strategy Claude for Enterprise Claude for Teams (Premium) Team members (seats) Tokens / team member / month 375 M ▶ How this is calculated & why it's only an example
Watch what happens as you drag it. At low usage the two doors barely differ, and Enterprise is actually the cheaper one, which is why none of this matters for a light team. Push the usage up and the metered line runs away. Routing pulls a third to a half straight back off it, and at the top end even moving to a flat Teams seat starts to win. But you can only make either move once you can see the bill clearly enough to compare, and most teams can’t.
Which projects actually paid for themselves
Routing fixes what you pay per task. The harder question is what you bought with it, and you can’t read that off the invoice at all. Cost is only the half people argue about. Attribution is the half that quietly costs more.
When someone spends $300 of Opus this month, the question isn’t which model, it’s which project . If you can’t answer that, every dollar lands in the same undifferentiated OpEx bucket and gets expensed the moment it’s spent. Finance sees a charge from Anthropic and a number, can’t tie it to a person or a piece of work, and so can do nothing with it but watch it grow. It’s a second payroll with no cost centres.
A bill without context is just a bill, a number that went up. With context it turns into a map. You can see that the team building the new billing flow is burning $40k of model time a month, while an experiment nobody signed off on is burning $60k. You can see which features cost more to ship than they’ll ever earn back, and which cheap ones are quietly carrying the roadmap. That isn’t cost-cutting; it’s knowing where your leverage is, which work to feed and which to starve. Attributed spend stops being the number finance dreads and becomes the sharpest read you’ve got on where value is actually being made.
And it changes the accounting, not just the reporting. AI spend that goes into building new software can be capitalised and amortised over its useful life, just like traditional software development under IAS 38 or ASC 350-40 3 . The blocker was never the accounting rules; it was the lack of attribution. You can’t capitalise what you can’t attribute, and the provider’s invoice attributes nothing. Flowstate ties every call to a person, a project, a model and a cost class, so the work building real value stops hiding in OpEx.
And the more of your work qualifies, the bigger this gets. If 70% of your development effort is genuinely building new product (and for a lot of teams it is), attribution shifts the bulk of that AI spend off this quarter’s P&L and onto the balance sheet, to be amortised over the years the software earns. On a seven-figure AI bill that isn’t housekeeping; it’s the difference between a margin hit now and an asset you recoup later. (Whether a given project qualifies is a judgement for your finance and audit team, not a blog post.)
Flowstate is an intelligent proxy: think Zscaler, but for AI traffic. We don’t pool accounts and we don’t hold your contracts; you keep your own keys and your own deal with every provider your team uses. We sit in the request path and do three things to each call as it passes through: route it to the model the task actually needs, inspect it for the things that should never leave (source code, customer PII heading somewhere it shouldn’t), and log it against a person, a project and a cost class. That’s the visibility Enterprise charges a premium for, without the premium, and without handing anyone your contracts.
Because we’re a proxy and not an account pool, where you sit on a provider’s terms stays your decision, made with the whole picture in front of you instead of in the dark. You can see what each deployment really costs, route the spend down, and turn the taps up or down per team based on how much risk you’re willing to carry. The two leaks above are the same machine doing two jobs: sending each request to the right model, and making the deployment you’ve chosen legible enough to manage.
A couple of caveats, plainly. Flowstate makes a deployment observable and controllable; it doesn’t rewrite your contract. If you need a BAA, data residency or a contractual no-training clause, that’s the Enterprise door, and there our job is the routing and the ledger: keeping the metered bill from running away from you. And the whole thing is a heavy-usage story: for a light team the metered bill never gets near the point where any of this pays for itself, as the calculator shows the moment you drag the usage down.
For years, the trade looked binary: let people reach for whatever model is in front of them and eat the bill, or lock the whole thing down and police every prompt by hand.
This shouldn’t be a binary choice between eating the bill and grinding your team to a halt with usage limits. Route the task, and you stop paying Opus prices for Sonnet work. Attribute the spend, and AI stops being an undifferentiated margin hit. You just need a proxy in the middle to give you the controls. 4
Ding et al., Hybrid LLM: Cost-Efficient and Quality-Aware Query Routing , ICLR 2024, reports up to 40% fewer calls to the large model with no drop in response quality. Ong et al., RouteLLM: Learning to Route LLMs with Preference Data , reports cost reductions of over 2× on parts of its benchmark without compromising quality. Vendor routers advertise higher (40–70%); I’ve modelled to the peer-reviewed figure. arxiv.org/abs/2404.14618 · arxiv.org/abs/2406.18665 ↩
Anthropic decoupled Enterprise seat fees from token usage in 2026: per the Claude Help Center, “Usage isn’t included in the seat fee… Every token your team uses — in chat, Claude Code, or Cowork — is billed at standard API rates on top of your seat cost.” Per-seat session limits differ by tier, with a Team Premium seat about 6.25× Pro’s per-session limit. Published base seats are $20 (Team Standard and Enterprise) and $100 (Team Premium); true Enterprise pricing is sales-negotiated. support.claude.com/en/articles/9797531 · support.claude.com/en/articles/9266767 · claude.com/pricing ↩
IAS 38 Intangible Assets (development phase) and ASC 350-40 (internal-use software) govern when development cost may be capitalised rather than expensed. Eligibility is a matter of judgement and audit, not assertion; nothing here is accounting advice. ↩
I co-founded Flowstate , so the obvious conflict-of-interest disclosure applies. The routing figures above are modelled from public pricing and the cited research, not from a customer account. ↩
Moore's Law for AI is officially dead
Agentic coding is changing the engineering workforce—just not how you think
How do you explain the £20M engineering spend?
Email address Subscribe On this page You’re paying Opus prices for Sonnet work
Which projects actually paid for themselves
