---
source: "https://github.com/letterstory/lettertrace"
hn_url: "https://news.ycombinator.com/item?id=49299569"
title: "Show HN: Measure your AI search with BYOK and OSS (free)"
article_title: "GitHub - letterstory/lettertrace: Open source BYOK AEO telemetry · GitHub"
author: "mathewpregasen"
captured_at: "2026-08-14T15:42:26Z"
capture_tool: "hn-digest"
hn_id: 49299569
score: 2
comments: 0
posted_at: "2026-08-14T14:57:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Measure your AI search with BYOK and OSS (free)

- HN: [49299569](https://news.ycombinator.com/item?id=49299569)
- Source: [github.com](https://github.com/letterstory/lettertrace)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T14:57:16Z

## Translation

タイトル: HN を表示: BYOK と OSS を使用して AI 検索を測定する (無料)
記事タイトル: GitHub - レターストーリー/レタートレース: オープンソース BYOK AEO テレメトリ · GitHub
説明: オープンソースの BYOK AEO テレメトリ。 GitHub でアカウントを作成して、letterstory/letterrace の開発に貢献してください。
HN テキスト: 私たちが作った楽しいものを共有します:)

記事本文:
GitHub - レターストーリー/レターレース: オープンソース BYOK AEO テレメトリ · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
レターストーリー
/
文字跡
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
226 コミット 226 コミット .github/ workflows .github/ workflows app app cli cli コンポーネント コンポーネント docs docs lib lib public public scripts scripts supabase supabas

e .dockerignore .dockerignore .env.example .env.example .eslintrc.json .eslintrc.json .gitignore .gitignore .gitkeep .gitkeep .nvmrc .nvmrc Dockerfile Dockerfile ライセンス ライセンス README.md README.md middleware.ts middleware.ts next.config.mjs next.config.mjs package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json vercel.json vercel.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI アシスタントの回答にブランドがどのように表示されるかをオープンソースの持ち込みキーでモニタリングします。
トピックを追跡する · 人々が実際に尋ねる質問を AI に自動生成する · 長期的な傾向を監視する · 競合他社の発言権の割合をベンチマークする。
Lettertrace は、自己ホスト可能な AEO ツールであり、AI メンション (別名: アンサー エンジン最適化/生成エンジン最適化) の診断と監視にのみ焦点を当てています。あなたのブランドといくつかのトピックについて説明します。 Lettertrace は、ユーザーが ChatGPT または Claude に尋ねる可能性のある現実的なプロンプトを生成し、独自の API キーを使用してそれらのモデルに対してプロンプトを実行し、ブランドと競合他社がいつ言及されたかを検出し、可視性、感情、声のシェアが時間の経過とともにどのように変化するかをグラフ化します。
🔓 オープンソース (MIT) と BYOK では、独自の Anthropic / OpenAI / Google / Perplexity キー、または代わりに単一の LLM ルーター キー ( Concentrate ) を使用します。いずれの場合も、保存時に暗号化され、インフラストラクチャから離れることはありません。
🧠 マルチモデル、クエリ Claude (Anthropic)、ChatGPT (OpenAI)、Gemini および Google AI 概要 (両方とも Google キー上)、および Perplexity Sonar。プロバイダーを簡単に追加できます。
🧩 トピック → バリエーション 、各トピックについて人々が AI に尋ねるさまざまな質問を自動生成します。
📈 時間の経過に伴う傾向、知名度、声のシェア、知名度、

実行中のセンチメント。
⚔️ 競合他社のベンチマーク 、競合他社を取り込み、それぞれがどのくらいの頻度で現れるかを確認します。
🏢 複数の組織、1 つのアカウントで多くのブランド/ドメインを追跡し、サイドバーからそれらを切り替えることができます。
🔎 Web 検索 + ソース アトリビューション : ネイティブ Web 検索をオンにしてモデルをクエリし、引用されている正確なソースをキャプチャします。これにより、名前が付けられていない場合でも、どの投稿が回答の原因となったか、自分のサイトが使用されているかどうかを確認できます。
⏱️ スケジュールされた監視、cron エンドポイント経由で毎日/毎週実行されます。
コンセプト
それは何ですか
組織（プロジェクト）
ブランドのワークスペース: ブランド名、エイリアス、ドメイン、デフォルト モデル、スケジュール。 1 つのアカウントに複数のアカウントを持つことができ、サイドバー セレクターによってダッシュボード全体がアカウント間で切り替わり、＋ 新しい組織ではセットアップ ウィザードが再度開きます。
競合他社
ベンチマーク対象のブランド (名前 + 別名)。
トピックス
監視したい対象（例：「プロジェクト管理ソフトウェア」）。
プロンプト（バリエーション）
トピックに対して生成された自然な質問。実際にモデルに送信されるクエリ。
走る
1 回の実行: すべてのアクティブなプロンプト → モデル → メンションの検出 → 保存。
言及
回答内で検出されたブランド/競合他社への言及。回数、知名度、センチメント、推奨されたかどうかが含まれます。
メンション検出の仕組み
モデルが返す各答えに対して、Lettertrace は次のようになります。
決定的検出 、ブランド名と各競合他社の名前 + エイリアス (単語境界、大文字と小文字を区別しない) を照合し、出現数と最初の位置 (目立つ) を記録します。
LLM エンリッチメント。言及されたエンティティについては、構造化された呼び出しによってセンチメントと、回答がエンティティを推奨したかどうかが分類されます。
集計、可視性 (メンション率)、発言のシェア、平均注目度、センチメントが実行ごとに計算され、時間の経過とともに傾向が表示されます。
Next.js 14 (アプリルート)

r、TypeScript) · Tailwind CSS · Recharts
Supabase、Postgres、認証、および行レベルのセキュリティ
保存時に AES-256-GCM で暗号化された BYOK プロバイダー キー
Anthropic ( @anthropic-ai/sdk ) + OpenAI ( openai ) SDK アダプターに加え、Google Gemini (Gemini モデル + Google AI 概要、Google 検索基盤経由) および Perplexity Sonar (常に検索基盤の実際のソース URL) 用の依存関係のない REST アダプター
supabase.com でプロジェクトを作成します。 「設定」→「API 取得」から:
プロジェクトURL → NEXT_PUBLIC_SUPABASE_URL
anon 公開鍵 → NEXT_PUBLIC_SUPABASE_ANON_KEY
service_role キー → SUPABASE_SERVICE_ROLE_KEY (スケジュールされた実行の場合にのみ必要)
Supabase SQL エディターを開き、 supabase/schema.sql の内容を実行します。すべてのテーブル、インデックス、行レベル セキュリティ ポリシー、およびサインアップ時にプロファイルを自動作成するトリガーを作成します。再実行しても安全です。
電子メールの確認: 最もスムーズなローカル エクスペリエンスを実現するには、「認証」→「プロバイダー」→「電子メール」で「電子メールの確認」を無効にするか、リンク ( /auth/callback によって処理されます) を介して確認します。
サインイン画面には、電子メール + パスワードとともに Google と GitHub が表示されます。どちらもオプションです。Supabase でプロバイダーが有効になっていない場合、そのボタンをクリックするとエラーが発生するだけなので、設定する予定がない場合は、app/login/auth-form.tsx の oauthProviders から削除してください。
新しい環境変数は関係しません。クライアント シークレットは、このリポジトリではなく、Supabase に存在します。
1. 各プロバイダーにアプリを登録します。どちらもあなたのコールバックではなく、Supabase のコールバックを指します。つまり、1 つの登録でローカルの開発と運用がカバーされます。
https://<プロジェクト参照>.supabase.co/auth/v1/callback
GitHub → 設定 → 開発者設定 → OAuth アプリ → 新しい OAuth アプリ。 「デバイスフローを有効にする」はチェックを外したままにしておきます。
Google → Cloud Console → API とサービス → 認証情報 → OAuth クライアント ID

(Webアプリケーション)。 https://<project-ref>.supabase.co を承認された JavaScript オリジンとして追加します。同意画面では、デフォルトの機密性のないスコープ ( email 、 profile 、 openid ) のみをリクエストします。これらのスコープには検証レビューは必要ありませんが、100 人を超えるユーザーを受け入れるにはアプリを本番環境に公開する必要があり、同意画面のリンク内のドメインは Search Console で検証する必要があります。
2. Supabase でプロバイダーを有効にします。認証 → プロバイダー → Google / GitHub: オンに切り替えて、各クライアント ID とシークレットを貼り付けます。プロジェクトの所有者または管理者の役割が必要です。他の役割の場合、フィールドはグレー表示されます。 「電子メールのないユーザーを許可する」をオフのままにします。GitHub は、プライベート電子メールを持つユーザーに対しても @users.noreply.github.com アドレスを返すため、 handle_new_user は常に profile に書き込む内容を持ちます。
3. 独自のリダイレクト URL を許可リストに登録します。認証→URL設定→リダイレクトURL。これは 2 番目のホップ (Supabase → アプリ) であり、ステップ 1 とは別のものです。これが欠けていると、サインインが行き詰まってしまう最も一般的な原因になります。
http://localhost:3000/auth/callback
https://your-domain.com/auth/callback
その画面でサイト URL を運用オリジンに設定し、NEXT_PUBLIC_SITE_URL がそれに一致することを確認します。/auth/callback はリクエスト オリジンよりもそれを優先します。そうでない場合、プロキシの背後にある内部デプロイメント ホストに解決されます。
パスワードを使用してサインアップし、その後同じアドレスでソーシャル プロバイダーを使用するユーザーは、プロバイダーが電子メールを認証済みとして報告する場合、Supabase によって 1 つのアカウントにリンクされます。
cp .env.example .env.local
Supabase の値を入力し、シークレットを生成します。
# 保存時の BYOK プロバイダー キーを暗号化する 32 バイトのキー
openssl rand -base64 32 # -> ENCRYPTION_KEY
# スケジュールされた実行エンドポイントの共有シークレット
openssl rand -hex 32 # -> CR

オン_シークレット
デフォルトではサードパーティによる追跡はありません。ホストされているlettertrace.comは、
RB2B のビジネス訪問者識別ピクセルが公開されています
マーケティング ページのみ ( / 、 /privacy 、 /terms )、ゲートオン
NEXT_PUBLIC_RB2B_KEY 。 .env.example と同様に、それを未設定のままにしておきます。
セルフホスト展開では、サードパーティの追跡は不要です。それは決して実行されません
認証されたアプリ ( /dashboard 、 /login 、…)。
代わりにコンテナとして実行しますか? Docker を使用して実行するまでスキップしてください
— ステップ 1 ～ 3 は引き続き適用されますが、ステップ 4 は単一の Docker run になります。
ノード 22 LTS を使用します ( .nvmrc を参照)。ノード 23+ には undic があります
応答中にプロバイダー接続が断続的に切断される回帰
（「時期尚早な終了」）。 nvm の場合:
nvm install # .nvmrc を選択します (ノード 22)
NVMの使用
npmインストール
npm 実行開発
http://localhost:3000 を開いてアカウントを作成すると、ダッシュボードが表示されます。
[設定] → Anthropic、OpenAI、Google、Perplexity API キー (保存時に検証され、保存時に暗号化される) または 1 つの LLM ルーター キーを追加し、ブランドとプロジェクト (名前、エイリアス、および監視する応答エンジン (Gemini または Google AI 概要など)) を入力します。端末の方がいいですか?レタートレース キー セット anthropic も同じことを行います。
競合他社 → ベンチマークの対象となるブランドを追加します。
「トピック」 → トピックを追加し、「バリエーションの生成」をクリックしてプロンプトを自動作成します (または独自のプロンプトを追加します)。
「実行」→「今すぐモニターを実行」。完了すると、概要に可視性、声のシェア、センチメント、トピックごとの内訳が表示されます。
前提条件: supabase/schema.sql を含む Supabase プロジェクト
適用済み - ステップ 1 および 2
上記、無料枠で約 2 分。 Supabase は認証を提供し、
データベースなので、コンテナの外に存在します。
docker run -p 3000:3000 \
-e NEXT_PUBLIC_SUPABASE_URL= " https://xxxx.supabase.co " \
-e NEXT_PUBLIC_SUPABASE_ANON_KEY= " eyJ..

。 " \
-e SUPABASE_SERVICE_ROLE_KEY= " eyJ... " \
-e ENCRYPTION_KEY= " $( openssl rand -base64 32 ) " \
-e NEXT_PUBLIC_SITE_URL= " http://localhost:3000 " \
ghcr.io/letterstory/letterrace
http://localhost:3000 を開き、アカウントを作成して追加します
いつものように設定でプロバイダー キーを入力します。どちらの方法でも BYOK なので、画像は
モデル認証情報は含まれていない状態で出荷されます。
すでに .env をお持ちですか? --env-file は入力する手間が少なく、キーを入力する必要がなくなります。
シェルの履歴:
docker run -p 3000:3000 --env-file .env ghcr.io/letterstory/letterrace
重要
ENCRYPTION_KEY は、データベースに保存されているプロバイダー キーを復号化します。生成する
一度保存してください。紛失した場合、保存されているすべての BYOK キーは無効になります。
判読できないため再入力する必要があります。変えても同じです。
何が必要で何が不必要なのか
変数
必要な
NEXT_PUBLIC_SUPABASE_URL
すべて
NEXT_PUBLIC_SUPABASE_ANON_KEY
すべて
SUPABASE_SERVICE_ROLE_KEY
スケジュールされた実行、管理画面
暗号化キー
BYOK プロバイダー キーの保存 (32 バイト、base64)
NEXT_PUBLIC_SITE_URL
認証リダイレクトと OG メタデータを修正します。ユーザーが実際にアクセスする URL に設定します。
CRON_SECRET
スケジュールされた実行を設定した場合のみ (下記)
.env.example 内の他のものはすべてオプションです: TRIAL_*
キーは、自分のプロバイダー アカウントで無料実行を配布するために存在します。
ADMIN_* / RESEND_API_KEY 値は演算子 al のみを有効にします

[切り捨てられた]

## Original Extract

Open source BYOK AEO telemetry. Contribute to letterstory/lettertrace development by creating an account on GitHub.

Sharing something fun we made :)

GitHub - letterstory/lettertrace: Open source BYOK AEO telemetry · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
letterstory
/
lettertrace
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
226 Commits 226 Commits .github/ workflows .github/ workflows app app cli cli components components docs docs lib lib public public scripts scripts supabase supabase .dockerignore .dockerignore .env.example .env.example .eslintrc.json .eslintrc.json .gitignore .gitignore .gitkeep .gitkeep .nvmrc .nvmrc Dockerfile Dockerfile LICENSE LICENSE README.md README.md middleware.ts middleware.ts next.config.mjs next.config.mjs package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json vercel.json vercel.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Open-source, bring-your-own-key monitoring for how your brand shows up in AI assistant answers.
Track topics · auto-generate the questions people actually ask AI · watch trends over time · benchmark competitors' share of voice.
Lettertrace is a self-hostable AEO tool, focused purely on diagnosing and monitoring AI mentions (a.k.a. Answer Engine Optimization / Generative Engine Optimization). You describe your brand and a few topics; Lettertrace generates realistic prompts a person might ask ChatGPT or Claude, runs them against those models with your own API key , detects when your brand and your competitors get mentioned, and charts how your visibility, sentiment, and share of voice move over time.
🔓 Open source (MIT) and BYOK , you bring your own Anthropic / OpenAI / Google / Perplexity keys — or a single LLM router key ( Concentrate ) instead. Either way they're encrypted at rest and never leave your infrastructure.
🧠 Multi-model , query Claude (Anthropic), ChatGPT (OpenAI), Gemini and Google AI Overviews (both on your Google key), and Perplexity Sonar. Add more providers easily.
🧩 Topics → variations , auto-generate the different questions people ask AI about each topic.
📈 Trends over time , visibility, share of voice, prominence, and sentiment across runs.
⚔️ Competitor benchmarking , ingest competitors and see how often each shows up.
🏢 Multiple organizations , one account can track many brands/domains and switch between them from the sidebar.
🔎 Web search + source attribution , query the models with their native web search on and capture the exact sources they cite, so you can see which posts drove an answer, and whether your own site is being used, even when you aren't named.
⏱️ Scheduled monitoring , daily/weekly runs via a cron endpoint.
Concept
What it is
Organization (project)
A brand's workspace: brand name, aliases, domain, default model, schedule. An account can have several, the sidebar selector switches the whole dashboard between them, and ＋ New organization re-opens the setup wizard.
Competitors
Brands you benchmark against (name + aliases).
Topics
Subjects you want to monitor (e.g. "project management software").
Prompts (variations)
Natural questions generated for a topic, the queries actually sent to the model.
Runs
One execution: every active prompt → the model → detect mentions → store.
Mentions
A detected reference to your brand/competitor in an answer, with count, prominence, sentiment, and whether it was recommended.
How mention detection works
For each answer the model returns, Lettertrace:
Deterministic detection , matches your brand's and each competitor's name + aliases (word-boundary, case-insensitive), recording occurrence count and first position (prominence).
LLM enrichment , for the entities that were mentioned, a structured call classifies sentiment and whether the answer recommended them.
Aggregation , visibility (mention rate), share of voice , average prominence, and sentiment are computed per run and trended over time.
Next.js 14 (App Router, TypeScript) · Tailwind CSS · Recharts
Supabase , Postgres, Auth, and Row Level Security
BYOK provider keys encrypted with AES-256-GCM at rest
Anthropic ( @anthropic-ai/sdk ) + OpenAI ( openai ) SDK adapters, plus dependency-free REST adapters for Google Gemini (Gemini models + Google AI Overviews, via Google Search grounding) and Perplexity Sonar (always search-grounded, real source URLs)
At supabase.com , create a project. From Settings → API grab:
Project URL → NEXT_PUBLIC_SUPABASE_URL
anon public key → NEXT_PUBLIC_SUPABASE_ANON_KEY
service_role key → SUPABASE_SERVICE_ROLE_KEY (only needed for scheduled runs)
Open the Supabase SQL Editor and run the contents of supabase/schema.sql . It creates all tables, indexes, Row Level Security policies, and a trigger that auto-creates a profile on sign-up. It's safe to re-run.
Email confirmation: for the smoothest local experience, disable "Confirm email" under Authentication → Providers → Email , or confirm via the link (handled by /auth/callback ).
The sign-in screen offers Google and GitHub alongside email + password. Both are optional — if a provider isn't enabled in Supabase, its button will simply error when clicked, so remove it from oauthProviders in app/login/auth-form.tsx if you don't plan to configure it.
No new environment variables are involved: the client secrets live in Supabase, not in this repo.
1. Register the app with each provider. Both point at Supabase's callback, not yours — which means one registration covers local development and production:
https://<project-ref>.supabase.co/auth/v1/callback
GitHub → Settings → Developer settings → OAuth Apps → New OAuth App. Leave "Enable Device Flow" unchecked.
Google → Cloud Console → APIs & Services → Credentials → OAuth client ID (Web application). Add https://<project-ref>.supabase.co as an authorized JavaScript origin. On the consent screen, request only the default non-sensitive scopes ( email , profile , openid ) — those need no verification review, but the app must be published to In production to accept more than 100 users, and any domain in your consent-screen links must be verified in Search Console.
2. Enable the providers in Supabase. Authentication → Providers → Google / GitHub: toggle on and paste each client ID and secret. Requires the Owner or Administrator role on the project; other roles see the fields greyed out. Leave "Allow users without an email" off — GitHub returns a @users.noreply.github.com address even for users with private emails, so handle_new_user always has something to write into profiles .
3. Allowlist your own redirect URLs. Authentication → URL Configuration → Redirect URLs. This is the second hop (Supabase → your app) and is separate from step 1; missing it is the most common cause of a sign-in that dead-ends:
http://localhost:3000/auth/callback
https://your-domain.com/auth/callback
Set Site URL to your production origin while you're on that screen, and make sure NEXT_PUBLIC_SITE_URL matches it — /auth/callback prefers it over the request origin, which would otherwise resolve to the internal deployment host behind a proxy.
Users who sign up with a password and later use a social provider with the same address are linked into one account by Supabase, provided the provider reports the email as verified.
cp .env.example .env.local
Fill in Supabase values and generate secrets:
# 32-byte key that encrypts BYOK provider keys at rest
openssl rand -base64 32 # -> ENCRYPTION_KEY
# shared secret for the scheduled-run endpoint
openssl rand -hex 32 # -> CRON_SECRET
No third-party tracking by default. The hosted lettertrace.com runs the
RB2B business-visitor identification pixel on its public
marketing pages only ( / , /privacy , /terms ), gated on
NEXT_PUBLIC_RB2B_KEY . Leave that unset — as .env.example does — and a
self-hosted deployment ships zero third-party tracking. It never runs on
the authenticated app ( /dashboard , /login , …).
Running it as a container instead? Skip to Run it with Docker
— steps 1–3 still apply, step 4 becomes a single docker run .
Use Node 22 LTS (see .nvmrc ). Node 23+ has an undici
regression that intermittently drops provider connections mid-response
("Premature close"). With nvm :
nvm install # picks up .nvmrc (Node 22)
nvm use
npm install
npm run dev
Open http://localhost:3000 , create an account, and you'll land on the dashboard.
Settings → add your Anthropic, OpenAI, Google, and/or Perplexity API key (verified on save, encrypted at rest), or one LLM router key , then fill in your brand & project (name, aliases, and the answer engine to monitor with, including Gemini or Google AI Overviews). Prefer a terminal? lettertrace keys set anthropic does the same thing.
Competitors → add the brands you want to benchmark against.
Topics → add a topic and click Generate variations to auto-create prompts (or add your own).
Runs → Run monitor now . When it finishes, the Overview fills in with visibility, share of voice, sentiment, and per-topic breakdowns.
Prerequisites: a Supabase project with supabase/schema.sql
applied — steps 1 and 2
above, about two minutes on the free tier. Supabase provides auth and the
database, so it lives outside the container.
docker run -p 3000:3000 \
-e NEXT_PUBLIC_SUPABASE_URL= " https://xxxx.supabase.co " \
-e NEXT_PUBLIC_SUPABASE_ANON_KEY= " eyJ... " \
-e SUPABASE_SERVICE_ROLE_KEY= " eyJ... " \
-e ENCRYPTION_KEY= " $( openssl rand -base64 32 ) " \
-e NEXT_PUBLIC_SITE_URL= " http://localhost:3000 " \
ghcr.io/letterstory/lettertrace
Open http://localhost:3000 , create an account, and add
your provider key in Settings as usual — it is BYOK either way, so the image
ships with no model credentials in it.
Already have a .env ? --env-file is less to type and keeps keys out of your
shell history:
docker run -p 3000:3000 --env-file .env ghcr.io/letterstory/lettertrace
Important
ENCRYPTION_KEY decrypts the provider keys stored in your database. Generate
it once and keep it. If you lose it, every stored BYOK key becomes
unreadable and has to be re-entered; if you change it, the same.
What's required and what isn't
Variable
Needed for
NEXT_PUBLIC_SUPABASE_URL
Everything
NEXT_PUBLIC_SUPABASE_ANON_KEY
Everything
SUPABASE_SERVICE_ROLE_KEY
Scheduled runs, admin surfaces
ENCRYPTION_KEY
Storing BYOK provider keys (32 bytes, base64)
NEXT_PUBLIC_SITE_URL
Correct auth redirects and OG metadata — set it to the URL users actually visit
CRON_SECRET
Only if you wire up scheduled runs (below)
Everything else in .env.example is optional: the TRIAL_*
keys exist to hand out free runs on your own provider account, and the
ADMIN_* / RESEND_API_KEY values only enable operator al

[truncated]
