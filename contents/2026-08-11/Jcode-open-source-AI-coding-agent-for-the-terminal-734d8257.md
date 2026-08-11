---
source: "https://jcode.sh/"
hn_url: "https://news.ycombinator.com/item?id=49255865"
title: "Jcode – open-source AI coding agent for the terminal"
article_title: "jcode - open-source AI coding agent for the terminal"
author: "jonbaer"
captured_at: "2026-08-11T10:42:21Z"
capture_tool: "hn-digest"
hn_id: 49255865
score: 1
comments: 0
posted_at: "2026-08-11T10:20:23Z"
tags:
  - hacker-news
  - translated
---

# Jcode – open-source AI coding agent for the terminal

- HN: [49255865](https://news.ycombinator.com/item?id=49255865)
- Source: [jcode.sh](https://jcode.sh/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T10:20:23Z

## Translation

タイトル: Jcode – 端末用のオープンソース AI コーディング エージェント
記事のタイトル: jcode - 端末用のオープンソース AI コーディング エージェント
説明: jcode は、Rust で書かれたターミナル用のオープンソース AI コーディング エージェントです。耐久性のあるメモリ、バックグラウンド タスク、エージェント群を 1 つのコマンドでインストールできます。

記事本文:
jcode - ターミナル用のオープンソース AI コーディング エージェント
jコード
Rust で書かれたオープンソースのターミナル コーディング エージェント。
カール -fsSL https://jcode.sh/install | bash コピー macOS Linux Windows すべてのビルド · ドキュメント · 価格 · GitHub
並列処理はコーディングの生産性を左右する最大の要因です。コーディング中に決して待ってはいけません。エージェントが作業しているすべての瞬間は、次のタスクを配布している可能性がある瞬間です。セッションが終了するのを監視する代わりに、別のセッションを開始すると、スループットは配布できる限り多くのタスクになります。 Jcode を使用すると、数十のエージェントを並行して実行することができ、12 個のエージェントの出力は 1 個のエージェントの数十倍になります。
ハーネスはモデルと同じくらい重要です。同じモデルでも、周囲にあるもの、つまり、到達できるツール、認識するコンテキスト、保持する記憶、間違いを発見するフィードバック ループに応じて、非常に異なる結果が生成されます。モデルの進行状況は独自のスケジュールに従って到着します。活用の進歩は私たちが行うべきものであり、それはモデルのリリースごとに増大します。
開発ツールはオープンソースである必要があります。コーディング エージェントはコードを読み取り、ファイルを編集し、マシン上でコマンドを実行します。これほど広範囲に及ぶツールは検査可能であり、変更可能である必要があります。最高の開発ツールは常にユーザーが開いて変更できるものであり、エージェントのソースはカスタマイズできるものである必要があります。 Jcode は MIT ライセンスを取得しており、私たちが測定し学習したものはすべて公開されています。
大規模並列処理のボトルネックはリソース効率です。数十のエージェントと RAM 消費バルーンを生成しますが、これはほとんどのコンシューマ マシンでは実現できません。各エージェントが数百メガバイトを消費する場合、いくつかのエージェントを実行し、その背後で作業をキューに入れる必要があります。 Jcodeはそれを修正しました。私たちは徹底的に最適化しているので、別のエージェントを生成するかどうかは決定ではなく、数十のセッションを実行することは実際には可能です。

スピードを犠牲にしたものは何もありません。以下のすべての数値は、同じマシン上での実際のエージェントの実際の起動からサンプリングされています。
追加の比例メモリ (PSS) は、クライアントがすでに実行されている場合、追加されるクライアントごとに追加されます。 10 個の jcode セッションのコストは約 100 MB で、1 つのクロード コードの半分未満です。
~10.4 MB ベースライン クロード コード ~212.7 MB ~212.7 MB 20.5 倍の RAM 完全な比較を表示 完全な比較を非表示 追加セッションあたりのメモリ、すべてのツール
~9.9 MB ベースライン jcode ~10.4 MB ~10.4 MB 1.1 倍の RAM Codex CLI ~21.6 MB ~21.6 MB 2.2 倍の RAM pi ~76.5 MB ~76.5 MB 7.7 倍の RAM Antigravity CLI ~86.4 MB ~86.4 MB 8.7 倍の RAM カーソル エージェント ~157.5 MB ~157.5 MB 15.9 倍の RAM GitHub Copilot CLI ~158.1 MB ~158.1 MB 16.0 倍の RAM クロード コード ~212.7 MB ~212.7 MB 21.5 倍の RAM OpenCode ~318.4 MB ~318.4 MB 32.2 倍の RAM 最初の入力までの時間
実際に入力できるまでの時間: 入力したプローブ テキストがレンダリングされた画面に表示されるまでの時間。対話型 PTY がそれぞれ 10 回起動します。 Antigravity は、サインイン画面でプローブ エコーが抑制されるため、内部の入力準備完了ログ マーカーを使用します。
48.7 ミリ秒ベースライン反重力 CLI 383.7 ミリ秒 383.7 ミリ秒 7.9 倍遅い pi 596.4 ミリ秒 596.4 ミリ秒 12.2 倍遅い Codex CLI 905.8 ミリ秒 905.8 ミリ秒 18.6 倍遅い OpenCode 1047.9 ミリ秒 1047.9 ミリ秒 21.5 倍GitHub Copilot CLI 1583.4 ms 1583.4 ms 32.5 倍遅い Cursor Agent 1978.7 ms 1978.7 ms 40.6 倍遅い Claude Code 3512.8 ms 3512.8 ms 72.2 倍遅い 最初のフレームまでの時間
何かがレンダリングされ、同じ 10 のインタラクティブ PTY が起動するまでの時間。
14.0 ミリ秒ベースライン反重力 CLI 383.5 ミリ秒 383.5 ミリ秒 27.4 倍遅い pi 590.7 ミリ秒 590.7 ミリ秒 42.2 倍遅い Codex CLI 882.8 ミリ秒 882.8 ミリ秒 63.1 倍遅い OpenCode 1035.9 ミリ秒 1035.9 ミリ秒 74.0× GitHub Copilot CLI 1518.6 ミリ秒 1518.6 ミリ秒 108.5 倍遅い カーソル エージェント 1949.7 ミリ秒 1949.7 ミリ秒 139.3 倍遅い クロード コード 3436.9

ms 3436.9 ms 245.5 倍遅い パフォーマンス デモを見る jcode パフォーマンス デモを見る swarm セッションを見る 20 のエージェントを並行して実行する インテリジェンス
これらのハーネスに対して標準の eval を実行すると、エージェントが実際に実行できる内容がいかに不十分に表現されているかがわかりました。 Existing benchmarks face a forced choice.公開ベンチマークはそのソリューションをトレーニング コーパスに漏洩するため、スコアは暗記力を測定し、能力を誤って表現します。プライベートベンチマークはこれに抵抗しますが、隠されたテストセットに対する信頼を要求し、透明性が低く、アクセスするのが困難です。その選択以外にも、彼らには 2 つの構造的欠陥があります。タイムアウトが課されるため、エージェントは問題により多くの時間を割り当てるとペナルティを受けます。これはまさに、エージェントに開発してもらいたい長期的な行動です。そして、それらは粗く離散的な境界でスコアを付けます。2 つの非常に異なるモデル間のギャップが、単一のタスクの完了として現れることがあります。ベンチの 70% を解決するモデルは、はるかに高い能力レベルを要求する残りの 30% の隣に位置する可能性があり、そのデルタ全体は個別のタスク間では見えません。そこで、次の仕様を備えた新しいクラスのベンチマークを設計しました。
構造によって汚染されない: 暗記すべき答えはないので、すべてを公開できる
構造によって飽和しにくい: 最適化の深さに基づいてスコアが付けられるため、フロンティアを越えても天井は開いたままになります。
構造的に決定的: 同じ提出物、同じスコア、常に
構築によって定量化可能: 指標はタスクの定義から外れており、ルーブリックやジャッジもありません
連続的な構造: 合格/不合格ではなくアナログスコア軸を使用するため、能力の差がすべてのレベルのスコア差として表示され、時間は制限されるのではなく記録されます。
構造的にチート耐性があります: 開始点として優れたリファレンス実装が提供されているため、既製のソリューションを Web で検索できます。

ons は、より低い機能レベルでのみ不正行為を行っています。フロンティアでは、有用な組み換え以外のコピーすべき実装は残っていない
このデザインはすべてが公開されているため、すぐに信頼できます。以前のモデルのトランスクリプトでのトレーニングは少し役立つかもしれませんが、モデルがそこから真に一般化されていない限り、スコアは向上しません。トランスクリプトを再生すると古いスコアと一致するだけで、それを上回るには、記憶されたものを超えて最適化する必要があり、ベンチマークがそもそも測定する能力です。
jcode ベンチ v1 は、この仕様に準拠した 3 つのタスク (float-print、json-unescape、utf16-transcode) を提供します。これらのタスクを選択したのは、これらのタスクが数秒で評価され、便利で最適化の真のメリットとなる現実世界の関数 (実際にライブラリに出荷される作業で、エージェントが登ることができるタイトなループ) を純粋にコーディングしているためです。実装がテスト スイートにオーバーフィットできないようにタスクも選択されます。正確性はサンプルではなく入力空間全体で検証されます。 Linux リポジトリ上でタイミングを計って grep の速度の測定を比較します。この場合、実装は特に Linux リポジトリにオーバーフィットする可能性があります。上では、単一のユーザー プロンプトからの、時間の経過に伴うエージェントのスコアをプロットしています。これらのタスクではツールの呼び出し時間は無視できるため、時間はコストおよびトークンの使用率と 1 対 1 の相関があることに注意してください。どちらの実行も、高度な思考を備えた同一のプロンプトを備えた Opus 4.8 を使用し、同時に起動され、それぞれが独自の判断で停止しました。フロートプリントでは、どちらも完全な 23² 正しさのゲートを通過しました。 jcode は +8.64 (398 倍の高速化) で終了しましたが、Claude Code は +7.17 (144 倍) でした。
汚染されないベンチマーク設計・jcode ベンチ
ハーネスに依存しないフロンティア モデルの結果、完全な結果とエージェントの記録
DeepSWE v1.1 ハーネスの比較
ハーネスの比較は完全に

一致: 同じ 113 個の DeepSWE v1.1 タスク、GPT-5.6 Sol、xhigh 推論労力、時間制限なしの自然完了、および k=1。ハーネスだけが変わります。現在の結果は完全に同点です。相互にスコア付けされた 112 個のタスクでは、Jcode が勝った 11 個のタスクで Codex が失敗し、Codex が勝った 10 個のタスクで Jcode が失敗しました。
すべてのタスク、両方の実行から監査されたタスクごとの結果
DataCurve DeepSWE v1.1 リファレンス
DataCurve は、すべてのミニ SWE エージェントのロールアウトも公開します。公式の GPT-5.6 Sol 結果は、改訂された v1.1 グレーダーと 4 つのベンチマーク全体の実行を使用しているため、有用なコンテキストですが、上記の v1 k=1 ハーネス結果に対してランク付けされていません。
公式の 95% の実行間隔: 最大 ±2.8 ポイント。 ±1.4ポイント高い。 DataCurve リーダーボード · 派生したタスクごとのデータと来歴
89 のターミナル タスク、claude-opus-4-8 がモーダルのハーバー ハーネスを介して実行されます。私たちの最良のセルは、Claude Code が公表している 78.9% に対して、中程度の労力で 77.8% (k=2) です。 Medium は xhigh よりも優れていますが、トライアルあたりのコストが低く、タスクが 30% 早く終了するため、このベンチマークで Medium を超えると、より多くの考えが元をとらなくなります。また、失敗した xhigh トライアル 28 件を読み取って分類しました。内訳は、実際のミスが 22 件、インフラ フレークが 3 件、境界線の閾値が 3 件、採点者の論争が 0 件です。
すべての実行で、コストとタスクごとの記録に対する精度が向上します
Jcode の todo ツールは、項目が割り当てられたときと完了とマークされたときの両方で、エージェントに各タスク項目に対する信頼度を評価するよう求めます。ターミナルベンチの実行全体でこれらのスコアを追跡すると、明確なパターンがわかりました。信頼度は割り当て後は常に高くなりますが、割り当て前に低くなることもあり、エージェントが割り当て時に自信を持っていたタスクから失敗することはほとんどありません。割り当て時の低いスコアは実際のシグナルですが、最後に 100 点にジャンプするのはそうではありません。そのため、信頼度が大幅に上昇した場合は、エージェントを受け入れるのではなく、元に戻って作業を確認するように強制します。

主張をする。理想的には、実装中に検証が行われるにつれて信頼性が段階的に高まり、テストに合格するたびにステップアップが得られます。上図は、ターミナルベンチ 2.1 上の両方のハーネスです。各行は、割り当て時の確信度から完了時の確信度までの 1 つのトライアルです。追加のチェックは有益です。時間内に終了するトライアルはより頻繁に通過し (92% 対 88%)、ベンチマークの 15 分の制限によって打ち切られたトライアルでも、すでに正しい作業が含まれている可能性が高くなります (47% 対 42%)。
完全な信頼度調査・すべての治験記録
モデルはこの作品がどの程度の坂を登れると考えましたか?
n = 2,012 ゴールスコアの提出
エージェントは、山を登る指標を持っているときに最大限の能力を発揮します。これは、強化学習によって報酬シグナルに対して最適化するようにモデルがトレーニングされるため、測定可能なフィードバックによってエージェントが長期にわたって改善を続けることができ、追加の生産的な反復がより良い結果につながるためであると考えられます。ほとんどのユーザーはこれを最大限に活用していません。 Jcode は、該当する場合、ユーザーに対してこれを実行します。すべてのエージェントの目標は、実際の進捗状況の定量化と反復可能性に基づいて、0 から 100 までの登坂可能性の評価を受けます。達成するための明確な目標がなければ、高いスコアは信頼できません。目標のスコアが低い場合、ハーネスは押し戻します。つまり、目標を検証可能な目標に再構成し、それを測定するハーネスを構築します。これがなければ、オープンエンド タスクではエージェントは反復する対象が何もなくなり、次の試行を前回よりも改善するために使用できるシグナルもなくなります。
登坂可能性の実装を確認する · 集計データをダウンロードする
永続性はインテリジェンスの残りの半分です。エージェントの失敗のほとんどは、間違った答えではなく、早期終了です。モデルは勝利を宣言するのが大好きです。 Jcode は最初に Todo リストをチェックします。 Todo が完了していない状態でターンが終了すると、ハーネスは

s は、モデルを自動的に動作できるように戻します。 Poke は失敗についても賢明です。一時的なネットワーク エラーは再試行されますが、再試行不可能なエラーはトークンを書き込む代わりにループを停止します。同じメカニズムがヘッドレス実行を駆動するため、「jcode run」は作業が実際に終了するまでターンをまたいでタスクを繰り返し続けます。
永続化の実装を参照する · 完全な実行トランスクリプトを読む
同じ規律がユーザーとモデルの間のすべてに適用されます。すべてのトークン、すべてのキャッシュ ヒット、すべてのバックグラウンド プロセスは、モデルがオーバーヘッドではなくユーザーの問題にその容量を費やすように設計されています。
追加専用のコンテキスト エンジニアリングにより、モデルの KV キャッシュがターンごとにホットに維持されます。
レイテンシとコストに関する最大の要因はプロンプト キャッシュです。プロバイダーは、モデルがすでに確認したトークンに対して一部の料金を請求します。また、トークンを読み戻すほうが、アテンションを最初から再計算するよりもはるかに高速です。 Jcode はキャッシュがほとんど壊れないように構築されています。プロンプト プレフィックスは安定しており、会話は厳密に追加専用であるため、可能な限り長く共有されたプレフィックスは、サイレントに無効化されるのではなく、ターン間で存続します。通常キャッシュを破壊する可能性のあるものはプレフィックスから除外されます。ツール スキーマは固定ディスク上のキャッシュから取得され、MCP ツールは広告です。

[切り捨てられた]

## Original Extract

jcode is an open-source AI coding agent for the terminal, written in Rust. Durable memory, background tasks, and agent swarms, installed with one command.

jcode - open-source AI coding agent for the terminal
jcode
An open source terminal coding agent, written in Rust.
curl -fsSL https://jcode.sh/install | bash copy macOS Linux Windows All builds · Docs · Pricing · GitHub
Parallelism is the biggest lever on coding productivity. You should never wait while coding. Every moment an agent is working is a moment you could be handing out the next task: spin up another session instead of watching this one finish, and your throughput becomes as many tasks as you can give out. Jcode makes it possible to run dozens of agents in parallel, and a dozen agents is a dozen times the output of one.
The harness matters as much as the model. The same model produces very different results depending on what surrounds it: the tools it can reach, the context it sees, the memory it keeps, and the feedback loops that catch its mistakes. Model progress arrives on its own schedule. Harness progress is ours to make, and it compounds with every model release.
Dev tools must be open source. A coding agent reads your code, edits your files, and runs commands on your machine. A tool with that much reach has to be inspectable, and it has to be modifiable: the best dev tools have always been the ones their users could open up and change, and your agent's source should be yours to customize. Jcode is MIT licensed, and everything we measure and learn is published.
The bottleneck to massive parallelism is resource efficiency. Spawn dozens of agents and RAM consumption balloons, unfeasible on most consumer machines: when each agent takes hundreds of megabytes, you run a handful and queue your work behind them. Jcode fixed that. We optimize to the bone so spawning another agent is a non-decision and running dozens of sessions is actually possible, with none of it traded for speed. Every number below is sampled from real launches of real agents on the same machine.
Extra proportional memory (PSS) each additional client adds once one is already running. Ten jcode sessions cost about 100 MB, less than half of one Claude Code.
~10.4 MB baseline Claude Code ~212.7 MB ~212.7 MB 20.5× more RAM Show the full comparison Hide the full comparison Memory per additional session, all tools
~9.9 MB baseline jcode ~10.4 MB ~10.4 MB 1.1× more RAM Codex CLI ~21.6 MB ~21.6 MB 2.2× more RAM pi ~76.5 MB ~76.5 MB 7.7× more RAM Antigravity CLI ~86.4 MB ~86.4 MB 8.7× more RAM Cursor Agent ~157.5 MB ~157.5 MB 15.9× more RAM GitHub Copilot CLI ~158.1 MB ~158.1 MB 16.0× more RAM Claude Code ~212.7 MB ~212.7 MB 21.5× more RAM OpenCode ~318.4 MB ~318.4 MB 32.2× more RAM Time to first input
How long until you can actually type: time until typed probe text appears on the rendered screen, 10 interactive PTY launches each. Antigravity uses its internal input-ready log marker because its sign-in screen suppresses probe echo.
48.7 ms baseline Antigravity CLI 383.7 ms 383.7 ms 7.9× slower pi 596.4 ms 596.4 ms 12.2× slower Codex CLI 905.8 ms 905.8 ms 18.6× slower OpenCode 1047.9 ms 1047.9 ms 21.5× slower GitHub Copilot CLI 1583.4 ms 1583.4 ms 32.5× slower Cursor Agent 1978.7 ms 1978.7 ms 40.6× slower Claude Code 3512.8 ms 3512.8 ms 72.2× slower Time to first frame
How long until anything renders, same 10 interactive PTY launches.
14.0 ms baseline Antigravity CLI 383.5 ms 383.5 ms 27.4× slower pi 590.7 ms 590.7 ms 42.2× slower Codex CLI 882.8 ms 882.8 ms 63.1× slower OpenCode 1035.9 ms 1035.9 ms 74.0× slower GitHub Copilot CLI 1518.6 ms 1518.6 ms 108.5× slower Cursor Agent 1949.7 ms 1949.7 ms 139.3× slower Claude Code 3436.9 ms 3436.9 ms 245.5× slower Watch the performance demo jcode performance demonstration Watch the swarm session 20 agents in parallel Intelligence
Running the standard evals against these harnesses showed us how poorly they represent what agents can actually do. Existing benchmarks face a forced choice. Public benchmarks leak their solutions into training corpora, so scores measure memorization and misrepresent capability. Private benchmarks resist that, but they demand trust in a hidden test set, offer less transparency, and are hard to get access to. Beyond that choice, they share two structural flaws. They impose timeouts, so an agent is penalized for allocating more time to a problem, exactly the long-horizon behavior we want agents to develop. And they score on a coarse, discrete boundary: the gap between two very different models can show up as a single task completion. A model that solves 70% of a bench may sit next to a remaining 30% that demands a far higher capability level, and that entire delta is invisible between the discrete tasks. So we designed a new class of benchmarks with a spec:
uncontaminatable by construction: there is no answer to memorize, so everything can be public
hard to saturate by construction: scored on optimization depth, so the ceiling stays open past the frontier
deterministic by construction: same submission, same score, always
quantifiable by construction: the metric falls out of the task definition, no rubric, no judge
continuous by construction: an analog score axis instead of pass/fail, so capability differences show up as score differences at every level, and time is recorded rather than capped
cheat-resistant by construction: we provide a good reference implementation as the starting point, so web searching for ready-made solutions is only cheating at lower capability levels. At the frontier, there are no implementations left to copy that aren't useful recombinations
This design is immediately trustable because everything is public. Training on a previous model's transcript may even help a little, but it doesn't score better unless the model has genuinely generalized from it: replaying a transcript only matches the old score, and beating it requires optimizing past what was memorized, which is the capability the benchmark measures in the first place.
jcode bench v1 provides three tasks that follow this spec: float-print, json-unescape, and utf16-transcode. We chose these tasks because they grade in seconds and are pure coding on real world functions that are useful and genuinely benefit from optimization: a tight loop the agent can climb, on work that actually ships in libraries. The tasks are also chosen so the implementation cannot be overfit to the test suite: correctness is verified over the entire input space, not a sample. Compare measuring grep's speed by timing it on the Linux repo, where an implementation can be overfit to the Linux repo specifically. Above, we plot the score of the agent over time, from a single user prompt. Note that time is 1:1 correlated with cost and token utilization, because tool call time is negligible in these tasks. Both runs used Opus 4.8 with high thinking, identical prompts, launched concurrently, each stopping on its own judgment. On float-print both passed the full 2³² correctness gate; jcode finished at +8.64 (398x speedup) versus Claude Code's +7.17 (144x).
The uncontaminatable benchmark design · jcode bench
harness-agnostic frontier model results · full results and agent transcripts
DeepSWE v1.1 harness comparison
The harness comparison is fully matched: the same 113 DeepSWE v1.1 tasks, GPT-5.6 Sol, xhigh reasoning effort, untimed natural completion, and k=1. Only the harness changes. The current result is an exact tie: on the 112 mutually scored tasks Jcode won 11 tasks that Codex failed and Codex won 10 that Jcode failed.
Every task, with audited per-task outcomes from both runs
DataCurve DeepSWE v1.1 reference
DataCurve also publishes every mini-SWE-agent rollout. Their official GPT-5.6 Sol results use the revised v1.1 grader and four whole-benchmark runs, so they are useful context but are not ranked against the v1 k=1 harness results above.
Official 95% run-to-run intervals: max ±2.8 points; high ±1.4 points. DataCurve leaderboard · derived per-task data and provenance
89 terminal tasks, claude-opus-4-8, run through the Harbor harness on Modal. Our best cell is medium effort at 77.8% (k=2), against Claude Code's published 78.9%. Medium beats xhigh while costing less per trial and finishing tasks 30% sooner, so more thinking stops paying for itself past medium on this benchmark. We also read and classified 28 failing xhigh trials: 22 real misses, 3 infra flakes, 3 borderline thresholds, and 0 grader disputes.
Every run, with accuracy against cost and per-task transcripts
Jcode's todo tool asks the agent to rate its confidence in each task item, both when the item is assigned and when it's marked done. Tracing those scores across our Terminal-Bench runs showed a clear pattern: confidence is always high after, but sometimes low before, and failures rarely come from tasks the agent was confident about at assignment. The low scores at assignment are real signal, and the jump to 100 at the end is not. So when we see a large spike in confidence, we force the agent to go back and check its work instead of accepting the claim. Ideally, confidence rises incrementally as validation happens during implementation, with each passing test earning a step up. Above, both harnesses on Terminal-Bench 2.1: each line is one trial, from confidence at assignment to confidence at completion. The extra checking pays: trials that finish in time pass more often (92% vs 88%), and even trials cut off by the benchmark's 15 minute limit are more likely to already contain correct work (47% vs 42%).
the full confidence study · every trial transcript
How hill-climbable did the model think the work was?
n = 2,012 goal-score submissions
Agents are at their most capable when they have a metric to hill-climb. This is likely because reinforcement learning trains models to optimize against reward signals, so measurable feedback helps agents keep improving over longer time horizons, and those additional productive iterations lead to better results. Most users never take full advantage of this. Jcode does this for the user when applicable. Every agent goal receives a hill-climbability rating from 0 to 100, based on how quantifiable and iterable its progress really is. A high score is not credible without a stated objective to climb toward. When a goal scores low, the harness pushes back: reframe the goal into a verifiable objective and build the harness that measures it. Without this, an open-ended task leaves the agent with nothing to iterate against, no signal it can use to make its next attempt better than its last.
See the hill-climbability implementation · Download the aggregate data
Persistence is the other half of intelligence: most agent failures are not wrong answers but early exits. Models love to declare victory. Jcode checks the todo list first. When a turn ends with incomplete todos, the harness pokes the model back to work automatically. Pokes are smart about failure too: transient network errors are retried, while non-retryable errors stop the loop instead of burning tokens. The same mechanism drives headless runs, so `jcode run` keeps iterating on a task across turns until the work is actually finished.
See the persistence implementation · Read the full run transcript
The same discipline is applied to everything between you and the model: every token, every cache hit, and every background process is engineered so the model spends its capacity on your problem instead of on overhead.
Append-only context engineering keeps the model's KV cache hot turn after turn.
The single biggest lever on latency and cost is the prompt cache: providers charge a fraction for tokens the model has already seen, and reading them back is far faster than recomputing attention from scratch. Jcode is built so that cache almost never breaks. The prompt prefix is stable and the conversation is strictly append-only, so the longest possible shared prefix survives between turns instead of being silently invalidated. Things that would normally bust the cache are kept out of the prefix: tool schemas come from a fixed on-disk cache, MCP tools are ad

[truncated]
