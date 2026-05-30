---
source: "https://blog.calif.io/p/an-ai-audit-of-freebsd"
hn_url: "https://news.ycombinator.com/item?id=48329215"
title: "An AI Audit of FreeBSD"
article_title: "An AI audit of FreeBSD - Calif"
author: "wslh"
captured_at: "2026-05-30T11:38:16Z"
capture_tool: "hn-digest"
hn_id: 48329215
score: 8
comments: 0
posted_at: "2026-05-29T21:01:49Z"
tags:
  - hacker-news
  - translated
---

# An AI Audit of FreeBSD

- HN: [48329215](https://news.ycombinator.com/item?id=48329215)
- Source: [blog.calif.io](https://blog.calif.io/p/an-ai-audit-of-freebsd)
- Score: 8
- Comments: 0
- Posted: 2026-05-29T21:01:49Z

## Translation

タイトル: FreeBSD の AI 監査
記事のタイトル: FreeBSD の AI 監査 - カリフォルニア
説明: 3 つの RCE、5 つの LPE、1 つの bhyve エスケープを含む 15 のカーネル バグ。

記事本文:
FreeBSD の AI 監査 - カリフォルニア
カリフォルニア
FreeBSD の AI 監査
15 個のカーネル バグ (3 個の RCE、5 個の LPE、1 個の bhyve エスケープを含む)。
12 1 シェア AI を使用してインターネットをハッキングするこのキャンペーンを開始して以来、多くの人がすでに知っていることを学びました。それは、インターネットはボランティアによって運営されているということです。インターネットのセキュリティと文化にとって重要なプロジェクトには、少人数のグループ、場合によっては 1 人でスタッフが配置されています。インターネット上のほぼすべてのリモート シェルを保護する OpenSSH は、1 人のオーストラリア人 (やあ、ダミアン!) が率いる小規模なチームによって保守されています。
私たちはメンテナーたちに何か借りがあると感じています。インターネットと、それを実行するオープンソース ソフトウェアがなければ、私たちは学んだことを学び、友達を作り、今日のキャリアを持つことはできなかったでしょう。そこで私たちは、専門家と AI を、そのヘルプを利用できるオープンソース プロジェクトと組み合わせることにしました。 FreeBSD が私たちの出発点です。
3 月末に、私たちは最初の AI 支援 FreeBSD リモート カーネル エクスプロイトを公開しました。今月初めに、 exeCVE で CVE を報告しました。また、めったに使用されないモジュールで 3 つの RCE が発生したことも報告しました。チームが手薄になっているのを見て、私たちは単にチームを増やすだけではなく、もっと助けられるように努めるべきだと考え、彼らに手を差し伸べました。チームは私たちに何に焦点を当てるべきかを教えてくれたので、AI に任せました。
その作業の最初の数週間以内に、監査によりさらに多くのバグが明らかになりました。
少数のメモリ漏洩と DoS
合計 15 件のバグが報告されました。すべてはカーネル内にあります。また、それらの一部を見つけるために使用した監査スキルもチームと共有しました。
この投稿は、そこに至るまでの経緯についてです。
私たちが FreeBSD チームと話し合ったとき、次の 2 つのことに同意しました。
FreeBSD のバグを見つけるのにコストがかかります。
私たちがいなくなった後も、FreeBSD チームがさらなるバグを発見、排除、防止できるよう支援してください。
CVE 番号を追求するつもりはありません

バー数または投稿バグ数。私たちは、プロジェクトを実行している人々の役に立ちたいだけです。
FreeBSD のような広く使用されているオープンソース プロジェクトのメンテナはレポートに溺れており、彼らの関心はこの企業全体で最も高価なリソースです。役に立つための第一のルールは、無駄にしないことです。私たちがまとめたいくつかの点:
重大なバグまたは重大なバグのみを送信します。当社は、重大な脆弱性または重大な脆弱性と思われるものに焦点を当てて送信レポートを作成します。場合によっては、高いと思われるバグが、メンテナによる詳細な検査によってダウングレードされることもありますが、私たちは議論するのではなく、主にメンテナ自身のスコアに従っています。
レポートは短くしてください。誰もが短いレポートを好みます。たった 1 行の概要と PoC は、15 ページにわたる蛇行分析よりもはるかに優れています。誰かがそれを求めれば、詳細な調査をフォローアップすることができます。
パッチを提案しますが、それを強要しないでください。メンテナの中には、提案されたパッチを受け取るのが大好きな人もいます。自分で修正を書くことを好む人もいます。デフォルトでは、提案として明確にラベル付けされたパッチをレポートに含めます。そのため、メンテナはやり取りすることなく、それを採用、変更、または無視できます。
人々と一緒に時間を過ごしましょう。電子メールとトラッカー チケットは必要ですが、早い段階で 1 回のビデオ通話を行うことは、注意深い問題のテンプレートをいくつも作成するよりも効果的です。 FreeBSD チームとの最初の会議の後、私たちは彼らとの直接のチャネルを設定しました。それ以来報告されたバグの多くは、報告から数日で修正されるまでになりました。
FreeBSD は、私たちがこのようなコラボレーションについて公に書いている初めてのものですが、それだけではありません。インターネットの稼働を維持する他のプロジェクトでも同様の作業がすでに進行中であり、それらの取り組みが成熟するにつれて、さらに多くのことを共有する予定です。
MAD Bugs の投稿にはウェアーズ ドロップが含まれている必要があるため、今日は 3 つの LPE のエクスプロイトと書き込みを公開します。

:
setcred (CVE-2026-45250) : kern_setcred_copyin_supp_groups での 1 文字のサイズの混乱が、 user_setcred のフレームでのスタック オーバーフロー、そしてローカル ルート シェルに変わります。 14.3 と 15.0 には同じソースのバグが存在しますが、悪用できるのは FreeBSD 14.4 のみです。
ptrace (CVE-2026-45253) : ptrace(PT_SC_REMOTE) は、リダイレクトされた syscall 番号の境界チェックをスキップし、LPE にチェーンする sysent テーブルに境界外のインデックスを作成します。
procdesc (CVE-2026-45251) : procdesc_free() は、ポーリング ウェイターを排出せずに、埋め込まれた pd_selinfo を持つ構造体 procdesc を解放します。 SCM_RIGHTS ファイル子孫を使用してスロットを再利用し、2 つの古い TAILQ_REMOVE を起動して、任意のカーネル ポインター書き込みを取得します。
エクスプロイトと書き込みは AI によって書かれました。私たちは、2026 年の AI 脆弱性研究の様子を示す歴史的成果物として、AI テキストを現状のまま保持することにしました。一方、エクスプロイトはすべて私たちによって検証されており、機能します。これらを公開することで、より多くの人がこれらのテクニックから学び、FreeBSD にさらなる支援をもたらすことを願っています。監査で得られた残りのバグは、FreeBSD チームが修正を出荷する際にリリースされる予定です。
好奇心旺盛な読者のために、リポジトリにはいくつかのボーナス エクスプロイトも含まれています。これらのほとんどは、動作する PoC なしで出荷された公開 FreeBSD アドバイザリから AI によって調理されたものです。
私たちと協力し、真剣に取り組んでくれた FreeBSD チームに感謝します。トークンのために OpenAI と Anthropic に。そして、わずかな信用とわずかな人手でインターネットを運営し続けているすべてのメンテナーに感謝します。
12 1 シェア この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

15 kernel bugs, including 3 RCEs, 5 LPEs, and 1 bhyve escape.

An AI audit of FreeBSD - Calif
Calif
Subscribe Sign in An AI audit of FreeBSD
15 kernel bugs, including 3 RCEs, 5 LPEs, and 1 bhyve escape.
12 1 Share Since we started this campaign of hacking the Internet with AI , we’ve learned something many of you already knew: the Internet runs on volunteers. Projects that are critical to Internet security and culture are staffed by tiny groups of people, sometimes one person. OpenSSH, which protects almost every remote shell on the Internet, is maintained by a small team led by a single Aussie (Hi Damien!).
We feel like we owe these maintainers something. Without the Internet, and the open source software that runs it, we would not have learned what we learned, made the friends we made, or had the careers we have today. So we decided to pair our experts and our AI with open source projects that could use the help. FreeBSD is where we started.
At the end of March we published the first AI-assisted FreeBSD remote kernel exploit . Earlier this month we reported a CVE in exeCVE . We also reported 3 RCEs in a rarely used module. Seeing the team stretched thin, we thought we should try to help more than just adding to the pile, and reached out to them. The team told us what to focus on, and we let the AI go brr.
Within the first few weeks of that work, the audit surfaced more bugs:
a handful of memory disclosures and DoS
In total, we have reported 15 bugs. All in the kernel. We have also shared the audit skill we used to find some of them with the team.
This post is about how we got there.
When we sat down with the FreeBSD team, we agreed on two things:
Make finding bugs in FreeBSD more expensive.
Help the FreeBSD team find, eliminate and prevent more bugs after we are no longer around.
We are not trying to chase CVE numbers or post bug counts. We just want to be useful to the people running the project.
Maintainers of widely-used open source projects like FreeBSD are drowning in reports, and their attention is the most expensive resource in this whole enterprise. The first rule of being useful is to not waste it. A few things we have converged on:
Send only high or critical bugs. We focus our outbound reports on what we believe are high or critical vulnerabilities. Sometimes a bug we think is high gets downgraded by the maintainers on closer inspection, and we largely follow their own scoring rather than arguing.
Keep reports short. Everyone likes a short report. A one-liner and a PoC is much better than fifteen pages of meandering analysis. The deep dive can go in a follow-up if anyone asks for it.
Suggest patches, but do not insist on them. Some maintainers love receiving suggested patches; some prefer to write the fix themselves. We default to including a patch in the report, clearly labeled as a suggestion, so the maintainer can take it, modify it, or ignore it without any back-and-forth.
Spend time with people. Email and tracker tickets are necessary, but a single video call early on does more for the working relationship than any number of careful issue templates. After our first meeting with the FreeBSD team, we set up a direct channel with them, and many of the bugs we have reported since then have gone from report to fix in days.
FreeBSD is the first such collaboration we are writing about publicly, but it is not the only one. Similar work is already underway with other projects that keep the Internet running, and we plan to share more as those efforts mature.
A MAD Bugs post must include some warez drops, so today we are publishing exploits and writeups for three of the LPEs:
setcred (CVE-2026-45250) : a one-character sizeof confusion in kern_setcred_copyin_supp_groups turns into a stack overflow in user_setcred 's frame and then a local root shell. Only FreeBSD 14.4 is exploitable, despite the same source bug being present in 14.3 and 15.0.
ptrace (CVE-2026-45253) : ptrace(PT_SC_REMOTE) skips a bounds check on the redirected syscall number, giving out-of-bounds indexing into the sysent table that we chain into LPE.
procdesc (CVE-2026-45251) : procdesc_free() frees a struct procdesc with an embedded pd_selinfo without draining poll waiters. We reclaim the slot with SCM_RIGHTS filedescents, fire two stale TAILQ_REMOVE s, and get arbitrary kernel-pointer writes.
The exploits and the writeups were written by AI. We have decided to keep the AI text as-is, as a historical artifact showing what AI vulnerability research looked like in 2026. The exploits, on the other hand, are all verified by us, and they work. By publishing them, we hope more people can learn from these techniques and bring more help to FreeBSD. The remaining bugs from the audit will be released as the FreeBSD team ships the fixes.
For curious readers, the repository also contains a few bonus exploits, mostly cooked by the AI from public FreeBSD advisories that shipped without working PoCs.
To the FreeBSD team, for working with us and for taking the work seriously. To OpenAI and Anthropic, for the tokens. And to all maintainers who keep the Internet running with very little credit and very few hands: thank you.
12 1 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
