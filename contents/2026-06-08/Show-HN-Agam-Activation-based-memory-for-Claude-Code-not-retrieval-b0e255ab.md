---
source: "https://github.com/CrypticCortex/agam"
hn_url: "https://news.ycombinator.com/item?id=48441197"
title: "Show HN: Agam – Activation-based memory for Claude Code, not retrieval"
article_title: "GitHub - CrypticCortex/agam · GitHub"
author: "aghoraguru"
captured_at: "2026-06-08T04:36:23Z"
capture_tool: "hn-digest"
hn_id: 48441197
score: 1
comments: 0
posted_at: "2026-06-08T04:00:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agam – Activation-based memory for Claude Code, not retrieval

- HN: [48441197](https://news.ycombinator.com/item?id=48441197)
- Source: [github.com](https://github.com/CrypticCortex/agam)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T04:00:51Z

## Translation

タイトル: Show HN: Agam – 検索ではなく、クロード コードのアクティベーション ベースのメモリ
記事タイトル: GitHub - CrypticCortex/agam · GitHub
説明: GitHub でアカウントを作成して、CrypticCortex/agam の開発に貢献します。

記事本文:
GitHub - CrypticCortex/agam · GitHub
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
クリプティックコーテックス
/
アガム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル

レ
43 コミット 43 コミット .github .github ナレッジ ナレッジ プロンプト プロンプト スクリプト スクリプト src/ agam src/ agam テンプレート テンプレート テスト テスト .coderabbit.yaml .coderabbit.yaml .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードの上にナレッジ グラフを利用したアイデンティティとコンテキスト インジェクション レイヤーを追加します。 Agam は、UserPromptSubmit フックを介して関連エンティティ (プロジェクト、サービス、意思決定、バグ、レッスン) をすべてのクロード コード セッションに自動挿入するため、モデルは毎回ファイルを検索するのではなく、履歴から応答します。以前のセッションのトランスクリプトに対するブートストラップ パスを介してグラフを 1 回シードすると、Agam がバックグラウンドでグラフを保温し続けます。
~/.claude/agam/ (AGAM.md、THISAI.md、MUGAM.md) にある ID ファイルにより、Claude Code は、あなたが誰であるか、そして何に取り組んでいるのかを永続的に把握できます。
~/.claude/knowledge/graph.db にある SQLite ナレッジ グラフには、エンティティと FTS5 検索との関係が保存されます。
グラフ呼び出しフック (UserPromptSubmit) は、プロンプト内のエンティティ名を照合し、モデルが応答する前にそのコンテキストをインラインで挿入します。
グラフ更新フック (Stop) は、終了したセッションをキューに入れ、launchd で管理されるウォッチドッグによるバックグラウンド処理を実行します。
agam bootstrap コマンドは、 ~/.claude/projects/ にある既存のクロード コードのトランスクリプトからグラフをシードします。
Anthropic API キーは必要なく、サポートされていません。 Agam が行うすべての LLM 呼び出しは、 docker exec を介して既存の claude-code devcontainer 内で実行され、すでに認証されている OAuth 資格情報を再利用します。
macOS。 v1 でサポートされているプラ​​ットフォームのみ。
Claude コードがインストールされ、認証されました (claude を対話的に 1 回実行します)。 Agam は、Claude Code がすでに持っている認証を再利用します。どこに生息しているかを探す必要はありません。
パイソン3

.11 以降 (持っていない場合は uv が取得します)。
オプションですが強くお勧めします: クロードコードの開発コンテナーが実行されている Docker デスクトップ。ブートストラップ パイプラインとバックグラウンド ウォッチドッグは両方とも docker exec をシェルアウトします。 ID ファイルとグラフリコール フックは Docker なしで動作するため、Docker の準備がまだ整っていないマシンに Agam をインストールできます。
インストーラは必要な前提条件をすべて検証し、何かが不足している場合は有用なエラーを表示して救済します。
git clone < リポジトリ URL > ~ /coding/agam
cd ~ /coding/agam
./install.sh
install.sh は最低限のことを行います。
uv 、 claude 、 Docker (存在しない場合は警告)、および macOS をチェックします。認証は、資格情報が保存されている場所を推測することによってではなく、後で agam Doctor によって検証され、最初の実際の claude -p 呼び出し時に検証されます。
uv sync を実行して Python 環境を具体化します。
uv へのデリゲートは、実際のインストーラーである agam init を実行します。
agam init は対話型の質問ウィザードです。いくつかの質問 (名前、主な目標、プロジェクト ディレクトリ、コンテナ モード) が行われ、次のことが行われます。
templates/*.template を ~/.claude/agam/ (AGAM.md、THISAI.md、MUGAM.md、CLAUDE.md スニペット) にレンダリングします。
Agam フックを ~/.claude/settings.json にマージします (最初に既存ファイルのタイムスタンプ付きバックアップを作成します)。
watchdog launchd plist を ~/Library/LaunchAgents/com.agam.watchdog.plist に書き込み、ロードします。
~/.claude/knowledge/graph.db が存在しない場合は、FTS5 スキーマを使用して作成します。
インストーラーを再実行しても安全です。デフォルトでは、既存の ~/.claude/agam/ の上書きは拒否されます。 --force を渡して、以前のインストールのタイムスタンプ付きバックアップで上書きします。
uv run agam init --force
YAML 応答ファイルを指定してウィザードを非対話的に実行することもできます。
uv run agam init --answers my-answers.yaml
ブートストラップのウォークスルー
ブートストラップ パスはオプションですが、強くお勧めします。こう書かれています

クロード コード セッションのトランスクリプトを作成し、ナレッジ グラフにエンティティと関係を追加します。
agam ブートストラップ -- 30 日目
請求前にコストのプレビューが表示されます。
[agam ブートストラップ] プロジェクト ディレクトリ: /Users/you/.claude/projects
[agam ブートストラップ] トランスクリプト: 47 (日フィルター: 30)
[agam ブートストラップ] 推定トークン: ~412,000
[agam ブートストラップ] 推定コスト: ~$0.4536
続行しますか? [y/N]
y と答えて実行します。パイプラインには 2 つのフェーズがあります。
抽出 -- Haiku は各トランスクリプトのチャンクを読み取り、候補となるエンティティと関係を出力します。
調整 -- Sonnet は重複をマージし、名前のバリエーションを解決し、グラフに書き込まれるクリーンなペイロードを返します。
状態は、トランスクリプトごとにチェックポイントが作成されます。プロセスが中断された場合 ( Ctrl-C 、クラッシュ、ラップトップのスリープ)、同じコマンドを再実行するだけです。中断したところから正確に再開されます。
agam bootstrap --days 30 # デフォルトで再開
agam bootstrap --no-resume # クリーン実行を強制します
便利なフラグ:
コスト見積もりでは、トークンあたり 4 文字のヒューリスティックを使用し、Haiku 入力トークン 100 万あたり 0.80 ドル、Sonnet 入力トークン 100 万あたり 3.00 ドル、トークン予算の約 10% が調整にルーティングされます。これらのレートは src/agam/bootstrap.py のデフォルトです。価格が異なる場合は、そこで上書きしてください。
インストーラーによって ~/.claude/agam/config.yaml に書き込まれます。
このファイルを直接編集すると、次のセッションで変更が反映されます。ウィザードでテンプレートからすべてを再生成する場合は、agam init --force を再実行します。
変数
デフォルト
目的
AGAM_INVOKER
設定を解除する
カスケードを単一の呼び出し元 (ホストまたはコンテナー) に固定します。自動検出したい場合はスキップしてください。
AGAM_CONTAINER_PATTERN
クロードコード
正規表現は docker ps 行と照合され、クロード コード コンテナーを検出します。
AGAM_CONTAINER_NAME
設定を解除する
正確なコンテナ名。上記の正規表現を破ります。
AGAM_WATCHDOG_MODE
設定を解除する
AGAM_INVOK のレガシー エイリアス

えー。バックコンパットで表彰されました。
AGAM_HOME
~/.クロード/アガム
ID + ログ ディレクトリのルート。
AGAM_KG_PATH
~/.claude/knowledge/graph.db
SQLite グラフへのパス。
AGAM_PROMPTS_DIR
バンドルされた
ブートストラップ プロンプト テンプレートを保持するディレクトリ。
最後の 3 つが必要になることはほとんどありません。これらはテスト用と、デフォルト以外のレイアウトから Agam を実行するコントリビューター用に存在します。
Agam は、グラフをブートストラップするとき、または終了したセッションを処理するときに、どこかで claude -p を呼び出す必要があります。この場所は実行時に自動的に検出されます。インストール時に「モード」を選択する必要はありません。
検出はカスケードです。 Agam はリストを順番に調べて、健全性を調べた最初のリストを使用します。
AGAM_INVOKER (または従来の AGAM_WATCHDOG_MODE ) を介して固定 -- 設定されている場合、その呼び出し者が唯一の候補になります。
名前付きコンテナ -- AGAM_CONTAINER_NAME が実行中のコンテナを指している場合。
検出されたコンテナ -- イメージ名が AGAM_CONTAINER_PATTERN (デフォルトの claude-code ) と一致する最初の docker ps 行。
ホスト claude -- PATH 上の claude 。 (認証は、クロード コードが配置する場所 (macOS ホスト上のキーチェーン、コンテナ内のファイル) に存在します。プローブはクロード コードを後から推測しません。CLI が PATH 上にある場合は対象となり、実際の認証失敗は実行時にクロード自身のエラーとして表面化します。)
「コンテナ」と「ホスト」はどちらもファーストクラスです。 Agam のバックグラウンド呼び出しは対話型のクロード コード セッションから分離されて実行されるため、両方が利用可能な場合はコンテナーが推奨されます。使用可能なものが 1 つだけの場合、カスケードは尋ねることなくそれを選択します。
具体的には、セッションが終了すると次のようになります。
Stop フックは、キュー エントリを ~/.claude/agam/queue/ に書き込みます。
ウォッチドッグ launchd エージェントは数分ごとに動作し、呼び出し元を解決し、キューを空にします。コンテナー: docker exec <検出された名前> claude -p ... 。ホスト: claude -p ... 直接。
健全な呼び出し元が存在しない: キューはそのまま残り、

~/.claude/agam/logs/watchdog.log には、すべてのプローブの失敗をリストした no-invoker 行が記録されます。次のティックで再試行されます。
Agam が現在利用できると考えているものを確認するには:
アガム医師
これにより、実行者候補ごとに 1 行に PASS/WARN/FAIL とその理由が出力されます。 「自動学習が停止した」場合の最初の停止として使用します。通常は、答えが表示されます (コンテナーが停止した、OAuth トークンの有効期限が切れた、クロードが PATH 上にない)。
Agam は Anthropic API キーを取得しません。すべての claude -p 呼び出しは、クロード コードが配置することを選択した場所 (キーチェーン、 ~/.claude/.credentials.json など) の既存のクロード コード OAuth を経由します。 API キーが必要な場合は、間違ったツールを使用しています。
アガムステータス
これにより、Agam ホーム パス、ナレッジ グラフ サイズ、キューの深さ、ブートストラップの再開状態、および検出されたコンテナー名 (存在する場合) が出力されます。グラフやキューには触れないため、いつでも安全に実行できます。
コンテナ: (何も検出されません) -- クロードコードの開発コンテナが実行されていません。それを開始し、 agam status を再実行します。それでも検出が失敗し、カスタム名がある場合は、 AGAM_CONTAINER_NAME を docker ps で表示される正確な名前に設定します。
~/.claude/agam/logs/watchdog.log の no-container 行 -- コンテナがダウンしているときは常に予想されます。キューは、コンテナーを開始した後の次のティックで排出されます。
キューがスタックしています / ~/.claude/agam/queue-errors/ のエントリ -- セッションの処理が失敗しました。エラー ペイロードを開いて、基になる例外を確認します。ログ:
tail -n 100 ~ /.claude/agam/logs/watchdog.log
ls ~ /.claude/agam/queue-errors/
エラー: agam ブートストラップから実行されているクロード コード コンテナーがありません -- 上記と同じ修正です。コンテナを起動して再実行します。ブートストラップが自動的に再開されます。
graph-recall は何も挿入していません。最初に、フックが hooks.UserPromptSubmit の下の ~/.claude/settings.json にあることを確認します。次に、グラップを確認します

h にはエンティティがあります: sqlite3 ~/.claude/knowledge/graph.db 'エンティティから count(*) を選択します;' 。ブートストラップを実行せずに新しくインストールされた Agam は空です。それが最も一般的な原因です。
最初に保存しておきたいものはすべてバックアップします。
cp -r ~ /.claude/agam ~ /agam-backup- $( date +%Y%m%d )
cp ~ /.claude/knowledge/graph.db ~ /graph-backup- $( date +%Y%m%d ) .db
次に、分解します。
# ウォッチドッグを停止して削除します
launchctl unload ~ /Library/LaunchAgents/com.agam.watchdog.plist
rm ~ /ライブラリ/LaunchAgents/com.agam.watchdog.plist
# settings.json から Agam フックを削除します
# エディターで ~/.claude/settings.json を開き、エントリを削除します
# コマンド パスは agam または ~/.claude/agam/ を参照します。
# 白紙の状態にしたい場合は、アイデンティティ ファイルとナレッジ グラフを削除します
rm -rf ~ /.claude/agam
rm -f ~ /.claude/knowledge/graph.db
~/coding/agam のリポジトリは独立しています。ソースが不要になった場合は、個別に削除してください。
コマンド
目的
アガム初期化
Agam スキャフォールディングを ~/.claude/ にインストールします。 --force を使用して上書きし、--answers FILE.yaml をスクリプトに使用します。
アガムブートストラップ
トランスクリプトをスキャンし、コストを見積もり、抽出してナレッジ グラフに調整します。再開可能。
アガムステータス
インストールの健全性の印刷: パス、グラフ サイズ、キューの深さ、コンテナーの検出、再

[切り捨てられた]

## Original Extract

Contribute to CrypticCortex/agam development by creating an account on GitHub.

GitHub - CrypticCortex/agam · GitHub
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
CrypticCortex
/
agam
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
43 Commits 43 Commits .github .github knowledge knowledge prompts prompts scripts scripts src/ agam src/ agam templates templates tests tests .coderabbit.yaml .coderabbit.yaml .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Knowledge-graph-powered identity and context injection layer on top of Claude Code. Agam auto-injects relevant entities (projects, services, decisions, bugs, lessons) into every Claude Code session via a UserPromptSubmit hook, so the model answers from your history instead of searching files every time. You seed the graph once via a bootstrap pass over prior session transcripts, then Agam keeps it warm in the background.
Identity files at ~/.claude/agam/ (AGAM.md, THISAI.md, MUGAM.md) give Claude Code a persistent sense of who you are and what you are working on.
A SQLite knowledge graph at ~/.claude/knowledge/graph.db stores entities + relationships with FTS5 search.
The graph-recall hook (UserPromptSubmit) matches entity names in your prompts and injects their context inline, before the model answers.
The graph-update hook (Stop) enqueues finished sessions for background processing by a launchd-managed watchdog.
The agam bootstrap command seeds the graph from your existing Claude Code transcripts in ~/.claude/projects/ .
No Anthropic API key is needed or supported. Every LLM call Agam makes runs inside your existing claude-code devcontainer via docker exec , reusing the OAuth credentials you already authenticated with.
macOS. Only platform supported in v1.
Claude Code installed and authenticated (run claude interactively once). Agam reuses whatever auth Claude Code already has; you do not need to find where it lives.
Python 3.11 or newer (uv will fetch one if you do not have it).
Optional but strongly recommended: Docker Desktop with a running claude-code devcontainer. The bootstrap pipeline and the background watchdog both shell out to docker exec . The identity files and the graph-recall hook work without Docker, so you can install Agam on a machine where Docker is not ready yet.
The installer verifies every required prerequisite and bails with a useful error if something is missing.
git clone < repo-url > ~ /coding/agam
cd ~ /coding/agam
./install.sh
install.sh does the minimum:
Checks for uv , claude , Docker (warns if absent), and macOS. Auth is verified later by agam doctor and at the first real claude -p call, not by guessing where credentials are stored.
Runs uv sync to materialize the Python environment.
Delegates to uv run agam init , which is the real installer.
agam init is an interactive questionary wizard. It asks a few questions (name, primary goal, projects directory, container mode), then:
Renders templates/*.template into ~/.claude/agam/ (AGAM.md, THISAI.md, MUGAM.md, CLAUDE.md snippet).
Merges Agam hooks into ~/.claude/settings.json (creating a timestamped backup of the existing file first).
Writes the watchdog launchd plist to ~/Library/LaunchAgents/com.agam.watchdog.plist and loads it.
Creates ~/.claude/knowledge/graph.db with the FTS5 schema if it does not exist.
Re-running the installer is safe. By default it refuses to overwrite an existing ~/.claude/agam/ . Pass --force to overwrite with a timestamped backup of the previous install:
uv run agam init --force
You can also drive the wizard non-interactively by feeding it a YAML answer file:
uv run agam init --answers my-answers.yaml
Bootstrap walkthrough
The bootstrap pass is optional but strongly recommended. It reads your Claude Code session transcripts and populates the knowledge graph with entities + relationships.
agam bootstrap --days 30
You will see a cost preview before anything bills:
[agam bootstrap] projects-dir: /Users/you/.claude/projects
[agam bootstrap] transcripts: 47 (days filter: 30)
[agam bootstrap] estimated tokens: ~412,000
[agam bootstrap] estimated cost: ~$0.4536
Proceed? [y/N]
Answer y to run. The pipeline has two phases:
Extraction -- Haiku reads each transcript chunk and emits candidate entities and relationships.
Reconciliation -- Sonnet merges duplicates, resolves name variants, and returns a clean payload that is written into the graph.
State is checkpointed after every transcript. If the process is interrupted ( Ctrl-C , crash, laptop sleep), just re-run the same command; it picks up exactly where it left off:
agam bootstrap --days 30 # resumes by default
agam bootstrap --no-resume # force a clean run
Useful flags:
Cost estimation uses a 4-chars-per-token heuristic, $0.80 per 1M Haiku input tokens, and $3.00 per 1M Sonnet input tokens with ~10% of the token budget routed to reconciliation. These rates are defaults in src/agam/bootstrap.py ; override them there if your pricing differs.
Written into ~/.claude/agam/config.yaml by the installer:
Edit this file directly and the next session picks up the changes. Re-run agam init --force if you want the wizard to regenerate everything from templates.
Variable
Default
Purpose
AGAM_INVOKER
unset
Pin the cascade to a single invoker: host or container . Skip if you want auto-detect.
AGAM_CONTAINER_PATTERN
claude-code
Regex matched against docker ps rows to discover your claude-code container.
AGAM_CONTAINER_NAME
unset
Exact container name. Beats the regex above.
AGAM_WATCHDOG_MODE
unset
Legacy alias for AGAM_INVOKER . Honored for back-compat.
AGAM_HOME
~/.claude/agam
Root of identity + log directories.
AGAM_KG_PATH
~/.claude/knowledge/graph.db
Path to the SQLite graph.
AGAM_PROMPTS_DIR
bundled
Directory holding bootstrap prompt templates.
You will rarely need the last three. They exist for tests and for contributors running Agam out of a non-default layout.
Agam needs to call claude -p somewhere when it bootstraps the graph or processes a finished session. That somewhere is auto-detected at run time -- you do not pick a "mode" at install time.
The detection is a cascade. Agam walks the list in order and uses the first one that probes healthy:
Pinned via AGAM_INVOKER (or the legacy AGAM_WATCHDOG_MODE ) -- if set, that invoker is the only candidate.
Named container -- if AGAM_CONTAINER_NAME points to a running container.
Discovered container -- the first docker ps row whose image name matches AGAM_CONTAINER_PATTERN (default claude-code ).
Host claude -- claude on your PATH . (Auth lives wherever Claude Code put it -- Keychain on macOS host, a file in a container. The probe doesn't second-guess Claude Code; if the CLI is on PATH it's eligible, and a real auth failure surfaces at run time with claude's own error.)
Both "container" and "host" are first-class. Container is preferred when both are available because Agam's background calls then run isolated from your interactive Claude Code session. If only one is available, the cascade picks it without asking.
Concretely, when a session closes:
The Stop hook writes a queue entry into ~/.claude/agam/queue/ .
The watchdog launchd agent ticks every few minutes, resolves the invoker, and drains the queue. Container: docker exec <discovered-name> claude -p ... . Host: claude -p ... directly.
No healthy invoker: the queue stays untouched and ~/.claude/agam/logs/watchdog.log records a no-invoker line listing every probe failure. The next tick tries again.
To see what Agam thinks is available right now:
agam doctor
That prints one line per candidate invoker with PASS/WARN/FAIL and the reason. Use it as the first stop when "auto-learning stopped happening" -- usually it tells you the answer (container stopped, OAuth token expired, claude not on PATH).
Agam never takes an Anthropic API key. Every claude -p invocation goes through your existing Claude Code OAuth -- wherever Claude Code chose to put it (Keychain, ~/.claude/.credentials.json , etc.). If you need an API key, you are using the wrong tool.
agam status
That prints the Agam home path, knowledge graph size, queue depth, bootstrap resume state, and the detected container name (if any). It does not touch the graph or the queue, so it is always safe to run.
Container: (none detected) -- your claude-code devcontainer is not running. Start it, then re-run agam status . If detection still fails and you have a custom name, set AGAM_CONTAINER_NAME to the exact name shown by docker ps .
no-container lines in ~/.claude/agam/logs/watchdog.log -- expected whenever the container is down. The queue will drain on the next tick after you start the container.
Queue stuck / entries in ~/.claude/agam/queue-errors/ -- a session failed processing. Open the error payload to see the underlying exception. Logs:
tail -n 100 ~ /.claude/agam/logs/watchdog.log
ls ~ /.claude/agam/queue-errors/
ERR: no claude-code container running from agam bootstrap -- same fix as above. Start the container and re-run; bootstrap resumes automatically.
graph-recall is not injecting anything -- first confirm the hook is in ~/.claude/settings.json under hooks.UserPromptSubmit . Then confirm the graph has entities: sqlite3 ~/.claude/knowledge/graph.db 'select count(*) from entities;' . A freshly installed Agam with no bootstrap run is empty; that is the most common cause.
Back up anything you want to keep first:
cp -r ~ /.claude/agam ~ /agam-backup- $( date +%Y%m%d )
cp ~ /.claude/knowledge/graph.db ~ /graph-backup- $( date +%Y%m%d ) .db
Then tear down:
# Stop + remove the watchdog
launchctl unload ~ /Library/LaunchAgents/com.agam.watchdog.plist
rm ~ /Library/LaunchAgents/com.agam.watchdog.plist
# Remove Agam hooks from settings.json
# Open ~/.claude/settings.json in your editor and delete the entries
# whose command paths reference agam or ~/.claude/agam/.
# Delete identity files + knowledge graph if you want a clean slate
rm -rf ~ /.claude/agam
rm -f ~ /.claude/knowledge/graph.db
The repo at ~/coding/agam is independent; remove it separately if you no longer want the source.
Command
Purpose
agam init
Install Agam scaffolding into ~/.claude/ . Use --force to overwrite, --answers FILE.yaml to script.
agam bootstrap
Scan transcripts, estimate cost, extract + reconcile into the knowledge graph. Resumable.
agam status
Print install health: paths, graph size, queue depth, container detection, re

[truncated]
