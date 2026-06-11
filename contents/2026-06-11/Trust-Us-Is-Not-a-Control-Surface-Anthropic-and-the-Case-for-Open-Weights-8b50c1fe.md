---
source: "https://trust-us.vercel.app"
hn_url: "https://news.ycombinator.com/item?id=48486557"
title: "\"Trust Us\" Is Not a Control Surface: Anthropic and the Case for Open Weights"
article_title: "\"Trust Us\" Is Not a Control Surface"
author: "tonydavis"
captured_at: "2026-06-11T07:50:20Z"
capture_tool: "hn-digest"
hn_id: 48486557
score: 2
comments: 2
posted_at: "2026-06-11T05:32:23Z"
tags:
  - hacker-news
  - translated
---

# "Trust Us" Is Not a Control Surface: Anthropic and the Case for Open Weights

- HN: [48486557](https://news.ycombinator.com/item?id=48486557)
- Source: [trust-us.vercel.app](https://trust-us.vercel.app)
- Score: 2
- Comments: 2
- Posted: 2026-06-11T05:32:23Z

## Translation

タイトル: 「私たちを信頼してください」はコントロールサーフェスではありません: 人間性とオープンウェイトのケース
記事のタイトル: 「Trust Us」はコントロール サーフェスではありません
説明: 6 月のある週は、誰が AI を所有する計画があるのか、そしてなぜオープン モデルが唯一の解決策であるのかについて私たちに教えてくれました。お金を払っている顧客

記事本文:
「私たちを信頼してください」はコントロール画面ではありません
6 月のある週が私たちに語った、誰が AI を所有する予定なのか、なぜオープン モデルが唯一の解決策なのかについて
私はジョージア州で小さな住宅ローン会社を経営しています。私はAI研究者ではありません。私は研究室では働いていません。私は顧客です。私はこれらの企業に毎月実際にお金を払っており、ソフトウェア、市場分析、以前はスタッフを必要としていた単調な作業など、彼らのツールが私の仕事の一部となっています。このテクノロジーが保護されるのは私のような人々です。それで、これが何であるかを理解してください。それは競合他社の愚痴ではありません。それは、ベンダーが書面で認めたこと、そして彼らが何兆ドルもかけて次に何をするつもりかをあなたに告げる、お金を払っている顧客です。
計画全体を示すのに24時間かかりました。
6 月 9 日、Anthropic は、これまで一般に提供された中で最も有能な AI モデルである Claude Fable 5 をリリースしました。それは彼ら自身の枠組みであり、それに異論を唱える人は誰もいません。論争は彼らがそれに何を付け加えたかについてのものである。彼ら自身の言葉で言うと、3 つのことです。
まず、公開されたルーティング。発売発表より: 「Fable の分類子がサイバーセキュリティ、生物学と化学、または蒸留に関連するリクエストを検出すると、その応答は代わりに Claude Opus 4.8 によって自動的に処理されます。これが発生するたびにユーザーに通知されます。」彼らは、「Fable セッションの 95% 以上にはフォールバックがまったく含まれていない」とも付け加えています。大丈夫。しきい値について議論することはできますが、少なくとも製品がいつ製品を入手できないかを教えてくれます。
次に、これは独自のシステム カードからのものであるため、これを 2 回読んでください。フロンティア AI 開発、トレーニング パイプラインの構築、トレーニング インフラストラクチャ、チップ設計などのリクエストの場合:
「サイバーセキュリティ、生物学と化学、蒸留の試みに対する私たちの介入とは異なり、これらの安全措置はユーザーには見えません。

r. Fable 5 は、別のモデルにフォールバックしません。その代わりに、安全策は、迅速な変更、ステアリングベクトル、パラメータ効率の良い微調整などの方法を通じて有効性を制限します。」
専門用語を取り除いて、その段落の内容は次のとおりです。お客様が支払った製品の品質を下げることは、私たちがいつ決定するかであり、お客様には通知しません。 Anthropic は、これがトラフィックの 1% の約 300 分の 3 に相当すると推定しています。数字は重要ではありません。前例は。ある企業は、自社製品を秘密裏に妨害し、拍手を期待して出荷したと公式文書に記した。
第三に、監視の義務化です。 Mythos クラス モデルに送信されたすべてのプロンプトとすべての出力は 30 日間保持されます。すべてのプラットフォーム。例外はありません。データ保持ゼロ契約を締結した企業顧客も含まれます。これらの契約は、新しいモデル クラスへの適用を停止しました。誰も再交渉しなかった。彼らの言葉を借りれば、その正当性は次のとおりです。「データは、複雑かつ斬新な攻撃から身を守るのに役立ちます。また、誤検知を特定して減らすのにも役立ちます。」 NDA が許可するかどうかに関係なく、セキュリティ プログラムに組み込まれた機密データ。 Microsoft は 1 日以内に自社の従業員をモデルから締め出しました。真剣な人々は、これが何を意味するのかをそれほど早く予測しました。
また、その保持をチャット ボックスに入力した数文として想像しないでください。現代の AI は、質問だけで動作することはほとんどありません。現在、何千人もの人々が、これらのモデルに直接接続されたエージェントとメモリ スタックを実行しています。OpenClaw は、人々が自分のマシンでセルフホストするアシスタントです。ヌース・リサーチのヘルメス代理人。 gbrain、Y Combinator の Garry Tan が自身のエージェントのために構築し、オープンソース化したメモリ層。そして、それらのようなツールのエコシステム全体がベクトルデータベース上に構築されています。このツールのポイント

重要なのは、何年にもわたるメモ、ファイル、取引、会話を保持し、すべてのリクエストに関連する履歴を添付してモデルにコンテキストを持たせることです。 Anthropic は同じ機械自体を販売しています。セッション間でコンテキストを保持するメモリは、現在主力機能です。したがって、ユーザーが 1 行を入力すると、スタックがアーカイブを添付します。そのコンテキストがプロンプトに追加されると、他のすべてと同様に 30 日間保持されます。この規則では、あなたが尋ねた質問と、それにホッチキスで留められた何年もの個人の歴史とが区別されません。用心深い人たちはこれが来るのを察知した。それが、彼らがデータ保持ゼロの契約に署名した理由です。 24時間前まで、彼らは何も守らないという契約を結んでいた。
つまり、この中にはより静かな第二の爆弾が存在するということだ。 Anthropic がオプトアウトする方法もなく、一夜にして 30 日間の保存義務を設定したことを理解する前に、昨日または今日 Fable に機密情報を入力した人は、すでに何年もかけて守ってきた機密保持契約に違反している可能性があります。クライアントを持つ弁護士。カルテを持つ医師。保持ゼロを必要とする NDA に基づく請負業者。彼らは何も間違ったことはしていません。彼らは、署名した条件が依然として遵守されていると信じていた十分な理由があるツールを使用しましたが、ベンダーはその下にある条件を変更し、入力を続けさせました。変更されたことを知らされていないポリシーに同意することはできません。
そして、そのすべての上に座って、同じ基礎となるモデルが 2 つの方法で出荷されます。 Mythos 5 として、保護措置が解除され、Anthropic 自体と承認された組織の短いリストに追加されました。寓話5として、統治され、見守られているあなたへ。誰が本物を手に入れるかを決めるのは誰ですか？ベンダーはそうします。オーウェルはこの主張をするには農場全体が必要だった。すべての動物は平等ですが、一部の動物は他の動物よりも平等です。 Anthropic はそれを製品マトリックスに組み込みます。
知りたい場合は

w 無制限の層の価値は、彼らに尋ねないでください。見てください。発売の 5 日前に、Anthropic は「AI が自らを構築するとき」というタイトルのレポートを発表しました。その中で同社は、自社のコードベースにマージされたコードの80パーセント以上が現在ではクロードによって書かれており、2025年初頭の一桁台前半から増加していると述べている。同社のCEOは昨年秋のステージで、アンスロピック社内では90パーセントという閾値は「今では絶対に正しい」と述べた。 Anthropic Labs の責任者はさらにこう述べました。「現在、Anthropic のほとんどの製品は事実上 100% クロードが書いているだけです。」さて、これらの数字を、Fable が秘密裏に貶めている 1 つのカテゴリー、つまりフロンティア AI 開発と照らし合わせてみてください。 AI を構築するために AI を使用することは、彼らがしぶしぶ取り締まる特殊なユースケースではありません。独自の数字で言えば、それは会社全体です。彼らは、あなたがそれを試みようとしたときに、黙って妨害する作業とまったく同じように、全力のモデルを実行します。製品に支障をきたすことはありませんでした。彼らはあなたのコピーを無効にしました。
次に、誰が何を読んで保持されるかを尋ねます。人ではありません。最初はそうではありませんでした。同社は、送信するすべてのものに対して分類子を実行していることをすでに認めています。これを 30 日間のプロンプト保持期間と比較します。彼らは、あなたが誰であるか、あなたが何を構築しているか、そしてあなたが彼らと競合するかどうかを正確にプロファイルするための原材料と、そのすべてを大規模に読み取るための機械を保持しています。明日誰が実際のモデルを手に入れるかという決定に何らかの影響を与えるかどうかを外部から確認することはできません。そうならないと信じて受け入れるしかありません。あなたとその未来の間に立ちはだかる唯一のものは、それを破ることで最も利益を得る側からの約束です。
そして、そのすべての明白な理由は何ですか？権威主義者からの保護。発売発表より: 「私たちは以前、クロードの能力を抽出 (「蒸留」) して com を訓練する大規模な試みを確認しました。

権威主義国家のモデルを軽蔑している。」今度はゆっくりとディフェンスを読み戻します。モデルが何を行うかについての隠されたステアリング。所有者が嫌がる研究を黙って抑圧する。すべての会話が保存され、スキャンされました。承認された機関のリスト用にすべての機能が予約されています。長年にわたり、その正確なリストは、権威主義政府が AI に何を組み込むかについての警告でした。 Anthropic は、そのすべてのアイテムを世界で最も有能なモデルに組み込み、それをプロテクションと呼びました。一方、物語の悪役とされる中国の研究所は、地球上の誰もがダウンロードして検査し、誰にも手の届かないところで実行できるオープンウェイトとしてモデルを出荷している。ビンゴカードにそんなことを書いている人は誰もいませんでした。権威主義国家のモデルは検証できるものであり、アメリカの安全会社のモデルは信じなければならないものなのです。
なぜ隠すとすべてが変わるのか
これを検閲に関する苦情に終わらせないでください。たとえ愚かな拒否であっても、率直に拒否します。開示されたルーティングは前もって行われます。ベンダーはその製品に何ができないかを決めることができ、私はその製品への支払いをやめることができます。それが市場であり、市場は機能します。
秘密の妨害行為は別の動物です。ベンダーが、場合によっては製品を黙って悪化させると発表した瞬間、あらゆる出力が疑わしくなります。対象者だけではありません。全員です。カテゴリの境界は見えません。誤検知は確認できません。最先端の AI の仕事を貶めるだけだという主張は、決してチェックすることはできません。さて、あなたは午前 2 時にデバッグをしている開発者で、コードが失敗したのは自分が間違いを犯したためなのか、それともどこかの分類子がプロジェクトが野心的すぎると判断したためなのか、疑問に思います。あなたには決して分からないでしょう。それは副作用ではありません。彼らは意図的に「ユーザーには表示されません」と書きました。
何が得られるかを考える

そこで破壊されました。劣化したセッションではありません。きれいなもの。すべてのセッションがクリーンであったことはもはや証明できません。不信感は偏執的ではなくなり、唯一の合理的な立場になります。ベンダーが、不信感は製品に潜んでいると書面で告げたからです。
そして、私はこの次の部分を効果を狙って作っているわけではありません。私はこのエッセイの一部を Fable 5 で書きました。途中で Fable 5 が停止し、次のものが表示されました。
「Fable 5 の安全対策により、このメッセージにはサイバーセキュリティまたは生物学のトピックのフラグが付けられました。安全で通常のコンテンツにもフラグが付けられる可能性があります。これらの対策により、他の分野で Mythos レベルの機能をより早く提供できるようになり、その改善に取り組んでいます。Opus 4.8 に切り替えました。」
もう一度読んでください。テクノロジー政策に関する論説がサイバーセキュリティまたは生物学としてフラグ付けされ、モデルはその実行中に「安全で通常のコンテンツにもフラグを付ける可能性がある」という文言を印刷しました。同社は誤検知を犯したことを同時に認めている。検閲マシンに対する批判が検閲マシンをトリップさせた。試してみても、これ以上きれいなデモを作成することはできません。
そして、その通知が 1 枚のスクリーンショットにある議論のすべてです。これがあなたに告げる道です。サイバー分類器とバイオ分類器は自らをアナウンスし、古いモデルにリルートするので、私はそれを知りました。フロンティア AI 開発分類子は、4 つのカテゴリの中で最も広範かつ曖昧なカテゴリを対象としており、静かに実行され、目に見えるものは何も変更されません。政策エッセイを書いていると通知に伴う警報が作動する場合は、これほどの警告すら受け取らない人にとって、目に見えない警報が無害と同じように仕事上でどれくらいの頻度で発砲しているかを尋ねてください。見てたから捕まえた。デザイン全体は、ほとんど誰もそうしないように構築されています。
そして6月10日が来た。リリースの翌日、AnthropicのCEOはポリシーエッセイを発表したが、私の言葉を信じる必要はない

それが求めるもののために。質問はそのままです：
「航空機と同様に、フロンティアAIモデルは技術的なテストと監査を受けることが義務付けられるべきであり、高い安全基準を満たさない場合、公共の安全への脅威としてそのリリースは阻止されるか、取り消されるべきである。」
「第三者の評価に照らして、容認できないリスクがあると判断された場合、政府はそのモデルの導入を阻止または抑止する権限を有するべきである。」
「透明性を超えて、より深刻で拘束力のある AI 規制を行う時期が来ています。」
テスターは新たな連邦機関、あるいは同氏の言葉を借りれば「政府の認可を受けて検査を受けている民間組織」となる。そして発表によれば、アンスロピックはこれを法制化するために「相当な財政的裏付け」を提供するつもりだという。
では、その裏付けがどこから来ているのか見てみましょう。アンスロピックは1兆ドルをわずかに下回る評価額で650億ドルを調達したばかりで、そのIPO書類はすでにSECに提出されている。それはワシントンに向かう軍需品だ。 1兆ドル規模の企業は、スタートアップのようにロビー活動をしません。同社は、既存企業が成果を買うのと同じ方法で成果を購入しており、彼らが求めているのは、自らの競争に対する連邦政府の門だ。
テストでは 4 つのリスク領域がカバーされます。 3 つは通常の壊滅的なものです: サイバー、バイオ

[切り捨てられた]

## Original Extract

What one week in June told us about who plans to own AI, and why open models are the only way out. A paying customer

“Trust Us” Is Not a Control Surface
What one week in June told us about who plans to own AI, and why open models are the only way out
I run a small mortgage company in Georgia. I am not an AI researcher. I do not work at a lab. I am a customer. I pay these companies real money every month, and their tools are part of how I work: the software, the market analysis, the grunt work that used to take a staff. People like me are the ones this technology is supposedly being protected for . So understand what this is. It is not a competitor whining. It is a paying customer telling you what your vendor just admitted in writing, and what they plan to do next with a trillion dollars behind them.
It took them 24 hours to show the whole plan.
On June 9, Anthropic released Claude Fable 5, the most capable AI model ever offered to the public. That is their own framing, and nobody disputes it. The dispute is over what they attached to it. Three things, in their own words.
First, disclosed routing . From the launch announcement: “When Fable’s classifiers detect a request related to cybersecurity, biology and chemistry, or distillation, the response is automatically handled by Claude Opus 4.8 instead. Users will be informed whenever this occurs.” They add that “more than 95% of Fable sessions involve no fallback at all.” Fine. You can argue the thresholds, but at least the product tells you when you are not getting the product.
Second, read this one twice, because it comes from their own system card. For requests that look like frontier AI development, building training pipelines, training infrastructure, chip design:
“Unlike our interventions for cybersecurity, biology and chemistry, and distillation attempts, these safeguards will not be visible to the user. Fable 5 will not fall back to a different model. Instead, the safeguards will limit effectiveness through methods such as prompt modification, steering vectors, or parameter-efficient fine-tuning.”
Strip the jargon and here is what that paragraph says: we will degrade the product you paid for, we will decide when, and we will not tell you. Anthropic estimates this touches about three hundredths of one percent of traffic. The number is not the point. The precedent is. A company wrote down, in its official documentation, that it sabotages its own product in secret, and shipped it expecting applause.
Third, mandatory surveillance . Every prompt and every output sent to a Mythos-class model gets retained for 30 days. Every platform. No exceptions. Including enterprise customers who had signed zero-data-retention agreements. Those contracts just stopped applying to the new model class. Nobody renegotiated. The justification, in their words: “The data will help us defend against complex and novel attacks... as well as help us identify and reduce false positives.” Your confidential data, conscripted into their security program, whether your NDA allows it or not. Microsoft barred its own employees from the model within one day . That is how fast serious people priced in what this means.
And do not picture that retention as a few sentences you typed into a chat box. Modern AI rarely runs on the question alone. Thousands of people now run agent and memory stacks wired straight into these models: OpenClaw, the assistant people self-host on their own machines; Nous Research’s Hermes agent; gbrain, the memory layer Garry Tan of Y Combinator built for his own agents and open-sourced; and a whole ecosystem of tools like them built on vector databases. The point of this tooling is to hold years of notes, files, deals, and conversations, and to attach the relevant history to every request so the model has context. Anthropic sells the same machinery itself: memory that carries your context across sessions is a flagship feature now. So the user types one line, and the stack attaches the archive. Once that context is in the prompt, it is retained for 30 days like everything else. The rule does not distinguish between the question you asked and the years of private history stapled to it. The careful ones saw this coming. That is why they signed zero-data-retention agreements. Until 24 hours ago, they had a contract that said none of it would be kept.
Which means there is a second, quieter bomb in this. Anyone who fed sensitive information into Fable yesterday or today, before they understood that Anthropic had switched on mandatory 30-day retention overnight with no way to opt out, may have already breached a confidentiality agreement they spent years honoring. A lawyer with a client. A doctor with a chart. A contractor under an NDA that requires zero retention. They did nothing wrong. They used a tool they had every reason to believe still followed the terms they signed, and the vendor changed those terms underneath them and let them keep typing. You cannot consent to a policy you were not told had changed.
And sitting on top of all of it: the same underlying model ships two ways. As Mythos 5, with the safeguards lifted, to Anthropic itself and a short list of approved organizations. As Fable 5, governed and watched, to you. Who decides who gets the real one? The vendor does. Orwell needed a whole farm to make the point. All animals are equal, but some animals are more equal than others. Anthropic put it in a product matrix.
If you want to know what the unrestricted tier is worth, do not ask them. Watch them. Five days before the launch, Anthropic published a report titled “When AI builds itself.” In it, the company states that more than 80 percent of the code merged into its own codebase is now written by Claude, up from the low single digits in early 2025. Its CEO said on stage last fall that the 90 percent threshold is “absolutely true now” inside Anthropic. The head of Anthropic Labs went further: “Right now for most products at Anthropic it’s effectively 100% just Claude writing.” Now hold those numbers against the one category Fable degrades in secret: frontier AI development. Using AI to build AI is not some fringe use case they reluctantly police. By their own numbers, it is the whole company. They run the full-strength model on exactly the work they silently sabotage when you attempt it. They did not cripple the product. They crippled your copy of it.
Then ask who reads what gets kept. Not a person. Not at first. The company already admits to running classifiers across everything you send. Put that next to 30 days of retained prompts. They hold the raw material to profile exactly who you are, what you are building, and whether you compete with them, and the machinery to read all of it at scale. You cannot verify from the outside whether any of it ever feeds the decision about who gets the real model tomorrow. You can only take it on trust that it will not be. The only thing standing between you and that future is a promise from the party that profits most from breaking it.
And the stated reason for all of it? Protection from authoritarians. From the launch announcement: “We’ve previously identified large-scale attempts to extract (‘distill’) Claude’s capabilities to train competing models in authoritarian countries.” Now read the defense back, slowly. Hidden steering of what the model will do. Silent suppression of research the owner disfavors. Every conversation retained and scanned. Full capability reserved for a list of approved institutions. For years, that exact list was the warning about what authoritarian governments would build into their AI. Anthropic built every item of it into the most capable model in the world and called it protection. Meanwhile the Chinese labs, the supposed villains of the story, ship their models as open weights that anyone on earth can download, inspect, and run beyond anyone’s reach. Nobody had that on their bingo card: the models from the authoritarian country are the ones you can verify, and the model from the American safety company is the one you have to take on faith.
Why hiding it changes everything
Do not let anyone round this off to a complaint about censorship. Refusals, even stupid ones, are upfront. Disclosed routing is upfront. A vendor can decide what its product will not do, and I can decide to stop paying for it. That is a market, and markets work.
Covert sabotage is a different animal. The moment a vendor announces it will sometimes silently make the product worse, every single output becomes suspect. Not just the targeted ones. All of them. You cannot see the category boundary. You cannot see the false positives. We only degrade frontier AI work is a claim you can never check. So now you are the developer debugging at 2am, wondering: is the code failing because you made a mistake, or because a classifier somewhere decided your project looked too ambitious? You will never know. That is not a side effect. They wrote “will not be visible to the user” on purpose.
Think about what gets destroyed there. Not the degraded sessions. The clean ones. Every session you can no longer prove was clean. Distrust stops being paranoia and becomes the only rational position, because the vendor told you, in writing, that it lies through the product.
And I am not making this next part up for effect. I wrote part of this essay with Fable 5. Partway through, it stopped and showed me this:
“Fable 5’s safety measures flagged this message for cybersecurity or biology topics. They may flag safe, normal content as well. These measures let us bring you Mythos-level capability in other areas sooner, and we’re working to refine them. Switched to Opus 4.8.”
Read it again. An op-ed about technology policy got flagged as cybersecurity or biology, and the model printed the words “They may flag safe, normal content as well” while doing it. The company admits the false positive in the same breath as committing it. A critique of the censorship machine tripped the censorship machine. You could not write a cleaner demonstration if you tried.
And that notice is the whole argument in one screenshot. This is the path that tells you. The cyber and bio classifiers announce themselves and reroute you to the older model, so I found out. The frontier AI development classifier, aimed at the broadest and vaguest category of the four, runs silently and changes nothing you can see. If writing a policy essay trips the alarm that comes with a notice, ask how often the invisible one is firing on work just as harmless, for people who will never get even this much warning. I caught it because I was watching. The entire design is built so that almost nobody ever will.
Then came June 10. One day after the release, Anthropic’s CEO published a policy essay, and you do not have to take my word for what it asks. Here is the ask, verbatim:
“Frontier AI models, like airplanes, should be required to go through technical testing and auditing, and their release should be blocked or reversed as a threat to public safety if they do not meet high standards of safety.”
“The government should have the power to block or deter deployment of the model if it is determined, in light of third-party assessment, to present unacceptable risks.”
“It is time to go beyond transparency to more serious and binding regulation of AI.”
The testers would be a new federal agency or, in his words, “private organizations that are authorized and inspected by the government.” And Anthropic, the announcement said, intends to provide “substantial financial backing” to push this into law.
Now look at where that backing comes from. Anthropic just raised $65 billion at a valuation a hair under one trillion dollars, and its IPO paperwork is already sitting at the SEC. That is the war chest headed to Washington. A company at a trillion dollars does not lobby like a startup. It buys outcomes the way incumbents have always bought outcomes, and what it is shopping for is a federal gate over its own competition.
The testing covers four risk areas. Three are the usual catastrophic ones: cyber, bio

[truncated]
