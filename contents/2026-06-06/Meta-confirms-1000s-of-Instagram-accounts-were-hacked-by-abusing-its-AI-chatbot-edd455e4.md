---
source: "https://this.weekinsecurity.com/meta-confirms-thousands-of-instagram-accounts-were-hacked-by-abusing-its-ai-chatbot/"
hn_url: "https://news.ycombinator.com/item?id=48427643"
title: "Meta confirms 1000s of Instagram accounts were hacked by abusing its AI chatbot"
article_title: "Meta confirms thousands of Instagram accounts were hacked by abusing its AI chatbot"
author: "speckx"
captured_at: "2026-06-06T21:27:51Z"
capture_tool: "hn-digest"
hn_id: 48427643
score: 178
comments: 67
posted_at: "2026-06-06T18:35:27Z"
tags:
  - hacker-news
  - translated
---

# Meta confirms 1000s of Instagram accounts were hacked by abusing its AI chatbot

- HN: [48427643](https://news.ycombinator.com/item?id=48427643)
- Source: [this.weekinsecurity.com](https://this.weekinsecurity.com/meta-confirms-thousands-of-instagram-accounts-were-hacked-by-abusing-its-ai-chatbot/)
- Score: 178
- Comments: 67
- Posted: 2026-06-06T18:35:27Z

## Translation

タイトル: Meta、数千の Instagram アカウントが AI チャットボットを悪用してハッキングされたことを確認
記事タイトル: Meta、数千の Instagram アカウントが AI チャットボットを悪用してハッキングされたことを確認
説明: Meta は、Meta AI チャットボットを騙して 2 要素認証を持たない Instagram アカウントのパスワードをリセットできるバグを修正しました。

記事本文:
サインイン
購読する
2026 年 6 月 6 日
4 分で読めます
記事
Meta、数千の Instagram アカウントが自社の AI チャットボットを悪用してハッキングされたことを確認
Metaは、同社のAIチャットボットが数か月に渡って悪用され、ハッカーが繰り返し騙してアカウントを乗っ取った際にInstagramアカウントがハイジャックされた数千人に通知している。
今週セキュリティ分野で公開された新しいデータ侵害通知レターの中で、Meta は長期にわたるハッキング キャンペーンの一環としてアカウントをハイジャックされた人の数を初めて明らかにしました。このキャンペーンは今週初めに発見され、404 Media ($) と TechCrunch ($) によって最初に報告されました。影響を受けたアカウントの数から、このハッキング キャンペーンがどれほど広範囲に及んだのか、またどのくらいの期間にわたって活動していたのかがある程度明確になります。
金曜日遅くにメイン州司法長官事務所に提出されたデータ侵害通知によると、メタはメイン州の30人を含む少なくとも2万225人にアカウントが侵害されたと通知した。
この侵害により、ハッカーは連絡先情報、生年月日、プロフィール情報の取得を含む、その人物のインスタグラム全体とそのリンクされたアカウントを乗っ取ることができたほか、その人物の投稿、ダイレクトメッセージ、アカウントアクティビティへのアクセスも可能になったと通知には書かれている。
Metaの通知では、この侵害が「InstagramのAI支援アカウント回復システムの脆弱性」に関連しており、この脆弱性が悪用されて「Instagramのユーザーアカウントのパスワードリセットを実行する」ことが確認されている。
以前に報告されたように、ハッカーは Meta のチャットボットの欠陥を悪用し、2 要素認証がオンになっていないアカウントのパスワードを誰でもリセットできるようにしました。このバグにより、チャットボットはアカウントではなくハッカーが管理する電子メール アドレスに検証コードを送信するように仕向けられました。

尋ねるだけで、ファイルに登録されている所有者の電子メール アドレスを知ることができます。とにかくチャットボットは従った。
「ツール自体は適切に動作し、意図したとおりに機能していましたが、別のコードパスのバグにより、システムは、パスワードのリセットを要求した個人が提供した電子メールアドレスが、そのユーザーのInstagramアカウントに関連付けられた電子メールアドレスと一致するかどうかを適切に検証しませんでした」とMetaは侵害通知の中で述べている。
「その結果、個人が以前にアカウントに関連付けられていない電子メールアドレスを提供した場合、システムはリクエストを拒否するのではなく、その関連付けられていない電子メールに誤ってパスワードリセットリンクを送信してしまいました。これにより、権限のない第三者が、自分が所有していないアカウントのパスワードリセットリンクを受信する可能性がありました」と同社は付け加えた。
メタ氏によると、現時点ではハッカーが誰かのパスワードをリセットし、あたかも正当な所有者であるかのようにアカウントを乗っ取る可能性があるという。
～今週のセキュリティ～は、あなたのような読者に支えられている私の週刊サイバーセキュリティ ニュースレターです。限定記事や分析などをご覧いただくために、月額 10 ドルから始まる有料サブスクリプションへのサインアップをご検討ください。
または、サポートを示すために 1 回限りのチップを送信することもできます。
メタ社は、ハッキング中に個人情報がアクセスされたとしても、その内容については「把握していない」と述べた。 （この件について明確にするようメタ社のプレスラインに送った電子メールは土曜日早朝の時点で返信されていない。）
メイン州のリストによると、ハッキングは4月17日頃に始まり、メタ社がチャットボットを確保したと発表した今週まで続いた。ハッキングが進行中であると一部で報告されていたにもかかわらず、インスタグラムは今週初めにパスワードリセット通知を送信することで影響を受けた個人への通知を開始したと伝えられている。
メタ社は通知の中で、「影響を受けるユーザーにパスワードをリセットするよう指示した」として、ユーザーにアカウントを保護するよう警告したことも認めた。

安全で検証されたチャネルを通じて再認証します。」
Metaは、現時点ではAIチャットボットを無効にし、チャットボットによるユーザーアカウントのリセットを可能にしたコードパスを削除したと述べ、また、インシデントの再発を防ぐためにプラットフォーム全体の他のチャットボットもチェックしていると述べた。どのような状況でチャットボットが悪用されたのかはまだ明らかではないが、同社が AI の活用を倍増させる中、Meta が数千人の従業員を解雇し、経営陣に株式インセンティブを与えた直後に発生した。
～今週のセキュリティ～をお読みいただき、誠にありがとうございます。この記事が気に入ったら、ぜひシェアしてください！この記事に関するフィードバック、質問、コメントがありましたら、お気軽に this@weekinsecurity.com までご連絡ください。
知っておくべきサイバー ニュースをすべて毎週配信します。
Zack Whittaker の週刊サイバーセキュリティ ニュースレターにサインアップしてください。手書き、傾きゼロ。
電子メールの開封やリンクの追跡はありません。いつでも購読を解除してください。
オウラ氏は、ユーザーデータに対する政府の要求を取得していると述べた。何人でシェアするのでしょうか？
4 分で読めます
20
5月
'26
AI はバグや欠陥を見つけることができますが、サイバーセキュリティの基本を忘れないでください
5 分で読めます
11
5月
'26
2026 年のインターネットに対する最も危険な脅威
04
5月
'26
どんでん返し: 私は司法省とFBIを告訴します
3 分で読めます
02
5月
'26
すべての組織がセキュリティ上の欠陥を簡単に報告できるようにする必要がある理由
8 分で読めます
～今週のセキュリティ～ © 2026
購読する

## Original Extract

Meta fixed the bug that let anyone trick its Meta AI chatbot into resetting the password on Instagram accounts that didn't have two-factor authentication.

Sign in
Subscribe
06 Jun 2026
4 min read
Articles
Meta confirms thousands of Instagram accounts were hacked by abusing its AI chatbot
Meta is notifying thousands of people whose Instagram accounts were hijacked during the months-long abuse of the company's AI chatbot, which hackers repeatedly tricked into taking control of a person's account.
In a new data breach notification letter , seen by this week in security , Meta has revealed for the first time how many people had their accounts hijacked as part of the long-running hacking campaign, which was discovered earlier this week and first reported by 404 Media ($) and TechCrunch ($) . The number of affected accounts gives some clarity as to how widespread this hacking campaign was, and for how long it operated.
According to the data breach notice filed with Maine's attorney general's office late on Friday, Meta notified at least 20,225 people that their accounts had been compromised, including 30 people in Maine.
The compromises allowed the hackers to take over the person's entire Instagram and any linked accounts, including obtaining contact information, dates of birth, and profile information, as well as the ability to access the person's posts, direct messages, and account activity, the notice reads.
Meta's notice confirmed that the breach relates to "a vulnerability in an AI-assisted account recovery system for Instagram," which was exploited to "perform password resets on Instagram user accounts."
As previously reported, hackers abused a flaw in Meta's chatbot that allowed anyone to reset the password of any account that did not have two-factor authentication switched on. The bug tricked the chatbot into sending a verification code to an email address controlled by the hacker, rather than the account holder's email address on file, simply by asking it. The chatbot complied anyway.
"The tool itself worked properly and functioned as intended; however due to a bug in a separate code path, the system did not properly verify that the email address provided by the individual requesting a password reset matched the email address associated with that user’s Instagram account," said Meta in its breach notice.
"As a result, when an individual provided an email address not previously associated with the account, the system incorrectly sent a password reset link to that unassociated email rather than rejecting the request. This allowed unauthorized third parties to receive a password reset link for accounts they did not own," the company added.
At this point, Meta says, the hackers could reset someone's password and take over their account as if they were the rightful owner.
~this week in security~ is my weekly cybersecurity newsletter supported by readers like you. Please consider signing up for a paying subscription starting at $10/month for exclusive articles, analysis, and more.
Or, you can submit a one-time tip to show your support!
Meta said that it is "unaware" of what, if any, personal information was accessed during the hacks. (An email to Meta's press line asking for clarity on this was unreturned as of early Saturday.)
According to Maine's listing, the hacks began around April 17 and lasted until this week, when Meta said that it had secured the chatbot. Instagram reportedly started notifying affected individuals earlier this week by sending a password reset notification, even as some reported that the hacks were ongoing .
Meta also confirmed in the notice that it alerted users to secure their accounts, saying it "instructed impacted users to reset their passwords and re-authenticate through secure, verified channels."
Meta said that it has disabled the AI chatbot for now and removed the code path that allowed the chatbot to reset user accounts, and said it's also checking other chatbots across its platforms to prevent a repeat incident. It's not yet clear what circumstances led up to the chatbot being abused, but comes soon after Meta laid off thousands of employees while rewarding top executives with stock incentives , as the company continues to double-down on AI.
Thank you so much for reading ~this week in security~. If you liked this article, please share it! Feel free to reach out with any feedback, questions, or comments about this article: this@weekinsecurity.com .
Get all the cyber news you need to know, delivered weekly.
sign up for Zack Whittaker's weekly cybersecurity newsletter. Hand-written, zero slop.
No email open or link tracking. Unsubscribe anytime.
Oura says it gets government demands for user data. Will it share how many?
4 min read
20
May
'26
AI can find bugs and flaws, but don't forget the cybersecurity basics
5 min read
11
May
'26
The most dangerous threats to the internet in 2026
04
May
'26
Plot twist: I'm suing the Justice Department and FBI
3 min read
02
May
'26
Why every organization should make it easy to report security flaws
8 min read
~this week in security~ © 2026
Subscribe
