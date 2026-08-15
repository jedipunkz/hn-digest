---
source: "https://xenodium.com/agent-shell-0-73-updates#agent-shell-enters-the-chat"
hn_url: "https://news.ycombinator.com/item?id=49309755"
title: "Show HN: Agent-shell – vendor-neutral chat with AI agents in Emacs"
article_title: "agent-shell 0.73 updates"
author: "xenodium"
captured_at: "2026-08-15T12:17:57Z"
capture_tool: "hn-digest"
hn_id: 49309755
score: 3
comments: 1
posted_at: "2026-08-15T11:42:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agent-shell – vendor-neutral chat with AI agents in Emacs

- HN: [49309755](https://news.ycombinator.com/item?id=49309755)
- Source: [xenodium.com](https://xenodium.com/agent-shell-0-73-updates#agent-shell-enters-the-chat)
- Score: 3
- Comments: 1
- Posted: 2026-08-15T11:42:16Z

## Translation

タイトル: Show HN: Agent-shell – Emacs での AI エージェントとのベンダー中立のチャット
記事のタイトル: エージェントシェル 0.73 アップデート
説明: もう 1 か月、エージェント シェルの更新が行われます。前回のアップデートを見逃した場合は、0.63 アップデートをご覧ください。この投稿では最新のハイライトを紹介します...

記事本文:
ゼノジウム.com
██ ██ ███████ ███ ██ ██████ ██████ ██ ██ ██ ███ ███
██ ██ ██ ████ ██ ██ ██ ██ ██ ██ ██ ██ ████ ████
███ █████ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ████ ██
██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██
██ ██ ███████ ██ ████ ██████ ██████ ██ ██████ ██ ██
2026 年 8 月 15 日
エージェントシェル 0.73 アップデート
もう 1 か月、エージェント シェルの更新がまた行われます。前回のアップデートを見逃した場合は、0.63 アップデートをご覧ください。この投稿では最新のハイライトを紹介しますが、変更点の完全なリストはここで説明する内容よりもはるかに膨大です。
Agent-Shell は、ACP ( Agent Client Protocol ) を利用した AI エージェントと対話するためのネイティブ Emacs モードです。
昨年 9 月の開始以来 (約 1 年になります)、agent-shell は comint モードを利用したシェルのようなエクスペリエンスを特徴としています。より集中したエクスペリエンスを希望する場合は、ビューポート モード (agent-shell-prefer-viewport-interaction 経由) もあり、チャット モードも追加されました。
チャット モードは、comint モードと、より伝統的なチャットのようなラベル付けエクスペリエンスを融合します。
私たちはここで生活しているため、チャット モードがデフォルトで有効になっています。 OK、それほどエッジの効いたものではありませんが、かなり安全であり (パワードオーバーレイ)、 (setq Agent-shell-chat-mode-enabled nil) によって完全に無効にすることができます。チャット モード自体はマイナー モードなので、 M-x Agent-shell-chat-mode を介していつでもオンとオフを切り替えることができます。
前回の投稿では、ツール呼び出しや

エージェントの思考は、デフォルトではすべて崩壊しています。それがあまりにも静かすぎる場合は、デフォルトで (setq Agent-shell-activity-group-expand-by-default t) によって展開できます。これら 2 つの設定が静かすぎる、またはおしゃべりすぎると思われる場合は、 (setq Agent-shell-activity-group-expand-by-default 'latest) による 3 番目の代替手段が用意されています。設定すると、グループ化された最新のアクティビティのみがデフォルトで展開され、エージェントが別のアクティビティに移動すると自動的に折りたたまれます。
私はこの機能がとても気に入っています (@nhojb の PR に感謝します)。これもデフォルトで有効になっています。
キューイングにいくつかの改善が加えられました。関連するコマンドは、agent-shell-prompt-queue.el の下に統合されました。エージェントがビジー状態のときにプロンプ​​トをキューに入れ、 M-x Agent-shell-prompt-queue 、 M-x Agent-shell-prompt-queue-resume 、および M-x Agent-shell-prompt-queue-remove を使用して保留中のプロンプトを表示、再開、または削除できます。保留中のキューは、新しい送信のたびに表示されるようになりました。
M-x Agent-Shell-prompt-compose は、プロンプトを作成するための専用バッファーを開き、シェルからより独立したものになりました。任意のバッファから呼び出すことができ (適切なシェルに解決されます)、C-c C-c を送信して、実行していたことに戻ります (起動して忘れます)。C-u C-c C-c はプロンプトを送信し、キューに入れられた別のプロンプトをすぐに作成できるようにします。
使用しているエージェントによっては、シェルの初期化に 1 ～ 2 秒かかる場合があります。つまり、新しいシェルへの入力を開始する前に、初期化を待つ必要がありました。不要なので、もうそんなことはありません。シェルプロンプトができるだけ早く提供されるようになったので、入力を開始できます。
Markdown リストは特別なレンダリングを行わなくても簡単に理解できますが、それよりも優れた処理ができるため、正規化されたパディング、インデント、そしてもちろん文明化されたビューを使用して、より適切な処理を行うようになりました。

レッツ。
TAB ナビゲーションの改善 + ヒント
TAB ナビゲーションは、かなり早い段階でエージェント シェルに組み込まれました。 TAB キーを押してバッファ内の任意のセクションに移動し、RET キーを押して折りたたみを切り替えることができるのが気に入っています。それはそれで素晴らしいことですが、Markdown ソース ブロック、リンク、画像などをナビゲーション パーティに取り入れることで、ナビゲーション エクスペリエンスをより豊かにすることができます。特にこれらがなぜでしょうか?もちろん、それらはすべて RET によっても実行可能です。
RET は Emacs でローカル テキスト ファイルへのリンクを開き、必要に応じてブラウザにデリゲートすると想定するのが当然ですが、Markdown ソース ブロックとイメージも同様の処理を受けます。それぞれの RET アクションはそれほど明白ではないかもしれませんが、確実に見つけやすくすることができるため、ヒントを追加します。アクション可能なアイテムの着地点で、何ができるか、どのキーがそれを実行するかのヒントがエコーされるようになりました (たとえば、ソース ブロックでは「RET を押してコピー」、画像では「+ を押して拡大」など)。マウスオーバーイベントでもヒントが表示されます。
Agent-Shell には、画像サイズをカスタマイズするための Agent-Shell-markdown-image-max-width が用意されていますが、これはかなり制限的です。すべての画像が同じというわけではありません。なぜすべての画像を同じ制約内に強制的に収める必要があるのでしょうか。 Agent-shell-markdown-image-max-width は引き続き優先デフォルトを提供しますが、エージェントに Pandoc スタイルのリンク属性で Markdown 画像に注釈を付けるようにすることで、画像を異なるスケールで拡大できるようになりました。
属性ブロックは画像の直後に配置され、幅および/または高さをピクセルまたはパーセンテージで受け取ります。
![alt](image.png){width=300}
![alt](image.png){幅=50% 高さ=200}
些細なことで汗をかくかもしれませんが、これは本当に私の歯車を磨く作業でした。私たちは素晴らしい Markdown テーブル レンダリングを備えています。LLM は完全に整列したテーブルを作成するのが必ずしも得意ではないので、これを実現できるのは嬉しいことです。最良の場合、LLM はテーブルを完璧に調整します。

そうですが、Emacs ウィンドウには幅が広すぎます。幸いなことに、私たちの美しいレンダリングでは、セルがウィンドウに収まるようにセルもラップされています。問題は、Emacs フレームのサイズを変更するか、単にウィンドウを分割した瞬間に、その美しいレンダリングがすべて消えてしまい、次のような巨大な結果になるということです。
わかっています。ここでは細かいことで汗をかきますが、まあ、私たちはこのように生きる必要はありません。 Emacs には世界中のあらゆるフックが備わっているので、ウィンドウの変更を追跡し、文明化された世界に再び参加しましょう。
はるかに小さいスケールでは、画像の自動サイズ変更も必要だったので、ここに示します。
レンダリング時に画像サイズに影響を与えることはできますが、それでも望ましくない画像サイズが生成される可能性があるため、バッファ内のすべての画像をオンデマンドで再スケールできるようになりました。
場合によってはハンマーが実際には適切なツールではないため、ポイントで画像を再スケールすることもできるようになりました。
(リンク、画像、または @file への言及から) ローカル ファイルを開くと、カスタマイズできる標準の表示バッファ アクションである、agent-shell-file-display-action を経由するようになりました。デフォルトでは、すでにファイルが表示されているウィンドウを再利用するか、現在のウィンドウを引き継ぎます。
別のウィンドウ配置が必要な場合は、次のようなことができます。
(setq エージェントシェルファイル表示アクション
'(ディスプレイバッファポップアップウィンドウ))
外部参照のサポート
ローカル ファイル リンクをたどると、オリジンが xref のマーカー スタックにプッシュされるようになりました。そのため、他の Emacs ジャンプと同じように、 xref-go-back ( M-, ) を使用すると、 Agent-Shell の元の場所にすぐに戻ります。
お気に入りの折りたたみキーバインディングを使用する
折りたたみ可能なフラグメントは、agent-shell-ui-fragment-map を使用し、agent-shell-ui-toggle-fragment をバインドするようになりました。折りたたみを切り替える RET バインディングのファンではない場合は、お好みのバインディングを使用できます。
TAB がジャムの場合は、次のようなことができます。
(with-eval-after-load 'agent-shell-ui
(キーエージェントシェルUIフラグメントマップの定義
(kbd "TAB") #'年齢

nt-shell-ui-toggle-fragment))
スタイリングの考え方
スタイリング エージェントの考え方を別の方法で好みたい場合は、新しいエージェント シェル 思考ボディ フェイスを使用すると、それが可能になります。
ストリーミングパフォーマンスも好評を博しました。プロファイリングと改善については @suhail-singh に、#757 のトレース分析とベンチマークについては @Scott-Guest と @claytharrison に感謝します。
「利用可能な構成オプション」セクションに、可能な値が表示されるようになりました。
Agent-shell-agent-configs を関数に設定できるようになり、静的なリストをハードコーディングするのではなく、利用可能なエージェント構成を動的に計算できるようになりました。この関数はアクセスのたびに呼び出されるため、コードがリロードされても最新の状態が維持されます。
利用可能なエージェントのみをリストしたい場合があります。以下に大まかな抜粋を示します。
(setq エージェント-シェル-エージェント-configs
(ラムダ()
(シーケンスフィルター
(ラムダ(メーカー)
(when-let* ((client (ignore-errors
(funcall (map-elt (funcall Maker) :client-maker)
(現在のバッファ))))
(コマンド (map-elt クライアント :コマンド)))
(実行可能ファイル検索コマンド)))
(エージェント-シェル-デフォルト-エージェント-構成-メーカー))))
エージェントメッセージチャンクをサブスクライブする
Agent-Shell-subscribe-to は、エージェント メッセージ チャンク イベントをブロードキャストするようになりました。これは、ストリーミング出力を観察したい外部統合に便利です。
executable-find がリモート ホスト (#742 by @CeleritasCelery) で正しく動作するようになり、TRAMP 主導のリモート エージェントがスムーズになります。
Agent-Shell-hq がファミリーに加わり、複数の Agent-Shell セッションを管理するためのインターフェイスを提供します。
コミット ログを覗いてみると、私が毎日 Agent-Shell で作業し、プロジェクトの流入に対応していることがわかります。先月以降、27 件の問題がクローズされ、13 件のプル リクエストがマージされました。今日の時点で、バックログは 11 件の未解決の課題と 5 件の未解決の PR となっています (前回は 13 件と 4 件でした)。優先してほしいことがあれば、お気軽に ping してください。
エージェントシェルのニーズ

あなたのサポート
ベンダー中立のツールの重要性はこれまで以上に高まっており、エージェント シェルの継続を支援する方法がいくつかあります。お金がかかるものもあれば、クリックするだけのものもあります。すべて感謝します ;)
エージェントシェルの寿命を延ばすためのスポンサーシップ
Agent-Shell はインディー開発者である私に過ぎませんが、それと競合するツールには潤沢な資金を持ったチームが背後にいます。 Agent-Shell に費やす時間は、料金を支払うための仕事から離れられる時間です。そのため、それが役立つ場合は、プロジェクトのスポンサーになることを検討してください。また、雇用主がエージェント シェルの使用によって利益を得ている場合は、雇用主にも協力するよう促してください。通常、雇用主は個人ではできない規模で貢献できます。
GitHub のスターは露出を促進し、新しいユーザーや潜在的なスポンサーを惹きつけます。 Agent-Shell のスター付けには費用はかからず、より多くの資金調達に役立つ可能性があるため、数回のクリックを気にしなければ、プロジェクトは実際に別の GitHub スターを使用できます。
これらの改善に貢献してくれたすべての貢献者に感謝します。
#730 : セッション/プッシュ中に送信されたリクエストをキューに入れ、終了時に送信します ( @catern )
#737 : グースの回避策を追加 ( @bergmannf )
#740 : when-active オプションを Agent-shell-activity-group-expand-by-default に追加します ( @nhojb )
#742 : リモートホスト上の実行可能ファイルの検索を修正 (@CeleritasCelery)
#743 : 太字/斜体/取り消し線をレンダリングするときに「要求された」領域を保持する (@alberti42)
#746 : システムスリープ時のロード試行の繰り返しを避ける (@liaowang11)
#748 : エスケープされた Markdown 句読点のプロパティを保持する ( @Scott-Guest )
#752 : ガードグループのメンバーが非前進ブロック範囲に向かって歩く (@hamza-m-masood )
#756 : テーブルレンダリング中にポイントを保持する (@Lenbok)
#762 : ドキュメント ビューポート ワークフロー (@KarimAziev)
#763 : 生のツール出力をレンダリングする (@mrychlik)
#765 : 構文の強調表示によって他のバッファーのモードフックが遅延するのを防ぐ ( @Scott-Guest )
#766 : ビューポートヘッダーを更新

nil の場合でも、agent-shell-prefer-viewport-interaction ( @catern )
持続可能なものにし、エージェントシェルのスポンサーとなる
エージェントシェルが好きですか?それが進化するのを見たいですか？この取り組みを後援することを検討してください。
プライバシーポリシー・利用規約

## Original Extract

Another month, another agent-shell update. If you missed the last one, have a look at the 0.63 update. While this post showcases the latest highlights...

xenodium.com
██ ██ ███████ ███ ██ ██████ ██████ ██ ██ ██ ███ ███
██ ██ ██ ████ ██ ██ ██ ██ ██ ██ ██ ██ ████ ████
███ █████ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ████ ██
██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██ ██
██ ██ ███████ ██ ████ ██████ ██████ ██ ██████ ██ ██
August 15, 2026
agent-shell 0.73 updates
Another month, another agent-shell update. If you missed the last one, have a look at the 0.63 update . While this post showcases the latest highlights, the full list of changes is far chunkier than what we'll cover.
agent-shell is a native Emacs mode to interact with AI agents powered by ACP ( Agent Client Protocol ).
Since inception in September last year (yikes nearly a year), agent-shell has featured a shell-like experience, powered by comint mode . There's also viewport mode (via agent-shell-prefer-viewport-interaction ), if you prefer a more focused experience, and now we have chat mode .
Chat mode fuses comint mode with a more traditional chat-like labelling experience.
We're living on the edge here, so chat mode is now enabled by default. Ok not really that edgy, it's fairly safe (powered overlays ) and can be disabled entirely via (setq agent-shell-chat-mode-enabled nil) . Chat mode itself is a minor mode, so you can always toggle it on and off via M-x agent-shell-chat-mode .
In the last post, we talked about making agent-shell less chatty with more grouping for the likes of tool calls and agent thinking, all collapsed by default. If that's far too quiet, you can expand by default via (setq agent-shell-activity-group-expand-by-default t) . If you found these two settings either too quiet or too chatty, we now have a third alternative via (setq agent-shell-activity-group-expand-by-default 'latest) . When set, only the latest grouped activity is expanded by default, and automatically collapsed when the agent moves on to something else.
I've become quite fond of this feature (thank you @nhojb for the PR ), so yes. It's also enabled by default.
Queuing received some improvements. The related commands have been consolidated under agent-shell-prompt-queue.el . You can queue prompts while the agent is busy, then view, resume, or drop pending prompts via M-x agent-shell-prompt-queue , M-x agent-shell-prompt-queue-resume , and M-x agent-shell-prompt-queue-remove . The pending queue is now shown after each new submission.
M-x agent-shell-prompt-compose opens a dedicated buffer for crafting a prompt, and it's now more independent of the shell. You can invoke it from any buffer (it resolves to the right shell), C-c C-c sends and returns you to whatever you were doing (fire and forget), and C-u C-c C-c submits the prompt and immediately lets you craft another queued prompt.
Shell initialization may take a second or two, depending on what agent you're using, which meant you had to wait for initialization before you could start typing into your new shell. Unnecessary, so that's no longer the case. Shell prompts are now offered as soon as possible, so one can get typing.
While Markdown lists are easily digestible without special rendering, we can do better than that, so we now give them a better treatment with normalized padding, indentation, and of course, civilized bullets.
TAB navigation improvements + hints
TAB navigation made it into agent-shell fairly early on. I love being able to TAB my way into any section in the buffer and press RET to toggle folding. That's great and all, but we can make the navigation experience richer by welcoming the likes of Markdown source blocks, links, and images to the navigation party. Why these in particular? They are all actionable by RET too, of course.
While you'd rightly assume RET opens links to local text files in Emacs and delegates to browsers when needed, Markdown source blocks and images get a similar treatment. While their respective RET actions may not be as obvious, we can certainly make them much more discoverable, so we now add hints. Landing point on an actionable item now echoes a hint of what you can do and which key does it (say, "Press RET to copy" on a source block, "Press + to enlarge" on an image, and so on). Hints are also shown on mouse over events.
While agent-shell offers agent-shell-markdown-image-max-width for customizing image sizes, it's fairly restrictive. Not all images are the same, so why force them all to fit within the same constraint? agent-shell-markdown-image-max-width continues to offer a preferred default, but you can now scale images differently by getting the agent to annotate Markdown images with Pandoc-style link attributes .
The attribute block goes right after the image, taking a width and/or height in pixels or percentages:
![alt](image.png){width=300}
![alt](image.png){width=50% height=200}
I may be sweating the small stuff here, but this was really grinding my gears. We have lovely Markdown table rendering, which I'm glad we do as LLMs aren't always great at producing perfectly aligned tables. In the best of cases, the LLMs align the table perfectly, but it's just too wide for our Emacs window. Luckily, our lovely rendering also wraps cells to make them fit into our window. The thing is, all that lovely rendering goes out the door the moment you either resize your Emacs frame or merely split your window, resulting in a monstrosity like this:
I know. I'm sweating the small stuff here, but hey we don't have to live like this. Emacs has all the hooks in the world, so let's track window changes and rejoin the civilized world.
On a much smaller scale, I also wanted auto resize for images, so here you have it…
While we can influence image size at render time, this can still generate undesirable image dimensions, so we can now rescale all images in buffer on demand.
Sometimes a hammer really isn't the right tool, so we can now also rescale an image at point.
Opening a local file (from a link, image, or @file mention) now routes through agent-shell-file-display-action , a standard display-buffer action you can customize. The default reuses a window already showing the file, or takes over the current one.
If you'd like a different window arrangement, you could do something like:
(setq agent-shell-file-display-action
'(display-buffer-pop-up-window))
Xref support
Following a local file link now pushes your origin onto xref 's marker stack, so xref-go-back ( M-, ) brings you right back to where you were in your agent-shell , just like other Emacs jumps.
Use your favorite folding key binding
Foldable fragments now use agent-shell-ui-fragment-map and bind agent-shell-ui-toggle-fragment . If you're not a fan of the RET binding to toggle folding, you can now use your preferred binding.
If TAB is your jam, you can do something like:
(with-eval-after-load 'agent-shell-ui
(define-key agent-shell-ui-fragment-map
(kbd "TAB") #'agent-shell-ui-toggle-fragment))
Styling thoughts
If you prefer styling agent thoughts differently, a new agent-shell-thought-body face lets you do just that.
Streaming performance also received some love. Thanks to @suhail-singh for the profiling and improvements , and to @Scott-Guest and @claytharrison for the trace analysis and benchmarking in #757 .
The "Available config options" section now displays possible values.
agent-shell-agent-configs can now be set to a function, letting you compute the available agent configurations dynamically rather than hard-coding a static list. The function is called on every access, so it stays current across code reloads.
Maybe you'd like to list only available agents. Here's a rough snippet.
(setq agent-shell-agent-configs
(lambda ()
(seq-filter
(lambda (maker)
(when-let* ((client (ignore-errors
(funcall (map-elt (funcall maker) :client-maker)
(current-buffer))))
(command (map-elt client :command)))
(executable-find command)))
(agent-shell-default-agent-config-makers))))
Subscribe to agent message chunks
agent-shell-subscribe-to now broadcasts an agent-message-chunk event, handy for external integrations that want to observe streamed output.
executable-find now works correctly on remote hosts ( #742 by @CeleritasCelery ), smoothing out TRAMP-driven remote agents.
agent-shell-hq joins the family, offering an interface for managing multiple agent-shell sessions.
If you peeked at the commit logs , you'll notice I've been working daily on agent-shell , keeping up with project inflow. Since last month, 27 issues have been closed and 13 pull requests merged. As of today, the backlog sits at 11 open issues and 5 open PRs (versus 13 and 4 last time around). If there's something you'd like me to prioritize, feel free to ping.
agent-shell needs your support
Vendor-neutral tooling matters more than ever, and there are a couple of ways to help keep agent-shell going. Some cost money, others just a click. All are appreciated ;)
Sponsorships for agent-shell longevity
agent-shell is just me, an indie dev, while the tools it competes with have well-funded teams behind them. Time spent on agent-shell is time away from work that pays the bills, so if it's useful to you, please consider sponsoring the project. And if your employer benefits from your agent-shell use, nudge them to chip in too, they can typically contribute at a scale individuals can't.
GitHub stars help with exposure, attracting new users and potential sponsors. Starring agent-shell costs nothing and can potentially help bring in more funding, so if you don't mind a couple of clicks, the project can really use another GitHub star .
Thank you to all contributors for these improvements!
#730 : Queue requests sent during session/push and submit them when it ends ( @catern )
#737 : Add a workaround for goose ( @bergmannf )
#740 : Add a when-active option to agent-shell-activity-group-expand-by-default ( @nhojb )
#742 : Fix executable-find on remote host ( @CeleritasCelery )
#743 : Preserve "claimed" regions when rendering bold / italic / strikethrough ( @alberti42 )
#746 : Avoid repeated system-sleep load attempts ( @liaowang11 )
#748 : Preserve properties on escaped Markdown punctuation ( @Scott-Guest )
#752 : Guard group member walk against non-advancing block range ( @hamza-m-masood )
#756 : Preserve point during table rendering ( @Lenbok )
#762 : Document viewport workflow ( @KarimAziev )
#763 : Render raw tool output ( @mrychlik )
#765 : Prevent syntax highlighting from delaying other buffers' mode hooks ( @Scott-Guest )
#766 : Refresh viewport header even with nil agent-shell-prefer-viewport-interaction ( @catern )
Make it sustainable, sponsor the agent-shell
Liking agent-shell ? Would like to see it evolve? Consider sponsoring the effort.
privacy policy · terms of service
