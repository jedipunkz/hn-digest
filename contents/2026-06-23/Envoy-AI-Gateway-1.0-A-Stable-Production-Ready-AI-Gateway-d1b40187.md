---
source: "https://aigateway.envoyproxy.io/blog/v1.0-release-announcement/"
hn_url: "https://news.ycombinator.com/item?id=48652415"
title: "Envoy AI Gateway 1.0 – A Stable, Production-Ready AI Gateway"
article_title: "Announcing Envoy AI Gateway 1.0 — A Stable, Production-Ready AI Gateway | Envoy AI Gateway"
author: "zoltano"
captured_at: "2026-06-23T23:23:44Z"
capture_tool: "hn-digest"
hn_id: 48652415
score: 3
comments: 0
posted_at: "2026-06-23T22:32:00Z"
tags:
  - hacker-news
  - translated
---

# Envoy AI Gateway 1.0 – A Stable, Production-Ready AI Gateway

- HN: [48652415](https://news.ycombinator.com/item?id=48652415)
- Source: [aigateway.envoyproxy.io](https://aigateway.envoyproxy.io/blog/v1.0-release-announcement/)
- Score: 3
- Comments: 0
- Posted: 2026-06-23T22:32:00Z

## Translation

タイトル: Envoy AI ゲートウェイ 1.0 – 安定した実稼働対応 AI ゲートウェイ
記事のタイトル: Envoy AI Gateway 1.0 の発表 — 安定した実稼働対応 AI ゲートウェイ | Envoy AI ゲートウェイ
説明: Envoy AI Gateway 1.0 は、CNCF Envoy Gateway 上に構築されたオープンソース AI ゲートウェイの最初の安定リリースです。コミットされた安定したコントロール プレーン API、16 の AI プロバイダー、MCP ゲートウェイ、マルチモーダル サポート、エンタープライズ オブザーバビリティはすべて、成長するコミュニティによって支えられています。

記事本文:
メイン コンテンツにスキップ Envoy AI ゲートウェイ ブログ ドキュメント 1.0 次へ
毎週の会議メモ (月曜日)
Envoy AI Gateway 1.0 を発表 — 安定した実稼働対応 AI ゲートウェイ
Envoy AI ゲートウェイ コントロール プレーンのスケーリングのベンチマーク
Envoy AI ゲートウェイを使用した MCP トラフィック ルーティングの現実とパフォーマンス
Envoy AI Gateway でのモデル コンテキスト プロトコルのサポートの発表
AI ゲートウェイの可観測性の強化 - OpenTelemetry トレースが Envoy AI ゲートウェイに到着
Envoy AI ゲートウェイ v0.3 リリースの発表
Envoy AI ゲートウェイがエンドポイント ピッカー サポートを導入
Envoy AI ゲートウェイの導入者向けのリファレンス アーキテクチャ
最初の Envoy AI ゲートウェイ リリースを発表 - コミュニティのマイルストーン!
KubeCon 2024 でのエンドユーザー基調講演
Envoy AI Gateway 1.0 を発表 — 安定した実稼働対応 AI ゲートウェイ
Dan Sun Envoy AI ゲートウェイ メンテナ - ブルームバーグ 米田武史 Envoy AI ゲートウェイ メンテナ - Netflix Aaron Choo Envoy AI ゲートウェイ メンテナ - ブルームバーグ Yao Weng Envoy AI ゲートウェイ メンテナ - ブルームバーグ Xunzhuo (Bit) Liu Envoy AI ゲートウェイ メンテナ - Tencent Ignasi Barrera 創設エンジニア - Tetrate Johnu George Envoy AI ゲートウェイ メンテナ - Nutanix Gavrish Prabhu Envoy AI ゲートウェイ メンテナー - Nutanix
本日、私たちは Envoy AI Gateway 1.0 (最初の安定版、一般提供版) を発表できることを嬉しく思います。
CNCF の Envoy Gateway 上に構築されたオープンソース AI ゲートウェイのリリース。
2025 年 2 月に v0.1 を出荷したとき、
私たちはその投稿を「1.0 に向かって！」という 3 つの言葉で締めくくりました。 16 か月後、何回もリリースされましたが、
業界全体のメンテナと採用者のコミュニティに支えられて、私たちはここにいます。 1.0 はできることを意味します
Envoy AI Gateway 上に自信を持って構築できます: 安定した実行を維持することに注力しているコントロール プレーン API
すでに世界最大規模の本番トラフィックを支えているのと同じ、歴戦の Envoy 基盤上で

企業。
🔒 1.0 の意味: 構築できる安定性
1.0 の見出しは単一の新機能ではなく、約束です。私たちの
リリースポリシーでは常にこう言っています
最初の安定したコントロール プレーン API が完成したら、メジャー v1.0.0 リリースを終了します。その瞬間がここにあります。
安定した API の場合、取り組みはシンプルかつ強力です。
重大なセキュリティ問題がない限り、API を破壊することはありません。また、常に API を提供します。
移行パスが必要な場合は、リリース ノートに記載してください。
これは企業が求めてきた基盤です。つまり、単一のコンポーネントを標準化する機能です。
移動するターゲットに賭けることなく、プロバイダーに依存しない AI ゲートウェイを実現します。
v0.1 は、2 つのプロバイダーと必須要素の前に統合された API を備えてリリースされました: アップストリーム
認可とトークンベースのレート制限。 1.0は別の動物です。プロジェクトがどこまで進んでいるかは次のとおりです
来てください：
OpenAI、Azure を含む、単一の OpenAI 互換インターフェイスを通じて統合された 16 のプロバイダー
OpenAI、Google Gemini、Google Vertex AI、AWS Bedrock、Anthropic、Mistral、Cohere、Groq、Togetter AI、
DeepInfra、DeepSeek、Hunyuan、SambaNova、Grok、および Tetrate Agent Router サービス。
アプリケーションを単一の OpenAI 互換エンドポイントに向けて、ゲートウェイに処理を任せます。
プロバイダー固有の変換、認証、ルーティング。タッチせずにプロバイダーを切り替えたり組み合わせたりできます
アプリケーション コード — モデルの仮想化を使用して、変更をルーティングする際にアプリ コードを安定させます。
その下:
バックエンド参照:
- 名前：openai - バックエンド
モデル名オーバーライド: "gpt-4o"
- 名前 : anthropic - バックエンド
モデル名オーバーライド: "claude-opus-4"
これは、A/B テスト、段階的な移行、マルチプロバイダー戦略、およびセキュリティ保護の鍵となります。
ベンダーロックイン。
プロバイダー認証、あなたのために処理されます
BackendSecurityPolicy はプロバイダーの資格情報をアプリケーションに漏らさないようにします

アップストリーム認証を一元化します —
API キー、AWS、Azure、GCP のクラウドネイティブ ID (Workload Identity を含む) はすべて、
ゲートウェイ。
エージェント時代の MCP ゲートウェイ
複数のモデル コンテキスト プロトコル サーバーを 1 つのエンドポイントの背後に集約し、ツールをルーティングおよびフィルター処理します。
クライアントは、きめ細かい CEL ベースの認証を確認して強制できるため、ツール/リストのみが常に有効になります
呼び出し元が実際に使用を許可されているものを返します。
エンタープライズ可観測性を組み込み
トークン対応メトリクス、OpenInference との互換性を備えた OpenTelemetry トレース (次のような評価ツール用)
Arize Phoenix) を使用し、推論トークンを個別に会計処理することで、コストの管理と可視性を実現します。
AI ワークロードが要求するもの。
Kubernetes Gateway API と
ゲートウェイ API 推論拡張機能、Envoy AI ゲートウェイ
Envoy Gateway の追加レイヤーです。Envoy が変更することなく GenAI トラフィックに対して実行できる機能を拡張します。
すでに導入して運用している方法。
🌍 CNCF Envoy 上でコミュニティによって構築
1.0 は、真の意味で業界を超えたコミュニティの作品です。メンテナは Tetrate、Bloomberg から来ています。
Tencent、Netflix、Nutanix に加え、週刊誌に参加する独立系寄稿者の名簿も増えています
コミュニティミーティング、ファイルの問題、および出荷コード。
同様に重要なのは、このプロジェクトが実際の使用をもとに形成されていることです。感謝の意を表します
途中でフィードバックを使用、テスト、共有した組織 (LY を含む)
Corporation、Alan by Comma Soft、および National Research Platform。
GitHub にアクセスし、リポジトリにスターを付けてサポートを示してください。
安定した API はゴールラインではなくスタートラインです。ロードマップ上:
専用の MCPBackend CRD 。MCP バックエンド構成を MCPRoute から切り離します。
ツール、リソース、プロンプトにわたるより深い MCP 認証。
自動的に設定される、より完全なクォータ認識ルーティング

レート制限のあるアップストリームの周りにあります。
プロバイダーの変換パスが増加し、マルチモーダル サポートが拡張されました。
このロードマップはコミュニティ主導型であり、その策定にご協力をお待ちしております。
1.0 は、私たちをここに導いたすべての人、つまりコードとレビューを書いたメンテナーと貢献者、
プレリリースをテストして問題点を教えてくれた早期採用者と、より広範なゲートウェイ API、
Envoy および CNCF コミュニティの標準を構築しています。ありがとう。
AI インフラストラクチャの未来は、オープンで安定しており、コミュニティ主導型です。あなたが何を期待するか楽しみです
その上に構築します。
Envoy AI ゲートウェイ 1.0 は現在利用可能です。詳細なリリース ノート、API の変更、アップグレード ガイダンスについては、
リリース ノート ページにアクセスしてください。
🔒 1.0 の意味: 構築できる安定性
✨ 1.0 の内容 1 つの API、すべてのプロバイダー
プロバイダー認証が自動的に処理されます
エージェント時代の MCP ゲートウェイ
エンタープライズ可観測性が組み込まれています
🌍 CNCF Envoy 上でコミュニティによって構築されました

## Original Extract

Envoy AI Gateway 1.0 is here — the first stable release of the open source AI gateway built on CNCF Envoy Gateway. A committed-stable control-plane API, 16 AI providers, an MCP gateway, multimodal support, and enterprise observability, all backed by a growing community.

Skip to main content Envoy AI Gateway Blog Docs 1.0 Next
Weekly Meeting Notes (Mondays)
Announcing Envoy AI Gateway 1.0 — A Stable, Production-Ready AI Gateway
Benchmarking Envoy AI Gateway Control Plane Scaling
The Reality and Performance of MCP Traffic Routing with Envoy AI Gateway
Announcing Model Context Protocol Support in Envoy AI Gateway
Enhancing AI Gateway Observability - OpenTelemetry Tracing Arrives in Envoy AI Gateway
Announcing the Envoy AI Gateway v0.3 Release
Envoy AI Gateway Introduces Endpoint Picker Support
A Reference Architecture for Adopters of Envoy AI Gateway
Announcing the first Envoy AI Gateway Release – A Community Milestone!
End User Keynote at KubeCon 2024
Announcing Envoy AI Gateway 1.0 — A Stable, Production-Ready AI Gateway
Dan Sun Envoy AI Gateway Maintainer - Bloomberg Takeshi Yoneda Envoy AI Gateway Maintainer - Netflix Aaron Choo Envoy AI Gateway Maintainer - Bloomberg Yao Weng Envoy AI Gateway Maintainer - Bloomberg Xunzhuo (Bit) Liu Envoy AI Gateway Maintainer - Tencent Ignasi Barrera Founding Engineer - Tetrate Johnu George Envoy AI Gateway Maintainer - Nutanix Gavrish Prabhu Envoy AI Gateway Maintainer - Nutanix
Today we're thrilled to announce Envoy AI Gateway 1.0 — the first stable, generally available
release of the open source AI gateway built on CNCF's Envoy Gateway .
When we shipped v0.1 in February 2025 ,
we closed that post with three words: "Onward to 1.0!" Sixteen months and many releases later,
backed by a community of maintainers and adopters across the industry, we're here. 1.0 means you can
build on Envoy AI Gateway with confidence: a control-plane API we're committing to keep stable, running
on the same battle-tested Envoy foundation that already powers production traffic at the world's largest
companies.
🔒 What 1.0 Means: Stability You Can Build On ​
The headline of 1.0 isn't a single new feature — it's a promise . Our
release policy has always said that
we'd cut the major v1.0.0 release once we had a first stable control-plane API . That moment is here.
For stable APIs, the commitment is simple and strong:
We will never break the APIs unless there is a critical security issue , and we will always provide
a migration path in the release notes if we ever must.
This is the foundation enterprises have been asking for — the ability to standardize on a single,
provider-agnostic AI gateway without betting on a moving target.
v0.1 launched with a unified API in front of two providers and the essentials: upstream
authorization and token-based rate limiting. 1.0 is a different animal. Here's how far the project has
come:
Sixteen providers, integrated through a single OpenAI-compatible interface, including OpenAI, Azure
OpenAI, Google Gemini, Google Vertex AI, AWS Bedrock, Anthropic, Mistral, Cohere, Groq, Together AI,
DeepInfra, DeepSeek, Hunyuan, SambaNova, Grok, and the Tetrate Agent Router Service.
Point your application at a single OpenAI-compatible endpoint and let the gateway handle
provider-specific translation, authentication, and routing. Switch or mix providers without touching
application code — and use model virtualization to keep your app code stable while routing changes
underneath:
backendRefs :
- name : openai - backend
modelNameOverride : "gpt-4o"
- name : anthropic - backend
modelNameOverride : "claude-opus-4"
This is the key to A/B testing, gradual migrations, multi-provider strategies, and safeguarding against
vendor lock-in.
Provider authentication, handled for you ​
BackendSecurityPolicy keeps provider credentials out of your application and centralizes upstream auth —
API keys, AWS, Azure, and GCP cloud-native identity (including Workload Identity), all managed at the
gateway.
An MCP gateway for the agent era ​
Aggregate multiple Model Context Protocol servers behind one endpoint, route and filter the tools
clients can see, and enforce fine-grained, CEL-based authorization — so tools/list only ever
returns what a caller is actually allowed to use.
Enterprise observability, built in ​
Token-aware metrics, OpenTelemetry tracing with OpenInference compatibility (for evaluation tools like
Arize Phoenix), and separate accounting for reasoning tokens give you the cost control and visibility
that AI workloads demand.
Built on the Kubernetes Gateway API and the
Gateway API Inference Extension , Envoy AI Gateway
is an additive layer on Envoy Gateway — it expands what Envoy can do for GenAI traffic without changing
how you already deploy and operate it.
🌍 Built by a Community, on CNCF Envoy ​
1.0 is the work of a genuinely cross-industry community. Maintainers come from Tetrate, Bloomberg,
Tencent, Netflix, and Nutanix , alongside a growing roster of independent contributors who join our weekly
community meetings , file issues, and ship code.
Just as importantly, the project has been shaped by real-world use. Our thanks to the
organizations who used, tested, and shared feedback along the way — including LY
Corporation, Alan by Comma Soft, and the National Research Platform .
Visit us on GitHub and star the repo to show your support.
A stable API is a starting line, not a finish line. On the roadmap:
A dedicated MCPBackend CRD , decoupling MCP backend configuration from MCPRoute .
Deeper MCP authorization across tools, resources, and prompts.
Fuller quota-aware routing that automatically steers around rate-limited upstreams.
More provider translation paths and expanded multimodal support.
The roadmap is community-driven — and we'd love your help shaping it.
1.0 belongs to everyone who got us here: the maintainers and contributors who wrote the code and reviews,
the early adopters who tested pre-releases and told us what broke, and the broader Gateway API,
Envoy, and CNCF communities whose standards we build on. Thank you.
The future of AI infrastructure is open, stable, and community-driven . We can't wait to see what you
build on it.
Envoy AI Gateway 1.0 is available now. For detailed release notes, API changes, and upgrade guidance,
visit our release notes page .
🔒 What 1.0 Means: Stability You Can Build On
✨ What's in 1.0 One API, every provider
Provider authentication, handled for you
An MCP gateway for the agent era
Enterprise observability, built in
🌍 Built by a Community, on CNCF Envoy
