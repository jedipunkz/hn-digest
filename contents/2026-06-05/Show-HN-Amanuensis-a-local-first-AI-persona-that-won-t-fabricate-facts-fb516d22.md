---
source: "https://github.com/msalsas/amanuensis"
hn_url: "https://news.ycombinator.com/item?id=48414891"
title: "Show HN: Amanuensis – a local-first AI persona that won't fabricate facts"
article_title: "GitHub - msalsas/amanuensis: A local-first AI persona for Mastodon and Bluesky — drafts by machine, approved by a human, nothing it can't ground gets published. · GitHub"
author: "msalsas"
captured_at: "2026-06-05T18:46:34Z"
capture_tool: "hn-digest"
hn_id: 48414891
score: 1
comments: 0
posted_at: "2026-06-05T16:34:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Amanuensis – a local-first AI persona that won't fabricate facts

- HN: [48414891](https://news.ycombinator.com/item?id=48414891)
- Source: [github.com](https://github.com/msalsas/amanuensis)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T16:34:28Z

## Translation

タイトル: Show HN: Amanuensis – 事実を捏造しないローカルファーストの AI ペルソナ
記事のタイトル: GitHub - msalsas/amanuensis: マストドンとブルースカイのためのローカルファースト AI ペルソナ — 機械による草案、人間による承認、根拠のないものは公開されます。 · GitHub
説明: マストドンとブルースカイのためのローカルファーストの AI ペルソナ — 機械によって草稿され、人間によって承認され、根拠のないものは公開されます。 - ムサルサス/アマヌエンシス

記事本文:
GitHub - msalsas/amanuensis: Mastodon と Bluesky のためのローカルファースト AI ペルソナ — 機械による草稿、人間による承認、根拠のないものは公開されません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
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
ムサルサ
/
アマヌエンシス
公共
通知
c にサインインする必要があります

ハンゲの通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット アダプター アダプター config config core core docs/ img docs/ img profiles/ alexa profiles/ alexa scripts scripts services services social social testing testing .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md main_batch.py main_batch.py main_dispatcher.py main_dispatcher.py main_reply_listener.py main_reply_listener.py main_telegram_listener.py main_telegram_listener.py Manual_post.py Manual_post.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
人間の拒否権を無視して執筆する、ローカルファーストの AI ペルソナ。草稿を作成し、あなたが承認すれば、根拠のないものは出版されます。
Mastodon と Bluesky で AI ペルソナを実行するためのローカルファーストのパイプライン。パイロットのペルソナである AlexaPavlova は、技術ニュースやオープンソースについて辛口な意見を述べ、皮肉を込めたベルリンの上級開発者として投稿します。
すべてはローカル GPU マシン上で実行されます。クラウド LLM 呼び出しはありません。
ライブでご覧ください: Mastodon · Bluesky — 現在は AI として公開されており、投稿は停止されています。
すべての投稿は公開前に携帯電話で確認され、承認、テキストや画像の再生成、キャンセルが行われます。
ステータス: これは実験であり、実際の製品ではありません。コードは MIT ライセンスであり、エンドツーエンドで動作します。フォークして、そこから学び、独自のペルソナを実行します。問題や PR については返答が得られない場合があります。
難しい部分はテキストを生成することではなく、モデルが技術的な詳細を組み立てることを停止することでした。短いバージョン: 事実のみのソース概要、LLM 判定前の決定論的なクリーンアップ、LLM グラウンディング チェック前の正規表現の事前スクリーニング、タイトルのみのメモリ、Telegram 上のすべての投稿を人間が承認する。
設計と途中で何が壊れたかの完全な書き込み: 書き込み

。
pip install -e " .[dev] "
cp .env.example .env
2. ローカルサービスを開始する
LMStudio をダウンロードし、命令調整されたモデル (Mistral-7B-Instruct などでテスト済み) をロードします。
[ローカル サーバー] に移動し、ポート 1234 でサーバーを起動します。
.env で LMSTUDIO_BASE_URL=http://localhost:1234 を設定します。
SwarmUI をインストールし、イメージ モデル + 41ex4_p4v10v4 LoRA をロードします
.env に SWARMUI_BASE_URL=http://localhost:7801 を設定します。
メッセージ @BotFather → /newbot → トークンを TELEGRAM_BOT_TOKEN にコピーします
@userinfobot にメッセージを送信 → 数値 ID を TELEGRAM_CHAT_ID にコピーします
新しいボットにメッセージを送信して、メッセージを返信できるようにします
python main_batch.py --dry-run
これにより、実際のストーリーが取得され、ローカル サービスを使用して投稿と画像が生成されます。データベースには何も書き込まれず、Telegram にも何も送信されません。これで 8 件の投稿が出力される場合、ローカル スタックは機能しています。
ソーシャル認証情報を .env (Mastodon および/または Bluesky — どちらもオプション、以下の表を参照) に追加し、3 つのターミナルを開きます。
# ターミナル 1 — 今日の投稿を生成し、承認のために Telegram に送信します
Python main_batch.py
# ターミナル 2 — Telegram の承認を聞く
Python main_telegram_listener.py
# ターミナル 3 — 承認された投稿をスケジュールされた時間に公開します
Python main_dispatcher.py
Telegram の投稿を承認します。ディスパッチャーがそれらを取得して公開します。終わり。
長時間実行されるセットアップの場合は、再起動しても存続できるように、systemd またはvisord の下で 3 つの永続プロセスを実行します。 main_batch.py​​ はワンショット スクリプトです。cron 経由で、または毎日手動で実行します。
ペルソナの画像は、SwarmUI を通じて生成された、Juggernaut XL "Ragnarok" (SDXL) 上でトレーニングされたカスタム LoRA から取得されます。 LoRA はこのリポジトリには含まれていません。これにはトレーニングされたデルタのみが含まれており、基本モデルは含まれていません。
LoRA: Hugging Face — msalsas/alexa-lora からダウンロードします。トリガーワード 41ex4_p4v10v4 、ウェイ

ght 0.3 、768×1024 で生成。
ベースモデル: RunDiffusion の Juggernaut XL "Ragnarok" — ここでは配布されず、別途入手してください。
LoRA は完全に合成されたデータセット (Juggernaut XL で生成された画像) でトレーニングされました。このキャラクターは実在の人物に基づいたものではありません。
画像を含む alexa プロファイルを実行するには、SwarmUI に Juggernaut XL をロードし、この LoRA を適用する両方が必要です。正確なプロンプト形式については、Hugging Face のモデル カードを参照してください。
アダプター (HN、ロブスター、BearBlog、AskHN)
└── キュレーター (URL + タイトル + サブレディットによる削除、禁止トピック フィルター)
━── バッチファクトリー
§── Brain (LMStudio → 投稿テキスト + 画像プロンプト)
└── ImageService (SwarmUI → LoRA経由のPNG)
└── スケジューラー (ジッターのある UTC 時間ウィンドウ)
━── QueueService（SQLite）
└── TelegramNotifier (写真 + 承認キーボード)
§── マストドンパブリッシャー（承認中）
└── BlueskyPublisher (承認中)
応答パイプライン (並列実行):
ReplyListener (ランダム投票、投稿ごとに 30 分～2 時間)
§── MastodonCommentFetcher / BlueskyCommentFetcher
§── Brain.evaluate_relevance() → スキップまたは返信の下書き
§── Brain.generate_reply()
└── TelegramNotifier (REPLY_APPROVE / REPLY_CANCEL)
§── MastodonPublisher.publish_reply() (承認時)
└── BlueskyPublisher.publish_reply() (承認時)
要件
ローカルで実行される LMStudio (OpenAI 互換 API)
41ex4_p4v10v4 LoRA がロードされた状態でローカルで実行される SwarmUI
Telegram ボット トークン + チャット ID (承認ワークフロー用)
Mastodon および/または Bluesky 認証情報 (公開用)
OpenWeatherMap API キー (プロンプト内のアンビエント コンテキスト用の無料利用枠)
pip install -e " .[dev] "
cp .env.example .env
# .env にトークンとサービス URL を入力します
.env リファレンス
変数
説明
LMSTUDIO_BASE_URL
LMStudio API ベース、例: http://localhos

t:1234
SWARMUI_BASE_URL
SwarmUI ベース、例: http://ローカルホスト:7801
TELEGRAM_BOT_TOKEN
@BotFather からのボット トークン
TELEGRAM_CHAT_ID
個人チャット ID (@userinfobot を使用して見つけます)
アクティブ_プロフィール
プロファイルスラッグ、デフォルトのアレクサ
WEATHER_API_KEY
OpenWeatherMapキー（無料）
ALEXA_MASTODON_ACCESS_TOKEN
Alexa プロファイルの Mastodon トークン
ALEXA_BLUESKY_APP_PASSWORD
AlexaプロファイルのBlueskyアプリパスワード
マストドン_アクセス_トークン
グローバル フォールバック マストドン トークン (プレフィックス付きの var が見つからない場合に使用されます)
MASTODON_INSTANCE_URL
フォールバック —identity.yaml で mastodon_instance_url を設定することを推奨します。
ブルースカイハンドル
フォールバック — アイデンティティ.yaml で bluesky_handle を設定することを推奨します。
ブルースカイ_アプリ_パスワード
グローバル フォールバック Bluesky アプリのパスワード
日々のワークフロー
main_batch.py は 8 つの投稿と画像を生成し、承認のためにそれぞれを Telegram に送信します (cron 経由で実行)。
電話で承認/再生成/キャンセルを行います。
main_dispatcher.py は、承認された投稿をスケジュールされた時間に公開します (永続ループ)。
main_reply_listener.py は、公開された投稿をポーリングし、受信したコメントへの返信を下書きし、承認のために Telegram 経由で送り返します (永続ループ)。
3 つの永続ループ (dispatcher、telegram_listener、reply_listener) は systemd またはvisord の下に属します。
何も書かずにプレビューする
python main_batch.py --dry-run
実際のストーリーを取得し、ローカル サービス経由でテキストと画像を生成し、すべてを標準出力に出力します。 DB 書き込み、テレグラム、副作用はありません。
python main_dispatcher.py --dry-run
ソーシャル API 呼び出しを行わずに、承認された投稿ごとに公開される内容をログに記録します。
スクリプト
目的
main_batch.py
1 日 1 回実行 — ストーリーを取得し、投稿を生成し、メモリとキューに保存し、Telegram に通知します
main_dispatcher.py
永続ループ — 承認された投稿をスケジュールされた UTC 時間に公開します
main_telegram_listener.py
永続的

ループ — APPROVE / REGEN / CANCEL / REPLY_APPROVE / REPLY_CANCEL コールバックを処理します
main_reply_listener.py
永続ループ — 公開された投稿に新しいコメントをポーリングし、返信の下書きを作成し、Telegram の承認を得るために送信します。
スロットの分配
毎日のバッチごとに、デフォルトで 8 つの投稿が生成されます ( profiles/alexa/identity.yaml )。
mkdir -p プロファイル/マルコ/プロンプト プロファイル/マルコ/生成
cp プロファイル/alexa/identity.yaml プロファイル/marco/
cp プロファイル/alexa/prompts/ * .j2 プロファイル/marco/prompts/
# profiles/marco/identity.yaml と .j2 テンプレートを編集します
# 新しいプロファイルで実行する
ACTIVE_PROFILE=マルコ Python main_batch.py --dry-run
ディレクトリ名 (スラッグ) は ACTIVE_PROFILE と一致する必要があります。これは、identity.yaml の名前フィールドとは別のものです (「alexa」と「AlexaPavlova」)。
構成/
schemas.py # RawStory、Post、ProfileConfig (Pydantic v2)
settings.py # .env からの Pydantic 設定
コア/
Brain.py # LMStudio 呼び出し、テキスト クリーニング、切り捨て
curator.py # URL + タイトル + サブレディットによる重複除去;禁止されたトピックのフィルター
Factory.py # アダプター → キュレーター → 脳 → 画像を調整します
scheduler.py # ランダムなジッターを伴う UTC 対応の時間ウィンドウ
profile_loader.py # profiles/{slug}/identity.yaml をロードします
アダプター/
hn_adapter.py # ハッカーニュースのトップ記事 (TECH)
lobsters_adapter.py # Lobste.rs で最もホットな (TECH)
bearblog_adapter.py # BearBlog Discover RSS (TECH)
ask_hn_adapter.py # Algolia 経由で HN ディスカッション投稿に質問する (個人)
reddit_adapter.py # Reddit (OAuth2 認証情報が必要; デフォルトでは使用されません)
サービス/
Memory_service.py # プロファイルごとの SQLite 投稿履歴 + プラットフォーム ID
queue_service.py # 承認キュー (SQLite)
image_gen.py # SwarmUI REST クライアント
notification.py # テレグラム sendPhoto/sendMessage、承認キーボード、ロングポーリング
comment_service.py # comments.sqlite: コメント、pending_replies、poll_state
ソーシャル/
マストドン_パブリッシャー.py #

マストドン REST — パブリッシュ +publish_reply
bluesky_publisher.py # AT プロトコル XRPC — パブリッシュ +publish_reply
mastodon_comment_fetcher.py # /api/v1/statuses/{id}/context 経由で応答を取得します
bluesky_comment_fetcher.py # app.bsky.feed.getPostThread 経由で応答を取得します
プロフィール/
アレクサ/
Identity.yaml # ペルソナ設定: スロット、ソース、禁止されたトピック、イメージ モデル
プロンプト/*.j2 # Jinja2 テンプレート: システム プロンプト + ムードごと + プラットフォームごと
generated/ # 出力画像 (slot_NNN.png)
Memory.sqlite # 各プロンプトにコンテキストとして挿入された投稿履歴
queue.sqlite # 承認キュー
comments.sqlite # コメント、保留中の返信、投票状態
テスト
pytest testing/ -v # 312 個のテスト、すべてモック化 - ライブ サービスは必要ありません
すべての HTTP 呼び出し (LMStudio、SwarmUI、Telegram、HN、Algolia など) は respx でモック化されます。
マストドンとブルースカイのためのローカルファーストの AI ペルソナ — 機械によってドラフトされ、人間によって承認され、根拠のないものは公開されます。
dev.to/msalsas/building-an-ai-persona-that-doesnt-lie-the-parts-nobody-bothers-to-build-33ii
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A local-first AI persona for Mastodon and Bluesky — drafts by machine, approved by a human, nothing it can't ground gets published. - msalsas/amanuensis

GitHub - msalsas/amanuensis: A local-first AI persona for Mastodon and Bluesky — drafts by machine, approved by a human, nothing it can't ground gets published. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
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
msalsas
/
amanuensis
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits adapters adapters config config core core docs/ img docs/ img profiles/ alexa profiles/ alexa scripts scripts services services social social tests tests .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md main_batch.py main_batch.py main_dispatcher.py main_dispatcher.py main_reply_listener.py main_reply_listener.py main_telegram_listener.py main_telegram_listener.py manual_post.py manual_post.py pyproject.toml pyproject.toml View all files Repository files navigation
A local-first AI persona that writes under a human's veto. It drafts, you approve, and nothing it can't ground gets published.
A local-first pipeline for running an AI persona on Mastodon and Bluesky. The pilot persona, AlexaPavlova , posts as a sarcastic senior Berlin dev with dry takes on tech news and open source.
Everything runs on a local GPU machine. No cloud LLM calls.
See it live: Mastodon · Bluesky — now disclosed as AI, no longer posting.
Every post is reviewed on a phone before it publishes — approve, regenerate text or image, or cancel.
Status: this was an experiment, not an active product. The code is MIT-licensed and works end-to-end — fork it, learn from it, run your own persona. Issues and PRs may not get a response.
The hard part wasn't generating text, it was stopping the model from fabricating technical detail. The short version: factual-only source summaries, deterministic cleanup before any LLM judgment, a regex pre-screen in front of an LLM grounding check, titles-only memory, and a human approving every post over Telegram.
Full write-up of the design and what broke along the way: write-up .
pip install -e " .[dev] "
cp .env.example .env
2. Start local services
Download LMStudio and load any instruction-tuned model (tested with Mistral-7B-Instruct and similar)
Go to Local Server → start the server on port 1234
Set LMSTUDIO_BASE_URL=http://localhost:1234 in .env
Install SwarmUI and load the image model + 41ex4_p4v10v4 LoRA
Set SWARMUI_BASE_URL=http://localhost:7801 in .env
Message @BotFather → /newbot → copy the token into TELEGRAM_BOT_TOKEN
Message @userinfobot → copy your numeric ID into TELEGRAM_CHAT_ID
Send any message to your new bot so it can message you back
python main_batch.py --dry-run
This fetches real stories and generates posts + images using your local services. Nothing is written to any database and nothing is sent to Telegram. If this prints 8 posts, your local stack is working.
Add your social credentials to .env (Mastodon and/or Bluesky — both optional, see table below), then open three terminals :
# Terminal 1 — generate today's posts and send to Telegram for approval
python main_batch.py
# Terminal 2 — listen for your Telegram approvals
python main_telegram_listener.py
# Terminal 3 — publish approved posts at their scheduled time
python main_dispatcher.py
Approve posts in Telegram. The dispatcher picks them up and publishes. Done.
For long-running setups, run the three persistent processes under systemd or supervisord so they survive reboots. main_batch.py is a one-shot script — run it via cron or manually each day.
The persona's images come from a custom LoRA trained on top of Juggernaut XL "Ragnarok" (SDXL), generated through SwarmUI. The LoRA is not included in this repo — it only contains trained deltas, not the base model.
LoRA: download from Hugging Face — msalsas/alexa-lora . Trigger word 41ex4_p4v10v4 , weight 0.3 , generated at 768×1024.
Base model: Juggernaut XL "Ragnarok" by RunDiffusion — get it separately, not distributed here.
The LoRA was trained on a fully synthetic dataset (images generated with Juggernaut XL); the character is not based on any real person.
To run the alexa profile with images you need both: load Juggernaut XL in SwarmUI and apply this LoRA. See the model card on Hugging Face for the exact prompt format.
Adapters (HN, Lobsters, BearBlog, AskHN)
└── Curator (dedup by URL + title + subreddit, banned-topic filter)
└── BatchFactory
├── Brain (LMStudio → post text + image prompt)
└── ImageService (SwarmUI → PNG via LoRA)
└── Scheduler (UTC time windows with jitter)
└── QueueService (SQLite)
└── TelegramNotifier (photo + approval keyboard)
├── MastodonPublisher (on APPROVE)
└── BlueskyPublisher (on APPROVE)
Reply pipeline (runs in parallel):
ReplyListener (random poll, 30 min – 2 h per post)
├── MastodonCommentFetcher / BlueskyCommentFetcher
├── Brain.evaluate_relevance() → skip or draft reply
├── Brain.generate_reply()
└── TelegramNotifier (REPLY_APPROVE / REPLY_CANCEL)
├── MastodonPublisher.publish_reply() (on APPROVE)
└── BlueskyPublisher.publish_reply() (on APPROVE)
Requirements
LMStudio running locally (OpenAI-compatible API)
SwarmUI running locally with the 41ex4_p4v10v4 LoRA loaded
A Telegram bot token + chat ID (for the approval workflow)
Mastodon and/or Bluesky credentials (for publishing)
An OpenWeatherMap API key (free tier, for ambient context in prompts)
pip install -e " .[dev] "
cp .env.example .env
# Fill in .env with your tokens and service URLs
.env reference
Variable
Description
LMSTUDIO_BASE_URL
LMStudio API base, e.g. http://localhost:1234
SWARMUI_BASE_URL
SwarmUI base, e.g. http://localhost:7801
TELEGRAM_BOT_TOKEN
Bot token from @BotFather
TELEGRAM_CHAT_ID
Your personal chat ID (use @userinfobot to find it)
ACTIVE_PROFILE
Profile slug, default alexa
WEATHER_API_KEY
OpenWeatherMap key (free)
ALEXA_MASTODON_ACCESS_TOKEN
Mastodon token for the alexa profile
ALEXA_BLUESKY_APP_PASSWORD
Bluesky app password for the alexa profile
MASTODON_ACCESS_TOKEN
Global fallback Mastodon token (used if no prefixed var found)
MASTODON_INSTANCE_URL
Fallback — prefer setting mastodon_instance_url in identity.yaml
BLUESKY_HANDLE
Fallback — prefer setting bluesky_handle in identity.yaml
BLUESKY_APP_PASSWORD
Global fallback Bluesky app password
Daily workflow
main_batch.py generates 8 posts + images and sends each to Telegram for approval (run via cron).
You APPROVE / REGEN / CANCEL on your phone.
main_dispatcher.py publishes approved posts at their scheduled time (persistent loop).
main_reply_listener.py polls published posts, drafts replies to incoming comments, and sends them back through Telegram for approval (persistent loop).
The three persistent loops ( dispatcher , telegram_listener , reply_listener ) belong under systemd or supervisord.
Preview without writing anything
python main_batch.py --dry-run
Fetches real stories, generates text and images via local services, prints everything to stdout. No DB writes, no Telegram, no side effects.
python main_dispatcher.py --dry-run
Logs what would be published for each approved post without making any social API calls.
Script
Purpose
main_batch.py
Run once daily — fetches stories, generates posts, saves to memory + queue, notifies Telegram
main_dispatcher.py
Persistent loop — publishes APPROVED posts at their scheduled UTC time
main_telegram_listener.py
Persistent loop — handles APPROVE / REGEN / CANCEL / REPLY_APPROVE / REPLY_CANCEL callbacks
main_reply_listener.py
Persistent loop — polls published posts for new comments, drafts replies, sends for Telegram approval
Slot distribution
Each daily batch generates 8 posts by default ( profiles/alexa/identity.yaml ):
mkdir -p profiles/marco/prompts profiles/marco/generated
cp profiles/alexa/identity.yaml profiles/marco/
cp profiles/alexa/prompts/ * .j2 profiles/marco/prompts/
# Edit profiles/marco/identity.yaml and the .j2 templates
# Run with the new profile
ACTIVE_PROFILE=marco python main_batch.py --dry-run
The directory name (slug) must match ACTIVE_PROFILE . It is separate from the name field in identity.yaml ( "alexa" vs "AlexaPavlova" ).
config/
schemas.py # RawStory, Post, ProfileConfig (Pydantic v2)
settings.py # Pydantic-settings from .env
core/
brain.py # LMStudio calls, text cleaning, truncation
curator.py # Dedup by URL + title + subreddit; banned-topic filter
factory.py # Orchestrates adapters → curator → brain → image
scheduler.py # UTC-aware time windows with random jitter
profile_loader.py # Loads profiles/{slug}/identity.yaml
adapters/
hn_adapter.py # Hacker News top stories (TECH)
lobsters_adapter.py # Lobste.rs hottest (TECH)
bearblog_adapter.py # BearBlog Discover RSS (TECH)
ask_hn_adapter.py # Ask HN discussion posts via Algolia (PERSONAL)
reddit_adapter.py # Reddit (requires OAuth2 credentials; not used by default)
services/
memory_service.py # Per-profile SQLite post history + platform IDs
queue_service.py # Approval queue (SQLite)
image_gen.py # SwarmUI REST client
notification.py # Telegram sendPhoto/sendMessage, approval keyboards, long-poll
comment_service.py # comments.sqlite: comments, pending_replies, poll_state
social/
mastodon_publisher.py # Mastodon REST — publish + publish_reply
bluesky_publisher.py # AT Protocol XRPC — publish + publish_reply
mastodon_comment_fetcher.py # Fetches replies via /api/v1/statuses/{id}/context
bluesky_comment_fetcher.py # Fetches replies via app.bsky.feed.getPostThread
profiles/
alexa/
identity.yaml # Persona config: slots, sources, banned topics, image model
prompts/*.j2 # Jinja2 templates: system prompt + per-mood + per-platform
generated/ # Output images (slot_NNN.png)
memory.sqlite # Post history injected as context into each prompt
queue.sqlite # Approval queue
comments.sqlite # Comments, pending replies, poll state
Tests
pytest tests/ -v # 312 tests, all mocked — no live services required
All HTTP calls (LMStudio, SwarmUI, Telegram, HN, Algolia, etc.) are mocked with respx .
A local-first AI persona for Mastodon and Bluesky — drafts by machine, approved by a human, nothing it can't ground gets published.
dev.to/msalsas/building-an-ai-persona-that-doesnt-lie-the-parts-nobody-bothers-to-build-33ii
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
