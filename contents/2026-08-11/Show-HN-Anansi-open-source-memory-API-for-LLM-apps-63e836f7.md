---
source: "https://github.com/g-33-L/anansi"
hn_url: "https://news.ycombinator.com/item?id=49256253"
title: "Show HN: Anansi – open-source memory API for LLM apps"
article_title: "GitHub - g-33-L/anansi: A self-hostable memory engine for AI agents that can distinguish what was true from what the agent knew at the time · GitHub"
author: "jsuleiman"
captured_at: "2026-08-11T11:38:22Z"
capture_tool: "hn-digest"
hn_id: 49256253
score: 1
comments: 0
posted_at: "2026-08-11T11:02:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Anansi – open-source memory API for LLM apps

- HN: [49256253](https://news.ycombinator.com/item?id=49256253)
- Source: [github.com](https://github.com/g-33-L/anansi)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T11:02:19Z

## Translation

タイトル: 表示 HN: Anansi – LLM アプリ用のオープンソース メモリ API
記事タイトル: GitHub - g-33-L/anansi: 真実とエージェントが当時知っていたことを区別できる AI エージェント用の自己ホスト可能なメモリ エンジン · GitHub
説明: AI エージェント用の自己ホスト可能なメモリ エンジンで、真実とエージェントがその時点で知っていたことを区別できる - g-33-L/anansi

記事本文:
GitHub - g-33-L/anansi: 真実とエージェントが当時知っていたことを区別できる AI エージェント用の自己ホスト可能なメモリ エンジン · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
g-33-L
/
アナンシ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .github .github apps apps docs docs e2e e2e サンプル サンプル パッケージ

パッケージ スクリプト スクリプト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSEライセンス LICENSE-EE LICENSE-EE README.md README.md SECURITY.md SECURITY.md バージョン バージョン docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js package.json package.json playwright.config.ts playwright.config.ts pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml Railway.toml Railway.toml tsconfig.json tsconfig.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用の自己ホスト可能なメモリ エンジンで、真実とエージェントがその時点で知っていたことを区別できます。
アナンシは AI システムに組織が実際にどのように機能するかを永続的に理解させます
動作とそれが時間の経過とともにどのように変化したか。
会社がすでに生み出している会話、ドキュメント、チケット、
会議の記録。アナンシはそれをエージェントが読み取れる構造化記憶に変換します
応答する前に、すべてのバージョンを保持します。したがって、単に「何があるか」を尋ねることができます。
私たちのエスカレーションプロセスは？」しかし、「3月にそれがあったと思っていました、そしてそれはいつでしたか
変わりますか？」 — 引用付きの回答が得られます。
TL;DR: MIT ライセンスのセルフホスト可能なコア。商用エンタープライズ/ホスト層を備えています。 「ライセンス」を参照してください。
POST /v1/ingest # これを覚えておいてください
GET /v1/context # 現時点で関連性があることを知っていますか?
ingest はすぐに 202 を返し、高価な作業をキューで実行するため、決して
応答パスに存在します。 context は、コンパクトな既に合成されたプロファイルを返します。
システム プロンプトに直接貼り付けることができます。自分自身をランク付けするためのチャンクの山ではありません。
自己ホスト型で MIT ライセンスを取得しており、Postgres および Redis 上で実行されます。 Wまで約5分

オーキング
インスタンス、サインアップはありません。
その引用によって実際に得られるもの — エンティティ グラフには 2 つの独立した要素が含まれています
時間軸なので、何が真実で、何がシステムが知っていたかを別々に尋ねることができます。
エンティティ グラフと時間クエリ結果 (上記の 2 軸推論) は Pro+ の機能です。セルフホスト型インストールでは、デフォルトでエンタープライズ プランが設定され (「プランの制限」を参照)、自動的に取得されます。ホスト型サービスでは有料レベルが必要です。
クイックスタート — セルフホスト型、約 5 分
弊社からのアカウントや API キーはありません。以下のすべてがマシン上で実行されます。
タイミングは目標ではなく測定されます。API イメージはソースから約 1 時間でビルドされます。
ウォーム Docker では 90 秒、埋め込みモデルのダウンロードは 274 MB です。の
5 分の数字は、ステップ 2 のオプション A または B を想定しています。オプション C は 7 GB を引き出します。
イメージであり、かなり時間がかかります。
git クローン https://github.com/g-33-L/anansi.git
CDアナンシ
ドッカー構成 -d
これにより、PostgreSQL、Redis、および API が起動します。移行は最初に自動的に実行されます
ブーツ。 API とドキュメントは http://localhost:3000 で提供されます。
2. 埋め込むものを与える
Anansi には埋め込みモデルが必要ですが、Compose は埋め込みモデルを開始しません。これをスキップしてください
ステップと取り込みは成功したように見えます。埋め込みが無効であるため、202 が返されます。
非同期 — 取得は 503 で失敗します。開く
/ステータス ;埋め込みバックエンドを明示的に報告します。
したがって、問題を推測するのではなく、見ることができます。
それを満たすための 3 つの方法を、実際にかかる時間の順に並べます。最初の 2 つだけを保持します
このクイックスタートは 5 分以内に完了します。
オラマプル nomic-embed-text
他に設定するものはありません。Compose はすでにコンテナをポイントしています。
host.docker.internal:11434 。これが最速のパスであり、次の場合に推奨されるパスです。
オラマがインストールされました。
printf ' DEPLOYMENT_MODE=ハイブリッド\nINFERENCE_LOCATION=ローカル\nEMBEDDING_LOCATIO

N=クラウド\nNOMIC_API_KEY=your_key_here\n ' >> .env
docker compose up -d api
4 行すべてが必要です。 DEPLOYMENT_MODE のデフォルトは local であり、これは禁止します
クラウド プロバイダーを完全に — ハイブリッドに切り替えることなく NOMIC_API_KEY を提供することは、
意図的な起動失敗であり、見落としではないため、コンテナは起動を拒否し、
あなたもそうですよ。ハイブリッドでは、ローカル推論とクラウド埋め込みを組み合わせることができます。
有効になったことを確認します。起動ログには、解決されたモードが記載されています。
[起動] デプロイメントモード: ハイブリッド (推論=ローカル、埋め込み=クラウド、テレメトリ=許可)
これにより、取り込んだテキストが Nomic に送信されることに注意してください。それが重要な場合は、A または C を使用してください。
ここはゆっくりとした道です。 ollama/ollama:latest は GPU を搭載しているため約 7 GB
通常の接続では、プルだけで 10 分以上かかります。右
Compose にすべてを置き、ホストには何も必要ない場合の選択 - 正しい選択ではありません
アナンシを初めて試す場合。
echo " OLLAMA_BASE_URL=http://ollama:11434 " >> .env
docker compose --profile local-ai up -d # 7 GB イメージをプルします
docker compose exec ollama ollama pull nomic-embed-text # 274 MB
OLLAMA_BASE_URL がすでに設定されている場合でも追加は安全です - Compose は最後のものを取得します
定義。 API がすでに実行されていた後に設定した場合は、次のようにして再起動します。
docker compose up -d api により、新しいアドレスが取得されます。
どちらを選択しても、取り込みと取得に必要なのは nomic-embed-text (274 MB) だけです。の
はるかに大きなチャット モデル ( llama3.1:8b 、~4.7 GB) は、静的メッセージとチャット メッセージを合成するためにのみ使用されます。
動的プロファイル — 後で必要なときに ollam pull llama3.1:8b を使用してプルします。
キーは独自のデータベース内に存在します。これはホストされているサービスには接続しません。
docker compose exec api ノード dist/scripts/seed-dev-key.js you@example.com
ans_ で始まるキーを出力します。エクスポートします:
経験値

ort ANAANSI_API_KEY=ans_...
各電子メール アドレスには独自のワークスペースがあり、メモリがそれらの間で交差することはありません。
同じ電子メールを使用してコマンドを再実行すると、同じワークスペースに別のキーが発行されます。
別のメールアドレスで実行すると、保存したものを参照できないキーが提供されます
もっと前に。データが消えたように見える場合は、キーがどの電子メールから送信されたものかを確認してください。
curl -X POST http://localhost:3000/v1/ingest \
-H " 認可: ベアラー $ANAANSI_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"userId":"user_123","content":"ユーザーは音声エージェントを構築しています。TypeScript を好みます。4 人のチームです。","sourceType":"conversation"} '
すぐに 202 を返します。埋め込みはバックグラウンドで行われます。
カール -G http://localhost:3000/v1/context \
-H " 認可: ベアラー $ANAANSI_API_KEY " \
--data-urlencode " userId=user_123 " \
--data-urlencode " q=ユーザーは何を構築していますか? "
関連性のある値が設定されて戻ります。
{ "relevant" : [ { "content" : " ユーザーは音声エージェントを構築しています。TypeScript を好みます。4 人のチームです。 " ,
「類似性」 : 0.4821 } ]、「静的」 : []、「動的」 : [] }
この応答については、次の 2 つのことを知っておく価値があります。
類似性: 0 は、埋め込みがまだランディングされていないことを意味します。埋め込みは非同期であるため、取り込み後 1 秒以内に発行されたクエリにはキーワード検索だけで回答できます。もう一度尋ねると、実際のコサインスコアが表示されます。検索はハイブリッドであるため、空の結果ではなく、どちらの方法でも答えが得られます。
静的および動的は、チャット モデルが合成に使用できるようになる (ステップ 2) まで空のままです。それは想定内であり、失敗ではありません。
セマンティック部分が本当に機能していることを証明するには、言葉を共有しない何かを質問してください。
保存したもの:
カール -G http://localhost:3000/v1/context \
-H " 認可: ベアラー $ANAANSI_API_KEY " \
--data-urlencode " userId=user_123 " \
--data-urlencode " q=コーディング LAN

彼らはゲージが好きですか？ 」
このスコアは、キーワードが重複する質問よりも高くなります ( 0.5915 )。
文字通りに一致しますが、意味的にのみ一致します。
最初に /status を確認します。Postgres、Redis、
キュー、および埋め込みバックエンドのいずれかがダウンしている場合は 503 を返します。 A503
/v1/context から、失敗した依存関係の名前と、それを直接修正する方法
レスポンスボディ。
ローカル (キーなし、アカウントなし)
オプション/外部
API、Postgres、Redis、ワーカー
✅ docker compose up -d によって開始される
埋め込み + 合成 (Ollama)
✅ 上記のオプションエアコン
Nomic ホスト型エンベディング (オプション B)、合成用の Cerebras/GitHub モデル
上記のクイックスタートのすべて
✅ 完全にローカル
Slack / Notion / Google ドキュメント / リニア コネクタ
独自の OAuth 資格情報が必要
ホスト型ハイブリッド検索のヒント、エラー報告
Sentry ( SENTRY_DSN ) はオプションであり、デフォルトではオフになっています
完全なローカル パス (Docker + Ollama、コネクタが構成されていない) では、取り込まれたコンテンツがマシンから送信されることはありません。クラウド埋め込み/LLM プロバイダーまたはコネクタを追加すると、そのサーフェスのデータがボックスから削除されます。各 DEPLOYMENT_MODE で許可される内容については、「セキュリティ」を参照してください。
Compose のデフォルトは意図的に開発専用となっており、開発を開始するには十分です。
.env を作成せずに使い捨てのローカルスタック。焼き付けられた暗号値
docker-compose.yml はこのリポジトリで公開されているため公開されており、決して使用しないでください。
彼らは地域開発の外にいます。永続的なものはコピーしてください
.env.example 、 ENCRYPTION_KEY に個別の値を生成します。
API_KEY_HMAC_SECRET 、 CSRF_SIGNING_KEY 、および QUERY_API_KEY
( openssl rand -hex それぞれ 32)、読み取ります
docs/enterprise/self-hosting.md 。
最初のインストール後は決して ENCRYPTION_KEY を変更しないでください。保存されているコネクタ トークンはすべて変更されません。
それによって暗号化されます。
local-ai profi 経由ではなくホスト上で Ollama を実行する場合

ル、作成
デフォルト ( host.docker.internal ) はすでにそれを指しています。ホスト実行用に作成された .env
pnpm dev には localhost:11434 が含まれます。これはコンテナー内にコンテナーを意味します
それ自体 - Compose はそのファイルを補間するため、API は黙ってホストに到達できません
オラマ。 503 本体は、試行したアドレスに名前を付けます。これが、これを特定する方法です。
pnpm dev を使用してソースから実行することをお勧めしますか? COTRIBUTING.md を参照してください — 注
pnpm テストにはシェルに DATABASE_URL と REDIS_URL が必要であること、およびスイート
TRUNCATE はローカル データベースなので、関心のあるものを指定しないでください。
TypeScript SDK ( package/sdk ) の使用:
anansiMemory を「anansi-memory」からインポートします。
const メモリ = 新しいアナンシメモリ ( {
apiKey: プロセス。環境 。 ANAANSI_API_KEY 、
BaseUrl : "http://localhost:3000" , // セルフホスティングの場合は必須 — 以下を参照
} ) ;
メモリを待ちます。摂取 ( {
ユーザー ID : "user_123" 、
content : 「ユーザーは音声エージェントを構築しています。TypeScript を好みます。4 人のチームです。」 、
ソースタイプ: "会話" 、
} ) ;
const ctx = メモリを待機します。 context ( { userId : "user_123" , q : "ユーザーは何を構築していますか?" } ) ;
const systemPrompt = `あなたは役に立つアシスタントです。\n\n ${ メモリ . formatForPrompt (ctx) } ` ;
セルフホスティング者: ベース URL を設定します。すべてのクライアントは、デフォルトでホストされた API を使用します。
https

[切り捨てられた]

## Original Extract

A self-hostable memory engine for AI agents that can distinguish what was true from what the agent knew at the time - g-33-L/anansi

GitHub - g-33-L/anansi: A self-hostable memory engine for AI agents that can distinguish what was true from what the agent knew at the time · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
g-33-L
/
anansi
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .github .github apps apps docs docs e2e e2e examples examples packages packages scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE-EE LICENSE-EE README.md README.md SECURITY.md SECURITY.md VERSION VERSION docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js package.json package.json playwright.config.ts playwright.config.ts pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml railway.toml railway.toml tsconfig.json tsconfig.json turbo.json turbo.json View all files Repository files navigation
A self-hostable memory engine for AI agents that can distinguish what was true from what the agent knew at the time.
Anansi gives an AI system a durable understanding of how an organization actually
works — and how that changed over time.
You feed it the exhaust your company already produces: conversations, docs, tickets,
meeting transcripts. Anansi turns that into structured memory your agent can read
before it answers, and keeps every version of it. So you can ask not just "what is
our escalation process?" but "what did we think it was in March, and when did it
change?" — and get an answer with a citation.
TL;DR: MIT-licensed self-hostable core, with a commercial enterprise/hosted layer. See License .
POST /v1/ingest # remember this
GET /v1/context # what do you know that's relevant right now?
ingest returns 202 immediately and does the expensive work on a queue, so it never
sits in your response path. context returns a compact, already-synthesized profile you
can paste straight into a system prompt — not a pile of chunks to rank yourself.
Self-hosted, MIT licensed, runs on Postgres and Redis. About five minutes to a working
instance, no signup.
What that citation actually buys you — the entity graph carries two independent
time axes, so you can ask what was true and what the system knew, separately:
The entity graph and temporal query results (the two-axis reasoning above) are a Pro+ feature. Self-hosted installs default to the enterprise plan (see Plan limits ) and get it automatically; on the hosted service it requires a paid tier.
Quickstart — self-hosted, about 5 minutes
No account and no API key from us. Everything below runs on your machine.
Timings are measured, not aspirational: the API image builds from source in about
90 seconds on a warm Docker, and the embedding model is a 274 MB download. The
five-minute figure assumes option A or B in step 2 — option C pulls a 7 GB
image and takes considerably longer.
git clone https://github.com/g-33-L/anansi.git
cd anansi
docker compose up -d
That brings up PostgreSQL, Redis, and the API. Migrations run automatically on first
boot. The API and docs serve at http://localhost:3000 .
2. Give it something to embed with
Anansi needs an embedding model, and Compose does not start one for you. Skip this
step and ingest will appear to succeed — it returns 202 because embedding is
asynchronous — while retrieval fails with 503 . Open
/status ; it reports the embedding backend explicitly,
so you can see the problem rather than infer it.
Three ways to satisfy it, ordered by how long they actually take. Only the first two keep
this quickstart inside five minutes.
ollama pull nomic-embed-text
Nothing else to configure: Compose already points the container at
host.docker.internal:11434 . This is the fastest path and the one to prefer if you have
Ollama installed.
printf ' DEPLOYMENT_MODE=hybrid\nINFERENCE_LOCATION=local\nEMBEDDING_LOCATION=cloud\nNOMIC_API_KEY=your_key_here\n ' >> .env
docker compose up -d api
All four lines are required. DEPLOYMENT_MODE defaults to local , which forbids
cloud providers outright — supplying NOMIC_API_KEY without switching to hybrid is a
deliberate startup failure, not an oversight, so the container refuses to boot and tells
you so. hybrid is what lets you mix local inference with cloud embeddings.
Confirm it took effect — the startup log states the resolved mode:
[startup] Deployment mode: hybrid (inference=local, embedding=cloud, telemetry=allowed)
Note that this sends the text you ingest to Nomic. Use A or C if that matters.
This is the slow path. ollama/ollama:latest is about 7 GB because it ships GPU
runtimes, and on a normal connection the pull alone takes ten minutes or more. Right
choice if you want everything in Compose and nothing on your host — not the right choice
if you are trying Anansi for the first time.
echo " OLLAMA_BASE_URL=http://ollama:11434 " >> .env
docker compose --profile local-ai up -d # pulls the 7 GB image
docker compose exec ollama ollama pull nomic-embed-text # 274 MB
Appending is safe even if OLLAMA_BASE_URL is already set — Compose takes the last
definition. If you set it after the API was already running, restart it with
docker compose up -d api so it picks up the new address.
Whichever you pick, nomic-embed-text (274 MB) is all that ingest and retrieval need. The
much larger chat model ( llama3.1:8b , ~4.7 GB) is only used to synthesize the static and
dynamic profiles — pull it later with ollama pull llama3.1:8b when you want those.
Keys live in your own database — this does not contact any hosted service:
docker compose exec api node dist/scripts/seed-dev-key.js you@example.com
It prints a key beginning ans_ . Export it:
export ANANSI_API_KEY=ans_...
Each email address gets its own workspace , and memory never crosses between them.
Re-running the command with the same email issues another key into the same workspace;
running it with a different email gives you a key that cannot see anything you stored
earlier. If your data seems to have vanished, check which email the key came from.
curl -X POST http://localhost:3000/v1/ingest \
-H " Authorization: Bearer $ANANSI_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"userId":"user_123","content":"User is building a voice agent. Prefers TypeScript. Team of 4.","sourceType":"conversation"} '
Returns 202 immediately — embedding happens in the background.
curl -G http://localhost:3000/v1/context \
-H " Authorization: Bearer $ANANSI_API_KEY " \
--data-urlencode " userId=user_123 " \
--data-urlencode " q=what is the user building? "
relevant comes back populated:
{ "relevant" : [ { "content" : " User is building a voice agent. Prefers TypeScript. Team of 4. " ,
"similarity" : 0.4821 } ], "static" : [], "dynamic" : [] }
Two things are worth knowing about that response:
similarity: 0 means the embedding had not landed yet. Embedding is asynchronous, so a query issued within a second of ingest can be answered by keyword search alone. Ask again and you will see a real cosine score. Search is hybrid, so you get an answer either way rather than an empty result.
static and dynamic stay empty until a chat model is available for synthesis (step 2). That is expected, not a failure.
To prove the semantic half is genuinely working, ask something that shares no words with
what you stored:
curl -G http://localhost:3000/v1/context \
-H " Authorization: Bearer $ANANSI_API_KEY " \
--data-urlencode " userId=user_123 " \
--data-urlencode " q=which coding language do they like? "
That scores higher ( 0.5915 ) than the keyword-overlapping question, because nothing in
it matches literally — only in meaning.
Check /status first: it reports Postgres, Redis, the
queue, and the embedding backend, and returns 503 when any of them is down. A 503
from /v1/context names the failing dependency and how to fix it directly in the
response body.
Local (no keys, no account)
Optional / external
API, Postgres, Redis, workers
✅ started by docker compose up -d
Embedding + synthesis (Ollama)
✅ option A/C above
Nomic hosted embeddings (option B), Cerebras/GitHub Models for synthesis
Everything in Quickstart above
✅ fully local
Slack / Notion / Google Docs / Linear connectors
require their own OAuth credentials
Hosted hybrid search hints, error reporting
Sentry ( SENTRY_DSN ) is optional and off by default
The full local path — Docker + Ollama, no connectors configured — never sends ingested content off your machine. Once you add a cloud embedding/LLM provider or a connector, that surface's data leaves the box; see Security for exactly what each DEPLOYMENT_MODE allows.
The Compose defaults are deliberately development-only and are sufficient to start a
disposable local stack without creating .env . The cryptographic values baked into
docker-compose.yml are published in this repository and therefore public — never use
them outside local development. For anything persistent, copy
.env.example , generate distinct values for ENCRYPTION_KEY ,
API_KEY_HMAC_SECRET , CSRF_SIGNING_KEY , and QUERY_API_KEY
( openssl rand -hex 32 each), and read
docs/enterprise/self-hosting.md .
Never change ENCRYPTION_KEY after first install — all stored connector tokens are
encrypted with it.
If you run Ollama on your host rather than via the local-ai profile, the Compose
default ( host.docker.internal ) already points at it. A .env written for host-run
pnpm dev will contain localhost:11434 , which inside a container means the container
itself — Compose interpolates that file, so the API silently cannot reach your host
Ollama. The 503 body names the address it tried, which is how you spot this.
Prefer running from source with pnpm dev ? See CONTRIBUTING.md — note
that pnpm test needs DATABASE_URL and REDIS_URL in your shell, and that the suite
TRUNCATE s the local database, so do not point it at anything you care about.
Using the TypeScript SDK ( packages/sdk ):
import AnansiMemory from "anansi-memory" ;
const memory = new AnansiMemory ( {
apiKey : process . env . ANANSI_API_KEY ,
baseUrl : "http://localhost:3000" , // required when self-hosting — see below
} ) ;
await memory . ingest ( {
userId : "user_123" ,
content : "User is building a voice agent. Prefers TypeScript. Team of 4." ,
sourceType : "conversation" ,
} ) ;
const ctx = await memory . context ( { userId : "user_123" , q : "what is the user building?" } ) ;
const systemPrompt = `You are a helpful assistant.\n\n ${ memory . formatForPrompt ( ctx ) } ` ;
Self-hosters: set the base URL. Every client defaults to the hosted API at
https

[truncated]
