---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48524387"
title: "Story of How Im Running an Unlimited $6/Month AI Provider on 4x RTX 3090s"
article_title: ""
author: "yolo-auto"
captured_at: "2026-06-14T07:43:08Z"
capture_tool: "hn-digest"
hn_id: 48524387
score: 3
comments: 1
posted_at: "2026-06-14T05:11:33Z"
tags:
  - hacker-news
  - translated
---

# Story of How Im Running an Unlimited $6/Month AI Provider on 4x RTX 3090s

- HN: [48524387](https://news.ycombinator.com/item?id=48524387)
- Score: 3
- Comments: 1
- Posted: 2026-06-14T05:11:33Z

## Translation

タイトル: 4 台の RTX 3090 で月額 6 ドルの無制限 AI プロバイダーを実行している方法の物語
HN テキスト: この投稿は、私がどのようにして待機リストに載っている約 60 人の誇大宣伝者に無制限の LLM プロバイダーを立ち上げ、すぐに完全に機能不全に陥った死のループ モデルを彼らに提供し、どのようにしてほとんどの人が、非常に合理的に消滅したかについての物語です。しかし、それでも非常に親切な数人の人々がとにかく残ってくれたおかげで、私たちはプロジェクトを存続させ、まだかなり混沌としていますが、勢いを増してきました。少し補足しますと、AI エージェントの要点は、働き続けることだと私は信じています。ファイルを読み取り、再試行、検索、コーディング、要約、ツールの実行を行い、ジョブが完了するまでループする必要があります。雇用主が費用を支払っている場合、コストなど誰にも気にされませんが、私の個人的なお金や趣味となると、すべてのループが小さな経済的イベントのように感じられると、エージェントを使用する代わりにエージェントの子守を始めてしまい、面白くありません。一方で、従量制料金では使いすぎが心配になります。使用量サブスクリプションを使用すると、最後の魔法の % をすべて使用する必要があるように感じられます。そうしないと、「無駄」になってしまいます。無制限のプロバイダーが存在していれば... その後、私は AMD 開発者プログラムに参加しました。独自の MI300x を立ち上げるためのクレジットを取得し、AMD で提供される vllm/sglang 推論をいじり始めました。 AMD MI300x について学んだ後、私はちょっと計算してみました。MI300x を 1 時間 2.00 でレンタル = 月額 ~1500 ドルです。 qwen-35b-3a などの小規模な MOE モデルを使用すると、おそらく約 150 人のユーザーをサポートできるか、おそらくそれ以上のユーザーをサポートできます。 1500 / 150 = 月額 $10.00 で、私たちは皆、少額の料金でエージェントと遊ぶことができます。少しオーバーサブスクライブできるため、2x 世代スロット、128k コンテキスト、トークン制限なし、レート制限なしで、ユーザーあたり月額 6 ドルに落ち着きました。サイトとルーターを構築し、待機リストを作成してから、MI300x を vllm ベンチが最適化されるまで過剰に最適化しました。

3k+ の出力と 40k+ のスループット.... しかし、最終的な config/serve コマンドをテストしませんでした... そして、そこで災害起動を行いました。ループしたりバグったりせずにプロンプ​​トを表示することはできません。それは呪われていました。そしてそこで私たちは多くの人を失いました。幸いなことに、私の友人は 3090 を数台持っていたので、私に救命ボートを投げ、2x 3090 で qwen をホストし始めました。そしてついに、なんと 3 人のユーザーに対して 1 時間あたり 2.00 ドルもかからない運用モデルが完成しました。ユーザーが増え始めたので、最大 4x 3090 に移行しました。これには、より多くのユーザーを受け入れる余地が十分にありますが、それでも、それ以来、vllm の設定を 15 回ほど間違っています。
GPUが死んでしまった
私たちは力を失った
openclaw、hermes、pi-mono をワンクリックで起動できるようにたくさんしましたが、どれも実際には正しく動作せず、おそらく人々が遠ざかってしまうでしょう。それらは現在も私たちのサイトにあります。 ...しかし、自分たちが何をしているのかを知っている人々は、この価格帯を本当に気に入っているようです。全体として、稼働時間は 98% ほどです。 1ヶ月ほど経ちました。私たちは二人とも多くのことを学びましたが、すでに SWE/SE/AI のバックグラウンドを持っていても、数人の有料ユーザーを抱えているため、彼らに良い製品を提供することに真剣に集中する必要がありました。そして今、電力/ホスティング料金の支払いが近づいているので、赤字ではないのではないかと思います (3,090 の設備投資を含めると、まだ赤字です)。私たちの損益分岐点は、MI300x を最大限に活用するためにクラウドに移行することであり、MI300x は現在調整されており、ユーザーを獲得したらすぐに使用できるようになりました。そして、一部の地域では、モデルを実行するよりもサービスに登録する方が安いことに気づきました（しかし、ローカルモデルを愛する者として、私はそれを完全に理解しています）。それ以来、私は qwen のような小規模なモデルで実際に動作するデスクトップ エージェントの開発に取り組んできました。これは、壊れた 1 クリック スタートに代わるものです。必要最小限の機能ですが、すぐに使える機能を備えています。できました

オープン ソースです。ここで私が何を話しているのかをご覧ください: https://github.com/yolo-auto-org/yolo-auto-desktop 。私たちは yolo-auto.com にいますが、それが機能することを証明するために、とんでもない無料枠があります。とにかく、笑ったり、面白かったりしていただければ幸いです。質問があればお書きください。

## Original Extract

This submission is a tale about how I launched an unlimited LLM provider to about 60 hyped people on the waitlist, then immediately served them a fully dysfunctional death-loop model, and how most people, very reasonably, disappeared, but thanks to a few extremely nice people stuck around anyway, we kept the project alive and its still pretty chaotic but gaining traction. To back up a little bit-- I believe that the whole point of AI agents is that they should keep working. They should read files, retry, search, code, summarize, run tools, and loop until the job is done. When your employer is paying for it, who cares about cost, but when it comes to my personal money/hobbies, if every loop feels like a tiny financial event, you start babysitting the agent instead of using it, and its not fun. On the other hand, metered pricing makes me worry about using too much. Usage subscriptions make me feel like I need to use every last magical % or I'm are "wasting it". If only an unlimited provider existed.... Then I joined the AMD developer program - I got some credits to spin up my own MI300x and started tinkering with vllm/sglang inference serving on AMD. After learning about AMD MI300x , i did some napkin math: Renting MI300x at 2.00 an hour = ~$1500 a month . It can probably support about 150 users using a small MOE model, like qwen-35b-3a , maybe more. 1500 / 150= $10.00 per month, and we all get to play with agents for a small price. You can oversubscribe a bit, so i landed on $6 per month, per user, for 2x generation slots, 128k context, no token limits, no rate limits. I built the site, router, made a waitlist, and then over-optimized the MI300x to the point where vllm bench had like 3k+ output and 40k+ throughput.... But i didn't test the final config/serve commands... And that's where i did a disaster launch. You couldn't prompt the thing without it looping or bugging out, it was cursed. And that's where we lost alot of people. Luckily, my buddy had a few 3090s, so he threw me a life boat and began hosting qwen for us on 2x 3090s and we finally had an operational model that wasn't costing $2.00 an hour for our whopping 3 users. We started gaining a more users, so we moved up to 4x 3090s. Which we have plenty of room for more users, but even so, since then: we've configured vllm wrong like 15 times
a GPU died
we lost power
I made a bunch of one-click starts for openclaw,hermes,pi-mono and none of them really work right and that probably drives people away. Those are still on our site right now. ...but people that know what they are doing seem to really be liking the price point. All in all we have like 98% up time. Its been about a month. We've both learned a ton, even already having backgrounds in SWE/SE/AI , being on the hook for a couple paying users forced us to really focus on delivering them a good product. And now i think we might be close to paying the power/hosting bill so we're not operating at a loss (if u include 3090 capex were still at aloss). Our break-even point is moving to the cloud to max out a MI300x, which is now tuned and ready to go once we get the users. And im finding in some areas, subscribing to our service is cheaper than running the model (but as someone who loves local models, i totally get it). Since then, I've been working on a desktop agent that actually works with small models like qwen -- thats going to replace the broken 1 click starts. It's barebones, but its something out of the box that just works. I made it open source, you can see what im talking about here: https://github.com/yolo-auto-org/yolo-auto-desktop , we're at yolo-auto.com and we have an abysmal free tier to prove it works! Anyway, hope you got a laugh or found it interesting! Drop a question if you have any.

