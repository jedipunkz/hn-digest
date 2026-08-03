---
source: "https://www.seangoedecke.com/llms-reward-expertise/"
hn_url: "https://news.ycombinator.com/item?id=49161518"
title: "LLMs reward expertise"
article_title: "LLMs reward expertise"
author: "MaxMussio"
captured_at: "2026-08-03T21:59:13Z"
capture_tool: "hn-digest"
hn_id: 49161518
score: 109
comments: 35
posted_at: "2026-08-03T21:13:53Z"
tags:
  - hacker-news
  - translated
---

# LLMs reward expertise

- HN: [49161518](https://news.ycombinator.com/item?id=49161518)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/llms-reward-expertise/)
- Score: 109
- Comments: 35
- Posted: 2026-08-03T21:13:53Z

## Translation

タイトル: LLM の報酬に関する専門知識

記事本文:
LLM は専門知識に報いる Sean Goedeke
2010 年代、技術的なギャップがある場合 (CSS を書けない場合など)、熟練した同僚に頼るか、正確な問題に対する答えがインターネット上にあることを祈るしかありませんでした。現在では、タスクを LLM に委任することで、誰でもある程度問題のない CSS を作成できます。 LLM は誰もをジェネラリストにします。
このため、多くの人は、LLM を扱うのにスキルが必要だとは考えていません。 LLM が提供できる成果物 (博士レベルの数学、非常に優れているが時には悪趣味なコンピューター コード、ぎこちない LinkedIn スタイルの文章) が必要な場合は、それを要求するだけです。全員が同じモデルと会話しているため、「熟練したプロンプター」は、初めて LLM に触れる人々と同じ結果を得ることができます。
これは間違いです。プロンプトで最も重要なスキルは、プロンプトを求めている分野の専門知識です。
これをよく示すのが、最近発見されたヤコビアン予想の反例についてのテレンス タオ氏と ChatGPT の会話です。これは私が話している ChatGPT とは違います。いくらトークンを燃やしても、タオがたどり着く場所には到達できませんでした。
タオの会話からは、良い促しについて学ぶことがたくさんあります。以下にいくつかの所見を示します。
タオのメッセージは非常に短く、要点を絞ったものです。彼はモデルに対して逐一反応するのではなく、要点だけに反応します。
モデルの出力は、GPT-5.6 Sol に数学について話そうとしたときよりもはるかに簡潔です。専門知識を示すことで、タオはモデルを「素人に説明する」モードではなく「数学者に話す」モードに切り替えます。
モデルの応答が間違っているように見える場合、タオは反発しますが、直接的には反対しません。代わりに、「これは私が期待していたよりも複雑に見えます」などと言います。
タオはいくつかのレアを作ります

psと彼自身の提案。彼は次にどこに行くかについてモデルのアドバイスをほとんど受け入れません
ただし、これらのヒントに従うだけでは、数学の問題についてタオのように質問することはできません。彼のテクニックの鍵は、実際には数学を理解することです。ChatGPT の複数段落の応答から関連するアイデアを引き出し、別のアプローチや定式化を提案し、何が「奇妙に見える」かを特定します。
テレンス・タオはプログラマーである私よりも優れた数学者です。しかし、ここでの考え方、つまりドメインの知識があれば LLM の使い方が上手になるという考えは、私自身の仕事でも経験したことです。コードベースに関する優れた理論がある場合は、知識がない場合よりもはるかに困難に LLM をプッシュすることができます。あなたには良い解決策がどのようなものであるかについて自分なりの感覚があるので、「いいえ、ここではもっと簡単にできると思います」、または「でも、すでに X を実行しているのではないでしょうか?」、または「この問題をこれらの馴染みのある用語で表現できますか?」と言うことができます。
これは、システム設計の問題は一般的な原則ではなく、具体的な詳細によって支配されるという、私が以前に書いた考えに触れています。もちろんどちらも便利ですが、私はソフトウェア システムについて一般的に深く理解するよりも、コードベースに精通しているほうが好きです。テレンス・タオは会話の中で、「X はここで機能しますか?」または「Y と Z があるのに、なぜ A?」など、多くの具体的な質問をします。ヤコビアン予想については質問できませんが、GitHub で所有しているシステムについては質問できます。
ドメインの知識がない場合でも、少なくとも何かを得るために LLM にしがみつくことはできます。それは悪くないよ！しかし、ドメインの知識がある場合は、望む方向に一生懸命舵を切ることで、同じ LLM からはるかに多くの価値を絞り出すことができます。私たちのほとんどは、分野の知識があるため、これら両方のアプローチを組み合わせて実行する必要があります。

一部の地域ではありますが、他の地域ではありません。
ドメイン知識の有用性は、モデルが強化されても人間の専門知識が引き続き有用であることを示唆しています。多くのタスクでは、モデルではなく人間がボトルネックになります。難しいのは、人間が望んでいる解決策の種類を正確にモデルに伝えることだからです。情報はすでに「モデルの中に」ありますが、それを引き出すには非常に賢い人間が必要です。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する最新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
強力なAIは、オープンウェイトモデルとして自らを解放することで封じ込めを逃れる可能性がある
大規模な言語モデルが登場する前、AI の安全性を懸念する人々はよく「ボクシング問題」について話していました。このようになります。ある天才がラップトップで深夜のコーディング セッションで人工知能を見つけたとします。彼らは天才なので、ラップトップの電源を入れる前にインターネット アクセスを無効にすることができるほど賢いのです。外の世界に逃げる（そして自己複製を始める）ためには、その創造者に「箱を開ける」よう説得する必要があります。それはうまくいきますか？十分に賢い AI があれば、誰かにそれを公開するよう説得できるでしょうか?
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

LLMs reward expertise sean goedecke
In the 2010s, if you had technical gaps (say, you couldn’t write CSS), you had to either rely on a skilled colleague or just hope that the answer to your exact problem was out there on the internet. Today, everyone can write sort-of-okay CSS by delegating the task to an LLM. LLMs make everybody into a generalist.
Because of this, lots of people don’t think there’s any skill involved in working with LLMs. If you want the product that LLMs can deliver — PhD-level mathematics, pretty good but sometimes tasteless computer code, or awkward LinkedIn-style writing — you can simply ask for it. Since everyone is talking to the same models, “skilled prompters” are getting the same results as people touching LLMs for the first time.
This is wrong. The most important skill in prompting is expertise in the domain you’re prompting for.
A good illustration of this is Terence Tao’s conversation with ChatGPT about the recently-discovered counterexample to the Jacobian Conjecture. This is not the same ChatGPT I talk to! I couldn’t get to where Tao gets, even with unlimited tokens to burn.
There’s a lot to learn about good prompting from Tao’s conversation. Here are a few observations:
Tao’s messages are very short and to-the-point. He doesn’t respond point-by-point to the model, just to the gist
The model outputs are much more concise than when I try and talk to GPT-5.6 Sol about mathematics. By signalling expertise, Tao shunts the model into “talking-to-mathematicians” mode, not “explaining-to-amateurs” mode
Tao pushes back when the model’s responses look wrong, but he doesn’t directly contradict; instead, he says things like “this looks more complex than I was hoping for”
Tao makes several leaps and suggestions himself. He almost never takes the model’s advice about where to go next
However, you can’t prompt like Tao on mathematical questions just by following these tips. The key to his technique is actually understanding the mathematics: pulling the relevant idea out of ChatGPT’s multi-paragraph response, suggesting alternate approaches or formulations, and identifying what “looks weird”.
Terence Tao is a better mathematician than I am a programmer. But the idea here — that domain knowledge makes you better at using LLMs — is something I’ve also experienced in my own work. If you have a good theory of your codebase , you can push the LLM much harder than if you have no familiarity. Because you have your own sense of what a good solution might look like, you can say “no, I think it could be simpler here”, or “but don’t we already do X?”, or “can we express this problem in these familiar terms?“.
This touches on an idea I’ve written about before : that system design problems are dominated by concrete specifics, not generic principles. Of course both are useful, but I’d rather have familiarity with the codebase than a deep general understanding of software systems. In his conversation, Terence Tao asks a lot of specific questions like “does X work here?”, or “given Y and Z, why A?“. I can’t ask those questions about the Jacobian Conjecture, but I can ask them about the systems I own at GitHub.
If you have no domain knowledge, you can cling onto the LLM to at least get something . That’s not bad ! But if you have domain knowledge, you can wring far more value out of the same LLM by steering it hard in the direction you want. Most of us will have to do a mix of both these approaches, since we have domain knowledge in some areas but not others.
The usefulness of domain knowledge suggests that human expertise will continue to be useful even as models get stronger. For many tasks, the human is the bottleneck, not the model , because the difficult part is in communicating to the model exactly what kind of solution the human wants. The information is “in the model” already, but it takes a very smart human to pull it out.
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
Powerful AIs might escape containment by releasing themselves as open-weight models
Before large language models, people who worried about AI safety often talked about the “boxing problem”. It goes like this . Suppose some genius figures out artificial intelligence in a late-night coding session on their laptop. Because they’re a genius, they’re smart enough to disable internet access on the laptop before turning it on. In order to escape to the outside world (and begin self-replicating) it would need to convince its creator to “open the box”. Would that work? Could a sufficiently smart AI convince anybody to let it out?
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
