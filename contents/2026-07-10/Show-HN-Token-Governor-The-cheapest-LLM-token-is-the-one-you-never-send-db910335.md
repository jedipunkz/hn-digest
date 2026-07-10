---
source: "https://github.com/shipitamit/llm-token-governor"
hn_url: "https://news.ycombinator.com/item?id=48861389"
title: "Show HN: Token Governor – The cheapest LLM token is the one you never send"
article_title: "GitHub - shipitamit/llm-token-governor: Re-engineer LLM prompts to spend the fewest tokens that still answer. A governing gateway with auth, budgets, caching, and streaming. Works with Anthropic, OpenAI, Gemini, and any OpenAI-compatible provider. · GitHub"
author: "mavster26"
captured_at: "2026-07-10T16:15:25Z"
capture_tool: "hn-digest"
hn_id: 48861389
score: 1
comments: 0
posted_at: "2026-07-10T15:35:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Token Governor – The cheapest LLM token is the one you never send

- HN: [48861389](https://news.ycombinator.com/item?id=48861389)
- Source: [github.com](https://github.com/shipitamit/llm-token-governor)
- Score: 1
- Comments: 0
- Posted: 2026-07-10T15:35:55Z

## Translation

タイトル: HN を表示: トークン ガバナー – 最も安価な LLM トークンは決して送信しないものです
記事のタイトル: GitHub - shipitamit/llm-token-governor: LLM を再設計すると、応答するトークンを最小限に抑えるよう求められます。認証、予算、キャッシュ、ストリーミングを備えた管理ゲートウェイ。 Anthropic、OpenAI、Gemini、および任意の OpenAI 互換プロバイダーで動作します。 · GitHub
説明: LLM を再設計すると、応答するトークンを最小限に抑えるように求められます。認証、予算、キャッシュ、ストリーミングを備えた管理ゲートウェイ。 Anthropic、OpenAI、Gemini、および任意の OpenAI 互換プロバイダーで動作します。 - shipitamit/llm-token-governor

記事本文:
GitHub - shipitamit/llm-token-governor: LLM プロンプトを再設計し、応答するトークンを最小限に抑えるように要求します。認証、予算、キャッシュ、ストリーミングを備えた管理ゲートウェイ。 Anthropic、OpenAI、Gemini、および任意の OpenAI 互換プロバイダーで動作します。 · GitHub
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
却下エール

RT
{{ メッセージ }}
シピタミット
/
llm-トークン-ガバナー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット docs docs サンプル サンプル ノード モジュール ノード モジュール スクリプト スクリプト サーバー サーバー src src テスト test Dockerfile Dockerfile ライセンス ライセンス README.md README.md ゲートウェイ.callers.example.json ゲートウェイ.callers.example.json ゲートウェイ.価格.example.json ゲートウェイ.価格.example.json インデックス.html インデックス.html パッケージ-ロック.json package-lock.json package.json package.json vite.config.js vite.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM を再設計すると、応答するトークンを最小限に抑えるよう求められます。
認証、予算、キャッシュ、ストリーミングを備えた管理ゲートウェイ
Anthropic、OpenAI、Gemini、および任意の OpenAI 互換プロバイダー。さらにライブA/B、
保存された制約プリセット、およびプロンプトにフラグを付けるバッチトリアージ
モデルが必要だった。
llm · トークン · コスト最適化 · ai ゲートウェイ · llmops · openai · anthropic · gemini · ストリーミング · プロンプトエンジニアリング
シングル モード — 制約付きリクエストを構築し、予測される節約額を確認し、ラッパーのコストが節約額を上回る場合の判定を取得します。
プリセット — 制約設定を保存/ロードします。これらはブラウザ内に残ります。
バッチ モード — プロンプトのリストをスコアリングし、確定的なプロンプトにフラグを立て、CSV/マークダウンをエクスポートし、フラグの付いた行ごとに推奨されるスクリプトを取得します。
ライブ A/B — サーバー側プロキシ経由でプロンプトを 2 回 (生と管理) 実行し、設定したプロバイダーを使用して実際のトークン数を報告します。
固定の 1 画面ダッシュボード (ページ スクロールなし - パネルは内部でスクロールします)。頂上
バーにはモード タブと価格設定/スケール入力が表示されます。以下に 3 つのパネルを示します。
プロンプト＆コン

Strats — プロンプトに簡潔さ、出力形状、温度、トグルを加えたもの。
予測される節約 — ゲージ、「価値がある/節約よりコストが高い」の判定、およびコストの受け取り前/後。
組み立てられたリクエストとライブ A/B — コピー ボタンを備えた生成されたシステム プロンプト、および返された実際のトークン数をレポートする生の A/B と管理された A/B です。
バッチ モードに切り替えると、決定的なフラグを立てるトリアージ テーブルになります。
プロンプトを表示します。 (画像はレイアウトのモックアップであり、スクリーンショットではありません。アプリを実行してください。
ライブ、インタラクティブ バージョン。)
Node.js 18+ (組み込みの fetch を使用)
1 つの LLM プロバイダーの API キー (ライブ A/B パネルにのみ必要)
npmインストール
cp .env.example .env # 次に、.env を編集します (下記の「プロバイダー」を参照)
プロバイダー
サーバーはプロバイダーに依存しません。これらを .env に設定します。
フロントエンドは決して変更されません。正規化されたリクエストを /api/messages に POST します。
そして、server/providers.js 内のアダプターは、選択したプロバイダーとの間で変換されます。
そして正規化された応答を返します。ネイティブ プロバイダーを追加するには、その 1 つのファイルを編集します。
現在、ほぼすべてのプロバイダーが OpenAI 互換 API を提供しているため、設定することで機能します。
ベース URL とモデル — コード変更なし。アクセスしたいプロバイダーのみ
独自のネイティブ プロトコルにはアダプターが必要です。
すべての OpenAI 互換プロバイダーは LLM_PROVIDER=openai を使用します。
#OpenAI
LLM_MODEL=gpt-4o-mini
# Gemini (Google AI Studio) — 注: 末尾のスラッシュはありません (コードは /chat/completions を追加します)
LLM_BASE_URL=https://generative language.googleapis.com/v1beta/openai
LLM_MODEL=gemini-2.5-flash # キーがアクセスできるモデルを使用します
#グロク
LLM_BASE_URL=https://api.groq.com/openai/v1
#オープンルーター
LLM_BASE_URL=https://openrouter.ai/api/v1
# 地元のオラマ
LLM_BASE_URL=http://localhost:11434/v1 LLM_API_KEY=ollama
注:
ジェミニの末尾のスラッシュ。 OpenAI SDK のため、Google 自身のドキュメントでは .../v1beta/openai/ を使用します。

チャット/完了を追加します。このサーバーは /chat/completions を追加するため、末尾のスラッシュを削除しないと、二重スラッシュ 404 が表示されます。
一部の新しい OpenAI モデルは max_completion_tokens を想定しています。そのエラーが発生した場合は、OPENAI_MAX_TOKENS_PARAM=max_completion_tokens を設定します。
モデルの入手可能性はキーによって異なります。モデル 404s の場合は、プロバイダーの現在のリストから 1 つ選択します。 npm run check:live は、端から端までの配線を確認します。
ネイティブ Gemini SDK (generateContent) はサポートされていません。アプリが OpenAI-compat モードではなく Google GenAI SDK のネイティブ API を使用している場合は、クライアントを OpenAI-compat モードに切り替えるか、ネイティブ Gemini アダプター (ボディマッパーと :streamGenerateContent パススルー、Anthropic アダプターと並行) を追加します。
npm 実行開発
http://localhost:5173 の Vite フロントエンド
http://localhost:8787 の Express API (Vite プロキシ /api から)
npm run build # 出力 dist/
npm start # Express は PORT で dist/ AND /api を提供します (デフォルトは 8787)
導入前チェック
発送前にキーとモデルが配線されていることを確認してください。
npm run check # 実行中のサーバーの /api/health にヒットします
npm run check:start # サーバー自体を起動し、チェックし、停止します
npm run check:live # 小さな実際のプロンプトも送信します (いくつかのトークンがかかります)
localhost の代わりにデプロイされたインスタンスを確認します。
HEALTH_URL=https://your-app.example.com/api/health npm 実行チェック
正常性エンドポイントが応答すること、キーが存在すること、モデルと
ベース URL が設定されています。 --live は、実際のラウンドトリップをさらに確認します。
終了コードは成功した場合は 0、失敗した場合は 1 なので、そのまま CI に落ちます。
ランタイム ゲートウェイ (ライブ アプリと統合)
サーバーは分析 UI を超えて、管理ゲートウェイとしても機能します。ポイントしてください
既存の SDK のbase_url とすべての呼び出しが自動的に管理されます。
制約が挿入され、max_tokens に制限があり、決定的なプロンプトが選択されます

イオン的に
ブロックされた後、すぐに戻ってきました。アプリのコードはほとんど変更されません。
ヘッダーによるリクエストごとの制御:
ポリシーは、server/policies.js (命令、 maxTokensCap 、
力温度、検出）。独自のものを編集または追加します。
# OpenAI SDK
openaiインポートからOpenAI
client = OpenAI (base_url = "http://localhost:8787/v1" 、api_key = "unused" )
クライアント。チャット 。完成品。作成(
モデル = "gpt-4o-mini" 、
ストリーム = True 、
メッセージ = [{ "ロール" : "ユーザー" , "コンテンツ" : "..." }],
extra_headers = { "x-governor-policy" : "簡潔" },
)
# Anthropic SDK
anthropic インポートより Anthropic
client = Anthropic (base_url = "http://localhost:8787" 、api_key = "unused" )
クライアントと。メッセージ。ストリーム (モデル = "claude-sonnet-5" 、 max_tokens = 1024 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : "..." }]) として s :
s 内のテキストの場合。テキストストリーム :
print ( text , end = "" )
(ゲートウェイはその環境から実際のキーを提供します。SDK の api_key は無視されます。)
stream: true は、サーバー送信イベントとして変更されずに渡されます。
背圧は維持されます。ゲートウェイはトークンの使用状況を記録するためだけにコピーを盗聴します。
クライアントが受け取るバイトを変更することはありません。試してみてください:
ノード例/stream.mjs # OpenAI プロトコルのストリーミング デモ
カール -N http://localhost:8787/v1/chat/completions \
-H ' コンテンツ タイプ: application/json ' -H ' x-governor-policy: 簡潔 ' \
-d ' {"モデル":"gpt-4o-mini","ストリーム":true,"メッセージ":[{"役割":"ユーザー","コンテンツ":"hi"}]} '
注意事項
ブロックはオプトインです。厳密なポリシー (または x-governor-detect: block ) のみが決定的なプロンプトをブロックします。提案されたスクリプトを含むプロバイダー形式の 400 を返します。検出器はヒューリスティックです。最初に warn で実行します。
max_tokens は唯一のハードレバーです。指示は勧告であり、上限は強制されます。
使用状況ログは、呼び出しごとに [使用状況] … 行を出力します。 OpenAI ストリーミングの場合は、 を設定します

tream_options.include_usage=true (デモではそうなっています)、またはストリームでは使用量がゼロとして読み取られます。
アップストリームでは同じプロトコル。アップストリームに一致するエンドポイントを使用します。ゲートウェイはプロトコル間で変換するのではなく、ネイティブに転送します (これが /api/messages の目的です)。
回復力。アップストリーム呼び出しには接続タイムアウトがあり、429/5xx でバックオフを使用して再試行します ( GATEWAY_TIMEOUT_MS 、 GATEWAY_RETRIES )。タイムアウトはストリーム本体ではなく接続/ヘッダーをカバーするため、長いストリームは切断されません。
マルチインスタンス。 REDIS_URL (および npm i redis ) を設定して、レプリカ間でバジェットとキャッシュを共有します。それ以外の場合、状態はプロセスごとです。
ゲートウェイを通過するジェミニ。 OpenAI アップストリームを Gemini の OpenAI-compat ベースに向けて /v1/chat/completions を呼び出します。 OPENAI_BASE_URL=https://generative language.googleapis.com/v1beta/openai (末尾のスラッシュなし) と OPENAI_API_KEY=<gemini key> を設定してから、gemini-* モデルを使用します。完全な互換性リストについては、プロバイダーの表を参照してください。
コントロール プレーン: 認証、バジェット、キャッシュ
ゲートウェイはコスト管理プレーンでもあります。 3 つはすべてオプションであり、重ねて使用できます。
同じ /v1/* エンドポイント (アプリ自体の /api/messages は影響を受けません)。
認証。呼び出し元をそれぞれベアラー キーとオプションの固定ポリシーで定義し、
予算、gateway.callers.json (gateway.callers.example.json を参照):
{ "発信者" : [
{ "id" : " web-app " 、 "key" : " gw_live_... " 、 "policy" : " terse " 、 "budget" : { "tokens" : 2000000 、 "usd" : 20 } },
{ "id" : " ワーカー " 、 "キー" : " env:WORKER_KEY " 、 "予算" : { "トークン" : 500000 } }
] }
呼び出し元は Authorization: Bearer <key> を送信します。 env:VAR のキー値が取得されます
環境から (つまり、注入されたシークレットから取得される可能性があります)。発信者がいない場合
設定されたゲートウェイはオープン (匿名) で実行されます - ローカルでは問題ありませんが、ローカルでは問題ありません
生産。呼び出し元の固定ポリシーには権威があるため、クライアントは緩めることができません

それ
x-governor-* ヘッダー付き。
予算。発信者ごとの毎日のトークンおよび/または USD の上限 (UTC リセット)。発信者がいるとき
以上で、次の呼び出しは Retry-After を使用して 429 を取得します。キャッシュヒットと現在
カウントは無料です。 USD の上限には価格設定ファイル (gateway.pricing.example.json →
GATEWAY_PRICING_FILE );トークンキャップは何も必要ありません。使用状況を確認してください:
カール -H ' 認証: ベアラー gw_live_... ' http://localhost:8787/v1/governor/usage
執行はソフトであり（各通話前にチェックされ、通話後に記録されます）、台帳は
メモリ内 — 単一インスタンス。複数のレプリカの場合は、server/budget.js をバックアップします。
共有ストア (Redis など)。
キャッシュ。非ストリーミングの確定的呼び出しのための完全一致キャッシュ
(温度 <= GATEWAY_CACHE_MAX_TEMP 、デフォルトは 0)。ヒットすると、保存されているものが返されます
x-governor-cache による応答: hit 、プロバイダーをスキップし、料金はかかりません
予算。サンプリングされた (高温の) 呼び出しとストリーミング呼び出しは決してキャッシュされないため、
古くなって非決定的な答えが得られることはありません。リクエストごとにバイパスする
x-governor-cache: オフ 。 GET /v1/governor/info でプレーン全体をイントロスペクトします。
npmテスト
依存関係のない (node --test) コントロール プレーンのカバレッジ — 認証トークン
チェック、予算執行、キャッシュの決定性/キーイング — なしで実行可能
プロバイダーキー。 npm run check / check:live runniを検証する

[切り捨てられた]

## Original Extract

Re-engineer LLM prompts to spend the fewest tokens that still answer. A governing gateway with auth, budgets, caching, and streaming. Works with Anthropic, OpenAI, Gemini, and any OpenAI-compatible provider. - shipitamit/llm-token-governor

GitHub - shipitamit/llm-token-governor: Re-engineer LLM prompts to spend the fewest tokens that still answer. A governing gateway with auth, budgets, caching, and streaming. Works with Anthropic, OpenAI, Gemini, and any OpenAI-compatible provider. · GitHub
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
shipitamit
/
llm-token-governor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits docs docs examples examples node_modules node_modules scripts scripts server server src src test test Dockerfile Dockerfile LICENSE LICENSE README.md README.md gateway.callers.example.json gateway.callers.example.json gateway.pricing.example.json gateway.pricing.example.json index.html index.html package-lock.json package-lock.json package.json package.json vite.config.js vite.config.js View all files Repository files navigation
Re-engineer LLM prompts to spend the fewest tokens that still answer — a
governing gateway with auth, budgets, caching, and streaming that works with
Anthropic, OpenAI, Gemini, and any OpenAI-compatible provider. Plus a live A/B,
saved constraint presets, and a batch triage that flags prompts which never
needed a model.
llm · tokens · cost-optimization · ai-gateway · llmops · openai · anthropic · gemini · streaming · prompt-engineering
Single mode — build a constrained request, see projected savings, and get a verdict when the wrapper costs more than it saves.
Presets — save/load constraint configs; they persist in the browser.
Batch mode — score a list of prompts, flag deterministic ones, export CSV/Markdown, and get a suggested script for each flagged row.
Live A/B — runs your prompt twice (raw vs. governed) through a server-side proxy and reports real token counts, using whichever provider you configure.
A fixed one-screen dashboard (no page scroll — panels scroll internally). The top
bar carries the mode tabs and pricing/scale inputs; below sit three panels:
Prompt & constraints — the prompt plus succinctness, output shape, temperature, and toggles.
Projected savings — a gauge, a "worth it / costs more than it saves" verdict, and a before/after cost receipt.
Assembled request & live A/B — the generated system prompt with copy buttons, and a raw-vs-governed A/B that reports the real token counts returned.
Switch to Batch mode and this becomes a triage table that flags deterministic
prompts. (The image is a layout mockup, not a screenshot — run the app for the
live, interactive version.)
Node.js 18+ (uses the built-in fetch )
An API key for one LLM provider (only needed for the live A/B panel)
npm install
cp .env.example .env # then edit .env (see "Providers" below)
Providers
The server is provider-agnostic. Set these in .env :
The frontend never changes: it POSTs a normalized request to /api/messages ,
and the adapter in server/providers.js translates to/from the chosen provider
and returns a normalized response. To add a native provider, edit that one file.
Almost every provider now offers an OpenAI-compatible API, so it works by setting
a base URL and model — no code change. Only providers you want to hit on their
own native protocol need an adapter.
All OpenAI-compatible providers use LLM_PROVIDER=openai :
# OpenAI
LLM_MODEL=gpt-4o-mini
# Gemini (Google AI Studio) — NOTE: no trailing slash (the code appends /chat/completions)
LLM_BASE_URL=https://generativelanguage.googleapis.com/v1beta/openai
LLM_MODEL=gemini-2.5-flash # use a model your key can access
# Groq
LLM_BASE_URL=https://api.groq.com/openai/v1
# OpenRouter
LLM_BASE_URL=https://openrouter.ai/api/v1
# Local Ollama
LLM_BASE_URL=http://localhost:11434/v1 LLM_API_KEY=ollama
Notes:
Gemini trailing slash. Google's own docs use .../v1beta/openai/ because the OpenAI SDK appends chat/completions . This server appends /chat/completions , so drop the trailing slash or you'll get a double-slash 404.
Some newer OpenAI models expect max_completion_tokens — set OPENAI_MAX_TOKENS_PARAM=max_completion_tokens if you hit that error.
Model availability changes and varies by key; if a model 404s, pick one from your provider's current list. npm run check:live confirms the wiring end to end.
Native Gemini SDK ( generateContent ) is not supported. If your app uses the Google GenAI SDK's native API rather than its OpenAI-compat mode, either switch that client to OpenAI-compat mode or add a native Gemini adapter (a body-mapper plus a :streamGenerateContent passthrough, parallel to the Anthropic adapter).
npm run dev
Vite frontend on http://localhost:5173
Express API on http://localhost:8787 (Vite proxies /api to it)
npm run build # outputs dist/
npm start # Express serves dist/ AND /api on PORT (default 8787)
Pre-deploy check
Confirm the key and model are wired before you ship:
npm run check # hits a running server's /api/health
npm run check:start # boots the server itself, checks, stops
npm run check:live # also sends a tiny real prompt (costs a few tokens)
Check a deployed instance instead of localhost:
HEALTH_URL=https://your-app.example.com/api/health npm run check
It verifies the health endpoint responds, the key is present, and a model and
base URL are configured; --live additionally confirms a real round-trip.
Exit code is 0 on success, 1 on any failure — so it drops straight into CI.
Runtime gateway (integrate with a live app)
Beyond the analysis UI, the server is also a governing gateway . Point your
existing SDK's base_url at it and every call is governed automatically —
constraints injected, max_tokens capped, deterministic prompts optionally
blocked — then streamed straight back. Your app code barely changes.
Per-request control via headers:
Policies live in server/policies.js (instructions, maxTokensCap ,
forceTemperature , detect ). Edit or add your own.
# OpenAI SDK
from openai import OpenAI
client = OpenAI ( base_url = "http://localhost:8787/v1" , api_key = "unused" )
client . chat . completions . create (
model = "gpt-4o-mini" ,
stream = True ,
messages = [{ "role" : "user" , "content" : "..." }],
extra_headers = { "x-governor-policy" : "terse" },
)
# Anthropic SDK
from anthropic import Anthropic
client = Anthropic ( base_url = "http://localhost:8787" , api_key = "unused" )
with client . messages . stream ( model = "claude-sonnet-5" , max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "..." }]) as s :
for text in s . text_stream :
print ( text , end = "" )
(The gateway supplies the real key from its env; the SDK's api_key is ignored.)
stream: true is passed through as Server-Sent Events, unmodified, with
backpressure preserved. The gateway sniffs a copy only to log token usage —
it never alters the bytes your client receives. Try it:
node examples/stream.mjs # OpenAI-protocol streaming demo
curl -N http://localhost:8787/v1/chat/completions \
-H ' content-type: application/json ' -H ' x-governor-policy: terse ' \
-d ' {"model":"gpt-4o-mini","stream":true,"messages":[{"role":"user","content":"hi"}]} '
Notes
Blocking is opt-in. Only the strict policy (or x-governor-detect: block ) blocks deterministic prompts; it returns a provider-shaped 400 with a suggested script. The detector is a heuristic — run it in warn first.
max_tokens is the only hard lever — instructions are advisory, the cap is enforced.
Usage logging prints a [usage] … line per call. For OpenAI streaming, set stream_options.include_usage=true (the demo does) or usage will read as zero on streams.
Same-protocol upstreams. Use the endpoint that matches your upstream; the gateway forwards natively rather than translating between protocols (that's what /api/messages is for).
Resilience. Upstream calls have a connect timeout and retry with backoff on 429/5xx ( GATEWAY_TIMEOUT_MS , GATEWAY_RETRIES ). The timeout covers connect/headers, not the stream body, so long streams aren't cut off.
Multi-instance. Set REDIS_URL (and npm i redis ) to share budgets and cache across replicas; otherwise state is per-process.
Gemini through the gateway. Point the OpenAI upstream at Gemini's OpenAI-compat base and call /v1/chat/completions : set OPENAI_BASE_URL=https://generativelanguage.googleapis.com/v1beta/openai (no trailing slash) and OPENAI_API_KEY=<gemini key> , then use a gemini-* model. See the Providers table for the full compatibility list.
Control plane: auth, budgets, cache
The gateway is also a cost-control plane. All three are optional and layer onto
the same /v1/* endpoints (the app's own /api/messages is unaffected).
Auth. Define callers, each with a bearer key and optional pinned policy and
budget, in gateway.callers.json (see gateway.callers.example.json ):
{ "callers" : [
{ "id" : " web-app " , "key" : " gw_live_... " , "policy" : " terse " , "budget" : { "tokens" : 2000000 , "usd" : 20 } },
{ "id" : " worker " , "key" : " env:WORKER_KEY " , "budget" : { "tokens" : 500000 } }
] }
Callers send Authorization: Bearer <key> . A key value of env:VAR is pulled
from the environment (so it can come from an injected secret). With no callers
configured the gateway runs open (anonymous) — fine for local, not for
production. A caller's pinned policy is authoritative: clients can't loosen it
with x-governor-* headers.
Budgets. Per-caller daily token and/or USD caps (UTC reset). When a caller is
over, the next call gets 429 with Retry-After . Cache hits and the current
count are free. USD caps need a pricing file ( gateway.pricing.example.json →
GATEWAY_PRICING_FILE ); token caps need nothing. Check usage:
curl -H ' Authorization: Bearer gw_live_... ' http://localhost:8787/v1/governor/usage
Enforcement is soft (checked before each call, recorded after), and the ledger is
in-memory — single-instance. For multiple replicas, back server/budget.js with
a shared store (e.g. Redis).
Cache. Exact-match cache for non-streaming, deterministic calls
( temperature <= GATEWAY_CACHE_MAX_TEMP , default 0). A hit returns the stored
response with x-governor-cache: hit , skips the provider, and doesn't charge
budget. Sampled (higher-temperature) and streaming calls are never cached, so you
never get a stale non-deterministic answer. Bypass per request with
x-governor-cache: off . Introspect the whole plane at GET /v1/governor/info .
npm test
Dependency-free ( node --test ) coverage of the control plane — auth token
checks, budget enforcement, and cache determinism/keying — runnable without a
provider key. npm run check / check:live validate a runni

[truncated]
