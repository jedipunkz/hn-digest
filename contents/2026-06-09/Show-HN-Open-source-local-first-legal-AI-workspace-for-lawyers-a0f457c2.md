---
source: "https://github.com/rohasnagpal/AI-Blueprint"
hn_url: "https://news.ycombinator.com/item?id=48457883"
title: "Show HN: Open-source, local-first legal AI workspace for lawyers"
article_title: "GitHub - rohasnagpal/AI-Blueprint: Open source AI-native infrastructure for law firms. · GitHub"
author: "rohasnagpal"
captured_at: "2026-06-09T10:19:25Z"
capture_tool: "hn-digest"
hn_id: 48457883
score: 2
comments: 0
posted_at: "2026-06-09T07:39:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source, local-first legal AI workspace for lawyers

- HN: [48457883](https://news.ycombinator.com/item?id=48457883)
- Source: [github.com](https://github.com/rohasnagpal/AI-Blueprint)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T07:39:33Z

## Translation

タイトル: Show HN: 弁護士向けのオープンソース、ローカルファーストの法律 AI ワークスペース
記事のタイトル: GitHub - rohasnagpal/AI-Blueprint: 法律事務所向けのオープンソース AI ネイティブ インフラストラクチャ。 · GitHub
説明: 法律事務所向けのオープンソース AI ネイティブ インフラストラクチャ。 - ロハスナグパル/AI ブループリント

記事本文:
GitHub - rohasnagpal/AI-Blueprint: 法律事務所向けのオープンソース AI ネイティブ インフラストラクチャ。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ロハスナグパル
/
AI ブループリント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインBr

anches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
101 コミット 101 コミット .github .github アプリ アプリ ドキュメント ドキュメント マイグレーション マイグレーション パブリック パブリック ルート ルート スクリプト スクリプト テスト テスト .dockerignore .dockerignore .env.production.example .env.production.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE OPERATIONS.md OPERATIONS.md README.md README.md SECURITY.md SECURITY.md ai-blueprint-multiuser-plugin-plan.html ai-blueprint-multiuser-plugin-plan.html alembic.ini alembic.ini database.py database.py docker-compose.example.yml docker-compose.example.yml main.py main.py multi-agent-contract-review-plan.html multi-agent-contract-review-plan.html 要件-local.txt 要件-ローカル.txt 要件.txt 要件.txt webtools.py webtools.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
文書ベースのチャット、ライブ音声支援、法的草案、契約レビュー、準備ワークフロー、翻訳、電子メール草案、および案件文書管理のための、ローカルファーストの法務 AI ワークスペースです。
AI ブループリントは、機密文書、問題のコンテキスト、情報源の根拠付け、レビュー ワークフロー、監査可能性の処理に AI を必要とする弁護士、会社、法務チーム、法務運用チーム向けに構築されています。
目標は法的判断に代わることではありません。目標は、データ、モデル、検索範囲をより強力に制御して、文書のレビュー、紛争と交渉の準備、作業成果物の草案、法的資料の翻訳、問題知識の整理のための構造化されたワークスペースを弁護士に提供することです。
文書に基づいた法的チャットと RAG
スコープ指定された ret を使用して、アップロードされたファイル、案件文書、契約書、電子メール、判例、ポリシー、法令、規制、承認された情報源全体にわたって質問します。

ライバル。
法的草案作成ワークスペース
事実、当事者、条項、指示、管轄権、論調、およびオプションのソース文書から、法的通知、合意、返答、取締役会文書、条項、顧客向け草案、その他の法的作業成果物を生成します。ドラフトには、印刷可能な HTML、プレーン テキスト、前提条件、欠落情報、レビュー警告、ソースの使用状況、進行状況イベント、および保存されたドラフト履歴が含まれます。
契約レビューのワークフロー
インデックス付きの契約書をワークスペースと案件から直接確認します。レビューの深さ、プレイブック、選択したソース契約、およびレビュー手順を使用して、条項分析、リスク、レッドライン提案、要約、警告、ソース参照、およびダウンロード可能なマークダウン出力を生成します。
準備ワークフロー
インデックス付き案件文書から構造化された仲裁、訴訟、調停、交渉の成果物を作成します。準備の実行では、事件のスナップショット、問題マップ、年表、証拠マトリックス、手順タスク、証人または参加者の準備、損害と救済の分析、リスク、ギャップ、追跡データ、警告、ソースベースを作成できます。
法的調査支援
チャット、ドキュメント モード、ドラフト、および準備のワークフローを使用して、法的問題を調査し、アップロードされた典拠を分析し、調査メモを作成し、議論をテストします。現在のプライマリ ナビゲーションでは、個別の法律調査画面が表示されません。
電子メールとクライアントとのコミュニケーションの下書き
IMAP 受信ボックスをポーリングし、ペルソナとドキュメント コンテキストを使用して AI 支援による返信を生成し、下書きを確認して、承認された返信を SMTP 経由で送信します。
法律文書およびビジネス文書の翻訳
貼り付けられたテキストまたはアップロードされた文書を法律、ビジネス、技術、リテラル、および平文モードで翻訳し、警告、翻訳者メモ、保存された用語、およびダウンロード可能な HTML 出力を確認します。
生の音声アシスタント
OpenAI リアルタイム音声を音声による法的支援に使用する

ランス、案件のウォークスルー、ドキュメントの質問、およびドキュメント モードでのドキュメント検索を含むワークフロー ガイダンス。
ワークスペース、案件、ユーザー、権限
ドキュメント、実行、メンバー、ロール、出力、アクセスをワークスペースと案件ごとに整理します。
現在のナビゲーション
新しいチャット、ドキュメントの追加、ドキュメントの表示、準備、ワークフロー、ペルソナ、設定から作業します。準備メニューには、仲裁準備、訴訟準備、調停準備、交渉準備が含まれます。ワークフロー メニューには、契約レビュー、草案、電子メール、翻訳が含まれます。
文書管理
ファイルのアップロード、URL の取り込み、ローカル フォルダーの接続、フォルダー ソースの同期、ドキュメントのインデックス作成、ドキュメント ライブラリの検索、削除の管理を行います。
ペルソナと役割固有の行動
契約レビュー担当者、訴訟アソシエイト、法律調査担当者、パートナーレビュー担当者、平易な英語での説明担当者、証拠分析担当者、クライアント電子メール起草者など、再利用可能な法的役割を構成します。
監査可能性、ジョブ、および信頼性の制御
実行、ジョブの進行状況、エクスポート、エスカレーション、監査イベント、プロバイダー/モデルの使用、暗号化されたシークレット、レビュー警告、人間による承認ポイントを追跡します。
ローカルファーストでプライバシーを意識した展開
明示的なランタイム パス、安全な Cookie、CORS コントロール、移行、暗号化されたシークレット、保護されたアップロード/データベース ストレージを使用して、ローカルで実行することも、チーム用にデプロイすることもできます。
事実、当事者、管轄権、論調、および選択されたソース文書に基づいて法的通知の草案を作成します。
プレイブックに照らして契約をレビューし、リスク、フォールバック言語、概要、警告、ソース参照を作成します。
問題、時系列、証拠、証人、損害、手続き上のリスクをマッピングして、仲裁の準備をします。
弁論、証拠、電子メール、メモから訴訟準備パッケージを構築します。
立場、利益、情報ギャップ、論点、

譲歩と和解オプション。
正確な質問、管轄区域、事実、アップロードされたソース資料から法的調査メモの草稿を作成します。
法的文書を、保存された条件を使用して別の言語に翻訳し、警告を確認します。
文書に基づいたクライアントの電子メール返信を生成して承認します。
契約書のバージョン間で条項を比較します。
成果物を送信する前に、パートナーレビュー担当者のペルソナを使用します。
AI ブループリントは 2 つの方法で使用できます。
ローカル バージョン: 単独の弁護士、実験、文書レビュー、草案、調査、翻訳、ローカル ファーストのワークフローのために、AI ブループリントを 1 台のマシンでプライベートに実行します。
ネットワーク/サーバー バージョン: 共有ワークスペース、案件アクセス、権限、監査証跡、再利用可能なワークフロー、一元化された構成を使用して、企業、チーム、または法務部門向けに AI ブループリントを実行します。
ローカル バージョンは最も早く開始できる方法です。ネットワーク バージョンは、マルチユーザーの法律実務の方向性です。
選択したモデル、音声、取得モードに応じたプロバイダー API キー
git clone https://github.com/rohasnagpal/AI-Blueprint.git
cd AI ブループリント
python3 -m venv .venv
ソース .venv/bin/activate
pip install -r 要件.txt
Python main.py
開く:
http://127.0.0.1:8000
初めて使用するときは、セットアップ フローを完了し、ワークスペースを作成または選択し、[設定] を開いて必要なモデル キーまたはプロバイダー キーを追加します。
一時的なブートストラップ アカウントが必要なクローズド ローカル デプロイメントの場合は、サーバーを起動する前に AI_BLUEPRINT_BOOTSTRAP_DEFAULT_ADMIN=true を設定し、ブートストラップ資格情報を使用してサインインし、すぐに変更します。パブリック展開にはブートストラップ アカウント パスを使用しないでください。
ローカルのみのベータ版の場合は、ユーザー自身のマシンで AI ブループリントを実行し、次の方法でのみ開きます。
http://127.0.0.1:8000
http://ローカルホスト:8000
アプリをパブリック インターフェイスにバインドしたり、ローカル インターフェイスを共有したりしないでください。

意図的にネットワーク/サーバー展開を実行している場合を除き、LAN 上のポートを使用しないでください。
推奨されるローカルのみの設定:
AI_BLUEPRINT_ENV=開発
AI_BLUEPRINT_SECURE_COOKIES=false
AI_BLUEPRINT_CORS_ORIGINS=http://127.0.0.1:8000,http://localhost:8000
同じマシンのローカルホストの使用には HTTPS は必要ありません。 AI ブループリントが実際のドメイン、リバース プロキシ、トンネル、LAN ホスト名、またはパブリック サーバーを通じて公開される場合は、HTTPS および AI_BLUEPRINT_SECURE_COOKIES=true が必要です。
ローカル ユーザーは自分自身のバックアップに対して責任を負います。これらのランタイム ファイルを一緒にバックアップします。
ai_blueprint.db
ai_blueprint_v2.db
アップロード/
アップロード_v2/
クロマデータベース/
.secret_key
.secret_key_v2
秘密キー ファイルは、復元後に保存された資格情報を復号化するために必要です。バックアップ アーカイブをリポジトリの外に保管し、クライアントの機密データとして扱います。
AI ブループリントはローカルで実行されますが、構成された外部プロバイダーは引き続きデータを受信する可能性があります。有効な設定に応じて、モデル、埋め込み、取得、電子メール、または音声プロバイダーは、プロンプト、ドキュメントの抜粋、埋め込み、認証情報、音声、出力、または関連するメタデータを受信する場合があります。
クライアントの機密資料を使用する前に、各プロバイダーのデータ使用および保持ポリシーを確認してください。外部プロバイダーが構成されていない場合、ローカル ドキュメントとローカル データベースはユーザーのマシン上に残ります。
最速で成功するパスは次のとおりです。
サインインするか、初回実行の管理者セットアップを完了します。
クライアント、紛争、取引、研究プロジェクト、または内部ファイル用の事項を作成します。
設定でプロバイダー API キーを追加します。
関連するドキュメントをアップロードまたは取り込みます。
ドキュメントのインデックスが作成されるまで待ちます。
チャット、音声、準備、契約レビュー、草案、翻訳、または電子メールを使用します。
クライアント、裁判所、規制、取引業務で使用する前に、すべての出力を確認してください。
OpenAI、Anthropic、Groq、OpenRouter、G などの構成可能なプロバイダー統合

emini、Perplexity、Mistral、xAI、Ollama
Chroma、LlamaIndex、センテンストランスフォーマーを使用したオプションのローカル RAG 依存関係
バニラの HTML、CSS、および JavaScript フロントエンド
AI ブループリントは、構成可能なモデル プロバイダー、ドキュメント取得設定、Web 検索プロバイダー、ペルソナ、電子メール設定、ワークスペース、案件、アップロード制限、アプリのブランディング、ユーザー、アクティビティ ログ、およびアプリ内および環境変数を介したデプロイメント制御をサポートします。
合法的に使用するには、この問題に基づいてシステムを構成します。
どのペルソナ、ドラフト設定、準備ワークフロー、または契約レビュー設定を使用する必要があるか
どのユーザーが案件にアクセスできるか
使用前にどの出力を確認する必要があるか
どのモデルプロバイダーがプロンプト、ドキュメントスニペット、埋め込み、出力を受け取ることができるか
運用環境の設定については、「 導入ガイド 」を参照してください。
AI ブループリントは、ローカルファーストの法的 AI ワークスペースとして設計されています。アップロードされたドキュメント、チャット履歴、下書き履歴、翻訳履歴、ローカル データベース、ベクター インデックス、ログ、API キーなどの機密性の高いランタイム データは、バージョン管理にコミットしないでください。
外部モデル、埋め込み、取得、電子メール、または音声プロバイダーを使用する場合は、ドキュメントのテキスト、プロンプト、埋め込み、資格情報、音声、出力がどこにあるかを確認してください。

[切り捨てられた]

## Original Extract

Open source AI-native infrastructure for law firms. - rohasnagpal/AI-Blueprint

GitHub - rohasnagpal/AI-Blueprint: Open source AI-native infrastructure for law firms. · GitHub
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
rohasnagpal
/
AI-Blueprint
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
101 Commits 101 Commits .github .github app app docs docs migrations migrations public public routes routes scripts scripts tests tests .dockerignore .dockerignore .env.production.example .env.production.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE OPERATIONS.md OPERATIONS.md README.md README.md SECURITY.md SECURITY.md ai-blueprint-multiuser-plugin-plan.html ai-blueprint-multiuser-plugin-plan.html alembic.ini alembic.ini database.py database.py docker-compose.example.yml docker-compose.example.yml main.py main.py multi-agent-contract-review-plan.html multi-agent-contract-review-plan.html requirements-local.txt requirements-local.txt requirements.txt requirements.txt webtools.py webtools.py View all files Repository files navigation
A local-first legal AI workspace for document-grounded chat, live voice assistance, legal drafting, contract review, preparation workflows, translation, email drafting, and matter document management.
AI Blueprint is built for lawyers, firms, legal teams, and legal operations teams that need AI to work with confidential documents, matter context, source grounding, review workflows, and auditability.
The goal is not to replace legal judgment. The goal is to give lawyers a structured workspace for reviewing documents, preparing disputes and negotiations, drafting work product, translating legal material, and organizing matter knowledge with stronger control over data, models, and retrieval scope.
Document-grounded legal chat and RAG
Ask questions across uploaded files, matter documents, contracts, emails, precedents, policies, statutes, regulations, and approved sources with scoped retrieval.
Legal drafting workspace
Generate legal notices, agreements, replies, board documents, clauses, client-facing drafts, and other legal work product from facts, parties, terms, instructions, jurisdiction, tone, and optional source documents. Drafts include printable HTML, plain text, assumptions, missing information, review warnings, source usage, progress events, and saved draft history.
Contract review workflows
Review indexed contracts directly from a workspace and matter. Use review depth, playbooks, selected source contracts, and review instructions to produce clause analysis, risks, redline suggestions, summaries, warnings, source references, and downloadable Markdown output.
Preparation workflows
Prepare structured arbitration, litigation, mediation, and negotiation work product from indexed matter documents. Prep runs can produce case snapshots, issue maps, chronologies, evidence matrices, procedural tasks, witness or participant preparation, damages and remedies analysis, risks, gaps, trace data, warnings, and source basis.
Legal research support
Use Chat, Documents mode, Draft, and Prep workflows to research legal questions, analyze uploaded authorities, produce research memos, and test arguments. Current primary navigation does not expose a separate Legal Research screen.
Email and client communication drafting
Poll IMAP inboxes, generate AI-assisted replies using personas and document context, review drafts, and send approved replies through SMTP.
Translation for legal and business documents
Translate pasted text or uploaded documents with legal, business, technical, literal, and plain-language modes, review warnings, translator notes, preserved terms, and downloadable HTML output.
Live voice assistance
Use OpenAI Realtime voice for spoken legal assistance, matter walkthroughs, document questions, and workflow guidance, including document search in Documents mode.
Workspaces, matters, users, and permissions
Organize documents, runs, members, roles, outputs, and access by workspace and matter.
Current navigation
Work from New Chat, Add Document, View Documents, Prep, Workflows, Personas, and Settings. The Prep menu includes Arbitration Prep, Litigation Prep, Mediation Prep, and Negotiation Prep. The Workflows menu includes Contract Review, Draft, Email, and Translate.
Document management
Upload files, ingest URLs, connect local folders, sync folder sources, index documents, search document libraries, and manage deletion.
Personas and role-specific behavior
Configure reusable legal roles such as Contract Reviewer, Litigation Associate, Legal Researcher, Partner Reviewer, Plain-English Explainer, Evidence Analyst, and Client Email Drafter.
Auditability, jobs, and trust controls
Track runs, job progress, exports, escalations, audit events, provider/model use, encrypted secrets, review warnings, and human approval points.
Local-first and privacy-conscious deployment
Run locally or deploy for teams with explicit runtime paths, secure cookies, CORS controls, migrations, encrypted secrets, and protected upload/database storage.
Draft a legal notice from facts, parties, jurisdiction, tone, and selected source documents.
Review a contract against a playbook and produce risks, fallback language, summaries, warnings, and source references.
Prepare for arbitration by mapping issues, chronology, evidence, witnesses, damages, and procedural risks.
Build a litigation preparation package from pleadings, exhibits, emails, and notes.
Prepare mediation or negotiation strategy with positions, interests, information gaps, talking points, concessions, and settlement options.
Draft a legal research memo from a precise question, jurisdiction, facts, and uploaded source materials.
Translate a legal document into another language with preserved terms and review warnings.
Generate and approve a document-grounded client email reply.
Compare clauses across versions of an agreement.
Use a Partner Reviewer persona before sending work product.
AI Blueprint can be used in two ways:
Local version: run AI Blueprint privately on one machine for solo lawyers, experiments, document review, drafting, research, translation, and local-first workflows.
Network/server version: run AI Blueprint for a firm, team, or legal department with shared workspaces, matter access, permissions, audit trails, reusable workflows, and centralized configuration.
The local version is the fastest way to start. The network version is the direction for multi-user legal practice.
Provider API keys depending on the models, voice, and retrieval mode you choose
git clone https://github.com/rohasnagpal/AI-Blueprint.git
cd AI-Blueprint
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
python main.py
Open:
http://127.0.0.1:8000
On first use, complete the setup flow, create or select a workspace, and open Settings to add the model or provider keys you need.
For closed local deployments that need a temporary bootstrap account, set AI_BLUEPRINT_BOOTSTRAP_DEFAULT_ADMIN=true before starting the server, then sign in with the bootstrap credentials and immediately change them. Do not use the bootstrap account path for public deployments.
For a local-only beta, run AI Blueprint on the user's own machine and open it only through:
http://127.0.0.1:8000
http://localhost:8000
Do not bind the app to a public interface or share the local port on a LAN unless you are intentionally running a network/server deployment.
Recommended local-only settings:
AI_BLUEPRINT_ENV=development
AI_BLUEPRINT_SECURE_COOKIES=false
AI_BLUEPRINT_CORS_ORIGINS=http://127.0.0.1:8000,http://localhost:8000
HTTPS is not required for same-machine localhost use. HTTPS and AI_BLUEPRINT_SECURE_COOKIES=true are required when AI Blueprint is exposed through a real domain, reverse proxy, tunnel, LAN hostname, or public server.
Local users are responsible for their own backups. Back up these runtime files together:
ai_blueprint.db
ai_blueprint_v2.db
uploads/
uploads_v2/
chroma_db/
.secret_key
.secret_key_v2
The secret key files are required to decrypt stored credentials after restore. Keep backup archives outside the repository and treat them as confidential client data.
AI Blueprint runs locally, but configured external providers may still receive data. Depending on enabled settings, model, embedding, retrieval, email, or voice providers may receive prompts, document excerpts, embeddings, credentials, audio, outputs, or related metadata.
Review each provider's data usage and retention policy before using confidential client material. If no external provider is configured, local documents and local databases remain on the user's machine.
The fastest successful path is:
Sign in or complete first-run admin setup.
Create a matter for the client, dispute, transaction, research project, or internal file.
Add provider API keys in Settings.
Upload or ingest relevant documents.
Wait until documents are indexed.
Use Chat, Voice, Prep, Contract Review, Draft, Translate, or Email.
Review all outputs before using them in client, court, regulatory, or transaction work.
Configurable provider integrations including OpenAI, Anthropic, Groq, OpenRouter, Gemini, Perplexity, Mistral, xAI, and Ollama
Optional local RAG dependencies with Chroma, LlamaIndex, and sentence-transformers
Vanilla HTML, CSS, and JavaScript frontend
AI Blueprint supports configurable model providers, document retrieval settings, web search providers, personas, email settings, workspaces, matters, upload limits, app branding, users, activity logs, and deployment controls from inside the app and through environment variables.
For legal use, configure the system around the matter:
which persona, draft settings, prep workflow, or contract review settings should be used
which users can access the matter
which outputs need review before use
which model providers may receive prompts, document snippets, embeddings, and outputs
For production settings, see Deployment Guide .
AI Blueprint is designed as a local-first legal AI workspace. Sensitive runtime data such as uploaded documents, chat history, draft history, translation history, local databases, vector indexes, logs, and API keys should not be committed to version control.
When using external model, embedding, retrieval, email, or voice providers, review where document text, prompts, embeddings, credentials, audio, and outputs

[truncated]
