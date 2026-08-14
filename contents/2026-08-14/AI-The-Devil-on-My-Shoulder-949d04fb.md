---
source: "https://myrkvi.no/ai-the-devil-on-my-shoulder/"
hn_url: "https://news.ycombinator.com/item?id=49300591"
title: "AI: The Devil on My Shoulder"
article_title: "AI: The Devil on My Shoulder – Myrkvi"
author: "speckx"
captured_at: "2026-08-14T16:41:03Z"
capture_tool: "hn-digest"
hn_id: 49300591
score: 2
comments: 0
posted_at: "2026-08-14T16:00:54Z"
tags:
  - hacker-news
  - translated
---

# AI: The Devil on My Shoulder

- HN: [49300591](https://news.ycombinator.com/item?id=49300591)
- Source: [myrkvi.no](https://myrkvi.no/ai-the-devil-on-my-shoulder/)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T16:00:54Z

## Translation

タイトル: AI: 肩の上の悪魔
記事のタイトル: AI: 肩の上の悪魔 – ミルクヴィ
説明: これは、私の見解と AI の使用の両方について、一種の回顧を目的としています。しばらくの間、それについての私の考えと経験を書き留めておきたいと思いました、最も...

記事本文:
これは、私の見解と AI の使用の両方について、一種の回顧を目的としています。私はしばらくの間、主に自分自身のために、それに関する私の考えと経験を書き留めたいと思っていました。また、学生以来、本格的に（厳密にはドキュメントではありませんが）たくさん書いたのも初めてです。これは非常に思考回路の話であり、あまり洗練されていないので、自己責任で読んでください。
私も懐疑論者であると考えるほとんどの開発者と同じように、あるいはもう少し熱心に AI の発展を追ってきました。初めのうちは、いじってみるのは楽しかったのですが、よほど簡単な仕事でない限り、どんな仕事を与えられても失敗せずに競争することが（当時は）不可能であることは明らかでした。
私は主に、より日常的なことや定型文に行または複数行の補完を使用しました。本当にできることはそれくらいだと思いました。これは、予測パターンを持つ、反復的でありふれたものにとって素晴らしい後押しとなりました。
Junie が発売されたときに試してみましたが、うまくいきました…ある程度は大丈夫でした。繰り返しになりますが、物事がうまくいかなかったので、問題が発生することがありましたが、それは改善でした。まだ技術デモのような感じでした。確かに、これは人々が真剣に使用できるものではありませんでした。その後、モデルは改良を続けました。仕事を通じてクロードを手に入れた。いろいろ遊んでみることができました。特にしばらくすると、状況は改善し続けました。おそらく、より大きなタスクを与えて完了できる段階に達しました。
私はそれに課題を与えました。私の Discord ボット趣味プロジェクトに Web インターフェイスを提供するようにしました。ずっとやりたいと思っていたのですが、時間がかかりすぎると感じていました。それをAIに任せればいいのです。最初の反復ではフロントエンドとして Svelte を使用しましたが、Go またはその一般的なものを使用できたとしても、(開発) 依存関係として「フロントエンドのすべて」を追加するのは多すぎることがわかりました。

tes Goのコード。最終的に実行するまでに 1 か月ほどかかりましたが、それでも、それをメイン ブランチにマージし直すまでには少し時間がかかりました。私はそれが何を管理したかに本当に感銘を受けました。すべての設定を Web インターフェイスから管理できるようになりました。さらに、Web インターフェイスでのみ利用可能ないくつかの新しい機能が追加されました。そして最後に、サーバー用の監査ログ ビューアです。
まだ何かが少し違和感を感じていた、あるいは少なくとも何かが不安だった。変更点とコードをざっと読んだことはありましたが、手書きで書いた場合のように 1 行ずつ実際に読むことはありませんでした。代わりに、コードベースに 7,000 行の新しい行が追加され、その後、私自身が書いていない新しい行がさらに 5,000 行追加されました。なんだか汚くて、そこにあってはいけないものがあるような気がした。白いシャツについた食べ物のシミのようなもの。同時に、これらの機能を短期間で実現できたことに信じられないほどの感動を覚えました。しかし、何か違和感を感じたという事実を払拭することはできませんでした。
プロジェクトをコントロールできなくなったように感じました。そして率直に言って、私はそう思っていました。私はかなりのタスクを AI エージェントに委任しました。以前持っていたコントロールの一部を失いました。それまではほとんどソロプロジェクトだったので。しかし、他の人がプロジェクトに貢献していたときは、このようには感じませんでした。それは、範囲が違ったからでしょうか、推論ができ、その仕事に対して責任を負える人が私にいたからでしょうか？たぶん両方の部分が少しあるかもしれない。おそらくこれは、私のプロジェクト リーダーが私たちにプロジェクトを引き渡すとき、または少なくとも最初に開発者からプロジェクト リーダーに役割を変更したときの気持ちです。
おそらくそれは、これが私にとって非常に個人的なプロジェクトであり、これほど重要な仕事が他人の手に委ねられるという事実、つまり、何であれそこに魂を込めることのできない機械の手に委ねられるという事実なのかもしれない。

つまり。おそらくそれは、LLM や AI エージェントの使用が膨大な量の電力と (飲料!) 水を消費しているという事実、または彼らがこれらのモデルを他の人々の労働に基づいてトレーニングしたという事実であり、彼らの仕事は 1 セントもこのために利用されることは決してありません。
なぜ私がこのように感じるのかについては、単一の答えはないかもしれませんが、それは問題ありません。上記のすべてをうまく組み合わせたものだと思います。扱うのが汚い気がするし、もし将来的にこれを細かく管理してより綿密にレビューするとしたら（そうするべきですし、そうするつもりです）、気分が良くなるとは思えません。コードを下位レベルで見ると、おそらくさらに多くの不一致や問題が見つかり、おそらく繰り返し修正しなければならないのがイライラするでしょう。しかしそれでも、数週間にわたるプロジェクトを週末のプロジェクトに圧縮できるということには、うんざりするほど満足感があります。私はそれが良いことだとはあまり思いません。私たちが仕事をするスピードで仕事をするのには理由があります。私たちが行うすべてのことを書き、処理し、考えるには時間がかかりますが、それを加速することはできません。
そうは言っても、AI には別の懸念もあります。電気や水の使用量に影響を及ぼし、データセンターは住宅地の近くに建設され、多くの騒音公害を引き起こしています。まさに同じ製品が、誤った情報を広め、さらに際限のないコンテンツで人々を疲弊させるために使用されています。おそらくこれらの懸念が、もし必要になった場合に、私がそれを使い続けるべきではない本当の理由だと思います。
それにもかかわらず、私たちは職場でAIエージェントをますます活用するよう求められています。さまざまな理由から、まだそうなっていませんが、これはある種の祝福のように感じています。なぜなら、私の勤務時間がクロードまたは ChatGPT に何をすべきかを指示し、彼らの「アンダー」を確認することで構成されている場合、これについてどう感じるか本当にわかりません。

タスク、その仕様書、およびコードの「立ち読み」。
結局のところ、それは本当に、ある種の悪魔が私の肩に座って、降参して、できればもっともっと使い続けるようにと言っているように感じます。しかし、そうなると自分の一部を失うような気がします。
[ コメントを表示またはこの投稿へのコメント ]
#あい
#chatgpt
#クロード
#コーデックス
#llm
#考え

## Original Extract

This is meant to be a sort of retrospective, both in my view and use of AI. I have wanted to write down my thoughts and experience with it for a while, mostl...

This is meant to be a sort of retrospective, both in my view and use of AI. I have wanted to write down my thoughts and experience with it for a while, mostly for my own sake. It's also my first time really writing much (that isn't strictly documentation) since school. It is very much a train of thought thing and not very polished, read at your own peril.
I have followed the development of AI like most developers who consider themselves skeptics, maybe a bit more enthusiastically so. In the beginning, it was fun to play around with, but it was obvious that it (at the time) could not be handed any task and compete it without failing, unless it was a very simple one.
I mostly used the line or multi-line completion for more mundane things and boilerplate. I figured that was really the extent of what it could do. It was a nice boost for repetitive, mundane stuff that had a predictive pattern.
I tried Junie when it came out, and it worked... somewhat okay. Again, it struggled with me come things and would therefore get things wrong, but it was an improvement. It still just felt like a tech demo. Surely this wasn't something people could use seriously. Then, models continued to improve. I got Claude through work. I could play around with it. Things continued to improve, especially after a while. It reached a point where I could probably give it a larger task that it could complete.
I gave it a task. I made it give my Discord bot hobby project a web interface. It was something I had wanted to do for a long time, but that felt too time-consuming. I could just hand it off to the AI. The first iteration used Svelte as the front-end, and I found it to be too much to add "all of front-end" as a (development) dependency, if I could just use Go, or something that generates Go code. A month or so went by before I finally did it, and even then, it took some time before I dared merging it back into the main branch. I was genuinely impressed by what it had managed. All the settings could now be managed from the web interface. Additionally, some new things only available in the web interface were added. And finally, an audit log viewer for a server.
Something still felt kinda off, or at least something was unsettling. I had skimmed over the changes and the code, but I never truly read it line by line, like I would have had to if I wrote it by hand. Instead, 7,000 new lines were added to the codebase, then another 5,000, that I had not written myself. It felt dirty somehow, like it was something that wasn't supposed to be there. Like a food stain on a white shirt. At the same time, it felt incredible to have made those features in the short amount of time it took. But, I couldn't shake the fact that something felt off.
It felt like I had lost some control of the project. And frankly, I had; I delegated a sizeable task to an AI agent. I had lost some of that control I held previously. Because it was mostly a solo project prior to this. Yet, I did not feel like this the times other had contributed things to the project. Is it because the scope was different, because I had someone I knew could reason and that could be held accountable for their work? Maybe a bit of both. Maybe this is how my project lead feels whenever he hands off projects for us to work on, or at least when he first changed roles from a developer to a project lead.
Maybe it is the fact that it has been such a personal project to me, and that having such a significant task be put in someone else's hands-- a machine's hands, with no real soul to put into it, whatever that would mean. Maybe it's the fact that the use of LLMs and AI agents is consuming vast amounts of electricity and (drinking!) water, or that they have trained these models on other people's labour, who'll never see a cent of their work being utilised for this.
There might not be a single answer to why this feels the way it does to me, and that's fine. I think it's a good mix of all of the above. It feels dirty to work with, and if I were to micromanage and review it more closely in the future (which I should, and will), I don't think I will feel better about it. I will probably see more of the inconsistencies and issues when I see the code at a lower level, which will be frustrating to have to correct, likely repeatedly. But even then, there is something disgustingly satisfying about being able to compress a multi-week project into a weekend thing. I don't really think that it is a good thing; we do the tasks at the speed we do for a reason. It takes time to write, process, and think about all that we do, and we can't just accelerate that.
With all that said, there are also the other concerns of AI. The impact it has on electricity and water usage, data centres being built near residental areas, creating a lot of noise pollution. The very same products being used to spread misinformation and exhaust people with even more endless content. I think that those concerns are probably the real reson why I shouldn't be continuing to use it, if I ever needed some.
Nevertheless, we are being pushed at work to use AI agents more and more. Due to various reasons, we have not yet, which to some extent has felt like some kind of blessing, because I really do not know how I would feel about this when my workdays will consist of telling Claude or ChatGPT what to do, and reviewing their "understanding" of tasks, their spec documents, and their code.
In the end, it really does feel like some sort of devil sitting on my shoulder, telling me to give in and just (continue to) use it, preferably more and more. But I do feel like I will lose a part of myself if I do.
[ View comments or comment on this post ]
#ai
#chatgpt
#claude
#codex
#llm
#thougts
