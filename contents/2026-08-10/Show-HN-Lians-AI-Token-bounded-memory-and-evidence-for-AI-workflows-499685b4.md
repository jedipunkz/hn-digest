---
source: "https://github.com/Lians-ai/Lians"
hn_url: "https://news.ycombinator.com/item?id=49248029"
title: "Show HN: Lians AI, Token-bounded memory and evidence for AI workflows"
article_title: "GitHub - Lians-ai/Lians: Open-source decision evidence and reconstruction for production AI agents. Bitemporal memory, Decision Receipts, runtime gates, and blast-radius investigation. · GitHub"
author: "ebeirne"
captured_at: "2026-08-10T19:49:18Z"
capture_tool: "hn-digest"
hn_id: 49248029
score: 1
comments: 0
posted_at: "2026-08-10T18:54:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lians AI, Token-bounded memory and evidence for AI workflows

- HN: [49248029](https://news.ycombinator.com/item?id=49248029)
- Source: [github.com](https://github.com/Lians-ai/Lians)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T18:54:01Z

## Translation

タイトル: Show HN: Lians AI、トークン境界メモリと AI ワークフローの証拠
記事のタイトル: GitHub - Lians-ai/Lians: 本番 AI エージェントのためのオープンソースの意思決定証拠と再構築。バイテンポラル記憶、意思決定受領書、実行時ゲート、爆発半径調査。 · GitHub
説明: オープンソースの意思決定証拠と本番 AI エージェントの再構築。バイテンポラル記憶、意思決定受領書、実行時ゲート、爆発半径調査。 -リアンアイ/リアンズ
HN テキスト: メモリを改善し、クロード コードとコーデックスのトークン使用量を削減するためにこれを構築しました。

記事本文:
GitHub - Lians-ai/Lians: 本番 AI エージェントのためのオープンソースの意思決定証拠と再構築。バイテンポラル記憶、意思決定受領書、実行時ゲート、爆発半径調査。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
リアンアイ
/
リアン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
279 コミット 279 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github Agentmem

Agentmem デモ デモ ドキュメント ドキュメント統合 統合 k8s k8s 出力/ pdf 出力/ pdf パッケージ/ Agentmem-backtest-check パッケージ/ Agentmem-backtest-check ペーパー/ Regulated-memory-eval ペーパー/ Regulated-memory-eval プラグイン プラグイン スクリプト スクリプト sdk/ python sdk/ Python スキル スキル tmp/ pdfs/ Regulatory-memory-eval tmp/ pdfs/regulated-memory-eval .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff Dockerfile Dockerfile Dockerfile.glama Dockerfile.glama LAUNCHGUIDE.md LAUNCHGUIDE.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md バージョン バージョン fly.toml fly.toml lians-gtm-plan.md lians-gtm-plan.md lians-logo-white-on-black.png lians-logo-white-on-black.png llms-install.md llms-install.md pyproject.toml pyproject.toml render.yaml render.yaml server.json server.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ウェブサイト
-
ドキュメント
-
インストール
-
クイックスタート
-
スターリアン
再現可能なベンチマーク証拠とオフライン品質ゲート
RIAD-1: 意思決定再構築ベンチマーク
· CI 領収書
Lians はクロスプラットフォームです
規制された AI の意思決定証拠と再構築レイヤー。それは与える
コンプライアンス、モデルリスク、オペレーショナルリスクの各チームがエージェントの役割を 1 つの記録に記録
何を取得したか、どのポリシーが管理しているか、どのツールが実行されているか、誰が実行しているかを知っていました。
それを見直して、その後何が変わったのか。
耐久性のある堀は中立です。企業は、Bedrock、Azure 全体でエージェントを実行できます
OpenAI、Anthropic ダイレクト、オープンソース ランタイムを移植性を維持しながら実現
すべてのプロバイダーの外部にある証拠記録。
すべての書き込みは管理された一時記録として保存され、
型付きメモリアーティファクト。すべてのリコールは、高速、詳細、または再構築で実行できます。
モードに設定し、自動的にバインドできるコンテンツアドレス付きのレシートを返します。
決定封筒。参照
決定証拠

そして再建、
規範的完全性グレード、
証拠パック署名キーの保管、
管理されたメモリエンジンと
再現可能な証拠ゲート。
このプラットフォームは 1 つの証拠ワークフローを公開します。
キャプチャ : デシジョン エンベロープを開き、メモリ、トレース、ポリシーをバインドします
決定、プロンプト、ツール、アクションが発生したときの人間によるレビュー。
再構築 : ポイントインタイムの知識と実行パスを再現します。
正確な決定論的な再生が不可能な場合。
Verify : すべての決定を記録済み、再構築可能、検証可能、または
不足しているすべての要件に名前を付けて再生可能。
監視: ソース、ポリシー、またはモデルが変更されたときに、公開されているすべての脆弱性を特定します。
決定し、爆発半径警報を発します。
メモリは依然として中核的な証拠ソースであり、パフォーマンスの原始要素です。それはそうではありません
それ自体は商業カテゴリーです。
エージェントの記憶は記録を失うことなく改善されるはずです
Lians は、エージェントに事実、コンテキスト、意思決定、結果にわたる永続的な記憶ループを提供します。
そして授業の復習をしました。 Memory 製品はコンテキストを最新かつ有用に保ちます。の
製品の動作と監視をオープンで検証可能なイベント形式で記録します。
ほとんどのメモリ層は、保存と取得で停止します。 Lians は次のようなチーム向けに構築されています。
また、エージェントが何を知っていたのか、いつ知ったのか、その事実がどこから来たのかを知る必要もあります。
どのような結果が続いたか、誰が閲覧を許可されたか、古いか消去されたか
コンテンツは将来の文脈から切り離されました。
それが、メモリのデモとチームが信頼できるメモリ システムとの間にあるギャップです。
特に金融、医療、法律環境における生産。
調節された記憶が証明しなければならないこと
一般的なエージェントの記憶により、パーソナライゼーションとリコールが最適化されます。規制対象物質
メモリには別の役割があります。エージェントのコンテキストを正確かつ最新に保つ必要があります。
分離されており、再現可能であり、レビュー中は防御可能です。
リアンズがデザインした

組織内で重要な障害モードについては、次のとおりです。
古い事実の汚染 - 古い料金、古いガイダンス、古い薬の投与量、
古い損害賠償額の見積りや、クライアントの古い事実が黙って文脈に入り込んではなりません。
ポイントインタイムの再構築 - 検査者、臨床医、パートナー、またはリスク
委員会は、エージェントが特定のタイムスタンプで何を知っていたかを尋ねる場合があります。
情報障壁 - 1 つのデスク、ケア チーム、または問題チームは読んではなりません
アプリケーション層のバグにより、別のチームのメモリが失われます。
監査存続による消去 - プライベート コンテンツは削除せずに削除できる必要があります。
保管記録、監査ハッシュ、法的保存証拠の破壊。
関係コンプライアンスチェック - 利益相反、関連当事者
露出および紹介ネットワークはグラフの質問であり、単純なベクトル検索ではありません。
ランタイム ベンダーが自社のクラウドについて説明します。リアンズは移植可能な決定を保持します
それらすべてにわたる証拠。
垂直
リアンズが証明したこと
プロダクトプリミティブ
金融機関
古い事実や将来の事実が決定に影響を与えることはありません。机上の障壁が保持されている。監査状態は再構築可能です
バイテンポラルリコール、バックテスト汚染チェック、SEC/FINRA監査エクスポート、RLS情報バリア、関連者グラフパス
医療機関
PHI アクセスは範囲指定されています。ケアチームの記憶は再構築可能です。患者の消去は証明可能
被験者ごとの暗号化、暗号シュレッド証明書、HIPAA セーフガード マッピング、ケア ネットワーク グラフ、エアギャップ モード
法的機関
物質の壁は保持されています。権限の遮断は再現可能です。保管管理は消去されても存続する
事項レベルの障壁、特権日のリコール_アット、監査の再構築、利益相反グラフのパス
調達および技術検討資料:
MCP - あらゆる AI クライアントのネイティブ ツール
Lians は公式 MCP レジストリにリストされています。任意の MCP 互換ホスト - Claude Desktop、Cursor、VS Code、Windsurf など - c

ローカルの永続メモリをすぐに使用するか、ホストされている Lians サーバーに接続します。ローカル モードでは、SDK コード、カスタム アダプター、Docker サービス、URL、または API キーは必要ありません。
エージェントは次の 8 つのツールを自動的に取得します。
クロード デスクトップ/カーソル/ウィンドサーフィン
claude_desktop_config.json (または同等の MCP 構成) に追加します。
{
"mcpサーバー": {
"リアン" : {
"コマンド" : " uvx " 、
"args" : [ " --from " 、 " lians-sdk[mcp] " 、 " lians-mcp " ]
}
}
}
クライアントを再起動すると、Lians メモリ ツールがすぐに表示されます。ローカル モードは ~/.lians/mcp.db まで存続します。代わりにホスト型デプロイメントを使用するには、 LIANS_URL 、 LIANS_API_KEY 、およびオプションで LIANS_AGENT_ID を設定します。
uvx --from ' lians-sdk[mcp] ' lians-mcp
ローカル モードでは環境変数は必要ありません。リモートサーバーを使用するには、 LIANS_URL 、 LIANS_API_KEY 、およびオプションで LIANS_AGENT_ID を設定します。
pip install lians-sdk[local] # SQLite と実際のローカル セマンティック埋め込み、Docker なし
lians から LocalLiansClient をインポート
from datetime import datetime 、タイムゾーン
mem = LocalLiansClient()
メモリ。追加(
エージェント ID = "アナリスト-1" ,
content = "NVDA の 2026 年度の収益見通しは 400 億ドルに引き上げられました" ,
events_time = datetime ( 2025 , 11 , 19 , 16 , tzinfo = タイムゾーン . utc )、
メタデータ = { "ティッカー" : "NVDA" , "メトリック" : "収益ガイダンス" },
)
# 置き換えられたファクトは DB 層で除外されます。LLM には決して到達しません
結果 = mem 。リコール (agent_id = "analyst-1" 、クエリ = "NVDA 収益ガイダンス" )
# 計画と研究のためのより深い多面的想起
結果 = mem 。思い出してください（
エージェント ID = "アナリスト-1" ,
query = "ガイダンスの何が変更され、その理由は何ですか?" 、
モード = "ディープ" 、
)
# 時点: 3 月 1 日に私たちは何を知っていましたか? (コンプライアンスグレードの回答)
結果 = mem 。リコール_at (
エージェント ID = "アナリスト-1" ,
query = "NVDA 収益ガイダンス" ,
as_of = datetime ( 2025 , 3 , 1 , tzinfo = タイムゾーン . utc )、
)
# すべての解像度

ult には、receipt_sha256、provinance_coverage、および
# 解決されたサービス提供モードとレイテンシ バジェット。
1 行でホストされたサーバーに切り替えます: from lians import LiansClient as LocalLiansClient
from datetime import datetime 、タイムゾーン
lians から AsyncLiansClient をインポート
AsyncLiansClient (base_url = LIANS_URL 、api_key = LIANS_API_KEY) を lians として非同期:
エンベロープ = リアンを待ちます。 open_decion_envelope (
Agent_id = "アンダーライター-1" ,
Decision_type = "credit_application" ,
レジーム = "ECOA_REG_B" ,
completeness_profile = "regulated_recordkeeper" ,
Knowledge_as_of = datetime 。現在 (タイムゾーン .utc)、
)
context = await lians 。思い出してください（
Agent_id = "アンダーライター-1" ,
query = "確認された申請者の収入" ,
Decision_envelope_id = エンベロープ [ "id" ],
)
封印 = リアンを待ちます。シール_決定_エンベロープ (
エンベロープ [ "id" ]、
結果 = "manual_review" ,
決定済み_at = datetime 。現在 (タイムゾーン .utc)、
input_hash = INPUT_SHA256 、
出力ハッシュ = OUTPUT_SHA256 、
)
# 過剰な要求は禁止: 欠落している要件には、ブロックするグレードの名前が付けられます。
印刷（封印[「完全性」]）
エージェント ハーネス — ドロップイン メモリ ループ
LiansMemoryHarness は、すべてのメモリ拡張エージェントが必要とする 2 つの操作をラップします。
remember-before と remember-after — コンプライアンスの範囲を指定して 1 つのオブジェクトにまとめる
(件名、ソース、イベント時間、情報バリア) 規制された展開が必要です。
任意の同期クライアント ( LiansClient または LocalLiansClient ) および任意のモデルで動作します。
lians import LiansClient 、 LiansMemoryHarness から
harness = LiansMemoryHarness ( mem 、agent_id = "research-desk" 、domain = "finance" )
# 1 回の呼び出し: コンテキストを呼び出し、モデルを実行し、応答を永続化します。
答え＝ハーネス。 run_turn (
「NVDA の現在の収益見通しはどのくらいですか?」 、
生成 = ラムダ コンテキスト 、クエリ : call_model ( f" { コンテキスト } \n \n ユーザー: { クエリ } " )、
)
#

または、各ステップを制御します。
コンテキスト = ハーネス 。 remember_context ( "NVDA 収益ガイダンス" ) # 注入する準備ができました
ハーネス。覚えておいてください ( "デスクノート: 現在のガイダンスは $40B" ) # ターン後に書きます
規制されたスコープは、すべての書き込みを 1 つのデータ主体と情報バリアに結び付けます。
ハーネス = LiansMemoryHarness (
mem 、agent_id = "ケアチーム-3" 、
subject_id = "MRN-00042" , # 件名ごとのキー — 暗号化シュレッド ターゲット
Barrier_group = "oncology" , # 情報バリアタグ
ドメイン = "ヘルスケア" 、
)
実行可能なエンドツーエンドのデモ：agentmem/examples/harness_demo.py 。
関係グラフ — 本質的に関係性のあるコンプライアンスに関する質問
一部のコンプライアンス チェックはグラフ クエリです。リアンはバイテンポラル関係を保存します
事実とエッジ - 同じ監査チェーン、同じ情報障壁、グラフなし
データベース — その時点で回答できるようになります。
法務 — 利益相反の到達可能性 (ABA 1.7/1.9): 弁護士です
敵対者と関係があるのか？
財務 — 関連者/受益者所有権 (SEC、AML/KYC):
制限されたエンティティから N ホップ以内に相手はいますか?
ヘルスケア — ケアネットワーク/紹介パターン (反キックバック) 分析。
メモリ。関連 ( "分析者-1" 、 src_entity = "弁護士" 、 rel_type = "代理人" 、
dst_entity = "ClientX" 、偶数

[切り捨てられた]

## Original Extract

Open-source decision evidence and reconstruction for production AI agents. Bitemporal memory, Decision Receipts, runtime gates, and blast-radius investigation. - Lians-ai/Lians

we built this for improved memory and reduced token usage for claude code and codex

GitHub - Lians-ai/Lians: Open-source decision evidence and reconstruction for production AI agents. Bitemporal memory, Decision Receipts, runtime gates, and blast-radius investigation. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Lians-ai
/
Lians
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
279 Commits 279 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github agentmem agentmem demo demo docs docs integrations integrations k8s k8s output/ pdf output/ pdf packages/ agentmem-backtest-check packages/ agentmem-backtest-check paper/ regulated-memory-eval paper/ regulated-memory-eval plugins plugins scripts scripts sdk/ python sdk/ python skills skills tmp/ pdfs/ regulated-memory-eval tmp/ pdfs/ regulated-memory-eval .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff Dockerfile Dockerfile Dockerfile.glama Dockerfile.glama LAUNCHGUIDE.md LAUNCHGUIDE.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md VERSION VERSION fly.toml fly.toml lians-gtm-plan.md lians-gtm-plan.md lians-logo-white-on-black.png lians-logo-white-on-black.png llms-install.md llms-install.md pyproject.toml pyproject.toml render.yaml render.yaml server.json server.json View all files Repository files navigation
Website
-
Docs
-
Install
-
Quickstart
-
Star Lians
Reproducible benchmark evidence and offline quality gates
RIAD-1: decision reconstruction benchmark
· CI receipts
Lians is the cross-platform
decision evidence and reconstruction layer for regulated AI . It gives
compliance, model-risk, and operational-risk teams one record of what an agent
knew, what it retrieved, which policy governed it, which tools ran, who
reviewed it, and what changed later.
The durable moat is neutrality. A firm can run agents across Bedrock, Azure
OpenAI, Anthropic direct, and open-source runtimes while keeping one portable
evidence record outside every provider.
Every write is preserved as a governed temporal record and compiled into a
typed memory artifact. Every recall can run in fast , deep , or reconstruct
mode and returns a content-addressed receipt that can bind automatically to a
Decision Envelope. See
decision evidence and reconstruction , the
normative completeness grades ,
Evidence Pack signing key custody , the
governed memory engine and
reproducible evidence gates .
The platform exposes one evidence workflow:
Capture : open a Decision Envelope and bind memory, traces, policy
decisions, prompts, tools, and human review as the action happens.
Reconstruct : reproduce the point-in-time knowledge and execution path even
when exact deterministic replay is impossible.
Verify : grade every decision as Recorded, Reconstructable, Verifiable, or
Replayable, with every missing requirement named.
Monitor : when a source, policy, or model changes, identify every exposed
decision and emit a blast-radius alert.
Memory remains a core evidence source and performance primitive. It is not the
commercial category by itself.
Agent memory should improve without losing the record
Lians gives agents a durable memory loop across facts, context, decisions, outcomes,
and reviewed lessons. The Memory product keeps context current and useful; the
Records product captures behavior and oversight in an open, verifiable event format.
Most memory layers stop at storage and retrieval. Lians is built for teams that
also need to know what the agent knew, when it knew it, where the fact came from,
which outcomes followed, who was allowed to see it, and whether stale or erased
content was kept out of future context.
That is the gap between a memory demo and a memory system teams can trust in
production, especially in financial, medical, and legal environments.
What regulated memory must prove
Generic agent memory optimizes for personalization and recall. Regulated agent
memory has a different job: it must keep the agent's context correct, current,
segregated, reproducible, and defensible under review.
Lians is designed for the failure modes that matter in institutions:
Stale fact contamination - old rates, old guidance, old medication doses,
old damages estimates, or old client facts must not silently enter context.
Point-in-time reconstruction - an examiner, clinician, partner, or risk
committee may ask what the agent knew at a specific timestamp.
Information barriers - one desk, care team, or matter team must not read
another team's memory because of an application-layer bug.
Erasure with audit survival - private content must be removable without
breaking custody records, audit hashes, or legal retention evidence.
Relational compliance checks - conflicts of interest, related-party
exposure, and referral networks are graph questions, not plain vector search.
Runtime vendors explain their own cloud. Lians preserves portable decision
evidence across all of them.
Vertical
What Lians proves
Product primitives
Financial institutions
No stale or future facts influenced a decision; desk barriers held; audit state is reconstructable
Bitemporal recall, backtest contamination checks, SEC/FINRA audit export, RLS information barriers, related-party graph paths
Healthcare organizations
PHI access is scoped; care-team memory is reconstructable; patient erasure is provable
Per-subject encryption, crypto-shred certificates, HIPAA safeguard mapping, care-network graph, air-gap mode
Legal institutions
Matter walls held; privilege cutoffs are reproducible; chain-of-custody survives erasure
Matter-level barriers, recall_at for privilege dates, audit reconstruction, conflict-of-interest graph paths
Procurement and technical review materials:
MCP - Native tool in any AI client
Lians is listed on the official MCP Registry . Any MCP-compatible host - Claude Desktop, Cursor, VS Code, Windsurf, and others - can use local persistent memory immediately or connect to a hosted Lians server. No SDK code, custom adapter, Docker service, URL, or API key is required for local mode.
Your agents get eight tools automatically:
Claude Desktop / Cursor / Windsurf
Add to your claude_desktop_config.json (or equivalent MCP config):
{
"mcpServers" : {
"lians" : {
"command" : " uvx " ,
"args" : [ " --from " , " lians-sdk[mcp] " , " lians-mcp " ]
}
}
}
Restart your client and Lians memory tools appear immediately. Local mode persists to ~/.lians/mcp.db . To use a hosted deployment instead, set LIANS_URL , LIANS_API_KEY , and optionally LIANS_AGENT_ID .
uvx --from ' lians-sdk[mcp] ' lians-mcp
No environment variables are needed for local mode. Set LIANS_URL , LIANS_API_KEY , and optionally LIANS_AGENT_ID to use a remote server.
pip install lians-sdk[local] # SQLite plus real local semantic embeddings, no Docker
from lians import LocalLiansClient
from datetime import datetime , timezone
mem = LocalLiansClient ()
mem . add (
agent_id = "analyst-1" ,
content = "NVDA FY2026 revenue guidance raised to $40B" ,
event_time = datetime ( 2025 , 11 , 19 , 16 , tzinfo = timezone . utc ),
metadata = { "ticker" : "NVDA" , "metric" : "revenue_guidance" },
)
# Superseded facts are excluded at the DB layer — never reach the LLM
results = mem . recall ( agent_id = "analyst-1" , query = "NVDA revenue guidance" )
# Deeper multi-facet recall for planning and research
results = mem . recall (
agent_id = "analyst-1" ,
query = "What changed in the guidance and why?" ,
mode = "deep" ,
)
# Point-in-time: what did we know on March 1? (compliance-grade answer)
results = mem . recall_at (
agent_id = "analyst-1" ,
query = "NVDA revenue guidance" ,
as_of = datetime ( 2025 , 3 , 1 , tzinfo = timezone . utc ),
)
# Every result includes receipt_sha256, provenance_coverage, and the
# resolved serving mode and latency budget.
Switch to the hosted server with one line: from lians import LiansClient as LocalLiansClient
from datetime import datetime , timezone
from lians import AsyncLiansClient
async with AsyncLiansClient ( base_url = LIANS_URL , api_key = LIANS_API_KEY ) as lians :
envelope = await lians . open_decision_envelope (
agent_id = "underwriter-1" ,
decision_type = "credit_application" ,
regime = "ECOA_REG_B" ,
completeness_profile = "regulated_recordkeeping" ,
knowledge_as_of = datetime . now ( timezone . utc ),
)
context = await lians . recall (
agent_id = "underwriter-1" ,
query = "verified applicant income" ,
decision_envelope_id = envelope [ "id" ],
)
sealed = await lians . seal_decision_envelope (
envelope [ "id" ],
outcome = "manual_review" ,
decided_at = datetime . now ( timezone . utc ),
input_hash = INPUT_SHA256 ,
output_hash = OUTPUT_SHA256 ,
)
# No overclaiming: every missing requirement names the grade it blocks.
print ( sealed [ "completeness" ])
Agent harness — drop-in memory loop
LiansMemoryHarness wraps the two operations every memory-augmented agent needs —
recall-before and remember-after — into one object, with the compliance scoping
(subject, source, event-time, information barrier) regulated deployments require.
Works with any sync client ( LiansClient or LocalLiansClient ) and any model.
from lians import LiansClient , LiansMemoryHarness
harness = LiansMemoryHarness ( mem , agent_id = "research-desk" , domain = "finance" )
# One call: recall context, run your model, persist the response.
answer = harness . run_turn (
"What is NVDA's current revenue guidance?" ,
generate = lambda context , query : call_model ( f" { context } \n \n User: { query } " ),
)
# Or control each step:
context = harness . recall_context ( "NVDA revenue guidance" ) # ready to inject
harness . remember ( "Desk note: guidance now $40B" ) # write after the turn
Regulated scoping ties every write to one data subject and an information barrier:
harness = LiansMemoryHarness (
mem , agent_id = "care-team-3" ,
subject_id = "MRN-00042" , # per-subject key — the crypto-shred target
barrier_group = "oncology" , # information-barrier tag
domain = "healthcare" ,
)
Runnable end-to-end demo: agentmem/examples/harness_demo.py .
Relationship graph — compliance questions that are inherently relational
Some compliance checks are graph queries. Lians stores bitemporal relationship
edges alongside facts — same audit chain, same information barriers, no graph
database — so you can answer them point-in-time:
Legal — conflict-of-interest reachability (ABA 1.7/1.9): is an attorney
connected to an adverse party?
Finance — related-party / beneficial-ownership (SEC, AML/KYC): is a
counterparty within N hops of a restricted entity?
Healthcare — care-network / referral-pattern (anti-kickback) analysis.
mem . relate ( "analyst-1" , src_entity = "Attorney" , rel_type = "represented" ,
dst_entity = "ClientX" , even

[truncated]
