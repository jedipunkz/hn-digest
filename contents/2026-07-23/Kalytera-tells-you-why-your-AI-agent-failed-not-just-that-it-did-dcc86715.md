---
source: "https://kalytera.dev/"
hn_url: "https://news.ycombinator.com/item?id=49028390"
title: "Kalytera – tells you why your AI agent failed, not just that it did"
article_title: "Kalytera — Find where your agent fails"
author: "mathurpriya19"
captured_at: "2026-07-23T21:57:30Z"
capture_tool: "hn-digest"
hn_id: 49028390
score: 1
comments: 0
posted_at: "2026-07-23T21:37:16Z"
tags:
  - hacker-news
  - translated
---

# Kalytera – tells you why your AI agent failed, not just that it did

- HN: [49028390](https://news.ycombinator.com/item?id=49028390)
- Source: [kalytera.dev](https://kalytera.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T21:37:16Z

## Translation

タイトル: Kalytera – AI エージェントが失敗したことだけでなく、なぜ失敗したかを説明します
記事のタイトル: Kalytera — エージェントの失敗箇所を見つける

記事本文:
カリテラ
ドキュメント
価格設定
研究
無料で始めましょう →
✕
無料プラン
アカウントを作成する
10,000 セッション/月、永久無料。クレジットカードは必要ありません。
名前または会社名
仕事用メール
アカウントを作成 →
無料プラン
API キー
今すぐコピーしてください。再度表示されなくなります。
APIキー
kly_live_...
コピー
これを大切に保管してください。これにより、エージェントが Kalytera に対して認証されます。紛失した場合は、priya@kalytera.ai に電子メールを送信して交換してください。
完了 — 次のステップを表示 →
✓
準備は完了です
キーが保存されました。 2 分未満で最初にスコアが付けられたインタラクション:
1
pip インストール kalytera
2
kalytera.configure(api_key= "kly_live_..." ) と kalytera.trace() をエージェントに追加します
3
ダッシュボードを開く — 最初のトレースから 30 秒以内にスコアが表示されます
ダッシュボードを開く →
クイックスタート ドキュメントを読む
↗
チェックアウトへのリダイレクト
支払いを完了するには、Stripe にリダイレクトされます。 API キーはすでに無料枠で有効です。
実稼働エージェントの LLM 評価
エージェントの失敗をキャッチする
ユーザーが行う前に。
すべてのステップを自動的に評価し、障害の背後にあるパターンを明らかにし、痕跡を掘り下げることなく、わかりやすい根本原因を取得します。
追加するのは 1 行です。 LangChain、CrewAI、AutoGen、または任意のカスタム フレームワークで動作します。 10,000 セッション/月は無料。
品質レイヤーがないということは、障害が静かに増大することを意味します 既存のツールはトラフィックをサンプリングし、最終出力のみをチェックします。ワークフローの途中で発生した障害は、最後に痕跡を残しません。開発者は顧客からの苦情から問題を発見し、何時間もかけてトレースを行って、原因となったステップを特定します。
企業は ROI を測定できますが、ROI の原動力は測定できません。 実験フレームワークは、エージェントの導入によって成果が向上したかどうかを示すことができます。しかし、品質レイヤーがなければ、どの失敗パターンがパフォーマンスを低下させたのかは誰にも分かりません。全体的なリターンが目に見えてわかります。原因はそうではありません。
一般的な合否スコア

エージェントが実際に求めているものと一致しない ヘルスケア エージェントとマーケティング エージェントを同じ方法で評価するべきではありません。ほとんどの評価ツールでは、ユースケースにとって品質が実際に何を意味するかを設定する方法がなく、1 つの一般的なスコアが提供されます。
エージェントが同じ間違いを繰り返す次のインタラクション 次のフェーズ エージェントが再び同じステップに陥る前に、失敗についてエージェントに警告するメカニズムはありません。これを修正するには再トレーニングが必要です。そのためには、まだ持っていないラベル付きデータが必要です。
本日生産中です
AI エージェントは通常のソフトウェアとは異なる失敗をします。それを捕捉するには、最終出力だけでなくすべてのステップを監視する、別の種類の可観測性が必要です。
ほとんどのチームは Kalytera に接続し、何日も黙って実行されていた障害を最初の 1 時間以内に発見します。
30 日後、状況はより鮮明になります。ツールが変わったからではなく、損失パターンが現れるまでに時間がかかったからです。
ユーザーから情報を得るのはやめましょう。
追加するのは 1 行です。失敗パターンは 30 秒以内に表面化します。月間 10,000 セッションまでは無料です。
評価されたセッションに対して支払います。座席ライセンスはありません。驚くべき請求はありません。無料利用枠は永久に有効です。さらに必要な場合はアップグレードしてください。
Kalytera の背後にある障害分類法と評価方法論は公開されています。この方法論は研究であり、独自のロックインではありません。
Kalytera は、サンプルや事後ではなく、すべてのインタラクションをあらゆるステップでリアルタイムに評価し、何が、どこで、そしてなぜ壊れたのかを正確に教えてくれます。
Enterprise AI エージェントは、複雑な複数ステップのワークフローを実行します。従来のソフトウェアとは異なり、エージェントはすべての対話において異なる考え方と行動をします。標準の品質チェックはそのために構築されていません。 Eval プラットフォームと可観測性プラットフォームは存在しますが、それらはサンプリングされたデータと事後分析を使用します。ワークフロー途中の障害（大惨事となる可能性のある 1 回限りの障害も含む）を見逃す

ロフィック。
すべてのインタラクションをすべてのステップで、それが発生した瞬間に評価します
すべてのインタラクションに、業界に合わせて調整された品質スコアを提供します
失敗フィードで 1 回限りの致命的な障害を即座に明らかにします
繰り返される失敗を、単純な根本原因を含む名前付きパターンにグループ化します。
pip インストールから最初のスコアリングされたインタラクションまで 2 分以内に完了します。
$ pip インストール kalytera
2. 設定する
Agent.py インポート kalytera
カリテラ。設定(
api_key= "kly_live_..." ,
api_endpoint= "https://api.kalytera.dev" ,
Agent_id= "請求エージェント" 、 # オプション
)
3. エージェントを監視する
エージェント.py @ kalytera.watch
def support_agent (user_input):
... # 他には何も変わりません
4. または各ステップで手動でトレースします
Agent.py カリテラ。トレース (
session_id= "セッション-001" ,
ステップ番号= 1 、
step_name= "classify_intent" ,
入力=ユーザーメッセージ、
出力=エージェント_応答、
)
決してブロックせず、trace() を呼び出すこともありません。ただちに戻ります。 Kalytera に到達できない場合、エージェントは実行を続けます。失敗した送信はローカル キューから自動的に再試行されます。
← 前へ はじめに 次へ → 品質スコアリング
品質スコアリング
すべてのインタラクションには、LLM 審査員が各ステップで評価する 4 つの側面から計算された 0 ～ 100 の品質スコアが与えられます。
total_score ≥ 0.7 の場合、インタラクションは合格します。ダッシュボードからエージェントごとに設定可能。
医療エージェントとマーケティング エージェントを同じ方法で評価すべきではありません。 Kalytera は、業界ごとに調整された品質の重み付けを備えて出荷されます。
失敗したすべてのインタラクションは、7 つの失敗タイプのいずれかに分類されます。
3 つのビュー。それぞれが 1 つの質問に答えます。つまり、エージェントは健康ですか、現在何が問題になっているのか、そしてこの 1 つのインタラクションで正確に何が起こったのかです。
7 日間にわたる品質スコアの傾向、今日の合格率、アクティブな失敗数、および頻度順にランク付けされた上位 3 つの失敗タイプ。
すべての失敗はそれが起こった瞬間に起こります。一回限りの

失敗は個別に表示されます。 5 つ以上の失敗が同じタイプとステップを共有し、根本原因が 1 文になると、繰り返しパターンが自動的にグループ化されます。
完全なステップバイステップのトレース。ステップごとの品質スコア。失敗したステップが強調表示されます。失敗の理由は、JSON ダンプではなく、平易な英語の 1 文で記述します。
Kalytera を認証し、デフォルトを設定します。起動時に、trace() を呼び出す前に 1 回呼び出します。
カリテラ。設定(
API_key: str 、
api_endpoint: str = "http://localhost:8000" ,
エージェント ID: str |なし = なし、
) -> なし
api_key str Kalytera API キー (kly_live_...)。 KALYTERA_API_KEY 環境変数からも読み取ります。
api_endpoint str Kalytera サーバーの URL。 KALYTERA_API_ENDPOINT 環境変数からも読み取ります。
エージェントID str |なし ダッシュボード内のエージェントのオプションの安定した名前 (例: 「billing-agent」)。
← 前へ 3 つのビュー 次へ → @watch
APIリファレンス
@kalytera.watch
あらゆるエージェント機能を計測するデコレータ。入力、出力、レイテンシーを自動的にキャプチャします。
@kalytera.watch
def your_agent (user_input: str ) -> str :
...
あらゆる Python 呼び出し可能なもの (LangChain チェーン、CrewAI エージェント、AutoGen エージェント、またはプレーン関数) で動作します。
インタラクションの 1 つのステップを手動で記録します。エージェントを @watch でラップできない場合に使用します。
カリテラ。トレース (
session_id: str 、
ステップ番号: int 、
ステップ名: str 、
入力: str 、
出力: str 、
ツールコール: リスト |なし = なし、
メタデータ: dict |なし = なし、
session_ended: bool = False 、
) -> なし
← 前へ @watch 次へ → データモデル
APIリファレンス
データモデル
テーブルが4つ。 Kalytera が保管するものはすべて、それ以外は何もありません。

## Original Extract

Kalytera
Docs
Pricing
Research
Get started free →
✕
Free plan
Create your account
10,000 sessions/month, free forever. No credit card required.
Name or company
Work email
Create account →
Free plan
Your API key
Copy it now — it won't be shown again.
API key
kly_live_...
Copy
Store this securely. It authenticates your agent to Kalytera. If you lose it, email priya@kalytera.ai for a replacement.
Done — show me next steps →
✓
You're all set
Your key is saved. First scored interaction in under 2 minutes:
1
pip install kalytera
2
Add kalytera.configure(api_key= "kly_live_..." ) and kalytera.trace() to your agent
3
Open the dashboard — scores appear within 30 seconds of your first trace
Open dashboard →
Read the quickstart docs
↗
Redirecting to checkout
You'll be redirected to Stripe to complete payment. Your API key is already active on the free tier.
LLM evaluation for production agents
Catch agent failures
before your users do.
Evaluate every step automatically, surface the patterns behind failures, and get a plain-English root cause — without digging through traces.
One line to add. Works with LangChain, CrewAI, AutoGen, or any custom framework. Free for 10K sessions/month.
No quality layer means failures compound silently Existing tools sample traffic and check only the final output. Failures that start mid-workflow leave no trace at the end. Developers find out from customer complaints, then spend hours in traces to find which step caused it.
Enterprises can measure ROI — but not what's driving it Experiment frameworks can show whether agent deployment improved outcomes. But without a quality layer, nobody knows which failure patterns dragged performance down. The overall return is visible. The causes are not.
Generic pass/fail scores don't match what your agent is actually for A healthcare agent and a marketing agent should not be graded the same way. Most eval tools give you one generic score with no way to configure what quality actually means for your use case.
Your agent repeats the same mistake next interaction Next phase There's no mechanism to warn an agent about a failure before it hits the same step again. Fixing it requires retraining — which needs labeled data you don't have yet.
In production today
AI agents fail differently than normal software. You need a different kind of observability to catch it — one that watches every step, not just the final output.
Most teams connect Kalytera and find failures within the first hour that had been running silently for days.
After 30 days the picture gets sharper — not because the tool changed, but because the loss patterns have had time to emerge.
Stop finding out from your users.
One line to add. Failure patterns surface in 30 seconds. Free for 10,000 sessions per month.
Pay for sessions evaluated. No seat licenses. No surprise bills. Free tier works forever — upgrade when you need more.
The failure taxonomy and evaluation methodology behind Kalytera are published openly. The methodology is research, not proprietary lock-in.
Kalytera evaluates every interaction at every step, in real time — not a sample, not after the fact — and tells you exactly what broke, where, and why.
Enterprise AI agents run complex multi-step workflows. Unlike traditional software, agents think and act differently in every interaction — standard quality checks aren't built for that. Eval and observability platforms exist, but they use sampled data and after-the-fact analysis. They miss failures mid-workflow, including one-off failures that could be catastrophic.
Evaluates every interaction at every step, the moment it happens
Gives every interaction a quality score calibrated to your industry
Surfaces one-off catastrophic failures immediately in the Failure Feed
Groups repeating failures into named patterns with plain English root cause
From pip install to your first scored interaction in under two minutes.
$ pip install kalytera
2. Configure
agent.py import kalytera
kalytera. configure (
api_key= "kly_live_..." ,
api_endpoint= "https://api.kalytera.dev" ,
agent_id= "billing-agent" , # optional
)
3. Watch your agent
agent.py @ kalytera.watch
def support_agent (user_input):
... # nothing else changes
4. Or trace manually at each step
agent.py kalytera. trace (
session_id= "session-001" ,
step_number= 1 ,
step_name= "classify_intent" ,
input=user_message,
output=agent_reply,
)
Never blocks, never raises trace() returns immediately. If Kalytera is unreachable, your agent keeps running. Failed sends are retried automatically from a local queue.
← Previous Introduction Next → Quality scoring
Quality scoring
Every interaction gets a quality score from 0 to 100, computed from four dimensions evaluated by an LLM judge on every step.
An interaction passes when overall_score ≥ 0.7 . Configurable per agent from the dashboard.
A healthcare agent and a marketing agent shouldn't be graded the same way. Kalytera ships with quality weightings tuned per industry.
Every failed interaction is classified into one of seven failure types.
Three views. Each answers one question: is my agent healthy, what's failing right now, and exactly what happened in this one interaction.
Quality score trend over 7 days, today's pass rate, active failure count, and top 3 failure types ranked by frequency.
Every failure the moment it happens. One-off failures shown individually. Repeating patterns grouped automatically once 5+ failures share the same type and step, with a one-sentence root cause.
Full step-by-step trace. Per-step quality score. The failing step highlighted. Failure reason in one plain English sentence — not a JSON dump.
Authenticates Kalytera and sets defaults. Call once at startup before any trace() calls.
kalytera. configure (
api_key: str ,
api_endpoint: str = "http://localhost:8000" ,
agent_id: str | None = None ,
) -> None
api_key str Your Kalytera API key (kly_live_...). Also reads from KALYTERA_API_KEY env var.
api_endpoint str URL of your Kalytera server. Also reads from KALYTERA_API_ENDPOINT env var.
agent_id str | None Optional stable name for your agent in the dashboard (e.g. "billing-agent").
← Previous The three views Next → @watch
API reference
@kalytera.watch
A decorator that instruments any agent function. Captures input, output, and latency automatically.
@ kalytera.watch
def your_agent (user_input: str ) -> str :
...
Works with any Python callable — LangChain chains, CrewAI agents, AutoGen agents, or a plain function.
Manually log one step of an interaction. Use when your agent can't be wrapped with @watch.
kalytera. trace (
session_id: str ,
step_number: int ,
step_name: str ,
input: str ,
output: str ,
tool_calls: list | None = None ,
metadata: dict | None = None ,
session_ended: bool = False ,
) -> None
← Previous @watch Next → Data model
API reference
Data model
Four tables. Everything Kalytera stores, nothing more.
