---
source: "https://github.com/alexander-akhmetov/sigil-hermes"
hn_url: "https://news.ycombinator.com/item?id=48433422"
title: "Show HN: Grafana Cloud observability plugin for Hermes Agent"
article_title: "GitHub - alexander-akhmetov/sigil-hermes: Grafana AI observability plugin for Hermes Agent · GitHub"
author: "oboroten"
captured_at: "2026-06-07T12:40:06Z"
capture_tool: "hn-digest"
hn_id: 48433422
score: 2
comments: 0
posted_at: "2026-06-07T10:12:45Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Grafana Cloud observability plugin for Hermes Agent

- HN: [48433422](https://news.ycombinator.com/item?id=48433422)
- Source: [github.com](https://github.com/alexander-akhmetov/sigil-hermes)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T10:12:45Z

## Translation

タイトル: HN を表示: Hermes Agent 用の Grafana Cloud 可観測性プラグイン
記事のタイトル: GitHub - alexander-akhmetov/sigil-hermes: Hermes Agent 用 Grafana AI 可観測性プラグイン · GitHub
説明: Hermes Agent 用の Grafana AI 可観測性プラグイン。 GitHub でアカウントを作成して、alexander-akhmetov/sigil-hermes の開発に貢献してください。

記事本文:
GitHub - alexander-akhmetov/sigil-hermes:Hermes Agent 用 Grafana AI 可観測性プラグイン · GitHub
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
アレクサンダー・アフメトフ
/
印章ヘルメス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
C

頌歌
アレクサンダー・アフメトフ/シギル・エルメス
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .github/ workflows .github/ workflows src/ hermes_plugin_sigil src/ hermes_plugin_sigil テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md img.png img.png llms.txt llms.txt plugin.yaml plugin.yaml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Hermes Agent 用の Grafana AI Observability プラグイン。 LLM 呼び出しとツールの実行を Sigil 世代として記録し、OTel トレースとメトリクスを出力します。
推奨: エージェントに任せてください
これをHermes (またはURLを取得できるClaude / Codex / Cursor / 同様のエージェント)に貼り付けます。
次のようにして、Grafana AI Observability プラグインをインストールして構成します。
https://raw.githubusercontent.com/alexander-akhmetov/sigil-hermes/main/llms.txt
エージェントは、 pip install、 ~/.hermes/config.yaml 、および Grafana Cloud からの資格情報の収集について説明します。また、デフォルトでどのような会話データが流れるか、および何かをオンにする前にそれを調整する方法についても説明します。
pip install git+https://github.com/alexander-akhmetov/sigil-hermes
hermes が実行されるのと同じ Python 環境にインストールします (どの hermes を確認するか)。次に、 ~/.hermes/config.yaml でプラグインを有効にします。
プラグイン:
有効:
- 印章
Hermes のプラグインを有効にする CLI は、pip でインストールされたプラグインをまだ認識しません。~/.hermes/plugins/ とバンドルされたディレクトリのみをスキャンします。回避策は、YAML を直接編集することです。
2 つの独立したチャネル (それぞれオプション): 正規の SIGIL_* スキーマに基づく世代、標準 OpenTelemetry OTEL_* スキーマに基づくトレースとメトリック。 URL とトークンは、Grafana アカウント https://grafana.com/orgs/{org} で見つけることができます。
Grafana Cloud アカウントをお持ちでない場合は、h で無料で作成できます。

ttps://grafana.com/auth/sign-up/create-user/ 。このプラグインを実行するには無料枠で十分です。
# 世代 → Sigil API (会話)
import SIGIL_ENDPOINT= " https://sigil-prod-<地域>.grafana.net "
エクスポート SIGIL_PROTOCOL=http
エクスポート SIGIL_AUTH_MODE=基本
import SIGIL_AUTH_TENANT_ID= " <grafana-cloud-stack-id> "
# スタック情報でこのトークンを見つけます → 「AI Observability」カード
# https://grafana.com/orgs/{org-id}/stacks/{stack-id}
import SIGIL_AUTH_TOKEN= " <sigil:書き込みトークン> "
# トレース + メトリクス → Grafana Cloud OTLP ゲートウェイ (標準 OTel 環境)
import OTEL_EXPORTER_OTLP_ENDPOINT= " https://otlp-gateway-prod-<地域>.grafana.net/otlp "
# OTEL_EXPORTER_OTLP_HEADERS はオプションです: 設定されていない場合、プラグインは派生します
# Authorization=Basic Base64("$SIGIL_AUTH_TENANT_ID:$SIGIL_AUTH_TOKEN") プラス
# X スコープの組織 ID。これは、OTLP ゲートウェイの基本認証ユーザー名が一致している場合にのみ機能します。
# SIGIL_AUTH_TENANT_ID と同じです。 OTLPインスタンスIDが異なる場合は設定してください
# そのユーザー名を明示的に使用します (シグナルごとにオーバーライドします)
# OTEL_EXPORTER_OTLP_TRACES_HEADERS / _METRICS_HEADERS):
# Base64 of "<otlp-instance-id>:<grafana-cloud-otlp-token>" — スタックの内容を参照してください
# 「OpenTelemetry」カード。
import OTEL_EXPORTER_OTLP_HEADERS= " Authorization=Basic <base64> "
オプション
変数
デフォルト
説明
SIGIL_AGENT_NAME
エルメス
世代ごとの gen_ai.agent.name
OTEL_SERVICE_NAME
エルメス
OTel リソース service.name 。これと OTEL_RESOURCE_ATTRIBUTES の service.name が両方とも設定されていない場合、プラグインのデフォルトは hermes です。
SIGIL_CONTENT_CAPTURE_MODE
いっぱい
フル / ツールコンテンツなし / メタデータのみ 。プラグインのデフォルトは完全であるため、ツールの引数と結果が表示されます。SDK 自体のデフォルトは no_tool_content で、エージェントの会話が空のままになります。
SIGIL_DEBUG
偽
詳細な SDK ログ
SIGIL_エルメス_SAMPLE_RATE
1.0
LLM とツールの割合

録音する呼び出し数、0.0 – 1.0
SIGIL_エルメス_MAX_CHARS
12000
編集されたペイロードの文字列ごとの切り捨ての上限
SIGIL_エルメス_OTEL_AUTO
本当の
アプリケーションがすでに TracerProvider / MeterProvider をインストールしている場合は、false を設定します。
SIGIL_エルメス_AGENT_VERSION
—
各世代にeffect_versionとしてスタンプされます（Sigilはバージョンごとのドリフトを追跡します）
検証する
SIGIL_DEBUG=真のエルメス
~/.hermes/logs/agent.log には次のように表示されます。
hermes-plugin-sigil: OTLP HTTP エクスポーターを使用して TracerProvider をインストールしました
hermes-plugin-sigil: OTLP HTTP エクスポーターを使用して MeterProvider をインストールしました
hermes-plugin-sigil: Sigil クライアントが初期化されました (世代 = 構成、otel = 構成)
hermes に何か質問してから、 Grafana Cloud -> Observability -> AI -> Conversations を確認してください。
Hermes Agent 用 Grafana AI 可観測性プラグイン
github.com/grafana/sigil-sdk
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
0.3.0
最新の
2026 年 6 月 7 日
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Grafana AI observability plugin for Hermes Agent. Contribute to alexander-akhmetov/sigil-hermes development by creating an account on GitHub.

GitHub - alexander-akhmetov/sigil-hermes: Grafana AI observability plugin for Hermes Agent · GitHub
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
alexander-akhmetov
/
sigil-hermes
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
alexander-akhmetov/sigil-hermes
main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .github/ workflows .github/ workflows src/ hermes_plugin_sigil src/ hermes_plugin_sigil tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md img.png img.png llms.txt llms.txt plugin.yaml plugin.yaml pyproject.toml pyproject.toml View all files Repository files navigation
Grafana AI Observability plugin for Hermes Agent . Records LLM calls and tool executions as Sigil generations and emits OTel traces + metrics.
Preferred: let your agent do it
Paste this into Hermes (or any Claude / Codex / Cursor / similar agent that can fetch URLs):
Install and configure the Grafana AI Observability plugin for me by following
https://raw.githubusercontent.com/alexander-akhmetov/sigil-hermes/main/llms.txt
The agent will walk you through pip install, ~/.hermes/config.yaml , and the credential collection from Grafana Cloud. It will also explain what conversation data flows by default and how to tune it before turning anything on.
pip install git+https://github.com/alexander-akhmetov/sigil-hermes
Install into the same Python environment hermes runs from ( which hermes to check). Then enable the plugin in ~/.hermes/config.yaml :
plugins :
enabled :
- sigil
Hermes's plugins enable CLI does not see pip-installed plugins yet — it only scans ~/.hermes/plugins/ and the bundled directory. Editing the YAML directly is the workaround.
Two independent channels, each optional: generations under the canonical SIGIL_* schema, traces and metrics under the standard OpenTelemetry OTEL_* schema. You can find URLs and tokens in your Grafana account: https://grafana.com/orgs/{org} .
If you do not have a Grafana Cloud account, you can create one for free at https://grafana.com/auth/sign-up/create-user/ . The free tier is enough to run this plugin.
# Generations → Sigil API (Conversations)
export SIGIL_ENDPOINT= " https://sigil-prod-<region>.grafana.net "
export SIGIL_PROTOCOL=http
export SIGIL_AUTH_MODE=basic
export SIGIL_AUTH_TENANT_ID= " <grafana-cloud-stack-id> "
# Find this token in your stack info → "AI Observability" card at
# https://grafana.com/orgs/{org-id}/stacks/{stack-id}
export SIGIL_AUTH_TOKEN= " <sigil:write token> "
# Traces + metrics → Grafana Cloud OTLP gateway (standard OTel envs)
export OTEL_EXPORTER_OTLP_ENDPOINT= " https://otlp-gateway-prod-<region>.grafana.net/otlp "
# OTEL_EXPORTER_OTLP_HEADERS is optional: when unset, the plugin derives
# Authorization=Basic base64("$SIGIL_AUTH_TENANT_ID:$SIGIL_AUTH_TOKEN") plus
# X-Scope-OrgID. That only works when the OTLP gateway's basic-auth username
# equals SIGIL_AUTH_TENANT_ID. If the OTLP instance ID differs, set it
# explicitly with that username (override per signal with
# OTEL_EXPORTER_OTLP_TRACES_HEADERS / _METRICS_HEADERS):
# Base64 of "<otlp-instance-id>:<grafana-cloud-otlp-token>" — see your stack's
# "OpenTelemetry" card.
export OTEL_EXPORTER_OTLP_HEADERS= " Authorization=Basic <base64> "
Optional
Variable
Default
Description
SIGIL_AGENT_NAME
hermes
Per-generation gen_ai.agent.name
OTEL_SERVICE_NAME
hermes
OTel resource service.name . Plugin defaults to hermes when this and OTEL_RESOURCE_ATTRIBUTES 's service.name are both unset.
SIGIL_CONTENT_CAPTURE_MODE
full
full / no_tool_content / metadata_only . Plugin defaults to full so tool args and results are visible — the SDK's own default is no_tool_content , which leaves agent conversations looking empty.
SIGIL_DEBUG
false
Verbose SDK logs
SIGIL_HERMES_SAMPLE_RATE
1.0
Fraction of LLM and tool calls to record, 0.0 – 1.0
SIGIL_HERMES_MAX_CHARS
12000
Per-string truncation cap for redacted payloads
SIGIL_HERMES_OTEL_AUTO
true
Set false if your application already installs a TracerProvider / MeterProvider
SIGIL_HERMES_AGENT_VERSION
—
Stamped on each generation as effective_version (Sigil tracks per-version drift)
Verify
SIGIL_DEBUG=true hermes
In ~/.hermes/logs/agent.log you should see:
hermes-plugin-sigil: installed TracerProvider with OTLP HTTP exporter
hermes-plugin-sigil: installed MeterProvider with OTLP HTTP exporter
hermes-plugin-sigil: Sigil client initialized (generations=configured, otel=configured)
Ask hermes anything, then check Grafana Cloud -> Observability -> AI -> Conversations .
Grafana AI observability plugin for Hermes Agent
github.com/grafana/sigil-sdk
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
0.3.0
Latest
Jun 7, 2026
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
