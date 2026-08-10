---
source: "https://github.com/thecolourfoundation/rune"
hn_url: "https://news.ycombinator.com/item?id=49247037"
title: "Hi HN Rune – A software intelligence runtime for AI coding assistants"
article_title: "GitHub - thecolourfoundation/rune · GitHub"
author: "malixp"
captured_at: "2026-08-10T17:43:37Z"
capture_tool: "hn-digest"
hn_id: 49247037
score: 1
comments: 0
posted_at: "2026-08-10T17:39:22Z"
tags:
  - hacker-news
  - translated
---

# Hi HN Rune – A software intelligence runtime for AI coding assistants

- HN: [49247037](https://news.ycombinator.com/item?id=49247037)
- Source: [github.com](https://github.com/thecolourfoundation/rune)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T17:39:22Z

## Translation

タイトル: Hi HN Rune – AI コーディング アシスタント用のソフトウェア インテリジェンス ランタイム
記事タイトル: GitHub - thecolourfoundation/rune · GitHub
説明: GitHub でアカウントを作成して、colourfoundation/rune の開発に貢献します。

記事本文:
GitHub - thecolourfoundation/rune · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
カラーファンデーション
/
ルーン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット bin bin python python scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ソフトウェア インテリジェンス ランタイム。
カーソル、コーデックス、

Claude Code — 彼らはすべて、セッションごとにコードベースをコールドで読み取り、セッションが終了するとすべてを忘れます。 Rune はそれらの下で継続的に実行されるものであるため、すべてのセッションはコードベースを再派生するのではなく、すでに認識し始めます。
フローチャート LR
A1[「新しい AI セッション」] --> A2[「ファイルをコールドで読み取る」]
A2 --> A3[「理由と答え」]
A3 --> A4["セッション終了"]
A4 -.->|「理解が失われています」| A1
読み込み中
ルーンの場合:
フローチャート LR
B1[「コード変更」] --> B2[「ルーンウォッチは自動的に再構築される」]
B2 --> B3["常に最新の理解"]
B3 --> B4[「新しい AI セッション」]
B4 --> B5["MCP 上のクエリ ルーン"]
B5 --> B6[「証拠に裏付けられた回答、ファイル + 行に引用」]
B6 -.->|"最新のまま"| B3
読み込み中
機能 · 仕組み · インストール · クイックスタート · 実際の出力の確認 · CLI · MCP ツール · 制限事項 · セキュリティ
常に最新 — ルーン ウォッチは、誰かがスキャンを再実行したことを覚えているときだけでなく、編集中でも理解グラフを最新の状態に保ちます。これはランタイムであり、ワンショット CLI ではありません。
永続的 — 理解は 1 回のセッションを過ぎても持続します。接続するすべての AI は、ソフトウェアをコールドで再度読み取るのではなく、すでにソフトウェアを認識し始めます。
共有 — 1 つの理解、多数の AI クライアント。クロード、GPT、ジェミニなど、来月実行するものはすべて、コードの一貫性のない個別のメンタル モデルを維持するのではなく、すべて同じ Rune インスタンスを要求します。
説明可能 — Rune が AI に伝えるものは何もブラックボックスです。主張を正当化するように要求すると、その背後にある正確な情報源の場所が表示されます。
読み取り専用、提案専用 — Rune はソースに書き込むことはありません。それが実際の堀です。この領域の他のすべてのものは自律的なコード作成に向けて競争しています。ルーンは、厳密にその線を読み、理解し、そして（後に）示唆する側に留まります。黙って再確認する必要は決してありません

コードを王様にしましょう。
2 つの層があり、すべての結論は証拠に遡ります。主張されず、常に検査可能です。
フローチャート LR
F1["事実: インポート反応、UserCard.jsx 行 1"] --> D["派生: React コンポーネント UserCard"]
F2["事実: 関数 UserCard、UserCard.jsx 3 行目"] --> D
F3["事実: useState フック、UserCard.jsx 4 行目"] --> D
D --> C["AI の答え: このプロジェクトには UserCard コンポーネントがあります"]
C -.->|"rune_explain"| F1
C -.->|"rune_explain"| F2
C -.->|"rune_explain"| F3
読み込み中
何かに対して rune Explain <id> (または rune_explain MCP ツール) を呼び出して、そのチェーンを取り戻します。 Rune を使用する AI は、アーキテクチャを推測しません。Rune も同様です。
npm インストール @pypl100/rune
Node.js を必要としないチーム向けに、完全な機能同等 (同じファクト/派生モデル、同じ検出器、同じ MCP ツール サーフェス) を備えた Python ディストリビューションが python/ に存在します。 pip install North-rune として公開されています (CLI コマンドは依然として rune です。 python/README.md を参照)。
両方 (@pypl100/rune (npm 経由でグローバルに、north-rune 経由で pip 経由)) を同じマシンにインストールした場合、実際に PATH 上に存在する rune コマンドは 1 つだけです。最後にインストールしたものではなく、シェルが最初に見つけたものになります。どの -a rune を使用して確認します。複数のパスがリストされている場合は、それが理由です。ここにはバージョン検出の魔法はありません。これは、同じ名前のバイナリを偶然インストールする無関係な 2 つのツールと同じ、単純な OS PATH 解決です。特定のものが必要な場合は、裸の rune に依存するのではなく、フルパスで呼び出してください。
プロジェクトの cd
npm install -g @pypl100/rune # または: npx -p @pypl100/rune rune <コマンド>
rune init # プロジェクトに Rune を設定します
rune watch & # 作業中にバックグラウンドで最新の理解を維持します
runeserve # MCP サーバーを起動し、任意の AI クライアントに公開します
ルーンウォッチは推奨されるデフォルトです - それが何ですか

Rune は、再実行を覚えておく必要があるツールの代わりにランタイムになります。 (CI またはワンショット チェックの場合は、代わりにルーン スキャンを使用してください。)
シーケンス図
あなたとしての参加者開発者
参加者 ルーンウォッチとして見る
参加者グラフをグラフファイルとして表示
AI クライアントとしての参加者 AI (クロード コード、カーソルなど)
Dev->>Watch: ファイルを保存する
Watch->>Watch: 変更を検出 (デバウンス)
ウォッチ - >> グラフ: リビルド + 書き込み
AI->>グラフ: rune_search / rune_explain (ルーンサーブ経由)
グラフ-->>AI: 現在の証拠に裏付けられた答え
読み込み中
MCP 互換クライアント (Claude Desktop、Claude Code、カスタム エージェントなど) をルーン サーブ プロセスに向けます。接続されているすべての AI は、同じ継続的な最新の理解を共有します。再起動や手動の再スキャンは必要ありません。
MCP クライアント設定エントリの例:
{
"mcpサーバー": {
「ルーン」: {
"コマンド" : " npx " ,
"args" : [ " -p " 、 " @pypl100/rune " 、 " rune " 、 "serve " 、 " /absolute/path/to/your-project " ]
}
}
}
これは実際にはどのように見えるか
このリポジトリのフィクスチャ プロジェクトからの実際の出力 - ステージングされていません:
$ルーンスキャン。
[ルーン] 34 ミリ秒で 4 つのファイルをスキャンしました
[ルーン]事実: 10、導き出された結論: 4
[ルーン] /project/.rune/graph.json に書き込まれるグラフ
$ ルーン説明コンポーネント_2
{
"種類": "機能",
"id": "コンポーネント_2",
"タイプ": "反応コンポーネント",
"ファイル": "コンポーネント/UserCard.jsx",
「行」: 3、
"名前": "ユーザーカード",
"証拠": "エクスポート関数 UserCard({ ユーザー }) {"
}
これが、一例における信頼モデル全体です。Rune は単に「UserCard コンポーネントがある」と言うだけではなく、正確なファイル、正確な行、および一致した正確なテキストを指します。何か説明を求めれば、請求ではなく領収書を受け取ることになります。
コマンド
何をするのか
ルーン初期化 [ディレクトリ]
現在の (または指定された) プロジェクトに .rune/ を設定します
ルーンスキャン[ディレクトリ]
理解グラフを一度構築（または再構築）します。
ルーンウォッチ[ディレクトリ]
理解グラフを最新の状態に保ちます

ファイルの変更に応じて (停止するには Ctrl+C)
ルーンサーブ[ディレクトリ]
MCPサーバーを起動します
ルーン説明 <id>
事実や結論の背後にある証拠の痕跡を印刷します
ルーン --バージョン
インストールされている Rune のバージョンを出力します
構成
rune init は .rune/config.json を書き込みます。現在の唯一の設定は無視です。組み込みのデフォルト (node_modules 、 .git 、 dist 、 build 、 .next 、coverage など、およびすべての dotfiles 無条件) に加えて、スキャンから除外するディレクトリ/ファイル名の配列です。
{
"無視" : [ " レガシー " 、 " ベンダー " ]、
「バージョン」：1
}
MCP ツールが公開される
ツール
目的
ルーン取得_概要
アーキテクチャの概要 — ここから始めてください
ルーンリストコンポーネント
検出されたすべての React コンポーネント
ルーンリストルート
Unified Express + Next.js ルート リスト
ルーン検索
名前、ファイル、またはルートの部分文字列によってファクト/派生ノードを検索します
ルーン説明
あらゆるIDの完全な証拠追跡
rune_get_file_dependency
ファイルの内部インポートグラフ
ルーン再スキャン
コード変更後のオンデマンドでの再スキャン
現在の範囲と正直な限界
Rune は現在意図的に狭いです:
フレームワークのサポート: React、Next.js (ページ + アプリ ルーター)、Express。それ以外の場合は、汎用ファイル/インポート スキャンのみが行われます。
抽出方法: ヒューリスティックな正規表現ベースのパターン マッチング — 完全な AST パーサーではありません。これにより、スキャナーは依存関係がなく高速に保たれ、すべてのファクトはファイル/行/一致テキストの証拠を保持しますが、異常なコード形状 (例: JSX を使用せずに React.createElement 経由で返されたコンポーネント、動的に構築されたルート文字列、深く再エクスポートされたコンポーネント) を見逃します。本物の AST ベースの抽出プログラムは当然の次のアップグレードです。実際、スキーマはダウンストリームに何も触れずに抽出方法を交換できるように設計されています。
ファイル間のルートプレフィックス解決はありません。Express ルートは app.use('/api', router) を介して 1 つのファイルにマウントされ、別のファイルで定義されます。

まだ単一のパスに縫い込まれていません。
読み取り専用: Rune はソース ファイルに書き込むことはありません。独自のグラフを .rune/graph.json に書き込むだけです。
単一プロセス、stdio MCP トランスポート - マルチクライアント デーモンはまだありません。
Rune は、例外なく、ドットファイルまたはドット ディレクトリ ( .env 、 .env.*.js 、 .git 、 .ssh 、エディタ設定など) をスキャンしません。これはスキャナーで強制され、回帰テストによってカバーされます。一致するコード拡張子 (例: .env.js ) を持つ構成ファイルは、その内容が証拠として読み取られたり埋め込まれたりすることはありません。
シンボリックリンクはトラバースされないため、プロジェクト ルートの外側を指すシンボリックリンクは外部ファイルをスキャンに取り込むことができません。
.rune/graph.json は rune init によって git から除外されます (まだ存在しない場合は .gitignore が作成されます)。
Rune が書き込むのはグラフ自体だけです。 .rune/graph.json を AI クライアントと共有すると、その内容はすべて共有されることになります。ソースのスニペットを引用する他のファイルと同様に扱います。
単体テスト スイート ( npm test ) はスキャンとグラフの理解に対応しますが、実際の MCP クライアントは起動しません。 scripts/verify-mcp-server.mjs は、実際の子プロセスとして rune を生成し、実際の JSON-RPC ハンドシェイク (initialize → notification/initialized → tools/list → tools/call ) を介して駆動し、予期されるすべてのツールが公開されていることを確認し、真に無効な呼び出しがサーバーをクラッシュさせるのではなく完全に拒否されていることを確認します。
npmインストール
npm run verify:mcp
# または、バンドルされたフィクスチャの代わりに実際のプロジェクトに対して:
npm run verify:mcp -- /path/to/your-project
ロードマップ
AST ベースの抽出 (正規表現スキャナーのスワップイン置換)
フロントエンド呼び出しとバックエンドルート間のデータフロートレース
真の増分再スキャン — ルーン ウォッチは現在も存在しますが、変更のたびに完全な再構築が行われます。どの ch の差分ではありません。

怒った。小規模から中規模のプロジェクトには適しています。非常に大きなものでは重要になります。
提案ツール (例: 既知の脆弱性のある依存関係パターンに証拠を付けてフラグを立てる、レビュー可能な修正を提案する) - 読み取り専用モデルに厳密に付加的であり、自動適用されることはありません
サーカスをサポートしたい場合: ETH 0xbc0979dde621c353737d21f6d7b4eb361f7bc11f
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to thecolourfoundation/rune development by creating an account on GitHub.

GitHub - thecolourfoundation/rune · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
thecolourfoundation
/
rune
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits bin bin python python scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
The Software Intelligence Runtime.
Cursor, Codex, Claude Code — they all read your codebase cold, every session, and forget everything when it ends. Rune is the thing that runs underneath them, continuously, so every session starts already knowing your codebase instead of re-deriving it.
flowchart LR
A1["New AI session"] --> A2["Reads files cold"]
A2 --> A3["Reasons, answers"]
A3 --> A4["Session ends"]
A4 -.->|"understanding lost"| A1
Loading
With Rune:
flowchart LR
B1["Code changes"] --> B2["rune watch rebuilds automatically"]
B2 --> B3["Understanding always current"]
B3 --> B4["New AI session"]
B4 --> B5["Queries Rune over MCP"]
B5 --> B6["Evidence-backed answer, cited to file + line"]
B6 -.->|"stays current"| B3
Loading
What it does · How it works · Install · Quickstart · See real output · CLI · MCP tools · Limitations · Security
Always current — rune watch keeps the understanding graph up to date as you edit, not just when someone remembers to re-run a scan. It's a runtime, not a one-shot CLI.
Persistent — understanding survives past any single session. Every AI you connect starts already knowing your software, instead of re-reading it cold.
Shared — one understanding, many AI clients. Claude, GPT, Gemini, whatever you're running next month — they all ask the same Rune instance instead of maintaining separate, inconsistent mental models of your code.
Explainable — nothing Rune tells an AI is a black box. Ask it to justify any claim and it shows the exact source location behind it.
Read-only, suggest-only — Rune never writes to your source. That's the actual moat: everything else in this space is racing toward autonomous code-writing; Rune stays strictly on the reading, understanding, and (later) suggesting side of that line. It never becomes the thing you have to double-check for silently breaking your code.
Two layers, and every conclusion traces back to evidence — not asserted, always inspectable:
flowchart LR
F1["Fact: import react, UserCard.jsx line 1"] --> D["Derived: React component UserCard"]
F2["Fact: function UserCard, UserCard.jsx line 3"] --> D
F3["Fact: useState hook, UserCard.jsx line 4"] --> D
D --> C["AI answer: this project has a UserCard component"]
C -.->|"rune_explain"| F1
C -.->|"rune_explain"| F2
C -.->|"rune_explain"| F3
Loading
Call rune explain <id> (or the rune_explain MCP tool) on anything and get that chain back. An AI using Rune isn't guessing about your architecture — and neither is Rune.
npm install @pypl100/rune
A Python distribution exists at python/ with full feature parity (same fact/derived model, same detectors, same MCP tool surface) for teams who'd rather not require Node.js — published as pip install north-rune (the CLI command is still rune ; see python/README.md ).
If you install both ( @pypl100/rune via npm globally, and north-rune via pip) on the same machine, only one rune command will actually be on your PATH — whichever your shell finds first, not whichever you installed most recently. Check with which -a rune ; if it lists more than one path, that's why. There's no version-detection magic here — it's plain OS PATH resolution, same as any two unrelated tools that happen to install a same-named binary. If you need a specific one, invoke it by its full path rather than relying on bare rune .
cd your-project
npm install -g @pypl100/rune # or: npx -p @pypl100/rune rune <command>
rune init # sets up Rune in your project
rune watch & # keeps the understanding current in the background as you work
rune serve # starts an MCP server exposing it to any AI client
rune watch is the recommended default — it's what makes Rune a runtime instead of a tool you have to remember to re-run. (For CI or a one-shot check, use rune scan instead.)
sequenceDiagram
participant Dev as You
participant Watch as rune watch
participant Graph as graph file
participant AI as AI client (Claude Code, Cursor, etc.)
Dev->>Watch: save a file
Watch->>Watch: detect change (debounced)
Watch->>Graph: rebuild + write
AI->>Graph: rune_search / rune_explain (via rune serve)
Graph-->>AI: current, evidence-backed answer
Loading
Point any MCP-compatible client (Claude Desktop, Claude Code, custom agents, etc.) at the rune serve process. Every connected AI shares the same, continuously current understanding — no restart, no manual rescan.
Example MCP client config entry:
{
"mcpServers" : {
"rune" : {
"command" : " npx " ,
"args" : [ " -p " , " @pypl100/rune " , " rune " , " serve " , " /absolute/path/to/your-project " ]
}
}
}
What this actually looks like
Real output, from the fixture project in this repo — not staged:
$ rune scan .
[rune] scanned 4 file(s) in 34ms
[rune] facts: 10, derived conclusions: 4
[rune] graph written to /project/.rune/graph.json
$ rune explain component_2
{
"kind": "function",
"id": "component_2",
"type": "react_component",
"file": "components/UserCard.jsx",
"line": 3,
"name": "UserCard",
"evidence": "export function UserCard({ user }) {"
}
That's the whole trust model in one example: Rune doesn't just say "there's a UserCard component" — it points at the exact file, the exact line, and the exact text it matched. Ask it to explain anything, and you get the receipts, not a claim.
Command
What it does
rune init [dir]
Sets up .rune/ in the current (or given) project
rune scan [dir]
Builds (or rebuilds) the understanding graph, once
rune watch [dir]
Keeps the understanding graph current as files change (Ctrl+C to stop)
rune serve [dir]
Starts the MCP server
rune explain <id>
Prints the evidence trail behind any fact or conclusion
rune --version
Prints the installed Rune version
Configuration
rune init writes .rune/config.json . The only setting today is ignore — an array of directory/file names to exclude from scanning, on top of the built-in defaults ( node_modules , .git , dist , build , .next , coverage , etc., and all dotfiles unconditionally):
{
"ignore" : [ " legacy " , " vendor " ],
"version" : 1
}
MCP tools exposed
Tool
Purpose
rune_get_overview
Architecture summary — start here
rune_list_components
All detected React components
rune_list_routes
Unified Express + Next.js route list
rune_search
Find facts/derived nodes by name, file, or route substring
rune_explain
Full evidence trail for any id
rune_get_file_dependencies
Internal import graph for a file
rune_rescan
Re-scan on demand after code changes
Current scope and honest limitations
Rune is intentionally narrow right now:
Framework support: React, Next.js (pages + app router), Express. Everything else gets generic file/import scanning only.
Extraction method: heuristic, regex-based pattern matching — not a full AST parser. This keeps the scanner dependency-free and fast, and every fact still carries file/line/matched-text evidence, but it will miss unusual code shapes (e.g. components returned via React.createElement with no JSX, dynamically constructed route strings, deeply re-exported components). A real AST-based extractor is the natural next upgrade — the fact schema is designed so extraction method can be swapped without touching anything downstream.
No cross-file route-prefix resolution: an Express route mounted via app.use('/api', router) in one file and defined in another isn't stitched into a single path yet.
Read-only: Rune never writes to your source files. It only ever writes its own graph to .rune/graph.json .
Single-process, stdio MCP transport — no multi-client daemon yet.
Rune never scans dotfiles or dot-directories ( .env , .env.*.js , .git , .ssh , editor configs, etc.), with no exceptions. This is enforced in the scanner and covered by a regression test — a config file with a matching code extension (e.g. .env.js ) will not have its contents read or embedded as evidence.
Symlinks are not traversed, so a symlink pointing outside the project root can't pull external files into the scan.
.rune/graph.json is excluded from git by rune init (it creates .gitignore if one doesn't already exist).
The graph itself is the only thing Rune writes. If you share .rune/graph.json with an AI client, you're sharing everything in it — treat it like any other file that quotes snippets of your source.
The unit test suite ( npm test ) covers scanning and the understanding graph, but doesn't spin up a real MCP client. scripts/verify-mcp-server.mjs does: it spawns rune serve as a real child process and drives it through the actual JSON-RPC handshake ( initialize → notifications/initialized → tools/list → tools/call ), checks that all expected tools are exposed, and confirms a genuinely invalid call is rejected cleanly rather than crashing the server.
npm install
npm run verify:mcp
# or against a real project instead of the bundled fixture:
npm run verify:mcp -- /path/to/your-project
Roadmap
AST-based extraction (swap-in replacement for the regex scanner)
Data-flow tracing between frontend calls and backend routes
True incremental re-scan — rune watch exists now, but it still does a full rebuild on every change, not a diff of just what changed. Fine for small-to-medium projects; will matter on very large ones.
Suggestion tools (e.g. flagging a known-vulnerable dependency pattern with evidence, proposing a reviewable fix) — strictly additive to the read-only model, never auto-applied
If you want to support the circus: ETH 0xbc0979dde621c353737d21f6d7b4eb361f7bc11f
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
