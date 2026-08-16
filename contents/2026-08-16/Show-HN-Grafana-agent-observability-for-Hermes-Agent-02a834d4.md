---
source: "https://github.com/alexander-akhmetov/grafana-agento11y-hermes"
hn_url: "https://news.ycombinator.com/item?id=49318128"
title: "Show HN: Grafana agent observability for Hermes Agent"
article_title: "GitHub - alexander-akhmetov/grafana-agento11y-hermes: Unofficial Grafana Agent Observability plugin for Hermes Agent · GitHub"
author: "eventuallyacat"
captured_at: "2026-08-16T09:17:59Z"
capture_tool: "hn-digest"
hn_id: 49318128
score: 2
comments: 0
posted_at: "2026-08-16T08:40:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Grafana agent observability for Hermes Agent

- HN: [49318128](https://news.ycombinator.com/item?id=49318128)
- Source: [github.com](https://github.com/alexander-akhmetov/grafana-agento11y-hermes)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T08:40:50Z

## Translation

タイトル: HN の表示: Hermes エージェントの Grafana エージェントの可観測性
記事タイトル: GitHub - alexander-akhmetov/grafana-agento11y-hermes: Hermes Agent 用の非公式 Grafana Agent Observability プラグイン · GitHub
説明: Hermes Agent 用の非公式 Grafana Agent Observability プラグイン - alexander-akhmetov/grafana-agento11y-hermes

記事本文:
GitHub - alexander-akhmetov/grafana-agento11y-hermes: Hermes Agent 用の非公式 Grafana Agent Observability プラグイン · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アレクサンダー・アフメトフ
/
グラファナ-アジェント11y-エルメス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
39 コミット 39 コミット .github .github スクリプト スクリプト src/ grafana_agento11y_hermes src/ grafana_agento11y_hermes テスト テスト .gitignore .gitignore CHANGELOG.md

CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md img.png img.png llms.txt llms.txt plugin.yaml plugin.yaml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Hermes Agent 用の Grafana Agent Observability プラグイン。 LLM 呼び出しとツールの実行を世代として記録し、OTel トレースとメトリックを生成します。
推奨: エージェントに任せてください
これをHermes (またはURLを取得できるClaude / Codex / Cursor / 同様のエージェント)に貼り付けます。
次のようにして、Grafana Agent Observability プラグインをインストールして構成します。
https://raw.githubusercontent.com/alexander-akhmetov/grafana-agento11y-hermes/main/llms.txt
エージェントは、 pip install、 ~/.hermes/config.yaml 、およびエージェント監視設定ページからの資格情報を案内します。また、デフォルトでどのような会話データが流れるか、および何かをオンにする前にそれを調整する方法についても説明します。
pip インストール grafana-agento11y-hermes
hermes が実行されるのと同じ Python 環境にインストールします (どの hermes を確認するか)。次に、 ~/.hermes/config.yaml でプラグインを有効にします。
プラグイン:
有効:
- エージェント11y
hermes のプラグインを有効にすると、CLI は pip でインストールされたプラグインをまだ認識しません。 ~/.hermes/plugins/ とバンドルされたディレクトリのみをスキャンします。回避策は、YAML を直接編集することです。
hermes-plugin-sigil からのアップグレード
パッケージ、モジュール、エントリポイント キー、および環境変数はすべて名前変更されました。
pip アンインストール hermes-plugin-sigil
pip インストール grafana-agento11y-hermes
アンインストールが必要です。新しいパッケージには名前が異なるため、pip は古いパッケージを置き換えるのではなく、一緒にインストールし、両方でプラグインを登録します。
~/.hermes/config.yaml のキーを sigil から Agento11y に変更します。古いキーは解決されなくなり、これがないと hermes はプラグインをロードしません。
SIG の名前を変更する

IL_* 環境変数を AGENTO11Y_* に変更し、サフィックスを維持します ( SIGIL_ENDPOINT は AGENTO11Y_ENDPOINT になります)。以下の [設定] のセットアップ ページでは、新しい名前の新しいブロックが提供されます。プラグインは今のところ古い名前を読み取り、名前を変更する内容をログに記録するため、アップグレードしても何も問題はありません。 SDK 自体はこれらを無視するため、SDK が修正されるとこのフォールバックはなくなります。
すべてはスタック内の 1 ページから得られます。
https://<スタック>.grafana.net/a/grafana-agento11y-app/setup
「環境変数としてコピー」をクリックします。
hermesが開始する環境にブロックを置きます。
AGENTO11Y_ENDPOINT=https://agento11y- < ... > .grafana.net
AGENTO11Y_PROTOCOL=http
AGENTO11Y_AUTH_MODE=基本
AGENTO11Y_AUTH_TENANT_ID=123456
AGENTO11Y_AUTH_TOKEN=glc_...
OTEL_EXPORTER_OTLP_ENDPOINT=https://otlp-gateway- < ... > .grafana.net/otlp
OTEL_EXPORTER_OTLP_HEADERS= ' Authorization=Basic <base64 of "123456:glc_..."> '
Grafana Cloud アカウントをお持ちでない場合は、 https://grafana.com/auth/sign-up/create-user/ で作成してください。無料枠で十分です。
AGENTO11Y_DEBUG=真のエルメス
~/.hermes/logs/agent.log には次のように表示されます。
grafana-agento11y-hermes: OTLP HTTP エクスポーターを使用して TracerProvider をインストールしました
grafana-agento11y-hermes: OTLP HTTP エクスポーターを使用して MeterProvider をインストールしました
grafana-agento11y-hermes: クライアントが初期化されました (世代 = 構成、otel = 構成)
hermes に何か質問してから、 Grafana Cloud -> Observability -> AI -> Conversations を確認してください。
非公式の Grafana Agent Observability プラグイン (Hermes Agent 用)
grafana.com/docs/grafana-cloud/observe-and-act/agent-observability/ トピック
Readme Apache-2.0 ライセンス アクティビティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Unofficial Grafana Agent Observability plugin for Hermes Agent - alexander-akhmetov/grafana-agento11y-hermes

GitHub - alexander-akhmetov/grafana-agento11y-hermes: Unofficial Grafana Agent Observability plugin for Hermes Agent · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
alexander-akhmetov
/
grafana-agento11y-hermes
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
39 Commits 39 Commits .github .github scripts scripts src/ grafana_agento11y_hermes src/ grafana_agento11y_hermes tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md img.png img.png llms.txt llms.txt plugin.yaml plugin.yaml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Grafana Agent Observability plugin for Hermes Agent . Records LLM calls and tool executions as generations and emits OTel traces + metrics.
Preferred: let your agent do it
Paste this into Hermes (or any Claude / Codex / Cursor / similar agent that can fetch URLs):
Install and configure the Grafana Agent Observability plugin for me by following
https://raw.githubusercontent.com/alexander-akhmetov/grafana-agento11y-hermes/main/llms.txt
The agent will walk you through pip install, ~/.hermes/config.yaml , and the credentials from the Agent Observability setup page. It will also explain what conversation data flows by default and how to tune it before turning anything on.
pip install grafana-agento11y-hermes
Install into the same Python environment hermes runs from ( which hermes to check). Then enable the plugin in ~/.hermes/config.yaml :
plugins :
enabled :
- agento11y
Hermes's plugins enable CLI does not see pip-installed plugins yet. It only scans ~/.hermes/plugins/ and the bundled directory. Editing the YAML directly is the workaround.
Upgrading from hermes-plugin-sigil
The package, the module, the entry-point key and the env vars were all renamed.
pip uninstall hermes-plugin-sigil
pip install grafana-agento11y-hermes
The uninstall is required. The new package has a different name, so pip installs it alongside the old one instead of replacing it, and both would register a plugin.
Change the key in ~/.hermes/config.yaml from sigil to agento11y . The old key no longer resolves, and hermes will not load the plugin without this.
Rename your SIGIL_* env vars to AGENTO11Y_* , keeping the suffix ( SIGIL_ENDPOINT becomes AGENTO11Y_ENDPOINT ). The setup page in Configure below gives you a fresh block with the new names. The plugin still reads the old names for now and logs what to rename, so nothing breaks the moment you upgrade. The SDK itself ignores them, so this fallback goes away once the SDK is fixed.
Everything comes from one page in your stack:
https://<stack>.grafana.net/a/grafana-agento11y-app/setup
Click Copy as environment variables .
Put the block in the environment hermes starts from.
AGENTO11Y_ENDPOINT=https://agento11y- < ... > .grafana.net
AGENTO11Y_PROTOCOL=http
AGENTO11Y_AUTH_MODE=basic
AGENTO11Y_AUTH_TENANT_ID=123456
AGENTO11Y_AUTH_TOKEN=glc_...
OTEL_EXPORTER_OTLP_ENDPOINT=https://otlp-gateway- < ... > .grafana.net/otlp
OTEL_EXPORTER_OTLP_HEADERS= ' Authorization=Basic <base64 of "123456:glc_..."> '
If you do not have a Grafana Cloud account, create one at https://grafana.com/auth/sign-up/create-user/ . The free tier is enough.
AGENTO11Y_DEBUG=true hermes
In ~/.hermes/logs/agent.log you should see:
grafana-agento11y-hermes: installed TracerProvider with OTLP HTTP exporter
grafana-agento11y-hermes: installed MeterProvider with OTLP HTTP exporter
grafana-agento11y-hermes: client initialized (generations=configured, otel=configured)
Ask hermes anything, then check Grafana Cloud -> Observability -> AI -> Conversations .
Unofficial Grafana Agent Observability plugin for Hermes Agent
grafana.com/docs/grafana-cloud/observe-and-act/agent-observability/ Topics
Readme Apache-2.0 license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
