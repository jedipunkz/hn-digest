---
source: "https://github.com/louisbrulenaudet/monorepo-template"
hn_url: "https://news.ycombinator.com/item?id=49034424"
title: "Show HN: A monorepo where AI agents can safely build and maintain applications"
article_title: "GitHub - louisbrulenaudet/monorepo-template: A concise, AI native and production-grade starter template for full-stack applications using pnpm, Turborepo and Cloudflare Workers for edge-native execution 🚚⛅ · GitHub"
author: "brulenaudet"
captured_at: "2026-07-24T12:12:54Z"
capture_tool: "hn-digest"
hn_id: 49034424
score: 1
comments: 0
posted_at: "2026-07-24T12:10:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A monorepo where AI agents can safely build and maintain applications

- HN: [49034424](https://news.ycombinator.com/item?id=49034424)
- Source: [github.com](https://github.com/louisbrulenaudet/monorepo-template)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T12:10:40Z

## Translation

タイトル: Show HN: AI エージェントがアプリケーションを安全に構築および保守できるモノリポジトリ
記事のタイトル: GitHub - louisbrulenaudet/monorepo-template: エッジネイティブ実行に pnpm、Turborepo、Cloudflare Workers を使用したフルスタック アプリケーション用の簡潔な AI ネイティブおよび運用グレードのスターター テンプレート 🚚⛅ · GitHub
説明: エッジネイティブ実行用に pnpm、Turborepo、Cloudflare Workers を使用するフルスタック アプリケーション用の簡潔な AI ネイティブかつ運用グレードのスターター テンプレート 🚚⛅ - louisbrulenaudet/monorepo-template
HN テキスト: 皆さん、こんにちは。私はアプリケーションを迅速に運用環境に導入するための理想的なテンプレートの開発に何十時間も費やしました。これは、AI ネイティブの方法で、Cloudflare インフラストラクチャ上で低コストで、スケーラブルなプロトタイプを迅速に提供したい人にとって興味深いものになるでしょう。このテンプレートには次のものが含まれています。 - Zod を使用した厳密に型指定された TypeScript の pnpm モノリポジトリにより、コントラクトを複製することなく、React アプリケーションとバックエンド (ここでは Hono とネイティブ Zod バリデーターを使用) の両方を開発できます。
- AI エージェント (現時点では主にクロード コードとカーソル) の機能を最大限に活用するように完全に設計されたセットアップ。コンテキストを汚染することなく規則を適用する正確でパス スコープのルール、超高速フォーマットとリンティングのための OXC ベースのフック、特に Context7 と Cloudflare MCP のおかげで最適なコンテキスト収集を可能にするサブエージェントを備えています。
- Vite 8、TanStack Router、TanStack Query、および React 19 の次世代構成。
- 人間と AI エージェントの両方の高速かつ直感的な開発を可能にする一連の make コマンド。
- 数十のマイクロサービスの導入をサポートすることを目的とした明確なドキュメント。マルチプル間でサポートされる RPC プロトコルのおかげで、アプリケーション間の通信をゼロコストで簡素化します。

ル労働者など。
- スケーリングコードベースでの開発に必要なすべてのコマンドの並列実行と、エッジでの数十のサービスの迅速なデプロイメントを可能にするターボレポ構成。私はすでにこのテンプレートを実稼働環境で、認証付きの MCP 開発、個人アプリケーション、ハッカソンのコンテキストで実戦テストすることができました。特に、このテンプレートのおかげで、Hexa で開催された {Tech: Europe} パリ AI ハッカソンでファイナリストとしてフィニッシュすることができました。これは、ハーネスへの投資のおかげで、セキュリティとベスト プラクティスの一定の保証を維持しながら、起業家がコードやアプリケーションのプロトタイプを迅速に作成するのに非常に役立ち、摩擦を最小限に抑え、バックエンドとフロントエンド間または複数のサービス (キュー、RPC など) 間でコードの重複を回避することを懸念するさまざまな開発者にとって魅力的であると私は信じています。このリポジトリは、最新の比較的十分に文書化されたテクノロジにネイティブに基づいており、TypeScript を日常的に使用しない人を含め、可能な限り最高の開発者エクスペリエンスを提供することを目的としています (実際には、使用されている詳細な構成により、最小限の変更でバックエンド マイクロサービスと React アプリケーションの両方の開発に摩擦なく、比較的安全な方法で適応することができます)。 Cloudflare インフラストラクチャとモノリポジトリの概念に精通している場合は、特に AI エージェントのセットアップに関して検討することを強くお勧めします。私が本当に楽しく使っているこのツールをさらに改善するための提案を歓迎します。このテンプレートで使用されているさまざまな依存関係の提供者全員に、素晴らしい体験を提供していただき感謝します。また、このモノリポジトリで何が提供できるのかを楽しみにしています。

記事本文:
GitHub - louisbrulenaudet/monorepo-template: エッジネイティブ実行用に pnpm、Turborepo、Cloudflare Workers を使用したフルスタック アプリケーション用の簡潔な AI ネイティブかつ運用グレードのスターター テンプレート 🚚⛅ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アルを解雇する

えーっと
{{ メッセージ }}
ルイブルノーデ
/
モノリポジトリテンプレート
パブリックテンプレート
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
47 コミット 47 コミット .agents/ スキル .agents/ スキル .claude .claude .cursor .cursor .github .github .husky .husky .vscode .vscode アプリ アプリ アセット アセット フック フック パッケージを作成する パッケージ .cursorignore .cursorignore .cursorindexingignore .cursorindexingignore .editorconfig .editorconfig .gitignore .gitignore .mcp.json .mcp.json .npmrc .npmrc .nvmrc .nvmrc .oxfmtrc.json .oxfmtrc.json .oxlintrc.json .oxlintrc.json .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml skill-lock.json skill-lock.json tsconfig.json tsconfig.jsonturbo.jsonturbo.json すべて表示ファイル リポジトリ ファイルのナビゲーション
Cloudflare、Hono、React、Vite、Tailwind を備えた pnpm に基づく Monorepo スターター 🚚⛅
Turborepo、Cloudflare Workers、Hono、React (Vite)、Tailwind CSS v4 、および TanStack Router/Query を備えた pnpm ワークスペース上に構築された、最小限の本番指向のモノリポ スターターです。 AI 対応、エッジ展開向けに設計され、スケールする実稼働プロジェクト向けに構造化されています。
今日のスターター アプリ (worker-api 、front-app)。以下のプレフィックスは、リポジトリがどのように成長するかを示しています。
モノレポ/
§── apps/ # ワーカーとフロントエンド
│ §── worker-api/ # REST API ゲートウェイ
│ └── フロントアプリ/ # React SPA (Vite + TanStack)
§── パッケージ/ # 共有 @repo/* パッケージ
│ §── dtos-common/ # Zod ワイヤーコントラクト (api / rpc / queue /

ウェブフック)
│ §── enums-common/ # 制約付き文字列値の共有 (`as const`)
│ └── typescript-config/ # TypeScript 設定のプリセット
§── make/ # 共有 Makefile フラグメント (make/README.md を参照)
§── フック/ # AI エージェント フック (Husky ではありません - フック/README.md を参照)
§── package.json # ルートパッケージの設定
§── pnpm-workspace.yaml # ワークスペースの設定
§──turbo.json # ターボレポ設定
━── tsconfig.json # ルート TypeScript 設定
アーキテクチャコンポーネント
モノリポジトリは、バックエンド サービスとフロントエンド アプリケーション、さらに共通機能の共有パッケージという 2 つの主要カテゴリに編成されています。
フローチャートTB
サブグラフエントリ [パブリックエントリ]
LR方向
フロント["フロント-* :517x"]
Ext["外部プロバイダ"]
McpClients["MCP クライアント"]
終わり
サブグラフ publicWorkers [公務員]
LR方向
ゲートウェイ["ワーカー-API :8700"]
Webhook["Webhook-* :876x"]
Mcp["mcp-* :878x"]
終わり
サブグラフ privateWorkers [プライベート ワーカー]
LR方向
Biz["worker-* (RPC のみ)"]
キュー["キュー-*"]
終わり
サブグラフ共有 [共有パッケージ]
LR方向
Enums["@repo/enums-common"]
DTO["@repo/dtos-common"]
列挙型 --> DTO
終わり
フロント --> ゲートウェイ
内線 --> Webhook
McpClients --> Mcp
ゲートウェイ --> ビジネス
Webhook --> ビジネス
MCP --> Biz
ゲートウェイ --> キュー
Webhook --> キュー
ビジネス --> キュー
フロント -.-> 共有
publicWorkers -.-> 共有
privateWorkers -.-> 共有
読み込み中
バックエンドサービス
Cloudflareワーカーはランタイムロールごとに編成されています。
worker-api - パブリック HTTP ゲートウェイ (Hono): CORS、検証、ルーティング。 RPC 経由で内部ワーカーを調整します。
worker-* - サービス バインディング RPC のみを介したビジネス ロジック (本番環境ではパブリック ルートはありません)。 src/db/ の下の Drizzle スキーマとそのデータベースのバインディングを所有する可能性があります (排他的所有者)。
queue-* - キュー専用コンシューマー (queue() ハンドラー)。メッセージはプロ仕様にすることができます

worker-api 、 worker-* 、または webhook-* によって引き起こされます。ローカル HTTP デバッグ パスが役立つ場合は、デュアル ハンドラー レイアウトを使用します。
webhook-* - 外部プロバイダー コールバックのパブリック HTTP イングレス。 RPC またはキュー経由で作業を転送します。
mcp-* - パブリック HTTP MCP サーバー。 RPC 経由で worker-* を呼び出すシン ツール。
共有パッケージ/db-* スキーマ パッケージを作成しないでください。 Drizzle スキーマを所有するアプリの src/db/ の下に置き、1 つの DB バインディング所有者を保持します。他のアプリはサービス バインディング RPC (またはキュー) 経由でそのデータにアクセスします。同じ DB バインディングを複数のアプリにアタッチしないでください。
フロントアプリ - Cloudflare Workers にデプロイされた React SPA (Vite 8、Tailwind v4、TanStack Router/Query)。 worker-api と通信するのは HTTP 経由のみであり、サービス バインディングを経由することはありません。
タスク
場所
新しい API ルート
apps/worker-api/src/routes/<feature>.ts → src/index.ts にマウント
HTTP Zod スキーマ
パッケージ/dtos-common/src/api/<機能>.ts
RPC / キュー / Webhook スキーマ
Packages/dtos-common/src/<layer>/<feature>.ts (レイヤー: rpc 、 queue 、または webhook )
共有文字列値セット
パッケージ/enums-common/src/
ワーカーローカルの値セット
apps/<ワーカー>/src/enums/
DB スキーマ (1 所有者)
apps/<owner>/src/db/ - パッケージ/db-* は使用しません
フロントエンドAPIクライアント
apps/front-app/src/services/worker-api/<feature>.ts
フロントエンド ページ + ルート
apps/front-app/src/pages/ + src/routes/
ローカル ワーカーの秘密
apps/<worker>/.dev.vars ( .dev.vars.example から)
フロントエンド環境
apps/front-app/.env.local ( .env.example から)
はじめに
Node.js 22+ (ルート package.json エンジンを参照)。バージョン管理には fnm をお勧めします
ルートの packageManager フィールド経由の pnpm (Corepack を推奨)
Cloudflareアカウントは、ログイン/デプロイ/リモートワーカー機能が必要な場合のみ
インストールする
ログイン番号をオプションにする - リモート Wrangler 機能の Cloudflare 認証
make prepare # Husky プリコミットフック
最初の実行前に env テンプレートをコピーします。
W

オーカー: apps/<worker>/.dev.vars.example → .dev.vars
フロントエンド: apps/front-app/.env.example → .env.local
ワークスペース リンクの一貫性を保つために、raw pnpm install よりも make install を優先します。
このリポジトリは、ルート package.json の packageManager を介して pnpm を固定します。
最初の成功した実行 (ローカルで確認)
すべての開発サーバーをリポジトリ ルートから起動します。
開発を行う
API を確認します: GET http://localhost:8700/api/v1/health
フロントエンドを開きます: http://localhost:5174
1 つのパッケージに集中して作業します: make dev SCOPE=worker-api (「スコープ」を参照)。
コマンド
説明
インストールする
ワークスペース パッケージのインストールとリンク
インストール凍結
凍結されたロックファイル (CI) を使用してインストールする
ログイン
Cloudflare にログインします (リポジトリに固定された Wrangler)
アップデート
依存関係を最新のものに更新します (pnpm カタログを書き換えます)
チェックする
lint + 形式チェック (型チェックなし)
シ
lint + フォーマット + チェックタイプ (ローカル PR ゲート)
デプロイする
すべてのアプリ/ワーカーをデプロイします (Turborepo 経由)
構築する
すべてのパッケージとアプリをビルドする (Turborepo 経由)
フォーマット
oxfmt による書式設定の自動修正
糸くず
oxlint を使用して lint の問題を自動修正する
開発者
すべての開発サーバーを起動します (Turborepo 経由)
プレビュー
実稼働ビルドをローカルでプレビューする (Turborepo 経由)
種類
アプリでworker-configuration.d.tsを生成する
チェックタイプ
すべてのワーカーとパッケージにわたる TypeScript
準備する
Husky git フックをインストールまたは再インストールする
ハスキーステータス
ハスキーフックのステータスを表示
スキルアップデート
ロックされたエージェント スキルを更新します (make/README.md を参照)
スコーピング (pnpm / Turborepo)
ターボバックアップされたルート ターゲットのオプションの変数 ( dev 、 build 、 ci など):
ニーモニック: 87xx = ワーカー (ゲートウェイ → ビジネス → キュー → Webhook → MCP → 予約)。フロントエンドは Vite の 51xx / 41xx を使用します。
サービス
パス
開発者
プレビュー
ワーカー API
apps/worker-api/wrangler.jsonc
8700
-
フロントアプリ
apps/front-app/vite.config.ts
5174
4174
注:
ワーカー: wrangler.jsonc に dev.port を設定し、 package.json に monorepo.devPort を設定します。インスペクターポを使用する

rt: 0 。
フロントエンド: Vite の server.port /review.port を strictPort: true に設定します。
ロールの範囲内で次の空きポートを割り当てます。 RPC およびキュー専用アプリはスタンドアロンの Wrangler dev 用のローカル ポートを引き続き取得しますが、運用環境にはパブリック URL がありません。
バインディングをテストするときは、マルチ構成のローカル実行を優先します (最初の -c は HTTP プライマリです)。
1. 新しいCloudflareワーカーを作成する
目的
プレフィックス
例
HTTPゲートウェイ
ワーカー API (スティッキー)
ワーカー API
ビジネスロジック（RPC）
労働者-
ワーカーアカウント
キュー専用コンシューマ
待ち行列-
キューメール
Webhook イングレス
ウェブフック-
Webhook の例
MCPサーバー
mcp-
mcpツール
フロントエンドアプリケーション
フロント-
フロントアプリ
主な特徴
ゲートウェイ ( worker-api ): パブリック HTTP のみ。リクエストを検証し、RPC 経由で worker-* を呼び出します。
ビジネス ワーカー ( worker-* ): 本番環境では RPC のみ ( WorkerEntrypoint )。 src/db/ の下に Drizzle スキーマとそのデータベースのバインディング (排他的) を所有している可能性があります。キューも消費する場合は、このプレフィックスを保持し、デュアル ハンドラー レイアウトを使用します。
キューのみ ( queue-* ): 本番環境にパブリック HTTP を持たない queue() コンシューマー。そのデータの唯一の書き込み者である場合、スキーマを所有することができます。
Webhook ワーカー ( webhook-* ): 外部コールバック用のパブリック HTTP。 RPC またはキュー経由で転送します。
MCP サーバー ( mcp-* ): パブリック HTTP MCP トランスポート。 RPC 経由で worker-* を呼び出すシン ツール - このサーフェス上で長期間有効な資格情報をローテーションしないでください。
フロントエンド (front-*): React + Vite;ゲートウェイへの HTTP のみ - サービス バインディングは行いません。
共有パッケージ/db-* スキーマ パッケージを作成しないでください。
足場チェックリスト (既存のアプリからコピー)
ジェネレーター CLI はありません。 apps/ の下に最も近い兄弟をコピーし、モノリポジトリに接続します。
apps/worker-api (または別の最も近い一致) を apps/<prefix-name> にコピーします (例: apps/worker-account )。
package.json name 、wrangler.jsonc name 、および表示文字列の名前を変更します。
ポートを割り当てる

開発ポート レジストリ - wrangler.jsonc で dev.port /inspector_port: 0 を設定し、 package.json で monorepo.devPort を設定します。
make/app.mk を含む 4 行の Makefile を保存します ( make/README.md を参照)。
@repo/typescript-config/workers.json (または一致するプリセット) を拡張します。 .dev.vars.example を追加します。
インストールしてgenと入力します。
インストールする
型を作る
既存のアプリから Wrangler パターン (compatibility_flags 、observability 、env.staging / env.production ) をコピーします。 .cursor/rules/backend/workers-config.mdc を参照してください。
リポジトリ ルートからスコープを指定した make を優先します。
make dev SCOPE=ワーカー名
または、ワーカーの Makefile を使用します。
cd アプリ/ワーカー名
開発を行う
これにより、apps/worker-name/package.json で定義された開発スクリプトが実行されます。
ターミナルに表示されているポートを開きます (例: http://localhost:8721 )
各ワーカー Makefile は、 make dev 、 make format 、 make lint 、 make types 、 make check-types 、 makedeploy を公開します。
ワーカー間のサービス バインディングのテスト
単一の複数構成の Wrangler dev を優先します (最初の -c は HTTP プライマリです)。
wrangler dev -c apps/worker-api/wrangler.jsonc -c apps/worker-account/wrangler.jsonc
または、各ワーカーを独自のターミナルで実行し ( cd apps/worker-account && make dev 、次に cd apps/worker-api && make dev )、サービス バインディングが Wrangler 出力に接続済みとして表示されていることを確認します。
このレイアウトは、キューを使用するキュー* アプリとワーカー* アプリに使用します。

[切り捨てられた]

## Original Extract

A concise, AI native and production-grade starter template for full-stack applications using pnpm, Turborepo and Cloudflare Workers for edge-native execution 🚚⛅ - louisbrulenaudet/monorepo-template

Hello everyone, I have spent dozens of hours working on my ideal template to quickly bring applications into production, which could be of interest to those who want to rapidly deliver a scalable prototype, in an AI-native manner and at low cost on Cloudflare infrastructures. In this template, you will find: - A pnpm monorepo in strictly typed TypeScript with Zod, allowing you to develop both your React applications and your backend (here with Hono and the native Zod validator) without duplicating contracts;
- A fully designed setup to maximize the capabilities of AI agents (mainly Claude Code and Cursor for now), with precise and path-scoped rules to enforce conventions without polluting the context, hooks based on OXC for ultra-fast formatting and linting, and sub-agents enabling optimal context collection, notably thanks to Context7 and Cloudflare MCPs;
- A next-generation configuration for Vite 8, TanStack Router, TanStack Query, and React 19;
- A set of make commands enabling fast and intuitive development for both humans and AI agents;
- Clear documentation aimed at supporting the introduction of dozens of microservices, simplifying inter-application communication at zero cost thanks to the RPC protocol supported between multiple workers, etc.;
- A turborepo configuration enabling the parallel execution of all essential commands for development on a scaling codebase and the rapid deployment of dozens of services at the edge. I have already been able to battle-test this template in production for the development of MCPs with authentication, personal applications, as well as in hackathon contexts, and it notably enabled us to finish as finalists in the {Tech: Europe} Paris AI Hackathon organized at Hexa. I believe it could greatly help entrepreneurs to vibe code or prototype applications quickly while preserving certain guarantees of security and best practices thanks to the investments made in the harness, and it can appeal to a wide variety of developers concerned with minimizing friction and avoiding duplicated code between backend and frontend or across multiple services (queues, RPC, etc.). This repo is natively based on modern and relatively well-documented technologies, with the intention of offering the best possible developer experience, including for those who do not work daily with TypeScript (in practice, the granular configuration used makes it possible to adapt with minimal changes to the development of both backend microservices and React applications without friction and in a relatively secure manner). I strongly encourage you, if you are familiar with Cloudflare infrastructure and the concept of a monorepo, to take a look at it, particularly regarding the setup for AI agents, as I welcome any recommendations to further improve this tool that I genuinely enjoy working with. A big thank you to all the contributors of the various dependencies used in this template for the great experience they provide, and I look forward to seeing what you will be able to ship with this monorepo.

GitHub - louisbrulenaudet/monorepo-template: A concise, AI native and production-grade starter template for full-stack applications using pnpm, Turborepo and Cloudflare Workers for edge-native execution 🚚⛅ · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
louisbrulenaudet
/
monorepo-template
Public template
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
47 Commits 47 Commits .agents/ skills .agents/ skills .claude .claude .cursor .cursor .github .github .husky .husky .vscode .vscode apps apps assets assets hooks hooks make make packages packages .cursorignore .cursorignore .cursorindexingignore .cursorindexingignore .editorconfig .editorconfig .gitignore .gitignore .mcp.json .mcp.json .npmrc .npmrc .nvmrc .nvmrc .oxfmtrc.json .oxfmtrc.json .oxlintrc.json .oxlintrc.json .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml skills-lock.json skills-lock.json tsconfig.json tsconfig.json turbo.json turbo.json View all files Repository files navigation
Monorepo starter based on pnpm with Cloudflare, Hono, React, Vite and Tailwind 🚚⛅
A minimal, production-oriented monorepo starter built on pnpm workspaces with Turborepo, Cloudflare Workers, Hono, React (Vite), Tailwind CSS v4 , and TanStack Router/Query . AI-ready, designed for edge deployment, and structured for production projects that scale.
Starter apps today ( worker-api , front-app ). Prefixes below describe how the repo grows.
monorepo/
├── apps/ # Workers and frontends
│ ├── worker-api/ # REST API gateway
│ └── front-app/ # React SPA (Vite + TanStack)
├── packages/ # Shared @repo/* packages
│ ├── dtos-common/ # Zod wire contracts (api / rpc / queue / webhook)
│ ├── enums-common/ # Shared constrained string values (`as const`)
│ └── typescript-config/ # TypeScript configuration presets
├── make/ # Shared Makefile fragments (see make/README.md)
├── hooks/ # AI agent hooks (not Husky - see hooks/README.md)
├── package.json # Root package configuration
├── pnpm-workspace.yaml # Workspace configuration
├── turbo.json # Turborepo configuration
└── tsconfig.json # Root TypeScript configuration
Architecture Components
The monorepo is organized into two main categories: Backend Services and Frontend Applications , plus Shared Packages for common functionality.
flowchart TB
subgraph entry [Public entry]
direction LR
Front["front-* :517x"]
Ext["External providers"]
McpClients["MCP clients"]
end
subgraph publicWorkers [Public Workers]
direction LR
Gateway["worker-api :8700"]
Webhook["webhook-* :876x"]
Mcp["mcp-* :878x"]
end
subgraph privateWorkers [Private Workers]
direction LR
Biz["worker-* (RPC only)"]
Queue["queue-*"]
end
subgraph shared [Shared packages]
direction LR
Enums["@repo/enums-common"]
DTOs["@repo/dtos-common"]
Enums --> DTOs
end
Front --> Gateway
Ext --> Webhook
McpClients --> Mcp
Gateway --> Biz
Webhook --> Biz
Mcp --> Biz
Gateway --> Queue
Webhook --> Queue
Biz --> Queue
Front -.-> shared
publicWorkers -.-> shared
privateWorkers -.-> shared
Loading
Backend Services
Cloudflare Workers are organized by runtime role:
worker-api - Public HTTP gateway (Hono): CORS, validation, routing; coordinates internal Workers via RPC.
worker-* - Business logic over service-binding RPC only (no public routes in production). May own Drizzle schema under src/db/ and that database’s binding (exclusive owner).
queue-* - Queue-only consumers ( queue() handler). Messages can be produced by worker-api , worker-* , or webhook-* . Use dual-handler layout when a local HTTP debug path is useful.
webhook-* - Public HTTP ingress for external provider callbacks; forward work via RPC or queues.
mcp-* - Public HTTP MCP servers; thin tools that call worker-* over RPC.
Do not create shared packages/db-* schema packages. Put Drizzle schema under the owning app’s src/db/ and keep one DB binding owner . Other apps reach that data via service-binding RPC (or a queue) - do not attach the same DB binding to multiple apps.
front-app - React SPA (Vite 8, Tailwind v4, TanStack Router/Query) deployed on Cloudflare Workers. Talks to worker-api over HTTP only - never via service bindings.
Task
Location
New API route
apps/worker-api/src/routes/<feature>.ts → mount in src/index.ts
HTTP Zod schemas
packages/dtos-common/src/api/<feature>.ts
RPC / queue / webhook schemas
packages/dtos-common/src/<layer>/<feature>.ts (layer: rpc , queue , or webhook )
Shared string value set
packages/enums-common/src/
Worker-local value set
apps/<worker>/src/enums/
DB schema (one owner)
apps/<owner>/src/db/ - never packages/db-*
Frontend API client
apps/front-app/src/services/worker-api/<feature>.ts
Frontend page + route
apps/front-app/src/pages/ + src/routes/
Local Worker secrets
apps/<worker>/.dev.vars (from .dev.vars.example )
Frontend env
apps/front-app/.env.local (from .env.example )
Getting Started
Node.js 22+ (see root package.json engines ); we recommend fnm for version management
pnpm via the root packageManager field (Corepack recommended)
Cloudflare account only if you need make login / deploy / remote Worker features
make install
make login # optional - Cloudflare auth for remote Wrangler features
make prepare # Husky pre-commit hooks
Copy env templates before the first run:
Workers: apps/<worker>/.dev.vars.example → .dev.vars
Frontend: apps/front-app/.env.example → .env.local
Prefer make install over raw pnpm install so workspace links stay consistent.
This repo pins pnpm via packageManager in the root package.json .
First successful run (verify locally)
Start all dev servers from the repo root:
make dev
Verify the API: GET http://localhost:8700/api/v1/health
Open the frontend: http://localhost:5174
Focused work on one package: make dev SCOPE=worker-api (see Scoping ).
Command
Description
install
Install and link workspace packages
install-frozen
Install with frozen lockfile (CI)
login
Login to Cloudflare (repo-pinned Wrangler)
update
Update dependencies to latest (rewrites pnpm catalog)
check
Lint + format check (no typecheck)
ci
Lint + format + check-types (local PR gate)
deploy
Deploy all apps/workers (via Turborepo)
build
Build all packages and apps (via Turborepo)
format
Auto-fix formatting with oxfmt
lint
Auto-fix lint issues with oxlint
dev
Start all dev servers (via Turborepo)
preview
Preview production builds locally (via Turborepo)
types
Generate worker-configuration.d.ts in apps
check-types
TypeScript across all workers and packages
prepare
Install or reinstall Husky git hooks
husky-status
Show Husky hooks status
skills-update
Refresh locked agent skills (see make/README.md)
Scoping (pnpm / Turborepo)
Optional variables on any turbo-backed root target ( dev , build , ci , …):
Mnemonic: 87xx = Workers (gateway → business → queue → webhook → MCP → reserve). Frontends use Vite’s 51xx / 41xx .
Service
Path
Dev
Preview
worker-api
apps/worker-api/wrangler.jsonc
8700
-
front-app
apps/front-app/vite.config.ts
5174
4174
Notes:
Workers: set dev.port in wrangler.jsonc and monorepo.devPort in package.json . Use inspector_port: 0 .
Frontends: set Vite server.port / preview.port with strictPort: true .
Assign the next free port in the role’s range. RPC and queue-only apps still get a local port for standalone wrangler dev , but have no public URL in production.
Prefer multi-config local runs when testing bindings (first -c is HTTP-primary).
1. Create a New Cloudflare Worker
Purpose
Prefix
Example
HTTP gateway
worker-api (sticky)
worker-api
Business logic (RPC)
worker-
worker-account
Queue-only consumer
queue-
queue-email
Webhook ingress
webhook-
webhook-example
MCP server
mcp-
mcp-tools
Frontend application
front-
front-app
Key Distinctions
Gateway ( worker-api ): Public HTTP only; validates requests and calls worker-* over RPC.
Business Workers ( worker-* ): RPC-only in production ( WorkerEntrypoint ); may own Drizzle schema under src/db/ and that database’s binding (exclusive). If they also consume queues, keep this prefix and use the dual-handler layout.
Queue-only ( queue-* ): queue() consumers with no public HTTP in production; may own schema when they are the sole writer for that data.
Webhook Workers ( webhook-* ): Public HTTP for external callbacks; forward via RPC or queues.
MCP Servers ( mcp-* ): Public HTTP MCP transport; thin tools that call worker-* over RPC - never rotate long-lived credentials on this surface.
Frontends ( front-* ): React + Vite; HTTP to the gateway only - never service bindings.
Do not create shared packages/db-* schema packages.
Scaffold checklist (copy from an existing app)
There is no generator CLI. Copy the closest sibling under apps/ and wire it into the monorepo:
Copy apps/worker-api (or another closest match) to apps/<prefix-name> (e.g. apps/worker-account ).
Rename package.json name , wrangler.jsonc name , and any display strings.
Assign ports from the Development ports registry - set dev.port / inspector_port: 0 in wrangler.jsonc and monorepo.devPort in package.json .
Keep the 4-line Makefile that includes make/app.mk (see make/README.md ).
Extend @repo/typescript-config/workers.json (or the matching preset); add .dev.vars.example .
Install and typegen:
make install
make types
Copy wrangler patterns ( compatibility_flags , observability , env.staging / env.production ) from the existing app - see .cursor/rules/backend/workers-config.mdc .
Prefer scoped make from the repo root:
make dev SCOPE=worker-name
Or use the worker's Makefile:
cd apps/worker-name
make dev
This runs the dev script defined in apps/worker-name/package.json
Open the port shown in your terminal (for example, http://localhost:8721 )
Each worker Makefile exposes make dev , make format , make lint , make types , make check-types , make deploy
Testing Service Bindings Between Workers
Prefer a single multi-config wrangler dev (first -c is HTTP-primary):
wrangler dev -c apps/worker-api/wrangler.jsonc -c apps/worker-account/wrangler.jsonc
Or run each Worker in its own terminal ( cd apps/worker-account && make dev , then cd apps/worker-api && make dev ) and confirm service bindings show as connected in the wrangler output.
Use this layout for queue-* apps and for worker-* apps that also consume queue

[truncated]
