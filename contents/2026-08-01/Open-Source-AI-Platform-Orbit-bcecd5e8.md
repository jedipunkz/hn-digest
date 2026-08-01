---
source: "https://github.com/schmitech/orbit"
hn_url: "https://news.ycombinator.com/item?id=49134457"
title: "Open-Source AI Platform Orbit"
article_title: "GitHub - schmitech/orbit: Self-hosted, OpenAI-compatible AI gateway for private RAG, natural-language data access, and tool-calling agents. · GitHub"
author: "r_martinez"
captured_at: "2026-08-01T14:25:36Z"
capture_tool: "hn-digest"
hn_id: 49134457
score: 2
comments: 0
posted_at: "2026-08-01T13:51:40Z"
tags:
  - hacker-news
  - translated
---

# Open-Source AI Platform Orbit

- HN: [49134457](https://news.ycombinator.com/item?id=49134457)
- Source: [github.com](https://github.com/schmitech/orbit)
- Score: 2
- Comments: 0
- Posted: 2026-08-01T13:51:40Z

## Translation

タイトル: オープンソース AI プラットフォーム Orbit
記事のタイトル: GitHub - schmitech/orbit: プライベート RAG、自然言語データ アクセス、およびツール呼び出しエージェント用のセルフホスト型の OpenAI 互換 AI ゲートウェイ。 · GitHub
説明: プライベート RAG、自然言語データ アクセス、およびツール呼び出しエージェント用の自己ホスト型の OpenAI 互換 AI ゲートウェイ。 - シュミテック/オービット

記事本文:
GitHub - schmitech/orbit: プライベート RAG、自然言語データ アクセス、およびツール呼び出しエージェント用の自己ホスト型の OpenAI 互換 AI ゲートウェイ。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
アラートを閉じる
{{ メッセージ }}
シュミテック
/
軌道
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2,144 コミット 2,144 コミット bin bin クライアント クライアント config config docker docker docs docs 例 例 インストール インストール サーバー サーバー utils utils .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md conformance-baseline.yml conformance-baseline.yml env.example env.example ruff.toml ruff.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つの OpenAI 互換 API を通じてプライベート データと内部ツールを接続する
ファイル、データベース、ベクター ストア、モデル、API、MCP ツールを接続します。認証、可観測性、ガバナンスが組み込まれており、ローカルまたはクラウドで実行できます。
⚡ ライブサンドボックスを試してみる
•
クイックスタート
•
デモを見る
•
チュートリアル
•
ドキュメント
マルチモーダル.mp4
PDF、スプレッドシート、画像をアップロードし、会話全体で保存されたコンテキストとともにそれらをクエリします。
👉 ブラウザでライブで試してみる →
⭐ ORBIT のクローンを作成しますか?役立つと思われる場合は、リポジトリにスターを付けます。これは、他の開発者がプロ​​ジェクトを発見し、新しいモデル、データソース、エージェントの統合への投資を継続する必要があることを示すのに役立ちます。
プライベート AI アプリケーション用の 1 つのバックエンド
ORBIT が提供するもの
何でもつなげる
ファイル、SQL、NoSQL、ベクター ストア、Elasticsearch、REST/GraphQL API、MCP ツールを複数の言語で自然言語でクエリします。
どのモデルでも使用可能
Ollama、llama.cpp、vLLM などのローカル モデル、または OpenAI、Anthropic、Gemini、Bedrock、Azure などのクラウド プロバイダーにわたって 1 つの API コントラクトをルーティングします。
安全に操作してください
API キー、RBAC、SSO、クォータ、モデレーション、

フォールバック、メトリクス、監査ログ、管理パネルを自分で組み立てるのではなく、
ORBIT は、アプリケーションと、アプリケーションに必要なモデル、データ、ツールの間に位置します。 YAML でアダプターを定義し、1 つの OpenAI 互換エンドポイントを通じて公開し、アーキテクチャを置き換えることなく、ローカル プロトタイプから管理されたデプロイメントに移行します。
どこに当てはまりますか? ORBIT は、AI ゲートウェイと検索およびツールの実行を組み合わせます。これは単なるチャット UI ではなくバックエンド API であり、アプリケーション コードに任せるのではなく運用コントロールが含まれています。 「ORBIT 対 Open WebUI」および「ORBIT 対 LiteLLM」を参照してください。
ORBIT は積極的に維持されます。リリース履歴、変更履歴、およびコミット履歴を参照してください。
モデル選択.mp4
config/adapters/multimodal.yaml で定義されているアクティブなアダプター ポリシーを使用して、会話中に許可されたモデルに切り替えます。
🚀 クイックスタート
インストールする前に、ORBIT で何ができるかを確認してください。ライブ サンドボックスを即座に探索できます。ダウンロード、Docker、セットアップは必要ありません。
それ以外の場合は、クローンと構成ファイルの編集をスキップし、フレーバー イメージをプルして実行します。 ORBIT、orbitchat Web UI、最小限のドキュメント チャット セットアップがすべて内蔵されており、数分で使い始めることができます。
前提条件: Docker、4 GB の空き RAM、および 3 GB のディスク容量。
docker pull schmitech/orbit-ollama:latest
docker run -d --name orbit -p 5173:5173 -p 3000:3000 \
-v 軌道データ:/軌道/データ \
-v orbit-models:/orbit/models \
シュミテック/オービット-オラマ:最新
最初の実行では、コンテナー内のローカル チャット/ビジョン モデル ( gemma4:e2b 、約 7.2 GB) がダウンロードされ、インターネット接続速度に応じて起動が完了するまでに時間がかかります。取得したら、http://localhost:5173 を開いてチャットを開始します。PDF、スプレッドシート、または画像をアップロードして、それについて質問します。クラウド アカウントや API キーは必要ありません。
エクスポート OPENAI_API_KEY=sk-...
港湾労働者

pull schmitech/orbit-openai:最新
docker run -d --name orbit -p 5173:5173 -p 3000:3000 \
-e OPENAI_API_KEY \
-v 軌道データ:/軌道/データ \
シュミテック/オービット-オープンナイ:最新
モデル
チャット
gpt-5.4-mini (選択可能: gpt-5.4 、 gpt-5.4-nano )
ビジョン
gpt-5.5
埋め込み
テキスト埋め込み-3-small
オプション 3: Gemini ホスト型モデル
エクスポート GOOGLE_API_KEY=...
docker pull schmitech/orbit-gemini:latest
docker run -d --name orbit -p 5173:5173 -p 3000:3000 \
-e GOOGLE_API_KEY \
-v 軌道データ:/軌道/データ \
シュミテック/オービットジェミニ:最新
モデル
チャット
gemini-3.1-pro-preview (gemini-3.6-flash も選択可能)
ビジョン
ジェミニ-3.6-フラッシュ
埋め込み
gemini-embedding-2-プレビュー
ヒント
これらは単なるデフォルトです。管理パネルのアダプター設定からアダプターごとのアクティブなモデルを変更します (再起動後も保持され、orbit-data ボリュームに保存されます)。または、実行中のコンテナー ( docker exec -it orbit sh ) 内の /orbit/config-runtime/adapters/multimodal.yaml で解決された YAML を直接検査/編集します。コンテナーの再起動によりこのファイルがイメージのデフォルトから再生成されるため、そこでの編集は再起動後はそのままでは存続しないことに注意してください。
-e OPENAI_API_KEY (no =value ) は、その変数がシェル内ですでに設定されているものをすべて通過します。最初にエクスポートしてください。キーを -e OPENAI_API_KEY=sk-... としてインラインに貼り付けないでください。これにより、キーがシェル履歴に残ることになります。各クラウド フレーバーには 1 つの認証情報が必要です。同じキーがチャット、ビジョン、埋め込みを強化するため、別のプロバイダーに暗黙的にフォールバックすることはありません。 docker pull は資格情報を必要としたり、受け取ったり、保持したりすることはありません。 docker run のみが実行します。
ポート 5173 はチャット UI で、ORBIT を直接呼び出す場合は 3000 が OpenAI 互換 API です。
curl -X POST http://localhost:3000/v1/chat/completions \
-H ' コンテンツ タイプ: application/json ' \
-H ' X-API キー: マルチモーダル ' \
-H

' X セッション ID: ローカルテスト ' \
-d ' {"messages":[{"role":"user","content":"ORBIT は何に接続できますか?"}]} '
http://localhost:3000/admin で管理パネルにアクセスすることもできます (デフォルトの資格情報: ユーザー名 admin 、パスワード admin123 — 管理パネル内で変更できます)。
これらのイメージには、初回実行時に便利なように、デフォルトのデータベースと API キーが付属しています。 ORBIT をローカルホストを超えて公開する前に、デフォルトの API キー/管理者パスワードをローテーションします。
実稼働デプロイメントの場合は、常に最新の安定リリースを使用してください。
Docker ガイド、チュートリアル、または Windows ガイドに従うこともできます。
目標
オービットハンドル
プライベートドキュメントを使ってチャットする
PDF、オフィス文書、スプレッドシート、画像、音声をアップロードします。会話全体で関連するコンテキストを取得します。チュートリアルを試してみる →
複数の言語でデータベースにクエリを実行する
SQL、MongoDB、Elasticsearch、および複合データソースにわたって安全なクエリを生成および実行します。 SQL デモを試してみる →
ツールを使用したエージェントの構築
限定された複数ステップのサーバー側ツール ループを使用して、モデルに MCP サーバーへの限定されたアクセスを提供します。 MCP ガイドを読む→
1 つの管理された AI エンドポイントを提供する
キーごとのアクセス、クォータ、フォールバック、モデレーション、メトリクス、監査機能を備えたローカル モデルとクラウド モデルをルーティングします。最初のキーを作成します →
データベースに複数の言語で質問する
db-クエリ.mp4
ORBIT はクエリを生成し、データベースに対して実行し、チャットで結果をグラフ化します。
👉 SQL Database クエリのデモをライブで試してみる →
dbチャット.mp4
SQL データベース、API、またはデータ レイクに基づいたスピーチツースピーチ音声 - 応答中に中断すると、停止してすぐに応答します。
mcp-tool-demo.mp4
エージェント フレームワークを追加せずに、ファイル システム、Slack、Postgres、GitHub、Jira、およびその他の MCP サーバーに接続します。
👉 MCP ツール呼び出しデモをライブで試してみる →
mcp-ビジネス.mp4
9 つの合成 MCP にわたるマルチステップのエージェント推論も可能

ls: CRM の健全性、テレメトリ シートの使用率、P1 サポートのエスカレーション、チャーン リスクのシミュレーション。
👉 Business & Revenue Intelligence MCP デモをライブで試してみる →
管理パネル.mp4
API キー、クォータ、レート制限の背後にある健全性、レイテンシー、コスト、トークン、セッション、アダプター、およびログを監視します。
必要な場合は…
ORBIT が提供するのは…
モデルルーティング以上のもの
RAG、構造化データの取得、Web 検索、ゲートウェイの背後でのツールの実行。
チャットインターフェイス以上のもの
ORBIT Chat または OpenAI 互換 API を呼び出すことができるクライアントと連携するバックエンド。
プロトタイプのフレームワークを超えたもの
認証、RBAC、SSO、クォータ、モデレーション、サーキット ブレーカー、フォールバック、メトリクス、監査ログ。
プライベート展開
ローカル推論、暗号化されたファイル ストレージ、クラウド シークレット マネージャー、および完全なオフライン操作。
オーケストレーション コードの削減
YAML 定義のアダプター、データソース、プロンプト、プロバイダー ルーティング、およびガードレール。
能力
能力
付属
モデルゲートウェイ
37 を超えるローカルおよびクラウド プロバイダー、OpenAI 互換 API、キーごとのルーティング、モデルの切り替え、再試行、フォールバック。
検索
Vector RAG、ファイルおよびマルチモーダル RAG、SQL、MongoDB、Elasticsearch、REST、GraphQL、Web 検索、およびマルチソースの回答。
エージェントとプロトコル
MCP ツール呼び出し、境界付きマルチステップ ループ、自然言語スキル ルーティング、A2A、および非同期 RabbitMQ リクエスト。
メディア
画像、ビデオ、音声、PDF、Word、Excel、PowerPoint、CSV、マークダウンの生成。
セキュリティ
API キー、RBAC、Entra ID および Auth0 SSO、レート制限、クォータ、モデレーション、AES-256-GCM ファイル暗号化、およびクラウド シークレット マネージャー。
運営
管理 UI、ヘルスチェック、メトリクス、監査ログ、リクエストごとのトークンと推定コストの追跡、支出分析、サーキット ブレーカー、データソース プーリング、ホット アダプターのリロード。
すべてのアダプターを参照する · プロバイダー構成を参照する · 構成を読む

参照
REST、OpenAI 互換、MCP、A2A、またはメッセージ キュー リクエストを認証し、モデル、プライベート データ、ツールにルーティングします。
ここから始めましょう
リソース
オービットを学ぶ
チュートリアル · 初めてのチャット · HTTP API
アダプターの構成
アダプターの概要 · 構成ガイド
プライベートデータを接続する
ファイル · ベクター ストア · SQL
ビルドエージェント
MCPツール・自動スキルルーティング・A2A
本番環境で実行する
認証 · 使用量とコストの追跡 · レート制限 · フォールト トレランス
クライアントを使用する
ORBIT チャット · Node.js SDK · API キーと Python の例
貢献する
新しい取得者とプロバイダーの統合、展開ガイド、テスト、修正、ドキュメントなどの貢献を歓迎します。 CONTRIBUTING.md を読み、未解決の問題を選択するか、ディスカッションを開始します。
Remsy Schmilinsky によって保守されています。
ORBIT は、Apache License 2.0 に基づいてライセンスされています。
プライベート RAG、自然言語データ アクセス、ツール呼び出しエージェント用の自己ホスト型の OpenAI 互換 AI ゲートウェイ。
github.com/schmitech/orbit トピック
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
56 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted, OpenAI-compatible AI gateway for private RAG, natural-language data access, and tool-calling agents. - schmitech/orbit

GitHub - schmitech/orbit: Self-hosted, OpenAI-compatible AI gateway for private RAG, natural-language data access, and tool-calling agents. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
schmitech
/
orbit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2,144 Commits 2,144 Commits bin bin clients clients config config docker docker docs docs examples examples install install server server utils utils .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md conformance-baseline.yml conformance-baseline.yml env.example env.example ruff.toml ruff.toml View all files Repository files navigation
Connect private data and internal tools through one OpenAI-compatible API
Connect files, databases, vector stores, models, APIs, and MCP tools. Run locally or in your cloud—with authentication, observability, and governance built in.
⚡ Try Live Sandbox
•
Quick start
•
Watch the demo
•
Tutorial
•
Documentation
multimodal.mp4
Upload PDFs, spreadsheets, and images, then query them together with context preserved across the conversation.
👉 Try this live in your browser →
⭐ Cloning ORBIT? If it looks useful, star the repository . It helps other developers discover the project and signals that we should keep investing in new model, datasource, and agent integrations.
One backend for private AI applications
What ORBIT gives you
Connect anything
Query files, SQL, NoSQL, vector stores, Elasticsearch, REST/GraphQL APIs, and MCP tools in natural language across multiple languages.
Use any model
Route one API contract across local models such as Ollama, llama.cpp, and vLLM or cloud providers such as OpenAI, Anthropic, Gemini, Bedrock, and Azure.
Operate it safely
Ship with API keys, RBAC, SSO, quotas, moderation, fallbacks, metrics, audit logs, and an admin panel instead of assembling them yourself.
ORBIT sits between your applications and the models, data, and tools they need. Define adapters in YAML, expose them through one OpenAI-compatible endpoint, and move from a local prototype to a governed deployment without replacing the architecture.
Where does it fit? ORBIT combines an AI gateway with retrieval and tool execution. It is a backend API rather than just a chat UI, and it includes production controls rather than leaving them to application code. See ORBIT vs. Open WebUI and ORBIT vs. LiteLLM .
ORBIT is actively maintained. See the release history , changelog , and commit history .
model-choosing.mp4
Switch to any allowed model mid-conversation, with the active adapter policy defined in config/adapters/multimodal.yaml .
🚀 Quick Start
See what ORBIT can do before you install: explore the live sandbox instantly—no download, Docker, or setup required.
Otherwise, skip the clone and config-file editing — pull a flavor image and run it. ORBIT, the orbitchat web UI, and a minimalistic document-chat setup are all inside to get you started in minutes.
Prerequisites: Docker, 4 GB of free RAM, and 3 GB of disk space.
docker pull schmitech/orbit-ollama:latest
docker run -d --name orbit -p 5173:5173 -p 3000:3000 \
-v orbit-data:/orbit/data \
-v orbit-models:/orbit/models \
schmitech/orbit-ollama:latest
The first run downloads the local chat/vision model ( gemma4:e2b , ~7.2 GB) inside the container and will take some time to complete startup depending on your internet connection speed. Once pulled, open http://localhost:5173 and start chatting — upload a PDF, a spreadsheet, or an image and ask about it. No cloud account or API key required.
export OPENAI_API_KEY=sk-...
docker pull schmitech/orbit-openai:latest
docker run -d --name orbit -p 5173:5173 -p 3000:3000 \
-e OPENAI_API_KEY \
-v orbit-data:/orbit/data \
schmitech/orbit-openai:latest
Model
Chat
gpt-5.4-mini (also selectable: gpt-5.4 , gpt-5.4-nano )
Vision
gpt-5.5
Embeddings
text-embedding-3-small
Option 3: Gemini Hosted Model
export GOOGLE_API_KEY=...
docker pull schmitech/orbit-gemini:latest
docker run -d --name orbit -p 5173:5173 -p 3000:3000 \
-e GOOGLE_API_KEY \
-v orbit-data:/orbit/data \
schmitech/orbit-gemini:latest
Model
Chat
gemini-3.1-pro-preview (also selectable: gemini-3.6-flash )
Vision
gemini-3.6-flash
Embeddings
gemini-embedding-2-preview
Tip
These are just the defaults. Change the active model per adapter from the Admin Panel's adapter settings (persists across restarts, stored in the orbit-data volume), or inspect/edit the resolved YAML directly at /orbit/config-runtime/adapters/multimodal.yaml inside the running container ( docker exec -it orbit sh ) — note that a container restart regenerates this file from the image default, so edits there don't survive a restart on their own.
-e OPENAI_API_KEY (no =value ) passes through whatever that variable is already set to in your shell — export it first, don't paste the key inline as -e OPENAI_API_KEY=sk-... , which would leave it sitting in your shell history. Each cloud flavor needs exactly one credential — the same key powers chat, vision, and embeddings, so nothing silently falls back to a different provider. docker pull never needs, receives, or persists a credential; only docker run does.
Port 5173 is the chat UI, 3000 is the OpenAI-compatible API if you want to call ORBIT directly:
curl -X POST http://localhost:3000/v1/chat/completions \
-H ' Content-Type: application/json ' \
-H ' X-API-Key: multimodal ' \
-H ' X-Session-ID: local-test ' \
-d ' {"messages":[{"role":"user","content":"What can ORBIT connect to?"}]} '
You can also access the Admin Panel at http://localhost:3000/admin (default credentials: username admin , password admin123 — which can be changed inside the admin panel).
These images ship with a default database and API key for first-run convenience. Rotate the default API key/admin password before exposing ORBIT beyond localhost.
For production deployments, ALWAYS use the latest stable release .
You can also follow the Docker guide , tutorial , or Windows guide .
Goal
ORBIT handles
Chat with private documents
Upload PDFs, office documents, spreadsheets, images, and audio; retrieve relevant context across a conversation. Try the tutorial →
Query databases in multiple languages
Generate and execute safe queries across SQL, MongoDB, Elasticsearch, and composite datasources. Try the SQL demo →
Build tool-using agents
Give models scoped access to MCP servers with bounded, multi-step server-side tool loops. Read the MCP guide →
Offer one governed AI endpoint
Route local and cloud models with per-key access, quotas, fallbacks, moderation, metrics, and auditability. Create your first key →
Ask database questions in multiple languages
db-query.mp4
ORBIT generates the query, runs it against the database, and charts the result in chat.
👉 Try the SQL Database query demo live →
dbchat.mp4
Speech-to-speech voice grounded in SQL databases, APIs, or data lakes — interrupt it mid-answer and it stops and responds immediately.
mcp-tool-demo.mp4
Connect filesystem, Slack, Postgres, GitHub, Jira, and other MCP servers without adding an agent framework.
👉 Try the MCP tool calling demo live →
mcp-business.mp4
Multi-step agent reasoning across 9 synthetic MCP tools: CRM health, telemetry seat utilization, P1 support escalations, and churn risk simulation.
👉 Try the Business & Revenue Intelligence MCP demo live →
admin-panel.mp4
Monitor health, latency, costs, tokens, sessions, adapters, and logs behind API keys, quotas, and rate limits.
If you need…
ORBIT gives you…
More than model routing
RAG, structured-data retrieval, web search, and tool execution behind the gateway.
More than a chat interface
A backend that works with ORBIT Chat or any client that can call an OpenAI-compatible API.
More than a prototype framework
Authentication, RBAC, SSO, quotas, moderation, circuit breakers, fallbacks, metrics, and audit logs.
Private deployment
Local inference, encrypted file storage, cloud secret managers, and fully offline operation.
Less orchestration code
YAML-defined adapters, datasources, prompts, provider routing, and guardrails.
Capabilities
Capability
Included
Model gateway
37+ local and cloud providers, OpenAI-compatible APIs, per-key routing, model switching, retries, and fallbacks.
Retrieval
Vector RAG, file and multimodal RAG, SQL, MongoDB, Elasticsearch, REST, GraphQL, web search, and multi-source answers.
Agents and protocols
MCP tool calling, bounded multi-step loops, natural-language skill routing, A2A, and asynchronous RabbitMQ requests.
Media
Image, video, speech, PDF, Word, Excel, PowerPoint, CSV, and markdown generation.
Security
API keys, RBAC, Entra ID and Auth0 SSO, rate limits, quotas, moderation, AES-256-GCM file encryption, and cloud secret managers.
Operations
Admin UI, health checks, metrics, audit logs, per-request token and estimated-cost tracking, spend analytics, circuit breakers, datasource pooling, and hot adapter reloads.
Browse all adapters · See provider configuration · Read the configuration reference
Authenticate and route REST, OpenAI-compatible, MCP, A2A, or message-queue requests to models, private data, and tools.
Start here
Resource
Learn ORBIT
Tutorial · Your first chat · HTTP APIs
Configure adapters
Adapter overview · Configuration guide
Connect private data
Files · Vector stores · SQL
Build agents
MCP tools · Auto skill routing · A2A
Run in production
Authentication · Usage and cost tracking · Rate limiting · Fault tolerance
Use a client
ORBIT Chat · Node.js SDK · API key and Python examples
Contributing
Contributions are welcome: new retrievers and provider integrations, deployment guides, tests, fixes, and documentation. Read CONTRIBUTING.md , pick an open issue , or start a discussion.
Maintained by Remsy Schmilinsky .
ORBIT is licensed under the Apache License 2.0 .
Self-hosted, OpenAI-compatible AI gateway for private RAG, natural-language data access, and tool-calling agents.
github.com/schmitech/orbit Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
56 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
