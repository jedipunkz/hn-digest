---
source: "https://github.com/kelu124/Resume_RDF_update/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48478278"
title: "Show HN: Building a Resume ontology and tooling for LLMs CV management"
article_title: "GitHub - kelu124/Resume_RDF_update: Parse a CV into a structured Turtle RDF knowledge graph using the ResumeRDF ontology via a password-protected Streamlit web app or a standalone Python script. · GitHub"
author: "kelu124"
captured_at: "2026-06-10T16:20:05Z"
capture_tool: "hn-digest"
hn_id: 48478278
score: 1
comments: 0
posted_at: "2026-06-10T15:56:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Building a Resume ontology and tooling for LLMs CV management

- HN: [48478278](https://news.ycombinator.com/item?id=48478278)
- Source: [github.com](https://github.com/kelu124/Resume_RDF_update/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T15:56:32Z

## Translation

タイトル: HN の表示: LLM の履歴書管理のための履歴書オントロジーとツールの構築
記事タイトル: GitHub - kelu124/Resume_RDF_update: パスワードで保護された Streamlit Web アプリまたはスタンドアロン Python スクリプト経由で ResumeRDF オントロジーを使用して、CV を構造化された Turtle RDF ナレッジ グラフに解析します。 · GitHub
説明: パスワードで保護された Streamlit Web アプリまたはスタンドアロン Python スクリプトを介して ResumeRDF オントロジーを使用して、CV を構造化された Turtle RDF ナレッジ グラフに解析します。 - kelu124/Resume_RDF_update

記事本文:
GitHub - kelu124/Resume_RDF_update: パスワードで保護された Streamlit Web アプリまたはスタンドアロン Python スクリプトを介して ResumeRDF オントロジーを使用して、CV を構造化された Turtle RDF ナレッジ グラフに解析します。 · GitHub
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
ケル124
/
再開_RDF_更新
公共
N

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
66 コミット 66 コミット .streamlit .streamlit データ データの例 例resume_rdfresume_rdf shire shire .gitignore .gitignore LLM.md LLM.md README.md README.md app.py app.py application_example.ipynb application_example.ipynb application_example.py application_example.py claude_memory.md claude_memory.md claude_recommendations.md claude_recommendations.md process_cvs.sh process_cvs.sh pyproject.toml pyproject.toml要件.txt要件.txt要件_full.txt要件_full.txtすべてのファイルを表示リポジトリ ファイルのナビゲーション
ResumeRDF オントロジーを Python ライブラリ、CLI ツール、または Streamlit Web アプリとして使用して、CV を解析して構造化された Turtle RDF ナレッジ グラフにします。
このプロジェクトは、Anthropic Claude API を使用して CV (PDF またはプレーン テキスト) を Turtle RDF に変換します。コア ロジックは、resume_rdf Python パッケージ内にあり、Web アプリや CLI とは独立してインストールおよびインポートできます。
このグラフは、職歴やスキルだけでなく、プロジェクト レベルの詳細 (クライアント名、役割、活動、福利厚生、ドメイン タグに加えて、MOOC、認定資格、個人プロジェクト、出版物など) もキャプチャします。
LLM / エージェント ユーザー: すべての CLI、Python API、パイプライン オーケストレーション、出力構造、および一般的な注意点を網羅した簡潔なガイドについては、LLM.md を参照してください。
このツールは、履歴書用のスマートな読み取りアシスタントと考えてください。履歴書（任意の形式）を渡すと、文書全体が読み取られ、その人がどこで働いていたか、そこで何をしたか、どのプロジェクトを主導したか、どのようなスキルを使用したか、どのような結果をもたらしたかなど、あらゆる意味のある情報が抽​​出されます。次に、それらすべてを

コンピューターが検索、比較、推論できる構造化フォーマット。
実際の結果: 人間のみが読み取ることができる PDF の代わりに、生きたクエリ可能なプロファイルが得られます。 「再生可能エネルギープロジェクトに取り組んだ人は誰ですか?」などの質問をすることができます。または「当社の従業員の中で 10 人以上のチームを率いた人は誰ですか?」一度に数十、数百の履歴書から正確な回答を即座に得ることができます。このツールはまた、ギャップ (「この役割には終了日がありません」) にフラグを立て、シンプルなチャット インターフェイスを通じてギャップを埋めることができ、必要なときにクリーンでフォーマットされた履歴書をエクスポートして戻すこともできます。
Web アプリを使用するためにコーディングは必要ありません。共有を選択しない限り、データが環境から離れることはありません。
コンサルティング会社のリソース マネージャー向け
コンサルティング会社の死活は、人材配置の決定の質とスピードによって決まります。このツールは、そのプレッシャーを念頭に置いて構築されました。
これにより解決される問題: コンサルタントの履歴書は通常、Word 文書または PDF であり、構造化されておらず、フォーマットに一貫性がなく、大規模なクエリは不可能です。入札にふさわしい人を見つけるには、メールでやり取りしたり、記憶に頼ったり、常に最新でないスプレッドシートを維持したりする必要があります。
すべてのコンサルタントの構造化された検索可能なプロフィール (スキル、クライアント、セクター、プロジェクトの役割、認定資格、成果など) はすべて既存の履歴書から自動的に抽出されます。
クロスポートフォリオ検索 — 単一のドキュメントを読むことなく、特定のクライアント、セクター、またはテクノロジーに関連する経験を持つ人を即座に特定します。
自動ギャップ検出 — このツールは、不足している情報 (プロジェクトの説明なし、終了日なし、成果の明示なし) にフラグを立て、会話型インターフェイスを通じてこれらのギャップを埋めることができます。
履歴書の統合 — コンサルタントがさまざまな対象者向けに複数の履歴書バージョンを持っている場合、ツールは情報を失うことなくそれらを 1 つの完全なプロファイルに統合します。

メーション。
一貫したエクスポート — あらゆるプロファイルからクリーンで標準化されたマークダウン CV を生成し、入札またはフレームワークの送信に合わせて調整できます。
その結果、スキル データベースが既存のドキュメントから自動的に構築され、コンサルタントが履歴書を更新しても常に最新の状態に保たれ、BD チームとスタッフィング チームにベンチ全体にわたる真の検索機能が提供されます。
resume_rdf/ ← インポート可能な Python ライブラリ
§── __init__.py パブリック API — すべてのキーシンボルを再エクスポートします
§── api.py generated_graph_from_file / _from_bytes
§──cache.py ファイルベースの SHA-256 応答キャッシュ
§── cli.py cv-to-rdf エントリポイント
§── data.py データセットのダウンロード + 反復ヘルパー
§──export.py ttl_to_markdown / cv-to-md CLI
§── merge.py consolidate_ttls / cv-merge CLI
§── ontology.py SYSTEM_PROMPT + 名前空間定数
§── parsing.py Turtle ヘルパー (count_triples、extract_person_name …)
§── qa.py Audit_ experience、update_field / cv-audit、cv-update CLI
§── reconcile.py エンティティ調整 / cv-reconcile CLI
━── viz.py Visualize_cv / cv-graph CLI
app.py Streamlit Web アプリ (薄いラッパー)
application_example.py スタンドアロンのエンドツーエンド デモ スクリプト
pyproject.toml pip インストール可能なパッケージ定義
data/master_resumes.jsonl (1 866 レコード、MIT)
インストール
pip インストール 。
# オプションの追加機能あり:
pip install " .[app] " # streamlit を追加
pip install " .[validate] " # Turtle 検証用の rdflib を追加します
pip install " .[dataset] " # データセット +huggingface_hub を追加します
pip install " .[all] " # すべて
ソース (開発) から
git clone https://github.com/kelu124/Resume_RDF_update.git
cd 再開_RDF_update
pip install -e " .[all] "
図書館の利用状況
fromresume_rdf importgenerate_graph_from_file 、generate_graph_from_bytes
# ファイルパスから(CLI/バッチ使用)
カメ、

使用法 = generated_graph_from_file (
"my_cv.pdf" ,
api_key = "sk-ant-..." , # または ANTHROPIC_API_KEY を設定します
extra_context = "エネルギー部門、英語ラベル。" 、
モデル = "claude-sonnet-4-6" , # デフォルト
)
print ( Turtle ) # 有効な Turtle RDF
# バイトから (Web / インメモリ使用)
open ( "my_cv.pdf" , "rb" ) を f として使用:
タートル、使用法 =generate_graph_from_bytes (
f 。 read (), "my_cv.pdf" , api_key = "sk-ant-..."
)
print ( f"~ { 使用法 [ 'output_tokens' ]:, } 出力トークン" )
公共事業
resume_rdf import count_triples 、 extract_person_name 、 validate_turtle から
print ( count_triples ( タートル )) # ~420
print ( extract_person_name ( Turtle )) # "ジェーン・スミス"
validate_turtle ( Turtle ) # True (rdflib が必要)
オントロジー定数
resume_rdf import から NAMESPACES 、 SYSTEM_PROMPT
# NAMESPACES は辞書です: {"cv": "http://...", "cvx": "http://...", ...}
NAMESPACES の prefix 、uri の場合。項目 ():
print ( f"@prefix { prefix } : < { uri } > ." )
キャッシング
応答は、(システム プロンプト、
モデル、ユーザーコンテンツ）トリプル。キャッシュを無効にする .json ファイルを削除します。
特定の入力を使用するか、LLM_CACHE_DIR を設定してキャッシュ ディレクトリをオーバーライドします。
pip install 後。次の 7 つの CLI コマンドが使用可能です。
履歴書から rdf へ my_cv.pdf # → my_cv.ttl
cv-to-rdf my_cv.pdf --outputgraph.ttl
cv-to-rdf my_cv.pdf --context 「エネルギー部門、出力は英語。」
cv-to-rdf my_cv.pdf --model claude-opus-4-7 --max-tokens 80000
cv-to-rdf my_cv.pdf --validate --quiet
cv-to-rdf my_cv.pdf --overwrite # 既存の .ttl を置き換えます
デフォルトでは、出力 .ttl がすでに存在する場合、cv-to-rdf は API を呼び出さずに終了します。 --overwrite を渡して置き換えます。
API キーは、環境変数 (推奨) またはフラグを介して渡すことができます。
import ANTHROPIC_API_KEY= " sk-ant-... "
# または
cv-to-rdf my_cv.pdf --api-key sk-ant-...
ストリームリ

ウェブアプリ
streamlit で app.py を実行
http://localhost:8501 で開きます。 4 ステップのパイプライン:
mkdir -p .streamlit
cp Secrets.toml.example .streamlit/secrets.toml
.streamlit/secrets.toml を編集します。
[アプリ]
パスワード = " アプリのパスワード "
[人類]
api_key = " sk-ant-... "
オントロジー
このグラフは、確立された語彙と軽量のカスタム拡張機能を組み合わせたものです。
RDFクラス
経由でリンクされました
説明
foaf:人物 + cv:CV
—
ID と CV ルート
CV:職歴
CV:職歴あり
雇用形態
CV:会社
履歴書: 勤務先
雇用主と顧客
CVX:プロジェクト
cvx:hasプロジェクト
クライアントとの関わり
CV：スキル
CV:スキルあり
技術的および専門的なスキル
履歴書:教育
履歴書:学歴あり
正式な学位
CVX:MOOC
cvx:hasMOOC
オンラインコース
CVX:トレーニング
cvx:hasトレーニング
ワークショップ、認定資格
cvx:パーソナルプロジェクト
cvx:hasPersonalProject
オープンソース、ハードウェア、コミュニティ プロジェクト
bibo:学術記事 / …
cvx:hasPublication
論文、レポート、特許
プロジェクトには、そのエンゲージメントで使用される cv:Skill ノードへの cvx:usesSkill リンクも含まれます。
Turtle CV を視覚的なナレッジ グラフとしてレンダリングします。
pyvis (HTML) または networkx + matplotlib (PNG/SVG) が必要です。
pip install "resume-rdf[viz]" 。
cv-graph my_cv.ttl # → my_cv.html (対話型、デフォルト)
cv-graph my_cv.ttl --outputgraph.png # → 静的 PNG
cv-graph my_cv.ttl --outputgraph.svg # → 静的 SVG
Python API
からresume_rdfインポートvisual_cv
out = Visualize_cv ( "my_cv.ttl" ) # → my_cv.html
out = Visualize_cv ( "my_cv.ttl" , "graph.png" ) # →graph.png
print ( f"保存済み: { out } " )
グラフには、個人 → 雇用主 → プロジェクト → スキル、教育が表示されます。
トレーニング/MOOC、個人プロジェクト、およびその人から発信される出版物
ノード。任意のノードにマウスを置くと、完全なツールチップ (名前、役割、日付、説明) が表示されます。
Turtle RDF CV をクリーンで人間が読める形式に変換します

マークダウン。
rdflib が必要です: pip install "resume-rdf[export]" 。
cv-to-md my_cv.ttl # 標準出力に出力
cv-to-md my_cv.ttl --output cv.md # ファイルに書き込みます
Python API
からresume_rdfインポートttl_to_markdown
md = ttl_to_markdown ( "my_cv.ttl" )
印刷 (MD)
pathlibインポートパスから
パス (「my_cv.md」)。 write_text ( ttl_to_markdown ( "my_cv.ttl" ))
出力は、手書きのコンサルタント履歴書の構造を反映しています。
個人ヘッダー、コアスキル、ネストされたプロジェクトでの専門的経験
(活動 + 成果 + 使用されたスキル)、教育、認定およびトレーニング、
個人的なプロジェクトや出版物。内容のない部分は省略しています。
履歴書の品質監査と現場の最新情報
グラフを生成した後、ギャップを監査して埋めることができます。
プログラム的に。
rdflib が必要です: pip install "resume-rdf[qa]" 。
cv-audit — 欠落しているフィールドを検索する
cv-audit my_cv.ttl
出力例:
3 つの質問が見つかりました:
[wh_acme_2019] 終了日
Acme Corp を辞めたのはいつですか、それとも現在の役割 (YYYY-MM-DD または「現在」) ですか?
[proj_smartgrid_2022] の特典の配信
プロジェクト「スマート グリッド アナリティクス」はどのような成果や利益をもたらしましたか?
[proj_smartgrid_2022] スキルの使用
プロジェクト「スマート グリッド アナリティクス」ではどのスキルが使用されましたか? (スキルのスラッグまたは名前を入力してください)
機械可読出力用に --json を追加します。
cv-audit my_cv.ttl --json
# → [{"slug": "wh_acme_2019", "field": "end

[切り捨てられた]

## Original Extract

Parse a CV into a structured Turtle RDF knowledge graph using the ResumeRDF ontology via a password-protected Streamlit web app or a standalone Python script. - kelu124/Resume_RDF_update

GitHub - kelu124/Resume_RDF_update: Parse a CV into a structured Turtle RDF knowledge graph using the ResumeRDF ontology via a password-protected Streamlit web app or a standalone Python script. · GitHub
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
kelu124
/
Resume_RDF_update
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
66 Commits 66 Commits .streamlit .streamlit data data examples examples resume_rdf resume_rdf shire shire .gitignore .gitignore LLM.md LLM.md README.md README.md app.py app.py application_example.ipynb application_example.ipynb application_example.py application_example.py claude_memory.md claude_memory.md claude_recommendations.md claude_recommendations.md process_cvs.sh process_cvs.sh pyproject.toml pyproject.toml requirements.txt requirements.txt requirements_full.txt requirements_full.txt View all files Repository files navigation
Parse a CV into a structured Turtle RDF knowledge graph using the ResumeRDF ontology — as a Python library , a CLI tool, or a Streamlit web app.
This project converts a CV (PDF or plain text) into Turtle RDF using the Anthropic Claude API . The core logic lives in the resume_rdf Python package, which can be installed and imported independently of the web app or CLI.
The graph captures not just employment history and skills, but also project-level detail : client names, roles, activities, benefits, and domain tags — plus MOOCs, certifications, personal projects, and publications.
LLM / agent users: see LLM.md for a concise guide covering all CLIs, the Python API, pipeline orchestration, output structure, and common gotchas.
Think of this tool as a smart reading assistant for CVs. You give it a CV — any format — and it reads through the whole document and extracts every meaningful piece of information: where the person worked, what they did there, which projects they led, what skills they used, and what results they delivered. It then stores all of that in a structured format that a computer can search, compare, and reason over.
The practical result: instead of a PDF that only a human can read, you get a living, queryable profile. You can ask questions like "who has worked on renewable energy projects?" or "which of our people have led teams of more than ten?" and get instant, accurate answers — across dozens or hundreds of CVs at once. The tool also flags gaps ("this role has no end date"), lets you fill them in through a simple chat interface, and can export a clean, formatted CV back out when you need it.
No coding required to use the web app. No data leaves your environment unless you choose to share it.
For resource managers in consulting firms
Consulting firms live and die by the quality and speed of their staffing decisions. This tool was built with that pressure in mind.
The problem it solves: consultant CVs are usually Word documents or PDFs — unstructured, inconsistently formatted, and impossible to query at scale. Finding the right person for a bid means emailing around, relying on memory, or maintaining a spreadsheet that's always out of date.
A structured, searchable profile for every consultant — skills, clients, sectors, project roles, certifications, and outcomes, all extracted automatically from existing CVs.
Cross-portfolio search — instantly identify who has relevant experience for a specific client, sector, or technology, without reading a single document.
Automatic gap detection — the tool flags missing information (no project description, no end date, no stated outcomes) and lets you fill those gaps through a conversational interface.
CV consolidation — when a consultant has multiple CV versions for different audiences, the tool merges them into one complete profile without losing any information.
Consistent exports — generate clean, standardised Markdown CVs from any profile, ready to tailor for a bid or framework submission.
The result is a skills database that builds itself from documents you already have, stays up to date as consultants update their CVs, and gives your BD and staffing teams a genuine search capability across the whole bench.
resume_rdf/ ← importable Python library
├── __init__.py public API — re-exports all key symbols
├── api.py generate_graph_from_file / _from_bytes
├── cache.py file-based SHA-256 response cache
├── cli.py cv-to-rdf entry-point
├── data.py dataset download + iteration helpers
├── export.py ttl_to_markdown / cv-to-md CLI
├── merge.py consolidate_ttls / cv-merge CLI
├── ontology.py SYSTEM_PROMPT + namespace constants
├── parsing.py Turtle helpers (count_triples, extract_person_name …)
├── qa.py audit_experience, update_field / cv-audit, cv-update CLIs
├── reconcile.py entity reconciliation / cv-reconcile CLI
└── viz.py visualize_cv / cv-graph CLI
app.py Streamlit web app (thin wrapper)
application_example.py standalone end-to-end demo script
pyproject.toml pip-installable package definition
data/ master_resumes.jsonl (1 866 records, MIT)
Installation
pip install .
# With optional extras:
pip install " .[app] " # adds streamlit
pip install " .[validate] " # adds rdflib for Turtle validation
pip install " .[dataset] " # adds datasets + huggingface_hub
pip install " .[all] " # everything
From source (dev)
git clone https://github.com/kelu124/Resume_RDF_update.git
cd Resume_RDF_update
pip install -e " .[all] "
Library usage
from resume_rdf import generate_graph_from_file , generate_graph_from_bytes
# From a file path (CLI / batch use)
turtle , usage = generate_graph_from_file (
"my_cv.pdf" ,
api_key = "sk-ant-..." , # or set ANTHROPIC_API_KEY
extra_context = "Energy sector, English labels." ,
model = "claude-sonnet-4-6" , # default
)
print ( turtle ) # valid Turtle RDF
# From bytes (web / in-memory use)
with open ( "my_cv.pdf" , "rb" ) as f :
turtle , usage = generate_graph_from_bytes (
f . read (), "my_cv.pdf" , api_key = "sk-ant-..."
)
print ( f"~ { usage [ 'output_tokens' ]:, } output tokens" )
Utilities
from resume_rdf import count_triples , extract_person_name , validate_turtle
print ( count_triples ( turtle )) # ~420
print ( extract_person_name ( turtle )) # "Jane Smith"
validate_turtle ( turtle ) # True (requires rdflib)
Ontology constants
from resume_rdf import NAMESPACES , SYSTEM_PROMPT
# NAMESPACES is a dict: {"cv": "http://...", "cvx": "http://...", ...}
for prefix , uri in NAMESPACES . items ():
print ( f"@prefix { prefix } : < { uri } > ." )
Caching
Responses are cached in cache/<sha256>.json keyed on the (system prompt,
model, user content) triple. Delete any .json file to bust the cache for
a specific input, or set LLM_CACHE_DIR to override the cache directory.
After pip install . seven CLI commands are available:
cv-to-rdf my_cv.pdf # → my_cv.ttl
cv-to-rdf my_cv.pdf --output graph.ttl
cv-to-rdf my_cv.pdf --context " Energy sector, output in English. "
cv-to-rdf my_cv.pdf --model claude-opus-4-7 --max-tokens 80000
cv-to-rdf my_cv.pdf --validate --quiet
cv-to-rdf my_cv.pdf --overwrite # replace existing .ttl
By default, if the output .ttl already exists, cv-to-rdf exits without calling the API. Pass --overwrite to replace it.
The API key can be passed via environment variable (recommended) or flag:
export ANTHROPIC_API_KEY= " sk-ant-... "
# or
cv-to-rdf my_cv.pdf --api-key sk-ant-...
Streamlit web app
streamlit run app.py
Opens at http://localhost:8501 . Four-step pipeline:
mkdir -p .streamlit
cp secrets.toml.example .streamlit/secrets.toml
Edit .streamlit/secrets.toml :
[ app ]
password = " your-app-password "
[ anthropic ]
api_key = " sk-ant-... "
Ontology
The graph combines established vocabularies with a lightweight custom extension:
RDF Class
Linked via
Description
foaf:Person + cv:CV
—
Identity and CV root
cv:WorkHistory
cv:hasWorkHistory
Employment positions
cv:Company
cv:employedIn
Employers and clients
cvx:Project
cvx:hasProject
Client engagements
cv:Skill
cv:hasSkill
Technical and professional skills
cv:Education
cv:hasEducation
Formal degrees
cvx:MOOC
cvx:hasMOOC
Online courses
cvx:Training
cvx:hasTraining
Workshops, certifications
cvx:PersonalProject
cvx:hasPersonalProject
Open-source, hardware, community projects
bibo:AcademicArticle / …
cvx:hasPublication
Papers, reports, patents
Projects also carry cvx:usesSkill links to the cv:Skill nodes used on that engagement.
Render a Turtle CV as a visual knowledge graph.
Requires pyvis (HTML) or networkx + matplotlib (PNG/SVG):
pip install "resume-rdf[viz]" .
cv-graph my_cv.ttl # → my_cv.html (interactive, default)
cv-graph my_cv.ttl --output graph.png # → static PNG
cv-graph my_cv.ttl --output graph.svg # → static SVG
Python API
from resume_rdf import visualize_cv
out = visualize_cv ( "my_cv.ttl" ) # → my_cv.html
out = visualize_cv ( "my_cv.ttl" , "graph.png" ) # → graph.png
print ( f"Saved: { out } " )
The graph shows: Person → Employer → Project → Skill , plus Education,
Training/MOOCs, Personal Projects, and Publications radiating from the person
node. Hover over any node for the full tooltip (name, role, dates, description).
Convert a Turtle RDF CV back to clean, human-readable Markdown.
Requires rdflib : pip install "resume-rdf[export]" .
cv-to-md my_cv.ttl # print to stdout
cv-to-md my_cv.ttl --output cv.md # write to file
Python API
from resume_rdf import ttl_to_markdown
md = ttl_to_markdown ( "my_cv.ttl" )
print ( md )
from pathlib import Path
Path ( "my_cv.md" ). write_text ( ttl_to_markdown ( "my_cv.ttl" ))
The output mirrors the structure of a hand-written consultant CV:
person header, core skills, professional experience with nested projects
(activities + outcomes + skills used), education, certifications & training,
personal projects, and publications. Sections with no content are omitted.
CV quality audit & field update
After generating a graph you can audit it for gaps and fill them in
programmatically.
Requires rdflib : pip install "resume-rdf[qa]" .
cv-audit — find missing fields
cv-audit my_cv.ttl
Example output:
Found 3 question(s):
[wh_acme_2019] endDate
When did you leave Acme Corp, or is this your current role (YYYY-MM-DD or 'present')?
[proj_smartgrid_2022] benefitsDelivered
What outcomes or benefits did project 'Smart Grid Analytics' deliver?
[proj_smartgrid_2022] usesSkill
Which skills were used on project 'Smart Grid Analytics'? (provide skill slugs or names)
Add --json for machine-readable output:
cv-audit my_cv.ttl --json
# → [{"slug": "wh_acme_2019", "field": "end

[truncated]
