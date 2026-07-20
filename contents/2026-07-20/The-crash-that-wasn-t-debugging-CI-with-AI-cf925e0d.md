---
source: "https://abbyssoul.github.io/engineering/2026/07/19/the-crash-that-wasnt-debugging-kinjo-ci-with-ai.html"
hn_url: "https://news.ycombinator.com/item?id=48973271"
title: "The crash that wasn't: debugging CI with AI"
article_title: "The crash that wasn’t: debugging Kinjo’s CI with AI | Abyss of Software Engineering"
author: "sagevoid"
captured_at: "2026-07-20T02:01:01Z"
capture_tool: "hn-digest"
hn_id: 48973271
score: 2
comments: 0
posted_at: "2026-07-20T01:01:23Z"
tags:
  - hacker-news
  - translated
---

# The crash that wasn't: debugging CI with AI

- HN: [48973271](https://news.ycombinator.com/item?id=48973271)
- Source: [abbyssoul.github.io](https://abbyssoul.github.io/engineering/2026/07/19/the-crash-that-wasnt-debugging-kinjo-ci-with-ai.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-20T01:01:23Z

## Translation

タイトル: クラッシュではなかった: AI を使用した CI のデバッグ
記事タイトル: クラッシュではなかった: AI を使用した金城の CI のデバッグ |ソフトウェアエンジニアリングの深淵
説明: 人間の直感と AI 支援の高速実験により、誤った Kinjo クラッシュが CI の tmux バージョンの違いに起因することがどのように追跡されたか。

記事本文:
そうではないクラッシュ: AI を使用した金城の CI のデバッグ |ソフトウェアエンジニアリングの深淵
コンテンツにスキップ
ソフトウェアエンジニアリングの深淵
動作するソフトウェアを構築するための実践的なメモ
執筆
トピックス
プロジェクト
について
エンジニアリングノート
そうでなかったクラッシュ: AI を使用した金城の CI のデバッグ
CI のみの失敗、人間の予感、AI 支援による高速実験、製品の一部となった診断作業。
CIの煙検査の結果、金城が墜落したことが判明した。実際、金城は無事退場した。
この区別には数時間かかり、数回の AI セッションと少数の廃棄が必要でした。
仮説と小さなクラッシュキャプチャツールキットを明らかにします。
Kinjo で依存関係を更新していましたが、
ターミナル UI からローカル DNS-SD および mDNS サービスを参照するための私のサイド プロジェクト。
その CI スイートには、tmux 内で Kinjo を起動し、いくつかのデータを送信するスモーク テストが含まれています。
キーストロークを確認し、正常に終了することを確認します。断続的に、代わりにテスト
報告した:
drive-tui: シグナルによってペインが終了しました (終了コードなし)
典型的なアプリケーションのクラッシュのように見えました。問題は 1 つだけありました。
ローカルで再現しないでください。 250 回以上の実行後ではなく、AddressSanitizer を使用していない場合でも、
タイミングを促進するためにマシンを単一の CPU コアに制限することはありません。
失敗。
これは CI にのみ存在する問題のデバッグに関する話ですが、
生産的な AI 支援エンジニアリングが実際にどのようなものであるかについて説明します。 AIがそれぞれ作った
仮説をテストするのが安価です。決定的な直観は得られなかった。ほとんど
重要なのは、誤ったクラッシュの調査に費やされた作業が消えたわけではありません。
次の本物を診断するためのはるかに優れた機械を金城に残しました。
AI により実験ループが高速化
AI を活用する上で最も有益な部分は、単一の素晴らしい答えではありませんでした。それは
推測テストのループの速度。
仮説を立てられるかもしれない、ch

テストハーネスを調整し、実験を進めます
CI、証拠を調べて次に進みます。 TUI の標準エラー出力を個別に取得しました。
通常の出口と信号死を区別するためにドライバーを変更しました。私たち
二分化された依存関係の変更。コア ダンプを有効にし、gdb を収集するようにしました
バックトレース、カーネル メッセージ、ジャーナル エントリ、およびアプリケーション ログ
失敗。
それらの実験のいくつかは、原因を見つけるのではなく、物事を排除しました。あ
依存性の疑いには責任がありませんでした。変更された他の依存関係を元に戻す
障害が発生したが解消されなかった頻度。これは、次の場合に役立つ手がかりでした。
それ自体: 私たちはおそらく微妙なタイミングまたはレイアウト効果を見ていたのですが、悪くはありませんでした。
特定のクレート内のコード パス。
これは、AI が本当に可能だと感じた点です。すべてのバリエーションを手作業で準備すると、
はるかに時間がかかりました。この機械によって調査作業が不要になったわけではありません。それ
アイデアとそのアイデアに関する証拠の間の機械的時間を短縮しました。
決定的な推測は人間によるものだった
何度試してもローカル マシンではエラーが再現しなかったので、最終的には疑問に思いました。
tmux サーバー自体が異なっていたかどうか。
それはランダムな実装の詳細ではありませんでした。煙テストには実際のテストが必要です
端末 UI を駆動するための疑似端末であり、tmux はその環境の一部です。もし
tmux の動作が異なると、テストの動作も異なります。たとえ Kinjo が
同じバイナリです。
バージョン チェックでは、まさにその違いが明らかになりました。
GitHub Actions は tmux 3.4 の Ubuntu パッケージを使用していました。
私の開発マシンは tmux 3.5a を使用していました。
その後、AI がこの疑惑を管理された実験に変えるのに役立ちました。同じ上に
Ubuntu 24.04 ランナーでは、各 tmux でナロー スモーク テストを 12 回ループしました。
バージョン:
同じランナー、同じアプリケーション、同じテスト、異なる tmux 。失敗の軌跡

d
tmux のバージョン。それはまた、何百もの地元での試みが私に教えてくれた理由の説明にもなりました。
何もありません: 私のローカル環境には、
失敗。
これは、AI 支援デバッグについて話すときに維持したいバランスです。
AI は仮説を検証し、実験を迅速に構築することに優れていました
彼らに挑戦するために必要でした。どちらの違いを選択するかについては、私が依然として関与していました
調査するために。役に立つ単位は「人間の直感」でも「AIの出力」でもなかった
隔離。それは二人が予感を結果に変えることができる速さでした。
死亡したプロセスは金城ではありませんでした
バージョンの相関関係を見つけることで、どこを確認すればよいかが説明されましたが、最も満足のいくものは
証拠の一部は、私たちが追加した前景を保存する小さなランチャーから来ました。
ハーネス。
通常、スモークテスト中のKinjoはtmux paneプロセスです。それがペインの運命を左右します。
アプリケーションの運命も同じように見えます。発射装置が彼らを分離した。それ
ペインプロセスになり、ターミナルフォアグラウンドで子供として金城を開始し、
そして両方のプロセスがどのように終了したかを記録しました。
実行が失敗すると、Kinjo の子がステータス 0 で終了したことが記録されました。
launcher (tmux ペイン プロセス) がシグナルを受信するプロセスでした。
つまり、金城は quit コマンドを処理し、正しくシャットダウンしました。ふー。
tmux 3.4 は、ティアダウン中にペイン プロセスに信号を送りました。元のテストでは、ここで
金城自身がその役割を果たし、信号はきれいな出口を駆け抜け、偽装した
アプリケーションのクラッシュとして。
私たちは正確な信号を特定したことはありませんでしたし、その時点ではその必要はありませんでした。それ
コア ダンプは生成されず、カーネル セグメンテーション違反やメモリ不足イベントも発生しませんでした。
さらに信号処理を追加すると、タイミングが変更されて失敗する可能性がありました。
消える。制御された tmux バージョンの実験

d 子の終了ステータスは
原因を特定するには十分です。
CI の修正は簡単でした。スモーク テスト ジョブとキャッシュ用に tmux 3.5a をビルドしました。
それ。クラッシュするはずだったテストは決定的でグリーンなものになりました。
これを、AI によってデバッグが楽になるという話に変えるのは簡単でしょう。
それは正直ではなく、単なる誇大広告になります。捜査は容易ではなかった。まだまだ大変な作業でした。
調査では膨大な数のトークンが消費され、数回にわたって行われました。
セッション。新しいセッションがそれぞれ継承されるように、引き継ぎ文書を維持しました。
前回の仮説、実験、分岐状態、および結論
1つ。まだ CI を待ち、あいまいな結果を解釈し、決定を下す必要がありました。
次に何を試そうか。
ただし、従来のデバッグにも問題があります。あらゆる実験は準備が必要で、
誰かによってレビューされ、実行され、解釈されます。商業プロジェクトでは、これらは次のとおりです
請求対象となるエンジニアリング時間。サイドプロジェクトでは、彼らは夜になり、注目を集めます。
エンジニアがすべての診断バリエーションを手動で構築すると、簡単にコストがかかる可能性があります
特にトークンでは再現できない障害の場合、トークンよりも効果的です。
自分たちのマシン。
だからといって、トークンのコストを無視すべきというわけではありません。それは、比較する必要があることを意味します
正直に言ってください。価値は調査が無料であることではありませんでした。その価値とは、
AI のおかげで、限られた条件下で、はるかに広範囲かつ高速な一連の実験を実行できるようになりました。
私に与えられた人間の時間。自分の時間。
診断作業が製品の一部になりました
最良の結果は、単に CI が再び緑色になることではありません。
この失敗を見つけるために、金城氏は再利用可能なクラッシュ キャプチャ ツールキット、ランチャーを入手しました。
アプリケーションのターミナルラッパーからの終了を区別する方法です。
コア ダンプ、自動 GDB バックトレース、stderr キャプチャ、および
関連情報を収集する障害レポート

アリシステムの証拠を隠さずに
元のテスト結果。
これらはいずれも tmux 3.4 に固有のものではありません。同じメカニズムをオンにすることができます。
将来の CI フレーク、およびこのアイデアは、金城氏が開発したときにユーザー向けのバグ レポートをサポートできます。
自分がコントロールできない環境では失敗します。誰かに説明を求めるのではなく
記憶から「クラッシュした」、プロジェクトは多くのことに答える証拠を収集できる
もっと良い質問: 金城さんは信号を受信しましたか?正常に終了しましたか？ありましたか？
コアダンプ？標準エラーには何が含まれていましたか?オペレーティング システムは何を観察しましたか?
最初の失敗は、アプリケーションに対する虚偽の告発でした。それを追いかけて
これにより、実際の障害が発生したときにアプリケーションを診断しやすくなりました。
バグハントは修正以上のものを生み出す可能性があります
CI のみの障害は、再現が難しいと言われることがよくあります。この場合はそうでした
関連する差異が存在しないため、ローカルで再現することは不可能です。
金城全然。それはその周りのtmuxサーバーにありました。
AI は、仮説から仮説に至るまでの道のりを短縮することで、この調査を実用的なものにしました。
実験。人間の文脈と直観が依然として疑問を打ち破った
ケースが開いた状態。このプロセスには、従来のプロセスと同様に、リアルタイムで実際のトークンがかかります。
バグハントには実際のエンジニアリング時間がかかります。それを価値のあるものにしたのは、次の両方です。
学習のスピードとその後に何が残ったか。
金城選手は落車はなかった。クラッシュレポートメカニズムが改善されました。
ランタイム構造とユーザーに基づいた、スクリプト、対話型アプリ、サービス、ジョブの実用的な分類。
(スタートアップで) テストする時間がありませんか?
自動テストが顧客価値の提供の一部であり、将来の変更を手頃な価格で維持するためのセーフティ ネットである理由。
これが分散システムが役立つ理由です (私も分散システムを構築しています)
問題が時間、メモリ、マシン全体にわたってスケールする必要がある場合、分散システムにはその複雑さの価値がある理由

NES、または失敗。
ソフトウェア エンジニアは、信頼性、分散システム、テスト、実際のプロジェクトに隠された教訓について執筆しています。
質問、修正、さまざまな視点は大歓迎です。
ソフトウェアの信頼性、分散システム、テスト、C++、AI 支援エンジニアリングに関する Ivan Ryabov による実践的なエッセイ。

## Original Extract

How human intuition and fast AI-assisted experiments traced a false Kinjo crash to a tmux version difference in CI.

The crash that wasn’t: debugging Kinjo’s CI with AI | Abyss of Software Engineering
Skip to content
Abyss of Software Engineering
Practical notes on building software that has to work
Writing
Topics
Projects
About
Engineering notes
The crash that wasn't: debugging Kinjo's CI with AI
A CI-only failure, a human hunch, fast AI-assisted experiments, and diagnostic work that became part of the product.
A CI smoke test told me Kinjo had crashed. Kinjo had actually exited successfully.
That distinction took a few hours, several AI sessions, a handful of discarded
hypotheses, and a small crash-capture toolkit to uncover.
I was updating dependencies in Kinjo ,
my side project for browsing local DNS-SD and mDNS services from a terminal UI.
Its CI suite includes a smoke test that starts Kinjo inside tmux, sends it some
keystrokes, and checks that it quits cleanly. Intermittently, the test instead
reported:
drive-tui: pane exited via a signal (no exit code)
It looked like a classic application crash. There was only one problem: I could
not reproduce it locally. Not after more than 250 runs, not under AddressSanitizer,
and not while restricting the machine to a single CPU core to encourage a timing
failure.
This is a story about debugging a problem that only existed in CI, but it is also
about what productive AI-assisted engineering actually looks like. AI made each
hypothesis cheap to test. It did not supply the decisive intuition. Most
importantly, the work spent investigating the false crash did not disappear: it
left Kinjo with much better machinery for diagnosing the next real one.
AI made the experimental loop fast
The most useful part of working with AI was not a single brilliant answer. It was
the speed of the guess-test loop.
We could form a hypothesis, change the test harness, push an experiment through
CI, inspect the evidence, and move on. We captured the TUI’s stderr separately.
We changed the driver to distinguish a normal exit from a signal death. We
bisected dependency changes. We enabled core dumps and arranged to collect gdb
backtraces, kernel messages, journal entries, and application logs after a
failure.
Several of those experiments ruled things out rather than finding the cause. A
suspected dependency was not responsible. Reverting other dependencies changed
how often the failure appeared but did not remove it, which was a useful clue in
itself: we were probably looking at a delicate timing or layout effect, not a bad
code path in one particular crate.
This is where AI felt genuinely enabling. Preparing every variation by hand would
have taken much longer. The machine did not eliminate the investigative work; it
compressed the mechanical time between an idea and evidence about that idea.
The decisive guess was a human one
Because the failure would not reproduce on my local machine, despite many attempts, I eventually wondered
whether the tmux server itself was different.
That was not a random implementation detail. The smoke test needs a real
pseudo-terminal to drive a terminal UI, and tmux is part of that environment. If
tmux behaves differently, then the test behaves differently — even when Kinjo is
the same binary.
The version check exposed exactly that difference:
GitHub Actions was using the Ubuntu package of tmux 3.4.
My development machine was using tmux 3.5a.
AI then helped turn the suspicion into a controlled experiment. On the same
Ubuntu 24.04 runner, we looped a narrow smoke test twelve times with each tmux
version:
Same runner, same application, same test, different tmux . The failure tracked
the tmux version. That also explained why hundreds of local attempts had told me
nothing: my local environment did not contain the component that produced the
failure.
This is the balance I want to preserve when talking about AI-assisted debugging.
AI was excellent at validating hypotheses and rapidly building the experiments
needed to challenge them. I was still instrumental in choosing which difference
to investigate. The useful unit was neither “human intuition” nor “AI output” in
isolation; it was the speed with which the two could turn a hunch into a result.
The process that died was not Kinjo
Finding the version correlation explained where to look, but the most satisfying
piece of evidence came from a small foreground-preserving launcher we added to
the harness.
Normally, Kinjo under a smoke test is the tmux pane process. That makes the pane’s fate and the
application’s fate look like the same thing. The launcher separated them. It
became the pane process, started Kinjo as a child in the terminal foreground,
and recorded how both processes ended.
On a failing run it recorded that the Kinjo child exited with status 0. The
launcher—the tmux pane process—was the process that received a signal.
In other words, Kinjo had handled the quit command and shut down correctly! Phew.
tmux 3.4 then signalled the pane process during teardown. In the original test, where
Kinjo itself occupied that role, the signal raced its clean exit and masqueraded
as an application crash.
We never pinned down the exact signal, and at that point we did not need to. It
did not produce a core dump, there was no kernel segfault or out-of-memory event,
and adding more signal handling changed the timing enough to make the failure
disappear. The controlled tmux-version experiment and the child exit status were
enough to establish the cause.
The CI fix was straightforward: build tmux 3.5a for the smoke-test job and cache
it. The supposedly crashing test became deterministic and green.
It would be easy to turn this into a story about AI making debugging effortless.
That would not be honest, and just hype. The investigation was not effortless. It was still a lot of work.
The investigation consumed a non-trivial number of tokens and spanned several
sessions. We maintained a handover document so that each new session inherited
the hypotheses, experiments, branch state, and conclusions from the previous
one. There was still waiting for CI, interpreting ambiguous results, and deciding
what to try next.
But conventional debugging has a bill too. Every experiment has to be prepared,
reviewed, run, and interpreted by somebody. On a commercial project those are
billable engineering hours; on a side project they are evenings and attention.
An engineer manually building every diagnostic variation could easily have cost
more than the tokens, especially for a failure they could never reproduce on
their own machine.
That does not mean token cost should be ignored. It means the comparison should
be honest. The value was not that the investigation was free. The value was that
AI let me run a much wider and faster sequence of experiments with the limited
human time I had available. My own time.
The diagnostic work became part of the product
The best outcome is not merely that CI is green again.
To find this failure, Kinjo gained a reusable crash-capture toolkit: a launcher
that distinguishes the application’s exit from its terminal wrapper, a way to
enable and collect core dumps, automatic gdb backtraces, stderr capture, and a
failure report that gathers the relevant system evidence without hiding the
original test result.
None of that is specific to tmux 3.4. The same mechanism can be switched on for a
future CI flake, and the ideas can support user-facing bug reports when Kinjo
fails in an environment I do not control. Instead of asking someone to describe
“it crashed” from memory, the project can collect evidence that answers much
better questions: Did Kinjo receive a signal? Did it exit normally? Was there a
core dump? What did stderr contain? What did the operating system observe?
The original failure was a false accusation against the application. Chasing it
made the application easier to diagnose when a real failure arrives.
A bug hunt can produce more than a fix
A CI-only failure is often described as hard to reproduce. In this case it was
impossible to reproduce locally because the relevant difference was not in
Kinjo at all; it was in the tmux server around it.
AI made this investigation practical by shortening the path from hypothesis to
experiment. Human context and intuition still supplied the question that broke
the case open. The process cost real time and real tokens, just as a traditional
bug hunt costs real engineering hours. What made it worthwhile was both the
speed of learning and what remained afterwards.
Kinjo did not have a crash. It now has a better crash-reporting mechanism.
A practical taxonomy of scripts, interactive apps, services, and jobs based on runtime structure and users.
No time for testing (in a start-up)?
Why automated testing is part of delivering customer value—and the safety net that keeps future changes affordable.
This is why distributed systems are useful (and I am building one)
Why distributed systems are worth their complexity when problems must scale across time, memory, machines, or failures.
Software engineer writing about reliability, distributed systems, testing, and the lessons hidden inside real projects.
Questions, corrections, and different perspectives are welcome.
Practical essays on software reliability, distributed systems, testing, C++, and AI-assisted engineering by Ivan Ryabov.
