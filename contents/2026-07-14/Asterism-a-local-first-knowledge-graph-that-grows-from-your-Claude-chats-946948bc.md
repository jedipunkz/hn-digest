---
source: "https://github.com/biditdas18/asterism"
hn_url: "https://news.ycombinator.com/item?id=48911279"
title: "Asterism – a local-first knowledge graph that grows from your Claude chats"
article_title: "GitHub - biditdas18/asterism: Local-first personal knowledge graph with Hebbian TTL decay, rendered as a living constellation. · GitHub"
author: "biditdas18"
captured_at: "2026-07-14T19:02:47Z"
capture_tool: "hn-digest"
hn_id: 48911279
score: 1
comments: 0
posted_at: "2026-07-14T18:40:24Z"
tags:
  - hacker-news
  - translated
---

# Asterism – a local-first knowledge graph that grows from your Claude chats

- HN: [48911279](https://news.ycombinator.com/item?id=48911279)
- Source: [github.com](https://github.com/biditdas18/asterism)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T18:40:24Z

## Translation

タイトル: アステリズム – クロード チャットから成長するローカルファーストのナレッジ グラフ
記事のタイトル: GitHub - biditdas18/asterism: ヘビアン TTL 減衰を使用したローカルファーストの個人知識グラフ。生きた星座としてレンダリングされます。 · GitHub
説明: ヘビアン TTL 減衰を使用したローカルファーストの個人知識グラフ。生きた星座としてレンダリングされます。 - ビディダス18/アステリズム

記事本文:
GitHub - biditdas18/asterism: ヘビアン TTL 減衰を伴うローカルファーストの個人知識グラフ。生きた星座としてレンダリングされます。 · GitHub
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
ビディダス18
/
アステリズム
公共
通知
通知設定を変更するにはサインインする必要があります
追加

最後のナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット docs docs .env.example .env.example .gitignore .gitignore README.md README.md RESEARCH.md RESEARCH.md app.py app.py cli.py cli.py config.py config.py context.py context.py クローラ.py クローラ.py db.py db.py Decay_scheduler.pydecay_scheduler.pydemo_seed.pydemo_seed.pyextractor.pyextractor.pygraph.pygraph.pyllm.pyllm.pypyproject.tomlpyproject.tomlrender.pyrender.pyrequirements.txtrequirements.txtschema.sqlschema.sqlsetup.shsetup.shtest_foundation.pytest_foundation.py test_llm.py test_llm.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
脳のように考え、星座のように見えるローカルファーストの個人知識グラフ。
クロードとの会話はすべて痕跡を残します。アステリズムは、それらの痕跡を生きた星の地図にマッピングします。何かについて考えれば考えるほど、それはより明るく輝きます。それについて考えるのをやめれば、それは暗闇に消えていきます。
Hebbian 学習 — LLM が概念を横断するたびに、概念間のエッジが強化されます (横断ごとに重み += 0.2)。再考されない概念はセッションの露出時間を蓄積します。横断せずに 3 時間連続して暴露すると、それらは減衰してグラフから消えます。
あなたは中心ノードです - あなたのユーザー ノード (Bidit) は常に最大の明るさでコンスタレーションの中心に位置します。それは決して朽ちることはありません。
ローカル SQLite ストレージ — グラフ全体は ~/.asterism/asterism.db に存在します。クラウドも同期もアカウントもありません。
LLM コンテキスト挿入 — すべてのメッセージで、最も関連性の高い上位 N 個のノードとエッジが暗黙的なコンテキストとしてクロード プロンプトに挿入され、モデルが過去の思考を意識して応答できるようにします。
トリプル抽出 - 各交換は高速抽出によって処理されます

(エンティティ、関係、エンティティ) を 3 倍にしてグラフに書き込むモデル (ローカルの Ollama または Anthropic Haiku)。
git clone https://github.com/biditdas18/asterism
cd アステリズム
chmod +x setup.sh && ./setup.sh
次に:
asterism init の実行 — ガイド付きセットアップ ウィザード
クロード データをエクスポートします: claude.ai → 設定 → データのエクスポート → メールを確認 → zip をダウンロード
アステリズム クロール --source claude --path path/to/conversations.json
星座が自動的に開きます
それだけです。あなたの心は星座です。
グラフがマシンから離れることはありません。外部呼び出しは、Claude 応答用の Anthropic API と (オプションで) Haiku を利用したトリプル抽出のみです。 LLM は、ローカル グラフから明示的に挿入されたもののみを参照します。生のデータベースにはアクセスできません。 ~/.asterism/ を削除して、白紙の状態に戻します。テレメトリ、分析、アカウントはありません。
Asterism は、個人の AI 記憶システムにいくつかの新しい貢献を導入します。
ヘビアン知識グラフ — エッジは横断すると強化され、使用されなくなると減衰します
セッションベースの TTL — 減衰は実時間ではなく、アクティブなセッション時間のみをカウントします。
インデックスとしてのグラフ — グラフは、LLM フラット メモリ上で加重 B ツリー インデックスとして機能します。
4 つの生物学的力学 - 強化、枝刈り、経路の最適化、連鎖治癒
因果連鎖トポロジー — トピックの類似性だけでなく、アイデアがどのように発展したかによって会話が整理されました
正式な研究論文は準備中です。
詳細については、RESEARCH.md を参照してください。
レイヤー
技術
ストレージ
SQLite ( ~/.asterism/asterism.db )
グラフ
ネットワークX
視覚化
Vanilla JS 強制シミュレーション (依存関係なし)
LLM
Anthropic SDK — claude-sonnet-4-6
抽出
Ollama llama3.2:3b (ローカル) または Anthropic Haiku (クラウド)
UI
ストリームライト
CLI
クリック
著者
ヘビアン TTL 減衰を伴うローカルファーストの個人知識グラフ。生きた星座としてレンダリングされます。
T

読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
アステリズム v0.1.0 — 初期リリース
最新の
2026 年 6 月 26 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-first personal knowledge graph with Hebbian TTL decay, rendered as a living constellation. - biditdas18/asterism

GitHub - biditdas18/asterism: Local-first personal knowledge graph with Hebbian TTL decay, rendered as a living constellation. · GitHub
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
biditdas18
/
asterism
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits docs docs .env.example .env.example .gitignore .gitignore README.md README.md RESEARCH.md RESEARCH.md app.py app.py cli.py cli.py config.py config.py context.py context.py crawler.py crawler.py db.py db.py decay_scheduler.py decay_scheduler.py demo_seed.py demo_seed.py extractor.py extractor.py graph.py graph.py llm.py llm.py pyproject.toml pyproject.toml render.py render.py requirements.txt requirements.txt schema.sql schema.sql setup.sh setup.sh test_foundation.py test_foundation.py test_llm.py test_llm.py View all files Repository files navigation
A local-first personal knowledge graph that thinks like a brain and looks like a constellation.
Every conversation you have with Claude leaves a trace. Asterism maps those traces into a living star map — the more you think about something, the brighter it glows. Stop thinking about it, and it fades into the dark.
Hebbian learning — edges between concepts strengthen each time the LLM traverses them ( weight += 0.2 per traversal). Concepts that aren't revisited accumulate session exposure time; after 3 hours of uninterrupted exposure without traversal they decay and vanish from the graph.
You are the central node — your user node (Bidit) sits at the centre of the constellation at full brightness, always. It never decays.
Local SQLite storage — the entire graph lives in ~/.asterism/asterism.db . No cloud, no sync, no accounts.
LLM context injection — on every message, the top-N most relevant nodes and edges are injected into the Claude prompt as implicit context, letting the model answer with awareness of your past thinking.
Triple extraction — each exchange is processed by a fast extraction model (local Ollama or Anthropic Haiku) that pulls (entity, relationship, entity) triples and writes them to the graph.
git clone https://github.com/biditdas18/asterism
cd asterism
chmod +x setup.sh && ./setup.sh
Then:
Run asterism init — guided setup wizard
Export your Claude data: claude.ai → Settings → Export Data → check email → download zip
asterism crawl --source claude --path path/to/conversations.json
Your constellation opens automatically
That's it. Your mind as a constellation.
Your graph never leaves your machine. The only external calls are to the Anthropic API for Claude responses and (optionally) Haiku-powered triple extraction. The LLM only sees what you explicitly inject from your local graph — it has no access to the raw database. Delete ~/.asterism/ to get a clean slate. No telemetry, no analytics, no accounts.
Asterism introduces several novel contributions to personal AI memory systems:
Hebbian knowledge graph — edges strengthen on traversal, decay on disuse
Session-based TTL — decay counts only active session time, not wall-clock time
Graph-as-index — the graph acts as a weighted B-tree index over LLM flat memory
Four biological mechanics — strengthening, pruning, pathway optimization, chain healing
Causal chain topology — conversations organized by how ideas evolved, not just topic similarity
A formal research paper is in preparation.
See RESEARCH.md for full details.
Layer
Tech
Storage
SQLite ( ~/.asterism/asterism.db )
Graph
NetworkX
Visualization
Vanilla JS force simulation (zero dependencies)
LLM
Anthropic SDK — claude-sonnet-4-6
Extraction
Ollama llama3.2:3b (local) or Anthropic Haiku (cloud)
UI
Streamlit
CLI
Click
Author
Local-first personal knowledge graph with Hebbian TTL decay, rendered as a living constellation.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Asterism v0.1.0 — Initial Release
Latest
Jun 26, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
