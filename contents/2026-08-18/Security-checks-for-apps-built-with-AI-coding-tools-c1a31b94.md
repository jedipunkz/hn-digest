---
source: "https://lukatasolutions.com"
hn_url: "https://news.ycombinator.com/item?id=49351074"
title: "Security checks for apps built with AI coding tools"
article_title: "Lukata"
image: "https://lukatasolutions.com/og-blueprint.png"
author: "LukataSolutions"
captured_at: "2026-08-18T19:21:26Z"
capture_tool: "hn-digest"
hn_id: 49351074
score: 2
comments: 1
posted_at: "2026-08-18T19:13:40Z"
tags:
  - hacker-news
  - translated
---

# Security checks for apps built with AI coding tools

- HN: [49351074](https://news.ycombinator.com/item?id=49351074)
- Source: [lukatasolutions.com](https://lukatasolutions.com)
- Score: 2
- Comments: 1
- Posted: 2026-08-18T19:13:40Z

## Translation

タイトル: AI コーディング ツールで構築されたアプリのセキュリティ チェック
記事タイトル: ルカタ
説明: AI コーディング ツールで構築したものを、実際に実際に使用できるソフトウェアに変えるための無料の青写真。データ、権限、ビジネス ルール、アップロード、アカウント、支払い、AI の安全性、クレーム、プライバシー、起動、回復。

記事本文:
ルカタ ブループリントにスキップ LUKATA スタート チェック ガイド 起動 概要 無料で読めます。無料でご利用いただけます。
AI を使って何かを構築したいと考えていますか?普段見落としがちな重要なことを見つけてください!
逆転思考とは、問題を逆向きに考えることです。成功する方法ではなく、失敗する可能性があることから始めます。
アイデアに対して逆転の発想を活用しましょう。どのように失敗する可能性があるのか​​、何が原因でユーザーが 5 秒以内にサイトから離れてしまうのかを考えてください。問題を特定し、修正を繰り返します。構築する前に、自分のアイデアをストレステストしてください。意図的に失敗しなくなるまで失敗させて、自分のソフトウェアを破壊してみます。ランディング ページのすべてが真実であり、事実であることを確認してください。
最初の 1 時間はここから始まります。これらは個人データの漏洩や法外な請求を捕らえます。
60 枚の小切手を読まないつもりですか?これを AI に貼り付けると、構築したものにどれが適用されるかを判断します。
1 ある人が他の人のものを開けることはできますか?これは、ある顧客が別の顧客の個人情報を閲覧してしまう最も一般的な方法です。ステップ 1 自分で確認してください 2 つのアカウントを作成します。それらを A と B と呼びます。
まずBとしてログインします。注文やメッセージなど、B に属するものを 1 つ開きます。 Web アドレスは通常、数字で終わります。その番号を書き留めてください。
ログアウト。 A としてログインします。同じ種類のページを開き、アドレスに A の代わりに B の番号を入力します。
B の架空の数値ではなく、実際の数値を使用します。存在しない番号は「見つかりません」と返され、何もテストしなかった場合はパスしたように見えます。
A が B のものを見ることができれば、それはバグです。そのページのボタンも試してみてください。編集、保存、削除。機能するものはすべて、同じバグのさらに悪いバージョンです。毎回エラーまたは空のページが表示されるはずです。
2 ブラウザのコードにパスワードまたはキーが含まれていますか?これなら誰でも2クリックくらいで読めます。ステップ1チャンネル

自分で確認してください 訪問者のブラウザで実行されているコードで、パスワードやキーのようなものを検索します。次に、git 履歴も検索します。 Git 履歴は、コードのこれまでのすべてのバージョンのコピーが保存されているため、昨年削除した行がまだそこに残っています。
訪問者のブラウザが独自に行うすべての決定をリストします。お金、誰かが何をすることが許されているか、誰が何を閲覧できるかについて、あらゆるものを探してください。
3 一人で一晩で請求書を使い切ることができますか?誰も何も盗む必要はありません。彼らはあなたに仕事の対価を支払わせればいいだけです。ステップ 1 自分で確認する アカウントをまったく持たずに、1 人があなたのアプリに 1 分間で何をさせることができるかを尋ね、それにどれくらいの費用がかかるかを計算してください。
製品にお金がかかる場合は、13 を確認してください。 AI があなたが書いていないものを読み取った場合は、チェック 17 を行ってください。
実際にアプリを使用する前に確認すべき 5 つのこと。
すべての機能が動作しているように見えても、これらは見落とされがちです。
ユーザー入力ではデータベースリクエストを書き換えることはできません
ユーザー入力をブラウザコードに変換することはできません
他のサイトはサインインしているユーザーとしてコマンドを送信できません
サインイン Cookie は盗んだり再利用したりするのが困難です
管理者は 2 番目のサインイン手順を使用します
ログインすると、承認されたページにのみユーザーが戻ります
エラーページには非公開の技術詳細が隠されている
プライベートであるはずのファイルはパブリックではありません
ライブアプリには独自の秘密鍵があります
アプリは承認されたアドレスにのみ接続します
他のサービスからの応答は使用前にチェックされます
接続がタイムアウトになり、リダイレクトが制御されたままになる
承認されたユーザーのみが変更を公開できます
コードはリリース前にキーの漏洩がチェックされます
バックアップが正常に復元されました
複数の顧客が 1 つのシステムを共有します。すべてのデータベース クエリ、ジョブ、キャッシュ、検索、およびファイル パス内に、これがどの顧客に含まれるかをチェックします。
ソーシャル ログイン というワンタイム コードを使用して、アプリ内でログインが実際に開始されたことを確認します。

PKCE。自分でリストした Web アドレスにのみユーザーを送り返してください。
AI による非公開文書の検索 検索自体は 1 人の顧客の文書のみを対象とします。それらの文書が存在する場所はどこでもロックダウンしてください。すべての回答の横にソースを表示します。
AI がアクションを実行する アクションを実行する AI 機能は、エージェントと呼ばれることがよくあります。作業は独自のスペースに保管してください。インターネット上でアクセスできる範囲を削減します。覚えておきたいこと、やりたいこと、大きなことはすべて承認します。それを止めるための迅速な方法を確保してください。
AI は顧客の機密データを処理します。以下にリンクされている OWASP の AI アプリ チェックリストのレベル 2 を使用します。
さらに詳しいレビューが必要ですか? OWASP は、無料のセキュリティ チェックリストを発行する非営利団体です。まずは Web アプリのチェックリストから始めます。最終製品に AI が使用されている場合は、AI チェックリストを追加します。
安全であることとそれを証明することは別の仕事です。
まずは自分に当てはまるものを見つけてください。次に、適切なレビューを取得します。このガイドは作業の一部に役立ちます。あなたの会社を認証したり、法律を決定したりすることはできません。
SOC 2 外部の専門家が貴社の情報保護方法を確認し、レポートを作成します。重要な場合 多くの場合、企業顧客が、会社が情報をどのように保護しているかについて、あなた以外の誰かからの証拠を求めている場合が重要です。
CPAと呼ばれる公認会計士がチェックを行い、報告書を作成します。このガイドは準備に役立つかもしれません。これは SOC 2 レポートではなく、認定でもありません。仕事のほとんどは証拠を集めて、あなたが一年中同じことをし続けたことを示すことですが、Vanta のようなツールがその証拠を収集します。彼らはクラウド、コード、ラップトップからデータを取得し、監査人が要求するものと照らし合わせて、一度ではなく一年中監視します。彼らはあなたを従わせることはなく、報告書を書くこともありません。 Vanta は、独立したアウディを紹介すると自ら主張しています。

監査は依然として会計事務所によって行われています。
HIPAA US は、HIPAA が適用される場合、保護される健康情報を保護します。重要な場合 これは、医療プラン、医療情報交換所、特定の医療提供者、およびそれらに代わって保護された医療情報を扱う企業に適用されます。
多くの消費者向け健康アプリは HIPAA の対象になっていませんが、他の規則が適用される場合があります。 HHS はアプリや企業を HIPAA 準拠として認定しません。
ISO/IEC 27001 コードだけでなく、会社全体にわたってセキュリティを実行するための文書化された方法。重要な場合 コードだけでなく、社内のあらゆる場所でセキュリティが適切に実行されていることを顧客が確認したい場合が重要です。
認定の取得は任意であり、あなたではなく外部機関が行います。 ISO自体は企業を認証するものではありません。このガイドでは作業の一部のみを説明します。
ISO/IEC 42001 社内全体で AI を実行するための明文化された方法。したがって、すべてが偶然に任せられることはありません。重要な場合 製品に AI が使用されており、顧客が期待するのではなく、問題が発生する可能性について計画を立てていることを望んでいる場合が重要です。
認定の取得は任意であり、あなたではなく外部機関が行います。 ISO自体は企業を認証するものではありません。このガイドでは作業の一部のみを説明します。
PCI DSS ペイメント カード データを保護するための業界標準。重要な場合 アプリがカードの詳細を保存、処理、送信する場合、または作成したものがそれらの詳細の安全性に影響を与える可能性がある場合は重要です。
外部レジを使用すると作業は削減できますが、すべての責任を引き継ぐわけではありません。どのような証明が必要かを決済会社に問い合わせてください。
欧州連合で事業を展開している場合、欧州連合の人々に製品を提供している場合、または欧州連合での行動を追跡している場合は、チェックしてください。
米国の 13 歳未満の子供からオンラインで個人情報を収集する前に確認してください。
FTC の健全性侵害通知

ルールは、HIPAA の範囲外の健康アプリに適用できます。
構築時に使用する 60 個のチェック。
リストを下に向かって進みます。リスク、テスト、修正、証明のためにいずれかを開きます。
1 – 4 データとアクセス 4 つのチェック チェック 1 テーブルを構築する前にデータをマップする
余計なものを保管すると、漏れる可能性のあるものが 1 つ増えることになります。実際に使用しない電話番号は機能ではありません。それはただそこに座ってあなたを困らせるのを待っているだけです。
今まで構築できなかった機能のために電話番号を収集します。 2 年後、データベースが漏洩し、今度は電話番号が何の理由もなく漏洩することになりました。
誰かがあなたに自分のアカウントを削除するように要求しましたが、その人のデータがどこに行ったのかはまったくわかりません。
持っているすべてのテーブルとすべての列を書き留めます。
それぞれの横に、削除すると何が壊れるかを書き留めます。
何も壊れなければ、その必要はありません。何かを取り出す前にテーブルをバックアップし、一度に 1 列ずつ取り出します。
まず地図を作成します。それが何であるか、なぜ必要か、どこに存在するか、誰がそれを読み、誰が変更し、誰が削除するか、AI がそれを認識するかどうか、他の誰がそれを受け取るか、どのくらいの期間保管するか、そして誰かがそれをどのように削除するか。
正当化できないものは削除してください。
次に、その逆ではなく、マップからテーブルを構築します。
設定されているすべての列が、その横に理由とともにマップ上に表示されます。
これをクロード コード、コーデックス、カーソル、またはビルドに使用するものに貼り付けます。何かを変更する前に結果を求めるため、背後で何も書き換えられることはありません。
ログインしていることと許可されていることは同じではありません
ログインすると、建物の正面玄関から入ることができます。内部のすべてのアパートの鍵を渡すわけではありません。
ユーザー A は、/invoice/1043 で自分の請求書を開き、番号を 1044 に変更して、ユーザー B の請求書を読みます。
この間違いは、OWASP の API セキュリティ リストの一番上の項目です。それはディレクトリです

サインインしている人が別の人の情報にアクセスする方法など。
アカウントを 2 つ作ります。それらを A と B と呼びます。
まずBとしてログインします。注文やメッセージなど、B に属するものを 1 つ開きます。 Web アドレスは通常、数字で終わります。その番号を書き留めてください。
ログアウト。 A としてログインします。同じ種類のページを開き、アドレスに A の代わりに B の番号を入力します。
B の架空の数値ではなく、実際の数値を使用します。存在しない番号は「見つかりません」と返され、何もテストしなかった場合はパスしたように見えます。
A が B のものを見ることができれば、それはバグです。そのページのボタンも試してみてください。編集、保存、削除。機能するものはすべて、同じバグのさらに悪いバージョンです。毎回エラーまたは空のページが表示されるはずです。
サーバー上では、何かを返す前に、質問者が実際にその行を所有していることを証明してください。
すべてのエンドポイントでそれを実行します。エンドポイントは、アプリがリクエストに応答する単一のアドレスです (/invoice や /account など)。 1 つでも欠けていると、すべてが欠けているのと同じです。
推測しにくい ID は多少は役に立ちますが、それ自体では何も止まりません。推測するのは難しいですが、保護されているのと同じではありません。
ユーザー A としてログインすると、どの ID を入力しても、ユーザー B に属するものを読み取り、変更、削除することはできません。
壊れたアクセス制御 · IDOR · 安全でない直接オブジェクト参照 · BOLA · 壊れたオブジェクトレベル認証
これをクロード コード、コーデックス、カーソル、またはビルドに使用するものに貼り付けます。何かを変更する前に結果を求めるため、背後で何も書き換えられることはありません。
現時点で邪魔になっているのはアプリのコードだけで、急いで新しい画面を追加してチェックを忘れる日まで機能します。データベースにもルールがあれば、忘れても致命的ではありません。一部のシステムでは、これを行レベル セキュリティ (RLS) と呼びます。これで、人々がそれを何を意味するかがわかりました。
あなたは追加します

急いで新しい画面を表示し、所有権のチェックを忘れると、静かに全員の行を返します。
サーバー向けのキーがブラウザーに送信されるため、ルールを邪魔することなく誰でもテーブルを直接クエリできるようになります。
各テーブルについて、通常のサインイン者がアプリを経由せずにデータベースに直接要求した場合に何が返されるかを尋ねます。それを試す方法がわからない場合は、設定を読むことで、以下のプロンプトで同じ答えが得られます。
誰かのデータを保持するすべてのテーブルに対してルールをオンにします。
機能する最小のルールを作成します。自分が所有する行だけが表示され、他には何も表示されません。
3 人の異なる人物としてテストしてください。サインアウトしたユーザー A とユーザー B。
ユーザー A としてクエリすると、テーブルはユーザー A の行のみを返します。サインアウトすると何も返されません。
これをクロード コード、コーデックス、カーソル、またはビルドに使用するものに貼り付けます。何かを変更する前に結果を求めるため、背後で何も書き換えられることはありません。
自分自身のデータ、破ってはいけないルール
最後の 2 つのチェックは、他人のものに到達することに関するものでした。これは違います。アカウントは彼らのものであり、命令も彼らのものですが、それでも彼らがそのようなことをすることは許されるべきではありません。
どの製品にもそのようなルールがあります。注文ごとに 1 つの割引。誰かがすでに持っているスロットを予約することはできません。

[切り捨てられた]

## Original Extract

A free blueprint for turning something you built with an AI coding tool into software that is actually ready for real people. Data, permissions, business rules, uploads, accounts, payments, AI safety, claims, privacy, launch, recovery.

Lukata Skip to the blueprint LUKATA Start Checks Guides Launch About Free to read. Free to use.
Thinking of building something with AI? Find out the important things it usually misses!
Inversion thinking means working a problem backwards. You start with how it could fail rather than how it could succeed.
Use inversion thinking on your idea. Think about how it could fail, and what would make someone leave your site within five seconds. Figure out the issues, iterate fixes. Stress-test your own ideas before building. Try to break your own software by purposely trying to make it fail until it doesn't anymore. Make sure everything on your landing page is true and factual.
Your first hour starts here. These catch private data leaks and runaway bills.
Not going to read 60 checks? Paste this into your AI and it will work out which ones apply to what you built.
1 Can one person open someone else’s stuff? This is the most common way one customer ends up seeing another customer’s private information. Step 1 Check it yourself Make two accounts. Call them A and B.
Log in as B first. Open one thing that belongs to B, like an order or a message. The web address usually ends in a number. Write that number down.
Log out. Log in as A. Open the same kind of page, then put B’s number in the address instead of A’s.
Use B’s real number, not a made up one. A number that does not exist comes back not found, and that reads like a pass when you tested nothing.
If A can see B’s thing, that is the bug. Try the buttons on that page too. Edit, save, delete. Anything that works is a worse version of the same bug. A should get an error or an empty page every time.
2 Is a password or key sitting in your browser code? If it is, anyone can read it in about two clicks. Step 1 Check it yourself Search the code that runs in the visitor’s browser for anything that looks like a password or a key. Then search your git history too. Git history is the saved copy of every version your code has ever been, so a line you deleted last year is still sitting in there.
List every decision the visitor’s browser makes on its own. Look for anything about money, about what somebody is allowed to do, and about who can see what.
3 Can one person run your bill up overnight? Nobody has to steal anything. They only have to make you pay for the work. Step 1 Check it yourself Ask what one person could make your app do in a minute with no account at all, then work out what that costs you.
If your product takes money, do check 13 . If its AI reads anything you did not write, do check 17 .
5 things to check before real people use your app.
These are easy to miss even when every feature appears to work.
User input cannot rewrite a database request
User input cannot turn into browser code
Other sites cannot send commands as a signed-in user
Sign-in cookies are hard to steal or reuse
Admins use a second sign-in step
Login only returns people to approved pages
Error pages hide private technical details
Files meant to be private are not public
The live app has its own secret keys
The app connects only to approved addresses
Replies from other services are checked before use
Connections time out and redirects stay controlled
Only approved people can publish changes
Code is checked for leaked keys before release
A backup has been restored successfully
Several customers share one system Put a check for which customer this is inside every database query, job, cache, search and file path.
Social login Make sure the login actually started in your app, using a one-time code called PKCE. Only send people back to web addresses you listed yourself.
AI searches private documents Make the search itself look at one customer’s documents only. Lock down wherever those documents live. Show the source next to every answer.
AI takes actions An AI feature that takes actions is often called an agent. Keep its work in its own space. Cut down what it can reach on the internet. Approve anything it wants to remember, and anything big it wants to do. Keep a fast way to stop it.
AI handles sensitive customer data Use level 2 of OWASP’s AI app checklist, linked below.
Want a deeper review? OWASP is a nonprofit that publishes free security checklists. Start with its web app checklist. Add its AI checklist if your finished product uses AI.
Being safe and proving it are two different jobs.
First find out which of these applies to you. Then get the right review. This guide can help with part of the work. It cannot certify your company, and it cannot decide the law for you.
SOC 2 An outside expert checks how your company protects information, then writes a report. When it matters It often matters when a business customer wants proof from somebody other than you about how your company protects information.
A licensed accountant, called a CPA, does the checking and writes the report. This guide may help you get ready. It is not a SOC 2 report and it is not certification. Most of the work is gathering proof and showing you kept doing the same things all year, and tools like Vanta gather that proof for you. They pull it from your cloud, your code and your laptops, line it up against what an auditor asks for, and watch it all year instead of once. They do not make you compliant and they do not write the report. Vanta says itself that it introduces you to independent auditors, and the accounting firm still runs the audit.
HIPAA US safeguards for protected health information when HIPAA applies. When it matters It applies to health plans, health care clearinghouses, certain health care providers and companies handling protected health information for them.
Many consumer health apps are not covered by HIPAA, but other rules may still apply. HHS does not certify apps or companies as HIPAA compliant.
ISO/IEC 27001 A written way of running security across your whole company, not just your code. When it matters It can matter when a customer wants to see that security is run properly everywhere in your company, not only in the code.
Getting certified is optional, and an outside body does it, not you. ISO itself does not certify companies. This guide covers only part of the work.
ISO/IEC 42001 A written way of running AI across your company, so none of it is left to chance. When it matters It can matter when your product uses AI and a customer wants to see you planned for what could go wrong, instead of hoping.
Getting certified is optional, and an outside body does it, not you. ISO itself does not certify companies. This guide covers only part of the work.
PCI DSS An industry standard for protecting payment card data. When it matters It matters when your app stores, handles or sends card details, or when anything you built could affect how safe those details are.
Using an outside checkout can cut down your work, but it does not hand off every responsibility. Ask your payment company what proof you need.
Check it if you operate in the European Union, offer a product to people there or track their behavior there.
Check it before collecting personal information online from children under 13 in the United States.
The FTC Health Breach Notification Rule can apply to health apps that are outside HIPAA.
60 checks to use as you build.
Work down the list. Open any one for its risk, test, fix and proof.
1 – 4 Your data and access 4 checks Check 1 Map your data before you build tables
Every extra thing you store is one more thing that can leak. A phone number you never actually use isn’t a feature. It is just something sitting there waiting to embarrass you.
You collect a phone number for a feature you never got around to building. Two years later the database leaks, and now you have leaked phone numbers for nothing.
Somebody asks you to delete their account and you genuinely can’t say where all of their data went.
Write down every table and every column you have.
Next to each one, write down what breaks if you delete it.
If nothing breaks, you do not need it. Back the table up before you take anything out, and take one column at a time.
Do the map first. What it is, why you need it, where it lives, who reads it, who changes it, who deletes it, whether AI sees it, who else receives it, how long you keep it, and how somebody removes it.
Delete anything you can’t justify.
Then build your tables from the map, instead of the other way around.
Every column you have shows up on the map with a reason next to it.
Paste this into Claude Code, Codex, Cursor, or whatever you build with. It asks for findings before it changes anything, so nothing gets rewritten behind your back.
Being logged in isn’t the same as being allowed
Logging in gets you through the front door of the building. It doesn’t hand you a key to every apartment inside.
User A opens their own invoice at /invoice/1043, changes the number to 1044, and reads User B’s invoice.
That mistake is the number one item on OWASP’s API security list. It is a direct way for one signed-in person to reach another person's information.
Make two accounts. Call them A and B.
Log in as B first. Open one thing that belongs to B, like an order or a message. The web address usually ends in a number. Write that number down.
Log out. Log in as A. Open the same kind of page, then put B’s number in the address instead of A’s.
Use B’s real number, not a made up one. A number that does not exist comes back not found, and that reads like a pass when you tested nothing.
If A can see B’s thing, that is the bug. Try the buttons on that page too. Edit, save, delete. Anything that works is a worse version of the same bug. A should get an error or an empty page every time.
On the server, before you hand anything back, prove the person asking actually owns that exact row.
Do it on every endpoint. An endpoint is any single address your app answers requests on, like /invoice or /account. Missing one is the same as missing all of them.
IDs that are hard to guess help a little, but on their own they stop nothing. Hard to guess isn’t the same as protected.
Logged in as User A, you can’t read, change or delete anything belonging to User B, no matter what ID you put in.
Broken access control · IDOR · Insecure direct object reference · BOLA · Broken object level authorization
Paste this into Claude Code, Codex, Cursor, or whatever you build with. It asks for findings before it changes anything, so nothing gets rewritten behind your back.
Right now your app code is the only thing standing in the way, and that works right up until the day you add a new screen in a hurry and forget the check. If the database has the rule too, forgetting isn’t fatal. Some systems call this row level security, or RLS. Now you know what people mean by it.
You add a new screen in a hurry, forget the ownership check, and it quietly hands back everybody’s rows.
A key meant for your server ends up in the browser, and now anyone can query the table directly with no rules in the way.
For each table, ask what a normal signed in person would get back if they asked the database for it directly, without going through your app. If you are not sure how to try that, the prompt below reaches the same answer by reading your settings.
Turn the rules on for every table that holds somebody’s data.
Write the smallest rule that works. You see the rows you own and nothing else.
Test it as three different people. Signed out, User A, and User B.
Queried as User A, the table hands back only User A’s rows. Signed out, it hands back nothing.
Paste this into Claude Code, Codex, Cursor, or whatever you build with. It asks for findings before it changes anything, so nothing gets rewritten behind your back.
Their own data, rules they still shouldn’t break
The last two checks were about reaching somebody else’s stuff. This one is different. The account is theirs, the order is theirs, and they still shouldn’t be allowed to do the thing.
Every product has rules like that. One discount per order. You can’t book a slot somebody already has.

[truncated]
