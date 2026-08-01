---
source: "https://cloud.google.com/blog/topics/developers-practitioners/generosity-under-conditions-hardening-google-cloud-access-management/"
hn_url: "https://news.ycombinator.com/item?id=49131063"
title: "Hardening Google Cloud IAM with CEL Conditions and Deny Policies"
article_title: "Generosity Under Conditions: Hardening Google Cloud Access Management | Google Cloud Blog"
author: "minherz"
captured_at: "2026-08-01T05:18:53Z"
capture_tool: "hn-digest"
hn_id: 49131063
score: 1
comments: 0
posted_at: "2026-08-01T04:33:50Z"
tags:
  - hacker-news
  - translated
---

# Hardening Google Cloud IAM with CEL Conditions and Deny Policies

- HN: [49131063](https://news.ycombinator.com/item?id=49131063)
- Source: [cloud.google.com](https://cloud.google.com/blog/topics/developers-practitioners/generosity-under-conditions-hardening-google-cloud-access-management/)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T04:33:50Z

## Translation

タイトル: CEL 条件と拒否ポリシーによる Google Cloud IAM の強化
記事のタイトル: 条件下の寛大さ: Google Cloud アクセス管理の強化 | Google クラウド ブログ
説明: IAM 条件、CEL、拒否ポリシーを使用して Google Cloud のアクセス管理を強化し、環境内で最小権限の原則を効果的に適用する方法を学びます。

記事本文:
条件下の寛大さ: Google Cloud アクセス管理の強化 | Google Cloud ブログ コンテンツに移動 クラウド ブログ 営業担当者へのお問い合わせ 無料で始めましょう クラウド ブログ ソリューションとテクノロジー AI と機械学習
開発者と実務者の条件下での寛大さ: Google Cloud アクセス管理の強化
シニア デベロッパー リレーションズ エンジニア
Google Cloud では、Identity and Access Management（IAM）は、クラウド リソースと運用に対するアクセス制御を維持するのに役立ちます。他の機能も含まれていますが、これが主な目的です。アプリケーションのセキュリティを強化しようとしたことがあるなら、最小特権の原則 (PoLP) の重要性をご存知でしょう。つまり、ユーザーとワークロードにタスクの実行を許可する絶対最小限の権限を付与するということです。これには、事前定義されたロールとカスタム ロールを使用し、プロジェクト、フォルダー、または組織レベルで許可と拒否の IAM ポリシーの組み合わせを設定することで達成できます。リソース階層に沿って許可ポリシーと拒否ポリシーを組み合わせて使用​​することは、アクセスを制御する効果的な方法です。このアプローチにより、さまざまなシナリオにわたって PoLP を適用できます。
プロジェクト内のリソースが複数のワークロード間で共有されている場合、または複数のチームによって使用されている場合、既存の柔軟な制御では不十分になる可能性があります。このようなシナリオの多くでは、IAM ポリシーをプロジェクト内の特定のリソースにバインドできます。たとえば、プロジェクトにロール Artifact Registry Editor (roles/artifactregistry.editor) を付与する場合と、プロジェクト内の特定のリポジトリにロールを付与する場合の違いを考えてみましょう。前者の場合、プロジェクト内の任意のリポジトリへのアクセスが許可されます。後者の場合、ユーザーは特定のリポジトリへの編集者アクセスのみを持ちます。ただし、IAM ポリシーをリソースまたはサービス レベルにバインドすることはできません。

いつでも可能なわけではありません。このような場合に IAM 条件を使用します。アクセス管理を強化する際の条件の力を示す 2 つの異なる例を見てみましょう。1 つは従来の管理役割に関するもので、もう 1 つは最新の AI 統合に関するものです。
使用例 1: 管理者の権限を制限する
このケースでは、広範な IAM ロールが実行を許可されている特定の操作を制限する方法を示します。プロジェクト レベルで「リソース作成者」ロールを付与し、選択したリソースに対して編集者ロールを付与することで、プロジェクト内の特定のリソースを管理するための管理者権限の範囲を簡単に設定できます。特定のリソースではなく、操作へのアクセスを許可することを目的とした IAM 管理者ロールを制限することは、はるかに困難です。代表的な例は、IAM 管理者ロール (roles/iam.admin) です。このロールを付与されたユーザーは、自分自身に他のロールを付与したり、新しいロールを作成したりできます。それは実際のニーズを大幅に超えています。最初のステップは、プロジェクトのレベルでのみ管理者権限を提供するプロジェクト IAM 管理者ロール (roles/resourcemanager.projectIamAdmin ) を使用してアクセスを制限することです。
ただし、付与される権限をさらに制限することも可能です。たとえば、リソースを作成してワークロードをデプロイするビルダー サービス アカウントにプロジェクト IAM 管理者ロールを付与するとします。ワークロードに必要なのは、BigQuery およびエージェント プラットフォーム API (旧称 Vertex API) へのアクセスと、ログとトレースを書き込む権限のみです。このような場合は、次の gcloud CLI コマンド、または Terraform の代替コマンドを使用できます。
読み込み中... gcloud プロジェクト add-iam-policy-binding "${PROJECT_ID}" \
--member="サービスアカウント:${SA_MAIL}" \
--role="roles/resourcemanager.projectIamAdmin" \
--condition="^:^\
title=限定IAMAdmin:\
式=api.getAttribute('iam.goo

gleapis.com/modifiedGrantsByRole', [])\
.hasOnly([\
'roles/aiplatform.user',\
'roles/bigquery.jobUser',\
'roles/bigquery.dataViewer',\
'roles/cloudtrace.agent',\
'roles/logging.logWriter'\
])" 条件パラメータの値は、共通表現言語 (CEL) 構文を使用して定義されます。まず、フィールド区切り文字がコンマではなくコロンになるようにカスタマイズされ、次に条件フィールド title とexpression が記述されます。expression フィールドは、API 属性の関数を使用して付与されているロールを識別し、カンマ区切りリスト内のロールのみを付与できるようにします。Terraform での同じ操作は非常によく似ています。環境変数の代わりに入力変数を使用すると、次のようになります。これ:
読み込み中... リソース "google_project_iam_member" "limited_project_iam_admin" {
プロジェクト = var.project_id
役割 = "roles/resourcemanager.projectIamAdmin"
メンバー = "serviceAccount:${var.sa_email}"
条件 {
title = "限定IAM管理者"
式 = <<-EOT
api.getAttribute('iam.googleapis.com/modifiedGrantsByRole', []).hasOnly([
'roles/aiplatform.user',
'roles/bigquery.jobUser',
'roles/bigquery.dataViewer',
'roles/cloudtrace.agent',
'roles/logging.logWriter'
])
EOT
}
使用例 2: MCP サーバー アクセスの制御
このケースは、単一の権限セットの背後にある特定のサービスへのアクセスを強化することに関するものです。
Google は、モデル コンテキスト プロトコル (MCP) エンドポイントを公開する MCP サーバーを介して、クラウド リソースとサービスのサブセットへのアクセスを公開します。これらのサーバーへのアクセスは、事前定義された MCP ツール ユーザー (roles/mcp.toolUser) ロールを使用して付与されます。このロールは、使用可能なすべての MCP サーバーへのアクセスを許可します (IAM ポリシーが設定されているプロジェクトの場合)。条件を使用すると、特定の MCP サーバーへのアクセスを絞り込むことができます。
読み込み中... gcloud プロジェクト add-iam-policy-binding $PROJECT_ID \
--メンバー=

"サービスアカウント:$SA_EMAIL" \
--role="roles/mcp.toolUser" \
--condition="^:^\
title=bigquery_mcp_server_only:\
Expression=resource.service == 'bigquery.googleapis.com'" resource.service 属性と比較される値は、MCP サーバーのエンドポイント ( bigquery.googleapis.com/mcp ) ではなく、サービスのエンドポイントであることに注意してください。アクセス範囲をさらに特定の MCP ツールのレベルまで狭めることができます。このためには、API 属性を再度使用する必要があります。次の式は、サービス アカウントのアクセスを 2 つの BigQuery MCP ツールのみのレベルに制限します。
読み込み中...expression=api.getAttribute('mcp.googleapis.com/tool.name', '') in [\
'mcp_bigquery-mcp_execute_sql',\
'mcp_bigquery-mcp_execute_sql_readonly'\
] MCP ツール レベルで IAM ポリシー バインディングを条件付けする場合は、resource.service 属性を検証する必要がないことに注意してください。
MCP サーバー アクセスを実験するには、「Getting Started with Google MCP Servers」コードラボを使用し、その gcloud プロジェクトの add-iam-policy-binding コマンドを変更できます。
IAM 条件を使用すると、事前定義されたロールを使用するときに正確な制御を強制できるほか、リクエストの時間に基づいてアクセス管理を作成できます。たとえば、次の条件式は平日の日中のみアクセスを許可します。
読み込み中...expression=request.time.getHours('ヨーロッパ/ベルリン') >= 9 &&\
request.time.getHours('ヨーロッパ/ベルリン') <= 17 &&\
request.time.getDayOfWeek('ヨーロッパ/ベルリン') >= 1 &&\
request.time.getDayOfWeek('Europe/Berlin') <= 5 この式は、月曜日から金曜日までの「ヨーロッパ/ベルリン」タイムゾーンに従って、朝の 9 時から夕方 5 時までのアクセスを制限します (曜日の範囲は日曜日から始まり 0 から 6 です)。
IAM 条件により、プリンシパル属性を使用してアクターの ID を制御できるようになります

s 。ただし、それは簡単にアンチパターンになる可能性があります。推奨される方法は、条件を使用するのではなく、IAM ポリシーのプリンシパルのリストを通じて、ポリシーの使用を許可されるアクターの ID を制御することです。
IAM 条件により、許可ポリシーよりも正確な精度が得られますが、IAM 拒否ポリシーを使用すると、多層防御戦略をさらに進めることができます。拒否ポリシーを使用すると、許可ポリシーを持つ事前定義された IAM ロールを使用してアクセスを許可し、ロールの過剰な権限を削除して PoLP を強制できます。拒否ポリシーの詳細については、次のリソースを参照してください。
拒否ポリシーでサポートされている権限を特定します。
拒否ポリシーのプリンシパル識別子の形式を取得します。
拒否ポリシーを使用したアクセス問題のトラブルシューティング方法を確認してください。
プリンシパルへのアクセスの拒否について詳しくは、こちらをご覧ください。
多層防御の構築に関するブログ投稿をお読みください。
Google スキルを使用して、IAM ポリシーを実際に体験できます。
コーディング エージェントを使用してエージェント開発ライフサイクルを自動化する
シャバム・サブー著 • 7 分で読めます
AI アプリが本番環境で失敗する理由 (そして Google がそれをどのように解決したか)
ステファニー・ウォン著 • 4 分で読めます
Gemini Enterprise Agent Platform 上に構築する 13 の実践的なデモ
シャバム・サブー著 • 6 分で読めます
AI トークンノミクス ガイド: トークンの効率的なソフトウェア エンジニアリングのための 11 の原則
Alex "Sandu" Astrum著 • 4 分で読めます
言語 英語 ドイツ語 フランス語 한국어 日本語

## Original Extract

Learn how to harden Google Cloud access management using IAM conditions, CEL, and Deny policies to effectively enforce the principle of least privilege in your environment.

Generosity Under Conditions: Hardening Google Cloud Access Management | Google Cloud Blog Jump to Content Cloud Blog Contact sales Get started for free Cloud Blog Solutions & technology AI & Machine Learning
Developers & Practitioners Generosity Under Conditions: Hardening Google Cloud Access Management
Senior Developer Relations Engineer
In Google Cloud, Identity and Access Management (IAM) helps you maintain access control over your cloud resources and operations. While it includes other features, this is its primary purpose. If you ever tried to harden security over your application, you know the importance of the Principle of Least Privilege ( PoLP ) ‒ grant the absolute minimum permissions to your users and workloads to allow them to perform their tasks. You reach it through use of predefined roles and custom roles and setting up a combination of Allow and Deny IAM policies at project, folder, or organization level. Using a combination of Allow and Deny policies along the resource hierarchy is an effective way to control access. This approach lets you enforce PoLP across many different scenarios.
The existing flexible control can be insufficient when resources in the project are shared between multiple workloads or used by more than one team. In many such scenarios, it is possible to bind IAM policies to a specific resource in the project. For example, consider the difference between granting the role Artifact Registry Editor ( roles/artifactregistry.editor ) on a project vs. granting it on a specific repository in the project. In the former case, the access is granted to ANY repository in the project. In the latter case, users will have the editor access only to a specific repository. However, binding IAM policies to a resource or service level isn't always possible. This is when it is time to use IAM conditions . Let’s look at two distinct examples that demonstrate the power of conditions when hardening access management: one for traditional administrative roles, and one for modern AI integrations.
Use Case 1: Constraining the Power of Admins
This case demonstrates how to restrict the specific operations that broad IAM roles are authorized to perform. You can easily scope administrative privileges for managing specific resources in a project by granting a "resource creator" role at the project level and an editor role on a selected resource. It is far more challenging to constrain IAM Admin Roles that are intended to grant access to operations rather than specific resources. A representative example would be the IAM Admin role ( roles/iam.admin ). Users granted this role can grant themselves any other role or create a new one. It greatly exceeds practical needs. The first step is to narrow the access by using the Project IAM Admin role ( roles/resourcemanager.projectIamAdmin ) that provides administrative privileges only at the level of the project.
It is possible, however, to restrict the granted privileges even further. For example, suppose you grant the Project IAM Admin role to your builder service account that creates resources and deploys workloads. The workloads only need access to the BigQuery and Agent Platform APIs (formerly Vertex APIs) and permission to write logs and traces. For such a case you can use the following gcloud CLI command or its alternative in Terraform:
Loading... gcloud projects add-iam-policy-binding "${PROJECT_ID}" \
--member="serviceAccount:${SA_MAIL}" \
--role="roles/resourcemanager.projectIamAdmin" \
--condition="^:^\
title=LimitedIAMAdmin:\
expression=api.getAttribute('iam.googleapis.com/modifiedGrantsByRole', [])\
.hasOnly([\
'roles/aiplatform.user',\
'roles/bigquery.jobUser',\
'roles/bigquery.dataViewer',\
'roles/cloudtrace.agent',\
'roles/logging.logWriter'\
])" The value of the condition parameter is defined using Common Expression Language ( CEL ) syntax . First it customizes a field delimiter to be a colon instead of a comma and then describes the condition fields title and expression . The expression field uses functions for API attributes to identify which roles are being granted to allow granting only the roles in the comma delimited list. The same operation in Terraform will look very similar. Using input variables instead of environment variables, it will look like this:
Loading... resource "google_project_iam_member" "limited_project_iam_admin" {
project = var.project_id
role = "roles/resourcemanager.projectIamAdmin"
member = "serviceAccount:${var.sa_email}"
condition {
title = "LimitedIAMAdmin"
expression = <<-EOT
api.getAttribute('iam.googleapis.com/modifiedGrantsByRole', []).hasOnly([
'roles/aiplatform.user',
'roles/bigquery.jobUser',
'roles/bigquery.dataViewer',
'roles/cloudtrace.agent',
'roles/logging.logWriter'
])
EOT
}
} Use Case 2: Control over MCP Server Access
This case is about hardening access to specific services behind a single set of permissions.
Google exposes access to a subset of cloud resources and services via MCP Servers that expose Model Context Protocol (MCP) endpoints. The access to these servers is granted using the predefined MCP Tool User ( roles/mcp.toolUser ) role. This role grants access to ALL available MCP servers (for a project where an IAM policy is set). Using conditions helps to narrow the access to a specific MCP server.
Loading... gcloud projects add-iam-policy-binding $PROJECT_ID \
--member="serviceAccount:$SA_EMAIL" \
--role="roles/mcp.toolUser" \
--condition="^:^\
title=bigquery_mcp_server_only:\
expression=resource.service == 'bigquery.googleapis.com'" Notice that the value compared to the resource.service attribute is not the MCP server endpoint (which is bigquery.googleapis.com/mcp ) but the endpoint of the service. It is possible to narrow the access scope further to the level of the specific MCP tools. For this you will need to use API attributes again. The following expression limits the service account access to the level of only two BigQuery MCP tools.
Loading... expression=api.getAttribute('mcp.googleapis.com/tool.name', '') in [\
'mcp_bigquery-mcp_execute_sql',\
'mcp_bigquery-mcp_execute_sql_readonly'\
] Note that if you condition the IAM policy binding at the MCP tool level, you don't need to validate the resource.service attribute.
For experimenting with MCP server access you can use the Getting Started with Google MCP Servers codelab and modify its gcloud projects add-iam-policy-binding commands.
Besides enforcing precise control when using predefined roles, IAM conditions let you craft access management based on the time of the request. For example, the following condition's expression allows access only during daytime on weekdays:
Loading... expression=request.time.getHours('Europe/Berlin') >= 9 &&\
request.time.getHours('Europe/Berlin') <= 17 &&\
request.time.getDayOfWeek('Europe/Berlin') >= 1 &&\
request.time.getDayOfWeek('Europe/Berlin') <= 5 The expression limits access from 9 o'clock in the morning to 5 o'clock in the evening according to the "Europe/Berlin" timezone from Monday to Friday (days of the week range from 0 to 6, starting with Sunday).
IAM conditions allow controlling the identity of the actor using the principal attributes . However, it can easily become an anti-pattern. The recommended practice is to control the identity of actors allowed to use the policy through the list of the IAM policy's principals instead of using the conditions.
While IAM conditions give you surgical precision over Allow policies, you can take your defense-in-depth strategy even further with IAM Deny policies . With Deny Policies you can grant access using the predefined IAM roles with Allow policies and remove excessive permissions of the role to enforce PoLP. See the following resources for additional information about Deny policies:
Identify the permissions that are supported in deny policies .
Get the format of principal identifiers in deny policies .
Find out how to troubleshoot access issues with deny policies .
Learn more about denying access to principals .
Read the blog post about Build defense in depth .
You can use Google Skills for hands-on experience with IAM policies.
Automate your agent development lifecycle using any coding agent
By Shubham Saboo • 7-minute read
Why AI apps fail in production (And how Google solved it)
By Stephanie Wong • 4-minute read
13 hands-on demos to build on Gemini Enterprise Agent Platform
By Shubham Saboo • 6-minute read
Guide to AI Tokenomics: Eleven Principles for Token Efficient Software Engineering
By Alex "Sandu" Astrum • 4-minute read
Language ‪English‬ ‪Deutsch‬ ‪Français‬ ‪한국어‬ ‪日本語‬
