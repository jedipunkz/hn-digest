---
source: "https://github.com/roandejager/Hillock"
hn_url: "https://news.ycombinator.com/item?id=48532640"
title: "Hillock – Local, brain-inspired AI memory using SQLite and HDC"
article_title: "GitHub - roandejager/Hillock · GitHub"
author: "roandejager"
captured_at: "2026-06-14T21:32:24Z"
capture_tool: "hn-digest"
hn_id: 48532640
score: 1
comments: 0
posted_at: "2026-06-14T20:59:50Z"
tags:
  - hacker-news
  - translated
---

# Hillock – Local, brain-inspired AI memory using SQLite and HDC

- HN: [48532640](https://news.ycombinator.com/item?id=48532640)
- Source: [github.com](https://github.com/roandejager/Hillock)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T20:59:50Z

## Translation

タイトル: Hillock – SQLite と HDC を使用した、脳にインスピレーションを得たローカル AI メモリ
記事タイトル: GitHub - roandejager/Hillock · GitHub
説明: GitHub でアカウントを作成して、roandejager/Hillock の開発に貢献します。

記事本文:
GitHub - roandejager/Hillock · GitHub
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
ロアンデジャガー
/
ヒロック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

dファイル
41 コミット 41 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md config.py config.py database.py database.py Evaluate_hillock_PROTO_ish.py Evaluate_hillock_PROTO_ish.py ingestor.py ingestor.py main.py main.py Plasticity.py Plasticity.py Required.txt Required.txt reservoir.py reservoir.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
こんにちは！これは Hillock です。これは基本的に私がハッキングしてきたローカルの個人的なメモリ システムです。標準的なベクトル データベースは、自分のコンピューターでオフラインのチャットボットを簡単に実行するには重すぎて複雑すぎると常に感じていたからです。
⚠️ 注意: このプロジェクトは現在進行中ですが、正直に言うと、それだけではありません。これは、脳をヒントにした数学を使用してローカル AI の記憶力を向上させることができるかどうかを確認するために私が取り組んでいる楽しい個人的な実験です。それは明らかに完成した製品ではないため、不格好な部品や奇妙なバグがいくつかあることが予想されます。
私はこのプロトタイプを、複雑な文構造、深い気を散らす要素、およびトリッキーな「ハード ネガティブ」クエリを使用した、大規模で非常に厳密な 30 文の科学ベンチマークを実行しました。小さなローカルの Qwen 1.5B モデルを実行すると、次のようになります。
取得精度: 30.0% (非常に複雑なクエリの一部については正しいファクトが取得されましたが、小さなモデルが抽出中に他のものを見逃しました)。
ゲート精度 : 30.0% (小さなモデル抽出エラーによりいくつかのリークが発生しましたが、多くの答えられない/幻覚的なクエリをブロックすることに成功しました)。
(これらのメトリクスの技術的な詳細と、複雑な文法で小さな 1.5B モデルを実行することが実際に非常に難しい理由については、下部のベンチマーク セクションを確認してください。)
⚙️ 仕組み (一般的な流れ)
データがシステム内をどのように移動するかを簡単に説明します。
[生のテキスト/PDF]
│
▼ (並列インジェスター)
[ オラマ (Qwen2) ]

│ │
▼ ▼
[SQLite グラフ] [ヘビアン メモリ]
│ │
━━━━┬──────┘
▼
[VSA/HDC リザーバー] ──► [ゲート コントローラー (ヒロック)]
(注: この ASCII 図は AI で作成されたものであるため、100% 正確ではない、または完全に一致しているわけではない可能性がありますが、物事がどのように接続されているかについての一般的な概念を示しています。)
基本的に、作業をいくつかの異なるレイヤーに分割します。
💾 SQLite グラフ : 永続的で確実な事実を単純なトリプル (Marie_Curie -> Born_in -> Poland など) として保存するため、システムは確かなグラウンド トゥルースを持ちます。
⚡ Hebbian Plasticity : チャットで話されているエンティティを動的に追跡し、単純なデジタル シナプスのようにエンティティ間の接続を強化します。
🌀 超次元コンピューティング (HDC) : 会話履歴で常に更新される 10,000 次元のベクトルを使用します。これは、システムが代名詞 (「彼」や「彼女」など) を解決し、幻覚を防ぐためにクエリをブロックするタイミングを決定するのに役立ちます。
この不格好なプロトタイプを実際に実行してみたい場合は、グローバル パッケージを台無しにしないように、クリーンな Python 仮想環境をセットアップすることを強くお勧めします。 Ollama をローカルにインストールして実行することも必要です。
git クローン https://github.com/roandejager/Hillock.git
cd ヒロック
2. 仮想環境のセットアップ
# 環境を作成する
Python -m venv .venv
# アクティブ化する (Windows)
.venv \S cripts \a アクティブ化
# 有効化する (Mac/Linux)
ソース .venv/bin/activate
3. 依存関係のインストールとモデルのプル
pip install -r 要件.txt
オラマ プル qwen2:1.5b
4. チャットコンソールを起動する
Python main.py
コンソール内で次のコマンドを使用できます。
/ingest [filepath] — ローカルの .txt または .pdf ファイルにインデックスを付けます。
/mode [strict/framed/conversational] — AI の会話性を変更します。
/reset — SQLite データベースをワイプし、HDC メモリ空間をリセットします。

。
📊 詳細な技術ベンチマーク
以下は、アップグレードされた非常に厳密な評価スクリプト (evaluate_hillock_PROTO_ish.py) からの正確な診断出力です。
--------------------------------------------------
* 抽出精度 : 10.6% (正しく構造化された事実ノード)
* 抽出再現率 : 22.7% (インデックス化された関係の完全性)
* 検索精度 : 30.0% (回答可能なクエリに対する事実の精度)
※ゲート精度：30.0％（幻覚防御率）
--------------------------------------------------
スコアがこのようになる理由:
10.6% の抽出精度と 22.7% の再現率: 評価セットを、量子物理学、コンピューター サイエンス、宇宙探査、哲学にまたがる 30 の膨大な複雑で複数の主題の文に押し上げました。小さな 1.5B パラメーター モデル ( qwen2:1.5b ) は、混乱せずにこれほど高密度のテキストを解析するには小さすぎます。 [James_Watson] - [discovered] -> [double-helix_model_of_DNA] または [Grace_Hopper] - [became_a_pioneer] -> [development_the_first_compiler] のような関係を幻覚させました。
「ニュートン / ガリレオ / アリストテレス」ブロック: 1.5B モデルは並列取り込みフェーズ中にクリーンな関係を解析できなかったため、これらの質問はステップ 2 で安全にブロックされました (回答できないものについては正しいブロックが生成されますが、回答可能なものについては誤ったブロックが生成されます)。
「エジソン / ファインマン」のリーク : 1.5B モデルは取り込み中にノイズの多い関係 ([ハインリッヒ・ヘルツ] -[生まれ年] -> [ハンブルク,_ドイツ] など) を抽出したため、言及されていないこと (ヘルツが誰と協力したかなど) について質問されると、出生の事実に対してゲートが開き、厳格なテスト スイートの下で「リーク」が発生しました。
ベクトル正規化 : レトリーバー マッチング自体は数学的に非常に安定しています。すべての候補ファクトを 3 つの固有のコンポーネントに厳密にバインドすることで、

(主語、目的語、および最も一致する述語の単語) により、短い事実の類似性スコアが人為的に高くなることを防ぎます。
config.py — すべてのハイパーパラメータ (HDC 寸法、減衰率など) を保持します。
Database.py — シンボリック ファクト ストレージ用の SQLite インターフェイス。
ingestor.py — ドキュメントをチャンクして解析するための並列ワーカー スレッドを生成します。
Plasticity.py — 間のヘビアン共活性化重みを追跡する
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to roandejager/Hillock development by creating an account on GitHub.

GitHub - roandejager/Hillock · GitHub
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
roandejager
/
Hillock
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
41 Commits 41 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md config.py config.py database.py database.py evaluate_hillock_PROTO_ish.py evaluate_hillock_PROTO_ish.py ingestor.py ingestor.py main.py main.py plasticity.py plasticity.py requirements.txt requirements.txt reservoir.py reservoir.py View all files Repository files navigation
Hi! This is Hillock , which is basically a local, personal memory system I've been hacking on because standard vector databases always felt way too heavy and complicated just to run a quick, offline chatbot on my own computer.
⚠️ Heads up: This project is very much a work in progress, and honestly, it isn't all that. It's just a fun personal experiment I'm working on to see if we can use brain-inspired math to make local AI memory better. It is definitely not a finished, production-ready product, so expect some clunky parts and weird bugs.
I put this prototype through a massive, highly rigorous 30-sentence scientific benchmark with complex sentence structures, deep distractors, and tricky "hard negative" queries. Running a tiny local Qwen 1.5B model, here is how it did:
Retrieval Accuracy : 30.0% (It retrieved the correct facts for some of the highly complex queries, but the tiny model missed others during extraction).
Gate Accuracy : 30.0% (It successfully blocked many unanswerable/hallucinatory queries, though some leaks occurred due to tiny model extraction errors).
(For a more detailed technical breakdown of these metrics and why running a tiny 1.5B model on complex grammar is actually quite hard, check out the Benchmark section at the bottom.)
⚙️ How It Works (The General Flow)
Here is a quick look at how data moves through the system:
[Raw Text / PDFs]
│
▼ (Parallel Ingestor)
[ Ollama (Qwen2) ]
│ │
▼ ▼
[SQLite Graph] [Hebbian Memory]
│ │
└─────┬──────┘
▼
[VSA/HDC Reservoir] ──► [Gating Controller (Hillock)]
(Note: This ASCII diagram was made with AI, so it might not be 100% correct or perfectly aligned, but it shows the general idea of how things connect.)
Basically, it splits the work into a few different layers:
💾 SQLite Graph : Stores the permanent, hard facts as simple triples (like Marie_Curie -> born_in -> Poland ) so the system has a solid ground truth.
⚡ Hebbian Plasticity : Dynamically tracks which entities are being talked about in the chat and strengthens the connections between them, like a simple digital synapse.
🌀 Hyperdimensional Computing (HDC) : Uses a 10,000-dimensional vector that constantly updates with conversational history, which helps the system resolve pronouns (like "he" or "she") and decide when to block a query to prevent hallucinations.
If you actually want to try running this clunky prototype, it is highly recommended to set up a clean Python virtual environment so you do not mess up your global packages. You will also need Ollama installed and running locally.
git clone https://github.com/roandejager/Hillock.git
cd Hillock
2. Set Up Virtual Environment
# Create the environment
python -m venv .venv
# Activate it (Windows)
.venv \S cripts \a ctivate
# Activate it (Mac/Linux)
source .venv/bin/activate
3. Install Dependencies & Pull Model
pip install -r requirements.txt
ollama pull qwen2:1.5b
4. Start the Chat Console
python main.py
Inside the console, you can use these commands:
/ingest [filepath] — Index a local .txt or .pdf file.
/mode [strict/balanced/conversational] — Change how conversational the AI is.
/reset — Wipe the SQLite database and reset the HDC memory space.
📊 Detailed Technical Benchmarks
Here is the exact diagnostic output from the upgraded, highly rigorous evaluation script ( evaluate_hillock_PROTO_ish.py ):
--------------------------------------------------
* Extraction Precision : 10.6% (Correctly structured factual nodes)
* Extraction Recall : 22.7% (Completeness of indexed relations)
* Retrieval Accuracy : 30.0% (Factual accuracy on answerable queries)
* Gate Accuracy : 30.0% (Hallucination defense rate)
--------------------------------------------------
Why the scores are what they are:
The 10.6% Extraction Precision & 22.7% Recall : We pushed the evaluation set to a massive 30 complex, multi-subject sentences spanning Quantum Physics, Computer Science, Space Exploration, and Philosophy. A tiny 1.5B parameter model ( qwen2:1.5b ) is simply too small to parse this much dense text without getting confused. It hallucinated relationships like [James_Watson] -[discovered]-> [double-helix_model_of_DNA] or [Grace_Hopper] -[became_a_pioneer]-> [developed_the_first_compiler] .
The "Newton / Galileo / Aristotle" Blocks : Because the 1.5B model failed to parse their clean relations during the parallel ingestion phase, those questions were safely blocked during step 2 (resulting in correct blocks for unanswerable ones but false blocks for answerable ones).
The "Edison / Feynman" Leaks : Because the 1.5B model extracted noisy relations during ingestion (like [Heinrich_Hertz] -[born_in]-> [Hamburg,_Germany] ), when asked about unmentioned things (like who Hertz collaborated with), the gate opened on the birth fact, resulting in "leaks" under the strict test suite.
Vector Normalization : The retriever matching itself is mathematically highly stable. By keeping all candidate facts strictly bound to exactly 3 unique components (Subject, Object, and best-matching Predicate word), we prevent shorter facts from having artificially higher similarity scores.
config.py — Holds all the hyperparameters (HDC dimensions, decay rates, etc.).
database.py — The SQLite interface for symbolic fact storage.
ingestor.py — Spawns parallel worker threads to chunk and parse documents.
plasticity.py — Tracks Hebbian co-activation weights betwee
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
