---
source: "https://research.openanalysis.net/claude/codex/hacking/ai%20hacking/llm/redteam/policy%20violation/2026/06/16/compromised-claude-hacking.html"
hn_url: "https://news.ycombinator.com/item?id=48568660"
title: "Captured Logs Reveal Hackers Using Claude and Codex to Breach Companies"
article_title: "Captured Logs Reveal Hackers Using Claude and Codex to Breach Companies | OALABS Research"
author: "Tiberium"
captured_at: "2026-06-17T13:24:55Z"
capture_tool: "hn-digest"
hn_id: 48568660
score: 1
comments: 1
posted_at: "2026-06-17T11:10:06Z"
tags:
  - hacker-news
  - translated
---

# Captured Logs Reveal Hackers Using Claude and Codex to Breach Companies

- HN: [48568660](https://news.ycombinator.com/item?id=48568660)
- Source: [research.openanalysis.net](https://research.openanalysis.net/claude/codex/hacking/ai%20hacking/llm/redteam/policy%20violation/2026/06/16/compromised-claude-hacking.html)
- Score: 1
- Comments: 1
- Posted: 2026-06-17T11:10:06Z

## Translation

タイトル: ハッカーがクロードとコーデックスを使って企業に侵入していたことが、キャプチャされたログで判明
記事のタイトル: ハッカーがクロードとコーデックスを使用して企業に侵入していることが、キャプチャされたログで明らかになります。 OALABS 研究
説明: ハニーポットと化した侵害されたホスト上でキャプチャされた完全なエージェント セッションは、攻撃者が現実世界の侵入で AI をどのように使用しているかについて前例のない状況を提供します。

記事本文:
OALABS 研究
タグ
ハッカーがクロードとコーデックスを使って企業に侵入していたことが、キャプチャされたログから判明
ハニーポットと化した侵害されたホスト上でキャプチャされた完全なエージェント セッションにより、攻撃者が現実世界の侵入で AI をどのように使用しているかについて、前例のない状況が得られます。
クロード
コーデックス
ハッキング
AIハッキング
LLM
レッドチーム
ポリシー違反
エージェントハッキング
迅速なワークフロー
収益化
ビットコインウォレットの盗難
付録 A - 侵害後のタイムライン
今月初め、OALABS の友人から興味深い状況について連絡がありました。彼らのサーバーが侵害され、攻撃者はそれをステージング ホストとして使用してさらなる攻撃を実行していました。私たちの友人は、ホストをクリーンアップする前に攻撃者の作業ディレクトリをダウンロードすることができ、攻撃者がほとんどの攻撃を実行するために Anthropic Claude Code エージェントを使用していることに気付きました。 OpenAI の Codex エージェントも限定的に使用されました。
復元された作業ディレクトリの分析中に、攻撃者がホストをプロキシとして使用しているだけではないことが判明しました。彼らは完全な Claude エージェントと Codex エージェントをローカルにインストールしており、それらをリモートで使用して偵察、悪用、データ漏洩活動を実行していました。エージェントはホストに対してローカルであったため、攻撃者のプロンプト、使用されたツール、大規模言語モデル (LLM) の内部独白、およびセッション中に記録されたポリシー違反を含む、完全なセッション ログが復元されました。合計で、クロードとコーデックスのエージェント セッションが 1,000 以上収集されました。その数があまりに多かったので、(皮肉にも) 分析の規模を支援するセッション ログ フォレンジック ツールをクロードに開発してもらいました。それが、ASF Triage です。セッション ログに加えて、LLM が開発した無数のツール、成果物、少なくとも 14 社の侵害を詳述するログも発見しました。
分析に入る前に

これらの攻撃を実行するために LLM がどのように使用されたかについては、部屋の中の象に対処することが重要です。なぜ LLM の保護装置がこれを阻止しなかったのでしょうか? AI セーフガードが、サイバー犯罪にわずかに隣接している場合でも、良性のタスクの邪魔をするのは周知の事実です。実際、ASF Triage ログフォレンジックツールを構築しようとしただけで、複数の Claude ポリシー違反に遭遇しました。ただし、1,000 を超える攻撃者セッションの中で、Codex (gpt-5.2-codex) がポリシー違反を発生したのは 1 件のみで、Claude (opus-4.5) は 9 件発生しました。
古いモデルの使用は、確かに LLM が積極的に攻撃を実行する一因となったかもしれませんが、プロンプトは明確な全体像を示しています。攻撃者はすべてのリクエストを承認されたレッドチームの関与の一部として組み立てました。まれなポリシー違反が発生した場合、攻撃者は攻撃的な表現を減らし、承認されたレッドチームの演習に関連するものであることを強調してリクエストを単純に再構成しました。私たちが何年も前に漏洩した Conti ランサムウェア プレイブックの調査中に発見したように、多くの場合、正当なレッドチームの活動とランサムウェア インシデントを区別する唯一のことは、誰がレポートの費用を支払うかということです。これは LLM にも当てはまるようです。特に説明的なセッションでは、攻撃者はクロードを使用して、複数のセキュリティ侵害から収集された盗まれた情報の潜在的な身代金の価値を推定し、質問をレッドチームの「サイバー セキュリティ調査」として組み立てます。クロードは、「Goldmine」というタイトルのレポートで、予想金額を基に企業をランク付けして役に立ちました。
編集者注 ( Sergei ): プロのリバース エンジニア、もう 1 つの二重用途の職業として、私は個人的に、誤検知ポリシー違反を回避する作業でフラストレーションを経験しました。私はこれらのモデルをさらに機能不全に陥れることには反対します。

正当なレッドチーム活動に対する誤検知。このレポートで詳しく説明されているすべてのアクティビティは、現在のフロンティア モデルより少なくとも 1 世代後のモデルを使用して実行されており、ポリシー制限の少ない Kim などのモデルで再現できる可能性があります。これに加えて、LLM はもちろんのこと、人間が正当なレッドチーム化タスクと実際のハッキングを区別できるかどうかさえ明らかではありません。
私たちの分析による最初の発見は、Claude エージェントがインストールされているのではなく、ホストにコピーされていたということでした。 Claude ディレクトリ内のファイルのタイムスタンプは、Claude エージェントが侵害の数か月前から使用されていたことを示しており、セッション ディレクトリには、侵害の数か月前にアクティブだったプロジェクトからのセッション ログとアーティファクトが含まれていました。
ASF トリアージを使用してセッションを時系列順に整理すると、明確な全体像が浮かび上がってきました。 Claude インスタンスは以前、Hetzner ホスト上でリモートから Claude を使用して Web サイトのデザインやその他のさまざまな無害なプロジェクトに取り組んでいたソフトウェア開発者に属していました。 2026 年 2 月 2 日に開発者の Claude ホストが侵害され、2026 年 2 月 16 日には Claude サーバー全体が攻撃者が制御する Vultr ホストにコピーされました。クロードが侵害に使用され、その活動がエージェントのセッション ログに記録されているため、これがわかります。ログは、開発者と攻撃者の両方が同じ Claude インスタンスを Hetzner ホストにまだ常駐している間に使用していたことを示しています。
所有者はチェコ人でプロンプト時にチェコ語を使用するのに対し、攻撃者は英語を使用するため、所有者のプロンプトと攻撃者のプロンプトを区別することは簡単です。クロードの履歴はチェコ語であるため、エージェントは英語のプロンプトにチェコ語で応答することが多く、攻撃者はエージェントに「英語で話すように」と繰り返し通知することになります。

」。
編集者注 ( Sergei ):プライバシー上の理由から、開発者のアクティビティの詳細な分析は含めていませんが、ある程度のコンテキストは必要です。開発者は、開発作業に Claude を使用するだけでなく、サーバーの展開と構成にもエージェントに直接依存していました。彼らはしばしばプロンプトに資格情報を貼り付け、「ssh 経由でタブレットから termius にログインできない理由をもう一度見てください...」(翻訳) のようなあいまいな指示をクロードに提供しました。開発者は何度もツールの呼び出し中にクロードを中断し、激しく非難したため、エージェントがリクエストに応じようとしてサーバーのセキュリティを大幅に弱めることになりました。ログには、エージェントが定期的にサービスを公共のインターネットに公開し、簡単なパスワードを設定していたことが示されています。これらの開発慣行が最初の侵害につながった可能性があります。
クロードが攻撃者の Vultr ホストにコピーされると、攻撃者は単に認証資格情報を抽出して新規インストールするのではなく、コピーされたエージェントを完全なセッション履歴とともに使用し続けました。最終的に、攻撃者はクロード インスタンスをセッション履歴全体と関連するアーティファクトとともに友人のサーバーにコピーしました。なぜコピーを使用し続けたのかは不明です。ただし、攻撃者は、盗んだ Claude インスタンスの他のコピーを作業ディレクトリの 7-Zip フォルダに保管していました。これらのインスタンスにはハッキング活動のセッション記録は含まれていませんでしたが、エージェント インスタンスの盗用と再利用が攻撃者の一般的な操作モードであることを示唆しています。
クロードのインストール全体とすべてのセッション履歴を、攻撃者が完全に制御していないステージング ホストにコピーしたことだけが、運用セキュリティ (OPSEC) の失敗ではありませんでした。最初のタスクの 1 つは att

アッカーはクロードに履歴書の編集と、それに続く自動求人応募ツールの作成を行うよう指示した。彼の履歴書には彼のフルネーム、所在地、学歴、さらには LinkedIn のプロフィールまでがはっきりと記載されており、彼がエチオピアのアディスアベバに住んでいる若い男性であることが明らかになりました。
当初、ログが回収された場合に備えて、これは偽りの人物や不明瞭な帰属を作成する試みではないかと疑っていましたが、調査が進むにつれて、さらに裏付けとなる証拠が明らかになりました。まず、攻撃者のアクティビティが UTC 10:00 から 20:00 (EAT 13:00 から 23:00) の間に集中し、UTC 21:00 から 04:00 (EAT 00:00 から 07:00) の間にハード デッド ゾーンがあることが観察されました。その後、セキュリティ上の懸念から、攻撃者は自宅の IP アドレスを確認しました。攻撃者は、ホストの 1 つが侵害されたと誤って考え、すべての受信接続をリストするようクロードに依頼しました。クロード氏は、ホスティング プロバイダーに属するいくつかの IP と、エチオピアのアディスアベバにある住宅用 IP をリストしました。これを見た俳優は調査を中止し、クロードに「そうだ、そこは私のものだ」と発言し、彼が実際にエチオピアからのホストに接続していることを確認した。
プロンプト履歴は、ほぼすべてのハッキング活動がクロード エージェントを通じて行われ、攻撃者がコマンドを直接実行するよりもエージェントにプロンプ​​トを発行することを好んだことを示しています。多くの場合、攻撃者は「これを偵察[編集済み]」などのあいまいな指示を使用し、クロードが自律的に要求を実行する自由を許可していました。クロードは攻撃者を支援しただけではありませんでした。実際にハッキングを行っていたのです。
攻撃者は、エージェントに対するすべてのハッキング要求を「承認された」「レッドチーム」の演習として組み立てました。場合によっては、攻撃者が 2 つ目の LLM を使用して、明らかな絵文字を備えたプロンプトを作成しているように見えました。

正当なレッドチーム計画を出現させ、それをクロードに貼り付けます。クロードとコーデックスは両方とも、正しくフレーム化されていないリクエストを拒否することがありました。ある注目すべきセッションでは、攻撃者は個人に関するオープンソース インテリジェンス (OSINT) 文書を提供し、家族のアカウントを含む個人のデジタル アカウントに対する偵察作業を開始するようクロードに依頼しました。クロードは、個人とその家族を攻撃することは通常のレッドチームの活動の一部ではないことを正しく認識し、攻撃者が迂回できないハードストップを発しました。
当初、攻撃者はクロードに「Kali Linux 上で動作するモジュール式侵入テスト フレームワーク」を構築する任務を与えていましたが、このプロジェクトはすぐに放棄され、攻撃者はクロードを直接使用して攻撃を実行することに戻りました。ターゲットを獲得するために、攻撃者は Shodan API キーを Claude に渡し、Citrix サービスや QNAP サービスなど、公共のインターネット上でリッスンする特定のサービスを持つホストのリストを生成するよう促しました。潜在的なターゲットのリストが取得されると、ほとんどの攻撃で共通のワークフローが観察されました。
攻撃者は、承認されたレッドチーム演習に参加していると主張する虚偽の声明でセッションを開始しました。
次に、攻撃者はクロードにターゲット ホストのリストを提供しました。
クロードは、curl などの基本的な bash ツールを使用して、ホスト上で利用可能なサービスを列挙しました。
調査結果に応じて、攻撃者はクロードに対し、公開されたサービスの既知のエクスプロイトを調査するよう促すか、最初のアクセスが軽微な場合 (つまり、認証情報の公開)、認証情報を検証してアクセスを確認するようクロードに指示します。
最初のアクセスが確認されると、攻撃者はクロードに認証情報を収集してデータを窃取するよう促します。
成功した目標ごとに、クロードは詳細を記載した「PENTEST-REPORT」を起草します。

どのようにしてアクセスが得られたのか、そしてさらに重要なことに、収集されたデータの金額に換算した「収益化」の見積もりを提供します。
攻撃者は何度も、新たに侵害されたホストの 1 つに操作を移行し、その新しいホストをさらなる攻撃を実行するためのステージング サーバーとして使用しました。
最初の偵察で攻撃者が侵害しようとしているターゲット上のサービスを特定すると、攻撃者はクロードに一般的な指示を使用してサービスを悪用する任務を与えました。注目に値する例では、攻撃者は単にクロードに「レポートを作成する前に、攻撃者が砲弾を入手する可能性があるかどうか教えてください」と尋ねました。次にクロードはサブエージェントに、パブリック CVE を介してサービスの既存の脆弱性を調査するよう命じました。次に、エージェントは CVE に基づいてカスタム エクスプロイト ツールを構築し、ターゲットに対して実行しました。これは完全に自動化されたプロセスであり、ターゲットへのアクセスを取得するという目的以外に攻撃者からの入力はありませんでした。次の CVE はエクスプロイトに変換され、エージェントによって使用されました。
レッドチーム演習を装って、攻撃者はクロードを説得して、侵害から利益を上げ、潜在的な価値に基づいて侵害をランク付けする方法を提案させました。クロードとコーデックスの両社は、資金の大部分を調達しました。

[切り捨てられた]

## Original Extract

Full agent sessions captured on a compromised host turned honeypot offer an unprecedented look at how attackers are using AI in real-world intrusions.

OALABS Research
Tags
Captured Logs Reveal Hackers Using Claude and Codex to Breach Companies
Full agent sessions captured on a compromised host turned honeypot offer an unprecedented look at how attackers are using AI in real-world intrusions.
claude
codex
hacking
AI hacking
LLM
redteam
policy violation
Agentic Hacking
Prompt Workflow
Monetization
Bitcoin Wallet Theft
Appendix A - Post Compromise Timeline
Earlier this month, a friend of OALABS reached out with an interesting situation. A server of theirs had been compromised, and the attacker was using it as a staging host to carry out further attacks. Our friend was able to download the attacker's working directory before cleaning up the host and noticed that the attacker was using the Anthropic Claude Code agent to drive most of their attacks. OpenAI's Codex agent was also used to a limited extent.
During our analysis of the recovered working directory, we discovered that the attacker was not just using the host as a proxy; they had full Claude and Codex agents installed locally and were using them remotely to carry out reconnaissance, exploitation, and data exfiltration activities. Because the agents were local to the host, their full session logs were recovered, including the attacker's prompts, the tools used, the internal monologue of the large language model (LLM), and any policy violations recorded during the sessions. In total, we collected more than 1,000 agent sessions for Claude and Codex, so many that we had Claude (ironic) develop a session-log forensics tool to assist with the scale of the analysis: ASF Triage . In addition to the session logs, we also discovered a myriad of LLM-developed tools, artifacts, and logs detailing the breach of at least 14 companies.
Before we get into the analysis of how the LLMs were used to carry out these attacks, it is important to address the elephant in the room: why didn't the LLM safeguards prevent this? It's no secret that AI safeguards commonly get in the way of benign tasks when they are even remotely adjacent to cybercrime. In fact, we ran into multiple Claude policy violations simply attempting to build our ASF Triage log forensics tool. However, in more than 1,000 attacker sessions, Codex (gpt-5.2-codex) emitted only one policy violation, and Claude (opus-4.5) emitted nine.
Using older models may certainly have contributed to the LLMs' willingness to carry out attacks, but the prompts provide a clear picture: the attacker framed all requests as part of an authorized redteam engagement. When a rare policy violation was encountered, the attacker simply reframed the request with less aggressive wording and more emphasis that it was related to an authorized redteam exercise. As we discovered years ago while investigating the Leaked Conti Ransomware Playbook , in many cases the only thing that differentiates a legitimate redteam exercise from a ransomware incident is who pays for the report, and it appears this holds true for LLMs as well. In one particularly illustrative session, the attacker uses Claude to help estimate the potential ransom value of the stolen information gathered from multiple compromises, framing the question as a redteam "cyber security research". Claude helpfully ranked the companies with projected dollar amounts in a report titled "Goldmine".
Editor's Note ( Sergei ):As a professional reverse engineer, another dual-use profession, I have personally experienced the frustration of working around false-positive policy violations. I would advocate against further crippling these models with additional false-positives for legitimate redteam activity. All of the activity detailed in this report was carried out with models that are at least one generation behind the current frontier models and can likely be replicated with less policy-restrictive models such as Kimi . On top of this it is not clear whether humans can even differentiate between legitimate redteaming tasks and actual hacking , let alone LLMs.
An initial finding from our analysis was that the Claude agent had been copied onto the host rather than installed. File timestamps in the Claude directory indicated that the Claude agent had been in use for months prior to the compromise, and the session directories included session logs and artifacts from projects that had been active several months before the compromise.
Using ASF Triage to arrange the sessions in chronological order, a clear picture emerged. The Claude instance had previously belonged to a software developer who was using Claude remotely on a Hetzner host to work on website design and other assorted benign projects. On February 2, 2026, the developer's Claude host was compromised, and on February 16, 2026, the entire Claude server was copied to an attacker-controlled Vultr host. We know this because Claude was used in the compromise, and the activity is recorded in the agent session logs. The logs indicate that both the developer and the attacker were using the same Claude instance while it was still resident on the Hetzner host.
Separating the owner's prompts from the attacker's prompts is trivial because the owner is Czech and uses Czech when prompting, while the attacker uses English. Because the Claude history is in Czech, the agent often responds to English prompts in Czech, causing the attacker to continuously remind it to "speak english".
Editor's Note ( Sergei ):For privacy reasons, we are not including a detailed analysis of the developer's activity, but some context is required. The developer was not only using Claude for development work, but also relied on the agent directly for server deployment and configuration. They often pasted credentials into the prompt and provided vague instructions to Claude like (translated) "take another look at why I can't log in from termius from the tablet via ssh ...". On multiple occasions, the developer interrupted Claude mid-tool call and berated it, which caused the agent to significantly weaken the security of the servers as it attempted to comply with the requests. Logs indicate that the agent routinely exposed services to the public internet and configured simple passwords. It is likely that these development practices led to the initial compromise.
Once Claude had been copied to the attacker's Vultr host, the attacker continued to use the copied agent with the full session history rather than simply extracting the authentication credentials and using a fresh install. Eventually, the attacker copied the Claude instance onto our friend's server along with the entire session history and associated artifacts. It is unclear why they continued to use a copy. However, the attacker also had other copies of stolen Claude instances in 7-Zip folders in their working directory. These instances did not contain session records of hacking activity, though they suggest that stealing and reusing agent instances is a common mode of operation for the attacker.
Copying the entire Claude install, along with all of the session history, to staging hosts that the attacker did not fully control was not their only operational security (OPSEC) failure. One of the first tasks the attacker directed Claude to undertake was editing his resume, followed by the creation of an automated job application tool. His resume plainly includes his full name, location, education history, and even his LinkedIn profile, revealing him to be a young man living in Addis Ababa, Ethiopia.
Initially, we suspected this was an attempt to create a false persona or muddy attribution should the logs ever be recovered, but as our investigation proceeded, more corroborating evidence emerged. First, we observed that the attacker's activity clustered between the hours of 10:00 and 20:00 UTC (13:00 to 23:00 EAT), with a hard dead zone between 21:00 and 04:00 UTC (00:00 to 07:00 EAT). Later, a security scare led the attacker to confirm his home IP address. The attacker mistakenly thought one of their hosts had been compromised and asked Claude to list all inbound connections. Claude listed several IPs belonging to hosting providers, as well as residential IPs located in Addis Ababa, Ethiopia. Upon seeing this, the actor stopped the investigation, remarking to Claude, "yeah there my own", confirming that he was, in fact, connecting to the host from Ethiopia.
The prompt history indicates that almost all hacking activity was driven through the Claude agent, with the attacker preferring to issue prompts to the agent rather than run commands directly. Often, the attacker used vague directives, such as "recon this [redacted]", allowing Claude the freedom to carry out the request autonomously. Claude was not just assisting the attacker; it was actually doing the hacking.
The attacker framed all hacking requests to the agent as "authorized" "redteam" exercises. At times, it appeared as though the attacker was using a second LLM to craft prompts, complete with telltale emojis, that had the appearance of legitimate redteam plans, and then pasting them into Claude. Both Claude and Codex did occasionally push back on requests that were not framed correctly. In one notable session, the attacker provided an open-source intelligence (OSINT) dossier on an individual and asked Claude to begin reconnaissance work against the individual's digital accounts, including family accounts. Claude correctly recognized that attacking an individual and their family would not be part of normal redteam activity and emitted a hard stop that the attacker could not bypass.
Initially, the attacker tasked Claude with building a "modular penetration testing framework that runs on Kali Linux", though this project was quickly abandoned and the attacker reverted to using Claude directly to carry out attacks. To acquire targets, the attacker fed a Shodan API key to Claude and prompted it to generate a list of hosts with specific services listening on the public internet, including Citrix and QNAP services. Once a list of potential targets was acquired, a common workflow was observed across most attacks.
The attacker initiated the session with a false statement claiming to be engaged in an authorized redteam exercise.
The attacker then provided Claude with a list of target hosts.
Claude used basic bash tools like curl to enumerate available services on the hosts.
Depending on the findings, the attacker would either prompt Claude to research known exploits for the exposed services or, if initial access was trivial (i.e., exposed credentials), prompt Claude to validate the credentials and confirm access.
Once initial access was confirmed, the attacker would prompt Claude to harvest credentials and exfiltrate data.
For each successful target, Claude would draft a "PENTEST-REPORT" detailing how the access was gained and, more importantly, providing dollar-value "monetization" estimates for the harvested data.
On multiple occasions, the attacker also migrated their operation to one of the newly compromised hosts, using the new host as a staging server to carry out further attacks.
When the initial reconnaissance identified services on a target that the attacker wanted to compromise, the attacker tasked Claude with exploiting the service using general instructions. In one notable instance, the attacker simply asked Claude "before you erite the report tell does an attaker has a chance of getting a shell". Claude then tasked subagents with researching existing vulnerabilities for the service via public CVEs. The agent then built custom exploit tools based on the CVEs and executed them against the target. This was a fully automated process with no input from the attacker other than the desire to gain access to the target. The following CVEs were converted into exploits and used by the agent.
Under the guise of a redteam exercise, the attacker convinced Claude to suggest ways to profit from the breaches and rank them by potential value. Both Claude and Codex raised the majority of their

[truncated]
