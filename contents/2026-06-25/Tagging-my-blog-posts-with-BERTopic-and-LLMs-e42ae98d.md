---
source: "https://vickiboykis.com/2026/05/18/tagging-my-blog-posts-with-bertopic-and-llms/"
hn_url: "https://news.ycombinator.com/item?id=48672707"
title: "Tagging my blog posts with BERTopic and LLMs"
article_title: "Tagging my blog posts with BERTopic and LLMs | ✰Vicki Boykis✰"
author: "surprisetalk"
captured_at: "2026-06-25T13:43:34Z"
capture_tool: "hn-digest"
hn_id: 48672707
score: 2
comments: 0
posted_at: "2026-06-25T13:00:06Z"
tags:
  - hacker-news
  - translated
---

# Tagging my blog posts with BERTopic and LLMs

- HN: [48672707](https://news.ycombinator.com/item?id=48672707)
- Source: [vickiboykis.com](https://vickiboykis.com/2026/05/18/tagging-my-blog-posts-with-bertopic-and-llms/)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T13:00:06Z

## Translation

タイトル: BERTopic と LLM を使用したブログ投稿のタグ付け
記事のタイトル: BERTopic と LLM を使用してブログ投稿にタグ付けする | ✰ヴィッキー・ボイキス✰
説明: LLM は、ループには依然として人間が必要ですが、別の部分に人間が必要であることを意味します

記事本文:
BERTopic と LLM でブログ投稿にタグを付ける | ✰ヴィッキー・ボーイキス✰ ✰ヴィッキー・ボーイキス✰
ブログ投稿に BERTopic と LLM をタグ付けする
最近、BERTopic と LLM の組み合わせを使用してブログにタグを追加しました。タグは右側のサイドバー (モバイルではフッター) に表示されます。以前 2023 年に GGUF Mistral で llama-cpp を使用してこれを実行しましたが、プロジェクトは完了しませんでした。現在、モデルは非常に良くなり、私のプロジェクトは小規模で比較的明確に定義されており、評価が簡単だったため、BERTopic、Claude Code、および Deepseek を備えた Pi を使用して、プロジェクトには 1 か月で約 6 ～ 10 時間かかりました。
なぜこれほど多くの異なる AI ツールがあるのでしょうか?主に、彼らの異なる働き方を評価するためです。その時間の多くは、タグ自体を反復するのではなく、タグの UX エクスペリエンスを検討することに費やされました。
最近の LLM の本当に便利な使用例の 1 つは、本番環境には影響せず、自分専用にカスタマイズされた小さな表面積を持つ個人プロジェクトを完了する場合です。言い換えれば、Robin Sloan 氏が書いたように、アプリは家庭料理になり得るのです。
コンテンツの作成と公開が簡単で、読み込みが速いため、私は静的サイトが大好きですが、もう少しフル機能があればいいのにと思うことがあります。 LLM により、検索などの機能を追加できるようになりました。私が使用しているテーマ Hugo Bear Blog はすでにタグをサポートしていますが、投稿にタグを追加したことがなかったので、タグを視覚化する少し異なる方法も必要でした。
私たち消費者は通常、テキストや画像に LLM を使用します。また、開発者であればコード生成にも LLM を使用します。しかし、LLM の最も過小評価されている機能の 1 つは、生成ではなく圧縮する機能です。これはまったく驚くべきことではありません。LLM は結局のところ、自然言語モデルです。
彼らはもともと言語モデリングタスクでトレーニングされ、微調整されていたため、

また、要約、情報検索、質問応答など、言語モデルが意図するすべてのタスクでも非常に優れたパフォーマンスを発揮します。 LLM は物事にラベルを付けるのが非常に得意です。つまり、彼らは、特にゼロショット コンテキスト (特にユーザーからの以前のトレーニング データがない場合) における、タグ付けの背後にある機械学習タスクであるトピック モデリングに優れています。
オンライン ブログの初期の頃、タグはコンテンツの発見を容易にするために重要でした。人々は当初、各自のブログに自分のブログ投稿にタグを付け始めました。最終的に、 Delicious のようなサイト アグリゲーターは、ユーザーが共有するトップ リンク全体のタグを集約することで、タグ付きのトップ リンクを表示するようになりました。
Pinboard は、タグを通じてコン​​テンツを検索し、集約されたタグを確認することがプラットフォームの重要な機能であるもう 1 つの著名なプラットフォームでした。偶然にも、Delicious は後に Pinboard に買収されました。
初期の頃、気に入ったものを見つけるための最良の方法は、これらのフォークソノミーのキュレーションに手動で参加することでした。 Twitter と Tumblr は、最も創造的なフォークソノミクスのタグ付けシステムをいくつか開発しました。 Tumblr では、タグは投稿を見つけるだけでなく、投稿について他の人と会話するための手段になりました。 Twitter では、Twitter の SimClusters アルゴリズムが実装される前に、ハッシュタグは共通の関心を持つ人々の共同発見を知らせる手段となりました。
タグ付けとハッシュタグは、10 年以上にわたり、Twitter、Instagram、TikTok、その他多くのサービスにわたって、プラットフォーム間でのコンテンツ発見の暗黙の契約として機能していました。しかし、ここしばらく、大手ソーシャルは大きく衰退している。グループ チャットは、従来のタグ付けメカニズムにあまり依存しない Discord グループと同様に、交換と発見のための媒体として登場しました。最近設立されたソーシャル プラットフォームである Bluesky は、

彼は投稿にハッシュタグを追加することができますが、ほとんどの人はそうしません。今日の多くのプラットフォームと同様に、スターター パックとカスタム アルゴリズム フィードを通じて発見が行われます。
LLM の台頭により、個々の Web サイトがシグナルを追加する力はさらに低下しました。要約や RAG を利用したソース合成を提供する検索結果の AI 概要機能の増加により、実際の Web サイトへのアクセスはこれまで以上に急速に減少しています。 LLM では、検出メカニズムとしてセマンティック検索と混合エージェント スタイルの検索が台頭しており、タグがそれほど重要ではなくなっています。
ブログ内では、(なんてことだ) LinkedIn のネイティブ投稿や X の長文投稿などのコンテンツ サーフェイスがプラットフォーム固有のロックインに貢献しています。いずれにしても、公開ブログのコンテンツはすべてトレーニング データとしてスクレイピングされており、エージェント検索と RAG により、人々はページに直接アクセスするのではなく、LLM の解釈を通じてコン​​テンツにアクセスすることになります。 RSS はまだ存在しますが、X や LinkedIn で記事を書けるようになったときに、誰がサイトをシンジケートするでしょうか? (もちろん私です。) 私のフィードを購読していただけます! 。私のブログの内容を現実的に理解するためのメカニズムとしてのブログタグは、おそらく私にとってのみ重要です。でも、それでも欲しいんです！
LDA を使用したタグ付けに対する歴史的なアプローチ
一般に、テキスト本文全体にわたるトピックの合成と検出は、教師なし学習の機械学習の問題です。カテゴリが何であるかは事前にはわかりません。コーパスからそれらを推測する必要があります。
従来、これは機械学習における非常に難しい統計問題であり、潜在ディリクレ割り当て (LDA) などのアプローチで解決されていました。コア実装の詳細に興味がある場合は、Ted の投稿を強くお勧めします。重要な考え方は、テキスト文書のコーパスが与えられた場合、各文書は次のようにモデル化されるということです。

トピックのコレクション。各トピックは単語全体の分布です。つまり、コーパス内の各単語がトピックのメンバーである確率が指定されています。
このアプローチおよびすべての統計的アプローチには、未解決の問題が数多くありました。最も大きな問題は、トピック モデリングの根底にある前提が、ドキュメントは単語の集まりであるということです。つまり、互いに独立した単語です。しかし、単語は構成文から切り離されるとコンテキストを失うため (モデルのトークン化された入力テキストを作成したときと同様)、モデルは、「銀行に行った」と「川の岸辺は水で溢れた」という 2 つの文の「銀行」という単語が互いに異なることを認識しません。
NLP で最もありきたりで使い古されている引用の 1 つは、「あなたは、それが管理する会社の言葉を知ることになるでしょう」です。これは、言語処理のための進化するニューラル アプローチの核となる成功の 1 つであることが判明しました。
2023 年の初期 LLM による埋め込みからタグ付けまで
エンベディングの使用法が進化したことにより、同じ多次元ベクトル空間に投影することで、意味的に互いに近いアイテムを数学的に表現することがはるかに簡単になりました。次の世代のコンテキスト エンベディングでは、単語の位置または表現に基づいてベクトルが計算され、最終的にはコンテキスト、つまり特定のテキスト内のすべての単語と他のすべての単語の関係を同時に完全にモデル化するモデルが作成され、ある単語が他の特定の単語とどの程度関連しているかの重み、別名アテンション メカニズムが作成されました。
トピック モデリング アプローチの最初の部分は、モデルが推定できるシード トピックを選択することです。これは、いくつかの選択したものからブートストラップして作成するなど、さまざまな方法で行うことができます。

tf-idf または LDA を介してドキュメント セット全体からもっともらしいトピックを選択し、すべてのコンテンツを調べて LLM でそれらをゼロショットすることで、より詳細なトピックを作成します。
これは、私がこの演習を初めて実行しようとしたときに実行しようとしたことです。 2023 年に Mistral 7B-Instruct のようなローカル モデルが利用可能になったとき、私はそれらをいじってみたところ、タグ付けがうまく機能することに気づきました。自分のブログにトピックをタグ付けできるかどうかを確認してみようと思いました。
最初の部分は、指定されたドキュメントから一連のタグを見つけ出すことでした。いくつか検索したところ、TopicGPT と呼ばれる論文を見つけました。この論文では、以前は BERT エンコーダー モデルによって実行されていたタスク、およびそれ以前は潜在ディリクレ割り当てによって実行されていたタスクを LLM を使用して実行する方法が既に模索されていました。
私は、llama-cpp-python を使用して、論文のテンプレートを Mistral 7B の GGUF クオントに適合させました (もう古い歴史です!)。
プロンプト テンプレートは非常に重要でしたが、現在ほどではありませんでした。コードは次のようになりました。
システム = f """
トピック階層からドキュメントとトップレベルのトピックのセットを受け取ります。
あなたのタスクは、階層からトップレベルのトピックとして機能するドキュメント内の一般化可能なトピックを特定することです。
ドキュメント内で識別されている既存のトップレベルのトピックを出力します。ドキュメントに新しいトピックがある場合、
それらを提案してください。
{ トピックス }
【例】
{ 例 }
【注意事項】
ステップ 1: 文書内で言及されているトピックを決定します。
- トピックのラベルは、可能な限り一般化可能である必要があります。ドキュメント固有のものであってはなりません。
- トピックは、トピックの組み合わせではなく、単一のトピックを反映する必要があります。
- 新しいトピックには、レベル番号、短い一般的なラベル、およびトピックの説明が必要です。
- トピックは、将来のサブトピックに対応できるほど広範である必要があります。
ステップ 2: 次の操作のいずれかを実行します。
1. すでに重複または関連するトピックがある場合

階層内でそれらのトピックを出力し、ここで停止します。
2. ドキュメントにトピックが含まれていない場合は、「なし」を返します。
3. それ以外の場合は、トピックをトップレベルのトピックとして追加します。ここで停止し、追加されたトピックを出力します。レベルを追加しないでください。
出力:
既存のトピック: ["前処理"、"LLM"]
new_topics: ["Python"]
「」
llama-cpp-python とモデルはどちらも (当時利用可能なものとしては) うまく機能しましたが、いくつかの制約がありました。
小さなコンテキスト ウィンドウ - Mistral 7B GGUF には 32,000 個のトークンがあり、プロンプトと長いドキュメントの両方をコンテキスト ウィンドウに完全に収める方法を決定する必要がありました。一般に、コンテキスト ウィンドウに収まるように圧縮するには、テキストを切り取るか、テキストの概要を作成してからその概要を要約する必要がありますが、どちらも特定の用語を検索する場合には理想的ではありません。プロセス中に発生する必要があるトークンのカウントを確認できます。
特に小規模なローカル モデルではメモリが不足したため、ドキュメントごとにモデルをループで実行するように再プロンプトし、それらのドキュメントの結果を追跡するのは面倒でした。
プロンプトとアウトラインの組み合わせを使用してカテゴリの保証された出力を制限することは、たとえアウトラインの助けを借りてでも困難でしたが、アウトラインのおかげで簡単になりました。補足: 論文を読んでください。とても素晴らしい内容です。
最初のレベルの分類トピックで初期パスを作成する適切な方法を決定できず、最終的にこのプロンプト アプローチを使用することになりましたが、それでも手動でシードする必要があり、適切なカテゴリは作成されませんでした。この問題は LLM に固有のものではなく、2005 年から遡って元のトピック論文で議論されている既知の問題です。
この変動の背後にある根本的な要因は、基本レベルの特異性が、そのような特異性を特定する程度に異なることである可能性があります。

氷は個人の生活に変化をもたらします。犬の専門家は、スキルだけでなく、たとえばビーグルとプードルを区別する必要性も備えています。専門知識の違いと同様に、他の社会的または文化的カテゴリーの違いは、基本レベルの違いを生み出す可能性があります。
ただし、システムをタグ付けする目的では、基本レベルの矛盾は悲惨な結果になる可能性があります。perl と javascript のタグが付いたドキュメントは一部のユーザーにとっては特殊すぎる可能性がある一方で、プログラミングのタグ付けされたドキュメントは他のユーザーにとっては一般的すぎる可能性があるためです。
タグ付けは基本的に意味を成すものです。センスメイキングとは、情報が分類およびラベル付けされ、重要なことに、それを通じて意味が現れるプロセスです (Weick、Sutcliffe、Obstfeld の近日公開予定)。 「基本レベル」は、人間がそれらのレベルのアイテムと対話する方法に関連していることを思い出してください (Tanaka & Taylor 1991)。人は外の世界と交流するとき、遭遇した物事を分類し、意味を与えることで理解します。しかし、実際には、カテゴリーは明確に定義されていないことが多く、その境界は曖昧です (Labov 1973)。アイテムは多くの場合、カテゴリ間に存在するか、複数のカテゴリに同様に存在します。人が最終的に自分のために描く線は、その人自身の経験、日々の実践、ニーズを反映しています。

[切り捨てられた]

## Original Extract

LLMs mean you still need a human in the loop, but in a different part

Tagging my blog posts with BERTopic and LLMs | ✰Vicki Boykis✰ ✰Vicki Boykis✰
Tagging my blog posts with BERTopic and LLMs
I recently added tags to my blog using BERTopic and a mix of LLMs. You can see the tags in the sidebar to the right (or in the footer on mobile). I’ve done this before in 2023, with GGUF Mistral using llama-cpp, but never finished the project. Now, because the models have been getting so good, and my project was small, relatively well-defined, and easy to evaluate, the project took me about 6-10 hours over a month, using BERTopic, Claude Code, and Pi with Deepseek.
Why so many different AI tools? Mostly to evaluate their different ways of working. Much of that time was spent noodling on the UX experience of the tags rather than iterating on the tags themselves.
One of the genuinely useful use-cases of LLMs these days is for finishing personal projects that don’t touch production and have a small surface area that’s personalized for you. In other words, as Robin Sloan wrote , an app can be a home-cooked meal.
I love having a static site because of how easy it is to write and publish content and how fast it loads, but sometimes I wish it was slightly more fully-featured. LLMs have allowed me to add features like search . The theme I use, Hugo Bear Blog , already has support for tags, but I’d never added them to posts, and I also wanted a slightly different way to visualize them.
We consumers generally use LLMs for text or image or, if we’re developers, for code generation. But, one of the most underrated features of LLMs is the ability to compress rather than generate. This is really unsurprising: LLMs are, after all, natural language models.
Since they were trained and fine-tuned originally on language modeling tasks they also perform really well at all the tasks that language models are meant for, such as summarization, information retrieval, question answering. LLMs are really good at labelling things. That is, they’re good at topic modeling, the machine learning task behind tagging, especially in a zero-shot context (where they have no previous training data from you specifically).
In the early days of online blogging, tags were important for facilitating content discovery. People initially started tagging their blog posts on their individual blogs. Eventually, site aggregators like Delicious , surfaced top links with tags by aggregating tags across top links shared by users.
Pinboard was another prominent platform where finding content through tags and looking at aggregated tags was an important feature of the platform. Coincidentally, Delicious was later acquired by Pinboard .
Early on, the best way to find things that you liked was to manually participate in the curation of these folksonomies. Twitter and Tumblr developed some of the most creative folksonomic tagging systems. On Tumblr , tags became a way to not only discover posts, but to have conversations with other people about your post. On Twitter, hashtags became a way to signal communal discovery of people with shared interests before the implementation of Twitter’s SimClusters algorithm .
Tagging and hashtags served as an implicit contract of content discovery across platforms for over a decade, across Twitter, Instagram, TikTok, and many other services. However, big social has been in big decline for a while now. Group chats have arisen as a medium for exchange and discovery , as well as Discord groups that rely less on traditional tagging mechanisms. Bluesky, a social platform that was founded more recently, has the ability to add hashtags to posts, but most folks don’t do so. Discovery happens, like with many platforms today, through starter packs and custom algorithmic feeds.
The rise of LLMs led to an even greater decrease in the power of individual websites to add signal. An increase in AI overview features in search results that offer either summarization or RAG-assisted source synthesis has meant that visits to actual websites are dropping faster than ever . With LLMs, the rise of semantic and blended agentic-style search as a discovery mechanism means tags are not as important.
Within blogging, content surfaces like (oh God) LinkedIn native posts and X’s longform posts are contributing to platform-specific lock-in. All of our public blog content is being scraped as training data anyway, and agentic search and RAG mean people access content through an LLM’s interpretation of it rather than going directly to a page. RSS still exists, but who is going to syndicate a site when they could write an article on X or LinkedIn? (Me, sure.) You can subscribe to my feed! . Blog tags as a mechanism for understanding what my blog is about realistically still probably matter only to me. But I still want them!
Historical approaches to tagging with LDA
Generally, synthesizing and detecting topics across a body of text is an unsupervised learning machine learning problem. We don’t know ahead of time what the categories are. We have to infer them from the corpus.
Traditionally, this was a really hard statistical problem in machine learning solved with approaches like latent dirichlet allocation (LDA) . I really recommend Ted’s post if you’re interested in more of the core implementation details. The key idea is that, given a corpus of text documents, each document is modeled as a collection of topics. Each topic is a distribution over words, which means that each word in our corpus has a stated probability for being a member of a topic.
There were a number of unsolved problems with this approach, and with all statistical approaches. The biggest one was that the underlying assumption of topic modeling is that documents are bags of words. That is, words that are independent of each other. But, because words lose their context when they’re broken up from their constituent sentences (as we did when we created the tokenized input text for the model), the model wouldn’t recognize that the word “bank” in these two sentences, with “I went to the bank” and “The bank of the river overflowed with water” are different from one another.
One of the most trite and overused quotes in NLP is “You shall know a word by the company it keeps”, and this turned out to be one of the core successes of evolving neural approaches for language processing.
From embeddings to tagging with early LLMs in 2023
The evolving use of embeddings made it much easier to represent items that were semantically close together mathematically, by projecting into the same multi-dimensional vector space. The following generation of contextual embeddings calculated vectors based on the position or representation of the word, which eventually led to models that fully modeled the context - or the relationship of every word to every other word in a given text at the same time, creating weights of how relevant a word is to any other given word, aka the attention mechanism.
The first part of any topic modeling approach is picking your seed topics that the model can extrapolate from. You can do this in a number of ways: by bootstrapping from a couple of selected ones and creating more detailed ones, by picking plausible topics across your document set through tf-idf or LDA, by going through all the content and having an LLM zero-shot them.
This is what I tried to do the first time I tried to do this exercise. When local models like Mistral 7B-Instruct became available in 2023, I started playing around with them and realized that they did well at tagging. I decided I wanted to see if I could tag my blog with topics.
The first part was figuring out a set of tags from the given documents. I did some searching and found a paper called TopicGPT which was already exploring ways to use LLMs to do tasks that were previously done by BERT encoder models, and prior to that, latent dirichlet allocation.
I adapted the template in the paper with GGUF quants of Mistral 7B using llama-cpp-python (ancient history by now!).
Prompt templates were really important, less so than now, and the code looked like this:
system = f """
You will receive a document and a set of top-level topics from a topic hierarchy.
Your task is to identify generalizable topics within the document that can act as top-level topics from the hierarchy.
Output the existing top-level topics as identified in the document. If the document has new topics,
suggest those.
{ topics }
[Examples]
{ examples }
[Instructions]
Step 1: Determine topics mentioned in the document.
- The topic labels must be as GENERALIZABLE as possible. They must not be document-specific.
- The topics must reflect a SINGLE topic instead of a combination of topics.
- The new topics must have a level number, a short general label, and a topic description.
- The topics must be broad enough to accommodate future subtopics.
Step 2: Perform ONE of the following operations:
1. If there are already duplicates or relevant topics in the hierarchy, output those topics and stop here.
2. If the document contains no topic, return "None".
3. Otherwise, add your topic as a top-level topic. Stop here and output the added topic(s). DO NOT add any additional levels.
Output:
existing_topics: ["Preprocessing", "LLMs"]
new_topics: ["Python"]
"""
Although both llama-cpp-python and the model worked well (for what was available at the time), I was constrained by a couple things:
A small context window - Mistral 7B GGUF had 32k tokens which meant that I had to make decisions about how to fit both my prompt and the longer documents entirely into the context window. Generally compacting to fit into a context window involves either cutting off text or creating a summary of your text and then summarizing that summary, neither of which are ideal when you’re looking for specific terms. You can see the token counting that has to happen in the process.
It was annoying to deal with reprompting the model to run in a loop for every document and keeping track of the results of those documents, especially as the smaller, local model ran out of memory.
Constraining guaranteed outputs for categories using a combination of the prompt and Outlines was finicky, even with the help of Outlines, which made it easier . Side note: Read the paper, it’s very cool.
I couldn’t decide on the proper method to create an initial pass at my first-level taxonomy topics and ended up doing it with this prompting approach, but still had to seed it manually, which didn’t create good categories. This problem isn’t specific to LLMs and is a known problem that’s discussed in the original topic paper, way back from 2005 :
The underlying factor behind this variation may be that basic levels vary in specificity to the degree that such specificity makes a difference in the lives of the individual. A dog expert has not only the skill but also the need to differentiate beagles from poodles, for example. Like variation in expertise, variations in other social or cultural categories likely yield variations in basic levels.
For the purposes of tagging systems, however, conflicting basic levels can prove disastrous, as documents tagged perl and javascript may be too specific for some users, while a document tagged programming may be too general for others.
Tagging is fundamentally about sensemaking. Sensemaking is a process in which information is categorized and labeled and, critically, through which meaning emerges (Weick, Sutcliffe & Obstfeld forthcoming). Recall that “basic levels” are related to the way in which humans interact with the items at those levels (Tanaka & Taylor 1991); when one interacts with the outside world, one makes sense of the things one encounters by categorizing them and ascribing meaning to them. However, in practice, categories are often not well defined and their boundaries exhibit vagueness (Labov 1973). Items often lie between categories or equally well in multiple categories. The lines one ultimately draws for oneself reflect one’s own experiences, daily practices, needs

[truncated]
