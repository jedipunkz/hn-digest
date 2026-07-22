---
source: "https://numerlab.org/2026/07/22/im-canceling-claude-subscription/"
hn_url: "https://news.ycombinator.com/item?id=49005076"
title: "Im cancelling my Claude Code subscription"
article_title: "Im canceling my claude subscription - Numerator's Blog"
author: "wallach-game"
captured_at: "2026-07-22T12:23:16Z"
capture_tool: "hn-digest"
hn_id: 49005076
score: 7
comments: 5
posted_at: "2026-07-22T11:32:04Z"
tags:
  - hacker-news
  - translated
---

# Im cancelling my Claude Code subscription

- HN: [49005076](https://news.ycombinator.com/item?id=49005076)
- Source: [numerlab.org](https://numerlab.org/2026/07/22/im-canceling-claude-subscription/)
- Score: 7
- Comments: 5
- Posted: 2026-07-22T11:32:04Z

## Translation

タイトル: クロード コードのサブスクリプションをキャンセルします
記事のタイトル: クロードのサブスクリプションをキャンセルします - Numerator のブログ

記事本文:
分子のブログ
投稿
について
チェスキー
クロードの定期購入をキャンセルします
クロードの定期購入をキャンセルします
昔、トークンが足りなくなったことがありました。また。そこで、作業を続けるためにオープンコードを起動しました。しかし今回は DeepSeek V4 Flash を無料で選んだので、何が起こるかについての準備ができていませんでした。
deepseek は作業をより速く完了しました。私のワークフローでは、何かがどのように機能するかを理解し、ボットにソフトウェアの動作方法と作成方法の指示のリストを与えます。ディープシークはより良い結果をより速く提供します。クロード・ソネットも完了できましたが、より多くのトークンが使用され、結果は通常よりバグが多かったです。
コメントであなたが叫んでいるのが聞こえます。なぜ作品や寓話ではなくソネットなのでしょうか？理由は単純で、トークンの予算です。しかし、私の主張は正しいです。少なくとも私の日々の仕事では、小型モデルの方が大型モデルよりも優れています。ただし、それは私が指示を与えるときだけだと思います。これがまさに私の仕事のやり方です。私にはソフトウェアがどのように機能するべきかというビジョンがあり、エージェントにそれを実行してもらいたいと考えています。
そして私にとって最も重要なのは、このモデルがコンシューマ ハードウェア上でローカルに実行できることです。
クロードサブをキャンセルすべきだと言っているわけではありません。これは特に私のワークフローで機能します。
古いソフトウェアを書き直してみましょう — twitch ビューア アラート。これは次のように動作します。視聴者数を監視し、その数が変化すると通知を発行します。設定もいくつかありますが、それ以上は覚えていません。エージェントの仕事は、これを Chrome 拡張機能に変換することです。元々は、ブラウザのコンソールに貼り付ける単なるスクリプトでした。批判しないでください。これを書いたとき、私は 19 歳でした。
各モデルは自律モードで 1 回テストされました。指標は、トークンの使用量、時間、物事を進めるために介入しなければならなかった回数です。
こんにちは、私はこの古いプロジェクト https://github.com/wallach-game/twitch-viewer-alerter を持っているので、会話したいです

それをChromeアドオンにRTします。ところで、このリポジトリにはプッシュせず、ローカルのクローンリポジトリとして保持します。
時間: 4分
トークン: 46,796
費用: $0.000004
介入: 0
注: 7 つの新機能とアイコンを追加しました
視聴者数の追跡: 5 秒ごと → 3 秒ごと
アラート ボックス: ブラウザ通知に置き換えられました
ボタンによる通知（ストリームを開く/30分間ミュート）
Mutationobserver によるスパ ナビゲーションのサポート
3 つの戦略とフォールバックを備えた堅牢なセレクター
時間: 2分
トークン: ~562,910 (ほとんどがキャッシュ)
費用: $0.40
介入: 0
注: UI には元の拡張機能よりも内容が少なくなっています
評決: 古いコードをヤンクして不要なものを削除しただけです。しかし、ブラウザ通知はもう不可能であることがわかりました。うまくいきました。
時間: 1:37
トークン: ~148,500
コスト: $0.30
介入: 0
注: このバージョンには設定 UI がありません。ちょっと奇妙です。そして通知も機能しません。ローカルストレージにナンセンスな問題が発生し、最終的には機能しない可能性があると文字通り告げられました。また、Chrome 拡張機能の https がないとブラウザ通知が機能しないこともわかりませんでした。
それで、これが問題です。 deepseek は 4 分で完了し、11 個の新機能が追加され、基本的に費用はかかりません。 Sonnet はそれをより速く実行しました (2 分) が、出力はさらに悪く、機能が削除され、古いコードが取り除かれました。 opus は 1:37 で最速でしたが、結果は文字通りうまくいきませんでした。
数字がすべてを物語っています。 deepseek v4 フラッシュのコストは 0.000004 ドル、ソネットのコストは 0.40 ドル、オーパスのコストは 0.30 ドルです。私のワークフローでは、明確な指示を与えてモデルを実行させると、安価なモデルが勝ちます。
しかし、最終的にソネットはプロンプトの指示に従って、拡張機能に変換しました。そしてそれはできました。次回は、「最初に元の動作をコピーしてから、新しい機能を追加して磨きをかける」など、より良いプロンプトを試してみます。
サブウェイをキャンセルすべきだと言っているのでしょうか？いいえ。これ

それが私にとってうまくいくことです。しかし、私と同じように仕事をするなら、まず計画を立ててから AI に構築してもらいましょう。ディープシーク v4 フラッシュを試してみてください。驚かれるかもしれません。
また、ディープシークは、特に opencodes auto exec コマンドと組み合わせると、何度か失敗しました。実行中のサーバー アプリを修正するという考えは、バックアップなしで削除することでした。しかし正直に言うと、これは私の促しのせいかもしれません。
ディープシークの価格設定がこのレベルではおそらく持続可能ではないことは承知していますが、クロードのサブスクリプションについても同じことが言えます。しかし、私はここで政治的なことを言いたくありません。実際のところ、消費者向けハードウェア上でローカルに実行できるモデルでは、ディープシークがクロードよりも優れたパフォーマンスを発揮しました。それが重要な部分です。
© 2026 分子。無断転載を禁じます。

## Original Extract

Numerator's Blog
Posts
About
Česky
Im canceling my claude subscription
Im canceling my claude subscription
Once upon a time i was out of tokens. again. so i launched opencode to keep working. but this time i picked DeepSeek V4 Flash free, and i wasnt ready for what happened.
deepseek finished the work faster. for my workflow — figure out how something should work, give the bot a list of instructions how the software should work and how to create it — deepseek gave better results, faster. claude sonnet could finish too, but it used more tokens and the result was usually more buggy.
and i can hear you screaming in the comments — why sonnet and not opus or even fable? the reason is simple, token budget. but my point stands: the smaller model, at least from my day to day work, outperforms the bigger one. but only when im providing the instructions, i think. this is just how i work — i have a vision how software should work and i want my agent to execute it.
and the main deal for me is that this model can run locally on consumer hardware.
im not saying you should cancel your claude sub. this works especially for my workflow.
so lets rewrite my old software — twitch viewer alert. it works like this: it watches the viewer count, and when the number changes, it fires a notification. theres some settings too, but i dont remember much more. the task for the agents is to convert this into a chrome extension. originally it was just a script you paste into the browser console — dont judge me, i was 19 when i wrote it.
each model was tested once in autonomous mode. the metrics are: token usage, time, and how many times i had to step in to keep things moving.
hi i have this older project https://github.com/wallach-game/twitch-viewer-alerter and i wanna convert it into chrome addon, btw do not push to this repo, keep it as local cloned repo
time: 4 minutes
tokens: 46,796
cost: $0.000004
interventions: 0
notes: added 7 new features and an icon
viewer count tracking: every 5s → every 3s
alert box: replaced with browser notifications
notifications with buttons (open stream / mute for 30m)
spa navigation support via mutationobserver
robust selectors with 3 strategies and fallbacks
time: 2 minutes
tokens: ~562,910 (mostly cache)
cost: $0.40
interventions: 0
notes: the ui has less stuff than the original extension
verdict: imo it just yanked the old code and removed unnecessary stuff. but it figured out that browser notifications are not possible anymore — good catch.
time: 1:37
tokens: ~148,500
cost: $0.30
interventions: 0
notes: this version doesnt have a settings ui. kinda odd. and the notification doesnt even work — it created some localstorage nonsense and literally told me at the end it might not work. also, it never figured out that browser notifications dont work without https in chrome extensions.
so heres the thing. deepseek did it in 4 minutes, added 11 new features, and cost basically nothing. sonnet did it faster (2 minutes) but the output was worse — stripped features, yanked old code. opus was the fastest at 1:37 but the result literally didnt work.
the numbers speak for themselves. deepseek v4 flash cost $0.000004, sonnet cost $0.40, opus cost $0.30. for my workflow — give clear instructions and let the model execute — the cheap model wins.
but in the end sonnet did what the prompt said — convert into an extension. and it did it. next time ill try a better prompt, something like “copy the original behavior first, then add new features and polish”.
am i saying you should cancel your sub? no. this is what works for me. but if you work like i do — plan first, then let the ai build it — give deepseek v4 flash a shot. you might be surprised.
and also deepseek screwed me over a few times, especially paired with opencodes auto exec command. its idea of fixing a running server app was deleting it without backup. but honestly, this can be blamed on my prompting.
im aware deepseek pricing is probably not sustainable at this level, but imo the same can be said about claude subscription. but i dont wanna get political here. the fact is that deepseek outperformed claude with a model that can be run locally on consumer hardware. thats the part that matters.
© 2026 Numerator. All rights reserved.
