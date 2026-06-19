---
source: "https://www.unsiloed.ai/blog/guides/claude-code-codex-one-pipeline"
hn_url: "https://news.ycombinator.com/item?id=48595599"
title: "Claude Code and Codex as one pipeline"
article_title: "Claude Code + Codex as One Pipeline: Claude Code + Codex as One Pipeline: A Technical Guide to Running Both Instead of Choosing | Unsiloed AI"
author: "ritzaco"
captured_at: "2026-06-19T07:59:57Z"
capture_tool: "hn-digest"
hn_id: 48595599
score: 1
comments: 0
posted_at: "2026-06-19T06:51:56Z"
tags:
  - hacker-news
  - translated
---

# Claude Code and Codex as one pipeline

- HN: [48595599](https://news.ycombinator.com/item?id=48595599)
- Source: [www.unsiloed.ai](https://www.unsiloed.ai/blog/guides/claude-code-codex-one-pipeline)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T06:51:56Z

## Translation

タイトル: 1 つのパイプラインとしての Claude Code と Codex
記事のタイトル: Claude Code + Codex を 1 つのパイプラインとして: Claude Code + Codex を 1 つのパイプラインとして: 選択する代わりに両方を実行するためのテクニカル ガイド |サイロ化されていない AI
説明: ベンチマーク、コンテキスト ウィンドウの動作、トークン エコノミクス、Claude Code と OpenAI Codex を単一のコーディング パイプラインとして実行するための MCP 配線。

記事本文:
Claude Code + Codex を 1 つのパイプラインとして: Claude Code + Codex を 1 つのパイプラインとして: 選択する代わりに両方を実行するためのテクニカル ガイド |サイロ化されていない AI
[ NEW ] Unsiloed が olmOCR-Bench で #1 ランクを達成 → Unsiloed Playground ドキュメント 価格設定 リソース 採用情報 ログイン デモを予約 ← ブログに戻る エンジニアリング クロード コード + コーデックスを 1 つのパイプラインとして
Claude Code + Codex を 1 つのパイプラインとして: どちらかを選択するのではなく両方を実行するためのテクニカル ガイド
ベンチマーク、コンテキスト ウィンドウの動作、トークン エコノミクス、Claude Code と OpenAI Codex を単一のコーディング パイプラインとして実行するための MCP 配線。
現在、AI コーディング コミュニティで最もよくある質問は、「クロード コードですか、それともコーデックスですか?」というものです。 40k 行の Rust サービスと 12k 行の React フロントエンドの両方を 2 か月間実行した後、これは間違った質問だと思います。これらのツールは相反する設計哲学に基づいて構築されており、その相反するからこそ、別々に使用するよりも連携して機能させることができるのです。
この記事では、ベンチマークが実際に何を示しているか、コンテキスト ウィンドウがいっぱいになるときに各ツールがどのように動作するか、実際のコストを決定するトークン エコノミクス、そして最も重要なことに、それらを単一のパイプラインとして実行するための具体的な MCP 配線について説明します。ここにあるものはすべて、現在のドキュメントに対して検証可能です。バージョン番号は急速に変化するため、実装する際には最新のリリースと照らし合わせて確認してください。
ローカル対クラウドのメンタル モデルの使用をやめる
時代遅れの枠組みは、Claude Code がローカルの端末ツールであり、Codex がクラウドのツールであるということです。その区別は崩れました。 Anthropic は現在、ターミナル、IDE、デスクトップ、Slack、および Web サーフェスにわたって Claude Code を配布しています。 OpenAI は、アプリ、IDE、CLI、クラウド全体に Codex を提供します。どちらもローカル実行と非同期実行にまたがります。
現在も維持されている区別は、監視型と自律型です。
Claude Code はライブで操作できるように設計されています。あなたは計画を見直し、それを観察します

推論し、編集が行われるたびに承認します。
Codex は委任用に設計されています。スコープ指定されたタスクを渡すと、サンドボックス内で動作し、後で結果を確認します。
これは機能のギャップではありません。これは意図するワークフローの違いであり、どのツールがパイプラインのどのステージを所有するかが決まります。
2026 年半ばの同じ時間枠に合わせて:
パターンは一貫しています。Codex はターミナルとシェルの作業に強いです。クロードは、複数ファイルの深い推論に優れています。これは、上記の監視付きと自律型の区別に直接対応します。
見落としがちな方法論上の注意点が 1 つあります。それは、各ツールのモデルがほぼ数週間ごとに変更されることです。 OpenAI は数か月かけて GPT-5.3、5.4、および 5.5-Codex に移行しました。 Anthropic は同じ期間内で Opus 4.6、4.7、および 4.8 を移行し、Sonnet 4.6 を標準価格の 1M トークンのコンテキストに移行しました。どのベンチマークも、移動するターゲットのスナップショットです。数値を方向性のあるものとして扱い、それに依存する前に再検証してください。
コンテキスト ウィンドウの動作: 「私の指示が無視された」ことを説明する詳細
1M トークンのコンテキスト ウィンドウは、そのウィンドウ全体で均一な品質を意味するものではありません。ウィンドウがいっぱいになると、取得の信頼性が低下します。広く引用されている GitHub の問題では、この曲線について文書化されています。0 ～ 20% の範囲では信頼性の高いパフォーマンスがあり、それを超えると徐々に低下し、100 万トークン付近では取得のおよそ 4 件に 1 件が失敗します。有効な信頼できる範囲は 200 ～ 256,000 トークンに近くなります。
これは、エージェントが長いセッションの途中で「コーディング ガイドラインに従わなくなる」というよくある苦情の説明になります。命令は無視されているわけではありません。飽和したコンテキストの奥深くから命令を取得するのは困難になってきています。実際的な緩和策:
タスクを切り替えるときにコンテキストをリセットするには、/clear を使用します。
/init を使用して CLAUDE.md からプロジェクト メモリを再構築します。
個々のセッションをうまく維持する

指示の順守が重要な場合は、最大値を下回ります。
関連するメモ: 2026 年初頭の一定期間、Ultrathink / 「より深く考える」トリガーは表面的なものになりました。Anthropic エンジニアの公開確認によると、これらは依然として視覚効果をレンダリングしますが、推論の深さは増加しません。それらに依存している場合は、代わりにプラン モードをお勧めします。
トークン経済学が現実世界のコストを決定する
サブスクリプション価格は重要な指標ではありません。指標は、1 日に取得するエージェント セッションの数と、それらのセッションをどれくらいの速さで消費するかです。これを推進するのは 2 つの事実です。
同一のタスクにおいて、Claude Code は Codex の約 4 倍のトークンを使用して測定されました。より深い推論にはコストがかかります。
マルチエージェントのワークフローにより、消費量が増加します。 Claude Code のエージェント チームは、プラン モードで単一セッションの約 7 倍のトークンを実行します。 Codex では、サブエージェントの上限を開発者あたり 8 に設定しています。クロードのエージェント チームにはハードキャップはありませんが、生成されるエージェントの数に応じて消費量が調整されます。
実際の結果は、開発者フィードバックの大規模なサンプルにわたって一貫して報告されています。20 ドル層では、単一の複雑なプロンプトがクロード コードの使用時間帯の大部分を消費する可能性があるのに対し、同等層の Codex は終日使用を維持します。広く繰り返されている要約は、Claude Code の方が品質は高いものの制限による制約があり、Codex の方が品質はわずかに低いものの、より継続的に使用できるというものです。
この経済的な非対称性自体が、ワークフローを分割する根拠となっています。つまり、大量の作業をより安価で高速なツールにルーティングし、コストに見合った作業用に高価なツールを予約しておきます。
統合層はモデル コンテキスト プロトコルです。 Claude Code は MCP クライアントであり、Codex CLI は MCP サーバーとして動作できます。つまり、端末を離れることなくタスクを一方から他方にルーティングできます。複雑さの高い順に 3 つのパターン。
パット

ern 1: コミット時のモデル間レビュー
最も収益が高く、労力が最も少ないパターン。 Claude Code が計画と実装を作成します。コミットする前に、独立したレビューのために差分を Codex に送信し、フィードバックを組み込みます。レビュー モデルは最初のモデルのアプローチに注力していないため、実際に修正されたバグではなく合格するように修正されたテストなど、1 人の自己レビュー エージェントが対応する問題を確実に捕捉します。
Codex を MCP サーバーとして登録します。
クロード mcp add --scope user codex-subagent \ --transport stdio -- uvx codex-as-mcp@latest
次に、グローバル CLAUDE.md でポリシーをエンコードします。
# ポリシーのレビュー コミットの前に、独立したレビューのためにステージングされた差分をコーデックス MCP サーバーに送信します。反対意見をインラインで表面化し、「git commit」を実行する前に解決します。複数ファイルの変更に対して独自の実装を自動的に受け入れないでください。
パターン2：強さで分ける
端末を多用する作業やサンドボックスでの並列ファーストパス実装には Codex を使用してください。高速でベンチマークが優れています。結果を Claude Code に取り込み、詳細なリファクタリング、セキュリティ レビュー、マルチエージェントによる調整レビューを行い、事前にベンチマークを行います。これはレースではなく組み立てラインです。各段階は、その段階に最適なツールによって処理されます。
パターン 3: 統合されたマルチエージェント
より大きなタスクの場合、Codex カスタム エージェントは AGENTS.md ファイルから共有規約を読み取ります。 OpenAI 独自のガイダンスでは、小規模で高速なモデルを大量のサブエージェントに固定し、判断が最も重要なエージェントのためにフラッグシップを予約します。一般的なパターンでは、プル リクエストのレビューが 3 つの並列エージェントに分割されます。1 つはコードをマッピングし、1 つはコードをレビューし、もう 1 つは外部 API をライブ ドキュメントと照合して検証します。
クロード側では、2 つのメカニズムの違いを理解してください。
サブエージェントは単一セッション内でのみ実行されます。

y 親エージェントに報告します。
エージェント チーム (実験版、Opus 4.6 に同梱) は、メールボックス システムを介してピアツーピアで通信し、共有タスク リストを介して調整する永続的な独立したインスタンスです。
エージェント チームは、 settings.json のフラグの背後で有効になります。
{ "環境": { "CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS": "1" } }
より広範なエコシステムは、ハーネス中立な構成に向けて移行しています。現在、マーケットプレイスは、Claude Code、Codex、Cursor、Gemini CLI、および Copilot によってネイティブに使用されるエージェント、スキル、およびコマンドの単一の Markdown ソースを公開しており、Codex または Gemini が移植せずに既存の .claude/agents/ 定義を再利用できるようにするツールが存在します。単一のハーネスを中心にワークフローを構築することは、ツールがどこに向かっているのかを賭けたものです。
いくつかの問題は発生しやすく、文書化されることはほとんどありません。
サイズが大きすぎる命令ファイルは、サイレントにデグレードされます。約 500 行を超えると、命令に従う能力が有限であるため、構成ファイルの大部分が従わなくなります。集中した 50 行のファイルは、広大な 1,000 行のファイルよりもパフォーマンスが優れています。 Codex CLI はさらに進んで、 project_doc_max_bytes を超えるコンテンツを暗黙的に切り捨てるため、サイズが大きすぎるファイルは警告なしに積極的にコンテキストを失います。信頼できる情報源を 1 つ AGENTS.md に保持し、ルールを複製するのではなく、ツール固有のファイルからそれを参照させます。
自動生成された構成は避けてください。 /init または自動ジェネレーターによって生成されるファイルは、汎用的で肥大化する傾向があります。構成を手動で記述して、各行が実際に観察された問題に対処できるようにします。リンターまたはフォーマッタがすでに処理しているルールを含めないでください。
MCP ツールのコンテキストを管理します。多数の MCP サーバーが構成されている場合、ツール定義が大量のコンテキストを消費する可能性があります。デフォルトで有効になっているツール検索は、必要になるまでツール スキーマを延期するため、セッション開始時に名前のみが読み込まれます。 ENABLE_TOOL_SEARCH= を設定すると、スキーマが事前に自動ロードされます。

n それらはコンテキスト ウィンドウの一部に収まり、残りは延期されます。
プラットフォームの不安定性を考慮します。どちらのツールも、リクエストの重要なシェアに影響を与えるルーティングのバグや、導入されて数日以内にロールバックされた少なくとも 1 つの確認されたハーネスの回帰など、インフラストラクチャ インシデントを繰り返し発生しました。出力品質が突然低下した場合は、問題が構成にあると考える前に、プラットフォームのステータスを確認してください。
作業が端末やインフラストラクチャに負荷がかかる場合、エントリー層で十分な制限を設けたサブエージェントの並列処理が必要な場合、またはすでに ChatGPT の料金を支払っており、エージェント コーディングを評価するための摩擦の少ない方法が必要な場合は、Codex を単独で使用してください。
複雑な複数ファイルの問題に対して最高の精度が必要な場合、真のピアツーピア調整を備えたエージェント チームが必要な場合、および上位層が予算内にあるため、使用制限が制約にならない場合は、クロード コードを単独で使用します。
確実に不正確な変更を行うとコストがかかる、本番環境に不可欠な作業を出荷する場合は、両方を使用してください。 Codex を介して高速な最初のパスを実行し、Claude Code を介して詳細なレビューとリファクタリングを実行し、コミット時にクロスモデル レビューを実行します。これにより、結果が変わる場合にのみ高価なトークンが消費されます。
実稼働ソフトウェアを出荷しているほとんどのチームにとって、3 番目のオプションが最も防御可能です。
これは、2 つのコードベース、1 人の開発者、特定のプロンプト スタイル、および「完了」の特定の定義を反映しています。グリーンフィールドでの単独作業では、チームで大規模なシステムを維持する場合とは異なる方法でトレードオフを考慮します。引用されているすべてのバージョン番号の有効期間は短いです。ベンチマークは指向性プロキシであり、独自のリポジトリでのテストの代替ではありません。導入とトークンの数値は、アナリストの推定、ベンダーのレポート、および部分的な可視性を備えた可観測性ツールに基づいています。耐荷重のあるものはそれに依存する前に確認してください。
「クロード・コード対コーデックス」

カテゴリエラーであるため、解決できません。 1 つのツールは監視された深さ向けに構築され、もう 1 つは自律的な委任向けに構築されており、その対立こそが、関係を破るのではなく、うまく構成する理由です。ベンチマークは、タスクの種類ごとにきれいに分割されていることを示しています。トークンエコノミクスによれば、どちらかのツールですべてを実行させるよりも、分割ワークフローの方がコストが低くなります。そして、統合ツール (MCP ブリッジ、クロスモデル レビュー、ハーネス中立エージェント定義) は、結合されたパイプラインを例外ではなくデフォルトにするために構築されています。
チームにとってより有益な質問は、どのツールを標準化するかではなく、パイプラインがどのようなもので、各ツールがどのステージを所有するかということです。それに答えると、選択肢は二者択一ではなくなります。
Unsiloed がドキュメントから何を抽出するかを確認する
サンプル PDF をアップロードすると、数分で構造化された出力が得られます。統合は必要ありません。
マルチモーダルな非構造化データを、高精度で構造化された LLM 対応資産に変換します。

## Original Extract

Benchmarks, context-window behavior, token economics, and the MCP wiring for running Claude Code and OpenAI Codex as a single coding pipeline.

Claude Code + Codex as One Pipeline: Claude Code + Codex as One Pipeline: A Technical Guide to Running Both Instead of Choosing | Unsiloed AI
[ NEW ] Unsiloed Achieves #1 Rank on olmOCR-Bench → Unsiloed Playground Docs Pricing Resources Careers Log in Book a demo ← Back to Blog Engineering Claude Code + Codex as One Pipeline
Claude Code + Codex as One Pipeline: A Technical Guide to Running Both Instead of Choosing
Benchmarks, context-window behavior, token economics, and the MCP wiring for running Claude Code and OpenAI Codex as a single coding pipeline.
The most common question in AI coding communities right now is "Claude Code or Codex?" After running both on a 40k-line Rust service and a 12k-line React frontend over two months, I think it is the wrong question. The tools are built on opposite design philosophies, and that opposition is precisely why they work better together than apart.
This article covers what the benchmarks actually say, how each tool behaves as its context window fills, the token economics that determine real-world cost, and most importantly, the concrete MCP wiring to run them as a single pipeline. Everything here is verifiable against current documentation; version numbers move quickly, so confirm them against the latest releases when you implement.
Stop using the local-vs-cloud mental model
The outdated framing is that Claude Code is the local terminal tool and Codex is the cloud one. That distinction has collapsed. Anthropic now ships Claude Code across terminal, IDE, desktop, Slack, and web surfaces; OpenAI ships Codex across app, IDE, CLI, and cloud. Both span local and async execution.
The distinction that still holds is supervised vs. autonomous :
Claude Code is designed to be steered live. You review the plan, observe the reasoning, and approve edits as they happen.
Codex is designed for delegation. You hand it a scoped task, it works in a sandbox, and you review the result later.
This is not a feature gap. It is a difference in intended workflow, and it determines which tool should own which stage of your pipeline.
Aligned to the same time window in mid-2026:
The pattern is consistent: Codex is stronger on terminal and shell work; Claude is stronger on deep multi-file reasoning. This maps directly onto the supervised-vs-autonomous distinction above.
One methodological caveat that is easy to miss: the model under each tool changes almost every few weeks. OpenAI moved through GPT-5.3, 5.4, and 5.5-Codex in months; Anthropic moved through Opus 4.6, 4.7, and 4.8 in the same window and shifted Sonnet 4.6 to a 1M-token context at standard pricing. Any benchmark is a snapshot of a moving target. Treat the numbers as directional and re-verify before relying on them.
Context-window behavior: the detail that explains "it ignored my instructions"
A 1M-token context window does not mean uniform quality across that window. Retrieval reliability degrades as the window fills. A widely cited GitHub issue documented the curve: reliable performance in the 0–20% range, progressive degradation beyond that, and roughly 1 in 4 retrievals failing near 1M tokens . The effective reliable range is closer to 200–256K tokens.
This explains the common complaint that the agent "stops following my coding guidelines" partway through a long session. The instructions are not being ignored — they are becoming hard to retrieve from deep in a saturated context. Practical mitigations:
Use /clear to reset context when switching tasks.
Use /init to rebuild project memory from CLAUDE.md .
Keep individual sessions well under the maximum if instruction adherence matters.
A related note: for a period in early 2026, the ultrathink / "think harder" triggers became cosmetic — they still render the visual effect but no longer increase reasoning depth, per an Anthropic engineer's public confirmation. If you have been relying on them, prefer plan mode instead.
Token economics determine real-world cost
Subscription price is not the metric that matters. The metric is how many agent sessions you get per day and how quickly you consume them. Two facts drive this:
On identical tasks, Claude Code has been measured using roughly 4x the tokens of Codex . Deeper reasoning has a cost.
Multi-agent workflows multiply consumption. Claude Code's Agent Teams run approximately 7x the tokens of a single session in plan mode. Codex caps subagents at 8 per developer; Claude's Agent Teams have no hard cap but scale consumption with the number of agents spawned.
The practical consequence, reported consistently across large samples of developer feedback: at the $20 tier , a single complex prompt can consume a large fraction of a Claude Code usage window, while Codex at the equivalent tier sustains all-day use. The widely repeated summary is that Claude Code is higher quality but constrained by limits, while Codex is slightly lower quality but more continuously usable.
This economic asymmetry is itself the argument for a split workflow: route high-volume work to the cheaper, faster tool and reserve the expensive tool for work that justifies the cost.
The integration layer is the Model Context Protocol. Claude Code is an MCP client, and Codex CLI can operate as an MCP server, which means you can route tasks from one to the other without leaving your terminal. Three patterns, in increasing order of complexity.
Pattern 1: Cross-model review on commit
The highest-return, lowest-effort pattern. Claude Code writes the plan and implementation; before committing, it sends the diff to Codex for an independent review and incorporates the feedback. Because the reviewing model is not invested in the first model's approach, it reliably catches issues a single self-reviewing agent waves through, including tests modified to pass rather than bugs actually fixed.
Register Codex as an MCP server:
claude mcp add --scope user codex-subagent \ --transport stdio -- uvx codex-as-mcp@latest
Then encode the policy in your global CLAUDE.md :
# Review policy Before any commit, send the staged diff to the codex MCP server for an independent review. Surface its objections inline and resolve them before running `git commit`. Do not auto-accept your own implementation on multi-file changes.
Pattern 2: Split by strength
Use Codex for terminal-heavy work and parallel first-pass implementation in its sandbox, where it is fast and benchmarks ahead. Bring the result into Claude Code for deep refactoring, security review, and coordinated multi-agent review, where it benchmarks ahead. This is an assembly line, not a race — each stage is handled by the tool best suited to it.
Pattern 3: Orchestrated multi-agent
For larger tasks, Codex custom agents read shared conventions from an AGENTS.md file. OpenAI's own guidance is to pin a small, fast model to high-volume subagents and reserve the flagship for the agent whose judgment matters most. A common pattern splits a pull-request review across three parallel agents: one maps the code, one reviews it, one verifies external APIs against live documentation.
On the Claude side, understand the difference between two mechanisms:
Subagents run within a single session and only report back to the parent agent.
Agent Teams (experimental, shipped with Opus 4.6) are persistent, independent instances that communicate peer-to-peer through a mailbox system and coordinate via a shared task list.
Agent Teams are enabled behind a flag in settings.json :
{ "env": { "CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS": "1" } }
The broader ecosystem is moving toward harness-neutral configuration: marketplaces now publish a single Markdown source of agents, skills, and commands consumed natively by Claude Code, Codex, Cursor, Gemini CLI, and Copilot, and tools exist to let Codex or Gemini reuse existing .claude/agents/ definitions without porting. Building a workflow around a single harness is a bet against where the tooling is heading.
Several issues are easy to hit and rarely documented:
Oversized instruction files are silently degraded. Past roughly 500 lines, much of a config file stops being followed because instruction-following capacity is finite; a focused 50-line file outperforms a sprawling 1,000-line one. Codex CLI goes further and silently truncates content past project_doc_max_bytes , so an oversized file actively loses context with no warning. Keep one source of truth in AGENTS.md and have tool-specific files reference it rather than duplicating rules.
Avoid auto-generated config. Files produced by /init or auto-generators tend to be generic and bloated. Write configuration by hand so every line addresses a real, observed problem. Do not include rules a linter or formatter already handles.
Manage MCP tool context. With many MCP servers configured, tool definitions can consume significant context. Tool Search, enabled by default, defers tool schemas until needed so only names load at session start. Setting ENABLE_TOOL_SEARCH=auto loads schemas upfront when they fit within a fraction of the context window and defers the rest.
Account for platform instability. Both tools have experienced repeated infrastructure incidents, including routing bugs affecting a meaningful share of requests and at least one confirmed harness regression that was introduced and rolled back within days. When output quality drops suddenly, check platform status before assuming the problem is your configuration.
Use Codex alone if your work is terminal- and infrastructure-heavy, you want subagent parallelism with generous limits at the entry tier, or you already pay for ChatGPT and want a low-friction way to evaluate agentic coding.
Use Claude Code alone if you need top accuracy on complex multi-file problems, you want Agent Teams with genuine peer-to-peer coordination, and a higher tier is within budget so usage limits stop being the constraint.
Use both if you ship production-critical work where a confidently incorrect change is costly. Run the fast first pass through Codex, the deep review and refactor through Claude Code, and cross-model review on commit. This spends expensive tokens only where they change the outcome.
For most teams shipping production software, the third option is the most defensible.
This reflects two codebases, one developer, a particular prompting style, and a specific definition of "done." Greenfield solo work weighs the tradeoffs differently than maintaining a large system on a team. Every version number cited has a short shelf life. Benchmarks are directional proxies, not a substitute for testing on your own repository. Adoption and token figures originate from analyst estimates, vendor reporting, and observability tools with partial visibility. Verify anything load-bearing before depending on it.
"Claude Code vs. Codex" resists resolution because it is a category error. One tool is built for supervised depth, the other for autonomous delegation, and that opposition is the reason they compose well rather than a tie to be broken. The benchmarks show they split cleanly by task type. The token economics show a split workflow costs less than forcing either tool to do everything. And the integration tooling — MCP bridges, cross-model review, harness-neutral agent definitions — is being built to make the combined pipeline the default rather than the exception.
The more useful question for your team is not which tool to standardize on, but what your pipeline looks like and which stage each tool owns. Answer that, and the choice stops being a binary.
See what Unsiloed extracts from your documents
Upload a sample PDF and get structured output in minutes. No integration required.
Transform multimodal unstructured data into structured, LLM-ready assets with high accuracy.
