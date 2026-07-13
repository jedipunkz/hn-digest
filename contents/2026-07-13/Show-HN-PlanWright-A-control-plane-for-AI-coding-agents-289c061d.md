---
source: "https://planwright.tools"
hn_url: "https://news.ycombinator.com/item?id=48897969"
title: "Show HN: PlanWright – A control plane for AI coding agents"
article_title: "PlanWright — The control plane for autonomous software labor"
author: "dudemanAtl"
captured_at: "2026-07-13T20:51:30Z"
capture_tool: "hn-digest"
hn_id: 48897969
score: 5
comments: 2
posted_at: "2026-07-13T19:59:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PlanWright – A control plane for AI coding agents

- HN: [48897969](https://news.ycombinator.com/item?id=48897969)
- Source: [planwright.tools](https://planwright.tools)
- Score: 5
- Comments: 2
- Posted: 2026-07-13T19:59:50Z

## Translation

タイトル: 表示 HN: PlanWright – AI コーディング エージェント用のコントロール プレーン
記事のタイトル: PlanWright — 自律的なソフトウェア レイバーのためのコントロール プレーン
説明: ポストカンバン、エージェントネイティブのプランニング。人間は目標を書きます。 Claude Code、Cursor、Codex は分解、実行、チェックインされます。すべての決定が署名されます。すべての変更が監査されます。
HN テキスト: Agentic Engineering 用の MCP 駆動のコントロール プレーン。 Claude Desktop から計画し、Codex で実装し、カスタム トリアージ エージェントでレビューします。すべて MCP 経由で、各エージェントが行ったすべての決定を完全に文書化して記録および追跡されます。

記事本文:
PlanWright — 自律型ソフトウェア レイバーのコントロール プレーン PlanWright Claude Desktop ドキュメントの比較 価格 GitHub サインイン エージェント レイバーのコントロール プレーン
コーディング エージェントの速度が 10 倍になりました。あなたのパイプラインはそうではありませんでした。
スピードは中央に移り、セレモニーを後にした。 Planwright は、依然として人間のペースにある 2 つを逆転させます。記録、資料、電子メール、Slack などの混乱を目的に変え、受け入れをトリアージして、判断が必要なものだけを判断できるようにします。計画と承認は最終的にエージェントのスピードで進み、すべての裁定が署名されます。
あなたはエージェントを買収しました。
スピードはどこへ行った？
以前はコーディングがボトルネックでした。もうない。これで、制約は自動化していないすべてのものになります。
人間が UI の速度でタスクを手入力し、数千倍の速度で実行されるマシンにデータを供給します。
上級エンジニアがプル リクエストを一度に 1 つずつ読んでいます。これはまさに速度が失われるところです。
中央を自由かつ無限に速くすると、端は天井になります。それはツールのギャップではありません。それはアムダールの法則です。
コーディングエージェントが削除できない 2 つの儀式を逆転させます。
フロントブックエンド反転プランニング: カオスをイン、目標をアウト。
計画のインプットは決して整然としたものではありません。それは会議の議事録、資料、電子メールのスレッド、Slack などです。 Planwright は、Claude Desktop 内から、その生の信号を機械チェック可能な許容基準を備えた構造化された目標に合成します。古い儀式では、人間が計画を書き、ツールがそれを記録していました。プランライトはそれを逆転させます。エージェントが草案を作成し、あなたが承認するのです。あなたが望むものを説明するスピードで。
バックブックエンド 逆受容: トリアージ、溺れないでください。
目標が完了すると、プランライトはそれらを拾い上げ、クロード・コワークで受け入れをトリアージします。つまり、単純で機械的な基準を精査し、安全にクリアできるものをクリアし、次に重要で判断に縛られた目標をルーティングします。

emsをあなたに。マシンは受け付けません。それはソートです。すべての差分を読むのをやめ、実際に人間が必要なものだけを判断し始めるため、乏しい判断力が氾濫するのではなく集中してしまいます。すべての判決は署名されています。
機械はそれを受け入れません - 分類します。
ステートマシンは、作業が「完了」する前に、すべての判断を人間が行うことを保証します。機械的な基準は自動的にクリアされます。重要な決定はあなたに委ねられます。あなたのサインオフは、SOC 2 監査人が探している変更承認のコントロール ポイントです。
人間が目標を設定し、エージェントの迎えのスケジュールを設定します。
進行中のエージェント エージェントは、MCP 経由で作業、計画、コード、および記録の差分を要求します。
承諾 ヒューマンエージェントが承諾を要求します。人間が差分を確認し、テストし、承認します。
完了 人間 人間は承認します。監査記録は封印され、ハッシュチェーン化されます。
レーンの移行ごとに、ハッシュチェーンされた暗号署名された監査レコードが生成されます。チェーンはパブリック トラスト ページで独立して検証できます。
3 つのプリミティブ。 1 つのスループットのロックを解除します。
Claude Desktop では、Planwright がトランスクリプト、資料、電子メール、Slack を機械チェック可能な承認基準を備えた構造化された目標に統合します。タスクを手動で作成するのではなく、意図を承認します。
トリアージの受け入れ、それに溺れないでください。
目標が完了すると、プランライトはクロード・コワークの機械的基準を精査し、安全にクリアできるものをクリアし、判断の呼び出しのみをあなたにルーティングします。機械はそれを受け入れません - 分類します。
機械によるスループットの責任は維持されます。人間による各決定は、アドホックな PR 読み取りではなく、ハッシュ チェーンされたレコードです。
目標は流れます。エージェントが実行します。人間は承認します。
ECDSA P-256 監査署名を実装する
ワークスペースの RLS ポリシーを追加する
望む結果を説明してください。ステップではありません。実装ではありません。エージェントが行うべき明確で原子的な目標

独自のエンドツーエンド。
Claude Code は Planwright ボードに接続し、スケジュールされた目標を取得して作業を開始します。チケットのコピー＆ペーストはできません。コンテキスト切り替えはありません。
エージェントが計画、実行、承認を要求
エージェントは目標を分解し、コードを作成し、テストを実行して、それを受け入れレーンに移動します。監査チェーン内のすべてのステップが記録されます。
すべての承認は暗号署名です。チェックボックスではありません。あなたのアイデンティティが結びついた本当の承認。
すべてのステップはハッシュチェーンされ、署名されています
目標の作成から最終的なマージまでの不変の監査証跡。それ自体を書き込む SOC 2 の証拠。
SOC 2 監査人は、AI が生成したコードの水準を引き上げています。
2026 年のトラスト サービス基準に基づき、監査人は、AI 主導の変更に対する不変で改ざんの明らかな監査証跡と文書化された変更承認コントロールをますます期待しています。ほとんどの PM ツールではこれらを生成できません。 Planwright は、最初の目的から最終マージまで、これをコア プリミティブ (ハッシュ チェーン、ECDSA P-256 署名) として出荷します。
トークンはありません。構成ファイルを手動で編集する必要はありません。インストール コマンドを実行します。ブラウザが GitHub ログイン用に開き、接続されます。
クロード mcp add planwright --transport http https://mcp.planwright.tools/mcp Zero-config をコピーします。このコマンドを実行すると、GitHub ログイン用のブラウザが開きます。接続は将来のセッションのためにキャッシュされます。
ブラウザーを開けない CI およびサービス アカウントの場合は、サインイン後に [設定] → [MCP トークン] で静的トークンを生成します。監査レコードはトークン作成者に属します。
さらに詳しい情報が必要ですか? Claude Code セットアップガイドの全文を読む→
無料で始めましょう。準備ができたらスケールします。
すべてのプランには、MCP サーバー、監査チェーン、およびボードが含まれています。シートと保持深さの料金を支払います。
エージェントネイティブのプランニングを試みる個人開発者向け。
小規模チームの場合、エージェントと毎日出荷します。
コンプライアンスが必要なチーム向け

箱。
規制要件のある組織向け。
SOC 2 認証レター (ロードマップ)
AI を搭載して出荷するチームが尋ねる質問。
Claude Desktop から Planwright ボードを管理できますか? Cursor または Codex は Planwright とどのように統合されますか?これは Linear の組み込みエージェントとどう違うのですか? 「AI クライアント ネイティブ」とは実際には何を意味しますか?監査チェーンは AI 生成コードに対する SOC 2 トラスト サービス基準を満たしていますか? HIPAA 規制のチームに対する受け入れゲートはどのように機能しますか? FedRAMP 中程度のワークロードに Planwright を使用できますか?これを規制チームのリニアやプレーンと比較するとどうですか?監査チェーンと git ログの違いは何ですか? FINRA 規制のチームは Planwright の AI エージェントを使用できますか?エアギャップ環境用のセルフホスト型オプションはありますか?いくらかかりますか？ PlanWright ポストカンバン時代に向けて構築されました。
© 2026 プランライト。無断転載を禁じます。

## Original Extract

Post-kanban, agent-native planning. Humans write objectives. Claude Code, Cursor, and Codex decompose, execute, and check in. Every decision signed. Every change audited.

MCP driven control plane for Agentic Engineering. Plan from Claude Desktop, implement in Codex, review in a custom triage agent. All via MCP, all logged and tracked with full documentation of all decisions made by each agent.

PlanWright — The control plane for autonomous software labor PlanWright Claude Desktop Compare Docs Pricing GitHub Sign in The Control Plane for Agent Labor
Your coding agents got 10x faster. Your pipeline didn't.
Speed moved to the middle and left the ceremonies behind. Planwright inverts the two that are still human-paced: it turns your chaos — transcripts, decks, email, Slack — into objectives, and triages acceptance so you rule only on what needs judgment. Planning and acceptance finally move at agent speed, and every ruling is signed.
You bought the agents.
Where did the speed go?
Coding used to be the bottleneck. Not anymore. Now the constraint is everything you didn't automate:
A human hand-typing tasks at UI speed to feed machines running thousands of times faster.
A senior engineer reading pull requests one at a time — which is exactly where velocity goes to die.
Make the middle free and infinitely fast, and the edges become the ceiling. That's not a tooling gap. It's Amdahl's law.
Invert the two ceremonies coding agents can't remove.
Front bookend Invert planning: chaos in, objectives out.
Your planning input was never tidy — it's meeting transcripts, decks, email threads, Slack. From inside Claude Desktop, Planwright synthesizes that raw signal into structured objectives with machine-checkable acceptance criteria. The old ceremony had a human write the plan and the tool record it. Planwright flips it: the agent drafts, you approve — at the speed of describing what you want.
Back bookend Invert acceptance: triage, don't drown.
When objectives complete, Planwright picks them up and triages acceptance in Claude Cowork — combing the plain, mechanical criteria and clearing what it safely can, then routing the critical, judgment-bound items to you. The machine doesn't accept; it sorts. You stop reading every diff and start ruling only on what actually needs a human, so your scarce judgment gets concentrated instead of flooded. Every ruling is signed.
The machine doesn't accept — it sorts.
The state machine ensures a human rules on every judgment call before work reaches “done.” Mechanical criteria get cleared automatically; critical decisions route to you. Your sign-off is the change-approval control point your SOC 2 auditor is looking for.
Human sets the objective and schedules it for agent pickup.
In Progress Agent Agent claims the work, plans, codes, and records diffs via MCP.
Acceptance Human Agent requests acceptance. A human reviews the diff, tests, and signs off.
Done Human Human approves. The audit record is sealed and hash-chained.
Every lane transition produces a hash-chained, cryptographically signed audit record. The chain is independently verifiable on the public trust page .
Three primitives. One throughput unlock.
In Claude Desktop, Planwright synthesizes your transcripts, decks, email, and Slack into structured objectives with machine-checkable acceptance criteria. You approve intent instead of hand-authoring tasks.
Triage acceptance, don't drown in it.
When objectives complete, Planwright combs the mechanical criteria in Claude Cowork and clears what it safely can, routing only the judgment calls to you. The machine doesn't accept — it sorts.
Machine-cleared throughput stays accountable: each human decision is a hash-chained record, not an ad-hoc PR read.
Objectives flow. Agents execute. Humans approve.
Implement ECDSA P-256 audit signing
Add RLS policies for workspaces
Describe the outcome you want. Not the steps. Not the implementation. A clear, atomic objective your agent can own end-to-end.
Claude Code connects to your Planwright board, picks up scheduled objectives, and starts working. No copy-pasting tickets. No context-switching.
Agent plans, executes, requests acceptance
The agent decomposes the objective, writes code, runs tests, and moves it to the acceptance lane. Every step recorded in the audit chain.
Every acceptance is a cryptographic signature. Not a checkbox. A real approval with your identity bound to it.
Every step hash-chained and signed
Immutable audit trail from objective creation to final merge. SOC 2 evidence that writes itself.
SOC 2 auditors are raising the bar on AI-generated code.
Under the 2026 Trust Services Criteria, auditors increasingly expect immutable, tamper-evident audit trails and documented change-approval controls for AI-driven changes. Most PM tools can't produce them. Planwright ships this as a core primitive — hash-chained, ECDSA P-256-signed, from the first objective to the final merge.
No tokens. No config files to hand-edit. Run the install command — your browser opens for GitHub login, and you're connected.
claude mcp add planwright --transport http https://mcp.planwright.tools/mcp Copy Zero-config. Run this command and your browser opens for GitHub login. The connection is cached for future sessions.
For CI and service accounts that can't open a browser, generate a static token in Settings → MCP Token after signing in. Audit records attribute to the token creator.
Need more detail? Read the full Claude Code setup guide →
Start free. Scale when you're ready.
Every plan includes the MCP server, audit chain, and board. Pay for seats and retention depth.
For solo developers trying agent-native planning.
For small teams shipping with agents every day.
For teams that need compliance out of the box.
For organizations with regulatory requirements.
SOC 2 attestation letter (roadmap)
Questions teams shipping with AI ask.
Can I manage my Planwright board from Claude Desktop? How does Cursor or Codex integrate with Planwright? How is this different from Linear's built-in agents? What does 'AI-client-native' mean in practice? Does the audit chain satisfy SOC 2 Trust Services Criteria for AI-generated code? How does the acceptance gate work for HIPAA-regulated teams? Can we use Planwright for FedRAMP Moderate workloads? How does this compare to Linear or Plane for regulated teams? What's the difference between the audit chain and a git log? Can FINRA-regulated teams use AI agents with Planwright? Is there a self-hosted option for air-gapped environments? How much does it cost? PlanWright Built for the post-kanban era.
© 2026 PlanWright. All rights reserved.
