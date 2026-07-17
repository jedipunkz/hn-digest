---
source: "https://github.com/orbsh/wiki/blob/main/tail-prompt-optimization-en.md"
hn_url: "https://news.ycombinator.com/item?id=48942361"
title: "Weird but Effective LLM Tricks: Cache Tree and Tail Prompt Optimization"
article_title: "wiki/tail-prompt-optimization-en.md at main · orbsh/wiki · GitHub"
author: "orbsh"
captured_at: "2026-07-17T01:47:30Z"
capture_tool: "hn-digest"
hn_id: 48942361
score: 1
comments: 0
posted_at: "2026-07-17T01:21:17Z"
tags:
  - hacker-news
  - translated
---

# Weird but Effective LLM Tricks: Cache Tree and Tail Prompt Optimization

- HN: [48942361](https://news.ycombinator.com/item?id=48942361)
- Source: [github.com](https://github.com/orbsh/wiki/blob/main/tail-prompt-optimization-en.md)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T01:21:17Z

## Translation

タイトル: 奇妙だが効果的な LLM のトリック: キャッシュ ツリーとテール プロンプトの最適化
記事のタイトル: メインの wiki/tail-prompt-optimization-en.md · orbsh/wiki · GitHub
説明: GitHub でアカウントを作成して、orbsh/wiki の開発に貢献します。

記事本文:
メインの wiki/tail-prompt-optimization-en.md · orbsh/wiki · GitHub
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
オーブシュ
/
ウィキ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
テールプロンプト最適化-en.md
パスをコピーする

ファイル アクションを責める さらにファイル アクションを追加する 最新のコミット
履歴 履歴 123 行 (91 loc) · 5.53 KB メイン ブレッドクラム
テールプロンプト最適化-en.md
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション キャッシュ ツリーとテール プロンプトの最適化
複数のターンが共通の KV キャッシュ プレフィックスを共有し、ツリー構造を形成します。
[システムプロンプト + 会話履歴] ← トランク (共有キャッシュ)
／＼
[分岐A:問題A] [分岐B:問題B] [分岐C:問題C] ←分岐(新たな計算)
コアプロパティ:
共有トランク: すべてのブランチが親の KV キャッシュを再利用し、冗長な計算は不要です。
独立したブランチ : 各ブランチの新しいトークンは独立して計算されます。
キャッシュの再利用: ブランチを切り替えても、共通のプレフィックス キャッシュは引き続き利用可能です (TTL 以内)
典型的なシナリオ: IM プラットフォームのスレッド メカニズム。各スレッドは、メイン スレッドのプレフィックス キャッシュを共有するブランチです。通常、基本情報は最初にキャッシュに事前に入力され (例: コーディング シナリオでのプロジェクト コード/ドキュメントのロード)、次に異なるブランチが異なるタスクを実行します。各ブランチはキャッシュされたプロジェクト コンテキストを継承し、独自のタスク トークンを計算するだけで済みます。
メインスレッド（幹）
§── thread_A: バグ修正 → トランクキャッシュを継承 + バグ説明
§── thread_B: 機能追加 → トランクキャッシュ + 要件記述を継承
└── thread_C: テストの書き込み → トランク キャッシュ + テスト スコープを継承
各スレッドのコスト ≈ タスク説明トークン (プロジェクト コンテキストはキャッシュを使用します)。
プロンプトの最後に挿入される一時的な命令。KV キャッシュ バイパス分岐メカニズムを利用して特定のタスクを実行します。
テール プロンプトは、キャッシュ ツリーの一種である Cache Vine です。 1 本の主幹が継続的に成長し、定期的に葉が発芽します。葉が落ちた後も幹は成長を続けます。

[システムプロンプト+会話履歴]──トランク（継続的に蓄積）
│
§── [テールプロンプトA + 質問] ← 葉(1ターン、落ちる)
│
幹は成長を続けています…
│
§── [尾プロンプトB + 質問] ← 葉(1ターン、落ちる)
│
幹は成長を続けています…
│
└── [テールプロンプトC + 質問] ← 葉 (1ターン、落ちる)
キャッシュ ツリーとの違い:
KVキャッシュ(変更なし): [会話履歴] ← トランク
分岐をバイパス (新しい計算): [末尾のプロンプト + ユーザーの質問] ← 葉
→ LLM は両方を 1 回のターンで完了します: ユーザーに応答し、テールプロンプトタスクを実行します。
→ターン終了後、葉が落ち（テールプロンプトは破棄）、結果のみが残る
TCO (末尾呼び出し最適化) への類似: 末尾呼び出しは現在のスタック フレームを再利用します。テールプロンプトは現在の KV キャッシュを再利用します。どちらも既存の状態を再利用する「末尾」操作です。
キャッシュ使用率: 履歴部分は KV キャッシュを使用します。末尾のプロンプトのみが新たに計算されます
バイパスブランチ : メイントランク (会話履歴) は変更せず、最後に一時的な命令のみを追加します。
ワンショット : ターン終了後にプロンプトから削除され、セッションには書き込まれません
制御 : インジェクターは、いつ、何を注入するかを決定します。エージェントは実行のみ
圧縮シナリオ: リーフが落下した場合のその場での交換
テールプロンプトはつるのパターンを使用します。つまり、1 本の主幹から葉が芽生え、その後成長し続けます。圧縮シナリオの特殊なケースです。リーフが落ちると、その前のトランクも置き換えられます (チェックポイントが古い履歴を置き換えます)。
KV キャッシュ (変更なし):
[ckpt_0] + [msg_101..msg_150]
新しく追加されたメッセージ (キャッシュされていない部分のみ):
テールプロンプト: 「まずユーザーの質問に答えてから、上記の会話を分析します。
設定/事実を抽出し、チェックポイントをマークするには、memory_store を呼び出します。」
ユーザー メッセージ: 「なぜ Fluxora のコンポーネント セットは閉じられているのですか?」

エージェントは 1 回のターンで 3 つのことを完了します (順序に注意してください)。
1.「Fluxora のコンポーネント セットは、次の理由で閉じられています。」 ← 最初にユーザーの質問に答えます
2.tool_call:memory_store(...) ←メモリ操作
3.tool_call:memory_checkpoint(...) ←最終的にチェックポイントをマークします
ターン終了後、セッションはクリーンな履歴を保存します。
セッションに保存されます:
[checkpoint_1] + ["Fluxora のコンポーネント セットが閉じられているのはなぜですか?"] + ["Fluxora のコンポーネント セットが閉じられている理由は..."]
そうではない:
[checkpoint_1] + [テールプロンプト + "Fluxora のコンポーネント セットが閉じられているのはなぜですか?"] + ["Fluxora のコンポーネント セットが閉じられている理由は..."]
テールプロンプトは破棄され、残留物はありません。
従来のアプローチとの比較
次元
従来の圧縮
テールプロンプト
LLMコール
独立した API 呼び出し
エージェントの通常のターンを再利用する
KVキャッシュ
最初から計算します (ヒット率 0%)
履歴部分はキャッシュを使用します (最大 99% のヒット率)
注意
圧縮タスクに完全に集中
分割 (ユーザーに応答 + タスクを実行)
制御
圧縮モジュールの制御
インジェクターコントロール (プロンプト + ツール経由)
参考文献
Agent-memory.md — プレフィックス チェックポイントのテール プロンプト アプリケーション
chart-memory.md — グラフ構造メモリ内の抽出プロンプト
LLM キャッシュ破壊パターン — キャッシュ分岐パターン (スレッド) と破壊パターン
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to orbsh/wiki development by creating an account on GitHub.

wiki/tail-prompt-optimization-en.md at main · orbsh/wiki · GitHub
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
orbsh
/
wiki
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
tail-prompt-optimization-en.md
Copy path Blame More file actions Blame More file actions Latest commit
History History 123 lines (91 loc) · 5.53 KB main Breadcrumbs
tail-prompt-optimization-en.md
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Cache Tree and Tail Prompt Optimization
Multiple turns share a common KV cache prefix, forming a tree structure:
[system prompt + conversation history] ← trunk (shared cache)
/ \
[branch A: question A] [branch B: question B] [branch C: question C] ← branches (new computation)
Core properties:
Shared trunk : All branches reuse the parent's KV cache, no redundant computation
Independent branches : Each branch's new tokens are computed independently
Cache reuse : When switching branches, the common prefix cache remains available (within TTL)
Typical scenario: IM platform thread mechanisms. Each thread is a branch sharing the main thread's prefix cache. Usually, base information is pre-filled into cache first (e.g., loading project code/docs in coding scenarios), then different branches execute different tasks — each branch inherits the cached project context and only needs to compute its own task tokens.
Main thread (trunk)
├── thread_A: fix bug → inherits trunk cache + bug description
├── thread_B: add feature → inherits trunk cache + requirement description
└── thread_C: write tests → inherits trunk cache + test scope
Each thread's cost ≈ task description tokens (project context uses cache).
A temporary instruction injected at the end of a prompt, leveraging the KV cache bypass branch mechanism to perform specific tasks.
Tail prompts are a variant of cache trees — Cache Vine . One main trunk grows continuously, with leaves sprouting periodically; after leaves fall, the trunk continues growing.
[system prompt + conversation history] ─── trunk (accumulates continuously)
│
├── [tail prompt A + question] ← leaf (one turn, falls off)
│
trunk continues growing...
│
├── [tail prompt B + question] ← leaf (one turn, falls off)
│
trunk continues growing...
│
└── [tail prompt C + question] ← leaf (one turn, falls off)
Differences from cache trees:
KV cache (unchanged): [conversation history] ← trunk
bypass branch (new computation): [tail prompt + user question] ← leaf
→ LLM completes both in a single turn: answer user + execute tail prompt task
→ After turn ends, leaf falls off (tail prompt discarded), only results remain
Analogy to TCO (Tail Call Optimization): tail calls reuse the current stack frame; tail prompts reuse the current KV cache. Both are "tail" operations that reuse existing state.
Cache utilization : History portion uses KV cache; only the tail prompt is newly computed
Bypass branch : Does not modify the main trunk (conversation history), only appends temporary instructions at the end
One-shot : Removed from prompt after turn ends, not written to session
Control : Injector decides when and what to inject; Agent only executes
Compression Scenario: In-Place Replacement When Leaf Falls
Tail prompts use the vine pattern — one main trunk sprouts leaves, then continues growing. The compression scenario's special case — when the leaf falls, the preceding trunk is also replaced (checkpoint replaces old history).
KV cache (unchanged):
[ckpt_0] + [msg_101..msg_150]
Newly appended messages (only uncached portion):
tail prompt: "First answer the user's question, then analyze above conversation,
call memory_store to extract preferences/facts, mark checkpoint."
user message: "Why is Fluxora's component set closed?"
Agent completes three things in one turn (note order):
1. "Fluxora's component set is closed because..." ← answer user question first
2. tool_call: memory_store(...) ← then memory operations
3. tool_call: memory_checkpoint(...) ← finally mark checkpoint
After turn ends, session stores clean history:
Stored in session:
[checkpoint_1] + ["Why is Fluxora's component set closed?"] + ["Fluxora's component set is closed because..."]
NOT:
[checkpoint_1] + [tail prompt + "Why is Fluxora's component set closed?"] + ["Fluxora's component set is closed because..."]
Tail prompt discarded, no residue.
Comparison with Traditional Approaches
Dimension
Traditional Compression
Tail Prompt
LLM Call
Independent API call
Reuses Agent's normal turn
KV cache
Computes from scratch (0% hit rate)
History portion uses cache (~99% hit rate)
Attention
Fully concentrated on compression task
Split (answer user + execute task)
Control
Compression module controls
Injector controls (via prompt + tools)
References
agent-memory.md — Tail prompt application in Prefix Checkpoint
graph-memory.md — Extraction prompts in graph-structured memory
LLM Caching Destruction Patterns — Cache branch patterns (thread) vs destruction patterns
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
