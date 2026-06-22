---
source: "https://github.com/av29nassh-sketch/PreFlight"
hn_url: "https://news.ycombinator.com/item?id=48632054"
title: "Show HN: PreFlight – A local AST scanner to catch AI architectural drift"
article_title: "GitHub - av29nassh-sketch/PreFlight: The local security gate for AI-generated code. · GitHub"
author: "Avenassh"
captured_at: "2026-06-22T16:33:44Z"
capture_tool: "hn-digest"
hn_id: 48632054
score: 3
comments: 0
posted_at: "2026-06-22T16:03:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PreFlight – A local AST scanner to catch AI architectural drift

- HN: [48632054](https://news.ycombinator.com/item?id=48632054)
- Source: [github.com](https://github.com/av29nassh-sketch/PreFlight)
- Score: 3
- Comments: 0
- Posted: 2026-06-22T16:03:31Z

## Translation

タイトル: Show HN: PreFlight – AI アーキテクチャのドリフトを捕捉するローカル AST スキャナー
記事のタイトル: GitHub - av29nassh-sketch/PreFlight: AI 生成コードのローカル セキュリティ ゲート。 · GitHub
説明: AI 生成コード用のローカル セキュリティ ゲート。 GitHub でアカウントを作成して、av29nassh-sketch/PreFlight の開発に貢献してください。

記事本文:
GitHub - av29nassh-sketch/PreFlight: AI 生成コードのローカル セキュリティ ゲート。 · GitHub
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
av29nassh-スケッチ
/
プリフライト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
79 コミット 79 コミット .claude .claude .preflight-playground/ app/ api/ users .preflight-playground/ app/ api/ users .vscode .vscode アセット アセット beta-reasoning-fixture beta-reasoning-fixture docs docshardcore-test-suite Hardcore-test-suite playground-ast-only playground-ast-only playground playground preflight-dashboard preflight-dashboard preflight-proxy preflight-proxy qa/ preflight-gauntlet qa/ preflight-gauntlet real-flow-fixture real-flow-fixture スクリプト スクリプト サーバー/ チェックアウト サーバー/ チェックアウト src src supabase supabase テスト テスト テスト テスト wasm wasm .cursorrules .cursorrules .gitignore .gitignore .npmignore .npmignore .vercelignore .vercelignore LICENSE.md LICENSE.md PreFlight-Founder-Playbook.md PreFlight-Founder-Playbook.md PreFlight-Founder-Playbook.pdf PreFlight-Founder-Playbook.pdf README.md README.md cli.js cli.js デモ.gif デモ.gif e2e-license-live-test.js e2e-license-live-test.js Generate-tests.js Generate-tests.js Index.js Index.js logger.js logger.js mcp-instructions.md mcp-instructions.md package-lock.json package-lock.json package.json package.json remediationEngine.js remediationEngine.jsリセット-デモ.js リセット-デモ.js scaffoldEngine.js scaffoldEngine.js taintTracker.js taintTracker.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング ドリフトが本番環境の技術的負債になる前に阻止します。 PreFlight は、AI 生成コード用のローカル ファースト セーフティ ゲートであり、安全でない認証、RLS、SQL、SSRF、コマンド実行、依存関係、シークレット処理の変更をコミットされる前に検出するように構築されています。
ウェブサイト: https://preflight-vibe.vercel.app
PreFlight は、コードベースのニーズに応じて 2 つの異なる層で実行されます。
機能: 無制限のローカル スキャンと、ローカルのデータ全体にわたる 10 個の無料パッチ アプリケーション

rministic 修正とプロキシ支援 AI 修正。
セットアップ: スキャン用の設定は不要です。 Pro キーは、10 個の無料パッチを使用した後にのみ必要になります。
npm install -g プリフライトプロ
プリフライト初期化
プリフライトスキャン。 --修正
preflight-pro をインストールすると、シェルでユニバーサル プリフライト コマンドが公開されます。
機能: 無制限のスキャンと無制限の修正。これには、複雑な複数ファイルのアーキテクチャ上の欠陥に対する深い推論による修復、テナント分離ロジック、パラメトリック SQL インジェクションが含まれます。
セットアップ: アクティブな PREFLIGHT_PRO_KEY またはプリフライト auth から保存されたキーが必要です。
$ env: PREFLIGHT_PRO_KEY = " PREFLIGHT-BETA-XXXXX "
プリフライトスキャン。 -- 修正します
バッシュ/macOS:
import PREFLIGHT_PRO_KEY= " PREFLIGHT-BETA-XXXXX "
プリフライトスキャン。 --修正
インストールの流れ
PreFlight は、ターミナルファーストのワークフローと IDE ファーストのワークフローの両方をサポートします。どちらのパスも preflight init で終わります。これは、このウィザードがエディター、MCP クライアント、および Pro/Beta キーを 1 か所で接続するためです。
npm install -g プリフライトプロ
プリフライト初期化
次に、任意のプロジェクトをルートからスキャンします。
プリフライトスキャン。 --修正
パス B: VS コード / カーソル
グローバル CLI コマンドをインストールします。 VSIX はエディター内 UI を提供しますが、拡張機能は引き続きグローバル プリフライト コマンドを使用して The Eye デーモンを起動し、修正を実行します。
npm install -g プリフライトプロ
PreFlight Companion VSIX 拡張機能をダウンロードしてインストールします。
PreFlight Web サイトから VSIX をダウンロードする
または、GitHub リリースを開いて、最新のプリフライト コンパニオン VSIX をインストールします。
VS Code または Cursor で、[拡張機能] パネルを開き、[...] メニューをクリックして、[VSIX からインストール...] を選択し、ダウンロードしたファイルを選択します。
プリフライト初期化
IDE でプロジェクトを開きます。この拡張機能は The Eye を自動的に起動し、ファイルの保存を監視し、PreFlight アラートをエディター内に表示します。
The Eye: VS Code/Cursor 拡張機能は、PreFlight のローカル デーモンを自動的に起動します。

ファイルの保存を監視し、AI が生成したコードによってハードブロックの問題が発生すると、エディタ内でアラートを生成します。
MCP ブリッジ: preflight init は、サポートされている AI エディターに preflight mcp を接続することもできるため、エージェントはコーディング フローを離れることなく PreFlight ツールを呼び出すことができます。
無料ユーザーは無制限のスキャンと、ローカル修正およびプロキシ支援 AI 修正にわたる合計 10 個のパッチを利用できます。 10 個の無料パッチを使用した後は、無制限の修正を行うには Pro/Beta キーが必要です。
プリフライト init 中にキーを追加することも、直接アクティブ化することもできます。
プリフライト認証 PREFLIGHT-BETA-XXXXX
1 つのターミナル セッションについては、手動で設定することもできます。
$ env: PREFLIGHT_PRO_KEY = " PREFLIGHT-BETA-XXXXX "
import PREFLIGHT_PRO_KEY= " PREFLIGHT-BETA-XXXXX "
価格設定
無料利用枠: 無制限のスキャン、10 個の無料パッチ (ローカル + Deep-Reasoning AI)。
Solo Pro: 無制限のスキャンと修正が月額 19 ドル。
チーム: チームの展開、共有オンボーディング、無制限のスキャンと修正は月額 49 ドル/シートです。
PreFlight は、より深いローカル解析プリミティブを利用するようになりました。
Micro-Fuzzer: SQL インジェクション、コマンド インジェクション、認証バイパス、SSRF、パス トラバーサルなどの危険なデータ フロー パスに焦点を当てたセキュリティ ペイロードを生成します。
量子化 CPG (コード プロパティ グラフ): 構文、制御フロー、データ フローのコンパクトなメモリ内グラフを構築することで、PreFlight は脆弱な文字列マッチングに頼るのではなく、信頼できない入力を危険なシンクに追跡できます。
Eye デーモン: CLI/拡張ワークフローを通じてローカルで実行され、ファイルの保存を監視して、AI コーディング セッションがアクティブな間に問題が発生するようにします。
これはコアの PreFlight 信号です。すべてのスキャンは 3 つの明確な結果のいずれかに解決されるため、停止するか、レビューするか、出荷するかを判断できます。
PreFlight は、厳密な順序で修正を実行します。
フェーズ 1: オフライン ローカル AST スイープ
PreFlight は、まず超高速オフライン構造パスを完了し、決定論的な LO を適用します。

cal fix は安全に解決できます。
フェーズ 2: PreFlight Pro Deep Reasoning ハンドオフ
残りの SQL、ファザー、および複雑なアーキテクチャ上の欠陥は、パッチでより深いコンテキストが必要な場合に、安全なプロキシを利用した推論パスを通じて引き渡されます。
最初の 10 個のパッチ適用は、どちらのフェーズでも無料です。その後、PREFLIGHT_PRO_KEY が必要になります。
PreFlight は、VS Code/Cursor 拡張機能を介してターミナル内で直接実行することも、AI ネイティブ エディターの MCP サーバーとして実行することもできます。
プリフライトMCP
利用可能な MCP ツールには次のものがあります。
scan_project は引き続き無料で無制限です。 preflight_fix は、PREFLIGHT_PRO_KEY が必要になる前に、グローバルな 10 パッチの無料割り当てを共有します。
PreFlight は、ワンショット スキャナーではなく、閉ループとして使用されるように設計されています。
AI コーディング アシスタントを使用してコードを生成または変更します。
プリフライト スキャンを実行します。変化を Tri-State Risk Score に分類します。
PreFlight が Hard Block を返した場合は、先に進む前に停止して構造的な問題を修復してください。
PreFlight が High-Risk Drift を返した場合は、プリフライト スキャンを実行します。 --修正を適用し、提案されたすべての修正を適用する前に検査します。
プリフライト スキャンを再実行します。受け入れられた各修正の後、リポジトリが Pass に落ち着くことを確認します。
最終検証パスが緑色になり、構造受領書が意図したアーキテクチャ境界と一致した場合にのみ出荷されます。
この検証ループが成果物です。つまり、スキャン、レビュー、パッチ、再スキャン、そして自信を持った導入です。
AI 生成コード用のローカル セキュリティ ゲート。
github.com/av29nassh-sketch/PreFlight#readme
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
PreFlight スカベンジャー v0.1.0
最新の
2026 年 6 月 2 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション

エーション
私の個人情報を共有しないでください

## Original Extract

The local security gate for AI-generated code. Contribute to av29nassh-sketch/PreFlight development by creating an account on GitHub.

GitHub - av29nassh-sketch/PreFlight: The local security gate for AI-generated code. · GitHub
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
av29nassh-sketch
/
PreFlight
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
79 Commits 79 Commits .claude .claude .preflight-playground/ app/ api/ users .preflight-playground/ app/ api/ users .vscode .vscode assets assets beta-reasoning-fixture beta-reasoning-fixture docs docs hardcore-test-suite hardcore-test-suite playground-ast-only playground-ast-only playground playground preflight-dashboard preflight-dashboard preflight-proxy preflight-proxy qa/ preflight-gauntlet qa/ preflight-gauntlet real-flow-fixture real-flow-fixture scripts scripts server/ checkout server/ checkout src src supabase supabase test test tests tests wasm wasm .cursorrules .cursorrules .gitignore .gitignore .npmignore .npmignore .vercelignore .vercelignore LICENSE.md LICENSE.md PreFlight-Founder-Playbook.md PreFlight-Founder-Playbook.md PreFlight-Founder-Playbook.pdf PreFlight-Founder-Playbook.pdf README.md README.md cli.js cli.js demo.gif demo.gif e2e-license-live-test.js e2e-license-live-test.js generate-tests.js generate-tests.js index.js index.js logger.js logger.js mcp-instructions.md mcp-instructions.md package-lock.json package-lock.json package.json package.json remediationEngine.js remediationEngine.js reset-demo.js reset-demo.js scaffoldEngine.js scaffoldEngine.js taintTracker.js taintTracker.js View all files Repository files navigation
Stop AI Coding Drift before it becomes production technical debt. PreFlight is a local-first safety gate for AI-generated code, built to catch unsafe auth, RLS, SQL, SSRF, command execution, dependency, and secret-handling changes before they get committed.
Website: https://preflight-vibe.vercel.app
PreFlight runs in two distinct tiers depending on what your codebase needs.
What it does: Unlimited local scanning plus 10 free patch applications across local deterministic fixes and proxy-backed AI fixes.
Setup: Zero config for scanning. A Pro key is only required after the 10 free patches are used.
npm install -g preflight-pro
preflight init
preflight scan . --fix
Installing preflight-pro exposes the universal preflight command in your shell.
What it does: Unlimited scans and unlimited fixes, including deep reasoning remediation for complex multi-file architectural flaws, tenant isolation logic, and parametric SQL injections.
Setup: Requires an active PREFLIGHT_PRO_KEY or a saved key from preflight auth .
$ env: PREFLIGHT_PRO_KEY = " PREFLIGHT-BETA-XXXXX "
preflight scan . -- fix
Bash / macOS:
export PREFLIGHT_PRO_KEY= " PREFLIGHT-BETA-XXXXX "
preflight scan . --fix
Installation Flow
PreFlight supports both a terminal-first workflow and an IDE-first workflow. Both paths end with preflight init , because that wizard connects your editor, MCP clients, and Pro/Beta key in one place.
npm install -g preflight-pro
preflight init
Then scan any project from its root:
preflight scan . --fix
Path B: VS Code / Cursor
Install the global CLI command. The VSIX gives you the in-editor UI, but the extension still uses the global preflight command to start The Eye daemon and run fixes.
npm install -g preflight-pro
Download and install the PreFlight Companion VSIX extension:
Download VSIX from the PreFlight website
Or open GitHub Releases and install the latest preflight-companion VSIX.
In VS Code or Cursor, open the Extensions panel, click the ... menu, choose Install from VSIX... , and select the downloaded file.
preflight init
Open your project in the IDE. The extension starts The Eye automatically, watches file saves, and surfaces PreFlight alerts in-editor.
The Eye: The VS Code/Cursor extension starts PreFlight's local daemon automatically. It watches file saves and raises in-editor alerts when AI-generated code introduces a hard-block issue.
MCP bridge: preflight init can also wire preflight mcp into supported AI editors so agents can call PreFlight tools without leaving the coding flow.
Free users get unlimited scans and 10 total patches across local fixes and proxy-backed AI fixes. After the 10 free patches are used, unlimited fixes require a Pro/Beta key.
You can add your key during preflight init , or activate it directly:
preflight auth PREFLIGHT-BETA-XXXXX
For one terminal session, you can also set it manually:
$ env: PREFLIGHT_PRO_KEY = " PREFLIGHT-BETA-XXXXX "
export PREFLIGHT_PRO_KEY= " PREFLIGHT-BETA-XXXXX "
Pricing
Free Tier: Unlimited scans, 10 Free Patches (Local + Deep-Reasoning AI).
Solo Pro: $19/mo for unlimited scans and fixes.
Teams: $49/seat/mo for team rollout, shared onboarding, and unlimited scans and fixes.
PreFlight is now powered by deeper local analysis primitives:
Micro-Fuzzer: Generates focused security payloads for risky data-flow paths, such as SQL injection, command injection, auth bypass, SSRF, and path traversal.
Quantized CPG (Code Property Graph): Builds a compact in-memory graph of syntax, control flow, and data flow so PreFlight can trace untrusted input into dangerous sinks instead of relying on brittle string matching.
The Eye daemon: Runs locally through the CLI/extension workflow and watches file saves so issues appear while the AI coding session is still active.
This is the core PreFlight signal. Every scan resolves into one of three clear outcomes so you know whether to stop, review, or ship.
PreFlight runs fixes in a strict sequence:
Phase 1: Offline Local AST Sweep
PreFlight completes an ultra-fast offline structural pass first and applies any deterministic local fixes it can resolve safely.
Phase 2: PreFlight Pro Deep Reasoning Handoff
Remaining SQL, fuzzer, and complex architectural flaws are handed off through the secure proxy-backed reasoning path when a patch requires deeper context.
The first 10 patch applications are free across both phases. After that, a PREFLIGHT_PRO_KEY is required.
PreFlight can run directly in the terminal, through the VS Code/Cursor extension, or as an MCP server for AI-native editors.
preflight mcp
Available MCP tools include:
scan_project remains free and unlimited. preflight_fix shares the global 10-patch free allowance before a PREFLIGHT_PRO_KEY is required.
PreFlight is designed to be used as a closed loop, not a one-shot scanner:
Generate or modify code with your AI coding assistant.
Run preflight scan . to classify the change under the Tri-State Risk Score.
If PreFlight returns Hard Block , stop and repair the structural issue before moving forward.
If PreFlight returns High-Risk Drift , run preflight scan . --fix and inspect every proposed fix before applying it.
Re-run preflight scan . after each accepted fix to confirm the repository settles into Pass .
Ship only after the final verification pass is green and the structural receipt matches the architecture boundary you intended.
This verification loop is the product: scan, review, patch, re-scan, then deploy with confidence.
The local security gate for AI-generated code.
github.com/av29nassh-sketch/PreFlight#readme
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
PreFlight Scavenger v0.1.0
Latest
Jun 2, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
