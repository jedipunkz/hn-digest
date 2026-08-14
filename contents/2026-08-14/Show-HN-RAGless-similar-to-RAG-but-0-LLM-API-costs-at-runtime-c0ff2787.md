---
source: "https://github.com/EmilResearch/RAGless"
hn_url: "https://news.ycombinator.com/item?id=49296595"
title: "Show HN: RAGless – similar to RAG, but $0 LLM API costs at runtime"
article_title: "GitHub - EmilResearch/RAGless: RAGless is a semantic retrieval system that answers questions about your documentation, without using an LLM at runtime. · GitHub"
author: "emilianoc"
captured_at: "2026-08-14T09:58:08Z"
capture_tool: "hn-digest"
hn_id: 49296595
score: 2
comments: 1
posted_at: "2026-08-14T09:51:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RAGless – similar to RAG, but $0 LLM API costs at runtime

- HN: [49296595](https://news.ycombinator.com/item?id=49296595)
- Source: [github.com](https://github.com/EmilResearch/RAGless)
- Score: 2
- Comments: 1
- Posted: 2026-08-14T09:51:11Z

## Translation

タイトル: Show HN: RAGless – RAG に似ていますが、実行時の LLM API コストは $0
記事のタイトル: GitHub - EmilResearch/RAGless: RAGless は、実行時に LLM を使用せずにドキュメントに関する質問に答えるセマンティック検索システムです。 · GitHub
説明: RAGless は、実行時に LLM を使用せずにドキュメントに関する質問に答えるセマンティック検索システムです。 - エミールリサーチ/RAGless

記事本文:
GitHub - EmilResearch/RAGless: RAGless は、実行時に LLM を使用せずにドキュメントに関する質問に答えるセマンティック検索システムです。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
エミール研究
/
ラグレス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
26 コミット 26 コミット ソース source .env.example .env.example .gitignore .gitignor

e ライセンス ライセンス README.md README.md chatbot.py chatbot.py config.py config.py ingest_to_qdrant.py ingest_to_qdrant.py prepare_data.py prepare_data.py ragless.png ragless.png 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
RAGless — 決定的検索専用 Q&A システム
RAGless は、実行時に LLM 呼び出しを行わない、検索専用の質問応答システムです。
ソース文書は自己完結型の情報ブロックに変換され、ローカル ベクトル データベース (Qdrant) にインデックス付けされ、非対称 Gemini 埋め込みを介してクエリされます。
実行時の幻覚ゼロ。遅延を最小限に抑えます。クエリあたりのコストはほぼゼロです。
プロジェクトは 3 つの独立したスクリプトで構成されています。
実行時に LLM を使用しない — チャットボットはベクトルの取得のみに依存します。ユーザー操作中に高価な API 呼び出しは必要ありません。
Answer_id 集約による堅牢な Q-Q マッチング — 同じ回答に対する複数の質問バリアントがクエリに一致する場合、それらのスコアが合計されます。これにより、結果は従来の「トップ 1」検索よりもはるかに安定します。
非対称 Gemini 埋め込み - Google が推奨する、取り込み時の RETRIEVAL_DOCUMENT と取得時の RETRIEVAL_QUERY。
組み込み Qdrant — Docker サーバーやクラウド サービスはありません。データはローカルのディスク ( ./qdrant_data ) に保存されます。
スマート チャンキング — ドキュメントは、モデルの実際のトークナイザーで測定されたトークンしきい値を超えた場合にのみチャンク化されます。
オプションのジャッジ検証 — prepare_data.py は、ソース テキスト ( --judge ) でサポートされていないブロックを破棄する 2 番目の LLM パスを有効にできます。
欠落したクエリのログ — しきい値を下回るクエリは、後で分析できるように、missed_queries.log に自動的に記録されます。
冪等性の保証 — ingest_to_qdrant.py を実行するたびにコレクションが最初から再作成され、隠れた重複が防止されます。
Gemini API キー (generou の無料利用枠)

の制限)
Python の依存関係 (「インストール」セクションを参照)
リポジトリを複製またはダウンロードし、プロジェクト フォルダーに移動します。
仮想環境を作成します (推奨):
Python -m venv venv
ソース venv/bin/activate # Linux/macOS
# または
venv \S cripts \a ctivate # Windows
pip install litellm qdrant-client pypdf python-dotenv tqdm
cp .env.example .env
# .env を編集して GEMINI_API_KEY を挿入します
。
§──source/ # ソースドキュメント（.pdf、.txt、.md）が含まれるフォルダー
§── config.py # 一元的な設定 (モデル、しきい値、パス)
§── prepare_data.py # スクリプト 1: Q&A ブロックの抽出
§── ingest_to_qdrant.py # スクリプト 2: 埋め込みとインデックス作成
§── chatbot.py # スクリプト 3: CLI チャットボット
§── data.json # prepare_data.py の出力（検証済みブロック）
§── qdrant_data/ # ローカルベクトルデータベース（自動作成）
§── failed_chunks/ # 有効な JSON の生成に失敗したチャンク (デバッグ)
└── miss_queries.log # しきい値を下回ったクエリのログ
使用法
1. データを準備します ( prepare_data.py )
ドキュメントを source/ フォルダー ( .pdf 、 .txt 、 .md をサポート) に配置し、次のコマンドを実行します。
Python prepare_data.py
オプションのジャッジ検証を使用すると (時間はかかりますが、より正確です):
python prepare_data.py --judge
機能:
各ファイルを読み取り、トークンをカウントします。
ドキュメントが短い場合 (≤ 10,000 トークン)、ドキュメント全体を LLM に送信します。それ以外の場合は、チャンクに分割します。
Answer 、questions、category 、source_quote を含む JSON ブロックを抽出します。
ブロックを検証し、 data.json に保存します。
2. ベクター データベースへのインデックス作成 ( ingest_to_qdrant.py )
Python ingest_to_qdrant.py
機能:
各ブロックを質問のバリエーションと同じ数の行に「分解」します。
LiteLLM + Gemini を介してバッチでエンベディングを生成します。
Qdrant コレクションを再作成し、決定的な UUID5 を持つベクトルを挿入します。
3. 起動

h チャットボット ( chatbot.py )
Pythonチャットボット.py
利用可能なオプション:
python chatbot.py --threshold 0.75 # 集計スコアの最小しきい値を変更します
python chatbot.py --debug # 内部スコアと集計テーブルを表示する
インタラクション:
あなた> チェックインはどのように行われますか?
[情報] 関連する回答が 2 件見つかり、上位 1 件を表示:
─────────────────────
--- 回答 1 (該当性: 1.85) ---
チェックインは午後 3 時から午後 8 時までご利用いただけます。 20時以降にご到着の場合は、
事前にフロントまでご連絡ください。...
─────────────────────
出典:source/regulations.txt
終了するには、 exit 、 quit 、または :q を入力します。
チャットボットの中核は、answer_id による集計です。
ユーザークエリは、 task_type=RETRIEVAL_QUERY で埋め込まれます。
Qdrant は、TOP_K_RETRIEVAL で最も類似したポイント (質問) を返します。
同じanswer_idを指すポイントのスコアが合計されます。
候補は次の場合にのみ表示されます。
集計スコアが DEFAULT_THRESHOLD を超えている、または
ベスト シングル ヒットが SINGLE_HIT_THRESHOLD を超える
かつ、最良の単一ヒットは > 0.68 (最低品質)
このメカニズムにより、システムは堅牢になります。たとえ 1 つの質問のバリエーションが完全に一致しない場合でも、同じ回答に対する複数の弱い一致の合計により、その答えが正しく表示されます。
RAGless とクラシックな生成 RAG
Classic RAG は、コンテキストを取得し、LLM に応答を合成するように指示することで、実行時に応答を生成します。 RAGless では生成ステップが完全に排除されます。回答は取り込み中に事前に生成され、取得されます

実行時に逐語的に。
利点
説明
はるかにシンプルなパイプライン
回答生成のためのプロンプトエンジニアリング、コンテキストウィンドウ管理、出力解析はありません。埋め込み、検索し、返すだけです。
Q-Q マッチングの信頼性がはるかに高い
クエリと質問 (Q-Q) のマッチングは、クエリとドキュメント チャンク (Q-D) やクエリと回答 (Q-A) よりも意味的に簡単で堅牢です。回答ごとに複数の質問バリエーションがあるため、冗長性が得られます。
決定論
いつも同じ質問、同じ答え。動作は再現可能でテスト可能です。
実行時の幻覚ゼロ
LLM はクエリ時に応答を生成しません。返されるテキストは事前に生成されており、変更できません。
クエリあたりのコストはゼロ
取り込み後は、API 呼び出しはありません。取得は純粋にローカルな計算です。
低遅延
クエリ埋め込み + ベクトル検索。秒ではなくミリ秒です。
検証可能な回答
すべてのブロックには、source_quote と source_file があります。監査証跡を完了します。
再現可能なバグ
答えが間違っている場合、それは 100% 間違っています。見つけて修正するのが簡単です。
有限の出力曲面
考えられる答えの数はわかっています ( data.json )。徹底的にテスト可能。
適度なハードウェアで動作
Qdrant が組み込まれており、メモリ内に LLM はありません。数 GB の RAM を搭載した CPU で動作します。
完全なプライバシー
取り込み後にユーザー データがマシンから流出することはありません。
依存関係のドリフトがない
埋め込みモデルは変わる可能性がありますが、答えは変わりません。 LLM プロバイダーの可用性や価格に拘束されることはありません。
トレードオフと制限事項
制限事項
説明
リアルタイムの柔軟性がない
新しい答えを合成したり、ブロック全体で情報を組み合わせたり、トーンを動的に適応させたりすることはできません。あなたが摂取したものはあなたが得るものです。
取り込みコストが高い
LLM を使用して Q&A ブロックを生成すると、単純な埋め込みよりもコストがかかります。以下のコスト比較を参照してください。
摂取に限定された適用範囲
prepare_data.py 中にトピックが抽出されなかった場合、システムはそれに応答できません。いいえ、「理由」

隙間周りはng」。
メンテナンスには再摂取が必要です
回答を更新するには、プロンプトの編集だけでなく、パイプライン全体を再実行する必要があります。
コストとパフォーマンスの比較
メトリック
クラシックな生成 RAG
Q-Qシステム（本プロジェクト）
違い・利点
取り込みコスト (1 回限り)
~0.01 ドル (100,000 トークンの埋め込みのみ)
~$1.50 (LLM は 100,000 トークンから Q&A を生成します)
+$1.49 (初期のデメリットはありますが、コストは無視できます)
ランタイム API コスト (月額)
~$157.50 (1,000 クエリ/日、~1,000 コンテキスト トークン + 150 出力トークン)
$0.00 (1,000 個の短いクエリの埋め込みのみ、月額 < $0.02)
~$157.50/月を節約 (コストゼロの拡張性)
レイテンシー (クエリごと)
2.5 ～ 4 秒 (LLM テキスト生成)
~0.15 秒 (純粋なベクトル検索)
>15 倍高速 (即時応答)
幻覚率 (実行時)
低いが常に > 0%
0%
リスクの排除 (静的で確定的な出力)
幻覚のトレードオフ
RAGless は、制御不能で最も危険な実行時の幻覚を排除します。取り込み時のリスクは依然として存在しますが、ナレッジ ベースが本番環境に移行する前に人によるレビューが可能となり、管理された環境でオフラインでリスクが軽減されます。
RAGless は、幻覚リスクを実行時から取り込みに移します。生成は prepare_data.py 中に行われます。そのため、オプションの --judge パスとべき等性のための決定論的な UUID が追加されました。
トレードオフは、オフラインでの検証可能性と引き換えにリアルタイムの柔軟性を放棄することです。リスクの高いドメインの場合、私は予測できない幻覚よりも、ログで捕捉できる幻覚を好みます。
すべての調整可能な定数は config.py にあります。
問題
解決策
GEMINI_API_KEY が見つかりません
GEMINI_API_KEY=... を使用して .env ファイルを作成します。
コレクションが見つかりません
最初に python ingest_to_qdrant.py を実行します
チャンク内の不正な形式の JSON
failed_chu を確認してください

nks/生テキスト用フォルダー
空の LLM 応答
ジェミニの安全ブロックの可能性。ソーステキストを減らすか変更してみてください
Qdrant ロックファイル
クライアントは自動的に終了します。クラッシュが発生した場合は、qdrant_data/ のロックを手動で削除します。
テクニカルノート
LiteLLM は、補完と埋め込みの両方のために Gemini を呼び出すための統合プロキシとして使用されます。
path= モードの Qdrant クライアントはすべてをローカル ファイルに保存します。サーバー プロセスは必要ありません。
Qdrant ポイント ID は決定的な UUID5 ( uuid5(NAMESPACE_DNS, Answer_id:question_text) ) であるため、取り込みを再実行しても論理的な重複は作成されません。
このプロジェクトは、GNU Affero General Public License v3.0 (AGPLv3) に基づいてライセンスされています。
RAGless は、実行時に LLM を使用せずにドキュメントに関する質問に答えるセマンティック検索システムです。
Readme MIT ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

RAGless is a semantic retrieval system that answers questions about your documentation, without using an LLM at runtime. - EmilResearch/RAGless

GitHub - EmilResearch/RAGless: RAGless is a semantic retrieval system that answers questions about your documentation, without using an LLM at runtime. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
EmilResearch
/
RAGless
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
26 Commits 26 Commits source source .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md chatbot.py chatbot.py config.py config.py ingest_to_qdrant.py ingest_to_qdrant.py prepare_data.py prepare_data.py ragless.png ragless.png requirements.txt requirements.txt View all files Repository files navigation
RAGless — Deterministic Retrieval-Only Q&A System
RAGless is a retrieval-only question-answering system with zero LLM calls at runtime .
Source documents are converted into self-contained informational blocks, indexed into a local vector database (Qdrant), and queried via asymmetric Gemini embeddings.
Zero hallucinations at runtime. Minimal latency. Near-zero cost per query.
The project consists of three independent scripts:
No LLM at runtime — The chatbot relies purely on vector retrieval. No expensive API calls during user interaction.
Robust Q-Q matching via answer_id aggregation — If multiple question variants for the same answer match the query, their scores are summed. This makes the result far more stable than classic "top-1" retrieval.
Asymmetric Gemini embeddings — RETRIEVAL_DOCUMENT at ingestion time and RETRIEVAL_QUERY at retrieval time, as recommended by Google.
Embedded Qdrant — No Docker server, no cloud service. Data is stored locally on disk ( ./qdrant_data ).
Smart chunking — Documents are chunked only if they exceed a token threshold, measured with the model's real tokenizer.
Optional Judge verification — prepare_data.py can enable a second LLM pass to discard blocks not supported by the source text ( --judge ).
Missed query logging — Below-threshold queries are automatically logged to missed_queries.log for later analysis.
Guaranteed idempotency — Every run of ingest_to_qdrant.py recreates the collection from scratch, preventing hidden duplicates.
Gemini API Key (free tier with generous limits)
Python dependencies (see Installation section)
Clone or download the repository and navigate to the project folder.
Create a virtual environment (recommended):
python -m venv venv
source venv/bin/activate # Linux/macOS
# or
venv \S cripts \a ctivate # Windows
pip install litellm qdrant-client pypdf python-dotenv tqdm
cp .env.example .env
# Edit .env and insert your GEMINI_API_KEY
.
├── source/ # Folder with source documents (.pdf, .txt, .md)
├── config.py # Centralized configuration (models, thresholds, paths)
├── prepare_data.py # Script 1: Q&A block extraction
├── ingest_to_qdrant.py # Script 2: embedding and indexing
├── chatbot.py # Script 3: CLI chatbot
├── data.json # Output of prepare_data.py (validated blocks)
├── qdrant_data/ # Local vector database (auto-created)
├── failed_chunks/ # Chunks that failed to produce valid JSON (debug)
└── missed_queries.log # Log of below-threshold queries
Usage
1. Prepare the data ( prepare_data.py )
Place your documents in the source/ folder (supports .pdf , .txt , .md ), then run:
python prepare_data.py
With optional Judge verification (slower but more accurate):
python prepare_data.py --judge
What it does:
Reads each file and counts tokens.
If the document is short (≤ 10,000 tokens), sends it whole to the LLM; otherwise splits it into chunks.
Extracts JSON blocks with answer , questions , category , source_quote .
Validates blocks and saves them to data.json .
2. Index into the vector database ( ingest_to_qdrant.py )
python ingest_to_qdrant.py
What it does:
"Explodes" each block into as many rows as its question variants.
Generates embeddings in batches via LiteLLM + Gemini.
Recreates the Qdrant collection and inserts vectors with deterministic UUID5s.
3. Launch the chatbot ( chatbot.py )
python chatbot.py
Available options:
python chatbot.py --threshold 0.75 # Change the minimum aggregated score threshold
python chatbot.py --debug # Show internal scores and aggregation table
Interaction:
You> How does check-in work?
[INFO] Found 2 relevant answers, showing top 1:
──────────────────────────────────────────────────────────────────────
--- Answer 1 (Pertinence: 1.85) ---
Check-in is available from 3:00 PM to 8:00 PM. If you arrive after 8:00 PM,
please contact reception in advance...
──────────────────────────────────────────────────────────────────────
Source: source/regulations.txt
Type exit , quit , or :q to leave.
The core of the chatbot is aggregation by answer_id :
The user query is embedded with task_type=RETRIEVAL_QUERY .
Qdrant returns the TOP_K_RETRIEVAL most similar points (questions).
Scores of points pointing to the same answer_id are summed .
A candidate is shown only if:
the aggregated score exceeds DEFAULT_THRESHOLD OR
the best single hit exceeds SINGLE_HIT_THRESHOLD
AND the best single hit is > 0.68 (minimum quality)
This mechanism makes the system robust: even if no single question variant is a perfect match, the sum of multiple weak matches on the same answer can make it emerge correctly.
RAGless vs. Classic Generative RAG
Classic RAG generates answers at runtime by retrieving context and prompting an LLM to synthesize a response. RAGless eliminates the generative step entirely. Answers are pre-generated during ingestion and retrieved verbatim at runtime.
Advantage
Explanation
Much simpler pipeline
No prompt engineering for answer generation, no context window management, no output parsing. Just embed, search, return.
Q-Q matching is far more reliable
Matching query-to-question (Q-Q) is semantically easier and more robust than query-to-document-chunk (Q-D) or query-to-answer (Q-A). Multiple question variants per answer provide redundancy.
Determinism
Same question, same answer, always. Behavior is reproducible and testable.
Zero hallucinations at runtime
No LLM generates answers at query time. Returned text is pre-generated and immutable.
Zero cost per query
After ingestion, there are no API calls. Retrieval is purely local computation.
Low latency
Query embedding + vector search. Milliseconds, not seconds.
Verifiable answers
Every block has source_quote and source_file . Complete audit trail.
Reproducible bugs
If an answer is wrong, it is 100% wrong. Easy to find and fix.
Finite output surface
The number of possible answers is known ( data.json ). Exhaustively testable.
Runs on modest hardware
Embedded Qdrant, no LLM in memory. Works on CPU with a few GB of RAM.
Total privacy
No user data leaves the machine after ingestion.
No dependency drift
The embedding model can change, but the answers do not. You are not tied to an LLM provider's availability or pricing.
Trade-offs and Limitations
Limitation
Explanation
No real-time flexibility
Cannot synthesize novel answers, combine information across blocks, or adapt tone dynamically. What you ingest is what you get.
Higher ingestion cost
Using an LLM to generate Q&A blocks costs more than simple embedding. See cost comparison below.
Coverage bounded by ingestion
If a topic was not extracted during prepare_data.py , the system cannot answer it. No "reasoning" around gaps.
Maintenance requires re-ingestion
Updating answers requires re-running the full pipeline, not just editing a prompt.
Cost & Performance Comparison
Metric
Classic Generative RAG
Q-Q System (This Project)
Difference / Advantage
Ingestion Cost (One-time)
~$0.01 (embedding 100,000 tokens only)
~$1.50 (LLM generates Q&A from 100,000 tokens)
+$1.49 (Initial disadvantage, but negligible cost)
Runtime API Cost (Monthly)
~$157.50 (1,000 queries/day, ~1,000 context tokens + 150 output tokens)
$0.00 (only embedding 1,000 short queries, < $0.02/month)
~$157.50/month saved (Zero-cost scalability)
Latency (per query)
2.5 – 4 seconds (LLM text generation)
~0.15 seconds (pure vector search)
>15x faster (Instant response)
Hallucination Rate (Runtime)
Low, but always > 0%
0%
Risk eliminated (Static, deterministic output)
The Hallucination Trade-off
RAGless eliminates hallucinations at runtime , where they are most dangerous because they are uncontrollable. The risk during ingestion still exists, but it is mitigated — and crucially — it happens offline, in a controlled environment, with the possibility of human review before the knowledge base goes to production.
RAGless shifts the hallucination risk from runtime to ingestion. Generation happens during prepare_data.py , which is why the optional --judge pass and deterministic UUIDs for idempotency were added.
The trade-off is: you give up real-time flexibility in exchange for offline verifiability. For high-risk domains, I prefer hallucinations I can catch in a log over ones I cannot predict.
All tunable constants are in config.py :
Problem
Solution
GEMINI_API_KEY not found
Create .env file with GEMINI_API_KEY=...
Collection not found
Run python ingest_to_qdrant.py first
Malformed JSON in chunks
Check the failed_chunks/ folder for raw text
Empty LLM response
Possible Gemini safety block; try reducing or modifying source text
Qdrant lockfile
Client closes automatically; in case of crash, manually remove the lock in qdrant_data/
Technical Notes
LiteLLM is used as a unified proxy to call Gemini for both completions and embeddings.
Qdrant Client in path= mode stores everything in local files: no server process is required.
Qdrant point IDs are deterministic UUID5s ( uuid5(NAMESPACE_DNS, answer_id:question_text) ), so re-running ingestion does not create logical duplicates.
This project is licensed under the GNU Affero General Public License v3.0 (AGPLv3) .
RAGless is a semantic retrieval system that answers questions about your documentation, without using an LLM at runtime.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
