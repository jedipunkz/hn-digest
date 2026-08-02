---
source: "https://ankursethi.com/blog/prevent-cognitive-debt-by-manually-retyping-llm-generated-code/"
hn_url: "https://news.ycombinator.com/item?id=49146214"
title: "Prevent cognitive debt by manually retyping LLM-generated code"
article_title: "Prevent cognitive debt by manually retyping LLM-generated code — Ankur Sethi's Lab Notebook"
author: "philipportner"
captured_at: "2026-08-02T17:54:10Z"
capture_tool: "hn-digest"
hn_id: 49146214
score: 2
comments: 0
posted_at: "2026-08-02T17:00:36Z"
tags:
  - hacker-news
  - translated
---

# Prevent cognitive debt by manually retyping LLM-generated code

- HN: [49146214](https://news.ycombinator.com/item?id=49146214)
- Source: [ankursethi.com](https://ankursethi.com/blog/prevent-cognitive-debt-by-manually-retyping-llm-generated-code/)
- Score: 2
- Comments: 0
- Posted: 2026-08-02T17:00:36Z

## Translation

タイトル: LLM で生成されたコードを手動で再入力することで認知的負債を防ぐ
記事のタイトル: LLM で生成されたコードを手動で再入力することで認知的負債を防ぐ — Ankur Sethi の Lab Notebook

記事本文:
アンクル・セティの研究ノート
ホーム
LLM で生成されたコードを手動で再入力することで認知的負債を防ぐ
4 月に述べたにもかかわらず、私はまだ個人プロジェクトでコーディング アシスタントを使用しています。
機能全体をワンショットで実行するためにこれらを使用すると、満足できず混乱してしまいますが、プロジェクトの退屈な部分を早送りするためにこれらを使用するのは楽しいです。
しかし、プロジェクト内でコーディングアシスタントを自由に歩き回らせると、膨大な量の認知的負債が残ります。 Web サイトにタグ付けを追加する方法を理解するために Django のドキュメントをじっくり読むという考えは嫌いかもしれませんが、基本的にはタグ付けがどのように機能するかを理解したいと思っています。問題が退屈だからといって、解決策についての理解を完全にマシンに任せたいわけではありません。
もちろん、LLM が生成するコードを 1 行ずつレビューすることもできます。それが、この呪われた 2026 年にほとんどの開発者が行うことを期待されているものです。ロボットが PR を作成し、人間が PR をレビューします。それは素晴らしい新しい世界です。
しかし、AI が生成した PR をレビューするのは好きではありません。過度に防御的で、ひどいコメントがあり、微妙に間違っているコードを何百行もじっくりと読むのは楽しいことではありません。雇用主のためにはしぶしぶそうするかもしれませんが、その雇用主ができるだけ早く元雇用主になるように配慮しながらも、個人的なプロジェクトのためには絶対にやらないでしょう。個人的なプロジェクトは何よりも楽しくなければなりません。個人的なプロジェクトに取り組む喜びは、結果ではなくプロセスから得られます。
それで、男の子は何をすればいいでしょうか？自分の作業と認識の制御をスロップマシンに譲らずに、退屈な作業を LLM にオフロードするにはどうすればよいでしょうか?
私は、ひどく非効率で、おそらく少し滑稽な解決策を思いつきました。コーディング アシスタントにチャットでコードを生成するよう依頼し、すべての編集を自分で手動で行います。
これらの指示はすべてのエージェントのファイルにあります

私の個人的なプロジェクトでは次のようになります。
このプロジェクトに含まれるコードのすべての行を理解したいと考えています。私から明示的に要求しない限り、プロジェクト ファイルを作成、編集、移動、名前変更、または削除しないでください。代わりに、チャット内で提案された編集をすべて表示して、手動で入力できるようにします。
私が明示的にアクションを要求しない限り、プロジェクト ファイルの変更、依存関係のインストール、またはリポジトリの状態の変更を行うコマンドを実行しないでください。代わりに、チャットでこれらのコマンドを表示して、手動で実行できるようにしてください。
私は経験豊富な開発者です。明示的に求められない限り、構文、API、プログラミングの概念、実装の詳細については説明しないでください。
この方法で LLM を使用すると、LLM をまったく使用しない場合よりも作業が速くなりますが、それでもマシンに考えさせようとする人よりは遅いです。 10 倍ではなく、おそらく 2 倍しか速くありません。しかし、速度という点では失われますが、コードをより深く理解するという点では得られます。
LLM で生成されたコードを 1 行ずつ手動でエディターに入力しながら、それがどのように機能し、既存のコードベースに適合するかについてのメンタル モデルを構築します。 API やアルゴリズムが理解できない場合は、立ち止まって調べたり、LLM に説明を求めたりすることができます。
自分でコードを入力すると速度が低下するため、LLM が行った可能性のある幻覚や不適切な設計選択を検出する可能性が高くなります。コードを再編成、リファクタリング、コメントの追加など、作業を進めながらクリーンアップして、自分の好みに合わせて調整することができます。
最も重要なのは、このワークフローによりコードベースの空間マップを構築できることです。私はあらゆる機能がコードベースのどこに存在するかを知っています。変更が必要なとき、どこを変更する必要があるのか​​が正確にわかります。これにより、プロジェクト内での作業が迅速化されるだけでなく、LLM へのプロンプトや指示がより簡単になります。

未来。
私が 10 代の頃にコーディングを学んでいたとき、経験豊富なプログラマーは、コードを決してコピーしてプロジェクトに貼り付けないようにとよく言いました。本で学習している場合は、すべての例をコンピューターにコピーして、実行できることを確認するようにアドバイスされました。ブログ投稿やフォーラムの回答から学んでいる場合は、完全に理解できるように、それを入力してコードベースに適応させるようにアドバイスされました。
LLM で生成されたコードをコードベースに手動で入力するのは、まったく同じ学習プロセスのように感じられます。これは LLM を使用する最も効率的な方法ではないかもしれませんが、私は生産性よりも理解を重視しています。私はこれを数か月間続けていますが、私にとってはうまくいきました。私はできる限りこのワークフローを使い続けるつもりです。
ソフトウェア業界は多額の認知的負債を抱えており、近いうちに返済しなければならないのではないかと心配しています。デジタルインフラの大部分がどのように組み立てられているかを理解できなくなる日が来るでしょう。私個人として業界全体の方向性を変えることはできないかもしれませんが、少なくとも自分が世に送り出したソフトウェアを完全に理解することはできます。それ以外の場合は業務上の過失になります。

## Original Extract

Ankur Sethi's Lab Notebook
Home
Prevent cognitive debt by manually retyping LLM-generated code
Despite what I said in April , I'm still using coding assistants on my personal projects.
Using them to one-shot entire features leaves me unsatisfied and disoriented, but I do enjoy using them to fast-forward through the boring parts of my projects.
However, allowing my coding assistant to roam free in my projects leaves me with a colossal amount of cognitive debt. I might hate the idea of poring over the Django documentation to figure out how to add tagging to my website, but I still fundamentally want to understand how it works. Just because a problem is boring doesn't mean I want to fully offload my understanding of the solution to a machine.
Of course, I could review every single line of code the LLM produces. That's what most developers are expected to do in this cursed year of 2026. Robots raise PRs, humans review them. It's a brave new world.
But I don't enjoy reviewing AI-generated PRs. Poring over hundreds of lines of overly-defensive, badly-commented, subtly incorrect code is not fun. I might grudgingly do it for an employer—while making sure said employer becomes an ex-employer as soon as possible—but I'm sure as hell not doing it for my personal projects. Personal projects must be fun above all else. The joy of working on personal projects comes from the process, not from the outcome .
So what's a boy to do? How do I offload the boring work to LLMs without ceding control of my own work and cognition to the slop machine?
I've come up with a solution that's grossly inefficient and perhaps slightly comical: I ask my coding assistant to generate code in the chat, then manually make all the edits myself.
I have these instructions in all the agents files in my personal projects:
I want to understand every line of code that goes into this project. Never create, edit, move, rename, or delete project files unless I explicitly ask you to do so. Instead, show me every proposed edit in the chat so I can type it in manually.
Do not run commands that modify project files, install dependencies, or change repository state unless I explicitly request that action. Instead, show me those commands in the chat so I can run them manually.
I'm an experienced developer. Do not explain syntax, APIs, programming concepts, or implementation details unless explicitly asked.
Using LLMs this way allows me to work faster than not using LLMs at all, but I'm still slower than those who are willing to allow the machine to think for them. Instead of being 10x faster, I'm probably only 2x faster. But what I lose out on in terms of speed, I gain in terms of a deeper understanding of my code.
As I manually type every single line of LLM generated code into my editor, I build up a mental model of how it works and fits into my existing codebase. If I don't understand an API or algorithm, I can stop to look it up, or just ask the LLM to explain it.
Typing the code myself forces me to slow down, which means I'm more likely to detect hallucinations or bad design choices the LLM might have made. I can clean up the code as I go, reorganizing it, refactoring it, adding comments, and generally adapting it to my own taste.
Most importantly, this workflow allows me to build a spatial map of my codebase. I know where every bit of functionality lives in the codebase. When I need to make a change, I know exactly where I need to make it. It not only helps me work faster within my projects, it also makes it easier for me to better prompt and instruct the LLM in the future.
When I was learning to code as a teenager, experienced programmers would often tell me to never copy and paste code into my projects. If I was learning from a book, I was advised to copy all the examples into my computer and make sure I could run them. If I was learning from a blog post or forum answer, I was advised to type it out and adapt it to my codebase so I understood it completely.
Manually typing LLM-generated into my codebase feels like the exact same learning process. It might not be the most efficient way to work with an LLM, but I value comprehension over productivity. I've been doing this for a few months now, and it's been working well for me. I plan to continue using this workflow for as long as I can.
I fear the software industry is taking on a large amount of cognitive debt that we'll have to pay back very soon. There will come a time when we no longer understand how large parts of our digital infrastructure are put together. I might not personally be able to change the course of the entire industry, but I can at least make sure I completely understand the software I put out into the world. Anything else would be professional malpractice.
