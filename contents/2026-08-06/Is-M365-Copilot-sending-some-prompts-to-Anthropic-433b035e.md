---
source: "https://thatrobot.ai/claude-residency-gap-m365-copilot/"
hn_url: "https://news.ycombinator.com/item?id=49190951"
title: "Is M365 Copilot sending some prompts to Anthropic?"
article_title: "Your Copilot lives in Australia. Your Claude does not. – That Robot"
author: "soundworlds"
captured_at: "2026-08-06T01:30:31Z"
capture_tool: "hn-digest"
hn_id: 49190951
score: 1
comments: 1
posted_at: "2026-08-06T00:39:11Z"
tags:
  - hacker-news
  - translated
---

# Is M365 Copilot sending some prompts to Anthropic?

- HN: [49190951](https://news.ycombinator.com/item?id=49190951)
- Source: [thatrobot.ai](https://thatrobot.ai/claude-residency-gap-m365-copilot/)
- Score: 1
- Comments: 1
- Posted: 2026-08-06T00:39:11Z

## Translation

タイトル: M365 副操縦士は Anthropic にプロンプ​​トを送信していますか?
記事のタイトル: あなたの副操縦士はオーストラリアに住んでいます。あなたのクロードはそうではありません。 – あのロボット

記事本文:
あなたの副操縦士はオーストラリアに住んでいます。あなたのクロードはそうではありません。
Microsoft は、2025 年後半にオーストラリア国内での Copilot 処理を利用できるようにしました。ほとんどの中堅企業はこの見出しを聞き、それに応じてプライバシー登録を更新しました。 Anthropic は、2026 年 3 月 10 日にアジア太平洋地域で 4 番目となるシドニー オフィスを開設し、オーストラリア企業からの最も一般的な要望としてデータ常駐を挙げました。どちらの動きも同じ方向を指します。
どちらも、Copilot セッションを Claude に切り替えた瞬間に開くギャップを埋めることはできません。
クロードに関する Microsoft Foundry のページでは、これについて異常に明確に説明されています。 「Anthropic (Microsoft ではありません) がデータの処理者です」とドキュメントには書かれています。プロンプトと出力は、顧客の地域外で処理される場合があります。 M365 副操縦士内のクロードは、明示的に EU データ境界の範囲外であり、同等のオーストラリアの保証もまだ存在しません。 Microsoft の主権の話に基づいて Copilot をテストし、その後同じブランドで Claude を有効にしたバイヤーにとって、これは調達デッキがフラグを立てるべき文です。
M365 の旗の下で何が変わったのか
この統合は現実的で便利です。 Anthropic は 2026 年 1 月に Microsoft のサブプロセッサーとなり、Foundry カタログには Microsoft のファーストパーティ モデルと並んで Claude Opus 4.6 が掲載され、エンタープライズ購入者はテナント インターフェイスを離れることなく特定の Copilot タスクを Claude にルーティングできるようになりました。マルチモデル導入を検討しているオーストラリアの組織にとって、これは信頼できる購入オプションです。
静かに変化したのはデータ パスです。 IDM Magazine の報道では、実際的な意味が明らかにされています。標準的な M365 Copilot インタラクションは、オーストラリア国境内で処理および保存できます。クロード経由でルーティングされたリクエストはすべて、そのコミットメントの対象外となります。ホストテナントは見た目が悪い

変わりました。データ フローが移動しました。
Anthropic 独自の地域コンプライアンス ページには、ヨーロッパの Microsoft Foundry が「2026 年に登場」と記載されています。オーストラリアはまだその段階にまったく入っていません。 2026 年 4 月の Microsoft Foundry フォーラムの Q&A スレッドでは、実際的な立場が明確にされています。現在、顧客が EU リージョンを選択している場合でも、推論の実行は Azure が運営するデータ センター内で完全に行われるわけではありません。同様のことがオーストラリアにも当てはまりますが、オーストラリアと地域の間に言及すべき約束がないという追加の事実もあります。
なぜギャップが来年ではなく今重要なのか
オーストラリアの 3 つのグループのバイヤーにとって、そのギャップは抽象的なものではありません。
1 つ目は APRA 規制対象事業体です。 APRA が業界に宛てた 2026 年 4 月 30 日付の書簡では、AI の 4 つの観察分野の 1 つとしてサプライヤーのリスクを挙げています。これは、マテリアル、サードパーティ、およびサードパーティの依存関係を完全にマッピングするようにボードに指示します。オフショアで処理された Anthropic 経由でプロンプトをサイレントにルーティングする Copilot の展開は、そのサプライ チェーン内に配置されます。規制当局が現在期待している方法でサプライヤー登録に表示されることはほとんどありません。
2 つ目は政府および SOCI 規制対象の団体です。 2026 年 3 月 23 日に発行された産業科学資源省のデータセンターおよび AI インフラストラクチャ開発者への期待では、ソブリン ホスティングと IRAP に準拠した調達がデフォルトとして定められています。ファウンドリ内でオーストラリア国外で加工されたクロードは、今日ではその基準を満たしていません。保護された分類のワークロードの場合はできません。
3 つ目は、重要な非公開顧客データを扱うプロフェッショナル サービス会社です。ここでの暴露は契約上のものであり、規制上のものではありません。過去 18 か月間に署名されたほとんどの契約書やマネージド サービス契約では、ツールとして Copilot が名指しされており、データはオーストラリアに保管されると想定されています。 M365 内のクロード パスによってそれが変更されるため、クライアントはそれを通知しない限りわかりません。

彼ら。
サプライヤーデッキが静かにスキップしているもの
M365 Copilot に関するベンダー説明会のほとんどは、EU データ境界の話とオーストラリア国内での処理に関する取り組みから始まります。どちらも本物で耐久性があります。どちらも自動的に Anthropic に拡張されるわけではありません。この資料には、ほとんどの場合、Microsoft が管理する推論の境界図が示されており、残りは購入者に推測してもらいます。
私たちが次のベンダー会議に出席するとしたら、次の 4 つの質問をテーブルに出すでしょう。
どの Copilot の機能とポリシーがプロンプトを Anthropic にルーティングしますか? それは現在のテナント トラフィックのどの部分を占めていますか?答えがゼロになることはほとんどなく、ほとんどの購入者は管理センター内からはわかりません。
Microsoft 管理のパスではなく、Anthropic パスの常駐義務は具体的にどのようなものですか?オーストラリアにとっての正直な答えは、現時点では「データは米国、ヨーロッパ、アジア、またはオーストラリアで処理され、米国に保存される可能性がある」です。 Anthropic 自身のプライバシー ページにはこれについて明確に記載されています。
Foundry と Copilot のどの構成で、機密分類の Anthropic パスを無効にしたり、ルーティングしたり、デフォルトで拒否したりできますか?構成は存在します。デフォルトでオンになることはほとんどありません。
Anthropic がオーストラリア地域で処理を提供するスケジュールと契約上の約束はどのようなものですか?また、出荷後にはどのような証拠が作成されますか? 「来る2026年」がウェブサイトに掲載されています。それは契約上の約束ではありません。
これらを書面で順番に質問することで、次回のサプライヤー審査開始時に取締役会が必要とする監査証跡が作成されます。
あのロボットが着地する場所
マルチモデル Copilot は健全なアーキテクチャ パターンです。クロードは強力な量産モデルです。 Anthropic のシドニー オフィスとローカル コンピューティングの探求は、オーストラリア市場にとって正しい方向です。そのどれにも異論はありません。
議論されているのはそのアイデアです

Microsoft の主権境界では、M365 サーフェスの下のすべてのモデルがカバーされます。そうではありません。このギャップは狭く、修正可能ですが、購入者のみが行うことができます。サプライヤーが独自に調達することはありません。彼らは、そのエッジケースをマッピングすることではなく、統合を販売するように求められています。
SOCI および NIST AI RMF に基づく主権についてすでに真剣に考えている組織にとって、アクション リストは短いものです。今日はルートを監査してみましょう。ポリシーではなく構成によって機密分類に対する人為的パスを制限します。次回の APRA、OAIC、または ASIC の作業の前に、欠落しているサプライヤー登録エントリを追加してください。地域の製品の更新を待つのではなく、質問をしたらサプライヤーが同意する契約上の約束を構築します。
主権の話だけで Copilot に同意した組織の場合、アクション リストはより短くなります。停止。ルートを地図にします。他の人が決定を下す前に、どのワークロードが今日人間の道に進むことができ、どのワークロードがそうではないのかを決定してください。 EU GPAI 実践規範は、国境を越えた購入者に対しても同じ方向性を示しています。地元の同等品が登場しつつあります。
M365 内のクロード パスは、データが実際にどこに行くのかという問題に関して、オーストラリアのバイヤーが 2 年間で行った最も有用なテスト ケースです。その答えは、サプライヤーデッキが示唆する答えではありません。質問をすれば、修正は簡単です。
当社は、オーストラリアの規制対象企業内で Copilot、Foundry、および Claude のデータ フローをマッピングし、監査対象のアーキテクチャを防御するために必要なサプライヤーのアンケート、技術管理マップ、および証拠パックを作成しました。当社が主権を実現するためにどのように構築するかをご覧ください。今週中にどの構成がギャップを埋めるのか、またどの構成がギャップを埋める前にサプライヤーからの書面によるコミットメントが必要なのかを説明します。
ポール・ウィルソン
あのロボの創始者

t. Microsoft、SAP、Deloitte、NRI で 27 年間データと AI システムを構築してきました。実際に動くAIについて書いています。
この柱について詳しく読む: 主権と安全 、または AI プライバシーの執行が次にどこに到達するかを読んでください。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。

## Original Extract

Your Copilot lives in Australia. Your Claude does not.
Microsoft made in-country Copilot processing available in Australia in late 2025. Most mid-market organisations heard the headline and updated their privacy register accordingly. Anthropic opened its Sydney office on 10 March 2026 , the fourth in Asia-Pacific, and named data residency as the single most common request from Australian enterprises. Both moves point in the same direction.
Neither closes the gap that opens the moment you switch a Copilot session to Claude.
The Microsoft Foundry page on Claude is unusually clear about it. “Anthropic (not Microsoft) is the processor of the data,” the documentation reads. Prompts and outputs may be processed outside the customer’s region. Claude inside M365 Copilot is explicitly out of scope for the EU Data Boundary, and the equivalent Australian assurance does not yet exist either. For a buyer who tested Copilot under the Microsoft sovereignty story and then enabled Claude under the same brand, that is the sentence the procurement deck should have flagged.
What changed under the M365 banner
The integration is real and useful. Anthropic became a Microsoft sub-processor in January 2026, the Foundry catalogue now lists Claude Opus 4.6 alongside Microsoft’s first-party models , and enterprise buyers can route specific Copilot tasks to Claude without leaving their tenant interface. For Australian organisations standing up multi-model deployments, this is a credible buying option.
What changed quietly is the data path. IDM Magazine’s coverage sets out the practical implication. Standard M365 Copilot interactions can be processed and stored within Australian borders. Any request routed through Claude falls outside that commitment. The host tenant looks unchanged. The data flow has moved.
Anthropic’s own regional compliance page lists Microsoft Foundry in Europe as “Coming 2026” . Australia is not yet on that page at all. The Q&A thread on the Microsoft Foundry forum from April 2026 makes the practical position explicit: even where the customer selects an EU region today, inference execution does not occur fully inside Azure-operated data centres . The same is true for Australia, with the additional fact that there is no Australian-region commitment to point to.
Why the gap matters now, not next year
For three groups of Australian buyers, the gap is not abstract.
The first is APRA-regulated entities. APRA’s 30 April 2026 letter to industry names supplier risk as one of four AI observation areas. It tells boards to map material, third-party and fourth-party dependencies in full. A Copilot deployment that silently routes prompts through Anthropic, processed offshore, sits inside that supply chain. It rarely appears in the supplier register the way the regulator now expects.
The second is government and SOCI-regulated entities. The Department of Industry, Science and Resources Expectations for data centres and AI infrastructure developers, published 23 March 2026, frame sovereign hosting and IRAP-aligned procurement as a default. Claude inside Foundry, processed outside Australia, does not meet that bar today. For Protected-classification workloads it cannot.
The third is professional services firms handling material non-public client data. The exposure here is contractual, not regulatory. Most engagement letters and managed-service deals signed in the last 18 months name Copilot as the tool and assume the data stays in Australia. The Claude path inside M365 changes that, and the client will not know unless you tell them.
What the supplier deck quietly skips
Most vendor briefings on M365 Copilot lead with the EU Data Boundary story and the Australian in-country processing commitment. Both are real and durable. Neither extends automatically to Anthropic. The deck almost always shows the boundary diagram for Microsoft-managed inference and lets the buyer infer the rest.
If we were sitting in the next vendor meeting, the four questions we would put on the table are these:
Which Copilot features and policies route prompts to Anthropic, and what fraction of our tenant traffic does that represent today? The answer is rarely zero, and most buyers cannot tell from inside the admin centre.
What is the residency commitment for the Anthropic path specifically, not the Microsoft-managed path? The honest answer for Australia is currently “data may be processed in the US, Europe, Asia or Australia, and stored in the US”. Anthropic’s own privacy page is explicit on this .
Which configurations in Foundry and Copilot let us disable, route, or default-deny the Anthropic path for sensitive classifications? The configuration exists. It is rarely on by default.
What is the timeline and contractual commitment for Anthropic to deliver Australian-region processing, and what evidence will be produced once it ships? “Coming 2026” is on the website. It is not a contractual commitment.
Asking these in order, in writing, produces the audit trail a board will need when the next supplier review opens.
Where That Robot lands on this
Multi-model Copilot is a sound architectural pattern. Claude is a strong production model. Anthropic’s Sydney office and exploration of local compute are the right direction of travel for the Australian market. None of that is in dispute.
What is in dispute is the idea that the Microsoft sovereignty boundary covers every model under the M365 surface. It does not. The gap is narrow and fixable, but only by the buyer. Suppliers will not raise it on their own. They have been asked to sell the integration, not to map its edge cases.
For organisations that already think hard about sovereignty under SOCI and NIST AI RMF , the action list is short. Audit the routes today. Restrict the Anthropic path for sensitive classifications by configuration, not policy. Add the missing supplier register entries before the next APRA, OAIC, or ASIC engagement. Build the contractual commitment your supplier will agree to once you ask the question, rather than waiting for a regional product update.
For organisations that signed off on Copilot on the sovereignty story alone, the action list is shorter. Stop. Map the routes. Decide which workloads can sit on the Anthropic path today, and which cannot, before someone else makes the decision for you. The EU GPAI Code of Practice signals the same direction for cross-border buyers. The local equivalent is on its way.
The Claude path inside M365 is the most useful test case Australian buyers have had in two years for the question of where their data actually goes. The answer is not the answer the supplier deck implies. The fix is straightforward once the question is asked.
We have mapped Copilot, Foundry and Claude data flows inside Australian regulated businesses and produced the supplier questionnaire, technical control map, and evidence pack required to defend the architecture under audit. See how we build for sovereignty and we will walk you through which configurations close the gap inside the week, and which need a written commitment from your supplier before they close at all.
Paul Wilson
Founder of That Robot. 27 years building data and AI systems at Microsoft, SAP, Deloitte and NRI. Writing about AI that actually works.
Read more on this pillar: Sovereignty and Safety , or read where AI privacy enforcement is landing next .
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
