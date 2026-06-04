---
source: "https://github.com/sharb1235-hash/attow-nexus"
hn_url: "https://news.ycombinator.com/item?id=48397993"
title: "Attow Nexus – Git for AI agent state"
article_title: "GitHub - sharb1235-hash/attow-nexus: A local coordination daemon and Git-like state ledger for polyglot AI agents. · GitHub"
author: "Suleiman_Harb"
captured_at: "2026-06-04T13:13:25Z"
capture_tool: "hn-digest"
hn_id: 48397993
score: 1
comments: 0
posted_at: "2026-06-04T12:55:12Z"
tags:
  - hacker-news
  - translated
---

# Attow Nexus – Git for AI agent state

- HN: [48397993](https://news.ycombinator.com/item?id=48397993)
- Source: [github.com](https://github.com/sharb1235-hash/attow-nexus)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T12:55:12Z

## Translation

タイトル: Attow Nexus – AI エージェント状態のための Git
記事のタイトル: GitHub - sharb1235-hash/attow-nexus: 多言語 AI エージェント用のローカル調整デーモンおよび Git のような状態台帳。 · GitHub
説明: 多言語 AI エージェント用のローカル調整デーモンおよび Git のような状態台帳。 - sharb1235-ハッシュ/attow-nexus

記事本文:
GitHub - sharb1235-hash/attow-nexus: 多言語 AI エージェント用のローカル調整デーモンおよび Git のような状態台帳。 · GitHub
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
sharb1235-ハッシュ
/
アトウネクサス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のn

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .github .github cli cli デーモン デーモン ダッシュボード ダッシュボード ドキュメント ドキュメントの例 例 proto/ nexus/ v1 proto/ nexus/ v1 スクリプト スクリプト sdks sdks テストフィクスチャ テストフィクスチャ .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile ライセンスLICENSE Makefile Makefile NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md buf.gen.yaml buf.gen.yaml buf.yaml buf.yaml docker-compose.yml docker-compose.yml justfile justfile すべてのファイルを表示リポジトリ ファイルのナビゲーション
多言語 AI エージェント用のローカル調整デーモンと Git のような状態台帳。
ローカルファーストの AI インフラストラクチャ スタックの一部として Attow, Inc. によって構築されました。
1 つのローカル デーモン、1 つの台帳、1 つの CLI、同じ実行をフィードする 3 つのフレームワーク サーフェス。
別のエージェント フレームワークではなく、その下にある共有状態とデバッグ基板です。
Nexus-IPC は、多言語 AI エージェント用のローカル調整デーモンです。 NexusLedger は、その上に構築された Git のような状態履歴です。 Attow Nexus Console は、ライブ デバッガーおよびダッシュボードです。これらを組み合わせることで、エージェントはリアルタイムで状態を共有できるようになり、開発者は重要な状態遷移ごとにリプレイ、ロールバック、差分分析、ループ検出、デバッグを行うことができます。
Attow Nexus ローカル デーモン、ユニバーサル デモ、ダッシュボードを 2 つの端末で実行します。
ターミナル 1 - これを実行し続けます。
cd <リポジトリ>
docker 構成 -- ビルド
Windows PowerShell 上のターミナル 2:
cd <リポジトリ>
.\scripts\setup.ps1
PowerShell が実行される場合

イオン ポリシーがローカル スクリプトをブロックする場合は、次を実行します。
powershell - NoProfile - ExecutionPolicy Bypass - ファイル .\scripts\setup.ps1
macOS/Linux 上のターミナル 2:
cd <リポジトリ>
chmod +x スクリプト/setup.sh
./scripts/setup.sh
ブラウザ:
セットアップ スクリプトによって出力された Vite Local URL (通常は http://127.0.0.1:5173 または http://127.0.0.1:5174 ) を開きます。
セットアップ スクリプトは API とメトリクスをチェックし、編集可能モードでローカル Python SDK をインストールし、キーなしのユニバーサル翻訳デモを実行し、TypeScript/Vercel スタイルのデモを実行し、CLI インスペクション出力を出力し、Vite ダッシュボード サーバーを起動します。クラウド アカウントや API キーは必要ありません。
Docker Compose は、ターミナル 1 でデーモン/API/メトリクスを実行します。ダッシュボードは、 Vite dev モードを通じて、 dump/ から個別に実行されます。
Attow Nexus は、v0.1 ではデーモン、CLI、SDK、またはローカル ダッシュボード ランタイム テレメトリを収集しません。ローカル ランタイムのプライバシー境界については、PRIVACY.md を参照してください。
基板をネイティブに検査する
これらのコマンドはクイックスタートには必要ありません。これらは、ローカル デーモン、レジャー、および CLI を直接検査する場合に便利です。
カール.exe http://127.0 。 0.1 : 7822 / API / ヘルス
カール.exe http://127.0 。 0.1 : 7823 / メトリクス
カーゴラン - p nexus -- ステータス
カーゴラン - p nexus -- エージェント
カーゴ ラン - P ネクサス -- チャネル
カーゴ ラン - p nexus -- ログ -- ユニバーサル実行 - デモ
カーゴ ラン - p nexus -- diff < commit_a > < commit_b >
カーゴ ラン - p nexus -- 再生 < commit_b >
macOS/Linux:
カール http://127.0.0.1:7822/api/health
カール http://127.0.0.1:7823/metrics
カーゴ ラン -p nexus -- ステータス
カーゴ ラン -p nexus -- エージェント
カーゴ ラン -p nexus -- チャネル
カーゴ run -p nexus -- log --run universal-demo
カーゴ ラン -p nexus -- diff < commit_a > < commit_b >
カーゴ ラン -p nexus -- 再生 < commit_b >
壊れたエージェントの回復デモ
ステップ 14 でエージェントがクラッシュした場合は、ステップ 1 ～ 13 の再実行を停止します。この奴ら

o は、不正な状態遷移を生成するローカル マルチエージェント ワークフロー、Nexus が不正なコミットを識別、開発者が最後の良好な状態を比較/再生し、修正されたコミットから回復するワークフローを示しています。
Git により、開発者はコードのバージョン管理を行うことができました。 Attow Nexus は、開発者にエージェントの状態のバージョン管理を提供します。
docker 構成 -- ビルド
ターミナル 2:
py - m pip install - e sdks\python
py 例\壊れた - エージェント - 回復\run_demo.py
次に、スクリプトによって出力されたコミット ID を検査します。
カーゴの実行 - P ネクサス -- ログ -- 壊れた実行 - エージェント - デモ
カーゴ ラン - p nexus -- diff < last_good_commit > < bad_commit >
カーゴ ラン - p ネクサス -- リプレイ <final_commit>
デモは決定的かつローカルです。LLM 呼び出し、API キー、クラウド サービスはありません。 Nexus はキャプチャされた論理状態を再生します。アダプターが補償アクションを提供しない限り、ファイルの書き込み、電子メール、API 呼び出し、購入、デプロイメント、データベースの変更などの実際の副作用を自動的に元に戻すことはありません。
Docker Compose は、Attow Nexus デーモン、HTTP API、およびメトリクス エンドポイントを起動します。現在、 / にある本番環境の React ダッシュボードは提供されていません。
API ヘルス: http://127.0.0.1:7822/api/health
メトリクス: http://127.0.0.1:7823/metrics
ダッシュボード開発サーバー: Dashboard/ から Vite によって出力される URL、通常は http://127.0.0.1:5173 または http://127.0.0.1:5174
CLI ステータス、エージェント、チャネル、ログ、差分、および再生。
Python/TypeScript フィクスチャのパリティ。
/api/events の正規の取り込み。
LLM を使用しない実際の例を含む LangGraph ラッパー。
CrewAI スタイルのラッパーは、偽/パブリック コールバック形状で検証されました。
Vercel AI SDK スタイルのラッパーは、偽のgenerateText / streamText で検証されました。
Docker での運用ダッシュボードのバンドル。
フレームワーク アダプターは初期のものであり、パブリック フック、コールバック、ミドルウェア、チェックポイント プログラム、または明示的なラッパーを使用する必要があります。
深いfr

amework ネイティブのチェックポイント作成者/ストア ブリッジ。
本番環境のリモート同期。
セキュリティ強化はローカルファースト/開発ファーストです。リモート使用には展開固有のレビューが必要です。
クラウドとチームの機能はロードマップのみです。
フローチャート LR
A["Python エージェント"] --> IPC["Nexus-IPC デーモン"]
B["TypeScript エージェント"] --> IPC
C["Rust/Go/カスタム エージェント"] --> IPC
IPC --> BUS["メモリ内状態バス"]
BUS --> SUB["加入者"]
IPC --> LEDGER["NexusLedger 追加専用 DAG"]
LEDGER --> SQLITE["SQLite WAL ストア"]
LEDGER --> ART[「アーティファクトストア」]
IPC --> CONSOLE["Attow Nexus コンソール"]
IPC --> MCP["オプションの MCP ブリッジ"]
読み込み中
3層
Nexus-IPC は、ライブ ローカル調整バスです。エージェントは、フレームワーク コードを結合することなく、登録、ハートビート、構造化状態デルタの公開、チャネルへのサブスクライブ、ブロードキャストの受信、キャプチャされたコンテキストの共有を行います。
NexusLedger は、耐久性のある Git のような状態グラフです。永続的なデルタは、差分、再生、フォーク、エクスポート、検査できるコンテンツアドレス指定のコミットになります。
Attow Nexus Console はローカル ダッシュボードです。接続されているエージェント、アクティブなチャネル、最近の差分、コミット履歴、ツール呼び出し、ループ警告、再生出力、フォーク コントロール、ロールバック警告、およびメトリックが表示されます。
エージェント フレームワークは、1 つのアプリケーション内でのオーケストレーションに優れています。実際のシステムには、複数のエージェント、言語、ランタイム、ローカル プロセス、ツールが含まれることがよくあります。 Attow Nexus は、エージェント フレームワークの下に共有ローカル基盤を提供するため、バージョン管理されたプロトコル メッセージと永続的な論理エージェント状態を通じて調整できます。
なぜ LangGraph メモリだけではだめなのでしょうか?
LangGraph のメモリとチェックポイントは、LangGraph アプリケーションにとって優れています。 Attow Nexus は、フレームワークネイティブのメモリ、永続性、またはオーケストレーションを置き換えようとはしていません。
Attow Nexus は、論理エージェントの状態が複数のフレームワーク、言語、

ランタイム、ローカル プロセス、ツール、またはカスタム ワーカー。これは、エージェント フレームワークの下でニュートラルなローカル ステート バスと Git のような台帳として機能するため、LangGraph、CrewAI、AutoGen、Microsoft Agent Framework、カスタム Python ワーカー、TypeScript サービス、およびその他のプロセスは、すべて同じアプリケーション フレームワークを採用することなく調整できます。
意図された関係は補完的です。各アプリ内で最適に動作するフレームワーク ネイティブ メモリを使用し続け、構造化状態のデルタの共有、永続的なコミット、プロセス間の可観測性、キャプチャされた論理状態の差分、再生、ロールバックがフレームワークの境界を越える必要がある場合は、Attow Nexus を使用します。
ユニバーサル翻訳層のデモ
v0.1 開発者プレビューの主張は単純です。1 つのローカル デーモン、1 つの台帳、1 つの CLI、同じ実行をフィードする 3 つのフレームワーク サーフェスです。
Attow Nexus は、複数のフレームワークからの状態イベントを 1 つのローカル バスと台帳に正規化します。現在のアダプターはパブリック サーフェス ラッパーです。これらはノードの汚染や状態スキーマの変更を必要とせず、ディープ フレームワーク ネイティブ ストアであるふりをしません。深いフレームワーク固有のチェックポイント作成者/ストア ブリッジは将来の作業です。
共有イベント コントラクトは docs/protocol.md に文書化されており、HTTP JSON サーフェスは docs/http-api.md に文書化されています。 Python アダプターと TypeScript アダプターは、 test-fixtures/universal-events の同じ正規フィクスチャに対してテストされます。
Examples/universal-translation-demo でローカルのキーなしデモを実行します。
cd <リポジトリ>
docker 構成 -- ビルド
別の PowerShell では次のようになります。
cd <リポジトリ>
cd sdks\Python
py -m pip install -e 。
py - m pip インストール langgraph
CD ..\..
py 例\ユニバーサル - 翻訳 - デモ\langgraph_planner.py
py 例\ユニバーサル - 翻訳 - デモ\crewai_researcher.py
3 番目の PowerShell では次のようになります。
cd < リポジトリ > \examples\universal - 翻訳 - デモ
npm。

cmd インストール
npm.cmd で vercel を実行 - デモ
次に、共有実行を検査します。
cd <リポジトリ>
カーゴラン - p nexus -- エージェント
カーゴ ラン - P ネクサス -- チャネル
カーゴ ラン - p nexus -- ログ -- ユニバーサル実行 - デモ
フレームワークの表面
言語
アダプターAPI
検証レベル
注意事項
ランググラフ
パイソン
楽器_言語グラフ
偽のグラフテストを使用した検証済みラッパー。 langgraph がインストールされている場合の実際の LLM なしの例
ノードの変更はありません
CrewAI
パイソン
楽器クルーワイ
偽のクルーテストを使用した検証済みラッパー。公開 step_callback コンポジション (利用可能な場合)
タスクの変更はありません
Vercel AI SDK
TypeScript
楽器ストリームテキスト / 楽器生成テキスト
偽の streamText /generateText で検証
コールバックを保持します
4 行の LangGraph
nexus_ipc から NexusClient をインポート
nexus_ipc から。アダプター。ランググラフのインポート、instrument_langgraph
クライアント = NexusClient 。接続する()
グラフ =instrument_langgraph (グラフ、クライアント = クライアント、run_id = "my-run")
結果 = グラフ。 invoke (input_state , config = { "configurable" : { "thread_id" : "thread-1" }})
ノード関数は Nexus をインポートせず、LangGraph 状態スキーマは変更されません。 Attow Nexus が場所を記録している間、ラップされたグラフは通常の .invoke() 、 .stream() 、 .ainvoke() 、および .astream() の動作を維持します。

[切り捨てられた]

## Original Extract

A local coordination daemon and Git-like state ledger for polyglot AI agents. - sharb1235-hash/attow-nexus

GitHub - sharb1235-hash/attow-nexus: A local coordination daemon and Git-like state ledger for polyglot AI agents. · GitHub
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
sharb1235-hash
/
attow-nexus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github .github cli cli daemon daemon dashboard dashboard docs docs examples examples proto/ nexus/ v1 proto/ nexus/ v1 scripts scripts sdks sdks test-fixtures test-fixtures .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md buf.gen.yaml buf.gen.yaml buf.yaml buf.yaml docker-compose.yml docker-compose.yml justfile justfile View all files Repository files navigation
A local coordination daemon and Git-like state ledger for polyglot AI agents.
Built by Attow, Inc. as part of a local-first AI infrastructure stack.
One local daemon, one ledger, one CLI, three framework surfaces feeding the same run.
Not another agent framework - the shared state and debugging substrate underneath them.
Nexus-IPC is a local coordination daemon for polyglot AI agents. NexusLedger is the Git-like state history built on top of it. Attow Nexus Console is the live debugger and dashboard. Together, they let agents share state in real time while giving developers replay, rollback, diffing, loop detection, and debugging for every important state transition.
Get the Attow Nexus local daemon, universal demo, and dashboard running with two terminals.
Terminal 1 - keep this running:
cd < repo >
docker compose up -- build
Terminal 2 on Windows PowerShell:
cd < repo >
.\scripts\setup.ps1
If your PowerShell execution policy blocks local scripts, run:
powershell - NoProfile - ExecutionPolicy Bypass - File .\scripts\setup.ps1
Terminal 2 on macOS/Linux:
cd < repo >
chmod +x scripts/setup.sh
./scripts/setup.sh
Browser:
Open the Vite Local URL printed by the setup script, usually http://127.0.0.1:5173 or http://127.0.0.1:5174 .
The setup script checks the API and metrics, installs the local Python SDK in editable mode, runs the no-key universal translation demo, runs the TypeScript/Vercel-style demo, prints CLI inspection output, and starts the Vite dashboard server. No cloud account or API key is required.
Docker Compose runs the daemon/API/metrics in Terminal 1. The dashboard runs separately through Vite dev mode from dashboard/ .
Attow Nexus does not collect daemon, CLI, SDK, or local dashboard runtime telemetry in v0.1. See PRIVACY.md for the local-runtime privacy boundary.
Inspecting the Substrate Natively
These commands are not required for the quickstart. They are useful when you want to inspect the local daemon, ledger, and CLI directly.
curl.exe http: // 127.0 . 0.1 : 7822 / api / health
curl.exe http: // 127.0 . 0.1 : 7823 / metrics
cargo run - p nexus -- status
cargo run - p nexus -- agents
cargo run - p nexus -- channels
cargo run - p nexus -- log -- run universal - demo
cargo run - p nexus -- diff < commit_a > < commit_b >
cargo run - p nexus -- replay < commit_b >
macOS/Linux:
curl http://127.0.0.1:7822/api/health
curl http://127.0.0.1:7823/metrics
cargo run -p nexus -- status
cargo run -p nexus -- agents
cargo run -p nexus -- channels
cargo run -p nexus -- log --run universal-demo
cargo run -p nexus -- diff < commit_a > < commit_b >
cargo run -p nexus -- replay < commit_b >
Broken Agent Recovery Demo
When an agent crashes at step 14, stop rerunning steps 1-13. This demo shows a local multi-agent workflow producing a bad state transition, Nexus identifying the bad commit, the developer diffing/replaying the last good state, and the workflow recovering from a fixed commit.
Git gave developers version control for code. Attow Nexus gives developers version control for agent state.
docker compose up -- build
Terminal 2:
py - m pip install - e sdks\python
py examples\broken - agent - recovery\run_demo.py
Then inspect with the commit IDs printed by the script:
cargo run - p nexus -- log -- run broken - agent - demo
cargo run - p nexus -- diff < last_good_commit > < bad_commit >
cargo run - p nexus -- replay < final_commit >
The demo is deterministic and local: no LLM calls, no API keys, and no cloud services. Nexus replays captured logical state. It does not automatically undo real-world side effects such as file writes, emails, API calls, purchases, deployments, or database mutations unless an adapter provides compensating actions.
Docker Compose starts the Attow Nexus daemon, HTTP API, and metrics endpoint. It does not currently serve the production React dashboard at / .
API health: http://127.0.0.1:7822/api/health
Metrics: http://127.0.0.1:7823/metrics
Dashboard dev server: the URL printed by Vite from dashboard/ , usually http://127.0.0.1:5173 or http://127.0.0.1:5174
CLI status , agents , channels , log , diff , and replay .
Python/TypeScript fixture parity.
/api/events canonical ingestion.
LangGraph wrapper with a real no-LLM example.
CrewAI-style wrapper verified with fake/public callback shape.
Vercel AI SDK-style wrapper verified with fake generateText / streamText .
Production dashboard bundling in Docker.
Framework adapters are early and should use public hooks, callbacks, middleware, checkpointers, or explicit wrappers.
Deep framework-native checkpointer/store bridges.
Production remote synchronization.
Security hardening is local-first/dev-first; remote use needs a deployment-specific review.
Cloud and team features are roadmap only.
flowchart LR
A["Python agent"] --> IPC["Nexus-IPC daemon"]
B["TypeScript agent"] --> IPC
C["Rust/Go/custom agent"] --> IPC
IPC --> BUS["In-memory state bus"]
BUS --> SUB["Subscribers"]
IPC --> LEDGER["NexusLedger append-only DAG"]
LEDGER --> SQLITE["SQLite WAL store"]
LEDGER --> ART["Artifact store"]
IPC --> CONSOLE["Attow Nexus Console"]
IPC --> MCP["Optional MCP bridge"]
Loading
Three Layers
Nexus-IPC is the live local coordination bus. Agents register, heartbeat, publish structured state deltas, subscribe to channels, receive broadcasts, and share captured context without coupling their framework code.
NexusLedger is the durable Git-like state graph. Durable deltas become content-addressed commits that can be diffed, replayed, forked, exported, and inspected.
Attow Nexus Console is the local dashboard. It shows connected agents, active channels, recent deltas, commit history, tool calls, loop warnings, replay output, fork controls, rollback warnings, and metrics.
Agent frameworks are good at orchestration inside one application. Real systems often have multiple agents, languages, runtimes, local processes, and tools. Attow Nexus provides the shared local substrate beneath agent frameworks so they can coordinate through versioned protocol messages and durable logical agent state.
Why Not Just LangGraph Memory?
LangGraph memory and checkpointing are excellent for LangGraph applications. Attow Nexus is not trying to replace framework-native memory, persistence, or orchestration.
Attow Nexus is useful when logical agent state spans multiple frameworks, languages, runtimes, local processes, tools, or custom workers. It acts as a neutral local state bus plus Git-like ledger underneath agent frameworks, so LangGraph, CrewAI, AutoGen, Microsoft Agent Framework, custom Python workers, TypeScript services, and other processes can coordinate without all adopting the same application framework.
The intended relationship is complementary: keep using the framework-native memory that works best inside each app, and use Attow Nexus when shared structured state deltas, durable commits, cross-process observability, diffing, replay, and rollback of captured logical state need to cross framework boundaries.
Universal Translation Layer Demo
The v0.1 developer-preview claim is simple: one local daemon, one ledger, one CLI, three framework surfaces feeding the same run.
Attow Nexus normalizes state events from multiple frameworks into one local bus and ledger. The current adapters are public-surface wrappers. They do not require node pollution or state schema changes, and they do not pretend to be deep framework-native stores. Deep framework-specific checkpointer/store bridges are future work.
The shared event contract is documented in docs/protocol.md , and the HTTP JSON surface is documented in docs/http-api.md . Python and TypeScript adapters are tested against the same canonical fixtures in test-fixtures/universal-events .
Run the local no-key demo in examples/universal-translation-demo :
cd < repo >
docker compose up -- build
In another PowerShell:
cd < repo >
cd sdks\python
py - m pip install - e .
py - m pip install langgraph
cd ..\..
py examples\universal - translation - demo\langgraph_planner.py
py examples\universal - translation - demo\crewai_researcher.py
In a third PowerShell:
cd < repo > \examples\universal - translation - demo
npm.cmd install
npm.cmd run vercel - demo
Then inspect the shared run:
cd < repo >
cargo run - p nexus -- agents
cargo run - p nexus -- channels
cargo run - p nexus -- log -- run universal - demo
Framework surface
Language
Adapter API
Verified level
Notes
LangGraph
Python
instrument_langgraph
Verified wrapper with fake graph tests; real no-LLM example when langgraph is installed
No node changes
CrewAI
Python
instrument_crewai
Verified wrapper with fake crew tests; public step_callback composition when available
No task changes
Vercel AI SDK
TypeScript
instrumentStreamText / instrumentGenerateText
Verified with fake streamText / generateText
Preserves callbacks
LangGraph In Four Lines
from nexus_ipc import NexusClient
from nexus_ipc . adapters . langgraph import instrument_langgraph
client = NexusClient . connect ()
graph = instrument_langgraph ( graph , client = client , run_id = "my-run" )
result = graph . invoke ( input_state , config = { "configurable" : { "thread_id" : "thread-1" }})
Your node functions do not import Nexus, and your LangGraph state schema does not change. The wrapped graph keeps its normal .invoke() , .stream() , .ainvoke() , and .astream() behavior while Attow Nexus records loca

[truncated]
