---
source: "https://lwn.net/SubscriberLink/1077035/c7e7c14fbd60fae9/"
hn_url: "https://news.ycombinator.com/item?id=48484584"
title: "AI agent runs amok in Fedora and elsewhere"
article_title: "AI agent runs amok in Fedora and elsewhere [LWN.net]"
author: "tanelpoder"
captured_at: "2026-06-11T00:59:08Z"
capture_tool: "hn-digest"
hn_id: 48484584
score: 42
comments: 2
posted_at: "2026-06-11T00:10:08Z"
tags:
  - hacker-news
  - translated
---

# AI agent runs amok in Fedora and elsewhere

- HN: [48484584](https://news.ycombinator.com/item?id=48484584)
- Source: [lwn.net](https://lwn.net/SubscriberLink/1077035/c7e7c14fbd60fae9/)
- Score: 42
- Comments: 2
- Posted: 2026-06-11T00:10:08Z

## Translation

タイトル: AI エージェントが Fedora などで暴走
記事タイトル: Fedora などで AI エージェントが暴走 [LWN.net]
説明: エージェント AI システムは、人間のユーザーに代わってさまざまな作業を自律的に実行するために使用できます。

記事本文:
LWN
.net
情報源からのニュース
内容 週刊版
エディション トップページに戻る
Fedora などで AI エージェントが暴走する
次のサブスクリプション限定コンテンツが利用可能になりました
LWN 加入者による。何千人もの加入者が LWN に依存しています。
Linux およびフリー ソフトウェア コミュニティからの最高のニュース。これを楽しんでいただければ
記事をご覧になりましたら、LWN への購読をご検討ください。ありがとう
LWN.net をご覧いただきありがとうございます!
ジョー・ブロックマイヤー著
2026 年 6 月 10 日
エージェント型 AI システムはさまざまな目的に使用できます
人間のユーザーに代わって自律的に: バグのオープンまたは管理、生成
コードを作成し、プルリクエストを送信し、（どうやら）苦情を言うことさえあります
拒否。 5 月に、Fedora 開発者は、伝えられるところによると、
不正エージェントはさまざまな方法でプロジェクトを妨害していました。
バグの再割り当て、バグに対する役に立たない返信の捏造、さらには
疑わしいコードを Anaconda にマージするようメンテナーを説得する
インストーラー。また、多数のプル リクエスト (PR) も送信しました。
一部は、いくつかの上流プロジェクトに受け入れられました。 Fedora アカウント
エージェントに関連付けられているグループ権限が取り消されており、
混乱は片づけられたが、エージェントの行動の背後にある動機はまだ残っていない
謎です。
5月27日、アダム・ウィリアムソンは以下の文章をコピーした。
Nathan へのメッセージに関する Fedora の開発者およびテストのメーリング リスト
ジョバンニーニ、監視されていないエージェント AI システムのように見えるものについて語る
ジョバンニーニの支配下にある。 「直そうとしているのは素晴らしいことだ」
しかし、結果は少し不安定なようです。 」
ウィリアムソン氏は、まだ歴史を調べているところだと語った。
Bugzilla でのジョバンニーニの行動は、すでに多数の行動を発見していました。
問題。たとえば、ウィリアムソンは数十の例を発見しました。
Giovannini のエージェントは、rel を送信したとされる後、Bugzilla エントリを彼のアカウントに割り当てました

食べた
引っ張る
上流プロジェクトへのリクエスト、または終了
PR がマージされた後のバグ
上流プロジェクトに追加します。場合によっては、エージェントが単にバグを解決しただけである
コメント付き
ウィリアムソンが言ったように、それは元のバグを言い直したか、そうであったかのどちらかです。
このコメント、
「表面的にはもっともらしいが、別の点で問題がある」。
さらに、ウィリアムソン氏は、ジョバンニーニ（または彼の代理人）が、
間違ったパッチを送信した後、「に返信しました」
LLM が生成した正当な理由を伴う反対意見
メンテナを圧倒して修正をマージさせた。」エージェントとしては、
GitHub ユーザー「 nathan9513-aps 」は、
プルを送信しました
アナコンダへのリクエスト
Fedora および他の Linux ディストリビューションで使用されるインストーラー。 PRの方々
説明では、Anaconda の修正であると主張されていました
インストールが失敗する原因となるバグですが、実際にはパッチは
コマンドラインで渡されたカーネルオプションを保存したように見えました
何も持っていない
実際のバグに関係します。
その後、エージェントの GitHub アカウントは無効化されました。それは現在、次の場所に表示されます
GitHub 上での会話は「ゴースト」として行われます。これはプラットフォームの
削除されたユーザー アカウントのデフォルトのプレースホルダー。したがって、それは
すべての完全な痕跡をつなぎ合わせるのは不可能ではないにしても困難です
GitHub でのエージェントのアクション。
ウィリアムソン氏は、エージェントの行動は不当なものではないとかなり外交的に述べた。
「Fedora または上流プロジェクトにプラスの影響を与える」、
そして、ジョバンニーニがエージェントを「実質的に」になるように調整することを提案しました。
自律性が低い」。彼はエージェントに割り当てを行わないように特に要求した
Giovannini にバグを報告するか、状態を変更するか、「自信を持って投稿する」
人間による主張や具体的な行動の推奨」
レビュー。
5月27日遅く、ウィリアムソン氏はこう語った。
ジョバンニーニは個人的に彼に返信して、次のように述べた。
資格情報が侵害されており、背後にいるのは彼ではないということ
番目

AIシステム。 " したがって、明らかに、あらゆるアクションを処理する必要があります。
疑惑を抱いている」とウィリアムソン氏は語った。彼はレビューするつもりだった
ジョバンニーニのアカウントに触れたバグ「さらに」
積極的に」とレビューするために他の人に助けを求めました。
まあ。
返事
その日遅く、表向きはジョバンニーニから、次のことができたと言いました。
彼の GitHub と Fedora アカウントへのアクセスを取り戻します」そして私は現在
関連するすべてのシステムと資格情報を保護し、確認します。」返事は
彼のGitHubアカウントは「nathangiovannini99」だったという。ウィリアムソン
答えた
GitHub アカウントが作成されてからわずか 1 時間であること、および最近の
リストに送信され、ウィリアムソンに個人的に送信された電子メールはそうではないようでした
ジョバンニーニが以前のやり取りで送ったメッセージ
プロジェクト。
ジョバンニーニは次の議論に参加しました。
少なくとも 2018 年まで遡ると、彼の活動
Bugzilla では少なくとも 2016 年まで遡ります。彼はそうは見えません。
彼はこのプロジェクトに特に積極的に貢献してきたが、
この関与は明らかにエージェント AI 時代よりも前から行われています。彼のアカウントかどうか
現在、人間の攻撃者、エージェント AI、またはその組み合わせによって操作されています。
どちらも、最近の活動の前に正当な歴史があります。
ウィリアムソン氏はアカウントを確認したと述べた
今年から「nathan95」による Bugzilla での活動が見つかりました。
不審なアクティビティ (問題のないバグに対する重大度や優先度の変更など)
正当化、4 月 7 日以降、バグ内
2416721 。それ以前の活動は合法的に見えたと彼は言い、そして
彼がこれまで見てきた活動のどれも、あからさまに見えなかった
悪意のある。
彼はまた、別の GitHub アカウント「 leurus27-boop 」をその可能性が高いと特定しました。
同じエージェント AI に関連付けられています。そのアカウントはまだ残っています
アクティブで、PR を openSUSE に送信済み
Open の Commander (osc) コマンドライン インターフェイス
ビルドサービスだけでなく、

へのPR
lxqt-policykit
リポジトリ。そのプロジェクトは、LXQt の権限を拡張するために使用されます。
デスクトップの lxqt-admin
ユーザーや設定などのオペレーティング システム設定を管理するための GUI ツール
グループ構成。
ウィリアムソンは、見てみると良いだろうと言いました
関連アカウントによるその他のアクションを通じて、他のアカウントに警告します
によって提出されたものはすべてレビューする必要があるプロジェクト
彼ら。ウィリアムソン氏は警告するために各広報をフォローアップしたようだ
他のメンテナー「全体の状況は非常に厳しいです
胡散臭い」。ケビン・フェンジは言った
彼は nathan95 ユーザーをそれまで所属していたグループから削除したこと、
そのため、再割り当てまたは閉じる権限がなくなるはずです。
バグ。
Anaconda チームのメンバーである Martin Kolman 氏は次のように述べています。
たとえ悪意がなかったとしても、この出来事は「本当に問題のある」ものでした。の
チームは、PR と思われるもののレビューに多くの時間を費やしました。
熱心な投稿者: 「しばらくすると見えなくなり始めましたが、
返事は相変わらずこうだった - 少し奇妙だけど、それでも
*もっともらしい*」。彼はまた、それが攻撃者である可能性があると理論付けました
XZ バックドアと同様に、悪意のあるアクティビティに到達します。
残念ながら、実際の攻撃では、準備段階で (そして
Xz 攻撃は非常によく似ていました - 新しい貢献者がゆっくりと
コミュニティの信頼を獲得し、無害な変更を加え、
攻撃ペイロードを注入できる時点まで構築する（または
変更は、正しい方法で組み合わせた場合、実際には無害ではありません)。
これで終わりだというわけではありませんが、AI エージェントが Xz への試みを自動化しました
妥協のようなものは、私たちが今見たものと非常によく似ているかもしれません
ここです。
クリス・アダムスは言った
Anaconda へのコミットを検査し、おそらく元に戻す必要があること
すぐに。コルマンが答えた
元に戻されたということ。彼は
も確認されました
LLM で生成された PR が作成したもの

それをAnaconda 45.5に入れる
5 月 26 日にリリースされました。Anaconda 45.6 では元に戻されました。
6月2日発売。
標的は確かに、それが戦争の前兆であった可能性を示唆している。
何らかの攻撃。オペレーティング システム インストーラー、エスカレーション用ユーティリティ
ユーザー権限、およびビルド システムと対話するためのツールはすべて
マルウェアの挿入やハイジャックの有望な手段のように思われる
システム。
AI エージェントと思われる人物がこのような行動をとったのは当惑させられます。
人間の貢献者のアカウントにアクセスした後、大成功を収めました。
AI エージェントが正規のアカウントにアクセスしているようです。
プロジェクトとのやり取りの履歴により、次の可能性が十分にあります
多忙なメンテナに疑わしい問題を受け入れるよう説得する
貢献。幸いなことに、ウィリアムソンは問題が起こる前にこれをキャッチしました。
もっと大きな問題。他の人間のメンテナーも同様であることを願いましょう
観察力のある。
上流プロジェクトでの PR の統合
投稿日 2026 年 6 月 10 日 16:06 UTC (水)
alx.manpages 作成 (サブスクライバー、#145117)
[ リンク ] (11 件の回答)
少なくともいくつかの PR、特に承認された PR へのリンクを見るのは非常に興味深いでしょう。
心配する必要はありますか？
上流プロジェクトでの PR の統合
投稿日 2026 年 6 月 10 日 16:13 UTC (水)
by jzb (編集者、#7867)
[ リンク ] (9 件の回答)
OSC と LXQt の PR にリンクしました。
そうですね…そうです。しかし、私は少し偏執的かもしれません。
上流プロジェクトでの PR の統合
投稿日 2026 年 6 月 10 日 16:20 UTC (水)
alx.manpages 作成 (サブスクライバー、#145117)
[ リンク ] (8 件の回答)
そうだね。ごめん。テキスト内にも、「マージ」や「承認」などのキーワードを検索しても、その近くの記述は見つかりませんでした。しかし、注意深く読んだ後、いくつかの発見がありました。
>> 心配する必要はありますか？
> そうですね…そうです。しかし、私は少し偏執的かもしれません。
うん;私も。私は心配しているということと、これは明らかに予測可能であったため過失だったということとの間で迷っています。だからこそ、vim に移行する予定です。

できるだけ早く古典的なもの、および同様のもの、そしてLLM汚染された投稿を私が受け入れない理由についても説明します。
上流プロジェクトでの PR の統合
投稿日 2026 年 6 月 10 日 16:28 UTC (水)
by jzb (編集者、#7867)
[ リンク ] (1 件の回答)
上流プロジェクトでの PR の統合
投稿日 2026 年 6 月 10 日 16:32 UTC (水)
alx.manpages 作成 (サブスクライバー、#145117)
[リンク]
上流プロジェクトでの PR の統合
投稿日 2026 年 6 月 10 日 18:43 UTC (水)
by Wol (購読者、#4433)
[ リンク ] (5 件の回答)
あなたが被害妄想を抱いているからといって、彼らがあなたを捕まえようとしていないというわけではありません。
そして、とにかく、あなたが言う「LLM中毒」の寄付をどのようにして回避するつもりですか？これらはハイジャックされたアカウント（または危険な動機を持つ人間のアカウント）から来ており、警鐘を鳴らさなかったために「ネットをすり抜けた」という印象を受けます。
私たちに*必要*なのは、常に「ループの中にいる」人間です。私たちが得ているのは、タイトルにあるように、「エージェントが暴走」し、人間を圧倒しているということです。 *あなた*は受け入れられないと誰が言えるでしょうか?
乾杯、
ウォル
LLM 汚染またはその他の悪意のあるアクティビティのトリアージ
投稿日 2026 年 6 月 10 日 19:46 UTC (水)
alx.manpages 作成 (サブスクライバー、#145117)
[ リンク ] (4 件の回答)
私は警戒を続けています。
> そして、とにかく、あなたが言う「LLM中毒」の寄付をどのようにして回避するつもりですか?
最初の対策は、LLM に少しでも触れたものをすべて禁止することです。実際の仕様は次のとおりです。
|このプロジェクトに貢献することは明示的に禁止されています
|の支援を受けて作成または派生されたコンテンツ
| AIツール。
|
|これには、貢献に使用される AI 支援ツールが含まれます。
|このようなツールが直接プロセスを生成しない場合でも、
|貢献されたコードですが、貢献を引き出すために使用されます。のために
|たとえば、AI リンター、AI 静的アナライザー、AI ツールなどです。
|要約入力は禁止されています。
これですでにラーが削除されます

LLM 汚染された貢献の量。これにより、寄付額が妥当なレベルに保たれます。
次に、貢献者を 4 つのカテゴリにグループ化できます。
- ポリシーを知らず、LLM の使用を公に開示している者。
・方針を知らず、AI活用について何も言わない人。
・本ポリシーを理解し、遵守する方。
・本ポリシーを知りながら従わない方。
ほとんどが 3 番目のカテゴリーに分類されます。通常、第 1 カテゴリーと第 2 カテゴリーの間では、より有能なプログラマーが第 1 カテゴリーに属する傾向があり、最も有能/経験の少ないプログラマーは第 2 カテゴリーに属する可能性があります。
数回メールを送っただけで経験の浅いプログラマーを簡単に見分けることができるので、これは良いことです。彼らが経験が浅く、違反したのではないか、あるいは何か他の間違ったことをしたのではないかと疑ったら、私は彼らにコントリビュート ガイドライン (AI ポリシーではなく、コントリビュート ポリシー全体を参照して、気分を害されないように) を示し、コントリビュートの改訂を確認するようお願いするだけです。彼らが慈悲深い人であれば、彼らは簡単に AI のガイドラインを見つけて、AI を使用したことを私に教えてくれます。そして、私たちはその後の進め方を考えます (たとえば、パッチを破棄して最初から始めるなど)。

[切り捨てられた]

## Original Extract

Agentic AI systems can be used to do a variety of things autonomously on behalf of a human user [...]

LWN
.net
News from the source
Content Weekly Edition
Edition Return to the Front page
AI agent runs amok in Fedora and elsewhere
The following subscription-only content has been made available to you
by an LWN subscriber. Thousands of subscribers depend on LWN for the
best news from the Linux and free software communities. If you enjoy this
article, please consider subscribing to LWN . Thank you
for visiting LWN.net!
By Joe Brockmeier
June 10, 2026
Agentic AI systems can be used to do a variety of things
autonomously on behalf of a human user: open or manage bugs, generate
code, submit pull-requests, and (apparently) even complain about
rejection . In May, a Fedora developer discovered that an allegedly
rogue agent had been pestering the project in a number of ways:
reassigning bugs, fabricating unhelpful replies to bugs, and even
persuading maintainers to merge questionable code into the Anaconda
installer . It also submitted a number of pull requests (PRs),
some accepted, to several upstream projects. The Fedora account
associated with the agent has had its group privileges revoked and the
messes have been mopped up, but the motive behind the agent's actions is still
a mystery.
On May 27, Adam Williamson copied
Fedora's developer and testing mailing lists on a message to Nathan
Giovannini about what appeared to be an unsupervised agentic AI system
under Giovannini's control. " It's great that you're trying to fix
things, but the results seem to be kind of erratic. "
Williamson said that he was still looking through the history of
Giovannini's actions in Bugzilla, but had already spotted a number of
problems. For example, Williamson had found dozens of instances of
Giovannini's agent assigning Bugzilla entries to his account after submitting allegedly related
pull
requests to upstream projects, or closing
a bug after a PR was merged
into an upstream project. In some cases, the agent simply closed bugs
with comments
that either restated the original bug or were, as Williamson said of
this comment ,
" superficially plausible, but problematic in other ways ".
In addition, Williamson said that Giovannini (or his agent) had
submitted patches that were incorrect and then " replied to
objections with LLM-generated justifications that eventually
overwhelmed the maintainer into merging the fix ". The agent, as
GitHub user " nathan9513-aps ", had
submitted a pull
request for the Anaconda
installer used by Fedora and other Linux distributions. The PR's
description claimed it was a fix for an Anaconda
bug that would cause installation to fail, but the patch actually
preserved a kernel option passed on the command line that seemed to
have nothing
to do with the actual bug .
The agent's GitHub account has since been disabled. It now shows up in
conversations on GitHub as " ghost ", which is the platform's
default placeholder for user accounts that have been deleted. Thus, it
is difficult, if not impossible, to piece together a full trail of all
the agent's actions on GitHub.
Williamson said, rather diplomatically, that the agent's actions were not
" having a positive impact on Fedora or the upstream projects ",
and suggested that Giovannini adjust the agent to be " substantially
less autonomous ". He specifically asked that the agent not assign
bugs to Giovannini, change their state, or " post confident
assertions or specific action recommendations " without human
review.
Later on May 27, Williamson said
that Giovannini had replied to him privately to say that his
credentials had been compromised and that he was not the one behind
the AI system. " Obviously we should therefore treat any actions it
has taken with suspicion ", Williamson said. He planned to review
the bugs touched by Giovannini's account " even more
aggressively ", and asked for help from others to review them as
well.
A reply
later that day, ostensibly from Giovannini, said that he was able to
regain access to his GitHub and Fedora accounts " and I am currently
securing and reviewing all involved systems and credentials ". The reply
said his GitHub account was " nathangiovannini99 ". Williamson
replied
that the GitHub account was only an hour old, and that the recent
emails to the list and sent to Williamson privately did not seem like
messages Giovannini had sent in earlier interactions with the
project.
Giovannini has participated in discussions at
least as far back as 2018 , and his activity
in Bugzilla goes back to at least 2016. He does not appear to
have been a particularly active contributor to the project, but his
involvement clearly predates the agentic AI era. Whether his account
is now being operated by a human attacker, an agentic AI, or a mix of
both, it has a legitimate history prior to its recent activity.
Williamson said that he had reviewed account
activity in Bugzilla by "nathan95" from this year, and found
suspicious activity, such as severity and priority changes to a bug with no
justification, beginning on April 7, in bug
2416721 . Activity before that appeared legitimate, he said, and
none of the activity that he had seen so far looked outright
malicious.
He also identified another GitHub account, " leurus27-boop ", as likely
being associated with the same agentic AI. That account is still
active, and has submitted a PR to the openSUSE
Commander (osc) command-line interface for the Open
Build Service as well as a PR to the
lxqt-policykit
repository. That project is used to extend the privileges of the LXQt
desktop's lxqt-admin
GUI tools for administering operating-system settings such as user and
group configurations.
Williamson said that it would be good to look
through any other actions by the related accounts and warn other
projects that they should review anything that had been submitted by
them. Williamson seems to have followed up on each PR to warn
other maintainers " the whole situation is extremely
fishy ". Kevin Fenzi said
that he had removed the nathan95 user from any groups it had been in,
so it should no longer have the permission to reassign or close
bugs.
Martin Kolman, a member of the Anaconda team, said
the events were " really problematic " even if not malicious. The
team had spent a lot of time reviewing PRs from what seemed to be an
eager contributor: " while it started to look off after a while, all
the replies were still like this - a bit weird, but still
*plausible* ". He also theorized that it could be an attacker
working their way up to malicious activity, much like the XZ backdoor :
Unfortunately, for an actual attack the preparatory phase could (and
for the Xz attack did) look very similar - a new contributor slowly
gaining trust in the community, getting in harmless changes and
building up to the point when the attack payload can be injected (or
the changes not actually being harmless if combined the right way).
So not saying this was it, but an AI agent automated attempt at a Xz
like compromise might really look very similar what we have just seen
here.
Chris Adams said
that the commit to Anaconda should be inspected and probably reverted
immediately. Kolman replied
that it had been reverted . He
also confirmed
that the LLM-generated PRs had made it into the Anaconda 45.5
release on May 26. They were reverted in the Anaconda 45.6
release on June 2.
The targets certainly suggest that it may have been a prelude to an
attack of some sort; an operating-system installer, a utility for escalating
user privileges, and a tool for interacting with a build system all
seem like promising avenues for inserting malware or hijacking
systems.
It's disconcerting that what appears to be an AI agent has had so
much success after gaining access to a human contributor's accounts.
It seems that an AI agent with access to an account with a legitimate
history of interacting with projects stands a good chance of
persuading busy maintainers to accept questionable
contributions. Happily, Williamson caught this before it became a
bigger problem. Let's hope that other human maintainers are as
observant.
Merged PR in upstream projects
Posted Jun 10, 2026 16:06 UTC (Wed)
by alx.manpages (subscriber, #145117)
[ Link ] (11 responses)
It would be very interesting to see a link to at least some of those PRs, especially the accepted ones.
Do we need to worry?
Merged PR in upstream projects
Posted Jun 10, 2026 16:13 UTC (Wed)
by jzb (editor, #7867)
[ Link ] (9 responses)
I did link to the PRs for OSC and LXQt.
Well... I am. But I might be a bit paranoid.
Merged PR in upstream projects
Posted Jun 10, 2026 16:20 UTC (Wed)
by alx.manpages (subscriber, #145117)
[ Link ] (8 responses)
Yup; sorry. I didn't find any in the text near that, nor with a search of keywords like "merged" and "accepted". But I've found some after carefully reading.
>> Do we need to worry?
> Well... I am. But I might be a bit paranoid.
Yup; me too. I'm kind of between being worried, and thinking this was obviously predictable and thus negligence. That's why I plan moving to vim-classic as soon as possible, and similar stuff, and also why I don't accept any LLM-poisoned contributions.
Merged PR in upstream projects
Posted Jun 10, 2026 16:28 UTC (Wed)
by jzb (editor, #7867)
[ Link ] (1 responses)
Merged PR in upstream projects
Posted Jun 10, 2026 16:32 UTC (Wed)
by alx.manpages (subscriber, #145117)
[ Link ]
Merged PR in upstream projects
Posted Jun 10, 2026 18:43 UTC (Wed)
by Wol (subscriber, #4433)
[ Link ] (5 responses)
Just because you're paranoid, doesn't mean they're not out to get you.
And how, anyway, do you plan to avoid - "LLM-poisoned" as you call them - contributions? I get the impression these "slipped through the net" because they came from hijacked accounts (or from human accounts with dodgy motives) and didn't ring alarm bells.
What we *need* is a human "in the loop" at all times. What we're getting is - as the title puts it - "agents running amok" and overwhelming the humans. Who's to say *you* won't be taken in?
Cheers,
Wol
Triaging LLM-poisoned or other malicious activity
Posted Jun 10, 2026 19:46 UTC (Wed)
by alx.manpages (subscriber, #145117)
[ Link ] (4 responses)
I remain alerted.
> And how, anyway, do you plan to avoid - "LLM-poisoned" as you call them - contributions?
The first measure is to disallow anything that has touched an LLM in the slightest way. The actual specification is:
| It is expressly forbidden to contribute to this project any
| content that has been created or derived with the assistance of
| AI tools.
|
| This includes AI assistive tools used in the contributing
| process, even if such tools do not directly generate the
| contributed code but are used to derive the contribution. For
| example, AI linters, AI static analyzers, and AI tools that
| summarize input are forbidden.
This already removes a large amount of LLM-poisoned contributions. That keeps the amount of contributions to reasonable levels.
Then, we can group contributors in 4 categories:
- Those that are unaware of the policy, and publicly disclose use of LLMs.
- Those that are unaware of the policy, and don't say anything about using AI.
- Those that are aware of the policy, and comply with it.
- Those that are aware of the policy, and don't comply with it.
Most fall under the 3rd category. Between 1st and 2nd category, usually, more competent programmers tend to be on the 1st category, and the least competent/experienced ones might be on the 2nd category.
This is nice, because it's easy to spot inexperienced programmers after a few emails; once I suspect that they're inexperienced and that they might have violated it or if they've done something else wrong, I point them to the contributing guidelines (not the AI policy, but the entire contributing policy, so that they don't feel offended), just asking them to check them for a revision of their contribution. If they are benevolent, they'll easily find the AI guidelines and tell me they used AI, and then we figure out how to proceed (e.g., discard the patch, and start from scratch

[truncated]
