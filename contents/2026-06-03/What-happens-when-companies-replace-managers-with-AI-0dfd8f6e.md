---
source: "https://analysis.infocentral.net/replacing-managers-with-ai.html"
hn_url: "https://news.ycombinator.com/item?id=48370221"
title: "What happens when companies replace managers with AI?"
article_title: "What will most likely happen when companies replace managers with AI"
author: "felineflock"
captured_at: "2026-06-03T00:45:40Z"
capture_tool: "hn-digest"
hn_id: 48370221
score: 7
comments: 3
posted_at: "2026-06-02T13:46:42Z"
tags:
  - hacker-news
  - translated
---

# What happens when companies replace managers with AI?

- HN: [48370221](https://news.ycombinator.com/item?id=48370221)
- Source: [analysis.infocentral.net](https://analysis.infocentral.net/replacing-managers-with-ai.html)
- Score: 7
- Comments: 3
- Posted: 2026-06-02T13:46:42Z

## Translation

タイトル: 企業がマネージャーを AI に置き換えると何が起こるか?
記事タイトル: 企業がマネージャーを AI に置き換えると、最も起こりやすいこと
説明: トレーニング データ、組織構造、および中間管理職を AI に置き換えることについての歴史的証拠が示唆するものについて。

記事本文:
エッセイ
経営とAI
企業がマネージャーを AI に置き換えると、おそらく何が起こるでしょうか
私は LLM 内の管理トレーニング データを探してみました。私が見つけたもの、あるいは見つけられなかったものは、「中間管理職を AI に置き換える」という議論についての私の考え方を変えました。
トレーニング ミックスが公開されている主要な LLM はすべて、同じストーリーを伝えています。一部のドメイン、特にコード、生物医学文書、法律文書、数学、学術文献は、明示的なソースまたはカテゴリとして表示されます。経営陣は通常そうではありません。
コードは専用の厳選されたソースから取得されます。医薬品は専用の調達が行われます。ローは専用の調達を行います。管理は、どの項目にも専用のカテゴリとして表示されません。
公開されているモデル ペーパーとデータセット マニフェストからの数値は次のとおりです。
コードには、GitHub スクレイピング、The Stack (358 言語で許可されたソース コード 3.1 TB)、および StarCoder の追加の 35B Python トークンが含まれています。 Llama 3 は、15 兆のトレーニング トークンの 17% をコードに割り当てます。コードには、HumanEval、MBPP などの専用ベンチマークもあります。
Medicine は、The Pile (EleutherAI の広く使用されているオープン トレーニング データセット) の 14.4% で PubMed Central を獲得し、さらに 3.1% で PubMed Abstracts を獲得しています。生物医学コンテンツの合計: そのコーパスの約 17.5%。
Law は、The Pile の 6.1% で FreeLaw (米国判例法) を取得し、さらに USPTO の特許背景が 3.7% で取得されています。
管理 - 主要な公開された事前トレーニング ミックスには、専用の管理コーパスが含まれていません。明示的な割り当てはありません。コード用の HumanEval や MBPP に匹敵する、広く採用されている標準的なベンチマークはありません。これらのモデルが持つ管理知識は、マーケティング ブログ、電子商取引のコピー、および一般的なビジネス コンテンツと混合した、Common Crawl Web フィルタリングを生き残るために起こったあらゆることから来ています。シェアを定量化しようとした出版された研究はありません。
LLM はほぼ確実に経営陣を見たことがある

書き込み。経営上のアドバイス、人事方針、リーダーシップに関する書籍、ビジネス ブログ、キャリア ガイダンスはすべて、ブロードウェブや書籍コーパスのどこかに掲載されています。
しかし、「マネジメント」は、開示されている主要なトレーニング構成の中では第一級のトレーニングカテゴリーとしては見られません。コードには、GitHub、The Stack、HumanEval、MBPP、および専用のコード モデルがあります。医学には PubMed と PubMed Central があります。 Law には FreeLaw と特許データがあります。数学と推論は、Llama 3 のようなモデルで明示的にバランスが取れています。対照的に、管理は未分化の Web、書籍、フォーラム コンテンツの中に埋もれています。
これは、証拠がより狭い主張を裏付けることを意味します。現在の LLM には管理知識が含まれている可能性がありますが、管理者の実際の行動を定義する判断、状況認識、および対人機能について、フロンティア モデルが意図的にトレーニングされ、バランスがとれ、ベンチマークされたという公的証拠はありません。
Google はこの質問に関して最もクリーンな自然実験を実施しました。 2001 年頃、創設者のラリー ペイジとサーゲイ ブリンは、エンジニアは各自の判断で任せるのが最善であると信じ、エンジニアリング マネージャーの役​​割をすべて廃止しました。実験は数か月続きました。ペイジとブリンのもとには、経費報告、人間関係の衝突、プロジェクトの優先順位付けに関する問い合わせがすぐに殺到しました。従業員はサポートや指導が不足していると不満を述べた。
Googleはマネージャーを復帰させた。そして 2008 年に、なぜマネージャーが重要なのかを理解するためにプロジェクト オキシジェンを立ち上げました。彼らは、100 以上の変数にわたって 10,000 を超える観測値を収集しました。この調査結果は、Google の人事担当副社長さえも驚かせたものでした。技術的な専門知識は、コーチング、コミュニケーション、権限付与、チームメンバーへの配慮、キャリア開発などの行動よりも下位にランクされていました。 1 位にランクされたのは、懸念される人間関係の行動と従業員の幸福へのサポートでした。
これらは最低限の機能です v

開示された事前トレーニングソース分類法では可能ですが、現在の AI ベンチマークで評価するのが最も困難です。
中間管理層を排除しようとした企業は Google だけではありません。このパターンは驚くほど一貫しています。
Zappos (2013) は、マネージャーや役職のないシステムであるホラクラシーを採用しました。数週間以内に、従業員の 14% (約 210 人) が新しいシステムの下で働くのではなく、買収を受け入れました。従業員は、誰が責任者なのか、何をするべきなのか、報酬の仕組みについて混乱していると報告した。同社は2023年までにシステムを大幅に変更し、別の名前で管理機能を密かに復活させた。
ゲーム会社 Valve は、「マネージャーがいない」文化で有名です。元従業員は別の話をします。元エンジニアのジェリ・エルズワース氏は、「社内で権力を獲得した人気のある子供たち」と評した。 Glassdoor のレビューでは、生き残るためには強力な内部グループに所属する必要があると説明されています。学術研究 (Foss and Klein、2022、 https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4289193 ) は、Valve で正式な階層が放棄されたとき、それは目に見えず、説明責任がなく、組織的な役割ではなく社会的支配によって運営されている非公式の階層に置き換えられたと結論付けています。
ミディアムは2013年にホラクラシーを採用したが、2016年に「仕事の邪魔になる」として廃止した。具体的な問題は、部門間の調整に時間がかかり、規模が大きくなると意見の対立が生じることです。
GitHub は 2016 年までフラットな構造で運営されていました。その期間中に、ある著名な開発者が共同創設者とその妻によるハラスメントを報告しました。この事件は、フラットな構造と説明責任、特に経営陣と個人の貢献者の間に管理層が存在しないことについての疑問を引き起こしました。 GitHub が部門別ファイルを復活させた

そして、それがより良い調整とより迅速な製品リリースを可能にしたと評価しました。
Buffer はフラットな構造を放棄し、共同創設者の Leo Widrich は「私たちが間違っていたこと」というタイトルの投稿を書きました。 ( https://buffer.com/resources/self-management-hierarchy/ ) Treehouse はマネージャーに戻り、上司不在の試みを「世間知らず」と呼びました。
どのケースでも、マネージャーを解任するために述べられた理論的根拠は同じでした：効率性、権限付与、官僚主義からの解放。いずれの場合も、実際の結果は、調整の失敗、目に見えない権力構造、従業員保護の喪失、そして最終的には異なる用語の下での管理層の復帰という組み合わせでした。
1945 年、フランスの政治哲学者ベルトラン・ド・ジュヴネルは、歴史を通じて中央集権的な権力がどのように成長してきたかを分析した『権力について』 (Du Pouvoir、 https://dn790003.ca.archive.org/0/items/onpoweritsnature00injouv/onpoweritsnature00injouv.pdf ) を出版しました。彼の中心的な見解は、主権者は中央権力と個人の間にある仲介機関、封建領主、ギルド、教会、専門職の命令、地方議会などを破壊することによってその権威を拡大するというものだ。
主権者のレトリックは常に解放である。私たちはあなたたちをこれらのつまらない暴君、非効率な仲介業者、官僚的な門番から解放します。
実際の結果はその逆です。中間層が取り除かれると、個人は、緩衝材も擁護者も、現場の特定の状況を理解する地元の保護者もいない状態で、中央集権的な権力にさらされることになる。仲介機関は乱雑で、時には腐敗していましたが、地元の現実を理解するための近さと、中央に対抗する立場の両方を備えた唯一の組織でした。
「中間管理職を AI に置き換える」ことと構造的に類似していることを無視するのは困難です。
T

経営層（主権者）は、AI システム（排除のメカニズム）を使用して中間管理職（中間権力）を排除し、個々の貢献者（国民）の効率化と権限付与を約束することを提案しています。しかし、AI システムは、それを構成した人、つまりエグゼクティブ層に報告します。エグゼクティブ層が定義するメトリクスを最適化します。チームへの忠誠心も、部門の存続に関わる出世権も、上からの不当な要求に抵抗する立場もありません。
中間管理職は、階層間の翻訳を行い、経営陣の戦略を実行可能な現地の意思決定に変換し、現場レベルの現実をリーダーが行動できる言語に変換します。彼らは曖昧さを吸収し、政策と現実の間のギャップで判断を下します。彼らはチームを上から守り、期限を交渉し、リソースを主張して保護します。彼らは状況に応じた知識を保持しています。つまり、クライアントのスケジュールがずれたため第 3 四半期の期限が柔軟であることや、新しいポリシーが静かな撤退のきっかけとなることを知っています。
これらの機能はいずれも、Web スクレイピング テキストのパターン マッチングによって解決できる単なる情報処理問題ではありません。それらは関係性があり、状況に応じて説明され、説明責任があります。
「フラット組織」の実験が実際に証明したもの
「成功した」フラットな組織はすべて、調べてみると、さまざまな名前で仲介権限を持っていることが判明します。 W.L.ゴア氏は工場の規模を約200人に制限しており、有機的に現れる「リーダー」を擁している。モーニングスターには、労働者間の正式な二国間約束制度があります。 Buurtzorg はコーチ付きの看護師 10 ～ 12 名でチームを構成します。それらのどれも中間層を排除しませんでした。彼らはそれを再設計しました。
失敗した実験の方が有益です。いずれの場合も、形式的な階層を削除しても階層は解消されませんでした。それにより、階層が目に見えなくなり、説明責任を負わなくなりました。 V

アルヴェの「人気キッズ」。 「依然として階層的に配置されていた」ザッポスのサークル。 GitHub のハラスメント論争は、管理層が不在の場合の説明責任についての疑問を引き起こしました。
ド・ジュヴネルは 1945 年に次のように予測しました。権力は、その形式的な構造を廃止しても消滅するわけではありません。それは非公式で不透明になり、異議を唱えるのが難しくなります。
これが「AI がマネージャーに取って代わる」というテーゼにとって何を意味するか
証拠が実際に何を裏付けているのかを明確にするために、この記事は異なるレベルの裏付けを持つ 3 つの異なる主張に基づいています。
1 つ目はデータ上の主張です。公開されている主要な事前トレーニング ミックスの中で、管理職が第一級のカテゴリーとして孤立しているわけではありません。これは最も強力な主張です。公開されているモデルペーパーやデータセットマニフェストに対して経験的に検証可能です。誰も反例を公開していません。それは、LLM が経営について何も知らないという意味ではありません。これは、フロンティア モデルがコード、医学、法律の場合と同じように意図的にトレーニングされ、バランスがとれ、ベンチマークが行われていることを誰も実証していないことを意味します。
2 つ目は能力に関する主張です。今日の AI は、管理者の業務の全範囲において依然として弱いままです。これはもっともらしいですが、あまり証拠がありません。これは、トレーニング データのギャップと管理作業の性質 (関係性、文脈、説明責任) から導き出されますが、現実的な管理タスク (組織再編の舵取り、離職リスクのある従業員の雇用維持、チームの対立の調停など) に関する AI パフォーマンスを測定する直接的な研究はほとんどありません。 Google の Project Oxygen は、マネジメントとは何かを教えてくれます。それ自体では、AI にはそれができないとは言えません。この主張は、証明された所見ではなく、合理的な推論として保持されるべきです。
3 つ目は、政治組織の主張です。中間層を置き換えると中央に権力が集中します。これは、ド・ジュヴネルから引き出された分析フレームワークであり、

その中で発見された因果律ではなく、組織のフラットな証拠に適用されます。歴史的な事件は実際にあります。パターンは一貫しています。しかし、同じ力関係が AI を介した管理でも繰り返されるという解釈は予測であり、確実なことではありません。それは保証としてではなく、入手可能な証拠の中で最ももっともらしい解釈として評価されるべきです。
これら 3 つの主張以外にも、注目に値する新たな経験的シグナルがあります。
悪いコードは目に見えてクラッシュします。悪い経営陣は文化を侵食し、離職率を高め、数か月または数年にわたって意思決定の質を低下させます。これはまさに、因果関係を特定するのが難しい種類のゆっくりとした損害です。結果が明らかになる頃には、解雇された管理者が保持していた組織上の知識は失われており、すぐに再構築することはできません。
2025 年の調査では、この懸念が単なる理論上のものではないことが示唆されています。ドンら。 (マックス プランク人間開発研究所 / ユトレヒト / トゥールーズ経済学部、arxiv.org/abs/2505.21752) は、カスタマイズされた Minecraft 職場で 382 人の労働者を人間、AI、またはハイブリッドの管理下に置きました。このプラットフォームが選ばれた理由は、リアルタイムの行動追跡、実際の偶発給による評価サイクルの繰り返し、実質に近い十分なタスクの自律性が可能であるためです。

[切り捨てられた]

## Original Extract

On training data, organizational structure, and what the historical evidence suggests about replacing middle management with AI.

Essay
Management & AI
What will most likely happen when companies replace managers with AI
I went looking for management training data inside LLMs. What I found - or didn't find - changed how I think about the "replace middle management with AI" conversation.
Every major LLM with a published training mix tells the same story: some domains, especially code, biomedical text, legal text, math, and academic literature, appear as explicit sources or categories. Management usually does not.
Code gets dedicated, curated sourcing. Medicine gets dedicated sourcing. Law gets dedicated sourcing. Management does not appear as a dedicated category in any of them.
Here are the numbers from published model papers and dataset manifests:
Code has GitHub scrapes, The Stack (3.1TB of permissively licensed source code in 358 languages), and StarCoder's additional 35B Python tokens. Llama 3 allocates 17% of its 15 trillion training tokens to code. Code also has purpose-built benchmarks: HumanEval, MBPP, and others.
Medicine gets PubMed Central at 14.4% of The Pile (EleutherAI's widely-used open training dataset), plus PubMed Abstracts at another 3.1%. Combined biomedical content: roughly 17.5% of that corpus.
Law gets FreeLaw (US case law) at 6.1% of The Pile, plus USPTO patent backgrounds at 3.7%.
Management - no major disclosed pretraining mix contains a dedicated management corpus. No explicit allocation. No canonical, broadly adopted benchmark comparable to HumanEval or MBPP for code. Whatever management knowledge these models have comes from whatever happened to survive Common Crawl web filtering - mixed in with marketing blogs, e-commerce copy, and generic business content. No published study has attempted to quantify the share.
LLMs almost certainly have seen management writing. Management advice, HR policies, leadership books, business blogs, and career guidance all appear somewhere inside broad web and book corpora.
But "management" is not visible as a first-class training category in the major disclosed mixes. Code has GitHub, The Stack, HumanEval, MBPP, and dedicated code models. Medicine has PubMed and PubMed Central. Law has FreeLaw and patent data. Math and reasoning are now explicitly balanced in models like Llama 3. Management, by contrast, is buried inside undifferentiated web, books, and forum content.
That means the evidence supports a narrower claim: current LLMs may contain management knowledge, but there is no public evidence that frontier models were deliberately trained, balanced, and benchmarked for the judgment, contextual awareness, and interpersonal functions that define what managers actually do.
Google ran the cleanest natural experiment on this question. Around 2001, founders Larry Page and Sergey Brin eliminated all engineering manager roles, believing engineers were best left to their own devices. The experiment lasted a few months. Page and Brin were immediately flooded with requests about expense reports, interpersonal conflicts, and project prioritization. Employees complained about the lack of support and guidance.
Google reinstated managers. Then, in 2008, they launched Project Oxygen to understand why managers mattered. They collected over 10,000 observations across 100+ variables. The finding that surprised even Google's VP of People Operations: technical expertise ranked below behaviors such as coaching, communication, empowerment, concern for team members, and career development. What ranked first were relational behaviors of concern and support for employee well-being.
Those are capabilities least visible in disclosed pretraining-source taxonomies and hardest to evaluate with current AI benchmarks.
Google is not the only company that tried to remove the intermediary management layer. The pattern is remarkably consistent:
Zappos (2013) adopted holacracy - a system with no managers or job titles. Within weeks, 14% of the workforce (roughly 210 people) accepted a buyout rather than work under the new system. Employees reported confusion about who was in charge, what they were supposed to do, and how compensation worked. By 2023, the company had significantly modified the system, quietly reinstating management functions under different names.
Valve , the gaming company, is famous for its "no managers" culture. Former employees tell a different story. Jeri Ellsworth, a former engineer, described "popular kids that have acquired power in the company." Glassdoor reviews describe needing to belong to powerful in-groups to survive. Academic research (Foss and Klein, 2022, https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4289193 ) concluded that when formal hierarchy was abandoned at Valve, it was replaced by informal hierarchy - invisible, unaccountable, and operating through social dominance rather than institutional role.
Medium adopted holacracy in 2013 and abandoned it in 2016, stating it was "getting in the way of the work." The specific problem: coordinating across functions became time-consuming and divisive at scale.
GitHub operated with a flat structure until 2016. During that period, a prominent developer reported harassment by a co-founder and his wife - a case that raised questions about flat structures and accountability, particularly regarding the absence of management layers between executives and individual contributors. GitHub reinstated departmental leadership and credited it with enabling better coordination and faster product releases.
Buffer gave up its flat structure, with co-founder Leo Widrich writing a post titled "What we got wrong." ( https://buffer.com/resources/self-management-hierarchy/ ) Treehouse returned to managers, calling the no-boss attempt "naive."
In every case, the stated rationale for removing managers was the same: efficiency, empowerment, liberation from bureaucracy. In every case, the actual outcome was some combination of coordination failure, invisible power structures, loss of employee protection, and eventual reinstatement of the management layer under different vocabulary.
In 1945, the French political philosopher Bertrand de Jouvenel published "On Power" (Du Pouvoir, https://dn790003.ca.archive.org/0/items/onpoweritsnature00injouv/onpoweritsnature00injouv.pdf ), analyzing how centralized power grows across history. His core observation: the sovereign expands its authority by destroying intermediary bodies - feudal lords, guilds, churches, professional orders, local assemblies - that stand between the central authority and the individual.
The sovereign's rhetoric is always liberation. We are freeing you from these petty tyrants, these inefficient middlemen, these bureaucratic gatekeepers.
The actual result is the opposite. Once the intermediary layer is removed, the individual stands exposed to centralized power with no buffer, no advocate, no local protector who understands the specific conditions on the ground. The intermediary bodies were messy and sometimes corrupt - but they were the only structures with both the proximity to understand local reality and the standing to push back against the center.
The structural parallel to "replace middle management with AI" is hard to ignore.
The executive layer (the sovereign) proposes to eliminate middle management (intermediary powers) using AI systems (the mechanism of removal), promising efficiency and empowerment for individual contributors (the people). But the AI system reports to whoever configured it - which is the executive layer. It optimizes for metrics the executive layer defines. It has no loyalty to the team, no career stake in the department's survival, no standing to push back on unreasonable demands from above.
Middle managers translate between layers - converting executive strategy into actionable local decisions, and converting ground-level reality back into language the leadership can act on. They absorb ambiguity, making judgment calls in the gap between policy and reality. They protect - shielding teams from above, negotiating deadlines, advocating for resources. They hold contextual knowledge - knowing that the Q3 deadline is flexible because the client's timeline slipped, or that the new policy will trigger quiet departures.
None of these functions are merely information-processing problems solvable by pattern-matching over web-scraped text. They are relational, contextual, and accountable.
What the "flat org" experiments actually prove
Every "successful" flat organization turns out, on inspection, to have intermediary powers under different names. W.L. Gore caps plant size at roughly 200 people and has "leaders" who emerge organically. Morning Star has a formal bilateral commitment system between workers. Buurtzorg caps teams at 10-12 nurses with a coach. None of them eliminated the intermediary layer. They redesigned it.
The failed experiments are more instructive. In every case, removing formal hierarchy did not eliminate hierarchy. It made hierarchy invisible and unaccountable. Valve's "popular kids." Zappos's circles that were "still arranged hierarchically." GitHub's harassment controversy that raised questions about accountability in the absence of management layers.
De Jouvenel predicted this in 1945: power does not disappear when you abolish its formal structures. It becomes informal, opaque, and harder to challenge.
What this means for the "AI-replaces-managers" thesis
To be clear about what the evidence actually supports, this article rests on three distinct claims with different levels of support.
The first is a data claim: management is not isolated as a first-class category in any major disclosed pretraining mix. This is the strongest claim. It is empirically verifiable against published model papers and dataset manifests. Nobody has published a counterexample. It does not mean LLMs know nothing about management. It means no one has demonstrated that frontier models were deliberately trained, balanced, and benchmarked for it the way they were for code, medicine, or law.
The second is a capability claim: AI today remains weak at the full scope of what managers do. This is plausible but less well-evidenced. It follows from the training data gap and from the nature of management work (relational, contextual, accountable), but direct studies measuring AI performance on realistic management tasks - navigating a reorg, retaining a flight-risk employee, mediating a team conflict - are scarce. Google's Project Oxygen tells us what management is. It does not tell us, on its own, that AI cannot do it. This claim should be held as a reasonable inference, not a proven finding.
The third is a political-organizational claim: replacing the intermediary layer concentrates power at the center. This is an analytical framework drawn from de Jouvenel and applied to the flat org evidence, not a causal law discovered in it. The historical cases are real. The pattern is consistent. But the interpretation - that the same dynamic will repeat with AI-mediated management - is a projection, not a certainty. It should be weighed as the most plausible reading of the available evidence, not as a guarantee.
Beyond those three claims, there is an emerging empirical signal worth watching.
Bad code crashes visibly. Bad management erodes culture, increases turnover, and degrades decision quality over months or years - exactly the kind of slow damage that is hard to attribute causally. By the time the consequences are obvious, the institutional knowledge held by the displaced managers is gone and cannot be rebuilt quickly.
A 2025 study suggests this concern is not just theoretical. Dong et al. (Max Planck Institute for Human Development / Utrecht / Toulouse School of Economics, arxiv.org/abs/2505.21752 ) placed 382 workers under human, AI, or hybrid management in a customized Minecraft workplace - a platform chosen because it allows real-time behavioral tracking, repeated evaluation cycles with actual contingent pay, and enough task autonomy to approximate real

[truncated]
