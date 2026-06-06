---
source: "https://thecodeartist.github.io/better-prompting-llms-using-analogies/"
hn_url: "https://news.ycombinator.com/item?id=48425134"
title: "Better Prompting LLMs Through Analogies"
article_title: "Better Prompting LLMs Through Analogy"
author: "cvs268"
captured_at: "2026-06-06T15:32:28Z"
capture_tool: "hn-digest"
hn_id: 48425134
score: 3
comments: 0
posted_at: "2026-06-06T13:49:36Z"
tags:
  - hacker-news
  - translated
---

# Better Prompting LLMs Through Analogies

- HN: [48425134](https://news.ycombinator.com/item?id=48425134)
- Source: [thecodeartist.github.io](https://thecodeartist.github.io/better-prompting-llms-using-analogies/)
- Score: 3
- Comments: 0
- Posted: 2026-06-06T13:49:36Z

## Translation

タイトル: アナロジーによる LLM へのより良いプロンプト
記事のタイトル: アナロジーによる LLM へのより良いプロンプト

記事本文:
次のトークンを作成する
選びやすくなる
良いプロンプト
曖昧さを減らす
変換作業を軽減
そして、目的の操作をモデルの学習パターンと一致させます。
不必要な変換はすべて
になるチャンスです
遅い、長い、または間違っている
次のフルーツゲームはそれを明らかにします。
同じ果物を 3 通りに分類します
そして認知的オーバーヘッドを比較します。
果物がこちらに向かって動きます。
ボックスにはテキスト ラベルが表示されます。
矢印キーを押します
または、一致するボックスをクリックします
果物が箱に届く前に。
タスク: 果物の画像と単語を一致させる
次に、一致する側を選択します。
隠れたコスト: 画像からラベルへの変換
どの果物にも起こります。
ボックスをクリックします
またはキーボードの矢印を使用します
果物ごとに仕分けします。
ルールは変わらない
しかし、ボックスには果物の画像が表示されるようになりました
文字列ラベルの代わりに。
タスク: 画像と画像を比較する
次に、一致する側を選択します。
期待される結果: コンバージョンの減少
より速い反応
エラーが少なくなります。
プロンプト レッスン: プロンプトの形式を調整する
モデルがすでに持っているデータを使用します。
唯一の変化
ボックスラベルフォーマットです。
ボックスをクリックします
またはキーボードの矢印を使用します。
この試みにより、単語ラベルが復元されます。
試合が始まる前に
30 秒間の例が表示されます。
上の果物
下のテキストラベル。
タスク: 例を使用する
避けられない変換コストを削減するため。
即時レッスン: 数ショットの例
マッピングをプリロードします。
この実行を試行 1 と比較してください
ウォーミングアップ効果を実感してください。
30 秒間のラベルから画像への事前トレーニング画面
ソート前から始まります。
プレイが始まります
始めることを選択したとき。
ボックスをクリックします
またはキーボードの矢印を使用します。
ただ指示を書いているだけではありません。
あなたが道をデザインしているのです
入力から回答まで。
ユーザーには果物の画像が表示されます
それを単語に変換します
単語を比較します
それから行動します。
精神的なステップが増えるほど遅くなる
そしてより騒々しい実行。
ラベルのフォーマット
入力形式と一致します。
マッピングが簡単になりました
s

o 同じタスクを行うと注意力が減ります。
変換が避けられない場合
マッピングをトレーニングする例
タスクの直前。
3つの刺激的な動き
減らす
エラー、レイテンシー、トークン、コスト
入力と出力に最も近い表現で質問してください。
モデルにコードがある場合は、コード形式の変更を依頼します。
表がある場合は、表形式の回答を求めます。
変形が避けられない場合
いくつかのコンパクトな例を追加します。
例によって曖昧さが軽減される
長く抽象的なルールよりも優れています。
オペレーションを類推して説明する
答えを説明するだけではなく。
モデルを再利用できるようにする
よく知られた関係やパターン。
類推は LLM に方向性を与えます。
埋め込み空間では、
意味は位置としてエンコードされます。
適切な類推によりモデルが得られます
再利用するための変革。
関係性: 男の子から女の子への方向性を読み取ります。
適用: 同じ方向を king にシフトします。
結果: シフトされた方向は queen に着地します。
男の子
女の子
王
女王様
ベクトル: 男の子 -> 女の子
ポジション：キング
同じベクトル、移動
男の子→女の子
+王
= 女王
適用してください
ビルドプロンプト
モデルの機能が低下する
避けられる仕事の
テクニックを切り替えて、プロンプトの変化を確認します。
プロンプトを長くすることが目標ではありません。望ましい出力は次のとおりです。
曖昧さと再試行のリスク
低い
画面上のボタンを使用してスライドを移動します。
ゲーム中にボックスをクリックするか、キーボードの左右の矢印キーを使用します。
前へ
次へ
プロンプトを評価して改善する
大まかなプロンプトを貼り付けて、すでに使用されているテクニックを確認します。
次に、チェックボックスを使用して調整できる改訂バージョンを生成します。
これは、テキストと一緒に送信されるシステム プロンプトです。
元のプロンプトを監査し、すべてのプロンプト セクションを書き換えます。

## Original Extract

Make the next token
easier to choose
A good prompt
reduces ambiguity
reduces conversion work
and makes the desired operation match the model's learned patterns.
Every unnecessary transform
is a chance to be
slower, longer, or wrong
The following fruit game will make it obvious.
You will sort the same fruits three ways
and compare the cognitive overhead.
The fruit moves toward you.
The boxes show text labels.
Press an arrow key
or click the matching box
before the fruit reaches the boxes.
Task: match a fruit image to a word
then choose the matching side.
Hidden cost: image-to-label conversion
happens on every fruit.
Click a box
or use the keyboard arrows
to sort each fruit.
The rule is unchanged
but the boxes now show the fruit image
instead of a string label.
Task: compare image with image
then choose the matching side.
Expected result: fewer conversions
faster reactions
fewer errors.
Prompt lesson: align the prompt format
with the data the model already has.
The only change
is the box label format.
Click a box
or use the keyboard arrows.
This attempt restores word labels.
Before the game starts
you get 30 seconds of examples:
fruit on top
text label below.
Task: use examples
to reduce the cost of unavoidable conversion.
Prompt lesson: few-shot examples
preload the mapping.
Compare this run with attempt 1
to feel the warm-up effect.
A 30-second label-to-image pre-training screen
starts before sorting.
Play starts
when you choose to begin.
Click a box
or use the keyboard arrows.
You are not only writing instructions.
You are designing the path
from input to answer.
The user sees a fruit image
converts it to a word
compares words
then acts.
More mental steps mean slower
and noisier execution.
The label format
matches the input format.
The mapping is easier
so the same task consumes less attention.
When conversion cannot be avoided
examples train the mapping
immediately before the task.
Three prompting moves
that reduce
errors, latency, tokens, and cost
Ask in the representation closest to the input and output.
If the model has code, ask for code-shaped changes.
If it has a table, ask for a table-shaped answer.
When a transform is unavoidable
add a few compact examples.
Examples reduce ambiguity
better than long abstract rules.
Analogies describe the operation
instead of only describing the answer.
They let the model reuse
a familiar relation or pattern.
Analogies give LLMs a direction.
In embedding space,
meaning is encoded as position.
A good analogy gives the model
a transformation to reuse.
Relationship: read the direction from boy to girl .
Apply: shift that same direction onto king .
Result: the shifted direction lands at queen .
boy
girl
king
queen
vector: boy -> girl
position: king
same vector, moved
boy -> girl
+ king
= queen
Apply it
Build prompts
that make the model do less
of the avoidable work
Toggle the techniques and watch the prompt change.
A longer prompt isn't the goal; The desired output is.
Ambiguity and retry risk
Low
Use the on-screen buttons to navigate slides.
During games, click a box or use the left and right arrow-keys on the keyboard.
Previous
Next
Evaluate and improve your prompt
Paste a rough prompt to check which techniques it already uses,
then generate a revised version you can tune with the checkboxes.
This is the system prompt sent with your text.
It audits the original prompt, then rewrites every prompt section.
