---
source: "https://github.com/acunningham-ship-it/leakproof"
hn_url: "https://news.ycombinator.com/item?id=48576074"
title: "Show HN: Leakproof – local secret-egress firewall for Claude Code/Cursor/aider"
article_title: "GitHub - acunningham-ship-it/leakproof: Local secret firewall for AI coding assistants (Claude Code/Cursor/aider) — scans + redacts secrets before anything leaves the machine. Local-first, zero cloud. MIT. · GitHub"
author: "TrustLayerDev"
captured_at: "2026-06-17T21:42:48Z"
capture_tool: "hn-digest"
hn_id: 48576074
score: 1
comments: 1
posted_at: "2026-06-17T20:05:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Leakproof – local secret-egress firewall for Claude Code/Cursor/aider

- HN: [48576074](https://news.ycombinator.com/item?id=48576074)
- Source: [github.com](https://github.com/acunningham-ship-it/leakproof)
- Score: 1
- Comments: 1
- Posted: 2026-06-17T20:05:13Z

## Translation

タイトル: Show HN: Leakproof – クロード コード/カーソル/エイダー用のローカル シークレット出力ファイアウォール
記事のタイトル: GitHub - acuningham-ship-it/leakproof: AI コーディング アシスタント (クロード コード/カーソル/エイダー) 用のローカル シークレット ファイアウォール — 何かがマシンから流出する前にシークレットをスキャンして秘匿化します。ローカルファースト、ゼロクラウド。マサチューセッツ工科大学· GitHub
説明: AI コーディング アシスタント (クロード コード/カーソル/エイダー) 用のローカル シークレット ファイアウォール — 何かがマシンから流出する前にシークレットをスキャンし、秘匿化します。ローカルファースト、ゼロクラウド。マサチューセッツ工科大学- アキュニンガム-シップイット/漏れ防止

記事本文:
GitHub - acuningham-ship-it/leakproof: AI コーディング アシスタント (クロード コード/カーソル/エイダー) 用のローカル シークレット ファイアウォール — 何かがマシンから流出する前にシークレットをスキャンして秘匿化します。ローカルファースト、ゼロクラウド。マサチューセッツ工科大学· GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
アカニンガム・シップ・イット
/
私

防水性
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット ドキュメント ドキュメント ランディング ランディング src/ リークプルーフ src/ リークプルーフ テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md conftest.py conftest.py pyproject.toml pyproject.toml pytest.ini pytest.ini すべてのファイルを表示リポジトリ ファイルのナビゲーション
AI コーディング アシスタント用のローカル ファースト シークレット ファイアウォール。
セキュリティ チームは、データの流出を理由にクロード コードまたはカーソルを禁止しました。これは、それらをオンに戻すためのローカル技術制御です。
Leakproof はツールとモデル API の間に位置し、すべての送信リクエストをマシンから送信される前に読み取ります。シークレットを見つけると、それを編集するか、リクエストを強制終了します。雲には何も当たりません。決定はラップトップ上で行われます。これが自滅的ではない唯一のセットアップです。見知らぬ人にキーを渡して、それがキーかどうかを尋ねることはありません。
# AI ツールをラップします。送信されるものはすべて最初にスキャンされ、クリーンアップされます。
漏れ防止の実行 -- クロード
漏れ防止の実行 -- 補助者
# リポジトリ自体を保護します: シークレットがコミットに達する前に停止します
漏れ防止取り付けフック
誰のためのものか
SOC 2 / HIPAA / ITAR / GDPR に基づくコンプライアンス重視のチーム。そのセキュリティ チームは、AI コーディング ツールをブロックしました。これは、ツールが作業ツリーのコンテキスト (開いているファイル内のあらゆるシークレットを含む) をクラウド API に漏洩するためでした。漏洩防止は、反対意見を満たすローカルの技術管理および監査証跡です。
代替ツール (GitGuardian の ggshield に最近追加された Claude Code と Cursor フック) にはクラウド アカウントが必要です。スキャン メタデータはマシンから離れます。これは構造的に、これを最も必要とする店舗にとっては想定外のことだ。リークプルーフにはクラウド依存性がありません

y — アカウント、API キー、テレメトリはなく、建物からは何も出ません。
24 件の敵対的スイートを含む 148 件のテスト。ルールのみのパス: 15/15 の仕掛けられたリークが捕捉され、0/9 がおとりでの誤検知 (AWS doc-example キー、git SHA、リテラルなしの env 読み取り - すべて正しく無視されました)。
最初のパスでキャッチします (ローカル モデルは必要ありません): AWS アクセス キーと秘密キー、GitHub/OpenAI/Anthropic/Stripe トークン、JWT、PEM 秘密キー、生の .env 値、高エントロピー BLOB、メール、電話、カード番号。
2 番目のパスはオプションです。変数名ではなく値を読み取るローカル モデルのセマンティック チェック (ollam 経由の qwen2.5:1.5b) です。ここで、キーワード スキャナーは機能しません。
detect-secrets は、一般的なコミット前のベースラインです。キーワード マッチングと行ごとのエントロピーを使用します。
正直なフレーミング: Leakproof は、変数名が無害である場合にスキャナーが見逃したキーワードをキャッチします。ローカル モデル セマンティック パスはオプトインであり、追加的です。これを使用しても、使用しなくても、完全な正規表現 + エントロピー レイヤーを取得できます。
pipx インストール漏れ防止
# またはインストールせずに実行します。
uvx リークプルーフ ラン -- クロード
Python 3.10以降。プロキシ サーフェスには aiohttp が必要です。 pipx install 'leakproof[proxy]' または uvx 'leakproof[proxy]' run -- claude でインストールします。
漏れ防止の実行 -- クロードは ANTHROPIC_BASE_URL (または aider の OPENAI_API_BASE) を 127.0.0.1:8747 のローカル プロキシに設定し、ツールを起動します。プロキシは各リクエストの本文を読み取り、スキャナを実行し、編集されたコピーを上流に転送し、応答をそのままストリームバックします。インストールする証明書も、システム全体のプロキシも、ラップするように要求していないものを傍受することもありません。
すべてのキャッチは、 ~/.local/share/leakproof/audit.jsonl にある追加専用の監査ログに記録されます。漏れ防止時計はそれを後押しします:
$漏れ防止時計
14:02:11 claude-code → api.anthropic.com が aws_secret_key を編集しました (重要)

14:02:11 claude-code → api.anthropic.com が .env から STRIPE_SECRET_KEY を編集しました
14:06:48 aider → api.openai.com が private_key (PEM) をブロックしました
このセッション: 3 つのシークレットが停止し、0 つはクラウドに到達しました
モード
監視 — ログのみ。何も変化しません。ワークフローを中断することなく、何が残っているかを確認するには、まずこれを使用してください。
redact — 各検出結果をプレースホルダーと交換し、クリーンアップされたリクエストを転送します。デフォルト。
ブロック — リクエストを 403 で完全に拒否し、漏洩した可能性があるものを指定します。
CLI は Apache-2.0 で、無料です。開発者は 1 人、アカウントも壁もありません。
リークプルーフ チームは、ラップトップごとのファイル以上のものを必要とするコンプライアンス ショップ向けです。チーム全体が継承する共有リダクション ポリシー、マシン間で集約された中央監査ログ、シークレットが配布されるはずのときにビルドを失敗させる CI ゲート、SOC 2 または HIPAA フォルダーに直接ドロップできる署名付き監査証拠エクスポートが追加されます。
早期アクセスと価格: hamstudios101@gmail.com
現在動作します: Claude Code および aider (ベース URL 環境変数を尊重するツール)。 Cursor と Copilot は、実際の HTTPS インターセプト プロキシと証明書のインストールを必要とする独自のバックエンドを使用します。これは、v1 ではなく v1.1 です。マシンは 1 台で、デーモンもテレメトリもありません。
Apache-2.0。ハムスタジオによって構築されました。問題やPRを歓迎します。
AI コーディング アシスタント (クロード コード/カーソル/エイダー) 用のローカル シークレット ファイアウォール — 何かがマシンから流出する前にシークレットをスキャンし、秘匿化します。ローカルファースト、ゼロクラウド。マサチューセッツ工科大学
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local secret firewall for AI coding assistants (Claude Code/Cursor/aider) — scans + redacts secrets before anything leaves the machine. Local-first, zero cloud. MIT. - acunningham-ship-it/leakproof

GitHub - acunningham-ship-it/leakproof: Local secret firewall for AI coding assistants (Claude Code/Cursor/aider) — scans + redacts secrets before anything leaves the machine. Local-first, zero cloud. MIT. · GitHub
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
acunningham-ship-it
/
leakproof
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits docs docs landing landing src/ leakproof src/ leakproof tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md conftest.py conftest.py pyproject.toml pyproject.toml pytest.ini pytest.ini View all files Repository files navigation
Local-first secret firewall for AI coding assistants.
Your security team banned Claude Code or Cursor over data egress. Here's the local technical control that lets you turn them back on.
leakproof sits between the tool and the model API and reads every outbound request before it leaves the machine. Finds a secret, it redacts it or kills the request. Nothing hits the cloud. The decision happens on your laptop, which is the only setup that isn't self-defeating — you don't hand a key to a stranger to ask them whether it's a key.
# wrap your AI tool: everything it sends gets scanned + cleaned first
leakproof run -- claude
leakproof run -- aider
# guard the repo itself: stop secrets before they reach a commit
leakproof install-hook
Who it's for
Compliance-bound teams under SOC 2 / HIPAA / ITAR / GDPR whose security team blocked AI coding tools because the tools exfiltrate working-tree context — including any secrets in open files — to a cloud API. leakproof is the local technical control and audit trail that satisfies the objection.
The alternative tools (GitGuardian's ggshield recently added Claude Code and Cursor hooks) require a cloud account: scan metadata leaves the machine. That's structurally off the table for the shops that most need this. leakproof has zero cloud dependency — no account, no API key, no telemetry, nothing leaves the building.
148 tests, including a 24-case adversarial suite. Rules-only pass: 15/15 planted leaks caught, 0/9 false-positives on decoys (AWS doc-example keys, git SHAs, env reads without literals — all correctly ignored).
Catches on the first pass (no local model needed): AWS access keys and secret keys, GitHub/OpenAI/Anthropic/Stripe tokens, JWTs, PEM private keys, raw .env values, high-entropy blobs, email, phone, card numbers.
The second pass is optional — a local-model semantic check (qwen2.5:1.5b via ollama) that reads the value rather than the variable name. That's where keyword scanners break down.
detect-secrets is a common pre-commit baseline. It uses keyword matching plus entropy on a per-line basis.
The honest framing: leakproof catches what keyword scanners miss when the variable name is benign . The local-model semantic pass is opt-in and additive — you get the full regex+entropy layer with or without it.
pipx install leakproof
# or run without installing:
uvx leakproof run -- claude
Python 3.10+. The proxy surface needs aiohttp — install with pipx install 'leakproof[proxy]' or uvx 'leakproof[proxy]' run -- claude .
leakproof run -- claude sets ANTHROPIC_BASE_URL (or OPENAI_API_BASE for aider) to a local proxy on 127.0.0.1:8747 , then launches the tool. The proxy reads each request body, runs the scanner, forwards a redacted copy upstream, and streams the response back untouched. No certificate to install, no system-wide proxy, no interception of anything you didn't ask it to wrap.
Every catch lands in an append-only audit log at ~/.local/share/leakproof/audit.jsonl . leakproof watch tails it:
$ leakproof watch
14:02:11 claude-code → api.anthropic.com redacted aws_secret_key (critical)
14:02:11 claude-code → api.anthropic.com redacted STRIPE_SECRET_KEY from .env
14:06:48 aider → api.openai.com blocked private_key (PEM)
this session: 3 secrets stopped, 0 reached the cloud
Modes
monitor — logs only, nothing changes. Use this first to see what's been leaving without disrupting your workflow.
redact — swaps each finding for a placeholder and forwards the cleaned request. Default.
block — rejects the request outright with a 403 and names what would have leaked.
The CLI is Apache-2.0 and free. One developer, no account, no wall.
leakproof Team is for compliance shops that need more than a per-laptop file. It adds: a shared redaction policy your whole team inherits, a central audit log aggregated across machines, a CI gate that fails the build when a secret would have shipped, and signed audit-evidence exports you can drop straight into your SOC 2 or HIPAA folder.
Early access and pricing: hamstudios101@gmail.com
Works today: Claude Code and aider (any tool that honors a base-URL env var). Cursor and Copilot use proprietary backends that need a real HTTPS intercept proxy and a cert install — that's v1.1, not v1. One machine, no daemon, no telemetry.
Apache-2.0. Built by hamstudios . Issues and PRs welcome.
Local secret firewall for AI coding assistants (Claude Code/Cursor/aider) — scans + redacts secrets before anything leaves the machine. Local-first, zero cloud. MIT.
Apache-2.0 license
Contributing
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
