---
source: "https://www.danielvaughn.dev/posts/huzzah/"
hn_url: "https://news.ycombinator.com/item?id=49378768"
title: "Show HN: Huzzah – a novel approach to coding with AI"
article_title: "Huzzah"
image: "https://www.danielvaughn.dev/opengraph-image.jpg"
author: "danielvaughn"
captured_at: "2026-08-20T19:24:14Z"
capture_tool: "hn-digest"
hn_id: 49378768
score: 3
comments: 0
posted_at: "2026-08-20T19:05:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Huzzah – a novel approach to coding with AI

- HN: [49378768](https://news.ycombinator.com/item?id=49378768)
- Source: [www.danielvaughn.dev](https://www.danielvaughn.dev/posts/huzzah/)
- Score: 3
- Comments: 0
- Posted: 2026-08-20T19:05:36Z

## Translation

タイトル: Show HN: Huzzah – AI を使用したコーディングへの新しいアプローチ
記事タイトル：ハッザ
説明: AI を使用してコーディングする新しい実験的な方法
HN本文：皆さんこんにちは。私は Huzzah という実験的なエディターに取り組んでいます。今年の 1 月以来、私はコーディング エージェントとほぼ独占的に仕事をしてきましたが、ここ数か月間、コーディング エージェントに完全に疲れ果てていると感じ始めました。それらは素晴らしいものですが、必要な変更ごとに完全な文章を書くのがますます面倒だと感じています。それだけでなく、コードベースには複雑さの制限があるようです。ある点を超えると、エージェント自体が混乱し始めます。コードの作成に戻りたいと思っていますが、完全な手動コーディングに戻りたくないのです。そこで私は、次のような対話パラダイムを考え出しました。 1. 自分にとって最も意味のある方法で疑似コードを作成します。
2. 保存時に、エディターは作業内容を実際のソース コードと同期します。
3. 疑似コードは生成されたコードとともに永続化されるため、プロンプトは事実上、意図の保存された記録となります。
すべてのユースケースで機能するとは限りませんが、最初にプレイした限りでは、非常に楽しいことがわかりました。現時点では、これは単なる概念実証です。インストール手順は Readme にあります: https://github.com/danielvaughn/hz また、実際の動作のビデオもここで見ることができます: https://x.com/danielvaughn/status/2090456808431165715 乾杯!

記事本文:
AI を使用してコーディングする新しい実験的な方法
あなたが私のようなソフトウェア エンジニアであれば、2026 年の最初の数か月は信じられないほど素晴らしいものでした。
コーディング エージェントが突然十分な性能を発揮するようになり、手動でコードを記述する必要がなくなりました。
しかし、あなたが私と同じなら、いつか壁にぶつかることになるでしょう。
ハネムーン期間は終わり、目新しさは消え去った。
ドーパミンの発作はもうありません。
8月だというのに、すっかり疲れてしまった。
正直に言うと、コードベースに加えたいすべての変更を説明するために長文の英語を書くのは死ぬほどうんざりしています。
ただし、すべてのコードを手動で記述する作業には戻りたくありません。
その練習には本当に退屈なことがありましたが、それは私にとって…いや、残りの人生で避けたいものです。
それでも、自分のコードが何をしているのかをよりよく理解し、制御する必要があると感じています。
自分の出力が高品質で信頼性の高いソフトウェアであることを知りたいです。
プロとして自分に自信を持ちたい。
それで私もケーキを食べて食べる方法を見つけようとしています。
コーディングエージェントに関する私の問題は、
人間の意図についての信頼できる記録はありません。プロンプトは破棄され、コードは AI によって生成された場合とそうでない場合があります。人間が機械に求めるものを表現する中心的な権威を私たちは失ってしまったので、その事実と戦うことが重要だと思います。
AI チャットは、アプリケーション自体ではなく、アプリケーションへの変更を説明する必須の段階的な指示です。これは、開発過程で命令が頻繁に繰り返されるため、トークンを消費することを意味します。これは非効率的です。
自然言語の多くは、情報提供のためではなく、社会的な理由で存在します。平均的な文章には本当の情報がほとんどありません。この方法でマシンに書き込むのは面倒です。
これらの問題に対処するために、私は実験的なエディターを構築しています。
私はそれを「ハッザ」と呼んでいます。

d これは、LLM を操作するための代替パラダイムを提案します。
コーディング エージェントの場合、プロンプトは (a) 長い形式、(b) 命令型、および (c) 一時的なものになります。
Huzzah では、プロンプトは (a) 疑似コード、(b) 宣言的、および (c) 永続的です。
見せていただくだけなら簡単ですよ。
お使いのブラウザは埋め込みビデオをサポートしていません。できます
Huzzah のデモンストレーションを直接見る
非常に単純な例を考えてみましょう。AI を使用してシュワシュワとした音を作りたいとします。これを 2 回実行します。1 回目はコーディング エージェントを使用し、もう 1 回目は Huzzah を使用します。
選択したツールでチャットを開始し、次のような内容を入力します。
100 回ループする関数を作成します。数値が 3 で割り切れる場合は、「fizz」と出力します。数値が 5 で割り切れる場合は、「buzz」と出力します。数値が両方で割り切れる場合 (たとえば 15 など)、「fizzuzz」と出力します。
編集が必要な場合は、チャットにフォローアップ メッセージを送信します。
100 回ループする代わりに、関数は数値入力を受け取り、その回数だけループする必要があります。
満足するまでこのプロセスを繰り返します。
fizz_buzz.hz という新しいファイルを作成します。
その中で、好きなように疑似コード表現を記述します。
個人的にはこうします:
fizz_buzz()
ループ100
モジュロ 3 ? 「フィズ」
5? 「バズ」
両方？ 「フィズバズ」
ファイルを保存すると、Huzzah はそのファイルから実際のコードを自動的に生成します。
編集が必要な場合は、ファイルを更新するだけです。
fizz_buzz(n)
ループn
モジュロ 3 ? 「フィズ」
5? 「バズ」
両方？ 「フィズバズ」
ファイルを保存すると、Huzzah は差分をキャプチャし、LLM へのプロンプトとして使用します。これにより、影響を受けたソース コードが再生成されます。
他のシナリオでこれがどのようになるかをよりよく理解できるように、いくつかの代替例を次に示します。
リストカート
在庫をリストアップする
mock_data = // モックデータを含めます
初期化()
在庫.fill(モックデータ)

追加アイテム(id)
cart.add(ID によるアイテム)
削除アイテム(id)
cart.filter(ID によるアイテム)
チェックアウト()
cart.sum (価格による商品) を返し、価格としてフォーマットします
2.Todoリスト
藤堂{
ID: 整数
テキスト: str
完了: ブール値
}
add_todo(テキスト)
todos.add(テキスト、完了 = false)
トグル_todo(id)
todo = todos.id で取得
todo.completed = NOT .completed
削除_todo(id)
todos.id によるフィルター
利点
すでにいくつかのメリットが感じられるはずです。長い形式のプロンプトよりも擬似コードの方がはるかに簡潔で読みやすいことに注目してください。さらにいくつかあります:
この方法でプロンプトを作成すると、コードの形状を設計しているように感じるため、集中力が高まります。
好きなだけ簡潔にも冗長にもできます。
擬似コードは、人間が意図を表現するために作成したものであるため、開発者ドキュメントとして機能します。
言語に依存しない疑似コードを作成し、それを複数の言語または環境ターゲットの基礎として使用できます。 CRDT のような複雑なアルゴリズムを考えてみましょう。
もちろん、特効薬はありません。
一部の例外:
このアプローチを大規模に行うと問題が発生する可能性は十分にあります。
これは明らかに、既存のコードベースよりも新しいコードベースの方が理想的です。
専門知識が不足している場合は、おそらく自然言語の方が簡単な対話方法です。
ファイル間の依存関係など、確実に表現するのが難しいものもあります。
LSP タイプの機能は利用できません (ただし、これは生成される可能性があります)。
Huzzah は積極的に開発が進められており、現時点では実験段階にのみ存在します。
ソース コードとセットアップ手順はここで見つけることができます。
ぜひ試してみて、ご意見をお聞かせください。

## Original Extract

A new experimental way to code with AI

Hello everyone. I've been working on this experimental editor called Huzzah. I've been working almost exclusively with coding agents since January of this year, and over the past few months I began to feel utterly exhausted by them. They're great, but I'm finding it more and more tedious to write full sentences for every change I want. Not only that, but it seems there's a complexity limit for codebases - beyond a certain point the agent begins confusing itself. I'd like to go back to writing code, but I don't want to go all the way back to fully manual coding. So I've come up with this interaction paradigm where you: 1. write pseudocode in whatever way makes the most sense to you
2. on save, the editor synchronizes your work to real source code
3. the pseudocode is persisted alongside the generated code, making your prompt effectively a stored record of intent.
It may not work for every use case, but in my initial playthroughs I've found it very enjoyable. Right now it's just a proof of concept - installation instructions are here in the readme: https://github.com/danielvaughn/hz You can also watch a video of it in action here: https://x.com/danielvaughn/status/2090456808431165715 Cheers!

A new experimental way to code with AI
If you’re a software engineer like me, the first few months of 2026 were incredible.
Coding agents suddenly became good enough that we no longer needed to manually write code.
But if you’re like me, then sometime later you hit a wall.
The honeymoon period ended, and the novelty wore off.
No more dopamine hits.
It’s August, and I feel utterly fatigued .
To be honest, I’m sick to death of writing longform English to describe every change I want to my codebase.
However, I also don’t want to go back to writing all my code manually.
There was real tedium in that practice that I’d prefer to avoid for…well, the rest of my life.
And yet, I sense that I need to have better insight and control over what my code is doing.
I want to know that my output is high quality, reliable software.
I want to feel good about myself as a professional.
So I’m trying to find a way to have my cake and eat it too.
My problem with coding agents is that
There’s no reliable record of human intent. Prompts are discarded, and the code may or may not have been generated by AI. We’ve lost the central authority that expresses what the human wants out of the machine, and I think it’s important to contend with that fact.
AI chats are imperative, step-by-step instructions that describe changes to the application, not the application itself. This means instructions are often repeated, and thus consume tokens, many times over the course of development. This is inefficient.
Much of natural language exists for social reasons, not informational. The average sentence is scarce in real information. Writing in this manner, to a machine, is cumbersome.
To address these problems, I’m building an experimental editor.
I’m calling it Huzzah, and it poses an alternative paradigm for working with LLMs.
With coding agents, prompts are (a) longform, (b) imperative, and (c) transient.
With Huzzah, prompts are (a) pseudocode, (b) declarative, and (c) persistent.
It’s easier if I just show you.
Your browser does not support embedded video. You can
watch the Huzzah demonstration directly
Let’s take a very simple example - say you want to use AI to create fizz buzz . We’ll do this twice - once with coding agents and another with Huzzah.
You start a chat in your tool of choice, and type something like the following:
Create a function that loops 100 times. If the number is divisible by 3, print “fizz”. If the number is divisible by 5, print “buzz”. If the number is divisible by both (like 15 for example), print “fizz buzz”.
If you need to make an edit, you’d send a follow up message to the chat:
Instead of looping 100 times, the function should take a number input and the function should loop that amount of times.
You repeat this process until you’re satisfied.
You create a new file called fizz_buzz.hz .
In it, you write a pseudocode representation, however you like.
This is how I’d do it, personally:
fizz_buzz()
loop 100
modulo 3 ? "fizz"
5 ? "buzz"
both ? "fizz buzz"
You save the file, and Huzzah automatically generates real code from it.
If you need to make an edit, simply update your file:
fizz_buzz(n)
loop n
modulo 3 ? "fizz"
5 ? "buzz"
both ? "fizz buzz"
When you save the file, Huzzah captures the diff and uses it as the prompt to the LLM. The affected source code is thus regenerated.
To give you a better sense for what this could look like in other scenarios, here are some alternative examples.
list cart
list inventory
mock_data = // include some mock data
init()
inventory.fill(mock_data)
add_item(id)
cart.add(item by id)
remove_item(id)
cart.filter(item by id)
checkout()
return cart.sum(item by price) and format as price
2. Todo List
Todo {
id: int
text: str
completed: bool
}
add_todo(text)
todos.add(text, completed = false)
toggle_todo(id)
todo = todos.get by id
todo.completed = NOT .completed
remove_todo(id)
todos.filter by id
Benefits
You should be able to see some benefits already. Notice how much more terse and readable the pseudocode is than the longform prompts? Here are some more:
Writing prompts this way engages your mind, because it feels much more like you’re designing the shape of the code.
You can be as terse or as verbose as you like.
The pseudocode acts as developer documentation because a human wrote it to express their intent.
You could write a language agnostic pseudocode and use it as the basis for multiple language or environmental targets. Think complex algorithms, like a CRDT .
There are no silver bullets, of course.
Some exceptions:
It’s entirely possible that there are issues with this approach at scale.
This is obviously more ideal for new codebases than existing ones.
If you lack domain expertise, natural language is probably the easier interaction method.
Some things may be more difficult to reliably express, like cross-file dependencies.
LSP-type features would not be available (though this could plausibly be generated).
Huzzah is actively being developed, and exists only in an experimental state for now.
You can find the source code and setup instructions here .
Please give it a spin and let me know what you think!
