---
source: "https://lanes.sh/blog/claude-fable-5"
hn_url: "https://news.ycombinator.com/item?id=48475266"
title: "Show HN: Lanes now runs Claude Fable 5 across parallel agent sessions"
article_title: "Claude Fable 5 is here. You can run it in Lanes today. | Blog | Lanes"
author: "s-xyz"
captured_at: "2026-06-10T13:19:07Z"
capture_tool: "hn-digest"
hn_id: 48475266
score: 3
comments: 0
posted_at: "2026-06-10T12:25:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lanes now runs Claude Fable 5 across parallel agent sessions

- HN: [48475266](https://news.ycombinator.com/item?id=48475266)
- Source: [lanes.sh](https://lanes.sh/blog/claude-fable-5)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T12:25:56Z

## Translation

タイトル: Show HN: Lanes が並列エージェント セッション全体で Claude Fable 5 を実行するようになりました
記事タイトル：クロード・ファブル5はこちらです。今すぐレーンで実行できます。 |ブログ |レーン
説明: Anthropic は最初の Mythos クラス モデルを出荷しました。 v0.43.0 からは、他の Claude と同様に Lanes モデル ピッカーに配置されます。それが何であるか、そのコストは何か、そしてなぜその長期的な自律性が並行セッションのボードに適合するのか。

記事本文:
クロード・ファブル5はここにあります。今すぐレーンで実行できます。 |ブログ | Lanes Lanes Docs スタート ブログ / 業界 クロード ファブル 5 はここにあります。今すぐレーンで実行できます。
Anthropic は昨日、一般公開された最初の Mythos クラス モデルである Claude Fable 5 をリリースしました。料金は 100 万トークンあたり 10 ドル、50 ドルで、6 月 22 日まで追加料金なしでクロードのサブスクリプションに含まれます。 v0.43.0 では、Lanes モデル ピッカーにも含まれています。
2 つの名前、1 つのモデル。 Mythos 5 は、許可されたユーザーのみを対象とした無制限バージョンです。 Fable 5 は同じ重みで、前面に安全装置があり、それが誰もが得られるバージョンです。
斬新な部分は安全装置です。分類子がリクエストにフラグを立てた場合、Fable 5 は拒否しません。Claude Opus 4.8 にフォールバックして、そこで応答します。 Anthropic によれば、それが影響するセッションは 5% 未満です。
入力トークン 100 万件あたり 10 ドル、出力 100 万件あたり 50 ドルで、Mythos プレビューの費用の半分以下です。 API モデル ID は claude-fable-5 です。サブスクリプションは随時行われます。Pro、Max、Team、およびシートベースの Enterprise では 6 月 22 日まで無料で、6 月 23 日からは使用量クレジットから引き出されます。
繰り返す価値のある証拠: Stripe は、「数か月のエンジニアリングを数日に圧縮」し、2 か月の範囲で 5,000 万行の Ruby の移行を 1 日で完了したと述べています。ビジョンは最先端です。永続的なファイルベースのメモを考慮すると、Slay the Spire では Opus 4.8 の管理より 3 倍も進歩しました。このスレッドは短い質問に対する鋭い回答ではありません。それは、有能で監督のない長時間にわたる作業です。
Lenny のニュースレターには早期アクセスがありました。評決は、「ベンチマークを破る」が、「設計上トークンを大量に消費」し、強力な計画とフォロースルーを躊躇する「実行に関して保守的」であるというものだ。 2 つのことが続きます。出力トークンが実際の請求額となるため、50 ドル側の方が重要です

10ドルよりも。そして、警戒心を議論するのではなく、努力のレベルでそれを操縦します。
実際に何に使うのか
初期のレポートでは 4 つの機能が循環し続けており、レーンはそれぞれに移動します。
監視なしでの長時間の実行。人間的なリーダーは、これまでのクロードを上回る自律性を備えています。並列セッションとして表示されるボード上: 各 Fable 5 は独自のワークツリーで実行され、必要なときにベルが鳴ります。リードを長く持つほど、より長く飛び続けることができます。
リポジトリ全体の移行。 Stripe は、2 か月に及ぶ 5,000 万行の Ruby の移行を 1 日に短縮しました。それを 1 つのセッションに任せることはありません。モジュールごとに 1 つの問題に分割し、各ワークツリーにわたって Fable 5 セッションを展開し、差分を Git Changes に並べて配置します。
強力な計画、躊躇する実行。レニーの解釈は、「綿密に計画を立ててから、その後の進捗を後から推測する」というものです。レーンはすでにこれらのステップを分割しています。計画モードを実行して計画を取得し、実行モードに切り替えて計画を押し進め、計画が行き詰った場合は、慎重に議論する代わりに努力ノブに手を伸ばします。
長期にわたって生き残るコンテキスト。永続的なファイルのメモを考慮すると、Slay the Spire では Opus 4.8 よりも 3 倍進歩しています。すべての Lanes セッションはデフォルトでそのように設定されています。メモを保持するための独自のワークツリーがあり、再開すると再起動後に完全な履歴が復元されるため、学習した内容は実行間で蒸発しません。
v0.43.0 以降、claude-fable-5 はセッション設定にあり、v0.41 で導入されたのと同じモデルおよびエフォート ピッカーです。セッションを起動するときに選択するか、 .lanes/cli-flags.json 内のプロジェクトに固定します。努力レベルも設定します。これはノブが何かをするモデルです。
そして、Anthropic が先月請求を分割したときと同じ理由で、請求は維持されます。 Lanes は実際の端末で公式のクロード CLI を実行します。Anthropic はサブウーファー側でこれをインタラクティブな使用としてカウントします。

スクリプトはまだカバーしています。 6 月 22 日まで、レーン セッションの Fable 5 は Pro または Max プランで無料です。
Mythos クラスは現在、研究プレビューではなく層であり、フォールバックのアイデアはコピーされるでしょう。より安全なモデルで答えることは、拒否することに勝ります。実際のポイントはもっと簡単です。 Anthropic が出荷した最強のロングホライズン コーディング モデルはモデル ピッカーに組み込まれており、今後 12 日間の計画に含まれています。
問題を選択し、モデルを claude-fable-5 に設定し、長いリードでどうなるかを確認します。初めての方はここから始めるか、Discord に参加して見つけたものを教えてください。
2026 年 6 月 4 日 v0.42 の新機能: データベースの内部を見てみる
v0.42 では、データベース エクスプローラーがコンテキスト サイドパネルに追加されます。作業フォルダー内の SQLite データベースの自動検出、テーブルとビューの参照、組み込みエディターからの読み取り専用クエリの実行をすべて、レーンを離れることなく実行できます。さらに、重複問題アクションと一連の安定性修正も含まれます。
2026 年 5 月 21 日 v0.41 の新機能: 多数のセッション、1 つの問題
同じ問題に対して複数のエージェント セッションを並行して実行します。実行する計画を連鎖させて、レビューしたり、アプローチを並べて比較したり、単に作業を整理したりすることができます。クイック アクションはセッション セレクターから起動できるようになり、セッション設定では CLI、モデル、作業量を事前に選択できるようになりました。さらに、入力待ち、トークン取り消し、プロジェクト切り替え全体の安定性も修正されました。

## Original Extract

Anthropic just shipped its first Mythos-class model. From v0.43.0 it sits in the Lanes model picker like any other Claude. What it is, what it costs, and why its long-horizon autonomy fits a board of parallel sessions.

Claude Fable 5 is here. You can run it in Lanes today. | Blog | Lanes Lanes Docs Get Started Blog / Industry Claude Fable 5 is here. You can run it in Lanes today.
Anthropic released Claude Fable 5 yesterday, the first Mythos-class model to reach general availability. It costs $10 in and $50 out per million tokens, and through June 22 it is included in Claude subscriptions at no extra cost. As of v0.43.0 it is in the Lanes model picker too.
Two names, one model. Mythos 5 is the unrestricted version, for authorized users only. Fable 5 is the same weights with safeguards in front, and that is the version everyone gets.
The novel part is the safeguards. When a classifier flags a request, Fable 5 does not refuse: it falls back to Claude Opus 4.8 and answers there. Anthropic says that touches fewer than 5% of sessions.
$10 per million input tokens, $50 per million output, less than half of what Mythos Preview cost. The API model ID is claude-fable-5 . Subscriptions are on a clock: free on Pro, Max, Team, and seat-based Enterprise through June 22, then drawn from usage credits from June 23.
The proof points worth repeating: Stripe says it "compressed months of engineering into days," finishing a 50-million-line Ruby migration in a day scoped at two months. Vision is state of the art. Given persistent file-based notes, it got three times further in Slay the Spire than Opus 4.8 managed. The thread is not sharper answers to short questions. It is longer stretches of competent, unsupervised work.
Lenny's Newsletter had early access. The verdict: it "crushes benchmarks," but it is "token-intensive by design" and "conservative on execution," with strong plans and hesitant follow-through. Two things follow. Output tokens are the real bill, so the $50 side matters more than the $10. And you steer it with effort levels, not by arguing with its caution.
What you'd actually use it for
Four capabilities the early reports keep circling, and the Lanes move for each.
Long, unsupervised runs. Anthropic leads with autonomy that outlasts any previous Claude. On a board that shows up as parallel sessions: each Fable 5 run in its own worktree, a bell when one needs you. The longer it holds the leash, the more you keep in flight.
Repo-wide migrations. Stripe compressed a two-month, 50-million-line Ruby migration into a day. You don't hand that to one session. Split it one issue per module, fan a Fable 5 session across each worktree, and let the diffs land side by side in Git Changes.
Strong plans, hesitant execution. Lenny's read: it plans well, then second-guesses the follow-through. Lanes already splits those steps. Run plan mode to get the plan, switch to implement mode to push it through, and reach for the effort knob when it stalls instead of arguing with its caution.
Context that survives the long haul. Given persistent file notes, it got three times further in Slay the Spire than Opus 4.8. Every Lanes session is that setup by default: its own worktree to keep notes in, and resume restores the full history across restarts, so nothing it learned evaporates between runs.
From v0.43.0, claude-fable-5 is in Session Settings, the same model and effort picker that arrived in v0.41. Choose it when you launch a session, or pin it for a project in .lanes/cli-flags.json . Set the effort level too; this is a model where that knob does something.
And the billing holds up, for the same reason it did when Anthropic split the bill last month. Lanes runs the official claude CLI in a real terminal, which Anthropic counts as interactive use, the side your subscription still covers. Through June 22, Fable 5 in a Lanes session is free on your Pro or Max plan.
Mythos-class is a tier now, not a research preview, and the fallback idea will get copied: answering with a safer model beats refusing. The practical takeaway is simpler. The strongest long-horizon coding model Anthropic has shipped is in your model picker, included on your plan for the next twelve days.
Pick an issue, set the model to claude-fable-5 , and see what it does with a long leash. Start here if you are new, or join our Discord and tell us what you find.
Jun 4, 2026 What's New in v0.42: Look Inside Your Databases
v0.42 adds a Databases Explorer to the context sidepanel: auto-discover SQLite databases in your working folder, browse tables and views, and run read-only queries from a built-in editor, all without leaving Lanes. Plus a Duplicate Issue action and a run of stability fixes.
May 21, 2026 What's New in v0.41: Many Sessions, One Issue
Run multiple agent sessions on the same issue in parallel. Chain plan to implement to review, compare approaches side by side, or just keep work organised. Quick Actions now launch from the session selector, and Session Settings lets you pick CLI, model, and effort upfront. Plus stability fixes across awaiting-input, token revocation, and project switching.
