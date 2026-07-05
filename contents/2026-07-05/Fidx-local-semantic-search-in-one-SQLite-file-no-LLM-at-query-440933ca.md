---
source: "https://github.com/williamliu-ai/fidx"
hn_url: "https://news.ycombinator.com/item?id=48791740"
title: "Fidx – local semantic search in one SQLite file, no LLM at query"
article_title: "GitHub - williamliu-ai/fidx: Local AI search engine for your files and agents — hybrid BM25 + vector semantic search (RAG) in one SQLite file. Millisecond queries, CPU-only, offline, no LLM in the query path. · GitHub"
author: "williamliu_ai"
captured_at: "2026-07-05T06:30:40Z"
capture_tool: "hn-digest"
hn_id: 48791740
score: 1
comments: 0
posted_at: "2026-07-05T06:27:15Z"
tags:
  - hacker-news
  - translated
---

# Fidx – local semantic search in one SQLite file, no LLM at query

- HN: [48791740](https://news.ycombinator.com/item?id=48791740)
- Source: [github.com](https://github.com/williamliu-ai/fidx)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T06:27:15Z

## Translation

タイトル: Fidx – 1 つの SQLite ファイルでのローカル セマンティック検索、クエリ時に LLM なし
記事のタイトル: GitHub - williamliu-ai/fidx: ファイルとエージェント用のローカル AI 検索エンジン — 1 つの SQLite ファイル内のハイブリッド BM25 + ベクトル セマンティック検索 (RAG)。ミリ秒クエリ、CPU のみ、オフライン、クエリ パスに LLM なし。 · GitHub
説明: ファイルとエージェント用のローカル AI 検索エンジン — 1 つの SQLite ファイル内のハイブリッド BM25 + ベクトル セマンティック検索 (RAG)。ミリ秒クエリ、CPU のみ、オフライン、クエリ パスに LLM なし。 - ウィリアムリューアイ/fidx

記事本文:
GitHub - williamliu-ai/fidx: ファイルとエージェント用のローカル AI 検索エンジン — 1 つの SQLite ファイル内のハイブリッド BM25 + ベクトル セマンティック検索 (RAG)。ミリ秒クエリ、CPU のみ、オフライン、クエリ パスに LLM なし。 · GitHub
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
ウィリアムリウアイ
/
フィックス

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット .github .github .well-known .well-known ライセンス ライセンス ベンチ ベンチ docker docker docs docs スクリプト スクリプト src/ fidx src/ fidx テスト テスト .dockerignore .dockerignore .gitignore .gitignore .python-version .python-version AI_ATTRIBUTION.md AI_ATTRIBUTION.md ATTRIBUTIONS.md ATTRIBUTIONS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md SOURCE_HEADER_EXAMPLES.md SOURCE_HEADER_EXAMPLES.md ai-attribution-policy.json ai-attribution-policy.json codemeta.json codemeta.json pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ファイルとエージェント用のローカル AI 検索エンジン。
マークダウン、テキスト、チャットのエクスポートのための、ローカルのみの高速なセマンティック + キーワード検索
そしてコード。 CPUのみ。クラウドも GPU も API キーもありません。 1 つの SQLite ファイルには、
完全なインデックス。 CLI、プライベート RAG リトリーバー、または検索レイヤーとして使用します。
エージェント メモリの背後: ミリ秒のウォーム クエリ、JSON 出力、およびコレクション
スコーピング。
デモ コーパスは、ベンチマークの合成チャット データです (
一緒にリリースする
ドキュメント コーパスとクエリ セットを使用);そこにあるdemo-driver.shは、
録音中。
ローカル セマンティック検索ツールは、遅延を伴うリコールを購入する傾向があります: LLM クエリ
拡張と LLM 再ランキングにより、単一のクエリが CPU で最大 10 秒かかります。フィックス
別の取引を行う — ハイブリッド BM25 + 768 次元ベクトル検索と融合
相互ランク融合 (RRF)、クエリ パス内に LLM 呼び出しはありません。 1 つの ONNX
クエリごとの埋め込みパスが唯一のモデル作業です。

ハイブリッド リコール — FTS5 BM25 は正確な名前と識別子を捕捉します。 768 次元
埋め込みは「インデックス作成プロジェクトについて議論したドキュメント」をキャッチします。 RRF は両方を融合します。
ミリ秒クラスのクエリ — ウォーム デーモンがハイブリッド検索に応答します。
2k ～ 19k-doc コーパスで 18 ～ 49 ミリ秒 (p50、シングル ONNX スレッド)。コールド CLI 呼び出し
1秒未満で十分に滞在してください。
1 つのファイル — ドキュメント、BM25 インデックス、およびベクターが単一の SQLite 内に存在します
データベース (FTS5 + sqlite-vec )。コピーし、バックアップし、削除します。
範囲指定された検索 — ソースを名前付きコレクションにグループ化します ( -c email )
言いたいことだけを検索してください。
Python 3.11 または 3.12 (その sqlite3 は読み込み可能な拡張機能をサポートしており、
FTS5 (fidx は sqlite-vec 拡張機能をロードします)。 fidx Doctor を実行して確認します。
以下の検証済みプラットフォーム用に事前構築されたホイールが存在します。コンパイラは必要ありません。
最初の fidx インデックスは、埋め込みモデルを 1 回ダウンロードします (その後、完全にオフラインになります)。
パッケージの名前が fmdidx なのはなぜですか?プロジェクトとそのコマンドは、
fidx ですが、PyPI 名 fidx はすでに無関係な名前によって使用されています
パッケージ — したがって、fidx は fmdidx (「高速マークダウン インデックス」) として配布されます。
名前が異なるのはそこだけです: uv tools install fmdidx installs
fidxコマンド。
推奨されるインストーラーは uv です。
uv マネージド Python には、Linux および Windows 上でロード可能な sqlite 拡張機能が同梱されています。
uv ツール install fmdidx # または: pipx install fmdidx
fidx Doctor # ホストを確認する
macOS: uv にバンドルされている Python (および python.org ビルド) には sqlite3 が同梱されています
ロード可能な拡張機能のサポートがないため、 Homebrew Python を使用します。
Pythonを醸造インストールする
uv ツールのインストール --python " $( brew --prefix python ) /libexec/bin/python " fmdidx
# または: pipx install --python "$(brew --prefix python)/libexec/bin/python3" fmdidx
フィックスドクター
構築されたホイールから (現在動作し、名前は関係ありません):
UVビルド
pip install --only-binary=:all: dist/ * .whl # ma で Homebrew Python を使用する

コス
フィックスドクター
チェックアウト (開発) から:
UV 同期 && UV 実行 FIDX ドクター
fidx Doctor が失敗を報告すると、不足しているものとその方法を正確に出力します。
修正してください。「トラブルシューティング」を参照してください。
# ディレクトリを名前付きコレクションとして登録します
fidx コレクション add ~ /notes --name メモ
fidx コレクション add ~ /mail/export --name メール --glob " **/*.txt "
# スキャン + チャンク + 埋め込み (増分; 最初の実行で ONNX モデルをダウンロード)
フィックスインデックス
# 検索 (デフォルトではハイブリッド BM25 + ベクトル)
fidx search "最新のインデックス作成プロジェクトについて説明したドキュメント"
fidx search " Grace Hopper " --mode lexical # 正確な名前の検索、モデルのロードなし
fidx search " デプロイメント チェックリスト " -c ノート # スコープは 1 つのコレクション
# エージェントフレンドリーな出力
fidx 検索 " エラー処理 " --json -n 10
fidx 検索 " auth " --files --min-score 0.02
# パスまたは docid でドキュメントを取得します
fidx は「notes/meeting.md」を取得します
fidx 取得 " #a1b2c3 "
ウォーム デーモン (エージェントに推奨)
fidxserve & # UNIX ソケット上でモデル + インデックスをホットに保ちます
fidx search " ... " # すべての検索にミリ秒かかるようになりました
CLI は、実行中にデーモンを自動的に使用します。 --no-daemon はオプトアウトします。
fidx はエージェント フレームワークではありません。その統合面は CLI です: register
ローカル ファイルを作成し、インデックスを作成し、デーモンをウォーム状態に保ち、エージェントまたはワークフローに任せます。
fidx search --json と fidx get を呼び出します。
# メモ、メモリ エクスポート、ドキュメント、またはコードを個別の検索可能な範囲としてインデックス付けする
fidx コレクション add ./memory --name メモリ \
--glob " **/*.md " --glob " **/*.txt " --glob " **/*.jsonl "
fidx コレクション add ./docs --name docs --glob " **/*.md " --glob " **/*.txt "
fidx コレクションの追加 ./src --name code \
--glob " **/*.py " --glob " **/*.ts " --glob " **/*.go " --glob " **/*.rs "
フィックスインデックス
フィックスサーブ&
# エージェント メモリ ルックアップ: ツール呼び出し用の構造化 JSON
fidx検索「どうしたのw」

再試行の処理について決定しますか? " -c メモリ --json -n 5
# ローカル RAG 取得: ドキュメントを検索し、選択したソース テキストを取得します。
fidx 検索 " sqlite ベクター バックエンド構成 " -c docs --json -n 5
fidx get --head " #a1b2c3 "
# コーディング エージェント コンテキスト: フォローアップ読み取りに一致するパスのみを返します
fidx search " リクエストのタイムアウトはどこで処理されますか? " -c コード --files -n 20
MCP サーバー、LangChain/LangGraph ツール、LlamaIndex レトリーバー、ワークフロー用
ノード、またはシェルベースのコーディング エージェントの場合、通常は同じコントラクトで十分です。
fidx search --json は、スニペット、パス、ドキュメント ID を含むランク付けされた結果を返します。
スコア; fidx get は、選択された結果を docid または collection/path によって展開します。
fidx ドクター # ホスト機能レポート (終了 0 = 準備完了)
# インストールされている CLI に対する、約 1,000 件のドキュメント コーパスの完全なエンドツーエンド ベンチマーク:
python scripts/e2e_smoke.py # コーパスの構築、インデックス、検索、ゲートリコール
# 初期状態の Docker コンテナーでのクリーンマシンの証明 (Linux):
scripts/verify-install.sh # ホイールを構築し、インストールして 3.11 および 3.12 で e2e を実行します
同じ e2e が、Linux、macOS (arm64)、Windows × Python 3.11/3.12 上の CI で実行されます。
(インストール マトリックス ワークフロー) — 構築されたホイールを最初からインストールし、
リコール@10 をアサートします。
ファイル ──> ドキュメント (SQLite) ──> FTS5 (BM25、ポーター) ──┐
│ §─> RRF 融合 ─> 結果
━──> チャンク ──> ONNX 埋め込み ──> sqlite-vec ┘
チャンク分割は最適な構造上のブレーク (見出し、コードフェンス) で分割されます。
境界、空白行）1800 文字以内のターゲット付近で 15% オーバーラップ、なし
コードフェンスの中。チャンクはコピーではなくオフセットを保存します。
fastembed /ONNX (CPU) による埋め込み。デフォルトのプロファイルは 768-dim です。
小さなコーパスには、より小さなプロファイルが存在します。
検索は BM25 とベクトル KNN を並行して実行し、RRF と融合します
(k=60);結果は、最良のチャンク スニペットを含むドキュメント レベルです。
エナ

ble_load_extension / "sqlite3 はロード可能な拡張機能なしでビルドされました
support" — Python の sqlite は sqlite-vec をロードできません。これがデフォルトです
macOS システムの Python および uv/python.org macOS ビルド。修正: fidx をインストールする
Homebrew Python を使用します (上記の macOS のインストールを参照)。 fidxの医師が確認した
修正します。
sqlite-vec のロードに失敗したか、アーキテクチャが間違っています - sqlite-vec を確認してください
あなたのプラットフォーム用のホイールが存在します: pip install --only-binary=:all: sqlite-vec 。
最初の検索が遅い/オフラインで使用 - 埋め込みモデルは一度ダウンロードされる
最初のインデックス。 fidx エアギャップを使用するには、FASTEMBED_CACHE_PATH を事前にシードします。
bench/ は、fidx と QMD を 4 つ比較する再現可能なハーネスです
既知の項目クエリを含むコーパス (CPU のみ、ウォーム エンジン、アイドル ボックス)。結果
品質には、トレードオフとなる 2 つの軸があります: 再現率 (これは、
トップ 10) と純度 — ノイズ @10 (LLM から返された結果のシェア)
審査員は無関係と評価、低いほど良い）およびクリーン@10（クエリの割合）
結果にはノイズが含まれておらず、高いほど優れています)。そして問い合わせも入ってくる
2 つの体制: 文書自体の言葉から構築された既知の項目のクエリ
(ほとんどの検索: 名前、識別子、記憶されたフレーズ) および言い換え
ターゲットと語彙をほとんど共有しないクエリ (純粋な意味論)
思い出してください）。既知の項目の見出し: vs QMD の LLM ハイブリッド、fidx が両方に勝利
すべてのコーパスの純度メトリクスとドキュメントとコードのリコール - リコールの関係
docs-small であり、チャットでは 0.004 遅れています。つまり、
doc/chat コーパスと 92k ファイル コード コーパスで約 65 倍。
fidx 行は、組み込みの決定論的な結果の切り捨てを使用して測定されます。
有効 ( --truncate 膝 ; デフォルトで出荷されます。これなしでは fidx trade が行われます)
再現のための純度、例:コード R@10 0.900、ノイズ 0.297) および e5-768
profile ( fidxindex --profile e5-768 ; インストールのデフォルトは次のとおりです
のみ

c-768-q 、コード上は同じスコアでした — BENCHMARKS.md を参照してください)。
「n/m」 = 測定されません。
ハイブリッド vs ハイブリッド (fidx vs QMD クエリ): fidx がノイズ @10 で勝利し、
すべてのコーパスで clean@10 を実行し、ドキュメントとコードでリコールします (docs-small に関連付けます)。
チャットでは 0.908 対 0.912) — 例:コード再現率 +8 ポイント、ノイズが 5.6 倍減少 —
クエリ パスのどこにも LLM がありません。
QMD が勝てる場所: その FTS 検索モードはチャットにおける純粋なチャンピオンです
(クリーン 0.964) およびビッグ コード コーパスのレイテンシ チャンピオン (121 ミリ秒 vs
fidx の 92,000 ベクトルにわたる 452 ミリ秒のブルートフォース KNN)、しかし再現では劣る
重要な場所 (ドキュメント、コード)。 QMD の純粋ベクトル モードは 0.048 に崩壊します
コード上の R@10。
fidx は最も弱いコーパスでも 1 秒未満にとどまり、fidx よりも約 65 倍高速です
QMD の LLM モードがあります。
セマンティック体制 (言い換えクエリ)。独立した査読者の後
既知の項目のクエリが語彙の重複をもたらすことを正しく認識したため、
LLM で作成され、個別に LLM で検証された言い換えクエリ セット (~0.1 クエリ→ドキュメント)
単語の重複。 bench/data/ にチェックインされました - 同じターゲット、コピーされていません
特徴的な用語。 @10 を思い出してください:
正直な見方: BM25 は、共通の条件がなければ文字通りゼロになります。フィックスの
ベクトルアームはミリ秒のレイテンシで散文に対して実際の意味論的な作業を行います。 QMD
LLM 拡張は、測定された場所で最高の意味的再現を実現します (~33 秒)。
クエリごとに 20 ミリ秒。フィ

[切り捨てられた]

## Original Extract

Local AI search engine for your files and agents — hybrid BM25 + vector semantic search (RAG) in one SQLite file. Millisecond queries, CPU-only, offline, no LLM in the query path. - williamliu-ai/fidx

GitHub - williamliu-ai/fidx: Local AI search engine for your files and agents — hybrid BM25 + vector semantic search (RAG) in one SQLite file. Millisecond queries, CPU-only, offline, no LLM in the query path. · GitHub
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
williamliu-ai
/
fidx
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits .github .github .well-known .well-known LICENSES LICENSES bench bench docker docker docs docs scripts scripts src/ fidx src/ fidx tests tests .dockerignore .dockerignore .gitignore .gitignore .python-version .python-version AI_ATTRIBUTION.md AI_ATTRIBUTION.md ATTRIBUTIONS.md ATTRIBUTIONS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SOURCE_HEADER_EXAMPLES.md SOURCE_HEADER_EXAMPLES.md ai-attribution-policy.json ai-attribution-policy.json codemeta.json codemeta.json pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Local AI search engine for your files and agents.
Fast, local-only semantic + keyword search for markdown, text, chat exports
and code. CPU-only. No cloud, no GPU, no API keys. One SQLite file holds the
full index. Use it as a CLI, a private RAG retriever, or the search layer
behind agent memory: millisecond warm queries, JSON output, and collection
scoping.
The demo corpus is the benchmark's synthetic chat data ( attached to the
release along
with the doc corpora and query sets); demo-driver.sh there reproduces the
recording.
Local semantic search tools tend to buy recall with latency: LLM query
expansion and LLM reranking push a single query to ~10 seconds on CPU. fidx
takes a different trade — hybrid BM25 + 768-dim vector search fused with
reciprocal-rank fusion (RRF), and no LLM calls in the query path . One ONNX
embedding pass per query is the only model work.
Hybrid recall — FTS5 BM25 catches exact names and identifiers; 768-dim
embeddings catch "that doc that discussed the indexing project"; RRF fuses both.
Millisecond-class queries — a warm daemon answers hybrid searches in
18–49 ms (p50, single ONNX thread) on 2k–19k-doc corpora; cold CLI calls
stay well under a second.
One file — documents, BM25 index and vectors live in a single SQLite
database (FTS5 + sqlite-vec ). Copy it, back it up, delete it.
Scoped search — group sources into named collections ( -c emails )
and search only what you mean.
Python 3.11 or 3.12 whose sqlite3 supports loadable extensions and
FTS5 (fidx loads the sqlite-vec extension). Run fidx doctor to verify.
Prebuilt wheels exist for the verified platforms below — no compiler needed .
First fidx index downloads the embedding model once (then fully offline).
Why is the package named fmdidx ? The project and its command are
fidx , but the PyPI name fidx was already taken by an unrelated
package — so fidx is distributed as fmdidx ("fast markdown index").
That is the only place the name differs: uv tool install fmdidx installs
the fidx command.
The recommended installer is uv , because a
uv-managed Python ships loadable sqlite extensions on Linux and Windows.
uv tool install fmdidx # or: pipx install fmdidx
fidx doctor # verify your host
macOS: uv's bundled Python (and the python.org build) ship a sqlite3
without loadable-extension support, so use Homebrew Python :
brew install python
uv tool install --python " $( brew --prefix python ) /libexec/bin/python " fmdidx
# or: pipx install --python "$(brew --prefix python)/libexec/bin/python3" fmdidx
fidx doctor
From a built wheel (works today, name-independent):
uv build
pip install --only-binary=:all: dist/ * .whl # use Homebrew Python on macOS
fidx doctor
From a checkout (development):
uv sync && uv run fidx doctor
If fidx doctor reports a failure, it prints exactly what is missing and how to
fix it — see Troubleshooting .
# Register directories as named collections
fidx collection add ~ /notes --name notes
fidx collection add ~ /mail/export --name emails --glob " **/*.txt "
# Scan + chunk + embed (incremental; first run downloads the ONNX model)
fidx index
# Search (hybrid BM25 + vector by default)
fidx search " the doc that discussed the most recent indexing project "
fidx search " Grace Hopper " --mode lexical # exact-name lookup, no model load
fidx search " deployment checklist " -c notes # scope to one collection
# Agent-friendly output
fidx search " error handling " --json -n 10
fidx search " auth " --files --min-score 0.02
# Fetch a document by path or docid
fidx get " notes/meeting.md "
fidx get " #a1b2c3 "
Warm daemon (recommended for agents)
fidx serve & # keeps the model + index hot on a unix socket
fidx search " ... " # all searches now take milliseconds
The CLI uses the daemon automatically when it is running; --no-daemon opts out.
fidx is not an agent framework. Its integration surface is the CLI: register
local files, index them, keep the daemon warm, then let an agent or workflow
call fidx search --json and fidx get .
# Index notes, memory exports, docs, or code as separate searchable scopes
fidx collection add ./memory --name memory \
--glob " **/*.md " --glob " **/*.txt " --glob " **/*.jsonl "
fidx collection add ./docs --name docs --glob " **/*.md " --glob " **/*.txt "
fidx collection add ./src --name code \
--glob " **/*.py " --glob " **/*.ts " --glob " **/*.go " --glob " **/*.rs "
fidx index
fidx serve &
# Agent-memory lookup: structured JSON for a tool call
fidx search " what did we decide about retry handling? " -c memory --json -n 5
# Local RAG retrieval: search docs, then fetch the selected source text
fidx search " sqlite vector backend configuration " -c docs --json -n 5
fidx get --head " #a1b2c3 "
# Coding-agent context: return only matching paths for follow-up reads
fidx search " where is request timeout handled? " -c code --files -n 20
For MCP servers, LangChain/LangGraph tools, LlamaIndex retrievers, workflow
nodes, or shell-based coding agents, the same contract is usually enough:
fidx search --json returns ranked results with snippets, paths, docids and
scores; fidx get expands a selected result by docid or collection/path .
fidx doctor # host capability report (exit 0 = ready)
# Full end-to-end benchmark on a ~1,000-doc corpus against the installed CLI:
python scripts/e2e_smoke.py # builds corpus, indexes, searches, gates recall
# Clean-machine proof in pristine Docker containers (Linux):
scripts/verify-install.sh # builds the wheel, installs + runs e2e on 3.11 & 3.12
The same e2e runs in CI on Linux, macOS (arm64) and Windows × Python 3.11/3.12
(the install-matrix workflow) — installing the built wheel from scratch and
asserting recall@10 .
files ──> documents (SQLite) ──> FTS5 (BM25, porter) ─┐
│ ├─> RRF fusion ─> results
└──> chunks ──> ONNX embeddings ─> sqlite-vec ┘
Chunking splits at the best structural break (headings, code-fence
boundaries, blank lines) near a ~1800-char target with 15% overlap, never
inside a code fence. Chunks store offsets, not copies.
Embeddings via fastembed /ONNX (CPU). The default profile is 768-dim;
smaller profiles exist for small corpora.
Search runs BM25 and vector KNN in parallel and fuses with RRF
(k=60); results are document-level with best-chunk snippets.
enable_load_extension / "sqlite3 was built without loadable-extension
support" — your Python's sqlite cannot load sqlite-vec . This is the default
on macOS system Python and uv/python.org macOS builds. Fix: install fidx
with Homebrew Python (see macOS install above). fidx doctor confirms the
fix.
sqlite-vec failed to load / wrong architecture — ensure a sqlite-vec
wheel exists for your platform: pip install --only-binary=:all: sqlite-vec .
First search is slow / offline use — the embedding model downloads once on
first index. Pre-seed FASTEMBED_CACHE_PATH to use fidx air-gapped.
bench/ is a reproducible harness comparing fidx against QMD on four
corpora with known-item queries (CPU-only, warm engines, idle box). Result
quality has two axes that trade off : recall (is the right document in
the top-10) and purity — noise@10 (share of returned results an LLM
judge rated irrelevant, lower is better) and clean@10 (share of queries
whose results contain zero noise, higher is better). And queries come in
two regimes : known-item queries built from the document's own words
(most lookups: names, identifiers, remembered phrases) and paraphrase
queries that share almost no vocabulary with the target (pure semantic
recall). The known-item headline: vs QMD's LLM hybrid, fidx wins both
purity metrics on every corpus and recall on docs and code — recall ties on
docs-small and is 0.004 behind on chat — at ~300–1000× lower latency on the
doc/chat corpora and ~65× on the 92k-file code corpus.
fidx rows are measured with its built-in deterministic result truncation
enabled ( --truncate knee ; ships off by default — without it fidx trades
purity for recall, e.g. code R@10 0.900 at noise 0.297) and the e5-768
profile ( fidx index --profile e5-768 ; the install default is
nomic-768-q , which scored identically on code — see BENCHMARKS.md).
"n/m" = not measured.
Hybrid vs hybrid (fidx vs QMD query ): fidx wins noise@10 and
clean@10 on every corpus, and recall on docs and code (tie on docs-small;
0.908 vs 0.912 on chat) — e.g. code recall +8 pts with 5.6× less noise —
with no LLM anywhere in its query path.
Where QMD wins: its FTS search mode is the purity champion on chat
(clean 0.964) and the latency champion on the big code corpus (121 ms vs
fidx's 452 ms brute-force KNN over 92k vectors), but trails on recall
where it matters (docs, code). QMD's pure-vector mode collapses to 0.048
R@10 on code.
fidx stays sub-second even on its weakest corpus and is ~65× faster than
QMD's LLM modes there.
The semantic regime (paraphrase queries). After an independent reviewer
correctly noted that known-item queries reward lexical overlap, we built an
LLM-written, separately-LLM-validated paraphrase query set (~0.1 query→doc
word overlap; checked into bench/data/ ) — same targets, no copied
distinctive terms. Recall@10 there:
Honest readings: BM25 gets literally zero without shared terms; fidx's
vector arm does real semantic work on prose at millisecond latency; QMD's
LLM expansion buys the best semantic recall where measured — at ~33 s vs
20 ms per query; fi

[truncated]
