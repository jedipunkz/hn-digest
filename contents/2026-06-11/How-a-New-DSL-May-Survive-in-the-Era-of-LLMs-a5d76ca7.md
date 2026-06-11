---
source: "https://www.williamcotton.com/articles/how-a-new-dsl-survives-in-the-era-of-llms"
hn_url: "https://news.ycombinator.com/item?id=48490939"
title: "How a New DSL May Survive in the Era of LLMs"
article_title: "How a New DSL May Survive in the Era of LLMs - William Cotton"
author: "williamcotton"
captured_at: "2026-06-11T16:28:08Z"
capture_tool: "hn-digest"
hn_id: 48490939
score: 1
comments: 0
posted_at: "2026-06-11T14:35:41Z"
tags:
  - hacker-news
  - translated
---

# How a New DSL May Survive in the Era of LLMs

- HN: [48490939](https://news.ycombinator.com/item?id=48490939)
- Source: [www.williamcotton.com](https://www.williamcotton.com/articles/how-a-new-dsl-survives-in-the-era-of-llms)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T14:35:41Z

## Translation

タイトル: LLM の時代に新しい DSL はどのように生き残るか
記事のタイトル: LLM の時代に新しい DSL はどのように生き残るか - William Cotton
説明: LLM の時代に新しい DSL がどのように生き残るか

記事本文:
LLM の時代に新しい DSL はどのように生き残れるのか - William Cotton
ウィリアムコットン.com
LLM の時代に新しい DSL が生き残るには
過去数十年にわたり、Python、Rust、Ruby、その他の「レガシー」ソフトウェア言語で書かれたプロジェクトが数え切れないほどありました。このコードはすべて、LLM にとって素晴らしい材料となっています。これらのモデルが進歩するにつれて、幻覚の事例は劇的に減少しました。
しかし、トレーニングの対象となるのは、単なるソース コードの量だけではありません。これは、これらの言語に関する高度なツールです。型チェッカー、リンター、言語サーバー、コンパイラー、インタープリター、テスト ハーネスなど、何でもあります。これらのツールはすべて、ソフトウェアを実際に接地し、LLM エージェントに即時にフィードバックを提供します。たとえば、幻覚は、コードが実行される前に型チェッカーによって検出されます。
このため、将来のモデルをトレーニングするためのさらに多くのコンテンツを作成するために、これらのレガシー ソフトウェア言語がさらに多く使用されています。
私たちはある種のフィードバックループに陥っているようです。
では、LLM の時代に生き残るために新しい言語は何をすべきでしょうか?
その答えは、優れたドキュメント、優れたマーケティング、優れたツールなど、かつての物事の仕組みとそれほど変わりません。将来の言語ユーザーはどのようにして新しい言語について知るのでしょうか?彼らはどうやって乗り込むのでしょうか？この言語は既存のツールやワークフローとどのように統合されますか?最新の言語には、堅牢な言語サーバーが必要になります。適切なオンボーディング フローを備えた広範なドキュメント セットが必要になります。
新しい言語も、LLM エージェントとうまく連携するにはもう少し必要になります。
1 つのアプローチは、新しい言語でバイナリ自体から AGENTS.md タイプのファイルを作成することです。次のようなものです。
これは、私が取り組んできた実験的な Web アプリケーション DSL である Web Pipe で使用される LLM テンプレートです。

最近続いています。
Web パイプには他の DSL にはない多くの利点があるため、これは少し不公平です。 jq、Lua、JavaScript、SQL、その他いくつかの言語が埋め込まれています。 LLM はこれらの言語にすでに慣れているため、パイプライン指向のアプローチの構文とセマンティクスについては、それほど学ぶ必要はありません。
私はコーデックスでのワンショット プロンプトで、この 1 つの AGENTS.md テンプレート ファイルだけをガイダンスとして使用し、Web パイプを使用してデモ アプリケーションを作成することにすでに成功しています。
新しい言語の目的と使用例を早く理解できれば、それだけ良いでしょう。誰かにその言語を使って遊んでもらうのは早ければ早いほど良いです。新しい言語用の WASM ランタイム環境を作成するのはかつてないほど簡単になっているため、 Datafarm と呼ばれる私の別のプロジェクトの場合と同様に、ランディング ページの上部にインタラクティブ エディターを追加することで大きなメリットが得られます。
したがって、CLI ツールのように 1 つのランタイムだけをターゲットにしないでください。ブラウザ ランタイムもターゲットにします。
優れた診断が必要になります。コンパイル時から実行時、リンティングに至るまで、この問題にはキッチンのシンクを丸投げする必要があります。言語サーバーが必要であり、基礎となる診断ツールとの複数のインターフェイスが必要になります。そしてもちろん、これを支援するエージェント プログラミング ツールが多数あります。
私が遭遇した 1 つのパターンは、ランタイムおよび言語サーバーとして動作する単一のバイナリを作成することです。これにより、すべての診断フィードバックが 2 つの間でインラインに保たれます。さらに、LSP API から診断を分離することは、Monaco のようなブラウザーに埋め込み可能なコンポーネント用の WASM 診断ツールを提供できることを意味します。ツールやランタイムに関係なく、すべてのタイプミスや構文エラーの下に赤い波線が表示されます。
というのが私の意見ですw

リフトオフに必要な基礎をカバーすることがますます容易になるため、今後数年間で、特に DSL の種類の新しい言語が爆発的に増加するでしょう。

## Original Extract

How a New DSL Survives in the Era of LLMs

How a New DSL May Survive in the Era of LLMs - William Cotton
williamcotton.com
How a New DSL May Survive in the Era of LLMs
There have been an untold amount of projects written in Python, Rust, Ruby, and other "legacy" software languages over the last few decades. All of this code has made for great fodder for LLMs. As these models have progressed the instances of hallucinations have decreased dramatically.
But it's not just the mere volume of source code to train on. It's the advanced tooling around these languages. Type checkers, linters, language servers, compilers, interpreters, testing harnesses, you name it. All of these tools ground software in reality and give an LLM agent immediate feedback. For example, hallucinations are caught by the type checker before the code is even run.
Because of this even more of these legacy software languages are being used to create even more content for future models to train on.
It seems like we're stuck in a feedback loop of sorts.
So what's a new language to do to become viable in the era of LLMs?
The answer isn't too dissimilar to how things used to work: great documentation, great marketing, and great tooling. How does a prospective language user find out about a new language? How do they onboard? How does the language integrate with their existing tools and workflows? A modern language is going to need a robust language server. It is going to need an extensive set of documentation with a good onboarding flow.
A new language also going to need a little bit more to play well with LLM agents.
One approach is to have your new language create an AGENTS.md type file from the binary itself, something along the lines of:
Here's the LLM template used by Web Pipe , an experimental web application DSL that I've been working on lately.
It's a bit unfair because Web Pipe has a number of advantages that other DSLs might not have. It embeds other languages like jq, Lua, JavaScript, SQL and a few more. LLMs are already going to have familiarity with these languages so the syntax and semantics of the pipeline-oriented approach are not that much more to learn.
I've already had a lot of luck with one-shot prompts in codex to create demo applications using Web Pipe with just this single AGENTS.md template file as guidance.
The quicker you can get across the purpose and use case of a new language the better. The quicker you can get someone playing with the language the better. Since it has never been easier to create WASM runtime environments for new languages you can benefit immensely from adding an interactive editor right at the top of your landing page, as is the case with another project of mine called Datafarm .
So don't just target one runtime like a CLI tool. Target a browser runtime as well!
You're going to need great diagnostics. From compile time, to runtime, to linting, you name it, you're going to need to throw the kitchen sink at this problem. You will need a language server and you will need multiple interfaces with the underlying diagnostic tooling. And of course we've got a plethora of agentic programming tools out there to assist with this!
One pattern I've come across is to create a single binary that operates as the runtime and the language server. This keeps all of the diagnostic feedback inline between the two. In addition, separating the diagnostics from the LSP APIs means you can ship WASM diagnostic tools for a browser-embeddable component like Monaco. Red squiggles under all of your typos and syntactical errors, no matter the tool or runtime, FTW!
It's my opinion that we're going to see an explosion of new languages, especially of the DSL variety, over the next few years as it becomes easier and easier to cover the bases needed for liftoff.
