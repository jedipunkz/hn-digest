---
source: "https://github.com/perminder-klair/subwave"
hn_url: "https://news.ycombinator.com/item?id=48674071"
title: "Personal internet radio: Agentic AI DJ"
article_title: "GitHub - perminder-klair/subwave: Personal internet radio: Agentic AI DJ · GitHub"
author: "fidotron"
captured_at: "2026-06-25T14:57:42Z"
capture_tool: "hn-digest"
hn_id: 48674071
score: 1
comments: 0
posted_at: "2026-06-25T14:34:54Z"
tags:
  - hacker-news
  - translated
---

# Personal internet radio: Agentic AI DJ

- HN: [48674071](https://news.ycombinator.com/item?id=48674071)
- Source: [github.com](https://github.com/perminder-klair/subwave)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T14:34:54Z

## Translation

タイトル：パーソナルインターネットラジオ：Agentic AI DJ
記事タイトル: GitHub - perminder-klair/subwave: パーソナル インターネット ラジオ: Agentic AI DJ · GitHub
説明: パーソナル インターネット ラジオ: Agentic AI DJ。 GitHub でアカウントを作成して、perminder-klair/subwave の開発に貢献してください。

記事本文:
GitHub - perminder-klair/subwave: パーソナル インターネット ラジオ: Agentic AI DJ · GitHub
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
パーミンダー・クレア
/
サブウェーブ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ブランチタグの開発

ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
930 コミット 930 コミット .claude/ スキル .claude/ スキル .github .github アプリ アプリ bin bin cli cli コントローラー コントローラー docker docker docs docs infra/ cli-installer infra/ cli-installer 液体石鹸 液体石鹸 mcp-subwave mcp-subwave スクリプト スクリプト サウンド サウンド テンプレート テンプレート ツール/ jamendo ツール/ jamendo web web .agents .agents .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .prettierrc.json .prettierrc.json .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md DEPLOY.md DEPLOY.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md ca_profile.xml ca_profile.xml docker-compose.byo.yml docker-compose.byo.yml docker-compose.dev.yml docker-compose.dev.yml docker-compose.tts-heavy-gpu.yml docker-compose.tts-heavy-gpu.yml docker-compose.yml docker-compose.yml install.sh install.sh package-lock.json package-lock.json package.json package.json release-please-config.json release-please-config.json すべて表示ファイル リポジトリ ファイルのナビゲーション
個人向けインターネットラジオ局。 1 つの Icecast ストリーム、1 つのブロードキャスト。
すべてのリスナーは同時に同じことを聞きます。 AI DJ が選ぶ
トラックとそれらの間の会話: 駅の ID、時間チェック、天気、
次にリリースされるものについての簡単な紹介。簡単に音楽をリクエストできます
言語。 DJ はあなたの意図を理解して、それを取り入れます。
プレイリストではなくラジオです。リスナーごとのシャッフル、スキップ ボタン、なし
「次はあなたの番です。」チャンネルを合わせると、流れているものは何でも聞こえます。
サブウェーブ.ショーリール.1080p.mp4
ライブデモ
デモプレーヤー — getsubwave.com/listen
モバイル アプリ — ネイティブ プレーヤー — App Store の iOS、Google Play の Android
セットアップのウォークスルー

— getsubwave.com/setup
オペレーターマニュアル — getsubwave.com/manual
リスナープレイヤー。共有ブロードキャストが 1 つあり、アプリ内で曲のリクエストが可能です。
管理コンソール。オペレーターがステーションを運営する場所。
図書館天文台。 DJ がタグ付けしたすべてのトラックのデータ アート マップ。ジャンルごとに配置され、エネルギーによって照らされます。ポイントをクリックすると、BPM、キー、ムード、エンベディング、最近傍などの完全なドシエが表示されます。 /observatory で開きます。
1 つの共有 Icecast ストリーム。すべてのリスナーは同じ放送を同時に聞きます。
選んでしゃべるAI DJ。追跡、イントロの作成、ステーション ID、時刻、天気の読み取りを管理します。
わかりやすい言葉でのリクエスト。 「もっと明るい曲をかけて」または「レディオヘッドの曲なら何でも」でも大丈夫です。
あなただけの音楽ライブラリ。 Subsonic API を介して Navidrome から取得します。外部カタログはありません。
交換可能な LLM プロバイダー。 Ollama、Anthropic、OpenAI、Google、DeepSeek、OpenRouter、Vercel AI Gateway、または任意の OpenAI 互換サーバー。再デプロイせずに管理 UI から変更します。
5基のTTSエンジン。 Piper と Kokoro は高速ローカルスピーチ用にインプロセスであり、オプションの tts-heavy サイドカー ( docker compose --profile tts-heavy up -d ) により、Chatterbox (ゼロショット音声クローン作成) と PocketTTS (6 倍のリアルタイム、EN/FR/DE/IT/ES/PT) が追加されます。クラウド(OpenAI/イレブンラボ)もご利用いただけます。音声の種類ごとに異なるエンジンを選択します。
複数の DJ ペルソナ。最大 10 人の魂がローテーションし、それぞれが独自の声と文体を持ちます。
デュアルコーデックブロードキャスト。 MP3 128 kbps (Sonos、ハードウェア ラジオ、自動車用)。最新のブラウザでは Ogg-Opus 96 kbps。 Web プレーヤーが自動的に選択します。
ネイティブアプリとPWA。ネイティブ iOS (App Store 上) および Android (Google Play 上) プレーヤー - バックグラウンド オーディオ、ロック画面 / CarPlay / Android Auto コントロール、マルチステーション - 加えて、電話とデスクトップにインストール可能な PWA。
予定されているショー。 24×7 のグリッド。

各スロットには独自のペルソナ、雰囲気、スキルがあります。
プラグイン可能なスキル。 DJ のトラック間のセグメント (天気、ニュース、交通情報、そして自分自身の情報) はスキルです。ビルトインは初回起動時に state/skills/<kind>/ の下に編集可能なファイルとしてスキャフォールディングされるため、管理コンソールから直接概要を書き直したり、ニュース フィード (BBC → 独自の RSS) を変更したりすることができます。コードや再デプロイは必要ありません。 SKILL.md (およびオプションのデータ取得コード) を state/skills/ にドロップし、Rescan をクリックして有効にすることで、独自のものを追加します。 docs/custom-skills.md を参照してください。
気分を意識したローテーション。時間帯、天気、フェスティバルの日によって、何が演奏されるか、DJ がどのように話すかが偏ります。
時間ごとのアーカイブ。 1 時間ごとに MP3 として保存され、後で再生できます。
クロスフェード+音声ダッキング。トラックはスムーズにブレンドされます。音楽はDJのスピーチの下で沈み、そして再び立ち上がります。
管理コンソール。ライブ ステータス、キュー、ブース ログ、ペルソナ、ショー、スキル、統計、および最近の LLM コールのデバッグ ビュー。
図書館天文台。 /observatory にあるタグ付けされたすべてのトラックのフルスクリーンのデータアート マップ。ジャンルごとに配置され、エネルギーで照らされ、トラックごとの完全なドシエ (BPM、キー、ムード、埋め込まれた指紋、ベクトル空間内の最近傍) が表示されます。数百から数万のトラックまで拡張できます。
MCPサーバー。外部エージェント (Claude Desktop、Cursor など) は曲をリクエストし、DJ を駆動できます。
自己ホスト型。 1 つの Docker が 1 つの Linux ホスト上で -d を構成します。 TLS 用に前面にあるオプションの Cloudflare。
プレイリストはあなたが管理するリストです。ラジオはあなたが参加する放送です。サブ/ウェーブ
2 番目の種類です:
1 つの共有ストリーム。全員が接続する単一の Icecast マウント。
誰もが同じ瞬間に同じ音声を聞きます。それがそれを可能にするのです
ジュークボックスの代わりにステーション。
スキップはありません。トラックエンドは唯一の自然な移行です。 DJ — 人間が厳選した
ペルソナと LLM — リスナーではなくペーシングを所有します。 （オペラ

トルズ缶
管理 API 経由でスキップします。リスナーはできません。）
カタログではなく、DJとしてのAI。音楽はあなた自身のライブラリであり、提供されます
Subsonic API 経由の Navidrome による。 LLM は次のことを選択して話し合います
トラックの間。それは音楽を生成しませんし、あなたの好みに取って代わることもありません。
自己ホスト型で交換可能。 Cloudflareの背後にある1つのLinuxボックス上で実行されます。の
LLM プロバイダーは実行時に交換可能です (Ollama、Anthropic、OpenAI、Google、
OpenRouter、Vercel AI Gateway) を再デプロイする必要はありません。
クイック スタート (CLI - 推奨)
カール -fsSL https://cli.gestbwave.com | sh # インストールし、init + start を提案します
サブウェーブセットアップ # Navidrome + LLM を接続
インストーラー中に 2 つの Enter プロンプト (「Run subwave init now?」) が表示されます。
今すぐスタックを上げますか? ) とスタックがオンエアされます。サブウェーブのセットアップ
Navidrome と LLM を接続するか、ブラウザで同じことを行います。
http://localhost:7700/onboarding 。
ホスト上にクローンもノードもありません。サブウェーブステータス/ログ/ドクター/アップデート
その後はどこからでも作業できます。
クイック スタート (CLI なし、生の Docker)
ホスト上のバイナリをスキップして docker compose にこだわりたい場合は、次のようにします。
mkdir サブウェーブ && cd サブウェーブ
カール -O https://raw.githubusercontent.com/perminder-klair/subwave/main/docker-compose.yml
カール -O https://raw.githubusercontent.com/perminder-klair/subwave/main/.env.example
mv .env.example .env
# .env を編集: ADMIN_USER、ADMIN_PASS、SITE_URL (3 つの変数、それだけです) を設定します。
ドッカー構成 -d
# 次に、https://your-host/onboarding を開きます。 Web ウィザードは Navidrome を収集します。
# LLM、TTS、DJ ペルソナ、およびジングルのレンダリングを提供します。
機能的に同一: 同じイメージ、同じ状態レイアウト、同じ永続性。
CLI は、curl-and-edit ダンスを保存し、サブウェーブ ログを提供するだけです。
残りのライフサイクルでは、 subwave Doctor など。
Chatterbox (ゼロショット音声クローン作成) と PocketTTS (高速多言語) は、

約 5 ～ 6 GB の PyTorch を追加する別の subwave-tts-heavy サイドカー
デフォルトでは開始されません。有効にするには:
docker compose --profile tts-heavy up -d
コントローラーはサイドカーを自動的に検出するように接続されています。やめてください
もう一度 docker compose --profile tts-heavy stop tts-heavy を使用します。残りの部分
スタックは実行を続け、Chatterbox/PocketTTS ペルソナは静かにフォールバックします。
パイパーに。古い docker build --build-arg WITH_CHATTERBOX=1 パスはまだ残っています
カスタムビルドされたコントローラーイメージがすでにある場合に機能します — を参照してください。
docker/Dockerfile.controller 。
git clone https://github.com/perminder-klair/subwave.git && cd subwave
./scripts/setup.sh # 3 変数の root .env + state/ を足場にします
docker compose -f docker-compose.dev.yml up -d # ブロードキャスト (icecast2 + liquidsoap) + コントローラー
cd web && npm install && npm run dev # Web UI on :7700、個別およびホットリロード
# 次に http://localhost:7700/onboarding で構成を完了します。
Dev compose は、コントローラー/src/、radio.liq、sounds/ をバインドマウントします。
レポ。コントローラーは tsx watch の下で実行されるため、src/** は内部でホットリロードを編集します
コンテナ。 radio.liq の編集には、 docker compose -f docker-compose.dev.yml restart Broadcasting だけが必要です。
スタンドアロンの subwave CLI は、複製されたリポジトリ内でも動作します。 cd subwave && subwave start dev は正しいことを行います。コントリビューターの利便性は npm start です。これは tsx -CLI ソースを直接実行するため、未リリースの変更は
運動した。同じコマンド、同じフラグ、npm install -g は必要ありません。
同じ CLI が、ステーションを実行するためのコンソールとしても機能します。 npm startを実行する
ステータス認識メニューの場合。すべてのメニュー操作はワンショットのサブコマンドでもあります。
npm start の後に追加 -- :
npm start # 対話型オペレーターコンソール (ステータス認識メニュー)
npm start -- setup # 初回起動ウィザード: Navidrome、LLM、admin、env ファイル
npm start -- ステータス # コンポジット

環境、サービス、再生中、最近のイベント
npm start -- ドクター # 完全な診断スイープ
npm start -- start dev # docker compose up -d (dev または prod)
npm start -- ブロードキャストを再起動 # プレーン再起動 (radio.liq は開発にバインドマウントされています)
npm start -- コントローラーを再起動 # リビルド + 再作成 (ソースはビルド時に COPY-d です)
npm start -- ログコントローラー # 末尾 1 サービス
npm start -- listen # ブラウザで Web プレーヤーを開きます
npm start --admin # ブラウザで管理コンソールを開きます
npm start -- stop # docker compose down (最初に確認)
実稼働デプロイ
単一の Linux ホスト、Cloudflare 終端 TLS、4 つの内部への Caddy ルーティング
サービス。上記の CLI を使用しないクイックスタートは次のとおりです。
正規パス: 2 つのファイルをカールし、3 つの変数を入力し、 docker compose up -d を実行し、ブラウザでのセットアップを完了します。ホストについては DEPLOY.md を参照してください。
前提条件、Cloudflareのセットアップ、アップデート、バックアップ。
アンレイドでは？コミュニティ アプリケーションからのワンクリック インストール
([アプリ] タブで SUB/WAVE を検索) — または、完全な分割コンテナ スタックを実行します。
Compose Manager Plus プラグイン。両方とも docs/unraid.md にあります。
独自のリバース プロキシを持参してください。すでに Traefik、nginx、または
ホームラボで独自の Caddy を使用する場合は、バンドルされた Caddy compose を BYO バリアントに交換します。
docker compose -f docker-compose.byo.yml アップ -d
あの博覧会

[切り捨てられた]

## Original Extract

Personal internet radio: Agentic AI DJ. Contribute to perminder-klair/subwave development by creating an account on GitHub.

GitHub - perminder-klair/subwave: Personal internet radio: Agentic AI DJ · GitHub
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
perminder-klair
/
subwave
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
develop Branches Tags Go to file Code Open more actions menu Folders and files
930 Commits 930 Commits .claude/ skills .claude/ skills .github .github app app bin bin cli cli controller controller docker docker docs docs infra/ cli-installer infra/ cli-installer liquidsoap liquidsoap mcp-subwave mcp-subwave scripts scripts sounds sounds templates templates tools/ jamendo tools/ jamendo web web .agents .agents .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .prettierrc.json .prettierrc.json .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md DEPLOY.md DEPLOY.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md ca_profile.xml ca_profile.xml docker-compose.byo.yml docker-compose.byo.yml docker-compose.dev.yml docker-compose.dev.yml docker-compose.tts-heavy-gpu.yml docker-compose.tts-heavy-gpu.yml docker-compose.yml docker-compose.yml install.sh install.sh package-lock.json package-lock.json package.json package.json release-please-config.json release-please-config.json View all files Repository files navigation
A personal internet radio station. One Icecast stream, one broadcast.
Every listener hears the same thing at the same time. An AI DJ picks the
tracks and talks between them: station idents, time checks, the weather,
a quick intro for whatever's going out next. You can ask for music in plain
language; the DJ works out what you meant and slots it in.
It's radio , not a playlist. No per-listener shuffle, no skip button, no
"up next for you." You tune in and hear whatever is on.
SUBWAVE.Showreel.1080p.mp4
Live demo
Demo player — getsubwave.com/listen
Mobile apps — native players — iOS on the App Store , Android on Google Play
Setup walkthrough — getsubwave.com/setup
Operator manual — getsubwave.com/manual
The listener player. One shared broadcast, with in-app song requests.
The admin console. Where the operator runs the station.
The Library Observatory. A data-art map of every track the DJ has tagged, placed by genre and lit by energy. Click a point for its full dossier: BPM, key, mood, embeddings, and nearest neighbours. Open it at /observatory .
One shared Icecast stream. Every listener hears the same broadcast at the same time.
AI DJ that picks and talks. Curates tracks, writes intros, and reads station idents, the time, and the weather.
Plain-language requests. "Play something more upbeat" or "anything by Radiohead" works.
Your own music library. Pulls from Navidrome over the Subsonic API. No external catalogue.
Swappable LLM provider. Ollama, Anthropic, OpenAI, Google, DeepSeek, OpenRouter, Vercel AI Gateway, or any OpenAI-compatible server. Change it from the admin UI with no redeploy.
Five TTS engines. Piper and Kokoro in-process for fast local speech, plus an optional tts-heavy sidecar ( docker compose --profile tts-heavy up -d ) that adds Chatterbox (zero-shot voice cloning) and PocketTTS (6× real-time, EN/FR/DE/IT/ES/PT). Cloud (OpenAI / ElevenLabs) is also available. Pick a different engine per kind of speech.
Multiple DJ personas. Up to 10 souls in rotation, each with its own voice and writing style.
Dual-codec broadcast. MP3 128 kbps for Sonos, hardware radios, and cars; Ogg-Opus 96 kbps for modern browsers. The web player picks automatically.
Native apps and PWA. Native iOS (on the App Store) and Android (on Google Play) players — background audio, lock-screen / CarPlay / Android Auto controls, multi-station — plus an installable PWA on phone and desktop.
Scheduled shows. A 24×7 grid; each slot has its own persona, mood, and skills.
Pluggable skills. The DJ's between-track segments — weather, news, traffic, and your own — are skills. The built-ins are scaffolded as editable files under state/skills/<kind>/ on first boot, so you can rewrite a brief or change the news feed (BBC → your own RSS) right from the admin console — no code, no redeploy. Add your own by dropping a SKILL.md (plus optional data-fetching code) into state/skills/ , hitting Rescan, and enabling it. See docs/custom-skills.md .
Mood-aware rotation. Time of day, weather, and festival days bias what gets played and how the DJ talks.
Hourly archives. Every hour saved as MP3 for later replay.
Crossfade + voice ducking. Tracks blend smoothly; the music ducks under DJ speech and lifts back up.
Admin console. Live status, queue, booth log, personas, shows, skills, stats, and a debug view of recent LLM calls.
Library Observatory. A full-screen, data-art map of every tagged track at /observatory — placed by genre, lit by energy, with a full dossier per track (BPM, key, mood, embedding fingerprints, and nearest-in-vector-space neighbours). Scales from a few hundred to tens of thousands of tracks.
MCP server. External agents (Claude Desktop, Cursor, etc.) can request songs and drive the DJ.
Self-hosted. One docker compose up -d on a single Linux host. Optional Cloudflare in front for TLS.
A playlist is a list you control. Radio is a broadcast you join. SUB/WAVE
is the second kind:
One shared stream. A single Icecast mount everyone connects to.
Everyone hears the same audio at the same instant. That's what makes it
a station instead of a jukebox.
No skip. Track-end is the only natural transition. The DJ — human-curated
personas plus an LLM — owns the pacing, not the listener. (Operators can
skip via the admin API; listeners cannot.)
AI as the DJ, not the catalogue. The music is your own library, served
by Navidrome over the Subsonic API. The LLM picks what's next and talks
between tracks. It doesn't generate music and it doesn't replace your taste.
Self-hosted and swappable. Runs on one Linux box behind Cloudflare. The
LLM provider is swappable at runtime (Ollama, Anthropic, OpenAI, Google,
OpenRouter, Vercel AI Gateway) with no redeploy.
Quick start (CLI — recommended)
curl -fsSL https://cli.getsubwave.com | sh # installs, then offers to init + start
subwave setup # connect Navidrome + LLM
Two Enter prompts during the installer ( Run subwave init now? , then
Bring the stack up now? ) and the stack is on-air. subwave setup
connects Navidrome and your LLM, or do the same in the browser at
http://localhost:7700/onboarding .
No clone, no Node on the host. subwave status / logs / doctor / update
work from anywhere afterwards.
Quick start (no CLI, raw docker)
If you'd rather skip our binary on your host and stick to docker compose :
mkdir subwave && cd subwave
curl -O https://raw.githubusercontent.com/perminder-klair/subwave/main/docker-compose.yml
curl -O https://raw.githubusercontent.com/perminder-klair/subwave/main/.env.example
mv .env.example .env
# Edit .env: set ADMIN_USER, ADMIN_PASS, SITE_URL (three vars, that's it).
docker compose up -d
# Then open https://your-host/onboarding. The web wizard collects Navidrome,
# LLM, TTS, DJ persona, and offers to render jingles.
Functionally identical: same images, same state layout, same persistence.
The CLI just saves you the curl-and-edit dance and gives you subwave logs ,
subwave doctor , etc. for the rest of the lifecycle.
Chatterbox (zero-shot voice cloning) and PocketTTS (fast multilingual) live in
a separate subwave-tts-heavy sidecar that adds ~5–6 GB of PyTorch and is
not started by default. To enable:
docker compose --profile tts-heavy up -d
The controller is wired up to discover the sidecar automatically. Stop it
again with docker compose --profile tts-heavy stop tts-heavy ; the rest of
the stack keeps running and Chatterbox/PocketTTS personas silently fall back
to Piper. The old docker build --build-arg WITH_CHATTERBOX=1 path still
works if you already have a custom-built controller image — see
docker/Dockerfile.controller .
git clone https://github.com/perminder-klair/subwave.git && cd subwave
./scripts/setup.sh # scaffolds a 3-var root .env + state/
docker compose -f docker-compose.dev.yml up -d # Broadcast (icecast2 + liquidsoap) + Controller
cd web && npm install && npm run dev # web UI on :7700, separate and hot-reloading
# Then http://localhost:7700/onboarding to finish configuration.
Dev compose bind-mounts controller/src/ , radio.liq , and sounds/ from the
repo. Controller runs under tsx watch so src/** edits hot-reload inside
the container; radio.liq edits just need a docker compose -f docker-compose.dev.yml restart broadcast .
The standalone subwave CLI works inside the cloned repo too. cd subwave && subwave start dev does the right thing. The contributor convenience is npm start , which tsx -runs the CLI source directly so unreleased changes are
exercised. Same commands, same flags, no npm install -g needed.
The same CLI doubles as the console for running the station. Run npm start
for a status-aware menu; every menu action is also a one-shot subcommand,
appended after npm start -- :
npm start # interactive operator console (status-aware menu)
npm start -- setup # first-boot wizard: Navidrome, LLM, admin, env files
npm start -- status # compose env, services, now-playing, recent events
npm start -- doctor # full diagnostic sweep
npm start -- start dev # docker compose up -d (dev or prod)
npm start -- restart broadcast # plain restart (radio.liq is bind-mounted in dev)
npm start -- restart controller # rebuild + recreate (source is COPY-d at build)
npm start -- logs controller # tail one service
npm start -- listen # open the web player in a browser
npm start -- admin # open the admin console in a browser
npm start -- stop # docker compose down (confirms first)
Production deploy
Single Linux host, Cloudflare terminating TLS, Caddy routing to four internal
services. The no-CLI quickstart above is
the canonical path: curl two files, fill in three vars, docker compose up -d , finish setup in the browser. See DEPLOY.md for host
prerequisites, Cloudflare setup, updates, and backup.
On Unraid? One-click install from Community Applications
(search SUB/WAVE in the Apps tab) — or run the full split-container stack via
the Compose Manager Plus plugin. Both in docs/unraid.md .
Bring your own reverse proxy. If you already run Traefik, nginx, or your
own Caddy in your homelab, swap the bundled-Caddy compose for the BYO variant:
docker compose -f docker-compose.byo.yml up -d
That expo

[truncated]
