---
source: "https://idp-software.com/news/the-76-percent-wall/"
hn_url: "https://news.ycombinator.com/item?id=48712212"
title: "Why frontier LLMs can't read the hard documents without experts involved"
article_title: "The 76% Wall: Where the Model Labs Stop Eating IDP"
author: "chelm"
captured_at: "2026-06-28T22:20:46Z"
capture_tool: "hn-digest"
hn_id: 48712212
score: 1
comments: 0
posted_at: "2026-06-28T22:08:25Z"
tags:
  - hacker-news
  - translated
---

# Why frontier LLMs can't read the hard documents without experts involved

- HN: [48712212](https://news.ycombinator.com/item?id=48712212)
- Source: [idp-software.com](https://idp-software.com/news/the-76-percent-wall/)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T22:08:25Z

## Translation

タイトル: なぜフロンティア LLM は専門家の関与なしに難しい文書を読むことができないのか
記事のタイトル: 76% の壁: モデル研究所が IDP を食べるのをやめる場所
説明: Gemini、GPT、および Claude は、2026 年 6 月にドキュメント抽出をコモディティ化しました。手書き、スパース テーブル、および規制により、依然としてすべてのフロンティア モデルの上限が 76% 近くに達しています。

記事本文:
メイン コンテンツにスキップ IDP-Software Vendor Finder について ニュースレターの投稿 IDP-Software Vendor Finder ベンダーの能力評価ガイド ニュース Vendor Finder ベンダーの概要
76%の壁: モデルラボがIDPを食べるのをやめる場所
公開日 2026-06-28 · 更新日 2026-06-28 · 8 分で読めます
価格がセントになった月
研究室は抽出物を販売していません。彼らは机を売っています。
同じ4週間以内に配管が開通した
購入する場合、これは何を意味しますか
価格がセントになった月
研究室は抽出物を販売していません。彼らは机を売っています。
同じ4週間以内に配管が開通した
購入する場合、これは何を意味しますか
Christopher Helm は、インテリジェント文書処理市場をカバーし、ベンダーの開発、テクノロジーのトレンド、企業の導入パターンを追跡しています。彼の分析は、AI を活用した自動化が金融サービス、ヘルスケア、保険、物流などの業界全体でドキュメント ワークフローをどのように再構築するかに焦点を当てています。
価格がセントになった月
2026 年 6 月、機械で文書を読むコストは管理する価値のある項目ではなくなりました。独立したコスト分析によると、Gemini Flash の直接抽出は 1,000 ページあたり約 0.17 ドルです。これに対し、AWS Textract の基本的な OCR は 1.50 ドル、Google 独自のレガシー Document AI Form Parser は 30 ドルです。標準的な請求書の束を抱えた開発者は、もはやインテリジェントな文書処理ベンダーを評価しません。彼らはすでに料金を支払ったモデルにファイルを貼り付け、返された JSON を読み取ります。
既存のクラウド プラットフォームは、独自のリリース ノートでその点を認めています。 Google の Document AI は現在、Gemini 3 Flash および Gemini 3 Pro を介してレイアウト パーサーをルーティングしており、画像とテーブルの注釈は 5 月 27 日に一般公開され、ID 用の 2022 年以前のすべてのレガシー プロセッサが提供されます。

税金、住宅ローン、および調達書類は、2026 年 6 月 30 日をもって廃止されました。10 年間にわたってクラウド OCR を定義してきたこの製品は、開発者が直接呼び出すことができる同じ一般モデルのラッパーになりました。このカテゴリを発明したプラットフォームがチャット モデルを優先して独自の専用プロセッサを廃止すると、専用レイヤーは終わります。
独立したベンチマークも一致しています。 Nanonets IDP Leaderboard は、9,000 を超える実際のドキュメントで実行されており、重要な情報抽出では Gemini 3 Flash が 91.1% でトップを占め、OCR、テーブル、ビジュアル質問応答では Gemini 3.1 Pro がトップとなっています。市場の簡単な半分（クリーンな請求書、印刷フォーム、構造化 PDF）にとって、モデルはパイプラインです。 5 月に説明したエージェント ホットフォルダーのループの価格はセント単位で、モデルの作成者によって出荷されます。
研究室は抽出物を販売していません。彼らは机を売っています。
OCR の価格改定は小さな話です。最も大きな問題は、フロンティアラボがチャットの売り込みをやめ、ナレッジワークデスクの売り込みを開始したことだ。ナレッジワークデスクとは、文書の多い専門的な仕事が行われているのと同じデスクである。
OpenAI は、カスタム GPT の後継となる Workspace Agent を 4 月 22 日に出荷しました。これには、Google Drive、SharePoint、Box、Salesforce に接続し、メモリを通じて機能を向上させるスケジュールされたエージェントが含まれています。この発表では、OpenAI 自身の社内利用、つまり 24,771 件の K-1 納税フォームが処理され、毎週のビジネス レポートが自動化されたことが引用されました。 2 日後にリリースされた GPT-5.5 は、1 年前であれば IDP ベンダーのカテゴリーで主張されていたであろう文言で位置づけられています。「ドキュメントとスプレッドシートを作成し、ソフトウェアを操作し、タスクが完了するまでツール間を移動する」というものです。そして 5 月 27 日、OpenAI と Thrive は、30 社以上の会計事務所のネットワークを活用して、自己改善型の税務 AI を実稼働環境に導入し、K-1 フォームの精度が 97% であると報告しました。

7,000往復のパイロットを通過します。あれはデモではありません。それは、大量の文書が規定され、自動化され、番号が付けられた専門的な作業です。
Anthropic も反対側から同じ動きをしています。 2 月 24 日に発売された Claude Cowork は、DocuSign 契約から重要な用語を読み取って抽出し、数式を使用してレシート画像から構造化データを Excel に抽出する「自律型デジタル コワークラ」として販売され、KYC ファイル スクリーニングとピッチブック作成用の 10 個の財務エージェント テンプレートが同梱されています。 Anthropic の金融サービス ポジショニングでは、複雑な Excel タスクで 83% の精度があり、Vals AI Finance Agent ベンチマークでリードしていると主張しています。サム・アルトマン氏の構図は明確だ。エージェントは「人間の上級従業員と同じように積極的に活動し、数日、数週間にわたるタスク」を処理すると信頼されている。
これは国内避難民中間層が織り込んでいない脅威だ。モデルラボは、より優れた抽出者になろうとしているわけではありません。彼らは、ドキュメントを読み取って推論し、下流の作業、つまり Coupa や UiPath のような垂直プラットフォームも必要とするエージェント ワークフロー層を実行するエージェントを所有しようとしています。抽出はそのエージェントの機能であり、それに付随する製品ではありません。
同じ4週間以内に配管が開通した
研究所が最下位を取得した一方、スタックの残りの部分は中央を標準化しました。ドキュメント形式自体は、いくつかのベンダーが堀として扱っていた独自の中間表現であり、6 月にオープンスタンダードになりました。
6 月 23 日、ABBYY は DocLang の輸出サポートを備えた FineReader Engine 12.8 を出荷し、Linux AI & Data Foundation は、ABBYY、IBM、NVIDIA、Red Hat、HumanSignal が共同設立したベンダー中立の AI ネイティブ ドキュメント標準として DocLang を発表しました。互いに競合する 5 社が自社製品のフォーマットを共同作成すると、そのフォーマットは 1 つのフォーマットではなくなります。

差別化を図り、インフラとなる。同じ週、IBMはDocling for watsonxを1,000ページあたり4ドルで一般提供し、同等のベンダーより20％安いと主張し、既にパイプライン実務者が説明するパイプラインの半分以下に収まっているオープンソースツールキットを包含した。 Docugami は 6 月 17 日に DGML をオープンソース化し、Inveniam パートナーシップと連携して、抽出されたデータをドキュメント レベルではなく個々のデータ ポイント レベルで検証可能にしました。
モデル層も同じスケジュールでコモディティ化されました。 Mistral は 6 月 23 日に OCR 4 をリリースしました。構造を認識し、単一コンテナで自己ホスト可能、1,000 ページあたり 2 ～ 4 ドル、170 言語にわたる OmniDocBench でのスコアは 93.07 でした。 Baidu は 6 月 22 日、MIT ライセンスに基づく 30 億パラメータの専門家混合モデルである Unlimited-OCR をオープンソース化しました。これらの発表がいずれも「信頼性」「検証可能」「AI ネイティブ」を売りにする理由は、フォーマットが標準化されつつある理由と同じです。つまり、抽出されたデータに作用するエージェントには、乱雑なドキュメントと確率モデルの間に決定論的な層が必要です。その層はまさに、私たちが 4 月に書いた精度と知識の問題であり、業界はそれを公に解決することを決定しました。
これは、ラボのマーケティングで省略されている部分であり、生計のために難しい文書を作成している人を安心させる部分です。フロンティアモデルは壁にぶつかりますが、その壁は測定可能です。
クリーンなドキュメントで Gemini の栄冠を獲得したのと同じ IDP リーダーボードでは、3 つのフロンティア ラボすべてを含む、テストしたすべてのモデルで手書き OCR が 75.5% に制限されていることがわかりました。ほとんどのモデルで、疎な非構造化テーブルのスコアは 55% 未満でした。チャート分析は特定の危険な方法で失敗しました。「軸の値が桁違いに読み取られ、間違ったバーが選択され、狭い間隔のデータ ポイントで 1 つずつずれたエラー」。これらは特殊なケースではありません。それらはbです

物流における船積みの煩わしさ、保険における手書きの請求書、医療における臨床記録や紹介状、決して簡単なものではなかった、規制され、一か八かの、変動の大きい文書。
そして、コスト調整後の精度においては、スペシャリストは依然として重要な部分で勝利を収めています。リーダーボードでは、Nanonets OCR-3 が全体の 85.9% でリードしており、Gemini 3.1 Pro の 83.2%、GPT-5.4 の 81.0% を上回っています。チャートの質問応答では、スペシャリストである Nanonets OCR2+ のスコアは GPT-5.4 の 77% に対して 87% で、1,000 ページあたり 10 ドルに対して 28 ドルです。専用モデルが、困難な課題においてフロンティア モデルを 4 分の 1 のコストで上回るということは、きれいなコモディティ化の物語が予測するものではありません。リーダーボード独自の結論は、すべての購入者が守るべきラインです。「記載されている問題はモデルの問題ではなく、アーキテクチャ、統合、および信頼の問題です。」
研究室自体は静かにタイムラインを遡っていきました。 5月下旬、アルトマン氏とダリオ・アモデイ氏はともに、最も積極的な自動化予測から撤退し、アルトマン氏は、初級レベルのホワイトカラーの仕事はすでになくなるという「間違いを喜んでいる」と述べた。 IDP のハードハーフに対するフロンティアの脅威は方向性があり、到達するものではありません。今は底が食べられています。頂点が約束されている。
ループがセントでフォーマットがオープンであれば、お金はそうでない両端に移動します。今四半期はまさにその方向に動きました。
さあ、エージェントオーケストレーションへ。 Coupa は、5 月 12 日の Rossum 買収に続き、9 日後に 2 回目の契約を締結し、ワークフローを取り込むスタートアップの Tonkean を買収して、吸収したばかりのドキュメント層のトップに座らせました。ダウン、生の容量へ：6月24日、PEの支援を受けたDaidaはScan-Opticsを買収した。Scan-Opticsは1968年に設立され、1日あたり400万ページを処理できる施設を運営している。

モデルにまったく依存していませんでした。
生き残った純粋なプレイは、サバイバルがどのようなものかを行っています。 Hyperscience は、Forrester が 2026 年第 2 四半期 Wave で「唯一の専用 IDP ピュアプレイ」と呼んだ唯一のベンダーであり、そこで最高の Current Offering スコアを獲得し、6 月 26 日には 2 年連続で IDP プラットフォーム オブ ザ イヤーに選ばれました。 Forrester 自身がそのレポートで購入者に警告しているのは、市場全体を一言で表したもので、MVP に達するまでに数週間ではなく、約 6 か月かかると予想しています。一方、UiPath は 5 月 28 日に史上初の GAAP ベースで黒字の四半期を発表しました。収益は 4 億 1,800 万ドル、GAAP 営業利益は 2,800 万ドルで、モデル ラボが下から押し上げているときにプラットフォームが見出す規律です。
購入する場合、これは何を意味しますか
最高のドキュメントではなく、最悪のドキュメントでテストしてください。フロンティアモデルはクリーンハーフで優れているので、そこで使用する必要があります。ベンダーを決定するのは、手書き、まばらなテーブル、および不良スキャンで何が起こるかです。これを含む誰かのリーダーボードを信じる前に、自分のコーパスで OCR と LLM のトレードオフを実行してください。
正直に簡単な半分の価格を設定します。ドキュメントが標準的で安定している場合、1,000 ページあたり数セントの直接モデル呼び出しは信頼できるアーキテクチャですが、同じ仕事に対する 6 桁のプラットフォーム契約は信頼できません。契約する前に、自分が市場のどの半分にいるのかを知ってください。
エージェントがどこに住んでいるのか尋ねてください。ラボ、Coupa、UiPath はすべて、ドキュメントに作用するエージェントを所有したいと考えています。抽出を購入する場合は、誰かのエージェントの機能を購入することになります。意図的に誰の代理人を決めるか。
ループではなく壁を購入してください。防御的に支払えるのは、モデルが読み取れない 24% に加えて、検証、監査証跡、例外キュー、および間違ったフィールドに関する説明責任です。それが本当のビジネスなのです。それはより小さくてより正直なものです

「文書を作成します」よりも。
コストの数値は独立した分析とベンダーのページからのものであり、方向性があります。長い複数ページの PDF で単純なフロンティア モデルを使用すると、1,000 ページあたり 100 ドル以上の費用がかかる可能性があるため、「安い」かどうかは呼び出し周りのアーキテクチャに完全に依存します。精度の数値は、公開されている IDP Leaderboard とラボ独自のベンチマークの主張からのものであり、ラボの主張はマーケティングとして解釈されるべきです。 6月のM&Aはプライマリリリースをソースとしています。当社は IDP ソフトウェアを販売しておらず、ここで指定されているベンダーから報酬を受け取っておりません。
きれいな請求書を読み取るモデルは、どんどん安くなり、性能も向上しています。手書きの請求フォームは関係ありません。
この投稿で引用されている情報源: Parsli LLM/OCR コスト分析 ; Google Document AI リリース ノート ; Nanonets IDP Leaderboard 1.5 とリーダーボードの詳細 ; OpenAI ワークスペース エージェント ; OpenAI GPT-5.5 ; OpenAI + Thrive Tax AI ;人間的なクロード・コワークと金融サービス担当のクロード。 Altman によるエージェント タスクと Altman/Amodei のウォークバックについて。アビー・ドクラング ; Watsonx 用の IBM Docling ;ドキュガミDGML ;ミストラルOCR4; Baidu Unlimited-OCR ;クーパがトンキーンを買収。 Daida が Scan-Optics を買収 ; Forrester Wave ドキュメント マイニング 2026 年第 2 四半期 ;ハイパーサイエンス IDP プラットフォーム オブ ザ イヤー ; UiP

[切り捨てられた]

## Original Extract

Gemini, GPT and Claude made document extraction a commodity in June 2026. Handwriting, sparse tables and regulation still cap every frontier model near 76%.

Skip to main content IDP-Software About Vendor Finder Contribute Newsletter IDP-Software Vendor Finder Vendors Capabilities Evaluations Guides News Vendor Finder Vendors Overview
The 76% Wall: Where the Model Labs Stop Eating IDP
Published 2026-06-28 · Updated 2026-06-28 · 8 min read
The month the price went to cents
The labs are not selling extraction. They are selling the desk.
The plumbing went open in the same four weeks
What this means if you are buying
The month the price went to cents
The labs are not selling extraction. They are selling the desk.
The plumbing went open in the same four weeks
What this means if you are buying
Christopher Helm covers the intelligent document processing market, tracking vendor developments, technology trends, and enterprise adoption patterns. His analysis focuses on how AI-powered automation reshapes document workflows across industries including financial services, healthcare, insurance, and logistics.
The month the price went to cents
In June 2026 the cost of reading a document with a machine stopped being a line item worth managing. Independent cost analysis puts direct Gemini Flash extraction at roughly $0.17 per 1,000 pages , against $1.50 for AWS Textract's basic OCR and $30 for Google's own legacy Document AI Form Parser. A developer with a stack of standard invoices no longer evaluates an intelligent document processing vendor. They paste the file into a model they already pay for and read the JSON that comes back.
The incumbent cloud platforms have conceded the point in their own release notes. Google 's Document AI now routes its Layout Parser through Gemini 3 Flash and Gemini 3 Pro , with image and table annotations reaching general availability on May 27, and every legacy pre-2022 processor for identity, tax, mortgage and procurement documents deprecated effective June 30, 2026. The product that defined cloud OCR for a decade is now a wrapper around the same general model a developer can call directly. When the platform that invented the category retires its own purpose-built processors in favor of a chat model, the purpose-built layer is over.
The independent benchmarks agree. The Nanonets IDP Leaderboard , run across more than 9,000 real documents, has Gemini 3 Flash leading key information extraction at 91.1% and Gemini 3.1 Pro on top for OCR, tables and visual question answering. For the easy half of the market (clean invoices, printed forms, structured PDFs), the model is the pipeline. The loop the agentic hotfolder post described in May is now priced in cents and shipped by the people who make the models.
The labs are not selling extraction. They are selling the desk.
The repricing of OCR is the small story. The large one is that the frontier labs have stopped pitching chat and started pitching the knowledge-work desk: the same desk where document-heavy professional work lives.
OpenAI shipped Workspace Agents on April 22, the successor to Custom GPTs, with scheduled agents that connect to Google Drive, SharePoint, Box and Salesforce and improve through memory. The launch cited OpenAI's own internal use: 24,771 K-1 tax forms processed, weekly business reports automated. GPT-5.5, released two days later , is positioned in language that would have been a category claim for an IDP vendor a year ago: "creating documents and spreadsheets, operating software, and moving across tools until a task is finished." Then on May 27, OpenAI and Thrive put a self-improving tax AI into production with a network of thirty-plus accounting firms, reporting 97% accuracy on K-1 forms across a 7,000-return pilot. That is not a demo. That is document-heavy regulated professional work, automated, with a number attached.
Anthropic is making the same move from the other side. Claude Cowork , launched February 24, is sold as an "autonomous digital colleague" that reads and extracts key terms from DocuSign agreements, pulls structured data out of receipt images into Excel with formulas, and ships with ten finance agent templates for KYC file screening and pitchbook building. Anthropic's financial-services positioning claims an 83% accuracy on complex Excel tasks and a lead on the Vals AI Finance Agent benchmark. Sam Altman's framing for where this goes is explicit: agents trusted to handle "multi-day and multi-week tasks, operating proactively much like a senior human employee."
This is the threat the IDP middle has not priced. The model labs are not trying to be a better extractor. They are trying to own the agent that reads the document, reasons over it, and does the downstream work, the agentic workflow layer that vertical platforms like Coupa and UiPath also want. Extraction is a feature of that agent, not a product alongside it.
The plumbing went open in the same four weeks
While the labs took the bottom, the rest of the stack standardized the middle. The document format itself, the proprietary intermediate representation that several vendors treated as a moat, became an open standard in June.
On June 23, ABBYY shipped FineReader Engine 12.8 with export support for DocLang , and the Linux AI & Data Foundation announced DocLang as a vendor-neutral AI-native document standard co-founded by ABBYY, IBM, NVIDIA, Red Hat and HumanSignal. When five companies that compete with each other co-author the format their products emit, the format has stopped being a differentiator and become infrastructure. The same week, IBM made Docling for watsonx generally available at $4 per 1,000 pages, claiming 20% below comparable vendors, wrapping the open-source toolkit that already sits under half the pipelines practitioners describe. Docugami open-sourced DGML on June 17, paired with an Inveniam partnership to make extracted data verifiable at the individual data-point level rather than the document level.
The model layer commoditized on the same schedule. Mistral released OCR 4 on June 23: structure-aware, self-hostable in a single container, $2 to $4 per 1,000 pages, scoring 93.07 on OmniDocBench across 170 languages. Baidu open-sourced Unlimited-OCR on June 22, a 3-billion-parameter mixture-of-experts model under an MIT license. The reason every one of these announcements sells "trust" and "verifiable" and "AI-native" is the same reason the format is being standardized at all: an agent acting on extracted data needs a deterministic layer between the messy document and the probabilistic model. That layer is exactly the accuracy and knowledge problem we wrote about in April, and the industry has decided to solve it in the open.
Here is the part the lab marketing leaves out, and the part that should reassure anyone running hard documents for a living. The frontier models hit a wall, and the wall is measurable.
The same IDP Leaderboard that crowns Gemini on clean documents found handwriting OCR capped at 75.5% across every model tested, all three frontier labs included. Sparse and unstructured tables scored below 55% for most models. Chart analysis failed in specific, dangerous ways: "axis values misread by orders of magnitude, the wrong bar selected, off-by-one errors on closely spaced data points." Those are not edge cases. Those are bills of lading in logistics, handwritten claims forms in insurance, clinical notes and referral forms in healthcare, the regulated, high-stakes, high-variance documents that were never the easy half.
And on cost-adjusted accuracy the specialists still win where it counts. The leaderboard has Nanonets OCR-3 leading overall at 85.9%, ahead of Gemini 3.1 Pro at 83.2% and GPT-5.4 at 81.0%. On chart question-answering, the specialist Nanonets OCR2+ scores 87% to GPT-5.4's 77%, at $10 per 1,000 pages against $28. A purpose-built model beating a frontier model on the hard task at a quarter of the cost is not what a clean commoditization story predicts. The leaderboard's own conclusion is the line every buyer should keep: "the problems described are not model problems, they are architecture, integration, and trust problems."
The labs themselves quietly walked the timeline back. In late May, both Altman and Dario Amodei retreated from their most aggressive automation predictions, with Altman saying he was "delighted to be wrong" that entry-level white-collar work would already be gone. The frontier threat to the hard half of IDP is directional, not arrived. The bottom is being eaten now. The top is being promised.
If the loop is cents and the format is open, the money moves to the two ends that are not. It moved in exactly those directions this quarter.
Up, into agentic orchestration. Coupa followed its May 12 Rossum acquisition with a second deal nine days later , buying the workflow-intake startup Tonkean to sit on top of the document layer it had just absorbed. Down, into raw capacity: on June 24, PE-backed Daida acquired Scan-Optics , a firm founded in 1968 that runs a facility capable of four million pages a day, a roll-up of the unglamorous scanning and BPO tier that never depended on a model at all.
The surviving pure-play is doing what survival looks like. Hyperscience was the only vendor Forrester called "the only dedicated IDP pure play" in its Q2 2026 Wave, where it took the highest Current Offering score, and on June 26 it was named IDP Platform of the Year for the second year running. Forrester's own warning to buyers in that report is the whole market in one sentence: expect about six months to reach an MVP, not a few weeks. Meanwhile UiPath posted its first-ever GAAP-profitable quarter on May 28: $418M in revenue, $28M GAAP operating income, the discipline a platform finds when the model labs are pressing up from underneath.
What this means if you are buying
Test on your worst documents, not your best. The frontier models are excellent on the clean half and you should use them there. The question that decides your vendor is what happens on the handwriting, the sparse tables, and the bad scans. Run the OCR-versus-LLM trade-off on your own corpus before you believe anyone's leaderboard, including this one.
Price the easy half honestly. If your documents are standard and stable, a direct model call at cents per thousand pages is credible architecture and a six-figure platform contract for the same job is not. Know which half of the market you are in before you sign.
Ask where the agent lives. The labs, Coupa and UiPath all want to own the agent that acts on the document. If you are buying extraction, you are buying a feature of someone's agent. Decide whose agent, on purpose.
Buy the wall, not the loop. What you can defensibly pay for is the 24% the models cannot read, plus the validation, the audit trail, the exception queue, and the accountability around a wrong field. That is a real business. It is a smaller and more honest one than "we do documents."
The cost figures are from independent analyses and vendor pages and are directional; naive frontier-model use on long multi-page PDFs can run $100 or more per 1,000 pages, so "cheap" depends entirely on the architecture around the call. The accuracy numbers are from the public IDP Leaderboard and the labs' own benchmark claims, and the labs' claims should be read as the marketing they are. The June M&A is sourced to primary releases. We do not sell IDP software and we are not paid by any vendor named here.
The model that reads the clean invoice keeps getting cheaper and better. The handwritten claims form does not care.
Sources cited in this post: Parsli LLM/OCR cost analysis ; Google Document AI release notes ; Nanonets IDP Leaderboard 1.5 and leaderboard details ; OpenAI Workspace Agents ; OpenAI GPT-5.5 ; OpenAI + Thrive tax AI ; Anthropic Claude Cowork and Claude for Financial Services ; Altman on agentic tasks and Altman/Amodei walkback ; ABBYY DocLang ; IBM Docling for watsonx ; Docugami DGML ; Mistral OCR 4 ; Baidu Unlimited-OCR ; Coupa acquires Tonkean ; Daida acquires Scan-Optics ; Forrester Wave Document Mining Q2 2026 ; Hyperscience IDP Platform of the Year ; UiP

[truncated]
