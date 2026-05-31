---
source: "https://linuxstans.com/ai-found-3900-critical-open-source-bugs-ibm-is-paying-5-billion-to-fix-them/"
hn_url: "https://news.ycombinator.com/item?id=48338895"
title: "AI Found 3,900 Critical Open Source Bugs. IBM Is Paying $5B to Fix Them"
article_title: "IBM Just Committed $5 Billion to Fix Open Source Security. The Linux Community Has Complicated Feelings About It."
author: "hochmartinez"
captured_at: "2026-05-31T00:27:03Z"
capture_tool: "hn-digest"
hn_id: 48338895
score: 4
comments: 0
posted_at: "2026-05-30T17:50:42Z"
tags:
  - hacker-news
  - translated
---

# AI Found 3,900 Critical Open Source Bugs. IBM Is Paying $5B to Fix Them

- HN: [48338895](https://news.ycombinator.com/item?id=48338895)
- Source: [linuxstans.com](https://linuxstans.com/ai-found-3900-critical-open-source-bugs-ibm-is-paying-5-billion-to-fix-them/)
- Score: 4
- Comments: 0
- Posted: 2026-05-30T17:50:42Z

## Translation

タイトル: AI がオープンソースの重大なバグを 3,900 個発見。 IBMはそれらを修正するために50億ドルを支払っている
記事のタイトル: IBM はオープンソース セキュリティの修正に 50 億ドルを投入しました。 Linux コミュニティはこれについて複雑な感情を抱いています。
説明: 2024 年に公開される CVE は 40,000 件。2026 年までに 59,000 件になると予測されています。IBM と Red Hat は、オープンソースのセキュリティ危機に対して 50 億ドルの解決策があると考えています。こちらが実際に構築しているものです。

記事本文:
IBM はオープンソース セキュリティの修正に 50 億ドルを投入しました。 Linux コミュニティはこれについて複雑な感情を抱いています。
コンテンツにスキップ
について
記事を書いてください – 記事を送信してください
検索
検索
検索…
検索
検索…
AI がオープンソースの重大なバグを 3,900 個発見。 IBMはそれらを修正するために50億ドルを支払っている
IBM の Project Lightwell の発表には、現時点で注目されている以上に注目に値する多くのものが埋もれています。 Anthropic の Mythos Preview AI モデルは、オープンソース ソフトウェアをスキャンし、3,900 件近くの高重大度または重大度の脆弱性を特定しました。それは何年にもわたる監査の遅れの結果ではありません。これは、あるフロンティア AI モデルがプレビュー実行で発見したものです。そしてモデルは改良されるばかりです。
それが、IBM と Red Hat が構築している世界です。 2026 年 5 月 28 日、両社は Project Lightwell への 50 億ドルの出資を発表しました。Project Lightwell はエンタープライズ オープン ソース ソフトウェアのセキュリティ クリアリングハウスであり、20,000 人のエンジニアと、攻撃者が脆弱性を武器化する前に脆弱性を発見して修正するように設計された AI ツールが支援しています。銀行はすでに登録済みです。 Linux コミュニティは非常に注意深く監視しています。
問題は現実であり、数字は醜くなっている
2024 年には 40,000 件を超える CVE が公開されました。IBM は、その数が 2026 年までに 59,000 件に増加する可能性があると予測しています。ソフトウェアが粗雑になっているため、その数が加速することはありません。 AI を活用した脆弱性発見が、人間のセキュリティ チームでは太刀打ちできないほどに拡大しているために、このようなことが起こっています。
Fortune 500 企業の 90% 以上がオープンソース ソフトウェアで運営されています。これらの CVE はいずれも、銀行、病院、電力網の実稼働システムへの潜在的な経路となります。これらの環境を動かすソフトウェアは、多くの場合、ボランティア、愛好家、および処理するための予算や帯域幅を持たずに活動する小規模なチームによって保守されています。

毎月何百件もの脆弱性レポートを作成しながら、機能の出荷やサポートの対応も行っています。
修復ギャップ、つまり脆弱性を発見してから、影響を受けるすべての運用環境に実際にパッチを適用するまでの距離は、個々の組織が単独でギャップを埋めることができるよりも速いスピードで拡大しています。それがプロジェクト・ライトウェルが埋めようとしているギャップだ。
ライトウェルプロジェクトが実際に行うこと
プレスリリースの文言を取り除くと、ここでは 3 つの具体的なことが起こっています。
調整されたセキュリティ情報交換所
企業は、安全な仲介フレームワークを通じて、機密の脆弱性を一般公開前に IBM および Red Hat に報告できます。 IBM は問題を検証し、企業独自のアプリケーション ソース コードにアクセスすることなく修正プログラムを開発します。修正は、顧客が管理するリポジトリに配信されます。
その後は上流へ向かいます。オープンソース プロジェクトはパッチを取得します。これは、より広範なエコシステムにとって最も重要な部分であり、IBM はそれについて明確に述べています。クリアリングハウス モデルは、上流のコミュニティを迂回するのではなく、上流のコミュニティを強化するように設計されています。
すでに実行しているものへのバックポート
これは、ほとんどのエンタープライズ チームが実際に気にする部分です。 Project Lightwell は、セキュリティ修正を取得するために依存関係をアップグレードするように組織に指示しません。実稼働環境ですでに実行されている正確なバージョンに修正をバックポートします。
企業のアプリケーションが 2022 年以降の特定の Java ライブラリ バージョンに固定されている場合、IBM はそのバージョンにパッチを適用します。強制的なアップグレードはありません。互換性のリスクはありません。 IBM は、pom.xml などの依存関係マニフェストに基づいて動作し、署名され検証されたパッケージを顧客が管理するリポジトリに配信します。当初の焦点は Maven と Java で、ロードマップには PyPI、npm、Go が含まれます。
大規模な AI 支援エンジニアリング
IBM は 20,000 人のエンジニアを配置しています

Red Hat と IBM と高度な AI ツールを組み合わせます。 AI は、大量の脆弱性のトリアージ、優先順位付け、初期パッチ開発を処理します。エンジニアは、実際に上流のプロジェクトや顧客環境に導入されるものをレビュー、形成し、出荷します。
IBM はすでに 62,000 を超えるオープンソース パッケージを使用しており、そのうち 10,000 を超えるパッケージについて深い専門知識を維持しています。対象範囲は、Linux、Kubernetes、Java、Kafka、Ansible、Terraform、Flink、Cassandra などです。 Lightwell は、そのモデルを Red Hat の従来の製品フットプリントを超えて、より広範なアプリケーション依存関係ツリーに拡張します。
早期導入者リストは冗談ではありません
IBMは、バンク・オブ・アメリカ、BNY、シティ、ゴールドマン・サックス、JPモルガン・チェース、マスターカード、モルガン・スタンレー、ロイヤル・バンク・オブ・カナダ、ステート・ストリート、ビザ、ウェルズ・ファーゴがすでにプロジェクト・ライトウェルで協力していると発表した。これらの組織は、プレスリリースが説得力があるため、登録していません。彼らが登録しているのは、広く使用されている Java ライブラリにパッチが適用されていない脆弱性があり、規制上および評判に悪影響を与える大惨事が待っているためであり、マネージド ソリューションに支払う予算があるからです。
彼らが最初からプログラムの形成に関与するということは、複雑なサプライチェーンをめぐる現実世界のエッジケースが早期に解決されることを意味します。これは、開始してから大規模にその限界が判明するプログラムに比べて、大きな利点です。
Linux コミュニティの実際の懸念
r/linux に対する反応は敵対的ではありませんでしたが、それ自体が IBM の発表としては注目に値します。深刻な批判のほとんどは 3 つのカテゴリーに分類されます。
AI が生成したパッチは上流で受け入れられるのに十分なものでしょうか?
AI モデルを使用して脆弱性を特定するのは簡単です。プロジェクトの既存のコード スタイルに一致し、テスト スイートに合格し、コード レビューに耐える修正を作成します。

過重労働で独断的なメンテナが実際に上流のプロジェクトにマージされる場合は、まったく別の問題です。
懸念されるのは、ボリュームのプレッシャーにより、技術的には CVE に対処するものの、基礎となるコードベースの保守が困難になるパッチの洪水になってしまうことです。 Red Hat の経験を持つ人々はこれに反発しました。20,000 人のエンジニアは経験豊富なオープンソース貢献者であり、上流のコミュニティがどのように運営されているかを知っています。記載されているモデルは、AI がトリアージと最初のパッチ生成を担当し、人間がレビューと貢献を担当します。 CVE をより速く処理するという商業的圧力の下でこの比率が維持されるかどうかは未解決の問題です。
「商用サブスクリプション」とは、有料の顧客が最初に修正プログラムを入手することを意味しますか?
商用サブスクリプションに関する発表の文言は、Lightwell に料金を支払っている企業がアップストリーム コミュニティよりも先にパッチ適用済みのパッケージを受け取り、悪用の余地が生まれるのではないかという懸念を即座に引き起こしました。
Red Hat の長年にわたるポリシーはアップストリーム第一であり、Red Hat に直接の経験を持つ複数の寄稿者がこの意見に強く反発しました。商用サブスクリプションには、バックポート、検証、SLA コミットメント、ライフサイクル管理が含まれます。セキュリティ修正への早期アクセスではありません。欧州サイバーレジリエンス法は、欧州で販売されるソフトウェアに対して上流優先の要件を法的に施行する予定であり、そのモデルからの逸脱をさらに抑制します。
メンテナがいないプロジェクトはどうなるのですか?
この質問は議論の中で十分な注目を集めませんでした。オープンソースの依存関係グラフの大部分は、転職した人、燃え尽きた人、または単純にやめた人によって維持されています。リポジトリに配信されるバックポートされたパッチは便利です。マージするアクティブな上流メンテナがいない修正では、断片化と長期にわたるメンテナンス負債が生じます。
IBM には n

オープンソース スタックの、放棄されているが広く依存している層については完全には対処できませんでした。それが現在の発表との大きなギャップだ。
そのタイミングは偶然ではありません。 IBM と Red Hat は、CVE を扱うすべてのセキュリティ チームに警鐘を鳴らしている同じ AI 脆弱性の波に注目しています。情報開示の量は増加しています。自動化された悪用はますます洗練されています。そして、脆弱性の公開と積極的な悪用との間の期間は縮小しています。
IBM は、企業がこれまで存在しなかった範囲でマネージド オープンソース セキュリティに料金を支払うようになるだろうと、計算の上で賭けています。 Red Hat は、モデルが Linux と OpenShift で機能することをすでに証明しています。 Lightwell は、エンタープライズ アプリケーションが取り込んでいるものの、誰も正式にメンテナンスしていないすべての独立したライブラリと AI フレームワークを含む、完全なアプリケーション依存関係ツリーにその賭けを拡張します。
より広範なオープンソース エコシステムに対する最終的な影響は、上流の貢献部分が実際にどのように機能するかに大きく依存します。 20,000 人のエンジニアが高品質のパッチを提供し、プロジェクトを共同保守し、小規模なメンテナーが単独で処理できない CVE フラッドに対処できるように支援している場合、IBM が最上位の商用検証レイヤーにいくら請求しているかに関係なく、コミュニティの利益は現実のものとなります。
ライトウェル計画は、より小規模で機能してきたモデルを使って真の危機に対処しています。 AI によって生成されたパッチの品質についての懐疑論は正当なものであり、プログラムが成熟するにつれ、IBM と Red Hat が直接回答するに値します。上流第一の取り組みは現実のものであり、法的に強化されています。
最も難しい問題は、そのテクノロジーが機能するかどうかではありません。重要なのは、ビジネスが拡大するにつれて、50 億ドルの商用プログラムがオープンソース コミュニティの関心と一致し続けることができるかどうかです。 Red Hat はこの点で他のほとんどの企業よりも優れた実績を持っています。

BM は、Red Hat エンジニアリングを分離し、独自の基準に従って運用することに取り組んでいます。
AI セキュリティの波は将来の問題ではありません。メンテナは現在、問題キューでそれを確認しています。何かを変えなければなりません。ライトウェル計画が正しいかどうかは、注意深く観察する価値がある。
詳細については、ibm.com/products/lightwell を参照してください。
IBMとRed HatがProject Lightwellに50億ドルを投入
IBM と Red Hat の 20,000 人のエンジニアがプログラムに取り組んでいます
オープンソース ソフトウェアだけで、Anthropic の Mythos Preview によって 3,900 件の高度または重大な脆弱性が発見されました
2024 年に 40,000 以上の CVE が公開される
IBM の推定によると、2026 年までに 59,000 の CVE が予測される
IBM は現在 62,000 以上のオープンソース パッケージを使用しており、10,000 以上に関する深い専門知識を備えています。
大手金融機関11社がすでに早期導入者として契約している
Ubuntu が AI 統合計画を発表したがコミュニティは不満
2008 年以来、あなたの毎日のドライバーとなってきたこのディストリビューションが、あなたが望まない方向に変わろうとしているかもしれません。いつ覚えておいてください[…]
AI が新たな Linux ゼロデイを発見し、セキュリティ研究者が激怒
コピーフェイル、ダーティ フラグ、フラグネシアのせいでセキュリティ サーカスは終わったと思っているなら、考え直してください。 Linux カーネルはさらに進化を遂げました […]
Linux サーバーを保護する方法
Linux サーバーは、企業がデータを保存および共有するための優れたコスト効率の高い方法です。 Linux はオープンソースであるため、多くの機能が提供されています […]
セキュリティとプライバシーに最適な Linux ディストリビューション
インターネットのセキュリティとプライバシーに関しては、Linux オペレーティング システム (ディストリビューション) が最良の選択肢です。しかし、非常に多くの Linux ディストリビューションがあるため […]
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
AI に関するリカルド氏、別の Linux ゼロデイを発見、セキュリティ研究者は安心

キングアウト
AI に関するメタが新たな Linux ゼロデイを発見し、セキュリティ研究者が慌てふためく
AI に関する KSaff が別の Linux ゼロデイを発見し、セキュリティ研究者が大慌て
Leoh on Framework の Ubuntu ラップトップは Windows を上回っていますが、そんなことは起こるはずがありません
David Collier-Brown 氏、Ubuntu について AI 統合計画を発表したばかりだが、コミュニティは満足していない
Flathub が AI に強硬線を引く: Vibe コード化されたアプリは許可されていない
AI がオープンソースの重大なバグを 3,900 個発見。 IBMはそれらを修正するために50億ドルを支払っている
カリフォルニア州で Linux が殺されそうになった後、誰かが実際にコードを読んだ
HP、デルとレノボに年間 10 万ドルの LVFS スポンサーシップで参加
AI が新たな Linux ゼロデイを発見し、セキュリティ研究者が激怒
スタンになろう！ニュースレターに参加する
あなたの電子メール アドレスは、ニュースレターと Linux Stans の活動に関する情報を送信するためにのみ使用されます。ニュースレターに含まれる購読解除リンクはいつでも使用できます。
前の投稿
カリフォルニア州で Linux が殺されそうになった後、誰かが実際にコードを読んだ
次の投稿
Flathub が AI に強硬線を引く: Vibe コード化されたアプリは許可されていない
Linux ファンによる万人向けのコンテンツ。
Linux Stans は、Linux Fans Stans によって作成および管理されている Web サイトです。リンに関連するあらゆる種類のユニークなコンテンツを読む

[切り捨てられた]

## Original Extract

40,000 CVEs published in 2024. 59,000 projected by 2026. IBM and Red Hat think they have a $5 billion answer to the open source security crisis. Here is what they are actually building.

IBM Just Committed $5 Billion to Fix Open Source Security. The Linux Community Has Complicated Feelings About It.
Skip to content
About
Write for Us – Submit an Article
Search
Search
Search …
Search
Search …
AI Found 3,900 Critical Open Source Bugs. IBM Is Paying $5 Billion to Fix Them
There is a number buried in IBM’s Project Lightwell announcement that deserves more attention than it is getting right now. Anthropic’s Mythos Preview AI model scanned open source software and identified nearly 3,900 high or critical-severity vulnerabilities. That is not the result of years of slow auditing. That is what one frontier AI model found in a preview run. And the model is only getting better.
That is the world IBM and Red Hat are building for. On May 28, 2026, the two companies announced a $5 billion commitment to Project Lightwell: a security clearinghouse for enterprise open source software, backed by 20,000 engineers and AI tooling designed to find and fix vulnerabilities before attackers can weaponize them. Banks are already signed up. The Linux community is watching very carefully.
The Problem Is Real and the Numbers Are Getting Ugly
More than 40,000 CVEs were published in 2024. IBM projects that number could climb to 59,000 by 2026. That acceleration is not happening because software is getting sloppier. It is happening because AI-driven vulnerability discovery is scaling in ways human security teams cannot match.
More than 90% of Fortune 500 companies run on open source software. Every one of those CVEs is a potential path into production systems at a bank, a hospital, a power grid. The software powering those environments is maintained, in many cases, by volunteers, hobbyists, and small teams operating without the budget or bandwidth to process hundreds of vulnerability reports a month while also shipping features and handling support.
The remediation gap, the distance between discovering a vulnerability and actually patching it across every affected production environment, is growing faster than any individual organization can close it on its own. That is the gap Project Lightwell is trying to fill.
What Project Lightwell Actually Does
Strip out the press release language and there are three concrete things happening here.
A Coordinated Security Clearinghouse
Enterprises can report sensitive vulnerabilities to IBM and Red Hat before public disclosure through a secure intermediary framework. IBM validates the issue and develops a fix without requiring access to the enterprise’s own application source code. The fix gets delivered to repositories the customer controls.
Then it goes upstream. The open source project gets the patch. That is the part that matters most for the broader ecosystem, and IBM has been explicit about it: the clearinghouse model is designed to strengthen upstream communities, not bypass them.
Backporting to What You Already Run
This is the piece most enterprise teams will actually care about. Project Lightwell does not tell organizations to upgrade their dependencies to get a security fix. It backports the fix to the exact versions they are already running in production.
If a company’s application is pinned to a specific Java library version from 2022, IBM patches that version. No forced upgrade. No compatibility risk. IBM works from dependency manifests like pom.xml and delivers signed, validated packages to repositories the customer controls. The initial focus is Maven and Java, with PyPI, npm, and Go on the roadmap.
AI-Assisted Engineering at Scale
IBM is deploying 20,000 engineers from Red Hat and IBM alongside advanced AI tooling. The AI handles high-volume vulnerability triage, prioritization, and initial patch development. The engineers review, shape, and ship what actually lands in upstream projects and customer environments.
IBM already uses more than 62,000 open source packages and maintains deep expertise across more than 10,000 of them. The reach covers Linux, Kubernetes, Java, Kafka, Ansible, Terraform, Flink, Cassandra, and more. Lightwell extends that model to the broader application dependency tree beyond Red Hat’s traditional product footprint.
The Early Adopter List Is Not a Joke
IBM announced that Bank of America, BNY, Citi, Goldman Sachs, JPMorganChase, Mastercard, Morgan Stanley, Royal Bank of Canada, State Street, Visa, and Wells Fargo are already collaborating on Project Lightwell. These organizations are not signing up because the press release was compelling. They are signing up because an unpatched vulnerability in a widely-used Java library is a regulatory and reputational catastrophe waiting to happen, and they have the budgets to pay for a managed solution.
Their involvement in shaping the program from the start means the real-world edge cases around complex supply chains will get worked out early. That is a meaningful advantage over a program that launches and then discovers its limitations at scale.
The Linux Community’s Actual Concerns
The reaction on r/linux was not hostile, which is itself notable for an IBM announcement. Most of the serious criticism fell into three categories.
Will the AI-Generated Patches Be Good Enough for Upstream Acceptance?
Identifying a vulnerability with an AI model is tractable. Writing a fix that matches a project’s existing code style, passes its test suite, survives code review from an overworked and opinionated maintainer, and actually gets merged into the upstream project is a different problem entirely.
The worry is that volume pressure turns this into a flood of patches that technically address a CVE but make the underlying codebase harder to maintain. People with Red Hat experience pushed back on this: the 20,000 engineers are experienced open source contributors who know how upstream communities operate. The stated model is AI for triage and initial patch generation, humans for review and contribution. Whether that ratio holds under commercial pressure to process CVEs faster is the open question.
Does “Commercial Subscription” Mean Paying Customers Get Fixes First?
The announcement language about commercial subscriptions triggered immediate concern that enterprises paying for Lightwell would receive patched packages before the upstream community does, creating a window for exploitation.
Red Hat’s long-standing policy is upstream first, and multiple contributors with direct Red Hat experience pushed back hard on this reading. The commercial subscription covers backporting, validation, SLA commitments, and lifecycle management. Not early access to security fixes. The European Cyber Resilience Act will legally enforce the upstream-first requirement for software sold in Europe anyway, which further constrains any drift from that model.
What Happens to Projects With No Maintainers?
This question did not get enough attention in the discussion. A significant portion of the open source dependency graph is maintained by people who have moved on, burned out, or simply stopped. A backported patch delivered to a repository is useful. A fix with no active upstream maintainer to merge it creates fragmentation and long-term maintenance debt.
IBM has not fully addressed the abandoned-but-widely-depended-upon layer of the open source stack. That is a real gap in the current announcement.
The timing is not accidental. IBM and Red Hat are watching the same AI vulnerability wave that is alarming every security team that handles CVEs. The volume of disclosures is climbing. The sophistication of automated exploitation is increasing. And the window between vulnerability disclosure and active exploitation is shrinking.
IBM is making a calculated bet that enterprises will pay for managed open source security at a scope that has not existed before. Red Hat already proved that model works for Linux and OpenShift. Lightwell extends that bet into the full application dependency tree, including all the independent libraries and AI frameworks that enterprise applications pull in but nobody officially maintains for them.
The net effect on the broader open source ecosystem depends heavily on how the upstream contribution piece actually plays out. If 20,000 engineers are contributing high-quality patches, co-maintaining projects, and helping small maintainers handle a CVE flood they cannot process alone, the community benefit is real regardless of what IBM charges for the commercial validation layer on top.
Project Lightwell is addressing a genuine crisis with a model that has worked at smaller scale. The skepticism about AI-generated patch quality is legitimate and deserves a direct answer from IBM and Red Hat as the program matures. The upstream-first commitment is real and legally reinforced.
The hardest question is not whether the technology can work. It is whether a $5 billion commercial program can stay aligned with open source community interests as the business scales. Red Hat has a better track record on that than most, and IBM has committed to keeping Red Hat engineering separate and operating by its own norms.
The AI security wave is not a future problem. Maintainers are seeing it in their issue queues right now. Something has to change. Whether Project Lightwell is the right something is worth watching closely.
More at ibm.com/products/lightwell .
$5 billion committed by IBM and Red Hat to Project Lightwell
20,000 engineers from IBM and Red Hat working on the program
3,900 high or critical vulnerabilities found by Anthropic’s Mythos Preview in open source software alone
40,000+ CVEs published in 2024
59,000 CVEs projected by 2026, per IBM estimates
62,000+ open source packages IBM currently uses, with deep expertise in 10,000+
11 major financial institutions already signed on as early adopters
Ubuntu Just Announced AI Integration Plans and the Community Is NOT Happy
The distro that’s been your daily driver since 2008 might be about to change in ways you really won’t like. Remember when […]
AI Just Found Another Linux Zero-Day and Security Researchers Are Freaking Out
If you thought the security circus was over after copyfail, dirty frag, and fragnesia, think again. The Linux kernel just took another […]
How to Secure Your Linux Server
Linux servers are a great, cost-effective way for businesses to store and share data. Linux is open-source, so it provides plenty of […]
Best Linux Distros for Security and Privacy
When it comes to internet security and privacy, Linux operating systems (distros) are the best option. But with so many Linux distros […]
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
Ricardo on AI Just Found Another Linux Zero-Day and Security Researchers Are Freaking Out
Meta on AI Just Found Another Linux Zero-Day and Security Researchers Are Freaking Out
KSaff on AI Just Found Another Linux Zero-Day and Security Researchers Are Freaking Out
Leoh on Framework’s Ubuntu Laptops Are Outselling Windows-And That’s Not Supposed to Happen
David Collier-Brown on Ubuntu Just Announced AI Integration Plans and the Community Is NOT Happy
Flathub Draws a Hard Line on AI: No Vibe-Coded Apps Allowed, Period
AI Found 3,900 Critical Open Source Bugs. IBM Is Paying $5 Billion to Fix Them
California Almost Killed Linux, Then Someone Actually Read the Code
HP Joins Dell and Lenovo in $100K Annual LVFS Sponsorship
AI Just Found Another Linux Zero-Day and Security Researchers Are Freaking Out
Become a Stan! Join our Newsletter
Your e-mail address is only used to send you our newsletter and information about the activities of Linux Stans. You can always use the unsubscribe link included in the newsletter.
Previous post
California Almost Killed Linux, Then Someone Actually Read the Code
Next post
Flathub Draws a Hard Line on AI: No Vibe-Coded Apps Allowed, Period
Content by Linux Fans for Everyone.
Linux Stans is a website created and maintained by Linux Fans Stans . Read all kinds of unique content related to Lin

[truncated]
