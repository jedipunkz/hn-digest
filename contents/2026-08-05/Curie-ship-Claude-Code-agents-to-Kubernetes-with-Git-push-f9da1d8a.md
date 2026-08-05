---
source: "https://github.com/curie-eng/curie"
hn_url: "https://news.ycombinator.com/item?id=49183972"
title: "Curie – ship Claude Code agents to Kubernetes with Git push"
article_title: "GitHub - curie-eng/curie · GitHub"
author: "athusoo"
captured_at: "2026-08-05T15:08:27Z"
capture_tool: "hn-digest"
hn_id: 49183972
score: 1
comments: 0
posted_at: "2026-08-05T15:03:01Z"
tags:
  - hacker-news
  - translated
---

# Curie – ship Claude Code agents to Kubernetes with Git push

- HN: [49183972](https://news.ycombinator.com/item?id=49183972)
- Source: [github.com](https://github.com/curie-eng/curie)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T15:03:01Z

## Translation

タイトル: Curie – Git プッシュを使用して Claude Code エージェントを Kubernetes に出荷する
記事タイトル: GitHub - curie-eng/curie · GitHub
説明: GitHub でアカウントを作成して、curie-eng/curie の開発に貢献します。

記事本文:
GitHub - curie-eng/curie · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
キュリーエンジニア
/
キュリー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,648 コミット 1,648 コミット .github .github apps apps charts/ curie charts/ curie cli cli compose compose docs docs 例 例 otel otel パッケージ パッケージ プロトタイプ プロトタイプ リリース リリース ランナー

ランナー スクリプト スクリプト テスト テスト ツール ツール .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml .gitleaksignore .gitleaksignore .mcp.json .mcp.json AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md TRADEMARKS.md TRADEMARKS.md compose.dev.yaml compose.dev.yaml get-curie.sh get-curie.sh llms.txt llms.txt pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンソース (Apache 2.0)、実稼働 AI エージェント用の自己ホスト型配信プラットフォーム。接続する
Slack — 最初のチャネルであり、次は電子メールと Teams — クロード コード形式の作成者
プラグイン バンドル (スキル + ツール + MCP)、バージョン管理されたボット ID としてデプロイし、どこでも実行できます -
ラップトップ上の開発環境、または独自の Kubernetes クラスター上の運用環境で。
モデルを構成して、エージェントに Anthropic、OpenRouter、またはローカル モデルを指定できるようにします。
オラマ。トレース、評価、予算、Git 駆動のデプロイを無料で入手できます。 1 つの CLI curie がすべてを制御します。
それ。
ここは新しいですか？クイックスタートを使用すると、最初のエージェントの応答が数分以内に得られます。
エージェントがラップトップから離れると壊れる理由
ローカル環境と運用環境は通常異なります。つまり、Python のバージョンが異なったり、ツールが欠落していたり、
ある場所に存在し、別の場所には存在しない資格情報。キュリーは 1 人のメカニックでそのギャップを埋めます。
同じプラグイン バンドルは 3 つの階層に上がります。
スキルは、前にプラットフォームを持たず、単一のコンテナとして直接実行します。
local は、Docker Compose を介して完全なプラットフォームで実行します。
クラスターは、Kubernetes 上の同じ完全なプラットフォームを通じてそれを実行します。
環境が違う

ce はこれらの層を進む間にバグとして現れますが、驚くことではありません
ユーザーがヒットする - ローカルで高速に反復処理を行い、自信を持って出荷できるようになります。
3 つの層すべてで不変のバンドル スナップショットが実行されます。ローカルとクラスターはそのスナップショットを
バージョンを識別しますが、スキルはコンテンツのダイジェストによって識別します。スキルが高速ループになる理由は次のとおりです。
1 つのコマンドで、スナップショットを作業ディレクトリから直接パックしてブートします。
正面のプラットフォーム。
キュリーは 3 つの階層を登る際に環境保証を提供します。それは行為ではありません
保証: 運用トラフィックは依然としてテスト ケースとは異なる動作をする可能性があり、プラットフォームは異なる動作をする可能性があります。
正直にそうでないと約束します。
プラットフォームのアーキテクチャと各要素がどのように組み合わされるかについては、ARCHITECTURE.md を参照してください。
各層が実際に実行する内容については、以下のターゲット表を参照してください。
Slack にボットが必要なだけですか? docs/your-first-slack-agent.md
は短いパスです。最初に取得するもの、4 つのコマンド、および 6 つの間違い
人々に1時間の負担を与える。以下のウォークスルーは長いもので、次のことを教えています。
そのままパリティラダーを実行します。
Docker + Compose v2 : 開発スタックとローカル ランナー コンテナー用。
kubectl + helm : クラスターのインストール パスのみ。
Curie を使用した最初のエージェントの構築と展開
Anthropic API キーを取得して、一度エクスポートします。
エクスポート CURIE_CREDENTIALS=sk-ant-...
以下のすべての手順では、この同じ認証情報と同じバンドルを再利用します。 Docker がインストールされていない場合は、
Docker をインストールし、実行されていることを確認します。これはステップ 1 ～ 2 で必要です。
カール -fsSL https://raw.githubusercontent.com/curie-eng/curie/main/get-curie.sh |バッシュ
curie init my-agent && cd my-agent
足場ができたものを見てください。
木-a
。
§── .claude/
│ └── スキル/
│ └── キュリーを使用/
│ ━─ SKILL.md
§── .claude-plugin/
│ └── plugin.js

に
§── .gitignore
§── .mcp.json
§── AGENTS.md
§── evals/
│ └──cases.json
└── スキル/
└── my-agent/
━── SKILL.md
skill/my-agent/SKILL.md - エージェントの指示: エージェントが何を行うか、いつどのツールを使用するか。
.mcp.json - このエージェントが呼び出すことができる MCP サーバー (ツール)。
evals/cases.json - すべての層でこのエージェントの動作を評価する eval ケース。
AGENTS.md - このバンドルで動作するコーディング エージェントのルール。
.claude/skills/using-curie/SKILL.md - エージェントにキュリー ハーネスの操作方法を教える入門スキル (キュリー ガイドと同じ内容)。
これは、Claude Code プラグインの形式です。そのままの形式です。packages/plugin-format/README.md を参照してください。
Curie が検証する形状に対して。
キュリーのスキルアップ
キュリーのスキルメッセージ「こんにちは、いますか？」
実際の返信は、Slack もプラットフォームもまだありません。スキルアップは不変のスナップショットを実行します
バンドルなので、skills/my-agent/SKILL.md を編集した後、実行中のランナーにのみ変更が適用されます。
curie skill up --replace で再起動します。完了したら、次のコマンドを実行します
キュリーのスキルダウン
これは開発の最も速い内部ループであり、反復してスキルを構築することができます。
2. フルプラットフォームでもラップトップを使用可能
次に、エージェントを完全なバックエンドに接続して、メッセージが実際のキュー -> ワーカー -> サンドボックス -> 応答パスを実行するようにします。
キュリーローカルアップ
キュリーローカルデプロイ --plugin-dir 。 --slack-channel C0123ABCD --api-url http://localhost:28000
キュリーのローカルメッセージ「こんにちは、そこにいるの？」
その後、この会話スレッドを続けます
キュリー ローカル メッセージ --続き「2 + 2 とは何ですか?」
これは実際の Slack @mention と同じパスです
テイク - ライブで試す準備ができたら、docs/slack-local-runbook.md を参照してください。
http://localhost:28080/ ? API=1
会話全体、そのトレース、メトリクス、コストを確認します。同じコンソ

le はログも表示します。
承認とメモリ。このプラグインが運用環境の Kubernetes にデプロイされると、より関連性が高まります。
完了したら、次のコマンドを実行します
キュリーローカルダウン
3. 本物のKubernetes
最後に、バンドルを Kubernetes にデプロイします。
kubectl / helm をクラスターに指定します - k3s が永続的な推奨事項です。クラスターがない場合
便利です。minikube をインストールして次のコマンドを実行します
(k3s と minikube の間のトレードオフについては、docs/operations.md を参照してください):
curie クラスタアップ --allow-egress-host anthropic --set security.gvisor.mode=off
キュリークラスターデプロイ --plugin-dir 。 --repo < 所有者 > / < 名前 >
キュリークラスターメッセージ「こんにちは、いますか？」
--repo は、このエージェントを、ステップ 4 でこのバンドルをプッシュする GitHub リポジトリにバインドします。
リポジトリ バインディングによってエージェントと照合されるため、バインディングなしでデプロイされたエージェントのプッシュは
何も一致せず、無視されると応答され、バージョンにはなりません。そのバインディングは次の場合にのみ設定されます
エージェントは最初に作成され、後で変更することはできないため、自分の所有者/名前を置き換えてください
実行する前に: 間違ったやり方をすると、エージェントを削除して再作成することになります。参照
ライフサイクル動詞については、cli/README.md を参照してください。
--set security.gvisor.mode=off は、実際のモデルがインストールする gVisor の追加のカーネル分離をスキップします。
それ以外の場合は、minikube が必要であり、デフォルトでは minikube が同梱されません。runsc があるクラスターにドロップします。
インストールされています。
--allow-egress-host はモデル呼び出しを開きます。クラスターサンドボックスは
デフォルトではフェイルクローズされます (スキル/ローカルはそうではありません)。
クラスターの前提条件と完全な出力モデルについては、docs/operations.md を参照してください。
この最初のクラスターは使い捨てです。運用環境では、耐久性のあるデータをクラスターの外に保持します: セット
postgres.deploy: マネージド Postgres の場合は false、minio.deploy: S3 互換オブジェクトの場合は false
保存し、その cr を提供します

既存の Kubernetes Secret を通じて edentials を取得します。チャートはそのチャートをリダイレクトします
消費者をそれらのストアに誘導します (バンドルの取得も含む)。を参照してください。
チャート/キュリー実稼働ストレージ構成
完全なバッキング ストア設定については、
その後、この会話スレッドを続けます
キュリークラスターメッセージ --続き「2 + 2 とは何ですか?」
完了したら、次のコマンドを実行します
キュリークラスターダウン --はい
4. 発送します: CI/CD
ステップ 3 で示したように、Git-flow には --repo で作成されたエージェントが必要です。クラスターを次の時点で破棄した場合
ステップ 3 が終了したら、それを元に戻し、プッシュする前に同じクラスター展開行を再実行します。
バインディングはリリースのデータベース内に存在します。残っているのは、Curie API を GitHub に公開して配線することです
Webhook とそのシークレット ( docs/operations.md を参照)、次に:
gitプッシュオリジン開発
すべてのプッシュは不変のバージョン管理されたバンドルとして保存され、開発ボットの下にデプロイされます。
自動的に。 prod にマージすると、リビルドではなく同じバージョンが昇格されるため、常に
何がライブであるかを正確に把握し、任意のバージョンにロールバックできます。
Slack イベントループ、キュー、書き込み用のサンドボックス配管はありません。同じバンドルを要求しました
どこか大きな場所に走らせてから、自動的に出荷するように指示しました。
実現する準備はできていますか?接続するには docs/slack-local-runbook.md を参照してください。
このバンドルを実際の Slack ワークスペースに組み込みます。
これが dev または prod で公開されたら、チームメイトと同様に Slack でボットを @メンションします - を参照してください。
実際のワークスペースをデプロイされたリリースに接続するための apps/dispatcher/README.md の Runbook。
オフライン --fake-model パスについては、QUICKSTART.md を参照してください。
サンプル/バンドル、およびソースからの Curie のビルド。
環境に関わるすべての CLI コマンドは、ターゲット名詞を
中央: スキル 、 ローカル 、または クラスター 。あなたのニーズに応える最も軽いものを選んでください
質問。 curie init は例外です。これはディスク上でプラグイン バンドルをスキャフォールディングします。
環境をターゲットにしません。の

3 つの目標のポイントは同じです
プラグインバンドル形式であり、同じ evals/cases.json がすべてのプラグインで実行されるため、
スキルの促進 → ローカル → クラスターは 3 つの個別のラダーではなく、パリティ ラダーです
セットアップ。ラップトップではパスし、クラスターでは失敗する評価は次のようになります。
ノイズではなく信号。各ターゲットの eval コマンドはそれに沿って文書化されています
cli/README.md にあります。
この表には対象となる動詞がリストされていますが、ターゲットが持つすべての動詞ではありません。
カルテット アップ / ダウン / ステータス / メッセージは 3 つのターゲットすべてにあり、同様に
eval 、ローカル/クラスターの追加デプロイ中。参照
各ターゲットの動詞については、cli/README.md を参照してください。パリティという意味はありません
すべての機能がすべての層で実装されます。すべての動詞がすべての層で応答されます。
サポートされていない概念が決定論的な理由と
で定義されている代替案
ADR 0041 。
重要な違い: スキルはランナーのみのループです - それが起動します
ランナー コンテナだけであり、何もせずに ACI HTTP サーフェスと直接通信します。
正面のプラットフォーム。ローカルとクラスターは完全なプラットフォーム (キュー、
ワーカー、サンドボックス) を同一のランナーと ACI の前に置きます。に関するメッセージ
したがって、どちらも実際の Slack メンションがたどる道をたどります。
ターゲットごとの完全なコマンド リファレンスについては、cli/README.md を参照してください。
(スキル、
地元の、
クラスター)、
ローカル Slack に接続するための docs/slack-local-runbook.md、および
d

[切り捨てられた]

## Original Extract

Contribute to curie-eng/curie development by creating an account on GitHub.

GitHub - curie-eng/curie · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
curie-eng
/
curie
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,648 Commits 1,648 Commits .github .github apps apps charts/ curie charts/ curie cli cli compose compose docs docs examples examples otel otel packages packages prototypes prototypes release release runner runner scripts scripts tests tests tools tools .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml .gitleaksignore .gitleaksignore .mcp.json .mcp.json AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md TRADEMARKS.md TRADEMARKS.md compose.dev.yaml compose.dev.yaml get-curie.sh get-curie.sh llms.txt llms.txt pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Open-source (Apache 2.0), self-hostable delivery platform for production AI agents. Connect
Slack — the first channel it speaks, with email and Teams next — author a Claude-Code-format
plugin bundle (skills + tools + MCP), deploy it as a versioned bot identity and run it anywhere -
in your development environment on your laptop or in production on your own Kubernetes cluster.
Configure your model, so you can point an agent at Anthropic, OpenRouter, or a local model through
Ollama. Get traces, evals, budgets, and git-driven deploys for free. One CLI, curie , drives all of
it.
New here? Quickstart gets you a first agent reply in a few minutes.
Why your agent breaks when it leaves your laptop
Local and production environments are usually different - a different Python version, a missing tool,
a credential that exists in one place and not the other. Curie closes that gap with one mechanic:
the same plugin bundle climbs three tiers.
skill runs it directly, as a single container, no platform in front.
local runs it through the full platform via Docker Compose.
cluster runs it through that same full platform on Kubernetes.
An environment difference then shows up as a bug while progressing through these tiers, not a surprise
your users hit - letting you iterate fast locally and ship with confidence.
All three tiers run an immutable bundle snapshot: local and cluster assign that snapshot a
version, while skill identifies it by its content digest. What makes skill the fast loop is that
it packs and boots that snapshot straight from your working directory in one command, with no
platform in front.
Curie provides an environment guarantee while climbing the three tiers. It is not a behavior
guarantee: production traffic can still behave differently than your test cases, and no platform can
honestly promise otherwise.
See ARCHITECTURE.md for the platform architecture and how the pieces fit together.
See the target table below for what each tier actually runs.
Just want a bot in Slack? docs/your-first-slack-agent.md
is the short path: what to get first, four commands, and the six mistakes that
cost people an hour. The walkthrough below is the longer one, and teaches the
parity ladder as it goes.
Docker + Compose v2 : for the dev stack and the local runner container.
kubectl + helm : only for the cluster-install path.
Building and deploying your first agent with Curie
Get an Anthropic API key and export it once:
export CURIE_CREDENTIALS=sk-ant-...
Every step below reuses this same credential and the same bundle. If you don't have Docker installed,
install Docker and make sure it's running - it is needed for steps 1-2.
curl -fsSL https://raw.githubusercontent.com/curie-eng/curie/main/get-curie.sh | bash
curie init my-agent && cd my-agent
Take a look at what got scaffolded:
tree -a
.
├── .claude/
│ └── skills/
│ └── using-curie/
│ └── SKILL.md
├── .claude-plugin/
│ └── plugin.json
├── .gitignore
├── .mcp.json
├── AGENTS.md
├── evals/
│ └── cases.json
└── skills/
└── my-agent/
└── SKILL.md
skills/my-agent/SKILL.md - the agent's instructions: what it does, and when to use which tool.
.mcp.json - the MCP servers (tools) this agent can call.
evals/cases.json - the eval cases that grade this agent's behavior, at every tier.
AGENTS.md - the rules for a coding agent working in this bundle.
.claude/skills/using-curie/SKILL.md - a primer skill that teaches the agent to drive the Curie harness (same content as curie guide ).
This is the Claude Code plugin format, verbatim - see packages/plugin-format/README.md
for the shape Curie validates against.
curie skill up
curie skill message " hello, are you there? "
A real reply streams back: no Slack, no platform yet. skill up runs an immutable snapshot of the
bundle, so after you edit skills/my-agent/SKILL.md the change only reaches a running runner once you
restart it with curie skill up --replace . When done run the following command
curie skill down
This is the fastest inner loop of development and enables you to iterate and build your skills.
2. Full platform, still your laptop
Next hook up the agent into the full backend so that the message runs the real queue -> worker -> sandbox -> reply path.
curie local up
curie local deploy --plugin-dir . --slack-channel C0123ABCD --api-url http://localhost:28000
curie local message " hello, are you there? "
Then continue this conversation thread
curie local message --continue " what's 2 + 2? "
This is the same path a real Slack @mention
takes - see docs/slack-local-runbook.md when you're ready to try it live.
http://localhost:28080/ ? api=1
to see the whole conversation, its traces, metrics, and cost. The same console also surfaces logs,
approvals, and memory, which get more relevant once this plugin is deployed on Kubernetes in production.
When done run the following command
curie local down
3. Real Kubernetes
Finally, deploy the bundle on Kubernetes.
Point kubectl / helm at a cluster - k3s is the lasting recommendation; if you don't have a cluster
handy, install minikube and run the following command
(See docs/operations.md for the tradeoffs between k3s and minikube):
curie cluster up --allow-egress-host anthropic --set security.gvisor.mode=off
curie cluster deploy --plugin-dir . --repo < owner > / < name >
curie cluster message " hello, are you there? "
--repo binds this agent to the GitHub repository you will push this bundle to in step 4. A push
is matched to an agent by its repo binding, so a push for an agent deployed without the binding
matches nothing and is answered ignored , never becoming a version. That binding is set only when
the agent is first created and cannot be changed afterwards, so substitute your own owner/name
before running it: getting it wrong means deleting the agent and recreating it. See
cli/README.md for the lifecycle verbs.
--set security.gvisor.mode=off skips gVisor's extra kernel isolation, which a real-model install
otherwise requires and minikube doesn't ship by default - drop it on a cluster that has runsc
installed.
--allow-egress-host opens the model call; a credential alone doesn't, since the cluster sandbox is
fail-closed by default (skill/local aren't).
See docs/operations.md for cluster prerequisites and the full egress model.
This first cluster is disposable. For production, keep durable data outside the cluster: set
postgres.deploy: false for managed Postgres and minio.deploy: false for an S3 compatible object
store, then supply their credentials through existing Kubernetes Secrets. The chart redirects its
consumers to those stores, including bundle fetches. See the
charts/curie production storage configuration
for the complete backing store settings.
Then continue this conversation thread
curie cluster message --continue " what's 2 + 2? "
When done run the following command
curie cluster down --yes
4. Ship it: your CI/CD
Git-flow needs the agent created with --repo , as step 3 showed. If you tore the cluster down at the
end of step 3, bring it back up and re-run that same cluster deploy line before pushing, because the
binding lives in the release's database. What is left is exposing the Curie API to GitHub and wiring
the webhook and its secret (see docs/operations.md ), then:
git push origin dev
Every push is stored as an immutable, versioned bundle and deployed under your dev bot
automatically. Merging to prod promotes that same version, not a rebuild, so you always
know exactly what's live and can roll back to any version.
No Slack event loop, no queue, no sandbox plumbing to write. You asked the same bundle
to run somewhere bigger, then told it to ship itself.
Ready to make it real? See docs/slack-local-runbook.md to wire
this bundle into an actual Slack workspace.
Once this is live in dev or prod , @mention the bot in Slack like any teammate - see
apps/dispatcher/README.md 's runbook for connecting a real workspace to a deployed release.
See QUICKSTART.md for the offline --fake-model path,
the examples/ bundles, and building Curie from source.
Every CLI command that touches an environment takes a target noun in the
middle: skill , local , or cluster . Pick the lightest one that answers your
question. curie init is the exception: it scaffolds a plugin bundle on disk and
targets no environment. The point of the three targets is that the same
plugin bundle format and the same evals/cases.json run across all of them, so
promoting skill → local → cluster is a parity ladder, not three separate
setups. An eval that passes on your laptop and fails on the cluster is
signal, not noise; each target's eval command is documented alongside it
in cli/README.md .
The table lists the verbs it covers, not every verb a target has: the universal
quartet up / down / status / message is on all three targets, and so is
eval , while local / cluster add deploy . See
cli/README.md for each target's verbs. Parity does not mean
every capability is implemented at every tier: every verb is answered at every
tier, with unsupported concepts returning a deterministic reason and an
alternative, as defined in
ADR 0041 .
The distinction that matters: skill is the runner-only loop — it boots
just the runner container and talks straight to its ACI HTTP surface with no
platform in front. local and cluster put the full platform (queue,
worker, sandbox) in front of the identical runner and ACI. A message on
either therefore walks the same path a real Slack mention would take.
See cli/README.md for the full command reference per target
( skill ,
local ,
cluster ),
docs/slack-local-runbook.md for connecting local Slack, and
d

[truncated]
