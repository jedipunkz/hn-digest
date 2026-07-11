---
source: "https://mentedb.com"
hn_url: "https://news.ycombinator.com/item?id=48875892"
title: "Persistent memory for Claude Code that survives context compaction"
article_title: "MenteDB | The Mind Database for AI Agents"
author: "mentedb"
captured_at: "2026-07-11T21:39:56Z"
capture_tool: "hn-digest"
hn_id: 48875892
score: 1
comments: 0
posted_at: "2026-07-11T21:07:20Z"
tags:
  - hacker-news
  - translated
---

# Persistent memory for Claude Code that survives context compaction

- HN: [48875892](https://news.ycombinator.com/item?id=48875892)
- Source: [mentedb.com](https://mentedb.com)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T21:07:20Z

## Translation

タイトル: コンテキスト圧縮後も存続するクロード コードの永続メモリ
記事タイトル: MenteDB | AIエージェントのためのマインドデータベース
説明: シングルパストランスの消費に関する知識を事前にダイジェストした Rust データベース。ラッパーではありません。考えるための革新的なストレージ エンジン。

記事本文:
メンテDB | AI エージェント用の Mind データベース Mente DB ベータ版の機能 例 価格設定 ベンチマーク MCP の比較 ドキュメント デモ GitHub はじめる オープン ソース、Apache 2.0 ベータ版 AI のメモリ、
ゼロから設計された
シングルパストランスの消費に関する知識を事前にダイジェストした Rust データベース。ラッパーではありません。考えるための革新的なストレージ エンジン。
GitHub で見る デモを試す 始める 問題
現在の AI システムは、単一パスでコンテキストを消費しますが、欠落しているものを再検討、反映、認識する機能はありません。
AI は単一のコンテキスト ウィンドウを取得します。再読もフォローアップもありません。すべてのトークンはカウントされる必要があります。
AIは知らないことを感知することができません。不確実性を感じたり、自分自身の知識のギャップを検出したりすることはできません。
フラットなコンテキスト、連想的思考
AI は直線的に読み取りますが、連想的に参加します。生のテキストダンプはノイズによる無駄な計算を行います。
静的命令ファイルは小規模なプロジェクトで機能します。しかし、コンテキストが拡大し、意思決定が蓄積され、さまざまなデバイスで作業するようになると、知識をアクティブに管理するものが必要になります。
何かが変更されるたびに、agent.md を手動で更新します
過去の決定を見つけるには、フラット ファイルを grep する必要があります
矛盾は何かが壊れるまで気づかれない
過去の過ちは忘れられ、繰り返される
コンテキストは 1 台のマシン上にのみ存在します
ファイルはコンテキスト ウィンドウを超えるまで大きくなります
AI は事実、好み、意思決定を自動的に抽出します
セマンティック検索により、セッション全体で関連する記憶が検索されます
矛盾は現れた瞬間にフラグが立てられます
間違いを繰り返す前に痛みの警告が現れる
クラウド同期により、すべてのデバイスが同じ認識を保つことができます
数千の記憶、関連するものだけを検索
AI がすべての会話を読み取り、何を覚えておく価値があるかを判断します。事実、好み、決定、修正は、何もしなくても保存されます。
「私たちは何をしたの

ecide about auth?" は、文字列マッチングではなく、埋め込みによって機能します。MenteDB は、別の単語を使用した場合でも、関連するメモリを見つけます。
先月は Postgres、今月は MySQL と言った場合、MenteDB は不整合にフラグを立てます。静的ファイルでは、決定が変更されたことを検出できません。
過去の間違いは、同じことを繰り返そうとしたときに自動的に表面化します。記録したアンチパターンは、現在のコンテキストと照合されます。
ラップトップとデスクトップは同じ脳を共有します。クラウド モードでは、すべてのデバイスがすべてのメモリを瞬時に確認できることを意味します。
Agent.md は、コンテキスト ウィンドウに対して大きすぎるまで機能します。 MenteDB は何千ものメモリを保存し、ベクトル検索によって関連するメモリのみを取得します。
単なるストレージではないコグニティブ エンジン
MenteDB をデータベースから AI の推論に積極的に参加させる 15 のコア システム。
バックグラウンド パイプラインは、意味論的な事実を抽出し、エンティティをリンクし、コミュニティを検出し、生の会話からユーザー プロファイルを構築します。脳が睡眠中に記憶を定着させるのと同じです。
会話が展開されるにつれてリアルタイムで信念が更新される継続的な記憶の取り込み。
クエリ時ではなく書き込み時に、保存されたメモリから新しい知識を自動的に導き出します。
トピック空間内の会話パスをマッピングして、会話がどこに向かっているかを予測します。
知識のギャップを検出し、AI が何を知らないかを認識できるようにプレースホルダー メモリを作成します。
矛盾する信念を分離することで、矛盾した記憶が文脈を汚染するのを防ぎます。
AI が間違いを繰り返すのを防ぐために、否定的なフィードバックと感情的なトリガーを記録します。
今後のクエリを予測し、ナレッジの分岐予測などのコンテキスト ウィンドウを事前に構築します。
メモリは valid_from/valid_until タイムスタンプを保持します。ポイントインタイム クエリを使用した、削除ではなく一時的な無効化。
LLM は、新しいメモリが無効か更新かどうかを判断します。

または既存のものと互換性があります。 62 のテスト ケースで 100% の精度。
BM25 + HNSW ベクター + RRF の融合により、両方の長所を生かした検索が可能。キーワードの精度は意味の理解を満たします。
クロード コードのフックにより、毎ターン記憶が自動的に行われます。 claude.ai コネクタ、ChatGPT、Cursor、および Copilot は 1 つの URL で接続します。
ネイティブ パフォーマンスを備えた Rust、Python、TypeScript バインディング。 pip install、npm install、または Cargo add を実行してビルドを開始します。
ダッシュボードから API キーを作成、取り消し、監視します。チームおよび運用環境向けの範囲限定のアクセス制御。
思い出が変わったときにリアルタイムで通知を受け取ります。配信ログ、ステータス追跡、自動再試行を備えた HMAC 署名ペイロード。
同じユーザーです。同じ質問です。まったく異なる体験。
メモリなしで実稼働環境にデプロイ
✗ デプロイに失敗しました - 間違った AWS リージョン
先週言いましたよね、ここは us-west-2 です!
申し訳ありませんが、以前のセッションの内容がわかりません。
MenteDB を使用する場合 [MenteDB コンテキストがロードされる] 本番環境へのデプロイ
⚠️ ペインシグナル: 「間違った AWS リージョンが原因でデプロイが失敗しました」 📋 保存: デプロイターゲットは us-west-2、アカウント 4821 us-west-2 (運用アカウント) にデプロイ中...
会話ターンごとに 1 回の通話。 MenteDB は、抽出、保存、取得、矛盾検出、バックグラウンド エンリッチメントを自動的に処理します。
コピー使用 mentedb::MenteDb; mentedb::process_turn::ProcessTurnInput を使用します。 mentedb_context::DeltaTracker を使用します。 let db = MenteDb::open( "./agent-memory" )?; mut デルタ = DeltaTracker::default(); とします。 let result = db.process_turn(&ProcessTurnInput { user_message: "Postgres から SQLite に切り替えました" .into()、assistant_response: Some( "わかりました!" .into())、turn_id: 0、project_context: なし、agent_id: なし、session_id: なし、 }, & mut delta)?; // result.context — プロンプトの準備ができた記憶 // result.facts_extracted —

何が学んだか // result.conflict_count — 競合が検出されました // 睡眠時間の強化がバックグラウンドで自動的に実行されます: // → 会話からセマンティックな事実が抽出されました // → エンティティがリンクされ、重複が排除されました // → コミュニティの概要が生成されました // → ユーザー プロファイルが構築および更新されました アーキテクチャ
ストレージ ページからコグニティブ処理に至るまで、すべてのレイヤーが AI メモリ用に設計された専用スタック。
ツールは32個。サーバーは 1 つ。設定ゼロ。
MenteDB は、6 つのカテゴリにわたる 32 のツールを備えた運用 MCP サーバーを出荷します。 Claude、Cursor、ChatGPT、または任意の MCP クライアントに数秒で接続できます。標準入出力経由でローカルに、またはストリーミング可能な HTTP 経由でリモートに接続できます。
クラウドエンドポイント: https://api.mentedb.com/mcp
LLM プロバイダーが構成されている場合、process_turn はスリープタイムの強化もトリガーします
$ npx mentedb-mcp@最新のセットアップ セットアップ
# 以下も: claude-code (フック)、claude、cursor ログイン (クラウド同期を有効にする)
$ npx mentedb-mcp@latestlogin ローカル モード: 標準入出力経由で任意の MCP 互換クライアントで動作します: デフォルトではクロード デスクトップ、カーソル、クロード コード。明示的なセットアップ ターゲットを介して CLI、カーソル、VS コード、カスタム エージェントをコパイロットします。
クラウド モード: https://api.mentedb.com/mcp でストリーミング可能な HTTP MCP 経由で接続します。ChatGPT、リモート クライアント、および HTTP MCP をサポートするあらゆる環境で動作します。
無料で始めましょう。さらに必要な場合はアップグレードしてください。
AI エージェント用の永続メモリの使用を開始します。
より多くの容量を必要とするパワー ユーザーとチーム向け。
すべての主張は再現可能なテストによって裏付けられています。品質はコミットごとに検証され、パフォーマンスは Criterion で測定されます。
LongMemEval (ICLR 2025) は長期会話記憶の標準ベンチマークです。長いマルチセッション履歴 (それぞれ最大 115,000 トークン) にわたる 6 つの推論タイプにわたる 500 の質問です。公式審査員による採点、修正なし。
置き換えられたメモリはグラフのエッジによって正しく除外されました
90.7%再

20ターンにわたる記憶回復トークンの減少
100 ターン、3 プロジェクト、0% 古いリターン、0.29 ミリ秒の挿入
U カーブ順序は 100% の LLM 準拠を維持します
100% 有用な記憶 vs 80% 単純 (+20pp 改善)
古い信念は 100% 正しく、総当たりスキャンよりも 4.8 倍高速
10,000 の記憶、6/6 の信念の変化を追跡、古い返品は 0
同じ 25 個のメモリ、異なるフォーマット
コマンドは 1 つです。永続的な記憶。
デフォルトでクロード コード フックを設定します。また、claude.ai (コネクタ)、Copilot、Cursor、および任意の MCP クライアントでも動作します。
ターミナルコピー $ npx mentedb-mcp@latest setup # クロードコードはデフォルトでフックします。または副操縦士、クロード、カーソル ✓ MCP 設定が書き込まれている ✓ エージェント命令がインストールされている ✓ 準備完了 — AI に永続メモリが追加されました。 # オプション: ログインしてデバイス間でメモリを同期 $ npx mentedb-mcp@latest login ✓ クラウド同期を有効にする または HTTP 経由で接続します: https://api.mentedb.com/mcp
MenteDB ドキュメント 規約 プライバシー 貢献 Rust で構築。アパッチ2.0。

## Original Extract

A Rust database that pre digests knowledge for single pass transformer consumption. Not a wrapper. A ground up storage engine that thinks.

MenteDB | The Mind Database for AI Agents Mente DB Beta Features Examples Pricing Benchmark Compare MCP Docs Demo GitHub Get Started Open source, Apache 2.0 Beta Your AI's memory,
engineered from scratch
A Rust database that pre digests knowledge for single pass transformer consumption. Not a wrapper. A ground up storage engine that thinks.
View on GitHub Try the Demo Get Started The Problem
Current AI systems consume context in a single pass with no ability to revisit, reflect, or recognize what they are missing.
AI gets a single context window. No re-reading, no follow ups. Every token must count.
AI can't sense what it doesn't know. It can't feel uncertainty or detect its own knowledge gaps.
Flat Context, Associative Thinking
AI reads linearly but attends associatively. Raw text dumps waste compute on noise.
Static instruction files work for small projects. But as context grows, decisions accumulate, and you work across devices, you need something that actively manages knowledge for you.
You manually update agent.md every time something changes
Finding past decisions means grep-ing through a flat file
Contradictions go unnoticed until something breaks
Past mistakes are forgotten and repeated
Context lives on one machine only
File grows until it blows the context window
AI extracts facts, preferences, and decisions automatically
Semantic search finds relevant memories across sessions
Contradictions are flagged the moment they appear
Pain warnings surface before you repeat a mistake
Cloud sync keeps every device on the same page
Thousands of memories, only the relevant ones retrieved
Your AI reads every conversation and decides what is worth remembering. Facts, preferences, decisions, and corrections are stored without you lifting a finger.
"What did we decide about auth?" works via embeddings, not string matching. MenteDB finds relevant memories even when you use different words.
If you said Postgres last month and MySQL this month, MenteDB flags the inconsistency. A static file cannot detect when your decisions change.
Past mistakes surface automatically when you are about to repeat them. Anti-patterns you recorded are matched against your current context.
Your laptop and desktop share the same brain. Cloud mode means every device sees every memory instantly.
An agent.md works until it is too big for the context window. MenteDB stores thousands of memories and retrieves only the relevant ones via vector search.
A cognitive engine, not just storage
Fifteen core systems that transform MenteDB from a database into an active participant in your AI's reasoning.
Background pipeline extracts semantic facts, links entities, detects communities, and builds user profiles from raw conversations. Like the brain consolidating memories during sleep.
Continuous memory ingestion with real time belief updates as conversations unfold.
Automatically derives new knowledge from stored memories at write time, not query time.
Maps conversation paths through topic space to predict where dialogue is heading.
Detects knowledge gaps and creates placeholder memories so the AI knows what it does not know.
Prevents contradictory memories from polluting context by isolating conflicting beliefs.
Records negative feedback and emotional triggers to prevent the AI from repeating mistakes.
Predicts upcoming queries and pre builds context windows, like branch prediction for knowledge.
Memories carry valid_from/valid_until timestamps. Temporal invalidation instead of deletion, with point-in-time queries.
LLM judges whether new memories invalidate, update, or are compatible with existing ones. 100% accuracy on 62 test cases.
BM25 + HNSW vector + RRF fusion for best-of-both-worlds retrieval. Keyword precision meets semantic understanding.
Claude Code hooks make memory automatic on every turn. The claude.ai connector, ChatGPT, Cursor, and Copilot connect with one URL.
Rust, Python, and TypeScript bindings with native performance. pip install, npm install, or cargo add and start building.
Create, revoke, and monitor API keys from the dashboard. Scoped access control for teams and production deployments.
Get real-time notifications when memories change. HMAC-signed payloads with delivery logs, status tracking, and automatic retries.
Same user. Same question. Completely different experience.
Without Memory Deploy to production
✗ Deploy failed — wrong AWS region
I told you LAST WEEK it's us-west-2!
Sorry, I don't have context from previous sessions.
With MenteDB [MenteDB context loaded] Deploy to production
⚠️ Pain signal: "Wrong AWS region caused deploy failure" 📋 Stored: Deploy target is us-west-2, account 4821 Deploying to us-west-2 (your production account)...
One call per conversation turn. MenteDB handles extraction, storage, retrieval, contradiction detection, and background enrichment automatically.
Copy use mentedb::MenteDb; use mentedb::process_turn::ProcessTurnInput; use mentedb_context::DeltaTracker; let db = MenteDb::open( "./agent-memory" )?; let mut delta = DeltaTracker::default(); let result = db.process_turn(&ProcessTurnInput { user_message: "I switched from Postgres to SQLite" .into(), assistant_response: Some( "Got it!" .into()), turn_id: 0, project_context: None, agent_id: None, session_id: None, }, & mut delta)?; // result.context — memories ready for your prompt // result.facts_extracted — what was learned // result.contradiction_count — conflicts detected // Sleeptime enrichment runs automatically in the background: // → semantic facts extracted from conversations // → entities linked and deduplicated // → community summaries generated // → user profile built and updated Architecture
A purpose built stack where every layer is designed for AI memory, from storage pages to cognitive processing.
32 tools. One server. Zero config.
MenteDB ships a production MCP server with 32 tools across 6 categories. Connect Claude, Cursor, ChatGPT, or any MCP client in seconds — locally via stdio or remotely via Streamable HTTP.
Cloud endpoint: https://api.mentedb.com/mcp
process_turn also triggers sleeptime enrichment when an LLM provider is configured
$ npx mentedb-mcp@latest setup setup
# also: claude-code (hooks), claude, cursor Login (enables cloud sync)
$ npx mentedb-mcp@latest login Local mode: Works with any MCP-compatible client via stdio: Claude Desktop, Cursor, Claude Code by default. Copilot CLI, Cursor, VS Code, custom agents via explicit setup targets.
Cloud mode: Connect via Streamable HTTP MCP at https://api.mentedb.com/mcp — works with ChatGPT, remote clients, and any environment that supports HTTP MCP.
Start free. Upgrade when you need more.
Get started with persistent memory for your AI agents.
For power users and teams who need more capacity.
Every claim backed by reproducible tests. Quality validated on every commit, performance measured with Criterion.
LongMemEval (ICLR 2025) is the standard benchmark for long-term conversational memory: 500 questions across six reasoning types over long, multi-session histories (~115K tokens each). Graded by the official judge, unmodified.
Superseded memories correctly excluded via graph edges
90.7% reduction in memory retrieval tokens over 20 turns
100 turns, 3 projects, 0% stale returns, 0.29ms insert
U-curve ordering maintains 100% LLM compliance
100% useful memories vs 80% naive (+20pp improvement)
100% correct on stale beliefs, 4.8x faster than brute-force scan
10,000 memories, 6/6 belief changes tracked, 0 stale returns
Same 25 memories, different formats
One command. Persistent memory.
Sets up Claude Code hooks by default. Also works with claude.ai (connector), Copilot, Cursor, and any MCP client.
Terminal Copy $ npx mentedb-mcp@latest setup # Claude Code hooks by default; or copilot, claude, cursor ✓ MCP config written ✓ Agent instructions installed ✓ Ready — your AI now has persistent memory. # Optional: log in to sync memory across devices $ npx mentedb-mcp@latest login ✓ Cloud sync enabled Or connect via HTTP: https://api.mentedb.com/mcp
MenteDB Docs Terms Privacy Contributing Built with Rust. Apache 2.0.
