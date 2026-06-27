---
source: "https://github.com/shirozuki/claude-cli"
hn_url: "https://news.ycombinator.com/item?id=48702113"
title: "Show HN: Claude-CLI – Run Claude Code in a throwaway Docker container"
article_title: "GitHub - shirozuki/claude-cli · GitHub"
author: "shirozuki"
captured_at: "2026-06-27T22:21:38Z"
capture_tool: "hn-digest"
hn_id: 48702113
score: 1
comments: 0
posted_at: "2026-06-27T21:51:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude-CLI – Run Claude Code in a throwaway Docker container

- HN: [48702113](https://news.ycombinator.com/item?id=48702113)
- Source: [github.com](https://github.com/shirozuki/claude-cli)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T21:51:56Z

## Translation

タイトル: HN を表示: Claude-CLI – 使い捨て Docker コンテナでクロード コードを実行
記事タイトル: GitHub - hirozuki/claude-cli · GitHub
説明: GitHub でアカウントを作成して、hirotsuki/claude-cli の開発に貢献します。

記事本文:
GitHub - シロズキ/claude-cli · GitHub
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
城月
/
クロード・クリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
14 コミット 14 コミット テスト testing .gitignore .gitignore ライセンス ライセンス README.md README.md claude-cli.sh claude-cli.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Docker コンテナ内で Claude コードを実行する軽量のシェル ラッパー。セッション間で Claude 構成を永続化しながら、ホスト システムをクリーンな状態に保ちます。
dmenu — モード選択プロンプト ( -f を使用する場合は不要)
fzf — ディレクトリピッカー。 -f によるモード選択にも使用されます
現在のユーザーは docker グループのメンバーである必要があります
git clone https://github.com/hirozuki/claude-cli
chmod +x クロード-cli/クロード-cli.sh
ln -s " $PWD /claude-cli/claude-cli.sh " ~ /.local/bin/claude-cli
使用法
クロード-cli [-b] [-c] [-f] [-h] [DIR...]
オプション
説明
-b
claude-cli:latest Docker イメージをビルド (またはリビルド) します。
-c
すべての claude-cli コンテナを削除します。オプションで画像を削除します
-f
モード選択プロンプトには dmenu の代わりに fzf を使用してください
-h
ヘルプを表示する
オプションや引数を指定せずに実行すると、モードを選択するための dmenu プロンプトが起動します。代わりに fzf を使用するには -f を渡します。dmenu のないセットアップやターミナル ピッカーを使用したい場合に便利です。
引数としてディレクトリを渡す
コマンドラインでディレクトリを直接渡すことで、メニューとピッカーを完全にスキップできます。
claude-cli ~ /projects/app # 単一ディレクトリモード
claude-cli ~ /projects/app ~ /shared # マルチディレクトリ モード (最初の引数は作業ディレクトリです)
1 つの引数は、そのディレクトリで単一ディレクトリ モードを実行します。
2 つ以上の引数はマルチディレクトリ モードを実行し、すべてのディレクトリをマウントし、最初の引数を作業ディレクトリとして使用します。
相対パスは現在の作業ディレクトリに対して解決されるため、 claude-cli となります。 $PWD をマウントします。この方法でディレクトリを渡す場合、dmenu と fzf は必要ありません。
プロジェクト ディレクトリをマウントせずに Claude を起動します。一般的な質問、簡単なタスク、

またはローカル ファイルを公開せずにクロードの機能を探索することもできます。
現在の作業ディレクトリ $PWD を直接マウントします。
fzf ピッカーを開いて、ホーム ツリーから 1 つのディレクトリを選択します (最大 3 レベルの深さ)。選択したディレクトリがマウントされ、コンテナ内の作業ディレクトリとして設定されます。
fzf を複数選択モードで開き、複数のディレクトリを選択し (Tab を使用して選択します)、そのうちの 1 つを作業ディレクトリとして指定するように求められます。選択したすべてのディレクトリが同時にマウントされます。複数のリポジトリにまたがって作業する場合や、共通の構成ディレクトリを共有する場合に便利です。
Claude の構成 ( ~/.claude/ および ~/.claude.json ) は実行のたびにコンテナーにバインドマウントされるため、アカウント、設定、セッション履歴はコンテナーの再起動後も残ります。
マウント ソースは次の順序で解決されます。
$XDG_CONFIG_HOME/claude/ — $XDG_CONFIG_HOME が設定されている場合
次の環境変数を設定すると、スクリプトを編集せずにデフォルトをオーバーライドできます。
CLAUDE_IMAGE=my-claude:dev クロード-cli
# クロードコード自体にフラグを渡します
CLAUDE_CLI_FLAGS= " --resume $session_id --dangerously-skip-permissions " クロード-cli
# カスタム Dockerfile からイメージをビルドする
CLAUDE_CLI_DOCKERFILE= ~ /my-claude.dockerfile クロード-cli -b
CLAUDE_CLI_FLAGS は単語分割されるため、スペースで区切られた各トークンは
別の議論。スペースを含む引数を引用符で囲むことはサポートされていません - パス
値にスペースが含まれていないフラグのみ。
最初の実行時 (または -b の後)、スクリプトはグローバルにインストールされた @anthropic-ai/claude-code を使用して、node:lts に基づいて Docker イメージを構築します。バインド マウントされたボリュームでのファイル権限の問題を回避するために、コンテナ ユーザーはホスト ユーザーと同じ UID/GID で作成されます。
Dockerfile は優先順位に従って選択されます。
カスタム Dockerfile パス ( CLAUDE_CLI_DOCKERFILE )
スクリプトの隣の Dockerfile

( クロード-cli-dockerfile )
スクリプトに組み込まれたインライン Dockerfile (デフォルト)
カスタマイズは純粋に追加的です。上記のセットが何もない場合、スクリプトは組み込みの Dockerfile を /tmp に書き込み、そこからビルドするため、claude-cli.sh は自己完結型のままで、完全に単独でフェッチして実行できます。
イメージをカスタマイズするには、スクリプトの隣に claude-cli-dockerfile をドロップします。存在すると、それが優先されます。
# 好みに応じて claude-cli-dockerfile を作成し、再構築します
クロード-cli -b
あるいは、ディスク上の任意の場所にある Dockerfile を CLAUDE_CLI_DOCKERFILE で指定します。
イメージには claude-cli:latest のタグが付けられ、 -b を使用して明示的に再構築するまで、後続の実行で再利用されます。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to shirozuki/claude-cli development by creating an account on GitHub.

GitHub - shirozuki/claude-cli · GitHub
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
shirozuki
/
claude-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md claude-cli.sh claude-cli.sh View all files Repository files navigation
A lightweight shell wrapper that runs Claude Code inside a Docker container, keeping your host system clean while persisting your Claude configuration across sessions.
dmenu — mode selection prompt (not required when using -f )
fzf — directory picker; also used for mode selection with -f
Current user must be a member of the docker group
git clone https://github.com/shirozuki/claude-cli
chmod +x claude-cli/claude-cli.sh
ln -s " $PWD /claude-cli/claude-cli.sh " ~ /.local/bin/claude-cli
Usage
claude-cli [-b] [-c] [-f] [-h] [DIR...]
Option
Description
-b
Build (or rebuild) the claude-cli:latest Docker image
-c
Remove all claude-cli containers; optionally remove the image
-f
Use fzf instead of dmenu for the mode selection prompt
-h
Show help
Running without any option or argument launches a dmenu prompt to choose a mode. Pass -f to use fzf instead — useful on setups without dmenu or when you prefer a terminal picker.
Passing directories as arguments
You can skip the menu and pickers entirely by passing directories directly on the command line:
claude-cli ~ /projects/app # single-dir mode
claude-cli ~ /projects/app ~ /shared # multi-dir mode (first arg is the working dir)
One argument runs single-dir mode with that directory.
Two or more arguments run multi-dir mode, mounting every directory and using the first as the working directory.
Relative paths are resolved against the current working directory, so claude-cli . mounts $PWD . When directories are passed this way, dmenu and fzf are not required.
Launches Claude without mounting any project directories. Useful for general questions, quick tasks, or exploring Claude's capabilities without exposing local files.
Mounts the current working directory $PWD directly.
Opens an fzf picker to select one directory from your home tree (up to 3 levels deep). The selected directory is mounted and set as the working directory inside the container.
Opens fzf in multi-select mode to pick several directories (use Tab to select), then asks you to designate one of them as the working directory. All selected directories are mounted simultaneously — useful when working across multiple repos or sharing common config directories.
Claude's configuration ( ~/.claude/ and ~/.claude.json ) is bind-mounted into the container on every run, so your account, settings, and session history survive container restarts.
The mount source is resolved in this order:
$XDG_CONFIG_HOME/claude/ — if $XDG_CONFIG_HOME is set
The following environment variables can be set to override defaults without editing the script:
CLAUDE_IMAGE=my-claude:dev claude-cli
# Pass flags through to Claude Code itself
CLAUDE_CLI_FLAGS= " --resume $session_id --dangerously-skip-permissions " claude-cli
# Build the image from a custom Dockerfile
CLAUDE_CLI_DOCKERFILE= ~ /my-claude.dockerfile claude-cli -b
CLAUDE_CLI_FLAGS is word-split, so each space-separated token becomes a
separate argument. Quote arguments that contain spaces is not supported — pass
only flags whose values have no spaces.
On first run (or after -b ), the script builds a Docker image based on node:lts with @anthropic-ai/claude-code installed globally. The container user is created with the same UID/GID as the host user to avoid file permission issues on bind-mounted volumes.
The Dockerfile is selected in order of precedence:
Custom Dockerfile path ( CLAUDE_CLI_DOCKERFILE )
Dockerfile next to the script ( claude-cli-dockerfile )
Inline Dockerfile baked into the script (the default)
The customization is purely additive: with none of the above set, the script writes its built-in Dockerfile to /tmp and builds from that, so claude-cli.sh stays self-contained and can be fetched and run entirely on its own.
To customize the image, drop a claude-cli-dockerfile next to the script and it will be preferred when present:
# create claude-cli-dockerfile to taste, then rebuild
claude-cli -b
Alternatively, point CLAUDE_CLI_DOCKERFILE at a Dockerfile anywhere on disk.
The image is tagged claude-cli:latest and reused on subsequent runs until you explicitly rebuild with -b .
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
