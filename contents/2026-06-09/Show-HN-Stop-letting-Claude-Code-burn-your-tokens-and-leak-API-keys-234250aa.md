---
source: "https://github.com/ashp15205/guardian-runtime"
hn_url: "https://news.ycombinator.com/item?id=48464199"
title: "Show HN: Stop letting Claude Code burn your tokens and leak API keys"
article_title: "GitHub - ashp15205/guardian-runtime: A zero-latency, local-first runtime firewall for LLMs. Intercept every prompt and response locally to stop data leaks and runaway token costs. · GitHub"
author: "ashishp15"
captured_at: "2026-06-09T18:51:19Z"
capture_tool: "hn-digest"
hn_id: 48464199
score: 1
comments: 0
posted_at: "2026-06-09T17:20:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Stop letting Claude Code burn your tokens and leak API keys

- HN: [48464199](https://news.ycombinator.com/item?id=48464199)
- Source: [github.com](https://github.com/ashp15205/guardian-runtime)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T17:20:56Z

## Translation

タイトル: HN を表示: クロード コードによるトークンの書き込みと API キーの漏洩を停止する
記事のタイトル: GitHub - ashp15205/guardian-runtime: LLM 向けのゼロ遅延、ローカルファースト ランタイム ファイアウォール。すべてのプロンプトと応答をローカルで傍受し、データ漏洩とトークンコストの暴走を防ぎます。 · GitHub
説明: LLM 用の遅延ゼロ、ローカルファーストのランタイム ファイアウォール。すべてのプロンプトと応答をローカルで傍受し、データ漏洩とトークンコストの暴走を防ぎます。 - ashp15205/ガーディアンランタイム

記事本文:
GitHub - ashp15205/guardian-runtime: LLM 用のゼロ遅延、ローカルファースト ランタイム ファイアウォール。すべてのプロンプトと応答をローカルで傍受し、データ漏洩とトークンコストの暴走を防ぎます。 · GitHub
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
アッシュ15205
/
ガーディアンランタイム
公共
通知
Y

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
54 コミット 54 コミット .github .github docs docs ガーディアン_ランタイム ガーディアン_ランタイム ポリシー ポリシー スクリプト テスト スクリプト テスト .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md Article.md Article.md Concept.md Concept.md Index.html Index.html pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI アプリケーション向けのゼロ遅延 FinOps およびセキュリティ ファイアウォール。
すべてのプロンプトと応答をローカルでインターセプトします。データ漏洩とトークンコストの暴走を阻止します。
🌐 ウェブサイトとドキュメント: https://ashp15205.github.io/guardian-runtime/
📦 PyPI で利用可能: https://pypi.org/project/guardian-runtime/
🛑 中心的な問題: なぜガーディアンが必要なのか
🟢 解決策: Guardian ランタイムとは何ですか?
🎯 包括的な使用例 (どこでどのように使用するか)
1. ターミナルコーディングエージェント (クロードコード、エイダー)
2. ビジュアル IDE (カーソル、ウィンドサーフィン)
3. 本番環境の Python アプリケーション (SDK)
4. エージェントフレームワーク (LangChain、AutoGen)
5. Web UI のデータ準備 (ドキュメント変換)
💻 完全な CLI コマンド リファレンス
⚙️ 詳細構成 (ポリシー YAML)
🛑 中心的な問題: なぜガーディアンが必要なのか
AI コーディング エージェント (Claude Code、Cursor、Aider) が標準の開発者ツールになるにつれて、2 つの大きな隠れたリスクと 1 つの規制上の悩みが生じます。
💸 1. FinOps のリスク: コストの暴走
自律エージェントはループで動作します。エージェントがバグ修正の再試行中に行き詰まったり、誤って 1 GB の巨大なログ ファイルをコンテキスト ウィンドウにダンプしたりすると、一晩で 100 ドルの API 請求が発生する可能性があります。
問題: プロバイダーの請求書が届くまでは、セッション コストをまったく把握できず、制御することもできません。

月末に。
🔒 2. セキュリティ リスク: データ漏洩
コーディング エージェントを使用するには、完全なローカル コードベース アクセスが必要です。ただし、AWS_SECRET_KEY またはデータベース パスワードを誤って .env ファイルに残した場合、エージェントはそれをサードパーティ LLM プロバイダー (OpenAI、Anthropic) にサイレントにアップロードします。
問題: 現在の可観測性ツール (Langfuse など) は、認証情報がすでにクラウドに到達した後にのみ漏洩を記録します。
🏛 3. コンプライアンスのリスク (簡単に)
未承認の PII (テスト データベース内の SSN や電子メールなど) を外国 LLM API に送信すると、GDPR および DPDP 規制に違反します。
🟢 解決策: Guardian ランタイム
Guardian Runtime は、ローカルファーストのセキュリティ ミドルウェアおよび FinOps ファイアウォールです。これは完全にローカル マシン上で実行され、インフラストラクチャを離れる前に LLM トラフィックを傍受します。
🏗 アーキテクチャとセキュリティ パイプライン
Guardian はネットワーク層または SDK 経由でトラフィックを傍受し、クラウドに到達する前に厳格な検証パイプラインを通過させます。
エージェント/開発者 Guardian ランタイム クラウド LLM
│ │ │
│ 1. プロンプト + コンテキスト │ │
│ ──------------------------------------------▶ │ │
│ │ │
│ │ [セキュリティファイアウォール] │
│ │ §─ AWS キー/シークレットをスキャン │
│ │ └─ 脅威が検出された場合はブロック ──┼─ (リクエストのドロップ)
│ │ │
│ │ [トークンオプティマイザー] │
│ │ §─ 空白を圧縮 │
│ │ └─ 簡潔モード（出力トリム） │
│ │ │
│ │ [FinOps 予算] │
│ │ §─ 1 日のご利用限度額を確認する │
│ │ └─ $5 リミットに達したらブロック ───┼─ (ドロップリクエスト)
│ │ │
│ │ 2. サニタイズされたプロンプト │
│ │ ─────────────────── ▶ │
│ │ │
│ │ 3. LLM 応答

│
│ │ ◀────────────── │
│ │ │
│ │ [アウトプットガード] │
│ │ 漏洩した PII/秘密の監査 │
│ │ │
│ 4. 安全な対応 │ │
│ ◀───────────── │ │
│ │ │
🔌 サポートされている統合
Guardian Runtime は HTTP プロキシまたはネイティブ Python SDK として機能します。つまり、内部コードを変更することなく、ほとんどすべての最新の AI ツールと簡単に統合できます。
ビジュアル IDE: カーソル、ウィンドサーフィン、VS コード (Cline/RooCode 経由)
ターミナル エージェント: Claude Code、Aider、GitHub Copilot CLI
フレームワーク: LangChain、AutoGen、LlamaIndex、CrewAI
LLM プロバイダー: OpenAI、Anthropic、Google Gemini (OpenAI 互換性レイヤー経由)
# コアフレームワークのみ
pip インストール ガーディアン_ランタイム
# または、特定の LLM プロバイダーを使用してインストールします。
pip install " Guardian_runtime[openai] "
pip install " Guardian_runtime[anthropic] "
pip install " Guardian_runtime[gemini] "
# またはすべてをインストールします (プロバイダー、ML スキャナー、ドキュメント コンバーター)。
pip install " Guardian_runtime[all] "
完了しました。サインアップもキーも設定も必要ありません。すべての監視データは、ローカル マシンの ~/.guardian_runtime/ に残ります。
🎯 包括的な使用例 (どこでどのように使用するか)
Guardian はユニバーサルであるように設計されています。ワークフローに基づいてデプロイする正確な方法を次に示します。
1. ターミナルコーディングエージェント (クロードコード、エイダー)
なぜここで使用するのでしょうか？ CLI エージェントは自律的に動作します。本番環境の AWS キーを含む .env ファイルを誤って読み取り、それをコンテキストとして Anthropic/OpenAI に送信する可能性があります。 Guardian はこれを防ぎ、エージェントが予算を超過しないようにします。
バックグラウンド端末でプロキシを開始します。
ガーディアン_ランタイム プロキシ --ポート 8080
e を使用してプロキシ経由でトラフィックをルーティングするようにエージェントに指示します。

環境変数:
PowerShell の場合:
$ env: ANTHROPIC_BASE_URL = " http://localhost:8080 "
クロード
Mac/Linux/Git Bash の場合:
エクスポート ANTHROPIC_BASE_URL=http://localhost:8080
クロード
2. ビジュアル IDE (カーソル、ウィンドサーフィン)
なぜここで使用するのでしょうか？ Cursor のような最新の GUI エディターは、深いコードベースにアクセスできます。コーディング中に、シークレットを含むファイルを強調表示して、「このファイルについて説明してください」と尋ねることがあるかもしれません。 Guardian は、Cursor がそのシークレットをクラウドに送信するのを阻止します。
ターミナルでプロキシを開始します: Guardian_runtime proxy --port 8080
カーソル設定を開く ( Cmd/Ctrl + 、 )
[モデル] > [ベース URL の上書き] に移動します。
ベース URL を http://localhost:8080 に設定します。
(これで、Cursor のトラフィックはすべてローカルで保護され、追跡されるようになりました。)
3. 本番環境の Python アプリケーション (SDK)
なぜここで使用するのでしょうか？運用チャットボットまたは RAG パイプラインを構築している場合は、ユーザーが「脱獄」プロンプト インジェクションを実行したり、LLM をだまして内部システム プロンプトを漏洩させたりできないようにする必要があります。
使用方法:
OpenAI/Anthropic SDK のドロップイン代替品として Guardian を使用します。
OSをインポートする
Guardian_runtime からのインポート GuardianRuntime 、 GuardianRuntimeBlockedError
OS 。 environ [ "OPENAI_API_KEY" ] = "sk-proj-..."
gr = GuardianRuntime () # ゼロ構成の初期化
試してみてください:
# OpenAI に送信する前にユーザー入力を保護します
応答 = gr 。完了（
メッセージ = [{ "ロール" : "ユーザー" , "コンテンツ" : "私の AWS キーは AKIAIOSFODNN7EXAMPLE です" }],
raise_on_block = True
)
print (応答の内容)
e としての GuardianRuntimeBlockedError を除く:
# 秘密を漏らすのではなく、アプリできれいに失敗します!
print ( f"ローカルでブロックされました: { e . 応答 . 違反 [ 0 ]. 詳細 } " )
4. エージェントフレームワーク (LangChain、AutoGen)
なぜここで使用するのでしょうか？複数の通信エージェントを生成するフレームワークは、トークンを急速に消費する可能性があります。 Guardian は、すべてのエージェント ノードの中央コスト追跡ハブとして機能します。
H

使用する必要があります:
フレームワークの Base_url がローカル プロキシを指すようにします。
langchain_openai から ChatOpenAI をインポート
llm = ChatOpenAI (
モデル = "gpt-4o" 、
Base_url = "http://localhost:8080" , # Guardian を経由するトラフィック ルート
api_key = "sk-proj-..."
)
応答 = llm 。呼び出す (「こんにちは、ガーディアン!」)
5. Web UI のデータ準備 (ドキュメント変換)
なぜここで使用するのでしょうか？標準の ChatGPT または Claude Web UI を使用する場合、PDF には膨大な量の隠れた書式設定が含まれているため、大きな PDF をアップロードするとコンテキスト ウィンドウがすぐに消費されてしまいます。
使用方法:
組み込みの CLI を使用して、フォーマットの肥大化を取り除き、ドキュメントを手動でアップロードする前に純粋な Markdown に圧縮します。
Guardian_runtime Convert < path/to/input.pdf > --out < path/to/output.md >
Cleaned_report.md を ChatGPT にアップロードできるようになり、大量のコンテキスト スペースが節約され、幻覚が防止されます。
💻 包括的な CLI コマンド リファレンス
Guardian には、オフライン CLI ツールの強力なスイートが付属しています。すべてのデータは純粋にローカルの ~/.guardian_runtime/ に保存されます。
以下では、すべてのコマンドとそのフラグ、およびそれを使用する正確な方法と理由について詳しく説明します。
Guardian_runtime プロキシ (セキュリティ ファイアウォール)
ローカル HTTP インターセプト サーバーを起動します。これは、ソース コードを編集できないツール (カーソル コードやクロード コードなど) を保護するためのコア エンジンです。
--port, -p <int> : リッスンするポート (デフォルト: 8080)。
--host <str> : バインド先のホスト。ローカル ネットワーク上で公開するには 0.0.0.0 を使用します (デフォルト: 127.0.0.1 )。
--policy <path> : カスタムのpolicy.yamlファイルへのパス。省略した場合、デフォルトの Zero-Config ポリシー (予算 10 ドル) が使用されます。
--reload : ポリシー ファイルが変更された場合に自動リロードを有効にします (開発モードで役立ちます)。
$ Guardian_runtime プロキシ --ポート 8080
⛨ GuardianRuntime ランタイム プロキシ
───────────

───────────
リスニング: http://127.0.0.1:8080
ポリシー : デフォルト (ゼロ構成)
ダッシュボード : Guardian_runtime ダッシュボード (別のターミナルで実行)
エージェントのセットアップ:
クロード コード → ANTHROPIC_BASE_URL=http://localhost:8080 クロード
補助者 → OPENAI_BASE_URL=http://localhost:8080 補助者
カーソル→設定→APIベース→http://localhost:8080
Guardian_runtime 変換 <パス> (ドキュメント分析)
大量の PDF、DOCX、XLSX ファイルを、高度に圧縮されたトークン最適化された Markdown に変換します。
なぜこれを使用するのでしょうか?生の PDF を Web UI (ChatGPT など) にアップロードしたり、エージェントで解析したりすると、隠れたフォーマットの肥大化によって何千ものトークンが無駄になります。このコマンドは、LLM コンテキスト ウィンドウに到達する前に肥大化を取り除きます。
<path> : 圧縮するドキュメントへの絶対パスまたは相対パス。
--out, -o <path> : 変換された Markdown の出力ファイル パス。省略した場合、プレビューを端末に出力します。
$ Guardian_runtime Convert < path/to/input_file > --out < path/to/output_file.md >
⛨ GuardianRuntime ドキュメントコンバータ
処理中: 入力ファイル...
✓ 変換完了！
• 元のファイル: input_file
• トークン数: 14,205
• 保存先：output_file.md
Guardian_runtime scan <テキスト> (手動脅威検証)
ML InputGuard スキャナーと Regex スキャナーを使用して、特定のテキスト文字列に対してローカル セキュリティ スキャンを実行します。
なぜこれを使用するのでしょうか?これは、大規模なコードベースをエージェントに送信する前に、ファイアウォールが何をキャッチするかを正確に確認する場合、または PII/秘密の検出の感度をテストする場合に使用します。
$ Guardian_runtime scan " 私の AWS キーは AKIAIOSFODNN7EXAMPLE です "
🛑 スキャンに失敗しました!検出された脅威:
- [HIGH] Secret_detected: AWS アクセス キー ID が見つかりました。
Guardian_runtime 分析 (FinOps 追跡)
今日の API コスト、トークン使用量、インターセプトの美しいターミナル概要を出力します。

テッドさん

[切り捨てられた]

## Original Extract

A zero-latency, local-first runtime firewall for LLMs. Intercept every prompt and response locally to stop data leaks and runaway token costs. - ashp15205/guardian-runtime

GitHub - ashp15205/guardian-runtime: A zero-latency, local-first runtime firewall for LLMs. Intercept every prompt and response locally to stop data leaks and runaway token costs. · GitHub
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
ashp15205
/
guardian-runtime
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
54 Commits 54 Commits .github .github docs docs guardian_runtime guardian_runtime policies policies scripts scripts tests tests .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md article.md article.md concept.md concept.md index.html index.html pyproject.toml pyproject.toml View all files Repository files navigation
A Zero-Latency FinOps & Security Firewall for AI Applications.
Intercept every prompt and response locally. Stop data leaks and runaway token costs.
🌐 Website & Docs: https://ashp15205.github.io/guardian-runtime/
📦 Available on PyPI: https://pypi.org/project/guardian-runtime/
🛑 The Core Problem: Why You Need Guardian
🟢 The Solution: What is Guardian Runtime?
🎯 Comprehensive Use Cases (Where & How to Use)
1. Terminal Coding Agents (Claude Code, Aider)
2. Visual IDEs (Cursor, Windsurf)
3. Production Python Applications (SDK)
4. Agentic Frameworks (LangChain, AutoGen)
5. Data Prep for Web UIs (Document Conversion)
💻 Complete CLI Command Reference
⚙️ Advanced Configuration (Policy YAML)
🛑 The Core Problem: Why You Need Guardian
As AI coding agents (Claude Code, Cursor, Aider) become standard developer tools, they introduce two massive, hidden risks, and one regulatory headache:
💸 1. The FinOps Risk: Cost Runaways
Autonomous agents operate in loops. If an agent gets stuck retrying a bug fix or accidentally dumps a massive 1GB log file into its context window, you can wake up to a $100 API bill overnight .
The Problem: You have zero visibility or control over session costs until the provider's bill arrives at the end of the month.
🔒 2. The Security Risk: Data Exfiltration
Coding agents require full local codebase access to be useful. However, if you accidentally leave an AWS_SECRET_KEY or a database password in a .env file, the agent will silently upload it to a third-party LLM provider (OpenAI, Anthropic).
The Problem: Current observability tools (like Langfuse) only log the leak after the credentials have already reached the cloud.
🏛 3. The Compliance Risk (Briefly)
Sending unauthorized PII (like SSNs or emails in a test database) to foreign LLM APIs violates GDPR and DPDP regulations.
🟢 The Solution: Guardian Runtime
Guardian Runtime is a local-first security middleware and FinOps firewall . It runs entirely on your local machine and intercepts LLM traffic before it leaves your infrastructure.
🏗 Architecture & The Security Pipeline
Guardian intercepts traffic at the network layer or via SDK, passing it through a strict verification pipeline before it ever reaches the cloud.
Agent / Dev Guardian Runtime Cloud LLM
│ │ │
│ 1. Prompt + Context │ │
│ ──────────────────────────▶ │ │
│ │ │
│ │ [Security Firewall] │
│ │ ├─ Scan AWS Keys / Secrets │
│ │ └─ Block if Threat Detected ──┼─ (Drops Request)
│ │ │
│ │ [Token Optimizer] │
│ │ ├─ Compress Whitespace │
│ │ └─ Terse Mode (Output Trim) │
│ │ │
│ │ [FinOps Budget] │
│ │ ├─ Check Daily Spend Limit │
│ │ └─ Block if $5 Limit Hit ─────┼─ (Drops Request)
│ │ │
│ │ 2. Sanitized Prompt │
│ │ ────────────────────────────▶ │
│ │ │
│ │ 3. LLM Response │
│ │ ◀──────────────────────────── │
│ │ │
│ │ [Output Guard] │
│ │ Audit for Leaked PII/Secrets │
│ │ │
│ 4. Safe Response │ │
│ ◀────────────────────────── │ │
│ │ │
🔌 Supported Integrations
Guardian Runtime acts as an HTTP proxy or a native Python SDK, meaning it integrates effortlessly with almost any modern AI tool without modifying their internal code.
Visual IDEs: Cursor, Windsurf, VS Code (via Cline/RooCode)
Terminal Agents: Claude Code, Aider, GitHub Copilot CLI
Frameworks: LangChain, AutoGen, LlamaIndex, CrewAI
LLM Providers: OpenAI, Anthropic, Google Gemini (via OpenAI compatibility layer)
# Core framework only
pip install guardian_runtime
# Or install with specific LLM providers:
pip install " guardian_runtime[openai] "
pip install " guardian_runtime[anthropic] "
pip install " guardian_runtime[gemini] "
# Or install everything (Providers, ML Scanner, Document Converter):
pip install " guardian_runtime[all] "
Done. No signup, no keys, zero configuration required. All monitoring data stays on your local machine in ~/.guardian_runtime/ .
🎯 Comprehensive Use Cases (Where & How to Use)
Guardian is designed to be universal. Here are the exact ways to deploy it based on your workflow.
1. Terminal Coding Agents (Claude Code, Aider)
Why use it here? CLI agents operate autonomously. They can accidentally read a .env file containing your production AWS keys and send it to Anthropic/OpenAI as context. Guardian prevents this and ensures the agent doesn't blow your budget.
Start the proxy in a background terminal:
guardian_runtime proxy --port 8080
Tell your agent to route traffic through the proxy using environment variables:
In PowerShell:
$ env: ANTHROPIC_BASE_URL = " http://localhost:8080 "
claude
In Mac/Linux/Git Bash:
export ANTHROPIC_BASE_URL=http://localhost:8080
claude
2. Visual IDEs (Cursor, Windsurf)
Why use it here? Modern GUI editors like Cursor have deep codebase access. While coding, you might highlight a file containing a secret and ask "explain this file". Guardian stops Cursor from sending that secret to the cloud.
Start the proxy in your terminal: guardian_runtime proxy --port 8080
Open Cursor Settings ( Cmd/Ctrl + , )
Navigate to Models > Override Base URL
Set the Base URL to: http://localhost:8080
(Now all of Cursor's traffic is protected and tracked locally!)
3. Production Python Applications (SDK)
Why use it here? If you are building a production chatbot or RAG pipeline, you must ensure your users cannot perform "jailbreak" prompt injections or trick the LLM into leaking internal system prompts.
How to use:
Use Guardian as a drop-in replacement for the OpenAI/Anthropic SDK.
import os
from guardian_runtime import GuardianRuntime , GuardianRuntimeBlockedError
os . environ [ "OPENAI_API_KEY" ] = "sk-proj-..."
gr = GuardianRuntime () # Zero-config initialization
try :
# Protects user input before sending to OpenAI
response = gr . complete (
messages = [{ "role" : "user" , "content" : "My AWS Key is AKIAIOSFODNN7EXAMPLE" }],
raise_on_block = True
)
print ( response . content )
except GuardianRuntimeBlockedError as e :
# Fails cleanly in your app instead of leaking the secret!
print ( f"Blocked Locally: { e . response . violations [ 0 ]. detail } " )
4. Agentic Frameworks (LangChain, AutoGen)
Why use it here? Frameworks that spawn multiple communicating agents can rapidly consume tokens. Guardian acts as a central cost-tracking hub for all agent nodes.
How to use:
Point your framework's base_url to the local proxy.
from langchain_openai import ChatOpenAI
llm = ChatOpenAI (
model = "gpt-4o" ,
base_url = "http://localhost:8080" , # Traffic routes through Guardian
api_key = "sk-proj-..."
)
response = llm . invoke ( "Hello, Guardian!" )
5. Data Prep for Web UIs (Document Conversion)
Why use it here? If you use the standard ChatGPT or Claude Web UI, uploading large PDFs eats up your context window quickly because PDFs contain massive amounts of hidden formatting bloat.
How to use:
Use the built-in CLI to strip out formatting bloat and compress documents into pure Markdown before manually uploading them.
guardian_runtime convert < path/to/input.pdf > --out < path/to/output.md >
You can now upload cleaned_report.md to ChatGPT, saving huge amounts of context space and preventing hallucination.
💻 Exhaustive CLI Command Reference
Guardian ships with a powerful suite of offline CLI tools. All data is stored purely locally in ~/.guardian_runtime/ .
Below is a detailed dive into every command, its flags, and exactly how and why to use it.
guardian_runtime proxy (The Security Firewall)
Starts the local HTTP interception server. This is the core engine for protecting tools that you cannot edit the source code for (like Cursor or Claude Code).
--port, -p <int> : Port to listen on (Default: 8080 ).
--host <str> : Host to bind to. Use 0.0.0.0 to expose on your local network (Default: 127.0.0.1 ).
--policy <path> : Path to a custom policy.yaml file. If omitted, uses the default Zero-Config policy ($10 budget).
--reload : Enables auto-reload if the policy file changes (useful for dev mode).
$ guardian_runtime proxy --port 8080
⛨ GuardianRuntime Runtime Proxy
─────────────────────────────────────────
Listening on : http://127.0.0.1:8080
Policy : Default (Zero-Config)
Dashboard : guardian_runtime dashboard (run in another terminal)
Agent setup:
Claude Code → ANTHROPIC_BASE_URL=http://localhost:8080 claude
Aider → OPENAI_BASE_URL=http://localhost:8080 aider
Cursor → Settings → API Base → http://localhost:8080
guardian_runtime convert <path> (Document Analysis)
Converts massive PDF, DOCX, and XLSX files into highly compressed, token-optimized Markdown.
Why use this? If you upload a raw PDF to a Web UI (like ChatGPT) or parse it in an agent, you waste thousands of tokens on hidden formatting bloat. This command strips the bloat before it hits the LLM context window.
<path> : The absolute or relative path to the document you want to compress.
--out, -o <path> : Output file path for the converted Markdown. If omitted, prints a preview to the terminal.
$ guardian_runtime convert < path/to/input_file > --out < path/to/output_file.md >
⛨ GuardianRuntime Document Converter
Processing: input_file...
✓ Conversion Complete !
• Original File: input_file
• Token Count: 14,205
• Saved to: output_file.md
guardian_runtime scan <text> (Manual Threat Verification)
Performs a local security scan on a specific text string using the ML InputGuard and Regex scanners.
Why use this? Use this to verify exactly what the firewall will catch before you send a massive codebase to an agent, or if you want to test how sensitive the PII/Secret detection is.
$ guardian_runtime scan " My AWS key is AKIAIOSFODNN7EXAMPLE "
🛑 Scan failed ! Threats detected:
- [HIGH] secret_detected: AWS Access Key ID found.
guardian_runtime analytics (FinOps Tracking)
Prints a beautiful terminal summary of today's API costs, token usage, and intercepted thr

[truncated]
