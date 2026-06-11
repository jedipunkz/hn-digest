---
source: "https://github.com/TheWinci/mimirs"
hn_url: "https://news.ycombinator.com/item?id=48491918"
title: "Show HN: Mimirs – persistent local memory for AI coding agents (MCP)"
article_title: "GitHub - TheWinci/mimirs: Local MCP server that gives AI coding agents persistent, searchable memory of your codebase · GitHub"
author: "winci"
captured_at: "2026-06-11T16:27:14Z"
capture_tool: "hn-digest"
hn_id: 48491918
score: 2
comments: 0
posted_at: "2026-06-11T15:42:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Mimirs – persistent local memory for AI coding agents (MCP)

- HN: [48491918](https://news.ycombinator.com/item?id=48491918)
- Source: [github.com](https://github.com/TheWinci/mimirs)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T15:42:48Z

## Translation

タイトル: 表示 HN: Mimirs – AI コーディング エージェント (MCP) の永続的なローカル メモリ
記事のタイトル: GitHub - TheWinci/mimirs: AI コーディング エージェントにコードベースの永続的で検索可能なメモリを提供するローカル MCP サーバー · GitHub
説明: AI コーディング エージェントにコードベースの永続的で検索可能なメモリを提供するローカル MCP サーバー - TheWinci/mimirs

記事本文:
GitHub - TheWinci/mimirs: AI コーディング エージェントにコードベースの永続的で検索可能なメモリを提供するローカル MCP サーバー · GitHub
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
ザウィンチ
/
ミミル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

n 個のオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
177 コミット 177 コミット .claude-plugin .claude-plugin ベンチマーク ベンチマーク ベンチマーク ベンチマーク ドキュメント ドキュメント フック フック スクリプト スクリプト スキル スキル src src テスト テスト wiki wiki .gitignore .gitignore BENCHMARKS.md BENCHMARKS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md bun.lock bun.lockデモ.gif Demon.gif デモ.テープ デモ.テープ mimirs-logo-2.png mimirs-logo-2.png package.json package.json tsconfig.json tsconfig.json wiki-screen.png wiki-screen.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
北欧の知恵と知識の神ミーミルにちなんで名付けられました。
AI コーディング エージェント用の永続的なプロジェクト メモリ。セットアップするコマンドは 1 つだけで、維持する必要はありません。
エージェントはすべてのセッションをブラインドで開始します。ファイル名を推測し、キーワードを grep し、無関係なファイルのコンテキストを書き込み、昨日話し合った内容をすべて忘れます。
ある実際のプロジェクトでは、典型的なプロンプトは 380K トークンを書き込み、エンドツーエンドで約 12 秒かかりました。
mimirs によるインデックス作成後: 91K トークン、約 3 秒 - コードベースで 76% 減少。数値はリポジトリのサイズ、クエリ、モデルによって異なります。
API キーはありません。雲はありません。ドッカーはありません。
bun と SQLite だけです。
セマンティック検索 ·
自動生成された Wiki
セッション間の記憶 ·
依存関係グラフ ·
注釈
対応機種: Claude Code、Cursor、Windsurf、JetBrains (Junie)、GitHub Copilot、任意の MCP クライアント
クイックスタート
Bun (curl -fsSL https://bun.sh/install | bash )、および macOS では最新の SQLite — Apple にバンドルされている SQLite は拡張機能をサポートしていません。
醸造インストールSQLite
Linux と Windows には、互換性のある SQLite がすでに付属しています。
2. エディターをセットアップします (自動)
bunx mimirs init --ide claude # または: カーソル、ウィンドサーフィン、副操縦士、ジェットブレイン、すべて
これにより、MCP サーバー構成が作成されます

、エディター ルール、 .mimirs/config.json 、および .gitignore エントリ。 --ide all を指定して実行すると、サポートされているすべてのエディターを一度にセットアップできます。
init は、Claude Code、Cursor、Windsurf、Copilot、および JetBrains (Junie) をカバーします。それ以外のすべて (Codex、Zed、カスタム クライアント) については、以下のスニペットの 1 つをコピーします。
3. エディターをセットアップする (マニュアル参照)
mimirs MCP サーバーは標準入出力上で実行されます。すべてのクライアントには同じ 3 つのものが必要です。コマンド ( bunx )、引数 ( ["mimirs@latest", "serve"] )、およびプロジェクト ルートを指す RAG_PROJECT_DIR 環境変数です。
{
"mcpサーバー": {
"ミミル" : {
"コマンド" : " bunx " ,
"args" : [ " mimirs@latest " , "serve " ],
"環境" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
カーソル — プロジェクトルートの .cursor/mcp.json
{
"mcpサーバー": {
"ミミル" : {
"コマンド" : " bunx " ,
"args" : [ " mimirs@latest " , "serve " ],
"環境" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
ウィンドサーフィン — ~/.codeium/windsurf/mcp_config.json (グローバル)
Windsurf は、プロジェクトではなくホーム ディレクトリから MCP サーバーを読み取ります。 JetBrains プラグインのバリアントは ~/.codeium/mcp_config.json を使用します。
{
"mcpサーバー": {
"ミミル" : {
"コマンド" : " bunx " ,
"args" : [ " mimirs@latest " , "serve " ],
"環境" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
JetBrains (Junie) — プロジェクトルートの .junie/mcp.json
{
"mcpサーバー": {
"ミミル" : {
"コマンド" : " bunx " ,
"args" : [ " mimirs@latest " , "serve " ],
"環境" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
GitHub Copilot — プロジェクト ルートの .vscode/mcp.json
VS Code の Copilot は、サーバー マップ ( mcpServers ではない) とタイプ フィールドを使用します。
{
"サーバー" : {
"ミミール" : {
"タイプ" : " stdio " ,
"コマンド" : " bunx " ,
"args" : [ " mimirs@latest " , "serve " ],
"環境" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project

」
}
}
}
}
コーデックス — ~/.codex/config.toml (グローバル)
Codex は JSON ではなく TOML を使用し、 ~/.codex/config.toml から読み取ります。プロジェクトごとに 1 つのブロック — 複数のリポジトリ ( mimirs-frontend 、 mimirs-api など) を接続する場合は、一意のテーブル名を選択します。
[ mcp_servers .ミミルズ]
コマンド = " bunx "
args = [ "mimirs@latest" , "serve " ]
env = { RAG_PROJECT_DIR = " /absolute/path/to/your/project " }
または、同様に、拡張された env テーブルを使用します。
[ mcp_servers .ミミルズ]
コマンド = " bunx "
args = [ "mimirs@latest" , "serve " ]
[ mcp_servers .ミミール。環境]
RAG_PROJECT_DIR = " /absolute/path/to/your/project "
読み取り専用のプロジェクト ディレクトリですか?インデックスをリダイレクトする
プロジェクトが読み取り専用マウントにある場合は、RAG_DB_DIR を書き込み可能な場所に設定します。インデックスは <project>/.mimirs/ の代わりにそこに存在します。
{
"mcpサーバー": {
"ミミール" : {
"コマンド" : " bunx " ,
"args" : [ " mimirs@latest " , "serve " ],
"環境" : {
"RAG_PROJECT_DIR" : " /read/only/project " ,
"RAG_DB_DIR" : " /home/me/.cache/mimirs/myproject "
}
}
}
}
4. 最初のインデックス
MCP サーバーは最初のクエリに対して遅延インデックスを作成するため、接続が完了したら、エージェントに何か質問するだけで済みます。完全なインデックスを事前に強制的に作成するには (大規模なリポジトリに役立ちます):
bunx mimirs インデックス # 現在のディレクトリ
bunx mimirs ステータス # ファイル、チャンク、埋め込みの数
5. デモを試してみる (オプション)
ブンクス・ミミールのデモ
手動ワークフロー ( init なし)
init は便利です。エディター (MCP 構成、エージェント ルール、 .gitignore 、 .mimirs/config.json ) を接続します。インデックスは構築されません。また、それ以下のものはインデックスを必要としません。インデックスとデフォルト設定は、初めてインデックスまたはクエリを実行するときに自動的に作成されます。
1. MCP サーバーを手動で追加します。上記のマニュアル参照からクライアントのスニペットをドロップします: command: "bunx" 、args: ["mimirs@latest", "serve"] 、および pr を指す RAG_PROJECT_DIR

オブジェクトルート。これが MCP セットアップ全体です。
init がないとエージェント ルール ファイルがないため、アシスタントはツールの存在を認識できません。プロンプトで mimirs について言及するか、CLAUDE.md からエディターのルールにツール リストをコピーします。
2. インデックスを作成します。 MCP サーバーは最初のツール呼び出しで遅延インデックスを作成するため、エージェントを介してこの手順をスキップできます。事前にインデックスを作成するには (大規模なリポジトリに推奨され、CLI 検索の前に必要です / 以下をお読みください):
bunx mimirs インデックス # 現在のディレクトリ
bunx mimirs Index /path/to/repo # 特定のディレクトリ
bunx mimirs Index --patterns " src/**/*.ts,*.md " # glob に制限する
bunx mimirs ステータス ファイル、チャンク、埋め込み数
init ファイルや構成ファイルは必要ありません。デフォルトが適用され、インデックスが <project>/.mimirs/ に書き込まれます。
3. CLI からクエリを実行します。 2 つの読み取りコマンド。どちらも現在のディレクトリのインデックスに対して実行されます (他の場所を指すには --dir を使用します)。
＃どこにありますか？ — ランク付けされたファイルパス + スニペットのプレビュー
bunx mimirs 検索「認証はどこで処理されるか」 -- トップ 10
＃それは何ですか？ — 実際に一致するコードチャンク (関数、クラス、セクション)
bunx mimirs read " jwt validation " --top 8 --threshold 0.3
--ext .ts,.tsx 、 --in src,packages/core 、または --exclude testing のいずれかを使用してスコープを指定します。注: CLI の検索/読み取りでは自動インデックスは作成されません。最初に mimirs インデックスを実行します (MCP サーバーのみがオンデマンドでインデックスを作成します)。
より深く統合するには、mimirs を Claude Code プラグインとしても利用できます。クロード コード セッション:
/plugin マーケットプレイスに https://github.com/TheWinci/mimirs.git を追加します
/plugin インストール mimirs
このプラグインは、MCP サーバー、3 つのフック (SessionStart (コンテキスト サマリー)、PostToolUse (編集時の自動再インデックス)、SessionEnd (自動チェックポイント))、および一般的なジョブのツールを調整する一連のワークフロー スキル (explorer、review、debug、catch-up、recall、d) を接続します。

oc-gaps 、および wiki 。
プラグインを使わずにスキルを習得したいですか?これらは、 skill/ の下にあるプレーンな SKILL.md ファイルです。プロジェクトの .claude/skills/<name>/ (リポジトリと共有) または ~/.claude/skills/<name>/ (すべてのプロジェクト) に好きなものをコピーすると、Claude Code が次のセッションでそれらを選択します。スキルはクロード コードの機能であるため、他のエディターには適用されませんが、MCP ツール自体はどこでも機能します。
89 ～ 97% のリコール @10、97 ～ 100% のリコール @20、MRR 0.69 ～ 0.77。 3 つの言語にわたる 4 つの実際のコードベースで、階層化された難易度混合のクエリ セット (それぞれ 72 ～ 120 のクエリ、約 1/3 の難易度) を使用してベンチマークを実施し、現在のパイプラインで 2026 年 6 月 4 日に再測定されました。完全な方法論は BENCHMARKS.md にあります。
より大きなリポジトリ (Kubernetes、Excalidraw) は十分な大きさがあるため、一部の正しいファイルはトップ 10 をわずかに超えてランク付けされます。リコールは上位 20 までに 97 ～ 100% に達するため、大規模なリポジトリでは searchTopK: 15 ～ 20 を設定します。
vs コーディングエージェント (ContextBench)
また、ContextBench (ゴールドコンテキスト) で mimirs も実行しました。
実際のリポジトリでの取得）、他のエントリは完全なコーディング エージェントです - マルチステップ
エクスプローラー — シングルコールツールではありません。焦点を絞ったクエリを指定すると (LLM が読み取り後に送信するもの)
この問題)、1 つの mimirs 検索呼び出しは、エージェントの軌跡全体に対して次のようにランク付けされます。
mimirs は、両方のカバレッジ メトリックを 1 回の呼び出しとしてリードします。ファイルの精度は最後です
目的: ゴールド ファイルの欠落は致命的です (LLM は修正すべきコードを決して認識しません)、余分なファイル
参照のフィルタリングは安価です。そのため、mimirs は再現率を最大化し、モデルに次の処理を実行させます。
精密パス。 n=15 サンプル対エージェントの 500 セット — 方向性。完全なリーダーボード、
注意事項、および BENCHMARKS.md のグラフ回復のストーリー。
ミミル
ツールなし (grep + Read)
コンテキストスタッフィング
クラウド RAG サービス
セットアップ
1つのコマンド
何もない
何もない
APIキー、アカウント
トークンコスト
～91K/プロンプト
~380K/プロンプト
コードベース全体
さまざま
検索品質

y
89 ～ 97% の再現率 @10
キーワードに応じて
N/A (すべてロード済み)
さまざま
コードの理解
AST 対応 (24 言語)
ラインレベル
なし
通常はラインレベル
セッション間の記憶
会話 + チェックポイント
なし
なし
一部
プライバシー
完全にローカル
ローカル
ローカル
データはマシンから出ます
価格
無料
無料
高額なトークン請求書
$10-50/月 + トークン
なぜ既存のツールを使わないのでしょうか?
Continue.dev の @codebase — 最も近い重複 (ローカル RAG、オープン ソース) ですが、取得はエディター拡張機能内に存在します。 Mimirs は、エージェントが計画できる明示的なツール ( search 、 read_relevant 、 project_map 、 search_conversation 、 annotate ) を備えたスタンドアロンの MCP サーバーに加え、会話のテーリングと Wiki ジェネレーターが組み込まれています。
Aider のリポジトリマップ — リポジトリの静的なツリーシッター概要、埋め込みなし。賢くて軽量ですが、概要は検索ではありません。mimirs はベクトル + BM25 を使用してクエリごとにチャンクをランク付けし、グラフの中心性によってブーストします。
Sourcegraph Cody / OpenCtx — コード検索には優れていますが、インデックス作成はクラウド インフラとアカウントに依存します。 Mimirs は 1 バンクス離れたところにあり、マシンから離れることはありません。
llama-index / LangChain / roll-your-own — これらはライブラリです。 Mimirs にはバッテリーが含まれており、AST 対応のチャンキング、ハイブリッド検索、ファイル ウォッチャー、会話テール、および既に接続されている注釈が含まれています。
解析とチャンク — を使用してコンテンツを分割します

[切り捨てられた]

## Original Extract

Local MCP server that gives AI coding agents persistent, searchable memory of your codebase - TheWinci/mimirs

GitHub - TheWinci/mimirs: Local MCP server that gives AI coding agents persistent, searchable memory of your codebase · GitHub
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
TheWinci
/
mimirs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
177 Commits 177 Commits .claude-plugin .claude-plugin benchmark benchmark benchmarks benchmarks docs docs hooks hooks scripts scripts skills skills src src tests tests wiki wiki .gitignore .gitignore BENCHMARKS.md BENCHMARKS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md bun.lock bun.lock demo.gif demo.gif demo.tape demo.tape mimirs-logo-2.png mimirs-logo-2.png package.json package.json tsconfig.json tsconfig.json wiki-screen.png wiki-screen.png View all files Repository files navigation
Named after Mímir , the Norse god of wisdom and knowledge.
Persistent project memory for AI coding agents. One command to set up, nothing to maintain.
Your agent starts every session blind — guessing filenames, grepping for keywords, burning context on irrelevant files, and forgetting everything you discussed yesterday.
On one real project, a typical prompt was burning 380K tokens and ~12 seconds end-to-end .
After indexing with mimirs: 91K tokens, ~3 seconds — a 76% drop on that codebase. Your numbers will vary with repo size, query, and model.
No API keys. No cloud. No Docker.
Just bun and SQLite.
Semantic Search ·
Auto-generated Wiki
Cross-session Memory ·
Dependency Graphs ·
Annotations
Works with: Claude Code · Cursor · Windsurf · JetBrains (Junie) · GitHub Copilot · any MCP client
Quick start
Bun ( curl -fsSL https://bun.sh/install | bash ) and, on macOS, a modern SQLite — Apple's bundled one doesn't support extensions:
brew install sqlite
Linux and Windows ship with a compatible SQLite already.
2. Set up your editor (automatic)
bunx mimirs init --ide claude # or: cursor, windsurf, copilot, jetbrains, all
This creates the MCP server config, editor rules, .mimirs/config.json , and .gitignore entry. Run with --ide all to set up every supported editor at once.
init covers Claude Code, Cursor, Windsurf, Copilot, and JetBrains (Junie). For everything else — Codex, Zed, custom clients — copy one of the snippets below.
3. Set up your editor (manual reference)
The mimirs MCP server runs over stdio. Every client needs the same three things: a command ( bunx ), args ( ["mimirs@latest", "serve"] ), and a RAG_PROJECT_DIR env var pointing at your project root.
{
"mcpServers" : {
"mimirs" : {
"command" : " bunx " ,
"args" : [ " mimirs@latest " , " serve " ],
"env" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
Cursor — .cursor/mcp.json in project root
{
"mcpServers" : {
"mimirs" : {
"command" : " bunx " ,
"args" : [ " mimirs@latest " , " serve " ],
"env" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
Windsurf — ~/.codeium/windsurf/mcp_config.json (global)
Windsurf reads MCP servers from your home directory, not the project. JetBrains plugin variant uses ~/.codeium/mcp_config.json .
{
"mcpServers" : {
"mimirs" : {
"command" : " bunx " ,
"args" : [ " mimirs@latest " , " serve " ],
"env" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
JetBrains (Junie) — .junie/mcp.json in project root
{
"mcpServers" : {
"mimirs" : {
"command" : " bunx " ,
"args" : [ " mimirs@latest " , " serve " ],
"env" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
GitHub Copilot — .vscode/mcp.json in project root
VS Code's Copilot uses a servers map (not mcpServers ) and a type field.
{
"servers" : {
"mimirs" : {
"type" : " stdio " ,
"command" : " bunx " ,
"args" : [ " mimirs@latest " , " serve " ],
"env" : {
"RAG_PROJECT_DIR" : " /absolute/path/to/your/project "
}
}
}
}
Codex — ~/.codex/config.toml (global)
Codex uses TOML, not JSON, and reads from ~/.codex/config.toml . One block per project — pick a unique table name if you wire up multiple repos ( mimirs-frontend , mimirs-api , etc).
[ mcp_servers . mimirs ]
command = " bunx "
args = [ " mimirs@latest " , " serve " ]
env = { RAG_PROJECT_DIR = " /absolute/path/to/your/project " }
Or, equivalently, with an expanded env table:
[ mcp_servers . mimirs ]
command = " bunx "
args = [ " mimirs@latest " , " serve " ]
[ mcp_servers . mimirs . env ]
RAG_PROJECT_DIR = " /absolute/path/to/your/project "
Read-only project directory? Redirect the index
If the project lives in a read-only mount, set RAG_DB_DIR to a writable location. The index lives there instead of <project>/.mimirs/ .
{
"mcpServers" : {
"mimirs" : {
"command" : " bunx " ,
"args" : [ " mimirs@latest " , " serve " ],
"env" : {
"RAG_PROJECT_DIR" : " /read/only/project " ,
"RAG_DB_DIR" : " /home/me/.cache/mimirs/myproject "
}
}
}
}
4. First index
The MCP server indexes lazily on the first query, so once it's wired up you can just ask your agent something. To force a full index up front (useful for large repos):
bunx mimirs index # current directory
bunx mimirs status # how many files, chunks, embeddings
5. Try the demo (optional)
bunx mimirs demo
Manual workflow (without init )
init is a convenience: it wires up your editor (MCP config, agent rules, .gitignore , .mimirs/config.json ). It does not build the index, and nothing below needs it — the index and a default config are created automatically the first time you index or query.
1. Add the MCP server by hand. Drop the snippet for your client from the manual reference above: command: "bunx" , args: ["mimirs@latest", "serve"] , and RAG_PROJECT_DIR pointing at your project root. That is the entire MCP setup.
Without init there's no agent-rules file, so your assistant won't know the tools exist. Either mention mimirs in your prompt, or copy the tool list from CLAUDE.md into your editor's rules.
2. Build the index. The MCP server indexes lazily on the first tool call, so through an agent you can skip this step. To index up front (recommended for large repos, and required before the CLI search / read below):
bunx mimirs index # current directory
bunx mimirs index /path/to/repo # a specific directory
bunx mimirs index --patterns " src/**/*.ts,*.md " # restrict to globs
bunx mimirs status # files, chunks, embeddings
No init and no config file required — defaults are applied and the index is written to <project>/.mimirs/ .
3. Query from the CLI. Two read commands, both running against the index in the current directory (use --dir to point elsewhere):
# Where is it? — ranked file paths + snippet previews
bunx mimirs search " where is auth handled " --top 10
# What is it? — the actual matching code chunks (functions, classes, sections)
bunx mimirs read " jwt validation " --top 8 --threshold 0.3
Scope either with --ext .ts,.tsx , --in src,packages/core , or --exclude tests . Note: the CLI search / read do not auto-index — run mimirs index first (only the MCP server indexes on demand).
For deeper integration, mimirs is also available as a Claude Code plugin. In a Claude Code session:
/plugin marketplace add https://github.com/TheWinci/mimirs.git
/plugin install mimirs
The plugin wires the MCP server, three hooks — SessionStart (context summary), PostToolUse (auto-reindex on edit), SessionEnd (auto-checkpoint) — and a set of workflow skills that orchestrate the tools for common jobs: explore , review , debug , catch-up , recall , doc-gaps , and wiki .
Want the skills without the plugin? They're plain SKILL.md files under skills/ . Copy any you like into your project's .claude/skills/<name>/ (shared with the repo) or ~/.claude/skills/<name>/ (all your projects) and Claude Code picks them up next session. Skills are a Claude Code feature, so they don't apply to other editors — but the MCP tools themselves work everywhere.
89–97% Recall@10, 97–100% Recall@20, MRR 0.69–0.77. Benchmarked on four real codebases across three languages with stratified, difficulty-mixed query sets (72–120 queries each, ~⅓ hard), re-measured 2026-06-04 on the current pipeline. Full methodology in BENCHMARKS.md .
The larger repos (Kubernetes, Excalidraw) are big enough that some correct files rank just past the top-10; recall reaches 97–100% by top-20, so set searchTopK: 15–20 on large repos.
vs coding agents (ContextBench)
We also ran mimirs on ContextBench (gold-context
retrieval on real repos), whose other entries are full coding agents — multi-step
explorers — not single-call tools. Given a focused query (what an LLM sends after reading
the issue), one mimirs retrieval call ranks like this against whole agent trajectories:
mimirs leads both coverage metrics as a single call. File precision is last on
purpose: a missed gold file is fatal (the LLM never sees the code to fix), an extra file
reference is cheap to filter — so mimirs maximizes recall and lets the model do the
precision pass. n=15 sample vs the agents' 500-set — directional. Full leaderboards,
caveats, and the graph-recovery story in BENCHMARKS.md .
mimirs
No tool (grep + Read)
Context stuffing
Cloud RAG services
Setup
One command
Nothing
Nothing
API keys, accounts
Token cost
~91K/prompt
~380K/prompt
Entire codebase
Varies
Search quality
89–97% Recall@10
Depends on keywords
N/A (everything loaded)
Varies
Code understanding
AST-aware (24 langs)
Line-level
None
Usually line-level
Cross-session memory
Conversations + checkpoints
None
None
Some
Privacy
Fully local
Local
Local
Data leaves your machine
Price
Free
Free
High token bills
$10-50/mo + tokens
Why not an existing tool?
Continue.dev's @codebase — closest overlap (local RAG, open source), but retrieval lives inside the editor extension. Mimirs is a standalone MCP server with explicit tools ( search , read_relevant , project_map , search_conversation , annotate ) the agent can plan around, plus conversation tailing and a wiki generator built in.
Aider's repo-map — static tree-sitter summary of the repo, no embeddings. Clever and lightweight, but a summary isn't retrieval — mimirs ranks chunks per query with vector + BM25 and boosts by graph centrality.
Sourcegraph Cody / OpenCtx — excellent at code search, but indexing leans on cloud infra and an account. Mimirs is one bunx away and never leaves your machine.
llama-index / LangChain / roll-your-own — those are libraries. Mimirs is batteries-included: AST-aware chunking, hybrid retrieval, file watcher, conversation tail, and annotations already wired together.
Parse & chunk — Splits content usin

[truncated]
