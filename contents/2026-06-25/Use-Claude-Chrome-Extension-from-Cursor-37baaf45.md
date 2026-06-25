---
source: "https://github.com/bebrws/claude-chrome-extension-from-cursor"
hn_url: "https://news.ycombinator.com/item?id=48667203"
title: "Use Claude Chrome Extension from Cursor"
article_title: "GitHub - bebrws/claude-chrome-extension-from-cursor · GitHub"
author: "bebrws"
captured_at: "2026-06-25T00:31:39Z"
capture_tool: "hn-digest"
hn_id: 48667203
score: 1
comments: 1
posted_at: "2026-06-25T00:20:09Z"
tags:
  - hacker-news
  - translated
---

# Use Claude Chrome Extension from Cursor

- HN: [48667203](https://news.ycombinator.com/item?id=48667203)
- Source: [github.com](https://github.com/bebrws/claude-chrome-extension-from-cursor)
- Score: 1
- Comments: 1
- Posted: 2026-06-25T00:20:09Z

## Translation

タイトル: カーソルから Claude Chrome 拡張機能を使用する
記事のタイトル: GitHub - bebrws/claude-chrome-extension-from-cursor · GitHub
説明: GitHub でアカウントを作成して、bebrws/claude-chrome-extension-from-cursor の開発に貢献します。

記事本文:
GitHub - bebrws/claude-chrome-extension-from-cursor · GitHub
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
ビブルズ
/
クロード-Chrome-カーソルからの拡張機能
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
bebrws/クロード-クローム-ext

カーソルからのテンション
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット カーソル カーソル README.md README.md claude-current-tab.zsh claude-current-tab.zsh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Chrome 拡張機能のカーソルの使用
Claude Chrome の拡張機能は、Cursor の現在のブラウザ統合よりもはるかに強力であることがわかりました。たとえば、SSO を使用するシステムにログインする必要があるものをデバッグする必要がある場合は特にそうです。このリポジトリを使用すると、すでに開いているデフォルトのブラウザで現在選択されているタブに対して Cursor から Claude の Chrome 拡張機能を使用できるようになります。 Cursor からの Claude の Chrome 拡張機能を利用できるようにこれを作成しました。Cursor はローカル ヘルパー スクリプトを呼び出し、スクリプトは既存のデフォルトの Chromium ブラウザ セッションをアクティブにし、Claude のサイド パネルを開き、プロンプトを送信し、macOS アクセシビリティを通じて Claude の最終応答を読み取り、その応答を Cursor に返して、Cursor がファイルのコーディングや編集を続行できるようにします。
デフォルトのブラウザで関連ページを開き、それが現在選択されているタブであることを確認します。次に、Cursor に Claude Chrome 拡張機能を使用するように依頼します。
/use-claude-extension を使用して、do_something ジョブの現在の実行のログを分析します。そして、@do_something.py ジョブを改善するための提案をしてください。
Cursor エージェントはシステム ブラウザをアクティブにし、現在開いているタブで Claude Chrome 拡張機能を使用します。カーソルは、このリポジトリに含まれるスクリプトを使用して、Claude Chrome 拡張機能からの応答を待ちます。
Cursor はリポジトリ内のコード編集には優れていますが、コードを適切に変更するために必要な情報が認証されたブラウザ セッションにのみ存在する場合があります。
このプロジェクトでは、Cursor エージェントが Claude の Chrome 拡張機能に現在の設定を検査するように依頼できます。

ブラウザ タブを選択すると、クロードの回答が Cursor に戻され、エージェントがその情報を使用してコードを変更できるようになります。
これは意図的に claude --chrome 、Playwright、または新しく起動したブラウザを使用しません。既存のデフォルトのブラウザ ウィンドウ、現在のタブ、インストールされている Claude Chrome 拡張機能、およびログインしたブラウザ セッションが使用されます。
macOS のデフォルトのブラウザを解決します。
既存のブラウザウィンドウをアクティブにします。
現在選択されているタブを使用します。
Claude Chrome 拡張機能のサイド パネルを開きます。
カーソルからブラウザのクロードにプロンプ​​トを送信します。
クロードの最終返答を待ちます。
macOS アクセシビリティを通じて応答を読み取ります。
Cursor がコーディング タスクを続行できるように、Cursor に応答を返します。
。
§── README.md
§── クロードカレントタブ.zsh
└── カーソル/
§── install.sh
§── ビン/
│ └── クロード-current-tab.zsh
§── ルール/
│ └── use-claude-chrome-extension.mdc
└── コマンド/
└── use-claude-extension.md
含まれるファイル
claude-current-tab.zsh — オリジナルのヘルパー スクリプト。
Cursor/bin/claude-current-tab.zsh — インストーラーによって使用されるバンドルされたヘルパー スクリプト。
Cursor/rules/use-claude-chrome-extension.mdc — このワークフローをいつどのように使用するかを Cursor に教えるグローバル Cursor ルール。
Cursor/commands/use-claude-extension.md — オプションのカーソル スラッシュ コマンド。インストール後に /use-claude-extension として使用できます。
Cursor/install.sh — ヘルパー コマンド、グローバル カーソル ルール、およびスラッシュ コマンドのワンステップ インストーラー。
カーソル ルールとスラッシュ コマンドは両方とも、インストール後にこのヘルパー コマンドを呼び出します。
~ /.local/bin/claude-current-tab
要件
Brave Browser や Google Chrome などの、Chromium 互換のデフォルトのシステム ブラウザ。
Claude Chrome 拡張機能は、デフォルトのブラウザ/プロファイルにインストールされています。
を切り替えるための Claude 拡張機能のショートカット

サイドパネルが利用可能である必要があります。テストされたセットアップでは、拡張コマンドは toggle-side-panel であり、 Cmd+E にバインドされています。
クロードに分析させたいブラウザ タブは、デフォルトのブラウザで現在選択されているタブである必要があります。
ローカル ヘルパー スクリプトは実行可能である必要があります。
Safari と Firefox は Claude Chrome 拡張機能をホストできないため、このワークフローには使用できません。
このリポジトリのルートから:
./cursor/install.sh
インストーラーは安全に再実行できます。それは以下をインストールします:
~/.local/bin/claude-current-tab
~/.cursor/rules/use-claude-chrome-extension.mdc
~/.cursor/commands/use-claude-extension.md
次に、Cursor をリロードまたは再起動して、新しいルールとコマンドを選択します。
ヘルパーが現在のデフォルト ブラウザー タブを読み取ることができることを確認します。
~ /.local/bin/claude-current-tab --print-context
デフォルトのブラウザ、バンドル ID、アクティブなタブの URL、およびアクティブなタブのタイトルを含む JSON が表示されるはずです。
install.sh を使用したくない場合は、これらの部分を手動でインストールしてください。
このリポジトリのルートから:
chmod +x カーソル/bin/claude-current-tab.zsh
mkdir -p ~ /.local/bin
ln -sf " $( pwd ) /cursor/bin/claude-current-tab.zsh " ~ /.local/bin/claude-current-tab
~/.local/bin が PATH 上にあることを確認してください。必要に応じて、これをシェル プロファイルに追加します。
エクスポート PATH= " $HOME /.local/bin: $PATH "
確認:
~ /.local/bin/claude-current-tab --print-context
2. グローバル カーソル ルールをインストールする
mkdir -p ~ /.cursor/rules
cp カーソル/ルール/use-claude-chrome-extension.mdc ~ /.cursor/rules/
3. カーソル スラッシュ コマンドをインストールします。
mkdir -p ~ /.cursor/commands
cp カーソル/コマンド/use-claude-extension.md ~ /.cursor/commands/
ルールとコマンドをインストールした後、Cursor をリロードまたは再起動します。
スラッシュ コマンドは、次のようにカーソル チャットで使用できるようになります。
/use-claude-extension
セットアップに関する重要な注意事項: macOS のアクセス許可
このワークフローは既存のワークフローを制御するため、

ブラウザ UI を使用し、macOS アクセシビリティを通じてクロード サイド パネルを読み取る場合は、必要な macOS 権限を付与する必要があります。
システム設定
→ プライバシーとセキュリティ
→ アクセシビリティ
ヘルパー スクリプトを実行する可能性のあるすべてのアプリに対してアクセシビリティ アクセスを有効にします。
Cursor — Cursor エージェントがワークフローを実行するときに必要です。
Terminal 、 iTerm 、 Warp 、またはヘルパーを手動で実行するために使用するその他のターミナル アプリ。
スクリプトを起動する可能性のある別のローカル エージェントまたはハーネス アプリ。
macOS では、カーソルまたは端末がシステム イベントとブラウザを制御できるかどうかを尋ねるオートメーション プロンプトが表示される場合もあります。これらのプロンプトを許可します。
--debug-screenshots でデバッグ スクリーンショットを使用する場合は、以下も開きます。
システム設定
→ プライバシーとセキュリティ
→ 画面録画
カーソルやターミナルなどのヘルパーを実行しているアプリの画面録画を有効にします。
ワークフローがクロード サイド パネルを開いたものの入力に失敗した場合、またはプロンプトを送信したが応答を読み取ることができなかった場合、最も考えられる原因は、ヘルパーを起動したアプリのアクセシビリティ権限が欠落していることです。
デフォルトのブラウザで現在選択されているタブで関連ページを開きます。
カーソルには、Claude Chrome 拡張機能を使用するリクエストが表示されます。
~ /.local/bin/claude-current-tab --response-only " <Claude Chrome 拡張機能のタスク> "
ヘルパーは既存のデフォルト ブラウザをアクティブにし、現在のタブのクロード サイド パネルを開きます。
クロードは、ログインしたブラウザ セッションを使用してブラウザのみのコンテキストを分析します。
ヘルパーはクロードの最終的な答えを待って、それを Cursor に返します。
カーソルは結果を要約し、関連するリポジトリ ファイルを検査し、必要なコードまたはスクリプトの変更を加えます。
複雑なデバッグ タスクの場合は、次のようなプロンプトを使用します。
現在選択されているログ タブで Claude Chrome 拡張機能を使用します。ジョブがクラッシュする理由を分析します。 cを返す

実装に焦点を当てた簡潔なレポート:
- ログから観察された証拠
- おそらく根本原因
- カーソルが検査する必要がある正確なリポジトリ ファイル/関数/スクリプト
- 推奨されるコード変更
- あらゆる不確実性
次に、応答を使用して、このリポジトリに必要な変更を加えます。
インストールされたグローバル カーソル ルールは、ブラウザ分析をクロードに委任するときに、次の一般的なプロンプト形状を使用します。
カーソルから使用されています。 Claude Chrome 拡張機能とログインしたブラウザ セッションを使用して、現在選択されているブラウザ タブを分析します。観察された証拠、考えられる根本原因、カーソルで検査または変更する必要がある正確なファイル/関数/スクリプト、推奨されるコード変更、および不確実性を含む、実装に焦点を当てた簡潔なレポートを返します。ブラウザで要求されない限り、自分でコードを変更しないでください。回答を読んだ後、カーソルはコードの変更を適用します。
タスク: <ユーザータスク>
ヘルパーの使用法
現在のブラウザタブのコンテキストを出力します。
~ /.local/bin/claude-current-tab --print-context
プロンプトを送信し、クロードの回答のみを返します
~ /.local/bin/claude-current-tab --response-only \
" 現在選択されているタブを分析し、実用的な結果を要約します。"
プロンプトを送信して JSON メタデータを返す
~ /.local/bin/claude-current-tab --wait-response \
" 現在選択されているタブを分析し、実用的な結果を要約します。"
スクリーンショットを使用してデバッグする
~ /.local/bin/claude-current-tab --response-only --debug-screenshots \
--スクリーンショット ディレクトリ /tmp/claude-current-tab-debug \
「現在選択されているタブを分析します。」
スクリーンショットは視覚的なデバッグのみを目的としています。長いまたはスクロールするクロードの応答は、スクリーンショット OCR ではなく、クロードのサイド パネルから macOS アクセシビリティを通じて読み取られます。
推奨される統合順序:
グローバル カーソル ルール — 最適な即時統合。これにより、Cursor に使用を要求したときに何をすべきかを自動的に知らせます。

Claude Chrome 拡張機能、現在のブラウザー タブ、デフォルトのブラウザー セッション、ログインしているブラウザー アカウント。
スラッシュ コマンド — 便利な手動ヘルパー。このワークフローを明示的にトリガーする場合は、/use-claude-extension を使用します。
MCP ツール — Cursor に ask_claude_chrome_extension_current_tab などの正式な名前付きツールとして認識させたい場合に長期的に最適です。これにはより多くのセットアップが必要ですが、よりクリーンなツールスタイルの動作が得られます。
デフォルトのブラウザは Chromium と互換性があり、Claude Chrome 拡張機能がインストールされている必要があります。
ワークフローでは、現在選択されているブラウザー タブが使用されます。 Cursor に拡張機能の使用を依頼する前に、Claude に分析させたいタブが選択されていることを確認してください。
これにより、新しいブラウザは開かれず、 claude --chrome も使用されません。
長いまたはスクロールするクロードの応答は、スクリーンショット OCR ではなく、クロードのサイド パネルから macOS アクセシビリティを通じて読み取られます。
クロードがブラウザ側の許可を求めた場合は、ブラウザでそれを承認し、Cursor に続行するように指示します。
--response-only は、Cursor がクロードの回答を消費してコーディングを続行する必要がある場合に最適です。
デフォルトでは、応答待機は無制限であるため、長時間にわたるブラウザー分析が中断されることはありません。有限の上限が必要な場合は、 --response-timeout <秒> を明示的に渡します。
--wait-response は、アクティブなタブの URL/タイトルとクロードを含む JSON メタデータを返します。

[切り捨てられた]

## Original Extract

Contribute to bebrws/claude-chrome-extension-from-cursor development by creating an account on GitHub.

GitHub - bebrws/claude-chrome-extension-from-cursor · GitHub
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
bebrws
/
claude-chrome-extension-from-cursor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
bebrws/claude-chrome-extension-from-cursor
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit cursor cursor README.md README.md claude-current-tab.zsh claude-current-tab.zsh View all files Repository files navigation
Cursor Use of Claude Chrome Extension
I find Claude Chrome's extension to be much more powerful than Cursor's current browser integration. Especially if I need to debug something that requires being logged in to a system that uses SSO for example. This repository will allow you to use Claude's Chrome extension from Cursor against the currently selected tab in your already-open default browser. I created this so that I could take advantage of Claude's Chrome extension from Cursor: Cursor calls a local helper script, the script activates the existing default Chromium browser session, opens the Claude side panel, sends a prompt, reads Claude's final response through macOS Accessibility, and returns that response to Cursor so Cursor can continue coding or editing files.
Open the relevant page in your default browser and make sure it is the currently selected tab. Then ask Cursor to use the Claude Chrome extension.
/use-claude-extension to analyze the the logs for the current run of the do_something job. And make suggestions to improve the @do_something.py job.
The Cursor agent will then activate your system browser and in the currently open tab use the Claude Chrome extension. Cursor will wait for the response from the Claude Chrome extension using the script included in this repository.
Cursor is excellent at editing code in a repository, but sometimes the information needed to make the right code change only exists in an authenticated browser session.
This project lets a Cursor agent ask Claude's Chrome extension to inspect the currently selected browser tab , then brings Claude's answer back into Cursor so the agent can use that information to modify code.
It intentionally does not use claude --chrome , Playwright, or a newly launched browser. It uses your existing default browser window, current tab, installed Claude Chrome extension, and logged-in browser session.
Resolves your macOS default browser.
Activates the existing browser window.
Uses the currently selected tab.
Opens the Claude Chrome extension side panel.
Sends a prompt from Cursor to Claude in the browser.
Waits for Claude's final response.
Reads the response through macOS Accessibility.
Returns the response to Cursor so Cursor can continue the coding task.
.
├── README.md
├── claude-current-tab.zsh
└── cursor/
├── install.sh
├── bin/
│ └── claude-current-tab.zsh
├── rules/
│ └── use-claude-chrome-extension.mdc
└── commands/
└── use-claude-extension.md
Included files
claude-current-tab.zsh — the original helper script.
cursor/bin/claude-current-tab.zsh — bundled helper script used by the installer.
cursor/rules/use-claude-chrome-extension.mdc — Global Cursor rule that teaches Cursor when and how to use this workflow.
cursor/commands/use-claude-extension.md — optional Cursor slash command, available as /use-claude-extension after installation.
cursor/install.sh — one-step installer for the helper command, Global Cursor rule, and slash command.
The Cursor rule and slash command both invoke this helper command after install:
~ /.local/bin/claude-current-tab
Requirements
A Chromium-compatible default system browser, such as Brave Browser or Google Chrome.
The Claude Chrome extension installed in that default browser/profile.
The Claude extension shortcut for toggling the side panel must be available. In the tested setup, the extension command is toggle-side-panel and is bound to Cmd+E .
The browser tab you want Claude to analyze must already be the currently selected tab in the default browser.
The local helper script must be executable.
Safari and Firefox cannot host the Claude Chrome extension, so they cannot be used for this workflow.
From the root of this repository:
./cursor/install.sh
The installer is safe to re-run. It installs:
~/.local/bin/claude-current-tab
~/.cursor/rules/use-claude-chrome-extension.mdc
~/.cursor/commands/use-claude-extension.md
Then reload or restart Cursor so it picks up the new rule and command.
Verify the helper can read your current default-browser tab:
~ /.local/bin/claude-current-tab --print-context
You should see JSON containing the default browser, bundle id, active tab URL, and active tab title.
If you prefer not to use install.sh , install the pieces manually.
From the root of this repository:
chmod +x cursor/bin/claude-current-tab.zsh
mkdir -p ~ /.local/bin
ln -sf " $( pwd ) /cursor/bin/claude-current-tab.zsh " ~ /.local/bin/claude-current-tab
Make sure ~/.local/bin is on your PATH . If needed, add this to your shell profile:
export PATH= " $HOME /.local/bin: $PATH "
Verify:
~ /.local/bin/claude-current-tab --print-context
2. Install the Global Cursor rule
mkdir -p ~ /.cursor/rules
cp cursor/rules/use-claude-chrome-extension.mdc ~ /.cursor/rules/
3. Install the Cursor slash command
mkdir -p ~ /.cursor/commands
cp cursor/commands/use-claude-extension.md ~ /.cursor/commands/
After installing the rule and command, reload or restart Cursor.
The slash command should be available in Cursor chat as:
/use-claude-extension
Important setup note: macOS permissions
Because this workflow controls the existing browser UI and reads the Claude side panel through macOS Accessibility, you must grant the required macOS permissions.
System Settings
→ Privacy & Security
→ Accessibility
Enable Accessibility access for every app that may run the helper script:
Cursor — required when a Cursor agent runs the workflow.
Terminal , iTerm , Warp , or any other terminal app you use to run the helper manually.
Any separate local agent or harness app that may launch the script.
macOS may also show Automation prompts asking whether Cursor or your terminal can control System Events and your browser. Allow those prompts.
If you use debug screenshots with --debug-screenshots , also open:
System Settings
→ Privacy & Security
→ Screen Recording
Enable Screen Recording for the app running the helper, such as Cursor or Terminal.
If the workflow opens the Claude side panel but fails to type into it, or if it sends a prompt but cannot read the response, the most likely cause is missing Accessibility permission for the app that launched the helper.
You open the relevant page in the currently selected tab of your default browser.
Cursor sees your request to use the Claude Chrome extension.
~ /.local/bin/claude-current-tab --response-only " <task for Claude Chrome extension> "
The helper activates the existing default browser and opens the Claude side panel for the current tab.
Claude analyzes the browser-only context using your logged-in browser session.
The helper waits for Claude's final answer and returns it to Cursor.
Cursor summarizes the findings, inspects the relevant repo files, and makes any necessary code or script changes.
For complex debugging tasks, use a prompt like this:
Use the Claude Chrome extension on the currently selected logs tab. Analyze why the job is crashing. Return a concise implementation-focused report with:
- observed evidence from the logs
- likely root cause
- exact repo files/functions/scripts Cursor should inspect
- recommended code changes
- any uncertainty
Then use the response to make the necessary changes in this repo.
The installed Global Cursor rule uses this general prompt shape when delegating browser analysis to Claude:
You are being used from Cursor. Analyze the currently selected browser tab using the Claude Chrome extension and my logged-in browser session. Return a concise, implementation-focused report with: observed evidence, likely root cause, exact files/functions/scripts Cursor should inspect or modify, recommended code change, and any uncertainty. Do not make code changes yourself unless asked in the browser; Cursor will apply code changes after reading your answer.
Task: <user task>
Helper usage
Print the current browser-tab context
~ /.local/bin/claude-current-tab --print-context
Send a prompt and return only Claude's answer
~ /.local/bin/claude-current-tab --response-only \
" Analyze the currently selected tab and summarize the actionable findings. "
Send a prompt and return JSON metadata
~ /.local/bin/claude-current-tab --wait-response \
" Analyze the currently selected tab and summarize the actionable findings. "
Debug with screenshots
~ /.local/bin/claude-current-tab --response-only --debug-screenshots \
--screenshot-dir /tmp/claude-current-tab-debug \
" Analyze the currently selected tab. "
Screenshots are for visual debugging only. Long or scrolling Claude responses are read through macOS Accessibility from the Claude side panel, not by screenshot OCR.
Recommended integration order:
Global Cursor rule — best immediate integration. It lets Cursor automatically know what to do when you ask it to use the Claude Chrome extension, current browser tab, default browser session, logged-in browser accounts.
Slash command — useful manual helper. Use /use-claude-extension when you want to explicitly trigger this workflow.
MCP tool — best long-term if you want Cursor to see this as a formal named tool, such as ask_claude_chrome_extension_current_tab . This requires more setup but gives cleaner tool-style behavior.
The default browser must be Chromium-compatible and have the Claude Chrome extension installed.
The workflow uses the currently selected browser tab. Make sure the tab you want Claude to analyze is selected before asking Cursor to use the extension.
This does not open a new browser and does not use claude --chrome .
Long or scrolling Claude responses are read through macOS Accessibility from the Claude side panel, not by screenshot OCR.
If Claude asks for browser-side permission, approve it in the browser and then tell Cursor to continue.
--response-only is best when Cursor needs to consume Claude's answer and continue coding.
By default, response waiting is unlimited so long browser analyses are not cut off. If you want a finite cap, pass --response-timeout <seconds> explicitly.
--wait-response returns JSON metadata, including the active tab URL/title, along with the Claude

[truncated]
