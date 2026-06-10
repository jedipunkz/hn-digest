---
source: "https://derickrethans.nl/humans-in-the-llm-loop.html"
hn_url: "https://news.ycombinator.com/item?id=48481213"
title: "Humans in the LLM Loop"
article_title: "Humans in the LLM Loop — Derick Rethans"
author: "speckx"
captured_at: "2026-06-10T21:46:22Z"
capture_tool: "hn-digest"
hn_id: 48481213
score: 1
comments: 0
posted_at: "2026-06-10T19:13:18Z"
tags:
  - hacker-news
  - translated
---

# Humans in the LLM Loop

- HN: [48481213](https://news.ycombinator.com/item?id=48481213)
- Source: [derickrethans.nl](https://derickrethans.nl/humans-in-the-llm-loop.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T19:13:18Z

## Translation

タイトル: LLM ループ内の人間
記事のタイトル: LLM ループの中の人間 — Derick Rethans

記事本文:
ここ数週間、私は Xdebug のいくつかのバグ レポートに取り組んできました。その結果、Xdebug 3.5.3 リリースが完成しました。
これらのバグ レポートは、人間のみから送信されたものではなく、セキュリティ関連の問題に焦点を当て、2 つの異なるソースと方法論から LLM アシスタント ツールを使用する複数の人間から送信されました。
これらすべての問題は実際にバグであり、現在修正していますが、セキュリティへの影響が低いとさえ分類できるものはないと思います。
しかし、これらのレポートには他にも多くの問題がありました。報告書自体は不自然に冗長になる可能性があり、また「被害者」や「攻撃者」などの用語を使用してかなり警戒心を強める場合もあります。レポートに記載されているテストは、多くの場合最小限であり、時には不完全であり、提案された修正の一部も同様でした。
レポートを転送した担当者は、問題トラッカーに突然レポートが殺到しないように注意し、最初に私に連絡してくれました。彼らは、報告された問題について話し合うのにも役立ちました。
最初の 4 つのケースは、長年 PHP に貢献し、最近復帰した Ilia Alshanetsky によって報告されました。
最初のレポート #2421 は、IDE と Xdebug が通信するために使用する DBGp プロトコルを介したコマンドで間違ったオプション文字を送信することを扱います。
Xdebug は、ラテンアルファベットの小文字 26 文字と - 文字を表す、これらのうち 27 個の配列を割り当てます。 Xdebug が実行しなかったことは、オプション文字が実際に [a-z-] の範囲内にあり、 -@ または -\x00 を喜んで受け入れるかどうかを確認することでした。これにより、メモリ内の場所を上書きすることが可能になります。
提案されたパッチは問題ありませんでしたが、それに伴うテストは非常に読みにくかったです。ステップ デバッガーのテストにも既存のフレームワークを使用しませんでした。独自の .html を追加する必要がありました。
私もこの問題は次のような影響を与えると信じています

ld はファザーによって簡単に見つけられるので、これも追加しました。ファザーは同じ問題を約 5 秒以内に発見しましたが、幸いにも他には何も見つかりませんでした。
2 つ目の問題 #2422 は、ネットワークからコマンドを読み取るデバッガーのコードに制限がないことを訴えていました。
パッチはほとんど問題ありませんでしたが、テストがまた完全に面倒でした。既存のテスト フレームワークも使用していませんでした。また、DBGp コマンドには 64MB という面白いほど大きな (任意の) 制限が設定されており、64KB あれば簡単に十分です。ほとんどの場合、256 バイトで問題ありません。
3 番目の問題 #2423 では、プロファイリング ファイルやトレース ファイルを作成するときに Xdebug がシンボリックリンクをたどるべきではないと主張しています。
パッチは問題なく、簡単なものでしたが、テストは再び非常に読みにくく、何をしようとしているのかを理解するのが困難でした。また、テストをスキップするために一部の既存のヘルパーも利用していませんでした。それは次のことを思いつきました：
--スキップ--
<?php
if (PHP_OS_FAMILY === 'Windows') die('skip: Linux のみのシンボリックリンク セマンティクス');
?>
他の場所で使用されているものの代わりに:
--スキップ--
<?php
__DIR__ が必要です。 '/../utils.inc';
check_reqs('!win');
?>
Ilia の 4 番目で最後の問題 #2424 は、パーサーが空または大きなコマンド パケットを正しく処理しない Xdebug のコントロール ソケット機能を扱います。 LLM が提案したパッチは症状を修正しましたが、この問題の実際の原因は修正しませんでした。
2 番目のレポート セットは、PHP Foundation のエコシステム セキュリティ チームの取り組みの一環として、Volker によって非公開の要点として私に共有されました。
最初のものはバグ #2421 の複製でした。テストはステップ デバッガーではなくコントロール ソケット機能に焦点を当てましたが、根本的な問題と修正は同じでした。
2 番目の問題はバグ #2433 として追加しました。トレースを使用して xdebug.collect_assignments を有効にすると、Xdebug はこれを再作成する必要があります。

この名前を読みやすい方法で表示するために、いくつかのオペコードからの変数名を取得します。
ただし、この問題はリアルタイムの問題ではありません。これは、コマンドラインで -r オプションを使用して PHP コードを実行し、 xdebug.start_with_request が yes である場合にのみ発生します。何らかの理由で、PHP が -r を介してコードを実行すると、CLI バイナリは EXT_STMT オペコードを生成しません (Xdebug は、ステップ デバッグ中の中断にこれらを使用します)。そうしないと、境界外のメモリ読み取りが発生しなくなります。
また、LLM ツールは、変数名の再構築を担当する関数の 3 番目の引数が常に NULL であり、したがって余分であることにも気づいていませんでした。
同じコミットを通じてこれらの両方に対処し、テストを追加しましたが、これでもほとんどの状況で問題は発生しません。期待される結果を文書化しておくことは依然として良いことです。
別のレポートでは、Xdebug のトラッカーに 2 つの問題が見つかりました。 #2427 は、 xdebug.file_link_format 設定が単独の % で終わる場合の不正なメモリ読み取りに対処し、 #2429 は同様のレポートですが、 xdebug.trace_output_name に関するものです。
レポートでは 3 つの場所について言及していましたが、付随するテストでは、これが問題となる 3 つの状況のうち 1 つだけをカバーしました。つまり、トレース ファイル名の場合であり、 xdebug.file_link_format によるリンク ファイルのフォーマットやファイルのプロファイリングではありませんでした。
提案されたパッチも間違っており、末尾の % を保持する代わりに削除するものでした。末尾の % が処理されなかった 3 つの場所のうち 1 つは内部のみであったため、構成エラーによって引き起こされることはありませんでした。
このレポートに付属のテストは、問題を明らかにするのに役に立ちませんでした。問題を表示するには AddressSanitizer に依存していましたが、それを発生させることはできませんでした。このツールによるすべてのテストでは、バグの存在をテストするテストも提供されました。

、正しい結果がどうあるべきかではありません。
幸いなことに、valgrind ツールで Xdebug テスト スイートを使用すると、問題が発生しました。
さらなるレポート #2430 では、IDE がステップ デバッグ プロトコルを通じて、または開発者が直接 ::: という名前の「変数」の内容を要求した場合に問題が発生することが示されました。ステップ デバッガーは、「このクラスのすべての静的変数」を示すために :: を使用しますが、その後に : を続けることは無効です。
修正は良好でしたが、壊れた動作をテストするため、テスト ケースを直接使用することはできませんでした。報告されたテスト ケースの再現ケースが正しかったため、テストを書くのはかなり簡単でした。
そして、2 番目のリストの最後の問題である #2431 は、DBGp テストを実行するための独自の方法を再度発明し、動作することを示すテストの代わりに、動作が間違っていたことをテストしました。
コードが修正されたとしても、新しく正しいテストを実行すると、Xdebug がディレクトリをファイルとして開いて失敗するという別の問題も表面化します。
結論 : LLM ツールにはバグが見つかりましたが、特に画期的なものではありませんでした。一部のバグはファジングによって発見され、そのプロセスで使用されるリソースがはるかに少なくなりました。
クラッシュや潜在的なセキュリティ問題のほとんどは、コードが実行されるマシンに攻撃者がまだアクセスしていない場合、または Xdebug と通信する IDE を持っていない場合にのみ問題になります。
マシンにアクセスできる場合、これらのバグが存在しない場合は、さらに悪いことが起こる可能性があります。 DBGp を介して Xdebug にクライアント アクセスできる場合は、ファイル システム上のすべてのファイルの読み取りやコードの実行など、PHP が提供するすべての機能を利用できます。
生成されたテスト ケースは一般的に読みにくいか、不完全でした。モデルが考え出したパッチは、必ずしも包括的であったり、正しいものではありませんでした。私もSPです

AddressSanitizer に何かをさせるのに時間がかかりすぎて、失敗しました。
問題の原因とその理由が提供されていれば、私もこれらのパッチや実際のテスト ケースを自分ですぐに作成できたと思います。
これらのツールを自分で使えるようにするのに時間を費やすつもりはありませんが、自分たちが何をしているのかを知っている適切な人々の手を借りれば、対処する必要がある問題を見つけることができます。しかし、その価値は人間が結果を解釈することから生まれます。
この記事には短縮 URL が用意されています: https://drck.me/xdebug-llm-reports-k8y
これによる興味深い下流効果は、メンテナによるトリアージへの取り組み方さえも変化する可能性があることです。自動レポートが日常的になると、ベースラインの想定が「これはおそらく有効だろう」から「とにかくこれを独自に検証する必要がある」に変わり、ワークフローのコストが変化します。その結果、ワークフローのコストは、実際に経験するまでは予測するのが困難です。
Amazon の欲しいものリストはここにあります。
人間が作ったウェビング
前 |
ランダム |
次へ
`GET /icons/blank.gif HTTP/1.1" 200 398 " https://downloads.php.net/~windows/pec l/releases/?utm_source=chatgpt.com `
一体、それには `utm_source` が必要なのでしょうか?
キャッシュの有効期限が 900±15 秒のようにあいまいになればいいのにと思うことがあります。そうですね、雷を散らす群れ効果を防ぐためです。
エントランスとメインエントランスを作成しました
解決された問題 #2424: 無効なコマンドを含むコントロール ソケットがクラッシュする
作成されたトレースファイルを必ずクリーンアップしてください
解決された問題 #2433: トレース ファイル内の変数名を検索するリファクタリング…
ソングツグミのツイート: おはよう
# BirdPhotography # BirdsOfFediverse # 鳥 # 自然 # 写真 # ロンドン
#写真撮影 #鳥 #BirdPhotogaphy #BirdsOfMastodon #自然 #ロンドン #BirdsOfFediverse
ベンチを作成しました。ベンチを削除しました

## Original Extract

In the last few weeks, I have been working through some bug reports for Xdebug, that resulted in the Xdebug 3.5.3 release .
These bug reports did not come solely from humans, but rather from a mix of humans using LLM assistant tools, focussing on security related problems, from two different sources and methodologies.
Although all of these issues where indeed bugs that I now have fixed, I don't think any of them can be classified even as having a low security impact.
But there was a whole host of other issues with these reports. The reports themselves can be unnaturally verbose — and also fairly alarmist using terms like "victim" and "attacker". The tests that were present in the reports were often minimal, and sometimes incomplete, and so were some of their suggested fixes.
The humans forwarding the reports took care not to flood the issue tracker with reports out of the blue, and reached out to me first. They've also been helpful discussing the reported issues.
The first four cases were reported by Ilia Alshanetsky , a long time, and recently returned, contributor to PHP.
The first report, #2421 , deals with sending wrong option characters with commands through the DBGp protocol, that an IDE and Xdebug use to communicate.
Xdebug would allocate an array for 27 of these, representing the 26 lower case letters of the Latin alphabet, and the - character. What Xdebug did not do is to make sure the option letters were indeed in the range [a-z-] , and would happily accept -@ or -\x00 . This makes it possible to overwrite locations in memory.
The suggested patch was fine, but the test that went with it was very hard to read. It didn't use already exiting framework for testing the step debugger either — I had to add my own .
I also believe that this issue could as easily have been found by a fuzzer, which I added now as well . The fuzzer found the same problem in about five seconds, and luckily, nothing else either.
The second issue, #2422 complained that there was no limit in the debugger's code that reads commands from the network.
The patch was mostly fine, but the test was wholly cumbersome again — it didn't use the already existing testing framework either. It also picked a funnily large (arbitrary) limit of 64MB for DBGp commands, where 64KB would easily suffice — in most cases, 256 bytes would have been fine.
The third issue, #2423 , argues that Xdebug shouldn't follow symlinks when creating profiling or trace files.
The patch was OK and trivial, but the test was again very hard to read making it hard to figure out what it as trying to do. It also did not make use of some existing helpers to skip tests. It came up with:
--SKIPIF--
<?php
if (PHP_OS_FAMILY === 'Windows') die('skip: Linux-only symlink semantics');
?>
Instead of what is used everywhere else:
--SKIPIF--
<?php
require __DIR__ . '/../utils.inc';
check_reqs('!win');
?>
The fourth and last issue through Ilia, #2424 , deals with Xdebug's Control Socket functionality, where its parser would not handle empty or large command packets correctly. The LLM proposed patch fixed the symptoms, but not the actual cause of this issue.
The second set of reports were shared with me in a private gist by Volker, as part of the PHP Foundation's Ecosystem Security Team effort.
The first one was a duplicate of bug #2421 . The test focussed on the Control Socket functionality instead of the Step Debugger, but the underlying issue and fix were the same.
The second issue I added as bug #2433 . When you enable xdebug.collect_assignments with tracing, Xdebug needs to re-create the variable name from several opcodes in order to show this name in a readable way.
But the issue is not a real time problem, insofar this can only happen if you run PHP code on the command line through the -r option, and xdebug.start_with_request is yes . For some reason, when PHP runs code through -r , the CLI binary does not generate EXT_STMT opcodes (Xdebug uses these for breaking during step-debugging), which would otherwise prevent the out-of-bound memory read from happening.
The LLM tool also hadn't realised that the third argument to the function responsible for reassembling the variable name was always NULL , and hence superfluous.
I addressed these both through the same commit, and added a test, which would not exhibit the problem in most situations either. It is still good to have the expected outcome documented.
Another report resulted in two issues in Xdebug's tracker. #2427 addresses an incorrect memory read if the xdebug.file_link_format setting ends with a lone % , and #2429 , a similar report, but then for xdebug.trace_output_name .
Although the report mentioned three locations, the accompanying test only covered one of three situations where this was a problem: for trace file names, but not formatting link files through xdebug.file_link_format , nor profiling files.
The patch it suggested was also wrong, as it would remove the trailing % instead of keeping it. One of three locations where the trailing % was not handled, was internal only, and hence couldn't be triggered by making configuration errors.
The test that came with this report did not help me trying to show the problem. It relied on AddressSanitizer to show any problems, but I could not get that to happen. All the tests through this tool also provided tests that tested that the bug was present, and not what the correct result ought to be.
Luckily, using the Xdebug test suite with the valgrind tool showed the problem.
A further report, #2430 showed a problem if either an IDE through the step debugging protocol, or a developer directly, would request the contents of a "variable" named ::: . The step debugger uses :: to indicate "all the static variables for this class", and following that up with a : isn't valid.
The fix was good, but I couldn't directly use the test case, as it tested for the broken behaviour. The test was fairly trivial to write as the reproduce case in the reported test case was correct.
And the last issue from the second list, #2431 , again reinvented its own way for doing DBGp tests, and also tested that the behaviour was wrong , instead of a test to show that it now works.
Even with the code fixed, the new correct test would also surface another issue, as it would have resulted in Xdebug to open a directory as it was a file, and then fail.
Conclusion : Although the LLM tools did find bugs, they were not particularly groundbreaking. Some of the bugs would also have been found by fuzzing, and used a lot less resources in that process.
Most of the crashes and potential security issues would only be a problem if an attacker didn't already have access to the machine that the code runs on itself, or have an IDE talking to Xdebug already.
If you have access to the machine, you can do worse without these bugs present. If you have client access to Xdebug through DBGp, you would have all the functionality that PHP provides, including reading all files on the file system and running code.
The generated test cases were generally hard to read, or incomplete. The patches that the models came up with were not always comprehensive, or correct. I also spent too much time getting AddressSanitizer to do anything, unsuccessfully.
I think I would have been as quick writing these patches and actual test cases myself, when provided with the issues' causes and the reasoning that was provided.
I don't think I'll be spending time trying to get these tools to work myself, but in the right hands with people that know what they're doing, they can find issues that needs to be addressed. But the value comes from the humans interpreting their results.
This article has a short URL available: https://drck.me/xdebug-llm-reports-k8y
The interesting downstream effect of this is that it might shift how maintainers even approach triage. When automated reports become routine, the baseline assumption shifts from "this is probably valid" to "I need to verify this independently anyway," which changes the workflow cost in ways that are hard to predict until you are living through it.
My Amazon wishlist can be found here .
Human-Made Webring
previous |
random |
next
`GET /icons/blank.gif HTTP/1.1" 200 398 " https:// downloads.php.net/~windows/pec l/releases/?utm_source=chatgpt.com `
WTF does that need an `utm_source` for?
I sometimes wish cache expiries could be fuzzy, like 900±15 seconds. Y'know, to prevent the thundering herd effect.
Created an entrance and a main entrance
Fixed issue #2424: Control-socket with invalid commands crashes
Make sure to clean up the created tracefile
Fixed issue #2433: Refactor finding variable names in trace files wit…
Song Thrush Tweets: Good Morning
# BirdPhotography # BirdsOfFediverse # Birds # Nature # Photography # london
# photography # Birds # BirdPhotogaphy # BirdsOfMastodon # nature # london # BirdsOfFediverse
Created a bench; Deleted a bench
