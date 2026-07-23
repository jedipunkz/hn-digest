---
source: "https://github.com/tomlin7/biscuit"
hn_url: "https://news.ycombinator.com/item?id=49027355"
title: "Show HN: Biscuit – Native, extensible, AI code editor. <20 mb in size"
article_title: "GitHub - tomlin7/biscuit: biscuit is a native, extensible, AI code editor. lightweight <20 mb in size. install and start using in seconds. · GitHub"
author: "tomlin7"
captured_at: "2026-07-23T20:12:12Z"
capture_tool: "hn-digest"
hn_id: 49027355
score: 1
comments: 0
posted_at: "2026-07-23T20:10:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Biscuit – Native, extensible, AI code editor. <20 mb in size

- HN: [49027355](https://news.ycombinator.com/item?id=49027355)
- Source: [github.com](https://github.com/tomlin7/biscuit)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T20:10:23Z

## Translation

タイトル: Show HN: Biscuit – ネイティブで拡張可能な AI コード エディター。サイズが 20 MB 未満
記事タイトル: GitHub - tomlin7/biscuit: biscuit は、ネイティブで拡張可能な AI コード エディターです。サイズが 20 MB 未満の軽量。数秒でインストールして使用を開始できます。 · GitHub
説明: biscuit は、ネイティブで拡張可能な AI コード エディターです。サイズが 20 MB 未満の軽量。数秒でインストールして使用を開始できます。 - tomlin7/ビスケット

記事本文:
GitHub - tomlin7/biscuit: biscuit は、ネイティブで拡張可能な AI コード エディターです。サイズが 20 MB 未満の軽量。数秒でインストールして使用を開始できます。 · GitHub
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
トムリン7
/
ビスケット
公共
ああ、ああ！
エラーが発生しました

ルをロードしています。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,859 コミット 1,859 コミット .biscuit .biscuit .github .github .ted .ted .vscode .vscode config config docs docs リソース リソース スクリプト スクリプト src/ biscuit src/ biscuit テスト テスト www www .deepsource.toml .deepsource.toml .editorconfig .editorconfig .gitignore .gitignore .prettierignore .prettierignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md logo.svg logo.svg mkdocs.yml mkdocs.yml poureted.lock votes.lock pyproject.toml pyproject.toml要件.txt要件.txt uv.lock uv.lock vercel.json vercel.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
biscuit は、エージェントを備えた高速で拡張可能なネイティブ コード エディターです。サイズが 20 MB 未満の軽量。数秒でインストールして使用を開始できます。
人気の拡張機能を探索し、私たちが取り組んでいる新しい拡張機能マーケットプレイスにアクセスしてください
開発者/ユーザー ガイドと API リファレンスについては、ドキュメントと WIP を確認してください。そのため、deepwiki を確認することをお勧めします。
超能力が満載、機能のリスト
以下を実行して最新リリースをインストールします。
> pip インストールビスケットエディター
biscuit path/to/src を使用してプロジェクトをすぐに開き、編集を開始します。必要に応じて、他のインストール方法 (pyinstaller など) を参照してください。
プロジェクト構造の概要と環境のセットアップについては、ドキュメントと寄稿ガイドを確認してください。
新しい拡張機能を作成するには、拡張機能のドキュメントを読んでください:>
gemini、人間 API サポート ( claude-4-5-opus/sonnet/haiku 、 gemini-2-5-flash/pro )
タスクリストを備えた計画エージェント
ファイル読み取りツール
biscu を通じて LLM プロバイダーを追加する

それの拡張子
チャットにコンテキストを追加するためのファイルを添付する
LLM プロバイダー拡張の例 (古いものは現在非推奨です)
ollam 拡張機能を使用してローカル LLM を実行する (非推奨)
ビスケット端末内での LLM 呼び出し (# 端末内でプロンプトを使用し、応答を受け入れる/拒否します)
ツリーシッターベースの高速解析とハイライト
エディター内のコード補完 (アイコン付き)
カーソルを合わせるとシンボル定義/ドキュメント文字列が表示されます (ハイライト + マークダウンでレンダリング)
開いているエディターでシンボルをナビゲートするためのシンボル アウトライン サイドバー パネル
コマンド パレット Ctrl + J によるシンボル検索
シンボルの定義/宣言にジャンプするフローティング ピーク ウィジェット
開いているエディターでのシンボル参照
ビスケット拡張機能を使用して言語サーバーを追加する
より多くの言語サーバーが拡張機能を介して登録されています。参考として、 Rust および Clangd 拡張機能を参照してください。
変更/段階的変更のための分割差分ビューア
簡単にアクセスできる重要な Git 操作 (プッシュ、プル、コミット、ステージング、ステージング解除、ブランチの切り替え)
リポジトリのクローンを作成し、アクティブなウィンドウまたは新しいウィンドウですぐに開きます
エディター内で gitHub の問題/prs を表示します (TODO: 無効になっている rn、拡張機能に変換されます)
ripgrep ベースの高速検索、ステータスバーからすぐにアクセス可能
出現箇所を個別に、または一度にすべて置き換える
正規表現のサポート、大文字と小文字を区別した検索、その他のカスタマイズ
フローティング検索/置換ウィジェットを使用して、開いているエディタ内を検索します
ファイル全体にブレークポイントを設定する
すべてのランタイム変数の検査パネル
デバッグ中にランタイム変数を変更する
呼び出しスタックの視覚化と例外の追跡
デバッガの追加は、ビスケット拡張機能を通じて登録できます。
GUI を介して利用可能なすべての拡張機能をインストールおよび管理する
Biscuit 内での拡張子検索
拡張機能ブートストラップ cli コマンドとテンプレート
拡張機能マーケットプレイスの Web サイト: ここにアクセスしてください
分割マークダウン エディター、プレーン HTML レンダラー
トグル

相対行番号付けのサポート
ビスケット拡張機能を使用してフォーマッタを追加する
フォーマッタ拡張子: black [非推奨]、ruff [非推奨]、YAPF [非推奨]、autopep8 [非推奨] (参考までに)。
ドラッグアンドドロップしてビスケット内のファイルまたはフォルダーを開きます
プロジェクトの editorconfig サポート
洗練されたコマンド パレット ( src/biscuit/commands にある静的コマンドの完全なリスト)
biscuit は MIT ライセンスを使用します。LICENSE ファイルを参照してください。
biscuit は、ネイティブで拡張可能な AI コード エディターです。サイズが 20 MB 未満の軽量。数秒でインストールして使用を開始できます。
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
33 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

biscuit is a native, extensible, AI code editor. lightweight <20 mb in size. install and start using in seconds. - tomlin7/biscuit

GitHub - tomlin7/biscuit: biscuit is a native, extensible, AI code editor. lightweight <20 mb in size. install and start using in seconds. · GitHub
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
tomlin7
/
biscuit
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,859 Commits 1,859 Commits .biscuit .biscuit .github .github .ted .ted .vscode .vscode config config docs docs resources resources scripts scripts src/ biscuit src/ biscuit tests tests www www .deepsource.toml .deepsource.toml .editorconfig .editorconfig .gitignore .gitignore .prettierignore .prettierignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md logo.svg logo.svg mkdocs.yml mkdocs.yml poetry.lock poetry.lock pyproject.toml pyproject.toml requirements.txt requirements.txt uv.lock uv.lock vercel.json vercel.json View all files Repository files navigation
biscuit is a fast, extensible, native code editor with agents. lightweight <20 mb in size. install and start using in seconds.
explore popular extensions, visit the new extension marketplace we've been working on
for developer/user guides & API reference, check documentation , WIP so i recommend checking deepwiki.
packed with superpowers, list of features
install the latest release by running:
> pip install biscuit-editor
quickly open up a project using biscuit path/to/src and start editing. see other installation methods if you'd like to (like pyinstaller).
please check the docs and contributing guide for a quick tour of the project structure and to set up the environment.
to make a new extension, read the extension docs :>
gemini, anthropic API support ( claude-4-5-opus/sonnet/haiku , gemini-2-5-flash/pro )
planning agent with task list
ReadFileTool
add more LLM providers through biscuit extensions
attach files for adding context in chat
LLM provider extension examples (old ones are now deprecated)
run local LLMs with ollama extension (deprecated)
LLM calls inside biscuit terminals (use # your prompt inside terminal, then accept/decline response)
fast tree-sitter based parsing and highlights
code completions within editor (with icons)
hover for symbol definition/docstring (rendered with highlights + markdown)
symbol outline sidebar panel for navigating symbols in open editor
symbol search through command palette Ctrl + J )
floating peek widget to jump-to-definition/declaration of symbols
symbol references in open editor
adding more language servers through biscuit extensions
more language servers are registered through extensions, see the rust , clangd extensions for reference.
split diff viewer for changes/staged changes
essential git operations easily accessible (push, pull, commit, stage, unstage, switch branches)
clone repositories and immediately open in active window, or new window
view gitHub issues/prs within editor (TODO: disabled rn, will be converted to an extension)
ripgrep based fast search, quickly accessible from statusbar
replace occurrences individually or all at once
regex support, case sensitive search and more customization
search within open editors with floating find-replace widget
setting breakpoints across files
inspection panel for all runtime variables
modify runtime variables while debugging
call stack visualization and exception tracing
add debuggers can be registered through biscuit extensions.
install and manage all available extensions though a gui
extension search within biscuit
extension bootstrapping cli commands and templates
extensions marketplace website: visit here
split markdown editor, plain HTML renderer
toggle relative line numbering support
add formatters through biscuit extensions
formatter extensions: black [DEPRECATED], ruff [DEPRECATED], YAPF [DEPRECATED], autopep8 [DEPRECATED] for reference.
drag-n-drop to open files or folders in biscuit
editorconfig support for projects
sophisticated command palette (full list of static commands in src/biscuit/commands )
biscuit uses the MIT License, see LICENSE file.
biscuit is a native, extensible, AI code editor. lightweight <20 mb in size. install and start using in seconds.
Readme MIT license Code of conduct
Security policy Activity Stars
33 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
