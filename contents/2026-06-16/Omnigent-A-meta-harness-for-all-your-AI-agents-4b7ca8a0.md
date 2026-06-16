---
source: "https://github.com/omnigent-ai/omnigent"
hn_url: "https://news.ycombinator.com/item?id=48554929"
title: "Omnigent: A meta-harness for all your AI agents"
article_title: "GitHub - omnigent-ai/omnigent: A meta-harness for all your AI agents. Omnigent provides a common layer over Claude Code, Codex, Pi, and the agents you write yourself: swap or combine harnesses without rewriting, keep them in check with policies and sandboxing, and collaborate in real time on the sam\n[truncated]"
author: "lobo_tuerto"
captured_at: "2026-06-16T13:55:40Z"
capture_tool: "hn-digest"
hn_id: 48554929
score: 1
comments: 0
posted_at: "2026-06-16T13:24:13Z"
tags:
  - hacker-news
  - translated
---

# Omnigent: A meta-harness for all your AI agents

- HN: [48554929](https://news.ycombinator.com/item?id=48554929)
- Source: [github.com](https://github.com/omnigent-ai/omnigent)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:24:13Z

## Translation

タイトル: Omnigent: すべての AI エージェントのためのメタハーネス
記事のタイトル: GitHub -omnigent-ai/omnigent: すべての AI エージェントのためのメタハーネス。 Omnigent は、Claude Code、Codex、Pi、および自分で作成したエージェントに共通のレイヤーを提供します。書き換えることなくハーネスを交換または結合し、ポリシーやサンドボックスでハーネスを監視し、サムでリアルタイムに共同作業します。
[切り捨てられた]
説明: すべての AI エージェント用のメタハーネス。 Omnigent は、Claude Code、Codex、Pi、および自分で作成したエージェントに共通のレイヤーを提供します。書き換えることなくハーネスを交換または結合し、ポリシーやサンドボックスでハーネスを管理し、同じライブ セッション上でどのデバイスからでもリアルタイムで共同作業できます。
[切り捨てられた]

記事本文:
GitHub -omnigent-ai/omnigent: すべての AI エージェントのためのメタハーネス。 Omnigent は、Claude Code、Codex、Pi、および自分で作成したエージェントに共通のレイヤーを提供します。書き換えることなくハーネスを交換または結合し、ポリシーやサンドボックスでハーネスを管理し、同じライブ セッション上でどのデバイスからでもリアルタイムで共同作業できます。 · GitHub
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
そうだね

別のタブまたはウィンドウでかゆいアカウントを表示します。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
オムニジェントアイ
/
全能の
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
227 コミット 227 コミット .claude/ スキル .claude/ スキル .github .github ap-web ap-web デプロイ デプロイ デザイン デザイン 開発 dev ドキュメント ドキュメントの例 例 オムニジェント オムニジェント スクリプト スクリプト sdks sdks テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス通知 NOTICE README.md README.md SECURITY.md SECURITY.md openapi.json openapi.json pyproject.toml pyproject.toml pyrefly.toml pyrefly.toml Railway.toml Railway.toml render.yaml render.yaml setup.py setup.py uv.lock uv.lock uv.toml uv.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すべての AI エージェント向けのメタハーネス
Omnigent は、Claude Code、Codex、Cursor、Pi、および自分で作成したエージェントに共通のレイヤーを提供します。書き換えることなくハーネスを交換または結合し、ポリシーやサンドボックスでハーネスを管理し、同じライブ セッション上でどのデバイスからでもリアルタイムで共同作業できます。
omnigent.ai · ⬇️ macOS デスクトップ アプリをダウンロード
📱 携帯電話など、あらゆるデバイスからエージェントと連携できます。セッション
あなたをフォローします: 端末で開始し、ブラウザで続行し、で取得します
あなたの電話。メッセージ、サブエージェント、端末、ファイルは同期を保ちます。
🤖 複数のエージェントを監督します。クロード コード、コーデックス、Pi、およびカスタムを使用する
エージェント (YAML で定義) を同じセッション内で一緒に使用します。エージェントに次のことを依頼してください
他の人の作業をレビューしたり、それぞれが得意とするエージェント間でタスクを分割したりできます。
さまざまなこと。
🔌 任意のモデルを使用します。ファーストパーティAP

I キー、Claude/ChatGPT サブスクリプション、
または互換性のあるゲートウェイ。どれも一級品。
🤝 協力してください。セッションを共有すると、チームメイトがエージェントとチャットできるようになります
動作をライブで観察したり、マシン上で共同駆動したり、フォークしたりできます。
会話は独自に続行されます。
☁️ クラウドサンドボックスでエージェントを実行します。ラップトップは必要ありません: セッションを実行します
使い捨てモーダル、デイトナ、または
Islo サンドボックス。CLI から起動するか、によってプロビジョニングされます。
セッションごとのサーバー (管理対象ホスト)。
🛡️ エージェントを管理します。作成
承認のために一時停止するポリシー
危険なアクションの前に、支出を制限したり、エージェントが利用できるツールを制限したりできます。
これらはサーバー全体、1 つのエージェント、または 1 つのチャットに適用されます。
1 つのコマンドで Omnigent とそれに必要なものがすべてインストールされます。
カール -fsSL https://raw.githubusercontent.com/omnigent-ai/omnigent/main/scripts/install_oss.sh |しー
手動でインストールしたいですか?
Omnigent には Python 3.12 以降が必要です。オムニジェント パッケージをインストールします。
uv ツール installomnigent # または: pip install "omnigent"
または Homebrew を使用して:
brew installomnigent-ai/tap/omnigent
または、リポジトリから直接インストールします。
uv ツールのインストール -q --python 3.12 git+https://github.com/omnigent-ai/omnigent.git
ツールチェーンと前提条件 (インストーラーが不足しているツールを報告する場合)
紫外線 (必須)。 https://docs.astral.sh/uv/getting-started/installation/
インストーラーはこれを設定するよう提案します。
npm を備えた Node.js 22 LTS 以降 (Claude、Codex、Pi 用)
コーディングハーネス。 omn​​igent run は、選択したハーネス CLI をインストールします。
https://docs.npmjs.com/downloading-and-installing-node-js-and-npm
tmux 、ネイティブのオムニジェント クロード/オムニジェント コーデックスで必要
ラッパー ( brew install tmux / apt install tmux ; インストーラーが提供するもの
インストールしてください)。
バブルラップ ( bwrap )、Linux のみ。原住民の全知のクロード /
オムニジェント コーデックスと PI ハーネスが各エージェント端末をラップします。

ブラップ
OS サンドボックス。 Linux では分離が必須であるため、bwrap が欠落しています
バイナリにより、これらの端末は起動に失敗します ( apt install bubblewrap ;
インストーラーがインストールを提案します)。 macOS は内蔵シートベルトを使用します
サンドボックスなので追加のものは何も必要ありません。
データブリック (オプション)。 Databricks ワークスペースをモデルとして使用するには
プロバイダーの場合は、Databricks エクストラを使用して Omnigent をインストールします。
UVツール「omnigent[databricks]」をインストールします。ワークスペースにもサインイン
Databricks CLI を使用します。
新しいリリースが PyPI にある場合、Omnigent は 1 行の通知を表示します (1 回につき 1 回)
リリース）ここを指しています。更新するには:
オムニ アップグレード # インストール方法を検出し、ローカル ファイルを排出して停止します
# サーバーを選択し、一致するアップグレード コマンドを実行します
omni upgrade --check # 新しいリリースが利用可能かどうかを報告するだけです
オムニ アップグレードは、実行中のエージェント セッションが終了するのを待ってから、
ローカルサーバー (即座に停止するには --force を渡します);次のオムニコマンド
サーバーを新しいバージョンに戻します。ソース チェックアウトが更新されます
代わりに git pull します。 OMNIGENT_NO_UPDATE_CHECK=1 を使用して通知を沈黙させます。
このチェックは、UV_INDEX_URL / を尊重して、設定されたパッケージ インデックスをクエリします。
PIP_INDEX_URL と uv.toml / pip.conf (デフォルトの PyPI) なので非公開です
ミラーはすぐに使用できます。必要に応じて、OMNIGENT_INDEX_URL でオーバーライドします。
omn​​igent はモデルを選択し、端末でセッションを開始します。それ
http://localhost:6767 でローカル Web UI も起動し、同じ内容が表示されます
ブラウザまたはネットワーク上の電話でセッションを実行します (ステップ 4)。の
デスクトップ アプリは同じ UI をラップします
ネイティブ ウィンドウで OS 通知とドック バッジを追加します —
macOS 用にダウンロードします。
インストールでは、同じ CLI に対して 2 つの名前 (omnigent と ) が PATH に設定されます。
短いオムニ 。交換可能です。
最初の実行時に、Omnigent はモデルの認証情報を取得します。

あなたの中で準備ができています
環境 ( ANTHROPIC_API_KEY / OPENAI_API_KEY 、または claude /
ログインしている codex CLI) がデフォルトとして提供されます。
全能の
または、特定のエージェント ランタイムまたは独自のエージェントを起動します。
omnigent claude # クロード コード、チームが参加できるセッション
全能コーデックス # コーデックス
omnigent run path/to/agent.yaml # 独自のエージェント (「独自のエージェントの作成」を参照)
🐙 ポリーと 🟠🔵 デビー
リポジトリには 2 つのサンプル エージェントが付属しており、最初のセッションは適切に行われます。
オムニジェント実行の例/polly/
オムニジェントの実行例/debby/
# 別のハーネスでオーケストレーターを実行します (サブエージェントは独自のハーネスを保持します)。
オムニジェント実行例/polly/ --harness pi
オムニジェントの実行例/debby/ --harness openai-agents
オムニジェント実行例/polly/ --harness カーソル # カーソル CLI (cursor-agent + CURSOR_API_KEY が必要)
🐙 Polly はマルチエージェントのコーディング オーケストレーターであり、自分自身ではコードを作成しません。
彼女は技術リーダーです。彼女は計画を立て、コーディングのサブエージェントに作業を委任します。
(Claude Code、Codex、または Pi) を並列 git worktree で実行し、各 diff をルーティングします。
執筆者とは別のベンダーのレビューアに送信します。あなたは合併します。
🟠🔵 デビーは、クロードと GPT の 2 つの頭を持つブレインストーミング パートナーです。
あなたが尋ねるすべての質問は両方の頭に浮かび、彼女は2つの答えを提示します
並んで。 /debate と入力すると、いくつかの項目についてヘッドがお互いを批判します。
収束する前にラウンドします。 (彼女はクロードと OpenAI の両方の資格情報を必要とします。
ステップ 3 を参照してください。)
ブラウザの方がいいですか?サーバーを起動し、マシンをホストとして登録します。
オムニジェントサーバーの起動 # ローカルサーバーと Web UI をバックグラウンドで起動します
omnigent host # (別の端末) このマシンをホストとして登録します
Web UI で、[新しいチャット] をクリックし、マシンを選択して次に進みます。ステータスを確認する
オムニジェントサーバーのステータス。全集中停止ですべてを停止します。
オムニジェントセットアップ
資格情報を追加する

ial、デフォルトを設定、またはエージェントごとにグループ化して削除します。オムニジェント
は 4 種類の認証情報を使用します。
デフォルトはエージェントごとに設定されるため、Claude のデフォルトと Codex のデフォルトが共存します。あなた
/model コマンドを使用して、セッションの途中でモデルを切り替えることもできます。
ゲートウェイ資格情報を追加すると、オムニジェント セットアップはベース URL を要求します
そして鍵。ベース URL は、どのエージェントを指すかによって異なります。
Claude コードの場合、OpenRouter の Anthropic 互換エンドポイントをポイントします。
( …/api/v1 ではなく …/api )。 Codex と OpenAI エージェント ハーネスの場合は、次を使用します。
OpenAI 互換の …/api/v1 。
4. サーバーをデプロイします (そして携帯電話から使用します📱)
安定した URL を使用してサーバー上で Omnigent を実行する
(deploy/README.md が完全なガイドです) とセッション
携帯電話を含め、どこからでもアクセスできるようになります。 Web UI は次のために構築されています。
モバイルなので、同じチャット、サブエージェント、端末、ファイルを同期して利用できます
ラップトップと一緒に。
1 つの Docker 構成が、お持ちの任意のホスト (VPS、ホーム) 上でサーバーを実行します。
サーバー);ワンクリックでレンダリングをデプロイします。 Fly.io、鉄道、ハグフェイススペース、
と Modal もカバーされています。サーバーは、次のようにクラウド サンドボックスをプロビジョニングすることもできます。
セッション (管理対象ホスト) なので、ラップトップをオンラインにしておく必要はありません。のフルメニュー
ターゲット、データベース オプション、およびサンドボックスのセットアップが存在します。
デプロイ/README.md 。
サーバーが起動したら、サインインしてラップトップをホストとして登録します。
オムニジェント ログイン https://your-host # 1 回サインインします。実行/接続/ホストでトークンを再利用する
オムニジェント ホスト https://your-host # 新しいセッションがこのマシンで実行できるようになりました
ヒント
独自のネットワークでは、デプロイする必要はありません。マシンの LAN を開きます
携帯電話のアドレス (例: http://192.168.x.x:6767 )。
Omnigent は、1 つの環境で制御されるマルチユーザー アカウントをサポートします
変数:
OMNIGENT_AUTH_ENABLED=1 オムニジェント サーバーの起動
ステップ 4 の Docker デプロイ
あなたのためにそれをオンにします

ou (OMNIGENT_AUTH_ENABLED のデフォルトは 1 です)。
Web UI (ローカルでは http://localhost:6767、またはホストの URL) を開き、
管理者としてサインインします。最初の実行ではパスワードが出力され、ローカルに保存されます。それから
「管理者」→「メンバー」→「招待」を開き、使い捨ての招待リンクを作成します。いいえ
電子メールサーバーが必要です。それを送ってください。チームメイトがそれを開いてパスワードを設定し、
サインアップは招待制です。
チームメイトはサーバーにアクセスできる必要があります。ローカルサーバーのみ
ネットワーク上で到達可能。外部にいる人のために、常時接続のホストを展開します
(ステップ 4 を参照)。
ライブセッションを共有します。 Web UI で [共有] をクリックし、リンクを送信します。
チームメイトはエージェントの作業を監視し、リアルタイムでチャットします。
コドライブ。チームメイトがあなたのランニングセッションに参加します。彼らの
メッセージはマシン上で実行されます。ペアリングや手渡しに最適です。
調査中にドメインの専門家にキーボードを渡します。
オムニジェント アタッチ < session_id >
フォーク。自分のマシンに会話のクローンを作成して続行します
フォークポイントから独立して。
omnigent run --fork < session_id >
チームがすでに持っているログイン情報 (Google、
GitHub、Okta、Microsoft )? OMNIGENT_OIDC_ISSUER とクライアント ID を設定します
デプロイされたサーバー上でシークレットとシークレットを実行し、再起動します。完全なウォークスルー、
ドメイン許可リスト、およびプロキシ専用ヘッダー認証モード

[切り捨てられた]

## Original Extract

A meta-harness for all your AI agents. Omnigent provides a common layer over Claude Code, Codex, Pi, and the agents you write yourself: swap or combine harnesses without rewriting, keep them in check with policies and sandboxing, and collaborate in real time on the same live session, from any device
[truncated]

GitHub - omnigent-ai/omnigent: A meta-harness for all your AI agents. Omnigent provides a common layer over Claude Code, Codex, Pi, and the agents you write yourself: swap or combine harnesses without rewriting, keep them in check with policies and sandboxing, and collaborate in real time on the same live session, from any device. · GitHub
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
omnigent-ai
/
omnigent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
227 Commits 227 Commits .claude/ skills .claude/ skills .github .github ap-web ap-web deploy deploy designs designs dev dev docs docs examples examples omnigent omnigent scripts scripts sdks sdks tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md openapi.json openapi.json pyproject.toml pyproject.toml pyrefly.toml pyrefly.toml railway.toml railway.toml render.yaml render.yaml setup.py setup.py uv.lock uv.lock uv.toml uv.toml View all files Repository files navigation
A meta-harness for all your AI agents
Omnigent provides a common layer over Claude Code, Codex, Cursor, Pi, and the agents you write yourself: swap or combine harnesses without rewriting, keep them in check with policies and sandboxing, and collaborate in real time on the same live session, from any device.
omnigent.ai · ⬇️ Download the macOS desktop app
📱 Work with agents from any device, including your phone. Sessions
follow you: start in your terminal, continue in the browser, pick it up on
your phone. Messages, sub-agents, terminals, and files stay in sync.
🤖 Supervise multiple agents. Use Claude Code, Codex, Pi, and custom
agents (defined in YAML) together in the same session. Ask one agent to
review another's work, or split a task across agents that are each good at
different things.
🔌 Use any model. A first-party API key, a Claude/ChatGPT subscription,
or any compatible gateway. All first-class.
🤝 Collaborate. Share a session so teammates can chat with your agent
and watch it work live, co-drive it on your machine, or fork the
conversation to continue on their own.
☁️ Run agents in cloud sandboxes. No laptop required: run sessions in
disposable Modal , Daytona , or
Islo sandboxes, launched from the CLI or provisioned by
the server per session ( managed hosts ).
🛡️ Govern your agents. Create
policies to pause for your approval
before risky actions, cap spend, or limit which tools an agent reaches.
They apply to the whole server, one agent, or a single chat.
One command installs Omnigent and everything it needs:
curl -fsSL https://raw.githubusercontent.com/omnigent-ai/omnigent/main/scripts/install_oss.sh | sh
Prefer to install manually?
Omnigent needs Python 3.12+ . Install the omnigent package:
uv tool install omnigent # or: pip install "omnigent"
Or with Homebrew :
brew install omnigent-ai/tap/omnigent
Or install straight from the repo:
uv tool install -q --python 3.12 git+https://github.com/omnigent-ai/omnigent.git
Toolchain and prerequisites (if the installer reports a missing tool)
uv (required). https://docs.astral.sh/uv/getting-started/installation/
The installer offers to set this up for you.
Node.js 22 LTS or newer with npm , for the Claude, Codex, and Pi
coding harnesses. omnigent run installs the harness CLI you pick.
https://docs.npmjs.com/downloading-and-installing-node-js-and-npm
tmux , required by the native omnigent claude / omnigent codex
wrappers ( brew install tmux / apt install tmux ; the installer offers
to install it for you).
bubblewrap ( bwrap ), Linux only . The native omnigent claude /
omnigent codex and pi harnesses wrap each agent terminal in a bwrap
OS-sandbox; on Linux that isolation is mandatory, so a missing bwrap
binary makes those terminals fail to start ( apt install bubblewrap ; the
installer offers to install it for you). macOS uses the built-in seatbelt
sandbox and needs nothing extra.
Databricks (optional). To use a Databricks workspace as your model
provider, install Omnigent with the databricks extra:
uv tool install "omnigent[databricks]" . Signing in to the workspace also
uses the Databricks CLI .
When a newer release is on PyPI, Omnigent shows a one-line notice (once per
release) pointing here. To update:
omni upgrade # detects how you installed, drains & stops the local
# server, then runs the matching upgrade command
omni upgrade --check # just report whether a newer release is available
omni upgrade waits for in-flight agent sessions to finish before stopping the
local server (pass --force to stop them immediately); the next omni command
brings the server back up on the new version. Source checkouts update with
git pull instead. Silence the notice with OMNIGENT_NO_UPDATE_CHECK=1 .
The check queries your configured package index — honoring UV_INDEX_URL /
PIP_INDEX_URL and your uv.toml / pip.conf (default PyPI), so private
mirrors work out of the box; override with OMNIGENT_INDEX_URL if needed.
omnigent picks a model with you and starts a session in your terminal. It
also launches a local web UI at http://localhost:6767 that shows the same
session in the browser, or on a phone on your network (step 4). The
desktop app wraps that same UI
in a native window and adds OS notifications and a dock badge —
download it for macOS .
The install puts two names for the same CLI on your PATH: omnigent and
the shorter omni . They're interchangeable.
On first run, Omnigent picks up model credentials already in your
environment (an ANTHROPIC_API_KEY / OPENAI_API_KEY , or a claude /
codex CLI you're logged into) and offers one as the default.
omnigent
Or launch a specific agent runtime, or your own agent:
omnigent claude # Claude Code, in a session your team can join
omnigent codex # Codex
omnigent run path/to/agent.yaml # your own agent (see "Write your own agent")
🐙 Polly and 🟠🔵 Debby
Two example agents ship with the repo, and they make good first sessions:
omnigent run examples/polly/
omnigent run examples/debby/
# Run an orchestrator on a different harness (sub-agents keep their own):
omnigent run examples/polly/ --harness pi
omnigent run examples/debby/ --harness openai-agents
omnigent run examples/polly/ --harness cursor # Cursor CLI (needs cursor-agent + CURSOR_API_KEY)
🐙 Polly is a multi-agent coding orchestrator who writes no code herself.
She's the tech lead: she plans, delegates the work to coding sub-agents
(Claude Code, Codex, or Pi) in parallel git worktrees, then routes each diff
to a reviewer from a different vendor than the one that wrote it. You merge.
🟠🔵 Debby is a brainstorming partner with two heads, one Claude and one GPT.
Every question you ask goes to both heads, and she lays the two answers out
side by side. Type /debate and the heads critique each other for a few
rounds before converging. (She needs both a Claude and an OpenAI credential;
see step 3.)
Prefer the browser? Start a server and register your machine as a host:
omnigent server start # start the local server and web UI in the background
omnigent host # (separate terminal) register this machine as a host
In the web UI, hit New Chat , pick your machine, and go. Check status with
omnigent server status ; stop everything with omnigent stop .
omnigent setup
Add a credential, set a default, or remove one, grouped by agent. Omnigent
works with four kinds of credentials:
Defaults are per agent, so a Claude default and a Codex default coexist. You
can also switch models in the middle of a session with the /model command.
When you add a Gateway credential, omnigent setup asks for a base URL
and a key. The base URL depends on which agent you point it at:
For Claude Code, point at OpenRouter's Anthropic-compatible endpoint
( …/api , not …/api/v1 ). For Codex and the OpenAI-agents harness, use
the OpenAI-compatible …/api/v1 .
4. Deploy a server (and use it from your phone📱)
Run Omnigent on a server with a stable URL
( deploy/README.md is the full guide) and your sessions
become reachable from anywhere, including your phone. The web UI is built for
mobile, so you get the same chat, sub-agents, terminals, and files, in sync
with your laptop.
One docker compose up runs the server on any host you have (a VPS, a home
server); Render deploys with one click; Fly.io, Railway, Hugging Face Spaces,
and Modal are covered too. The server can also provision a cloud sandbox per
session ( managed hosts ), so no laptop has to stay online. The full menu of
targets, the database options, and the sandbox setup live in
deploy/README.md .
Once the server is up, sign in and register your laptop as a host:
omnigent login https://your-host # sign in once; run / attach / host reuse the token
omnigent host https://your-host # new sessions can now run on this machine
Tip
On your own network you don't need a deploy. Open your machine's LAN
address on your phone (e.g. http://192.168.x.x:6767 ).
Omnigent supports multi-user accounts , controlled by one environment
variable:
OMNIGENT_AUTH_ENABLED=1 omnigent server start
The Docker deploy in step 4
turns it on for you ( OMNIGENT_AUTH_ENABLED defaults to 1 there).
Open the web UI ( http://localhost:6767 locally, or your host's URL) and
sign in as admin ; first run prints the password and saves it locally. Then
open Admin → Members → Invite to create a single-use invite link, no
email server needed. Send it over; your teammate opens it, sets a password,
and they're in. Signup is invite-only.
Teammates need to be able to reach the server. A local server is only
reachable on your network; for anyone off it, deploy an always-on host
(see step 4 ).
Share a live session. Hit Share in the web UI and send the link;
teammates watch your agent work and chat with it in real time.
Co-drive. A teammate co-attaches to your running session; their
messages execute on your machine. Great for pairing or handing the
keyboard to a domain expert mid-investigation.
omnigent attach < session_id >
Fork. Clone a conversation onto your own machine and continue
independently from the fork point.
omnigent run --fork < session_id >
Want your team to sign in with the logins they already have ( Google,
GitHub, Okta, Microsoft )? Set OMNIGENT_OIDC_ISSUER plus a client ID
and secret on your deployed server and restart. The full walkthrough,
domain allowlists, and the proxy-only header auth mode a

[truncated]
