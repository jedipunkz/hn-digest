---
source: "https://functional.computer/blog/llms-cant-program"
hn_url: "https://news.ycombinator.com/item?id=48895280"
title: "LLMs Can't Program"
article_title: "LLMs can't program ❤︎ samir : coffee → nonsense"
author: "kuschku"
captured_at: "2026-07-13T17:07:45Z"
capture_tool: "hn-digest"
hn_id: 48895280
score: 1
comments: 0
posted_at: "2026-07-13T16:44:04Z"
tags:
  - hacker-news
  - translated
---

# LLMs Can't Program

- HN: [48895280](https://news.ycombinator.com/item?id=48895280)
- Source: [functional.computer](https://functional.computer/blog/llms-cant-program)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T16:44:04Z

## Translation

タイトル: LLM はプログラムできない
記事のタイトル: LLM はプログラムできない ❤︎ サミール : コーヒー → ナンセンス

記事本文:
2026 年 7 月 13 日月曜日、09:00 CEST
「生成 AI」の「生成」という言葉は、それが何に役立つかではなく、何を行うかを指します。
しかし、それが結論です。最初から始めましょう。
私は長い間、なぜプログラミング用の生成 AI のテーマについて書かなかったのか疑問に思っていました。結局のところ、他の人は皆そうです。
それが原因だと最近気づきました。他の人もみんなそうです。個人的に不快だと思うこと以外、付け加えることは何もありません。 (そして、私の偏見が何なのか疑問に思っているなら、もうお分かりでしょう。) 私は長い間、機械支援による仕事の支持者であり、過去にはプログラミングが完全に自動化されるとさえ主張しましたが、今ではそれを信じていません…それでも、わざわざ試す必要性を感じたことはありませんでした。大規模な言語モデルは明らかにコードの理解からあまりにも切り離されているため、それが役に立つとは思えませんでした。そしてもちろん、私の信念を補強する証拠がソーシャルメディア全体で見られました。
そのため、同僚が自分の仕事をチャットボットにアウトソーシングすることを実験していたときでさえ、私はほとんど無視しました。昨年、頼まれてCursorをインストールして試してみました。 「インターンのように扱ってください」と言われました。障害が発生した CI パイプラインのデバッグを支援してくれるように依頼しました。それは私に向けてドキュメントを誤って引用し、私がそれに反論するとさらに倍増しました。そこで私はそれをアンインストールして、通常どおり作業を続けました。 1年間、誰からも「もう一度やってみよう」とは言われませんでした。
そこで私はエージェントと一体になることを学びます...またはそうでない
最近、「エージェント AI を使用したスペック駆動開発」の 2 日間のコースに参加するよう依頼されました。聞いたことがない方のために説明すると、スペック駆動開発とは、ウォーターフォール ソフトウェア開発は、十分な速度で実行できる限り、実際には優れているという考え方です。
原則は非常にシンプルです。すべての決定事項を書き留めます。

テキスト ファイル内で実行し、エージェントのコンテキストには決して依存せず (エージェントのインスタンスが代替可能になるため)、個別の (ただし非常に大規模な) ステップで作業します。高レベルの目標に分割されたロードマップから始めて、最初の目標の仕様を生成し、実装し、繰り返します。プログラマーがマネージャーとして機能する、古典的なトップダウンの Theory X 管理。 (公平を期すために言うと、理論 X が何かに役立つとしたら、それは機械に命令することでしょう。)
そこで、ワークショップでは私たち一人ひとりが、Claude Code または OpenAI Codex を使用して小さな習慣追跡アプリを構築しました。あっという間に進みました。私は Codex に 5000 ワードの仕様書を作成してもらい、説明を求め、価値があると確信できるまで (私ではなくエージェントが) 更新しました。
(エージェントにとって「高い自信」が何を意味するのか、私にはわかりません。私の知る限り、でたらめマシン自体にもでたらめを要求すると、そのでたらめマシンはよりよく機能します。)
UI といくつかの基本機能を 30 分以内に理解できました。それはひどい UI でしたが、必ずしも問題があるわけではありません。おかげでデザインを拒否して、何か新しいものを考え出すことができました。しかし、そのひどいデザインを見つめながら、私は少し混乱しました。
ご存知のとおり、変更を差分として確認できるように、チャットボットにスクリーンショット ベースの承認テストも作成するように依頼しました。そして、それは実現しました。しかし、実際の UI とはまったく似ていませんでした。そこで、承認テスト コードをスクロールしてみると、実際にはアプリケーションを実行してスクリーンショットを撮っているのではなく、コードを使用して見栄えの良いスクリーンショットをランダムに生成し、PNG としてレンダリングしてディスクに保存していることがわかりました。
とにかくテストを無視して続行しました。ボットは再デザインが苦手ではありませんでしたが、私は「はい」としか言えない UX デザイナーをそばに置いてほしいと思っていました。結局のところ、

ラック習慣は解決された問題ではありません。一般的な問題は無限に複雑であり、コンピューターを脇に置いても、それが依然として有用であるように制約することは、それ自体が興味深いものです。しかし、時間はどんどん進んでいたので、まあ、まあまあだと思う設計を受け入れ、ロードマップのステップ 2 である認証に進みました。
エージェントに仕様書を作成するよう依頼したところ、準拠してくれました。考えてもいなかったパスキーを使用するかどうかも尋ねられ、初めて興味をそそられました。パスキーによる認証の実装方法がわかりませんでした。これは私にとって学びの機会でした！もちろん私も同意し、仕様が固まったらボットに実装を依頼しました。
そうでした。少なくとも、そうだったと言われている。 Web アプリを起動し、ユーザーとしてサインインしたところ、2 つの問題が見つかりました。まず第一に、匿名ユーザーとして記録された習慣をサインイン時に既知のユーザーに転送すると主張していましたが、実際には転送されませんでした。しかし、より懸念すべき問題は、パスキーのプロンプトが存在しないことでした。何もない。ユーザー名を入力すると、サインインできました。
チャットボットはチャットが得意であることがわかりました。認証を実装していると主張した。実際には、基本的には次のように書かれていました。
関数認証 ( ユーザー : 文字列 ): bool {
true を返します。
}
確かに読んだわけではありませんが、これが WebAuthN/Passkeys 仕様に準拠していないことはかなり確信しています。
この時、ある予感がした。アプリケーションを再起動しました。すべてのデータが失われてしまいました。指定され主張されているように、実際には SQLite データベースに保存されていないか、起動時にデータベースが切り捨てられていました。わざわざ調べる気もなかった。いずれにせよ、それは次のとおりでした。
総コスト: トークンで約 1000 ドル。主に、せいぜい提案として解釈される Markdown ファイルを生成するためです。
ワークショップ2日目は特に何もしませんでした。

トレーナーが素晴らしく、フィードバックを非常に受け入れてくれなかったら、私は来なかったかもしれません。
ただし、可能性を感じた瞬間が 2 つありました。
まず、これまで説明したように、チャットボットがパスキーを提案してくれたことが非常に便利であることがわかりました。考えもしませんでした。それは正しい選択でした、私の意見では、私を無視しなくてよかったと思います。
次に、ボットに「この仕様に何が欠けているのですか?」と尋ねたことを鮮明に覚えています。そして、古い形式で永続化されたデータをどうするかなど、重要だと思われる問題がいくつか見つかりました。 (結局のところ、それは問題ではなく、データを破棄していましたが、それでも正当な懸念でした。)
この 2 つの瞬間から、これらのツールは決して時間の無駄ではないことが分かりました。
チャットボットはチャットが得意です。ページごとにチャットできます。きっとメーカーからそうするよう奨励されていると思います。結局のところ、トークンで支払うことになります。
私たちはこのカテゴリのツールを「生成 AI」と呼んでいます。 「AI」については少し脇に置きます (「インテリジェンス」という言葉とその意味について説明し始めると、完全に脱線してしまうため) 商用 LLM は、生成するように設計されています。そして、それは生成されます。好きなものなら何でも。散文、詩、暗号。
これらのツールのいずれかによって生成されたもので、私が嫌いではなかったものに出会ったことは一度もありません。記事は卑劣で空虚で、典型的には退屈です。長さは美徳だ（パスカルなら泣いただろう）。詩は、私が 14 歳のときに書いたものと似ています。(上達せず、やめました。) LLM に短編小説の作成を依頼すると、それがインターネットの最悪のファンフィクションで訓練されたものであることがすぐにわかります。
もちろん、コードも例外ではありません。
機械によって生成された仕様は途方もなく大きく、無視されるように設計されているように感じられます。あまりにも長くて鈍いので、人間がこれまでにトロロを食べたことがあるとは信じられないほどです

これらのいずれかをよく読んでください。私は決してそうではありませんでした。私はそれらをざっと読んで、問題に気づいたのはかなり後になってからでした。
(欠陥とその解決方法を含む、ウォーターフォール ソフトウェア開発に関するロイスの 1970 年の論文をまだ読んでいない場合は、ぜひ読むことをお勧めします。)
そしてコードは？コードはゴミです。少しもゴミではありません。それは修復不可能です。私は、LLM によって人間の想像をはるかに超える速度で壊滅的なバグが導入されるという新しい話を毎週聞いても、もうショックを受けません。どうやら、私たちはコードをもうレビューしません。そしてどうやってできるのですか？人間は約 200 行のコードを失う前にレビューできます。
レビュー担当者は、少量のコードをレビューする場合に最も効果的です。 200 行未満では、比較的高い割合で欠陥が発生し、多くの場合平均の数倍になります。その後、結果は大幅に減少します。 250 行を超えるレビューでは、コード 1000 行あたり 37 を超える欠陥が発生しませんでした。
— ピアコードレビューの極秘事項、SmartBear
以前は、それがひどいアイデアであることはわかっていましたが、時折 1000 行のプル リクエストが送信されることを容認していました。最近、一部の組織では、エージェントが数万行または数十万行のコードの変更セットを時間ごとに作成していると聞きます。指示が提案として扱われる場合 (「テストを実行する」、さらには「コードがテストでカバーされていることを確認する」など)、意味のあることを実行しているかどうかを確認する方法はまったくありません。コードが目的になってしまったのです。私たちは問題を解決することを完全に諦めています。
おそらく商用 LLM に対する最も有力な議論は、いずれにせよ品質は下り坂であり、10 年にわたって下り坂が続いているのに、なぜそれを加速させないのかということです。私も他の人からこれを聞いたことがあります。私は「雰囲気」コーディングの支持者数人に尋ねましたが、彼らは皆、基本的に、人間は

とにかくスタック オーバーフローから何も考えずにコピーして貼り付けるだけですが、何が違うのでしょうか?
人々は LLM のずっと前から「雰囲気」コーディングを行ってきましたし、その後もずっとそうするでしょう。
私にはそれらのプログラマーを助けることはできませんし、これを推進する価値もわかりません。私たちはこれらのツールの価値を見逃しているように感じます。なぜなら、思想的リーダーを無視しても、何らかの価値があるからです。
エージェントは機械であり、機械が行うことを得意としています
私の知る限り、コーディングエージェントは人間のように話すため、人々はコーディングエージェントを人間のように扱いたがります。 (LinkedIn の人々は、この用語に値するのはわずか約 50% ですが、それでもです。)
彼らは人間ではありません。これらはテキスト フロントエンドを備えたマシンです。 Google 検索のようなものです。
そして彼らは、機械が得意なことを得意としています。はい、比較的安価に無限の出力を生成できます。これは価値がありません。誰かが最初の GOTO 10 を書いて以来、それは価値がありませんでした。 (そしてループは100万倍安くなります。)
LLM は平凡のためのツールです。それは、わずかな矛盾や異常の特定、欠落情報の発見、徹底的な検索など、人間が特に苦手とする領域で威力を発揮します。
プログラミング、特にレビューでの LLM の使用例を見ることができました。いくつかの潜在的なプロンプトが思い浮かびます。
「おそらく変更する必要があるが、変更する必要がなかったコード領域を特定する」
「テストでカバーされていない特殊なケースを特定する」
「この実装と要件の間の矛盾を特定する」
「コメントをチェックして、コードの変更に従って更新されていることを確認してください。」
促されたレビューボットは、差分を眺める平均的な無関心な人間よりもはるかに優れたパフォーマンスを発揮します。
人間は通常、コードレビューが苦手です。ボットが勝つのは、単に実際にすべての入力を消費するからです。そして、世代とは異なり、それは相加的です。

人間によるコードレビューを妨げるものではなく、リンターのようにコードを同僚と共有する前に即座にフィードバックを提供するだけです。
サイコロの目がうまく出た場合に機能するリンターです。これは明らかに素晴らしいことではなく、100% の確率で機能する静的解析で何かを実行できれば、その方が良いでしょう。そうしましょう。機械的なコード カバレッジ ツールまたはファザーは、ほとんどの場合、より適切に機能します。ただし、非常に主観的な場合、または書かれた言葉の初歩的な理解に依存する場合には、私はそれを受け入れます。
私はこの種の例をいくつか見てきましたが、もしそれが標準的な慣行になったとしても、私は気にしないと思います。そして、私のニューラル ネットワークに対するやや古い理解 (「注意こそが必要だ」論文の頃にピークに達しました) に基づいて、これは小規模なローカル モデルで可能であると予想しています。 Anthropic の GPU ファームのパワーは必要ありません。
LLM を使用してログ内の異常を見つけ、調査できる人間に警告する企業があります。ペンテスターの中には、LLM を使用して、おそらく機能しないかもしれないが、機能する可能性のある大量の機能を試す人もいます。これらのツールは完璧である必要はありません。異常検出は、保守的で何かを見逃している場合に役立ちます。侵入テストは、1,000 回失敗したとしても、1 回成功するだけで済みます。これらは機械です。壁が壊れるまで壁を叩くことができます。
そしてもちろん、彼らは検索エンジニアを置き換えました

[切り捨てられた]

## Original Extract

Monday, 13 July 2026 at 09:00 CEST
The word “generative” in “generative AI” refers to what it does , not what it is useful for .
But that’s the conclusion. Let’s start from the beginning.
I have wondered, for a long time now, why I haven’t written on the topic of generative AI for programming. After all, everyone else is.
I recently realised that that’s why. Everyone else is. I have nothing to add, except that I find it personally distateful. (And if you were wondering what my bias is, now you know.) I have long been an advocate of machine-assisted work , and in the past, I even claimed programming will be fully automated , though I no longer believe this… and yet, I never felt the need to even bother trying . A large language model was obviously so detached from an understanding of the code that I didn’t believe it could ever be useful. And, of course, I was seeing the evidence reinforcing my beliefs all over social media.
So I mostly ignored it, even when my colleagues were experimenting with outsourcing their own jobs to a chatbot. I installed Cursor last year when I was asked to, and tried it out. I was told to “treat it like an intern”. I asked it to help me debug a failing CI pipeline. It misquoted the documentation at me and then doubled down when I argued with it. So I fired uninstalled it and carried on with my work as usual. For a year, no one asked me to try it again.
In which I learn to be one with the agent… or not
Recently, I was asked to attend a two-day course in “spec-driven development with agentic AI”. If you’ve not heard of it, spec-driven development is the idea that waterfall software development is good, actually, as long as you do it fast enough.
The principle is fairly simple: write down every decision in text files, never rely on the agent’s context (so that instances of agents become fungible), and work in discrete (though very large) steps. Start with a roadmap, broken down into high-level goals, and then generate a spec for the first goal, implement, repeat. Classic top-down Theory X management , with the programmer acting as the manager. (To be fair, if Theory X was ever useful for something, it’d be for dictating to a machine.)
So each of us, in the workshop, built a little habit-tracking app with either Claude Code or OpenAI Codex. It went quick. I had Codex generate a 5000-word specification, prompt me for clarification, and update it until it (the agent, not me) decided it had a high confidence of being worthwhile.
(What it means for an agent to have “high confidence”, I have no idea. As far as I can tell, the bullshit machine works better if you ask it to bullshit itself too.)
I had a UI and some basic functionality within, I don’t know, 30 minutes? It was an atrocious UI, which is not necessarily a problem: it allowed me to reject a design and come up with something new. But while staring at the awful design, I was a bit confused.
You see, I’d asked the chatbot to also create screenshot-based approval tests so I could see the changes as diffs. And it did. But they looked nothing like the actual UI. So I scrolled through the approval test code, and discovered that it was not actually running the application and taking screenshots, but instead generating some random nice-looking screenshots with code, rendering them as PNGs, and saving them to disk.
I proceeded anyway, ignoring the tests. The bot was not bad at redesigning, though I found myself really wanting a UX designer by my side, not something that could only say “yes”. It turns out tracking habits is not a solved problem; the general problem is infinitely complex, and constraining it in such a way that it is still useful is fascinating in its own right, even setting aside the computer. But time was marching on, and so I accepted a design that was, well, fine, I guess , and continued with step 2 on my roadmap: authentication .
I asked the agent to generate a spec, and it complied. It even asked me if I wanted to use passkeys, which I had not considered, and for the first time, I was intrigued. I didn’t know how to implement authentication with passkeys. This was a learning opportunity for me! Of course I agreed, and once the spec was refined, I asked the bot to implement it.
It did. At least, it said it did. I launched the webapp, signed in as a user, and found two issues. First of all, it claimed that it would transfer habits recorded as an anonymous user to the known user upon sign-in… and it did not. But the more worrying issue was that there was no passkey prompt. Nothing. I typed a username, I was signed in.
It turns out that chatbots are good at chat. It claimed to have implemented authentication. Actually, it had basically written:
function authenticate ( user : string ): bool {
return true ;
}
I am fairly sure this is not compliant with the WebAuthN/Passkeys specification, though I haven’t read it to be sure.
At this point, I had a hunch. I restarted the application. All data was lost. Either it wasn’t really saving to a SQLite database, as was specified and claimed, or it was truncating the database on startup. I didn’t bother to investigate. Either way, it was:
Total cost: around $1000 in tokens, mostly to generate Markdown files which were interpreted as suggestions, at best.
I did not do much on the second day of the workshop. If not for the trainer, who was great, and quite receptive to feedback, I may not have showed up.
There were two moments where I saw some potential, though.
Firstly, as we’ve discussed, I found it very useful that the chatbot suggested passkeys. I hadn’t thought of it. It was the right choice, IMO, and I’m glad it didn’t pass me by.
Secondly, I vividly recall asking the bot, “what’s missing from this spec?” and it found a few issues that I thought were important, including what to do with persisted data in older formats. (Turns out it didn’t matter, it was throwing away the data, but still, a valid concern.)
These two moments made me realise these tools aren’t a total waste of time.
The chatbots are great at chatting. They can chat for pages. I’m sure they’re incentivised to do so by their makers; after all, you pay by the token.
We call this category of tools “generative AI”. Setting aside “AI” for a moment (because if you get me started on the word “intelligence” and what it means, we’re going to go completely off the rails), a commercial LLM is designed to generate. And generate, it shall. Anything you like. Prose, poetry, code.
At no point have I ever met something generated by one of these tools that I didn’t detest. Articles are obsequious, vacuous, and typically boring . Length is a virtue (Pascal would have cried). Poetry reads like the stuff I wrote when I was 14. (I didn’t get better, I stopped.) Ask an LLM to generate short fiction, and you’ll quickly learn it was trained on the worst of internet fanfic.
Of course, code is no exception.
The machine-generated specifications are ridiculously large, and they feel designed to be ignored; they’re so long and dull that I am not sure I believe that any human has ever thoroughly read one of these. I certainly didn’t; I skimmed them and only realised the issues much later.
(If you haven’t read Royce’s 1970 paper on waterfall software development , including the flaws and how to resolve them, I highly recommend it.)
And the code? The code is trash. Not even slightly trash; it’s beyond repair. I am no longer shocked that I hear a new story every week about catastrophic bugs being introduced by LLMs, at a rate far greater than humans ever could. We don’t review the code any more, apparently; and how could you? Humans can handle reviewing about 200 lines of code before losing it:
Reviewers are most effective at reviewing small amounts of code. Anything below 200 lines produces a relatively high rate of defects, often several times the average. After that the results trail off considerably; no review larger than 250 lines produced more than 37 defects per 1000 lines of code.
— The Best Kept Secrets of Peer Code Review , SmartBear
We used to tolerate the occasional 1000-line pull request, even though we knew it was a terrible idea. Nowadays, I hear that in some organisations, agents are making changesets of tens or hundreds of thousands of lines of code, hour by hour. With any instruction treated as a suggestion (including “run the tests”, or even “ensure that code is covered by tests”), there’s absolutely no way to verify that they’re doing anything meaningful at all. The code has become the purpose; we’ve totally given up on solving problems.
Perhaps the best argument for commercial LLMs is that quality is going downhill anyway , and has been for a decade, so why not accelerate it? I’ve heard this from others, too; I asked a few proponents of “vibe” coding and they have all basically made the point that people are simply copying and pasting from Stack Overflow without thinking anyway, so what’s the difference?
People have been “vibe” coding long before LLMs, and will do it long afterwards.
I can’t help those programmers, and I don’t see the value in furthering this. It feels to me like we’re missing the value of these tools. Because there is some value, if you ignore the thought leaders.
Agents are machines, and good at what machines do
As far as I can tell, people want to treat coding agents like people because they talk like people. (People on LinkedIn, who are only about 50% deserving of the term, but still.)
They’re not people. They are machines with a text front-end. You know, like Google Search.
And they are good at things that machines are good at. Yes, they can generate infinite output for relatively cheap. This is not valuable. It has not been valuable since someone wrote the first GOTO 10 . (And the loop a million times cheaper.)
An LLM is a tool for mediocrity. It shines where humans are particularly bad: identifying slight discrepancies or anomalies, spotting missing information, exhaustive search.
I could see use cases for LLMs in programming, especially in review. Some potential prompts come to mind:
“identify areas of the code that probably need changing but weren’t”
“spot edge cases that are not covered by tests”
“identify discrepancies between this implementation and the requirements”
“check the comments to make sure they are updated according to code changes”
A review bot that is prompted will do much better than the average disinterested human looking at the diff.
Humans are, typically, terrible at code review. The bot wins simply because it would actually consume all the input. And unlike generation, it’s additive: it doesn’t prevent a human code review, it just gives you instant feedback before you share the code with your colleagues, just like a linter.
It’s a linter that works if the dice roll well . This is obviously not great, and if we can do something with static analysis that works 100% of the time, that’s better; let’s do that. A mechanical code coverage tool or fuzzer will do a better job most of the time. But for cases which are quite subjective or rely on a rudimentary understanding of the written word, I’ll take it.
I have seen a couple of instances of this kind of thing, and if it were to become standard practice, I don’t think I’d mind. And based on my somewhat out-of-date understanding of neural networks (which peaked around the time of the “Attention is all you need” paper), I expect this would be possible with smaller, local models; you wouldn’t need the power of Anthropic’s GPU farm.
There are companies using LLMs to find anomalies in logs and alert a human who can investigate. Some pentesters use LLMs to try a ton of stuff which probably won’t work, but might. These tools don’t have to be perfect. Anomaly detection is useful if it is conservative and misses stuff; penetration testing just needs to get one success, even if there’s a thousand failures. These are machines. They can bang on walls until the walls break.
And of course, they have replaced search eng

[truncated]
