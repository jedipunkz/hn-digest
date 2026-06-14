---
source: "https://learn.microsoft.com/en-us/microsoft-365/copilot/connect-to-ai-subprocessor"
hn_url: "https://news.ycombinator.com/item?id=48530288"
title: "Anthropic Models in Microsoft Online Services"
article_title: "Anthropic models in Microsoft Online Services | Microsoft Learn"
author: "sntran"
captured_at: "2026-06-14T18:36:09Z"
capture_tool: "hn-digest"
hn_id: 48530288
score: 2
comments: 0
posted_at: "2026-06-14T17:47:04Z"
tags:
  - hacker-news
  - translated
---

# Anthropic Models in Microsoft Online Services

- HN: [48530288](https://news.ycombinator.com/item?id=48530288)
- Source: [learn.microsoft.com](https://learn.microsoft.com/en-us/microsoft-365/copilot/connect-to-ai-subprocessor)
- Score: 2
- Comments: 0
- Posted: 2026-06-14T17:47:04Z

## Translation

タイトル: Microsoft Online Services における人類モデル
記事のタイトル: Microsoft Online Services における人類モデル | Microsoft Learn
説明: Microsoft Online Services の人体モデルについて学習します。

記事本文:
メインコンテンツにスキップ
「質問する」にスキップしてチャット体験を学習する
このブラウザはサポートされなくなりました。
Microsoft Edge にアップグレードして、最新の機能、セキュリティ更新プログラム、テクニカル サポートを利用してください。
このページにアクセスするには認証が必要です。サインインするか、ディレクトリを変更してみてください。
このページにアクセスするには認証が必要です。ディレクトリを変更してみてください。
Microsoft Online Services の人体モデル
適用対象: ✅ Microsoft 365 コパイロット
Microsoft は、Microsoft Online Services の一部として Anthropic AI モデルを備えた新しい製品を導入し、組織内で Anthropic モデルを安全かつ責任を持って使用できるようにするエンタープライズ グレードのコミットメントと保護手段を提供します。
この変更を有効にするために、Anthropic は Microsoft サブプロセッサとして採用されました。この変更により、Microsoft のエンタープライズ フレームワークの下でエクスペリエンスが簡素化され、コンプライアンスとセキュリティが強化されます。 Microsoft 顧客著作権コミットメント (CCC) は、Microsoft 365 Copilot や Copilot Studio など、CCC の対象となる製品内で使用される Anthropic モデルに適用されます。
Anthropic は副処理者として、契約上の保護措置と適切な技術的および組織的手段を通じて Microsoft の監督を受けて業務を遂行します。モデルに「データ保持付きプレビュー モデル」というラベルが付いている場合を除き、Microsoft のエンタープライズ オンライン サービスを通じた Anthropic モデルの使用には、Microsoft 製品条項および Microsoft Data Protection Addendum (DPA) が適用されます。このような使用も、当社のエンタープライズ データ保護の対象となります。
サブプロセッサのデータ アクセスの詳細については、「 Microsoft データ アクセス管理 」を参照してください。 Microsoft 副処理者のリストを確認するには、Service Trust Portal を参照してください。
Microsoft は、商用クラウドのほとんどの顧客に対して、Anthropic モデルをデフォルトで有効にします (EU/EFTA および英国を除く)。これまで

この日付により、組織内のユーザーは、Microsoft 365 Copilot、Researcher、Copilot Studio、Power Platform、Microsoft 365 アプリの Copilot など、Microsoft 製品で複数の AI モデルを使用できるようになります。これは、エンタープライズ グレードのセキュリティとコンプライアンスを維持しながら、主要な AI モデルから選択肢を提供するという Microsoft の取り組みを裏付けるものです。
Microsoft 365 Copilot、Researcher、Copilot Studio、Power Platform、Microsoft 365 アプリの Copilot などの Microsoft 製品にデプロイされている人為モデルは、現在 EU データ境界から除外されており、該当する場合は国内処理コミットメントからも除外されています。 EU データ境界内の顧客および英国の顧客は、デフォルトで Anthropic モデルが無効になります。
Anthropic モデルは現在、ガバメント クラウド (GCC、GCC High、DoD) またはソブリン クラウドでは使用できません。
サブプロセッサとしての Anthropic は段階的に導入されており、まだすべての組織で利用できるわけではありません。この段階的なロールアウト中は、組織によっては一部の機能が制限される場合があります。 2026 年 3 月末までに完全に利用可能になる予定です。
時折、Anthropic の一部の新しく利用可能な高度なモデルは、Microsoft テナント管理者が Anthropic の別個の商業条件およびデータ処理契約に基づいてプレビュー モデルの使用をオプトインできる別個のコントロールを使用して引き続き利用できる場合があります。たとえば、データ保持を備えた Anthropic Preview モデルには異なる条件が適用され、他の Anthropic モデルがデフォルトでオンになっている場合でも、これらのモデルは常に個別に有効にする必要があります。詳細については、「Microsoft Online Services で AI モデルをプレビューする」を参照してください。
データ保持を備えた Anthropic Preview モデル
一部の Anthropic モデルは、Anthropic による顧客データの保持および保管の対象となります (「データ保持付きプレビュー モデル」)。のために

これらのモデルでは、Anthropic は Microsoft のサブプロセッサではなく、独立したデータ プロセッサとして機能します。データ保持機能を備えたプレビュー モデルは、一部の Microsoft オンライン サービス内から使用できますが、Anthropic の商用利用規約および Anthropic のデータ保護補遺が適用されます。組織によるデータ保持付き Anthropic Preview モデルの使用はオプションです。データ保持を備えたプレビュー モデルには、Microsoft 365 管理センターで追加の制御があり、商用クラウドでさまざまな Anthropic モデルがデフォルトでオンになっている地域や、Anthropic が Microsoft サブプロセッサである地域など、すべてのシナリオでデフォルトでオフのままです。
Microsoft 365 管理センターで Anthropic のクロード モデル設定を管理する
Microsoft は、特定の地域で Anthropic モデルをデフォルトで利用できるようにしています。 Microsoft 365 Copilot (Web、デスクトップ、モバイル) では、クロード モデルが使用されているときに UI インジケーターが表示されます。 Copilot Studio では、作成者はエージェントの作成時にモデルを選択する必要があります。 Microsoft 365 アプリや Researcher の Copilot を使用した編集などの機能では、ユーザーは クロード を選択できます。
一部の地域では、Anthropic のモデルはデフォルトでは利用できません。これらの領域では、トグルが表示されますが、デフォルトは Off に設定されています。これらの地域には、欧州連合 (EU)、欧州自由貿易連合 (EFTA)、および英国 (UK) が含まれます。
さらに、FedRAMP 認定がまだ導入されていないため、Anthropic モデルはガバメント クラウド (GCC、GCC High、DoD) では利用できません。また、他のソブリン クラウドでも利用できません。ガバメント クラウドまたはソブリン クラウドにはトグルは存在しません。
2026 年 4 月 3 日、Microsoft は EU/EFTA および英国で Anthropic モデルを使用した M365 アプリの新しい Microsoft 365 管理センター設定 Copilot を導入し、Anthropic をデフォルト モデルとして有効にしました。

Microsoft 365 アプリのコパイロット。詳細については、「人為的モデルを使用した Microsoft 365 アプリのコパイロット」を参照してください。
Anthropic のモデルを使用するためにオプトインします
サブプロセッサーとして Anthropic がデフォルトでオフに設定されている地域に組織がある場合は、Anthropic のモデルを組織で利用できるようにオプトインすることを選択できます。このタスクを実行するには、グローバル管理者ロールのメンバーである必要があります。詳細については、「管理者の役割について」および「全体管理者」を参照してください。
Microsoft 365 管理センターに移動し、[Copilot] -> [設定] -> [すべて表示] を選択します。
Microsoft サブプロセッサとして動作する AI プロバイダーを選択します。
[Microsoft サブプロセッサとして動作している AI プロバイダー] ページの [組織で利用可能なサブプロセッサ] で、 [Anthropic] を選択し、 [保存] を選択します。
[Copilot および生成 AI エクスペリエンスのための人体モデルにアクセスできるユーザーの選択] で、ユーザーまたはグループを選択し、保存します。
Microsoft 365 管理センターで特定のユーザーまたは Microsoft Entra ID セキュリティ グループにアクセス許可を割り当てることで、AI プロバイダーのサブプロセッサへのユーザー アクセスを制限できます。これらの割り当てはプロバイダー レベルで適用され、Microsoft 365 Copilot および Copilot Studio エクスペリエンス全体に適用されます。ユーザーまたはグループのメンバーシップによってアクセスが制限されている場合、割り当てられたユーザーのみが、その AI プロバイダーに依存する Copilot 機能またはエージェントを使用できます。既存のユーザーまたはグループの割り当てを確認し、必要に応じてポリシーまたは構成を更新します。ユーザーとセキュリティ グループのアクセスの詳細については、「 AI プロバイダーのアクセスをユーザーとグループに割り当てる 」を参照してください。セキュリティ グループの作成の詳細については、「セキュリティ グループの作成」を参照してください。
組織が欧州連合 (EU)、欧州自由貿易連合 (EFTA)、または英国に所属しており、以前に Anthropic の sep に基づいて Anthropic モデルを使用することをオプトインしている場合

商業条件とデータ処理契約を確認するには、再度オプトインする必要があります。トグルはデフォルトでオフに設定されます。
一部の機能は、人間モデルが有効になっている場合にのみ使用できます。 Anthropic をサブプロセッサとしてオフにすると、特定の機能にアクセスできなくなる可能性があります。
データ保持付きのプレビュー モデルを含むプレビュー モデルの使用を許可するオプトイン
プレビュー モデルは、探索とテストには利用できる最新のモデルですが、運用環境での使用は推奨されません。テナントのデータ保持を備えた Anthropic Preview モデルの使用を有効にする前に、プレビュー モデルの制限を確認してください。
データ保持を備えた Anthropic Preview モデルは、Anthropic によるデータ保持を必要とする高度なモデル (Claude Fable 5 や Claude Mythos 5 など) です。これは、データが Anthropic によって保存され、製品条件や DPA での約束を含む Microsoft 顧客契約の対象ではないことを意味します。データ保持を備えた Anthropic プレビュー モデルは、Anthropic がサブプロセッサとして有効になっている場合でも、テナントではデフォルトでオフになっています。テナント管理者がデータ保持付きのプレビュー モデルの使用を有効にすることを明示的にオプトインするまで、ユーザーはこれらのモデルにアクセスできません。データ保持付きの Anthropic Preview モデルを使用するには、Anthropic の商用利用規約および Anthropic のデータ保護補遺に同意する必要があります。テナント管理者は、どのユーザー (すべてまたは特定のユーザーまたはグループ) がこれらのモデルにアクセスして使用できるかを選択する必要もあります。
組織内のユーザーにデータ保持機能を備えたプレビュー モデルへのアクセスが許可されると、ユーザーは製品エクスペリエンスで、これらのモデルが Anthropic によるデータ保持を必要とすることを識別できるようになります。
データ保持付きの Anthropic プレビュー モデルの使用を有効にすると、Anthropic のコマーシャルで説明されているように、Anthropic はデータを保持します。

ata 保持ポリシー: Mythos クラス モデルのデータ保持 これは、Anthropic (Microsoft ではなく) がほとんどの入力と出力を削除するまで最大 30 日間保存することを意味します。 Anthropic の信頼および安全性分類子が Anthropic の使用ポリシーの潜在的な違反を特定した場合、Anthropic はコンテンツ (入力および出力) を最長 2 年間、信頼性および安全性分類スコアを最長 7 年間保持することがあります。 Anthropic は、ユーザーの明示的な許可なしに、保持されたデータをモデルのトレーニングに使用しません。詳細については、「API とデータ保持 - Claude API ドキュメント」を参照してください。
Power Platform 管理センターでの Copilot Studio と Power Platform の追加コントロール
Microsoft 365 管理センターで有効にすると、Microsoft Power Platform 管理センター (PPAC) で追加の管理コントロールが利用可能になり、Copilot Studio および Power Platform で Anthropic を使用できるようになります。詳細については、「生成応答に外部ラージ言語モデル (LLM) を許可する」を参照してください。
Anthropic のモデルへの接続を無効にする
サブプロセッサとしての Anthropic が既定で [オン] に設定されているリージョンに組織が存在する場合は、Microsoft 365 管理センターで Anthropic モデルを制限することをオプトアウトできます。このタスクを実行するには、グローバル管理者ロールのメンバーである必要があります。詳細については、「管理者の役割について」および「全体管理者」を参照してください。
Microsoft 365 管理センターに移動し、[Copilot] -> [設定] を選択します。
[ユーザー アクセス] ページで、Microsoft サブプロセッサとして動作する AI プロバイダーを選択します。
[Microsoft サブプロセッサとして動作している AI プロバイダー] ページの [組織で利用可能なサブプロセッサ] で、 [Microsoft サブプロセッサとして Anthropic を無効にする] を選択します。
AI サブプロセッサとしての Anthropic を無効にすると、ユーザーは Anthropic の AI モデルを使用できなくなります。 Anthropic モデルを有効にするかどうかを選択できます

ご希望に応じて後日。
従来の Anthropic 管理者の非推奨トグル
Anthropic の個別の商業条件とデータ処理契約にオプトインする従来の Anthropic トグルは廃止され、この新しい Anthropic のサブプロセッサ管理者トグルに置き換えられました。
2025 年 12 月 8 日: Anthropic モデルの新しい管理者切り替え機能が Microsoft 365 管理センターに表示されます。ほとんどの商用クラウド顧客の場合、デフォルトでオンになります (ただし、EU/EFTA および英国の顧客はデフォルトでオフになります)。一部の顧客および地域の除外が適用されます。
2026 年 1 月 7 日: Microsoft サブプロセッサとしての Anthropic が組織内で有効になります。 Anthropic の商業条件とデータ処理契約にオプトインするための従来の管理者切り替え機能は廃止されました。
以前に従来の Anthropic トグルをオプトインしており、新しいトグルがデフォルトでオフに設定されている地域にいる場合、Anthropic のモデルを使用するには、新しいサブプロセッサ トグルをオプトインする必要があります。
Microsoft Online Services の AI 機能とモデルを理解する
Microsoft 365 Copilot のデータ、プライバシー、セキュリティ
Microsoft 365 Copilot の AI サブプロセッサの概要
サプライヤーのセキュリティとプライバシーの保証
Ask Learnを使ってクラリフしてみたい

[切り捨てられた]

## Original Extract

Learn about Anthropic models in Microsoft Online Services.

Skip to main content
Skip to Ask Learn chat experience
This browser is no longer supported.
Upgrade to Microsoft Edge to take advantage of the latest features, security updates, and technical support.
Access to this page requires authorization. You can try signing in or changing directories .
Access to this page requires authorization. You can try changing directories .
Anthropic models in Microsoft Online Services
Applies to: ✅ Microsoft 365 Copilot
Microsoft is introducing a new offering with Anthropic AI models as part of Microsoft Online Services, delivering enterprise-grade commitments and safeguards to ensure secure and responsible use of Anthropic models within your organization.
To enable this change, Anthropic has onboarded as a Microsoft subprocessor. This change simplifies the experience and strengthens compliance and security under Microsoft's enterprise framework. The Microsoft Customer Copyright Commitment (CCC) applies to Anthropic models used within products covered by the CCC, including Microsoft 365 Copilot and Copilot Studio.
As a subprocessor, Anthropic will operate with Microsoft oversight through contractual safeguards and appropriate technical and organizational measures. Unless models are labeled "Preview models with Data Retention," the Microsoft Product Terms and Microsoft Data Protection Addendum (DPA) apply to use of Anthropic models through Microsoft's enterprise Online Services. Such use is also covered under our Enterprise Data Protection .
For more information about subprocessor data access, see Microsoft Data Access Management . To see a list of Microsoft subprocessors, see Service Trust Portal .
Microsoft will enable Anthropic models on by default for most customers in commercial cloud (excluding EU/EFTA and UK). This update gives users in your organization the ability to use multiple AI models in their Microsoft offerings, such as in Microsoft 365 Copilot, Researcher, Copilot Studio, Power Platform, and Copilot in Microsoft 365 apps. This affirms Microsoft's commitment to offering choice between leading AI models while maintaining enterprise-grade security and compliance.
Anthropic models deployed in Microsoft offerings such as Microsoft 365 Copilot, Researcher, Copilot Studio, Power Platform, and Copilot in Microsoft 365 apps are currently excluded from the EU Data Boundary, and when applicable, in-country processing commitments. Customers within the EU Data Boundary and customers in the UK will have Anthropic models disabled by default.
Anthropic models aren't currently available for use in government clouds (GCC, GCC High, DoD) or sovereign clouds.
Anthropic as a subprocessor is being introduced gradually and isn't yet available to all organizations. During this phased rollout, some features may be limited for your organization. Full availability is expected by the end of March 2026.
From time to time, some newly available and advanced models from Anthropic may still be made available with separate controls that allow Microsoft tenant admins to opt in to use preview models under Anthropic's separate commercial terms and data processing agreement . For example, different terms apply to Anthropic Preview models with Data Retention, and these models must always be separately enabled even where other Anthropic models are on by default. For more information, see Preview AI models in Microsoft Online Services .
Anthropic Preview models with Data Retention
Some Anthropic models are subject to retention and storage of Customer Data by Anthropic ("Preview models with Data Retention"). For these models, Anthropic acts as an independent data processor, not a Microsoft subprocessor. Preview models with Data Retention are made available for use from within some Microsoft Online Services but subject to Anthropic's Commercial Terms of Service and Anthropic's Data Protection Addendum . Your organization's use of Anthropic Preview models with Data Retention is optional . Preview models with Data Retention have additional controls in the Microsoft 365 admin center and remain default-off for all scenarios, including in regions when different Anthropic models are on by default in commercial cloud and for which Anthropic is a Microsoft subprocessor.
Manage Anthropic's Claude model settings in the Microsoft 365 admin center
Microsoft is making Anthropic models available by default in certain regions. In Microsoft 365 Copilot (web, desktop, and mobile), UI indicators will show when Claude models are in use. In Copilot Studio, creators must select the model during agent creation. In capabilities such as Edit with Copilot in Microsoft 365 apps or Researcher, users can select Claude .
In some regions, Anthropic's models aren't available by default. For these regions, the toggle will appear but the default is set to Off . These regions include the European Union (EU), the European Free Trade Association (EFTA) , and the United Kingdom (UK).
In addition, Anthropic models aren't available in government clouds (GCC, GCC High, DoD) as there's no FedRAMP certification in place yet. They're also not available in other sovereign clouds. No toggle will be present for government or sovereign clouds.
On April 3, 2026, Microsoft introduced a new Microsoft 365 admin center setting Copilot in M365 apps with Anthropic models in EU/EFTA and UK to enable Anthropic as the default model for Copilot in Microsoft 365 apps. For more information, see Copilot in Microsoft 365 apps with Anthropic models .
Opt-in to use Anthropic's models
If your organization is in a region that has Anthropic as a subprocessor set to Off by default, you can choose to opt in so Anthropic's models are available for your organization. You must be a member of the global administrator role to perform this task. For more information, see About admin roles and Global administrator .
Go to the Microsoft 365 admin center and select Copilot -> Settings -> View all .
Select AI providers operating as Microsoft subprocessors .
On the AI providers operating as Microsoft subprocessors page, under Available subprocessors for your organization , select Anthropic , and Save .
Under Choose who can access Anthropic models for Copilot and generative AI experiences , select your users or groups and Save .
You can restrict user access to AI provider subprocessors by assigning permissions to specific users or Microsoft Entra ID security groups in the Microsoft 365 admin center. These assignments are applied at the provider level and enforced across Microsoft 365 Copilot and Copilot Studio experiences. When access is limited by user or group membership, only the assigned users can use Copilot features or agents that rely on that AI provider. Review existing user or group assignments and update policies or configurations as needed. For more information on user and security group access, see Assign AI provider access to users and groups . For more information on creating security groups, see Create a security group .
If your organization is in the European Union (EU), the European Free Trade Association (EFTA), or the United Kingdom and you previously opted in to use Anthropic models under Anthropic's separate commercial terms and data processing agreement, you'll need to opt in again. The toggle will be set to Off by default.
Some features are only available when Anthropic models are enabled. If you turn off Anthropic as a subprocessor, certain features may no longer be accessible.
Opt-in to allow use of Preview models, including Preview models with Data Retention
Preview models are the latest models that are available for exploration and testing but aren't recommended for production use. Review the limitations of preview models before enabling use of Anthropic Preview models with Data Retention for your tenant.
Anthropic Preview models with Data Retention are advanced models (such as Claude Fable 5 and Claude Mythos 5 ) that require data retention by Anthropic. This means data is stored by Anthropic and not subject to your Microsoft Customer Agreement including commitments in the Product Terms and DPA. Anthropic Preview models with Data Retention are default-off for your tenant, even if Anthropic is enabled as a subprocessor. No users will have access to these models until the tenant admin explicitly opts in to enable use of Preview models with Data Retention. Use of Anthropic Preview models with Data Retention is subject to acceptance of Anthropic's Commercial Terms of Service and Anthropic's Data Protection Addendum . The tenant admin must also choose which users (all or specific users or groups) may access and use these models.
When users within your organization are granted access to Preview models with Data Retention, users will also be able to identify in the product experience that these models require data retention by Anthropic.
When you enable use of Anthropic Preview models with Data Retention, Anthropic retains data as explained in Anthropic's commercial data retention policy: Data retention for Mythos-class models This means that Anthropic (not Microsoft) stores most inputs and outputs for up to 30 days before deleting them. If Anthropic's trust and safety classifiers identify potential violations of Anthropic's Usage Policy, Anthropic may retain content (inputs and outputs) for up to 2 years and trust and safety classification scores for up to 7 years. Anthropic doesn't use retained data for model training without your express permission. For more information, see API and data retention - Claude API Docs .
Additional controls for Copilot Studio and Power Platform in the Power Platform Admin Center
Once enabled in the Microsoft 365 admin center, additional admin controls are available in the Microsoft Power Platform admin center (PPAC) to allow Anthropic to be used in Copilot Studio and Power Platform. For more information, see Allow external large language models (LLMs) for generative responses .
Disable connection to Anthropic's models
If your organization is in a region that has Anthropic as a subprocessor set to On by default, you can opt-out to restrict Anthropic models in the Microsoft 365 admin center. You must be a member of the global administrator role to perform this task. For more information, see About admin roles and Global administrator .
Go to the Microsoft 365 admin center and select Copilot -> Settings .
On the User access page, select AI providers operating as Microsoft subprocessors .
On the AI providers operating as Microsoft subprocessors page, under Available subprocessors for your organization , select Disable Anthropic as a Microsoft subprocessor .
Once you disable Anthropic as an AI subprocessor, users won't have the option to use Anthropic's AI models. You can choose to enable Anthropic models at a later date if desired.
Deprecation of legacy Anthropic admin toggle
The legacy Anthropic toggle to opt in to Anthropic's separate commercial terms and data processing agreement has been deprecated and replaced by this new Anthropic as a subprocessor admin toggle.
December 8, 2025: New administrator toggle for Anthropic models appears in the Microsoft 365 admin center. For most commercial cloud customers, it will be enabled On by default (except customers in EU/EFTA and UK will be Off by default). Some customer and region exclusions apply .
January 7, 2026: Anthropic as a Microsoft subprocessor becomes enabled in your organization. The legacy admin toggle to opt-in to Anthropic's commercial terms and data processing agreement is deprecated.
If you previously opted in to the legacy Anthropic toggle, and you're in a region where the new toggle is set to off by default, you need to opt in to the new subprocessor toggle to use Anthropic's models.
Understanding AI functionality and models in Microsoft Online Services
Data, Privacy, and Security for Microsoft 365 Copilot
Overview of AI subprocessors in Microsoft 365 Copilot
Supplier Security & Privacy Assurance
Want to try using Ask Learn to clarif

[truncated]
