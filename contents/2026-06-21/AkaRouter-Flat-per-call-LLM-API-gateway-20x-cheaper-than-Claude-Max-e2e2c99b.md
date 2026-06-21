---
source: "https://akarouter.dev"
hn_url: "https://news.ycombinator.com/item?id=48620183"
title: "AkaRouter – Flat per-call LLM API gateway (20x cheaper than Claude Max)"
article_title: "AkaRouter — AI Inference Gateway"
author: "mrdedatn"
captured_at: "2026-06-21T16:31:43Z"
capture_tool: "hn-digest"
hn_id: 48620183
score: 1
comments: 0
posted_at: "2026-06-21T16:19:59Z"
tags:
  - hacker-news
  - translated
---

# AkaRouter – Flat per-call LLM API gateway (20x cheaper than Claude Max)

- HN: [48620183](https://news.ycombinator.com/item?id=48620183)
- Source: [akarouter.dev](https://akarouter.dev)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T16:19:59Z

## Translation

タイトル: AkaRouter – コールごとにフラットな LLM API ゲートウェイ (Claude Max より 20 倍安い)
記事のタイトル: AkaRouter — AI 推論ゲートウェイ
説明: OpenAI 互換の AI 推論ゲートウェイ。すべてのモデルに 1 つのキー。

記事本文:
AkaRouter — AI 推論ゲートウェイの機能 モデル 価格 テーマの切り替え サインイン 開始する [SYS_STATUS]: GATEWAY_ONLINE_READY CLAUDE MAX 20X
$200/月 $20/月
同じOpus 4.8。同じプロンプト サイズ制限。 91% 安くなります。リクエストごとに支払います。 API キーは 1 つです。あらゆるフロンティアモデル。
* サインアップごとに、1 ポイント モデルで 100 無料ポイントが付属します (1 日あたり 10 ポイントが上限)。クレジットカードはありません。
openaiインポートからOpenAI
クライアント = OpenAI(
Base_url="https://api.akarouter.dev/v1",
api_key="akar_your_key_here"
)
応答 = client.chat.completions.create(
モデル="ステップ-37-フラッシュ",
messages=[{"role": "user", "content": "Hello AkaRouter!"}]
) Opus 4.8 コールあたり $0.08 11 のフロンティア モデル 0¢ コンテキスト長の不安 1 つの KEY 統合 API 残酷な数学 なぜ同じモデルに 5 倍から 90 倍も支払うのでしょうか?
同じ上流モデル。同じプロンプトが表示されます。同じ応答品質。 AkaRouter は、大手企業が使用しているのと同じプロバイダーにあなたをルーティングします。ただし、50 倍に値上げするわけではありません。
マルチプロバイダー (Anthropic + OpenAI + Google) 通話ごとに定額 (トークン計算なし) 5,000 または 200,000 のプロンプト混合でも同価格 無料のフロンティア モデルが含まれる 1 つのキー、すべてのモデル クロードのみ OpenAI のみ OpenAI 互換 (任意のクライアント) ヘビー クロード コード ユーザー クロード Max 20x: $200/月 OpenRouter ダイレクト: $540/月 AkaRouter Ultra: $50/月 750 Opus 4.8通話 + 無料フロンティア + 無制限の格安モデル。 Max 20x より 91% 安い。
1 つのキーで 250 オーパス + 500 ソネット + 5000 以上の格安通話。 2 つのサブスクリプションを置き換えます。
すべてを 1 つのゲートウェイ経由でルーティングします。同じワークロードで小売 API の支出が 87% オフになります。
理にかなった請求モデル リクエストごとに支払う。トークンごとではありません。
ほとんどの API は 100 万トークンごとに料金がかかります。料金は通話ごとに定額です。プロンプトが 500 ワードであっても 200,000 トークンであっても同じ料金です。
トークンごとの支払い (OpenRouter、直接 API) 5,000 トークン プロンプト $0.025 50,000 トークン プロンプト (中程度のクロード コード) $0.25 200,000 トークン プロンプト (フル)

context Opus) $1.00+ 同じ質問、コストの差異は 40 倍 コンテキストにより多くのコードを詰め込むたびに、より多くの料金を支払うことになります。すべての長いドキュメント。すべての大きなリポジトリのクローン。形だけの不安は本物だ。
コールごとのフラット、すべてのサイズ 5,000 トークン プロンプト 10 ポイント ($0.08) 50,000 トークン プロンプト 10 ポイント ($0.08) 200,000 トークン プロンプト 10 ポイント ($0.08) 同じ質問、同じ価格 プロンプトがモデルのコンテキスト ウィンドウに収まる限り、同じ料金を支払います。 Opus 4.8 は 200K トークンに適合します。全部使ってください。
トークン計算なし 入力/出力トークンの分割を計算しません。リクエストの前に毎回コストを見積もらないでください。モデルに電話するだけです。
フルコンテキストを使用する コードベース全体を詰め込みます。10 個の PDF をドロップします。電卓を使わずに 200K Opus ウィンドウ全体を使用します。
予測可能な請求額は、100 Opus コール = 8 ドルです。いつも。月曜日も同じ、日曜日も同じ。驚くべき超過はありません。
* コールごとの価格は、プロンプトがモデルの文書化されたコンテキスト ウィンドウ内に収まる限り適用されます。限界に達しましたか？リクエストを分割するか、より大きなウィンドウを備えたモデルにアップグレードしてください。
すべてのモデルにおける価格設定の完全な透明性。各ポイントのコスト。
隠し階層はありません。 「プレミアム」マークアップはありません。全メニューを実際に支払う価格で。
プロ プランには 2,500 ポイント/月が付属しています。 Ultra は 7,500 個で出荷されます。自由に組み合わせて組み合わせることができます。モデルのロックはありません。
高可用性と低遅延の推論ワークロードを実現するためにゼロから構築されました。
スマートなロード バランシング リアルタイムのヘルス重み付けと動的な実行中の同時実行追跡​​を備えたラウンドロビン ルーティング。
高可用性フェイルオーバー 自動リクエスト再試行およびホットスワップ ルーティング。ルーティング ターゲットがダウンした場合、トラフィックはすぐに再割り当てされます。
クォータ管理 サブスクリプション層の詳細なレート制限、スライド式トークン予算、API キーごとに記録されるコスト分析。
単一の API キーを通じてすべてのモデルにアクセスできます。トークンごとおよびリクエストごとの課金をサポートします。
すべて トークンごと リクエストごと 任意の t

ext ビジョン コード 11 モデル中 10 を表示 |トークンごとに 0 |リクエストごとに 11 クロード ソネット 4.6 claude-sonnet-46
リクエストごとの Anthropic 200K ctx チャット完了強力なコーディングと推論を備えたバランスのとれたクロードのバリアント。
Per-Request OpenAI 1M ctx チャット完了 強力な総合性能を備えたフラッグシップ GPT モデル。
リクエストごとの OpenAI 922K ctx チャット完了 最大の機能と推論を備えた最上位の GPT モデル。
リクエストごとの NVIDIA 262K ctx テキスト 速度と推論のために最適化されたフロンティア オープンソース LLM。
リクエストごと その他 100 万 ctx チャット完了 強力な推論機能とコード機能を備えたフロンティアのクローズドソース LLM。
リクエストごとの Anthropic 200K ctx チャット完了 日常業務に適した、高速かつ手頃な価格のクロード バリアント。
リクエストごとの Google 100 万 ctx チャット完了 Google のフラッグシップ Pro モデルで、強力な根拠を備えています。
リクエストごとの OpenAI 272K ctx チャット完了コード 生成とリファクタリング用に最適化されたコードに特化したバリアント。
リクエストごとの Stepfun 256K ctx テキスト ビジョン 幅広い汎用機能を備えた超高速会話モデル。
リクエストごとの MiniMax 1M ctx テキスト 1M トークン コンテキスト モデルは、長い形式の推論とエージェント ワークフロー向けに最適化されています。
実稼働スループットの需要に合わせて、柔軟なサブスクリプション プランを選択してください。
スターター + クロード ソネット 4.6、GPT-5.4、Gemini 3.1 Pro
スターター + Owl Alpha、Nemotron Ultra、ステップ 3.7 フラッシュ
Pro + GPT-5.5、GPT-5.3 コーデックスパーク、クロード オーパス 4.8
* カスタム制限またはエンタープライズ設定が必要ですか? Telegram サポート グループ (t.me/akarouter_support) に参加してください。私たちは電子メールではなく、グループのみを行っています。
© 2026 AkaRouter by HAMDI.カーネル: v1.0.0-STABLE。

## Original Extract

OpenAI-compatible AI inference gateway. One key, every model.

AkaRouter — AI Inference Gateway Features Models Pricing Toggle theme Sign in Get Started [SYS_STATUS]: GATEWAY_ONLINE_READY CLAUDE MAX 20X
$200/MO $20/MO
Same Opus 4.8. Same prompt size limits. 91% cheaper. Pay-per-request. One API key. Every frontier model.
* Every signup ships with 100 free points on 1pt models (10/day cap). No credit card.
from openai import OpenAI
client = OpenAI(
base_url="https://api.akarouter.dev/v1",
api_key="akar_your_key_here"
)
response = client.chat.completions.create(
model="step-37-flash",
messages=[{"role": "user", "content": "Hello AkaRouter!"}]
) $0.08 per Opus 4.8 call 11 frontier models 0¢ context-length anxiety 1 KEY unified API the brutal math Why pay 5x–90x more for the same model?
Same upstream models. Same prompts. Same response quality. AkaRouter routes you to the same providers the big guys use — we just don't mark it up 50x.
Multi-provider (Anthropic + OpenAI + Google) Flat per-call (no token math) Same price for 5K or 200K prompt mixed Free frontier model included One key, every model Claude only OpenAI only OpenAI-compatible (any client) Heavy Claude Code user Claude Max 20x: $200/mo OpenRouter direct: $540/mo AkaRouter Ultra: $50/mo 750 Opus 4.8 calls + free frontier + unlimited cheap models. 91% cheaper than Max 20x.
250 Opus + 500 Sonnet + 5000+ cheap calls in ONE key. Replace two subscriptions .
Route everything through ONE gateway. 87% off retail API spend at the same workload.
the billing model that makes sense Pay per request. Not per token.
Most APIs charge per million tokens. We charge per call — flat. Same price whether your prompt is 500 words or 200,000 tokens.
Pay-per-token (OpenRouter, direct APIs) 5K-token prompt $0.025 50K-token prompt (medium Claude Code) $0.25 200K-token prompt (full context Opus) $1.00+ Same question, 40x cost variance Every time you stuff more code into context, you pay more. Every long doc. Every large repo clone. Token anxiety is real.
Flat per-call, every size 5K-token prompt 10 pts ($0.08) 50K-token prompt 10 pts ($0.08) 200K-token prompt 10 pts ($0.08) Same question, same price As long as your prompt fits in the model's context window , you pay the same. Opus 4.8 fits 200K tokens. Use them all.
No token math Don't calculate input/output token splits. Don't estimate cost before every request. Just call the model.
Use full context Stuff the whole codebase in. Drop in 10 PDFs. Use the full 200K Opus window without a calculator.
Predictable bills 100 Opus calls = $8. Always. Same on Monday, same on Sunday. No surprise overages.
* Per-call pricing applies as long as your prompt fits within the model's documented context window. Hit the limit? Split your request — or upgrade to a model with a bigger window.
full pricing transparency Every model. Every point cost.
No hidden tiers. No "premium" markups. The whole menu, at the price you'll actually pay.
Pro Plan ships with 2,500 points/month. Ultra ships with 7,500. Mix and match freely — no model locking.
Built from the ground up for high availability and low-latency inference workloads.
SMART LOAD BALANCING Round-robin routing with real-time health weighting and dynamic in-flight concurrency tracking.
HIGH-AVAILABILITY FAILOVER Automatic request retry and hot-swap routing. If a routing target goes down, traffic is immediately re-allocated.
QUOTA MANAGEMENT Granular subscription tier rate limits, sliding token budgets, and cost analytics logged per API key.
All models accessible through a single API key. Supports per-token and per-request billing.
All Per-Token Per-Request Any text vision code Showing 10 of 11 models | 0 per-token | 11 per-request Claude Sonnet 4.6 claude-sonnet-46
Per-Request Anthropic 200K ctx chat completion Balanced Claude variant with strong coding and reasoning.
Per-Request OpenAI 1M ctx chat completion Flagship GPT model with strong general performance.
Per-Request OpenAI 922K ctx chat completion Top-tier GPT model with maximum capability and reasoning.
Per-Request NVIDIA 262K ctx text Frontier open-source LLM optimized for speed and reasoning.
Per-Request Other 1M ctx chat completion Frontier closed-source LLM with strong reasoning and code capabilities.
Per-Request Anthropic 200K ctx chat completion Fast, affordable Claude variant for everyday tasks.
Per-Request Google 1M ctx chat completion Google flagship Pro model with strong reasoning.
Per-Request OpenAI 272K ctx chat completion code Code-specialized variant optimized for generation and refactoring.
Per-Request Stepfun 256K ctx text vision Ultra-fast conversational model with broad general capability.
Per-Request MiniMax 1M ctx text 1M token context model optimized for long-form reasoning and agentic workflows.
Choose a flexible subscription plan that matches your production throughput demands.
Starter + Claude Sonnet 4.6, GPT-5.4, Gemini 3.1 Pro
Starter + Owl Alpha, Nemotron Ultra, Step 3.7 Flash
Pro + GPT-5.5, GPT-5.3 Codex Spark, Claude Opus 4.8
* Need custom limits or an enterprise setup? Join our Telegram support group at t.me/akarouter_support — we don't do email, just the group.
© 2026 AkaRouter by HAMDI. Kernel: v1.0.0-STABLE.
