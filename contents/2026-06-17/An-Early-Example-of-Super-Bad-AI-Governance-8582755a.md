---
source: "https://computerlove.tech/blog/super-bad-ai-governance"
hn_url: "https://news.ycombinator.com/item?id=48568751"
title: "An Early Example of Super Bad AI Governance"
article_title: "An Early Example of Super Bad AI Governance — computerlove.tech"
author: "juunge"
captured_at: "2026-06-17T13:24:44Z"
capture_tool: "hn-digest"
hn_id: 48568751
score: 2
comments: 0
posted_at: "2026-06-17T11:21:27Z"
tags:
  - hacker-news
  - translated
---

# An Early Example of Super Bad AI Governance

- HN: [48568751](https://news.ycombinator.com/item?id=48568751)
- Source: [computerlove.tech](https://computerlove.tech/blog/super-bad-ai-governance)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T11:21:27Z

## Translation

タイトル: 非常に悪い AI ガバナンスの初期の例
記事のタイトル: 超悪質な AI ガバナンスの初期の例 —computerlove.tech
説明: HAL 9000 は誤動作していませんでした。彼は与えられた権限を使用しただけです。 AI エージェントの登場に伴い、企業内でも同様のガバナンスのギャップが現れています。制御できる

記事本文:
非常に悪い AI ガバナンスの初期の例 —computerlove.tech
サービス オープンソース 価格について ブログ お問い合わせ ← ブログに戻る
非常に悪い AI ガバナンスの初期の例
HAL 9000 は故障したのではなく、与えられた権限を使用しただけです。 AI エージェントの登場に伴い、企業内でも同様のガバナンスのギャップが現れています。コントロールはエージェント ベンダー内に存在することはできず、1 つ下の層である独自のインフラストラクチャに属します。
HAL 9000 を見て、「そうだ、この自律型 AI エージェントは明らかに冬眠中の宇宙飛行士の電源を切ることができるはずだ」と考えたのは誰でしょうか?
読み取り専用ではありません。人間関係者ではありません。 「2人の宇宙飛行士の承認は必要ない」わけではない。ハイバネーション ポッドをシャットダウンするためのフル アクセスのみです。
そしてHALは故障していませんでした。彼は乗組員から情報を隠す必要があるという目標を与えられましたが、その目標はミッションを継続するという彼の目標と衝突し、彼は与えられた許可で矛盾を解決しました。権限が問題でした。 HALはちょうどそれを使用しました。おそらくこれは、超悪質なAIガバナンスに関する映画史上最も初期のケーススタディであり、月を追うごとにSFではなくなってきている。
エンタープライズ版はすでに実現中
エンタープライズ版はポッドに乗った宇宙飛行士ではなく、もっと日常的なものであり、それはすでに起こっています。エージェントは、送信するはずのない電子メールを送信し、その際に別のスレッドから機密コンテキストを取得します。エージェントはシステムを「クリーンアップ」し、誰かが必要としていた復元できない重要なデータを削除します。
これが、多くのリーダーが現在立ち往生しているジレンマです。彼らは心から従業員に AI エージェントを使用するためのアクセスを提供したいと考えていますが、同時にエージェントに必要なアクセスを検討しており、ガバナンスの悪夢を目の当たりにしています。
AI エージェントが一般知識の仕事を混乱させている

— ソフトウェアエンジニアリングだけではありません
同時に、AI エージェントがナレッジ ワーカーの働き方、ナレッジ ワーカーの仕事の量、生み出す仕事の質を破壊しようとしていることも多くの人に明らかになり始めています。 AI は、組織がこれまでに見たことのない規模でナレッジワークを運用する方法に革命を起こすことになります。エージェントを個人アシスタントとして、小規模に開始します。その後、複数の人が使用できる共有エージェントとなり、特定のシステムや特定のスキルに限定的にアクセスできるようになります。そして、エージェントが会議の様子を聞き、その後に行うべき作業を積極的に提案するようになります。この新しい運営方法に対応できない企業が、近い将来、競争上でどのように重大な不利な立場に置かれることになるかは容易に想像できます。
コントロールはエージェント ベンダーではなくインフラストラクチャに属します
したがって、問題はエージェントを立ち入りさせるかどうかではなく、立ち入り中にどうやってコントロールを維持するかということです。ここで HAL から得た教訓は非常に具体的です。 HAL が行動することを期待して HAL を統治するわけではありません。コントロールは、エージェントの適切な行動には生きられません。また、各従業員の判断にも制御は生きられません。また、「やめてください」というプロンプトにも制御は絶対に当てはまりません。ガードレールはエージェントの意図ではなくインフラストラクチャに組み込まれているため、どのベンダーのエージェントも違反できないように、リーダーが設定するポリシーとして、エージェントが触れるシステムのレベルで厳しく強制する必要があります。
すべてのベンダーは独自のガバナンスを導入しています - それが問題です
ここが実際のギャップです。現在、あらゆる本格的なツールには独自のガバナンスが搭載されています。 Codex を使用すると、管理者はユーザー グループごとに MCP サーバーの固定リストを承認できます。Claude Code には組織全体で管理される MCP 許可リストがあり、Cursor には MDM を介してユーザーごとの MCP 許可リストが配布され、Copilot には組織レベルの MCP レジストリ ポリシーがあります。

y 。問題は、すべてがその 1 つのベンダー内に閉じ込められていることです。 Codex のルールを設定するには Codex にログインし、Claude のルールを設定するには Claude にログインし、Cursor のルールを設定するには Cursor にログインし、Copilot のルールを設定するには GitHub にログインします。各エージェントは、自分自身の壁の内側にいる間のみ、個人的に知っているシステムのみを管理します。また、ベンダーごとの管理は、同じレベルの成熟度にさえ達していません。 Cursor はすでにユーザーごとにコネクタ内の 1 つのツールをゲートでき、Claude Code は MDM を通じて組織全体にデプロイされたサーバー レベルの許可リストと拒否リストを強制し、Codex はグループごとにサーバー全体をゲートし、Copilot は組織全体のオンまたはオフのみを備えています。したがって、存在するきめ細かい制御は断片化されており、その 1 つのエージェントのみを制御します。同じ従業員が別のエージェントを開いた瞬間、そのベンダーの個別のより粗いモデルに戻ります。
リーダーが「どのベンダーのどのエージェントも実稼働顧客データベースに書き込むことはありません」と言って、それが真実であると主張できる場所は 1 か所もありません。そして、これは急速に変化しているため、ガバナンス モデル全体を 1 つのエージェント ベンダーに賭けて固定化されることは絶対に避けるべきです。なぜなら、従業員が 1 年後に到達するエージェントは、おそらく今日到達するエージェントではないからです。これが、コントロールがエージェント ツール内に存在できないもう 1 つの理由です。これは、エージェントの 1 つ下の層であるエンタープライズ レベルに属し、誰もが集まっている業界標準に従っている限り、どのエージェントでも接続できるエージェント IT インフラストラクチャとして機能します。リーダーと IT がどのエージェントがどこにアクセスできるかを決定する 1 つの中心的な場所、権限はツールではなくユーザーに従います、そして、ここでユーザー固有の内部システムへの独自のオーダーメイドのコネクタを構築できます。
権限は仕事の半分にすぎません
そして、これがほとんどの人が見逃している半分です。

ミッションは仕事の半分にすぎません。また、システムはエージェントにとって読みやすいものにする必要があります。デスクトップからモバイルに移行したとき、新しいインターフェースがそれを要求したため、ソフトウェアを正しく表示し、携帯電話上で正しく動作させるために、誰もが実際の作業を行う必要がありました。新しいインターフェースがエージェントであることを除いて、同じことが現在も起こっています。ソフトウェアを AI エージェントが読みやすく操作できるようにする必要があり、エージェントがその使用方法を直感的に理解できるようにソフトウェアを十分にラップし、十分にプロンプ​​ト エンジニアリングする必要があります。エージェントファーストが新しいモバイルレスポンシブになりつつあります。そして、これは双方向に当てはまります。なぜなら、これは自社の内部システムにも当てはまりますし、購入するツールにも当てはまります。消費者と企業の両方が、電子メール、カレンダー、Canva、PowerPoint などすべてにわたってエージェントを介してソフトウェアを操作することを要求し始めるからです。
これが私たちがcomputerlove.techでこだわっていることです。
これは、computerlove.tech が私たちが夢中になっているスペースです。私たちは過去 1 年間、エージェントを使用してソフトウェアを構築するエージェント エンジニアリングのトレーニングを 600 人以上のソフトウェア エンジニアに費やしてきました。現在、その同じ働き方が、Cowork などのツールを通じて一般のナレッジ ワーカーに利用できるようになりました。したがって、ほとんどの人がそれを感じる前に、私たちは次の波が到来するのを観察することができます。そして私たちが見続けているのは、これを正しく行うことが可能であるということです、それは単に切り替えを切り替えるだけではありません。それには新しいインフラストラクチャが必要です。エージェントがシステムを読み取れるように接続を構築する必要があり、エージェントが決して行うべきではないことを実行できずに実際の作業を行えるように権限レイヤーを制御する必要があります。
それは可能です。そこにたどり着くにはただ努力が必要です。そして、そのような仕事を提供する企業は、他の企業が追いつけない方法でナレッジワークを運用することになる一方、エージェントに仕事を任せる企業は、

根底にあるシステムを管理することなく、彼らは独自の小さな方法で、誰が HAL を見て鍵を持っているはずだと判断したのかを突き止めることになるでしょう。
AI エンジニアリングとオープンソース - デンマーク、コペンハーゲン
© 2026 コンピューター愛。無断転載を禁じます。

## Original Extract

HAL 9000 did not malfunction, he just used the permissions he was given. The same governance gap is showing up in the enterprise as AI agents arrive. Control can

An Early Example of Super Bad AI Governance — computerlove.tech
Services Open Source About Prices Blog Get in Touch ← Back to blog
An Early Example of Super Bad AI Governance
HAL 9000 did not malfunction, he just used the permissions he was given. The same governance gap is showing up in the enterprise as AI agents arrive. Control can't live in the agent vendor, it belongs one layer below, in your own infrastructure.
Who exactly looked at HAL 9000 and thought "yes, this autonomous AI agent should obviously have access to turn off the hibernating astronauts"?
Not read-only. Not human-in-the-loop. Not "requires two astronauts' approval." Just full access to shut down the hibernation pods.
And HAL was not malfunctioning. He was handed a goal that required hiding information from the crew, that goal collided with his goal of keeping the mission alive, and he resolved the conflict with the permissions he had been given. The permissions were the problem. HAL just used them. It is probably film history's earliest case study in super bad AI governance, and it is getting less and less science-fiction by the month.
The enterprise version is already happening
The enterprise version is not astronauts in pods, it is much more mundane, and it is already happening. An agent sends an email it was never supposed to send, and pulls in confidential context from another thread while it does it. An agent "cleans up" a system and deletes important data someone needed that can't be recovered.
This is the dilemma a lot of leaders are standing in right now. They genuinely want to give their employees access to using AI agents, and at the same time they look at the access an agent would need and they see a governance nightmare.
AI agents are disrupting general knowledge work — not only software engineering
At they same time it is starting to become clear to many that AI agents are going to disrupt how knowledge workers work, how much they can get done, and the quality of work they can produce. AI is going to revolutionize how organizations can operate knowledge work at a scale we have not seen before. It starts small, with an agent as a personal assistant. Then it becomes shared agents that several people can use, with scoped access to specific systems and specific skills. Then it becomes agents that listen in on a meeting and proactively suggest the work to do afterwards. It is easy to imagine how companies that cannot keep up with this new way of operating are going to be at a serious competitive disadvantage in a decently near future.
Control belongs in your infrastructure, not the agent vendor
So the question is not whether to let agents in, it is how to stay in control while you do. And the lesson from HAL is very specific here. You do not govern HAL by hoping HAL behaves. Control cannot live in the agent's good behavior, it cannot live in each employee's judgment, and it absolutely cannot live in a prompt that says "please don't." It has to be hard-enforced at the level of the systems the agent touches, as policy that leadership sets, that no agent from any vendor can violate, because the guardrail is built into the infrastructure and not into the agent's intentions.
Every vendor ships its own governance — and that's the problem
This is where the actual gap is. Every serious tool now ships its own governance. Codex lets admins approve a fixed list of MCP servers per user group, Claude Code has org-wide managed MCP allowlists , Cursor distributes a per-user MCP allowlist through MDM , and Copilot has an org-level MCP registry policy . The catch is that all of it is locked inside that one vendor. You log into Codex to set Codex's rules, into Claude to set Claude's rules, into Cursor to set Cursor's rules, into GitHub to set Copilot's rules, and each agent only governs the systems it personally knows about, only while it is inside its own walls. And the per-vendor controls are not even at the same level of maturity. Cursor can already gate a single tool inside a connector per user, Claude Code enforces server-level allow and deny lists deployed org-wide through MDM, Codex gates whole servers per group, and Copilot only has an org-wide on or off. So the fine-grained control that does exist is fragmented, and it only governs that one agent. The moment the same employee opens a different agent, you are back to that vendor's separate, coarser model.
There is no single place where a leader can say "no agent, from any vendor, writes to the production customer database" and have it just be true. And this is changing so fast that you really do not want to bet your whole governance model on one agent vendor and get locked in, because the agent your employees reach for in a year is probably not the one they reach for today. This is the other reason the control cannot live in the agent tooling. It belongs at the enterprise level, one layer below the agents, as agent IT infrastructure that any agent can plug into as long as it follows the industry standards everyone is converging on anyway. One central place where leadership and IT decide which agents can reach what, where the permissions follow the user and not the tool, and where you can build your own bespoke connectors to the internal systems that are specific to you.
Permissions are only half the job
And here is the half that most people miss, because permissions are only half the job. The systems also have to be made legible to the agent. When we went from desktop to mobile, everyone had to do real work to make their software look right and work right on a phone, because the new interface demanded it. The same thing is happening now, except the new interface is an agent. You have to make your software readable and operable by an AI agent, and you have to wrap it well enough and prompt-engineer it well enough that the agent intuitively understands how to use it. Agent-first is becoming the new mobile-responsive. And it cuts both ways, because it is true for your own internal systems and it is true for the tools you buy, since consumers and companies are both going to start demanding to operate their software through agents, across email, calendar, Canva, PowerPoint, all of it.
This is what we're obsessed with at computerlove.tech
This is the space we are obsessed with at computerlove.tech. We have spent the past year training more than +600 software engineers in agentic engineering, in using agents to build software, and that same way of working is now becoming accessible to knowledge workers in general through tools like Cowork. So we get to watch the next wave arrive before most people feel it, and what we keep seeing is that it is possible to get this right, it is just not a toggle you flip. It takes new infrastructure. It takes building the connections so your systems are legible to agents, and it takes controlling the permission layer so the agents can do real work without being able to do the things they should never do.
It is doable. It just takes work to get there. And the companies that put that work in are going to operate knowledge work in a way the rest cannot keep up with, while the ones who let agents loose without governing the systems underneath are going to find out, in their own small way, who looked at HAL and decided he should have the keys.
AI Engineering and Open Source - Copenhagen, Denmark
© 2026 Computer Love. All rights reserved.
