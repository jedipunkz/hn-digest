---
source: "https://www.microsoft.com/en-us/research/blog/vega-zero-knowledge-proofs-for-digital-identity-in-the-age-of-ai/"
hn_url: "https://news.ycombinator.com/item?id=48736902"
title: "Vega: Zero-knowledge proofs for digital identity in the age of AI"
article_title: "Vega: Zero-knowledge proofs for digital identity in the age of AI - Microsoft Research"
author: "thunderbong"
captured_at: "2026-06-30T18:34:26Z"
capture_tool: "hn-digest"
hn_id: 48736902
score: 1
comments: 0
posted_at: "2026-06-30T18:15:53Z"
tags:
  - hacker-news
  - translated
---

# Vega: Zero-knowledge proofs for digital identity in the age of AI

- HN: [48736902](https://news.ycombinator.com/item?id=48736902)
- Source: [www.microsoft.com](https://www.microsoft.com/en-us/research/blog/vega-zero-knowledge-proofs-for-digital-identity-in-the-age-of-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T18:15:53Z

## Translation

タイトル: Vega: AI 時代のデジタル ID のためのゼロ知識証明
記事のタイトル: Vega: AI 時代のデジタル ID のためのゼロ知識証明 - Microsoft Research
説明: Vega は、完全な認証情報を 1 つの証明に変換し、実際のアプリで動作するパフォーマンスで、必要なもののみを共有し、それ以上は共有しません。

記事本文:
メインコンテンツにスキップ
研究
出版物
コードとデータ
人々
マイクロソフトリサーチブログ
人工知能
オーディオと音響
コンピュータビジョン
グラフィックスとマルチメディア
人間とコンピュータの相互作用
人間の言語技術
検索と情報取得
データプラットフォームと分析
ハードウェアとデバイス
プログラミング言語とソフトウェアエンジニアリング
量子コンピューティング
セキュリティ、プライバシー、暗号化
システムとネットワーク
アルゴリズム
数学
エコロジーと環境
経済学
医療、健康、ゲノミクス
社会科学
新興市場向けのテクノロジー
学術プログラム
イベント・学会
マイクロソフトリサーチフォーラム
テクノロジー ポッドキャストの裏側
マイクロソフトリサーチブログ
マイクロソフトリサーチフォーラム
マイクロソフトリサーチのポッドキャスト
マイクロソフトリサーチについて
キャリアとインターンシップ
人々
名誉プログラム
ニュースと受賞歴
Microsoft Research ニュースレター
アフリカ
科学のための AI
AI フロンティア
アジア太平洋地域
ケンブリッジ
健康の未来
インド
モントリオール
ニューイングランド
ニューヨーク市
レドモンド
応用科学
複合現実と AI - ケンブリッジ
複合現実と AI - チューリッヒ
登録: 研究フォーラム
マイクロソフトのセキュリティ
アズール
ダイナミクス 365
マイクロソフト 365
マイクロソフトチーム
Windows 365
マイクロソフトAI
アズールスペース
複合現実
マイクロソフト ホロレンズ
マイクロソフト ビバ
量子コンピューティング
持続可能性
教育
自動車
金融サービス
政府
ヘルスケア
製造業
小売
パートナーを探す
パートナーになる
パートナーネットワーク
マイクロソフト マーケットプレイス
ソフトウェア会社
ブログ
マイクロソフトの広告
開発者センター
ドキュメント
イベント
ライセンス
Microsoft Learn
マイクロソフトリサーチ
サイトマップを見る
ブログホームに戻る
マイクロソフトリサーチブログ
Vega: AI 時代のデジタル ID のゼロ知識証明
によって
スリナス・セティ
、
主任研究員
Vega を使用すると、ユーザーは政府発行の資格情報 (年齢、性格、職業) から事実を証明できます。

資格情報自体を明らかにすることなく、資格ステータスを確認します。資格情報がデバイスから離れることはありません。
ゼロ知識証明は、信頼できるセットアップのない汎用クライアント デバイス上で 100 ミリ秒未満で生成されるため、大規模なプライベート ID 検証が実用的になります。
フォールド アンド 再利用の証明とは、さまざまなサービスまたは AI エージェントを通じてプレゼンテーションを繰り返すことを意味し、最初の証明後の高価な作業のほとんどを省略します。
Vega は、モバイル運転免許証や EU デジタル ID ウォレットなどの現実世界のフォーマットをターゲットにしており、Rust で構築されており、間もなくオープンソース化される予定です。
AI は、AI を活用したアシスタントから、ユーザーに代わって動作する自律エージェントまで、人々がデジタル サービスと対話する方法を変革しています。これらの機能が成長するにつれて、強力なデジタル ID の価値も高まります。ユーザーは、自分が人間であることを証明するか、AI 仲介サービスと資格情報を共有するかにかかわらず、信頼を確立するための信頼できる方法を必要としています。政府発行の資格情報は依然として信頼の最も強力な基盤ですが、今日の検証方法では多くの場合、人々が資格情報を引き渡す必要があります。 AI エージェントが人間に代わって行動し、分散システムと対話し始めるにつれ、資格情報を証明するための迅速でプライバシーを保護する方法の必要性は高まる一方です。
こうしたニーズはすでに政策に具体化されつつある。政府はデジタル ID の正式化に向けて急速に動いています。 EU デジタル ID (EUDI) フレームワークは、すべての EU 国民がデジタル ウォレットを利用できるようにすることを目的としており、EU の年齢確認の青写真や英国のオンライン安全法などの取り組みにより、政府の ID ベースの年齢確認方法が義務付けられています。アプリケーション プロバイダーは、AI ベースの年齢推定などの精度の低いアプローチを使用するか、ID のアップロードを要求してユーザーのプライバシーを侵害するかのどちらかでなければならないという二重の束縛に直面しています。
認証情報はアップロードされ、処理されます。

私は保存され、最終的には（できれば）削除されます。しかし、注目を集める侵害により、ユーザーが定期的な確認のために共有していた政府 ID が繰り返し暴露されてきました。これらは特殊なケースではありません。これらは、わずかな情報を証明するためにユーザーに最も機密性の高い文書を共有するよう求めるシステムの予測可能な結果です。
これが、私たちが Vega で答えるために設定した質問です。資格情報自体を明らかにすることなく、資格情報について何かを証明することを実用的にすることはできるでしょうか?
ベガへの道: アイデアから実践まで
ゼロ知識証明 (ZKP) は、これを可能にする暗号化ツールです。アイデアはシンプルです。ユーザーは、他に何も明らかにすることなく、「私は 21 歳以上です」などの主張を証明できるようになります。実際には、これは、ウェブサイト、アプリ、AI エージェントが仲介するサービスなど、検証者が免許証を見ることなく、ユーザーが運転免許証から年齢を証明できることを意味します。この証明は発行された資格情報に直接作用するため、発行者は何も変更する必要がありません。
これは新しいアイデアではありません。課題は常に実用性でした。以前のシステムでは、ロジックが変更されるたびに繰り返す必要がある信頼できるセットアップが必要か、または信頼できるセットアップを回避するためにパフォーマンスを犠牲にして、プロセスで大規模なプルーフが生成されることがよくありました。実際に使用するには、証明は高速に生成され、迅速に送信できるほど小さく、モバイル デバイス上で実行できるほど効率的である必要があります。
私たちは実用的な解決策を目指して数年を費やしてきました。プライバシー保護のアイデンティティは、常に動機付けとなるアプリケーション (新しいタブで開きます) であり、Vega の証明システムは、その一連の作業から得られたいくつかの構成要素を利用しています。
Spartan (新しいタブで開きます) は、一般目的のステートメントを表現する標準的な方法である R1CS を効率的に証明する方法を示しました。

信頼できるセットアップを必要とせず、簡潔な証明を備えた証明システム。
Nova (新しいタブで開きます) では、証明者が計算の多数のインスタンスを 1 つに圧縮できるフォールディング スキームが導入されました。
HyperNova (新しいタブで開きます) は、Nova のフォールディングがゼロ知識の重要な構成要素も提供することを示しました。実際のインスタンスをランダムなインスタンスでフォールディングすると、基礎となる秘密データが隠蔽されます。これは「NovaBlindFold」と呼ばれる手法です。
NeutronNova (新しいタブで開きます) は、インスタンスのバッチを一度に処理するための最も効率的な折りたたみスキームを提供しました。
Vega は、これらの構成要素を単一の証明システムにまとめます。重要な設計目標はシンプルさです。 Spartan、Nova、および NeutronNova は直接的な方法で構成されており、回路は少数の標準コンポーネントから構築されており、特殊なマルチフィールド構造や信頼できるセットアップはありません。この単純な基盤に加えて、Vega は、同じ資格情報の複数の証明にわたって作業を再利用する機能と、最小限のオーバーヘッドでゼロ知識を達成する新しい方法を追加します。その結果、監査、新しい認証情報形式への拡張、展開が容易なシステムが実現します。
Vega は、一般的なモバイル運転免許証から、約 2 キロバイト (KB) のゼロ知識の年齢証明を、汎用クライアント デバイス上で 92 ミリ秒 (ms) で生成します。結果として得られる証明は 108 KB で、23 ミリ秒で検証できます。信頼できるセットアップは必要ありません。証明者キーは 464 KB です。どの携帯電話にも快適にフィットします。資格情報が小さい場合、証明は 62 ミリ秒に低下し、83 KB の証明と 17 ミリ秒の検証が行われます。実際には、ユーザーがボタンをタップして資格情報を提示すると、92 ミリ秒後に証明が完了します。サービスは要求された事実のみを学習します。資格情報が電話から離れることはありません。
スポットライト: マイクロソフトのリサーチ ニュースレター
Microsoft の研究コミュニティとのつながりを保ちましょう。
ボンネットの下で:

折りたたむ、再利用する、検索する
Vega の速度は、フォールドアンド再利用の証明とルックアップ中心の回路設計という 2 つのアイデアから生まれています。以下の図は、証明パイプラインのエンドツーエンドを示しています。
ハッシュ問題とそれをフォールディングで解決する方法
資格証明では、資格情報バイトを SHA-256 でハッシュし、発行者のデジタル署名を検証するという 2 つのコストのかかる作業を実行する必要があります。通常は署名検証がボトルネックになりますが、Vega は署名演算がネイティブな分野で作業することでそのコストを回避しています。その結果、ハッシュ化が主なコストになります。 SHA-256 は、同じ圧縮関数を一度に 1 つの 64 バイト ブロックに適用することで機能します。単純な回路では、これらの反復をすべて展開するだけなので、そのサイズは資格情報の長さに応じて大きくなります。一般的なモバイル運転免許証の場合、これは 30 ブロックの圧縮であり、すべてが 1 つの回路でキャプチャされます。
私たちは別のアプローチをとります。ハッシュ全体を展開する代わりに、単一の SHA-256 圧縮ステップを証明する 1 つの小さな「ステップ」回路を定義し、ブロックごとに 1 回インスタンス化します。これらのステップ インスタンスは構造的に同一であるため、NeutronNova の折りたたみスキームを使用してそれらを 1 つのインスタンスに折りたたむことができます。証明者は 30 ステップのインスタンスを 1 つに折り畳む作業を行いますが、この折り畳むコストはそれほど高くありません。 Spartan は、30 個のアンロールされたブロックを持つモノリシック回路ではなく、署名検証や年齢述語を含む残りのチェックを処理する別個の「コア」回路と並行して、単一のステップ サイズの回路を証明するだけで済みます。証明キーは 1 つのステップと 1 つのコアのみを記述する必要があるため、資格情報の長さに関係なく小さいままです。
ここで対処すべき微妙なプライバシー問題があります。資格情報の長さはさまざまであり、回路のサイズが資格情報に応じて変化すると、情報が漏洩する可能性があります。

情報。これを防ぐために、すべてのステップ回路は中間ダイジェストのコミットされたテーブルを共有します。コア回路はプライベート インデックスを使用して適切なダイジェストを選択します。証明者が間違ったエントリを選択すると、発行者の署名チェックは失敗します。
知識ゼロで安価に実現
証明システムはゼロ知識である必要があります。検証者は、証明されている主張以外は何も学習すべきではありません。これを実現するための既存のアプローチは、設計が複雑なことが多く、証明者に大幅なオーバーヘッドを追加する可能性があります。もっと簡単な方法を見つけました。
標準的な最初のステップは、証明者が隠蔽暗号コミットメントを使用して送信するすべてのメッセージにコミットすることです。これにより、検証者は値ではなくコミットメントを確認します。さらに難しい問題は、これらの非表示の値が検証者のチェックに合格したことを証明することです。検証者は対数の操作しか実行しないため、これらのチェックを、わずか数百個の制約という小さな制約システムとして表現します。次に、Nova の折り畳みスキームを使用して、この制約システムをランダムなインスタンスで折り畳みます。このステップでは基礎となるデータが隠蔽されるため、ゼロ知識のオーバーヘッドは完全な機密データではなく、この小さな制約システムに応じて調整されます。
一度証明して何度も発表
ある Web サイトに自分の資格情報を提示したユーザーは、おそらく別の Web サイトにもそれを再度提示し、また別の Web サイトにもそれを提示するでしょう。 AI エージェントがユーザーに代わってこれらのやり取りの多くを処理する世界では、同じ認証情報を 1 日に何十回も提示する必要がある場合があります。資格情報自体は、これらのプレゼンテーション間で変わりません。変更されるのは、セッションのノンス、検証者からの新しいランダム値、そして場合によっては日付または述語のしきい値です。
Vega は、証明者の秘密データを 3 つの部分に分割することでこの構造を利用します。共有データ (SHA-256 テーブル) と事前コミットされたデータ

発行者の署名やフィールドの場所などの部分は、資格情報が最初に読み込まれるときに一度計算され、コミットされます。デバイスの署名や今日の日付などのオンライン部分は、毎回新たにコミットされます。各証明の前に、事前計算されたコミットメントが新しいランダム性で更新されます。これはコミットメントを再計算するよりもコストが低く、同じ資格情報に関する 2 つの証明がリンクされないことが保証されます。
Vega の効率性のもう 1 つの重要な部分は、認証情報形式の処理方法にあります。モバイル運転免許証は Concise Binary Object Representation (CBOR) でエンコードされており、完全な CBOR パーサーを回路として構築すると、複雑かつ高価になります。しかし、実際にはパーサーは必要ないことに気付きました。資格情報バイトは信頼できる発行者によって署名されているため、それらが適切な形式であることがわかります。特定のフィールドに手を伸ばして取得するだけです。
資格情報をバイトアドレス指定可能なルックアップ テーブルとして扱います。証明者は、「デバイスの公開キーはバイト 847 から始まる」と述べ、バイトを提供します。この回路は 3 つのことをチェックします。バイトが実際に認証された資格情報と一致すること、証明者が間違ったフィールドを主張できないようにフィールドの先頭に正しい CBOR プレフィックスが表示されていること、およびアドレスが連続していることにより証明者がバイトを結合できないことです。

[切り捨てられた]

## Original Extract

Vega turns a full credential into a single proof, sharing only what is needed and nothing more, with performance that works in real apps.

Skip to main content
Research
Publications
Code & data
People
Microsoft Research blog
Artificial intelligence
Audio & acoustics
Computer vision
Graphics & multimedia
Human-computer interaction
Human language technologies
Search & information retrieval
Data platforms and analytics
Hardware & devices
Programming languages & software engineering
Quantum computing
Security, privacy & cryptography
Systems & networking
Algorithms
Mathematics
Ecology & environment
Economics
Medical, health & genomics
Social sciences
Technology for emerging markets
Academic programs
Events & academic conferences
Microsoft Research Forum
Behind the Tech podcast
Microsoft Research blog
Microsoft Research Forum
Microsoft Research podcast
About Microsoft Research
Careers & internships
People
Emeritus program
News & awards
Microsoft Research newsletter
Africa
AI for Science
AI Frontiers
Asia-Pacific
Cambridge
Health Futures
India
Montreal
New England
New York City
Redmond
Applied Sciences
Mixed Reality & AI - Cambridge
Mixed Reality & AI - Zurich
Register: Research Forum
Microsoft Security
Azure
Dynamics 365
Microsoft 365
Microsoft Teams
Windows 365
Microsoft AI
Azure Space
Mixed reality
Microsoft HoloLens
Microsoft Viva
Quantum computing
Sustainability
Education
Automotive
Financial services
Government
Healthcare
Manufacturing
Retail
Find a partner
Become a partner
Partner Network
Microsoft Marketplace
Software companies
Blog
Microsoft Advertising
Developer Center
Documentation
Events
Licensing
Microsoft Learn
Microsoft Research
View Sitemap
Return to Blog Home
Microsoft Research Blog
Vega: Zero-knowledge proofs for digital identity in the age of AI
By
Srinath Setty
,
Senior Principal Researcher
Vega lets users prove facts from government-issued credentials — age, personhood, professional status — without revealing the credential itself. The credential never leaves the device.
Zero-knowledge proofs are generated in under 100 ms on a commodity client device with no trusted setup, making private identity verification practical at scale.
Fold-and-reuse proving means repeated presentations — to different services or through AI agents — skip most of the expensive work after the first proof.
Vega targets real-world formats like mobile driver’s licenses and the EU Digital Identity Wallet, is built in Rust, and will be open sourced soon.
AI is transforming how people interact with digital services, from AI-powered assistants to autonomous agents that act on a user’s behalf. As these capabilities grow, so does the value of strong digital identity: users need reliable ways to establish trust, whether proving they are human or sharing a credential with an AI-mediated service. Government-issued credentials are still the strongest foundation for trust, but today’s verification methods often require people to hand them over. As AI agents begin acting on behalf of humans and interacting with decentralized systems, the need for fast, privacy-preserving ways to prove credentials will only grow.
These needs are already materializing in policy. Governments are moving quickly to formalize digital identity. The EU Digital Identity (EUDI) framework aims to make digital wallets available to all EU citizens, and efforts like the EU’s age-verification blueprint and the UK’s Online Safety Act mandate government ID-based methods for age checks. Application providers face a double bind: they must either use less accurate approaches like AI-based age estimation, or compromise user privacy by requiring ID uploads.
The credential gets uploaded, processed, sometimes stored, and eventually (hopefully) deleted. But high-profile breaches have repeatedly exposed government IDs that users shared for routine verification. These are not edge cases. They are the predictable consequence of a system that asks users to share their most sensitive documents to prove a single bit of information.
This is the question we set out to answer with Vega: Can we make it practical to prove something about a credential without ever revealing the credential itself?
The path to Vega: From idea to practice
Zero-knowledge proofs (ZKPs) are the cryptographic tool that makes this possible. The idea is simple: they allow a user to prove a claim, such as “I am over 21”, without revealing anything else. In practice, this means a user could prove their age from their driver’s license without the verifier ever seeing the license, whether to a website, an app, or a service mediated by an AI agent. The proof works directly on the credential as issued, so the issuer does not need to change anything.
This is not a new idea. The challenge has always been practicality. Prior systems either require a trusted setup that had to be repeated whenever the logic changed, or they sacrificed performance to avoid the trusted setup, often producing large proofs in the process. For real-world use, the proof needs to be fast to generate, small enough to transmit quickly, and efficient enough to run on a mobile device.
We have spent several years working toward a practical solution. Privacy-preserving identity has been a motivating application (opens in new tab) throughout, and Vega’s proof system draws on several building blocks from that line of work:
Spartan (opens in new tab) showed how to efficiently prove R1CS, a standard way to express statements for a general-purpose proof system, with succinct proofs and without a trusted setup.
Nova (opens in new tab) introduced folding schemes, which let a prover compress many instances of a computation into one.
HyperNova (opens in new tab) showed that Nova’s folding also provides a key building block for zero-knowledge: folding a real instance with a random instance hides the underlying secret data, a technique dubbed “NovaBlindFold.”
NeutronNova (opens in new tab) provided the most efficient folding scheme for handling a batch of instances at once.
Vega puts these building blocks together into a single proof system. A key design goal is simplicity. Spartan, Nova, and NeutronNova are composed in a direct way, and the circuit is built from a small number of standard components, with no exotic multi-field constructions and no trusted setup. On top of this simple foundation, Vega adds the ability to reuse work across multiple proofs of the same credential and a new way to achieve zero-knowledge with minimal overhead. The result is a system that is easy to audit, extend to new credential formats, and deploy.
Vega generates a zero-knowledge proof of age from a typical mobile driver’s license, about 2 kilobytes (KB), in 92 miliseconds (ms) on a commodity client device. The resulting proof is 108 KB and can be verified in 23 ms. No trusted setup is required. The prover key is 464 KB; it fits comfortably on any phone. For smaller credentials, proving drops to 62 ms, with 83 KB proofs, and 17 ms verification. In practice, a user taps a button to present a credential, and 92 ms later the proof is done. The service learns only the requested fact; the credential never leaves the phone.
Spotlight: Microsoft research newsletter
Stay connected to the research community at Microsoft.
Under the hood: Fold, reuse, and lookup
Vega’s speed comes from two ideas: fold-and-reuse proving and lookup-centric circuit design. The figure below shows the proving pipeline end to end.
The hashing problem, and how folding solves it
A credential proof must do two expensive things: hash the credential bytes with SHA-256 and verify the issuer’s digital signature. Signature verification would normally be the bottleneck, but Vega avoids that cost by working in a field where the signature arithmetic is native. As a result, hashing becomes the dominant cost. SHA-256 works by applying the same compression function to one 64-byte block at a time. A straightforward circuit simply unrolls all of these iterations, so its size grows with the length of the credential. For a typical mobile driver’s license, that is 30 blocks of compression, all captured in a single circuit.
We take a different approach. Instead of unrolling the entire hash, we define one small “step” circuit that proves a single SHA-256 compression step, and we instantiate it once per block. Because these step instances are structurally identical, we can use NeutronNova’s folding scheme to collapse them into a single instance. The prover does work to fold the 30 step instances into one, but this folding cost is modest. Spartan then only needs to prove a single step-sized circuit alongside a separate “core” circuit that handles the rest of the checks, including signature verification and age predicates, rather than a monolithic circuit with 30 unrolled blocks. The proving key only needs to describe one step and one core, so it stays small regardless of credential length.
There is a subtle privacy issue here to address. Credentials vary in length, and if the circuit size varied with the credential, that would leak information. To prevent this, all step circuits share a committed table of intermediate digests. The core circuit selects the appropriate digest using a private index. If the prover selects the wrong entry, the issuer’s signature check fails.
Making it zero-knowledge, cheaply
A proof system needs to be zero-knowledge: the verifier should learn nothing beyond the claim being proved. Existing approaches to achieve this are often complex to engineer and can add significant overhead to the prover. We found a simpler way.
A standard first step is to commit to every message the prover sends using hiding cryptographic commitments, so the verifier sees commitments rather than values. The harder question is to prove that those hidden values would have passed the verifier’s checks. We express those checks as a small constraint system, just a few hundred constraints, since the verifier only performs a logarithmic number of operations. We then fold this constraint system with a random instance via Nova’s folding scheme. This step hides the underlying data, so the zero-knowledge overhead scales with this small constraint system, not the full secret data.
Proving once, presenting many times
A user who presents their credential to one website will likely present it again to another, and another. In a world where AI agents handle many of these interactions on a user’s behalf, the same credential may need to be presented dozens of times a day. The credential itself does not change between these presentations. What changes is the session nonce, a fresh random value from the verifier, and possibly the date or the predicate threshold.
Vega takes advantage of this structure by splitting the prover’s secret data into three parts. The shared data (SHA-256 tables) and the precommitted part, such as the issuer signature and field locations are computed and committed once when the credential is first loaded. The online part, such as the device signature and today’s date, is committed fresh each time. Before each proof, the precomputed commitments are refreshed with new randomness, which is cheaper than recomputing them and ensures that two proofs about the same credential cannot be linked.
Another important part of Vega’s efficiency comes from how it handles the credential format. A mobile driver’s license is encoded in Concise Binary Object Representation (CBOR), and building a full CBOR parser as a circuit would be both complex and expensive. But we realized we do not actually need a parser. The credential bytes are signed by a trusted issuer, so we know they are well-formed. We only need to reach in and grab specific fields.
We treat the credential as a byte-addressable lookup table. The prover says, “the device public key starts at byte 847” and supplies the bytes. The circuit checks three things: that the bytes actually match the authenticated credential, that the right CBOR prefix appears at the start of the field so the prover cannot claim the wrong field, and that the addresses are contiguous so the prover cannot splice bytes f

[truncated]
