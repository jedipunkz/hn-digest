---
source: "https://github.com/mario-guerra/orchid-trace"
hn_url: "https://news.ycombinator.com/item?id=48660800"
title: "Show HN: Orchid – Local-first record and replay for AI agent debugging"
article_title: "GitHub - mario-guerra/orchid-trace: Orchid - Orchestration interactive debugger - Record, inspect, & replay AI agents · GitHub"
author: "brightmonkey"
captured_at: "2026-06-24T14:55:55Z"
capture_tool: "hn-digest"
hn_id: 48660800
score: 1
comments: 0
posted_at: "2026-06-24T14:48:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Orchid – Local-first record and replay for AI agent debugging

- HN: [48660800](https://news.ycombinator.com/item?id=48660800)
- Source: [github.com](https://github.com/mario-guerra/orchid-trace)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T14:48:05Z

## Translation

タイトル: Show HN: Orchid – AI エージェントのデバッグのためのローカルファーストの記録と再生
記事のタイトル: GitHub - mario-guerra/orchid-trace: Orchid - オーケストレーション インタラクティブ デバッガー - AI エージェントの記録、検査、および再生 · GitHub
説明: Orchid - オーケストレーション対話型デバッガー - AI エージェントの記録、検査、および再生 - mario-guerra/orchid-trace
HN テキスト: Orchid (オーケストレーション インタラクティブ デバッガー) は、エージェント パイプライン内のすべての API および LLM 呼び出しをキャプチャし、実行全体をローカルで段階的に検査および再生できるゼロインストルメンテーション プロキシです。計測機器、ベンダー ロックイン、クラウドへの依存はありません。また、ビジュアル インスペクターと MCP サーバーも提供されるため、セッションを自分で検査したり、お気に入りのエージェント コーディング IDE を使用してエージェントの実行をデバッグしたりできます。私がこれを構築したのは、ログを grep してエージェントの障害をデバッグするのにうんざりしたためであり、利用可能な AI 可観測性ツールはすべて、侵入型の計測やプロンプトと応答をクラウド サービスに送信する必要があるように見えました。ベンダー ロックインやデータ プライバシーを心配することなく、ローカルで実行されるエージェントをデバッグできるものが必要でした。オーキッドはそのツールです。少なくとも私のユースケースでは、通話検査機能は非常にうまく機能しますが、おそらく再生機能の方が興味深いでしょう。これにより、高価な API 呼び出しをモックしたり再実行したりすることなく、LLM パイプライン テストが決定的になります。無料、自己ホスト型、マシンまたはインフラストラクチャ上で実行: https://github.com/mario-guerra/orchid-trace マルチステップのエージェント システムを構築している人、または非決定論的な LLM テストの失敗に悩んでいる人からのフィードバックをお待ちしています。

記事本文:
GitHub - mario-guerra/orchid-trace: Orchid - オーケストレーション インタラクティブ デバッガー - AI エージェントの記録、検査、および再生 · GitHub
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
マリオゲッラ
/
蘭の跡
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
46 コミット 46 コミット .github/ workflows .github/ workflows アセット アセット docs docs sdk sdk LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Orchid — AI エージェントの記録、検査、再生
ログの grep を停止します。 Orchid は、エージェントのネットワーク トラフィック (LLM 呼び出し、ツールの呼び出し、エージェントが通信するその他の API) を、ゼロ計測プロキシを通じて記録します。次に、次のことが可能になります。
完了した実行を段階的にタイムトラベルします
すべてのプロンプト、応答、トークン数、コストを検査する
組み込みの Web UI または IDE からの MCP ツール経由でエラーをデバッグします
記録された実行をオフラインで再生する - API コストゼロの確定的テスト
Orchid は、コーディング エージェントに AI アプリをデバッグする機能を提供します。プロキシには MCP サーバーが組み込まれているため、LLM 呼び出しがスタックの奥深く (フレームワーク、キュー、3 つの抽象化層の背後) に埋め込まれている場合、Cursor、VS Code、または Claude Code の AI アシスタントが記録されたトラフィックを直接クエリできます。印刷ステートメントやログの探索はありません。エージェントに「なぜこの実行が失敗したのですか?」と尋ねてください。そして、原因を調べて解明し、修正してくれるのです。
記録する量を選択します。軽量の検査のために LLM トラフィックのみをプロキシ経由でルーティングするか、全体像を把握するためにすべてをキャプチャします。実行を完全な忠実度で再生するには、エージェントのすべてのネットワーク トラフィックがプロキシを経由する必要があります。再生は記録された応答を返すことによって機能するため、記録されていないものは再生できません。
重要な注意 - データがインフラストラクチャから流出することはありません。
プロキシは、アプリがすでに呼び出していたアップストリーム API にのみリクエストを転送します。
記録されたものはすべて、コンテナー内のローカル SQLite データベース (またはマウントされたボリューム) に残ります。

自分）。フォンホーム、テレメトリ、クラウド バックエンドはありません。
シークレットはディスクに何かが書き込まれる前にメモリ内でスクラブされます。認可ヘッダーは変更されずに上流に転送されますが、保存されることはありません。また、シークレットのような名前 (キー、トークン、パスワード、認証情報、Cookie) を持つヘッダー、クエリ文字列、および本文フィールドは [編集済み] として保存されます。
正直な注意点が 1 つあります。編集は、プロンプトの内容をスキャンすることではなく、フィールド名 ( api_key や authorization など) を認識することによって機能します。プロンプトと完了テキストはそのまま記録されます。これが Orchid の重要な点です。そのため、シークレットがプロンプトに貼り付けられた場合、それはプロンプト テキストの残りの部分とともに保存されます。
このリポジトリには、オープンソースの Orchid SDK とユーザー ドキュメントが含まれています。 orchid-proxy コンテナは、GitHub Container Registry 経由で配布されます (以下を参照)。ここのコンテンツはメインの開発リポジトリから自動的に同期されます。問題やディスカッションは歓迎です。プル リクエストは直接マージされるのではなく、移植される場合があります。
独自のアプリをインストルメント化せずに Orchid が動作している様子を見てみたいですか? Orchid LangGraph のデモをチェックしてください。
このリポジトリは、Orchid にネイティブに接続された複雑なマルチエージェント AI アーキテクチャ (OpenAI、Anthropic、Google Vertex AI を使用) を紹介します。これには完全なトレース フィクスチャが含まれており、設定なしでオフライン再生モードで敵対的エージェント パイプライン全体を実行できます。API キーは必要なく、アップストリーム コストもかかりません。これは、決定論的な AI テストを体験し、ビジュアライザー ダッシュボードをテストするための絶対に最速の方法です。
フローチャート LR
subgraph env["あなたの環境 (ラップトップ、オンプレミス、または独自のクラウド)"]
app["アプリケーション<br/>(Python、TS、...)"]
プロキシ["オーキッドプロキシ<br/>:4320 (プロキシ)<br/>:4321 (クエリ/UI/MCP)"]
db[("orchid.db<br/>(SQLite)")]
アプリ -- "HTTP" --> pr

オキシ
プロキシ -- 「レコード」 --> データベース
データベース - 。 "リプレイ" .-> プロキシ
終わり
上流["アップストリーム API<br/>OpenAI、Anthropic、<br/>Gemini、ツール、あらゆる API ..."]
プロキシ -- 「HTTPS<br/>(リプレイ モードではスキップ)」 --> アップストリーム
読み込み中
非侵入的傍受 (Thin SDK)
すべてのクライアント初期化をラップしたり、AST を変更した重い SDK を使用したりする必要がある従来の LLM 可観測性ツールとは異なり、Orchid は APM スタイルの Thin SDK を使用します。 SDK は、基本的な HTTP トランスポート層 (Python では httpx / リクエスト、Node ではフェッチ) にパッチを適用するため、標準クライアント ライブラリによって行われるすべての LLM 呼び出しは、プロンプト処理や生成コードを変更することなく、ローカルまたはリモートの Orchid プロキシを通じて自動的にルーティングされます。
プロキシはアプリケーションの状態を追跡しません。 X-Orchid-* HTTP ヘッダー (SDK によって挿入されるか、任意の言語から手動で設定される) を読み取り、各リクエストの処理方法を決定します。
passthrough : 透過的なリバース プロキシ。リクエストを転送し、ディスクに何も書き込まずに応答を返します。
Capture : リクエストを転送し、完全なリクエスト/レスポンス ペイロード (ストリーミング チャンクを含む) をシリアル化し、コストを計算して、特定の Session ID でローカル SQLite データベースに保存します。
replay : すべての送信ネットワーク トラフィックをブロックします。受信リクエストをハッシュし、正確に一致する記録された応答を SQLite から提供します。一致するものが見つからない場合は、決定的な模擬エラーを返します。
キャプチャされたすべての LLM コール (または「Exchange」) は以下を記録します。
メタデータの要求 : システム プロンプト、ユーザー プロンプト、温度、top-p、およびカスタム タグ。
応答テレメトリ: 完全な完了テキスト、使用トークン (入力/出力)、および遅延。
コスト計算 : ユーザー指定のモデル価格マップに基づくリアルタイムの USD コスト帰属 ( docs/configuration.md および docs/api_reference.md の /pricing エンドポイントを参照)。
ストリーム・レアス

sembly : ストリーミング補完の場合、Orchid は SSE チャンクをメモリにバッファリングしてクライアントに即座に提供し、完全に再構築された補完本体を SQLite に書き込みます。
テストで LLM 呼び出しのモックを作成することは、脆弱で退屈なことで有名です。 Orchid は、モック管理を単純な記録フローに変換します。
テスト スイートをキャプチャ モードで 1 回実行して、JSON フィクスチャを生成します。
フィクスチャをリポジトリにコミットします。
フィクスチャを使用して CI をリプレイ モードで実行します。テストはオフラインで即座に実行され、API コストはゼロになります。
リプレイはローカル記録からの応答をほぼゼロのレイテンシーで提供するため、独自のコードのパフォーマンスも分離します。ネットワーク呼び出しとアップストリーム API の差異を考慮してエージェント ロジックをプロファイルまたはベンチマークし、実行ごとに再現可能な数値を取得します。
プロキシはポート 4321 に React ベースのダッシュボードを埋め込みます。追加でインストールするものは何もありません。モデル、プロバイダー、ステータス、またはプロンプト キーワードによって交換を検索およびフィルタリングします。セッション全体でのトークンの使用量とコストを比較します。セッションを移植可能な JSON フィクスチャとしてエクスポートします。
内蔵 MCP サーバー (SSE) を使用すると、Cursor、VS Code、Claude Code などの AI アシスタントが記録されたトラフィックを直接クエリできます。プロンプトのパフォーマンスを分析したり、トークン/コスト統計を取得したり、コード編集のコンテキストとしてペイロード サンプルを取得したりできます。
クイック スタート: Orchid プロキシを実行する
プロキシは、マルチ アーキテクチャ コンテナ イメージ (Apple Silicon arm64 および Linux amd64 ) として出荷されます。
安定版: ghcr.io/mario-guerra/orchid-proxy:latest
docker pull ghcr.io/mario-guerra/orchid-proxy:latest
1. APIキーを生成する
docker run --rm ghcr.io/mario-guerra/orchid-proxy:latest generated-api-key
2.プロキシを開始します
コンテナを起動し、生成されたキーを ORCHID_API_KEY 環境変数として渡します。
API キーは Docker では必須です : Orchid プロキシ コンテナは 0.0.0.0 にバインドされます (または

CHID_BIND_HOST=0.0.0.0 ) )、デフォルトでネットワーク トラフィックを受信できるようになります。ローカルホスト以外のアドレスにバインドされるため、Docker 内で実行する場合は ORCHID_API_KEY の設定が必須です。 ORCHID_API_KEY を設定せずにコンテナを起動しようとすると、セキュリティ上の理由から、プロキシは起動時にクラッシュ終了します。
パブリック ルートを除外する: ヘルス チェック エンドポイント ( /health ) とクエリ ポート ( 4321 ) 上の静的ビジュアライザー Web アセット (HTML、JS、CSS) にはキーは必要ありません。これにより、ビジュアライザ UI をブラウザに読み込むことができますが、セッション データを読み込むにはキーが必要です (画面にキーの入力を求めるプロンプトが表示されます)。ポート 4320 上のすべてのプロキシ トラフィックとデータ API エンドポイント (ポート 4321 の /v1/* および /api/* の下) は厳密に認証ゲートされており、キーが必要です。
docker run -d \
--name 蘭プロキシ \
-p 4320:4320 \
-p 4321:4321 \
-v 蘭データ:/データ \
-e ORCHID_API_KEY=あなたの安全な API キー \
-e ORCHID_DB_PATH=/data/orchid.db \
ghcr.io/mario-guerra/orchid-proxy:最新
記録プロキシ : http://localhost:4320/v1 (X-Orchid-Api-Key: your-secure-api-key を渡します — Authorization ヘッダーはそのまま上流プロバイダーに転送されます)。 LLM エンドポイントと、それを経由してルーティングされるその他の HTTP API で機能します。
クエリ API / ビジュアライザー UI : http://localhost:4321
3. アプリをプロキシに向けます
リクエストをインターセプトして記録するには、ローカル プロキシ経由でアウトバウンド HTTP トラフィックをルーティングするようにアプリケーションを構成する必要があります。これは 2 つの方法で行うことができます。
軽量 SDK をインストールし、アプリケーションのライフサイクルの最初に初期化ヘルパーを呼び出します。これにより、グローバル HTTP/HTTPS トランスポート クライアントに自動的にパッチが適用されます。
Python : pip install orchid-sdk
輸入蘭
蘭。 init () # LLM クライアントを初期化する前に呼び出す必要があります
TypeScript/ノード: npm install orchid-sdk
輸入する

「orchid-sdk」からの t { init } ;
init () を待ちます。 // LLM クライアントを初期化する前に呼び出す必要があります
SDK を使用したくない場合 (または Go、Java、または Ruby を使用している場合)、LLM クライアントを直接構成できます。 API ベース URL をプロキシ アドレス ( http://localhost:4320/v1 ) に設定し、X-Orchid-Api-Key ヘッダーを使用して生成された API キーを渡します。
Python の例 (OpenAI クライアント):
openaiインポートからOpenAI
クライアント = OpenAI (
Base_url = "http://localhost:4320/v1" ,
default_headers = { "X-Orchid-Api-Key" : "あなたの安全な API キー" }
）
その他の構成オプションと高度なセットアップ (クラウド展開テンプレートを含む) については、 docs/getting_started.md および docs/configuration.md を参照してください。
4. AI アシスタント (MCP) を接続する
プロキシの MCP サーバーを Cursor、VS Code、または Claude Code に接続すると、コーディング エージェントはアプリの記録された LLM トラフィックを直接確認できるようになります。
以下の構成を IDE の mcp_config.json (例: ~/.gemini/antigravity-ide/mcp_config.json または ~/Library/Application Support/Claude/mcp_config.json ) に追加します。
{
"mcpサーバー": {
"オーキッドローカル" : {
"コマンド" : " docker " ,
"引数" : [
" 実行 " 、
" -i " 、
" orchid-proxy " ,
" orchid-proxy " ,
「 --mcp 」
】
}
}
}
これにより、実行中のまたは内部で MCP サーバー プロセスが直接実行されます。

[切り捨てられた]

## Original Extract

Orchid - Orchestration interactive debugger - Record, inspect, & replay AI agents - mario-guerra/orchid-trace

Orchid (Orchestration interactive debugger) is a zero-instrumentation proxy that captures every API & LLM call in your agent pipeline, then lets you inspect and replay the entire run locally, step by step. No instrumentation, no vendor lock-in, no cloud dependency. It also provides a visual inspector and MCP server, so you can inspect the session yourself or use your favorite agentic coding IDE to debug your agent runs. I built it because I was tired of debugging agent failures by grepping through logs, and the available AI observability tools all seemed to require intrusive instrumentation and/or sending my prompts and responses to a cloud service. I wanted something that would let me debug agent runs locally, without having to worry about vendor lock-in or data privacy. Orchid is that tool. The call inspection features work extremely well, at least for my use cases, but the replay feature is perhaps more interesting. It makes LLM pipeline testing deterministic without mocking or re-running expensive API calls. Free, self-hosted, runs on your machine or infrastructure: https://github.com/mario-guerra/orchid-trace Would love feedback from anyone building multi-step agentic systems or struggling with non-deterministic LLM test failures.

GitHub - mario-guerra/orchid-trace: Orchid - Orchestration interactive debugger - Record, inspect, & replay AI agents · GitHub
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
mario-guerra
/
orchid-trace
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
46 Commits 46 Commits .github/ workflows .github/ workflows assets assets docs docs sdk sdk LICENSE LICENSE README.md README.md View all files Repository files navigation
Orchid — Record, Inspect, Replay AI Agents
Stop grepping logs. Orchid records your agent's network traffic — LLM calls, tool invocations, and any other API your agent talks to — through a zero-instrumentation proxy. Then it lets you:
Time-travel through completed runs, step by step
Inspect every prompt, response, token count, and cost
Debug failures in the built-in web UI or via MCP tools from your IDE
Replay recorded runs offline — deterministic tests with zero API cost
Orchid gives your coding agent the ability to debug your AI app. The proxy has a built-in MCP server, so when an LLM call is buried deep in your stack — behind a framework, a queue, three layers of abstraction — your AI assistant in Cursor, VS Code, or Claude Code can query the recorded traffic directly. No print statements, no log spelunking: ask your agent "why did this run fail?" and it can go look, figure out why, and fix it for you.
You choose how much to record: route only your LLM traffic through the proxy for lightweight inspection, or capture everything for the full picture. To replay a run with perfect fidelity, all of the agent's network traffic must go through the proxy — replay works by serving back the recorded responses, so anything that wasn't recorded can't be replayed.
IMPORTANT NOTE - YOUR DATA NEVER LEAVES YOUR INFRASTRUCTURE!
The proxy forwards requests only to the upstream APIs your app was already calling.
Everything recorded stays in a local SQLite database inside the container (or your mounted volume). No phone-home, no telemetry, no cloud backend.
Secrets are scrubbed in memory before anything is written to disk: Authorization headers are forwarded untouched to the upstream but never stored, and headers, query strings, and body fields with secret-like names (keys, tokens, passwords, credentials, cookies) are stored as [REDACTED] .
One honest caveat: redaction works by recognizing field names (like api_key or authorization ), not by scanning the contents of your prompts. Prompt and completion text is recorded verbatim — that's the whole point of Orchid — so if a secret is pasted into a prompt, it will be stored along with the rest of the prompt text.
This repository contains the open-source Orchid SDKs and user documentation. The orchid-proxy container is distributed via the GitHub Container Registry (see below). Content here is synced automatically from the main development repository — issues and discussions are welcome; pull requests may be ported rather than merged directly.
Want to see Orchid in action without instrumenting your own app? Check out the Orchid LangGraph Demo .
This repository showcases a complex, multi-agent AI architecture (using OpenAI, Anthropic, and Google Vertex AI) natively hooked into Orchid. It includes a complete trace fixture, allowing you to execute the entire adversarial agent pipeline in offline replay mode with zero configuration—no API keys required and zero upstream cost. It's the absolute fastest way to experience deterministic AI testing and test-drive the visualizer dashboard.
flowchart LR
subgraph env["Your Environment (laptop, on-prem, or your own cloud)"]
app["Application<br/>(Python, TS, ...)"]
proxy["orchid-proxy<br/>:4320 (proxy)<br/>:4321 (query / UI / MCP)"]
db[("orchid.db<br/>(SQLite)")]
app -- "HTTP" --> proxy
proxy -- "record" --> db
db -. "replay" .-> proxy
end
upstream["Upstream APIs<br/>OpenAI, Anthropic,<br/>Gemini, tools, any API ..."]
proxy -- "HTTPS<br/>(skipped in replay mode)" --> upstream
Loading
Non-Intrusive Interception (Thin SDK)
Unlike traditional LLM observability tools that require wrapping every client initialization or using heavy SDKs with AST modifications, Orchid uses an APM-style Thin SDK . The SDK patches the foundational HTTP transport layer ( httpx / requests in Python, fetch in Node), so every LLM call made by standard client libraries is automatically routed through the local or remote Orchid Proxy — without changing your prompt-handling or generation code.
The proxy does not keep track of application state. It reads X-Orchid-* HTTP headers (injected by the SDK, or set manually from any language) to decide how to process each request:
passthrough : Transparent reverse proxy. Forwards the request and returns the response without writing anything to disk.
capture : Forwards the request, serializes the complete request/response payloads (including streaming chunks), calculates costs, and saves them to a local SQLite database under a specific Session ID .
replay : Blocks all outbound network traffic. Hashes the incoming request and serves the exact matching recorded response from SQLite. If no match is found, returns a deterministic mock error.
Every captured LLM call (or "Exchange") records:
Request Metadata : System prompts, user prompts, temperature, top-p, and custom tags.
Response Telemetry : Complete completion text, usage tokens (input/output), and latency.
Cost Calculation : Real-time USD cost attribution based on user-supplied model pricing maps (see docs/configuration.md and the /pricing endpoint in docs/api_reference.md ).
Stream Reassembly : For streaming completions, Orchid buffers SSE chunks in memory, serving them to the client instantly, and writes the fully reassembled completion body to SQLite.
Writing mocks for LLM calls in tests is notoriously fragile and tedious. Orchid converts mock management into a simple recording flow:
Run your test suite once in capture mode to generate a JSON fixture.
Commit the fixture to your repository.
Run CI in replay mode using the fixture. Your tests now execute instantly, offline, and with zero API cost.
Because replay serves responses from the local recording with near-zero latency, it also isolates your own code's performance : profile or benchmark your agent logic with network calls and upstream API variance taken out of the equation, and get reproducible numbers run after run.
The proxy embeds a React-based dashboard on port 4321 — nothing extra to install. Search and filter exchanges by model, provider, status, or prompt keywords; compare token usage and costs across sessions; export sessions as portable JSON fixtures.
A built-in MCP server (SSE) lets AI assistants like Cursor, VS Code, or Claude Code query the recorded traffic directly: analyze prompt performance, pull token/cost statistics, or fetch payload examples as context for editing code.
Quick Start: Run the Orchid Proxy
The proxy ships as a multi-arch container image (Apple Silicon arm64 and Linux amd64 ):
Stable : ghcr.io/mario-guerra/orchid-proxy:latest
docker pull ghcr.io/mario-guerra/orchid-proxy:latest
1. Generate an API key
docker run --rm ghcr.io/mario-guerra/orchid-proxy:latest generate-api-key
2. Start the proxy
Start the container and pass your generated key as the ORCHID_API_KEY environment variable.
API Key is Mandatory in Docker : The Orchid Proxy container binds to 0.0.0.0 ( ORCHID_BIND_HOST=0.0.0.0 ) by default so that it can receive network traffic. Because it binds to a non-localhost address, setting ORCHID_API_KEY is mandatory when running inside Docker. If you attempt to start the container without setting ORCHID_API_KEY , the proxy will crash-exit on startup for security reasons.
Exempt Public Routes : The health check endpoint ( /health ) and the static visualizer web assets (HTML, JS, CSS) on the query port ( 4321 ) do not require the key. This allows the visualizer UI to load in your browser, but it requires the key to load any session data (the screen will prompt you to enter the key). All proxying traffic on port 4320 and data API endpoints (under /v1/* and /api/* on port 4321 ) are strictly auth-gated and require the key.
docker run -d \
--name orchid-proxy \
-p 4320:4320 \
-p 4321:4321 \
-v orchid-data:/data \
-e ORCHID_API_KEY=your-secure-api-key \
-e ORCHID_DB_PATH=/data/orchid.db \
ghcr.io/mario-guerra/orchid-proxy:latest
Recording Proxy : http://localhost:4320/v1 (pass X-Orchid-Api-Key: your-secure-api-key — the Authorization header is forwarded untouched to the upstream provider). Works for LLM endpoints and any other HTTP API you route through it.
Query API / Visualizer UI : http://localhost:4321
3. Point your app at the proxy
To intercept and record requests, you must configure your application to route its outbound HTTP traffic through the local proxy. You can do this in two ways:
Install the lightweight SDK and call the initialization helper at the very beginning of your application lifecycle. This automatically patches the global HTTP/HTTPS transport clients.
Python : pip install orchid-sdk
import orchid
orchid . init () # Must be called before initializing any LLM client
TypeScript/Node : npm install orchid-sdk
import { init } from "orchid-sdk" ;
await init ( ) ; // Must be called before initializing any LLM client
If you prefer not to use the SDK (or are using Go, Java, or Ruby), you can configure your LLM client directly. Set the API base URL to the proxy address ( http://localhost:4320/v1 ) and pass your generated API key using the X-Orchid-Api-Key header:
Python example (OpenAI Client) :
from openai import OpenAI
client = OpenAI (
base_url = "http://localhost:4320/v1" ,
default_headers = { "X-Orchid-Api-Key" : "your-secure-api-key" }
)
For more configuration options and advanced setups (including cloud deployment templates), see docs/getting_started.md and docs/configuration.md .
4. Connect your AI assistant (MCP)
Hook the proxy's MCP server into Cursor, VS Code, or Claude Code and your coding agent gets direct visibility into your app's recorded LLM traffic.
Add the configuration below to your IDE's mcp_config.json (e.g., ~/.gemini/antigravity-ide/mcp_config.json or ~/Library/Application Support/Claude/mcp_config.json ):
{
"mcpServers" : {
"orchid-local" : {
"command" : " docker " ,
"args" : [
" exec " ,
" -i " ,
" orchid-proxy " ,
" orchid-proxy " ,
" --mcp "
]
}
}
}
This runs the MCP server process directly inside the running or

[truncated]
