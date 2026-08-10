---
source: "https://codapult.dev/blog/why-cloning-beats-prompting-ai-saas"
hn_url: "https://news.ycombinator.com/item?id=49248411"
title: "I spent $200 in API credits asking AI agents to scaffold a Next.js starter"
article_title: "I Spent $200 in API Credits Prompting an Agent to Build a Starter. Here’s Why I Still Clone Boilerplates."
author: "vladzoff"
captured_at: "2026-08-10T19:48:35Z"
capture_tool: "hn-digest"
hn_id: 49248411
score: 1
comments: 0
posted_at: "2026-08-10T19:20:39Z"
tags:
  - hacker-news
  - translated
---

# I spent $200 in API credits asking AI agents to scaffold a Next.js starter

- HN: [49248411](https://news.ycombinator.com/item?id=49248411)
- Source: [codapult.dev](https://codapult.dev/blog/why-cloning-beats-prompting-ai-saas)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T19:20:39Z

## Translation

タイトル: AI エージェントに Next.js スターターのスキャフォールディングを依頼するために API クレジットを 200 ドル費やしました
記事のタイトル: エージェントにスターターの構築を促す API クレジットに 200 ドルを費やしました。私が今でもボイラープレートのクローンを作成する理由は次のとおりです。
説明: Opus 4.8 または GPT-5.6 に複数ページの Next.js 16 アプリのスキャフォールディングを要求すると、API トークンとエージェントの子守りの時間に数百ドルが費やされます。クローン作成がプロンプトよりも優れている理由の背後にある計算は次のとおりです。

記事本文:
エージェントにスターターの構築を促す API クレジットに 200 ドルを費やしました。私が今でもボイラープレートのクローンを作成する理由は次のとおりです。 Codapult 価格設定プラグイン ブログ ドキュメント デモ ビルダー向け SaaS 定型文
© 2026 コダパルト.無断転載を禁じます。
ブログに戻る July 30, 2026 · 6 min read · Codapult チーム 私は、エージェントにスターターの構築を促す API クレジットに 200 ドルを費やしました。私が今でもボイラープレートのクローンを作成する理由は次のとおりです。
Opus 4.8 または GPT-5.6 に複数ページの Next.js 16 アプリのスキャフォールディングを要求すると、API トークンと何時間ものエージェントの子守りに数百ドルが費やされます。クローン作成がプロンプトよりも優れている理由の背後にある計算は次のとおりです。
先月、私は単純な仮説を検証することにしました。CLI エージェント内で実行されている Claude Opus 4.8 や GPT-5.6 のような最上位モデルでは、SaaS ボイラープレートは正式に廃止されたのでしょうか?
新しいディレクトリを開き、エージェント ループを起動し、Better-Auth、Stripe サブスクリプション、Tailwind v4、および Drizzle ORM を使用して実稼働対応の Next.js 16 アプリをスキャフォールディングするように依頼しました。
3 時間後、180 ドルの Anthropic 請求書を受け取り、私は答えを得ました。
複数ファイルコンテキストの残酷な計算
「AI コーディングは基本的に無料です」と人々が言うとき、彼らは通常、単一の UI コンポーネントまたはスタンドアロンのユーティリティ関数の生成について話しています。それには何ペニーもかかります。
この計算は、相互接続された 30 以上のファイルにわたるマルチページ アーキテクチャの構築をエージェントに依頼した瞬間に完全に崩れてしまいます。
エージェント ワークフローはコードを記述するだけではなく、ディレクトリ ツリーを検査し、ターミナル チェックを実行し、型定義を読み取り、複数のパスにわたってファイルを編集します。エージェントがツール呼び出しを行ったり、 npm run build を実行したり、サーバー コンポーネントが再レンダリングに失敗した理由を確認したりするたびに、リポジトリ コンテキスト全体がモデルに再送信されます。
パスごとに数千の推論トークンを追加し、認証、セッション Cookie、および St を結び付ける単一のマルチターン セッションを追加します。

成熟した Webhook は、生の API クレジットで 150 ドルから 300 ドル以上を簡単に消費します。
エージェントに標準の認証フローをゼロから繰り返し再作成させるために、事実上、AI プロバイダーに数百ドルを支払っていることになります。
凝集の問題: 70 以上のモジュールが一緒に「プロンプト」を表示しない
Opus 4.8 や GPT-5.6 などのモデルは、局所的なロジックの生成に優れています。しかし、実稼働 SaaS は分離されたファイルのコレクションではなく、密接に結合された依存関係の網です。
実際の基盤には、まったく同じ言語を話す 70 個以上のモジュールと 30 個以上のデータベース テーブルが必要です。
認証とアイデンティティ: パスキー、2FA、SSO、およびマルチテナント組織向けに設定されたより良い認証。
請求ライフサイクル: Stripe と LemonSqueezy の両方の Webhook ハンドラー、サブスクリプション シート、猶予期間、使用量測定。
システム層: 管理ダッシュボード、詳細な RBAC、および i18n ローカリゼーション ルーティング。
エージェントにこのグラフを最初から作成するよう促すと、タイプ ドリフトは避けられません。データベース スキーマはセッション ミドルウェアとは若干異なります。 RBAC チェックはサーバー アクションのエッジ ケースを見逃します。ワークスペース ID が文字列ではなく未定義として渡されると、支払い Webhook が中断します。
これらのモジュール間統合のバグを修正するには、基本的な機能を取得して互いの破壊を防ぐためだけにトークンを焼き尽くす、継続的なプロンプトのやり取りが必要です。
エージェントは、 npm run dev を介して localhost:3000 で実行されるものを構築することに優れています。これらは、実稼働インフラストラクチャ用のコードを準備する際の信頼性が低いことで知られています。
プロジェクトを本番環境に向けて準備するには、Infrastructure as Code (IaC) をセットアップする必要があります。
ローカル運用ミラー用の有効な Docker Compose ファイル。
クリーンな Terraform または Pulumi マニフェスト。
すぐにデプロイできる Kubernetes Helm チャート。
PostgreSQL とエッジ対応 Turso SQLit 間の交換など、マルチデータベースの柔軟性

e 環境変数 ( DB_PROVIDER ) を介して。
アプリ コードと一緒に有効な Helm チャートまたは Terraform マニフェストを生成するようエージェントに依頼すると、構成キーが幻覚を示したり、古い構文が出力されたりすることがよくあります。 AI エージェントを使用して失敗したデプロイメント パイプラインをデバッグするには、API コンピューティングだけで 30 ～ 50 ドルの費用がかかる可能性があり、YAML ファイルと HCL ファイルの軽微な構文エラーの修正にのみ費やされます。
プロンプト セッションによって生成されたコードは、時間の静的なスナップショットです。リネージも上流リポジトリもありません。
今から 6 か月後、Next.js に重大な API 変更が導入されたり、Tailwind v4 がメジャー パッチをリリースしたり、認証の依存関係にセキュリティの脆弱性が表面化したりすると、AI が生成したスターターは即座に技術的負債となります。唯一の手段は、コードベース全体をエージェントにフィードバックし、リファクタリング中にカスタム ビジネス ロジックが破壊されないことを祈ることです。
専用のボイラープレートは、Git アップストリーム リモートを提供します。コアのセキュリティ パッチ、フレームワークの移行、または新機能が削除された場合は、 git pullupstream を使用して直接プルします。あなたは、保守不可能な単発生成ではなく、アクティブなメンテナによってサポートされる生きたコードベースを所有しています。
API トークンが完全に無料だったとしても、時間的なコストがかかります。
現在のエージェントは高速ですが、完全なアプリケーションが自動的に構築される間、プロンプトを 1 つ作成して立ち去ることができるほど自律的ではありません。あなたは端末の前に座ってプロセスの子守りをすることを余儀なくされます。
差分を確認して、エージェントが Next.js 16 の非同期リクエスト API ( cookies() 、 headers() 、または Route params ) を台無しにしていないことを確認します。
Webhook ルートが署名チェックに失敗したため、ターミナル実行の承認とビルドの監視が失敗します。
サーバーとクライアントの状態間の微妙なハイドレーションの不一致を修正することでモデルをガイドします。
ビルド中にセットアップをカスタマイズすることにした場合は、

たとえば、組織の請求を維持しながらマルチテナンシーを削除します。エージェントに新たに作成したコードをリファクタリングするよう依頼すると、通常、インポートが壊れたり、スキーマ テーブルが孤立したりする結果になります。
これを、 npx create-codapult のような対話型 CLI インストーラーの実行と比較してください。このインストーラーは、機能フラグの入力を事前に要求し、クリーンでトリミングされたコードベースを数秒で即座に出力します。
本当に重要なことのためにコンピューティング予算を節約しましょう
最上位モデルは、製品を際立たせる複雑でユニークなビジネス ロジックを記述するのに優れています。高度な推論コンピューティングを使用して、標準の Stripe サブスクリプション ライフサイクル、Tailwind v4 構成、DB スキーマを再生成することは資本の無駄です。
AI に対して、レンガの基礎を築くのと同等の割増料金を支払っていることになります。
これがまさに私が Codapult を構築した理由です。
これは、すぐに疲れたりコンテキスト ドレインから解放されるように設計された、クリーンなモジュール式の Next.js 16 SaaS スターターです。
完全な基盤: Better-Auth (パスキー、SSO、2FA)、Stripe & LemonSqueezy 課金、Tailwind v4、Drizzle ORM、およびすぐに使える i18n。
実稼働インフラストラクチャ: 事前構成された Docker、Terraform、Kubernetes Helm チャート、および柔軟な DB ルーティング (PostgreSQL / Turso SQLite)。
拡張可能なアーキテクチャ: 後で特殊な機能が必要ですか?コア フレームワークには触れずに、AI Kit 、 CRM 、または Helpdesk などの公式モジュールをドロップインします。
保守可能なライフタイム: Git アップストリーム CLI ワークフローを通じて、簡単に最新情報を維持できます。
git clone を実行するか、 npx create-codapult を使用して .env 変数をドロップし、すぐに実際の機能の構築を開始します。
コア製品用に上位トークンを保存してください。
Codapult には、すぐに運用可能な完全な Next.js 16 スタックが含まれており、Better-Auth、Stripe/LemonSqueezy 請求ライフサイクル、Tailwind v4、Drizzle ORM、IaC 構成、モジュラー プラグインが含まれており、コストを節約できます。

API コンピューティングに数百ドル、エージェント オーケストレーションに数時間かかります。ドキュメントでアーキテクチャを調べてください。

## Original Extract

Prompting Opus 4.8 or GPT-5.6 to scaffold a multi-page Next.js 16 app burns hundreds of dollars in API tokens and hours of agent baby-sitting. Here is the math behind why cloning beats prompting.

I Spent $200 in API Credits Prompting an Agent to Build a Starter. Here’s Why I Still Clone Boilerplates. Codapult Pricing Plugins Blog Docs Demo The SaaS Boilerplate for Builders
© 2026 Codapult . All rights reserved.
Back to blog July 30, 2026 · 6 min read · Codapult Team I Spent $200 in API Credits Prompting an Agent to Build a Starter. Here’s Why I Still Clone Boilerplates.
Prompting Opus 4.8 or GPT-5.6 to scaffold a multi-page Next.js 16 app burns hundreds of dollars in API tokens and hours of agent baby-sitting. Here is the math behind why cloning beats prompting.
Last month I decided to test a simple hypothesis: with top-tier models like Claude Opus 4.8 and GPT-5.6 running inside CLI agents, is the SaaS boilerplate officially dead?
I opened a fresh directory, booted up an agentic loop, and asked it to scaffold a production-ready Next.js 16 app with Better-Auth, Stripe subscriptions, Tailwind v4, and Drizzle ORM.
Three hours and a $180 Anthropic invoice later, I had my answer.
The Brutal Math of Multi-File Context
When people say "AI coding is basically free," they are usually talking about generating a single UI component or a standalone utility function. That costs pennies.
That math completely falls apart the second you ask an agent to build a multi-page architecture across 30+ interconnected files.
Agentic workflows don't just write code—they inspect directory trees, run terminal checks, read type definitions, and edit files across multiple passes. Every time the agent makes a tool call, runs npm run build , or checks why a Server Component failed to re-render, it re-sends the entire repository context back into the model.
Add thousands of reasoning tokens per pass, and a single multi-turn session to wire up auth, session cookies, and Stripe webhooks easily burns through $150 to $300+ in raw API credits .
You are effectively paying AI providers hundreds of dollars to have an agent repeatedly re-invent standard authentication flows from scratch.
The Cohesion Problem: 70+ Modules Don't Just "Prompt" Together
Models like Opus 4.8 and GPT-5.6 are brilliant at generating localized logic. But a production SaaS isn't a collection of isolated files—it is a tightly coupled web of dependencies.
A real foundation requires 70+ modules and 30+ database tables all speaking the exact same language:
Auth & Identity: Better-Auth configured for Passkeys, 2FA, SSO, and multi-tenant Organizations.
Billing Lifecycles: Webhook handlers for both Stripe and LemonSqueezy, subscription seats, grace periods, and usage metering.
System Layer: Admin dashboards, granular RBAC, and i18n localization routing.
When you try to prompt an agent into building this graph from scratch, type drift is inevitable. The database schema slightly diverges from the session middleware; the RBAC checks miss an edge case in server actions; the payment webhooks break when a workspace ID is passed as undefined instead of a string.
Fixing these cross-module integration bugs requires constant back-and-forth prompts, burning through tokens just to get basic features to stop breaking each other.
Agents excel at building things that run on localhost:3000 via npm run dev . They are notoriously unreliable when preparing code for production infrastructure.
Getting a project ready for production means setting up Infrastructure as Code (IaC):
Valid Docker Compose files for local production mirrors.
Clean Terraform or Pulumi manifests.
Ready-to-deploy Kubernetes Helm charts .
Multi-database flexibility, like swapping between PostgreSQL and edge-ready Turso SQLite via an environment variable ( DB_PROVIDER ).
When you ask an agent to generate valid Helm charts or Terraform manifests alongside app code, it frequently hallucinates configuration keys or outputs outdated syntax. Debugging a failing deployment pipeline through an AI agent can cost $30 to $50 in API compute alone, purely spent fixing minor syntax errors in YAML and HCL files.
Code generated by a prompt session is a static snapshot in time. It has no lineage and no upstream repository.
Six months from now, when Next.js introduces breaking API changes, Tailwind v4 releases a major patch, or security vulnerabilities surface in your auth dependencies, your AI-generated starter becomes instant technical debt. Your only recourse is to feed the whole codebase back into an agent and hope it doesn't break your custom business logic while refactoring.
A purpose-built boilerplate provides a Git upstream remote . When core security patches, framework migrations, or new features drop, you pull them directly with git pull upstream . You own a living codebase backed by an active maintainer, not an unmaintainable single-shot generation.
Even if API tokens were completely free, there is the time cost.
Current agents are fast, but they aren't autonomous enough for you to write one prompt and walk away while a full application builds itself. You are forced to sit at your terminal baby-sitting the process:
Reviewing diffs to make sure the agent didn't mess up Next.js 16's async request APIs ( cookies() , headers() , or route params ).
Approving terminal executions and watching builds fail because a webhook route missed a signature check.
Guiding the model through fixing subtle hydration mismatches between server and client states.
If you ever decide to customize the setup mid-build—for instance, removing multi-tenancy while keeping organization billing—asking an agent to refactor its own freshly written code usually results in broken imports and orphaned schema tables.
Compare that to running an interactive CLI installer like npx create-codapult , which prompts you for feature flags upfront and instantly outputs a clean, trimmed codebase in seconds.
Save Your Compute Budget for What Actually Matters
Top-tier models are incredible at writing complex, unique business logic that sets your product apart. Using high-reasoning compute to re-generate standard Stripe subscription lifecycles, Tailwind v4 configs, and DB schemas is a waste of capital.
You’re paying premium rates for the AI equivalent of laying a brick foundation.
This is exactly why I built Codapult .
It is a clean, modular Next.js 16 SaaS starter designed to save you from prompt fatigue and context drain:
Complete Foundation: Better-Auth (Passkeys, SSO, 2FA), Stripe & LemonSqueezy billing, Tailwind v4, Drizzle ORM, and i18n out of the box.
Production Infrastructure: Pre-configured Docker, Terraform, Kubernetes Helm charts, and flexible DB routing (PostgreSQL / Turso SQLite).
Extensible Architecture: Need specialized capabilities later? Drop in official modules like the AI Kit , CRM , or Helpdesk without touching the core framework.
Maintainable Lifetime: Stay updated effortlessly via the Git upstream CLI workflow.
You run git clone or use npx create-codapult , drop in your .env variables, and start building actual features immediately.
Save your high-tier tokens for your core product.
Codapult includes a complete, production-ready Next.js 16 stack out of the box — Better-Auth, Stripe/LemonSqueezy billing lifecycles, Tailwind v4, Drizzle ORM, IaC configs, and modular plugins — saving you hundreds of dollars in API compute and hours of agent orchestration. Explore the architecture in the documentation .
