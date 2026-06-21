---
source: "https://github.com/GoogleCloudPlatform/knowledge-catalog/tree/main/okf"
hn_url: "https://news.ycombinator.com/item?id=48620068"
title: "Googles specification (and tooling) for the LLM wiki"
article_title: "knowledge-catalog/okf at main · GoogleCloudPlatform/knowledge-catalog · GitHub"
author: "Parkkeeper"
captured_at: "2026-06-21T16:32:06Z"
capture_tool: "hn-digest"
hn_id: 48620068
score: 2
comments: 1
posted_at: "2026-06-21T16:06:05Z"
tags:
  - hacker-news
  - translated
---

# Googles specification (and tooling) for the LLM wiki

- HN: [48620068](https://news.ycombinator.com/item?id=48620068)
- Source: [github.com](https://github.com/GoogleCloudPlatform/knowledge-catalog/tree/main/okf)
- Score: 2
- Comments: 1
- Posted: 2026-06-21T16:06:05Z

## Translation

タイトル: LLM wiki の Google の仕様 (およびツール)
記事のタイトル: メインのKnowledge-catalog/okf · GoogleCloudPlatform/knowledge-catalog · GitHub
説明: Google Cloud ナレッジ カタログのツールとサンプル。 GitHub でアカウントを作成して、GoogleCloudPlatform/ナレッジ カタログの開発に貢献します。

記事本文:
メインのKnowledge-catalog/okf · GoogleCloudPlatform/knowledge-catalog · GitHub
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
Googleクラウドプラットフォーム
/
知識カタログ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
さらなるオプション

ns ディレクトリのアクション
歴史 歴史メイン ブレッドクラム
コピーパスのトップフォルダーとファイル
.. バンドル バンドル サンプル サンプル src/reference_agent src/reference_agent テスト テスト .gitignore .gitignore LICENSE.md LICENSE.md README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml すべてのファイルを表示 README.md
概要オープンナレッジフォーマット(OKF)
📖 Open Knowledge Format v0.1 仕様を読む → SPEC.md
このリポジトリは主に Open Knowledge Format に関するものです
(OKF) 。
OKF は、ナレッジを表現するためのベンダー中立の汎用形式です。
YAML フロントマターを使用したプレーンなマークダウン ファイルとして。何にも縛られていない
特定のエージェント、フレームワーク、モデルプロバイダー、またはサービス提供システム。の
目標はシンプルです:
誰でも OKF を作成できます。人間が手作業でオーサリングし、エージェントはその上に構築されます。
任意のフレームワーク (Google ADK、LangChain、カスタム)、からパイプラインをエクスポート
既存のカタログ (Dataplex、Unity Catalog、Collibra など)、またはスクリプト
データベースを歩き回ります。
誰でも OKF (静的ファイル サーバー) を提供および利用できます。
ナレッジ管理 UI (Obsidian、Notion、MkDocs)、LLM ロード
ファイルをコンテキスト、検索インデックス、またはグラフ ビューアに追加する
このリポジトリにバンドルされています。
以下のエージェントは、次の 1 つの方法を示す概念実証です。
OKF バンドルを自動的に生成します。フォーマット自体は、
貢献。このエージェントとビジュアライザーはフォーマットを作成するために存在します。
生産と消費の両端で具体的なものになります。
OKF の実際を参照してください — これによって生成された 3 つのすぐに閲覧できるバンドル
エージェント、bundles/ にチェックインされました:
Bundles/ga4/ — GA4 電子商取引データセット
( viz.html )
バンドル/stackoverflow/ — スタック・オーバーフロー
パブリック データセット ( viz.html )
バンドル/crypto_bitcoin/ — ビットコイン
ブロック/トランザクション ( viz.html )
OKF は、カタログの知識を YAML を使用したプレーンなマークダウン ファイルとして表現します
フロントマター、ディレクトリ階層に編成

アーチー。この選択により、いくつかのロックが解除されます
サービス所有のメタデータ ストアから取得するのが難しいプロパティ:
人間とエージェントが読み取り可能。 SDK やクエリ言語は、
読者も内容も。エンジニアはコンセプトを理解できます。 LLM は取り込むことができます
それをそのまま文脈に取り入れます。
すぐにバージョン管理が可能。バンドルは git に存在します。引く
リクエスト、行ごとの差分、非難、レビューのワークフローは正常に機能します。
知識のキュレーションが通常のソフトウェア エンジニアリング活動になります。
ポータブルでロックインフリー。バンドルはディレクトリです。として出荷します
tarball、任意のリポジトリでホスト、任意のファイルシステムからマウント、または同期
ファイルを話すあらゆるシステム。独自の API はあなたとの間に立ちはだかることはありません
あなたのメタデータ。
構造化データと非構造化データを意図的に混合します。フロントマターを使用する
クエリ、フィルター、またはインデックスを作成するいくつかのフィールドに対して ( type 、
リソース、タグ、タイムスタンプ);散文にはマークダウン本文を使用し、
スキーマ、LLM と人間が実際に読み取るクエリの例。
最小限の意見で、自由に拡張可能。必要なものの小さなセット
キーは相互運用性を保証しますが、バンドルは任意の追加情報を運ぶことができます
フロントマターキーと任意のボディセクションを壊さずに
消費者。
既存のツールを使用して構成します。多くのナレッジ ツール — Notion、
Obsidian、MkDocs、Hugo、Jekyll — すでに Markdown と YAML について話しています
フロントマターなので、バンドルを参照、編集、レンダリングすることができます。
カスタムUI。
プログレッシブ開示機能が組み込まれています。自動生成されたindex.mdファイル
エージェントまたは人間が階層を一度に 1 レベルずつ移動できるようにする
バンドル全体をコンテキストにロードするのではなく。
単なるツリー状ではなく、グラフ状です。コンセプトは相互にリンクします。
通常のマークダウンリンクよりも豊かな関係を表現します。
ディレクトリレイアウトによって暗示される親/子。
最終的な影響は、参照エージェント、消費年齢です。

NTと人間
すでにコラボレーションしているのと同じ方法で、同じ成果物でコラボレーションする
ソースコードについて。
python3.13 -m venv .venv
.venv/bin/pip install --index-url https://pypi.org/simple/ -e .[dev]
資格情報
BigQuery: gcloud auth アプリケーションのデフォルト ログインと課金用のプロジェクト
( gcloud config set project <id> )。公開データセットは読み取り可能ですが、
呼び出し元のプロジェクトはクエリ バイトに対して課金されます。
Gemini: GEMINI_API_KEY (AI Studio) を設定するか、設定によって Vertex AI を使用します
GOOGLE_GENAI_USE_VERTEXAI=true 、 GOOGLE_CLOUD_PROJECT=<id> 、および
GOOGLE_CLOUD_LOCATION=<地域> 。
参照エージェントは 2 つのパスで実行されます。 BQ パスは 1 つの OKF を書き込みます
BigQuery メタデータのみを使用して、ソースが宣伝するコンセプトごとのドキュメント。
次に、Web パスは LLM を独自のクローラーとして実行します。
シード URL のリスト ( --web-seed または --web-seed-file で提供)、
fetch_url ツールを介してシードを取得し、どのアウトバウンドかを決定します。
リンクは信頼できるものであるかどうかに基づいてフォローする価値があります
既存の概念に関するドキュメント。フェッチするページごとに、
エージェントは、(a) 1 つ以上の既存の概念ドキュメントを充実させる、(b) ミントすることを選択します。
スタンドアロンのreferences/<slug> doc、または (c) スキップします。難しい
--web-max-pages の上限と同じドメインの許可されたホスト フィルター
( --web-allowed-host 経由で構成可能) はツール内で強制されます。
したがって、エージェントがオーバーランすることはありません。 Web パスをスキップするには、--no-web を使用します。
最小限の呼び出し — BigQuery データセットとバンドル出力をポイントします
ディレクトリ。 Web パスのシードは明示的です。それらを省略する（または渡す）
--no-web ) BQ のみを実行するには:
.venv/bin/python -mreference_agent エンリッチ \
--ソース bq \
--dataset <プロジェクト>.<データセット> \
--web-seed-file <path/to/seeds.txt> \
--out ./bundles/<名前>
--concept <type>/<name> を追加して、単一のコンセプトを反復処理します (例:
--コンセプト テーブル/イベント_ );再現可能。
それぞれ

サンプルは、レシピ (samples/<name>/) とシード URL を組み合わせます。
正確なエンリッチ コマンド) と生成されたバンドル (bundles/<name>/)
レシピが生成したもの。再現するレシピを開いてください。バンドルを開く
結果を直接参照します。
GA4 Google Merchandise Store — パブリック e コマース データセット、シード済み
正規の GA4 BigQuery エクスポート ドキュメント URL を使用します。
・レシピ
・バンドル
· viz.html
Stack Overflow — パブリック データセット (Stack Exchange データのミラー)
Dump)、コミュニティの正規スキーマ参照がシードされます。
横断的なドキュメント ページから複数の概念を強化します。
・レシピ
・バンドル
· viz.html
ビットコイン (暗号) — 公開データセット (ブロック、トランザクション、入力、
出力）ビットコイン-ETLパイプラインから。クロステーブル演習
散文における外部キー関係。
・レシピ
・バンドル
· viz.html
Visualize サブコマンドは、OKF バンドルを自己完結型としてレンダリングします。
インタラクティブ HTML ファイル — 1 つのファイル、バックエンドなし、へのインストールなし
見る側。最新のブラウザで開き、アーティファクトとして共有し、
静的ファイルサーバーでホストするか、バンドルの隣にコミットします (
このリポジトリはそうなります)。
ビューア自体は OKF の概念実証コンシューマであり、ミラーリングを行います。
参照エージェントが概念実証プロデューサーである方法。 OKF
バンドルは、マークダウンを読み取るものであれば何でも使用できます。これはただの
一つの形。
バンドル内のすべての概念の力指向のグラフ。
ノードのタイプ (データセット、テーブル、参照など) と方向ごとに色分けされます。
マークダウン ボディの各クロスリンクから描画されたエッジ。
選択したコンセプトの詳細パネルにその前部分が表示されます
(説明、リソースリンク、タグ) とそのレンダリングされたマークダウン本文 —
内部 […](/path/to/concept.md) リンクを再配線して移動します
パスをたどるのではなく、ビューア内で。
各 co の下にある「引用者」バックリンクのリスト

ncept (から計算
リンク グラフの逆)。
検索ボックス (タイトル、コンセプト ID、タグに一致)、タイプ
フィルター、および切り替え可能なグラフ レイアウト (cose/同心円/
幅優先/円/グリッド)。
.venv/bin/python -mreference_agent Visualize --bundle ./bundles/<名前>
これにより、 Bundles/<name>/viz.html が書き込まれます。フラグ:
出力を別の場所に書き込み、ヘッダーをオーバーライドする例:
.venv/bin/python -mreference_agent視覚化\
--バンドル ./bundles/crypto_bitcoin \
--out /tmp/btc.html \
--name "ビットコイン OKF"
作り方
HTML はバンドルを JSON BLOB として埋め込み、次を使用します。
グラフ用の Cytoscape.js と
ブラウザ内マークダウン レンダリング用にマークされている、
どちらも CDN からロードされます。ページからデータが流出することはありません。バンドルが解析される
生成時に一度実行され、ファイルにシリアル化されます。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Google Cloud Knowledge Catalog Tools and Samples. Contribute to GoogleCloudPlatform/knowledge-catalog development by creating an account on GitHub.

knowledge-catalog/okf at main · GoogleCloudPlatform/knowledge-catalog · GitHub
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
GoogleCloudPlatform
/
knowledge-catalog
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. bundles bundles samples samples src/ reference_agent src/ reference_agent tests tests .gitignore .gitignore LICENSE.md LICENSE.md README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml View all files README.md
Outline Open Knowledge Format (OKF)
📖 Read the Open Knowledge Format v0.1 specification → SPEC.md
This repository is primarily about the Open Knowledge Format
(OKF) .
OKF is a universal, vendor-neutral format for representing knowledge
as plain markdown files with YAML frontmatter. It is not tied to any
particular agent, framework, model provider, or serving system . The
goal is simple:
Anyone can produce OKF — humans authoring by hand, agents built on
any framework (Google ADK, LangChain, custom), export pipelines from
existing catalogs (Dataplex, Unity Catalog, Collibra, …), or scripts
walking a database.
Anyone can serve and consume OKF — a static file server, a
knowledge-management UI (Obsidian, Notion, MkDocs), an LLM loading
files into context, a search index, or a graph viewer like the one
bundled in this repo.
The agent below is a proof of concept demonstrating one way to
produce OKF bundles automatically. The format itself is the
contribution; this agent and the visualizer exist to make the format
tangible at both ends — production and consumption.
See OKF in practice — three ready-to-browse bundles produced by this
agent, checked into bundles/ :
bundles/ga4/ — GA4 e-commerce dataset
( viz.html )
bundles/stackoverflow/ — Stack Overflow
public dataset ( viz.html )
bundles/crypto_bitcoin/ — Bitcoin
blocks/transactions ( viz.html )
OKF represents catalog knowledge as plain markdown files with YAML
frontmatter, organized in a directory hierarchy. That choice unlocks a few
properties that are hard to get from a service-owned metadata store:
Human- and agent-readable. No SDK or query language stands between a
reader and the content. An engineer can cat a concept; an LLM can ingest
it verbatim into context.
Version-controllable out of the box. Bundles live in git. Pull
requests, line-by-line diffs, blame, and review workflows just work —
knowledge curation becomes a normal software-engineering activity.
Portable and lock-in free. A bundle is a directory. Ship it as a
tarball, host it in any repo, mount it from any filesystem, or sync it to
any system that speaks files. No proprietary API stands between you and
your metadata.
Mixes structured and unstructured data deliberately. Use frontmatter
for the few fields you want to query, filter, or index on ( type ,
resource , tags , timestamp ); use the markdown body for the prose,
schemas, and example queries that LLMs and humans actually read.
Minimally opinionated, freely extensible. A small set of required
keys ensures interoperability, but bundles can carry arbitrary extra
frontmatter keys and arbitrary body sections without breaking
consumers.
Composes with existing tooling. Many knowledge tools — Notion,
Obsidian, MkDocs, Hugo, Jekyll — already speak markdown plus YAML
frontmatter, so bundles can be browsed, edited, or rendered without
custom UI.
Progressive disclosure built in. Auto-generated index.md files
let an agent or human navigate the hierarchy one level at a time
instead of loading the entire bundle into context.
Graph-shaped, not just tree-shaped. Concepts link to each other via
normal markdown links, expressing relationships richer than the
parent/child implied by the directory layout.
The net effect is that reference agents, consumption agents, and humans
collaborate on the same artifacts in the same way they already collaborate
on source code.
python3.13 -m venv .venv
.venv/bin/pip install --index-url https://pypi.org/simple/ -e .[dev]
Credentials
BigQuery: gcloud auth application-default login plus a project for billing
( gcloud config set project <id> ). Public datasets are readable, but the
caller's project is billed for query bytes.
Gemini: set GEMINI_API_KEY (AI Studio) or use Vertex AI by setting
GOOGLE_GENAI_USE_VERTEXAI=true , GOOGLE_CLOUD_PROJECT=<id> , and
GOOGLE_CLOUD_LOCATION=<region> .
The reference agent runs in two passes. The BQ pass writes one OKF
doc per concept the source advertises, using BigQuery metadata alone.
The web pass then runs the LLM as its own crawler: it receives a
list of seed URLs (provided via --web-seed or --web-seed-file ),
fetches the seeds via the fetch_url tool, and decides which outbound
links are worth following based on whether they look like authoritative
documentation for the existing concepts. For each page it fetches, the
agent chooses to (a) enrich one or more existing concept docs, (b) mint
a standalone references/<slug> doc, or (c) skip. A hard
--web-max-pages cap and a same-domain allowed-hosts filter
(configurable via --web-allowed-host ) are enforced inside the tool,
so the agent cannot overrun. Use --no-web to skip the web pass.
Minimum invocation — point at a BigQuery dataset and a bundle output
directory. Seeds for the web pass are explicit; omit them (or pass
--no-web ) to run BQ-only:
.venv/bin/python -m reference_agent enrich \
--source bq \
--dataset <project>.<dataset> \
--web-seed-file <path/to/seeds.txt> \
--out ./bundles/<name>
Iterate on a single concept by adding --concept <type>/<name> (e.g.
--concept tables/events_ ); repeatable.
Each sample pairs a recipe ( samples/<name>/ , with the seed URLs and
exact enrich command) with the produced bundle ( bundles/<name>/ )
that the recipe generated. Open the recipe to reproduce; open the bundle
to browse the result directly.
GA4 Google Merchandise Store — public e-commerce dataset, seeded
with canonical GA4 BigQuery Export documentation URLs.
· recipe
· bundle
· viz.html
Stack Overflow — public dataset (mirror of the Stack Exchange Data
Dump), seeded with the community's canonical schema references.
Exercises multi-concept enrichment from cross-cutting docs pages.
· recipe
· bundle
· viz.html
Bitcoin (crypto) — public dataset (blocks, transactions, inputs,
outputs) from the bitcoin-etl pipeline. Exercises cross-table
foreign-key relationships in prose.
· recipe
· bundle
· viz.html
The visualize subcommand renders any OKF bundle as a self-contained
interactive HTML file — one file, no backend, no install on the
viewing side. Open it in any modern browser, share it as an artifact,
host it on a static file server, or commit it next to the bundle (as
this repo does).
The viewer is itself a proof-of-concept consumer of OKF, mirroring
the way the reference agent is a proof-of-concept producer . OKF
bundles can be consumed by anything that reads markdown; this is just
one shape.
A force-directed graph of every concept in the bundle, with
colored nodes by type (datasets, tables, references, …) and directed
edges drawn from each cross-link in the markdown bodies.
A detail panel for the selected concept showing its frontmatter
(description, resource link, tags) and its rendered markdown body —
with internal […](/path/to/concept.md) links rewired to navigate
within the viewer instead of following the path.
A "Cited by" backlinks list under each concept (computed from the
reverse of the link graph).
A search box (matches title, concept id, and tags), a type
filter , and switchable graph layouts (cose / concentric /
breadth-first / circle / grid).
.venv/bin/python -m reference_agent visualize --bundle ./bundles/<name>
That writes bundles/<name>/viz.html . Flags:
Example, writing the output somewhere else and overriding the header:
.venv/bin/python -m reference_agent visualize \
--bundle ./bundles/crypto_bitcoin \
--out /tmp/btc.html \
--name "Bitcoin OKF"
How it's built
The HTML embeds the bundle as a JSON blob and uses
Cytoscape.js for the graph and
marked for in-browser markdown rendering,
both loaded from a CDN. No data leaves the page; the bundle is parsed
once at generation time and serialized into the file.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
