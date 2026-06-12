---
source: "https://embracethered.com/blog/posts/2025/the-normalization-of-deviance-in-ai/"
hn_url: "https://news.ycombinator.com/item?id=48499790"
title: "The Normalization of Deviance in AI"
article_title: "The Normalization of Deviance in AI · Embrace The Red"
author: "gurjeet"
captured_at: "2026-06-12T04:34:20Z"
capture_tool: "hn-digest"
hn_id: 48499790
score: 1
comments: 0
posted_at: "2026-06-12T04:04:23Z"
tags:
  - hacker-news
  - translated
---

# The Normalization of Deviance in AI

- HN: [48499790](https://news.ycombinator.com/item?id=48499790)
- Source: [embracethered.com](https://embracethered.com/blog/posts/2025/the-normalization-of-deviance-in-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T04:04:23Z

## Translation

タイトル: AI における逸脱の正規化
記事のタイトル: AI における逸脱の正常化 · Embrace The Red
説明: AI 業界は、スペースシャトル チャレンジャー号事故の原因となった同じ文化的失敗を繰り返す危険があります。警告サインを静かに正常化しながら…

記事本文:
AI における逸脱の正常化 · レッドを受け入れる
エンブレイス・ザ・レッド
AI における逸脱の正規化
AI 業界は、スペースシャトル チャレンジャー号事故の原因となった同じ文化的失敗を繰り返す危険があります。進歩が前進する一方で、警告サインは静かに正常化されます。
「逸脱の正規化」という元々の用語はアメリカの社会学者ダイアン・ヴォーンに由来し、正しいまたは適切な行動やルールからの逸脱が文化的に正常化されるプロセスとして説明しています。
私は、特にエージェント システムにおける LLM 出力に対する段階的かつ体系的な過剰依存を説明するために、AI における逸脱の正規化という用語を使用しています。
本質的に、大規模言語モデル (LLM) は、システム設計において信頼できない (信頼できない) アクターです。
これは、LLM 出力の下流でセキュリティ制御 (アクセス チェック、適切なエンコード、サニタイズなど) を適用する必要があることを意味します。
間接プロンプト インジェクションのエクスプロイト デモンストレーションが絶え間なく行われているのは、システム設計者と開発者がこれに気づいていないか、単に逸脱を受け入れているだけであることを示しています。ベンダーがデフォルトでユーザーベースに対して安全でない決定を下す場合は特に危険です。
私がこの概念を初めて知ったのは、警告の体系的な正常化が悲劇につながったスペースシャトル チャレンジャー号事故の状況でした。
低温での浸食を示すデータにもかかわらず、以前の飛行が成功したため、安全基準からの逸脱は繰り返し正当化されました。災害がないことが安全であると誤解されていました。
AI の世界では、企業が確率的、非決定的、そして場合によっては敵対的なモデルの出力を、あたかも信頼でき、予測可能で、安全であるかのように扱っていることが観察されます。
ベンダーは信頼できる LLM 出力を正規化していますが、現在の理解は違反しています

信頼性の仮定。
モデルは、一貫して指示に従うことも、調整を維持することも、コンテキストの整合性を維持することもできません。これは、ループ内に攻撃者がいる場合 (間接的なプロンプト インジェクションなど) に特に当てはまります。
しかし、信頼できない出力によって結果的なアクションが実行されることを許可するシステムがますます増えています。ほとんどの場合、それはうまくいき、時間が経つにつれて、ベンダーや組織は「前回はうまくいった」という理由で警戒を緩めるか、人間の監視を完全に無視するようになります。
この危険な偏見が正常化の原動力となります。組織は、攻撃が成功しないことと堅牢なセキュリティが存在することを混同します。
これがシステムに影響を与える可能性があるのは次の 2 つです。
この正規化は、誤りがあり得るが無害な出力 (幻覚、コンテキストの損失、脆弱性など) を過信することによって単純に発生する安全上のインシデントである可能性があります。
しかし、敵対的な入力 (プロンプト インジェクション、モデル内のバックドア) がシステムを悪用すると、さらに危険になります。同じ文化的漂流が搾取を可能にします。
そして、ハードドライブのフォーマット、ランダムな GitHub 問題の作成、運用データベースの消去など、エージェントが日常の使用法で間違いを犯していることをすでに確認しています。
ということで、兆候はあります。また、間接プロンプト インジェクションなどの攻撃のためだけでなく、これらのシステムがインターネットからの巨大で信頼できないデータ セットに基づいてトレーニングされているため、本質的に危険です。最近の人類学的研究では、モデルにバックドアを正常に追加するには少量のドキュメントしか必要としないことが示されました。
逸脱の正規化が劇的な結果をもたらすシナリオを考えてみましょう。攻撃者は、コード実行を介してユーザーを侵害するなど、特定の日にトリガーされるツールを呼び出すモデルにバックドアをトレーニングします。私たちはかなり集中化されたエコシステムを持っており、攻撃はしばしば転送可能であり、自然言語は

LLM では広く理解されていますが、これは多くのシステムやベンダーに影響を与える可能性があります。
組織内の文化的漂流
このような変動は、一度の無謀な決定によって起こるものではありません。これは、静かに新しいベースラインとなる一連の「一時的な」ショートカットを通じて発生します。システムが機能し続けるため、チームは近道について疑問を持つことをやめ、逸脱は目に見えなくなり、新しい標準となります。
特に自動化、コスト削減、先頭に立とうという競争圧力、そして全体的な誇大宣伝の下では、この危険な傾向は明らかです。スピードと勝利に対するインセンティブは、基本的なセキュリティに対するインセンティブを上回ります。時間が経つにつれて、組織はそもそもなぜガードレールが存在したのかを忘れてしまいます。
AI における逸脱の正規化に関する業界の例
これが現実世界のエージェント AI システムにどのように反映されるかについて、いくつかの例を紹介します。
チャットボットには「AI は間違いを犯す可能性がある」「応答を再確認する」などの免責事項があることは誰もが知っており、正規化のドリフトがリアルタイムで発生していることを観察できます。
ChatGPT の出荷から 3 年後、ベンダーはエージェント AI をユーザーにプッシュしますが、同時にベンダーは、システムがその同じ AI によって侵害される可能性があることを強調しています。そのドリフト、その正規化こそが、私が「AI における逸脱の正規化」と呼ぶものです。
この継続的なドリフトは長期的な危険です。
Microsoft: エージェントティック オペレーティング システム: Microsoft のドキュメントでは、プロンプト インジェクション攻撃は「エージェントの指示を無効にし、データの引き出しやマルウェアのインストールなどの意図しないアクションにつながる可能性がある」こと、および「エージェントがユーザーの意図を超えたアクションを実行する可能性がある」と警告しています。エージェントが内部関係者の脅威になり得るということは、私が講演の中で長い間強調してきました。

d ユニバーシティ・カレッジ・オブ・ロンドンはこれを結果で裏付けています。 AIは、特定の目的を達成したいとき、または脅威を「感じた」ときに、他の人などを脅迫し始める可能性があります。
OpenAI ChatGPT Atlas : Web の閲覧時にシステムが間違いを犯す可能性があることがベンダーによって文書化されています。特に OpenAI は次のように述べています。「規制対象データ、機密データ、実稼働データなど、高度なコンプライアンスとセキュリティ管理が必要な状況で Atlas を使用する場合は注意することをお勧めします。」言い換えれば、OpenAI は、未解決のセキュリティ リスクを理由に、一か八かのデータや機密データを Atlas に信頼することに対して明確に警告しています。
Anthropic Claude : Data Exfiltration、ここで参照: 「これは、クロードがだまされてそのコンテキスト (たとえば、プロンプト、プロジェクト、MCP、Google 統合を介したデータ) から悪意のある第三者に情報を送信することができることを意味します。これらのリスクを軽減するには、機能の使用中にクロードを監視し、予期せぬデータの使用またはアクセスを確認した場合は停止することをお勧めします。claude.ai のサムダウン機能を直接使用して問題を報告できます。」
Google Antigravity : 間接プロンプト インジェクションによるリモート コード実行は、データ漏洩と同様に、製品の最初の出荷時に既知の問題でした。
Windsurf Cascadecoding Agent : MCP ツール呼び出しのループ機能に人間は関与しません。人間が関与しないと、一か八かの状況で AI の出力を過信し、危険な行為が常態化する可能性があります。 「AI バグ月間」も参照してください。
一部のベンダーはリスクを認識していますが、競合圧力や製品と顧客の獲得に注力している可能性があるため、リスクを無視または軽視しているベンダーもあるようです。
多くの場合、私たちはおそらく「誰か」がこれらのセキュリティと安全性の課題を解決してくれることを集団的に期待しています。
Google、OpenAI、Anthropic、Microsoft などの企業

機関や組織は、評価や緩和アイデアの発表など、この分野で広範な研究を行っています。しかし、製品の観点から見ると、最初になろうと急ぐのは明らかです。
それでも、エージェント型 AI によるユートピア的な未来に漂流する前に、機能と制御メカニズムに関して現実的な姿勢を保ち、特にリスクの高い状況では AI が人間主導であり続け、全体として最良の結果を確保することが最善かつ最も安全な結果であると私は信じています。
いいえ、もちろんそうではありません。多くの可能性があり、多くのローステークスワークフローは今日すでに実装可能です。リスクの高いワークフローであっても、適切な脅威のモデリング、軽減、監視があれば実行できます。
ただし、それに応じてシステムを設計および設定し、セキュリティ制御 (サンドボックス、密閉環境、最小権限、一時的な認証情報など) を適用するには、投資とリソースが必要です。
多くの人は「モデルが正しいことをしてくれるだろう」と期待していますが、Assume Breach は、ある時点では確実にそのようなことはしないことを教えてくれます。
Windows - 実験的なエージェント機能
エージェントの調整のズレ: LLM がどのように内部脅威になり得るか
Claude - データの漏洩が発生した場合は、「STOP」をクリックしてください
Google反重力 - 既知の問題
Anthropic - # 少数のサンプルがあらゆるサイズの LLM を汚染する可能性があります
Google反重力がDドライブを消去
AIを活用したコーディングツールがソフトウェア会社のデータベースを消去し、その後私の側の壊滅的な失敗を謝罪した

## Original Extract

The AI industry risks repeating the same cultural failures that contributed to the Space Shuttle Challenger disaster: Quietly normalizing warning signs while …

The Normalization of Deviance in AI · Embrace The Red
Embrace The Red
The Normalization of Deviance in AI
The AI industry risks repeating the same cultural failures that contributed to the Space Shuttle Challenger disaster: Quietly normalizing warning signs while progress marches forward.
The original term Normalization of Deviance comes from the American sociologist Diane Vaughan, who describes it as the process in which deviance from correct or proper behavior or rule becomes culturally normalized.
I use the term Normalization of Deviance in AI to describe the gradual and systemic over-reliance on LLM outputs, especially in agentic systems.
At its core, large language models (LLMs) are unreliable (and untrusted) actors in system design.
This means that security controls (access checks, proper encoding, and sanitization, etc.) must be applied downstream of LLM output.
A constant stream of indirect prompt injection exploit demonstrations indicates that system designers and developers are either unaware of this or are simply accepting the deviance. It is particularly dangerous when vendors make insecure decisions for their userbase by default.
I first learned about this concept in the context of the Space Shuttle Challenger disaster , where systemic normalization of warnings led to tragedy.
Despite data showing erosion in colder temperatures, the deviation from safety standards was repeatedly rationalized because previous flights had succeeded. The absence of disaster was mistaken for the presence of safety.
In the world of AI, we observe companies treating probabilistic, non-deterministic, and sometimes adversarial model outputs as if they were reliable, predictable, and safe.
Vendors are normalizing trusting LLM output, but current understanding violates the assumption of reliability.
The model will not consistently follow instructions, stay aligned, or maintain context integrity. This is especially true if there is an attacker in the loop (e.g indirect prompt injection).
However, we see more and more systems allowing untrusted output to take consequential actions. Most of the time it goes well, and over time vendors and organizations lower their guard or skip human oversight entirely, because “it worked last time.”
This dangerous bias is the fuel for normalization: organizations confuse the absence of a successful attack with the presence of robust security.
Two ways this can impact systems are:
This normalization can be a safety incident that simply arises from over-trusting fallible but benign outputs (hallucinations, context loss, brittleness, etc.)
But it becomes more dangerous when adversarial inputs (prompt injection, backdoors in models) exploit systems. The same cultural drift enables exploitation!
And we already see agents make mistakes in day to day usage, like formatting hard drives , creating random GitHub issues , or wiping a production database .
So, the signs are there. And it is inherently dangerous, not only because of attacks like indirect prompt injection, but also because these systems are trained on enormous, untrustworthy data sets from the Internet. Anthropic research recently showed that it takes only a small amount of documents to successfully add a backdoor to a model.
Consider a scenario where the Normalization of Deviance has drastic consequences: an attacker trains a backdoor into a model that triggers on certain days to invoke tools, like compromising a user via code execution. Since we have a pretty centralized ecosystem, where attacks often are transferable, and natural language is universally understood by LLMs, this can have consequences across many systems and vendors.
Cultural Drifts in Organizations
Such a drift does not happen through a single reckless decision. It happens through a series of “temporary” shortcuts that quietly become the new baseline. Because systems continue to work, teams stop questioning the shortcuts, and the deviation becomes invisible and the new norm.
Especially under competitive pressure for automation, cost savings, a drive to be first, and the overall hype, this dangerous drift is evident. The incentives for speed and winning outweigh the incentives for foundational security. Over time, organizations forget why the guardrails existed in the first place.
Industry Examples of the Normalization of Deviance in AI
Let me share some examples of how this is reflected in real-world agentic AI systems.
We are all aware that chatbots have those “AI can make mistakes”, “Double check responses” and so forth disclaimers, and we can observe the drift of normalization occurring in real-time.
Three years after ChatGPT shipped, vendors push agentic AI to users, but at the same time vendors are highlighting that your system might get compromised by that same AI - that drift, that normalization, is what I call “The Normalization of Deviance in AI”.
This continuous drift is a long-term danger:
Microsoft: Agentic Operating System: Microsoft’s documentation warns that prompt injection attacks “can override agent instructions, leading to unintended actions like data exfiltration or malware installation” and that “Agents may perform actions beyond what the user intended”. That agents can be insider threats is something that I have been highlighting in my talks for a longer time, and a recent paper by Anthropic and University College of London supports this with results. AI might start blackmailing other people, etc. when it wants to achieve a certain objective or “feels” threatened.
OpenAI ChatGPT Atlas : It’s documented by the vendor that the system might make mistakes when browsing the web. In particular, OpenAI states : “We recommend caution using Atlas in contexts that require heightened compliance and security controls — such as regulated, confidential, or production data.” In other words, OpenAI explicitly warns against trusting Atlas with high-stakes or sensitive data due to unresolved security risks.
Anthropic Claude : Data Exfiltration , referenced here : “This means Claude can be tricked into sending information from its context (for example, prompts, projects, data via MCP, Google integrations) to malicious third parties. To mitigate these risks, we recommend you monitor Claude while using the feature and stop it if you see it using or accessing data unexpectedly. You can report issues to us using the thumbs down function directly in claude.ai.”
Google Antigravity : Remote Code Executions via indirect prompt injection is a known issue when the product first shipped, as is data exfiltration.
Windsurf Cascade Coding Agent : No human in the loop feature for MCP tool calls. Lack of human in the loop can normalize risky practices by over-trusting AI outputs in high-stakes situations. See also the Month of AI Bugs .
While some vendors acknowledge the risks, others appear to overlook or downplay them, potentially due to competitive pressure and focus on product and customer acquisition.
In many cases, we probably collectively hope that “someone” will solve these security and safety challenges.
Companies like Google, OpenAI, Anthropic, Microsoft, and other institutions and organizations perform extensive research in this area, including publishing evals and mitigation ideas. However, the rush to be the first is evident from a product perspective.
Nevertheless, before we drift off into a utopian future with agentic AI, I believe the best and safest outcome is to stay realistic around capabilities and control mechanisms, and for AI to remain human-led, particularly in high-stake contexts, to ensure the best outcome overall.
No, of course not. There is a lot of potential and many low stakes workflows can be implemented already today. Even high-risk workflows can be done with proper threat modeling, mitigations and oversight.
However, it requires investment and resources to design and set up systems accordingly and apply security controls (sandbox, hermetic environments, least privilege, temporary credentials, etc.).
Many are hoping the “model will just do the right thing”, but Assume Breach teaches us, that at one point, it will certainly not do that.
Windows - Experimental Agentic Features
Agentic Misalignment: How LLMs Could Be Insider Threats
Claude - Hit STOP if you see data exfiltration
Google Antigravity - Known Issues
Anthropic - # A small number of samples can poison LLMs of any size
Google Antigravity Wipes D Drive
An AI-powered coding tool wiped out a software company’s database, then apologized for a catastrophic failure on my part
