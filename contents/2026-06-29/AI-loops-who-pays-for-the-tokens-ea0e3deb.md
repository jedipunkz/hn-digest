---
source: "https://www.jackfranklin.co.uk/blog/claude-code-ai-feedback-skill/"
hn_url: "https://news.ycombinator.com/item?id=48721493"
title: "AI loops: who pays for the tokens?"
article_title: "Tracking feedback with Claude Code - Jack Franklin"
author: "jackfranklin"
captured_at: "2026-06-29T16:59:20Z"
capture_tool: "hn-digest"
hn_id: 48721493
score: 1
comments: 0
posted_at: "2026-06-29T16:39:35Z"
tags:
  - hacker-news
  - translated
---

# AI loops: who pays for the tokens?

- HN: [48721493](https://news.ycombinator.com/item?id=48721493)
- Source: [www.jackfranklin.co.uk](https://www.jackfranklin.co.uk/blog/claude-code-ai-feedback-skill/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T16:39:35Z

## Translation

タイトル: AI ループ: トークンの支払いは誰が行うのか?
記事のタイトル: Claude Code を使用したフィードバックの追跡 - Jack Franklin

記事本文:
Claude Code によるフィードバックの追跡
私は、ゲーム OnTrack をプレイテストする際に、ペア プログラマ、デバッガ、およびフィードバック コレクターとして Claude Code を使用してきました。このセットアップは非常にうまく機能したので、拡張するために小さなカスタム スキルを構築することになりました。この投稿では、それがどのように機能するかを説明します。
(最後の DevBlog は 2024 年 (!) でしたが、私はゲームに取り組み続けました。AI のおかげで、父親になってから全体的には時間が減ったにもかかわらず、私のペースと進歩は速くなりました。)
プレイテストのために開発を一時停止する
開発期間が終了したら、プレイテストを行うのが好きです。私は何かを修正することに抵抗があり、代わりに 1 時間ゲームをプレイし、バグ、考え、提案を書き留めます。
以前は Trello ボードでこれを行っていましたが、最近、Claude をバグ収集の同僚として使用して、マークダウン ファイル内のフィードバックを照合してみることにしました。これは本当にうまくいきました！ Claude を使用する利点の 1 つは、類似したフィードバックや重複を検出したり、古いフィードバックを書き直して関連する新しいポイントを組み込んだりできることです。また、UI、メカニクス、または「即効性」などのさまざまなカテゴリに分けてグループ化しました。
大きなMarkdownファイルの問題
その後、これらのプレイテスト セッションに修正を加えながら、フィードバックを別のドキュメントに移動して完了としてマークするようクロードに依頼しました。 1 時間のプレイテストで、修正したい問題が 50 以上も簡単に見つかりました (ゲームはまだ完成には程遠いです!)。ある日、不完全なアイテムを含む Markdown ファイルの長さは 3,000 行で、完了したアイテムを含むマークダウン ファイルの長さは 4,000 行であることに気付きました。
クロードはこれをうまく管理していましたが、時間が経つにつれて、このリストの管理が難しくなり、より多くのトークンを消費し、非効率になることがわかりました。クロードに未解決の問題、またはすべての問題をリストするよう依頼します。

特定のカテゴリでは、サブセットについてのみ質問している場合でも、Markdown ファイル全体を読み取る必要がありました。
クロードがより効率的にクエリを実行できるように、このデータをより構造化された形式で保存する方法が必要でした。たとえば、説明全体を読まなくてもバグのすべてのタイトルを取得したり、特定のカテゴリにあるすべての項目をリストしたり、項目を完了としてマークし、明示的に要求されない限り再度解析する必要がなくなるなどです。
この表は単純です。各行には、 title 、 詳細 、 priority 、 status 、 category 、 project 、および doned フラグがあります。最後の項目が重要です。完了した項目は削除されることはなく、デフォルトのリスト出力から非表示になるだけなので、必要に応じて完全な履歴が得られます。
そこで、まさにそれを行うフィードバック スキルを作成し、プレイテスト セッションからのアイデアとフィードバックを記録、処理、追跡する方法をクロードに教えました。
その中心となるのは、SQLite データベースと通信する Deno CLI です。このデータベースにより、クロードはフィードバック項目をクエリし、新しい項目を追加できます。
なぜデノなのか？正直に言うと、いつも Node に手を伸ばすのではなく、その中で何かを構築してみたかったからです。ただし、ここでは Node (または他のランタイム/言語) で問題ありません。それに、とにかくクロードがそれを書きました:)
プロジェクトごとに 1 つのデータベースを作成するのではなく、すべてのデータを 1 つのファイル (つまり、 .gitignore 内) に保存します。 AI が項目を検索または追加するときは、項目に --project フラグを付ける必要があります。
AI リストのプロジェクトは次のとおりです (routemaster は OnTrack の古い内部名です)。
run list --project Routemaster 次に、特定の項目の詳細を確認します。このフィードバックに AI を使用することが優れているもう 1 つの理由は次のとおりです。AI にフィードバックを与えると、コードとシステムを分析して、実装に関するメモを作成することもできます。
ショー32を実行します
[ 32 ] #024 — プレーヤーが演奏するときに視覚的に確認できない

アクション (トースト通知システム)
プロジェクト: ルートマスター
優先度: 高
ステータス: オープン
作成日: 2026-04-22 07:08:06
---
観察: プレイヤーが駅を建設したり、ルートを決定したり、橋を設置したり、購入したりするとき
アップグレードの場合、一時的な視覚的な確認はありません。アクションはサイレントに成功します。ルート
特に作成は混乱を招きます。完全なファイナライズ モーダル フローを経た後、
ルートが実際に作成されたかどうかの確認はありません。
分析: 通知システムはすでに存在します。ロガー ( src/log.ts ) はシングルトン EventTarget です
LogMessage オブジェクトを保存し、NewLogMessage イベントを送出します。ログ出力
( src/views/log-output.ts ) はこれをサブスクライブし、除外リストをレンダリングします。
Remove-old-messages 属性が設定されている場合、メッセージは 10 秒後に自動的に期限切れになります。現在使用されているのは、
契約の完了と街の成長。そして項目を追加します (ほぼ同一に見える既存の項目を ID で編集することもできます)。
add --project Routemaster を実行 \
--title "列車の詳細ウィンドウ: 列車が遅延すると旅行進行状況バーが消える" \
--detail "列車が遅れている場合 (赤で表示)、列車の詳細ウィンドウの旅行進行状況バーのレンダリングが完全に停止します。
インジケーターと進行状況バーは状態を共有しているように見え、一方が他方を上書きします。任意の列車をさらに遅延させることで再現可能
2駅以上です。」 \
-- 優先度高 \
--category bug 最後に、項目を完了としてマークします。
実行完了 75 フィードバック SKILL.md ファイル
CLI は直観的で文書化されているため、実際のスキルは非常に簡単です。クロードに CLI コマンドを教え、それらを上手に使用する方法についてのガイダンスを追加します。最も影響力のある部分は、新しいフィードバックを保存する前に重複アイテムや関連アイテムをチェックするように指示していることです。よく説明することに気づきました

気づかないうちに、根本的には同じ問題が 2 つの異なる点で存在しており、クロードはそれを認識します。
他にも注目すべきヒントがいくつかあります。
ショーの前に必ずリストを表示します。タイトルと ID は、完全なアイテムの詳細よりもはるかに少ないトークンを使用するため、スキルはクロードに、詳細に入る前にその土地の状況を把握するように指示します。必要になるまで各タスクの詳細をすべて表示する必要があったくありませんでした。デフォルトで詳細をリストするようにした場合、Markdown ファイルのアプローチでトークンを保存することはありません。
文脈からプロジェクトを推測します。クロードは毎回私に尋ねるのではなく、会話から私たちがどのプロジェクトに取り組んでいるかを把握し、本当に曖昧な場合にのみ質問します。
私が最も驚いたのは、この構造がトークンの節約以外にも、「UI カテゴリの優先度の高いバグは何ですか?」と尋ねることができるという点です。迅速かつ正確な回答が得られると、バックログの処理方法が大きく変わります。 GitHub からスキルを自由に取得して、同様のものを構築した場合はお知らせください。
この投稿が気に入ったら、Bluesky で私をフォローしてみてください。

## Original Extract

Tracking feedback with Claude Code
I've been using Claude Code as a pair programmer, debugger, and feedback collector while I playtest my game, OnTrack . The setup has worked so well that I ended up building a small custom skill to make it scale and in this post I'll take you through how it works.
(That last DevBlog was in 2024 (!), but I have continued to work on the game — with AI my pace and progress has quickened, even if overall I've less time since becoming a dad .)
Pausing development to playtest
After a period of development, I like to do a playtest. I resist fixing things, and instead I play the game for an hour, noting down bugs, thoughts and suggestions as I go.
I used to do this in a Trello board, but I recently decided to try using Claude as my bug collecting coworker and collate feedback in a markdown file. This worked really well! One of the benefits of using Claude is that it was able to detect similar feedback, or duplicates, or re-write old feedback to incorporate new points that were related. It also grouped them as it went under different categories such as UI, mechanics, or "quick wins".
The problem of large Markdown files
As I then followed these playtesting sessions with fixes, I asked Claude to mark feedback as complete by moving it into a separate document. In one hour of playtesting I would easily pick up on 50+ issues I wanted to fix (the game is far from done!) and one day I realised that the Markdown file with incomplete items was 3,000 lines long, and the one with completed items was 4,000 lines long!
Claude was managing this fine - but I could see that as time grew, this list would become hard to manage, eat up more of my tokens, and become inefficient. Asking Claude to list the outstanding issues, or all the issues under a particular category, required it to read the entire Markdown file, even if I was only asking about a subset.
I needed a way to store this data in a more structured form where Claude could query it more efficiently — for example, getting all the titles of bugs without having to read the entire description, or listing all items under a particular category, or marking items as done and never having to parse them again unless explicitly asked.
The table is straightforward — each row has a title , detail , priority , status , category , project , and a done flag. That last one is the key bit: completed items are never deleted, just hidden from the default list output, so there's a full history if you ever want it.
So, I made a feedback skill that does just that and teaches Claude how to log, process and track ideas and feedback from the playtesting session.
At its core is a Deno CLI that talks to an SQLite database. This database lets Claude query feedback items and add new ones.
Why Deno? Honestly, because I wanted to try building something in it rather than always reaching for Node. But Node (or any other runtime/language) would be fine here. Plus, Claude wrote it anyway :)
Rather than one database per project, I keep all the data in a single file (that is in my .gitignore ). When the AI looks up or adds items, it must tag them with a --project flag.
Here is the AI listing projects ( routemaster is the old, internal name for OnTrack):
run list --project routemaster And then looking at detail for a particular item. Here is another reason that using an AI for this feedback shines - as you give it feedback it can also analyse your code and system to make some potential implementation notes as you go:
run show 32
[ 32 ] #024 — No visual confirmation when player performs an action (toast notification system)
Project: routemaster
Priority: high
Status: open
Created: 2026 -04-22 07:08:06
---
Observation: When the player builds a station, finalises a route, places a bridge, or purchases
an upgrade, there is no transient visual confirmation. The action succeeds silently. Route
creation in particular is confusing — after going through the full finalise modal flow there is
no confirmation that the route was actually created.
Analysis: A notification system already exists. Logger ( src/log.ts ) is a singleton EventTarget
that stores LogMessage objects and dispatches NewLogMessage events. LogOutput
( src/views/log-output.ts ) subscribes to it and renders a dismissing list — when
remove-old-messages attribute is set, messages auto-expire after 10 seconds. Currently used for
contract completions and town growth. And adding an item (you can also edit an existing item by ID which looks nearly identical):
run add --project routemaster \
--title "Train detail window: journey progress bar disappears when train is delayed" \
--detail "When a train is running late (shown in red), the journey progress bar in the train detail window stops rendering entirely. The delay
indicator and the progress bar appear to share state and one clobbers the other. Reproducible by letting any train fall behind schedule by more
than 2 stops." \
--priority high \
--category bug And finally, marking an item as complete:
run done 75 The feedback SKILL.md file
Because the CLI is intuitive and well documented, the actual skill is fairly straightforward — it teaches Claude the CLI commands and adds some guidance on how to use them well. The most impactful part has been telling it to check for duplicates or related items before storing new feedback. I found I'd often describe the same underlying issue in two different ways without realising, and Claude catches that.
A couple of the other hints worth calling out:
Always list before show — titles and IDs use far fewer tokens than full item details, so the skill tells Claude to get the lay of the land before diving into specifics. I didn't want it to have to view all the details of each task until I need it. If I made it list the detail by default, I wouldn't be saving any tokens over the Markdown file approach.
Infer the project from context — rather than asking me every time, Claude figures out which project we're working in from the conversation and only asks when it's genuinely ambiguous.
What surprised me most is how much the structure helps beyond the token savings — being able to ask "what are all the high priority bugs in the UI category?" and get a fast, accurate answer genuinely changes how you work through a backlog. Feel free to grab the skill from GitHub and let me know if you build something similar.
If you enjoyed this post, you might like to follow me on Bluesky .
