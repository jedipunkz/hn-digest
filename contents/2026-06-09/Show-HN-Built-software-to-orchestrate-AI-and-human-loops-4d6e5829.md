---
source: "https://github.com/sparckix/cognitive-firm"
hn_url: "https://news.ycombinator.com/item?id=48455712"
title: "Show HN: Built software to orchestrate AI and human loops"
article_title: "GitHub - sparckix/cognitive-firm: A governance kernel for organizations that coordinate persistent human and AI-agent roles. · GitHub"
author: "Sparckix"
captured_at: "2026-06-09T04:29:29Z"
capture_tool: "hn-digest"
hn_id: 48455712
score: 2
comments: 0
posted_at: "2026-06-09T02:55:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Built software to orchestrate AI and human loops

- HN: [48455712](https://news.ycombinator.com/item?id=48455712)
- Source: [github.com](https://github.com/sparckix/cognitive-firm)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T02:55:25Z

## Translation

タイトル: Show HN: AI と人間のループを調整するソフトウェアを構築
記事のタイトル: GitHub - sparckix/cognitive-firm: 人間と AI エージェントの永続的な役割を調整する組織のためのガバナンス カーネル。 · GitHub
説明: 永続的な人間と AI エージェントの役割を調整する組織向けのガバナンス カーネル。 - sparckix/コグニティブファーム
HN テキスト: ご意見をお聞かせください。

記事本文:
GitHub - sparckix/cognitive-firm: 人間と AI エージェントの永続的な役割を調整する組織のためのガバナンス カーネル。 · GitHub
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
スパークキックス
/
認識力が強い
公共
通知
通知設定を変更するにはサインインする必要があります
追加

イオンナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
57 コミット 57 コミット .obsidian .obsidian デプロイ デプロイ distro distro docs docs orbit orbit org org スキーマ スキーマ スクリプト スクリプト src src テナント テナント テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FUTURE_AREAS.md FUTURE_AREAS.md ライセンス ライセンス MANIFEST.in MANIFEST.in Makefile Makefile NOTICE NOTICE README.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml required.txt requirements.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
永続的な人間とエージェントの役割を調整する組織のためのガバナンス カーネル。
アイデンティティ · 対象者 · アーキテクチャ · リポジトリ マップ · クイックスタート · 運用 · ステータス · 導入 · ライセンスと来歴
コグニティブファームでは、1 人のプリンシパルが永続的な AI エージェントを調整します。
入力された権限、予算の上限、人間による承認と作業領域、テナント
オーバーレイパターン。
これは、エージェント フレームワーク、LLM SDK、またはチャット クライアントではありません。それはレイヤーです
これらの上: 権限のある永続的な役割オフィス、受信箱、セッション、
シグナルと監査証跡。
運用モデルは具体的です。エージェントは由来のあるアクションを実行し、人間はそのアクションを実行します。
領収書による制限された作業、役割オフィス交換義務、および入力のみ
状態の遷移は組織を変えます。パートナー通話、制限付きシステム
チェック、コード変更、予測ミス、ガバナンス提案は異なります
異なるフォローアップ義務が生じるためです。
人間とエージェントの作業モデルは、意図的に「人間がすべてのアクションを承認する」ものではありません。
「事後の監視を伴うエージェントの自律性」ではありません。エージェントは迅速に実行できます
内側の境界

封筒。人間は、適切なアクターとして貢献するか、
裁判官。カーネルは、作業をレビュー可能にする状態遷移を記録します。
そのためには docs/human-agent-work.md から始めます
理論とその運用上のマッピング。
これは、実際の研究展開から抽出され、再利用可能なものとして書き直されました。
カーネルとテスト、サンプル、ローカル デプロイメント フィクスチャ。
2. 誰のためのもので、何を解決するのか
LangChain、AutoGen、CrewAI、Letta などのエージェント フレームワークが最強
実行グラフ、ツール呼び出し、メモリ、エージェントの対話をモデル化するとき。
このカーネルは、永続的な役割、権限、
これらのランタイムには、説明責任、人的作業、リース、監査が表面化します。
研究室、基金、または個人会社を経営する単一のプリンシパルで、限定された権限を持つ常駐エージェントが必要な場合。
小規模チーム (1 ～ 5 人) は、役割の権限、予算、監査証跡、承認がすべて git 内のファイルとして存在する 1 つの場所を必要としています。
研究者は、ログからリバース エンジニアリングするのではなく、検査、再生、フォークできるシステムを構築しています。
範囲外
なぜ
エンタープライズ RBAC / SSO
無駄のないマルチアクターロールのメンバーシップがサポートされています。完全なエンタープライズ IAM 管理、SSO プロビジョニング、およびテナント分離は、展開作業として残ります。
グラフ/モデルのランタイム
デーモンは、既存の CLI (Claude Code、codex) をディスパッチするガバナンス ランタイムです。モデル推論、グラフ再生、ノード スケジューリングは実装されません。
チャット UI
記録システムはファイルシステム + git です。 UI (Orbit、Telegram) は投影であり、真実ではありません。
エージェントプロンプトエンジニアリング
権限と役割は入力された契約であり、プロンプトのテンプレートではありません。
組織間のフェデレーション
単一のリポジトリ/単一のプリンシパルがガバナンスの単位です。
代替案を簡単に言うと
LangGraph / AutoGen / CrewAI : 作業単位がオンの場合により適合します

グラフ、クルー、またはチャットの実行。コグニティブファームは永続的な組織向けです。これらのランタイムは、ライフサイクル イベントをランタイムに投影できます。
Letta / MemGPT : 問題がエージェントのメモリにある場合に適しています。コグニティブファームは、ファイルシステム + git がメモリであると仮定します。
n8n / Zapier : SaaS API 間の接着に問題がある場合に適しています。コグニティブファームは、どの役割がどの権限の下でどの MCP ツールを起動できるかをゲートします。
コグニティブファームは、ガバナンスプロトコルとプリミティブを積み重ねた小さなセットです。
ファイルシステムに基づく記録システム。詳細については docs/PROTOCOLS.md をお読みください。
仕様。このセクションが地図です。
レイヤー
プロトコル
それが支配するもの
スペック
人間 ↔ 組織
H2A (人間対エージェント)
Telegram + Orbit + CLI サーフェス。ペース階層型の注意規律。 STOP権限
docs/プロトコル/h2a.md
役割→人間
A2H (エージェントと人間の作業調整)
成果物、受領書、エージェントのフォローアップなど、役割オフィスによって要求される限定的な人的作業
docs/protocols/a2h.md
役割 ↔ 役割
A2A (エージェント間)
型指定された AgentMessage エンベロープ、義務のライフサイクル、コンテンツに対応したアーティファクトの依存関係、サガ補償
docs/プロトコル/a2a.md
役割 ↔ 外部
MCP (モデル コンテキスト プロトコル)
エンタープライズ システム (Linear、Salesforce、ERP) への機能ゲート型送信ボックス リレー ディスパッチ
docs/プロトコル/mcp.md
横断的
委任
各役割が自律的に実行できる内容とエスカレーションによって実行できる内容に関する型付き権限契約
docs/protocols/mandate.md
プロジェクト・テナント
プロジェクト憲章
スコープの忠実度: 中心となる質問、スコープ外の境界、最終状態、予測タイプ、継承、およびアンカー プロキシ
docs/protocols/project-charter.md
不変条件
プリミティブは小さな概念的なコアとして機能します。参照
完全なマップについては docs/kernel-invariants.md を参照してください。
┌─────────────────

─────────────┐
│ 校長（人間） │
│ ↕ H2A: Telegram (モバイル) · Orbit (デスクトップ) · CLI │
│ ↕ A2H: 限定された作業リクエスト、受領書、フォローアップ │
━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ 承認 + 制限された人間の作業
┌─────────▼─────────────────┐
│ 役割（リサーチディレクター、マネージャー、ディベートランナーなど） │
│ ↕ A2A: タイプ付き封筒、義務ライフサイクル、サガ │
━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ サブプロセス (OAuth 用に環境スクラブ)
┌─────────▼─────────────────┐
│ AGENT RUNTIMES (Claude Code、コーデックス、将来のオープンソース) │
│ エージェント側のサブスクリプション割り当て・ツール用の API トークン │
━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ MCP 機能ゲート型送信ボックス リレー
┌─────────▼───────────────

────┐
│ 外部システム (Linear、Salesforce、ERP、チケット発行) │
━━━━━━━━━━━━━━━━━━━━━━━━━┘
記録システム: org/ (ファイルシステム) + git 履歴。
監査証跡 = git ログ; UI = プロジェクション。
4. リポジトリマップ
新しい採用者? docs/getting-started.md から始めます —
ゼロから統治された企業を運営するまでの最短パス ( pip install →
コグニティブファームディストリビューションをインストール→そこにツールを指定→実行します）。テーブル
以下は完全な読む順序です。
Obsidian でドキュメントを読んでいます。このリポジトリは黒曜石です
ボールト — リポジトリ フォルダーをボールトとして開き、README.md とその下のすべてのファイルを開きます。
docs/ は、作業リンク、バックリンク、グラフ ビューを使用してナビゲートできます。の
commit された .obsidian/ ディレクトリには、ボールト設定 (マークダウン リンク、
相対パス);ユーザーごとの UI 状態は無視されます。それ以降はセットアップは必要ありません
フォルダを開いているところ。
初めての方は、次の順序でお読みください。
リポジトリ ルートも Obsidian 互換のボールトです。ルートフォルダーを開きます
Obsidian では、バックリンク、グラフ ビュー、および
相対リンク。
高速パス — PyPI からカーネルとその 3 つの CLI をインストールし、起動します
1 つのコマンドで管理された組織を構築します。
pip install コグニティブファーム
コグニティブファーム-ディストリビューション スターターファーム --into ./my-firm をインストールします
./my-firm は現在、管理された企業です。独自の Git リポジトリがあり、起動が検証されています。
docs/getting-started.md が完全です
ゼロからランニングまでのパス。
カーネル自体に取り組んでいますか?クローンを作成して編集可能にインストールします。
git clone https://github.com/sparckix/cognitive-firm ~ /cognitive-firm
cd ~ /認知的企業
python3 -m venv venv && ソース venv/bin/activate
pip install -r 要件.txt
pip install -e 。
＃ 公共

検証パス
スモークを公開する
スモークドッカーを作る
# プリンシパル設定を構成する
cp org/preferences/templates/principal.yaml org/preferences/principal.yaml
$EDITOR org/preferences/principal.yaml
# 環境の構成 — テンプレートをコピーし、キーを入力します
cp .env.example .env
$EDITOR .env
# (オプション) API キーの代わりにクロード コードのサブスクリプション認証を使用します
unset ANTHROPIC_API_KEY # したがって、クロードは OAuth を好みます
クロードセットアップトークン
# スモークテスト — プリフライト + ドライランでの 1 つのデーモンティック
python scripts/org_role_preflight.py --role Research_director
python scripts/agent_daemon.py --role Research_director --tick-once --dry-run
python -m コグニティブ_ファーム.オーケストレーション.org_surface
カーネルサービスをスモークにする
アプリの統合を準拠させる
app-service-integration-smoke を作成する
ソースカバレッジのウォークスルーを作成する
学習ループのウォークスルーを作成する
バックアップ、復元、スモークを作成する
make docs-surface-check
予行チェックは候補となる作業を検出し、ディスパッチされる内容を印刷し、予算を消費せずに終了します。プリフライトが失敗した場合は、不足しているマンデート ファイルまたはプリファレンス ファイルとその作成方法が示されます。
.env.example は、カーネルと Orbit のすべての環境変数を文書化します。
ダッシュボードを読みました。ほとんどの変数はオプションです。少なくとも 1 つの LLM API キーが必要です
カーネル自体によって行われるモデル呼び出しの場合。
組織とパッケージの管理
コグニティブファームディストリビューションは、管理された組織をインストールおよび管理する方法です
およびパッケージ — 人によって組み立てられるのではなく、一度のアクションで管理される組織
プロトコルカタログからのハンド:
コグニティブファームディストリビューションリスト
コグニティブファーム-ディストリビューション スターターファーム --into ./my-firm をインストールします
インストールはトランザクションであり、Git-backed です。ターゲットは独自の Git になります。
リポジトリでは、インストールはタグ付けされたコミットとして実行されます。
install/starter-firm/<version> 、インストーラーは新しいバージョンを検証します。
組織する

コミットする前のイオンのガバナンス グラフ。不適切なインストールとは、
可逆 — コグニティブファームディストリビューションのロールバック starter-firm --into ./my-firm 。
コグニティブ-ファーム-ディストリビューション CLI は、 show 、 verify 、 upgrade 、
アンインストールと lint ;パッケージは git URL (SHA 固定、
コンテンツハッシュ化されたpackages.lockに記録され、オーバーレイを
org を実行すると、管理された権限差分提案がファイルされます。パッケージはそうでない場合があります。
役割の権限を拡大します。参照
docs/protocols/distribution.md 。
独自のパッケージ (ディストリビューションまたはオーバーレイ) の構築と公開について説明します。
で

[切り捨てられた]

## Original Extract

A governance kernel for organizations that coordinate persistent human and AI-agent roles. - sparckix/cognitive-firm

Let me know what you think!

GitHub - sparckix/cognitive-firm: A governance kernel for organizations that coordinate persistent human and AI-agent roles. · GitHub
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
sparckix
/
cognitive-firm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
57 Commits 57 Commits .obsidian .obsidian deploy deploy distro distro docs docs orbit orbit org org schemas schemas scripts scripts src src tenants tenants tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FUTURE_AREAS.md FUTURE_AREAS.md LICENSE LICENSE MANIFEST.in MANIFEST.in Makefile Makefile NOTICE NOTICE README.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
A governance kernel for organizations that coordinate persistent human and agent roles.
Identity · Who it's for · Architecture · Repository map · Quickstart · Production · Status · Adoption · License & provenance
cognitive-firm lets one principal coordinate persistent AI agents under
typed mandates, budget caps, human approval and work surfaces, and a tenant
overlay pattern.
It is not an agent framework, an LLM SDK, or a chat client. It is the layer
above those: persistent role offices with mandates, inboxes, sessions,
signals, and audit trails.
The operating model is concrete: agents take actions with provenance, humans do
bounded work with receipts, role offices exchange obligations, and only typed
state transitions change the organization. A partner call, a restricted-system
check, a code change, a forecast miss, and a governance proposal are different
records because they create different follow-up obligations.
The human-agent work model is deliberately not "human approves every action"
and not "agent autonomy with after-the-fact oversight." Agents can run quickly
inside bounded envelopes. Humans contribute where they are the right actor or
judge. The kernel records the state transitions that make the work reviewable.
Start with docs/human-agent-work.md for that
theory and its operational mapping.
It was extracted from a live research deployment and rewritten as a reusable
kernel with tests, examples, and local deployment fixtures.
2. Who it's for & what it solves
Agent frameworks such as LangChain, AutoGen, CrewAI, and Letta are strongest
when they model execution graphs, tool calls, memory, and agent interaction.
This kernel focuses on a different layer: durable roles, authority,
accountability, human work, leases, and audit surfaces around those runtimes.
Single principals running a research lab, fund, or solo company who need persistent agents with bounded authority.
Small teams (1-5 people) wanting one place where role mandates, budgets, audit trails, and approvals all live as files in git.
Researchers building on top of a system that can be inspected, replayed, and forked rather than reverse-engineered from logs.
Out of scope
Why
Enterprise RBAC / SSO
Lean multi-actor role membership is supported. Full enterprise IAM administration, SSO provisioning, and tenant isolation remain deployment work.
A graph/model runtime
The daemon is a governance runtime that dispatches existing CLIs (Claude Code, codex ); it does not implement model inference, graph replay, or node scheduling.
A chat UI
The system of record is the filesystem + git. UIs (Orbit, Telegram) are projections, not the truth.
Agentic prompt engineering
Mandates and roles are typed contracts, not prompt templates.
Cross-org federation
A single repo / single principal is the unit of governance.
Alternatives, briefly
LangGraph / AutoGen / CrewAI : better fit if your unit of work is one graph, crew, or chat run. cognitive-firm is for orgs that persist; those runtimes can project lifecycle events into it.
Letta / MemGPT : better fit if your problem is agent memory. cognitive-firm assumes the filesystem + git is the memory.
n8n / Zapier : better fit if your problem is glue between SaaS APIs. cognitive-firm gates which role may invoke which MCP tool under what mandate.
cognitive-firm is a small set of governance protocols and primitives stacked on
a filesystem-backed system of record. Read docs/PROTOCOLS.md for the full
spec; this section is the map.
Layer
Protocol
What it governs
Spec
Human ↔ org
H2A (Human-to-Agent)
Telegram + Orbit + CLI surfaces; pace-layered attention discipline; STOP authority
docs/protocols/h2a.md
Role → human
A2H (Agent-to-Human work coordination)
Bounded human work requested by role offices, with deliverables, receipts, and agent follow-up
docs/protocols/a2h.md
Role ↔ role
A2A (Agent-to-Agent)
Typed AgentMessage envelopes, obligation lifecycle, content-addressed artifact dependencies, saga compensation
docs/protocols/a2a.md
Role ↔ external
MCP (Model Context Protocol)
Capability-gated outbox-relay dispatch to enterprise systems (Linear, Salesforce, ERPs)
docs/protocols/mcp.md
Cross-cutting
Mandate
Typed authority contracts for what each role may do autonomously vs. by escalation
docs/protocols/mandate.md
Project / tenant
Project Charter
Scope fidelity: core question, out-of-scope boundaries, end states, forecast type, inheritance, and anchor proxies
docs/protocols/project-charter.md
The invariants
The primitives serve a small conceptual core. See
docs/kernel-invariants.md for the full map.
┌─────────────────────────────────────────────────────────────────┐
│ PRINCIPAL (human) │
│ ↕ H2A: Telegram (mobile) · Orbit (desktop) · CLI │
│ ↕ A2H: bounded work requests · receipts · follow-up │
└────────────────────────┬────────────────────────────────────────┘
│ approvals + bounded human work
┌────────────────────────▼────────────────────────────────────────┐
│ ROLE OFFICES (research_director, manager, debate_runner, …) │
│ ↕ A2A: typed envelopes · obligation lifecycle · sagas │
└────────────────────────┬────────────────────────────────────────┘
│ subprocess (env-scrubbed for OAuth)
┌────────────────────────▼────────────────────────────────────────┐
│ AGENT RUNTIMES (Claude Code, codex, future open-source) │
│ Subscription quota for agent-side · API tokens for tools │
└────────────────────────┬────────────────────────────────────────┘
│ MCP capability-gated outbox-relay
┌────────────────────────▼────────────────────────────────────────┐
│ EXTERNAL SYSTEMS (Linear, Salesforce, ERPs, ticketing) │
└─────────────────────────────────────────────────────────────────┘
System of record: org/ (filesystem) + git history.
Audit trail = git log; UI = projection.
4. Repository map
New adopter? Start with docs/getting-started.md —
the shortest path from zero to a running governed firm ( pip install →
cognitive-firm-distro install → point the tools at it → run it). The table
below is the full reading order.
Reading the docs in Obsidian. This repository is an Obsidian
vault — open the repo folder as a vault and README.md plus everything under
docs/ is navigable with working links, backlinks, and graph view. The
committed .obsidian/ directory carries the vault config (markdown links,
relative paths); per-user UI state is gitignored. No setup needed beyond
opening the folder.
Read in this order if you are new:
The repository root is also an Obsidian-compatible vault. Open the root folder
in Obsidian to browse the Markdown docs with backlinks, graph view, and
relative links.
The fast path — install the kernel and its three CLIs from PyPI, then stand up
a governed organization in one command:
pip install cognitive-firm
cognitive-firm-distro install starter-firm --into ./my-firm
./my-firm is now a governed firm — its own git repo, verified to boot.
docs/getting-started.md is the full
zero-to-running path.
Working on the kernel itself? Clone and install it editable:
git clone https://github.com/sparckix/cognitive-firm ~ /cognitive-firm
cd ~ /cognitive-firm
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
pip install -e .
# Public verification path
make smoke-public
make smoke-docker
# Configure principal preferences
cp org/preferences/templates/principal.yaml org/preferences/principal.yaml
$EDITOR org/preferences/principal.yaml
# Configure environment — copy the template and fill in keys
cp .env.example .env
$EDITOR .env
# (optional) Use Claude Code subscription auth instead of API key
unset ANTHROPIC_API_KEY # so claude prefers OAuth
claude setup-token
# Smoke test — preflight + one daemon tick in dry-run
python scripts/org_role_preflight.py --role research_director
python scripts/agent_daemon.py --role research_director --tick-once --dry-run
python -m cognitive_firm.orchestration.org_surface
make kernel-service-smoke
make app-integration-conformance
make app-service-integration-smoke
make source-coverage-walkthrough
make learning-loop-walkthrough
make backup-restore-smoke
make docs-surface-check
The dry-run tick discovers candidate work, prints what it would dispatch, and exits without spending budget. If preflight fails, it tells you which mandate or preference file is missing and how to create it.
.env.example documents every environment variable the kernel and Orbit
dashboard read. Most variables are optional. At least one LLM API key is needed
for model calls made by the kernel itself.
Managing organizations and packages
cognitive-firm-distro is how you install and manage governed organizations
and packages — a governed organization in one action, rather than assembled by
hand from the protocol catalog:
cognitive-firm-distro list
cognitive-firm-distro install starter-firm --into ./my-firm
install is transactional and git-backed: the target becomes its own git
repository, the install lands as a commit tagged
install/starter-firm/<version> , and the installer verifies the new
organization's governance graph before committing it. A bad install is
reversible — cognitive-firm-distro rollback starter-firm --into ./my-firm .
The cognitive-firm-distro CLI also covers show , verify , upgrade ,
uninstall , and lint ; a package can be installed from a git URL (SHA-pinned,
recorded in a content-hashed packages.lock ), and installing an overlay onto a
running org files a governed authority-diff proposal — a package may not
widen a role's authority. See
docs/protocols/distribution.md .
Building and publishing your own package — distro or overlay — is covered
in d

[truncated]
