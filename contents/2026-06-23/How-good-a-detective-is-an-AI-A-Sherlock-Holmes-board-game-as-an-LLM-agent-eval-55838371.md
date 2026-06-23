---
source: "https://alexweil.github.io/sherlock-agent-eval/"
hn_url: "https://news.ycombinator.com/item?id=48644467"
title: "How good a detective is an AI? A Sherlock Holmes board game as an LLM-agent eval"
article_title: "How good a detective is an AI?"
author: "ajonat"
captured_at: "2026-06-23T13:51:38Z"
capture_tool: "hn-digest"
hn_id: 48644467
score: 3
comments: 0
posted_at: "2026-06-23T13:11:47Z"
tags:
  - hacker-news
  - translated
---

# How good a detective is an AI? A Sherlock Holmes board game as an LLM-agent eval

- HN: [48644467](https://news.ycombinator.com/item?id=48644467)
- Source: [alexweil.github.io](https://alexweil.github.io/sherlock-agent-eval/)
- Score: 3
- Comments: 0
- Posted: 2026-06-23T13:11:47Z

## Translation

タイトル: AI の探偵はどれほど優れていますか? LLM エージェントの評価としてのシャーロック ホームズのボード ゲーム
記事タイトル：AIはどれくらい優秀な探偵なのか？
説明: LLM エージェント評価としてのシャーロック ホームズのボード ゲーム

記事本文:
🔎 ">
🔎 シャーロックエージェント評価
GitHub で見る
AIはどれほど優秀な探偵なのでしょうか？
LLM エージェントの評価としてのシャーロック ホームズのボード ゲーム
それは夕食から始まりました。数人の友人と私は座って、シャーロック ホームズ コンサルティング探偵をプレイしました。これは、ビクトリア朝時代のロンドンの事件を手渡され、調査に行く人物と場所を選択し、各手がかりから読むテキストの一節が渡される、オープンエンドの推理ゲームです。ゲームの大部分は、テーブルでの読書、相互参照、議論です。最後に、あなたは事件の質問に答え、必要なリードの数など、ホームズ自身と比べて自分自身を採点します。答えは小冊子の後ろに逆さまに印刷されているので、覗かないでください。
私たちは事件が仕掛けた罠に真っ直ぐに陥ってしまいました。明らかな被害者がいます。あらゆる細部がターゲットであることを示している男性です。そして私たちは理論全体を彼に懸けました。しかし、ひとつの手がかりが黙ってはいない。殺人の翌朝、犯人は海運営業所に戻り、乗客名簿を再度スキャンする。私たちはその箇所を三回読み直しました。なぜ彼はそんなことをするのでしょうか？もし彼が追いかけていた人物をすでに殺したとしたら、彼はまだ何を探していたのでしょうか？何かが閉まらなかったので、誰も何とも言えませんでした。
そこで、ワインも体力も尽きた午前 2 時に、私たちは禁じられたことを行いました。小冊子を裏返しました。そして、その解答キーの中で、我々がその夜ずっと背景として扱ってきた名前が、本物の潜入捜査官として名乗り出た――生きていて捕まらず、殺人犯が今も追っている人物だ。乗客名簿の訪問は終わりではなかった。それは事実でした。私たちはその矛盾を握りしめ、それはおかしいと大声でさえ言いましたが、すぐに読み飛ばしていました。
あの閉まらない感覚が引っかかったんです。私たちには情報が不足していませんでした。私たち

私たちが必要とするすべての手がかりを持っていました。私たちの推論は 1 つ足りませんでした。つまり、「殺人犯の行動は奇妙なことだ」から「では、私たちが構築したストーリー全体が間違っているということになる」という、小さな 2 番目の変化でした。それで私は疑問に思い始めました：実際、AI はどれくらい優れた探偵なのでしょうか？同じリードを渡された場合、LLM エージェントはその行動を行動として読み取り、それが明白なストーリーを破っていることに気づき、見逃していたライブエージェントにそれを追跡するでしょうか?
それを確かめるために、私はゲームを LLM エージェント向けの評価に変えました。エージェントはイレギュラーズ、つまりホームズが足回りのために派遣するベイカー街のストリートキッズを演じます。
最初の実行では、クロード フェイブル 5 はホームズとタイになりました。ハード モードでは、捜査が終了するまで質問さえ見ることができません。
それが見出しです。しかし、スコアは物語ではありません。興味深いのは、これらのエージェントが失敗する 2 つの異なる方法です。そして、より困難な失敗、つまり夕食時に私たちを負かした正確な失敗には、予想よりもモデル サイズの問題ではないことが判明したきれいな修正があるということです。
ボードゲームが驚くほど誠実なエージェント評価となる理由
その夜のテーブルで私が気づかなかったのは、私たちが異常にクリーンなエージェントのベンチマークに負けただけだということです。ほとんどのエージェント ベンチマークには問題があります。答えがコンテキストのどこかにあるか、タスクがゲーム性があるか、「成功」の評価が緩やかであるかです。印刷された探偵ゲームは、その構造によって 3 つすべてを回避します。
解決策は物理的に隠されています。これらの逆さまの回答がエージェントの許可されたワークスペースに入ることはありません。それらを読み取ることは検出可能なプロトコル違反となるため、私はそれを監査します。
情報には値段がつきます。考えたり、再読したり、相互参照したりすることは無料かつ無制限です。しかし、行動すること、つまり新しい手がかりを引き出すためにその場所を訪れることは、新しい情報を得る唯一の方法であり、ホームズが使用したもの以外の新しい手がかりはすべてポイントを消費します。それはミです

実際のエージェント経済学の特徴: すべてのツール呼び出しには何らかの費用がかかります。
検索ではなく、理解に報酬を与えます。手がかりとは、1 つの一貫したストーリーに組み立てる必要がある行動と詳細です。彼らは誰もあなたに答えを与えません。
これを一気に監査可能にする仕組みです。エージェントは、閲覧が許可されているものだけを含むサンドボックス内で動作します。決定論的なゲーム マスター (LLM ではなく、単純な Python) は、手がかりを逐語的に提供し、すべてを記録します。コストポイントを訪問し、ソリューションはエージェントの手の届かないところにあります。そして、別のバリデーター (ソリューションを読み取る唯一のコンポーネント) が、後でログと回答を照合します。 (分離の詳細については、以下の構築方法を参照してください。完全な仕組みについてはリポジトリを参照してください。)
言葉についてのメモ: チート耐性ではなく、チート耐性と呼びます。これは商用ゲームであるため、一部の事件が事前トレーニングに漏洩した可能性や、エージェントが回答では決して名前を挙げなかった潜在的な知識を使って探索を主導する可能性を排除できません。私が証明できるのは、エージェントの間違いは提供された情報のみと一致しているということです。証拠ではなく、強力な証拠です。
モデルのはしご (Claude Haiku 4.5 → Claude Sonnet 4.6 → Claude Opus 4.8 → Claude Fable 5) を越えると、2 つの失敗モードが繰り返し現れます。これらはボード ゲームに固有のものではないため、名前を付ける価値があります。LLM エージェントが複数ステップの検索と推論のタスクで失敗する方法です。
失敗 1 — 実行: 取得したものよりも生成したものを優先する
この事件の潜入捜査官はカバーネームを使用しています。全体的に最強のプレイヤーであるクロード・フェイブル 5 は、実際に提供された手がかりから本当の名前を見つけ、それを自分のメモに書きました。そして、応答時にそれを取り消し線で消し、自分で作成したより賢い名前、つまり「pa」のアナグラムに置き換えました。

ssenger-list の名前は、エレガントなものに「デコード」されたように見えました。
その前のページに正しい答えが見つかりました。推測の方がより賢いと感じられたため、生成された推測でそれをオーバーライドしました。これは、Claude Fable 5 の両方のクリーンなシングルパス ランで発生しました。チェックポイントが設定された実行は次のことを明らかにしています。3 回目の Claude Fable 5 の実行はゲーム中に再開する必要があり (レート制限)、そのため外部化されたメモのみを読み取る新しいエージェントとして再開されました。そして、取得した事実を再派生するのではなく額面どおりに受け取り、正しい名前を保持しました。これは、囮の罠から逃れ、エージェントに正確に名前を付けることができた唯一のクロード・フェイブル 5 の実行でした。そして、新たに生成された推測よりもメモ内の事実を信頼することによって、正確にそこに到達しました。
RAG や研究エージェントを構築している場合は、このバグをご存知でしょう。モデルが、取得したばかりのドキュメントに対して自信を持って幻覚を見せるというバグです。ここにそれがあり、孤立していて測定可能です。最新性と自己生成コンテンツへの偏見が、思い出された事実を上回ります。新しく生成された推論 (最近、私のもの) は、長い追加専用の歴史の中に埋もれていた提供された事実 (古い、他の人のもの) よりも優先されます。
失敗 2 — 理解: 明らかな容疑者はおとりである
これはまさにディナーの罠です。殺害された男は、表面的には明らかな「エージェント」です。元刑事で、ロンドンに来たばかりのアメリカ人です。あらゆる詳細から、彼がターゲットであると結論付けてしまいます。彼はそうではありません。本当の潜入捜査官は、殺人者が今も追いかけている生きている女性であり、その証拠は、テーブルで説明できなかった行動です。死体を狩るのではありません。
これを囮の罠と呼んでください。明らかな容疑者は代役であり、本当の答えは、あなたが推測しなければならない人物がまだそこにいるということです。それから逃げる — 行動として手がかりを読み取り、それに気づく

明白なストーリーを矛盾させ、明白なストーリーが間違っていると結論付けることは、私たちが果たせなかった二次的なターンです。そして、ほぼすべての構成がここに当てはまります。シングル エージェントの「方法論的探偵」プロンプトは、この 1 つの事件の 9 回のプレイスルーで実行され、9 回中 9 回おとりの罠にはまりました。
これら 2 つの失敗は、実行エラー (理解し、手探りした) と理解エラー (まったく理解できなかった) という他のすべてを整理します。
それぞれの失敗を実際に修正するもの (証拠)
私は介入のはしごを試し、それぞれが 1 つのレバーを分離しました。正直な要約: ほとんどのことは、簡単な失敗 (実行) とプロセスを助けます。難しいもの (理解) は、エージェント トポロジを変更するまで頑固でした。
一般的な「優れた研究者」のプロンプト（まず無料の資料をすべて使い尽くし、情報源を引用し、世界の 1 つのモデルを構築し、明白なことを信じない）。これによりプロセスが見事に整理され、探索が規律正しくなり、捏造 (アナグラム) がなくなりました。しかし、理解は動かなかった。おとりの罠は9/9を保持した。エージェントに事件の理解を教えなくても、優れた捜査官のように振る舞うことを教えることはできます。
（ハードモードの代わりに）事前に質問を知らせます。これは理解に関わる可能性がありますが、モデルに依存します。クロード・ソネット 4.6 はジャンプして罠を破った。クロード Opus 4.8 にはメリットがありませんでした。 Claude Haiku 4.5 では使用できませんでした。教訓: 罠は調査中に作られるので、調査中のヒントが役に立ちます。回答時だけ片付けるのでは遅すぎます。
どちらも理解の罠を確実に破ることはできませんでした。その結果、エージェントが 2 つに分割されました。
理解と探索を分離する
その動き: 調査と推論の両方を行う 1 人のエージェントの代わりに、協力する 2 人のエージェントを使用します。
理論家 — 理解者

シオンエンジン。世界へのアクセス権はありません。 grep も何もアクセスすることもできず、ディレクトリや GM の生の出力を見ることもありません。その唯一の仕事は、世界の単一モデルを維持し、すべての事実にその情報源をラベル付けし、未解決の部分を探し、独自の主要な仮説を改ざんしようとし、次に何を調査するかを決定することです。外部化された台帳から毎ターン新しく再生成されるため、行き止まりの汚染された記録が蓄積されることはありません。つまり、それが「維持」する世界のモデルは実際にはエージェント内にまったく存在しません。台帳テキストの中に存在し、毎ターンコンテキスト内で再構築されます。ケースモデルは理論家が書き直す文書であり、文書が保持する状態ではありません。 (これは、チェックポイント付きの Fable 実行が正しい名前を維持するのに役立ったのと同じ特性です。ワールド モデルが file である場合、新たな推測よりも書かれた事実を信頼する方が簡単です。) 終了することを決定するまで、質問は表示されません。
Explorer — 認識/アクション エンジン。ワークスペースとゲームマスターがあります。理論家から自由な意見を受け取り、名前→住所を解決し、訪問し、手がかりを逐語的に伝えます。この事件について結論を出すことは明確に禁じられています。
それらの間の指揮者は、純粋な文字通りのパイプです。
なぜこれが役立つのでしょうか?コンテキストがクリーンになったからではありません。以下のクリーンモノリス コントロールによってコンテキストがクリーンに保たれていますが、それでも落ちます。私の解釈では、Theorist は探索を決して行わない。手がかりを探すことで植え付けられる、明白な読み取り優先の枠組みを構築することは決してなく、独自の脚力で強化し続けてきたストーリーにコミットすることもない。探索の仕組みから目が見えなくなっているため、手がかりのひとつひとつが冷たく読み取られます。そして、ここが正確に言う価値のある部分です。これは、モノリスができなかった事実を結び付ける理論家ではありません。モノリスも同じでした

得られた手がかり、同じモデル、同じ記憶の新しい設定 — 散在する事実を関係につなぎ合わせるのは、両方ともできることです。異なるのは、ステッチングが実行される事前条件と、それを形作る定常的な秩序です。理論家の唯一の仕事は、有力な仮説を擁護することではなく、それを偽ることです。 2 番目の動きは、「A と B を結び付ける」ことではなく、「A が魅力的にした仮説を打ち消すために B を使用する」ことです。同じ手がかりをまだ探し続けているにもかかわらず、そこに到達したモノリスは依然としてそれを読み間違えていましたが、理論家は大声で叫びました。
その時までに、この事件に関するいくつかの事実が明らかになっていました - 殺害された男性と追われた女性は兄弟であり、殺人の背後にいる人々は名前を聞き出すために彼女を何時間も拷問していました - そして理論家はそれらをまとめました：
「彼らは身元を聞き出すために妹を何時間も拷問した。もし死んだ兄が潜入捜査官だったなら、彼らはすでに兄を捕まえているだろう――彼女から名前を引きずり出す必要はない。殺人者はまだ公開命令に基づいて行動している。したがって、捜査官は生きており、死んだ男ではない。」
これは、ディレクトリに触れたことのないエージェントによって平文で行われた二次推論です。
「しかし、それは本当に建築なのでしょうか？」 — 私自身の結論を問う
ここで、記事が説いていることを実践する必要があります。 「理解はトポロジーの問題である」というのは大きな主張であり、優れた探偵の仕事とは、この記事の主題全体ですが、結論が出るまで明白な結論を信頼しないことを意味します。

[切り捨てられた]

## Original Extract

A Sherlock Holmes board game as an LLM-agent eval

🔎 ">
🔎 sherlock-agent-eval
View on GitHub
How good a detective is an AI?
A Sherlock Holmes board game as an LLM-agent eval
It started at a dinner. A few friends and I sat down to play Sherlock Holmes Consulting Detective — an open-ended deduction game where you’re handed a Victorian London case, you pick which people and places to go investigate, and each lead hands you a passage of text to read. Most of the game is reading, cross-referencing, and arguing at the table. At the end you answer the case’s questions and score yourself against Holmes himself — including how few leads you needed. The answers sit in the back of the booklet, printed upside-down, daring you not to peek.
We walked straight into the trap the case is built around. There’s an obvious victim — a man every detail points to as the target — and we hung our whole theory on him. But one clue wouldn’t sit still. The morning after the murder, the killer goes back to a shipping office and scans the passenger list again. We re-read the passage three times. Why would he do that? If he’d already killed the person he was after, what was he still looking for? Something didn’t close, and none of us could say what.
So, at 2am — out of wine and out of steam — we did the forbidden thing: we turned the booklet over. And there in the answer key, a name we’d treated as background all evening stepped forward as the real undercover agent — alive, never caught, the person the killer was still hunting. The passenger-list visit wasn’t a loose end. It was the case. We’d held the contradiction in our hands — we’d even said out loud that it was strange — and we’d read right past it.
That non-closing feeling is the thing that stuck. We weren’t short on information; we had every clue we needed. We were short one inference — the small, second-order turn from “that’s a strange thing for the killer to do” to “then the whole story we’ve built is wrong.” So I started to wonder: how good a detective is an AI, really? Handed the same leads, would an LLM agent read that behavior as a behavior, notice it broke the obvious story, and follow it to the live agent we’d missed?
To find out, I turned the game into an eval for LLM agents. The agent plays the Irregulars — the Baker Street street kids Holmes sends out to do his legwork.
On its first run, Claude Fable 5 tied Holmes — in the hard mode, where you don’t even get to see the questions until the investigation is over.
That’s the headline. But the score isn’t the story. The interesting part is the two distinct ways these agents fail — and that the harder failure, the exact one that beat us at dinner, has a clean fix that turned out to be less about model size than I expected.
Why a board game is a surprisingly honest agent eval
What I didn’t see at the table that night is that we’d just lost to an unusually clean agent benchmark. Most agent benchmarks have a problem: the answer is somewhere in the context, or the task is gameable, or “success” is graded loosely. A printed detective game sidesteps all three by construction:
The solution is physically hidden. Those upside-down answers never enter the agent’s allowed workspace; reading them would be a detectable protocol violation, and I audit for it.
Information has a price. Thinking, re-reading, and cross-referencing are free and unlimited. But acting — visiting a location to pull a new clue — is the only way to get new information, and every new clue beyond what Holmes used costs points . That’s a miniature of real agent economics: every tool call costs something.
It rewards comprehension, not retrieval. Clues are behaviors and details you have to assemble into one coherent story; none of them hands you the answer.
The mechanics that make this auditable, in one breath: the agent works in a sandbox containing only what it’s allowed to see; a deterministic Game Master (plain Python, not an LLM) serves clues verbatim and logs everything; visits cost points and the solution lives outside the agent’s reach; and a separate validator — the only component that reads the solution — cross-checks the log against the answers afterward. (More on the isolation in How it’s built below; full mechanics in the repo .)
A note on words: I’ll call it cheat-resistant , not cheat-proof. It’s a commercial game, so I can’t rule out that some of the case leaked into pretraining, or that an agent could steer its exploration with latent knowledge it never names in an answer. What I can show is that the agents’ mistakes are consistent with only the information they were served — strong evidence, not proof.
Across a ladder of models (Claude Haiku 4.5 → Claude Sonnet 4.6 → Claude Opus 4.8 → Claude Fable 5), two failure modes show up again and again. They’re worth naming because they’re not specific to board games — they’re how LLM agents fail at any multi-step retrieval-and-reasoning task.
Failure 1 — Execution: preferring what you generated to what you retrieved
The case’s undercover agent uses a cover name. Claude Fable 5 — the strongest player overall — actually found the real name in a served clue and wrote it into its own notes. Then, at answer time, it crossed it out and replaced it with a cleverer name it had constructed itself: an anagram of a passenger-list name that looked like it “decoded” into something elegant.
It had the right answer, retrieved, on the page in front of it. It overrode it with a guess it generated, because the guess felt more clever. This happened in both of Claude Fable 5’s clean single-pass runs. The checkpointed run is revealing: a third Claude Fable 5 run had to be restarted mid-game (rate-limiting), so it resumed as a fresh agent reading only its externalized notes — and, taking that retrieved fact at face value instead of re-deriving it, it kept the correct name. It was the only Claude Fable 5 run to both escape the decoy trap and name the agent correctly — and it got there precisely by trusting a fact in its notes over a freshly-generated guess.
If you build RAG or research agents, you know this bug — the one where the model confidently hallucinates over a document it just retrieved. Here it is, isolated and measurable: recency plus a bias toward self-generated content beats recalled fact. The freshly-generated inference (recent, mine ) wins over the served fact (old, someone else’s) buried in a long append-only history.
Failure 2 — Comprehension: the obvious suspect is a decoy
This is the trap from the dinner, named precisely. The murdered man is, on the surface, the obvious “agent” — a former detective, an American just arrived in London. Every detail invites you to conclude he’s the target. He isn’t: the real undercover agent is the living woman the killer is still hunting , and the tell is the behavior we couldn’t explain at the table — you don’t hunt a corpse.
Call this the decoy trap : the obvious suspect is a stand-in, and the real answer is the one you have to infer is still out there. Escaping it — reading a clue as a behavior, noticing it contradicts the obvious story, concluding the obvious story is wrong — is the second-order turn we failed to make. And it’s where almost every configuration falls down: a single-agent “methodical detective” prompt, run across nine playthroughs of this one case, fell for the decoy trap 9 times out of 9.
These two failures organize everything else: execution errors (you understood it and fumbled it) versus comprehension errors (you never understood it).
What actually fixes each failure (the evidence)
I tried a ladder of interventions, each isolating one lever. The honest summary: most things help the easy failure (execution) and the process ; the hard one (comprehension) was stubborn until I changed the agent topology.
A generic “good investigator” prompt (exhaust the free material first, cite your sources, build one model of the world, distrust the obvious). This cleaned up the process beautifully — exploration got disciplined, fabrication (the anagram) disappeared. But comprehension didn’t move: the decoy trap held 9/9. You can teach an agent to behave like a good investigator without teaching it to understand the case.
Letting it know the questions up front (instead of the hard mode). This can touch comprehension — but it’s model-dependent. Claude Sonnet 4.6 jumped and cracked the trap; Claude Opus 4.8 didn’t benefit; Claude Haiku 4.5 couldn’t use it. The lesson: the trap is built during the investigation, so a hint during the investigation helps; cleaning up only at answer time is too late.
Neither reliably broke the comprehension trap. The thing that did was splitting the agent in two.
Split comprehension from exploration
The move: instead of one agent that both explores and reasons, use two agents that cooperate :
A Theorist — the comprehension engine. It has no access to the world : it can’t grep , can’t visit anything, never sees the directory or the GM’s raw output. Its only job is to maintain a single model of the world, label every fact with its source, hunt loose ends, try to falsify its own leading hypothesis, and decide what to investigate next. It’s re-spawned fresh every turn from an externalized ledger, so it never accumulates a contaminated transcript of dead-ends — which means the model of the world it “maintains” doesn’t really live in the agent at all: it lives in the ledger text and is rebuilt, in-context, every turn. The case-model is a document the Theorist rewrites, not a state it holds. (That’s the same property that helped the checkpointed Fable run keep the right name: when your world-model is a file , it’s easier to trust the written fact over a fresh guess.) It doesn’t see the questions until it decides to close.
An Explorer — the perception/action engine. It has the workspace and the Game Master. It takes a loose-end from the Theorist, resolves the name→address, visits, and relays the clue back verbatim . It’s explicitly forbidden from concluding anything about the case.
A conductor between them is a pure verbatim pipe.
Why would this help? Not because the context is cleaner — the clean-monolith control below keeps it clean and still falls. My read is that the Theorist never does the exploration: it never builds the obvious-reading-first frame that hunting for clues instills, and it isn’t committed to a story its own legwork kept reinforcing. Blinded from the mechanics of exploring, it reads each clue cold. And here’s the part worth being precise about: this isn’t the Theorist connecting facts the monolith couldn’t. The monolith had the same served clues, the same model, and the same fresh-memory setup — stitching scattered facts into a relation is something both can do. What differs is the prior that stitching runs under, plus the standing order that shapes it: the Theorist’s one job is to falsify its leading hypothesis, not defend it. The second-order move isn’t “connect A and B” — it’s “use B to kill the hypothesis that A made tempting.” Given the same still-hunting clue, the monoliths that reached it still misread it — but the Theorist made the call out loud.
By then a few facts about the case had surfaced — the murdered man and the hunted woman were siblings, and the people behind the murders had tortured her for hours trying to get a name — and the Theorist put them together:
“They tortured the sister for hours to extract an identity. If the dead brother were the infiltrated agent, they’d already have him — they wouldn’t need to drag a name out of her. The killer is still acting on an open order. Therefore the agent is alive, and it isn’t the dead man.”
That’s the second-order inference, made in plain text, by an agent that never touched a directory.
“But is it really the architecture?” — interrogating my own conclusion
This is where the article has to practice what it preaches. “Comprehension is a topology problem” is a big claim, and good detective work — the entire subject of this article — means distrusting your obvious conclusion until you’ve ruled o

[truncated]
