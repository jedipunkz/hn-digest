---
source: "https://github.com/shaharia-lab/agento"
hn_url: "https://news.ycombinator.com/item?id=49283885"
title: "Agento: The missing dashboard for Claude Code"
article_title: "GitHub - shaharia-lab/agento: Production ready Personal AI Agent Platform using Claude Code CLI. Focused on productivity, reliability and security at it's core · GitHub"
author: "shahariaa"
captured_at: "2026-08-13T10:53:51Z"
capture_tool: "hn-digest"
hn_id: 49283885
score: 1
comments: 0
posted_at: "2026-08-13T10:23:05Z"
tags:
  - hacker-news
  - translated
---

# Agento: The missing dashboard for Claude Code

- HN: [49283885](https://news.ycombinator.com/item?id=49283885)
- Source: [github.com](https://github.com/shaharia-lab/agento)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T10:23:05Z

## Translation

タイトル: Agento: クロード コードの欠落したダッシュボード
記事のタイトル: GitHub - shaharia-lab/agento: Claude Code CLI を使用した本番環境対応パーソナル AI エージェント プラットフォーム。生産性、信頼性、セキュリティを中核に重点を置く · GitHub
説明: Claude Code CLI を使用した、実稼働対応のパーソナル AI エージェント プラットフォーム。生産性、信頼性、セキュリティを中核に重点を置く - shaharia-lab/agento

記事本文:
GitHub - shaharia-lab/agento: Claude Code CLI を使用した本番環境対応のパーソナル AI エージェント プラットフォーム。生産性、信頼性、セキュリティを中核に重点を置く · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
シャハリアラボ
/
代理人
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
155 コミット 155 コミット .claude .claude .github .github エージェント エージェント

ts cmd cmd docs docs e2e e2e フロントエンド フロントエンド 内部 内部 .env.example .env.example .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .mockery.yaml .mockery.yaml .pre-commit-config.yaml .pre-commit-config.yaml CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md 資産.go 資産.go 資産_dev.go 資産_dev.go go.mod go.mod go.sum go.sum main.go main.go sonar-project.properties sonar-project.properties すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code の欠落しているダッシュボード。
Claude Code は終了した瞬間にすべてを忘れます。 Agento は、すでにディスクに書き込まれているセッション ファイルを読み取り、それらをコスト分析、生産性に関する洞察、およびすべての実行の検索可能な履歴に変換します。また、エージェントの構築、スケジュール設定、使用するツールへの接続のためのブラウザ UI も提供します。
One Go バイナリ。 API キー、アカウント、テレメトリはありません。あなたの履歴、エージェント、分析はあなたのマシン上に残ります。
醸造インストール シャハリアラボ/タップ/エージェント
エージェントウェブ
それが全体のセットアップです。 Agento は http://localhost:8990 で開き、クロード コードの履歴を見つけて、ダッシュボードの構築を開始します。
Agento によってお金や時間を節約できるのであれば、プロジェクトに対してできる最も役立つことはスターです。 1 秒かかりますが、他のクロード コード ユーザーがそれを見つける方法です。
要件: Claude Code CLI がインストールされ、認証されている。 claude がターミナルで実行されている場合、Agento は機能します。 Agento は Claude Code がすでに持っている認証を使用するため、Anthropic API キーは必要ありません。
醸造インストール シャハリアラボ/タップ/エージェント
エージェントウェブ
直接ダウンロード
# リリースからプラットフォームのアーカイブを取得します
tar -xzf Agento_Linux_x86_64.tar.gz
sudo mv エージェント /usr/local/bin/
エージェントウェブ
ビン

Linux (x86_64、arm64)、macOS (Intel、Apple Silicon)、および Windows 用のaries は、「リリース」ページにあります。
便利なフラグ: ポートを変更する場合はagento web --port 3000、タブを開くのをスキップする場合は--no-browser。再起動後もバックグラウンドで実行し続けるには:
Agento サービスのインストール # その後: ステータス |停止 |開始 |再起動 |ログ |アンインストールする
得られるもの
すべてのトークンの種類、適切な価格設定
入力、出力、キャッシュ読み取り、キャッシュ書き込みの料金は大きく異なるため、Agento では合計に 1 つの価格を掛けるのではなく、それらを区別しています。その結果、ほとんどの人が驚くべきグラフが得られます。トークンが最も多いモデルは、多くの場合、お金を奪うモデルではありません。
コストは、サブエージェント内で行われる作業も含めて、コストを費やしたモデルに起因するため、より安価なモデルに委任すると、実際の節約として現れます。
効果が上がっているかどうかを確認する
Insights ページは生のカウントを超えています。セッションに必要なターン数、クロードが何かを尋ねるまでにどれだけ進んだか、待機させた時間、キャッシュ ヒット率とツール エラー率をすべて前の期間と比較して追跡するため、方向性がわかります。次に、すべてのツール呼び出しがスキル、プラグイン、MCP サーバー、またはサブエージェントの責任であると判断します。これにより、呼び出しの 3 分の 1 を静かに実行しているスキルを見つけることができます。
期間は壁時計ではなく、アクティブな時間を意味します。クロード コードのセッションは再開可能であるため、1 週間後にセッションを開始すると、1 週間の作業が報告されます。制御するしきい値を超えるアイドル ギャップは、期間が表示されるすべての場所から除外されます。
自分の作業パターンを理解する
1 日あたりのセッション数、モデルの組み合わせ、最も混雑する日、プロジェクトあたりのコスト、およびセッションが終了した時間だけでなく、実行中の時間ごとにセッションをカウントするアクティビティ ヒートマップ。
これまでに実行したすべてのセッションを参照して検索します
フィルタリングおよびページングされた i

n SQL なので、セッションが 50 であっても 5,000 であっても高速に動作します。タイトルとコンテンツを検索し、プロジェクト、モデル、日付、コスト、期間でフィルターし、リンクされたプル リクエスト、Git ブランチ、権限モードを各行で確認します。
セッションをステップバイステップで再生します
ジャーニー ビューは、実行の完全なタイムラインを再構築します。つまり、すべてのプロンプト、応答、ツール呼び出し、結果が順番に再構築され、各サブエージェントのステップは、それを生成した委任の下にネストされます。長時間の自動運転で問題が発生した場合、ここで問題が発生した場所がわかります。
各セッションには、ターン、ターンあたりのステップ数、最長自律チェーン、アクティブ期間、クロードが作業に費やした時間、ユーザー自身の平均応答時間、ツール エラー率などの独自のメトリクスも含まれます。
コードを書かずにエージェントを構築
エージェントに名前、システム プロンプト、モデル、思考モード、および使用できるツールの明示的なリストを与えます。一度保存すると、ブラウザ、スケジューラ、または CLI から再利用できます。 {{current_date}} のようなテンプレート変数は実行時に入力されます。
ここでは権限モードが重要です。エージェントは、最初に計画を立て、ユーザーの承認後にのみ動作するように設定できます。これは、ファイルの書き込みやコマンドの実行が可能になった時点で必要なことです。
cron 式、固定間隔、または特定の時間に 1 回、任意のエージェントを実行します。すべての実行はステータス、継続時間、完全な出力とともに記録されるため、外出中に何が実行されたかを正確に確認できます。
すでに使用しているツールを接続する
各統合はインプロセス MCP サーバーとして実行されるため、動作する追加のデーモンはありません。一度設定すれば、どのエージェントでも使用できるようになります。
Google (カレンダー、Gmail、ドライブ)、GitHub、Slack、Jira、Confluence、Telegram、WhatsApp が組み込まれています。他の MCP サーバーは、stdio、ストリーミング可能な HTTP、または SSE 経由で ~/.agento/mcps.yaml を介して追加できます。
価格カタログを正直に保つ
Anthropic、Moonshot、Z.ai、Alibab の料金が配送されます

料金を追加すると、過去の使用料金は請求された価格のままになりますが、料金を修正すると書き換えられます。レートが公開されていないモデルは、他のものとして静かに価格設定されるのではなく、不明として報告されます。
Agento は ~/.claude を読み取り、結果を ~/.agento/agento.db のローカル SQLite データベースにキャッシュします。何もアップロードされず、アカウントもサーバー コンポーネントもありません。数字から除外したいプロジェクトはすべてのレポートで非表示にすることができ、期間メトリクスの背後にあるアイドルしきい値は自分で設定できます。
💬 チャットとタブ付きマルチチャット ワークスペース
構築したエージェントと複数ターンの会話を行います。応答はサーバー送信イベントを通じてライブでストリーミングされ、セッションはローカルに保持され、お気に入りにしたり名前を変更したりできます。ファイルをドラッグ アンド ドロップするか、画像を入力に直接貼り付けます。
マルチチャット ワークスペースは複数の会話を並行して実行し、各タブには独自のエージェントとセッション状態があり、ページのリロード後も存続します。
デモ.webm
マルチタブチャット-デモ.webm
💻 CLI: ターミナルからエージェントを実行
エージェントは「今日のリポジトリで何が変更されましたか?」と尋ねます。
Agento ask --agent code-reviewer " ステージングされた差分を確認する "
Agento ask --agent code-reviewer " フォローアップ " < セッション ID >
会話を続けるにはセッション ID を渡します。スクリプトとシェル パイプラインに役立ちます。
すべての HTTP リクエスト、エージェントの実行、ツール呼び出し、ストレージ操作が計測されます。 [監視設定] タブから OTLP gRPC エクスポーターまたは Prometheus プル エンドポイントを構成すると、ホットリロードが行われ、再起動や構成ファイルは必要ありません。構造化ログは、セッションごとのログとともに ~/.agento/logs/system.log に書き込まれます。
タスク完了およびエージェント イベントの SMTP 配信を構成し、UI からテスト メッセージを送信し、通知ログを参照します。スケジュールされた実行はすべてジョブ h に保持されます

開始時刻、期間、終了ステータス、および完全な出力を含むストーリー。
複数の名前付きクロード設定プロファイル ( ~/.claude/settings_<slug>.json として保存) を保持し、エージェントごとまたはチャットごとに切り替えます。デフォルトのプロファイルは、初回起動時に既存の ~/.claude/settings.json から作成されます。暗いテーマと明るいテーマ、フォント サイズ、フォント ファミリーは、UI 全体に即座に適用されます。
Agento は起動時に新しいリリースをチェックし、利用可能なリリースがある場合はバナーを表示します。適切な場所にアップグレードするには、agento update を実行します。バックグラウンド サービスとしてインストールされている場合、サービスは自動的に再起動されます。
何も設定する必要はありません。以下はすべてオプションであり、環境変数が設定 UI よりも優先されます。
Agento Web [--port int] [--no-browser] Web UI を開始します
Agento ask [--agent slug] [--no- Thinking] エージェントに 1 回限りの質問をする
<質問> [セッションID]
Agento update [-y] [--no-restart] 最新リリースへのアップデート
エージェント サービス <インストール|アンインストール|開始|停止|再起動|ステータス|ログ>
Agento サービスは、macOS に LaunchAgent ( ~/Library/LaunchAgents/com.shaharialab.agento.plist ) をインストールするか、Linux に systemd ユーザー ユニット ( ~/.config/systemd/user/agento.service ) をインストールするため、Agento はログアウトして再起動しても存続します。
Go 1.25 以降と Node.js が必要です。
git clone https://github.com/shaharia-lab/agento.git
CDエージェント
ビルドする
アーキテクチャの概要と開発ワークフローについては、docs/development.md を参照してください。
別のデバイスから Agento にアクセスする
Agento はデフォルトではループバックでのみリッスンし、認証はありません。作業中のマシン上で実行することを目的としています。 API はエージェントを作成して実行できるため、API に到達できるものはすべてそのマシン上でコマンドを実行できます。
携帯電話、タブレット、または別のコンピュータから使用するには:
AGENTO_BIND=0.0.0.0 エージェント Web
信頼できるネットワーク上でのみこれを行うか、認証を行うプロキシを前に置きます

それの。 IP ではなくホスト名で (リバース プロキシまたはトンネル経由で) Agento にアクセスする場合は、[設定] のパブリック URL (または AGENTO_PUBLIC_URL ) をそのアドレスに設定しないと、リクエストは拒否されます。
アップグレード中ですか？これは、すべてのインターフェイスでリッスンしていました。別のデバイスから Agento にアクセスして機能しなくなった場合は、 AGENTO_BIND=0.0.0.0 を設定します。起動ログには、バインドされたインターフェイスの名前が記録されます。
はじめに : セットアップと最初の実行のウォークスルー
クロード セッション: スキャンされる内容、コストと期間の測定方法、およびその上に構築される分析
エージェント : システム プロンプト、モデル、ツール、テンプレート変数
タスク: スケジュールに従ってエージェントを実行し、ジョブ履歴を実行します。
統合 : Google、GitHub、Slack、Jira、Confluence、Telegram、WhatsApp の接続
価格設定: コストの計算方法とカタログの管理方法
セキュリティ : ネットワークへの露出、API ガード、データの保存場所
モニタリング : OpenTelemetry トレース、メトリクス、ログ
開発: アーキテクチャと貢献ガイドライン
問題やプルリクエストは大歓迎です。機能がありませんか?問題を開いて、それを何に使用するかを教えてください。
マサチューセッツ工科大学。 Shaharia Lab によって維持されています。
Agento は無料でオープンソースです。
役立つ場合は、リポジトリにスターを付けてください。

[切り捨てられた]

## Original Extract

Production ready Personal AI Agent Platform using Claude Code CLI. Focused on productivity, reliability and security at it's core - shaharia-lab/agento

GitHub - shaharia-lab/agento: Production ready Personal AI Agent Platform using Claude Code CLI. Focused on productivity, reliability and security at it's core · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
shaharia-lab
/
agento
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
155 Commits 155 Commits .claude .claude .github .github agents agents cmd cmd docs docs e2e e2e frontend frontend internal internal .env.example .env.example .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .mockery.yaml .mockery.yaml .pre-commit-config.yaml .pre-commit-config.yaml CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md assets.go assets.go assets_dev.go assets_dev.go go.mod go.mod go.sum go.sum main.go main.go sonar-project.properties sonar-project.properties View all files Repository files navigation
The missing dashboard for Claude Code.
Claude Code forgets everything the moment it exits. Agento reads the session files it already writes to your disk and turns them into cost analytics, productivity insights and a searchable history of every run. It also gives you a browser UI for building agents, scheduling them, and connecting them to the tools you use.
One Go binary. No API key, no account, no telemetry. Your history, your agents and your analytics stay on your machine.
brew install shaharia-lab/tap/agento
agento web
That is the whole setup. Agento opens at http://localhost:8990 , finds your Claude Code history, and starts building your dashboards.
If Agento saves you money or time, a star is the single most useful thing you can do for the project. It takes a second and it is how other Claude Code users find it.
Requirements: the Claude Code CLI , installed and authenticated. If claude runs in your terminal, Agento works. No Anthropic API key is needed, because Agento uses the authentication Claude Code already has.
brew install shaharia-lab/tap/agento
agento web
Direct download
# grab the archive for your platform from Releases
tar -xzf agento_Linux_x86_64.tar.gz
sudo mv agento /usr/local/bin/
agento web
Binaries for Linux (x86_64, arm64), macOS (Intel, Apple Silicon) and Windows are on the Releases page .
Useful flags: agento web --port 3000 to change the port, --no-browser to skip opening a tab. To keep it running in the background across reboots:
agento service install # then: status | stop | start | restart | logs | uninstall
What you get
Every token type, priced properly
Input, output, cache reads and cache writes bill at very different rates, so Agento keeps them apart instead of multiplying one total by one price. The result is the chart most people find surprising: the model with the most tokens is often not the model taking your money.
Cost is attributed to the model that spent it, including work done inside sub-agents, so delegating to a cheaper model shows up as an actual saving.
Find out whether you are getting more effective
The Insights page goes past raw counts. It tracks how many turns a session needed, how far Claude got before it had to ask you something, how long you kept it waiting, cache hit rate and tool error rate, all compared against the previous period so you can see the direction. It then attributes every tool call to the skill, plugin, MCP server or sub-agent responsible, which is how you find the skill quietly burning a third of your calls.
Durations mean active time, not wall clock. Claude Code sessions are resumable, so one picked up a week later would otherwise report a week of work. Idle gaps beyond a threshold you control are excluded everywhere a duration is shown.
Understand your working patterns
Sessions per day, model mix, busiest days, cost per project, and an activity heatmap that counts a session in every hour it was running rather than only the hour it finished.
Browse and search every session you have ever run
Filtered and paged in SQL, so it stays fast whether you have 50 sessions or 5,000. Search across titles and content, filter by project, model, date, cost or duration, and see linked pull requests, git branch and permission mode on every row.
Replay any session step by step
The journey view reconstructs the full timeline of a run: every prompt, response, tool call and result in order, with each sub-agent's steps nested under the delegation that spawned it. When a long autonomous run goes wrong, this is where you find out where.
Each session also carries its own metrics: turns, steps per turn, longest autonomous chain, active duration, time Claude spent working, your own average reply time, and tool error rate.
Build agents without writing code
Give an agent a name, a system prompt, a model, a thinking mode and an explicit list of tools it may use. Save it once and reuse it from the browser, the scheduler or the CLI. Template variables like {{current_date}} are filled in at runtime.
Permission modes matter here: an agent can be set to plan first and act only after you approve, which is what you want the moment it can write files or run commands.
Run any agent on a cron expression, a fixed interval, or once at a specific time. Every execution is recorded with its status, duration and full output, so you can see exactly what ran while you were away.
Connect the tools you already use
Each integration runs as an in-process MCP server, so there is no extra daemon to operate. Configure it once and any agent can use it.
Google (Calendar, Gmail, Drive), GitHub, Slack, Jira, Confluence, Telegram and WhatsApp are built in. Any other MCP server can be added through ~/.agento/mcps.yaml over stdio, streamable HTTP or SSE.
Keep the pricing catalog honest
Rates ship for Anthropic, Moonshot, Z.ai and Alibaba models, and they are effective-dated: adding a rate leaves past usage priced at what it was charged, while correcting one rewrites it. A model with no published rate is reported as unknown instead of being quietly priced as something else.
Agento reads ~/.claude and caches results in a local SQLite database at ~/.agento/agento.db . Nothing is uploaded, there is no account, and there is no server component. Projects you would rather leave out of the numbers can be hidden from every report, and the idle threshold behind the duration metrics is yours to set.
💬 Chats and a tabbed multi-chat workspace
Hold multi-turn conversations with any agent you have built. Responses stream live over Server-Sent Events, sessions persist locally, and you can favourite or rename them. Drag and drop files or paste images straight into the input.
The multi-chat workspace runs several conversations in parallel, each tab with its own agent and session state, and it survives a page reload.
demo.webm
Multi-Tab.Chat.-.Demo.webm
💻 CLI: run agents from the terminal
agento ask " What changed in the repo today? "
agento ask --agent code-reviewer " Review the staged diff "
agento ask --agent code-reviewer " Follow up " < session-id >
Pass a session ID to continue a conversation. Useful for scripts and shell pipelines.
Every HTTP request, agent run, tool call and storage operation is instrumented. Configure an OTLP gRPC exporter or a Prometheus pull endpoint from the Monitoring settings tab and it hot-reloads, no restart and no config file. Structured logs are written to ~/.agento/logs/system.log , with per-session logs beside them.
Configure SMTP delivery for task completion and agent events, send a test message from the UI, and browse the notification log. Every scheduled run is kept in job history with its start time, duration, exit status and full output.
Keep several named Claude settings profiles (stored as ~/.claude/settings_<slug>.json ) and switch between them per agent or per chat. A default profile is created from your existing ~/.claude/settings.json on first launch. Dark and light themes, font size and font family apply instantly across the UI.
Agento checks for new releases on startup and shows a banner when one is available. Run agento update to upgrade in place. If it is installed as a background service, the service is restarted for you.
Nothing needs configuring. Everything below is optional, and environment variables win over the Settings UI.
agento web [--port int] [--no-browser] Start the web UI
agento ask [--agent slug] [--no-thinking] Ask an agent a one-off question
<question> [session-id]
agento update [-y] [--no-restart] Update to the latest release
agento service <install|uninstall|start|stop|restart|status|logs>
agento service installs a LaunchAgent on macOS ( ~/Library/LaunchAgents/com.shaharialab.agento.plist ) or a systemd user unit on Linux ( ~/.config/systemd/user/agento.service ), so Agento survives logout and reboot.
Requires Go 1.25+ and Node.js.
git clone https://github.com/shaharia-lab/agento.git
cd agento
make build
See docs/development.md for the architecture overview and the development workflow.
Reaching Agento from another device
Agento listens on loopback only by default, and has no authentication — it is meant to run on the machine you are working at. The API can create an agent and run it, so anything that can reach it can run commands on that machine.
To use it from a phone, tablet or another computer:
AGENTO_BIND=0.0.0.0 agento web
Only do that on a network you trust, or put a proxy that authenticates in front of it. If you reach Agento under a hostname rather than an IP — through a reverse proxy or a tunnel — set Public URL in Settings (or AGENTO_PUBLIC_URL ) to that address, or requests will be refused.
Upgrading? This used to listen on every interface. If you reach Agento from another device and it stopped working, set AGENTO_BIND=0.0.0.0 . The startup log names the interface it bound.
Getting started : setup and a first-run walkthrough
Claude sessions : what is scanned, how cost and duration are measured, and the analytics built on top
Agents : system prompts, models, tools and template variables
Tasks : running agents on a schedule, and job history
Integrations : connecting Google, GitHub, Slack, Jira, Confluence, Telegram and WhatsApp
Pricing : how cost is calculated and how to maintain the catalog
Security : network exposure, the API guards, and where your data lives
Monitoring : OpenTelemetry traces, metrics and logs
Development : architecture and contribution guidelines
Issues and pull requests are welcome. Missing a feature? Open an issue and tell us what you would use it for.
MIT. Maintained by Shaharia Lab .
Agento is free and open source.
If it is useful to you, star the repository so more Claude

[truncated]
