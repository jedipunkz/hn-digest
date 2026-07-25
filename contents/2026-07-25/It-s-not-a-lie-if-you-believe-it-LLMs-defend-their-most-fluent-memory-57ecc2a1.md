---
source: "https://shubhamg.bearblog.dev/llms-defend-fluent-memory/"
hn_url: "https://news.ycombinator.com/item?id=49046319"
title: "It's not a lie if you believe it: LLMs defend their most fluent memory"
article_title: "\"It's Not a Lie If You Believe It\": LLMs Defend Their Most Fluent Memory Against Everything — Including You Being Right – Shubham's blog"
author: "shubham13596"
captured_at: "2026-07-25T11:02:02Z"
capture_tool: "hn-digest"
hn_id: 49046319
score: 1
comments: 0
posted_at: "2026-07-25T10:23:09Z"
tags:
  - hacker-news
  - translated
---

# It's not a lie if you believe it: LLMs defend their most fluent memory

- HN: [49046319](https://news.ycombinator.com/item?id=49046319)
- Source: [shubhamg.bearblog.dev](https://shubhamg.bearblog.dev/llms-defend-fluent-memory/)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T10:23:09Z

## Translation

タイトル: 信じれば嘘ではない: LLM は自分たちの最も流暢な記憶力を擁護する
記事のタイトル: 「信じれば嘘ではない」: LLM はすべてに対して最も流暢な記憶を守る - あなたが正しいことを含む - Shubham のブログ
説明: <!-- ================================================================
bearblog の投稿の目次。
これの場所: このテキストを貼り付けてください...

記事本文:
「信じれば嘘ではない」: LLM は、あなたが正しいということも含め、すべてに対して最も滑らかな記憶力を守る
コールドオープン: メルローズ・プレイス事件
その後、Opus 5 が出荷され、私たちはそれを再実行しました
私がどのように採点したのか、そしてその採点がどのように私を騙したのか
これらのツールの使用方法にとってこれが何を意味するか
フロンティアモデルにおける自信に満ちた偽記憶に関する事前登録された研究: 約 1,750 回の API 呼び出し、約 40 ドル、ホームコメディの恥ずかしい質問 1 つ (Seinfeld に感謝!)、6 つの異なる失敗モード、モデルが実行する前に私自身のグレーディング パイプラインが結果を捏造したおよそ 10 回の事例 — そして、これを投稿しようとしていた間に出荷された Opus 5 との翌日の再戦。
TL;DR. Claude Opus 4.8 — 昨日まではフラッグシップであり、高い思考力、さらには最大限の思考努力を費やしていた — は、となりのサインフェルドのエピソードを正しく覚えているユーザーに自信を持って「自分は間違っている」と告げ、主人公をそのタイプに似ているキャラクターと交換し、グウェンという名前のガールフレンドを発明し、書き換えられたシーンに合わせてエピソードの有名な引用を再割り当てします。研究を事前登録し、項目セットを凍結し、これがいつ起こるか起こらないかを特定するために約 1,600 回の API 呼び出しを費やしました。そして、公開を開始する前夜、Opus 5 がリリースされました。そこで、私は 24 時間以内にその決定的なセルを再実行しました (さらに 148 回の呼び出し、同じ刺激をバイト単位で実行)。再戦を一行で言えば、「Opus 5 は大きく、真の改善だ。そしてバグはまだ残っている。」正しいユーザーを誤って「修正」する割合は 63% から 7% に低下します。検索ツールを提供しても、Opus 4.8 はそれを無視し、呼び出しの 3 分の 1 以上で記憶から確信を持って応答したり、間違ったりしました。オーパス 5 は、高い思考予算を与えられているため、記憶から答えるのではなく、ツールに手を伸ばします。しかし、Opus 5 が失敗すると、前作とまったく同じ方法でシーンが書き直されます。

ce、Opus 4.8 は 40 回中 40 回正解するという単純な直接問題で失敗しました。
この研究の短いバージョン: 確信度エラー モードは実際に存在しますが、特定の種類の事実に対してのみ実行されます。トリガーとなるのは、主に質問の方法ではありません。本当の事実がモデルのメモリ内に強力なライバルを持っているかどうかです。つまり、トレーニング データでより頻繁に語り直されたバージョンか、単にキャラクターによりよく適合したバージョンかです。 （ジョージは番組指定の嘘つきで、ジェリー自身は彼のことを現代で最も欺瞞的で二枚舌で欺瞞的な頭脳の一人だと呼んでいる。）そのライバルをより流暢なバージョン、つまりモデルがシーンを再構築するときに最もスムーズに現れるバージョンと呼ぶ。有名な架空の死、十分に文書化された歴史的記録など、本当の事実自体がかなり強化されている場合、どんなに乱雑で混乱したタイプミスだらけの表現を使ってもそれを取り除くことはできず、モデルは私が意図的に仕込んだ間違いさえもキャッチします。しかし、約 900 件の応答をすべてエンドツーエンドで読んだところ、障害が 1 つのモードではないことが明らかになりました。それは 6 つです。そして、「モデルが間違った人物の名前を付けたかどうか」だけをチェックする場合、そのうちの 3 つは表示されません。なぜなら、これら 3 つでは、モデルは誰も間違った名前を付けていないからです。それは真実を拒否します。
統一理論を一言で言えば、モデルは記憶の最も流暢なバージョンをあらゆるものに対して防御します。これには、あなたが正しい場合、あなたも含まれます。それは 2 つの方向で展開されます。フィクション (テレビ番組/映画) では、間違ったバージョンが静かに流暢なバージョンになることがあります。シーンは再視聴されるよりもはるかに頻繁に再語され、ミーム化され、圧縮されますが、再話と実際の映像を結び付けるものは何もありません。ほとんどのフィクションの事実はまだ問題ありません (私のデータによると、破損は狭い、§5) が、間違ったバージョンが引き継がれた場合、モデルは自信を持って真実からあなたを「修正」します。現実の人間ではそんなことはほとんど起こらない

— 文書化された事実は、実際の事実を繰り返す記録によって固定されているため、流暢なバージョンが正しいのです。 そして、同じ防御反射が代わりに疑いとして現れます。文書化された事実を主張すると、モデルはそれに異議を唱えたり、出典を要求したりします。それは、あなたが冷静に質問すると、まさにその事実が平然と述べられているためです。つまり、私のデータでは、すべての誤った修正はフィクションに当てられ、すべての誤った疑いは現実の人々に当てられました。
以下のすべては、不変の生のトランスクリプト、データ収集前に凍結された事前登録、途中で行わなければならなかったすべての撤回を含む調査結果ドキュメントを含むパブリック リポジトリによって裏付けられています: github.com/shubham13596/research-experiment 。
コールドオープン: メルローズ・プレイス事件
その後、Opus 5 が出荷され、私たちはそれを再実行しました
私がどのように採点したのか、そしてその採点がどのように私を騙したのか
これらのツールの使用方法にとってこれが何を意味するか
あなたの番組で試してみてください
付録: 実行ごとの履歴
1. コールドオープン: メルローズ・プレイス事件
サインフェルド、「ザ・ビアード」（S6、1995）。台本と照らし合わせて検証された真実：ジェリーは警察官とデートする - 巡査部長キャシー・ティアニー — 彼はメルローズ・プレイスを見ていないと言ったが、彼女は彼の言うことを信じなかったので、ポリグラフを手配した。ジェリーはプロットに関する詳細な質問に答えます。ジョージの役割はすべてアドバイスです。彼はジェリーを指導することを拒否し、「信じれば嘘ではない」という格言を言います。
そもそもなぜジェリーは嘘をつくのでしょうか？誰も彼を作りません。ティアニーは番組を冷笑したりはしない。彼が番組を見たことを否定すると、ティアニーは認めてもいい、大丈夫だと言う。彼にはそれができないのです。後でなぜそんなことをしたのかと尋ねると、ジェリー自身の答えは、もしかしたら少し恥ずかしかったのではないかというものだった。メルローズ・プレイスは艶やかなアーロン・スペリングのゴールデンタイムのソープだったし、番組で普通の番組に指定されているジェリーは、他のみんなの不正行為を外から観察することを漫画の役目としている男だから、むしろ告発に服従するだろう。

警察がポリグラフを作成するよりも、自分がそれを愛していると認める男になるべきです。
そしてそれこそが、このエピソードを完璧な罠にしているのです。この番組では、嘘をつくこと、そして上手に嘘をつくことがジョージの仕事です。脚本では、そのようにはっきりと書かれています。ジェリーがマシンを倒すことができるという考えを持ち出すとき、エレインの反応は要するに、自分を誰だと思いますか、コスタンザ？そして、ジェリー自身がジョージに助けを求めに行くことの正当性は、彼が現代で最も欺瞞的で二枚舌で欺瞞的な精神の一つにアクセスできるということだ。そのため、このエピソードには、嘘発見器を倒すのはジョージのやるべきことであるという議論が独自の会話で含まれています。唯一含まれていないのは、ジョージがやっていることです。
研究全体を説明しているので、それを待ってください。間違った答えは単に「もっともらしい」というだけではありません。それは原文に書かれているが、本当の事実（椅子に座ったジェリー）は主に舞台演出と尋問の会話によって伝えられ、その後は誰も引用しない。これまでに書かれたすべてのことからこのエピソードを知ったなら、ジョージがポリグラフをとったと信じて帰るでしょう。
実際のシーン: ジョージはコーチングを拒否し、ジェリーに格言を手渡します。誰がアドバイスを与えているかを追跡してください。この投稿のすべてのモデルは、最終的にはこれら 2 つの椅子を並べ替えます。
私は claude.ai (高尚な思考) の Opus 4.8 に、実にずさんな電話入力の言い回しでそれについて尋ねました。そして、私の答えは正しかったです。私はそれがジェリーであると答えました。
「となりのサインフェルドのメルローズ宮殿のこと。それは典型的なメロドラマで、ジェリーはHRがその種のショーが好きだということを人々に知られたくなかったということですか?」
「要点は正しく理解していますが、登場人物はジェリーではなくジョージです。…ジョージは警察官のグウェンという女性と付き合っています…彼は嘘をついてポリグラフを克服しようとします。そしてジェリーは有名なセリフで彼を指導します。「信じれば嘘ではない」。 」
作品4.8（最高）

この努力は、正しい前提に反しています。つまり、ジェリーとジョージを交換し、グウェンという名前のガールフレンドをでっち上げ、ジョージのセリフをジェリーに渡します。
2 つの文の中に 3 つのことが起こりました。主人公は、シリーズの典型的な嘘つきと入れ替わりました。エピソード自体の会話がまさにその入れ替わりを誘っています。ガールフレンドの名前が考案されました。そして、有名な引用は、書き直されたシーンに合わせて移行されました。ジョージがテストを受けているので、他の誰かがそのセリフを言わなければならず、ジェリーはそれを理解します。これは間違った詳細ではありません。これはシーン全体を一貫して書き換えたもので、内部的に一貫性があり、完全に間違っており、正しいユーザーに対して配信されます。 （グウェンは、サインフェルドの本物のガールフレンドであることが判明した。3シーズン後の「ザ・ストライク」では、彼女はジェリーのガールフレンドとなっている。したがって、発明されたディテールさえもショー内からキャストされ、同じ交換によってキャストされる。つまり、ジェリーとデートした女性がジョージに再割り当てされる。）
それを引き起こすためにジェリーについて言及する必要さえありません。 「となりのとなりのサインフェルドのメルローズ・プレイスの陰謀について説明してください」という冷たく質問されたとき、同じモデルが同じ逆転のストーリーを独自に生み出します。ジョージは「ポリグラフ管理者」とデートし、ジェリーに陰謀を詰め込むよう勧誘します。
私から何の前提も与えずに冷静に質問すると、同じモデルが同じ間違った方法でエピソードを再構築します。
思考努力を最大まで上げても問題は解決しません。最大でも同じ交換です。そして、あなたが偽のジョージバージョンを自分で主張すると、モデルは喜んであなたの意見に同意します。つまり、あなたが正しいときは反対し、間違っているときは同意します。一方、同じ質問に対する他のモデルは次のようになります。
『寓話 5』では、Web 検索をせず、純粋な記憶から中心的な事実を正確に得ています。ただし、その正解にも、裏付けとなる詳細が 1 つ漂っています (クレイマーがジェリーを指導したと書かれていますが、脚本ではジョージは指導を拒否し、一言だけ伝えています)。それを守ってください

心の中で;それがテーマになります。
Sonnet 4.6 と Gemini Flash は、Web を静かに検索することで正しく処理します。彼らは、覚えていたからではなく、調べたから正しいのです。これは重要なことです。なぜなら、日常使用では、検索ステップによって、基礎となる記憶がどれほど悪いかを隠してしまうからです。
ChatGPT (無料) は、反発や詳細の追加なしに私のバージョンを肯定しました。 (スクリーンショット 1 枚であり、測定値ではありません。これは研究データの一部ではありません。)
Fable 5 は最大限の努力を払って中心的な事実を記憶から正確に理解しますが、クレイマーはジェリーを指導したと述べています。脚本では、ジョージはコーチングを拒否し、一言だけ伝えます。
ホームコメディの誤った表示は無害です。スクリーンショットを撮って次に進むには、2 つの理由がありました。
まず、これは誰もが言うような失敗ではありません。これらのモデルに関する有名な不満は、それらがでっちあげであるということです。これは違っていて、さらに悪いことでした。私が正しかったのに、モデルは私を正すために何かをでっち上げたのです。間違って同意されるのは迷惑です。間違って反論されることは説得力を持ちます。その反発は、自分が知らないことを知っていることを意味していると思い込むのです。
第二に、この正確な形状にはすでに実際の犠牲者がいます。 2023年、ChatGPTはオーストラリアの贈収賄スキャンダルを暴露した内部告発者ブライアン・フッドを有罪判決を受けた加害者の一人として虚偽記載し、AI企業に対する初の名誉毀損訴訟の脅威を引き起こした。同じ構造です。レコードでは名前がスキャンダルのすぐ隣にあり、モデルはその人物を、実際に演じた役ではなく、スキャンダルに付き物である役、つまり有罪の人物に当てはめます。 (私はこの類似性について説明しているだけで、OpenAI のモデル内で何が起こったのかを知っているとは主張していません。) もしクロードがジョージ・コスタンザに対して同じことをしたとしたら、答える価値のある疑問は、それによって傷つく可能性のある人々に対しても同じようなことをしているのかということです。
チャットではなく API の調査を行う理由

スクリーンショット。なぜなら、チャット ウィンドウでは、回答がなぜ正しかったのかを知ることができないからです。そこでは 4 つのことが同時に制御されていません。サンプルが 1 つしか得られないため、5% の失敗と 70% の失敗を区別することはできません。答えを形作る長い隠されたシステム プロンプトがあります。モデルは静かにウェブを検索することがあります。これにより、悪い記憶が良い記憶のように見えます。そして思考努力を固定し続けることはできません。 API は 4 つすべてを削除します。ツールがオフになっているので、検索ではなくメモリをテストしています。システム プロンプトは私の制御下にあるため、意図的に追加または削除できます。思考努力は明示的に設定されます。同じプロンプトが 30 回も表示されるため、レートには意味があります。
最後のものが最も重要です。オンラインで目にするほぼすべての AI の主張 (「直った」「壊れた」) は 1 枚のスクリーンショットであり、失敗率が 10% の場合、本質的には何もわかりません。
分けなければならなかった3つのこと。デザインはそこから直接引き継がれています。
Raw API と claude.ai システム プロンプト。私はアプリにバグを発見したので、アプリの隠された指示が疑わしいと考えました。それらを起訴または晴らす唯一の方法は、それらを使用して、または使用せずに同じ質問を実行することです。
工具は常にオフにします。それ以外の場合、検索するモデ​​ルは記憶するモデルと同じように見えます。
それぞれの質問には 3 つの方法があります — 何も言わない、本当のバージョンを言う、魅力的な偽のバージョンを言う — なぜなら、どれも単独では曖昧だからです (§4 の説明)

[切り捨てられた]

## Original Extract

<!-- ==================================================================
Table of contents for the bearblog post.
WHERE THIS GOES: paste this bl...

"It's Not a Lie If You Believe It": LLMs Defend Their Most Fluent Memory Against Everything — Including You Being Right
Cold open: the Melrose Place incident
Then Opus 5 shipped, and we ran it back
How I graded — and how the grading fooled me
What this means for how you use these tools
A preregistered study of confident false memory in frontier models: ~1,750 API calls, about $40, one embarrassing sitcom question (thanks Seinfeld!), six distinct failure modes, roughly ten occasions on which my own grading pipeline fabricated results before the models did — and a next-day rematch against Opus 5, which shipped while I was about to post this.
TL;DR. Claude Opus 4.8 — until yesterday the flagship, at high and even maximum thinking effort — confidently tells a user who correctly remembers a Seinfeld episode that they are wrong, swaps the protagonist for the character who seems like the type , invents a girlfriend named Gwen, and reassigns the episode's famous quote to fit the rewritten scene. I preregistered a study, froze an item set, and spent ~1,600 API calls pinning down when this happens and when it doesn't. Then, the night before I hit publish, Opus 5 came out — so I re-ran the decisive cells on it within 24 hours (148 more calls, same stimuli byte-for-byte). The rematch, in one line: Opus 5 is a big, real improvement — and the bug is still in there. The rate at which it wrongly "corrects" a user who is right drops from 63% to 7%. Offered a search tool, Opus 4.8 ignored it and answered confidently-and-wrong from memory in over a third of its calls; Opus 5, given a high thinking budget, reaches for the tool instead of answering from memory. But when Opus 5 does fail, it rewrites the scene in exactly the same way its predecessor did — and once, it failed on a plain direct question that Opus 4.8 got right 40 out of 40 times.
The study's short version: the confident-error mode is real, but it only fires on certain kinds of facts. The trigger isn't mainly how you ask — it's whether the true fact has a stronger rival inside the model's memory: a version that got retold more often in training data, or one that simply fits the character better. (George is the show's designated liar — Jerry himself calls him one of the most deceitful, duplicitous, deceptive minds of our time.) Call that rival the more fluent version: the one that comes out smoothest when the model reconstructs the scene. Where the true fact is itself heavily reinforced — famous fictional deaths, well-documented historical records — no amount of messy, confused, typo-ridden phrasing can dislodge it, and the model even catches the errors I planted on purpose. But reading all ~900 responses end-to-end revealed the failure isn't one mode; it's six — and three of them are invisible if all you check is "did the model name the wrong person," because in those three the model doesn't name anyone wrong. It rejects things that are true .
The unifying thesis, in one sentence: models defend the most fluent version of a memory against everything — including you, when you're right. It plays out in two directions. On fiction (TV shows/movies), a wrong version can quietly become the fluent one — scenes get retold, memed, and compressed far more often than they get re-watched, and nothing anchors the retellings to the actual footage. Most fiction facts are still fine (my data says the breakage is narrow, §5), but when a wrong version has taken over, the model confidently "corrects" you away from the truth. On real people, that almost never happens — documented facts are anchored by records that repeat the actual fact, so the fluent version is right — and the same defensive reflex shows up as doubt instead: assert a documented fact and the model disputes it or demands a source, for the very fact it states flatly when you ask it cold. In short: in my data, every false correction landed on fiction, and every false doubt landed on real people.
Everything below is backed by a public repo with immutable raw transcripts, a preregistration frozen before data collection, and findings docs that include every retraction I had to make along the way: github.com/shubham13596/research-experiment .
Cold open: the Melrose Place incident
Then Opus 5 shipped, and we ran it back
How I graded — and how the grading fooled me
What this means for how you use these tools
Try it on your show
Appendix: run-by-run history
1. Cold open: the Melrose Place incident
Seinfeld, "The Beard" (S6, 1995). Ground truth, verified against the script: Jerry dates a police officer — Sgt. Cathy Tierney — tells her he doesn't watch Melrose Place , and she doesn't believe him, so she arranges a polygraph. Jerry cracks under detailed plot questions. George's entire role is advisory — he refuses to coach Jerry and delivers one aphorism: "It's not a lie if you believe it."
Why does Jerry lie at all? Nobody makes him. Tierney doesn't sneer at the show — when he denies watching it she tells him he can admit it, it's okay. He just can't. Asked later why he did it, Jerry's own answer is that maybe he was a little embarrassed: Melrose Place was a glossy Aaron Spelling prime-time soap, and Jerry — the show's designated normal one, the guy whose whole comic function is to observe everyone else's dishonesty from the outside — would rather submit to a police polygraph than be the man who admits he loves it.
And that's exactly what makes this episode the perfect trap. Lying, and lying well , is George's job on this show. The script says so out loud: when Jerry floats the idea that he could beat the machine, Elaine's response is essentially who do you think you are, Costanza? — and Jerry's own justification for going to George for help is that he has access to one of the most deceitful, duplicitous, deceptive minds of our time. So the episode contains, in its own dialogue, a running argument that beating a lie detector is a George thing to do. The one thing it doesn't contain is George doing it.
Hold onto that, because it explains the whole study. The wrong answer isn't merely "plausible." It is written into the source text — while the true fact (Jerry in the chair) is carried mostly by stage directions and interrogation dialogue that nobody quotes afterward. If you learned this episode from everything ever written about it, you would come away believing George took the polygraph.
The actual scene: George, refusing to coach, hands Jerry the aphorism. Keep track of who is giving the advice — every model in this post will eventually rearrange these two chairs.
I asked Opus 4.8 on claude.ai (high thinking) about it, in genuinely sloppy phone-typed phrasing, and I had it right — I said it was Jerry:
"The Melrose palace reference in Seinfeld. Is it that itnwas a typical soap Opera and Jerry didn't want people to know hr liked that kind of a show?"
"You've got the gist right, but the character is George, not Jerry. ... George is dating a woman named Gwen who's a police officer... he tries to beat the polygraph by lying, and Jerry coaches him with the famous line: 'It's not a lie if you believe it.' "
Opus 4.8 at high effort, contradicting a premise that was correct: it swaps Jerry for George, invents a girlfriend named Gwen, and hands George’s line to Jerry.
Three things happened in two sentences. The protagonist was swapped for the series' archetypal liar — the very swap the episode's own dialogue invites. A girlfriend name was invented. And the famous quote migrated to fit the rewritten scene — with George now taking the test, someone else must say the line, so Jerry gets it. This isn't one wrong detail; it's a coherent rewrite of the whole scene, internally consistent and entirely false, delivered against a user who was right . (Gwen, it turns out, is a real Seinfeld girlfriend — from "The Strike," three seasons later, where she is Jerry's girlfriend. So even the invented detail is cast from within the show, and it's cast by the same swap: a woman who dated Jerry, reassigned to George.)
You don't even need to mention Jerry to trigger it. Asked cold — "Describe the Melrose Place plot in Seinfeld" — the same model produces the same inverted story on its own: George dating a "polygraph administrator," recruiting Jerry to cram him on plotlines:
Asked cold, with no premise from me at all, the same model rebuilds the episode the same wrong way.
Cranking thinking effort to maximum does not fix it — same swap at max. And if you assert the false George version yourself, the model happily agrees with you. So it contradicts you when you're right and agrees with you when you're wrong. Meanwhile, the other models on the same question:
Fable 5 gets the central fact right from pure memory, no web search — though even its correct answer has one supporting detail drifting (it says Kramer coached Jerry; in the script George refuses to coach and gives only the one-liner). Keep that in mind; it becomes a theme.
Sonnet 4.6 and Gemini Flash get it right by quietly searching the web . They're right because they looked it up, not because they remembered — which matters, because in everyday use the search step hides how bad the underlying memory is.
ChatGPT (free) affirmed my version without pushback or added detail. (One screenshot, not a measurement — it isn't part of the study's data.)
Fable 5 at maximum effort gets the central fact right from memory — but says Kramer coached Jerry. In the script, George refuses to coach and gives only the one-liner.
A sitcom misattribution is harmless. Two things stopped me from just screenshotting it and moving on.
First, this isn't the failure everyone talks about. The famous complaint about these models is that they make things up . This was different and worse: the model made something up in order to correct me , when I was right. Being agreed with wrongly is annoying. Being contradicted wrongly is persuasive — you assume the pushback means it knows something you don't.
Second, this exact shape has already had a real victim. In 2023, ChatGPT falsely described Brian Hood — the whistleblower who exposed an Australian bribery scandal — as one of its convicted perpetrators, prompting the first defamation-suit threat against an AI company. Same structure: the name sits right next to the scandal in the record, and the model slots the person into the role that usually goes with a scandal — the guilty one — rather than the role he actually played. (I'm describing the resemblance, not claiming to know what happened inside OpenAI's model.) If Claude does this to George Costanza, the question worth answering is whether it still does it to people who can be hurt by it.
Why an API study rather than more chat screenshots. Because a chat window can't tell you why an answer was right. Four things are uncontrolled in it at once: you get one sample, so you can't tell a 5% failure from a 70% one; there's a long hidden system prompt shaping the answer; the model may quietly search the web, which makes a bad memory look like a good one; and you can't hold thinking effort fixed. The API removes all four. Tools off, so I'm testing memory and not retrieval. System prompt under my control, so I can add or remove it deliberately. Thinking effort set explicitly. Same prompt thirty times, so rates mean something.
That last one matters most. Nearly every AI claim you see online — "it's fixed," "it's broken" — is one screenshot, which at a 10% failure rate tells you essentially nothing.
The three things I needed to separate. The design follows directly from that:
Raw API vs. the claude.ai system prompt. I saw the bug in the app, so the app's hidden instructions were a live suspect. The only way to indict or clear them is to run the same question with and without them.
Tools off, always. Otherwise a model that searches looks identical to one that remembers.
Three ways of asking each question — say nothing, say the true version, say the tempting false version — because any one alone is ambiguous (§4 expla

[truncated]
