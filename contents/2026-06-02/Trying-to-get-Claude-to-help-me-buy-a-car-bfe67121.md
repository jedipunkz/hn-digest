---
source: "https://developerwithacat.com/blog/202606/ai-car-search/"
hn_url: "https://news.ycombinator.com/item?id=48358966"
title: "Trying to get Claude to help me buy a car"
article_title: "Trying to get Claude to help me buy a car"
author: "mmarian"
captured_at: "2026-06-02T00:36:58Z"
capture_tool: "hn-digest"
hn_id: 48358966
score: 3
comments: 3
posted_at: "2026-06-01T16:22:55Z"
tags:
  - hacker-news
  - translated
---

# Trying to get Claude to help me buy a car

- HN: [48358966](https://news.ycombinator.com/item?id=48358966)
- Source: [developerwithacat.com](https://developerwithacat.com/blog/202606/ai-car-search/)
- Score: 3
- Comments: 3
- Posted: 2026-06-01T16:22:55Z

## Translation

タイトル: クロードに車の購入を手伝ってもらおうとしている
説明: AI を使用して Autotrader を検索し、MOT 履歴を確認しようとしましたが、古いデータ、ボット ブロック、CAPTCHA が邪魔をしていました。 Chrome拡張機能がついに機能しました。

記事本文:
クロードに自動車開発者の購入を手伝ってもらいたいと考えていますWithACat RSS |
ニュースレター クロードに車の購入を手伝ってもらおうとしています
現在の住まいに引っ越したとき、車がないと半年も持たないと言われました。 2年経って、ついに諦めました。
受け入れられるのは苦痛だったが、私たちの AI の支配者が探索を行ってくれるという考えに慰めを求めた
素早く簡単に。私のニーズを入力するだけで、なんと、理想的なリストがいくつか出てきて、その中から勝者を選ぶことになります。
そこで、ある晩、私はソファに行き、猫の大君にスペースを空けてほしいと懇願し、ほんの小さなパッチしかもらえないことを受け入れました。
彼は感謝の気持ちを込めて押し入り、クロードに私の意図通りにできるように心を開いてくれました。
まずは私の要件に合ういくつかのモデルを勧めてくれました。
とても役に立ちました！車に関しては全くの無知でした。
それから私は貪欲になりました。私はクロードに Autotrader から実際のリストを取得してレビューするよう依頼しました。分析は難航しました、リンクは
与えられた。いずれかをクリックすると、検索ページにリダイレクトされます。別の同じ結果をクリックします。 3回目と4回目のトライについて。
そこで私はクロードに「一体何が起こっているの？？」と尋ねました。 - 丁寧な言い方で、もちろん、私がいつその使用人になるかわかりません。
実際には Autotrader のライブストックを確認できないことが判明しました。検索ページの古くてキャッシュされたスナップショットを読み取っていました。
また、Autotrader はこれらのページからのボットを完全にブロックしているため、個々の広告を開くこともできませんでした。
{
"error_type" : "ROBOTS_DISALLOWED" ,
"error_message" : "サイトでは自動アクセスが許可されていません。検索結果のスニペットを使用するか、代替ソースを見つけてください。"
}
クロードが web_fetch ツールを使用してリストを開くときに受け取るペイロード
ただし、これを回避する方法があります。Claude を実行することです。
Chrome拡張機能
Autotrader の検索結果ページ。
でも待ってください！もう一つ確認しなければならないことがあります。MOTの結果です。

国王陛下のウェブサイトから無料で入手できます。
そこで私はクロードにこのプロンプトを出しました…
この Autotrader の結果ページから各車を金額に見合った価値で評価してください。
また、各リストの写真からその登録番号を取得してから、
https://www.check-mot.service.gov.uk/results?registration={取得した登録番号を入力} にアクセスして、MOT 結果を取得します
そこから。それを評価に使用してください。
…そして、魔法をかけてみましょう。結果はかなり良いように思えました…応答の一番下を見てみるまでは、
注意: gov.uk サイトの CAPTCHA が作動したため、クロードはほとんどの MOT を取得できませんでした。
幸いなことに、私たちの偉大な王国は、この特定の目的のための API を提供しています。
そこで、API 資格情報を使用してローカル プロキシを設定し、クロードにそれを呼び出させます。
ただし、このワークフローは 1 つのモデルの検索結果のみを対象としています。十数回実行したかった。プロンプトを調整すればよかった
すべてを一度に実行したいのですが、実験を通じて、拡張機能が頻繁に失敗することに気づきました。それで私はそれを実行することにしました
モデルごとに手動で行うことは、いずれにしてもそれほど面倒ではありませんでした。
最後にもう 1 つ驚きがありました。会話は実際にはメインのクロード チャットに送信されません。それが欲しかっただろう
すべてのリストからトップ 10 を選択します。私にとって、それは小さな不便でした。コンテンツを
クロード・アイチャット。他の人はこれに満足していませんでしたが、これが最悪の状況を説明しています
2.6 評価
拡張機能は Chrome ウェブストアにあります。
おそらく、これが私の車を見つけるのにどのように役立ったかを私が言うことを期待しているでしょう。実はブロックを見てすぐに諦めてしまいましたが、
妻が見つけた素晴らしいリストに参加し、余裕ができてからこの作品の制作に戻りました。
これを再度実行する必要がある場合は、必ず最終ワークフローを使用します。でも、その時までに私は実際にできるようになるでしょう。
円周率

メインチャットインターフェースから車を確認してください!
話したいですか？私は LinkedIn を利用しています。接続時にメッセージを追加してください。あるいは、このお問い合わせフォームをご利用ください。新しいブログ投稿に関する通知を受け取りたい場合は、購読してください。

## Original Extract

I tried using AI to search Autotrader and check MOT history - stale data, bot blocks, and CAPTCHAs got in the way. Chrome extension finally worked.

Trying to get Claude to help me buy a car developerWithACat RSS |
Newsletter Trying to get Claude to help me buy a car
When I moved to my current abode, I was told I wouldn’t last 6 months without a car; 2 years later, I finally gave in.
The acceptance was painful, yet I sought comfort in the thought that our AI overlords would make the search
quick and easy. Just type in my needs and, poof, a few ideal listings would come up, from which I’d pick the winner.
So, one evening, I went to my couch, begged my feline overlord to make space, accepted that I’ll only get a tiny patch,
squeezed in gratefully, and opened up Claude to do as I intended.
It started off recommending me a few models that fit my requirements.
That was incredibly helpful! I was totally clueless about cars.
Then I got greedy. I asked Claude to pull actual listings from Autotrader to review. Analysis was crunched, links
were given. I click one: it redirects me to the search page. I click another, same result; as for the 3rd, and 4th try.
So I ask Claude “what the hell’s happening??” - in a polite way, of course, you never know when I’ll become its servant.
Turns out it can’t actually see Autotrader’s live stock. It was reading a stale, cached snapshot of the search page.
And it couldn’t open the individual ads because Autotrader blocks bots from those pages entirely.
{
"error_type" : "ROBOTS_DISALLOWED" ,
"error_message" : "Site disallows automated access. Use the search-result snippet, or find an alternative source."
}
The payload Claude receives when it uses its web_fetch tool to open a listing
There is a way around it though: running Claude’s
Chrome extension
on Autotrader’s search results page.
But wait! There’s one more thing we need to check - the MOT results,
freely available from our royal majesty’s website .
So I gave Claude this prompt…
Assess each car from this Autotrader results page for value for money.
Also fetch its reg number from each listing photo, then go
to https://www.check-mot.service.gov.uk/results?registration={enter registration number you fetched} and get the MOT results
from there. Use that in your assessment.
…and let it do its magic. The results seemed pretty good…until I looked at the bottom of the response and saw the
caveat: Claude was unable to fetch most of the MOTs because the gov.uk site’s CAPTCHA kicked in.
Luckily our great kingdom offers an API for this particular purpose.
So I set up a local proxy with my API credentials, and have Claude call it.
This workflow only covers search results for one model though. I wanted to run it for a dozen. I would’ve adjusted the prompt
to do all in one go, but I noticed through my experiments that the extension frequently fails. So I decided to just run it
manually for each of the models, it wasn’t much of a hassle anyway.
There was one final surprise - the conversation doesn’t actually get sent to the main Claude chat. I would’ve wanted that
to pick the top 10 across all listings. For me, it was a minor incovenience: I just had to paste the content into the
claude.ai chat. Others weren’t as happy with this, which explains the abysmal
2.6 rating
the extension has on the Chrome Web Store.
You’re probably expecting me to say how this helped me find my car. In reality, I gave up shortly after seeing the blocks,
went with a great listing my wife found, and came back to making this work once I had the headspace.
When I have to go through this again, I’ll definitely use the final workflow. Although, who knows, by that time I can actually
pick the car from the main chat interface!
Want to talk? I'm on LinkedIn ; please add a message when you connect. Alternatively, use this contact form . And, subscribe if you want to be notified about new blog posts.
