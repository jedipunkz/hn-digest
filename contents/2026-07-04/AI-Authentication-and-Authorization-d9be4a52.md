---
source: "https://fusionauth.io/articles/ai/ai-authentication-authorization"
hn_url: "https://news.ycombinator.com/item?id=48789910"
title: "AI Authentication and Authorization"
article_title: "AI Authentication and Authorization | FusionAuth Docs"
author: "mooreds"
captured_at: "2026-07-04T23:53:52Z"
capture_tool: "hn-digest"
hn_id: 48789910
score: 1
comments: 0
posted_at: "2026-07-04T23:10:55Z"
tags:
  - hacker-news
  - translated
---

# AI Authentication and Authorization

- HN: [48789910](https://news.ycombinator.com/item?id=48789910)
- Source: [fusionauth.io](https://fusionauth.io/articles/ai/ai-authentication-authorization)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T23:10:55Z

## Translation

タイトル: AIの認証と認可
記事のタイトル: AI の認証と認可 | FusionAuth ドキュメント
説明: この記事では、RAG、MCP、エージェント システムという 3 つの一般的な AI ユース ケースについて説明します。また、認証と認可がこれらのユースケースとどのように相互作用するかについても説明します。

記事本文:
AIの認証と認可 | FusionAuth ドキュメント / 記事
ログイン デモを入手 メイン メニューを開く 認証 認証システム ロックインの回避 FusionAuth での SSO と MFA の組み合わせ 一般的な認証実装のリスクとその軽減方法 FedCM: Web でのログイン方法を変える可能性がある新しい提案された ID 標準 SSO はどのように機能しますか?シングル サインオンについて FusionAuth が MFA 要件への準拠を簡素化する方法について説明しました パスワードレス SSO: 安全でシームレスな認証 MFA によるユーザー エクスペリエンスの保護 シングル サインオンの開発者の利点 Kubernetes Auth のパスワードレス認証タイプのセキュリティへの影響 WebAuthn について説明しました 多要素認証 (MFA) とは何ですか?パスワードレス認証とは何ですか?またその仕組みは何ですか? WebAuthn とは何ですか?またその仕組みは何ですか?ログイン失敗がなぜ重要なのか WebAuthn の時代にパスワードレス認証が重要なのか CIAM 認証ボトルネック アーキテクチャ - アプリケーション全体での認証の合理化 認証ファサード パターン - 環境全体での ID 管理の簡素化 CIAM の課題 - セキュリティ、使いやすさ、コンプライアンスのナビゲート CIAM と IAM - 顧客 ID と ID アクセス管理の説明 GDPR 開発者ガイドCIAM へのコンプライアンスの達成 CIAM によるサインアップ フローの最適化: 摩擦を軽減し、成長を促進する 認証のスケーリング - パフォーマンスとセキュリティのベスト プラクティス 顧客 ID およびアクセス管理 (CIAM) とは何ですか? CIAM は、CIAM がサードパーティ Cookie の終焉を生き抜く鍵である理由を説明しました。 ID の基本 登録フォームのベスト プラクティス 認証プロバイダーはどの程度完全である必要がありますか?パスワードレス認証は規制を遵守する準備ができていますか? Magic Links - パスワードレス認証のガイド マルチテナントとシングルテナント

-テナント IDaaS ソリューション オープン ソースと商用認証プロバイダーの比較 認可セキュリティの最適化: アクセス コントロール モデルのガイド 認証ベンダーに対するデュー デリジェンスの実行 ユーザー データの移行が遅い User-Cent をコミットする前に認証プロバイダーを試す価値
[切り捨てられた]
人間のアイデンティティは AI の権威の源です。
あなたが考えていることはわかります。AI セキュリティに関する別の記事?ついてきてください。これは、業界がエージェントの出荷を急ぐあまりに忘れ続けている、単純かつほぼ明白な真実に基づいているため、これとは異なります。つまり、2010 年代の API ブームを保護したのと同じ ID と認証パターンが、今日の AI システムを保護するためにまさに必要なものなのです。
OAuth 統合を構築したり、API キーを管理したり、ロールベースのアクセス制御を設定したりしたことがある場合は、必要な知識のほとんどをすでに持っています。 AI 認証は新しい分野ではありません。これは既存のベスト プラクティスの拡張です。
「ランダムな数字を生成する算術的方法を考える人は、もちろん罪を犯しています。」
— ジョン・フォン・ノイマン
「決定的な安全策を講じずに AI にリソースへのアクセスを許可する人は、当然ながら愚かな状態にあります。」
— 著者
フォン・ノイマンの引用は、信頼できるものを使用しても非決定的な出力が得られると仮定することに対する古典的な警告です。ここでは逆の原理が適用されます。 AI システムは確率論的であり、推論し、幻覚を示し、即興で行動します。しかし、誰のために行動するのか、何が許可されるのかを管理するアイデンティティ層は決定的でなければなりません。アイデンティティは「雰囲気」を与えるものではありません。
この記事では、次の 3 つの AI ユースケースについて説明します。
検索拡張世代 (RAG)
そして、認証、認可、アイデンティティ管理のレンズを通してそれらを検査します。 FusionAuth の例を使用していますが、標準ベースのソリューションが存在する場所についても説明しています。

ns。
全体を通して 1 つの実行例を使用します。あなたは銀行のエンジニアリング マネージャーで、従業員と顧客の両方に対するサポート デスクの業務を AI で改善したいと考えています。
ユースケースの概要 #
本題に入る前に、何を扱うのかを定義しましょう。
検索拡張生成 (RAG) は、クエリ時にドキュメントをフィードすることにより、AI モデルで利用可能なデータを拡張します。銀行員または顧客が質問すると、RAG システムは関連する内部文書を取得し、それを LLM に提供して、LLM の回答を根拠付けます。認証に関する重要な懸念事項は、すべてのユーザーがすべてのドキュメントを見る必要があるわけではないということです。顧客は窓口担当者から異なる書類を見ることになり、窓口担当者は副社長から異なる書類を見ることになります。
ツール (MCP および API) を使用すると、AI システムがデータベースからの読み取り、顧客レコードの更新、外部サービスの呼び出しなどのアクションを実行できるようになります。 Model Context Protocol (MCP) は、AI ツールをサービスに接続するための新しい標準ですが、豊富なドキュメントを備えたプレーンな API も機能します。認証に関する重要な点は、各ツールが誰に代わって何を実行できるかを制御することです。
エージェント システムは、データを読み取り、複数のシステムにわたってアクションを実行し、必要に応じて人間の入力を求めることができる、半自律的なタスク指向のワークフローです。これらは、推論ステップを連鎖させる非決定的なソフトウェア コンポーネントです。認証に関する主な懸念事項は、ワークフローを承認した人間から実行されるすべてのアクションに至るまで一連の ID を維持することと、エージェントのアクセスを制限することです。
これらが ID プロバイダーが支援できる内容とどのように対応しているかは次のとおりです。
次に、これらの各ユースケースを詳しく見てみましょう。
RAG: モデルが見るべきではないものを見ないようにする #
ローン契約、顧客契約、コンプライアンス ポリシーなど、カスタマー サポート業務に関連する銀行文書があります。

管理戦略、および不正調査手順。顧客や従業員が AI インターフェイスを通じてクエリを実行できるようにしたいと考えています。ただし、すべてのドキュメントをすべてのユーザーが利用できるようにする必要はありません。顧客サポート、詐欺とセキュリティ、紛争とチャージバック、ローン返済チームはそれぞれ、異なるドキュメント セットにアクセスする必要があります。そして顧客自身のことも忘れてはいけません。
LinkedIn 、 DoorDash 、 Vimeo などの企業は、すでに運用環境で RAG を使用しています。パターンがしっかりと確立されています。
RAG にとってアイデンティティが重要な理由 #
クエリに答えるとき、LLM はユーザーがアクセスすべきではないドキュメントを決して表示すべきではありません。巧妙なプロンプトを作成する必要はありません。秘密を守るためにモデルに依存しているわけではありません。適切な認証フレームワークを使用すると、ドキュメントがモデルに到達する前にフィルタリングできます。
これは主に認証の問題です。ユーザーを認証し (本人であることを証明し)、クエリを処理し、ベクター データストアからドキュメントを取得して、ユーザーがクエリを許可されているドキュメントに基づいてドキュメントをフィルタリングします。
モデルは、承認チェックに合格したドキュメントのみを受け取ります。
実装は単純なパイプラインに従います。
ドキュメントをベクトル検索に適したセグメントに分割します。
ユーザーとロールをドキュメント アクセスにマップする承認スキーマを構築します。
どの役割、部門、ユーザーが各チャンクにアクセスできるかなど、メタデータをドキュメント チャンクと一緒に Vector データベースに保存します。
取得時にユーザーを認証し、ID クレームを取得します。
結果を LLM に渡す前に、手順 3 で保存したユーザーおよびドキュメントの属性でフィルタリングします。
一部のフレームワークでは認証に JWT を使用します。他のものは API キーを使用します。フィルタリング メカニズムは、RAG フレームワークにも依存します。例えば、

LangChain を使用すると、結果を返す前に承認サービスを呼び出す取得ラッパーを構築できます。
認可チェックには、きめ細かい認可 (FGA) システムを使用します。 Permify の FusionAuth FGA はオプションの 1 つです。データの安全性を確保するためにオンサイトに導入できる確定的な承認を提供し、ニーズに合わせて拡張できます。
使用している RAG フレームワークに関係なく、承認ロジックは一元化され、単一の信頼できる情報源である必要があります。これを利用して、確率論的ではなく決定論的なフィルターが必要です。
以下は、読み込み中に適切なメタデータがドキュメントに保存されたときのリクエスト フローの簡略図です。
フローチャート LR
サブグラフ User["User"]
A[認証する]
終わり
サブグラフ RAG["RAG パイプライン"]
B[クエリベクターDB]
C[ユーザー属性とドキュメント属性を備えた FGA を使用したフィルター]
D[承認されたチャンクを返す]
終わり
サブグラフ LLM["LLM"]
E[チャンクを含むレスポンスを生成]
終わり
A --> B --> C --> D --> E
しかし、そのメタデータを取得する場合はどうすればよいでしょうか?ドキュメントは常に特定のアクセス レベルに適切にマッピングされるわけではなく、一部のドキュメントはチャンクごとに異なるアクセス権を持つ場合があります。チャンク化によりメタデータが失われる可能性があります。
たとえば、コンプライアンス PDF には、法務チームに限定されたセクションとともに、すべての従業員がアクセスできるセクションが含まれる場合があります。チャンク パイプラインがこれを処理できることを確認してください。
したがって、RAG プロセスの一部としてユーザーをキャプチャし、メタデータにアクセスすることを計画してください。ユーザーがアクセスすべきではないドキュメントを LLM に決して表示させないようにするには、ユーザーとアクセス データが利用可能であることを確認する必要があります。
顧客サービス チームのメンバーが AI ツールを使用して銀行の顧客情報 (連絡先詳細、口座設定、サービス リクエスト) を更新できるようにしたいとします。ただし、役割ごとに使用できるツールが異なり、同じツールであっても使用できる機能が異なります。

nt ユーザーには異なる制限があります。 Tier 1 サポート エージェントは、電話番号を更新できても、与信限度額を調整できない場合があります。
モデル コンテキスト プロトコル (MCP) は、あらゆる API またはサービスを構造化された方法で AI ツールにアクセスできるようにする新しい標準です。 Block、Bloomberg、Amazon などの企業はすでに社内で MCP を使用しています。ただし、MCP が唯一の選択肢ではありません。プレーンな API も適切に機能します。 AI モデルは、優れたドキュメントから API セマンティクスを理解できます。
公開時点での MCP の最新バージョンでは、AI システムまたはツールの認証に OAuth 2.1 と認可コードグラントが使用されています。クライアント資格情報付与を使用するための拡張機能も開発中です。
API は、API キーまたはアクセス トークンといった従来の認証方法を再利用します。
REST API 時代から使用してきたのと同じゲートウェイ パターンは、MCP サーバーまたは API サーバーのアクセスのレート制限や監視に役立ちます。
ID を使用して MCP を設定する方法は次のとおりです。
既存の API とサービスの上に MCP サーバーを構築します。 OAuth 2.1 をサポートする ID プロバイダーを指すように MCP サーバーを構成します。 MCP クライアントは、事前に登録するか、動的に作成する必要があります。
MCP クライアントが MCP サーバーにアクセスしようとすると、MCP サーバーは設定された ID プロバイダーにリダイレクトする必要があります。これにより、MCP クライアントを駆動するユーザーが認証され、トークンが発行されます。その後、トークンが MCP サーバーに提示されます。
MCP と実装について詳しくは、こちらをご覧ください。
OAuth スコープが提供する範囲を超えた詳細な制御が必要な場合は、MCP サーバーがアクセスしているサービスに詳細な承認を追加する必要がある場合があります。
API アクセスの場合、パターンはさらに単純です。
既存の API とサービスを使用します。 MCP サーバーは必要ありません。
ID プロバイダーを使用してユーザーを認証します。
AI が Web ツールを持っている場合、AI はアクセスできます。

REST 呼び出しを使用して API を呼び出し、トークンを渡します。
API を使用した SDK も利用可能にすることを検討してください。繰り返しになりますが、OAuth スコープが提供する範囲を超えた詳細な制御が必要な場合は、API とサービスに詳細な承認を追加する必要がある場合があります。
おそらく、再利用できる可能性のある認証と API に関するインフラストラクチャがいくつかあるでしょう。たとえば、複数の API ゲートウェイが FusionAuth で動作します。
エージェントシステム: 前進して仕事をする #
ここが興味深いところであり、AI 認証における新しい考え方が必要となるところです。
エージェントは、さまざまなレベルの自律性でタスクを完了するように指示できる非決定的なソフトウェア コンポーネントです。これらは数十または数百のインスタンスにスケールし、人間、API、および MCP ツールと対話し、推論ステップを連鎖させます。
あなたの銀行は、新しい法人口座の設定を自動化したいと考えています。新しいビジネスには、当座預金口座、普通預金口座、販売者サービス、給与計算の設定が必要です。エージェントは次のことを行う必要があります。
ビジネスの種類を評価し、パッケージを推奨します
ビジネス文書 (EIN、定款) を文書ストアから収集する
API経由で信用度を確認
カレンダー サービスを介してリレーションシップ マネージャーとのオンボーディング セッションをスケジュールする
これは複数のステップです

[切り捨てられた]

## Original Extract

This article covers three common AI use cases: RAG, MCP and agentic systems. It also discusses how authentication and authorization interact with these use cases.

AI Authentication and Authorization | FusionAuth Docs / Articles
Log In Get a demo Open main menu Authentication Avoiding Authentication System Lock-In Combining SSO and MFA in FusionAuth Common Authentication Implementation Risks And How To Mitigate Them FedCM: A New Proposed Identity Standard That Could Change How We Log In on the Web How Does SSO Work? Single Sign-On Explained How FusionAuth Simplifies Compliance with MFA Requirements Passwordless SSO: Secure and Seamless Authentication Securing Your User Experience with MFA The Developer Benefits of Single Sign-On The Security Implications Of Passwordless Authentication Types Of Kubernetes Auth WebAuthn Explained What Is Multi-Factor Authentication (MFA)? What is Passwordless Authentication and How Does It Work? What Is WebAuthn and How Does It Work? Why Login Failures Matter Why Passwordless Authentication Matters in the Age of WebAuthn CIAM Auth Bottleneck Architecture - Streamline Authentication Across Applications Auth Facade Pattern - Simplifying Identity Management Across Environments Challenges of CIAM - Navigating Security, Usability, and Compliance CIAM vs. IAM - Customer Identity vs. Identity Access Management Explained GDPR Developer's Guide to Achieving Compliance with Your CIAM Optimizing Sign-Up Flows with CIAM: Reduce Friction, Drive Growth Scaling Your Auth - Best Practices for Performance and Security What is Customer Identity and Access Management (CIAM)? CIAM Explained Why CIAM is the Key to Surviving the Demise of Third-Party Cookies Identity Basics Best Practices For Registration Forms How Complete Does Your Authentication Provider Need To Be? Is Passwordless Authentication Ready for Regulatory Compliance? Magic Links - A Guide to Passwordless Authentication Multi-Tenant vs. Single-Tenant IDaaS Solutions Open Source vs. Commercial Auth Providers Optimizing Authorization Security: A Guide to Access Control Models Performing Due Diligence On Authentication Vendors Slow Migration of User Data The Value of Trying Your Auth Provider Before You Commit User-Cent
[truncated]
Human identity is the source of AI authority.
I know what you're thinking: another article about AI security? Stick with me. This one is different because it's grounded in a simple, almost obvious truth that the industry keeps forgetting in its rush to ship agents: the same identity and authorization patterns that secured the API boom of the 2010s are exactly what you need to secure AI systems today.
If you've built OAuth integrations, managed API keys, or set up role-based access control, you already have most of the knowledge you need. AI auth is not a new discipline. It's an extension of existing best practices.
"Anyone who considers arithmetical methods of producing random digits is, of course, in a state of sin."
— John von Neumann
"Anyone who lets AI access resources without deterministic safeguards is, of course, in a state of folly."
— The Author
The Von Neumann quote is a classic warning about assuming you can take something reliable and get non-deterministic outputs. The inverse principle applies here. AI systems are probabilistic: they reason, hallucinate, and improvise. But the identity layer that governs who they act for and what they're allowed to do must be deterministic. Identity is not something to "vibe."
This article walks through three AI use cases:
retrieval-augmented generation (RAG)
And examines them through the lens of authentication, authorization, and identity management. It uses FusionAuth examples, but also notes where there are standards-based solutions.
We'll use a single running example throughout: you are an engineering manager at a bank, looking to improve support desk operations for both employees and customers, with AI.
A Quick Overview of the Use Cases #
Before we dive in, let's define what we're working with.
Retrieval-augmented generation (RAG) augments the data available to an AI model by feeding it documents at query time. Your bank employees or customers ask a question, and the RAG system retrieves relevant internal documents and then provides it to an LLM to ground the LLM's answer. The key auth concern: not every user should see every document. A customer is going to see different documents from a teller, who will see different ones from a VP.
Tool use (MCP and APIs) allows AI systems to take actions like reading from a database, updating a customer record, or calling an external service. The Model Context Protocol (MCP) is an emerging standard for connecting AI tools to services, but plain APIs with rich documentation work too. The key auth concern: controlling what each tool can do, and on whose behalf.
Agentic systems are semi-autonomous, task-oriented workflows that can read data, take action across multiple systems, and ask for human input when needed. They are non-deterministic software components that chain together reasoning steps. The key auth concern: maintaining a chain of identity from the human who authorized the workflow all the way through to every action taken, as well as limiting agents' access.
Here's how these map to what an identity provider can help with:
Now let's dig into each of these use cases.
RAG: Making Sure the Model Never Sees What It Shouldn't #
You have bank documents related to customer support tasks, such as loan agreements, customer agreements, compliance policies, wealth management playbooks, and fraud investigation procedures. You want to make them available for customers and employees to query through an AI interface. But not all documents should be available to every user. Customer support, fraud and security, disputes and chargebacks, and loan servicing teams each need access to different document sets. And don't forget customers themselves.
Companies like LinkedIn , DoorDash , and Vimeo already use RAG in production. The pattern is well-established.
Why Identity Matters for RAG #
When answering a query, the LLM should never even see documents the user shouldn't have access to. You don't have to craft some clever prompt. You're not relying on the model to keep secrets. With the right authorization framework, you're filtering documents before they reach the model.
This is primarily an authorization problem. You authenticate the user (prove they are who they claim to be), process their query, pull documents from the vector datastore, then filter the documents based on which documents the user is allowed to query.
The model only receives documents that pass the authorization check.
The implementation follows a straightforward pipeline:
Chunk your documents into segments suitable for vector search.
Build an authorization schema that maps users and roles to document access.
Store metadata alongside your document chunks in the vector database, including which roles, departments, or users can access each chunk.
On retrieval, authenticate the user and get their identity claims.
Filter by user and document attributes that you stored in step 3 before passing results to the LLM.
For authentication, some frameworks use JWTs for authentication; others use API keys. The filtering mechanism depends on your RAG framework as well. For example, LangChain allows you to build a retriever wrapper which calls out to an authorization service before returning results.
For the authorization checks, use a fine-grained authorization (FGA) system. FusionAuth FGA by Permify is one option. It provides deterministic authorization that can be deployed on-site for data safety and scales with your needs.
Your authorization logic should be centralized and a single source of truth, regardless of which RAG framework you're using. You want a filter to leverage this and be deterministic, not probabilistic.
Here's a simplified diagram of the request flow, when the proper metadata has been stored on the documents during the loading.
flowchart LR
subgraph User["User"]
A[Authenticate]
end
subgraph RAG["RAG Pipeline"]
B[Query Vector DB]
C[Filter using FGA with User and Document Attributes]
D[Return Authorized Chunks]
end
subgraph LLM["LLM"]
E[Generate Response Including Chunks]
end
A --> B --> C --> D --> E
But what about capturing that metadata? Documents don't always cleanly map to a given access level, and some documents may have different access for different chunks. Chunking may lose metadata.
For instance, a compliance PDF might contain sections accessible to all employees alongside sections restricted to the legal team. Make sure your chunking pipeline can handle this.
So, plan to capture the user and access metadata as part of your RAG process. If you want the LLM to never see documents the user shouldn't access, you have to make sure the user and access data is available.
Suppose you want to allow customer service team members to use AI tools to update bank customer information — contact details, account preferences, service requests. But different tools are available to different roles, and even with the same tools, different users have different limits. A tier-one support agent might be able to update a phone number but not adjust a credit limit.
The Model Context Protocol (MCP) is an emerging standard that makes any API or service accessible to AI tooling in a structured way. Companies like Block, Bloomberg, and Amazon are already using MCP internally. But MCP isn't the only option — plain APIs work well too. AI models are capable of figuring out API semantics from good docs.
The most recent version of MCP at the time of publishing uses OAuth 2.1 and the authorization code grant for authentication of an AI system or tool. There are also extensions under development for use of the client credentials grant.
APIs re-use traditional authentication methods: API keys or access tokens.
The same gateway patterns you've been using since the REST API era can help rate-limit or monitor access for either MCP or API servers.
Here's how to set up MCP with identity:
Build an MCP server on top of your existing APIs and services. Configure your MCP server to point to an identity provider which supports OAuth 2.1. MCP clients should be either preregistered or created dynamically.
When an MCP client tries to access an MCP server, the MCP server should redirect to the configured identity provider, which will authenticate the user driving the MCP client and then issue a token. The token is then presented to the MCP server.
Learn more about MCP and implementation .
You may need to add fine-grained authorization to the services the MCP server is accessing if you need granular control beyond what OAuth scopes provide.
For API access, the pattern is even simpler:
Use your existing APIs and services; no MCP server required.
Authenticate users with your identity provider.
If the AI has a web tool, it can access the API using REST calls, passing the token.
Consider making an SDK using your API available as well. Again, you may need to add fine-grained authorization to your APIs and services if you need granular control beyond what OAuth scopes provide.
You probably have some infrastructure around authentication and your APIs that you might be able to re-use. For example, multiple API gateways work with FusionAuth.
Agentic Systems: Go Forth And Do Work #
This is where things get interesting and where new thinking in AI auth needs to happen.
Agents are non-deterministic software components that can be prompted to complete a task with varying levels of autonomy. They scale to tens or hundreds of instances, interact with humans, APIs, and MCP tools, and chain together reasoning steps.
Your bank wants to automate new business account setup. A new business needs checking accounts, savings accounts, merchant services, and payroll setup. An agent needs to:
Assess the business type and recommend a package
Gather business documents (EIN, articles of incorporation) from a document store
Check creditworthiness via an API
Schedule an onboarding session with a relationship manager via a calendar service
This is a multi-step

[truncated]
