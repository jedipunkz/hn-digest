---
source: "https://github.com/blitzcrieg1/agentmetry"
hn_url: "https://news.ycombinator.com/item?id=48909516"
title: "Agentmetry, catch your AI coding agent reading –/.ssh and phoning home"
article_title: "GitHub - blitzcrieg1/agentmetry: Agentmetry — a local-first SIEM and flight recorder for AI coding agents. Captures every tool call, approval, and denial from Cursor, Claude Code, Codex and Antigravity into a MITRE ATT&CK-mapped audit trail, with correlated sequence detection (credential access -> n\n[truncated]"
author: "blitzcrieg1"
captured_at: "2026-07-14T17:03:57Z"
capture_tool: "hn-digest"
hn_id: 48909516
score: 1
comments: 0
posted_at: "2026-07-14T16:41:48Z"
tags:
  - hacker-news
  - translated
---

# Agentmetry, catch your AI coding agent reading –/.ssh and phoning home

- HN: [48909516](https://news.ycombinator.com/item?id=48909516)
- Source: [github.com](https://github.com/blitzcrieg1/agentmetry)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T16:41:48Z

## Translation

タイトル: Agentmetry、AI コーディング エージェントが –/.ssh を読んで家に電話しているのをキャッチ
記事のタイトル: GitHub - blitzcrieg1/agentmetry: Agentmetry — ローカルファーストの SIEM および AI コーディング エージェント用のフライト レコーダー。 Cursor、Claude Code、Codex、Antigravity からのすべてのツール呼び出し、承認、拒否を、相関シーケンス検出 (資格情報アクセス -> n
[切り捨てられた]
説明: Agentmetry — AI コーディング エージェント用のローカル ファースト SIEM およびフライト レコーダー。 Cursor、Claude Code、Codex、Antigravity からのすべてのツール呼び出し、承認、拒否を MITRE ATT&CK マップ監査証跡にキャプチャし、相関シーケンス検出 (資格情報アクセス -> ネットワーク出力) とフォワーダーを使用して、
[切り捨てられた]

記事本文:
GitHub - blitzcrieg1/agentmetry: Agentmetry — ローカルファーストの SIEM および AI コーディング エージェント用のフライト レコーダー。 Cursor、Claude Code、Codex、Antigravity からのすべてのツール呼び出し、承認、拒否を MITRE ATT&CK マップ監査証跡にキャプチャし、相関シーケンス検出 (資格情報へのアクセス -> ネットワーク出力) と Loki、Elastic、Splunk へのフォワーダーを行います。 Apache-2.0。ゼロエグレス。 · GitHub
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
別の場所でサインアウトしました

タブまたはウィンドウ。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
電撃戦1
/
エージェントメトリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
119 コミット 119 コミット .agents .agents .claude .claude .codex .codex .github/ workflows .github/ workflows アダプター アダプター アプリ アプリ docs docs infra/ loki infra/ loki ポリシー/ dlp ポリシー/ dlp スクリプト スクリプト テスト テスト vault vault .cursorignore .cursorignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md docker-compose.loki.yml docker-compose.loki.yml docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントメトリ: AI エージェント用の SIEM
AI エージェント ツール用のオープンソース フライト レコーダーとセキュリティ レイヤー。
すべてのツール呼び出し、すべての拒否、すべての人間による承認は、ハッシュ化され、関連付けられ、所有する JSONL 証跡に保存されます。
オンデマンドで再生。 SIEM が必要な場合は、Loki、Elastic、または Splunk に転送します。
クイックスタート ·
ドキュメント ·
スキーマ ·
ロードマップ ·
セキュリティ
すべての AI エージェント ツール呼び出しには、発生時に MITRE ATT&CK のタグが付けられます。
無害に見える一連の通話が攻撃につながる場合、相関アラートを発します。
🚧 パブリック アルファ : コアのキャプチャ、リプレイ、SIEM 転送は、初期の探索に使用できます。 API と統合サーフェスは急速に進化する可能性があります。
自律エージェントがツールを実行すると、ほとんどのスタックにはインシデント対応者に渡せるものは何も保存されません。ログにはプロセスが表示されます。意図、セッションの境界、人間が承認する内容は示されません。

d .
Agentmetry は、AI エージェント用のオープンソース エンドポイント フライト レコーダーです。完全にマシン上で実行されるように構築されており、オプションで既に運用している SIEM への転送も可能です。
管理された AI エージェントのための不変のオペレーター所有の監査証跡 - ベンダー クラウドではなく、IDE ライフサイクル境界および MCP ワイヤーでツールの実行をキャプチャします。
IDE ライフサイクル フック (Cursor、Claude Code、Codex、Antigravity) および MCP stdio 監査プロキシを介したエージェント ツール呼び出しのインターセプト
MITRE ATT&CK エンリッチメントと SHA-256 引数ハッシュを使用して、すべてのイベントを正規スキーマ v1.1.0 に正規化します。
単一のイベントでは明らかにできない、相関する動作シーケンスの検出 (資格情報のエクスポート、ガードレールのバイパス、偵察と取得)
ローカル正規表現 DLP エンジン (ログ モードまたはブロック モード) を使用したフック境界でのシークレットと PII のブロック
クラウドを記録システムにすることなく、同じ JSONL 証跡を Loki、Elastic ECS、Splunk HEC、または汎用 Webhook に転送します
Agentmetry は CASB やシャドウ AI スパイではありません。接続したエージェントを記録します。問題がブラウザーで管理されていない ChatGPT である場合は、フライト レコーダーではなく、ネットワーク/エンドポイント ポリシーが必要です。
Agentmetry は完全にローカルで実行されます。監査証跡は、明示的に転送しない限り、マシンから離れることはありません。
最初に何かをキャッチするのを見てください (30 秒)
サーバーも API キーも設定もありません。クローンを作成して実行します。
git clone https://github.com/blitzcrieg1/agentmetry.git && cd Agentmetry
pip install -r apps/orchestrator/requirements.txt
Python スクリプト/demo.py
実際の取り込み API を通じてエージェント セッションを再生します。エージェントは SSH を読み取ります。
秘密キーを取得し、AWS キーを含むコマンドを実行して、URL を取得します。エージェントメトリ
各呼び出しを MITRE ATT&CK でタグ付けし、DLP で AWS キーをキャッチします (ルールを保存し、
決して値ではありません)、そして、尋ねられることなく、読み取られたキーを次のキーと関連付けます。

ネットワークを呼び出し、CRITICAL credential-exfil 検出を起動します。
これらのイベントのどれもアラートではありません。順序は次のとおりです。それが全体です
製品を 1 つの画面で表示します。
ダッシュボードにストーリーが含まれているのを確認する
python scripts/demo_dashboard.py # 5 セッション + 4 検出をシード、http://127.0.0.1:8010/ を提供
1 つのコマンドで現実的なデモ証跡を作成し、ダッシュボードをローカルで提供します。
API キー、クラウドなし。フラグの付いたセッションをクリックすると、すぐに検出が行われます。
その下に完全なイベントがドリルで開きます。
各ビューが何をどのように表示するかについては、ダッシュボード ツアーをご覧ください。
それを読むために。
要件
バージョン
パイソン
3.10+
Node.js
18+
1. クローンを作成してインストールする
git clone https://github.com/blitzcrieg1/agentmetry.git
CD エージェントメトリ
# Python オーケストレーター
CD アプリ\オーケストレーター
Python - m venv .venv
.\.venv\Scripts\activate
pip install - e " .[dev] "
.env.example .env をコピー
CD ..\..
# Next.js ダッシュボード
cd アプリ\ダッシュボード
npmインストール
CD ..\..
2. フライトレコーダーを起動する
スクリプト\start-dev.bat
ダッシュボード → http://localhost:3000 ・Orchestrator API → http://localhost:8000
powershell - 実行ポリシーのバイパス - ファイル scripts\install_cursor_hooks.ps1
powershell - 実行ポリシーのバイパス - ファイル scripts\install_claude_hooks.ps1
フックが読み込まれるように、カーソル/クロード コードを完全に終了して再起動します。
Python スクリプト\agentmetry_ingest.py セルフテスト
イベントは数秒以内にダッシュボードのフライト レコーダーに表示されます。
エージェントがツールを実行すると、Agentmetry は自動的に次のことを行います。
引数がフック プロセスを離れる前に、ライフサイクル フックまたは MCP ツール/呼び出しをインターセプトします。
ツール引数 (SHA-256) をハッシュし、コマンド文字列内のインライン シークレットをスクラブします。
MITREの戦術/テクニックのマッピングとセッションの相関関係で各イベントを強化します
正規の JSONL をローカルに保存します ( Audit-forward.jsonl ) — フック パスの記録システム
マルチSを検出

セッションのタイムライン全体にわたる行動パターンを観察する
SIEM シンクとアラート Webhook に転送します (オプション、ベストエフォート)
フローチャートTB
サブグラフ Capture["キャプチャ レイヤ (層 A + B)"]
HOOKS["IDE ライフサイクル フック<br/>カーソル · クロード · コーデックス · 反重力"]
PROXY["MCP 監査プロキシ<br/>mcp_audit_proxy.py"]
終わり
サブグラフ ゲート["ローカル セキュリティ ゲート"]
DLP["DLP スキャナー<br/>正規表現ルール"]
HASH["Arg ハッシュ + 秘密のスクラブ"]
終わり
サブグラフ Core["オーケストレーター:8000"]
INGEST["POST /api/v1/audit/ingest"]
CANON["正規スキーマ v1.1.0<br/>MITRE エンリッチメント"]
DETECT["シーケンス検出エンジン"]
OUTBOX[("SQLite アウトボックス<br/>イベント.db")]
終わり
サブグラフ Output["出力"]
JSONL["audit-forward.jsonl"]
DASH[「ダッシュボード<br/>フライトレコーダー + アナリティクス」]
SIEM["Loki · Elastic · Splunk · Webhook"]
終わり
フック --> DLP
プロキシ --> DLP
DLP -->|許可|ハッシュ
DLP -->|拒否|摂取する
ハッシュ --> 取り込み
インジェスト --> キヤノン
キヤノン --> アウトボックス
キヤノン --> JSONL
キヤノン --> 検出
JSONL --> ダッシュ
JSONL --> SIEM
読み込み中
キャプチャパス
フローチャート LR
サブグラフ TierB["Tier B — IDE フック"]
C[「カーソル」]
CL[「クロード・コード」]
AG[「反重力」]
CX[「コーデックス」]
終わり
サブグラフ TierA["Tier A — MCP プロキシ"]
MCP[「任意の MCP クライアント」]
WRAP["監査プロキシがサーバーコマンドをラップする"]
終わり
INGEST["agentmetry_ingest.py → /audit/ingest"]
C --> 取り込み
CL --> 取り込み
AG --> 摂取
CX --> 取り込み
MCP --> ラップ --> 摂取
読み込み中
コンポーネント
パス
役割
フッククライアント
スクリプト/agentmetry_ingest.py
IDE ライフサイクル イベントを正規ペイロードにマップします。プロセス中のハッシュ引数
MCPプロキシ
apps/orchestrator/tools/mcp_audit_proxy.py
任意の stdio MCP サーバーをラップします。すべてのツール/呼び出し + エラーをログに記録します
APIの取り込み
コア/監査/ingest.py
ペイロードの正規化、承認の推論 ( inferred:* )、シンクの書き込み
DLP エンジン
コア/監査/dlp/
ツール引数の正規表現スキャン (バリデータ、例: Luhn)。実行前にブロックまたはログに記録する
検出

イオンエンジン
コア/監査/検出/
相関シーケンスがセッションのイベント タイムラインを支配します
シンク
コア/監査/シンク.py
ファイル、Webhook、Elastic ECS、Splunk HEC
リプレイ
コア/監査/再生.py
ローカル送信ボックスからの ASCII タイムラインの再構築
正規のイベント
実行するたびに、型指定された SIEM 対応の JSON が出力されます。単一の tools_called 行:
{
"schema_version" : " 1.1.0 " 、
"correlation_id" : " thread-8892 " 、
"timestamp_utc" : " 2026-07-12T09:14:22.041+00:00 " ,
"アクター" : { "タイプ" : " ユーザー " 、 "id" : " dev_01 " 、 "ロール" : " オペレーター " },
"アクション" : { "タイプ" : " 呼び出されたツール " 、 "結果" : " 成功 " },
"エージェント" : { "名前" : " カーソル " 、 "スキル ID" : " " }、
「ツール」: {
"qualified" : " vault_fs.read_file " ,
"サーバー" : " vault_fs " ,
"input_hash" : " e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 " ,
"parameters_redacted" : true 、
"mitre" : { "tactic" : " 収集 (TA0009) " , "technique" : " ローカル システムからのデータ (T1005) " }
}、
"モデル" : { "id" : " クロード-3-5-ソネット " 、 "プロバイダー" : " 人間性 " }
}
完全なスキーマ → docs/agentmetry-event-schema.md
Agentmetry は、接続したエージェント (IDE フックまたは MCP プロキシ) を記録します。目に見えないものに対しては正直です。
🎥 フライトレコーダー
動的な列、ドラッグ アンド ドロップ レイアウト、CSV エクスポート、セッション ドリルダウンを備えたライブ監査テール
📊 分析とプロセスツリー
セッションレベルのチャート、MITRE戦術の内訳、React Flowの水平タイムライン
🔍 行動検出
相関シーケンス ルール — 資格情報の Exfil、ガードレール バイパス、Recon-then-grab
🛡️ ローカル DLP
Regex スキャナーは、ツールの実行前に AWS キー、GitHub トークン、Slack トークン、PII をブロックします
🎯 MITRE ATT&CK マッピング
すべての正規イベントのツールごとの戦術/技術タグ
🔐 引数のハッシュ化
デフォルトでツール引数の SHA-256 — プレーンテキストがフックからネットワークを通過することはありません
📡

SIEMネイティブのエクスポート
Elastic ECS、Splunk HEC、Loki/LogQL、汎用 Webhook、拒否時のアラート Webhook
🔁 リプレイと証拠
ASCII セッション タイムライン + 改ざん証拠パックのエクスポート
👥 マルチIDEのサポート
Cursor、Claude Code、Codex、Antigravity — グローバルフックインストールスクリプト
統合
カテゴリ
本日サポートされました
ロードマップ
IDE/エージェントホスト
カーソル · クロード コード · コーデックス · 反重力
ウィンドサーフィン · VS Code 副操縦士
MCPトランスポート
Stdio 監査プロキシ (任意の MCP サーバー コマンドをラップ)
SSE / ストリーミング可能な HTTP プロキシ
可観測性 / SIEM
Loki、Grafana、Elastic ECS、Splunk HEC、汎用 Webhook
Datadog · New Relic
検出フォーマット
エンジン内シーケンス ルール · LogQL · Elastic · Splunk · Sigma パック
STIX/TAXII輸出
ポリシーエンジン
正規表現 DLP マニフェスト (policy/dlp/)
OPA / Rego のコードとしてのポリシー
コンプライアンス文書
ISO 42001 マッピング・AI 法のチェックリスト
SOC 2 証拠テンプレート
Agentmetry はコミュニティによって構築されています。未解決の問題またはロードマップを参照します。
イベントごとの MITRE タグは、単一のツール呼び出しが何であるかを示します。検出エンジンは、一連の呼び出しが何を意味するのか、つまりエージェントのセッション境界がなかったために EDR が認識できない信号を示します。
イベントが到着するとルールが実行されます。起動ルールは、ファーストクラスの正規イベント ( action.type: detect 、 action.outcome: <severity> ) としてセッションごとに 1 回発行され、同じシンクダウンします。

[切り捨てられた]

## Original Extract

Agentmetry — a local-first SIEM and flight recorder for AI coding agents. Captures every tool call, approval, and denial from Cursor, Claude Code, Codex and Antigravity into a MITRE ATT&CK-mapped audit trail, with correlated sequence detection (credential access -> network egress) and forwarders to
[truncated]

GitHub - blitzcrieg1/agentmetry: Agentmetry — a local-first SIEM and flight recorder for AI coding agents. Captures every tool call, approval, and denial from Cursor, Claude Code, Codex and Antigravity into a MITRE ATT&CK-mapped audit trail, with correlated sequence detection (credential access -> network egress) and forwarders to Loki, Elastic and Splunk. Apache-2.0. Zero egress. · GitHub
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
blitzcrieg1
/
agentmetry
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
119 Commits 119 Commits .agents .agents .claude .claude .codex .codex .github/ workflows .github/ workflows adapters adapters apps apps docs docs infra/ loki infra/ loki policies/ dlp policies/ dlp scripts scripts tests tests vault vault .cursorignore .cursorignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md docker-compose.loki.yml docker-compose.loki.yml docker-compose.yml docker-compose.yml View all files Repository files navigation
Agentmetry: SIEM for AI Agents
The open-source flight recorder and security layer for AI agent tool-use.
Every tool call, every denial, every human approval — hashed, correlated, and stored in a JSONL trail you own.
Replay on demand; forward to Loki, Elastic, or Splunk when you want a SIEM.
Quickstart ·
Docs ·
Schema ·
Roadmap ·
Security
Every AI-agent tool call, tagged with MITRE ATT&CK as it happens — and a
correlated alert when a sequence of innocent-looking calls adds up to an attack.
🚧 Public Alpha : Core capture, replay, and SIEM forwarding are usable for early exploration. APIs and integration surfaces may evolve rapidly.
When an autonomous agent runs a tool, most stacks keep nothing you could hand to an incident responder. Logs show a process; they do not show intent , session boundaries , or what the human approved .
Agentmetry is the open-source endpoint flight recorder for AI agents — built to run entirely on your machine, with optional forwarding to the SIEM you already operate.
an immutable, operator-owned audit trail for governed AI agents — capturing tool execution at the IDE lifecycle boundary and the MCP wire, not in a vendor cloud
Intercepting agent tool calls through IDE lifecycle hooks (Cursor, Claude Code, Codex, Antigravity) and an MCP stdio audit proxy
Normalizing every event into a canonical schema v1.1.0 with MITRE ATT&CK enrichment and SHA-256 argument hashing
Detecting correlated behavioral sequences a single event cannot reveal (credential exfil, guardrail bypass, recon-then-grab)
Blocking secrets and PII at the hook boundary with a local regex DLP engine ( log or block mode)
Forwarding the same JSONL trail to Loki, Elastic ECS, Splunk HEC, or a generic webhook — without making the cloud the system of record
Agentmetry is not a CASB or shadow-AI spy. It records the agents you wire in. If your problem is unmanaged ChatGPT in the browser, you need network/endpoint policy — not a flight recorder.
Agentmetry runs fully locally. The audit trail never leaves your machine unless you explicitly forward it.
See it catch something first (30 seconds)
No server, no API key, no config. Clone and run:
git clone https://github.com/blitzcrieg1/agentmetry.git && cd agentmetry
pip install -r apps/orchestrator/requirements.txt
python scripts/demo.py
It replays an agent session through the real ingest API: the agent reads an SSH
private key, runs a command containing an AWS key, then fetches a URL. Agentmetry
tags each call with MITRE ATT&CK, catches the AWS key with DLP (storing the rule,
never the value), and then — without being asked — correlates the key read with
the network call and fires a CRITICAL credential-exfil detection.
No single one of those events is an alert. The sequence is. That is the whole
product in one screen.
See the dashboard with a story in it
python scripts/demo_dashboard.py # seeds 5 sessions + 4 detections, serves http://127.0.0.1:8010/
One command seeds a realistic demo trail and serves the dashboard locally — no
API key, no cloud. Click a flagged session and the detection is right there, with
the full event drilled open beneath it:
See the dashboard tour for what each view shows and how
to read it.
Requirement
Version
Python
3.10+
Node.js
18+
1. Clone and install
git clone https: // github.com / blitzcrieg1 / agentmetry.git
cd agentmetry
# Python orchestrator
cd apps\orchestrator
python - m venv .venv
.\.venv\Scripts\activate
pip install - e " .[dev] "
copy .env.example .env
cd ..\..
# Next.js dashboard
cd apps\dashboard
npm install
cd ..\..
2. Boot the flight recorder
scripts\start-dev.bat
Dashboard → http://localhost:3000 · Orchestrator API → http://localhost:8000
powershell - ExecutionPolicy Bypass - File scripts\install_cursor_hooks.ps1
powershell - ExecutionPolicy Bypass - File scripts\install_claude_hooks.ps1
Fully quit and restart Cursor / Claude Code so hooks load.
python scripts\agentmetry_ingest.py selftest
Events should appear in the dashboard Flight Recorder within a few seconds.
When an agent runs a tool, Agentmetry automatically:
Intercepts the lifecycle hook or MCP tools/call before arguments leave the hook process
Hashes tool arguments (SHA-256) and scrubs inline secrets in command strings
Enriches each event with MITRE tactic/technique mappings and session correlation
Stores canonical JSONL locally ( audit-forward.jsonl ) — the system of record for the hook path
Detects multi-step behavioral patterns across the session timeline
Forwards to your SIEM sinks and alert webhook (optional, best-effort)
flowchart TB
subgraph Capture["Capture Layer (Tier A + B)"]
HOOKS["IDE Lifecycle Hooks<br/>Cursor · Claude · Codex · Antigravity"]
PROXY["MCP Audit Proxy<br/>mcp_audit_proxy.py"]
end
subgraph Gate["Local Security Gate"]
DLP["DLP Scanner<br/>regex rules"]
HASH["Arg Hash + Secret Scrub"]
end
subgraph Core["Orchestrator :8000"]
INGEST["POST /api/v1/audit/ingest"]
CANON["Canonical Schema v1.1.0<br/>MITRE enrichment"]
DETECT["Sequence Detection Engine"]
OUTBOX[("SQLite Outbox<br/>events.db")]
end
subgraph Output["Outputs"]
JSONL["audit-forward.jsonl"]
DASH["Dashboard<br/>Flight Recorder + Analytics"]
SIEM["Loki · Elastic · Splunk · Webhook"]
end
HOOKS --> DLP
PROXY --> DLP
DLP -->|allow| HASH
DLP -->|deny| INGEST
HASH --> INGEST
INGEST --> CANON
CANON --> OUTBOX
CANON --> JSONL
CANON --> DETECT
JSONL --> DASH
JSONL --> SIEM
Loading
Capture paths
flowchart LR
subgraph TierB["Tier B — IDE Hooks"]
C["Cursor"]
CL["Claude Code"]
AG["Antigravity"]
CX["Codex"]
end
subgraph TierA["Tier A — MCP Proxy"]
MCP["Any MCP Client"]
WRAP["Audit Proxy wraps server command"]
end
INGEST["agentmetry_ingest.py → /audit/ingest"]
C --> INGEST
CL --> INGEST
AG --> INGEST
CX --> INGEST
MCP --> WRAP --> INGEST
Loading
Component
Path
Role
Hook client
scripts/agentmetry_ingest.py
Maps IDE lifecycle events to canonical payloads; hashes args in-process
MCP proxy
apps/orchestrator/tools/mcp_audit_proxy.py
Wraps any stdio MCP server; logs every tools/call + errors
Ingest API
core/audit/ingest.py
Normalizes payloads, infers approvals ( inferred:* ), writes sinks
DLP engine
core/audit/dlp/
Regex scan of tool arguments (validators, e.g. Luhn); block or log before execution
Detection engine
core/audit/detection/
Correlated sequence rules over a session's event timeline
Sinks
core/audit/sinks.py
File, webhook, Elastic ECS, Splunk HEC
Replay
core/audit/replay.py
ASCII timeline reconstruction from the local outbox
The canonical event
Every run emits typed, SIEM-ready JSON. A single tool_called line:
{
"schema_version" : " 1.1.0 " ,
"correlation_id" : " thread-8892 " ,
"timestamp_utc" : " 2026-07-12T09:14:22.041+00:00 " ,
"actor" : { "type" : " user " , "id" : " dev_01 " , "role" : " operator " },
"action" : { "type" : " tool_called " , "outcome" : " success " },
"agent" : { "name" : " cursor " , "skill_id" : " " },
"tool" : {
"qualified" : " vault_fs.read_file " ,
"server" : " vault_fs " ,
"input_hash" : " e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 " ,
"parameters_redacted" : true ,
"mitre" : { "tactic" : " Collection (TA0009) " , "technique" : " Data from Local System (T1005) " }
},
"model" : { "id" : " claude-3-5-sonnet " , "provider" : " anthropic " }
}
Full schema → docs/agentmetry-event-schema.md
Agentmetry records agents you wire in — IDE hooks or the MCP proxy . It is honest about what it cannot see.
🎥 Flight Recorder
Live audit tail with dynamic columns, drag-and-drop layout, CSV export, and session drill-down
📊 Analytics & Process Tree
Session-level charts, MITRE tactic breakdown, horizontal React Flow timeline
🔍 Behavioral Detection
Correlated sequence rules — credential exfil, guardrail bypass, recon-then-grab
🛡️ Local DLP
Regex scanner blocks AWS keys, GitHub tokens, Slack tokens, and PII before tool execution
🎯 MITRE ATT&CK mapping
Per-tool tactic/technique tags on every canonical event
🔐 Argument hashing
SHA-256 of tool args by default — plaintext never crosses the wire from hooks
📡 SIEM-native export
Elastic ECS, Splunk HEC, Loki/LogQL, generic webhook, alert webhook on denials
🔁 Replay & evidence
ASCII session timeline + tamper-evident evidence pack export
👥 Multi-IDE support
Cursor, Claude Code, Codex, Antigravity — global hook install scripts
Integrations
Category
Supported today
Roadmap
IDE / Agent hosts
Cursor · Claude Code · Codex · Antigravity
Windsurf · VS Code Copilot
MCP transport
Stdio audit proxy (wrap any MCP server command)
SSE / streamable HTTP proxy
Observability / SIEM
Loki · Grafana · Elastic ECS · Splunk HEC · generic webhook
Datadog · New Relic
Detection formats
In-engine sequence rules · LogQL · Elastic · Splunk · Sigma pack
STIX/TAXII export
Policy engines
Regex DLP manifest ( policies/dlp/ )
OPA / Rego policy-as-code
Compliance docs
ISO 42001 mapping · AI Act checklist
SOC 2 evidence templates
Agentmetry is community-built. Browse open issues or the roadmap .
Per-event MITRE tags say what a single tool call is. The detection engine says what a sequence of calls means — the signal an EDR cannot see because it never had the agent's session boundary.
Rules run as events arrive . A firing rule is emitted once per session as a first-class canonical event ( action.type: detection , action.outcome: <severity> ) down the same sinks a

[truncated]
