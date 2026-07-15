---
source: "https://lowlatencyclub.ai/blog/posts/ai-phone-story-hotline-python"
hn_url: "https://news.ycombinator.com/item?id=48918668"
title: "Build an AI Phone Story Hotline – Interactive Storytelling with Telnyx Voice AI"
article_title: "Build an AI Phone Story Hotline — Interactive Storytelling with Telnyx Voice AI | Low Latency Club"
author: "harpreetseehra"
captured_at: "2026-07-15T11:22:34Z"
capture_tool: "hn-digest"
hn_id: 48918668
score: 1
comments: 0
posted_at: "2026-07-15T10:25:42Z"
tags:
  - hacker-news
  - translated
---

# Build an AI Phone Story Hotline – Interactive Storytelling with Telnyx Voice AI

- HN: [48918668](https://news.ycombinator.com/item?id=48918668)
- Source: [lowlatencyclub.ai](https://lowlatencyclub.ai/blog/posts/ai-phone-story-hotline-python)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T10:25:42Z

## Translation

タイトル: AI 電話ストーリー ホットラインの構築 – Telnyx Voice AI を使用したインタラクティブなストーリーテリング
記事のタイトル: AI 電話ストーリー ホットラインを構築する — Telnyx Voice AI を使用したインタラクティブなストーリーテリング |低遅延クラブ
説明: 電話番号に電話し、ジャンルを選択して、AI が生成したインタラクティブなストーリーを聞き、あなたの選択が物語を形作ります。 Python は 104 行。

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
AI 電話ストーリー ホットラインを構築する — Telnyx Voice AI を使用したインタラクティブなストーリーテリング
電話番号に電話し、ミステリー、SF、ファンタジー、ホラー、ロマンスなどのジャンルを選択し、ユーザーの選択にリアルタイムで適応する AI のナレーションを聞くことを想像してみてください。 1を押すと軋むドアが開きます。 2 を押してウィンドウを確認します。ストーリーは分岐し、AI は続行し、すべての通話が新しい冒険になります。
これは AI Phone Story Hotline です。Telnyx 通話制御と AI 推論で構築された 104 行の Python アプリです。ゲーム エンジンも分岐スクリプトも、事前に作成されたダイアログ ツリーもありません。 AI はユーザーが進むにつれてストーリーを生成し、携帯電話のキーパッドが次に何が起こるかを形作ります。
このチュートリアルでは、これを最初から構築します。リポジトリのクローンを作成し、電話番号を構成し、数分でデプロイします。
インタラクティブなストーリーを開始するために誰でも電話できる電話番号:
発信者がダイヤルイン — Telnyx が応答し、ジャンル メニューで挨拶します
発信者はジャンルを選択します — ミステリー、SF、ファンタジー、ホラー、またはロマンスの場合は 1 ～ 5 を押します
AIg

第 1 章を生成します — 2 つの選択肢で終わる 3 ～ 4 文のストーリーセグメント
発信者が選択します — 1 または 2 を押すか、選択内容を声に出して言います
ストーリーは続きます – AI が選択に基づいて次の章を生成し、完全な会話のコンテキストを維持します
5 章後 — AI が物語を満足のいく結末に導きます
インタラクション全体は音声主導で行われます。Text-to-Speech が各章を読み上げ、発信者は DTMF キーを押すか音声で応答します。 AI モデル (Telnyx AI 推論経由の Llama 3.3 70B) がストーリーテリングを処理し、通話制御が電話を処理します。
AI 電話のデモのほとんどは、テーブルの予約、潜在顧客の評価、パスワードのリセットなど、ビジネス ユース ケースです。これは違います。それは創造的な AI です。モデルはリアルタイムでフィクションを書き、人間の入力に基づいて分岐し、それを電話で配信します。ストーリーテリング形式 (短い章、それぞれ 2 つの選択肢) により、AI の応答が緊密になり、通話が魅力的になります。
また、あらゆる会話型 AI アプリで機能するパターン (Webhook 駆動のステート マシン + 会話メモリを備えた LLM + 音声 I/O) も示しています。ストーリー ホットラインがどのように機能するかを確認したら、ストーリーテリング プロンプトを任意のドメインに置き換えることができます。つまり、自分で選択するアドベンチャーのオンボーディング フロー、インタラクティブなクイズ、分岐トレーニング シミュレーションなどです。
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
Flask アプリ (app.py、104 行)
│
§──► 通話開始 → 電話に出る
§──► 電話応答 → TTS グリーティング + ジャンルメニュー
§──► call.gather.ended → DTMF/音声入力 → AI推論
§──► call.speak.ending → 次の選択肢を集める → AI Inferen

CE
└──► call.hangup → クリーンアップセッション
│
▼
Telnyx AI 推論 (Llama 3.3 70B)
│
▼
ストーリーの章 (発信者に TTS を返す)
このアプリは、Telnyx Webhook イベントによって駆動されるステート マシンです。各イベントは、応答する、話す、集まる、電話を切るなどの次のアクションをトリガーします。 AI 推論呼び出しはその中間に位置し、実行中の会話履歴からストーリーの章を生成します。
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-phone-story-hotline-python
cp .env.example .env
pip install -r 要件.txt
資格情報を使用して .env を編集します。
TELNYX_API_KEY=KEY0123456789ABCDEF # portal.telnyx.com/api-keys から
TELNYX_PUBLIC_KEY= # portal.telnyx.com/api-keys から (公開キー)
STORY_NUMBER=+13105551234 # Telnyx の電話番号
AI_MODEL=メタ-ラマ/ラマ-3.3-70B-命令
ステップ 2: コードを理解する
アプリ全体は、1 つのファイル app.py に 104 行あります。主要な部分は次のとおりです。
Webhook署名の検証
すべての Telnyx Webhook は Ed25519 キーで署名されます。アプリはイベントを信頼する前に署名を検証します。
@app.route("/webhooks/voice",methods=["POST"])
def handle_voice():
試してみてください:
client.webhooks.unwrap(request.get_data(as_text=True), headers=dict(request.headers))
例外を除く:
return jsonify({"エラー": "無効な署名"})、401
これにより、なりすましの Webhook 呼び出しが呼び出しアクションをトリガーするのを防ぎます。
各呼び出しは、 call_control_id をキーとするメモリ内の辞書で追跡されます。
アクティブコール = {}
ステート マシンは 5 つのイベントを処理します。
通話が開始されました — 通話状態を保存して応答します。
イベントタイプ == "call.initiated" かつ p.get("direction") == "incoming" の場合:
active_calls[ccid] = {"状態": "ジャンル選択", "会話": [], "チャプター": 0}
client.calls.actions.answer(ccid)
電話に応答しました — テキスト読み上げを使用して、ジャンル メニューで発信者に応答します。

elifevent_type == "call.answered":
client.calls.actions.speak(ccid,
payload="ストーリー ホットラインへようこそ! 冒険を選択してください。ミステリーの場合は 1、SF の場合は 2、ファンタジーの場合は 3、ホラーの場合は 4、ロマンスの場合は 5 を押してください。",
voice="女性"、言語コード="en-US")
通話が終了しました — TTS が終了したら、発信者の入力を収集します。ジャンルの選択中に、DTMF 桁を収集します。ストーリー中に音声または DTMF を収集します。
elifevent_type == "call.speak.ended" で次のように呼び出します。
if call["state"] == "genre_select":
client.calls.actions.gather(ccid, input_type="dtmf", timeout_secs=10, min_digits=1, max_digits=1)
それ以外の場合:
client.calls.actions.gather(ccid, input_type="speech dtmf", end_silence_timeout_secs=3, timeout_secs=20, language_code="en-US")
ストーリーテリングのプロンプト
発信者がジャンルを選択すると、アプリは通話の残りの部分で AI をガイドするシステム プロンプトを構築します。
ジャンル = {"1": "ミステリー"、"2": "SF"、"3": "ファンタジー"、"4": "ホラー"、"5": "ロマンス"}
if call["state"] == "genre_select":
ジャンル = GENRES.get(数字, "ミステリー")
call["状態"] = "ストーリー"
call["会話"] = [{"役割": "システム", "コンテンツ":
f「あなたは、電話ホットラインの対話型 {ジャンル} ストーリーテラーです。」
f「短い章 (それぞれ 3 ～ 4 文) で魅力的なストーリーを語ってください。」
f「各章の最後には、正確に 2 つの選択肢があります:「1 を押して...」または「2 を押して...」。
f「鮮やかで映画的なものにしてください。5 章を経て、物語を満足のいく結末に導きます。」
}]
story_start = call_inference(call["会話"] + [{"role": "user", "content": "ストーリーを開始します。"}])
call["会話"].append({"役割": "アシスタント", "コンテンツ": story_start})
client.calls.actions.speak(ccid, payload=story_start, voice="女性", language_code="en-US")
プロンプトは次の 3 つのことを行います。
形式の制約 - 電話での会話に自然に収まる短い章 (3 ～ 4 文)
選択構造

e — 収集ステップで DTMF をキャプチャできるように、「Press 1」または「Press 2」で終わる 2 つのオプション
終了条件 — 5 章後にストーリーを終了する
call_inference ヘルパーは、完全な会話履歴を Telnyx AI Inference に送信し、モデルの応答を返します。
def call_inference(messages, max_tokens=250):
resp = request.post(INFERENCE_URL,
headers={"Authorization": f"Bearer {TELNYX_API_KEY}", "Content-Type": "application/json"},
json={"モデル": AI_MODEL、"メッセージ": メッセージ、"max_tokens": max_tokens、"温度": 0.9, },
タイムアウト=20)
resp.raise_for_status()
return resp.json()["選択"][0]["メッセージ"]["コンテンツ"]
温度は 0.9 に設定されています。創造的なストーリーテリングには十分に高く、物語の一貫性を維持するには十分に低くなります。モデルは Llama 3.3 70B Instruct で、OpenAI 互換 API を備えた Telnyx AI Inference で利用できます。
発信者が選択を行うたびに、アプリはそれを会話に追加し、AI に次の章について尋ねます。
elif call["状態"] == "ストーリー":
選択 = 数字または音声
call["conversation"].append({"role": "user", "content": f"オプション {choice} を選択します"})
call["章"] += 1
call["chapters"] >= 5の場合:
call["conversation"][-1]["content"] += "。物語を劇的な結末に導きます。"
継続 = call_inference(call["会話"])
call["会話"].append({"役割": "アシスタント", "コンテンツ": 続き})
client.calls.actions.speak(ccid, payload=continuation, voice="女性", language_code="en-US")
5 章が終わると、アプリはストーリーを終わらせるための最終命令を挿入します。 AI はすべての通話で完全な会話履歴を確認するため、すべての章にわたって物語の連続性が維持されます。
Python app.py
別のターミナルで、ngrok を使用してローカル サーバーを公開します。
ngrok http 5000
HTTPS URL をコピーし、Telnx Por で設定します。

タル:
通話制御アプリケーションに移動します
アプリケーションを作成または編集する
Webhook URL を https://<your-ngrok-url>.ngrok.app/webhooks/voice に設定します。
Telnyx 電話番号をまだ割り当てていない場合は、この通話制御アプリケーションに割り当てます。
どの電話からでも Telnyx 番号に電話をかけます。次のように聞こえます:
「ストーリー ホットラインへようこそ! 冒険を選択してください。ミステリーの場合は 1、SF の場合は 2、ファンタジーの場合は 3、ホラーの場合は 4、ロマンスの場合は 5 を押してください。」
キーを押します。 AI が最初の章を生成し、読み上げます。最後に 2 つの選択肢が表示されます。 1 または 2 を押すか、選択を話すと、ストーリーが続きます。
5 章が終わると AI がストーリーをまとめ、通話は終了します。
ストーリーテリング システム プロンプトはコントロール サーフェスです。小さな変更が非常に異なるエクスペリエンスを生み出します。
ジャンル = {"1": "ノワール探偵"、"2": "スペース オペラ"、"3": "ポスト黙示録"、"4": "ゴシック ホラー"、"5": "コージー ロマンス"}
章ごとに選択肢を追加します。
f「各章は、正確に 3 つの選択肢で終了します:「1 を押して...」、「2 を押して...」、または「3 を押して...」。
クイズ、セラピー セッション、トレーニング シナリオなど、別の形式にします。
f「あなたは、電話ホットラインのインタラクティブなコンプライアンス クイズの主催者です。章ごとに 1 つの多肢選択式の質問 (3 ～ 4 文) を出題します。「A の場合は 1、B の場合は 2、C の場合は 3 を押してください。」で終わります。 10 個の質問が終わったら、発信者を採点し、フィードバックを返します。」
このパターン (Webhook ステート マシン + 会話履歴を備えた LLM + 音声 I/O) は、分岐するあらゆる会話エクスペリエンスに機能します。
この例では、簡単にするためにインメモリ ストレージを使用します。運用環境の場合:
データベース - メモリ内の active_calls dict を Redis または PostgreSQL に置き換え、再起動後も呼び出し状態が維持されるようにします。
同時実行 — 複数のワーカーを使用して gunicorn の背後でアプリを実行します
エラー回復 - 推論タイムアウトを処理し、再試行または SM で正常に呼び出し失敗を処理します。

S フォールバック
プロンプトチューニング — ジャンルに合わせてさまざまなシステムプロンプトと温度設定をテストします
レート制限 - Webhook エンドポイントを悪用から保護します
モニタリング — 構造化されたログ記録と通話の成功/失敗率に関するアラートを追加します。
別の AI モデルを使用できますか?はい。 Telnyx AI Inference は OpenAI と互換性があります。 .env の AI_MODEL をプラットフォームで利用可能な任意のモデルに設定します。 Llama 3.3 70B Instruct は、クリエイティブなストーリーテリングに適したデフォルトです。
通話料金はいくらですか? Telnyx Call Control は 1 分ごとに料金がかかります。 AI 推論の価格は 1,000 トークンごとに設定されます。 5 章のストーリーでは通常、最大 2,000 トークンと最大 3 分の通話時間が使用され、通話ごとに数セントがかかります。
複数の人が同時に電話をかけることはできますか?はい。各呼び出しは、 call_control_id をキーとする active_calls 内に独自のエントリを取得します。アプリは干渉することなく同時通話を処理します。
発信者が選択しなかった場合はどうなりますか?収集ステップは、10 秒 (ジャンルの選択) または 20 秒 (ストーリーの選択) 後にタイムアウトします。アプリは発信者に 1 または 2 を押すよう再度促します。
発信者はキーを押す代わりに話すことができますか?はい。ストーリーの選択中、収集は DTMF と音声の両方を受け入れます。発信者はキーを押す代わりに、「one」または「two」と言うことができます。
低遅延の音声 AI を構築する準備はできていますか?
リアルタイム会話の未来を構築する開発者に参加しましょう

## Original Extract

Call a number, choose a genre, and listen to an AI-generated interactive story where your choices shape the narrative. 104 lines of Python.

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
Build an AI Phone Story Hotline — Interactive Storytelling with Telnyx Voice AI
Imagine calling a phone number, picking a genre — mystery, sci-fi, fantasy, horror, or romance — and listening to an AI narrate a story that adapts to your choices in real time. Press 1 to open the creaking door. Press 2 to check the window. The story branches, the AI continues, and every call is a new adventure.
This is the AI Phone Story Hotline — a 104-line Python app built with Telnyx Call Control and AI Inference. No game engine, no branching script, no pre-written dialogue trees. The AI generates the story as you go, and your phone keypad shapes what happens next.
In this walkthrough, you'll build it from scratch. Clone the repo, configure a phone number, and deploy in minutes.
A phone number that anyone can call to start an interactive story:
Caller dials in — Telnyx answers and greets them with a genre menu
Caller picks a genre — press 1–5 for mystery, sci-fi, fantasy, horror, or romance
AI generates chapter one — a 3–4 sentence story segment ending with two choices
Caller chooses — press 1 or 2, or speak the choice aloud
Story continues — AI generates the next chapter based on the choice, keeping full conversation context
After 5 chapters — the AI brings the story to a satisfying ending
The whole interaction is voice-driven: Text-to-Speech reads each chapter, and the caller responds with DTMF keypresses or speech. The AI model (Llama 3.3 70B via Telnyx AI Inference) handles the storytelling, and Call Control handles the telephony.
Most AI phone demos are business use cases — book a table, qualify a lead, reset a password. This one is different. It's creative AI — the model is writing fiction in real time, branching based on human input, and delivering it over a phone call. The storytelling format (short chapters, two choices each) keeps the AI responses tight and the call engaging.
It also demonstrates a pattern that works for any conversational AI app: webhook-driven state machine + LLM with conversation memory + voice I/O. Once you see how the story hotline works, you can swap the storytelling prompt for any domain — a choose-your-own-adventure onboarding flow, an interactive quiz, a branching training simulation.
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
Flask app (app.py, 104 lines)
│
├──► call.initiated → answer the call
├──► call.answered → TTS greeting + genre menu
├──► call.gather.ended → DTMF/speech input → AI Inference
├──► call.speak.ended → gather next choice → AI Inference
└──► call.hangup → cleanup session
│
▼
Telnyx AI Inference (Llama 3.3 70B)
│
▼
Story chapter (TTS back to caller)
The app is a state machine driven by Telnyx webhook events. Each event triggers the next action — answer, speak, gather, or hang up. The AI Inference call sits in the middle, generating story chapters from the running conversation history.
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-phone-story-hotline-python
cp .env.example .env
pip install -r requirements.txt
Edit .env with your credentials:
TELNYX_API_KEY=KEY0123456789ABCDEF # from portal.telnyx.com/api-keys
TELNYX_PUBLIC_KEY= # from portal.telnyx.com/api-keys (public key)
STORY_NUMBER=+13105551234 # your Telnyx phone number
AI_MODEL=meta-llama/Llama-3.3-70B-Instruct
Step 2: Understand the Code
The entire app is 104 lines in a single file: app.py . Here are the key pieces.
Webhook Signature Verification
Every Telnyx webhook is signed with an Ed25519 key. The app verifies the signature before trusting any event:
@app.route("/webhooks/voice", methods=["POST"])
def handle_voice():
try:
client.webhooks.unwrap(request.get_data(as_text=True), headers=dict(request.headers))
except Exception:
return jsonify({"error": "invalid signature"}), 401
This prevents spoofed webhook calls from triggering call actions.
Each call is tracked in an in-memory dict keyed by call_control_id :
active_calls = {}
The state machine handles five events:
Call initiated — store the call state and answer:
if event_type == "call.initiated" and p.get("direction") == "incoming":
active_calls[ccid] = {"state": "genre_select", "conversation": [], "chapters": 0}
client.calls.actions.answer(ccid)
Call answered — greet the caller with the genre menu using Text-to-Speech:
elif event_type == "call.answered":
client.calls.actions.speak(ccid,
payload="Welcome to Story Hotline! Choose your adventure. Press 1 for Mystery, 2 for Sci-Fi, 3 for Fantasy, 4 for Horror, 5 for Romance.",
voice="female", language_code="en-US")
Speak ended — after TTS finishes, gather the caller's input. During genre selection, gather DTMF digits. During the story, gather speech or DTMF:
elif event_type == "call.speak.ended" and call:
if call["state"] == "genre_select":
client.calls.actions.gather(ccid, input_type="dtmf", timeout_secs=10, min_digits=1, max_digits=1)
else:
client.calls.actions.gather(ccid, input_type="speech dtmf", end_silence_timeout_secs=3, timeout_secs=20, language_code="en-US")
The Storytelling Prompt
When the caller picks a genre, the app builds the system prompt that guides the AI for the rest of the call:
GENRES = {"1": "mystery", "2": "sci-fi", "3": "fantasy", "4": "horror", "5": "romance"}
if call["state"] == "genre_select":
genre = GENRES.get(digits, "mystery")
call["state"] = "story"
call["conversation"] = [{"role": "system", "content":
f"You are an interactive {genre} storyteller on a phone hotline. "
f"Tell a gripping story in short chapters (3-4 sentences each). "
f"End each chapter with exactly TWO choices: 'Press 1 to...' or 'Press 2 to...'. "
f"Make it vivid and cinematic. After 5 chapters, bring the story to a satisfying ending."
}]
story_start = call_inference(call["conversation"] + [{"role": "user", "content": "Begin the story."}])
call["conversation"].append({"role": "assistant", "content": story_start})
client.calls.actions.speak(ccid, payload=story_start, voice="female", language_code="en-US")
The prompt does three things:
Format constraint — short chapters (3–4 sentences) that fit naturally in a phone call
Choice structure — exactly two options ending with "Press 1" or "Press 2" so the gather step can capture DTMF
Ending condition — after 5 chapters, wrap up the story
The call_inference helper sends the full conversation history to Telnyx AI Inference and returns the model's response:
def call_inference(messages, max_tokens=250):
resp = requests.post(INFERENCE_URL,
headers={"Authorization": f"Bearer {TELNYX_API_KEY}", "Content-Type": "application/json"},
json={"model": AI_MODEL, "messages": messages, "max_tokens": max_tokens, "temperature": 0.9, },
timeout=20)
resp.raise_for_status()
return resp.json()["choices"][0]["message"]["content"]
Temperature is set to 0.9 — high enough for creative storytelling, low enough to keep the narrative coherent. The model is Llama 3.3 70B Instruct, available on Telnyx AI Inference with an OpenAI-compatible API.
Each time the caller makes a choice, the app appends it to the conversation and asks the AI for the next chapter:
elif call["state"] == "story":
choice = digits or speech
call["conversation"].append({"role": "user", "content": f"I choose option {choice}"})
call["chapters"] += 1
if call["chapters"] >= 5:
call["conversation"][-1]["content"] += ". Bring the story to a dramatic conclusion."
continuation = call_inference(call["conversation"])
call["conversation"].append({"role": "assistant", "content": continuation})
client.calls.actions.speak(ccid, payload=continuation, voice="female", language_code="en-US")
After 5 chapters, the app injects a final instruction to bring the story to an end. The AI sees the full conversation history on every call, so it maintains narrative continuity across all chapters.
python app.py
In a separate terminal, expose your local server with ngrok:
ngrok http 5000
Copy the HTTPS URL and configure it in the Telnx Portal :
Go to Call Control Applications
Create or edit your application
Set the Webhook URL to https://<your-ngrok-url>.ngrok.app/webhooks/voice
Assign your Telnyx phone number to this Call Control Application if you haven't already.
Call your Telnyx number from any phone. You'll hear:
"Welcome to Story Hotline! Choose your adventure. Press 1 for Mystery, 2 for Sci-Fi, 3 for Fantasy, 4 for Horror, 5 for Romance."
Press a key. The AI generates the first chapter and reads it aloud. At the end, you'll hear two choices. Press 1 or 2 — or speak your choice — and the story continues.
After 5 chapters, the AI wraps up the story and the call ends.
The storytelling system prompt is the control surface. Small changes produce very different experiences:
GENRES = {"1": "noir detective", "2": "space opera", "3": "post-apocalyptic", "4": "gothic horror", "5": "cozy romance"}
Add more choices per chapter:
f"End each chapter with exactly THREE choices: 'Press 1 to...', 'Press 2 to...', or 'Press 3 to...'."
Make it a different format — a quiz, a therapy session, a training scenario:
f"You are an interactive compliance quiz host on a phone hotline. Ask one multiple-choice question per chapter (3-4 sentences). End with 'Press 1 for A, 2 for B, 3 for C.' After 10 questions, score the caller and give feedback."
The pattern — webhook state machine + LLM with conversation history + voice I/O — works for any branching conversational experience.
This example uses in-memory storage for simplicity. For a production deployment:
Database — replace the in-memory active_calls dict with Redis or PostgreSQL so call state survives restarts
Concurrency — run the app behind gunicorn with multiple workers
Error recovery — handle inference timeouts and call failures gracefully with retry or SMS fallback
Prompt tuning — test different system prompts and temperature settings for your genre
Rate limiting — protect your webhook endpoint from abuse
Monitoring — add structured logging and alerting on call success/failure rates
Can I use a different AI model? Yes. Telnyx AI Inference is OpenAI-compatible. Set AI_MODEL in .env to any model available on the platform. Llama 3.3 70B Instruct is a good default for creative storytelling.
How much does a call cost? Telnyx Call Control is priced per minute. AI Inference is priced per 1K tokens. A 5-chapter story typically uses ~2,000 tokens and ~3 minutes of call time — a few cents per call.
Can multiple people call at the same time? Yes. Each call gets its own entry in active_calls keyed by call_control_id . The app handles concurrent calls without interference.
What happens if the caller doesn't choose? The gather step times out after 10 seconds (genre selection) or 20 seconds (story choices). The app re-prompts the caller to press 1 or 2.
Can the caller speak instead of pressing keys? Yes. During story choices, the gather accepts both DTMF and speech. The caller can say "one" or "two" instead of pressing keys.
Ready to build with low-latency voice AI?
Join developers building the future of real-time conversations
