---
source: "https://lowlatencyclub.ai/blog/posts/ai-negotiation-practice-phone-python"
hn_url: "https://news.ycombinator.com/item?id=48921718"
title: "Practice Negotiations with an AI Phone Agent Real-Time Roleplay with Telnyx"
article_title: "Practice Negotiations with an AI Phone Agent — Real-Time Roleplay with Telnyx Voice AI | Low Latency Club"
author: "harpreetseehra"
captured_at: "2026-07-15T15:19:14Z"
capture_tool: "hn-digest"
hn_id: 48921718
score: 2
comments: 0
posted_at: "2026-07-15T14:50:14Z"
tags:
  - hacker-news
  - translated
---

# Practice Negotiations with an AI Phone Agent Real-Time Roleplay with Telnyx

- HN: [48921718](https://news.ycombinator.com/item?id=48921718)
- Source: [lowlatencyclub.ai](https://lowlatencyclub.ai/blog/posts/ai-negotiation-practice-phone-python)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T14:50:14Z

## Translation

タイトル: AI 電話エージェントとの交渉の練習 Telnyx を使用したリアルタイム ロールプレイ
記事のタイトル: AI 電話エージェントとの交渉の練習 — Telnyx Voice AI を使用したリアルタイム ロールプレイ |低遅延クラブ
説明: 携帯電話を手に取り、番号にダイヤルし、採用担当マネージャーを演じる AI を相手に給与交渉の練習をするところを想像してください。AI はあなたの背中を押してくれます。

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
AI 電話エージェントとの交渉の練習 — Telnyx Voice AI を使用したリアルタイム ロールプレイ
携帯電話を手に取り、番号にダイヤルし、採用担当マネージャーを演じる AI を相手に給与交渉の練習をするところを想像してみてください。AI は最初のオファーを押し返し、予算の制約で反論し、電話を切った後にあなたのテクニックを採点します。模擬面接のスケジュールを立てたり、コーチにお金を払ったり、同僚と気まずいロールプレイをしたりする必要はありません。
これは AI Negotiation Practice Phone です。Telnyx 通話制御と AI 推論で構築された 110 行の Python アプリです。 3 つのシナリオ (給与、販売契約、ベンダー契約)、音声による会話、および通話のたびに提供される構造化されたパフォーマンス スコア。 AI は性格を維持し、ユーザーのアプローチに適応し、実用的なフィードバックを提供します。
このチュートリアルでは、これを最初から構築します。リポジトリのクローンを作成し、電話番号を設定すれば、数分で練習を開始できます。
交渉を練習するために誰でも電話できる電話番号:
発信者がダイヤルイン - Telnyx が応答し、

メニューを提供しています
シナリオの選択 — 給与交渉の場合は 1、販売契約の場合は 2、ベンダー契約の場合は 3 を押してください
AI がオープン – AI が反対の役割 (採用マネージャー、エンタープライズ バイヤー、ベンダー アカウント マネージャー) を果たし、ポジションを獲得します
ライブ交渉 – 発信者は自然に話し、AI はキャラクターに応じて反応し、押し返し、反論し、適応します
通話後のスコアリング - 電話を切ると、AI が 5 つの次元にわたってネゴシエーションをスコアリングし、構造化された JSON を返します。
セッション履歴 — すべての練習セッションは保存され、GET /sessions 経由でアクセスできます。
インタラクション全体は音声主導で行われます。Text-to-Speech が AI のセリフを読み上げ、発信者は自然な音声で応答します。このモデル (Telnyx AI 推論経由の Llama 3.3 70B) は、リアルタイムのロールプレイと通話後の評価の両方を処理します。
ほとんどの交渉トレーニング ツールは、テキストベースのチャットボットまたは静的なビデオ コースです。どちらも、生の会話のプレッシャー、つまり一時停止、反発、立ち上がって考えなければならない瞬間を捉えていません。これは、予算、役割、および隠れた制約 (「最大 15% 割引」または「予算は 15 万 5,000 ドルで、16 万 5,000 ドルまでの柔軟性」など) を備えた AI と実際の電話通話を行うことになります。
採点システムにより、単なる会話以上のものになります。電話を切った後、AI はあなたのパフォーマンスを 5 つの側面 (アンカリング、譲歩戦略、積極的な傾聴、創造性、自信) に加えて全体的なスコアと具体的な長所/改善点で評価します。同じシナリオを 5 回練習して、テクニックが向上したかどうかを追跡できます。
また、DTMF + 音声パターンも示します。アプリは、シナリオの選択とネゴシエーション自体の音声認識に DTMF (キーパッド入力) を使用します。この 2 つの入力パターンは、多くの実際の IVR ワークフローに現れます。
残高が入金された Telnyx アカウント
Telnyx 電話番号

音声が有効になっているバー
Webhook URL で構成された通話制御アプリケーション
ローカルサーバーを Telnyx Webhook に公開するための ngrok
電話
│
▼
Telnyx 通話制御 (Webhook イベント)
│
▼
Flask アプリ (app.py、110 行)
│
§──► 通話開始 → 電話に出る
§──► 電話応答 → TTS メニュー (1、2、または 3 を押します)
§──► call.speak.ended → DTMF収集（シナリオ選択）
§──► call.gather.ended → DTMF → シナリオ選択 → AI 開始ライン
│ └─► 音声→AI推論→TTS応答（交渉ループ）
└──► call.hangup → スコアネゴシエーションを JSON として → 保存
│
▼
Telnyx AI 推論 (Llama 3.3 70B)
│
▼
構造化スコア (5 次元の JSON + フィードバック)
このアプリは、選択 (DTMF メニュー) とネゴシエーション (音声会話) の 2 つのフェーズを持つステート マシンです。各 Telnyx Webhook イベントは次のアクションを引き起こします。電話を切った後、構造化された JSON を返すスコアリング プロンプトとともに会話全体が AI に送信されます。
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-negotiation-practice-phone-python
cp .env.example .env
pip install -r 要件.txt
資格情報を使用して .env を編集します。
TELNYX_API_KEY=your_telnyx_api_key_here
TELNYX_PUBLIC_KEY=...
AI_MODEL=メタ-ラマ/ラマ-3.3-70B-命令
PRACTICE_NUMBER=+13105551234
ステップ 2: コードを理解する
すべては app.py (110 行) にあります。各部分の動作は次のとおりです。
アプリには 3 つの組み込みネゴシエーション シナリオが付属しており、それぞれにロールと非表示のコンテキストが含まれています。
シナリオ = {
"1": {"role": "採用マネージャー", "context": "候補者は 180,000 ドルを希望しています。予算は 155,000 ドルですが、165,000 ドルまで柔軟に対応できます。経験レベルを押し戻します。代替案として株式または契約ボーナスを提供できます。"},
"2": {"役割": "エンタープライズバイヤー", "コンテ

xt": "あなたは、同社の SaaS 製品を年間 5 万ドルで評価しています。 35,000 ドルの競合オファーがあります。予算は 45,000 ドルです。ボリュームディスカウントとより長い支払い期間を求めてください。"},
"3": {"role": "ベンダー アカウント マネージャー", "context": "クライアントは契約を 40% 削減したいと考えています。トップ 10 に入るアカウントです。最大 15% の割引を提供することも、別の条件で取引を再構築することもできます。」}
}
コンテキストはシステム プロンプトに挿入されるため、AI はその制約を認識しますが、呼び出し元はそれを知りません。実際の交渉と同じように、会話を通じて相手の予算や柔軟性を発見します。
ステートマシン: 選択 → 交渉
Webhook ハンドラーには 2 つの状態があります。選択状態では、DTMF 数字を収集してシナリオを選択します。
if call["state"] == "選択":
client.calls.actions.gather(ccid, input_type="dtmf", timeout_secs=10, min_digits=1, max_digits=1)
発信者が 1、2、または 3 を押すと、アプリはシナリオをロードし、システム プロンプトを構築し、AI に冒頭のセリフを要求し、交渉状態に切り替えます。
if call["state"] == "選択":
シナリオ = SCENARIOS.get(数字, SCENARIOS["1"])
call["状態"] = "交渉中"
call["シナリオ"] = シナリオ
call["conversation"] = [{"role": "system", "content": f"あなたは交渉における {scenario['role']} です。 {scenario['context']} キャラクターのままでいてください。毅然とした態度で、しかし公平に。最初のオファーを押し返します。回答は 2 文以内にしてください。 6 回のやり取りの後、まとめを開始します。"}]
starting = call_inference(call["conversation"] + [{"role": "user", "content": "交渉が始まります。開始位置を決めてください。"}])
call["会話"].append({"役割": "アシスタント", "コンテンツ": オープニング})
client.calls.actions.speak(ccid, payload=opening, voice="女性", language_code="en-US")
交渉状態では、発言を収集し、会話を実行します。

p:
それ以外の場合:
client.calls.actions.gather(ccid, input_type="speech", end_silence_timeout_secs=2, timeout_secs=20, language_code="en-US")
会話のループ: 話す→集まる→推測する→話す。 AI はターンごとに完全な会話履歴を確認するため、コンテキストを維持し、ユーザーが提供した内容を記憶します。
def call_inference(messages, max_tokens=200):
resp = request.post(INFERENCE_URL, headers={"Authorization": f"Bearer {TELNYX_API_KEY}", "Content-Type": "application/json"},
json={"モデル": AI_MODEL、"メッセージ": メッセージ、"max_tokens": max_tokens、"温度": 0.7}、タイムアウト=15)
resp.raise_for_status()
return resp.json()["選択"][0]["メッセージ"]["コンテンツ"]
Telnyx AI 推論に対する OpenAI 互換のチャット完了呼び出し。温度 0.7 — 交渉者に毎回完全に一貫性を持たせるのではなく、もう少し創造的で予測不能な行動をとらせるため、価格見積エージェント (0.5) よりも高くなります。
発信者が電話を切ると、アプリはスコアリング プロンプトとともに会話全体を AI に送信します。
elif イベントタイプ == "call.hangup":
call = active_calls.pop(ccid, なし)
if call と len(call.get("conversation", [])) > 3:
core_prompt = [{"role": "system", "content": "この交渉慣行にスコアを付けます。返される JSON: アンカーリング (1 ～ 10)、concession_strategy (1 ～ 10)、active_listening (1 ～ 10)、創造性 (1 ～ 10)、自信 (1 ～ 10)、全体 (1 ～ 10)、強み (リスト)、改善点 (リスト)、deal_outcome (文字列)。"},
{"role": "user", "content": chr(10).join(f"{m['role']}: {m['content']}" for m in call["conversation"] if m["role"] != "system")}]
試してみてください:
スコア = json.loads(call_inference(score_prompt, max_tokens=400))
session.append({"scenario": call.get("scenario", {}).get("role"), "score": スコア, "duration": int(time.time() - call["start"])})
例外を除く:
パスする
AI は 5 つの構造化されたスコアを返します。

次元、強みのリスト、改善点のリスト、および取引の結果。アプリは、それをシナリオと通話時間と一緒に保存します。
@app.route("/セッション", メソッド=["GET"])
def list_sessions():
return jsonify({"セッション": セッション[-20:]}), 200
GET /sessions を押して、過去 20 回の練習セッションを JSON として表示します。時間の経過とともに進捗状況を追跡したり、シナリオを比較したり、分析ダッシュボードにパイプしたりできます。
アプリには、古い通話状態をクリーンアップするバックグラウンド スレッドが含まれています。
def _start_ttl_cleanup(*ストア、ttl_秒 = 3600、間隔 = 300):
def _cleanup():
True の場合:
_ttl_time.sleep(間隔)
カットオフ = _ttl_time.time() - ttl_秒
店舗での保管の場合:
期限切れ = [k の k、store.items() の v
if isinstance(v, dict) および v.get("_ts", _ttl_time.time()) < カットオフ]
期限切れの k の場合:
store.pop(k, なし)
threading.Thread(target=_cleanup, daemon=True).start()
クラッシュまたは放棄された通話は 1 時間後に自動的にクリーンアップされます。長時間実行される展開でもメモリ リークは発生しません。
Python app.py
別のターミナルで、ngrok を使用してローカル サーバーを公開します。
ngrok http 5000
HTTPS URL をコピーし、Telnyx ポータルで設定します。
通話制御アプリケーションに移動します
アプリケーションを作成または編集する
Webhook URL を https://<your-ngrok-url>.ngrok.app/webhooks/voice に設定します。
Telnyx 電話番号をまだ割り当てていない場合は、この通話制御アプリケーションに割り当てます。
どの電話からでも Telnyx 番号に電話をかけます。次のように聞こえます:
「交渉練習！給与交渉は1、売買契約は2、ベンダー契約は3を押してください。」
1 を押します。AI (採用マネージャー役) が次のような内容で開きます。
「お電話いただきありがとうございます。あなたの経験を検討しましたが、正直に言います。通常、このレベルの候補者は 15 万 5,000 ドルで採用されます。18 万ドルと言ったことは知っていますが、経験の差を考えると、正当な理由が言えるかわかりません。

その範囲です。あなたにプレミアムを付ける価値があるのは何でしょうか?」
交渉してください。押し戻します。自分の主張を主張してください。 AI はこれに対抗し、制約を提示し、最終的には合意に達するか、堅持します。
電話を切った後、スコアを確認してください。
カール http://localhost:5000/sessions | python3 -m json.tool
次のような構造化されたスコア オブジェクトが表示されます。
{
"シナリオ": "採用マネージャー",
「スコア」: {
「アンカリング」: 7、
「譲歩戦略」: 5、
「アクティブリスニング」: 8、
「創造性」: 6、
「自信」: 7、
「全体」: 6.6、
"強み": ["18万ドルの強力なオープニングアンカー"、"基本給の代替としての株式についての質問"],
"改善": ["基本給に関してあまりにも早く譲歩した", "契約ボーナスのオプションを検討しなかった"],
"deal_outcome": "ベース 162,000 ドル + 契約ボーナス 10,000 ドルで決済されました"
}、
「期間」: 184
}
シナリオのカスタマイズ
SCENARIOS dict はコントロール サーフェスです。小さな変更により、まったく異なる練習セッションが生成されます。
不動産交渉を追加します。
"4": {"role": "売主の代理人", "context": "その家は 65 万ドルで上場されています。売り手は最低 $620,000 を受け入れます。購入者は興味はあるものの、価格には敏感なようです。要求に近いものを求めますが、締め切り日については柔軟な姿勢を示します。」}
AI をより攻撃的にします。
call["conversation"] = [{"role": "system", "content": f"あなたは交渉における {scenario['role']} です。 {scenario['context']} キャラクターのままでいてください。アグレッシブに、そして一生懸命にプッシュしてください

[切り捨てられた]

## Original Extract

Imagine picking up your phone, dialing a number, and practicing a salary negotiation against an AI that plays the hiring manager — one that pushes back on yo...

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
Practice Negotiations with an AI Phone Agent — Real-Time Roleplay with Telnyx Voice AI
Imagine picking up your phone, dialing a number, and practicing a salary negotiation against an AI that plays the hiring manager — one that pushes back on your first offer, counters with budget constraints, and then scores your technique after you hang up. No scheduling a mock interview, no paying a coach, no awkward roleplay with a colleague.
This is the AI Negotiation Practice Phone — a 110-line Python app built with Telnyx Call Control and AI Inference. Three scenarios (salary, sales deal, vendor contract), voice-driven conversation, and a structured performance score delivered after every call. The AI stays in character, adapts to your approach, and gives you actionable feedback.
In this walkthrough, you'll build it from scratch. Clone the repo, configure a phone number, and start practicing in minutes.
A phone number anyone can call to practice negotiations:
Caller dials in — Telnyx answers and offers a menu
Scenario selection — press 1 for salary negotiation, 2 for sales deal, 3 for vendor contract
AI opens — the AI plays the opposing role (hiring manager, enterprise buyer, or vendor account manager) and makes an opening position
Live negotiation — caller speaks naturally, AI responds in character, pushes back, counters, and adapts
Post-call scoring — on hangup, the AI scores the negotiation across 5 dimensions and returns structured JSON
Session history — every practice session is stored and accessible via GET /sessions
The whole interaction is voice-driven: Text-to-Speech reads the AI's lines, and the caller responds with natural speech. The model (Llama 3.3 70B via Telnyx AI Inference) handles both the real-time roleplay and the post-call evaluation.
Most negotiation training tools are text-based chatbots or static video courses. Neither captures the pressure of a live conversation — the pauses, the pushback, the moment you have to think on your feet. This one puts you on a real phone call with an AI that has a budget, a role, and a hidden constraint (like "max 15% discount" or "budget is $155K with flexibility to $165K").
The scoring system makes it more than a conversation. After you hang up, the AI evaluates your performance across five dimensions — anchoring, concession strategy, active listening, creativity, and confidence — plus an overall score and specific strengths/improvements. You can practice the same scenario five times and track whether your technique improves.
It also demonstrates the DTMF + speech pattern: the app uses DTMF (keypad input) for scenario selection and speech recognition for the negotiation itself. This two-input pattern shows up in many real IVR workflows.
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
Flask app (app.py, 110 lines)
│
├──► call.initiated → answer the call
├──► call.answered → TTS menu (press 1, 2, or 3)
├──► call.speak.ended → gather DTMF (scenario selection)
├──► call.gather.ended → DTMF → pick scenario → AI opening line
│ └──► speech → AI inference → TTS response (negotiation loop)
└──► call.hangup → score negotiation as JSON → store
│
▼
Telnyx AI Inference (Llama 3.3 70B)
│
▼
Structured score (JSON with 5 dimensions + feedback)
The app is a state machine with two phases: selection (DTMF menu) and negotiating (speech conversation). Each Telnyx webhook event drives the next action. After hangup, the full conversation is sent to the AI with a scoring prompt that returns structured JSON.
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-negotiation-practice-phone-python
cp .env.example .env
pip install -r requirements.txt
Edit .env with your credentials:
TELNYX_API_KEY=your_telnyx_api_key_here
TELNYX_PUBLIC_KEY=...
AI_MODEL=meta-llama/Llama-3.3-70B-Instruct
PRACTICE_NUMBER=+13105551234
Step 2: Understand the Code
Everything lives in app.py (110 lines). Here's what each piece does.
The app ships with three built-in negotiation scenarios, each with a role and a hidden context:
SCENARIOS = {
"1": {"role": "hiring manager", "context": "The candidate wants $180K. Your budget is $155K with flexibility to $165K. Push back on experience level. You can offer equity or signing bonus as alternatives."},
"2": {"role": "enterprise buyer", "context": "You're evaluating their SaaS product at $50K/year. You have a competing offer at $35K. Your budget is $45K. Ask for volume discounts and longer payment terms."},
"3": {"role": "vendor account manager", "context": "The client wants to reduce their contract by 40%. They're a top-10 account. You can offer 15% discount max, or restructure the deal with different terms."}
}
The context is injected into the system prompt so the AI knows its constraints — but the caller doesn't know them. You discover the other side's budget and flexibility through the conversation, just like a real negotiation.
The State Machine: Selection → Negotiating
The webhook handler has two states. In select state, it gathers DTMF digits to pick a scenario:
if call["state"] == "select":
client.calls.actions.gather(ccid, input_type="dtmf", timeout_secs=10, min_digits=1, max_digits=1)
When the caller presses 1, 2, or 3, the app loads the scenario, builds the system prompt, asks the AI for an opening line, and switches to negotiating state:
if call["state"] == "select":
scenario = SCENARIOS.get(digits, SCENARIOS["1"])
call["state"] = "negotiating"
call["scenario"] = scenario
call["conversation"] = [{"role": "system", "content": f"You are a {scenario['role']} in a negotiation. {scenario['context']} Stay in character. Be firm but fair. Push back on their first offer. Keep responses under 2 sentences. After 6 exchanges, start wrapping up."}]
opening = call_inference(call["conversation"] + [{"role": "user", "content": "The negotiation begins. Make your opening position."}])
call["conversation"].append({"role": "assistant", "content": opening})
client.calls.actions.speak(ccid, payload=opening, voice="female", language_code="en-US")
In negotiating state, it gathers speech and runs the conversation loop:
else:
client.calls.actions.gather(ccid, input_type="speech", end_silence_timeout_secs=2, timeout_secs=20, language_code="en-US")
The conversation loop: speak → gather → infer → speak. Each turn, the AI sees the full conversation history, so it maintains context and remembers what you offered.
def call_inference(messages, max_tokens=200):
resp = requests.post(INFERENCE_URL, headers={"Authorization": f"Bearer {TELNYX_API_KEY}", "Content-Type": "application/json"},
json={"model": AI_MODEL, "messages": messages, "max_tokens": max_tokens, "temperature": 0.7}, timeout=15)
resp.raise_for_status()
return resp.json()["choices"][0]["message"]["content"]
OpenAI-compatible chat completions call against Telnyx AI Inference. Temperature 0.7 — higher than the price quote agent (0.5) because you want the negotiator to be a bit more creative and unpredictable, not perfectly consistent every time.
When the caller hangs up, the app sends the full conversation to the AI with a scoring prompt:
elif event_type == "call.hangup":
call = active_calls.pop(ccid, None)
if call and len(call.get("conversation", [])) > 3:
score_prompt = [{"role": "system", "content": "Score this negotiation practice. Return JSON: anchoring (1-10), concession_strategy (1-10), active_listening (1-10), creativity (1-10), confidence (1-10), overall (1-10), strengths (list), improvements (list), deal_outcome (string)."},
{"role": "user", "content": chr(10).join(f"{m['role']}: {m['content']}" for m in call["conversation"] if m["role"] != "system")}]
try:
score = json.loads(call_inference(score_prompt, max_tokens=400))
sessions.append({"scenario": call.get("scenario", {}).get("role"), "score": score, "duration": int(time.time() - call["start"])})
except Exception:
pass
The AI returns a structured score with five dimensions, a list of strengths, a list of improvements, and the deal outcome. The app stores it alongside the scenario and call duration.
@app.route("/sessions", methods=["GET"])
def list_sessions():
return jsonify({"sessions": sessions[-20:]}), 200
Hit GET /sessions to see the last 20 practice sessions as JSON — track your progress over time, compare scenarios, or pipe into an analytics dashboard.
The app includes a background thread that cleans up stale call state:
def _start_ttl_cleanup(*stores, ttl_seconds=3600, interval=300):
def _cleanup():
while True:
_ttl_time.sleep(interval)
cutoff = _ttl_time.time() - ttl_seconds
for store in stores:
expired = [k for k, v in store.items()
if isinstance(v, dict) and v.get("_ts", _ttl_time.time()) < cutoff]
for k in expired:
store.pop(k, None)
threading.Thread(target=_cleanup, daemon=True).start()
Calls that crash or get abandoned are automatically cleaned up after an hour — no memory leaks in long-running deployments.
python app.py
In a separate terminal, expose your local server with ngrok:
ngrok http 5000
Copy the HTTPS URL and configure it in the Telnyx Portal :
Go to Call Control Applications
Create or edit your application
Set the Webhook URL to https://<your-ngrok-url>.ngrok.app/webhooks/voice
Assign your Telnyx phone number to this Call Control Application if you haven't already.
Call your Telnyx number from any phone. You'll hear:
"Negotiation Practice! Press 1 for salary negotiation, 2 for sales deal, 3 for vendor contract."
Press 1. The AI (playing a hiring manager) opens with something like:
"Thanks for calling in. I've reviewed your experience, and I have to be honest — we typically bring in candidates at $155K for this level. I know you mentioned $180K, but given the experience gap, I'm not sure I can justify that range. What would make you worth the premium?"
Negotiate. Push back. Make your case. The AI will counter, bring up constraints, and eventually either reach a deal or hold firm.
After you hang up, check your score:
curl http://localhost:5000/sessions | python3 -m json.tool
You'll see a structured score object like:
{
"scenario": "hiring manager",
"score": {
"anchoring": 7,
"concession_strategy": 5,
"active_listening": 8,
"creativity": 6,
"confidence": 7,
"overall": 6.6,
"strengths": ["Strong opening anchor at $180K", "Asked about equity as alternative to base salary"],
"improvements": ["Conceded too quickly on base salary", "Didn't explore signing bonus option"],
"deal_outcome": "Settled at $162K base + $10K signing bonus"
},
"duration": 184
}
Customizing the Scenarios
The SCENARIOS dict is the control surface. Small changes produce very different practice sessions:
Add a real estate negotiation:
"4": {"role": "seller's agent", "context": "The house is listed at $650K. The seller will accept $620K minimum. The buyer seems interested but price-sensitive. Push for close to asking but signal flexibility on closing dates."}
Make the AI more aggressive:
call["conversation"] = [{"role": "system", "content": f"You are a {scenario['role']} in a negotiation. {scenario['context']} Stay in character. Be aggressive and push hard

[truncated]
