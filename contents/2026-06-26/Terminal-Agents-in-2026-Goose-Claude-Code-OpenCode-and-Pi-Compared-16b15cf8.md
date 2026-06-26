---
source: "https://outofcontext.dev/blog/goose-claude-code-opencode-pi/"
hn_url: "https://news.ycombinator.com/item?id=48683357"
title: "Terminal Agents in 2026: Goose, Claude Code, OpenCode, and Pi Compared"
article_title: "Terminal Agents in 2026: goose, Claude Code, OpenCode, and Pi Compared"
author: "leianixcheese"
captured_at: "2026-06-26T07:13:53Z"
capture_tool: "hn-digest"
hn_id: 48683357
score: 1
comments: 0
posted_at: "2026-06-26T07:12:24Z"
tags:
  - hacker-news
  - translated
---

# Terminal Agents in 2026: Goose, Claude Code, OpenCode, and Pi Compared

- HN: [48683357](https://news.ycombinator.com/item?id=48683357)
- Source: [outofcontext.dev](https://outofcontext.dev/blog/goose-claude-code-opencode-pi/)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T07:12:24Z

## Translation

タイトル: 2026 年のターミナル エージェント: Goose、Claude Code、OpenCode、Pi の比較
記事のタイトル: 2026 年のターミナル エージェント: グース、クロード コード、OpenCode、Pi の比較
説明: 4 つのターミナル コーディング エージェント (Goose、Claude Code、OpenCode、Pi) をモデル ロックイン、コスト、安全性、拡張性に関して比較しました。実践者選択チャート…

記事本文:
2026 年のターミナル エージェント: goose、Claude Code、OpenCode、Pi の比較 OutOfContext.dev
Stack Overflow で Dzhuneyt を表示 GitHub で Dzhuneyt を表示 2026 年 6 月 25 日 2026 年のターミナル エージェント: goose、Claude Code、OpenCode、Pi の比較
2026 年に端末コーディング エージェントを選択する場合、実際にはモデルを選択しているわけではありません。フロンティア モデルはほぼ統合されているため、それらに巻き付けられるハーネスが日々のエクスペリエンスを決定します。過去 18 か月間で 4 つの製品が真剣に注目されましたが、それらは 1 つの方針に沿って分かれました。Claude Code は独自のプラットフォーム製品です。他の 3 つはオープンソースですが、異なる方向を向いています。goose は財団が管理するジェネラリスト、OpenCode は IDE グレードのコード インテリジェンスを備えたフル機能の候補、Pi は自分で拡張する最小限のコアです。
これはドキュメントに基づいた比較であり、スコア付けされたベイクオフではありません。各ツールは、モデルのロックイン、コストと認証、安全性のデフォルト、ビルドと継承のセットアップの量など、日々の要因を実際に決定する軸に基づいて、主要なソース (公式ドキュメント、リポジトリ、リリース ノート) に対して測定されます。
範囲に関する注: 以下の数値は、2026 年 6 月 25 日時点の一次情報源と照合して検証されています。星の数は毎日変化します。価格、権限モード、サブスクリプション ポリシーはリリースごとに変更されます。
斧はピックを壊す頻度によって並べられます。複数項目のセルは共有機能 (例: MCP · サブエージェントが最初) を先頭にしているため、列をまっすぐ下にスキャンして、どのハーネスが特定の機能を追加または欠落しているかを確認できます。領収書については、 goose 、 Claude Code 、 OpenCode 、または Pi にジャンプします。
goose は、2025 年初めに「コード名 goose」として Block のオープンソース プログラム内で開始されました。2026 年 4 月、Block はそれを Linux Foundation (Anthropic の MCP と同じ傘下) 傘下の Agentic AI Foundation に寄付しました。

仕様と OpenAI の AGENTS.md。リポジトリは github.com/aaif-goose/goose にあります (2026 年 6 月 25 日の時点で最大 50,000 個のスター)。
際立っているのは、グースが「オープンなクロード規範」として位置づけていないことだ。これは、たまたまコーディングが得意な、汎用のローカルファーストエージェントです。レシピ (共有および自動化できるポータブル YAML ワークフロー テンプレート) は、サブスクリプション エージェントから切り替えるときに人々が挙げる機能です。強力な API モデルと BYO キーを組み合わせると、コミュニティ スレッドではタスクあたりのコストが Claude Max 層よりも 1 桁下回ることがよく報告されますが、これはモデルの選択とループの長さに依存します。
Auth: goose は Anthropic の API キーのみです。 Claude Pro/Max サブスクリプション ログインはありません。Anthropic OAuth に対する長年の機能リクエストは出荷されずに終了しました。 ANTHROPIC_API_KEY (または 15 以上の他のプロバイダーのいずれか) を持参し、トークンごとに支払います。
拡張性: MCP は統合サーフェスです。 Block のドキュメントには 70 以上のファーストパーティ拡張パターンがリストされています。コミュニティ MCP カタログははるかに大きいです。サブエージェントと YAML レシピが表面を仕上げます。デスクトップ アプリは、TUI に住んでいないチームメイトの端末の障壁を下げます。
荒削りな点: ハッカー ニュースのスレッドは、レシピとコストの柔軟性を賞賛しますが、オーケストレーションが基礎となるモデルに遅れをとっていると警告します。ハーネスは、同じフロンティア モデルを両方で指している場合でも、Claude Code や OpenCode よりも洗練されていないように感じます。セットアップ (プロファイル、プロバイダー、拡張機能) には、独自の CLI にサインインするよりも手間がかかります。
評決: ベンダー中立のガバナンス、あらゆるモデルの柔軟性、1 回限りのチャット セッションを超えたワークフローのパッケージ化が必要な場合に最適です。最もスムーズなすぐに使えるコーディング UX が必要な場合は最悪です。
Claude Code は、洗練された商用ベンチマークです。 Anthropic のターミナル、VS Code/JetBrains、デスクトップ、Web、および GitHub/GitLab PR ワークフロー

一つの製品表面。すでにクロードにお金を払っていて、摩擦を最小限に抑えたい場合は、これがデフォルトの既存企業に対する対策です。
権限モードは機能チェックリストよりも重要です。ドキュメントには、現在の権限モード ページに対して検証された 6 つがリストされています。default (質問せずに読み取ります。編集とシェルでプロンプトが表示されます)、acceptEdits 、 plan (読み取り専用、編集しないことを提案)、auto (分類子が各アクションをレビューします。重複した詳細説明ではなく、カーソル並列についての自動レビューと YOLO の投稿を参照してください)、dontAsk (ロックダウンされた CI の場合、事前に承認されたツールのみ)、および bypassPermissions (すべて。コンテナ/VM のみ)。マーケティングで無視されているニュアンスが 1 つあります。Shift+Tab サイクルに含まれるのは、デフォルト→編集を受け入れる→プランだけです。 auto はアカウントとバージョンが条件を満たすと表示され (リサーチ プレビュー、v2.1.83 以降)、 dontAsk は --permission-mode 経由で設定され、 bypassPermissions は起動フラグの後ろでのみロックを解除します。
拡張性: MCP、スキル、フック、サブエージェント、および動的ワークフロー (リサーチ プレビュー) はファーストパーティです。エコシステムは人間の形をしており、強力ですが、エージェントが何をすべきかについては、1 つの研究室の意見にとらわれています。
請求とサブスクリプション (2026 年 6 月): インタラクティブな Claude Code セッションはサブスクリプション プールから引き出され、Claude Code は Anthropic が承認したわずか 2 つのサブスクリプション クライアントのうちの 1 つ (もう 1 つは claude.ai) であるため、その OAuth /login は、サードパーティのハーネスとは異なり安定しています。プログラムによる使用 — Agent SDK、claude -p 、GitHub Actions — は、6 月 15 日に別の月次 API クレジットに移行する予定でした。Anthropic は初日にその変更を一時停止しました。タイムラインがソフトなままである間、分割を方向的に true として扱います。ワークフローのほとんどが対話型のターミナル セッションである場合、これは CI でヘッドレス エージェントを実行する場合ほど重要ではありません。
ロックインは見た目よりもソフトです。クロードコードのデフォルト

Anthropic モデルには適用されますが、固定的に組み込まれているわけではありません。 Anthropic Messages API を話す任意のゲートウェイで ANTHROPIC_BASE_URL をポイントし、 ANTHROPIC_AUTH_TOKEN を設定すると、Bedrock、Vertex、LiteLLM または vLLM プロキシ、またはセルフホスト モデルをフロントにすることができます (LLM ゲートウェイのドキュメント)。これにより、トラフィックと請求がサブスクリプションではなくエンドポイントにルーティングされるため、「Anthropic のみ」の上限は壁ではなくデフォルトになります。実際の結合は、厳密にその重みではなく、Anthropic の製品の意見 (モード、スキル、フック) との関係です。
評決: 最高の洗練とクロードのモデルの改善との最も緊密な統合。最悪のすぐに使えるモデルの柔軟性 – ゲートウェイを介して回避できますが、それを魅力的にしているサブスクリプションの経済性を放棄する必要があります。
SST/Anomaly チームの OpenCode は、「なぜ希望するモデルにクロード コードのエルゴノミクスを搭載できないのか?」に対するオープンソースの答えです。リポジトリ anomalyco/opencode には、2026 年 6 月 25 日時点で約 178,000 の GitHub スターが付いていました。スターの速度は、TUI と権限モデルが自分の作業方法に一致するかどうかよりも重要です。
LSP が差別化要因です。 OpenCode は、数十のファイル タイプに対して言語サーバーを起動し、診断をエージェント ループにフィードバックできます (LSP docs)。 LSP はデフォルトでは無効になっています。 OpenCode 自身のドキュメントでは、言語サーバーが非同期になり、メモリを消費し、セッションが遅くなる可能性があると警告しています。これは、コンパイラのフィードバックが生のファイルの差分よりも優れている、TypeScript のような型指定されたコードベースで最も効果的です。小さなスクリプトでは、オーバーヘッドは価値がないかもしれません。代わりに、エージェントに bash 経由で pnpm check を実行させます。
エージェント: タブは、ビルド (完全なツール) と計画 (編集および bash のデフォルトは ask ) を切り替えます。サブエージェント — General、Explore、Scout — は、メイン セッションを汚染することなく並行調査を処理します。権限はツールごとに細かく設定されます ( edit 、 bash 、 lsp 、 MCP ワイルドカード)

）。
モデル、認証、およびプライバシー: OpenCode は、75 を超えるプロバイダー (BYO キー) を使用して Models.dev を経由します。 Claude ユーザーにとっての鋭い点の 1 つは、サブスクリプション認証がなくなったことです。 2026 年 3 月、Anthropic がサードパーティのサブスクリプションの使用を法的に推し進めた後、OpenCode は Anthropic OAuth プラグインを公式ビルドから削除しました。Pro/Max ログインではなく API キーを使用して Claude を認証します。コードの常駐性は、実行するパスによって異なります。BYO キーとセルフホスト パスでは、トラフィックが独自のインフラとプロバイダーに維持されます。ホスト層 (OpenCode Go、月額 $5、Zen) はサーバー経由でルーティングします。実際に使用するモードを確認してください。
トレードオフ: コミュニティのレポートによると、同じモデルのクロード コードよりも反復が遅く感じる可能性があり、ハーネスのオーバーヘッドが増加し、構成面が増加します。貢献者の数は膨大です。ブレイクチェンジは頻繁に土地を変える。
評決: LSP がコストに見合う価値がある場合、実際のコード インテリジェンスを備えたオープンソースのクロード コード人間工学。 MIT ライセンス、任意のモデル、および Anthropic ロックインのない IDE グレードのエージェント モードが必要な場合、サブスクリプションではなくトークンごとに Claude を支払うことができる場合に最適な選択肢です。
Pi は、Mario Zechner の意図的に小さなハーネスです。libGDX と OpenClaw 内のコア ループの背後にあるのと同じエンジニアリングの考え方です。 2026 年 4 月、Zechner は Earendil Works を発表しました。 MIT コアは MIT のままです。リポジトリ Earendil-works/pi は、2026 年 6 月 25 日に約 65,000 個のスターにありました。
インストールに関する注意: Pi は npm パッケージとして出荷されますが、範囲は会社とともに移動しました。現在のパッケージは @earendil-works/pi-coding-agent です。古い @mariozechner/pi-coding-agent スコープは非推奨となり 0.73.1 で凍結され、0.74.0 が新しいホーム (移行後) での最初のリリースとなります。古いピンは引き続き解決されますが、新しいインストールでは Earendil スコープを使用する必要があります。
設計上最小限: Pi にはデフォルトで 4 つのツールが同梱されています — read 、wr

ite 、 edit 、 bash (著者投稿)。組み込みのサブエージェント、プラン モード、または MCP はありません。システム プロンプトは意図的に短くしています。 Zechner 氏の主張は、フロンティア モデルはコーディング用にすでに RL トレーニングされており、追加の命令はほとんどがコンテキストを焼き尽くすというものです。書き込みアクセスをさらに制限したい場合は、オプションの読み取り専用ツール ( grep 、 find 、 ls ) が存在します。
拡張性は製品です: TypeScript 拡張機能のホットリロード。スキルはプログレッシブ開示を使用するため、プロンプトキャッシュは暖かい状態を保ちます。 Pi パッケージは、npm または git 経由で拡張機能を配布します。 50 を超えるサンプルがリポジトリ内に同梱されています。 MCPが欲しいですか？拡張機能を作成するか、Pi をドキュメントに指定してスキャフォールディングさせます。ベンダー マーケットプレイス ゲートはありません。
認証 — 信頼する前にこれをお読みください。Pi は、Claude Pro/Max および ChatGPT サブスクリプションの /login フローを BYO API キーとともに保持します。サブスクリプションパスは実際のものですが、脆弱です。 Anthropic は、2026 年初めにサードパーティのハーネスがクロードのサブスクリプションを利用することを制限し、4 月 4 日にそれを施行し (サードパーティの通話は「プランの制限ではなく、追加使用量からの引き出し」として拒否されました)、その後 6 月 16 日にメーターの変更を一時停止しました。そのため、Anthropic が今年すでに 3 回変更したポリシーに基づいて、サブスクリプション ログインは現在も機能しています。安定したフォールバックとしての API キーの予算。
モード: インタラクティブ TUI、スクリプト用の印刷/JSON、埋め込み用の RPC、独自のエージェントを構築するための SDK。 OpenClaw 内の Pi に関する Armin Ronacher の記事は、パワー ユーザーの訴えを捉えています。拡張コードを読むと、その動作を所有するのはあなたです。
トレードオフ: 大衆向けの洗練されていない。バックグラウンド bash なし — 完全な可観測性を備えた長時間実行プロセスが必要な場合、Pi は tmux を期待します。サブスクリプション /ログイン パスにより API の摩擦は軽減されますが、認証面とポリシーのリスクが増加するため、受け入れる必要があります。
評決: ハーネスの形状を変更し、製品を採用しないでください。あなたがいるときに最適です

エージェント ループをバージョン管理するインフラストラクチャとして扱います。
勝者は 1 人ではありません。モデルに巻き付ける製品の量について 4 つの異なる賭けが行われます。
すでに Anthropic のサブスクリプションを利用しており、セットアップの手間を最小限に抑え、ファーストパーティの洗練、フック、PR 統合を重視する場合は、Claude Code を選択してください。また、ここで Anthropic が保証するのはそのサブスクリプション認証だけであることに注意してください。
オープン ソース (MIT)、任意のモデル、ビルド/プラン エージェントのエルゴノミクス、型指定されたコードベースでの LSP による診断が必要な場合は、OpenCode を選択してください。また、Claude にサブスクリプションではなく API キーを支払うことで問題ありません。
組織にとって基盤ガバナンスが重要である場合、反復可能なチーム間ワークフローのレシピが必要な場合、または 15 を超えるプロバイダーにわたる BYO キーを使用した純粋なコーディングを超えたジェネラリスト エージェントが必要な場合は、グースをお選びください。
エージェント インフラストラクチャ (最小限のコア、所有する TS 拡張機能、トークン効率の高いプロンプト) を構築している場合は、Pi を選択してください。セットアップ時間を 1 回投資すれば、ロックインを永久に節約できます。
まだほとんどの編集を IDE で行う場合は、これらを唯一のツールとして選択しないでください。これらは端末のデフォルトであり、カーソル タブの補完や言語サーバーのインライン修正に代わるものではありません。
チームのデフォルトを移行する前にプレッシャーテストを行う軸の 1 つは、コストと認証の予測可能性です。自律ループを実行するまでは、サブスクリプション エージェントは安く感じられます。 BYO API は、モデル化するまで高価に感じます。

[切り捨てられた]

## Original Extract

Four terminal coding agents — goose, Claude Code, OpenCode, and Pi — compared on model lock-in, cost, safety, and extensibility. A practitioner pick chart…

Terminal Agents in 2026: goose, Claude Code, OpenCode, and Pi Compared OutOfContext.dev
View Dzhuneyt on Stack Overflow View Dzhuneyt on GitHub Jun 25, 2026 Terminal Agents in 2026: goose, Claude Code, OpenCode, and Pi Compared
Pick a terminal coding agent in 2026 and you are not really picking a model — the frontier models have largely converged, so the harness wrapped around them decides the daily experience. Four have earned a serious look in the last eighteen months, and they split along one line: Claude Code is the proprietary platform product; the other three are open source but pull in different directions — goose is the foundation-governed generalist, OpenCode is the full-featured contender with IDE-grade code intelligence, and Pi is the minimal core you extend yourself.
This is a documentation-grounded comparison, not a scored bake-off. Each tool is measured against its primary sources — official docs, repos, and release notes — on the axes that actually decide a daily driver: model lock-in, cost and auth, safety defaults, and how much setup you inherit versus build.
Scope note: Figures below are verified against primary sources as of June 25, 2026. Star counts move daily; pricing, permission modes, and subscription policies change between releases.
Axes are ordered by how often they break a pick. Multi-item cells lead with the shared capabilities (e.g. MCP · subagents first) so you can scan straight down a column and see which harness adds or misses a given feature. Jump to goose , Claude Code , OpenCode , or Pi for receipts.
goose started inside Block’s open-source program as “codename goose” in early 2025. In April 2026 Block donated it to the Agentic AI Foundation under the Linux Foundation — same umbrella as Anthropic’s MCP spec and OpenAI’s AGENTS.md. The repo lives at github.com/aaif-goose/goose (~50K stars as of June 25, 2026).
What stands out: goose is not positioning as “the open Claude Code.” It is a general-purpose, local-first agent that happens to be good at coding. Recipes — portable YAML workflow templates you can share and automate — are the feature people cite when they switch from a subscription agent. Pair a strong API model with BYO keys and community threads often report per-task cost an order of magnitude below Claude Max tiers, though that depends on model choice and loop length.
Auth: goose is API-key-only for Anthropic. There is no Claude Pro/Max subscription login — a long-standing feature request for Anthropic OAuth was closed without shipping. You bring an ANTHROPIC_API_KEY (or any of 15+ other providers) and pay per token.
Extensibility: MCP is the integration surface. Block’s docs list 70+ first-party extension patterns; the community MCP catalog is much larger. Subagents and YAML recipes round out the surface. The desktop app lowers the terminal barrier for teammates who will not live in a TUI.
Rough edges: Hacker News threads praise recipes and cost flexibility but flag orchestration that lags the underlying model — the harness feels less refined than Claude Code or OpenCode even when you point the same frontier model at both. Setup (profiles, providers, extensions) takes more fiddling than signing into a proprietary CLI.
Verdict: Best when you want vendor-neutral governance, any-model flexibility, and workflow packaging beyond one-off chat sessions. Worst when you want the smoothest out-of-box coding UX.
Claude Code is the polished commercial benchmark. Anthropic ships terminal, VS Code/JetBrains, desktop, web, and GitHub/GitLab PR workflows from one product surface. If you are already paying for Claude and want the least friction, this is the default incumbents measure against.
Permission modes matter more than feature checklists. The docs list six, verified against the current permission-modes page : default (reads without asking; prompts on edits and shell), acceptEdits , plan (read-only, propose don’t edit), auto (a classifier reviews each action — see our auto-review vs YOLO post for the Cursor parallel, not a duplicate deep-dive), dontAsk (only pre-approved tools, for locked-down CI), and bypassPermissions (everything; containers/VMs only). One nuance the marketing skips: only default → acceptEdits → plan sit in the Shift+Tab cycle. auto appears once your account and version qualify (research preview, v2.1.83+), dontAsk is set via --permission-mode , and bypassPermissions unlocks only behind a launch flag.
Extensibility: MCP, skills, hooks, subagents, and dynamic workflows (research preview) are first-party. The ecosystem is Anthropic-shaped — powerful, but you are inside one lab’s opinion of what an agent should do.
Billing and subscription (June 2026): Interactive Claude Code sessions draw from your subscription pool, and Claude Code is one of only two Anthropic-approved subscription clients (the other is claude.ai), so its OAuth /login is stable in a way third-party harnesses’ is not. Programmatic usage — Agent SDK, claude -p , GitHub Actions — was scheduled to move to a separate monthly API credit on June 15. Anthropic paused that change on day one ; treat the split as directionally true while the timeline stays soft. If your workflow is mostly interactive terminal sessions, this matters less than if you run headless agents in CI.
Lock-in is softer than it looks. Claude Code defaults to Anthropic models, but it is not hard-wired to them. Point ANTHROPIC_BASE_URL at any gateway that speaks the Anthropic Messages API and set ANTHROPIC_AUTH_TOKEN , and you can front Bedrock, Vertex, a LiteLLM or vLLM proxy, or a self-hosted model ( LLM gateway docs ). That routes traffic — and billing — to your endpoint instead of your subscription, so the “Anthropic-only” ceiling is a default, not a wall. The real coupling is to Anthropic’s product opinions (modes, skills, hooks), not strictly its weights.
Verdict: Best polish and tightest integration with Claude’s model improvements. Worst out-of-box model flexibility — escapable through a gateway, but only by giving up the subscription economics that make it attractive in the first place.
OpenCode from the SST/Anomaly team is the open-source answer to “why can’t I have Claude Code ergonomics on whatever model I want?” The repo anomalyco/opencode had ~178K GitHub stars on June 25, 2026 — star velocity matters less than whether the TUI and permission model match how you work.
LSP is the differentiator. OpenCode can start language servers for dozens of file types and feed diagnostics back into the agent loop ( LSP docs ). LSP is disabled by default; OpenCode’s own docs warn that language servers can desync, eat memory, and slow sessions. It pays off most on typed codebases like TypeScript, where compiler feedback beats raw file diffs; on small scripts the overhead may not be worth it — let the agent run pnpm check via bash instead.
Agents: Tab switches between Build (full tools) and Plan (edits and bash default to ask ). Subagents — General, Explore, Scout — handle parallel research without polluting the main session. Permissions are granular per tool ( edit , bash , lsp , MCP wildcards).
Models, auth, and privacy: OpenCode routes through Models.dev with 75+ providers (BYO key). One sharp edge for Claude users: subscription auth is gone. In March 2026, after Anthropic’s legal push against third-party subscription use, OpenCode stripped its Anthropic OAuth plugin from the official build — you authenticate Claude with an API key, not a Pro/Max login. Code residency depends on the path you run: BYO-key and self-hosted paths keep traffic on your own infra and provider; hosted tiers (OpenCode Go at $5/mo, Zen) route through their servers — check the mode you actually use.
Tradeoffs: Community reports suggest iteration can feel slower than Claude Code on the same model — more harness overhead, more configuration surface. The contributor count is enormous; breaking changes land frequently.
Verdict: Open-source Claude Code ergonomics with real code intelligence when LSP is worth the cost. The standout choice when you want MIT license, any model, and IDE-grade agent modes without Anthropic lock-in — provided you are fine paying Claude per-token rather than on a subscription.
Pi is Mario Zechner’s deliberately tiny harness — the same engineering mind behind libGDX, and the core loop inside OpenClaw. In April 2026 Zechner announced Earendil Works ; the MIT core stays MIT. The repo earendil-works/pi sat at ~65K stars on June 25, 2026.
Install note: Pi ships as an npm package, but the scope moved with the company. The current package is @earendil-works/pi-coding-agent ; the old @mariozechner/pi-coding-agent scope is deprecated and frozen at 0.73.1 , with 0.74.0 the first release under the new home ( migration post ). Old pins still resolve, but new installs should use the Earendil scope.
Minimal by design: Pi ships four tools by default — read , write , edit , bash ( author post ). No built-in subagents, plan mode, or MCP. The system prompt is intentionally short; Zechner’s argument is that frontier models are already RL-trained for coding — extra instructions mostly burn context. Optional read-only tools ( grep , find , ls ) exist if you want to restrict write access further.
Extensibility is the product: TypeScript extensions hot-reload. Skills use progressive disclosure so prompt cache stays warm. Pi packages distribute extensions via npm or git. Fifty-plus examples ship in-repo. Want MCP? Write an extension or point Pi at the docs and let it scaffold one — there is no vendor marketplace gate.
Auth — read this before you rely on it: Pi keeps a /login flow for Claude Pro/Max and ChatGPT subscriptions alongside BYO API keys. The subscription path is real but fragile. Anthropic restricted third-party harnesses from drawing on Claude subscriptions in early 2026, enforced it on April 4 (third-party calls were rejected with “draws from your extra usage, not your plan limits”), then paused the metering change on June 16 . So subscription login works today, under a policy Anthropic has already changed three times this year. Budget for an API key as the stable fallback.
Modes: Interactive TUI, print/JSON for scripts, RPC for embedding, SDK for building your own agent. Armin Ronacher’s write-up on Pi inside OpenClaw captures the power-user appeal: you read the extension code, you own the behavior.
Tradeoffs: Not mass-market polish. No background bash — Pi expects tmux if you want long-running processes with full observability. The subscription /login paths reduce API friction but add auth surface — and policy risk — you must accept.
Verdict: Reshape the harness, do not adopt a product. Best when you treat the agent loop as infrastructure you version-control.
No single winner — four different bets on how much product you want wrapped around the model.
Pick Claude Code if you already live in Anthropic’s subscription, want the least setup friction, and value first-party polish, hooks, and PR integrations — and note its subscription auth is the only one here Anthropic guarantees.
Pick OpenCode if you want open source (MIT), any model, Build/Plan agent ergonomics, and LSP-fed diagnostics on typed codebases — and you are fine paying Claude with an API key, not a subscription.
Pick goose if foundation governance matters to your org, you need recipes for repeatable cross-team workflows, or you want a generalist agent beyond pure coding with BYO keys across 15+ providers.
Pick Pi if you are building agent infrastructure — minimal core, TS extensions you own, token-efficient prompts — and you will invest setup time once to save lock-in forever.
Pick none of them as your only tool if you still do most editing in an IDE. These are terminal defaults, not replacements for Cursor tab-complete or a language server’s inline fixes.
One axis to pressure-test before migrating a team default: cost and auth predictability. Subscription agents feel cheap until you run autonomous loops; BYO API feels expensive until you model per-

[truncated]
