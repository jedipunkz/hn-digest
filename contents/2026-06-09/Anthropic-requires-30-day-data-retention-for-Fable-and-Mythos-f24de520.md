---
source: "https://support.claude.com/en/articles/15425996-data-retention-practices-for-mythos-class-models"
hn_url: "https://news.ycombinator.com/item?id=48464258"
title: "Anthropic requires 30 day data retention for Fable and Mythos"
article_title: "Data retention practices for Mythos-class models | Claude Help Center"
author: "lebovic"
captured_at: "2026-06-09T18:51:07Z"
capture_tool: "hn-digest"
hn_id: 48464258
score: 5
comments: 0
posted_at: "2026-06-09T17:23:40Z"
tags:
  - hacker-news
  - translated
---

# Anthropic requires 30 day data retention for Fable and Mythos

- HN: [48464258](https://news.ycombinator.com/item?id=48464258)
- Source: [support.claude.com](https://support.claude.com/en/articles/15425996-data-retention-practices-for-mythos-class-models)
- Score: 5
- Comments: 0
- Posted: 2026-06-09T17:23:40Z

## Translation

タイトル: Anthropic では、寓話と神話について 30 日間のデータ保持が必要です
記事のタイトル: Mythos クラス モデルのデータ保持プラクティス |クロード ヘルプセンター

記事本文:
Mythos クラス モデルのデータ保持プラクティス |クロード ヘルプ センター メイン コンテンツにスキップ API ドキュメント リリース ノート サポートを受ける方法 English Français Deutsch Bahasa India Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English API ドキュメント リリース ノート サポートを受ける方法 English Français Deutsch Bahasa India Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English 記事の検索... すべてのコレクション
Mythos クラス モデルのデータ保持プラクティス
Mythos クラス モデルのデータ保持プラクティス
Mythos クラスのモデルを責任を持ってデプロイしていることを確認するために、安全作業の一環としてデータの保持とレビューを制限することが求められています。 Mythos クラス モデルに送信されたプロンプトと、Mythos クラス モデルによって生成された出力は、信頼性と安全性の目的で、これらのモデルが提供されているすべてのプラットフォームで 30 日間保持されます。
これは、Mythos クラスのモデルと、対象モデルとして指定されている同様の機能を備えた将来のモデルに適用されます。他のすべてのモデルについては、使用するものはすべて影響を受けず、現在の条件が適用されます。
以下に説明するこのポリシーは、2026 年 6 月 9 日に発効します。保持されるデータの脅威モデルおよび関連するプライバシー制御の詳細については、弊社のトラスト センターにある対応する技術ホワイト ペーパーを参照してください。
Claude.ai や Claude Code を含む、Web、デスクトップ、モバイル アプリ全体のコンシューマー プラン (Claude Free、Pro、Max) は、これらのサーフェス上で安全性を目的として入力と出力をすでに保持しているため、このアップデートの影響を受けません。消費者向けプランのデータをどのように保持するかについて詳しくご覧ください。
この変更は、Claude Console でデータ保持ゼロ (ZDR) のワークスペースを設定し、Cla で ZDR を備えた Claude Code を使用している組織にのみ適用されます。

ude Enterprise を使用するか、AWS Bedrock、Google Cloud Agent Platform、または ZDR を備えた Microsoft Foundry を介して Claude にアクセスします。この記事の残りの部分は、これらの組織にのみ適用されます。
Claude Mythos 5 では、モデルの機能が大幅に向上しており、その一部は無害な目的と悪意のある目的の両方に使用できます。 Claude Fable 5 は、Claude Mythos 5 と同じ基礎モデルを共有していますが、特にサイバー領域とバイオ領域において追加の安全対策が施されています。これらの保護手段により、このインテリジェンスをより広範囲に共有できるようになりますが、このクラスのモデルの誤用パターンを探すことができる保守的なアプローチを採用しています。一部の攻撃は、複数のリクエストにわたってのみ表示されます。たとえば、Best-of-N ジェイルブレイクでは、プロンプトが機能することを期待して、プロンプトのわずかなバリエーションを何百も送信します。国家支援のスパイ活動やデータ恐喝キャンペーンなど、より大規模な悪用パターンは、当社の保護分類器が多くのリクエストにわたってズームアウトできる場合にのみ表面化します。これらの脅威を検出するには、プロンプトと出力を一時的に保持して、一度に 1 つずつ分析するのではなく、まとめて分析できるようにする必要があります。
Anthropic の従業員は、深刻な危害の可能性があるとしてフラグが立てられるか、顧客の書面による要求がない限り、あなたの会話にアクセスすることはできません。これらのレビューは、エクスポート、コピー、ダウンロードを禁止するツールを使用して、承認された少数のレビュー担当者のみが実行できます。アクセスのすべてのインスタンスは、レビュー担当者が抑制または変更できない改ざん防止ログに記録されます。 30 日が経過すると、安全性調査の一環である場合や法的に保管が義務付けられている場合を除き、データは自動的に削除されます。対象となる組織には、顧客管理の暗号化キーを追加し、透明性監査ログにアクセスするオプションもあります。
人間は維持します

顧客データのセキュリティ、機密性、完全性を保護するために設計された技術的および組織的対策を備えた文書化された情報セキュリティ プログラム。当社のリスクベースのプログラムは、既知および予想される脅威モデルを防御するために構築および進化しており、定期的にテストされています。詳細については、トラスト センターのテクニカル ホワイト ペーパーを参照してください。
私たちのアプローチについてセキュリティ リーダーが言っているのは次のとおりです。
「Databricks を使用すると、企業は重要なデータを推論するエージェントを構築できます。Anthropic の新しいフロンティア モデルは、お客様がそれをより効果的に行うのに役立ちます。Anthropic は、新しいモデルの機能に合わせて安全策を継続的に進化させることで、私たち一人ひとりが責任を持ってインテリジェンスを拡張できるようにしています。私たちは、この新しい種類の安全策を業界に導入するという Anthropic の取り組みを強力にサポートしています。」
- Fermin Serna 氏、Databricks CISO
「金融エコシステムの信頼は強力なセキュリティに依存しており、Stripe は顧客に代わって防御を強化するために長年にわたって AI を使用してきました。フロンティア AI がセキュリティの脆弱性を発見する能力が向上するにつれて、それらの機能が防御的に使用されるようにすることがより重要になります。Anthropic はデータの処理と保持に透明性を提供することで、Mythos に強力な標準を設定しています。」
- Matthew Kemelhar 氏、Stripe CISO
何か設定が必要な場合、何を設定する必要がありますか?
この変更は、Claude Console でデータ保持ゼロ (ZDR) のワークスペースを設定している組織、Claude Enterprise で ZDR を使用して Claude Code を使用している組織、または ZDR を使用して AWS Bedrock、Google Cloud Agent Platform、または Microsoft Foundry を介して Claude にアクセスしている組織にのみ適用されます。他のすべての組織の場合、変更はなく、構成するものは何もありません。このセクションの残りの部分は、Claude にアクセスする組織を対象としています。

現在、データ保持は限界に達しており、指定されたモデルが利用可能になったときに使用するには、データ保持を設定する必要があります。
開発者が Claude API を使用している場合
Claude Platform を介して Anthropic から直接: 開発者コンソールで対象モデルを使用するワークスペースの保持をオンにします ( [ワークスペース] > [管理] > [プライバシー コントロール])。他の ZDR 対応ワークスペースでは ZDR が維持されます。ドキュメントについては、Anthropic Trust Center を参照してください。
AWS 上の Claude Platform 経由: 保持は直接の Claude API と同じように機能します。これはワークスペース レベルで構成され、保持されたデータは同じ制御の下で Anthropic によって処理されます。
Amazon Bedrock 経由: 新しい対象モデルにアクセスするには保持を有効にする必要があり、保持されたデータは AWS 環境に残ります。モデルが利用可能になると、オンボーディングの詳細が共有されます。
Google Cloud のエージェント プラットフォーム経由: 新しい対象モデルに対して保持を有効にする必要があり、保持されたデータは GCP 環境に残ります。モデルが利用可能になると、オンボーディングの詳細が共有されます。
Azure Foundry のクロード経由: 保持期間は Azure サブスクリプションごとに構成されます。ゼロ データ保持が構成されている場合、これらのモデルにアクセスするには、別の Azure サブスクリプションを作成して使用する必要があります。
Anthropic API を通じて: Claude Code のデータ処理慣行は、それが動作するワークスペースによって管理されます。そのワークスペースで保持が有効になっている場合、Claude Code は指定されたモデルを使用できます。直接サインインする開発者の場合は、組織のクロード コード ワークスペースでの保持を有効にします。
Amazon Bedrock または Google Cloud Agent Platform 経由: Claude Code はクラウド認証情報を使用するため、クラウド環境の保持設定に従います。クラウド環境で保持を有効にし、保持する必要があります。

データはプロバイダーの環境に残ります。同じことが、Amazon Bedrock または Google Cloud のエージェント プラットフォームを通じてアクセスされる Cowork にも当てはまります。
ZDR を使用した Claude Enterprise を通じて: プライマリ オーナーが保持設定を直接変更できるように、管理コンソールのコントロールをリリースします。まだ本番組織には触れたくない場合は、別の Sandbox 組織のセットアップをお手伝いします。
チームが Claude for Enterprise を介して Claude チャットまたは Cowork を使用している場合
これらのサーフェスはすでに標準の保持機能で動作しているため、新しいモデルが利用可能になったらすぐにアクセスできるようになります。
公共部門に関するよくある質問 サードパーティ プラットフォームで Claude for Microsoft 365 を使用する Claude 対象モデルでのリアルタイム サイバー セーフガード これはあなたの質問の答えになりましたか? 😞 😐 😃 目次 製品

## Original Extract

Data retention practices for Mythos-class models | Claude Help Center Skip to main content API Docs Release Notes How to Get Support English Français Deutsch Bahasa Indonesia Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English API Docs Release Notes How to Get Support English Français Deutsch Bahasa Indonesia Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English Search for articles... All Collections
Data retention practices for Mythos-class models
Data retention practices for Mythos-class models
To ensure we’re responsibly deploying Mythos-class models, we are requiring limited data retention and review as part of our safety work. Prompts submitted to, and outputs generated by, Mythos-class models are retained for 30 days for trust and safety purposes, on every platform where these models are offered.
This applies to Mythos-class models and future models with similar capabilities that we designate as covered models . For all other models, everything you use is unaffected and stays under the current terms.
This policy, described below, goes into effect on June 9, 2026. For more information on the threat model for retained data and associated privacy controls, please see the corresponding technical white paper on our Trust Center.
Consumer plans (Claude Free, Pro, and Max) across our web, desktop, and mobile apps—including Claude.ai and Claude Code—are unaffected by this update, since we already retain inputs and outputs for safety purposes on these surfaces. Learn more about how we retain data for consumer plans.
This change only applies to organizations that have set up workspaces with zero data retention (ZDR) in Claude Console, use Claude Code with ZDR in Claude Enterprise, or access Claude through AWS Bedrock, Google Cloud Agent Platform, or Microsoft Foundry with ZDR. The rest of this article applies only to these organizations.
Claude Mythos 5 represents a substantial increase in model capabilities, some of which can be used for both benign and malicious purposes. Claude Fable 5 shares the same underlying model as Claude Mythos 5, but with additional safeguards, particularly in the cyber and bio domains. While these safeguards allow us to share this intelligence more broadly, we are taking a conservative approach that allows us to look for patterns of misuse with this class of model. Some attacks only become visible across multiple requests. Best-of-N jailbreaking , for example, sends hundreds of slight variations of a prompt in the hope that one will work. Larger patterns of misuse, such as state-sponsored espionage or data extortion campaigns , only surface when our safeguards classifiers can zoom out across many requests. Detecting these threats requires temporarily retaining prompts and outputs so they can be analyzed together, rather than one at a time.
Anthropic employees cannot access your conversations unless they are flagged for potential serious harm or upon a customer’s written request. These reviews can only be performed by a small set of approved reviewers through tooling that prevents export, copying, or downloading. Every instance of access is recorded in a tamper-proof log that reviewers cannot suppress or modify. After 30 days, the data is deleted automatically, except in the rare cases where it's part of a safety investigation or we're legally required to keep it. Eligible organizations also have the option to add customer-managed encryption keys and access transparency audit logs.
Anthropic maintains a documented information security program with technical and organizational measures that are designed to protect the security, confidentiality, and integrity of customer data. Our risk-based program is built for and evolves to defend against known and anticipated threat models and is tested regularly. For more information, see the technical white paper in our Trust Center.
Here is what security leaders are saying about our approach:
“Databricks enables enterprises to build agents that reason over critical data, and new frontier models from Anthropic help our customers do that more effectively. By continuously evolving their safeguards to match new model capabilities, Anthropic is making it possible for each of us to scale intelligence responsibly. We strongly support Anthropic’s initiative to bring this new class of safeguards to the industry."
- Fermin Serna, Databricks CISO
“Trust in the financial ecosystem depends on strong security, and Stripe has used AI for years to help strengthen our defenses on behalf of our customers. As frontier AI becomes better at finding security vulnerabilities, it becomes more important to ensure those capabilities are used defensively. Anthropic is setting a strong standard for Mythos by providing transparency into data handling and retention.“
- Matthew Kemelhar, Stripe CISO
What, if anything, do I need to configure?
This change only applies to organizations that have set up workspaces with zero data retention (ZDR) in Claude Console, use Claude Code with ZDR in Claude Enterprise, or access Claude through AWS Bedrock, Google Cloud Agent Platform, or Microsoft Foundry with ZDR. For all other organizations, there is no change and there's nothing to configure. The rest of this section is for organizations that access Claude without data retention today and need to set up data retention in order to use designated models when they become available.
If your developers use the Claude API
Directly from Anthropic through Claude Platform: Turn on retention for the workspaces where you want to use covered models in the developer console ( Workspace > Manage > Privacy Controls ). Your other ZDR-enabled workspaces keep ZDR. Refer to the Anthropic Trust Center for documentation .
Through Claude Platform on AWS: Retention works the same way as the direct Claude API. It's configured at the workspace level, and retained data is handled by Anthropic under the same controls.
Through Amazon Bedrock: Retention will need to be enabled to access your new covered model, and retained data stays in your AWS environment. When models become available, onboarding details will be shared.
Through Google Cloud's Agent Platform: Retention will need to be enabled for your new covered model, and retained data stays in your GCP environment. When models become available, onboarding details will be shared.
Through Claude in Azure Foundry: Retention is configured for each Azure Subscription. If you have Zero Data Retention configured, then you will need to create and use a separate Azure Subscription to access these models.
Through the Anthropic API: Claude Code’s data handling practices are governed by the workspace it operates in. If that workspace has retention enabled, Claude Code can use designated models. For developers who sign in directly, enable retention at your organization’s Claude Code workspace.
Through Amazon Bedrock or Google Cloud Agent Platform: Claude Code uses your cloud credentials, so it follows your cloud environment's retention setting. Retention must be enabled in your cloud environment, and retained data stays in your provider's environment. The same applies to Cowork accessed through Amazon Bedrock or Google Cloud’s Agent Platform.
Through Claude Enterprise with ZDR: We're releasing controls in the admin console so your Primary Owner can change the retention setting directly. If you'd rather not touch your production org yet, we can help you set up a separate sandbox org.
If your team uses Claude chat or Cowork through Claude for Enterprise
These surfaces already operate with standard retention, so you'll have access to the new models as they become available.
Public Sector FAQs Use Claude for Microsoft 365 with third-party platforms Real-time cyber safeguards on Claude Covered Models Did this answer your question? 😞 😐 😃 Table of contents Product
