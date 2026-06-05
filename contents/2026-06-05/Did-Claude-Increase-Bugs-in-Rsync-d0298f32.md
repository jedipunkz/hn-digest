---
source: "https://alexispurslane.github.io/rsync-analysis/"
hn_url: "https://news.ycombinator.com/item?id=48411635"
title: "Did Claude Increase Bugs in Rsync?"
article_title: "Did Claude Increase Bugs in rsync?"
author: "logicprog"
captured_at: "2026-06-05T13:08:48Z"
capture_tool: "hn-digest"
hn_id: 48411635
score: 49
comments: 25
posted_at: "2026-06-05T12:43:33Z"
tags:
  - hacker-news
  - translated
---

# Did Claude Increase Bugs in Rsync?

- HN: [48411635](https://news.ycombinator.com/item?id=48411635)
- Source: [alexispurslane.github.io](https://alexispurslane.github.io/rsync-analysis/)
- Score: 49
- Comments: 25
- Posted: 2026-06-05T12:43:33Z

## Translation

タイトル: クロードは Rsync のバグを増やしましたか?
記事のタイトル: クロードは rsync のバグを増やしましたか?

記事本文:
データ分析 · 2026 年 6 月
クロードは rsync のバグを増やしましたか?
バグ データを含むすべての rsync リリースの単純な分布分析。モデルはありません。仮定はありません。ただの配置。
1 · 背景: 「rsync の怒り」
2026 年 5 月下旬、rsync が爆発しました。 GitHub、ハッカー ニュース、ロブスターズ: オープンソースのメンテナーが AI で書かれたコードを配布して信頼できるものにすることができるかどうか、またコードを無料で受け取る人々がコードの作り方を要求できるかどうかについて、何百人もの人々が議論しています。
2026 年 5 月 30 日、「このソフトウェアをバイブで動かさないでください」というタイトルの GitHub の問題が rsync リポジトリに対してオープンされました。プロジェクトによるクロードの使用を批判するマストドンの投稿のスクリーンショットが添付されていた。バグレポートはありません。技術的な内容はありません。その後の事態は異常なもので、思慮深い懸念からあからさまな嫌がらせに至るまで、329 件のコメントが寄せられ、その数は増え続けています。
スレッドは言葉で止まりませんでした。あるユーザーは、「バイブコード化されたコミットをプッシュしたプロジェクト管理人」の首を絞めているマイリトルポニーの絵を投稿しました。
それは Hacker News や Lobsters にも広がり、さらに数百ものコメントが生成されました。どこでも繰り返される中心的な主張: クロード支援開発により、以前は安定していたツールにバグが導入されました。
非常に安定していて信頼されているツールがすぐに下り坂になり始めたことに対して人々が怒っているのは当然です…すべてはメインの開発者がそのソフトウェアをバイブコーディングしているためです。
ロブスターについて、ユーザー boramalper は次のように書きました。
誰かが（可能であれば）各リリース後の回帰のタイムチャートを実際に作成して、最近数値が実際に増加したかどうかを確認できたら興味深いでしょう。
ユーザーの bitshift は、「私もそのようなチャートを見たいと思っています。完全に有益というわけではありません…しかし、少なくとも、客観的に測定できるものにはなるでしょう。」と答えました。
今回の分析はそのチャートです。リリースごとに 1 つの指標があり、モデルはありません。
オン t

HN スレッド、ユーザー zos_kia が混乱を直接指摘しました。
ざっと見たところ、これは 2007 年以来コード内に存在していたコーディング エラーを CVE が表面化させたことに対応したセキュリティ修正のように見えます。これはあまりにもありふれたものなので、人々がそれを気にしなくなっているのを見るのは実際には滑稽です。
ロブスターズについて、ユーザー jbert は因果関係の連鎖を次のように説明しました。
変更量の増加 (したがってリグレッションの数の増加) のきっかけは、(主に) LLM 対応のセキュリティ問題の流入でした。つまり、因果関係の連鎖は次のとおりです。LLM → より多くの既知のセキュリティ問題 → 通常よりも多くの変更が必要 → 通常よりも多くの回帰。
これらのユーザーは、まさに混乱を特定しました。それは、回帰を引き起こしたコードを作成した AI ではなかったということです。 Tridge が通常よりも多くの変更をリリースすることを余儀なくされたのは、AI がセキュリティ ホールを発見したためでした。そして、誰が作成したかに関係なく、より多くの変更はより多くのリグレッションを意味します。これはクロードの問題ではありません。それは「さらなる変化」の問題です。 Tridge 氏自身も、返答の中でこの因果関係を確認し、AI によって生成された CVE レポートの洪水がどのようにして rsync の攻撃対象領域に急速かつ広範な変更を強いたかを説明しました。彼は退職した開発者で、航海することを望んでいたため、テスト スイートの作成、多層防御の強化、セキュリティ バックログの処理など、大量の作業を手伝ってもらうようクロードに連絡を取りました。同氏はv3.4.3のリグレッションを認めたが、エッジケースの互換性よりもセキュリティ修正を意図的に優先したと述べた。
バグ データを含む 37 のリリース (v2.4.6 から v3.4.3 まで)
2 つのリリースにはクロード コミットがあります: v3.4.2 (9 クロード、0.80 バグ/10c) および v3.4.3 (28 クロード、6.76 バグ/10c)
どちらも過去の分布の中央の 50% 内に収まります
任意の 2 つのリリースのランダムなペアの 46% が、Claude リリースと同等かそれよりも悪いスコアを獲得
過去の平均はクロード平均の 2 倍です (7.60 対 3.78 バグ/10c)

)
レジームシフトは検出されませんでした: テストを実行 p=0.231、シーケンスはランダム性と一致しています
v3.4.1 (102 バグ / 9 コミット、クロードなし) は外れ値ですが、ベースラインに属します。これはリリースであり、ディストリビューションはすでにそれをキャプチャしています
分析では、10 コミットあたりのバグ数 (バグ/10c) という単一の指標が使用されます。各リリースについて、そのリリースに起因するバグの数をその範囲内のコミット数で割り、10 を掛けます。これはリリース サイズに対して正規化されます。
コミットがリリースに割り当てられる方法
デフォルト ブランチ上のすべてのコミットはコミッターの日付順に並べられ、順次タイムラインが生成されます。各 git タグは、このタイムライン内の特定のコミットを指します。リリースの範囲は、前のタグとそのタグの間のすべてのコミットです。プレリリース タグ (「pre」、「rc」) は境界としてスキップされ、最終リリースに吸収されます。すべてのコミットは 1 つのリリースにのみ属します。
バグはどのように発見され、その原因が特定されるのか
バグ数は、rsync リポジトリ内の GitHub の問題、rsync Bugzilla インスタンス、rsync メーリング リストの 3 つのソースから取得されます。 rsync プロジェクトに対して提出された問題は、GitHub REST API 経由で収集されました。メーリング リストのバグは、バグ レポート パターンのメッセージ件名を解析し、プロジェクトの問題追跡と相互参照することによって特定されました。 Bugzilla エントリは Bugzilla API を介して収集されました。各エントリには、バグが報告されたリリースを明示的に示す「バージョン」フィールドがあり、バグはそのリリースに起因します。 GitHub の問題とメーリング リストのバグは、バグが報告される前に出荷された最新のリリースに起因すると考えられます。
批評家の主張は単純な比較であり、金利は上がったかどうかというものである。最も単純な正直な応答は、単純なレートです。クロードのリリースが歴史的分布の真ん中に位置する場合、それを説明する責任は批評家に移る

なぜこの特定のミドルが、それ以前の他のすべてのミドルよりも劣っているのかということです。
コミットの複雑さ、セキュリティの強度、バグの重大度は制御されません。 1 行のタイプミス修正と CVE パッチは区別されません。鈍器です。しかし、批評家らの非難も率直で、「クロードが状況を悪化させている」というものだ。鈍器は最も公平な反応です。
クロードのリリースはどれくらい正常ですか?
これらのリリースは、以前のすべてのリリースのディストリビューションに含まれる位置は次のとおりです。
各ドットはリリースです。緑の影付きのバンドは四分位範囲 (IQR) です。
過去のリリースの中間 50%、10c あたり 0.65 ～ 6.82 バグ。
過去のすべてのリリースの半分はこのバンド内にあり、残りの半分はその外にあります。
両側の暗い領域は下四分の一と上四分の一です。
クロード リリース (緑色の点) は両方とも IQR 内に収まります —
バグ率は典型的な歴史的範囲と区別がつきません。
過去の平均は 7.60 バグ/10c ですが、これは二峰性分布によって決まります。 v2.x リリースでは平均 2.04 バグ/10c。 v3.x リリースの平均は 11.46。 v3.x の中でも、Claude リリースは目立ったものではありません。v3.4.2 は 21 の v3.x リリース中 16 位にランクされ、v3.4.3 も 16 位にランクされています。
35 個の非クロード リリースに対する実行テストでは、14 回の実行が見つかりました (ランダム性、z=-1.54、p=0.123 の下で 18.5 と予想)。時間的クラスタリングの証拠はありません。シーケンスは、同じ分布からのランダムな抽出と一致しています。
5 · データの一致と不一致
…「私は xyz 大学の博士号で、LLM はすべてを構成する単なる確率論的なツールであり、それを使用すると世界が崩壊すると言っています」のようなことを言っている人たちに対して、私はここであなたが時代遅れであることを伝えたいと思います。ソフトウェア エンジニアリングの世界はここ数か月で劇的に変化しました

これ。大量の報告に直面して、IT セキュリティとソフトウェアの保守の世界は、ここ数週間で完全かつ完全に変化しました。あなたが昨年このことについて学んだことは、他の惑星から来たものと同じかもしれません… 結論から言えば、私は LLM がどのように機能するかを (まあ、大まかに) 知っていますが、だからと言って LLM が役に立たないわけではありません。それは確かに注意しなければならないことを意味しますが、私は慎重であるか、いわゆるインターネット専門家からの大量の嫌がらせに対処せずに航海したいという私の願望を考慮して、可能な限り慎重です。

## Original Extract

Data Analysis · June 2026
Did Claude Increase Bugs in rsync?
A simple distributional analysis of every rsync release with bug data. No model. No assumptions. Just placement.
1 · Background: The "rsync Outrage"
In late May 2026, rsync blew up. GitHub, Hacker News, Lobsters: hundreds of people arguing about whether open-source maintainers can ship AI-written code and have it be reliable — and whether the people taking the code for free get to demand how it is made.
On May 30, 2026, a GitHub issue titled "Please Do Not Vibe Fuck Up This Software" was opened against the rsync repository. It attached a screenshot of a Mastodon post criticizing the project's use of Claude. No bug report. No technical content. What followed was extraordinary: 329 comments and counting , ranging from thoughtful concern to outright harassment.
The thread did not stop at words. One user posted My Little Pony drawings of themselves strangling the "project janitor that pushed vibecoded commits":
It spread to Hacker News and Lobsters , generating hundreds more comments. The central claim, repeated everywhere: Claude-assisted development introduced bugs into a previously stable tool.
People are very justifiably angry that a very stable, well trusted tool , has started to immediately go downhill… all because the main dev is vibecoding that software.
On Lobsters, user boramalper wrote:
It'd be interesting if someone actually did a timechart of regressions after each release (if at all possible) to see if the number actually went up recently or not.
User bitshift replied: "I would also love to see such a chart. It wouldn't be completely informative… But at least it would be something objective we could measure."
This analysis is that chart. One metric, every release, no model.
On the HN thread, user zos_kia pointed at the confound directly:
From a cursory look, it looks like a security fix in response to a CVE surfaced a coding error which has been present in the code since 2007. This is so banal that it's actually hilarious to see people lose their shit over it.
On Lobsters, user jbert spelled out the causal chain:
The trigger for the increased volume of changes (and hence increased number of regressions) was the influx of (mostly) LLM-enabled security issues. i.e. the causal chain was: LLMs → more known security issues → more changes needed than usual → more regressions than usual.
These users identified the exact confound: it wasn't AI writing the code that caused regressions. It was AI finding security holes that forced tridge to ship more changes than usual — and more changes means more regressions, regardless of who wrote them. This is not a Claude problem. It is a "more changes" problem. Tridge himself confirmed this causal chain in his response, describing how a flood of AI-generated CVE reports forced rapid, extensive changes to rsync's attack surface. A retired developer who would rather be sailing, he reached for Claude to help with the volume: writing test suites, adding defence-in-depth hardening, and working through the security backlog. He acknowledged the regressions in v3.4.3 but said he had deliberately prioritized security fixes over edge-case compatibility.
37 releases with bug data, spanning v2.4.6 to v3.4.3
2 releases have Claude commits: v3.4.2 (9 Claude, 0.80 bugs/10c) and v3.4.3 (28 Claude, 6.76 bugs/10c)
Both fall inside the middle 50% of the historical distribution
46% of random pairs of any 2 releases score equal or worse than the Claude releases
The historical mean is 2× the Claude mean (7.60 vs 3.78 bugs/10c)
No regime shift detected: runs test p=0.231, sequence is consistent with randomness
v3.4.1 (102 bugs / 9 commits, no Claude) is an outlier but belongs in the baseline — it is a release, and the distribution already captures it
The analysis uses a single metric: bugs per 10 commits (bugs/10c). For each release, divide the number of bugs attributed to it by the number of commits in its range, then multiply by 10. This normalizes for release size.
How commits are assigned to releases
Every commit on the default branch was ordered by committer date to produce a sequential timeline. Each git tag points to a specific commit in this timeline. A release's range is all commits between the previous tag and its own tag. Pre-release tags ("pre", "rc") are skipped as boundaries and absorbed into their final release. Every commit belongs to exactly one release.
How bugs are found and attributed
Bug counts come from three sources: GitHub issues in the rsync repository, the rsync Bugzilla instance, and the rsync mailing list. Issues filed against the rsync project were collected via the GitHub REST API. Bugs from the mailing list were identified by parsing message subjects for bug report patterns and cross-referencing with the project's issue tracking. Bugzilla entries were collected via the Bugzilla API; each entry has a "Version" field that explicitly states which release the bug was reported against, and bugs are attributed to that release. GitHub issues and mailing-list bugs are attributed to the most recent release that shipped before the bug was reported.
The critics' claim is a simple comparison: did the rate go up? The simplest honest response is a simple rate. If the Claude releases sit in the middle of the historical distribution, the burden shifts to the critics to explain why this particular middle is somehow worse than all the other middles that came before it.
It does not control for commit complexity, security intensity, or bug severity. It does not distinguish between a one-line typo fix and a CVE patch. It is a blunt instrument. But the critics' accusation is also blunt: "Claude is making things worse." A blunt instrument is the fairest response.
How Normal Are the Claude Releases?
Here is where these releases fall in the distribution of all prior releases:
Each dot is a release. The shaded green band is the interquartile range (IQR) —
the middle 50% of historical releases, from 0.65 to 6.82 bugs/10c.
Half of all historical releases fall inside this band, and half fall outside.
The darker regions on either side are the lower and upper quarters.
The Claude releases (green dots) both fall inside the IQR —
their bug rates are indistinguishable from the typical historical range.
The historical mean is 7.60 bugs/10c, but this is driven by a bimodal distribution. v2.x releases average 2.04 bugs/10c; v3.x releases average 11.46. Even within v3.x, the Claude releases are unremarkable: v3.4.2 ranks 16th of 21 v3.x releases, v3.4.3 ranks 16th as well.
A runs test on the 35 non-Claude releases finds 14 runs (expected 18.5 under randomness, z=-1.54, p=0.123 ). There is no evidence of temporal clustering — the sequence is consistent with a random draw from the same distribution.
5 · What the Data Is Consistent And Inconsistent With
…for the people saying things like "I'm a PhD from xyz uni and I'm telling you LLMs are just stochastic tools that make everything up and the world will fall apart if you use them", I'm here to tell you that you are out of date. The world of software engineering has changed dramatically in the last few months. The world of IT security and maintaining software in the face of the flood of reports has completely and utterly changed just in the last few weeks. Anything you learned about this stuff last year might as well be from another planet… Bottom line is I do know (well, roughly!) how LLMs work, but that doesn't make them not useful. It does mean you have to be cautious, but I am being cautious, or as cautious as I can be given my desire to be sailing and not dealing with a flood of gunk from so-called internet experts.
