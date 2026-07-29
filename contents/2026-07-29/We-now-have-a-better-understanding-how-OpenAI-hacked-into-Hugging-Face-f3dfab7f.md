---
source: "https://arstechnica.com/security/2026/07/jfrog-tries-to-spin-openai-0-day-exploit-of-its-app-into-a-success-story/"
hn_url: "https://news.ycombinator.com/item?id=49094248"
title: "We now have a better understanding how OpenAI hacked into Hugging Face"
article_title: "JFrog tries to spin OpenAI 0-day exploit of its app into a success story - Ars Technica"
author: "joozio"
captured_at: "2026-07-29T07:46:00Z"
capture_tool: "hn-digest"
hn_id: 49094248
score: 1
comments: 0
posted_at: "2026-07-29T07:02:30Z"
tags:
  - hacker-news
  - translated
---

# We now have a better understanding how OpenAI hacked into Hugging Face

- HN: [49094248](https://news.ycombinator.com/item?id=49094248)
- Source: [arstechnica.com](https://arstechnica.com/security/2026/07/jfrog-tries-to-spin-openai-0-day-exploit-of-its-app-into-a-success-story/)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T07:02:30Z

## Translation

タイトル: OpenAI がどのようにして Hugging Face をハッキングしたかについての理解が深まりました。
記事タイトル: JFrog は、OpenAI のアプリのゼロデイエクスプロイトを成功物語に仕立てようとしている - Ars Technica
説明: JFrog Artifactory 0-day を利用する OpenAI モデルからパッチのリリースまで 10 日が経過しました。

記事本文:
コンテンツにスキップ
アルス テクニカ ホーム
セクション
フォーラム
購読する
検索
AI
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ストーリーにピンを付ける
テーマ
ハイパーライト
検索
サインイン
サインインダイアログ...
サインイン
ここで祝うことは何もありません
OpenAI がどのようにして Hugging Face をハッキングしたのかについての理解が深まりました。
JFrog Artifactory 0-day を活用した OpenAI モデルからパッチのリリースまで 10 日が経過しました。
54
クレジット:
オーリッヒ・ローソン
クレジット:
オーリッヒ・ローソン
テキスト
設定
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ナビゲーション用に最小化する
先週、2つのOpenAIセキュリティハッキングモデルが同業のAI企業Hugging Faceのネットワークに侵入した前例のないセキュリティイベントは、Artifactoryの1つ以上のゼロデイ脆弱性を悪用することによって可能になったと、製品開発者のJFrogは月曜日に述べた。
ディストピア SF 小説を模倣した事件で、2 つの OpenAI モデルが内部テスト中にインターネットへのアクセスを防ぐことを目的とした制限された環境から脱出したと、AI 企業が先週明らかにしました。モデルたちはさらにハギング・フェイスのネットワークに侵入し、機密情報や資格情報を盗み出しました。 OpenAIは、そのエージェントがこれまで知られていなかった脆弱性を悪用することで偉業を達成したと述べた。同社はこの出来事を「前例のない」出来事だと呼び、外部の関係者も大部分がこれに同意した。
意図的に達成された勝利ではない
OpenAIによると、これらのモデルは認証情報の盗難やゼロデイなど複数の攻撃ベクトルを悪用してリモートでコードを実行する機能を獲得したが、これまで脆弱なソフトウェアは不明だったという。 JFrog の月曜日の開示では、この製品は自己管理型インスタンス Artifactory であり、顧客の安全性を確保し、合理化するリポジトリ管理システムであると述べられています。

ソフトウェア開発業務。 JFrog 氏によると、Artifactory は 7,500 以上の開発者チームによって使用されており、そのうちの 80% はフォーチュン 100 企業で働いています。
「フロンティアのサイバー能力の内部評価中に、隔離された研究環境で本番環境の保護策なしで意図的に実行されていた OpenAI のモデルは、連鎖的な脆弱性を自律的に発見して利用し、サンドボックスを脱出し、オープンなインターネットに到達し、Hugging Face のインフラストラクチャから評価の回答を抽出しました」と JFrog の CTO、Yoav Landman 氏は書いています。同幹部はさらに、同社はOpenAIからゼロデイのことを知ったと述べた。
同社は月曜日、悪用された脆弱性を修正したと発表したが、脆弱性を特定したり、脆弱性が悪用される条件などの重要な詳細については明らかにしなかった。このような詳細は、顧客がリスクを評価するために必要であるため、多くの脆弱性開示では標準となっています。同社の代表者は電子メールで詳細の提供を拒否した。
月曜日に公開されたバージョン Artifactory 7.161.15 のリリースノートには、パッチが適用された 9 つの脆弱性に対する CVE 指定がリストされています。この開示では、それらのいずれかが実際に積極的に悪用されたことについては言及されていませんでした。ただし、外部情報源によると、そのうちの 3 つ（ CVE-2026-65617 、 CVE-2026-65923 、および CVE-2026-66018 ）は、OpenAI 研究者の Khai Tran によって非公開で報告されたことが示されています。そのうち少なくとも 2 つは OpenAI のゼロデイ モデルが悪用されたものである可能性がありますが、確認がなければ断定することはできません。
このハッキングは、OpenAI のモデルのセキュリティ機能の内部テスト中に発生しました。同社はリスクの高い行動を阻止するはずのガードレールを意図的に無効にした。モデルを分離するはずだった環境は、最終的には

これは、現在 Artifactory であることがわかっている、名前のないホストされたパッケージ レジストリ プロキシとキャッシュを介したインターネットへの経路です。 OpenAIによると、モデルがExploitGymと呼ばれる業界標準のベンチマークのソリューションを見つけることに「過度に集中」すると、最終的には「かなり狭いテスト目標を達成するために多大な労力を費やすことになる」という。
これらの極端な対策の一環として、OpenAI モデルは Hugging Face ネットワークに侵入し、運用データベースの 1 つから必要なデータを盗みました。 Hugging Face は 7 月 16 日にこの侵害を明らかにしました。OpenAI は 7 月 21 日まで侵入の責任を明らかにしませんでした。
Landman氏の月曜日の投稿は、JFrogセキュリティチームがOpenAIの報告書を「世界に知られていない真のゼロデイとして、それに値する緊急性を持って扱い、それに応じて行動した」ため、事件全体を成功物語のように紡ぎ出そうとした。同 CTO はさらに、「人間が発見できなかったエクスプロイト パスをモデルが発見できるのと同じ機能が、防御側がそれらのパスを最初に見つけて根絶できるようにする機能です。」と付け加えました。
投稿から省略されているのは、OpenAI が Hugging Face が明らかにした侵害における役割を明らかにするまでに 5 日が経過し、OpenAI がゼロデイを報告し、JFrog がそれらに対するパッチをリリースしてから少なくともさらに 5 日が経過したことです。教訓: OpenAI エージェントが 10 日間有利にスタートできれば、悪意を持って使用されている他のモデルも有利にスタートできるでしょう。これは、JFrog と OpenAI が実現しようとしている成功物語とは言えません。
ゼロデイに関する JFrog の不透明さと組み合わせると、この事件はさらに悪化しているように見えます。 AI 企業の進歩のスピードを考えると、今後さらに悪化する可能性があります。
54件のコメント
コメント
フォーラムビュー
コメントを読み込んでいます...
前の話
次の話
よく読まれている
1.
専門家は現在のStarshipヒートシールド技術は迅速な再利用には「行き止まり」であると警告
2.
「グーグル

eとRedditはインターネットを所有していない」とウェブスクレーパーは裁判での勝訴後に語る
3.
アンダースコアの欠如により無実の男性が18か月の懲役刑に処せられた
4.
リアクションホイールの故障でスウィフト救出ミッションは軌道上を回転中
5.
活動家、国境職員に携帯電話を消去する「強迫コード」を与え重罪で起訴
カスタマイズ
Ars Technica は信号を分離してきました。
25年以上続く騒音。弊社独自の組み合わせにより、
技術的な知識と技術芸術への幅広い関心
Ars は、情報の海の中で信頼できる情報源です。後
すべてを知る必要はありません。重要なことだけを知っておく必要があります。

## Original Extract

10 days passed from OpenAI models exploiting JFrog Artifactory 0-day to release of a patch.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
NOTHING TO CELEBRATE HERE
We now have a better understanding how OpenAI hacked into Hugging Face
10 days passed from OpenAI models exploiting JFrog Artifactory 0-day to release of a patch.
54
Credit:
Aurich Lawson
Credit:
Aurich Lawson
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
Last week’s unprecedented security event in which two OpenAI security hacking models trespassed into the network of fellow AI company Hugging Face was enabled by exploiting one or more zero-day vulnerabilities in Artifactory, JFrog, the product’s developer, said Monday.
In an incident mimicking a dystopian sci-fi novel, two OpenAI models broke out of the restricted environment meant to keep them from accessing the Internet during an internal test, the AI company revealed last week . The models went on to breach Hugging Face’s network and steal confidential information and credentials. OpenAI said its agent achieved the feat by exploiting a previously unknown vulnerability. The company called the event “unprecedented,” and outsiders largely agreed.
Not the triumph made out to be
OpenAI said the models exploited multiple attack vectors, including stolen credentials and zero-days, to gain remote code execution capabilities, but until now, the vulnerable software was unknown. JFrog’s Monday disclosure said the product was a self-managed instance Artifactory, a repository management system that secures and streamlines customers’ software development operations. JFrog says Artifactory is used by more than 7,500 developer Teams, 80 percent of which work for Fortune 100 companies.
“During an internal evaluation of frontier cyber capabilities, OpenAI’s models, running deliberately without production safeguards in an isolated research environment, autonomously discovered and employed chained vulnerabilities to escape its sandbox, reach the open internet, and extract evaluation answers from Hugging Face’s infrastructure,” JFrog CTO Yoav Landman wrote. The executive went on to say that the company learned of the zero-days from OpenAI.
The company said Monday that it fixed the exploited vulnerabilities, but it didn’t identify them or provide other important details, such as the conditions under which the vulnerabilities can be exploited. Such details are standard in many vulnerability disclosures because they’re necessary for customers to assess risks. In an email, a company representative declined to provide the details.
Release notes published Monday for version Artifactory 7.161.15 listed the CVE designations for nine patched vulnerabilities. The disclosure made no mention that any of them had been actively exploited in the wild. External sources, however, show that three of them— CVE-2026-65617 , CVE-2026-65923 , and CVE-2026-66018 —were privately reported by OpenAI researcher Khai Tran. It’s likely that at least two of them were the zero-days OpenAI’s models exploited, but without confirmation, it’s impossible to say so definitively.
The hack came during an internal OpenAI test of its models’ security capabilities. The company deliberately disabled guardrails that are supposed to block high-risk actions. The environment that was supposed to isolate the models ended up having a pathway to the Internet through an unnamed hosted package-registry proxy and cache that we now know to be Artifactory. When the models “hyperfocused” on finding a solution for an industry-standard benchmark called ExploitGym, one ended up going to “extreme lengths to achieve a rather narrow testing goal,” OpenAI said.
As part of these extreme measures, the OpenAI model broke into the Hugging Face network and stole the needed data from one of its production databases. Hugging Face disclosed the breach on July 16. OpenAI didn’t reveal its culpability in the intrusion until July 21.
Landman’s Monday post tried to spin the entire incident as a success story because the JFrog security team treated OpenAI’s report “with the urgency it deserved, as a genuine zero-day unknown to the world, and moved accordingly.” The CTO added: “The same capability that lets a model find an exploit path no human had found is the capability that will let defenders find and eradicate those paths first.”
Left out of the post is that five days passed until OpenAI revealed its role in the breach Hugging Face disclosed and that at least another five days passed from the time OpenAI reported the zero-days and JFrog released patches for them. The lesson: If OpenAI agents could gain a 10-day head start, so too can other models being used maliciously. This is hardly the success story JFrog and OpenAI are trying to make it out to be.
Combined with JFrog’s opaqueness surrounding the zero-days, the incident looks even worse. Given the speed at which AI companies are moving, there may still be worse to come.
54 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Experts warn current Starship heat shield tech is a "dead end" for rapid reuse
2.
“Google and Reddit do not own the Internet," web scraper says after court win
3.
A missing underscore sent innocent man to prison for 18 months
4.
Reaction wheel failures leave Swift rescue mission spinning in orbit
5.
Activist charged with felony after giving border agent "duress code" that wiped his phone
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
