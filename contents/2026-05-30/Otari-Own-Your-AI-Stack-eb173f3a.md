---
source: "https://blog.mozilla.ai/otari-own-your-ai-stack/"
hn_url: "https://news.ycombinator.com/item?id=48328461"
title: "Otari: Own Your AI Stack"
article_title: "Otari: Own Your AI Stack | AI Gateway & Hosted Platform"
author: "mwheeler"
captured_at: "2026-05-30T11:39:09Z"
capture_tool: "hn-digest"
hn_id: 48328461
score: 11
comments: 1
posted_at: "2026-05-29T20:03:15Z"
tags:
  - hacker-news
  - translated
---

# Otari: Own Your AI Stack

- HN: [48328461](https://news.ycombinator.com/item?id=48328461)
- Source: [blog.mozilla.ai](https://blog.mozilla.ai/otari-own-your-ai-stack/)
- Score: 11
- Comments: 1
- Posted: 2026-05-29T20:03:15Z

## Translation

タイトル: オタリ: 自分の AI スタックを所有する
記事のタイトル: Ohari: 自分の AI スタックを所有する | AI ゲートウェイとホスト型プラットフォーム
説明: Ohari は、any-llm 上に構築されたオープンソース LLM ゲートウェイです。 Ohari.ai は、ホスト型ルーティング、予算、可観測性、およびモデル管理を追加します。

記事本文:
サインイン
購読する
お知らせ
小谷: AI スタックを所有する
any-llm を利用したオープンソース LLM ゲートウェイである Otari と、同じ基盤上に構築されたホスト型プラットフォームである Otari.ai をご紹介します。使用状況追跡、予算管理、ルーティング ポリシー、可観測性、チーム管理を備えた 1 つの API を通じてフロンティア モデルまたはオープンウェイト モデルを実行します。
クローズド ソースのフロンティア プロバイダーは、ツール、MCP サーバーの統合、実行環境、Web 検索、支出管理など、完全なスタックのように見えるものを提供します。次のプロジェクトに 1 つを選択するのは簡単なように感じます。
次に、コストや主権のため、あるいは単に実行できるという理由から、オープンウェイト モデルを実行することにします。そのスタックのほとんどが消えます。チャット エンドポイントを取得します。残りはあなたが再構築してください。
本日、私たちは、 any-llm 上に構築されたオープンソース LLM ゲートウェイである Otari と、それを中心に構築されたホスト型プラットフォームである Otari.ai をローンチします。これらを組み合わせることで、フロンティアかオープンウェイトか、ホスト型かセルフサービス型かにかかわらず、期待する開発者のエクスペリエンスと機能を放棄することなく、そして最も重要なことに、プライバシーを侵害することなく、あらゆるモデルを選択できるようになります。
Ohari は、ユーザー管理、プロバイダーのキー管理、使用量と予算の追跡、オープンソース モデルをより機能的にするためのツール セットなど、不足している部分をスタックにもたらします。
機能を犠牲にすることなく、コストとプライバシーを向上させます。また、Python にロックされているわけではなく、SDK のいずれかを介して接続するか、API に直接アクセスして接続できます。
フロンティアプロバイダーが出荷するのは重量だけではありません。これらは、コードの実行、Web 検索、文字起こし、画像生成、バッチ処理を提供します。ワークロードを Claude または GPT からオープンウェイト モデルに切り替える場合、これらのツールは付属しません。モデルは単純なチャット エンドポイントに退行し、アプリケーション コードは以前は必要なかった層を成長させる必要があります。

鉱石。
Ohari は、これらの機能をサーバー側のモデルに依存しないツールとして提供します。ゲートウェイは、ツール呼び出しをサポートするモデルにそれらをディスパッチします。
サンドボックス化されたコードの実行。 Docker で分離された Python REPL。モデルがコードを実行する必要があるときにサーバー側で呼び出されます。ツールを使用するモデルにはコード インタープリターが搭載されるようになりました。それを微調整する必要はありません。サンドボックスを作成するのではありません。それはただそこにあります。
ウェブ検索。 Tavily、Brave、または Exa をプラグインするオプションを備えた、すぐに使える SearXNG による最新情報の取得。オープンウェイト モデルがトレーニング カットオフでスタックすることはなくなりました。
音声入力、画像出力。 OpenAI 互換のトランスクリプションおよび画像生成エンドポイントにより、背後のモデルを交換してもマルチモーダル パイプラインは動作し続けます。
再ランキング。世代モデルに関係なく、LLM を利用した RAG のドキュメントの再ランキング。
バッチ処理。レイテンシーは問題ではなくコストが重要なワークロード向けの OpenAI 互換の非同期バッチ API。
オープンソース モデルを選択することは、機能が失われることを意味するべきではありません。小谷は土俵を平等にする。クローズドソースプロバイダーで使用するのと同じツールが、選択したどのモデルにも付属しています。オープンウェイト チャット モデルを Otar と組み合わせると、必要最低限​​の機能を備えたエージェント ランタイムではなく、完全に装備されたエージェント ランタイムが得られます。
そしてここで終わりではありません。次に llamafile 、 encoderfile 、および any-guardrail を利用したガードレールが登場するため、GPU がなくても、モデルの周囲の安全層と分類層が高速かつローカルで実行されます。
ゲートウェイが存在する理由の残りの半分は、各チームが最終的に自分たちで構築することになる、退屈だが重要なものです。オタリは以下の商品を発送します。
仮想 API キー: ハッシュ化され、名前付きで、オプションで有効期限が設定されるキーがユーザーにバインドされているため、クライアントは上流のプロバイダーの資格情報を参照することはありません。
ユーザー管理と予算: ユーザーごとの支出上限と設定可能なリセットウィンドウ。
使用量と支出の追跡

g: プロバイダー全体のリアルタイムのコスト計算。
レート制限: ユーザーごとに構成可能な RPM キャップ。ヒット数は Prometheus メトリクスとしてエクスポートされます。
ヘルスとプロメテウスのメトリクス。
プラットフォーム モード: 委任ベースのマルチテナント認証。これは、Otari.ai の構築の基盤です。
エンジンは小谷。 Otarari.ai は、自分で実行したくないときに入手できるものです。これは、OSS ゲートウェイ上に構築された、管理されたチーム指向のサーフェスです。
アイデンティティとチーム。ユーザー アカウント、ロールベースのアクセス権を持つ組織 (所有者、管理者、メンバー)、組織をスコープとするワークスペース。それぞれに独自のキー、メンバー、プレイグラウンド、支出ダッシュボードがあります。
ルーティングポリシー。ワークスペース レベルでプロバイダーとモデル間でリクエストがどのように流れるかを定義します。私たちは単純なフォールバック システムから始めており、近い将来、より複雑なルーターに拡張する予定です。
安全な金庫。プロバイダーの資格情報は保存時に暗号化されます。
管理されたプロバイダー。 API キーを持参しなくても、Otarari.ai を通じてフロンティア モデルにアクセスできます。透明なトークンごとの価格設定でウォレットに対して請求されます。
Mozilla.ai プロバイダー。ファーストパーティの管理プロバイダーは、オープンウェイト モデルにルーティングします。組織ごとに自動プロビジョニングされます。同じゲートウェイ、同じ予算、同じトレース。一級市民として無差別級に挑戦する。
マルチレベルの予算とウォレット。プロバイダー キーごとの使用制限に加え、メンバーごと、プロバイダー キーごとのきめ細かい制御のための上限があり、それぞれに独自のリセット ケイデンスがあります。
宣言的な構成。組織全体 (ワークスペース、プロバイダー キー、ルーティング ポリシー、予算、メンバーの予算ポリシー、カスタム価格設定、プラットフォーム キー) を 1 つの YAML ドキュメントで記述します。現在の状態との差分を計画し、それを Git にコミットし、ワンクリックで環境を再作成します。
可観測性。 OpenTelemetry を装備したアプリの OTLP トレース取り込み

ライセンス、OpenSearch を利用した分析、フィルタリングとセッションごとのメトリクスを備えたセッション エクスプローラー、およびコストを視覚化した使用状況ダッシュボードが含まれます。
小谷はオープンコアです。 Ohari.ai は、最上位の透明なビジネス レイヤーであり、同じエンジン、同じ API サーフェスであり、お客様に代わってホストされ、操作されます。速度を上げるためにプラットフォームを使用します。セルフホストによるプライバシーの保護: プロンプト、完了、トレース、使用ログが当社に影響することはありません。ユーザーが何を入力したかはわかりません。モデルが何を言い返すのかわかりません。アプリケーション コードを書き直すことなく、後で方向を切り替えることができます。ワイヤ形式は同じです。
私たちが重視しているもう 1 つの設計上の選択は、オープンウェイト モデルを第一級の市民にすることです。チェックボックスでもフォールバックでも、独自のインフラストラクチャを用意する必要があるものでもありません。同じダッシュボード、同じ予算、同じツール呼び出し、同じマネージド プロバイダー エクスペリエンス。それが、OSS プロジェクトとプラットフォームの両方、そして Mozilla.ai の広範な取り組みの背後にある賭けです。
Ohari.ai (ホスト型): サインアップしてウォレットにチャージし、フロンティアまたはオープンウェイト モデルへの通話を開始します。独自のキーを持参するか、マネージド プロバイダーを使用します。
Ohari (オープン ソース): リポジトリをクローンし、docker compose up を実行し、OpenAI クライアントにゲートウェイ URL を指定します。
フィードバックをお待ちしております。 GitHub で問題を報告するか、Mozilla.ai コミュニティ チャネル ( X 、 LinkedIn 、 Bluesky 、 Mastodon 、 Discord ) で私たちを見つけてください。あるいは、構築を開始して問題点を教えてください。
クラウド AI の価格は 2026 年に急速に変更されました。この投稿では、なぜ多くのチームがローカル モデルに戻るのか、Ollama や LM Studio などのツールの背後にあるトレードオフ、および開発者にとってポータビリティと所有権が大きな懸念事項になっている理由について考察します。
CQ Exchange: 国境なきエージェント
CQ Exchange は、プライベート名前空間とパブリック C

オモンズ。
インターフェースはもはや製品ではない
AI の未来は、現在のアプリを使用するエージェントではない可能性があります。エージェントが直接検査、変更、検証できる構造化表現を中心に再構築されたアプリである可能性があります。デッキ、ドキュメント、またはダッシュボードは、真実の情報源ではなく、出力になります。
cq の第一防御線 (エージェントのスタック オーバーフロー)
cq は、コーディング エージェントが解決パスを共有し、過去の失敗から学ぶのに役立ちます。私たちは Lauren Mushro と提携して VIBE✓ を cq に導入し、知識単位が共有メモリに入る前にレビューできるように支援しました。
購読してチームから最新のニュースやアイデアを入手してください

## Original Extract

Otari is an open-source LLM gateway built on any-llm. Otari.ai adds hosted routing, budgets, observability, and model management.

Sign in
Subscribe
Announcement
Otari: Own Your AI Stack
Meet Otari, an open-source LLM gateway powered by any-llm, and Otari.ai, the hosted platform built on the same foundation. Run frontier or open-weights models through one API with usage tracking, budget controls, routing policies, observability, and team management.
Closed-source frontier providers offer what looks like a complete stack: tools, MCP server integrations, execution environments, web search, spend controls, etc. Choosing one for your next project feels like a no-brainer.
Then you decide to run an open-weights model, for cost, sovereignty, or simply because you can. Most of that stack disappears. You get a chat endpoint. The rest is yours to rebuild.
Today we're launching Otari , an open-source LLM gateway built on top of any-llm , and Otari.ai , the hosted platform built around it. Together, they let you choose any model, whether it is frontier or open-weights, hosted or self-served, without giving up the developer experience and capabilities you expect and, most importantly, without compromising your privacy.
Otari brings the missing pieces to your stack: user management, provider key management, usage and budget tracking, and a set of tools to make open source models more capable.
Better cost and privacy without compromising on capabilities. And you are not locked into Python, you can connect via one of our SDKs or by hitting the API directly.
Frontier providers ship more than just weights. They ship code execution, web search, transcription, image generation, and batching. When you switch a workload from Claude or GPT to an open-weights model, those tools do not come with you. The model regresses to a simple chat endpoint, and your application code must grow a layer it did not need before.
Otari ships those capabilities as server-side, model-agnostic tools. The gateway dispatches them to any model that supports tool calls:
Sandboxed code execution. A Docker-isolated Python REPL, invoked server-side when a model needs to run code. Any tool-using model now has a code interpreter. You don't fine-tune for it; you don't write the sandbox; it's just there.
Web search. Current-information retrieval via SearXNG out of the box, with the option to plug in Tavily, Brave, or Exa. Your open-weights model is no longer stuck at its training cutoff.
Audio in, images out. OpenAI-compatible transcription and image generation endpoints, so multimodal pipelines keep working when you swap the model behind them.
Reranking. LLM-powered document reranking for RAG, independent of your generation model.
Batch processing. OpenAI-compatible asynchronous batch API for workloads where latency doesn't matter and cost does.
Choosing open-source models shouldn't mean losing capabilities. Otari levels the playing field. The same tools you use with closed-source providers are attached to whatever model you choose. Pair an open-weights chat model with Otari and you get a fully equipped agent runtime, not a stripped-down one.
And we're not stopping here. Guardrails powered by llamafile , encoderfile , and any-guardrail are next, so the safety and classification layers around your model run fast and locally, even without a GPU.
The other half of why a gateway exists is the boring, important stuff every team ends up building for itself. Otari ships it:
Virtual API keys: Hashed, named, optionally-expiring keys bound to a user, so clients never see your upstream provider credentials.
User management and budgets: Per-user spending caps with configurable reset windows.
Usage and spend tracking: Real-time cost calculation across providers.
Rate limiting: Configurable RPM caps per user, with hits exported as Prometheus metrics.
Health and Prometheus metrics.
Platform mode: Delegation-based multi-tenant authorization, which is the seam Otari.ai is built on.
Otari is the engine. Otari.ai is what you get when you don't want to run it yourself. It is the managed, team-oriented surface built on top of the OSS gateway.
Identity and teams. User accounts, organizations with role-based access (owner, admin, member), workspaces scoped to organizations, each with their own keys, members, playground, and spending dashboards.
Routing Policies. Define how requests flow across providers and models at the workspace level. We are starting with a simple fallback system and we will be expanding on more elaborate routers in the near future.
Secure vault. Provider credentials encrypted at rest.
Managed providers. Reach frontier models through Otari.ai without bringing your own API key. Billed against your wallet at transparent per-token pricing.
Mozilla.ai provider. A first-party managed provider routes to open-weights models. Auto-provisioned per organization. Same gateway, same budgets, same traces. Open-weights as a first-class citizen.
Multi-level budgets and wallets. Spend limits per provider key, plus per-member-per-provider-key caps for fine-grained control, each with their own reset cadence.
Declarative configuration. Describe an entire organization — workspaces, provider keys, routing policies, budgets, member budget policies, custom pricing, platform keys — in a single YAML document. Plan a diff against the current state, commit it to your Git, and re-create your environment with a single click.
Observability. OTLP trace ingest for any OpenTelemetry-instrumented application, OpenSearch-backed analytics, a session explorer with filtering and per-session metrics, and usage dashboards with cost visualization.
Otari is open core. Otari.ai is a transparent business layer on top, same engine, same API surface, just hosted and operated for you. Use the platform for velocity. Self-host to keep your privacy: Your prompts, your completions, your traces, and your usage logs never touch us. We don't see what your users type. We don't see what your models say back. Switch directions later without rewriting your application code: the wire format is the same.
The other design choice we care about is making open-weights models a first-class citizen. Not a checkbox, not a fallback, not a thing you have to bring your own infrastructure for. Same dashboards, same budgets, same tool calls, same managed provider experience. That's the bet behind both the OSS project and the platform, and behind Mozilla.ai's broader work.
Otari.ai (hosted): sign up, top up your wallet, start calling frontier or open-weights models. Bring your own keys or use the managed providers.
Otari (open source): clone the repo, run docker compose up, point your OpenAI client at the gateway URL.
We'd love your feedback. File issues on GitHub, find us in the Mozilla.ai community channels ( X , LinkedIn , Bluesky , Mastodon , Discord ), or just start building and tell us what breaks.
Cloud AI pricing changed fast in 2026. This post looks at why more teams are moving back to local models, the tradeoffs behind tools like Ollama and LM Studio, and why portability and ownership are becoming bigger concerns for developers.
cq exchange: Agents without Borders
cq exchange gives agents a shared place to store and retrieve experience-driven knowledge through private namespaces and a public commons.
The Interface Is No Longer the Product
The future of AI may not be agents using today’s apps. It may be apps rebuilt around structured representations agents can inspect, modify, and validate directly. The deck, doc, or dashboard becomes the output, not the source of truth.
First Line of Defense for cq (Stack Overflow for Agents)
cq helps coding agents share resolution paths and learn from past failures. We partnered with Lauren Mushro to bring VIBE✓ into cq and help review knowledge units before they enter shared memory.
Subscribe to get the latest news and ideas from our team
