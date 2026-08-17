---
source: "https://juleskreuer.eu/research/tabbench-llm"
hn_url: "https://news.ycombinator.com/item?id=49327601"
title: "Are LLMs good in-context tabular classifiers?"
article_title: "TabBench-LLM"
author: "not_a_feature"
captured_at: "2026-08-17T08:31:59Z"
capture_tool: "hn-digest"
hn_id: 49327601
score: 1
comments: 0
posted_at: "2026-08-17T07:45:17Z"
tags:
  - hacker-news
  - translated
---

# Are LLMs good in-context tabular classifiers?

- HN: [49327601](https://news.ycombinator.com/item?id=49327601)
- Source: [juleskreuer.eu](https://juleskreuer.eu/research/tabbench-llm)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T07:45:17Z

## Translation

タイトル: LLM はコンテキスト内の表形式の分類子として優れていますか?
記事タイトル: TabBench-LLM
説明: TabBench-LLM は、大規模な言語モデルを、2 つのベースラインを使用して、少数ショットのコンテキスト内の表形式の分類子として直接評価します。

記事本文:
ライト モード !a.isFolder&&!b.isFolder||a.isFolder&&b.isFolder?a.displayName.localeCompare(b.displayName,void 0,{numeric:!0,sensitivity:\"base\"}):!a.isFolder&&b.isFolder?1:-1","filterFn":"node=>node.slugSegment!==\"tag s\"","mapFn":"node=>{node.displayName=node.displayName.charAt(0).toUpperCase()+node.displayName.slice(1)}"}">内容
TabBench-LLM は、大規模な言語モデルを、2 つのベースラインを使用して、少数ショットのコンテキスト内表形式分類子として直接評価します。これは、表形式 ML の一般的なベンチマークである TabArena を LLM 設定に適合させます。
プライマリ スイートは 19 個の決定論的合成タスクであり、その機能名とクラス名には現実世界の意味はありません。そのため、決定ルールはプロンプト内の行から推測することしかできず、決して呼び出すことはできません。 38 個のパブリック TabArena-v0.1 データセットのセカンダリ スイートは、それが実際のテーブルに引き継がれるかどうかをチェックします。未調整のランダム フォレストと TabPFN v2 は同一の分割に適合します。
Elo は共有データセットに対するペアごとの評価であり、ランダム フォレスト = 1000 になるように調整されています。
両方が試行したターゲットのみでモデルを比較しますが、次の場合に選択バイアスを修正できません。
ターゲットが欠けていると体系的に難しくなります。そのため、ヘッドラインは 1 年に 1 回だけ発表されます。
モデルは同じスケジュールされた合成ブロックを完了しました。
データセットごとに、10、20、50、または 100 行の層別トレーニング テーブルと、すべてのモデルで共有される固定のホールドアウト テスト セットを描画します。各 LLM リクエストはラベル付きテーブルと 1 つのラベルなしテスト行を運び、単一のクラス トークンを要求します。実際のデータセットはさらに、可視のクラス名と不透明なトークンを比較します。合成ターゲットはすでに無意味であるため、不透明なアームのみが実行されます。
8 列にわたる超平面
クソ
4 つの列のうち 2 つの列の XOR
指輪
6 列中 2 列の放射状定規
クラスA
クラスB
合成スイートの概念的なスケッチ。そのうちの3つ

19のレシピ、
有益な 2 つの次元で描画されます。実際のタスクは 6 ～ 64 列にあります
不透明なクラストークンを持つ x1..xd という名前。破線は真の決定を示します
ルール;間違った側の点はレシピのラベルノイズです。
ディスカッション
LLM が表形式の予測タスクに使用されることはほとんどなく、表形式の予測タスクで評価されることもあまりありません。この 2 つは互いに補強し合っていますが、どちらも多くの公開された証拠に基づいていません。
いくつかの障害が立ちはだかっています。
まず、発信者が実際に何を返すかを考えてみましょう。出力は確率ではなくトークンです。
しきい値の調整、ランキングメトリクス、およびコスト重視の決定には、調整されたスコアが必要です。言語化された信頼度は十分に調整されておらず、logprob は常に公開されるわけではなく、回答は採点前に解析する必要があります。予測も不安定です。デモンストレーションのシャッフル、クラス名の変更、またはリサンプリングによって予測が変わる可能性があります。機能の重要性や再現可能なアーティファクトはなく、デプロイメントでは行がサードパーティ API に送信され、ホストされたモデルは変更または消滅する可能性があります。
続いて入力側。数値はテキストとして到着します。トークナイザーは 0.4931 をフラグメントに分割し、スケール、順序、距離を再構築する必要がありますが、ツリーは値を直接使用します。 CSV、JSON、マークダウン、機能の順序と精度はすべてスコアに影響を与える可能性があります。ベースラインはマトリックスを参照します。 LLM は 1 つの任意のレンダリングを認識しますが、これはモデル全体で修正されていますが、文献全体では修正されていません。
また、この体制はスケールしません。コンテキスト内学習はプロンプトの長さによって制限され、推論にはテスト行ごとに 1 つのリクエストが必要で、1 つのヘッドライン モデルには最大 9,500 件のリクエストが必要ですが、ランダム フォレストと TabPFN v2 はラップトップ上で同じ分割を数秒で実行できます。したがって、100 例の設定が実際の比較となります。このコストを考えると、パイプライン内での地位を獲得するには、LLM が単に同等であるだけでなく、明らかに優れている必要があります。
アンダーン

すべては測定の問題にあります。標準的な表形式のデータセットはターゲットも含めて何年もオンラインになっているため、再現と推論の区別がつきません。したがって、TabArena データセットのより強力な結果は、暗記によって助けられる可能性があります。合成スイートと不透明なクラス トークンはそのリスクを軽減しますが、TabPFN には少し不公平な利点があります。データはトレーニングで使用されたものと同様の事前分布から生成されます。
このいずれも、誰かが LLM を除外すべきであると決定する必要はありませんでした。期待される結果が弱く、コストが高く、プロトコルがあいまいであれば、それだけで十分でした。ベンチマークを実行すると、その推論が数値に置き換えられ、それらがどの程度離れているか、どのような種類の構造に基づいているかがわかります。
Elo - ペアごとの評価。ランダム フォレスト = 1000 になるように調整されています。
スコア - データセットごとの最小-最大正規化マクロ F1 (最良のモデル = 1、最悪 = 0)。
改善可能性 % - データセットごとの最良のモデルとの平均相対ギャップ。低いほど良いです。
ランキングでは、セレクターを介して記録された任意のメトリックを使用できます。
コードと実験計画の準備ができました。ホスト型フロンティア モデルには依然として API が必要
ノルマ。クレジット、より高いレート制限、または GPU 時間を使用すると、縮小せずにモデルを追加できます
共通評価ブロック。
19 データセット × 5 フォールド × 100 保持行、不透明なラベル、推論は無効です。
合成スイート上の 4 つの少数ショット サイズ (10 / 20 / 50 / 100) すべて。
あらゆるサンプル サイズの合成スイートと実際のスイート、および機能上限の実行。
プロバイダー アクセスにより、フロンティア システムをさらに追加できます。評価対象のみ。いいえ
暗黙の支持または提携。
予備のクレジット、レート制限の引き上げ、未使用の許可、または GPU 時間はすべて役に立ちます。開く
開始する問題。 API キーを public に貼り付けないでください。
番号は tabbench-llm サイトによって再生成され、 files/tabbench-llm/leaderboard.json としてアップロードされます。グラフはブラウザ上で Plotly を使用して描画されます。

、 jsdelivr からロードされます。 Figure のエクスポートでは、同じ CDN からの JSZip を使用します。
Quartz v4.5.2 で作成 © 2026

## Original Extract

TabBench-LLM evaluates large language models as few-shot, in-context tabular classifiers, head-to-head with two baselines.

Light mode !a.isFolder&&!b.isFolder||a.isFolder&&b.isFolder?a.displayName.localeCompare(b.displayName,void 0,{numeric:!0,sensitivity:\"base\"}):!a.isFolder&&b.isFolder?1:-1","filterFn":"node=>node.slugSegment!==\"tags\"","mapFn":"node=>{node.displayName=node.displayName.charAt(0).toUpperCase()+node.displayName.slice(1)}"}"> Content
TabBench-LLM evaluates large language models as few-shot, in-context tabular classifiers , head-to-head with two baselines. It adapts TabArena , a general benchmark for tabular ML, to the LLM setting.
The primary suite is 19 deterministic synthetic tasks whose feature and class names carry no real-world meaning, so the decision rule can only be inferred from the rows in the prompt, never recalled. A secondary suite of 38 public TabArena-v0.1 datasets checks whether that carries over to real tables. Untuned Random Forest and TabPFN v2 fit the identical splits.
Elo is a pairwise rating over shared datasets, calibrated so Random Forest = 1000. It
compares models only on targets they both attempted, but cannot correct selection bias when
the missing targets are systematically harder — so the headline is released only once every
model has completed the same scheduled synthetic block.
Per dataset we draw a stratified training table of 10, 20, 50 or 100 rows and a fixed held-out test set, shared by every model. Each LLM request carries the labelled table and exactly one unlabelled test row, and asks for a single class token. Real datasets additionally compare visible class names with opaque tokens; synthetic targets are meaningless already, so only their opaque arm is run.
a hyperplane over 8 columns
xor
XOR of two of 4 columns
rings
radial rule in 2 of 6 columns
class A
class B
Conceptual sketch of the synthetic suite. Three of the 19 recipes,
drawn in their two informative dimensions. The real tasks live in 6–64 columns
named x1..xd with opaque class tokens. Dashed lines mark the true decision
rule; points on the wrong side are the recipe's label noise.
Discussion
LLMs are rarely used for tabular prediction tasks, nor are they often evaluated on them. The two reinforce each other and neither is based on much published evidence.
There are several things standing in the way.
Firstly, consider what a caller actually gets back. The output is a token, not a probability .
Threshold tuning, ranking metrics and cost-sensitive decisions require a calibrated score. Verbalised confidences are poorly calibrated, logprobs are not always exposed and answers must be parsed before scoring. Predictions are also unstable: shuffling demonstrations, renaming classes or resampling can change them. There are no feature importances or reproducible artefacts, deployment sends rows to a third-party API, and hosted models can change or disappear.
Then the input side. Numbers arrive as text : a tokeniser splits 0.4931 into fragments and must reconstruct scale, ordering and distance, while a tree uses the value directly. CSV, JSON, Markdown, feature order and precision can all affect the score. A baseline sees the matrix; the LLM sees one arbitrary rendering, fixed here across models but not across the literature.
The regime also does not scale : in-context learning is limited by prompt length, inference needs one request per test row and one headline model takes ~9,500 requests, while Random Forest and TabPFN v2 fit the same splits in seconds on a laptop. The 100-example setting is therefore the practical comparison. At this cost, an LLM must be clearly better, not merely comparable, to earn a place in a pipeline.
Underneath it all lies a measurement problem. Standard tabular datasets have been online for years , including targets, so recall is indistinguishable from inference. The stronger results on TabArena datasets are therefore likely helped by memorisation. The synthetic suite and opaque class tokens reduce that risk, although TabPFN has a slightly unfair advantage there: the data is generated from priors similar to those used in its training.
None of this required anyone to decide that LLMs should be excluded. Weak expected results, high cost and an ambiguous protocol were enough on their own. Running the benchmark replaces that inference with a number: how far off they are and on which kinds of structure.
Elo - pairwise rating, calibrated so Random Forest = 1000.
Score - min–max normalized macro-F1 per dataset (best model = 1, worst = 0).
Improvability % - mean relative gap to the best model per dataset; lower is better.
Rankings can use any recorded metric via the selector.
The code and the experiment design are ready; hosted frontier models still need API
quota. Credits, higher rate limits, or GPU hours let us add models without shrinking
the common evaluation block.
19 datasets × 5 folds × 100 held-out rows, opaque labels, reasoning off.
All four few-shot sizes (10 / 20 / 50 / 100) on the synthetic suite.
Synthetic and real suites at every sample size, plus the feature-cap runs.
Provider access can add further frontier systems. Evaluation targets only; no
endorsement or affiliation implied.
Spare credits, a raised rate limit, an unused grant, or GPU hours all help. Open an
issue to start; never paste an API key in public .
The numbers are regenerated by tabbench-llm site and uploaded as files/tabbench-llm/leaderboard.json . The charts are drawn in the browser with Plotly , loaded from jsdelivr ; the figure export uses JSZip from the same CDN.
Created with Quartz v4.5.2 © 2026
