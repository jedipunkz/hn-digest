---
source: "https://mtgautodeck.com/articles/mtg-bench/"
hn_url: "https://news.ycombinator.com/item?id=48492177"
title: "MTG Bench: Testing how well LLMs can play magic"
article_title: "MTG Bench: Testing how well LLMs can play magic | MTG Auto Deck"
author: "CallumFerg"
captured_at: "2026-06-11T16:26:40Z"
capture_tool: "hn-digest"
hn_id: 48492177
score: 2
comments: 0
posted_at: "2026-06-11T16:01:01Z"
tags:
  - hacker-news
  - translated
---

# MTG Bench: Testing how well LLMs can play magic

- HN: [48492177](https://news.ycombinator.com/item?id=48492177)
- Source: [mtgautodeck.com](https://mtgautodeck.com/articles/mtg-bench/)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T16:01:01Z

## Translation

タイトル: MTG ベンチ: LLM がどれだけ上手にマジックをプレイできるかをテストする
記事のタイトル: MTG ベンチ: LLM がどれだけ上手にマジックをプレイできるかをテストする | MTGオートデッキ
説明: MTG ベンチは、LLM がどれだけマジックをプレイできるかをテストします。

記事本文:
MTG ベンチ: LLM がどれだけ上手にマジックをプレイできるかをテストする | MTG オートデッキ MTG オートデッキ ログイン サインアップ MTG ベンチ: LLM がどれだけ上手にマジックをプレイできるかをテストする
総合スコア
（高いほど良い）
上のチャートをクリックして、各ベンチマークのシミュレーションを表示します。
Fable 5 は占術ランドをプレイし、デッキの一番上のカードを確認します
Gemini 3.5 フラッシュは、占術、発見、チューター効果で複雑なターンを実行します
Opus 4.8 は誤ってカードをデッキに戻し、その間違いを自己報告します
Gpt 5.5は発見で追放されたカードをデッキに戻すのを忘れ、間違いを自己報告します
Fabel 5 はツールミスを犯し、その後黙ってターンを再開しようとします (後で評価によって発見されました)
主な考え方は、LLM が優れたマジックを実行できるほど賢い場合、LLM はルール エンジンを必要としないほど賢いということです。
法的措置を強制するルール エンジンはパフォーマンス フロアを向上させるでしょうが、シミュレーションの全体的な品質が向上するとは思いません。
各 LLM 呼び出しは、プリミティブ ライブラリ操作を使用して MCP サーバーにアクセスできます。デッキの一番上からカードを引く、カードをデッキの一番下に戻す、シャッフルするなどのことができます。
より高度な操作をシミュレートするには、次のようにします。
占術
または
監視、
複数のライブラリ ツール呼び出しを使用できます。
ライブラリ以外はすべて LLM によって管理されます。ベンチマークの合法性チェックとスコアリングはすべて gpt-5.5 (中) で行われました。
私のテストによると、LLM は実際に合法的なターンのシミュレーションを実行するよりも、シミュレートされたターンが合法的かどうかを評価する方がはるかに優れていました。
MCP サーバーの使用を選択した理由
すべてのデータと LLM API 呼び出しを完全に制御できるのに、なぜ基本的な関数/ツール呼び出しの代わりに MCP を使用するのでしょうか?
主な理由は、OpenAI と Anthropic では API リクエストでリモート MCP サーバーの URL を提供できるためです。
これは、OpenAI または

エージェント ループを人為的に処理します。これには 2 つの大きな利点があります。
これは 1 回の API 呼び出しであるため、ツールを使用するたびにキャッシュされた入力トークンのコストを支払う必要はありません (少なくとも OpenAI の場合は後で説明します)。
バッチ API を使用すると、ツール呼び出しのたびに新しいバッチを送信する必要がなく、50% 節約できます。
私の意見では、キャッシュされた入力トークンの課金方法は、エージェント ループにとっては意味がありません。この価格設定は、独立したリクエストにとって合理的です。
複数の独立した API 呼び出しが同じ大規模なシステム プロンプトで開始される場合、入力キャッシュにより、無料または少額のキャッシュ料金で割引が受けられます。
ただし、エージェント ループを使用すると、ツール呼び出しのたびに、大規模なシステム プロンプトに対するキャッシュされた入力コストが課金されます。
例を考えてみましょう。システム プロンプトがすでにキャッシュされており、ツール呼び出しによるトークンの使用が無視できる程度であると仮定します。
大規模なシステム プロンプト = 10,000 トークン
エージェントは 10 個のツール関数を呼び出します (並列ではありません)
請求対象のキャッシュされた入力トークン = 10,000 + 10,000 * 10 = 110,000 トークン
ツール機能の結果を待っている間に LLM がほんの一瞬しか停止していない場合、エージェントがターンするたびにシステム プロンプトに料金を請求するのは意味がないと思います。
これは、ツールを呼び出すために出力トークンがどのように必要になるか、ツール関数の結果は依然として入力トークンとして処理される必要があるなど、いくつかの詳細を見落としています。しかし、私の場合、
API コストの大部分は、エージェントがターンするたびにキャッシュされた入力トークンとして課金される大規模なシステム プロンプトによって決まります。
エージェント ループの料金は、アプリケーション コードにエージェント ループがあり、ツール呼び出しのたびに新しい API 呼び出しを行う場合に理解できます。
ただし、リモート MCP サーバーを提供し、エージェント ループを自分で処理しない場合は、さらに意味がありません。 OpenAI はそれを正しく処理します。
リモート MCP サーバーを使用した OpenAI への 1 つの API 呼び出しでは、入力プロンプトに対してのみ料金が発生します。

一度。リモート MCP を使用した Anthropic API 呼び出し
ただし、server は前の例と同様に動作します。
実際の数値として、gpt-5.5 (中) ベンチマークでは、マジック ターンあたりの平均入力トークンは 11,386 でした。 claude-fable-5 (中) の平均は 51,610 でした。
このベンチマークは、ツールを呼び出すことに熱心すぎるモデルを、ほとんどのベンチマークよりも厳しく罰します。
多くの場合、ツール呼び出しは情報を取得するだけなので、モデルがあまりにも多くのツールを呼び出すと、唯一の欠点が生じます。
入力トークンとツール結果のコンテキスト ウィンドウが無駄になります。ツールが状態を変更する場合でも、通常は
最終結果が正しくなるように元に戻してください。
シミュレーションマジックの場合はそうではありません。カードを引いた場合、それが間違いだったと気づいても、ただカードを戻すことはできません。
たとえカードをデッキに戻したとしても、そのカードが何であるかがわかっているので、シミュレーションは違法です。
一般的な障害モードは、モデルがツール呼び出しを開始した後、それが間違いであることに気づき、それを修正する方法がないというものでした。
すべてのライブラリ MCP 関数には、必須の理由フィールドがあります。見てみると
この例は Opus 4.8 からのもので、
「Draw forturn」という理由でそのターンにカードを引いた後、「No-op check は不要、キャンセル」という理由でカードを山札に戻していることがわかります。
次に、理由「noop」で「x」という名前のカードをデッキに戻し、次に理由「stop」で再度戻します。
私が作りました
MTGオートデッキ
バイブコーディングを試す方法として。私は LLM ベースのコーディングの現状についていけず、最終的にこのプロジェクトとベンチマークを何もせずに作成してしまいました。
一行のコードを手作業で書きます。
実装が速かったので、アカウントと支払いを備えたライブバージョンのみを作成しました。
の
プロジェクトは GitHub 上にあります
実行して独自の API キーまたはローカルの llama.cpp を使用したい場合
実際のところ、ライブバージョンにお金を払うことはお勧めしません。カレーと一緒に

魔法を正確に演奏できるモデルのコストと速度、
このアプリはあまり役に立ちません。一度に 1 ターンずつシミュレーションすると、オンラインのいずれかを使用してデッキを手動で金魚させるよりも時間がかかります。
ツール。また、何十ものシミュレーションを並行して実行して概要を得るのはコストがかかりすぎます。
より優れた安価な LLM がリリースされると、このアプリの便利なバージョンが見つかると思います。何百ものランニングを想像することができます
シミュレーションを行って、どのカードが良いか悪いかについての統計結果を示します。または、デッキを交換することで自動的に最適化します。
あなたのためのカード。

## Original Extract

MTG Bench tests how well LLMs can play Magic.

MTG Bench: Testing how well LLMs can play magic | MTG Auto Deck MTG Auto Deck Log in Sign up MTG Bench: Testing how well LLMs can play magic
Overall Score
(higher is better)
Click on the charts above to view each benchmark's simulations.
Fable 5 plays a scry land and looks at the top card of the deck
Gemini 3.5 flash performs complex turn with scry, discover, and tutor effects
Opus 4.8 erroneously returns a card to the deck then self reports the mistake
Gpt 5.5 forgets to return cards exiled with discover to the deck and self reports the mistake
Fabel 5 makes a tool mistake, then silently tries to restart the turn (caught by evaluation later)
The main idea is that if an LLM is smart enough to play good magic, then it is also smart enough to not need a rules engine.
A rules engine that enforces legal actions would improve the performance floor, but I don't think it would improve the overall quality of the simulation.
Each LLM call has access to an MCP server with primitive library operations. It can do things like draw a card from the top of the deck, return card to bottom of deck, and shuffle.
To simulate more advanced operations, like
scry
or
surveil ,
it can use multiple library tool calls.
Everything other than the library is managed by the LLM. Legality checks and scoring for the benchmarks was all done with gpt-5.5 (medium).
From my testing, LLMs were much better at evaluating if a simulated turn was legal than they were at actually performing a legal turn simulation.
Why I choose to use an MCP server
I have full control over all of the data and the LLM api calls, so why use MCP instead of basic function/tool calling?
The main reason is that OpenAI and Anthropic allow you to provide a remote MCP server url in an api request.
This means that OpenAI or Anthropic handle the agent loop. This has two major benefits.
Since it is one api call, you don't pay for the cached input token cost after each tool use (at least with OpenAI. more on that later)
You can use the batch api for 50% savings without having to submit a new batch after every tool call
In my opinion, the way cached input tokens are charged does not make sense for agent loops. The pricing makes sense for independent requests.
If multiple independent api calls start with the same large system prompt, input caching gets you a discount for free, or for a small caching fee.
With an agent loop, however, you are charged the cached input cost for a large system prompt after every tool call.
Consider an example. Assume the system prompt is already cached and tool calls result in negligible token use.
Large system prompt = 10k tokens
Agent calls 10 tool functions (not parallel)
Billed cached input tokens = 10k + 10k * 10 = 110k tokens
I don't think it makes sense to charge for the system prompt after every agent turn if the LLM is only pausing for a fraction of a second while waiting for a tool function result.
This is overlooking some details, like how it takes output tokens to call a tool, and the tool function result still needs to be processed as input tokens. But in my case,
the api cost is dominated by the large system prompt being charged as cached input tokens after every agent turn.
The pricing for an agent loop is understandable when your application code has the agent loop, and is making a new api call after each tool call.
But it makes even less sense when you provide a remote MCP server and do not handle the agent loop yourself. OpenAI handles it correctly.
A single api call to OpenAI with a remote MCP server will only ever charge you for the input prompt once. An Anthropic api call with remote MCP
server, however, works like the previous example.
Some real numbers, the gpt-5.5 (medium) benchmark had an average input tokens per magic turn of 11,386. The average for claude-fable-5 (medium) was 51,610.
This benchmark punishes models that are too eager to call tools more than most benchmarks.
In many cases, tool calls are only retrieving information, so if a model calls too many tools, the only downside
is wasted input tokens and context window for the tool results. Even if the tool mutates state, it can usually
be undone so the final result is correct.
This is not the case when simulation magic. If you draw a card, then realize that was a mistake, you can't just put it back.
Even if you do return the card to the deck, you now know what that card is, so the simulation is illegal.
A common failure mode was the model starting a tool call, then realizing it was a mistake and having no way to correct it.
All the library MCP functions have a required reason field. If you look at
this example from Opus 4.8 ,
you can see that it draws a card for turn with reason "Draw for turn", then returns the card to the deck with reason "No-op check not needed; cancel".
It then proceeds to return a card named "x" to the deck with reason "noop", then again with reason "stop".
I made
MTG Auto Deck
as a way to try out vibe coding. I had not been keeping up with the state of LLM based coding, and I ended up making this project and the benchmark without
writing a single line of code by hand.
I only made a live version with accounts and payments because of how quick it was to implement.
The
project is on GitHub
if you want to run it and use your own api keys or local llama.cpp
I wouldn't actually recommend paying for the live version. With the current cost and speed of models that that can accurately play magic,
the app does not provide much utility. Simulating turns one at a time is slower than manually goldfishing your deck using one of the online
tools. And it is too expensive to run dozens of simulations in parallel and give you a summary.
As better cheaper LLMs get released, I think there is some version of this app that would be useful. I can imagine running hundreds
of simulations, then giving statistical results about which cards are good and bad. Or automatically optimizing a deck by swapping out
cards for you.
