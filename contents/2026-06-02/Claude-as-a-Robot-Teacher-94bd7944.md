---
source: "https://irvin.quest/distilling-zero-shot-control/"
hn_url: "https://news.ycombinator.com/item?id=48357704"
title: "Claude as a Robot Teacher"
article_title: "Claude as a Robot Teacher | Irvin Hwang"
author: "evakhoury"
captured_at: "2026-06-02T00:38:16Z"
capture_tool: "hn-digest"
hn_id: 48357704
score: 2
comments: 0
posted_at: "2026-06-01T14:59:52Z"
tags:
  - hacker-news
  - translated
---

# Claude as a Robot Teacher

- HN: [48357704](https://news.ycombinator.com/item?id=48357704)
- Source: [irvin.quest](https://irvin.quest/distilling-zero-shot-control/)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T14:59:52Z

## Translation

タイトル: ロボット教師としてのクロード
記事のタイトル: ロボット教師としてのクロード |アービン・ファン
説明: クロードとワールド モデルを抽出して、より高速なゼロショット ロボット コントローラーを作成する

記事本文:
ロボット教師としてのクロード |アービン・ファン ロボット教師としてのアービン・ファン・クロード
前回の私の目標は、ロボット アームのビュー内でオブジェクトを中心に置くという制御タスクにワールド モデルを使用することでした。結果は次のとおりです。
これを機能させるには 2 つの段階がありました。 1 つ目は、クロード ベースのコントローラーを作成することでした。これは、ワールド モデルを使用して各候補の動き (左にパン、右にパン、またはホールド) の結果を予測し、クロードにレビューしてどのアクションがより良いかを決定させることで機能しました。これは非常に遅かったので、第 2 段階では、このクロード コントローラーを使用してデータを生成し、そのデータを使用して、ビデオで見られるより高速なコントローラーをトレーニングしました。その過程で、私はロボットのトレーニングプロセスの多くを AI に制御させる実験をすることになりましたが、それはあまり成功しませんでした。
私の元のワークフローは、データ収集とモデル トレーニングの 2 つの別々のプロセスで構成されていました。私は、これらのプロセスを組み合わせて、いつそれらの間を行き来するかをクロードに決定させることを実験してみることにしました。高レベルでは、自律学習器と呼ぶ統合プロセスには 5 つの状態があります。データの収集、検証 (新しいデータでモデルをテスト)、トレーニング、思考、人間の介入のための一時停止 (シーン内でオブジェクトを手動で移動する) です。自律学習器は主にオーケストレーション用であり、既存のデータ収集およびトレーニング リポジトリからのコードを利用して、データの収集、トレーニング、および状態の検証を行います。思考状態はこのプロセスの新しい部分であり、基本的に他の状態からのデータ (学習曲線、収集されたデータに関する情報、構成パラメーターなど) をクロードに求めるプロンプトです。クロードはこのデータをレビューし、自律学習者にとって次の状態がどうあるべきか、また行うべき調整を決定します。チャンギ

モデルのアーキテクチャ、微調整とゼロからのトレーニング、またはどのくらいの量のデータを収集し、どの範囲の動作を行うかなどです。自律学習器により、データ収集とトレーニングのプロセス全体が容易になりましたが、クロードによる意思決定は、前回の投稿でトレーニング プロセスを自動化したのと同じように、より良いモデルにはつながりませんでした。クロードはしばしば合理的な分析を行います。データに多様性が欠如していると判断しても、人間が場面を再配置するために学習者に一時停止を指示することはできません。データ収集とトレーニングをループに入れることで利便性が高まり、データ収集とトレーニングの間の移行がよりスクリプト化されるようになりました。最終モデルは、約 10 億のパラメータを持つ拡散トランスでした。これには、トレーニング データ内の事実に反するアクションを予測する優れた機能がありました。
推論をライブで実行すると、アクションの条件付けが一部明らかになりましたが、結果はそれほど良好ではありませんでした。
当初の計画は、以前のプロジェクトで行ったのと同様に、ビジョン言語モデル (VLM) をワールド モデルと組み合わせてロボットを制御することでした。このアイデアは、ワールド モデルを使用してさまざまなアクションに基づいて予測フレームを取得し、VLM を使用して目的の最終状態に最も近い予測を選択することでした。前回のようにデータを手動で収集して VLM を微調整したくなかったので、Gemma や Qwen のさまざまなバージョンなど、いくつかの異なるローカル モデルを試しました。どれもそれほど良好なパフォーマンスを発揮しなかったため、最終的には、Claude Code CLI ( claude -p --model opus --effort max ) を介して、最大エフォートを使用して Claude への呼び出しを試みることになりました。これは最初に試したときはうまくいきましたが、非常に遅かったです (アクションごとに約 25 ～ 50 秒)。最近、ライブ推論のために再試行しましたが、うまくいきませんでした。

これは、照明の変化やその他の配分外の要因が原因である可能性があります。
私は、大規模なモデルをより効率的な小さなモデルに蒸留できるという漠然とした考えを持っていたため、Claude と相談した後、動作の複製を通じて高速モデルをトレーニングしてみることにしました。このアイデアは、トレーニング データのフレームを使用し、それをクロード ロボット コントローラーにフィードすることでした。内部では、これはスロー コントローラーと同じ予測してからランク付けするループを再利用します。ワールド モデルは、すべての候補アクションの下で各サンプリングされたフレームを前方にロールし、クロードはコングを最も中心に置く予測を選択します。これにより、単純な ResNet18 + 多層パーセプトロン (~11M パラメーター)、画像からアクションへの分類器をトレーニングするために使用される新しいデータセットが作成されました。これは最終的にうまく機能し、アクションまでの時間が約 25 秒から約 1 秒に短縮されました。現在は、計算ではなく腕が物理的に動くことによって制限されています。また、ライブではワールド モデルよりもはるかにうまく機能しました。これは、より高速なモデルを教えるために、ワールド モデルをあまり一般化する必要がないことを意味しているのかもしれないので、興味深いです。また、ワールド モデルに事前トレーニングされたバックボーンを使用してみる必要があるということも意味するかもしれません。
ここから進むべき方向性はたくさんありますが、中長期的な主な目標は、ある種のピックアンドプレースタスクを実行し、短期的には物体に触れる腕を訓練すること、つまりプロセスを複数の関節に一般化することだと思います。
完璧なツールを目指す曲がりくねった道。
YouTube GitHub にメールで送信する

## Original Extract

Distilling Claude and a world model into a fast(er) zero-shot robot controller

Claude as a Robot Teacher | Irvin Hwang Irvin Hwang Claude as a Robot Teacher
My goal last time was to use the world model for the control task of centering an object in the view of the robot arm. Here’s the result.
There were two stages to get this to work. The first was making a Claude based controller, which worked by using the world model to predict the outcome of each candidate move (pan left, pan right, or hold) then having Claude review and decide which action was better. This was really slow so the second stage was to use this Claude controller to generate data that was then used to train a faster controller that you see in the video. Along the way I ended up experimenting with letting AI control more of the robot training process, although that was somewhat less successful.
My original workflow consisted of two separate processes of data collection and model training . I decided to experiment with combining these processes and having Claude make decisions on when to move back and forth between them. At a high level the unified process, which I’ll call the autonomous learner , has five states: collect data, validate (test the model on new data), train, think, and pause for human intervention (manually move objects around in the scene). The autonomous learner is mainly for orchestration and utilizes code from the existing data collection and training repos for the collect data, train, and validate states. The think state is the new part of this process and is essentially a prompt to Claude with data from the other states (e.g. learning curve, info on data collected, configuration parameters, etc). Claude reviews this data and then decides what the next state should be for the autonomous learner and any adjustments that should be made e.g. changing the architecture of the model, fine-tuning vs training from scratch, or how much data to collect and over what range of motion. The autonomous learner did make the entire process of collecting data and training easier, but the decision making by Claude didn’t lead to better models in the same way automating the training process did in the last post. Claude would often make a reasonable analysis e.g. decide there was a lack of diversity in data, but then would fail to direct the learner to pause for a human to rearrange the scene. The convenience came more from putting data collection and training into a loop, and the transitions between data collection and training ended up being more scripted. The final model was a diffusion transformer with about 1 billion parameters. It had a good ability to predict counterfactual actions within the training data.
Some action conditioning was evident when running inference live, but the results were not nearly as good.
The initial plan was to use a vision language model (VLM) in conjunction with the world model to control the robot similar to what I did in my earlier project . The idea was to use the world model to get predicted frames based on different actions then use the VLM to choose the prediction that was closest to the desired end state. I didn’t want to manually collect data to fine tune the VLM like last time so I tried a few different local models like various versions of Gemma and Qwen. None of them performed that well so I ended up trying calls to Claude using max effort — via the Claude Code CLI ( claude -p --model opus --effort max ). This worked when I initially tried it, but was extremely slow (about 25-50 seconds per action). It’s worth noting I recently tried it again for live inference and it did not work well, which may be due to changes in lighting or other out of distribution factors.
I had a vague notion that large models could be distilled into smaller more efficient models so after some consultation with Claude I decided to try training a fast model through behavior cloning . The idea was to use frames from the training data and feed them into the Claude robot controller. Under the hood this reuses the same predict-then-rank loop as the slow controller — the world model rolls each sampled frame forward under every candidate action and Claude picks the prediction that best centers the Kong. This created a new dataset that was used to train a simple ResNet18 + multi layer perceptron (~11M parameters), image-to-action classifier. This ended up working well and cut the time to action from ~25 seconds to about a second, now limited by the arm physically moving rather than compute. It also worked much better live than the world model, which is interesting because maybe it implies we don’t need to get the world model to generalize as much to still be useful for teaching the faster model. It could also mean I should try using a pre-trained backbone for the world model.
There are a lot of potential directions to go from here, but I think the main medium to longterm goal is to perform some sort of pick and place task and in the shorter term training the arm to touch an object i.e. generalize the process to multiple joints.
A meandering path towards the perfect tool.
Email YouTube GitHub
