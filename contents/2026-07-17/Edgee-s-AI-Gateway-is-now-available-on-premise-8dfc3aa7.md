---
source: "https://www.edgee.ai/blog/posts/edgee-on-premise-gateway"
hn_url: "https://news.ycombinator.com/item?id=48948141"
title: "Edgee's AI Gateway is now available on-premise"
article_title: "Edgee On-Premise: The Gateway, Behind Your Firewall - Edgee Blog"
author: "kokakiwi"
captured_at: "2026-07-17T15:06:11Z"
capture_tool: "hn-digest"
hn_id: 48948141
score: 2
comments: 0
posted_at: "2026-07-17T14:55:47Z"
tags:
  - hacker-news
  - translated
---

# Edgee's AI Gateway is now available on-premise

- HN: [48948141](https://news.ycombinator.com/item?id=48948141)
- Source: [www.edgee.ai](https://www.edgee.ai/blog/posts/edgee-on-premise-gateway)
- Score: 2
- Comments: 0
- Posted: 2026-07-17T14:55:47Z

## Translation

タイトル: Edgee の AI ゲートウェイがオンプレミスで利用可能になりました
記事のタイトル: Edgee オンプレミス: ファイアウォールの背後にあるゲートウェイ - Edgee ブログ
説明: 同じ圧縮、ルーティング、監視ゲートウェイを独自のインフラストラクチャ内に完全に導入できるようになりました。

記事本文:
Edgee オンプレミス: ファイアウォールの背後にあるゲートウェイ - Edgee ブログの機能 チーム ターボ モデル トークン圧縮 オンプレミス リーダーボード モデル ブログ 価格 ドキュメント ログイン メイン メニューを開く Edgee オンプレミス: ファイアウォールの背後にあるゲートウェイ
金融取引、臨床記録、政府事件ファイルなどの機密性の高いデータを処理する場合、データの保存場所を完全に制御することがセキュリティ要件の中核となります。これらの組織では、すべての推論は信頼できるインフラストラクチャ内で行われる必要があります。
現在、Edgee の AI ゲートウェイは完全なオンプレミス展開として利用可能です。当社のクラウド製品を強化するものと同じ高度な圧縮、ルーティング、監視機能を利用できますが、完全に独自の環境内で実行されます。これは、データ、プロバイダー キー、または可観測性の制御を手放すことなく、最新の AI モデルに接続し、プロバイダー間のトラフィックを管理できることを意味します。
セルフホスティングは、製品とのトレードオフではなく、導入の選択肢になりました。
単一の VM の Docker Compose セットアップ、または Kubernetes 1.24 以降のクラスターの Helm チャート。
1 つのコンテナー イメージ。読み取り専用のルート ファイル システムを使用して非ルートで実行されます。
ホスト型ゲートウェイと同じトークン圧縮、プロバイダー ルーティング、リクエスト レベルの可観測性。オンプレミス パスでは何も取り除かれません。
オンプレミスは、Edgee の共有クラウドおよび専用テナント ホスティング オプションと並ぶ、別個の層です。デプロイ シークレットは、組織管理者によって、コンソール → 組織設定 → オンプレミスを通じて組織ごとに発行されます。
ネットワークを離れずに一元化された構成
デフォルトの設定は接続モードです。ゲートウェイは、Edgee から組織の構成 (モデル レジストリ、ルーティング ルール、支出制限) を 15 秒ごとに定期的に同期し、使用状況を報告して、チーム ダッシュボードを常に稼働状態に保ちます。
プロンプトと

プロバイダー キーがインフラストラクチャから離れることはありません。唯一の送信トラフィックは次のとおりです。
LLM プロバイダーへ (プロンプト推論リクエストと応答用)
Edgee のコントロール プレーンへ (構成の同期と使用状況メトリクスやヘルス チェックなどのテレメトリ データ用)
プロンプト、入力完了、プロバイダー API キーが Edgee に送信されることはありません。管理者は、インフラストラクチャ上で実行されているゲートウェイを指すだけで、すでに使用しているのと同じコンソールを維持します。
完全にエアギャップされた環境では、Edgee への発信通話が許可されず、ヘッドレス モードではすべての同期が無効になります。構成はローカルで管理され、設定はインフラストラクチャ上のファイルに保存されます。使用状況データが Edgee に送信されることはなく、ゲートウェイはネットワーク内でのみ通信します。可観測性を確保するために、使用状況ログを独自の内部 OTLP エンドポイントにルーティングし、データを外部に共有することなく完全な可視性を維持できます。
コントロールは既存の場所に留まります
独自のプロバイダー API キー (BYOK): Edgee がそれらを保持することはありません。
独自のモデルでは、リストとチームごとまたはユーザーごとの支出上限をローカルで作成できます。
Edgee ダッシュボードではなくそこにメトリクスをルーティングしたい場合は、独自の可観測性バックエンド。
独自のシークレット管理: Helm チャートは、既存の Kubernetes シークレットからライセンス キーと署名キーを取得することをサポートしているため、値ファイルやバージョン管理には機密情報は何も含まれていません。
コンプライアンスとセキュリティの責任者は、「データはどこに行くのか」に対する明確な答えを得ることができます。つまり、すでに許可されていない場所には行くことはできません。プロンプトとプロバイダー キーは、管理するインフラストラクチャ上に残ります。
プラットフォームとインフラの所有者は、スタック内の他のものと同じようにインストールできるゲートウェイを取得します: helm upgrade --install または docker compose up 、バージョン管理されたリリース、予期せぬアップグレードはありません。
オンプレミスを評価している場合は、「Edgee オンプレミス」ページで詳細をご覧ください。
私たちを見つけたいですか

Edgee について詳しく知りたい場合は、サービスや今後の機能をテストしてください。ぜひご意見をお待ちしております。以下のフォームにご記入ください。ご連絡させていただきます。
ニュースレターを購読する 最新のニュース、記事、リソースが受信箱に送信されます。
X GitHub Discord © 2026 Edgee Cloud 無断複写・転載を禁じます。

## Original Extract

The same compress, route, and observe gateway, now deployable entirely inside your own infrastructure.

Edgee On-Premise: The Gateway, Behind Your Firewall - Edgee Blog Features Team Turbo Models Token Compression On-Premise Leaderboard Models Blog Pricing Documentation Login Open main menu Edgee On-Premise: The Gateway, Behind Your Firewall
When processing highly sensitive data, such as financial transactions, clinical records, or government case files, maintaining full control over data residency is a core security requirement. For these organisations, all inference must occur within a trusted infrastructure.
Now, Edgee's AI Gateway is available as a fully on-premise deployment. You get the same advanced compress, route, and observe capabilities that power our cloud offering, but running entirely inside your own environment. This means you can connect to the latest AI models and manage traffic across providers, without ever relinquishing control over your data, provider keys, or observability.
Self-hosting is now a deployment choice, not a trade-off against the product.
A Docker Compose setup for a single VM, or a Helm chart for any Kubernetes 1.24+ cluster.
One container image, running non-root with a read-only root filesystem.
The same token compression, provider routing, and request-level observability as the hosted gateway: nothing stripped down for the on-prem path.
On-Premise is a distinct tier alongside Edgee's Shared Cloud and Dedicated Tenant hosting options. Deployment secrets are issued per-org through Console → Org settings → On-Premise, by an org admin.
Centralised config, without leaving your network
The default setup is connected mode. The gateway regularly syncs your organisation's configuration (model registry, routing rules, spend limits) from Edgee every 15 seconds and reports usage back, keeping your team dashboard live.
Prompts and provider keys never leave your infrastructure. The only outbound traffic is:
To your LLM provider (for prompt inference requests and responses)
To Edgee's control plane (for configuration sync and telemetry data, such as usage metrics and health checks)
No prompts, completions, or provider API keys are ever sent to Edgee. Your admins keep the same console they already use, just pointed at a gateway running on your infrastructure.
In fully air-gapped environments, where no outbound calls to Edgee are allowed, headless mode disables all syncing. Configuration is managed locally, with settings stored in a file on your infrastructure. No usage data is ever sent to Edgee, and the gateway communicates only within your network. For observability, you can route usage logs to your own internal OTLP endpoint, maintaining full visibility without sharing any data externally.
Control stays where it already is
Your own provider API keys (BYOK): Edgee never holds them.
Your own model allows list and per-team or per-user spend caps, authored locally.
Your own observability backend, if you'd rather route metrics there instead of the Edgee dashboard.
Your own secrets management: the Helm chart supports pulling license and signature keys from an existing Kubernetes secret, so nothing sensitive sits in a values file or in version control.
Compliance and security leads get a straight answer to "where does our data go": nowhere it isn't already allowed to go. Prompts and provider keys stay on infrastructure you control.
Platform and infra owners get a gateway that installs like anything else in the stack: helm upgrade --install or docker compose up , versioned releases, no surprise upgrades.
If you're evaluating on-premises, learn more on the Edgee On-Premise page.
Would you like to find out more about Edgee, test our services or our upcoming features? We’d love to hear from you. Please fill in the form below and we’ll be in touch.
Subscribe to our newsletter The latest news, articles, and resources, sent to your inbox.
X GitHub Discord © 2026 Edgee Cloud All rights reserved.
