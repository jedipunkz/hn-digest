---
source: "https://blog.spill.coffee/p/trust-the-harbor-pilot"
hn_url: "https://news.ycombinator.com/item?id=48861719"
title: "Lessons from designing MCP tools for AI agents"
article_title: "Trust the harbor pilot - by Gene Beal"
author: "spillcoffee"
captured_at: "2026-07-10T16:14:47Z"
capture_tool: "hn-digest"
hn_id: 48861719
score: 1
comments: 0
posted_at: "2026-07-10T16:00:06Z"
tags:
  - hacker-news
  - translated
---

# Lessons from designing MCP tools for AI agents

- HN: [48861719](https://news.ycombinator.com/item?id=48861719)
- Source: [blog.spill.coffee](https://blog.spill.coffee/p/trust-the-harbor-pilot)
- Score: 1
- Comments: 0
- Posted: 2026-07-10T16:00:06Z

## Translation

タイトル: AI エージェント用の MCP ツールの設計から得た教訓
記事のタイトル: 港のパイロットを信頼する - ジーン・ビール著
説明: AI エージェント用の MCP ツールの設計から得た教訓

記事本文:
港のパイロットを信頼してください - ジーン・ビール著
購読 サインイン 港のパイロットを信頼してください
AI エージェント用の MCP ツールの設計から得た教訓
Gene Beal 2026 年 7 月 9 日 共有モデル コンテキスト プロトコル (MCP) は、達成するのが十分に簡単なタスクのように思えます。 UI が使用するアプリケーションの API がすでにあり、コードの再利用を通じて関連する部分をリサイクルし、AI が消化できるように適切にラップします。 AI ダイジェスト MCP 実装は、独自のツールを使用するのではなく、ツールを利用することだけを知っておく必要があります。ユーザーがインストールしたんですよね？
新しい読者向けの簡単なコンテキスト: FlurryPORT は、安定した URL で受信 Webhook をキャプチャし、署名をそのままにしてローカル アプリに再生します。私は AI を使用した設計セッションに着手し、実装すべき主要な領域を考え出しました。セキュリティ上の懸念について話し合い、それを計画に組み込みました。 FlurryPORT.dev/try にある無料サービスからのアップグレード パスを精査するのに午後丸々費やしたこともありました。数回のコーディング セッションの後、私の製品には効果的な MCP stdio が実装されました。そう思っていました。
私は、コーディング補助として使用したものに対してライバルのコーディング AI を起動し、プロンプトを出し始めました。私は、プロンプトをよく考えていないものの、タスクの進め方については大まかに知っているユーザーを模倣するために、誘導的ではない質問から始めました。 AI は、私の実装で提供されたものを除く、Webhook を送受信するために考えられるすべてのツールを把握していることがすぐにわかりました。
私はテストを中止して、「この MCP 実装はどうでしたか、何が改善されるでしょうか?」と直接尋ねることにしました。要件を集めるのがこんなに簡単だったことに驚きましたし、いくつかのことを学びました。次に、次のバージョンの設計をテスト AI エージェントに渡し、ビルドする前にレビューしてもらいました。
投票をやりがいのあるものにする — 表現力のない当たり障りのないデータは AI にとって魅力的ではありません。これ

AI は単なるコンピューターではなく、生の数値からあらゆる種類の有用性を引き出すべきではないという点で、私は直感に反していると感じました。これは直観に反すると私は思いました。AI は単なるコンピューターではなく、生の数値が与えられればあらゆる種類の有用性を導き出せるはずではないでしょうか。しかし、このモデルは計算しているのではなく、最も確実にタスクを前進させるものに手を伸ばしており、必要最低限​​の数字だけでは何も行動することができません。違いはペイロードを見ると簡単に分かります。
{
「キャプチャカウント」: 12、
"キャプチャ残り": 88、
"バースト": { "毎分": 30, "使用済みこの分": 22 },
"notices": ["セッション上限まで 8 キャプチャ"],
「アクション」: [
{ "label": "キャプチャを保持するにはこのセッションを要求します"、"cost": "free"、"effect": "キャップを削除します"、"url": "https://flurryport.io/claim/..." }
】
}
アフォーダンスは散文に勝る — AI にサーバーへのリクエストのペースを調整する必要があると通知する段落は無視されました。シンプルな Capture_count ツール (その説明にその目的が記載されています) が、バッチごとに呼び出されます。もしあなたがガイダンスの散文を書いていることに気づいたら、そのガイダンスを不要にするツールは何かを尋ねてください。
構造化された事実を返す — ラベル/コスト/効果/URL を含む領収書、診断コード、および action[] が、実行のたびに逐語的にユーザーに中継されました。うまく練られた散文による定型的な説明が、モデルの役割を果たそうとする私たちのツールでした。モデルにユーザーにナレーションを与え、理解できる事実を与えます。
洗練されたコピーよりも安定したコード — エージェントは、文章の表現ではなく、 at_cap または cap_would_exceed をキーにします。機械可読な状態により、散文によって AI が即興で進行方法を決定する動作がより決定的になりました。エラー文字列は UI です。コードはAPIです。
行動する前に拒否する — 部分的に成功すると、審査担当者が「説明責任」と呼ぶものが生じます。 10 件のイベントのうち 7 件を送信すると、エージェントはユーザーにナレーションを行う必要があります。きれいな拒否

at では、「最大 N を送信」と指定すると、即座にクリーンな動作が生成されます。リクエストを処理する前に、利用可能な予算と照らし合わせてリクエスト全体を確認してください。
移行にはマップが必要 — ツールが変更された場合、または変換イベント (FlurryPORT の匿名からアカウントへの反転など) に基づいてデータが移行された場合、その変更マッピングを AI エージェントに提供します。ブレッドクラムがない場合、AI エージェントは 4 分間ツールを探索して、何が変更されたかを再発見しました。マッピングが提供された場合、同じタスクが 21 秒で実行されました。これに続くものとして、移行後も名前を維持するツールは互換性のあるスキーマも維持する必要があります。 AI クライアントはスキーマを不完全にキャッシュします。
ワンショット通知が失われる - session_claimed 通知がバッチ途中のポーリングによって失われ、エージェントはアップグレードが成功したことをユーザーに通知しませんでした。エージェントが見逃してはならないものは、単一のイベントではなく、すべての応答で永続的なフィールドである必要があります。
デーモンの実行をエージェントに依存しない - エージェントは、ユーザーのためにレシーバーを起動するためにシェル プロセスのバックグラウンド処理に 1 分 46 秒を費やしました。これを実装内で利用可能なツールに置き換えることで、その時間が 3 秒に短縮されました (例: start_echo_server )。
ツールの説明はルーティング層です。説明内の「CALL THIS FIRST」は、ngrok または Stripe CLI で 4 回のコールド ランが到達しなかった理由です。見つけやすさはドキュメントではなく説明に書かれています。
エージェントに質問する — エージェントに質問することは、ビルドの最良のレビュー担当者です。ビルド前の設計レビューではスキーマが再設計され、ペーシング セマンティクスとカット機能が追加されました。事後フィードバックによりバグが発見されます。ビルド前のレビューにより、悪い形状を防ぎます。
MCP 実装の最適化に時間をかけることは、有意義な時間を費やします。
80 イベントの一括実行中に拒否された送信 — 最適化前: ~50 (429 秒)、最適化後: 0
完全な匿名ファネル、壁時計 — 最適化前

乾燥: ~20 分、後: ~7 分
サインアップ後の再オリエンテーション — 最適化前: ~4 分、最適化後: 21 秒
ローカル エコー レシーバーのセットアップ — 最適化前: デーモンのシェル実行 1 分 46 秒、後: 3 秒、1 回のツール呼び出し
手動シェル コマンド — 最適化前: 多数、最適化後: 0
そして、審査中の AI の評決は、そのままの意味で「私は最初に ngrok や Stripe CLI に手を伸ばしたわけではありません。私はキャプチャして、見て、再生して、説明しました。それがまさにあなたが望んでいる製品の動きです。」
エージェントにそれを指摘する前に、細かい点を確認してください。サインアップなしセッションはプレーンテキストで有効期限が切れます (約 90 分、キャプチャ 100 件)。ツールはエージェントに本番データや PII データを送信しないよう指示します。セッションによってすべてが暗号化されたプロジェクトに移動されると主張します。テスト イベント シグネチャは現実的な形状であり、実際のシグネチャではありません。ツールはそのように表示します。現時点では stdio のみであるため、Claude Code、Codex CLI、および Cursor は動作します。 ChatGPT と claude.ai Web にはまだありません。エージェントが取得するトークンは読み取り専用で、デフォルトで PII が編集され、サインアップ リンクはエージェントではなくブラウザに送信されます。
ご親切にあなたの経験についてコメントしていただければ、ぜひ聞かせていただきたいと思います。
クロード mcp add flurryport -- npx -y flurryport mcp codex mcp add flurryport -- npx -y flurryport mcp [mcp_servers.flurryport]
コマンド = "npx"
args = ["-y", "flurryport", "mcp"] Webhook ハンドラーをデバッグし、どのツールが選択されるかを監視するように要求します。ドキュメントは flurryport.io/docs/cli にあります。
シェア 前 この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

Lessons from designing MCP tools for AI agents

Trust the harbor pilot - by Gene Beal
Subscribe Sign in Trust the harbor pilot
Lessons from designing MCP tools for AI agents
Gene Beal Jul 09, 2026 Share Model Context Protocol (MCP) seems like a simple enough task to achieve. You already have an API for your application that your UI uses, recycle relevant pieces through code re-use and wrap them nicely for AI to digest. An AI digesting MCP implementation should just know to reach for your tool over rolling their own; their user installed it didn’t they?
Quick context for new readers: FlurryPORT captures your incoming webhooks at a stable URL and replays them into your local app, signatures intact. I set out in a design session with AI to come up with the key areas that should be implemented. Discussed security concerns and wrapped them into the plan. I even spent an entire afternoon scrutinizing an upgrade path from my free offering at FlurryPORT.dev/try . A few coding sessions later and my product had an effective MCP stdio implementation — or so I thought.
I fired up a rival coding AI to what I used as a coding aide and started prompting. I started with non-leading questions to mimic a user who hadn’t thought through their prompts, but loosely knew about how to go about their task. I quickly found that AI was grasping at every tool it could imagine to send/receive webhooks except those provided by my implementation.
I decided to stop my test and just ask directly: how was this MCP implementation, what would make it better? I was surprised how easy it was to gather requirements and I learned a few things. I then handed the next version of design back to the test AI agent to review before building it.
Make polling rewarding — Having bland data that isn’t expressive is not appealing to AI. This I found counter intuitive, in that isn’t AI just a computer and shouldn’t it just derive all sorts of usefulness give raw numbers. This I found counterintuitive: isn't AI just a computer, and shouldn't it derive all sorts of usefulness given raw numbers? But the model isn't calculating, it's reaching for whatever most reliably moves the task forward, and a bare number gives it nothing to act on. The difference is easy to see in the payload.
{
"captureCount": 12,
"capturesRemaining": 88,
"burst": { "perMinute": 30, "usedThisMinute": 22 },
"notices": ["8 captures until session cap"],
"actions": [
{ "label": "Claim this session to keep captures", "cost": "free", "effect": "removes cap", "url": "https://flurryport.io/claim/..." }
]
}
Affordances beat prose — A paragraph informing AI that it should pace requests to the server was ignored. A simple capture_count tool, whose description says what it's for, got called between every batch. If you find yourself writing guidance prose, ask what tool would make the guidance unnecessary.
Return structured facts — Receipts, diagnosis codes and an action[] with label/cost/effect/url got relayed to the user verbatim on every run. A canned explanation with well crafted prose was our tool trying to do the model’s job. Let the model narrate for its user and give it understandable facts.
Stable codes over polished copy — An agent will key on at_cap or cap_would_exceed , not sentence wording. Having machine readable state made behavior more deterministic where prose set AI off improvising on how to proceed. Error strings are UI; codes are API.
Refuse before acting — A partial success creates what the reviewing agent called “explanation debt”. Sending 7 of 10 events forces an agent to narrate to its user. A clean refusal that says “send at most N” produced clean behavior instantly. Check the whole request against available budget before processing any of it.
Transitions need a map — If tooling changed or data was migrated based on a conversion event (like FlurryPORT’s anonymous-to-account flip) provide that change mapping to the AI agent. Without a breadcrumb the AI agent spent 4 minutes exploring tools to rediscover what just changed. When mapping was provided it performed the same task in 21 seconds. As a follow up to this one, tools that keep their names across a transition should keep compatible schemas too; AI clients cache schemas imperfectly.
One shot notifications get lost — A session_claimed notice was lost by a mid-batch poll and the agent never informed the user their upgrade was successful. Anything that an agent must not miss must be a durable field on every response, not on a single event.
Don't rely on agents to run daemons — An agent spent 1m 46s backgrounding a shell process to stand up a receiver for its user. Replacing it with tooling available within the implementation brought that time down to 3s (eg. start_echo_server ).
Tool descriptions are your routing layer — “CALL THIS FIRST” in a description is a reason four cold runs never reached for ngrok or the Stripe CLI. Discoverability is written into descriptions not the docs.
Ask the agent — Asking an agent is the best reviewer of your build. The pre-build design review redesigned schemas, added pacing semantics and cut features. P ost-hoc feedback finds bugs; pre-build review prevents bad shapes.
Taking time to optimize your MCP implementation is time well spent.
Rejected sends during an 80-event bulk run — Before optimizing: ~50 (429s), After: 0
Full anonymous funnel, wall clock — Before optimizing: ~20 min, After: ~7 min
Post-signup re-orientation — Before optimizing: ~4 min, After: 21 seconds
Local echo receiver setup — Before optimizing: 1m 46s of shelling a daemon, After: 3 seconds, one tool call
Hand-rolled shell commands — Before optimizing: many, After: zero
And the reviewing AI’s verdict, verbatim: "I did not reach for ngrok or Stripe CLI first. I captured, watched, replayed, and explained. That's exactly the product motion you want."
Before you point an agent at it, the fine print. The no-signup session is plaintext and expires (about 90 minutes, 100 captures), and the tools tell the agent not to send production or PII data there; claiming the session moves everything into an encrypted project. Test-event signatures are realistic shapes, not real signatures, and the tool says so. And it's stdio only for now, so Claude Code, Codex CLI, and Cursor work; ChatGPT and claude.ai web don't yet. The token your agent gets is read-only and redacts PII by default, and signup links go to your browser, never the agent.
If you would be kind enough to comment on your experience I would love to hear it.
claude mcp add flurryport -- npx -y flurryport mcp codex mcp add flurryport -- npx -y flurryport mcp [mcp_servers.flurryport]
command = "npx"
args = ["-y", "flurryport", "mcp"] Ask it to debug a webhook handler and watch which tools it picks. Docs at flurryport.io/docs/cli .
Share Previous Discussion about this post Comments Restacks Top Latest Discussions No posts
