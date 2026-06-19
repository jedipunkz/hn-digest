---
source: "https://github.com/Byron2306/EdgeK-BEAST"
hn_url: "https://news.ycombinator.com/item?id=48597071"
title: "Beast – governed output gateway for AI coding agents"
article_title: "GitHub - Byron2306/EdgeK-BEAST: Governed output gateway for agentic coding tools — enforces output contracts, repairs non-compliant patches, and learns which tool calls are worth making. · GitHub"
author: "Byron230686"
captured_at: "2026-06-19T13:19:50Z"
capture_tool: "hn-digest"
hn_id: 48597071
score: 2
comments: 0
posted_at: "2026-06-19T10:43:41Z"
tags:
  - hacker-news
  - translated
---

# Beast – governed output gateway for AI coding agents

- HN: [48597071](https://news.ycombinator.com/item?id=48597071)
- Source: [github.com](https://github.com/Byron2306/EdgeK-BEAST)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T10:43:41Z

## Translation

タイトル: Beast – AI コーディング エージェント向けの管理された出力ゲートウェイ
記事のタイトル: GitHub - Byron2306/EdgeK-BEAST: エージェント コーディング ツールの管理された出力ゲートウェイ — 出力コントラクトを強制し、非準拠のパッチを修復し、どのツール呼び出しを行う価値があるかを学習します。 · GitHub
説明: エージェント コーディング ツール用の管理された出力ゲートウェイ — 出力コントラクトを強制し、非準拠のパッチを修復し、どのツール呼び出しを行う価値があるかを学習します。 - Byron2306/EdgeK-BEAST

記事本文:
GitHub - Byron2306/EdgeK-BEAST: エージェント コーディング ツール用の管理された出力ゲートウェイ — 出力コントラクトを強制し、非準拠のパッチを修復し、どのツール呼び出しを行う価値があるかを学習します。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
バイロン2306
/
EdgeK-BEAST
公共
通知

ns
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット アプリ アプリ ベンチマーク ベンチマーク bin bin デプロイ/ 生成 デプロイ/ 生成されたドキュメント ドキュメント ポリシー スクリプト スクリプト テスト テスト vscode-extension vscode-extension .gitignore .gitignore BEAST UI.png BEAST UI.png BEAST logo.png BEAST logo.png BEAST マスコット透過性.png BEAST マスコット透過性.png BEAST マスコット.png BEASTマスコット.png EdgeK BEAST VS CODE IDE.md EdgeK BEAST VS CODE IDE.md EdgeK BEAST VS CODE IDE.txt EdgeK BEAST VS CODE IDE.txt EdgeK_BEAST_Meta_Optimization_Whitepaper.md EdgeK_BEAST_Meta_Optimization_Whitepaper.md Gemini_Generated_Image_z6vjayz6vjayz6vj.png Gemini_Generated_Image_z6vjayz6vjayz6vj.png README.md README.md README_BEAST_UI_CONTROLS_PATCH.md README_BEAST_UI_CONTROLS_PATCH.md README_BEAST_UI_PATCH.md README_BEAST_UI_PATCH.md README_BEAST_UI_SPP_PATCH.md README_BEAST_UI_SPP_PATCH.md conftest.py conftest.py pytest.ini pytest.ini 要件-統合.txt 要件-統合.txt 要件-litellm.txt 要件-litellm.txt 要件-semantic.txt要件-セマンティック.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
BEAST - 効率的なエージェント システムとツールのブローカー
エージェントコーディングツールの管理された出力ゲートウェイ。
BEAST は、AI コーディング エージェント (Cursor、Claude Code、VS Code Copilot) と LLM プロバイダーの間に位置します。これは、出力コントラクトを強制し、ファイルシステムに触れる前に非準拠のパッチを修復し、どのツールを呼び出す価値があるかを学習するなど、何が入って何が出てくるかを管理します。
AIコーディングエージェントは注意を払っていません。 3 行が必要な場合は、ファイル全体を読み取ります。彼らは自分のパスに書き込みます

そうはなりません。彼らは冗長な検索にトークンの予算を費やします。プロバイダーが不正な形式の JSON を返すと、何も通知せずに失敗するか、コードが破損します。
入力ガバナンス — コンテキストの圧縮、ツールの遅延学習、予算の執行、サーキット ブレーカー
出力ガバナンス — すべてのモデル応答は、ディスクにアクセスする前に、型指定された出力コントラクト ( Beast.action_intent.v1 ) に対して解析されます。非準拠のパッチはローカルで修復され、検証されます。検証が失敗した場合は、何も書き込まれません。
決定的 — 10 タスク、5 レーン
レーン
完了しました
中央値トークン
対生
生（BEASTなし）
0 / 10
47,661
—
コンテキストのみ
0 / 10
44
−99.9%
ラグ
8/10
296
−99.4%
ラグ + ツール
10 / 10
326
−99.3%
フルビースト
10 / 10
390
−99.2%
生のコンテキストは、モデルがスコープ付きの問題を推論する前にトークン バジェットに達します。 BEAST は 400 トークン未満でタスクを 100% 完了し、pytest スイートに合格することで検証されています。
ライブプロバイダー — 20 のプロバイダールートにわたる 192 のタスク
結果
カウント
BEAST のエンドツーエンドの完了
192 / 192
クリーンプロバイダーの完了
36 / 192
BEAST によって救出された完成品
156 / 192
プロバイダーの生の出力の 79% が、準拠していないか、不正な形式であるか、不完全でした。 BEASTは彼ら全員を救出しました。出力ガバナンスがなければ、これらの 156 のタスクは通知なく失敗するか、破損したパッチを書き込むことになるでしょう。
ランク
プロバイダー
役割
クリーン
フィットネス
レイテンシ
1
オフクラウド
候補パッチプロバイダー
5/10
0.663
14秒
2
puter_deepseek
候補パッチ (高遅延)
4/10
0.619
13秒
3
一致する
候補パッチプロバイダー
4/10
0.614
6.7秒
4
ディープインフラ
候補パッチ (高遅延)
4/10
0.612
32秒
5
抱きしめる顔
救助支援型アクション IR
3/10
0.583
1.6秒
6
nスケール
救助支援型アクション IR
3/10
0.581
7.8秒
7
ミストラル
救助支援（コードストラル）
2/10
0.545
4.1秒
8
オープンルーター
迅速な救助支援アクション IR
2/10
0.544
3.8秒
9
サンバノバ
迅速な救助支援アクション IR
1/10
0.512

3.0秒
10
クラウドフレア
エッジ/マイクロタスク
1/10
0.483
2.1秒
11–14
大脳、フェザーレス、nvidia_nim、ジェミニ
スカウト/セレクター
0～2/10
0.33～0.42
変化する
15–16
グロク、llm7
スカウトのみ
0/10
0.23
速い
17–18
aion_labs 、ノビタ
レート制限/レスキュー
1/10
0.39～0.51
変化する
19–20
双曲線、fal
使用しないでください（認証/請求）
0/10
—
—
注目すべき発見:
コンピューター経由の DeepSeek は、無料のプロキシ ルート上で 4 つのクリーン パスを達成しました (有料プロバイダーに匹敵します)。 BEAST は、ガバナンスを通じて型破りな無料ルートを運用可能にすることができます。
LLM7 は 100% のタスクで有効な JSON を返しましたが、出力スキーマを通過したのは 10% のみでした。出力ガバナーがなければ、機能しているように見えます。そうではありません。
NVIDIA NIM はすべてのタスクで出力コントラクトに失敗しました。 BEAST は両方の対象タスクを修復し、救助しました。サイレント障害ゼロ。
DeepInfra が観測したコスト: 検証済みの管理されたコード修正ごとに ~0.000332 ドル。
コーディングエージェント（カーソル/クロードコード/VSコード）
│
▼
┌─────────────────────┐
│ ビーストゲートウェイ │
│ │
│ 入力側 出力側 │
│ ──────────── │
│ コンテキストエコノミー アウトプットコントラクト │
│ ツールの怠惰 ローカル検証者 │
│ 予算台帳パッチコンパイラ │
│ サーキットブレーカー アンカーレゾルバ │
│ ワークスペースグラフ修復エンジン │
│ MCP ブローカー サンドボックスバリデーター │
│ │
│ メモリ: L0 ポリシー → L4 フォレンジック アーカイブ│
━━━━━━━━━━━━━━━━━━━━━━┘
│
▼
任意の LLM プロバイダー (20 個以上のテスト済み)
出力ガバナンス ループ
すべてのモデル応答は以下を通過します。
コントラクト解析 - 応答は Beast.action_intent.v1 に準拠する必要があります
アンカー解像度 — アンカー

or_ref フィールドは正確なコードの位置を解決します。コピー＆ペーストによる書き込みは禁止
パスの検証 - 許可されたパス外の書き込みはコンパイル前に拒否されます。
ローカルパッチのコンパイル - ActionIR → ResolvedAction → ステージングされたファイルの書き込み
サンドボックス検証 - コンパイルされたパッチは、ディスクコミットの前に pytest に対して実行されます。
修復 - 検証が失敗した場合、ローカル検証者は諦める前に修復を試みます。
法医学記録 - すべての結果 (クリーン、修復、拒否) がクロニクルに書き込まれます
プロバイダー固有の出力プロファイルはモデルの癖を処理します。NVIDIA NIM は refs_only=True を取得します。 HuggingFace は Repair_attempts=2 を取得します。
レイヤー
名前
内容
L0
メタルール
支出上限、シェル許可リスト、ブロックされたパス - 不変
L1
インサイトインデックス
セッション状態、キャッシュハンドル、回線状態
L2
ワークスペースグラフ
シンボルマップ、依存関係エッジ、セマンティックチャンク
L3
スキルツリー
推進および検証されたワークフローとルートカード
L4
法医学アーカイブ
追加専用クロニクル — すべてのリクエスト、すべての結果
インストール
git clone https://github.com/Byron2306/EdgeK-BEAST
cdエッジK-BEAST
pip install -r 要件.txt
オプション (セマンティック RAG、大型 ML ホイール):
pip install -r 要件-semantic.txt
オプション (LiteLLM プロキシ サポート):
pip install -r 要件-litellm.txt
ゲートウェイを起動します。
uvicorn app.main:app --host 0.0.0.0 --port 8005
コーディング エージェントにプロバイダーではなく BEAST を直接指定します。
# OpenAI対応（カーソル、クロードコードなど）
エクスポート OPENAI_BASE_URL=http://localhost:8005/v1
# 人類互換
エクスポート ANTHROPIC_BASE_URL=http://localhost:8005
プロバイダーのセットアップ
使用するプロバイダーを設定します。
エクスポート HF_TOKEN= ' ... '
エクスポート HF_INFERENCE_BASE_URL= ' https://router.huggingface.co/v1 '
エクスポート OPENROUTER_API_KEY= ' ... '
エクスポート GEMINI_API_KEY= ' ... '
エクスポート NVIDIA_API_KEY= ' ... '
エクスポート COHERE_API_KEY= ' ... '
博覧会

rt MISTRAL_API_KEY= ' ... '
# ローカル
エクスポート LOCAL_NIM_BASE_URL= ' http://localhost:8000/v1 '
BEAST は、フィットネス マップに従ってプロバイダー間でルーティング、管理、フォールバックを行います。設定していないプロバイダーは完全にスキップされます。
# ゲートウェイの健全性
GET /健康
GET /edgek/state
# BEAST Cockpit (ライブ運用ダッシュボード)
GET /ui
# 推論 (ドロップイン置換)
POST /v1/chat/completions # OpenAI 互換
POST /v1/messages # Anthropic 互換
POST /hf/v1/chat/completions # HuggingFace ルーター
POST /litellm/v1/chat/completions # LiteLLM プロキシ
# コンテキストとワークスペース
POST /edgek/tools/intercept # セマンティックツール呼び出しインターセプト
GET /edgek/workspace # ワークスペースグラフの状態
POST /edgek/workspace/index # リポジトリにインデックスを付ける
# 予算と実行時間
GET /edgek/runtime/state
GET /edgek/runtime/attempts
POST /edgek/runtime/circuit-breakers/{provider}/reset
# MCP ブローカー
POST /edgek/mcp/evaluate
POST /edgek/mcp/execute
GET /edgek/mcp/audit
# スキルと昇進
/edgek/skills/昇進候補者を取得する
POST /edgek/スキル/プロモート
# エンタープライズ
POST /edgek/enterprise/チーム
POST /edgek/enterprise/virtual-keys
GET /edgek/enterprise/observability
API ドキュメントの完全なエンドポイント リファレンス。
events/default.yaml はすべてを制御します。
プロバイダーおよびチームごとの支出上限とトークン予算
シェルコマンドのホワイトリストとブロックリスト
ツールの遅延学習パラメータ
自分でベンチマークを実行する
# 決定論的ベンチマーク (API 呼び出しは必要ありません)
パイソンパス=。 python3 benchmarks/run_benchmark.py --lanes all --tasks 10
# ライブプロバイダーベンチマーク
パイソンパス=。 python3 ベンチマーク/run_live_benchmark.py --providers hf、openrouter、cohere
# プロバイダー エッジの比較 (クラウドとローカル NIM)
パイソンパス=。 python3 benchmarks/provider_edge_compare.py --repeats 3
結果は benchmark/results/ に書き込まれます。
BEAST は LiteLLM を生成します

アクティブなポリシーから直接 Nginx 構成を取得します。
パイソンパス=。 python3 scripts/generate_deploy_configs.py --out デプロイ/生成
Nginx は /tool-calls/* を BEAST のセマンティック インターセプターにルーティングします。ファイル読み取りリクエストは完全なソース ファイルではなく、上位 3 つの関連スニペットを返します。
GitHub ツール呼び出し、Postgres 統合、プロンプト キャッシュ キープアライブ セットアップを含む完全な Runbook については、deployment_integrations.md を参照してください。
LLM プロバイダーを置き換えるものではありません。エージェントとプロバイダー間のトラフィックを管理します。
ほとんどのタスクで発生する遅延は発生しません。出力ガバナンスにより、ローカルでマイクロ秒が追加されます。プロバイダーの遅延が支配的です。
GPUは必要ありません。ガバナンスとコンパイルのパイプライン全体が CPU 上で実行されます。
自宅に電話はしません。ワークスペース グラフ、予算台帳、フォレンジック アーカイブ、スキル ツリーなど、すべてがローカル SQLite の追加専用ファイルです。
積極的な開発。コアガバナンスパイプライン (インプットエコノミー + アウトプット契約 + ローカル検証) は安定しており、ベンチマークされています。 V2 ロードマップは、Chronicle エンジン、ルート カード、スキル プロモーション ループに焦点を当てています。 BEAST_V2_ROADMAP.md を参照してください。
貢献、問題、プロバイダーのベンチマーク結果を歓迎します。
エージェント コーディング ツールの管理された出力ゲートウェイ — 出力コントラクトを強制し、非準拠のパッチを修復し、どのツール呼び出しが価値があるかを学習します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 ギット

[切り捨てられた]

## Original Extract

Governed output gateway for agentic coding tools — enforces output contracts, repairs non-compliant patches, and learns which tool calls are worth making. - Byron2306/EdgeK-BEAST

GitHub - Byron2306/EdgeK-BEAST: Governed output gateway for agentic coding tools — enforces output contracts, repairs non-compliant patches, and learns which tool calls are worth making. · GitHub
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
Byron2306
/
EdgeK-BEAST
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits app app benchmarks benchmarks bin bin deploy/ generated deploy/ generated docs docs policies policies scripts scripts tests tests vscode-extension vscode-extension .gitignore .gitignore BEAST UI.png BEAST UI.png BEAST logo.png BEAST logo.png BEAST mascot transparent.png BEAST mascot transparent.png BEAST mascot.png BEAST mascot.png EdgeK BEAST VS CODE IDE.md EdgeK BEAST VS CODE IDE.md EdgeK BEAST VS CODE IDE.txt EdgeK BEAST VS CODE IDE.txt EdgeK_BEAST_Meta_Optimization_Whitepaper.md EdgeK_BEAST_Meta_Optimization_Whitepaper.md Gemini_Generated_Image_z6vjayz6vjayz6vj.png Gemini_Generated_Image_z6vjayz6vjayz6vj.png README.md README.md README_BEAST_UI_CONTROLS_PATCH.md README_BEAST_UI_CONTROLS_PATCH.md README_BEAST_UI_PATCH.md README_BEAST_UI_PATCH.md README_BEAST_UI_SPP_PATCH.md README_BEAST_UI_SPP_PATCH.md conftest.py conftest.py pytest.ini pytest.ini requirements-integrations.txt requirements-integrations.txt requirements-litellm.txt requirements-litellm.txt requirements-semantic.txt requirements-semantic.txt requirements.txt requirements.txt View all files Repository files navigation
BEAST - Broker for Efficient Agentic Systems and Tooling
Governed output gateway for agentic coding tools.
BEAST sits between your AI coding agent (Cursor, Claude Code, VS Code Copilot) and any LLM provider. It governs what goes in and what comes out — enforcing output contracts, repairing non-compliant patches before they touch your filesystem, and learning which tool calls are worth making.
AI coding agents are not careful. They read entire files when they need three lines. They write to paths they shouldn't. They spend your token budget on redundant lookups. When a provider returns malformed JSON, they fail silently or corrupt your code.
Input governance — context compression, tool laziness learning, budget enforcement, circuit breakers
Output governance — every model response is parsed against a typed output contract ( beast.action_intent.v1 ) before anything touches disk. Non-compliant patches are repaired locally and verified. If verification fails, nothing is written.
Deterministic — 10 tasks, 5 lanes
Lane
Completed
Median tokens
vs raw
Raw (no BEAST)
0 / 10
47,661
—
Context only
0 / 10
44
−99.9%
RAG
8 / 10
296
−99.4%
RAG + Tools
10 / 10
326
−99.3%
Full BEAST
10 / 10
390
−99.2%
Raw context hits the token budget before the model can reason about the scoped problem. BEAST completes 100% of tasks at under 400 tokens, verified by passing pytest suites.
Live providers — 192 tasks across 20 provider routes
Result
Count
BEAST end-to-end completions
192 / 192
Clean provider completions
36 / 192
BEAST-rescued completions
156 / 192
79% of raw provider outputs were non-compliant, malformed, or incomplete. BEAST rescued every one of them. Without output governance, those 156 tasks would have silently failed or written corrupted patches.
Rank
Provider
Role
Clean
Fitness
Latency
1
ovhcloud
candidate patch provider
5/10
0.663
14s
2
puter_deepseek
candidate patch (high latency)
4/10
0.619
13s
3
cohere
candidate patch provider
4/10
0.614
6.7s
4
deepinfra
candidate patch (high latency)
4/10
0.612
32s
5
huggingface
rescue-backed action IR
3/10
0.583
1.6s
6
nscale
rescue-backed action IR
3/10
0.581
7.8s
7
mistral
rescue-backed (Codestral)
2/10
0.545
4.1s
8
openrouter
fast rescue-backed action IR
2/10
0.544
3.8s
9
sambanova
fast rescue-backed action IR
1/10
0.512
3.0s
10
cloudflare
edge / microtask
1/10
0.483
2.1s
11–14
cerebras , featherless , nvidia_nim , gemini
scout / selector
0–2/10
0.33–0.42
varies
15–16
groq , llm7
scout only
0/10
0.23
fast
17–18
aion_labs , novita
rate-limited / rescue
1/10
0.39–0.51
varies
19–20
hyperbolic , fal
do not use (auth/billing)
0/10
—
—
Notable findings:
Puter-routed DeepSeek achieved 4 clean passes on a free proxied route — matching paid providers. BEAST can make unconventional free routes production-viable through governance.
LLM7 returned valid JSON on 100% of tasks but passed the output schema on only 10%. Without an output governor, it looks like it's working. It isn't.
NVIDIA NIM failed the output contract on every task. BEAST repaired and rescued both targeted tasks. Zero silent failures.
DeepInfra observed cost: ~$0.000332 per verified, governed code fix.
Coding agent (Cursor / Claude Code / VS Code)
│
▼
┌─────────────────────────────────────────┐
│ BEAST Gateway │
│ │
│ Input side Output side │
│ ───────── ─────────── │
│ Context economy Output contract │
│ Tool laziness Local verifier │
│ Budget ledger Patch compiler │
│ Circuit breakers Anchor resolver │
│ Workspace graph Repair engine │
│ MCP broker Sandbox validator │
│ │
│ Memory: L0 policy → L4 forensic archive│
└─────────────────────────────────────────┘
│
▼
Any LLM provider (20+ tested)
The output governance loop
Every model response passes through:
Contract parse — response must conform to beast.action_intent.v1
Anchor resolution — anchor_ref fields resolve to exact code locations; no copy-paste writes
Path validation — writes outside allowed paths are rejected before compilation
Local patch compile — ActionIR → ResolvedAction → staged file writes
Sandbox verification — compiled patches run against pytest before disk commit
Repair — if verification fails, the local verifier attempts repair before giving up
Forensic record — every outcome (clean, repaired, rejected) is written to the Chronicle
Provider-specific output profiles handle model quirks: NVIDIA NIM gets refs_only=True ; HuggingFace gets repair_attempts=2 .
Layer
Name
Contents
L0
Meta Rules
Spend caps, shell allowlists, blocked paths — immutable
L1
Insight Index
Session state, cache handles, circuit state
L2
Workspace Graph
Symbol maps, dependency edges, semantic chunks
L3
Skill Tree
Promoted, verified workflows and route cards
L4
Forensic Archive
Append-only Chronicle — every request, every outcome
Installation
git clone https://github.com/Byron2306/EdgeK-BEAST
cd EdgeK-BEAST
pip install -r requirements.txt
Optional (semantic RAG, large ML wheels):
pip install -r requirements-semantic.txt
Optional (LiteLLM proxy support):
pip install -r requirements-litellm.txt
Start the gateway:
uvicorn app.main:app --host 0.0.0.0 --port 8005
Point your coding agent at BEAST instead of your provider directly:
# OpenAI-compatible (Cursor, Claude Code, etc.)
export OPENAI_BASE_URL=http://localhost:8005/v1
# Anthropic-compatible
export ANTHROPIC_BASE_URL=http://localhost:8005
Provider setup
Set whichever providers you use:
export HF_TOKEN= ' ... '
export HF_INFERENCE_BASE_URL= ' https://router.huggingface.co/v1 '
export OPENROUTER_API_KEY= ' ... '
export GEMINI_API_KEY= ' ... '
export NVIDIA_API_KEY= ' ... '
export COHERE_API_KEY= ' ... '
export MISTRAL_API_KEY= ' ... '
# Local
export LOCAL_NIM_BASE_URL= ' http://localhost:8000/v1 '
BEAST will route, govern, and fall back across providers according to the fitness map. Providers you haven't configured are skipped cleanly.
# Gateway health
GET /health
GET /edgek/state
# BEAST Cockpit (live ops dashboard)
GET /ui
# Inference (drop-in replacements)
POST /v1/chat/completions # OpenAI-compatible
POST /v1/messages # Anthropic-compatible
POST /hf/v1/chat/completions # HuggingFace router
POST /litellm/v1/chat/completions # LiteLLM proxy
# Context and workspace
POST /edgek/tools/intercept # Semantic tool-call interception
GET /edgek/workspace # Workspace graph state
POST /edgek/workspace/index # Index a repository
# Budget and runtime
GET /edgek/runtime/state
GET /edgek/runtime/attempts
POST /edgek/runtime/circuit-breakers/{provider}/reset
# MCP broker
POST /edgek/mcp/evaluate
POST /edgek/mcp/execute
GET /edgek/mcp/audit
# Skills and promotion
GET /edgek/skills/promotion-candidates
POST /edgek/skills/promote
# Enterprise
POST /edgek/enterprise/teams
POST /edgek/enterprise/virtual-keys
GET /edgek/enterprise/observability
Full endpoint reference in the API docs .
policies/default.yaml controls everything:
Spend caps and token budgets per provider and per team
Shell command allowlists and blocklists
Tool laziness learning parameters
Running the benchmark yourself
# Deterministic benchmark (no API calls needed)
PYTHONPATH=. python3 benchmarks/run_benchmark.py --lanes all --tasks 10
# Live provider benchmark
PYTHONPATH=. python3 benchmarks/run_live_benchmark.py --providers hf,openrouter,cohere
# Provider edge compare (cloud vs local NIM)
PYTHONPATH=. python3 benchmarks/provider_edge_compare.py --repeats 3
Results are written to benchmarks/results/ .
BEAST generates LiteLLM and Nginx configs directly from your active policy:
PYTHONPATH=. python3 scripts/generate_deploy_configs.py --out deploy/generated
Nginx routes /tool-calls/* into BEAST's semantic interceptor — file read requests return the top 3 relevant snippets instead of full source files.
See deployment_integrations.md for the full runbook including GitHub tool calls, Postgres integration, and prompt-cache keepalive setup.
It does not replace your LLM provider. It governs the traffic between your agent and your provider.
It does not add latency you'll notice for most tasks. Output governance adds microseconds locally; provider latency dominates.
It does not require a GPU. The entire governance and compilation pipeline runs on CPU.
It does not phone home. Everything — workspace graph, budget ledger, forensic archive, skill tree — is local SQLite and append-only files.
Active development. Core governance pipeline (input economy + output contracts + local verification) is stable and benchmarked. V2 roadmap focuses on the Chronicle engine, route cards, and skill promotion loop. See BEAST_V2_ROADMAP.md .
Contributions, issues, and provider benchmark results welcome.
Governed output gateway for agentic coding tools — enforces output contracts, repairs non-compliant patches, and learns which tool calls are worth making.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 Git

[truncated]
