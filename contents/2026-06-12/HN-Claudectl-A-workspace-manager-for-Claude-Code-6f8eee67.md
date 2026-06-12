---
source: "https://github.com/babarmuhammad/claudectl"
hn_url: "https://news.ycombinator.com/item?id=48497618"
title: "HN: Claudectl – A workspace manager for Claude Code"
article_title: "GitHub - babarmuhammad/claudectl · GitHub"
author: "babarmuhammad1"
captured_at: "2026-06-12T01:02:35Z"
capture_tool: "hn-digest"
hn_id: 48497618
score: 1
comments: 0
posted_at: "2026-06-11T23:03:51Z"
tags:
  - hacker-news
  - translated
---

# HN: Claudectl – A workspace manager for Claude Code

- HN: [48497618](https://news.ycombinator.com/item?id=48497618)
- Source: [github.com](https://github.com/babarmuhammad/claudectl)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T23:03:51Z

## Translation

タイトル: HN: Claudectl – クロード コードのワークスペース マネージャー
記事タイトル: GitHub - babarmuhammad/claudectl · GitHub
説明: GitHub でアカウントを作成して、babarmuhammad/claudectl の開発に貢献します。

記事本文:
GitHub - babarmuhammad/claudectl · GitHub
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
ババルムハマド
/
クラウディクトル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクション メニューを開く 折りたたむ

ファイルとファイル
8 コミット 8 コミット .github/ workflows .github/ workflows claude_sessions claude_sessions テスト テスト .gitattributes .gitattributes .gitignore .gitignore リポジトリ cmd.bat を開く リポジトリ cmd.bat を開く README.md README.md クロード フォルダ.ico クロード フォルダ.ico claude-sessions.py claude-sessions.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code 用の Windows ワークスペース マネージャー — 高速ターミナル UI でのプロジェクト メモリ、MCP 認識、およびマルチプロジェクト ワークフロー。
Claude Code はあなたの作品をチャットのコレクションとして扱います。 claudectl は各プロジェクトを永続的なワークスペースとして扱います。セッションは参照可能および検索可能であり、プロジェクトのコンテキストは維持された CLAUDE.md ファイル内に存在し、MCP サーバーは一目で確認でき、すべての起動はプロジェクトごとに構成されます。プロジェクトを切り替えると、エージェントの記憶を失うような感覚はなくなります。
セッション ブラウザ — すべての Claude Code プロジェクトとセッション、最新順に並べ替え
クイック再開 — メイン画面の★/☆ ショートカットは、すべてのプロジェクトの最近のセッションに直接戻ります
検索 - ライブセッションを名前またはプレビューでフィルタリングするために入力します
名前変更 / 削除 / フォーク — R / D / F キーを使用して個々のセッションを管理します
Scaffold CLAUDE.md (C) — git リポジトリ、最近のコミット、README、および以前のセッションのトピックから機械的にプロジェクト コンテキストを構築します
AI CLAUDE.md 生成 (A) — Claude 自体がコードベースを詳細に分析し、包括的な CLAUDE.md を作成または更新します。何かが書かれる前にレビューする
システム プロンプト (S) — 起動のたびに挿入されるプロジェクトごとのシステム プロンプトを AI で生成または手動で編集します
MCP ステータス — 接続されているサーバーが起動時にフッターに表示されます
MCP ドキュメント — MCP サーバーのツールを分析し、ドキュメントをグローバル ~/.claude/CLAUDE.md に書き込むことで、クロードがセッションごとにドキュメントを認識できるようになります。
工数/モデルセレクター

— 各起動前に推論の労力とモデルのオーバーライドを選択します。プロジェクトごとに記憶される最後の選択
追加の PATH エントリ — クロードの環境に注入されるプロジェクトごとの PATH ディレクトリ
AI セッション タイトル - マニュアル名のないセッションには、AI によって生成されたトランスクリプトのタイトルが表示されます。
設定画面 (⚙) — エディター、claude.exe パス、およびデフォルトの起動オプション ( ~/.claude/claudectl.json ) を構成します。
ヘルプ画面 — ? を押します。キーボードのリファレンスとして
Claude Code CLI がインストールされている (%USERPROFILE%\.local\bin\claude.exe または PATH で自動検出され、設定で上書き可能)
任意のテキスト エディタ — Notepad++ / VS Code は自動検出され、Windows メモ帳がフォールバックになります (設定で上書き可能)
pipx インストール claudectl
クラウディクトル
以上です。claudectl はセッション ブラウザを起動し、Claude を直接起動します。
git クローン https://github.com/babarmuhammad/claudectl.git
cd クローディクトル
「Open Repo cmd.bat」をダブルクリックします (またはターミナルから実行します)。
[リポジトリ cmd.bat を開く] → [送信] → [デスクトップ (ショートカットの作成)] を右クリックします。
オプション: タスクバーにピン留めする (Windows 11)
Windows 11 では .bat ショートカットを直接固定することはできません。ショートカットは cmd.exe を指す必要があります。 PowerShell でリポジトリ フォルダーからこれを 1 回実行します。
$shell = 新しいオブジェクト - ComObject WScript.Shell
$lnk = $shell .CreateShortcut ( " $ env: USERPROFILE \Desktop\Open Repo Claude.lnk " )
$lnk .TargetPath = " C:\Windows\System32\cmd.exe "
$lnk .Arguments = " /c `" $PWD \Open Repo cmd.bat `" "
$lnk .WorkingDirectory = " $PWD "
$lnk .IconLocation = " $PWD \claude フォルダー.ico, 0 "
$lnk .Save ()
次に、デスクトップのショートカットを右クリック→タスクバーにピン留めします。
起動時に、claudectl は、Claude Code がこれまでに開いたすべてのプロジェクトを、最近使用した順に並べて表示します。
クイック再開アイテムは上部に表示されます (★ = 最新のセッション、☆ = 古いセッション)
他のすべてのプロジェクトがこれに続き、最近の順に並べ替えられます
の

バックグラウンド チェックが完了すると、MCP ステータス フッターに接続された MCP サーバーが表示されます
下部にある ⚙ Global CLAUDE.md / MCP Analysis を選択して、グローバル コンテキスト メニューを開きます。
これらは、すべてのプロジェクトで最近使用された 5 つのセッションです。 1 つを選択すると、プロジェクトのセッション リストに移動することなく、そのセッションがすぐに再開されます。 ★ は単一の最新セッションを示します。 ☆は古いエントリを示します。
⚙ グローバル CLAUDE.md / MCP 分析
接続されているすべての MCP サーバーをリストするサブメニューを開きます。 MCP のツール/リスト エンドポイントを呼び出し、結果をマークダウンとしてフォーマットするプロンプトを表示してクロードを実行するサーバーを選択します。出力はサーバーごとのセンチネル ブロック内の ~/.claude/CLAUDE.md に書き込まれるため、後続の実行時にきれいに更新できます。このメニューから Notepad++ でグローバル CLAUDE.md を直接開くこともできます。
キー
アクション
↑ / ↓
ナビゲート
入る
プロジェクトの選択
ESC
終了
セッション画面（プロジェクトのセッションリスト）
キー
アクション
↑ / ↓
ナビゲート
入る
選択・確認
ESC
戻る / キャンセル (アクティブな場合は最初にフィルターをクリアします)
R
セッション名の変更
D
セッションを削除します (確認のプロンプトが表示されます)
F
フォークセッション
C
Scaffold CLAUDE.md (git + セッション)
あ
AI 生成 CLAUDE.md (クロード CLI)
S
システムプロンプトの編集/生成
P
追加の PATH エントリを管理する
?
ヘルプ/キーボードリファレンス
バックスペース
最後のフィルター文字を削除
テキストを入力します
ライブセッションを名前またはプレビューでフィルタリングします
起動オプション画面
キー
アクション
↑ / ↓
工数フィールドとモデルフィールドを切り替える
← / →
選択したフィールドの値を循環します
入る
選択したオプションで起動する
ESC
メインメニューに戻る (起動なし)
プロジェクトごとのファイル
各プロジェクトは ~/.claude/projects/<encoded-name>/ にフォルダーを取得します。 claudectl は、そこでいくつかのファイルを読み書きします。
C — 足場 (高速、機械的)
プロジェクトとリンクの深さ 2 レベルまでの Git リポジトリが見つかります

d 追加のパス
各リポジトリからの最後の 7 つのコミット ( git log --oneline -7 )
各リポジトリの README の最初の 15 行
すべてのセッションのトピック (蓄積され、決して破棄されません)
既存のファイルでは、<!-- AUTOGEN:START -->…<!-- AUTOGEN:END --> ブロックと <!-- SESSIONS:START -->…<!-- SESSIONS:END --> ブロックのみが置き換えられます。これらのブロックの外側にあるものはすべて正確に保存されます。
A — AI 分析 (低速、包括的)
完全なディレクトリ ツリー、git 履歴、README、追加のパス、セッション履歴を含む豊富なプロンプトを使用して claude.exe -p を実行します。クロードは CLAUDE.md 全体を作成します。ファイルが書き込まれる前に、ポケットベルでそれを確認し、承認または拒否します。
既存のファイルでは、現在のコンテンツがグラウンド トゥルースとして渡され、明らかに変更された事実のみを更新するように指示されます。
生成後、<!-- AUTOGEN:START/END --> ブロックと <!-- SESSIONS:START/END --> ブロックが機械的に挿入され、<!-- AI:ANALYZED --> が 2 行目に挿入されるため、今後の実行はフレッシュ モードではなく更新モードになります。
~/.claude/CLAUDE.md は、すべてのプロジェクトのすべてのセッションでクロード コードによってロードされます。 claudectl はこれを使用して MCP ツールのドキュメントを保存します。
各 MCP サーバーは、センチネルで区切られた独自のセクションを取得します。
<!-- MCP:Notion:START -->
## MCP: 概念
…ツールリスト…
<!-- MCP:Notion:END -->
同じサーバーに対して分析を再実行すると、そのセクションのみが更新されます。他のコンテンツは変更されていません。
アクセス方法: メイン画面 → ⚙ Global CLAUDE.md / MCP Analysis
症状
修正
起動時の「claude.exe が見つかりません」画面
Claude Code をインストールするか、⚙ 設定でパスを設定します。
生成されたファイルがエディターで開かない
⚙ 設定でエディターのパスを設定します (Notepad++、VS Code を自動検出し、Notepad にフォールバックします)。
ウィンドウがエラーですぐに閉じてしまう
%TEMP%\claudectl_crash.log を確認します。クラッシュ ハンドラーはそこにトレースバックを書き込みます。
プロジェクト

リストにありません
プロジェクト フォルダーが移動または削除されたか、パスをデコードできません。以下のセッション エンコードを参照してください。
設定場所
~/.claude/claudectl.json — 手動で編集することも、削除してリセットすることも安全です
セッションのエンコーディング
Claude コードは、パス区切り文字を -- に、特定の特殊文字を - に置き換えることにより、プロジェクト パスを ~/.claude/projects/ の下のフォルダー名としてエンコードします。たとえば:
D:\Projects\my-app → D--Projects-my-app
paths.py の claudectl の find_actual_path() は、ファイルシステムを調べてエンコードされたコンポーネントを照合し、ディレクトリ名の _ 、 + 、 - 、 # などのエッジ ケースを処理することでこれを逆にします。
.\クラウデクト\
§── claude-sessions.py # ランチャースタブ + クラッシュハンドラー
§── リポジトリを開く cmd.bat # バットランチャー
§── README.md
§── .gitignore
└── claude_sessions\ # パッケージ
§── __init__.py
§── config.py # 定数、パス、センチネル文字列
§── paths.py # encode_component, find_actual_path
§── session.py # セッション解析、永続化ヘルパー
§── ui.py # text_input、menu、paths_menu、launch_options_menu
§── claude_md.py # scaffold_claude_md、ai_scaffold_claude_md、ヘルパー
§── system_prompt.py # edit_system_prompt, ai_generate_system_prompt
§── session_menu.py # session_menu
§── mcp.py # MCP バックグラウンド ポーリング、global_claude_md_menu
└── main.py # run() — プロジェクトの検出とメインループ
について
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

Contribute to babarmuhammad/claudectl development by creating an account on GitHub.

GitHub - babarmuhammad/claudectl · GitHub
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
babarmuhammad
/
claudectl
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github/ workflows .github/ workflows claude_sessions claude_sessions tests tests .gitattributes .gitattributes .gitignore .gitignore Open Repo cmd.bat Open Repo cmd.bat README.md README.md claude folder.ico claude folder.ico claude-sessions.py claude-sessions.py pyproject.toml pyproject.toml View all files Repository files navigation
A Windows workspace manager for Claude Code — project memory, MCP awareness, and multi-project workflows in a fast terminal UI.
Claude Code treats your work as a collection of chats. claudectl treats each project as a persistent workspace : sessions stay browsable and searchable, project context lives in maintained CLAUDE.md files, MCP servers are visible at a glance, and every launch is configured per project. Switching projects stops feeling like losing the agent's memory.
Session browser — every Claude Code project and session, sorted by recency
Quick-resume — ★/☆ shortcuts on the main screen jump straight back into recent sessions across all projects
Search — type to filter sessions live by name or preview
Rename / Delete / Fork — manage individual sessions with R / D / F keys
Scaffold CLAUDE.md (C) — build project context mechanically from git repos, recent commits, READMEs, and prior session topics
AI CLAUDE.md generation (A) — Claude itself deep-analyzes the codebase and writes or updates a comprehensive CLAUDE.md; you review before anything is written
System prompts (S) — AI-generate or hand-edit a per-project system prompt injected on every launch
MCP status — connected servers shown in the footer on startup
MCP documentation — analyze any MCP server's tools and write the docs into the global ~/.claude/CLAUDE.md so Claude knows them in every session
Effort / model selector — choose reasoning effort and model override before each launch; last choice remembered per project
Extra PATH entries — per-project PATH dirs injected into Claude's environment
AI session titles — sessions without a manual name show their AI-generated transcript title
Settings screen (⚙) — configure editor, claude.exe path, and default launch options ( ~/.claude/claudectl.json )
Help screen — press ? for a keyboard reference
Claude Code CLI installed (auto-detected at %USERPROFILE%\.local\bin\claude.exe or on PATH; overridable in Settings)
Any text editor — Notepad++ / VS Code are auto-detected, Windows Notepad is the fallback (overridable in Settings)
pipx install claudectl
claudectl
That's it — claudectl launches the session browser and starts Claude directly.
git clone https://github.com/babarmuhammad/claudectl.git
cd claudectl
Double-click Open Repo cmd.bat (or run it from a terminal).
Right-click Open Repo cmd.bat → Send to → Desktop (create shortcut) .
Optional: Pin to taskbar (Windows 11)
Windows 11 can't pin .bat shortcuts directly — the shortcut must point to cmd.exe . Run this once in PowerShell from the repo folder:
$shell = New-Object - ComObject WScript.Shell
$lnk = $shell .CreateShortcut ( " $ env: USERPROFILE \Desktop\Open Repo Claude.lnk " )
$lnk .TargetPath = " C:\Windows\System32\cmd.exe "
$lnk .Arguments = " /c `" $PWD \Open Repo cmd.bat `" "
$lnk .WorkingDirectory = " $PWD "
$lnk .IconLocation = " $PWD \claude folder.ico, 0 "
$lnk .Save ()
Then right-click the Desktop shortcut → Pin to taskbar .
On launch, claudectl shows all projects Claude Code has ever opened, sorted by most recently used.
Quick-resume items appear at the top (★ = most recent session, ☆ = older sessions)
All other projects follow, sorted by recency
The MCP status footer shows connected MCP servers once the background check completes
Select ⚙ Global CLAUDE.md / MCP Analysis at the bottom to open the global context menu
These are the 5 most recently used sessions across all projects. Selecting one immediately resumes that exact session without navigating into the project's session list. ★ marks the single most recent session; ☆ marks older entries.
⚙ Global CLAUDE.md / MCP Analysis
Opens a sub-menu listing all connected MCP servers. Select any server to run Claude with a prompt that calls the MCP's tools/list endpoint and formats the result as markdown. The output is written into ~/.claude/CLAUDE.md inside a per-server sentinel block so it can be cleanly updated on subsequent runs. You can also open the global CLAUDE.md directly in Notepad++ from this menu.
Key
Action
↑ / ↓
Navigate
ENTER
Select project
ESC
Exit
Sessions screen (session list for a project)
Key
Action
↑ / ↓
Navigate
ENTER
Select / confirm
ESC
Back / cancel (clears filter first if active)
R
Rename session
D
Delete session (prompts for confirmation)
F
Fork session
C
Scaffold CLAUDE.md (git + sessions)
A
AI-generate CLAUDE.md (Claude CLI)
S
Edit / generate system prompt
P
Manage extra PATH entries
?
Help / keyboard reference
BACKSPACE
Delete last filter character
Type text
Filter sessions live by name or preview
Launch options screen
Key
Action
↑ / ↓
Switch between Effort and Model fields
← / →
Cycle through values for the selected field
ENTER
Launch with selected options
ESC
Back to main menu (no launch)
Per-project files
Each project gets a folder at ~/.claude/projects/<encoded-name>/ . claudectl reads and writes several files there:
C — Scaffold (fast, mechanical)
Git repos found up to 2 levels deep in the project and any linked extra paths
Last 7 commits from each repo ( git log --oneline -7 )
First 15 lines of each repo's README
All session topics (accumulated, never discarded)
On an existing file, only the <!-- AUTOGEN:START -->…<!-- AUTOGEN:END --> and <!-- SESSIONS:START -->…<!-- SESSIONS:END --> blocks are replaced. Everything outside those blocks is preserved exactly.
A — AI analyze (slower, comprehensive)
Runs claude.exe -p with a rich prompt containing the full directory tree, git history, READMEs, extra paths, and session history. Claude writes the entire CLAUDE.md. You review it in a pager and approve or reject before any file is written.
On an existing file, the current content is passed as ground truth with instructions to update only facts that have clearly changed.
After generation the <!-- AUTOGEN:START/END --> and <!-- SESSIONS:START/END --> blocks are injected mechanically, and <!-- AI:ANALYZED --> is inserted on line 2 so future runs enter update mode rather than fresh mode.
~/.claude/CLAUDE.md is loaded by Claude Code in every session across all projects. claudectl uses it to store MCP tool documentation:
Each MCP server gets its own sentinel-delimited section:
<!-- MCP:Notion:START -->
## MCP: Notion
… tool listing …
<!-- MCP:Notion:END -->
Re-running the analysis for the same server updates only that section; other content is untouched.
Access via: main screen → ⚙ Global CLAUDE.md / MCP Analysis
Symptom
Fix
"claude.exe not found" screen on startup
Install Claude Code , or set the path in ⚙ Settings
Generated files don't open in an editor
Set your editor path in ⚙ Settings (auto-detects Notepad++, VS Code, falls back to Notepad)
Window closes instantly with an error
Check %TEMP%\claudectl_crash.log — the crash handler writes the traceback there
Projects missing from the list
The project folder was moved/deleted, or the path can't be decoded — see Session encoding below
Settings location
~/.claude/claudectl.json — safe to edit by hand or delete to reset
Session encoding
Claude Code encodes project paths as folder names under ~/.claude/projects/ by replacing path separators with -- and certain special characters with - . For example:
D:\Projects\my-app → D--Projects-my-app
claudectl's find_actual_path() in paths.py reverses this by walking the filesystem and matching encoded components, handling edge cases like _ , + , - , # in directory names.
.\claudectl\
├── claude-sessions.py # launcher stub + crash handler
├── Open Repo cmd.bat # bat launcher
├── README.md
├── .gitignore
└── claude_sessions\ # package
├── __init__.py
├── config.py # constants, paths, sentinel strings
├── paths.py # encode_component, find_actual_path
├── sessions.py # session parsing, persistence helpers
├── ui.py # text_input, menu, paths_menu, launch_options_menu
├── claude_md.py # scaffold_claude_md, ai_scaffold_claude_md, helpers
├── system_prompt.py # edit_system_prompt, ai_generate_system_prompt
├── session_menu.py # sessions_menu
├── mcp.py # MCP background poll, global_claude_md_menu
└── main.py # run() — project discovery and main loop
About
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
