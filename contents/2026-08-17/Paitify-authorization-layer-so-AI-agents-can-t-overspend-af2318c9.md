---
source: "https://paitify.io"
hn_url: "https://news.ycombinator.com/item?id=49327033"
title: "Paitify – authorization layer so AI agents can't overspend"
article_title: "Paitify — Agent Spend Policy Engine"
image: "https://paitify.io/og-image.png"
author: "rahmankapucu"
captured_at: "2026-08-17T06:34:02Z"
capture_tool: "hn-digest"
hn_id: 49327033
score: 2
comments: 0
posted_at: "2026-08-17T06:05:30Z"
tags:
  - hacker-news
  - translated
---

# Paitify – authorization layer so AI agents can't overspend

- HN: [49327033](https://news.ycombinator.com/item?id=49327033)
- Source: [paitify.io](https://paitify.io)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T06:05:30Z

## Translation

タイトル: Paitify – AI エージェントが過剰な支出を防ぐための認証レイヤー
記事のタイトル: Paitify — エージェント支出ポリシー エンジン
説明: AI エージェントに対するリアルタイムの支出承認とポリシーの適用。エージェントが支出できる金額を事前承認、監査、制御します。

記事本文:
Paitify — エージェント支出ポリシー エンジンを読み込み中…
リアルタイム認証 AI エージェントはガードレールを使用します。
Paitify は、自律型 AI エージェント用の支出ポリシー エンジンです。エージェントの速度を低下させることなく、詳細なルールを定義し、リアルタイムでトランザクションを承認し、完全な監査証跡を維持します。
API を参照してください。 auto_approve < $500 · human_approval ≥ $500 ·hard_cap $3,000 金額 (USD) $ Merchant Authorize 金額を入力し、[Authorize] をクリックすると、実際の決定が表示されます。
90 秒: AI エージェントが 49 ドルで自動承認され、600 ドルで停止する様子を観察します。
AI エージェントはお金を使うことができます。たくさんのお金。
最新の LLM エージェントは、インフラストラクチャを参照、購入、スピンアップし、有料 API を呼び出します。多くの場合、自律的に 24 時間年中無休で行われます。一元的な支出管理がなければ、設定を誤った単一のエージェントが、誰にも気づかれないうちに予算を使い果たしたり、コンプライアンス ポリシーに違反したり、未承認の販売者と取引したりする可能性があります。
従来の支払い管理は人間向けに設計されていました。 Paitify はマシン向けに構築されており、リアルタイムの決定、コードとしてのポリシー、および支払いレイヤーが自宅に電話することなく確認できるマシン可読 JWT トークンを備えています。
エージェントが何を購入し、どのようなレールで支払ったとしても、Paitify はポリシーと監査層としてイエスかノーを決定し、それを記録します。
部門の予算に基づいてソフトウェア、消耗品、またはクラウド リソースを購入するエージェント。
従量制 API、LLM トークンを呼び出すコーディング エージェント、またはハード シーリングに対する計算。
エージェントは、1 日の予算やキャンペーンの予算の上限内で入札を調整したり、キャンペーンを開始したりします。
ポジションと値の制限内で取引を実行するエージェントが、再構築のためにログに記録されます。
各支払いの前に 1 回の API 呼び出し。ポリシーが施行されました。すべての決定が記録されます。
エージェントごとまたは会社全体で、ダッシュボードで支出制限、販売者ルール、承認しきい値を設定します。
エージェントは前に 1 つの API 呼び出しを行います

支払いごとに。 Paitify は理由を付けてリアルタイムで承認または拒否します。
すべての決定は自動的に記録されます。ダッシュボードで完全な監査証跡を確認します。
インポートリクエスト
応答 = リクエスト.post(
"https://api.paitify.io/v1/integration/authorize",
headers={"X-API-Key": PAITIFY_API_KEY},
json={
"エージェントId": "調達エージェント-01",
「金額」: 249.99、
"通貨": "USD",
"販売者名": "AWS",
"mccコード": "7372",
}
)
結果 = 応答.json()
# result["state"] は、決定の信頼できる唯一の情報源です。
if result["state"] == "AUTHORIZED":
チャージカード(金額=249.99)
confirm_capture(result["authorizationId"]) # POST /v1/integration/authorize/{id}/capture
elif result["状態"] == "捕捉":
Charge_card(amount=249.99) # AUTO キャプチャ モード: すでに記録されています — スキップ /capture
elif result["state"] == "REQUIRES_APPROVAL":
wait_for_approval(result["authorizationId"]) # ポーリング GET /v1/integration/authorizations/{id}
それ以外の場合:
print(f"支出が拒否されました: {result['reasonCode']}") 実稼働 AI システム用に構築
すべての機能は、自律エージェントを大規模に展開するという実際の運用上の課題に合わせて設計されています。
Redis を利用した支出カウンターとメモリ内ルールの評価。バッチ遅延はありません。すべての承認はライブでチェックされます。
自動: 承認時にすぐにコミットを消費し、失敗した場合にのみ /cancel を呼び出します。明示的: AUTHORIZED で予約、/capture でコミット、/cancel で解放。
クロードおよびあらゆる MCP 互換クライアントですぐに使用できます。
すべての決定は、完全なルール評価結果、エージェント ID、販売者データ、およびタイムスタンプとともに記録されます。
高額な取引にフラグを立てて人間の承認を得る。エージェントは待ちます。人間が決める。予算は解決されるまで保持されます。
承認は RSA-2048 署名付きトークンです。決済処理業者は JWKS を使用してオフラインで検証します。コールバックの遅延はありません。
エージェントごとまたは会社全体。金額・期間による制限

、クライアント センター コード、通貨、速度、販売者、営業時間。
スライディング ウィンドウごとのレート制限トランザクション。カードの制限に達する前に、エージェントのループを停止します。
REST API。あらゆるスタック、あらゆる言語、あらゆる支払いプロセッサで動作します。
すべての決定は連鎖し、証明可能です。
Paitify が行うすべての承認 (承認、拒否、人間によるレビューのために送信) は、決定自体と同じトランザクションで追加専用の監査ログに書き込まれます。
各エントリはその前のエントリにチェーンされるため、レコードが改ざんされたり欠落したりするとチェーンが切断され、すぐに表示されます。論理的に削除可能な行はありません。それは台帳です。
Audit_log — 追加専用エントリ 0000 … 0000 ← 0000 … 0000 ← ジェネシス 0000…0000 エントリ 0000 … 0000 ← 0000 … 0000 ← ジェネシス 0000…0000 エントリ 0000 … 0000 ← 0000 … にチェーン0000 ← ジェネシス 0000…0000 シンプルで透明性の高い価格設定
無料で始めましょう。エージェントの成長に合わせて拡張します。
個人およびサイドプロジェクト向け。
最初のエージェントを導入する小規模チーム向け。
より高い制限やカスタム要件が必要ですか?
エージェントにガードレールを設置する準備はできていますか?
無料で始めましょう。クレジットカードは必要ありません。最初の 1,000 件の承認は私たちが行います。
Paitify 自律型 AI エージェント用の支出ポリシー エンジン。

## Original Extract

Real-time spend authorization and policy enforcement for AI agents. Pre-authorize, audit, and control what your agents can spend.

Paitify — Agent Spend Policy Engine Loading…
Real-time authorization Your AI agents spend with guardrails.
Paitify is the spend policy engine for autonomous AI agents. Define granular rules, authorize transactions in real time, and maintain a complete audit trail — without slowing your agents down.
See the API auto_approve < $500 · human_approval ≥ $500 · hard_cap $3,000 Amount (USD) $ Merchant Authorize Enter an amount and hit Authorize to see a live decision.
90 seconds: watch an AI agent get auto-approved at $49 — and stopped at $600.
AI agents can spend money. A lot of money.
Modern LLM agents browse, purchase, spin up infrastructure, and call paid APIs — often autonomously, 24/7. Without centralized spend controls, a single misconfigured agent can drain budgets, violate compliance policies, or transact with unauthorized merchants before anyone notices.
Traditional payment controls were designed for humans. Paitify is built for machines — real-time decisions, policy as code, and machine-readable JWT tokens your payment layer can verify without calling home.
Whatever your agent buys, and whatever rail it pays on, Paitify is the policy and audit layer that decides yes or no — and logs it.
An agent buying software, supplies, or cloud resources against a department budget.
A coding agent calling metered APIs, LLM tokens, or compute against a hard ceiling.
An agent adjusting bids or launching campaigns within daily and campaign budget caps.
An agent executing trades within position and value limits, logged for reconstruction.
One API call before every payment. Policies enforced. Every decision logged.
Set spend limits, merchant rules, and approval thresholds in the dashboard — per agent or company-wide.
Your agent makes one API call before each payment. Paitify approves or denies in real time, with a reason.
Every decision is logged automatically. Review the full audit trail in your dashboard.
import requests
response = requests.post(
"https://api.paitify.io/v1/integration/authorize",
headers={"X-API-Key": PAITIFY_API_KEY},
json={
"agentId": "procurement-agent-01",
"amount": 249.99,
"currency": "USD",
"merchantName": "AWS",
"mccCode": "7372",
}
)
result = response.json()
# result["state"] is the sole source of truth for the decision.
if result["state"] == "AUTHORIZED":
charge_card(amount=249.99)
confirm_capture(result["authorizationId"]) # POST /v1/integration/authorize/{id}/capture
elif result["state"] == "CAPTURED":
charge_card(amount=249.99) # AUTO capture mode: already recorded — skip /capture
elif result["state"] == "REQUIRES_APPROVAL":
wait_for_approval(result["authorizationId"]) # poll GET /v1/integration/authorizations/{id}
else:
print(f"Spend denied: {result['reasonCode']}") Built for production AI systems
Every feature designed for the real operational challenges of deploying autonomous agents at scale.
Redis-backed spend counters and in-memory rule evaluation. No batch delay — every authorization is checked live.
AUTO: spend commits immediately on authorize, call /cancel only on failure. EXPLICIT: reserve with AUTHORIZED, commit with /capture, or release with /cancel.
Works with Claude and any MCP-compatible client out of the box.
Every decision logged with full rule evaluation results, agent ID, merchant data, and timestamps.
Flag high-value transactions for human approval. Agents wait; humans decide; budget holds until resolved.
Approvals are RSA-2048 signed tokens. Payment processors verify offline using JWKS — no callback latency.
Per-agent or company-wide. Limits by amount, period, MCC code, currency, velocity, merchant, and business hours.
Rate-limit transactions per sliding window. Stop looping agents before they hit your card limits.
REST API. Works with any stack, any language, any payment processor.
Every decision, chained and provable.
Every authorization Paitify makes — approved, denied, or sent for human review — is written to an append-only audit log in the same transaction as the decision itself.
Each entry chains to the one before it, so a tampered or missing record breaks the chain and shows immediately. Nothing is a soft-deletable row — it's a ledger.
audit_log — append-only entry 0000 … 0000 ← chained to 0000 … 0000 ← genesis 0000…0000 entry 0000 … 0000 ← chained to 0000 … 0000 ← genesis 0000…0000 entry 0000 … 0000 ← chained to 0000 … 0000 ← genesis 0000…0000 Simple, transparent pricing
Start free. Scale as your agents grow.
For individuals and side projects.
For small teams deploying their first agents.
Need higher limits or custom requirements?
Ready to put guardrails on your agents?
Start for free. No credit card required. Your first 1,000 authorizations are on us.
Paitify The spend policy engine for autonomous AI agents.
