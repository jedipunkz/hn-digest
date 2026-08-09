---
source: "https://swiftrocks.com/using-ai-without-losing-critical-thinking"
hn_url: "https://news.ycombinator.com/item?id=49232014"
title: "Using AI without losing your critical thinking"
article_title: "Using AI without losing your critical thinking"
author: "rockbruno"
captured_at: "2026-08-09T15:21:35Z"
capture_tool: "hn-digest"
hn_id: 49232014
score: 2
comments: 1
posted_at: "2026-08-09T14:58:29Z"
tags:
  - hacker-news
  - translated
---

# Using AI without losing your critical thinking

- HN: [49232014](https://news.ycombinator.com/item?id=49232014)
- Source: [swiftrocks.com](https://swiftrocks.com/using-ai-without-losing-critical-thinking)
- Score: 2
- Comments: 1
- Posted: 2026-08-09T14:58:29Z

## Translation

タイトル: 批判的思考を失わずに AI を使用する
説明: ありません

記事本文:
ブログ
について
話す
プロジェクト
本の記録
ゲーム記録
ブログ
について
話す
プロジェクト
本の記録
ゲーム記録
批判的思考を失わずに AI を使用する
批判的思考を失わずに AI を使用する
業界では、AI の言うことすべてを盲目的に信じる人を表すために「AI 精神病」という用語が使用されており、AI が真実であると主張したという理由だけで、その人が現実にまったく根拠のない概念を強く信じてしまう状況につながります。ソフトウェア エンジニアリングの一例としては、AI が生成した非常に長い (そして非常に自信に満ちた) 問題の説明と、PR の提案がその問題に対する正しい解決策である理由を記載したプル リクエストを開いた人が挙げられます。しかし、一度深く観察するだけで、その説明の中に実際には何も意味がないことがわかるようになります。 PR は実際には何も解決せず、せいぜい問題を別の場所に移すだけです。 PR の主張はすべてでっち上げであることが判明し、一見すると説得力があるように見える単なる大げさなサラダになっていますが、時間をかけて分析してみると完全にナンセンスです。
これは非常に現実的な問題であり、私は個人的にそれが毎日起こっているのを見てきました。しかし、精神病は病状であり、これについて医学的な意味は何もないため、「精神病」という言い方は非常に悪い表現だと思います。それは単に、誰かが AI とは何か、AI を動かすテクノロジーがどのように機能するかを最初に理解せずに AI を使用しようとした結果にすぎません。
したがって、個人的には、認知的降伏（または最近作られた、よりカジュアルな等価肉代理）という用語のほうが、問題を完璧に捉えているので、この問題を説明するのにはるかに優れていると思います。これは、人々が「おかしくなっている」ということではなく、むしろ、多くの場合、単純に、AI の応答に対する批判的思考の適用に失敗していることです。

AI は間違いを犯す可能性があります (そして常に間違いを犯すでしょう)。
私は過去に、ソフトウェアエンジニアリングには正しいも間違いもないという記事を書きましたが、これは明らかな例外だと私は考えています。 AI の肉の代理人になることが誰かにとって良いアイデアとなり得る理由は 1 つも思いつきません。人が AI の発言を単に代理するだけでは、テーブルに何の価値ももたらしていないからです。今では誰でも AI を利用できるようになりました。そのため、そこから得られるものに対して批判的思考を適用できる能力こそが、あなたに価値をもたらし、他のエンジニアとの違いを生むのです。誰かがそれをせずに、AIの言うことをすべて無視することを決めた場合、その人と対話する人は、AIに直接プロンプトを表示して、その人を完全にスキップする方が簡単ではないでしょうか?
したがって、AI を使用する場合は、テクノロジーの欠点を認識し、それぞれを管理する戦略を立てることが重要です。
AIを正しく効果的に活用する方法
このような状況を踏まえて、この問題を回避し AI を正しく使用する方法に関する私の個人的なガイドラインとアドバイスを以下に示します。
パラダイムは AI に依存しない開発ではなく、AI 支援開発であることを理解する
研究室が最新モデルの能力をどれほど主張しても、どの AI モデルも定期的に結論を急ぎ、指示を忘れ、情報を幻覚するという同じ問題に悩まされています。新しいモデルではこれは解決されません。これはテクノロジとしての LLM の欠陥であり、基礎となるテクノロジ自体が変わらない限り存在し続けるでしょう。研究室がこれらの問題のいずれかを解決したと主張しているのを見ると、それは純粋なマーケティングです。 OpenAI が GPT-2 は危険すぎてリリースできないと言ったのを覚えていますか?同じモデルは、プロンプト以上の列に並ぶのがやっとであることが判明しましたが、それはすべて単なるマーケティングでした。
このメア

AI が多くの調査を行うことが予想される複雑なプロンプトの場合は、常に AI の動作を監視し、正しい方向に進んでいることを確認する必要があります。おかしな仮定をし始めたら、それを中断して正しい方向に戻す必要があります。 AI を監視なしで単独で動作させることは、単純なリクエストにはうまく機能する可能性がありますが、私の経験では、これが複雑な作業にはほとんど機能しません。どれだけ多くの詳細を提供しても、ある時点で間違った仮定を立てて、間違ったものを構築し始めることはほぼ確実です。
LLM がこれらの問題に悩まされる理由をより深く理解したい場合は、LLM を自分で構築するより良い方法はないと思います。 Andrej Karpathy (おそらくこのテーマに関しては世界一の専門家) は、LLM の仕組みとその構築方法についての教育ビデオとチュートリアルをいくつか公開しています。それらはすべて信じられないほどよく書かれており、理解しやすいものです。
AIの主張には常に異議を唱える
また、AI が生成するすべてのものに挑戦し、再確認する習慣も必要です。もともと私は「どうやってこれを知っていますか？」などと尋ねることによってこれを行いました。あるいは「証拠を見せて」と言われますが、最近では、AI が何かを見つけたと主張するたびに、停止フックを使用して自動的にそれを行っています。これにより、AI は結論を急がず、実際に調査を行って証拠を提供するようになります。ここで私たちが取り組んでいる問題の 1 つは、AI が指示を無視する傾向があるため、システム プロンプトの使用は機能しないことに注意してください。この場合、システム プロンプトだけでは信頼できません。
AI が主張の証拠を示したら、それを盲目的に信じたいという衝動に耐えてください。たとえばリンクの場合は、実際にリンクを開いて、そこに情報があること (そして、それが AI によって生成されたものではないこと) を確認します。これはベックです

だって、その「証明」もAIが完全にでっち上げたものである場合も多いでしょう。したがって、この困難なプロセスは、証拠に議論の余地がなくなるまで続けられます。
(コード変更の場合) 生成されたコードを 1 行ずつ読み取ります
この問題は、新しいモデルでは改善されると思いますが、コーディングの知識を失わないようにするためには、常に行うことが良いと思います。
前の 2 つのガイドラインに従えば、最終的には機能するコードが作成されます。ただし、それはコードが最善の方法で構造化されていることを意味するものではありません。少なくとも今日のモデルでは、結果はおそらく非常に複雑になり、重複に満ち、関心事の分離がゼロになり、エッジケースの誤った処理など、基本的に初心者が犯しそうな間違いばかりになるでしょう。
次に、生成されたコードを確認することをお勧めします。次の 2 つの理由があります。1) コードが正しく構造化されており、必要なエッジ ケースを処理していることを確認するため、2) コードの何が変更されたかを実際に理解していることを確認するためです。この 2 番目の理由が特に重要なのは、多くの人が AI の使用の結果、プログラミング スキルを失いつつあると感じているためです。そのため、新しいモデルによってコードの品質の問題が解消されたとしても、少なくともスキルや知識の喪失を防ぎ、コードがどのように機能するかについての理解を維持するには、コードを見直すことが常に役立つのではないかと思います。
© 2026 ブルーノ・ロシャ
ホーム / すべての投稿を見る

## Original Extract

It doesn

blog
about
talks
projects
book recs
game recs
blog
about
talks
projects
book recs
game recs
Using AI without losing your critical thinking
Using AI without losing your critical thinking
The industry has been using the term AI Psychosis to describe a person who blindly believes everything that AI says, leading to situations where said person develops strong beliefs in concepts that have zero grounding in reality whatsoever, simply because the AI claimed it was true. One example from software engineering would be someone opening a pull request with a very long (and very confident!) AI-generated description of a problem and why the proposal in the PR is the correct solution for it, but one deep look is all it takes to realize that nothing in that description actually makes sense. The PR doesn't really solve anything, at best it's just moving the problem elsewhere. Every claim made by the PR turns out to be made up, making it just a big word salad that looks convincing at a glance but is complete nonsense if you really take the time to analyze it.
This is a very real problem and personally I've been seeing it happen every day. But I think "psychosis" is a really bad way to describe it because psychosis is a medical condition, and there's nothing medical about this. It's simply the consequence of someone trying to use AI without first understanding what AI "is" and how the tech that powers it works.
So personally I think the term cognitive surrender (or the recently coined more casual equivalent meat proxy ) is a much better way to describe it because it perfectly captures the problem: it's not that people are getting "crazy", but rather that they are failing to apply critical thinking to the AI's responses, often simply due to a lack of understanding that the AI can (and will, constantly ) make mistakes.
I've written posts in the past about how there's no right or wrong in software engineering, but this is something that I think is a clear exception. I cannot think of a single reason why being a meat proxy for the AI could be a good idea for someone, because if a person simply proxies what the AI says, then they are not bringing any value to the table. Anyone can prompt AI nowadays, so it's your ability to apply critical thinking to what comes out of it that gives you value and sets you apart from other engineers. If someone doesn't do that and decides to just bounce whatever the AI tells them, wouldn't it be easier for someone interacting with said person to just prompt the AI directly and skip the person entirely?
It's then critical that when you work AI, you need to be aware of the technology's shortcomings and have strategies to manage each of them.
How to work with AI correctly and effectively
Given this context, here are my personal guidelines and advice on how to avoid this problem and use AI correctly:
Understand the paradigm is AI-assisted development, not AI-independent development
It doesn't matter how capable a lab claims a latest model is, every single AI model suffers from the same problems of regularly jumping to conclusions, forgetting instructions, and hallucinating information. Newer models will not solve this , this is a flaw of LLMs as a technology and will continue to exist for as long as the underlying tech itself doesn't change. Any time you see a lab claim that they solved any of these problems is pure marketing. Remember when OpenAI said that GPT-2 was too dangerous to be released ? The same model turned out to barely be able to stay in line for more than a prompt, it was all just marketing.
This means that for complex prompts where you expect the AI to do a lot of research, you should always watch the AI as it does its thing and make it sure it's going the right way. If it starts to make weird assumptions, you should interrupt it and guide it back the right way. I find that letting the AI work on its own with zero oversight can work well for simple requests, but in my experience this rarely works for complex pieces of work. No matter how many details you provide it's pretty much guaranteed that it will make wrong assumptions at some point and start building the wrong thing.
If you want to develop a deeper understanding of why LLMs suffer from these problems, I find that there's nothing better than building one yourself. Andrej Karpathy (arguably the number 1 expert on the topic in the world) has several educational videos and tutorials on how LLMs work and how to build them, and they are all incredibly well written and easy to follow.
Always challenge the AI's claims
You should also have a habit of challenging and double-checking everything the AI produces. I originally did this by asking things such as "how do you know this?" or "give me proof", but nowadays I'm using a stop hook to do that automatically every time the AI claims to have found something. This nudges the AI to not jump to conclusions and actually perform research and provide proof. Note that using a system prompt will not work because one of the issues we're fighting here is that the AI tends to sometimes ignore instructions, so system prompts alone are not reliable in this case.
When the AI provides proof for a claim, endure the urge to blindly believe it. If it's a link for example, actually open the link and confirm that the information is there (and that it's not also AI-generated). This is because there will be many cases where the "proof" is also something the AI completely made up. This challenging process thus goes on until the proof is indisputable.
(For code changes) Read the generated code line by line
This issue is something that I do think newer models will improve, but I think it will always be good to do regardless to prevent you from losing your coding knowledge.
If you followed the previous two guidelines, you will end up with code that works. However, that doesn't mean the code is structured the best way it could. At least with today's models, the result will probably be massively overcomplicated, full of duplication, zero separation of concerns, wrong handling of edge cases, basically all the mistakes you'd do as a beginner.
I then think it's a good idea to go over the generated code for two reasons: 1) to make sure the code is structured correctly and handles the neccesary edge cases, and 2) to make sure you actually understand what has changed in the code. This second reason is particularly important because many have been feeling that they are losing their programming skills as a result of their AI usage, so I think that even though newer models might eliminate the code quality problem, going over the code may always be useful at least to prevent you from losing your skills and knowledge and retain your understanding of how the code works.
© 2026 Bruno Rocha
Home / See all posts
