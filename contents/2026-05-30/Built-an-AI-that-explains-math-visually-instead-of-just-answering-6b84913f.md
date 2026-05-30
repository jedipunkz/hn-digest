---
source: "https://github.com/arzumanabbasov/claw-learn"
hn_url: "https://news.ycombinator.com/item?id=48328761"
title: "Built an AI that explains math visually instead of just answering"
article_title: "GitHub - arzumanabbasov/claw-learn: AI-powered visual math tutor, inspired by 3Blue1Brown. · GitHub"
author: "Arzuman"
captured_at: "2026-05-30T11:38:43Z"
capture_tool: "hn-digest"
hn_id: 48328761
score: 2
comments: 0
posted_at: "2026-05-29T20:26:26Z"
tags:
  - hacker-news
  - translated
---

# Built an AI that explains math visually instead of just answering

- HN: [48328761](https://news.ycombinator.com/item?id=48328761)
- Source: [github.com](https://github.com/arzumanabbasov/claw-learn)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T20:26:26Z

## Translation

タイトル: 数学をただ答えるだけでなく視覚的に説明する AI を構築
記事のタイトル: GitHub - arzumanabbasov/claw-learn: 3Blue1Brown からインスピレーションを得た、AI を活用したビジュアル数学家庭教師。 · GitHub
説明: 3Blue1Brown からインスピレーションを得た、AI を活用したビジュアル数学の家庭教師。 - アルズマナバソフ/爪学習

記事本文:
GitHub - arzumanabbasov/claw-learn: 3Blue1Brown からインスピレーションを得た、AI を活用したビジュアル数学家庭教師。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アルズマナバソフ
/
爪学習
公共
通知
通知設定を変更するにはサインインする必要があります
広告

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット アプリ アプリコンポーネント コンポーネント docs docs フック フック lib lib public パブリック タイプ タイプ .env.local.example .env.local.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md auth.ts auth.ts eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs proxy.ts proxy.ts tsconfig.json tsconfig.json vercel.json vercel.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claw Learn は、ElementalLabs 音声エンジンを搭載したリアルタイム音声インターフェイスを備えた、AI を活用したビジュアル数学家庭教師です。音声またはテキストで数学や物理の質問をすると、ブラウザ上でリアルタイムに生成される同期されたアニメーションの説明が表示されます。
ライブデモ · バグを報告 · 機能をリクエスト
Claw Learn は、イレブンラボの音声エンジンと AI シーン プランナーおよびカスタム キャンバス レンダラーを組み合わせて、数学の質問を同期ナレーション付きのライブ アニメーションの説明に変換します。
音声エンジンはエクスペリエンスの中核です。音声入力と音声出力の両方を WebRTC 経由で処理するため、キーボードに触れることなく、質問を話したり、説明の途中で中断したり、フォローアップをしたりすることができます。音声エンジンが構成されていない場合、アプリは REST TTS とブラウザベースの音声認識にフォールバックします。
スライドはありません。教科書はありません。事前に録画されたビデオはありません。すべての説明は、あなたの質問に正確に対応して新たに生成されます。
あなた: 「なぜ導関数は傾きを表すのですか?」
アプリ: → Celebrity Speech Engine が WebRTC 経由であなたの音声をキャプチャします。
→ AIが10シーンのビジュアル指導計画を生成
→C

anvas レンダリング: 軸、放物線、接線、傾斜公式
→ 音声エンジンがアニメーションと同期して各シーンのナレーションを行います
→ いつでも中断してフォローアップを求めてください - 話すだけです
デモ
「行列の乗算はどのように機能するのですか?」
「フーリエ変換を視覚的に説明する」
「積分とは何ですか?なぜ領域を求めるのですか?」
「オイラーの公式 e^(iπ) + 1 = 0 を見せて」
「重力はどのようにして軌道を作るのですか?」
レイヤー
テクノロジー
フレームワーク
Next.js 16 (アプリルーター、ターボパック)
UI
React 19、Tailwind CSS v4、Framer Motion
AI
任意の OpenAI 互換 API (Gemini、OpenAI、Ollama など)
音声入出力
イレブンラボ音声エンジン (WebRTC)
TTS フォールバック
イレブンラボ REST API
STT フォールバック
ウェブスピーチAPI
アニメーション
カスタム 2D Canvas レンダラー
言語
タイプスクリプト 5
導入
ヴェルセル
はじめに
OpenAI 互換 API キー — Gemini (aistudio.google.com で無料)、OpenAI、または互換性のあるプロバイダー
Google OAuth 認証情報 — ログインに必要 ( console.cloud.google.com )
Upstash Redis — レート制限に推奨 ( console.upstash.com 、無料枠)
イレブンラボ — elevenlabs.io のオプションの無料枠
git クローン https://github.com/arzumanabbasov/claw-learn.git
CD クローラーン
npmインストール
2. 環境変数を設定する
cp .env.local.example .env.local
.env.local を開いてキーを入力します。
# ── AI プロバイダー (必須) ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─
OPENAI_API_KEY = your_api_key_here
OPENAI_BASE_URL = https://generative language.googleapis.com/v1beta/openai
OPENAI_MODEL = ジェミニ-2.5-フラッシュ
# ── 認証 (必須) ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ 認証（必須）

─
# シークレットを生成します: openssl rand -base64 32
AUTH_SECRET = your_auth_secret_here
# Google OAuth — https://console.cloud.google.com/
GOOGLE_CLIENT_ID = your_google_client_id
GOOGLE_CLIENT_SECRET = your_google_client_secret
# ── レート制限 — Upstash Redis (推奨) ──--------------------------------------------------------
# これらがないと、レート制限はメモリ内にフォールバックします (サーバーの再起動時にリセットされます)。
# https://console.upstash.com/ で無料の Redis DB を作成します
UPSTASH_REDIS_REST_URL = https://your-db.upstash.io
UPSTASH_REDIS_REST_TOKEN = your_token_here
# ── イレブンラボボイス（オプション） ──── ──── ────
ELEVENLABS_API_KEY = your_elevenlabs_api_key_here
ELEVENLABS_VOICE_ID = pNInz6obpgDQGcFmaJgB
# 音声エンジン — 完全な WebRTC 音声 I/O
# https://elevenlabs.io/app/conversational-ai でエージェントを作成します
ELEVENLABS_SPEECH_ENGINE_ID = エージェント_xxxxxxxxxxxxxxxxxxxx
# ── セキュリティ ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─
ALLOWED_ORIGIN = https://your-domain.com
3. 走る
npm 実行開発
http://localhost:3000 を開きます。
Claw Learn は、Google OAuth で NextAuth.js v5 を使用します。すべてのルートには有効なセッションが必要です。認証されていないユーザーは /login にリダイレクトされます。
console.cloud.google.com → API とサービス → 認証情報に移動します
OAuth 2.0 クライアント ID の作成 (Web アプリケーション)
ドメインを承認された JavaScript オリジンに追加し、https://your-domain.com/api/auth/callback/google を承認されたリダイレクト URI に追加します。
クライアント ID とシークレットを環境変数にコピーします。
AUTH_S を生成する

openssl rand -base64 を使用した ECRET 32
ローカル開発の場合は、承認されたオリジンとして http://localhost:3000 を追加し、リダイレクト URI として http://localhost:3000/api/auth/callback/google を追加します。
認証された各ユーザーは 1 日あたり 3 つの質問を受け取り、Google ユーザー ID によって追跡され、UTC 午前 0 時にリセットされます。
レート制限では、実稼働環境で Upstash Redis (TTL が現在の UTC 日の終わりに設定されたアトミック INCR) を使用します。これはサーバーレスで安全であり、すべての Vercel エッジ インスタンスで機能します。
Upstash 資格情報がないと、アプリはサーバーが再起動するたびにリセットされるメモリ内ストアにフォールバックします (ローカル開発には問題ありませんが、運用環境には適していません)。
console.upstash.com で無料の Redis データベースを作成します。
REST URL とトークンを UPSTASH_REDIS_REST_URL と UPSTASH_REDIS_REST_TOKEN にコピーします。
残りの質問数は上部バーにライブバッジとして表示され、毎日自動的にリセットされます。
Claw Learn は OpenAI 互換の API 形式を使用するため、それをサポートするあらゆるプロバイダーで動作します。
OPENAI_API_KEY = your_gemini_api_key
OPENAI_BASE_URL = https://generative language.googleapis.com/v1beta/openai
OPENAI_MODEL = ジェミニ-2.5-フラッシュ
OpenAI
OPENAI_API_KEY = your_openai_api_key
OPENAI_BASE_URL = https://api.openai.com/v1
OPENAI_MODEL = gpt-4o
オラマ (地元)
OPENAI_API_KEY = オラマ
OPENAI_BASE_URL = http://localhost:11434/v1
OPENAI_MODEL = ラマ3.1
音声モード
モード 1 — イレブンラボ音声エンジン (推奨)
Speech Engine は WebRTC 経由で Celebrities Conversational AI エージェントに接続し、Claw Learn の主要な音声インターフェイスです。単一の低遅延接続で入力と出力の両方を処理します。
音声入力 — 入力する必要がなく、質問を自然に話します
ストリーミング TTS — 各シーンの再生時に イレブンラボからオーディオを直接ストリーミングします。
中断 — 説明の途中で話して、リダイレクトするかフォローアップを尋ねます
ロー

wer latency — WebRTC は REST フォールバックよりも大幅に高速です
elevenlabs.io/app/conversational-ai に移動します。
システム プロンプトを次のように設定します。「あなたは数学のナレーション音声です。ユーザーが送信した内容を明確で自然なナレーションとして正確に読んでください。」
エージェント ID をコピーし、.env.local で ELEVENLABS_SPEECH_ENGINE_ID として設定します。
上部バーの音声ボタンは、音声エンジンの接続と切断を行います。接続すると、緑色のインジケーターが点滅し、セッションがライブであることを示します。
ELEVENLABS_API_KEY が設定されていても音声エンジンが設定されていない場合、各シーンのナレーションが Celebrities REST API に送信され、オーディオとして再生されます。このモードでは音声入力はありません。
このアプリは、イレブンラボの設定を行わなくても、テキスト入力とサイレント アニメーションのみで完全に動作します。
npxヴェルセル
Vercel ダッシュボードの [設定] → [環境変数] で次の環境変数を設定します。
リポジトリ内の vercel.json は事前に構成されています。
npm ビルドを実行する
npmスタート
Node.js 18 以降と上記の環境変数が必要です。
キャンバス レンダラーは 30 以上の要素タイプをサポートしています。
座標系: 原点が中心、x が右、y が上。一般的な可視範囲: x ∈ [-6, 6]、y ∈ [-4, 4]。
クローラーン/
§── アプリ/
│ §── api/
│ │ §── Explain/route.ts # POST — AI シーンプラン生成
│ │ §── narrate/route.ts # POST — イレブンラボ REST TTS
│ │ └─ speech-engine/token/ # GET — WebRTC 会話トークン
│ §── page.tsx # ルート — 着陸 ↔ 家庭教師ルーター
│ §──layout.tsx # フォント、メタデータ、グローバルCSS
│ └── globals.css # トークン、アニメーションのデザイン
│
§── コンポーネント/
│ §── LandingPage.tsx # マーケティングページ
│ §── TutorApp.tsx # アプリシェル
│ §──AnimationCanvas.tsx # Canvas + シーンシーケンサー
│ §── ConversationPanel.tsx # チャット履歴
│ §── クエ

stionInput.tsx # 入力バー
│ └── NarrationSubtitle.tsx # キャンバス下の字幕
│
§── フック/
│ §── useTutor.ts # コアオーケストレーション
│ §── useSpeechEngine.ts # Celebrity Speech Engine (WebRTC)
│ └── useVoice.ts # Web Speech API フォールバック
│
§── lib/
│ §── openai.ts # OpenAI 互換クライアント + システムプロンプト
│ §── animeEngine.ts # Canvas レンダラー (30 以上の要素)
│ §── elevenlabs.ts # イレブンラボ REST ヘルパー
│ └─ voiceRecognition.ts # Web Speech API ラッパー
│
§── 種類/
│ └── scene.ts # シーンプラン TypeScript の種類
│
§── .env.local.example # 環境変数テンプレート
§── COTRIBUTING.md # 投稿ガイド
§── ライセンス # MIT
└── vercel.json # Vercel デプロイメント構成
セキュリティ
API キーはサーバー側のみであり、ブラウザーに公開されることはありません
入力は長さ制限があり、すべての API ルートで検証されます
CORS は本番環境では ALLOWED_ORIGIN にロックされています
キャンバス レンダラーは、安全な再帰降下数学パーサーを使用します。eval や new Function は使用しません。
セキュリティ ヘッダー ( X-Frame-Options 、 X-Content-Type-Options 、 Referrer-Policy 、 Permissions-Policy ) がすべての応答に設定されます
完全なセキュリティ ポリシーと脆弱性の報告方法については、SECURITY.md を参照してください。
永続性なし - 会話履歴はメモリ内にあり、ページを更新すると消去されます
音声入力 - Web Speech API フォールバックには Chrome または Edge が必要です。音声エンジンは最新のすべてのブラウザで動作します
JSON の切り捨て — 非常に複雑なトピックでは、AI が切り捨てられた JSON を返す可能性があります。パーサーは最後の完全なシーンを見つけて回復を試みます
イレブンラボの無料枠 — 10,000 文字/月。割り当てを超過すると、アプリはナレーションなしで静かに続行されます
貢献は大歓迎です。まず COTRIBUTING.md をお読みください。
＃フォーク、

n:
git checkout -b feat/your-feature
# 変更を加える
npx tsc --noEmit # 渡す必要があります
git commit -m " feat: あなたの機能 "
git Push Origin feat/your-feature
# プルリクエストを開く
謝辞
3Blue1Brown と manim アニメーション ライブラリからインスピレーションを得た
エレブで構築

[切り捨てられた]

## Original Extract

AI-powered visual math tutor, inspired by 3Blue1Brown. - arzumanabbasov/claw-learn

GitHub - arzumanabbasov/claw-learn: AI-powered visual math tutor, inspired by 3Blue1Brown. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
arzumanabbasov
/
claw-learn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits app app components components docs docs hooks hooks lib lib public public types types .env.local.example .env.local.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md auth.ts auth.ts eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs proxy.ts proxy.ts tsconfig.json tsconfig.json vercel.json vercel.json View all files Repository files navigation
Claw Learn is an AI-powered visual math tutor with a real-time voice interface — powered by the ElevenLabs Speech Engine. Ask any math or physics question by voice or text, and watch a synchronized animated explanation generate live in the browser.
Live Demo · Report a Bug · Request a Feature
Claw Learn combines the ElevenLabs Speech Engine with an AI scene planner and a custom canvas renderer to turn math questions into live animated explanations with synchronized narration.
The Speech Engine is the core of the experience — it handles both voice input and audio output over WebRTC, so you can speak your question, interrupt mid-explanation, and ask follow-ups without ever touching a keyboard. When the Speech Engine isn't configured, the app falls back to REST TTS and browser-based speech recognition.
No slides. No textbooks. No pre-recorded videos. Every explanation is generated fresh for your exact question.
You: "Why does the derivative represent slope?"
App: → ElevenLabs Speech Engine captures your voice over WebRTC
→ AI generates a 10-scene visual teaching plan
→ Canvas renders: axes, parabola, tangent line, slope formula
→ Speech Engine narrates each scene in sync with the animation
→ Interrupt at any time to ask a follow-up — just speak
Demo
"How does matrix multiplication work?"
"Explain the Fourier transform visually"
"What is integration and why does it find area?"
"Show me Euler's formula e^(iπ) + 1 = 0"
"How does gravity create orbits?"
Layer
Technology
Framework
Next.js 16 (App Router, Turbopack)
UI
React 19, Tailwind CSS v4, Framer Motion
AI
Any OpenAI-compatible API (Gemini, OpenAI, Ollama, etc.)
Voice I/O
ElevenLabs Speech Engine (WebRTC)
TTS fallback
ElevenLabs REST API
STT fallback
Web Speech API
Animations
Custom 2D Canvas renderer
Language
TypeScript 5
Deployment
Vercel
Getting Started
An OpenAI-compatible API key — Gemini (free at aistudio.google.com ), OpenAI, or any compatible provider
Google OAuth credentials — required for login ( console.cloud.google.com )
Upstash Redis — recommended for rate limiting ( console.upstash.com , free tier)
ElevenLabs — optional, free tier at elevenlabs.io
git clone https://github.com/arzumanabbasov/claw-learn.git
cd claw-learn
npm install
2. Configure environment variables
cp .env.local.example .env.local
Open .env.local and fill in your keys:
# ── AI Provider (required) ────────────────────────────────────────────────────
OPENAI_API_KEY = your_api_key_here
OPENAI_BASE_URL = https://generativelanguage.googleapis.com/v1beta/openai
OPENAI_MODEL = gemini-2.5-flash
# ── Auth (required) ───────────────────────────────────────────────────────────
# Generate a secret: openssl rand -base64 32
AUTH_SECRET = your_auth_secret_here
# Google OAuth — https://console.cloud.google.com/
GOOGLE_CLIENT_ID = your_google_client_id
GOOGLE_CLIENT_SECRET = your_google_client_secret
# ── Rate limiting — Upstash Redis (recommended) ───────────────────────────────
# Without these, rate limiting falls back to in-memory (resets on server restart)
# Create a free Redis DB at https://console.upstash.com/
UPSTASH_REDIS_REST_URL = https://your-db.upstash.io
UPSTASH_REDIS_REST_TOKEN = your_token_here
# ── ElevenLabs Voice (optional) ───────────────────────────────────────────────
ELEVENLABS_API_KEY = your_elevenlabs_api_key_here
ELEVENLABS_VOICE_ID = pNInz6obpgDQGcFmaJgB
# Speech Engine — full WebRTC voice I/O
# Create an agent at https://elevenlabs.io/app/conversational-ai
ELEVENLABS_SPEECH_ENGINE_ID = agent_xxxxxxxxxxxxxxxxxxxx
# ── Security ──────────────────────────────────────────────────────────────────
ALLOWED_ORIGIN = https://your-domain.com
3. Run
npm run dev
Open http://localhost:3000 .
Claw Learn uses NextAuth.js v5 with Google OAuth. All routes require a valid session — unauthenticated users are redirected to /login .
Go to console.cloud.google.com → APIs & Services → Credentials
Create an OAuth 2.0 Client ID (Web application)
Add your domain to Authorized JavaScript origins and https://your-domain.com/api/auth/callback/google to Authorized redirect URIs
Copy the Client ID and Secret into your env vars
Generate AUTH_SECRET with openssl rand -base64 32
For local dev, add http://localhost:3000 as an authorized origin and http://localhost:3000/api/auth/callback/google as a redirect URI.
Each authenticated user gets 3 questions per day , tracked by their Google user ID and reset at UTC midnight.
Rate limiting uses Upstash Redis in production — an atomic INCR with a TTL set to the end of the current UTC day. This is serverless-safe and works across all Vercel edge instances.
Without Upstash credentials, the app falls back to an in-memory store that resets whenever the server restarts (fine for local dev, not suitable for production).
Create a free Redis database at console.upstash.com
Copy the REST URL and token into UPSTASH_REDIS_REST_URL and UPSTASH_REDIS_REST_TOKEN
The remaining question count is shown in the top bar as a live badge and resets automatically each day.
Claw Learn uses the OpenAI-compatible API format, so it works with any provider that supports it.
OPENAI_API_KEY = your_gemini_api_key
OPENAI_BASE_URL = https://generativelanguage.googleapis.com/v1beta/openai
OPENAI_MODEL = gemini-2.5-flash
OpenAI
OPENAI_API_KEY = your_openai_api_key
OPENAI_BASE_URL = https://api.openai.com/v1
OPENAI_MODEL = gpt-4o
Ollama (local)
OPENAI_API_KEY = ollama
OPENAI_BASE_URL = http://localhost:11434/v1
OPENAI_MODEL = llama3.1
Voice Modes
Mode 1 — ElevenLabs Speech Engine (recommended)
The Speech Engine connects via WebRTC to an ElevenLabs Conversational AI agent and is the primary voice interface for Claw Learn. It handles both input and output in a single low-latency connection:
Voice input — speak your questions naturally, no typing needed
Streaming TTS — audio streams directly from ElevenLabs as each scene plays
Interruption — speak mid-explanation to redirect or ask a follow-up
Lower latency — WebRTC is significantly faster than the REST fallback
Go to elevenlabs.io/app/conversational-ai
Set the system prompt to: "You are a math narration voice. Read exactly what the user sends you as clear, natural narration."
Copy the Agent ID and set it as ELEVENLABS_SPEECH_ENGINE_ID in your .env.local
The Voice button in the top bar connects and disconnects the Speech Engine. When connected, a pulsing green indicator shows the session is live.
When ELEVENLABS_API_KEY is set but no Speech Engine is configured, each scene's narration is sent to the ElevenLabs REST API and played back as audio. No voice input in this mode.
The app works fully without any ElevenLabs configuration — text input and silent animations only.
npx vercel
Set these environment variables in the Vercel dashboard under Settings → Environment Variables :
The vercel.json in the repo is pre-configured.
npm run build
npm start
Requires Node.js 18+ and the environment variables above.
The canvas renderer supports 30+ element types:
Coordinate system: origin at center, x right, y up. Typical visible range: x ∈ [-6, 6], y ∈ [-4, 4].
clawlearn/
├── app/
│ ├── api/
│ │ ├── explain/route.ts # POST — AI scene plan generation
│ │ ├── narrate/route.ts # POST — ElevenLabs REST TTS
│ │ └── speech-engine/token/ # GET — WebRTC conversation token
│ ├── page.tsx # Root — landing ↔ tutor router
│ ├── layout.tsx # Fonts, metadata, global CSS
│ └── globals.css # Design tokens, animations
│
├── components/
│ ├── LandingPage.tsx # Marketing page
│ ├── TutorApp.tsx # App shell
│ ├── AnimationCanvas.tsx # Canvas + scene sequencer
│ ├── ConversationPanel.tsx # Chat history
│ ├── QuestionInput.tsx # Input bar
│ └── NarrationSubtitle.tsx # Subtitle below canvas
│
├── hooks/
│ ├── useTutor.ts # Core orchestration
│ ├── useSpeechEngine.ts # ElevenLabs Speech Engine (WebRTC)
│ └── useVoice.ts # Web Speech API fallback
│
├── lib/
│ ├── openai.ts # OpenAI-compatible client + system prompt
│ ├── animationEngine.ts # Canvas renderer (30+ elements)
│ ├── elevenlabs.ts # ElevenLabs REST helpers
│ └── voiceRecognition.ts # Web Speech API wrapper
│
├── types/
│ └── scene.ts # Scene plan TypeScript types
│
├── .env.local.example # Environment variable template
├── CONTRIBUTING.md # Contribution guide
├── LICENSE # MIT
└── vercel.json # Vercel deployment config
Security
API keys are server-side only — never exposed to the browser
Input is length-limited and validated on every API route
CORS is locked to ALLOWED_ORIGIN in production
The canvas renderer uses a safe recursive-descent math parser — no eval or new Function
Security headers ( X-Frame-Options , X-Content-Type-Options , Referrer-Policy , Permissions-Policy ) are set on all responses
See SECURITY.md for the full security policy and how to report vulnerabilities.
No persistence — conversation history is in-memory, cleared on page refresh
Voice input — Web Speech API fallback requires Chrome or Edge; the Speech Engine works in all modern browsers
JSON truncation — very complex topics may cause the AI to return truncated JSON; the parser attempts recovery by finding the last complete scene
ElevenLabs free tier — 10,000 characters/month; the app continues silently without narration when quota is exceeded
Contributions are welcome. Please read CONTRIBUTING.md first.
# Fork, then:
git checkout -b feat/your-feature
# Make changes
npx tsc --noEmit # must pass
git commit -m " feat: your feature "
git push origin feat/your-feature
# Open a pull request
Acknowledgments
Inspired by 3Blue1Brown and the manim animation library
Built with the Eleve

[truncated]
