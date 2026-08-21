---
source: "https://allaboutcoding.ghinda.com/a-week-of-using-codex-more-than-claude/"
hn_url: "https://news.ycombinator.com/item?id=49393051"
title: "Quick impressions: A week of using Codex more than Claude"
article_title: "Quick impressions: A week of using Codex more than Claude | All about coding"
image: "https://allaboutcoding.ghinda.com/assets/images/posts/a-week-of-using-codex-more-than-claude/og.png"
author: "speckx"
captured_at: "2026-08-21T20:14:36Z"
capture_tool: "hn-digest"
hn_id: 49393051
score: 2
comments: 0
posted_at: "2026-08-21T19:51:48Z"
tags:
  - hacker-news
  - translated
---

# Quick impressions: A week of using Codex more than Claude

- HN: [49393051](https://news.ycombinator.com/item?id=49393051)
- Source: [allaboutcoding.ghinda.com](https://allaboutcoding.ghinda.com/a-week-of-using-codex-more-than-claude/)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T19:51:48Z

## Translation

タイトル: 簡単な感想: クロードよりも Codex を使用した 1 週間
記事のタイトル: 簡単な感想: クロードよりも Codex を使用した 1 週間 |コーディングに関するすべて
説明: クロードは、質問されたことをはるかに超えて、あなたが望むものを推測します。 Codex は、ユーザーの指示どおりに実行し、実行される可能性がある最初の兆候で停止します。 Codex をさらに 1 週​​間使用してからの 10 件の感想。

記事本文:
Lucian Ghinda が執筆した Ruby と Rails の技術コンテンツ
簡単な感想: クロードよりも Codex を使用した 1 週間
今週、Claude よりも Codex を使用したことによる非常に個人的な感想をいくつか簡単に述べます (できれば週末に完全な分析を行うつもりです)。
(1) 今年は、同じプラグイン/スキルのセットなどを持ち、クロードとコーデックスを同等に保つように努めましたが、いくつかのセッションからスキルを作成し、そのすべてがコーデックスに移植されなかったため、クロードの方がより多くのスキルを持っていました。これを修正するのは簡単です。Codex にクロード スキル フォルダーを指定し、クロード用に変換するよう依頼します。
(2) 急いでいるとき (緊急だと感じたものをデバッグするときなど)、それでもクロードを開いたのは、どういうわけかクロードのほうが居心地が良かったからです。それがより優れていたとは言いませんが、慣れ親しんだものでした。デバッグの際には、私が知っているツールを使用することが重要でした。
(3) Codex によって作成された変更では、Ruby/Ruby on Rails コードのコメントが減りました。私はそれがとても気に入ったので、これに関して私が行ったいくつかの実験を近々共有したいと思います。
(4) Codex エージェント ハーネスの出力は、Claude からの出力よりもはるかに「技術的」です。クロードは、あなたにメッセージを送っているタプル セッションの同僚のように感じますが、コーデックスはスタートレックのデータのバージョンのように感じます。
(5) 以前のようにクロードの大きなセッションを開く代わりに、Codex のセッションをもっと多く開き、集中力を維持したいと考えています。これは Codex に固有のことではないかもしれませんが、Codex を使用しているときに気づきました。
(6) Codex の方がクロードよりも早く変更を行うように感じます。しかし、主な変更を行った後、プル リクエストを完了するには、多くのテストの再実行やレビューなど、多くの時間がかかりました。この徹底ぶりは好感が持てるが、結局タイム差で勝負はならなかった。
(7) Codex は、Claude よりもコード アーキテクチャの点ではるかに単純なソリューションを作成したように感じました。

。クロードは通常、抽象化、概念、ソルベ シグネチャ、型エイリアスなど、多くのものを作成し続けます。 Codex はもう少し封じ込められ、作成される量は減りました。今週は、コード調査 -> 設計変更 -> 変更レビュー -> 実装 -> 検証という改善されたフローもテストしました。しかし、私は両方に同じドキュメントを使用して同じ要件を実装させました。クロードのコードはもう少し複雑ですが、ケースは処理されました。
(8) コーデックスにもいくつかの間違いがありました。クロードは、他の仕事から手を広げ、それらを同期させたいという私の意図を理解してくれました。 Codex は、ブランチ A がメインをターゲットとするブランチ B をターゲットにするなど、厄介なことを行いました。リベースを依頼すると、メインでリベースされ、4000 以上の追加を含む PR が作成されました。明示的に、ターゲットのみを使用してリベースするように依頼する必要がありました。
(9) Codex の場合、MCP ではなく CLI ツールを使用する私の環境では、Jira と Atlassian を使用するのが面倒でした。 JIRA を開いてログインを求め、CLI に切り替えてからブラウザに戻りました。この場合、クロードは、以前のセッションに基づいて、私が望むものを手に入れ、私が望むようにそれを実行しようとすることにはるかに熱心でした。
(10) MCP を使用する場合、私は Codex CLI アプローチの方が好きです。Codex mcp ログインを実行するように求められ、毎回適切な認証と認可のフローが開かれます。クロードは時々ターン中に自動的に実行しようとしますが、スタックすることがあります。
私が感じるクロードとコーデックスの主な違いは、クロードは求められたことを超えて、あなたが何を望んでいるのかを推測し、それを直接実行しようとするのに対し、コーデックスは言われたことはやるけど無理はしない仲間のようなところだと思います。それが完了する可能性があるという最初の兆候で停止します。
私はこの記事を書き、Grammarly を使用して校正と修正を行いました。
ステータス行の設定

Codex 内のすべてのプロジェクト
2026 年 8 月 20 日
Ruby を使用してクロード コードにステータス ラインを設定できます
2026 年 8 月 19 日
© 2026 ルシアン・ギンダ ·
RSS・
ギンダ.com
Jekyll で構築 · テーマに基づいて
ジキル・ブラッグ
Mike Vormwald 著 (オリジナルは Cassidoo )。

## Original Extract

Claude goes above and beyond what is asked and guesses what you might want. Codex does what you tell it and stops at the first sign that it might be done. Ten impressions from a week of using Codex more.

Ruby and Rails technical content written by Lucian Ghinda
Quick impressions: A week of using Codex more than Claude
Some quick and very personal impressions from using Codex more than Claude this week (I will do a full analysis during the weekend hopefully).
(1) While I tried this year to keep Claude and Codex on par, having the same set of plugins/skills and so on, Claude had more skills, as I created skills out of some sessions and not all of them were ported to Codex. Fix for this is simple: Point Codex at the Claude skills folder and ask it to transform them for Claude
(2) When I was in a rush (like debugging something that felt urgent), I still opened Claude as somehow I felt more at home with it. I am not saying it was better, but it was familiar, and when debugging, using tools that I know is important.
(3) Changes created by Codex had fewer comments in Ruby/Ruby on Rails code. I liked that a lot, and I will soon share some experiments I ran on this.
(4) The output of the Codex agent harness is much more “technical” than the one from Claude. Claude feels more like your colleague in a Tuple session writing to you while Codex feels more like a version of Data from Star Trek.
(5) I want to open many more sessions of Codex and keep them focused instead of a big session of Claude as I was doing before. This may not be specific to Codex, but I noticed it while working with Codex.
(6) It feels to me that Codex does changes faster than Claude. But after making the main changes, it took a lot to finish the pull request: rerunning many tests, review, and so on. I like the thoroughness of this, but in the end, there was no win in terms of time difference.
(7) It felt to me that Codex created a much simpler solution in terms of code architecture than Claude. Claude usually goes on to create a lot of things: abstractions, concepts, Sorbet signatures, type aliases, and so on. Codex was a bit more contained and created less. This week I also tested an improved flow of code research -> design change -> review change -> implement -> verify . But I made both of them implement the same requirement using the same documents, and Claude’s code was a bit more complex but handled cases.
(8) Codex also made some mistakes. Claude could understand my intention to branch out from other work and keep them in sync. Codex did some nasty things like branch A targets branch B that targets main, and when I asked it to rebase, it rebased with main, which created some PR with 4000+ additions. I had to be explicit and ask it to rebase only with the target.
(9) For Codex, working with Jira and Atlassian was a hassle in my environment where I use the CLI tool and not the MCP. It opened JIRA to prompt me to log in, then switched to the CLI, then back to the browser. In this case, Claude was much more eager to try to get what I want and to do it the way I want it done, based on previous sessions.
(10) Working with MCPs, I like the Codex CLI approach more, where it asks me to execute codex mcp login and every time it opens the right authentication and authorization flow. Claude sometimes tries to run it automatically in a turn, and it can get stuck.
I think the main difference I feel between Claude and Codex is that Claude tries to go above and beyond what is asked and guess what you might want and then directly do it, while Codex is more like a companion that does what you tell it but will not overdo it. It will stop at the first sign that it might be done.
I wrote this article and used Grammarly to proofread and fix it.
Setting the status line for all projects in Codex
Aug 20th, 2026
You can use Ruby to set the status line in Claude Code
Aug 19th, 2026
© 2026 Lucian Ghinda ·
RSS ·
ghinda.com
Built with Jekyll · theme based on
jekyll-blahg
by Mike Vormwald (original by Cassidoo ).
