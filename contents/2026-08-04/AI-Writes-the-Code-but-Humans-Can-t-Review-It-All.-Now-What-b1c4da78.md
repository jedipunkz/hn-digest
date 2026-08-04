---
source: "https://www.sylvainkalache.com/blog/ai-writes-the-code-but-humans-cant-review-it-all"
hn_url: "https://news.ycombinator.com/item?id=49174469"
title: "AI Writes the Code, but Humans Can't Review It All. Now What?"
article_title: "AI Writes the Code, But Humans Can't Review It All. Now What? — Sylvain Kalache"
author: "jjtang1"
captured_at: "2026-08-04T20:18:57Z"
capture_tool: "hn-digest"
hn_id: 49174469
score: 1
comments: 0
posted_at: "2026-08-04T20:17:37Z"
tags:
  - hacker-news
  - translated
---

# AI Writes the Code, but Humans Can't Review It All. Now What?

- HN: [49174469](https://news.ycombinator.com/item?id=49174469)
- Source: [www.sylvainkalache.com](https://www.sylvainkalache.com/blog/ai-writes-the-code-but-humans-cant-review-it-all)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T20:17:37Z

## Translation

タイトル: AI がコードを作成しますが、人間がすべてをレビューすることはできません。さて、何ですか？
記事のタイトル: AI がコードを書くが、人間がすべてをレビューすることはできない。さて、何ですか？ — シルヴァン・カラシュ
説明: AI は人間がレビューするよりも速くコードを生成できます。エンジニアリング チームは、レビュー、可逆性、説明責任を再設計する必要があります。

記事本文:
AI がコードを作成しますが、人間がすべてをレビューすることはできません。さて、何ですか？ — Sylvain Kalache ~/ ブログ エージェント セクション
お問い合わせ ブログに戻る AI がコードを作成しますが、人間がすべてをレビューすることはできません。さて、何ですか？
AI は人間がレビューするよりも速くコードを生成できます。エンジニアリング チームは、レビュー、可逆性、説明責任を再設計する必要があります。
AI が製品コードを記述できるかどうかに関する議論は終わりました。企業は多数のエージェントに仕事を引き渡しており、多くの企業では、本番環境に出荷するコードのほとんどをエージェントが作成しています。
次の課題は、エンジニアリング組織全体がこのようにフルスピードで実行されたときに起こるすべてのことです。コードを 10 倍速く生成したチームでも、依然として人間のスピードでコードをレビューしており、その不一致が制約となっています。開発者がエージェント プロセスを信頼しすぎるようになったことで、コードの所有権も問題になりつつあります。エージェントが制作を中断した場合、誰が責任を負いますか?
この記事は、エンジニアリング リーダーを特集し、Anthropic の技術スタッフのメンバーである Ian Sinnott が司会を務める、エンタープライズ規模での AI の実行に関する最近のパネルからの洞察に触発されています。
ボトルネックはコードの作成からレビューに移りました #
エージェントが人間の精査よりも早くコードを作成すると、コードレビューが制約になります。最近の調査によると、AI 導入率の高いチームは 2025 年に 1 年前と比べて 98% 多くのプル リクエストをマージしましたが、PR レビュー時間は 91% 増加し、PR サイズは 154% 増加しました。では、エンジニアを消耗させないスケーラブルなソリューションとは何でしょうか?
AI の増加がその解決策の一部であることは明らかですが、AI のレビューだけでは AI のリスクを解決できません。 Twingate のエンジニアリング担当副社長である Eran Kampf 氏は、「可逆性と判断密度に着目している」という 2 つの特定の指標に重点を置いていると述べました。依存関係のアップグレードなどの元に戻せる操作の場合、エージェントはコミットできます。壊れたら、それは

おそらくテストまたはステージングで壊れるでしょう。一方、データベース移行のような操作の場合は、人間がコードをレビューする必要があります。
徹底的にレビュープロセス自体を追求しました。最初は小規模なプル リクエストを強制しようとしましたが、「AI エージェントはそのようには機能しません」と CTO の Quentin Rousseau 氏は語りました。人間が段階的に考えて実行する一方で、コーディング エージェントは移行、モデル、サービス、テスト全体をワンショットで生成します。それを 5 つの PR に分割すると、レビュー担当者が必要とするコンテキストが分散するだけです。
Rootly は、すべての PR にリスク ラベルを添付し、人間がレビューする必要があるかどうかを決定するエージェントを構築しました。そのリスク レベルは、変更の大きさではなく、PR が本番環境を破壊した場合に何が起こるか、つまり影響の重大度と範囲に基づいています。各 PR では、作成者がその理由と内容を説明する必要がある特定のテンプレートも使用されており、アクセス制御、記録された例外、変更を安全に元に戻すためのロールバック計画をカバーするチェックリストが含まれています。
迅速なレビューではなく、可逆性を最適化します #
ロールバックに賭けることは、エージェント時代の必勝戦略のようです。ルソーは、人間によるレビューをエージェントの出力に合わせようとするのではなく、間違いを軽減することに投資しました。 「私たちは、自動化、ロールバック、カナリア デプロイメントを改善して CI/CD パイプラインを改善しました。そうすることで、人々ははるかに速く出荷できるようになりました。」現在、ほとんどのレビューはエージェントが処理します。 「正直に言うと、彼らは私よりもコードをよくレビューしています」と彼は言い、専任のセキュリティおよびベストプラクティスエージェントがすべての PR に対して組織のルールを適用しています。
パターンはパネルの外側でも保持されます。 Intercom のエンジニアリング チームは、デプロイとリリースを区別する機能フラグを使用して 1 日に約 180 回出荷し、1 分以内に機能をオフに切り替えることができます。デプロイメントがどのように行われているかを知るために、チームはモニターを停止しました

システムのみ。代わりに、結果に焦点を当てました。展開後にハートビート メトリック (システムの健全性ではなく顧客の結果) が低下した場合、自動ロールバックが起動されます。その甲斐あって、重大な変更によるダウンタイムは 35% 減少し、デプロイ量は 2 倍になりました。
物が壊れたら誰が責任を取るのですか？ #
可逆性は間違いを軽くしますが、より難しい質問には答えません。エージェントが事件を起こした場合、誰が責任を負うのでしょうか?
「身元を特定するだけでなく、エージェントの行動に対する説明責任も設定したいと考えています」と Port の CEO、Zohar Einy 氏は述べています。 「エージェントに人間の所有者を設定し、エージェントが行っていることに責任を負う人がいることを確認します。誰も責任を負わない場合は、エージェントを責めることができます。」
Cyber​​Ark の 2025 年のアイデンティティ セキュリティ ランドスケープでは、平均的な組織の人間ごとに 82 個のマシン ID が存在することが判明しており、ほとんどの企業はそれらすべてが何をしているのかを言うことができません。
Descope の AI 戦略および開発者エクスペリエンス責任者の Kevin Gao 氏は、その数字の下にある混乱について説明しました。 「10 ほどの異なるチームがあり、それぞれがまったく異なることを行うエージェントを構築しています。どのエージェントがどのサービスに接続しているかを簡単に確認できる統一された中央の場所はありません。」
ここでは、最初にエージェントに変更を指示した開発者に責任を負わせるのが常識です。しかし、プロセス全体が自動化されたらどうなるでしょうか? Einy 氏は、「所有者はそのコンポーネント、つまりエージェントが触れるシステムの部分を担当するチームになるだろう」と主張します。説明責任は、プロンプトを作成した人ではなく、コンポーネントの所有者にあります。
誇大広告のギャップ: 「AI ネイティブ」の人はまだいますか? #
聴衆はパネルに簡単な質問をしました。「今、エンジニアに何を期待すべきですか?」
コンセンサスは、速度はもはや重要ではないということでした。

エンジニアリングリーダーが求めているもの。代わりに品質と提供される価値を優先します。 「1 日に最も多くのコード行を出荷しているのは誰なのかということから、実際に価値を出荷しているのは誰なのかということになりました」とルソー氏は述べています。 Kampf 氏は、「今では誰でもいい加減な AI を実行できるようになり、たくさん出荷されていますが、それは価値のあるものですか、それともくだらないものですか?」と率直に言いました。
これは、Uber の最高執行責任者 (COO) である Andrew Macdonald 氏の言葉と同じです。彼は最近、Uber のコードの 70% は AI によって生成されており、製品価値への結びつきは「まだそこには至っていない」と述べました。同氏の言葉を借りれば、「これらの統計の 1 つと、『現在、実際に消費者向けの便利な機能が 25% 増加している』との間に線を引くのは非常に困難です。」
ほとんどのエンジニアリング組織は、エージェント時代に満足しているとは言えません。この機能は本物ですが、それを信頼性の高い実稼働グレードのエンジニアリング実践に変える作業は、ほとんどがまだ先のことです。 Einy にとって、ほとんどの企業はまだ「AI 支援」の段階にあります。
AI エンジニアリング、信頼性、そしてその仕事に携わる人々について書いています。何か役に立つことがあるときだけメールを送ります。
RSS 経由でフォローしてください。関連する執筆と講演

## Original Extract

AI can generate code faster than humans can review it. Engineering teams must redesign review, reversibility, and accountability.

AI Writes the Code, But Humans Can't Review It All. Now What? — Sylvain Kalache ~/ Blog Agent Section
Contact Back to the blog AI Writes the Code, But Humans Can't Review It All. Now What?
AI can generate code faster than humans can review it. Engineering teams must redesign review, reversibility, and accountability.
The debate about whether AI can write production code is over. Companies are handing work to fleets of agents, and at many of them, agents write most of the code that ships to production.
The next challenge is everything that happens once an entire engineering organization runs this way, at full speed. Teams that generate code 10 times faster still review it at human speed, and that mismatch is now the constraint. Code ownership is also becoming an issue as developers learn to trust agentic processes a little too much. When an agent breaks production, who is responsible?
This article is inspired by insights from a recent panel on running AI at enterprise scale , featuring engineering leaders and moderated by Ian Sinnott, a Member of Technical Staff at Anthropic.
The bottleneck moved from writing code to reviewing it #
When agents write code faster than humans can vet it, code review becomes the constraint. A recent survey found that high-AI-adoption teams merged 98% more pull requests in 2025 than a year prior, but PR review time rose 91% and PR size grew 154%. So what is a scalable solution that does not burn engineers out?
More AI is obviously part of the answer, but AI review alone does not fix AI risk. Twingate 's VP of Engineering, Eran Kampf , shared that they focus on two specific metrics: “We look at reversibility and judgment density.” For reversible operations like upgrading a dependency, the agent can commit; if it breaks, it will probably break in tests or staging. For operations like a database migration, on the other hand, you want a human to review the code.
Rootly went after the review process itself. It first tried to enforce small pull requests , but “AI agents don't work that way,” shared its CTO, Quentin Rousseau . While humans think and execute incrementally, coding agents generate an entire migration, model, service, and tests in one shot. Splitting that into five PRs just scatters the context a reviewer needs.
Rootly built an agent that attaches a risk label to every PR, which determines whether a human should review it. That risk level is based on what happens if the PR breaks production—the severity and scope of impact—rather than how big the change is. Each PR also uses a specific template that requires the author to explain the why and the what, and includes a checklist covering access control, logged exceptions, and a rollback plan for safely undoing the change.
Optimize for reversibility, not faster review #
Betting on rollbacks seems to be a winning strategy for the agentic era. Rather than trying to make human review keep pace with agent output, Rousseau invested in making mistakes cheaper. “We improved our CI/CD pipeline with better automation, rollbacks, canary deployments, and people can ship much faster that way.” Agents now handle most reviews. “To be honest, they review code better than I do,” he said, with dedicated security and best-practices agents enforcing the organization's rules on every PR.
The pattern holds outside the panel. Intercom's engineering team ships around 180 times a day , with feature flags that separate deploy from release and can switch a feature off in under a minute. To know how a deployment is doing, the team stopped monitoring systems only. Instead, it focused on outcomes: automatic rollbacks fire when heartbeat metrics— customer outcomes rather than system health —dip after a deployment. It paid off, with downtime from breaking changes dropping by 35% while deployment volume doubled.
Who is responsible when things break? #
Reversibility makes mistakes cheap, but it does not answer a harder question. When an agent does cause an incident, who is accountable?
“Beyond the identity, you also want to set the accountability for the agent's behavior,” said Port CEO Zohar Einy . “Set a human owner to the agent, making sure that there is someone accountable for what the agent is doing. When no one is accountable, you can blame the agent.”
CyberArk's 2025 Identity Security Landscape found 82 machine identities for every human in the average organization, and most companies cannot say what all of them are doing.
Descope 's Head of AI Strategy and Developer Experience, Kevin Gao , described the mess underneath that number. “You have like 10 different teams who have all built agents doing vastly different things, and there's no unified central place where I can just look and see which agents are connecting to which services.”
Here, common sense is to hold the developer who initially directed the agent to make the change accountable. But what happens when the whole process is automated? Einy argues that “the owner would be the team in charge of that component, the part in the system that the agent is touching.” Accountability stays with whoever owns the component, not whoever wrote the prompt.
The hype gap: is anyone “AI-native” yet? #
An audience member asked the panel a simple question: what should you expect from engineers now?
The consensus was that velocity is no longer what engineering leaders look for; they prioritize quality and delivered value instead. “It went from who is shipping the most lines of code per day to who is actually shipping value,” said Rousseau. Kampf was blunter: “Now that anyone can do slop AI, you ship a lot, but is it valuable stuff or crap?”
That echoes Uber COO Andrew Macdonald, who recently shared that 70% of Uber's code is AI-generated and that the link to product value is “not there yet.” As he put it, “It's very hard to draw a line between one of those stats and ‘now we're actually producing 25% more useful consumer features.’”
Most engineering organizations are far from being comfortable with the agentic era. The capability is real, but the work of turning it into a reliable, production-grade engineering practice is mostly still ahead. For Einy, most companies are still in the “AI-assisted” stage.
I write about AI engineering, reliability, and the people doing the work. I’ll only email when I have something useful to say.
follow via RSS . Related writing and talks
