---
source: "https://faridsaid.com/en/blog/event-loop-starvation-trading.html"
hn_url: "https://news.ycombinator.com/item?id=48616819"
title: "I built a real-time trading platform with an AI; then the clock started lying"
article_title: "I built a real-time trading platform with an AI. Then the clock started lying. - Farid Saïd"
author: "fawraw"
captured_at: "2026-06-21T10:18:34Z"
capture_tool: "hn-digest"
hn_id: 48616819
score: 1
comments: 0
posted_at: "2026-06-21T08:24:49Z"
tags:
  - hacker-news
  - translated
---

# I built a real-time trading platform with an AI; then the clock started lying

- HN: [48616819](https://news.ycombinator.com/item?id=48616819)
- Source: [faridsaid.com](https://faridsaid.com/en/blog/event-loop-starvation-trading.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T08:24:49Z

## Translation

タイトル: AI を使用してリアルタイム取引プラットフォームを構築しました。それから時計は嘘をつき始めた
記事タイトル：AIを使ったリアルタイム取引プラットフォームを構築してみました。それから時計は嘘をつき始めました。 - ファリド・サイード
説明: リアルタイム取引プラットフォームでの非同期イベント ループの枯渇を診断した方法: 1 つの遅いクライアントがループ全体を停止させました。デバッグの話と 2 つの修正。

記事本文:
FS
AIを使ったリアルタイム取引プラットフォームを構築してみました
ブログに戻る
AIを使ったリアルタイム取引プラットフォームを構築しました。それから時計は嘘をつき始めました。
前回の記事では、AI に本番インフラストラクチャへの読み取り専用 SSH アクセスを許可し、監査、文書化、監視することについて説明しました。論理的な次のステップは、単に読み取るだけでなく、ビルドできるようにすることでした。
それで私はそうしました。クロードを共同パイロットとして、私は実際のアプリケーション、つまり OTC 金利スワップの価格マッチング プラットフォームを構築しました。リアルタイム、WebSocket、SSO 認証、複数の取引相手の相互利益を自動的に組み合わせるエンジン。 FastAPI、asyncio、SQLite。パイロットブローカーと約 15 の銀行取引相手とともに本番環境に移行しました。
するとブローカーは私に「まだ時間が残っているうちにセッションは打ち切られた」とメッセージを送ってきた。また、「トレーダーの概念は点滅し、現れたり消えたりします。」そして、「タイミングがおかしい。セッションが長すぎる。」
私の時計は嘘をつき始めていた。ここでは、私がその理由を調べた方法と、「非同期」という言葉について何を学んだかを説明します。
症状: セッションが時間通りに終了しない
プラットフォームの中心となるのは、固定期間のマッチング セッションです。たとえば 10 分間に設定すると、カウントダウンが始まり、一致するウィンドウが表示され、終了します。誰もが同じ時計を見ます。
いいえを除いて。運用環境では、600 秒に構成されたセッションは 1202 (倍) 持続しました。もう一つ、1803（トリプル）。約 40 回のセッションを遡ってみましたが、明確な 2 の因数は見つかりませんでした。これを見つけました:
10 分のセッションが 1 時間近く続くこともあります。そして誰もその理由を理解できませんでした、まず私がそうでした。
最初の間違った方向：「タイマーです」
明らかな反射です。タイマー コードにはバグがあります。双腕になったり、止まり忘れたり、何か。
それを単独で分離してテストしました。 2秒

カウントダウン、3 秒間のマッチング: 正確、単一トランジション、ダブルアームなし。タイマーは単独で完璧でした。
さらに悪いことに（あるいはそれ以上に）、本番環境ではカウントダウンは正確で、30 秒でした。整合位相のみがずれました。違いは何ですか?カウントダウン中はほとんど何も起こりません。マッチング中、それは生きています。注文は WebSocket 経由で到着し、すべてのクライアントは 3 秒ごとに API をポーリングし、価格フィードは 1 秒間に数回データをプッシュし、サーバーは状態を継続的に全員にブロードキャストします。
タイマーは壊れていませんでした。それは飢えていた。
すべてを解く手がかり エラーの形
診断を覆した詳細は次のとおりです。それ以来、これが私のルールになっています。
個別のバグ、つまりリセットされ 2 回カウントされるタイマーは、個別のエラー、つまりクリーンな整数の倍数を生成します。 ×2、×3、決して×2.67。しかし、私は整数倍を持っていませんでした。連続値は 1.2x、1.41x、2.67x、3.5x、5.51x でした。連続ランプ。
また、連続的なランプはロジックのバグとは思えません。それは競合のように見えます。負荷が大きいほど、比例して遅くなります。ストレッチ係数はセッションのアクティビティを追跡しました。そこからはタイマーのバグを狩ることはなくなりました。私はループをブロックするものを探していました。
原因: 1 つの遅いクライアントが全員をブロックした
asyncio は単一のスレッドで実行されます。サーバー全体、タイマー、オーダー、ブロードキャスト、WebSocket ハートビートが 1 つのイベント ループを共有します。このループは協調的です。コードの一部が (実際に制御を譲渡する await を使用して) 譲歩するまで、他には何も実行されません。
私のタイマーは、asyncio.sleep(1) 呼び出しを待ってカウントしました。理論的には、各ループ = 1 秒です。実際には、sleep(1) は、ループがコールバックする時間がある場合にのみ再開します。ループが他の場所でビジーな場合、タイマーの各「秒」は 1 秒にラグを加えたものになります。十分な遅延ループを数えて、10 分間

ute セッションは 50 回実行されます。
主な原因はブロードキャスト機能です。非同期として宣言されているのは問題ありませんが、タイムアウトなしで各クライアントに順番に送信され、タイマーティックごと、取引ごとに待機されました。
ブロードキャスト ループ全体をブロックするには、1 つの低速または半停止状態のクライアント、フリーズしたタブ、有効期限切れのトークン、ネットワーク側の完全な TCP バッファーを使用して、そのクライアントへの送信を送信するだけで済みました。そして、そのループがブロックされている限り、タイマーの刻みは待ち続けられました。他のすべてのクライアントの心拍数も同様でした。
本当の根本原因の利点は、1 つの症状を説明できないことです。それらすべてを説明します。
「セッションの実行時間が長すぎます」 → 飢餓によりタイマーが延長されます。
「時間どおりに切断されました」 → WebSocket ハートビートの到着が遅れました。ブラウザは接続が切れて切断されたと考えました。猶予期間の後、クライアントは排除され、クロックは最後のティックでフリーズしました。
「観念がちらつく」 → 放送が遅れて乱れます。
ログに「トークンの有効期限が切れました」という嵐が発生し、リクエストが遅いため、クライアントはループでの再試行を余儀なくされ、ループの負荷がさらに増大しました。
5 つの異なる訴え、1 つの病気。バグは5つもありませんでした。私はマスクを5枚付けて1枚持っていました。
1. ループカウンターではなく、壁掛け時計。 sleep(1) 呼び出しのカウントを停止しました。各フェーズの開始時に、絶対的な期限を固定します: Deadline = time.monotonic() +duration 。各ティックの残り時間は ceil(deadline - now) であり、フェーズは now >= Deadline になるとすぐに終了します。ループが飽和していてティックが遅れている場合、期限は移動しません。セッションは実際に設定された時刻、期間に終了します。ループのラグによって時間が伸びることはなくなり、追いつきます。 (そして、 datetime.now() ではなく、 time.monotonic() です。NTP 調整の影響を受けず、絶対に逆行しないクロックが必要です。)
2. 短所

現在のブロードキャスト、クライアントごとに制限されています。すべてのクライアントに並行して送信するようにブロードキャストを書き直しました ( asyncio.gather )。各送信は asyncio.wait_for(..., timeout=2s) でラップされます。遅いクライアントですか？そのラウンドではメッセージはスキップされますが、切断されることはなく、速度が遅くなり、停止することもありません。本当に死んだクライアント?削除されました。結果: ブロードキャストは、すべての送信の合計ではなく、最悪の場合でも最大 2 秒に制限されます。
勝利を宣言する前の証拠
私は「展開して指を交差させる」ことをしたくありませんでした。私は、オンデマンドで再現されるティック、スターベーションごとに 1 秒間の同期ブロッキングを挿入するテストを作成しました。
修正前（ループカウンター）：一致×2。
修正後 (壁時計): 挿入されたブロックに関係なく、正確な継続時間。
そして、3 つのクライアント (高速クライアント、低速クライアント (5 秒)、デッドクライアント) を使用した 2 番目のテストで、制限されたブロードキャストの動作を確認します。低速クライアントはスキップされ、デッドクライアントは削除され、高速クライアントは誰も待機しません。どちらのテストも合格します。それから私は展開しました。
「非同期」は「同時」を意味するものではありません。それがこの物語全体の中心にある概念的な間違いです。 asyncio を使用すると、同時実行の可能性が得られます。それはあなたに無料で与えられるわけではありません。各操作を次々に待機するループは、単純な for ループと同じように順次的であり、丁寧にブロックするだけです。
1 つのクライアントの遅さが共有状態に影響を与えないようにしてください。ネットワークと通信するものにはタイムアウトが必要です。 1 人のピアのバックプレッシャーによって全員の時計が人質に取られることがあってはなりません。
時間を反復ではなく時計で測定します。 sleep(1) 呼び出しをカウントすることは、ループが決して遅れないことを賭けます。結局のところ、それはいつもそうなのです。
エラーの形状が診断になります。整数の倍数をクリーンアップ → 離散ロジックのバグを探します。連続体 → 競合を探します。そのたった 1 つの違いが私を何日も救ってくれました。
そしてメタ、それはこのスレッドだから

s ブログ: 本業は開発者ではない IT 責任者が、AI を使用してこのプラットフォームを実稼働環境に導入しました。 AI がコードの大部分を書きました。しかし、このバグはコード生成によって解決されたのではなく、連続体に従い、飢餓仮説を形成し、「壁時計」と「カウンター」のどちらかを選択するという推論によって解決されました。 AI は、原因がわかったら、計測し、測定し、修正を作成するための優れたパートナーでした。理解力は人間のままでした。
まさに私が3ヶ月前に言った通りです。 AI はエンジニアに取って代わるものではありません。それは彼らに理解という本当の仕事をするための時間を取り戻します。
IT 責任者。IT インフラストラクチャ、ネットワーキング、セキュリティに 14 年の経験があります。

## Original Extract

How I diagnosed asyncio event-loop starvation on a real-time trading platform: one slow client stalled the whole loop. The debugging story, and the two fixes.

FS
I built a real-time trading platform with an AI
Back to blog
I built a real-time trading platform with an AI. Then the clock started lying.
In my previous article, I described giving an AI read-only SSH access to my production infrastructure: auditing, documenting, monitoring. The logical next step was to let it not just read , but build .
So I did. With Claude as a co-pilot, I built a real application: a price-matching platform for OTC interest-rate swaps . Real-time, WebSocket, SSO authentication, an engine that automatically pairs the complementary interests of several counterparties. FastAPI, asyncio, SQLite. Shipped to production with a pilot broker and around fifteen bank counterparties behind him.
Then the broker messaged me: “The session cut off while there was still time on the clock.” And also: “A trader's notionals are flickering, appearing and disappearing.” And: “The timing is weird, the sessions run too long.”
My clock had started lying. Here's how I found out why, and what it taught me about the word “asynchronous.”
The symptom: sessions that don't end on time
The heart of the platform is a fixed-duration matching session. You configure it for, say, ten minutes: a countdown, then a matching window, then it's over. Everyone sees the same clock.
Except no. In production, a session configured for 600 seconds lasted 1202 (double). Another, 1803 (triple). Going back through about forty sessions, I didn't find a clean factor of two. I found this:
A session meant to last ten minutes could run for almost an hour. And nobody understood why, me first.
The first wrong turn: “it's the timer”
The obvious reflex: the timer code has a bug. It double-arms, it forgets to stop, something.
I isolated and tested it on its own. Two-second countdown, three-second matching: exact, a single transition, no double-arming. The timer was perfect in isolation.
Worse (or better): in production, the countdown was exact : 30 seconds, dead on. Only the matching phase drifted. What's the difference? During the countdown, almost nothing happens. During matching, it's alive: orders arrive over WebSocket, every client polls the API every three seconds, a price feed pushes data several times a second, and the server broadcasts state to everyone continuously.
The timer wasn't broken. It was starved .
The clue that unlocked everything: the shape of the error
Here's the detail that flipped the diagnosis, and it has been my rule ever since:
A discrete bug, a timer that resets, that counts twice, produces discrete errors: clean integer multiples. ×2, ×3, never ×2.67. But I didn't have integer multiples. I had a continuum : 1.2×, 1.41×, 2.67×, 3.5×, 5.51×. A continuous ramp.
And a continuous ramp doesn't look like a logic bug. It looks like contention : the more load, the slower, proportionally. The stretch factor tracked the session's activity. From there, I was no longer hunting a timer bug. I was hunting whatever blocked the loop.
The cause: one slow client blocked everyone
asyncio runs on a single thread . The whole server, the timer, the orders, the broadcasts, the WebSocket heartbeats, shares one event loop. That loop is cooperative: until a piece of code yields (with an await that actually cedes control), nothing else runs.
My timer counted await asyncio.sleep(1) calls. In theory, each loop = one second. In practice, sleep(1) only resumes when the loop has time to call it back. If the loop is busy elsewhere, every “second” of the timer lasts one second plus the lag . Count enough late loops, and your ten-minute session runs fifty.
The main culprit: the broadcast function. It was declared async all right, but it sent to each client sequentially, one after another, with no timeout , and it was awaited on every timer tick and every trade .
It took just one slow or half-dead client , a frozen tab, an expired token, a full TCP buffer on the network side, for the send to that client to block the entire broadcast loop. And as long as that loop blocked, the timer tick waited. And so did every other client's heartbeat.
The beauty of a real root cause is that it doesn't explain one symptom. It explains all of them.
“The session runs too long” → the timer is stretched by starvation.
“Cut off with time on the clock” → the WebSocket heartbeat arrived late; the browser thought the connection was dead and disconnected; after a grace period the client was ejected, the clock frozen on its last tick.
“The notionals flicker” → broadcasts arrived late and out of order.
The storm of “token expired” in the logs → slow requests pushed clients to retry in a loop, which loaded the loop even more.
Five different complaints, one single disease. I didn't have five bugs. I had one, wearing five masks.
1. A wall clock, not a loop counter. I stopped counting sleep(1) calls. At the start of each phase, I freeze an absolute deadline: deadline = time.monotonic() + duration . On each tick, the time left is ceil(deadline - now) , and the phase ends as soon as now >= deadline . If the loop is saturated and a tick is late, the deadline doesn't move: the session ends at the real configured time , period. The loop's lag no longer stretches time, it catches up. (And it's time.monotonic() , not datetime.now() : you want a clock that never goes backwards, immune to NTP adjustments.)
2. A concurrent broadcast, bounded per client. I rewrote the broadcast to send to all clients in parallel ( asyncio.gather ), each send wrapped in an asyncio.wait_for(..., timeout=2s) . A slow client? Its message is skipped for that round, not disconnected, it's slow, not dead. A truly dead client? Removed. The result: the broadcast is bounded to ~2 seconds worst case, instead of the sum of all sends.
Proof before declaring victory
I didn't want to “deploy and cross my fingers.” I wrote a test that injects one second of synchronous blocking on every tick , starvation, reproduced on demand.
Before the fix (loop counter): matching ×2.
After the fix (wall clock): exact duration, no matter the injected blocking.
And a second test with three clients, a fast one, a slow one (5 s), a dead one, to check the bounded broadcast behaves: the slow one is skipped, the dead one removed, the fast one waits for no one. Both tests pass. Then I deployed.
“Asynchronous” does not mean “concurrent.” That's the conceptual mistake at the heart of the whole story. asyncio gives you the possibility of concurrency; it doesn't give it to you for free. A loop that await s each operation one after another is as sequential as a plain for loop, it just blocks politely.
Never let one client's slowness touch shared state. Anything that talks to the network must have a timeout. One peer's backpressure must never be able to hold everyone's clock hostage.
Measure time with a clock, not with iterations. Counting sleep(1) calls bets that the loop is never late. It always is, eventually.
The shape of your errors is a diagnosis. Clean integer multiples → look for a discrete logic bug. A continuum → look for contention. That single distinction saved me days.
And the meta, because it's the thread of this blog: a Head of IT who isn't a developer by trade put this platform into production with an AI. The AI wrote much of the code. But this bug wasn't solved by code generation, it was solved by reasoning : following the continuum, forming the starvation hypothesis, choosing between “wall clock” and “counter.” The AI was an excellent partner to instrument, measure and write the fix once the cause was understood . The understanding stayed human.
It's exactly what I said three months ago. AI doesn't replace the engineer. It gives them back the time to do the real work: understanding.
Head of IT, 14 years of experience in IT infrastructure, networking and security.
