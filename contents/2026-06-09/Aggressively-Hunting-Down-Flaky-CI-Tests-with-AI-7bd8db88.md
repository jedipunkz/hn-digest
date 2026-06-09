---
source: "https://openwisp.org/blog/aggressively-hunting-down-flaky-ci-tests-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=48463235"
title: "Aggressively Hunting Down Flaky CI Tests with AI"
article_title: "Aggressively Hunting Down Flaky CI Tests with AI - OpenWISP Blog"
author: "nemesisdesign"
captured_at: "2026-06-09T18:52:51Z"
capture_tool: "hn-digest"
hn_id: 48463235
score: 5
comments: 0
posted_at: "2026-06-09T16:25:58Z"
tags:
  - hacker-news
  - translated
---

# Aggressively Hunting Down Flaky CI Tests with AI

- HN: [48463235](https://news.ycombinator.com/item?id=48463235)
- Source: [openwisp.org](https://openwisp.org/blog/aggressively-hunting-down-flaky-ci-tests-with-ai/)
- Score: 5
- Comments: 0
- Posted: 2026-06-09T16:25:58Z

## Translation

タイトル: AI を使用して不安定な CI テストを積極的に追い詰める
記事のタイトル: AI を使用して不安定な CI テストを積極的に追い詰める - OpenWISP ブログ

記事本文:
機能 よくある質問について ソース コード ロードマップ 歴史 チーム パートナー スポンサーシップ サポート コミュニティ サポート 商用サポート プロのように展開 ドキュメント ブログ デモ 展開 AI を使用して脆弱な CI テストを積極的に追跡
公開日: 2026 年 6 月 8 日、フェデリコ・カポアーノ カテゴリー: ニュース タグ: テスト CI Devops AI 不安定な CI クラッシュの AI 支援デバッグ OpenWISP に貢献したことがある方は、おそらく不安定な CI に会ったことがあるでしょう。テストはマシン上で合格し、CI 上で失敗し、ジョブを再実行すると再び合格します。ほとんどの OpenWISP モジュールは、Selenium ブラウザ テストを備えた Django アプリケーションであり、時が経つにつれて、これらの断続的な障害は小さな迷惑以上のものになりました。赤いビルドが 1 分後には緑色になったり、コントリビュータはプル リクエストのブロックを解除するためにジョブを再実行したり、メンテナは赤いビルドが本物なのか、それとも別のゴーストなのかを再確認する方法を学習したりします。自己修復ボットとそれが隠した負債
不安定な実行によって全員がブロックされるのを防ぐために、私は Sarthak Tyagi に、共有ツールに小さなセーフティ ネットを構築するのを手伝ってくれるように頼みました。それは、失敗した実行を検査し、潜在的な修正の推奨事項を提供し、既知の不安定な障害の兆候を認識し、失敗したビルドを自動的に再開する openwisp-utils の再利用可能な GitHub Actions ワークフローです。これについては、自動化された CI 障害ボットのドキュメントで読むことができます。ボットはパイプラインを動かし続け、手動での再実行を大幅に節約しました。しかし、それは病気を治療するのではなく、症状を治療するものでした。本当のバグはまだ存在しており、いくつかのモジュールに分散していたため、それらを追跡する時間が見つかりませんでした。息抜きは便利ですが、根本原因に立ち返らなければ、それは静かに借金となります。私が使用したプロセス
そこで、コアに余裕のあるマシン上で、疲れを知らないアシスタントとして AI コーディング エージェントを使用することにしました。私はまだ懐疑的でした。数モン

以前、私はまさにこの種のデバッグで AI エージェントを試しましたが、役に立ちませんでした。自信を持って間違っており、間違った手がかりを追いかけることに熱心で、長時間の調査を一緒に行うことができませんでした。それ以来、ツールは改善され、私のプロンプトも改善されましたが、最も困難なバグに賭ける前に、そのアプローチがまったく機能するという証拠が必要でした。有益だったのは、エージェントが特別な洞察力を持っていたことではありません。私は、失敗した CI ログを読み取り、パターンを比較し、修正を試み、テストを再度実行するという、自分自身で行っていたであろう繰り返し作業を実行できるようにしました。これを手動で行うと、少なくとも丸 1 日は集中して作業する必要がありました。代わりに、Opus 4.8 の Claude Code を高エフォート モードで使用し、他のことで忙しい間も最小限の監視で動作させ続けました。プロセスは大まかに次のようになります。失敗し、CI 失敗ボットによって再起動された不安定な CI ビルドの出力を取得します。これらのログを分析し、最も一般的な障害を特定し、それぞれに対して考えられる解決策を探します。調査結果を、検査できるローカルの Markdown レポートに書き込みます。このレポートの簡単な成果から始めましょう。変更によってターゲットとしていた不安定な障害が実際に減少したことをエージェントが示すまで、ローカルでテストを実行し続けます。ブランチを GitHub にプッシュし、GitHub Actions CI ジョブを再起動し続けることで、1 つの環境だけを信頼するのではなく、ローカルのストレス実行と繰り返しの CI 実行を比較できます。 CodeRabbit からのフィードバックに対処します。エージェントが別の繰り返し発生する障害パターンを見つけた場合は、常にレポートを更新します。同じ不安定な障害が引き続き発生する場合は、対象とした変更を無効として扱い、別の解決策を探すためにエージェントを送り返します。更新されたレポートを時々読み、エージェントにヒントを与え、漂流が始まったときに優先順位や方法論を導きます。ループは

シンプル: 失敗を収集し、それらをランク付けし、最も簡単で役立つ修正を試みます。証明
[切り捨てられた]
私は openwisp-controller から始めました。不安定な障害は煩わしいものの、より扱いやすく、最も恐ろしいクラッシュは当面放置しておきました。予想以上にうまくいきました。私は問題を渡さず、答えを待ちました。私はエージェントを使用して、自分自身の通常のデバッグ ループの時間のかかる部分を実行しました。つまり、テストを実行し続け、障害を検査し、小さな変更を試し、それぞれの変更が本当に変化したかどうかを確認しました。このプロセスにより、いくつかの不安定なテストが修正され、さらに興味深いことに、その一部の背後にある本当のバグが明らかになりました。バックグラウンド タスクが、削除された行を野良 INSERT で復活させていたのです。結果として生じる FOREIGN KEY エラーにより、SQLite テスト状態が破損し、Selenium ブラウザが回復できないページを待機したままになる可能性があります。同じ作業により、セッション ストレージも共有キャッシュから移動され、本番環境の強化を兼ねたテスト修正が行われました。
それは何も魔法ではありませんでした。以前の試みから変わったのは、エージェントを神託者のように扱うのではなく、予備のマシン上で非常に忍耐強い手のように扱うことです。私は、再現し、何かを変更し、再度実行し、行き止まりを反証し、前進し続けるという、私自身がたどったであろう同じループを押し続けました。このエージェントは、私がチェックインし、軌道を修正し、他の作業を続けている間、何時間もそのループ内にとどまることができるため、役に立ちました。この成功のおかげで、私は同じ設定を、私が静かに無視していた別のケース、つまり openwisp-monitoring における頻繁で残忍な不安定な失敗グループに指摘する勇気を得ることができました。最初の修正セットではまだクラッシュが発生していませんでした。これは、時系列テストにおける InfluxDB UDP タイミング レースのクラスターでした。 UDP リスナーは書き込みを非同期にフラッシュするため、テストによって書き込まれたデータは常にフラッシュされるわけではありません。

同じテストでそれを読み戻すときにクエリ可能です。修正内容は、固定スリープの信頼を停止し、読み取りが安定するまでポーリングし、バーストがドロップされないようにテスト UDP 読み取りバッファーを拡大し、チェックを実行する前に Wi-Fi クライアント データを待機し、テストで信頼性の高い待機を挿入できない場合に UDP 実行からいくつかの読み取り後書き込みアサーションを除外することでした。修正不可能と思われた厄介なセグメンテーション違反
OpenWISP モニタリングで不安定なテストの修正に取り組んでいるときに、おなじみの障害が再び発生しました。二重解放または破損の後に致命的な Python エラーが発生しました: セグメンテーション フォールト。これはアサーション、タイムアウト、またはターゲットを外した Selenium のクリックではありませんでした。それはCレベルの激しいクラッシュでした。実行中にライブ テスト サーバーを強制終了し、Selenium テスト スイート全体も一緒にダウンさせました。おそらく 14 回の実行に 1 回、一部のジョブで発生し、他のジョブでは発生しませんでした。まさに、ログを読み直しても特定するのがほぼ不可能な種類のバグです。私はこれまでに何度もこれを目にし、すでにデバッグを試みていましたが、有益な結論に十分早く到達することができませんでした。他の仕事が優先され続けたので、手放さなければなりませんでした。今回、私は最近の勝利に元気づけられ、夜に燃やすための予備のトークンがたくさんあったので、営業時間外にエージェントを実行したままにすることにしました。ブルートフォース、そして分析
リラックスしてテレビシリーズを見たり、ゲームをしたりしていると、ラップトップのファンがおかしくなり始めました。上の htop スクリーンショットは、エージェントがラップトップのハードウェア リソースをフル容量で使用していることを示しています。エージェントは私に代わって、総当たり攻撃と各実行の分析という重労働を行ってくれました。私の指示の下、Selenium スイートを何度も実行し、各ワーカーが互いに干渉しないように分離し、CI 環境に一致する専用の Python 3.13 ツールチェーンを立ち上げる並列テスト ハーネスの構築に役立ちました。

オンメント。このプロセス中に作成されたサンプル テスト スクリプトは次のとおりです: #!/bin/bash
# TestDashboardMap で SQLite ダブルフリーをローカルに再現します。
cd /home/nemesis/Code/ow-oss/openwisp-monitoring ||出口1
OUT = /home/nemesis/Code/ow-oss/openwisp-monitoring/repro_results.md
SIG = 'ダブルフリー|フリー\(\):|破損|致命的なPythonエラー|セグメンテーション違反|コアダンプ'
エクスポート SELENIUM_HEADLESS = 1 DJANGO_SETTINGS_MODULE = openwisp2.settings
TARGET = openwisp_monitoring.tests.test_selenium.TestDashboardMap
ITERS = " ${ 1 :- 25 } "
: > " $OUT "
echo "# ローカル再現 - $ITERS iters - $( date -u + '%H:%M:%S' ) UTC" >> " $OUT "
for n in $( seq 1 " $ITERS " ) ;する
rm -f テスト/openwisp-monitoring-tests.db* テスト/openwisp-monitoring.db* 2 >/dev/null
log = $( タイムアウト 600 Python テスト/manage.py テスト " $TARGET " --noinput 2 > & 1 )
c = $( printf '%s' " $log " | grep -ciE " $SIG " )
ok = $( printf '%s' " $log " | grep -cE '^OK$' )
echo "- iter $n : crash= $c ok= $ok $( date -u +%H:%M:%S ) " >> " $OUT "
if [" $c " -gt 0 ] ;それから
echo " >>> iter $n <<<" >> " $OUT " でクラッシュ
printf '%s' " $log " | grep -iE " $SIG " | sed -E 's/^/ /' >> " $OUT "
休憩
フィ
完了しました
echo "DONE $( date -u + '%H:%M:%S' ) UTC" >> " $OUT "
利点は創造性ではありませんでした。それは、私が作業の方向性を監督している間、機械は退屈な部分を続けてくれるということでした。数百回の実行後、ついに実際の測定可能なシグナルが得られました。クラッシュは Python 3.13 でのみ約 7% の確率で再現されました。すべての Python 3
[切り捨てられた]
原因はGeoDjangoに隠れていた。 Django の SpatiaLite バックエンドは、新しい接続用の SpatiaLite ライブラリ パスを構築する際に ctypes.util.find_library("spatialite") を呼び出します。 Linux では、find_library が ldconfig サブプロセスをフォークする場合があります。マルチスレッドのライブサーバーでそれを繰り返し行うのは悪い組み合わせです

国家: 1 つのスレッドが SQLite 接続のクリーンアップ中に別のスレッドがフォークする可能性があり、C コードの奥深くにあるため、不適切なインターリーブは、友好的な Python トレースバックではなくヒープ破損として表示されます。ほとんどの場合はそれで済みます。 Python 3.13 では、「ほとんどの場合」が断続的な二重解放になる程度にタイミングが変更されました。修正は小さくて少し退屈ですが、これが最良の修正です。 SpatiaLite バックエンド ラッパーでは、サブプロセスが何度もフォークされないように、ライブラリ ルックアップをメモ化します。 Selenium テスト ミックスインでは、デフォルトのバックエンドが SQLite または SpatiaLite の場合に、Django の BaseDatabaseWrapper.connect および BaseDatabaseWrapper._close を中心に接続のオープンとクローズをシリアル化します。そのため、接続ライフサイクルがプロセス全体に及び、ライブサーバー テストの間、退屈なシリアル化が行われます。両方の要素を適切に配置すると、90 回を超える実行でクラッシュ率が 7% からゼロになり、その後、共有修正の CI 再実行が繰り返されてもクラッシュ率は維持されました。最良の部分: 共有される

## Original Extract

Features About FAQ Source Code Roadmap History Team Partners Sponsorship Support Community Support Commercial Support Deploy Like a Pro Docs Blog Demo Deploy Aggressively Hunting Down Flaky CI Tests with AI
Published on Mon 08 June 2026 by Federico Capoano Category: news Tags: testing ci devops ai AI-assisted debugging of a flaky CI crash If you have ever contributed to OpenWISP, you have probably met flaky CI. A test passes on your machine, fails on CI, then passes again when you re-run the job. Most OpenWISP modules are Django applications with Selenium browser tests, and over time those intermittent failures became more than a minor annoyance: red builds that were green a minute later, contributors re-running jobs to unblock their pull requests, and maintainers learning to double-check whether a red build was real or just another ghost. A self-healing bot, and the debt it hid
To stop flaky runs from blocking everyone, I asked Sarthak Tyagi to help me build a small safety net into our shared tooling: a reusable GitHub Actions workflow in openwisp-utils that inspects a failed run, offers recommendations for potential fixes, recognizes the signatures of known flaky failures, and automatically restarts the failed build. You can read about it in the docs for the automated CI failure bot . The bot kept the pipeline moving and saved us a lot of manual re-runs. But it treated the symptom, not the disease. The real bugs were still there, scattered across several modules, and we never found the time to chase them down. Breathing room is useful, but if you never come back to the root cause it quietly turns into debt. The process I used
So I decided to use an AI coding agent as a tireless assistant on a machine with cores to spare. I was still skeptical. A few months earlier I had tried an AI agent on exactly this kind of debugging and it had been useless: confidently wrong, eager to chase the wrong clue, and unable to hold a long investigation together. The tools had improved since then, and so had my prompting, but before betting on the hardest bug I wanted proof that the approach could work at all. The useful part was not that the agent had some special insight. I made it do the repetitive work I would have done myself: read failed CI logs, compare patterns, try fixes, and run the tests again. Doing that manually would have taken me at least a full day of focused work. Instead, I used Claude Code with Opus 4.8 in high-effort mode and kept it working with minimal supervision while I was busy with other things. The process looked roughly like this: Fetch the output of flaky CI builds that had failed and then been restarted by the CI failure bot. Analyze those logs, identify the most common failures, and look for possible solutions to each one. Write the findings into a local Markdown report that I could inspect. Start with the low-hanging fruit from that report. Keep running tests locally until the agent could show that a change really reduced the flaky failures it was targeting. Push branches to GitHub and keep restarting GitHub Actions CI jobs, so we could compare local stress runs with repeated CI runs instead of trusting just one environment. Address feedback from CodeRabbit . Update the report whenever the agent found another recurring failure pattern. Treat a targeted change as invalid if the same flaky failure kept showing up, then send the agent back to look for a different solution. Read the updated report from time to time, give the agent hints, and steer priorities or methodology when it started drifting. The loop was simple: collect failures, rank them, try the easiest useful fix, prov
[truncated]
I started with openwisp-controller , where the flaky failures were annoying but more tractable, and left the scariest crash alone for the moment. It went better than expected. I did not hand the problem over and wait for an answer. I used the agent to do the time-consuming part of my own normal debugging loop: keep running tests, inspect failures, try small changes, and verify whether each change really moved the needle. That process fixed several flaky tests and, more interestingly, exposed a real bug behind some of them: background tasks were resurrecting already-deleted rows with a stray INSERT . The resulting FOREIGN KEY error could corrupt the SQLite test state and leave the Selenium browser waiting for a page that would never recover. The same work also moved session storage out of the shared cache, a test fix that doubled as production hardening.
Nothing about it was magic. What changed since my earlier attempt was that I treated the agent like a very patient pair of hands on a spare machine, not like an oracle . I kept pushing it through the same loop I would have followed myself: reproduce, change one thing, run it again, disprove the dead end, keep going. The agent was useful because it could stay inside that loop for hours while I checked in, corrected course, and kept working on other things. That success is what made me brave enough to point the same setup at another case I had been quietly ignoring: a group of frequent and brutal flaky failures in openwisp-monitoring . The first set of fixes there was not the crash yet. It was a cluster of InfluxDB UDP timing races in the timeseries tests. The UDP listener flushes writes asynchronously, so data written by a test was not always queryable when the same test read it back. The fixes were to stop trusting fixed sleeps, poll until reads became stable, enlarge the test UDP read buffer so bursts were not dropped, wait for wifi-client data before running checks, and exclude a few read-after-write assertions from the UDP run when the test could not insert a reliable wait. A nasty segmentation fault that seemed unfixable
While working on the fixes for the flaky tests in OpenWISP Monitoring, a familiar failure showed up again: double free or corruption followed by Fatal Python error: Segmentation fault . This was not an assertion, a timeout, or a Selenium click that missed its target. It was a hard C-level crash. It killed the live test server mid-run and dragged the entire Selenium test suite down with it. It happened maybe once in fourteen runs, on some jobs and not others. Exactly the kind of bug that is almost impossible to pin down by re-reading logs. I had seen this many times before and had already tried to debug it, but I never reached a useful conclusion quickly enough. Other work kept taking priority, so I had to let it go. This time I was galvanized by the recent wins and had more spare tokens to burn at night, so I decided to leave the agent running after hours. Brute force, then analysis
While I was relaxing, watching TV series, and playing games, my laptop fan started going crazy. The htop screenshot above shows the agent using my laptop's hardware resources at full capacity. The agent was doing the heavy work for me: brute-forcing and analyzing each run. Under my direction, it helped build a parallel test harness that hammered the Selenium suite over and over, isolated each worker so they would not clobber one another, and stood up a dedicated Python 3.13 toolchain to match the CI environment. Here's a sample test script created during this process: #!/bin/bash
# Reproduce the SQLite double-free in TestDashboardMap locally.
cd /home/nemesis/Code/ow-oss/openwisp-monitoring || exit 1
OUT = /home/nemesis/Code/ow-oss/openwisp-monitoring/repro_results.md
SIG = 'double free|free\(\):|corruption|Fatal Python error|Segmentation fault|core dumped'
export SELENIUM_HEADLESS = 1 DJANGO_SETTINGS_MODULE = openwisp2.settings
TARGET = openwisp_monitoring.tests.test_selenium.TestDashboardMap
ITERS = " ${ 1 :- 25 } "
: > " $OUT "
echo "# Local repro - $ITERS iters - $( date -u + '%H:%M:%S' ) UTC" >> " $OUT "
for n in $( seq 1 " $ITERS " ) ; do
rm -f tests/openwisp-monitoring-tests.db* tests/openwisp-monitoring.db* 2 >/dev/null
log = $( timeout 600 python tests/manage.py test " $TARGET " --noinput 2 > & 1 )
c = $( printf '%s' " $log " | grep -ciE " $SIG " )
ok = $( printf '%s' " $log " | grep -cE '^OK$' )
echo "- iter $n : crash= $c ok= $ok $( date -u +%H:%M:%S ) " >> " $OUT "
if [ " $c " -gt 0 ] ; then
echo " >>> CRASH at iter $n <<<" >> " $OUT "
printf '%s' " $log " | grep -iE " $SIG " | sed -E 's/^/ /' >> " $OUT "
break
fi
done
echo "DONE $( date -u + '%H:%M:%S' ) UTC" >> " $OUT "
The advantage was not creativity. It was that the machine could keep doing the boring part while I supervised the direction of the work. After a few hundred runs we finally had a real, measurable signal: the crash reproduced about 7% of the time, and only on Python 3.13 . Every Python 3
[truncated]
The cause was hiding in GeoDjango. Django's SpatiaLite backend calls ctypes.util.find_library("spatialite") while building the SpatiaLite library paths for new connections. On Linux, find_library may fork an ldconfig subprocess. Doing that repeatedly in a multi-threaded live server is a bad combination: one thread can be forking while another thread is in the middle of SQLite connection cleanup, deep enough in C code that a bad interleaving shows up as heap corruption rather than a friendly Python traceback. Most of the time you get away with it. Python 3.13 changed the timing just enough to turn "most of the time" into an intermittent double free. The fix is small and a little boring, which is the best kind. In our SpatiaLite backend wrapper, we memoize the library lookup so the subprocess is not forked over and over. In the Selenium test mixin, we serialize connection open and close around Django's BaseDatabaseWrapper.connect and BaseDatabaseWrapper._close when the default backend is SQLite or SpatiaLite. That makes the connection lifecycle process-wide and boringly serialized for the duration of the live-server tests. With both pieces in place, the crash rate went from 7% to zero across more than ninety runs , and then held green across repeated CI reruns of the shared fix. The best part: it is shared
