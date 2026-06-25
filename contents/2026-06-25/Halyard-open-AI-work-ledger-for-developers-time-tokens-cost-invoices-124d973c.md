---
source: "https://github.com/Kormiloio/Halyard"
hn_url: "https://news.ycombinator.com/item?id=48667366"
title: "Halyard – open AI work ledger for developers (time, tokens, cost, invoices)"
article_title: "GitHub - Kormiloio/Halyard · GitHub"
author: "mcamaj"
captured_at: "2026-06-25T01:40:46Z"
capture_tool: "hn-digest"
hn_id: 48667366
score: 1
comments: 0
posted_at: "2026-06-25T00:47:03Z"
tags:
  - hacker-news
  - translated
---

# Halyard – open AI work ledger for developers (time, tokens, cost, invoices)

- HN: [48667366](https://news.ycombinator.com/item?id=48667366)
- Source: [github.com](https://github.com/Kormiloio/Halyard)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T00:47:03Z

## Translation

タイトル: Halyard – 開発者向けのオープン AI 作業台帳 (時間、トークン、コスト、請求書)
記事タイトル: GitHub - Kormiloio/Halyard · GitHub
説明: GitHub でアカウントを作成して、Kormiloio/Halyard の開発に貢献します。

記事本文:
GitHub - コルミロイオ/ハリヤード · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
コルミロイオ
/
ハリヤード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインBr

anches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
328 コミット 328 コミット .agents/ スキル .agents/ スキル .github .github 資産/ アイコン/ ハリヤード-航海 アセット/ アイコン/ ハリヤード-航海 ドキュメント ドキュメント openspec openspec 価格 価格設定 プロンプト プロンプト サンプル サンプル src/ ハリヤード src/ ハリヤード テスト テスト vscode-extension vscode-extension .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ハリヤードとは帆を上げるための綱のことです。引っ張ると帆が上がります。これを引くと、AI の仕事に焦点が当てられます。
AI の仕事には痕跡が残ります。ハリヤードは、その証跡を読みやすく、監査可能で、クライアントにとって安全なものにします。
すべての AI セッション (時間、トークン、モデル、コスト、プロジェクト) は、作業が行われた場所でキャプチャされ、
あなたが所有するあなたのマシンにプレーンテキストとして保存されます。アカウントがありません。クラウドサービスはありません。いいえ
プロンプトまたはコードキャプチャ。これまで。 MITライセンス取得済み。
ステータス: アルファ版、オープンソース — キャプチャ ループ、レポート、請求書、The Bridge、および TUI
日常の使用において。
あなたは AI 支援の仕事をしています。スプリント、1 か月、またはクライアントとのエンゲージメントの終了時に、
次の 3 つの基本的な質問に答えることができません。
このプロジェクトで AI が実際に費やしたのはいくらですか?
AI は何を生み出すのに役立ちましたか?そしてそれを証明できますか?
AI への支出は正しい方向に進んでいますか?
あなたのツールはこれを記録しません。ハリヤードはそうです。
これは、Claude Code、Cursor、および Gemini CLI 内の軽量フックとして実行されます。
パブリック AI セッション フックがない VS Code などのツールのマニュアル/エディター タスク キャプチャ
存在します。すべてのセッションは、所有するプレーンテキストのログに 1 行を書き込みます。そのログから:
コストの内訳、プロジェクトの属性

使用料、請求書の証拠、そして最終的には署名された、
クライアントに渡すことができる検証可能な AI 作品の付録。
プライバシーの約束は無条件です: Halyard はセッション メタデータをキャプチャし、プロンプトを表示することはありません
コンテンツ、コードコンテキスト、ファイルコンテンツ、またはトランスクリプト。
個人の開発者とフリーランサーが、現時点での主な対象者です。
あなたの時間、AI 支出、請求書の証拠は、プレーンテキストとしてコンピュータに保存されます。
ラップトップ。 Halyard は、プロンプトやコードを公開することなく、何が起こったのかを証明するのに役立ちます。
Git を追加したり、バックアップしたり、好きなように同期したりできます。アカウントがありません。 SaaS はありません。いいえ
独自の形式。
小規模な AI ショップ — チーム全体で同じローカル台帳形式を共有します。
プロジェクト支出、信頼ラベル付きのコスト配分、クライアントに安全な付録の構築
個々の平文ログから。
エンタープライズ — 同じフォーマットでガバナンス、コストセンター、および
クロスツール AI Work Intelligence については後で説明します。その層は加算的であり、ゲートオンされています
design-partner のプルであり、ローカル ファイルの意味は変わりません。
コレクション — AI の作業が行われる場所で実行される軽量のフック。クロード・コード、
Cursor、Windsurf、および Gemini CLI フックはセッションを自動的にキャプチャします。 VSコード
拡張機能は編集時間、ブランチ、コードデルタ — トークン数をキャプチャします
VS Code/Copilot がパブリック セッション フックを公開するまでは利用できません。 v3.5以降、
クロード コード セッションには、アドバイザリー クライアント サーフェスのタグが付けられます
( cli/desktop/ide ) がローカル環境から検出されました。キャプチャされたフィールド
時間、トークン（利用可能な場合）、モデル、コスト、プロジェクト、ブランチが含まれます。書かれた
あなたが所有するプレーンテキストのログに。新しいセッションが追加され、現在の
Hardening Track は修正を明示的かつ監査可能にしています。何も失われません
静かに。
インテリジェンス — そのログに基づいて構築された分析。ローカル CLI レポート、プロジェクトごとのコストの内訳、モデルごとの支出、予算アラート、および信頼ラベル付き合計 (c

取得、計算、割り当て）。オフラインで動作し、アカウントは必要ありません。
AI Work Ledger — シートのサブスクリプションとクレジット プランのコスト割り当て。 Claude Max に月額 200 ドルを支払うと、Halyard はそのコストをアクティブな時間数、セッション数、またはクレジットの使用量に基づいてプロジェクト全体に比例配分するため、各クライアントとのエンゲージメントに実際にかかるコストがわかります。 ai-sessions.log および ai-plans.toml の上で実行されます。生のログには何も書き戻されません。
証拠となるアーティファクト — 今日の請求書の証拠と、署名された証明可能な AI 作品
次に付録。目標は、AI 支援による作業を証明する、クライアントに安全な成果物です。
プロンプト、トランスクリプト、ソースコード、またはファイルの内容は表示されません。
The Bridge — キャプチャの発生をリアルタイムで監視するためのローカル ダッシュボード。ハリヤード プロジェクト内でハリヤード ダッシュボードを実行します。
豊富なセッション テレメトリ — ツールが公開する場合、Halyard はコストを超えた運用メタデータ (ツールの呼び出し回数、エラー率、稼働時間とアクティブ エージェント時間の比較、コード デルタ、モデルごとの内訳など) を収集します。 Gemini CLI セッションには、履歴ファイルからの完全なマルチモデルの内訳が含まれています。これらのシグナルは、生産性スコアではなく、セッション形状の正直なシグナルとして、TUI と The Bridge に仕事の健康指標として現れます。
Honors — 生の時間ではなく、正確な証拠を表彰するサービス記録。ランクは属性セッション (デッキハンド → コモドール) で上がり、ストライプはウォッチの連続記録を追跡し、8 つのメダルは重要な行動を認識します (最初のウォッチの完了、マニフェストのクリーンな維持、漂流セッションの救出など)。ハリヤード・オナーズを実行して記録を確認してください。
Friends of the Sea — 完了したプロジェクトごとに 1 匹の海の生き物が、個性によって自動的に割り当てられます。プロジェクトは航海の段階を経て進みます（錨を寄せる → 前進する → マークを丸くする → フライングカラーズ → 船の形）

· セッションが蓄積されると係留されます。ターゲットがヒットするか非アクティブになるとオートコンプリートします。ハリヤード航海を実行して名簿を確認します。ブリッジの船長室にはパスポートが表示されます。使用した AI ツールごとに 1 つのスタンプが押されます。
Halyard は Python 3.11 以降のローカルファースト CLI およびダッシュボードであり、ホスト型課金ではありません
サービス。ハリヤード コマンドは Typer アプリ、レポートはリッチ、ターミナルを使用します
ダッシュボードはテキストを使用し、ブリッジは小規模な 127.0.0.1 HTTP です。
サーバー。耐久性のあるデータ モデルは、プロジェクト フォルダー内のプレーン テキストです。 SQLite は
クエリを高速化するための再構築可能な読み取りモデル キャッシュのみ。
キャプチャ パイプラインは意図的に単純になっています。
AIツールフック/インポーター/マニュアルコマンド
-> 正規化された AiSession オブジェクト
-> 追加に重点を置いた ai-sessions.log 行
-> レポート、元帳割り当て、ダッシュボード、請求書の証拠
では、ハリヤードは「丸太を見ているだけ」なのでしょうか？正確には違います。最良のパブリックを使用します
各ツールが公開するシグナル:
Claude Code : UserPromptSubmit フックと Stop フックをインストールします。スタートフック
セッションの開始時刻と git SHA を記録します。ストップフックは構造化された情報を読み取ります。
フックペイロード。新しいクロード コード形式では、トークン/モデルを集約することもできます
フックによって渡されたローカル トランスクリプト JSONL パスからのメタデータ。
Cursor : beforeSubmitPrompt をインストールし、フックを停止します。ストップを読み上げます
ペイロードであり、アトリビューションには workspace_roots を優先します。
実際のエディターのワークスペース、必ずしもシェル CWD ではありません。
Gemini CLI : SessionStart 、 AfterModel 、および AfterAgent フックをインストールします。
AfterModel は、 useMetadata からトークンの使用量を蓄積します。アフターエージェント
セッションを終了します。 OpenTelemetry (OTLP) と深く統合して測定します。
正確な API およびツールの実行時間 ( api_seconds 、 tools_seconds ) と強化
Gemini のローカル履歴ファイルから、正確なマルチモデル トークンの内訳、ツール呼び出し数、
そして決定的なコスト。
コーデックスデスク

op : ローカル ~/.codex/sessions/.../rollout-*.jsonl をインポートします
ファイルを作成し、タイミング/モデル/トークンのメタデータを抽出し、インポートされたセッション ID を記録します
そのため、インポートを繰り返してもエントリが重複することはありません。
VS Code / GitHub Copilot : ローカル VS Code 拡張機能 ( vscode-extension/ )
アクティブな編集時間を追跡し、git 経由でブランチとコードデルタをキャプチャし、書き込みます
ハリヤード記録セッションによるセッション。ローカルから拡張機能をインストールする
.vsix ビルド;パブリック Copilot セッション フックはまだ存在しないため、トークン数は
利用できません。
すべてのコレクターは、同じ正規化されたレコード形式 (タイムスタンプ、ツール、
モデル、トークン (利用可能な場合)、キャッシュ トークン、コスト、請求タイプ、プロジェクト、
ブランチ、キャプチャソース、および帰属の出所。ハリヤードは保管しない
プロンプト、ソースコード、ファイルの内容、または ai-sessions.log の完全なトランスクリプト。
コレクターがローカルのトランスクリプトまたは履歴ファイルを一時的に読み取る場合、
セッションメタデータの抽出のみを目的としています。
ログは追加に重点を置いています。セッションレコードは s ... 行です。修正は
元の行のハッシュをキーとする個別の修正レコード。作家
排他的な OS レベルのファイル ロックを保持します (POSIX では fcntl.flock、msvcrt.locking)
Windows の場合)、同時フックは書き込みをインターリーブしません。不正な形式のレコード
レポート生成をクラッシュさせる代わりに隔離されます。
コスト処理は信頼について明示的に行われます。直接 API の使用状況をキャプチャすることも、
トークンと現地の価格表から計算されます。座席プランまたはクレジットプランは、
ai-plans.toml からのレポート時に、アクティブな分数、セッション数、
またはクレジット。レポートでは、結果にキャプチャ、計算、割り当て、
推測、混合、または割り当てられていないため、クライアントに直面する証拠は、
推定値は測定値です。
プラットフォーム: macOS、Linux、Windows。ハリヤードサービスのインストール
コマンドは macOS のみです。他のプラットフォームでもハリヤード ダを実行できます

板
代わりに長寿命のターミナルを使用します。
pipx ハリヤードをインストールする
cd ~ /businesss/my-freelance
ハリヤード初期化
# ガイド付きセットアップでは、サポートされているフックがインストールされ、準備が整っているかどうかがチェックされます。
ハリヤードのセットアップ
# フックを確認して最初にキャプチャします。
ハリヤードドクター --最初の捕獲
# ダッシュボードを開く
ハリヤードダッシュボード
日々のワークフロー
# 人間タイマーを開始します (AI セッションはバックグラウンドで自動キャプチャされます)
halyard start acme/auth-migration
# ...仕事をしてください...
ハリヤードストップ
# ターミナルUI
ハリヤード・トゥイ
# 自然言語で質問する — 2 つの方法、どちらも読み取り専用:
ハリヤード ログ「今月は何を使いましたか?」 # CLI から (ローカルで
# デフォルト; --エージェント クロード|openai オプション)
# …または、MCP サーバー経由のコーディング エージェント経由 (以下を参照)
# AI 使用証拠の付録を含む請求書を生成する
ハリヤード請求書 acme --month 2026-05 --include-ai-evidence
完全なコマンド リファレンスを表示する
# 統計フォワード分析 - セッション、連続記録、ピーク時間、モデルの組み合わせ
ハリヤードの使用法 -- 射程 30d
ハリヤードの使用法 --range 7d --json # 機械可読
# CI + スクリプトのレポート/使用状況/予算/ステータス/証拠/健康/医師に関する --json
ハリヤードレポート --all --json | jq ' .totals.cost_usd '
ハリヤードの予算 --json | jq ' .[] | select(.month.state=="over") ' # ゲートの支払い
# バックグラウンドの launchd サービスをインストールします (macOS):
ハリヤード・サービ

[切り捨てられた]

## Original Extract

Contribute to Kormiloio/Halyard development by creating an account on GitHub.

GitHub - Kormiloio/Halyard · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Kormiloio
/
Halyard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
328 Commits 328 Commits .agents/ skills .agents/ skills .github .github assets/ icons/ halyard-nautical assets/ icons/ halyard-nautical docs docs openspec openspec pricing pricing prompts prompts samples samples src/ halyard src/ halyard tests tests vscode-extension vscode-extension .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A halyard is the line that raises the sails. Pull on it, the sails go up. Pull on this one, your AI work comes into focus.
Your AI work leaves a trail. Halyard makes that trail legible, auditable, and client-safe.
Every AI session — time, tokens, model, cost, project — captured where the work happens,
stored as plain text on your machine, owned by you. No account. No cloud service. No
prompt or code capture. Ever. MIT licensed.
Status: alpha, open source — capture loop, reports, invoices, The Bridge, and TUI
in daily use.
You're doing AI-assisted work. At the end of a sprint, a month, or a client engagement,
you can't answer three basic questions:
What did AI actually cost on this project?
What did AI help produce — and can you prove it?
Is your AI spend going in the right direction?
Your tools don't record this. Halyard does.
It runs as lightweight hooks inside Claude Code, Cursor, and Gemini CLI, with
manual/editor-task capture for tools like VS Code where no public AI-session hook
exists. Every session writes one line to a plain-text log you own. From that log:
cost breakdowns, project attribution, invoice evidence, and eventually a signed,
verifiable AI work appendix you can hand to a client.
The privacy promise is unconditional: Halyard captures session metadata, never prompt
content, code context, file contents, or transcripts.
Individual developers and freelancers — your primary audience right now.
Your time, your AI spend, and your invoice evidence live as plain text on your
laptop. Halyard helps you prove what happened without exposing prompts or code.
Git it, back it up, sync it however you want. No account. No SaaS. No
proprietary format.
Small AI shops — share the same local ledger format across a team.
Project spend, trust-labeled cost allocation, and client-safe appendices built
from individual plain-text logs.
Enterprise — the same format will support governance, cost centers, and
cross-tool AI Work Intelligence later. That layer is additive, gated on
design-partner pull, and will not change what local files mean.
Collection — Lightweight hooks that run where AI work happens. Claude Code,
Cursor, Windsurf, and Gemini CLI hooks capture sessions automatically; a VS Code
extension captures editing time, branch, and code delta — token counts
unavailable until VS Code/Copilot exposes a public session hook. Since v3.5,
Claude Code sessions are tagged with an advisory client surface
( cli / desktop / ide ) detected from the local environment. Captured fields
include time, tokens when available, model, cost, project, and branch. Written
to a plain-text log you own. New sessions are appended, and the current
hardening track is making corrections explicit and auditable. Nothing is lost
silently.
Intelligence — Analytics built on that log. Local CLI reports, cost-by-project breakdowns, per-model spend, budget alerts, and trust-labeled totals (captured vs. calculated vs. allocated). Works offline, no account required.
AI Work Ledger — Cost allocation for seat subscriptions and credit plans. If you pay $200/month for Claude Max, Halyard allocates that cost across your projects proportionally — by active minutes, session count, or credit usage — so you know what each client engagement actually costs. Runs on top of ai-sessions.log and ai-plans.toml ; nothing is written back to the raw log.
Proof Artifacts — Invoice evidence today, and a signed attestable AI work
appendix next. The goal is a client-safe artifact that proves AI-assisted work
without showing prompts, transcripts, source code, or file contents.
The Bridge — A local dashboard for watching capture happen in real time. Run halyard dashboard inside any Halyard project.
Rich Session Telemetry — Where tools expose it, Halyard captures operational metadata beyond cost: tool call counts, error rates, wall time vs. active agent time, code delta, and per-model breakdowns. Gemini CLI sessions include full multi-model breakdowns from the history file. These signals surface in the TUI and The Bridge as work-health indicators — not productivity scores, but honest signals of session shape.
Honors — A service record that rewards clean proof, not raw hours. Ranks advance on attributed sessions (Deckhand → Commodore), stripes track your watch streak, and eight medals recognize behaviors that matter: completing your first watch, keeping a clean manifest, rescuing adrift sessions, and more. Run halyard honors to see your record.
Friends of the Sea — One sea creature per completed project, auto-assigned by personality. Projects move through nautical voyage stages (Anchors Aweigh → Making Headway → Rounding the Mark → Flying Colors → Shipshape · Moored) as sessions accumulate. Auto-completes on target hit or inactivity. Run halyard voyage to see the roster. Your Captain's Quarters on The Bridge shows a Passport — one stamp per AI tool you've used.
Halyard is a Python 3.11+ local-first CLI and dashboard, not a hosted billing
service. The halyard command is a Typer app, reports use Rich, the terminal
dashboard uses Textual, and The Bridge is a small 127.0.0.1 HTTP
server. The durable data model is plain text in the project folder; SQLite is
only a rebuildable read-model cache for faster queries.
The capture pipeline is intentionally simple:
AI tool hook/importer/manual command
-> normalized AiSession object
-> append-focused ai-sessions.log line
-> reports, ledger allocation, dashboard, invoice evidence
So: is Halyard "just looking at logs"? Not exactly. It uses the best public
signal each tool exposes:
Claude Code : installs UserPromptSubmit and Stop hooks. The start hook
records session start time and git SHA. The stop hook reads the structured
hook payload; for newer Claude Code formats it can also aggregate token/model
metadata from the local transcript JSONL path passed by the hook.
Cursor : installs beforeSubmitPrompt and stop hooks. It reads the stop
payload and prefers workspace_roots for attribution because that is the
actual editor workspace, not necessarily the shell CWD.
Gemini CLI : installs SessionStart , AfterModel , and AfterAgent hooks.
AfterModel accumulates token usage from usageMetadata ; AfterAgent
finalizes the session. It integrates deeply with OpenTelemetry (OTLP) to measure
exact API and tool-execution durations ( api_seconds , tool_seconds ) and enriches
from Gemini's local history file for accurate multi-model token breakdowns, tool-call counts,
and deterministic cost.
Codex Desktop : imports local ~/.codex/sessions/.../rollout-*.jsonl
files, extracts timing/model/token metadata, and records imported session IDs
so repeated imports do not duplicate entries.
VS Code / GitHub Copilot : a local VS Code extension ( vscode-extension/ )
tracks active editing time, captures branch and code-delta via git, and writes
sessions through halyard record-session . Install the extension from a local
.vsix build; no public Copilot session hook exists yet so token counts are
not available.
Every collector writes the same normalized record shape: timestamps, tool,
model, tokens when available, cache tokens, cost, billing type, project,
branch, capture source, and attribution provenance. Halyard does not store
prompts, source code, file contents, or full transcripts in ai-sessions.log .
When a collector temporarily reads a local transcript or history file, it is
only to extract session metadata.
The log is append-focused. Session records are s ... lines; corrections are
separate amendment records keyed by a hash of the original line. Writers
hold an exclusive OS-level file lock ( fcntl.flock on POSIX, msvcrt.locking
on Windows) so concurrent hooks do not interleave writes. Malformed records
are quarantined instead of crashing report generation.
Cost handling is explicit about trust. Direct API usage can be captured or
calculated from tokens and the local pricing table. Seat or credit plans are
allocated at report time from ai-plans.toml by active minutes, session count,
or credits. Reports label the result as captured, calculated, allocated,
inferred, mixed, or unallocated so client-facing evidence does not pretend an
estimate is a measurement.
Platform: macOS, Linux, and Windows. The halyard service install
command is macOS-only; other platforms can run halyard dashboard
in a long-lived terminal instead.
pipx install halyard
cd ~ /businesses/my-freelance
halyard init
# Guided setup installs supported hooks and checks readiness:
halyard setup
# Verify your hooks and first capture:
halyard doctor --first-capture
# Open the dashboard
halyard dashboard
Daily Workflow
# Start your human timer (AI sessions are auto-captured in the background)
halyard start acme/auth-migration
# ... do work ...
halyard stop
# Terminal UI
halyard tui
# Ask in natural language — two ways, both read-only:
halyard log " what did I spend this month? " # from the CLI (local by
# default; --agent claude|openai optional)
# …or via your coding agent through the MCP server (see below)
# Generate an invoice with an AI usage evidence appendix
halyard invoice acme --month 2026-05 --include-ai-evidence
View Full Command Reference
# Stats-forward analytics — sessions, streaks, peak hour, model mix
halyard usage --range 30d
halyard usage --range 7d --json # machine-readable
# --json on report/usage/budget/status/evidence/health/doctor for CI + scripts
halyard report --all --json | jq ' .totals.cost_usd '
halyard budget --json | jq ' .[] | select(.month.state=="over") ' # spend gate
# Install the background launchd service (macOS):
halyard servi

[truncated]
