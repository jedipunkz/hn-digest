---
source: "https://github.com/robinduckett/leclaude"
hn_url: "https://news.ycombinator.com/item?id=49161957"
title: "Show HN: Leclaude – A little badge for your Claude Code projects"
article_title: "GitHub - robinduckett/leclaude: Leclaude shows a robot badge in Explorer on each folder that has Claude Code session history. · GitHub"
author: "robinduckett"
captured_at: "2026-08-03T21:58:38Z"
capture_tool: "hn-digest"
hn_id: 49161957
score: 1
comments: 0
posted_at: "2026-08-03T21:52:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Leclaude – A little badge for your Claude Code projects

- HN: [49161957](https://news.ycombinator.com/item?id=49161957)
- Source: [github.com](https://github.com/robinduckett/leclaude)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T21:52:32Z

## Translation

タイトル: HN を表示: Leclaude – クロード コード プロジェクトのための小さなバッジ
記事のタイトル: GitHub - robinduckett/leclaude: Leclaude は、クロード コードのセッション履歴がある各フォルダーのエクスプローラーにロボット バッジを表示します。 · GitHub
説明: Leclaude は、クロード コードのセッション履歴がある各フォルダーのエクスプローラーにロボット バッジを表示します。 - ロビンダケット/ルクロード

記事本文:
GitHub - robinduckett/leclaude: Leclaude は、クロード コードのセッション履歴がある各フォルダーのエクスプローラーにロボット バッジを表示します。 · GitHub
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
ロビンダケット
/
ルクロード
公共
通知
c にサインインする必要があります

ハンゲの通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .github/ workflows .github/ workflows 資産 アセット ドキュメント ドキュメント スクリプト スクリプト src src テスト テスト .gitignore .gitignore CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CMakePresets.json CMakePresets.json ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Leclaude は、Windows エクスプローラーのクロード コード セッション履歴が含まれる各フォルダーにロボット バッジを表示します。
ディスクを見ると、Claude Code を使用したフォルダーがすぐにわかります。
Leclaude は、x64 および ARM64 上の Windows 10 および Windows 11 用のシェル拡張機能です。
Claude Code は、開いている各フォルダーを %USERPROFILE%\.claude\projects\ に記録します。
Leclaude は、エクスプローラーにアイコン オーバーレイ ハンドラーを登録します。
ハンドラーは、表示されているフォルダーごとに、メモリ内のセット内で 1 回の高速検索を実行します。
フォルダーにセッション履歴がある場合、エクスプローラーはフォルダー アイコンにロボット バッジを描画します。
ウォッチャー スレッドはクロード コード データを追跡し、バッジを更新します。
設計ドキュメントには詳細が記載されています。
リリース ページからプロセッサ用の zip ファイルをダウンロードします: x64 または arm64。
ファイルが本物であることを確認してください (オプション、以下を参照)。
解凍したフォルダーで管理者 PowerShell を開きます。
スクリプトは DLL を %ProgramFiles%\Leclaude にコピーし、ハンドラーを登録して、エクスプローラーを再起動します。
Leclaude を削除するには、「.\uninstall.ps1」と入力します。
リリースが本物であることを確認する
リリース ワークフローは、GitHub アーティファクト証明書を使用して各 zip ファイルを証明します。
Sigstore は証明書に署名します。チェックを行うには、次のように入力します。
gh 認証検証 Leclaude-<tag>-x64.zip --owner robinduckett
このコマンドは、ワークフローとファイルを作成したコミットを表示します。
ブイ

ld は Ninja ジェネレーターで CMake を使用します。 Visual Studio には両方のツールが含まれています。
Visual Studio の開発者 PowerShell を開きます。
「cmake --preset x64-release」と入力します。
「cmake --build --preset x64-release」と入力します。
テストを実行するには、ctest --test-dir build/x64-release --output-on-failure と入力します。
インストールには管理者権限が必要です。 Windows は、HKLM でのみオーバーレイ登録を受け入れます。
Windows では、システム全体で最大 15 のオーバーレイ タイプが表示されます。クラウド ストレージ ツールではその多くが使用されます。競合するツールが多すぎる場合、Windows は Leclaude バッジを無視することがあります。 ShellExView ツールは、システム上の競合を表示します。
クラウド同期フォルダー (OneDrive フォルダーなど) では、Windows 11 はサードパーティのバッジを非表示にすることができます。
更新後、バッジには古いアイコンが表示されますが、一部のサイズでのみ表示されます。
エクスプローラーは、アイコンをディスク上のキャッシュ ファイルに、サイズごとに 1 つのファイルとして保持します。
Explorer を再起動しても、これらのファイルは削除されません。それらを削除するには、PowerShell に次のコマンドを入力します。
プロセスの停止 - 名前エクスプローラー - 強制
Remove-Item " $ env: LOCALAPPDATA \Microsoft\Windows\Explorer\iconcache_*.db " - 強制 - ErrorAction SilentlyContinue
if ( -not ( Get-Process Explorer - ErrorAction SilentlyContinue)) { Start-Process Explorer.exe }
コンピュータを再起動すると、この問題も修正されます。
バッジがまったく表示されない場合:
Windows では、最大 15 のオーバーレイ タイプが表示されます。 ShellExView ツールを使用して、ロードされたハンドラーを確認します。
レジストリ キー ShellIconOverlayIdentifiers の下の最初の 15 個の名前に「Leclaude」という名前が含まれていることを確認します。
Leclaude は、Claude Code のセッション履歴が含まれる各フォルダーのエクスプローラーにロボット バッジを表示します。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Leclaude shows a robot badge in Explorer on each folder that has Claude Code session history. - robinduckett/leclaude

GitHub - robinduckett/leclaude: Leclaude shows a robot badge in Explorer on each folder that has Claude Code session history. · GitHub
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
robinduckett
/
leclaude
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .github/ workflows .github/ workflows assets assets docs docs scripts scripts src src tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CMakePresets.json CMakePresets.json LICENSE LICENSE README.md README.md View all files Repository files navigation
Leclaude shows a robot badge in Windows Explorer on each folder that has Claude Code session history.
When you look through your disk, you can immediately see the folders where you did work with Claude Code.
Leclaude is a shell extension for Windows 10 and Windows 11, on x64 and ARM64.
Claude Code records each opened folder in %USERPROFILE%\.claude\projects\ .
Leclaude registers an icon-overlay handler with Explorer.
For each visible folder, the handler does one fast search in a set in memory.
When the folder has session history, Explorer draws the robot badge on the folder icon.
A watcher thread follows the Claude Code data and refreshes the badges.
The design documents give the full details:
Download the zip file for your processor from the releases page : x64 or arm64 .
Make sure that the file is authentic (optional, see below).
Open an administrator PowerShell in the unpacked folder.
The script copies the DLL to %ProgramFiles%\Leclaude , registers the handler, and restarts Explorer.
To remove Leclaude, enter: .\uninstall.ps1
Make sure that a release is authentic
The release workflow attests each zip file with GitHub artifact attestation .
Sigstore signs the attestation. To do a check, enter:
gh attestation verify Leclaude-<tag>-x64.zip --owner robinduckett
The command shows the workflow and the commit that made the file.
The build uses CMake with the Ninja generator. Visual Studio contains both tools.
Open a Developer PowerShell for Visual Studio.
Enter: cmake --preset x64-release
Enter: cmake --build --preset x64-release
To do the tests, enter: ctest --test-dir build/x64-release --output-on-failure
The installation needs administrator rights. Windows accepts overlay registrations in HKLM only.
Windows shows a maximum of 15 overlay types for the full system. Cloud-storage tools use many of them. When too many tools compete, Windows can ignore the Leclaude badge. The ShellExView tool shows the competition on your system.
In a cloud sync folder (for example, a OneDrive folder), Windows 11 can hide third-party badges.
The badge shows an old icon after an update, but only at some sizes:
Explorer keeps the icons in cache files on the disk, with one file for each size.
An Explorer restart does not delete these files. To delete them, enter these commands in PowerShell:
Stop-Process - Name explorer - Force
Remove-Item " $ env: LOCALAPPDATA \Microsoft\Windows\Explorer\iconcache_*.db " - Force - ErrorAction SilentlyContinue
if ( -not ( Get-Process explorer - ErrorAction SilentlyContinue)) { Start-Process explorer.exe }
A restart of the computer also corrects this.
The badge does not show at all:
Windows shows a maximum of 15 overlay types. Use the ShellExView tool to see the loaded handlers.
Make sure that the name " Leclaude" is in the first 15 names under the registry key ShellIconOverlayIdentifiers .
Leclaude shows a robot badge in Explorer on each folder that has Claude Code session history.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
