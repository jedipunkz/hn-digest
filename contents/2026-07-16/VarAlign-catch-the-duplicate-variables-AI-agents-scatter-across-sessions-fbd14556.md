---
source: "https://github.com/greenyogainc/varalign"
hn_url: "https://news.ycombinator.com/item?id=48937123"
title: "VarAlign – catch the duplicate variables AI agents scatter across sessions"
article_title: "GitHub - greenyogainc/varalign: VarAlign — catch the duplicate, drifted, and misaligned variables AI coding agents scatter across sessions. VS Code extension (source-available, BSL 1.1). · GitHub"
author: "andreab67"
captured_at: "2026-07-16T17:04:10Z"
capture_tool: "hn-digest"
hn_id: 48937123
score: 1
comments: 0
posted_at: "2026-07-16T17:00:31Z"
tags:
  - hacker-news
  - translated
---

# VarAlign – catch the duplicate variables AI agents scatter across sessions

- HN: [48937123](https://news.ycombinator.com/item?id=48937123)
- Source: [github.com](https://github.com/greenyogainc/varalign)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T17:00:31Z

## Translation

タイトル: VarAlign – AI エージェントがセッション間で散らばる重複変数を捕捉する
記事のタイトル: GitHub - greenyogainc/varalign: VarAlign — AI コーディング エージェントがセッション全体に散在する、重複、ドリフト、および不整合な変数を捕捉します。 VS Code 拡張機能 (ソースが入手可能、BSL 1.1)。 · GitHub
説明: VarAlign — AI コーディング エージェントがセッション全体に散在する、重複、ドリフト、および不整合な変数を捕捉します。 VS Code 拡張機能 (ソースが入手可能、BSL 1.1)。 - greenyogainc/varalign

記事本文:
GitHub - greenyogainc/varalign: VarAlign — AI コーディング エージェントがセッション全体に散在する、重複、ドリフト、および不整合な変数を捕捉します。 VS Code 拡張機能 (ソースが入手可能、BSL 1.1)。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。 P

このページをリロードしてください。
グリーンヨガインク
/
バラライン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット メディア メディア スクリプト スクリプト src src .gitignore .gitignore .vscodeignore .vscodeignore ライセンス ライセンス README.md README.md esbuild.js esbuild.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json 表示すべてのファイル リポジトリ ファイルのナビゲーション
VarAlign VS Code 拡張機能のパブリック ソースで利用可能なミラー (ビジネス ソース ライセンス 1.1)。 VS Code Marketplace / Open VSX からインストールします。
VS Code マーケットプレイスにもあります。
AI コーディング エージェントがセッション間で散らばる重複変数、ドリフト変数、位置ずれ変数をエディタ内で直接捕捉します。 100% ローカル: コードがマシンから離れることはありません。
AIアシスタントは忘れてしまいます。セッション間で、すでに使用されている変数が再導入されます。
別の名前で存在する場合、先週書き込まれた値から値を変更します。
または、ファイルの移動時に定義が取り残されます。 VarAlign は、すべての割り当てを追跡します。
アシスタントが書き込み、重複とドリフトにスコアを付けて、すぐに貼り付けられるファイルを渡します
プロンプトを修正します。
VarAlign 上のネイティブ ツリー ビュー — 重複、変数、セッション —
エンジン。依存関係のない Python で実行される検出、スコアリング、永続化
拡張機能内に同梱されるエンジン。ビューは薄い、読み飛ばされます
クライアント。すべてがマシン上で実行されます。クラウドもテレメトリもコードも必要ありません
アップロード - そのため、ロックダウンされたエアギャップ環境でも機能します。
重複 — 高 / 中 / 低グループ。ペアを展開すると両側が表示されます。
右クリックして [閉じる (重複ではない)] 、 [確認] 、または [閉じる] をクリックします。
注意事項評決は継続し、メンバーの 1 人を解任すると自動的に沈黙する

全体
各パターンを一度確認してください。
変数 — 追跡されたすべての割り当て。ファイルごとにグループ化され、その内容ごとに色分けされます。
最悪の重複レベル。クリックすると定義にジャンプします。
セッション — 各 AI セッションで導入または変更されたもの。
ステータス バー — VarAlign: N 高 ;クリックして「重複」ビューにフォーカスします。
修正プロンプトの生成 - 新しいタブにリポジトリを対象とした修復プロンプトが表示されます。
アシスタントに貼り付ける準備ができました。
AI で修正 — ターゲットを絞った統合プロンプトを Claude Code に渡すか、
Kilo コード、開いているものを選択します。
ストアが変更されると（フックまたは CLI がストアに書き込んだ場合）、自動更新されます。
Python 3.11+ が PATH 上にあることを確認してください。拡張機能にはエンジンがバンドルされています。
したがって、他にインストールしたり指定したりするものは何もありません。
リポジトリを開き、アクティビティ バーの VarAlign チップをクリックします。 VarAlign
ワークスペースをスキャンして追跡を開始します。
それだけです。ローカルで実行されており、すべてのバイトがマシン上に残ります。
VarAlign は、追跡データをリポジトリ ルートの .varmem/ フォルダーに保存します。の
拡張子は .varmem/ を .gitignore に自動的に追加します (git リポジトリのみ)。
それができない場合は、次の行を自分で追加してください。
Pro は変数のマージのロックを解除します。重複したペアを右クリックし、VarAlign を選択します。
正規名を指定し、参照を書き換えて、重複を削除します。
定義 — エディター、マシン上で。
ライセンスはオフラインで検証されます (Ed25519 署名付きキー、ローカルでチェックされます)
有効期限を過ぎても 14 日間の猶予期間 — どこにも何も送信されないため、Pro が機能します
エアギャップもある）。 VarAlign でアクティブ化: License を入力します。いつでもチェックしてください
VarAlign: ライセンス ステータス。
設定
デフォルト
意味
varalign.pythonPath
パイソン
ローカルエンジンを実行するために使用されるインタープリタ
varalign.corePath
(同梱)
varmem.py へのパス。空 = バンドルされたエンジン
varalign.minLevel
中程度
最低の重複レベルが表示されています
varalign.showDismissed
ファ

そうだね
却下/自動停止されたペアを含む
varalign.aiTool
自動
AI による修正のアシスタント (自動 / クロード / キロ)
varalign.licenseKey
「」
Pro ライセンス キー (VL1.…)、オフラインで検証済み
varalign.apiUrl
「」
オプションのセルフホスト API (以下を参照)
varalign.apiToken
「」
その API のベアラー トークン
varalign.apiプロジェクト
「」
その API のプロジェクト ID
varalign.apiAllowInsecure
偽
TLS 検証をスキップ (内部 CA のみ)
オプション: セルフホスト API からの読み取り
エンジンを集中的に実行するチーム ( python varmem.py は独自に機能します)
インフラストラクチャ) varalign.apiUrl を設定することでビューを指すことができます。
varalign.apiToken および varalign.apiProject 。その後、ビューが読み取られます
/v1/プロジェクト/{id}/レポート ;レビュー判定は API 経由で読み取り専用のままです。休暇を取る
完全にローカルなエクスペリエンスを実現するには、apiUrl を空 (デフォルト) にします。
cd拡張子
npmインストール
npm run apply # 拡張機能開発ホストの場合は F5 キーを押します
.vsix (vscode:prepublish による縮小バンドル + エンジン) をビルドします。
npx @vscode/vsce パッケージ
ライセンス
VarAlign VS Code 拡張機能は、ビジネス ソースからソースで入手できます。
ライセンス 1.1 (LICENSE ファイルを参照): 使用および変更は自由、再販または再販は禁止
競合するホスト型/組み込み型製品は、2030 年 7 月 15 日に Apache-2.0 に変換されます。の
基礎となる VarAlign エンジンは、Apache-2.0 に基づいて個別にライセンスされます。
VarAlign — AI コーディング エージェントがセッション全体に散在する、重複、ドリフト、および不整合な変数を捕捉します。 VS Code 拡張機能 (ソースが入手可能、BSL 1.1)。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

VarAlign — catch the duplicate, drifted, and misaligned variables AI coding agents scatter across sessions. VS Code extension (source-available, BSL 1.1). - greenyogainc/varalign

GitHub - greenyogainc/varalign: VarAlign — catch the duplicate, drifted, and misaligned variables AI coding agents scatter across sessions. VS Code extension (source-available, BSL 1.1). · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
greenyogainc
/
varalign
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits media media scripts scripts src src .gitignore .gitignore .vscodeignore .vscodeignore LICENSE LICENSE README.md README.md esbuild.js esbuild.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Public source-available mirror of the VarAlign VS Code extension (Business Source License 1.1). Install from the VS Code Marketplace / Open VSX.
Also on the VS Code Marketplace .
Catch the duplicate, drifted, and misaligned variables AI coding agents scatter across sessions — right in your editor. 100% local: your code never leaves your machine.
AI assistants forget. Across sessions they re-introduce a variable that already
exists under another name, let a value drift from the one they wrote last week,
or strand a definition when a file moves. VarAlign tracks every assignment your
assistant writes, scores the duplicates and drift, and hands you a ready-to-paste
fix prompt.
Native tree views — Duplicates · Variables · Sessions — over the VarAlign
engine. Detection, scoring, and persistence run in a zero-dependency Python
engine that ships inside the extension; the views are a thin, read-through
client. Everything runs on your machine — no cloud, no telemetry, no code
upload — so it works in locked-down and air-gapped environments.
Duplicates — High / Medium / Low groups. Expand a pair to see both sides;
right-click to Dismiss (not a duplicate) , Confirm , or Dismiss with
note . Verdicts persist, and dismissing one member auto-quiets the whole
family so you review each pattern once.
Variables — every tracked assignment, grouped by file and coloured by its
worst duplicate level; click to jump to the definition.
Sessions — what each AI session introduced or changed.
Status bar — VarAlign: N high ; click to focus the Duplicates view.
Generate Fix Prompt — a repo-scoped remediation prompt in a new tab,
ready to paste back to your assistant.
Fix with AI — hands a targeted consolidation prompt to Claude Code or
Kilo Code, whichever you have open.
Auto-refreshes when the store changes (a hook or the CLI wrote to it).
Make sure Python 3.11+ is on your PATH — the extension bundles the engine,
so there's nothing else to install or point at.
Open a repo and click the VarAlign chip in the activity bar. VarAlign
scans the workspace and starts tracking.
That's it — you're running locally, and every byte stays on your machine.
VarAlign keeps its tracking data in a .varmem/ folder at your repo root. The
extension adds .varmem/ to your .gitignore automatically (git repos only).
If it can't, add this line yourself:
Pro unlocks Merge Variables : right-click a duplicate pair and VarAlign picks
the canonical name, rewrites the references, and removes the duplicate
definition — in your editor, on your machine.
Licenses are verified offline (an Ed25519-signed key, checked locally with a
14-day grace period past expiry — nothing is ever sent anywhere, so Pro works
air-gapped too). Activate with VarAlign: Enter License ; check anytime with
VarAlign: License Status .
Setting
Default
Meaning
varalign.pythonPath
python
interpreter used to run the local engine
varalign.corePath
(bundled)
path to varmem.py ; empty = the bundled engine
varalign.minLevel
medium
lowest duplicate level shown
varalign.showDismissed
false
include dismissed / auto-quieted pairs
varalign.aiTool
auto
assistant for Fix with AI ( auto / claude / kilo )
varalign.licenseKey
""
Pro license key ( VL1.… ), verified offline
varalign.apiUrl
""
optional self-hosted API (see below)
varalign.apiToken
""
bearer token for that API
varalign.apiProject
""
project id on that API
varalign.apiAllowInsecure
false
skip TLS verify (internal CA only)
Optional: read from a self-hosted API
Teams that run the engine centrally ( python varmem.py serve on their own
infrastructure) can point the views at it by setting varalign.apiUrl ,
varalign.apiToken , and varalign.apiProject . The views then read
/v1/projects/{id}/report ; review verdicts stay read-only over the API. Leave
apiUrl empty (the default) for the fully local experience.
cd extension
npm install
npm run compile # press F5 for an Extension Development Host
Build a .vsix (minified bundle + engine via vscode:prepublish ):
npx @vscode/vsce package
License
The VarAlign VS Code extension is source-available under the Business Source
License 1.1 (see the LICENSE file): free to use and modify, no reselling or
competing hosted/embedded offering, converting to Apache-2.0 on 2030-07-15. The
underlying VarAlign engine is licensed separately under Apache-2.0.
VarAlign — catch the duplicate, drifted, and misaligned variables AI coding agents scatter across sessions. VS Code extension (source-available, BSL 1.1).
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
