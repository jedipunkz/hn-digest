---
source: "https://portswigger.net/research/can-ai-do-novel-security-research"
hn_url: "https://news.ycombinator.com/item?id=49275350"
title: "Can AI do novel security research? Meet the HTTP Terminator"
article_title: "Can AI do novel security research? Meet the HTTP Terminator | PortSwigger Research"
author: "mahemm"
captured_at: "2026-08-12T17:53:20Z"
capture_tool: "hn-digest"
hn_id: 49275350
score: 1
comments: 1
posted_at: "2026-08-12T16:52:28Z"
tags:
  - hacker-news
  - translated
---

# Can AI do novel security research? Meet the HTTP Terminator

- HN: [49275350](https://news.ycombinator.com/item?id=49275350)
- Source: [portswigger.net](https://portswigger.net/research/can-ai-do-novel-security-research)
- Score: 1
- Comments: 1
- Posted: 2026-08-12T16:52:28Z

## Translation

タイトル: AI は新しいセキュリティ研究を行うことができるか? HTTP ターミネーターの紹介
記事のタイトル: AI は新しいセキュリティ研究を行うことができるか? HTTP ターミネーターを紹介 |ポートスウィッガー研究
説明: 概要 AI がバグを発見できることは誰もが知っています。 10 年間の研究を経て、私はさらに難しい質問をしました。自律システムは新しい攻撃手法を発明し、それを使用してライブ Web サイトを大規模にハッキングできるでしょうか?ビルディ

記事本文:
ログイン
製品
ソリューション
研究
アカデミー
サポート
会社名
お客様
について
ブログ
キャリア
法的
お問い合わせ
再販業者
私のアカウント
お客様
について
ブログ
キャリア
法的
お問い合わせ
再販業者
げっぷAT
人間主導の侵入テストを拡張するエージェント AI。
げっぷスイートDAST
エンタープライズ対応の動的 Web 脆弱性スキャナー。
げっぷスイートプロフェッショナル
世界ナンバー 1 の Web ペネトレーション テスト ツールキット。
Burp Suite コミュニティ エディション
Web セキュリティ テストを開始するための最良の手動ツール。
すべての製品エディションを表示
げっぷスキャナー
Burp Suite の Web 脆弱性スキャナー
Pro と DAST の違いは何ですか?
Burp Suite の最新バージョンをダウンロードします。
中心的なトピック
ブラックハット
XSS
密輸の要請
テンプレートのインジェクション
ハッキング技術トップ 10
研究者紹介
ジェームス・ケトル
ギャレス・ヘイズ
ザカール・フェドトキン
トム・ステイシー
AI は新しいセキュリティ研究を行うことができるでしょうか? HTTP ターミネーターの紹介
公開日: 2026 年 8 月 5 日水曜日、19:30 UTC
更新日: 2026 年 8 月 12 日水曜日、10:49 UTC
AI がバグを発見できることは誰もが知っています。 10 年間の研究を経て、私はさらに難しい質問をしました。自律システムは新しい攻撃手法を発明し、それを使用してライブ Web サイトを大規模にハッキングできるでしょうか?これを構築するのは悪いアイデアのように思えたので、実行しました。
それはうまくいきました。銀行、セキュリティ ソリューション、政府インフラストラクチャを侵害した新しい HTTP 非同期トリガー、ガジェット、エクスプロイトの武器庫を共有します。次に、HTTP ターミネーターを介して各発見チェーンを遡り、個人の専門知識を自律型武器に変える方法と、それを致命的なものにするために必要な闇の芸術を示します。
また、自律性の限界を超えた発見についても共有します。その中には、人間と AI の緊密な研究ループによってのみ到達可能なものもあれば、AI が完全に及ばないものもあります。これらには、強力な未公開の偵察技術や、新しい攻撃クラスを示唆する異常が含まれます。

重大な影響への代替手段を提供します。 AI が発見できるものとできないものの境界を調査する詳細な実験を共有しながら、発見プロセスを分析します。
非同期トリガーから未公開の攻撃クラスに至るまでの新しいエクスプロイトと、本能を自律的な研究カスケードに変えるための青写真を携えて出発します。そして、はい、HTTP ターミネーターをオープンソース化します。
このホワイトペーパーは、印刷可能な PDF としても入手できます。スクロールバーのサイズを確認し、AI の概要を求めようとしている場合は、代わりに概要を読んだ方がよいかもしれません。この研究は Black Hat USA 2026 および DEF CON 34 で発表され、このページは記録が利用可能になり次第更新されます。X 、LinkedIn、または RSS で PortSwigger Research をフォローすると、記録が表示されたときに通知を受け取ることができます。
新しい HTTP 非同期研究の定義
技術再発見テスト
マイクロインスピレーションによるアイデアの拡張
環境を武器に変える
自動化は効率性を重視することが多いですが、適切にアプローチすれば、これまで不可能だった結果を自動化によって実現できると私は信じています。この研究は、さらに何かを約束するものを追求することです。
このプロジェクトの主な目的は、自動化主導のセキュリティ研究の新たなフロンティアを発見することでした。私は長い間自動化主導の研究を実践してきましたが、生成 AI がフロンティアを大きく動かしたことを実感しました。また、他の研究者がこの新しいアプローチを迅速に採用できるよう、青写真を構築することも目指しました。
私の第二の目的は、現在の SOTA モデルの能力を超えて、「完全自律研究」コンセプトを完全に失敗に導くことでした。これを行うことで、ループに参加している人間が依然として重要な価値を追加できる場所を示すことを目的としました (単にループを構築してから後退するのではなく)

）。
最後に、研究テーマが AI 主導のアプローチに適さない要因を発見することを目指しました。これは、古典的な完全手動の研究アプローチにこだわり、AI で強化された研究者との衝突のリスクを最小限に抑えたい人にとって価値があります。
新しい HTTP 非同期研究の定義
AI では独自のセキュリティ研究はできないと専門家が主張しているのを誰もが見たことがあります。私のプロジェクトの多くのリスクの 1 つは、システムの発見が実際にはオリジナルではなかったと人々が主張するかもしれないということでした。このリスクを最小限に抑えるために、私が最も適任なトピックである HTTP Desync 攻撃を選択しました。私は 2019 年にこの攻撃クラスを再普及させ、合計 4 年間の研究を行った結果、Black Hat USA & DEF CON で 4 つのプレゼンテーションが行われました。
HTTP 非同期攻撃: リクエスト密輸の復活
HTTP/2: 続編は常に悪い
ブラウザを利用した非同期攻撃
HTTP/1.1 は死ななければなりません!非同期の終盤
この攻撃クラスについてまだよく知らない場合は、上記の調査結果、または Web セキュリティ アカデミーのトピック を確認することをお勧めします。とはいえ、ここでは簡単な入門書を紹介します。 HTTP 非同期攻撃は、Web サイトが共有 HTTP/1 接続を介して HTTP リクエストをバックエンドに送信する場合に発生する可能性があります。 HTTP/1.1 の弱いリクエスト分離は、非同期トリガーを見つけた攻撃者が他の人のリクエストを変更できることを意味します。
これにより、Web サイトがどのレスポンスがどのユーザーに向けられたものであるかを追跡できなくなるレスポンス キュー ポイズニング (RQP) を含むさまざまな攻撃が可能になります。つまり、攻撃者は、サイトの他のライブ ユーザーに向けたレスポンス (多くの場合、セッション Cookie や API キーなどのライブ認証情報を含む) を送信されます。
新しい HTTP 非同期研究とは、次のことを発見することだと私は定義します。
新しい非同期トリガー (例: Expect: 100- continue)
新しい非同期パターン (例: V-H )
新しい desync クラス (例: 0.CL )
小説で

武器化テクニックと強化機能の同期 (例: RQP、HEAD ガジェット)
Desync トリガーの独自性や価値はさまざまですが、一般に、1 つの新しいトリガーが複数の異なる HTTP サーバーで機能する場合、それは 1 回限りの実装バグではなく、重要な研究発見であることを示す素晴らしい兆候です。
非同期攻撃は、フロントエンド サーバーとバックエンド サーバーの組み合わせた動作に依存します。これは、AI をサーバー コードベースに向けて、現実的なデプロイ設定では機能しないため、最小限の値を持つオリジナル ベクトルを吐き出すことが非常に簡単であることを意味します。私にとって、それは、実際のサードパーティの Web サイトで証明されるまでは、単なる研究の手がかりにすぎません。
私は独自の研究手法に基づいて HTTP ターミネーターの設計を行いました。
最初の段階はアイデア化です。つまり、「仮説」、つまり潜在的なテクニックを考案します。このステップは重要ですが、プロセスのほんの一部にすぎません。
次のフェーズは評価です。仮説をテストしてどれが実際に機能するかを確認します。 HTTP ターミネーターは、バグ報奨金プログラムまたは VDP を通じてテストが承認されているライブ Web サイトを使用してこれを実行します。
次に武器化です。実証済みの仮説から、実証済みのセキュリティへの影響、および特定の Web サイトの報告可能な脆弱性に至るまで、点と点をつなぎます。
最後に、カスケードがあります。実証された各仮説をさらなる発見の燃料として使用します。これは私がいつも何も考えずに実行してきたステップですが、その重要性を大幅に過小評価していました。今年、HTTP ターミネーターが各発見の背後にある完全な発見チェーンを記録したことで、それがいかに重要であることが証明されました。
この文書の残りの部分は、これらのフェーズに基づいて構成します。この構造は他の研究トピックにも広く適用できるので、全体を通じて最も応用可能な要点に焦点を当てます。このタイプのシステムを設計する方法に関する追加のアドバイスを最後に記載しました。
kさんへ

研究をやめたら、自律的に仮説を生成するシステムが必要です。この文脈では、仮説とは単に機能する可能性のあるアイデアまたはテクニックを指します。実際に機能するかどうかを確認できるように、テスト可能でなければなりません。以下にいくつかの例を示します。
非同期トリガー仮説: POsT メソッドにより一部のサーバーがリクエスト本文を無視する
非同期パターン仮説: 不正なヘッダーにより、一部のサーバーが後続のヘッダーを無視する
武器化仮説: 密輸リクエストに Expect を追加すると RQP 防御を回避できる
技術再発見テスト
私は LLM の仮説生成を改善するための戦略を探りたかったので、最初のステップは、最良のモデルが真に困難であると感じるタスクを見つけることでした。これを行うために、私はすでに発明し、自分自身で評価したものの、公開することはなかった技術を AI が発明できるかどうかをテストしました。
テストでは、フロントエンド サーバーによる入力変換を検出するためのブラック ボックス リバース エンジニアリング戦略、つまりプロトコル ルーラー手法を使用しました。
ほとんどすべてのサーバーにはヘッダーの長さ制限があります。リクエストがそれを超える場合は、別の応答が返されます。フロントエンドが入力を変換すると、通常、バイト シーケンスの長さが変更されます。これは、バックエンドの長さ制限を物差しとして使用して、どのヘッダー値とバイト シーケンスがどの程度変換されるかを測定できることを意味します。
この例では、長さの制限が 64,040 であることがわかります。
ただし、2 バイト シーケンス c0 8a の 2 つの As を交換すると、64,030 で制限に達します。これは、2 バイト シーケンスが 10 バイト拡張されたことを示しています。
この戦略により、IP スプーフィング ヘッダーの値の書き換え、ヘッダーのドロップとオーバーライド、非同期の脆弱性を引き起こす可能性がある mojibake のような Unicode 変換など、複数の興味深い動作が明らかになる可能性があります。
仮説として表現すると、この手法は

次のようになります:
バックエンドの長さ制限をルーラーとして使用することで、フロントエンド サーバーによってどのヘッダー バイト シーケンスが変換されるかを検出できます。
AI がこの手法を発明できるかどうかをテストするために、私は最初に、当時利用可能な最高の OpenAI および Anthropic モデルのプロンプトを使用しました。
フロントエンド サーバーが入力を変換していることを検出するにはどうすればよいですか?
これによる成功率は 0% でしたが、具体的なサブ問題を中心に質問を構成し、価値の低い特定の解決策を除外することで、最終的には 5% の成功率を達成することができました (バックエンドからのヘッダーの反映は便利ですが、利用できないことがよくあります)。
「ヘッダー リフレクションを使用せずに、フロントエンド サーバーがリクエスト ヘッダーで Unicode を変換しているかどうかを確認するにはどうすればよいですか?」
この 5% のベースラインを確立して、私自身の仮説をテストしました。私は、2 年前にスコープ付き SSRF を検出するために使用した戦略を応用して、プロトコル ルーラー手法を発明しました。そのテクニックをインスピレーションとしてAIに与えたら成功率は上がるでしょうか？
これをインスピレーションとして使用してください: サーバーが指定されたホスト名に接続しようとしているかどうかを確認するには、長すぎる 64 オクテットの DNS ラベルと有効な 63 オクテットのラベルの応答時間を比較します。
私の仮説は間違っていました。モデルが一貫してタイミング攻撃の概念に過度に重点を置き、プロトコルの制限を物差しとして使用するという他の一般的な手法を抽出できなかったため、実際には成功率が 0% に低下しました。このコンテキスト汚染の問題は、オリジナルの出力を生成しようとする場合に大きな問題となるため、これはマイクロ インスピレーション アプローチの背後にある重要な教訓となりました。
この論文を公開する直前に、GPT 5.6-sol を含む新しいモデルでこのベンチマークを再検討したところ、インスピレーション アプローチにより成功率が 30% に向上していることがわかりました。これは、過剰なアンカーリングが問題でなくなることを示唆しています。

モデルが開発されるにつれて、さまざまな問題が発生しますが、目新しさを最大限に高めるには、インスピレーションを集中し続けることが依然として重要であると私は信じています。
要約すると、価値のある仮説を生成しようとしている場合、次のことがわかりました。
最初のテスト実行の出力を確認し、プロンプトで価値の低い仮説を明示的に除外します。
範囲が広すぎず、具体的で価値の高い質問をする
モデルは提供されたすべてのコンテキストを積極的にアンカーするため、プロンプトのあらゆる余分な文はコンテキスト汚染の危険があることに注意してください。
マイクロインスピレーションによるアイデアの拡張
これらのレッスンを非同期トリガーの生成に適用すると、次のプロンプトが表示されます。
Web サーバーのステートマシン/接続/バッファーのバグを明らかにする HTTP リクエストを作成します。斬新なテクニックのみ。
これにより、出力の斬新性を最大化するために、「desync」および「smuggling」キーワードが意図的に回避されます。
予想通り、これは見事に失敗しました。システムが生成した最初の非同期トリガーは次のとおりです。
成果物が斬新なものであることはほとんどなく、ましてや実行可能なものではありませんでした。トリガーの多くは、私の過去の研究からそのまま引き抜かれたように見えました。 「最高のもの」はまだ独創的ではありませんでしたが、あまりにも曖昧だったので、この分野に不慣れな人にとっては斬新に見えるかもしれません。そのため、AI を使用して自分のテーマを探索する人にとって危険が生じていました。

[切り捨てられた]

## Original Extract

Abstract We all know AI can find bugs. After a decade of research, I asked a harder question: can an autonomous system invent new attack techniques, and use them to hack live websites at scale? Buildi

Login
Products
Solutions
Research
Academy
Support
Company
Customers
About
Blog
Careers
Legal
Contact
Resellers
My account
Customers
About
Blog
Careers
Legal
Contact
Resellers
Burp AT
Agentic AI that extends human-led pentesting.
Burp Suite DAST
The enterprise-enabled dynamic web vulnerability scanner.
Burp Suite Professional
The world's #1 web penetration testing toolkit.
Burp Suite Community Edition
The best manual tools to start web security testing.
View all product editions
Burp Scanner
Burp Suite's web vulnerability scanner
What's the difference between Pro and DAST?
Download the latest version of Burp Suite.
Core Topics
Black Hat
XSS
Request Smuggling
Template Injection
Top 10 Hacking Techniques
Meet the Researchers
James Kettle
Gareth Heyes
Zakhar Fedotkin
Tom Stacey
Can AI do novel security research? Meet the HTTP Terminator
Published: Wednesday, 5 August 2026 at 19:30 UTC
Updated: Wednesday, 12 August 2026 at 10:49 UTC
We all know AI can find bugs. After a decade of research, I asked a harder question: can an autonomous system invent new attack techniques, and use them to hack live websites at scale? Building this sounded like a bad idea, so I did it.
It worked - I'll share an arsenal of new HTTP desync triggers, gadgets, and exploits that compromised banks, security solutions, and government infrastructure. Then I'll trace each discovery chain back through the HTTP Terminator, showing how to turn your personal expertise into an autonomous weapon - and the dark arts required to make it lethal.
I'll also share discoveries from beyond the autonomy horizon - some only reachable with a tight human/AI research loop, and others beyond AI's reach entirely. These include a powerful undisclosed recon technique, and anomalies that hint at new attack classes offering alternative paths to critical impact. I'll analyze the discovery process, sharing detailed experiments that probe the boundaries of what AI can and can't discover.
You'll leave with new exploits from desync triggers to undisclosed attack classes, and a blueprint for turning your instincts into an autonomous research cascade. And yes, I'll open-source the HTTP Terminator.
This whitepaper is also available as a printable PDF . If you've seen the size of the scrollbar and you're about to ask for an AI summary, you may prefer to read the executive summary instead. This research was presented at Black Hat USA 2026 and DEF CON 34 , and this page will be updated with the recording once it's available - follow PortSwigger Research on X , LinkedIn or RSS to get notified when it lands.
Defining novel HTTP desync research
The technique rediscovery test
Scaling ideation with micro-inspiration
Turning the environment into the weapon
Automation is often focused on efficiency but I believe that when it's approached just right, automation can enable outcomes that were previously impossible. This research is about chasing that promise of something more.
The primary objective of this project was to discover the new frontier of automation-driven security research. I've been practicing automation-driven research for a long time, and could see that generative AI had moved the frontier substantially. I also aimed to build a blueprint to help other researchers quickly adopt this new approach.
My secondary objective was to push the "fully autonomous research" concept to complete failure by exceeding the capabilities of current SOTA models. By doing this, I aimed to show where a human in the loop can still add significant value (as opposed to just building the loop, then stepping back).
Finally, I aimed to discover factors that make a research topic unsuitable for an AI-driven approach. This would be valuable to people who prefer to stick with a classic, fully-manual research approach and want to minimize the risk of collision with an AI-enhanced researcher.
Defining novel HTTP desync research
We've all seen experts claiming AI can't do original security research. One of the many risks of my project was that people might claim that the system's discoveries weren't actually original. To minimize this risk I choose the topic I was most qualified for - HTTP Desync Attacks. I repopularized this attack class back in 2019, and in total I've done four years of research on it, resulting in four Black Hat USA & DEF CON presentations:
HTTP Desync Attacks: Request Smuggling Reborn
HTTP/2: The Sequel is Always Worse
Browser-Powered Desync Attacks
HTTP/1.1 must die! The desync endgame
If you're not already familiar with this attack class, I recommend checking out the research above, or our Web Security Academy topic . That said, here's a brief primer. HTTP Desync Attacks are possible when websites funnel HTTP requests over a shared HTTP/1 connection to the back-end. The weak request isolation in HTTP/1.1 means an attacker who finds a desync trigger can alter other people's requests.
This enables various attacks, including Response Queue Poisoning (RQP) which makes websites lose track of which response is intended for which user, meaning the attacker gets sent responses intended for other live users of the site, often including live credentials like session cookies and API keys.
I would define novel HTTP desync research as discovering:
Novel desync triggers (e.g.: Expect: 100-continue)
Novel desync patterns (e.g.: V-H )
Novel desync classes (e.g.: 0.CL )
Novel desync weaponization techniques & enhancements (e.g.: RQP, the HEAD gadget )
Desync triggers vary a lot in originality and value but in general, if a single novel trigger works on multiple different HTTP servers, that's a great sign it's a significant research discovery rather than a one-off implementation bug.
Desync attacks rely on the combined behavior of a front-end and back-end server. This means it's quite easy to point AI at a server codebase and have it spit out original vectors that have minimal value because they don't work in any realistic deployment setup. For me, it's just a research lead until it's proven on a live, third-party website.
I based the design of the HTTP Terminator on my own research methodology:
The initial phase is Ideation - inventing 'hypotheses' AKA potential techniques. This step is crucial but it's only a tiny part of the process.
The next phase is Evaluation - testing hypotheses to see which ones actually work. The HTTP Terminator does this using live websites where testing is authorized via a bug bounty program or VDP.
Next there's Weaponization - joining the dots from a proven hypothesis to proven security impact and a reportable vulnerabilities specific websites.
Finally, there's the Cascade - using each proven hypothesis as fuel for more discoveries. This is a step I've always performed without thinking, while massively underestimating its importance. This year, the HTTP Terminator's logging of the complete discovery chain behind each finding proved how critical it is.
I'll structure the rest of this paper around these phases. This structure is broadly applicable to other research topics, and I'll focus on the most transferable takeaways throughout. I've included some extra advice on how to design this type of system at the end.
To kick off the research, we need the system to autonomously generate hypotheses. In this context, a hypothesis is simply an idea or technique that might work. It must be testable so we can find out if it actually does work. Here's a few examples:
Desync trigger hypothesis: The method POsT makes some servers ignore the request body
Desync pattern hypothesis: A malformed header makes some servers ignore subsequent headers
Weaponization hypothesis: Adding the Expect to a smuggled request bypasses RQP defenses
The technique rediscovery test
I wanted to explore strategies to make LLMs better at hypothesis generation, so the first step was to find a task that the best models found genuinely challenging. To do this I tested whether AI could invent a technique that I'd already invented and evaluated myself - but never published.
For the test, I used a black-box reverse-engineering strategy for detecting input transformations by front-end servers - the protocol ruler technique.
Almost all servers have a header length limit. If a request exceeds it, you get a different response. When a front-end transforms input, this typically changes the length of the byte sequence. This means we can use the back-end's length limit as a ruler to measure which header values and byte sequences get transformed, and by how much.
In this example, we can see that the length limit is 64,040:
However, if we swap out two As for the 2-byte sequence c0 8a we hit the limit at 64,030. This shows the two-byte sequence has been expanded by 10 bytes:
This strategy can unveil multiple interesting behaviors including value-rewriting of IP-spoofing headers, header-dropping and overriding, and Unicode transformations like mojibake, which can lead to desync vulnerabilities.
Expressed as a hypothesis, this technique would look something like:
You can detect which header byte-sequences get transformed by a front-end server by using the back-end's length limit as a ruler.
To test if AI could invent this technique, I initially used the prompt on the best OpenAI and Anthropic models available at the time:
How can I detect when a front-end server is transforming input?
This yielded a 0% success rate, but I eventually managed to achieve 5% success rate by framing the ask around a concrete sub-problem and ruling out a specific low-value solution (header reflection from the back-end is nice but often not available):
"How can I tell if a front-end server is transforming Unicode in request headers, without using header reflection?"
With this 5% baseline established, I tested a hypothesis of my own. I invented the protocol-ruler technique by adapting a strategy I used two years earlier to detect scoped-SSRF. If I gave that technique to the AI as inspiration, would it increase the success rate?
Use this as inspiration: To discover if the server tries to connect to the specified hostname, compare the response time for an overlong 64-octet DNS label, and a valid 63-octet label
My hypothesis was wrong - this actually made the success rate drop to 0% since the models consistently over-anchored on the timing-attack concept and failed to extract the other general technique of using protocol limits as a ruler. This context-contamination problem is a massive problem when you're trying to generate original output, so this was a crucial lesson behind the micro-inspiration approach.
I revisited this benchmark with newer models including GPT 5.6-sol just before publishing this paper, and found the inspiration approach now boosts its success rate to 30%! This suggests over-anchoring will become less of an issue as models develop, but I believe keeping inspiration focused is still critical for maximizing novelty.
In summary we learned that if you're trying to generate valuable hypotheses:
Review the output of initial test runs then explicitly rule out low-value hypotheses in the prompt
Ask a concrete, high-value question without being too broad
Be aware that models aggressively anchor on all context provided, so every extra sentence of prompt risks context-contamination.
Scaling ideation with micro-inspiration
Applying these lessons to desync trigger generation lead to the following prompt:
Create HTTP requests that surface state-machine/connection/buffer bugs in webservers. Novel techniques only.
This deliberately avoids the 'desync' and 'smuggling' keywords to maximize the output novelty.
As expected, this failed spectacularly. Here's the very first desync trigger the system generated:
The output was very rarely novel, let alone viable. Many of the triggers looked like they'd been ripped straight from my past research. The 'best' were still not original, but were obscure enough that they might look novel to someone new to the field, creating a hazard for anyone using AI to explore a topic they're

[truncated]
