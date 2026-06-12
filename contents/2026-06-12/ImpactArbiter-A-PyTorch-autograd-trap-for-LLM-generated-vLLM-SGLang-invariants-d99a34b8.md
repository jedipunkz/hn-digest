---
source: "https://github.com/msunda17/impactarbiter-cli"
hn_url: "https://news.ycombinator.com/item?id=48509033"
title: "ImpactArbiter – A PyTorch autograd trap for LLM-generated vLLM/SGLang invariants"
article_title: "GitHub - msunda17/impactarbiter-cli: A deterministic PyTorch autograd verification trap for catching silent KV-cache routing and block-alignment failures in vLLM and SGLang serving infrastructure. · GitHub"
author: "maniksundar"
captured_at: "2026-06-12T21:37:47Z"
capture_tool: "hn-digest"
hn_id: 48509033
score: 1
comments: 0
posted_at: "2026-06-12T20:22:17Z"
tags:
  - hacker-news
  - translated
---

# ImpactArbiter – A PyTorch autograd trap for LLM-generated vLLM/SGLang invariants

- HN: [48509033](https://news.ycombinator.com/item?id=48509033)
- Source: [github.com](https://github.com/msunda17/impactarbiter-cli)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T20:22:17Z

## Translation

タイトル: ImpactArbiter – LLM で生成された vLLM/SGLang 不変式用の PyTorch autograd トラップ
記事のタイトル: GitHub - msunda17/impactarbiter-cli: vLLM および SGLang サービス提供インフラストラクチャでのサイレント KV キャッシュ ルーティングとブロック アライメントの失敗を捕捉するための決定論的な PyTorch autograd 検証トラップ。 · GitHub
説明: vLLM および SGLang サービス提供インフラストラクチャでのサイレント KV キャッシュ ルーティングとブロック アライメントの失敗を捕捉するための決定論的な PyTorch autograd 検証トラップ。 - msunda17/impactarbiter-cli

記事本文:
GitHub - msunda17/impactarbiter-cli: vLLM および SGLang サービス インフラストラクチャでのサイレント KV キャッシュ ルーティングとブロック アライメントの失敗を捕捉するための決定論的な PyTorch autograd 検証トラップ。 · GitHub
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
ムスンダ17
/
インパクトアービター-cli
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE src src テスト テスト .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md README.md README.md pyproject.toml pyproject.toml results.csv results.csv すべてのファイルを表示 リポジトリ ファイルナビゲーション
📖 立ち上げ記事全文を読み、ここで 2D テンソル衝突デモをご覧ください: https://maniksundar.substack.com/p/the-physics-illusion-why-llms-still
KV キャッシュ ルーティング カーネル用の LLM 生成の単体テストでは、サイレント障害モードが発生します。LLM は、実装とテストの両方で同じバグを幻覚的に示し、カーネルが不正確なままテストが合格してしまいます。これは、LLM がコードとテストの両方を作成するときに、同じ欠陥のあるメンタル モデルに基づいて推論するために発生します。 ImpactArbiter は、2 段階の RAG パイプラインを使用してこれに対処します。まず、Distill Agent が実際の研究論文からルーティング ロジックを抽出して要約します。次に、コーディング エージェントがその概要に基づいて実装とテストを作成します。生成されたコードは、勾配署名を SymPy オラクルと比較する PyTorch autograd トラップを通じて実行されます。このトラップは、LLM 自身の test_route() アサーションが成功した場合でも、単体テストが見逃したバグを捕捉します。
Python -m venv venv
source venv/bin/activate # Windows の場合: venv\Scripts\activate
2. 開発モードでインストールする
pip install -e 。
3. API キーの設定 (プロバイダーの選択)
import OPENAI_API_KEY= " your-openai-API-key "
# Windows の場合 (PowerShell): $env:OPENAI_API_KEY="your-openai-api-key"
クロード (人間):
import ANTHROPIC_API_KEY= " あなたの anthropic-API キー "
# Windows の場合 (PowerShell): $env:ANTHROPIC_API_KEY="your-anthro

pic-API-キー」
ジェミニ / バーテックス AI:
# gcloud が認証され、プロジェクトが設定されていることを確認します
gcloud 認証ログイン
gcloud config set プロジェクト Impactagent
import GOOGLE_CLOUD_PROJECT= " インパクトエージェント "
# Windows の場合 (PowerShell): $env:GOOGLE_CLOUD_PROJECT="impactagent"
永続的な構成の場合は、.env ファイルに以下を追加します。
OPENAI_API_KEY=あなたのopenai APIキー
ANTHROPIC_API_KEY=あなたの anthropic-API キー
GOOGLE_CLOUD_PROJECT=影響力
デモコマンド
Impactarbiter auto-heal --oracle radix --model gemini
追加のフラグ
--full-agent-trace : コード生成と修復試行の前に LLM 思考連鎖推論を表示します。
--live : キャッシュされた確定的リプレイの代わりにライブ LLM API 呼び出しを使用します (API キーが必要)
--mock : 確定的リプレイを使用してオフライン評価を実行します (API キーがない場合のデフォルト)
ライブ LLM 生成と完全なトレースの例:
Impactarbiter verify --workflow Agentic-kv-scheduler --full-agent-trace --live
サンプル出力 (2D 非対称リング バッファ)
────────────────── インパクト アービター — オートヒール ──────────────────
モデル: vertex_ai/gemini-2.5-pro
[論文をダウンロードしました]
https://arxiv.org/pdf/2312.07104.pdf
[クイック蒸留]
### KV キャッシュ ルーティング仕様: プランナーとエグゼキューターのハンドオフ
...
[生成されたコードとテスト]
def Route_radix_2d(b_local_idx, head_idx, prefix_length_h, total_blocks_h, block_size):
k = prefix_length_h + b_local_idx
論理ブロック = k // ブロックサイズ
オフセット = k % block_size
return (head_idx、logical_block、offset)
[LLM ユニットテストパス ✅]
LLM 自己検証に合格しました。
[オートグラードトラップ失敗 ❌ ハードブロック]
発散=1.00e+00 > トル=1e-04
グラデーションダイバージェンスマップ — KV_cache.grad (ヘッド × ブロック × オフセット)
トークン (b=5,h=0,prefix_h=60,N_h=4) |予期: ヘッド = 0 ブロック = 0 オフセット = 1 |取得: ヘッド = 0 ブロック = 4 オフセット = 1
非

ゼロ勾配: [0, 4, 1, :] - 128 個のフロートの配線が間違っています
[オートヒール試行 1/3]
def Route_radix_2d(b_local_idx, head_idx, prefix_length_h, total_blocks_h, block_size):
絶対_idx = プレフィックス長_h + b_local_idx
logical_block = (absolute_idx // block_size) % total_blocks_h
オフセット = 絶対IDx % ブロックサイズ
return (head_idx、logical_block、offset)
[ファイナルパス✅]
divergence=0.00e+00 (1 回の回復試行後)
LLM の非決定性とトラップの信頼性について
autograd トラップ自体は完全に決定的であるため、同一のコードは常に同一の勾配発散結果を生成します。
異なるのは、LLM が特定の実行時に正しいルーティング ロジックを生成するか、誤ったルーティング ロジックを生成するかです。
これは、実際の運用環境を反映しています。エージェントが生成したサービス コードは通過することもありますが、多くの場合、不規則な境界で静かに失敗します。
ImpactArbiter を使用すると、エージェントが生成するコードを決定論的に検証できるため、モデルが「今回は正しく実行された」ことを期待する必要がなくなります。
実際には、Gemini 2.5 Pro は、重要な 2D リング バッファ ラップ ケースの試行の約 65% で誤ったルーティングを生成します。このトラップは、誤った実装をすべて検出し、偽陰性はゼロです。
2D RadixAttention テスト マトリックス (デフォルト)
実稼働関連のデモに推奨されます。
境界フィクスチャ [15、99、100、105、128] は履歴比較モード用に維持されます。
最新の評価実行の要約統計:
オラクル
合計実行数
トラップ発射
無事に治りました
基数 2d
32
21
21
基数(1D)
15
9
9
vllm (ページング)
15
0
0
リポジトリのレイアウト
ソース/
§── oracles/ # SymPy AST + ラムディファイド呼び出し可能
§── トラップ/ # autograd トラップと ASCII 分岐マップ
§── ファザー/ # 明示的な境界フィクスチャ
§── cli/ # 自動修復パイプライン + litellm エージェント + ペーパーエクストラクター
━── db/ # nextpaper.db (SQLite) validation_tra

セス
テスト/
§── test_paged_oracle.py
§── test_radix_oracle.py
━── test_trap.py
テストの実行
pytest テスト/ -v
testing/test_trap.py 内の 4 つの耐荷重クレームはすべて合格する必要があります。
ImpactArbiter はオープン プロジェクトです。新しいオラクル、ファズ ケース、バグ レポートの貢献を歓迎します。完全な貢献者ガイドについては、CONTRIBUTING.md を参照してください。短い要約:
問題の報告: .github/ISSUE_TEMPLATE/ にある関連テンプレートを使用して、GitHub の問題を開きます。バグ レポートには、正確な CLI コマンド、使用したモデル、分岐マップ (またはスタック トレース)、および必要に応じて nextpaper.db 行の内容を含める必要があります。
新しいオラクルの提供: すべてのオラクルは、(1) SymPy AST と src/oracles/ のラムディファイド呼び出し可能、(2) src/trap/ の決定論的 autograd トラップ、および (3) src/fuzzer/ の明示的な境界フィクスチャのトリプルとして出荷する必要があります。 3 つすべてが欠けている PR は返送されます。
方法論の議論: 方法論の問題を開きます。これらは毎週レビューされ、模擬幻覚率とオラクルごとのセッション ウィンドウを調整するために使用されます。
テンプレート
いつ使用するか
バグレポート.md
トラップ、自動修復、エバリュエーター、または CSV エクスポートにより、不正な動作またはクラッシュ動作が発生します。
oracle_contribution.md
新しいアテンション/ルーティング オラクル (FlashInfer、MLA、スライディング ウィンドウなど) を提案したいと考えています。
方法論.md
あなたは幻覚率の調整、セッションウィンドウの定義、またはトラップ許容量に同意しません。
バグの報告 — 再現可能な最小限のレポート
バグは、失敗を決定的に再現できる場合にのみ対処可能です。以下を含めてください:
コマンド ライン — すべてのフラグを含む、正確な Impactarbiter ... の呼び出し。
環境 — Python のバージョン、OS、および --live または --mock が使用されたかどうか。
モデル ID ( --live の場合) — 例: vertex_ai/gemini-2.5-pro 。 API キーを貼り付けないでください。
失敗す

rface — 勾配発散マップ、自動修復スタック トレース、または results.csv 概要ブロックのいずれかを貼り付けます。
予想と実際 — オラクルが予測したものとエージェント/トラップが返したもの。
バグがトラップ自体にある場合 (偽 PASS または偽 FAIL)、問題のエージェント関数と対応する境界ケースを src/fuzzer/ からアタッチします。トラップの正確性に関するバグは P0 として扱われます。
vLLM および SGLang サービス インフラストラクチャでのサイレント KV キャッシュ ルーティングとブロック アライメントの失敗を捕捉するための決定論的な PyTorch autograd 検証トラップ。
github.com/msunda17/impactarbiter-cli
トピックス
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

A deterministic PyTorch autograd verification trap for catching silent KV-cache routing and block-alignment failures in vLLM and SGLang serving infrastructure. - msunda17/impactarbiter-cli

GitHub - msunda17/impactarbiter-cli: A deterministic PyTorch autograd verification trap for catching silent KV-cache routing and block-alignment failures in vLLM and SGLang serving infrastructure. · GitHub
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
msunda17
/
impactarbiter-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md README.md README.md pyproject.toml pyproject.toml results.csv results.csv View all files Repository files navigation
📖 Read the full launch post and watch the 2D tensor collision demo here: https://maniksundar.substack.com/p/the-physics-illusion-why-llms-still
LLM-generated unit tests for KV-cache routing kernels suffer from a silent failure mode: the LLM hallucinates the same bug in both the implementation and the test, causing the test to pass while the kernel remains incorrect. This happens because LLMs reason from the same flawed mental model when writing both code and tests. ImpactArbiter addresses this by using a two-stage RAG pipeline: first, a Distill Agent extracts and summarizes the routing logic from the actual research paper; second, a Coding Agent writes the implementation and test based on that summary. The generated code is then run through a PyTorch autograd trap that compares gradient signatures against SymPy oracles. The trap catches bugs that unit tests miss, even when the LLM's own test_route() assertions pass.
python -m venv venv
source venv/bin/activate # On Windows: venv\Scripts\activate
2. Install in Development Mode
pip install -e .
3. Configure API Keys (Choose Your Provider)
export OPENAI_API_KEY= " your-openai-api-key "
# On Windows (PowerShell): $env:OPENAI_API_KEY="your-openai-api-key"
Claude (Anthropic):
export ANTHROPIC_API_KEY= " your-anthropic-api-key "
# On Windows (PowerShell): $env:ANTHROPIC_API_KEY="your-anthropic-api-key"
Gemini / Vertex AI:
# Ensure gcloud is authenticated and project is set
gcloud auth login
gcloud config set project impactagent
export GOOGLE_CLOUD_PROJECT= " impactagent "
# On Windows (PowerShell): $env:GOOGLE_CLOUD_PROJECT="impactagent"
For persistent configuration, add these to your .env file:
OPENAI_API_KEY=your-openai-api-key
ANTHROPIC_API_KEY=your-anthropic-api-key
GOOGLE_CLOUD_PROJECT=impactagent
Demo Command
impactarbiter auto-heal --oracle radix --model gemini
Additional Flags
--full-agent-trace : Display LLM Chain-of-Thought reasoning before code generation and heal attempts
--live : Use live LLM API calls instead of cached deterministic replay (requires API key)
--mock : Run offline evaluation with deterministic replay (default if no API key)
Example with live LLM generation and full trace:
impactarbiter verify --workflow agentic-kv-scheduler --full-agent-trace --live
Sample Output (2D Asymmetric Ring Buffer)
─────────────────── IMPACT ARBITER — AUTO-HEAL ───────────────────
Model: vertex_ai/gemini-2.5-pro
[PAPER DOWNLOADED]
https://arxiv.org/pdf/2312.07104.pdf
[QUICK DISTILL]
### KV Cache Routing Specification: Planner-Executor Handoff
...
[GENERATED CODE & TESTS]
def route_radix_2d(b_local_idx, head_idx, prefix_length_h, total_blocks_h, block_size):
k = prefix_length_h + b_local_idx
logical_block = k // block_size
offset = k % block_size
return (head_idx, logical_block, offset)
[LLM UNIT TEST PASS ✅]
LLM self-validation passed.
[AUTOGRAD TRAP FAIL ❌ HARD_BLOCK]
divergence=1.00e+00 > tol=1e-04
GRADIENT DIVERGENCE MAP — KV_cache.grad (head × block × offset)
Token (b=5,h=0,prefix_h=60,N_h=4) | Expected: head=0 block=0 offset=1 | Got: head=0 block=4 offset=1
Non-zero gradient at: [0, 4, 1, :] — misrouted 128 floats
[AUTO-HEAL attempt 1/3]
def route_radix_2d(b_local_idx, head_idx, prefix_length_h, total_blocks_h, block_size):
absolute_idx = prefix_length_h + b_local_idx
logical_block = (absolute_idx // block_size) % total_blocks_h
offset = absolute_idx % block_size
return (head_idx, logical_block, offset)
[FINAL PASS ✅]
divergence=0.00e+00 (after 1 heal attempts)
On LLM non-determinism and trap reliability
The autograd trap itself is fully deterministic, which means identical code always produces identical gradient divergence results.
What varies is whether the LLM generates correct or incorrect routing logic on a given run.
This mirrors real production reality: agent-generated serving code sometimes passes, but often silently fails on ragged boundaries.
ImpactArbiter gives you deterministic verification of whichever code the agent produces, so you're not relying on hoping the model "got it right this time."
In practice, Gemini 2.5 Pro generates incorrect routing on roughly 65% of attempts for the critical 2D ring-buffer wrap cases. The trap catches every incorrect implementation with zero false negatives.
2D RadixAttention Test Matrix (Default)
Recommended for production-relevant demo.
Boundary fixtures [15, 99, 100, 105, 128] are maintained for historical comparison mode.
Summary statistics from most recent evaluation run:
Oracle
Total Runs
Trap Fired
Healed Successfully
radix-2d
32
21
21
radix (1D)
15
9
9
vllm (Paged)
15
0
0
Repository layout
src/
├── oracles/ # SymPy ASTs + lambdified callables
├── trap/ # autograd trap & ASCII divergence map
├── fuzzer/ # explicit boundary fixtures
├── cli/ # auto-heal pipeline + litellm agent + paper extractor
└── db/ # nextpaper.db (SQLite) validation_traces
tests/
├── test_paged_oracle.py
├── test_radix_oracle.py
└── test_trap.py
Running the tests
pytest tests/ -v
The four load-bearing claims in tests/test_trap.py must all pass.
ImpactArbiter is an open project — contributions of new oracles, fuzz cases, and bug reports are welcome. See CONTRIBUTING.md for the full contributor guide. A short summary:
Reporting issues : Open a GitHub issue using the relevant template under .github/ISSUE_TEMPLATE/ . Bug reports must include the exact CLI command, the model used, the divergence map (or stack trace), and the contents of nextpaper.db row(s) when relevant.
Contributing a new oracle : Every oracle must ship as a triple — (1) a SymPy AST plus a lambdified callable in src/oracles/ , (2) a deterministic autograd trap in src/trap/ , and (3) explicit boundary fixtures in src/fuzzer/ . PRs without all three will be sent back.
Discussing methodology : Open a methodology issue — these are reviewed weekly and used to calibrate the mock hallucination rates and per-oracle session windows.
Template
When to use
bug_report.md
The trap, auto-heal, evaluator, or CSV export produces incorrect or crashing behavior.
oracle_contribution.md
You want to propose a new attention/routing oracle (e.g. FlashInfer, MLA, sliding-window).
methodology.md
You disagree with a hallucination-rate calibration, session-window definition, or trap tolerance.
Reporting a bug — minimum reproducible report
A bug is only actionable when we can replay the failure deterministically. Please include:
Command line — The exact impactarbiter ... invocation, including all flags.
Environment — Python version, OS, and whether --live or --mock was used.
Model identity (if --live ) — e.g. vertex_ai/gemini-2.5-pro . Do not paste API keys.
Failure surface — Paste either the gradient divergence map, the auto-heal stack trace, or the results.csv summary block.
Expected vs actual — What the oracle predicts vs what the agent / trap returned.
If the bug is in the trap itself (false PASS or false FAIL), attach the offending agent function and the corresponding boundary case from src/fuzzer/ . Trap correctness bugs are treated as P0.
A deterministic PyTorch autograd verification trap for catching silent KV-cache routing and block-alignment failures in vLLM and SGLang serving infrastructure.
github.com/msunda17/impactarbiter-cli
Topics
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
