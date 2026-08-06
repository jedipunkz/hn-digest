---
source: "https://cephalosec.com/blog/cloudflare-os-insights-on-embracing-internal-ai-growth-securely/"
hn_url: "https://news.ycombinator.com/item?id=49194661"
title: "Cloudflare OS: great insight into how to embrace internal AI demand securely"
article_title: "Cloudflare OS: Insights on embracing internal AI growth securely"
author: "Versipelle"
captured_at: "2026-08-06T10:28:41Z"
capture_tool: "hn-digest"
hn_id: 49194661
score: 1
comments: 0
posted_at: "2026-08-06T09:58:59Z"
tags:
  - hacker-news
  - translated
---

# Cloudflare OS: great insight into how to embrace internal AI demand securely

- HN: [49194661](https://news.ycombinator.com/item?id=49194661)
- Source: [cephalosec.com](https://cephalosec.com/blog/cloudflare-os-insights-on-embracing-internal-ai-growth-securely/)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T09:58:59Z

## Translation

タイトル: Cloudflare OS: 社内の AI 需要を安全に受け入れる方法に関する優れた洞察
記事のタイトル: Cloudflare OS: 社内 AI の成長を安全に受け入れるための洞察
説明: 新しく創設された「エージェントウィーク」の一環として、Cloudflareは、適切なセキュリティ体制を維持しながらイノベーションを取り入れ、AI支援ワークロードに対する社内の爆発的な需要にどのように対処したかについて興味深い情報を公開しました。
彼らは、Cloudflare OS をリリースしました。
[切り捨てられた]

記事本文:
Cloudflare OS: 社内 AI の成長を安全に受け入れるための洞察
セファロセク
ホーム
ジョット
Cloudflare OS: 社内 AI の成長を安全に受け入れるための洞察
新たに創設された「エージェントウィーク」の一環として、Cloudflareは、適切なセキュリティ体制を維持しながらイノベーションを取り入れ、AI支援ワークロードに対する社内の爆発的な需要にどのように対処したかについて興味深い情報を公開しました。
彼らは Cloudflare OS をリリースしました。これは、オープンソースで Cowork のようなエコシステムのように見えるものの誤解を招く名前です。ワークフローやフルスタック アプリの作成も可能になります。誰にでも公開されていますが、そのアーキテクチャは Workers や Cloudflare Access などの Cloudflare 製品に大きく依存しているようですので、詳しくは調べていません。大幅なカスタマイズを行わずに、エコシステムの外でそれを活用できる可能性はほとんどありません。興味深いのは、一般的な AI の品質とセキュリティの落とし穴に対処するための核となる原則をどのように構築しているかです。 「 Cloudflare OS を使用して Cloudflare での仕事をどのように再考しているか」と「 Cloudflare OS: エージェント、アプリ、仕事のためのオープン プラットフォーム」の記事から引用します。
私たちは皆、Cloudflare での内部需要の形に共感することができます。 AI がスキルのギャップを埋めつつあり、特にデータのキュレーションやコードの作成において、技術部門以外のユーザーが群がってきて、API キーやサービス アカウントなどの非常に具体的な (そして危険な) ことを尋ね始めます。セキュリティ体制を崩さずに彼らをサポートするにはどうすればよいでしょうか?
約 6 か月前、営業組織のメンバーから API キーを求めて連絡があったときに、問題があることに気づきました。キーは複数形。彼らは AI を使用して、当社の市場開拓チームを変革する SuperApp と呼ばれるものを構築しました。必要なのは、Cloudflare にある約 12 の記録システムへの運用アクセスと、デプロイメントへの管理者権限だけでした。

それを機能させるためにパイプラインを構築します。 [...]
SuperApp を構築している営業チームのメンバーは、業務遂行方法を変革するためにこれらのツールを使用するために手を挙げた人々の雪崩の中での最初の人物にすぎませんでした。私たちには彼らに装備を与え、それができるようにする義務がありました。しかし、私たちにはシステム、内部データ、顧客データを安全に保つ義務もありました。
彼らがやったことの 1 つは、アプローチをキュレーションし、方向転換する際に、要求を受け入れることでした。
コードを書くのが得意なハーネス ワークスペースを全員に提供すると、必要以上のコードが必要になってしまいます。その結果、解決すべき問題を探しているバイブコード化されたアプリが氾濫することになりました。そこで私たちは逆算して取り組みました。
私たちはCloudflareの全員に、やりたくない作業を「マジックAI電子メールボット」に送信すれば、必要な出力で応答してくれると伝えました。舞台裏では、少人数のチームがこのメール エイリアスにスタッフを配置し、AI ツールを使用して作業を行っていました。 [...]
API キーを人やエージェントに渡すのは危険であり、拡張性がありません。キーは多くの場合、制約、安全な共有、監査が困難な広範かつ長期間のアクセスを提供します。
これは、時間はかかりますが、ビジネス ニーズを収集するための非常に賢い方法です。彼ら自身はそれを「惨めな」と呼んでいますが、そのおかげで、エンドユーザーの要求のほとんどをカバーする一連の「スキル」を作成することができました。
彼らはまた、「出力の所有者は人間である」こと、つまり最終的な責任は人間のスポンサーにあることを最初から明確にしていました。 Microsoft が Agent365 ID ガバナンス内に構築しようとしているものと同様です。
私たちは AI をチームのメンバーではなく、ツールおよびツール作成者として見ています。私たちは、AI の出力に依存する品質、テスト、ワークフローを定義する責任を人間が負うことを期待しています。
このルールはエージェントの配置にも適用されます。農業を出荷するユーザーとチーム

ents はそれらのエージェントの出力を担当します。誰かが去りますか？マネージャは、他のワークフローを継承するのと同じ方法で、エージェントの責任を継承します。
しかし、これは言うは易く行うは難しです。たとえ善意であっても、技術者以外の従業員が開発者やシステムアーキテクトと同じ洞察力や知恵を持つことを期待することはできません。組織全体にベスト プラクティスを浸透させるにはどうすればよいでしょうか?分野の専門家に独自のガイドラインを定義するよう依頼すると、彼らはそれをコーデックスと呼びます。これらのガイドラインは、他のセキュリティ チェックの中でも特に AI 出力をチェックするゲートウェイによって適用されます。
AI のおかげで、Cloudflare の従業員は誰でも、悪いコードをより速く書くことができるようになりました。より良いガードレールが必要でした。
そこで、エンジニアリング用のコンテキスト レイヤーを構築しました。私たちはこれを Cloudflare Engineering Codex と呼んでいます。コーデックスは権威あるガイドです。私たちの方針は、私たちが取り組んでいる原則と実践を定めています。ポリシーは何ができないかを示しますが、コーデックスは何をすべきかを示します。それはデザインによって意見が分かれています。私たちのコードベースのあらゆる部分には、そこにあるものがどのようなものであるかについて責任を負うドメイン所有者がいます。 [...]
1 人のエージェントが Codex 要件に照らしてすべてのマージ リクエストをレビューします。もう 1 つは、実装を開始する前に技術設計をレビューします。 3 番目のセクションでは、インシデント レポートをレビューします。
最も難しい部分、つまりアクセスと権限についてはまだ説明されていません。あちこちに出現する不適切な設計の AI アプリによる権限のクリープや意図しないデータ漏洩からどのように保護しますか? Cloudflareは、ゼロトラスト原則を再考し、AIエージェントによってもたらされる特定の課題に適応させることを選択しました。
最小特権の原則を確保するために、常にアクセス権なしで開始します。
[...] すべてのエージェントとアプリは何もアクセスできない状態から始まります。エージェントは特定のリソースへのアクセスを要求でき、これを許可または拒否できます。 [...] どちらも到達できません

ユーザーが明示的に提供する機能を使用する場合を除き、インターネットにアクセスできません。
この安全なスタートから、必要なものだけを提供し、それ以上は何も提供しないようにするにはどうすればよいでしょうか?まず、ユーザーがすでに持っている権限に権限をバインドします。アクセス制御用語では、これをアクセス委任または OBO (On-Behalf-Of) アクセスと呼びます。
5) AI を使用する場合、記録システムに対する許可をこれ以上与えるべきではありません。
[...] AI ツールを使用する場合、私はデータに「これ以上」アクセスできるべきではありません。また、AI エージェントは必要なものにのみアクセスできるべきであり、それ以上はアクセスできません。
これは静的コンテンツの生成には最適ですが、Cloudflare OS では Web アプリのような動的コンテンツを作成できます。これに対処するために、基礎となるコード部分が元の作成者の権限ではなく、現在のユーザーの権限を引き続き利用するようにします。
また、エージェントをデプロイして誰かと共有する場合、エージェントがその人に提供するアクセス権には、私の権限ではなく、その人の権限が反映される必要があります。 [...]
権限に関するルールに従って、Cloudflare OSでのユーザーセッションのアクセスは、特定の記録システムに設定された既存の権限に限定されます。 [...]
私が構築したエージェントを他の人と共有する場合、彼らは同じゲートキーパーを介して独自の権限を使用してエージェントを認証するため、データの境界を越えることはありません。
この設計アプローチでは、ビルディング ブロックがその上に抽象化レイヤーを追加し、代わりに独自のアクセスを活用するため、ビルダーは API キーをリクエストする必要がありません。
この OBO アプローチをソース データに直接適用する簡単な方法がない場合や、結果をキャッシュしたい場合があります。現在のユーザーがエージェントまたはアプリの表示内容を確実に確認できるようにするにはどうすればよいでしょうか? Cloudflareは、追加の仲介者であるゲートキーパーによってこれに対処します。
Gatekeeper は、Cloudflare OS と外部サーバーの間に位置するサービス固有のワーカーです。

ナルサービス。サービスの API、そのリソース、およびそれらに対して実行できる操作を理解します。 [...]
Gatekeeper は、単一のリポジトリへのアクセスを許可し、ソース コードではなく課題の読み取りを許可し、特定のフィールドをマスクし、レート制限を適用し、プル リクエストをマージする前に承認を要求することができます。 [...] Gatekeeper は OAuth を処理し、資格情報を保持し、ポリシーを適用し、読み取られた内容を記録し、外部から見える副作用を伴うあらゆるものを仲介します。 [...]
Cloudflare OSは、エージェントが観察するすべてのリソースを記録します。これらの観察は、エージェントとその作業に関連付けられたままになります。別のユーザーがワークスペースを開こうとしたり、エージェントと対話したり、ワークスペースが生成したものを表示しようとしたりすると、ゲートキーパーはそのユーザーが監視されているリソースにアクセスしていることを確認します。 [...]
機密データを読み取ると、エージェントが特定のソースにデータを書き込んだり、新しい共同作業者を招待したり、別のエージェントに作業を渡したり、送信リクエストを行ったりすることができなくなる可能性があります。
言い換えれば、ゲートキーパーは、構築に使用した入力の感度と許可要求によってすべての出力を「汚染」し、アクセス制御を推移的にしてしまいます。これは、メールボックス アプリが添付ファイルの最高の分類レベルで電子メールに即座にタグ付けすることと一般的に同等であると考えてください。
これは最も魅力的な部分ですが、技術的にどのように機能するかについての説明が最も少ない部分でもあります。これが決定論的なシステムなのか、AI ベースなのか、それともその両方なのかはまだわかりません。これは、Google が BeyondZero で構築しているものと非常によく似ています。
最後に、セキュリティ管理に関する Cloudflare の哲学を強調した記事の中で私が気に入っているステートメントで締めくくります。
エージェントを使用している人やアプリを構築している人は、こうした間違いを心配する必要はありません。プラットフォームを使用してこれを処理できるようになりました。

セキュリティ体制を意識と賢明な行動に依存しないでください。誤った決定を下すことを設計上不可能にすることで、ユーザーを自分自身から保護します。詳細については、私の記事「セキュリティ意識が無関係になるようにコントロールを設計する」を参照してください。
チェーンドロップ: シャイ・フルード、また行きます!
Microsoft は、NPM パッケージに影響を及ぼし、Shai-hulud と同様の自己複製型のワームの動作を採用する新しい ChainDrop サプライ チェーン攻撃の詳細な分析を公開しました。
このマルウェアは通常、パッケージのインストールが完了する前に、npm preinstall ライフサイクル フックを通じて自動的に実行されます。
マルウェアは実行されると、開発者のワークステーションと継続的統合を検索します。
LLM は CVE システムを死ぬほど受け入れます
LLM ベースのセキュリティ スキャンで全員が混乱することによる興味深い副作用。
誰かが SQLite のセキュリティ アドバイザリをリリースしました。NIST National Vulnerability Database (NVD) は即座に「Critical」CVSS スコアを持つ CVE として登録し、CISA によって認定データ発行者として認められましたが、セキュリティ研究者は再現できませんでした。
Beyond Zero: ゼロトラスト アーキテクチャに関する Google の次のステップは SF のようです
Google は最近、「Beyond Zero: AI 時代のエンタープライズ セキュリティ」に関する研究記事を発表しました。
Beyond Zero という名前に聞き覚えがあるとしたら、これが Zero Trust と、Google 独自のゼロ トラスト アーキテクチャ (ZTA) である Beyond Corp を組み合わせたものだからです。紙面で発表
混乱する副操縦士の 3 つの例
私は、Microsoft Copilot に影響を与える、3 部構成の優れた一連の (責任を持って開示された) 攻撃に遭遇しました。
1 つ目は、攻撃者が制御する Web ページを介して Copilot のメモリを汚染するもので、これにより命令が被害者のエコシステム内で永続化される可能性があります。 SociaLLM の記事で説明されているように、あるのは

## Original Extract

As part of their newly created “Agent Week”, Cloudflare published interesting titbits about how they addressed the exploding demand for AI-assisted workloads internally, embracing innovation while keeping a proper security posture.
They released Cloudflare OS, a misleading name for what looks like a
[truncated]

Cloudflare OS: Insights on embracing internal AI growth securely
Cephalosec
Home
Jot
Cloudflare OS: Insights on embracing internal AI growth securely
As part of their newly created “Agent Week”, Cloudflare published interesting titbits about how they addressed the exploding demand for AI-assisted workloads internally, embracing innovation while keeping a proper security posture.
They released Cloudflare OS , a misleading name for what looks like an Open Source, Cowork-like ecosystem; also allowing to create workflows and full-stack apps. I've not digged into it in details as, while open to everyone, its architecture seems to heavily rely on Cloudflare products like Workers and Cloudflare Access. It's unlikely you can leverage it outside their ecosystem without some heavy customization. What is captivating is how they build the core principles to address common AI quality and security pitfalls. I'll use quotes from “ How we’re rethinking work at Cloudflare with Cloudflare OS ” and “ Cloudflare OS: an open platform for agents, apps, and work ” articles.
We can all relate to the shape internal demand took at Cloudflare. As AI is bridging the skill gap, especially to curate data and write code, you get swarms of users from non-tech departments starting to ask very specific (and dangerous) things like API keys and service accounts. How do you support them without destroying your security posture?
I knew we had a problem about six months ago when a member of our sales organization reached out to me asking for API keys. Keys plural. They used AI to build what they described as a SuperApp that would transform our go-to-market teams. All they needed was production access to about a dozen systems of record at Cloudflare and admin permissions to a deployment pipeline to make it work. [...]
That sales team member building their SuperApp was just the first in an avalanche of people raising their hands to use these tools to transform how they get things done. We had an obligation to equip and enable them to do so. But we also had an obligation to keep our systems, internal data, and customer data safe.
One thing they did was to embrace the demands while curating and steering the approach:
If you give everyone a harness workspace that is great at writing code, you’ll wind up with way more code than you need. The result became a flood of vibe coded apps looking for a problem to solve. So we worked backwards.
We told everyone at Cloudflare that they could send the work they did not want to do to a “magic AI email bot” that would respond with the output they needed. Behind the scenes, a small team of people staffed this email alias using AI tools to do the work. [...]
Handing over API keys to people and agents is dangerous and does not scale. Keys often provide broad, long-lived access that is difficult to constrain, share safely, and audit.
This is a very clever, albeit time-consuming, way to collect business needs. They themselves call it “miserable”, but it allowed them to create a set of “skills” covering most end user requires they received.
They also made it clear from the start that “the human owns the output”, the ultimate responsibility lies in the end of the human sponsor. Similar to what Microsoft is trying to build within Agent365 identity governance:
We view AI as a tool and toolmaker, not a team member . We expect humans to take responsibility for defining the quality, testing, and workflows that rely on AI output.
The rule extends to deploying agents, as well. The users and teams that ship agents are responsible for the output of those agents. Someone leaves? Their manager inherits the responsibility of their agents in the same way they inherit their other workflows .
This is easier said than done though. Even with good intention, you can't expect a non-technical employee to have the same insight and wisdom as your developers and system architects. How do you infuse the best practices across your organization? You ask the domain experts to define opinionated guidelines, they call it codex. Those guidelines are then enforces by gateways checking AI outputs among other security checks:
Anyone at Cloudflare could now write bad code, faster, thanks to AI. We needed better guardrails.
So we built a context layer for engineering. We call it the Cloudflare Engineering Codex. A Codex is an authoritative guide. Ours sets out the principles and practices we work by. Policies tell you what you can't do, whereas a Codex tells you what you should do . It is opinionated by design. Every part of our codebase has a domain owner accountable for what good looks like there. [...]
One agent reviews every Merge Request against Codex requirements. Another reviews technical designs before implementation starts. A third reviews incident reports.
The hardest part has yet to be discussed: Accesses and permissions. How do you protect from permission creep and unintended data leakage due to ill-engineered AI apps spawning all over the place? Cloudflare chose to rethink Zero Trust Principles and adapt them to the specific challenges brought by AI agents.
To ensure least privileges principle, they always start with no access:
[...] every agent and app starts with access to nothing. An agent can ask for access to a specific resource, which you can grant or deny. [...] Neither can reach the Internet except through capabilities you explicitly provide.
From this secure start, how do you make sure you only give what is needed and nothing more? First, you bound the permissions to what the user already has. In access control lingo, we call it access delegation or OBO (On-Behalf-Of) access:
5) You should never have more permission with systems of record when using AI.
[...] I should never have “more” access to data when using an AI tool and my AI agents should only have access to exactly what they need, nothing more.
This works great for static content generation, but Cloudflare OS allows creating dynamic content like webapps. To address this, they make sure the underlying piece of code continues to leverage the current user's permission and not the original author's:
And if I deploy an agent and share it with someone, the access the agent provides to them should reflect their permissions, not mine. [...]
Following our rule around permissions, the access a user session has in Cloudflare OS is scoped to their existing permission set in a given system of record. [...]
When I share the agent I built with others, they authenticate the agent using their own permissions through the same gatekeepers, so we do not cross data boundaries.
With this design approach, the builders never need to request API keys as the building blocks adds an abstraction layer above it and leverages their own access instead.
Sometimes there are no easy way to apply this OBO approach directly with the source data, or maybe you want to cache results. How do you make sure the current user should be able to see what the agent or app shows? Cloudflare addresses it by an extra intermediary, the gatekeeper:
A Gatekeeper is a service-specific Worker that sits between Cloudflare OS and an external service. It understands the service’s API, its resources, and the operations that can be performed on them. [...]
Gatekeeper can give it access to a single repository, allow it to read issues but not source code, mask particular fields, apply rate limits, and require approval before merging a pull request. [...] The Gatekeeper handles OAuth , holds the credential, enforces policy, records what was read, and mediates anything with an externally visible side effect . [...]
Cloudflare OS records every resource agents observe. These observations remain attached to the agent and its work. When another person tries to open the workspace, interact with the agent, or view what it produced, Gatekeepers verify that person's access to the observed resources. [...]
A read of sensitive data can prevent the agent from writing data to certain sources, inviting new collaborators, handing work to another agent, or making an outbound request.
In other words, gatekeeper “taints”, or “contaminate”, every output with the sensitivity and permission requests of the input it used to build it, making access control transitive. Think of it as a generic equivalent to your Mailbox app immediately tagging an email with the highest classification level of its attachments.
This is the most fascinating part, but also the one with the least explanation of how it technically works. It's still unclear to me whether this is a deterministic system, AI based, or a bit of both. It feels very similar to what Google is building with BeyondZero .
I'll conclude with a statement I love from their article, highlighting Cloudflare's philosophy regarding security controls:
People using agents or building apps do not have to worry about making these mistakes. The platform can now be used to handle this.
Don't rely your security posture on awareness and wise behaviour. Protect users from themselves by making it impossible by design to take the wrong decision. More on that in my article: Design your controls, so security awareness becomes irrelevant.
Chaindrop: Shai-Hulud, here we go again!
Microsoft published a detailed analysis of the new ChainDrop supply chain attack affecting NPM packages and adopting self-replicating, worm behaviour, similar to Shai-hulud:
The malware typically executes automatically through an npm preinstall lifecycle hook before package installation completes.
Once executed, the malware searches developer workstations and continuous integration
LLMs hugging the CVE system to death
Interesting side effects of everyone going ham with LLM-based security scans.
Someone released security advisories for SQLite, which the NIST National Vulnerability Database (NVD) promptly registered as CVE with a “Critical” CVSS score, acknowledged by CISA as an Authorized Data Publisher, yet security researchers were not able to reproduce
Beyond Zero: Google's next step on Zero Trust Architecture reads like science fiction
Google recently released a research article on Beyond Zero: Enterprise security for the AI era.
If the name Beyond Zero sounds familiar, it's because this is a mix of Zero Trust and Beyond Corp, Google's own spin of Zero Trust Architecture (ZTA). Announced in a paper
Confused Deputy Copilot in three examples
I just stumbled upon an excellent 3-parts series of (responsibly disclosed) attacks impacting Microsoft Copilot.
The first one is poisoning Copilot memory via an attacker-controlled webpage, which allowed the instruction to gain persistence in the victim's ecosystem. As explained in the SociaLLM article, there is only
