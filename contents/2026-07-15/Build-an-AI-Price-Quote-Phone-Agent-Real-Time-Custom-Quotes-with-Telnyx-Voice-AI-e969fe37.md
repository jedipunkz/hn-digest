---
source: "https://lowlatencyclub.ai/blog/posts/ai-price-quote-phone-agent-python"
hn_url: "https://news.ycombinator.com/item?id=48919737"
title: "Build an AI Price Quote Phone Agent Real-Time Custom Quotes with Telnyx Voice AI"
article_title: "Build an AI Price Quote Phone Agent — Real-Time Custom Quotes with Telnyx Voice AI | Low Latency Club"
author: "harpreetseehra"
captured_at: "2026-07-15T13:13:41Z"
capture_tool: "hn-digest"
hn_id: 48919737
score: 1
comments: 0
posted_at: "2026-07-15T12:23:35Z"
tags:
  - hacker-news
  - translated
---

# Build an AI Price Quote Phone Agent Real-Time Custom Quotes with Telnyx Voice AI

- HN: [48919737](https://news.ycombinator.com/item?id=48919737)
- Source: [lowlatencyclub.ai](https://lowlatencyclub.ai/blog/posts/ai-price-quote-phone-agent-python)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T12:23:35Z

## Translation

タイトル: Telnyx Voice AI を使用した AI 価格見積り電話エージェントのリアルタイム カスタム見積りの構築
記事のタイトル: AI 価格見積電話エージェントの構築 — Telnyx Voice AI を使用したリアルタイムのカスタム見積 |低遅延クラブ
説明: 電話番号に電話し、必要なものを説明すると、明細項目と月次合計を含む AI 生成の価格見積もりがリアルタイムで表示されます。 Python は 102 行。

記事本文:
メインコンテンツにスキップ
イベント
動画
ポッドキャスト
統合
リソース
製品
お問い合わせ
Slack に参加してください
ブログ
コードファーストの技術的な詳細
用語集
音声 AI と通信用語
開発者ドキュメント
API リファレンスと統合ガイド
GitHub
オープンソースのリポジトリと SDK
記事
チュートリアルとサポート記事
AIエージェント
音声 AI エージェントを構築する
音声API
プログラム可能な音声通話
クローハウス
AIを活用した音声アシスタント
クロードトーク
会話型AIプラットフォーム
ブログ
コードファーストの技術的な詳細
用語集
音声 AI と通信用語
開発者ドキュメント
API リファレンスと統合ガイド
GitHub
オープンソースのリポジトリと SDK
記事
チュートリアルとサポート記事
AIエージェント
音声 AI エージェントを構築する
音声API
プログラム可能な音声通話
クローハウス
AIを活用した音声アシスタント
クロードトーク
会話型AIプラットフォーム
AI 価格見積り電話エージェントを構築する — Telnyx Voice AI を使用したリアルタイムのカスタム見積り
顧客があなたの会社に電話して見積もりを求めてきたと想像してください。 AI エージェントは、顧客を保留にしたり、営業に転送したり、コールバックを約束したりするのではなく、応答を受け取り、3 ～ 4 つの条件付きの質問をし、数量と単価を含む項目の見積もりを作成し、通話中に合計金額を提示します。フォームも待ち時間も摩擦もありません。
これは AI Price Quote Phone Agent です。Telnyx Call Control と AI Inference で構築された 102 行の Python アプリです。コールセンターもCRM統合も事前構築された価格設定エンジンもありません。 AI は製品カタログと価格を把握し、適切な質問をして、構造化された見積もりをリアルタイムで生成します。
このチュートリアルでは、これを最初から構築します。リポジトリのクローンを作成し、電話番号を構成し、数分でデプロイします。
誰でもカスタム価格の見積もりを得るために電話できる電話番号:
発信者がダイヤルイン — Telnyx が応答して挨拶します
AI が何を必要としているかを尋ねる「どのようなコミュニケーション サービスがあるか」

探していますか？」
会話 - 電話をかけてきた人がニーズを説明し、AI が量を見積もるために 3 ～ 4 つのフォローアップの質問をします。
見積書の配信 - AI が明細項目と月次合計を集計します
構造化抽出 — 電話を切ると、AI が完全な見積書を JSON として抽出します (品目、数量、単価、小計、月次合計、メモ)
保存された見積 — GET /quotes エンドポイント経由でアクセス可能
インタラクション全体は音声主導で行われます。Text-to-Speech が各質問と最後の引用を読み上げ、発信者は自然な音声で応答します。 AI モデル (Telnyx AI 推論経由の Llama 3.3 70B) が会話と見積もりの生成を処理し、通話制御が電話を処理します。
ほとんどの価格見積もりフローは静的です。Web フォームに記入し、販売メールを待ちます。これは動的であり、AI が発信者の発言に基づいて質問を調整するリアルタイムの会話です。音声通話と SMS が必要ですか? AI は予想されるボリュームを尋ねます。フリーダイヤル番号が必要なだけですか? AI は無関係な質問をスキップし、直接見積もりに進みます。
また、チャットを超えたパターンも示しています。AI は単に会話するだけではなく、会話から構造化データを抽出します。通話が終了すると、アプリは「見積もりを JSON として抽出」というプロンプトとともに完全なトランスクリプトをモデルに送り返します。明細項目と月次合計を含むクリーンで機械可読な見積オブジェクトを取得し、すぐに保存したり、電子メールで送信したり、請求システムにパイプしたりできます。
残高が入金された Telnyx アカウント
音声が有効になっている Telnyx 電話番号
Webhook URL で構成された通話制御アプリケーション
ローカルサーバーを Telnyx Webhook に公開するための ngrok
電話
│
▼
Telnyx 通話制御 (Webhook イベント)
│
▼
Flask アプリ (app.py、102 行)
│
§──► 通話開始 → 電話に出る
§──► 電話応答 → TTS gree

ティン
§──► call.gather.ended → 音声 → AI推論 → TTS応答
§──► call.speak.ended → 次の入力を収集
└──► call.hangup → 引用符をJSONとして抽出 → 保存
│
▼
Telnyx AI 推論 (Llama 3.3 70B)
│
▼
構造化見積 (項目を含む JSON)
このアプリは、Telnyx Webhook イベントによって駆動されるステート マシンです。各イベントは、応答する、話す、集まる、電話を切るなどの次のアクションをトリガーします。 AI 推論呼び出しは、呼び出し元と会話して見積もりを作成する作業と、通話終了後に構造化された JSON 見積もりを抽出する作業の 2 つのジョブを処理します。
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-price-quote-phone-agent-python
cp .env.example .env
pip install -r 要件.txt
資格情報を使用して .env を編集します。
TELNYX_API_KEY=KEY019C8F8C3D268774F8F560946B984D9E_...
TELNYX_PUBLIC_KEY=...
AI_MODEL=メタ-ラマ/ラマ-3.3-70B-命令
QUOTE_NUMBER=+13105551234
ステップ 2: コードを理解する
すべては app.py (102 行) にあります。各部分の動作は次のとおりです。
AI には、ユニットごとの価格が記載された製品カタログが組み込まれています。
PRICING = {"voice_ minutes": 0.005、"sms_messages": 0.004、"local_numbers": 1.00、"toll_free_numbers": 2.00、
「sip_trunking_channel」: 2.50、「ai_inference_1k_tokens」: 0.01、「cloud_storage_gb」: 0.05、「fax_pages」: 0.07}
これはシステム プロンプトに挿入されるため、AI はユニットごとに請求すべき金額を正確に認識します。
SYSTEM_PROMPT = f「あなたは価格設定のスペシャリストです。利用可能な製品と単位あたりの価格設定: {json.dumps(PRICING)}。発信者に必要なものを尋ね、数量を見積もり、見積もりを作成します。回答は 2 文以内にまとめてください。要件 (通常は 3 ～ 4 つの質問) を収集した後、品目と月次合計を含めて見積もりを要約します。」
プロンプトはコントロールサーフェスです。製品カタログを交換したり、質問を変更したりできます

スタイルを変更したり、要約する前に AI が尋ねる質問の数を調整したりできます。
Webhook ハンドラーはコア ステート マシンです。各 Telnyx イベントは次のアクションをトリガーします。
イベントタイプ == "call.initiated" かつ p.get("direction") == "incoming" の場合:
active_calls[ccid] = {"呼び出し元": p.get("from"), "会話": [{"役割": "システム", "コンテンツ": SYSTEM_PROMPT}], "開始": time.time()}
client.calls.actions.answer(ccid)
elifevent_type == "call.answered":
client.calls.actions.speak(ccid, payload="こんにちは! 今すぐカスタム見積もりを作成できます。どのような種類のコミュニケーション サービスをお探しですか?", voice="女性", language_code="en-US")
elifevent_type == "call.speak.ended" で次のように呼び出します。
client.calls.actions.gather(ccid, input_type="speech", end_silence_timeout_secs=2, timeout_secs=20, language_code="en-US")
elifevent_type == "call.gather.ended" で次のように呼び出します。
speech = p.get("スピーチ", {}).get("結果", "")
call["会話"].append({"役割": "ユーザー", "内容": スピーチ})
応答 = call_inference(call["会話"])
call["会話"].append({"役割": "アシスタント", "コンテンツ": 応答})
client.calls.actions.speak(ccid, payload=response, voice="女性", language_code="en-US")
会話のループ: 話す→集まる→推測する→話す。 AI はターンごとに完全な会話履歴を確認するため、すべてのやり取りのコンテキストを維持します。
def call_inference(messages, max_tokens=250):
resp = request.post(INFERENCE_URL, headers={"Authorization": f"Bearer {TELNYX_API_KEY}", "Content-Type": "application/json"},
json={"モデル": AI_MODEL、"メッセージ": メッセージ、"max_tokens": max_tokens、"温度": 0.5}、タイムアウト=15)
resp.raise_for_status()
return resp.json()["選択"][0]["メッセージ"]["コンテンツ"]
Telnyx AI 推論に対する OpenAI 互換の単純なチャット完了呼び出し。温度0.5の短所

一貫性 — 引用は創造的ではなく、予測可能であることが必要です。
ここが興味深い部分です。発信者が電話を切ると、アプリは別のプロンプトを使用して会話全体を AI に送り返します。構造化された引用を抽出します。
elif イベントタイプ == "call.hangup":
call = active_calls.pop(ccid, なし)
if call と len(call.get("conversation", [])) > 3:
extract = [{"role": "system", "content": "価格見積を抽出します。返される JSON: line_items ({product、quantity、unit_price、小計} のリスト)、monthly_total (数字)、notes (文字列)。"},
{"role": "user", "content": chr(10).join(f"{m['role']}: {m['content']}" for m in call["conversation"] if m["role"] != "system")}]
quote = json.loads(call_inference(extract, max_tokens=400))
quote["発信者"] = call["発信者"]
quote["タイムスタンプ"] = time.strftime("%Y-%m-%dT%H:%M:%SZ")
quotes.append(引用)
AI は、品目、数量、単価、小計、月次合計を含むクリーンな JSON を返します。アプリは発信者の番号とタイムスタンプを添付して保存します。
@app.route("/quotes",methods=["GET"])
def list_quotes():
return jsonify({"quotes": quotes[-20:]}), 200
GET /quotes を押すと、最後の 20 個の引用符が JSON として表示され、CRM、請求システム、または分析ダッシュボードにパイプする準備が整います。
Python app.py
別のターミナルで、ngrok を使用してローカル サーバーを公開します。
ngrok http 5000
HTTPS URL をコピーし、Telnyx ポータルで設定します。
通話制御アプリケーションに移動します
アプリケーションを作成または編集する
Webhook URL を https://<your-ngrok-url>.ngrok.app/webhooks/voice に設定します。
Telnyx 電話番号をまだ割り当てていない場合は、この通話制御アプリケーションに割り当てます。
どの電話からでも Telnyx 番号に電話をかけます。次のように聞こえます:
「こんにちは！ 今すぐカスタム見積もりを作成させていただきます。どのような通信サービスをお探しですか?」
次のようなことを言います。「必要があります

約 50 の電話番号、月間 10,000 分の音声通話、数千の SMS メッセージ。」
AI は、SMS の数、フリーダイヤル番号、SIP トランキングが必要かなどのフォローアップの質問をし、品目と月次合計を含む見積もりを要約します。
電話を切った後、抽出された見積もりを確認してください。
カール http://localhost:5000/quotes | python3 -m json.tool
次のような構造化された JSON オブジェクトが表示されます。
{
"line_items": [
{"製品": "ローカル番号", "数量": 50, "単価": 1.00, "小計": 50.00},
{"製品": "音声分", "数量": 10000, "単価": 0.005, "小計": 50.00},
{"製品": "sms_messages"、"数量": 5000、"単価": 0.004、"小計": 20.00}
]、
"monthly_total": 120.00、
"notes": "お客様は、第 3 四半期に SMS の量が増加する可能性があると指摘しました。",
"発信者": "+12125551234",
"タイムスタンプ": "2026-07-15T14:30:00Z"
}
エージェントのカスタマイズ
価格カタログとシステム プロンプトはコントロール サーフェスです。小さな変更により、まったく異なるエージェントが生成されます。
PRICING = {"consulting_hour": 200.00、"audit_report": 2500.00、"retainer_monthly": 8000.00、"training_session": 1500.00}
質問スタイルを変更します。
SYSTEM_PROMPT = f"あなたは価格設定のスペシャリストです。利用可能な製品: {json.dumps(PRICING)}。一度に 1 つずつ質問してください。会話的でありながら効率的であること。正確に 3 つの質問を行った後、項目と月ごとの合計を含む見積もりを要約します。電話をかけてきた人が割引を求めてきた場合は、標準の価格帯について説明してください。」
割引ロジックを追加します。
SYSTEM_PROMPT = f"あなたは価格設定のスペシャリストです。製品: {json.dumps(PRICING)}。 10,000 ユニットを超えるボリュームの場合は、15% 割引が適用されます。 50,000 ユニットを超える場合は 25% が適用されます。概要に割引について記載してください。」
電話後に SMS 経由で見積もりを送信します。
# 引用文を抽出したらSMSで送信
Telnyx からのインポートメッセージ
quote_text = f"あなたの見積もり: {quote['monthl

y_total']}/月。 {len(quote['line_items'])} 件の項目。」
Message.create(to=call["caller"], from_=QUOTE_NUMBER, text=quote_text)
このパターン (製品カタログを使用した会話型 AI + 通話後の構造化抽出) は、あらゆる見積もりや注文のフローに機能します。
この例では、簡単にするためにインメモリ ストレージを使用します。運用環境の場合:
データベース - メモリ内の引用リストを PostgreSQL または Redis に置き換え、再起動後も引用が存続するようにします。
同時実行 — 複数のワーカーを使用して gunicorn の背後でアプリを実行します
エラー回復 - 再試行または SMS フォールバックにより、推論タイムアウトと呼び出し失敗を適切に処理します。
認証 — /quotes エンドポイントに API キー検証を追加します
Webhook 検証 — アプリはすでに Telnyx Ed25519 署名を検証しています。公開キーが最新であることを確認してください
プロンプト調整 - 製品カタログのさまざまなシステム プロンプトと温度設定をテストします。
レート制限 - Webhook エンドポイントを悪用から保護します
モニタリング — 構造化されたログ記録と通話の成功/失敗率に関するアラートを追加します。
CRM 統合 — 抽出された見積もりを Salesforce、HubSpot、または請求システムにパイプします
別の AI モデルを使用できますか?はい。 Telnyx AI Inference は OpenAI と互換性があります。 .env の AI_MODEL をプラットフォームで利用可能な任意のモデルに設定します。 Llama 3.3 70B Instruct は、構造化引用符のデフォルトとして適切です。
通話料金はいくらですか?テルニクス

[切り捨てられた]

## Original Extract

Call a number, describe what you need, and get an AI-generated price quote with line items and monthly total in real time. 102 lines of Python.

Skip to main content
Events
Videos
Podcast
Integrations
Resources
Products
Contact us
Join our Slack
Blog
Code-first technical deep dives
Glossary
Voice AI & telecom terminology
Developer Docs
API references & integration guides
GitHub
Open-source repos & SDKs
Articles
Tutorials & support articles
AI Agents
Build voice AI agents
Voice API
Programmable voice calls
ClawHouse
AI-powered voice assistant
ClawdTalk
Conversational AI platform
Blog
Code-first technical deep dives
Glossary
Voice AI & telecom terminology
Developer Docs
API references & integration guides
GitHub
Open-source repos & SDKs
Articles
Tutorials & support articles
AI Agents
Build voice AI agents
Voice API
Programmable voice calls
ClawHouse
AI-powered voice assistant
ClawdTalk
Conversational AI platform
Build an AI Price Quote Phone Agent — Real-Time Custom Quotes with Telnyx Voice AI
Imagine a customer calls your business and asks for a quote. Instead of putting them on hold, transferring to sales, or promising a callback, an AI agent picks up — asks 3–4 qualifying questions, builds a line-item quote with quantities and unit prices, and delivers the total on the call. No forms, no waiting, no friction.
This is the AI Price Quote Phone Agent — a 102-line Python app built with Telnyx Call Control and AI Inference. No call center, no CRM integration, no pre-built pricing engine. The AI has your product catalog and pricing, asks the right questions, and generates a structured quote in real time.
In this walkthrough, you'll build it from scratch. Clone the repo, configure a phone number, and deploy in minutes.
A phone number that anyone can call to get a custom price quote:
Caller dials in — Telnyx answers and greets them
AI asks what they need — "What kind of communication services are you looking for?"
Conversation — caller describes their needs, AI asks 3–4 follow-up questions to estimate quantities
Quote delivered — AI summarizes with line items and monthly total
Structured extraction — on hangup, AI extracts the full quote as JSON (line items, quantities, unit prices, subtotals, monthly total, notes)
Quote stored — accessible via GET /quotes endpoint
The whole interaction is voice-driven: Text-to-Speech reads each question and the final quote, and the caller responds with natural speech. The AI model (Llama 3.3 70B via Telnyx AI Inference) handles the conversation and quote generation, and Call Control handles the telephony.
Most price quote flows are static: fill out a web form, wait for a sales email. This one is dynamic — a real-time conversation where the AI adapts its questions based on what the caller says. Need voice minutes and SMS? The AI asks about expected volume. Just want a toll-free number? The AI skips irrelevant questions and gets straight to the quote.
It also demonstrates a pattern that goes beyond chat: the AI doesn't just converse — it extracts structured data from the conversation. After the call ends, the app sends the full transcript back to the model with a "extract the quote as JSON" prompt. You get a clean, machine-readable quote object with line items and a monthly total — ready to store, email, or pipe into your billing system.
A Telnyx account with funded balance
A Telnyx phone number with voice enabled
A Call Control Application configured with your webhook URL
ngrok for exposing your local server to Telnyx webhooks
Phone Call
│
▼
Telnyx Call Control (webhook events)
│
▼
Flask app (app.py, 102 lines)
│
├──► call.initiated → answer the call
├──► call.answered → TTS greeting
├──► call.gather.ended → speech → AI Inference → TTS response
├──► call.speak.ended → gather next input
└──► call.hangup → extract quote as JSON → store
│
▼
Telnyx AI Inference (Llama 3.3 70B)
│
▼
Structured quote (JSON with line items)
The app is a state machine driven by Telnyx webhook events. Each event triggers the next action — answer, speak, gather, or hang up. The AI Inference call handles two jobs: conversing with the caller to build the quote, and extracting a structured JSON quote after the call ends.
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-price-quote-phone-agent-python
cp .env.example .env
pip install -r requirements.txt
Edit .env with your credentials:
TELNYX_API_KEY=KEY019C8F8C3D268774F8F560946B984D9E_...
TELNYX_PUBLIC_KEY=...
AI_MODEL=meta-llama/Llama-3.3-70B-Instruct
QUOTE_NUMBER=+13105551234
Step 2: Understand the Code
Everything lives in app.py (102 lines). Here's what each piece does.
The AI has a built-in product catalog with per-unit pricing:
PRICING = {"voice_minutes": 0.005, "sms_messages": 0.004, "local_numbers": 1.00, "toll_free_numbers": 2.00,
"sip_trunking_channel": 2.50, "ai_inference_1k_tokens": 0.01, "cloud_storage_gb": 0.05, "fax_pages": 0.07}
This gets injected into the system prompt so the AI knows exactly what to charge per unit:
SYSTEM_PROMPT = f"You are a pricing specialist. Available products and per-unit pricing: {json.dumps(PRICING)}. Ask what the caller needs, estimate quantities, build a quote. Keep responses under 2 sentences. After gathering requirements (usually 3-4 questions), summarize the quote with line items and monthly total."
The prompt is the control surface. You can swap the product catalog, change the questioning style, or adjust how many questions the AI asks before summarizing.
The webhook handler is the core state machine. Each Telnyx event triggers the next action:
if event_type == "call.initiated" and p.get("direction") == "incoming":
active_calls[ccid] = {"caller": p.get("from"), "conversation": [{"role": "system", "content": SYSTEM_PROMPT}], "start": time.time()}
client.calls.actions.answer(ccid)
elif event_type == "call.answered":
client.calls.actions.speak(ccid, payload="Hi! I can put together a custom quote for you right now. What kind of communication services are you looking for?", voice="female", language_code="en-US")
elif event_type == "call.speak.ended" and call:
client.calls.actions.gather(ccid, input_type="speech", end_silence_timeout_secs=2, timeout_secs=20, language_code="en-US")
elif event_type == "call.gather.ended" and call:
speech = p.get("speech", {}).get("result", "")
call["conversation"].append({"role": "user", "content": speech})
response = call_inference(call["conversation"])
call["conversation"].append({"role": "assistant", "content": response})
client.calls.actions.speak(ccid, payload=response, voice="female", language_code="en-US")
The conversation loop: speak → gather → infer → speak. Each turn, the AI sees the full conversation history, so it maintains context across all exchanges.
def call_inference(messages, max_tokens=250):
resp = requests.post(INFERENCE_URL, headers={"Authorization": f"Bearer {TELNYX_API_KEY}", "Content-Type": "application/json"},
json={"model": AI_MODEL, "messages": messages, "max_tokens": max_tokens, "temperature": 0.5}, timeout=15)
resp.raise_for_status()
return resp.json()["choices"][0]["message"]["content"]
Straightforward OpenAI-compatible chat completions call against Telnyx AI Inference. Temperature 0.5 for consistency — you want quotes to be predictable, not creative.
This is the interesting part. When the caller hangs up, the app sends the full conversation back to the AI with a different prompt — extract the structured quote:
elif event_type == "call.hangup":
call = active_calls.pop(ccid, None)
if call and len(call.get("conversation", [])) > 3:
extract = [{"role": "system", "content": "Extract the price quote. Return JSON: line_items (list of {product, quantity, unit_price, subtotal}), monthly_total (number), notes (string)."},
{"role": "user", "content": chr(10).join(f"{m['role']}: {m['content']}" for m in call["conversation"] if m["role"] != "system")}]
quote = json.loads(call_inference(extract, max_tokens=400))
quote["caller"] = call["caller"]
quote["timestamp"] = time.strftime("%Y-%m-%dT%H:%M:%SZ")
quotes.append(quote)
The AI returns clean JSON with line items, quantities, unit prices, subtotals, and a monthly total. The app attaches the caller's number and timestamp, then stores it.
@app.route("/quotes", methods=["GET"])
def list_quotes():
return jsonify({"quotes": quotes[-20:]}), 200
Hit GET /quotes to see the last 20 quotes as JSON — ready to pipe into a CRM, billing system, or analytics dashboard.
python app.py
In a separate terminal, expose your local server with ngrok:
ngrok http 5000
Copy the HTTPS URL and configure it in the Telnyx Portal :
Go to Call Control Applications
Create or edit your application
Set the Webhook URL to https://<your-ngrok-url>.ngrok.app/webhooks/voice
Assign your Telnyx phone number to this Call Control Application if you haven't already.
Call your Telnyx number from any phone. You'll hear:
"Hi! I can put together a custom quote for you right now. What kind of communication services are you looking for?"
Say something like: "I need about 50 phone numbers, 10,000 voice minutes a month, and a few thousand SMS messages."
The AI asks follow-up questions — how many SMS, any toll-free numbers, do you need SIP trunking — then summarizes the quote with line items and a monthly total.
After you hang up, check the extracted quote:
curl http://localhost:5000/quotes | python3 -m json.tool
You'll see a structured JSON object like:
{
"line_items": [
{"product": "local_numbers", "quantity": 50, "unit_price": 1.00, "subtotal": 50.00},
{"product": "voice_minutes", "quantity": 10000, "unit_price": 0.005, "subtotal": 50.00},
{"product": "sms_messages", "quantity": 5000, "unit_price": 0.004, "subtotal": 20.00}
],
"monthly_total": 120.00,
"notes": "Customer indicated potential for higher SMS volume in Q3.",
"caller": "+12125551234",
"timestamp": "2026-07-15T14:30:00Z"
}
Customizing the Agent
The pricing catalog and system prompt are the control surface. Small changes produce very different agents:
PRICING = {"consulting_hour": 200.00, "audit_report": 2500.00, "retainer_monthly": 8000.00, "training_session": 1500.00}
Change the questioning style:
SYSTEM_PROMPT = f"You are a pricing specialist. Available products: {json.dumps(PRICING)}. Ask one question at a time. Be conversational but efficient. After exactly 3 questions, summarize the quote with line items and monthly total. If the caller asks for a discount, explain standard pricing tiers."
Add discount logic:
SYSTEM_PROMPT = f"You are a pricing specialist. Products: {json.dumps(PRICING)}. For volumes over 10,000 units, apply a 15% discount. For over 50,000 units, apply 25%. Mention the discount in your summary."
Send the quote via SMS after the call:
# After extracting the quote, send it as SMS
from telnyx import Message
quote_text = f"Your quote: {quote['monthly_total']}/month. {len(quote['line_items'])} line items."
Message.create(to=call["caller"], from_=QUOTE_NUMBER, text=quote_text)
The pattern — conversational AI with a product catalog + post-call structured extraction — works for any quoting or ordering flow.
This example uses in-memory storage for simplicity. For a production deployment:
Database — replace the in-memory quotes list with PostgreSQL or Redis so quotes survive restarts
Concurrency — run the app behind gunicorn with multiple workers
Error recovery — handle inference timeouts and call failures gracefully with retry or SMS fallback
Authentication — add API key validation on the /quotes endpoint
Webhook verification — the app already validates Telnyx Ed25519 signatures; ensure your public key is current
Prompt tuning — test different system prompts and temperature settings for your product catalog
Rate limiting — protect your webhook endpoint from abuse
Monitoring — add structured logging and alerting on call success/failure rates
CRM integration — pipe extracted quotes into Salesforce, HubSpot, or your billing system
Can I use a different AI model? Yes. Telnyx AI Inference is OpenAI-compatible. Set AI_MODEL in .env to any model available on the platform. Llama 3.3 70B Instruct is a good default for structured quoting.
How much does a call cost? Telnyx

[truncated]
