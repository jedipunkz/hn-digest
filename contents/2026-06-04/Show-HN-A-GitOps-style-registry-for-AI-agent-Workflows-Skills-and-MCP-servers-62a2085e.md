---
source: "https://github.com/Friz-zy/ai-capability-registry"
hn_url: "https://news.ycombinator.com/item?id=48397625"
title: "Show HN: A GitOps-style registry for AI agent Workflows, Skills and MCP servers"
article_title: "GitHub - Friz-zy/ai-capability-registry: Curated AI Workflows, Skills & MCP servers for secure multi-agent development · GitHub"
author: "frizzy"
captured_at: "2026-06-04T13:14:10Z"
capture_tool: "hn-digest"
hn_id: 48397625
score: 1
comments: 0
posted_at: "2026-06-04T12:21:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A GitOps-style registry for AI agent Workflows, Skills and MCP servers

- HN: [48397625](https://news.ycombinator.com/item?id=48397625)
- Source: [github.com](https://github.com/Friz-zy/ai-capability-registry)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T12:21:17Z

## Translation

タイトル: HN を表示: AI エージェントのワークフロー、スキル、および MCP サーバー用の GitOps スタイルのレジストリ
記事のタイトル: GitHub - Friz-zy/ai-capability-registry: 安全なマルチエージェント開発のための厳選された AI ワークフロー、スキル、MCP サーバー · GitHub
説明: 安全なマルチエージェント開発のための厳選された AI ワークフロー、スキル、MCP サーバー - Friz-zy/ai-capability-registry

記事本文:
GitHub - Friz-zy/ai-capability-registry: 安全なマルチエージェント開発のための厳選された AI ワークフロー、スキル、MCP サーバー · GitHub
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
フリズジー
/
AI 機能レジストリ
公共
通知
通知設定を変更するにはサインインする必要があります
アディ

ナビゲーション オプション
コード
Friz-zy/ai-capability-registry
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
55 コミット 55 コミット .github/ workflows .github/ workflows external external mcp-catalog.d mcp-catalog.d mcp mcp レジストリ レジストリ ロール ロール スクリプト スクリプト スキル-カタログ.d スキル-カタログ.d スキル スキル テンプレート テンプレート ワークフロー ワークフロー .gitignore .gitignore .gitmodules .gitmodules AGENTS.full-registry.md.template AGENTS.full-registry.md.template AGENTS.mcp.md.template AGENTS.mcp.md.template AGENTS.md AGENTS.md AGENTS.skills.md.template AGENTS.skills.md.template AGENTS.workflows-skills.md.template AGENTS.workflows-skills.md.template AGENTS.workflows.md.template AGENTS.workflows.md.template COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md ai-agents.md ai-agents.md Capability-routing.md capability-routing.md mcp-catalog.md mcp-catalog.md skill-catalog.md skill-catalog.md workflows-catalog.md workflows-catalog.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
警告
AGENTS.full-registry.md.template を介してこのレジストリに接続すると、エージェントは実行時にルーティング インデックスと選択された機能ガイドを読み取るようになります。これにより、最小限の静的エージェント設定よりも大幅に多くのモデル コンテキストとトークンが使用される可能性があります。コンテキストの予算が重要な場合は、タスク、役割、キーワード、またはプロジェクト固有のサブセットを優先します。
AI Capability Registry は、AI エージェント向けの動的機能ルーティングの実験です。現在のタスクで必要な場合にのみ、適切なスキル、MCP サーバー、ワークフロー、統合指示をロードします。
このリポジトリは、すべてのエージェントに巨大な静的プロンプト、ランダムなツールの山、または利用可能なすべての MCP サーバーを詰め込むのではなく、エージェントの機能をバージョン管理されたインフラストラクチャとして扱います。エージェントはユーザーを解決できます

のリクエストをタスク、ロール、キーワードごとに処理し、関連するスキル手順または MCP 接続ガイダンスのみを段階的にロードします。
これは、すべてのセットアップにとって理論的に最適な値ではありません。通常、1 つの特定のタスクに対して、最小限のスキル セットと MCP サーバーを備えた慎重に手動で調整されたエージェントのほうが、よりシンプルで高速で、推論が容易になります。レジストリは、より一般的な中間点、つまり多くのタスクにわたって多くのエージェントを実行し、「万が一に備えて」すべてのエージェントに対してスキルと MCP サーバーの広範なリストを有効にするよりも優れたものを必要とするチームとユーザーを対象としています。多くのエージェント ツールは、その広範な接続パスを簡単にします。このプロジェクトでは、エージェントごとの完全に最小限の構成と、デフォルトですべてが有効になる機能の広がりとの間の、より明示的なルーティング層を調査します。
レジストリを最小限の静的セットアップの開始点として使用することもできます。ユーザーは、AI アシスタントに共有レジストリを検査して、特定の役割、タスク、またはエージェント構成に役立つ最小のスキル セットと MCP サーバーを組み立てるよう依頼できます。このモードでは、レジストリはランタイム層そのものではありません。これは、タスク固有のエージェント設定が導出される、レビューされた機能の共通のソースです。
目標は、AI エージェントの機能管理をよりモジュール化し、再現可能で、レビュー可能にすることです。
上流のソースからスキル、MCP サーバー、ワークフロー、エージェントの統合を発見して厳選します。
実行時ルーティングのタスク、役割、キーワードごとに機能を整理します。
有効な機能をローカル エージェントの状態で非表示にするのではなく、Git で明示的に維持します。
信頼できる機能、レビューされた機能、および候補の機能を分離します。
ホストされたエンドポイントや制約のある Docker コンテナなど、より安全な MCP ランタイムを好みます。
プロンプト/ツール構成を複製する代わりに、マルチエージェント設定で同じ機能レジストリを共有できるようにします。
これは

pository は、AI エージェント機能のスプロールを再現可能な GitOps スタイルのレジストリに変えます。上流のソースをピン留めされたサブモジュールとしてダウンロードし、検出されたすべてのスキルと MCP エントリをインベントリし、役割、タスク、キーワードごとに編成されたルーティング カタログとシンボリックリンク パックを生成します。
git clone --recurse-submodules < リポジトリ URL >
cd ai-capability-registry
./scripts/bootstrap.sh
複数のリポジトリ間で共有して使用するには、このレジストリを ~/.ai-registry に保持します。
git clone --recurse-submodules < リポジトリ URL > ~ /.ai-registry
cd ~ /.ai-registry
./scripts/bootstrap.sh
ブートストラップ スクリプトは次のことを行います。
外部サブモジュールを同期してコミットを修正する
レジストリ構成を検証する
skill-catalog.d/ の下のプロバイダー チャンクを検出されたスキルと同期します
ワークフロー ランタイム命令、ルーティング インデックス、人間が判読できるワークフロー カタログを registry/workflows.yaml から生成します。
結合されたスキル/ルーティング カタログとシンボリックリンク パックを skill-catalog.d/ から生成します。
mcp-catalog.d/ から MCP ルーティング カタログを生成します。
エージェントは、このレジストリを 2 つの方法で使用できます。実行時の動的ルーティング、または特定のエージェント、役割、またはタスクのレジストリから派生した最小限の静的セットアップです。
AGENTS.full-registry.md.template を使用して、推奨されるワークフロー優先セットアップで別のリポジトリを ~/.ai-registry に保存されている共有レジストリに接続します。
cp ~ /.ai-registry/AGENTS.full-registry.md.template /path/to/project/AGENTS.md
フルレジストリ テンプレートは意図的にブートローダーとしてのみ使用されます。これはエージェントに、capability-routing.md を読んで従い、 workflows/workflow.md から開始し、ワークフローの選択が必要な場合にのみ workflows/routing.md を使用し、ワークフローで選択されたロールを role/ からロードし、必要な場合にのみそれらのロールをスキルと MCP にルーティングするように指示します。より具体的なローカル リポジトリの指示は、引き続き共有ガイダンスをオーバーライドします。
ルートテンプレート

より狭いセットアップで使用できます。
レジストリが別の場所にインストールされている場合は、コピーした AGENTS.md 内のパスを編集します。
エージェントが具体的なスキルまたは MCP サーバーをロードする前にプロセス ガイダンスを選択する必要がある場合は、生成されたワークフロー ランタイム命令を使用します。
<レジストリへのパス>/workflows/workflow.md
ワークフロー スコープが割り当てられていない場合に完全なワークフロー ルーティングを行うには、次を使用します。
<レジストリへのパス>/workflows/routing.md
リポジトリ間で使用する場合は、部分的なワークフローのみの命令を追加するよりも、AGENTS.full-registry.md.template または AGENTS.workflows.md.template をインストールすることを優先します。
ワークフロー ルーティングは、スキル ルーティングの前に読む必要があります。ワークフローは、タスクにどのようにアプローチするか、どのロールが参加する必要があるか、どのステージとゲートが適用されるか、どの成果物が予想されるかを答えます。次に、役割は、ワークフローの自分の部分にどのスキルをロードするかを決定します。
エージェントがリクエストごとにスキルを動的に選択する必要がある場合は、生成されたスキル ランタイム命令を使用します。
<レジストリへのパス>/skills/skills.md
レジストリ スコープが割り当てられていない場合に完全なロールとタスク機能のルーティングを行うには、次を使用します。
<レジストリへのパス>/skills/routing.md
動的ルーティングの場合、エージェントは通常、ワークフローで選択された役割プロンプトを通じて skill/skills.md にアクセスする必要があります。スキルのみの直接使用は、AGENTS.skills.md.template または同等のブートローダーを通じて引き続きサポートされます。
直接エージェント構成の場合、より狭い静的セットアップが必要な場合は、レジストリ全体ではなく、生成された 1 つのパックをアタッチします。
<レジストリへのパス>/skills/packs/roles/<role-id>/
<レジストリへのパス>/skills/packs/tasks/<タスクID>/
<レジストリへのパス>/skills/packs/keywords/<キーワード>/
skill/packs/all/ は存在しますが、主に検査や実験に役立ちます。実際のエージェントにはタスク、ロール、キーワード、またはカスタムの最小パックを優先します。
エージェントが必要な場合は、生成された MCP ランタイム命令を使用します。

タスクで必要な場合にのみ MCP サーバーを選択します。
<レジストリへのパス>/mcp/mcp.md
レジストリ スコープが割り当てられていない場合に完全なロールとタスク機能のルーティングを行うには、次を使用します。
<レジストリへのパス>/mcp/routing.md
動的ルーティングの場合、エージェントは通常、ワークフローで選択された役割のプロンプトまたはワークフロー ステージのガイダンスを通じて mcp/mcp.md に到達する必要があります。 MCP のみの直接使用は、AGENTS.mcp.md.template または同等のブートローダーを通じて引き続きサポートされます。
直接エージェント構成の場合は、生成された接続メタデータをエージェントの MCP 構成スキーマに適合させます。
<レジストリへのパス>/mcp/servers/<サーバーID>/connection.json
<レジストリへのパス>/mcp/servers/<サーバーID>/SKILL.md
ホストされた MCP サーバーは mcp/web.md に従う必要があります。 Docker MCP サーバーは mcp/docker.md に従う必要があり、特権コンテナー、ホスト ネットワーキング、Docker ソケット マウント、無制限のホスト マウント、SSH キー マウント、およびクラウド資格情報ディレクトリ マウントを避ける必要があります。
生成されたファイルを直接編集しないでください。ソースチャンクを編集して再生成します。
スキルについては、 skill-catalog.d/*.yaml 内の有効なキーワードのみを編集します。新しく検出されたスキルはデフォルトで無効に追加され、削除されたアップストリーム スキルは存在します: false のままになります。
MCP の場合は、手動で厳選したエントリを mcp-catalog.d/manual.yml に保存します。 official-mcp-registry.yml および docker-mcp-registry.yml から自動検出されたエントリは、レビューされ、有用なキーワードが割り当てられ、enabled: true で明示的に昇格されるまで無効のままにする必要があります。
重要な MCP フィールドは、enabled、exists、trust、runtime、transport.type、default_mode、およびキーワードです。
コマンド
使用する
./scripts/bootstrap.sh
完全同期、検証、ワークフロー生成、スキル生成、MCP 生成
./scripts/generate-workflows.py
workflows/workflow.md 、 workflows/routing.md 、 workflows-catalog.md 、およびワークフロー カタログを registry/workflows.yaml から再生成します。
./スクリプト

/discover-skills.py
skill-catalog.d/ からスキル カタログとパックを再生成します。
./scripts/discover-mcp.py
無効になっている候補 MCP エントリを設定済みの上流ソースからインポートします
./scripts/generate-mcp.py
mcp-catalog.d/ から mcp/ と mcp-catalog.md を再生成します。
./scripts/validate-registry.py
レジストリの YAML 構造とポリシー制約を検証する
信頼とポリシー
信頼レベル: セキュリティでレビューされた公式または信頼できるソースについては信頼され、手動でレビューされた既知のメンテナについてはレビューされ、検出専用エントリの候補となります。
デフォルトでは、信頼できる機能またはレビューされた機能を使用します。
発見された候補者はレビューされるまで無効のままにしておきます。
ホスト型 MCP または制約付き Docker MCP を優先します。
直接ノード、 npx 、 python 、 pip 、および uvx MCP ランナーを拒否します。
カールは決して使用しないでください | sh 、特権付き Docker、ホスト ネットワーキング、Docker ソケット マウント、または無制限のホスト マウント。
スキルと MCP サーバーは、現在のタスクに関連する場合にのみロードします。
現在、信頼できるスキル ソースには、Anthropic スキルとナレッジワーク プラグイン、OpenAI スキル、Vercel Agent Skills、Agent Skills 仕様、Superpowers 開発方法論スキル、Trail of Bits セキュリティ ワークフロー、および Kilo Marketplace スキルが含まれます。
厳選された AI ワークフロー、スキル、MCP

[切り捨てられた]

## Original Extract

Curated AI Workflows, Skills & MCP servers for secure multi-agent development - Friz-zy/ai-capability-registry

GitHub - Friz-zy/ai-capability-registry: Curated AI Workflows, Skills & MCP servers for secure multi-agent development · GitHub
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
Friz-zy
/
ai-capability-registry
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Friz-zy/ai-capability-registry
main Branches Tags Go to file Code Open more actions menu Folders and files
55 Commits 55 Commits .github/ workflows .github/ workflows external external mcp-catalog.d mcp-catalog.d mcp mcp registry registry roles roles scripts scripts skill-catalog.d skill-catalog.d skills skills templates templates workflows workflows .gitignore .gitignore .gitmodules .gitmodules AGENTS.full-registry.md.template AGENTS.full-registry.md.template AGENTS.mcp.md.template AGENTS.mcp.md.template AGENTS.md AGENTS.md AGENTS.skills.md.template AGENTS.skills.md.template AGENTS.workflows-skills.md.template AGENTS.workflows-skills.md.template AGENTS.workflows.md.template AGENTS.workflows.md.template CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md ai-agents.md ai-agents.md capability-routing.md capability-routing.md mcp-catalog.md mcp-catalog.md skill-catalog.md skill-catalog.md workflows-catalog.md workflows-catalog.md View all files Repository files navigation
Warning
Connecting this registry through AGENTS.full-registry.md.template makes agents read routing indexes and selected capability guides at runtime. This can use substantially more model context and tokens than a minimal static agent setup. Prefer task-, role-, keyword-, or project-specific subsets when context budget matters.
AI Capability Registry is an experiment in dynamic capability routing for AI agents: load the right skills, MCP servers, workflows, and integration instructions only when the current task needs them.
Instead of stuffing every agent with a huge static prompt, a random pile of tools, or every available MCP server, this repository treats agent capabilities as versioned infrastructure. An agent can resolve the user's request by task, role, and keywords, then progressively load only the relevant skill instructions or MCP connection guidance.
This is not the theoretical optimum for every setup. A carefully hand-tuned agent with the smallest possible set of skills and MCP servers for one specific task will usually be simpler, faster, and easier to reason about. The registry is meant for the more common middle ground: teams and users who run many agents across many tasks and need something better than enabling a broad list of skills and MCP servers for every agent "just in case". Many agent tools make that broad attachment path easy; this project explores a more explicit routing layer between fully minimal per-agent configuration and everything-enabled-by-default capability sprawl.
You can also use the registry as a starting point for minimal static setups. A user can ask their AI assistant to inspect the shared registry and assemble the smallest useful set of skills and MCP servers for a specific role, task, or agent configuration. In that mode, the registry is not the runtime layer itself; it is the common source of reviewed capabilities from which task-specific agent setups are derived.
The goal is to make AI-agent capability management more modular, reproducible, and reviewable:
discover and curate skills, MCP servers, workflows, and agent integrations from upstream sources;
organize capabilities by tasks, roles, and keywords for runtime routing;
keep enabled capabilities explicit in Git rather than hidden in local agent state;
separate trusted, reviewed, and candidate capabilities;
prefer safer MCP runtimes such as hosted endpoints or constrained Docker containers;
let multi-agent setups share the same capability registry instead of duplicating prompt/tool configuration.
This repository turns AI-agent capability sprawl into a reproducible GitOps-style registry. It downloads upstream sources as pinned submodules, inventories every discovered skill and MCP entry, and generates routing catalogs plus symlink packs organized by roles, tasks, and keywords.
git clone --recurse-submodules < repo-url >
cd ai-capability-registry
./scripts/bootstrap.sh
For shared use across multiple repositories, keep this registry at ~/.ai-registry :
git clone --recurse-submodules < repo-url > ~ /.ai-registry
cd ~ /.ai-registry
./scripts/bootstrap.sh
The bootstrap script will:
Sync external submodules to correct commits
Validate registry configuration
Sync provider chunks under skill-catalog.d/ with discovered skills
Generate the workflow runtime instructions, routing index, and human-readable workflow catalog from registry/workflows.yaml
Generate the combined skills/ routing catalogs and symlink packs from skill-catalog.d/
Generate MCP routing catalogs from mcp-catalog.d/
Agents can use this registry in two ways: dynamic routing at runtime or a minimal static setup derived from the registry for a specific agent, role, or task.
Use AGENTS.full-registry.md.template to connect another repository to a shared registry stored at ~/.ai-registry with the recommended workflow-first setup:
cp ~ /.ai-registry/AGENTS.full-registry.md.template /path/to/project/AGENTS.md
The full-registry template is intentionally only a bootloader. It tells agents to read and follow capability-routing.md , start with workflows/workflow.md , use workflows/routing.md only when workflow selection is needed, load workflow-selected roles from roles/ , then let those roles route to skills and MCP only when needed. More specific local repository instructions still override shared guidance.
Root templates are available for narrower setups:
If the registry is installed somewhere else, edit the paths in the copied AGENTS.md .
Use the generated workflow runtime instructions when the agent should choose process guidance before loading concrete skills or MCP servers:
<path-to-registry>/workflows/workflow.md
For full workflow routing when no workflow scope is assigned, use:
<path-to-registry>/workflows/routing.md
For cross-repository use, prefer installing AGENTS.full-registry.md.template or AGENTS.workflows.md.template rather than adding partial workflow-only instructions.
Workflow routing should be read before skill routing. Workflows answer how to approach a task, which roles should participate, which stages and gates apply, and which artifacts are expected. Roles then decide which skills to load for their part of the workflow.
Use the generated skill runtime instructions when the agent should dynamically choose skills per request:
<path-to-registry>/skills/skills.md
For full role and task capability routing when no registry scope is assigned, use:
<path-to-registry>/skills/routing.md
For dynamic routing, agents should usually reach skills/skills.md through workflow-selected role prompts. Direct skill-only use is still supported through AGENTS.skills.md.template or an equivalent bootloader.
For direct agent configuration, attach one generated pack instead of the whole registry when you want a narrower static setup:
<path-to-registry>/skills/packs/roles/<role-id>/
<path-to-registry>/skills/packs/tasks/<task-id>/
<path-to-registry>/skills/packs/keywords/<keyword>/
skills/packs/all/ exists, but it is mainly useful for inspection or experiments. Prefer task, role, keyword, or custom minimal packs for real agents.
Use the generated MCP runtime instructions when the agent should choose MCP servers only when a task needs them:
<path-to-registry>/mcp/mcp.md
For full role and task capability routing when no registry scope is assigned, use:
<path-to-registry>/mcp/routing.md
For dynamic routing, agents should usually reach mcp/mcp.md through workflow-selected role prompts or workflow stage guidance. Direct MCP-only use is still supported through AGENTS.mcp.md.template or an equivalent bootloader.
For direct agent configuration, adapt the generated connection metadata to your agent's MCP config schema:
<path-to-registry>/mcp/servers/<server-id>/connection.json
<path-to-registry>/mcp/servers/<server-id>/SKILL.md
Hosted MCP servers should follow mcp/web.md . Docker MCP servers should follow mcp/docker.md and must avoid privileged containers, host networking, Docker socket mounts, unrestricted host mounts, SSH key mounts, and cloud credential directory mounts.
Do not edit generated files directly. Edit source chunks, then regenerate.
For skills, edit only enabled and keywords in skill-catalog.d/*.yaml . New discovered skills are added disabled by default, and removed upstream skills remain with exists: false .
For MCP, keep hand-curated entries in mcp-catalog.d/manual.yml . Auto-discovered entries from official-mcp-registry.yml and docker-mcp-registry.yml should stay disabled until reviewed, assigned useful keywords, and explicitly promoted with enabled: true .
Important MCP fields are enabled , exists , trust , runtime , transport.type , default_mode , and keywords .
Command
Use
./scripts/bootstrap.sh
Full sync, validation, workflow generation, skills generation, and MCP generation
./scripts/generate-workflows.py
Regenerate workflows/workflow.md , workflows/routing.md , workflows-catalog.md , and workflow catalogs from registry/workflows.yaml
./scripts/discover-skills.py
Regenerate skill catalogs and packs from skill-catalog.d/
./scripts/discover-mcp.py
Import disabled candidate MCP entries from configured upstream sources
./scripts/generate-mcp.py
Regenerate mcp/ and mcp-catalog.md from mcp-catalog.d/
./scripts/validate-registry.py
Validate registry YAML structure and policy constraints
Trust And Policy
Trust levels: trusted for official or reputable security-reviewed sources, reviewed for manually reviewed known maintainers, and candidate for discovery-only entries.
Use trusted or reviewed capabilities by default.
Keep discovered candidates disabled until reviewed.
Prefer hosted MCP or constrained Docker MCP.
Deny direct node , npx , python , pip , and uvx MCP runners.
Never use curl | sh , privileged Docker, host networking, Docker socket mounts, or unrestricted host mounts.
Load skills and MCP servers only when they are relevant to the current task.
Trusted skill sources currently include Anthropic skills and knowledge-work plugins, OpenAI skills, Vercel Agent Skills, the Agent Skills specification, Superpowers development methodology skills, Trail of Bits security workflows, and Kilo Marketplace skills.
Curated AI Workflows, Skills & MCP

[truncated]
