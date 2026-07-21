---
source: "https://www.superserve.ai/"
hn_url: "https://news.ycombinator.com/item?id=48999489"
title: "Show HN: Superserve – Firecracker microVM sandboxes for long-running AI agents"
article_title: "Superserve | The sandbox for long-running agents"
author: "pavitrabhalla"
captured_at: "2026-07-21T23:48:48Z"
capture_tool: "hn-digest"
hn_id: 48999489
score: 3
comments: 0
posted_at: "2026-07-21T22:59:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Superserve – Firecracker microVM sandboxes for long-running AI agents

- HN: [48999489](https://news.ycombinator.com/item?id=48999489)
- Source: [www.superserve.ai](https://www.superserve.ai/)
- Score: 3
- Comments: 0
- Posted: 2026-07-21T22:59:06Z

## Translation

タイトル: Show HN: Superserve – 長時間実行される AI エージェント用の Firecracker microVM サンドボックス
記事のタイトル: スーパーサーブ |長期稼働エージェント向けのサンドボックス
説明: 長時間実行されるエージェント用のオープンソース サンドボックス。
HN テキスト: やあ、HN、私は Superserve を構築しました。これは、AI エージェントがセッション時間制限なしで分離された Firecracker microVM 内に存在できるようにするコンピューティング レイヤーです。私が遭遇し続けた問題: ほとんどのサンドボックス プロバイダーは 24 時間後にエージェントを強制終了します。コードベースのリファクタリングやループ内でのテストの実行など、何日も動作する必要がある自律的なものを実行している場合、タイムアウトと状態の再構築と常に格闘することになります。 Superserve を使用すると、実行中の VM のスナップショットをいつでも作成し、それを並列ブランチにフォークして、中断したところから正確に再開できます。各エージェントは独自の VM を取得し、共有カーネルは使用しません。資格情報ブローカーもあるので、エージェントが生の API キーに直接触れることはなく、下りは指定した URL にロックされます。独自のエージェント (Claude Code、Codex など) を持参します。私たちはその下で耐久性のあるコンピューティングを処理します。 TypeScript および Python SDK、MCP サポート、オープンソース、自己ホスト可能。特に長期稼働エージェントを構築している方からのフィードバックをお待ちしています。

記事本文:
スーパーサーブ |長期稼働エージェント用のサンドボックス ドキュメント 料金表 ブログ 採用情報 お問い合わせ 418 はじめに メニューを開く [新規] Superserve 上のクロード管理対象エージェント Superserve 上のクロード管理対象エージェントの紹介 → オープンソース サンドボックス
長期にわたるエージェント向け。
無期限に一時停止し、即座に再開し、エージェントが必要とする限り実行できる耐久性の高い Firecracker VM。
あなたのデータ SUPERSERVE SDK あなたのアプリ s3 gcs r2 github gitlab SUPERSERVE SANDBOX あなたのハーネス claude code codex opencode kilo code hermes openclaw [01] 使用例 エージェントには専用のコンピュータが必要
人間と協働するエージェント向けのステートフルで安全なサンドボックス
制御されたネットワークと資格情報アクセスを備えた、分離されたステートフル サンドボックス内でエージェントを実行します。
エージェントにファイル、パッケージ、bash、動作状態への完全なアクセス権を与えます。
AI によって生成されたコード、ブラウザー、またはシステム タスクをサンドボックスで実行します。
ハーネスを外側に置き、ツール呼び出し用にサンドボックスを起動します。
ワークスペースをフォークし、分離されたサンドボックスで並列エージェントを実行し、最良の結果を選択します。
任意のステップで VM のスナップショットを作成し、多くの環境に分岐してから、以前の状態から再開します。
// Agent.ts — 長時間実行ハーネス import { Sandbox } from "@superserve/sdk" const sbx = await Sandbox .create({ name: "claude-sandbox" }) await sbx .commands.run( "npm i -g @anthropic-ai/claude-agent-sdk" ) await sbx .commands.run( "claude -p 'refactor src/**/*.ts'" ) await sbx .pause() // 同じ状態で後で再開 [02] 耐久性のある環境 エージェントの実際の動作に一致する並列および柔軟なコンピュータ
完全にカスタマイズ可能なテンプレートから
チェックポイント、フォーク、ロールバックまで
中断したところから正確に
すべてのエージェントとツールの呼び出しは独自の microVM で実行されます。共有カーネルやクロステナント状態はなく、コンテナやローカル ワークツリーよりも強力な分離が行われます。
代理人が所持してはいけないもの

資格。当社のブローカーはエージェントとエージェントが呼び出す API の間に位置し、仲介されたアクセスによる漏洩のリスクを排除します。
指定された URL および IP への出力トラフィックを制御します。すべての送信接続がログに記録され、すべてのルールがあなたのものになります。
セキュリティ ドキュメントとコンプライアンス レポートはトラスト センターで入手できます。
superserve-ai/superserve Star on GitHub → [05] プログラム制御SDK・API クラス最高のエージェントエクスペリエンス
リアルタイム出力ストリーミングを使用して、分離された環境でコードとコマンドを実行します。
セッション間でファイルの読み取り、書き込み、管理を行います。状態は、エージェントが終了したときとまったく同じように保持されます。
エージェントがボックス内からオンザフライで CPU/メモリ リソースを要求できるようにします。
インスタント プレビューのために、サンドボックス内からパブリック URL にポートを公開します。
MCP 互換エージェントからモデル コンテキスト プロトコルを介してサンドボックスに接続します。
エージェントが依存するツールを妥協することなく、サンドボックス内で Docker コンテナをネイティブに実行します。
無料で始めて、規模に応じて料金を支払います。
クレジットカードなしで始められます。秒単位の課金。
初めてのサンドボックス
数秒で。
オープン ソース 418 ✓ BYOC が利用可能 Superserve — 長時間実行されるエージェント用のオープンソース サンドボックス。

## Original Extract

The open-source sandbox for long-running agents.

Hey HN, I built Superserve, a compute layer that lets AI agents live inside isolated Firecracker microVMs with no session time limits. The problem I kept running into: most sandbox providers kill your agent after 24 hours. If you're running something autonomous that needs to work for days — refactoring a codebase, running tests in a loop — you're constantly fighting timeouts and rebuilding state. Superserve lets you snapshot a running VM at any point, fork it into parallel branches, and resume exactly where you left off. Each agent gets its own VM, no shared kernels. There's also a credentials broker so agents never touch raw API keys directly, and egress is locked to URLs you specify. You bring your own agent (Claude Code, Codex, etc.). We handle the durable compute underneath. TypeScript and Python SDKs, MCP support, open-source and self-hostable. Would love to hear feedback, especially from anyone building long-running agents.

Superserve | The sandbox for long-running agents Docs Pricing Blog Careers Contact 418 Get started Open menu [ new ] Claude Managed Agents on Superserve Introducing Claude Managed Agents on Superserve → The open-source sandbox
for long-running agents.
Durable Firecracker VMs that pause indefinitely, resume instantly, and run as long as your agents need.
YOUR DATA SUPERSERVE SDK YOUR APP s3 gcs r2 github gitlab SUPERSERVE SANDBOX YOUR HARNESS claude code codex opencode kilo code hermes openclaw [01] Use cases Your agents need their own computers
Stateful, secure sandboxes for agents that work alongside humans
Run the agent inside an isolated, stateful sandbox with controlled network and credential access.
Give the agent full access to files, packages, bash, working state.
Execute AI-generated code, browsers, or system tasks in a sandbox.
Keep the harness outside, spin up sandboxes for tool calls.
Fork a workspace, run parallel agents in isolated sandboxes, pick the best result.
Snapshot the VM at any step, branch out to many environments, then resume from a previous state.
// agent.ts — long-running harness import { Sandbox } from "@superserve/sdk" const sbx = await Sandbox .create({ name: "claude-sandbox" }) await sbx .commands.run( "npm i -g @anthropic-ai/claude-agent-sdk" ) await sbx .commands.run( "claude -p 'refactor src/**/*.ts'" ) await sbx .pause() // resume later with the same state [02] Durable environments Parallel · Flexible Computers that match how agents actually work
from fully customizable templates
to checkpoint, fork and rollback
from exactly where you left off
Every agent and tool call runs in its own microVM. No shared kernels or cross-tenant state — stronger isolation than containers or local worktrees.
Agents should not possess credentials. Our broker sits between your agent and the APIs it calls, eliminating exfiltration risk with brokered access.
Control egress traffic to specified URLs and IPs. Every outbound connection is logged, every rule is yours.
Security documentation and compliance reports available in the trust center.
superserve-ai/superserve Star on GitHub → [05] Programmatic control SDK · API Best in class agent experience
Execute code and commands in isolated environments with real-time output streaming.
Read, write, and manage files across sessions. State persists exactly as the agent left it.
Let your agents request CPU/memory resources on the fly from within the box.
Expose ports from inside a sandbox to a public URL for instant previews.
Connect to your sandboxes via the Model Context Protocol from any MCP-compatible agent.
Run Docker containers natively inside the sandbox, with no compromises on the tooling your agents depend on.
Start for free, pay as you scale.
No credit card to start. Per-second billing.
Your first sandbox
in seconds.
Open source 418 ✓ BYOC available Superserve — The open-source sandbox for long-running agents.
