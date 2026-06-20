---
source: "https://github.com/agentark-ai/AgentArk"
hn_url: "https://news.ycombinator.com/item?id=48606186"
title: "Show HN: AgentArk – open-source self-hosted AI agent OS"
article_title: "GitHub - agentark-ai/AgentArk: Personal AI OS, secure first, self learning via GEPA · GitHub"
author: "debankad"
captured_at: "2026-06-20T04:33:12Z"
capture_tool: "hn-digest"
hn_id: 48606186
score: 2
comments: 0
posted_at: "2026-06-20T03:45:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentArk – open-source self-hosted AI agent OS

- HN: [48606186](https://news.ycombinator.com/item?id=48606186)
- Source: [github.com](https://github.com/agentark-ai/AgentArk)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T03:45:06Z

## Translation

タイトル: Show HN: AgentArk – オープンソースのセルフホスト型 AI エージェント OS
記事タイトル: GitHub - Agentark-ai/AgentArk: パーソナル AI OS、セキュアファースト、GEPA 経由の自己学習 · GitHub
説明: パーソナル AI OS、セキュアファースト、GEPA 経由の自己学習 - Agentark-ai/AgentArk

記事本文:
GitHub - Agentark-ai/AgentArk: パーソナル AI OS、セキュアファースト、GEPA 経由の自己学習 · GitHub
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
エージェントアーク-ai
/
エージェントアーク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
88 コミット 88 コミット .cargo .cargo .github .github アセット アセット アセット ブリッジ ブリッジ クライアント/ コンパニオン クライアント/ コンパニオン フロントエンド フロントエンド スクリプト スクリプト src src テスト テスト .dockerignore .dockerignore .gitignore .gitignore ACKNOWLEDGMENTS.md ACKNOWLEDGMENTS.md API.md API.md ARCHITECTURE.md ARCHITECTURE.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile Dockerfile.lowmem Dockerfile.lowmem LICENSE-APACHE LICENSE-APCHE LICENSE-MIT LICENSE-MIT README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md VERIFY.md VERIFY.md エージェントアーク エージェントアーク エージェントアーク.bat エージェントアーク.bat build.ps1 ビルド.ps1 ビルド.rs ビルド.rs ビルド.sh ビルド.sh 拒否.toml 拒否.toml docker-compose.dev.yml docker-compose.dev.yml docker-compose.lowmem-build.yml docker-compose.lowmem-build.yml docker-compose.lowmem.yml docker-compose.lowmem.yml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントではありません。エージェントのための箱舟: プロンプトとツールから構築し、アプリ、オートメーション、またはウォッチャーとしてデプロイし、ノイズの多いコンテキストを抽出し、あらゆるアクションを監視し、あらゆる境界を保護し、使用状況から自己進化します。
エージェントのライフサイクル全体に対応するセルフホスト型ランタイム。
構造化されたプロンプト、ツール、統合からエージェントを構築します。これらをライブ アプリ、スケジュールされた自動化、条件付きウォッチャー、またはチャット セッションとして展開します。
アクション トレース、障害分類、ドリフト検出を使用して、Sentinel のすべてのステップを監視します。インテントの分類、出力ガード、承認ゲート、アクションごとの承認により、あらゆる機能の境界を保護します。
ArkDistill でコンテキストを保存: ノイズの多いブラウザー ページ、ログ、トレース、HTML、および int の前に確定的なツール出力圧縮

放出ダンプがモデルに到達し、多くの場合、ノイズの多い出力が 60 ～ 90% 削減されます。
プロンプト、分類子、ルーティング ポリシー、スペシャリストの動作、およびコンテキスト保存プロファイルを独自の使用方法から自己進化させます。
Reflect を通じて、日、週、または月を確認します。チャット、ArkOrbit、アプリ、目標、ウォッチャー、メモリ、バックグラウンド エージェント、使用状況、および学習したワークフローがクラスター化されている場所のローカルの視覚的なパノラマです。
チャット、メモリ、デバイス、統合、レビュー可能なアクションがすべて 1 か所にまとめられ、すべてマシン上にあり、デフォルトでは非公開です。
~3.1GB の Docker イメージ · アイドル状態で ~500MB、負荷下で定常状態で ~1GB RAM (コンテナー 5 個、埋め込みが読み込まれています) · AES-256-GCM 暗号化 · モデルに依存しない
インストール・
特徴・
アークコア・
構成・
建築・
セキュリティ ·
API・
貢献する ·
ディープウィキ ·
不和
AgentArk はベータ版であり、本番環境向けではありません。間違いを犯し、ワークスペース内のファイルを上書きする可能性があります。 Docker 境界によりホスト ファイル システムから切り離されますが、コンテナーにマウントするものはすべてスコープ内に含まれます。承認を維持し、データをバックアップし、結果を確認します。
バグや粗い部分が予想されます。 AgentArk は 1 人 ( @debankadas ) によって公開で構築および保守されているため、表面積が広く、船が見つかるとすぐに修正されます。何かが壊れた場合は、問題を開いてください。再現、ログ、スクリーンショットが非常に役に立ちます。
> 毎週平日の午前 9 時に、天気を含むその日の概要を送ってください。
カレンダー、緊急メール、期限を過ぎたタスクなど。
> 私は Telegram で簡潔な回答と毎日の更新を好むことを覚えておいてください。
> 受信箱にクライアントからの緊急メッセージがないか監視し、返信しない場合は警告を発します。
> このメッセージに対する返信の下書きを作成し、送信する前に質問してください。
> 新しいプロジェクトのランディング ページを作成してください。パブリック URL を使用してデプロイします。
> マルチエージェント アーキテクチャに関する最近の論文を Web で検索します。
上位 3 つを要約し、ドキュメントに保存します。

> Linear 統合をインストールし、割り当てられた問題をリストします。
> Google カレンダーに接続すると、会議の 10 分前にリマインダーが表示されます。
> Stripe 支払いアラートを Telegram に投稿する Webhook を設定します。
返事だけでは終わりません。設定の保存、フォローアップのスケジュール、概要の配信、返信の下書き、更新の監視、統合の接続、または作業を永続的なタスクに昇格させて後で戻ってくることができます。
AgentArk はエージェントではありません。それはエージェントのための方舟です。 Ark はセキュリティ層です。内部のすべてのエージェントが実行できる内容を含み、監視し、強制するラッパーと、すべてのアクションがレビュー可能になる監査面です。エージェントは、Ark 内で実行されるものです。チャット ハンドラー、デプロイされたアプリ、スケジュールされた自動化、条件付きウォッチャー、ルーターによって派遣される専門のサブエージェントです。 Ark を使用すると、それらのユーザーが実際のデータを安全に参照できるようになります。
その境界内で、AgentArk はまた、要求されたエージェントを構築し、パブリック URL、オートメーション、またはウォッチャーを備えたアプリとして展開し、すべてのステップを監視し、モデル コンテキストを拡張する前にノイズの多いツール出力を抽出し、使用状況からプロンプト、ポリシー、およびコンテキスト保存プロファイルを自己進化させます。チャット、メモリ、タスク、統合、ドキュメント、コンパニオン デバイス、監査証跡は、マシン上の 1 つのプライベート ワークスペースに一緒に存在します。ユーザーの好みを追跡し、毎日の概要を配信し、チャネル全体でフォローアップし、ルーチンをスケジュールし、バックグラウンドで監視し、アプリを構築し、要求に応じて安全にアクションを実行できます。
ユーザーとともに進化するように作られています。受け入れられた作業、ユーザーの修正、繰り返されるルーチン、ライブ ツールの結果はローカル メモリ、プロンプト、ルーティング、および戦略に反映されるため、OS は毎回のセッションのように動作するのではなく、ワークフローとより連携します。

e.
返信を短く書き換え続けると、デフォルトで簡潔さを保つよう学習します
特定のツールパスがタスクで成功し続ける場合、そのパスを再度選択する可能性が高くなります。
ブラウザーのページ、ログ、またはトレースがコンテキストを無駄にし続ける場合、Evolve は必要なフィールドを維持しながらそれらを縮小する ArkDistill プロファイルを改善できます。
ブリーフィング、ルーティング、フォローアップの方法を修正すると、今後の実行にはその修正が反映されます
データは常に手元に残ります。あなたの秘密は暗号化されています。危険な行動について最終決定権を握るのはあなたです。
注: AgentArk は現在、1 つのグローバル ワークスペースとして実行されます。プロジェクト固有のワークスペースとプロジェクト スコープの UI/API の動作は、意図的にフェーズ 2 に延期されます。
あなたがする場所に住んでいます。マシン上の Docker です。メモリ、シークレット、統合トークン、会話履歴、監査証跡はすべてローカル ボリュームにあり、他人のクラウドに保存されることはありません。依存する管理対象バックエンドも、保持する必要のあるアカウントも、オプトアウトする必要のあるテレメトリもありません。
私たちではなく、あなたがモデルに支払います。 AgentArk を Ollama または任意のローカル モデルに指定すると、インストール後のすべてのプロンプトは完全に無料です。料金制限や突然の請求書はありません。独自の Anthropic、OpenAI、Gemini、または Groq キーを持参すると、プロバイダーの公開料金を直接支払います。 AgentArk は、単一のトークンをプロキシしたり、仲介したり、マークアップしたりすることはありません。サブスクリプション、シートごと、最低価格はありません。
デザインに縛られる。世界に関わるあらゆるアクションは許可ゲートを通過します。エージェントは、事前承認されていないものに対する承認キューを備えた Docker 境界内で実行されます。ホスト ファイルシステムは、参照したいものを明示的にマウントしない限り、立ち入り禁止のままになります。
あなたに適応します。受け入れられた作業、修正、およびライブツールの結果は、ローカルメモリ、プロンプト、およびルーティングにフィードバックされます。何週間も使用すると、OS は実際の作業方法、つまりあなたの好みに応じて形作られます。

ローアップ スタイル、ルーティング設定、タスクを継続的に成功させるツール パスなど、他のすべてのユーザーの一般的な組み合わせによるものではありません。
オープンで検査可能。 MIT と Apache 2.0。すべての行を読み、フォークし、実行します。すべてのアクションの監査証跡により、チャット、自動化、ウォッチャー、デプロイされたアプリ、統合にわたって、エージェントがいつ何を行ったのか、なぜ行ったのかを常に確認できます。
クイック スタート (Docker イメージ、ソース クローンなし)
カール -sSL https://raw.githubusercontent.com/agentark-ai/AgentArk/main/scripts/install.sh |バッシュ
Windows:
irm https://raw.githubusercontent.com/agentark - ai/AgentArk/main/scripts/install.ps1 |アイエックス
インストーラーは、Docker をインストールする前に、Docker が見つからないかどうかを確認し、必要に応じて Docker デスクトップを起動し、Compose/ランタイム ヘルパー ファイルのみをダウンロードし、公開された AgentArk イメージを取得して、スタックを開始します。通常の使用には Git クローンは必要ありません。
http://localhost:8990 を開き、[設定] で LLM プロバイダーを選択し、チャットを開始します。
AgentArk を構築またはコーディングしている場合にのみ Git を使用してください。
git clone https://github.com/agentark-ai/AgentArk.git && cd AgentArk
AGENTARK_IMAGE=agentARK:dev ./scripts/start.sh ビルド
Windows ソース チェックアウトの場合:
git clone https://github.com/agentark-ai/AgentArk.git && cd AgentArk
scripts\start.bat ビルド
ソース ビルドは、公開された AgentArk ランタイム イメージを GHCR からプルしません。 Docker ビルドのベース イメージと、ローカル イメージのコンパイルに必要なパッケージの依存関係を引き続きダウンロードします。
Web UIを使用します。 AgentArk は、 http://localhost:8990 にある Docker Compose スタックと Mission Control を通じて実行されるように設計されています。
サポートされているインストール パスでは、Docker Compose のデフォルトと実行時状態の名前付き Docker ボリュームが使用され、更新後もそれらのボリュームが保持されます。 AgentArk は、ルート プロジェクト .env を作成したり、必要としたりしません。生成されたアプリにはフレームワークが所有する環境が含まれる場合があります

ファイルは必要に応じて独自のアプリ ディレクトリ内に保存されますが、秘密キーは AgentArk の管理された秘密ストレージまたはランタイム インジェクション パスに残ります。
Pulse は、フレームワーク管理のバックアップを自動的に作成します。デフォルトでは、AgentArk は 14 日ごとに新しい管理バックアップをチェックし、Sentinel がシステムがアイドル状態であると判断した場合にのみバックアップを作成します。チャット、アプリの作業、ブラウザ セッション、サンドボックス コンテナ、または重いバックグラウンド作業がアクティブな場合、バックアップは延期され、後で再試行されます。バックアップ作業は、メイン API リクエスト パスではなく、バックグラウンド タスクと子プロセスで実行されます。
バックアップは、タイムスタンプ付きのアーティファクトとして /app/data/backups に書き込まれます。
Agentark-managed-*.dump - 会話、メッセージ、タスク、ウォッチャー、設定、メモリ/ドキュメント インデックス、トレース、ログ、およびその他の DB にバックアップされた状態の Postgres 論理ダンプ。
Agentark-managed-*.data.tar.gz - バックアップ ディレクトリ自体を除く /app/data のアーカイブ。
Agentark-managed-*.config.tar.gz - 構成ボリュームが存在する場合の /app/config のアーカイブ。
AgentArk 自体がバックアップ ディレクトリを作成します。バックアップの作成が失敗した場合、Pulse はデータの安全性に関する重大な結果を報告し、ユーザーに通知します。ユーザーにバックアップ フォルダーを手動で作成するよう求めるべきではありません。
完全なインストールを回復するには、オペレーター ボリュームも保持してください

[切り捨てられた]

## Original Extract

Personal AI OS, secure first, self learning via GEPA - agentark-ai/AgentArk

GitHub - agentark-ai/AgentArk: Personal AI OS, secure first, self learning via GEPA · GitHub
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
agentark-ai
/
AgentArk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
88 Commits 88 Commits .cargo .cargo .github .github assets assets bridges bridges clients/ companion clients/ companion frontend frontend scripts scripts src src tests tests .dockerignore .dockerignore .gitignore .gitignore ACKNOWLEDGMENTS.md ACKNOWLEDGMENTS.md API.md API.md ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile Dockerfile.lowmem Dockerfile.lowmem LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md VERIFY.md VERIFY.md agentark agentark agentark.bat agentark.bat build.ps1 build.ps1 build.rs build.rs build.sh build.sh deny.toml deny.toml docker-compose.dev.yml docker-compose.dev.yml docker-compose.lowmem-build.yml docker-compose.lowmem-build.yml docker-compose.lowmem.yml docker-compose.lowmem.yml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh View all files Repository files navigation
Not an agent. An Ark for agents: build from prompts and tools, deploy as apps, automations, or watchers, distill noisy context, monitor every action, secure every boundary, self-evolve from your usage.
A self-hosted runtime for the full agent lifecycle.
Build agents from structured prompts, tools, and integrations. Deploy them as live apps, scheduled automations, conditional watchers, or chat sessions.
Monitor every step through Sentinel with action traces, failure classification, and drift detection. Secure every capability boundary with intent classification, output guards, approval gates, and per-action authorization.
Save context with ArkDistill: deterministic tool-output compaction before noisy browser pages, logs, traces, HTML, and integration dumps reach the model, often cutting noisy outputs by 60-90%.
Self-evolve prompts, classifiers, routing policies, specialist behavior, and context-saving profiles from your own usage.
Review your day, week, or month through Reflect: a local visual panorama of where chat, ArkOrbit, apps, goals, watchers, memory, background agents, usage, and learned workflows clustered.
Chat, memory, devices, integrations, and reviewable actions, all in one place, all on your machine, private by default.
~3.1GB Docker image · ~500MB idle, ~1GB RAM steady-state under load (5 containers, embeddings loaded) · AES-256-GCM encrypted · model-agnostic
Install ·
Features ·
Ark Core ·
Configuration ·
Architecture ·
Security ·
API ·
Contributing ·
DeepWiki ·
Discord
AgentArk is in beta — not for production. It can make mistakes and overwrite files inside its workspace. The Docker boundary keeps it off your host filesystem, but anything you mount into the containers is in scope. Keep approvals on, back up data, verify results.
Bugs and rough edges are expected. AgentArk is built and maintained by one person ( @debankadas ) in the open, so the surface area is large and fixes ship as they're found. Please open an issue when something breaks — repros, logs, and screenshots help a lot.
> Every weekday at 9am, send me a daily brief with weather,
calendar, urgent email, and overdue tasks.
> Remember that I prefer concise answers and daily updates in Telegram.
> Watch my inbox for urgent client messages and alert me if I do not reply.
> Draft a reply to this message and ask before sending it.
> Build me a landing page for my new project. Deploy it with a public URL.
> Search the web for recent papers on multi-agent architectures,
summarize the top 3, and save them to my documents.
> Install the Linear integration and list my assigned issues.
> Connect my Google Calendar and remind me 10 minutes before every meeting.
> Set up a webhook that posts Stripe payment alerts to my Telegram.
It does not stop at a reply. It can save the preference , schedule the follow-up , deliver the brief , draft the reply , watch for updates , connect an integration , or promote the work into a durable task and come back later.
AgentArk is not an agent. It is an Ark for agents. The Ark is the security layer: the wrapper that contains, observes, and enforces what every agent inside it is allowed to do, and the audit surface where every action becomes reviewable. Agents are the things that run inside the Ark - chat handlers, deployed apps, scheduled automations, conditional watchers, specialist sub-agents dispatched by the router. The Ark is what makes any of them safe to point at your real data.
Inside that boundary AgentArk also builds the agents you ask for, deploys them as apps with public URLs, automations, or watchers, monitors every step, distills noisy tool output before it expands the model context, and self-evolves prompts, policies, and context-saving profiles from your usage. Chat, memory, tasks, integrations, documents, companion devices, and audit trails live together in one private workspace on your machine. It can keep track of your preferences, deliver a daily brief, follow up across channels, schedule routines, monitor things in the background, build apps, and take action safely when you ask.
It is built to evolve with you. Accepted work, user corrections, repeated routines, and live tool outcomes are reflected into local memory, prompts, routing, and strategy so the OS gets more aligned with your workflow instead of acting like every session is day one.
If you keep rewriting replies to be shorter, it learns to stay concise by default
If a certain tool path keeps succeeding for a task, it becomes more likely to choose that path again
If browser pages, logs, or traces keep wasting context, Evolve can improve ArkDistill profiles that shrink them while preserving required fields
If you correct how it briefs, routes, or follows up, future runs reflect that correction
Your data stays with you. Your secrets are encrypted. You keep the final say on risky actions.
Note: AgentArk currently runs as one global workspace. Project-specific workspaces and project-scoped UI/API behavior are intentionally deferred to phase 2.
Lives where you do. Docker on your machine, period. Memory, secrets, integration tokens, conversation history, audit trails — all in local volumes, never in someone else's cloud. No managed backend you depend on, no account you have to keep, no telemetry you have to opt out of.
You pay your model, not us. Point AgentArk at Ollama or any local model and every prompt after install is genuinely free — no rate limits, no surprise invoice. Bring your own Anthropic, OpenAI, Gemini, or Groq key and you pay the provider's published rate directly; AgentArk never proxies, intermediates, or marks up a single token. No subscription, no per-seat, no minimum.
Bounded by design. Every action that touches the world goes through a permission gate. The agent runs inside a Docker boundary with an approval queue for anything not pre-authorized. Your host filesystem stays off-limits unless you explicitly mount what you want it to see.
Adapts to you . Accepted work, your corrections, and live tool outcomes feed back into local memory, prompts, and routing. Over weeks of use the OS gets shaped by how you actually work — your follow-up style, your routing preferences, the tool paths that keep succeeding for your tasks — not by a generic mix of every other user.
Open and inspectable. MIT and Apache 2.0. Read every line, fork it, run it. Audit trails on every action mean you can always see what the agent did, why, and when — across chat, automations, watchers, deployed apps, and integrations.
Quick start (Docker image, no source clone)
curl -sSL https://raw.githubusercontent.com/agentark-ai/AgentArk/main/scripts/install.sh | bash
Windows:
irm https: // raw.githubusercontent.com / agentark - ai / AgentArk / main / scripts / install.ps1 | iex
The installer asks before installing Docker if Docker is missing, starts Docker Desktop when needed, downloads only the Compose/runtime helper files, pulls the published AgentArk image, and starts the stack. No Git clone is needed for normal use.
Open http://localhost:8990 , pick your LLM provider in Settings, start chatting.
Use Git only if you are building or coding AgentArk:
git clone https://github.com/agentark-ai/AgentArk.git && cd AgentArk
AGENTARK_IMAGE=agentark:dev ./scripts/start.sh build
On Windows source checkouts:
git clone https://github.com/agentark-ai/AgentArk.git && cd AgentArk
scripts\start.bat build
Source builds do not pull the published AgentArk runtime image from GHCR. They still download Docker build base images and package dependencies needed to compile the local image.
Use the Web UI. AgentArk is designed to run through the Docker Compose stack and Mission Control at http://localhost:8990 .
The supported install path uses Docker Compose defaults plus named Docker volumes for runtime state and preserves those volumes across updates. AgentArk does not create or require a root project .env . Generated apps may have framework-owned env files inside their own app directories when required, but secret keys stay in AgentArk's managed secret storage or runtime injection path.
Pulse creates framework-managed backups automatically. By default, AgentArk checks for a fresh managed backup every 14 days and only creates one when Sentinel sees the system as idle; if chats, app work, browser sessions, sandbox containers, or heavy background work are active, the backup is deferred and retried later. Backup work runs in background tasks and child processes, not on the main API request path.
Backups are written under /app/data/backups as timestamped artifacts:
agentark-managed-*.dump - Postgres logical dump for conversations, messages, tasks, watchers, settings, memory/document indexes, traces, logs, and other DB-backed state.
agentark-managed-*.data.tar.gz - archive of /app/data , excluding the backup directory itself.
agentark-managed-*.config.tar.gz - archive of /app/config when that config volume is present.
AgentArk creates the backup directory itself. If backup creation fails, Pulse raises a critical data-safety finding and notifies the user; users should not be asked to create the backup folder manually.
For full install recovery, also keep an operator volume

[truncated]
