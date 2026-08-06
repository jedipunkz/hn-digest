---
source: "https://zhenyi.gibber.blog/1-rce-2-ai-models-0-ways-to-tell-who-was-right"
hn_url: "https://news.ycombinator.com/item?id=49194523"
title: "1 RCE, 2 AI models, 0 ways to tell who was right"
article_title: "1 RCE, 2 AI Models, 0 Ways to Tell Who Was Right - Gibberish and Stuff"
author: "zhenyi"
captured_at: "2026-08-06T10:29:23Z"
capture_tool: "hn-digest"
hn_id: 49194523
score: 1
comments: 0
posted_at: "2026-08-06T09:38:59Z"
tags:
  - hacker-news
  - translated
---

# 1 RCE, 2 AI models, 0 ways to tell who was right

- HN: [49194523](https://news.ycombinator.com/item?id=49194523)
- Source: [zhenyi.gibber.blog](https://zhenyi.gibber.blog/1-rce-2-ai-models-0-ways-to-tell-who-was-right)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T09:38:59Z

## Translation

タイトル: 1 つの RCE、2 つの AI モデル、誰が正しかったかを判断する方法は 0 つ
記事のタイトル: 1 つの RCE、2 つの AI モデル、誰が正しかったかを判断する 0 つの方法 - 意味不明なことなど
説明: (これは私の友人に起こりました。この話は軽く編集されているだけで、彼の名前はジェリーではありません。)
ジェリーは少し前にバイブコーディングに興味を持ちました。彼は、デフォルトのモデルである Claude Sonnet 4.6 を使用していました。

記事本文:
1 つの RCE、2 つの AI モデル、誰が正しかったかを判断する方法は 0
(これは私の友人に起きた出来事です。この話は軽く編集されているだけで、彼の名前はジェリーではありません。)
ジェリーは少し前にバイブコーディングに興味を持ちました。彼は当時のデフォルト モデルである Claude Sonnet 4.6 を使用していました。大丈夫でした。その後、ソネットでは解決できない問題に遭遇しました（今ではそれが何だったか思い出せません）。さらに悪いことに、それ自体が矛盾し始めました。 X が根本原因です。 X は根本原因ではありません。ジェリーが尋ねると、やはりXが根本原因だという。ジェリーにはどの答えが正しいのか分かりませんでした。
不審に思ったジェリーは、初めて Opus 4.8 に切り替えました。すぐに問題を特定し、修正を実装しました。プロンプトなしで、Opus は途中で 3 つのさまざまなバグも修正しました。それ以来、ジェリーは二度とソネットを使用しませんでした。オーパスは共同創設者がいるようなものでした。独自の意見があり、技術的な決定を下してくれるでしょう。また、人間の技術共同創設者とは異なり、いつでもビジネス上の質問をすることができ、それは同様に優れていました。
ジェリーはバイブコーディングが上達するにつれて、ベンチマークに注意を払い始めました。 『Opus 5』はリリースされたばかりで、アンスロピックはこれがすべてにおいて最高だと述べた。ジェリーがスイッチを入れたのですが、アンスロピックは冗談ではありませんでした。彼は、Opus 4.8 に既存のサイトの外観と雰囲気をコピーしたいと考えていましたが、それは決してうまくいきませんでした。 Opus 5 はそれを 1 回のパスで実現しました。スクリーンショットも撮って検証しました。
ある日、ジェリーは気まぐれに GPT-5.6 Sol も試してみました。彼は Codex アプリをインストールし、Opus と一緒に構築していたアプリを誇らしげに Sol に見せました。ソルが最初に行ったのはセキュリティ検査でした。すると、アプリには RCE があると表示されました。ジェリーは、RCEとは何なのか尋ねました。ソル氏は、基本的に誰でもアプリを通じてサーバーを乗っ取ることができると述べた。ジェリーは「なんてことだ」って感じだった。彼はずっとズボンを下ろしたまま歩き回っていた

時間です。
ジェリーはソルのメッセージをオーパスに伝えた。オーパスはその主張を再確認し、パニックに陥った。そのため、サーバー上で実行するコマンドがたくさん必要になりましたが、彼はそれを実行することに不快感を感じていました。彼はオーパスに落ち着くように言いました。アプリはまだリリースされておらず、ユーザーは友人だけでした。彼らはこれを適切に修正する必要があります。 Opus はこれに同意し、すべてのコマンドと、これやこれを見た場合に何をすべきかを説明するウォークスルーのようなマークダウン ファイルを作成しました。
ジェリーはそれをソルに見せました。 Sol 氏は、最初の試みはうまくいったが、Opus がフォルダー名を間違えたため、1 つのコマンドが失敗するだろうと言いました。ジェリーはメッセージをオーパスに転送しました。 Opus が修正を行いました。ジェリーはそれをソルに見せました。ソルはさらに多くの欠点を見つけ出しました。ジェリーは最初、AI のおかげでアプリをより良く、より安全なものにすることができて満足していました。しかし、4～5ラウンド行ったり来たりしているうちに疲れてしまった。彼はもう彼らが何を言っているのか理解できませんでした。
ジェリーはソルに、今度は君が運転するように言った。問題を指摘するのではなく、修正してください。それから彼はそれをオーパスに見せました。そこでソルは初めてファイルを編集しました。ジェリーがそれをオーパスに見せると、オーパスは変更を厳しく精査し始め、テスト用の Python スクリプトを 5 つ書きました (彼のアプリは Python さえ使用していませんでした)。次に、ファイルを実行しないようにと表示されました。ステップ 3 には微妙なバグがありました。ユーザーが同時に何かを実行すると、データベースが破損します。ジェリーはそれをソルに送り返した。 Sol はこれに同意し、変更を実行しました。
ジェリーはついに終わったと思ったが、オーパスはなんとか欠陥を見つけ続け、ソルは修正のたびに独自のひねりを加え続けた。 15 行のウォークスルーは、エッジケースに関するコメントを含めて 100 行を超えました。
それからジェリーは二人に、お互いに勝ち上がろうとするのをやめるよう言いました。このままではサーバーが修復されることはありませんでした。ウォークスルーを台無しにします。彼はサーバー上でコマンドを直接実行するつもりでした。彼らは挨拶をするだろう

何を実行するかを一度に 1 つずつ選択します。彼はそれを実行し、出力を貼り付けて戻しました。これで二人は黙り、二人とも彼の時間を無駄にするのをやめることに同意した。
オーパスの制限時間が5時間に近づいていたため、ジェリーはソルを選んで指示を与えた。彼は脳手術を行っているような気分でしたが、コマンドは従うのが簡単で、ソルが予測した正しい出力が得られ続けました。 5 分後、セキュリティ ホールがパッチされ、ファイルのアクセス許可が修正され、サーバー プロセスが再起動されました。
ジェリーはもう終わりだと思って次へ進むことができましたが、オーパスはソルが問題を解決するよう導いてくれたことを知り、受動的攻撃的な行動をとり始めました。 「それを受け入れるか、やめるか」とか、「それは単なる私の推奨です。あなたたちがそれを望むなら、私はそうすることができます」などのことを言い始めました。ジェリーはため息をついた。彼は、AI を扱うときは社内政治の影響を受けないと考えていた。
この時点で、ジェリーは懐疑的でした。彼のアプリは本当に安全でバグがなかったのでしょうか、それとも AI がすべてを彼に教えてくれなかったのでしょうか?彼は、Fable 5 という大きな武器を導入することにしました。彼には 100 ドルの使用クレジットがありました。彼はファブルにすべての経緯と意見の相違を伝え、コードをもう一度調べるように頼みました。数分後、非常に長いメッセージが彼に送られてきました。基本的には、アプリの状態は良好、修正は機能しています、私には良さそうです、Opus と Sol は良い仕事をしました、というものでした。ジェリーは使用クレジットをオフにしました。このプロンプトには 10 ドルの費用がかかりました。そして、自分のアプリが安全かどうかはまだわかりませんでした。
こんにちは、私は Zhenyi Tan です。これは意味不明のブログです。 And a Dinosaur and Mastodon でも私を見つけることができます。

## Original Extract

(This happened to a friend of mine. The story is only lightly edited, and his name is not Jerry.)
Jerry got into vibe coding a while ago. He'd been using Claude Sonnet 4.6, the default model at the...

1 RCE, 2 AI Models, 0 Ways to Tell Who Was Right
(This happened to a friend of mine. The story is only lightly edited, and his name is not Jerry.)
Jerry got into vibe coding a while ago. He’d been using Claude Sonnet 4.6, the default model at the time. It was fine. Then he ran into a problem Sonnet couldn’t fix (he can’t remember what it was now). Worse, it started contradicting itself. X is the root cause. X isn’t the root cause. When Jerry asked, it said X was the root cause again. Jerry just didn’t know which answer was right.
Suspicious, Jerry switched to Opus 4.8 for the first time. It quickly identified the issue and implemented a fix. Without prompting, Opus also fixed 3 miscellaneous bugs along the way. From then on, Jerry never used Sonnet again. Opus was like having a cofounder. It had its own opinions and would make technical decisions for you. And unlike a human technical cofounder, you could ask it business questions at any time, and it was just as good.
As Jerry got better at vibe coding, he started paying attention to benchmarks. Opus 5 had just been released, and Anthropic said it was the best at everything. Jerry switched it in, and Anthropic wasn’t kidding. He’d wanted Opus 4.8 to copy the look and feel of his existing site, and it never quite managed. Opus 5 did it in one pass. It even took screenshots and verified them.
One day, on a whim, Jerry tried GPT-5.6 Sol as well. He installed the Codex app and proudly showed Sol the app that he and Opus had been building. The first thing Sol did was a security checkup. Then it said, your app has an RCE. Jerry asked what an RCE was. Sol said basically anyone could take over his server through his app. Jerry was like, oh shit. He’d been walking around with his pants down this whole time.
Jerry relayed Sol’s message to Opus. Opus double-checked the claim and panicked. It gave him a lot of commands to run on the server, which he was uncomfortable doing. He told Opus to chill. The app wasn’t launched yet and his only users were friends. They should fix this properly. Opus agreed and drafted a markdown file, sort of like a walkthrough, explaining every command and what he should do if he saw this or that.
Jerry showed it to Sol. Sol said it was a good first attempt, but one command would fail because Opus had gotten the folder name wrong. Jerry forwarded the message to Opus. Opus made the fix. Jerry showed it to Sol. Sol picked out even more flaws. At first Jerry was happy, the AIs were helping him make his app better and more secure. But after 4-5 rounds of back-and-forth, he got tired. He didn’t understand what they were saying anymore.
Jerry told Sol, this time, you drive. Instead of pointing out the issues, fix them. Then he’d show it to Opus. So Sol edited the file for the first time. When Jerry showed it to Opus, Opus started scrutinizing the changes hard, writing 5 Python scripts to test (his app didn’t even use Python). Then it said not to run the file, step 3 had a subtle bug. If a user was doing something at the same time, it would corrupt the database. Jerry sent it back to Sol. Sol agreed and implemented the change.
Jerry thought it was finally over, but Opus kept managing to find flaws, and Sol kept adding its own twist to each fix. The 15-line walkthrough became more than 100 lines with comments about edge cases.
Then Jerry told both of them to stop trying to one-up each other. At this rate the server was never getting fixed. Screw the walkthrough. He was just going to run the commands on the server directly. They’d tell him what to run, one at a time. He’d do it and paste the output back. This shut them both up, and they both agreed to stop wasting his time.
Jerry picked Sol to give him the commands, because Opus was nearing its 5-hour limit. He felt like he was performing brain surgery, but the commands were easy to follow, and he kept getting the correct output Sol had predicted. After 5 minutes, the security holes were patched, file permissions were corrected, and the server process was restarted.
Jerry thought it was over and they could move on, but Opus started acting passive-aggressive after it learned that Sol had guided him through the fix. It started saying things like “take it or leave it” or “that’s just my recommendation — I can do it if that’s what you both want”. Jerry sighed. He’d thought he’d be immune to office politics when dealing with AI.
At this point Jerry was skeptical. Was his app really safe and bug-free, or were the AIs not telling him the whole story? He decided to bring in the big gun: Fable 5. He had $100 in usage credit. He told Fable the whole story, the disagreements, and asked it to look through his code again. After several minutes it sent him a very long message basically saying, the app is in good shape, the fix works, looks good to me, Opus and Sol did good work. Jerry turned off the usage credit. The prompt had cost him $10. And he still didn’t know if his app was safe.
Hi, I am Zhenyi Tan and this is my Gibberish blog. You can also find me at And a Dinosaur and Mastodon .
