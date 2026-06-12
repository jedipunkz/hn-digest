---
source: "https://labs.beconfident.app/papers/harness-engineering-survey"
hn_url: "https://news.ycombinator.com/item?id=48508618"
title: "The 98% Problem: A Survey of Harness Engineering for AI Agents"
article_title: "The 98% Problem: A Survey of Harness Engineering for AI Agents — BeConfident Labs"
author: "gdss"
captured_at: "2026-06-12T21:38:26Z"
capture_tool: "hn-digest"
hn_id: 48508618
score: 4
comments: 0
posted_at: "2026-06-12T19:44:11Z"
tags:
  - hacker-news
  - translated
---

# The 98% Problem: A Survey of Harness Engineering for AI Agents

- HN: [48508618](https://news.ycombinator.com/item?id=48508618)
- Source: [labs.beconfident.app](https://labs.beconfident.app/papers/harness-engineering-survey)
- Score: 4
- Comments: 0
- Posted: 2026-06-12T19:44:11Z

## Translation

タイトル: 98% の問題: AI エージェントのためのハーネス エンジニアリングに関する調査
記事のタイトル: 98% の問題: AI エージェントのためのハーネス エンジニアリングに関する調査 — BeConfident Labs
説明: エージェントの品質は、モデルの下位、つまりループ、コンテキスト エンジン、ツール サーフェス、安全スタック、および評価器内に存在します。 2026 年中頃のハーネス エンジニアリングの調査。オープンソースの自己保守型ナレッジ ハーネスである GROOM を紹介します。

記事本文:
98% の問題: AI エージェントのためのハーネス エンジニアリングの調査 — BeConfident Labs 98% の問題: AI エージェントのためのハーネス エンジニアリングの調査
(a) トリガー シーケンス (b) メンテナンス操作エージェント セッション ランチャー GROOM エージェント wiki/ チェック: 有効?最後の実行は 24 時間以上前ですか?スキルのアクティベーションは 100 ミリ秒未満で返されます。スポーン · デタッチされたすべてのページの読み取りページ コンテンツの編集 · 1 つの操作でジャーナル エントリを追加 t 会話が進行し、lint がブロックされることはありません。フォームを修正します。意味には決して触れないでください。 prune 繰り返しているものを切り取る。残す行数を減らします。展開 変更内容を取り込みます。最大 3 ～ 6 個のファイルをタッチします。研究 新しい論文を認めます。まず引用を要求します。サイクルごとに 1 つの演算、構成作者によって選択
以下の著者の寄稿文。
フロンティア モデルは 2023 年から 2026 年の間に収束しました。ほとんどの運用タスクでは、あるトップ モデル ファミリを別のモデル ファミリに交換しても、結果はあまり変わりません。勝てるシステムは、モデルの呼び出し方法、モデルに供給するコンテキスト、公開するツール、実行が行われる場所、結果の測定方法など、レイヤーの下で異なります。専門家はその層をハーネスと呼びます。
最初の番号がこの文書のタイトルになります。最も文書化されたプロダクションエージェントである Claude Code のコミュニティ分析では、コードの約 1.6% がモデルの動作を決定していると推定されています。残りの部分はコンテキストを組み立て、ツールをディスパッチし、権限をチェックし、サンドボックスを実行し、状態を保持し、障害から回復します [17] 。コード行は、機能が存在する場所ではなく、エンジニアリングが存在する場所を測定するため、比率を桁違いの主張として扱います。それは依然として実際の状態、つまり 98% 問題に名前を付けています。エージェントの品質を決定する層は誰もベンチマークを行わず、チームのスタッフもほとんどおらず、すべてのプロジェクトはゼロから再構築されます。
意味。ハーネスエンジニアリングは、

1 つまたは複数のモデルを信頼できるエージェント システムに変える、制御、実行、安全性、評価、およびトレーニングのインフラストラクチャ。プロンプトは 1 回の呼び出しに対する 1 つの入力です。ハーネスは、タスク全体にわたるすべての呼び出しを制御します。
この分野には現在、一次文献が存在します。 Anthropic は、エージェント設計 [3] 、コンテキスト エンジニアリング [4] 、ツール設計 [5] 、長期稼働ハーネス [6] に関する 4 つのエンジニアリング ガイドを発行しました。 OpenAI は Codex ハーネスの実践方法を作成しました [2]。 Böckeler は初期の実践者のフレームワークを発表し [1]、2 つの学術チームが実稼働システムをエンドツーエンドで分析しました [17]、[18]。ここでの調査とは、系統的なレビューではなく、主要な工学文献の総合を意味します。
1 つのメンタル モデルがフィールドを組織します。ハーネスをオペレーティング システムとして扱い、モデルをその内部のプロセスとして扱います。 OS は、プロセスがどのメモリを読み取るか、どのシステムコールが存在するか、どの呼び出しが成功するか、どこで実行が行われるか、プロセスが世界について何を学習するかを決定します。パターンはルールに圧縮されます。モデルが提案し、ハーネスが [ 3 ] 、 [ 17 ] を配置します。モデルに独自の権限を付与できる設計により、モデルにルートが与えられました。
前もって開示しておきたいのは、この調査の一部は、最後のセクションで説明されているシステムである GROOM によって調査および維持されているということです。
私たちは、公開されたエンドツーエンドの分析 [17] 、 [18] を使用して、実稼働エージェントをモデルを中心とする同じ 8 つのサブシステムに分解します。図 1 はそれらをマッピングしています。ジョブ、それを機能させるパターン、およびジョブをスキップしたときに発生する障害のサブシステムをクリックします。
MODEL 交換可能なコンポーネント オーケストレーター / ループ コンテキスト エンジン ツールと MCP 権限 サンドボックス メモリ サブエージェント 可観測性と評価 サブシステムをクリックして検査します。
各モデル呼び出しの高シグナルトークンの最小セットを厳選します。 Windows は 100 万トークンに成長しましたが、

xt rot（注意力が二次関数的に薄れるにつれて記憶力が低下すること）は消えませんでした。運用システムは階層化された圧縮を実行します。最初に安価なトリムを実行し、圧力がかかる場合にのみ完全な LLM 要約を実行し、完全な履歴を追加専用のレコードとして保存します。
図 1 · ハーネスの構造 (インタラクティブ)。小型の交換可能なモデルを中心とした 8 つのサブシステム。エージェントループ
ランタイムの形状は ReAct [7] から派生しています。つまり、コンテキストをアセンブルし、モデルを呼び出し、提案をゲートし、分離して実行し、観察し、停止条件に達するまで繰り返します。図 2 は、ハーネスが挿入される場所を示しています。
アセンブル 圧縮 · 取得 モデル呼び出しがアクションを提案 ゲート リスク スコア · 承認 サンドボックス内で実行 トレースへの追加が拒否されたことを観察: 理由がテキストとして返され、モデルは次の状態になるまでループを再計画します: 完了 · 予算が使い果たされた · 評価者の拒否権 · 中止 図 2 · エージェント ループ。 5 行の疑似コード。各矢印はサブシステムを隠します。 3 つの詳細により、プロダクション ループが教科書的な ReAct から分離されます。イングレス ガードレールは、モデルが読み取る前に機密性の高い入力をマスクします。ゲートは提案されたすべてのアクションをスコアリングし、拒否をテキストとしてフィードバックするため、ブロックされたアクションは衝突ではなくステアリングになります [17] 。オンライン評価者は、ストール、繰り返し状態、またはコスト上限でループを停止できます。
Böckeler は、すべてのハーネス コントロールを 2 つの軸で分類します。つまり、アクションの前に操縦するのか、それとも後でチェックするのか、決定論的な計算として実行するのかモデル推論として実行するのかということです [1] 。図 3 は、この文書で取り上げるすべてのメカニズムを分類したものです。
静的ガイド 推論ガイド 計算センサー 推論センサー スキーマ · ホワイトリスト規約ファイル · サンドボックス構成コスト ランタイム スキル · 取得されたサンプル エージェント メモリ適応、コスト トークン テスト · 型チェッカー · リンター 安価でドリフトフリーの LLM-as-judge · サブエージェント cov のレビュー

ers セマンティクスでは、決定論的モデル推論後のチェック前にキャリブレーション ステアリングが必要です。 図 3 · 2 つの制御軸 [1] 。成熟したハーネスは 4 つの象限すべてを満たします。センサーのないガイドは、ルールが機能したかどうかを決して学習しません。ガイドのないセンサーは同じ間違いを繰り返します。コンテキストエンジニアリング
Anthropic は目標を一文で述べています: 望む結果の確率を最大化する高信号トークンの最小セットを見つけることです [4] 。
100 万トークンのウィンドウでは制限は解除されませんでした。モデルは、ウィンドウがいっぱいになるにつれて、ウィンドウ内の情報をより悪く思い出します。主な説明: ソフトマックス アテンションにより、各トークンの影響がより多くの競合他社に分散され、トレーニング データの偏りが短くなるため、長距離の再現が不十分になります [4] 。メカニズムがまだ議論されている場合でも、その効果は確実に現れます。実際、位置 400K で読み取られたモデルは、動作を駆動するには希薄すぎる可能性があります。図 4 は直感を構築します。
25 % 50 % 75 % 100 % 8,000 30,000 100,000 300,000 ウィンドウ内の 1M トークン (対数スケール) 推定再現率 ハーネスなし ハーネスあり 実際に使用されたウィンドウ: 36,000 トークン 推定再現率: 生 46 % → 活用済み 93 % 直観のための説明モデル — パラメータは教育的なものであり、測定されたものではありません。図 4 · コンテキストの腐敗とその緩和策 (対話型)。曲線は教育モデルです。定性的な形状はエンジニアリング情報源の報告と一致しています。ウィンドウがいっぱいになると再現率が低下し、キュレーションがウィンドウ サイズを上回ります [ 4 ] 、 [ 17 ] 。 4つのパターン
実稼働システムは、それぞれ異なる障害に対処する 4 つの技術を組み合わせています [4] 、 [8] 。
圧縮チューニングには、名前を付ける価値のある非対称性があります。未解決のバグを削除するサマリーには、タスクのコストがかかります。冗長なツール出力を保持する要約にはトークンがかかります。リコールのためにチューニングしてから、[4] をトリミングします。
これら 4 つのパターンで窓の内側の腐敗を管理します。ノウ

エージェントがウィンドウにロードされるウェッジも、クロックが遅いと腐ってしまいますが、このリストにはこれに対処するループはありません。この論文は、GROOM に関するその問題に戻ります。
Claude Code はこの用語を一般化しています。パターン 2 の LLM 要約は、5 つの圧縮層のうち最後で最も高価な層にすぎず、すべてのモデル呼び出しの前にコスト順に実行され、圧力が残っている間のみ段階的に増加します [17] 。
レイヤー 4 には、この領域で最も深い設計アイデアが組み込まれています。つまり、読み取り時に投影を伴う追加専用の状態です。システムはトランスクリプトを決して破壊しません。圧縮するとビューが生成されます。 「なぜエージェントがそうしたのか」を数週間後でも答えることができますが、EU AI 法では、これがデバッグの利便性から義務に変わりました [24] 。
ツール、プロトコル、およびアクション サーフェス
Tool Design as a Measured Variable
生産レポートで最もよく挙げられる失敗は、機能が重複して肥大化したツールセットです。人間のリトマス試験紙はそれを早期に発見します。どのツールが状況に適合するかを判断できない場合、モデルも判断できず、実行時に間違ったツールが呼び出されて曖昧さが表面化します [5] 。
ソース全体に適用される設計ルール [5]、[17]:
トークン効率の高い出力。すべてのツールの結果は、残りの会話のコンテキストに反映されます。
1 つのツール、1 つの仕事。重なっているツールを結合するか、境界をシャープにします。
リスクの注釈。 readOnly 、 destructive 、 idempotent にタグを付けて、権限システムにそれらをルーティングさせます。
デコード時のスキーマの強制。制約されたデコードにより、不正な型、フィールドの欠落、無効な列挙など、不正な形式の呼び出しが表現できなくなります。パスが存在するかどうか、コマンドが安全かどうかなどの意味論的な有効性は、依然としてゲートに属します。
スキーマの読み込みの遅延。 Keep tool names visible, load full schemas on demand, and hundreds of tools stop consuming the window.
Claude Code には最大 54 個の組み込みツールが同梱されており、そのうち 19 個は常時稼働します

フィーチャー フラグの後ろには 35 あるため、モデルが常に見るサーフェスははるかに小さくなります [17] 。これを、厳選されたファーストパーティ サーフェスの参照点として扱います。それ以降、ディスパッチはプロトコルに属します。
相互運用性は 2024 年以降に標準化されます。重要なのは 2 つのプロトコルです。
2026 年にサードパーティ サービス向けに特注の配管を作成することは、通常、設計上の問題を示します。 MCP サーバーを使用します。主要なファーストパーティ ツールは依然として手作業で構築されており、そうあるべきです。プロトコル化のより深い結果として、推論、ツールトランスポート、フェデレーションが独立した標準を持つレイヤーに分離されるようになりました。
安全性: 権限、分離、多層防御
Claude Code は 7 つの独立した安全レイヤーを積み重ねます。単一のレイヤーはアクションをブロックできます [17] :
スキーマの事前フィルタリング。拒否リストにあるツールはモデルに表示されません。 It cannot attempt what it cannot see.
拒否第一のルール。許可ルールなしは拒否ルールよりも優先されます。
許可モード。読み取り専用のプランニングから完全な自律性までの大まかなダイヤル。
安全分類器。 2 段階のモデルは、ユーザーのプロンプトなしで危険なアクションを判断します。
サンドボックスの分離。上のレイヤーが失敗した場合のダメージを制限します。
許可なしの復元。セッション許可は再開後に期限切れになります。すべての承認は再評価されます。
ライフサイクルフック。外部ポリシーは、あらゆるアクションを拒否、変更、またはログに記録できます。
レイヤ 1 は強調に値します。エージェントを制限するということは、より小さなツール スキーマを使用してエージェントを構築することを意味します。動作を促すと、機能がそのまま残ります。
Claude Code 研究で報告された Anthropic 自動モード テレメトリによると、ユーザーは許可プロンプトの 93% を承認しています [17] 。このままでは、プロンプトは習慣として機能し、習慣は何も捕らえません。 The fix runs in two directions: prompt only for consequential actions, and add checks that work without user attention.同じテレメトリは、実績に応じて信頼が拡大していることを示しています: 自動承認

若いインストールではアクションの約 20% から、数百のセッション後には 40% 以上に増加します [17]。
シェルにアクセスできるモデルは、任意のコマンドを生成できます。 Isolation tiers bound what the worst command can do:
任意の層の内部: 最小権限、一時的な認証情報、明示的な出力制御。
メモリと長距離演算
運用システムは、MemGPT の OS の類似点を継承した 3 層モデルに落ち着きました [8]。
制作後の事後分析では、ある教訓が繰り返されます。それは、事実が保管庫に保管され、適切な瞬間にウィンドウに到達しなかったため、記憶を検索するときに失敗するということです。まずはアーキテクトの検索。そして、誰かが「エージェントが X を学習した」と言ったら、どの層がそれを学習したかを尋ねてください。モデルの重みはほとんどなく、ハーネス構成は時々、コンテキスト メモリがほとんどの場合です。
新しいセッションはそれぞれ空から開始されるため、多くのコンテキスト ウィンドウにまたがる作業にはハンドオフ プロトコルが必要です。 Anthropic は人間のシフト変更のパターンをモデル化します [6]。イニシャライザ セッションは、パス フラグを含む詳細な機能リスト、進行状況ログ、環境ブートストラップ スクリプト、ベースライン コミット、および規約ファイルなどの永続的な成果物を生成します。以降のすべてのセッションでは同じ手順に従います。
進行状況ログと git 履歴を読みます。
何かを触る前にテストスイートを実行してください。
を選択してください

[切り捨てられた]

## Original Extract

Agent quality lives below the model: in the loop, the context engine, the tool surface, the safety stack, and the evaluator. A survey of harness engineering as of mid-2026, introducing GROOM, an open-source self-maintaining knowledge harness.

The 98% Problem: A Survey of Harness Engineering for AI Agents — BeConfident Labs The 98% Problem: A Survey of Harness Engineering for AI Agents
(a) trigger sequence (b) maintenance ops Agent session Launcher GROOM agent wiki/ checks: enabled? last run ≥ 24 h ago? skill activation returns in < 100 ms spawn · detached read all pages page contents edits · one op append journal entry t conversation proceeds, never blocked lint Fix the form. Never touch the meaning. prune Cut what repeats. Leave fewer lines behind. expand Ingest what changed. Touch at most 3–6 files. research Admit new papers. Demand citations first. one op per cycle, chosen by config AUTHORS
Author contributions statement below.
Frontier models converged between 2023 and 2026. For most production tasks, swapping one top model family for another no longer changes the outcome much. The systems that win differ a layer down: in how they call the model, what context they feed it, which tools they expose, where execution happens, and how they measure outcomes. Practitioners call that layer the harness .
The first number gives this paper its title. A community dissection of Claude Code , the most documented production agent, estimates that about 1.6% of the code decides what the model does; the rest assembles context, dispatches tools, checks permissions, sandboxes execution, persists state, and recovers from failure [ 17 ] . Lines of code measure where the engineering lives, not where the capability lives, so treat the ratio as an order-of-magnitude claim. It still names a real condition: the 98% problem . The layer that decides agent quality is the one nobody benchmarks, few teams staff, and every project rebuilds from scratch.
Definition. Harness engineering is the design and operation of the control, execution, safety, evaluation, and training infrastructure that turns one or more models into a dependable agentic system. Prompts are one input to one call; the harness governs every call across a task.
The discipline now has primary literature. Anthropic published four engineering guides on agent design [ 3 ] , context engineering [ 4 ] , tool design [ 5 ] , and long-running harnesses [ 6 ] . OpenAI wrote up Codex harness practice [ 2 ] . Böckeler published an early practitioner framework [ 1 ] , and two academic teams dissected production systems end to end [ 17 ] , [ 18 ] . Survey here means a synthesis of that primary engineering literature, not a systematic review.
One mental model organizes the field. Treat the harness as an operating system and the model as a process inside it. The OS decides what memory the process reads, which syscalls exist, which calls succeed, where execution happens, and what the process learns about the world. The pattern compresses to a rule: the model proposes, the harness disposes [ 3 ] , [ 17 ] . A design that lets the model grant its own permissions has handed it root.
One disclosure up front: parts of this survey were researched and maintained by GROOM, the system its final sections describe.
We decompose the production agents with public end-to-end dissections [ 17 ] , [ 18 ] into the same eight subsystems around the model. Figure 1 maps them. Click a subsystem for its job, the pattern that makes it work, and the failure you get when you skip it.
MODEL interchangeable component Orchestrator / Loop Context Engine Tools & MCP Permissions Sandbox Memory Sub-agents Observability & Evals Click a subsystem to inspect it.
Curates the smallest set of high-signal tokens for each model call. Windows grew to 1M tokens but context rot — recall degradation as attention thins quadratically — did not vanish. Production systems run layered compaction: cheap trims first, full LLM summarization only under pressure, with full history preserved as an append-only record.
Figure 1 · Anatomy of a harness (interactive). Eight subsystems around a small, swappable model. The Agent Loop
The runtime shape descends from ReAct [ 7 ] : assemble context, call the model, gate the proposal, execute in isolation, observe, repeat until a stop condition. Figure 2 shows where the harness inserts itself.
Assemble compaction · retrieval Model call proposes an action Gate risk score · approval Execute inside a sandbox Observe append to trace denied: the reason returns as text and the model re-plans loop until: done · budget exhausted · evaluator veto · abort Figure 2 · The agent loop. Five lines of pseudocode. Each arrow hides a subsystem. Three details separate production loops from textbook ReAct. Ingress guardrails mask sensitive input before the model reads it. The gate scores every proposed action and feeds denials back as text, so a blocked action becomes steering rather than a crash [ 17 ] . Online evaluators can stop the loop on stalls, repeated states, or cost ceilings.
Böckeler classifies every harness control on two axes: does it steer before the action or check after, and does it run as deterministic computation or as model inference [ 1 ] . Figure 3 sorts every mechanism you will meet in this paper.
Static guides Inferred guides Computational sensors Inferential sensors schemas · allowlists convention files · sandbox config cost zero at runtime skills · retrieved examples agentic memory adapts, costs tokens tests · type checkers · linters cheap, drift-free LLM-as-judge · review subagents covers semantics, needs calibration steer before check after deterministic model inference Figure 3 · Two axes of control [ 1 ] . Mature harnesses fill all four quadrants. Guides without sensors never learn whether the rules worked. Sensors without guides repeat the same mistakes. Context Engineering
Anthropic states the goal in one sentence: find the smallest set of high-signal tokens that maximizes the probability of the outcome you want [ 4 ] .
Million-token windows did not remove the limit. A model recalls in-window information worse as the window fills. The leading explanations: softmax attention spreads each token’s influence across more competitors, and training data skews short, so long-range recall is undertrained [ 4 ] . The effect shows up reliably even where the mechanism is still debated. A fact the model read at position 400K can sit too dilute to drive behavior. Figure 4 builds the intuition.
25 % 50 % 75 % 100 % 8K 30K 100K 300K 1M tokens in the window (log scale) estimated recall no harness with harness Window actually used: 36K tokens Estimated recall: raw 46 % → harnessed 93 % Illustrative model for intuition — parameters are pedagogical, not measured. Figure 4 · Context rot and its mitigations (interactive). The curve is a pedagogical model. The qualitative shape matches what the engineering sources report: recall falls as windows fill, and curation beats window size [ 4 ] , [ 17 ] . Four Patterns
Production systems combine four techniques, each against a different failure [ 4 ] , [ 8 ] :
Compaction tuning has an asymmetry worth naming. A summary that drops an unresolved bug costs you the task. A summary that keeps redundant tool output costs you tokens. Tune for recall, then trim [ 4 ] .
These four patterns manage rot inside the window. The knowledge an agent loads into the window rots too, on a slower clock, and no loop in this list addresses it. The paper returns to that problem with GROOM.
Claude Code generalizes the term: the LLM summary of pattern two is only the last and most expensive of five compaction layers, run in cost order before every model call, escalating only while pressure remains [ 17 ] :
Layer 4 carries the deepest design idea in this area: append-only state with projection at read time. The system never destroys transcripts. Compaction produces a view. You keep the ability to answer “why did the agent do that” weeks later, which the EU AI Act turns from a debugging convenience into an obligation [ 24 ] .
Tools, Protocols, and the Action Surface
Tool Design as a Measured Variable
The most-cited failure in production reports: bloated tool sets with overlapping functions. Anthropic’s litmus test catches it early. If you cannot say which tool fits a situation, the model cannot either, and the ambiguity surfaces as wrong-tool calls at runtime [ 5 ] .
The design rules that hold up across sources [ 5 ] , [ 17 ] :
Token-efficient outputs. Every tool result lives in context for the rest of the conversation.
One tool, one job. Merge overlapping tools or sharpen the boundary.
Risk annotations. Tag readOnly , destructive , idempotent , and let the permission system route on them.
Schema enforcement at decode time. Constrained decoding makes malformed calls unrepresentable: wrong types, missing fields, invalid enums. Semantic validity, like whether the path exists or the command is safe, still belongs to the gate.
Deferred schema loading. Keep tool names visible, load full schemas on demand, and hundreds of tools stop consuming the window.
Claude Code ships up to 54 built-in tools, 19 always on and 35 behind feature flags, so the surface the model sees at any moment is far smaller [ 17 ] . Treat that as a reference point for a curated first-party surface. Past it, dispatch belongs to protocols.
Interoperability standardized after 2024. Two protocols matter:
Writing bespoke plumbing for a third-party service in 2026 usually signals a design problem; use its MCP server. Core first-party tools remain hand-built, and should be. The deeper consequence of protocolization: reasoning, tool transport, and federation now separate into layers with independent standards.
Safety: Permissions, Isolation, Defense in Depth
Claude Code stacks seven independent safety layers. Any single layer can block an action [ 17 ] :
Schema pre-filtering. Deny-listed tools never appear to the model. It cannot attempt what it cannot see.
Deny-first rules. No allow rule outranks a deny rule.
Permission modes. Coarse dials from read-only planning to full autonomy.
Safety classifier. A two-stage model judges risky actions without user prompts.
Sandbox isolation. Bounds the damage when the layers above fail.
Permission non-restoration. Session grants expire across resume; every approval re-evaluates.
Lifecycle hooks. External policy can deny, modify, or log any action.
Layer 1 deserves the emphasis. Restricting an agent means constructing it with a smaller tool schema. Prompting it to behave leaves the capability in place.
Users approve 93% of permission prompts, per Anthropic auto-mode telemetry reported in the Claude Code study [ 17 ] . At that rate a prompt works as a habit, and a habit catches nothing. The fix runs in two directions: prompt only for consequential actions, and add checks that work without user attention. The same telemetry shows trust widening with track record: auto-approval grows from roughly 20% of actions in young installations to above 40% after several hundred sessions [ 17 ] .
A model with shell access can generate any command. Isolation tiers bound what the worst command can do:
Inside any tier: least privilege, ephemeral credentials, explicit egress control.
Memory and Long-Horizon Operation
Production systems settled on a three-tier model that descends from MemGPT’s OS analogy [ 8 ] :
Production post-mortems repeat one lesson: memory fails at retrieval, since the fact sat in storage and never reached the window at the right moment. Architect retrieval first. And when someone says “the agent learned X,” ask which layer learned it: model weights almost never, harness configuration sometimes, contextual memory in most cases.
Work that spans many context windows needs a handoff protocol, since each new session starts blank. Anthropic models the pattern on human shift changes [ 6 ] . An initializer session produces durable artifacts: a granular feature list with pass flags, a progress log, an environment bootstrap script, a baseline commit, and a conventions file. Every later session follows the same steps:
Read the progress log and the git history.
Run the test suite before touching anything.
Pick th

[truncated]
