---
source: "https://techupdate24.com/aws-bedrock-agentcore-hijacked-ai-security/"
hn_url: "https://news.ycombinator.com/item?id=49384306"
title: "AWS Bedrock AgentCore enforces user context to prevent hijacked AI agents"
article_title: "Stop Hijacked AI Agents: AWS Bedrock AgentCore User Context"
image: "https://techupdate24.com/wp-content/uploads/2026/08/aws-bedrock-agentcore-hijacked-ai-security.jpg"
author: "sysadmin_diarie"
captured_at: "2026-08-21T06:28:52Z"
capture_tool: "hn-digest"
hn_id: 49384306
score: 2
comments: 0
posted_at: "2026-08-21T05:56:29Z"
tags:
  - hacker-news
  - translated
---

# AWS Bedrock AgentCore enforces user context to prevent hijacked AI agents

- HN: [49384306](https://news.ycombinator.com/item?id=49384306)
- Source: [techupdate24.com](https://techupdate24.com/aws-bedrock-agentcore-hijacked-ai-security/)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T05:56:29Z

## Translation

タイトル: AWS Bedrock AgentCore は、AI エージェントのハイジャックを防ぐためにユーザーコンテキストを強制します
記事のタイトル: ハイジャックされた AI エージェントを停止する: AWS Bedrock AgentCore ユーザーコンテキスト
説明: AWS Bedrock AgentCore がユーザー認証コンテキストと IAM ID 伝播を強制することで、ハイジャックされた AI エージェントを停止し、インジェクションを促す方法を学びます。

記事本文:
AWS Bedrock AgentCore がハイジャックされた AI エージェントを阻止する方法
企業は、データベース、ドキュメント リポジトリ、SaaS プラットフォームから取得してワークフローを自動化する AI エージェントの導入を急いでいます。しかし、利便性の裏には静かなリスクが潜んでいます。プロンプト インジェクションによって操作された侵害されたエージェントにより、要求元のユーザーに表示が許可されていないデータが引き渡される可能性があります。認可をエージェントのコードから AWS インフラストラクチャ自体に移すことで、このギャップを埋める方法を説明します。
このアーキテクチャの変更により、完全にハイジャックされたエージェントであっても、ユーザーの特定の権限に暗号的に制限されたままになります。最近の GitLab GraphQL の重大な欠陥の詳細で、不正なデータ操作の壊滅的な結果について説明したように、AI アクセスをアプリケーション レベルのフィルタリングに依存することは、重大な構造的弱点です。
エージェントレベルのフィルタリングが構造的なセキュリティ上の欠陥となるのはなぜですか?
Amazon Bedrock AgentCore はユーザー ID をどのように検証しますか?
DynamoDB、ナレッジベース、SaaS の統合をどのように保護するか?
エージェントレベルのフィルタリングが構造的なセキュリティ上の欠陥となるのはなぜですか?
AI エージェントの内部ロジックやプロンプト指示に依存してデータベース結果をフィルタリングすることは、重大な脆弱性です。攻撃者がプロンプト インジェクションを使用してエージェントのロジックをバイパスすると、エージェントの基盤となるデータセットがすべて公開されます。真のセキュリティには、LLM の推論とは関係なく、インフラストラクチャ層で認可を強制する必要があります。
従来の解決策は、AI エージェントに広範な管理資格情報を与え、結果をユーザーに返す前にエージェントが結果をフィルタリングすることを信頼することでした。 AWS はこれを根本的に欠陥のある設計であると認識しています。 AWS Well-Architected Agentic AI Lens の AGENTSEC03 ベストプラクティスでは、エージェントは純粋に または として機能する必要があると規定されています。

一方、アクセスの決定はダウンストリーム サービスによって厳密に適用されます。
Amazon Bedrock AgentCore はユーザー ID をどのように検証しますか?
Amazon Bedrock AgentCore は、カスタム部門クレームと AWS セッション タグで強化された JSON ウェブ トークン (JWT) を Amazon Cognito 経由でインターセプトすることでユーザー ID を検証します。 AgentCore ランタイムは到着時にこのトークンを検証し、AI エージェントのコードが実行を開始する前に不正な呼び出しリクエストをブロックします。
ユーザーがログインすると、トークン生成前の Lambda トリガーによって ID メタデータがトークンに直接挿入されます。そこから、AWS は、ダウンストリーム サービスに対して、インフラストラクチャによって強制される明確な ID 伝播パターンを実証し、エージェントの実行ロールが独自の直接のデータストア権限を保持しないようにします。
DynamoDB、ナレッジベース、SaaS の統合をどのように保護するか?
DynamoDB の場合、エージェントは AssumeRoleWithWebIdentity を使用してユーザー トークンをスコープ付き認証情報と交換し、IAM LeadingKeys 条件を有効にします。ナレッジベースの場合、部門メタデータフィルターが検索呼び出しに追加されます。 Salesforce の場合、RFC 8693 On-Behalf-Of トークン交換を利用してユーザー ID を安全に交換します。
公式 AWS セキュリティ ブログで詳しく説明されているように、インフラストラクチャ層で認可を強制すると、エージェントレベルの操作に関係なく、データ境界がそのまま維持されます。各リクエストには、有効期限が短く、暗号的に導出された資格情報が含まれますが、これらの資格情報はすぐに期限切れになり、悪意のあるプロンプトによって偽造することはできません。
プロンプトインジェクション攻撃は AWS インフラストラクチャの認証をバイパスできますか?
いいえ。アクセス制御は、暗号署名されたユーザー トークンに基づく IAM 条件とダウンストリーム SaaS ルールによって強制されるため、LLM プロンプトを操作しても、ユーザーの明示的な権限を超えて権限を昇格させることはできません。アマゾですか

n Bedrock AgentCore はサードパーティの ID プロバイダー (IdP) をサポートしますか?
はい、AWS アーキテクチャでは Amazon Cognito が強調されることが多いですが、トークン交換メカニズムとランタイム検証は標準の OIDC 準拠のサードパーティ ID プロバイダーと統合できます。 RFC 8693 On-Behalf-Of (OBO) トークン交換とは何ですか?
RFC 8693 は OAuth 2.0 拡張機能で、仲介者 (AI エージェントなど) が生の認証情報を公開することなく、受信ユーザー トークンをダウンストリーム サービス (Salesforce など) に合わせた新しいトークンと交換できるようにします。
カテゴリー Agentic AI & LLM タグ AI セキュリティ , Amazon Bedrock , AWS
GitLab GraphQL の重大な欠陥 (CVE-2026-19478) によりパブリック プロジェクトの削除が可能になる
コメントを残す 返信をキャンセル
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
AWS Bedrock AgentCore がハイジャックされた AI エージェントを阻止する方法
GitLab GraphQL の重大な欠陥 (CVE-2026-19478) によりパブリック プロジェクトの削除が可能になる
連鎖的な停止: AWS の手動変更による Terraform の状態ドリフトを修正する方法

## Original Extract

Learn how AWS Bedrock AgentCore stops hijacked AI agents and prompt injection by enforcing user authorization context and IAM identity propagation.

How AWS Bedrock AgentCore Stops Hijacked AI Agents
Enterprises are racing to deploy AI agents that pull from databases, document repositories, and SaaS platforms to automate workflows. But a quiet risk lurks beneath the convenience: a compromised agent manipulated via prompt injection could hand over data the requesting user was never authorized to see. We will show you how to close this gap by moving authorization out of the agent’s code and into the AWS infrastructure itself.
This architectural shift ensures that even a fully hijacked agent remains cryptographically constrained to the user’s specific permissions. Just as we discussed the catastrophic consequences of unauthorized data manipulation in our recent breakdown of the Critical GitLab GraphQL Flaw , relying on application-level filtering for AI access is a severe structural weakness.
Why is Agent-Level Filtering a Structural Security Flaw?
How Does Amazon Bedrock AgentCore Validate User Identity?
How Do We Secure DynamoDB, Knowledge Bases, and SaaS Integrations?
Why is Agent-Level Filtering a Structural Security Flaw?
Relying on an AI agent’s internal logic or prompt instructions to filter database results is a critical vulnerability. If an attacker bypasses the agent’s logic using prompt injection, the agent’s full underlying dataset is exposed. True security requires enforcing authorization at the infrastructure layer, independently of the LLM’s reasoning.
The traditional fix has been to give an AI agent broad administrative credentials and trust the agent to filter results before returning them to the user. AWS identifies this as a fundamentally flawed design. The AGENTSEC03 best practice in the AWS Well-Architected Agentic AI Lens dictates that the agent should purely act as an orchestrator, while access decisions are strictly enforced by the downstream services.
How Does Amazon Bedrock AgentCore Validate User Identity?
Amazon Bedrock AgentCore validates user identity by intercepting JSON Web Tokens (JWTs) enriched with custom department claims and AWS session tags via Amazon Cognito. The AgentCore Runtime verifies this token on arrival, blocking unauthorized invocation requests before the AI agent’s code even begins to execute .
When a user logs in, a pre-token generation Lambda trigger directly injects identity metadata into the token. From there, AWS demonstrates distinct infrastructure-enforced identity propagation patterns for downstream services, ensuring that the agent’s execution role holds no direct data-store permissions of its own.
How Do We Secure DynamoDB, Knowledge Bases, and SaaS Integrations?
For DynamoDB, the agent exchanges the user token for scoped credentials using AssumeRoleWithWebIdentity, enabling IAM LeadingKeys conditions. For Knowledge Bases, it appends department metadata filters to retrieval calls. For Salesforce, it utilizes RFC 8693 On-Behalf-Of token exchange to swap the user identity securely .
As detailed in the official AWS Security Blog , enforcing authorization at the infrastructure layer ensures data boundaries remain intact regardless of agent-level manipulation. Each request carries short-lived, cryptographically derived credentials that expire rapidly and cannot be forged by malicious prompts.
Can a prompt injection attack bypass AWS infrastructure authorization?
No. Because the access controls are enforced by IAM conditions and downstream SaaS rules based on cryptographically signed user tokens, manipulating the LLM prompt cannot escalate privileges beyond the user’s explicit permissions. Does Amazon Bedrock AgentCore support third-party Identity Providers (IdPs)?
Yes, while the AWS architecture often highlights Amazon Cognito, the token exchange mechanisms and runtime validations can be integrated with standard OIDC-compliant third-party identity providers. What is RFC 8693 On-Behalf-Of (OBO) token exchange?
RFC 8693 is an OAuth 2.0 extension that allows an intermediary (like an AI agent) to exchange an incoming user token for a new token tailored to a downstream service (like Salesforce), without exposing raw credentials.
Categories Agentic AI & LLMs Tags AI Security , Amazon Bedrock , AWS
Critical GitLab GraphQL Flaw (CVE-2026-19478) Allows Public Project Deletion
Leave a Comment Cancel reply
Save my name, email, and website in this browser for the next time I comment.
How AWS Bedrock AgentCore Stops Hijacked AI Agents
Critical GitLab GraphQL Flaw (CVE-2026-19478) Allows Public Project Deletion
The Cascading Outage: How to Fix Terraform State Drift from Manual AWS Changes
