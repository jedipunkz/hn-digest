---
source: "https://byteiota.com/claude-code-v2-1-172-sub-agents-can-now-spawn-sub-agents/"
hn_url: "https://news.ycombinator.com/item?id=48522449"
title: "Claude Code v2.1.172: Sub-Agents Can Now Spawn Sub-Agents"
article_title: "Claude Code v2.1.172: Sub-Agents Can Now Spawn Sub-Agents | byteiota"
author: "sscaryterry"
captured_at: "2026-06-14T01:02:29Z"
capture_tool: "hn-digest"
hn_id: 48522449
score: 2
comments: 0
posted_at: "2026-06-13T23:17:22Z"
tags:
  - hacker-news
  - translated
---

# Claude Code v2.1.172: Sub-Agents Can Now Spawn Sub-Agents

- HN: [48522449](https://news.ycombinator.com/item?id=48522449)
- Source: [byteiota.com](https://byteiota.com/claude-code-v2-1-172-sub-agents-can-now-spawn-sub-agents/)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T23:17:22Z

## Translation

タイトル: クロード コード v2.1.172: サブエージェントがサブエージェントを生成できるようになりました
記事のタイトル: クロード コード v2.1.172: サブエージェントがサブエージェントを生成できるようになりました |バイトイオタ

記事本文:
クロード コード v2.1.172: サブエージェントがサブエージェントを生成できるようになりました。バイトイオタ
AI と開発 コンピューター ビジョン
開発者エクスペリエンス 開発者ツール
ニュースと分析 業界分析
クロード コード v2.1.172: サブエージェントがサブエージェントを生成できるようになりました
2 年間、Anthropic は例外なくクロード コードの 1 つのルールを適用しました。それは、サブエージェントは別のサブエージェントを生成できないというものです。 6 月 10 日にリリースされたバージョン 2.1.172 は、それを静かに終わらせました。変更ログのエントリは 1 行です。エージェント ワークフローの設計方法には影響しません。
サブエージェントは、最大 5 レベルの深さでネストして独自のサブエージェントを起動できるようになりました。 5 レベルの上限はサーバー側で強制されるものであり、上限を上げたり下げたり無効にしたりする設定はありません。 Anthropic で Claude Code を率いる Boris Cherny 氏は、その動機をはっきりとこう説明しました。「エージェントが、より適切にコンテキストを管理する方法として、エージェントを解雇するのです。」
その枠組みは保持する価値があります。この機能は、より多くの作業を並行して実行することではありません。それは、主要な会話が見えない場所で騒々しい作業を実行することです。このリリースより前は、ログ分析を処理するサブエージェントは、その親のコンテキストに何千もの生のログ トークンを大量に送信していました。その後、親は追加のトークンを書き込み、有用な作業を続行する前に自らを再接地させます。ネストされたサブエージェントは、ログ リーダーを独自のコンテキスト フレームに分離し、概要のみを上向きに返すことでこの問題を解決します。
5 フレーム スタック: 持続するメンタル モデル
ネストされたサブエージェントは、5 フレームの深さ制限を持つ呼び出しスタックの再帰と考えてください。各フレームには、独自のシステム プロンプト、独自のモデル選択、および独自の 200K トークン コンテキスト ウィンドウが含まれます。親はリーフが返すもののみを読み取ります。検索、ファイルの読み取り、中間推論など、その間にあるすべての処理にトークンがかかり、その後、親の視界から消えます。
実践

l デバッグチェーンは次のようになります。
メインセッション (Opus) → トリアージリード (Opus) → repro-runner (Sonnet) → log-summariser (Haiku)
レイヤ 1 がインシデントを読み込みます。レイヤ 2 は、4 つの原因候補にわたってファンアウトします。ログの相関関係を調査するエージェント (ノイズが多くトークンを大量に使用するタスク) は、レイヤー 3 Haiku エージェントを生成して grep 作業を実行し、1 行の結果を返します。メインスレッドは、生のログ出力の 50,000 トークンではなく、4 つのクリーン判定を確認します。これが使用例です。エージェントを追加するのではなく、既存のエージェントからよりクリーンな結果が得られます。
最も有用なチェーンは深さ 2 ～ 3 に存在します。 5 レベルは目標ではなく天井です。
トークンの計算: 何かをネストする前にこれをお読みください
ネストすると、レベルごとのブランチごとにトークンの消費量が約 7 倍になります。これは大まかな見積もりではなく、エージェントのオーケストレーション、コンテキストの初期化、およびフレーム間のサマリーの受け渡しによって観察されるオーバーヘッドです。急速に化合します。
コミュニティ フォーラムの開発者は、気づく前に 1 分あたり 887,000 トークンに達していたと説明しました。金融サービス チームは、23 のサブエージェントとともに「シンプルな」コード品質プロジェクトを実行し、3 日後に 47,000 ドルの請求書を受け取りました。 5 レベルの深さのキャップにより、無限ループを防ぎます。午前 3 時に請求アラートがトリガーされるのを防ぐことはできません。
本番環境で何かをネストする前に、支出制限を設定します。 Claude Code の課金設定はセッションごとの上限をサポートしています。それらを使用してください。
ネストされたサブエージェントを構成する方法
サブエージェント定義は、プロジェクト レベルでは .claude/agents/<name>.md に、ユーザー スコープでは ~/.claude/agents/<name>.md に存在します。ツールの Agent() フィールドは、このリリースの新機能です。これは、このエージェントが生成できるサブエージェント タイプを制御する許可リストです。
---
名前: トリアージリード
モデル: クロード-作品-4-6
ツール:
- 読む
- バッシュ
- エージェント(再現実行者、ログ要約者)
---
あなたはデバッグです

トリアージのリーダーを務める。インシデントをロードし、再現ランナーに委任します
再現の場合は、ログの関連付けのために log-summariser に委任し、返します。
単一の判定: 確認済み、未確認、またはさらにデータが必要。
コストが重要な場合、レベル間のモデル階層化はオプションではありません。オーケストレーション層では Opus を実行し、中間レベルの実装作業には Sonnet を、ログ読み取り、grep、テスト生成などのリーフ タスクには Haiku を実行します。公式サブエージェントのドキュメントには、完全な構成仕様が記載されています。実際には、段階的ルーティングのコストはセッションあたり約 0.98 ドルですが、均一の Opus の場合は 2.02 ドルです。これはリーフ タスクの品質を損なうことなく 51% 削減されます。
円形スポーンによるトークンの爆発。深さの制限は、循環ロジックではなく、無限のネストを防止します。 1 つのプロダクション ケース: トリアージ エージェントが「一般研究者」を生成し、それが「調査スペシャリスト」を生成し、さらに「ログ リーダー」が生成され、さらに一般研究者が生成されます。深さは 4 レベル、循環型、有用な出力はゼロ、コストはかなりかかります。明示的な Agent() ホワイトリストを使用します。モデルが指定された子を持つタスクを完了できない場合は、タスクの説明を修正します。許可リストを拡大しないでください。
深さ 5 でサイレント障害。レベル 5 エージェントがレベル 6 を生成しようとすると、エラーが発生します。明示的に処理しないと、エラーは静かに伝播するか、トップレベルで予期しない出力として表面化します。複雑な階層を運用環境にデプロイする前に、カナリア チェーンを使用してネストの深さを明示的にテストします。
すでに短いタスクをネストする。ブランチごとのオーバーヘッドにより、約 1,000 トークン未満の出力を生成するタスクではネストのコストが高くなります。子タスクが短い場合は、インラインで実行します。 10 行の grep を独自のコンテキスト フレームにネストすると、分離の価値よりもオーケストレーションのオーバーヘッドが高くなります。
npm i 経由で v2.1.172 に更新します。

nstall -g @anthropic-ai/claude-code 。 GitHub の完全なリリース ノートには、Bedrock 領域の解像度の向上やアイドル状態の CPU の削減など、30 の変更すべてが記載されており、どちらも読む価値があります。 DevelopersIO の内訳には、ネストされたサブエージェントを可能にしたインフラストラクチャの変更に関する最も詳細な技術分析が含まれています。
ネストされたサブエージェントの場合: この機能は、スループットではなくコンテキスト汚染がボトルネックであるワークフローに最適です。 1 つのエージェントが生データに埋もれているためにメイン セッションがフォーカスを失っている場合は、ネストが適切な解決策です。より多くのタスクを並列実行するためにネストすることを考えている場合、それはフラット並列実行であり、別の機能です。
深さ 2 ～ 3 は、この機能がオーバーヘッドを得る場所です。必要なときに 5 つのレベルが表示されます。ほとんどのチームはしばらくはそれを必要としません。
RoguePlanet CVE-2026-47281: 今すぐ Windows Defender にパッチを適用してください
Kiro Web が GitLab をサポート: ブラウザーの仕様
Pyodide 314.0: Python ブラウザのパッケージ化がついに実現
Google Colab CLI: ターミナルから A100s および H100s を実行する
MetaMask エージェント ウォレット: AI エージェントはガードレール付きの暗号ウォレットを入手
Python 3.14 T-Strings: より安全な SQL、HTML、AI プロンプト
TensorZero がシャットダウン: OSS LLMOps が生き残れないもの
Gemini-SQL2: Google の Text-to-SQL は BIRD で 80% をクリア
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
Gemini 管理エージェント: 1 つの API 呼び出しで AI エージェントをデプロイ
Claude 管理エージェント: Cron スケジュールと Credential Vault の説明
Gemini 3.5 Pro: 2M コンテキスト、深い思考、および Flash 対 Pro の決定
Fable 5 の禁止: 米国政府の指令とクロード API の修正
Kim K2.7-Code: Moonshot の Open-Weight 1T コーディング エージェント
カーソル共有キャンバス: IDE からチーム ダッシュボードまで
AIと開発

コンピュータビジョン
開発者のエクスペリエンス
オープンソース
最小読み取り
© 2021 バイトイオタ | byteiotaによって設計および開発されました
読書中
クロード コード v2.1.172: サブエージェントがサブエージェントを生成できるようになりました
8 分で読めます
AI と開発 コンピューター ビジョン
開発者エクスペリエンス 開発者ツール
ニュースと分析 業界分析
AIと開発
コンピュータビジョン
開発者のエクスペリエンス
開発者ツール
ニュースと分析
業界分析
AIと開発
コンピュータビジョン
開発者のエクスペリエンス
開発者ツール
ニュースと分析
業界分析
Pyodide 314.0: Python ブラウザのパッケージ化がついに実現
Google Colab CLI: ターミナルから A100s および H100s を実行する
MetaMask エージェント ウォレット: AI エージェントはガードレール付きの暗号ウォレットを入手
Python 3.14 T-Strings: より安全な SQL、HTML、AI プロンプト
TensorZero がシャットダウン: OSS LLMOps が生き残れないもの

## Original Extract

Claude Code v2.1.172: Sub-Agents Can Now Spawn Sub-Agents | byteiota
AI & Development Computer Vision
Developer Experience Developer Tools
News & Analysis Industry Analysis
Claude Code v2.1.172: Sub-Agents Can Now Spawn Sub-Agents
For two years, Anthropic enforced one rule in Claude Code without exception: a sub-agent cannot spawn another sub-agent. Version 2.1.172, released June 10, quietly ended that. The changelog entry is a single line. The implications for how you architect agentic workflows are not.
Sub-agents can now launch their own sub-agents, nesting up to five levels deep. The five-level cap is hard — enforced server-side, with no setting to raise, lower, or disable it. Boris Cherny, who leads Claude Code at Anthropic, described the motivation plainly: “agents kicking off agents as a way to better manage context.”
That framing is worth holding onto. The feature is not about running more work in parallel. It is about running the noisy work somewhere the main conversation cannot see it. Before this release, a sub-agent handling log analysis would flood its parent’s context with thousands of raw log tokens. The parent then burned additional tokens re-grounding itself before continuing useful work. Nested sub-agents solve that by isolating the log-reader in its own context frame and returning only a summary upward.
The 5-Frame Stack: A Mental Model That Holds
Think of nested sub-agents as call-stack recursion with a five-frame depth limit. Each frame carries its own system prompt, its own model selection, and its own 200K token context window. The parent reads only what the leaf returns. Everything in between — the searches, the file reads, the intermediate reasoning — costs tokens and then disappears from the parent’s view.
A practical debugging chain looks like this:
main session (Opus) → triage-lead (Opus) → repro-runner (Sonnet) → log-summariser (Haiku)
Layer 1 loads the incident. Layer 2 fans out across four candidate causes. The agent investigating log correlations — a noisy, token-heavy task — spawns a Layer 3 Haiku agent to do the grep work and return a one-line result. The main thread sees four clean verdicts, not 50,000 tokens of raw log output. That is the use case: not more agents, but cleaner results from the agents you already have.
Most useful chains live at depth 2–3. Five levels is the ceiling, not the target.
Token Math: Read This Before You Nest Anything
Nesting multiplies token consumption at roughly 7× per branch per level. That is not a rough estimate — it is the observed overhead from agent orchestration, context initialization, and summary-passing between frames. It compounds fast.
A developer on the community forums described hitting 887,000 tokens per minute before noticing. A financial services team ran a “simple” code quality project with 23 sub-agents and received a $47,000 invoice three days later . The five-level depth cap prevents infinite loops. It does not prevent your billing alert from triggering at 3 AM.
Set a spend limit before nesting anything in production. Claude Code’s billing settings support per-session caps. Use them.
How to Configure Nested Sub-Agents
Sub-agent definitions live in .claude/agents/<name>.md at the project level or ~/.claude/agents/<name>.md for user scope. The Agent() field in tools is new as of this release — it is the allowlist controlling which sub-agent types this agent is permitted to spawn.
---
name: triage-lead
model: claude-opus-4-6
tools:
- Read
- Bash
- Agent(repro-runner, log-summariser)
---
You are a debugging triage lead. Load the incident, delegate to repro-runner
for reproduction, delegate to log-summariser for log correlation, and return
a single verdict: confirmed, unconfirmed, or needs-more-data.
Model tiering across levels is not optional if cost matters. Run Opus at the orchestration layer, Sonnet for mid-level implementation work, and Haiku for leaf tasks like log reads, grep, and test generation. The official sub-agents documentation covers the full configuration spec. In practice, tiered routing costs roughly $0.98 per session versus $2.02 for uniform Opus — a 51% reduction with no quality loss on leaf tasks.
Token explosion via circular spawning. The depth cap prevents infinite nesting, not circular logic. One production case: a triage agent spawned a “general researcher” that spawned an “investigation specialist” that spawned a “log reader” that spawned a general researcher again. Four levels deep, circular, zero useful output, substantial cost. Use explicit Agent() allowlists. If the model cannot complete the task with the specified children, fix the task description — do not widen the allowlist.
Silent failures at depth five. When a Level-5 agent attempts to spawn a Level-6, it receives an error. Without explicit handling, that error propagates silently or surfaces as unexpected output at the top level. Test your nesting depth explicitly with a canary chain before deploying complex hierarchies in production.
Nesting tasks that are already short. The per-branch overhead makes nesting expensive for any task producing less than roughly 1,000 tokens of output. If the child task is short, do it inline. Nesting a ten-line grep into its own context frame costs more in orchestration overhead than the isolation is worth.
Update to v2.1.172 via npm install -g @anthropic-ai/claude-code . The full release notes on GitHub cover all 30 changes, including Bedrock region resolution improvements and idle CPU reduction — both worth reading. The DevelopersIO breakdown has the most detailed technical analysis of the infrastructure changes that made nested sub-agents possible.
On nested sub-agents: the feature is well-suited for workflows where context pollution is the bottleneck, not throughput. If your main session is losing focus because one agent is drowning it in raw data, nesting is the right solution. If you are thinking about nesting to run more tasks in parallel, that is flat parallel execution — a different feature.
Depth 2–3 is where this feature earns its overhead. Five levels is there when you need it. Most teams will not need it for a while.
RoguePlanet CVE-2026-47281: Patch Windows Defender Now
Kiro Web Now Supports GitLab: Specs in Your Browser
Pyodide 314.0: Python Browser Packaging Just Got Real
Google Colab CLI: Run A100s and H100s From Your Terminal
MetaMask Agent Wallet: AI Agents Get a Crypto Wallet With Guardrails
Python 3.14 T-Strings: Safer SQL, HTML, and AI Prompts
TensorZero Shuts Down: What OSS LLMOps Can’t Survive
Gemini-SQL2: Google’s Text-to-SQL Clears 80% on BIRD
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
Gemini Managed Agents: Deploy an AI Agent with One API Call
Claude Managed Agents: Cron Schedules and Credential Vaults, Explained
Gemini 3.5 Pro: 2M Context, Deep Think, and the Flash-vs-Pro Decision
Fable 5 Banned: US Gov Directive and Your Claude API Fix
Kimi K2.7-Code: Moonshot’s Open-Weight 1T Coding Agent
Cursor Shared Canvases: From IDE to Team Dashboard
AI & Development
Computer Vision
Developer Experience
Open Source
min read
© 2021 Byteiota | Designed & Developed by byteiota
Now Reading
Claude Code v2.1.172: Sub-Agents Can Now Spawn Sub-Agents
8 min read
AI & Development Computer Vision
Developer Experience Developer Tools
News & Analysis Industry Analysis
AI & Development
Computer Vision
Developer Experience
Developer Tools
News & Analysis
Industry Analysis
AI & Development
Computer Vision
Developer Experience
Developer Tools
News & Analysis
Industry Analysis
Pyodide 314.0: Python Browser Packaging Just Got Real
Google Colab CLI: Run A100s and H100s From Your Terminal
MetaMask Agent Wallet: AI Agents Get a Crypto Wallet With Guardrails
Python 3.14 T-Strings: Safer SQL, HTML, and AI Prompts
TensorZero Shuts Down: What OSS LLMOps Can’t Survive
