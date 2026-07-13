---
source: "https://github.com/insightitsGit/ChorusGraph/"
hn_url: "https://news.ycombinator.com/item?id=48888185"
title: "Same agent tasks, 76% fewer LLM calls – we moved semantic cache inside the graph"
article_title: "GitHub - insightitsGit/ChorusGraph: Deterministic AI agent orchestration — comment via Discussions & Issues; code changes by maintainers only. · GitHub"
author: "insightits"
captured_at: "2026-07-13T05:27:04Z"
capture_tool: "hn-digest"
hn_id: 48888185
score: 2
comments: 0
posted_at: "2026-07-13T05:14:13Z"
tags:
  - hacker-news
  - translated
---

# Same agent tasks, 76% fewer LLM calls – we moved semantic cache inside the graph

- HN: [48888185](https://news.ycombinator.com/item?id=48888185)
- Source: [github.com](https://github.com/insightitsGit/ChorusGraph/)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T05:14:13Z

## Translation

タイトル: 同じエージェント タスクで、LLM コールが 76% 減少 – セマンティック キャッシュをグラフ内に移動しました
記事のタイトル: GitHub - InsightitsGit/ChorusGraph: 決定論的 AI エージェント オーケストレーション — ディスカッションと問題を介したコメント。コードの変更はメンテナのみによって行われます。 · GitHub
説明: 決定論的な AI エージェント オーケストレーション — ディスカッションと問題を介したコメント。コードの変更はメンテナのみによって行われます。 - 洞察力Git/ChorusGraph

記事本文:
GitHub - InsightitsGit/ChorusGraph: 決定論的 AI エージェント オーケストレーション — ディスカッションと問題を介したコメント。コードの変更はメンテナのみによって行われます。 · GitHub
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
洞察力Git
/
コーラスグラフ
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
100 コミット 100 コミット .cursor/ rules .cursor/ rules .github .github ベンチマーク ベンチマーク コーラスグラフ コーラスグラフ デプロイ デプロイ ドキュメント docs sbom sbom スクリプト スクリプト テスト テスト Web サイト Web サイト .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml 要件-lock.txt 要件-lock.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
セマンティック キャッシュ、スワップ可能な取得 (PrismRAG)、監査可能なメモリ、およびエンタープライズ強化を備えたネイティブ エージェント ランタイム — 1 つの pip インストール、5 つのプラグイン ポート。
pip インストールコーラスグラフ
コーラスグラフデモ
インタラクティブなデモ (プロダクト ハント / 起動): InsightitsGit.github.io/ChorusGraph/demo.html — クリックスルー ウォークスルー、ステップ 1 ～ 3 には API キーはありません。
ChorusGraph = ネイティブ エンジン + Prism スタック · LangGraph = A/B 比較のみのオプションのベースライン ( docs/TERMINOLOGY.md )
ChorusGraph は LangGraph ラッパーではありません。デフォルトでアタッチされた Prism 製品スタック (セマンティック キャッシュ、L2 取得、L3 メモリ、ルート台帳、チェックポイント、および可観測性) を備えたネイティブ BSP グラフ エンジン (chorusgraph.core.Graph) が同梱されています。
ノード、エッジ、条件付きルーティングはネイティブ エンジンで定義します。キャッシュ、取得、メモリ、およびツールは、ChorusStack の明示的なポートを介してプラグインされます。オーケストレーションを書き換えることなく、Redis、ベクター RAG、またはカスタム ツール レジストリを交換します。
ChorusGraph 自体のコードには、製品パスに対する LangGraph の依存関係はありません。スケジューラとすべてのプラグイン ポートは LangGraph をインポートしません。 (コア依存関係 prismlang は内部的に LangGraph を使用します)

r 独自のチェックポイント設定ユーティリティ — これは pip show に表示されますが、ChorusGraph エンジンがそれを呼び出すことはありません。) FL*/HL* 比較シナリオを実行する場合にのみ、chorusgraph[benchmark] をインストールします。
本番環境の LLM エージェントを構築するということは、通常、オーケストレーション、セマンティック キャッシュ、ベクター DB、リランカー、チェックポイント、監査ログの 6 つのシステムを結合することを意味します。 ChorusGraph は、それらを明示的なプラグイン ポートを備えた 1 つのランタイムとして出荷します。
pip インストールコーラスグラフ
from コーラスグラフ import Graph 、 START 、 END 、 ChorusStack
コーラスグラフから。コア。ノードインポート dict_node_adapter
スタック = ChorusStack 。デフォルト ( tenant_id = "デモ" )
g = グラフ ( tenant_id = "demo" 、graph_id = "hello" )
g 。 add_node (
「エコー」、
dict_node_adapter ( lambda s : { "reply" : f"Hello, { s . get ( 'name' , 'world' ) } " }, hop = "echo" ),
）
g 。 add_edge ( START , "エコー" )
g 。 add_edge ( "エコー" , END )
アウト = g 。コンパイル (スタック = スタック)。 invoke ({ "名前" : "ChorusGraph" })
print ( out ) # {'reply': 'こんにちは、ChorusGraph'}
Chorusgraph-demo # ルーティング + 台帳 (LLM フリー)
Chorusgraph-finance-patterns # ReAct / Plan-Solve / Reflection (GEMINI_API_KEY が必要)
Chorusgraph-audit --log your_queries.jsonl # シミュレートされたキャッシュ ヒット率 (API キーなし)
開発者ガイド: docs/DEVELOPER_GUIDE.md — 計画と推論、ドメインのパフォーマンス (金融と医療)、コード例。
完全なインストール ガイド: docs/INSTALL.md · AI IDE プロンプト: docs/AI_IDE_PROMPTS.md
特徴
説明
ネイティブグラフエンジン
BSP スケジューラ、エンベロープ チャネル、条件付きルーティング - プロダクト パスに LangGraph なし
セマンティック キャッシュ (L1)
2 段階ゲート: 粗リコール → 完全検証。ドメインごとの安全なリプレイポリシー
取得(L2)
キーワードのデフォルト。ベクター + 分類のための PrismRAGRetrievalBackend (オプトイン追加)
メモリ(L3)
PrismCortex 構造化された再生可能なメモリ
ルート台帳
ホップごとの監査証跡: キャッシュ ヒット

s、スコア、期間、rule_chain
チェックポイント
SQLite のデフォルト。 Postgres エンタープライズ永続性 (ライセンスゲート型)
ツールレジストリ
サンドボックスを備えたホワイトリストに登録されたツール。 MCP互換パターン
回復力
タイムアウト、再試行、サーキット ブレーカー、正常なノード障害
可観測性
構造化された JSON ログ、OpenTelemetry トレース、ヘルス/メトリクス エンドポイント
マルチテナントガード
テナントの分離、リソース制限、漏洩テスト
コールド監査 CLI
Chorusgraph-audit — クエリログからキャッシュ節約量を見積もる (LLM 呼び出しなし)
エージェントのパターン
Chorusgraph.agents.Agent による ReAct、Plan-Solve、Reflection — グラフ = plan
ベンチマークマトリックス
公平性開示付きの 8 つのシナリオ (FL/FC/HL/HC)
パッケージングの展開
Dockerfile、docker-compose、k8s マニフェスト
LangGraph および DIY スタックとの比較
ランググラフ単体
DIY スタック (オーケストレーター + Redis + ベクター DB + リランカー + ログ)
コーラスグラフ
オーケストレーション
✅ ステートグラフ
あなたは統合します
✅ ネイティブグラフ
セマンティックキャッシュ
❌ 自分でロールしてください
別途サービス+接着剤
✅ 内蔵 L1、交換可能
検索/RAG
❌ 外部
クロマ/松ぼっくり + カスタム コード
✅ RetrievalBackend ポート
監査/説明可能性
限定
カスタムロギング
✅ ホップごとのルート台帳
安全なキャッシュリプレイ
あなたの問題
あなたの問題
✅ ドメインプロファイル (例: 医療における事実のみ)
ベンチマークの証拠
該当なし
該当なし
✅ 公開された A/B と LangGraph の比較
LangGraph の依存関係
必須
オプション
製品パスには何もありません
ChorusGraph には、同じモデル、ツール、プロンプト、ワークロードを公平に比較するための LangGraph ベースライン ( benchmark/fl* 、 benchmark/hl* ) が含まれています。ネイティブ シナリオ ( benchmark/fc* 、 benchmark/hc* ) は、chorusgraph.core.Graph のみでコンパイルされます。
┌───────────────────────────┐
│あなたのg

raph — ノード、エッジ、条件付きルーティング │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ ChorusStack — 交換可能なポート │
│ ┌─────┬───┬───┬─────────┐ │
│ │ キャッシュ │ メモリ │ ツール │ 検索 (L2) │ │
│ │ Prism / │ Cortex │ レジストリ │ キーワード / PrismRAG │ │
│ │ Redis │ │ │ │
│ ━─────┴───┴─────┴─────────┘ │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ エンジン（固定）：BSPスケジューラー・エンベロープ・レゾナンス・JL │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ ルート台帳・チェックポイント・テナントガード・可観測性 │
━━━━━━━━━━━━━━━━━━━━━━━━┘
詳細: docs/COMPOSE.md · docs/DEVELOPER_GUIDE.md
ChorusStack 上の 4 つの交換可能なポート — エンジンとスケジューラは固定されたままです。
コーラスグラフから。インポートを作成する ChorusStack 、 PrismRAGRetrievalBackend 、 RedisCacheBackend
コーラスグラフから。エンベッダーインポート PrismlangOnnxEmbedder
バックエンド =

PrismRAGRetrievalBackend (
embedder = PrismlangOnnxEmbedder (),
マッピング = { "カテゴリ" : [...], "ルール" : [...]},
）
バックエンド。インデックス ( your_corpus )
スタック = (
コーラススタック 。デフォルト ( tenant_id = "acme" )
。 with_retrieval (バックエンド)
。 with_cache ( RedisCacheBackend ( tenant_id = "acme" 、 redis_url = "redis://localhost:6379/0" ))
）
完全なプラグイン ガイド: docs/PLUGINS.md
レイヤー
コンポーネント
役割
L0 — ホップ
プリズムラング
64-d 状態圧縮 + ルールチェーン監査
L1 — キャッシュ
プリズムキャッシュ
セマンティックゲート、共鳴スコアリコール
L2 — 知識
検索プラグイン
キーワードのデフォルト · ベクトル + 分類法のオプトイン
再ランク付け
プリズムレゾナンス
共有基板のリランク
L3 — メモリ
プリズムコーテックス
構造化された再生可能な記憶
輸送
コーラス / プリズムAPI
クロスノードエンベロープ・フェデレーションフック
ChorusGraph は、Prism ファミリの統合ランタイムです。PrismLang、PrismCache、PrismCortex、PrismRAG は、個別の科学プロジェクトではなく、デフォルトまたはオプトインの追加機能として出荷されます。
コンパニオン: PrismGuard (プロンプトインジェクションファイアウォール)
PrismGuard ( 0.1.4 ) は別のパッケージであり、ChorusStack ポートではありません。取得 / LLM ホップの前に監査可能なプロンプト挿入チェックが必要な場合は、ChorusGraph と一緒にインストールします。
pip インストールコーラスグラフ「プリズムガード[プリズム,ガードモデル]==0.1.4」
プリズムガード モデルのダウンロード # ~705 MB ONNX — ホイールには含まれていません。 GitHub リリース v0.1.2 より
プリズムガードから。統合。コーラスグラフのインポート (
create_checker_from_env 、
make_guard_handler 、
ルート_アフター_ガード 、
）
checker = create_checker_from_env () # プロセスごとに 1 回
ガード = make_guard_handler (チェッカー)
# START → ガード → [ブロック → END |取得→ …]
# Graph.add_node("guard", dict_node_adapter(guard, hop="guard")) と接続する
# キャッシュゲートホップの前にガードを配置して、ブロックされたプロンプトがキャッシュをシードしないようにします
リンク
URL
PyPI
https://pypi.org/project/prism

ガード/・0.1.4
GitHub
https://github.com/insightitsGit/PrismGuard
統合ガイド
https://github.com/insightitsGit/PrismGuard/blob/main/docs/integration-guide.md
ONNXモデルリリース
https://github.com/insightitsGit/PrismGuard/releases/tag/v0.1.2
docs/PLUGINS.md も参照してください。
公正な A/B と有能な LangGraph ベースライン — 同じモデル、ツール、プロンプト、ワークロード。
ここから開始します: docs/BENCHMARK_RESULTS.md · アーカイブ インデックス: benchmark/results/mvp_scenarios/README.md · マシン ポインター: benchmark/results/mvp_scenarios/latest.json
方法論: ベンチマーク/FAIRNESS_H9.md · 統合テーブル: ベンチマーク/結果/BENCHMARK_LATENCY_LLM_SUMMARY.md
2026 年 7 月の方法論修正 (ベンチマークのみ - ライブラリ リリースなし): FL2 研究者プロンプトは、annual_rate_pct を使用します (ツール スキーマと一致します)。比較スクリプトは、LangGraph の成功の分母としてエージェント/ツールのエラーをカウントします。 FL2 対 FC2 を膨らませたプレフィックス ランを置き換えます。無効な割り当て量の実行が重い_20260708_124337 を引用しないでください。
タスクの成功 (LangGraph → ChorusGraph) — 中層、n=100
シナリオ
ランググラフ
コーラスグラフ
デルタ
財務シングル (FL1→FC1)
87.0%
98.0%
+11.0ポイント
金融マルチ（FL2→FC2）
87.0%
94.0%
+7.0ポイント
ヘルスケアシングル（HL1→HC1）
74.0%
79.0%
+5.0ポイント
ヘルスケアマルチ（HL2→HC2）
59.0%
85.0%
+26pp
タスクの成功 - 重い層、n=300
シナリオ
ランググラフ
コーラスグラフ
デルタ
財務シングル (FL1→FC1)
90.0%
96.7%
+6.7ポイント
金融マルチ（FL2→FC2）
89.0%
93.0%
+4.0ポイント
ヘルスケアシングル（HL1→HC1）
73.7%
84.0%
+10.3ポイント
ヘルスケアマルチ（HL2→HC2）
62.3%
77.3%
+15pp
LLM コールとレイテンシー (中間層、タスクごとの平均)
シナリオ
LLM コール (L → C)
平均レイテンシ ミリ秒 (L → C)
キャッシュヒット(C)
FL1 / FC1
3.24 → 0.77 (−76%)
4760 → 1348 (−72%)
52%
FL2 / FC2
2.03 → 0.69 (−66%)
3269 → 1085 (−67%)
40%
HL1/HC1
3.00 → 1.56 (−48%)
7036 → 3990 (−43%)
6

0%
HL2/HC2
3.82 → 3.15 (−18%)
10296 → 10753 (引き分け)
51%
LLM コールとレイテンシー (重い層、タスクごとの平均)
シナリオ
LLM コール (L → C)
平均レイテンシ ミリ秒 (L → C)
キャッシュヒット(C)
FL1 / FC1
3.33 → 0.80 (−76%)
4972 → 1318 (−73%)
49.7%
FL2 / FC2
2.04 → 0.75 (−63%)
3081 → 1335 (−57%)
34.7%
HL1/HC1
2.94 → 1.33 (−55%)
7105 → 3812 (−46%)
72.7%
HL2/HC2
3.85 → 2.67 (−31%)
10354 → 9537 (−8%; p95 同率)
79.0%
ヘルスケア マルチでは、設計により LLM コールの節約が少なくなります (ファクトのみのキャッシュ、判断ホップが少なくなります)

[切り捨てられた]

## Original Extract

Deterministic AI agent orchestration — comment via Discussions & Issues; code changes by maintainers only. - insightitsGit/ChorusGraph

GitHub - insightitsGit/ChorusGraph: Deterministic AI agent orchestration — comment via Discussions & Issues; code changes by maintainers only. · GitHub
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
insightitsGit
/
ChorusGraph
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
100 Commits 100 Commits .cursor/ rules .cursor/ rules .github .github benchmark benchmark chorusgraph chorusgraph deploy deploy docs docs sbom sbom scripts scripts tests tests website website .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml requirements-lock.txt requirements-lock.txt View all files Repository files navigation
Native agent runtime with semantic cache, swappable retrieval (PrismRAG), auditable memory, and enterprise hardening — one pip install, five plug-in ports.
pip install chorusgraph
chorusgraph-demo
Interactive demo (Product Hunt / launch): insightitsGit.github.io/ChorusGraph/demo.html — click-through walkthrough, no API key for steps 1–3.
ChorusGraph = native engine + Prism stack · LangGraph = optional baseline for A/B comparison only ( docs/TERMINOLOGY.md )
ChorusGraph is not a LangGraph wrapper. It ships a native BSP graph engine ( chorusgraph.core.Graph ) with the Prism product stack attached by default: semantic cache, L2 retrieval, L3 memory, Route Ledger, checkpoints, and observability.
You define nodes, edges, and conditional routing on the native engine. Cache, retrieval, memory, and tools plug in through explicit ports on ChorusStack — swap Redis, vector RAG, or custom tool registries without rewriting orchestration.
ChorusGraph's own code has no LangGraph dependency on the product path. The scheduler and all plug-in ports never import LangGraph. (Core dependency prismlang uses LangGraph internally for its own checkpointing utilities — it appears in pip show , but the ChorusGraph engine never calls it.) Install chorusgraph[benchmark] only when running FL*/HL* comparison scenarios.
Building production LLM agents usually means gluing six systems: orchestration, semantic cache, vector DB, reranker, checkpointing, and audit logs. ChorusGraph ships them as one runtime with explicit plug-in ports.
pip install chorusgraph
from chorusgraph import Graph , START , END , ChorusStack
from chorusgraph . core . node import dict_node_adapter
stack = ChorusStack . defaults ( tenant_id = "demo" )
g = Graph ( tenant_id = "demo" , graph_id = "hello" )
g . add_node (
"echo" ,
dict_node_adapter ( lambda s : { "reply" : f"Hello, { s . get ( 'name' , 'world' ) } " }, hop = "echo" ),
)
g . add_edge ( START , "echo" )
g . add_edge ( "echo" , END )
out = g . compile ( stack = stack ). invoke ({ "name" : "ChorusGraph" })
print ( out ) # {'reply': 'Hello, ChorusGraph'}
chorusgraph-demo # routing + ledger (LLM-free)
chorusgraph-finance-patterns # ReAct / Plan-Solve / Reflection (needs GEMINI_API_KEY)
chorusgraph-audit --log your_queries.jsonl # simulated cache hit rate (no API key)
Developer guide: docs/DEVELOPER_GUIDE.md — planning & reasoning, domain performance (finance vs healthcare), code examples.
Full install guide: docs/INSTALL.md · AI IDE prompts: docs/AI_IDE_PROMPTS.md
Feature
Description
Native graph engine
BSP scheduler, envelope channels, conditional routing — no LangGraph on product paths
Semantic cache (L1)
Two-stage gate: coarse recall → full verify; safe replay policies per domain
Retrieval (L2)
Keyword default; PrismRAGRetrievalBackend for vector + taxonomy (opt-in extra)
Memory (L3)
PrismCortex structured, replayable memory
Route Ledger
Per-hop audit trail: cache hits, scores, durations, rule_chain
Checkpoints
SQLite default; Postgres enterprise persistence (license-gated)
Tool registry
Allowlisted tools with sandbox; MCP-compatible patterns
Resilience
Timeouts, retries, circuit breakers, graceful node failure
Observability
Structured JSON logs, OpenTelemetry traces, health/metrics endpoints
Multi-tenant guards
Tenant isolation, resource limits, leakage tests
Cold audit CLI
chorusgraph-audit — estimate cache savings from query logs (no LLM calls)
Agent patterns
ReAct, Plan-Solve, Reflection via chorusgraph.agents.Agent — graph = plan
Benchmark matrix
8 scenarios (FL/FC/HL/HC) with fairness disclosure
Deploy packaging
Dockerfile, docker-compose, k8s manifests
Comparison with LangGraph and DIY stacks
LangGraph alone
DIY stack (orchestrator + Redis + vector DB + reranker + logs)
ChorusGraph
Orchestration
✅ StateGraph
You integrate
✅ Native Graph
Semantic cache
❌ Roll your own
Separate service + glue
✅ Built-in L1, swappable
Retrieval / RAG
❌ External
Chroma/Pinecone + custom code
✅ RetrievalBackend port
Audit / explainability
Limited
Custom logging
✅ Route Ledger per hop
Safe cache replay
Your problem
Your problem
✅ Domain profiles (e.g. facts-only in healthcare)
Benchmark proof
N/A
N/A
✅ Published A/B vs LangGraph
LangGraph dependency
Required
Optional
None on product path
ChorusGraph includes LangGraph baselines ( benchmark/fl* , benchmark/hl* ) for fair apples-to-apples comparison — same model, tools, prompts, workload. Native scenarios ( benchmark/fc* , benchmark/hc* ) compile with chorusgraph.core.Graph only.
┌─────────────────────────────────────────────────────────────┐
│ Your graph — nodes, edges, conditional routing │
├─────────────────────────────────────────────────────────────┤
│ ChorusStack — swappable ports │
│ ┌──────────┬──────────┬──────────┬──────────────────────┐ │
│ │ Cache │ Memory │ Tools │ Retrieval (L2) │ │
│ │ Prism / │ Cortex │ Registry │ Keyword / PrismRAG │ │
│ │ Redis │ │ │ │ │
│ └──────────┴──────────┴──────────┴──────────────────────┘ │
├─────────────────────────────────────────────────────────────┤
│ Engine (fixed): BSP scheduler · envelopes · Resonance · JL │
├─────────────────────────────────────────────────────────────┤
│ Route Ledger · checkpoints · tenant guards · observability │
└─────────────────────────────────────────────────────────────┘
Details: docs/COMPOSE.md · docs/DEVELOPER_GUIDE.md
Four swappable ports on ChorusStack — engine and scheduler stay fixed.
from chorusgraph . compose import ChorusStack , PrismRAGRetrievalBackend , RedisCacheBackend
from chorusgraph . embedders import PrismlangOnnxEmbedder
backend = PrismRAGRetrievalBackend (
embedder = PrismlangOnnxEmbedder (),
mapping = { "categories" : [...], "rules" : [...]},
)
backend . index ( your_corpus )
stack = (
ChorusStack . defaults ( tenant_id = "acme" )
. with_retrieval ( backend )
. with_cache ( RedisCacheBackend ( tenant_id = "acme" , redis_url = "redis://localhost:6379/0" ))
)
Full plug-in guide: docs/PLUGINS.md
Layer
Component
Role
L0 — hop
PrismLang
64-d state compression + rule_chain audit
L1 — cache
PrismCache
Semantic gate, Resonance-scored recall
L2 — knowledge
Retrieval plug-in
Keyword default · vector + taxonomy opt-in
rerank
PrismResonance
Shared substrate rerank
L3 — memory
PrismCortex
Structured, replayable memory
transport
CHORUS / PrismAPI
Cross-node envelopes · federation hooks
ChorusGraph is the integration runtime for the Prism family — PrismLang, PrismCache, PrismCortex, PrismRAG ship as defaults or opt-in extras, not separate science projects.
Companion: PrismGuard (prompt-injection firewall)
PrismGuard ( 0.1.4 ) is a separate package — not a ChorusStack port. Install it alongside ChorusGraph when you want an auditable prompt-injection check before retrieve / LLM hops:
pip install chorusgraph " prismguard[prism,guard-model]==0.1.4 "
prismguard-model download # ~705 MB ONNX — not in the wheel; from GitHub Release v0.1.2
from prismguard . integrations . chorusgraph import (
create_checker_from_env ,
make_guard_handler ,
route_after_guard ,
)
checker = create_checker_from_env () # once per process
guard = make_guard_handler ( checker )
# START → guard → [blocked → END | retrieve → …]
# Wire with Graph.add_node("guard", dict_node_adapter(guard, hop="guard"))
# Place guard before cache-gated hops so blocked prompts never seed cache
Link
URL
PyPI
https://pypi.org/project/prismguard/ · 0.1.4
GitHub
https://github.com/insightitsGit/PrismGuard
Integration guide
https://github.com/insightitsGit/PrismGuard/blob/main/docs/integration-guide.md
ONNX model release
https://github.com/insightitsGit/PrismGuard/releases/tag/v0.1.2
See also docs/PLUGINS.md .
Fair A/B vs competent LangGraph baselines — same model, tools, prompts, workload.
Start here: docs/BENCHMARK_RESULTS.md · archive index: benchmark/results/mvp_scenarios/README.md · machine pointer: benchmark/results/mvp_scenarios/latest.json
Methodology: benchmark/FAIRNESS_H9.md · consolidated tables: benchmark/results/BENCHMARK_LATENCY_LLM_SUMMARY.md
July 2026 methodology fixes (benchmark-only — no library release ): FL2 researcher prompt uses annual_rate_pct (matches tool schema); comparison script counts agent/tool errors in LangGraph success denominators. Supersedes pre-fix runs that inflated FL2 vs FC2. Do not cite invalid quota run heavy_20260708_124337 .
Task success (LangGraph → ChorusGraph) — mid tier, n=100
Scenario
LangGraph
ChorusGraph
Delta
Finance single (FL1→FC1)
87.0%
98.0%
+11.0 pp
Finance multi (FL2→FC2)
87.0%
94.0%
+7.0 pp
Healthcare single (HL1→HC1)
74.0%
79.0%
+5.0 pp
Healthcare multi (HL2→HC2)
59.0%
85.0%
+26 pp
Task success — heavy tier, n=300
Scenario
LangGraph
ChorusGraph
Delta
Finance single (FL1→FC1)
90.0%
96.7%
+6.7 pp
Finance multi (FL2→FC2)
89.0%
93.0%
+4.0 pp
Healthcare single (HL1→HC1)
73.7%
84.0%
+10.3 pp
Healthcare multi (HL2→HC2)
62.3%
77.3%
+15 pp
LLM calls and latency (mid tier, mean per task)
Scenario
LLM calls (L → C)
Mean latency ms (L → C)
Cache hit (C)
FL1 / FC1
3.24 → 0.77 (−76%)
4760 → 1348 (−72%)
52%
FL2 / FC2
2.03 → 0.69 (−66%)
3269 → 1085 (−67%)
40%
HL1 / HC1
3.00 → 1.56 (−48%)
7036 → 3990 (−43%)
60%
HL2 / HC2
3.82 → 3.15 (−18%)
10296 → 10753 (tie)
51%
LLM calls and latency (heavy tier, mean per task)
Scenario
LLM calls (L → C)
Mean latency ms (L → C)
Cache hit (C)
FL1 / FC1
3.33 → 0.80 (−76%)
4972 → 1318 (−73%)
49.7%
FL2 / FC2
2.04 → 0.75 (−63%)
3081 → 1335 (−57%)
34.7%
HL1 / HC1
2.94 → 1.33 (−55%)
7105 → 3812 (−46%)
72.7%
HL2 / HC2
3.85 → 2.67 (−31%)
10354 → 9537 (−8%; p95 tie)
79.0%
Healthcare multi saves fewer LLM calls by design (facts-only cache, judgment hops re

[truncated]
