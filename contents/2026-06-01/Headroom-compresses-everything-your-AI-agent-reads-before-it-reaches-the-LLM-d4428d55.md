---
source: "https://pypi.org/project/headroom-ai/"
hn_url: "https://news.ycombinator.com/item?id=48349275"
title: "Headroom compresses everything your AI agent reads before it reaches the LLM"
article_title: "headroom-ai · PyPI"
author: "mooreds"
captured_at: "2026-06-01T00:25:26Z"
capture_tool: "hn-digest"
hn_id: 48349275
score: 2
comments: 1
posted_at: "2026-05-31T20:10:44Z"
tags:
  - hacker-news
  - translated
---

# Headroom compresses everything your AI agent reads before it reaches the LLM

- HN: [48349275](https://news.ycombinator.com/item?id=48349275)
- Source: [pypi.org](https://pypi.org/project/headroom-ai/)
- Score: 2
- Comments: 1
- Posted: 2026-05-31T20:10:44Z

## Translation

タイトル: ヘッドルームは、AI エージェントが読み取るすべてのものを LLM に到達する前に圧縮します。
記事タイトル: headroom-ai・PyPI
説明: LLM アプリケーションのコンテキスト最適化レイヤー - コストを 50 ～ 90% 削減

記事本文:
メインコンテンツにスキップ
モバイル版に切り替える
警告
サポートされていないブラウザを使用しているため、新しいバージョンにアップグレードしてください。
PyPIを検索
検索フォーカス#フォーカス検索フィールド"
data-search-focus-target="検索フィールド">
検索
ヘルプ
pip インストール ヘッドルーム-ai
PIP 命令のコピー
LLM アプリケーションのコンテキスト最適化レイヤー - コストを 50 ～ 90% 削減
ライセンス式: Apache-2.0
SPDX
ライセンスの表現
メンテナー: ヘッドルームの貢献者
タグ
llm
、
オープンナイ
、
人間的な
、
クロード
、
GPT
、
コンテキスト
、
トークン
、
最適化
、
圧縮
、
キャッシュ
、
プロキシ
、
あい
、
機械学習
追加の提供:
アグノ
、すべて
、いずれも
、岩盤
、ベンチマーク
、コード
、開発者
、評価
、html
、画像
、ラングチェーン
、MCP
、記憶
、メモリスタック
、ml
、オテル
、代理
、関連性
、レポート
、ストランド
、声
、ボイストレーニング
ライセンス
OSI 承認済み :: Apache ソフトウェア ライセンス
オペレーティングシステム
OSに依存しない
プログラミング言語
パイソン :: 3
トピック
科学/工学 :: 人工知能
ソフトウェア開発 :: ライブラリ :: Python モジュール
██╗ ██╗███████╗ █████╗ ██████╗ ██████╗ ██████╗ ██████╗ ███╗ ███╗
██║ ██║██╔════╝██╔══██╗██╔══██╗██ ╔══██╗██╔═══██╗██╔═══██╗████╗ ████║
███████║█████╗ ███████║██║ ██║██████╔╝██║ ██║██║ ██║██╔████╔██║
██╔══██║██╔══╝ ██╔══██║██║ ██║██╔══██╗██║ ██║██║ ██║██║╚██╔╝██║
██║ ██║███████╗██║

██║██████╔╝██║ ██║╚██████╔╝╚██████╔╝██║ ╚═╝ ██║
╚═╝ ╚═╝╚══════╝╚═╝ ╚═╝╚═════╝ ╚═╝ ╚═╝ ╚═════╝ ╚═════╝ ╚═╝ ╚═╝
AI エージェントのコンテキスト圧縮レイヤー
60 ～ 95% 少ないトークン · ライブラリ · プロキシ · MCP · 6 つのアルゴリズム · ローカルファースト · 可逆
ドキュメント ·
インストール・
証明・
エージェント ·
不協和音 ·
llms.txt
AI エージェント/LLM: ここで /llms.txt を読み取るか、ライブ インデックス/完全なドキュメント BLOB を取得します。
ヘッドルームは、AI エージェントが読み取るすべてのもの (ツール出力、ログ、RAG チャンク、ファイル、会話履歴) を LLM に到達する前に圧縮します。同じ答え、トークンの一部。
ライブ: 10,144 → 1,260 トークン — 同じ FATAL が見つかりました。
ライブラリ — Python または TypeScript で圧縮(メッセージ)、任意のアプリでインライン
プロキシ — ヘッドルーム プロキシ -- ポート 8787 、コード変更なし、任意の言語
エージェント ラップ — 1 つのコマンドでヘッドルームをラップ クロード|コーデックス|カーソル|エイダー|コパイロット
MCP サーバー - 任意の MCP クライアントの headroom_compress 、 headroom_retrieve 、 headroom_stats
クロスエージェント メモリ — クロード、コーデックス、ジェミニにわたる共有ストア、自動重複化
ヘッドルーム学習 — 失敗したセッションをマイニングし、修正を CLAUDE.md / AGENTS.md に書き込みます
可逆 (CCR) — オリジナルは削除されません。 LLM はオンデマンドで取得します
エージェント/アプリ
(クロード コード、カーソル、コーデックス、LangChain、Agno、Strands、独自のコード…)
│ プロンプト、ツール出力、ログ、RAG 結果、ファイル
▼
┌───────────────────────┐
│ ヘッドルーム (ローカルで実行 - データはここに残ります) │
│ ──────────

───────────── │
│ CacheAligner → ContentRouter → CCR │
│ §─ SmartCrusher (JSON) │
│ §─ コードコンプレッサー (AST) │
│ └─ Kompress ベース (テキスト、HF) │
│ │
│ エージェント間の記憶・ヘッドルーム学習・MCP │
━━━━━━━━━━━━━━━━━━━━━┘
│ 圧縮プロンプト + 検索ツール
▼
LLM プロバイダー (Anthropic · OpenAI · Bedrock · …)
ContentRouter — コンテンツ タイプを検出し、適切なコンプレッサーを選択します
SmartCrusher / CodeCompressor / Kompress-base — JSON、AST、または散文を圧縮します
CacheAligner — プロバイダーの KV キャッシュが実際にヒットするようにプレフィックスを安定化します。
CCR — オリジナルをローカルに保存します。 LLM は必要に応じて headroom_retrieve を呼び出します
→ アーキテクチャ ・ CCR 可逆圧縮 ・ Kompress ベースのモデルカード
#1 — インストールする
pip install "headroom-ai[all]" # Python
npm install headroom-ai # ノード / TypeScript
#2 — モードを選択する
headroom Wrap claude # コーディングエージェントをラップする
headroom proxy --port 8787 # ドロップイン プロキシ、コード変更なし
# または: ヘッドルームインポート圧縮 # インラインライブラリから
#3 — 節約効果を確認する
ヘッドルーム統計
詳細な追加機能: [proxy] 、 [mcp] 、 [ml] 、 [agno] 、 [langchain] 、 [evals] 。 Python 3.10以降が必要です。
実際のエージェントのワークロードを節約:
標準ベンチマークで維持される精度:
再現: python -m headroom.evals suite --tier 1 · 完全なベンチマークと方法論
コミュニティによって保存された 600 億以上のトークン — ライブリーダーボード →
OpenAI 互換クライアントはすべて、ヘッドルーム プロキシ経由で動作します。 MCP ネイティブ: ヘッドルーム mcp install 。
AI コーディング エージェントを毎日実行し、コードを変更せずにコストを節約したい
複数のエージェント間で作業し、共有メモリが必要である
必要です

可逆圧縮 - CCR 経由でいつでもオリジナルを取得可能
単一プロバイダーのネイティブ圧縮のみを使用し、エージェント間のメモリを必要としません。
ローカルプロセスを実行できないサンドボックス環境で作業する
SmartCrusher — ユニバーサル JSON: dict の配列、ネストされたオブジェクト、混合型。
CodeCompressor — Python、JS、Go、Rust、Java、C++ の AST 対応。
Kompress-base — エージェント トレースでトレーニングされた HuggingFace モデル。
画像圧縮 — トレーニングされた ML ルーターにより 40 ～ 90% 削減。
CacheAligner — Anthropic/OpenAI KV キャッシュが実際にヒットするようにプレフィックスを安定させます。
IntelligentContext — 学習された重要性に適合するスコアベースのコンテキスト。
CCR — 可逆圧縮。 LLM はオンデマンドでオリジナルを取得します。
エージェント間メモリ — 共有ストア、エージェントの出所、自動重複排除。
SharedContext — マルチエージェント ワークフロー間で受け渡される圧縮コンテキスト。
headroom learn — Claude、Codex、Gemini のプラグインベースの障害マイニング。
Headroom は、 compress() 、SDK、およびプロキシにわたる 1 つの安定したリクエスト ライフサイクルを公開します。
セットアップ → 開始前 → 開始後 → 入力受信 → 入力キャッシュ → 入力ルーティング → 圧縮入力 → 入力記憶 → 送信前 → 送信後 → 応答受信
変換は、CacheAligner、ContentRouter、SmartCrusher、CodeCompressor、Kompress-base、IntelligentContext / RollingWindow によって処理されます。
パイプライン拡張機能は、 on_pipeline_event(...) を介してライフサイクル ステージを監視またはカスタマイズします。
圧縮フックは、追加の拡張シームとして正規のライフサイクルに沿って配置されます。
プロキシ拡張機能は、引き続き ASGI ミドルウェア、ルート、起動ポリシーのサーバー/アプリ統合の継ぎ目です。
プロバイダーとツール固有の動作はヘッドルーム/プロバイダーの下に存在するため、コア オーケストレーションはライフサイクル、シーケンス、ポリシーに重点を置き続けます。
CLI/ツールスライス: headroom/providers/claude 、co

パイロット、コーデックス、オープンクロー
プロバイダー ランタイム スライス: headroom/providers/claude 、 gemini 、および headroom/providers/registry.py の共有バックエンド/ランタイム ディスパッチ
コア ファイルはオーケストレーション ファーストのままです: Wrap.py 、 client.py 、 cli/proxy.py 、および proxy/server.py は、プロバイダー固有の環境シェーピング、API ターゲットの正規化、バックエンドの選択、およびトランスポート ディスパッチを委任します。
pip install "headroom-ai[all]" # Python、すべて
npm install headroom-ai # TypeScript / ノード
docker pull ghcr.io/chopratejas/headroom:latest
詳細な追加機能: [proxy] 、 [mcp] 、 [ml] (Kompress-base)、 [agno] 、 [langchain] 、 [evals] 。 Python 3.10以降が必要です。
pipxを使用していますか?サポートされているインタープリターを明示的に選択します。
pipx install --python python3.13 "ヘッドルーム-ai[すべて]"
→ インストール ガイド — Docker タグ、永続サービス、PowerShell、devcontainers。
headroom learn — 失敗したセッションをマイニングし、修正を CLAUDE.md / AGENTS.md / GEMINI.md に書き込みます。
Headroom はローカルで実行され、あらゆるコンテンツ タイプをカバーし、あらゆる主要なフレームワークで動作し、可逆的です。
帰属。 Headroom には、シェル出力書き換え用の優れた RTK バイナリ ( git show --short 、scoped ls 、要約インストーラー) が同梱されています。 RTK チームに多大な感謝を申し上げます。彼らのツールは私たちのスタックの最上級の部分であり、Headroom はその下流のすべてを圧縮します。 Headroom は、選択した CLI コンテキスト ツールとして lean-ctx を使用することもできます。ヘッドルーム ラップ ... を実行する前に HEADROOM_CONTEXT_TOOL=lean-ctx を設定します。
git clone https://github.com/chopratejas/headroom.git && cd ヘッドルーム
pip install -e ".[dev]" && pytest
.devcontainer/ 内の Devcontainers (デフォルト + Qdrant および Neo4j のメモリスタック)。 COTRIBUTING.md を参照してください。
ライブリーダーボード — 600 億以上のトークンが保存され、増え続けています。
Discord — 質問、フィードバック、戦争の話。
HuggingFace をベースにした Kompress-base — テキスト圧縮の背後にあるモデル

。
ライセンス式: Apache-2.0
SPDX
ライセンスの表現
メンテナー: ヘッドルームの貢献者
タグ
llm
、
オープンナイ
、
人間的な
、
クロード
、
GPT
、
コンテキスト
、
トークン
、
最適化
、
圧縮
、
キャッシュ
、
プロキシ
、
あい
、
機械学習
追加の提供:
アグノ
、すべて
、いずれも
、岩盤
、ベンチマーク
、コード
、開発者
、評価
、html
、画像
、ラングチェーン
、MCP
、記憶
、メモリスタック
、ml
、オテル
、代理
、関連性
、レポート
、ストランド
、声
、ボイストレーニング
ライセンス
OSI 承認済み :: Apache ソフトウェア ライセンス
オペレーティングシステム
OSに依存しない
プログラミング言語
パイソン :: 3
トピック
科学/工学 :: 人工知能
ソフトウェア開発 :: ライブラリ :: Python モジュール
発売履歴
リリースのお知らせ |
RSSフィード
ご使用のプラットフォーム用のファイルをダウンロードします。どれを選択すればよいかわからない場合は、パッケージのインストールについて詳しくご覧ください。
名前、インタープリター、ABI、プラットフォームでファイルをフィルターします。
ファイル名の形式がわからない場合は、 ホイール ファイル名 について学習してください。
現在のフィルターへの直接リンクをコピーします。
コピー
アップロードされました
2026 年 5 月 21 日
CPython 3.13 多くの Linux: glibc 2.28+ ARM64
アップロードされました
2026 年 5 月 21 日
CPython 3.13 macOS 11.0+ ARM64
アップロードされました
2026 年 5 月 21 日
CPython 3.12 manylinux: glibc 2.28+ x86-64
アップロードされました
2026 年 5 月 21 日
CPython 3.12 manylinux: glibc 2.28+ ARM64
アップロードされました
2026 年 5 月 21 日
CPython 3.12 macOS 11.0+ ARM64
アップロードされました
2026 年 5 月 21 日
CPython 3.11 manylinux: glibc 2.28+ x86-64
アップロードされました
2026 年 5 月 21 日
CPython 3.11 manylinux: glibc 2.28+ ARM64
アップロードされました
2026 年 5 月 21 日
CPython 3.11 macOS 11.0+ ARM64
アップロードされました
2026 年 5 月 21 日
CPython 3.10 manylinux: glibc 2.28+ x86-64
アップロードされました
2026 年 5 月 21 日
CPython 3.10 manylinux: glibc 2.28+ ARM64
アップロードされました
2026 年 5 月 21 日
CPython 3.10 macOS 11.0+ ARM64
ファイル headroom_ai-0.22.3-cp313-cp313-manylinux_2_28_aarch64.whl の詳細。
ダウンロード URL: headroom_ai-0.22.3-cp313-cp313-manylinux_

2_28_aarch64.whl
タグ: CPython 3.13、多くのLinux: glibc 2.28+ ARM64
信頼できる出版を使用してアップロードしますか?はい
アップロード方法:twine/6.1.0 CPython/3.13.12
ハッシュの使用の詳細については、こちらをご覧ください。
次の構成証明バンドルは headroom_ai-0.22.3-cp313-cp313-manylinux_2_28_aarch64.whl 用に作成されました。
出版社:
Chopratejas/headroom の release.yml
声明:
ステートメントの種類: https://in-toto.io/Statement/v1
述語タイプ: https://docs.pypi.org/attestations/publish/v1
サブジェクト名: headroom_ai-0.22.3-cp313-cp313-manylinux_2_28_aarch64.whl
件名ダイジェスト: 13314984fae54e94ecbd2101197c529d62cab06b1831f37277681ded49a227c7
シグストアの透明性エントリ: 1595359591
シグストアの統合時間:
2026 年 5 月 21 日
パーマリンク:
Chopratejas/headroom@9536d6b48c32a30d05c47e34c2e12e200f3157bd
所有者: https://github.com/chopratejas
トークン発行者: https://token.actions.githubusercontent.com
ランナー環境: github ホスト型
出版ワークフロー:
release.yml@9536d6b48c32a30d05c47e34c2e12e200f3157bd
ファイル headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl の詳細。
ダウンロードURL: headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl
タグ: CPython 3.13、macOS 11.0+ ARM64
信頼できる出版を使用してアップロードしますか?はい
アップロード方法:twine/6.1.0 CPython/3.13.12
ハッシュの使用の詳細については、こちらをご覧ください。
次の認証バンドルは headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl 用に作成されました。
出版社:
Chopratejas/headroom の release.yml
声明:
ステートメントの種類: https://in-toto.io/Statement/v1
述語タイプ: https://docs.pypi.org/attestations/publish/v1
サブジェクト名: headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl
件名ダイジェスト: d55d9a59183118539203c1f17b366c71dd1a3583e61d390cd82fcdeb89e53d3e
シグストアの透明性エントリ: 1595359324
シグストアの統合時間:
2026 年 5 月 21 日
パーマリンク:
チョラテジャス/ヘッド

部屋@9536d6b48c32a30d05c47e34c2e12e200f3157bd
所有者: https://github.com/chopratejas
トークン発行者: https://token.actions.githubusercontent.com
ランナー環境: github ホスト型
出版ワークフロー:
release.yml@9536d6b48c32a30d05c47e34c2e12e200f3157bd
ファイル headroom_ai-0.22.3-cp312-cp312-manylinux_2_28_x86_64.whl の詳細。
ダウンロード URL: headroom_ai-0.22.3-cp312-cp312-manylinux_2_28_x86_64.whl
タグ: CPython 3.12、多くのLinux: glibc 2.28+ x86-64
信頼できる出版を使用してアップロードしますか?はい
アップロード方法:twine/6.1.0 CPython/3.13.12
ハッシュの使用の詳細については、こちらをご覧ください。
次の証明書バンドルは彼のために作成されました。

[切り捨てられた]

## Original Extract

The Context Optimization Layer for LLM Applications - Cut costs by 50-90%

Skip to main content
Switch to mobile version
Warning
You are using an unsupported browser, upgrade to a newer version.
Search PyPI
search-focus#focusSearchField"
data-search-focus-target="searchField">
Search
Help
pip install headroom-ai
Copy PIP instructions
The Context Optimization Layer for LLM Applications - Cut costs by 50-90%
License Expression: Apache-2.0
SPDX
License Expression
Maintainer: Headroom Contributors
Tags
llm
,
openai
,
anthropic
,
claude
,
gpt
,
context
,
token
,
optimization
,
compression
,
caching
,
proxy
,
ai
,
machine-learning
Provides-Extra:
agno
, all
, anyllm
, bedrock
, benchmark
, code
, dev
, evals
, html
, image
, langchain
, mcp
, memory
, memory-stack
, ml
, otel
, proxy
, relevance
, reports
, strands
, voice
, voice-train
License
OSI Approved :: Apache Software License
Operating System
OS Independent
Programming Language
Python :: 3
Topic
Scientific/Engineering :: Artificial Intelligence
Software Development :: Libraries :: Python Modules
██╗ ██╗███████╗ █████╗ ██████╗ ██████╗ ██████╗ ██████╗ ███╗ ███╗
██║ ██║██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔═══██╗██╔═══██╗████╗ ████║
███████║█████╗ ███████║██║ ██║██████╔╝██║ ██║██║ ██║██╔████╔██║
██╔══██║██╔══╝ ██╔══██║██║ ██║██╔══██╗██║ ██║██║ ██║██║╚██╔╝██║
██║ ██║███████╗██║ ██║██████╔╝██║ ██║╚██████╔╝╚██████╔╝██║ ╚═╝ ██║
╚═╝ ╚═╝╚══════╝╚═╝ ╚═╝╚═════╝ ╚═╝ ╚═╝ ╚═════╝ ╚═════╝ ╚═╝ ╚═╝
The context compression layer for AI agents
60–95% fewer tokens · library · proxy · MCP · 6 algorithms · local-first · reversible
Docs ·
Install ·
Proof ·
Agents ·
Discord ·
llms.txt
AI agents / LLMs: read /llms.txt here, or fetch the live index / full docs blob .
Headroom compresses everything your AI agent reads — tool outputs, logs, RAG chunks, files, and conversation history — before it reaches the LLM. Same answers, fraction of the tokens.
Live: 10,144 → 1,260 tokens — same FATAL found.
Library — compress(messages) in Python or TypeScript, inline in any app
Proxy — headroom proxy --port 8787 , zero code changes, any language
Agent wrap — headroom wrap claude|codex|cursor|aider|copilot in one command
MCP server — headroom_compress , headroom_retrieve , headroom_stats for any MCP client
Cross-agent memory — shared store across Claude, Codex, Gemini, auto-dedup
headroom learn — mines failed sessions, writes corrections to CLAUDE.md / AGENTS.md
Reversible (CCR) — originals never deleted; LLM retrieves on demand
Your agent / app
(Claude Code, Cursor, Codex, LangChain, Agno, Strands, your own code…)
│ prompts · tool outputs · logs · RAG results · files
▼
┌────────────────────────────────────────────────────┐
│ Headroom (runs locally — your data stays here) │
│ ─────────────────────────────────────────────── │
│ CacheAligner → ContentRouter → CCR │
│ ├─ SmartCrusher (JSON) │
│ ├─ CodeCompressor (AST) │
│ └─ Kompress-base (text, HF) │
│ │
│ Cross-agent memory · headroom learn · MCP │
└────────────────────────────────────────────────────┘
│ compressed prompt + retrieval tool
▼
LLM provider (Anthropic · OpenAI · Bedrock · …)
ContentRouter — detects content type, selects the right compressor
SmartCrusher / CodeCompressor / Kompress-base — compress JSON, AST, or prose
CacheAligner — stabilizes prefixes so provider KV caches actually hit
CCR — stores originals locally; LLM calls headroom_retrieve if it needs them
→ Architecture · CCR reversible compression · Kompress-base model card
# 1 — Install
pip install "headroom-ai[all]" # Python
npm install headroom-ai # Node / TypeScript
# 2 — Pick your mode
headroom wrap claude # wrap a coding agent
headroom proxy --port 8787 # drop-in proxy, zero code changes
# or: from headroom import compress # inline library
# 3 — See the savings
headroom stats
Granular extras: [proxy] , [mcp] , [ml] , [agno] , [langchain] , [evals] . Requires Python 3.10+ .
Savings on real agent workloads:
Accuracy preserved on standard benchmarks:
Reproduce: python -m headroom.evals suite --tier 1 · Full benchmarks & methodology
60B+ tokens saved by the community — live leaderboard →
Any OpenAI-compatible client works via headroom proxy . MCP-native: headroom mcp install .
run AI coding agents daily and want savings without changing your code
work across multiple agents and want shared memory
need reversible compression — originals always retrievable via CCR
only use a single provider's native compaction and don't need cross-agent memory
work in a sandboxed environment where local processes can't run
SmartCrusher — universal JSON: arrays of dicts, nested objects, mixed types.
CodeCompressor — AST-aware for Python, JS, Go, Rust, Java, C++.
Kompress-base — our HuggingFace model, trained on agentic traces.
Image compression — 40–90% reduction via trained ML router.
CacheAligner — stabilizes prefixes so Anthropic/OpenAI KV caches actually hit.
IntelligentContext — score-based context fitting with learned importance.
CCR — reversible compression; LLM retrieves originals on demand.
Cross-agent memory — shared store, agent provenance, auto-dedup.
SharedContext — compressed context passing across multi-agent workflows.
headroom learn — plugin-based failure mining for Claude, Codex, Gemini.
Headroom exposes one stable request lifecycle across compress() , the SDK, and the proxy:
Setup → Pre-Start → Post-Start → Input Received → Input Cached → Input Routed → Input Compressed → Input Remembered → Pre-Send → Post-Send → Response Received
Transforms do the work: CacheAligner, ContentRouter, SmartCrusher, CodeCompressor, Kompress-base, IntelligentContext / RollingWindow.
Pipeline extensions observe or customize lifecycle stages via on_pipeline_event(...) .
Compression hooks sit alongside the canonical lifecycle as an additional extension seam.
Proxy extensions remain the server/app integration seam for ASGI middleware, routes, and startup policy.
Provider and tool-specific behavior lives under headroom/providers/ so core orchestration stays focused on lifecycle, sequencing, and policy.
CLI/tool slices : headroom/providers/claude , copilot , codex , openclaw
Provider runtime slices : headroom/providers/claude , gemini , plus shared backend/runtime dispatch in headroom/providers/registry.py
Core files stay orchestration-first : wrap.py , client.py , cli/proxy.py , and proxy/server.py delegate provider-specific env shaping, API target normalization, backend selection, and transport dispatch.
pip install "headroom-ai[all]" # Python, everything
npm install headroom-ai # TypeScript / Node
docker pull ghcr.io/chopratejas/headroom:latest
Granular extras: [proxy] , [mcp] , [ml] (Kompress-base), [agno] , [langchain] , [evals] . Requires Python 3.10+ .
Using pipx ? Choose a supported interpreter explicitly:
pipx install --python python3.13 "headroom-ai[all]"
→ Installation guide — Docker tags, persistent service, PowerShell, devcontainers.
headroom learn — mines failed sessions, writes corrections to CLAUDE.md / AGENTS.md / GEMINI.md .
Headroom runs locally , covers every content type, works with every major framework, and is reversible .
Attribution. Headroom ships with the excellent RTK binary for shell-output rewriting — git show --short , scoped ls , summarized installers. Huge thanks to the RTK team; their tool is a first-class part of our stack, and Headroom compresses everything downstream of it. Headroom can also use lean-ctx as the selected CLI context tool; set HEADROOM_CONTEXT_TOOL=lean-ctx before running headroom wrap ... .
git clone https://github.com/chopratejas/headroom.git && cd headroom
pip install -e ".[dev]" && pytest
Devcontainers in .devcontainer/ (default + memory-stack with Qdrant & Neo4j). See CONTRIBUTING.md .
Live leaderboard — 60B+ tokens saved and counting.
Discord — questions, feedback, war stories.
Kompress-base on HuggingFace — the model behind our text compression.
License Expression: Apache-2.0
SPDX
License Expression
Maintainer: Headroom Contributors
Tags
llm
,
openai
,
anthropic
,
claude
,
gpt
,
context
,
token
,
optimization
,
compression
,
caching
,
proxy
,
ai
,
machine-learning
Provides-Extra:
agno
, all
, anyllm
, bedrock
, benchmark
, code
, dev
, evals
, html
, image
, langchain
, mcp
, memory
, memory-stack
, ml
, otel
, proxy
, relevance
, reports
, strands
, voice
, voice-train
License
OSI Approved :: Apache Software License
Operating System
OS Independent
Programming Language
Python :: 3
Topic
Scientific/Engineering :: Artificial Intelligence
Software Development :: Libraries :: Python Modules
Release history
Release notifications |
RSS feed
Download the file for your platform. If you're not sure which to choose, learn more about installing packages .
Filter files by name, interpreter, ABI, and platform.
If you're not sure about the file name format, learn more about wheel file names .
Copy a direct link to the current filters
Copy
Uploaded
May 21, 2026
CPython 3.13 manylinux: glibc 2.28+ ARM64
Uploaded
May 21, 2026
CPython 3.13 macOS 11.0+ ARM64
Uploaded
May 21, 2026
CPython 3.12 manylinux: glibc 2.28+ x86-64
Uploaded
May 21, 2026
CPython 3.12 manylinux: glibc 2.28+ ARM64
Uploaded
May 21, 2026
CPython 3.12 macOS 11.0+ ARM64
Uploaded
May 21, 2026
CPython 3.11 manylinux: glibc 2.28+ x86-64
Uploaded
May 21, 2026
CPython 3.11 manylinux: glibc 2.28+ ARM64
Uploaded
May 21, 2026
CPython 3.11 macOS 11.0+ ARM64
Uploaded
May 21, 2026
CPython 3.10 manylinux: glibc 2.28+ x86-64
Uploaded
May 21, 2026
CPython 3.10 manylinux: glibc 2.28+ ARM64
Uploaded
May 21, 2026
CPython 3.10 macOS 11.0+ ARM64
Details for the file headroom_ai-0.22.3-cp313-cp313-manylinux_2_28_aarch64.whl .
Download URL: headroom_ai-0.22.3-cp313-cp313-manylinux_2_28_aarch64.whl
Tags: CPython 3.13, manylinux: glibc 2.28+ ARM64
Uploaded using Trusted Publishing? Yes
Uploaded via: twine/6.1.0 CPython/3.13.12
See more details on using hashes here.
The following attestation bundles were made for headroom_ai-0.22.3-cp313-cp313-manylinux_2_28_aarch64.whl :
Publisher:
release.yml on chopratejas/headroom
Statement:
Statement type: https://in-toto.io/Statement/v1
Predicate type: https://docs.pypi.org/attestations/publish/v1
Subject name: headroom_ai-0.22.3-cp313-cp313-manylinux_2_28_aarch64.whl
Subject digest: 13314984fae54e94ecbd2101197c529d62cab06b1831f37277681ded49a227c7
Sigstore transparency entry: 1595359591
Sigstore integration time:
May 21, 2026
Permalink:
chopratejas/headroom@9536d6b48c32a30d05c47e34c2e12e200f3157bd
Owner: https://github.com/chopratejas
Token Issuer: https://token.actions.githubusercontent.com
Runner Environment: github-hosted
Publication workflow:
release.yml@9536d6b48c32a30d05c47e34c2e12e200f3157bd
Details for the file headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl .
Download URL: headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl
Tags: CPython 3.13, macOS 11.0+ ARM64
Uploaded using Trusted Publishing? Yes
Uploaded via: twine/6.1.0 CPython/3.13.12
See more details on using hashes here.
The following attestation bundles were made for headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl :
Publisher:
release.yml on chopratejas/headroom
Statement:
Statement type: https://in-toto.io/Statement/v1
Predicate type: https://docs.pypi.org/attestations/publish/v1
Subject name: headroom_ai-0.22.3-cp313-cp313-macosx_11_0_arm64.whl
Subject digest: d55d9a59183118539203c1f17b366c71dd1a3583e61d390cd82fcdeb89e53d3e
Sigstore transparency entry: 1595359324
Sigstore integration time:
May 21, 2026
Permalink:
chopratejas/headroom@9536d6b48c32a30d05c47e34c2e12e200f3157bd
Owner: https://github.com/chopratejas
Token Issuer: https://token.actions.githubusercontent.com
Runner Environment: github-hosted
Publication workflow:
release.yml@9536d6b48c32a30d05c47e34c2e12e200f3157bd
Details for the file headroom_ai-0.22.3-cp312-cp312-manylinux_2_28_x86_64.whl .
Download URL: headroom_ai-0.22.3-cp312-cp312-manylinux_2_28_x86_64.whl
Tags: CPython 3.12, manylinux: glibc 2.28+ x86-64
Uploaded using Trusted Publishing? Yes
Uploaded via: twine/6.1.0 CPython/3.13.12
See more details on using hashes here.
The following attestation bundles were made for he

[truncated]
