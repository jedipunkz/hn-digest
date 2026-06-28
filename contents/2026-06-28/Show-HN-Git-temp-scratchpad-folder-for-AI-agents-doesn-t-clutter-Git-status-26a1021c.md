---
source: "https://github.com/sebmellen/git-temp"
hn_url: "https://news.ycombinator.com/item?id=48703875"
title: "Show HN: Git-temp – scratchpad folder for AI agents; doesn't clutter Git status"
article_title: "GitHub - sebmellen/git-temp · GitHub"
author: "sebmellen"
captured_at: "2026-06-28T03:03:35Z"
capture_tool: "hn-digest"
hn_id: 48703875
score: 1
comments: 0
posted_at: "2026-06-28T02:48:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Git-temp – scratchpad folder for AI agents; doesn't clutter Git status

- HN: [48703875](https://news.ycombinator.com/item?id=48703875)
- Source: [github.com](https://github.com/sebmellen/git-temp)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T02:48:04Z

## Translation

タイトル: HN を表示: Git-temp – AI エージェント用のスクラッチパッド フォルダー。 Git ステータスを乱雑にしない
記事のタイトル: GitHub - sebmellen/git-temp · GitHub
説明: GitHub でアカウントを作成して、sebmellen/git-temp の開発に貢献します。

記事本文:
GitHub - sebmellen/git-temp · GitHub
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
セブメレン
/
git-temp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル

レ
4 コミット 4 コミット bin bin test test FUTURE_GAP_ANALYSIS.md FUTURE_GAP_ANALYSIS.md LICENSE LICENSE README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
目的: コミットせずに、Git リポジトリ内にローカル AI スクラッチパッドを作成します。
git-temp は temp/ フォルダーを作成し、 .git/info/exclude に追加して、エディターや AI ツールに表示できるようにします。つまり、ノイズの多い git status 、偶発的なコミット、およびインデックス付けや @ タグ付けからファイルを隠す .gitignore ルールがないことを意味します。
npx git-temp # temp/を作成します
npx git-temp Notes # メモを作成/
npx git-temp status # 非表示のスクラッチパッドの内容を表示します
npx git-temp clean -f # temp/ を空にして再作成します
npx git-tempintegrate # 既存のエージェント指示ファイルを更新します
それが生み出すもの
temp/scripts/ 1 回限りのスクリプトの場合
JSON、CSV、ログ、ペイロードの temp/dumps/
temp/drafts/ メモとドラフト仕様用
人間と AI エージェントにフォルダーを説明する temp/README.md
AI/エディタツールによるスクラッチ作品の検索とタグ付けを可能にします。
共有された .gitignore ではなく、 .git/info/exclude を通じてローカルでのスクラッチ作業を無視します。
誤って一時ファイルをコミットしないようにしてください。
* または /* を含むブロックされているネストされた temp/.gitignore ファイルを削除します。
npx の下では依存性ゼロで高速に保ちます。
git-temp [ディレクトリ] [--integrate]
スクラッチパッド、サブフォルダ、README、およびローカル除外エントリを作成します。
git-temp ステータス [ディレクトリ]
Git で非表示になったファイルのファイル数、ディレクトリ数、サイズ、および変更時間を表示します。
git-temp clean [ディレクトリ] [--force]
スクラッチパッドの内容を削除し、標準構造を再作成します。
git-temp 統合 [ディレクトリ]
AI スクラッチパッド ガイドラインを既存の命令ファイルのみに追加します。
.github/copilot-instructions.md
www.npmjs.com/package/git-temp
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
のために

クス
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to sebmellen/git-temp development by creating an account on GitHub.

GitHub - sebmellen/git-temp · GitHub
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
sebmellen
/
git-temp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits bin bin test test FUTURE_GAP_ANALYSIS.md FUTURE_GAP_ANALYSIS.md LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Purpose: create a local AI scratchpad inside a Git repo without committing it.
git-temp makes a temp/ folder, adds it to .git/info/exclude , and keeps it visible to editors and AI tools. That means no noisy git status , no accidental commits, and no .gitignore rule that hides files from indexing or @ tagging.
npx git-temp # creates temp/
npx git-temp notes # creates notes/
npx git-temp status # shows hidden scratchpad contents
npx git-temp clean -f # empties and recreates temp/
npx git-temp integrate # updates existing agent instruction files
What it creates
temp/scripts/ for one-off scripts
temp/dumps/ for JSON, CSV, logs, and payloads
temp/drafts/ for notes and draft specs
temp/README.md explaining the folder to humans and AI agents
Keep scratch work searchable and taggable by AI/editor tooling.
Ignore scratch work locally through .git/info/exclude , not shared .gitignore .
Avoid committing temporary files by accident.
Remove blocking nested temp/.gitignore files that contain * or /* .
Stay zero-dependency and fast under npx .
git-temp [directory] [--integrate]
Creates the scratchpad, subfolders, README, and local exclude entry.
git-temp status [directory]
Shows file count, directory count, size, and modified times for files Git now hides.
git-temp clean [directory] [--force]
Deletes scratchpad contents, then recreates the standard structure.
git-temp integrate [directory]
Appends an AI scratchpad guideline to existing instruction files only:
.github/copilot-instructions.md
www.npmjs.com/package/git-temp
Resources
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
