---
source: "https://github.com/harihkk/Debug-Brief"
hn_url: "https://news.ycombinator.com/item?id=48622061"
title: "Show HN: DebugBrief – turn debugging sessions into reports, no AI"
article_title: "GitHub - harihkk/Debug-Brief: Local-first CLI that turns a debugging session into an honest markdown brief for PRs, handoffs and incidents · GitHub"
author: "itshkrishna"
captured_at: "2026-06-21T20:30:28Z"
capture_tool: "hn-digest"
hn_id: 48622061
score: 3
comments: 0
posted_at: "2026-06-21T19:57:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: DebugBrief – turn debugging sessions into reports, no AI

- HN: [48622061](https://news.ycombinator.com/item?id=48622061)
- Source: [github.com](https://github.com/harihkk/Debug-Brief)
- Score: 3
- Comments: 0
- Posted: 2026-06-21T19:57:54Z

## Translation

タイトル: HN を表示: DebugBrief – AI を使用せずにデバッグ セッションをレポートに変換します
記事のタイトル: GitHub - harihkk/Debug-Brief: デバッグ セッションを PR、ハンドオフ、インシデントの正直なマークダウン ブリーフに変えるローカルファースト CLI · GitHub
説明: デバッグ セッションを PR、ハンドオフ、インシデントの正直なマークダウン ブリーフに変えるローカル ファースト CLI - harihkk/Debug-Brief

記事本文:
GitHub - harihkk/Debug-Brief: デバッグ セッションを PR、ハンドオフ、インシデントの正直なマークダウン ブリーフに変えるローカル ファースト CLI · GitHub
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
ハリクク
/
デバッグの概要
公共
通知
通知設定を変更するにはサインインする必要があります

イングス
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
163 コミット 163 コミット .github .github docs docs 例 例 src/ debugbrief src/ debugbrief テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
デバッグ中に行ったことを記録し、プル リクエスト、ハンドオフ、またはインシデント メモ用の証拠に裏付けられた Markdown レポートに変換します。
DebugBrief は、ユーザーが書いたメモと実行したコマンドをキャプチャし、実際に起こったこと (試したこと、失敗したこと、成功したこと、その間に変更されたファイル) からレポートを作成します。 AI は使用されず、根本原因を推測したり、得られなかったテスト結果を報告したりすることはありません。
pipx インストール デバッグブリーフ
# または
UV ツールのインストール デバッグブリーフ
プレーンな pip install デバッグブリーフも機能します。 DebugBrief には Python 3.9 以降が必要です。ネイティブ Git は、リポジトリのメタデータと変更されたファイルをキャプチャするために使用できる場合に使用されます。デバッグするプロジェクトは Python である必要はありません。 DebugBrief のみが Python 上で実行されます。
debugbrief start " 間違った結果を返す add() を修正 "
デバッグ簡単なメモ「 add() は加算ではなく減算を行います。テストでは 5 が期待されます。」
debugbrief run -- python -m pytest -q test_calc.py # 失敗します
# 修正を加えます
debugbrief redo # 同じテスト、パスしました
debugbrief end # PR レポートを書き込みます
-- 以降はすべて入力したとおりに実行され、出力はターミナルにライブでストリーミングされます。 DebugBrief フラグ ( --timeout 、 --shell 、 --no-redact 、 --verify ) は -- の前に置きます。 redo は最後にキャプチャされたコマンドを再実行し、end はデフォルトで pr レポート モードになります。
ヒント: 1 行のエイリアスを使用すると、日常使用ではキャプチャ プレフィックスが表示されなくなります。
エイリアス db= " debugbr

IEF を実行 -- "
db pytest -q
得られるもの
記録された証拠のみから作成されたレポート。実際の実行からの短い抜粋:
## 概要
失敗したチェック ` python -m pytest -q test_calc.py ` は 2 秒間で 2 回の試行後に合格し、calc.py に変更が加えられました。
## 赤から緑へ
チェックは 12:02:09 に失敗し、` python -m pytest -q test_calc.py ` は 12:02:10 (ウィンドウ 1s) に成功しました。
チェックが失敗してから合格するまでの間に、次のファイルが変更されました (相関関係があり、原因は証明されていません)。
- ` calc.py `
完全なサンプル: PR 、ハンドオフ、インシデント。
run は、出力ストリームがライブになるように疑似端末でコマンドを実行し、実際の終了コード、制限された出力プレビュー、期間、およびコマンドごとの git スナップショットを記録します。
合否は終了コードによってのみ決まります。コマンドは、認識されたテスト、ビルド、lint、またはタイプチェック ランナーが実際に終了した場合にのみ検証済みとしてカウントされます。
どの言語でも機能します。認識されたランナー (pytest、jest、vitest、go test、cargo test、dotnet test、make check など) は自動的に分類されます。他のコマンドは引き続きキャプチャされるため、 --verify を使用してチェックマークを付けます。
end は、赤から緑へのウィンドウ、再現コマンドと検証コマンド、タイムライン、観察されたエラー、および失敗した試行などのイベントからレポートを取得します。空のセクションは省略されます。
シークレットは、ディスクに何かが書き込まれる前に編集されます。編集はベストエフォートです。 --no-redact は、単一のコマンドをオプトアウトします。
完全なコマンド リファレンスと完全な認識ランナー リスト: docs/COMMANDS.md 。セキュリティ モデルと編集の詳細: SECURITY.md 。
レポートをプル リクエストに直接投稿します (GitHub CLI はオプション)。
デバッグブリーフ終了 --stdout | gh pr コメント --body-file -
コマンド
それぞれの完全なフラグと動作: docs/COMMANDS.md 。
DebugBrief は、Python 標準ライブラリとネイティブ git を使用します。 Python 3.11 以降では他に何も必要ありません

。 Python 3.9 および 3.10 では、小さな tomli パッケージを使用して、オプションの .debugbrief.toml を読み取ります。 DebugBrief 自体はネットワーク リクエストを行わず、AI も使用せず、テレメトリも収集しません。
Linux と macOS は、Python 3.9 から 3.14 までの CI でテストされています。他の Unix 系システムは動作する可能性がありますが、現在はテストされていません。ネイティブ Windows および PowerShell はサポートされていません。
キャプチャは debugbrief run を通じて明示的に行われます。制限されたプレビューがレポートに保存されている間、出力ストリームはライブになるため、完全なトランスクリプトは存在しません。
全画面 TUI (vim セッション、 htop ) は意味のある形でキャプチャされません。これらを直接実行し、結果を note に記録します。
編集は保守的かつ最善の努力で行われます。すべての秘密を把握できるわけではありません。
Git セクションにはネイティブ git が必要です。リポジトリの外では省略されます。
pip install -e " .[dev] "
pytest
セットアップと貢献のガイドライン全体については、CONTRIBUTING.md を参照してください。
デバッグ セッションを PR、ハンドオフ、インシデントの正直なマークダウン ブリーフに変えるローカル ファースト CLI
pypi.org/project/debugbrief/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
デバッグブリーフ 1.3.0
最新の
2026 年 6 月 17 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-first CLI that turns a debugging session into an honest markdown brief for PRs, handoffs and incidents - harihkk/Debug-Brief

GitHub - harihkk/Debug-Brief: Local-first CLI that turns a debugging session into an honest markdown brief for PRs, handoffs and incidents · GitHub
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
harihkk
/
Debug-Brief
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
163 Commits 163 Commits .github .github docs docs examples examples src/ debugbrief src/ debugbrief tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
Record what you do while debugging and turn it into an evidence-backed Markdown report for a pull request, a handoff, or an incident note.
DebugBrief captures the notes you write and the commands you run, then builds a report from what actually happened: what you tried, what failed, what passed, and which files changed in between. It does not use AI and does not infer a root cause or report a test result you did not get.
pipx install debugbrief
# or
uv tool install debugbrief
Plain pip install debugbrief works too. DebugBrief needs Python 3.9 or newer. Native Git is used when available to capture repository metadata and changed files. The project you debug does not need to be Python; only DebugBrief runs on Python.
debugbrief start " Fix add() returning wrong result "
debugbrief note " add() subtracts instead of adds; the test expects 5. "
debugbrief run -- python -m pytest -q test_calc.py # fails
# make your fix
debugbrief redo # same test, now passes
debugbrief end # writes the PR report
Everything after -- runs exactly as you typed it, with its output streaming live to your terminal. DebugBrief flags ( --timeout , --shell , --no-redact , --verify ) go before the -- . redo re-runs the last captured command, and end defaults to the pr report mode.
Tip: a one-line alias makes the capture prefix disappear in daily use.
alias db= " debugbrief run -- "
db pytest -q
What you get
A report built only from recorded evidence. A short excerpt from a real run:
## Summary
Failing check ` python -m pytest -q test_calc.py ` passed after 2 attempts over 2s, changes touched calc.py.
## Red to green
A check failed at 12:02:09 and ` python -m pytest -q test_calc.py ` passed at 12:02:10 (window 1s).
Between the failing and passing checks, these files changed (correlation, not proven cause):
- ` calc.py `
Full samples: PR , handoff , incident .
run executes a command under a pseudo-terminal so its output streams live, then records the real exit code, a bounded output preview, the duration, and a per-command git snapshot.
Pass or fail comes only from the exit code. A command counts as verified only when a recognized test, build, lint, or typecheck runner actually exits 0.
It works with any language. A recognized runner (pytest, jest, vitest, go test, cargo test, dotnet test, make check, and more) is classified automatically. Any other command is still captured, and you mark it a check with --verify .
end derives the report from those events: the red-to-green window, the reproduce and verify commands, a timeline, the observed error, and the failed attempts. Empty sections are omitted.
Secrets are redacted before anything is written to disk. Redaction is best effort; --no-redact opts out for a single command.
Full command reference and the complete recognized-runner list: docs/COMMANDS.md . Security model and redaction details: SECURITY.md .
Post a report straight to a pull request (GitHub CLI optional):
debugbrief end --stdout | gh pr comment --body-file -
Commands
Full flags and behavior for each: docs/COMMANDS.md .
DebugBrief uses the Python standard library and native git . On Python 3.11 and newer it needs nothing else. On Python 3.9 and 3.10 it uses the small tomli package to read an optional .debugbrief.toml . DebugBrief itself makes no network requests, uses no AI, and collects no telemetry.
Linux and macOS are tested in CI across Python 3.9 through 3.14. Other Unix-like systems may work but are not currently tested. Native Windows and PowerShell are not supported.
Capture is explicit through debugbrief run . Output streams live while a bounded preview is stored for the report, so there is no full transcript.
Full-screen TUIs (a vim session, htop ) are not meaningfully captured. Run those directly and record the outcome with note .
Redaction is conservative and best effort; it does not catch every secret.
Git sections need native git ; outside a repository they are omitted.
pip install -e " .[dev] "
pytest
See CONTRIBUTING.md for the full setup and contribution guidelines.
Local-first CLI that turns a debugging session into an honest markdown brief for PRs, handoffs and incidents
pypi.org/project/debugbrief/
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
DebugBrief 1.3.0
Latest
Jun 17, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
