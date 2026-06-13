---
source: "https://whissle.ai/gateway"
hn_url: "https://news.ycombinator.com/item?id=48514802"
title: "Whissle Gateway – Run Multi-Modal Voice AI Locally in a 500MB Docker"
article_title: "Whissle Gateway — Run VoiceAI Locally"
author: "ksingla025"
captured_at: "2026-06-13T10:04:08Z"
capture_tool: "hn-digest"
hn_id: 48514802
score: 2
comments: 0
posted_at: "2026-06-13T08:16:09Z"
tags:
  - hacker-news
  - translated
---

# Whissle Gateway – Run Multi-Modal Voice AI Locally in a 500MB Docker

- HN: [48514802](https://news.ycombinator.com/item?id=48514802)
- Source: [whissle.ai](https://whissle.ai/gateway)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T08:16:09Z

## Translation

タイトル: Whissle Gateway – 500MB Docker でマルチモーダル音声 AI をローカルで実行する
記事のタイトル: Whisle Gateway — VoiceAI をローカルで実行する
説明: 自己ホスト型 VoiceAI。 1 つの Docker コマンド。 ASR、TTS、日記、メタデータ、AI コーチング。 CPU または GPU の 6 つのバリエーション。モデルは自動的にダウンロードされます。

記事本文:
Whissle Gateway — VoiceAI をローカルで実行 Whissle — 研究、音声、日常業務用のパーソナル AI 🔒 サービスに関するお知らせ: 当社のクラウド サービスは一時的に停止しています。その間、オンプレミス AI を強化しています。こんにちは@whissle.ai
🔒 サービスに関するお知らせ: クラウド サービスは一時的に停止しています - オンプレミス AI を強化しています。連絡先: hello@whissle.ai 🔒 サービスに関するお知らせ: クラウド サービスは一時的に停止しています - オンプレミス AI を強化しています。連絡先: hello@whissle.ai
Product Edge AI Whissle Browser アンビエント音声インテリジェンスを備えた AI ブラウザ macOS アプリ オンデバイス AI プラットフォームを備えたデスクトップ アプリ プラットフォーム Whissle ゲートウェイ シングル GPU 上のセルフホスト型音声 AI Agents Studio クラウド上でマルチモーダル エージェントを構築および展開 - 近日提供予定
ASR、TTS、音声通話、日記作成、メタデータ、AI コーチング - 1 つの Docker コマンド。モデルは自動的にダウンロードされます。クラウドへの依存はありません。
$ docker run -d --name whissle \
-p 9000:9000 -p 8001:8001 -p 8003:8003 \
-v whissle-models:/models -v whissle-data:/data \
-e VARIANT=en-full \
-e ANTHROPIC_API_KEY=あなたのキー \
whissleasr/whissle-gateway:latest VARIANT= hinglish en-lite en-full multi-full multi-zh all DEVICE= cpu cuda en-full · 初回実行時に最大 2 GB をダウンロードします (後でキャッシュされます)
═======================= ═===========================
Whissle Gateway — en-full
═======================= ═===========================
GPU が検出されません → CPU を使用しています
共有モデル:
✓ スピーカーエンコーダー + VAD 26 MB
✓ 句読点 254 MB
✓ ITN (英語 + ヒングリッシュ) 1.5 MB
バリアント: en-full
✓ en-in-tech-misc (485 MB)
✓ KenLM 英語 (1.5 GB)
認証:
モード: ローカル
トークン: wh_a1b2c3d4e5f6... (管理者)
管理:curl -H '認可:

ベアラー ...' localhost:9000/auth/tokens
サービスを開始しています...
PostgreSQL: :5432 ●
ASR: :8001 ●
TTS: :8003 ●
エージェント: :8765 ●
パイプキャット: :8000 ●
ゲートウェイ: :9000 ● API
5 つのインターフェイス - バッチ REST、ストリーミング WebSocket、テキスト読み上げ、音声通話、インテリジェント エージェント。
POST ローカルホスト:8001/transcribe
$カール -X POST http://localhost:8001/transcribe \
-F "ファイル=@call.mp3" \
-F "diarize=true" \
-F "スピーカー数=2" \
-F "句読点=true" \
-F "metadata_prob=true" \
-F "要約=セールス_コーチング" \
-o result.json レスポンス — トランスクリプト + セグメントごとのメタデータ + AI 分析
{
「セグメント」: [
{
"スピーカー" : "SPEAKER_00" ,
"text" : "こんにちは、おはようございます。" 、
「開始」：1.0、「終了」：1.9、
「メタデータ」: {
"感情" : "EMOTION_NEUTRAL" ,
"動作" : "BEHAVIOR_DIRECT" ,
"役割" : "ROLE_INTERVIEWER" ,
"年齢" : "AGE_30_45" ,
"性別" : "GENDER_MALE"
}、
"単語" : [{ "単語" : "こんにちは", "開始": 1.0, "終了": 1.3}]
}
]、
「分析」: {
「総合スコア」：78、
"buyer_outcome" : "変換済み" ,
"実践" : { "フォロー中": 6, "合計": 8 },
「ハイライト」 : [...]
}
パラメータ
POST /transcribe のすべてのパラメータ。
任意の文字起こしに -F "summarize=mode" を追加します。日記化されたトランスクリプトとメタデータは、分析のためにクロードまたはジェミニに送信されます。
8 つのベストプラクティスがスコア化されました。担当者/購入者の識別情報。タイムスタンプ付きで強調表示されます。セグメントごとの動作ラベル。総合スコアは 0 ～ 100。
本人確認、理由の記載、金額の記載、嫌がらせなし。通話結果 (支払いの約束/紛争/困難)。次のアクション。
概要、参加者、主要なトピック、感情的なダイナミクス、エンティティ、結果。マークダウン形式。
任意のプロンプト文字列を渡します。 LLM は、指示とセグメントごとのメタデータを含む完全なトランスクリプトを受け取ります。
各モデルは、単一の ASR フォワード パスで異なるメタデータを抽出します。個別のモデルや API 呼び出しは必要ありません。
120Mp

arams、コーチング、セラピー、インタビューのための 26 の行動規範。 8つの評価ラベル。
1 億 1500 万パラメータ、債権回収の意図 - 返済、紛争、苦難。エージェント/顧客の役割の検出。
ヒンディー語 - 英語 · 5 頭、26 クラス
160M パラメータ、北/南方言検出付き中国語。
マンダリン · 3 頭、12 クラス
600M パラメータ、インライン アクション トークン。 31 の意図グループ、18,000 の語彙。
23 言語 · 5,500 以上のアクション トークン
非自己回帰的なテキスト読み上げ。 CPU では 200ms 未満の TTFB。常に含まれています。
句読点の復元と逆テキスト正規化。
英語 + ヒングリッシュ · 自動ダウンロード
すべてのセグメントにはこれらのタグが含まれます。すべてのモデルに共通のタグが表示されます。追加のタグはモデルによって異なります。
言語と品質のニーズに基づいてバリエーションを選択してください。 VARIANT= を変更して再起動することで切り替えます。キャッシュされたモデルは再利用されます。
ラップトップ (CPU) からデータセンターの GPU まで。同じ Docker、同じ API。 GPUを自動検出します。
┌─────────────────────────┐
│ Docker コンテナ │
│ │
│ ┌─────┐ ┌─────┐ ┌─────┐ ┌───────┐ │
│ │ ASR │ │ TTS │ │ Pipecat │ │ エージェント │ │
│ │ :8001 │ │ :8003 │ │ :8000 │ │ :8765 │ │
│ │ │ │ ココロ │ │ │ │ クロード / │ │
│ │ ONNX │ │ 82M │ │ WebRTC │ │ Gemini API │ │
│ │ +KenLM │ │ 55 音声 │ │ Twilio │ │ │
│ │ +ECAPA │ │ │ 音声AI │ │ 要約 │ │
│ │ +VAD │ │ │ │ │ コーチ │ │
│ │ +パンクト │ │ │ 認証 │ │ 分析 │ │
│ │ +ITN │

│ │ │ マルチ組織│ │ │
│ ━━━━━━━━━━━━━━━━━━┘ ━━━━━━━┘ │
│ │ │
│ ┌─────────┐ │
│ │ PostgreSQL │ │
│ │ :5432 │ │
│ ━━━━━━━┘ │
│ │
│ /models (Docker ボリューム - キャッシュされた ASR モデル) │
│ /data (Docker ボリューム — PostgreSQL、認証、会話) │
━━━━━━━━━━━━━━━━━━━━━┘ ホイッスルモデルの巻
ASR モデル、KenLM、句読点、ITN。最初の実行時にダウンロードされ、永久にキャッシュされます。コンテナーの再起動後も存続します。
会話、分析、エージェント構成、認証トークン。再起動後も持続します。 docker volume rm によってのみ削除されます。
コマンドは 1 つです。モデルは自動的にダウンロードされます。 2分で準備完了。
コンタクト センター、セールス インテリジェンス、行動 AI などのために構築されています。
あなた専用の AI に出会う準備はできましたか?
ブラウザをダウンロードし、Web アプリを試すか、オープンソース、自己ホスト型、プライバシー最優先の API を使用して構築します。
セルフホスト ゲートウェイ 製品 Whissle ブラウザ Whissle ゲートウェイ macOS アプリ セルフホスティング 開発者 開発者ドキュメント クイック スタート オープンソース ソリューション コンタクト センター セールス インテリジェンス 行動 AI 3D アバター エージェント 会社概要 研究ブログ プライバシー © 2026 Whissle Inc. 全著作権所有。

## Original Extract

Self-hosted VoiceAI. One Docker command. ASR, TTS, diarization, metadata, AI coaching. 6 variants, CPU or GPU. Models download automatically.

Whissle Gateway — Run VoiceAI Locally Whissle — Personal AI for Research, Voice, and Everyday Tasks 🔒 Service Notice: Our cloud services are temporarily down — meanwhile we're reinforcing our on-prem AI. hello@whissle.ai
🔒 Service Notice: Cloud services temporarily down — reinforcing our on-prem AI. Contact: hello@whissle.ai 🔒 Service Notice: Cloud services temporarily down — reinforcing our on-prem AI. Contact: hello@whissle.ai
Product Edge AI Whissle Browser AI browser with ambient voice intelligence macOS App Desktop app with on-device AI Platform Whissle Gateway Self-hosted voice AI on a single GPU Agents Studio Build and deploy multi-modal agents on cloud — coming soon
ASR, TTS, voice calling, diarization, metadata, AI coaching — one Docker command. Models download automatically. No cloud dependency.
$ docker run -d --name whissle \
-p 9000:9000 -p 8001:8001 -p 8003:8003 \
-v whissle-models:/models -v whissle-data:/data \
-e VARIANT=en-full \
-e ANTHROPIC_API_KEY=your-key \
whissleasr/whissle-gateway:latest VARIANT= hinglish en-lite en-full multi-full multi-zh all DEVICE= cpu cuda en-full · Downloads ~2 GB on first run (cached after)
═══════════════════════════════════════════════
Whissle Gateway — en-full
═══════════════════════════════════════════════
No GPU detected → using CPU
Shared models:
✓ speaker encoder + VAD 26 MB
✓ punctuation 254 MB
✓ ITN (English + Hinglish) 1.5 MB
Variant: en-full
✓ en-in-tech-misc (485 MB)
✓ KenLM ENGLISH (1.5 GB)
Auth:
Mode: local
Token: wh_a1b2c3d4e5f6... (admin)
Manage: curl -H 'Authorization: Bearer ...' localhost:9000/auth/tokens
Starting services...
PostgreSQL: :5432 ●
ASR: :8001 ●
TTS: :8003 ●
Agent: :8765 ●
Pipecat: :8000 ●
Gateway: :9000 ● API
Five interfaces — batch REST, streaming WebSocket, text-to-speech, voice calling, and an intelligent agent.
POST localhost:8001/transcribe
$ curl -X POST http://localhost:8001/transcribe \
-F "file=@call.mp3" \
-F "diarize=true" \
-F "num_speakers=2" \
-F "punctuation=true" \
-F "metadata_prob=true" \
-F "summarize=sales_coaching" \
-o result.json Response — transcript + metadata per segment + AI analysis
{
"segments" : [
{
"speaker" : "SPEAKER_00" ,
"text" : "Hello, good morning." ,
"start" : 1.0, "end" : 1.9,
"metadata" : {
"emotion" : "EMOTION_NEUTRAL" ,
"behavior" : "BEHAVIOR_DIRECT" ,
"role" : "ROLE_INTERVIEWER" ,
"age" : "AGE_30_45" ,
"gender" : "GENDER_MALE"
},
"words" : [{ "word" : "Hello", "start": 1.0, "end": 1.3}]
}
],
"analysis" : {
"overall_score" : 78,
"buyer_outcome" : "Converted" ,
"practices" : { "followed": 6, "total": 8 },
"highlights" : [...]
}
} Parameters
All parameters for POST /transcribe .
Add -F "summarize=mode" to any transcription. The diarized transcript + metadata is sent to Claude or Gemini for analysis.
8 best practices scored. Rep/buyer identification. Highlights with timestamps. Behavior labels per segment. Overall score 0–100.
Identity verification, reason stated, amount mentioned, no harassment. Call outcome (Promise to Pay / Dispute / Hardship). Next action.
Overview, participants, key topics, emotional dynamics, entities, outcome. Markdown format.
Pass any prompt string. The LLM receives your instructions + full transcript with per-segment metadata.
Each model extracts different metadata in a single ASR forward pass — no separate models or API calls.
120M params, 26 Behavioral codes for coaching, therapy, interviews. 8 evaluation labels.
115M params, Debt collection intents — pay-back, disputes, hardship. Agent/Customer role detection.
Hindi-English · 5 heads, 26 classes
160M params, Mandarin with North/South dialect detection.
Mandarin · 3 heads, 12 classes
600M params, inline action tokens. 31 intent groups, 18K vocabulary.
23 languages · 5,500+ action tokens
Non-autoregressive text-to-speech. Sub-200ms TTFB on CPU. Always included.
Punctuation restoration and inverse text normalization.
EN + Hinglish · Auto-downloaded
Every segment includes these tags. Common tags appear on all models. Additional tags depend on the model.
Choose your variant based on language and quality needs. Switch by changing VARIANT= and restarting. Cached models are reused.
From your laptop (CPU) to data center GPUs. Same Docker, same API. Auto-detects GPU.
┌──────────────────────────────────────────────────────────────┐
│ Docker Container │
│ │
│ ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────┐ │
│ │ ASR │ │ TTS │ │ Pipecat │ │ Agent │ │
│ │ :8001 │ │ :8003 │ │ :8000 │ │ :8765 │ │
│ │ │ │ Kokoro │ │ │ │ Claude / │ │
│ │ ONNX │ │ 82M │ │ WebRTC │ │ Gemini API │ │
│ │ +KenLM │ │ 55 voice │ │ Twilio │ │ │ │
│ │ +ECAPA │ │ │ │ Voice AI │ │ Summarize │ │
│ │ +VAD │ │ │ │ │ │ Coach │ │
│ │ +Punct │ │ │ │ Auth │ │ Analyze │ │
│ │ +ITN │ │ │ │ Multi-org│ │ │ │
│ └──────────┘ └──────────┘ └──────────┘ └──────────────┘ │
│ │ │
│ ┌──────────────┐ │
│ │ PostgreSQL │ │
│ │ :5432 │ │
│ └──────────────┘ │
│ │
│ /models (Docker volume — cached ASR models) │
│ /data (Docker volume — PostgreSQL, auth, conversations) │
└──────────────────────────────────────────────────────────────┘ whissle-models volume
ASR models, KenLM, punctuation, ITN. Downloaded on first run, cached forever. Survives container restarts.
Conversations, analytics, agent configs, auth tokens. Persists across restarts. Only deleted by docker volume rm .
One command. Models download automatically. Ready in 2 minutes.
Built for contact centers, sales intelligence, behavioral AI, and more.
Ready to meet your personal AI?
Download the browser, try the web app, or build with our APIs — open source, self-hostable, and privacy-first.
Self-Host Gateway Product Whissle Browser Whissle Gateway macOS App Self-Hosting Developers Developer Docs Quick Start Open Source Solutions Contact Centers Sales Intelligence Behavioral AI 3D Avatar Agents Company About Us Research Blog Privacy © 2026 Whissle Inc. All rights reserved.
