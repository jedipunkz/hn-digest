---
source: "https://www.theregister.com/ai-and-ml/2026/06/16/a-modest-proposal-reformat-everything-to-make-documents-more-palatable-to-ai/5255938"
hn_url: "https://news.ycombinator.com/item?id=48594009"
title: "A modest proposal: Reformat everything to make documents more palatable to AI"
article_title: "A modest proposal: Reformat everything to make documents more palatable to AI"
author: "gmays"
captured_at: "2026-06-19T04:37:15Z"
capture_tool: "hn-digest"
hn_id: 48594009
score: 2
comments: 1
posted_at: "2026-06-19T01:57:12Z"
tags:
  - hacker-news
  - translated
---

# A modest proposal: Reformat everything to make documents more palatable to AI

- HN: [48594009](https://news.ycombinator.com/item?id=48594009)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/06/16/a-modest-proposal-reformat-everything-to-make-documents-more-palatable-to-ai/5255938)
- Score: 2
- Comments: 1
- Posted: 2026-06-19T01:57:12Z

## Translation

タイトル: 控えめな提案: すべてを再フォーマットして、AI にとってより使いやすいドキュメントを作成する
説明: 何か

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
控えめな提案: すべてを再フォーマットして、AI にとってより使いやすい文書にする
Web サイトは AI モデルで利用できるように再設計されており、現在、連合はこの傾向をデジタル ドキュメントにも拡張したいと考えています。
Linux Foundation の下にある LF AI & Data Foundation は、企業がファイルを AI システムにフィードできるようにすることを目的とした AI フレンドリーなドキュメント形式である DocLang の開発を主導するワーキング グループを設立しました。
IBM、NVIDIA、Red Hat、ABBYY、HumanSignal、Forgis によって設立された DocLang グループは、PDF、Markdown、HTML、LaTeX などの既存の形式は AI ドキュメントの解析には適していないと主張しています。
2024 年後半、IBM は、Microsoft の MarkItDown または Marker プロジェクトと同様に、AI ドキュメントの解析を容易にする Docling と呼ばれるオープンソース ツールキットを開発しました。 Docling は、さまざまなファイル形式を構造化された AI 対応データに変換する方法を提供します。 DocLang は、さまざまなシステム間で構造化された出力を交換するための標準によってその基盤を拡張します。
「DocLangは、エンタープライズAIの根本的な問題の1つを解決するように設計されている。ドキュメントは機械ではなく人間のために作られているということだ」とAIオートメーション事業ABBYYのAI戦略担当副社長のMaxime Vermeir氏は声明で述べた。 「ドキュメントの構造、レイアウト、意味、ガバナンスに関する最小限で標準化された AI ネイティブの表現を導入することにより、DocLang は最新の AI システムのためのより決定的な基盤を構築します。」
仕様作成者らは、既存の形式はレンダリング用に設計されており、AI モデルがトークンに変換する際に意味情報、構造的関係、幾何学的なコンテキストが失われるため、新しい DocLang 形式が必要であると主張しています。仕様では次のことが説明されています。

Markdown には十分な範囲が欠けており、HTML は過度に冗長であり、LaTeX では多義性が許容されすぎています。
Cisco SD-WAN の make-me-root バグが攻撃を受けている
研究者によると、脱獄ではなく単に「このコードを修正してください」というプロンプトが表示されただけで、連邦当局は寓話5に激怒したとのこと
DARPA、将来のスターウォーズに役立つ交換可能な衛星を模索
Anthropic はクロードの潜水艦の ID を確認する権利を留保します
基本的に、DocLang は、DocLang 要素と LLM トークン間を 1 対 1 でマッピングするマークアップを通じて、LLM トークナイザー用に最適化されています。この仕様は、最適化されたプロンプトを生成するために LLM トークナイザーと連携する限られた XML 語彙に依存しています。ロスレスなので、AI 変換によって貴重な情報が失われることはありません。表、数式、グラフ、マルチモーダル コンテンツなどの一般的なグラフィック要素をサポートするように設計されています。そしてそれはオープンスタンダードです。
DocLang はコストの抑制にも役立ちます。 AI Cost Check によると、AI モデルに PDF の OCR スキャンを実行させるには、ベースラインとして約 1,200 の入力トークンと 150 の出力トークンが必要です。
これは企業の AI 顧客にとって 1 回限りでは重要ではありませんが、大規模になると注意が必要になります。また、AI モデルのトークンコストは大きく変動するため、企業は、特に文書が長くて複雑な場合、または高価なフロンティアモデルが使用されている場合、AI システムに PDF を取り込ませるために予想よりも多くの費用がかかっていることに気づく可能性があります。
「PDF は理解するためではなく、レンダリングするために設計されています」と、ABBYY の AI バリューおよびイネーブルメント リードの Jon Knisley 氏は The Register への電子メールで述べています。 「PDF が AI パイプラインに入るたびに、構造、意味、レイアウトが失われるため、最終的にはモデルの品質ではなくドキュメントの品質によってモデルの精度がボトルネックになってしまいます。チームはすべての統合ポイントでカスタム パーサーを構築することでそれを補っていますが、その結果、脆弱で 1 回限りの作業が発生し、

新しいドキュメントタイプごとにエンジニアリングスプリントを行います。」
Knisley 氏によれば、それには測定可能なコストがかかります。
「構造があいまいなため、モデルは推測に頼らざるを得ないため、幻覚のリスクが高まり、意味を抽出する代わりにレイアウトを解読するためにトークンが消費されてしまいます」と同氏は説明した。 「DocLang を使用すると、顧客は精度の向上、コストの削減、消費トークンの削減、パフォーマンスの向上、より一貫した出力を期待できます。正確な節約額はユースケースとドキュメントの複雑さによって異なりますが、私たちの初期ベンチマークでは、評価したモデルに応じて 4 倍から 30 倍以上のコスト削減が示されています。」
Knisley 氏はまた、文書が移動されると文書の出所データとメタデータが削除される可能性があることを指摘し、ガバナンスの利点についても言及しました。同氏によれば、DocLang はその情報を添付したままにしているという。
AI ドキュメント処理を提供する ABBYY は、DocLang ドキュメントを AI モデルにフィードすることでトークンを節約できる可能性を示すために、DocLang Interactive Benchmark を作成しました。たとえば、IBM の 2025 年年次報告書の PDF では、8,421 個の入力トークンと 512 個の出力トークンが生成されますが、DocLang バージョンでは 5,310 個の入力トークンと 498 個の出力トークンしか必要ありません。さらに、DocLang バージョンではレイテンシが低くなり (2.7 秒対 4.2 秒)、より高い品質が得られます (AI は 1 つのサブセクションを見逃し、PDF 内の表の結合を破壊しました)。
「まだ時期尚早であり、採用を誇張するつもりはない」とナイスリー氏は語った。 「この標準はオープンであり、自由に構築できるため、このグループはより多くのテクノロジープロバイダーや企業の参加を積極的に呼びかけています。初期の反応は心強いものであり、私たちは今後の展開について楽観視しています。」
オンプレミス
2,000 台の引退した Google Pixel スマートフォンがプライベート クラウドとして第二の人生を得る
このシステムには 2 キラピクセルの計算能力が詰め込まれていると言えるかもしれません。
ミッドジャーニーは、AI 画像生成から、患者が治療を行うボディ スキャン メディカル スパへと方向転換します。

「黄金の光」の中で
基礎となるテクノロジーは本物です...そして、会社が言及しなかったパートナーから借用したものです
ZTE と広東省中国電信がクロスベンダー IP ネットワーク シミュレーションのパイロットを推進し、インテリジェントなネットワーク運用への道を開く
パートナーコンテンツ: 95% を超えるデジタルツインの忠実度およびマルチベンダーのコラボレーションを活用して、ネットワーク変更のリスクを排除し、エラーゼロの O&M を達成します。
熱心な懐疑論者は、実際には最悪ではない新しい Amazon AI 製品に熱中していることに気づきました
エドからコーリーへのメモ: 安全な場合は 1 回まばたきをし、危険な場合は 2 回まばたきをしてください
PAAS と IAAS
Graviton 5 は感動的ですが、すべての神聖なるものへの愛を込めて、これを「AI チップ」と呼ぶのはやめてください
AWS は口よりもチップ工場の運営のほうが上手い
Citrix では、コストを意識した未公開株の投資家のように仮想デスクトップを実行できるようになりました
PC 価格の高騰により、ハードウェアの更新に代わる選択肢が興味深い
セキュリティ
研究者によると、脱獄ではなく単に「このコードを修正してください」というプロンプトが表示されただけで、連邦当局は寓話5に激怒したとのこと
オンプレミス
アマゾンは昨年自社のビット納屋で最大25億ガロンの水を使用している
科学
AI とブレイン コンピューター インターフェイスにより、言葉を話すことのできない ALS 患者がフルタイムで働けるようになります
公共部門
キャピタは公務員年金制度修正の期限を過ぎて航行しようとしている
仮想化
Tesco は急速な移行リスクにもかかわらず、VMware と Broadcom からの撤退に全力を注いでいます
基礎となるテクノロジーは本物です...そして、会社が言及しなかったパートナーから借用したものです
エドからコーリーへのメモ: 安全な場合は 1 回まばたきをし、危険な場合は 2 回まばたきをしてください
鳥ブランドのAIがストンキングスティングレイに乗ります
ハイブリッド システムはエッジでの効率向上をもたらす可能性がありますが、従来のインフラストラクチャではすぐには実現できません
患者の呼気をサンプリングすると、命と救急治療室のリソースが救われる可能性がある
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社は「W

あなたの情報をお持ちですか?』さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
フランスのデジタル主権推進はマイクロソフトの重力から逃れるのに苦戦している
Nextcloud のロールアウトは、ローカルで制御されるストレージが 1 つのことであることを示しています。ユーザーを Office から解放することはまったく別のことです
CentOS の歴史: 生化学者の Linux 趣味プロジェクトがどのようにして企業世界のデフォルトのオペレーティング システムになったのか
Red Hat が Windows が「おそらく適切な製品」であると述べた後、コミュニティが団結したとき
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
OpenBSD 7.9 が登場、あらゆる鋭いエッジを誇るダイヤモンドの原石
60 番目のリリースでは、その禁欲的な努力を失うことなく、より多くのコア、遅延休止状態、および基本的な Wi-Fi 6 が追加されています
Fedora: Microsoft は全力で取り組んでいるが、Deepin は見捨てられている
Red Hat の無料ディストリビューションはデスクトップを失うが、重要な新しい友人を作る
LocalSend はスニーカーネットを廃業にします
Apple ロックインを除いた AirDrop と同様
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

What

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
A modest proposal: Reformat everything to make documents more palatable to AI
Websites are being redesigned for consumption by AI models, and now a coalition wants to extend the trend to digital documents.
The LF AI & Data Foundation , under the Linux Foundation, has formed a working group to steer the development of DocLang , an AI-friendly document format that aims to help enterprises feed their files to AI systems.
The DocLang group, founded by IBM, NVIDIA, Red Hat, ABBYY, HumanSignal, and Forgis, contends that existing formats like PDF, Markdown, HTML, and LaTeX are ill-suited for AI document parsing.
In late 2024, IBM developed an open source toolkit called Docling to facilitate AI document parsing, not unlike Microsoft's MarkItDown or the Marker project. Docling provides a way to convert various file formats into structured AI-ready data. DocLang expands upon that foundation with a standard for exchanging structured output across different systems.
"DocLang is designed to solve one of the foundational problems in enterprise AI: documents were built for humans, not machines," said Maxime Vermeir, VP of AI Strategy at AI automation biz ABBYY in a statement. "By introducing a minimal, standardized, and AI-native representation of document structure, layout, meaning and governance, DocLang creates a far more deterministic foundation for modern AI systems."
The new DocLang format is necessary, the spec authors argue, because existing formats were designed for rendering and lose semantic information, structural relationships, or geometric context when AI models turn them into tokens. The specification explains that Markdown lacks sufficient scope, that HTML is excessively verbose, and that LaTeX allows too much ambiguity.
Cisco SD-WAN make-me-root bug under attack
Feds freaked over Fable 5 after simple 'fix this code' prompt, not jailbreak, says researcher
DARPA seeks swappable satellites to help with future star wars
Anthropic reserves right to check ID for Claude subs
Essentially, DocLang is optimized for LLM tokenizers through markup that maps between DocLang elements and LLM tokens on a 1-to-1 basis. The spec relies on a limited XML vocabulary that aligns with LLM tokenizers to produce optimized prompts. It is lossless, so the AI conversion doesn't do away with valuable info. It's designed to support common graphical elements like tables, formulas, charts, and multimodal content. And it's an open standard.
DocLang could also help keep costs under control. According to AI Cost Check , having an AI model conduct an OCR scan on a PDF requires about 1,200 input tokens and 150 output tokens as a baseline.
That's inconsequential to corporate AI customers on a one-off basis but demands attention at scale. And because AI models have highly variable token costs, companies may find they are spending more than they anticipated to have their AI system ingest PDFs, particularly if the documents are long and complicated or an expensive frontier model is used.
"PDFs were designed for rendering, not understanding," said Jon Knisley, AI Value and Enablement Lead at ABBYY, in an email to The Register . "Every time a PDF enters an AI pipeline, structure, meaning and layout get lost, so the model's accuracy ends up bottlenecked by document quality rather than model quality. Teams compensate by building custom parsers at every integration point, which results in brittle, one-off work, and a new engineering sprint for every new document type."
According to Knisley, that has measurable cost.
"Ambiguous structure forces the model into guesswork, which drives up hallucination risk and burns tokens deciphering layout instead of extracting meaning," he explained. "With DocLang, customers can expect better accuracy, lower costs, fewer tokens consumed, faster performance and more consistent outputs. The exact savings depend on the use case and document complexity, but our initial benchmarks show 4x to more than 30x lower cost depending on the model evaluated."
Knisley also cited governance advantages, noting that document provenance data and metadata can get stripped when documents gets moved. DocLang, he said, keeps that information attached.
ABBYY, which offers AI document processing, has created the DocLang Interactive Benchmark to illustrate the potential token savings of feeding DocLang documents to AI models. A PDF of IBM's 2025 annual report, for example, results 8,421 input tokens and 512 output tokens while a DocLang version requires only 5,310 input tokens and 498 output tokens. What's more, the DocLang version results in lower latency (2.7s vs 4.2s) and delivers better quality (the AI missed one subsection and mangled a table merger in the PDF).
"It's still early, and we won't overstate adoption," said Knisley. "The standard is open and free to build on, and the group is actively inviting more technology providers and enterprises to join. The early response has been encouraging, and we're optimistic about where it goes from here." ®
On-PREM
2,000 retired Google Pixel phones get a second life as a private cloud
You might say the system packs two kilapixels of compute
Midjourney pivots from AI image generation to body scanning medical spa where patients bathe in 'golden light'
The underlying technology is real...and borrowed from a partner the company failed to mention
ZTE and China Telecom Guangdong advance cross‑vendor IP network simulation pilots, paving the way for intelligent network operations
PARTNER CONTENT: Leveraging >95% digital twin fidelity and multi-vendor collaboration to eliminate network change risks and achieve zero-error O&M
Committed skeptic finds himself warming to new Amazon AI products that actually don't suck
Ed's note to Corey: Blink once if you're safe, twice if you're in danger
PAAS AND IAAS
Graviton 5 impresses, but please, for the love of all that's holy, stop calling them 'AI chips'
AWS better at running chip fabs than their mouths
Citrix now lets you run virtual desktops like a cost-conscious private equityeer
Soaring PC prices make alternatives to hardware refreshes interesting
security
Feds freaked over Fable 5 after simple 'fix this code' prompt, not jailbreak, says researcher
ON-PREM
Amazon owns up to using 2.5bn gallons of H2O in its bit barns last year
science
AI and brain-computer interface allow speechless ALS patient to work a full-time job
PUBLIC SECTOR
Capita is about to sail past deadline to fix civil service pensions scheme
virtualization
Tesco is sprinting to quit VMware and Broadcom despite rapid migration risks
The underlying technology is real...and borrowed from a partner the company failed to mention
Ed's note to Corey: Blink once if you're safe, twice if you're in danger
Bird-branded AI will ride on Stonking Stingray
Hybrid systems could bring efficiency gains at the edge, but conventional infrastructure isn't going anywhere fast
Sampling patients' breath may save lives and emergency room resources
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
France's digital sovereignty push is struggling to escape the Microsoft gravity well
Nextcloud rollout shows locally controlled storage is one thing; getting users off Office is quite another
History of CentOS: How a biochemist's Linux hobby project became the enterprise world's default operating system
When a community came together after Red Hat said Windows was 'probably the right product'
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
OpenBSD 7.9 arrives, a diamond in the rough proud of every sharp edge
Sixtieth release adds more cores, delayed hibernation, and basic Wi-Fi 6 without losing its ascetic streak
Fedora: Microsoft is all aboard, but Deepin is dumped
Red Hat’s free distro loses a desktop, but makes an important new friend
LocalSend puts your sneakernet out of business
Like AirDrop, minus the Apple lock-in
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
