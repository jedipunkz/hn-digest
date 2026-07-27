---
source: "https://laurentiugabriel.github.io/token-town/"
hn_url: "https://news.ycombinator.com/item?id=49068477"
title: "TokenTown –> Learn how LLM's work in a SimCity-like world"
article_title: "TokenTown: How an LLM Works, as a City"
author: "laurentiurad"
captured_at: "2026-07-27T12:51:24Z"
capture_tool: "hn-digest"
hn_id: 49068477
score: 9
comments: 3
posted_at: "2026-07-27T12:09:08Z"
tags:
  - hacker-news
  - translated
---

# TokenTown –> Learn how LLM's work in a SimCity-like world

- HN: [49068477](https://news.ycombinator.com/item?id=49068477)
- Source: [laurentiugabriel.github.io](https://laurentiugabriel.github.io/token-town/)
- Score: 9
- Comments: 3
- Posted: 2026-07-27T12:09:08Z

## Translation

タイトル: TokenTown –> シムシティのような世界で LLM がどのように機能するかを学びましょう
記事のタイトル: TokenTown: 都市としての LLM の仕組み
説明: ミニチュア変圧器をリアルタイムで実行するアイソメトリック都市: トークナイザー ドック、アテンション プラザ、KV キャッシュ倉庫、フィードフォワード ミル、サンプラー。

記事本文:
一度に 1 つのトークンを都市としてレイアウトした言語モデル
街は暇だ。以下のプロンプトを入力して、1 つのトークンが往復するのを確認してください。
実際のモデルが持つ 4,096 以上の数字を表す 12 の数字。青はマイナス、暖色はプラスです。
KV キャッシュ上のソフトマックスに注意
ボキャブラリー スタジアムで計算されます。
TokenTown は、すべての地区がトランスフォーマー言語モデルの 1 つの段階であるアイソメトリック都市です。
輸送船団は道路に沿って隠された状態を運びます。それは波止場でトークンに切り分けられ、
鋳造所でベクトルを取得し、その位置をスタンプしてから、レイヤー リングの周りを駆動します (注意、残留、
フィードフォワード、残差) をレイヤーごとに 1 回、スタジアムが確率分布に変換する前に、
サンプラーはトークンを 1 つ選択します。そのトークンがフィードバック ハイウェイを元に戻し、街全体が再び動き始めます。
ブラウザ内で実際に計算され、ライブで実行されます。トークナイザー分割。埋め込みルックアップ。
正弦波位置エンコーディング。レイヤノルム;因果マスキングを使用したマルチヘッドスケーリングドット積注意
実際に増大する KV キャッシュ上。残りは加算されます。 GELU フィードフォワード。および温度/トップポイントサンプリング。
トラック上のバ​​ーは実際のベクトルです。倉庫の上の梁は実際のソフトマックス ウェイトです。
Prefill は実際にすべてのプロンプト トークンを一度に処理しますが、decode は実際に 1 つだけを処理します。
スケールダウン: 数千ではなく 12 次元、数十ではなく 2 つの注目ヘッド、
80 の代わりに 2 ～ 12 のレイヤーがあり、10 万語以上ではなく数百語の語彙が含まれます。
意図的に偽造: 重みはランダムであり、ここでは何もトレーニングされていないため、ランダムな重みが使用されます。
モデルはノイズを発します。出力を読みやすく保つために、最終的なロジットは実際の隠れ状態をブレンドします。
小さな固定コーパスから事前に構築されたバイグラムによる投影。アテンションスコアも鋭くなり、

小さな最初のトークン (「シンク」) と最新性バイアスが与えられるため、マップはトレーニングされたモデルのパターンのように見えます
実際に生産します。この街が書くテキストを風景として扱います。メカニズムを教訓として扱います。
護送船団が初めて地区に到着したとき、それを読み取るのに十分な長さの時間停止します。
地区の説明、量に応じて9秒から26秒の間
言うこと。パネル テキストの下の進行状況バーには、ストップの残り時間が表示されます。
すべての地区の説明が終わると、街は代わりに目に見えるペースで進みます
読みやすいものであり、繰り返されるレイヤーは同じであるため早送りされます。
異なる重みを持つ道路。スペースは任意のストップを無期限に保持します、S
は一度に 1 段階ずつステップを実行し、速度スライダーはすべてをスケールします。
0.4倍から8倍まで読み取りが停止します。リセット (⟲) は、スローツアーを再生します。
始まり。 Run では、すでに読んだ内容が保持されます。
スペース再生/一時停止・S 1 ステージ進む・R ツアーをリセットしてリプレイ・F カメラをフォロー・L ラベル
ドラッグしてパン、スクロールしてズーム、任意の地区をクリックすると説明が表示されます。
景色は車列の近くから始まり、車列に沿って進みます。 ⤢ボタン
左側 (または地図上でダブルクリック) では、都市全体に戻ります。回す
オフ フォローすると、一人でさまようことができます。
PGSimCity (PostgreSQL の都市型モデル) の背後にあるアイデアからインスピレーションを得ています。ここにあるすべてのコード、アート、コピーはオリジナルです。

## Original Extract

An isometric city that runs a miniature transformer in real time: tokenizer docks, attention plaza, KV-cache warehouse, feed-forward mill and the sampler.

A language model laid out as a city, one token at a time
The city is idle. Type a prompt below and watch a single token make the round trip.
12 numbers standing in for the 4,096+ a real model carries. Blue is negative, warm is positive.
Attention softmax over the KV cache
Computed at the Vocabulary Stadium.
TokenTown is an isometric city where every district is one stage of a transformer language model.
A convoy carries a hidden state along the roads: it is cut into tokens at the docks, cast into a
vector at the foundry, stamped with its position, then driven around the layer ring (attention, residual,
feed-forward, residual) once per layer, before the stadium turns it into a probability distribution and
the sampler picks one token. That token drives back up the feedback highway and the whole city runs again.
Genuinely computed, live, in your browser: the tokenizer split; the embedding lookup;
sinusoidal positional encoding; LayerNorm; multi-head scaled dot-product attention with causal masking
over a real growing KV cache; the residual adds; a GELU feed-forward; and temperature / top-p sampling.
The bars on the truck are the actual vector. The beams over the warehouse are the actual softmax weights.
Prefill really does process every prompt token at once while decode really does process only one.
Scaled down: 12 dimensions instead of thousands, 2 attention heads instead of dozens,
2–12 layers instead of 80, and a vocabulary of a few hundred words instead of 100k+.
Deliberately faked: the weights are random, and nothing here was trained, so a random-weight
model would emit noise. To keep the output readable, the final logits blend the real hidden-state
projection with a bigram prior built from a small fixed corpus. The attention scores are also sharpened,
and given a small first-token ("sink") and recency bias, so the map looks like the patterns trained models
actually produce. Treat the text this city writes as scenery; treat the mechanism as the lesson.
The first time the convoy reaches a district it stops long enough to read that
district's explanation, between 9 and 26 seconds depending on how much there is
to say. A progress bar under the panel text shows how long the stop has left.
Once every district has been explained the city runs at a watchable pace instead
of a readable one, and the repeated layers fast-forward because they are the same
road with different weights. Space holds any stop indefinitely, S
steps one stage at a time, and the Speed slider scales everything, including the
reading stops, from 0.4× to 8×. Reset (⟲) replays the slow tour from the
beginning; Run keeps what you have already read.
Space play / pause · S advance one stage · R reset and replay the tour · F follow camera · L labels
Drag to pan, scroll to zoom, click any district for its explanation.
The view starts close on the convoy and rides along with it. The ⤢ button
on the left (or a double-click on the map) pulls back to the whole city; turning
off Follow lets you wander on your own.
Inspired by the idea behind PGSimCity (a city-shaped model of PostgreSQL); all code, art and copy here are original.
