---
source: "https://github.com/baristaGeek/minions-army-harness"
hn_url: "https://news.ycombinator.com/item?id=49090454"
title: "Show HN: Minions Army Harness – Small Version of Claude Tag"
article_title: "GitHub - baristaGeek/minions-army-harness · GitHub"
author: "baristaGeek"
captured_at: "2026-07-28T22:04:29Z"
capture_tool: "hn-digest"
hn_id: 49090454
score: 2
comments: 0
posted_at: "2026-07-28T21:57:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Minions Army Harness – Small Version of Claude Tag

- HN: [49090454](https://news.ycombinator.com/item?id=49090454)
- Source: [github.com](https://github.com/baristaGeek/minions-army-harness)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T21:57:06Z

## Translation

タイトル: 表示 HN: ミニオン アーミー ハーネス – クロード タグの小さいバージョン
記事タイトル: GitHub - baristaGeek/minions-army-harness · GitHub
説明: GitHub でアカウントを作成して、baristaGeek/minions-army-harness の開発に貢献します。

記事本文:
GitHub - バリスタギーク/ミニオンアーミーハーネス · GitHub
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
バリスタオタク
/
ミニオンアーミーハーネス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .github/ workflows .github/ workflows .vscode .vscode alembic alembic config_examples config_examples docs docs evals evals 実行 実行ログ ログ minions_army minions_army サンプルアプリ サンプルアプリのテスト テスト user_data user_data .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONSTITUTION.md CONSTITUTION.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.minion Dockerfile.minion ライセンス ライセンス Makefile Makefile QUICKSTART.md QUICKSTART.md README.md README.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml fly.minion.toml fly.minion.toml fly.toml fly.toml pyproject.toml pyproject.toml 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
責任あるバイブコーディングのための小型の Slack ファースト エージェント ハーネス。
平易な英語のリクエスト (「アプリにダーク テーマを追加する」、「最近のトランザクション テーブルに販売者の列を追加する」) を送信すると、それが仕様に変換され、分離されたサンドボックスにコードが書き込まれ、プル リクエストが開かれ、敵対的にレビューされ、許可されればデプロイされます。 CLI を学習したり、コーディング エージェントに直接触れたりする必要はありません。これは、「コード化ミニオンの艦隊」というアイデアを小型化した、自己ホスト可能なものです。
Anthropic/OpenAI/Kimi API キーと、所有するリポジトリの GitHub トークンが必要です (これをフォークします。
minion は、バンドルされているサンプルアプリ/ ) に対して動作します。
import ANTHROPIC_API_KEY=sk-ant-... // または OPENAI_API_KEY // または KIMI_API_KEY
エクスポート GITHUB_TOKEN=ghp_...
import REPOSITORY_NAME=your-username/minions-army-harness # トークンがプッシュできるリポジトリ
docker compose up --build # API、Postgresを起動し、ミニオンイメージをビルドします
# 別のシェルでリクエストを発行します。

Webhook:
カール -X POST http://localhost:8000/api/v1/webhooks/slack/me
-H " Content-Type: application/json " \
-d ' {"チャンネル":"C123","ユーザー":"U123","テキスト":"フッターを追加します。1"} '
ミニオンはローカルの Docker 兄弟コンテナとして起動し、フォークに対して OpenSpec を実行し、ゲートオンします。
npm run build し、実際のプルリクエストを開きます。 ultをレビューするのでデモ
PRで止まります。実行ごとに実際の LLM トークンがかかり、ボットがコミットしても構わないリポジトリが必要です。
構成は user_data/api/config.yml (API によって読み取られる) および user_data/orchestrator/config.yml にあります。
(ミニオン内を読んでください)。
2 つの Fly アプリ: API ( fly.toml ) とミフェメラルを所有する 1 つ
マシン ( fly.minion.toml )。
フライアプリはミニオンアーミーを作成します
フライアプリはミニオンを作成します-アーミー-ミニオン
# 1. ミニオン イメージをビルドしてプッシュします。 --image-label は以下の launcher.image と一致する必要があります。
flydeploy --config fly.minion.toml --remote-only --build-oatest
# 2. user_data/api/config.yml を編集します — API imaoy に組み込まれています:
# ランチャー:
# バックエンド: fly_machines
# 画像: registry.fly.io/minions-army-minion:latest
# fly_machine_app: minions-army-minion # e をホストします
# fly_app: <デプロイするアプリ> #deploy.mode が flyctl の場合は必須
# fly_api_token: ${FLY_API_TOKEN} # 環境フォールなし
# repository.name + verify.command/cwd -> リポジトリ
# reviewer.enabled: true
# デプロイ.モード: flyctl
#slack.allowed_channel_id: <ボットがリッスンするチャネル>
# 3. シークレットは ${VAR} プレースホルダーに入力されます。
フライ シークレット セット --app minions-army \
DATABASE_URL=postgres://... ANTHROPIC_API_KEY=sk-ant-...
SLACK_BOT_TOKEN=xoxb-... FLY_API_TOKEN= $( フライトークン作成デプロイ )
# 4. デプロイしてから移行します (fly.toml には release_command がありません)。
flydeploy --config fly.toml --remote-only
fly ssh console --app minions-army -C " alembic upgrade head "
次に、Slack アプリを作成します。イベント サブスクリプションをポイントします。

t
https://minions-army.fly.dev/api/v1/webhooks/slack/messages (エンドポイントは Slack の
チャレンジ )、 app_mention をサブスクライブし、ボットをスレッド内に提供します。
署名秘密のチェックはないため、slack.allowed_channel_id がゲートとなります。
すべてのメッセージは独自の一時的な Fly Machine を取得し、受信時に作成され、終了時に破棄されます。
プロセスごとに 1 つの YAML ファイル。読み込み時に環境から埋められる ${VAR} プレースホルダーが含まれます。両方
出荷時のデフォルト。最初から始めるには、cp config_exta/api/config.yml
(そのパスはローダーが読み取るパスです。 MINIONS_CONFIG_PATH でオーバーライドします)。キーノブ:
pip install -r 要件-dev.txt
ラフチェックミニオン_アーミーテスト
mypy ミニオンズアーミー
pytest
仕組み
メッセージは、Slack Webhook (または汎用 Web API エンドポイント) 経由で受信されます。
ミニオンは、ローカルの Docker 兄弟コンテナ (Tier 1) または
ランが終了すると破壊される一時的なフライ マシン (Tier 2)。
エージェントは、軽量の OpenSpec を実行します。
仕様駆動開発フロー — リクエストを明確にして具体的な仕様にし、それを実装します。
プロバイダーは Claude → Codex → Kim にフォールバックするため、1 回の停止やクレジット残高が空になっても実行は停止しません。
PR が開かれ、敵対的にレビューされます。 LLMはおべっかであり、自分自身を判断するのが苦手です。
機能するため、別のレビュー担当者が PR を採点します。ライターとレビューアーの両方が共有しているのは、
破壊的な操作を禁止するエンジニアリング憲法 (DROP TABLE 、
無条件の DELETE はなく、追加および可逆的な移行のみ)。
船とか旗とか。レビュー担当者を有効にすると、承認された PR を自動マージしてデプロイできます。それ以外の場合は
PRは人間に任せる。 (デフォルトではどちらもオフになっています。デモでは PR が開くだけです。)
PR を開く前に CONTRIBUTING.md を参照してください。
Readme MIT ライセンス
0 フォーク レポート リポジトリ Rel

緩和する
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to baristaGeek/minions-army-harness development by creating an account on GitHub.

GitHub - baristaGeek/minions-army-harness · GitHub
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
baristaGeek
/
minions-army-harness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github/ workflows .github/ workflows .vscode .vscode alembic alembic config_examples config_examples docs docs evals evals execution execution logs logs minions_army minions_army sample-app sample-app tests tests user_data user_data .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONSTITUTION.md CONSTITUTION.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.minion Dockerfile.minion LICENSE LICENSE Makefile Makefile QUICKSTART.md QUICKSTART.md README.md README.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml fly.minion.toml fly.minion.toml fly.toml fly.toml pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
A small, Slack-first agent harness for responsible vibe-coding .
Send it a plain-English request ("add a dark theme to the app", "add a merchant column to the recent transactions table") and it turns that into a spec, writes the code in an isolated sandbox, opens a pull request, reviews it adversarially, and — if you let it — deploys. No one has to learn a CLI or touch a coding agent directly. It's a miniature, self-hostable take on the "fleet of coding minions" idea.
You need an Anthropic/OpenAI/Kimi API key and a GitHub token for a repo you own (fork this one — the
minion works against the bundled sample-app/ ).
export ANTHROPIC_API_KEY=sk-ant-... // or OPENAI_API_KEY // or KIMI_API_KEY
export GITHUB_TOKEN=ghp_...
export REPOSITORY_NAME=your-username/minions-army-harness # a repo your token can push to
docker compose up --build # boots the API, Postgres, and builds the minion image
# In another shell, fire a request at the webhook:
curl -X POST http://localhost:8000/api/v1/webhooks/slack/me
-H " Content-Type: application/json " \
-d ' {"channel":"C123","user":"U123","text":"add a footer .1"} '
A minion spins up as a local Docker sibling container, runs OpenSpec against your fork, gates on
npm run build , and opens a real pull request . Review ult, so the demo
stops at the PR. It costs real LLM tokens per run and needs a repo you don't mind a bot committing to.
Config lives in user_data/api/config.yml (read by the API) and user_data/orchestrator/config.yml
(read inside the minion).
Two Fly apps: the API ( fly.toml ) and one that owns the miphemeral
machines ( fly.minion.toml ).
fly apps create minions-army
fly apps create minions-army-minion
# 1. Build and push the minion image. --image-label must match launcher.image below.
fly deploy --config fly.minion.toml --remote-only --build-oatest
# 2. Edit user_data/api/config.yml — baked into the API imaoy:
# launcher:
# backend: fly_machines
# image: registry.fly.io/minions-army-minion:latest
# fly_machine_app: minions-army-minion # hosts the e
# fly_app: <the app you deploy> # required when deploy.mode is flyctl
# fly_api_token: ${FLY_API_TOKEN} # no env fall
# repository.name + verification.command/cwd -> your repo
# reviewer.enabled: true
# deploy.mode: flyctl
# slack.allowed_channel_id: <channel the bot listens to>
# 3. Secrets fill the ${VAR} placeholders.
fly secrets set --app minions-army \
DATABASE_URL=postgres://... ANTHROPIC_API_KEY=sk-ant-...
SLACK_BOT_TOKEN=xoxb-... FLY_API_TOKEN= $( fly tokens create deploy )
# 4. Deploy, then migrate (fly.toml has no release_command).
fly deploy --config fly.toml --remote-only
fly ssh console --app minions-army -C " alembic upgrade head "
Then create a Slack app: point Event Subscriptions at
https://minions-army.fly.dev/api/v1/webhooks/slack/messages (the endpoint echoes Slack's
challenge ), subscribe to app_mention , and give the bot ly in-thread.
There is no signing-secret check, so slack.allowed_channel_id is the gate.
Every message now gets its own ephemeral Fly Machine, created on receipt and destroyed on exit.
One YAML file per process, with ${VAR} placeholders filled from the environment at load time. Both
ship working defaults; to start from scratch, cp config_exta/api/config.yml
(that path is the one the loader reads — override with MINIONS_CONFIG_PATH ). Key knobs:
pip install -r requirements-dev.txt
ruff check minions_army tests
mypy minions_army
pytest
How it works
A message comes in — via the Slack webhook (or the generic Web API endpoint).
A minion is spawned in an isolated sandbox: a local Docker sibling container (Tier 1) or an
ephemeral Fly Machine (Tier 2) that is destroyed when the run ends.
The agent runs OpenSpec — a lightweight
spec-driven-development flow — to disambiguate the request into a concrete spec, then implement it.
Providers fall back Claude → Codex → Kimi , so one outage or empty credit balance doesn't stop the run.
A PR is opened, then reviewed adversarially. LLMs are sycophantic and poor at judging their own
work, so a separate reviewer agent grades the PR. Both writer and reviewer share an
engineering constitution that forbids destructive operations (no DROP TABLE ,
no unconditional DELETE , additive-and-reversible migrations only).
Ship or flag. With the reviewer enabled, an approved PR can auto-merge and deploy; otherwise the
PR is left for a human. (Both are off by default — the demo just opens a PR.)
See CONTRIBUTING.md before opening a PR.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
