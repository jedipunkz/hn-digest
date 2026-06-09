---
source: "https://www.claudecertifiedarchitects.com/"
hn_url: "https://news.ycombinator.com/item?id=48460948"
title: "A free diagnostic for the Claude Certified Architect exam"
article_title: "Claude Certified Architect Practice Test | CCA Prep"
author: "jockorules"
captured_at: "2026-06-09T16:06:33Z"
capture_tool: "hn-digest"
hn_id: 48460948
score: 1
comments: 0
posted_at: "2026-06-09T13:34:52Z"
tags:
  - hacker-news
  - translated
---

# A free diagnostic for the Claude Certified Architect exam

- HN: [48460948](https://news.ycombinator.com/item?id=48460948)
- Source: [www.claudecertifiedarchitects.com](https://www.claudecertifiedarchitects.com/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T13:34:52Z

## Translation

タイトル: クロード認定アーキテクト試験の無料診断
記事タイトル: クロード認定建築士模擬試験 | CCAの準備
説明: 400 のシナリオベースの質問、時間指定された練習テスト、および 5 つの試験ドメインすべてにわたるドメイン加重スコアリングを使用して、CCA Foundations 試験の準備をします。

記事本文:
独立した試験準備 · Anthropic と提携、認可、承認されていません。 ✕
クロード建築士認定
ホーム
試験
ブログ
JS 駆動のボタンではなく、
静的ページへの同じサイトのリンクであれば、それ以上は何も必要ありません。 -->
無料診断を開始する
検定試験 ↗
レッスン
模擬試験
進捗状況
ログイン
無料でサインアップ
無料
ログアウト
×
ログイン
クロード認定建築家へようこそ!
無料アカウントの準備ができました。ほとんどの人はここから始めます。
クロード認定建築士模擬試験
試験形式に合わせて設計された現実的なシナリオベースの質問で、CCA Foundations 試験の準備をします。 5 つの試験分野すべてを学習して、自信を持って合格してください。
CCA Foundations の 5 つのドメインすべてをカバーする完全な独立した準備プログラムです。すべての質問は公式のシナリオベースの形式に一致するように作成されています。
公式 CCA Foundations 試験ガイドに基づいて作成 · 毎週更新 · 10 日間の返金保証
購入する前に質問の質を確認してください
実際のシナリオベースの質問には、すべての答えが完全に説明されています。まさに銀行のフルバンクの 400 問と同じです。
コーディング エージェントはループ内でツールを呼び出します。開発者はループをいつ終了するかを決定する必要があります。正確で堅牢な終了信号はどれですか?
アシスタントのテキストに「完了」などの完了単語が含まれなくなったら停止します。
B 状態に関係なく、最大 10 回の固定反復後に停止します。
C 応答の stop_reason が tools_use である間、ツールの実行とループを続けます。 stop_reason が end_turn の場合は停止します。
D いずれかのツールがエラーを返したらすぐに停止します。
API は、生成が停止した理由を stop_reason 経由で報告します。 tool_use は、ツールを実行し、結果を追加し、ループバックすることを意味します。 end_turn はクロードが終了したことを意味し、ループが終了します。補完単語のテキストを解析することは脆弱なアンチパターンです。固定イテレーションキャップは安全です

メインコントロールではなくバックストップ。最初のツール エラーで停止すると、エージェントは回復できなくなります。
チームはプル リクエストを自動レビューするために CI パイプラインにクロード コードを追加します。ジョブは無期限にハングし、出力を生成する代わりにタイムアウトになります。どの修正で解決しますか?
B -p (印刷) フラグを指定して claude を実行し、プロンプトを引数として渡します。
D /dev/null から標準入力をリダイレクトします。
-p を指定しないと、claude は対話型セッションを開始し、入力が来るまで stdin で待機するため、ハングします。 -p / --print フラグはヘッドレスで実行されます。プロンプトが表示され、結果が標準出力に出力され、終了します。 CLAUDE_HEADLESS と --batch は実際のフラグではありません。 /dev/null から stdin をリダイレクトすることは、実際に必要な出力と終了の動作を実現しないハックです。
あなたは、エージェントの外部インベントリ API をクエリするツールを設計しています。 API は、「アイテムが見つかりません」またはレート制限エラーを返すことがあります。エージェントの回復能力を最大限に高めるために、ツールはこれらをどのように処理すべきでしょうか?
A エージェント ループを終了する例外を発生させ、人間が介入できるようにします。
B ツールの結果としてエラーを説明する構造化された結果 (明確なメッセージとエラーの種類) を返すため、クロードはそれを読んで続行方法を決定できます。
C 空の文字列を返すので、クロードはデータが存在しないとみなします。
D API が最終的に成功するまで、ツール内でサイレントに再試行します。
ツールの結果はクロードにフィードバックされるため、明確な構造化エラーがある場合は、再試行、代替案の試行、または報告を行うことができます。エラーが発生するたびに終了すると、自己修正が削除されます。空の文字列は失敗を隠し、幻覚を招きます。サイレントで無限に再試行すると、エージェントがハングし、レート制限のある API が使用できなくなる可能性があります。
すべての CCA 試験ドメインをカバーする 5 つのモジュール
400 のシナリオベースの解説付き練習問題
時間制限付き練習セッション: 20 分、40 分、60 分
全長模擬試験（120分）

s、60 問）
実際の試験と一致するドメイン加重スコアリング (エージェント 27%、クロード コード 20%、プロンプト 20%、ツール/MCP 18%、コンテキスト 15%)
CCA 試験の進化に応じた生涯アップデート
Stripe を搭載 · 10 日間の返金保証 · 安全なチェックアウト
詳しい説明はこちらをご覧ください)。
両方の数値が同じライブ CSS 変数 nav 自体から読み取られるようになりました。
から構築されているため、このバナーは常にナビゲーションの画面と同じ高さでレンダリングされます。
実際の下端 - 決して重なり合ったり、隙間を残したりしません。 -->
支払いが成功しました！クロード認定建築家コースへようこそ。却下
模擬試験
時間制限のあるセッションを選択するか、完全な認定試験を受けてください。
どれから始めればよいかわからないですか?まずは無料の診断を受けてください。選択する前に最も弱いドメインを確認してください。
ランダムなトピックをカバーする 10 問の高速ドリル。毎日の練習に最適です。
コア アーキテクチャ パターンとモデルの選択についての 20 の質問セッション。
すべての領域にわたる 30 の質問からなる包括的なセッション。
完全な認定試験はロックされています
実際の認定試験をシミュレートした完全な 60 問の模擬テスト。
最初の練習セッションを受講してください。ここで、ドメインの習熟度、全体的な準備状況の推定、学習時間をどこに集中すべきか追跡を開始します。
0 回の演習試行からの非公式の推定値。各ドメインの実際の試験の割合で重み付けされています。
公式の CCA 試験では、720 / 1,000 の合格基準が使用されます。これは公式のスコアではなく、試験当日の結果と正確には一致しません。
これまでの最低スコアのドメインは、実際の試験の 0 % です。通常、ここで追加の学習時間を費やすと、全体のスコアが最も大きく変わります。
2026 年に取得する価値のある AI 認定資格に関する実用的なガイド - どの認定資格が本当のスキルを示すのか、どれを積み重ねるべきか、役割に基づいて優先順位を付ける方法。
CCA 認定の正直な内訳

実際のコスト、それによって得られるもの、そして 2026 年に誰がそれを追求すべきか（そしてすべきではないのか）。
CCA Foundations 試験の難易度についての現実的な見方 - 何が難しいのか、ほとんどの人がどこで不合格になるのか、合格するにはどのくらいのスコアが必要なのか。
CCA Foundations 試験は、Claude を使用して実稼働グレードのアプリケーションを設計および構築する専門家に対する Anthropic の公式認定資格です。エージェントティック アーキテクチャ、クロード コード構成、プロンプト エンジニアリング、ツール設計と MCP、コンテキスト管理の 5 つの領域にわたる知識をテストします。
CCA Foundations 試験は、制限時間 120 分のシナリオベースの多肢選択問題 60 問で構成されています。当社の練習プラットフォームは、5 つのドメインすべてにわたって 400 の独立した練習問題を提供します。
CCA Foundations 試験では、100 ～ 1,000 点の段階的な採点システムが使用されます。合格点は 1,000 点中 720 点です。これはスケールされたスコアであり、正解した質問の割合ではありません。
エージェントティック アーキテクチャが 27% と最も高い比重を占め、続いてクロード コード構成とプロンプト エンジニアリングがそれぞれ 20%、ツール設計と MCP が 18%、コンテキスト管理が 15% となっています。
ほとんどの学生は、毎日の練習セッションを使用して 1 ～ 2 週間で準備します。クイック スプリント モード (10 問、20 分) は毎日の練習に最適ですが、120 分間の完全な試験シミュレーションはテスト前の最後の数日間にお勧めします。
MCP (Model Context Protocol) は、Claude を外部ツールやデータ ソースに接続するための Anthropic のオープン スタンダードです。 CCA 試験の 18% を占め、MCP サーバー アーキテクチャ、ツール スキーマ設計、トランスポート層 (stdio および SSE)、およびセキュリティに関する考慮事項をカバーします。
PRECISE は、プロンプト エンジニアリング ドメイン (試験の 20%) でテストされるプロンプト エンジニアリング フレームワークです。これは、ペルソナ、役割、明示的な指示、コンテキスト、指示、ステップ、例を表します。

効果的なシステム プロンプトを設計するための構造化されたアプローチ。
すべての質問はシナリオベースです。あなたは現実世界の建築状況に置かれており、最善の決定を選択する必要があります。単純な想起や定義に関する質問はありません。トピックには、エージェント ループの設計、ツールの選択、プロンプトの最適化、およびコンテキスト ウィンドウの管理が含まれます。
CCA Foundations 試験は Anthropic の Claude に特化したもので、Constitutional AI、CLAUDE.md 構成システム、Claude Code、Model Context Protocol などの Claude 固有の機能を対象としています。一般的な AI 理論ではなく、実践的なアーキテクチャの決定に焦点を当てています。
CCA Foundations 認定の有効期間は 2 年間です。これは、実稼働グレードの Claude を利用したアプリケーションの構築における専門知識を実証しており、製品やワークフローに Anthropic の AI を採用している組織によって認められています。
これは独立した学習リソースです。 Anthropic と提携、認可、承認されていません。 「Claude」および「Claude Certified Architect」は、それぞれの所有者の商標です。公式認定試験の準備に役立つ非公式の練習資料を提供しています。

## Original Extract

Prepare for the CCA Foundations exam with 400 scenario-based questions, timed practice tests, and domain-weighted scoring across all 5 exam domains.

Independent exam prep · Not affiliated with, authorized by, or endorsed by Anthropic. ✕
Claude Architect Certification
Home
Exam
Blog
rather than a JS-driven button: it's a
same-site link to a static page, nothing more is needed. -->
Start the Free Diagnostic
Official Exam ↗
Lessons
Practice Tests
Progress
Log In
Sign Up Free
FREE
Log out
×
Log In
Welcome to Claude Certified Architects!
Your free account is ready. Here's where most people start:
Claude Certified Architect Practice Test
Prepare for the CCA Foundations exam with realistic, scenario-based questions designed to match the exam format. Study all 5 exam domains and pass with confidence.
A complete, independent prep program covering all five CCA Foundations domains — every question written to match the official scenario-based format.
Built from the official CCA Foundations exam guide · Updated weekly · 10-day money-back guarantee
See the question quality before you buy
Real scenario-based questions, every answer fully explained — just like the 400 in the full bank.
A coding agent calls tools inside a loop. A developer needs to decide when to terminate the loop. Which is the correct, robust termination signal?
A Stop when the assistant’s text no longer contains a completion word like “done.”
B Stop after a fixed maximum of 10 iterations, regardless of state.
C Keep executing tools and looping while the response’s stop_reason is tool_use ; stop when stop_reason is end_turn .
D Stop as soon as any tool returns an error.
The API reports why generation stopped via stop_reason . tool_use means run the tools, append results, and loop back; end_turn means Claude is finished, so the loop ends. Parsing text for completion words is a brittle anti-pattern. A fixed iteration cap is a safety backstop, not the main control. Stopping on the first tool error prevents the agent from recovering.
A team adds Claude Code to a CI pipeline to auto-review pull requests. The job hangs indefinitely and times out instead of producing output. Which fix resolves it?
B Run claude with the -p (print) flag and pass the prompt as an argument.
D Redirect stdin from /dev/null .
Without -p , claude starts an interactive session and waits on stdin for input that never comes, so it hangs. The -p / --print flag runs headless: take the prompt, print the result to stdout, exit. CLAUDE_HEADLESS and --batch are not real flags. Redirecting stdin from /dev/null is a hack that doesn’t give the print-and-exit behaviour you actually want.
You’re designing a tool that queries an external inventory API for an agent. The API sometimes returns “item not found” or rate-limit errors. How should the tool handle these to maximise the agent’s ability to recover?
A Raise an exception that terminates the agent loop so a human can intervene.
B Return a structured result describing the error (clear message and error type) as the tool result, so Claude can read it and decide how to proceed.
C Return an empty string so Claude assumes there’s no data.
D Retry silently inside the tool until the API eventually succeeds.
Tool results are fed back to Claude, so a clear structured error lets it retry, try an alternative, or report back. Terminating on every error removes self-correction. An empty string hides the failure and invites hallucination. Silent infinite retries can hang the agent and hammer a rate-limited API.
5 modules covering all CCA exam domains
400 scenario-based practice questions with explanations
Timed practice sessions: 20 min, 40 min, 60 min
Full-length practice exam (120 minutes, 60 questions)
Domain-weighted scoring matching real exam (Agentic 27%, Claude Code 20%, Prompting 20%, Tools/MCP 18%, Context 15%)
Lifetime updates as the CCA exam evolves
Powered by Stripe · 10-day money-back guarantee · Secure checkout
for the full explanation).
Both numbers are now read from the same live CSS variables nav itself
is built from, so this banner always renders flush against the nav's
actual bottom edge — never overlapping it, never leaving a gap. -->
Payment successful! Welcome to the Claude Certified Architect course. Dismiss
Practice Tests
Choose a timed session or take the full certification exam.
Not sure which to start with? Take the free diagnostic first — see your weakest domain before you choose.
A fast 10-question drill covering random topics. Great for daily practice.
A 20-question session hitting core architecture patterns and model selection.
A 30-question comprehensive session across all domains.
Full Certification Exam LOCKED
The complete 60-question practice test simulating the real certification exam.
Take your first practice session and we'll start tracking your domain mastery, an overall readiness estimate, and where to focus your study time — right here.
Informal estimate from 0 practice attempt(s), weighted by each domain's real exam percentage.
The official CCA exam uses a scaled 720 / 1,000 passing standard — this is not an official score and won't exactly match exam-day results.
Your lowest-scoring domain so far is 0 % of the real exam. Spending extra study time here typically moves your overall score the most.
A practical guide to the AI certifications worth pursuing in 2026 — which ones signal real skills, which to stack, and how to prioritise based on your role.
An honest breakdown of what the CCA certification actually costs, what it gets you, and who should (and shouldn't) pursue it in 2026.
A realistic look at the CCA Foundations exam difficulty — what makes it hard, where most people fail, and what score you need to pass.
The CCA Foundations exam is Anthropic's official certification for professionals who design and build production-grade applications with Claude. It tests knowledge across five domains: Agentic Architecture, Claude Code Configuration, Prompt Engineering, Tool Design & MCP, and Context Management.
The CCA Foundations exam consists of 60 scenario-based multiple choice questions with a 120-minute time limit. Our practice platform offers 400 independent practice questions across all five domains.
The CCA Foundations exam uses a scaled scoring system of 100–1,000. The passing score is 720 out of 1,000 — this is a scaled score, not a percentage of questions answered correctly.
Agentic Architecture carries the highest weight at 27%, followed by Claude Code Configuration and Prompt Engineering at 20% each, Tool Design & MCP at 18%, and Context Management at 15%.
Most students prepare in 1–2 weeks using daily practice sessions. Our Quick Sprint mode (10 questions, 20 minutes) is ideal for daily practice, while the full 120-minute exam simulation is recommended in the final days before your test.
MCP (Model Context Protocol) is Anthropic's open standard for connecting Claude to external tools and data sources. It accounts for 18% of the CCA exam and covers MCP server architecture, tool schema design, transport layers (stdio and SSE), and security considerations.
PRECISE is a prompt engineering framework tested in the Prompt Engineering domain (20% of the exam). It stands for Persona, Role, Explicit instructions, Context, Instructions, Steps, and Examples — a structured approach for designing effective system prompts.
All questions are scenario-based — you are placed in a real-world architectural situation and must choose the best decision. There are no simple recall or definition questions. Topics include agent loop design, tool selection, prompt optimization, and context window management.
The CCA Foundations exam is specific to Anthropic's Claude and covers Claude-specific capabilities like Constitutional AI, the CLAUDE.md configuration system, Claude Code, and the Model Context Protocol. It focuses on practical architecture decisions rather than general AI theory.
The CCA Foundations certification is valid for 2 years. It demonstrates expertise in building production-grade Claude-powered applications and is recognized by organizations adopting Anthropic's AI in their products and workflows.
This is an independent study resource. It is not affiliated with, authorized by, or endorsed by Anthropic. ‘Claude’ and ‘Claude Certified Architect’ are trademarks of their respective owner. We provide unofficial practice materials to help you prepare for the official certification exam.
