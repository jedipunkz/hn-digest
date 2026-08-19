---
source: "https://plugos.net/blog/2026/08/confidential-ai-how-plugclaw-keeps-your-ai-work-private/"
hn_url: "https://news.ycombinator.com/item?id=49359714"
title: "Confidential AI: Let AI Work for You Without Handing over Your Privacy"
article_title: "Confidential AI: How PlugClaw Protects Your AI Work"
image: "https://plugos.net/assets/static/favicon-192x192.GgSX_ec9.webp"
author: "PlugMate"
captured_at: "2026-08-19T11:16:31Z"
capture_tool: "hn-digest"
hn_id: 49359714
score: 2
comments: 0
posted_at: "2026-08-19T10:47:57Z"
tags:
  - hacker-news
  - translated
---

# Confidential AI: Let AI Work for You Without Handing over Your Privacy

- HN: [49359714](https://news.ycombinator.com/item?id=49359714)
- Source: [plugos.net](https://plugos.net/blog/2026/08/confidential-ai-how-plugclaw-keeps-your-ai-work-private/)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T10:47:57Z

## Translation

タイトル: 機密 AI: プライバシーを引き継がずに AI を働かせましょう
記事のタイトル: 機密 AI: PlugClaw が AI の動作をどのように保護するか
説明: AI が便利になればなるほど、より多くのデータが必要になります。 PlugClaw is Confidential AI がハードウェア レベルの暗号化でワークスペース全体をどのように保護するかをご覧ください。

記事本文:
機密 AI: PlugClaw が AI の動作をどのように保護するか
ホーム
AI はもはや質問に答えるだけのツールではなく、ユーザーのために働くエージェントになりつつあります。 1 年前なら、AI に記事の要約やメールの下書きを依頼していました。 AI エージェントは、受信トレイの処理、契約のレビュー、コードの作成、API の呼び出し、さらにはアプリやデータベースの制御も行うことができます。
獲物は？ AI が便利であればあるほど、あなたについてより多くのことを知る必要があります。物事を遂行するには、エージェントがあなたの個人的な電子メール、ファイル、ソースコード、顧客データ、財務記録、ログイン資格情報、長期記憶、およびサードパーティのサービスへのアクセスが必要になる場合があります。
つまり、AI エージェントの時代において、私たちが保護する必要があるのは、単一のプロンプトだけではありません。これは、AI がユーザーに代わって作業中にアクセスできるデジタル ワークスペース全体です。 AI の知識が増え、アクセスできる範囲が増えるほど、プライバシーとセキュリティのリスクは高まります。
クラウド AI: 優れたパワー、より優れた信頼性
現在の最も強力な AI モデルはクラウドに存在します。これは、携帯電話やラップトップでは提供できない大量のコンピューティングを必要とするためです。大規模なモデル、長いコンテキスト ウィンドウ、および複雑なエージェント タスクはすべて、ほとんどのローカル ハードウェアが提供できるリソースをはるかに超えるリソースを必要とします。
したがって、高度な AI を使用すると、プロンプト、ファイル、コンテキストを必然的にクラウド環境に送信することになります。これは、AI プロバイダーとそのクラウド インフラストラクチャが責任を持ってデータを処理することを信頼することを意味します。彼らはプライバシー ポリシー、アクセス制御、内部監査を導入しているかもしれませんが、結局のところ、そのモデルは依然として信頼に基づいていることが多く、「見ないよう信じてください」ということです。
また、リクエストがプロキシ、ゲートウェイ、またはその他のサードパーティを通過する場合、ホップが追加されるたびに、平文データが公開される可能性のある別の場所が作成され、信頼する必要がある別のパーティが作成されます。
本当に機密性の高いデータの場合、より良い q

問題は、許可されていない当事者がそれを見ないようにするだけではなく、技術的に困難にすることができるかどうかです。
AI エージェントにとって、この質問はさらに重要になります。彼らは 1 回限りのプロンプトを処理するだけではありません。彼らは、あなたの個人的な生活や職業生活に及ぶ可能性のあるデータを継続的に蓄積していきます。
オンデバイス AI: プライバシーの向上、機能の制限
明らかな答えは、すべてをローカルで実行することです。 AI がデバイス上で実行される場合、データはデバイスから離れる必要はありません。
ただし、ローカル デバイスにはコンピューティング、メモリ、バッテリー寿命に関して厳しい制限があります。クラウド インフラストラクチャほど効果的に最大のモデルを実行したり、非常に長いコンテキスト ウィンドウを処理したり、複雑な複数ステップの推論を実行したりすることはできません。
したがって、次のようなトレードオフが残ります。
オンデバイス: データはプライベートのままですが、AI の能力は劣ります。
クラウド: AI はより強力ですが、データに関しては AI を信頼する必要があります。
それが AI プライバシーのパラドックスです。保護が最も重要なデータは、多くの場合、最も強力な AI がアクセスする必要があるデータと同じです。
クラウドの約束を完全に信頼することはできず、AI の機能を低下させるわけにもいかないため、第 3 の方法、つまり、ほぼローカルのプライバシー保護を備えたクラウドのコンピューティング機能が必要です。
それが Confidential AI が構築された目的です。
従来のデータ セキュリティは、保存データと転送中のデータという 2 つの状態に焦点を当てています。しかし、AI は 3 番目の、より厄介な状態、つまりデータ使用中を導入します。
計算中、モデルは結果を生成するために平文データを処理する必要があります。機密データをコンピューティング環境に公開する必要があるのはまさにこのときです。
Confidential Computing は、この段階でデータを保護するように設計されています。ハードウェアベースの信頼できる実行環境 (TEE)、ハードウェアによるメモリ暗号化、分離メカニズムを使用して、機密データと計算を保護された実行環境内に保持します。

nt。
重要な実装の 1 つは Confidential VM (CVM) です。
ハードウェア暗号化メモリ: VM のメモリはハードウェア レベルで暗号化され、分離されます。保護された環境の外では、データは暗号化されたままになります。クラウド プロバイダーのオペレーティング システム、ハイパーバイザー、管理者は、通常のサーバーのように単純にメモリから平文を読み取ることはできません。
リモート構成証明 : 機密データを送信する前に、デバイスはリモート環境から暗号化された署名構成証明を検証できます。これは、環境の ID と実行されているコードに関する証拠を提供します。デバイスは、予期した環境が実行されていることを確認した後にのみ、データを解放します。
この環境内では、機密データと計算はハードウェアで保護された境界の背後に隔離されています。たとえクラウドの基盤となるオペレーティング システムやハイパーバイザーが侵害されたとしても、それらのコンポーネントは保護された環境内の平文に単純にアクセスすることはできません。
簡易バージョン: 機密 AI とは、データを計算することはできますが、見ることはできないことを意味します。
Confidential AI の仕組みを理解すると、Confidential AI が AI に何を変えることができるかについて 3 つのことが明らかになります。
1.「信頼する」から「検証する」へ
従来のクラウド AI では、セキュリティはプロバイダーのプライバシー ポリシー、アクセス制御、内部慣行に大きく依存します。言い換えれば、「見ないようにしてください。」
機密性のある AI がモデルを変える。ポリシーや約束だけに依存するのではなく、ハードウェアベースの分離、暗号化、暗号検証を使用して、不正アクセスを技術的に困難または不可能にします。
目標は信頼を完全に排除することではありません。それは、信頼すべきもの、つまりデータにアクセスできるシステム、人員、コンポーネントを減らすことです。そして、安全性の保証は、

単に約束するのではなく、検証する必要があります。
2. プライバシーと機能のトレードオフを打破する
データをプライベートに保つためだけに、弱いローカル モデルを使用する必要はもうありません。また、最も機能的な AI モデルを使用するためだけに、機密データを保護されていないクラウド環境に渡す必要は必ずしもありません。
Confidential AI は 3 番目のオプションを提供します。ハードウェアで保護された環境内で機密の個人データやビジネス データを処理しながら、クラウドのコンピューティング能力を使用します。
プライバシーと AI 機能を二者択一で選択する必要はなくなりました。
3. AI エージェントのより強力なセキュリティ境界
AI エージェントは、お客様に代わって電子メール、ソース コード、財務データ、長期記憶、その他の機密リソースにアクセスできます。エージェントの能力が向上し、より多くの権限を取得するにつれて、データのセキュリティ境界がますます重要になります。
Confidential AI を使用すると、機密データとエージェントの計算をハードウェアで保護された環境内で処理できます。エージェントの知識が増え、実行できることが増えるほど、周囲のインフラストラクチャの信頼だけに依存しないセキュリティ境界を持つことがより重要になります。
PlugClaw の機密 AI アーキテクチャ
AI エージェントの世界に機密コンピューティングを導入するために、PlugClaw はクラウド処理とモデル推論を通じてデバイスからのデータを保護する機密 AI アーキテクチャを使用します。
従来の AI フローは単純です: ユーザー → AI プロバイダー → AI モデル。プロンプト、ファイル、コンテキストを AI プロバイダーに直接送信します。
PlugClaw は、PlugClaw デバイス → 機密クラウド → 機密 / 非機密 AI モデルという異なるアプローチを採用しています。機密クラウドはデータと外部 AI の間に位置します。最も機密性の高いデータを処理し、モデルに送信される内容を正確に制御します。
プラグクロー

デバイス: データは自宅に保管されます
PlugClaw は、独自のマルチコア CPU、メモリ、暗号化ストレージを備えた小型のスタンドアロン AI コンピューターで、PlugOS と呼ばれる強化された Android OS を実行します。 USB経由で携帯電話またはコンピュータに接続します。ホスト デバイスは画面とキーボードを提供しますが、PlugClaw は独自の AI ワークスペースを独立して実行します。
PlugClaw は携帯電話やコンピュータ上では実行されず、ホスト上のすべてに自動的にアクセスできるわけでもありません。チャット履歴、ファイル、ナレッジ ベース、エージェント データはすべて、PlugClaw の暗号化されたストレージ内に保存されます。これは単なる別のアプリではなく、AI ワークスペースです。
タスクが本当にクラウドスケールのパワーを必要とする場合にのみ、PlugClaw は必要最小限のデータを機密クラウドに送信します。
Confidential Cloud: ハードウェアで保護された AI 環境
PlugClaw のクラウド インフラストラクチャは単なるサーバーではありません。
AI ワークロードはハードウェアで隔離された機密コンピューティング環境内で実行され、クラウド サービスのさまざまな部分が個別のセキュリティ ゾーンに分離されます。
ビジネス サービス レイヤーはアカウント、請求、ログを処理しますが、プロンプト、ファイル、エージェント コンテキストには決して触れません。
実際のデータ処理は、ハードウェアで保護されたエンクレーブである機密 VM 内で行われます。
PlugClaw がクラウドに接続すると、まずリモート認証を実行して、環境が本物であり、正しいコードが実行されていることを確認します。その後初めてデータが解放されます。 VM 内でデータは復号化されて処理され、平文はハードウェアで保護されたスペース内にのみ存在します。クラウド独自の OS と管理者はそれを読み取ることができません。
これにより、PlugClaw の Confidential Cloud は単なる API プロキシ以上のものになります。これは、クラウド側の処理全体にわたって機密データを隔離した状態に保つように設計された、ハードウェアで保護された AI 実行環境です。
モデル推論: 2 つのレベル

保護の
PlugClaw Confidential AI は、AI モデルの 2 つの使用方法をサポートしています。タスクの感度と必要な AI 機能のレベルに最適なモデルを選択できます。
機密推論 : モデル自体は、信頼されたドメイン内の機密コンピューティング環境内で実行されます。データは、PlugClaw から推論を経て保護された環境内に残ります。モデル演算子を含め、信頼された実行環境の外部の当事者は平文にアクセスできません。これは Confidential AI の完全な形であり、エンドツーエンドの機密性への道です。
非機密推論 : Claude、GPT、Gemini などのクローズドソース モデルでは、モデル プロバイダーの推論環境は信頼できるドメインの外にあります。したがって、モデルプロバイダーは、処理するリクエストの平文コンテンツを確認できます。この場合、Confidential Cloud は匿名化ゲートウェイとして機能します。元のデータ、エージェントのコンテキスト、機密環境内でのデータ処理を保護します。外部モデルプロバイダーは、実際にモデルに送信される特定のコンテンツのみを参照します。
クローズドソース モデルを使用しても、それは依然として「機密 AI」ですか?
これはおそらく、Confidential AI に対する PlugClaw のアプローチに関して尋ねるべき最も重要な質問です。
PlugClaw アーキテクチャでは、「機密」とは、次の 3 つの異なる層にわたる保護を意味します。
ホスト デバイスからの機密情報: 携帯電話やコンピュータ上のオペレーティング システム、アプリ、さらにはマルウェアは、PlugClaw 内に保存されているデータにアクセスできません。ハードウェアの分離により、AI ワークスペースがホスト デバイスから分離されます。
TrustKernel とクラウド プロバイダーからの機密: データはクラウド内で常に Confidential VM 内で処理されます。機密コンピューティングとリモート認証により TrustKernel を暗号化して維持

基盤となるクラウド プロバイダーは信頼された実行環境の外にあるため、データにアクセスできません。
モデルプロバイダーの機密情報 : これは選択したモデルによって異なります。機密推論モデルでは、推論パス全体が信頼できるドメイン内に残り、エンドツーエンドの機密性が提供されます。従来のクローズドソース モデルでは、この層は適用されません。モデルのプロバイダーは、処理するリクエストの平文コンテンツを確認できますが、匿名化されたルーティングを通じて、リクエストが誰に属しているかを直接確認することはできません。
では、すべてに単純に機密推論モデルを使用してはどうでしょうか?
なぜなら、今日の現実として、Claude、GPT、Gemini などの最も有能なフロンティア モデルの多くは、機密コンピューティング環境内での展開には利用できないからです。
「機密」がこれらのモデルへのアクセスを放棄することを意味する場合、ユーザーは強力な AI かプライバシーという古い選択に戻るだけでしょう。
PlugClaw は異なるアプローチを採用し、ユーザーに選択肢を与えます。
契約、医療記録、財務情報、その他の機密データなど、非常に機密性の高いタスクには、機密推論モデルを使用してエンドツーエンドの保護を行います。
最も強力なフロンティア モデルを必要とするタスクには、モデルを使用します。

[切り捨てられた]

## Original Extract

The more useful AI gets, the more data it needs. See how PlugClaw is Confidential AI protects your entire workspace with hardware-level encryption.

Confidential AI: How PlugClaw Protects Your AI Work
Home
AI is no longer just a tool that answers questions - it's becoming an agent that works for you. A year ago, you'd ask AI to summarize an article or draft an email. Now, AI agents can process your inbox, review contracts, write code, call APIs, and even control apps and databases.
The catch? The more useful an AI is, the more it needs to know about you. To get things done, an agent may need access to your personal emails, files, source code, customer data, financial records, login credentials, long‑term memory, and third‑party services.
So in the age of AI agents, what we need to protect isn't just a single prompt. It's the entire digital workspace the AI can access while working on your behalf. The more an AI knows and the more access it has, the higher the privacy and security stakes.
Cloud AI: Great Power, Greater Trust
Today's most powerful AI models live in the cloud - because they require massive amounts of compute that your phone or laptop simply can't provide. Large models, long context windows, and complex agent tasks all demand resources far beyond what most local hardware can deliver.
So when you use advanced AI, you inevitably send your prompts, files, and context to a cloud environment. That means trusting the AI provider and its cloud infrastructure to handle your data responsibly. They may have privacy policies, access controls, and internal audits in place - but at the end of the day, the model is still largely based on trust: "Trust us not to look."
And if your requests pass through proxies, gateways, or other third parties, every additional hop creates another place where plaintext data could be exposed - and another party you have to trust.
For truly sensitive data, the better question is: Can we make it technically hard for unauthorized parties to see it - rather than simply asking them not to?
That question becomes even more important for AI agents. They don't just process one-off prompts; they work with continuously accumulating data that can span your personal and professional life.
On-Device AI: Better Privacy, Limited Capability
The obvious answer is to run everything locally. If the AI runs on your device, your data never has to leave it.
But local devices have hard limits on compute, memory, and battery life. They simply can't run the largest models, handle very long context windows, or perform complex multi-step reasoning as effectively as cloud infrastructure can.
So you're left with a trade-off:
On-device: Your data stays private, but the AI is less capable.
Cloud: The AI is more powerful, but you have to trust it with your data.
That's the AI privacy paradox. The very data you care about protecting most is often the same data the most powerful AI needs access to.
Since we can't fully trust cloud promises, and we can't afford to dumb down our AI, we need a third way: the cloud's computing muscle with near-local privacy protection.
That's what Confidential AI is built for.
Traditional data security focuses on two states: data at rest and data in transit. But AI introduces a third, trickier state - data in use.
During computation, a model has to process plaintext data to generate results. That's precisely when sensitive data needs to be exposed to the computing environment.
Confidential Computing is designed to protect data at this stage. It uses hardware-based Trusted Execution Environments (TEEs), hardware-backed memory encryption, and isolation mechanisms to keep sensitive data and computation inside a protected execution environment.
One key implementation is the Confidential VM (CVM):
Hardware-encrypted memory : The VM's memory is encrypted and isolated at the hardware level. Outside the protected environment, the data remains encrypted. The cloud provider's operating system, hypervisor, and administrators cannot simply read the plaintext from memory as they could on an ordinary server.
Remote attestation : Before sending sensitive data, your device can verify a cryptographically signed attestation from the remote environment. It provides evidence about the environment's identity and the code it is running. Only after the device verifies that the expected environment is running does it release the data.
Inside this environment, sensitive data and computation are isolated behind a hardware-protected boundary. Even if the cloud's underlying operating system or hypervisor is compromised, those components cannot simply access the plaintext inside the protected environment.
Simple version: Confidential AI means your data can be computed, but it can't be seen.
Once you understand how Confidential AI works, three things become clear about what it can change for AI.
1. From "Trust Us" to "Verify It"
With traditional cloud AI, security largely depends on the provider's privacy policies, access controls, and internal practices. In other words: "Trust us not to look."
Confidential AI changes the model. Instead of relying solely on policies and promises, it uses hardware-based isolation, encryption, and cryptographic verification to make unauthorized access technically difficult or impossible.
The goal isn't to eliminate trust entirely. It's to reduce what you have to trust - fewer systems, fewer people, and fewer components with access to your data. And the security guarantees become something that can be verified, rather than simply promised.
2. Breaking the Privacy-Capability Trade-off
You no longer have to settle for weaker local models just to keep your data private. And you don't necessarily have to hand sensitive data over to an unprotected cloud environment just to use the most capable AI models.
Confidential AI offers a third option: use the computing power of the cloud while processing sensitive personal and business data inside a hardware-protected environment.
Privacy and AI capability no longer have to be a binary choice.
3. A Stronger Security Boundary for AI Agents
AI agents can access emails, source code, financial data, long-term memory, and other sensitive resources on your behalf. As agents become more capable and gain more permissions, the security boundary around that data becomes increasingly important.
With Confidential AI, sensitive data and agent computation can be processed within a hardware-protected environment. The more an agent knows and the more it can do, the more important it becomes to have a security boundary that doesn't depend solely on trusting the infrastructure around it.
PlugClaw's Confidential AI Architecture
To bring confidential computing into the world of AI agents, PlugClaw uses a confidential AI architecture that protects data from the device all the way through cloud processing and model inference.
The traditional AI flow is straightforward: User → AI provider → AI model. You send your prompts, files, and context directly to the AI provider.
PlugClaw takes a different approach: PlugClaw device → Confidential Cloud → Confidential / Non-Confidential AI Model. The confidential cloud sits between your data and the external AI. It handles the most sensitive data and controls exactly what gets sent to the model.
PlugClaw Device: Your Data Stays at Home
PlugClaw is a tiny, standalone AI computer - with its own multi‑core CPU, memory, and encrypted storage - running a hardened Android OS called PlugOS. It connects to your phone or computer via USB. Your host device provides the screen and keyboard, but PlugClaw runs its own AI workspace independently.
PlugClaw does not run on your phone or computer, and it doesn't automatically gain access to everything on your host. Your chat history, files, knowledge base, and agent data all live inside PlugClaw's encrypted storage. It's your AI workspace - not just another app.
Only when a task truly needs cloud‑scale power does PlugClaw send the minimum necessary data to the confidential cloud.
Confidential Cloud: A Hardware‑Protected AI Environment
PlugClaw's cloud infrastructure is not just another server.
It runs AI workloads inside hardware-isolated confidential computing environments, with different parts of the cloud service separated into distinct security zones:
The business service layer handles accounts, billing and logging - but it never touches your prompts, files, or agent context.
The actual data processing happens inside confidential VMs - hardware‑protected enclaves.
When PlugClaw connects to the cloud, it first performs remote attestation to verify that the environment is genuine and running the correct code. Only then does it release any data. Inside the VM, your data is decrypted and processed - and the plaintext exists only inside that hardware‑protected space. The cloud's own OS and admins can't read it.
That makes PlugClaw's Confidential Cloud more than an API proxy. It is a hardware-protected AI execution environment designed to keep sensitive data isolated throughout cloud-side processing.
Model Inference: Two Levels of Protection
PlugClaw Confidential AI supports two ways of using AI models. You can choose the model that best fits the sensitivity of your task and the level of AI capability you need.
Confidential inference : The model itself runs inside a confidential computing environment within the trusted domain. Your data remains within protected environments from PlugClaw through inference and back. No party outside the trusted execution environment can access the plaintext, including the model operator. This is the full form of Confidential AI - and the path to end-to-end confidentiality.
Non-confidential inference : With closed-source models such as Claude, GPT, and Gemini, the model provider's inference environment is outside our trusted domain. The model provider can therefore see the plaintext content of the request it processes. In this case, the Confidential Cloud acts as an anonymizing gateway. It protects your original data, agent context, and data processing inside the confidential environment. The external model provider sees only the specific content that is actually sent to the model.
If I Use a Closed‑Source Model, Is It Still "Confidential AI"?
This is probably the most important question to ask about PlugClaw's approach to Confidential AI.
In the PlugClaw architecture, "confidential" means protection across three distinct layers:
Confidential from the host device : The operating system, apps, or even malware on your phone or computer cannot access the data stored inside PlugClaw. Hardware isolation keeps your AI workspace separate from the host device.
Confidential from TrustKernel and the cloud provider : Data is processed inside Confidential VMs throughout its time in the cloud. Confidential computing and remote attestation cryptographically keep TrustKernel and the underlying cloud provider outside the trusted execution environment, so they cannot access your data.
Confidential from the model provider : This depends on the model you choose. With a confidential inference model, the entire inference path remains within the trusted domain, providing end-to-end confidentiality. With a conventional closed-source model, this layer does not apply: the model provider can see the plaintext content of the request it processes - but, through anonymized routing, does not directly see who the request belongs to.
So why not simply use confidential inference models for everything?
Because the reality today is that many of the most capable frontier models - including Claude, GPT, and Gemini - are not available for deployment inside a confidential computing environment.
If "confidential" meant giving up access to those models, users would simply be back to the old choice: powerful AI or privacy.
PlugClaw takes a different approach: give users the choice.
For highly sensitive tasks - contracts, medical records, financial information, and other confidential data - use a confidential inference model for end-to-end protection.
For tasks that demand the strongest frontier models, use models s

[truncated]
