---
source: "https://bunnyandcloud.com/"
hn_url: "https://news.ycombinator.com/item?id=48914764"
title: "Opensourcing Multiplayer AI in Discord"
article_title: "bunny: build in multiplayer with humans and agents on a shared environment"
author: "mehdim"
captured_at: "2026-07-15T01:23:45Z"
capture_tool: "hn-digest"
hn_id: 48914764
score: 2
comments: 0
posted_at: "2026-07-15T00:31:52Z"
tags:
  - hacker-news
  - translated
---

# Opensourcing Multiplayer AI in Discord

- HN: [48914764](https://news.ycombinator.com/item?id=48914764)
- Source: [bunnyandcloud.com](https://bunnyandcloud.com/)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T00:31:52Z

## Translation

タイトル: Discord でマルチプレイヤー AI をオープンソース化
記事タイトル: バニー: 共有環境上で人間とエージェントとのマルチプレイヤーを構築する
説明: バニー: 自己ホスト型のマルチプレイヤー開発ワークスペースで、人間と AI エージェントが同じマシンを共有し、共有シェル、ライブ プレビュー、チーム チャットが 1 か所で行われます。

記事本文:
ウサギと雲
なぜ
仕組み
コンテキスト
検証
コラボレーション
ガバナンス
GitHub
不和
AI時代の共同開発
チャットから配送まで、人間と AI エージェントを同じコンテキストに保ちます。
VM または Docker コンテナを、共有シェル、ライブ プレビュー、チャット ネイティブ ワークフローを備えた共有開発ステーションに変えます。デフォルトでは自己ホスト型。
チャットで話し合います。共有コンテキストで実行します。
コードレビュー、バニー AI エージェント (別のパス) -->
バニー -> GitHub (y=402 のレビュー行とは別にしてください) -->
バニー -> ブラウザ -->
バニー、その後バニー -> スレッド -->
スレッド
今日発送できますか？
製品版でのログインが不安定です
@bunny ライブ ストリーム プレビューを開始
色を変えるべきです
そしてレイアウト
はい、同意します。
ソーシャルログインを提案する必要があります
@bunny ライブプレビューを停止します
@bunny を修正して PR を開く
レビューは良さそうです
検証する
クリックしてください
拒否する
PR がメインに統合されました
AIエージェント
GitHub
ゲートプッシュ
コードレビュー
スレッド内
- ログイン(メールアドレス、パスワード)
+loginWithOAuth(プロバイダー)
session.create(ユーザー)を返す
この変更を出荷できますか?
OAuth フローは正しいようです
🤖 コードレビューが検証されました
AIエージェントによる
コードレビューが完了しました
ライブプレビュー
サインイン
+
🧠
リプレイ
同じ瞬間、唯一の真実の情報源
ターミナル
カーソル
GitHub
メイン？
v2?
2 つのアクション · 1 つの対立
対
ウサギと一緒に
SSH・VM
バニー
期間。
カーソル
クロード
コーデックス
GitHub
スーパーベース
✓ 同期済み
2 つのアクションと 1 つの状態
管理されたゲートウェイ、共有コンテキスト
ログインを修正する
展開しますか？
統治されていない · 文脈が失われた
それぞれのツール、それぞれのメモリ
対
ウサギと一緒に
クロード
カーソル
コーデックス
GitHub
スーパーベース
自己ホスト型ツール
マネージドサービス
ログインを修正する
展開しますか？
@バニーやります
管理されたゲートウェイ · メモリ内の共有コンテキスト
チャット ↔ バニー ↔ 各ツール
以前の提案
現在のバージョン
チームチャット
T
ターミナル
ブラウザ
コーデックス
Git
個別のコンテキスト
対
共有コンテキスト
チームチャット
T
バニー
ターミナル
ブラウザ
コーデックス
Git
1 つの共有コンテキスト
提案:

2 つの開発者とサイロ
開発者A
メイン？
開発者B
ログインを修正する
チームチャット
T
ログインしますか？
ターミナル
壊れたUI
ブラウザ
v2 ≠
コーデックス
合併する
Git
ログインを修正する
ログインを修正する
ログインを修正する
ログインを修正する
ログインを修正する
ログインを修正する
同じトピック、バージョンが競合している
対
関連するコンテキスト
共有コンテキスト
議論します
議論します
コマンド
@bunny テストを実行する
@bunny プレビューを開く
@bunny コミット修正
✓ ログインを修正する
✓ ログインを修正する
コントロールプレーン
✓ ログインを修正する
✓ ログインを修正する
✓ ログインを修正する
真実の情報源
✓ ログインを修正する
✓ ログインを修正する
✓ ログインを修正する
✓ ログインを修正する
開発者A
チームチャット
T
テストは失敗しますか？
ログインが壊れていますか？
その上で
@bunny テストを実行する
開発者B
バニー
ターミナル
ブラウザ
コーデックス
Git
チャットで議論する · コントロールツール · 共有コンテキスト
シフト
ローカル リポジトリから共有の実行可能コンテキストまで
各開発者がローカルで作業し、GitHub 経由で同期するのではなく、チームが共同作業します。
チャット、ロール、権限をそのままにして、共通のリモート環境に直接アクセスできます。
VPS またはコンテナはチームの永続的なワークスペースになります。ターミナル、ライブ プレビュー、
ストリーミングされたブラウザー、コード、フィードバック、実験が集まる統一されたタイムライン。
同じコンテキスト、同じコーディング エージェント、および異なるチャネルから作業する
エンジニア、デザイナー、オペレーター、技術者以外のメンバーなど、すべての貢献者に権利が与えられます。
プロンプトを通じて、すでに使用しているツール (シェル、CLI、スクリプト、Codex、Claude) をインストールします。
チームが管理する環境でワークフローを共有します。
ソフトウェアが実際にどのように作られるかを把握するバージョン管理
Git コミットは、多くの場合、プロジェクトの進化を説明する唯一の永続的な成果物です。しかし、
最新のワークフローは、チャット スレッド、エージェントのプロンプト、集団的な意思決定など、別の場所に存在することを意味します。
テストが実行され、エラーが発生し、トレードオフが受け入れられます。
バニーは、コードのバージョン管理の上にセマンティック レイヤーを追加し、重要な変更をそれぞれ
議論、目標、そして私

そのきっかけとなったやり取り。
単なるコミット ログではなく、1 つのコラボレーション グラフ。
あらゆる変更に対してクエリ可能なコンテキスト
強化されたメタデータは GraphRAG に取り込まれ、Git だけでは答えられない質問を解決できるようになります。
どのような議論がこの実装につながったのでしょうか?
この変更はどのような目的で行われたのでしょうか?
この決定に影響を与えたエージェントまたは貢献者は誰ですか?
長期的には、リモート環境がリポジトリの保護者になります。 Git はそうではありません
消える。それはプロジェクトの唯一のメンタルモデルではなくなります。
バイブコーディングの速度での CI フィードバック
現在のワークフローは AI 支援開発にはあまり適していません。チームはより速く反復し、信頼性を高める
エージェントについてはさらに詳しく説明しますが、CI は到着が遅すぎ、時間がかかりすぎ、場合によっては完全にバイパスされてしまいます。
bunny は、開発と並行して動作する検証エージェントを継続的に統合します。
変更をテストし、回帰を検出し、エラーを分析し、チャットで直接修正を提案します。
上書きや競合のない真のコラボレーション
共有 VM 上では、シェル内のエンジニア、複数の AI など、多くの貢献者が同時に作業します。
エージェントも並行して実行します。 1 回のチェックアウトでは、1 つのセッションで継続的に衝突が発生することになります。
編集中に他人のファイルを踏みつぶすこと。
bunny は、すべてのエージェントとすべてのユーザー シェルに git ワークツリーをプロビジョニングします。
各セッションは独自の作業ディレクトリを取得し、同じリポジトリ上でローカルにブランチします。
分離された編集、安全な並列処理、準備ができたらマージを行うためのマシン。
✓ エージェントおよびシェルセッションごとに 1 つのワークツリー
✓ 重複クローンのない共有オブジェクト ストア
✓ 通常の Git フローを通じて統合された並列ブランチ
.git / オブジェクト
git ワークツリー
エージェントA
偉業/認証
エージェントB
修正/ログイン
シェル
アリス/ウイ
シェル
ボブ/API
並列編集・ゼロ上書き
レポは1つ。たくさんの木々。全員が並行して出荷します。
全員に対するあらゆるアクションの承認
うさぎが間に座っています

チームと接続されているすべてのツールで検証と適用を行う
層。チームメイトまたは AI エージェントがコマンドを実行する前、PR を開く前、または統合に触れる前に、
bunny は、ポリシーと RBAC ルールに照らして承認をチェックします。
リクエストがチャット内の人間からのものであっても、自律的なものであっても、同じガバナンスが適用されます。
バックドアやシャドウ アクセスを持たないエージェント。ポリシーはコンテキストとともに GitHub、シェル、
ブラウザ、および MCP に接続されたすべてのサービス。
✓ 人間とエージェントの役割ベースのアクセス
✓ 接続されているすべてのツールに対するポリシーの適用
✓ チャットコンテキストでの一元的な監査証跡
チームメイト
AIエージェント
ゲスト
ガバナンス
RBAC + ポリシー
GitHub
シェル
MCP
認可が付与されました
ポリシーが拒否されました、アクセスできません
門は一つ。あらゆるツール。人間もエージェントも同様です。
リポジトリだけでなくコンテキストに基づいてコラボレーションする
人間、AI エージェント、環境、ディスカッション、
テストと意思決定は 1 つのフローでつながったままになります。
オープンソース · セルフホスト · 分散チーム向けに構築

## Original Extract

bunny: a self-hosted, multiplayer dev workspace where humans and AI agents share the same machine, with shared shells, live previews, and team chat in one place.

bunny and cloud
Why
How it works
Context
Validation
Collaboration
Governance
GitHub
Discord
Collaborative development for the AI era
From chat to shipping, keep humans and AI agents in the same context.
Turn a VM or Docker container into a shared dev station with shared shells, live previews, and chat-native workflows. Self-hosted by default.
Discuss in chat. Execute in shared context.
code review, bunny AI agent (separate paths) -->
bunny -> GitHub (keep separate from review line at y=402) -->
bunny -> browser -->
bunny, then bunny -> thread -->
Thread
Can we ship today?
Login is flaky in prod
@bunny launch live stream preview
We should change the color
and the layout
Yes, I agree, and we
should propose social login
@bunny stop live preview
@bunny fix and open PR
Review looks good
Validate
Click me
Reject
PR merged to main
AI agent
GitHub
gated push
Code review
in Thread
- login(email, password)
+ loginWithOAuth(provider)
return session.create(user)
Can we ship this change?
OAuth flow looks correct
🤖 Code review validated
by AI agent
Code review done
live preview
Sign in
+
🧠
Replay
Same moment, one source of truth
Terminal
Cursor
GitHub
main?
v2?
two actions · one conflict
vs
With bunny
ssh · VM
bunny
Term.
Cursor
Claude
Codex
GitHub
Supabase
✓ synced
two actions · one state
Governed gateway, shared context
fix login
deploy?
ungoverned · context lost
each tool, its own memory
vs
With bunny
Claude
Cursor
Codex
GitHub
Supabase
self-hosted tools
managed services
fix login
deploy?
@bunny do it
governed gateway · shared context in memory
chat ↔ bunny ↔ every tool
Earlier proposals
Current version
team chat
T
Terminal
Browser
Codex
Git
separate contexts
vs
Shared context
team chat
T
bunny
Terminal
Browser
Codex
Git
one shared context
Proposal: 2 devs and silos
Dev A
main?
Dev B
fix login
team chat
T
login?
Terminal
broken UI
Browser
v2 ≠
Codex
to merge
Git
fix login
fix login
fix login
fix login
fix login
fix login
same topic · conflicting versions
vs
Relevant context
shared context
discusses
discusses
commands
@bunny run tests
@bunny open preview
@bunny commit fix
✓ fix login
✓ fix login
control plane
✓ fix login
✓ fix login
✓ fix login
source of truth
✓ fix login
✓ fix login
✓ fix login
✓ fix login
Dev A
team chat
T
tests failing?
login broken?
on it
@bunny run tests
Dev B
bunny
Terminal
Browser
Codex
Git
discuss in chat · control tools · shared context
The shift
From local repos to shared, executable context
Instead of each developer working locally and synchronizing through GitHub, teams collaborate
directly around a common remote environment, with chat, roles, and permissions intact.
A VPS or container becomes the team's persistent workspace: terminals, live previews,
streamed browsers, and a unified timeline where code, feedback, and experiments converge.
Work from a channel with the same context, the same coding agent, and distinct
rights for every contributor, whether engineers, designers, operators, or non-technical members.
Install the tools you already use (shell, CLIs, scripts, Codex, Claude) through prompts
and shared workflows in an environment your team controls.
Versioning that captures how software is actually made
The Git commit is often the only durable artifact explaining a project's evolution. But in
modern workflows, meaning lives elsewhere: chat threads, agent prompts, collective decisions,
tests run, errors encountered, and trade-offs accepted.
bunny adds a semantic layer on top of code versioning, connecting each important change to
the discussions, goals, and interactions that led to it.
One collaboration graph, not just a commit log.
Queryable context for every change
Enriched metadata is ingested into a GraphRAG, enabling questions that Git alone cannot answer.
Which discussion led to this implementation?
What objective did this change serve?
Which agents or contributors influenced this decision?
In the long term, the remote environment becomes the guardian of the repository. Git doesn't
disappear. It stops being the project's only mental model.
CI feedback at the speed of vibe coding
Current workflows are poorly suited to AI-assisted development. Teams iterate faster, rely
more on agents, yet CI arrives too late, takes too long, and sometimes gets bypassed entirely.
bunny integrates a validation agent that works in parallel with development: continuously
testing changes, detecting regressions, analyzing errors, and suggesting fixes directly in chat.
True collaboration without overwrites or conflicts
On a shared VM, many contributors work at once: engineers in their shells, multiple AI
agents in parallel. A single checkout would mean constant collisions, with one session
stomping another's files mid-edit.
bunny provisions a git worktree for every agent and every user shell.
Each session gets its own working directory and branch on the same repository, locally on
the machine for isolated edits, safe parallelism, and merge when ready.
✓ One worktree per agent and per shell session
✓ Shared object store with no duplicate clones
✓ Parallel branches, integrated through your usual Git flow
.git / objects
git worktree
Agent A
feat/oauth
Agent B
fix/login
Shell
alice/ui
Shell
bob/api
Parallel edits · zero overwrite
One repo. Many trees. Everyone ships in parallel.
Authorization on every action, for everyone
bunny sits between your team and every connected tool with a verification and enforcement
layer. Before any teammate or AI agent runs a command, opens a PR, or touches an integration,
bunny checks authorizations against your policies and RBAC rules.
The same governance applies whether the request comes from a human in chat or an autonomous
agent, with no backdoors or shadow access. Policies travel with context across GitHub, shells,
browsers, and every MCP-connected service.
✓ Role-based access for humans and agents
✓ Policy enforcement on every connected tool
✓ Centralized audit trail in chat context
Teammate
AI agent
Guest
Governance
RBAC + policies
GitHub
Shell
MCP
Authorization granted
Policy denied, no access
One gate. Every tool. Humans and agents alike.
Collaborate around context, not just a repository
Software development as a living process where humans, AI agents, environments, discussions,
tests, and decisions stay connected in one flow.
Open source · Self-hosted · Built for distributed teams
