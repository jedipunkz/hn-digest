---
source: "https://manazir.dev/work/anamnesis-forecasting-memory-corruption"
hn_url: "https://news.ycombinator.com/item?id=49164643"
title: "LLM memory doesn't only get written wrong, it goes wrong later"
article_title: "ANAMNESIS: Forecasting and Self-Healing of Memory Corruption in LLM Agents"
author: "mnzrdev"
captured_at: "2026-08-04T06:25:27Z"
capture_tool: "hn-digest"
hn_id: 49164643
score: 1
comments: 0
posted_at: "2026-08-04T05:30:59Z"
tags:
  - hacker-news
  - translated
---

# LLM memory doesn't only get written wrong, it goes wrong later

- HN: [49164643](https://news.ycombinator.com/item?id=49164643)
- Source: [manazir.dev](https://manazir.dev/work/anamnesis-forecasting-memory-corruption)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T05:30:59Z

## Translation

タイトル: LLM メモリは間違って書き込まれるだけでなく、後で間違ってしまう
記事のタイトル: ANAMNESIS: LLM エージェントにおけるメモリ破損の予測と自己修復
説明: 最終年度のプロジェクトの生きた記録。 LLM の記憶は、間違って書かれているだけではなく、間違った方向に進みます。月曜日には真実だった事実が、木曜日には矛盾されたり、置き換えられたり、毒されたりし、次の統合は腐敗の上に築かれます。すべての出荷メモリ システムは、書き込み時にファクトを 1 回チェックします。アナムネシス
[切り捨てられた]

記事本文:
ANAMNESIS: LLM エージェントのメモリ破損の予測と自己修復 コンテンツへスキップ スリランカ、コロンボ 仕事について ブログ ギャラリー © 2026 · Manazir Ali サイトマップ RSS プロジェクト ANAMNESIS: LLM エージェントのメモリ破損の予測と自己修復
これは完成した論文ではありません。これは 2027 年 4 月までの最終年度のプロジェクトの作業記録であり、最後に一度書くのではなく、作業の進行に合わせて更新します。セクション 8 は変更ログであり、以下のステータス表は、今日の実際の内容を確認する最も簡単な方法です。
2 つの分野が論文からこのページに引き継がれています。測定されるまでは結果として何も主張されず、まだ保留中のものはすべてそのことを明白な言葉で示しています。この研究で私が間違っていることがすでに証明されている箇所については、読む価値のある部分なので、それもページに掲載されています。
記憶システムはあなたの会話を読み取り、後で思い出すことができる短い事実に抽出します。この抽出ステップは言語モデルによって行われ、言語モデルはあなたが決して語らなかったことを書き留めることができます。これくらいはよく知られています。
より困難な問題はその後にやってくる問題です。事実は完全に記録されても、真実でなくなる可能性があります。あなたが街を動かすのです。あなたは仕事を変えます。あなたはラップトップを売ります。何も幻覚は見られず、何も攻撃されませんでした。世界はただ動いただけで、記憶は動きませんでした。
どこかで読む前に、私は自分のシステムでこれをヒットしました。私はアシスタントに、MacBook に切り替えたことを伝えました。統合社は私の次の質問の13秒前に新事実を認めた。検索により、内容が文脈に引き込まれます。答えはやはり ThinkPad でした。
Enter またはスペースを押してノードを選択します。その後、矢印キーを使用してノードを移動できます。削除するには削除を押し、キャンセルするにはエスケープを押します。 Enter または Space を押してエッジを選択します。その後、削除を押して、

削除するかエスケープしてキャンセルしてください。通常の意味では、そのシーケンス内の何もバグではありません。すべてのコンポーネントがその役割を果たしました。ギャップは、古い事実が過ぎ去ったばかりであることに気付くコンポーネントが存在しないということであり、それがこのプロジェクトのギャップです。
言語モデルとの長時間にわたる会話のためのローカルファーストの記憶システム。記憶は 4 つの層で構成されており、認知神経科学が人間の長期記憶をどのように分解するかに大まかにマッピングされています。ライブ ターンの作業層、追加のみのエピソード記録、帯域外の統合によって構築された抽出された意味層、および応答パターンの手続き層です。ストア全体は 1 つの SQLite ファイルです。書き込みパス内の唯一のネットワーク呼び出しは、チャット モデル自体に対するものです。取得、含意チェック、リスク スコアリングはすべて CPU 上でローカルに実行されます。
すべての候補事実は、記憶になる前に同じ短い経路をたどります。
Enter またはスペースを押してノードを選択します。その後、矢印キーを使用してノードを移動できます。削除するには削除を押し、キャンセルするにはエスケープを押します。 Enter または Space を押してエッジを選択します。その後、削除を押して削除するか、エスケープしてキャンセルできます。負荷に耐える特性は、何も検査されずにセマンティック メモリに入ることがなく、何も削除されないことです。拒否された事実は監査記録がそのままの状態で隔離されるため、拒否を黙って飲み込むのではなく検査することができます。
3. 明白なバージョンが貢献ではない理由
これは、他の人がこのページを書いた場合、最初に読みたいセクションです。
私は 2026 年 5 月に、記憶の統合が忠実であるかどうかを誰も測定したりゲートしたりしていないという確信を持った主張から始めました。翌日の文献スキャンにより、それが解体されました。脳にインスピレーションを得た階層型記憶は何度も構築されてきました。永続的な単一会話メモリが出荷されました。ベンチマー

メモリシステムの幻覚専用に作られた HaluMem は 2025 年 11 月に公開されていました。書き込み時ゲートが指定されていました。
そこで主張は変わりました。誰もこれらの部分を統合して、それらが相互に役立つかどうかをテストしたことがありませんでした。それは約3週間生き延びました。 6月に私は2024年から2026年の分野について敵対的監査を実施し、特に私自身の絞り込まれた主張を潰そうとしましたが、それはうまくいきました。 ProMem はすでに書き込み時検証を構築し、HaluMem でベンチマークを実行し、それによって Mem0 を破っていました。他のいくつかのシステムでは、ゲート制御またはトレーニングされたメモリ操作が同じウィンドウ内で出荷されていました。
3 回目のパスで生き残ったのは、監査でどこにも見つからなかった部分です。これらのシステムはどれも現在形でメモリをチェックします。これは現在サポートされていますか。現在クリーンなメモリが腐る可能性を予測するものはなく、その予測に基づいて修復ループを閉じるものはありません。
それがこのプロジェクトの本質であり、この狭さは意図的なものです。
Enter またはスペースを押してノードを選択します。その後、矢印キーを使用してノードを移動できます。削除するには削除を押し、キャンセルするにはエスケープを押します。 Enter または Space を押してエッジを選択します。その後、削除を押して削除するか、エスケープしてキャンセルできます。 3 つのメカニズムとその正直な新規性の範囲:
GATE は候補事実を原子的な主張に分解し、ローカル含意モデルを使用して、すべての主張が引用するエピソードの結合によって含意される場合にのみそれを認めます。それは目新しいものではありませんし、私はそれを主張しません。ライブ実行からの有効な例: 統合は、ウィスカーという名前の猫が 2021 年 3 月に生まれたと提案しましたが、情報源はその推論については述べられていませんでした。分解によりアサーションが分離され、含意スコアは 0.055 に戻り、ファクトは拒否され、失敗したクレームが指定されて隔離されました。
汚職の予測には、ラベル付きの例が必要です

後に破損した、タイミングがわかった記憶。そのようなデータセットは公開されていないため、ラベルは情報を含むデータセットから暗黙的に無料で派生されます。
3 つのルールにより評価が公正に保たれます。すべてのデータセットには独自のゴールド メモリが付属しており、テスト対象となるのはメモリの構築であるため、どれもロードされることはありません。生の会話が再生され、システムが独自のメモリを構築します。ラベル自体が将来の情報であり、ランダムな分割では情報が漏洩してしまうため、分割はランダムではなく一時的なものです。また、すべての統計手順は実行前に修正されているため、分析が望ましい答えに向かってドリフトすることはありません。
評価全体は、100 ドルのハードバジェットの下、1 台のラップトップ上で実行され、書き込みパスは設計上自由です。
スイートには 220 のバックエンド テストと 222 のフロントエンド テストがあります。無料枠モデルのみで、記録された 211 件のモデル コール全体で、これまでの支出はゼロドルです。
6. 最初の数字が実際に何を示しているか
予報官は、最初から最後まで一度トレーニングを受けています。ラップトップの CPU で 3 秒、お金はかかりません。これは科学的な結果というよりも配管の検証であり、数値よりも正直に読み取ることが重要です。
厳密に後のテスト分割でのフォワード AUROC は 0.665 でした。これは、後で破損したメモリと破損しなかったメモリが 1 つあるとすると、モデルはそれらを 3 回中 2 回正しくランク付けすることを意味します。偶然よりは良いですが、役に立つには程遠いです。同一の特徴でトレーニングされたロジスティック回帰のスコアは 0.6647 でした。したがって、勾配ブースト モデルは最も単純な妥当なベースラインを 0.0007 上回っており、これは同点です。私は結果が出る前に、その結​​果を正確に報告することを書面で約束していました。
審査官に指摘してもらいたい 3 つの注意点:
3 番目のことは、このプロジェクトが私に教えてくれた最も有益なことです

これまでのところ、トレーニングの実行を信頼するのではなく、トレーニングされたモデルの機能をライブデータベーススキーマと照合して確認することでそれを見つけました。
7. 持ち帰る価値のある 3 つのもの
これらはこのプロジェクトを超えて一般化されているため、論文のみではなくここに記載されています。
明らかな特徴が間違っている可能性があります。異常検出器を構築する場合、本能的に埋め込み距離に手を伸ばし、外れ値を探すことになります。 MINJA の論文は、通常の会話によって植え付けられた毒された記憶が、埋め込み空間内で良性の記憶と絡み合って存在することを示しています。これは、攻撃が溶け込むことによって機能するためです。そのジオメトリに基づいて構築された検出器は、最もキャッチする必要がある攻撃を正確に認識できないため、ここで設定された機能は動作的かつ構造的なものであり、埋め込み距離はまったく含まれていません。同じ表現は検索に負荷がかかり、予測から意図的に除外されます。
要約統計は、縮退モデルによって操作できます。上記の校正数値は素晴らしく、意味がありませんでした。どのスカラーでも、1 つのことを言うことを学習したモデルを隠すことができます。その背後にある図ではそれができません。
データをテーブルにダンプして確認することは、コードを読むよりも簡単です。このプロジェクトの 2 つの実際のバグは、ソースを読むことによってではなく、その方法で発見されました。1 つは上記のショートカット学習の問題、もう 1 つは 14,823 行ごとに 3 つの機能を黙ってゼロにする文字列の不一致です。
ANAMNESIS は、ウェストミンスター大学付属情報学研究所での 2026 年から 2027 学年度に向けた私の最終年度プロジェクトであり、Shiham Farook が監督しています。これは、主権パーソナル AI に関する私の長期にわたる個人研究である Project Hydra と 1 つの研究領域を共有していますが、独自の範囲、成果物、期限があります。このスライスに名前を付けた Hydra シリーズの以前の記事は、

このアイデアは、セクション 3 で説明する再構築前の 2026 年 5 月に始まりました。
設計の背後にある厳選された情報源: Chen et al. (2025) HaluMem で。ドンら。 (2025) メモリインジェクション攻撃について。チャオら。 (2026) 暗黙の古さについて。ヤンら。 (2026) プロアクティブなメモリ抽出について。ミンら。 (2023) 原子的な事実の正確さについて。郭ら。 (2017) 校正について。アンゲロプロスとベイツ (2021) による等角予測。

## Original Extract

A living record of my final-year project. LLM memory is not only written wrong, it goes wrong: a fact that is true on Monday is contradicted, superseded, or poisoned by Thursday, and the next consolidation builds on the rot. Every shipping memory system checks a fact once, at write time. ANAMNESIS a
[truncated]

ANAMNESIS: Forecasting and Self-Healing of Memory Corruption in LLM Agents Skip to content Colombo, Sri Lanka About Work Blog Gallery © 2026 · Manazir Ali Sitemap RSS Projects ANAMNESIS: Forecasting and Self-Healing of Memory Corruption in LLM Agents
This is not a finished paper. It is the working record of a final-year project that runs to April 2027, and I update it as the work moves rather than writing it once at the end. Section 8 is the changelog, and the status table below is the fastest way to see what is actually true today.
Two disciplines carry over from the thesis into this page. Nothing is claimed as a result until it has been measured, and anything still pending says so in plain words. Where the work has already proved me wrong, that is on the page too, because those are the parts worth reading.
A memory system reads your conversation and distils it into short facts it can recall later. That distilling step is done by a language model, and a language model can write down something you never said. This much is well known.
The harder problem is the one that arrives afterwards. A fact can be recorded perfectly and still stop being true. You move city. You change jobs. You sell the laptop. Nothing was hallucinated and nothing was attacked; the world simply moved, and the memory did not.
I hit this in my own system before I read about it anywhere. I told the assistant I had switched to a MacBook. Consolidation admitted the new fact thirteen seconds before my next question. Retrieval pulled it into context. The answer still said ThinkPad.
Press enter or space to select a node. You can then use the arrow keys to move the node around. Press delete to remove it and escape to cancel. Press enter or space to select an edge. You can then press delete to remove it or escape to cancel. Nothing in that sequence is a bug in the ordinary sense. Every component did its job. The gap is that no component's job was to notice that an old fact had just been outlived, and that is the gap this project sits in.
A local-first memory system for long-running conversations with a language model. Memory is organised in four tiers, loosely mapped to how cognitive neuroscience decomposes human long-term memory: a working tier for the live turn, an append-only episodic record, a distilled semantic tier built by out-of-band consolidation, and a procedural tier for response patterns. The whole store is a single SQLite file. The only network call in the write path is to the chat model itself; retrieval, entailment checking, and risk scoring all run locally on CPU.
Every candidate fact travels the same short path before it becomes memory.
Press enter or space to select a node. You can then use the arrow keys to move the node around. Press delete to remove it and escape to cancel. Press enter or space to select an edge. You can then press delete to remove it or escape to cancel. The load-bearing property is that nothing enters semantic memory unexamined, and nothing is ever deleted. A rejected fact is quarantined with its audit record intact, so a refusal can be inspected rather than silently swallowed.
3. Why the obvious version is not the contribution
This is the section I would want to read first if someone else had written this page.
I started in May 2026 with a claim I was confident about: that nobody was measuring or gating whether memory consolidation is faithful. A literature scan the following day dismantled it. Brain-inspired tiered memory had been built several times over. Persistent single-conversation memory had shipped. A benchmark purpose-built for hallucination in memory systems, HaluMem, had been published in November 2025. Write-time gates had been specified.
So the claim moved: nobody had integrated these pieces and tested whether they help each other. That survived about three weeks. In June I ran an adversarial audit of the 2024 to 2026 field, specifically trying to kill my own narrowed claim, and it worked. ProMem had already built write-time verification, benchmarked it on HaluMem, and beaten Mem0 doing it. Several other systems had shipped gated or trained memory operations in the same window.
What survived the third pass is the part the audit could not find anywhere. Every one of those systems checks a memory in the present tense: is this supported right now. None of them forecast which currently-clean memory is going to rot, and none close a repair loop driven by that forecast.
That is what the project is, and the narrowness is deliberate.
Press enter or space to select a node. You can then use the arrow keys to move the node around. Press delete to remove it and escape to cancel. Press enter or space to select an edge. You can then press delete to remove it or escape to cancel. The three mechanisms and their honest novelty scoping:
GATE decomposes a candidate fact into atomic claims and admits it only if every claim is entailed by the union of the episodes it cites, using a local entailment model. It is not novel and I do not claim it. The worked example from a live run: consolidation proposed that a cat named Whiskers was born in March 2021, an inference the source turns never stated. Decomposition isolated the assertion, the entailment score came back at 0.055, and the fact was rejected and quarantined with the failing claim named.
Forecasting corruption needs labelled examples of memories that later became corrupted, with the timing known. No such dataset is published, so the labels are derived from datasets that contain the information implicitly, at no cost.
Three rules keep the evaluation honest. Every dataset ships its own gold memories and none of them are ever loaded, because memory construction is the thing under test; the raw conversations are replayed and the system builds its own memory. The splits are temporal rather than random, because the label is itself future information and a random split would leak it. And every statistical procedure was fixed before the runs, so the analysis cannot drift toward a preferred answer.
The whole evaluation runs under a hard budget of one hundred US dollars, on one laptop, with the write path free by design.
Suites stand at 220 backend and 222 frontend tests. Spend to date is zero dollars across 211 recorded model calls, on free-tier models only.
6. What the first numbers actually say
The forecaster has been trained once, end to end. Three seconds on a laptop CPU, no money spent. It is a plumbing verification rather than a scientific result, and the honest reading matters more than the number.
Forward AUROC on the strictly later test split came out at 0.665, meaning that given one memory that later corrupted and one that did not, the model ranks them correctly about two times in three. Better than chance, well short of useful. A logistic regression trained on identical features scored 0.6647, so the gradient-boosted model beat the simplest reasonable baseline by 0.0007, which is a tie. I had committed in writing to reporting exactly that outcome before the result existed.
Three caveats, all of which I would want an examiner to raise:
That third one is the most useful thing the project has taught me so far, and I found it by checking the trained model's features against the live database schema rather than trusting the training run.
7. Three things worth taking away
These generalise past this project, which is why they are here rather than only in the thesis.
The obvious feature can be the wrong one. Building an anomaly detector, the instinct is to reach for embedding distance and look for outliers. The MINJA paper shows that poisoned memories planted through ordinary conversation sit entangled with benign ones in embedding space, because the attack works by blending in. A detector built on that geometry would be blind to exactly the attack it most needs to catch, so the feature set here is behavioural and structural, and contains no embedding distance at all. The same representation is load-bearing for retrieval and deliberately excluded from forecasting.
A summary statistic can be gamed by a degenerate model. The calibration number above was excellent and meaningless. Any scalar can hide a model that has learned to say one thing; the diagram behind it cannot.
Dumping the data into a table and looking at it beats reading the code. Two real bugs in this project were found that way, not by reading source: the shortcut-learning problem above, and a string mismatch that silently zeroed three features across every one of 14,823 rows.
ANAMNESIS is my final-year project at the Informatics Institute of Technology, affiliated with the University of Westminster, for the 2026 to 2027 academic year, supervised by Shiham Farook. It shares one research territory with Project Hydra, my longer-running personal work on sovereign personal AI, but has its own scope, deliverables, and deadline. The earlier article in the Hydra series that named this slice is a snapshot of where the idea started in May 2026, before the reshaping described in section 3.
Selected sources behind the design: Chen et al. (2025) on HaluMem; Dong et al. (2025) on memory injection attacks; Chao et al. (2026) on implicit staleness; Yang et al. (2026) on proactive memory extraction; Min et al. (2023) on atomic factual precision; Guo et al. (2017) on calibration; Angelopoulos and Bates (2021) on conformal prediction.
