---
source: "https://github.com/valmishq/valmis"
hn_url: "https://news.ycombinator.com/item?id=48761096"
title: "Show HN: I built an open-source alternative to Claude Cowork"
article_title: "GitHub - valmishq/valmis: The AI Agent designed for work with security in mind. · GitHub"
author: "wayneshng"
captured_at: "2026-07-02T13:30:52Z"
capture_tool: "hn-digest"
hn_id: 48761096
score: 1
comments: 0
posted_at: "2026-07-02T13:17:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I built an open-source alternative to Claude Cowork

- HN: [48761096](https://news.ycombinator.com/item?id=48761096)
- Source: [github.com](https://github.com/valmishq/valmis)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T13:17:49Z

## Translation

タイトル: Show HN: Claude Cowork に代わるオープンソースの代替品を構築しました
記事のタイトル: GitHub - valmishq/valmis: セキュリティを念頭に置いて動作するように設計された AI エージェント。 · GitHub
説明: セキュリティを念頭に置いて動作するように設計された AI エージェント。 - ヴァルミシュク/ヴァルミス
HN テキスト: HN さん、数か月前、私は人気のある AI エージェント OpenClaw を使用して作業の一部を自動化しようとしましたが、それを API やサードパーティ サービスと安全に連携させることがいかに難しいかにすぐに気づきました。これは多くの仕事関連のタスクに不可欠です。その後、OpenClaw はむしろ個人的なアシスタントであり、同僚として実際の作業を行うように設計されていないことに気づきました。そこで私は、セキュリティを最優先にして、100 以上のアプリやサービスで動作する OpenClaw の代替となる Valmis の構築を開始しました。 Valmis は、プロキシ システムを設計することでセキュリティ問題に対処しています。Dockerized エージェント ランタイムは、関連する資格情報 ID を提供することによってホスト マシンに API リクエストを行うことのみを要求できます。次に、ホストは実際のリクエストを作成し、JSON データをエージェント ランタイムに返します。この設計では、エージェント コンテナの動作中に、そのインターネット アクセスをオフにすることもできます。当社のプロキシ システムは現在、すべての Google Workspace アプリ、Slack、Notion、HubSpot、Salesforce、Figma を含む 100 以上のビジネスと生産性の統合をサポートしています。 Valmis の最も優れた機能の 1 つは、自動化されたワークフローです。ワークフロー ビルダーを使用して、複数ステップのワークフローを自動化できます。各ワークフローは cron、webhook、アプリ イベントによってトリガーでき、条件とループをサポートしています。コメント欄で質問があれば喜んでお答えします。

記事本文:
GitHub - valmishq/valmis: セキュリティを念頭に置いて動作するように設計された AI エージェント。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ヴァルミシュク
/
ヴァルミス
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
94 コミット 94 コミット .github/ ワークフロー .github/ ワークフロー アプリ アプリ docker/ ブラウザ docker/ ブラウザ docs docs パッケージ パッケージ 静的/ スクリーンショット 静的/ スクリーンショット .dockerignore .dockerignore .env.example .env.example .env.min.example .env.min.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc Dockerfile Dockerfile ライセンス ライセンス README.md README.md docker-compose.build.yml docker-compose.build.yml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh package.json package.json pm2.config.cjs pm2.config.cjs pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.jsonturbo.jsonturbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Valmis - セキュリティを念頭に置き、仕事用に設計された OpenClaw の代替品
今日はProduct Huntをライブで開催します！
Valmis は本日 Product Hunt で開始されます。このプロジェクトが気に入っていただけましたら、ぜひサポートをお願いいたします。
Valmis は、本番作業用の AI エージェントを配布するためのクラウドベースのアプリケーションです。これにより、100 以上のビジネスと生産性の統合と対話できるエージェントのフリートを構築できます。このシステムはセキュリティを念頭に置いて設計されており、エージェントは隔離されたコンテナ内で実行されます。つまり、AI が API 資格情報やホスト ファイルにアクセスすることはありません。
Valmis は、AI を使用してワークフローを自動化するように設計されています。チャット インターフェイスを使用してエージェントと対話したり、複数ステップのワークフローを作成し、cron、Webhook、またはアプリ イベント (新しいメール、フォームの送信など) でトリガーするようエージェントに依頼したりできます。
OpenClaw はパーソナル アシスタントを作成するための優れたツールですが、仕事用ではありません。エージェントは資格情報をメモリに保存するため、最大の懸念はセキュリティです。

プレーン テキストとして認証情報を送信し、場合によっては LLM プロバイダーに直接認証情報を送信します。
Valmis は、プロキシ システムを設計することでこの問題に対処しています。Dockerized エージェント ランタイムは、関連する資格情報 ID を提供することによってホスト マシンに API リクエストを行うことのみを要求できます。次に、ホストは実際のリクエストを作成し、JSON データをエージェント ランタイムに返します。 LLM API 呼び出し自体もプロキシを使用して行われます。この設計では、エージェント コンテナの動作中に、そのインターネット アクセスをオフにすることもできます。
各エージェントには独自のファイル システムがあり、ホスト マシンや他のエージェントから完全に分離されています。
エージェントは、アプリに (安全に) アクセスできる場合にのみ作業できます。当社のプロキシ システムは現在、すべての Google Workspace アプリ、Slack、Notion、Hubspot、Salesforce、Figma を含む 100 以上のビジネスと生産性の統合をサポートしています。現在サポートされているすべてのアプリについては、統合フォルダーを参照してください。作成した各エージェントには、特定 (またはすべて) の資格情報へのアクセスを割り当てることができ、この境界はコード レベルで厳密に遵守されます。その後、エージェントと会話して特定のタスクを完了すると、エージェントはホスト マシンへのプロキシ リクエストを作成して実際にリクエストを送信します。
valmis-github-recording-sm.mp4
Youtube でイントロを見る
最後に、エージェントに作業を手動で依頼することは必ずしも必要ではありません。自動化機能を使用して、複数ステップのワークフローを自動化できます。各ワークフローは cron、webhook、アプリ イベントによってトリガーでき、条件とループをサポートします。ワークフロー ビルダー UI を使用して複数ステップのワークフローを作成することも、説明を提供してエージェントにワークフローの作成を依頼することもできます。
では、なぜこのプロジェクトが Valmis と呼ばれているのかということになります。 Valmis はエストニア語で「完了」または「完了」を意味します（フィンランド語でも同様）。これは、プロジェクトがインスピレーションを受けてデザインされたためです。

ヨーロッパのデジタル国家であるエストニアで開発されました。また、アイスランドの TLD を使用するドメイン名 valm.is も持っているため、このプロジェクトはかなり北欧的です。 (エストニアが北欧かどうかについて議論する問題を開かないでください:))
独立して動作する、または連携して動作するエージェントのフリートを作成できます。各エージェントには、異なる資格情報、異なるスキル、および異なる知識ベースを割り当てることができます。また、エージェントごとに異なる LLM プロバイダーを割り当てて、重要性の低いミッションをより安価なモデルで実行することもできます。
適切な権限が付与されている場合、一部のエージェントはチーム リーダーとして機能し、他のエージェントのワークフローを管理する権限を持ち、人間のユーザーが最終的に制御する意思決定ツリーを形成できます。
ワークフロー ビルダー キャンバスを使用して、エージェントが自動的に実行する複数ステップのワークフローを作成できます。これは、ワークフローが毎日繰り返される場合、または特定のイベント (フォーム送信、Slack へのメンションなど) によってトリガーされる場合に特に便利です。
データのセキュリティと制御を強化するために、エージェントが各ステップで使用できる資格情報とツールを制限できます。より効率的なデータ マッピングのために、各ステップの出力のスキーマを定義することもできます。ワークフローに条件とループを追加できます。条件は、スマート (人間の言語を使用して条件を記述し、AI が条件が満たされるかどうかを判断する) または厳密 (プログラミング ロジックを使用して値を厳密に比較する) で指定できます。
エージェントにはセッション間の記憶がある
各エージェントには独自の記憶システムがあり、エピソード的 (何が起こったのか)、意味的 (永続的な事実)、手続き的 (ルールと制約)、作業的 (短期間のコンテキスト) の 4 つのカテゴリに分類されます。これは認知記憶研究からインスピレーションを得たメカニズムです。
エージェントは、覚えておくべきことを伝えたときや、次のようなときに自動的に記憶を書き込むことができます。

将来役立つかもしれないものを発見します。エージェントのメモリはセッション間で持続するため、エージェントと対話すればするほど、エージェントはワークフローをより深く理解できるようになります。セッションが終了すると、エージェントは学習した内容を自動的に抽出して、次の会話をよりスマートに開始できるようにします。エージェントにメモリ項目を変更または削除するように指示することもできます。
Valmis システムは pgvector を使用してメモリを保存および取得します。各メモリ項目は埋め込みモデルを使用して埋め込まれ、検索はテキスト埋め込みを使用して意味論的に行われます。
プロキシを使用してツールに安全に接続する
API キーまたは Oauth2 認証を使用して、100 を超えるビジネス アプリや生産性アプリに接続できます。資格情報は AES-256-GCM で暗号化され、データベースに保存されます。 AI エージェントは資格情報自体にアクセスすることはありませんが、代わりにプロキシを使用してホスト マシン経由でこれらの API を呼び出します。理論的には、インターネットにアクセスできないようにエージェント ランタイムを構成することもできますが、それでも動作します。
ここでは、すでにサポートされているアプリのいくつかのプレビューを示します。
...その他、QuickBooks、Xero、Stripe、WooCommerce、BigCommerce、Intercom、Zendesk、
Freshdesk、Pipedrive、ActiveCampaign、Klaviyo、SendGrid、Twilio、Calendly、Cal.com、GitHub、Jira、
Confluence、ClickUp、monday.com、Todoist、Contentful、Webflow、WordPress、Ghost、Miro、Canva、
イレブンラボ、Pinecone、および汎用 HTTP コネクタ (ベーシック、ベアラー、ヘッダー、およびクエリ認証)
リストにないもの。すべての統合は 1 つの YAML ファイルです。
Packages/utils/src/integrations/settings/ なので、
カタログは簡単に拡張できます。
ブラウザの自動化が簡単に
有効にすると、エージェントはヘッドレス ブラウザの操作、ナビゲーション、フォームへの入力、クリック、ページの読み取り、スクリーンショットの撮影を行うことができます。ブラウザもホスト マシンによって管理されるため、エージェントは対話します。

プロキシを使用してそれらと連携します。エージェントは自分の閲覧履歴とセッション Cookie にアクセスでき、手動でリセットまたは管理できます。
ループ内の人間: 重要な決定を下す必要があるときはいつでも、エージェントは一時停止し、一連のオプションを人間に尋ねます。
任意の LLM プロバイダーを使用する: 任意の LLM プロバイダーに接続し、チャットまたは埋め込みモデルを柔軟に使用できます。当社はすでに 20 プロバイダーの 200 近くのモデルをサポートしています (OpenAI、
Anthropic、Google、Mistral、Cohere など)、OpenRouter を使用してより多くの選択肢を提供することもできます。
ナレッジ ベース: Google Drive、Dropbox、Notion を使用してエンタープライズ ナレッジ ベースに接続するか、単にファイルをアップロードします。知識ベース ファイルはエージェントの記憶としても処理され、知識をより迅速に思い出すことができます。
スキル システム: Github からサードパーティのスキルをインストールしたり、対話しながら学習する独自の自己進化スキルを作成したりできます。
Valmis はおそらく、正当な手で本物のチェスをプレイできる最初の AI エージェントです。 LLM はチェスが下手で、常に幻覚を見ていることで有名です。そこで、私たちはチェスエンジンと呼ばれるツールをエージェントに追加しました。これは基本的に、エージェントが手を生成するためにテキスト生成に依存するのではなく、組み込みの軽量チェスエンジンの計算に厳密に基づいて各手を生成することを要求します。そして、AI は優れた (そしてスポーツマンらしい!) チェスプレイヤーになることができます。
システム全体は単一の Docker Compose ファイル、つまりアプリ (フロントエンドは 3000、バックエンドは 4000) から実行されます。
pgvector 対応の PostgreSQL、およびエージェント ランタイム用の Docker ソケット プロキシ。
# すべてを設定したい場合は .env.example を使用してください
cp .env.min.example .env
# シークレットを入力します — 少なくとも次のとおりです。
# CREDENTIAL_ENCRYPTION_KEY、JWT_SECRET、PROXY_TOKEN_SECRET
# それぞれを次のように生成します。
# ノード -e "console.log(require('crypto').randomBytes(

32).toString('hex'))"
ドッカー構成 -d
次に、 http://localhost:3000 を開き、 /setup に最初の管理者ユーザーを作成します。
Valmis は、Apache License 2.0 に基づいてリリースされています。
AI エージェントは、セキュリティを念頭に置いて作業できるように設計されています。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
2
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The AI Agent designed for work with security in mind. - valmishq/valmis

Hey HN, A few months ago, I tried to automate some of my work with the popular AI agent OpenClaw, and then I quickly realized how difficult it is to get it to work with APIs and third-party services securely, which is essential for a lot of work-related tasks. Then I realized OpenClaw is more of a personal assistant and it was not designed to get actual work done as a coworker. So I started to build Valmis, an alternative to OpenClaw that works with more than 100 apps and services, with security being the priority. Valmis addresses the security issue by designing a proxy system: dockerized agent runtime can only request the host machine to make API requests by providing the relevant credential ID. The host then makes the actual request and returns the JSON data to the agent runtime. With this design, you can even turn off the internet access of the agent container while making it work. Our proxy system now supports 100+ business and productivity integrations, including all Google Workspace apps, Slack, Notion, HubSpot, Salesforce, and Figma. One of the coolest features of Valmis is the automated workflow. You can automate multi-step workflows using our workflow builder. Each workflow can be triggered by cron, webhooks, app events, and it supports conditions and loops. I'd be happy to answer any questions in the comment section.

GitHub - valmishq/valmis: The AI Agent designed for work with security in mind. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
valmishq
/
valmis
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
94 Commits 94 Commits .github/ workflows .github/ workflows apps apps docker/ browser docker/ browser docs docs packages packages statics/ screenshots statics/ screenshots .dockerignore .dockerignore .env.example .env.example .env.min.example .env.min.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc Dockerfile Dockerfile LICENSE LICENSE README.md README.md docker-compose.build.yml docker-compose.build.yml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh package.json package.json pm2.config.cjs pm2.config.cjs pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json turbo.json turbo.json View all files Repository files navigation
Valmis - The OpenClaw alternative designed for work, with security in mind
We are live on Product Hunt today!
Valmis is launching on Product Hunt today. If you like the project, we'd love your support.
Valmis is a cloud-based application for shipping AI agents for production work. It allows you to build a fleet of agents that can talk to 100+ business and productivity integrations. The system is designed with security in mind, and agents run in isolated containers, meaning AI never gets access to your API credentials or your host files.
Valmis is designed to automate workflows using AI. You can interact with your agent using the chat interface, or ask it to create multi-step workflows and trigger it with cron, webhook, or app events (new email, form submission, etc.).
OpenClaw is a great tool to create personal assistants, but it is not for work. The biggest concern is security, as agents store credentials in their memory as plain text and sometimes send credentials directly to LLM providers.
Valmis addresses this issue by designing a proxy system : dockerized agent runtime can only request the host machine to make API requests by providing the relevant credential ID. The host then makes the actual request and returns the JSON data back to the agent runtime. Even the LLM API calls themselves are made using proxy. With this design, you can even turn off the internet access of the agent container while making it work.
Each agent has its own file system and is completely isolated from the host machine or from other agents.
Agents can only work for you when they have access (safely) to your apps. Our proxy system now supports 100+ business and productivity integrations, including all Google Workspace apps, Slack, Notion, Hubspot, Salesforce, and Figma. See integrations folder for all currently supported apps. Each agent you create can be assigned access to specific (or all) credentials, and this boundary is strictly followed at the code level. You can then talk to the agent to complete certain tasks, and agents will formulate proxy requests to the host machine to actually send the requests.
valmis-github-recording-sm.mp4
Watch intro on Youtube
Finally, you don't always want to manually ask your agents to work. You can automate multi-step workflows using our automation feature. Each workflow can be triggered by cron, webhooks, app events and it supports conditions and loops. You can create multi-step workflows using our workflow builder UI, or you can simply ask your agent to create one by providing a description.
Then it comes to why the project is called Valmis. Valmis is an Estonian word that means "completed" or "done"(Same in Finnish). This is because the project was inspired and designed in Estonia, Europe's digital nation. It also has the domain name valm.is that uses the Icelandic TLD, so the project is pretty Nordic. (Please do not open issues discussing whether Estonia is Nordic :) )
You can create a fleet of agents that act independently or in collaboration. Each agent can be assigned different credentials, different skills, and different knowledge bases. You can also assign different LLM providers for each agent, making sure less critical missions are done by less expensive models.
When given the proper permission, some agents can act as a team lead and have the authority to manage the workflow of other agents, forming a decision tree that is ultimately controlled by the human user.
You can use our workflow builder canvas to create multi-step workflows that the agent will run automatically. This is especially useful when you have workflows that repeat every day or are triggered by specific events (form submission, Slack mentions etc.)
For better data security and control, you can limit the credentials and tools the agent can use in each step. You can also define the schema for the output of each step for more efficient data mapping. You can add conditions and loops to the workflow. Conditions can by smart (you describe the condition using human language and AI decides if the condition is met) or strict (compare values rigorously using programming logic).
Agents have cross-session memory
Each agent has its own memory system that is organized into four categories: episodic (what happened), semantic (durable facts), procedural (rules and constraints), and working (short-lived context). This is a mechanism inspired by cognitive-memory research.
Your agents are able to automatically write memory when you tell them anything worth remembering or when it discovers something that might be useful in the future. Agent memory is persistent across sessions, meaning the more you interact with the agent, the more it will know your workflow. When a session ends, the agent automatically distills what it learned so the next conversation starts smarter. You can also instruct the agent to modify or remove memory items.
Valmis system uses pgvector to store and fetch memories. Each memory item is embedded using embedding models, and the search is done semantically using text embedding.
Connect to your tools safely with proxies
You can connect to more than 100 business and productivity apps using an API key or Oauth2 authentication. The credentials are encrypted with AES-256-GCM and are stored in your database. AI agents never get access to the credentials themselves, but instead call these APIs through the host machine using a proxy. Theoretically, you can configure your agent runtime to have no access to the internet, and it will still work.
Here is a preview of some of the apps we already support.
...and many more, including QuickBooks, Xero, Stripe, WooCommerce, BigCommerce, Intercom, Zendesk,
Freshdesk, Pipedrive, ActiveCampaign, Klaviyo, SendGrid, Twilio, Calendly, Cal.com, GitHub, Jira,
Confluence, ClickUp, monday.com, Todoist, Contentful, Webflow, WordPress, Ghost, Miro, Canva,
ElevenLabs, Pinecone, and the generic HTTP connectors (Basic, Bearer, header, and query auth) for
anything not on the list. Every integration is one YAML file in
packages/utils/src/integrations/definitions/ , so the
catalog is easy to extend.
Browser automation made simple
When enabled, agents can operate a headless browser, navigate, fill forms, clicks, read pages, and take screenshots. Browsers are also managed by the host machine so agents interact with them using proxy. Agents have access to their own browsing history and session cookies, which you can manually reset or manage.
Human in the loop: Whenever there is a critical decision to make, the agent pauses and ask the human with a set of options.
Use any LLM provider: You can connect to any LLM provider and use their chat or embedding models flexibly. We already support nearly 200 models from 20 providers (OpenAI,
Anthropic, Google, Mistral, Cohere, and more), you can also use OpenRouter for more choices.
Knowledge base: Connect your enterprise knowledge base using Google Drive, Dropbox, Notion, or simply upload files. Knowledge base files are also processed as memories for agents to ensure quicker knowledge recall.
Skills system: You can install third-party skills from Github, or create your own self-evolving skills that learn from you as you interact with it.
Valmis is probably the first AI agent that is able to play real chess with legit moves. We all know LLMs are notoriously terrible at playing chess and always hallucinate moves . So we added a tool to the agent called chess-engine, which basically requires the agent not to rely on text generation to produce moves, but instead to produce each move strictly based on the calculation of a lightweight chess engine built in. And AI can be a great (and sportsmanlike!) chess player.
The whole system runs from a single Docker Compose file: the app (frontend on 3000, backend on 4000),
a pgvector-enabled PostgreSQL, and a Docker socket proxy for the agent runtime.
# Use .env.example if you want to configure everything
cp .env.min.example .env
# Fill in the secrets — at minimum:
# CREDENTIAL_ENCRYPTION_KEY, JWT_SECRET, PROXY_TOKEN_SECRET
# Generate each with:
# node -e "console.log(require('crypto').randomBytes(32).toString('hex'))"
docker compose up -d
Then open http://localhost:3000 and create your first admin user at /setup .
Valmis is released under the Apache License 2.0 .
The AI Agent designed for work with security in mind.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
2
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
