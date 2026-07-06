---
source: "https://d1k3.com"
hn_url: "https://news.ycombinator.com/item?id=48810289"
title: "Show HN: Dike is a compliance gateway for AI products in the EU"
article_title: "Dike – Compliance framework for teams shipping LLMs"
author: "orbanlevi"
captured_at: "2026-07-06T21:22:05Z"
capture_tool: "hn-digest"
hn_id: 48810289
score: 1
comments: 0
posted_at: "2026-07-06T20:46:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dike is a compliance gateway for AI products in the EU

- HN: [48810289](https://news.ycombinator.com/item?id=48810289)
- Source: [d1k3.com](https://d1k3.com)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T20:46:11Z

## Translation

タイトル: Show HN: Dike は EU における AI 製品のコンプライアンス ゲートウェイです
記事のタイトル: Dike – LLM を出荷するチームのためのコンプライアンス フレームワーク
説明: 1 つのベース URL 変更でチャットボットまたは RAG パイプラインを EU AI 法に準拠させるゲートウェイ プロキシ。監査レベルのログ記録、監視記録、インシデント報告。

記事本文:
Dike – LLM を出荷するチームのためのコンプライアンス フレームワーク
✓ 生成 · gpt-4o · 封印された a3f1… ✓ 取得 · 5 つのドキュメント · 封印された 91bc… ✓ human_review · 承認済み · 封印された d04e… ✓ 編集 · pii · 封印された 5e9f… ✓ インシデント · 報告された · 封印された 7c2a… ✓ エクスポート · 証拠 · 封印された b60d… 製品を出荷するチーム向け
監査人が要求するすべてのもの。チームが構築する必要があるものは何もありません。
Dike は、EU における AI 製品のコンプライアンス ゲートウェイです。 LLM トラフィックを当社のプロキシ経由でルーティングし、ベース URL を 1 回変更するだけで、監査レベルのロギング、人間による監視記録、インシデント レポートを取得できます。
コンプライアンスはオプションではありませんが、それを自分で構築するには、出荷ゼロのエンジニアリングに数か月かかる必要があります。
製品。それが私たちがあなたから外す部分です。
ベース URL の変更が 1 つあります。
その間にあるものはすべて証拠です。
OpenAI API 呼び出しは、OpenAI 互換の任意のデバイスに向かう途中で、Dike ゲートウェイを経由してルーティングされます。
プロバイダー。 PII は編集され、悪意のある通話はブロックされ、すべてのリクエストが監査になります
完全なトレースを記録します。
監査人が要求するすべてのもの。
チームが構築する必要があるものは何もありません。
AI 法は善意に報いるものではありません。それは証拠に報酬を与えます。 Dike はあなたの AI システムを変えます
日々の業務を自動的に証拠として記録します。
すべてのプロンプト、取得、および完了は、封印されたハッシュ チェーンされた監査レコードになります。構造による改ざん防止: 監査人は、何も変更または削除されていないことを暗号的に検証できます。
モデルが回答するために実際に使用した参照文書、つまり単なるチャットの記録ではなく、第 12 条 (2) が要求する正確な証拠を記録します。
承認、編集、上書きは第一級のイベントです。規制当局が人間が AI をどのように監視するか (第 14 条) と尋ねるとき、あなたは答えを書くのではなくエクスポートします。
重大なインシデントが発生すると、15 日間の報告時間が開始されます (

第73条）。ダイクは事件を開始し、時計を追跡し、市場監視当局への報告書の草案を作成します。
PII は、ストレージに何かが書き込まれる前に、ゲートウェイで編集されます。ストレージは EU のみであり、保持期間は法定最低 6 か月から始まり、完全に構成可能です。
設計によるフェールオープン: 監査ストレージにアクセスできない場合でも、ゲートウェイはリクエストをモデル プロバイダーに渡します。ミリ秒未満のオーバーヘッド、ストリーミングをサポート。
Dike は、モデル プロバイダーの前にあるゲートウェイ プロキシです。既存のクライアントに当社の拠点を案内してください
URL とすべての完了は、ストリーミングかどうかにかかわらず、改ざんが明らかな監査レコードになります。 SDK はありません、いいえ
コードが変更されます。
BaseURL を 1 つ変更するだけで、インストールするものは何もありません
OpenAI 互換の SDK またはフレームワークで動作します
RAG パイプラインとチャットボットはすぐにサポートされます
監査ストレージがダウンしている場合でもリクエストは通過します
# OpenAI 互換の SDK。ベース URL をゲートウェイにポイントする
# エンドポイント名をモデルとして使用します。それだけです
openaiインポートからOpenAI
クライアント = OpenAI (
Base_url= "https://gw.dike.dev/v1" ,
api_key=環境。 DIKE_KEY 、
)
応答 = client.chat.completions。作成(
model= "support-bot-v2" , # エンドポイント名
メッセージ=メッセージ、
) レイテンシーバジェットを代理します
気付かないだろう。
ゲートウェイは Rust に組み込まれており、応答が到着するとそれをストリーミングします。コンプライアンスの実行
クリティカルパス上ではなく、飛行中 — 堤防を通過することは、通過することと区別がつきません
直接。
監査人が見ているもの。目に見えるもの。
堤防ダッシュボード: すべてのイベントは封印され、すべての出来事は時計に表示され、すべてのエクスポートはワンクリックで行われます。
✓ 訴訟が開始され、証拠が自動添付される (38 件の記録)
○ 法的承認を待っており、m.kovacs に割り当てられています
📄 第 12 条ログ抜粋 6 月 1 日 – 6 月 30 日 · 2.4 GB · 署名済み
📄 人間による監視の概要 第 14 条・214 件のレビューイベント
📄 技術文書

イオンスケルトン Annex IV · プレフィルド
📄 保持と編集ポリシー GDPR · EU ストレージ証明書
シンプルで透明性のある価格設定。
使用した分だけ支払います。
AI のトレースと通話を追跡します。無料で始めて、成長に合わせて拡張します。
無制限のトレース、カスタムニーズ
苦労から数分で監査準備完了まで
1 つのベース URL を変更し、LLM トラフィック ルートが EU プロキシを経由するようにします。同日発送。
ハッシュチェーンされたイベントは PII 編集され、法律で要求される限り EU 内に保持されます。
ワンクリック: ログの抽出、技術文書のスケルトン、インシデントケースファイル。
競合他社よりも先にコンプライアンスを遵守しましょう。
私たちは、EU AI チームの小グループをクローズド ベータに参加させています。ベータ期間中は無料で、
設立チームによる統合の強力な支援。
スパムはありません。招待の準備ができたときに 1 通のメールが送信され、開始時に 1 通のメールが送信されます。
© 2026 ダイク。 EU製🇪🇺

## Original Extract

A gateway proxy that makes your chatbot or RAG pipeline EU AI Act compliant with one base-URL change. Audit-grade logging, oversight records, incident reporting.

Dike – Compliance framework for teams shipping LLMs
✓ generation · gpt-4o · sealed a3f1… ✓ retrieval · 5 docs · sealed 91bc… ✓ human_review · approved · sealed d04e… ✓ redaction · pii · sealed 5e9f… ✓ incident · reported · sealed 7c2a… ✓ export · evidence · sealed b60d… for teams who ship products
Everything the auditor will ask for. Nothing your team has to build.
Dike is a compliance gateway for AI products in the EU. Route your LLM traffic through our proxy and get audit-grade logging, human-oversight records and incident reporting, with a single base-URL change.
Compliance isn't optional, but building it yourself means months of engineering that ships zero
product. That's the part we take off your plate.
One base-URL change.
Everything in between is evidence.
Your OpenAI API calls route through the Dike gateway on their way to any OpenAI-compatible
provider. PII is redacted, malicious calls are blocked, and every request becomes an audit
record with full traces.
Everything the auditor will ask for.
Nothing your team has to build.
The AI Act doesn't reward good intentions. It rewards evidence. Dike turns your AI system's
day-to-day operation into the evidence, automatically.
Every prompt, retrieval and completion becomes a sealed, hash-chained audit record. Tamper-evident by construction: auditors can cryptographically verify that nothing was altered or deleted.
We record which reference documents your model actually used to answer: the exact evidence Article 12(2) asks for, not just a chat transcript.
Approvals, edits and overrides are first-class events. When the regulator asks how humans supervise your AI (Article 14), you export the answer instead of writing it.
Serious incidents start a 15-day reporting clock (Article 73). Dike opens the case, tracks the clock and drafts the report to your market surveillance authority.
PII is redacted at the gateway, before anything is written to storage. Storage is EU-only, retention starts at the 6-month legal minimum and is fully configurable.
Fail-open by design: if audit storage is unreachable, the gateway still passes your requests through to the model provider. Sub-millisecond overhead, streaming supported.
Dike is a gateway proxy in front of your model provider. Point your existing client at our base
URL and every completion, streaming or not, becomes a tamper-evident audit record. No SDK, no
code changes.
One baseURL change, nothing to install
Works with any OpenAI-compatible SDK or framework
RAG pipelines and chatbots supported out of the box
Requests pass through even if audit storage is down
# Any OpenAI-compatible SDK. Point the base URL at the gateway
# and use your endpoint name as the model. That's it
from openai import OpenAI
client = OpenAI (
base_url= "https://gw.dike.dev/v1" ,
api_key=env. DIKE_KEY ,
)
response = client.chat.completions. create (
model= "support-bot-v2" , # your endpoint name
messages=messages,
) A proxy your latency budget
won't even notice.
The gateway is built in Rust and streams responses through as they arrive. Compliance runs
in-flight, not in your critical path — going through Dike is indistinguishable from going
direct.
What the auditor sees. What you see.
The Dike dashboard: every event sealed, every incident on a clock, every export one click.
✓ Case opened, evidence auto-attached (38 records)
○ Awaiting legal sign-off, assigned to m.kovacs
📄 Article 12 log extract Jun 1 – Jun 30 · 2.4 GB · signed
📄 Human oversight summary Article 14 · 214 review events
📄 Technical documentation skeleton Annex IV · pre-filled
📄 Retention & redaction policy GDPR · EU storage attestation
Simple, transparent pricing.
Pay for what you use.
Track AI traces and calls. Start free, scale as you grow.
Unlimited traces, custom needs
From struggle to audit-ready in minutes
Change one base URL and your LLM traffic routes through our EU proxy. Ship the same day.
Hash-chained events, PII-redacted, retained in the EU for as long as the law requires.
One click: log extracts, technical documentation skeletons and incident case files.
Get compliant before your competitors do.
We're onboarding a small group of EU AI teams into the closed beta: free during beta, with
white-glove integration help from the founding team.
No spam. One email when your invite is ready, one when we launch.
© 2026 Dike. Made in the EU 🇪🇺
