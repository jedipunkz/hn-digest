---
source: "https://epho.io"
hn_url: "https://news.ycombinator.com/item?id=49376256"
title: "Show HN: Epho – run Claude Code with a curl"
article_title: "epho — agents as an api"
image: ""
author: "karakanb"
captured_at: "2026-08-20T16:23:13Z"
capture_tool: "hn-digest"
hn_id: 49376256
score: 4
comments: 0
posted_at: "2026-08-20T15:45:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Epho – run Claude Code with a curl

- HN: [49376256](https://news.ycombinator.com/item?id=49376256)
- Source: [epho.io](https://epho.io)
- Score: 4
- Comments: 0
- Posted: 2026-08-20T15:45:51Z

## Translation

タイトル: HN を表示: Epho – カールを使用してクロード コードを実行
記事のタイトル: epho — API としてのエージェント
説明: epho はクラウドでクロード コード、コーデックス、またはオープンコードを実行します。メッセージを投稿し、作業をストリームバックします。 SDK、デーモン、インフラはありません。
HN テキスト: 皆さん、Burak です。 Epho は、クラウドのサンドボックスで Claude Code、Codex、または Opencode を実行できるようにする API です。サンドボックスを抽象化し、単一の HTTP リクエストでコーディング エージェントを実行できるようにします。 Epho は、独自の AI アナリストを構築するという私たち自身の苦労から生まれました。
- サンドボックスでは、裸のマシンが提供されます。エージェントのワークロードに合わせて構成する必要があります。
- 各エージェントの動作は異なるため、それぞれとの統合を構築する必要があります。
- サンドボックス プロバイダーはあまり信頼できないため、障害を避けるためにマルチプロバイダー戦略を立てる必要があります。
- ロギング、アーティファクト、入出力、イベント ストリーミング、その他すべての運用面を把握する必要があります。私たちは自らその痛みを経験しなければなりませんでした。私たちは物事がかなり信頼できるところまで到達し、これはそれ自体がプリミティブであるべきであることがより明白になりました。つまり、POST リクエストを送信し、イベントをストリーミングして返します。 Epho は API 製品としてのエージェントです。リクエストを送信すると、サンドボックスが起動され、選択されたハーネスが構成され、リポジトリのクローンが作成され、エージェントが起動されます。さまざまなプロバイダー間での自動フォールバックを処理し、認証関連を処理し、イベントと出力をストリーミングバックするだけです。 Claude Code、Codex、Opencode をそのままサポートしており、それらがサポートするほぼすべてのモデルをそのままサポートしています。イベントをストリーミングバックし、添付ファイルと出力ファイルを処理し、さまざまなサンドボックスプロバイダーでのフォールバック、再試行、およびすべての認証を自動的に管理します。プロンプト、リポジトリ、必要な MCP サーバーを送信するだけです。

o それを使用すると、それらを実行します。実際の例を示すために、ここでデモを録画しました: https://youtu.be/HGfly1aytPA 私は Epho に非常に興奮しています。なぜなら、Epho は、今日よりもはるかに簡単にエージェントを製品に組み込むことができる新しいプリミティブだと思うからです。私たちは本番環境の Epho でエージェントを実行しているため、関係なくメンテナンスを継続し、独立した製品として出荷したいと考えていました。 Epho は無料で始めることができ、Opencode の無料モデルで実行して使い始めることができます。皆様のご意見を知りたいと思っておりますので、フィードバックをお待ちしております。乾杯、
ブラク

記事本文:
epho 仕組み 機能 ドキュメント 料金 よくある質問 主要なエージェントを入手
API 。
epho はクラウドで claude code 、 codex 、または opencode をスピンアップします。つまり、リポジトリが複製され、ファイルが添付され、mcp サーバーが接続されます。プロンプトを POST すると、作業が戻ってきます。 SDK、デーモン、インフラはありません。
プロバイダーキーが手元にない場合は、 opencode/…-free モデルは、オープンコードなしで実行されます。
POST /api/v1/chat
{
"ハーネス" : "クロード" ,
"モデル" : "クロード-ソネット-5" ,
"prompt" : "auth_test.go の不安定なテストを修正します。" 、
"リポジトリ" : [{"url": "…/acme/api"}] ,
"provider_api_key" : "sk-ant-••••••••”
} クロード code anthropic のエージェント codex openai のエージェント opencode オープンソース エージェントで動作 $ epho Explain --how-it-works
あなたは http を話します。
エージェントはリポジトリと対話します。
1 つのリクエストで、選択したハーネスを使用して新しいクラウド環境を起動し、そこにリポジトリを複製し、ツールを接続し、プロンプトを渡し、発生したときにすべてをパイプで戻します。
┌───────┐
│あなた│
└───────┘ スクリプト POST /api/v1/chat ┌────┐
│ エフォ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┐
│ クロードコード │
§─────────┤
│コーデックス│
§─────────┤
│ オープンコード │
└─────────┘ クラウド ボックス内のテキスト/イベント ストリーム — トークン、ツール呼び出し、差分 01
1 つのエンドポイント、1 つの JSON 本文。ハーネスとモデルを選択し、プロンプトを作成し、provider_api_key を渡します。
ハーネスがインストールされた状態でサンドボックスが起動し、リポジトリが複製され、input_files がステージングされ、mcp_servers が接続されます。
sse は、ツールの呼び出し、編集、最終的な答えなど、実行全体をストリームバックします。または /chat/async に POST して、後で結果を取得します。
チャットは退屈だ

サンドボックスを接続した会話。ボックスはターンが終了すると停止し、次のターン (同じファイルシステム、同じチェックアウト、同じエージェントセッション) で再開します。 chat_id を返して続行してください。
エージェントはあくまで
文脈としては良い。
プロンプト自体はパーティのトリックです。コード、ファイル、ツール、シークレットはすべて同じリクエストに乗っています。これは、それらを運ぶすべてのフィールドです。
小さな表面。
深刻な走行距離。
製品全体が 1 つのエンドポイントであり、デフォルトが適切です。以下のものはすべてそこから外れます。
それがAPIです。 json を記述できれば、ドキュメントは完成です。
ハーネスを 1 つのフィールドと交換します。同じリクエストの形状でも、頭脳は異なります。
プロバイダー キーはリクエストとともに送られます。トークン、プロバイダーの請求書 — epho がユーザーとモデルの間に介在することはありません。
-N を渡して監視します。最初に ID を関連付けて、次にすべてのエージェント イベントを調べてから、出力とアーティファクトを処理します。
エージェントは私たちのボックス上で実行されます。蓋を閉めて、飛行機に乗りましょう。入力は続きます。
それは単なる http です。それを Github アクション、cron、Discord ボット、5 行の bash にパイプします。
ボディは同じですが、エンドポイントは異なります。 /chat/async はミリ秒単位で ID を返し、そのまま続行します。webhook_url を設定すれば、それは完全に忘れられます。
-free で終わる opencode zen モデルは、プロバイダーの資格情報なしで実行されます。 1 つの epho キーとカールがセットアップ全体です。
パイプが落ちても、実行が失われたわけではありません
すべてのイベントは永続化されます。再接続すると、ライブテールが再開される前に見逃したものを取得できます。または、十分に確認したらターンを完全に停止します。
サンドボックスはあなたに箱を渡します。
epho はあなたにエージェントを渡します。
e2b、daytona、レンタルした VM など、未加工のサンドボックス上でクロード コードを自分で実行することもできます。人々はそうします。サインアップするものは次のとおりです。
完全なマシン制御が必要ですか?サンドボックスを使用します。仕事が必要ですか？エフォを使います。
実行するサンドボックスのリソースに対して料金を支払います

セス。メーターは起動から分解まで動作し、動作が完了すると停止します。何もアイドル状態にならず、何も保存されず、何も請求を継続しません。
デフォルトのインスタンス: 2 vCPU · 2 GiB · 10 GiB ≈ $0.0000438/s · 1 分間の実行 ≈ $0.0026
開始クレジット: サインアップ時に 10 ドル — デフォルト インスタンスの約 60 時間
モデル トークン: epho によって請求されることはありません。これらは、プロバイダーの価格で、provider_api_key に基づいて利用されます。
メーター: GET /credits でいつでも読み取ります。残高がゼロになると、ターンは 402 で拒否されます。
これが全体です
統合。
それをコピーし、キーを交換し、プロンプトを変更します。これで、コーディング エージェントをクラウドで実行できるようになりました。これ以降はすべてオプションです。
これらの 1 つを独自のコーディング エージェントに渡します。それぞれは上記のエンドポイント上に構築された実際のアプリで、午後には完了できるほど小さいものです。
リポジトリのオープンなプル リクエストをリストします。レビュー ボタンは、epho エージェントによる差分の読み取りをサイド パネルにストリーミングします。差分は入力ファイルとして渡され、ベース ブランチでリポジトリが複製されます。レビューは同時に実行されます。
エージェントとラベル付けされた Github の問題をプル リクエストに変換します。 POST /chat/async で実行を開始し、Webhook が到達したときに問題について結果をコメントバックします。
CSV をブラウザにドロップし、入力ファイルとして epho エージェントに渡します。アーティファクトとして書き戻されるチャートをレンダリングし、ストリーミング中にエージェントの推論を表示します。
スレッドごとに 1 つの epho チャットを備えた Slack ボット。エージェントの作業を発生時にスレッドにストリーミングし、chat_id を再利用して、フォローアップが同じ温かいサンドボックスに入るようにします。
すべてのコピーには、https://epho.io/docs/api へのポインタと、キーをサーバー側に保管するためのリマインダーが同梱されています。そのため、エージェントは何かを書き込む前に仕様を読み取ります。
必要ありません。単一の POST です。本当に型付きヘルパーが必要な場合は、お気に入りの llm にカールを貼り付けて尋ねてください。なんとかなりますよ。

あなたがやる。リクエストには Provider_api_key が含まれるため、トークンはプロバイダーの価格でプロバイダーの請求書に記載されます。マークアップも仲介者もいません。
サンドボックス時間は、ブートからティアダウンまで計測されます。インスタンス キーを使用してボックスのサイズを自分で設定します。すべての実行では、最終イベントで独自のcost_usdがレポートされます。
ハーネスがマシン上で行うことはすべて、リポジトリの読み取り、コードの作成、テストの実行、失敗時の反復などです。それをクラウドでオンデマンドで並行して実行するだけです。
はい — 各チャットには独自の隔離された環境があるため、好きなだけ広く展開できます。チームは 5 つのアクティブなターンから開始します。上限を超えたものはすべてキューに置かれ、キュー済みステータスが発行され、スロットが開くまでコストはかかりません。ただし、チャットごとに 1 つのライブ ターンが必要です。並行作業には並行チャットが必要です。
それがチャットというものです。新しいプロンプトで chat_id を返すと、同じファイルシステム、同じチェックアウト、同じエージェント セッションでサンドボックスが再開されます。基礎となるボックスがなくなった場合、epho は代替ボックスを起動し、セッションのスナップショットを復元するため、会話はマシン上で存続します。
repos 内の最大 32 のエントリを渡します。各エントリにはオプションのブランチとリポジトリごとのトークンが含まれます。 github、gitlab、bitbucket のトークンはすべて処理されます。リポジトリはターンごとであるため、会話の途中でトークンをローテーションしたり、新しいリポジトリを追加したりできます。
はい — mcp_servers は、ローカル stdio プロセス ( command + args ) またはリモート サーバー ( URL + headers ) のいずれかを受け取ります。リニア、figma、独自の内部サーバー。これらはチャットの作成時に設定され、その後のターンごとに継承されます。
いいえ。 /chat/async に POST すると、すぐに ID を含む 202 が得られ、その後 /chat/{id}/response から結果を収集するか、 webhook_url を設定します。ストリーミング エンドポイント上でもパイプがドロップされても実行が強制終了されることはありません。/events/subscribe に再接続すると、見逃したものが再生されます。
両方。 input_files はサンドボックスに最大 20 個のファイルを置きます — inl

Base64 またはフェッチした URL を取得し、そのパスをプロンプトに追加します。エージェントが生成するものはすべて、署名付きのダウンロード URL を持つアーティファクトとして返されます。有効期限は約 10 分間で、応答エンドポイントから再作成できます。
はい — -free で終わる opencode zen モデルは、-free なしで実行されます。 "harness": "opencode" と -free モデルを選択し、provider_api_key を完全に削除すると、epho キーがリクエスト内の唯一の認証情報になります。
それはその実行に使用され、それ以外には何も使用されません。再開を含むすべてのターンで再送信すると、ターンが終了した瞬間にサンドボックスが停止します。
ここで取り上げられていないことはありますか？ [メールで保護されています]
__
/\ \
__ _____\ \ \___ ___
/'__`\/\ '__`\ \ _ `\ / __`\
/\ __/\ \ \L\ \ \ \ \ \/\ \L\ \
\ \____\\ \ ,__/\ \_\ \_\ \____/
\/___/ \ \ \/ \/_/\/_/\/___/
\ \_\
\/_/ 鍵を手に入れてください。カールを保ちます。
今から 2 分後には、エージェントの作業を端末から見ているかもしれません。

## Original Extract

epho runs claude code, codex, or opencode in the cloud. POST a message, stream back the work. no sdk, no daemon, no infra on your side.

Hey folks, Burak here. Epho is an API that allows running Claude Code, Codex or Opencode in a sandbox in the cloud. It abstracts away sandboxes, and allows running coding agents with a single HTTP request. Epho came out of our own struggles with building our own AI analyst:
- Sandboxes give you bare machines; you need to configure them for agentic workloads.
- Each agent behaves differently, and you need to build integrations with each of them.
- Sandbox providers are not very reliable, which means you need to figure out a multi-provider strategy to avoid failures.
- Logging, artifacts, input/output, event streaming, and all of the other operational aspects need to be figured out. We had to go through the pain ourselves. We got to a point where things got quite reliable, and it became more obvious to us that this should be a primitive on its own: send a POST request, get the events streaming back to you. Epho is an agents-as-an-API product: you send a request, it spins up a sandbox, configures the chosen harness, clones your repos, and kicks off the agent. It takes care of automatic fallbacks across different providers, handles auth stuff, and just streams back the events and outputs. It supports Claude Code, Codex and Opencode out of the box, and pretty much all the models they support out of the box. It streams the events back, handles attachments and output files, automatically manages the fallbacks on different sandbox providers, retries, and all the auth stuff. You just send a prompt, your repo, MCP servers you want to use with it, and it runs them. I recorded a demo here to show a real example: https://youtu.be/HGfly1aytPA I am quite excited for Epho, simply because I think it is a new primitive that would allow building agents into product a lot easier than it is today. We are running our agents on Epho on prod, so we'll keep maintaining it regardless, and we wanted to ship it as an independent product. Epho is free to get started, and you can run it with Opencode's free models to get started with it. I am quite curious to hear what you'd think and would love to get your feedback. Cheers,
Burak

epho how it works features docs pricing faq get a key agents as
an api .
epho spins up claude code , codex , or opencode in the cloud — your repos cloned, your files attached, your mcp servers wired in. you POST a prompt, the work streams back. no sdk, no daemon, no infra on your side.
no provider key handy? the opencode/…-free models run without one.
POST /api/v1/chat
{
"harness" : "claude" ,
"model" : "claude-sonnet-5" ,
"prompt" : "Fix the flaky test in auth_test.go." ,
"repos" : [{"url": "…/acme/api"}] ,
"provider_api_key" : "sk-ant-••••••••"
} WORKS WITH claude code anthropic's agent codex openai's agent opencode the open-source agent $ epho explain --how-it-works
you talk http.
the agent talks to your repo.
one request boots a fresh cloud environment with the harness you picked, clones your repos into it, connects your tools, hands it your prompt, and pipes everything back as it happens.
┌───────┐
│ you │
└───────┘ your script POST /api/v1/chat ┌────────┐
│ epho │
└────────┘ boots & routes runs ┌─────────────┐
│ claude code │
├─────────────┤
│ codex │
├─────────────┤
│ opencode │
└─────────────┘ in a cloud box text/event-stream — tokens, tool calls, diffs 01
one endpoint, one json body. pick a harness and model , write a prompt , pass your provider_api_key .
a sandbox boots with the harness installed, your repos cloned, your input_files staged and your mcp_servers connected.
sse streams back the whole run — tool calls, edits, the final answer. or POST to /chat/async and pick the result up later.
a chat is a durable conversation with a sandbox attached. the box stops when a turn ends and resumes for the next one — same filesystem, same checkout, same agent session. hand back the chat_id and keep going.
an agent is only as
good as its context.
a prompt on its own is a party trick. the code, the files, the tools and the secrets all ride on the same request — this is every field that carries them.
small surface.
serious mileage.
the whole product is one endpoint and good defaults. everything below falls out of that.
that's the api. if you can write json, you've finished the docs.
swap harnesses with one field. same request shape, different brain.
your provider key rides along with the request. your tokens, your provider bill — epho never sits between you and the model.
pass -N and watch. ids first so you can correlate, then every agent event, then done with the output and artifacts.
agents run on our boxes. close your lid, board your flight — it keeps typing.
it's just http. pipe it into github actions, cron, a discord bot, five lines of bash.
same body, different endpoint. /chat/async hands back ids in milliseconds and gets on with it — set a webhook_url and forget about it entirely.
the opencode zen models ending in -free run without any provider credential. one epho key and a curl is the entire setup.
a dropped pipe isn't a lost run
every event is persisted. reconnect and you get what you missed before the live tail resumes — or stop a turn outright when you've seen enough.
a sandbox hands you a box.
epho hands you an agent.
you can absolutely run claude code yourself on a raw sandbox — e2b, daytona, a vm you rent. people do. here's what you sign up for.
need full machine control? use a sandbox. need work done? use epho.
you pay for the sandbox resources your run uses. the meter runs from boot to teardown and stops when the run does. nothing idles, nothing is stored, nothing keeps billing.
default instance: 2 vCPU · 2 GiB · 10 GiB ≈ $0.0000438/s · a one-minute run ≈ $0.0026
starting credit: $10 on signup — roughly 60 hours of the default instance
model tokens: never billed by epho. they ride on your provider_api_key, at your provider's price.
the meter: read it any time with GET /credits — a turn is refused with 402 once the balance hits zero.
this is the entire
integration.
copy it, swap in your key, change the prompt. you're now running coding agents in the cloud. everything after this is optional.
hand one of these to your own coding agent. each is a real app built on the endpoint above, small enough to finish in an afternoon.
List a repo’s open pull requests. A review button streams an epho agent’s read of the diff into a side panel — diff passed in as an input file, repo cloned at the base branch. Reviews run concurrently.
Turn any github issue labelled agent into a pull request. Fire the run with POST /chat/async and comment the result back on the issue when the webhook lands.
Drop a csv in the browser and hand it to an epho agent as an input file. Render whatever chart it writes back as an artifact, and show the agent’s reasoning as it streams.
A slack bot with one epho chat per thread. Stream the agent’s work into the thread as it happens, and reuse the chat_id so follow-ups land in the same warm sandbox.
every copy ships with a pointer to https://epho.io/docs/api and a reminder to keep your key server-side — so your agent reads the spec before it writes anything.
you don't need one — it's a single POST. if you really want typed helpers, paste the curl into your favorite llm and ask. it'll manage.
you do. the request carries your provider_api_key , so tokens land on your provider bill at your provider's price. no markup, no middleman.
sandbox time is metered from boot to teardown. size the box yourself with the instance key. every run reports its own cost_usd in the final event.
whatever the harness does on your machine: read repos, write code, run tests, iterate on failures. it just does it in the cloud, on demand, in parallel.
yes — each chat gets its own isolated environment, so fan out as wide as you like. teams start with five active turns; anything past the cap sits in a queue, emits a queued status and costs nothing until a slot opens. one live turn per chat, though — parallel work wants parallel chats.
that's what a chat is. pass the chat_id back with a new prompt and the sandbox resumes — same filesystem, same checkout, same agent session. if the underlying box is gone, epho boots a replacement and restores the session snapshot, so the conversation survives the machine.
pass up to 32 entries in repos , each with an optional branch and a per-repo token . github, gitlab and bitbucket tokens are all handled. repos are per-turn, so you can rotate a token or add a new repo halfway through a conversation.
yes — mcp_servers takes either a local stdio process ( command + args ) or a remote server ( url + headers ). linear, figma, your own internal server. they're set when the chat is created and every later turn inherits them.
no. POST to /chat/async and you get 202 with the ids immediately, then collect the result from /chat/{id}/response or set a webhook_url . even on the streaming endpoint a dropped pipe doesn't kill the run — reconnect to /events/subscribe and it replays what you missed.
both. input_files puts up to 20 files in the sandbox — inline base64 or a url we fetch — and appends their paths to your prompt. anything the agent produces comes back as artifacts with a presigned download url, good for about ten minutes and re-mintable from the response endpoint.
yes — the opencode zen models ending in -free run without one. pick "harness": "opencode" and a -free model, drop provider_api_key entirely, and your epho key is the only credential in the request.
it's used for that run and nothing else. you re-send it on every turn, including resumes, and the sandbox stops the moment the turn ends.
something not covered here? [email protected]
__
/\ \
__ _____\ \ \___ ___
/'__`\/\ '__`\ \ _ `\ / __`\
/\ __/\ \ \L\ \ \ \ \ \/\ \L\ \
\ \____\\ \ ,__/\ \_\ \_\ \____/
\/____/ \ \ \/ \/_/\/_/\/___/
\ \_\
\/_/ grab a key. keep the curl.
two minutes from now you could be watching an agent work from your terminal.
