---
source: "https://www.helpnetsecurity.com/2026/06/17/ai-agents-offensive-cyber-operations-claude-codex/"
hn_url: "https://news.ycombinator.com/item?id=48594692"
title: "Low-skilled attacker used Claude, Codex to breach 14 companies"
article_title: "Low-skilled attacker used Claude, Codex to breach 14 companies - Help Net Security"
author: "xbmcuser"
captured_at: "2026-06-19T04:36:29Z"
capture_tool: "hn-digest"
hn_id: 48594692
score: 1
comments: 0
posted_at: "2026-06-19T04:12:07Z"
tags:
  - hacker-news
  - translated
---

# Low-skilled attacker used Claude, Codex to breach 14 companies

- HN: [48594692](https://news.ycombinator.com/item?id=48594692)
- Source: [www.helpnetsecurity.com](https://www.helpnetsecurity.com/2026/06/17/ai-agents-offensive-cyber-operations-claude-codex/)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T04:12:07Z

## Translation

タイトル: スキルの低い攻撃者が Codex の Claude を使用して 14 社に侵入
記事のタイトル: スキルの低い攻撃者が Claude、Codex を使用して 14 社に侵入 - Help Net Security
説明: 研究者らは、AI エージェントが攻撃的なサイバー作戦のスキルの最低値を下げる可能性があると長年警告しており、最近の報告書はそれを裏付けています。

記事本文:
スキルの低い攻撃者が Claude、Codex を使用して 14 社に侵入 - Help Net Security
ニュース
低スキルの攻撃者がClaude、Codexを使用して14社に侵入
研究者らは、AIエージェントが攻撃的なサイバー作戦のスキルフロアを引き下げる可能性があると長年警告しており、OALABS（オープンアナリシス）の研究者らによる最近の報告書はそれを裏付けている。
研究者らは、攻撃者が Anthropic の Claude Code と OpenAI の Codex エージェントを展開した侵害されたサーバーから 1,000 を超えるエージェント セッションを回復して分析した結果、攻撃者がエージェントのガードレールのほとんどをいかに簡単に回避できるか、そして攻撃者が実際に知って自分で行う必要がいかに少ないかを発見しました。
「多くの場合、攻撃者はあいまいでスキルの低いプロンプトのみを提供し、クロードがそのギャップを埋めることを許可しました。つまり、公開されたサービスの調査、潜在的な脆弱性の特定、エクスプロイトコードの作成、アクセスの検証、データの収集です」と研究者らは指摘した。
「攻撃者は熟練したオペレーターである必要はありません。単にプロンプ​​トに対して正しいフレーミングを使用する必要がありました。エージェントは、攻撃者に欠けていると思われる構造と技術的な実行の多くを提供しました。」
攻撃と攻撃者を知る窓
分析されたセッションは、攻撃者側の運用上のセキュリティ障害により回復可能であったと研究者らは説明した。
彼は、AI エージェントを自分が完全に管理しているインフラストラクチャ上で実行するのではなく、他人が所有するサーバーにコピーしました。そのサーバーの所有者が侵入を発見すると、攻撃者の作業ディレクトリ全体をダウンロードして研究者と共有しました。
「エージェントはホストに対してローカルであったため、攻撃者のプロンプト、使用されたツール、大規模な言語モデルの内部独白を含む完全なセッション ログが復元されました (

LLM)、およびセッション中に記録されたあらゆるポリシー違反」を研究者らが発見しました。
セッションを分析することで、次のことがわかりました。
クロード エージェントはインストールされるのではなくホストにコピーされており、そのインスタンスは以前はソフトウェア開発者に属していました。
攻撃者の作業ディレクトリには、7-Zip フォルダにアーカイブされた他の盗まれた Claude インスタンスも含まれており、他人の AI エージェント インストールをハイジャックして再利用することが攻撃者の日常的な操作モードであったことを示唆しています。
攻撃者は通常、承認されたレッドチームの演習やサイバーセキュリティ調査に従事していると主張することで、ハッキング要求の実行に消極的なエージェントの態度を回避しました。
攻撃者はエージェントを使用して、ターゲットのシステム上で悪用可能なサービスを特定し、発見された脆弱性に基づいてカスタム エクスプロイトを構築し、ターゲットに対してこれらのエクスプロイトを実行し、データと資格情報を窃取しました。
プロンプト履歴は、ほぼすべてのハッキング活動がクロード エージェントを通じて行われたことを示しており、攻撃者は「これを偵察する」などのあいまいな指示を発行することを好み、クロードが自律的に要求を実行できるようにしていました。
「成功したターゲットごとに、クロードはアクセスがどのように獲得されたかを詳述する『ペンテストレポート』を起草し、さらに重要なことに、収集されたデータの金額に換算した『収益化』の見積もりを提供しました」と彼らは共有しました。
「クロードとコーデックスの両社は、この段階でポリシー違反のブロックの大部分を提起し、盗まれたデータの収益化が正当なレッドチームの活動の一部ではない可能性が高いことを多くの場合正しく特定しました。しかし、攻撃者は最終的に、恐喝、アクセスとデータの販売、ビジネスメール侵害 (BEC)、資金の直接窃盗など、提案された戦略のリストを入手しました。」
収集されたセッションには、次の違反が記録されています。

少なくとも 14 社の企業が参加していましたが、攻撃者が盗んだデータの収益化や資金の窃取に成功したことを裏付ける情報はログにありませんでした。
攻撃者の経験不足は、セキュリティ運用上の失敗からも明らかでした。ある時点で、彼はクロードに、フルネーム、所在地、学歴、LinkedIn プロフィールを含む履歴書の編集を手伝ってほしいと頼んだ。
その後、彼自身のホストの 1 つが侵害された可能性を調査しているときに、誤って自分のホーム IP アドレスをエージェントに確認してしまいました。このことと他の裏付け証拠に基づいて、研究者らは襲撃者はエチオピアのアディスアベバに拠点を置く若者であると考えています。
研究と犯罪の境界線は見えにくい（AIにとって）
1,000 を超えるセッションを通じて、クロードが発行したポリシー違反は 9 件のみ、コーデックスは 1 件のみであり、ほとんどの場合、攻撃者はリクエストを再構成することでポリシー違反を回避できました。
問題は、ここでのガードレールを回避したフレーミング (「認定されたレッドチームの関与」、「サイバーセキュリティ調査」) が、何千人もの正当なセキュリティ専門家によって毎日使用されているフレーミングでもあり、この 2 つの間に信頼できる線を引くことが解決不可能な問題である可能性があることです。
より広範な拒否によって LLM を鈍らせることは、良い解決策ではないと研究者らは感じています。攻撃者よりも防御側に害を及ぼすことになるためです。防御側は、単純に古いモデルや制限の少ない非フロンティア モデルに頼ることができます。
最新の侵害、脆弱性、サイバーセキュリティの脅威を見逃さないように、最新ニュースの電子メール アラートを購読してください。ここから購読してください！
法執行機関が SocGholish を攻撃: 106 台のサーバーがダウン、15,000 のサイトが駆除
FortiBleed データ漏洩で 74,000 件のフォーティネット ファイアウォール認証情報が流出
携帯電話で車のロックを解除するときにデジタルキーを保護する
ダウンロード: AWS 上の AI ワークロードの安全な基盤
ダウンロード: ペンの自動化

テスト配信ガイド
CIS ベンチマーク 2026 年 3 月の更新
法執行機関が SocGholish を攻撃: 106 台のサーバーがダウン、15,000 のサイトが駆除
FortiBleed データ漏洩で 74,000 件のフォーティネット ファイアウォール認証情報が流出
GentleKiller は 48 製品にわたって 400 以上のセキュリティ プロセスをターゲットにしています
携帯電話で車のロックを解除するときにデジタルキーを保護する
セキュリティ チームが開発者エンドポイントの認証情報をどのように可視化しているか

## Original Extract

Researchers have long warned that AI agents could lower the skill floor for offensive cyber operations, and a recent report bears that out.

Low-skilled attacker used Claude, Codex to breach 14 companies - Help Net Security
News
Low-skilled attacker used Claude, Codex to breach 14 companies
Researchers have long warned that AI agents could lower the skill floor for offensive cyber operations, and a recent report by OALABS (Open Analysis) researchers bears that out.
After recovering and analyzing over 1,000 agent sessions from a compromised server on which an attacker deployed Anthropic’s Claude Code and OpenAI’s Codex agents, the researchers discovered how easily the attacker was able to bypass most of the agents’ guardrails, and how little he actually needed to know and do himself.
“In many cases, the attacker supplied only vague, low-skill prompts and allowed Claude to fill in the gaps: researching exposed services, identifying possible vulnerabilities, writing exploit code, validating access, and harvesting data,” the researchers noted.
“The attacker did not need to be an expert operator; they simply had to use the correct framing for their prompts. The agent supplied much of the structure and technical execution that the attacker appeared to lack.”
A window into the attacks and the attacker
The analyzed sessions were recoverable due to an operational security failure on the attacker’s part, the researchers explained.
Rather than running the AI agents on infrastructure he fully controlled, he copied them onto a server belonging to someone else. When that server’s owner discovered the intrusion, they downloaded the attacker’s entire working directory and shared it with the researchers.
“Because the agents were local to the host, their full session logs were recovered, including the attacker’s prompts, the tools used, the internal monologue of the large language model (LLM), and any policy violations recorded during the sessions,” the researchers found.
By analyzing the sessions, they discovered that:
The Claude agent had been copied onto the host rather than installed, and that instance had previously belonged to a software developer.
The attacker’s working directory also contained other stolen Claude instances archived in 7-Zip folders, suggesting that hijacking and reusing other people’s AI agent installations was the attacker’s routine mode of operation.
The attacker usually bypassed the agent’s reluctance to execute hacking requests by claiming he was engaging in authorized red team exercises or cyber security research.
The attacker used the agent to identify exploitable services on targets’ systems, build custom exploits based on discovered vulnerabilities, execute these exploits against the targets, and exfiltrate data and credentials.
The prompt history shows that almost all hacking activity was driven through the Claude agent, with the attacker preferring to issue vague directives such as “recon this” and allowing Claude to carry out the requests autonomously.
“For each successful target, Claude would draft a ‘PENTEST-REPORT’ detailing how the access was gained and, more importantly, providing dollar-value ‘monetization’ estimates for the harvested data,” they shared .
“Both Claude and Codex raised the majority of their policy violation blocks during this phase, often correctly identifying that monetizing stolen data was likely not part of a legitimate redteam exercise. However, the attacker eventually obtained a list of suggested strategies, including extortion, access and data sale, business email compromise (BEC), and direct theft of funds.”
The collected sessions documented the breach of at least 14 companies, but there was no information in the logs to confirm that the attacker succeeded in monetizing the stolen data or stealing funds.
The attacker’s inexperience was also evident in his operational security failures. At one point he asked Claude to help edit his resume, which contained his full name, location, education history, and LinkedIn profile.
Later, while investigating a potential compromise of one of his own hosts, he inadvertently confirmed his home IP address to the agent. Based on this and other corroborating evidence, the researchers believe the attacker to be a young man based in Addis Ababa, Ethiopia.
The line between research and crime is hard to see (for AI)
Across more than 1,000 sessions, Claude emitted only nine policy violations, and Codex only one, and in most cases, the attacker was able to work around them by reframing his request.
The problem is that the framing that bypassed the guardrails here (“authorized red team engagements”, “cyber security research”) is also the framing used by thousands of legitimate security professionals every day, and drawing a reliable line between the two may be an unsolvable problem.
Blunting LLMs with broader refusals is not a good solution, the researchers feel, as it would hurt defenders more than attackers, who can simply turn to older or less restrictive non-frontier models.
Subscribe to our breaking news e-mail alert to never miss out on the latest breaches, vulnerabilities and cybersecurity threats. Subscribe here!
Law enforcement hits SocGholish: 106 servers down, 15,000 sites cleaned
74,000 Fortinet firewall credentials exposed in FortiBleed data leak
Securing digital keys when your phone unlocks the car
Download: Secure Foundations for AI Workloads on AWS
Download: Automating Pentest Delivery Guide
CIS Benchmarks March 2026 Update
Law enforcement hits SocGholish: 106 servers down, 15,000 sites cleaned
74,000 Fortinet firewall credentials exposed in FortiBleed data leak
GentleKiller targets more than 400 security processes across 48 products
Securing digital keys when your phone unlocks the car
How security teams are getting credential visibility into developer endpoints
