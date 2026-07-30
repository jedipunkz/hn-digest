---
source: "https://github.com/YoanWai/agent-manager"
hn_url: "https://news.ycombinator.com/item?id=49107749"
title: "Agent-Manager: A Tmux TUI for Running Claude Code, Codex and OpenCode"
article_title: "GitHub - YoanWai/agent-manager: Terminal UI to manage AI coding-agent sessions (Claude Code, OpenCode, Codex, Grok Build) in tmux: live status, group tree, live pane preview, resource gauges. · GitHub"
author: "yoanwaidev"
captured_at: "2026-07-30T10:13:43Z"
capture_tool: "hn-digest"
hn_id: 49107749
score: 1
comments: 0
posted_at: "2026-07-30T09:34:27Z"
tags:
  - hacker-news
  - translated
---

# Agent-Manager: A Tmux TUI for Running Claude Code, Codex and OpenCode

- HN: [49107749](https://news.ycombinator.com/item?id=49107749)
- Source: [github.com](https://github.com/YoanWai/agent-manager)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T09:34:27Z

## Translation

タイトル: エージェント マネージャー: クロード コード、コーデックス、および OpenCode を実行するための Tmux TUI
記事のタイトル: GitHub - YoanWai/agent-manager: tmux で AI コーディング エージェント セッション (Claude Code、OpenCode、Codex、Grok Build) を管理するためのターミナル UI: ライブ ステータス、グループ ツリー、ライブ ペイン プレビュー、リソース ゲージ。 · GitHub
説明: tmux で AI コーディング エージェント セッション (Claude Code、OpenCode、Codex、Grok Build) を管理するためのターミナル UI: ライブ ステータス、グループ ツリー、ライブ ペイン プレビュー、リソース ゲージ。 - YoanWai/エージェントマネージャー

記事本文:
GitHub - YoanWai/agent-manager: tmux で AI コーディング エージェント セッション (Claude Code、OpenCode、Codex、Grok Build) を管理するターミナル UI: ライブ ステータス、グループ ツリー、ライブ ペイン プレビュー、リソース ゲージ。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ヨアンW

あい
/
エージェントマネージャー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
193 コミット 193 コミット .github .github docs docs external 内部 .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml ライセンス ライセンス README.md README.md go.mod go.mod go.sum go.sum main.go main.go main_test.go main_test.go すべてのファイルを表示 リポジトリ ファイル ナビゲーション
1 つの端末からすべての AI コーディング エージェントを実行します。 Claude Code、Codex、OpenCode、および Grok は、それぞれ独自の tmux セッションで並行して実行されるため、マネージャーを終了した後も動作し続けます。
どのエージェントが完了し、どのエージェントが停止しているかを確認するためにターミナル タブを探す代わりに、すべてのセッションがライブ ステータスとともに 1 つのリストに表示され、折りたたんで並べ替えることができるプロジェクト ツリーにグループ化されます。添付せずにそれらのいずれかに回答します。スペースはプロンプトをセッションのペインに直接送信するか、選択したグループに新しいエージェントを生成します。停止したセッションは、 v で中断した場所から復活します。また、ctrl+r を押すと、エージェントが変更した内容の完全なファイルの差分が開き、構文が強調表示されます。行に残したコメントは、エージェントのペインに直接戻ります。
ステータス検出は現在、Claude Code 、OpenCode 、Codex 、および Grok Build をそのままサポートしています。他の CLI ツールはセッションとして実行できます。ステータス ルールを含む [tools.<name>] ブロックを追加して、ライブ ステータスを取得します (「設定」を参照)。
brew install yoanwai/tap/agent-manager
tmux が存在しない場合はそれをインストールします。
github.com/YoanWai/agent-manager@latest をインストールしてください
Go 1.26 以降と tmux が必要です。 $(go env GOPATH)/bin にインストールされます。
リリースからダウンロードします (macOS および Linux、amd64/arm64)。
WSL2 内で実行: エージェント マネージャーは、Linux/macOS ツールである tmux 上に存在します。 WSL シェルでは、Homebrew を使用してインストールするか、

e リリースからの Linux バイナリ。
マネージャーは GitHub リリースを 1 日に 1 回チェックし、新しいバージョンがリリースされるとヘッダーに ↑ vX.Y.Z の利用可能なバッジを表示します。取り付けた方法で引き抜きます。
brew upgrade yoanwai/tap/agent-manager # Homebrew
go install github.com/YoanWai/agent-manager@latest # Go
使用法
エージェントマネージャー
セッションは tmux (am_* 名前空間) 内で実行されるため、マネージャーが終了しても存続します。セッション内で、Ctrl+Q を押すとマネージャに戻ります。 Agent-manager --version はバージョンを出力します。
キー
アクション
n
新しいセッション (名前、ツール、ディレクトリ、オプションの開始プロンプト、グループ ピッカー)
g
新しいグループ (名前、親、デフォルトのパス)
入る
セッションのアタッチ/グループのフォールド
Ctrl+Q
セッション内: マネージャーに戻る
K / J (またはShift+↑ / SHIFT+↓)
表示されている兄弟間でセッションまたはグループを並べ替えます
メートル
セッションを別のグループに移動する
r
セッション/編集ツールの名前を変更します。グループ名とデフォルトのパスを編集します
v
死んだセッションを復活させます ( revive_command 、例: claude -- continue 、会話を再開します)
あ/う
セッション、またはグループとそのサブツリー全体をアーカイブ/復元します
d
セッション、またはグループとそのサブツリー全体を削除します
スペース
クイック プロンプト: 選択したセッションに応答するか、選択したグループにエージェントを生成します
Ctrl+R
選択したセッションの変更を確認します: フルスクリーンのファイル全体の差分、エージェントに送信された行コメント
f
折りたたみ/展開グループ
s
設定 (クイックスポーンツール、テーマ、レビューレイアウト)
t
アーカイブ済みビューの切り替え
e
空のグループを表示/非表示にする
/
検索
?
ヘルプ
q
終了 (セッションは継続して実行されます)
クイックプロンプト
スペースを押して、プロンプト バーをサイドバーの下部にドッキングします。バーが開いている間、ターゲットはカーソルに従います (↑↓ はまだ移動します)。
セッション行で Enter を押すと、入力されたテキストがセッションのペインに直接送信されるため、添付しなくてもエージェントはそれをユーザー メッセージとして取得します。バーは開いたままで消えます、読んでください

y は次の答えです。
グループ行で「」と入力すると、グループのデフォルトのパスを使用して、プロンプトが埋め込まれた新しいエージェントがそのグループに生成されます。スポーン ツールは設定 ( s ) のデフォルトで開始し、タブでそれを循環します (claude ↔ opencode ↔ 設定されたツール)。フッターには現在のピックが表示されます。エージェントはプロンプトに対してすぐに作業を開始します。
Esc でバーを閉じます。新しいセッション フォームのオプションのプロンプト フィールドでも、同じ方法でエージェントを起動します。 CLI がフラグの背後にあるプロンプトを取得するツールは、prompt_flag を使用してそれを宣言します (「設定」を参照)。
カスタム名なしで生成されたセッション (すべてのクイック生成、および名前が空白のままのフォーム) は、 claude-a1b2 のようなプレースホルダーを取得し、セッションの広範な機能 (単一のサブタスクではない) の短縮名を使用して、エージェント マネージャーの名前変更 "<name>" を 1 回実行するようエージェントに要求することにより、最初のプロンプトが開きます。このディレクティブは、要求しない限り名前を再度変更しないようにエージェントに指示します。最初のプロンプトでディレクティブを実行できない場合 (/slash コマンド、またはプロンプトがまったくない場合)、ツールの入力ボックスがペインに表示されると、マネージャーはそれを独自のメッセージとして送信します。サブコマンドは、名前をセッションごとのファイルにドロップします。マネージャーは次のポーリングでそれを取得し、サイドバーの行と tmux ステータス バーを更新します。これは、エージェントがプロンプトを読み取って 1 つのシェル コマンドを実行するだけで済むため、どのツールでも機能します。
自分で名前を付けたセッションでは、その名前が維持されます。最初のプロンプトでは、エージェントとマネージャーの名前変更は、要求に応じて後で使用できることが示されているだけで、エージェントに今すぐ名前を変更するように指示するものではありません。後からエージェントにセッション名の変更を依頼することも、セッション内のシェルから自分自身で Agent-Manager rename を実行することもできます。
レビュー中のリポジトリの宣言
セッションの作業ディレクトリは、多くのリポジトリを保持する包括的なフォルダーであることが多いため、レビューではエージェントがどのリポジトリを意味するのかを推測することしかできません。

どのリポジトリで作業しているかを知っているエージェントは、セッション内のシェルから Agent-manager review-repo <path> を実行することで、そのことを知ることができます。このサブコマンドは、パスが git リポジトリである (またはその中に存在する) ことを確認し、それをリポジトリ ルートに解決して、セッションごとのファイルにドロップします。マネージャーは次回の投票でそのリポジトリを選択し、次にリポジトリを開いたときにそのリポジトリでレビューが開きます。 git リポジトリ内にないパスは拒否されるため、宣言は推測ではなく常に事実になります。
エージェントは、ワークツリー内から Agent-manager review-base <ref> を実行することで、そのブランチの差分を宣言することもできます。ref はそのリポジトリで検証され、セッションおよびリポジトリごとに保存され、それ以降、「vs target」スコープでそれが使用されます。 Agent-Manager review-base --clear は自動検出に戻ります。レビュー中のエラーとしてサーフェスの解決を停止する保存された ref と、B がターゲット ピッカー (リポジトリのブランチと自動エントリ) を開いて手動で設定またはクリアします。
通常、エージェントは git ワークツリー (ワークツリーごとに 1 つのブランチ) で動作し、それらのワークツリーはディスク上のどこにでも存在できます。ワークツリーのルートである宣言されたパスは、それが存在する場所であればどこでも受け入れられるため、1 回の review-repo 呼び出しでリポジトリとレビュー対象のブランチの両方に名前を付けます。 Review は固定の順序でターゲットを解決します。つまり、マネージャーが実行されている間は r または b で手動で選択したリポジトリが優先され、次にエージェントが宣言したリポジトリ、次にランキングになります (最初にダーティ作業ツリー、次に最新のコミット)。選択または宣言されたパスが git リポジトリではなくなると、レビューのステータス行にその旨が表示され、r が正しいパスを選択します。
MCP: エージェントがこれらのコマンドを検出する方法
マネージャーが生成または復活するすべてのセッションには、エージェント マネージャー MCP サーバーが含まれるため、MCP 対応エージェントは、いつ呼び出すかを指示する説明付きのネイティブ ツールとして rename 、 review_repo 、 review_base を参照します。

それぞれ: プロンプト挿入やプロジェクトごとのセットアップはありません。サーバーは同じバイナリ (agent-manager mcp、stdio) 内に存在し、その環境を通じて呼び出しセッションを識別します。
登録はツールごとに行われます。組み込みの claude、codex、opencode、および grok ツールは自動的に登録されます。claude は生成された --mcp-config ファイルを取得し、codex は -c mcp_servers... オーバーライドを取得し、opencode は OPENCODE_CONFIG マージ ファイルを取得し、grok は初回起動時に 1 回限りの grok mcp add --scope ユーザー エントリを取得します。カスタム ツールは、構成セクションで mcp = "<style>" を使用してオプトインするか、 mcp = "none" を使用してオプトアウトします。 CLI サブコマンドは、MCP かどうかに関係なく、どこでも機能し続けます。
セッションで Ctrl+R を押すと、そのリポジトリの全画面レビューが開きます。左側には +/- カウントが付いた変更されたファイル、右側には構文が強調表示され、変更された行が色付けされたファイル全体が表示されるため、すべての編集がフルコンテキストで読み取られます。矢印キーと ctrl+d / ctrl+u でファイルをスクロール、 g / G で上下にジャンプ、J / K (または tab / shift+tab ) でファイルを切り替え、n / N で変更間をジャンプ、u で統合と並列を切り替え、 s でスコープ (未コミット、対ターゲット、最後のコミット、ステージング) を循環、スペースでレビュー済みのファイルをマークします。作業ディレクトリに複数のリポジトリが保持されている場合、r はフィルタリングするために入力したピッカーを開き、b は現在のリポジトリのワークツリーをブランチ名ごとにリストするため、キーを 1 回押すだけで別のブランチを確認できます。 B は、「vs target」スコープが比較するターゲット (merge-into ブランチ) を選択します。ヘッダー内の変更可能な値にはそれぞれ独自のキーが付けられているため、スコープ、レイアウト、リポジトリ、およびターゲットのピルは、一目で s 、 u 、 r 、 B の凡例として読み取れます。エージェントが編集を続けると、差分が更新されます。
コメントを書き込むには、行で c を押します。 C は、すべてのコメントを 1 つのレビュー プロンプトにまとめてエージェントのペインに直接送信するため、ユーザーがコメントを見ている間にエージェントがメモに対処し始めます。

彼は差分を更新しました。 esc でレビューを閉じます。
グループは、無制限の深さのツリーを形成するパス ( backend/api/auth ) です。セッションは、ルートを含む任意のノードに存在できます。 g でインラインでサブグループを作成し、 K / J (または SHIFT+↑↓ ; 順序は維持) でグループとセッションの両方を並べ替え、 f でサブツリーを折りたたみ、 e で空のグループを視覚的に非表示または復元し、 r でグループの名前とデフォルトのパスを編集します。セッション上で、r を押すとツールの名前が変更され、タブでツールが循環します (ステータス ルールと復活は新しいツールに従います。ペイン内の 1 つのエージェントを終了して別のエージェントを開始する場合に便利です)。
各セッションの tmux ペインがポーリングされ (デフォルトでは 2 秒ごと)、ステータスが取得されます。
各行にはそのステータスとツールがインラインで含まれており、折りたたまれたグループはステータスごとのカウントを保持するため、折りたたまれたサブツリーでも、何かが必要かどうかがわかります。セッションを選択すると、右側のペインの末尾が表示されます。これにより、待機中のエージェントからの実際の質問が添付されずに届くようになります。
検出は、ツールごとの正規表現ルールを表示ペインと照合し、最新のターンを分析して終了か待機かを判断し、ストリーミング出力 (ポーリング間で変化するコンテンツ) を動作中として扱います。ターン概要ラインが表示されずに終了したターンでも解決されます。作業ペインが静止すると、ターン カウントが表示されます。

[切り捨てられた]

## Original Extract

Terminal UI to manage AI coding-agent sessions (Claude Code, OpenCode, Codex, Grok Build) in tmux: live status, group tree, live pane preview, resource gauges. - YoanWai/agent-manager

GitHub - YoanWai/agent-manager: Terminal UI to manage AI coding-agent sessions (Claude Code, OpenCode, Codex, Grok Build) in tmux: live status, group tree, live pane preview, resource gauges. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
YoanWai
/
agent-manager
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
193 Commits 193 Commits .github .github docs docs internal internal .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum main.go main.go main_test.go main_test.go View all files Repository files navigation
Run every AI coding agent from one terminal. Claude Code, Codex, OpenCode, and Grok run side by side, each in its own tmux session, so they keep working after you quit the manager.
Instead of hunting through terminal tabs to see which agent is done and which is stuck, every session shows up in one list with live status, grouped into a project tree you can fold and reorder. You answer any of them without attaching: space sends a prompt straight into a session's pane, or spawns a new agent in the selected group. A dead session revives where it left off with v . And ctrl+r opens a full-file diff of what an agent changed, syntax-highlighted, where a comment you leave on a line goes straight back to the agent's pane.
Status detection currently supports Claude Code , OpenCode , Codex , and Grok Build out of the box. Any other CLI tool can run as a session; add a [tools.<name>] block with status rules to get live status for it (see Configuration ).
brew install yoanwai/tap/agent-manager
Installs tmux with it if missing.
go install github.com/YoanWai/agent-manager@latest
Requires Go 1.26+ and tmux; installs to $(go env GOPATH)/bin .
Download from Releases (macOS and Linux, amd64/arm64).
Run inside WSL2 : agent-manager lives on tmux, which is a Linux/macOS tool. In a WSL shell, install with Homebrew or grab the Linux binary from Releases.
The manager checks GitHub Releases once a day and shows a ↑ vX.Y.Z available badge in the header when a newer version is out. Pull it in the way you installed:
brew upgrade yoanwai/tap/agent-manager # Homebrew
go install github.com/YoanWai/agent-manager@latest # Go
Usage
agent-manager
Sessions run inside tmux ( am_* namespace), so they survive the manager quitting. Inside a session, Ctrl+Q detaches back to the manager. agent-manager --version prints the version.
Key
Action
n
New session (name, tool, directory, optional starting prompt, group picker)
g
New group (name, parent, default path)
enter
Attach session / fold group
ctrl+q
Inside a session: back to the manager
K / J (or shift+↑ / shift+↓ )
Reorder session or group among its visible siblings
m
Move session to another group
r
Rename session / edit tool; edit group name and default path
v
Revive a dead session ( revive_command , e.g. claude --continue , resumes the conversation)
a / u
Archive / restore a session, or a group and its entire subtree
d
Delete session, or a group + its entire subtree
space
Quick prompt: answer the selected session, or spawn an agent in the selected group
ctrl+r
Review the selected session's changes: full-screen whole-file diffs, line comments sent to the agent
f
Fold / unfold group
s
Settings (quick-spawn tool, theme, review layout)
t
Toggle archived view
e
Hide / show empty groups
/
Search
?
Help
q
Quit (sessions keep running)
Quick prompt
Press space to dock a prompt bar at the bottom of the sidebar. The target follows the cursor while the bar is open ( ↑↓ still navigate):
On a session row, enter sends the typed text straight into the session's pane, so the agent gets it as a user message without you attaching. The bar stays open and clears, ready for the next answer.
On a group row, enter spawns a new agent in that group with the prompt embedded, using the group's default path. The spawn tool starts at the Settings ( s ) default and tab cycles it (claude ↔ opencode ↔ any configured tool); the footer shows the current pick. The agent starts working on the prompt immediately.
esc closes the bar. The new-session form's optional prompt field launches an agent the same way; tools whose CLI takes the prompt behind a flag declare it with prompt_flag (see Configuration ).
Sessions spawned without a custom name (every quick spawn, and the form with the name left blank) get a placeholder like claude-a1b2 , and their first prompt opens by asking the agent to run agent-manager rename "<name>" once with a short name for the broad feature of the session (not a single subtask). The directive also tells the agent not to rename again unless you ask. When the first prompt cannot carry the directive (a /slash command, or no prompt at all), the manager sends it as its own message once the tool's input box appears in the pane. The subcommand drops the name into a per-session file; the manager picks it up on the next poll and updates the sidebar row and the tmux status bar. This works with any tool, since it only needs the agent to read its prompt and run one shell command.
Sessions you name yourself keep that name: the first prompt only notes that agent-manager rename is available later if you ask, and does not instruct the agent to rename now. You can still ask an agent to rename its session later, or run agent-manager rename yourself from a shell inside the session.
Declaring the repo under review
A session's working directory is often an umbrella folder holding many repos, so review can only guess which one the agent means. An agent that knows which repo it is working in can say so by running agent-manager review-repo <path> from a shell inside its session. The subcommand checks that the path is (or sits inside) a git repo, resolves it to the repo root, and drops it into a per-session file; the manager picks it up on the next poll and review opens on that repo the next time you open it. A path that is not inside a git repo is rejected, so a declaration is always a fact rather than a guess.
An agent can also declare what its branch diffs against by running agent-manager review-base <ref> from inside its worktree: the ref is validated in that repo, stored per session and repo, and the "vs target" scope uses it from then on. agent-manager review-base --clear returns to automatic detection. A stored ref that stops resolving surfaces as an error in review, and B opens a target picker (the repo's branches plus an auto entry) to set or clear it by hand.
Agents usually work in git worktrees, one branch per worktree, and those worktrees can live anywhere on disk. A declared path that is a worktree root is accepted wherever it lives, so one review-repo call names both the repo and the branch under review. Review resolves its target in a fixed order: a repo you picked by hand with r or b wins for as long as the manager is running, then the agent's declared repo, then the ranking (dirty working trees first, then most recent commit). When the picked or declared path stops being a git repo, review says so in the status line and r is there to pick the right one.
MCP: how agents discover these commands
Every session the manager spawns or revives carries the agent-manager MCP server, so MCP-capable agents see rename , review_repo and review_base as native tools with descriptions telling them when to call each: no prompt injection, no per-project setup. The server lives in the same binary ( agent-manager mcp , stdio) and identifies the calling session through its environment.
Registration is per tool. The built-in claude, codex, opencode and grok tools register automatically: claude gets a generated --mcp-config file, codex gets -c mcp_servers... overrides, opencode gets an OPENCODE_CONFIG merge file, and grok gets a one-time grok mcp add --scope user entry on its first launch. A custom tool opts in with mcp = "<style>" in its config section, or out with mcp = "none" . The CLI subcommands keep working everywhere, MCP or not.
Press ctrl+r on a session to open a full-screen review of its repo: changed files with +/− counts on the left, the whole file on the right with syntax highlighting and changed lines tinted, so every edit reads in full context. Arrow keys and ctrl+d / ctrl+u scroll the file, g / G jump to top and bottom, J / K (or tab / shift+tab ) switch files, n / N jump between changes, u toggles unified and side-by-side, s cycles the scope (uncommitted, vs target, last commit, staged), and space marks a file reviewed. When the working directory holds several repos, r opens a picker you type to filter, and b lists the current repo's worktrees by branch name so you can review another branch with one keypress. B picks the target (the merge-into branch) the "vs target" scope compares against. Each changeable value in the header wears its own key, so the scope, layout, repo, and target pills read as s , u , r , B legends at a glance. The diff refreshes as the agent keeps editing.
Press c on a line to write a comment; C flattens every comment into one review prompt and sends it straight into the agent's pane, so the agent starts addressing your notes while you watch the diff update. esc closes the review.
Groups are paths ( backend/api/auth ) forming a tree of unlimited depth. Sessions can live at any node, including the root. Create subgroups inline with g , reorder both groups and sessions with K / J (or shift+↑↓ ; the order persists), fold a subtree with f , hide or restore empty groups visually with e , and edit a group's name and default path with r . On a session, r renames it and tab cycles the tool (status rules and revive follow the new tool; useful when you quit one agent in the pane and start another).
Each session's tmux pane is polled (default every 2s) to derive a status:
Each row carries its status and tool inline, and a folded group keeps a count per status so a collapsed subtree still tells you whether anything needs you. Selecting a session shows the tail of its pane on the right, which is how a waiting agent's actual question reaches you without attaching.
Detection matches per-tool regex rules against the visible pane, analyzes the newest turn to tell finished from waiting , and treats streaming output (content changing between polls) as working . A turn that ends without any turn-summary line still resolves: when a working pane goes quiet, the turn count

[truncated]
