---
source: "https://github.com/Wisser/Jailer"
hn_url: "https://news.ycombinator.com/item?id=48539114"
title: "AI SQL Advisor integrated into a famous database tool"
article_title: "GitHub - Wisser/Jailer: Database Subsetting and Relational Data Browsing Tool. · GitHub"
author: "rwisser"
captured_at: "2026-06-15T11:11:08Z"
capture_tool: "hn-digest"
hn_id: 48539114
score: 1
comments: 0
posted_at: "2026-06-15T10:12:24Z"
tags:
  - hacker-news
  - translated
---

# AI SQL Advisor integrated into a famous database tool

- HN: [48539114](https://news.ycombinator.com/item?id=48539114)
- Source: [github.com](https://github.com/Wisser/Jailer)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T10:12:24Z

## Translation

タイトル: AI SQL Advisor を有名なデータベース ツールに統合
記事のタイトル: GitHub - Wisser/Jailer: データベースのサブセット化およびリレーショナル データの閲覧ツール。 · GitHub
説明: データベースのサブセット化およびリレーショナル データ参照ツール。 - ウィッサー/看守

記事本文:
GitHub - Wisser/Jailer: データベースのサブセット化およびリレーショナル データの閲覧ツール。 · GitHub
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
ウィッサー
/
看守
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知セットを変更するにはサインインする必要があります

ひずみ
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5,672 コミット 5,672 コミット .github .github 管理者 管理者 ブックマーク ブックマーク 構成 構成 データモデル データモデル デモ-DB デモ-DB ドキュメント ドキュメント サンプル サンプル 抽出モデル 抽出モデル レイアウト レイアウト lib lib maven-artifacts maven-artifacts src src テンプレート テンプレート tmp tmp .classpath .classpath .dockerignore .dockerignore .gitignore .gitignore .project .project Jailer.exe Jailer.exe README.md README.md build.xml build.xml driverlist.csv driverlist.csv ジェイラー.バット ジェイラー.バット ジェイラー.jar ジェイラー.jar ジェイラー.sh ジェイラー.sh ジェイラーGUI.バット ジェイラーGUI.バット ジェイラーGUI.sh ジェイラーGUI.sh ライセンス.txt License.txt releasenotes.txt releasenotes.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Jailer は、データベースのサブセット化とリレーショナル データの参照のためのツールです。
Subsetter は、データベースから小さなスライスを作成します (一貫性があり、参照が損なわれていない)
SQL (トポロジー的にソート)、DbUnit レコード、または XML として。
テスト データの小さなサンプルを作成したり、関連する運用データを使用してローカルな問題を分析したりするのに最適です。
データ ブラウザを使用すると、テーブル間の関係 (外部キー ベースまたはユーザー定義) に従ってデータベース内を移動できます。
一貫性があり、参照的に完全な行セットを本稼働データベースからエクスポートします。
そして、データを開発およびテスト環境にインポートします。
整合性を損なうことなく、古いデータを削除してアーカイブすることにより、データベースのパフォーマンスを向上させます。
トポロジー的にソートされた SQL-DML、階層構造の JSON、YAML、XML、および DbUnit データセットを生成します。
データの閲覧。外部キーベースまたはユーザー定義のリレーションシップに従って、データベースを双方向にナビゲートします。
コード補完、構文の強調表示、データベース メタデータを備えた SQL コンソール

視覚化。
AIクエリアシスタント。 Anthropic または OpenAI 互換 API を使用して、大規模なスキーマのスマートなテーブル選択を使用して、自然言語の質問から SQL を生成します。
SQL アドバイザ。 AI 支援を利用して既存の SQL クエリを分析、説明、リファクタリングします。インライン diff ビューには、何が変更されたのかが正確に表示されます。
構成の手間をかけずに第一印象を得ることができるデモ データベースが含まれています。
JDBC テクノロジーが使用されているため、原則としてあらゆる DMBS がサポートされます。ただし、最良の結果を得るには、特定の追加サポート機能が役立ちます。これらは次の用途に使用できます。
2026-05-31 SQL Advisor: AI 支援を利用して既存の SQL クエリを分析、説明、リファクタリングします。提案メニューには、最も一般的なタスク (説明、最適化、CTE としての書き換え、NULL 問題の検索など) が含まれています。結果は分割ビューとして表示され、左側に修正された SQL、右側に書式設定された説明が表示されます。インライン diff トグルは、SQL 領域で削除された行と追加された行を直接強調表示します。
2026-05-15 AI クエリ アシスタント: Anthropic または OpenAI 互換 API (OpenAI、Azure、Groq、Ollama、OpenRouter など) を使用し、大規模なスキーマのスマート テーブル選択を使用して自然言語の質問から SQL を生成します。
2024-07-04 データを構造化された JSON および YAML ファイルとしてもエクスポートできるようになりました。
2024-06-26 暗い環境での読みやすさを向上させるダーク UI テーマが導入されました。
2024-04-18 Liquibase ツールの統合により、データベース オブジェクトを作成するための DDL スクリプトを生成できるようになりました。これにより、オンボード手段のみを使用してサブセット データベースを最初から作成することが可能になります。
2023-02-03 ステートメントの詳細な分析のおかげで、SQL コンソールはクエリの結果をソース テーブルに関連付け、それに応じて表示できるようになりました。さらに、このテクニックにより、フィルター条件を動的にすることもできます。

任意の SQL クエリに追加されます。
2022-01-01 ユーザー インターフェイス全体の包括的な再設計と最新化。新しいルック アンド フィール FlatLaf 。
2021-02-04 親子関係のサイクルが検出され、壊れます。したがって、このようなデータは、Null 許容外部キーの挿入を延期することでエクスポートできます。
2020-02-04 Jailer エンジンが Maven リポジトリで公開されました。 https://mvnrepository.com/artifact/io.github.wisser/jailer-engine
2019-02-01 新しい「モデル移行ツール」を使用すると、この抽出モデルに対する最後の変更後にデータ モデルが拡張されている場合、新しく追加された関連付けを簡単に検索して編集できます。
2018-04-26 新機能「SQL 分析」は SQL ステートメントを分析し、関連付け定義を提案します。これにより、既存の SQL クエリに基づいてデータ モデルをリバース エンジニアリングできます。
2018-03-06 コード補完、構文の強調表示、データベース メタデータの視覚化を備えた SQL コンソール。
2017-05-10 新しい API により、データのエクスポートおよびインポート機能へのプログラムによるアクセスが提供されます。 https://wisser.github.io/Jailer/api.html
2017-03-30 フィルター管理を改善しました。テンプレートを使用すると、列にフィルターを割り当てるためのルールを定義できます。主キー列のフィルターは、対応する外部キー列に自動的に伝播されます。 https://wisser.github.io/Jailer/filters.html
2016-12-04 データを同じデータベースのスキーマに直接エクスポートできるようになりました。これにより、最適なパフォーマンスが保証されます。
2016-10-23 行を別の組み込みデータベースに収集することもできます。これにより、読み取り専用データベースからデータをエクスポートできるようになります。
2016-07-20 「例によるサブセット」機能を実装しました。データ ブラウザーを使用して抽出するすべての行を収集し、Jailer がそのサブセットのモデルを作成できるようにします。 https://wisser.github.io/Jailer/subset-by-example.html
2016-04-15 データの閲覧

erが導入されました。外部キーベースまたはユーザー定義のリレーションシップに従って、データベースを双方向にナビゲートします。
インストール ファイル「Jailer-database-tools-n.n.n.msi」(Windows の場合) または「jailer-database-tools_n.n.n-x64.deb」(Linux の場合) を使用します。
独自の Java インストールを使用する場合を除きます。または、コマンド ライン インターフェイス (CLI) を使用する場合も同様です。この場合、ファイル「jailer_n.n.n.zip」を解凍します。 https://wisser.github.io/Jailer/faq.html#multiuser も参照してください。
解凍された zip からツールを起動するには:
Windows プラットフォームでは、「Jailer.exe」を実行します。 「jailerGUI.bat」を起動することもできます。
Unix/Linux/macOS プラットフォームでは、スクリプト「jailerGUI.sh」を実行するか、「java -jarjailer.jar」を使用します。
git clone https://github.com/Wisser/Jailer.git
ツールを構築するには、ant を使用するだけです: ( https://ant.apache.org )
ホーム: https://github.com/Wisser/Jailer または http://jailer.sourceforge.net/
フォーラム: https://sourceforge.net/p/jailer/Discussion/
サポート: rwisser@users.sourceforge.net
このプロジェクトは、貢献してくださったすべての人々のおかげで存在します。
財政的に貢献して、コミュニティの維持にご協力ください。 [ 貢献 ]
あなたの組織でこのプロジェクトをサポートしてください。あなたのロゴが Web サイトへのリンクとともにここに表示されます。 [ 貢献 ]
データベースのサブセット化およびリレーショナル データ参照ツール。
wisser.github.io/看守
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
141
フォーク
レポートリポジトリ
リリース
260
看守 17.1.2
最新の
2026 年 6 月 15 日
+ 259 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
https://www.buymeacoffee.com/wisser
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Database Subsetting and Relational Data Browsing Tool. - Wisser/Jailer

GitHub - Wisser/Jailer: Database Subsetting and Relational Data Browsing Tool. · GitHub
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
Wisser
/
Jailer
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
5,672 Commits 5,672 Commits .github .github admin admin bookmark bookmark config config datamodel datamodel demo-db demo-db docs docs example example extractionmodel extractionmodel layout layout lib lib maven-artifacts maven-artifacts src src template template tmp tmp .classpath .classpath .dockerignore .dockerignore .gitignore .gitignore .project .project Jailer.exe Jailer.exe README.md README.md build.xml build.xml driverlist.csv driverlist.csv jailer.bat jailer.bat jailer.jar jailer.jar jailer.sh jailer.sh jailerGUI.bat jailerGUI.bat jailerGUI.sh jailerGUI.sh license.txt license.txt releasenotes.txt releasenotes.txt View all files Repository files navigation
Jailer is a tool for database subsetting and relational data browsing.
The Subsetter creates small slices from your database (consistent and referentially intact)
as SQL (topologically sorted), DbUnit records or XML.
Ideal for creating small samples of test data or for local problem analysis with relevant production data.
The Data Browser lets you navigate through your database following the relationships (foreign key-based or user-defined) between tables.
Exports consistent and referentially intact row-sets from your productive database
and imports the data into your development and test environment.
Improves database performance by removing and archiving obsolete data without violating integrity.
Generates topologically sorted SQL-DML, hierarchically structured JSON, YAML, XML and DbUnit datasets.
Data Browsing. Navigate bidirectionally through the database by following foreign-key-based or user-defined relationships.
SQL Console with code completion, syntax highlighting and database metadata visualization.
AI Query Assistant. Generate SQL from natural language questions using Anthropic or OpenAI-compatible APIs with smart table selection for large schemas.
SQL Advisor. Analyze, explain, and refactor existing SQL queries with AI assistance - inline diff view shows exactly what changed.
A demo database is included with which you can get a first impression without any configuration effort.
Thanks to the JDBC technology used, any DMBS is in principle supported. For best results, specific additional support features are useful, however. These are available for:
2026-05-31 SQL Advisor: Analyze, explain, and refactor existing SQL queries with AI assistance. A suggestions menu covers the most common tasks (explain, optimize, rewrite as CTEs, find NULL issues, and more). Results are shown as a split view with the revised SQL on the left and a formatted explanation on the right. An inline diff toggle highlights removed and added lines directly in the SQL area.
2026-05-15 AI Query Assistant: Generate SQL from natural language questions using Anthropic or OpenAI-compatible APIs (OpenAI, Azure, Groq, Ollama, OpenRouter, etc.) with smart table selection for large schemas.
2024-07-04 Data can now also be exported as structured JSON and YAML files.
2024-06-26 A dark UI theme has been introduced that improves readability in low light environments.
2024-04-18 DDL scripts for creating database objects can now be generated thanks to an integration of the Liquibase tool. This makes it possible to create subset databases from scratch using only on-board means.
2023-02-03 Thanks to deep analysis of statements, the SQL console can now relate the result of queries to the source tables and display them accordingly. In addition, this technique also allows filter conditions to be dynamically added to arbitrary SQL queries.
2022-01-01 Comprehensive redesign and modernization of the entire user interface. New Look & Feel FlatLaf .
2021-02-04 Cycles in parent-child relationships will be detected and broken. Thus, such data can be exported by deferring the insertion of nullable foreign keys.
2020-02-04 The Jailer engine is published in Maven repository. https://mvnrepository.com/artifact/io.github.wisser/jailer-engine
2019-02-01 The new "Model Migration Tool" allows you to easily find and edit the newly added associations if the data model has been extended after the last change to this extraction model.
2018-04-26 The new feature "Analyze SQL" analyzes SQL statements and proposes association definitions. This allows to reverse-engineer the data model based on existing SQL queries.
2018-03-06 SQL Console with code completion, syntax highlighting and database metadata visualization.
2017-05-10 New API provides programmatic access to the data export and import functionality. https://wisser.github.io/Jailer/api.html
2017-03-30 Improved filter management. Templates allows you to define rules for assigning filters to columns. Filters on primary key columns will automatically be propagated to the corresponding foreign key columns. https://wisser.github.io/Jailer/filters.html
2016-12-04 Data can now also be exported directly to a schema of the same database. This ensures optimal performance.
2016-10-23 Rows can alternatively be collected in a separate embedded database. This allows exporting data from read-only databases.
2016-07-20 Implemented the "Subset by Example" feature: Use the Data Browser to collect all the rows to be extracted and let Jailer create a model for that subset. https://wisser.github.io/Jailer/subset-by-example.html
2016-04-15 A Data Browser has been introduced. Navigate bidirectionally through the database by following foreign-key-based or user-defined relationships.
Use the installation file "Jailer-database-tools-n.n.n.msi" (for Windows) or "jailer-database-tools_n.n.n-x64.deb" (for Linux).
Unless you want to use your own Java installation. Or also if you want to use the command line interface (CLI). In this cases unzip the file "jailer_n.n.n.zip". See also https://wisser.github.io/Jailer/faq.html#multiuser
To start the tool from the unpacked zip:
On windows platform execute "Jailer.exe". You can also start "jailerGUI.bat".
On Unix/Linux/macOS platform execute the script "jailerGUI.sh" or use "java -jar jailer.jar"
git clone https://github.com/Wisser/Jailer.git
To build the tool you can just use ant: ( https://ant.apache.org )
Home: https://github.com/Wisser/Jailer or http://jailer.sourceforge.net/
Forum: https://sourceforge.net/p/jailer/discussion/
Support: rwisser@users.sourceforge.net
This project exists thanks to all the people who contribute.
Become a financial contributor and help us sustain our community. [ Contribute ]
Support this project with your organization. Your logo will show up here with a link to your website. [ Contribute ]
Database Subsetting and Relational Data Browsing Tool.
wisser.github.io/Jailer
Topics
There was an error while loading. Please reload this page .
141
forks
Report repository
Releases
260
Jailer 17.1.2
Latest
Jun 15, 2026
+ 259 releases
Sponsor this project
There was an error while loading. Please reload this page .
https://www.buymeacoffee.com/wisser
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
