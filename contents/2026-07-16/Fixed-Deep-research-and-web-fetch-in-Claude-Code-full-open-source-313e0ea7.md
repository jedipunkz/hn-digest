---
source: "https://github.com/kunal12203/webify-mcp"
hn_url: "https://news.ycombinator.com/item?id=48938858"
title: "Fixed Deep research and web fetch in Claude Code : full open source"
article_title: "GitHub - kunal12203/webify-mcp: Adaptive web research MCP skill for AI coding agents. pip install webify-mcp · GitHub"
author: "kunal122"
captured_at: "2026-07-16T19:57:44Z"
capture_tool: "hn-digest"
hn_id: 48938858
score: 1
comments: 0
posted_at: "2026-07-16T19:07:59Z"
tags:
  - hacker-news
  - translated
---

# Fixed Deep research and web fetch in Claude Code : full open source

- HN: [48938858](https://news.ycombinator.com/item?id=48938858)
- Source: [github.com](https://github.com/kunal12203/webify-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T19:07:59Z

## Translation

タイトル: クロード コードのディープ リサーチと Web フェッチを修正: 完全なオープン ソース
記事タイトル: GitHub - kunal12203/webify-mcp: AI コーディング エージェント向けのアダプティブ Web リサーチ MCP スキル。 pip インストール webify-mcp · GitHub
説明: AI コーディング エージェント用のアダプティブ Web リサーチ MCP スキル。 pip インストール webify-mcp - kunal12203/webify-mcp

記事本文:
GitHub - kunal12203/webify-mcp: AI コーディング エージェント用のアダプティブ Web リサーチ MCP スキル。 pip インストール webify-mcp · GitHub
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
クナル12203
/
webify-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション o

オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット docs docs .gitignore .gitignore BENCHMARKS.md BENCHMARKS.md ライセンス ライセンス README.md README.md benchmark_results.json benchmark_results.json mcp_server.py mcp_server.py pyproject.toml pyproject.toml webify.py webify.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AIコーディングエージェントのためのアダプティブウェブリサーチ
Deep Research の品質の 91%、コストの 5%、すべての MCP クライアントで動作
ドキュメント: 中文 · 日本語 · 한국어 · Español · हिन्दी · Français · Deutsch · Português · Русский
pip インストール webify-mcp
クロード mcp webify を追加 -- webify-mcp
それだけです。クロード コード、カーソル、VS コード、ウィンドサーフィン、および Zed で動作します。
フローチャートTB
エージェント[AI エージェント] -->|"web_find('query')"|ウェブ化
エージェント -->|"web_lookup(url, 'query')"|ウェブ化
Webify -->|"80–300 トークン"|エージェント
サブグラフ Webify[Webify MCP サーバー]
検索[検索\nBrave / DDG] --> グラフ[DOM 構造\nグラフ ビルダー]
グラフ --> 取得[BM25 + BFS\n取得]
取得 --> 合成[Haiku\nSynthesis]
終わり
スタイル Webify 塗りつぶし:#1a1a2e、ストローク:#16213e、カラー:#fff
スタイル エージェントの塗りつぶし:#0f3460、ストローク:#16213e、カラー:#fff
読み込み中
Web リサーチ用の 2 つのツール — どちらも全ページを読むよりも大幅に安価です。
web_find — マルチソース調査
フローチャート LR
A[クエリ] --> B[検索\nブレイブ / DDG]
B --> C1[ページ 1]
B --> C2[ページ 2]
B --> C3[ページ 3–6]
C1 --> D[DOM グラフ\n+ BM25]
C2 --> D
C3 --> D
D --> E[マルチアスペクト\ネクストラクション]
E --> F[俳句\n合成]
F --> G["答え\n(~800 トークン)"]
読み込み中
クエリの複雑さに合わせて深さを調整します。素朴な疑問が3つの情報源にヒットしました。多次元研究は、独立したサブアスペクト検索により 6+ まで拡張できます。深い調査レベルのカバレッジを得るために、焦点を絞ったサブクエリを使用して複数回呼び出します。
web_look

p — 単一ページの取得
フローチャート LR
A[URL + クエリ] --> B[ページを取得]
B --> C[DOM 構造\nグラフ]
C --> D[BM25 スコア]
D --> E[BFS トラバーサル]
E --> F["関連するサブツリー\n(80–300 トークン)"]
読み込み中
クエリに対してノードにスコアを付け、関連するサブツリーのみを返します。WebFetch がコンテキストに入れるフルページ テキストの 3,000 ～ 15,000 トークンではなく、80 ～ 300 トークンです。
Claude の Deep Research に対するブラインド A/B テスト — 15 個の未確認のクエリ、ランダムな順序、Sonnet ジャッジによるスコアリングの正確性 + 完全性 + 特異性 (各 1 ～ 5)。
Webify は常に正しい情報を見つけます。ギャップは常に完全性です — Deep Research の詳細をご覧ください。ほとんどのクエリでは、その違いは重要ではありません。徹底的に調査するには、web_find を複数回呼び出します。
インストールされると、AI は高価な組み込みツールの代わりに Webify を Web 調査に自動的に使用します。構成は必要ありません。設定ポリシーはパッケージ自体に埋め込まれています。
> Raft と Paxos のコンセンサスのトレードオフは何ですか?
→ クロードは web_find() を呼び出します — 検索、グラフの構築、答えの合成を行います
> GitHub API ドキュメントでレート制限を調べる
→ クロードは web_lookup() を呼び出します — そのページを取得し、関連するセクションのみを返します
ツール固有のセットアップ
pip インストール webify-mcp
クロード mcp webify を追加 -- webify-mcp
カーソル · ウィンドサーフィン · VS コード (コンティニュー/クライン) · ゼッド
{
"mcpサーバー": {
「ウェブ化」: {
"コマンド" : "webify-mcp"
}
}
}
設定ファイルの場所:
ウィンドサーフィン → ~/.windsurf/settings.json
VS Code / 続行 → ~/.Continue/config.json
Zed → ~/.config/zed/settings.json
pip install --upgrade webify-mcp
構成
環境変数
必須
説明
ANTHROPIC_API_KEY
web_find の場合
俳句総合
BRAVE_SEARCH_API_KEY
おすすめ
信頼性の高い検索 · 月あたり 2,000 回無料
WEBIFY_CACHE_DIR
いいえ
キャッシュディレクトリ · デフォルト ~/.cache/webify
検索: Brave API (

キーが設定されている場合) → DuckDuckGo Lite (無料フォールバック、キーは必要ありません)
macOS / Linux — ~/.zshrc または ~/.bashrc に追加します。
import ANTHROPIC_API_KEY= " sk-ant-... "
import BRAVE_SEARCH_API_KEY= " BSA... "
Windows (PowerShell):
[環境]::SetEnvironmentVariable( " ANTHROPIC_API_KEY " , " sk-ant-... " , " ユーザー " )
[環境]::SetEnvironmentVariable( " BRAVE_SEARCH_API_KEY " , " BSA... " , " ユーザー " )
MCP 構成内 (Webify にのみ適用):
{
"mcpサーバー": {
「ウェブ化」: {
"コマンド" : " webify-mcp " ,
"環境" : {
"ANTHROPIC_API_KEY" : " sk-ant-... " ,
"BRAVE_SEARCH_API_KEY" : "BSA..."
}
}
}
}
キーを取得します。
Anthropic → https://console.anthropic.com/settings/keys
Brave Search → https://brave.com/search/api/
# URL のグラフを作成する
python -m webify ビルド https://docs.python.org/3/library/json.html
# 特定の情報を調べる
python -m webify lookup https://docs.python.org/3/library/json.html " JSON 文字列を解析 "
Pythonライブラリ
Webify をインポートする
# マルチソース検索
結果 = Webify 。 web_find ( "mTLS はサービス メッシュでどのように機能しますか" )
print ( result [ "content" ]) # 合成された答え
print ( result [ "sources" ]) # [{url, タイトル, 信頼度, トークン}]
# 単一ページのルックアップ
結果 = Webify 。 Smart_lookup ( "https://docs.python.org/3/library/json.html" , "JSON を解析" )
print ( result [ "content" ]) # 関連セクションのみ (~376 トークン)
トラブルシューティング
webify-mcp # テストサーバー (終了するには Ctrl+C)
ls ~ /.cache/webify/ # キャッシュをチェックする
webify-mcp: コマンドが見つかりません → pip install webify-mcp を実行します
ツールが表示されない → 構成に追加した後、エディタを再起動します
web_find エラー → ANTHROPIC_API_KEY を設定
web_find は結果を返さない → DDG レート制限あり。 BRAVE_SEARCH_API_KEY を設定します
MIT · 著作権 © 2026 GrapeRoot
AI コーディング エージェント向けのアダプティブ Web リサーチ MCP スキル。 pip インストール webify-mcp

読み込み中にエラーが発生しました。このページをリロードしてください。
5
フォーク
レポートリポジトリ
リリース
1
Webify v0.0.2
最新の
2026 年 7 月 8 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Adaptive web research MCP skill for AI coding agents. pip install webify-mcp - kunal12203/webify-mcp

GitHub - kunal12203/webify-mcp: Adaptive web research MCP skill for AI coding agents. pip install webify-mcp · GitHub
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
kunal12203
/
webify-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits docs docs .gitignore .gitignore BENCHMARKS.md BENCHMARKS.md LICENSE LICENSE README.md README.md benchmark_results.json benchmark_results.json mcp_server.py mcp_server.py pyproject.toml pyproject.toml webify.py webify.py View all files Repository files navigation
Adaptive web research for AI coding agents
91% of Deep Research quality · 5% of the cost · Works in every MCP client
Docs: 中文 · 日本語 · 한국어 · Español · हिन्दी · Français · Deutsch · Português · Русский
pip install webify-mcp
claude mcp add webify -- webify-mcp
That's it. Works in Claude Code, Cursor, VS Code, Windsurf, and Zed.
flowchart TB
Agent[AI Agent] -->|"web_find('query')"| Webify
Agent -->|"web_lookup(url, 'query')"| Webify
Webify -->|"80–300 tokens"| Agent
subgraph Webify[Webify MCP Server]
Search[Search\nBrave / DDG] --> Graph[DOM Structural\nGraph Builder]
Graph --> Retrieve[BM25 + BFS\nRetrieval]
Retrieve --> Synthesize[Haiku\nSynthesis]
end
style Webify fill:#1a1a2e,stroke:#16213e,color:#fff
style Agent fill:#0f3460,stroke:#16213e,color:#fff
Loading
Two tools for web research — both dramatically cheaper than reading full pages:
web_find — multi-source research
flowchart LR
A[Query] --> B[Search\nBrave / DDG]
B --> C1[Page 1]
B --> C2[Page 2]
B --> C3[Page 3–6]
C1 --> D[DOM Graph\n+ BM25]
C2 --> D
C3 --> D
D --> E[Multi-aspect\nextraction]
E --> F[Haiku\nsynthesis]
F --> G["Answer\n(~800 tokens)"]
Loading
Adapts depth to query complexity. Simple questions hit 3 sources. Multi-dimensional research scales to 6+ with independent sub-aspect retrieval. Call it multiple times with focused sub-queries for deep-research-level coverage.
web_lookup — single-page retrieval
flowchart LR
A[URL + Query] --> B[Fetch page]
B --> C[DOM structural\ngraph]
C --> D[BM25 scoring]
D --> E[BFS traversal]
E --> F["Relevant subtree\n(80–300 tokens)"]
Loading
Scores nodes against your query, returns only the relevant subtree — 80–300 tokens instead of the 3,000–15,000 tokens of full page text WebFetch puts in context.
Blind A/B test against Claude's Deep Research — 15 unseen queries, randomized order, Sonnet judge scoring accuracy + completeness + specificity (1–5 each).
Webify finds correct information every time. The gap is always completeness — Deep Research reads more. For most queries that difference doesn't matter; for exhaustive research, call web_find multiple times.
Once installed, the AI automatically uses Webify for web research instead of expensive built-in tools — no configuration needed. The preference policy is embedded in the package itself.
> What are the tradeoffs between Raft and Paxos consensus?
→ Claude calls web_find() — searches, builds graphs, synthesizes answer
> Look up rate limits in the GitHub API docs
→ Claude calls web_lookup() — fetches that page, returns relevant sections only
Tool-specific setup
pip install webify-mcp
claude mcp add webify -- webify-mcp
Cursor · Windsurf · VS Code (Continue/Cline) · Zed
{
"mcpServers" : {
"webify" : {
"command" : " webify-mcp "
}
}
}
Config file locations:
Windsurf → ~/.windsurf/settings.json
VS Code / Continue → ~/.continue/config.json
Zed → ~/.config/zed/settings.json
pip install --upgrade webify-mcp
Configuration
Env var
Required
Description
ANTHROPIC_API_KEY
For web_find
Haiku synthesis
BRAVE_SEARCH_API_KEY
Recommended
Reliable search · free 2k/mo
WEBIFY_CACHE_DIR
No
Cache dir · default ~/.cache/webify
Search: Brave API (if key set) → DuckDuckGo Lite (free fallback, no key needed)
macOS / Linux — add to ~/.zshrc or ~/.bashrc :
export ANTHROPIC_API_KEY= " sk-ant-... "
export BRAVE_SEARCH_API_KEY= " BSA... "
Windows (PowerShell):
[ Environment ]::SetEnvironmentVariable( " ANTHROPIC_API_KEY " , " sk-ant-... " , " User " )
[ Environment ]::SetEnvironmentVariable( " BRAVE_SEARCH_API_KEY " , " BSA... " , " User " )
In your MCP config (applies only to Webify):
{
"mcpServers" : {
"webify" : {
"command" : " webify-mcp " ,
"env" : {
"ANTHROPIC_API_KEY" : " sk-ant-... " ,
"BRAVE_SEARCH_API_KEY" : " BSA... "
}
}
}
}
Get your keys:
Anthropic → https://console.anthropic.com/settings/keys
Brave Search → https://brave.com/search/api/
# Build a graph for a URL
python -m webify build https://docs.python.org/3/library/json.html
# Look up specific info
python -m webify lookup https://docs.python.org/3/library/json.html " parse JSON string "
Python library
import webify
# Multi-source search
result = webify . web_find ( "how does mTLS work in service meshes" )
print ( result [ "content" ]) # synthesized answer
print ( result [ "sources" ]) # [{url, title, confidence, tokens}]
# Single-page lookup
result = webify . smart_lookup ( "https://docs.python.org/3/library/json.html" , "parse JSON" )
print ( result [ "content" ]) # relevant sections only (~376 tokens)
Troubleshooting
webify-mcp # test server (Ctrl+C to exit)
ls ~ /.cache/webify/ # check cache
webify-mcp: command not found → Run pip install webify-mcp
Tool not showing up → Restart your editor after adding to config
web_find errors → Set ANTHROPIC_API_KEY
web_find returns no results → DDG rate-limited; set BRAVE_SEARCH_API_KEY
MIT · Copyright © 2026 GrapeRoot
Adaptive web research MCP skill for AI coding agents. pip install webify-mcp
There was an error while loading. Please reload this page .
5
forks
Report repository
Releases
1
Webify v0.0.2
Latest
Jul 8, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
