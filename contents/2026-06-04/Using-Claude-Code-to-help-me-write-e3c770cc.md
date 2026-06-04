---
source: "https://andrewpwheeler.com/2026/03/20/using-claude-code-to-help-me-write/"
hn_url: "https://news.ycombinator.com/item?id=48403318"
title: "Using Claude Code to help me write"
article_title: "Using Claude Code to help me write | Andrew Wheeler"
author: "apwheele"
captured_at: "2026-06-04T21:38:52Z"
capture_tool: "hn-digest"
hn_id: 48403318
score: 3
comments: 0
posted_at: "2026-06-04T19:17:05Z"
tags:
  - hacker-news
  - translated
---

# Using Claude Code to help me write

- HN: [48403318](https://news.ycombinator.com/item?id=48403318)
- Source: [andrewpwheeler.com](https://andrewpwheeler.com/2026/03/20/using-claude-code-to-help-me-write/)
- Score: 3
- Comments: 0
- Posted: 2026-06-04T19:17:05Z

## Translation

タイトル: クロード コードを使用して執筆を支援する
記事のタイトル: クロード コードを使用して執筆を支援する |アンドリュー・ウィーラー
説明: LLM を使用して作成を支援することは、多くの人にとって当然のことながら扱いにくいテーマです。 LLM ツールに考えさせて表面的には問題ないが、最終的にはゴミの散文を書くのは非常に簡単なので、現在ではかなりの AI のスロップが出てきています。私の最近の著書『LLMs for Mortals』では、ソネットを使用しました…

記事本文:
コース
Advanced Criminology (Undergrad) Crim 3302
Communities and Crime (Undergrad) Crim 4323
Crim 7301 – UT ダラス – 犯罪学の研究と分析のセミナー
Crime Science (Graduate) Crim 7381
GIS in Criminology/Criminal Justice (Graduate)
Crime Analysis (Special Topics) – Undergrad
Using Claude Code to help me write
LLM を使用して執筆を支援することは、多くの人にとって当然のことながら扱いにくいテーマです。 LLM ツールに考えさせて表面的には問題ないが、最終的にはゴミの散文を書くのは非常に簡単なので、現在ではかなりの AI のスロップが出てきています。
私の最近の著書『LLMs for Mortals』では、Sonnet 4.1 を使用して本の最初の草稿を書きました (約 5 ドル)。前の本は約 1 年かかりましたが、この本は約 2 か月で読み終えることができました。私は間違いなく大量のコピー編集を行いました (おそらく 1 章あたり平均約 20 ～ 30 時間) が、本の内容の約 50% はソネットが生成したオリジナルの散文だと思います。
LLM はツールです。使い方が下手かもしれませんが、かなり上手に使えると思います。 AI の書き込みを検出するために使用されるツールである Pangram は、モータルズ用の LLM 内のどのパッセージも AI が生成したものとしてフラグを立てません。
このブログ投稿では、クロード コードを使用して執筆を支援した方法についてのメモを説明します (ただし、実際には、Codex や Gemini などの現在のコーディング ツールのいずれにも当てはまります)。メタリファレンスとして、このブログ投稿は 100% 私自身が直接書いたものですが、参照フレームとして投稿の後半でクロード コードを使用して書かれたドラフトにリンクします。
まず、LLM に直接書いてもらうことに同意できないとしても、比較的議論の余地のない使用例があります。それは、LLM にあなたの作品のコピー編集パスを取得させることです。
これは私が最近これを使用した例です。Crime De-Coder に関するブログ投稿は非常に有益です。

API とローカル LLM の使用の適合性。この会話では、私のオリジナルの草案と、クロードのデスクトップ ツール (無料版) が提供した提案を見ることができます。
繰り返しますが、これは実際にはクロードに固有のものではありません (これは ChatGPT でも同様にうまく機能したはずです)。 LLM は、スペル ミスだけでなく、スペル チェックでは検出できない文法上の問題や、コンテンツに関する一般的なコピー編集のアドバイスにも適しています。
この点の 1 つは、私の設定を再現するには、プレーン テキストで記述する必要があるということです。私が書いているもののほとんどは何らかのマークダウン形式です (ブログ投稿の場合は単純なマークダウン、長いレポート/書籍などの場合は Quarto)。これにより、ツール、特に Claude Code などのコマンド ライン インターフェイス (CLI) ツールの使用がはるかに簡単になります。
現在、LLM の作成には 2 つの大きな問題があります。
現在の LLM の書き方には、それ自体が顕著になりつつある特定のスタイルがあります
最初の箇条書きでは、何が書かれているかを確認する必要があります。自分が専門とするコンテンツについて書いてもらうほうがはるかに簡単なので、レビューして間違いを見つけるのも簡単です。 (これは、コンピューター コードの作成を支援するツールの使用に関する現在の問題と同じです。ツールは上級者にとっては恩恵ですが、大量の駄文を作成する可能性があるため、初心者のプログラマーは問題を見つけるのが困難です。)
2 番目の項目、自分のスタイルを模倣することについて、ここで説明します。生成 AI LLM がどのように機能するかを高いレベルで理解することは価値があります。「質問 X に答えてください」と尋ねた場合と「ここに本があります、…、質問 X に答えてください」と尋ねた場合、LLM は異なる応答を生成します。前者のプロンプトの最初の部分、「ここに本があります、…」はコンテキストを参照しているものです。現在のモデルには、約 500,000 ワードのコンテキスト ウィンドウ (潜在的な入力の大きさ) があります (技術的には約 100 万トークンであり、1 つのワードは複数であることがよくあります)

ただしトークン）。
通常、コンテキスト ウィンドウを 100% 埋める必要はありませんが、500,000 ワードは非常に大きな数です。テキストを含めるだけでも数冊の本になります。もう 1 つの一般的なプロンプト手法は、k ショット例と呼ばれるものです。通常は次のようになります
入力例 1: ...テキスト... 期待される出力: ...まあ...
例 input2: ...diff text... Expected_output: ...blah2...
....
これをコンテキスト ウィンドウに配置し、通常のプロンプトを送信して、LLM にコンテンツを生成させます。最終出力がどのようになるかを LLM にガイドするのに役立つ以前の例が示されています。これはライティングでも同様に機能します。LLM に以前のライティングの例を提供して、将来のスタイルを模倣できるようにします。
わかりやすくするために、github に例を作成しました。基本的には、以前に書いたもの (テキストで!) を用意して、Claude Code に次のような質問をするだけです。
/blogposts フォルダー内の私の以前のブログ投稿を確認してください。テキストを確認した後、概要を踏まえてトピック X に関する新しいブログ投稿を書いてもらいます。
次に、以前の作業がコンテキスト ウィンドウに表示された後、書きたい内容のアウトラインを LLM に入力します。この例では、アウトラインを実際のテキスト ファイルに置き、次のように述べています。
ClaudeWritingPost フォルダーで、outline.txt を確認し、新しい md を作成します。
ClaudeWritingExample.md という名前のファイルに、以下に基づいてセクションを埋めます。
概要
次に、Claude Code はアウトラインを含むテキスト ファイルを確認して投稿を作成します。 github リポジトリには、この同じ投稿の元のアウトラインがあるので、並べて見ることができます。
技術的には、Claude Code (または他の CLI ツール) を使用してカスタム コマンドとスキルを記述して、2 つのプロンプトを入力する手順を省略できますが、ユーザーにとってわかりやすくするために、ここでは手動で 2 つの手順を示しているだけです。本当にその2ステップだけです

– 前の例をコンテキスト ウィンドウに表示し、他に何を書くべきかのアウトラインを入力します。
Github リポジトリでは、追加の Claude.md ファイルがいくつか表示されます。これらのファイルには追加の手順が含まれています。私がよく言うのは「絵文字を含めない」です。 LLM の記述は冗長になり、リストが過剰になる傾向があります。したがって、それらを避けるための指示もあります。
書かれたブログ投稿は悪くありません。概念実証として読んでみることをお勧めします (セッションをエクスポートしました。費用は約 50 セントです)。私が通常、ブログ投稿について心配しない理由の 1 つは、執筆の過程で頻繁に内容を追加したり変更したりするためです。したがって、私が個人的に書いた投稿は長く、さらにいくつかの要素が含まれていることがわかります。
では、いつ使用しますか? Python でチュートリアルを書くようなテクニカル ライティングは非常にうまく機能します。したがって、LLM ブックの最初のパスを書き込んで、コンテンツの 50% を保持させることができます。将来的にはブログ投稿に使用するかもしれません (毎日何かを書く必要があると感じた場合)。しかし、今のところはそのような思い切った行動をとるつもりはありません。
論文全体や本などの長い文章の場合は、詳細な概要を作成するだけでなく、LLM にそれを小さなセクションに分けて書いてもらうことをお勧めします。これは、コンテンツをレビューするだけでなく、編集/変更を加えた場合に LLM を軌道に乗せるのにも役立ちます。 (会話が長くなると、パフォーマンスが低下し、エラーが繰り返される可能性が高くなります。)
私はもう学術論文をあまり書いていませんが、LLM 執筆に関するもう 1 つの根本的な問題は、引用の幻覚です。テキストのマークダウン ファイルで記述する場合、私の提案は非常にシンプルです。引用したい論文を bibtex ファイルに記述し、マークダウンのインラインで、次の形式で論文のみを引用します。
引用すると、@item1 はなんとも言えません [@item1; @item2]。特定のページの引用の場合 [@item1

p. 34-35]。
私のアウトラインの書き方は、典型的には、論文 a、b、c を引用して、X についての段落を書くようなものです。したがって、アウトラインを徐々に埋めていく私の個人的なスタイルは、LLM とうまく機能します。
したがって、これは、あなたがすでに論文のリストを持っていることを前提としています (そして、LLM を使用して、まだ読んでいない論文に基づいて lit レビューを動的に書いているわけではありません)。次回、実際に学術論文を書く必要があるときは、Semantic Sc​​holar の API にクエリを実行して、優れた bibtex ファイルを作成するための MCP ツールを作成するかもしれません。
ただし、ここでの解決策は、やはり出力の正確性を確認する必要があるということです。これらのツールを持たない人は怠け者で、まだ読んでいないものを引用するため、今後もそのようなことが起こり続けるでしょう (ツールはそれを簡単にするだけです)。ただし、ツールを適切に使用する方法を理解している人は、より生産的な文章を書くことができます。
X で共有 (新しいウィンドウで開きます)
×
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
LinkedIn で共有 (新しいウィンドウで開きます)
リンクトイン
Pinterest で共有 (新しいウィンドウで開きます)
ピンタレスト
アンディ・ウィーラーによる投稿、2026 年 3 月 20 日
https://andrewpwheeler.com/2026/03/20/using-claude-code-to-help-me-write/
パングラムは良いです |アンドリュー・ウィーラー
実験を行う期間: ASEBP で話してください
テクノロジーコースへの関心を集める
LinkedIn プレミアムでは投稿が増えません
メール アドレスを入力してこのブログをフォローし、新しい投稿の通知をメールで受け取ります。
クロード コードを使用して執筆を支援する
Stata のグループベースの軌道モデル - いくつかのグラフと適合統計
Microsoft Wordで数式を書く
購読する
購読しました
アンドリュー・ウィーラー
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

Using LLMs to help you write is understandably a touchy subject for many. There is quite a bit of AI slop coming out now, as it is really easy to just have the LLM tools think for you and write superficially OK but ultimately garbage prose. My recent book, LLMs for Mortals, I used Sonnet…

Courses
Advanced Criminology (Undergrad) Crim 3302
Communities and Crime (Undergrad) Crim 4323
Crim 7301 – UT Dallas – Seminar in Criminology Research and Analysis
Crime Science (Graduate) Crim 7381
GIS in Criminology/Criminal Justice (Graduate)
Crime Analysis (Special Topics) – Undergrad
Using Claude Code to help me write
Using LLMs to help you write is understandably a touchy subject for many. There is quite a bit of AI slop coming out now, as it is really easy to just have the LLM tools think for you and write superficially OK but ultimately garbage prose.
My recent book, LLMs for Mortals , I used Sonnet 4.1 to write the initial draft of the book (for around $5). My prior book took around a year, whereas I was able to finish this book in around two months. I definitely did a ton of copy-editing (maybe around 20-30 hours per chapter on average), but I believe around 50% of the book material is the original Sonnet generated prose.
LLMs are a tool – they can be used poorly, but I think they can be used quite well. Pangram, a tool used to detect AI writing, does not flag any of the passages in LLMs for Mortals as AI generated .
This blog post goes over my notes on how I used Claude Code to help me write (although it really is applicable to any of the current coding tools, like Codex or Gemini as well). As a meta-reference, this blog post is 100% written by myself directly, but I will link to a draft written using Claude Code later in the post for a frame of reference.
First, even if you do not agree with having an LLM write for you directly, there is a use case that should be relatively uncontroversial – having an LLM take a copy-edit pass on your work.
Here is an example I used this for recently, the blog post on Crime De-Coder goes over the benefits of using an API vs local LLMs . In this conversation, you can see my original draft , and the suggestions that Claude’s desktop tool (the free version) gave.
Again this is not really specific to Claude (this would have worked fine in ChatGPT as well). LLMs are good for not only spelling errors, but grammatical issues that spell check will not catch, as well as just more general copy-editing advice on the content.
One point of this – to replicate my setup, you need to write in plain text. Most of the things I write are in some form of markdown (plain markdown for blog posts, and Quarto for longer reports/books/etc). This makes it much easier to use the tools, especially the command line interface (CLI) tools like Claude Code.
There are two big issues currently with LLM writing:
current LLM writing has a particular style that is itself becoming noticeable
The first bullet, you need to review what it writes. It is much easier to have it write on content you are an expert in, so it is easier to review and spot errors. (It is the same current problem with using the tools to help you write computer code – they are boons for seniors but can write a ton of slop that more neophyte coders have a hard time spotting issues.)
The second bullet, having the style mimic your own, is what I am going to discuss here. It is worth understanding at a high level how generative AI LLMs work – if you ask “answer question X” vs “here is a book, …., answer question X” the LLM will generate a different response. The first part in the former prompt, “here is a book, …” is what is referred to the context . Current models have context windows (how large of a potential input) at around 500,000 words (technically they are around 1 million tokens, one word is often multiple tokens though).
You generally do not want to fill up the context window 100%, but 500,000 words is a very large number – just including text it would be multiple books. Another common prompting technique is what is called k-shot examples. It will typically go like
example input1: ...text... expected_output: ...blah...
example input2: ...diff text... expected_output: ...blah2...
....
This is what you place in the context window, then submit your usual prompt, and have the LLM generate the content. It is giving prior examples to help guide the LLM what you expect the final output to look like. This works the same way with writing – give the LLM prior examples of your writing to help it mimic your future style.
To keep it simple, I have created an example on github to follow along . Basically just have your prior writing (in text!), and then ask Claude Code something like:
review my prior blog posts in folder /blogposts, I am going to have you write a new blog post on topic X given the outline *after* you review the text
Then after your prior work is in the context window, feed the LLM an outline for what you want to write. In this example, I put the outline in an actual text file and said:
In the ClaudeWritingPost folder, review the outline.txt, then create a new md
file, called ClaudeWritingExample.md, filling in the sections based on the
outline
Claude Code will then go and review the text file with the outline and write the post. In the github repo I have my original outline for this same post , so you can see side-by-side.
You can technically write custom commands and skills with Claude Code (or the other CLI tools) to save the steps of typing two prompts, but to keep it simple for folks I am just showing the two steps manually. It is really just those two steps – get your prior examples into the context window, and then feed an outline for what else to write.
In the Github repo you can see some additional Claude.md files – these are files that include additional instructions. A common one I say is “do not include emojis”. LLM writing also tends to be verbose and have excessive lists. So I have instructions to avoid those as well.
The written blog post is not bad – I would suggest to go and read it as a proof of concept (I exported the session, can see it cost around fifty cents). Part of the reason I do not typically worry about blog posts is that I often add in things/change things in the process of writing. So you can see my personally written post is longer and has a few more elements.
So when would you use it? Technical writing, like writing tutorials in python, it works very well. Hence I could have it write the first pass on my LLM book and keep 50% of the content. I may use it for blog posts in the future (if I felt compelled to write something every day). But will not take that plunge for now.
For longer pieces, like an entire paper or a book, I suggest to not only make a detailed outline, but to also have the LLM write it in smaller sections. This both helps with reviewing the content, as well as to keep the LLM on track if you make edits/changes as you go. (Longer conversations it is more likely to degrade and make repeated errors.)
I am not writing academic papers much anymore, but another fundamental problem with LLM writing is hallucinating citations. If you write in text markdown files, my suggestion is fairly simple – have the papers you want to cite in a bibtex file, and in-line in markdown, only cite papers in the form:
Citation, @item1 says blah [@item1; @item2]. For a specific page quote [@item1 p. 34-35].
The way I write my outlines, it typically is like write a paragraph about X, cite papers a,b,c . So my personal style of progressively filling in an outline works well with LLMs.
So this presumes you already have a list of papers (and are not using the LLM to dynamically write your lit review based on papers you have not read ). Next time I actually need to write an academic paper, I may write up an MCP tool to query Semantic Scholar’s API and create a nice bibtex file.
But the solution here is again you need to review the output for accuracy. People without these tools are lazy and cite things they have not read already, so that will continue to happen (the tools just make it easier). Those that figure out how to use the tools appropriately though can be much more productive writing.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Share on LinkedIn (Opens in new window)
LinkedIn
Share on Pinterest (Opens in new window)
Pinterest
Posted by Andy Wheeler on March 20, 2026
https://andrewpwheeler.com/2026/03/20/using-claude-code-to-help-me-write/
Pangram is good | Andrew Wheeler
How long to conduct your experiment: Talk at ASEBP
Gathering interest in tech courses
LinkedIn Premium Does Not Boost your Posts
Enter your email address to follow this blog and receive notifications of new posts by email.
Using Claude Code to help me write
Group based trajectory models in Stata - some graphs and fit statistics
Writing equations in Microsoft Word
Subscribe
Subscribed
Andrew Wheeler
Already have a WordPress.com account? Log in now.
