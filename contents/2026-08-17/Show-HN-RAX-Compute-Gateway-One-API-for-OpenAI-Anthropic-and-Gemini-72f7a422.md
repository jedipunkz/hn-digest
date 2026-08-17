---
source: "https://github.com/radium0090/Compute-Gateway"
hn_url: "https://news.ycombinator.com/item?id=49327446"
title: "Show HN: RAX Compute Gateway – One API for OpenAI, Anthropic, and Gemini"
article_title: "GitHub - radium0090/Compute-Gateway: The AI Compute Gateway - One API. Every AI Model. · GitHub"
image: "https://opengraph.githubassets.com/59647de737bdd2eb480b0c2a213b69e5a544acbbb1d8b9745a270deaa71b35dc/radium0090/Compute-Gateway"
author: "radium90"
captured_at: "2026-08-17T07:43:36Z"
capture_tool: "hn-digest"
hn_id: 49327446
score: 1
comments: 0
posted_at: "2026-08-17T07:21:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RAX Compute Gateway – One API for OpenAI, Anthropic, and Gemini

- HN: [49327446](https://news.ycombinator.com/item?id=49327446)
- Source: [github.com](https://github.com/radium0090/Compute-Gateway)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T07:21:37Z

## Translation

タイトル: Show HN: RAX Compute Gateway – OpenAI、Anthropic、および Gemini 用の 1 つの API
記事のタイトル: GitHub - radium0090/Compute-Gateway: AI コンピューティング ゲートウェイ - 1 つの API。すべての AI モデル。 · GitHub
説明: AI コンピューティング ゲートウェイ - 1 つの API。すべての AI モデル。 GitHub でアカウントを作成して、radium0090/Compute-Gateway の開発に貢献してください。

記事本文:
GitHub - radium0090/Compute-Gateway: AI コンピューティング ゲートウェイ - 1 つの API。すべての AI モデル。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ラジウム0090
/
コンピューティングゲートウェイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
107 コミット 107 コミット .github .github apps/ゲートウェイ apps/ゲートウェイ ベンチマーク ベンチマーク db/移行 db/移行 デプロイ デプロイ ドキュメント ドキュメントの例 例 openapi openapi パッケージ パッケージ スクリプト scri

pts sdk sdk .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .node-version .node-version .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md Dockerfile Dockerfile ライセンス ライセンス通知 README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts vitest.integration.config.ts vitest.integration.config.ts vitest.live.config.ts vitest.live.config.ts vitest.rc.config.ts vitest.rc.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
RAX コンピューティング ゲートウェイは、オープンソースの AI コンピューティング ゲートウェイです。アプリケーションに安定性を与え、
ルーティングを一元化しながら、複数のモデルプロバイダー向けの OpenAI 互換 API、
認証、再試行、制限、テレメトリ。
アプリケーション -> RAX コンピューティング ゲートウェイ -> OpenAI |人類 |ジェミニ |将来のプロバイダー
ステータス
RAX コンピューティング ゲートウェイ v0.2.0
オペレーターコンソール、ワンコマンドの自己ホスト型クイックスタート、およびパブリックを追加します。
以下に説明する悪用耐性のあるホスト型評価。リリース アーティファクトには、
署名されたマルチアーキテクチャイメージ、Helm チャート、OpenAPI コントラクト、チェックサム、SBOM、
そして来歴。このリポジトリ内の文書は、次のような規定がない限り、規範的なものであり続けます。
承認されたアーキテクチャ決定記録 (ADR) がそれらに優先します。
POST /v1/chat/completions (ストリーミングを含む)
OpenAI、Anthropic、および Gemini アダプター
安定したパブリック モデルのエイリアスと明示的なプロバイダー モデル
決定的なルーティング、フォールバック、タイムアウト、および再試行ポリシー
RAX Compute Gateway API キーと Bring-Your-own-provider-key (BYOK) 操作
ポストグレ

SQL メタデータ、オプションの Redis 調整、およびステートレス ゲートウェイ ノード
OpenTelemetry トレースとメトリクス、および構造化され編集されたログ
ローカル使用には Docker Compose、本番環境には Kubernetes マニフェスト/Helm
公的契約から生成された TypeScript および Python SDK
元の MVP の非目標には、ツールの呼び出しと構造化された出力が含まれます。
ホスト型課金システム、マーケットプレイス、GPU スケジューリング、マルチモーダル生成、
そして自律的な品質ベースのモデル選択。オペレーターコンソールとホストされた
評価は、承認された ADR によって管理される、狭い範囲の MVP 後の追加です。
api.rax-digital.com/demo を開き、次のように確認します。
GitHub に移動し、生成されたcurlをコピーします。完全なパブリックパスは
実稼働ゲートウェイに対して検証される: キーを要求し、コマンドを実行し、
正規化されたモデル応答を受け取ります。
このサービスは、5 分後に有効期限が切れる固有の API キーを発行します。あります
このリポジトリには共有公開キーがありません。トライアルは意図的に以下に限定されています
低コスト モデル 1 つ、非ストリーミング通話、入出力予算が小さい、1 つ
同時リクエスト、アカウントごとのクールダウン、グローバルな 1 日の予算。参照
ホストされたデモの設計と運用。
ホストされたサービスは評価のみを目的としています。アプリケーションは自己ホストするか、
試用版の利用可能性に依存するのではなく、通常の顧客キーを取得します。
前提条件: Docker Compose、curl、および OpenAI API キーを備えた Docker 26 以降
最初のリクエスト。フォークはコードを提供する場合にのみ必要です。試してみる
ゲートウェイで、アップストリーム リポジトリのクローンを作成し、次の 1 つのコマンドを実行します。
git clone https://github.com/radium0090/Compute-Gateway.git
cd コンピューティングゲートウェイ
sh スクリプト/quickstart.sh
スクリプトはプライベート .env を作成し、ローカル ゲートウェイ シークレットを生成し、プロンプトを表示します。
プロバイダー キーをエコーせずに取得すると、PostgreSQL、Redis、テレメトリが開始されます。
およびゲートウェイ、プロビジョニング

ローカルクライアントキーを使用し、最初のモデルを出力します
応答。構成された OpenAI 以外の場所に OpenAI キーを送信することはありません。
エンドポイント。 docker compose down でスタックを停止します。
実行が成功すると、正規化された JSON チャット応答とローカル
rcg_dev_... 資格情報が 1 回表示されます。プロバイダーが最終リクエストを拒否した場合、
ローカル スタックは診断のために実行されたままになります。プロバイダーの請求/モデルを確認する
にアクセスし、 .env を更新して、スクリプトを再実行します。
代わりに各操作を手動で理解または実行するには:
cp .env.example .env
# .env 内の偽の RCG シークレットと OPENAI_API_KEY を置き換えます。
docker 構成 --build --wait
docker compose exec postgres psql -U rcg -d compute_gateway -c \
" テナント (ID、名前、ステータス) の値 ('123e4567-e89b-42d3-a456-426614174000', 'local', 'active') に挿入します。競合時に何も行いません "
RCG_API_KEY= " $( docker compose run --rm ゲートウェイ キー 作成 \
--テナント ID 123e4567-e89b-42d3-a456-426614174000 \
--name local-app --environment dev --models ' rax/* ' --allow-streaming ) "
RCG_API_KEY をエクスポートする
テスト -n " $RCG_API_KEY "
key コマンドは新しい認証情報を 1 回発行します。コマンド置換により保持されます
ターミナルから取り出して現在のシェルに保存します。殻から出さないでください
履歴、ソース管理、ログ、URL。
Compose は PostgreSQL と Redis を提供します。外部運用、ゲートウェイの実行
RCG_REDIS_URL がない場合は、プロセスローカルの制限と回線状態を使用します。ごとに
本番レプリカには Redis の調整が必要であり、そうでない場合は起動に失敗します。
設定されています。
カール http://localhost:8080/v1/chat/completions \
-H " 認可: ベアラー $RCG_API_KEY " \
-H " Content-Type: application/json " \
-d ' {
"モデル": "rax/高速",
"messages": [{"role": "user", "content": "RAX Compute Gateway からこんにちは"}]
} '
RAX Compute Gateway は、ベース URL を変更することで OpenAI クライアントを受け入れます。
インポート

OS
openaiインポートからOpenAI
クライアント = OpenAI (
api_key = os 。環境 [ "RCG_API_KEY" ]、
Base_url = "http://localhost:8080/v1" ,
)
応答 = クライアント 。チャット 。完成品。作成(
モデル = "rax/高速" 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : "こんにちは" }],
)
オペレーターコンソール (v0.2)
RCG_ADMIN_ENABLED=true を設定し、正確な RCG_ADMIN_ORIGIN を構成して、
専用の RCG_ADMIN_SESSION_PEPPER 。移行後、最初の
標準入力で指定された一時パスワードを持つ管理者:
printf ' %s\n ' " $RCG_ADMIN_TEMPORARY_PASSWORD " | docker compose run --rm -T ゲートウェイ \
admins create --email owner@example.com --display-name ' ゲートウェイ所有者 '
http://localhost:8080/admin/ を開きます。一時パスワードは次の時点で置き換える必要があります。
最初のログイン。コンソールはテナントとワンタイム表示 API キーを管理し、
は、限定されたサービス/アクティビティのメタデータを示しています。プロバイダーの資格情報が公開されることはありません。
パスワード/セッション ハッシュ、API キー ハッシュ、プロンプト、または入力完了。
実行可能なcurl、Node.js、Pythonのサンプルが含まれています
ローカルで実行されているゲートウェイの場合。
実装ステータスと既知のギャップ
キーワード MUST 、 MUST NOT 、 SHOULD 、および MAY は定義どおりに使用されます。
RFC 2119 による。実装作業は次の優先順位に従う必要があります。
アーキテクチャと運用のガイダンス
ドキュメントが競合する場合は、マージする前に問題をオープンし、ADR で解決してください。
公共契約を変更する行為。
バグや限定された機能については GitHub の問題を使用し、質問や内容についてはディスカッションを使用してください。
設計の調査、およびレビューされた変更のプル リクエスト。読んでください
コードを送信する前に貢献してください。
ゲートウェイを実行するためだけにフォークする必要はありません。貢献するには、まず
リポジトリをフォークし、クローンを作成します
の説明に従って、フォークを作成し、機能ブランチを作成し、プル リクエストを開きます。
貢献ガイド。
プロジェクトはライセンスを取得しています

Apache License 2.0に基づいて。参照
依存関係と貢献のルールについては、licensing.md。
AI コンピューティング ゲートウェイ - 1 つの API。すべての AI モデル。
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The AI Compute Gateway - One API. Every AI Model. Contribute to radium0090/Compute-Gateway development by creating an account on GitHub.

GitHub - radium0090/Compute-Gateway: The AI Compute Gateway - One API. Every AI Model. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
radium0090
/
Compute-Gateway
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
107 Commits 107 Commits .github .github apps/ gateway apps/ gateway benchmarks benchmarks db/ migrations db/ migrations deploy deploy docs docs examples examples openapi openapi packages packages scripts scripts sdk sdk .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .node-version .node-version .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts vitest.integration.config.ts vitest.integration.config.ts vitest.live.config.ts vitest.live.config.ts vitest.rc.config.ts vitest.rc.config.ts View all files Repository files navigation
RAX Compute Gateway is an open-source AI Compute Gateway. It gives applications one stable,
OpenAI-compatible API for multiple model providers while centralizing routing,
authentication, retries, limits, and telemetry.
Application -> RAX Compute Gateway -> OpenAI | Anthropic | Gemini | future providers
Status
RAX Compute Gateway v0.2.0
adds the operator console, one-command self-hosted quickstart, and the public,
abuse-resistant hosted evaluation described below. Release artifacts include a
signed multi-architecture image, Helm chart, OpenAPI contract, checksums, SBOM,
and provenance. The documents in this repository remain normative unless an
accepted Architecture Decision Record (ADR) supersedes them.
POST /v1/chat/completions , including streaming
OpenAI, Anthropic, and Gemini adapters
Stable public model aliases and explicit provider models
Deterministic routing, fallback, timeout, and retry policies
RAX Compute Gateway API keys and bring-your-own-provider-key (BYOK) operation
PostgreSQL metadata, optional Redis coordination, and stateless gateway nodes
OpenTelemetry traces and metrics plus structured, redacted logs
Docker Compose for local use and Kubernetes manifests/Helm for production
TypeScript and Python SDKs generated from the public contract
Non-goals for the original MVP include tool calling and structured outputs, a
hosted billing system, a marketplace, GPU scheduling, multimodal generation,
and autonomous quality-based model selection. The operator console and hosted
evaluation are narrowly scoped post-MVP additions governed by accepted ADRs.
Open api.rax-digital.com/demo , verify with
GitHub, and copy the generated curl . The complete public path has been
verified against the production gateway: claim a key, run the command, and
receive a normalized model response.
The service issues a unique API key that expires after five minutes; there is
no shared public key in this repository. The trial is intentionally limited to
one low-cost model, non-streaming calls, small input/output budgets, one
concurrent request, per-account cooldown, and a global daily budget. See
hosted demo design and operation .
The hosted service is for evaluation only. Applications should self-host or
obtain a normal customer key rather than depend on trial availability.
Prerequisites: Docker 26+ with Docker Compose, curl , and an OpenAI API key for
the first request. A fork is needed only when contributing code; to try the
gateway, clone the upstream repository and run one command:
git clone https://github.com/radium0090/Compute-Gateway.git
cd Compute-Gateway
sh scripts/quickstart.sh
The script creates a private .env , generates local gateway secrets, prompts
for the provider key without echoing it, starts PostgreSQL, Redis, telemetry,
and the gateway, provisions a local client key, and prints the first model
response. It never sends the OpenAI key anywhere except the configured OpenAI
endpoint. Stop the stack with docker compose down .
A successful run ends with a normalized JSON chat response and a local
rcg_dev_... credential shown once. If the provider rejects the final request,
the local stack remains running for diagnosis; check provider billing/model
access, update .env , and rerun the script.
To understand or run each operation manually instead:
cp .env.example .env
# Replace the fake RCG secrets and OPENAI_API_KEY in .env.
docker compose up --build --wait
docker compose exec postgres psql -U rcg -d compute_gateway -c \
" INSERT INTO tenants (id, name, status) VALUES ('123e4567-e89b-42d3-a456-426614174000', 'local', 'active') ON CONFLICT DO NOTHING "
RCG_API_KEY= " $( docker compose run --rm gateway keys create \
--tenant-id 123e4567-e89b-42d3-a456-426614174000 \
--name local-app --environment dev --models ' rax/* ' --allow-streaming ) "
export RCG_API_KEY
test -n " $RCG_API_KEY "
The key command emits the new credential once; command substitution keeps it
out of the terminal and stores it in the current shell. Keep it out of shell
history, source control, logs, and URLs.
Compose supplies PostgreSQL and Redis. Outside production, running the gateway
without RCG_REDIS_URL uses process-local limits and circuit state; every
production replica requires Redis coordination and fails startup if it is not
configured.
curl http://localhost:8080/v1/chat/completions \
-H " Authorization: Bearer $RCG_API_KEY " \
-H " Content-Type: application/json " \
-d ' {
"model": "rax/fast",
"messages": [{"role": "user", "content": "Hello from RAX Compute Gateway"}]
} '
RAX Compute Gateway accepts the OpenAI client by changing its base URL:
import os
from openai import OpenAI
client = OpenAI (
api_key = os . environ [ "RCG_API_KEY" ],
base_url = "http://localhost:8080/v1" ,
)
response = client . chat . completions . create (
model = "rax/fast" ,
messages = [{ "role" : "user" , "content" : "Hello" }],
)
Operator console ( v0.2 )
Set RCG_ADMIN_ENABLED=true , configure the exact RCG_ADMIN_ORIGIN , and use a
dedicated RCG_ADMIN_SESSION_PEPPER . After migrations, create the first
administrator with a temporary password supplied on standard input:
printf ' %s\n ' " $RCG_ADMIN_TEMPORARY_PASSWORD " | docker compose run --rm -T gateway \
admins create --email owner@example.com --display-name ' Gateway Owner '
Open http://localhost:8080/admin/ . The temporary password must be replaced at
first login. The console manages tenants and one-time-display API keys and
shows bounded service/activity metadata. It never exposes provider credentials,
password/session hashes, API-key hashes, prompts, or completions.
Runnable curl, Node.js, and Python examples are included
for a locally running gateway.
Implementation status and known gaps
The keywords MUST , MUST NOT , SHOULD , and MAY are used as defined
by RFC 2119. Implementation work should follow this precedence order:
Architecture and operational guidance
If documents conflict, open an issue and resolve it with an ADR before merging
behavior that changes the public contract.
Use GitHub Issues for bugs and scoped features, Discussions for questions and
design exploration, and pull requests for reviewed changes. Please read
CONTRIBUTING before submitting code.
You do not need to fork merely to run the gateway. To contribute, first
fork the repository , clone
your fork, create a feature branch, and open a pull request as described in the
contribution guide.
The project is licensed under Apache License 2.0. See
licensing.md for dependency and contribution rules.
The AI Compute Gateway - One API. Every AI Model.
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
