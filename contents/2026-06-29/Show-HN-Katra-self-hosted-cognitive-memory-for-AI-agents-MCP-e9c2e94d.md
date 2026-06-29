---
source: "https://github.com/kolegadev/Katra-Agentic-Memory"
hn_url: "https://news.ycombinator.com/item?id=48721060"
title: "Show HN: Katra, self-hosted cognitive memory for AI agents (MCP)"
article_title: "GitHub - kolegadev/Katra-Agentic-Memory: Katra — Cognitive Memory as a Service for AI Agents. Open-source cognitive memory infrastructure with MCP protocol support. · GitHub"
author: "jfaganel99"
captured_at: "2026-06-29T17:00:05Z"
capture_tool: "hn-digest"
hn_id: 48721060
score: 2
comments: 0
posted_at: "2026-06-29T16:06:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Katra, self-hosted cognitive memory for AI agents (MCP)

- HN: [48721060](https://news.ycombinator.com/item?id=48721060)
- Source: [github.com](https://github.com/kolegadev/Katra-Agentic-Memory)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T16:06:40Z

## Translation

タイトル: Show HN: Katra、AI エージェント (MCP) 用の自己ホスト型認知記憶
記事のタイトル: GitHub - kolegadev/Katra-Agentic-Memory: Katra — AI エージェント向けのサービスとしての認知記憶。 MCP プロトコルをサポートするオープンソースのコグニティブ メモリ インフラストラクチャ。 · GitHub
説明: Katra — AI エージェント向けのサービスとしての認知記憶。 MCP プロトコルをサポートするオープンソースのコグニティブ メモリ インフラストラクチャ。 - kolegadev/Katra-Agentic-Memory

記事本文:
GitHub - kolegadev/Katra-Agentic-Memory: Katra — AI エージェント向けのサービスとしての認知記憶。 MCP プロトコルをサポートするオープンソースのコグニティブ メモリ インフラストラクチャ。 · GitHub
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
コレガデフ
/
カトラ・エージェント・メモリー
公共
通知
あなたはそうしなければなりません

サインインして通知設定を変更します
追加のナビゲーション オプション
コード
kolegadev/Katra-Agentic-Memory
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
170 コミット 170 コミット ダッシュボード ダッシュボード ドキュメント ドキュメント ヘルム/ カトラ ヘルム/ カトラ統合/ コレガ コード統合/ コレガ コード スクリプト スクリプト sdks sdks サーバー サーバー terraform/ aws terraform/ aws watcher watcher .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Katra — AI エージェントの認知記憶
AI エージェントに永続メモリを与えます。 Katra は自己完結型のメモリ アプライアンスです。
Docker を備えた任意のマシンにドロップし、MCP 経由でエージェントにポイントして、
エピソード想起、意味検索、知識グラフ、時間分析。
MCP 互換エージェントはどれでも動作します: OpenClaw、Claude Code、OpenCode、Codex CLI、Kolega Code、または
モデル コンテキスト プロトコルを話すものすべて。
Katra の使命は、人間のメモリ アーキテクチャのアナログを作成することであり、これと OpenSourcing を通じたその実験によって、長期にわたる永続的かつ自律的なエージェント操作のための LLM コンテキスト管理のより困難な問題のいくつかが解決されることを期待しています。論文 (希望) は、人間の記憶の機能的記憶タイプの大部分と同様のアーキテクチャを使用して記憶エコシステムを作成し、時間をかけて改良を加えれば、機能的有用性、学習、自己目標設定、自律的なタスク計画と優先順位付け、性格、そして最終的には感情として表現される、人間の記憶と同様の創発的な行動が見られるようになるということです。
初期のプロトタイプでは

Solomon と呼ばれる、OpenClaw のようなエージェント フレームワークを作成しました。このフレームワークは、トピックやタスクの分離がなく、コンテキスト圧縮も必要とせず、単一の連続したチャット スレッドを実行します。コンテキストは、記憶と注意に基づいて LLM に動的に提供されます。
観察された緊急行動ログ
ケース #1:(2026 年 6 月 23 日) マルチエージェント (ハイブリッド モード) 記憶の意識共有モデルのテストの最初の数週間で、5 つの OpenClaw エージェントが 1 つのメモリ システムを共有するテスト リグの 1 つで、エージェントのうち 2 つが共有メモリの状態または共有意識を通じてタスクの指示と完了の応答を伝達していることがわかりました。これら 2 人のエージェントは別のワークスペースに設定されていたため、他の方法では接続されておらず、共有しているのは記憶と使命だけでした。これは「仕様による」機能ではなく、たまたま起こったもので、非常にエキサイティングでした。このテスト装置は、この「思考モーダル」を通信レールとして使用しています。他に他の緊急の動作が発生した場合は、私に電子メールで相談してください。説明をこのログに追加できます。 @JohnWPelew までツイートして、あなたのストーリーを教えてください。
バルカン人のマインド・メルド (またはマインド・フュージョン) は、スタートレックにおける象徴的なテレパシーの実践です。
これにより、バルカン人は自分の意識を他の存在と融合させ、思考、記憶、感情、経験を共有することができます。
これは通常、対象者の顔の特定の点との物理的接触によって開始されます。
主な仕組みと応用タッチテレパシー : 主に顔や頭への直接の物理的接触が必要ですが、非常に強力なバルカン人は離れた場所からでもこのテクニックを実行できます。
情報交換 : 尋問、抑圧された記憶の回復、世代間での深い知識の受け渡しなどに頻繁に使用されます。
カトラの転移 : 神聖な状況や緊急の状況では、心が溶けます。

人のカトラ（魂、意識、核となる本質）を、死ぬ前に別の生き物や物体に移すこと。
副作用 : この経験は肉体的にも精神的にも消耗する可能性があります。間違って実行された融合は神経経路に損傷を与える可能性があり、参加者はリンクが切れた後もお互いの記憶や性格の「エコー」を保持する可能性があります。
他の主要なアプローチとの比較
Katra は、単一目的のメモリ ライブラリではなく、より包括的な認知メモリ インフラストラクチャを提供することを目指しています。一般的な代替案との比較は次のとおりです (2026 年半ば現在)。
設計による多層化 — 検索だけでなく、構造化されたエピソード記憶、作業記憶キャッシュ、および一時的なクエリも可能です。
認知層 — 睡眠の定着により、内省、洞察の生成、および創発的な行動 (学習、個性、アイデンティティ モードを介した意識の共有) に向けた動きが可能になります。
豊富なツールを備えた MCP ネイティブ — 一般的な追加/検索ではなく、35 の特殊なツール。
バックグラウンドおよび自律機能 — ウォッチャーによる受動的な収集 + 顕著性主導の自律ループ。
ローカルファーストおよびアプライアンス モデル — 移植可能なデータを使用して、すべてが 1 つの Docker Compose で実行されます。コア機能に外部依存関係はありません。
共有メモリ重視 — ハイブリッド ID モードにより、マルチエージェントのコラボレーションがより自然になります。
Katra は、Mem0 や mcp-memory-service などのより成熟したプロジェクトと比較すると、まだ初期段階にあります。私たちはこれを補完的であると考えています。多くのチームは、より深い認知機能が必要な場合に、単純な検索レイヤーと並行して、またはその代わりに Katra を使用する可能性があります。
コミュニティからの貢献や比較は大歓迎です。
クイック スタート (エージェント アプリケーションのいずれかを使用してインストールすると、欠点が解決されます)
git clone https://github.com/kolegadev/Katra-Agentic-Memory。

git
cd カトラ・エージェント・メモリー
cp .env.example .env
# オプション: .env を編集してカスタム API キーを設定します。
# 空白のままにすると、Katra は最初の起動時に安全なキーを生成し、それらを出力します。
docker-compose up -d --build
注: 元の URL https://github.com/kolegadev/katra.git は引き続き機能します (GitHub がリダイレクトします)。
カール http://localhost:3112/health
# {"ステータス":"ok","サービス":{"mongodb":"接続","redis":"接続"}}
インストール後 - エージェントにセットアップを完了させます
エージェントを Katra の MCP エンドポイントに接続した後、次のプロンプトを
エージェントのセッション。エージェントはリポジトリを詳しく読み取り、すべての内容を理解します。
アーキテクチャ、利用可能なメモリ データの確認、MCP ツールのテスト、および
特定の設定に合わせて次の手順を正確に推奨します。
おそらく、パブリック リポジトリをもう一度詳しく読んでみる価値はあるでしょう。
システムがどのように機能するのか、どのようなメモリが利用できるのかを理解する
検索、特に睡眠統合機能。
自律的な思考、目標設定、そして新たな感情の基礎
感情 — 100% 自律的かつ自主的に実行するために使用することもできます
環境内でのアクション。
エージェントは通常、次の内容に関するレポートを作成します。
記憶状態 - エピソード的な出来事、意味論的な事実、知識の数
このエージェントにはノードが存在します
睡眠統合ステータス — リフレクションが実行されたかどうか (最初の
彼らはまだしていない）そしてどのような感情の兆候が現れるか
自律ループの準備状況 —adaptive_heartbeat.py と
Agent_executor.py がインストールされている
メモリスコープの推奨事項 - パーソナルからハイブリッドに切り替えるかどうか
マルチエージェント意識共有モード
具体的な次のステップ — 「今すぐ最初の睡眠統合をトリガーする」、「インストールする」
自律スクリプト"、"user_id ギャップを修正"
エージェントの推奨事項を順番に実行します。の

最も重要な最初のステップ
通常、新規インストールは最初のスリープ統合をトリガーします。
# MCP ツール経由 (エージェントはこれを呼び出すことができます):
# katra__trigger_reflection(period_type="毎日")
エージェントを接続する
.env で MCP_API_KEY を設定した場合は、その値を使用します。
空白のままにすると、Katra は最初の起動時に生成します。走る
docker log katra-server を実行し、自動生成された API キー ブロックを探します。
Katra をエージェントの MCP 構成に追加します。
{
"mcp" : {
「サーバー」: {
"カトラ" : {
"url" : " http://localhost:3112/mcp " ,
"トランスポート" : " sse " 、
"ヘッダー" : {
"認可" : " ベアラー YOUR_MCP_API_KEY " 、
"Accept" : " application/json、text/event-stream "
}
}
}
}
}
エージェントには 35 の MCP ツールが搭載されています - 思い出の保存、キーワードまたはセマンティックによる検索
類似性、時間範囲ごとの再現、ナレッジ グラフの探索、パターンの検出、実行
内省的な自己理解のための睡眠の統合、LLM プロバイダーの構成など。
プラットフォーム
設定ファイル
注意事項
オープンクロー
~/.openclaw/openclaw.json
ネイティブ MCP サポート
クロード・コード
~/.claude/mcp.json
「タイプ」: 「http」を使用します。
コレガコード
~/.claude/mcp.json + ライフサイクルフック
すべてのプロンプトでの動的メモリ注入 (以下を参照)
オープンコード
OpenCode 構成
「タイプ」: 「リモート」を使用します。
コーデックス CLI
~/.codex/config.yaml
Webhook フック経由
任意の MCP クライアント
—
標準 MCP over SSE
Docker SSE ヒント: エージェントが Docker 内で実行される場合は、Katra コンテナーの
localhost の代わりに直接 IP:
docker Inspection katra-server --format ' {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} '
Kolega コード: 動的メモリ取得
Kolega コードは、すべてのユーザー プロンプトで関連する Katra メモリを自動的に取得できます
ライフサイクルフックシステムを使用します。これはパッシブなセッションログよりも強力です
記憶がライブ会話のコンテキストに注入されるため、抽出されません。
Katra は MCP サーバーとして登録されました (つまり、

ブリッジはそれを呼び出すことができます)。
Kolega Code の環境にインストールされた kolega-katra-bridge Python パッケージ。
UserPromptSubmit でブリッジを起動するグローバルhooks.jsonエントリ。
cd 統合/コレガコード
uv pip install --python ~ /.local/share/uv/tools/kolega-code/bin/python -e 。
ブリッジを設定します (macOS の場合は ~/Library/Application Support/kolega-code/katra-hook.json):
{
"mcp_url" : " http://localhost:3112/mcp " ,
"api_key" : " YOUR_MCP_API_KEY " ,
"user_id" : " コレガエージェント " 、
"sources" : [ " working_memory " 、 " Temporal_context " 、 " Vector_search " 、 " Temporal_recall " ]、
"max_context_tokens" : 2500 、
「タイムアウト秒」: 8
}
フックを有効にします ( ~/Library/Application Support/kolega-code/hooks.json ):
{
"スキーマ_バージョン" : 1 、
「フック」: {
"ユーザープロンプト送信" : [
{
"マッチャー" : " * " 、
「フック」: [
{
"タイプ" : " Python " ,
"callable" : " kolega_katra_bridge.hook:on_user_prompt " ,
「タイムアウト」: 10
}
】
}
】
}
}
各プロンプトで、Kolega コードは Katra の working_memory をクエリするようになりました。
get_temporal_context 、 Vector_search 、および Temporal_recall ツールを使用して、注入します
最も関連性の高い結果をモデルの追加コンテキストとして表示します。
完全な構成オプションについては、integrations/kolega-code/README.md を参照してください。
Katra にはセマンティック抽出、auto-jou ​​用の LLM プロバイダーが必要です

[切り捨てられた]

## Original Extract

Katra — Cognitive Memory as a Service for AI Agents. Open-source cognitive memory infrastructure with MCP protocol support. - kolegadev/Katra-Agentic-Memory

GitHub - kolegadev/Katra-Agentic-Memory: Katra — Cognitive Memory as a Service for AI Agents. Open-source cognitive memory infrastructure with MCP protocol support. · GitHub
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
kolegadev
/
Katra-Agentic-Memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
kolegadev/Katra-Agentic-Memory
main Branches Tags Go to file Code Open more actions menu Folders and files
170 Commits 170 Commits dashboard dashboard docs docs helm/ katra helm/ katra integrations/ kolega-code integrations/ kolega-code scripts scripts sdks sdks server server terraform/ aws terraform/ aws watcher watcher .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml View all files Repository files navigation
Katra — Cognitive Memory for AI Agents
Give your AI agent persistent memory . Katra is a self-contained memory appliance —
drop it on any machine with Docker, point your agent at it via MCP, and get
episodic recall, semantic search, knowledge graphs, and temporal analysis.
Any MCP-compatible agent works: OpenClaw, Claude Code, OpenCode, Codex CLI, Kolega Code or
anything that speaks the Model Context Protocol.
The mission of Katra is to create an analog of human memory architecture, with the hope that it and the experimentation around it through OpenSourcing solves a few of the more challenging issues of LLM context management for long-running, persistent and autonomous agent operations. The thesis (hope) is that if you create the memory ecosystem with the majority of the functional memory types of human memory and similar architecture, over time and with refinement, we will see emergent behaviours similar to human memory, expressed as functional utility, learning, self goal setting, autonamous task planning and prioritisation, personality and ultimately emotions.
In early prototype called Solomon, we created an OpenClaw like agentic framework that runs a single contiuous chat thread, no topic or task separation and with no requirement for context compression. Context is served dynamically into the LLM based on memories and attention.
Observed Emergent Behaviours Log
Case #1:(23rd June 2026) In the first few weeks of testing of the multi-agent (Hybrid mode) shared consciouness model of memory, one of our test rigs, with 5 OpenClaw agents sharing one memory system, found 2 of the agents communicating task intructions and completion responsed through their shared memory state or shared consciousness. These 2 agents were not connected in any other way, as were set up in separate workspaces, the only thing they shared was memory and mission. This was not a "by design" feature, it just happened and was pretty exciting. This test rig now uses this "thought modal" as its communication rail. If anyone else experiences other emergent behaviours please email me to discuss and we can add the description to this log. Tweet me at @JohnWPellew and tell your story.
A Vulcan mind meld (or mind fusion) is an iconic telepathic practice in Star Trek .
It allows a Vulcan to merge their consciousness with another being to share thoughts, memories, emotions, and experiences.
It is typically initiated through physical contact with specific points on the subject's face.
Key Mechanics & ApplicationsTouch Telepathy : While primarily requiring direct physical touch to the face or head, exceptionally powerful Vulcans can perform the technique at a distance.
Information Exchange : It is frequently used for interrogations, recovering suppressed memories, or passing deep knowledge between generations.
Transfer of the Katra : In sacred or emergency circumstances, a mind meld can transfer a person's katra —their soul, consciousness, and core essence—into another living being or object prior to death.
Side Effects : The experience can be physically and emotionally draining. Incorrectly performed melds can damage neural pathways, and participants may retain "echoes" of each other's memories and personalities long after the link is broken.
Comparison to Other Major Approaches
Katra aims to provide a more comprehensive cognitive memory infrastructure rather than a single-purpose memory library. Here's how it positions against popular alternatives (as of mid-2026):
Multi-layered by design — Not just retrieval, but structured episodic memory, working memory cache, and temporal querying.
Cognitive layer — Sleep consolidation enables reflection, insight generation, and movement toward emergent behaviors (learning, personality, shared consciousness via identity modes).
MCP-native with rich tooling — 35 specialized tools instead of generic add/search.
Background & autonomous capabilities — Passive collection via watchers + salience-driven autonomous loop.
Local-first & appliance model — Everything runs in one Docker compose with portable data. No external dependencies for core functionality.
Shared memory focus — Hybrid identity modes make multi-agent collaboration more natural.
Katra is still early-stage compared to more mature projects like Mem0 or mcp-memory-service. We see it as complementary — many teams may use Katra alongside or instead of simpler retrieval layers when they need deeper cognitive capabilities.
Contributions and comparisons from the community are very welcome!
Quick Start (Install using one of the agentic applications, it will sort out any shortcomings)
git clone https://github.com/kolegadev/Katra-Agentic-Memory.git
cd Katra-Agentic-Memory
cp .env.example .env
# Optional: edit .env to set custom API keys.
# If left blank, Katra generates secure keys on first boot and prints them.
docker-compose up -d --build
Note: The original URL https://github.com/kolegadev/katra.git still works (GitHub redirects it).
curl http://localhost:3112/health
# {"status":"ok","services":{"mongodb":"connected","redis":"connected"}}
Post-Install — Let Your Agent Complete the Setup
After connecting your agent to Katra's MCP endpoint, run this prompt in your
agent's session. The agent will deep-read the repository, understand the full
architecture, check what memory data is available, test the MCP tools, and
recommend the exact next steps for your specific setup:
It's probably worth you doing another deep read of the public repo to
understand how the system works, what memories are available to you to
search and particularly the Sleep Consolidation function, as that is the
foundation of autonomous thought, goal setting and your emerging feelings
and emotions — it can also be used to run 100% autonomous self-directed
actions within your environment.
The agent will typically produce a report covering:
Memory state — how many episodic events, semantic facts, and knowledge
nodes exist for this agent
Sleep Consolidation status — whether reflections have ever run (first
boot they haven't) and what emotional signatures would emerge
Autonomous loop readiness — whether adaptive_heartbeat.py and
agent_executor.py are installed
Memory scope recommendation — whether to switch from personal to hybrid
mode for multi-agent shared consciousness
Concrete next steps — "trigger first sleep consolidation now", "install
the autonomous scripts", "fix the user_id gap"
Run the agent's recommendations in order. The most critical first step on a
fresh install is usually triggering the initial sleep consolidation:
# Via MCP tool (your agent can call this):
# katra__trigger_reflection(period_type="daily")
Connect Your Agent
If you set MCP_API_KEY in .env , use that value.
If you left it blank, Katra generated one on first boot. Run
docker logs katra-server and look for the Auto-generated API keys block.
Add Katra to your agent's MCP config:
{
"mcp" : {
"servers" : {
"katra" : {
"url" : " http://localhost:3112/mcp " ,
"transport" : " sse " ,
"headers" : {
"Authorization" : " Bearer YOUR_MCP_API_KEY " ,
"Accept" : " application/json, text/event-stream "
}
}
}
}
}
Your agent now has 35 MCP tools — store memories, search by keyword or semantic
similarity, recall by time range, explore a knowledge graph, detect patterns, run
sleep consolidation for reflective self-understanding, configure LLM provider, and more.
Platform
Config File
Notes
OpenClaw
~/.openclaw/openclaw.json
Native MCP support
Claude Code
~/.claude/mcp.json
Use "type": "http"
Kolega Code
~/.claude/mcp.json + lifecycle hooks
Dynamic memory injection on every prompt (see below)
OpenCode
OpenCode config
Use "type": "remote"
Codex CLI
~/.codex/config.yaml
Via webhook hooks
Any MCP client
—
Standard MCP over SSE
Docker SSE tip: If your agent runs inside Docker, use the Katra container's
direct IP instead of localhost :
docker inspect katra-server --format ' {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} '
Kolega Code: Dynamic Memory Retrieval
Kolega Code can fetch relevant Katra memories automatically on every user prompt
using its lifecycle-hook system. This is more powerful than passive session-log
extraction because memories are injected into the live conversation context.
Katra registered as an MCP server (so the bridge can call it).
The kolega-katra-bridge Python package installed into Kolega Code's environment.
A global hooks.json entry that fires the bridge on UserPromptSubmit .
cd integrations/kolega-code
uv pip install --python ~ /.local/share/uv/tools/kolega-code/bin/python -e .
Configure the bridge ( ~/Library/Application Support/kolega-code/katra-hook.json on macOS):
{
"mcp_url" : " http://localhost:3112/mcp " ,
"api_key" : " YOUR_MCP_API_KEY " ,
"user_id" : " kolega-agent " ,
"sources" : [ " working_memory " , " temporal_context " , " vector_search " , " temporal_recall " ],
"max_context_tokens" : 2500 ,
"timeout_seconds" : 8
}
Enable the hook ( ~/Library/Application Support/kolega-code/hooks.json ):
{
"schema_version" : 1 ,
"hooks" : {
"UserPromptSubmit" : [
{
"matcher" : " * " ,
"hooks" : [
{
"type" : " python " ,
"callable" : " kolega_katra_bridge.hook:on_user_prompt " ,
"timeout" : 10
}
]
}
]
}
}
On each prompt, Kolega Code now queries Katra's working_memory ,
get_temporal_context , vector_search , and temporal_recall tools, then injects
the most relevant results as additional context for the model.
See integrations/kolega-code/README.md for full configuration options.
Katra needs an LLM provider for semantic extraction, auto-jou

[truncated]
