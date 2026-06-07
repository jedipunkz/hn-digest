---
source: "https://cognatoai.com"
hn_url: "https://news.ycombinator.com/item?id=48433857"
title: "Show HN: Version Control for AI Agents"
article_title: "Cognato AI - Version Controla and Agent Audit Platform"
author: "harsh020"
captured_at: "2026-06-07T12:39:24Z"
capture_tool: "hn-digest"
hn_id: 48433857
score: 1
comments: 0
posted_at: "2026-06-07T11:32:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Version Control for AI Agents

- HN: [48433857](https://news.ycombinator.com/item?id=48433857)
- Source: [cognatoai.com](https://cognatoai.com)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T11:32:21Z

## Translation

タイトル: HN の表示: AI エージェントのバージョン管理
記事のタイトル: Cognato AI - バージョン管理とエージェント監査プラットフォーム
説明: AI エージェントのバージョン管理、監査可能性、コンプライアンス
HN テキスト: Git/GitHub はエージェント時代に向けて進化していないため、構築中です

記事本文:
Cognato AI - バージョン コントロールとエージェント監査プラットフォーム AI エージェントの AI バージョン管理のための Cognato AI ログイン待機リスト コラボレーション ツール
エージェント セッションのプッシュ、プル、クローン作成。タスクの途中でモデルを切り替え、チーム全体で AI ワークフローに協力します。
mach を試してください (オープン ソース) エンタープライズ ウェイトリスト ログイン mach clone Origin/session_29a ブランチ: feature/auth-gen リモート状態の取得 モデル Gemini 1.5 の交換 エージェント タスクの stdout 再開 $ mach clone Origin/session_29a 'working-dir' へのクローン作成... リモート: エージェント状態の列挙: 45、完了。受信状態: 100% (45/45)、完了。 ✓ ブランチ「gemini-handoff」に切り替えました 問題 なぜ今なのか?
AI エージェント セッションは、孤立した不透明なブラック ボックスです。エージェントが壁にぶつかったとき、そのコンテキストを簡単に引き渡すことはできません。企業が間違いを犯した場合、企業はその理由を監査する方法がありません。従来のバージョン管理は、自律エージェントの意図ではなく、人間のコードを追跡します。
エージェントのメモリを git リポジトリのように扱います。
すべてのエージェントのアクションを暗号化して検証します。
共同エージェント ワークフローの分散プッシュ セッション: claude_ui_build
gemini-1.5-pro でタスクを再開しました
Mach は、エージェント メモリのシリアル化、バージョン管理、検証のための標準プロトコルを確立し、あらゆるモデルまたはプラットフォームにわたってシームレスな操作を可能にします。
エージェントは、コンテキスト、メモリ、アクションを Git のようなコミットとして決定的に記録する、Mach でラップされた環境で実行されます。
エージェントのすべてのステップはハッシュ化されてリンクされ、自律的な意図と実行に関する不変で検証可能な台帳が作成されます。
セッションをリモート レジストリにプッシュします。チーム メンバーは、さまざまなローカル モデルまたはリモート モデルにシームレスにプル、分岐、および交換できます。
エージェントのメモリが実行エンジンから切り離されると、まったく新しいワークフローのロックが解除されます。
クロードの長いコンテキストを使用して文書の下書きを開始します。コーディングブロッカーに遭遇しましたか?コミットする

セッションを実行し、それをプルし、Gemini または特殊な微調整モデルを使用してまったく同じコンテキスト ウィンドウを再開します。
GM Llama マルチエージェント オーケストレーション
専門の調査エージェントがセッション ツリーを生成できるようにし、独立したコーダー エージェントが特定のコミットから分岐して実装を並行して構築できるようにします。
エージェントが幻覚を起こして重要なコードを削除したのか?エラーの直前にセッション状態をチェックアウトし、正しい方向に導きます。
すべてのセッションは、入力、推論トレース、ツールの実行、出力という一連のステップの台帳です。任意のステップで分岐して、さまざまな結果を調べたり、並行プロンプトをテストしたり、タスクの途中でモデルを交換したりできます。
{ user_prompt } [ 推論 ] > ツール: read_file { session_merge } [ spawn: test_run ] > ツール: bash { 出力: 失敗 } [ swap: gemini-1.5 ] > ツール: write_file プラットフォームの概要
エージェントのメモリ、コンテキスト、タスクに対するローカルファーストのバージョン管理。
アクティブなセッションのプッシュ、プル、クローン作成
AI モデル間のタスクの引き継ぎ
プラットフォーム全体で状態を標準化する
ホストされたレジストリにより、エージェント セッションを簡単に保存、共有、追跡できます。
クラウドでホストされるセッション レジストリ
エージェントのワークフローをシームレスに共有
Webベースのセッション履歴ビューア
開発パイプラインとの統合
マルチエージェントの従業員に対する完全な監査可能性、コンプライアンス、およびアクセス制御。
ビジュアルノードグラフとリスクエンジン
AI エージェントと連携する準備はできていますか?
Enterprise プラットフォームへの早期アクセスの待機リストに参加してください。
© 2026 Cognato AI.無断転載を禁じます。

## Original Extract

Version Control, Auditability & Compliance for AI Agents

Git/GitHub did not evolve for agentic era, so we are building

Cognato AI - Version Controla and Agent Audit Platform Cognato AI Login Waitlist Collaboration Tool for AI Version Control for AI Agents
Push, pull, and clone agent sessions. Switch models mid-task and collaborate on AI workflows with your entire team.
Try mach (Open Source) Enterprise Waitlist Login mach clone origin/session_29a Branch: feature/auth-gen Fetching remote state Swapping model gemini 1.5 Resuming agent task stdout $ mach clone origin/session_29a Cloning into 'working-dir'... remote: Enumerating agent states: 45, done. Receiving state: 100% (45/45), done. ✓ Switched to branch 'gemini-handoff' The Problem Why now?
AI agent sessions are isolated, opaque black boxes. When an agent hits a wall, you can't easily hand off its context. When it makes a mistake, enterprises have no way to audit its reasoning. Traditional version control tracks human code, not autonomous agent intent.
Treat agent memory like a git repository.
Cryptographically verify every agent action.
Collaborative Agent Workflow Distributed Pushed session: claude_ui_build
Resumed task with: gemini-1.5-pro
Mach establishes a standard protocol for serializing, versioning, and verifying agent memory, allowing seamless operations across any model or platform.
Agents run in a Mach-wrapped environment that deterministically records context, memory, and actions as Git-like commits.
Every agent step is hashed and linked, creating an immutable, verifiable ledger of autonomous intent and execution.
Push sessions to the remote registry. Team members can pull, branch, and swap to different local or remote models seamlessly.
Unlock entirely new workflows when agent memory is decoupled from the execution engine.
Start drafting a document with Claude's long context. Hit a coding blocker? Commit the session, pull it, and resume the exact same context window using Gemini or a specialized fine-tuned model.
GM Llama Multi-Agent Orchestration
Allow a specialized researching agent to generate a session tree, and let independent coder agents branch from specific commits to build implementations in parallel.
Agent hallucinated and deleted critical code? Checkout the session state right before the error and steer it in the right direction.
Every session is a ledger of sequential steps: inputs, reasoning traces, tool executions, and outputs. Branch off at any step to explore different outcomes, test parallel prompts, or swap models mid-task.
{ user_prompt } [ reasoning ] > tool: read_file { session_merge } [ spawn: test_run ] > tool: bash { output: failed } [ swap: gemini-1.5 ] > tool: write_file Platform Overview
Local-first version control for agent memory, context, and tasks.
Push, pull, clone active sessions
Handoff tasks between AI models
Standardize states across platforms
Hosted registries to store, share, and track agent sessions effortlessly.
Cloud-hosted session registries
Seamless sharing of agent workflows
Web-based session history viewer
Integration with dev pipelines
Full auditability, compliance, and access control for your multi-agent workforce.
Visual node-graph & risk engine
Ready to collaborate with AI agents?
Join the waitlist for early access to the Enterprise platform.
© 2026 Cognato AI. All rights reserved.
