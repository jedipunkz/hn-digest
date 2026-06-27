---
source: "https://github.com/Guy-Sela/brytlog"
hn_url: "https://news.ycombinator.com/item?id=48696184"
title: "Show HN: Brytlog – AI logger"
article_title: "GitHub - Guy-Sela/brytlog: AI logger · GitHub"
author: "guy-sela"
captured_at: "2026-06-27T08:53:16Z"
capture_tool: "hn-digest"
hn_id: 48696184
score: 1
comments: 0
posted_at: "2026-06-27T08:06:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Brytlog – AI logger

- HN: [48696184](https://news.ycombinator.com/item?id=48696184)
- Source: [github.com](https://github.com/Guy-Sela/brytlog)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T08:06:02Z

## Translation

タイトル: HN を表示: Brytlog – AI ロガー
記事タイトル: GitHub - Guy-Sela/brytlog: AI ロガー · GitHub
説明: AIロガー。 GitHub でアカウントを作成して、Guy-Sela/brytlog の開発に貢献してください。

記事本文:
GitHub - Guy-Sela/brytlog: AI ロガー · GitHub
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
ガイ・セラ
/
ブライトログ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダ

RSとファイル
8 コミット 8 コミット .github/ workflows .github/ workflows src/ brytlog src/ brytlog テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md demo-no-brytlog.mp4デモ-no-brytlog.mp4 デモ-with-brytlog.mp4 デモ-with-brytlog.mp4 デモ.py デモ.py llms.txt llms.txt pyproject.toml pyproject.toml ruff.toml ruff.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Brytlog は生のログを AI サマリーに置き換えるため、開発者の時間、手間、コストを節約します。
エージェント ワークフローでは、brytlog はチーフ エージェントに対する安価で高速なプリプロセッサとして機能します。
たとえば、Claude Opus 4.8 (チーフ エージェント) は、プレーンな python run.py ではなく、 brytlog python run.py を実行する可能性があります。こうすることで、生の出力全体を独自に処理する必要 (遅く、高価で、コンテキストが肥大化する) の代わりに、Gemini-3-flash などの安価で高速なモデルによって生成された簡潔な概要のみを取得できます。
フェイルセーフとして、生のログが保存され、必要に応じてチーフ エージェントまたは開発者がアクセスできるようになります (これは設定で切り替え可能な機能です)。
安価なモデルでも完全な生ダンプを取得するのではなく、重要な部分のみを取得するため、さらに時間とお金を節約できます。
非エージェントの開発主導型ワークフローでは、brytlog を使用すると、開発者が生の出力を自分で分析したり、行をコピーしてコーディング アシスタントに貼り付けたりする時間と手間が省けます。
プラットフォーム、言語、LLM ベンダーに依存しない
最小限のセットアップ (独自のキーを持参するか、ローカルで実行するだけ)
既存のコードを変更する必要はありません (AGENTS.md に数行を追加するだけです)。
軽量 (~50 KB、~1,400 行のコード)
プライバシー重視 (brytlog はデータを収集せず、LLM に渡す前に機密情報を編集します)
デモのブライトログ.mp4
デモ-with-brytlog.mp4
デモ-json-output.mp4
使用方法
インストール後、どちらか

er は、 brytlog を使用してコマンド (例: brytlog ノード main.js ) を実行します。これにより、ターミナルでオンボーディング設定プロセスが起動されます。または、 brytlog --config を実行して、デフォルトのエディターで json 設定ファイルを開きます。
API キー (例: JQ.Ab9RN6W6QW7cmcnY92DIuoVtjCpKm_qfmO5T5oGzQmnwe5fjhw)
Google Gemini: Google AI Studio ( https://aistudio.google.com/ ) のキーを使用します。 Google Cloud Vertex AI はまだネイティブにサポートされていません。
カスタム プロバイダー: 標準の Authorization: Bearer ヘッダーを介して OpenAI 互換エンドポイント (Ollama、vLLM など) をサポートします。 Azure OpenAI はまだネイティブにサポートされていません。
カスタム モデルのみの追加必須フィールド (Ollama/Enterprise エンドポイントのオプションの上書き)
ベース URL (例: http://localhost:11434/v1 )
コマンドの前に brytlog を付けるだけです。
構文: brytlog [オプション] <コマンド> 。
ブライトログPython run.py
brytlog --api-key xyz... ノード build.js
brytlog --model クロード-haiku-4-5 ./deploy.sh
複数のコマンドを連鎖させる
複数のコマンドを一緒に実行するには、文字列全体を引用符で囲みます。
brytlog " pip install -rrequirements.txt && python run.py "
brytlog " ノード main.js
npm 実行サーブ "
エージェントワークフローでの使用
セッションの開始時にインラインでプロンプトを表示するか、これまたは同様のものを AGENTS.md に追加します。
brytlog は、生のターミナル出力を簡潔な AI 概要に置き換えます。
インタプリタ (例: `python`、`node`)、コンパイラ、ビルド ツール (例: `npm`、`make`)、およびテスト ランナー (例: `pytest`) に使用します。標準の OS ユーティリティ (例: `ls`、`cat`)、バージョン管理 (例: `git`)、または対話型 CLI ツール (例: `htop`) には使用しないでください。
構文: `brytlog [options] <command>` (例: `brytlog --json python run.py`)。
結果
生のログの代わりに、短いレポートが端末に出力されます。
───────────────

──────────
🧠📜 ブライトログクラッシュレポート
問題
整数を文字列で除算しようとすると、TypeError が発生してプログラムがクラッシュしました。構成の「items」値は文字列であるため、算術演算のために整数に変換する必要があります。
修正
/var/folders/ys/zfg72rwd661dn0cnrsqth5sh0000gn/T/tmph1e0wuse.py の 6 行目で、payload['items'] を整数に変換してから除算します。
平均 = ペイロード['合計'] / int(ペイロード['アイテム'])
完全なレポート → /path/to/project/brytlog-reports/2026-06-15T14-11-51.log
完全な生ログ → /path/to/project/brytlog-raw/2026-06-15T14-11-51.txt
─────────────────────
CLI フラグ
旗
説明
--バージョン
プログラムのバージョン番号を表示して終了します
--config
設定ファイルを開きます
--リセット
設定をデフォルトにリセットする
--アップグレード
PyPI で brytlog を最新バージョンにアップグレードする
--テスト
シミュレートされたクラッシュを実行して、LLM 構成をテストします。
--ログ
ローカルの brytlog-reports/ ディレクトリから最近のログを一覧表示します
--プロバイダー
LLM プロバイダー ( google 、 openai 、 anthropic 、 grok 、 ollama 、custom )
--モデル
使用する LLM モデル (例: gpt-4o-mini )
--APIキー
APIキーをインラインで渡す
--api-base-url
カスタムまたはollamプロバイダーのベースURL
--json
レポートをJSONとして出力します
--ログなし
ローカルの brytlog-reports/ ディレクトリへの AI レポートの書き込みを無効にする
--生のログなし
ローカルの brytlog-raw/ ディレクトリへの生ログの書き込みを無効にする
--静か
ターミナルでの生ログのライブ ストリームを抑制します (デフォルト)
--いいえ、静かに
ターミナルに生のログのライブ ストリームを表示します (静かなデフォルトをオーバーライドします)
環境変数
変数
デフォルト
説明
BRYTLOG_API_KEY
—
クラッシュ リポジトリに必要

rts (ローカル llm を使用しない場合)
BRYTLOG_PROVIDER
—
LLMプロバイダー
BRYTLOG_MODEL
—
LLMモデル
BRYTLOG_API_BASE_URL
—
カスタムプロバイダーに必須
BRYTLOG_SAVE_REPORT
本当の
AI レポート ファイルを無効にするには false に設定します。
BRYTLOG_SAVE_RAW_LOG
本当の
raw ログ ファイルを無効にするには false に設定します。
BRYTLOG_QUIET
本当の
true に設定すると、ライブ端末出力がミュートされます。
BRYTLOG_MAX_INPUT
4000
保持する端末出力の最大トークン (概算)
BRYTLOG_SYSTEM_PROMPT
(内蔵)
クラッシュレポート用のカスタムシステムプロンプト
BRYTLOG_TEMPERATURE
0.2
モデル温度
BRYTLOG_MAX_OUTPUT
1000
レポートの最大トークン数。注: API トークンの制限は、推論モデルに対応するために自動的に埋め込まれます (最小 2048)。
仕組み
brytlog は、指定されたコマンドを子プロセスとして実行し、デフォルトでその stdout と stderr をターミナルから外し、代わりに実行の簡潔な AI 概要のみを出力します。
クリーンな成功 (exit 0 、警告のようなキーワードなし): 成功を報告し、LLM には何も送信しません。
警告付きの成功 (exit 0 、 warning 、 deprecated 、 Skipped などのキーワード): 実行の短い概要として、サンプリングされた先頭/警告/末尾の抜粋を選択した LLM に送信します。
クラッシュ (ゼロ以外の終了): トークンで区切られた出力の先頭と末尾の抜粋を LLM に送信し、簡潔な問題/修正レポートを出力します。
brytlog はプロセスにフックするのではなくプロセスをラップするため、実行中のプログラムを変更することなく、あらゆる言語またはランタイムで機能します。
フローチャート TD
A[コマンドを実行] --> B[出力をキャプチャしてコードを終了]
B --> C{成功}
C -- はい --> D{警告が検出されました}
C -- いいえ --> F[ビルドの抜粋]
D -- いいえ --> E[印刷成功]
D -- はい --> F
F --> G[LLM に送信]
G --> H[AIレポート印刷]
B -.-> I["生のログを保存する (オプション)"]
H -.-> J["AI レポートの保存 (オプション)"]
読み込み中
注意事項
CLI フラグは環境変数より優先されます。

次に、設定ファイルよりも優先されます。
構成は ~/.brytlog.json (Windows では C:\Users\Name\.brytlog.json) にグローバルに保存されます。このファイルを直接編集することも、brytlog --config を実行して開くこともできます。
ログは、現在の作業ディレクトリの brytlog-reports/ および brytlog-raw/ にローカルに保存されます。 brytlog-*/ を .gitignore に追加するとよいでしょう。
明確にしておきます: トークンを節約し、コンテキストの肥大化を防ぐために、生のコマンド出力はデフォルトで人間と AI エージェントに対してミュート (Quiet モードで実行) されます。 Quiet モードでは、stdin は /dev/null にリダイレクトされます。つまり、対話型プロンプトはプログラムを即座にクラッシュさせ (例: EOFError )、無期限にハングするのではなく、有用な AI レポートを生成します。 Brytlog は、コマンドが失敗した場合にのみ、簡潔な AI レポートを出力します。通常どおりターミナルに生ログのライブ ストリームを表示し、対話型プロンプトに応答するには、 --no-quiet を使用します。
推論モデル (Gemini Flash-3-Preview など): これらのモデルは、出力予算の大部分を「思考」トークンに費やします。 Brytlog は、API レベルで 2048 トークンの下限を自動的に適用して、これらのモデルに考慮の余地を確保しながら、表示されるレポートを MAX_OUTPUT 制限以下に保つように指示します。
AGENTS.md や直接プロンプトに依存するのではなく、エージェントによる brytlog の使用を強制することも可能です (例: PATH 上のサブプロセス shim 経由)。ただし、このバージョンでは範囲外です。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
PyPI での最初の公開リリース
最新の
2026 年 6 月 24 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI logger. Contribute to Guy-Sela/brytlog development by creating an account on GitHub.

GitHub - Guy-Sela/brytlog: AI logger · GitHub
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
Guy-Sela
/
brytlog
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github/ workflows .github/ workflows src/ brytlog src/ brytlog tests tests .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md demo-no-brytlog.mp4 demo-no-brytlog.mp4 demo-with-brytlog.mp4 demo-with-brytlog.mp4 demo.py demo.py llms.txt llms.txt pyproject.toml pyproject.toml ruff.toml ruff.toml View all files Repository files navigation
Brytlog replaces raw logs with AI summary, thus saving developers time, trouble and money.
In agentic workflows brytlog acts as a cheap, fast pre-processor to the chief agent.
For example, Claude Opus 4.8 (chief agent) might run brytlog python run.py , rather than the plain python run.py . This way, instead of having to process the entire raw output on its own (slow, expensive, bloats context), it will only get a concise summary, generated by a cheaper, faster model, such as Gemini-3-flash.
As a fail-safe, raw logs are saved so they can be accessed by the chief agent or developer if still needed (this is a toggleable feature in config).
Even the cheaper model doesn't get the full raw dump, just the important parts, thus saving even more time and money.
In non-agentic, dev-driven workflows, brytlog simply saves the developer the time and trouble of analyzing raw output by himself, or copy-pasting lines into a coding assistant.
platform, language and llm vendor agnostic
minimal setup (just bring your own key, or run locally)
no need to change existing code (just add a couple of lines to AGENTS.md)
lightweight (~50 KB, ~1,400 lines of code)
privacy-minded (brytlog doesn't collect any data, and it redacts sensitive information before passing it to the LLM)
demo-no-brytlog.mp4
demo-with-brytlog.mp4
demo-json-output.mp4
How to use
After installation either run a command using brytlog (e.g. brytlog node main.js ), which will launch an on-boarding config process in the terminal, or run brytlog --config to open a json configuration file in your default editor.
API key (e.g. JQ.Ab9RN6W6QW7cmcnY92DIuoVtjCpKm_qfmO5T5oGzQmnwe5fjhw)
Google Gemini: Use keys from Google AI Studio ( https://aistudio.google.com/ ). Google Cloud Vertex AI is not natively supported yet.
Custom Providers: Supports OpenAI-compatible endpoints (Ollama, vLLM, etc.) via standard Authorization: Bearer headers. Azure OpenAI is not natively supported yet.
Additional required field for custom model only (Optional override for Ollama/Enterprise endpoints)
Base URL (e.g. http://localhost:11434/v1 )
Simply prefix any command with brytlog .
Syntax: brytlog [options] <command> .
brytlog python run.py
brytlog --api-key xyz... node build.js
brytlog --model claude-haiku-4-5 ./deploy.sh
Chaining Multiple Commands
Wrap the entire string in quotes to run multiple commands together.
brytlog " pip install -r requirements.txt && python run.py "
brytlog " node main.js
npm run serve "
Use in agentic workflow
Either prompt inline at the beginning of a session or add this or similar to AGENTS.md:
brytlog replaces raw terminal output with a concise AI summary.
Use it for interpreters (e.g.`python`, `node`), compilers, build tools (e.g. `npm`, `make`), and test runners (e.g. `pytest`). Do not use it for standard OS utilities (e.g.`ls`, `cat`), version control (e.g.`git`), or interactive CLI tools (e.g.`htop`).
Syntax: `brytlog [options] <command>` (e.g., `brytlog --json python run.py`).
Outcome
Instead of raw log, a short report is outputted to the terminal.
────────────────────────────────────────────────────────────
🧠📜 brytlog crash report
Problem
The program crashed due to a TypeError when attempting to divide an integer by a string. The 'items' value from the configuration is a string and needs to be converted to an integer for arithmetic operations.
Fix
Convert payload['items'] to an integer before division in /var/folders/ys/zfg72rwd661dn0cnrsqth5sh0000gn/T/tmph1e0wuse.py, line 6:
average = payload['total'] / int(payload['items'])
Full report → /path/to/project/brytlog-reports/2026-06-15T14-11-51.log
Full raw log → /path/to/project/brytlog-raw/2026-06-15T14-11-51.txt
────────────────────────────────────────────────────────────
CLI Flags
Flag
Description
--version
Show program's version number and exit
--config
Open the configuration file
--reset
Reset configuration to defaults
--upgrade
Upgrade brytlog to the latest version on PyPI
--test
Run a simulated crash to test the LLM configuration
--logs
List recent logs from the local brytlog-reports/ directory
--provider
LLM provider ( google , openai , anthropic , grok , ollama , custom )
--model
LLM model to use (e.g., gpt-4o-mini )
--api-key
Pass API key inline
--api-base-url
Base URL for custom or ollama providers
--json
Output the report as JSON
--no-log
Disable writing AI reports to the local brytlog-reports/ directory
--no-raw-log
Disable writing raw logs to the local brytlog-raw/ directory
--quiet
Suppress the live stream of raw logs in the terminal (default)
--no-quiet
Display the live stream of raw logs in the terminal (override quiet default)
Environment Variables
Variable
Default
Description
BRYTLOG_API_KEY
—
Required for crash reports (unless using local llm)
BRYTLOG_PROVIDER
—
LLM provider
BRYTLOG_MODEL
—
LLM model
BRYTLOG_API_BASE_URL
—
Required for custom provider
BRYTLOG_SAVE_REPORT
true
Set to false to disable AI report files
BRYTLOG_SAVE_RAW_LOG
true
Set to false to disable raw log files
BRYTLOG_QUIET
true
Set to true to mute live terminal output
BRYTLOG_MAX_INPUT
4000
Max tokens (approximate) of terminal output to keep
BRYTLOG_SYSTEM_PROMPT
(built-in)
Custom system prompt for the crash report
BRYTLOG_TEMPERATURE
0.2
Model temperature
BRYTLOG_MAX_OUTPUT
1000
Max tokens for the report. Note: API token limits are automatically padded (min 2048) to accommodate reasoning models.
How It Works
brytlog runs a given command as a child process, diverts its stdout and stderr away from the terminal by default, and instead only outputs a concise AI summary of the run.
Clean success (exit 0 , no warning-like keywords): reports success and sends nothing to the LLM.
Success with warnings (exit 0 , keywords such as warning , deprecated , or skipped detected): sends a sampled head / warnings / tail excerpt to the chosen LLM for a short summary of the run.
Crash (non-zero exit): sends a token-bounded head-and-tail excerpt of the output to the LLM and prints a concise problem/fix report.
Because brytlog wraps the process rather than hooking into it, it works for any language or runtime with no changes to the program being run.
flowchart TD
A[Run command] --> B[Capture output and exit code]
B --> C{Success}
C -- Yes --> D{Warnings detected}
C -- No --> F[Build excerpt]
D -- No --> E[Print success]
D -- Yes --> F
F --> G[Send to LLM]
G --> H[Print AI report]
B -.-> I["Save raw log (optional)"]
H -.-> J["Save AI report (optional)"]
Loading
Notes
CLI flags take precedence over environment variables, which in turn take precedence over the config file.
Configuration is saved globally to ~/.brytlog.json ( C:\Users\Name\.brytlog.json on Windows). You can edit this file directly or run brytlog --config to open it.
Logs are saved locally to brytlog-reports/ and brytlog-raw/ in the current working directory. You may want to add brytlog-*/ to your .gitignore .
To be clear: raw command output is muted (run in quiet mode) to humans and AI agents by default to save tokens and prevent context bloat. In quiet mode, stdin is redirected to /dev/null , meaning any interactive prompt will instantly crash the program (e.g. EOFError ) to generate a helpful AI report rather than hanging indefinitely. Brytlog only prints the concise AI report if a command fails. To display the live stream of raw logs in the terminal as usual and answer interactive prompts, use --no-quiet .
Reasoning models (e.g., Gemini Flash-3-Preview): These models spend a large portion of their output budget on "thinking" tokens. Brytlog automatically enforces a 2048-token floor at the API level to ensure these models have room to think, while still instructing them to keep the visible report under your MAX_OUTPUT limit.
Enforcing agent use of brytlog, rather than relying on AGENTS.md or direct prompting, is also possible (e.g. via subprocess shim on PATH), but out of scope in this version.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Initial public release on PyPI
Latest
Jun 24, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
