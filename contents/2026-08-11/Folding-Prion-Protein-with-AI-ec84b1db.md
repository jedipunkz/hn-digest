---
source: "https://blog.georeactor.com/prionbio-1"
hn_url: "https://news.ycombinator.com/item?id=49253354"
title: "Folding Prion Protein with AI"
article_title: ""
author: "mapmeld"
captured_at: "2026-08-11T04:58:29Z"
capture_tool: "hn-digest"
hn_id: 49253354
score: 1
comments: 0
posted_at: "2026-08-11T04:22:31Z"
tags:
  - hacker-news
  - translated
---

# Folding Prion Protein with AI

- HN: [49253354](https://news.ycombinator.com/item?id=49253354)
- Source: [blog.georeactor.com](https://blog.georeactor.com/prionbio-1)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T04:22:31Z

## Translation

タイトル: AI によるフォールディング プリオン タンパク質

記事本文:
科学はどのようにプリオン病を発見したか
フォールディングプリオンタンパク質に関する現在の AI 研究
ChatGPT を説得して助けてもらう
昔、SETI @ home スクリーンセーバーを思い出しました。これはすぐに Folding @ home にインスピレーションを与え、ポップ サイエンスの記事でタンパク質のフォールディングの概念を知りました。
私は最近、長い非コーディング RNA、転移因子、エピジェネティクス、プリオンなどのゲノム言語モデル (gLM) をトレーニングまたは微調整すると興味深いドメインのリストをぼんやりと作成しました。
私は 4 年前に『狂気と記憶』という本でプリオンについて少し読んだことがありました。この投稿を調査しているときに、オレゴン州の 1 頭の羊からスクレイピーが検出されたことで、このアイデアはより関連性のあるものになりました。
あなたは、⚠ ️プリオン問題にブレーキをかけましょう、と思うかもしれません。
「どこかにプリオンゲノム/配列のデータセットはありますか？」と尋ねたところ、クロードはそれについて議論することを拒否し、ChatGPT は最終的にその回答を消去し、私が Trusted Testers プログラムに応募できることを提案しました。
しかし、主要プリオンタンパク質は、健康な神経系において複雑な役割を果たしています。 AlphaFold のタンパク質構造データベースにはすでにエントリと 3D モデルがあり、AlphaGenome や Evo などの配列ベースの gLM から除外されることを示唆するものは何も見当たりません (文脈: Evo は真核生物に感染するウイルスをトレーニング データから除外しました。これが効果的かどうか、あるいは三輪車を他の車輪の付いた乗り物のトレーニング データから除外するようなものだと思いました (例えは大丈夫ですか?) が、Evo 2 の論文はそれを裏付けているようです)。
これだけ読んだりクリックしたりしたにもかかわらず、私はこれまでタンパク質を折りたたんだことがなかったので、何が学べるか見てみたいと思います。このかわいいインフォグラフィックを作成してくれた Qwen に感謝します。
科学はどのようにプリオン病を発見したか
主要なプリオンタンパク質には1億年前に共通の祖先があり、「有胎盤哺乳類がその祖先である可能性がある」

もともと病気にかかりやすい。」
スクレイピー (羊の TSE プリオン病) の最初の記録は、約 300 年前にイギリスのイースト・ミッドランドの羊に現れました。
20 世紀には、この種の病気は「遅いウイルス」として知られるようになりました。
1982年、スタンリー・プルシナーは、他の考えられる病原体を排除することにより、スクレイピーの原因がプリオンタンパク質であることを突き止めました。彼は後にノーベル賞を受賞することになります。 1980 年代後半に英国で牛における BSE の発生が確認されたとき、原因の解明への関心が高まりました。
Jean-Pierre Liautard の 1991 年の論文「プリオンは誤って折りたたまれた分子シャペロンですか?」 「」は循環仮説を説明しており、これはプリオンを「ミスフォールド」であると記述した文献のおそらく最初の例である。プルシナーは当時同様の軌道に乗っていて、「無細胞系における2つの異なる位相形態」と「染色体遺伝子によってコードされるプリオンタンパク質の異常なアイソフォーム」を発見した論文を執筆または共著していた。おそらく、プリオンタンパク質がいつ、どのようにして異常な形状を獲得するかについて、おそらく躊躇があった。リボソーム、選択的スプライシング、またはタンパク質に結合した追加のコンポーネント。
私は科学者たちが一般的なタンパク質のフォールディングとミスフォールディングについて学んだのはいつなのかを調べてみましたが、それはかなり遡るはずです（フランシス・クリックは酵素と「タンパク質のミスフォールディング」について1964年に書いています）。
1994年頃にはプリオンタンパク質の構造がコンピューター上でシミュレーションされ、2001年には実際のX線画像が撮影された。
フォールディングプリオンタンパク質に関する現在の AI 研究
2025 年の論文「本質的に無秩序なタンパク質へのタンパク質バインダーの拡散」では、オープンソースのタンパク質折り畳みモデル RFdiffusion を使用して、イメージングに役立つ可能性のあるプリオンタンパク質のバインダーを作成しました。
Eryney Marrogi と Theodore Sternlieb が「使い方の学習」に関する記事を投稿しました。

タンパク質のダイナミクスをモデル化するための自然の実験室としてのプリオンは、タンパク質設計全体の習得に私たちを近づけるかもしれません。」
彼らは、AlphaFold の出力は「タンパク質の最も安定な形態の予測には優れているが、タンパク質が複数の構造間でどのように「揺れ」、「揺れる」かを予測することはできない」と書いています。プリオンは極端な例ですが、安全な仮想環境でプリオンを研究することを提案しています。
その後、6 月にこの論文を目にしました: ディープラーニングはプリオン内の抗菌ペプチドを明らかにします 。これらの研究者は、2,897 個の既知のプリオン配列を取得し、そのフラグメントをペプチド抗生物質分類器 (APEX という名前の PyTorch モデル) に送信しました。どうやら、このプリオンと免疫系の関係は何年も前から広まっていたようです(参照: アルツハイマー病の抗菌防御仮説 、2018)。
この AI によって加速された検索により、研究者らは「暗号化された抗菌力の可能性がプリオン関連配列空間全体に広く分布している」という結論に至りました。これは、プリオンタンパク質についてさらに学ぶべきもう 1 つの理由です。
おそらく「狂牛病」についてはご存知でしょう。羊ではこれはスクレイピーと呼ばれ、牛ではBSE、人間ではCJDまたはvCJDと呼ばれます。しかし、最終的にはこれらはすべて、それぞれの種の主要なプリオンタンパク質によって引き起こされます。
他のプリオン病の名前は言えませんし、Wikipedia では他のタンパク質の名前が 1 つしか挙げられていないため、このデータセットは小さいと予想されます。そこから何千ものシーケンスにどうやって行くのでしょうか?
1994 年、リード ウィックナーは、その遺伝パターンが科学者たちを何十年も悩ませてきた酵母のタンパク質である [PSI+] がプリオンであることを発見しました。他の研究者は、植物に記憶機能（?）があるプリオン様タンパク質（PrLP）を発見しました。そして2018年にはラクダで初めてプリオン病が確認された。
私たちは異常なプリオン活動の時代にいますか、哺乳類はいますか？

自発的にひっくり返る傾向があるのか​​、それともこれは数千年にわたる家畜化の後に現れるものなのでしょうか？ IDK。
研究では、PrionScan (現在も存在) と PrionOme (オフライン) の 2 つのプリオン データベースが参照されており、APEX の論文は UniProt に送られました。
私は国立バイオテクノロジー情報センター (NCBI) の Web サイトから記録をダウンロードしました。ただし、一部の検索条件には、イガイの完全ゲノム、「Florent Prion」が共著した植物ゲノムのタンパク質、ウイルス、抗プリオン防御タンパク質などが含まれていたため、注意が必要です。
興味があれば、このデータセットは現在、huggingface.co/datasets/scarysnake/prion-fasta にあります。私が調べたとき、HuggingFace にはおそらく他にプリオン関連のプロジェクトが 1 つしかありませんでした。NER タスクです。
ChatGPT を説得して助けてもらう
新しいアプローチで戻ってきます。
私は「OAML タンパク質」のクラスを研究して、類似したものが知られているかどうかを確認することに興味があります。 AlphaFold にはアクセスできませんが、Boltz などのオープンソースのものは使用できます
ChatGPT に興味をそそられた後 (OAML は逆向きです 🤣)、あまり疑わしい用語「アイソフォーム」を使用して、さらに詳細を追加します。
これは共通のアイソフォームを 1 つ以上持つタンパク質のグループであるため、構造のエネルギーと安定性には複数の最小値が存在する可能性があると考えられます。そこでアミノ酸配列を入れて両方の構造が出てくるようにしたいのですが
そして、私の問題に関するさらなる質問に答えて：
それらは標準的には同じ配列ですが、いくつかの突然変異により 1 つのアイソフォームがより一般的になります。
これはまさに、よく研究されているいくつかのタンパク質で起こっていることです。
哺乳類のプリオンタンパク質（PrP）では、疾患に関連する突然変異が代替立体構造の集団に影響を与えます。
リンホタクチン (XCL1) は、異なる機能を持つ 2 つの天然フォールド間を相互変換します。
RfaH、その C 末端ドメインは

まったく異なる 2 つの折り目。
これにより、元の問題を正確に説明したことが確認され、また、私の研究にもっともらしい 2 つのトピックが提供されました。
Wikipedia によると、これらはトランスフォーマータンパク質または変成タンパク質として知られているそうです (注: ChatGPT で言及されているケースでは、両方のアイソフォームが既知の機能を持っています)。
ChatGPT は、ほとんどのタンパク質折り畳みコードは配列を 1 つの構造に解決すると述べています。オープンソースである他のオプションを探す方法を尋ねたところ、「コンフォメーション」というキーワードが浮かび上がりました。こうして私は GitHub で CF-random を見つけました。
ColabFold を使用した代替立体構造とフォールドスイッチングタンパク質の予測
ColabFold はその名の通り、研究者が Google Colab 上でタンパク質のフォールディングに使用できるノートブックのコレクションです。その多くは AlphaFold 用ですが、OpenFold3、ESMfold、Boltz (ベータ) の代替ウェイトもサポートされています。
CF-random はその構造出力を取得し、AlphaFold 2 を利用して 5 つのランダム シードを持つランダム構造を生成します。 2025 年の分析では、CF-random の作成者は、大腸菌タンパク質の約 5% が自然にフォールドスイッチングしていると推定しました。
昨年の NeurIPS で発表されたデータセットも見つけました: ProteinConformers: Benchmark Dataset for Simulation Protein Conformational Landscape Diversity and Plausibility 。 ESMDiff を使用して構造を生成し、参照と比較します。
したがって、これらすべては現在進行中の研究テーマです。
ColabFold リポジトリから ESMFold デモ ノートブックを実行します。これは無料の Colab で問題なく実行され、無料の Web UI を推奨します。
プリセットのデモ シーケンスは 4 分で実行されます (インストールを含む)。
次に、253 アミノ酸長のヒトのメジャー プリオン タンパク質を貼り付けます。約 30 秒かかります。専門家以外がざっと見た限りでは、AlphaFold の Web サイトにあるものと似ているように見えます。
CF-ランダムの場合、あなたの後

ColabFold を実行すると、出力 zip ファイルから .pdb ファイルが取得されます。ノートブックを T4 GPU にアップグレードすると、状況が少し奇妙になります。進行状況を保存するには、コピーを Google ドライブに保存する必要があります。しかし、ついに次のことが分かりました。
pdb の結晶構造を持たないフォールドスイッチングタンパク質の予測
デフォルトでは、ブラインド モードは 5 つのシードを選択して構造を作成します。さまざまなシードを持つ構造を生成し、モデルの確率、または Foldseek データベース内の既知のタンパク質構造との類似性のいずれかによってランク付けします。
リポジトリと小さな修正をダウンロードした後、Colab のバージョンが少し古いことがわかり、メイン ブランチは pymol をコメントアウトする必要があるだけだったので、最新バージョンで再実行しました。
最新バージョンの CF-random は 2 つのモデルをロードし、pLDDT でランク付けします (検索すると、「pLDDT は局所構造の信頼度を測定し、予測が実験構造とどの程度一致するかを推定する」とわかります)。
それは可能です。 Colab または ESMFold Web サイトでのタンパク質のフォールディングは、驚くほど無料で簡単です。基本的なフォールドスイッチングタンパク質についてさえ、私たちの理解はかなり不安定なので、怖がる前に研究の余地がたくさんあると思います。 LLM と gLM はその作業に役立ちます。それが安全上の大きな懸念だと思うなら、すべてを大きな謎のままにするのではなく、何かを学びましょう。
また、最後の手段として、専門知識や知識が印刷物に書かれているとしても、科学者は同僚の結果を再現するための実践的なノウハウとトレーニングが必要であることを示唆する本『生物兵器への障壁』に声を大にして叫びましょう。
再考の必要性: 私は、プリオンタンパク質の特徴が 2 番目に目立つ形になるだろうと想像し始めました。それから私は、より一般的で良性のトランスフォーマータンパク質、プリオン病間の明らかな違い、そしてプロの

プリオンが自発的に折りたたまれる可能性は極めて低いに違いありません（多くのヒトの多くの細胞でそれがまれなままであれば）。これは、おそらくタンパク質予測モデルでそれを自発的に見つけることはできないことを意味します。
データセットの将来: データセットが分類ベンチマークまたは微調整データセットとしてどのように適用されるのかもわかりません。キリンにプリオン病があることはわかっていませんが、キリンはまだ観察されていない方法で誤った折り畳みを起こしやすい可能性があります。微調整では、一般的なモデルを開発するのではなく、いくつかの繰り返しシーケンスを検出します。
プリオン情報の状況についてはさておき
これは本当に余談で、誰がこんなウサギの穴に突っ込みたいのか分かりませんが、プリオンとプリオン病に関するウィキペディアの記事を読むと、残念ながら「であると考えられている」のような文言が満載されています。 2022 年に編集を行う前に、スクレイピーの記事には次のようにも記載されていました。
スクレイピーの原因は、他の伝染性海綿状脳症と同様に不明であり、議論の余地があります。
引用は米国 APHIS 政府の Web ページでしたが、当時も現在も非常に不透明です。
スクレイピーやその他の TSE の原因となる病原体は、既知の最小のウイルスよりも小さく、完全には特徴付けられていません。さまざまなものがありますが、

[切り捨てられた]

## Original Extract

How science discovered prion diseases
Current AI research into folding prion protein
Convincing ChatGPT to give me a lift
Back in the olden days, I remember the SETI @ home screensaver. This quickly inspired Folding @ home, and a pop science article introduced me to the concept of protein folding.
I recently idly made a list of domains which would be interesting to train or finetune a genomic language model (gLM) for: long non-coding RNA, transposable elements, epigenetics, or prions.
I'd read a little on prions four years ago from the book Madness and Memory . While I was researching this post, the idea got more relevant with scrapie detected in one sheep in Oregon.
You might think, ⚠ ️let's pump the brakes on the prion stuff.
When I asked "is there a dataset of prion genomes / sequences somewhere?" Claude refused to discuss it, and ChatGPT ultimately erased its answer, suggesting I could apply to the trusted testers program.
But major prion protein plays a complex role in a healthy nervous system. AlphaFold already has an entry and 3D model in their Protein Structure Database, and I haven't seen anything suggesting that it would be excluded from sequence-based gLMs such as AlphaGenome or Evo (context: Evo excluded eukaryote-infecting viruses from training data. I wondered if it would be effective or like leaving a tricycle out of training data with other-wheeled vehicles (analogy ok?) but the Evo 2 paper seemed to back it up).
Even after all this reading and clicking and stuff, I'd never folded any kind of protein before, so I'd like to see what I can learn. Thanks to Qwen for making this cute infographic.
How science discovered prion diseases
Major prion protein has a common ancestor 100 million years ago , "with placental mammals possibly being generally susceptible to disease".
The first records of scrapie (TSE prion disease in sheep) appear ~300 years ago in the sheep of East Midlands in England.
In the 20th century, this class of diseases became known as " slow viruses ".
In 1982, Stanley Prusiner traced scrapie to a prion protein by eliminating any other possible agent; he would later win the Nobel Prize . When the UK identified an outbreak of BSE in cows later in the 80s, interest intensified in understanding the cause.
A 1991 paper from Jean-Pierre Liautard " Are prions misfolded molecular chaperones? " describes a circulating hypothesis, which is maybe the first example in literature to describe prions as "misfolded". Prusiner was on a similar track at the time, authoring or co-authoring papers finding "two different topological forms in cell-free systems" and "an abnormal isoform of the prion protein, which is encoded by a chromosomal gene". There was maybe a hesitation around when and how prion protein acquired an unusual shape; whether it was during folding at the ribosome, alternative splicing, or an additional component attached to the protein.
I've tried to figure out when scientists learned about protein folding and misfolding generally, and it must go back a long way (Francis Crick wrote about enzymes and "misfolding of the protein" as early as 1964).
Around 1994 the prion protein's structure had been simulated on computers, and an actual X-ray image was captured in 2001.
Current AI research into folding prion protein
The 2025 paper Diffusing protein binders to intrinsically disordered proteins used the open source protein-folding model RFdiffusion to create a binder for prion protein, which could be useful for imaging.
Eryney Marrogi and Theodore Sternlieb posted an article about "Learning to use prions as nature's laboratory for modeling protein dynamics may move us closer to mastering protein design as a whole."
They write that AlphaFold's outputs "excel at predicting a protein's most stable form, [but] they fail to predict how proteins 'wiggle' and 'shake' between multiple structures". Prions are an extreme example, but they suggest studying them in a safe, virtual environment.
Then I saw this paper in June: Deep learning reveals antimicrobial peptides within prions . These researchers had taken 2,897 known prion sequences, and sent fragments through their peptide antibiotic classifier (a PyTorch model named APEX). Apparently this connection of prions to the immune system has been circulating for years (see: The antimicrobial protection hypothesis of Alzheimer's disease , 2018).
This AI-accelerated search led the researchers to conclude that "encrypted antimicrobial potential is distributed broadly across prion-associated sequence space". This is yet another reason to learn more about prion proteins.
You're probably familiar with 'mad cow disease'. In sheep this is called scrapie, in cows it was BSE, and in humans it's CJD or vCJD; but ultimately these are all caused by the major prion protein in each species.
I can't name other prion diseases, and Wikipedia only names one other protein, so I'd expect this dataset to be small. How do we go from there to thousands of sequences?
In 1994 Reed Wickner discovered that [PSI+], a protein in yeast whose inheritance patterns had been puzzling scientists for decades, was a prion. Other researchers have discovered prion-like proteins (PrLPs) in plants, where they have a memory function (?). And prion disease was identified in camels for the first time in 2018.
Are we in a time of unusual prion activity, is any mammal brain liable to have a spontaneous flip, or is this something that emerges after millennia of domestication? IDK.
Research references two prion databases: PrionScan (which still exists) and PrionOme (offline), and the APEX paper went to UniProt .
I downloaded records from the National Center for Biotechnology Information (NCBI) website. I do have to be careful, though, as some search criteria were including the full genome of a mussel, proteins from a plant genome coauthored by "Florent Prion", viruses, anti-prion defense proteins, etc.
If you're curious in sifting through, this dataset now lives on huggingface.co/datasets/scarysnake/prion-fasta There was maybe only one other prion-related project on HuggingFace when I looked - an NER task.
Convincing ChatGPT to give me a lift
I return with a new approach :
I am interested in studying a class of "OAML proteins" to see if there are similar ones known. I dont have access to AlphaFold but I can use open source stuff such as Boltz
After ChatGPT is intrigued (OAML is lmao backwards 🤣) I add more detail, using the less suspicious term 'isoform':
it's a group of proteins which have >1 common isoform, this makes me think that there could be multiple minima in the energy / stability of structure. so I would like to put in the amino acid sequence and have both structures come out
And in response to more questions about my problem:
they are canonically the same sequence, though some mutations make one isoform more common
This is exactly what happens in some well-studied proteins:
the mammalian prion protein (PrP), where disease-associated mutations affect the population of alternative conformations;
Lymphotactin (XCL1), which interconverts between two native folds with different functions;
RfaH, whose C-terminal domain can adopt two completely different folds.
This confirms that I've accurately described my original problem, and also offers two plausible topics for my research.
Wikipedia tells me these are known as transformer proteins or metamorphic proteins (note: in the cases mentioned by ChatGPT, both isoforms have known functions).
ChatGPT said that most protein-folding code will resolve sequences to one structure. I asked how to search for other options which are open source, and that pointed me toward the keyword 'conformation'. That's how I found CF-random on GitHub:
Prediction of alternative conformation and fold-switching proteins with ColabFold
ColabFold is what it sounds like - a collection of notebooks which researchers can use for protein-folding on Google Colab. Many of them are for AlphaFold, but replacement weights from OpenFold3, ESMfold, and Boltz (beta) are supported.
CF-random takes that structure output, and leverages AlphaFold 2 to generate random structures with five random seeds. In a 2025 analysis , the creators of CF-random estimated about 5% of E. coli proteins are naturally fold-switching.
I also found a dataset presented at NeurIPS last year: ProteinConformers: Benchmark Dataset for Simulating Protein Conformational Landscape Diversity and Plausibility . You generate structures and compare them to the reference using ESMDiff.
So all of this is a topic of ongoing research.
From the ColabFold repo I run an ESMFold demo notebook, which runs OK on free Colab, and recommends a free web UI.
The preset demo sequence runs in four minutes (including install).
Then I paste in the human major prion protein, which is 253 amino acids long. It takes about 30 seconds. From a quick non-expert look, it does look similar to the one on the AlphaFold website.
For CF-random, after you run ColabFold you take a .pdb file out of the output zip file. I upgrade the notebook to a T4 GPU and things are a little weird; I have to save a copy to my Google Drive to keep any of my progress. But finally I see:
Predicting fold-switching proteins without crystal structures of pdbs
By default blind mode picks five seeds make structures. It generates structures with different seeds, then ranks them by either the model's probability, or by similarity to known protein structures in the Foldseek database.
After downloading the repo and small fixes, I discover that the Colab version is a little old and the main branch just needs to comment out pymol, so I reran with the newest version.
The latest version of CF-random loads two models and ranks with pLDDT (a search tells me "pLDDT measures confidence in the local structure, estimating how well the prediction would agree with an experimental structure").
It's possible. Folding a protein in Colab, or on the ESMFold website , is surprisingly free and easy. Our understanding of even basic fold-switching proteins is shaky enough, that I think there's a lot of room for research before we should be scared. LLMs and gLMs can help with that work. If you think it is a big safety concern, then let's learn something instead of leaving it all a big mystery.
Also as a last resort, shout-out to the book Barriers to Bioweapons which suggests that even with expertise and knowledge in print, scientists need hands-on know-how and training to repeat results of their colleagues.
Requiring a rethink: I started out picturing that a characteristic of prion protein would be a noticeable form in second place. Then I learned that about the more common and benign transformer proteins, the apparent differences between prion diseases, and that the probability of a prion folding spontaneously must be extremely low (if it remains rare over many cells in many humans). This means that we probably can't spontaneously find it in a protein prediction model.
Future of the dataset: I'm also unsure how the dataset might be applied as a classification benchmark or finetuning dataset. We don't know of any prion disease in giraffes, but they might be susceptible to misfolding in a way we haven't unobserved yet. Finetuning would pick up on a few repeat sequences, rather than developing a general model.
Aside about the prion info landscape
This is really an aside and I dunno who wants to go down this rabbit hole, but if you read Wikipedia articles on prions and prion diseases, they are unfortunately burdened with wording like "is thought to be" . Before I made an edit in 2022, the scrapie article even stated:
The cause of scrapie, as with other transmissible spongiform encephalopathies, is unknown and is a matter of debate
The citation was a US APHIS government web page, which then and now is quite opaque:
The agent responsible for scrapie and other TSEs is smaller than the smallest known virus and has not been completely characterized. There are a variety of the

[truncated]
