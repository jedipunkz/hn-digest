---
source: "https://github.com/dburnett11155-rgb/Tokencompress"
hn_url: "https://news.ycombinator.com/item?id=49317274"
title: "Tokencompress A sub-2ms Go CLI and MCP sidecar that prunes AI agent tool context"
article_title: "GitHub - dburnett11155-rgb/Tokencompress · GitHub"
author: "ProffessorD"
captured_at: "2026-08-16T06:22:23Z"
capture_tool: "hn-digest"
hn_id: 49317274
score: 2
comments: 0
posted_at: "2026-08-16T05:55:12Z"
tags:
  - hacker-news
  - translated
---

# Tokencompress A sub-2ms Go CLI and MCP sidecar that prunes AI agent tool context

- HN: [49317274](https://news.ycombinator.com/item?id=49317274)
- Source: [github.com](https://github.com/dburnett11155-rgb/Tokencompress)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T05:55:12Z

## Translation

タイトル: Tokencompress AI エージェント ツール コンテキストをプルーニングする 2 ミリ秒未満の Go CLI および MCP サイドカー
記事のタイトル: GitHub - dburnett11155-rgb/Tokencompress · GitHub
説明: GitHub でアカウントを作成して、dburnett11155-rgb/Tokencompress の開発に貢献します。

記事本文:
GitHub - dburnett11155-rgb/Tokencompress · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
dburnett11155-rgb
/
トークン圧縮
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット Makefile Makefile README.md README.md COMPLEX.jSON COMPLEX.JSON go.MOD GO.MOD mcp_client.py mcp_client.py tc tc tokencompress tokencompress tokencompress.conf tokencompress.conf tokencompress.go tokencompress.go tokencompre

ss_test.go tokencompress_test.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのコンテキスト ウィンドウに入る前に未加工のツール出力 (JSON、ターミナル ログ、HTML) をプルーニングする、依存関係のない、ミリ秒未満の Go CLI および MCP サイドカー。
2 回目の LLM 要約ターンを使用せずに、LLM コンテキスト トークンの消費を 60% ～ 80% 削減します。
AI エージェントがツールを実行すると (API の呼び出し、ターミナル コマンドの実行、または Web ページのスクレイピング)、数千行の未解析の生データを受け取ります。
大規模な JSON ペイロード: 1,000 項目の配列により、プロンプトに 15,000 以上のトークンが溢れます。
詳細なスタック トレース: フレームワークのノイズとnode_modules パスがルート例外をかき消します。
生の HTML: CSS、スクリプト、ナビゲーション メニューによりコンテキスト ウィンドウが肥大化します。
これにより、API コストが高くなり、応答時間が遅くなり、エージェントがタスク中に幻覚を見せたり、ツール呼び出しを繰り返したりするコンテキストの腐敗が発生します。
ツールとモデルの間の高速かつ決定論的なフィルターとして機能します。
JSON の切り捨て: 代表的なスキーマの例を保持し、冗長な配列項目をメタデータ (: N) に置き換えます。
ログ プルーニング: 内部フレームワーク パスを削除し、ルート エラー メッセージ、ユーザー ファイル パス、および実行行のみを返します。
HTML クリーニング: スクリプト、スタイル、ナビゲーション要素を削除し、コンテンツを読み取り可能なテキスト/マークダウンに変換します。
重複ループ検出: エージェントが重複したツールの結果を 2 回続けて受信した場合、セッションごとにツールの出力をハッシュし、警告ヘッダーを先頭に追加します。
ビルドする
sudo mv tokencompress /usr/local/bin/
大きな JSON 応答を圧縮する
猫ラージ_response.json | tokencompress --mode json
詳細なスタック トレース ログを整理する
猫のアプリログ | tokencompress --mode log --log-internal-marker "mycompany/internal"
カール -s https://example.com | tokencompress --mode html
3. MCP (モデル コンテキスト プロトコル) の統合
クロードに直接追加

デスクトップ構成 ():
{
"mcpサーバー": {
"トークン圧縮": {
"コマンド": "/usr/local/bin/tokencompress",
"args": ["--mode", "mcp"]
}
}
}
入力ペイロード
未加工のトークン数
圧縮されたトークンの数
トークン削減
実行時間
JSON配列(500項目)
~12,500 トークン
~850トークン
93.2%
Python スタック トレース
~3,200 トークン
~410トークン
87.1%
HTML Webスクレイピング
~18,000 トークン
~2,100 トークン
88.3%
💼 カスタム統合とコンサルティング
カスタムの高性能 MCP プロキシ、エンタープライズ ツール用の特殊なパーサー、またはエージェント インフラストラクチャ用のカスタム トークン最適化セットアップが必要ですか?
カスタム セットアップ ギグ: カスタム統合セットアップごとに 50 ～ 00。
カスタム ログ/AST パーサー: カスタム ドメイン パーサーごとに 50。
連絡先: このリポジトリで問題を開くか、GitHub プロファイル経由で直接連絡します。
MITライセンス。オープンソースおよび商用エージェントのセットアップで無料で使用できます。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to dburnett11155-rgb/Tokencompress development by creating an account on GitHub.

GitHub - dburnett11155-rgb/Tokencompress · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
dburnett11155-rgb
/
Tokencompress
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits Makefile Makefile README.md README.md compressed.json compressed.json go.mod go.mod mcp_client.py mcp_client.py tc tc tokencompress tokencompress tokencompress.conf tokencompress.conf tokencompress.go tokencompress.go tokencompress_test.go tokencompress_test.go View all files Repository files navigation
A zero-dependency, sub-millisecond Go CLI and MCP sidecar that prunes raw tool outputs (JSON, terminal logs, HTML) before they enter your AI agent's context window.
Cut LLM context token consumption by 60% to 80% without using a second LLM summarization turn.
When AI agents execute tools (calling APIs, running terminal commands, or scraping web pages), they receive thousands of lines of raw, unparsed data:
Massive JSON Payloads: A 1,000-item array floods the prompt with 15,000+ tokens.
Verbose Stack Traces: Framework noise and node_modules paths drown out the root exception.
Raw HTML: CSS, scripts, and navigation menus bloat the context window.
This leads to high API costs, slower response times, and context rot —where agents hallucinate or repeat tool calls mid-task.
acts as a high-speed, deterministic filter between your tools and your model:
JSON Truncation: Keeps representative schema examples and replaces redundant array items with metadata (: N).
Log Pruning: Strips internal framework paths and returns only the root error message, user file paths, and execution lines.
HTML Cleaning: Strips scripts, styles, and navigation elements, converting content to readable text/markdown.
Duplicate Loop Detection: Hashes tool outputs per session and prepends a warning header if an agent receives duplicate tool results twice in a row.
make build
sudo mv tokencompress /usr/local/bin/
Compress a large JSON response
cat large_response.json | tokencompress --mode json
Prune a verbose stack trace log
cat app.log | tokencompress --mode log --log-internal-marker "mycompany/internal"
curl -s https://example.com | tokencompress --mode html
3. MCP (Model Context Protocol) Integration
Add directly to your Claude Desktop config ():
{
"mcpServers": {
"tokencompress": {
"command": "/usr/local/bin/tokencompress",
"args": ["--mode", "mcp"]
}
}
}
Input Payload
Raw Token Count
Compressed Token Count
Token Reduction
Execution Time
JSON Array (500 items)
~12,500 tokens
~850 tokens
93.2%
Python Stack Trace
~3,200 tokens
~410 tokens
87.1%
HTML Web Scrape
~18,000 tokens
~2,100 tokens
88.3%
💼 Custom Integrations & Consulting
Need a custom high-performance MCP proxy, specialized parsers for enterprise tools, or custom token-optimization setups for your agent infrastructure?
Custom Setup Gigs: 50 - 00 per custom integration setup.
Custom Log/AST Parsers: 50 per custom domain parser.
Contact: Open an issue on this repository or contact directly via GitHub profile.
MIT License. Free to use in open-source and commercial agent setups.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
