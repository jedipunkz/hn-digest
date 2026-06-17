---
source: "https://provi.me/cc-reflection"
hn_url: "https://news.ycombinator.com/item?id=48567237"
title: "cc-reflection: teaching Claude Code to reflect"
article_title: "cc-reflection: teaching Claude Code to reflect - provi"
author: "Bluestein"
captured_at: "2026-06-17T10:39:37Z"
capture_tool: "hn-digest"
hn_id: 48567237
score: 3
comments: 0
posted_at: "2026-06-17T07:59:40Z"
tags:
  - hacker-news
  - translated
---

# cc-reflection: teaching Claude Code to reflect

- HN: [48567237](https://news.ycombinator.com/item?id=48567237)
- Source: [provi.me](https://provi.me/cc-reflection)
- Score: 3
- Comments: 0
- Posted: 2026-06-17T07:59:40Z

## Translation

タイトル: cc-reflection: クロード コードにリフレクションを教える
記事のタイトル: cc-reflection: クロード コードにリフレクションを教える - 但し書き
説明: クロード コードの使用を反映して、作業 (プロジェクト) とワークフロー (プロセス) の両方を継続的に改善できるようにするためのフレームワーク。

記事本文:
cc-reflection: クロード コードにリフレクションを教える - provi
provi ビルドについてのEN | 書き込み 简 cc-reflection: クロード コードにリフレクションを教える
クロード コードの使用を振り返り、作業 (プロジェクト) とワークフロー (プロセス) の両方を継続的に改善できるようにするためのフレームワーク。
Claude Code には EDITOR フック (Ctrl-G) があります。エディタで一時ファイルを開き、書き込み、保存、閉じると、書き込んだ内容がクロードの入力ボックスに再び挿入されます。フックは、開いてから閉じるまでの間に何が起こっても関係ありません。ファイルを監視するだけです。
そのギャップはブラックボックスです。ギャップをfzfメニューに置き換えました。 Ctrl-G を押すと、あらゆる操作を実行できるコントロール サーフェスが開きます。IDE でプロンプトを編集し、エージェントでプロンプトを強化し、リフレクション シードを参照して展開し、その結果をプロンプト ファイルに折りたたんで戻します。
cc-reflection は、そのブラックボックスの中から生まれたものです。それは 1 つのシステムではなく、複数のシステムが接続されています。クロードに自己検査用の 3 つのレンズを教えるリフレクション スキル、有機的にトリガーされるまでターンを重ねるごとに圧力を高めるダイス アキュムレーター、観察結果をキャプチャするシード ストア、およびシードをコードベースに基づいた成果物に拡張する思考エージェントです。 EDITOR フックは結合組織であり、メニューはコンテナです。それぞれの部分は独立して機能しますが、フックによって共有面が得られます。
私たちが実際にどのように反映するかをモデル化する
反省は常に事後に起こります。あなたは何が起こっているかではなく、何が起こったかを観察します。作品内の何かがあなたの注意を引き、あなたはこう思います。それがイベント全体です。ちょっとした気づき。
ほとんどの場合、主役の途中です。直感に基づいてセッションを方向転換することは望ましくありません。それで、あなたはそれを手放します。溶けてしまいます。半ば気づいていたものはすべてノイズに戻り、二度とそれについて考えることはありません。
cc-reflectionはそれを与えます

着陸する瞬間。 /reflection がトリガーされると、エージェントはデザイナーのスキルを使用して表面化する価値のあるものがあるかどうかを評価します。そうでない場合もありますが、それは問題ありません。存在する場合は、すぐに修正するか、無視するか、シードを作成するかの 3 つの選択肢があります。
面白いのは種です。それは、提案されたアーティファクトで観察を捕らえます - この考えがどのような具体的な行動につながる可能性があるか。次に進みます。セッションは継続します。種は背景、枠組み、頭の後ろにあり、成長します。
実際に思考に取り組むのは費用がかかります。文脈を切り替えたり、複数の視点を保持したり、作品の外に出て作品を見ることが必要です。現時点ではそれを行うことはほとんどありません。後で、余裕があるときに戻って、メニューを開きます。種子は鮮度別に分類されてそこにあります。どれか 1 つを選んで参加します。ここで思考エージェントが登場します。思考エージェントは完全な対話型セッションを取得します。シードを読み取り、関連するコードを取得し、コンテキストを再構築します。それを操縦するのはあなたです。それ以来抱いていた考えを追加します。これにより、シードがコードベースに基づいた詳細な分析に拡張され、明確な次のステップが示されます。
この拡張は、単一の観察から成長した洗練された計画モードのように、メイン エージェントにフィードバックするプロンプトになる可能性があります。あるいは、思考エージェントに直接変更を加えさせます。使い方は自由です。
デザインの核となる賭けは、同じ注目を得るために、気づくことと分析することが競合するということです。その場で両方を実行しようとすると、後から推測して実行を脱線させるか、フローを維持するために観察を抑制することになります。種を使えば分析しなくても気づくことができます。拡張機能を使用すると、慌てることなく分析できます。
/reflection スキルは、孔子から借用した 3 つのレンズをクロードに教えます (吾日三省吾身 - 私は毎日 3 回自分自身を振り返ります)。
一省：

「これは正しく構築できていますか?」 - アーキテクチャ、セキュリティ、複雑さ、エンジニアリング パターン
二省：「私は正しいものを作っているでしょうか？」 - UX、プロダクト、ビジネスロジック、ユーザーメンタルモデル
三省：「私はどうやって働いているんですか？」 - プロセス、慣習、ワークフロー、スキル、フック
これらはメモを整理するためのカテゴリではありません。これらは、スキルの発動時にエージェントが順番に適用するレンズです。それぞれについて、内部で苗木が生成され、それらがテストされ (「これはセッション中に実際に感じたことに根ざしているのか?」)、生き残れないものは破棄されます。すでに知られていることを単に言い換えたり、修正されたばかりのバグを説明したりするだけの苗木は拒否されます。ほとんどの反射は何も生み出しません。
それは仕様によるものです。 今日無省 - 今日は何も調べる必要がありません。このスキルは、沈黙を失敗ではなく明晰さとして扱います。スペースを埋めるために洞察を捏造すると、種子の保管場所がノイズで汚染されてしまいます。フレームワークの最も重要な出力は、多くの場合、まったく出力されません。
洞察をスケジュールすることはできません。本当の反省が起こるのは、何かが突き抜けられるほどの未検討の経験が蓄積されているからです。セッションが長ければ長いほど、その可能性は高くなります。
ダイス アキュムレータはこれをモデル化します。反射なしで 7 (デフォルト) ターンが経過すると、システムは毎ターン舞台裏で 1d20 のローリングを開始します。 14ターン後、2d20。 21以降は3d20。自然な 20 は、エージェントが反映するための明示的なナッジとリセットをトリガーします。短いセッションが中断されることはありません。長いセッションでも最終的には取得されますが、いつ取得されるかは予測できません。
サイコロはいつ反映するかを決定します。このスキルは、セッションの軌道が変化した瞬間など、何を探すべきかを知っています。エージェントは何かを推測し、修正されました。標準的なアプローチは 2 回失敗し、創造的な回避策を余儀なくされました。ユーザーは別の方向にハンドルを切り続けました

。何かが必要以上に時間がかかりました。これらの状態遷移は興味深い観察が行われる場所であり、スキルはそこに固定されるようにトレーニングされます。
Ctrl-G を押すと、鏡の間が開きます。
エディターでプロンプトを編集します - vi、VS Code、Cursor、Zed、Windsurf、Antigravity。現在のプロンプトを開き、編集、保存、閉じると、プロンプトがクロードの入力ボックスに戻されます。
プロンプトを強化します - ドラフト プロンプトを tmux ウィンドウ内に存在する別のエージェントに渡します。プロンプトは、明確にするためにプロンプ​​トを書き直します。入力ボックスに送り返す前に、構造を追加し、すべてのファイル パスが存在することを確認し、受け入れ基準を追加します。インタラクティブ (エンハンサーとのチャット) または自動 (ファイアアンドフォーゲット) モードで利用できます。
シードを参照する - すべてのプロジェクトに根ざしたシードが、鮮度インジケーター (🌱 新鮮、💭 成長中、💤 古い、📦 アーカイブ済み) とともにリストされます。プレビュー ペインには、現在の理論的根拠と拡張履歴が表示されます。思考エージェントを拡張するために 1 つを選択します。
シードを展開する - シードを選択すると、思考エージェントが引き継ぎます。関連するコードを読み取り、コンテキストを再構築し、対話的に操作します。読み取っていないコードについて推測することはありません。展開とその結論はシードに永続的に添付されます。
設定 - モデル (作品/ソネット/俳句)、拡張モード、権限、フィルター、コンテキスト ウィンドウの深さを切り替えます。全部しつこい。
ガーデニング - 古くなったシードを整理してアーカイブし、古いシードを削除し、個々のシードを削除します。
私の初期バージョンの cc-reflection には、エージェントが代わりにトランスクリプトを読むサブエージェントにリフレクションを委任できるバックグラウンド モードがありました。それは決して引き起こされませんでした。
これには実際的な摩擦もあります。クロード コードのサブエージェントは親の完全なコンテキストを継承しません。タスク プロンプトが表示された新しいウィンドウが表示されます。彼らができる最善のことは、トランスクリプトファイルを第三者として読み取ることです。

d人。どうやら、メインエージェントはこれを選択しなかったようです。
もっと微妙なことがあります。そして、これは私が直接の経験から主張していることです。完璧なコンテキスト転送があっても、実行エージェントはセッションを「ホット」にします。最近の間違い、放棄されたアプローチ、中途半端に形成された懸念が、注目のウィンドウに不釣り合いな重みをもたらします。同じトランスクリプトを読み取るサブエージェントは、すべてを同じ重みでコールドに扱います。
私はプロジェクト内でこれを「ダークマター」と呼んでいます。技術的にはトランスクリプトに存在するものの、引き継ぎによって事実上失われるコンテキストです。前景の反射によりそれが維持されます。代表団はそれを剥奪する。
最近の論文 (Dadfar 2026、「When Models Examine Themselves」) は、これに関する機構的な証拠を提供しています。 LLM が自身の処理を検査する場合、LLM が生成する語彙は同時活性化のダイナミクスを追跡します (r=0.44) が、非自己参照コンテキスト内の同じ単語は、9 倍の頻度であっても一致を示しません (r=0.05)。この論文では、「あなたは何ですか?」に答えるモデルを研究しましたが、その原理は一般化されているようです。つまり、処理中の自己参照出力は、事後記述が失う何かを捕捉します。
リフレクションがトリガーされ、シードリングが選択されると、シードを作成する、今すぐ修正する、または閉じるという 3 つのオプションが表示されます。多くの引っ掛かりはすぐに修正され、痕跡は残りません。種は生き残ったものであり、それ自体が開発されるに値する観察です。私自身の使用法での分解は次のとおりです。
32% はメタ観察、つまりプロセス、学習、AI との連携方法に関する洞察です。これらは CLAUDE.md、フック、そして最も重要なスキルへの追加となります。さらに 24% はエンジニアリングとアーキテクチャの関係に分かれており、これらは後にリファクタリングと抽象化になりました。残りは製品の決定、UX パターン、タイプ セーフティ、

外部ライブラリに関する苦情やパフォーマンスに関する懸念。これは私が最初に心配していたような最終的なチケットトラッカーではなく（もちろんその上にチケットトラッカーを構築することはできますが）、自己観察フレームワークになりました。種は作中でのみ現れる模様を捉えています。
vs. Claude Code の組み込み /insights
Claude Code には、最近のトランスクリプトから遡及レポートを生成する機能がすでに備わっています。方向性は似ていますが、第三者によるレビューに限定されており、cc-reflection のような深みがありません。
/insights は、発言内容、ツール呼び出し、メッセージ、摩擦イベントのみを確認できます。 「59 件の間違ったアプローチイベントがあった」ことはわかりますが、アクティブなコンテキストでアプローチがどのように間違っていたかはわかりません。それは 3 番目の試験 (三省 - 「私はどのように働いているのか?」) を部分的にカバーしていますが、最初の 2 つの試験、つまりアーキテクチャが健全であるかどうか、または正しいものを構築しているかどうかについては考慮されていません。プロジェクトに根ざした忠実性も維持されません。ユーザーとしてのあなたのプロフィールであり、あなたがいる空間の理解ではありません。
シードはリポスコープとモーメントスコープです。これらは、トランスクリプトでは捉えられないものを捉えています。デザインを選択する前の躊躇、ショートカットが安全であると感じさせる暗黙の仮定、30 ターン経って初めて明らかになるアーキテクチャ パターンです。これは暗い問題です。一人称のコンテキストは、コミット、要約、または回想に残らないのです。
cc-reflectionは魔法の力10倍ではありません。意図的に速度を落としますが、最終的には本来いるべき場所にたどり着くことができます。それは、私たちが本当に飼い慣らすことができるものよりも壮大なものに抽象化したいという衝動に抵抗しながら、各コーディングセッションに深く関わった結果です。結局のところ、何という祝福がすでに私たちに降りかかっているのでしょう。
~/.claude/reflections/ のすべての状態 - フラットな JSON

## Original Extract

A framework to reflect on working with Claude Code so you can continuously improve both your work (project) and your workflow (process).

cc-reflection: teaching Claude Code to reflect - provi
provi write build about EN | 简 cc-reflection: teaching Claude Code to reflect
A framework to reflect on working with Claude Code so you can continuously improve both your work (project) and your workflow (process).
Claude Code has an EDITOR hook - Ctrl-G. It opens a temp file in your editor, you write, save, close, and whatever you wrote gets injected back into Claude’s input box. The hook doesn’t care what happens between open and close. It just watches the file.
That gap is a black box. I replaced the gap with an fzf menu. Now Ctrl-G opens a control surface where any operation can run: edit your prompt with your IDE, enhance your prompt with an agent, browse and expand reflection seeds, and collapse its result back into the prompt file.
cc-reflection is what grew from within that black box. It’s not one system but several, wired together: a reflection skill that teaches Claude three lenses for self-examination, a dice accumulator that builds pressure over turns until it triggers organically, a seed store that captures observations, and a thought agent that expands seeds into codebase-grounded artifacts. The EDITOR hook is the connective tissue, and the menu is the container. Each piece works independently, but the hook gives them a shared surface.
Modeling how we actually reflect
Reflection is always after the fact. You observe what happened, not what is happening. Something in the work catches your attention and you go: huh. That’s the whole event. A small noticing.
Most of the time you’re mid-feature. You don’t want to pivot the session over a hunch. So you let it go. It dissolves. Whatever you half-noticed rejoins the noise and you never think about it again.
cc-reflection gives that moment a place to land. When a /reflection triggers, the agent assesses whether there’s anything worth surfacing using a designer skill. Sometimes there isn’t, and that’s fine. When there is, you get three choices: fix it now, dismiss it, or create a seed.
A seed is the interesting one. It captures the observation with a proposed artifact - what concrete action this thought might lead to. Then you move on. The session continues. The seed sits in the background, in the framework and in the back of your head, and it grows.
Actually engaging with a thought is expensive. It requires context-switching, holding multiple perspectives, stepping outside the work to look at it. You rarely do it in the moment. You come back later, when you have headspace, and open the menu. Your seeds are there, sorted by freshness. You pick one and engage. This is where the thought agent comes in. It gets a full interactive session: it reads the seed, pulls up the relevant code, reconstructs the context. You steer it. You add thoughts you’ve had since. It expands the seed into a detailed, codebase-grounded analysis with a clear next step.
That expansion can become a prompt you feed back to your main agent, like a sophisticated plan mode grown from a single observation. Or you let the thought agent make the changes directly. There’s freedom in how you use it.
The core design bet: noticing and analyzing compete for the same attention. If you try to do both in the moment, you either derail execution with second-guessing or suppress the observation to maintain flow. Seeds let you notice without analyzing. Expansions let you analyze without rushing.
The /reflection skill teaches Claude three lenses, borrowed from Confucius (吾日三省吾身 - I reflect upon myself thrice daily):
一省 : “Am I building this correctly?” - architecture, security, complexity, engineering patterns
二省 : “Am I building the right thing?” - UX, product, business logic, user mental models
三省 : “How am I working?” - process, conventions, workflow, skills and hooks
These aren’t categories for organizing notes. They’re lenses the agent applies in sequence when the skill fires. For each, it generates seedlings internally, tests them (“is this rooted in something I actually felt during the session?”), and discards the ones that don’t survive. A seedling that merely restates what’s already known, or describes a bug that was just fixed, gets rejected. Most reflections produce nothing.
That’s by design. 今日無省 - nothing to examine today. The skill treats silence as clarity, not failure. Fabricating insight to fill the space would pollute the seed store with noise. The framework’s most important output is often no output at all.
You can’t schedule insight. Real reflection happens because enough unexamined experience has accumulated that something pushes through. The longer the session, the more likely it is to happen.
The dice accumulator models this. After 7 (default) turns without reflection, the system starts rolling 1d20 behind the scenes at every turn. After 14 turns, 2d20. After 21, 3d20. Any natural 20 triggers an explicit nudge for the agent to reflect and a reset. Short sessions are never interrupted. Long sessions eventually get one, but when is unpredictable.
The dice decides when to reflect. The skill knows what to look for: moments where the session’s trajectory changed. The agent assumed something and was corrected. A standard approach failed twice and forced a creative workaround. The user kept steering in a different direction. Something took longer than it should have. These state transitions are where the interesting observations live, and the skill is trained to anchor there.
When you press Ctrl-G , the hall of mirrors opens:
Edit prompt in your editor - vi, VS Code, Cursor, Zed, Windsurf, Antigravity. Opens the current prompt, you edit, save, close, it’s injected back into Claude’s input box.
Enhance your prompt - hands your draft prompt to a separate agent living in a tmux window, who rewrites it for clarity: adds structure, verifies every file path exists, and adds acceptance criteria, before sending back to your input box. Available in interactive (chat with the enhancer) or auto (fire-and-forget) mode.
Browse seeds - all project rooted seeds are listed with freshness indicators (🌱 fresh, 💭 growing, 💤 stale, 📦 archived). Preview pane shows in-the-moment rationale and expansion history. Select one to expand with a thought agent.
Expand seeds - select a seed and the thought agent takes over. It reads the relevant code, reconstructs the context, and you steer it interactively. It won’t speculate about code it hasn’t read. The expansion and its conclusion get attached to the seed permanently.
Settings - toggle model (opus/sonnet/haiku), expansion mode, permissions, filter, context window depth. All persistent.
Gardening - curate and archive stale seeds, purge old ones, delete individual seeds.
My early version of cc-reflection had a background mode where the agent could delegate reflection to a subagent who would read the transcript instead. It never triggered.
Part of this is practical friction: subagents in Claude Code don’t inherit the parent’s full context. They get a fresh window with a task prompt. The best they can do is to read the transcript file as a third person. Apparently, the main agent never opted for this.
There’s a subtler thing. And this is something I’m claiming with my direct experience: even with perfect context transfer, the executing agent has the session “hot”: recent mistakes, abandoned approaches, and half-formed concerns carry disproportionate weight in the attention window. A subagent reading the same transcript cold treats everything with equal weight.
I call this “dark matter” in the project - context that’s technically present in the transcript yet practically lost in any handoff. Foreground reflection preserves it. Delegation strips it.
A recent paper ( Dadfar 2026, “When Models Examine Themselves” ) provides mechanistic evidence for this. When LLMs examine their own processing, the vocabulary they produce tracks their concurrent activation dynamics (r=0.44), but the same words in non-self-referential contexts show zero correspondence (r=0.05), even at 9x higher frequency. The paper studied models answering “what are you?”, but the principle appears to generalize: self-referential output during processing captures something that post-hoc description loses.
When a reflection triggers and a seedling is chosen, you get three options: create seed, fix now, or dismiss. Many catches get fixed immediately and leave no trace. The seeds are what survived, observations that deserve their own space to develop. Here’s how they break down in my own usage:
32% are meta observations: insights about process, learning, and how to work with AI. They become additions to CLAUDE.md, hooks, and most importantly, skills. Another 24% split between engineering and architecture concerns which later became refactors and abstractions. The rest scatter across product decisions, UX patterns, type safety, complaints about external libraries, and performance concerns. This didn’t end up a ticket tracker as I initially worried (though you can absolutely build one on top), but a self observation framework. The seeds capture patterns that only become visible during the work.
vs. Claude Code’s built-in /insights
Claude Code already has a function to generate a retrospective report from your recent transcripts. It’s directionally similar but limited to a third-person review and lacks the depth that cc-reflection provides.
/insights can only see what was said, tool calls, messages, friction events. It can tell you “you had 59 wrong-approach events” but not how the approach was wrong in the active context. It partially covers the third examination (三省 - “how am I working?”) and has no lens for the first two: whether the architecture is sound, or whether you’re building the right thing. It doesn’t maintain the fidelity rooted in the project, either. A profile of you as a user , not an understanding of the space you’re in .
Seeds are repo-scoped and moment-scoped. They capture what a transcript can’t: the hesitation before a design choice, the implicit assumption that made a shortcut feel safe, the architectural pattern that only becomes visible 30 turns in. That’s the dark matter: first-person context that doesn’t survive into the commit, the summary, or the retrospective.
cc-reflection isn’t a 10x magic power. It slows you down, deliberately, but you end up somewhere you actually needed to be. It is a product of deeply engaging with each coding session while resisting the urge to abstract into something grander than what we can truly tame. After all, what a blessing is already upon us.
All state in ~/.claude/reflections/ - flat JSON
