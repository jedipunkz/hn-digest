---
source: "https://github.com/Zaimwa9/themis"
hn_url: "https://news.ycombinator.com/item?id=48903007"
title: "Show HN: Themis – Self-hosted AI code reviews with your own keys and models"
article_title: "GitHub - Zaimwa9/themis: Reviewer Agent you can self-host · GitHub"
author: "Diwadoo"
captured_at: "2026-07-14T07:04:26Z"
capture_tool: "hn-digest"
hn_id: 48903007
score: 2
comments: 0
posted_at: "2026-07-14T06:34:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Themis – Self-hosted AI code reviews with your own keys and models

- HN: [48903007](https://news.ycombinator.com/item?id=48903007)
- Source: [github.com](https://github.com/Zaimwa9/themis)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T06:34:22Z

## Translation

タイトル: Show HN: Themis – 独自のキーとモデルを使用した自己ホスト型 AI コード レビュー
記事タイトル: GitHub - Zaimwa9/themis: セルフホストできるレビューアーエージェント · GitHub
説明: 自己ホストできるレビューアー エージェント。 GitHub でアカウントを作成して、Zaimwa9/themis の開発に貢献してください。
HN テキスト: HN さん、私は仕事や副業プロジェクトで使用しているコード レビュー ツールに満足していませんでした。騒々しい、表面のレビュー、高価（サイドとしてはやりすぎ）。
それで、ジェミニが日没に近づいていたので、私は自分のものを持っていることを模索しました。私は安っぽいネズミなので、この週末、Codex/Claude サブスクリプションで動作する Themis を構築しました (サブスクリプションでは「claude -p」が再び OK であることを理解しました)。他のインディーズ開発者や小規模チームを支援できるように、自己ホスト可能 (またはローカル)、サブスクライブから描画し、ほとんどの場合は構成可能であるため、独自のリポジトリでうまく機能するように、オープンソースにしてほしかったのです。ここには大きな野望も 10 億のスタートアップもありません。ただ楽しいだけです (だから、あーあなどと思わないでください)。数人でも役に立てば、それはそれで素晴らしいことです。フィードバックをいただければ幸いです (特にアプリの配布について!) 追伸: 1 番目の HN 投稿、それが正しい形式であることを願っています

記事本文:
GitHub - Zaimwa9/themis: セルフホストできるレビューアー エージェント · GitHub
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
ザイムワ9
/
テーマ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く

アクションメニュー フォルダーとファイル
106 コミット 106 コミット .github/ workflows .github/ workflows .themis .themis docs docsexamples/themisexamples/themissrc/themissrc/themis テストテスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .python-version .python-version .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml release-please-config.json release-please-config.json uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Themis は、独自の Codex 上で実行される自己ホスト型の GitHub PR レビュー ボットです。
クロード・マックスの定期購読。インライン結果と構造化されたレポートを使用してプル リクエストをレビューします。
概要 (評決、スコア表、重大度順のセクション)、回答
レビュー スレッドや PR 会話で質問し、レビューを受け付けます
.themis/ の下にある独自のリポジトリからの教義。
GitHub App Webhook は、PR およびコメント イベントを Themis に配信します。各イベント
メモリ内キュー上のジョブとなり、一度に 1 つずつ処理されます。労働者
PR ヘッドをシャロークローンし、設定されたエンジンを実行します ( codex exec 、または
claude -p — ネイティブまたは GLM の API モードで)、リポジトリのレビュー原則に対して、調査結果と
概要をアプリとして GitHub に戻します。 1 つのイメージが分離されたコントローラーとして実行されます
そしてエージェント。データベース、Redis、またはメッセージ ブローカーはまだありません。
Compose プラグインを備えた Docker (Docker Desktop、または Docker Engine + docker compose )
Codex アクセスのある OpenAI アカウント、または Claude Max サブスクリプション (エンジンを選択してください): 一致する CLI がインストールされており ( npm install -g @openai/codex または npm install -g @anthropic-ai/claude-code 、Node 22+)、Codex ログインまたは claude setup-token がマシン上で動作しています。クロード

Themis のデフォルトは Opus であるため、Pro はサポートされていません。
glm エンジンにはローカル CLI ログインは必要ありません。必要なのは、 .env 内の GLM_API_KEY (Z.ai GLM コーディング プラン) だけです。
GitHub アプリを作成できる GitHub アカウント (個人アカウントまたは組織)
クローンやビルドは必要ありません。Themis はビルド済みイメージとして出荷されます。
ghcr.io/zaimwa9/themis 。
クイックスタート (マニフェスト ブートストラップ - すぐに試せますが、長期使用には向いていません)
ブートストラップは、GitHub のアプリ マニフェスト フローを使用してアプリを作成し、そのアプリを生成します。
秘密キーと Webhook シークレットを取得し、要求されたリポジトリにインストールします。
すぐに実行できるデプロイメントを作成します。コピーする GitHub 設定はありません。 GitHub
引き続き、アカウント所有者にアプリの作成とリポジトリへのアクセスを承認するよう求めます。
なぜ長期ではないのでしょうか？マニフェスト フローはランダムなアプリ名を生成するため、
新しいデプロイメントごとに再実行されます。永続的にインストールするには、GitHub を作成します
安定した名前とスラッグを使用して手動でアプリを作成し、Themis にそれを指定します — を参照してください。
手動パスの docs/bootstrap.md。
以下のプロンプトを Claude Code、Codex、または実行可能なエージェントにコピーします。
シェルコマンド。ブートストラップ全体を自律的に処理します。
自動化された GitHub App Manifest ブートストラップを使用して、このリポジトリに Themis をセットアップします。
https://github.com/Zaimwa9/themis/blob/main/docs/bootstrap.md
Claude エンジンとバンドルされた ngrok トンネルを使用します。
GitHub アプリを手動で作成または構成しないでください。 `python -m themis init` を実行します。
「--エンジン クロード --トンネル」。
すべてを自律的に処理します。 GitHub が私の承認を必要とする場合、または次の場合にのみ一時停止します。
ngrok 認証トークンを提供するか、クロード認証を完了する必要があります。
セットアップ後、デプロイを開始し、動作することを確認して、ボットの @mention を教えてください。
レビューをトリガーする方法についても説明します。
1.コーデックスにログインします
代わりにクロード エンジンを使用しますか? clude setup-token を実行し、パスします
--engine claude をブートストラップに追加し、結果を置きます

トークンが入っています
生成された .env 内の CLAUDE_CODE_OAUTH_TOKEN 。 glmを使用していますか？ CLIなし
ログインが必要です: --engine glm をブートストラップに渡し、Z.ai GLM を配置します。
生成された .env の GLM_API_KEY のコーディング プラン キー。詳細は
エンジン。
CLI をまだインストールしていない場合はインストールします (ノード 22 以降)。
npm install -g @openai/codex
コーデックスログイン
これにより、資格情報が ~/.codex/auth.json に書き込まれます。ブートストラップはそれをコピーします
モード 0600 で生成されたデプロイメント。
既存のパブリック HTTPS URL を持つインスタンスの場合:
mkdir themis-deploy
docker run --rm -it \
--user " $( id -u ) : $( id -g ) " \
-p 127.0.0.1:8976:8976 \
-v " $PWD /themis-deploy:/output " \
-v " $HOME /.codex:/host-codex:ro " \
ghcr.io/zaimwa9/themis:最新 \
Python -m theis init \
--repo 所有者/リポジトリ \
--public-url https://themis.example.com \
--出力 /出力 \
--codex-auth /host-codex/auth.json \
--バインドホスト 0.0.0.0 \
--ブラウザなし
パブリック URL のないローカル マシンの場合は、ngrok トークンをエクスポートして置き換えます。
--public-url ... --tunnel を使用:
import NGROK_AUTHTOKEN= <あなたのトークン>
# docker run に `-e NGROK_AUTHTOKEN` を追加し、`--tunnel` を init に渡します。
ターゲットが組織によって所有されている場合は、 --organization OWNER も渡します。
生成されたアプリは次のとおりであるため、アプリはその組織の下で作成する必要があります。
プライベート。このコマンドはローカルホスト URL を出力します。それを開いて、事前に入力された内容を承認します
アプリを選択し、インストール画面で OWNER/REPO を選択します。秘密鍵は決して
ブートストラップ プロセスと生成された .env が残ります。成功画面と
ターミナルはボットが生成した @mention を表示します。にも保存されています
後で参照するための themis-info.json。
Docker を使用せずに、ソース チェックアウトから同じコマンドを実行できます。
uv run python -m themis init \
--repo 所有者/リポジトリ \
--public-url https://themis.example.com \
--output ./themis-deploy
詳しいオプション、トンネル

コマンド、リカバリ、および手動フォールバックは次のとおりです。
docs/bootstrap.md 。
代わりに GitHub アプリを手動で作成しますか?これらのリポジトリを設定します
権限:
pull_request 、 issue_comment 、および
pull_request_review_comment 。アクション権限は必要ありません。既存の
事前に、ステータスのチェックとコミットの権限を使用してアプリを更新する必要もあります
Themis はレビューに CI コンテキストを含めることができます。
cd themis-deploy
ドッカー構成 -d
ブートストラップで --tunnel が使用された場合は、 docker compose --profile tongue up -d を使用します。
Themis はトンネル URL を検出し、アプリ Webhook を自動的に更新します。
カールローカルホスト:8000/healthz
# {"ステータス":"ok"}
docker compose ログで、
アプリのスラッグを解決しました。アプリがインストールされているリポジトリに対してテスト PR を開きます。
トリガーで 👀 の反応を期待し、その後 PR で 🚀 の反応を期待します
レビューが始まり、その後レビュー自体が始まります。
スターター キットをターゲット リポジトリにコピーします。これには一時的な浅瀬が必要です
Themis のチェックアウト (デプロイ自体は行いません):
スターター= " $( mktemp -d ) "
git clone -- Depth 1 https://github.com/Zaimwa9/themis.git " $starter "
cp -r " $starter /examples/themis " .themis
.themis/review.md : レビューの原則、哲学、厳しさ
キャリブレーション、コードベースのマップ、ハウスルール。これを編集してください。読まれています
すべてのレビューで PR ブランチから直接送信されます。
.themis/config.yaml : 動作ノブ、すべてのキーはオプションです。
教義がどのように利用されるか、そして機能する教義をどのように書くか:
docs/doctrine.md 。このリポジトリ自体を独自のリポジトリでレビューします
.themis/review.md 。
PR でボットに話しかけます: @<app-slug> オンデマンドで再レビューします。
@<app-slug> review <focus> は、レビューを特定の領域 (
フォーカス テキストは、リポジトリ所有者、組織メンバー、および
協力者 — 他の人は単純なレビューを受けます)、@<app-slu

g> <質問>
質問をし、ボットがすでに投稿されているスレッド内で返信します。
自動的に応答されるため、言及する必要はありません。
Themis は、Codex、Claude Max、または GLM コーディング プランのサブスクリプションを使用して、エージェント CLI を通じてレビューを実行します。
.env の THEMIS_ENGINE でインスタンスのデフォルトを選択します。リポジトリはそれをオーバーライドできます
エンジンを含む .themis/config.yaml 内: それらのいずれかに設定します。あのエンジンなら
インスタンスに認証情報がない場合、Themis は代わりにその旨のコメントを投稿します。
静かに失敗する。 clude パスと glm パスにはボリュームは必要ありません: と入力します。
.env 、完了しました。
症状
修正
環境変数に名前を付けると起動時にクラッシュする
その変数を設定します。 Themis は、必要な構成が欠落しているか無効であるとすぐに失敗します。
GET /app 呼び出しの起動時にクラッシュする
THEMIS_GH_APP_CLIENT_ID が間違っているか、THEMIS_GH_APP_PRIVATE_KEY の形式が正しくありません。
コーデックスサンドボックスエラー
THEMIS_CODEX_SANDBOX=危険フルアクセスを設定します。 Landlock を使用しないランタイムでは、コンテナーがサンドボックス境界になります。
PR コメントには、使用制限に達したことが記載されています
ジョブを実行したエンジンのサブスクリプションは、使用期間に達しました。ボットがリセットされたら、もう一度ボットについてメンションします。
機能していた認証が数か月後に失敗し始める
codex ログインをローカルで実行し、「自動セットアップ: Codex 認証の更新」のコマンドを使用して永続的なエージェント認証情報を更新します。
レビューのコメントには、エンジンの認証情報が欠落していると記載されています
CLAUDE_CODE_OAUTH_TOKEN (claude)、GLM_API_KEY (glm) を設定するか、コーデックス認証ボリューム (codex) をシードするか、THEMIS_ENGINE / リポジトリのエンジン: キーを変更します。
Webhook 配信では、アプリの設定に 401 が表示されます
THEMIS_GH_WEBHOOK_SECRET はアプリの Webhook シークレットと一致しません。
ログはどこにありますか
docker compose ログ -f themis
再起動直前にキューに入れられたジョブが実行されなかった
メモリ内のキューは再起動後に存続しません。ボットをもう一度メンションすると、再トリガーされます。
PRを開いたが何も起こらなかった
順番にチェックしてください: PR は医師です

船尾 (準備完了とマークされるまでスキップ); PR 作成者はボット アカウントです (無視されます)。 auto_review: .themis/config.yaml の false ; GitHub アプリはそのリポジトリにインストールされていません。 Webhook 配信が失敗します (アプリ設定 > 詳細 > 最近の配信)。
ドキュメント
docs/server-deploy.md : 任意の Docker ホストまたは PaaS へのデプロイ、アップグレード。
docs/local-tunnel.md : ngrok トンネルの詳細なプロファイル。
docs/headless.md : 独自の Webhook ハンドラー、/api/review および /api/discuss コントラクトを使用します。
docs/doctrine.md : レビューの原則、その仕組み、および良いレビューの書き方。
docs/configuration.md : 完全な env および .themis/config.yaml リファレンス。
docs/security.md : 信頼モデルとボット側のガードレール。
docs/contributing-engines.md : 新しいエンジン/モデルプロバイダーを追加します。
Themis 自体の作業には uv と Python 3.12 が必要ですが、Docker は必要ありません。
uv sync --locked # .venv に deps をインストールします
uv run pytest # テストスイートを実行します
UVランラフチェック。 #糸くず
uv run python -m themis # サーバーをローカルで実行します (環境から THEMIS_* を読み取ります)
CI は、すべてのプッシュ リクエストとプル リクエストで同じ pytest コマンドと ruff コマンドを実行します。
自己ホストできるレビューアー エージェント
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
5
v0.3.0
緯度

[切り捨てられた]

## Original Extract

Reviewer Agent you can self-host. Contribute to Zaimwa9/themis development by creating an account on GitHub.

Hey HN, I wasn't happy with the code review tools we use at work and on my side projects. Noisy, reviews in surface, expensive (overkill for sides).
So since Gemini was getting sunsetted, I explored having my own. As I am a cheap rat, this weekend I built Themis, that works with my Codex/Claude subscriptions (I understood "claude -p" is ok again in subscriptions). I wanted it open-source so it can help other indie devs or small teams, self-hostable (or local), drawing from subs and mostly, configurable so it works well with your own repos. No big ambitions nor 1 billion startup here, just fun (so please don't roast ahah). If it is useful to even a few people, that's already great. Happy to have feedback (especially on distributing the app!) Ps: 1st HN contrib, hope it's the right format

GitHub - Zaimwa9/themis: Reviewer Agent you can self-host · GitHub
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
Zaimwa9
/
themis
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
106 Commits 106 Commits .github/ workflows .github/ workflows .themis .themis docs docs examples/ themis examples/ themis src/ themis src/ themis tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .python-version .python-version .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml release-please-config.json release-please-config.json uv.lock uv.lock View all files Repository files navigation
Themis is a self-hosted GitHub PR review bot that runs on your own Codex or
Claude Max subscription. It reviews pull requests with inline findings and a structured
summary (verdict, scoring table, severity-ordered sections), answers
questions in review threads and PR conversation, and takes its review
doctrine from your own repository, under .themis/ .
A GitHub App webhook delivers PR and comment events to Themis. Each event
becomes a job on an in-memory queue, processed one at a time. The worker
shallow-clones the PR head, runs the configured engine ( codex exec , or
claude -p — natively or in API mode for GLM) against your repo's review doctrine, and posts findings and a
summary back to GitHub as the App. One image runs as an isolated controller
and agent; there is still no database, Redis, or message broker.
Docker with the Compose plugin (Docker Desktop, or Docker Engine + docker compose )
An OpenAI account with Codex access, or a Claude Max subscription (pick your engine): the matching CLI installed ( npm install -g @openai/codex or npm install -g @anthropic-ai/claude-code , Node 22+) and codex login or claude setup-token working on your machine. Claude Pro is not supported because Themis defaults to Opus.
The glm engine needs no local CLI login: just a GLM_API_KEY (Z.ai GLM Coding Plan) in .env .
A GitHub account that can create a GitHub App (personal account or an org)
No clone or build needed: Themis ships as a prebuilt image,
ghcr.io/zaimwa9/themis .
Quickstart (manifest bootstrap — fast to try, not for long-term use)
The bootstrap uses GitHub's App Manifest flow to create the App, generate its
private key and webhook secret, install it on the requested repository, and
write a ready-to-run deployment. There are no GitHub settings to copy. GitHub
still asks the account owner to approve App creation and repository access.
Why not long-term? The manifest flow generates a random App name and must
be re-run on every fresh deployment. For a permanent install, create a GitHub
App manually with a stable name and slug, then point Themis at it — see
docs/bootstrap.md for the manual path.
Copy the prompt below into Claude Code, Codex, or any agent that can run
shell commands. It handles the entire bootstrap autonomously:
Set up Themis for this repository using the automated GitHub App Manifest bootstrap:
https://github.com/Zaimwa9/themis/blob/main/docs/bootstrap.md
Use the Claude engine and the bundled ngrok tunnel.
Do not manually create or configure a GitHub App. Run `python -m themis init` with
`--engine claude --tunnel`.
Handle everything autonomously. Only pause when GitHub requires my approval, or when
I need to provide an ngrok auth token or complete Claude authentication.
After setup, start the deployment, verify it works, and tell me the bot's @mention
and how to trigger a review.
1. Log in to Codex
Using the Claude engine instead? Run claude setup-token , pass
--engine claude to the bootstrap, and put the resulting token in
CLAUDE_CODE_OAUTH_TOKEN in the generated .env . Using glm? No CLI
login needed: pass --engine glm to the bootstrap and put your Z.ai GLM
Coding Plan key in GLM_API_KEY in the generated .env . Details in
Engines .
Install the CLI if you haven't already (Node 22+):
npm install -g @openai/codex
codex login
This writes credentials to ~/.codex/auth.json . The bootstrap copies it into
the generated deployment with mode 0600 .
For an instance with an existing public HTTPS URL:
mkdir themis-deploy
docker run --rm -it \
--user " $( id -u ) : $( id -g ) " \
-p 127.0.0.1:8976:8976 \
-v " $PWD /themis-deploy:/output " \
-v " $HOME /.codex:/host-codex:ro " \
ghcr.io/zaimwa9/themis:latest \
python -m themis init \
--repo OWNER/REPO \
--public-url https://themis.example.com \
--output /output \
--codex-auth /host-codex/auth.json \
--bind-host 0.0.0.0 \
--no-browser
For a local machine without a public URL, export an ngrok token and replace
--public-url ... with --tunnel :
export NGROK_AUTHTOKEN= < your-token >
# Add `-e NGROK_AUTHTOKEN` to docker run, then pass `--tunnel` to themis init.
If the target is owned by an organization, also pass --organization OWNER .
The App must be created under that organization because the generated App is
private. The command prints a localhost URL. Open it, approve the pre-filled
App, and select OWNER/REPO on the installation screen. The private key never
leaves the bootstrap process and the generated .env . The success screen and
terminal show the bot's generated @mention ; it is also saved in
themis-info.json for later reference.
The same command can run from a source checkout without Docker:
uv run python -m themis init \
--repo OWNER/REPO \
--public-url https://themis.example.com \
--output ./themis-deploy
Detailed options, the tunnel command, recovery, and the manual fallback are in
docs/bootstrap.md .
Creating the GitHub App manually instead? Configure these repository
permissions:
Subscribe to pull_request , issue_comment , and
pull_request_review_comment . Actions permission is not required. Existing
Apps must also be updated with the Checks and Commit statuses permissions before
Themis can include CI context in reviews.
cd themis-deploy
docker compose up -d
Use docker compose --profile tunnel up -d when the bootstrap used --tunnel .
Themis discovers the tunnel URL and updates the App webhook automatically.
curl localhost:8000/healthz
# {"status":"ok"}
Check docker compose logs themis for the startup line reporting the
resolved App slug. Open a test PR against a repo the App is installed on:
expect a 👀 reaction on the trigger, then a 🚀 reaction on the PR once the
review starts, then the review itself.
Copy the starter kit into the target repo. This needs a temporary shallow
checkout of Themis (the deployment itself does not):
starter= " $( mktemp -d ) "
git clone --depth 1 https://github.com/Zaimwa9/themis.git " $starter "
cp -r " $starter /examples/themis " .themis
.themis/review.md : the review doctrine, philosophy, severity
calibration, a map of your codebase, house rules. Edit this; it is read
straight from the PR branch on every review.
.themis/config.yaml : behavior knobs, every key optional.
How the doctrine is consumed and how to write one that works:
docs/doctrine.md . This repo reviews itself with its own
.themis/review.md .
Talk to the bot in a PR: @<app-slug> review re-reviews on demand,
@<app-slug> review <focus> steers the review toward a given area (the
focus text is honored only from repo owners, org members, and
collaborators — anyone else gets a plain review), @<app-slug> <question>
asks a question, and replies inside a thread the bot already posted in are
answered automatically, no mention needed.
Themis runs reviews through an agent CLI, using your Codex, Claude Max, or GLM Coding Plan subscription:
Pick the instance default with THEMIS_ENGINE in .env . A repo can override it
in .themis/config.yaml with engine: set to any of them; if that engine
has no credentials on the instance, Themis posts a comment saying so instead of
failing silently. The claude and glm paths need no volume: key in
.env , done.
Symptom
Fix
Crashes at startup naming an env var
Set that variable; Themis fails fast on missing or invalid required config.
Crashes at startup on a GET /app call
Wrong THEMIS_GH_APP_CLIENT_ID or malformed THEMIS_GH_APP_PRIVATE_KEY .
Codex sandbox errors
Set THEMIS_CODEX_SANDBOX=danger-full-access ; the container is the sandbox boundary on runtimes without Landlock.
PR comment says the usage limit was reached
The subscription of whichever engine ran the job has hit its usage window. Mention the bot again once it resets.
Auth that worked starts failing months later
Run codex login locally, then refresh the persistent agent credential using the command in Automated setup: Refreshing Codex authentication .
Review comment says engine credentials missing
Set CLAUDE_CODE_OAUTH_TOKEN (claude), GLM_API_KEY (glm), or seed the codex auth volume (codex), or change THEMIS_ENGINE / the repo's engine: key.
Webhook deliveries show 401 in the App's settings
THEMIS_GH_WEBHOOK_SECRET doesn't match the App's webhook secret.
Where are the logs
docker compose logs -f themis
A job queued right before a restart never ran
The in-memory queue doesn't survive restarts; mention the bot again to re-trigger.
Opened a PR, nothing happened
Check in order: PR is a draft (skipped until marked ready); PR author is a bot account (ignored); auto_review: false in .themis/config.yaml ; the GitHub App isn't installed on that repo; webhook deliveries are failing (App settings > Advanced > Recent Deliveries).
Documentation
docs/server-deploy.md : deploying to any Docker host or PaaS, upgrades.
docs/local-tunnel.md : the ngrok tunnel profile in depth.
docs/headless.md : bring your own webhook handler, the /api/review and /api/discuss contracts.
docs/doctrine.md : the review doctrine, how it works and how to write a good one.
docs/configuration.md : the full env and .themis/config.yaml reference.
docs/security.md : the trust model and bot-side guardrails.
docs/contributing-engines.md : adding a new engine / model provider.
Working on Themis itself needs uv and Python 3.12, no Docker:
uv sync --locked # install deps into .venv
uv run pytest # run the test suite
uv run ruff check . # lint
uv run python -m themis # run the server locally (reads THEMIS_* from the environment)
CI runs the same pytest and ruff commands on every push and pull request.
Reviewer Agent you can self-host
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
5
v0.3.0
Lat

[truncated]
