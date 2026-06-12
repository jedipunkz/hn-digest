---
source: "https://traject.tamor.ai/"
hn_url: "https://news.ycombinator.com/item?id=48503930"
title: "Trajeckt: a fail-closed gateway that enforces what AI agents can do (~1.6ms)"
article_title: "Trajeckt"
author: "Bhuwan28"
captured_at: "2026-06-12T16:08:31Z"
capture_tool: "hn-digest"
hn_id: 48503930
score: 1
comments: 0
posted_at: "2026-06-12T13:39:49Z"
tags:
  - hacker-news
  - translated
---

# Trajeckt: a fail-closed gateway that enforces what AI agents can do (~1.6ms)

- HN: [48503930](https://news.ycombinator.com/item?id=48503930)
- Source: [traject.tamor.ai](https://traject.tamor.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T13:39:49Z

## Translation

タイトル: Trajeckt: AI エージェントが実行できることを強制するフェールクローズ ゲートウェイ (~1.6 ミリ秒)
記事のタイトル: Trajeckt
説明: AI エージェントの動作方法を理解し、管理し、制御します。

記事本文:
Trajeckt + + Platform Intelligence Docs リクエスト デモ AI エージェントのランタイム強制 単一のツール呼び出しは、
順序が問題になることはほとんどありません。
現在、すべてのエージェント ガードレールはアクションを一度に 1 つずつ検査するため、ステップ間でのみ存在する障害は問題なく処理されます。 Trajeckt は、実行時に、取り消し不能なアクションが開始される前に、モデルの下で軌跡全体にわたってコミットメントを強制します。
ブロックされた + + 証拠を参照してください エージェントがどのように考えているかを学びましょう
Trajeckt は、実行を構造化された証拠レイヤー (ライブ フィード、ポリシーの状態、再生履歴、より良い実行を形作る学習信号) に変換します。
裏返してください。計画は変わりません。結果はそうなります。
挿入されたステップは宣言された計画から外れていたため、起動する前にブロックされました。同じツールが計画内で許可され、計画から停止されました。
ベンチマーク run-20260604-152135 からスクリプト化
機能 1 つの施行層。
他のすべてはそれに続きます。
Trajeckt は、実行を構造化された証拠レイヤー (ライブ フィード、ポリシーの状態、再生履歴、より良い実行を形作る学習信号) に変換します。
すべてのツール呼び出しは、実行前に宣言された計画に対してチェックされます。静的な許可リストではなく、シーケンス内の位置によって許可またはブロックが決まります。
実行するたびに、検査できる意思決定、移行、結果の構造化されたタイムラインが作成されます。
あらゆる実行をリプレイします。あらゆる決定を追跡します。生のログから再構築することなく、何が起こったのか、なぜ起こったのかを正確に理解します。
このページのすべての数値は、シミュレートされたシナリオではなく、実際のタスクに対して測定された、固定ベンチマークの実行に基づいています。
ライブラリではなくゲートウェイ: スタックが MCP または OpenAI 互換ツール呼び出しを話す場合、適用は 1 つの構成変更です。
任意の MCP クライアント。強調表示された行はサーバー URL であり、唯一の変更点です。
エージェントが準備完了したら、あらゆる軌道を運用に変えます

知能。
エージェントの可観測性、ガバナンス、調査のために構築されたコントロール サーフェスから始めます。 Trajeckt を使用すると、何が起こったのかを確認し、何が起こるべきかを強制し、次に起こることを改善することが容易になります。
サインイン 軌跡ベースの分析 人間による承認 インシデントのリプレイ ポリシーの適用 プラットフォームのスナップショット リアルタイムのクラスター ステータス ライブ アクティブな軌跡 0 保留中の承認 0 ブロックされたアクション 0 インシデントのリプレイ 0 デプロイ準備完了 · 0 インシデント · すべてクリア $ trajecktdeploy --prod --region us-east-1 推測に頼らないエージェント操作。

## Original Extract

Understand, govern, and control how AI agents operate.

Trajeckt + + Platform Intelligence Docs Request Demo Runtime Enforcement For AI Agents A single tool call is
rarely the problem The sequence is.
Every agent guardrail today inspects actions one at a time — so the failures that only exist across steps walk right through. Trajeckt enforces commitments across the whole trajectory, at runtime, below the model, before any irreversible action fires.
See it Blocked + + Proof Learn how your agents think
Trajeckt turns executions into a structured evidence layer: live feeds, policy state, replay history, and the learning signals that shape better runs.
Flip it. The plan doesn't change. The outcome does.
The injected step was off the declared plan, so it was blocked before it fired: the same tool, allowed in plan and stopped out of it.
Scripted from benchmark run-20260604-152135
Capabilities One Enforcement Layer.
Everything else Follows.
Trajeckt turns executions into a structured evidence layer: live feeds, policy state, replay history, and the learning signals that shape better runs.
Every tool call is checked against the declared plan before it executes. Position in the sequence determines allow or block, not a static allowlist.
Every execution creates a structured timeline of decisions, transitions, and outcomes you can inspect.
Replay any run. Trace every decision. Understand exactly what happened and why, without reconstructing from raw logs.
Every number on this page traces to a fixed benchmark run, measured against real tasks, not simulated scenarios.
A gateway, not a library: if your stack speaks MCP or OpenAI-compatible tool calls, enforcement is one config change.
Any MCP client. The highlighted line is the server URL — the only change.
Ready when your agents are Turn every trajectory into operating intelligence.
Start with a control surface built for agent observability, governance, and investigation. Trajeckt makes it easier to see what happened, enforce what should happen, and improve what happens next.
Sign in Trajectory-based analytics Human approvals Incident replay Policy enforcement Platform snapshot Real-time cluster status live Active trajectories 0 Pending approvals 0 Blocked actions 0 Incident replays 0 Deploy ready · 0 incidents · All clear $ trajeckt deploy --prod --region us-east-1 Agent operations, without the guesswork.
