---
source: "https://totaldebug.uk/posts/self-hosted-llms-ollama-open-webui-docker-compose/"
hn_url: "https://news.ycombinator.com/item?id=48828676"
title: "Run your own self-hosted LLMs with Docker Compose"
article_title: "Run your own ChatGPT: self-hosted LLMs with Docker Compose | TotalDebug"
author: "marksie1988"
captured_at: "2026-07-08T07:31:02Z"
capture_tool: "hn-digest"
hn_id: 48828676
score: 1
comments: 0
posted_at: "2026-07-08T07:22:46Z"
tags:
  - hacker-news
  - translated
---

# Run your own self-hosted LLMs with Docker Compose

- HN: [48828676](https://news.ycombinator.com/item?id=48828676)
- Source: [totaldebug.uk](https://totaldebug.uk/posts/self-hosted-llms-ollama-open-webui-docker-compose/)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T07:22:46Z

## Translation

タイトル: Docker Compose を使用して独自のセルフホスト LLM を実行する
記事のタイトル: 独自の ChatGPT を実行する: Docker Compose を使用した自己ホスト型 LLM |トータルデバッグ
説明: 私は毎日クロード コードを使用していますが、個人データ、個人的なメモ、中途半端なアイデアなど、すべてがパブリック ドメインに属するわけではありません。そういう人たちに対しては、会話をローカルに保ちたいと思っています。

記事本文:
独自の ChatGPT を実行する: Docker Compose を使用した自己ホスト型 LLM |トータルデバッグ
≡ メニュー
// テクノロジー。裏返しに
$ツリー〜
§─ ~/ ホーム h
© 2026 スティーブン・マークス
>_ で構築されました
>_ Docker · ホームラボ
独自の ChatGPT を実行する: Docker Compose を使用した自己ホスト型 LLM
私は毎日クロード コードを使用していますが、個人データ、個人的なメモ、途中で形成されたアイデアなど、すべてがパブリック ドメインに属するわけではありません。そういう人たちに対しては、会話をローカルに保ちたいと思っています。
そこで私はローカルアシスタントも運営しています。小さな Docker Compose スタックを使用すると、商用のチャット UI によく似たチャット UI を備えたプライベート LLM が得られます。この種のプロンプトでは、トークンごとの請求がなく、ローカル ネットワークから何も出ないことを意味します。
全体は 2 つのサービスで構成されています。モデルを実行する Ollama とモデルと対話する Open WebUI です。
Ollama はモデルをプルし、ディスク上に保持し、ポート 11434 上の単純な HTTP API 経由で提供します。
オープン WebUI がフロントエンドです。これは、会話、複数のモデル、ユーザー アカウント、ファイルのアップロードを備えた、洗練された自己ホスト型チャット インターフェイスです。 API を介して Ollama と通信するため、ChatGPT スタイルのエクスペリエンスをローカルで得ることができます。
2 つのコンテナー、1 つのネットワーク、永続化のための 2 つのボリューム。それが全体のデザインです。
ホームラボで何かを実行する場合は、おそらくすでにインストールされている Docker が必要です。
小規模モデル以外の場合は、ホストに NVIDIA GPU と NVIDIA Container Toolkit をインストールすることも必要です。ツールキットは、コンテナーが GPU を認識できるようにするものです。これがないと、Compose ファイル内の GPU ブロックは起動できず、代わりに CPU のみのバリアントが必要になります (これについては後ほど説明します)。ツールキットが動作していることを確認できます。
1
docker run --rm --gpus all nvidia/cuda:12.4.0-base-ubuntu22.04 nvidia-smi
GPU が表示されれば準備完了です。エラーが発生した場合は、まずツールキットを修正してください。LLM スタックはまさに​​その p に依存しているためです。

ぐずぐずする。
これがスタック全体です。重要な部分を下に分けて説明します。
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
サービス:
オラマ :
画像：オラマ/オラマ
コマンド: サーブ
公開する:
- 11434/tcp
ヘルスチェック:
テスト: ollam --version ||出口1
ボリューム:
- オラマ:/root/.ollama
デプロイ:
リソース:
予約：
デバイス:
- ドライバー: nvidia
device_ids : [ "すべて" ]
機能: [GPU]
ウェブイ：
画像：ghcr.io/open-webui/open-webui:main
ポート:
- 8080:8080/tcp
環境：
- OLLAMA_BASE_URL=http://ollama:11434
ボリューム:
- open-webui:/app/backend/data
依存関係 :
- オラマ
ボリューム:
オラマ :
オープンウェブイ:
これを docker-compose.yml として保存し、 docker compose up -d を実行してフルスタックをオンラインにします。
実行する前に、以下でもう少し詳しくその仕組みを理解する価値があります。
2 つのサービスがお互いを見つける方法
OLLAMA_BASE_URL=http://ollama:11434 に注目してください。その ollam は私が設定したホスト名ではなく、ollam のサービス名です。 Compose は、すべてのサービスにプロジェクトのデフォルト ネットワーク上の DNS エントリを与えます。そのため、Open WebUI は、ホスト ポートや IP アドレスを使用せずに、内部ネットワーク経由で名前によって Ollama に到達します。
これが、Ollama が ports ではなく opens を使用する理由でもあります。公開すると、11434 が同じネットワーク上の他のコンテナに到達可能になりますが、ホストには公開されません。私が実際に公開しているのは Open WebUI の 8080 だけです。これは人間が到達する必要がある 1 つのインターフェイスだからです。生のモデル API をホストから遠ざけることは、小さなことですが実際の衛生管理であり、それがなぜ重要なのかについてはまた改めて説明します。
GPU パススルー、誰もが引っかかる部分
このブロックは、GPU をコンテナーに渡すものです。
1
2
3
4
5
6
7
デプロイ:
リソース:
予約：
デバイス:
- ドライバー: nvidia
device_ids : [ "すべて"

】
機能: [GPU]
知っておく価値のあることがいくつかあります:
機能: [GPU] が必要です。これを省略すると、残りのブロックが存在しても Docker は GPU アクセスを許可しません。
device_ids: ["all"] はコンテナーにすべての GPU を与えます。複数のカードをお持ちで、Ollama を特定のカードに固定したい場合は、それを device_ids: ["0"] のようなものに置き換えます。
これは、 Compose で docker run --gpus all に相当します。これは、前のセクションの NVIDIA Container Toolkit がインストールされている場合にのみ機能します。 「始まらない」ということは、ほとんどの場合ここに遡ります。
これを正しく行うと、Ollama がモデルを VRAM にロードして GPU で実行します。これが、使用可能なアシスタントと待機中のアシスタントの違いです。
2 つの名前付きボリュームはモデルを永続化し、コンテナーが再起動されるたびに数ギガバイトのモデルが再ダウンロードされるのを防ぎます。
ollama:/root/.ollama はプルしたモデルを保持します。
open-webui:/app/backend/data には、アカウント、チャット、設定が保存されます。
Ollama のヘルスチェックは、「おそらくプロセスが実行中です」ではなく、実際の準備完了信号を Docker に提供します。これは単純ですが、depends_on と独自のツールがエンジンが実際に稼働しているかどうかを推論できることを意味します。
1
2
ドッカー構成 -d
ドッカー構成PS
Open WebUI は http://localhost:8080 にあります。最初に作成したアカウントが管理者となるため、それをどこかに公開する前に作成してください。
スタックは実行中ですが、Ollama にはまだモデルがありません。 1 つを引いてください:
1
docker compose exec ollama ollama pull llama3.1:8b
llama3.1:8b は、8 GB カードに快適に収まる一般的な開始点として適しています。 VRAM が厳しい場合は、 qwen2.5:3b などの小さいモデルを試してください。余裕がある場合は、より大きなものを引っ張ってください。ダウンすると、Open WebUI のモデル ピッカーに表示され、チャットを開始できるようになります。モデルを直接プルすることもできます。

UI ですが、コマンド ラインで一度実行すると、どこで作業が行われたかが明確になります。
これを試すのに GPU は必要ありません。 ollam サービスからdeploy: ブロック全体を削除します。他のものはすべて同じままです。
1
2
3
4
5
6
7
8
9
オラマ :
画像：オラマ/オラマ
コマンド: サーブ
公開する:
- 11434/tcp
ヘルスチェック:
テスト: ollam --version ||出口1
ボリューム:
- オラマ:/root/.ollama
CPU上では快適に動作します。現実的に考えてください。小さなモデルは動作が遅く、大きなモデルは壊れているように感じます。 CPU のみはタイヤを蹴るのに最適で、3B クラスのモデルでの軽い使用には問題ありません。毎日使用するものはすべて、GPU によって快適になります。
これを裸でインターネットに公開しないでください
Open WebUI には独自の認証があり、これは優れています。これはまだ、生のポートでインターネットに直接公開するようなものではなく、Ollama API には認証がまったくありません。まさにこれが、上記の Compose ファイルがホストに 11434 を公開しない理由です。
家の外からこれにアクセスしたい場合は、TLS を終端するリバース プロキシの背後に配置し、理想的には認証層を前面に配置します (私はホーム ネットワークの外部からアクセスせずにローカルのみで実行しています)。
実際に自分で実行する価値はあるのでしょうか?
私にとっては、1 つだけ注意点があれば、それだけの価値があると思います。完全なトレードオフは次のとおりです。
コスト。必要なハードウェアをすでに所有している場合は、基本的に無料です。トークンごとの請求や突然の請求はありません。
プライバシー。ネットワークからは何も出ません。あなたのプロンプトはあなたのものです。
待ち時間。まともな GPU では、最初のトークンのレイテンシーは非常に競争力があり、プロバイダーとのネットワークの往復はありません。
モデルの品質。自宅で実行できるオープン モデルは非常に優れており、さらに改良されていますが、より複雑なタスクでは依然として最大のフロンティア モデルが優位性を持っています。要約、草案、コーディングのヘルプ、一般的な質問については、l

オーカルモデルは十分すぎるほどです。
私個人としては、最先端の高度な作業にはクラウド ツールを使用し、個人的なものやプライベートなものはすべてローカル モデルに送られ、トラフィックが家から出ることはありません。上記のシンプルな積み重ねにより、プロジェクトではなく 5 分で意思決定が完了します。
2 つのサービス、1 つの Compose ファイル、そして独自のハードウェア上で実行されるプライベート ChatGPT のようなプロンプトを手に入れることができます。セットアップするための非常にシンプルなコンテナーがあります。サービス名で Open WebUI を Ollama に指定し、デプロイ ブロック (およびその背後にある NVIDIA ツールキット) で GPU を渡し、モデルとチャットを名前付きボリュームに保持し、推論 API を公開しません。これらを正しく理解すれば、あとはどのモデルを選択するかを選択するだけです。
> Homer ダッシュボードの構成 2022-10-16
> Docker を使用した Homer ダッシュボード 2022-10-15
> 実稼働用の CentOS を使用した Docker Overlay2 2020-05-05

## Original Extract

I use Claude Code every day but not everything belongs in the public domain: personal data, private notes, half-formed ideas. For those I’d rather keep the conversation local.

Run your own ChatGPT: self-hosted LLMs with Docker Compose | TotalDebug
≡ menu
// technology. inside out
$ tree ~
├─ ~/ home h
© 2026 Steven Marks
built with >_
>_ Docker · Homelab
Run your own ChatGPT: self-hosted LLMs with Docker Compose
I use Claude Code every day but not everything belongs in the public domain: personal data, private notes, half-formed ideas. For those I’d rather keep the conversation local.
So I also run a local assistant. A small Docker Compose stack gives me a private LLM with a chat UI that looks a lot like the commercial ones and for that kind of prompt it means no per-token bill and nothing leaving my local network.
The whole thing is two services: Ollama to run the models and Open WebUI to talk to them.
Ollama pulls models, keeps them on disk, and serves them over a simple HTTP API on port 11434 .
Open WebUI is the front end. It’s a polished, self-hosted chat interface with conversations, multiple models, user accounts and file uploads. It talks to Ollama over the API, so you get a ChatGPT-style experience locally.
Two containers, one network, two volumes for persistence. That’s the entire design.
You need Docker which is likely already installed if you run anything in a homelab.
For anything beyond small models you also want an NVIDIA GPU and the NVIDIA Container Toolkit installed on the host. The toolkit is what lets a container see the GPU. Without it, the GPU block in the Compose file will fail to start, and you’ll want the CPU-only variant instead (I cover that later). You can check the toolkit is working with:
1
docker run --rm --gpus all nvidia/cuda:12.4.0-base-ubuntu22.04 nvidia-smi
If that prints your GPU, you’re ready. If it errors, fix the toolkit first, because the LLM stack relies on exactly that plumbing.
Here’s the stack in full. I’ll break down the parts that matter underneath.
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
services :
ollama :
image : ollama/ollama
command : serve
expose :
- 11434/tcp
healthcheck :
test : ollama --version || exit 1
volumes :
- ollama:/root/.ollama
deploy :
resources :
reservations :
devices :
- driver : nvidia
device_ids : [ " all" ]
capabilities : [ gpu ]
webui :
image : ghcr.io/open-webui/open-webui:main
ports :
- 8080:8080/tcp
environment :
- OLLAMA_BASE_URL=http://ollama:11434
volumes :
- open-webui:/app/backend/data
depends_on :
- ollama
volumes :
ollama :
open-webui :
Save it as docker-compose.yml and run docker compose up -d to bring the full stack online.
Before you do it’s worth understanding how it works in a little more detail below.
How the two services find each other
Notice OLLAMA_BASE_URL=http://ollama:11434 . That ollama is not a hostname I’ve configured, it’s the service name for ollama. Compose gives every service a DNS entry on the default network for the project. So Open WebUI reaches the Ollama by name, over the internal network, with no host ports or IP addresses involved.
That’s also why Ollama uses expose rather than ports . expose makes 11434 reachable to other containers on the same network but does not publish it to the host. The only thing I actually publish is Open WebUI’s 8080 , because that’s the one interface a human needs to reach. Keeping the raw model API off the host is a small but real bit of hygiene and I’ll come back to why it matters.
GPU passthrough, the part everyone gets stuck on
This block is what hands the GPU to the container:
1
2
3
4
5
6
7
deploy :
resources :
reservations :
devices :
- driver : nvidia
device_ids : [ " all" ]
capabilities : [ gpu ]
A few things worth knowing:
capabilities: [gpu] is required. If you omit it, Docker will not grant GPU access even with the rest of the block present.
device_ids: ["all"] gives the container every GPU. If you have more than one and want to pin Ollama to a specific card, swap it for something like device_ids: ["0"] .
This is the Compose equivalent of docker run --gpus all . It only works if the NVIDIA Container Toolkit from the previous section is installed. This is where “it won’t start” almost always traces back to.
Get this right and Ollama will load models into VRAM and run them on the GPU, which is the difference between a usable assistant and one you wait on.
The two named volumes persist the models to stop re-downloading several gigabytes of model every time a container restarts:
ollama:/root/.ollama keeps your pulled models.
open-webui:/app/backend/data keeps your accounts, chats and settings.
The healthcheck on Ollama gives Docker a real readiness signal instead of “the process is running, probably”. It’s simple, but it means depends_on and your own tooling can reason about whether the engine is actually up.
1
2
docker compose up -d
docker compose ps
Open WebUI is now on http://localhost:8080 . The first account you create becomes the admin, so do that before you expose it anywhere.
The stack is running, but Ollama has no models yet. Pull one:
1
docker compose exec ollama ollama pull llama3.1:8b
llama3.1:8b is a good general starting point that fits comfortably on an 8 GB card. If you’re tight on VRAM, try a smaller model such as qwen2.5:3b ; if you’ve got headroom, pull something larger. Once it’s down, it shows up in the model picker in Open WebUI and you can start chatting. You can also pull models straight from the UI, but doing it once on the command line makes it obvious where the work happens.
You don’t need a GPU to try this. Drop the entire deploy: block from the ollama service and everything else stays the same:
1
2
3
4
5
6
7
8
9
ollama :
image : ollama/ollama
command : serve
expose :
- 11434/tcp
healthcheck :
test : ollama --version || exit 1
volumes :
- ollama:/root/.ollama
It will run happily on CPU. Be realistic about it, a small model will feel sluggish and a large one will feel broken. CPU-only is great for kicking the tyres and fine for light use with a 3B-class model. For anything you’ll use daily, the GPU is what makes it pleasant.
Don’t put this on the internet naked
Open WebUI has its own authentication, which is good. It is still not something I’d publish straight to the internet on a raw port, and the Ollama API has no auth at all, which is exactly why the Compose file above never publishes 11434 to the host.
If you want to reach this from outside the house, put it behind a reverse proxy that terminates TLS and ideally, an auth layer in front (I run mine local only with no access from outside my home network).
Is running it yourself actually worth it?
I find that it is worth it for me with one real caveat. Here’s the full trade-off:
Cost. If you already own the required hardware then its essentially free. No per-token billing, no surprise invoice.
Privacy. Nothing leaves your network. Your prompts are yours.
Latency. On a decent GPU, first-token latency is genuinely competitive, and there’s no network round trip to a provider.
Model quality. The open models you can run at home are very good and getting better, but the largest frontier models still have an edge on the more complex tasks. For summarising, drafting, coding help and general questions, local models are more than good enough.
For me personally, the cloud tools stay in the mix for the heavy, cutting-edge work, and anything personal or private goes to the local model instead, where the traffic never leaves the house. The simple stack above is what makes it a five-minute decision rather than a project.
Two services, one Compose file, and you’ve got a private ChatGPT-like prompt running on your own hardware. It has very simple containers to setup: point Open WebUI at Ollama by its service name, hand the GPU over with the deploy block (and the NVIDIA toolkit behind it), keep your models and chats in named volumes, and never publish the inference API. Get those right and the rest is just picking which models to pull.
> Configuring Homer Dashboard 2022-10-16
> Homer dashboard with Docker 2022-10-15
> Docker Overlay2 with CentOS for production 2020-05-05
