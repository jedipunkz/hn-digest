---
source: "https://github.com/samchon/ttsc/tree/master/packages/graph"
hn_url: "https://news.ycombinator.com/item?id=48750084"
title: "Show HN: I Made TS Compiler Graph MCP: 10x Fewer Tokens in Claude Code and Codex"
article_title: "ttsc/packages/graph at master · samchon/ttsc · GitHub"
author: "autobe"
captured_at: "2026-07-01T17:35:55Z"
capture_tool: "hn-digest"
hn_id: 48750084
score: 1
comments: 0
posted_at: "2026-07-01T17:09:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I Made TS Compiler Graph MCP: 10x Fewer Tokens in Claude Code and Codex

- HN: [48750084](https://news.ycombinator.com/item?id=48750084)
- Source: [github.com](https://github.com/samchon/ttsc/tree/master/packages/graph)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T17:09:53Z

## Translation

タイトル: HN を表示: TS コンパイラ グラフ MCP を作成しました: クロード コードとコーデックスのトークンが 10 分の 1 に減少
記事のタイトル: ttsc/packages/graph at master · samchon/ttsc · GitHub
説明: コンパイラを利用したプラグインとタイプセーフな実行のための「typescript-go」ツールチェーン + コンパイラに統合された 500 倍高速な lint - マスターの ttsc/packages/graph · samchon/ttsc

記事本文:
マスターの ttsc/packages/graph · samchon/ttsc · GitHub
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
サムチョン
/
ttsc
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

s
コード
その他のオプション ディレクトリアクション
歴史 歴史マスター ブレッドクラム
コピーパスのトップフォルダーとファイル
.. build build src src README.md README.md package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 README.md
@ttsc/graph は、コーディング エージェントに、MCP 経由で TypeScript コードベースのグラフを提供します。つまり、何が何を呼び出し、何が何に依存し、各部分がどこに存在するかがわかります。
これは実際の TypeScript コンパイラによって描画されるため正確であり、すべてのクレームは開くことができるファイルと行に固定されています。
エージェントは、ファイルごとにクロールするのではなく、グラフから構造的な質問に答えます。これにより、オープンエンドの「これはどのように機能するのですか?」に関してトークンが約 10 倍削減されます。質問。
これを構築した理由、詳細な動作方法、 codegraph 、 codebase-memory-mcp 、および serena との比較については、開始時の投稿を読んでください: https://ttsc.dev/blog/i-made-ts-compiler-graph-mcp
npm install -D ttsc @ttsc/graph typescript@rc
@ttsc/graph は、ttsc 型チェックされたプログラムからグラフを読み取るため、2 つを一緒にインストールします。
ttsc は、まだリリース候補である TypeScript-Go (TypeScript v7) ランタイム上で実行されるため、インストールでは typescript@rc が固定されます。安定版 TypeScript v6.x ではまだ動作しません。
サーバーをエージェントの MCP 構成に一度追加します。クロード コードの場合、それはプロジェクト ルートの .mcp.json です。
{
"mcpサーバー": {
"ttsc-グラフ" : {
"コマンド" : " npx " ,
"args" : [ " -y " , " @ttsc/graph " ]
}
}
}
サーバーが tsconfig.json を見つけられるように、プロジェクト ルートからエージェントを起動します。エージェントは独自にグラフをクエリします。決して手で呼び出すことはありません。
この例では Claude Code としていますが、MCP 対応のエージェントであればどれでも機能します (Codex、Cursor など)。
MCP サーフェス全体は、inspect_typescript_graph という 1 つのツールです。平易な言葉で質問すると、ツール内で必要な短い一連の思考 (質問、ドラフト、レビュー) が SM の計画を立てます。

すべてのクエリを実行し、1 つの操作を選択します。
ソース本体ではなくインデックスです。名前、エッジ、署名、およびスパンを返しますが、コードを返すことはありません。そのため、リポジトリが大きくなっても応答はフラットのままです。
実際のコンパイラ上に構築されています。 ttsc 型チェックされたプログラムを読み取るため、tsconfig エイリアス、pnpm monorepos、シンボリックリンク、および再エクスポートは正確に解決されますが、テキスト パーサーは推測することしかできません。
それ自体を強制するものではありません。グラフが適切なソースであることを示し、それ以外のすべてに対して最高のエスケープを提供します。
エラーや糸くずも。 tsc コンパイル エラーと @ttsc/lint およびプラグイン (typia、nesia) の結果は同じグラフに乗っているため、「ここで何が壊れているのか?」 1つの指標から答えを導き出します。
操作 (ツアー、エントリポイント、ルックアップ、トレース、詳細、概要、エスケープ) と完全なリクエストと結果のコントラクトは、設計ガイドに記載されています: https://ttsc.dev/docs/graph/design
8 つの実際のリポジトリにわたって、@ttsc/graph はフラットで低いトークンコストの中央値を維持していますが、代替案はリポジトリのサイズに応じて変動します。
完全なベンチマークには、インタラクティブなチャート、すべてのモデル、メソッドが含まれています: https://ttsc.dev/docs/benchmark/graph
これを独自のプロジェクトで実行して、ローカル ポートから提供されるグラフをブラウザで開きます。
npx @ttsc/グラフビュー
これは、種類ごとに色分けされた 3D の TypeORM (ライブ ビューアー) です。
Launch post : これを構築した理由、および codegraph 、 codebase-memory-mcp 、および serena との比較。
デザイン : 1 つのツール、そのリクエストと結果のブランチ、およびノー​​ドとエッジの種類。
比較: 他のグラフおよび言語サーバー MCP ツールとの直接比較。
ベンチマーク : インタラクティブなチャート、すべてのモデル、メソッド。
あなたの寄付はttscの発展を奨励します。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A `typescript-go` toolchain for compiler-powered plugins and type-safe execution + 500x faster lint integrated into compiler - ttsc/packages/graph at master · samchon/ttsc

ttsc/packages/graph at master · samchon/ttsc · GitHub
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
samchon
/
ttsc
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History master Breadcrumbs
Copy path Top Folders and files
.. build build src src README.md README.md package.json package.json tsconfig.json tsconfig.json View all files README.md
@ttsc/graph gives your coding agent a graph of your TypeScript codebase over MCP: what calls what, what depends on what, where each piece lives.
It is drawn by the real TypeScript compiler, so it is exact, and every claim is anchored to a file and line you can open.
The agent answers structural questions from the graph instead of crawling file by file, which cuts its tokens by roughly 10x on open-ended "how does this work?" questions.
For why I built it, how it works in depth, and how it compares to codegraph , codebase-memory-mcp , and serena , read the launch post: https://ttsc.dev/blog/i-made-ts-compiler-graph-mcp
npm install -D ttsc @ttsc/graph typescript@rc
@ttsc/graph reads the graph from the program ttsc type-checked, so install the two together.
ttsc runs on the TypeScript-Go (TypeScript v7) runtime, which is still a release candidate, so the install pins typescript@rc . It does not run on stable TypeScript v6.x yet.
Add the server to your agent's MCP config, once. For Claude Code, that is a .mcp.json in your project root:
{
"mcpServers" : {
"ttsc-graph" : {
"command" : " npx " ,
"args" : [ " -y " , " @ttsc/graph " ]
}
}
}
Start your agent from your project root so the server finds your tsconfig.json . The agent queries the graph on its own; you never call it by hand.
The example says Claude Code, but any MCP-capable agent works (Codex, Cursor, and others).
The whole MCP surface is one tool, inspect_typescript_graph . You ask in plain language, and a short required chain of thought inside the tool ( question , draft , review ) plans the smallest query and picks one operation.
An index, not source bodies. It returns names, edges, signatures, and spans, never code, so the response stays flat as the repository grows.
Built on the real compiler. It reads the program ttsc type-checked, so tsconfig aliases, pnpm monorepos, symlinks, and re-exports resolve exactly, where a text parser can only guess.
It does not force itself. It states when the graph is the right source and offers a first-class escape for everything else.
Errors and lint too. tsc compile errors and @ttsc/lint and plugin (typia, nestia) findings ride the same graph, so "what is broken here?" answers from one index.
The operations ( tour , entrypoints , lookup , trace , details , overview , escape ) and the full request and result contract are in the Design guide: https://ttsc.dev/docs/graph/design
Across eight real repositories, @ttsc/graph holds a flat, low median token cost while the alternatives swing with repository size.
The full benchmark has the interactive charts, every model, and the method: https://ttsc.dev/docs/benchmark/graph
Run this in your own project to open the graph in your browser, served from a local port:
npx @ttsc/graph view
This is TypeORM in 3D, colored by kind ( live viewer ):
Launch post : why I built it, and how it compares to codegraph , codebase-memory-mcp , and serena .
Design : the one tool, its request and result branches, and the node and edge kinds.
Comparison : the head-to-head with other graph and language-server MCP tools.
Benchmark : the interactive charts, every model, and the method.
Your donation encourages ttsc development.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
