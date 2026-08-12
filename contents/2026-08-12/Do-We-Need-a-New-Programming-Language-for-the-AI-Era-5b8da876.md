---
source: "https://mfallah.com/blog/do-we-need-a-new-programming-language-for-the-ai-era/"
hn_url: "https://news.ycombinator.com/item?id=49269405"
title: "Do We Need a New Programming Language for the AI Era?"
article_title: "Do We Need a New Programming Language for the AI Era?"
author: "dalmif"
captured_at: "2026-08-12T09:01:52Z"
capture_tool: "hn-digest"
hn_id: 49269405
score: 1
comments: 1
posted_at: "2026-08-12T08:37:55Z"
tags:
  - hacker-news
  - translated
---

# Do We Need a New Programming Language for the AI Era?

- HN: [49269405](https://news.ycombinator.com/item?id=49269405)
- Source: [mfallah.com](https://mfallah.com/blog/do-we-need-a-new-programming-language-for-the-ai-era/)
- Score: 1
- Comments: 1
- Posted: 2026-08-12T08:37:55Z

## Translation

タイトル: AI 時代には新しいプログラミング言語が必要ですか?
説明: AI によって生成されたコードと仕様ファイルは、プログラミング インターフェイスとしての人間の言語の限界を明らかにしています。新しい意図言語の提案と、それに対応する Reasoner。

記事本文:
AI時代には新しいプログラミング言語が必要か?モハマド・ファラー
AI時代には新しいプログラミング言語が必要か?
AI によって書かれた PR を受け取ると、特に Android 開発者ではないプロジェクト マネージャーからの場合、その変更によって人間が作成し、思慮深く設計されたコードが台無しになっていると感じると、本当に気分が悪くなります。
私のレビュー コメントはいくつかの小さな問題を修正することはできますが、アーキテクチャ全体を修正することはできません。もっと一般的なコメントを残そうとしましたが、多くの場合、ソリューション全体が書き直されることになり、最初から見直してコードをもう一度理解する必要があります。
AI 生成コードへの依存度が高まるにつれて、コードベースが従来の自動生成コード (Room KSP などのツールで生成される Kotlin や Java コードなど) に似てきているように感じます。ただし、生成されたコードを直接変更したり、レビューしたり、バージョン管理したりすることはありません。
これは、明確に定義された注釈と関数のセットから生成され、出力が完全に決定的であるためです。異なる時間に異なるマシンで生成しても、同じ結果が得られます。
現在の AI コード生成ワークフローには何かが欠けていることがはっきりとわかります。
プロジェクト マネージャーや開発者は、自分の知識と Jira チケットに基づいてプロンプトを作成し、その後、他の開発者が生成されたコードとその変更をレビューするように割り当てられます (場合によっては、元のプロンプトを見ることさえありません)。
このため、最初のプロンプトにもっと時間を費やし、それを Git に保存し、チームと共有し、開発者に意見を求めることを提案する人もいます。
では、このプロンプトは新しい言語になるのでしょうか?
私は、人々が AI を使用して仕様の .md ファイルを書き始めているのを見てきました。より簡単です。すべての文を熟考する必要はありません。

すべてを自分で入力するか、入力します。 Jira チケットを提供するだけで、優れた AI スキルを使用して仕様を生成できます。
しかし、人々は実際にそれらの .md ファイルを読んでいるのでしょうか?
情報が必要な場合でも、AI にファイルを要約してもらったり、その内容に基づいて質問に答えてもらったりします。
AI が単純なタスクの説明を複雑なものに変えるのは簡単です。冗長な言葉を使用したり、同じ要点を言い換えたり、いくつかの異なる方法で繰り返したりします。
これがまさに私たちが今いる状況です。人間が作成したのか、AI によって生成されたのかわからないため、読む価値のあるものは何も感じられず、見る価値のあるものも何も感じられません。作成者が公開する前にわざわざ読んだのか、それとも単純に AI を信頼して残りの人間を放っておいたのかさえわかりません。
長い文章を読むことよりも、作成することがはるかに簡単になりました。
これらの .md ファイルには何が含まれていますか?
これらには通常、コンテキスト、例、参考資料、要件、設計とアーキテクチャの決定が混在したものが含まれます。
これらのファイルを新しい言語として扱う場合、構造やアーキテクチャについてあまり考えずに、何かを書いて動作させることができただけで満足していた 1980 年代に戻ってしまうことになります。
.md ファイル内に記述されているアプリケーションのアーキテクチャについて話しているのではありません。私は .md ファイル自体のアーキテクチャと構造について話しています。
それらには何を含めるべきでしょうか?
相互に参照するにはどうすればよいでしょうか?
私の最初の問題は、それらは人間の言語で書かれており、人間の言語は非常に冗長であるということです。
例を挙げてみましょう。プログラマーとして、どちらを好みますか?
このアカウントの名前はマイクです。
現在アクティブなセッションに関連付けられた ID に基づいて、現在のユーザーの名前を設定できます。

はっきりとマイクであることがわかりました。
"名前" : "マイク"
わかりやすいですね。重要な 2 つの情報 ( name と Mike ) を抽出するためだけに、不要な単語のコレクションを読む必要はありません。
私も .md ファイル全体について同じように感じます。多くの場合、冗長で冗長で、誰も読みたくないような表現がされています。最終的には、大量のテキストから少量の有用なコンテキストのみを抽出することになります。
場合によっては、いくつかの単語や文を強調表示して、後で戻ったときに強調表示された部分だけをざっと見て、頭の中でコンテキストをすぐに復元できるようにする必要もあります。
制約は大規模なシステムの構築に役立ちます
多くのプログラミング言語は、特定のことを簡単にする一方で、他のことを意図的に難しくするように設計されています。それらは制約をもたらします。
BASIC などの言語を使用したプログラミングの初期には、メモリ内のどこにデータを配置し、どこにコードを記述するかを決定できました。私たちはメモリ空間のほぼ全体を制御できました。
それは良いことでしたか、それとも悪いことでしたか?
時間が経つにつれて、私たちはそれらの詳細を抽象化した言語を作成しました。たとえば、C では、関数、構造体、その他の抽象化が導入されました。
システムがさらに複雑になると、人々はオブジェクト指向プログラミングなどのアイデアを広く普及させました。以前よりもさらに多くの制約を追加しました。
次に、関数がどこに属するか、クラスがどのように通信するか、どのデザイン パターンを使用するかを考える必要があります。
これらの制約により、ますます複雑なシステムを設計できるようになりました。
意図を伝えるための言語
今日、私たちは AI によって、初期の「全記憶へのアクセス」の瞬間の別のバージョンを経験していると思います。
そこで私は新しい言語を提案したいと思っています。
ただし、今回は言語が異なります。
私たちは泣き言を言わないだろう

主にマシンや CPU と通信するためのコードです。私たちは他のプログラマーと通信するためのコードを書きました。
それがプログラマーの得意分野です。
関数が適切に実装され、適切な抽象化が使用されていると仮定すると、適切に記述された関数の実装をざっと調べて、その関数が何を行うのかを理解することができます。
AI を使用すると、この種の抽象化を以前よりもはるかに扱えるようになるかもしれません。
すべての実装の詳細を自分で書く必要はないかもしれません。 300行のコードを含む複雑な機能もAIによって実装できる可能性がある。 AI に適切なコンテキストを提供するだけで済みます。
プログラミングは、おそらくアノテーションを使用して、メタプログラミングの形式に移行する可能性があります。
@ context ( "このアプリは、API から天気データを取得してユーザーに表示する天気アプリです" )
クラスアプリ
@ context (「これはアプリケーションのメイン エントリ ポイントです」)
楽しいメイン () {
// $ で始まる関数は実際の関数ではありません。
// これらは、クラスまたはクラスの推移的なセットで置き換えることもできます。
// これらはポインタに似ています。
// アイデアは、関数を実装するのに十分なコンテキストを AI に与えることです。
print ( $get_weather ( "コペンハーゲン" ))
}
// オプション: AI に追加のコンテキストを提供できます。
// ここで、AI に WeatherRepository を生成または使用するように指示します。
// および GetWeatherUseCase。 AI は理解できるほど賢い
// それらが互いにどのように関係しているか。
@ context ($WeatherRepository、$GetWeatherUseCase)
@transitivelyDependsOn ( "retrofit" 、 "okHttp" 、 "kotlinx-serialization" )
@file:UseCase // この関数の実装時にAIが読み込むテキストファイルを指します。
楽しい $get_weather (都市): $Weather
// オプション: リポジトリの実装の詳細を提供することもできます。
楽しい＆WeatherRepository。 get_weather (都市): $Weather {
val 応答 = 天気サービス。ゲットウィ

アテル（都市）
応答を返します。 $toWeather ()
}
どこまで深くするかは開発者次第です。
関数を完全に実装することも、部分的に記述することも、利用可能なコンテキストに基づいて完全に AI に実装させることもできます。
このコードは Reasoner によって処理され、システムの構造 (モジュール、クラス、関数、関係) のテキスト表現が生成されます。関数は名前と責任によって説明されますが、必ずしも実際の実装は含まれません。
推論プロセスは AI モデルによって実行されます。
結果として得られる推論されたファイルは保存され、Git にコミットされます。異なるモデルは異なる推論を生成する可能性があるため、推論された各ファイルには、それを生成したモデルの名前とバージョンがタグ付けされます。
推論されたファイルは、AI を使用して実際のコードを生成するために使用されます。実装モデルは、推論を実行したのと同じモデルである場合もあれば、異なるモデルである場合もあります。
Reasoner の主な目的は、隠された推論プロセスを実装プロセスから分離することです。
モデルがコードを生成するたびにプライベートに推論できるようにする代わりに、その推論の結果を具体的な形式で保存します。
これらのファイルは、元の抽象定義よりもはるかに具体的になります。その結果、異なるモデルが同じ推論されたファイルから同様の実装を生成する可能性があります。
推論されたファイルは直接編集されません。これらを変更するには、開発者はコンテキストまたは抽象コードを更新し、Reasoner を再度実行します。
推論されたファイルは、JSON、TOML、または構造化された人間の言語テキストとして表すことができます。
【推理】
モデル = "サンプルモデル"
バージョン = "1.0"
[コンポーネント。ウェザーリポジトリ ]
責任 = 「フェット

リモート API からの ch 気象データ」
[コンポーネント。 GetWeatherUseCase ]
depend_on = [ "WeatherRepository" ]
戻り値 = "天気"
また、推論段階とコード生成段階の両方で AI モデルに適切なコンテキストを提供するように設計されたツールも存在するでしょう。
この言語の名前はすでに決めています。
私はそれを Ruz と呼んでいます。「ルーズ」と発音します。

## Original Extract

AI-generated code and spec files are exposing the limits of human language as a programming interface. A proposal for a new language of intent — and a Reasoner to go with it.

Do We Need a New Programming Language for the AI Era? Mohammad Fallah
Do We Need a New Programming Language for the AI Era?
When I receive a PR written by AI—especially when it comes from a project manager who is not an Android developer—and I feel that the changes have ruined human-crafted, thoughtfully engineered code, I genuinely feel bad.
My review comments can fix some small issues, but they cannot fix the entire architecture. I have tried leaving more general comments, but that often results in the whole solution being rewritten, which means I need to review it from the beginning and understand the code all over again.
As we rely more on AI-generated code, I feel that our codebases are starting to resemble traditional auto-generated code, such as the Kotlin or Java code produced by tools like Room KSP. However, we never directly modify, review, or even version-control that generated code.
That is because it is generated from a well-defined set of annotations and functions, and the output is fully deterministic. We can generate it on different machines at different times and still get the same result.
We can clearly see that something is missing from our current AI code-generation workflow.
A project manager or even a developer writes a prompt based on their own knowledge and a Jira ticket, and then other developers are assigned to review the generated code and its changes—sometimes without even seeing the original prompt.
Because of this, some people suggest spending more time on the initial prompt, storing it in Git, sharing it with the team, and asking developers for input.
So, is this prompt becoming a new language?
I have seen people start writing specification .md files using AI. It is easier: they do not need to think through every sentence or type everything themselves. They can simply provide a Jira ticket and use a good AI skill to generate the specification.
But are people actually reading those .md files?
Even when they need information from them, they ask AI to summarize the files or answer questions based on their contents.
It is easy for AI to turn a simple task description into something complicated—to use wordy language, rephrase the same point, and repeat it in several different ways.
This is exactly where we are now: nothing feels worth reading, and nothing feels worth watching, because you do not know whether it was created by a person or generated by AI. You do not even know whether the creator bothered to read it before publishing it, or whether they simply trusted the AI and left the rest of us alone with that shit.
Producing long texts has become much easier than reading them.
What do these .md files contain?
They usually include a mixture of context, examples, references, requirements, and design and architecture decisions.
If we are going to treat these files as a new language, then we are back in the 1980s, when we were simply happy that we could write something and make it work, without thinking much about structure or architecture.
I am not talking about the architecture of the application described inside the .md file. I am talking about the architecture and structure of the .md files themselves.
What should we include in them?
How should they reference one another?
My first problem is that they are written in human language, and human language is extremely verbose.
Let me give you an example. As a programmer, which one would you prefer?
The name of this account is Mike.
Based on the identity associated with the currently active session, the name of the current user can be confidently identified as Mike.
"name" : "Mike"
It is easy to understand. I do not need to read a collection of unnecessary words just to extract the only two pieces of information that matter: name and Mike .
I feel the same way about entire .md files. They are often verbose, wordy, and phrased in a way that nobody wants to read. In the end, you extract only a small amount of useful context from a large amount of text.
You might even need to highlight a few words and sentences so that, when you return later, you can scan only the highlighted parts and quickly restore the context in your mind.
Constraints Help Us Build Larger Systems
Many programming languages are designed to make certain things easier while intentionally making other things more difficult. They introduce constraints.
In the early days of programming, with languages such as BASIC, we could decide where to place data in memory and where to write the code. We had control over almost the entire memory space.
Was that a good thing or a bad thing?
Over time, we created languages that abstracted those details away. In C, for example, we introduced functions, structs, and other abstractions.
When systems became even more complicated, people brought ideas such as object-oriented programming into widespread use. We added even more constraints than before.
Now, you need to think about where a function belongs, how classes are supposed to communicate, and which design patterns should be used.
These constraints allowed us to design increasingly complicated systems.
A Language for Communicating Intent
Today, with AI, I believe we are experiencing another version of that early “access to the whole memory” moment.
So I want to propose a new language.
This time, however, the language would be different.
We would not write code primarily to communicate with machines and CPUs. We would write code to communicate with other programmers.
That is what programmers are good at.
We can scan the implementation of a well-written function and understand what it does—assuming that the function is properly implemented and uses the right abstractions.
With AI, we may be able to work with these kinds of abstractions much more than before.
We may not need to write every implementation detail ourselves. Complicated functions containing 300 lines of code could be implemented by AI. We would only need to provide the AI with good context.
Programming could move towards a form of metaprogramming, perhaps using annotations.
@ context ( "This app is a weather app that fetches weather data from an API and displays it to the user" )
class App
@ context (" This is the main entry point of the application ")
fun main () {
// Functions starting with $ are not real functions.
// They can be replaced by a class or even a transitive set of classes.
// They are more like pointers.
// The idea is to give AI enough context to implement the function.
print ( $get_weather ( "copenhagen" ))
}
// OPTIONAL: We can provide additional context to the AI.
// Here, we tell the AI to generate or use WeatherRepository
// and GetWeatherUseCase. The AI is smart enough to understand
// how they are related to each other.
@ context ($WeatherRepository, $GetWeatherUseCase)
@ transitivelyDependsOn ( "retrofit" , "okHttp" , "kotlinx-serialization" )
@file:UseCase // Refers to a text file that the AI reads when implementing this function.
fun $get_weather (city): $Weather
// OPTIONAL: We can also provide implementation details for the repository.
fun & WeatherRepository. get_weather (city): $Weather {
val response = weatherService. getWeather (city)
return response. $toWeather ()
}
It would be up to the developer to decide how deeply they want to go.
They could fully implement a function, partially describe it, or leave it entirely for AI to implement based on the available context.
This code would be processed by a Reasoner , which would produce a textual representation of the system’s structure: its modules, classes, functions, and relationships. The functions would be described by name and responsibility, without necessarily including their actual implementations.
The reasoning process would be performed by an AI model.
The resulting reasoned files would be preserved and committed to Git. Different models might produce different reasoning, so each reasoned file would be tagged with the name and version of the model that produced it.
The reasoned file would then be used to generate the actual code using AI. The implementation model could be the same model that performed the reasoning, or it could be a different one.
The main purpose of the Reasoner would be to separate the hidden reasoning process from the implementation process.
Instead of allowing the model to reason privately every time it generates code, we would preserve the result of that reasoning in a concrete form.
These files would be much more specific than our original abstract definitions. As a result, different models could potentially produce similar implementations from the same reasoned file.
The reasoned files would not be edited directly. To change them, developers would update the context or abstract code and run the Reasoner again.
The reasoned files could be represented as JSON, TOML, or structured human-language text.
[ reasoning ]
model = "example-model"
version = "1.0"
[ component . WeatherRepository ]
responsibility = "Fetch weather data from the remote API"
[ component . GetWeatherUseCase ]
depends_on = [ "WeatherRepository" ]
returns = "Weather"
There would also be tools designed to provide AI models with the right context during both the reasoning phase and the code-generation phase.
I have already chosen a name for this language.
I call it Ruz , pronounced “rooz.”
