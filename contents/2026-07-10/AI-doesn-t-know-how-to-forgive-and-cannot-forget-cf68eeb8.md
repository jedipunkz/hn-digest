---
source: "https://tejassuds.com/blog/ai-cannot-forget"
hn_url: "https://news.ycombinator.com/item?id=48859047"
title: "AI doesn't know how to forgive and cannot forget"
article_title: "AI doesn't know how to forgive and cannot forget | Tejas Parthasarathi Sudarshan"
author: "tejassuds"
captured_at: "2026-07-10T13:04:26Z"
capture_tool: "hn-digest"
hn_id: 48859047
score: 16
comments: 11
posted_at: "2026-07-10T12:37:40Z"
tags:
  - hacker-news
  - translated
---

# AI doesn't know how to forgive and cannot forget

- HN: [48859047](https://news.ycombinator.com/item?id=48859047)
- Source: [tejassuds.com](https://tejassuds.com/blog/ai-cannot-forget)
- Score: 16
- Comments: 11
- Posted: 2026-07-10T12:37:40Z

## Translation

タイトル: AIは許す方法を知らず、忘れることができない
記事タイトル：AIは許し方を知らず、忘れられない |テジャス パルタサラティ スダルシャン
説明: マシン メモリの実際の配管、なぜメモリが私たちのメモリのように減衰しないのか、およびどのアーキテクチャにも存在しない 1 つの操作について説明します。

記事本文:
AIは許す方法を知らず、忘れることができない |テジャス パルタサラティ スダルシャン テジャス パルタサラティ スダルシャン テジャス。研究ブログ サイドクエスト 映画 読んだもの freewrite バックエッセイ
AI は許す方法を知らず、忘れることもできない
マシンのメモリの実際の配管、なぜ私たちのメモリのように減衰しないのか、そしてどのアーキテクチャにも存在しない 1 つの操作を見てみましょう。
このフレーズは「許す」と「忘れる」であり、この 2 つはいとこであるかのように言います。そうではありません。忘れることはあなたに起こることです。許しはあなたがするものです。 1 つは基板の特性であり、もう 1 つはその上で学習されるスキルであり、機械にはそのどちらもありません。それは忘れることができません、なぜならその中には朽ちるように作られたものは何もないからです。そして、許し方を知りません。なぜなら、許しというのは、私たちが出荷するアーキテクチャには実際には備わっていないからです。
これを配管のレベルで文字通りに説明したいと思います。これは雰囲気のように聞こえる主張の 1 つであり、実際には単なるエンジニアリングだからです。
4つの思い出、どれも色褪せない
「AI は決して忘れない」と言うとき、人々は通常、1 つのシステムを指して、それが全体であると想像しています。それらを分離するのに役立ちます。最新の AI システムはあなたの過去を 4 つの異なる場所に保存し、それぞれが独自の壊れた方法で忘れるか、まったく忘れません。
重みは最も古い記憶です。モデルがトレーニング中に見たものは何十億ものパラメーターに塗りつぶされており、一度入ってしまうと、それを取り出す方法がわかりません。これは政策上のギャップではなく、未解決の研究問題です。機械の非学習 、最初から再トレーニングせずに特定のトレーニング データの影響を取り除くことは、実際の規模では未解決です。データベースから行を削除できます。残されたグラデーションは削除できません。これが「忘れられる権利」が衝突し続ける静かな理由だ

g 機械学習の場合: 法則では、忘れることは削除であると想定されており、基板には削除がありません。
コンテキスト ウィンドウは作業メモリであり、調光器ではなく、照明のスイッチのように忘れてしまいます。ウィンドウ内では、アテンションはすべてのトークンを同様に到達可能なものとして扱います。最初の文は最後の文と同じように親しみやすく、年齢を重ねても柔らかくなりません。そしてセッションが終了し、すべてが一度に完全に消えてしまいます。 Eidetic または void の 2 つの状態。人間は生涯をこの 2 つの間の空間で過ごし、その空間にはほぼすべての社会機構が存在します。
検索ストアは、これが本格的に行われる場所です。埋め込みはぼやけません。コサイン類似度には時間軸がありません。 3 年前の記憶が昨日の記憶とまったく同じ忠実度で再現され、遠い昔のフェルトセンスはまったく付加されません。ベクトル空間では、過去全体が同時に存在し、最近傍検索が 1 つ離れたところにあります。過去は外国ではありません。隣の列です。
そして、ログ、スナップショット、レプリカ、バックアップは、意図的な忘れであっても分散システム プロジェクトであることを意味します。つまり、廃棄、保存期間、バックアップは、実行したと思われる削除を超えて存続します。
イベント
後で
高い
低い
AI検索
壊滅的な忘却
人間の記憶
人間の記憶は減衰し、定着します。細部はすぐに消えていきますが、要点はフロアに残ります。機械の回収は決して上から外れることはありません。そして、ニューラルネットが忘れてしまっても、朽ちることはなく、崖から落ちてしまいます。
忘れる機械は間違った種類のものです
ここが「AIは忘れない」よりも奇妙に見える部分だ。ニューラル ネットワークは常に忘れますが、それが非常に苦手です。
新しいタスクでネットワークをトレーニングすると、古いタスクが大規模に上書きされる傾向があります。これは、1989 年にマクロスキーとコーエンによって命名された壊滅的な忘却であり、今でもあなたが忘れる理由となっています。

すでに知っていることを危険にさらすことなく、デプロイされたモデルに何か新しいことを気軽に教えないでください。それが上のグラフのテラコッタの崖です。突然の完全な喪失に至るまで完璧な保持力を維持します。
そして、単一のコンテキスト内であっても、変圧器は偶然に忘れてしまいます。 「Lost in the Middle」の研究 (Liu et al., 2023) では、モデルが長いコンテキストの最初と最後に配置された情報を、途中に埋もれた同じ情報よりもはるかによく思い出すことが示されました。誰もそれを設計しませんでした。それは出現したデッドゾーンであり、アーキテクチャがたまたま弱いところに降りかかる無原則の忘却です。
コンテキストの始まり
終わり
思い出す
デッドゾーン
うっかり忘れ物。同じ事実は、長い文脈の端ではよく思い出されますが、途中ではあまり思い出されません。これは仕様によるものではなく、注意の使い方の副作用としてです。
つまり、実際の状況はすべてを記憶する機械ではありません。それは完全に記憶するか、ほとんど何もせずにすべてを失うかのどちらかであるマシンです。そしてその中間こそが人間の記憶なのです。
忘れることは機能であり、バグではありません
私たちが「忘れる」と呼ぶものは、それ自体が腐敗するわけではありません。ヘルマン・エビングハウスは 1880 年代にこの曲線を測定し、意味のない音節を記憶し、それらが滑るのを観察しました。最初は速く、その後は横ばいになります。しかし、現代の読書は興味深いものです。リチャーズとフランクランドは、2017 年の論文『記憶の持続性と一時性』で、「忘却は一般化を助けるために存在する、能動的で適応的なプロセスである」と直接主張しました。あらゆる細部をフル解像度で維持していた脳は、自らの過去に過剰適合してしまいます。詳細を薄めることで、要点を再利用可能なものに昇格させることができます。忘れることは、目的を持って圧縮することです。
その根底には機械にはないメカニズムがあります。記憶を思い出すとき、あなたはそうではありません

それを読んで、あなたはそれを開きます。思い出すという行為は、記憶を一時的に不安定にし、その後、変化した記憶を再保存します。これは再統合 (Nader、Schafe、LeDoux、2000) であり、人間の記憶が基本的に読み取り、変更、書き込みであることを意味します。すべてのリコールは小さな編集です。マシンのメモリは読み取り専用です。取得は、何回取得したか、または現在わかっている内容に関係なく、バイトごとにそのままのベクトルを返します。
すべてのリコールは編集です
機械
取り出す→戻す
記録は決して触れられない
記憶を思い出すことで記憶を編集します。マシンはそれらを変更せずに返します。これが、人間の記憶の感情的な負荷がその内容とは別の軌道で消えていき、何かが痛かったことを再感じることなく思い出すことができる理由です。ベクター ストアではこれらを分離するものは何もありません。
最後のポイントは静かなポイントです。人間の場合、記憶に付随する感情とその事実は別々に減衰します。時間は起こったことを消してはくれません。それはそこから熱を排出します。読み取り専用ストアには熱の概念がまったくないため、熱を排出する方法がありません。レコードとそのチャージは同じ不変オブジェクトです。
タイトルの話に戻ります。エンジニアが実装できる方法で許しを定義しようとすると、驚くほど具体的なことがわかります。それは、記録を完全な忠実度で保持し、記録に動作をさせるオプションは完全に利用できる一方で、記録に動作をさせるのをやめるということです。簡単にできるときは、意図的にそれを維持し、再評価し、それに基づいて行動しないでください。
ここでツールボックスを見てください。これを行うものがないことがわかります。
マシンの学習解除は強制削除です。それは記憶喪失であり、許しではありません、あなたはそれを手放さなかった、あなたは記録を破壊しました。
TTL、減衰スケジュール、および期限切れのメモリは削除されます。タイマーをかけるのを忘れる。
RLHF と微調整は、すべてを一度に包括的に設定します。 ｐは無いよ

えー、記憶の猶予、いいえ、「これを最大の解像度で保持し、今回はそれに条件を付けないことを選択します。」
そのための操作はありません。私が出荷したり読んだりしたどのフレームワークにも含まれていません。最も近いのは口止めだ。事件を決して話題にしないように設定されたシステムだが、これも許しではない。なぜなら、許しが意味を持つのは、本当に怒りがあり、それを受け入れなかった場合だけだからだ。覚えているかどうかに関係なく同じように振る舞うストアには、その選択を表現する方法さえありません。
なぜ私たちはこれを取得でき、彼らは取得しなかったのか
正直な答えはコストです。生物学的記憶は代謝的に高価であり、ニューロンとシナプスは維持するためにエネルギーを消費するため、脳は剪定、圧縮、解放するという絶え間ないプレッシャーにさらされています。希少性が教師です。私たちが忘れてしまうのは、抱え続けるのはお金がかかるからであり、それと関連した理由で物事を手放すのだと私は思います。不満を抱え続けるのもお金がかかるからです。どちらの種類のグレースも、予算の下で進化しました。
機械にはそんな予算はありません。ストレージは実質的に無料です。保持されている記憶には何の費用もかかりませんし、保持されている恨みにも費用はかかりません。したがって、解放へのプレッシャーは決してありません。手放す必要があるほど高価なものはありません。
これは暗いというよりはむしろ本当に興味深いと思います。なぜなら、これは修正がどこから来なければならないかを示しているからであり、これはより大きなコンテキストウィンドウやより優れた価値観の文書ではないからです。機械がその中間、優雅な減衰、分離可能な感情の負荷、機械に動かされずに保持できる記憶を学習するとしたら、それは記憶と不満に何らかのコストを与えることから来るでしょう。希少性から。私たちと同じ先生でした。
それまでは、何を展開しているのかを正確に把握しておく価値があります。私たちは、不完全な想起を中心に静かに進化してきた人間関係の真ん中に、完全な想起システムを導入しています。二度目のチャンスは実際には政策ではありませんでした。それらは人工物だった

ハードウェアに同梱されている贈り物です。すべてを完全に忠実に記憶するものを永久に構築しても、より良いバージョンは得られません。あなたは、一度も最初の慈悲が与えられたことのないもの、求められることもなく忘れ去られてやってくるものを手に入れるのです。
表紙について: エッセイ全体が 1 コマに収められています。左側には、記憶ノードの密集した格子があり、すべてが同じ明るさで点灯し、接続は損なわれておらず、距離によって薄くなるものは何もありません、それが機械の記憶であり、同時に存在する過去です。右側に向かって格子がほつれ、ノードが薄くなり、接続が切断され、最後のノードが緩んで火花として漂います。それが人間の曲線であり、剪定のように忘れられ、形が生き残るためにディテールが手放されます。この図は議論です。一方は決して色褪せないが、もう一方は意図的に色褪せます。
読む価値のある情報源: 壊滅的な干渉に関する McCloskey & Cohen (1989)。劉ら。 （2023）、「道に迷った」。リチャーズ & フランクランド (2017)、「記憶の持続性と儚さ」。 Nader、Schafe、LeDoux (2000) 再統合について。オリジナルの曲線については Ebbinghaus (1885) が挙げられます。
好奇心旺盛な人からの連絡はいつでも嬉しいです。 tejas@winsen.ai

## Original Extract

A look at the actual plumbing of machine memory, why none of it decays the way ours does, and the one operation no architecture has.

AI doesn't know how to forgive and cannot forget | Tejas Parthasarathi Sudarshan Tejas Parthasarathi Sudarshan tejas . Research Blog Sidequests Movies Things I Read freewrite Back Essay
11 min read AI doesn't know how to forgive and cannot forget
A look at the actual plumbing of machine memory, why none of it decays the way ours does, and the one operation no architecture has.
The phrase is forgive and forget , and we say it like the two things are cousins. They aren't. Forgetting is something that happens to you. Forgiveness is something you do. One is a property of the substrate and the other is a skill learned on top of it, and a machine has neither. It cannot forget, because nothing in it was built to decay. And it doesn't know how to forgive, because forgiveness is an operation no architecture we ship actually has.
I want to make that literal, at the level of the plumbing, because this is one of those claims that sounds like a mood and is actually just engineering.
Four memories, none of them fade
When people say "AI never forgets," they're usually gesturing at one system and imagining it's the whole thing. It helps to separate them. A modern AI system holds your past in four different places, and each one forgets in its own broken way, or not at all.
The weights are the oldest memory. Whatever a model saw in training is smeared across billions of parameters, and once it's in, we do not know how to get it out. This isn't a policy gap, it's an open research problem. Machine unlearning , removing the influence of specific training data without retraining from scratch, is unsolved at any real scale. You can delete the row from your database. You cannot delete the gradient it left behind. This is the quiet reason "the right to be forgotten" keeps colliding with machine learning: the law assumes forgetting is a delete, and the substrate has no delete.
The context window is the working memory, and it forgets like a light switch, not a dimmer. Inside the window, attention treats every token as equally reachable; your first sentence is as addressable as your last, with no softening for age. Then the session ends and it is all gone, completely, at once. Two states, eidetic or void. Humans spend our entire lives in the space between those two, and that space is where nearly all of the social machinery lives.
The retrieval store is where this gets its teeth. Embeddings don't blur. Cosine similarity has no time axis. A memory from three years ago is returned at exactly the fidelity of one from yesterday, with no felt sense of long ago attached. In vector space the whole past is co-present, sitting one nearest-neighbor lookup away. The past isn't a foreign country. It's the adjacent row.
And the logs , the snapshots and replicas and backups, mean that even deliberate forgetting is a distributed-systems project: tombstones, retention windows, backups that outlive the delete you thought you ran.
event
later
high
low
AI retrieval
catastrophic forgetting
human memory
Human memory decays and settles. Detail falls away fast, gist survives at a floor. Machine retrieval never moves off the top. And when neural nets do forget, they don't decay, they fall off a cliff.
The forgetting machines do have is the wrong kind
Here's the part that makes the picture stranger than "AI never forgets." Neural networks forget constantly, they're just terrible at it.
Train a network on a new task and it tends to overwrite the old one, wholesale. This is catastrophic forgetting , named by McCloskey and Cohen back in 1989, and it is still the reason you can't casually teach a deployed model something new without risking what it already knew. That's the terracotta cliff in the chart above: perfect retention right up until a sudden, total loss.
And even within a single context, transformers forget by accident. The "Lost in the Middle" work (Liu et al., 2023) showed that models recall information placed at the beginning and end of a long context far better than the same information buried in the middle. Nobody designed that. It's an emergent dead zone, an unprincipled forgetting that lands wherever the architecture happens to be weak.
start of context
end
recall
the dead zone
An accidental forgetting. The same fact is recalled well at the edges of a long context and poorly in the middle, not by design, but as a side effect of how attention is spent.
So the real situation isn't a machine that remembers everything. It's a machine that either remembers perfectly or loses everything, with almost nothing in between. And that in between is exactly what human memory is.
Forgetting is a feature, not a bug
The thing we call "forgetting" isn't decay for its own sake. Hermann Ebbinghaus measured the curve in the 1880s, memorizing nonsense syllables and watching them slip: fast at first, then leveling off. But the modern reading is the interesting one. Richards and Frankland made the case directly in their 2017 paper The Persistence and Transience of Memory : forgetting is an active, adaptive process that exists to help you generalize. A brain that kept every detail at full resolution would overfit to its own past. Letting the specifics fade is how the gist gets promoted to something reusable. Forgetting is compression with a purpose.
There's a mechanism underneath it that machines simply don't have. When you recall a memory, you don't read it, you reopen it. The act of remembering makes the memory briefly labile and then re-stores it, changed. This is reconsolidation (Nader, Schafe, and LeDoux, 2000), and it means human memory is fundamentally read-modify-write. Every recall is a small edit. Machine memory is read-only : retrieval returns the vector untouched, byte for byte, no matter how many times you pull it or what you now know.
every recall is an edit
Machine
retrieve → return
the record is never touched
You edit memories by remembering them. A machine returns them unchanged. This is why the emotional charge of a human memory fades on a different track from its content, you can recall that something hurt without re-feeling it. Nothing in a vector store separates those.
That last point is the quiet one. In people, the feeling attached to a memory and the facts of it decay separately. Time doesn't delete what happened; it drains the heat out of it. A read-only store has no way to drain the heat, because it has no notion of heat at all. The record and its charge are the same immutable object.
Which brings us back to the title. If you try to define forgiveness in a way an engineer could implement, you get something surprisingly specific: keep the record, at full fidelity, and stop letting it drive the behavior, while the option to let it drive the behavior is still fully available. Retain, re-weight, and don't act on it, on purpose, when you easily could.
Now look at the toolbox and notice that nothing does this.
Machine unlearning is a forced delete . That's amnesia, not forgiveness, you didn't let it go, you destroyed the record.
TTLs, decay schedules, and expiring memories are eviction . Forgetting on a timer.
RLHF and fine-tuning set a blanket disposition over everything at once. There's no per-memory grace, no "hold this one at full resolution and choose, this time, not to condition on it."
There is no op for that. Not in any framework I've shipped or read. The closest anyone gets is muzzling, a system configured to never bring the incident up, which isn't forgiveness either, because forgiveness only means anything when resentment was genuinely available and you didn't take it. A store that behaves identically whether it remembers or not has no way to even express the choice.
Why we got this and they didn't
The honest answer is cost. Biological memory is metabolically expensive, neurons and synapses burn energy to maintain, so brains are under constant pressure to prune, compress, and let go. Scarcity is the teacher. We forget because holding is expensive, and I'd argue we let things go for a related reason: carrying a grievance is expensive too. Grace, of both kinds, evolved under a budget.
Machines have no budget like that. Storage is effectively free. A held memory costs nothing, a held grudge costs nothing, so there's no pressure toward release, ever. Nothing is expensive enough to need letting go of.
I find this genuinely interesting rather than gloomy, because it points at where the fix would have to come from, and it isn't a bigger context window or a better values document. If machines ever learn the in-between , graceful decay, separable emotional charge, memory you can hold without being run by it, it will come from making memory and grievance cost something. From scarcity. The same teacher we had.
Until then it's worth being precise about what we're deploying. We are putting perfect-recall systems into the middle of human relationships that quietly evolved around imperfect recall. Second chances were never really a policy. They were an artifact of biology, a gift that shipped with the hardware. Build something that remembers everything at full fidelity forever, and you don't get a better version of us. You get something that has never once been offered the first mercy, the one that arrives, unasked, as forgetting.
About the cover: it's the whole essay in one frame. On the left, a dense lattice of memory nodes, every one lit to the same brightness, connections intact, nothing dimming with distance, that's machine memory, the co-present past. Toward the right the lattice frays: nodes thin out, connections drop, and the last of them come loose and drift off as sparks. That's the human curve, forgetting as pruning, detail let go so the shape can survive. The picture is the argument: one side never fades, the other side fades on purpose.
Sources worth reading: McCloskey & Cohen (1989) on catastrophic interference; Liu et al. (2023), "Lost in the Middle"; Richards & Frankland (2017), "The Persistence and Transience of Memory"; Nader, Schafe & LeDoux (2000) on reconsolidation; and Ebbinghaus (1885) for the original curve.
Always happy to hear from curious people. tejas@winsen.ai
