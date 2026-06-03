---
source: "https://lyfe.ninja/news/#we-practice-what-we-preach"
hn_url: "https://news.ycombinator.com/item?id=48369939"
title: "Practicing What We Preach – Validating Our AI Assistant Content"
article_title: "Insights for ML Encoding & Digital Signature for AI Content Verification | lyfe.ninja"
author: "lyfeninja"
captured_at: "2026-06-03T00:46:38Z"
capture_tool: "hn-digest"
hn_id: 48369939
score: 2
comments: 0
posted_at: "2026-06-02T13:21:03Z"
tags:
  - hacker-news
  - translated
---

# Practicing What We Preach – Validating Our AI Assistant Content

- HN: [48369939](https://news.ycombinator.com/item?id=48369939)
- Source: [lyfe.ninja](https://lyfe.ninja/news/#we-practice-what-we-preach)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:21:03Z

## Translation

タイトル: 私たちが説いていることの実践 – AI アシスタントのコンテンツの検証
記事のタイトル: AI コンテンツ検証のための ML エンコーディングとデジタル署名に関する洞察 |ライフ忍者
説明: 機械学習のエンコードとデジタル署名、AI コンテンツの証明、検証、出所、信頼できるデータ システムへの新しいアプローチに関する最新情報と洞察。

記事本文:
-->
ホーム
会社名
について
お問い合わせ
私たちと一緒に働きましょう
ニュース
プロジェクト
ソリューション
ブロックシール
-->
サインイン
-->
ホーム
について
お問い合わせ
ニュース
コンサルティング -->
プロジェクト
私たちと一緒に働きましょう
サインイン
-->
-->
-->
lyfe.ninja AI アシスタント
lyfe.ninja の核となる信念の 1 つは、信頼は検証可能であるべきだということです。
コンテンツ検証について話すのは簡単です。毎日使用するシステムにそれを組み込むのは困難です。
過去数か月間、私たちは今後のコンテンツ署名および検証プラットフォームの背後にあるインフラストラクチャを開発してきました。これには、実稼働署名サービス、Python SDK、ブラウザ側の JavaScript および CSS ヘルパー、コンテンツの署名と独立した検証に必要なサポート システムが含まれます。
最近、私たちは重要な一歩を踏み出しました。私たちは同じインフラストラクチャを独自の AI アシスタントに直接統合しました。
すべてのアシスタントの応答は、ユーザーに配信される前に署名されるようになりました。検証は、外部開発者が最終的にアクセスできるのと同じサービスとライブラリを使用して独立して行われます。
この統合は非常に貴重なものでした。これにより、実際の状況下でプラットフォームをテストし、開発者のエクスペリエンスを検証し、広く利用可能になる前に改善点を特定することができます。
さらに重要なことは、他の人に期待するのと同じ基準を確実に維持できるようにすることです。 AI が生成したコンテンツは検証可能であるべきだと考えるのであれば、私たち自身の AI システムも検証可能であるべきです。
統合では現在以下を使用しています。
弊社の実稼働署名および検証サービス
ブラウザサイドの JavaScript ヘルパー
検証 UI コンポーネントとスタイリング ヘルパー
まだ初期段階ではありますが、このマイルストーンは、コンテンツ検証を現実世界のアプリケーションに簡単に統合できるようにするための重要なステップを表しています。まだやるべきことはありますが、基礎は整いつつあります。
チェックしてください

右上のチャットアイコンをクリックしてチャットを開始します。ブラウザ開発ツールを使用して AI によって生成されたコンテンツを改ざんして、何が起こるかを確認してください。
これを「エージェントを知る (KYA)」というジレンマに対する私たちの見解を考えてみましょう。当社の BlkBolt™ 技術で作成された署名は、この問題を独自に解決できることがわかりました。
AI エージェントは急速に進化しています。新しいモデル。新しいバージョン。新しい行動。ベンチマークが変化し、システムが更新され、先月導入したものが突然まったく同じではなくなります。
一般的なマントラは次のようです...「正しく見えます...おそらく大丈夫です。」
リスク管理とサイバーセキュリティの背景があるので、それは楽観的だと言えます。
私たちは少し異なるアプローチを検討してきました。
エージェントを信頼するだけでなく、エージェントが実際に何を生成したかを確認してください。
私たちはストリーミング システムでこのアイデアを検討する小さなデモを作成しました。
中心となる概念はシンプルです。
LLM/エージェントが応答を生成する
最終応答は作成時にソース システムによって署名されます。
クライアントは信頼する前に検証できます
なぜこれが重要なのか（特に今）
急速に変化する AI モデルとエージェントの世界では:
どのエージェントまたはバージョンが応答を生成したかをどのようにして知ることができますか?
応答が転送中に変更されていないことをどのように確認しますか?
過去の成果物を放置したくない場合はどうすればよいでしょうか?
実際にユーザーに表示された内容をどのように証明しますか?
2 つのシステムが同じものを生成しましたが、どちらが先ですか?
以前に発行された署名は無効になります
そのバージョンからの出力の信頼をやめることができます
あなたは、これまでに生み出されたものすべてに永久に結びついているわけではありません
🛒 エージェントコマース - AI エージェントがお客様に代わって購入を行います。
あなたのエージェントは実際にその取引を開始しましたか?
リクエストは途中で変更されましたか?
企業は同じことを検証できますか?
👶 教師エージェント - A ch

ild は AI 講師と対話しています。
表示されるコンテンツはシステムが生成したものとまったく同じですか?
ブラウザ内、拡張機能、または転送中に何か変更されましたか?
実際に表示されたものを監査できますか?
🏦 一か八かの顧客サービス - 銀行代理店が返金を行ったり、財務上のガイダンスを提供したりします。
ユーザーに提示または言われた内容を証明できますか?
後で監査してもらえますか？ (誰が、何を、どこで、いつ)
欠陥のあるバージョンからの出力の信頼を取り消すことはできますか?
これがどのように機能するか (本番環境)
目標は、非常にシンプルにすることです。
高いレベルで:
BlkBolt™ エンコーディング モデルを所有し、AI エージェントにリースします。
リースされたモデルは署名と検証に使用されます
エージェント システムは、コンテンツの作成時に 1 回の署名呼び出しを行います。
応答は署名アーティファクトとともに送信されます
クライアントは別の API 呼び出しを通じてバックグラウンドで自動的に検証します。
取り消し可能性という重要な差別化要因に加えて、BlkBolt™ で生成された署名には、次のような追加の利点もあります。
メタデータによる埋め込みポリシー
RSA、ECDSA などの従来の署名と比較して、従来のキー処理は不要です。
分散検証 - 署名がどうなるかを制御するのはあなたですが、私たちは署名の検証を支援するだけです
接続レベルの信頼 (TLS など) だけではなく、コンテンツ レベルの信頼
まだ初期で少し粗末ですが、核となるアイデアは機能します。エージェントがユーザーに代わって行動し、システムがより自律的になり、マルチエージェント システムが標準になり、信頼が「アプリ」から「出力」に移行するにつれて、これはより重要になると私たちは考えています。
あなたがどう思うか知りたいです!
このようなものは実際にどこで役立つのでしょうか?
やりすぎ？それともスタンダードになるものでしょうか？
まずこれが欲しい場所は何ですか？
フィードバックを共有してみませんか?お問い合わせ ！
デジタル署名はどこにでもあります。ソフトウェアアップデートから

金融取引に関しては、現代のシステムにおける真正性と完全性を検証する方法のバックボーンを形成しています。
何十年にもわたって、その基盤は RSA や ECDSA などの暗号化スキーム、確立された数学的原理と公開/秘密キー インフラストラクチャに基づいたアプローチに基づいて構築されてきました。
lyfe.ninja では、別の方向性を模索してきました。私たちは、BlkBolt™ のエンコード強度をより適切に評価するための研究の機会を積極的に追求していますが、実用的で即時の使用例としてデジタル署名にますます注目が集まっています。このことから、次のような単純な疑問が生まれました。
データは固定長にハッシュ化されます
信頼は難しい数学的問題に根ざしています
BlkBolt™ は異なるアプローチを採用しています。キーの代わりに、トレーニングされたモデルを使用してデータをエンコードおよび検証します。その結果、キーではなく入力のモデルベースの表現に関連付けられた署名アーティファクトが得られます。
大まかに言えば、どちらのアプローチも同じ中心的な目標を達成します。
整合性 (データは変更されていない)
信頼性 (予想されるソースからのものである)
最も重要な違い: 取り消し可能性
従来のデジタル署名は設計上永続的です。
無限に検証できる
たとえ秘密鍵が後で破棄されたとしても
失効は外部 (CRL、OCSP) で処理され、常に強制されるわけではありません
BlkBolt™ は別のモデルを導入します。署名の有効性は、発行後にアクティブに制御できます。
モデルを破棄 → 署名を検証できなくなります
アクセスの有効期限が切れる → 署名の検証が失敗する (これはリースを通じて行われます)
埋め込みメタデータ (リース、タイムスタンプなど) を使用してポリシーを適用 → 署名が条件付きで有効になります
現実の世界では、信頼が永続的なものであることはほとんどありません。
書類は受理されても、後で不正であることが判明する可能性があります。
発言や支持は行っても、後で撤回することができます。
アクセスしてくださいw

ある時点では有効だったとしても、もはや適切ではなくなる可能性があります。
キーを超えて: システムレベルの変化
秘密鍵を安全に保管
署名権限は一意のモデル インスタンスに関連付けられています
検証を厳密に制御できる
重要なコンポーネント (エンコーダやデコーダなど) は分離されたままにすることができます
敏感なコンポーネントの露出を制限する
管理された検証環境
明確にしておきますが、これは RSA や ECDSA を置き換えるものではありません。従来のデジタル署名は実証され、標準化され、広く信頼されており、相互運用可能です。
BlkBolt™ は、従来のシグネチャがあまり適していない状況に対して、別のアプローチを模索しているだけです。このアプローチは、次のようなシナリオで興味深いものになります。
一時的なシステムまたはセッションベースのシステム
永久有効性が義務となる状況
従来の署名は、「これは署名されましたか?」と答えます。
BlkBolt™ 署名は、「これはまだ有効ですか?」という扉を開きます。
私たちは、特に正式なセキュリティの観点から、BlkBolt のプロパティを評価する段階ではまだ初期段階にあります。
アイデアをプレッシャーテストすることに意欲的な人々
これらのカテゴリのいずれかに該当する場合は、ぜひご連絡をお待ちしております。お問い合わせページからご連絡いただくか、Work With Us フォームにご記入ください。
lyfe.ninja は、カーネギーメロン大学の CyLab Venture Network に参加したことを発表できることを嬉しく思います。
過去 18 か月間にわたり、lyfe.ninja は、分離と長期的な回復力に焦点を当てたデータ エンコーディングに対するニューラル ネットワーク ベースのアプローチを探求する研究主導のデータ保護システムである BlkBolt™ を開発してきました。 Venture Network は、このテクノロジーの実世界への応用を継続的に改良しながら、サイバーセキュリティ研究者、創設者、実践者のより広範なコミュニティと関わる機会を提供します。
CyLab Venture Network への参加は、lyfe.ninja の取り組みを反映しています。

ディープラーニングと次世代データ保護の交差点における素晴らしい技術対話、外部検証、コラボレーション。
私たちは、来年も CyLab コミュニティと関わることを楽しみにしています。
当社は、当社のコアテクノロジーである BlkBolt™ の 1 つの実験的ユースケースを紹介する新しいインタラクティブなデモを開始しました。これは、デジタル ファイル用の取り消し可能なモデル検証済み署名です。
デモでは次のことが可能です。
リースされた分離モデルを使用してファイルに署名する
信頼と検証がどのように変化するかをリアルタイムで確認する
設計上取り消し可能な署名
有効期間の長いキーではなく、分離されたモデルに関連付けられた検証
データの管理を第三者に渡すことなく信頼性を確保
ファーストパーティのコンテンツ検証 - 元のファイルを公開または保存せずに、コンテンツを作成したことを証明します。
サードパーティの承認と検証 - 組織、ブランド、または権威者が他の人が作成したコンテンツを承認または検証できるようにし、必要に応じて後でその承認を取り消すことができます。
AI コンテンツ認証 - ウォーターマークの脆弱性、一元化されたレジストリ、または取り消し不能な申し立てを回避しながら、コンテンツが AI によって生成されたものであること、どのモデルが生成したのか、そのモデルが誰に属しているのかを証明します。
デジタル資産の整合性と来歴 - ファイルまたはデジタル資産が時間の経過とともに変更されていないことを確認し、資産に関する重要なメタデータを提供し、信頼が適用されなくなった場合に明確に明らかにします。
本日、私たちは lyfe.ninja の大きなマイルストーンを公に共有できることを嬉しく思います。それは、「ランダム化トークン化と潜在空間表現を使用したニューラル ネットワーク ベースのデータ エンコーディングのシステムおよび方法」というタイトルの仮特許の提出です。
この申請書は、当社の長期にわたる研究イニシアチブであるプロジェクト BlackBox から生まれた新しいデータ保護テクノロジーである BlkBolt™ の基礎を表しています。
なぜB

lkBolt™ は存在します
このアプローチの違い
データは数学的に暗号化されるのではなく、エンコードされます
解釈には、再利用可能なキーではなく、特定のトレーニング済みモデルへのアクセスが必要です
各展開は分離されており、システム上のリスクが軽減されます
モデルは進化および回転できるため、システムの改善に応じて保護されたデータを再エンコードできるようになります。
今年も終わりに近づいてきたので、感謝の意を表し、ちょっとした最新情報を共有したいと思います。
過去 1 年間、私たちはコア テクノロジーとそれをサポートするインフラストラクチャにおいて、密かに大きな進歩を遂げてきました。反復、テスト、破壊、修正、迅速な学習など、この作業の多くは舞台裏で行われましたが、それが私たちが心から誇りに思う基盤を築きました。
今シーズンは、lyfe.ninja がビジネスになってからほぼ 1 年になりますが、それでも少し現実離れした感じがします。アイデアとして始まったものは、粘り強さ、好奇心、そしてたくさんの夜更かしによって形作られ、現実のものへと成長しました。
まだすべての詳細を共有する準備ができていません…しかし、重大な発表が間もなく行われることをお知らせできることを嬉しく思います。大きな出来事が動き始めており、次の章は重要なものになりつつあります。
今のところ、私たちは進歩、教訓、そしてその過程でのサポートに感謝しています。暖かく穏やかなホリデーシーズンを迎え、心躍る新年のスタートをお祈りいたします。
🎄✨
— ライフ

[切り捨てられた]

## Original Extract

Updates and insights on machine learning encoding and digital signatures, AI content attestation, verification, provenance, and emerging approaches to trusted data systems.

-->
Home
Company
About
Contact
Work With Us
News
Projects
Solutions
BlkSeal
-->
Sign In
-->
Home
About
Contact
News
Consulting -->
Projects
Work With Us
Sign In
-->
-->
-->
lyfe.ninja AI Assistant
One of our core beliefs at lyfe.ninja is that trust should be verifiable .
It's easy to talk about content verification. It's harder to build it into the systems you use every day.
Over the past several months, we've been developing the infrastructure behind our upcoming content signing and verification platform. That includes a production signature service, a Python SDK, browser-side JavaScript and CSS helpers, and the supporting systems needed to sign and independently verify content.
Recently, we took an important step. We integrated that same infrastructure directly into our own AI assistant.
Every assistant response is now signed before it is delivered to the user. Verification occurs independently using the same services and libraries that external developers will eventually have access to.
This integration has been invaluable. It allows us to test the platform under real-world conditions, validate our developer experience, and identify improvements before broader availability.
More importantly, it helps ensure we are holding ourselves to the same standard we expect from others. If we believe AI-generated content should be verifiable, then our own AI systems should be verifiable too.
The integration currently uses:
Our production signing and verification service
Browser-side JavaScript helpers
Verification UI components and styling helpers
While still early, this milestone represents an important step toward making content verification easier to integrate into real-world applications. We still have work to do, but the foundation is coming together.
Check it out by clicking the chat icon in the top right. Try tampering with the AI-generated content using your browser dev tools and see what happens.
Consider this our take on the “ Know Your Agent (KYA) ” dilemma. Turns out, signatures created with our BlkBolt™ tech are uniquely enabled to solve this problem.
AI agents are evolving fast. New models. New versions. New behaviors. Benchmarks change, systems get updated, and suddenly the thing you deployed last month isn’t quite the same anymore.
The general mantra seems to be…“ Looks right…probably fine. ”
With backgrounds in risk management and cyber security, let's just say, that feels… optimistic.
We’ve been thinking about a slightly different approach:
Don’t just trust the agent, verify what it actually produced.
We built a small demo exploring this idea in a streaming system.
The core concept is simple:
An LLM/agent generates a response
The final response is signed at creation time by the source system
The client can verify it before trusting it
Why This Matters (especially now)
In a world of rapidly changing AI models and agents:
How do you know which agent or version produced a response?
How do you confirm the response wasn’t altered in transit?
What happens when you don’t want to stand by past outputs?
How do you prove what was actually shown to a user?
Two systems produced the same thing, which came first?
Previously issued signatures will become invalid
You can stop trusting outputs from that version
You’re not permanently tied to everything it ever produced
🛒 Agentic Commerce - Your AI agent makes a purchase on your behalf.
Did your agent actually initiate that transaction?
Was the request altered anywhere along the way?
Can the business verify the same thing?
👶 Teacher agents - A child is interacting with an AI tutor.
Is the content displayed exactly what the system produced?
Was anything modified in the browser, by an extension, or in transit?
Can you audit what was actually shown?
🏦 High-stakes customer service - A banking agent issues a refund or provides financial guidance.
Can you prove what was presented or said to the user?
Can you audit later? (who, what, where, when)
Can you revoke trust in outputs from a faulty version?
How This Could Work (in production)
The goal is to make it very simple.
At a high level:
You own the BlkBolt™ encoding model/s and lease them to your AI agents
The leased model is used for signing and verification
The agent system makes one signing call when content is created
The response is sent with the signature artifact
The client verifies automatically in the background through separate API call
In addition to our key differentiator of revocability , BlkBolt™ produced signatures offer several additional benefits such as:
Embedded policies through metadata
No traditional key handling - compared to traditional signatures like RSA, ECDSA, etc.
Distributed verification - you control what happens to signatures, we just help verify them
Content-level trust - not just connection-level trust (e.g. TLS)
It’s early and a bit scrappy, but the core idea works. We think this becomes more relevant as agents act on behalf of users, systems get more autonomous, multi-agent systems become the norm, and trust shifts from “ the app ” to “ the output ”.
We’re curious what you think!
Where would something like this actually be useful?
Overkill? Or something that becomes standard?
What’s the first place you’d want this?
Want to share your feedback? Contact Us !
Digital signatures are everywhere. From software updates to financial transactions, they form the backbone of how we verify authenticity and integrity in modern systems.
For decades, that foundation has been built on cryptographic schemes like RSA and ECDSA, approaches grounded in well-established mathematical principles and public/private key infrastructure.
At lyfe.ninja, we’ve been exploring a different direction. While we are actively pursuing research opportunities to better evaluate the encoding strength of BlkBolt™, we’ve increasingly been drawn to digital signatures as a practical and immediate use case. This led us to a simple question:
Data is hashed to a fix length
Trust is rooted in hard mathematical problems
BlkBolt™ takes a different approach. Instead of keys, it uses trained models to encode and verify data. The result is a signature artifact tied not to a key, but to a model-based representation of the input.
At a high level, both approaches achieve the same core goals:
Integrity (data hasn’t changed)
Authenticity (it came from the expected source)
The Most Important Difference: Revocability
Traditional digital signatures are permanent by design.
It can be verified indefinitely
Even if the private key is later destroyed
Revocation is handled externally (CRLs, OCSP), and not always enforced
BlkBolt™ introduces a different model. Signature validity can be actively controlled after issuance.
Destroy the model → signatures can no longer be verified
Expire access → signatures fail validation (we do this through leases)
Enforce policies using embedded metadata (e.g., lease, timestamp) → signatures become conditionally valid
In the real world, trust is rarely permanent.
Documents can be accepted and later proven fraudulent.
Statements or endorsements can be made and later retracted.
Access that was valid at one point in time may no longer be appropriate.
Beyond Keys: A System-Level Shift
Secure storage of private keys
Signing authority is tied to a unique model instance
Verification can be tightly controlled
Critical components (like encoders or decoders) can remain isolated
Limited exposure of sensitive components
Controlled verification environments
To be clear, this isn’t about replacing RSA or ECDSA. Traditional digital signatures are proven, standardized, widely trusted, and interoperable.
BlkBolt™ is simply exploring a different approach for situations traditional signatures are not well suited for. This approach becomes interesting in scenarios like:
Ephemeral or session-based systems
Situations where permanent validity is a liability
Traditional signatures answer: “Was this signed?"
BlkBolt™ signatures open the door to: “Is this still valid?”
We’re still early in evaluating BlkBolt’s properties, especially from a formal security perspective.
People willing to pressure-test the ideas
If you fit in one of these categories we'd love to hear from you. Please reach out via our Contact page or by filling out our Work With Us form.
lyfe.ninja is pleased to announce that it has joined the CyLab Venture Network at Carnegie Mellon University!
Over the past 18 months, lyfe.ninja has been developing BlkBolt™ , a research-driven data protection system exploring neural-network-based approaches to data encoding with a focus on isolation, and long-term resilience. The Venture Network provides an opportunity to engage with a broader community of cybersecurity researchers, founders, and practitioners while continuing to refine real-world applications of this technology.
Participation in the CyLab Venture Network reflects lyfe.ninja’s commitment to thoughtful technical dialogue, external validation, and collaboration at the intersection of deep learning and next-generation data protection.
We look forward to engaging with the CyLab community over the coming year.
We’ve launched a new interactive demo showcasing one experimental use case of our core technology, BlkBolt™ : a revocable, model-verified signature for digital files.
The demo lets you:
Sign a file using a leased isolated model
See, in real time, how trust and verification change
Signatures that are revocable by design
Verification tied to isolated models , not long-lived keys
Authenticity without handing custody of your data to a third party
First-party content verification - Proving you authored a piece of content — without publishing or storing the original file.
Third-party endorsement and verification - Allowing an organization, brand, or authority to endorse or validate content created by someone else — and revoke that endorsement later if needed.
AI content attestation - Proving content was AI-generated, which model generated it, and who the model belonged to while avoiding watermark fragility, centralized registries, or irreversible claims.
Digital asset integrity & provenance - Verifying that file or digital assets haven’t changed over time, provide important metadata about the asset, and clearly surfacing when trust no longer applies.
Today, we’re excited to publicly share a major milestone for lyfe.ninja : the submission of a provisional patent titled “ System and Method for Neural Network-Based Data Encoding Using Randomized Tokenization and Latent-Space Representation. ”
This filing represents the foundation of BlkBolt™ , a new data protection technology born out of our long-running research initiative, Project BlackBox .
Why BlkBolt™ Exists
What’s Different About This Approach
Data is encoded , not mathematically ciphered
Interpretation requires access to a specific trained model , not a reusable key
Each deployment is isolated , reducing systemic risk
Models can evolve and rotate , allowing protected data to be re-encoded as the system improves
As the year winds down, we wanted to take a moment to say thank you and share a small update.
Over the past year, we’ve quietly made significant advancements in our core technology and the infrastructure that supports it. A lot of this work happened behind the scenes—iterating, testing, breaking things, fixing them, and learning fast—but it’s laid a foundation we’re genuinely proud of.
This season also marks nearly one year since lyfe.ninja became a business, which still feels a little surreal. What started as an idea has grown into something real, shaped by persistence, curiosity, and a lot of late nights.
We’re not ready to share all the details just yet… but we are excited to say that a major announcement is coming soon. Big things are in motion, and this next chapter is shaping up to be an important one.
For now, we’re grateful—for the progress, the lessons, and the support along the way. Wishing you a warm, restful holiday season and an exciting start to the new year.
🎄✨
— The lyfe

[truncated]
