---
source: "https://samihonkonen.com/posts/purpose-built-local-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48473010"
title: "Purpose-built local AI agents"
article_title: "Purpose-built local AI agents | Sami Honkonen"
author: "shonkone"
captured_at: "2026-06-10T10:30:24Z"
capture_tool: "hn-digest"
hn_id: 48473010
score: 1
comments: 0
posted_at: "2026-06-10T08:00:58Z"
tags:
  - hacker-news
  - translated
---

# Purpose-built local AI agents

- HN: [48473010](https://news.ycombinator.com/item?id=48473010)
- Source: [samihonkonen.com](https://samihonkonen.com/posts/purpose-built-local-ai-agents/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T08:00:58Z

## Translation

タイトル: 専用のローカル AI エージェント
記事のタイトル: 専用のローカル AI エージェント |サミ・ホンコネン
説明: アイドル状態の Mac Studio をプライベート AI バックエンドに変え、その上に専用のエージェントを構築した方法。

記事本文:
去年、私は主に Beata のボーカルをトラッキングする音楽プロジェクトのために Mac Studio を購入しました。ローカル LLM を実行する人が増えているという記事を読んだとき、自分が所有している強力なマシンは、ほとんどが電源を切った状態で机の上に置かれているだけであることに気付きました。音楽セッション後にシャットダウンするのではなく、実行したままにしてローカル モデルを実行したらどうなるでしょうか?
まず、 Tailscale を設定します。私のデバイス間にプライベート ネットワークが作成されるため、どちらがどこにいても、私の MacBook Air は安定したホスト名を介して Mac Studio にアクセスできます。次に、Studio で SSH を有効にしました。これで、Air 上で「ssh studio」と入力するだけで済みます。
モデルを選択するのは、whollm よりも簡単です。実行するだけでハードウェアが自動検出され、システムに適合する HuggingFace の上位モデルがリストされます。
モデルは LM Studio で実行します。 GUI を使用してモデルをダウンロードし、API トークンを設定しました。
その後はすべて CLI なのでリモートで実行できます。詳細をラップするために、Studio の ~/bin にいくつかの小さなスクリプトを書きました。起動スクリプトは大まかに次のようになります。
モデル = "qwen/qwen3.6-27b"
MODEL_IDENTIFIER = "スタジオ-llm"
ポート = "4000"
lms サーバーの開始 --port " $PORT " --bind 0.0.0.0
lms アンロード --all
lms ロード " $MODEL " \
--context-length 65536 \
--GPU 最大 \
--identifier " $MODEL_IDENTIFIER "
停止するのは lms server stop だけです。私の Air からは ssh Studio '~/bin/start-lmstudio.sh' または stop です。 --bind 0.0.0.0 を指定すると、ローカルホストだけでなくネットワーク上でリッスンするようになり、Air が Tailscale 経由でアクセスできるようになります。
クライアントが構成を変更する必要なく、基礎となるモデルを交換できるようにしたかったので、開始スクリプト内のモデル識別子を修正しました。 qwen/qwen3.6-27b をクライアントにアドバタイズする代わりに、サーバーは studio-llm をアドバタイズします。
アイデアは、それに接続するすべてのものは、

名前は studio-llm です。別のモデルを試したい場合は、別のモデル名で lms load を実行し、 --identifier studio-llm を保持します。他には何も変わりません。エンドポイントとモデル名は安定したままになります。このように、基礎となるモデルは、消費者さえ知らない単なる実装の詳細になります。
重い GUI ではなくコマンドラインから実行したいことがたくさんあります。もちろん、Claude Code を使用することも選択肢の 1 つですが、これはコーディングに特化しており、特にプロンプ​​トのほとんどが当面のタスクに関連しない場合、システム プロンプトはローカル LLM にとっては長くなります。
pi は最小限のアプローチを採用しているため、一般的なエージェントのフロントエンドに最適です。 LM Studio を指すことは、 ~/.pi/agent/models.json の構成エントリにすぎません。
{
「プロバイダー」: {
"lmstudio" : {
"baseUrl" : "http://studio:4000/v1" ,
"apiKey" : "<lmstudio から>" ,
"api" : "openai-completions" ,
"互換性" : {
"supportsDeveloperRole" : false ,
"supportsReasoningEffort" : false
}、
「モデル」: [
{
"id" : "studio-llm" ,
"名前" : "スタジオ LLM" ,
"入力" : [ "テキスト" , "画像" ],
"コンテキストウィンドウ" : 65536 、
"コスト" : { "入力" : 0 、 "出力" : 0 、 "キャッシュ読み取り" : 0 、 "キャッシュ書き込み" : 0 }
}
】
}
}
}
BaseUrl はホスト名として Studio を使用し、Tailscale で解決されます。セッションを開くたびにフラグを渡す必要がないように、lmstudio/studio-llm を pi のデフォルト モデルとして設定しました。
専用の AI エージェントの作成 #
特定の何かについて AI の助けが必要なときは、Air の ~/projects/ の下に新しいディレクトリを作成し、pi で作業を開始します。やりたいことを完了したら、そのプロセスを AGENTS.md に記録するように pi に指示します。それ以降、そのディレクトリで pi を開くたびにファイルが読み取られ、すぐに続行できるようになります。
その結果、専用のエージェントのコレクションがディスク上に配置されます。

特定のジョブに合わせて形成された ch。意図的にクラウド モデルに切り替えない限り、データは自分のマシン上に残ります (より大きなモデルを使用したい場合に切り替えることもあります)。ラップトップを詰まらせることなく、ローカル モデルが Studio 上で実行されます。

## Original Extract

How I turned an idle Mac Studio into a private AI backend, and built purpose-built agents on top of it.

Last year I bought a Mac Studio for music projects, mostly tracking vocals for Beata . When I read about more and more people running local LLMs, I realized I have a powerful machine that mostly just sits on a desk with power off. What if, instead of shutting it down after a music session, I just left it running and ran a local model on it?
First, I set up Tailscale . It creates a private network between my devices, so my MacBook Air can reach the Mac Studio over a stable hostname regardless of where either of us is. Then I enabled SSH on the Studio. Now I can just type ssh studio on my Air.
Picking a model doesn’t get easier than whichllm . You just run it and it autodetects your hardware and lists the top models from HuggingFace that fit your system.
I run the model with LM Studio . I used the GUI to download the model and to set an API token.
After that, everything is CLI so I can do it remotely. I wrote a couple of small scripts in ~/bin on the Studio to wrap the details. The start script looks roughly like this:
MODEL = "qwen/qwen3.6-27b"
MODEL_IDENTIFIER = "studio-llm"
PORT = "4000"
lms server start --port " $PORT " --bind 0.0.0.0
lms unload --all
lms load " $MODEL " \
--context-length 65536 \
--gpu max \
--identifier " $MODEL_IDENTIFIER "
Stopping it is just lms server stop . From my Air it’s ssh studio '~/bin/start-lmstudio.sh' or stop . The --bind 0.0.0.0 makes it listen on the network rather than just localhost, so my Air can reach it over Tailscale.
I wanted to be able to swap the underlying model without the clients needing to change their configuration so I fixed the model identifier in the start script. Instead of advertising qwen/qwen3.6-27b to clients, the server advertises studio-llm .
The idea is that everything connecting to it uses the name studio-llm . When I want to try a different model, I run lms load with a different model name and keep --identifier studio-llm . Nothing else changes. The endpoint and model name stay stable. This way the underlying model is just an implementation detail that the consumers don’t even know about.
There’s plenty of stuff I like to do from the command-line instead of a heavy GUI. Obviously using Claude Code is an option, but it’s heavily geared for coding and its system prompts are lengthy for a local LLM, especially if most of it is not relevant for the task at hand.
pi has a minimalistic approach and thus it’s a great fit for a general agent front end. Pointing it at LM Studio is just a config entry in ~/.pi/agent/models.json :
{
"providers" : {
"lmstudio" : {
"baseUrl" : "http://studio:4000/v1" ,
"apiKey" : "<from lmstudio>" ,
"api" : "openai-completions" ,
"compat" : {
"supportsDeveloperRole" : false ,
"supportsReasoningEffort" : false
},
"models" : [
{
"id" : "studio-llm" ,
"name" : "Studio LLM" ,
"input" : [ "text" , "image" ],
"contextWindow" : 65536 ,
"cost" : { "input" : 0 , "output" : 0 , "cacheRead" : 0 , "cacheWrite" : 0 }
}
]
}
}
}
The baseUrl uses studio as the hostname, which resolves over Tailscale. I set lmstudio/studio-llm as pi’s default model so I don’t have to pass a flag every time I open a session.
Creating a purpose-built AI agent #
Whenever I want AI help with something specific, I make a new directory under ~/projects/ on my Air and just start working with pi. Once I’ve done what I want to do, I tell pi to record the process in an AGENTS.md . From that point on, every time I open pi in that directory it reads the file and is immediately ready to continue.
The result is a collection of purpose-built agents sitting on disk, each shaped to a specific job. The data stays on my own machines unless I deliberately switch to a cloud model (which I sometimes do when I want to use a bigger model), and the local model runs on the Studio rather than clogging up my laptop.
