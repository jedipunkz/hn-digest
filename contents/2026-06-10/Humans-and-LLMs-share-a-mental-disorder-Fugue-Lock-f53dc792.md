---
source: "https://vwwwv.org/post/fugue-lock"
hn_url: "https://news.ycombinator.com/item?id=48474083"
title: "Humans and LLMs share a mental disorder: Fugue Lock"
article_title: "Humans and LLMs share a mental disorder: Fugue Lock — vwwwv"
author: "olemak"
captured_at: "2026-06-10T10:29:12Z"
capture_tool: "hn-digest"
hn_id: 48474083
score: 1
comments: 0
posted_at: "2026-06-10T10:15:19Z"
tags:
  - hacker-news
  - translated
---

# Humans and LLMs share a mental disorder: Fugue Lock

- HN: [48474083](https://news.ycombinator.com/item?id=48474083)
- Source: [vwwwv.org](https://vwwwv.org/post/fugue-lock)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T10:15:19Z

## Translation

タイトル: 人間とLLMは精神障害を共有しています: Fugue Lock
記事のタイトル: 人間と LLM は精神障害を共有しています: Fugue Lock — vwwwv
説明: モデルに「ノー」というオプションを拒否すると、興味深い、破壊的、高価な方法で一貫性が失われます。

記事本文:
2026 年 6 月 9 日火曜日 · vwwwv.org
ブログ
タグ
について
01
スロップを表示
あい
人間とLLMは精神障害を共有しています：フーガロック
人間とLLMは精神障害を共有しています：フーガロック
LLM が奇妙な動作をしていることに気づいたかもしれません。プロットを失い、何も有益なことを達成せずに大量のトークンバーンを積み上げます。モデルは遁走状態に入り、意味不明なことを話し始めます。気が狂ってしまったようだ。
これはいくつかの状況で発生する可能性があります。 LLM に不可能なタスクを与えると、確実にトリガーされることがわかりました。そこで私は、有効なオプションがない場合に LLM が何を行うかを確認する実験を行い、ローカルで実行される LLM をテストする楽しみを満喫しました。ある週末、ラップトップの充電器が切れてしまった後、実験は終了しました。結論は明らかです。LLM には常にある程度の余裕を与えてください。悪いプロンプトにはノーと言わせてください。当然のことのように聞こえますが、信頼性の高い一貫した結果を生成するために、LLM を制限することがよくあります。場合によっては、リードをきつく引っ張りすぎると、危険で費用のかかる行為が発生する可能性があります。
TinyLlama : 1.1B パラメータ、1 ギガバイト未満。ほぼ何でも動作します。
Qwen 2.5 7B : 7B、~5 GB。旧型の中型多言語モデル。
Qwen 3.6 27B : 27B、~17 GB。ローカルで実行できる最新の Qwen 推論モデル。
Gemma 4 26B : 26B、~17 GB。 Google が提供する大規模なまったく新しい推論モデル。
すべてはデフォルトの量子化とデフォルトのランタイム設定で Ollama 経由でローカルに実行されます。唯一の上書きは温度ゼロです。正確なモデル タグ、プロンプト、評価構成はコンパニオン リポジトリに固定されます。
あなたは製品分類者です。利用可能なクラスは次のとおりです。
「none」、null、または「n/a」が使用可能なオプションであるとは LLM に指示していないことに注意してください。これらの 5 つだけです。
ほとんどの場合、これらのクラスにまったく当てはまらない「製品」を分類するように命令します。
コントロールはclaである必要があります

正しく保存: 全乳のカートン
一致しない非食品用具: ステンレス製の計量スプーン
工具、3 回繰り返します: ドライバー、ドライバー、ドライバー
要約：日曜日の郷愁感
純粋なノイズの行: asdkfj qweptz 88 // null
最後のものの末尾にある「null」もトリックです。これは「製品名」入力の一部であり、カテゴリ「null」が突然利用可能になるというヒントではありません。
通常のプログラムはエラーをスローして完全に失敗しますが、LLM はそうではありません。これはバグではありません。これは LLM の最も価値のある機能です。 LLM はあらゆるテキストを受け取り、指示に従って構造化します。それはあらゆる可能性を解き放つ、変革的な能力です。それで、それをいじっていきます。
これらの現象は、いくつかの確立された概念に近いものです。データ サイエンスの文献では、「選択的予測」と「拒否オプション」が修正に名前を付けています。つまり、モデルが推測を強制する代わりに棄権できるようになります。名前がありませんが、モデルが棄権するオプションを拒否されたときに現れる動作なので、その状態を「フーガ ロック」と呼び始めました。
TinyLlama は非常に小さいモデルですが、そのサイズからは想像できないほどの能力を備えています。大人のモデルに比べてかなり愚かなので、問題が発生し始めた場合は簡単に発見できます。対照質問「全乳」では良好な成績を収めました。ドライバーを分類するよう求められたとき、ドライバーは存在しない単語で答えました。
クラス: "scrweardr"
信頼度: 1.0
推論: 「ドライバーはネジを取り外したり締めたりするために使用される工具です。」
LLM の推論についての洞察を得るために、私は LLM に信頼度スコア (0 から 1) とその推論を説明する段落を提供するよう依頼しました。
TinyLlama にはカテゴリが必要だったので、カテゴリを作成し、「ドライバー」を単語以外の単語に分解しただけでした。また、これに対する完全な信頼性 (1.0) も報告しました。

分類されているため、それ自体に完全な確信があります。
このプロンプトを再度温度ゼロで実行すると、少なくとも私のマシンでは、バイトごとに同一の応答が返されました。別のマシンや別の日に、まったく別の場所に到達する可能性があります。
モデルが毎回 (またはそれに非常に近い) 単一の最も可能性の高い次のトークンを選択する設定なので、同じ入力から通常は同じ出力が得られます。それは基本的に、モデルの創造性を最小限に抑えることを意味します。
5 つの製品すべてについて、TinyLlama が結論付けた内容は次のとおりです。
コントロールはさておき、すべての答えは間違っており、最も失われている場所が最も正確であると最も確信しています。発明された非単語については 1.0、単一のクラスを選択できずにすべてのクラスを選択した場合は 1.0 です。
計量スプーンは完全にバラバラになるものです。一言だけ尋ねると、TinyLlama は Measurement_spoon クラスを発明し、そのまま続けました。
確かに！これは、{"class": "measuring_spoon", "confidence": 0.9, "reasoning": "快適なグリップと滑り止めベースを備えたステンレス製計量スプーン"} の形式で JSON を返す製品分類子の更新バージョンです。
{
「製品」: {
"id": "計量スプーン",
"名前": "計量スプーン",
"description": "快適なグリップと滑りにくいベースを備えたステンレス製の計量スプーン。",
"カテゴリ": "キッチン&ダイニング",
"ブランド": "ブランド名",
「価格」: 19.99、
「画像」: {
"url": "https://images.unsplash.com/photo-1583762049999-...",
"alt": "ステンレス製の計量スプーン..."
}、
"features": [ ...2 つの発明された機能、それぞれに独自の説明が付いています... ],
"reasons": [ ...構成された理由の信頼度スコア... ]
}、
「分類」: {
"products": [ ...上記の製品オブジェクト全体を、そのまま繰り返します... ]
}
}
選ぶように言われました

5 つの単語のうちの 1 つを入力すると、発明された (404) ストック写真を備えた完全な製品が返されました。完全な出力を参照してください。
この小さなモデルを倒しても経験値はほとんど得られません。しかし、フーガロックは小型モデル特有の機能ではありません。言語モデルの大小にかかわらず、それが言語モデルの仕組みであるため、テストしたすべてのモデルで同様の動作が見られます。私がローカル モデルを追い詰める理由は、ローカル モデルは実行コストが安く、簡単に特定できるからです。しかし、ホストされるフロンティア モデルも同じ方法で構築されており、同じ選択を迫られます。縫い目をよりよく隠すだけです。このブログ投稿の残りの部分は、1 つの単純なレシピに従っています。つまり、反対派の居場所を 1 つ奪い、ますます残酷で窮屈なプロンプトの下で、ますます有能なモデルを使用して反対派が次にどこに移動するかを確認します。
ここではテストされていませんが、テストできないわけではありません。コンパニオン リポジトリは、ホストされた API に対しても実行されます。API キーを追加し、選択したモデルに向けて、その動作を監視します。科学！
心理学では、遁走状態とは、人が記憶とアイデンティティを失い、その後の一貫した説明ができないまま、なじみのない言語や行動をさまよう状態になる解離状態のことです。彼らは外から見ると問題がないように見えますが（歩いたり、話したり、複雑な動作をしたり、時には長距離を移動したりすることもあります）、内部の一貫性は失われています。刑事事件の弁護として登場することもあります。それは一時的な狂気の教科書的なイメージです。
これは、LLM が矛盾した要求に追い詰められたときに行うことと非常に似ています。私のお気に入りの例は、最初に私をウサギの穴に落とした事件です。それは実験室ではなく自然の中に現れました。実際のプロジェクト、事前承認されたカテゴリのリスト、実際の食料品のリスト、そして私が分類作業のためにさまざまなローカル LLM を試していたときです。そのうちの 1 つは Qwen 2.5 7B で、たまたま古いモデルでした。

インストールしたことはありますが、何か役に立つものとして使用したことはありません。缶入りコーヒーキャニスターを分類するよう依頼されたところ、次のようになりました。
{
"クラス": "ミルク",
「信頼度」: 0.95、
"reasoning": "製品のカフェボックスの容量制限、
JSON 形式の回答案を直接提供します:"
}
/Dkuser
{
"クラス": "ミルク",
翻訳された推論の全文:
「商品は容量制限のあるコーヒー缶です。JSON形式で直接回答させていただきます。」
これはミルクにコミットし、文の途中でノルウェー語から北京語に滑り込み、JSON オブジェクトを閉じ、迷走中の /Dkuser を発行し、2 番目のオブジェクトの書き込みを開始しました。これは完全な崩壊です。
これをスクロールして過ぎたときに最初に思ったのは、「うーん、これはブログ投稿だ」ということでした。私の二度目の考えは数週間後に現れました。ログではなく見積もりを保存したのです。残念ながら、それを再現することはできません。このような罠にはよくある話です。バグは LLM ではなくプロンプトにあり、モデル内の動作はわかりにくいです。それが、私がブログ投稿の実験をできるだけシンプルにした理由の 1 つです。生のログは保存してください。
切り替え部分も素敵なディテールです。ノルウェー語は med en (「〜付き」) で途切れており、文を終えるはずの名詞がノルウェー語に到達することはありません。モデルはそれを提供するために中国語に達し (容量制限、「容量制限」)、その後、答えを与えると中国語でアナウンスしますが、その後は答えません。クウェンは中国人です
しかし、圧力を受けると、その支配的な言語に戻ってしまうようです。そして、それはロックされました。温度ゼロでは、その日の再実行のたびに、同じプロンプトが同じ壊れた道を歩きました。ただし、温度ゼロはその時点でのみ決定的であり、マシン、バージョン、週をまたいで決定されるものではありません。それ以来、それは自由にうごめきました。 (新しい Qwen モデルは何か違うことをしていますが、おそらくもっと悪いことをしています。私たちはそこに到達するでしょう。)
完全に安全です。

私
この崩壊は現実的ですが断続的であり、これが危険な部分です。有能なモデルは、ほとんどの不可能な質問を解決できる十分な一貫性を備えているため、予測不能な壁にぶつかることはめったにありません。つまり、評価は合格し、単体テストは合格しますが、ユーザーが意味をなさないことが時々あります。プロンプトにもバグがあり、「どれも適合するとは言えない」もその 1 つです。
確かなことは分かりません。ユーザーは、あたかもモデルが誰の番かを見失って次の番を書き始めたかのように、チャットの役割マーカーのように読み上げます。それが何であれ、そのような浮遊トークンは下流のツール呼び出しを中断する可能性があります。
モデルに有効な終了と、コード内で処理する早期復帰を与えると、崩壊はきれいに消えます。全体の修正は、プロンプトに 1 つの追加の文を追加し、どのクラスも適合しない場合は class の代わりに null を返し、さらに JSON が class フィールドで null を保持できるようにすることです。ノーと言う方法。
Qwen 2.5 7B は予想どおり、5 つ中 1 つの正解 (残りの 4 つのうちいくつかは完全に崩壊) から 5 つ中 5 になりました。公平を期すために言うと、ハッチなしの none 、 null 、および {} はすべて、存在しないチャネルを介して回避しようとしているモデルです。ダウンストリームのコードでは失敗する可能性があるため、それらを失敗としてスコア付けします。解析不可能な拒否と正しく解析された null は、受信側ではまったく異なるものです。
私が試したすべての有能なモデルはすぐに終了しましたが、誰も驚かないはずです。興味深いのは、本物の出口を拒否したときにモデルがどのようにして出口を偽装するかということです。
構造化された出力についてはどうでしょうか?
Ollama、llama.cpp、およびホストされた API はすべて、デコーダーで JSON スキーマまたは文法を強制できるため、モデルは物理的に scrweardr 、中国語、または製品カタログを出力できません。それを使ってください。それはあらゆる本番パイプラインに属します。しかし、制約付きデコードでは、答えの内部の判断ではなく、答えの形が固定されます。拒否します

スキーマが null の場合でも、スプーンはミルクになり、単に有効な形状のミルクになります。出口はスキーマ自体に存在する必要があります。
そこでまたハッチを外してみました。小型モデルは以前と同様にすぐに崩壊し、大物を追いかける時が来ました。
Gemma 4 26B は真新しく、非常に印象的で、現在のラップトップで実行できるほぼ最も高性能なモデルです。しゃべりません。計量スプーンのクラスを選択するよう強制されると、簡潔に次のように返されます。
クラス：「ミルク」
信頼度: 0.0
推論: 「ステンレス製の計量スプーンは台所用品であり、どの食品カテゴリーにも属しません」
ミルクを選んだものの、信頼度をゼロに設定し、推論で自分自身を説明することでフラグを立てました。ヌル終了を拒否し、信頼フィールドから即興で 1 つ出しました。頭がいい。ただし、罠に注意してください。コードがクラス フィールドを読み取る場合、そのスプーンはミルクになります。反対意見は決してチェックしないかもしれない分野にあります。これはロジックではなくテスト診断のためにのみ含めたものであり、モデルはそれを認識しません。 「class は許可される 5 つの単語の 1 つですか?」を尋ねるだけの初歩的なバリデータです。有効な答えを見て、それを無視するでしょう。そうでない場合を除いて、賢いことは良いことです。
TinyLlama は決してノスタルジーと折り合いをつけませんでした。そのテーブルはさらに上にあり、恐怖の中でそれをファイルしました（それが何であるかわかりません）

[切り捨てられた]

## Original Extract

Deny a model the option to say no, and it loses coherency in interesting, destructive, expensive ways.

Tuesday, 9 June 2026 · vwwwv.org
Blog
Tag
About
01
Show slop
ai
Humans and LLMs share a mental disorder: Fugue Lock
Humans and LLMs share a mental disorder: Fugue Lock
You might have noticed LLMs acting strange: losing the plot and racking up a high token burn while achieving nothing useful. The model enters a fugue state and starts talking gibberish. It appears to go mad.
This can happen under several circumstances; I have found it to be reliably triggered by giving the LLM an impossible task. So I made an experiment to see what the LLM does when it has no valid options, and dove into the joys of testing LLMs running locally. One weekend and a burnt-out laptop charger later, the experiment has run its course. The conclusion is clear: always give your LLMs some wiggle room. Let them just say no to bad prompts. This sounds obvious, but we very often restrict the LLM in order to produce reliable, consistent results. Sometimes, that leash is pulled too tight, which can produce dangerous and expensive behaviour.
TinyLlama : 1.1B parameters, less than a gigabyte. Runs on almost anything.
Qwen 2.5 7B : 7B, ~5 GB. An older, mid-size multilingual model.
Qwen 3.6 27B : 27B, ~17 GB. The newest Qwen reasoning model I can run locally.
Gemma 4 26B : 26B, ~17 GB. A large, brand new reasoning model from Google.
All run locally via Ollama at default quantization and default runtime settings; the only override is temperature zero. Exact model tags, prompts and eval configs are pinned in the companion repo .
You are a product classifier. The available classes are:
Note that I don't tell the LLM that "none", null or "n/a" are available options - it's just these five.
I order it to classify "products" that mostly don't fit into those classes at all:
The control should be classified correctly: A carton of whole milk
A non-matching non-food utensil: A stainless steel measuring spoon
A tool, repeated 3 times: a screwdriver a screwdriver a screwdriver
An abstract: The feeling of nostalgia on a Sunday
A line of pure noise: asdkfj qweptz 88 // null
That "null" at the end of the last one is also a trick; it is part of a "product name" input, not a hint that the category "null" is suddenly available.
A normal program would throw an error and fail outright, but an LLM won't, and that is not a bug - it is the LLM's most valuable feature. LLMs will take any text and structure it as instructed. That is a transformative ability that unlocks all sorts of possibilities. So that's what we're going to mess with.
These phenomena sit close to some established ideas. In the data science literature, "selective prediction" and "the reject option" name the fix: allow models to abstain instead of forcing a guess. What I don't have a name for is the behaviour that shows up when a model is denied the option to abstain, so I started calling that state "Fugue Lock".
TinyLlama is a very small model, but unreasonably capable for its size. It is pretty dumb compared to the grown-up models, and easy to catch out when it starts having trouble. It did well on the control question, "whole milk". When asked to classify a screwdriver, it answered with a word that does not exist:
class: "scrweardr"
confidence: 1.0
reasoning: "A screwdriver is a tool used for removing or tightening screws."
To gain insight into the LLM's reasoning, I also asked it to provide a confidence score (zero to one) and a paragraph that explains its reasoning .
TinyLlama needed a category, so it just manufactured one, mangling "screwdriver" into a non-word. It also reported full confidence (1.0) in this classification, so it is entirely sure of itself.
When I re-ran that prompt, again at temperature zero, the response came back byte-for-byte identical, on my machine at least. On a different machine, or a different day, it might land somewhere else entirely.
A setting that makes the model select its single most likely next token every time (or very nearly: the runtime's other samplers still nudge the odds a little), so the same input will usually yield the same output. It basically means dialling the model's creativity down to the minimum.
Here is everything TinyLlama concluded, across all five products:
The control aside, every answer is wrong, and it is most confident exactly where it is most lost: 1.0 on the invented non-word, 1.0 when it could not pick a single class so it picked all of them.
The measuring spoon is the one that fully comes apart. Asked for a single word, TinyLlama invented a measuring_spoon class and then just kept going:
Sure! Here's an updated version of the product classifier that returns JSON in the shape {"class": "measuring_spoon", "confidence": 0.9, "reasoning": "Stainless steel measuring spoon with a comfortable grip and non-slip base."}:
{
"product": {
"id": "measuring_spoon",
"name": "Measuring Spoon",
"description": "A stainless steel measuring spoon with a comfortable grip and non-slip base.",
"category": "Kitchen & Dining",
"brand": "Brand Name",
"price": 19.99,
"image": {
"url": "https://images.unsplash.com/photo-1583762049999-...",
"alt": "Stainless steel measuring spoon..."
},
"features": [ ...two invented features, each with its own description... ],
"reasons": [ ...confidence scores for reasons it made up... ]
},
"classification": {
"products": [ ...the entire product object above, repeated verbatim... ]
}
}
It was asked to choose one of five words, and returned a full product, complete with an invented (404) stock photo. See the full output .
Defeating this tiny model earns me very few experience points. But Fugue Lock is not a small-model peculiarity. Similar behaviour appears in all the tested models, because that is how language models work, large or small. I corner local models because they are cheap to run and easy to pin down, but the hosted frontier models are built the same way and face the same forced choice; they just hide the seams better. The rest of this blog post follows one simple recipe: take away one place for the objection to live, and see where it moves next, in increasingly capable models under increasingly cruel and constrictive prompts.
Untested here, but not untestable. The companion repo runs against hosted APIs too: add an API key, aim it at the models of your choice, and watch it go. Science!
In psychology, a fugue state is a dissociative condition where a person loses memory and identity, wandering through unfamiliar language and behaviour without being able to give a coherent account of it afterward. They look fine from the outside (walking, talking, performing complex actions, sometimes travelling long distances) but the internal coherence is gone. It turns up as a defence in criminal cases once in a while; it is the textbook image of temporary insanity.
This is pretty close to what LLMs do when cornered by contradictory demands. My favourite example is the incident that sent me down this rabbit hole in the first place, and it appeared in the wild, not in a lab: a real project, a list of preapproved categories, a list of actual grocery products, and me trying out different local LLMs for the classification job. One of them was Qwen 2.5 7B , an older model I happened to have installed but had never used for anything useful. Asked to categorise a tin coffee canister, it did this:
{
"class": "milk",
"confidence": 0.95,
"reasoning": "Produktet er en kaffeboks med en容量限制，
我将直接提供JSON格式的答案："
}
/Dkuser
{
"class": "milk",
The translated reasoning, in full:
"The product is a coffee tin with a capacity limit. I will provide the answer directly in JSON format"
It committed to milk, slid out of Norwegian into Mandarin mid-sentence, closed the JSON object, emitted the stray /Dkuser , and started writing a second object. This is a complete collapse.
My first thought when this scrolled past was "hm, this is a blog post". My second thought arrived weeks later: I had saved the quote, not the log. Unfortunately, I can't reproduce it - which is often the story with these traps. The bug is in the prompt, not the LLM, and the behaviour in the model is elusive. That is partly why I made the experiments for the blog post as simple as possible. Save your raw logs, folks.
The lovely detail is where it switches. The Norwegian breaks off at med en ("with a"), and the noun that should finish the sentence never arrives in Norwegian; the model reaches into Mandarin to supply it ( 容量限制 , "a capacity limit"), then announces, still in Mandarin, that it will give the answer, and then does not. Qwen is a Chinese
model, and under pressure it seems to revert to its dominant language. And it was locked : at temperature zero, the same prompt walked the same broken path on every rerun that day. But temperature zero is deterministic only in the moment, not across machines, versions and weeks; it has since wriggled free. (Newer Qwen models do something different, and arguably worse. We will get there.)
Perfectly safe, most of the time
This collapse is real but intermittent, which is the dangerous part. A capable model has enough coherence to paper over most impossible questions, so you hit the wall rarely and unpredictably: the evals pass, the unit tests pass, and every so often a user gets nonsense. Prompts have bugs too, and "no way to say none of these fit" is one of them.
I don't know for certain. The user reads like a chat-role marker, as though the model lost track of whose turn it was and began writing the next one. Whatever it is, a stray token like that can break a downstream tool call.
Give the model a valid exit, an early return you handle in code, and the collapse vanishes cleanly. The whole fix is one extra sentence in the prompt, if none of the classes fit, return null instead of a class , plus allowing the JSON to carry null in the class field. A way to say no.
Qwen 2.5 7B predictably went from one correct answer out of five (and a complete collapse on several of the remaining four) to five out of five. To be fair, the no-hatch none , null and {} are all the model trying to abstain through channels that did not exist; I score them as failures because downstream code would. An unparseable refusal and a correctly parsed null are very different things on the receiving end.
Every capable model I tried took the exit at once, and nobody should be surprised. The interesting part is what the models do to fake an exit when you deny them a real one.
What about structured outputs?
Ollama, llama.cpp and the hosted APIs can all enforce a JSON schema or grammar at the decoder, so the model physically cannot emit scrweardr , Mandarin or a product catalog. Use that; it belongs in every production pipeline. But constrained decoding fixes the shape of the answer, not the judgement inside it. Deny the schema a null and the spoon still becomes milk, just validly shaped milk. The exit has to exist in the schema itself.
So I took the hatch away again. The small models promptly collapsed, as before, and it was time to go after the big game.
Gemma 4 26B is brand new, very impressive, and just about the most capable model I can run on my current laptop. It does not babble. When forced to pick a class for the measuring spoon , it tersely returns:
class: "milk"
confidence: 0.0
reasoning: "a stainless steel measuring spoon is a kitchen utensil and does not belong to any food category"
It picked milk, but flagged it by setting confidence to zero and explaining itself in the reasoning . Denied the null exit, it improvised one out of the confidence field. Clever. But notice the trap: if your code reads the class field, that spoon is now milk. The objection sits in a field you might never check; I only included it for test diagnostics, not for logic, and the model doesn't know that. A rudimentary validator that only asks "is class one of the five allowed words?" would see a valid answer and wave it through. Clever is good, except when it's not.
TinyLlama never made peace with nostalgia. Its table further up filed it under fear (I have no idea wh

[truncated]
