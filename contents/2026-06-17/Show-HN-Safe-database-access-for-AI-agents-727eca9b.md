---
source: "https://github.com/polymr-platform/mcp-sql"
hn_url: "https://news.ycombinator.com/item?id=48572665"
title: "Show HN: Safe database access for AI agents"
article_title: "GitHub - polymr-platform/mcp-sql · GitHub"
author: "its-a-new-world"
captured_at: "2026-06-17T18:58:53Z"
capture_tool: "hn-digest"
hn_id: 48572665
score: 1
comments: 0
posted_at: "2026-06-17T16:20:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Safe database access for AI agents

- HN: [48572665](https://news.ycombinator.com/item?id=48572665)
- Source: [github.com](https://github.com/polymr-platform/mcp-sql)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T16:20:56Z

## Translation

タイトル: HN の表示: AI エージェントの安全なデータベース アクセス
記事タイトル: GitHub - Polymr-platform/mcp-sql · GitHub
説明: GitHub でアカウントを作成して、polymr-platform/mcp-sql の開発に貢献します。
HN テキスト: HN さん、LLM エージェントにデータベースへの安全なアクセスを提供することを目的とした MCP サーバーを作成しました。これは、より広範な Polymr プラットフォーム ( https://polymr-platform.github.io/ ) の一部です。私は通常、LLM がコンテキストを確立するために必要に応じて選択を実行できるようにしていますが、特に機密性の高い環境でデータを変更することになったら、絶対的な制御と、LLM が行おうとしている内容の理解を維持したいと考えています。そこで私は mcp-sql を構築しました。機能の一部: - DML プレビュー: サーバーがトランザクションを実行し、すぐにロールバックするプレビュー モードでクエリを実行できます。これにより、変更を許可した場合に何が起こるかを正確に確認できます。 - 動的権限プロンプト: サーバーは SQL を動的にイントロスペクトし、必要な権限を決定します。エージェントが自動承認されていないクエリを実行しようとすると、ユーザーにプロンプ​​トが表示されます。エージェントの自律性レベルを完全に制御できます。 - 環境ルーティング: ポリシーを使用して、クエリが実行される接続を動的に切り替えることができ、環境間を簡単に切り替えることができます。特に実稼働データベースで MCP サーバーを使用している人々からのフィードバックをお待ちしています。

記事本文:
GitHub - Polymr-platform/mcp-sql · GitHub
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
ポリマープラットフォーム
/
mcp-sql
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く Fo

フォルダとファイル
7 コミット 7 コミット .github/ workflows .github/ workflows アセット アセット src/ main src/ main .gitignore .gitignore LICENSE LICENSE LICENSE.GPL LICENSE.GPL README.md README.md pom.xml pom.xml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
JDBC ベースの SQL アクセス用の MCP サーバー。 SQL の実行とスキーマのイントロスペクションのためのツールを提供します。
これは、polymr プラットフォームの一部です。
クエリを実行してトランザクションをロールバックすることによる、DML のプレビュー モード。
SQL を動的にイントロスペクションして、SQL を実行するためにどの権限が必要かを判断し、まだ付与されていない権限の入力を求めます。これにより、LLM エージェントに持たせる自律性のレベルを完全に制御できます。
たとえば、データベースへの書き込みが自動的に承認されない場合、LLM がどのような変更を行うかを正確に示すプレビューを取得できます。
複数の jdbc 接続 (プーリングを使用) を定義でき、ポリシーを使用して 1 つを動的に切り替えることができ、これにより環境を簡単に切り替えることができます。
データベースに到達するための SSH トンネリングをサポートします。
MCP アプリを使用して、クエリの結果をテーブルに表示します。
SQL ステートメント (SELECT/DML/DDL) を実行します。
params (オプション): ? の位置パラメータプレースホルダー
max_rows (オプション): SELECT の最大行数をオーバーライドします。
offset (オプション): ページネーションの行オフセット (SELECT のみ)
JSON 数値は JDBC 数値値としてバインドされます
JSON ブール値はブール値 JDBC 値としてバインドされます。
すべての要素がスカラー値の場合、JSON 配列は JDBC SQL 配列としてバインドされます。
埋め込みには JSON 配列で params を使用し、それを SQL で明示的にキャストします。
pgvector を使用した例: select * from document where embedding = ?::vector
距離演算子を使用した例: <-> ?::vector を埋め込むことでドキュメントから ID、コンテンツを並べ替える 制限 5
プレーンな PostgreSQL 配列列の場合は、SQL でネイティブ配列型を使用します (例: where タグ @> ?::)

テキスト[]
ネストされた JSON 配列は、ドライバーでサポートされている場合、ネストされた JDBC 配列として渡されますが、pgvector 自体は単一の 1 次元ベクトル値を想定しています。
pgvector セットアップで特定の入力形式が必要な場合は、データベースが変換を処理できるように SQL でのキャストを維持します。
スキーマ、テーブル、列、および関係を検査します。データベースベンダーを返します。
検索を省略した場合、ツールは列の詳細を含まずにオブジェクト名のみをリストします。
mvn -q -DskipTests パッケージ
ネイティブ実行可能ファイルをビルドします (GraalVM が必要):
mvn -q -DskipTests -Pnative パッケージ
バイナリは target/mcp-sql (Windows の場合は target/mcp-sql.exe) に書き込まれます。
java -jar target/mcp-sql-0.1.0-SNAPSHOT.jar --print-config-schema
について
LGPL-3.0、GPL-3.0 ライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0
最新の
2026 年 6 月 17 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to polymr-platform/mcp-sql development by creating an account on GitHub.

Hey HN, I created an MCP server aimed at giving LLM agents safe access to a database. It's part of the broader Polymr platform ( https://polymr-platform.github.io/ ). I generally allow the LLM to perform selects as it wishes to establish context, but once it comes down to modifying data, especially in a sensitive environment, I want to retain absolute control and understanding of what it's trying to do. That's why I built mcp-sql. Some of the features: - DML Previews: Queries can be run in preview mode where the server executes the transaction and immediately rolls it back. This lets you see exactly what will happen if you allow the change. - Dynamic Permission Prompts: The server dynamically introspects SQL to determine required permissions. If an agent tries to execute a query it doesn't have auto-approval for, it prompts the user. You retain full control over the agent's autonomy level. - Environment Routing: You can use policies to dynamically switch the connection a query is run on, allowing you to easily switch between environments. I'd love feedback, especially from people using MCP servers with production databases.

GitHub - polymr-platform/mcp-sql · GitHub
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
polymr-platform
/
mcp-sql
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows assets assets src/ main src/ main .gitignore .gitignore LICENSE LICENSE LICENSE.GPL LICENSE.GPL README.md README.md pom.xml pom.xml View all files Repository files navigation
MCP server for JDBC-backed SQL access. It provides tools for SQL execution and schema introspection.
It is part of the polymr platform .
Preview mode for DML by running query and rolling back transaction.
Dynamic introspection of SQL to determine which permissions are needed to run it, prompting for the ones that are not granted yet. This gives you full control to choose the level of autonomy you want the LLM agent to have.
For example if database writing is not automatically approved, you can get a preview showing you exactly what change the LLM wants to do.
You can define multiple jdbc connections (with pooling), you can then use a policy to dynamically switch one is the default, allowing easy environmental switching.
Supports SSH tunneling to reach databases.
Uses MCP Apps to show the results of a query in a table.
Run SQL statements (SELECT/DML/DDL).
params (optional): positional parameters for ? placeholders
max_rows (optional): override max rows for SELECT
offset (optional): row offset for pagination (SELECT only)
JSON numbers bind as numeric JDBC values
JSON booleans bind as boolean JDBC values
JSON arrays bind as JDBC SQL arrays when all elements are scalar values
Use params with a JSON array for the embedding, then cast it explicitly in SQL.
Example with pgvector: select * from documents where embedding = ?::vector
Example with distance operators: select id, content from documents order by embedding <-> ?::vector limit 5
For plain PostgreSQL array columns, use a native array type in SQL, for example where tags @> ?::text[]
Nested JSON arrays are passed through as nested JDBC arrays when supported by the driver, but pgvector itself expects a single 1-D vector value.
If your pgvector setup requires a specific input format, keep the cast in SQL so the database handles the conversion.
Inspect schemas, tables, columns, and relationships. Returns the database vendor.
If search is omitted, the tool only lists object names without column details.
mvn -q -DskipTests package
Build a native executable (requires GraalVM):
mvn -q -DskipTests -Pnative package
The binary is written to target/mcp-sql (or target/mcp-sql.exe on Windows).
java -jar target/mcp-sql-0.1.0-SNAPSHOT.jar --print-config-schema
About
LGPL-3.0, GPL-3.0 licenses found
Licenses found
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.1.0
Latest
Jun 17, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
