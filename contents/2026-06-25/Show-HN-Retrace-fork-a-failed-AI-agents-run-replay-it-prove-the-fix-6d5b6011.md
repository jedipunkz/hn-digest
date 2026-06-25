---
source: "https://retraceai.tech"
hn_url: "https://news.ycombinator.com/item?id=48676599"
title: "Show HN: Retrace fork a failed AI agents run, replay it, prove the fix"
article_title: "Retrace — The Execution Replay Engine for AI Agents — Retrace"
author: "Yashwanthbogam"
captured_at: "2026-06-25T17:51:43Z"
capture_tool: "hn-digest"
hn_id: 48676599
score: 1
comments: 0
posted_at: "2026-06-25T17:24:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Retrace fork a failed AI agents run, replay it, prove the fix

- HN: [48676599](https://news.ycombinator.com/item?id=48676599)
- Source: [retraceai.tech](https://retraceai.tech)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T17:24:39Z

## Translation

タイトル: HN の表示: 失敗した AI エージェントの実行をフォークし直し、それを再生し、修正を証明します
記事のタイトル: Retrace — AI エージェント用の実行再生エンジン — Retrace
説明: AI エージェントの実行を記録、再生、フォーク、共有します。すべての LLM 呼び出し、ツールの呼び出し、エージェントが発生するエラーを確認し、数秒でデバッグと反復を行います。 1,000 トレース/月までは無料。
HN テキスト: AI エージェントが実行した記録を再追跡することで、記録を段階的に再生し、任意の時点から修正にフォークして、結果をリンクとして共有できます。

記事本文:
Retrace — AI エージェント用の実行リプレイ エンジン — Retrace retrace 製品レコード 1 つのデコレータが、あらゆる決定からのすべての LLM 呼び出しのフォークとリプレイ ブランチをキャプチャし、将来をリプレイします 修正を証明する 修正を再実行して評決を取得します — うまくいきましたか?評価ゲート 不正なデプロイをブロックする CI/CD 品質ゲート ガードレール ランタイムの安全性 — 失敗する前にエージェントを停止する 執行サーキット ブレーカー — 暴走ループと予算のパンクをブロックする 障害分類法 MAST — エージェントの実行が失敗する理由、自動分類されたセッション 因果関係グラフによるマルチエージェント トレース 料金設定 Playground ドキュメント 変更ログ サインイン はじめに AI エージェントの動作のための CI
AI エージェントが失敗した理由を把握します。修正を証明してください。
Retrace は、すべての LLM 呼び出し、ツール呼び出し、およびエージェントによる決定を記録します。そのため、実行を再生し、問題が発生した正確なステップからフォークして、出荷前に修正を証明できます。
| npm i retrace-sdk OpenAI、Anthropic、Gemini — あらゆる LLM、あらゆるフレームワークで動作します。 // コアループ
エージェントが行ったことを記録します。再生し、フォークし、修正し、修正を証明します。
単一のデコレータから評決まで、Retrace がエンドツーエンドで所有する信頼性ループ。
1 つのデコレータがすべての LLM 呼び出し、ツール呼び出し、および決定をキャプチャするため、本番環境の障害は永続的な回帰テストになります。
記録された実行を再実行するか、中断した正確なステップからフォークします。これにより、症状ではなく本当の原因をデバッグできます。
問題が発生したステップ (プロンプト、ツール入力、またはモデル) を変更し、フォークを再実行すると、出荷前に修正されたパスが表示されます。
失敗した実行に対して変更を再実行し、判定を得ます。つまり、希望ではなく修正を出荷することになります。
↻ すべての実行が次の実行につながります - 記録、修正、繰り返し。
本物の録音です。自分でこすってください。
実際のトレース UI - 記録されたインシデント - OOM 症状で停止し、失敗とマークされた RCA 実行。

そのステップから出発して、本当の原因を追跡しました。デプロイによってバッチ サイズが 50 倍に増加しました。タイムラインをスクラブし、任意のスパンを開いて、フォークの分岐を観察します。
ステップ 4 でフォーク → 再実行 失敗した実行 → 判定: 検証済みの retrace.ai.generate 6.18s · $0.0001 を修正
retrace.ai.generate 5.83秒 · $0.0005エラー
↳ retrace.ai.generate 2.79秒 · $0.0003 OK
↳ retrace.ai.generate 2.01s · $0.0001 ok
6 スパン · ステップ 4 で 1 つのエラー · ステップ 4 で分岐 · 修正が確認されました
インタラクティブ — 製品からエクスポートされた実際のトレース (incident-rca · INC-4835)。スパン、コスト、タイミングは記録された値です。
ダッシュボードではなく、完全な信頼性プラットフォーム。
記録された実行ごとに、障害を検出し、制限を適用し、品質を評価し、マルチエージェント システムを理解します。
接地ギャップ、統計的ドリフト、障害クラスター、MAST 障害タイプに自動的にフラグを付けるため、実行が失敗したことだけでなく、ユーザーが失敗する前に、実行が失敗した理由を知ることができます。
予算、ループ、またはステップの制限で実行を停止し、承認待ちの通話前ゲートウェイ経由で実行される前に不正なアクションをブロックします。これにより、バグが静かにクラウド請求に波及することがなくなります。
評価 0 3 評価 評価
ゲートできる品質 — 評価、自動評価ルール、データセット、修正の証明の判定に加えて、回帰がユーザーに到達せずにビルドを失敗させるための CI ゲート。
わかる 0 4 わかる わかる
システム全体 (セッション、エージェント トポロジ、エージェント メモリ、セマンティック検索、プロンプト バージョニング、共有可能なテープ) を確認できるため、複数エージェントの障害がブラック ボックスではなくなります。
接地ギャップ、統計的ドリフト、障害クラスター、MAST 障害タイプに自動的にフラグを付けるため、実行が失敗したことだけでなく、ユーザーが失敗する前に、実行が失敗した理由を知ることができます。
バジェット、ループ、またはステップ制限で実行を停止し、hold-for-a を使用してコール前ゲートウェイ経由で不正なアクションを実行前にブロックします。

承認 — バグがクラウドの請求に静かに波及することはありません。
ゲートできる品質 — 評価、自動評価ルール、データセット、修正の証明の判定に加えて、回帰がユーザーに到達せずにビルドを失敗させるための CI ゲート。
システム全体 (セッション、エージェント トポロジ、エージェント メモリ、セマンティック検索、プロンプト バージョニング、共有可能なテープ) を確認できるため、複数エージェントの障害がブラック ボックスではなくなります。
可観測性により、何が壊れたかがわかります。 Retrace はそれを再実行し、失敗したステップをフォークし、修正を証明します。
集計メトリクスとダッシュボード
正確な実行を再生します。失敗したステップからフォークしてカスケード再実行するため、グラフではなく本当の原因をデバッグできます。
セマンティック検索 — 平易な言葉でバグを説明し、数秒で失敗した実行にジャンプします
すべてを再実行せずに修正をテストする方法はありません
修正の証明 — 失敗した実行に対して変更を再実行し、出荷前に判定を得る
予算を超過したりループしたりした場合に実行を停止するガードレールとサーキットブレーカー - バグによって静かに請求額が膨らむことはありません
実行時強制 - 不正なアクションを実行後に診断するのではなく、実行前にブロックします。
1 つのデコレータ — 任意の Python/TS エージェント (LangChain、CrewAI、LlamaIndex)、ロックインなし
ゼロから最初のトレースまで約 2 分で完了します。
GitHub でサインインし、~3 行をコピーし、エージェントを実行します。最初のトレースはライブでストリーミングされます。管理するインフラストラクチャがありません。
SDK をインストールし、デコレータを 1 つ追加します。 OpenAI、Anthropic、Gemini への呼び出しは自動的にキャプチャされます。
npm i retrace-sdk ↳ フレームワークに依存しない — LangChain、CrewAI、LlamaIndex で動作します。
1 2 3 4 5 6 インポート リトレース リトレース 。 configure(api_key = "rt_..." ) @retrace.record (name = "my-agent" ) def run_agent (prompt): エージェントを返します。 invoke (プロンプト) SDK が接続されました、classify_intent が記録されました
エージェントを実行します。すべての LLM 呼び出し、ツール呼び出し、コストおよびエラー領域

ミリ秒をタイムラインに表示します。
↳ または、ダッシュボードで段階的に再生される様子をご覧ください。
1 retrace トレース テール REC · ライブ assign_intent 0.4 秒 fetch_context 1.8 秒 Generate_response 2.4 秒
壊れた正確なステップからフォークし、入力を変更して再実行し、修正が実際に機能したことを証明します。
↳ verify-fix は、改善された、低下した、または変更されていないという判定を返します。
1 2 3 リトレース フォーク作成 --trace < id > --span < id > --input "グラウンデッド プロンプト" リトレース フォーク リプレイ < id > --wait リトレース トレース verify -fix < id >generate_response error ↳ fork ok verify-fix → 検証済み修正
AI エージェントに希望ではなくテストを提供します。
プロンプトの編集、ツールの変更、またはモデルのアップグレードにより、複数ステップの動作が静かに中断される可能性があります。 Retrace は、各本番環境の失敗を回帰テストに変換し、それをすべての PR の評価ゲートとして実行します。動作が壊れるとビルドは失敗します。各実行をゴールデンベースラインと比較するので、単なる赤や緑の点ではなく、「前回のリリースと比べて低下した関連性」を把握できます。
1 2 3 4-name: Eval Gate run: retrace eval Gate --evaluation $EVAL_ID --trace $TRACE_ID --threshold 0.8 env:RETRACE_API_KEY:${{秘密。 RETRACE_API_KEY } } チェック ✗ retrace / eval-gate の動作が低下 — ビルド失敗 ✓ retrace / eval-gate に合格 5/5 回の実行 — マージがブロックされていない eval ゲートの仕組み → 閉ループ
0 1 本番トレース 実際の実行は失敗します
0 2 キャプチャされた失敗は回帰テストになります
PR の 0 3 評価ゲートはそのテストを再実行します
0 4 緑色のチェック動作が保持 — マージ
単体テストでは制御できないもの
動作が壊れるとビルドは失敗します。そのため、正確に失敗した場合は再出荷するのが難しくなります。 (評価ゲート出口 1 を戻ります。)
評価ゲートとリプレイはモデル アカウントで実行されます。
Google (Gemini) API キーを設定に追加し、評価ゲート ジャッジとすべてのサーバー側リプレイ (フォーク、カスケード) を追加します。

e、および修正の証明 — キーを介してモデルを呼び出すと、トークンが独自のプロバイダー アカウントに請求されます。キーは保存時に検証され、保存時に暗号化され (AES-256-GCM)、最後の 4 文字としてのみ表示され、再度返されることはありません。これをいつでも削除すると、Retrace はプラットフォーム キーに戻ります。
Google · キー ••••a1b2 検証済み ✓ キーの機能
→ 評価ゲートの審査員が PR 実行ごとに採点する
→ フォークとカスケード リプレイによりエージェントが再実行されます
→ 修正の証明は判定を計算します
保存時にプロバイダーに対して検証されます。不正なキーは拒否され、保存されません。
保存時は AES-256-GCM とシークレットごとの派生キーを使用して暗号化されます。
最後の 4 文字としてのみ表示され、ログに記録されることも、返されることもありません。
これを削除すると、ダウンタイムなしでリプレイがプラットフォーム キーにフォールバックします。
クレジットカードは必要ありません。より多くのトレースまたは AI リクエストが必要な場合は、アップグレードしてください。
+ サンドボックス環境: 5,000 トレース/月
+ サンドボックス環境: 50,000 トレース/月
+ 無制限の修正の実行/月
+ サンドボックス環境: 無制限のトレース/月
2分未満。 SDK をインストールし、デコレータを 1 つ追加すると、すぐにストリームをトレースします。管理するインフラストラクチャがありません。
OpenAI、Anthropic、Google Gemini の自動インストルメンテーションを備えた Python および TypeScript SDK。 LangChain、CrewAI、Vercel AI SDK、AutoGen、LlamaIndex など、あらゆるエージェント フレームワークと連携します。
トレース内の任意のスパンを選択し、その入力を変更し、その時点からカスケード再生を再トレースします。フォークからのコンテキストは後続の LLM 呼び出しに流れます。コストとレイテンシの差分を並べて表示します。
エージェントをリアルタイムで監視するランタイム ポリシー。コスト予算、ループ検出、コンテキスト オーバーフロー制限、またはレイテンシー キャップを設定します。制限を超えると、エージェントは HALT コマンドを受け取ります。そのため、暴走ループや予算の超過は請求額を超過するのではなく、制限で停止します。
転送中のTLS

、保存時に暗号化されます。 API キーは SHA-256 ハッシュ化されています。 PII 自動秘匿化は、セキュリティ ベースラインとしてすべてのプランで実行されます。テナントの分離はアプリケーション層で強制されます。すべてのクエリはユーザーごとにスコープが設定され、ガードレール回帰テストによって裏付けられます。
はい。評価ゲート エンドポイント (POST /evaluations/:id/gate) は、しきい値に対する合格/失敗を返します。 CLI コマンド「retrace eval Gate」は失敗するとコード 1 で終了します。これは GitHub Actions に最適です。
LangSmith はトレースと可観測性に重点を置いています。 Retrace は、任意のステップからのインタラクティブなフォークとカスケード リプレイ、暴走エージェントを停止するランタイム ガードレール、接地検出、および修正検証の検証を追加します。
はい。各スパンにはエージェント ID、セッション グループのマルチターン会話が含まれ、エージェント トポロジ グラフにはエージェント間の順序付けとエージェント間の障害モードが表示されます。
推測するのはやめてください。
再生を開始します。
エージェントはエラーが表面化する 3 ステップ前に失敗しました。症状ではなく、本当の原因から分岐します。
AI エージェント実行のための Git ブランチを再トレースします。すべての実行を記録し、壊れたステップからフォークし、出荷前に修正を証明します。

## Original Extract

Record, replay, fork & share AI agent executions. See every LLM call, tool invocation, and error your agent makes — then debug and iterate in seconds. Free for 1,000 traces/mo.

Retrace records your AI agents runs so you can replay them step by step, fork from any point to a fix, and share the result as a link

Retrace — The Execution Replay Engine for AI Agents — Retrace retrace Product Record One decorator captures every LLM call Fork & Replay Branch from any decision, replay the future Prove the Fix Re-run a fix and get a verdict — did it work? Eval Gates CI/CD quality gates that block bad deploys Guardrails Runtime safety — halt agents before they fail Enforcement Circuit breakers — block runaway loops & budget blow-outs Failure Taxonomy MAST — why agent runs fail, auto-classified Sessions Multi-agent tracing with causal graphs Pricing Playground Docs Changelog Sign in Get started CI for AI agent behavior
Know why your AI agent failed. Prove the fix.
Retrace records every LLM call, tool call, and decision your agent makes — so you can replay any run, fork from the exact step it broke, and prove the fix before you ship.
| npm i retrace-sdk Works with OpenAI · Anthropic · Gemini — any LLM, any framework. // the core loop
Record what your agent did. Replay it, fork it, fix it — and prove the fix.
From a single decorator to a verdict — the reliability loop Retrace owns, end to end.
One decorator captures every LLM call, tool call, and decision — so a production failure becomes a permanent regression test.
Re-run any recorded run, or fork from the exact step it broke — so you debug the real cause, not the symptom.
Change the step that broke — prompt, tool input, or model — and re-run the fork, so you see the corrected path before you ship it.
Re-run a change against the failed run and get a verdict — so you ship the fix, not hope.
↻ every run feeds the next — record, fix, repeat.
A real recording. Scrub it yourself.
The actual trace UI — a recorded incident-RCA run that stopped at the OOM symptom and was marked failed, then forked from that step and traced the real cause: a deploy that raised the batch size 50×. Scrub the timeline, open any span, watch the fork diverge.
forked at step 4 → re-ran failed run → verdict: fix verified retrace.ai.generate 6.18s · $0.0001
retrace.ai.generate 5.83s · $0.0005 err
↳ retrace.ai.generate 2.79s · $0.0003 ok
↳ retrace.ai.generate 2.01s · $0.0001 ok
6 spans · 1 error at step 4 · forked at step 4 · fix verified
Interactive — a real trace exported from the product (incident-rca · INC-4835). Spans, costs and timings are the recorded values.
A full reliability platform — not a dashboard.
Detect failures, enforce limits, evaluate quality, and understand multi-agent systems — on every recorded run.
Flag groundedness gaps, statistical drift, failure clusters, and MAST failure types automatically — so you learn why a run failed, before a user does, not just that it did.
Halt a run at a budget, loop, or step limit, and block a bad action before it runs via a pre-call gateway with hold-for-approval — so a bug can't quietly cascade into a cloud bill.
Evaluate 0 3 evaluate Evaluate
Quality you can gate on — evaluations, auto eval-rules, datasets, and prove-the-fix verdicts, plus a CI gate so a regression fails the build instead of reaching users.
Understand 0 4 understand Understand
See the whole system — sessions, agent topology, agent memory, semantic search, prompt versioning, and shareable tapes — so a multi-agent failure isn't a black box.
Flag groundedness gaps, statistical drift, failure clusters, and MAST failure types automatically — so you learn why a run failed, before a user does, not just that it did.
Halt a run at a budget, loop, or step limit, and block a bad action before it runs via a pre-call gateway with hold-for-approval — so a bug can't quietly cascade into a cloud bill.
Quality you can gate on — evaluations, auto eval-rules, datasets, and prove-the-fix verdicts, plus a CI gate so a regression fails the build instead of reaching users.
See the whole system — sessions, agent topology, agent memory, semantic search, prompt versioning, and shareable tapes — so a multi-agent failure isn't a black box.
Observability shows you what broke. Retrace re-runs it, forks the failed step, and proves the fix.
Aggregate metrics & dashboards
Replay the exact run — fork from the failed step and cascade-re-execute, so you debug the real cause, not a chart
Semantic search — describe the bug in plain language and jump to the failing run in seconds
No way to test a fix without re-running everything
Prove-the-fix — re-run a change against the failed run and get a verdict before you ship
Guardrails + circuit breakers that halt a run when it blows a budget or loops — so a bug can't quietly run up your bill
Runtime enforcement — block a bad action before it runs, not diagnose it after
One decorator — any Python/TS agent (LangChain · CrewAI · LlamaIndex), no lock-in
From zero to your first trace in ~2 minutes.
Sign in with GitHub, copy ~3 lines, run your agent — the first trace streams in live. No infrastructure to manage.
Install the SDK and add one decorator. Calls to OpenAI, Anthropic and Gemini are captured automatically.
npm i retrace-sdk ↳ Framework-agnostic — works with LangChain, CrewAI, and LlamaIndex.
1 2 3 4 5 6 import retrace retrace . configure ( api_key = "rt_..." ) @retrace.record ( name = "my-agent" ) def run_agent ( prompt ) : return agent . invoke ( prompt ) SDK connected classify_intent recorded
Run your agent. Every LLM call, tool call, cost and error streams onto the timeline as it happens.
↳ Or watch it replay step-by-step in the dashboard.
1 retrace traces tail REC · live classify_intent 0.4s fetch_context 1.8s generate_response 2.4s
Fork from the exact step that broke, change the input, re-run — then prove the fix actually worked.
↳ verify-fix returns a verdict — improved, regressed, or unchanged.
1 2 3 retrace forks create --trace < id > --span < id > --input "grounded prompt" retrace forks replay < id > --wait retrace traces verify -fix < id > generate_response error ↳ fork ok verify-fix → fix verified
Ship AI agents with tests, not hope.
A prompt edit, tool change, or model upgrade can quietly break multi-step behavior. Retrace turns each production failure into a regression test and runs it as an eval gate on every PR — the build fails when behavior breaks. It diffs each run against a golden baseline, so you catch “relevance dropped vs last release,” not just a red or green dot.
1 2 3 4 - name : Eval Gate run : retrace eval gate --evaluation $ EVAL_ID --trace $ TRACE_ID --threshold 0.8 env : RETRACE_API_KEY : $ { { secrets . RETRACE_API_KEY } } Checks ✗ retrace / eval-gate behavior regressed — build failed ✓ retrace / eval-gate passed 5/5 runs — merge unblocked How the eval gate works → the closed loop
0 1 production trace a real run fails
0 2 captured failure becomes a regression test
0 3 eval gate on the PR re-runs that test
0 4 green check behavior holds — merge
what it gates on — that unit tests can't
The build fails when behavior breaks — so that exact failure is harder to ship again. (retrace eval gate exits 1.)
Eval gates and replays run on your model account.
Add your Google (Gemini) API key in Settings and the eval-gate judge and every server-side replay — fork, cascade, and prove-the-fix — call the model through your key, so the tokens are billed to your own provider account. The key is validated on save, encrypted at rest (AES-256-GCM), shown only as its last four characters, and never returned again. Remove it any time and Retrace falls back to the platform key.
Google · key ••••a1b2 validated ✓ what your key powers
→ eval-gate judge scores every PR run
→ fork & cascade replay re-executes the agent
→ prove-the-fix computes the verdict
Validated against the provider on save — a bad key is rejected, never stored.
Encrypted at rest with AES-256-GCM and a per-secret derived key.
Surfaced only as its last 4 characters — never logged, never returned.
Remove it and replays fall back to the platform key with no downtime.
No credit card required. Upgrade when you need more traces or AI requests.
+ Sandbox env: 5,000 traces/mo
+ Sandbox env: 50,000 traces/mo
+ Unlimited prove-the-fix runs/mo
+ Sandbox env: Unlimited traces/mo
Under 2 minutes. Install the SDK, add one decorator, and traces stream immediately. No infrastructure to manage.
Python and TypeScript SDKs with auto-instrumentation for OpenAI, Anthropic, and Google Gemini. Works with any agent framework — LangChain, CrewAI, Vercel AI SDK, AutoGen, LlamaIndex.
Select any span in a trace, modify its input, and Retrace cascade-replays from that point forward. Context from the fork flows into subsequent LLM calls. You get a side-by-side diff with cost and latency deltas.
Runtime policies that monitor your agent in real-time. Set cost budgets, loop detection, context overflow limits, or latency caps. When a limit is crossed, the agent receives a HALT command — so a runaway loop or budget blow-out stops at the limit instead of running up your bill.
TLS in transit, encrypted at rest. API keys are SHA-256 hashed. PII auto-redaction runs on every plan as a security baseline. Tenant isolation is enforced at the application layer — every query is scoped per user and backed by a guardrail regression test.
Yes. The eval gate endpoint (POST /evaluations/:id/gate) returns pass/fail against a threshold. The CLI command `retrace eval gate` exits with code 1 on failure — perfect for GitHub Actions.
LangSmith focuses on tracing and observability. Retrace adds interactive fork & cascade-replay from any step, runtime guardrails that halt runaway agents, groundedness detection, and prove-the-fix verification.
Yes. Each span carries an agent id, sessions group multi-turn conversations, and an agent topology graph shows cross-agent ordering and inter-agent failure modes.
Stop guessing.
Start replaying.
Your agent failed 3 steps before the error surfaced. Fork from the real cause — not the symptom.
retrace Git branching for AI agent execution — record every run, fork from the step that broke, and prove the fix before you ship.
