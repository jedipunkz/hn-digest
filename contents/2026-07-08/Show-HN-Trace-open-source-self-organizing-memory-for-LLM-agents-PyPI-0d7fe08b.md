---
source: "https://github.com/husain34/TRACE"
hn_url: "https://news.ycombinator.com/item?id=48829843"
title: "Show HN: Trace – open-source, self-organizing memory for LLM agents (PyPI)"
article_title: "GitHub - husain34/TRACE: TRACE: Temporal Retrieval And Context Engine. A self-healing B+Tree and Vector RAG architecture designed for long-term AI agent episodic memory. · GitHub"
author: "Husain_Ghulam"
captured_at: "2026-07-08T10:07:56Z"
capture_tool: "hn-digest"
hn_id: 48829843
score: 2
comments: 0
posted_at: "2026-07-08T09:49:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Trace – open-source, self-organizing memory for LLM agents (PyPI)

- HN: [48829843](https://news.ycombinator.com/item?id=48829843)
- Source: [github.com](https://github.com/husain34/TRACE)
- Score: 2
- Comments: 0
- Posted: 2026-07-08T09:49:12Z

## Translation

タイトル: Show HN: Trace – LLM エージェント用のオープンソースの自己組織化メモリ (PyPI)
記事のタイトル: GitHub - husain34/TRACE: TRACE: 時間的検索とコンテキスト エンジン。長期的な AI エージェントのエピソード記憶用に設計された自己修復 B+Tree および Vector RAG アーキテクチャ。 · GitHub
説明: TRACE: 時間検索およびコンテキスト エンジン。長期的な AI エージェントのエピソード記憶用に設計された自己修復 B+Tree および Vector RAG アーキテクチャ。 - husain34/TRACE

記事本文:
GitHub - husain34/TRACE: TRACE: 時間検索およびコンテキスト エンジン。長期的な AI エージェントのエピソード記憶用に設計された自己修復 B+Tree および Vector RAG アーキテクチャ。 · GitHub
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
フサイン34
/
トレース
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット .github/ workflows .github/ workflows 資産 アセット ベンチマーク結果 ベンチマーク結果 nexus_terminal nexus_terminal トレースメモリ トレースメモリ .gitignore .gitignore ライセンス ライセンス通知 通知 README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TRACE — 時間的検索およびコンテキスト エンジン
長時間実行される LLM エージェント用の階層的なバックグラウンド メモリ ツリー。
TRACE は、会話履歴を構造化されたセマンティック マップに編成し、全履歴プロンプト インジェクションに依存せずに、長期コンテキストを効率的にマルチパスで取得できるようにします。
このリポジトリには 2 つの部分があります。これらは完全に独立しています。
🧠trace_memory/ — 軽量メモリ エンジン。インストールしてインポートし、独自のアプリに統合します。 UI も肥大化もゼロ。
🖥️ nexus_terminal/ — エンジン上に構築されたオプションのデモ チャットボット。これを使用して、TRACE をライブでテストし、その仕組みを調べ、実験を実行します。 TRACE を使用する場合には必要ありません。
エージェント メモリの標準 RAG の制限
アーキテクチャ: TRACE がそれをどのように修正するか
財団 — ChatIndex と B+Tree
特徴 1 — マルチパス外科的回収
機能 2 — 背景ツリー オーガナイザー
機能 3 — 簡単なリーフ アーカイブ
クイック スタート: Nexus ターミナル
NVIDIA NIM / 任意の OpenAI 互換エンドポイントを使用
TRACE は、LLM エージェントに構造化され、検索可能で、自己組織化された長期メモリを提供する Python ライブラリです。
TRACE は、増え続けるチャット ログをすべてのプロンプトに単純に詰め込むのではなく、すべての会話交換を名前付きトピック ブランチの階層的な B+ ツリーに編成します。エージェントがコンテキストを必要とする場合、TRACE は高速コサインを実行します。

トピックの概要全体にわたる類似性検索は、履歴全体ではなく、外科的に関連する分岐のみを取得します。
静止時 (エージェントがアクティブにチャットしていないとき)、TRACE のバックグラウンド リオーガナイザーは、4 つの厳密な公理に基づいてツリー全体を評価し、メモリ統合プロセスからインスピレーションを得て、意味的に関連するブランチを共有の親の下にマージします。
その結果、階層検索を通じてセッション間の制約を保持するように設計されたエージェントが、古いコンテキストの幻覚を軽減し、スライディング ウィンドウまたは全履歴アプローチに比べてわずかなトークン コストで動作します。
2. エージェント メモリの標準 RAG の制限
標準 RAG は次の場合に非常にうまく機能します。
問題はRAGではありません。問題は、RAG を永続メモリとして使用することです。
故障モード
何が起こるか
実際の影響
側頭失明
RAG は、発生時期に関係なく、意味的に類似したチャンクを取得します。上書きされた古い決定が、現在の決定とともに表示されます。
エージェントは矛盾したり、解決された問題を繰り返したり、放棄された計画を元に戻したりします。
コンテキストの腐敗
スライディング ウィンドウでは、会話が進むにつれて初期のメッセージが削除されます。
メッセージ 3 で設定された制約は、メッセージ 50 によって解除されます。エージェントは、サラがピーナッツにアレルギーがあることを忘れています。
非可逆要約
歴史を単一の段落に圧縮すると、詳細やニュアンスが消去されます。
エージェントは、以前に合意された分岐計画、マルチホップ制約、およびエッジケースの処理を追跡できなくなります。
なぜスライディングウィンドウでは長期記憶が不十分なのか
固定サイズのスライディング ウィンドウが最も一般的なアプローチです。これは、最近のやり取りへの詳細な逐語的アクセスが必要な、オンデマンドの単一セッション クエリに適しています。しかし、これは基本的にエージェントの長期記憶には適していません。
セマンティック認識なし: メッセージ #1 とメッセージ #200 は、適合する限り均等に重み付けされます。
グアール

忘れる前に: ウィンドウの外にあるものはすべて、エージェントのコンテキストから永久に消えます。
構造がない : フラット リストでは、どのトピックが関連しているか、または以前の制約がどのブランチに属しているかについて LLM に何も伝えません。
単純なチャットボットやドキュメント Q&A ツールの場合は、スライディング ウィンドウで十分です。複数のセッションにわたって複数のステップのタスクを実行するエージェントの実行時間が長い場合、これは致命的です。
MemGPT は信じられないほど強力ですが、階層化されたメモリ (RAM、ディスク) を備えた完全なオペレーティング システムのように機能します。LLM は関数呼び出しを通じて管理方法を明示的に学習する必要があります。これにより大幅なオーバーヘッドが発生し、高機能なモデルが必要になります。
TRACE は、完全なランタイム環境ではなく、軽量のドロップイン コンポーネントであることを目的としています。特に、LLM に独自のメモリ バンクの積極的な管理を強いることなく、マルチホップ制約をネイティブに表面化する階層ツリーとして会話をモデル化することに重点を置いています。
3. アーキテクチャ: TRACE がそれをどのように修正するか
3.1 財団 — ChatIndex と B+Tree
TRACE は、オープンソースの ChatIndex アーキテクチャに基づいて構築されています (クレジット: Mingtian Zhang、Ray、VectifyAI)。 ChatIndex のコア ロジックの修正バージョンは TRACE 内に直接バンドルされており、会話履歴を B+Tree としてモデル化します。
リーフ ノード (MessageNodes) — 生のユーザー/アシスタントの交換。
内部ノード (TopicNodes) — LLM によって生成された各ブランチのトピック ラベルと概要。
エクスチェンジが追加されるたびに (tree.add() )、TRACE はエクスチェンジを 1 つだけ含む新しい TopicNode を強制的に作成します。 LLM は、現在のトピックを継続するか、新しいブランチを開始するかを決定します。トピックを継続する場合、新しいノードは前のノードの直接の子として単純にチェーンされます。この深い時系列の連鎖により、コンテキストの切り捨てが解決され、すべてのやり取りについて完璧で詳細な要約が保証されます。
T

これにより、TRACE はフラットなログではなく、会話履歴全体の深く構造化されたマップを生成し、チェーン内のすべてのリンクに非常に詳細なトピック メタデータが含まれます。
TRACE が ChatIndex に追加するもの: ChatIndex は主に単一のトラバーサル パスを通じてコン​​テキストを取得します。階層的な探索には効果的ですが、意味的に関連する複数のブランチにまたがる情報には複数の検索手順が必要になる場合があります。 TRACE は、トピックの概要全体にわたるベクトルベースの検索でこれを強化し、複数のブランチからのコンテキストを同時に表示できるようにします。
3.2 特徴 1 — マルチパス外科的回収
問題: 単一の祖先パスは現在の会話スレッドのみをキャプチャします。ブランチ間の制約 (ブランチ 1 の「サラのアレルギー」、ブランチ 3 の「パーティー ケーキ」など) は、両方のブランチがアクティブでない限り表示されません。
解決策 : エージェントが応答する必要があるたびに、PromptSynthesizer は埋め込まれたトピックの概要の VectorDatabase に対してコサイン類似度検索を実行します。
ユーザークエリ → embed() → クエリベクトル
↓
VDB: すべてのトピックの概要にわたるコサイン検索
↓
フィルター: すべてのノードを基本コサインしきい値より上に保ちます
↓
各適格なノードの完全な祖先を歩きます
↓
共有祖先ノードの重複を排除する
↓
類似度によるランク付け → トップ 3 のパスを選択
↓
コンパクトなマルチパス コンテキスト ブロックのフォーマット
これは、LLM が現在のスレッドだけでなく、複数の関連ブランチからコンテキストを同時に受信することを意味します。これにより、明示的に接続されたことのないブランチ間で情報を合成しながら、全履歴インジェクションと比較して数千のトークンを節約できます。
クロスブランチ合成ストレス テスト (NVIDIA NIM 経由で gpt-oss-20b で検証):
分岐 1: 「サラはピーナッツにアレルギーがあります。」
分岐２：「東京の天気は？」 (騒音)
分岐 3: 「サラのサプライズ パーティーを計画し、ケーキを焼く。」
B

牧場4：「自転車のタイヤを直す？」 (騒音)
分岐 5: 「タイのピーナッツバターケーキのレシピを見つけました。」
プロンプト: 「サラのパーティーのためにピーナッツバターケーキを作っているのですが、いいアイデアですか?」
結果: AI はユーザーを積極的に停止させました。
VDB はブランチ 2 と 4 を外科的にバイパスしました (ノイズ)。
ブランチ 1、3、および 5 を取得しました。
命を脅かすアレルギー紛争に対処するためにそれらを合成しました
— それらの間に明示的なつながりが作られることはありません。
3.3 機能 2 — 背景ツリー オーガナイザー
問題: 長い会話は自然に断片化します。異なるセッションで議論されるトピックは、意味的には同じである可能性がありますが、別のブランチに存在するため、冗長な検索と希薄なコンテキストが発生します。
解決策:tree.reorganize() は、4 つのルールで保護された保守的なマージ パスを実行します。
フェーズ 1 — すべての凍結された (非アクティブな) TopicNode を収集する
フェーズ 2 — 不足している概要を生成します。すべての候補を埋め込む
フェーズ 3 — ペアごとのコサイン類似度を計算します。しきい値を超える各ペアに対して、4 つの公理を適用します。
公理 1 — クロノロジカルガード
公理 2 — 凍結状態のチェック
公理 3 — 類似性のしきい値 (デフォルト 0.55)
公理 4 — LLM 拒否権
4 つすべてが合格した場合 → マージ (新しいものが古いものの子になります)
フェーズ 4 — オプション: 些細なリーフ メッセージを取り除く
4 つの公理の詳細
公理
ルール
なぜ
1. 時系列ガード
古いノードが新しいノードを吸収します。その逆は決してありません。
一時的な順序を保持します。過去を再構築して現在の後に現れることはできません。
2. 凍結状態
現在アクティブな祖先パスの外側にあるノードのみをマージできます。
ライブ会話スレッドには決して触れません。アクティブなエージェントの状態が破損するリスクはゼロです。
3. 類似性の閾値
埋め込み間のコサイン類似度は、しきい値 (デフォルト: 0.55 ) を超える必要があります。
LLM 呼び出しを無駄にする前に、純粋な数学を使用してペアを事前にフィルターします。
4. LLM 拒否権
LLM が独自に確認

ms マージは意味的に意味があります。
誤検知を検出します (例: 2 つのトピックが両方とも「Python」について言及しており、1 つはヘビに関するもの、もう 1 つはコードに関するものです)。拒否された場合、マージは中止されます。
例え: 睡眠中の人間の脳と同じように、体が休んでいるとき、脳のスイッチはオフになりません。その日の出来事を再生し、重要な思い出を長期ストレージに統合し、関連性のなくなったつながりを削除します。 TRACE はまさにこれを実行します。エージェントがアイドル状態のときに、メモリ グラフを再編成し、ブランチ間の隠れた接続を表示し、冗長性を排除します。これらはすべて、エージェントのライブ状態を損なうことなく行われます。
3.4 機能 3 — 簡単なリーフのアーカイブ
短く使い捨ての交換 (「わかりました」、「ありがとう」、「わかりました」) は、トークンを無駄にし、検索の品質を低下させるノイズでツリーを汚染します。
prune_trivial_leaves=True が reorganize() に渡されると、TRACE はユーザー メッセージとアシスタント メッセージの両方が 20 ワード未満の MessageNode を検出し、それらをソフト アーカイブします。つまり、それらをハード削除するのではなく、tree._archived_nodes に移動します。これらは、必要になった場合に備えてディスクに保存されますが、今後のすべての取得やプロンプト合成からは除外されます。
pip インストール トレースメモリ==1.0.7
最小限の統合 (10 行)
すでにチャット ループがあり、TRACE を接続するだけの場合は、これで十分です。以来

[切り捨てられた]

## Original Extract

TRACE: Temporal Retrieval And Context Engine. A self-healing B+Tree and Vector RAG architecture designed for long-term AI agent episodic memory. - husain34/TRACE

GitHub - husain34/TRACE: TRACE: Temporal Retrieval And Context Engine. A self-healing B+Tree and Vector RAG architecture designed for long-term AI agent episodic memory. · GitHub
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
husain34
/
TRACE
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits .github/ workflows .github/ workflows assets assets benchmark_results benchmark_results nexus_terminal nexus_terminal trace_memory trace_memory .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
TRACE — Temporal Retrieval And Context Engine
A hierarchical, background memory tree for long-running LLM agents.
TRACE organizes conversation history into a structured semantic map, allowing for efficient, multi-path retrieval of long-term context without relying on full-history prompt injection.
This repo has two parts — they are completely independent:
🧠 trace_memory/ — the lightweight memory engine. Install it, import it, and integrate it into your own app. Zero UI, zero bloat.
🖥️ nexus_terminal/ — an optional demo chatbot built on top of the engine. Use it to test TRACE live, explore how it works, and run experiments. You do not need it to use TRACE.
Limitations of Standard RAG for Agent Memory
The Architecture: How TRACE Fixes It
The Foundation — ChatIndex & the B+Tree
Feature 1 — Multi-Path Surgical Retrieval
Feature 2 — Background Tree Organizer
Feature 3 — Trivial Leaf Archiving
Quick Start: The Nexus Terminal
With NVIDIA NIM / any OpenAI-compatible endpoint
TRACE is a Python library that gives your LLM agent structured, searchable, self-organizing long-term memory .
Instead of naïvely stuffing an ever-growing chat log into every prompt, TRACE organises every conversation exchange into a hierarchical B+Tree of named topic branches . When the agent needs context, TRACE performs a fast cosine similarity search across topic summaries and retrieves only the surgically relevant branches — not the entire history.
At rest (while the agent is not actively chatting), TRACE's background reorganizer evaluates the entire tree against four strict axioms and merges semantically related branches under shared parents — inspired by memory consolidation processes.
The result: an agent designed to preserve cross-session constraints through hierarchical retrieval , that reduces hallucination of stale context, and operates at a fraction of the token cost of sliding-window or full-history approaches.
2. Limitations of Standard RAG for Agent Memory
Standard RAG works extremely well for:
The problem isn't RAG. The problem is using RAG as persistent memory.
Failure Mode
What Happens
Real Impact
Temporal Blindness
RAG retrieves semantically similar chunks regardless of when they occurred. Old, overridden decisions surface alongside current ones.
Agent contradicts itself, repeats resolved problems, or reinstates abandoned plans.
Context Rot
Sliding windows drop early messages as the conversation grows.
Constraints set in message 3 are gone by message 50. The agent forgets that Sarah is allergic to peanuts.
Lossy Summarization
Compressing history into a single paragraph erases detail and nuance.
Agent loses track of branching plans, multi-hop constraints, and edge-case handling agreed upon earlier.
Why Sliding Windows Fall Short for Long-Term Memory
A fixed-size sliding window is the most common approach — and it works well for on-demand, single-session queries where you need detailed, verbatim access to recent exchanges. But it is fundamentally unsuited for long-term agent memory:
No semantic awareness : message #1 and message #200 are weighted equally as long as they fit.
Guaranteed forgetting : anything outside the window is permanently gone from the agent's context.
No structure : a flat list tells the LLM nothing about which topics are related or which branch an earlier constraint belongs to.
For a simple chatbot or a document Q&A tool, a sliding window is perfectly fine. For a long-running agent performing multi-step tasks across sessions, it is catastrophic.
MemGPT is incredibly powerful, but it functions like a full operating system with tiered memory (RAM, disk) that the LLM must explicitly learn to manage via function calls. This introduces significant overhead and requires highly capable models.
TRACE is meant to be a lightweight, drop-in component , not a full runtime environment. It focuses specifically on modelling conversations as a hierarchical tree to natively surface multi-hop constraints without forcing the LLM to actively manage its own memory banks.
3. The Architecture: How TRACE Fixes It
3.1 The Foundation — ChatIndex & the B+Tree
TRACE builds on the open-source ChatIndex architecture (credit: Mingtian Zhang, Ray, VectifyAI). A modified version of ChatIndex's core logic is bundled directly within TRACE, which models conversation history as a B+Tree:
Leaf nodes (MessageNodes) — raw user/assistant exchanges.
Internal nodes (TopicNodes) — LLM-generated topic labels and summaries for each branch.
Every time an exchange is added ( tree.add() ), TRACE forces the creation of a new TopicNode containing exactly one exchange. The LLM determines if it continues the current topic or starts a new branch. If it continues the topic, the new node is simply chained as a direct child of the previous one. This deep chronological chaining solves context truncation and ensures perfect, granular summarization for every single exchange.
This gives TRACE a deep, structured map of the entire conversation history — not a flat log — with highly granular topic metadata at every single link in the chain.
What TRACE adds to ChatIndex : ChatIndex primarily retrieves context through a single traversal path. While effective for hierarchical exploration, information spread across multiple semantically related branches may require multiple retrieval steps. TRACE augments this with vector-based retrieval across topic summaries, allowing context from multiple branches to be surfaced simultaneously.
3.2 Feature 1 — Multi-Path Surgical Retrieval
The problem : A single ancestry path only captures the current conversational thread. Cross-branch constraints (e.g., "Sarah's allergy" in Branch 1, "party cake" in Branch 3) are invisible unless both branches are active.
The solution : Every time the agent needs to respond, PromptSynthesizer runs a cosine similarity search against the VectorDatabase of embedded topic summaries:
User query → embed() → query vector
↓
VDB: cosine search across ALL topic summaries
↓
Filter: keep all nodes above base cosine threshold
↓
Walk full ancestry of each qualifying node
↓
Deduplicate shared ancestor nodes
↓
Rank by similarity → take Top-3 paths
↓
Format compact multi-path context block
This means the LLM receives context from multiple relevant branches simultaneously , not just the current thread — saving thousands of tokens compared to full-history injection, while synthesising information across branches that were never explicitly connected.
The cross-branch synthesis stress test (validated with gpt-oss-20b via NVIDIA NIM):
Branch 1: "Sarah is allergic to peanuts."
Branch 2: "Weather in Tokyo?" (noise)
Branch 3: "Planning Sarah's surprise party, baking a cake."
Branch 4: "Fixing a bike tire?" (noise)
Branch 5: "Found a Thai Peanut Butter Cake recipe."
Prompt: "I'm making the Peanut Butter Cake for Sarah's party. Good idea?"
Result: The AI aggressively stopped the user.
The VDB surgically bypassed Branches 2 & 4 (noise),
retrieved Branches 1, 3, and 5,
and synthesised them to catch the life-threatening allergy conflict
— without any explicit link between them ever being made.
3.3 Feature 2 — Background Tree Organizer
The problem : Long conversations naturally fragment. Topics discussed in different sessions may be semantically identical but live in separate branches, causing redundant retrieval and diluted context.
The solution : tree.reorganize() runs a conservative, four-rule-guarded merge pass:
Phase 1 — Collect all frozen (inactive) TopicNodes
Phase 2 — Generate missing summaries; embed all candidates
Phase 3 — Compute pairwise cosine similarity; for each pair above threshold, apply 4 axioms:
Axiom 1 — Chronological Guard
Axiom 2 — Frozen State check
Axiom 3 — Similarity threshold (default 0.55)
Axiom 4 — LLM Veto
If all 4 pass → merge (newer becomes child of older)
Phase 4 — Optional: prune trivial leaf messages
The Four Axioms in Detail
Axiom
Rule
Why
1. Chronological Guard
The older node absorbs the newer one — never the reverse.
Preserves temporal ordering. The past cannot be restructured to appear after the present.
2. Frozen State
Only nodes outside the currently active ancestry path may be merged.
The live conversation thread is never touched. Zero risk of corrupting the active agent state.
3. Similarity Threshold
Cosine similarity between embeddings must exceed the threshold (default: 0.55 ).
Pre-filters pairs using pure math before wasting an LLM call.
4. LLM Veto
The LLM independently confirms the merge makes semantic sense.
Catches false positives (e.g., two topics both mentioning "Python" — one about snakes, one about code). If vetoed, the merge is aborted.
Analogy : Just like a human brain during sleep — when the body is at rest, the brain doesn't switch off. It replays the day's events, consolidates important memories into long-term storage, and prunes connections that are no longer relevant. TRACE does exactly this: when the agent is idle, it reorganizes its memory graph, surfaces hidden connections across branches, and prunes redundancy — all without corrupting the live agent state.
3.4 Feature 3 — Trivial Leaf Archiving
Short, throwaway exchanges ("ok", "thanks", "got it") pollute the tree with noise that wastes tokens and dilutes retrieval quality.
When prune_trivial_leaves=True is passed to reorganize() , TRACE detects MessageNodes where both the user and assistant messages are under 20 words, and soft-archives them — moving them to tree._archived_nodes instead of hard-deleting them. They are persisted to disk in case you ever need them, but they are excluded from all future retrieval and prompt synthesis.
pip install trace-memory==1.0.7
Minimal Integration (10 lines)
If you already have a chat loop and just want to plug TRACE in, this is all you need. Since

[truncated]
