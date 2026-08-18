---
source: "https://www.windmill.dev/blog/support-automation"
hn_url: "https://news.ycombinator.com/item?id=49347822"
title: "Automating quality support at scale: AI and human in the loop"
article_title: "Automating quality support at scale: AI and human in the loop | Windmill"
image: ""
author: "rubenfiszel"
captured_at: "2026-08-18T16:20:00Z"
capture_tool: "hn-digest"
hn_id: 49347822
score: 5
comments: 1
posted_at: "2026-08-18T16:06:07Z"
tags:
  - hacker-news
  - translated
---

# Automating quality support at scale: AI and human in the loop

- HN: [49347822](https://news.ycombinator.com/item?id=49347822)
- Source: [www.windmill.dev](https://www.windmill.dev/blog/support-automation)
- Score: 5
- Comments: 1
- Posted: 2026-08-18T16:06:07Z

## Translation

タイトル: 大規模な品質サポートの自動化: AI と人間の関与
記事のタイトル: 大規模な品質サポートの自動化: AI と人間の関与 |風車

記事本文:
メイン コンテンツにスキップ Windmill プラットフォーム スクリプト エディター フロー エディター アプリ ビルダー トリガー データ テーブル デプロイとバージョン管理 ローカル開発ワーカー AI サンドボックス 可観測性 RBAC No-ops セルフホスト 価格設定 ドキュメント Cloud Hub OpenAPI ブログ 変更ログ ロードマップ 大規模な品質サポートの自動化: AI と人間の関与
私たちはサポートに迅速かつ適切に回答できることを少し誇りに思っており、成長するにつれてどちらも失われることを望んでいませんでした。
また、製品の構築に費やす時間をサポートに費やすことも望ましくありませんでした。小規模なチームの場合、これら 2 つは通常互いに対立します。
そこで、できる限り多くの作業を自動化しました。
ハードルを下げるわけではありません。多くの場合、コードベース、ドキュメント、顧客のテレメトリを手にした AI は、スレッドをざっと読んでいる人間よりも根本原因を見つけるのが得意です。
しかし、顧客は AI が作成した返信を信頼するのは、その返信が良好なままであるため、送信前に人間がすべてを承認することになります。
顧客のメッセージからリリースされた修正までのリクエストは次のようになります。
私たちのサポートは、共有 Slack/Discord チャネル、電子メールの受信箱、GitHub の問題など、顧客やユーザーがすでにいる場所で行われます。
私たちはすべてを 1 つの Discord キューに集め、その上に独自の Windmill で実行するパイプラインを置きました。パイプラインはコンテキストを取得し、返信の下書きを作成し、多くの場合は修正の下書きも作成します。
すべてはクロードと一緒に地元で構築されました。
仕組みは次のとおりです。
多くのチャネルのうちの 1 つのキュー
各チャネルには独自の入口があり、すべて同じ入口フローを供給します。
Slack Connect チャネルが Webhook にヒットします。フローは新しいメッセージを自動的に結合するため、顧客の最初のメッセージを見逃すことはありません。
専用アドレスへの電子メールは Windmill の電子メール トリガーを介して受信され、他のものと同じ形式に正規化されます。
Discord メッセージは Windmill WebSocket トリガーを介して受信されるため、チャンネル内の新しいメッセージ

とスレッドはリアルタイムで着陸します。
オープンソース リポジトリ上の GitHub の問題は Linear にミラーリングされ、同じトリアージで実行されるため、利益を得るのは顧客にお金を支払うだけではありません。コミュニティのレポートは公開されているため、より厳格なガードレールで同じ扱いを受けます。
Slack メッセージやメールは、Slack やメール内に留まることはありません。これは Discord スレッドになり、実際にキューを処理するのは Discord です。
私たちはトリアージを読み、ドラフトに目を通し、承認を押し、修正が完了するまで、すべてを離れることなく実行します。
返信は、顧客がどこから書き込んだかに返信されるだけです。
したがって、この投稿の残りの部分で「スレッド」と書かれている場合は、顧客がどのチャネルで開始したかに関係なく、Discord を意味します。
この 1 つのキューは、設定以来約 3,100 件の会話を受け入れてきました。
この投稿の残りのスクリーンショットは、1 つのリクエストに従っており、実際のバグは、発見された瞬間から修正が送られるまで、私たち自身でパイプラインを介して実行されました。
チャネルが何であれ、エントリ フローは同じハウスキーピングを行います。つまり、ノイズを削除し (感謝の言葉や主題から外れた冗談は安っぽい分類ステップを通過することはありません)、添付ファイルをオブジェクト ストレージにステージングし、割り当てや閉じるなどのチーム コマンドを処理します。
次に、すべてを機能させるための要素、スレッド マッチングです。
新しいメッセージが必ずしも新しいチケットであるとは限りません。
私たちがスレッドを書き上げてから 3 日後に誰かが返信したり、すでに未解決のバグについて新たなメッセージを開始したりします。
きちんとした解決策は、全員が常にスレッド内で返信することですが、私はそれを強制することを諦めました...そこで、代わりにコードで実行します。軽量モデルが各受信メッセージを読み取り、最近の会話と比較し、既存のスレッドを継続するか新しいスレッドを開始するかを決定します。
一致すると、そのスレッドのコンテキストが再アタッチされ、リロードされます。ミスは新しいものを開きます。
会話st

顧客がどれだけ会話を分散させても、それは 1 つの会話です。
その後、メッセージを渡します。
エントリ フローは、高度な分析をインラインで実行しません。ステップは、後処理フローを別のジョブとしてキューに入れ、TypeScript クライアントで 1 回の runFlowByPath 呼び出しを行って、戻ります。
ハンドラーは 1 秒未満で終了するため、10 個のメッセージが一度に到達した場合でも、1 回の数分間の分析にそれらのメッセージが積み重なることはありません。
遅い思考は帯域外で発生します。
これらのステップのいくつか (ノイズ フィルター、スレッド マッチャー、オーナー ルーティング) は Windmill AI エージェント ステップ であり、モデルがツールを使用して小さな意思決定を行えるようにするために、任意のフローに組み込むことができる同じ構成要素です。
フローは 3 つの段階で実行され、次のいくつかのセクションが順番に続きます。チケットの分析、トリアージとドラフトのスレッドへの公開、その後、返信と修正の保護されたディスパッチが行われます。
非同期ステップは、必要なすべてのコンテキストを使用して、サンドボックス化された Windmill スクリプトで Claude を実行します。
最初に取得される顧客関係書類: プラン、サブスクリプション ステータス、最近の使用状況。請求や商業に関する質問の場合は、一目見ただけで十分に答えられることがよくあります。私たちは独自のデータからそれを組み立てるため、クロードがライブ認証情報に触れることはありません。
コードベース、ドキュメント、GitHub の問題への読み取りアクセスにより、問題がすでに既知の問題であるか、最近のリリースで修正されたかを確認し、Windmill が実際にどのように動作するかに基づいて答えを導き出すことができます。
その顧客自身のテレメトリ。その範囲はインスタンスに限定され、他には何も含まれません。発生したバグは新しいリリースですでに修正されている可能性があるため、最も有用なシグナルは、そのバージョンがどのバージョンであるかということです。クエリはそのスコープから逃れることはできないため、あるチケットが別の顧客のデータを読み取ることはできません。
メッセージの添付ファイル、インライン化されたテキスト、および画像が直接表示されました。
これらすべてのうち、分析により 1 つの小さな構造化オブジェクトが返されます。

:
{
概要: 文字列; // 1 つまたは 2 つの文
customer_draft ? ： 弦 ; // すぐに送信できる返信
修正リクエスト ? : { // 明らかな修正が 1 つある場合のみ
タイトル : 文字列 ;
問題: 文字列 ;
提案された修正: 文字列 ; // file:line と diff
ファイル: 文字列 [ ] ;
} ;
}
概要と下書きはスレッドに直接投稿されます。
したがって、人間が見た時点で、チケットはすでに誰かがチケットの作成を開始したように見えます。
Opus は綿密な分析を行います。軽量モデル (Sonnet、Haiku) は、安価なルーティングと分類を処理します。
顧客が書き戻すと、フローは以前のコンテキストを再ロードしてフォローアップ モードで実行し、コールドで開始する代わりにスレッドを再開します。
1 つの会話、1 つの記録が、どれだけ長く続いても成長します。
返信の下書きは自動的に作成されます。
私たちはそれらを自動的には送信しません。その行は私たちにとって重要です。
システム全体は顧客が下書きを信頼することで動作しており、間違った、またはずさんな AI 応答が初めて誰かに届いた時点で、その信頼は失われます。それ以降、すべての AI メッセージはノイズとして認識されます。
したがって、顧客向けの返信はすべて、最初に承認ステップを経ます。
下書きは、独自のサポート ボットによって投稿されたスレッドに表示され、そのまま送信するか最初に編集するためのボタンが付いています。
編集すると小さな Discord モーダルが開き、削除される前に文言を調整できるため、簡単な修正はスレッドから離れることを意味しません。
内部的には Windmill stop なので、フローは誰かがアクションを実行するまで、ワーカーを保持せずにただ待機します。
私たちは、返信がどのように組み立てられるかについても慎重に考えています。
これはボットから送信されますが、フッターにはレビューしたチームメイトの名前が記載されており、AI が作成したものであることが明確に示され、顧客に直接返信できることが示されています。
そのため、顧客は常に自分が何を取得しているのかを知っています。つまり、指名された人間が承認した AI ドラフトであり、誰かがすべての単語を入力しているわけではなく、ボットがそのふりをしているわけではありません。

ああ、人間でありなさい。
返信を待っている間に修正案を作成する
これら 2 つのブランチは並行して実行されます。つまり、返信がまだ承認キューに残っている間に修正の下書きが作成されます。
分析が 1 つの明らかな変更に到達すると、fix_request が発行され、2 番目のブランチがそれとともに実行されます。
リクエストは、バグそのもの、技術的な問題、提案された変更を中心に書かれており、リクエストを行った顧客ではないため、顧客固有の内容は含まれません。
ルーティングステップにより、チーム名簿から適切な所有者が選択され、リクエストは割り当てられたリニア課題になります。
そこから、並列 AI コーディング エージェント用のオープンソース ダッシュボードである webmux が引き継ぎます。
引き継ぎは緩いままです。webmux を呼び出しません。単に Linear の問題を Todo に移動し、webmux_oneshot ラベルを追加して、修正を所有する必要がある人にそれを割り当てます。
webmux は約 1 分ごとに Linear をポーリングし、その組み合わせを見つけると、担当者自身のマシン上で git ワークツリーをスピンアップし、自律エージェントを開始して PR をドラフトします。
したがって、修正は共有ボックスではなく、レビューするエンジニアのマシン上でドラフトされます。
PR が起動すると、そのリンクは元の Discord スレッドに戻ります。
2 人のエージェントはきれいに分かれました。
トリアージ ステップでは専用のアナライザーが実行されますが、実際に修正を作成するエージェントは、私たちが独自のコーディング セッションに使用するものとまったく同じ CLAUDE.md から実行されます。そのため、AI がどのように Windmill コードを作成するかについては、人間が開始したのか、チケットが開始したのかに関係なく、信頼できる情報源が 1 つあります。
チケットの場合、Webmux を「ワンショット」モードで実行します。これは基本的に追加の指示が 1 つあります。人間が関与しない状態で、チケットから直接プル リクエストに進むようにしてください。
ワンショットで完了できない場合、中途半端な試行はまだ Webmux ダッシュボードに残されており、誰かが引き継ぐことができます。
そして、それが詰まった場所は、通常、CLA の薄い何かを指します。

UDE.md なので、これを修正すると、サポート修正と独自のセッションの両方が改善されます。
すべてのチケットがそのバーをクリアするわけではありません。それが関門のポイントです。明らかな変更が 1 つある場合、確認されたバグ、または小規模な自己完結型機能がある場合にのみ、チケットが自動的にディスパッチされます。
また、パイプライン全体が信頼できない入力で実行されるため、ゲートの最初のチェックはプロンプト インジェクションです。エージェントを操作しようとするチケットには、追跡されるのではなくフラグが立てられます。
実際の製品や設計に関する問い合わせが必要なもの（ほとんどの機能リクエスト）は、すべて私たちにお任せします。
AI が抵抗しても、チームメイトは AI をスレッドから押し出すことができます。
より鮮明な再現で返信するか、適切なファイルを指定すると、そのコンテキストで再実行されます。または、構築することに決めた機能については、続行してディスパッチするように指示するだけです。
修正の所有者を @ メンションして選択することもできます。
いずれの場合でも、スレッドは、ディスパッチ時のリニア問題リンク、または保留された理由などのメモを返します。
これが上記の例がたどる道です。AI は、いくつかの妥当な修正のうちどれが正しいか確信が持てなかったため、修正を保留し、返信の草案のみを作成しました。
チームメイトがアプローチを確認し、ディスパッチするように指示すると、アプローチが再実行され、PR が webmux に渡されました。
そしてそれはPRだけにとどまりません。
webmux の進行状況コメントは、機能するとスレッドに転送され、修正がマージされると、顧客が最初に報告した場所に自動的に ping を送信して、出荷されたことを知らせることができます。
再開されれば、それもまた戻ってきます。
Discord チャネルにドロップされた 1 行は、レビューされたコード変更を間に挟んで「最新リリースで修正されました」という返信にまで到達することができ、そこに到達するまでに 4 つの異なるツール間で誰もそれを子守する必要はありませんでした。
5 月下旬に自動修正パスがオンになって以来、71 件のプル リクエストがオープンされ、そのうち 62 件はプル リクエストをオープンしました。

それらは合併しました。
残りの 9 件はレビューでクローズされました。通常、修正が必要でないか、一発で行うには大きすぎることが判明したためです。誰も開いて座っていません。
最近の例: 新しい AI チャット パネルの背後にダイアログが表示される UI のバグ。ここで修正されました。
これまでのすべては、顧客メッセージから PR に向けて実行されています。
しかし、変更の多くはその逆です。エンジニアが直接何かを修正するか、私たちが独自に機能の構築を開始し、先週顧客から問い合わせられた内容を PR がチケットも添付されずに静かに解決します。
修正はサポート リンクを通じて自動的に送信されます。
野良 PR は、待っていた人に何の関係もありません。私たちは、こうした PR が黙って出荷されることを本当に望んでいません。
したがって、GitHub pull_request Webhook は別のフローをフィードします。
チームメイトがまだ問題に関連付けられていない特技または修正 PR を開くと、モデルはそれを、まだ修正がリンクされていないオープン サポート リクエストのプールと照らし合わせて保持します。
これは意図的に偏執的です。通常、PR は 0 件または 1 件のリクエストに一致し、一致したものが見つかった場合は、何かが作成される前に Discord の承認を待っています。
それを承認すると、PR がその顧客のスレッドにリンクし、そこから同じマージとリリースの通知が引き継がれます。
キュ

[切り捨てられた]

## Original Extract

Skip to main content Windmill Platform Script editor Flow editor App builder Triggers Data tables Deployment & versioning Local dev Workers AI sandboxes Observability RBAC No-ops self-host Pricing Docs Cloud Hub OpenAPI Blog Changelog Roadmap Automating quality support at scale: AI and human in the loop
We're a little proud of how fast and how well we answer support, and we didn't want either to slip as we grew.
We also didn't want support to eat the time we spend building the product, and for a small team those two usually pull against each other.
So we automated as much of it as we could.
Not to lower the bar: an AI with the codebase, the docs and the customer's telemetry in hand is often better at finding the root cause than a person skimming a thread.
But customers only trust AI-drafted replies while those replies stay good, so a human still approves everything before it goes out.
Here's what a request looks like now, from the customer's message to the released fix:
Our support lives where our customers and users already are: shared Slack/Discord channels, an email inbox and GitHub issues.
We funneled all of it into one Discord queue and put a pipeline on top, running on our own Windmill: it pulls in context, drafts the reply, and often drafts the fix too.
The whole thing was built locally with Claude.
Here's how it works.
One queue out of many channels ​
Each channel has its own way in, and they all feed the same entry flow :
Slack Connect channels hit a webhook . The flow auto-joins new ones, so we never miss a customer's first message.
Email to a dedicated address comes in through Windmill's email trigger , normalized into the same shape as everything else.
Discord messages come in over a Windmill WebSocket trigger , so new messages in our channels and threads land in real time.
GitHub issues on the open-source repo get mirrored into Linear and run through the same triage, so it's not only paying customers who benefit: community reports get the same treatment, with tighter guardrails because they're public.
A Slack message or an email doesn't stay stuck in Slack or email; it becomes a Discord thread, and Discord is where we actually work the queue.
We read the triage, eyeball the draft, hit approve, and watch the fix land, all without leaving it.
The reply just goes back out to wherever the customer wrote from.
So when the rest of this post says "the thread," it means Discord, no matter which channel the customer started in.
That one queue has taken in around 3,100 conversations since we set it up.
The screenshots through the rest of this post follow one request, a real bug we ran through the pipeline ourselves, from the moment it lands to a fix getting dispatched.
Whatever the channel, the entry flow does the same housekeeping: it drops noise (a thank-you or some off-topic banter never gets past a cheap classification step), stages attachments to object storage , and handles team commands like assign and close.
Then the piece that makes it all work: thread matching.
A new message isn't always a new ticket.
Someone replies three days after we'd written a thread off, or starts a fresh message about a bug we already have open.
The tidy fix would be for everyone to always reply in-thread, but I gave up trying to enforce that... so we do it in code instead: a lightweight model reads each incoming message, compares it against recent conversations, and decides whether it continues an existing thread or starts a new one.
A match reattaches and reloads that thread's context; a miss opens a fresh one.
The conversation stays one conversation, however much the customer scatters it.
Then it hands the message off.
The entry flow doesn't run the heavy analysis inline: a step enqueues the post-processing flow as a separate job, one runFlowByPath call with the TypeScript client , and returns.
The handler finishes in well under a second, so when ten messages land at once they don't pile up behind a single multi-minute analysis.
The slow thinking happens out of band.
Several of those steps, the noise filter, the thread matcher, the owner routing, are Windmill AI agent steps , the same building block any flow can drop in to let a model make a small decision with tools.
The flow runs in three stages, and the next few sections follow them in order: analyze the ticket, publish the triage and draft into the thread, then a guarded dispatch of the reply and any fix.
The async step runs Claude in a sandboxed Windmill script , with all the context it needs:
A customer dossier , fetched first: plan, subscription status, recent usage. For a billing or commercial question that's often enough to answer at a glance. We assemble it from our own data, so Claude never touches live credentials.
Read access to our codebase, docs and GitHub issues , so it can check whether the problem is already a known issue or got fixed in a recent release, and ground its answer in how Windmill actually behaves.
That customer's own telemetry , scoped to their instances and nothing else; the most useful signal is which version they're on, since the bug they hit may already be fixed in a newer release. The query can't escape that scope, so one ticket can never read another customer's data.
The attachments from the message, text inlined and images looked at directly.
Out of all that, the analysis returns one small structured object:
{
summary : string ; // one or two sentences
customer_draft ? : string ; // a ready-to-send reply
fix_request ? : { // only when there's one obvious fix
title : string ;
problem : string ;
proposed_fix : string ; // file:line and a diff
files : string [ ] ;
} ;
}
The summary and the draft get posted straight into the thread.
So by the time a human looks, the ticket already reads like someone got started on it.
Opus does the heavy analysis; lighter models (Sonnet, Haiku) handle the cheap routing and classification.
If the customer writes back, the flow reloads the earlier context and runs in follow-up mode, picking the thread back up instead of starting cold.
One conversation, one growing record, however long it runs.
We draft replies automatically.
We do not send them automatically, and that line matters to us.
The whole system runs on customers trusting the drafts, and that trust is gone the first time a wrong or sloppy AI reply reaches someone; from then on every AI message reads as noise.
So every customer-facing reply goes through an approval step first.
The draft lands in the thread, posted by our own support bot, with buttons to send it as-is or edit it first.
Editing opens a small Discord modal where you can tweak the wording before it goes out, so a quick fix doesn't mean leaving the thread.
Under the hood it's a Windmill suspend , so the flow just waits, holding no worker, until someone acts on it.
We're also deliberate about how a reply is framed.
It goes out from the bot, but the footer names the teammate who reviewed it, says plainly that it was AI-drafted, and tells the customer they can reply directly.
So the customer always knows what they're getting: an AI draft that a named human signed off on, not someone typing every word, and not a bot pretending to be human.
Drafting the fix while the reply waits ​
Those two branches run in parallel: the fix gets drafted while the reply is still sitting in the approval queue.
When the analysis lands on one obvious change, it emits a fix_request and the second branch runs with it.
The request is written around the bug itself, the technical problem and a proposed change, not the customer who hit it, so nothing customer-specific travels with it.
A routing step picks the right owner off the team roster, and the request becomes an assigned Linear issue.
From there webmux , our open-source dashboard for parallel AI coding agents, takes over.
The handoff stays loose: we don't call webmux, we just move the Linear issue to Todo, add a webmux_oneshot label, and assign it to whoever should own the fix.
webmux polls Linear every minute or so, and when it sees that combination it spins up a git worktree on the assignee's own machine and starts an autonomous agent to draft the PR.
So the fix gets drafted on the machine of the engineer who'll review it, not on a shared box.
When the PR's up, its link drops back into the original Discord thread.
The two agents split cleanly.
The triage step runs a dedicated analyzer, but the agent that actually writes the fix runs off the very same CLAUDE.md we use for our own coding sessions, so there's one source of truth for how the AI writes Windmill code, whether a person or a ticket kicked it off.
For tickets we run webmux in "oneshot" mode, which is basically one extra instruction: try to go from the ticket straight to a pull request with no human in the loop.
When oneshot can't close it out, the half-finished attempt is still sitting in our webmux dashboard for someone to take over.
And wherever it got stuck usually points at something thin in our CLAUDE.md , so fixing that makes both the support fixes and our own sessions better.
Not every ticket clears that bar, and that's the point of the gate: it dispatches on its own only when there's one obvious change to make, a confirmed bug or a small, self-contained feature.
And because the whole pipeline runs on untrusted input, the gate's first check is for prompt injection: a ticket that tries to steer the agent gets flagged instead of followed.
Anything that needs a real product or design call, which is most feature requests, it leaves for us.
When the AI holds back, a teammate can still nudge it from the thread.
Reply with a sharper repro or point it at the right file and it re-runs with that context, or for a feature we've decided to build, just tell it to go ahead and dispatch.
You can also pick who owns the fix by @-mentioning them.
Either way the thread gets a note back: the Linear issue link when it dispatches, or the reason it held off.
That's the path the example above takes: the AI wasn't sure which of a few reasonable fixes was right, so it held the fix back and only drafted a reply.
Once a teammate confirmed the approach and told it to dispatch, it re-ran and handed the PR to webmux.
And it doesn't stop at the PR.
webmux's progress comments get forwarded into the thread as it works, and when the fix merges we can ping the customer automatically, right where they first reported it, to tell them it shipped.
If it ever gets reopened, that comes back too.
A single line dropped in a Discord channel can travel all the way to a "fixed in the latest release" reply, with a reviewed code change in between, and nobody had to babysit it across four different tools to get there.
Since the auto-fix path switched on in late May, it has opened 71 pull requests, and 62 of them have merged.
The other 9 were closed on review, usually because the fix turned out not to be needed or too big for a one-shot; none are sitting open.
One recent example: a UI bug where dialogs rendered behind the new AI chat panel, fixed here .
Everything so far runs from a customer message toward a PR.
But plenty of changes go the other way: an engineer fixes something directly, or we start building a feature on our own, and the PR quietly resolves something a customer asked about last week, with no ticket attached.
Fixes dispatched through support link themselves automatically.
A stray PR has nothing tying it back to the person who was waiting, and we really don't want those shipping in silence.
So a GitHub pull_request webhook feeds another flow.
When a teammate opens a feat or fix PR that isn't already tied to an issue, a model holds it up against the pool of open support requests that still have no fix linked.
It's deliberately paranoid: a PR usually matches zero or one request, and any match it does find waits behind a Discord approval before anything is created.
Approve it, and the PR links to that customer's thread, and from there the same merge-and-release notifications take over.
The cu

[truncated]
