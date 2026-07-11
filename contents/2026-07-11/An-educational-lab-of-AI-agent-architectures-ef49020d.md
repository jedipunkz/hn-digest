---
source: "https://github.com/Rudnik-Ilia/Agents-Sandbox"
hn_url: "https://news.ycombinator.com/item?id=48872922"
title: "An educational lab of AI agent architectures"
article_title: "GitHub - Rudnik-Ilia/Agents-Sandbox: A comprehensive educational lab for AI agent architectures built on LangChain & Ollama. Compare diverse patterns: simple chat memory, advanced tool calling (ReAct), and multiple RAG backends (Chroma, LlamaIndex, NumPy, Rerank). Study intelligent routing (Adapt/Se\n[truncated]"
author: "ilia53"
captured_at: "2026-07-11T15:45:00Z"
capture_tool: "hn-digest"
hn_id: 48872922
score: 1
comments: 0
posted_at: "2026-07-11T15:33:22Z"
tags:
  - hacker-news
  - translated
---

# An educational lab of AI agent architectures

- HN: [48872922](https://news.ycombinator.com/item?id=48872922)
- Source: [github.com](https://github.com/Rudnik-Ilia/Agents-Sandbox)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T15:33:22Z

## Translation

タイトル: AI エージェント アーキテクチャの教育ラボ
記事のタイトル: GitHub - Rudnik-Ilia/Agents-Sandbox: LangChain と Ollama 上に構築された AI エージェント アーキテクチャのための包括的な教育ラボ。単純なチャット メモリ、高度なツール呼び出し (ReAct)、および複数の RAG バックエンド (Chroma、LlamaIndex、NumPy、Rerank) など、さまざまなパターンを比較します。インテリジェントルーティングの検討（Adapt/Se）
[切り捨てられた]
説明: LangChain と Ollama に基づいて構築された AI エージェント アーキテクチャの包括的な教育ラボ。単純なチャット メモリ、高度なツール呼び出し (ReAct)、および複数の RAG バックエンド (Chroma、LlamaIndex、NumPy、Rerank) など、さまざまなパターンを比較します。ユニを介したインテリジェント ルーティング (Adapt/Semant) とエージェント ワークフローを学習します。
[切り捨てられた]

記事本文:
GitHub - Rudnik-Ilia/Agents-Sandbox: LangChain と Ollama 上に構築された AI エージェント アーキテクチャのための包括的な教育ラボ。単純なチャット メモリ、高度なツール呼び出し (ReAct)、および複数の RAG バックエンド (Chroma、LlamaIndex、NumPy、Rerank) など、さまざまなパターンを比較します。統合された CLI を介してインテリジェント ルーティング (Adapt/Semant) とエージェント ワークフローを学習します。 LLM が実際に何をすべきかをどのように決定するかを見てみましょう。 · GitHub
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
別の場所でサインアウトしました

r タブまたはウィンドウ。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ルドニク＝イリア
/
エージェント-サンドボックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット データ/コーパス データ/コーパス ルール ルール スキル スキル src/ localagent src/ localagent .env.example .env.example .gitignore .gitignore README.md README.md SOUL.md SOUL.md mcp.json.example mcp.json.example png.png png.png pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LangChain とローカル アーキテクチャに基づいて構築された AI エージェント アーキテクチャの教育ラボ
オラマサーバー。各バリアントは個別の実行可能な CLI なので、1 つずつ学習できます。
メカニズムを一度に確認し、ログを通じて動作を確認します。
モデル (世代): gemma4:e4b
モデル (埋め込み): mxbai-embed-large:latest
ホスト: http://10.100.102.10:11434 (構成可能)
チャットプロバイダー: デフォルトではローカルの Ollama、または LLM_PROVIDER=openrouter を指定した OpenRouter
カテゴリ
バリアント
コマンド
チャット+記憶
バッファがいっぱいです
チャットバッファ
チャット+記憶
サマリーバッファ
チャットの概要
ツール呼び出し
ReAct テキスト プロトコル
ツールの反応
ツール呼び出し
ネイティブバインドツール
ツールネイティブ
ラグのみ
純粋な数百余弦
ぼろぼろ
ラグのみ
クロマベクトルDB
ラグクロマ
ラグのみ
LlamaIndex (メモリ内)
ラグラマインデックス
ラグのみ
LlamaIndex + Chroma ストア
ラグ-ラマインデックス-クロマ
ラグのみ
干し草の山
ぼろ干し草の山
ラグのみ
ハイブリッド BM25+高密度 + クロスエンコーダーのリランク
ラグリランク
ハイブリッド (チャット + RAG)
セマンティックルーター
ハイブリッドセマンティック
ハイブリッド (チャット + RAG)
LLM 分類ルーター
ハイブリッド-llm
ハイブリッド (チャット + RAG)
アダプティブ RAG (LangGraph)
ハイブリッド対応
ハイブリッド (チャット + RAG)
修正 RAG (マルチグレーダー + rew)

儀式)
ハイブリッド アダプティブ プラス
エージェントティック RAG
ツール呼び出し + ツールとしての取得
エージェントラグ
チャット コマンドは --persist を受け入れて、実行全体で SQLite にメモリを節約します。
依存関係を仮想環境にインストールします。
UV同期
接続を構成します (ホストが異なる場合はコピーして編集します)。
.env.example .env をコピー
チャット生成に OpenRouter を使用するには、 LLM_PROVIDER=openrouter を設定します。
OPENROUTER_API_KEY 、および .env の OPENROUTER_MODEL。埋め込みはまだ使用されています
Ollama は EMBED_MODEL 経由なので、RAG コマンドには引き続きローカル埋め込みモデルが必要です。
すべてのエージェントは独自のコンソール スクリプトです。それらのいずれかを uv run <command> で実行します。
UV実行チャットバッファ
uv run chat-summary --persist
ツール呼び出し:
UV 実行ツール-反応
UV 実行ツールネイティブ
RAG (各バックエンド 1 つ):
UV 実行 rag-numpy
UV ラン ラグ クロマ
uv 実行 rag-llamaindex
uv 実行 rag-llamaindex-chroma
UV 実行 rag-haystack
UV 実行 rag-rerank
ハイブリッド (チャット + RAG):
UV 実行ハイブリッド セマンティック
UV 実行ハイブリッド LLM
UVランハイブリッド対応
UV ラン ハイブリッド アダプティブ プラス
エージェント RAG と評価:
UV 実行エージェントラグ
UV 実行 rag-eval
共通フラグ: --no-soul (すべてのエージェント)、--persist (チャット)、--no-index および --drop (RAG/ハイブリッド/エージェント)。
各コマンドの動作 (簡単に言うと)
chat-buffer - このチャットで発言した内容をすべて記憶するチャットボット。
chat-summary - チャットが長くなった場合に短い概要を記憶するチャットボット。
tools-react - ツール (電卓など) をテキストとして書き出すことで使用できるボット。
tools-native - 同じ考え方ですが、モデルがサポートする「公式」の方法でツールを呼び出します。
rag-numpy - 最初にドキュメントを読んで質問に答えます。最も単純なバージョン。
rag-chroma - 同じですが、読み取った内容を保存するので、毎回再読み込みする必要はありません。
rag-llamaindex / rag-haystack - LlamaIndex / Haystack ライブラリで構築された同じアイデア。
rag-llamaindex-chroma - LlamaIndex の実行

Chroma から保存されたストアを使用して読み取り値を取得します。
rag-rerank - よりスマートな検索: 2 つの方法で検索し、結果を再ソートして最良のものを選択します。
hybrid-semantic / hybrid-llm - 「ドキュメントを読むべきか、それともチャットするべきか?」を決定します。答える前に。
ハイブリッド適応 - 最初にドキュメントを読み、次に「これは役に立つか?」を確認します。そうでない場合はチャットします。
hybrid-adaptive-plus - 慎重なバージョン: 各ドキュメントと回答をチェックし、必要に応じて再試行します。
Agentic-rag - ボットは、いつドキュメントを検索するか、いつツールを使用するか、または単に応答するかを自ら決定します。文書の検索はその「ツール」の 1 つとして与えられています。
rag-eval - チャットではありません。これは、各 RAG バージョンが適切なドキュメントをどれだけうまく見つけたかをスコア化する小規模なテストであり、数値で比較できます。
REPL 内: メッセージを入力するか、スラッシュ コマンドを使用します。
/skills は、利用可能なスキルをリストします (番号付きメニュー)。
/<number> または /<name> は、アクティブなコンテキストにスキルをロードします。
/add [パス] は、再起動せずにテキスト ファイルを RAG ストアにライブで取り込みます。
パスがない (または無効な) 場合、data/corpus/ をスキャンし、新しいファイルを追加します。
まだインデックスされていません。 RAG およびハイブリッド エージェントで動作します。 Chroma では持続します。
/win は、現在モデルに送信されている完全なコンテキスト ウィンドウを出力します。
/active 、 /clear 、 /rules 、 /help 、 /exit 。
ツール呼び出しエージェント ( tools-react 、 tools-native ) は、任意のツールを使用できます。
モデル コンテキスト プロトコル サーバー。例をコピーして編集します。
mcp.json.example をコピー mcp.json
各エントリは、stdio サーバー ( command + args ) または HTTP サーバーのいずれかです。
( URL + トランスポート )。起動時にエージェントが接続し、サーバーのツールをロードし、
そしてそれらを組み込みとマージします (名前の衝突では組み込みが優先されます)。 mcp.jsonの場合
が存在しないか、サーバーに障害が発生した場合、エージェントは単純に組み込みツールを使用して実行されます。
必要に応じて、別の構成ファイルを指すように MCP_CONFIG を設定します。
SOUL.md (プロジェクトルート)

すべてのエージェントの先頭に追加されるグローバル ペルソナ/アイデンティティです
システムプロンプト。すべてのエージェントは --no-soul を受け入れてロードをスキップします。
エージェントは永続的なファクトを SOUL.md のマネージド ## メモリ セクションに書き込むことができます。
(メモリ:開始/メモリ:終了マーカー間、追加専用 + 重複排除):
任意のエージェントの /remember <fact> コマンド、または
ツール呼び出しエージェント。新しい記憶は次のターンに適用されます (再開はありません)。
rules/ は、すべてのエージェントのシステムに注入される常時オンのマークダウンを保持します。
自動的にプロンプトが表示されます。
Skills/ は、/ メニューを介してオンデマンドでロードされるオプトイン マークダウン スキルを保持します。をドロップします
skill/ 内の新しい .md ファイル (オプションで名前: / 説明: フロントマター)
そしてメニューに表示されます。
4 つの RAG バックエンドすべてが 1 つの Retriever インターフェイス (インデックス/検索) を実装します。
そして同じ RetrievedChunk(text,source,score) を返すため、エージェントは
後ろのエンジンに関係なく同じです。流れ:
data/corpus/ 内のすべてのファイルをロードし、チャンクに分割します。
チャンクを Ollama mxbai-embed-large に送信して、チャンクごとに 1 つのベクトルを取得します。
ベクトルをバックエンドに保存します (numpy 行列、Chroma、LlamaIndex、Haystack)。
同じモデルにクエリを埋め込みます。
コサイン類似度によってすべてのチャンクにスコアを付け、上位 k (デフォルトは 4) を取得します。
何も取得されない場合、エージェントは「わかりません」と応答します。
それ以外の場合は、Context: [source] chunk... / Question: ... を構築して送信します。
システムルールを使用した gemma4:e4b へ: コンテキストからのみ回答し、引用します
ソースを [ソース] として指定します。
検索ヒット (ソース + スコア) と LLM 呼び出しがログに記録されます。
2 つの特徴的な特徴: RAG 専用エージェントはステートレスです (各クエリは
独立しており、会話の記憶はありません）、埋め込みは別のモデルからのものです
( mxbai-embed-large ) は世代 ( gemma4:e4b ) よりも優れています。
pure-numpy ( rag-numpy ): 最も透明です。キャラクタウィンドウごとのチャンク
オーバーラ付き

p、L2 ベクトルを正規化し、総当たりの内積を実行します。
コサイン類似度の場合は (行列 @ query_vec )。データベースなし、クエリごとに O(N) - 素晴らしい
理解には及ばないが、規模感には劣る。
pure-chroma ( rag-chroma ): 同じ考え方ですが、ベクトルは永続的に存在します。
最近傍検索を実行する Chroma コレクション ( .chroma/ )。それは
生き残った唯一のバックエンドが再起動します。コレクションがすでにデータを保持している場合は、
インデックスが再作成されるのではなく、再利用されます。
llamaindex ( rag-llamaindex ): LlamaIndex の VectorStoreIndex を使用します。
文分割ツールと Ollama 埋め込み/LLM、メモリに保持されます。
llamaindex-chroma ( rag-llamaindex-chroma ): 同じ LlamaIndex オーケストレーション、
ただし、その背後には永続的な Chroma ベクトル ストアがあります。これは共通のことを示しています
プロダクション レイヤ化 - フレームワークがパイプラインを所有し、外部のスケーラブルなフレームワークがパイプラインを所有します。
ストアはベクトルを所有します。私たち自身のコードが行う rag-chroma と比較してください。
オーケストレーションと Chroma は直接駆動されます。
haystack ( rag-haystack ): Haystack コンポーネントを使用します (ドキュメント スプリッター、
インメモリ ストア、埋め込みレトリバー) が同じインターフェイスに接続されています。
チャンクが異なるため (文字ごとに分割された numpy/chroma、LlamaIndex と
干し草の山を文/単語ごとに分割）、同じ質問をわずかに取得できます
バックエンド間の異なるパッセージ - 比較するのは良いことです。
rerank ( rag-rerank ): 小さなローカル上の実稼働スタイルの取得スタック
スケール。ハイブリッド検索を実行します - 密集 (Ollama 埋め込み、コサイン) プラス
sparse (BM25 キーワード) - 2 つのランキングを相互ランク融合で融合し、
クロスエンコーダー (cross-encoder/ms-marco-MiniLM-L-6-v2) を使用して候補を再ランク付けします
デフォルトで)、top-k を維持します。クロスエンコーダーにより、より明確な関連性が得られます
生のコサインよりも分離 (右チャンクの明らかにプラスのスコア、強く
残りはマイナス）。モデル

初回実行時にダウンロード。経由で設定します
RERANK_MODEL と FUSION_CANDIDATES を介した候補者プール。
すべての RAG エージェント (およびすべてのハイブリッド エージェント) は以下を受け入れます。
--no-index - コーパスのインデックス作成/アップロードを行わずに開始します。インメモリ
バックエンドは空から始まります。 Chroma は、すでに永続化されているものをすべて使用します。
--drop - 開始前に永続化された RAG ストレージを削除します (Chroma のみ。
インメモリバックエンド)、クリーンリビルドを強制します。
デフォルトでは、メモリ内のバックエンドは実行のたびにコーパスにインデックスを付けますが、Chroma は
永続化されたコレクションが存在する場合は再利用します。
uv 実行 rag-chroma --drop
uv 実行 rag-chroma --no-index
エージェントティック RAG
Agentic-rag は、ツールと取得の組み合わせを別の観点から捉えたものです。
ルーターが RAG 対チャットを事前に決定し、取得はネイティブに公開されます。
search_knowledge_base ツールとしてのツール呼び出しエージェント (ハイブリッドによって支援される)
BM25+高密度+リランクエンジン)。モデルはターンごとに、
コーパス、別のツール (電卓、シェル、web_search、MCP) を使用するか、回答してください
直接的に。事実に関する質問については知識ベースを優先するように指示されています。
UV 実行エージェントラグ
プロダクション機能
人間参加型: 危険なツール (現在はシェル) は確認が必要です
走る前に。 REQUIRE_TOOL_A によって制御されます

[切り捨てられた]

## Original Extract

A comprehensive educational lab for AI agent architectures built on LangChain & Ollama. Compare diverse patterns: simple chat memory, advanced tool calling (ReAct), and multiple RAG backends (Chroma, LlamaIndex, NumPy, Rerank). Study intelligent routing (Adapt/Semant) and Agentic workflows via a uni
[truncated]

GitHub - Rudnik-Ilia/Agents-Sandbox: A comprehensive educational lab for AI agent architectures built on LangChain & Ollama. Compare diverse patterns: simple chat memory, advanced tool calling (ReAct), and multiple RAG backends (Chroma, LlamaIndex, NumPy, Rerank). Study intelligent routing (Adapt/Semant) and Agentic workflows via a unified CLI. See how LLMs actually decide what to do! · GitHub
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
Rudnik-Ilia
/
Agents-Sandbox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits data/ corpus data/ corpus rules rules skills skills src/ localagent src/ localagent .env.example .env.example .gitignore .gitignore README.md README.md SOUL.md SOUL.md mcp.json.example mcp.json.example png.png png.png pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
An educational lab of AI agent architectures, built on LangChain and a local
Ollama server. Each variant is a separate, runnable CLI so you can study one
mechanism at a time and watch it work through the logs.
Model (generation): gemma4:e4b
Model (embeddings): mxbai-embed-large:latest
Host: http://10.100.102.10:11434 (configurable)
Chat provider: local Ollama by default, or OpenRouter with LLM_PROVIDER=openrouter
Category
Variant
Command
Chat + memory
full buffer
chat-buffer
Chat + memory
summary buffer
chat-summary
Tool calling
ReAct text protocol
tools-react
Tool calling
native bind_tools
tools-native
RAG only
pure numpy cosine
rag-numpy
RAG only
Chroma vector DB
rag-chroma
RAG only
LlamaIndex (in-memory)
rag-llamaindex
RAG only
LlamaIndex + Chroma store
rag-llamaindex-chroma
RAG only
Haystack
rag-haystack
RAG only
hybrid BM25+dense + cross-encoder rerank
rag-rerank
Hybrid (chat + RAG)
semantic router
hybrid-semantic
Hybrid (chat + RAG)
LLM classifier router
hybrid-llm
Hybrid (chat + RAG)
adaptive RAG (LangGraph)
hybrid-adaptive
Hybrid (chat + RAG)
corrective RAG (multi-grader + rewrite)
hybrid-adaptive-plus
Agentic RAG
tool-calling + retrieval-as-a-tool
agentic-rag
The chat commands accept --persist to save memory to SQLite across runs.
Install dependencies into a virtual environment:
uv sync
Configure the connection (copy and edit if your host differs):
copy .env.example .env
To use OpenRouter for chat generation, set LLM_PROVIDER=openrouter ,
OPENROUTER_API_KEY , and OPENROUTER_MODEL in .env . Embeddings still use
Ollama via EMBED_MODEL , so RAG commands still need the local embedding model.
Every agent is its own console script. Run any of them with uv run <command> .
uv run chat-buffer
uv run chat-summary --persist
Tool calling:
uv run tools-react
uv run tools-native
RAG (one backend each):
uv run rag-numpy
uv run rag-chroma
uv run rag-llamaindex
uv run rag-llamaindex-chroma
uv run rag-haystack
uv run rag-rerank
Hybrid (chat + RAG):
uv run hybrid-semantic
uv run hybrid-llm
uv run hybrid-adaptive
uv run hybrid-adaptive-plus
Agentic RAG and evaluation:
uv run agentic-rag
uv run rag-eval
Common flags: --no-soul (all agents), --persist (chat), --no-index and --drop (RAG/hybrid/agentic).
What each command does (in plain words)
chat-buffer - a chatbot that remembers everything you said in this chat.
chat-summary - a chatbot that remembers a short summary when the chat gets long.
tools-react - a bot that can use tools (like a calculator) by writing them out as text.
tools-native - the same idea, but it calls tools the "official" way the model supports.
rag-numpy - answers questions by first reading your documents; the simplest version.
rag-chroma - the same, but it saves what it read so it does not re-read every time.
rag-llamaindex / rag-haystack - the same idea built with the LlamaIndex / Haystack libraries.
rag-llamaindex-chroma - LlamaIndex doing the reading, with the saved store from Chroma.
rag-rerank - a smarter search: it looks two ways and then re-sorts results to pick the best.
hybrid-semantic / hybrid-llm - decides "should I read the docs or just chat?" before answering.
hybrid-adaptive - reads the docs first, then checks "is this useful?" and chats if not.
hybrid-adaptive-plus - the careful version: checks each document, the answer, and retries if needed.
agentic-rag - the bot decides by itself when to search the documents, when to use a tool, or just answer. Searching the documents is given to it as one of its "tools".
rag-eval - not a chat. It is a small test that scores how well each RAG version finds the right document, so you can compare them with numbers.
Inside the REPL: type a message, or use slash commands.
/skills lists available skills (a numbered menu).
/<number> or /<name> loads a skill into the active context.
/add [path] ingests a text file into the RAG store live, without restarting.
With no path (or an invalid one), it scans data/corpus/ and adds any new files
not yet indexed. Works for RAG and hybrid agents; persists for Chroma.
/win prints the full context window currently sent to the model.
/active , /clear , /rules , /help , /exit .
The tool-calling agents ( tools-react , tools-native ) can use tools from any
Model Context Protocol server. Copy the example and edit it:
copy mcp.json.example mcp.json
Each entry is either a stdio server ( command + args ) or an HTTP server
( url + transport ). On startup the agents connect, load the servers' tools,
and merge them with the built-ins (built-ins win on name clashes). If mcp.json
is absent or a server fails, the agents simply run with the built-in tools.
Set MCP_CONFIG to point at a different config file if needed.
SOUL.md (project root) is a global persona/identity prepended to every agent's
system prompt. Every agent accepts --no-soul to skip loading it.
Agents can write durable facts into the managed ## Memory section of SOUL.md
(between memory:start / memory:end markers, append-only + deduped): use the
/remember <fact> command in any agent, or the remember tool in the
tool-calling agents. New memories apply on the next turn (no restart).
rules/ holds always-on markdown that is injected into every agent's system
prompt automatically.
skills/ holds opt-in markdown skills loaded on demand via the / menu. Drop a
new .md file in skills/ (optionally with name: / description: frontmatter)
and it appears in the menu.
All four RAG backends implement one Retriever interface ( index / search )
and return the same RetrievedChunk(text, source, score) , so the agents are
identical regardless of the engine behind them. The flow:
Load every file in data/corpus/ and split it into chunks.
Send the chunks to Ollama mxbai-embed-large to get one vector per chunk.
Store the vectors in the backend (numpy matrix, Chroma, LlamaIndex, Haystack).
Embed the query with the same model.
Score every chunk by cosine similarity and take the top-k (default 4).
If nothing is retrieved, the agent replies "I don't know".
Otherwise it builds Context: [source] chunk... / Question: ... and sends it
to gemma4:e4b with a system rule: answer only from the context and cite
sources as [source] .
Retrieval hits (sources + scores) and the LLM call are logged.
Two defining traits: the RAG-only agents are stateless (each query is
independent, no conversation memory), and embeddings come from a different model
( mxbai-embed-large ) than generation ( gemma4:e4b ).
pure-numpy ( rag-numpy ): the most transparent. Chunks by character window
with overlap, L2-normalizes vectors, and does a brute-force dot product
( matrix @ query_vec ) for cosine similarity. No database, O(N) per query - great
for understanding, poor for scale.
pure-chroma ( rag-chroma ): same idea but vectors live in a persistent
Chroma collection ( .chroma/ ) that performs the nearest-neighbor search. It is
the only backend that survives restarts: if the collection already holds data it
is reused instead of re-indexed.
llamaindex ( rag-llamaindex ): uses LlamaIndex's VectorStoreIndex with a
sentence splitter and Ollama embeddings/LLM, kept in memory.
llamaindex-chroma ( rag-llamaindex-chroma ): same LlamaIndex orchestration,
but with a persistent Chroma vector store behind it. This shows the common
production layering - a framework owns the pipeline while an external, scalable
store owns the vectors. Contrast it with rag-chroma , where our own code does
the orchestration and Chroma is driven directly.
haystack ( rag-haystack ): uses Haystack components (document splitter,
in-memory store, embedding retriever) wired into the same interface.
Because the chunking differs (numpy/chroma split by characters; LlamaIndex and
Haystack split by sentences/words), the same question can retrieve slightly
different passages across backends - a good thing to compare.
rerank ( rag-rerank ): a production-style retrieval stack on a small local
scale. It runs hybrid retrieval - dense (Ollama embeddings, cosine) plus
sparse (BM25 keyword) - fuses the two rankings with Reciprocal Rank Fusion, then
reranks the candidates with a cross-encoder ( cross-encoder/ms-marco-MiniLM-L-6-v2
by default) and keeps the top-k. The cross-encoder gives much sharper relevance
separation than raw cosine (clearly positive scores for the right chunk, strongly
negative for the rest). The model downloads on first run; configure it via
RERANK_MODEL and the candidate pool via FUSION_CANDIDATES .
Every RAG agent (and every hybrid agent) accepts:
--no-index - start without indexing/uploading the corpus. The in-memory
backends start empty; Chroma uses whatever is already persisted.
--drop - drop persisted RAG storage before starting (Chroma only; a no-op for
the in-memory backends), forcing a clean rebuild.
By default the in-memory backends index the corpus on every run, while Chroma
reuses its persisted collection if present.
uv run rag-chroma --drop
uv run rag-chroma --no-index
Agentic RAG
agentic-rag is a different take on combining tools and retrieval: instead of a
router deciding RAG-vs-chat up front, retrieval is exposed to a native
tool-calling agent as a search_knowledge_base tool (backed by the hybrid
BM25+dense+rerank engine). The model decides per turn whether to search the
corpus, use another tool (calculator, shell, web_search, MCP), or answer
directly. It is instructed to prefer the knowledge base for factual questions.
uv run agentic-rag
Production features
Human-in-the-loop : dangerous tools (currently shell ) require confirmation
before running. Controlled by REQUIRE_TOOL_A

[truncated]
