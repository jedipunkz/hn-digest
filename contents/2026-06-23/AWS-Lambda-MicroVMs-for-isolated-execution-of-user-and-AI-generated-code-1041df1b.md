---
source: "https://aws.amazon.com/about-aws/whats-new/2026/06/aws-lambda-microvms/"
hn_url: "https://news.ycombinator.com/item?id=48639498"
title: "AWS Lambda MicroVMs for isolated execution of user and AI-generated code"
article_title: "AWS introduces Lambda MicroVMs for isolated execution of user and AI-generated code - AWS"
author: "leemoore"
captured_at: "2026-06-23T03:03:20Z"
capture_tool: "hn-digest"
hn_id: 48639498
score: 9
comments: 1
posted_at: "2026-06-23T02:38:31Z"
tags:
  - hacker-news
  - translated
---

# AWS Lambda MicroVMs for isolated execution of user and AI-generated code

- HN: [48639498](https://news.ycombinator.com/item?id=48639498)
- Source: [aws.amazon.com](https://aws.amazon.com/about-aws/whats-new/2026/06/aws-lambda-microvms/)
- Score: 9
- Comments: 1
- Posted: 2026-06-23T02:38:31Z

## Translation

タイトル: ユーザーコードと AI 生成コードを分離して実行するための AWS Lambda MicroVM
記事のタイトル: AWS、ユーザーコードと AI 生成コードの分離実行のための Lambda MicroVM を導入 - AWS
説明: 内容について詳しく知る

記事本文:
AWS がユーザーコードと AI 生成コードを分離して実行するための Lambda MicroVM を導入 - AWS
メインコンテンツにスキップ
AWS がユーザーコードと AI 生成コードを分離して実行するための Lambda MicroVM を導入
AWS は、VM レベルの分離、ほぼ瞬時の起動と再開の速度、およびユーザーまたは AI が生成したコードを実行する状態の保存を提供する新しいサーバーレス コンピューティング プリミティブである Lambda MicroVM を導入します。各ユーザーまたはジョブに独自のコンピューティング環境を提供し、仮想化インフラストラクチャを管理したり、分離、速度、状態保持のいずれかを選択したりすることなく、コードを安全に実行できるようになりました。
開発者は、インタラクティブ コーディング環境、データ分析プラットフォーム、コーディング アシスタント、脆弱性スキャン プラットフォームなどのユースケースのために、エンド ユーザーまたは AI によって提供されたコードを実行するマルチテナント アプリケーションを構築することが増えています。これらのアプリケーションの場合、開発者は、同時に実行されている他のユーザーやジョブに対する不正または悪意のあるコードの影響を制限するために、ユーザーまたはセッションごとに個別の分離された実行環境を割り当てる必要があります。以前は、開発者はこれらのアプリケーションを構築する際に、強力な分離、高速な起動時間、状態保持のいずれかを選択する必要がありました。本日より、Lambda MicroVM はこれらの機能をトレードオフなしで提供します。 VM レベルの分離、ほぼ瞬時の起動速度、および最大 8 時間の実行の一時停止と再開の機能が得られます。 Lambda MicroVM は、月間 15 兆回を超える Lambda 関数呼び出しを実現するテクノロジーである Firecracker 仮想化に基づいて構築されています。
まず、Dockerfile から MicroVM イメージを作成し、そのイメージから MicroVM を起動します。各ユーザーまたはジョブに、HTTP/2、gRPC、WebSocket などの一般的な接続プロトコルをサポートする専用の HTTPS URL を使用して独自の MicroVM を与えます。

Lambda MicroVM は現在、米国東部 (バージニア北部)、米国東部 (オハイオ)、米国西部 (オレゴン)、アジアパシフィック (東京)、欧州 (アイルランド) の AWS リージョンでご利用いただけます。詳細については、AWS Lambda MicroVM 開発者ガイドとリリースに関するブログ投稿をご覧ください。 AWS Lambda コンソール、AWS CloudFormation、AWS クラウド開発キットを通じて MicroVM の使用を開始するか、好みのエージェント開発ツールで AWS 用エージェント ツールキットを使用します。 MicroVM の実行中にベースラインのコンピューティング リソースに対して料金が発生します。また、ワークロードがベースラインを超えた場合に消費される追加リソースのアクティブ期間に対してのみ料金が発生します。料金の詳細については、「Lambda MicroVM の料金」を参照してください。
英語
トップに戻る
Amazon は機会均等な雇用主であり、保護されている退役軍人ステータス、障害、またはその他の法的に保護されているステータスに基づいて差別することはありません。
×
フェイスブック
リンクイン
インスタグラム
けいれん
ユーチューブ
ポッドキャスト
電子メール
プライバシー
プライバシーに関する選択
カリフォルニア州消費者プライバシー法 (CCPA) オプトアウト アイコン

## Original Extract

Discover more about what

AWS introduces Lambda MicroVMs for isolated execution of user and AI-generated code - AWS
Skip to main content
AWS introduces Lambda MicroVMs for isolated execution of user and AI-generated code
AWS introduces Lambda MicroVMs, a new serverless compute primitive that provides VM-level isolation, near-instant launch and resume speeds, and state preservation for executing user or AI-generated code. You can now give each user or job their own compute environment to securely run code without managing virtualization infrastructure or choosing between isolation, speed, and state retention.
Developers are increasingly building multi-tenant applications that execute code supplied by end users or AI for use cases such as interactive coding environments, data analytics platforms, coding assistants, and vulnerability scanning platforms. For these applications, developers need to allocate a separate, isolated execution environment per user or session to limit the impact of incorrect or malicious code on other concurrently running users or jobs. Previously, developers needed to choose between strong isolation, fast launch times, and state retention when building these applications. Starting today, Lambda MicroVMs provides you these capabilities without any trade-offs. You get VM-level isolation, near-instant launch speeds, and the ability to suspend and resume execution for up to 8 hours. Lambda MicroVMs is built on Firecracker virtualization, the technology powering more than 15 trillion monthly Lambda Function invocations.
To get started, create a MicroVM image from your Dockerfile, then launch MicroVMs from that image. Give each user or job their own MicroVM with a dedicated HTTPS URL that supports popular connectivity protocols such as HTTP/2, gRPC, and WebSockets.
Lambda MicroVMs is available today in the following AWS Regions: US East (N. Virginia), US East (Ohio), US West (Oregon), Asia Pacific (Tokyo), and Europe (Ireland). To learn more, visit the AWS Lambda MicroVMs developer guide and the launch blog post . Get started with MicroVMs through the AWS Lambda console, AWS CloudFormation, AWS Cloud Development Kit, or use the Agent Toolkit for AWS with your preferred Agentic development tools. You pay for baseline compute resources while your MicroVM is running, and only for the active duration of additional resources consumed when your workload exceeds the baseline. To learn more about pricing, see Lambda MicroVMs pricing .
English
Back to top
Amazon is an equal opportunity employer and does not discriminate on the basis of protected veteran status, disability or other legally protected status.
x
facebook
linkedin
instagram
twitch
youtube
podcasts
email
Privacy
Your Privacy Choices
California Consumer Privacy Act (CCPA) Opt-Out Icon
