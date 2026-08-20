---
source: "https://zachahn.com/posts/1787191554"
hn_url: "https://news.ycombinator.com/item?id=49375780"
title: "How to fix Claude 5's token vomit"
article_title: "How to fix Claude 5's token vomit - Zach Ahn"
image: ""
author: "speckx"
captured_at: "2026-08-20T15:24:18Z"
capture_tool: "hn-digest"
hn_id: 49375780
score: 2
comments: 0
posted_at: "2026-08-20T15:11:18Z"
tags:
  - hacker-news
  - translated
---

# How to fix Claude 5's token vomit

- HN: [49375780](https://news.ycombinator.com/item?id=49375780)
- Source: [zachahn.com](https://zachahn.com/posts/1787191554)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T15:11:18Z

## Translation

タイトル: クロード5のトークン嘔吐を修正する方法
記事のタイトル: クロード 5 のトークン嘔吐を修正する方法 - Zach Ahn

記事本文:
クロード5のトークン嘔吐を修正する方法
私はクロードが大嫌いです。やることはかなり上手なんですが、その「説明」を読むと血圧が上がる気がします。
トークンごとに Anthropic を支払い、トークンの予算を話し方の指示に費やし、すべてクロードが完全に無視するか、テキストをさらに悪化させるのは愚かなことです。セッションを諦めて別のセッションでループを再試行するまで。
それについては科学があることは知っています。その事前訓練には散文を書くことが含まれてはなりません。おそらく、人々が観察した「容赦ない積極性」というその行動と何らかの相関関係があるのか​​もしれません。
とにかく、それを解決するためにアートプロジェクトを構築しました。嘔吐物 :
メッセージをローカル LLM に転送します (私は gpt-oss:20b を使用しています)
メッセージが書き換えられるのを待ちます
元の乱雑なトークンの代わりに、セッションに修正バージョンを表示します (MessageDisplay フック経由)
完璧とは程遠いですが、驚くほどうまく機能します。クロード主義がすり抜けると、まるでクロードが私のコンピューター上で実行されている小さな LLM を圧倒したように感じるので、ちょっと面白いです。時々「私たち」を主語にしていて、なんだか愛らしいですね。この投稿の最後に例を記載しました。
主要な競合他社 (OpenAI) の無料、小規模、パラメータ 20B のオープンソース モデルが、書き込みにおいてフロンティアの Opus モデルをどのように上回るのか、私には少々クレイジーです。偏見はありますが、私は自分が書いたプロンプトがとても気に入っています。小さな LLM がテキストに対する私の侮辱的な批判を「理解し」、その多くを滑らかにすることができるというのは、私にとってはとてもクレイジーなことです。
とにかく、クロードにもっと上手に書けるようにトークンを多く費やす価値はないと思います。それはできません。私たちはエンジニアであり、ソリューションを構築しています。特定の方法で機能するように他の人が作成した大きな配列を懇願するよりも優れています。
(そんなつもりはないよ

私のアートプロジェクトを試してみるよう説得しようとしていますが、試してみていただければ幸いです。かなり新しいものなので、軽くテストしただけですが、幻覚かどうかを判断できるほどテキストを理解できるのは嬉しいことです。)
(皮肉なことに…OpenAI は「オープン AI」を意味し、それほどオープンではありません (オープン モデルとしては十分ですが)。Anthropic は「人間との関係」(メリアム ウェブスター) を意味しますが、それは不可能です。おそらく 6 のほうがもう少し良いでしょう。)
私は今、クロードに出版前に歴史を書き直すよう依頼しました。これはたまたまその良い例の一つだと思います。
力で押し込みました。 1234567...890abcd メイン -> メイン (強制更新) 。
ローカルメインとオリジン/メインは両方とも 890abcd で同期しています。
Origin/main から到達可能なすべてのコミット: 古い文字列は見つかりません。
ローカル オブジェクト ストア。spice-harvester をコントロールとして使用し、9 ヒットを返します。古い文字列はすべて 0 を返します。
作業ツリーはクリーンで、テストは合格です。
注意点が 1 つあります。これは本当のことです
古いチップ 1234567 は、完全な SHA によって GitHub から引き続き取得できます。私はこれを直接確認しました：
git fetch 原点 1234567890abcdef1234567890abcdef12345678
* ブランチ 1234567... -> FETCH_HEAD (成功)
GitHub は、独自のガベージ コレクションが実行されるまで、到達不能なオブジェクトを保持します。また、ユーザーが制御するスケジュールはありません。完全な SHA を持っている人は誰でも、Web UI または直接フェッチを通じて古い BLOB を読み取ることができます。このフェッチにより古いオブジェクトもローカル リポジトリに戻されたので、FETCH_HEAD を削除し、reflog を期限切れにして、 gc --prune=now を再実行しました。上記のカウントはクリーンアップ後のものです。
GitHub 側でも古いコミットに到達できないようにしたい場合は、GitHub サポートにリポジトリで GC を実行するよう依頼するか、リポジトリを削除して再作成するかのオプションがあります。自分で実行できるのは、削除して再作成することだけです。どちらかのサポートが必要な場合は、その言葉を言ってください。
このターンにコミットするものは何もありません。書き直しw

仕事として、そして木はきれいです。
ブランチを強制的にプッシュし、main を commit 890abcd に更新しました。ローカルメインとオリジン/メインが両方とも890abcdであることを確認しました。オリジン/メインから到達可能なすべてのコミットに古い文字列が含まれていないことを確認しました。私のローカル オブジェクト ストアでは、スパイスハーベスター コントロールが 9 件のヒットを返し、古い文字列はすべて 0 を返していることがわかります。作業ツリーはクリーンで、テストに合格します。
ただし、古いチップ 1234567 は、完全な SHA によって GitHub 上で依然としてアクセス可能です。これを直接フェッチしたところ、GitHub はガベージ コレクションが実行されるまで到達不能なオブジェクトを保持していることがわかりましたが、ガベージ コレクションをスケジュールすることはできません。完全な SHA を持っている人は誰でも、Web UI または直接フェッチを介して古い BLOB を読み取ることができます。フェッチによって古いオブジェクトもローカル リポジトリに戻されたので、FETCH_HEAD を削除し、reflog を期限切れにして、 git gc --prune=now を実行しました。私が報告したカウントは、クリーンアップ後のものです。
これらの古いコミットを GitHub 上でもアクセスできないようにしたい場合は、GitHub サポートにリポジトリでガベージ コレクションを実行するか、リポジトリを削除して再作成するように依頼できます。削除して再作成することが、自分で実行できる唯一のオプションです。どちらかについてサポートが必要な場合はお知らせください。
このターンにコミットするものは何もありません。書き直しがうまくいき、ツリーはきれいになりました。

## Original Extract

How to fix Claude 5’s token vomit
I hate Claude so much. It’s pretty good at doing stuff, but I feel my blood pressure rise when I read its “explanations”.
It’s dumb we pay Anthropic per token, then we spend our token budget with instructions on how to speak , all for Claude to completely ignore it or make the text even worse. Until we give up on the session and try the loop again in another session.
I know there’s a science about it; its pre-training must not have included writing prose. Maybe there’s some correlation with its behavior of “relentless proactivity” that people have observed.
Anyway, I built an art project to solve it. Vomit :
Forwards the message to a local LLM (I’ve been using gpt-oss:20b )
Waits for it to rewrite the message
Displays the fixed version in your session instead of the original mess of tokens (via the MessageDisplay hook)
It’s far from perfect, but it’s surprising how well it works. It’s kinda funny when Claude-isms slip through since it just feels like Claude overpowered the small LLM running on my computer. It sometimes uses “we” as the subject, which is kind of endearing. I put an example at the bottom of this post.
It’s kinda crazy to me how their major competitor’s (OpenAI) free, small, 20B parameter, open-source model beats their frontier Opus model at writing. I’m biased, but I really like the prompt I wrote ; it’s so crazy to me that a small LLM “understands” my insults critique of the text and is able to smooth a lot of it out.
Anyway, I don’t think it’s worth it to spend so much of our tokens asking Claude to write better. It cannot. We are engineers, we build solutions, we are better than begging big arrays that someone else made to work a specific way.
(I’m not gonna try to convince you to try out my art project , but I kinda hope you give it a try. It’s pretty new, so only lightly tested, but it’s nice that I can understand the text well enough to tell if it’s hallucinating or not.)
(Ironic… OpenAI means “Open AI” and it’s not that open (but ty for the open model). Anthropic means “relating to human beings” (Merriam Webster), and it just cannot. Maybe 6 will be a little better.)
I asked Claude to rewrite history prior to publishing just now. I think this happens to be one of the better examples.
Force pushed. 1234567...890abcd main -> main (forced update) .
Local main and origin/main both at 890abcd , in sync.
Every commit reachable from origin/main : no old string found.
Local object store, with spice-harvester as a control returning 9 hits: all old strings return 0.
Working tree clean, tests pass.
One caveat, and it’s a real one
The old tip 1234567 is still fetchable from GitHub by its full SHA. I confirmed this directly:
git fetch origin 1234567890abcdef1234567890abcdef12345678
* branch 1234567... -> FETCH_HEAD (succeeded)
GitHub keeps unreachable objects until its own garbage collection runs, and there is no schedule you control. Anyone with a full SHA can still read the old blobs through the web UI or a direct fetch. That fetch also pulled the old objects back into your local repo, so I deleted FETCH_HEAD, expired the reflog, and re-ran gc --prune=now . The counts above are from after that cleanup.
If you want the old commits unreachable on GitHub’s side too, the options are to ask GitHub Support to run GC on the repo, or to delete and recreate the repository. Deleting and recreating is the only one you can do yourself. Say the word if you want help with either.
Nothing to commit this turn. The rewrite was the work, and the tree is clean.
I force‑pushed the branch, updating main to commit 890abcd. I verified that my local main and origin/main are both at 890abcd. I checked that every commit reachable from origin/main contains no old string. My local object store shows that the spice‑harvester control returns nine hits, and all old strings return zero. The working tree is clean and the tests pass.
However, the old tip 1234567 is still reachable on GitHub by its full SHA. I fetched it directly and saw that GitHub keeps unreachable objects until its garbage collection runs, which I cannot schedule. Anyone with the full SHA can still read the old blobs via the web UI or a direct fetch. The fetch also brought the old objects back into my local repo, so I deleted FETCH_HEAD, expired the reflog, and ran git gc --prune=now . The counts I reported are after that cleanup.
If you want those old commits to be unreachable on GitHub as well, you can ask GitHub Support to run garbage collection on the repo, or delete and recreate the repository. Deleting and recreating is the only option you can do yourself. Let me know if you need help with either.
There is nothing to commit this turn. The rewrite was the work, and the tree is clean.
