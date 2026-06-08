---
source: "https://www.slashdb.com/2026/06/08/slashdb-2-2-release-notes/"
hn_url: "https://news.ycombinator.com/item?id=48450005"
title: "SlashDB ver. 2.2 /w DuckDB, advanced filter and AI"
article_title: "SlashDB ver. 2.2 (Summer 2026)"
author: "agilevic"
captured_at: "2026-06-08T21:39:50Z"
capture_tool: "hn-digest"
hn_id: 48450005
score: 1
comments: 0
posted_at: "2026-06-08T19:07:16Z"
tags:
  - hacker-news
  - translated
---

# SlashDB ver. 2.2 /w DuckDB, advanced filter and AI

- HN: [48450005](https://news.ycombinator.com/item?id=48450005)
- Source: [www.slashdb.com](https://www.slashdb.com/2026/06/08/slashdb-2-2-release-notes/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T19:07:16Z

## Translation

タイトル：SlashDB ver. 2.2 /w DuckDB、高度なフィルター、AI
記事タイトル：SlashDB ver. 2.2 (2026 年夏)
説明: SlashDB 2.2 には、DuckDB サポート、高度なデータ検出フィルタリング、簡単なローカル セットアップ、およびモジュラー拡張フレームワーク上に構築された AI 拡張機能が追加されています。

記事本文:
スラッシュDB版2.2 (2026 年夏)
コンテンツにスキップ
ソリューション
HTML5 とモバイルアプリ
スラッシュDB版2.2 (2026 年夏)
SlashDB バージョン 2.2 がダウンロード可能になり、Azure 、AWS 、Docker ベースの環境で即時起動できるようになりました。
このリリースでは、分析ワークロードの DuckDB サポート、強化されたデータ検出エクスペリエンス、開発者向けの合理化されたローカル インストールと起動ワークフロー、パフォーマンス、診断、および使いやすさの継続的な向上など、製品全体にわたって実際的な改善が行われています。このリリースには合計 95 を超える作業項目が含まれています。
SlashDB アーキテクチャも進化し続けています。継続的なリファクタリングの取り組みにより、オプションでプラグインとしてインストールできる機能が増えており、コア プラットフォームをスリムでモジュール式にし、拡張しやすくするのに役立ちます。
アップグレード手順については、ユーザー ガイドの「アップグレード」の章を参照してください。
GET /db デモをスケジュールする
バージョン 2.2 の新機能は何ですか?
🔍 データディスカバリーの高度なフィルタリング
Data Discovery は、新しいフィルタリング パネルと拡張されたモディファイア コントロールを通じて、大幅に強力なフィルタリング エクスペリエンスを提供するようになりました。ユーザーは、SlashDB の API フィルタリング機能の多くに Web インターフェイスから直接アクセスできるため、リクエスト パスを手動で構築することなく、データセットを簡単に絞り込むことができます。
新しいエクスペリエンスは、高度なフィルタリングを日常の業務でより見つけやすく、より実践できるように設計されています。以下をサポートします。
列の値との一致 (全体および部分)
論理 (AND、OR、NOT) による正確な複数値フィルタリング
フィルタリングロジックを維持しながら、関連テーブル間のトラバーサル
結果セット内の列の選択
結果修飾子（個別、ストリーミング、転置、レコード数など）
テーブルヘッダーに、列の非表示、並べ替え順序の調整、高度なフィルターを開くためのコンテキスト メニューが追加されました。

鳴る。
🦆 分析ワークロードの DuckDB サポート
SlashDB 2.2 では、完全にサポートされるデータベース バックエンドとして DuckDB が追加されています。これにより、今日最も人気のあるオープンソース分析データベースの 1 つに構築された、新しい分析および組み込みのユースケースへの扉が開かれます。チームは、従来のリレーショナル バックエンドと同様に、SlashDB の API と Web インターフェイスを通じて DuckDB ベースのデータを公開できるようになりました。
📦 Windows へのローカルインストール
開発者とデータ エンジニアは、標準の pip ワークフローを使用して Windows ワークステーションに SlashDB をインストールし、必要なデータベースと機能の依存関係のみを組み込み、新しい init およびserve コマンドを使用してローカル インスタンスをより迅速に起動して実行できるようになりました。
これらのパッケージ化と CLI の改善により、SlashDB の評価が容易になり、プロトタイプの作成が迅速になり、最新の Python ベースの開発環境への統合が容易になります。以下は、わずか数回のキーストロークで Microsoft SQL Server および PostgreSQL データベースをサポートする Windows 上で SlashDB を起動する例です。
PS D:\> python -m venv sdbwin
PS D:\> .\sdbwin\Scripts\activate
(sdbwin) PS D:\> pip install --find-links https://downloads.slashdb.com/latestlashdb[mssql,postgresql]
リンクを参照: https://downloads.slashdb.com/latest
スラッシュデータベースの収集[mssql,postgresql]
[...]
alchemyjsonschema-0.8.0 attrs-25.4.0 beaker-1.13.0boltons-25.0.0 certifi-2025.11.12 cffi-2.0.0 charset-normalizer-2.1.1 click-8.3.3 colorama-0.4.6 concurrent-log-handler-0.9.29 が正常にインストールされました。 cryptography-48.0.0 dictknife-0.14.1 greenlet-3.2.4 hupper-1.12.1 idna-3.11 isodate-0.7.2 jsonschema-4.26.0 jsonschema-仕様-2025.9.1 jwcrypto-1.5.7 lxml-6.0.2 lz4-4.4.5 magicimport-0.9.1 mako-1.3.10 markupsafe-3.0.3 パッケージ化-26.2 passlib-1.7.4 ペースト-3.10.1 ペーストデプロイ-3.1.0 ペーストスクリプト-3.7.0 石膏-1.1.2 石膏ペースト

デプロイ-1.0.1 platformdirs-4.9.6 portalocker-3.2.0 psycopg2-2.9.12 pycparser-2.23 pyjwt-2.12.1 pyodbc-5.3.0 ピラミッド-1.10.8 ピラミッド-ビーカー-0.9 ピラミッド-exclog-1.1 ピラミッド-マコ-1.1.0 ピラミッド-マルチ認証-0.9.0 python-dateutil-2.9.0.post0 python-memcached-1.62 pytz-2025.2 pywin32-311 pyyaml-6.0.3 参照-0.36.2 リクエスト-2.34.2 rpds-py-0.27.1 setuptools-80.9.0 simplejson-4.1.1 six-1.17.0スラッシュdb-2.2.0 スラッシュdb-mssql-1.2.0 スラッシュdb-postgresql-1.2.0 sqlalchemy-2.0.49 sqlparse-0.5.5 strict-rfc3339-0.7 translationstring-1.4 testing-extensions-4.15.0 urllib3-1.26.20 venusian-3.1.1 waitress-3.0.2 webob-1.8.9 wsgicors-0.7.0 zope-deprecation-6.0 zope-interface-8.0.1 zxcvbn-4.4.28
(sdbwin) PS D:\>slashdb init
SlashDB をセットアップするディレクトリ [C:\Users\volx\AppData\Local\SlashDB]: d:\sdbwin\etc
ファイルを d:\sdbwin\etc にコピーしています
テンプレートの置き換え。
SlashDB は d:\sdbwin\etc で初期化されました
`slashdb-serve --ini d:\sdbwin\etc\slashdb.ini` を実行して SlashDB を起動します
(sdbwin) PS D:\>slashdb-serve --ini d:\sdbwin\etc\slashdb.ini
d:\sdbwin\etc\slashdb.ini の INI から SlashDB を起動する
http://0.0.0.0:6543 で配信中
PID 51908 でサーバーを起動しています。
🤖 AI アシスタント拡張機能の概念実証
SlashDB の拡張フレームワークは、AI 統合の概念実証を実証するためにも使用され、プラットフォームがサポートできるクライアント固有のカスタマイズの種類を強調しました。デモは非常に好評で、Google Gemini、Ollama、Microsoft Copilot、OpenAI ChatGPT との統合が含まれていました。
拡張フレームワークはコア SlashDB の一部ですが、AI 統合自体はその上に構築できるものの例です。お客様の環境でこれに興味がある場合、デモをご覧になりたい場合は、お問い合わせください。
AI エージェントには、新しいクエリ定義を生成する新しい機能が導入されました。

自然言語プロンプト、および既存のクエリを自然言語で説明します。ユーザーは、生成されたクエリを確認、承認、または変更することができ、完全に制御できるため、API エンドポイントとして即座に公開できます。
🔢 ユーザーインターフェイスのレコード数の切り替えとページネーションの改善
Data Discovery は、Web インターフェイスでレコード カウント機能を公開し、ページ数を自動的に計算するようになりました。
⚠️ 動作の変更と非推奨
特定の静的コンテンツの提供方法を変更します。以前のバージョンからアップグレードする場合は、必ず以下を編集してください。
nginx.conf からこのセクションを完全に削除します。
lashdb.ini で、static_path 変数を削除するか、次のように編集する必要があります。
static_path = スラッシュdb.views:static
アップグレード後も static_path の元の場所 (/var/lib/slashdb/www) がまだ存在する場合は、それをファイル システムから削除することもできます。
Python 3.9 はサポートされなくなりました。サポートされるバージョンは 3.10、3.11、および 3.12 になりました。
🛠️ 追加の改善と修正
転置された出力のスキーマにより、前回のリリースからのスキーマ拡張が完成しました
レコード数の表示と直接ページ ナビゲーションにより、UI のページネーションが改善されました。
サーバー側のカーソルとストリーミング コントロールは、Web UI でより柔軟でより明示的です。
データ検出と SQL パススルーの使いやすさが、パフォーマンスの向上やテーブル検索のエルゴノミクスなど、非常に大規模なスキーマ向けに改善されました。
データベース接続速度は、キャッシュされたスキーマでは速くなり、キャッシュされていないスキーマでは向上しました。
Slashdb-diag ツールキットは、診断出力の処理が改善され、非常に大きなログ ファイルに対してより効率的です。
リクエスト ステータスにより、中断されたリクエストまたは外部で失敗したリクエストの検出が改善され、リフレクト ステータス ポーリングのノイズが少なくなりました。
Databricks タイムアウト、in

有効なトークンとシャットダウンのシナリオでは、より意味のある動作とエラー処理が明らかになりました。
SQL パススルーとデータ検出は、転置された出力動作、フィルターとパラメーターの特殊文字、HTML 出力のベクトルと配列、列選択 URL の処理など、より多くのエッジ ケースを正しく処理します。
日付と時刻のフィールド、マッピングされたデータベース資格情報の欠落、無効なクエリの読み込み、現在のライセンスが無効な場合のライセンスの交換など、いくつかのワークフローでの検証およびエラー メッセージがより明確になります。
Windows 固有のさまざまな修正により、構成の保存、SQLite 接続、ローカル インストールの信頼性が向上します。
多数の UI と API のバグ修正により、構成、フィルタリング、データ入力、権限のワークフロー全体の一貫性が向上しました。
Linux カーネルの脆弱性 CopyFail、DirtyFrag/Fragnesia に対する防御策、依存関係の更新など、セキュリティ関連の修正。
ぜひご連絡ください。以下の方法でご連絡いただけます。
StackOverflow (質問を投稿するにはタグ「slashdb」を使用してください)
画像は Nano Banana 2 を使用して Google AI によって生成されました。画像の作成には実際の SlashDB コンテンツが使用されました。
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
X で共有 (新しいウィンドウで開きます)
×
LinkedIn で共有 (新しいウィンドウで開きます)
リンクトイン
Reddit で共有 (新しいウィンドウで開きます)
レディット
Pinterest で共有 (新しいウィンドウで開きます)
ピンタレスト
印刷（新しいウィンドウで開きます）
印刷する
デモをスケジュールする
© 著作権は VT Enterprise LLC にあります。無断転載を禁じます。

## Original Extract

SlashDB 2.2 adds DuckDB support, advanced Data Discovery filtering, easier local setup, and an AI extension built on its modular extension framework.

SlashDB ver. 2.2 (Summer 2026)
Skip to content
Solutions
HTML5 and Mobile Apps
SlashDB ver. 2.2 (Summer 2026)
SlashDB version 2.2 is now available for download and for instant launch on Azure , AWS , and Docker -based environments.
This release delivers a range of practical improvements across the product, including DuckDB support for analytical workloads, an enhanced Data Discovery experience, streamlined local installation and startup workflows for developers, and continued gains in performance, diagnostics, and usability. In total, this release includes over 95 work items.
SlashDB architecture also continues to evolve. Ongoing refactoring efforts are making more features optionally installable as plug-ins, helping keep the core platform lean, modular, and easier to extend.
For upgrade instructions please refer to Upgrading chapter in the User Guide.
GET /db Schedule a Demo
What’s new in version 2.2?
🔍 Advanced filtering in Data Discovery
Data Discovery now offers a significantly more powerful filtering experience through a new filtering panel and expanded modifier controls. Users can access more of SlashDB’s API filtering capabilities directly from the web interface, making it easier to refine datasets without manually building request paths.
The new experience is designed to make advanced filtering more discoverable and more practical in day-to-day work. It supports:
Column match to value (whole and partial)
Logical (AND, OR, NOT) for precise multi-value filtering
Traversal across related tables while preserving filtering logic
Column selection within result sets
Result modifiers, including distinct, streaming, transposing and record count
Table headers now feature a context menu for hiding columns, adjusting sort order and opening advanced filtering.
🦆 DuckDB support for analytical workdloads
SlashDB 2.2 adds DuckDB as a fully supported database backend. This opens the door to new analytical and embedded use cases built on one of today’s most popular open-source analytical databases. Teams can now expose DuckDB-based data through SlashDB’s API and web interface just as they do with more traditional relational backends.
📦 Local installation on Windows
Developers and data engineers can now install SlashDB with standard pip workflows on Windows workstations, include only the database and feature dependencies they need, and use the new init and serve commands to get local instances up and running more quickly.
These packaging and CLI improvements make SlashDB easier to evaluate, faster to prototype with, and simpler to integrate into modern Python-based development environments. Below is an example of launching SlashDB on Windows with support for Microsoft SQL Server and PostgreSQL databases in just a few keystrokes.
PS D:\> python -m venv sdbwin
PS D:\> .\sdbwin\Scripts\activate
(sdbwin) PS D:\> pip install --find-links https://downloads.slashdb.com/latest slashdb[mssql,postgresql]
Looking in links: https://downloads.slashdb.com/latest
Collecting slashdb[mssql,postgresql]
[...]
Successfully installed alchemyjsonschema-0.8.0 attrs-25.4.0 beaker-1.13.0 boltons-25.0.0 certifi-2025.11.12 cffi-2.0.0 charset-normalizer-2.1.1 click-8.3.3 colorama-0.4.6 concurrent-log-handler-0.9.29 cryptography-48.0.0 dictknife-0.14.1 greenlet-3.2.4 hupper-1.12.1 idna-3.11 isodate-0.7.2 jsonschema-4.26.0 jsonschema-specifications-2025.9.1 jwcrypto-1.5.7 lxml-6.0.2 lz4-4.4.5 magicalimport-0.9.1 mako-1.3.10 markupsafe-3.0.3 packaging-26.2 passlib-1.7.4 paste-3.10.1 pastedeploy-3.1.0 pastescript-3.7.0 plaster-1.1.2 plaster-pastedeploy-1.0.1 platformdirs-4.9.6 portalocker-3.2.0 psycopg2-2.9.12 pycparser-2.23 pyjwt-2.12.1 pyodbc-5.3.0 pyramid-1.10.8 pyramid-beaker-0.9 pyramid-exclog-1.1 pyramid-mako-1.1.0 pyramid-multiauth-0.9.0 python-dateutil-2.9.0.post0 python-memcached-1.62 pytz-2025.2 pywin32-311 pyyaml-6.0.3 referencing-0.36.2 requests-2.34.2 rpds-py-0.27.1 setuptools-80.9.0 simplejson-4.1.1 six-1.17.0 slashdb-2.2.0 slashdb-mssql-1.2.0 slashdb-postgresql-1.2.0 sqlalchemy-2.0.49 sqlparse-0.5.5 strict-rfc3339-0.7 translationstring-1.4 typing-extensions-4.15.0 urllib3-1.26.20 venusian-3.1.1 waitress-3.0.2 webob-1.8.9 wsgicors-0.7.0 zope-deprecation-6.0 zope-interface-8.0.1 zxcvbn-4.4.28
(sdbwin) PS D:\> slashdb init
Directory in which to set up SlashDB [C:\Users\volx\AppData\Local\SlashDB]: d:\sdbwin\etc
Copying files to d:\sdbwin\etc
Replacing templates.
SlashDB initialized at d:\sdbwin\etc
Run `slashdb-serve --ini d:\sdbwin\etc\slashdb.ini` to start SlashDB
(sdbwin) PS D:\> slashdb-serve --ini d:\sdbwin\etc\slashdb.ini
Starting SlashDB from INI at d:\sdbwin\etc\slashdb.ini
Serving at http://0.0.0.0:6543
Starting server in PID 51908.
🤖 AI assistant extension proof of concept
SlashDB’s extension framework was also used to demonstrate a proof‑of‑concept AI integration, highlighting the types of client‑specific customizations the platform can support. The demos were very well received and included integrations with Google Gemini, Ollama, Microsoft Copilot, and OpenAI ChatGPT.
The extension framework is part of core SlashDB, while the AI integration itself is an example of what can be built on top of it. If this is of interest for your environment, please contact us if you would like to see a demo .
AI agents were introduced with new capabilities to generate new query definitions from natural language prompts, as well as to explain existing queries in natural language. Users remain in full control, with the ability to review, accept, or modify generated queries, which can then be instantly exposed as API endpoints.
🔢 Record count toggle and improved pagination in user interface
Data Discovery now exposes the record count feature in the web interface and automatically calculates the number of pages.
⚠️ Behavior changes and deprecations
We are changing how certain static content is served. If you are upgrading from prior versions please make sure to edit the following.
From nginx.conf remove this section entirely.
In slashdb.ini you need to remove the static_path variable or edit the as follows:
static_path = slashdb.views:static
You can also remove the original location for the static_path from the file system (i.e. /var/lib/slashdb/www), if it still exists after the upgrade.
Python 3.9 is no longer supported. Supported versions are now 3.10, 3.11, and 3.12.
🛠️ Additional improvements and fixes
Schema for transposed output rounds out schema enhancements from last release
Pagination in the UI is improved with record count visibility and direct page navigation.
Server-side cursors and streaming controls are more flexible and more explicit in the web UI.
Data Discovery and SQL Pass-thru usability have been improved for very large schemas, including better performance and table search ergonomics.
Database connection speed is faster for cached schemas and improved for uncached schemas.
The slashdb-diag toolkit is more efficient on very large log files, with improved handling of diagnostic output.
Request Status has improved detection of interrupted or externally failed requests, and refelect status polling is less noisy.
Databricks timeout, invalid token and shutdown scenarios now surface more meaningful behavior and error handling.
SQL Pass-thru and Data Discovery handle more edge cases correctly, including transposed output behavior, special characters in filters and parameters, vectors and arrays in HTML output, and column selection URL handling.
Validation and error messages are clearer in several workflows, including date and time fields, missing mapped database credentials, invalid query loading, and license replacement when the current license is invalid.
A range of Windows-specific fixes improve configuration saving, SQLite connectivity, and local installation reliability.
Numerous UI and API bug fixes improve consistency across configuration, filtering, data entry, and permissions workflows.
Security-related fixes, including defensive measures against Linux kernel vulnerabilities CopyFail, DirtyFrag/Fragnesia and dependency updates.
We invite you to get in touch with us. You can reach us via:
StackOverflow (use tag “slashdb” to post questions)
The image ws generated by Google AI with Nano Banana 2. Actual SlashDB content was used in making of the image.
Share on Facebook (Opens in new window)
Facebook
Share on X (Opens in new window)
X
Share on LinkedIn (Opens in new window)
LinkedIn
Share on Reddit (Opens in new window)
Reddit
Share on Pinterest (Opens in new window)
Pinterest
Print (Opens in new window)
Print
Schedule a Demo
© Copyright by VT Enterprise LLC. All rights reserved.
