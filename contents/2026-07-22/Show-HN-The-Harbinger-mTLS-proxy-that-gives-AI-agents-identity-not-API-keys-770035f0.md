---
source: "https://github.com/n0tduck1e/theharbinger"
hn_url: "https://news.ycombinator.com/item?id=49012282"
title: "Show HN: The Harbinger- mTLS proxy that gives AI agents identity, not API keys"
article_title: "GitHub - n0tduck1e/theharbinger: Sane policies for insane agents · GitHub"
author: "not-duckie"
captured_at: "2026-07-22T20:07:23Z"
capture_tool: "hn-digest"
hn_id: 49012282
score: 3
comments: 1
posted_at: "2026-07-22T19:38:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: The Harbinger- mTLS proxy that gives AI agents identity, not API keys

- HN: [49012282](https://news.ycombinator.com/item?id=49012282)
- Source: [github.com](https://github.com/n0tduck1e/theharbinger)
- Score: 3
- Comments: 1
- Posted: 2026-07-22T19:38:16Z

## Translation

タイトル: HN を表示: API キーではなく AI エージェントに ID を与える Harbinger-mTLS プロキシ
記事のタイトル: GitHub - n0tduck1e/theharbinger: 非常識なエージェントのための健全なポリシー · GitHub
説明: 非常識なエージェントのための健全なポリシー。 GitHub でアカウントを作成して、n0tduck1e/theharbinger の開発に貢献してください。
HN テキスト: このモデルは他のツールが行っていることをより良くできると考えたので、これを構築しました。皆さんはどう思いますか？ Harbinger は、AI エージェントの送信トラフィックの前に位置する mTLS プロキシです。すべての AI エージェント、ボット、サービス アカウントは暗号化された ID を取得します。リクエストが行われるすべてのリクエストは、許可される前にポリシーに照らしてチェックされ、必要な実際のシークレットは独自の保管庫から取得され、リクエストが主張する場所に送信されることが確認された場合にのみ交換されます。エージェント自体が永続的な資格を保持することはありません。

記事本文:
GitHub - n0tduck1e/theharbinger: 非常識なエージェントのための健全なポリシー · GitHub
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
アラートを閉じる
{{ メッセージ }}
n0tduck1e
/
前兆
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

s
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット アセット 構成 構成 デプロイ/ docker デプロイ/ docker docs docs harbinger_admin harbinger_admin harbinger_agentd harbinger_agentd harbinger_gateway harbinger_gateway harbinger_ui harbinger_ui hbg_cli hbg_cli 内部 内部スキーマ スキーマ .dockerignore .dockerignore .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
非常識なエージェントのための健全な政策。
Harbinger は、非人間 ID (NHI) およびエージェント セキュリティ プラットフォームです。すべての AI エージェント、ボット、サービス アカウントは暗号化 ID (mTLS) を取得し、すべてのリクエストは許可される前にポリシーと照合してチェックされ、必要な実際の秘密はすべて独自の保管庫から取得され、エッジで交換されるため、エージェント自体が永続的な資格情報を保持することはありません。
このカテゴリは、人間以外のアイデンティティ セキュリティ全般 (API キー、サービス アカウント、ボット トークン) であり、AI エージェントはその最も鋭く急速に成長している例であり、AI 専用ツールではありません。
現在のほとんどの「AI エージェント セキュリティ」プラットフォームは、エージェントを書き換える必要がある SDK、使用方法を変更せずにどのような資格情報が存在するかを示す検出/ガバナンス ダッシュボード、またはある API キーから別の API キーに責任を移すだけのもののいずれかです。 Harbinger は別のアプローチをとります。
コード変更は不要: Harbinger は透過的な mTLS MITM プロキシです。エージェントのトラフィックをゲートウェイに向ければ、制御できないサードパーティのエージェント フレームワーク (今日のほとんどのエージェント コード) も含めてカバーされます。
秘密のアクセスは、手続き上だけでなく構造的にもゲートされています。資格情報を含む各リクエストは、次の 2 つの質問に対してチェックされます: エージェントはこれへのアクセスを許可されていますか?

の Web サイトにアクセスできますか? エージェントはこの Web サイトに対してこの資格情報を使用できますか?両方が true の場合にのみ、エージェントは資格情報を使用できます。
cd デプロイ/ドッカー
cp .env.example .env
# .env を編集し、実際の ADMIN_PASSWORD を設定します
docker compose ファイルを実行します。
docker 構成 --build
ビルドが完了すると、 pki-init は hbg_cli 、 harbinger_agentd 、および root_ca.crt をホスト上のdeploy/docker/client-bundle/に自動的にドロップします。エージェントを別のマシンで実行する場合は、最初にそのディレクトリ全体をそこにコピーします。
cd デプロイ/docker/クライアントバンドル
./hbg_cli cert install --cert root_ca.crt
./hbg_cli エージェントのインストール --agent my-agent --ip 10.0.0.5 --user admin
これにより登録リクエストが送信され、承認を待ちます。 Web UI (エージェント登録の下の http://localhost:8446 ) から承認するか、CLI を使用して、同じコマンドを再実行して登録を完了します。
./hbg_cli エージェントのインストール --agent my-agent
これにより、 my-agent.crt / .key が書き込まれ、ルート CA がキャッシュされ、agentd.yml が生成されます。これらはすべてデフォルトで ~/.config/harbinger_agentd/ の下にあります。次に、 harbinger_agentd の下でエージェントを実行し、ゲートウェイ経由でトラフィックをプロキシします。
harbinger_agentd run --agent my-agent -- < YOUR_AGENT_COMMAND >
# 例
harbinger_agentd run --agent my-agent --opencode
# ワークフローでない場合はエイリアスを作成する
alias opencode= " harbinger_agentd run --agent my-agent -- opencode "
エージェントのトラフィックにシークレットを挿入します (オプション)。
Web UI で、Vault Integration → Local に移動します。Secrets の下にシークレットを追加し、その下のセクションでターゲットを作成します。これにより、ポリシーが参照できる名前が、実際のシークレットが存在する場所と、そのシークレットの送信先が許可される場所にマップされます (local: プレフィックスが適用されます)。 CLI と同等のもの（スクリプトを作成する場合）:
hbg_cli admin target create --namelack-bot-token --secret-r

ef local:slack-bot-token --hostlack.com
ボールト側の設定 (ローカルではなく AWS を含む) については、docs/agent-secret-placeholders.md を参照してください。
次に、エージェントが現在実際のトークンを読み取っている場所に、代わりにプレースホルダーを置きます。
エクスポート SLACK_BOT_TOKEN= ' hbg-secret:slack-bot-token: '
Harbinger は、このターゲットを明示的に許可するポリシーを持つエージェントからのslack.comへのリクエストに対してのみ、実際の値を交換します。
バイナリ
何をするのか
ハービンジャー_管理者
コントロールプレーン。データベース、管理 CA、ユーザー/API キー認証、組織/チーム/RBAC、ポリシー/ボールト/ターゲット ストレージ、登録レビュー、トラフィック テレメトリ、およびゲートウェイへの署名付きエッジ構成配布を所有します。
ハービンジャー_ゲートウェイ
データプレーン。 mTLS MITM プロキシは、すべてのエージェントのトラフィックが通過します。エージェントの証明書を検証し、リクエストごとにポリシーを評価し、(構成されている場合は) ボールトからの実際のシークレットを置き換えます。登録時にエージェント証明書も発行します。
hbg_cli
コマンドラインクライアント。 PKI 生成、登録、ポリシー/ボールト/ターゲット管理、管理者ユーザー/組織/チーム管理のすべてを実行します。 PKI 操作を実行できる唯一のクライアント。
ハービンジャー_ui
PKI を除くすべての hbg_cli に代わるブラウザベースの代替手段。これは、独自のバックエンドを持たない単純な静的ファイル サーバーです。ブラウザは、CLI とまったく同じように、管理 API を直接呼び出します。
ハービンジャーエージェントd
オプションのラッパー デーモン: 既存のエージェント プロセスを子として実行し、そのトラフィックをゲートウェイ ( HTTPS_PROXY 経由) 経由で透過的にルーティングし、証明書の有効期限が切れる前に証明書を自動更新します。プロキシを指定できないエージェントに便利です。
管理サーバーとゲートウェイ サーバーは、ホスティングの柔軟性だけでなく攻撃対象領域の非対称性を考慮して、意図的に別個のプロセスになっています。ゲートウェイは、敵対的なエージェントのトラフィックとライブシークレットの取得に直面します。 admin は CA/ユーザー管理権限を保持しており、

エージェントのトラフィックに直接触れることはありません。
フローチャート LR
エージェント["エージェント / ボット<br/>(エージェント コード)"]
Sys[「あなたのシステム<br/>(Slack、GitHubなど)」]
管理者["Harbinger 管理者<br/>(コントロール プレーン)<br/>CA · DB · ポリシー/ボールト ストレージ<br/>登録レビュー · UI/CLI 認証"]
サブグラフ ゲートウェイ["ハービンジャー ゲートウェイ"]
TB方向
V["1. 証明書の検証"] --> P["2. ポリシーの確認"] --> I["3. シークレットの挿入"]
終わり
エージェント -- 「mTLS (エージェント証明書)」 --> ゲートウェイ
ゲートウェイ -- 「実際の API 資格情報<br/>エッジで置換」 --> Sys
管理者 -- 「署名付きエッジ構成<br/>(ポリシー、CRL、ターゲット)」 --> ゲートウェイ
読み込み中
エージェントには最初から実際の資格情報が与えられることはありません。これらはプレースホルダー ( hbg-secret:secret-mapping: ) で構成されており、ゲートウェイは、ポリシーがそのターゲットに対してそのエージェントを明示的に許可し、リクエストが登録された宛先であると主張しているだけでなく実際にその宛先に向かっていることを確認した後にのみ、実際の値を交換します。 2 番目のチェックが重要な理由を含む完全な仕組みについては、 docs/agent-secret-placeholders.md を参照してください (これは、プロンプト挿入されたエージェントが攻撃者が制御するホストにシークレットを漏洩するのを阻止するものです)。
すべてのエージェントは実際の mTLS 証明書を使用して登録します。
エージェントは独自のキーペアと CSR を生成し、それが属するユーザー アカウントに名前を付けて送信します。
そのユーザー (または管理者) がリクエストを確認し、承認します。
エージェントはゲートウェイへの登録を完了し、CSR 自身の署名と管理者の署名付き承認を検証し、実際の mTLS 証明書を発行します。
取り消しでは、更新された CRL が署名されたエッジ構成を介してすべてのゲートウェイにプッシュされるため、ローテーションされていない共有 API キーのように、アクセスが暗黙的にその有用性を超えて存続することはありません。
信頼は 4 層の CA 階層 (ルート CA -> 管理 CA / ゲートウェイ CA / エージェント CA) に分割されるため、1 つの CA が侵害されても他の CA に引き継がれることはありません。たとえば、l

漏洩したゲートウェイ CA はエージェント証明書を偽造できず、漏洩したエージェント CA は偽のゲートウェイを登録できません。
エージェント サンドボックスとしての harbinger_agentd: 現在、エージェントのネットワーク トラフィックをラップします。次に、エージェントがホスト上で実行しようとするコマンドもキャプチャし、危険なコマンドを実行前にブロックして、ポリシー適用を「呼び出せるもの」から「実行できるもの」に拡張します。
より多くのボールトプロバイダー: 現在、AWS とローカルシークレットがサポートされています。次に GCP Secret Manager と Bitwarden が続きます。
質問
ドクター
完全な REST API サーフェス (管理者 + ゲートウェイ) とは何ですか?
docs/api-reference.md
コードベースはどのように構成されていますか?またコードベースに追加するにはどうすればよいですか?
docs/architecture.md
シークレット注入用にエージェントを構成するにはどうすればよいですか?
docs/agent-secret-placeholders.md
Docker のデプロイメント
デプロイ/docker/README.md
ライセンス
Apache License、バージョン 2.0 に基づいてライセンスされています。
非常識なエージェントに対する健全なポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Sane policies for insane agents. Contribute to n0tduck1e/theharbinger development by creating an account on GitHub.

I built this as i thought this model could better to what other tools are doing. What do you guys think ? Harbinger is an mTLS proxy that sits in front of an AI agent's outbound traffic. Every AI agent, bot, and service account gets a cryptographic identity. Every request it makes is checked against policy before it's allowed, and any real secret it needs is pulled from your own vault and swapped in only when the request is verified to be going where it claims. The agent itself never holds a standing credential.

GitHub - n0tduck1e/theharbinger: Sane policies for insane agents · GitHub
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
n0tduck1e
/
theharbinger
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits assets assets configs configs deploy/ docker deploy/ docker docs docs harbinger_admin harbinger_admin harbinger_agentd harbinger_agentd harbinger_gateway harbinger_gateway harbinger_ui harbinger_ui hbg_cli hbg_cli internal internal schema schema .dockerignore .dockerignore .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Sane policies for insane agents.
Harbinger is a non-human identity (NHI) and agentic security platform. Every AI agent, bot, and service account gets a cryptographic identity (mTLS), every request it makes is checked against policy before it's allowed, and any real secret it needs is pulled from your own vault and swapped in at the edge, so the agent itself never holds a standing credential.
The category is non-human identity security in general (API keys, service accounts, bot tokens), with AI agents as the sharpest and fastest-growing example of it, not an AI-only tool.
Most "AI agent security" platforms today are either an SDK you have to rewrite your agent against, a discovery/governance dashboard that tells you what credentials exist without changing how they're used, or something that just shifts responsibility from one API key to another. Harbinger takes a different approach:
No code change required: Harbinger is a transparent mTLS MITM proxy. Point an agent's traffic at the gateway and it's covered, including third-party agent frameworks you don't control, which is most agent code today.
Secret access is structurally, not just procedurally, gated. Each request with credentials gets checked against two questions: is the agent allowed to access this website, and is the agent allowed to use this credential against this website? Only when both are true does the agent get to use the credential.
cd deploy/docker
cp .env.example .env
# edit .env and set a real ADMIN_PASSWORD
Run the docker compose file:
docker compose up --build
Once the build finishes, pki-init automatically drops hbg_cli , harbinger_agentd , and root_ca.crt into deploy/docker/client-bundle/ on the host. If the agent will run on a different machine, copy that whole directory there first.
cd deploy/docker/client-bundle
./hbg_cli cert install --cert root_ca.crt
./hbg_cli agent install --agent my-agent --ip 10.0.0.5 --user admin
This submits the enrollment request and waits for approval. Approve it from the web UI ( http://localhost:8446 , under Agent Enrollment) or use the CLI, then re-run the same command to complete enrollment:
./hbg_cli agent install --agent my-agent
This writes my-agent.crt / .key , caches the root CA, and generates agentd.yml , all under ~/.config/harbinger_agentd/ by default. Then run the agent under harbinger_agentd , which proxies its traffic through the gateway:
harbinger_agentd run --agent my-agent -- < YOUR_AGENT_COMMAND >
# example
harbinger_agentd run --agent my-agent -- opencode
# or create an alias if it's not a workflow
alias opencode= " harbinger_agentd run --agent my-agent -- opencode "
Inject a secret into the agent's traffic (optional).
In the web UI, go to Vault Integration → Local: add a secret under Secrets, then create a target in the section below it, this maps a name your policy can reference to where the real secret lives and which destination it's allowed to go to (the local: prefix is applied for you). CLI equivalent, if you'd rather script it:
hbg_cli admin target create --name slack-bot-token --secret-ref local:slack-bot-token --host slack.com
See docs/agent-secret-placeholders.md for the vault-side setup (including AWS instead of local).
Then wherever your agent currently reads its real token from, put the placeholder there instead:
export SLACK_BOT_TOKEN= ' hbg-secret:slack-bot-token: '
Harbinger swaps in the real value only for requests to slack.com from an agent whose policy explicitly allows this target.
Binary
What it does
harbinger_admin
Control plane. Owns the database, the Admin CA, user/API-key auth, org/team/RBAC, policy/vault/target storage, enrollment review, traffic telemetry, and signed edge-config distribution to gateways.
harbinger_gateway
Data plane. The mTLS MITM proxy every agent's traffic goes through: verifies the agent's certificate, evaluates policy per request, and (if configured) substitutes real secrets from a vault. Also issues agent certs at enrollment time.
hbg_cli
Command-line client. Does everything: PKI generation, enrollment, policy/vault/target management, admin user/org/team management. The only client that can do PKI operations.
harbinger_ui
Browser-based alternative to hbg_cli for everything except PKI. It's a plain static-file server with no backend of its own; the browser calls the admin API directly, exactly like the CLI does.
harbinger_agentd
Optional wrapper daemon: runs an existing agent process as a child, transparently routes its traffic through the gateway (via HTTPS_PROXY ), and auto-renews its certificate before it expires. Useful for agents you can't otherwise point at a proxy.
Admin and gateway servers are deliberately separate processes, not just for hosting flexibility but for attack-surface asymmetry. The gateway faces adversarial agent traffic and live secret-fetching; admin holds CA/user-management authority and never touches agent traffic directly.
flowchart LR
Agent["Agent / Bot<br/>(agent code)"]
Sys["Your systems<br/>(Slack, GitHub, etc.)"]
Admin["Harbinger Admin<br/>(control plane)<br/>CA · DB · policy/vault storage<br/>enrollment review · UI/CLI auth"]
subgraph Gateway["Harbinger Gateway"]
direction TB
V["1. verify cert"] --> P["2. check policy"] --> I["3. inject secret"]
end
Agent -- "mTLS (agent cert)" --> Gateway
Gateway -- "real API credential<br/>substituted at the edge" --> Sys
Admin -- "signed edge config<br/>(policy, CRL, targets)" --> Gateway
Loading
Agents are never given a real credential to begin with. They're configured with a placeholder ( hbg-secret:secret-mapping: ), and the gateway swaps in the real value only after policy explicitly authorizes that agent for that target and confirms the request is actually headed to the registered destination, not just claiming to be. See docs/agent-secret-placeholders.md for the full mechanics, including why that second check matters (it's what stops a prompt-injected agent from exfiltrating a secret to an attacker-controlled host).
Every agent enrolls with a real mTLS certificate:
The agent generates its own keypair + CSR and submits it, naming the user account it belongs to.
That user (or any admin) reviews and approves the request.
The agent completes enrollment with the gateway, which verifies the CSR's own signature plus the admin's signed approval, then issues a real mTLS certificate.
Revocation pushes an updated CRL to all gateways via signed edge config, so access doesn't silently outlive its usefulness the way an un-rotated shared API key does.
Trust is split across a four-tier CA hierarchy (Root CA -> Admin CA / Gateway CA / Agent CA), so a compromise of one CA doesn't hand over the others. For example, a leaked Gateway CA can't forge agent certificates, and a leaked Agent CA can't register a fake gateway.
harbinger_agentd as an agent sandbox: today it wraps an agent's network traffic; next it'll also capture the commands the agent tries to run on the host and block risky ones before they execute, extending policy enforcement from "what it can call" to "what it can do."
More vault providers: AWS and local secrets are supported today; GCP Secret Manager and Bitwarden are next.
Question
Doc
What's the full REST API surface (admin + gateway)?
docs/api-reference.md
How is the codebase organized, and how do I add to it?
docs/architecture.md
How do I configure an agent for secret injection?
docs/agent-secret-placeholders.md
Docker deployment
deploy/docker/README.md
License
Licensed under the Apache License, Version 2.0 .
Sane policies for insane agents
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
