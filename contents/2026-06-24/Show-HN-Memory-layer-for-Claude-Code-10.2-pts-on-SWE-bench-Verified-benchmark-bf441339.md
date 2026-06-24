---
source: "https://github.com/SaravananJaichandar/world-model-mcp"
hn_url: "https://news.ycombinator.com/item?id=48656153"
title: "Show HN: Memory layer for Claude Code(+10.2 pts on SWE-bench Verified benchmark)"
article_title: "GitHub - SaravananJaichandar/world-model-mcp: A production-ready MCP server that builds a world model for codebases, preventing hallucinations, repeated mistakes, and regressions in Claude Code. · GitHub"
author: "saravanan2294"
captured_at: "2026-06-24T07:07:36Z"
capture_tool: "hn-digest"
hn_id: 48656153
score: 2
comments: 0
posted_at: "2026-06-24T06:53:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Memory layer for Claude Code(+10.2 pts on SWE-bench Verified benchmark)

- HN: [48656153](https://news.ycombinator.com/item?id=48656153)
- Source: [github.com](https://github.com/SaravananJaichandar/world-model-mcp)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T06:53:54Z

## Translation

タイトル: Show HN: クロード コードのメモリ層 (SWE ベンチ検証済みベンチマークで +10.2 ポイント)
記事のタイトル: GitHub - SaravananJaichandar/world-model-mcp: コードベースのワールド モデルを構築し、クロード コードの幻覚、繰り返される間違い、および回帰を防ぐ実稼働対応の MCP サーバー。 · GitHub
説明: コードベースのワールド モデルを構築し、クロード コードの幻覚、繰り返しの間違い、および回帰を防止する実稼働対応の MCP サーバー。 - サラヴァナン・ジャイチャンダル/world-model-mcp

記事本文:
GitHub - SaravananJaichandar/world-model-mcp: コードベースのワールド モデルを構築し、クロード コードの幻覚、繰り返される間違い、およびリグレッションを防ぐ実稼働対応の MCP サーバー。 · GitHub
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
サラヴァナンジャイチャンダル
/
ワールドモデル-MC

p
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
サラヴァナンジャイチャンダル/world-model-mcp
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
45 コミット 45 コミット .github/ workflows .github/ workflows アダプター ベンチマーク ベンチマーク docs/ デプロイメント docs/ デプロイメントの例 例 フック フック スクリプト スクリプト テスト テスト world_model_server world_model_server .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.http Dockerfile.http ライセンス ライセンス QUICKSTART.md QUICKSTART.md README.md README.md RELEASE.md RELEASE.md RELEASE_NOTES.md RELEASE_NOTES.md docker-compose.yml docker-compose.yml Glama.json Glama.jsonマニフェスト.json マニフェスト.json pyproject.toml pyproject.toml サーバー.json サーバー.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントの強制、来歴、および中立的なメモリの利用。編集境界で学習した制約に対してコードの変更を検証し、圧縮後に関連するコンテキストを再挿入し、信頼度に重み付けされた解決策で矛盾を追跡し、クロード コード、カーソル、およびパイ全体で実行する時間的ナレッジ グラフ。
ステータス: v0.9.1 — 26 個の MCP ツール、19 個の CLI サブコマンド、375 個のテスト、SWE ベンチ 49 インスタンスにわたる +10.2 ポイントのペア デルタ (ドメイン内で +15.0 ポイント、クロスドメインで +6.9 ポイント)、105 ペアの矛盾解決ベンチマークで検証済みの反復ミス ベンチマーク。 v0.9 は経験的なウェッジ証明を同梱しています。これは、永続知識層がパブリック タスク コーパス上で繰り返されるコーディング エージェントの間違いを測定可能に減らすかどうかをテストする、ロックされた事前登録された方法論です。結果は、ドメイン内およびドメイン間の正の効果を確認し、外部での回帰は観察されませんでした。

f-ドメインタスク。完全なタスクごとのテーブル、2 つのクロスドメイン フリップのメカニズム分析 (sphinx-9461 が最もクリーンなケース)、および Benchmarks/repeat-mistake/RESULTS.md の正直な制限。 v0.8.1 では、矛盾解決ベンチマークが 19 カテゴリにわたる 105 ペアに拡張されました。 v0.8.0 では、証拠タイプごとの TTL、項目ごとの来歴フィールド source_tool およびconfirmer、スラッシュ コマンド書き込み操作、およびsolve_conflict のコンファーマー パラメータを使用したドメイン対応の信頼度減衰が追加されました。 Antigravity アダプターは、SDK の TransformCompactionHook を保留して 4 回連続のリリースを保留しました。次回は 2026 年 7 月 24 日に再確認します。 v0.7.6 では、/world-model スラッシュ コマンドと status-watch TUI ウィジェットが追加されました。 v0.7.5 には Codex CLI アダプターが追加されました。 v0.7.0 では、PostCompact 自動インジェクション、遅延強制層、信頼度に重み付けされた矛盾解決、および圧縮監査ログが導入されました。貢献は大歓迎です。
mcp-name: io.github.SaravananJaichandar/world-model-mcp
world-model-mcp が役に立った場合は、リポジトリにスターを付けるか、何が機能したか、何が機能しなかったかについて問題を開いてください。私はすべてを読み、フィードバックが次に出荷されるものを形作ります。
World Model MCP は、すべてのコーディング セッションから次のことを学習するコードベースの時間ナレッジ グラフを作成します。
幻覚の防止 -- 使用前に既知のエンティティに対して API/関数参照を検証します
繰り返される間違いを阻止 -- 修正から制約を学習し、将来のセッションに適用します
リグレッションを減らす -- バグ修正を追跡し、変更が重要な領域に達したときに警告します
Survive Compation -- エージェントのコンテキスト ウィンドウがリセットされた後、最上位の制約と最近のファクトを再挿入します。
矛盾を解決する -- 信頼性、最新性、またはソース数を使用して、矛盾する事実の間で勝者を選択します。
これは、Claude Code、Cursor、またはその他の MCP と並行して実行される長期記憶層と考えてください。

認識されたコーディングエージェント。
SWE ベンチ検証済みの反復ミス ベンチマーク — 中心的なウェッジの証明。 django、sympy、matplotlib、scikit-learn、sphinx にわたる 50 個の SWE ベンチ検証済みタスクを、ベースラインと治療のペア比較として実行します。方法論は 2026 年 6 月 17 日に (データが存在する前に) Benchmarks/repeat-mistake/DESIGN.md でロックされていたため、結果をゴールポストの移動で非難することはできません。
ヘッドラインの結果 — サブセット 1 (ドメイン内: django + sympy) ベースライン 15/20 = 75.0 パーセント、治療 18/20 = 90.0 パーセント、デルタ +15.0 ポイント、FAIL から PASS への反転が 4 件、回帰が 1 件ありました。サブセット 2 (クロスドメイン: matplotlib + scikit-learn + sphinx) ベースライン 18/29 = 62.1 パーセント、処理 20/29 = 69.0 パーセント、デルタ +6.9 ポイント (2 回のフリップおよび回帰ゼロ)。 49 インスタンスにわたるペアの結果の合計: 33/49 ～ 38/49、デルタ +10.2 ポイント。
クロスドメイン転送はきれいに分離されました — サブセット 2 処理アームは 4 つのサブセット 1 制約 (django および sympy ディレクティブ) のみをロードし、11 個のサブセット 2 制約を保持して、1 つのリポジトリ ファミリからの学習が別のリポジトリ ファミリに一般化するかどうかをテストしました。ロードされた制約に基づいた妥当なメカニズムの説明を伴う 2 つのクロスドメイン フリップ。 Sphinx-9461 は最も強力なケースです。sympy クラスメソッド制約が Sphinx クラスメソッド ラッパーのラップ解除バグに転送されました。
RESULTS.md に埋め込まれた正直な警告 — 単一トライアル設計、サブセット 1 での制約障害の重複、小さなクロスドメイン転送速度、上流の SWE ベンチの PIP フラグ問題による 1 つのドロップされたインスタンス、およびジャッジモデルの自己参照リスクを含む 7 つの明示的な制限。付録に隠されているのではなく、そのまま記載されています。
完全な再現性の成果物 - すべての進行状況 JSONL、予測 JSON、結果 JSONL、分類 JSONL、制約 JSON、およびハーネス レポート JSON コミットメント

ベンチマーク/repeat-mistake/ の d 。 Failure_classifier.py と learning_hook.py のジャッジ プロンプトをロックしました。両部門にわたるエージェントの合計コストは、Claude Code サブスクリプションで約 90 米ドルでした。
矛盾解決ベンチマークの拡張 -- v0.7.4 の 24 ペアのベンチマークは、19 のカテゴリにわたって手動で厳選された 105 ペアに拡大しました。 6 つの新しいカテゴリ (source_tool_corroboration 、confirmer_overrides_pending 、decay_advantage_session_vs_source 、decay_advantage_stale_session 、evidence_type_user_correction 、sett_beats_higher_confidence ) は、v0.8.0 スキーマを具体的に実行します。決定論的ランナー (ベンチマーク/矛盾-200/run.py)完全な戦略ごと + カテゴリごとの内訳は、 Benchmarks/conflicts-200/RESULTS.md にあります。
数値に関する正直な枠組み: 新しいデータセットは、v0.7.4 の 24 ペア セットよりも困難です。これは、新しいカテゴリが生の信頼度ランキングではなく、スキーマ認識 (確認者、証拠タイプ、減衰) を意図的にテストするためです。見出しの数値: keep_most_sources 99.0%、keep_higher_confidence 81.0%、auto 77.1%、keep_higher_confidence_decayed 90.5% (evidence_type が存在する 21 ペアで)、全戦略全体で 78.2%。元の 24 ペア v0.7.4 93.5% の数値は、ベンチマーク/矛盾/で変更されずに保存され、無効になりません。別の（より小さく、より簡単な）コーパスをテストしました。
ウェッジのベンチマークは v0.9 です。「学習ループは、パブリック タスク コーパスで繰り返されるコーディング エージェントの間違いを測定可能に減らしますか?」このリリースでの矛盾解決作業は、内部スキーマの正確性の検証です。公開されたエッセイの枠組みにマッピングされる経験的な成果物 (学習ループは耐久性のある層です) は、SWE ベンチ スタイルの反復ミス ベンチマークを備えた v0.9 に組み込まれています。
ドメイン認識の信頼度減衰 -- 指数を備えた新しい world_model_server/decay.py モジュール

半減期減衰は、evidence_type ごとに異なります。半減期: ソースコード 365 日、テスト 180 日、セッション 14 日、ユーザー修正 730 日、バグ修正 365 日。減衰は読み取り時に適用されるため (バックグラウンド タスクなし)、次の query_fact 呼び出しは時間補正された信頼度を返します。解決されたファクト (正規ステータス、または確認者 != NULL を持つファクト) は自動遷移しません。 0.2 信頼度を下回って減衰する合成ファクトと、0.1 信頼度を下回って減衰する裏付けられたファクトは、読み取り時に自動的に置き換えられ、次の圧縮注入までに腐敗が表面化します。
ファクトに関する項目ごとの来歴フィールド -- 3 つの追加列 (source_tool TEXT 、confirmer TEXT 、last_decay_at TIMESTAMP )、すべて NULL デフォルトで、バックフィルなし。 source_tool は、どのツールがファクトを書き込んだかを記録します (例: claude_code 、 codex 、cursor 、 pi 、 user )。確認者は、主張者とは別に、誰がそれを確認したかを記録します。 NULL は保留中を意味し、非 NULL は解決済みを意味します。どちらも Fact モデルで公開され、 create_fact を通じて伝播されます。 Patdolitse ( anthropics/claude-code#47023 ) と ferhimedmine ( openai/codex#19195 ) への公約を尊重します。
スラッシュ コマンドの書き込み操作 -- 2 つの新しいサブコマンド。 /world-modelsolve <id> は、矛盾を解決済みとしてマークします (手動。信頼度に重み付けされたピッキングには、resolve_conflict MCP ツールを使用します)。 /world-modelforget <id> は、ファクトに valid_at を設定します (監査ログに保存されます。現在のみの読み取りではそれ以降スキップされます)。どちらも冪等であり、未知の ID について明確に報告します。ヘルプ テキストには、v0.7.6 に同梱されている読み取り専用サブコマンドとともに両方がリストされるようになりました。
解決_矛盾は確認者を受け入れます -- 確認者引数が MCP ツールまたはその基礎となる解決関数に提供されると、勝利したファクトにはその値がスタンプされた確認者列が取得されます。これは、「アサーターは X と言う」と「X は c である」を区別する仕様プリミティブです。

作業グループのスケッチに従って、Y によって確認されました。
反重力アダプターは 3 回連続でリリースされました。 2026 年 6 月 13 日の再検証では、SDK で InspectHook として宣言されている OnCompactionHook が、TransformCompactionHook も追加のコンテキスト戻りフィールドも持たないことが判明しました。負荷を伴うメモリ注入コントラクトはまだ SDK に存在しません。次に 2026 年 6 月 27 日に再確認します。
エージェント内 /world-model スラッシュ コマンド -- エージェント ハーネス内でユーザーが入力すると、チャットを離れることなくワールド モデルの状態が表示されます。 v0.7.6 では読み取り専用 ( status 、矛盾、最近、ヘルプ)。書き込み操作 (resolve、forget) は v0.8 に導入されます。既存の inject_helper で UserPromptSubmit をインターセプトすることにより、Claude Code、Cursor、Codex、および pi 全体で動作します。 Codex が強制する厳密なキャメルケース形式 (deny_unknown_fields) で追加のコンテキストを返すため、ハーネスごとの分岐を持たずに、同じワイヤアップが 4 つのハーネスすべてにサービスを提供します。
world-model status-watch TUI ウィジェット -- エージェントと並行して実行され、5 秒ごとに更新されるターミナル ペイン。制約 (合計、重大度 = エラー、重大度 = 警告)、未解決の矛盾、事実 (標準 / 合成 / 置き換え)、および最後の圧縮時間を表示します。すでに依存関係にある豊富なライブラリに基づいて構築されています

[切り捨てられた]

## Original Extract

A production-ready MCP server that builds a world model for codebases, preventing hallucinations, repeated mistakes, and regressions in Claude Code. - SaravananJaichandar/world-model-mcp

GitHub - SaravananJaichandar/world-model-mcp: A production-ready MCP server that builds a world model for codebases, preventing hallucinations, repeated mistakes, and regressions in Claude Code. · GitHub
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
SaravananJaichandar
/
world-model-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
SaravananJaichandar/world-model-mcp
main Branches Tags Go to file Code Open more actions menu Folders and files
45 Commits 45 Commits .github/ workflows .github/ workflows adapters adapters benchmarks benchmarks docs/ deployment docs/ deployment examples examples hooks hooks scripts scripts tests tests world_model_server world_model_server .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.http Dockerfile.http LICENSE LICENSE QUICKSTART.md QUICKSTART.md README.md README.md RELEASE.md RELEASE.md RELEASE_NOTES.md RELEASE_NOTES.md docker-compose.yml docker-compose.yml glama.json glama.json manifest.json manifest.json pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Enforcement, provenance, and harness-neutral memory for AI coding agents. A temporal knowledge graph that validates code changes against learned constraints at the edit boundary, re-injects relevant context after compaction, tracks contradictions with confidence-weighted resolution, and runs across Claude Code, Cursor, and pi.
Status: v0.9.1 — 26 MCP tools, 19 CLI subcommands, 375 tests, SWE-bench Verified repeat-mistake benchmark with +10.2 pts paired delta across 49 instances (+15.0 pts within-domain, +6.9 pts cross-domain), 105-pair contradiction-resolution benchmark. v0.9 ships the empirical wedge proof: a locked, pre-registered methodology tested whether the persistent-knowledge layer measurably reduces repeated coding-agent mistakes on a public task corpus. Result confirms positive within-domain and cross-domain effects with zero observed regressions on out-of-domain tasks. Full per-task tables, mechanistic analysis of the two cross-domain flips (sphinx-9461 is the cleanest case), and honest limitations in benchmarks/repeat-mistake/RESULTS.md . v0.8.1 expanded the contradiction-resolution benchmark to 105 pairs across 19 categories. v0.8.0 added domain-aware confidence decay with per-evidence-type TTL, per-item provenance fields source_tool and confirmer , slash command write operations, and a confirmer parameter on resolve_contradiction . Antigravity adapter held for the fourth consecutive release pending a TransformCompactionHook in the SDK; next re-verify 2026-07-24. v0.7.6 added the /world-model slash command and status-watch TUI widget. v0.7.5 added the Codex CLI adapter. v0.7.0 introduced PostCompact auto-injection, the defer enforcement tier, confidence-weighted contradiction resolution, and a compaction audit log. Contributions welcome.
mcp-name: io.github.SaravananJaichandar/world-model-mcp
If world-model-mcp helped you, star the repo or open an issue with what worked or didn't. I read every one and the feedback shapes what ships next.
World Model MCP creates a temporal knowledge graph of your codebase that learns from every coding session to:
Prevent Hallucinations -- Validates API/function references against known entities before use
Stop Repeated Mistakes -- Learns constraints from corrections, applies them in future sessions
Reduce Regressions -- Tracks bug fixes and warns when changes touch critical regions
Survive Compaction -- Re-injects top constraints and recent facts after the agent's context window resets
Resolve Contradictions -- Picks a winner between conflicting facts using confidence, recency, or source count
Think of it as a long-term memory layer that runs alongside Claude Code, Cursor, or any MCP-aware coding agent.
Repeat-mistake benchmark on SWE-bench Verified — the central wedge proof. 50 SWE-bench Verified tasks across django, sympy, matplotlib, scikit-learn, and sphinx, run as a paired baseline-vs-treatment comparison. Methodology was locked at benchmarks/repeat-mistake/DESIGN.md on 2026-06-17 (before the data existed) so the result cannot be accused of goalpost-moving.
Headline results — Subset 1 (within-domain: django + sympy) baseline 15/20 = 75.0 percent, treatment 18/20 = 90.0 percent, delta +15.0 pts with 4 FAIL to PASS flips and 1 regression. Subset 2 (cross-domain: matplotlib + scikit-learn + sphinx) baseline 18/29 = 62.1 percent, treatment 20/29 = 69.0 percent, delta +6.9 pts with 2 flips and zero regressions. Combined paired result across 49 instances: 33/49 to 38/49, delta +10.2 pts.
Cross-domain transfer isolated cleanly — the Subset 2 treatment arm loaded ONLY the 4 Subset 1 constraints (django and sympy directives), holding out the 11 Subset 2 constraints to test whether learning from one repo family generalizes to a different one. Two cross-domain flips with plausible mechanistic explanations grounded in the loaded constraints. Sphinx-9461 is the strongest case: a sympy classmethod constraint transferred to a sphinx classmethod-wrapper unwrapping bug.
Honest caveats embedded in RESULTS.md — seven explicit limitations including single-trial design, constraint-failure overlap on Subset 1, the small cross-domain transfer rate, one dropped instance due to an upstream SWE-bench pip flag issue, and judge-model self-reference risk. Stated verbatim rather than hidden in an appendix.
Full reproducibility artifacts — every progress JSONL, predictions JSON, results JSONL, classification JSONL, constraints JSON, and harness report JSON committed in benchmarks/repeat-mistake/ . Locked judge prompts in failure_classifier.py and learning_hook.py . Total agent cost across both arms was approximately 90 USD on a Claude Code subscription.
Contradiction-resolution benchmark expansion -- the v0.7.4 24-pair benchmark grew to 105 hand-curated pairs across 19 categories. Six new categories exercise the v0.8.0 schema specifically: source_tool_corroboration , confirmer_overrides_pending , decay_advantage_session_vs_source , decay_advantage_stale_session , evidence_type_user_correction , settled_beats_higher_confidence . Deterministic runner at benchmarks/contradictions-200/run.py ; full per-strategy + per-category breakdown at benchmarks/contradictions-200/RESULTS.md .
Honest framing on the numbers : the new dataset is harder than v0.7.4's 24-pair set because the new categories deliberately test schema awareness (confirmer, evidence_type, decay) rather than raw confidence ranking. Headline numbers: keep_most_sources 99.0%, keep_higher_confidence 81.0%, auto 77.1%, keep_higher_confidence_decayed 90.5% (on the 21 pairs where evidence_type is present), overall 78.2% across all strategies. The original 24-pair v0.7.4 93.5% number is preserved unchanged at benchmarks/contradictions/ and is not invalidated; it tested a different (smaller, easier) corpus.
The wedge benchmark is v0.9 : "does the learning loop measurably reduce repeated coding-agent mistakes on a public task corpus?" The contradiction-resolution work in this release is internal schema-correctness validation. The empirical artifact that maps to the published essay framing — the learning loop is the durable layer — lands in v0.9 with a SWE-bench-style repeat-mistake benchmark.
Domain-aware confidence decay -- new world_model_server/decay.py module with exponential half-life decay per evidence_type . Half-lives: source_code 365d, test 180d, session 14d, user_correction 730d, bug_fix 365d. Decay applies on read (no background task), so the next query_fact call returns the time-corrected confidence. Settled facts ( canonical status, or any fact with confirmer != NULL ) never auto-transition. Synthesized facts that decay below 0.2 confidence and corroborated facts that decay below 0.1 confidence auto-supersede on read, surfacing rot to the next compaction injection.
Per-item provenance fields on facts -- three additive columns ( source_tool TEXT , confirmer TEXT , last_decay_at TIMESTAMP ), all NULL-defaulted, no backfill. source_tool records which tool wrote the fact (e.g. claude_code , codex , cursor , pi , user ). confirmer records who confirmed it, distinct from the asserter; NULL means pending, non-NULL means settled. Both are exposed on the Fact model and propagated through create_fact . Honors the public commitment to Patdolitse ( anthropics/claude-code#47023 ) and ferhimedamine ( openai/codex#19195 ).
Slash command write operations -- two new subcommands. /world-model resolve <id> marks a contradiction as resolved (manual; for confidence-weighted picking use the resolve_contradiction MCP tool). /world-model forget <id> sets invalid_at on a fact (preserved in the audit log; current-only reads skip it from then on). Both are idempotent and report cleanly on unknown ids. Help text now lists both alongside the read-only subcommands shipped in v0.7.6.
resolve_contradiction accepts confirmer -- when a confirmer argument is provided to the MCP tool or its underlying resolve function, the winning fact gets its confirmer column stamped with that value. This is the spec primitive that distinguishes "the asserter says X" from "X is confirmed by Y" per the working group sketch.
Antigravity adapter held for the third consecutive release. The 2026-06-13 re-verification found OnCompactionHook declared as InspectHook in the SDK with no TransformCompactionHook and no additional_context return field. The load-bearing memory-injection contract still does not exist in the SDK. Next re-verify 2026-06-27.
In-agent /world-model slash command -- typed by the user inside the agent harness, surfaces the world model state without leaving the chat. Read-only in v0.7.6 ( status , contradictions , recent , help ); write operations ( resolve , forget ) land in v0.8. Works across Claude Code, Cursor, Codex, and pi by intercepting UserPromptSubmit in the existing inject_helper . Returns additionalContext in the strict camelCase shape Codex enforces ( deny_unknown_fields ), so the same wire-up serves all four harnesses without a per-harness branch.
world-model status-watch TUI widget -- terminal pane that runs alongside the agent and refreshes every 5 seconds. Shows constraints (total, severity=error, severity=warning), unresolved contradictions, facts (canonical / synthesized / superseded), and last compaction time. Built on the rich library already in the depen

[truncated]
