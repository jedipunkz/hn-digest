---
source: "https://blog.skypilot.co/sandboxes/"
hn_url: "https://news.ycombinator.com/item?id=48506318"
title: "SkyPilot Sandboxes: Run Agent Code on Your Own Kubernetes, at Scale"
article_title: "SkyPilot Sandboxes: Run Agent Code on Your Own Kubernetes, at Scale | SkyPilot Blog"
author: "covi"
captured_at: "2026-06-12T18:47:05Z"
capture_tool: "hn-digest"
hn_id: 48506318
score: 6
comments: 0
posted_at: "2026-06-12T16:40:35Z"
tags:
  - hacker-news
  - translated
---

# SkyPilot Sandboxes: Run Agent Code on Your Own Kubernetes, at Scale

- HN: [48506318](https://news.ycombinator.com/item?id=48506318)
- Source: [blog.skypilot.co](https://blog.skypilot.co/sandboxes/)
- Score: 6
- Comments: 0
- Posted: 2026-06-12T16:40:35Z

## Translation

タイトル: SkyPilot サンドボックス: 独自の Kubernetes でエージェント コードを大規模に実行
記事のタイトル: SkyPilot サンドボックス: 独自の Kubernetes でエージェント コードを大規模に実行する |スカイパイロットのブログ
説明: SkyPilot は、すでに所有している Kubernetes クラスター上のサンドボックスで、信頼できない LLM 生成のコードを実行します。単一のクラスターは 50,000 以上のサンドボックスをサポートし、さらにマルチクラスターのサポートもサポートされます。個々は 1 秒以内に起動します。コードを使用すると、ホストされている代替手段のコストの最大 10 分の 1 で済みます。
[切り捨てられた]

記事本文:
SkyPilot サンドボックス: 独自の Kubernetes 上でエージェント コードを大規模に実行 |スカイパイロットのブログ
SkyPilot ブログの事例紹介
SkyPilot サンドボックス: 独自の Kubernetes 上でエージェント コードを大規模に実行
サンドボックスとは何ですか?なぜサンドボックスが必要なのでしょうか?
SkyPilot サンドボックスは独自のインフラ上で実行されます
例: サンドボックス報酬を使用したコード生成モデルの RL トレーニング
独自のクラスターで Modal と競合可能 パフォーマンス: 最初のコマンドまでの速度が速く、クラスターに合わせて拡張可能
すべてのエージェント、コーディング アシスタント、RL パイプラインは最終的に同じ壁にぶつかります。モデルがコードを作成し、今度は誰かがそれを実行する必要があります。現在、ほとんどのチームは、そのコードをホスト型サンドボックス ベンダーに渡し、プロンプト、テスト ケース、モデル出力をクラウドから送信しながら、他の人のマシン上で信頼できないコードを実行するために生のコンピューティングの倍数を支払っています。一方、彼らがすでに運用している Kubernetes クラスターはすぐそこにあり、一度に 50,000 のサンドボックスを実行できます。この投稿は、そのギャップを埋めることに関するものです。SkyPilot Sandboxes (BYOC コード実行レイヤー) と、RL ポストトレーニングの完全な例と、Modal に対する直接ベンチマークを示します。
完全な価格計算は、以下のコストセクションで計算されます。
サンドボックスとは何ですか?なぜサンドボックスが必要なのでしょうか? #
LLM はコードを生成します。それがエージェントであれ、コーディング アシスタントであれ、半分トレーニングされたモデルの出力をスコアリングする RL 報酬ループであれ、ある時点でそのコードを実行する必要がありますが、それを信頼することはできません。永遠にループしたり、メモリを使い果たしたり、ファイルを書き込んだり、プロセスを生成したり、電話をかけようとするものをインポートしたりする可能性があります。それを実行するには使い捨ての隔離された場所が必要で、通常は一度に大量のそれらが必要になります。
現在、それはホスト型サンドボックス ベンダーを利用することを意味します。それは機能しますが、取引は現実です。
コスト。すでに所有しているコンピューティングに加えて、ベンダーのサンドボックスごとの料金を支払います。
プライバシー。コードとデータ

(モデルの出力、テスト ケース、プロンプト) 環境をサードパーティに渡します。
米国以外のユーザーの待ち時間。ベンダーはそれぞれの地域で運営されています。別の場所から通話すると、通話ごとにネットワーク距離税が発生します。
SkyPilot サンドボックスは独自のインフラ上で実行されます #
SkyPilot Sandbox は軽量の独立したポッドで、オンデマンドで作成し、コマンドを実行して破棄し、既存の Kubernetes 上で実行します (BYOC: 独自のクラウドを使用する)。
ポッドごとの分離。各サンドボックスは、専用のイメージ、CPU、メモリを備えた独自のポッドです。不正な動作をするコードはポッドに格納され、完了するとポッドは破棄されます。
すごい平行線。 1 回の呼び出しで多くのサンドボックスを起動し、サンドボックス全体にコマンドを同時に送信します。
温水プールで 1 秒以内に打ち上げられます。プールは、事前にプロビジョニングされたポッドをアイドル状態で準備状態に保つため、サンドボックスを作成すると、Kubernetes のスケジューリングやイメージのプルを待つ代わりに、実行中のポッドを要求します。これにより、1 つのサンドボックスの起動時間が 50% 以上短縮されます。
あなたのインフラ、あなたのデータ。コードとデータがクラウドから離れることはありません。グレーディングに認証情報 (プライベート パッケージ インデックス、統合テスト用のデータベース) が必要な場合、認証情報は作成時に SkyPilot Secrets Manager から挿入され、イメージにベイクされることはありません。
モーダルスタイルのAPI。 create() 、 exec() 、 terminate() 、それぞれに大規模なファンアウトを実現する .aio 上の非同期兄弟が含まれています。ホスト型サンドボックス SDK を使用したことがある場合は、これについてはすでにご存知です。
sky.sandbox をインポートします
sb = 空。サンドボックス。 create (image = "python:3.12"、cpus = 1、memory_gb = 2)
結果 = sb 。 exec ( "python" 、 "-c" 、 "print(2 + 2)" )
print ( result [ "stdout" ]) # "4" (また: stderr、exit_code)
sb 。終了()
# 1 回の呼び出しで、分離されたサンドボックスのリストが返されます。
サンドボックス = 空 。サンドボックス。作成 ( image = "python:3.12" 、 num_sandboxes = 100 )
サンドボックス内の sb の場合:
sb

。 exec ( "pytest" 、 "-q" 、 "tests/" )
# すべてのエントリポイントには、`.aio` 上に非同期の兄弟があります。
サンドボックス = 空を待ちます。サンドボックス。作成する 。 aio ( image = "python:3.12" 、 num_sandboxes = 64 )
results = asyncio を待ちます。集まって（
* ( sb . exec . aio ( "python" , "-c" , code ) サンドボックス内の sb ))
asyncio を待ちます。収集 ( * ( sb . terminate . aio () サンドボックス内の sb 用 ))
例: サンドボックス化された報酬を使用したコード生成モデルの RL トレーニング #
大量の信頼できないコードは、強化学習で最も顕著に現れます。この例では、コード生成 LLM をポストトレーニングします。これは、プログラミング上の問題が与えられた場合に、それを解決する Python 関数を作成するポリシー モデルです。トレーニングの目標は簡単に言うと、モデルの生成された関数がより頻繁にテストに合格するようにすることです。
すべてのトレーニング ステップ、バッチ内のすべてのロールアウトで、半分トレーニングされたモデルが作成したばかりのコード (バグが多く、場合によっては無限ループになり、定義上信頼できない) が実行され、その実行はトレーニングのクリティカル パス上にあります。これは、HuggingFace の Open R1 が RL 報酬にホスト型サンドボックスを使用したときに発生した問題と同じ形です。ここでは、SkyPilot サンドボックスを介して独自の Kubernetes クラスター上で実行が実行されます。
標準的な分散 RL レイアウトを使用します。SkyPilot ジョブ グループ内の 5 つのサービスが HTTP 経由で通信します。
データ サーバーはプロンプト (非表示のテストを含む MBPP スタイルの問題) をロールアウト サーバーに提供します。
ロールアウト サーバー (SGLang) は、現在のポリシーに基づいて候補ソリューションを生成し、報酬サーバーに送信します。
サンドボックス報酬サーバーは各候補者にスコアを付けます。ここでサンドボックスが登場します。受信したバッチごとに、ウォーム プールからサンドボックスのバッチを取得し、独自のサンドボックス内の非表示のテストに対して各候補を実行し、1.0 (すべてのテストに合格) または 0.0 (その他) を返します。
PPO トレーナーはスコアを書き込みます

リプレイ バッファーへのロールアウト。
PPO トレーナー (GRPO) は報酬を使用してポリシーを更新し、ループが繰り返されます。
ロールアウト サーバーはコードを生成し、サンドボックス報酬サーバーは独自のポッドで各候補を実行して報酬を返します。PPO トレーナーはポリシーを更新し、スコア付けされたロールアウトをリプレイ バッファーに保存し、新しい重みをロールアウト サーバーに同期します。
報酬サーバー内。 PPO トレーナーはすでに、{プロンプト、応答、テスト} のバッチを報酬サーバー上の /batch_reward エンドポイントに POST しています。文字列マッチング報酬サーバーからの唯一の変更は、その内部で何が起こるか、つまりコードを実行することです。生成されたスクリプトごとに 1 つのサンドボックスを作成し、作成されたサンドボックスとスクリプトの各ペアに対してスコアリング関数を呼び出します。
非同期をインポートする
sky.sandbox をインポートします
非同期定義スコア_バッチ (項目):
# 1 回の呼び出しで、ウォーム プールから取得されたサンドボックスのリストが返されます。
サンドボックス = 空を待ちます。サンドボックス。作成する 。アイオ（
name = "reward" 、num_sandboxes = len ( items )、pool = POOL_NAME )
試してみてください:
# すべてのロールアウトを同時に 1 つのサンドボックスでスコア付けします。
報酬 = asyncio を待ちます。集まって（
* ( core_one ( sb , item ) for sb , zip 内の item ( Sandboxes , items )))
最後に：
# たとえ幹部が上に上がったとしても、常にサンドボックスを破壊してください。
asyncio を待ちます。収集 ( * ( sb . terminate . aio () サンドボックス内の sb 用 )、
return_Exceptions = True )
リターンリスト（報酬）
1 つのロールアウトをスコアリングすると、実行が行われます。コード ブロックを抽出し、セットアップ コードおよび非表示のテストと連結し、そのスクリプトをサンドボックスで実行します。報酬は可能な限りクリーンなシグナルです。exit 0 はすべてのテストが合格したことを意味し、それ以外のもの (アサーションの失敗、ランタイム エラー、タイムアウトに達した無限ループ、サンドボックス レベルのエラー) は報酬 0.0 です。ルールは、悪いロールアウトで報酬機能を決して引き上げてはいけないということです。

ション;トレーニングの初期段階では、ほとんどのロールアウトはうまくいかず、ループを続けなければなりません。
非同期定義スコア_one ( sb , item ):
コード = extract_code (項目 . 応答 )
code または item でない場合。テスト:
return RewardResponse (報酬 = 0.0、渡された = False)
script = build_test_script (コード、項目.セットアップコード、項目.テスト)
試してみてください:
result = asyncio を待ちます。 wait_for (
sb 。実行。 aio ( "python" 、 "-c" 、 script )、
タイムアウト = EXEC_TIMEOUT_SECONDS )
( asyncio . TimeoutError 、 Exception ) を除く:
# クラッシュやタイムアウトは 0.0 の報酬であり、決して例外ではありません。
# エスケープしてバッチを強制終了します。
return RewardResponse (報酬 = 0.0、渡された = False)
渡された = 結果 [ "exit_code" ] == 0 # stdout / stderr / exit_code
return RewardResponse (合格した場合は報酬 = 1.0、それ以外の場合は 0.0、合格 = 合格)
ウォーム プールはサーバーの起動時に 1 回作成され、共有セッションは停止時に 1 回解放されます。
空。サンドボックス。 create_pool (
名前 = POOL_NAME 、画像 = "python:3.11-slim" 、
CPU = 1、memory_gb = 2、レプリカ = 8)
# ...シャットダウン時:
空を待ってください。サンドボックス。近い（）
この報酬サーバーを文字列一致のものと交換しても、標準 GRPO パイプラインの残りの部分は変更されません。
独自のクラスターで Modal と競合 #
パフォーマンス: 最初のコマンドまでの速度が速く、クラスターに応じて拡張可能 #
私たちは、Modal の米国インフラストラクチャでホストされているマネージド マルチテナント サービスである Modal Sandbox に対して BYOC サンドボックスのベンチマークを行いました (内部ベンチマーク、2026 年 6 月)。 3 つのポイント。
スケールはクラスターによって決まります。単一の Kubernetes クラスターは、220 ノードにわたって約 50,000 の健全なサンドボックスを維持しました。クラスターを追加して上位に移動すると、SkyPilot は、容量のあるクラスターにリクエストをインテリジェントにルーティングします。
最初のコマンドまでの時間が最大 20% 短縮され、テールがよりタイトになります。重要な指標は、コマンドが実行されるまでの時間です。

新しいサンドボックスの n が戻ります。作成してすぐに実行します。 p50 では、SkyPilot サンドボックスは作成と最初の実行を ~1.0 秒であるのに対し、Modal の ~1.2 秒で完了し、尾部はさらに分岐します (p99 ~1.5 秒 vs ~2.0 秒)。 Modal の create() はすぐに戻りますが、まだ準備ができていないハンドルを返します。 readiness は最初の exec に適用され、そこに差異が存在します。 SkyPilot は準備状況を create() にフロントロードするため、最初の実行は迅速かつ予測可能です。
カール -fsSLO https://gist.githubusercontent.com/lloyd-brown/58bdefdea5ff15f1563efa81fbed272a/raw/benchmark.py
Pythonベンチマーク.py
ベンチマーク: プラットフォームあたり 200 回の作成、実行、終了サイクル、create() からエコーがインポート時間を返すまでの経過時間
モーダルをインポートする
試してみてください:
sky.sandbox をインポートします
ImportError を除く:
sky = なし # SkyPilot クライアントはありませんか?ベンチ モーダルのみ
print ( "SkyPilot + モーダルのベンチマーク" if sky else
「SkyPilot クライアントが見つかりません。モーダルのみをベンチマークします。」)
N = 200
アプリ = モーダル 。アプリ 。 lookup ( "ベンチ" 、 create_if_missing = True )
画像 = モーダル。画像 。 debian_slim ( python_version = "3.12" )
空の場合:
# 同等のスリムな Python 3.12 イメージ;ウォーム容量を 1 回事前にプロビジョニングします。
print ( "ウォーム プールを作成しています (1 回限り、期限なし)..." )
空。サンドボックス。 create_pool ( name = "ベンチ" 、画像 = "python:3.12-slim" 、
cpus = 1 、memory_gb = 2 、レプリカ = 5 、ブロッキング = True )
# プラットフォームごとに 1 つの時間制限のないウォームアップ サイクルがあるため、1 回限りのセットアップ (モーダル イメージ)
# 最初の使用時、クライアント セッションの初期化時の解決) は数値に反映されません。
print ( "プラットフォームごとのウォームアップ サイクル (時間制限なし)..." )
msb = モーダル 。サンドボックス。 create ( "スリープ" 、 "無限" 、 app = アプリ 、 画像 = 画像 )
msb 。 exec ( "echo" , "hi" ) 。待ってください()
msb 。終了()
空の場合:
sb = 空。サンドボックス。 create ( name = "ベンチウォームアップ" 、プール = "ベンチ" )
sb 。 exec ( "エコー" 、 "こんにちは" )
sb 。終了()
def pctl ( xs , q ):

並べ替えて返します ( xs )[round ( q / 100 * ( len ( xs ) - 1 ))]
print ( f "タイミング { N } プラットフォームごとに作成 -> 実行 -> サイクルを終了..." )
skypilot_s 、 modal_s = [], []
範囲 ( N ) 内の i の場合:
空の場合:
t0 = 時間。 perf_counter ()
sb = 空。サンドボックス。 create ( name = f "bench- { i } " , pool = "bench" ) # exec-ready
sb 。 exec ( "エコー" 、 "こんにちは" )
スカイパイロット_s 。 append ( time . perf_counter () - t0 )
sb 。終了()
t0 = 時間。 perf_counter ()
msb = モーダル 。サンドボックス。 create ( "スリープ" 、 "無限" 、 app = アプリ 、 画像 = 画像 )
msb 。 exec ( "echo" , "hi" ) 。 wait () # コンテナーの準備がここに到着します
modal_s 。 append ( time .perf_counter ()
[切り捨てられた]
独自のクラスターでは、マシンの料金のみを支払います。ここでは、Modal のコア秒ごとの料金と GiB 秒ごとの料金 (2026 年 6 月に公開されたオンデマンド価格) の両方に対する、独自の数値で再実行できる 2 つの完全に機能した比較を示します。 1 つは汎用ノードでの控えめな比較で、もう 1 つはバースト可能ノードでの 10 倍に近い無駄のない比較です。両方のシナリオ: 単一クラスターが 50,000 のサンドボックスを維持し、フリート全体に対して 1 時間あたりの料金がかかります。
保守的なケース、汎用ノードでは、最大 4 倍のコストがかかります。各サンドボックスには 2 つの vCPU と 4 GB のメモリが割り当てられます。主催:
50,000 x (2 コア x 0 ドル。

[切り捨てられた]

## Original Extract

SkyPilot runs untrusted, LLM-generated code in sandboxes on the Kubernetes clusters you already own. A single cluster sustains 50,000+ sandboxes, with multi-cluster support to go further. Individual launches in less than a second. All at up to one-tenth the cost of hosted alternatives, with your cod
[truncated]

SkyPilot Sandboxes: Run Agent Code on Your Own Kubernetes, at Scale | SkyPilot Blog
SkyPilot Blog Case Studies
SkyPilot Sandboxes: Run Agent Code on Your Own Kubernetes, at Scale
What is a sandbox, and why do you need one?
SkyPilot Sandboxes run on your own infra
Example: RL-training a code-generation model, with sandboxed reward
Competitive with Modal, on your own clusters Performance: faster to first command, scales with your clusters
Every agent, coding assistant, and RL pipeline eventually hits the same wall: the model wrote code, and now someone has to run it. Today, most teams hand that code to a hosted sandbox vendor paying a multiple of raw compute to execute untrusted code on someone else’s machines, while their prompts, test cases, and model outputs leave their cloud. Meanwhile, the Kubernetes cluster they already operate sits right there, capable of running 50,000 sandboxes at once. This post is about closing that gap: SkyPilot Sandboxes, a BYOC code execution layer, with a full RL post-training example and head-to-head benchmarks against Modal.
The full pricing math is worked out in the cost section below.
What is a sandbox, and why do you need one? #
LLMs generate code. Whether it is an agent, a coding assistant, or an RL reward loop scoring the output of a half-trained model, at some point you have to run that code, and you cannot trust it. It can loop forever, exhaust memory, write files, spawn processes, or import something that tries to phone home. You need a disposable, isolated place to run it, and you usually need a lot of them at once.
Today that means reaching for a hosted sandbox vendor. It works, but the trade is real:
Cost. You pay the vendor’s per-sandbox rate on top of the compute you already own.
Privacy. Your code and data (the model’s output, your test cases, your prompts) leave your environment for a third party.
Latency for non-US users. The vendor runs in their regions. Reach them from somewhere else and every call pays a network-distance tax.
SkyPilot Sandboxes run on your own infra #
A SkyPilot Sandbox is a lightweight, isolated pod you create on demand, run commands in, and tear down, running on the Kubernetes you already have (BYOC: bring your own cloud).
Per-pod isolation. Each sandbox is its own pod with a dedicated image, CPU, and memory. Code that misbehaves is contained to its pod, and the pod is destroyed when you are done.
Massively parallel. Launch many sandboxes in a single call and fan commands out across them concurrently.
Sub-second launches with warm pools. A pool keeps pre-provisioned pods idle and ready, so creating a sandbox claims a running pod instead of waiting on Kubernetes scheduling and an image pull. That cuts a single sandbox’s launch time by more than 50%.
Your infra, your data. Code and data never leave your cloud. If grading needs credentials (a private package index, a database for integration tests), they are injected from the SkyPilot Secrets Manager at create time, never baked into an image.
Modal-style API. create() , exec() , terminate() , each with an async sibling on .aio for massive fan-out. If you have used a hosted sandbox SDK, you already know this one.
import sky.sandbox
sb = sky . sandbox . create ( image = "python:3.12" , cpus = 1 , memory_gb = 2 )
result = sb . exec ( "python" , "-c" , "print(2 + 2)" )
print ( result [ "stdout" ]) # "4" (also: stderr, exit_code)
sb . terminate ()
# One call returns a LIST of isolated sandboxes.
sandboxes = sky . sandbox . create ( image = "python:3.12" , num_sandboxes = 100 )
for sb in sandboxes :
sb . exec ( "pytest" , "-q" , "tests/" )
# Every entrypoint has an async sibling on `.aio`.
sandboxes = await sky . sandbox . create . aio ( image = "python:3.12" , num_sandboxes = 64 )
results = await asyncio . gather (
* ( sb . exec . aio ( "python" , "-c" , code ) for sb in sandboxes ))
await asyncio . gather ( * ( sb . terminate . aio () for sb in sandboxes ))
Example: RL-training a code-generation model, with sandboxed reward #
Untrusted code at volume shows up most sharply in reinforcement learning. This example post-trains a code-generation LLM, a policy model that, given a programming problem, writes a Python function to solve it. The training goal is simple to state: make the model’s generated functions pass the tests more often.
On every training step, for every rollout in the batch, we execute code that a half-trained model just wrote (buggy, occasionally infinite-looping, untrusted by definition) and that execution sits on the critical path of training. This is the same shape of problem HuggingFace’s Open R1 hit when they used hosted sandboxes for their RL reward; here, the execution runs on your own Kubernetes cluster via SkyPilot Sandboxes.
We use a standard distributed RL layout: five services in a SkyPilot job group, talking over HTTP.
The Data Server serves prompts (MBPP-style problems with hidden tests) to the Rollout Server
The Rollout Server (SGLang) has the current policy generate candidate solutions and sends them to the reward server.
The Sandbox Reward Server scores each candidate. This is where sandboxes come in. For every batch it receives, it claims a batch of sandboxes from a warm pool, runs each candidate against its hidden tests in its own sandbox, and returns 1.0 (all tests passed) or 0.0 (anything else).
The PPO trainer writes the scored rollouts to the Replay Buffer.
The PPO Trainer (GRPO) uses the rewards to update the policy, and the loop repeats.
The rollout server generates code, the sandbox reward server runs each candidate in its own pod and returns a reward, and the PPO trainer updates the policy, stores scored rollouts in the replay buffer, and syncs new weights back to the rollout server.
Inside the reward server. The PPO trainer already POSTs a batch of {prompt, response, tests} to the /batch_reward endpoint on the reward server. The only change from a string-matching reward server is what happens inside it: we run code. We create one sandbox for each of the generated scripts and call the scoring function on each pair of created sandbox and script:
import asyncio
import sky.sandbox
async def score_batch ( items ):
# One call returns a LIST of sandboxes, claimed from the warm pool.
sandboxes = await sky . sandbox . create . aio (
name = "reward" , num_sandboxes = len ( items ), pool = POOL_NAME )
try :
# Score every rollout concurrently, one sandbox each.
rewards = await asyncio . gather (
* ( score_one ( sb , item ) for sb , item in zip ( sandboxes , items )))
finally :
# ALWAYS tear sandboxes down, even if an exec raised above.
await asyncio . gather ( * ( sb . terminate . aio () for sb in sandboxes ),
return_exceptions = True )
return list ( rewards )
Scoring one rollout is where execution happens. We extract the code block, concatenate it with the setup code and hidden tests, and run that script in the sandbox. The reward is the cleanest possible signal: exit 0 means every test passed, and anything else (an assertion failure, a runtime error, an infinite loop that hits the timeout, a sandbox-level error) is reward 0.0. The rule is that a bad rollout must never raise out of the reward function; early in training, most rollouts are bad, and the loop has to keep going.
async def score_one ( sb , item ):
code = extract_code ( item . response )
if not code or not item . tests :
return RewardResponse ( reward = 0.0 , passed = False )
script = build_test_script ( code , item . setup_code , item . tests )
try :
result = await asyncio . wait_for (
sb . exec . aio ( "python" , "-c" , script ),
timeout = EXEC_TIMEOUT_SECONDS )
except ( asyncio . TimeoutError , Exception ):
# A crash or a timeout is a 0.0 reward, never an exception that
# escapes and kills the batch.
return RewardResponse ( reward = 0.0 , passed = False )
passed = result [ "exit_code" ] == 0 # stdout / stderr / exit_code
return RewardResponse ( reward = 1.0 if passed else 0.0 , passed = passed )
The warm pool is created once when the server starts, and the shared session is released once when it stops:
sky . sandbox . create_pool (
name = POOL_NAME , image = "python:3.11-slim" ,
cpus = 1 , memory_gb = 2 , replicas = 8 )
# ... on shutdown:
await sky . sandbox . aclose ()
Swap this reward server in for a string-matching one and the rest of a standard GRPO pipeline does not change.
Competitive with Modal, on your own clusters #
Performance: faster to first command, scales with your clusters #
We benchmarked BYOC sandboxes against Modal Sandboxes , a managed, multi-tenant service hosted in Modal’s US infrastructure (internal benchmarks, June 2026). Three takeaways.
Scale is determined by your cluster. A single Kubernetes cluster sustained ~50,000 healthy sandboxes across 220 nodes. Add clusters to go higher and SkyPilot will intelligently route requests to clusters with capacity.
Time to first command is ~20% faster, with a much tighter tail. The metric that matters is how long until a command you run in a fresh sandbox comes back: create, then immediately exec. At p50, a SkyPilot sandbox completes create + first exec in ~1.0s vs Modal’s ~1.2s , and the tails diverge further (p99 ~1.5s vs ~2.0s). Modal’s create() returns quickly but hands back a not-yet-ready handle; readiness lands on the first exec, which is where its variance lives. SkyPilot front-loads readiness into create() , so the first exec is quick and predictable.
curl -fsSLO https://gist.githubusercontent.com/lloyd-brown/58bdefdea5ff15f1563efa81fbed272a/raw/benchmark.py
python benchmark.py
The benchmark: 200 create, exec, terminate cycles per platform, wall time from create() until an echo returns import time
import modal
try :
import sky.sandbox
except ImportError :
sky = None # no SkyPilot client? bench Modal only
print ( "Benchmarking SkyPilot + Modal" if sky else
"No SkyPilot client found; benchmarking Modal only" )
N = 200
app = modal . App . lookup ( "bench" , create_if_missing = True )
image = modal . Image . debian_slim ( python_version = "3.12" )
if sky :
# Comparable slim Python 3.12 image; pre-provision warm capacity once.
print ( "Creating warm pool (one-time, untimed)..." )
sky . sandbox . create_pool ( name = "bench" , image = "python:3.12-slim" ,
cpus = 1 , memory_gb = 2 , replicas = 5 , blocking = True )
# One untimed warmup cycle per platform, so one-time setup (Modal image
# resolution on first use, client session init) never lands in the numbers.
print ( "Warmup cycle per platform (untimed)..." )
msb = modal . Sandbox . create ( "sleep" , "infinity" , app = app , image = image )
msb . exec ( "echo" , "hi" ) . wait ()
msb . terminate ()
if sky :
sb = sky . sandbox . create ( name = "bench-warmup" , pool = "bench" )
sb . exec ( "echo" , "hi" )
sb . terminate ()
def pctl ( xs , q ):
return sorted ( xs )[ round ( q / 100 * ( len ( xs ) - 1 ))]
print ( f "Timing { N } create -> exec -> terminate cycles per platform..." )
skypilot_s , modal_s = [], []
for i in range ( N ):
if sky :
t0 = time . perf_counter ()
sb = sky . sandbox . create ( name = f "bench- { i } " , pool = "bench" ) # exec-ready
sb . exec ( "echo" , "hi" )
skypilot_s . append ( time . perf_counter () - t0 )
sb . terminate ()
t0 = time . perf_counter ()
msb = modal . Sandbox . create ( "sleep" , "infinity" , app = app , image = image )
msb . exec ( "echo" , "hi" ) . wait () # container readiness lands here
modal_s . append ( time . perf_counter ()
[truncated]
On your own cluster you pay only for the machines. Here are two fully worked comparisons you can rerun with your own numbers, both against Modal’s per-core-second and per-GiB-second billing (published on-demand pricing, June 2026): a conservative one on general-purpose nodes, and a leaner one on burstable nodes that approaches 10x. The scenario for both: the 50,000 sandboxes a single cluster sustains, priced per hour for the whole fleet.
The conservative case, on general-purpose nodes: ~4x cheaper. Each sandbox gets 2 vCPUs and 4 GB of memory. Hosted:
50,000 x (2 cores x $0.

[truncated]
