---
source: "https://www.bleepingcomputer.com/news/security/every-ai-agent-is-an-identity-most-organizations-dont-treat-them-that-way/"
hn_url: "https://news.ycombinator.com/item?id=48598282"
title: "Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way"
article_title: "Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way"
author: "ilreb"
captured_at: "2026-06-19T16:05:48Z"
capture_tool: "hn-digest"
hn_id: 48598282
score: 2
comments: 1
posted_at: "2026-06-19T13:23:13Z"
tags:
  - hacker-news
  - translated
---

# Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way

- HN: [48598282](https://news.ycombinator.com/item?id=48598282)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/every-ai-agent-is-an-identity-most-organizations-dont-treat-them-that-way/)
- Score: 2
- Comments: 1
- Posted: 2026-06-19T13:23:13Z

## Translation

タイトル: すべての AI エージェントはアイデンティティです。ほとんどの組織は彼らをそのように扱っていません
説明: AI エージェントは、多くの場合、ほとんど監視されることなく、データにアクセスし、ワークフローをトリガーし、コードをデプロイし、重要なビジネス システムと対話できます。 Token Security は、AI エージェントがアイデンティティとガバナンスの新たな課題となっている理由を分析します。

記事本文:
すべての AI エージェントはアイデンティティです。ほとんどの組織は彼らをそのように扱っていません
FortiBleed のリークにより、73,000 台のデバイスのフォーティネット VPN 認証情報が公開されました。
Microsoft、RoguePlanet ゼロデイ向けの Defender パッチの開発に取り組んでいる
悪意のある JetBrains Marketplace プラグインが開発者から AI API キーを盗む
Steam ワークショップが悪用され、Wallpaper Engine アプリ経由でマルウェアが拡散
すべての AI エージェントはアイデンティティです。ほとんどの組織は彼らをそのように扱っていません
ウェビナー: 攻撃者が MFA を回避する方法と防御者がどのように対応できるか
Microsoft: 2026 年 6 月の Windows アップデートでごみ箱プロンプトが表示されなくなる
ファイルを保持し、料金はかかりません - 1 TB の生涯ストレージはわずか 130 ドルです
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
すべての AI エージェントはアイデンティティです。ほとんどの組織は彼らをそのように扱っていません
すべての AI エージェントはアイデンティティです。ほとんどの組織は彼らをそのように扱っていません
セキュリティ チームは長年にわたり、ID を制御すればリスクも制御できるという単純な前提に基づいてプログラムを構築してきました。従業員は ID プロバイダーを通じて認証します。サービス アカウントはシステムを接続します。 API キーを使用すると、ワークロードがクラウド サービスやデータベースと通信できるようになります。
俳優たちは非常に予測可能でした。その結果、アイデンティティ セキュリティとガバナンス モデルはその予測可能性に従っています。今、この前提は崩れています。
AI エージェントは静かに企業に入り込み、会議を要約し、電子メールの下書きを作成し、従業員の情報検索を支援しました。ほとんどのセキュリティ チームは、最初はそれらについて深く考えていませんでした。製品のように見えました

アクティビティツール、それがまさにそれだったからです。
その後、組織は、Salesforce、Snowflake、GitHub、Jira、本番データベース、クラウド環境などの重要なビジネス サービスにそれらを接続し始めました。現在では、情報を取得し、ワークフローをトリガーし、レコードを更新し、コードを作成して展開し、複数のシステムにわたってアクションを実行します。
時には人間に代わって、時には自律的に、そして時にはどちらであるか本当に不明瞭な方法で。
これにより、AI エージェントは単なるツールではなくなります。それが彼らのアイデンティティとなり、ほとんどの企業には彼らのためのセキュリティとガバナンスのモデルがありません。
このパターンは組織全体で一貫しています。新しい ID レイヤーは、ID チームが過去 10 年間かけて導入してきた制御がほとんどない状態で、既存のインフラストラクチャ上に構築されます。エージェントは、あるチームによって作成され、別のチームによって使用され、5 つの異なるアプリケーションに接続され、まったく異なる目的のためにプロビジョニングされた資格情報で実行される可能性があります。
誰かがそれを動作させる必要があり、作業を遅らせたくなかったため、早期に広くアクセスされました。その結果、権限が高く可視性が低い攻撃者が無秩序に蔓延し、ほとんどのセキュリティ チームが管理することはおろか、管理することもできません。
恐怖で行動を遅らせないでください。トークン セキュリティを味方に付けて大規模な AI を実現します。
AI エージェントは、従来の IAM 制御を上回るマシン速度で ID を作成、使用、ローテーションします。
Token Security は、チームが AI エージェント ID のライフサイクル全体を管理し、修復によってリスクを軽減し、スピードを犠牲にすることなくガバナンスと監査の準備を維持するのに役立ちます。
私たち Token Security が委託した 2026 年の CSA 調査によると、組織の 82% が過去 1 年間にセキュリティ、IT、ガバナンス チームの知識なしに作成された AI エージェントを少なくとも 1 つ発見し、41% がこれを発見しました。

何度も起こっている。
ここで、セキュリティに関する議論が脇道に逸れてしまいました。 AI セキュリティに関する注目のほとんどは、プロンプト インジェクション、ジェイルブレイク、安全でない出力などのモデル リスクに集中しています。これらはすべてエージェント AI エコシステムの重要な部分ですが、企業のセキュリティ チームが必要とする全体像を描いているわけではありません。彼らが必要とする最も重要な部分は、エージェントが実際に何にアクセスできるのかを答える必要があります。
公的文書を要約するエージェントの爆発範囲は限られています。顧客記録、ソース コード、財務システム、管理者レベルのクラウド認証情報にエージェントが接続される場合は、まったく別の問題になります。
不正なプロンプト、侵害されたセッション、悪意のあるプラグイン、または誤って設定された統合により、過剰な権限を持つエージェントが、データの漏洩、破壊的なアクション、または接続されるはずのないシステムを経由した水平移動のパスに変わる可能性があります。
これはもはや理論上の話ではなく、過去 1 年間に組織の 65% が AI エージェントに関係するセキュリティ インシデントを経験し、61% がその結果として機密データの漏洩または誤った取り扱いを報告しました (ソース)。
コントロールの取得は可視化から始まります。セキュリティ チームは、実際に重要な質問に答えるために、単なる名前やプラットフォームを超えた AI エージェントの検出とインベントリを必要としています。
このエージェントの所有者は誰ですか?誰がそれを呼び出すことができますか?どのようなシステムに接続されていますか?どのような資格情報が使用されますか?各ターゲット アプリケーションで何を読み取り、書き込み、削除、または実行できますか?
表面が明確ではないため、これは思っているよりも難しいです。セキュリティ チームは、セールス アシスタントが管理者権限を持つ Snowflake サービス アカウントで実行されていることを知らずに、AI プラットフォームにセールス アシスタントが存在することを知っている可能性があります。どのシークレット、リポジトリ、CI/CD パイプラインが分からなくても、コーディング エージェントが開発者エンドポイントにインストールされていることは知っている可能性があります。

それは到達することができます。
エージェント自体は全体像の一部にすぎません。エージェントのアイデンティティが触れることができるものはすべて、実際の露出面です。
2番目の部分は目的です。 AI エージェントでは、セキュリティとガバナンスを純粋に権限ベースにすることはできません。エージェントの意図を考慮する必要があります。営業準備エージェントに必要なのは、CRM レコードへの読み取りアクセスのみです。データベーステーブルを削除する必要はありません。
財務ワークフロー エージェントは請求書のみを読む必要があります。新しい特権ユーザーを作成できないようにする必要があります。エージェントが何を行うべきかを理解すると、その権限がそのスコープに一致するかどうかを評価できます。そして、今日実際には、そうしたことはほとんどなく、そのギャップこそが本当のリスクの所在であり、最小特権ポリシーの変動によって時間の経過とともに拡大するだけです。
意図が理解されると強制執行が可能になります。エージェントの実際の目的に合わせて権限をトリミングしたり、過剰な権限を持つサービス アカウントを修正したり、未使用の資格情報をローテーションまたは削除したり、危険な接続をインシデントに発展する前に捕捉したりできます。
ほとんどのチームがつまずくのは、これらの作業が 1 回限りではないことです。アクセス レビューや監査は進歩しているように感じるかもしれませんが、それらは特定の時点のチェックボックスと誤ったセキュリティ感を提供するだけです。その理由は、エージェントが変わり、指示が更新され、ユーザーベースが変化し、統合が拡大するためです。
狭い内部ツールとして開始されたエージェントは、誰かが間違った決定をしたからではなく、スコープが忍び寄ったときに誰も見ていなかったために、接触するように設計されていなかったシステムにひっそりと接続されてしまう可能性があります。
そのため、通常のパターンから外れてアプリケーションにアクセスし始めたり、予期しない認証情報を使用したり、定められた目的に適合しないアクションを実行したりするエージェントを捕捉するために、ガバナンスを継続的に行う必要があります。
AIで成功する企業は生き残れない

エージェントを完全にブロックする es。これらはエージェントを管理可能にし、安全な AI イノベーションを促進するものになります。これは、所有者、アクセス、動作、リスク、ライフサイクル制御を備えたファーストクラスのアイデンティティとしてそれらを扱うことを意味します。
AI エージェントは特権的な内部関係者になりつつあります。セキュリティおよびアイデンティティ プログラムは、内部関係者が目に見えない攻撃経路になる前に追いつく必要があります。
Token Security がこの問題にどのように取り組んでいるかをぜひご紹介します。安全性を犠牲にすることなく拡張できるように、デモを予約して技術チームとチャットしてください。
Token Security によってスポンサーおよび執筆されました。
FortiBleed のリークにより、73,000 台のデバイスのフォーティネット VPN 認証情報が公開されました。
Steam ワークショップが悪用され、Wallpaper Engine アプリ経由でマルウェアが拡散
ランサムウェア集団が Microsoft Teams リレーを悪用して悪意のあるトラフィックを隠蔽
AI を本番環境で機能させるには何が必要ですか?オンデマンドのワークフローで調べてください。
パスワードのヘルスチェックの期限を過ぎましたか? Active Directory を無料で監査する
問題を AI エージェントのスプロールにマッピングするだけでなく、修正してください。
AI はデータ侵害の時限爆弾です: 新しいレポートを読む
Microsoft 365 はデータを完全には保護しません。 Acronis Cyber​​ Protect がどのようにギャップを埋めるかをご覧ください。
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

AI agents can access data, trigger workflows, deploy code, and interact with critical business systems, often with little oversight. Token Security breaks down why AI agents are becoming a new identity and governance challenge.

Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way
FortiBleed leak exposes Fortinet VPN credentials for 73,000 devices.
Microsoft working on Defender patch for RoguePlanet zero-day
Malicious JetBrains Marketplace plugins steal AI API keys from developers
Steam Workshop abused to spread malware via Wallpaper Engine app
Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way
Webinar: How attackers bypass MFA and how defenders can respond
Microsoft: June 2026 Windows updates break Recycle Bin prompts
Keep your files, lose the charges—1TB of lifetime storage is just $130
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way
Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way
For years, security teams built their programs around a simple premise of if you control the identities, you can control the risk. Employees authenticate through identity providers. Service accounts connect systems. API keys let workloads talk to cloud services and databases.
The actors have been very predictable. And as a result, the identity security and governance model have followed that predictability. Now, this premise is breaking.
AI agents entered the enterprise quietly, summarizing meetings, drafting emails, helping employees find information. Most security teams didn't think hard about them at first. They looked like productivity tools, because that is exactly what they were.
Then, organizations started connecting them to critical business services such as Salesforce, Snowflake, GitHub, Jira, production databases, and cloud environments. Now, they retrieve information, trigger workflows, update records, write and deploy code, and take actions across multiple systems.
Sometimes on the behalf of a human, sometimes autonomously, and sometimes in ways where it's genuinely unclear which.
This makes AI agents more than just tools. It makes them identities and most enterprises have no security and governance models for them.
The pattern is consistent across organizations. A new identity layer gets built on top of existing infrastructure with almost none of the controls that identity teams spent the last decade putting in place. An agent might be created by one team, used by another, connected to five different applications, and running on credentials that were provisioned for a completely different purpose.
It got broad access early because someone needed it to work and didn't want to slow things down. The result is a sprawl of high-privilege, low-visibility actors that most security teams can't inventory, let alone govern.
Don't let fear slow you down. AI at scale with Token Security on your side.
AI agents create, use, and rotate identities at machine speed, outpacing traditional IAM controls.
Token Security helps teams manage the full lifecycle of AI agent identities, reduce risk with remediation, and maintain governance and audit readiness without sacrificing speed.
According to a 2026 CSA survey commissioned by us here at Token Security, 82% of organizations discovered at least one AI agent created without the knowledge of security, IT, or governance teams in the past year, and 41% found this happening multiple times.
Here's where the security conversation has gone sideways. Most of the attention on AI security has landed on model risk, such as prompt injection, jailbreaks, unsafe outputs. While these are all an important part of the agentic AI ecosystem, they don’t paint the complete picture enterprise security teams require. The most important piece they need must answer what can the agent actually access?
An agent that summarizes public documentation has limited blast radius. An agent connected to customer records, source code, financial systems, and admin-level cloud credentials is a different problem entirely.
A bad prompt, a compromised session, a malicious plugin, or a misconfigured integration can turn an overprivileged agent into a path for data exfiltration, destructive action, or lateral movement through systems that were never meant to be connected.
This is no longer theoretical, 65% of organizations experienced a security incident involving an AI agent in the past year, with 61% reporting exposure or mishandling of sensitive data as a result ( source ).
Getting control starts with visibility. Security teams need AI agent discovery and inventory that extends beyond just names and platforms to answer questions that actually matter.
Who owns this agent? Who can invoke it? What systems is it connected to? What credentials does it use? What can it read, write, delete, or execute in each target application?
This is harder than it sounds, because the surface isn't obvious. A security team might know a sales assistant exists in an AI platform without knowing it runs on a Snowflake service account with admin privileges. They might know a coding agent is installed on developer endpoints without knowing which secrets, repositories, and CI/CD pipelines it can reach.
The agent itself is only part of the picture. Everything the agent's identities can touch is the actual exposure surface.
The second piece is purpose. Security and governance can't be purely permission-based with AI agents. It has to account for the agent’s intent. A sales prep agent only needs read access to CRM records. It doesn't need to delete database tables.
A finance workflow agent should only read invoices. It shouldn't be able to create new privileged users. When you understand what an agent is supposed to do, you can evaluate whether its permissions match that scope. And, in practice today, they rarely do and that gap is where the real risk lives and it only widens over time through least privilege policy drift .
Once intent is understood, enforcement becomes possible. Permissions can be trimmed to match the agent’s actual purpose, overprivileged service accounts remediated, unused credentials rotated or removed, and risky connections caught before they turn into incidents.
The part that trips up most teams is that none of this is a one-time exercise. An access review or an audit may feel like progress, but they just provide a point-in-time checkbox and a false sense of security. The reason is that agents change, instructions update, user bases shift, and integrations expand.
An agent that started as a narrow internal tool can quietly end up connected to systems it was never designed to touch, not because anyone made a bad decision, but because nobody was watching when the scope crept.
That's why governance needs to be continuous to catch agents that start accessing applications outside their normal pattern, use unexpected credentials, or take actions that don't fit their stated purpose.
The enterprises that succeed with AI will not be the ones that block agents entirely. They will be the ones that make agents governable and promote secure AI innovation. This means treating them as first-class identities with owners, access, behavior, risk, and lifecycle controls.
AI agents are becoming privileged insiders. Security and identity programs must now catch up before those insiders become invisible attack paths.
We’d love to show you how we’re tackling this at Token Security, book a demo to chat with our technical team so you can scale without sacrificing safety.
Sponsored and written by Token Security .
FortiBleed leak exposes Fortinet VPN credentials for 73,000 devices.
Steam Workshop abused to spread malware via Wallpaper Engine app
Ransomware gang abuses Microsoft Teams relays to hide malicious traffic
What does it take to make AI work in production? Find out with Workflow on-demand.
Overdue a password health-check? Audit your Active Directory for free
Don't just map the problem with AI agent sprawl, fix it.
AI is a data-breach time bomb: Read the new report
Microsoft 365 doesn’t fully protect your data. See how Acronis Cyber Protect fills the gap.
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.
