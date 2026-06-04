---
source: "https://developers.openai.com/api/docs/deprecations"
hn_url: "https://news.ycombinator.com/item?id=48391813"
title: "OpenAI Agent Builder Is Being Deprecated"
article_title: "Deprecations | OpenAI API"
author: "ILOVEZOA"
captured_at: "2026-06-04T01:06:10Z"
capture_tool: "hn-digest"
hn_id: 48391813
score: 1
comments: 0
posted_at: "2026-06-03T23:57:16Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Agent Builder Is Being Deprecated

- HN: [48391813](https://news.ycombinator.com/item?id=48391813)
- Source: [developers.openai.com](https://developers.openai.com/api/docs/deprecations)
- Score: 1
- Comments: 0
- Posted: 2026-06-03T23:57:16Z

## Translation

タイトル: OpenAI Agent Builder は非推奨になります
記事のタイトル: 非推奨 | OpenAI API
説明: OpenAI API の非推奨と推奨される代替品に関する情報を検索します。

記事本文:
非推奨 | OpenAI API
ホーム API ドキュメント OpenAI API のガイドと概念 API リファレンス エンドポイント、パラメーター、および応答 Codex ドキュメント Codex のガイド、概念、および製品ドキュメント 使用例 ワークフローとタスクの例 チームが Codex ChatGPT Apps SDK に渡す ChatGPT Commerce を拡張するアプリを構築する ChatGPT 広告でコマース フローを構築する ChatGPT 広告で広告を公開および測定する リソース ショーケース インスピレーションを得るためのデモ アプリ ブログ 開発者からの学習と経験 クックブックOpenAI モデルを使用して構築するためのノートブックの例 OpenAI を使用して構築するためのドキュメント、ビデオ、およびデモ アプリを学習する コミュニティ プログラム、交流会、ビルダーのサポート API ダッシュボードの検索を開始する API ドキュメントを検索する
統合と可観測性
MCP とコネクタによる安全な MCP トンネル
ファイルの検索と取得 ファイルの検索
文字起こし リアルタイム文字起こし
リアルタイムセッション 会話の管理
Webhook とサーバー側コントロール
Workload Identity フェデレーションの概要
最適化サイクルの微調整
直接的なプリファレンスの最適化
アシスタント API 移行ガイド
ブログ スキルを活用して OSS メンテナンスを加速する
Codex と Figma を使用したフロントエンド UI の構築
クックブック トレース、評価、コーデックスを使用してエージェント改善ループを構築する
Codex を使用して反復修復ループを構築する
Perplexity がリアルタイム API を使用して音声検索を何百万人ものユーザーに提供した方法
GPT-5.4 を使用した楽しいフロントエンドの設計
プロンプトから製品まで: 1 年間の対応
スキルを活用してOSSのメンテナンスを加速する
Codex と Figma を使用したフロントエンド UI の構築
統合と可観測性
MCP とコネクタによる安全な MCP トンネル
ファイルの検索と取得 ファイルの検索
文字起こし リアルタイム文字起こし
リアルタイムセッション 会話の管理
Webhook とサーバー側コントロール
Workload Identity フェデレーションの概要
最適化サイクルの微調整
直接優先

最適化以降
アシスタント API 移行ガイド
非推奨になった機能と推奨される代替機能を見つけます。
より安全でより機能的なモデルを発売するにつれて、古いモデルは定期的に廃止されます。 OpenAI モデルに依存するソフトウェアは、動作し続けるために時々更新する必要がある場合があります。影響を受けるお客様には、より大きな変更については常にメールとドキュメントでブログ投稿とともに通知されます。
このページには、すべての API の非推奨と、推奨される代替 API がリストされています。
モデルの非推奨通知期間
お客様が計画を立てて移行する時間を確保できるよう、モデルを廃止する前に事前に通知します。モデルの非推奨を発表する場合、モデルを積極的に使用している顧客に電子メールで通知し、このページに非推奨について文書化します。
安全性またはコンプライアンス上の懸念により、より迅速なスケジュールが必要な場合を除き、モデルの廃止前に次の最低通知期間を設けています。
一般提供モデル：最低6ヶ月。
一般的に入手可能なモデルの特殊バージョン: 少なくとも 3 か月。例には、 gpt-5.1-chat-latest などのチャット バリアント、 gpt-5.3-codex などの Codex バリアント、 o3-deep-research などのディープ リサーチ バリアントが含まれます。
プレビュー モデル: モデル名のプレビューで識別されるプレビュー モデルは、2 週間など、より短い通知で廃止される場合があります。例には、computer-use-preview および gpt-4o-audio-preview が含まれます。急遽移行できる場合を除き、ビジネス クリティカルな運用ワークロードにプレビュー モデルを使用することはお勧めしません。
安全性またはコンプライアンス上の理由により、モデルを早期に廃止する必要がある場合、当社は合理的に可能な限り多くの通知を提供します。
これらの通知期間により、お客様は、モデルが利用できなくなる前に、推奨される代替モデルを評価し、アプリケーションの動作をテストし、移行を完了する時間が与えられます。場合によっては、開発者は次のことを提供できる場合があります。

モデルのシャットダウン日以降も継続的にアクセスできる専用の容量。このオプションを検討するには、当社の営業チームにお問い合わせください。
モデルまたはエンドポイントを廃止するプロセスを指すために「非推奨」という用語を使用します。モデルまたはエンドポイントが非推奨になると発表すると、すぐに非推奨になります。非推奨のすべてのモデルとエンドポイントにもシャットダウン日が設定されます。シャットダウン時には、モデルまたはエンドポイントにアクセスできなくなります。
「サンセット」と「シャットダウン」という用語は、モデルまたはエンドポイントにアクセスできなくなったことを意味するために同じ意味で使用されます。
「レガシー」という用語は、アップデートを受け取らなくなったモデルとエンドポイントを指すために使用されます。私たちはエンドポイントとモデルをレガシーとしてタグ付けして、プラットフォームとして移行する方向と、新しいモデルまたはエンドポイントに移行する必要がある可能性があることを開発者に知らせます。レガシー モデルまたはエンドポイントは、将来のある時点で非推奨になることが予想されます。
今後の非推奨は以下にリストされており、最新の発表が一番上にあります。
2026 年 6 月 3 日、ダッシュボードと API で再利用可能なプロンプトを使用している開発者に、再利用可能なプロンプト オブジェクトが非推奨になることを通知しました。
移行するには、再利用可能なプロンプト コンテンツをアプリケーション コードに移動します。 「プロンプト オブジェクトからの移行」を参照してください。
2026 年 6 月 3 日、Evals プラットフォームを使用している開発者に、この製品が非推奨になることを通知しました。
評価ワークフローに関して文書化された採点者は、この移行の一部です。微調整関連のタイムラインについては、以下のセルフサービス微調整セクションで引き続き説明します。
移行パスについては、「OpenAI Evals から Promptfoo への移行」を参照してください。
2026 年 6 月 3 日、Agent Builder を使用している開発者に、この製品が非推奨になることを通知しました。 ChatKit は引き続き利用可能です。
Agent SDK または ChatGPT Workspace A を続行するには、「Agent Builder からの移行」を参照してください。

男性諸君。
2026-06-02: GPT イメージ モデルの非推奨
2026 年 6 月 2 日、古い GPT イメージ モデルを使用している開発者に、2026 年 12 月 1 日に廃止され、API から削除されることを通知しました。
2026-05-08: gpt-5.2-chat-latest および gpt-5.3-chat-latest モデルのスナップショット
2026 年 5 月 8 日に、gpt-5.2-chat-latest および gpt-5.3-chat-latest モデルのスナップショットを使用している開発者に、API からの廃止と削除について通知しました。
OpenAI のセルフサービス微調整の更新
2026 年 5 月 7 日に、OpenAI のセルフサービス微調整プラットフォームを使用している開発者に可用性の更新を通知しました。
微調整モデルの推論は、基本モデルが非推奨になるまで引き続き利用できます。
2026-04-22: レガシー GPT モデルのスナップショット
信頼性を向上させ、開発者が適切なモデルを選択しやすくするために、一連の古い OpenAI モデルを非推奨にします。これらのモデルへのアクセスは、以下の日付で停止されます。
また、以下のような微調整されたバージョンも削除されます。
2026-03-24: Sora 2 ビデオ生成モデルとビデオ API
2026 年 3 月 24 日に、Videos API と Sora 2 ビデオ生成モデルのエイリアスとスナップショットを使用している開発者に、非推奨となり、2026 年 9 月 24 日に API から削除されることを通知しました。
2025-11-14: DALL・E モデルのスナップショット
2025 年 11 月 14 日に、DALL·E モデルのスナップショットを使用している開発者に、2026 年 5 月 12 日に非推奨となり API から削除されることを通知しました。
2025-09-26: レガシー GPT モデルのスナップショット
信頼性を向上させ、開発者が適切なモデルを選択しやすくするために、使用量が減少している一連の古い OpenAI モデルを今後 6 ～ 12 か月間で非推奨とします。これらのモデルへのアクセスは、以下の日付で停止されます。
*特にレイテンシーに敏感で、推論を必要としないタスクの場合
Realtime API ベータ版は非推奨となり、削除されました。

2026 年 5 月 12 日の API。
リアルタイム ベータ API とリリースされた GA API のインターフェイスには、いくつかの重要な違いがあります。現在の GA インターフェイスの移行ガイドおよび関連するリアルタイム ドキュメントを参照してください。
2025 年 8 月 26 日に、アシスタント API を使用している開発者に、アシスタント API が非推奨となり、1 年後の 2026 年 8 月 26 日に API から削除されることを通知しました。
2025 年 3 月に Responses API をリリースしたとき、すべてのアシスタント API 機能を使いやすい Responses API に移行し、2026 年にサービスを終了する計画を発表しました。
現在の統合を Responses API および Conversations API に移行する方法の詳細については、アシスタントから会話への移行ガイドを参照してください。
2025-09-15: gpt-4o-realtime-preview モデル
2025 年 9 月に、gpt-4o-realtime-preview モデルを使用している開発者に、モデルが非推奨となり、6 か月以内に API から削除されることを通知しました。
過去の非推奨は以下にリストされており、最新の発表が一番上にあります。
2025-11-18: chatgpt-4o-最新のスナップショット
2025 年 11 月 18 日に、chatgpt-4o-最新モデルのスナップショットを使用している開発者に、非推奨となり、2026 年 2 月 17 日に API から削除されることを通知しました。
2025-11-17: codex-mini-最新モデルのスナップショット
2025 年 11 月 17 日に、codex-mini-latest モデルを使用している開発者に、このモデルが非推奨となり、2026 年 2 月 12 日に API から削除されることを通知しました。この非推奨の一環として、 codex-mini-latest でのみ使用できる従来のローカル シェル ツールはサポートされなくなります。新しいユースケースの場合は、最新のシェル ツールを使用してください。
2025-06-10: gpt-4o-realtime-preview-2024-10-01
2025 年 6 月 10 日に、gpt-4o-realtime-preview-2024-10-01 を使用している開発者に、その廃止と 3 か月以内の API からの削除を通知しました。
2025-06-10: gpt-4o-audio-preview-2024-10-01
2025 年 6 月 10 日に、当社は de

gpt-4o-audio-preview-2024-10-01 を使用しているベロッパーは、3 か月以内に非推奨となり API から削除されます。
2025 年 4 月 28 日に、テキスト モデレーションを使用している開発者に、6 か月以内に非推奨となり API から削除されることを通知しました。
2025-04-28: o1-preview と o1-mini
2025 年 4 月 28 日に、o1-preview と o1-mini を使用している開発者に、それぞれ 3 か月後と 6 か月後に非推奨となり API から削除されることを通知しました。
2025 年 4 月 14 日に、gpt-4.5-preview モデルが非推奨となり、今後数か月以内に API から削除されることを開発者に通知しました。
2024-10-02: アシスタント API ベータ版 v1
2024 年 4 月にアシスタント API の v2 ベータ版をリリースした際、v1 ベータ版へのアクセスは 2024 年末までに停止されると発表しました。v1 ベータ版へのアクセスは 2024 年 12 月 18 日に停止されます。
ツールの使用状況をアシスタント API の最新バージョンに移行する方法の詳細については、アシスタント API v2 ベータ移行ガイドを参照してください。
2024-08-29: babbage-002 および davinci-002 モデルのトレーニングの微調整
2024 年 8 月 29 日に、babbage-002 と davinci-002 を微調整する開発者に、これらのモデルでの新しい微調整トレーニングの実行は 2024 年 10 月 28 日以降サポートされなくなることを通知しました。
これらの基本モデルから作成された微調整モデルはこの非推奨の影響を受けませんが、これらのモデルを使用して新しい微調整バージョンを作成することはできなくなります。
2024-06-06: GPT-4-32K および Vision Preview モデル
2024 年 6 月 6 日に、gpt-4-32k と gpt-4-vision-preview を使用している開発者に、それぞれ 1 年後と 6 か月以内に非推奨になることを通知しました。 2024 年 6 月 17 日以降、これらのモデルの既存ユーザーのみが引き続き使用できるようになります。
2023-11-06: チャット モデルの更新
2023 年 11 月 6 日に、更新された GPT-3 のリリースを発表しました。

.5-Turbo モデル (デフォルトで 16k コンテキストが付属) と、 gpt-3.5-turbo-0613 および gpt-3.5-turbo-16k-0613 の非推奨。 2024 年 6 月 17 日以降、これらのモデルの既存ユーザーのみが引き続き使用できるようになります。
これらの基本モデルから作成された微調整モデルはこの非推奨の影響を受けませんが、これらのモデルを使用して新しい微調整バージョンを作成することはできなくなります。
2023-08-22: エンドポイントを微調整
2023 年 8 月 22 日に、新しい微調整 API ( /v1/fine_tuning/jobs ) と、元の /v1/fine-tunes API および従来のモデル (/v1/fine-tunes API で微調整されたモデルを含む) が 2024 年 1 月 4 日にシャットダウンされることを発表しました。これは、/v1/fine-tunes API を使用して微調整されたモデルにアクセスできなくなることを意味します。更新されたエンドポイントと関連する基本モデルを使用して、新しいモデルを微調整します。
シャットダウン日 システム 推奨代替品 2024-01-04 /v1/fine-tunes /v1/fine_tuning/jobs
2023-07-06: GPT とエンベディング
2023 年 7 月 6 日、完了エンドポイント経由で提供される古い GPT-3 および GPT-3.5 モデルの今後の廃止を発表しました。また、第 1 世代のテキスト埋め込みモデルが近日中に廃止されることも発表しました。 1月4日をもって閉店となります。

[切り捨てられた]

## Original Extract

Find information about OpenAI API deprecations and recommended replacements.

Deprecations | OpenAI API
Home API Docs Guides and concepts for the OpenAI API API reference Endpoints, parameters, and responses Codex Docs Guides, concepts, and product docs for Codex Use cases Example workflows and tasks teams hand to Codex ChatGPT Apps SDK Build apps to extend ChatGPT Commerce Build commerce flows in ChatGPT Ads Publish and measure ads in ChatGPT Resources Showcase Demo apps to get inspired Blog Learnings and experiences from developers Cookbook Notebook examples for building with OpenAI models Learn Docs, videos, and demo apps for building with OpenAI Community Programs, meetups, and support for builders Start searching API Dashboard Search the API docs
Integrations and observability
MCP and Connectors Secure MCP Tunnel
File search and retrieval File search
Transcription Realtime transcription
Realtime sessions Managing conversations
Webhooks and server-side controls
Workload identity federation Overview
Fine-tuning Optimization cycle
Direct preference optimization
Assistants API Migration guide
Blog Using skills to accelerate OSS maintenance
Building frontend UIs with Codex and Figma
Cookbooks Build an Agent Improvement Loop with Traces, Evals, and Codex
Build iterative repair loops with Codex
How Perplexity Brought Voice Search to Millions Using the Realtime API
Designing delightful frontends with GPT-5.4
From prompts to products: One year of Responses
Using skills to accelerate OSS maintenance
Building frontend UIs with Codex and Figma
Integrations and observability
MCP and Connectors Secure MCP Tunnel
File search and retrieval File search
Transcription Realtime transcription
Realtime sessions Managing conversations
Webhooks and server-side controls
Workload identity federation Overview
Fine-tuning Optimization cycle
Direct preference optimization
Assistants API Migration guide
Find deprecated features and recommended replacements.
As we launch safer and more capable models, we regularly retire older models. Software relying on OpenAI models may need occasional updates to keep working. Impacted customers will always be notified by email and in our documentation along with blog posts for larger changes.
This page lists all API deprecations, along with recommended replacements.
Model deprecation notice periods
We provide advance notice before retiring models so customers have time to plan and migrate. When we announce a model deprecation, we notify customers who are actively using the model by email and document the deprecation on this page.
Unless safety or compliance concerns require a faster timeline, we provide the following minimum notice periods before model retirement:
Generally available models: At least 6 months.
Specialized variants of generally available models: At least 3 months. Examples include chat variants such as gpt-5.1-chat-latest , Codex variants such as gpt-5.3-codex , and deep research variants such as o3-deep-research .
Preview models: Preview models, identified by preview in the model name, may be retired with much shorter notice, such as 2 weeks. Examples include computer-use-preview and gpt-4o-audio-preview . We don’t recommend using preview models for business-critical production workloads unless you can migrate on short notice.
If safety or compliance concerns require us to retire a model sooner, we will provide as much notice as reasonably possible.
These notice periods give customers time to evaluate recommended replacement models, test application behavior, and complete migrations before a model is no longer available. In some cases, developers may be able to provision dedicated capacity for continued access after a model’s shutdown date. To explore this option, contact our sales team .
We use the term “deprecation” to refer to the process of retiring a model or endpoint. When we announce that a model or endpoint is being deprecated, it immediately becomes deprecated. All deprecated models and endpoints will also have a shut down date. At the time of the shut down, the model or endpoint will no longer be accessible.
We use the terms “sunset” and “shut down” interchangeably to mean a model or endpoint is no longer accessible.
We use the term “legacy” to refer to models and endpoints that no longer receive updates. We tag endpoints and models as legacy to signal to developers where we’re moving as a platform and that they should likely migrate to newer models or endpoints. You can expect that a legacy model or endpoint will be deprecated at some point in the future.
Upcoming deprecations are listed below, with the most recent announcements at the top.
On June 3, 2026, we notified developers using reusable prompts in the dashboard and API that reusable prompt objects are being deprecated.
To migrate, move reusable prompt content into your application code. See Migrate from prompt objects .
On June 3, 2026, we notified developers using the Evals platform that the product is being deprecated.
Graders documented for eval workflows are part of this transition. Fine-tuning-related timelines remain covered in the self-serve fine-tuning section below.
See Moving from OpenAI Evals to Promptfoo for a migration path.
On June 3, 2026, we notified developers using Agent Builder that the product is being deprecated. ChatKit remains available.
See Migrate from Agent Builder to continue with the Agents SDK or ChatGPT Workspace Agents.
2026-06-02: GPT Image model deprecations
On June 2, 2026, we notified developers using older GPT Image models of their deprecation and removal from the API on December 1, 2026.
2026-05-08: gpt-5.2-chat-latest and gpt-5.3-chat-latest model snapshots
On May 8th, 2026, we notified developers using gpt-5.2-chat-latest and gpt-5.3-chat-latest model snapshots of their deprecation and removal from the API.
Update to OpenAI’s self-serve fine-tuning
On May 7th, 2026, we notified developers using OpenAI’s self-serve fine-tuning platform of updates to availability.
Inference on fine-tuned models will continue to be available until the base models are deprecated.
2026-04-22: Legacy GPT model snapshots
To improve reliability and make it easier for developers to choose the right models, we are deprecating a set of older OpenAI models. Access to these models will be shut down on the dates below.
We are also removing fine-tuned versions as below:
2026-03-24: Sora 2 video generation models and Videos API
On March 24th, 2026, we notified developers using the Videos API and Sora 2 video generation model aliases and snapshots of their deprecation and removal from the API on September 24, 2026.
2025-11-14: DALL·E model snapshots
On November 14th, 2025, we notified developers using DALL·E model snapshots of their deprecation and removal from the API on May 12, 2026.
2025-09-26: Legacy GPT model snapshots
To improve reliability and make it easier for developers to choose the right models, we are deprecating a set of older OpenAI models with declining usage over the next six to twelve months. Access to these models will be shut down on the dates below.
*For tasks that are especially latency sensitive and don’t require reasoning
The Realtime API Beta was deprecated and removed from the API on May 12, 2026.
There are a few key differences between the interfaces in the Realtime beta API and the released GA API. See the migration guide for the current GA interface and related Realtime docs.
On August 26th, 2025, we notified developers using the Assistants API of its deprecation and removal from the API one year later, on August 26, 2026.
When we released the Responses API in March 2025 , we announced plans to bring all Assistants API features to the easier to use Responses API, with a sunset date in 2026.
See the Assistants to Conversations migration guide to learn more about how to migrate your current integration to the Responses API and Conversations API.
2025-09-15: gpt-4o-realtime-preview models
In September, 2025, we notified developers using gpt-4o-realtime-preview models of their deprecation and removal from the API in six months.
Past deprecations are listed below, with the most recent announcements at the top.
2025-11-18: chatgpt-4o-latest snapshot
On November 18th, 2025, we notified developers using chatgpt-4o-latest model snapshot of its deprecation and removal from the API on February 17, 2026.
2025-11-17: codex-mini-latest model snapshot
On November 17th, 2025, we notified developers using codex-mini-latest model of its deprecation and removal from the API on February 12, 2026. As part of this deprecation, we will no longer support our legacy local shell tool, which is only available for use with codex-mini-latest . For new use cases, please use our latest shell tool.
2025-06-10: gpt-4o-realtime-preview-2024-10-01
On June 10th, 2025, we notified developers using gpt-4o-realtime-preview-2024-10-01 of its deprecation and removal from the API in three months.
2025-06-10: gpt-4o-audio-preview-2024-10-01
On June 10th, 2025, we notified developers using gpt-4o-audio-preview-2024-10-01 of its deprecation and removal from the API in three months.
On April 28th, 2025, we notified developers using text-moderation of its deprecation and removal from the API in six months.
2025-04-28: o1-preview and o1-mini
On April 28th, 2025, we notified developers using o1-preview and o1-mini of their deprecations and removal from the API in three months and six months respectively.
On April 14th, 2025, we notified developers that the gpt-4.5-preview model is deprecated and will be removed from the API in the coming months.
2024-10-02: Assistants API beta v1
In April 2024 when we released the v2 beta version of the Assistants API, we announced that access to the v1 beta would be shut off by the end of 2024. Access to the v1 beta will be discontinued on December 18, 2024.
See the Assistants API v2 beta migration guide to learn more about how to migrate your tool usage to the latest version of the Assistants API.
2024-08-29: Fine-tuning training on babbage-002 and davinci-002 models
On August 29th, 2024, we notified developers fine-tuning babbage-002 and davinci-002 that new fine-tuning training runs on these models will no longer be supported starting October 28, 2024.
Fine-tuned models created from these base models are not affected by this deprecation, but you will no longer be able to create new fine-tuned versions with these models.
2024-06-06: GPT-4-32K and Vision Preview models
On June 6th, 2024, we notified developers using gpt-4-32k and gpt-4-vision-preview of their upcoming deprecations in one year and six months respectively. As of June 17, 2024, only existing users of these models will be able to continue using them.
2023-11-06: Chat model updates
On November 6th, 2023, we announced the release of an updated GPT-3.5-Turbo model (which now comes by default with 16k context) along with deprecation of gpt-3.5-turbo-0613 and gpt-3.5-turbo-16k-0613 . As of June 17, 2024, only existing users of these models will be able to continue using them.
Fine-tuned models created from these base models are not affected by this deprecation, but you will no longer be able to create new fine-tuned versions with these models.
2023-08-22: Fine-tunes endpoint
On August 22nd, 2023, we announced the new fine-tuning API ( /v1/fine_tuning/jobs ) and that the original /v1/fine-tunes API along with legacy models (including those fine-tuned with the /v1/fine-tunes API) will be shut down on January 04, 2024. This means that models fine-tuned using the /v1/fine-tunes API will no longer be accessible and you would have to fine-tune new models with the updated endpoint and associated base models.
Shutdown date System Recommended replacement 2024-01-04 /v1/fine-tunes /v1/fine_tuning/jobs
2023-07-06: GPT and embeddings
On July 06, 2023, we announced the upcoming retirements of older GPT-3 and GPT-3.5 models served via the completions endpoint. We also announced the upcoming retirement of our first-generation text embedding models. They will be shut down on January 04,

[truncated]
