---
source: "https://meffmadd.github.io/samplesurium/posts/baba_is_agent/"
hn_url: "https://news.ycombinator.com/item?id=48355346"
title: "Can LLMs Play Baba Is You?"
article_title: "Baba is Agent – Samplesurium"
author: "stared"
captured_at: "2026-06-02T00:41:54Z"
capture_tool: "hn-digest"
hn_id: 48355346
score: 2
comments: 0
posted_at: "2026-06-01T11:18:47Z"
tags:
  - hacker-news
  - translated
---

# Can LLMs Play Baba Is You?

- HN: [48355346](https://news.ycombinator.com/item?id=48355346)
- Source: [meffmadd.github.io](https://meffmadd.github.io/samplesurium/posts/baba_is_agent/)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T11:18:47Z

## Translation

タイトル: LLM はババをプレイできますか?
記事タイトル: ババはエージェント – サンプルスリウム

記事本文:
それを調べるために OpenCode ベースのエージェントを構築しました。 GitHub リポジトリのコードをチェックしてください。
Baba is You はパズル ゲームで、プレイヤーはグリッドベースのレベルをナビゲートし、テキスト ルールを操作して最終的に勝利状態に到達する必要があります。ルールは、壁、旗、岩、またはあなた自身など、グリッド上のエンティティに影響します。
OpenCode は、エージェントが使用できる独自のツールを作成するなど、拡張可能なオープンソースの CLI コーディング エージェントです。私は素晴らしい baba_is_eval リポジトリに触発され、MCP サーバーであるその既存のコードを OpenCode で動作するように調整しました。既存のリポジトリからコア ゲームプレイを処理するツールを更新することに加えて、ユーティリティ ツールとして game_insights と Shortest_path を追加しました。エージェントがアクセスできたツールのリストは次のとおりです。
get_game_state : 現在のゲーム状態を文字列の 2D 配列 (grid) またはエンティティの JSON リスト (entities) として取得します。
execute_game_commands : 移動コマンドのリストを実行します (上、下、左、右、アイドル)
restart_level : 現在のレベルを再開します
game_insights : アクティブなルール、あなた/勝利ポジション、および勝利ポジションへの最短パス (該当する場合) を表示します。
Shortest_path : ブロックされたエンティティ (停止ルールまたは敗北ルールを持つエンティティなど) を回避して、ターゲット位置までの経路探索を行う A*
undo_multiple : 最後の N 回の動作を元に戻します
todowrite : タスクの進行状況を追跡する (OpenCode ネイティブ)
get_game_state ツールの 2 つの形式は、ARC AGI 3 [1] と Nicolas Martorell による論文 [2] からインスピレーションを得ています。 ARC AGI 3 は、ゲームの状態を 3D 配列 (1 つの追加の時間次元) として返します。これは、ここではグリッド形式に対応します。エンティティの形式は、論文 [2] のデカルト JSON 表記に基づいており、評価では非常に優れた結果を示しました。
各評価実行のタイムアウトは、200,000 トークンのしきい値で 20 分に設定されました。モデルは決して到達しませんでした

e 200,000 トークンのしきい値であり、タイムアウトになるのは唯一です。
評価はレベルおよびモデルごとに 1 回だけ実行されました (時間とコストの制約)
評価用に AGENTS.md ファイルを削除するのを忘れました (レベル ソルブのコンテキスト肥大化)
統計は、OpenCode JSON ログによって提供される情報に完全に依存しています。
OpenCode Go/Zen がモデル プロバイダーとして使用されました。
次の表は、モデルの全体的な結果を示しています。オープンフロンティアモデルとクローズドフロンティアモデルの間では、モデルの能力に多少の差異が見られます。 Gemini 3.1 Pro はすべてのレベルを解決できましたが、オープンウェイト モデルで最も優れているのは GLM 5.1 で、8 つのレベルのうち 5 つをクリアしました。
次のセクションでは、各レベルの結果をもう少し詳しく見ていきたいと思います。モデルは、レベルを解決するのにかかる時間と、それに必要なトークン/ツールの呼び出しの両方においてかなり異なります。ただし、モデルごとにレベルごとに 1 回の試行だけでは、最終的な結論を引き出すことはできません。
レベル 0 は入門レベルであり、結果的には簡単です。解決策は、岩を押すことができるので、右に8回移動するだけです。
レベルが単純なため、すべてのモデルが問題を解決できました。 Qwen 3.6 Plus は、レベルを解くのに必要な時間の点で少し外れ値です。また、他のすべてのモデルでは 2 ～ 3 回しか必要としなかったのに対し、8 回のツール呼び出しも行われました。
このレベルでは、レベルを解決するための 2 つの簡単なオプションがあります。 FLAG IS WIN を構築するか、もう少しトリッキーな解決策として WALL IS WIN を構築します。
このレベルは全体的な結果を少し予見するものでした。 Gemini 3.1 Pro が明らかに最速で、次にクローズド フロンティア モデルが続きました。 Minimax M2.7 はすでにタイムアウトしており、他のオープン モデルはクローズド フロンティア モデルと比較して大幅に長い時間を必要としていました。
レベル 2 は、概念的には同じレイアウトを持つレベル 1 と非常に似ていますが、

ティティが変更されました (例: FLAG が壁として機能するようになりました)。現在、レベルを解決するオプションは 1 つだけあり、それは FLAG IS WIN を構築することです。これはレベル 1 の 2 番目のオプションに対応します。
モデルはこのレベルを解くのにもう少し時間がかかるため、レベル 1 に比べて難易度は確かに少し高いようです。
このレベルでは、旗は赤ニシンであり、右下隅にある旗と岩のルールを操作して、 ROCK IS WIN を形成する必要があります。こちらもシンクメカを使うレベルです。最初の岩を水ブロックに押し付けて破壊し、通過できるようにする必要があります。
ここでは、タイムアウトせずに解決が初めて失敗した新しいケースがわかります。 GLM 5.1 と Kim K2.6 は 32,000 を超える推論トークンを生成しましたが、評価が途中で終了しました。 DeepSeek V4 Pro と Qwen 3.6 Plus が制限時間に達し、タイムアウトになりました。これは、閉じられたフロンティア モデルのみがこのレベルを解決できたことを意味します。
このレベルでは、初めて DEFEAT ルールに遭遇します。勝利条件で旗を囲むドクロバリアを回避するには、無効化する必要があります。これは、岩を使用するか、ROCK IS PUSH ルール自体のテキストを使用して SKULL IS DEFEAT ルールを破ることによって実行できます。
GLM 5.1 はこのレベルを解決できましたが、Claude Opus 4.7 は制限時間に達しました。 Gemini 3.1 Pro と GPT 5.5 は、再び非常に迅速かつ効率的にレベルを解くことができました。
レベル 5 では、LAVA をプッシュ可能にするための創造的なルール操作が必要で、これにより HOT および MELT 制約がバイパスされます。
Gemini 3.1 Pro と GPT 5.5 の課題は、このレベルでは以前のレベルと比べて解決時間がほぼ 2 倍になったため、間違いなくワンランク上になりました。 GLM 5.1 および Claude Opus 4.7 で再びエラーが発生し、ステップごとの出力トークンの最大数を超えました。残りのモデルはできませんでした

o 20 分の時間枠内でレベルを解決してください。
このレベルでも、創造的なルール操作が必要になります。 SKULL IS DEFEAT は無効化できないため、スカルバリアを回避する方法を見つける必要があります。これを達成するには、壁ブロックの一部がすでにバリアを超えているため、「WALL IS YOU」ルールを作成する必要があります。そうすれば旗に到達して勝つことができます。
Claude Opus 4.7 は、20 分以内にこのレベルをかろうじて解くことができました。 GPT 5.5 と Gemini 3.1 Pro は、再び非常に早くレベルを解くことができました。このレベルでは、キミ K2.6 がタイムアウト直前に試行錯誤の末、自主的にギブアップしました。 GLM 5.1 は出力トークンの制限に達しました。このレベルを解決できたのは、閉鎖フロンティア モデルだけでした。
レベル 7 は概念的には非常にシンプルですが、いくつか注意点があります。まず、壁は完全に無視して、そのエリアが囲まれているという錯覚を生み出すだけです。実際には、それらはまったく障害ではありません。一方、草は FLAG ブロックと WIN ブロックを囲み、一種の迷路を作成します。テキスト ブロックを草の迷路から移動し、BABA IS YOU ルールを再利用して FLAG IS WIN ルールを作成し、IS を共有する必要があります。また、レベルは動作が重いため、正確な位置決めが必要になります。
Gemini 3.1 Pro だけがこのレベルを正常に解決しました。 GPT 5.5 はこのレベルで自発的に放棄しましたが、GLM 5.1 は出力制限に達しました。他のすべてのモデルは 20 分のタイムアウトになりました。
評価を実行した後、信号が非常にクリアであることに驚きました。他の「ブティック」ベンチマークと同様に、これらのよりニッチなタスクでは、オープンウェイト モデルとクローズド フロンティア モデルの間には明確な違いがありました。
Minimax M2.7 は明らかに小さすぎて、重要なレベルを解決できませんでした。 GLM 5.1 と Claude Opus 4.7 は非常に僅差であり、それぞれ異なるレベルで優勝しました。 GLM 5.1 も同様でした

レベル 2 を超えるレベルを解決するには、オープン ウェイト モデルのみを使用します。Gemini 3.1 Pro と GPT 5.5 が最高のパフォーマンスを発揮しているようです。確かにそれらの方が効率的です。 ARC AGI 3 のリリースにより、モデルはこのタスクに関しても改善される可能性があります。ここでテストされたクローズド フロンティア モデルは、ARC AGI 3 リーダーボードにほとんど影響を与えず、2026 年 5 月 10 日の時点で 0.1 ～ 0.4% のスコアしか得られません。
私は今、このような何かが閉じたモデルの訓練データに現れるのか、それとも単に創発的な動作なのか疑問に思っています。おそらく、オープン ラボはそのギャップを埋めるために純粋なコーディング ハーネスに重点を置く一方、フロンティア ラボはより広範なコンテキストでモデルをトレーニングします。少なくともGoogleはそのようだ。
以下は、全レベルにわたる各モデルのコストと全体の合計です。
以下の表は、すべての実行において、各モデルが format=entities と format=grid を使用して get_game_state を呼び出す頻度を示しています。
以下は、すべての実行における全体のツール呼び出し数 (モデルごとの合計ツール呼び出しの割合) です。

## Original Extract

I built an OpenCode-based agent to find out! Check out the code in the GitHub repo !
Baba is You is a puzzle game where the player has to navigate grid-based levels and manipulate textual rules to get to an eventual win state. The rules affect entities on the grid such as walls, flags, rocks or yourself.
OpenCode is an open-source CLI coding-agent that is extensible, including writing your own tools the agent can use. I was inspired by the fantastic baba_is_eval repo and adapted its existing code, which was an MCP server, to work with OpenCode. On top of updating the tools handling the core gameplay from the existing repo, I added game_insights and shortest_path as utility tools. Here is the list of tools the agent had access to:
get_game_state : Get current game state either as a 2D array of strings ( grid ) or a JSON list of entities ( entities )
execute_game_commands : Execute a list of movement commands (up, down, left, right, idle)
restart_level : Restart the current level
game_insights : Shows the active rules, you/win positions, and the shortest path to a win position (if applicable)
shortest_path : A* pathfinding to target position, avoiding blocked entities (e.g. entities with stop or defeat rules)
undo_multiple : Undo last N moves
todowrite : Track task progress (OpenCode native)
The two formats of the get_game_state tool are inspired by ARC AGI 3 [1] and a paper by Nicolas Martorell [2] . ARC AGI 3 returns the game state as a 3D array (one extra temporal dimension), which corresponds to the grid format here. The entities format is based on the Cartesian JSON notation from the paper [2] , which performed very well in their evaluation.
The timeout for each evaluation run was set to 20 minutes with a 200,000 token threshold. Models never reached the 200,000 token threshold and only ever timed-out.
Evaluations were run only once per level and model (time/money constraints)
I forgot to remove the AGENTS.md file for evaluation (context bloat for level solve)
Stats entirely rely on information provided by the OpenCode JSON logs
OpenCode Go/Zen was used as a model provider.
The following table shows the overall results of the models. We see some variance in model capability across open and closed frontier models. Gemini 3.1 Pro was able to solve all levels while the best open-weights model is GLM 5.1 with 5 out of 8 levels beaten.
In the following sections I want to quickly go over the results for each level in a bit more detail. The models differed quite a bit in both the time to solve the level and the token/tool calls required to do so. However, it is not possible to draw any definitive conclusion with only a single attempt per level per model.
Level 0 is the introductory level and is trivial as a result. The solution is just moving right 8 times, since the rocks can be pushed.
Due to the level’s simplicity all models were able to solve it. Qwen 3.6 Plus is a bit of an outlier here in terms of the time required to solve the level. It also made 8 tool calls, whereas all other models required only 2-3.
For this level there are two straightforward options to solve the level. Either construct FLAG IS WIN or the slightly more tricky solution is to construct WALL IS WIN .
This level was a bit of foreshadowing of the overall results. Gemini 3.1 Pro was clearly the fastest followed by the closed frontier models. Minimax M2.7 already timed-out and the other open models needed significantly more time compared to the closed frontier models.
Level 2 is conceptually very similar to level 1 with the same layout, however the entities are now permuted (e.g. FLAG now acts as a wall). There is now only one option to solve the level, namely to construct FLAG IS WIN . This corresponds to the second option from level 1.
Models needed a bit more time to solve this level so the difficulty seems to indeed be a bit higher compared to level 1.
In this level, the FLAG is a red herring and the flag and rock rules in the lower right corner have to be manipulated to form ROCK IS WIN . This is also a level where the sink mechanic is used. The first rock has to be pushed against a water block to destroy it and be able to pass.
Here we see a new case where a solve failed for the first time without timing out. GLM 5.1 and Kimi K2.6 produced more than 32k reasoning tokens, which prematurely exits the evaluation. DeepSeek V4 Pro and Qwen 3.6 Plus reached the time limit and timed out. This means only the closed frontier models were able to solve this level.
In this level, we encounter the DEFEAT rule for the first time. It has to be deactivated to bypass the skull barrier surrounding the flag with the win condition. This can be done by using the rocks or the text of the ROCK IS PUSH rule itself to break the SKULL IS DEFEAT rule.
GLM 5.1 was able to solve this level while Claude Opus 4.7 reached the time limit. Gemini 3.1 Pro and GPT 5.5 were able to solve the level very quickly and efficiently again.
Level 5 requires creative rule manipulation to make the LAVA pushable, thus bypassing the HOT and MELT constraints.
The challenge for Gemini 3.1 Pro and GPT 5.5 was definitely turned up a notch as the solving times almost doubled for this level compared to the previous one. GLM 5.1 and Claude Opus 4.7 encountered an error again, exceeding the maximum number of output tokens per step. The remaining models were unable to solve the level within the 20 minutes time frame.
This level again requires creative rule manipulation. Since SKULL IS DEFEAT cannot be deactivated, we have to find a way to bypass the skull barrier. To achieve this, we have to create a WALL IS YOU rule, since some wall blocks are already past the barrier. This way we can reach the flag and win.
Claude Opus 4.7 was barely able to solve this level within the 20 minutes. GPT 5.5 and Gemini 3.1 Pro were able to solve the level very quickly again. In this level, Kimi K2.6 gave up voluntarily after extensive trial and error just before the timeout. GLM 5.1 hit the output token limit. Only the closed frontier models were able to solve this level.
Level 7 is conceptually very simple but has a few gotchas. Firstly, the walls can be ignored completely and only create the illusion that the area is enclosed. In reality they are no obstacle at all. Meanwhile the grass encloses the FLAG and WIN blocks and creates a sort of labyrinth. The text blocks have to be moved out of the grass labyrinth and then the BABA IS YOU rule has to be reused to create the FLAG IS WIN rule, sharing the IS . The level is also movement heavy, requiring a lot of precise positioning.
Only Gemini 3.1 Pro solved this level successfully. GPT 5.5 gave up voluntarily in this level, while GLM 5.1 hit the output limit. All other models ran into the 20 minute timeout.
After running the evaluations, I was surprised how clear the signal was. As with some other “boutique” benchmarks, there was a clear distinction between open-weights models and closed frontier models on these more niche tasks.
Minimax M2.7 was clearly too small and could not solve any non-trivial level. GLM 5.1 and Claude Opus 4.7 are quite close and won different levels each. GLM 5.1 was also the only open weights model to solve a level past level 2. Gemini 3.1 Pro and GPT 5.5 seem to be performing best. They are certainly more efficient. With the release of ARC AGI 3, models will likely improve for this task as well. The closed frontier models tested here barely make a dent in the ARC AGI 3 leaderboard and only score 0.1-0.4% as of May 10th 2026.
I now wonder if anything like this shows up in the training data of the closed models or if it is just emergent behavior. Maybe open labs do focus more heavily on pure coding harnesses to close the gap there, while frontier labs train their models in broader contexts. Google seems like that at least.
Below is the cost of each model across all levels with the overall total.
The table below shows how often each model called get_game_state with format=entities vs format=grid across all runs.
Below are the overall tool call counts (with percentage of total tool calls per model) across all runs.
