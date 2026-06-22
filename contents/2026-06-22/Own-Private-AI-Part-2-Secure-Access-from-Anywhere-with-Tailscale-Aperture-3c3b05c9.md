---
source: "https://10io.com/blog/private-ai-part-2-secure-access-via-tailscale-aperture"
hn_url: "https://news.ycombinator.com/item?id=48635255"
title: "Own Private AI, Part 2: Secure Access from Anywhere with Tailscale Aperture"
article_title: "10io — Actual Intelligence | Fractional CAIO / Chief AI Officer"
author: "anactofgod"
captured_at: "2026-06-22T20:02:29Z"
capture_tool: "hn-digest"
hn_id: 48635255
score: 1
comments: 0
posted_at: "2026-06-22T19:58:50Z"
tags:
  - hacker-news
  - translated
---

# Own Private AI, Part 2: Secure Access from Anywhere with Tailscale Aperture

- HN: [48635255](https://news.ycombinator.com/item?id=48635255)
- Source: [10io.com](https://10io.com/blog/private-ai-part-2-secure-access-via-tailscale-aperture)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T19:58:50Z

## Translation

タイトル: 独自のプライベート AI、パート 2: Tailscale Aperture を使用してどこからでも安全にアクセス
記事のタイトル: 10io — 実際のインテリジェンス | Fractional CAIO / 最高 AI オフィサー
説明: フルタイムのオーバーヘッドを発生させることなく、組織を AI 志向の組織から AI 対応組織に移行するための専門知識。フラクショナルおよび暫定最高 AI オフィサー (CAIO) / AI コンサルタント。

記事本文:
独自のプライベート AI、パート 2: Tailscale Aperture を使用してどこからでも安全にアクセス
Aperture AI ゲートウェイが前面にあり、公共のインターネットに公開されることのないプライベート Tailscale ネットワーク経由で、どこからでもセルフホスト LLM に安全にアクセスするための初心者向けのチュートリアルです。
パート 1 では、ステージ 1 と 2 を実行して、GB10 NVIDIA DGX Spark クラス マシンである「SparkStation」上で OpenAI 互換 API を提供する Qwen3.6-35B-A3B-FP8 Mixture of Experts (MoE) モデルを取得しました。ただし、モデルには、 localhost:8000 を介してマシン自体でのみアクセスできます。このパート 2 では、ステージ 3 から 5 までを実行して、モデルを公共のインターネットに公開することなく、選択した他のデバイスから安全にアクセスできるようにします。これらの手順は、SparkStation 以外のものによって提供される別のローカル AI モデルがある場合でも役立ちます。
私は主権 AI ソリューションを運用し、アクセスするためのより良い方法を常に探しています。 OpenAI や Anthropic のソリューションと同じくらい柔軟にアクセスできる、プライベート インフラストラクチャ上で実行されるセルフホスト型モデルが必要です。ただし、localhost:8000 で応答するモデルは、それが実行されているマシンでのみ使用できます。どこからでもアクセスできるようにする明白な方法は、ルーターを介してポートを転送することです。ただし、これは危険でもあります。認証されていない AI エンドポイントがオープン インターネット上に置かれ、誰でも見つけて悪用できるようになります。
代わりに、私の現在のアプローチには 2 つの層があります。
Tailscale — 承認された (「許可リストに登録された」) デバイス間の直接接続を、物理的にどこにあるかに関係なく、あたかも同じ LAN 上にあるかのように暗号化するプライベート メッシュ ネットワーク (「テールネット」) です。何も公に公開されることはありません。明示的に追加したデバイスのみが相互に接続できます。 Tailscale は非常に寛大な無料枠を提供しており、私はこれをセキュリティ対策に使用しています

年月を超えていない。走行距離は異なる場合があります。
Aperture (Tailscale 製) — テールネット上の 1 つ以上のモデルの前に配置される「AI ゲートウェイ」。呼び出し元の Tailscale ID によってすべてのリクエストを認証するため、API キーを配布する必要がなく、すべての使用状況を一元的に記録します。
このガイドに従うと、ローカルでホストされているモデルはプライベート ネットワーク経由でのみアクセス可能になり、ゲートウェイを経由するすべてのリクエストが識別されて記録されます。それは真にプライベートで安全なマネージド AI です。
コンセプトを一息に。テールネットはプライベート デバイス ネットワークです。 MagicDNS は、IP ではなく名前 (例: ゲートウェイ ) でデバイスをアドレス指定できるようにする Tailscale の機能です。 Aperture のプロバイダーは上流モデルです。許可とは、誰がどのモデルを使用できるかを示すルールです。以下でそれぞれに会います。
ステージ 3 — サーバーをプライベート ネットワーク上に配置する
まず、GB10 マシン (または「AI サーバー」として使用しているマシン) をテールネットに接続します。
3.1 サーバーに Tailscale をインストールして参加する
サーバーに Tailscale をインストールして起動します。
カール -fsSL https://tailscale.com/install.sh |しー
sudo テールスケールアップ
tailscale up は URL を出力します。任意のブラウザーで開き、サインインします (Google、GitHub、Microsoft、または電子メールはすべて機能します)。これにより、このマシンが認証され、テールネットに追加されます。サインインに使用するアカウントによってテールネットが定義されるため、どのアカウントを使用するかを覚えておいてください。すべてのデバイスが同じアカウントに参加する必要があります。
3.2 サーバーの Tailscale IP をメモします
テールスケール IP -4
Tailscale のプライベート スペースである 100.x.x.x の範囲のアドレスを取得します。以下では代用として 100.92.0.10 を使用します。自分のものに置き換えてください。これは、Aperture がモデルにアクセスするために使用するアドレスです。
vLLM はパート 1 で --network host を使用して起動されたため、すでにこの i でリッスンしています。

インターフェース — 変更の必要はありません。 ufw のようなファイアウォールを実行している場合は、Tailscale インターフェイス上のポートを許可します。 sudo ufw allowed in on tailscale0 から任意のポート 8000 まで。
ステージ 4 — Aperture を前面に配置
次に、ゲートウェイを追加します。 Aperture は、独自の Web ダッシュボードを備えたテールネット上で独自のノードとして実行されます。
aperture.tailscale.com にアクセスして、アクセス/サインアップをリクエストします。ベータ期間中は、Tailscale アカウントであれば無料です。プロビジョニングが完了すると、Aperture はホスト名を持つテールネット上のマシンとして表示され、次の場所にダッシュボードを提供します。
http://<あなたのアパーチャのホスト名>/ui
代替ホスト名としてゲートウェイを使用するので、ダッシュボードは http://gateway/ui になります。名前は独自の名前を持ち、Aperture のサインアップ フローと Tailscale 管理コンソールのマシンのリストに表示されます。
2 つの異なるダッシュボード - 混同しないでください。 login.tailscale.com/admin は、Tailscale 管理コンソール (ネットワーク: デバイス、ユーザー、アクセス ルールを管理します) です。 http://<aperture-host>/ui は、Aperture ダッシュボードです (モデル、プロバイダー、および使用状況を管理します)。以下のモデル構成は、Aperture ダッシュボードにあります。
4.2 モデルをプロバイダーとして追加する
Aperture ダッシュボードで [構成] を開き、生の HuJSON 構成 (Tailscale のコメント付き JSON 形式) を編集して、セルフホスト モデルをプロバイダーとして定義します。 「providers」: {...} ブロックを探します。サードパーティプロバイダー (Anthropic、Codex など) のデフォルトのリストがすでに含まれていることがわかります。 「providers」: {...} ブロック内の最初の行の直前に次の行を追加しました。
「主権者」: {
"baseurl": "http://100.92.0.10:8000",
"apikey": "ローカル認証なし",
"モデル": ["Qwen/Qwen3.6-35B-A3B-FP8"]
}、
重要な 3 つの詳細は、それぞれに午後を費やす可能性があります。
Baseurl には /v1 がありません。 Aperture は受信リクエストのパス (これはすべて

eady には /v1/chat/completions ) が Baseurl に含まれます。ここにも /v1 を追加すると、/v1/v1/... パスが壊れてしまいます。ステージ 3 のサーバーの Tailscale IP を 100.92.0.10 に置き換えることを忘れずに、ホストとポートのみを使用します。
apikeyは使い捨てです。 vLLM サーバーにはキーは必要ありませんが、Aperture のダッシュボードのテスト ボタンはキーがないと実行できません。空でない文字列は機能します。 vLLM はそれを無視します。
モデルは、vLLM が提供する正確なモデル ID と一致する必要があります (不明な場合は、サーバー上の http://localhost:8000/v1/models を確認してください)。
Aperture はデフォルトで拒否されています。管理者であっても、許可が与えられるまでモデルを呼び出すことはできません。したがって、同じ JSON 構成内で、「grants」: [...] ブロック (「providers」: {...} ブロックの後にあります) を見つけて、Aperture 経由でモデルにアクセスするための { "models": "**" } 機能が含まれていることを確認します。
「助成金」: [
{
"ソース": ["*"],
「アプリ」: {
"tailscale.com/cap/aperture": [
{ "役割": "管理者" },
{ "モデル": "**" }
】
}
}
】
これにより、テールネット ( "*" ) 上の全員に管理者ロールとすべてのモデル ( "**" ) へのアクセスが付与されます。これは個人的な設定には問題ありません。自分自身だけに制限するには、「*」を Tailscale のログイン名に置き換えます。設定を保存します (Aperture は保存時に警告をエラーとして扱うため、何か問題があるかどうかを通知します)。
4.4 ダッシュボードからルートをテストする
「モデル」タブを開きます。 Qwen/Qwen3.6-35B-A3B-FP8 の横に再生アイコンが表示されます。それをクリックします。緑色のチェックは、Aperture がテールネットを通じて vLLM サーバーに正常に到達したことを意味します。赤い X は、それができなかったことを意味します (通常はネットワーク アクセス ルール。トラブルシューティングを参照)。
ステージ 5 — クライアント コンピューターの接続
最後の部分: 別のデバイス (ラップトップ、私の場合は MacBook) からモデルにアクセスします。 OpenAI 互換ツールであればどれでも使用できますが、まずクライアントもテールネット上に存在する必要があります。
5.1 Tailscale を CLI にインストールして参加する

エント
macOS / Windows: tailscale.com/download (または Mac App Store) から Tailscale アプリをインストールし、起動して、サーバーで使用したのと同じアカウントでサインインします。
Linux:curl -fsSL https://tailscale.com/install.sh | sh 次に sudo tailscale up を実行し、同じアカウントでサインインします。
「同じアカウント」の部分は重要です。これは、クライアントをサーバーおよびゲートウェイと同じテールネットに置くものです。
クライアント上の端末から、まずゲートウェイを認識できることを確認し、その名前を解決します。
カール http://gateway/v1/models
モデルをリストした JSON が返された場合、MagicDNS は解決されており、テールネット パスは機能します。次に、真実の瞬間です。ラップトップからゲートウェイを介して、GB10 ボックスで実行されているモデルに実際のリクエストが送信されます。
カール http://gateway/v1/chat/completions \
-H "コンテンツ タイプ: application/json" \
-d '{"モデル":"Qwen/Qwen3.6-35B-A3B-FP8","messages":[{"role":"user","content":"ラップトップからこんにちは"}]}'
リクエストに API キーがありません。Tailscale ID が認証になります。応答が戻ってくると、ラップトップ → テールネット → Aperture → サーバー上の vLLM → 戻る という流れになります。 Aperture ダッシュボードの使用ログを確認すると、そのリクエストがあなたの ID とトークン数とともに記録されていることがわかります。
これは、必要に応じて機能するセットアップ全体です。非常に有能なプライベート モデルであり、ユーザーとユーザーのデバイスがどこにいてもアクセス可能で、ID によって認証され、一元的にログに記録され、公共のインターネットに公開されることはありません。
おめでとう！これで、独自のモデルへの安全なプライベート ゲートウェイが完成しました。自然な次のステップは、実際のアプリケーション (チャット フロントエンド、ノートブック ツール、コーディング アシスタント) をそれに向けることです。これらはすべて同じ 3 つの設定に減ります。
特定のアプリとさまざまなモデルの連携、およびアプリごと、モデルごとに生じる問題について調査する予定です。

— 今後の投稿で。しかし、主権 AI セットアップを試すことを妨げるものは何もありません。そして、ラインをドロップして、どのように楽しんでいるか、何に使用しているか、他の主権 AI プロジェクトへの提案を共有してください。
トラブルシューティングの要約 (パート 2)

## Original Extract

Expertise that moves your org from AI-curious to AI-capable without the full-time overhead. Fractional and interim Chief AI Officer (CAIO) / AI Consultant.

Your Own Private AI, Part 2: Secure Access from Anywhere with Tailscale Aperture
A beginner-friendly walkthrough for securely accessing your self-hosted LLM from anywhere — over a private Tailscale network fronted by the Aperture AI gateway, never exposed to the public internet.
In Part 1 we did Stages 1 and 2 to get a Qwen3.6-35B-A3B-FP8 Mixture of Experts (MoE) model serving an OpenAI-compatible API on a “SparkStation”, a GB10 NVIDIA DGX Spark-class machine. However, the model is only accessible on the machine itself via localhost:8000 . Here in Part 2, we run through Stages 3 to 5 to make the model securely reachable from any other devices you choose, without ever exposing it to the public internet. These instructions should be helpful even if you have a different local AI model being served by something other than a SparkStation.
I am constantly looking for a better way to operate and access sovereign AI solutions. I want self-hosted models running on private infrastructure that are as flexibly accessible as the solutions from OpenAI or Anthropic. But a model answering at localhost:8000 is only usable by the machine it runs on. The obvious way to make it accessible from anywhere is to forward a port through my router. However, this is also dangerous: it puts an unauthenticated AI endpoint on the open internet for anyone to find and abuse.
Instead, my current approach has two layers:
Tailscale — a private mesh network (“tailnet”) that encrypts direct connections between approved (“allowlisted”) devices, as if they were on the same LAN, no matter where they are physically. Nothing is exposed publicly; only devices that I’ve explicitly added can reach each other. Tailscale offers a very generous free tier, which I’ve been using for several months and have not yet exceeded. Your mileage may vary.
Aperture (by Tailscale) — an “AI gateway” that sits in front of one or more models on the tailnet. It authenticates every request by the caller’s Tailscale identity, so there are no API keys to distribute, and it logs all usage centrally.
If you follow this guide, your locally hosted models will be reachable only over your private network, and every request through the gateway will be identified and recorded. That’s genuinely private, secure, managed AI.
Concepts in one breath. A tailnet is your private device network. MagicDNS is Tailscale’s feature that lets you address devices by name (e.g. gateway ) instead of IP. A provider in Aperture is an upstream model. A grant is a rule saying who may use which models. You’ll meet each below.
Stage 3 — Put the server on your private network
First, get the GB10 machine (or whatever machine you are using as the “AI server”) onto your tailnet.
3.1 Install and join Tailscale on the server
On the server, install Tailscale and bring it up:
curl -fsSL https://tailscale.com/install.sh | sh
sudo tailscale up
tailscale up prints a URL — open it in any browser and sign in (Google, GitHub, Microsoft, or email all work). That authenticates this machine and adds it to your tailnet. The account you sign in with defines your tailnet, so remember which one you use — every device must join the same account.
3.2 Note the server’s Tailscale IP
tailscale ip -4
You’ll get an address in the 100.x.x.x range — Tailscale’s private space. I’ll use 100.92.0.10 as a stand-in below; replace it with your own . This is the address Aperture will use to reach your model.
Because vLLM was launched with --network host back in Part 1, it’s already listening on this interface — no change needed. If you run a firewall like ufw , allow the port on the Tailscale interface: sudo ufw allow in on tailscale0 to any port 8000 .
Stage 4 — Put Aperture in front
Now we add the gateway. Aperture runs as its own node on your tailnet with its own web dashboard.
Go to aperture.tailscale.com and request access / sign up. During the beta it’s free with any Tailscale account. Once provisioned, Aperture appears as a machine on your tailnet with a hostname, and serves a dashboard at:
http://<your-aperture-hostname>/ui
I’ll use gateway as the stand-in hostname, so my dashboard is http://gateway/ui . Yours will have its own name — you’ll find it in the Aperture sign-up flow and in your Tailscale admin console’s list of Machines.
Two different dashboards — don’t confuse them. login.tailscale.com/admin is the Tailscale admin console (manages your network: devices, users, access rules). http://<aperture-host>/ui is the Aperture dashboard (manages models, providers, and usage). The model configuration below lives in the Aperture dashboard.
4.2 Add your model as a provider
In the Aperture dashboard, open Configuration , and edit the raw HuJSON configuration (Tailscale’s JSON-with-comments format) to define your self-hosted model as a provider. Look for the "providers": {...} block. You may see it already has a default list of third-party providers (e.g., Anthropic, Codex). I just added the following lines inside the "providers": {...} block, right before the first of those lines:
"sovereign": {
"baseurl": "http://100.92.0.10:8000",
"apikey": "local-no-auth",
"models": ["Qwen/Qwen3.6-35B-A3B-FP8"]
},
Three details that matter, each of which can cost you an afternoon:
baseurl has no /v1 . Aperture appends the incoming request path (which already includes /v1/chat/completions ) to your baseurl . If you add /v1 here too, you get a broken /v1/v1/... path. Use just the host and port, remembering to substitute 100.92.0.10 for your server’s Tailscale IP from Stage 3.
apikey is a throwaway. Your vLLM server doesn’t require a key, but Aperture’s dashboard test button refuses to run without one. Any non-empty string works; vLLM ignores it.
models must match the exact model ID vLLM serves (check http://localhost:8000/v1/models on the server if unsure).
Aperture is deny-by-default : even as an admin, you can’t call a model until a grant says so. So, in the same JSON config, find the "grants": [...] block (it follows the "providers": {...} block), and make sure it includes a { "models": "**" } capability to access your model via Aperture:
"grants": [
{
"src": ["*"],
"app": {
"tailscale.com/cap/aperture": [
{ "role": "admin" },
{ "models": "**" }
]
}
}
]
This grants everyone on your tailnet ( "*" ) the admin role and access to all models ( "**" ) — fine for a personal setup. To restrict it to just yourself, replace "*" with your Tailscale login name. Save the config (Aperture treats warnings as errors on save, so it’ll tell you if anything’s off).
4.4 Test the route from the dashboard
Open the Models tab. You should see Qwen/Qwen3.6-35B-A3B-FP8 with a Play icon beside it. Click it — a green check means Aperture successfully reached your vLLM server through the tailnet. A red X means it couldn’t (usually a network access rule; see Troubleshooting).
Stage 5 — Connect a client computer
The final piece: reach the model from another device — a laptop, in my case a MacBook. Any OpenAI-compatible tool can use it, but first the client has to be on the tailnet too.
5.1 Install and join Tailscale on the client
macOS / Windows: install the Tailscale app from tailscale.com/download (or the Mac App Store), launch it, and sign in with the same account you used on the server.
Linux: curl -fsSL https://tailscale.com/install.sh | sh then sudo tailscale up , signing in with the same account.
The “same account” part is essential — it’s what puts the client on the same tailnet as the server and the gateway.
From a terminal on the client, first confirm it can see the gateway and resolve its name:
curl http://gateway/v1/models
If that returns JSON listing your model, MagicDNS is resolving and the tailnet path works. Then the moment of truth — a real request, from your laptop, through the gateway, to the model running on the GB10 box:
curl http://gateway/v1/chat/completions \
-H "Content-Type: application/json" \
-d '{"model":"Qwen/Qwen3.6-35B-A3B-FP8","messages":[{"role":"user","content":"hello from my laptop"}]}'
No API key in the request — your Tailscale identity is the authentication. When the reply comes back, it has traveled: laptop → tailnet → Aperture → vLLM on the server → back . Check the Aperture dashboard’s usage log and you’ll see that request recorded with your identity and a token count.
This is the whole setup working as required: a very capable private model, reachable from anywhere you and your devices are, authenticated by identity, logged centrally, and never exposed to the public internet.
Congratulations! You now have a secure, private gateway to your own model. A natural next step is to point real applications at it — chat front-ends, notebook tools, coding assistants — which all reduce to the same three settings:
I’m planning to explore wiring up specific apps and different models — and the per-app and per-model quirks that arise — in future posts. But, nothing is stopping you from trying out your sovereign AI setup. And, drop a line to share how you are enjoying it, what you are using it for, and suggestions for other sovereign AI projects!
Troubleshooting recap (Part 2)
