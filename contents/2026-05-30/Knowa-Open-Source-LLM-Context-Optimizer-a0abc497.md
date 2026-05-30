---
source: "https://github.com/zzorphcreator/knowa"
hn_url: "https://news.ycombinator.com/item?id=48328512"
title: "Knowa – Open-Source LLM Context Optimizer"
article_title: "GitHub - zzorphcreator/knowa: LLM Context Optimizer · GitHub"
author: "zzorphcreator"
captured_at: "2026-05-30T11:39:02Z"
capture_tool: "hn-digest"
hn_id: 48328512
score: 2
comments: 0
posted_at: "2026-05-29T20:07:49Z"
tags:
  - hacker-news
  - translated
---

# Knowa – Open-Source LLM Context Optimizer

- HN: [48328512](https://news.ycombinator.com/item?id=48328512)
- Source: [github.com](https://github.com/zzorphcreator/knowa)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T20:07:49Z

## Translation

タイトル: Knowa – オープンソース LLM コンテキスト オプティマイザー
記事のタイトル: GitHub - zzorphcreator/知識: LLM Context Optimizer · GitHub
説明: LLM コンテキスト オプティマイザー。 GitHub でアカウントを作成して、zzorphcreator/knowa の開発に貢献してください。

記事本文:
GitHub - zzorphcreator/know: LLM コンテキスト オプティマイザー · GitHub
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
ゾルフクリエイター
/
知ってる
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスターブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github .github knowa knowa 移行 移行スクリプト スクリプトテスト テスト .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md cli.py cli.py config.py config.py docker-compose.yml docker-compose.yml knowa.gif knowa.gif pyproject.toml pyproject.toml pytest.ini pytest.iniquestions.sample.jsonquestions.sample.json要件.txt要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
問題: LLM トークンのコストは急速に増加します
AI を活用したアプリを構築するための素朴なアプローチは、ドキュメントをプロンプトにロードし、
LLM に理解させてください。デモでは動作します。生産中に壊れてしまいます。
1,000 ページのナレッジ ベースは約 100 万トークンに相当します。現在の API 価格では、送信
すべてのリクエストでクエリごとにドルがかかります。 1 日あたり 10,000 クエリの場合、これは数十に相当します。
月額数千ドル — 質問の内容とは 95% 無関係な文脈です。
AI の使用がチームや製品にわたって拡大するにつれて、これが主要なコストラインになります。
Knowa の中心的な仕事は、これを解決することです。ドキュメントをベクトル チャンク、フルテキストとして 1 回インデックス付けします。
ページと名前付きエンティティ グラフ — その後、質問ごとに少数のチャンクのみが抽出されます
実際に関連するものは、通常、大規模なコーパスから 1,000 ～ 3,000 個のトークンです。規模的には、
回答の品質を損なうことなく、LLM に送信されるコンテキストが 90 ～ 99% 削減されます。すべてのクエリ応答
実稼働環境でこれを追跡できるように、測定されたトークン節約量の数値が含まれています。貯蓄というのは、
コーパス サイズに依存します — トークンの節約についてを参照してください。
さまざまなスケールで何が期待されるか。
これは今でも重要ですが、今後さらに重要になるでしょう

AI の使用はプロトタイプから製品、そして
全社的なインフラストラクチャ。精密な検索はあれば便利というわけではありません - それが違いです
拡張性のある AI 機能と予算を超えた AI 機能の間で。
Knowa は、ハイブリッド検索ライブラリおよびナレッジ ベース サーバーです。ローカルからドキュメントを取り込みます
ディレクトリ、Notion、Confluence、および接続するカスタム ソース - それらをベクターとして保存
PostgreSQL のチャンクとフルテキスト ページ、および PostgreSQL の名前付きエンティティ ナレッジ グラフとして
(デフォルト)、Neo4j、または Kuzu。質問ごとに、最も関連性の高いコンテキストのみを取得します。
3 つの表現すべてにわたって、LLM または AI パイプラインに挿入できる状態で返されます。
あなたが選びます。
ナレッジ グラフはインデックス作成中に構築され、検索に構造的な次元を与えます。
純粋なベクトル検索のミス — 人、製品、組織、コンセプトを世界中で結び付ける
ドキュメント全体のコーパスなので、エンティティ中心の質問 (「どのチームが X に取り組んでいますか?」、
「どのページに Y と Z の両方が記載されていますか?」) 正確で的を絞った回答が得られます。
グラフにどのように表示されるかは、必要なカバレッジの範囲によって異なります。
spaCy NER (デフォルト) — API コストゼロでインデックス作成時にローカルで実行されます。認識します
標準的なエンティティ タイプ (人、組織、場所、日付など)。交換できます
精度を向上させたり、特定のドメインをターゲットにしたモデル: en_core_web_trf (トランスフォーマーベース、
曖昧な名前の方が良い)、科学または生物医学テキストの科学的モデル、またはその他の
spaCy対応モデル。
LLM エンティティ強化 (オプトイン) — OpenAI 互換の任意の機能を使用した追加的な 2 番目のパス
モデル (gpt-4o-mini、Qwen3、Kimi2 など)。 LLM はドメイン固有のエンティティを抽出します
一般的な NER モデルには欠けているもの: 製品の機能、内部コード名、技術標準、
抽象的な概念 - 組織内で意味のあるものすべて

言語。同時実行
大規模なインデックス作成を高速に維持するためにページ間でインデックスを作成し、1,000 ページあたり約 0.30 ドルのコストがかかります。
gpt-4o-mini。
ドッカー構成 -d
これにより、 pgvector がプリインストールされた PostgreSQL 16 インスタンスが起動され、 localhost:5432 に公開されます。
データは再起動後も Docker ボリュームに保持されます。
3. 仮想環境の作成
Python -m venv .venv
ソース .venv/bin/activate # Windows: .venv\Scripts\activate
または - miniconda がインストールされている場合
conda create -n knowa python=3.12
conda が知っていることをアクティブ化します
4. インストール
pip install -r 要件.txt
python -m spacy ダウンロード en_core_web_sm
pip install -e 。 # `knowa` CLI コマンドを登録します
5. 設定する
cp .env.example .env
.env を編集し、必要な値を入力します。
python3 -c " シークレットをインポート; print(secrets.token_urlsafe(32)) "
6. ドキュメントのインデックスを作成する
knowa インデックス /path/to/docs --name " My Docs "
移行は最初のコマンドで自動的に実行されます。サポートされているファイルの種類: .md 、 .txt 、 .pdf 、 .docx
知っているチャット「何を知りたいですか?」
8. 管理UIを実行する
uvicorn knowa.api.main:app --reload --port 8000
http://localhost:8000/admin/ui を開く — ブラウザーからソースを検索、参照し、リビルドをトリガーします。
Python ライブラリとして — Knowa をアプリに直接埋め込みます。 LLM コールの所有者はあなたです。ノワ
インデックス作成、取得、およびコンテキストの書式設定を処理します。 Anthropic、OpenAI、Gemini で動作します。
ローカル モデル、LangChain、LlamaIndex、またはその他のフレームワーク。
REST API として — FastAPI サーバーを実行し、/query を押して完全な回答を取得します。
あらゆる言語またはツールからの引用。Python は必要ありません。
CLI ツールとして - knowa コマンドを使用してディレクトリのインデックスを作成し、知識とチャットします。
ターミナルから直接ベースを利用できるため、サーバーやコードは必要ありません。
大規模な場合は最大 90 ～ 99% のトークン削減 — 3 パスのハイブリッド取得 (ベクトル検索、
全文検索

、プロパティ グラフ）関連するものだけを外科的に抽出します。すべてのクエリ
トークンの節約量を測定したレポートを作成できるため、本番環境の効率を追跡できます。参照
現実的な期待に基づいたトークンの節約について理解する
さまざまなコーパスサイズで。
階層的なチャンク化 — 正確なベクトルを実現するために、ドキュメントは小さな子チャンクに分割されます。
検索してから、より大きな親チャンクに拡張して完全なコンテキストを表示し、関連性を最大化します
周囲の情報を失わずに
プラグイン可能なエンベッダー — OpenAIEmbedder (デフォルト) または SentenceTransformerEmbedder
(完全にローカルで、クエリ時に API キーは必要ありません);追加するEmbedderプロトコルを実装します。
あなた自身の
複数のソース - Notion、Confluence Cloud、ローカル ディレクトリ ( .md 、 .txt 、 .pdf 、 .docx )
増分同期 — 前回の実行以降に変更されたファイルのみを再処理します
デフォルトでは、インデックス時の LLM 呼び出しはゼロです。spaCy はエンティティ抽出を処理します。 OpenAIなし
ナレッジ ベースのサイズに関係なく、インデックス作成中に費用がかかります。より豊富なエンティティ カバレッジを実現するには、
オプションの LLM エンリッチメント パスを有効にすることができます (「グラフ エンティティの抽出」を参照)
管理 UI — ソースごとの同期/再構築コントロール、クエリ インターフェイス、トークン節約の追跡、
インタラクティブなエンティティ グラフの視覚化
CLI — サーバーを実行せずに完全なインデックスとチャット管理を行う
各クエリの後に表示されるトークン節約量の数値は、インデックス付けされたコーパスの量を測定します。
Knowa は LLM への送信を避けました。
節約率 % = 1 − (このクエリによって取得されたトークン / DB 内のすべての親チャンク トークン)
これはコーパスサイズ相対の指標です。コーパスが小さい場合、検索で大部分をカバーできる可能性があります。
すべてのクエリでそれが実行され、節約は 0% に近くなります。それは予想されることであり、正しいことです。それはありません
バグまたは設定ミス。
取得の仕組み (およびサイズが重要な理由)
デフォルトでは、Knowa はベクトル類似度 ( TOP_K_CHUNKS=5 ) によって上位 5 つの子チャンクを取得します。
それから

それぞれを親チャンク (~2,048 トークン) に展開します。つまり、約 10,000 トークン
コーパスのサイズに関係なく、クエリごとのコンテキスト。
節約できるのは、インデックス付けされたコンテンツの合計がその期間を大幅に超えた場合のみです。
デフォルト設定 — TOP_K_CHUNKS=5 、親チャンク サイズ ~2,048 トークン
一般的なドキュメント サイズ — Wiki ページ、Notion ページ、Confluence ページ: 500 ～ 3,000 ワード
→ それぞれ 1 ～ 4 個の親チャンク。短い README 形式のファイル (500 ワード未満) では、単一の
チャンク;長い PDF または DOCX ファイル (5,000 ワード以上) では、それぞれ 5 ～ 15 のチャンクが生成される場合があります。
ファイル タイプが混在するとブレークポイントが移動します。10 個の長い PDF のコーパスが 50 以上の PDF のように動作する可能性があります。
合計チャンク数の点で短いマークダウン ファイル。
コーパスサイズごとに期待される節約量
コーパス
約親チャンク
期待される節約額
5 つの短いドキュメント
5～10
0 ～ 10% — 検索はコーパスの大部分をカバーします
20 ～ 30 ページの充実したページ
30～50
40～60%
100ページ以上
100～200
70～85%
500ページ以上
500以上
85 ～ 95%
大規模な Wiki (1,000 ページ以上)
1,000以上
90～99%
0% 節約の実際の意味
節約率が 0% であっても、取得が失敗したり、Knowa が役に立たなかったりするわけではありません。
これは、コーパスが十分に小さいため、検索ウィンドウが本質的にすべてをカバーすることを意味します。
— それは正しい動作です。回答の質に関する利点 (適切なページを見つけて、
一貫した答えを合成すること）は、あらゆるコーパス サイズに存在します。
節約バッジは、ナレッジ ベースが確立されれば、有用な生産監視シグナルになります。
取得が真に選択的になるほど十分に大きくなります。通常は 20 ～ 30 個の相当量です
書類以上。
大規模なコーパスがあるにもかかわらず節約額が低い場合は、TOP_K_CHUNKS を削減することを検討してください。
あなたの .env 。デフォルトの 5 は保守的です。およそ半分の 3 つにまで下がります
コンテキストを理解して節約量を増やしますが、広範な質問に対する再現率はわずかに低下します。

CLI はデータベースに直接接続します。サーバーは必要ありません。
# ディレクトリにインデックスを付ける — デフォルトでは増分、最初の実行時に完全に再構築
インデックス /path/to/docs を知っていますか
knowaindex /path/to/docs --name " Engineering Docs " # 分かりやすいラベルを付ける
knowa Index /path/to/docs --full # 完全な再構築を強制する
knowa Index /path/to/docs --workers 4 # 並列インデックス作成 (4 スレッド)
# ページ/チャンク数とラベルを含むすべてのインデックス付きソースを表示する
知っているリスト
# インデックスとチャットする
知っているチャット # インタラクティブ REPL
知っているチャット「返金ポリシーは何ですか?」 # 単発
knowa chat --source " Engineering Docs " # 1 つのソースに限定
knowa chat --debug # 各回答の前に取得パスを表示します
knowa chat --no-index /path/to/docs # インデックスをバイパスし、ドキュメントを直接読み取ります
# グラフ バックエンドを検査する (エンティティ/エッジ数)
デバッグを知っています
# インデックス付きとインデックスなしの質問のベンチマーク
知っているベンチの質問.json
# ソースのインデックスをクリアする
/path/to/docs をクリアしてください
knowa clear --source-id < 概念ワークスペース ID >
knowa clear --source-id company.atlassian.net/ENG
わかっている -- すべて
サポートされているファイルの種類: .md 、 .txt 、 .pdf 、 .docx
グラフの視覚化をテストするためにナレッジ ベースをすばやく入力するには、次の手順を実行します。
python scripts/fetch_sample_docs.py --out /tmp/knowa_sample_docs
knowa インデックス /tmp/knowa_sample_docs --name " Knowa サンプル ドキュメント "
これにより、最大 20 個のダウンロードが行われます

[切り捨てられた]

## Original Extract

LLM Context Optimizer. Contribute to zzorphcreator/knowa development by creating an account on GitHub.

GitHub - zzorphcreator/knowa: LLM Context Optimizer · GitHub
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
zzorphcreator
/
knowa
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github .github knowa knowa migrations migrations scripts scripts tests tests .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md cli.py cli.py config.py config.py docker-compose.yml docker-compose.yml knowa.gif knowa.gif pyproject.toml pyproject.toml pytest.ini pytest.ini questions.sample.json questions.sample.json requirements.txt requirements.txt View all files Repository files navigation
The problem: LLM token costs compound fast
The naive approach to building AI-powered apps is to load your documents into the prompt and
let the LLM figure it out. It works in demos. It breaks in production.
A 1,000-page knowledge base is roughly 1 million tokens. At current API pricing, sending
that on every request costs dollars per query. At 10,000 queries a day that is tens of
thousands of dollars a month — for context that is 95% irrelevant to the question being asked.
As AI usage scales across teams and products, this becomes the dominant cost line.
Knowa's core job is to solve this. It indexes your documents once — as vector chunks, full-text
pages, and a named-entity graph — then for each question extracts only the handful of chunks
that are actually relevant, typically 1,000–3,000 tokens out of a large corpus. At scale that is
a 90–99% reduction in context sent to the LLM, with no loss in answer quality. Every query response
includes a measured token savings figure so you can track this in production. Savings are
corpus-size dependent — see Understanding token savings for
what to expect at different scales.
This matters now and will matter much more as AI usage goes from prototype to product to
company-wide infrastructure. Precision retrieval is not a nice-to-have — it is the difference
between an AI feature that scales and one that breaks your budget.
Knowa is a hybrid retrieval library and knowledge base server. It ingests documents from local
directories, Notion, Confluence, and any custom source you connect — storing them as vector
chunks and full-text pages in PostgreSQL, and as a named-entity knowledge graph in PostgreSQL
(default), Neo4j, or Kuzu. For each question it retrieves only the most relevant context
across all three representations and returns it ready to inject into any LLM or AI pipeline
you choose.
The knowledge graph is built during indexing and gives retrieval a structural dimension that
pure vector search misses — connecting people, products, organisations, and concepts across
your entire document corpus so entity-centric questions ("what teams work on X?",
"which pages mention both Y and Z?") get precise, targeted answers.
How the graph is populated depends on how much coverage you need:
spaCy NER (default) — runs locally at indexing time with zero API cost. Recognises
standard entity types (people, organisations, locations, dates, and more). You can swap the
model to improve accuracy or target a specific domain: en_core_web_trf (transformer-based,
better on ambiguous names), scispaCy models for scientific or biomedical text, or any
spaCy-compatible model.
LLM entity enrichment (opt-in) — an additive second pass using any OpenAI-compatible
model (gpt-4o-mini, Qwen3, Kimi2, and others). The LLM extracts domain-specific entities
that a general NER model misses: product features, internal codenames, technical standards,
abstract concepts — anything with meaning in your organisation's language. Runs concurrently
across pages to keep indexing fast at scale, and costs roughly $0.30 per 1,000 pages with
gpt-4o-mini.
docker compose up -d
This starts a PostgreSQL 16 instance with pgvector pre-installed, exposed on localhost:5432 .
Data is persisted in a Docker volume across restarts.
3. Create a virtual environment
python -m venv .venv
source .venv/bin/activate # Windows: .venv\Scripts\activate
OR - if you have miniconda installed
conda create -n knowa python=3.12
conda activate knowa
4. Install
pip install -r requirements.txt
python -m spacy download en_core_web_sm
pip install -e . # registers the `knowa` CLI command
5. Configure
cp .env.example .env
Edit .env and fill in the required values:
python3 -c " import secrets; print(secrets.token_urlsafe(32)) "
6. Index your documents
knowa index /path/to/docs --name " My Docs "
Migrations run automatically on the first command. Supported file types: .md , .txt , .pdf , .docx
knowa chat " What would you like to know? "
8. Run the admin UI
uvicorn knowa.api.main:app --reload --port 8000
Open http://localhost:8000/admin/ui — search, browse sources, and trigger rebuilds from the browser.
As a Python library — embed Knowa directly in your app. You own the LLM call; Knowa
handles indexing, retrieval, and context formatting. Works with Anthropic, OpenAI, Gemini,
local models, LangChain, LlamaIndex, or any other framework.
As a REST API — run the FastAPI server and hit /query to get complete answers with
citations from any language or tool, no Python required.
As a CLI tool — use the knowa command to index directories and chat with your knowledge
base directly from the terminal, no server or code required.
Up to 90–99% token reduction at scale — three-path hybrid retrieval (vector search,
full-text search, property graph) surgically extracts only what is relevant; every query
reports measured token savings so you can track efficiency in production. See
Understanding token savings for realistic expectations
at different corpus sizes.
Hierarchical chunking — documents are split into small child chunks for precise vector
search, then expanded to larger parent chunks for full context — maximising relevance
without losing surrounding information
Pluggable embedders — OpenAIEmbedder (default) or SentenceTransformerEmbedder
(fully local, no API key needed at query time); implement the Embedder protocol to add
your own
Multiple sources — Notion, Confluence Cloud, and local directories ( .md , .txt , .pdf , .docx )
Incremental sync — only re-processes files changed since the last run
Zero LLM calls at index time by default — spaCy handles entity extraction; no OpenAI
spend during indexing regardless of knowledge base size. For richer entity coverage, an
optional LLM enrichment pass can be enabled (see Graph entity extraction )
Admin UI — per-source sync/rebuild controls, query interface, token savings tracking,
and interactive entity graph visualization
CLI — full index and chat management without running the server
The token savings figure shown after each query measures how much of your indexed corpus
Knowa avoided sending to the LLM:
savings % = 1 − (tokens retrieved by this query / all parent-chunk tokens in the DB)
This is a corpus-size-relative metric . With a small corpus, retrieval may cover most of
it on every query — and savings will be near 0%. That is expected and correct; it is not a
bug or a misconfiguration.
How retrieval works (and why size matters)
By default Knowa retrieves the top 5 child chunks by vector similarity ( TOP_K_CHUNKS=5 ),
then expands each to its parent chunk (~2,048 tokens). That means roughly 10,000 tokens of
context per query, regardless of corpus size.
Savings only accumulate once your total indexed content significantly exceeds that window.
Default settings — TOP_K_CHUNKS=5 , parent chunk size ~2,048 tokens
Typical document sizes — wiki pages, Notion pages, Confluence pages: 500–3,000 words
→ 1–4 parent chunks each. Short README-style files (under 500 words) may produce a single
chunk; long PDFs or DOCX files (5,000+ words) may produce 5–15 chunks each.
Mixed file types shift the breakpoints: a corpus of 10 long PDFs can behave like 50+
short markdown files in terms of total chunk count.
Expected savings by corpus size
Corpus
Approx. parent chunks
Expected savings
5 short docs
5–10
0–10% — retrieval covers most of the corpus
20–30 substantial pages
30–50
40–60%
100+ pages
100–200
70–85%
500+ pages
500+
85–95%
Large wiki (1,000+ pages)
1,000+
90–99%
What 0% savings actually means
A 0% savings reading does not mean retrieval is broken or that Knowa is not helping.
It means your corpus is small enough that the retrieval window covers essentially all of it
— which is the correct behaviour. The answer quality benefit (finding the right pages and
synthesising a coherent answer) is present at any corpus size.
The savings badge becomes a useful production monitoring signal once your knowledge base
grows large enough that retrieval is genuinely selective — typically 20–30 substantial
documents or more.
If you have a large corpus but still see low savings, consider reducing TOP_K_CHUNKS in
your .env . The default of 5 is conservative; dropping to 3 roughly halves retrieved
context and increases savings, at the cost of slightly lower recall on broad questions.
The CLI connects directly to the database — no server needed.
# Index a directory — incremental by default, full rebuild on first run
knowa index /path/to/docs
knowa index /path/to/docs --name " Engineering Docs " # attach a friendly label
knowa index /path/to/docs --full # force full rebuild
knowa index /path/to/docs --workers 4 # parallel indexing (4 threads)
# See all indexed sources with page/chunk counts and labels
knowa list
# Chat with the index
knowa chat # interactive REPL
knowa chat " What is our refund policy? " # single-shot
knowa chat --source " Engineering Docs " # scoped to one source
knowa chat --debug # show retrieval path before each answer
knowa chat --no-index /path/to/docs # bypass index, read docs directly
# Inspect graph backend (entity/edge counts)
knowa debug
# Benchmark questions with and without index
knowa bench questions.json
# Clear index for a source
knowa clear /path/to/docs
knowa clear --source-id < notion-workspace-id >
knowa clear --source-id company.atlassian.net/ENG
knowa clear --all
Supported file types: .md , .txt , .pdf , .docx
To quickly populate the knowledge base for testing the graph visualization:
python scripts/fetch_sample_docs.py --out /tmp/knowa_sample_docs
knowa index /tmp/knowa_sample_docs --name " Knowa Sample Docs "
This downloads ~20

[truncated]
