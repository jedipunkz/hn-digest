---
source: "https://github.com/javaid-codes/rapunzel/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48560175"
title: "Show HN: Rapunzel – a tree-style tab terminal emulator for Codex Claude Gemini"
article_title: "GitHub - javaid-codes/rapunzel: Browser for agents. · GitHub"
author: "WasimBhai"
captured_at: "2026-06-16T19:18:31Z"
capture_tool: "hn-digest"
hn_id: 48560175
score: 2
comments: 0
posted_at: "2026-06-16T18:54:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rapunzel – a tree-style tab terminal emulator for Codex Claude Gemini

- HN: [48560175](https://news.ycombinator.com/item?id=48560175)
- Source: [github.com](https://github.com/javaid-codes/rapunzel/tree/main)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T18:54:01Z

## Translation

タイトル: Show HN: Rapunzel – Codex Claude Gemini 用のツリー スタイルのタブ ターミナル エミュレータ
記事のタイトル: GitHub - javaid-codes/rapunzel: エージェント用のブラウザ。 · GitHub
説明: エージェント用のブラウザ。 GitHub でアカウントを作成して、javaid-codes/rapunzel の開発に貢献してください。

記事本文:
GitHub - javaid-codes/rapunzel: エージェント用のブラウザ。 · GitHub
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
javaidコード
/
ラプンツェル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く

アクション メニュー フォルダーとファイル
11 コミット 11 コミット ラプンツェル ラプンツェル スクリプト スクリプト webui webui .gitignore .gitignore ARCHITECTURE_AND_PROGRESS.md ARCHITECTURE_AND_PROGRESS.md FEATURE_TRACKER.md FEATURE_TRACKER.md PROJECT_STATUS.md PROJECT_STATUS.md README.md README.md Rapunzel.command Rapunzel.command Rapunzel.ps1 Rapunzel.ps1 app.py app.py build_macos_app.sh build_macos_app.sh icon.png icon.png install.ps1 install.ps1 install.sh install.sh package-lock.json package-lock.json package.json package.json rapunzel_pyinstaller.spec rapunzel_pyinstaller.spec要件.txt 要件.txt ターミナルツリーアプリデザイン.md ターミナルツリーアプリデザイン.md test_models.py test_models.py test_session_close.py test_session_close.py test_session_platform.py test_session_platform.py test_state.py test_state.py test_store.py test_store.py test_terminal_screen.py test_terminal_screen.py ui.png ui.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Rapunzel はブラウザですが、エージェント向けです。
rapunzel を構築したのは、多くのエージェントがセッション間で実行されると、ターミナル タブのスケーリングが停止し、追跡が困難になるためです。 Firefox の Tree Style Tab Extension が気に入ったので、これを作成しました。私がこれを構築したのは、今は Chrome の時代ではなく、エージェントの時代だからです。できるだけシンプルにするために作りました。
macOS または Linux 上の新しいクローンから:
git clone < リポジトリ URL >
CD ラプンツェル
./install.sh
Windows PowerShell 上の新しいクローンから:
git clone < リポジトリ - URL >
CD ラプンツェル
.\install.ps1
インストール スクリプトは、ローカル仮想環境を作成し、Python 依存関係をインストールし、フロントエンド依存関係をインストールして、埋め込み Web UI を構築します。
./ラプンツェル.コマンド
Windows PowerShell の場合:
.\ラプンツェル.ps1
macOS では、Finder で Rapunzel.command をダブルクリックすることもできます。
ソース .venv/bin/activate
python3 アプリ.py

Rapunzel.command は、リポジトリローカル仮想環境が存在する場合はそれを優先するため、インストール後も起動パスは安定したままになります。
長時間実行されるエージェント セッションが多数あるとターミナル タブのスケーリングが停止する
階層ナビゲーションは、フラットなタブ ストリップよりも推論しやすいです。
目標は、巨大な IDE ではなく、エージェントの作業のためのシンプルなデスクトップ画面です。
ルートブランチと子ブランチを含むツリーベースのシェルセッション
1 つのアクティブなライブ端末会話面
名前変更、移動、折りたたみ、ブランチを閉じるアクション
~/.rapunzel/workspace.json へのワークスペースの永続化
起動時に保存されたワークスペース ツリーを復元する
ソースから実行するのではなく、クリック可能なアプリ バンドルが必要な場合:
ソース .venv/bin/activate
pip インストール pyinstaller
bash build_macos_app.sh
出力:
ビルドが完了したら、Finder から dist/Rapunzel.app を開くか、直接ダブルクリックします。
主要な製品パスは、埋め込まれた pywebview デスクトップ アプリです
シェルセッションは、macOS/Linux では POSIX PTY バックエンドを使用し、Windows では pywinpty を使用します。
古い Tk UI はフォールバック コードとしてまだ存在しますが、メインのアーキテクチャではなくなりました
Apple Silicon Mac では、Rapunzel シェル セッションが brew をラップするため、Homebrew コマンドはネイティブ arm64 経由で転送されます。
install.sh : macOS/Linux 上のソースからの初回実行セットアップ
install.ps1 : Windows 上のソースからの初回実行セットアップ
Rapunzel.command : macOS/Linux で通常のデスクトップで使用するためのソース ランチャー
Rapunzel.ps1 : Windows で通常のデスクトップで使用するためのソース ランチャー
rapunzel/ : アプリの状態、PTY ランタイム、ウィンドウ ホスト、およびフォールバック UI
webui/ : ツリーとターミナルの組み込みフロントエンド アセット
PROJECT_STATUS.md : 実装状況と次回の作業
ARCHITECTURE_AND_PROGRESS.md : アーキテクチャとデバッグの履歴
FEATURE_TRACKER.md : 実行中の機能ログ
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

Browser for agents. Contribute to javaid-codes/rapunzel development by creating an account on GitHub.

GitHub - javaid-codes/rapunzel: Browser for agents. · GitHub
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
javaid-codes
/
rapunzel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits rapunzel rapunzel scripts scripts webui webui .gitignore .gitignore ARCHITECTURE_AND_PROGRESS.md ARCHITECTURE_AND_PROGRESS.md FEATURE_TRACKER.md FEATURE_TRACKER.md PROJECT_STATUS.md PROJECT_STATUS.md README.md README.md Rapunzel.command Rapunzel.command Rapunzel.ps1 Rapunzel.ps1 app.py app.py build_macos_app.sh build_macos_app.sh icon.png icon.png install.ps1 install.ps1 install.sh install.sh package-lock.json package-lock.json package.json package.json rapunzel_pyinstaller.spec rapunzel_pyinstaller.spec requirements.txt requirements.txt terminal-tree-app-design.md terminal-tree-app-design.md test_models.py test_models.py test_session_close.py test_session_close.py test_session_platform.py test_session_platform.py test_state.py test_state.py test_store.py test_store.py test_terminal_screen.py test_terminal_screen.py ui.png ui.png View all files Repository files navigation
Rapunzel is a browser but for agents.
I built rapunzel because terminal tabs stop scaling once many agents are running across sessions and they become difficult to track. I built it because I love the Tree Style Tab Extension in Firefox. I built it because, it is the time for Chrome, but for agents. I made it, to make it as simple as possible and nothing more.
From a fresh clone on macOS or Linux:
git clone < repo-url >
cd rapunzel
./install.sh
From a fresh clone on Windows PowerShell:
git clone < repo - url >
cd rapunzel
.\install.ps1
The install script creates a local virtual environment, installs the Python dependencies, installs the frontend dependencies, and builds the embedded web UI.
./Rapunzel.command
On Windows PowerShell:
.\Rapunzel.ps1
On macOS, you can also double-click Rapunzel.command in Finder.
source .venv/bin/activate
python3 app.py
Rapunzel.command prefers the repo-local virtual environment when it exists, so the launch path stays stable after install.
terminal tabs stop scaling once you have many long-running agent sessions
hierarchical navigation is easier to reason about than a flat tab strip
the goal is a simple desktop surface for agent work, not a giant IDE
tree-based shell sessions with root and child branches
one active live terminal conversation surface
rename, move, collapse, and close branch actions
workspace persistence to ~/.rapunzel/workspace.json
restore of the saved workspace tree on launch
If you want a clickable app bundle instead of running from source:
source .venv/bin/activate
pip install pyinstaller
bash build_macos_app.sh
Output:
After the build finishes, open dist/Rapunzel.app from Finder or double-click it directly.
the primary product path is the embedded pywebview desktop app
shell sessions use the POSIX PTY backend on macOS/Linux and pywinpty on Windows
the older Tk UI still exists as fallback code, but it is no longer the main architecture
on Apple Silicon Macs, Rapunzel shell sessions wrap brew so Homebrew commands are forwarded through native arm64
install.sh : first-run setup from source on macOS/Linux
install.ps1 : first-run setup from source on Windows
Rapunzel.command : source launcher for normal desktop use on macOS/Linux
Rapunzel.ps1 : source launcher for normal desktop use on Windows
rapunzel/ : app state, PTY runtime, window host, and fallback UI
webui/ : embedded frontend assets for the tree and terminal
PROJECT_STATUS.md : implementation status and next work
ARCHITECTURE_AND_PROGRESS.md : architecture and debugging history
FEATURE_TRACKER.md : running feature log
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
