---
source: "https://github.com/corryl/FindMyPipe"
hn_url: "https://news.ycombinator.com/item?id=48353876"
title: "FindMyPipe – Query Apple Find My from Linux for AI Agents"
article_title: "GitHub - corryl/FindMyPipe: Local CLI bridge for Apple Find My / Dovè polling from Linux, designed for AI agents. · GitHub"
author: "AgataVire"
captured_at: "2026-06-02T00:44:03Z"
capture_tool: "hn-digest"
hn_id: 48353876
score: 1
comments: 0
posted_at: "2026-06-01T08:02:23Z"
tags:
  - hacker-news
  - translated
---

# FindMyPipe – Query Apple Find My from Linux for AI Agents

- HN: [48353876](https://news.ycombinator.com/item?id=48353876)
- Source: [github.com](https://github.com/corryl/FindMyPipe)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T08:02:23Z

## Translation

タイトル: FindMyPipe – AI エージェント用に Linux から Apple Find My をクエリする
記事のタイトル: GitHub - corryl/FindMyPipe: AI エージェント向けに設計された、Linux からの Apple Find My / Dovè ポーリング用のローカル CLI ブリッジ。 · GitHub
説明: AI エージェント向けに設計された、Linux からの Apple Find My / Dovè ポーリング用のローカル CLI ブリッジ。 - コリル/FindMyPipe

記事本文:
GitHub - corryl/FindMyPipe: AI エージェント向けに設計された、Linux からの Apple Find My / Dovè ポーリング用のローカル CLI ブリッジ。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
コリル
/
FindMyPipe
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット アセット アセット リファレンス リファレンス スクリプト スクリプト src/ findmy_agent_bridge src/ findmy_agent_bridge テスト テスト .gitignore .gitignore ライセンス ライセンス README.it.md README.it.md README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml すべてのファイルを表示リポジトリ ファイルのナビゲーション
AI エージェントと人間向けの、CLI 経由の Apple デバイスの位置情報。
FindMyPipe は、Linux から Apple Find My にクエリを実行し、デバイスの場所を構造化された JSON として返すローカル CLI ブリッジであり、AI エージェント、シェル スクリプト、自動化パイプラインに対応します。
findmypipe は、単にデバイスがどこにあるかを知るだけでなく、位置認識ワークフローを強化するように設計されています。デバイスの座標が与えられると、AI エージェントが外部サービスにクエリを実行して、近くの関心のある場所 (レストラン、ホテル、病院、交通機関の停留所) を検索し、より豊かでコンテキストを認識したエクスペリエンスを構築できます。
「私のiPhoneはどこですか？」 · 「私の MacBook は家にありますか?」・「AirPods の近くのコーヒーショップを見つけてください」 ・「iPad から一番近い病院はどこですか?」
特徴
説明
🔍 Apple デバイスを見つける
iPhone、iPad、Mac、AirPods — iCloud 経由のリアルタイム位置
📍 近くの観光スポット
POI API (Google Places、OSM、Foursquare) に座標をフィードして、デバイスの周囲にあるものを見つけます
🖥️ プロフェッショナル CLI
標準出力上のクリーンな構造化された JSON 出力、パイプと解析が簡単
🔐 プライバシー第一
ログ内の認証情報が編集され、デバイス ID が SHA-256 ハッシュに置き換えられます
📦 ゼロ構成
モックモードはすぐに使用できます - 探索するために Apple ID は必要ありません
⏱️ オプションのキャッシュ
iCloud API への負荷を回避するための TTL を構成可能
🧹 スマートフィルター
--skip-offline と --max-age により、最新の関連データのみを保持します。
🔄 完全な 2FA lo

ジン
完全な Apple 認証フロー、インタラクティブな 2FA を含む
🤖 AIエージェント対応
AgentSkills.io 準拠のスキルとして設計されており、任意のエージェント フレームワークにプラグインできます。
🤖 AgentSkills.io スキルとは何ですか?
FindMyPipe には、AgentSkills.io 標準に準拠した SKILL.md 定義が付属しています。これは次のことを意味します:
この標準をサポートする AI エージェント フレームワークは、このツールを自動的に検出して呼び出すことができます。
スキルは自己記述的です。入力、出力、機能を機械可読形式で宣言します。
Hermes 、カスタム LLM パイプライン、または任意のシェル対応オーケストレーターなどのエージェントとシームレスに統合します。
これは OpenAPI 仕様ですが、AI エージェント ツール用と考えてください。
🗺️ デバイスの近くの興味のあるスポットを見つけます
# iPhone の座標を取得する
COORDS= $( findmy-agent Karaoke " iPhone " --json | jq -r ' "\(.asset.latitude),\(.asset.longitude)" ' )
# 近くのレストランをクエリする (Google Places API の例)
カール " https://maps.googleapis.com/maps/api/place/nearbysearch/json \
?location= $COORDS &radius=500&type=restaurant&key= $GOOGLE_API_KEY " | jq ' .results[].name '
📍 デバイスが家にあるかどうかを確認する
findmy-agent は「MacBook」を検索します --json | jq ' .asset | {名前、緯度、経度、最後に見た} '
🔔 cronによる定期監視
# 15 分ごとに、すべてのオンライン デバイスをログに記録します
* /15 * * * * findmy-agent list --json --skip-offline >> /var/log/findmy.log
🤖 AIエージェントの統合（Hermes）
# ~/.hermes/config.yaml
ツール:
- 名前 : findmy_locate
cmd : " findmy-agent の場所 \" {{name}} \" --json "
- 名前 : findmy_list
cmd: " findmy-agent list --json --skip-offline "
- 名前 : findmy_doctor
cmd: " findmy-agent ドクター --json "
📦 インストール
「Find My」が有効になっている Apple ID (ライブモードのみ)
git clone https://github.com/corryl/FindMyPipe.git
cd FindMyPipe
python3 -m venv .venv
#モ

ck モードのみ (すぐに動作し、Apple ID は必要ありません)
.venv/bin/pip install -e ' .[dev] '
# iCloud のライブサポートあり
.venv/bin/pip install -e ' .[dev,live] '
# インストールを確認する
.venv/bin/findmy-agent ドクター --json
期待される出力:
{
"キャッシュ" : { "有効" : false 、 "状態" : "空" 、 "ttl_秒" : 0 }、
"live_probe_available" : true 、
"ok" : true 、
"プロバイダー" : "モック" ,
"secrets_redacted" : true 、
"トランスポート" : "ローカル"
}
⚙️ 構成
エクスポート FINDMY_AGENT_PROVIDER= " icloud "
エクスポート FINDMY_APPLE_ID= " [編集済] "
エクスポート FINDMY_APPLE_PASSWORD= " [編集済み] "
🔐 アプリ固有のパスワードを使用します — Apple ID → セキュリティ → アプリ固有のパスワード。
メインの Apple ID パスワードは決して使用しないでください。
エクスポート FINDMY_COOKIE_DIR= " $HOME /.local/state/findmypipe/icloud "
エクスポート FINDMY_CACHE_TTL= " 300 " # 秒 (0 = 無効)
import FINDMY_CACHE_FILE= " $HOME /.local/state/findmypipe/cache.json "
対話型ログイン (初回実行)
findmy-agent ログイン --json
# Apple から 2FA コードが送信されたら、それを入力します
🖥️ CLI リファレンス
すべてのコマンドは構造化出力の --json を受け入れます。
ブリッジのステータスとプロバイダーの可用性を確認します。
findmy-agent ドクター --json
findmy-agent Doctor --provider icloud --json
findmy-エージェントリスト
すべてのデバイスを位置、バッテリーレベル、最後に見た時刻とともにリストします。
findmy-agent リスト --json
findmy-agent list --json --skip-offline --max-age 30
findmy-agent list --json --include-raw
出力例
{
「資産」: [{
"id" : " icloud:a1b2c3d4e5f6a7b8 " ,
"名前" : " iPhone " ,
"種類" : "デバイス" ,
"プロバイダー" : " icloud " 、
「緯度」 : 45.1234 、
「経度」 : 9.5678 、
"精度_m" : 15.0 、
「バッテリー」：0.85、
"battery_status" : "充電済み" ,
"last_seen" : " 2025-05-30T12:34:56Z " ,
"location_is_old" : false
}]
}
findmy-agent を見つける
名前または ID で特定のデバイスを検索します (case-insens)

アクティブ）。
findmy-agent は「iPhone」を検索します --json
findmy-agent は「AirPods」を検索します --json --skip-offline --max-age 60
findmy-agent ログイン
2FA によるインタラクティブな iCloud 認証。
findmy-agent ログイン --json
共通オプション
オプション
説明
--プロバイダー
モック (デフォルト) または icloud
--json
構造化された JSON 出力
--include-raw
編集された生のペイロードを含める (デバッグ)
--max-age <分>
N 分より古い位置をフィルターします
--スキップオフライン
オフラインデバイスを除外する
エラーフォーマット
すべてのエラーは、一貫した構造化フォーマットを返します。
{ "error" : " FINDMY_APPLE_ID が設定されていません " 、 "error_type" : "configuration_error " 、 "ok" : false 、 "secret_safe" : true }
🔐 プライバシーとセキュリティ
HTTP サーバーなし - すべての I/O は標準入出力上にあり、オープンポートもデーモンもありません
Webhook なし - Apple サーバーへのアウトバウンド ポーリングのみ
認証情報は記録されません - パスワード、2FA コード、トークンは常に編集されます
ハッシュされたデバイス ID — すべての出力で実際の識別子が SHA-256 ハッシュに置き換えられます
生のペイロードはデフォルトで非表示になります - 明示的な --include-raw フラグが必要です
制限されたファイル権限 — ディレクトリ 0700 、ファイル 0600
~/.local/state/findmypipe/
§── icloud/ # Cookie とセッション (0700)
━──cache.json # オプションのキャッシュ (0600)
🏗️ アーキテクチャ
┌───────────────────┐
│ AIエージェント / シェルスクリプト / ターミナル │
━━━━━━━━━━━━━━━━━━┘
│ CLI (標準入力/標準出力/標準エラー出力)
┌───────▼─────────────┐
│findmyエージェント│
│ │
│ ┌─────┐ ┌───────┐ │
│ │ CLI │ │ キャッシュ

え │ │
│ │ (タイパー) │ │ (JSON ファイル) │ │
│ └───┬─────┘ └───────┘ │
│ │ │
│ ┌────▼──────┐ │
│ │ コア │ │
│ └────┬──────┘ │
│ │ │
│ ┌────▼──────┐ │
│ │ プロバイダー │ │
│ │ モック│iCld │ │
│ └───────┘ │
━━━━━━━━━━━━━━━━━━┘
│ HTTPS (送信のみ)
┌───────▼─────────────┐
│ Apple iCloud API │
━━━━━━━━━━━━━━┘
🧪 テスト
.venv/bin/pytest -q --tb=short
テスト/test_cache.py ......
テスト/test_cli.py ....
テスト/test_core.py ......
テスト/test_icloud_provider.py ...
テスト/test_provider_factory.py ...
31 合格
📋 制限とロードマップ
アスペクト
ステータス
iPhone、iPad、Mac、AirPods
✅ サポートされています
エアタグ / アイテム
⏳ まだサポートされていません
Linux
✅ テスト済み
macOS
✅ うまくいくはずです
インタラクティブな 2FA
✅ サポートされています
POI統合（組み込み）
⏳ 計画中
📄ライセンス
AI コミュニティ向けの ❤️ で構築 · AgentSkills.io · スキル定義
地元。安全な。プライベート。あなたのデータはあなたの管理下にあります。
Linux からの Apple Find My / Dovè ポーリング用のローカル CLI ブリッジ。AI エージェント用に設計されています。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人的なことを共有しないでください

情報

## Original Extract

Local CLI bridge for Apple Find My / Dovè polling from Linux, designed for AI agents. - corryl/FindMyPipe

GitHub - corryl/FindMyPipe: Local CLI bridge for Apple Find My / Dovè polling from Linux, designed for AI agents. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
corryl
/
FindMyPipe
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit assets assets references references scripts scripts src/ findmy_agent_bridge src/ findmy_agent_bridge tests tests .gitignore .gitignore LICENSE LICENSE README.it.md README.it.md README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml View all files Repository files navigation
Your Apple devices' location, via CLI — for AI agents and humans.
FindMyPipe is a local CLI bridge that queries Apple Find My from Linux and returns device locations as structured JSON , ready for AI agents, shell scripts, and automation pipelines.
Beyond simply knowing where your devices are, findmypipe is designed to power location-aware workflows : given the coordinates of a device, an AI agent can query external services to find nearby points of interest — restaurants, hotels, hospitals, transit stops — and build richer, context-aware experiences.
"Where's my iPhone?" · "Is my MacBook at home?" · "Find a coffee shop near my AirPods" · "What's the nearest hospital to my iPad?"
Feature
Description
🔍 Locate Apple devices
iPhone, iPad, Mac, AirPods — real-time position via iCloud
📍 Points of interest nearby
Feed coordinates to any POI API (Google Places, OSM, Foursquare) to find what's around your device
🖥️ Professional CLI
Clean structured JSON output on stdout, easy to pipe and parse
🔐 Privacy-first
Redacted credentials in logs, device IDs replaced by SHA-256 hash
📦 Zero configuration
Mock mode works out of the box — no Apple ID needed to explore
⏱️ Optional cache
Configurable TTL to avoid hammering the iCloud API
🧹 Smart filters
--skip-offline and --max-age to keep only fresh, relevant data
🔄 Full 2FA login
Complete Apple authentication flow, interactive 2FA included
🤖 AI agent-ready
Designed as an AgentSkills.io -compliant skill — plug it into any agent framework
🤖 What is an AgentSkills.io Skill?
FindMyPipe ships with a SKILL.md definition conforming to the AgentSkills.io standard. This means:
Any AI agent framework that supports the standard can automatically discover and invoke this tool
The skill is self-describing : it declares its inputs, outputs, and capabilities in a machine-readable format
It integrates seamlessly with agents like Hermes , custom LLM pipelines, or any shell-capable orchestrator
Think of it as an OpenAPI spec, but for AI agent tools.
🗺️ Find points of interest near your device
# Get your iPhone's coordinates
COORDS= $( findmy-agent locate " iPhone " --json | jq -r ' "\(.asset.latitude),\(.asset.longitude)" ' )
# Query nearby restaurants (example with Google Places API)
curl " https://maps.googleapis.com/maps/api/place/nearbysearch/json \
?location= $COORDS &radius=500&type=restaurant&key= $GOOGLE_API_KEY " | jq ' .results[].name '
📍 Check if a device is home
findmy-agent locate " MacBook " --json | jq ' .asset | {name, latitude, longitude, last_seen} '
🔔 Periodic monitoring via cron
# Every 15 minutes, log all online devices
* /15 * * * * findmy-agent list --json --skip-offline >> /var/log/findmy.log
🤖 AI agent integration (Hermes)
# ~/.hermes/config.yaml
tools :
- name : findmy_locate
cmd : " findmy-agent locate \" {{name}} \" --json "
- name : findmy_list
cmd : " findmy-agent list --json --skip-offline "
- name : findmy_doctor
cmd : " findmy-agent doctor --json "
📦 Installation
An Apple ID with Find My enabled (live mode only)
git clone https://github.com/corryl/FindMyPipe.git
cd FindMyPipe
python3 -m venv .venv
# Mock mode only (works immediately, no Apple ID needed)
.venv/bin/pip install -e ' .[dev] '
# With live iCloud support
.venv/bin/pip install -e ' .[dev,live] '
# Verify installation
.venv/bin/findmy-agent doctor --json
Expected output:
{
"cache" : { "enabled" : false , "state" : " empty " , "ttl_seconds" : 0 },
"live_probe_available" : true ,
"ok" : true ,
"provider" : " mock " ,
"secrets_redacted" : true ,
"transport" : " local "
}
⚙️ Configuration
export FINDMY_AGENT_PROVIDER= " icloud "
export FINDMY_APPLE_ID= " [REDACTED] "
export FINDMY_APPLE_PASSWORD= " [REDACTED] "
🔐 Use an app-specific password — Apple ID → Security → App-Specific Passwords.
Never use your main Apple ID password.
export FINDMY_COOKIE_DIR= " $HOME /.local/state/findmypipe/icloud "
export FINDMY_CACHE_TTL= " 300 " # seconds (0 = disabled)
export FINDMY_CACHE_FILE= " $HOME /.local/state/findmypipe/cache.json "
Interactive login (first run)
findmy-agent login --json
# Type your 2FA code when Apple sends it
🖥️ CLI Reference
All commands accept --json for structured output.
Check bridge status and provider availability.
findmy-agent doctor --json
findmy-agent doctor --provider icloud --json
findmy-agent list
List all devices with location, battery level, and last seen time.
findmy-agent list --json
findmy-agent list --json --skip-offline --max-age 30
findmy-agent list --json --include-raw
Example output
{
"assets" : [{
"id" : " icloud:a1b2c3d4e5f6a7b8 " ,
"name" : " iPhone " ,
"kind" : " device " ,
"provider" : " icloud " ,
"latitude" : 45.1234 ,
"longitude" : 9.5678 ,
"accuracy_m" : 15.0 ,
"battery" : 0.85 ,
"battery_status" : " charged " ,
"last_seen" : " 2025-05-30T12:34:56Z " ,
"location_is_old" : false
}]
}
findmy-agent locate
Search for a specific device by name or ID (case-insensitive).
findmy-agent locate " iPhone " --json
findmy-agent locate " AirPods " --json --skip-offline --max-age 60
findmy-agent login
Interactive iCloud authentication with 2FA.
findmy-agent login --json
Common options
Option
Description
--provider
mock (default) or icloud
--json
Structured JSON output
--include-raw
Include redacted raw payload (debug)
--max-age <min>
Filter positions older than N minutes
--skip-offline
Exclude offline devices
Error format
All errors return a consistent structured format:
{ "error" : " FINDMY_APPLE_ID not set " , "error_type" : " configuration_error " , "ok" : false , "secret_safe" : true }
🔐 Privacy & Security
No HTTP server — all I/O on stdio, no open ports, no daemon
No webhooks — outbound polling only toward Apple's servers
Credentials never logged — password, 2FA codes, and tokens always redacted
Hashed device IDs — real identifiers replaced by SHA-256 hash in all output
Raw payload hidden by default — requires explicit --include-raw flag
Restrictive file permissions — directories 0700 , files 0600
~/.local/state/findmypipe/
├── icloud/ # Cookies & sessions (0700)
└── cache.json # Optional cache (0600)
🏗️ Architecture
┌──────────────────────────────────────┐
│ AI Agent / Shell Script / Terminal │
└────────────┬─────────────────────────┘
│ CLI (stdin/stdout/stderr)
┌────────────▼─────────────────────────┐
│ findmy-agent │
│ │
│ ┌──────────┐ ┌──────────────────┐ │
│ │ CLI │ │ Cache │ │
│ │ (Typer) │ │ (JSON file) │ │
│ └────┬─────┘ └──────────────────┘ │
│ │ │
│ ┌────▼──────┐ │
│ │ Core │ │
│ └────┬──────┘ │
│ │ │
│ ┌────▼──────┐ │
│ │ Provider │ │
│ │ Mock│iCld │ │
│ └───────────┘ │
└────────────┬─────────────────────────┘
│ HTTPS (outbound only)
┌────────────▼─────────────────────────┐
│ Apple iCloud API │
└──────────────────────────────────────┘
🧪 Tests
.venv/bin/pytest -q --tb=short
tests/test_cache.py ..........
tests/test_cli.py ....
tests/test_core.py ........
tests/test_icloud_provider.py ...
tests/test_provider_factory.py ...
31 passed
📋 Limits & Roadmap
Aspect
Status
iPhone, iPad, Mac, AirPods
✅ Supported
AirTag / Items
⏳ Not yet supported
Linux
✅ Tested
macOS
✅ Should work
Interactive 2FA
✅ Supported
POI integration (built-in)
⏳ Planned
📄 License
Built with ❤️ for the AI community · AgentSkills.io · Skill Definition
Local. Secure. Private. Your data, under your control.
Local CLI bridge for Apple Find My / Dovè polling from Linux, designed for AI agents.
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
