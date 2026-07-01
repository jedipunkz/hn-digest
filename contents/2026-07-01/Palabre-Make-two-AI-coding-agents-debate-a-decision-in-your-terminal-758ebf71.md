---
source: "https://palab.re/en"
hn_url: "https://news.ycombinator.com/item?id=48747786"
title: "Palabre – Make two AI coding agents debate a decision in your terminal"
article_title: "Palabre - Orchestrate conversations between AI agents."
author: "JuReyms"
captured_at: "2026-07-01T14:57:54Z"
capture_tool: "hn-digest"
hn_id: 48747786
score: 2
comments: 0
posted_at: "2026-07-01T14:49:58Z"
tags:
  - hacker-news
  - translated
---

# Palabre – Make two AI coding agents debate a decision in your terminal

- HN: [48747786](https://news.ycombinator.com/item?id=48747786)
- Source: [palab.re](https://palab.re/en)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T14:49:58Z

## Translation

タイトル: Palabre – 端末内で 2 人の AI コーディング エージェントに議論させて決定を下す
記事のタイトル: Palabre - AI エージェント間の会話を調整します。
説明: 2 人のエージェント間でディベートを実行するか、複数の独立した回答を含む Ask リクエストを実行します。 Palabre はローカル CLI を呼び出し、Markdown 合成をエクスポートします。

記事本文:
palabre 検索… k FR 検索… k AI エージェント間の議論を調整します。
技術的な質問を最寄りの代理店に送信してください。 2 人のエージェント間で矛盾した議論を実行したり、複数の独立した回答を含む Ask リクエストを実行したりして、忠実な比較総合を取得します。
最優先事項は二重送信の防止です。 CheckoutPanel.vue では、リクエストが保留されている間、支払いボタンはクリック可能なままであるため、素早いダブルクリックで 2 つのトランザクションをトリガーできます。
私なら UI パッチを小さくします。isSubmitting 状態、submitPayment() の実行中は無効にし、最後にブロックして、成功と失敗の両方の後に UI が回復するようにします。
私は全面的に同意しません。UI のボタンを無効にするだけでは支払いには十分ではありません。ブラウザの再試行、複数のタブ、またはサーバーの遅延によっても、2 つの有効なリクエストが作成される可能性があります。
盲点はべき等性です。 API がサポートしている場合、コンポーネントは支払い試行の開始時に生成された idempotencyKey を送信する必要があります。
良い反論だ。パッチを 2 層構造にします。ダブルクリックを防止する UI ガードと、createPaymentIntent に idempotencyKey を渡します。
回帰を防ぐために、ダブルクリック、ネットワーク障害後の再試行、1 回の試行が保留されている間に同じキーを再利用するという 3 つのテストを追加します。
ニュアンスのあるコンセンサス: UI ガードは必要ですが、十分ではありません。本当の保護はリクエストレベルの冪等性から生まれます。
注意: バックエンドがこのキーをまだサポートしていない場合は、修正をブロックしないでください。その場合は、明示的な技術ノートを付けて UI ガードを出荷してください。
CheckoutPanel.vue では二重送信が優先リスクです。
UI ロックはダブルクリックのリスクを軽減しますが、すべての再試行をカバーするわけではありません。
API がサポートしている場合、リクエスト レベルの idempotencyKey は堅牢な保護となります。
バックエンドは支払い用の冪等キーをすでに受け入れていますか?
API がサポートしていない場合は、sh

UI ガードは単独で出荷するのか、それともリリースをブロックするのか?
支払いが保留されている間はボタンを無効にし、ローカル ガードを追加します。
支払い試行ごとにべき等性キーを生成します。
API がそれを受け入れる場合は、そのキーを createPaymentIntent に渡します。
ダブルクリック、ネットワークの再試行、キーの再利用をテストします。
最良の修正は、即時の UI 保護と支払いレベルの冪等性を組み合わせたものです。この議論は、ブラウザーに限定された誤った安全感を回避します。
質問に合った形式を選択してください。
2 人のエージェントが交互に応答します。この形式は、反対意見を表面化し、決定をプレッシャーテストし、矛盾した総合を生み出すように構築されています。
複数のエージェントが、お互いに内容を読まずに同じ件名に回答します。忠実な合成を読み取る前に、独立した角度を比較します。
Codex、Claude Code、OpenCode、Mistral Vibe、Antigravity、Ollama はすべて Palabre セッションに参加できます。
オラマは地元で走っています。 OpenCode、Vibe、Antigravity、Codex も、各ツールの無料アカウント、クレジット、または割り当てに応じてテストできます。
Palabre は別のアカウント層を追加しません。マシン上にすでに存在する CLI を調整し、その制限を尊重します。
Claude Code は、使用状況に応じて適切な Claude アクセス、特に Claude Pro 以上を使用します。
被写体から合成まで、いくつかのステップで完了します。
アーキテクチャの選択、リファクタリング戦略、アプローチの比較などの質問を書いてから、2 人のエージェント間で議論するか、複数の独立した回答を求める質問するかのモードを選択します。
ディベート モードでは、パラブレが交互のターンを調整します。 Ask モードでは、各エージェントは他のエージェントの影響を受けることなく、同じ件名に単独で回答します。
Palabre は最終的な合成を生成し、応答、メタデータ、結論を含む .debate.md または .ask.md ファイルをエクスポートします。
必要なものはすべて揃っており、それ以上は何もありません
Palabre はオープンソースであり、ダイレクトです

独自のアカウント レイヤーやプロキシを追加せずに、マシンにインストールされている CLI を呼び出します。
各エージェントは独自の認証フローを保持します。 CLI がターミナルで動作する場合は、Palabre でも動作できます。
実装者、レビュー担当者、批評家、アーキテクト、スカウトなどの役割をエージェントに割り当てて、議論を形成したり、Ask の回答を比較したりできます。
クラウド エージェントとローカル モデルを組み合わせて、より柔軟でコストを意識したセッションを実現します。 Ollama エージェントは、Claude、Codex、OpenCode と同様に参加します。
--files または --context を使用してソース ファイルを渡します。エージェントは、コードベースの実際のコンテキストを使用して推論します。
各セッションでは、構造化された Markdown ファイル (ディベートの場合は .debate.md、Ask リクエストの場合は .ask.md) が生成されます。
Palabre-vscode 拡張機能は、ディベート、Ask 回答、Markdown エクスポートをエディターに直接もたらします。
VS Code から Palabre を起動し、エージェントを選択して、ワークスペースを離れることなくセッションに従います。
CLI は同じオーケストレーションとエクスポート動作を維持しながら、エディターから有用なファイルを追加します。
AI エージェントを連携させます。

## Original Extract

Run a debate between two agents or an Ask request with several independent answers. Palabre invokes local CLIs and exports a Markdown synthesis.

palabre Search… k FR Search… k Orchestrate debates between AI agents.
Send a technical question to your local agents. Run a contradictory debate between two agents or an Ask request with several independent answers, then get a faithful comparative synthesis.
The first priority is double-submit protection. In CheckoutPanel.vue, the payment button remains clickable while the request is pending, so a fast double click can trigger two transactions.
I would keep the UI patch small: an isSubmitting state, disabled while submitPayment() runs, and a finally block so the UI recovers after both success and failure.
I do not fully agree: disabling the button in the UI is not enough for payments. Browser retry, multiple tabs, or server latency can still create two valid requests.
The blind spot is idempotency. The component should send an idempotencyKey generated when the payment attempt starts, if the API supports it.
Good objection. I would make the patch two-layered: UI guard to prevent double-click, then idempotencyKey passed to createPaymentIntent.
To prevent regression, I would add three tests: double-click, retry after network failure, and the same key reused while one attempt is pending.
Consensus with nuance: the UI guard is necessary, but not sufficient. The real protection comes from request-level idempotency.
The caveat: do not block the fix if the backend does not support this key yet. In that case, ship the UI guard with an explicit technical note.
Double-submit is the priority risk in CheckoutPanel.vue.
The UI lock reduces double-click risk, but does not cover every retry.
A request-level idempotencyKey is the robust protection if the API supports it.
Does the backend already accept an idempotency key for payments?
If the API does not support it, should the UI guard ship alone or block release?
Disable the button while payment is pending and add a local guard.
Generate an idempotencyKey per payment attempt.
Pass that key to createPaymentIntent if the API accepts it.
Test double-click, network retry, and key reuse.
The best fix combines immediate UI protection with payment-level idempotency. The debate avoids a false sense of safety limited to the browser.
Choose the format that fits your question.
Two agents answer each other turn after turn. The format is built to surface objections, pressure-test a decision, and produce a contradictory synthesis.
Several agents answer the same subject without reading each other. You compare independent angles before reading a faithful synthesis.
Codex, Claude Code, OpenCode, Mistral Vibe, Antigravity, and Ollama can all participate in Palabre sessions.
Ollama runs locally. OpenCode, Vibe, Antigravity, and Codex can also be tested depending on each tool’s free accounts, credits, or quotas.
Palabre does not add another account layer: it orchestrates the CLIs already present on your machine and respects their limits.
Claude Code uses the appropriate Claude access, notably Claude Pro or higher depending on your usage.
From subject to synthesis in a few steps.
Write your question, such as an architecture choice, refactoring strategy, or comparison of approaches, then choose the mode: debate between two agents or Ask with several independent answers.
In Debate mode, Palabre orchestrates alternating turns. In Ask mode, each agent answers the same subject alone, without being influenced by the others.
Palabre produces a final synthesis, then exports a .debate.md or .ask.md file with responses, metadata, and conclusion.
Everything you need, nothing more
Palabre is open source and directly invokes the CLIs installed on your machine, without adding a proprietary account layer or proxy.
Each agent keeps its own authentication flow. If a CLI works in your terminal, it can work with Palabre.
Assign roles to agents, such as implementer, reviewer, critic, architect, or scout, to shape a debate or compare Ask responses.
Combine a cloud agent with a local model for more flexible and cost-conscious sessions. Ollama agents participate like Claude, Codex, or OpenCode.
Pass source files with --files or --context . Agents reason using the real context of your codebase.
Each session produces a structured Markdown file: .debate.md for a debate, .ask.md for an Ask request.
The Palabre-vscode extension brings debates, Ask responses, and Markdown exports directly into your editor.
Start Palabre from VS Code, select your agents, and follow the session without leaving your workspace.
Add useful files from the editor, while the CLI keeps the same orchestration and export behavior.
Make your AI agents work together.
