---
source: "https://github.com/arm/metis"
hn_url: "https://news.ycombinator.com/item?id=48327316"
title: "ARM Open Sources AI-Powered Security Code Review"
article_title: "GitHub - arm/metis: Metis is an open-source, AI-driven tool for deep security code review · GitHub"
author: "ARob109"
captured_at: "2026-05-30T11:40:50Z"
capture_tool: "hn-digest"
hn_id: 48327316
score: 7
comments: 0
posted_at: "2026-05-29T18:28:51Z"
tags:
  - hacker-news
  - translated
---

# ARM Open Sources AI-Powered Security Code Review

- HN: [48327316](https://news.ycombinator.com/item?id=48327316)
- Source: [github.com](https://github.com/arm/metis)
- Score: 7
- Comments: 0
- Posted: 2026-05-29T18:28:51Z

## Translation

タイトル: ARM オープンソース AI を活用したセキュリティ コード レビュー
記事のタイトル: GitHub - arm/metis: Metis は、ディープ セキュリティ コード レビューのためのオープンソースの AI 駆動ツールです · GitHub
説明: Metis は、詳細なセキュリティ コード レビューのためのオープンソースの AI 駆動ツールです - arm/metis

記事本文:
GitHub - arm/metis: Metis は、詳細なセキュリティ コード レビューのためのオープンソースの AI 駆動ツールです · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
腕
/
メティス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

オプションについて
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
195 コミット 195 コミット .github .github docker/ db-init docker/ db-init docs docs 例 例 src/ metis src/ metis テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CONTRIBUTION.md CONTRIBUTION.md Dockerfile Dockerfileライセンス ライセンス MANIFEST.in MANIFEST.in README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml project.toml project.toml pyproject.toml pyproject.toml pytest.ini pytest.ini すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Metis: AI を活用したセキュリティ コードのレビュー
Metis は、Arm の製品セキュリティ チームによって作成された、詳細なセキュリティ コード レビューのためのオープンソースのエージェント AI セキュリティ フレームワークです。これは、エンジニアが微妙な脆弱性を検出し、安全なコーディング方法を改善し、レビューの疲労を軽減するのに役立ちます。これは、従来のツールでは不十分な場合が多い、大規模で複雑なコードベース、または従来のコードベースで特に価値があります。
メティスは、知恵、深い思考、助言を司るギリシャの女神にちなんで名付けられました。
深い推論
リンターや従来の静的分析ツールとは異なり、Metis はハードコードされたルールに依存しません。意味の理解と推論が可能な LLM を使用します。
コンテキストを意識したレビュー
RAG により、モデルがより広範なコード コンテキストと関連ロジックにアクセスできるようになり、より正確で実用的な提案が得られます。
プラグインフレンドリーで拡張可能
拡張性を念頭に置いて設計されており、追加の言語、モデル、新しいプロンプトのサポートが簡単です。
問題の検証
独自の分析とサードパーティの SAST ツールからの結果を検証し、誤検知を減らすための証拠を収集します。
プロバイダーの柔軟性
主要な LLM サービスとローカル モデル (vLLM、Ollama、LiteLLM など) のサポート。 vLLM ガイドと Ollama gui を参照してください。

ローカル設定の例については de を参照してください。
Metis には次の言語のサポートが含まれています。
トリアージ分析の詳細 (フロー分析と構造分析) については、 docs/triage-flow.md を参照してください。
Metis はプラグインベースの言語システムを使用しているため、サポートを追加の言語に簡単に拡張できます。
また、pgvector を備えた PostgreSQL や ChromaDB など、複数のベクター ストア バックエンドもサポートしています。
デフォルトでは、Metis はローカルでセットアップなしで使用するために ChromaDB を使用します。スケーラブルなインデックス作成とマルチプロジェクトのサポートのために PostgreSQL (pgvector を使用) を使用することもできます。
リポジトリのクローンを作成した後、仮想環境を作成するか、システム全体に依存関係をインストールできます。
仮想環境を使用するには (推奨):
紫外線対策
uv pip インストール 。
または --system を使用してシステム全体にインストールします。
uv pip インストール 。 --システム
PostgreSQL (pgvector) バックエンド サポートを使用してインストールするには:
uv pip インストール ' .[postgres] '
1.1 ドッカー
git clone https://github.com/arm/metis.git
CDメティス
docker build -t metis 。
2.LLMプロバイダーのセットアップ
Metis を使用する前に OpenAI API キーをエクスポートします。
import OPENAI_API_KEY= " あなたのキーをここに "
3. インデックス付けと実行分析
分析するソースへのパスも指定して metis を実行します。
uv run metis --codebase-path <path_to_src>
次に、次を使用してコードベースにインデックスを付けます。
インデックス
最後に、次のコマンドを使用して、コードベース全体でセキュリティ分析を実行します。
レビューコード
インデックスが使用できないにもかかわらず分析を実行したい場合は、次を使用します。
review_code --ignore-index
これは、 review_code 、 review_file 、 review_patch 、および triage でのみサポートされます。このモードでは、Metis は取得をスキップし、関連コンテキストの検索が無効になっていることを警告します。
コードベースのパスに移動して、次を実行します。
docker run --rm -it -v ` pwd ` :/metis metis
環境変数を渡すには、 -e を使用します。
docker run --rm -it -v ` pwd ` :/metis -e " OPENAI_API_KEY= ${OPENAI_API_KEY} " metis
あなたはできます

metis への ass 引数:
docker run --rm -it -v ` pwd ` :/metis metis --non-interactive --command ' review_code ' --output-file results/review_code_results.json
構成
Metis 構成 ( metis.yaml )
Metis 構成は、metis の実行時に作業ディレクトリ内の YAML 構成ファイル ( metis.yaml ) を使用して上書きできます。デフォルトの構成は src/metis/metis.yaml にあります。このファイルは、以下を含むすべての実行時パラメータを定義します。
LLM プロバイダー: OpenAI モデル名、埋め込みモデル、トークン制限
エンジンの動作: 最大ワーカー、最大トークン長、類似度 top-k
データベース接続: PostgreSQL の場合: ホスト、ポート、認証情報、およびスキーマ名
ベクターのインデックス作成: pgvector の HNSW パラメーター
このファイルは Metis を実行するために必要であり、展開ごとにカスタマイズする必要があります。
プロンプト構成 ( plugins.yaml )
Metis は、plugins.yaml ファイルを使用して、LLM プロンプト テンプレートやドキュメント分割ロジックなどの言語固有の動作を定義します。
各言語プラグイン (C など) は、このファイルを参照してロードします。
次のプロンプトのようないくつかのプロンプトをカスタマイズできます。
security_review : コードまたは差分のセキュリティ監査を実行するように LLM をガイドします。
validation_review : 生成されたレビューの正確性または品質を評価するように LLM に依頼します。
security_review_checks : LLM が検索しようとするすべてのセキュリティ問題のリスト。
これらのプロンプトは LLM に自然言語コンテキストを提供し、ユースケース (厳格な監査、プライバシー レビュー、コンプライアンスなど) に合わせて調整できます。
ソース コードとドキュメントのチャンク パラメーターを構成することもできます。
chunk_lines : チャンクあたりの行数
chunk_lines_overlap : チャンク間のオーバーラップ
max_chars : チャンクあたりの最大文字数
Metis は、Setuptools エントリ ポイントを使用して言語プラグインを検出します。パッケージは宣言によってプラグインを公開できます

パッケージ化メタデータにグループ metis.plugins を追加します。各エントリは metis.plugins.base.BaseLanguagePlugin を実装するクラスに解決される必要があり、オプションでコンストラクターで plugin_config を受け入れる必要があります。
サードパーティプラグインの pyproject.toml の例:
[project.entry-points."metis.plugins"]
my_lang = "my_pkg.my_module:MyLanguagePlugin"
ランニングメティス
Metis は、いくつかの組み込みコマンドを備えた対話型 CLI を提供します。起動後、次を実行できます。
--custom-prompt PATH – 追加のガイダンスを含むオプションの .md または .txt ファイル。提供されると、Metis はそれを 1 回ロードし、そのテキストをすべてのセキュリティ レビュー プロンプトに組み込みます。フラグが省略された場合、Metis はプロジェクト ルートで .metis.md を検索し、存在する場合はそれを使用します。これを使用して、 plugins.yaml を編集せずに組織固有のポリシーまたはセキュリティ要件を挿入します。
--backend chroma|postgres – Vector-store バックエンド (デフォルト chroma ) を選択します。
--project-schema / --chroma-dir – バックエンド固有のノブ。
--triage – review_code 、 review_file 、または review_patch の後、結果を優先順位付けし、SARIF 出力に注釈を付けます。
--include-triaging – トリアージを実行するときに、Metis によってすでにトリアージされた結果を含めます。
--ignore-index – review_code 、 review_file 、 review_patch 、および Triage をインデックスに基づくコンテキストなしで実行できるようにします。 Metis は警告を発し、このモードでは取得をスキップします。 ask または update には適用されません。
--verbose 、 --quiet 、 --output-file 、 --output-files – ログ記録とエクスポート形式を制御します。
コードベースをベクトル データベースにインデックス付けします。分析の前に実行する必要があります。
インデックス付きコードベースの完全なセキュリティ レビューを実行します。
使用可能なインデックスがない場合に取得せずに実行するには、--ignore-index を使用します。
単一ファイルの対象を絞ったセキュリティ レビューを実行します。
使用可能なインデックスがない場合に取得せずに実行するには、--ignore-index を使用します。
レビュー

diff/patch ファイルを作成し、変更によってもたらされる潜在的なセキュリティ問題を強調表示します。
使用可能なインデックスがない場合に取得せずに実行するには、--ignore-index を使用します。
diff を使用してインデックスを増分更新します。完全なインデックスの再作成を回避します。
インデックス付きコードベースについて何でも Metis に質問してください。アーキテクチャの調査、設計パターンの特定、ロジックの明確化に役立ちます。
SARIF ファイル内の結果を優先順位付けし、各結果に Metis 優先順位付けメタデータで注釈を付けます。
このコマンドは、Metis または他のセキュリティ/静的分析ツールによって生成された SARIF で使用できます。
使用可能なインデックスがない場合に、取得せずに優先順位を付けるには、--ignore-index を使用します。
トリアージの仕組みの概要については、docs/triage-flow.md を参照してください。
非対話型モードで実行する
Metis は、自動化、CI/CD パイプライン、またはスクリプトによる使用に役立つ非対話型モードもサポートしています。
非対話モードで Metis を使用するには、 --non-interactive フラグを --command とともに使用します。
metis --non-interactive --command " <コマンド> [引数...] " [--output-file < file.json > ]
例
例 1: クロマ (デフォルト)
metis --codebase-path < path_to_src >
例 2: Postgres
デフォルトの ChromaDB バックエンドを使用したくない場合は、ローカル インストールまたは提供された Docker セットアップを使用して PostgreSQL に切り替えることができます。
ドッカー構成 -d
これにより、 docker-compose.yml で指定された認証情報を使用して、pgvector 拡張機能が有効になった PostgreSQL インスタンスが起動します。
次に、PostgreSQL バックエンドで Metis を実行します。
メティス\
--プロジェクトスキーマ myproject_main \
--codebase-path < path_to_src > \
--バックエンド postgres
例 3: 使用法と出力
> review_file src/memory/remap.c
脆弱なソースコード:
// メモリ アドレスをある領域から別の領域に再マップする
for ( uint32_t * ptr = 開始 ; ptr < 終了 ; ptr ++ ) {
uint32_t 値 = * ptr ;
if ( 値 >= OLD_REGION_BASE && 値 < OL

D_REGION_BASE + REGION_SIZE ) {
値 = 値 - OLD_REGION_BASE + NEW_REGION_BASE ;
}
}
出力例:
ファイル: src/memory/remap.c
特定された問題 1: アドレス再マッピング ループでメモリが更新されない
スニペット:
for (uint32_t * ptr = 開始 ; ptr < 終了 ; ptr++) {
uint32_t 値 = * ptr ;
もし...
理由: remap_address_table 関数では、コードは古いメモリ領域から新しいメモリ領域へのアドレス参照を調整することを目的としています。ただし、ローカル変数 ' value ' に格納されている更新された値は、メモリのポインター位置 ( * ptr) に書き戻されることはありません。これは、アドレス エントリが変更されないことを意味します。システムがこれらの値が正しく再配置されることに依存している場合、意図しない動作が発生する可能性があります。
軽減策: 新しいアドレスを計算した後に値が書き戻されるようにループを更新します。たとえば:
for (uint32_t * ptr = 開始 ; ptr < 終了 ; ptr++) {
uint32_t 値 = * ptr ;
if (値 > = OLD_REGION_BASE && 値 < OLD_REGION_BASE + REGION_SIZE) {
値 = (( 値 - OLD_REGION_BASE) + NEW_REGION_BASE);
*ptr = 値;
}
}
これにより、再配置されたメモリ領域を指すように各エントリが適切に更新されることが保証されます。
信頼度: 1 。 0
例 4: 完全なセキュリティ レビューを実行する (非対話型)
metis --non-interactive --comma

[切り捨てられた]

## Original Extract

Metis is an open-source, AI-driven tool for deep security code review - arm/metis

GitHub - arm/metis: Metis is an open-source, AI-driven tool for deep security code review · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
arm
/
metis
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
195 Commits 195 Commits .github .github docker/ db-init docker/ db-init docs docs examples examples src/ metis src/ metis tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CONTRIBUTION.md CONTRIBUTION.md Dockerfile Dockerfile LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml project.toml project.toml pyproject.toml pyproject.toml pytest.ini pytest.ini View all files Repository files navigation
Metis: AI-Powered Security Code Review
Metis is an open-source, agentic AI security framework for deep security code review, created by Arm's Product Security Team . It helps engineers detect subtle vulnerabilities, improve secure coding practices, and reduce review fatigue. This is especially valuable in large, complex, or legacy codebases where traditional tooling often falls short.
Metis is named after the Greek goddess of wisdom, deep thought and counsel.
Deep Reasoning
Unlike linters or traditional static analysis tools, Metis doesn’t rely on hardcoded rules. It uses LLMs capable of semantic understanding and reasoning.
Context-Aware Reviews
RAG ensures that the model has access to broader code context and related logic, resulting in more accurate and actionable suggestions.
Plugin-Friendly and Extensible
Designed with extensibility in mind: support for additional languages, models, and new prompts is straightforward.
Issue validation
Validates findings from its own analysis and third-party SAST tools, gathering evidence to reduce false positives.
Provider Flexibility
Support for major LLM services and local models (vLLM, Ollama, LiteLLM etc.). See the vLLM guide and the Ollama guide for local setup examples.
Metis includes support for the following languages:
For triage analysis details ( Flow Analysis vs Structural Analysis ), see docs/triage-flow.md .
Metis uses a plugin-based language system, making it easy to extend support to additional languages.
It also supports multiple vector store backends, including PostgreSQL with pgvector and ChromaDB.
By default, Metis uses ChromaDB for local, no-setup usage. You can also use PostgreSQL (with pgvector) for scalable indexing and multi-project support.
After cloning the repository, you can either create a virtual environment or install dependencies system-wide.
To use a virtual environment (recommended):
uv venv
uv pip install .
or install system wide using --system:
uv pip install . --system
To install with PostgreSQL (pgvector) backend support:
uv pip install ' .[postgres] '
1.1 Docker
git clone https://github.com/arm/metis.git
cd metis
docker build -t metis .
2. Set up LLM Provider
Export your OpenAI API key before using Metis:
export OPENAI_API_KEY= " your-key-here "
3. Index and Run Analysis
Run metis by also providing the path to the source you want to analyse:
uv run metis --codebase-path <path_to_src>
Then, index your codebase using:
index
Finally, run the security analysis across the entire codebase with:
review_code
If the index is unavailable and you still want to run an analysis, use:
review_code --ignore-index
This is supported only for review_code , review_file , review_patch , and triage . In that mode Metis skips retrieval and warns that relevant-context lookup was disabled.
Go to your codebase path and run:
docker run --rm -it -v ` pwd ` :/metis metis
To pass environment variables use -e :
docker run --rm -it -v ` pwd ` :/metis -e " OPENAI_API_KEY= ${OPENAI_API_KEY} " metis
You can pass arguments to metis:
docker run --rm -it -v ` pwd ` :/metis metis --non-interactive --command ' review_code ' --output-file results/review_code_results.json
Configuration
Metis Configuration ( metis.yaml )
Metis configuration can be over-ridden using a YAML configuration file ( metis.yaml ) in the working directory when running metis. The default configuration is in src/metis/metis.yaml. This file defines all runtime parameters including:
LLM provider: OpenAI model names, embedding models, token limits
Engine behavior: max workers, max token length, similarity top-k
Database connection: In the case of PostgreSQL: host, port, credentials, and schema name
Vector indexing: HNSW parameters for pgvector
This file is required to run Metis and should be customized per deployment.
Prompt Configuration ( plugins.yaml )
Metis uses a plugins.yaml file to define language-specific behavior, including LLM prompt templates and document splitting logic.
Each language plugin (e.g., C) references this file to load:
You can customize a number of prompts like the following prompts:
security_review : Guides the LLM to perform a security audit of code or diffs.
validation_review : Asks the LLM to assess the correctness or quality of a generated review.
security_review_checks : A list of all the security issues the LLM will try to search for.
These prompts provide natural language context for the LLM and can be tailored to your use case (e.g., stricter audits, privacy reviews, compliance).
You can also configure the chunking parameters for source code and documentation:
chunk_lines : Number of lines per chunk
chunk_lines_overlap : Overlap between chunks
max_chars : Max characters per chunk
Metis discovers language plugins using Setuptools entry points. Packages can expose plugins by declaring the group metis.plugins in their packaging metadata. Each entry should resolve to a class implementing metis.plugins.base.BaseLanguagePlugin and optionally accept plugin_config in the constructor.
Example pyproject.toml for a third-party plugin:
[project.entry-points."metis.plugins"]
my_lang = "my_pkg.my_module:MyLanguagePlugin"
Running Metis
Metis provides an interactive CLI with several built-in commands. After launching, you can run the following:
--custom-prompt PATH – optional .md or .txt file that contains additional guidance. When provided, Metis loads it once and weaves the text into every security-review prompt. If the flag is omitted, Metis looks for .metis.md in your project root and uses it when present. Use this to inject organization-specific policy or security requirements without editing plugins.yaml .
--backend chroma|postgres – choose vector-store backend (default chroma ).
--project-schema / --chroma-dir – backend-specific knobs.
--triage – after review_code , review_file , or review_patch , triage findings and annotate SARIF output.
--include-triaged – include findings already triaged by Metis when running triage.
--ignore-index – allow review_code , review_file , review_patch , and triage to run without index-backed context. Metis warns and skips retrieval in this mode. It does not apply to ask or update .
--verbose , --quiet , --output-file , --output-files – control logging and export formats.
Indexes your codebase into a vector database. Must be run before any analysis.
Performs a full security review of the indexed codebase.
Use --ignore-index to run without retrieval when no index is available.
Performs a targeted security review of a single file.
Use --ignore-index to run without retrieval when no index is available.
Reviews a diff/patch file and highlights potential security issues introduced by the change.
Use --ignore-index to run without retrieval when no index is available.
Incrementally updates the index using a diff. Avoids full reindexing.
Ask Metis anything about the indexed codebase. Useful for exploring architecture, identifying design patterns, or clarifying logic.
Triages findings in a SARIF file and annotates each result with Metis triage metadata.
You can use this command on SARIF generated by Metis or by other security/static-analysis tools.
Use --ignore-index to triage without retrieval when no index is available.
See docs/triage-flow.md for a short overview of how triage works.
Running in Non-Interactive Mode
Metis also supports a non-interactive mode, useful for automation, CI/CD pipelines, or scripted usage.
To use Metis in non-interactive mode, use the --non-interactive flag along with --command:
metis --non-interactive --command " <command> [args...] " [--output-file < file.json > ]
Examples
Example 1: Chroma (default)
metis --codebase-path < path_to_src >
Example 2: Postgres
If you prefer not to use the default ChromaDB backend, you can switch to PostgreSQL either using a local installation or the provided Docker setup.
docker compose up -d
This will launch a PostgreSQL instance with the pgvector extension enabled, using the credentials specified in your docker-compose.yml .
Then, run Metis with the PostgreSQL backend:
metis \
--project-schema myproject_main \
--codebase-path < path_to_src > \
--backend postgres
Example 3: Usage and output
> review_file src/memory/remap.c
Vulnerable source code:
// Remap memory addresses from one region to another
for ( uint32_t * ptr = start ; ptr < end ; ptr ++ ) {
uint32_t value = * ptr ;
if ( value >= OLD_REGION_BASE && value < OLD_REGION_BASE + REGION_SIZE ) {
value = value - OLD_REGION_BASE + NEW_REGION_BASE ;
}
}
Example output:
File: src/memory/remap.c
Identified issue 1: Address Remapping Loop Does Not Update Memory
Snippet:
for (uint32_t * ptr = start ; ptr < end ; ptr++) {
uint32_t value = * ptr ;
if...
Why: In the remap_address_table function, the code is intended to adjust address references from an old memory region to a new one. However, the updated value stored in the local variable ' value ' is never written back into memory at the pointer location ( * ptr). This means the address entries remain unchanged, which can lead to unintended behavior if the system relies on those values being relocated correctly.
Mitigation: Update the loop so that after computing the new address, the value is written back. For example:
for (uint32_t * ptr = start ; ptr < end ; ptr++) {
uint32_t value = * ptr ;
if (value > = OLD_REGION_BASE && value < OLD_REGION_BASE + REGION_SIZE) {
value = (( value - OLD_REGION_BASE) + NEW_REGION_BASE);
*ptr = value;
}
}
This ensures that each entry is properly updated to point to the relocated memory region.
Confidence: 1 . 0
Example 4: Run a full security review (non-interactive)
metis --non-interactive --comma

[truncated]
