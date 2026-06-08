---
source: "https://runapi.ai/"
hn_url: "https://news.ycombinator.com/item?id=48451431"
title: "Show HN: RunAPI – one API for AI video, image, music, audio, and LLMs"
article_title: "RunAPI - Unified AI API for Video, Music, Image &amp; LLM Generation | Runapi.Ai"
author: "billy42"
captured_at: "2026-06-08T21:38:18Z"
capture_tool: "hn-digest"
hn_id: 48451431
score: 1
comments: 0
posted_at: "2026-06-08T20:27:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RunAPI – one API for AI video, image, music, audio, and LLMs

- HN: [48451431](https://news.ycombinator.com/item?id=48451431)
- Source: [runapi.ai](https://runapi.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T20:27:30Z

## Translation

タイトル: Show HN: RunAPI – AI ビデオ、画像、音楽、オーディオ、LLM 用の 1 つの API
記事のタイトル: RunAPI - ビデオ、音楽、画像、LLM 生成のための統合 AI API |るなぴ・あい
説明: 1 つの統合 API を通じてビデオ、音楽、画像、LLM モデルにアクセスします。 Kling、Suno、Flux、Claude、Gemini、DeepSeek を、予測可能な価格設定と開発者に優しい SDK で統合します。
HN テキスト: いくつかの SaaS アプリを構築した後、多くの AI モデル API を同時に接続すると、すぐに面倒になってしまうことがわかりました。そこで、1 つの API キーで画像、ビデオ、音楽/オーディオ、LLM モデルを呼び出すことができる RunAPI を構築し、SDK、CLI、スキル、MCP サーバーもリリースしました。これにより、バイブコーディングでこれらのモデルをアプリに統合することが簡単になります。 Claude Code、Codex 内で直接使用できるほか、OpenClaw やHermes などのエージェント内でも使用できます。ぜひお試しいただき、フィードバックを共有していただければ幸いです。

記事本文:
ビデオ、音楽、画像、LLM 用の統合 AI API
Kling、Suno、Flux、Claude、Gemini、DeepSeek などにアクセスするための 1 つの API キー。従量課金制。生産準備完了。
すべてのシステムが稼働中
·
10以上のモデル
·
4つのSDK
category-tabs#switch" data-pane="claude" type="button">クロード コード
category-tabs#switch" data-pane="codex" type="button">Codex
category-tabs#switch" data-pane="gemini" type="button">Gemini CLI
るなぴ.ai
catalog-copy#copy" data-catalog-copy-source-param="hero-claude" type="button">コピー
npx スキル追加 runapi-ai/kling -g
# クロード コードで次のプロンプトを表示します。
RunAPI を使用して、Kling ビデオ生成をこの Next.js アプリに追加します。
catalog-copy#copy" data-catalog-copy-source-param="hero-codex" type="button">コピー
npx スキル追加 runapi-ai/suno -g
# Codex で、次のプロンプトを表示します。
RunAPI を使用して、この Next.js アプリに Suno 音楽生成を追加します。
catalog-copy#copy" data-catalog-copy-source-param="hero-gemini" type="button">コピー
npx スキル追加 runapi-ai/flux-kontext -g
# Gemini CLI で、次のプロンプトを表示します。
RunAPI を使用して、この Next.js アプリに Flux イメージ生成を追加します。
モデルスキルのインストール -> プロンプトコーディングエージェント -> モデル統合
スキル
クリング
スノ
フラックス
シードドリーム
ヴェオ
クロード
ジェミニ
ディープシーク
GPT画像
シーダンス
ナノバナナ
10 以上の AI サービスを 1 つの API に統合
ビデオモデルを選択しますか?
Seedance 2.0、Kling 3.0、および Veo 3.1 API の比較
マルチモデル AI インフラストラクチャの退屈な部分を処理します。
単一の API キーを介してビデオ、音楽、画像、LLM モデルにアクセスします。これには、Suno (他では利用できる公式 API はありません) や Kling ビデオ生成が含まれます。
実稼働ワークロード向けに構築されています。非同期タスク管理、Webhook コールバック、自動再試行、予測可能なクレジットベースの請求。 Python、Node.js、Ruby、Go の SDK。
使った分だけお支払いください。サブスクリプションや隠れた料金はありません。 API を呼び出す前に、各世代のコストを正確に確認してください。
Thr

ビデオを生成したり、音楽を作成したり、画像を作成したりするためのコード行。
カール -X POST https://runapi.ai/api/v1/kling/text_to_video \
-H "認可: ベアラー $RUNAPI_KEY" \
-H "コンテンツ タイプ: application/json" \
-d '{
"モデル": "kling-3.0",
"prompt": "ビーチを走るゴールデンレトリバー",
「期間」：「5」
}'
# { "task_id": "task_abc123", "status": "pending" }
catalog-copy#copy" data-catalog-copy-source-param="qs-music" type="button">コピー
カール -X POST https://runapi.ai/api/v1/suno/text_to_music \
-H "認可: ベアラー $RUNAPI_KEY" \
-H "コンテンツ タイプ: application/json" \
-d '{
"モデル": "suno-v5",
"prompt": "深夜のコーディングについての明るいインディー ロック ソング"
}'
# { "task_id": "task_def456", "status": "pending" }
catalog-copy#copy" data-catalog-copy-source-param="qs-image" type="button">コピー
カール -X POST https://runapi.ai/api/v1/flux_2/text_to_image \
-H "認可: ベアラー $RUNAPI_KEY" \
-H "コンテンツ タイプ: application/json" \
-d '{
"モデル": "flux-2-pro-text-to-image",
"prompt": "技術系スタートアップのミニマルなロゴ、きれいなベクター アート",
「幅」: 1024、
「高さ」：1024
}'
# { "task_id": "task_ghi789", "ステータス": "保留中" }
ビデオ、音楽、画像に対応する 1 つの API
非同期
構築を開始するには API キーを取得してください
よくある質問
無料のアカウントにサインアップし、ダッシュボードから API キーを生成し、最初のリクエストを行います。ほとんどのチームは 10 分以内に実用的な統合を出荷します。
RunAPI は、ビデオ生成 (Kling、Veo、Seedance)、音楽生成 (Suno)、画像生成 (Flux、Seedream、GPT Image)、および LLM (Claude、Gemini、DeepSeek) へのアクセスを提供します。新しいモデルが定期的に追加されます。
RunAPI は、従量課金制のクレジットベースのシステムを使用します。サブスクリプションや隠れた料金はありません。生成した分だけお支払いいただき、クレジットの有効期限が切れることはありません。
はい。公式 SDK は Python、Node.js、R で利用可能です

ユービーとゴー。 REST API を任意の HTTP クライアントで直接使用することもできます。
世代が失敗するとどうなりますか?
失敗した世代には料金はかかりません。システムは一時的な障害を自動的に再試行し、障害の詳細を含む Webhook 通知を受け取ります。
無料の API キーを取得して、数分で生成を開始できます。
APIを実行する
ビデオ、画像、音楽、LLM など、AI モデルごとに 1 つの API。

## Original Extract

Access video, music, image, and LLM models through one unified API. Integrate Kling, Suno, Flux, Claude, Gemini, and DeepSeek with predictable pricing and developer-friendly SDKs.

After building a few SaaS apps, I found that connecting many AI model APIs at the same time gets messy very quickly. So I built RunAPI, which lets you call image, video, music/audio, and LLM models with one API key, and I also released SDKs, a CLI, skills, and an MCP server. This makes it easy for vibe coding to integrate these models into your app. You can use it directly inside Claude Code, Codex, and even agents like OpenClaw and Hermes. I’d love for you to try it and share feedback.

Unified AI API for Video, Music, Image & LLMs
One API key to access Kling, Suno, Flux, Claude, Gemini, DeepSeek and more. Pay-as-you-go. Production-ready.
All systems operational
·
10+ models
·
4 SDKs
catalog-tabs#switch" data-pane="claude" type="button">Claude Code
catalog-tabs#switch" data-pane="codex" type="button">Codex
catalog-tabs#switch" data-pane="gemini" type="button">Gemini CLI
runapi.ai
catalog-copy#copy" data-catalog-copy-source-param="hero-claude" type="button">Copy
npx skills add runapi-ai/kling -g
# In Claude Code, prompt:
Add Kling video generation to this Next.js app with RunAPI.
catalog-copy#copy" data-catalog-copy-source-param="hero-codex" type="button">Copy
npx skills add runapi-ai/suno -g
# In Codex, prompt:
Add Suno music generation to this Next.js app with RunAPI.
catalog-copy#copy" data-catalog-copy-source-param="hero-gemini" type="button">Copy
npx skills add runapi-ai/flux-kontext -g
# In Gemini CLI, prompt:
Add Flux image generation to this Next.js app with RunAPI.
install model skill -> prompt coding agent -> model integration
skills
Kling
Suno
Flux
Seedream
Veo
Claude
Gemini
DeepSeek
GPT Image
Seedance
Nano Banana
10+ AI services, unified under one API
Choosing a video model?
Compare Seedance 2.0, Kling 3.0, and Veo 3.1 APIs
The boring parts of multi-model AI infrastructure, handled.
Access video, music, image, and LLM models through a single API key — including Suno (no official API available elsewhere) and Kling video generation.
Built for production workloads. Async task management, webhook callbacks, automatic retries, and predictable credit-based billing. SDKs in Python, Node.js, Ruby, and Go.
Pay only for what you use. No subscriptions, no hidden fees. See exactly what each generation costs before you call the API.
Three lines of code to generate a video, create music, or produce an image.
curl -X POST https://runapi.ai/api/v1/kling/text_to_video \
-H "Authorization: Bearer $RUNAPI_KEY" \
-H "Content-Type: application/json" \
-d '{
"model": "kling-3.0",
"prompt": "A golden retriever running on the beach",
"duration": "5"
}'
# { "task_id": "task_abc123", "status": "pending" }
catalog-copy#copy" data-catalog-copy-source-param="qs-music" type="button">Copy
curl -X POST https://runapi.ai/api/v1/suno/text_to_music \
-H "Authorization: Bearer $RUNAPI_KEY" \
-H "Content-Type: application/json" \
-d '{
"model": "suno-v5",
"prompt": "An upbeat indie rock song about coding late at night"
}'
# { "task_id": "task_def456", "status": "pending" }
catalog-copy#copy" data-catalog-copy-source-param="qs-image" type="button">Copy
curl -X POST https://runapi.ai/api/v1/flux_2/text_to_image \
-H "Authorization: Bearer $RUNAPI_KEY" \
-H "Content-Type: application/json" \
-d '{
"model": "flux-2-pro-text-to-image",
"prompt": "A minimalist logo for a tech startup, clean vector art",
"width": 1024,
"height": 1024
}'
# { "task_id": "task_ghi789", "status": "pending" }
One API for video, music, and image
async
Get your API key to start building
FAQ
Sign up for a free account, generate an API key from the dashboard, and make your first request. Most teams ship a working integration in under 10 minutes.
RunAPI provides access to video generation (Kling, Veo, Seedance), music generation (Suno), image generation (Flux, Seedream, GPT Image), and LLMs (Claude, Gemini, DeepSeek). New models are added regularly.
RunAPI uses a credit-based system with pay-as-you-go pricing. No subscriptions or hidden fees. You only pay for what you generate, and credits never expire.
Yes. Official SDKs are available for Python, Node.js, Ruby, and Go. You can also use the REST API directly with any HTTP client.
What happens when a generation fails?
Failed generations are not charged. The system automatically retries transient failures, and you receive a webhook notification with the failure details.
Get your free API key and start generating in minutes.
Run API
One API for every AI model — video, image, music, and LLMs.
