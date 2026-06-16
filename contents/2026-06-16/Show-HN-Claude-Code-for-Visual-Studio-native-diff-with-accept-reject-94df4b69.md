---
source: "https://github.com/firish/claude_code_vs"
hn_url: "https://news.ycombinator.com/item?id=48548381"
title: "Show HN: Claude Code for Visual Studio (native diff with accept/reject)"
article_title: "GitHub - firish/claude_code_vs · GitHub"
author: "firish"
captured_at: "2026-06-16T01:08:53Z"
capture_tool: "hn-digest"
hn_id: 48548381
score: 9
comments: 1
posted_at: "2026-06-15T23:15:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Code for Visual Studio (native diff with accept/reject)

- HN: [48548381](https://news.ycombinator.com/item?id=48548381)
- Source: [github.com](https://github.com/firish/claude_code_vs)
- Score: 9
- Comments: 1
- Posted: 2026-06-15T23:15:56Z

## Translation

タイトル: HN を表示: Visual Studio のクロード コード (受け入れ/拒否を伴うネイティブの差分)
記事のタイトル: GitHub - firish/claude_code_vs · GitHub
説明: GitHub でアカウントを作成して、firish/claude_code_vs の開発に貢献します。
HN テキスト: VS Code と JetBrains には正式な Claude Code IDE 統合がありますが、Visual Studio にはありません。多くの +1 が付けられた未解決の GitHub の問題があるため、自分でビルドしました (視覚的な例については、GitHub リポジトリの GIF を確認してください)。公式プラグインが使用するものと同じプロトコルを実装しているため、Claude CLI は自動的に接続します。何も設定せず、拡張機能をインストールして「起動」をクリックするだけです。ターミナルで Claude を実行する場合に追加される主な点は、編集内容がターミナルで自動適用されたりプロンプトを表示されたりするのではなく、Visual Studio のネイティブの差分ビューアで開かれることです。そこで[承認]または[拒否]をクリックします。理由を付けて拒否することもでき、クロードは別のパスを受け取ります。また、コンパイラ エラー (C# および C++) と現在の選択内容を CLI と自動的に共有するため、何もコピーして貼り付けることなく、Claude にコンテキストが表示されます。他にもいくつか:
- セッションの接続ステータスとトークン/コスト統計を表示するドッキング可能なパネルがあります
- 差分を開かずに編集を自動的に受け入れる「run wild」トグル
既存の claude CLI で動作し、独自のモデル呼び出しはありません マーケットプレイス: https://marketplace.visualstudio.com/items?itemName=firish.b... GitHub: https://github.com/firish/claude_code_vs

記事本文:
GitHub - firish/claude_code_vs · GitHub
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
猟奇的な
/
クロード_コード_対
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
1 コミット 1 コミット デモ デモ docs docs スパイク スパイク src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md ClaudeCodeVS.slnx ClaudeCodeVS.slnx ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Claude コードを Visual Studio 2026 に導入 - 受け入れ/拒否、自動選択 + コンパイラ診断コンテキスト、およびライブ統計パネルを備えたネイティブの差分ウィンドウ。クロード CLI はエージェントの作業を行います。この拡張機能は、Claude Code の統合プロトコルの IDE の半分です。
ステータス: コミュニティ プロジェクト。Anthropic とは関係ありません。現時点では Visual Studio 2026 のみ。
Claude Code には、VS Code と JetBrains 用のファーストクラスの IDE 統合がありますが、Visual Studio にはありません - anthropics/claude-code#15942 を参照してください。この拡張機能は、同じ IDE 統合プロトコルを VS にネイティブに実装するため、ターミナルにコピー＆ペーストする代わりに、CLI が実際の Visual Studio の差分ウィンドウを駆動し、選択内容とビルド エラーを確認します。
単一の受け入れ/拒否ゲートを備えたネイティブ diff - クロードの編集は Visual Studio の差分ビューアで開き、承認することが唯一のステップです (ターミナルに重複する y/n プロンプトはありません)。
フィードバック付きで拒否 - 編集を拒否し、何を変更するかをクロードに伝えます。あなたのメモで再考します。
Run wild (auto-accept) - クックさせたい場合に、差分を開かずに編集を適用するためのパネル切り替えです。各セッションをリセットします。
診断の共有 - クロードは、Visual Studio のコンパイラ エラー/警告 (C# および C++) を読み取り、それらを修正します。
選択コンテ​​キスト - クロードは、あなたが見ているファイルと行を自動的に認識します。
ライブ パネル - ドッキング可能なクロード コード パネル: 接続ステータス、編集決定、トークン使用量 + 推定コスト (最新の呼び出しと累積セッション)。
Claude Code CLI はインストールされ、認証されています - Claude を参照してください。

コードドキュメント 。この拡張機能はモデル呼び出しを行わず、エージェント自体も動作しません。 CLI が必要です。
クロード 2.1.173 に対してテストされました。
マーケットプレイス: [拡張機能] -> [拡張機能の管理] で「Claude Code for Visual Studio」を検索するか、Visual Studio Marketplace からインストールします。
サイドロード: リリースから .vsix をダウンロードし、ダブルクリックします。
Visual Studio 2026 でプロジェクトまたはソリューションを開きます。
[クロード コード] パネル ( [表示] -> [その他のウィンドウ] -> [クロード コード]) を開き、[クロード コードの起動] をクリックします ([ツール] メニューからも使用できます)。
すでに IDE に接続されている claude を実行するターミナルが開きます (/ide は必要ありません)。パネルの丸印が緑色に変わります - 接続されました。
クロードに変更を依頼してください。その編集は差分として開きます - [承認] 、 [拒否] 、または [フィードバック付き拒否] をクリックします。
診断では、コンパイラーが分析できるように、ロードされたプロジェクト (オープンフォルダー モードのルーズ ファイルではない) が必要です。
これはプロトコル ブリッジであり、エージェントの再実装ではありません。拡張機能の起動時:
localhost WebSocket サーバーを起動し、 ~/.claude/ide/<port>.lock にロックファイルを書き込みます。
ENABLE_IDE_INTEGRATION とブリッジ ポートを使用して claude を起動するため、自動接続され、ソケット経由で MCP / JSON-RPC を通信します。
CLI が駆動する IDE ツール ( openDiff 、 openFile 、 getDiagnostics 、選択の更新、および diff-tab ライフサイクル) を実装します。
VS diff を単一の承認ゲートにするために、拡張機能は提案された編集を diff 経由でルーティングする小さな PreToolUse フックをワークスペースの .claude/settings.json にインストールします。 CLI はすべてのエージェントの作業を実行します。拡張機能はモデル呼び出しを行うことはありません。
ブリッジは 127.0.0.1 のみにバインドし、接続ごとに (ロックファイルからの) 認証トークンを検証します。トークンはログに記録されません。
この拡張機能は、ネットワーク呼び出しや独自の LLM 呼び出しを行いません。すべての AI 作業は、独自の認証のもとで、claude CLI によって行われます。
○

n Launch を実行すると、これらがワークスペースの .claude/ フォルダに書き込まれ、フック エントリが .claude/settings.json にマージされます (既存のコンテンツは保持されます)。
vs-permission-hook.ps1 - VS diff を介して編集/書き込み/マルチエディット編集をルーティングします。
vs-usage-hook.ps1 - パネルにトークンの統計を表示できるようにトランスクリプト パスを報告します。
トークンコストは推定値 (ハードコードされた層ごとの価格) であり、[推定コストの表示] をクリックした場合にのみ表示されます。
現時点では Visual Studio 2026 のみ (需要があれば VS 2022 バックフィルが計画されています)。
IDE 統合プロトコルは文書化されておらず、バージョンが脆弱であるため、クロードのアップデートにより変更される可能性があります。正常な動作: 2.1.173。
診断にはロードされたプロジェクトが必要です (エラー リスト/Roslyn はルーズ ファイルを分析しません)。
トークン統計は編集時に更新されるため (信頼性の高いフック トリガー)、チャットのみのターンではすぐに更新されない可能性があります。
コストの数値は見積もりであり、請求額ではありません。
パネルに「CLI を待機しています」と表示されます。[Launch Claude Code] をクリックするか、claude ターミナルで /ide を実行して Visual Studio を選択します。
新しいファイルが間違ったフォルダーに配置されます。拡張機能から起動するか (作業ディレクトリをワークスペースに固定します)、リポジトリ内から claude を実行します。
getDiagnostics は何も返しません。コードをプロジェクトとして開き、エラーが [エラー リスト] に表示されることを確認します。
バグを報告する: [出力] -> [Claude Code] ペインの内容と claude --version を含めます。
Visual Studio 拡張機能開発ワークロードを備えた Visual Studio 2026 が必要です。
msbuild src / ClaudeCodeVS / ClaudeCodeVS.csproj / t:Rebuild / p:Configuration = Release
F5 を押して (または devenv /rootsuffix Exp を起動して)、実験的インスタンスでデバッグします。アーキテクチャと貢献者のガイダンスは CLAUDE.md にあります。今後の作業は ROADMAP.md に保存されます。
問題やPRを歓迎します。プロトコル契約は文書化されておらず、以前にリグレッションしました - クロード CLI を使用した場合は、スパイクを実行してください

スモーク テスト (spike/) を実行し、バージョンをメモします。
MIT © 2026 リシ・グラティ。アンソロピックとは無関係です。 「Claude」および「Claude Code」は Anthropic の商標であり、ここでは相互運用性を説明するためにのみ使用されます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
v1.0.1
最新の
2026 年 6 月 15 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to firish/claude_code_vs development by creating an account on GitHub.

VS Code and JetBrains have official Claude Code IDE integration, but Visual Studio doesn't. There's an open GitHub issue with a lot of +1s so I built it myself (check out the gif in the GitHub repo for a visual example). It implements the same protocol the official plugins use, so the Claude CLI connects to it automatically. You don't configure anything, just install the extension and click Launch. The main thing it adds over running Claude in a terminal is that edits open in Visual Studio's native diff viewer instead of auto-applying or prompting you in the terminal. You click Accept or Reject right there. You can also reject with a reason, and Claude will take another pass. It also shares your compiler errors (C# and C++) and your current selection with the CLI automatically, so Claude has context without you having to copy-paste anything. A few other things:
- There's a dockable panel with connection status and token/cost stats for the session
- A "run wild" toggle to auto-accept edits without opening the diff
Works with the existing claude CLI, no model calls of its own Marketplace: https://marketplace.visualstudio.com/items?itemName=firish.b... GitHub: https://github.com/firish/claude_code_vs

GitHub - firish/claude_code_vs · GitHub
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
firish
/
claude_code_vs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit demo demo docs docs spike spike src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md ClaudeCodeVS.slnx ClaudeCodeVS.slnx LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md View all files Repository files navigation
Bring Claude Code into Visual Studio 2026 - a native diff window with accept/reject, automatic selection + compiler-diagnostics context, and a live stats panel. The claude CLI does the agent work; this extension is the IDE half of Claude Code's integration protocol.
Status: community project, not affiliated with Anthropic. Visual Studio 2026 only , for now.
Claude Code has first-class IDE integration for VS Code and JetBrains, but not Visual Studio - see anthropics/claude-code#15942 . This extension implements that same IDE-integration protocol natively for VS, so the CLI drives a real Visual Studio diff window and sees your selection and build errors - instead of you copy-pasting into a terminal.
Native diff with a single accept/reject gate - Claude's edits open in Visual Studio's diff viewer, and approving there is the only step (no duplicate y/n prompt in the terminal).
Reject with feedback - reject an edit and tell Claude what to change; it reconsiders with your note.
Run wild (auto-accept) - a panel toggle to apply edits without opening the diff, for when you want to let it cook. Resets each session.
Diagnostics sharing - Claude reads Visual Studio's compiler errors/warnings (C# and C++) and fixes them.
Selection context - Claude automatically knows the file and lines you're looking at.
Live panel - a dockable Claude Code panel: connection status, edit decisions, and token usage + estimated cost (latest call vs cumulative session).
The Claude Code CLI , installed and authenticated - see the Claude Code docs . This extension makes no model calls and does no agent work itself; it requires the CLI.
Tested against claude 2.1.173 .
Marketplace: search "Claude Code for Visual Studio" in Extensions -> Manage Extensions , or install from the Visual Studio Marketplace .
Sideload: download the .vsix from Releases and double-click it.
Open your project or solution in Visual Studio 2026.
Open the Claude Code panel ( View -> Other Windows -> Claude Code ) and click Launch Claude Code (also available on the Tools menu).
A terminal opens running claude , already connected to the IDE (no /ide needed). The panel pill turns green - Connected .
Ask Claude to make a change. Its edit opens as a diff - click Accept , Reject , or Reject with feedback… .
Diagnostics need a loaded project (not a loose file in Open-Folder mode) for the compiler to analyze it.
This is a protocol bridge , not a re-implementation of the agent. On launch the extension:
Starts a localhost WebSocket server and writes a lockfile at ~/.claude/ide/<port>.lock .
Launches claude with ENABLE_IDE_INTEGRATION and the bridge port, so it auto-connects and speaks MCP / JSON-RPC over the socket.
Implements the IDE tools the CLI drives - openDiff , openFile , getDiagnostics , selection updates, and diff-tab lifecycle.
To make the VS diff the single approval gate , the extension installs a small PreToolUse hook into your workspace's .claude/settings.json that routes proposed edits through the diff. The CLI does all agent work; the extension never makes model calls.
The bridge binds to 127.0.0.1 only and validates an auth token (from the lockfile) on every connection. The token is never logged.
The extension makes no network calls and no LLM calls of its own . All AI work is the claude CLI, under your own authentication.
On Launch , it writes these into your workspace's .claude/ folder and merges hook entries into .claude/settings.json (preserving existing content):
vs-permission-hook.ps1 - routes Edit/Write/MultiEdit edits through the VS diff.
vs-usage-hook.ps1 - reports the transcript path so the panel can show token stats.
Token cost is an estimate (hardcoded per-tier prices), shown only when you click Show est. cost .
Visual Studio 2026 only for now (a VS 2022 backfill is planned if there's demand).
The IDE-integration protocol is undocumented and version-fragile - a claude update could change it. Known-good: 2.1.173.
Diagnostics need a loaded project (the Error List / Roslyn won't analyze loose files).
Token stats refresh on edits (the reliable hook trigger), so a chat-only turn may not update them immediately.
Cost figures are estimates , not billing.
Panel says "Waiting for CLI": click Launch Claude Code , or run /ide in a claude terminal and pick Visual Studio .
New files land in the wrong folder: launch from the extension (it pins the working directory to your workspace), or run claude from inside the repo.
getDiagnostics returns nothing: open the code as a project and confirm the error appears in the Error List .
Filing a bug: include the Output -> Claude Code pane contents and your claude --version .
Requires Visual Studio 2026 with the Visual Studio extension development workload.
msbuild src / ClaudeCodeVS / ClaudeCodeVS.csproj / t:Rebuild / p:Configuration = Release
Press F5 (or launch devenv /rootsuffix Exp ) to debug in the experimental instance. Architecture and contributor guidance live in CLAUDE.md . Future work lives in ROADMAP.md .
Issues and PRs welcome. The protocol contract is undocumented and has regressed before - if you bump the claude CLI, please run the spike smoke test ( spike/ ) and note the version.
MIT © 2026 Rishi Gulati. Not affiliated with Anthropic. "Claude" and "Claude Code" are trademarks of Anthropic, used here only to describe interoperability.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
v1.0.1
Latest
Jun 15, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
