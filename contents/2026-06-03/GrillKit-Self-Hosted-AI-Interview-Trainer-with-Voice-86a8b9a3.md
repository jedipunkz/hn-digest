---
source: "https://github.com/grillkit/grillkit"
hn_url: "https://news.ycombinator.com/item?id=48370292"
title: "GrillKit – Self-Hosted AI Interview Trainer with Voice"
article_title: "GitHub - GrillKit/grillkit: Self-hosted AI technical interview trainer with voice and real-time scoring · GitHub"
author: "vitchenkokir"
captured_at: "2026-06-03T00:45:23Z"
capture_tool: "hn-digest"
hn_id: 48370292
score: 2
comments: 0
posted_at: "2026-06-02T13:53:53Z"
tags:
  - hacker-news
  - translated
---

# GrillKit – Self-Hosted AI Interview Trainer with Voice

- HN: [48370292](https://news.ycombinator.com/item?id=48370292)
- Source: [github.com](https://github.com/grillkit/grillkit)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:53:53Z

## Translation

タイトル: GrillKit – 音声付き自己ホスト型 AI 面接トレーナー
記事タイトル: GitHub - GrillKit/grillkit: 音声とリアルタイムのスコアリングを備えた自己ホスト型 AI 技術面接トレーナー · GitHub
説明: 音声とリアルタイムのスコアリングを備えた自己ホスト型 AI 技術面接トレーナー - GrillKit/grillkit

記事本文:
GitHub - GrillKit/グリルキット: 音声とリアルタイムのスコアリングを備えた自己ホスト型 AI 技術面接トレーナー · GitHub
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
グリルキット
/
グリルキット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コ

デ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .github .github alembic alembic アプリ アセット アセット データ データ 静的 静的テンプレート テンプレート テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンソースの AI 技術面接トレーナー。厳選された YAML 質問バンクから練習し、構造化されたスコアリングとフォローアップを取得し、オプションで独自の LLM (クラウドまたはローカル) で音声を使用します。
GrillKit を使用する理由 · クイック スタート · 変更履歴 · アーキテクチャ
なぜ GrillKit を使うのか (ChatGPT だけではない)
一般的なチャット アシスタントは柔軟性がありますが、面接を実行してくれるわけではありません。
構造化された練習 — トラック、難易度、トピックを選択します。 GrillKit は質問計画を作成し、その場限りのプロンプトではなく、セッション全体のスコアを維持します。
プライバシーと制御 — ラップトップまたはサーバー上の Docker 経由で実行します。 API キーとインタビュー履歴は ./data (gitignored) の下に残ります。 LLM プロバイダー以外のアカウントやサブスクリプションは必要ありません (クラウド モデルを使用している場合)。
デモビデオ — セットアップからフィードバックの採点までの完全なフロー
ダッシュボード — 最近のセッションとクイックスタート
面接のセットアップ — 質問バンクのトラック、レベル、トピック、セッションのオプション
面接セッション — AI 採点によるリアルタイムの Q&A と最終評価
インタビュー — マルチトラック設定、セッションごとに複数のトピック、WebSocket Q&A、AI スコア 1 ～ 5、最大

質問ごとに 2 回のフォローアップ
質問バンク — Python、データベース/SQL、システム設計、Kafka、RabbitMQ、Docker、Kubernetes、可観測性、Airflow など、data/questions/{track}/ の下にあります (ジュニア / ミドル / シニア (該当する場合))
タイマー — オプションのラウンドごとの時間制限。期限切れのラウンドのスコアは 0 となり、セッションは続行されます。
音声 - 入力した回答に対するオフラインのささやきディクテーション。質問を読み上げるためのオプションの Piper TTS
音声回答 — 設定されたモデルが音声入力をサポートし、Whisper の準備ができている場合、インタビュー ページから WAV 回答を録音して送信します。
セットアップ - /config のモデル カタログ、インタビュー ロケール (AI フィードバック言語)、UI からの Whisper/Piper ダウンロード
ダッシュボード - ホームページ上の最近のインタビュー履歴
デプロイ - ポート 8000 上の Docker Compose と構成、DB、およびモデル用の ./data ボリューム
クラウドプロバイダーまたはローカルの OpenAI 互換サーバー (Ollama、vLLM など) の API キー
git clone https://github.com/yourusername/grillkit.git
CDグリルキット
docker 構成 --build
http://localhost:8000 を開きます。
オプションの質問音声 (Piper TTS、同じアプリ コンテナ):
docker compose up を実行します (または、開発用に uvicorn app.main:app を実行します)。
/config を開き、 [質問を読み上げる] を有効にして保存します。
[構成] ページで、プロンプトが表示されたら、質問の音声をダウンロードします (Hugging Face からの音声はロケールごとに最大 60 MB) を使用します。
インタビューを開始します。質問は音声で再生できます。 WAV キャッシュは data/tts-cache/v2/{locale}/ の下にあります。
ホスト上の ./data には、 SQLite、 config.json 、 llm_models.json 、 Whisper/Piper モデル、および TTS キャッシュが保持されます。質問バンク、テンプレート、静的ファイルはイメージに同梱されています。
バインドマウントされた data/ が書き込み可能でない場合 (Linux UID の不一致):
PUID= $( id -u ) PGID= $( id -g ) docker compose up --build
初回の流れ
構成 ( /config ) — 1 つ以上の OpenAI 互換モデルを Catal に追加します。

たとえば、インタビュー モデルを選択し、インタビュー ロケールを設定します。接続をテストして保存します。
新しいインタビュー ( /setup ) — 1 つ以上の質問バンク トラック (トラックごとのレベル) を有効にし、複数のトピックを選択し、合計質問数を設定します (選択したトピックごとに少なくとも 1 つ。インタビュー ロケールは構成から読み取り専用です)。
インタビュー ( /interview/{id} ) — ページの読み込み履歴。テキストによる回答と補完は WebSocket 経由で行われます。
プロバイダー構成が保存されていない場合、 /setup は /config にリダイレクトされます。
寄稿者の場合: CONTRIBUTING.md を参照してください。クイック実行:
UV 同期 --extra dev
uv run uvicorn app.main:app --reload
http://127.0.0.1:8000 での初回フローと同じです。
OpenAI 互換の HTTP API はすべて機能します。
モデルをカタログに追加 — ベース URL、モデル名、オプションの API キー。モデルがマルチモーダル オーディオ (および文字起こし用の Whisper のダウンロード) をサポートしている場合にのみオーディオ入力を受け入れます。
インタビュー モデル — カタログから選択し、接続をテストし、保存します。
Locale — AI フィードバックと音声の言語 (data/config.json に保存され、gitignored)。
ウィスパー — サイズ (小、中、大) を選択し、UI からダウンロードしてディクテーションと音声による回答を行います。
質問を読み上げます — Piper を有効にし、音声をダウンロードします (約 60 MB)。
data/config.json 、 data/llm_models.json 、または API キーはコミットしないでください。
オプションの環境変数 ( ARCHITECTURE.md 内の完全なリスト):
セッション全体の制限時間（インタビューの合計時間）
さらに多くの質問バンクとカテゴリ
インタビュー UI のコード エディター
カスタム質問バンク、PWA / スタンドアロン フロントエンド
文書
内容
アーキテクチャ.md
レイヤー、HTTP/WebSocket ルート、データ フロー、永続性、質問バンク
貢献.md
開発セットアップ、テスト、ruff/mypy/pytest、コントリビューションワークフロー
変更ログ.md
発売履歴
セキュリティ
SECURITY.md の説明に従って脆弱性を報告します。セキュリティ上の問題については公開問題を開かないでください。
Apache ライセンス 2.0 (se

また注意してください)
音声とリアルタイムのスコアリングを備えた自己ホスト型 AI 技術面接トレーナー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
3
2026.5.31
最新の
2026 年 5 月 31 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted AI technical interview trainer with voice and real-time scoring - GrillKit/grillkit

GitHub - GrillKit/grillkit: Self-hosted AI technical interview trainer with voice and real-time scoring · GitHub
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
GrillKit
/
grillkit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .github .github alembic alembic app app assets assets data data static static templates templates tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Open-source AI technical interview trainer. Practice from curated YAML question banks, get structured scoring and follow-ups, and optionally use voice — with your own LLM (cloud or local).
Why GrillKit · Quick start · Changelog · Architecture
Why GrillKit (not just ChatGPT)
A general chat assistant is flexible, but it does not run an interview for you.
Structured practice — You pick tracks, difficulty, and topics; GrillKit builds a question plan and keeps score across the whole session, not a single ad-hoc prompt.
Privacy and control — Run via Docker on your laptop or server. API keys and interview history stay under ./data (gitignored). No account or subscription required beyond your LLM provider (if you use a cloud model).
Demo video — full flow from setup to scored feedback
Dashboard — recent sessions and quick start
Interview setup — question-bank tracks, levels, topics, and session options
Interview session — real-time Q&A with AI scoring and final evaluation
Interviews — multi-track setup, several topics per session, WebSocket Q&A, AI scoring 1–5, up to 2 follow-ups per question
Question banks — Python, Database/SQL, System Design, Kafka, RabbitMQ, Docker, Kubernetes, Observability, Airflow, and more under data/questions/{track}/ (junior / middle / senior where applicable)
Timer — optional per-round time limit; expired rounds score 0 and the session moves on
Voice — offline Whisper dictation for typed answers; optional Piper TTS to read questions aloud
Audio answers — when the configured model supports audio input and Whisper is ready, record and send a WAV answer from the interview page
Setup — model catalog on /config , interview locale (AI feedback language), Whisper/Piper downloads from the UI
Dashboard — recent interview history on the home page
Deployment — Docker Compose on port 8000 with ./data volume for config, DB, and models
API key for a cloud provider, or a local OpenAI-compatible server (Ollama, vLLM, …)
git clone https://github.com/yourusername/grillkit.git
cd grillkit
docker compose up --build
Open http://localhost:8000 .
Optional question voice (Piper TTS, same app container):
Run docker compose up (or uv run uvicorn app.main:app for development).
Open /config , enable Read questions aloud , save.
On the Configuration page, use Download question voice when prompted (~60 MB per locale voice from Hugging Face).
Start an interview — questions can play aloud; WAV cache lives under data/tts-cache/v2/{locale}/ .
./data on the host holds SQLite, config.json , llm_models.json , Whisper/Piper models, and TTS cache. Question banks, templates, and static files ship in the image.
If bind-mounted data/ is not writable (Linux UID mismatch):
PUID= $( id -u ) PGID= $( id -g ) docker compose up --build
First-time flow
Configuration ( /config ) — add one or more OpenAI-compatible models to the catalog, select an interview model, set interview locale; test connection, then save.
New interview ( /setup ) — enable one or more question-bank tracks (level per track), select multiple topics, set total question count (at least one per selected topic; interview locale is read-only from config).
Interview ( /interview/{id} ) — page loads history; text answers and completion go over WebSocket.
Without saved provider config, /setup redirects to /config .
For contributors: see CONTRIBUTING.md . Quick run:
uv sync --extra dev
uv run uvicorn app.main:app --reload
Same first-time flow at http://127.0.0.1:8000 .
Any OpenAI-compatible HTTP API works:
Add model to catalog — base URL, model name, optional API key; enable Accepts audio input only if the model supports multimodal audio (and download Whisper for transcription).
Interview model — pick from the catalog, Test Connection , save.
Locale — language for AI feedback and speech (stored in data/config.json , gitignored).
Whisper — choose size ( small , medium , large ), download from the UI for dictation and audio answers.
Read questions aloud — enable Piper, download a voice (~60 MB).
Do not commit data/config.json , data/llm_models.json , or API keys.
Optional environment variables (full list in ARCHITECTURE.md ):
Session-wide time limit (total interview duration)
More question banks and categories
Code editor in the interview UI
Custom question banks, PWA / standalone frontend
Document
Contents
ARCHITECTURE.md
Layers, HTTP/WebSocket routes, data flows, persistence, question banks
CONTRIBUTING.md
Dev setup, tests, ruff/mypy/pytest, contribution workflow
CHANGELOG.md
Release history
Security
Report vulnerabilities as described in SECURITY.md . Do not open public issues for security problems.
Apache License 2.0 (see also NOTICE )
Self-hosted AI technical interview trainer with voice and real-time scoring
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
3
2026.5.31
Latest
May 31, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
