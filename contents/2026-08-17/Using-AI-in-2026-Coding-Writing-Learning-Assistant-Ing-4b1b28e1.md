---
source: "https://blog.sshh.io/p/how-i-use-ai-in-2026-coding-writing"
hn_url: "https://news.ycombinator.com/item?id=49331280"
title: "Using AI in 2026 (Coding, Writing, Learning, Assistant-Ing)"
article_title: "How I use AI in 2026 (Coding, Writing, Learning, Assistant-ing)"
image: "https://substackcdn.com/image/fetch/$s_!MDmw!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F4c0a5c31-596a-4a91-b284-19ea9f6ae744_1730x784.png"
author: "sshh12"
captured_at: "2026-08-17T14:18:28Z"
capture_tool: "hn-digest"
hn_id: 49331280
score: 1
comments: 0
posted_at: "2026-08-17T14:11:01Z"
tags:
  - hacker-news
  - translated
---

# Using AI in 2026 (Coding, Writing, Learning, Assistant-Ing)

- HN: [49331280](https://news.ycombinator.com/item?id=49331280)
- Source: [blog.sshh.io](https://blog.sshh.io/p/how-i-use-ai-in-2026-coding-writing)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T14:11:01Z

## Translation

タイトル: 2026 年の AI の活用 (コーディング、ライティング、学習、アシスタント)
記事タイトル: 2026 年の AI の使い方 (コーディング、ライティング、学習、アシスタント)
説明: 「AI の使い方 (2025)」の個人的なアップデート。

記事本文:
2026 年の AI の使い方 (コーディング、ライティング、学習、アシスタント)
Shrivu のサブスタック
2026 年の AI の使い方 (コーディング、ライティング、学習、アシスタント)
「AI の使い方 (2025)」の個人的な更新。
Shrivu Shankar 2026 年 8 月 17 日 シェア AI を効果的に使用する方法を学ぶ最良の方法の 1 つは、「パワー ユーザー」の肩越しに眺め、これらのテクノロジの束を試して、誇大広告と意味のあるものを区別することです。
How I use AI (2025) の 1 年間のフォローアップでは、私が個人的に AI をいじっている最新の方法のスナップショットを撮りたいと思いました。
私はトークンのほとんどをコーディングと研究プロジェクトに費やしています。効果的には、次のようなランダムな質問を受けるだけです。
大勢のエージェントにハッキングを依頼したらどうなるでしょうか?
2026 年以降のフロンティア モデルを使用する最善の方法は何ですか?
私たちはカーバル宇宙計画へのプロンプトにどれくらい近づいていますか?
現在の私のワークフローは、どのプロジェクトでもほぼ同じに見えます。
CONCEPT.md を手書き (~段落) — 理論的な Hacker News タイトル、私のプロジェクトの論文、いくつかの散在する制約
ウルトラ コードの寓話「CONCEPT.md を具体化してください。曖昧な点は何ですか、質問してください。検討していないディメンションは何ですか、必要な API キーは何ですか…」と組み合わせます。
ウルトラ コード フェイブル (またはコーデックス ソル マックス) と組み合わせます。「TECH_PLAN.md に変換します。費やしてもよい金額は次のとおりです。… でホストします。API キーは次のとおりです…」
次に、文字通り「TECH_PLAN.md をビルドして検証してください」というプロンプトを表示し、次の 4 ～ 48 時間かけてすべてをビルドします。
私の典型的なコーディング設定。動的なワークフローを有効にするには、「ultracode」を使用することが重要です。これらの実行の場合:
私は完全にバニラコーデックスとクロードコーデです。カスタムスキル、プラグイン、設定はありません。サイドプロジェクトの場合、これらの機能のほとんどは、エージェントを報酬として使用するための補助輪であると考えています。

プログラミング ワークフロー — 私にとって、これは実際には存在すべきではないコーディング ワークフローです。また、いかなる種類の「サブエージェント ワークフロー」も意図的に設計しているわけではなく、実装プロンプトを起動するときに動的なワークフローに主導権を委ねているだけです。
コードの 95% 以上が最初のメガ ビルド実行で書かれます。左シフトがコードベースのスロップ (つまり SlopCodeBench ) に対する秘密兵器であることを、人々は理解していないと思います。ステップ (4) と同様に、ここでは実際にはバイナリです。ペアリングはなく、端末エージェントの発言を読み取ることさえありません。出力が間違っている場合は、それを完全に破棄し、CONCEPT.md に制約を追加します。多くのバイブコーダーにとって、最初のビルドプロンプトはコードの 5% を書き込みますが、それが実際に問題のほとんどの根底にあると思います。
「TECH_PLAN.md を構築して検証する」という 1 つの段階でプロンプトを表示する意図的な副作用として、プロジェクト履歴全体をハーバー スタイルの評価に変換し、新しいモデルのリリースに対して「実際の作業」をパルス チェックできるようになります。 Codex と Claude Code は現在では十分に近いので、元の実装に何を選択するかをラウンドロビンで決定します。
コードを少し読んでみました。多くの場合、形状 (つまり、ファイル ツリー) とエントリ ポイントです。コアアルゴリズムがある場合は、ソースを徹底的に調べるのではなく、.html の説明を求めます。もし私がコードを掘り下げることになったとしたら、それは実装に何らかの「不正行為」があったのではないかと疑うからです。
最終出力の書かれたテキストについては、計画内で任意の単語数を設定します。 「このアプリ全体には、ユーザー向けの単語が 500 語しか含まれていない可能性があります。」これは、内容を読みやすくするための最も効果的な方法であると思います（簡略化された英語や「簡潔にする」プロンプトと比較して）。
私は通常、午前 7 時頃に実装プロンプトを起動し (仕事中に実行させておきます)、

d 午後9時（私が寝ている間）。コーディング エージェントは常に自動モードに設定されており、技術計画は通常十分に明確であるため、中間ステップで人間による検証は必要ありません。クロード/コーデックスの「リモート コントロール」機能はあまり使用しません。中間出力でペアリングする必要があるのは私にとってアンチパターンだからです。
これらのプロジェクトの結果は、多くの場合、洞察や、もしもの質問に対する答えになります。コードやアプリの URL を共有することに意味があるとはほとんど考えられません。代わりに、私は通常、ループ全体とその成果物は一時的なものであると考え、X に関する洞察を共有するか、ブログ投稿で共有するだけです。
私は 3 つの異なるマシンタイプを使用しており、1 ～ 10 個のエージェント CLI 端末を同時に実行しています。
ゲーム PC (Nvidia 5090、Windows + WSL v2) — ML/RL 研究およびゲーム/グラフィックス関連プロジェクト用
Mac Mini — 日常的なプロジェクトのほとんどに。最も近いデバイスからクラウドフレアトンネルを介して SSH 接続します。
モーダル関数 — 極めて並列な CPU コンピューティングまたは大規模な GPU 研究プロジェクト用。多くの場合、PC 上で高速にスケールダウンした反復を実行し、最終的な $$$ 実行のためにクラスターにスケールアウトします。
Shrivu のサブスタックをお読みいただきありがとうございます。無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
私たちは、企業およびソーシャル AI のピークの時代に突入しました。私は、AI を使用して高品質のコンテンツを作成できると強く信じていますが、ほとんどの場合、最終的にはそうではないという現実にこの 1 年でより根ざしたものになりました。その結果、「この文章はAIによって書かれたのか」ということは事実上、「文章を書くために何らかの努力が払われたのか」と同じことになってしまいました。残念ですが、わかります。プラスの面としては、タイプミスや貧弱な文法が（限定的な範囲で）スタイルとして戻ってきたと思うので、個人的には、次のようなプレッシャーをあまり感じなくなりました。

「完璧な」テキスト。
その結果、人に向けた文章については、GenAI 以前のレベルの AI のタイプミスとサブセンテンスの文法チェックに戻り、自分が書いたものに労力が費やされたかどうかについて曖昧さがなくなりました。実際のところ、手入力にそれほど時間はかかりませんが、自分の文章が以前と同じようにうまくまとまっており、読者に最適であるという「確信」が少し薄れているだけです。現時点では、これは「人間が書いた」パングラムバッジと交換する価値があります。
私の最新のサブスタック投稿。 100% 認定された人間が書いたものです。誰もがメモを受け取るわけではありません。私は、(仕事でもそれ以外でも) ライティングや AI の使用に関する期待を設定することに、より快適になっていると感じています。 AI の使用を決して恥じることはありませんが、肥大化したテキストやレビューされていないテキストは AI の悪い使用法であり、読んで楽しくないことを明確に示しています。
私は .html ファイルを使って物事を学ぶことに夢中です (「HTML の理不尽な効果」を参照)。 (X、研究室/スタートアップのブログ投稿、またはハッカー ニュースを通じて) トピック、本、研究論文を見つけて、それらを「インタラクティブな遊び場」に変換します。通常:
X に関する注目の新しい研究論文をご覧ください
要約をざっと読んで、全文をクロード/コーデックスに放り込みます。「ここで何が目新しいのかを説明するために、インタラクティブなプレイグラウンド アーティファクトを構築します。私は技術者で、… についてはすでに知っていますが、… についてはあまり詳しくありません。」
更新された .html ファイルを生成するフォローアップの質問をいくつか行います。(3) に進みます。
これは技術的なトピックを学習するのに最適ですが、私は今でも時事問題 (例: インタラクティブな地図/デジタル博物館) や非技術書籍 (例: 構造化され、逐語的に開示される内容の各章を再フォーマットしたもの) に対して試してみることがよくあります。大学時代に受けた講義のほとんどはもっと効果的だったかもしれないとまで言いたいです（個人的には）

) は、巧みに作成されたインタラクティブな .html ファイルとして保存されます。
最近、バッチ推論のための GPU メモリ割り当てに興味があり、Claude にこの説明器を構築してもらいました。私は、ノブが何をするかを予測し、実際に何が起こるかを確認するためにノブをいじってみるのが、非常に厄介な学習戦略であると感じています。練習すれば、興味のあるどんなトピックでも習得できるようになります。急速に進化する AI コミュニティの多くが X 上に存在するため、私はカスタム CLI (動的ワークフローで使用) 経由で Grok X Search API をよく使用します。これは、あるトピックに関する先行技術を深く調査したり、信号の高い人々がフォローしたりするためです (興味深い事実: X API を直接使用するよりも、Grok 経由の方が 10 倍安価です)。
個人背景アシスタント
OpenClaw に関する誇大宣伝は少し下火になりましたが、自律型パーソナル アシスタントはかつてないほど優れており、安価になっています。
ここでも私はほとんどバニラです。既存の Claude サブスクリプションを使用して、Mac Mini に ssh で接続し、tmux セッションを開き、次のように Claude Code を起動するだけです。
$ claude --dangerously-skip-permissions “/start-ops-team”
ここで、「/start-ops-team」はカスタム スキルです。
「/start-ops-team」は、エージェントにいくつかの動作原則と、生成するサブエージェントとともに使用するためのローカル マークダウン ディレクトリ レイアウトを教えます。クロード コードの組み込み「/loop」を多用して、数週間にわたって継続的に実行し続けます。
私は効率的な並列ブラウザ自動化のために brw を使用しています。私がやりたいことのほとんどには MCP がなく、従来のブラウザを使用するとかなりコストがかかるか大ざっぱなので、エージェントが使用できるようにこれを構築しました。
私はカスタム WhatsApp プラグインを使用して、WhatsApp から直接チャットできるようにしています。私のアシスタントにも独自の実際の電話番号が設定されています。これは、ニッチだが非常に強力な「チャネル」MCP 機能を使用します。
信じられない

個人的な「司令センター」、またはジャービスのようなアシスタント。代わりに、私はバックグラウンドでエージェントを徹底的に使いこなしており、私がいなくてもアシスタントが実行できるタスクに集中させています。私は文字通り、緊急の例外がない限り、多くても週に 1 回連絡するように促しました。私もすぐに通知疲れを感じます。タスクには次のものが含まれます。
自動支払いを行わずに定期的な請求書を支払い、経費の領収書を転送します。
ソーシャルメディアインバウンドへの対応。特に、見知らぬ人を調査し、スケジュールされたコーヒーチャットやその他の直接チャネルに優先順位を付けることで、LinkedIn DM を監視します。アシスタントは実行中のランブックから対応方法を取得し、エッジケースに遭遇した場合には週次メッセージでエスカレーションします。私にとって重要なのは、アシスタントが実際には私ではないことを知らずにアシスタントと長々と会話をしている人がいないことです。そのため、アシスタントは適切なチャネルへのトリアージに大きく方向転換されています。
情報源として私を登録し、Google カレンダーを同期します (例: イベントに招待される → 行きたいと十分な確信を持って決定する → 登録する + カレンダーを保留して更新する)。これらは多くの場合、私が過去に会ったことがある人々からの出来事であり、アシスタントはそれを知っています。ヘアカットやその他の同様の形の定期的な予定も好きです。
AI 主導の LinkedIn エクスチェンジ。私が実際に見たのは、この人が誰なのか、そしてチャットに役立つ内容についてのコンテキストが含まれる金曜日の Google カレンダーの最後のイベントだけでした。アシスタントはスクリーニング後にメッセージの約 90% を無視し、10% のほとんどはカレンダーのリンクを提供します。 AI によって生成された応答は、事前に承認された簡潔な応答のランブックによって制限されます。コスト
奇妙なことに、今は 1 年前よりも月あたりの支出が減りました (800 ドル→ 500 ドル)。それは私が Anth だけに統合したことによって完全に推進されました

Ropic と OpenAI のサブスクリプションと、そこから得られる信じられないほどの使用量。これまでのコストの多くは API トークンの請求から来ており、現在はこれらのサブスクリプションを通じて戦術的にルーティングしています。私のナプキンの計算によると、ナプキンなしの場合、現時点での実際の使用コストは月あたり約 6,000 ドルになります。
Google AI Pro ($20/月) — GSuite の特典を備えた便利な AI ファミリー プラン
モーダル、鉄道、Netlify ($20-500+/月) — 実験のホスティングまたは実行用
削除: イレブンラボ、Suno、Cursor、Vast.ai、Perplexity、Gemini Ultimate
Fable/Sol ウルトラ モードの最大化にもかかわらず、通常、毎月の最大プランにはまだ少し余裕があります。 ChatGPT Pro の制限に達したことはありませんが、アイデアが多い週には Fable が足りなくなってしまうことがよくあります。
最後に、AI を最大限に活用するための私の最新の推奨事項をご紹介します。
チャットベースのアシスタントのような AI の使用をやめます。ほとんどの作業が最初のプロンプトで完了するように左にシフトし、自分を副操縦士というよりもマネージャーであると考えてください。中間のチャット メッセージではなく、結果を確認します。私がこれまでに行ったペア プロンプト セッションで最もよく見られた間違いは、完全な目標をそのまま文書に移し、エージェントにそこから（中断せずに！）調理させるだけではなく、より大きな目標を達成するために狭いタスクをチャット セッションに少しずつ持ち込んでしまうことです。
自分自身の AI の野心とスキルをスコアリングするための代理としてフロンティア モデルを使用します。 「AI は頭打ちになった」、あるいは

[切り捨てられた]

## Original Extract

A personal update to "How I use AI (2025)".

How I use AI in 2026 (Coding, Writing, Learning, Assistant-ing)
Shrivu’s Substack
Subscribe Sign in How I use AI in 2026 (Coding, Writing, Learning, Assistant-ing)
A personal update to "How I use AI (2025)".
Shrivu Shankar Aug 17, 2026 Share One of the best ways to learn how to use AI effectively is just to look over the shoulder of a “power user” and play with a bunch of these technologies to tease out what’s hype vs what meaningfully sticks.
In this one-year follow-up to How I use AI (2025) , I wanted to snapshot the latest ways I’m messing with AI personally.
I spend most of my tokens on coding and research projects. Effectively just taking random questions like:
What would happen if I asked a bunch of agents to hack me?
What’s the best way to use 2026+ frontier models?
How close are we to prompt-to-Kerbal Space Program?
My workflow right now looks nearly identical for every project:
Hand-write (~paragraph) a CONCEPT.md — the theoretical Hacker News title, my project thesis, some scattered constraints
Pair with ultra code fable “Flesh out CONCEPT.md, what’s ambiguous, ask me questions, what are dimensions I’m not considering, what API keys do you need…”
Pair with ultra code fable (or codex sol max) “Convert to TECH_PLAN.md, here’s how much I’m willing to spend, host on …, here’s some API keys …”
Then I will literally just prompt “Build and verify TECH_PLAN.md” and over the next 4-48 hours I’ll let it build everything out.
My typical coding setup. It’s critical to use “ultracode” to enable dynamic workflows. For these runs:
I’m completely vanilla Codex and Claude Code. No custom skills, plugins, or settings. For side-projects, I see most of those features as training wheels for using these agents as pair programming workflows — which to me is a coding workflow that shouldn’t really exist anymore. I’m also not intentionally designing any sort of “subagent workflows” and just letting dynamic workflows take the wheel when I fire off the implementation prompt.
95%+ of the code is written in that first mega build run. I don’t think folks appreciate how much shifting left is the secret weapon against codebase slop (i.e. SlopCodeBench ). Like step (4) really is binary here — there’s no pairing or even reading what the terminal agent says. If the output is wrong, I throw it completely away and add constraints to the CONCEPT.md. For many vibe coders out there, the first build prompt writes 5% of the code and I think that actually underlies most of their issues.
An intentional side-effect of prompting with a single stage “Build and verify TECH_PLAN.md” is that I am also turning my entire project history into harbor-style evals which allow me to pulse check “real work” against new model releases. Codex and Claude Code are close enough now that I’ll round-robin what I pick for the original implementation.
I read the code a little bit. Often the shape (i.e. file tree) and entry points. If there’s some core algorithm, I’ll ask for a .html explainer rather than digging through the source. If I do end up digging into the code, it’s because I suspect some sort of “cheating” in the implementation.
For any written text in the final output I set arbitrary word counts in the plan. “This entire app may only have 500 user-facing words”. I find this to be the most effective way to keep things readable (vs simplified English or “be concise” prompts).
I fire off the implementation prompts usually around 7 am (letting them run while at work) and around 9 pm (while I’m sleeping). The coding agents are always set to auto-mode and the tech plan is usually clear enough that there’s no human-verification required at intermediate steps. I don’t really use claude/codex ‘remote control’ features that much because to me it’s an anti-pattern to need to pair on intermediate outputs.
The outcome of these projects is often an insight or the answer to the what-if question. Rarely does it make sense for me to share the code or even the app URL. Instead I typically consider the entire loop and its artifacts ephemeral and just share the insight on X or with a blog post.
I use three different machine types, with one to ~ten agent CLI terminals running at the same time:
A gaming PC (Nvidia 5090, Windows + WSL v2) — for ML/RL research and gaming/graphics-related projects
A Mac Mini — for most day-to-day projects. I ssh over a cloudflared tunnel from whatever device is closest to me.
Modal functions — for extremely parallel CPU compute or for big boy GPU research projects. Often doing fast scaled-down iteration on my PC and then scaling it out to a cluster for a final $$$ run.
Thanks for reading Shrivu’s Substack! Subscribe for free to receive new posts and support my work.
We have entered an era of peak corporate and social AI-slop. I firmly believe that you can use AI to write high-quality content but have over the last year become more grounded in the reality that most of the time that’s not what ends up happening. As a result “was this text written by AI” has de facto become the same as “was any effort put into the writing”. It’s unfortunate, but I get it. On the plus side, I think typos and poor grammar (to a limited extent) have come back into style so I do personally feel much less pressure to have “perfect” text.
So as a result, for human-facing writing, I’ve gone back to pre-GenAI-level AI typo and sub-sentence grammar checking so there’s no ambiguity as to whether what I wrote had effort put into it. Hand typing really doesn’t take that much more time though I do just feel slightly less “sure” that my writing is as well synthesized and audience optimal as before. At this point, it’s a worthy trade-off for the “human-written” Pangram badge.
My most recent Substack post. 100% certified human written! Not everyone gets the memo. I do find myself getting more comfortable setting writing and AI-use expectations (at work and outside of it). Never shaming someone for using AI but explicitly making it clear that bloated and/or unreviewed text is a bad use of AI and is not enjoyable to read.
I’m obsessed with learning things with .html files (see The unreasonable effectiveness of HTML ). I’ll discover (through X, lab/startup blog posts, or Hacker News) some topic, book, or research paper and just convert them into “interactive playgrounds”. Typically:
See hot new research paper on X
Skim the abstract, throw the full text into Claude/Codex, “build an interactive playground artifact to explain what’s novel here, I’m a technical person who already knows …, I’m less familiar with …”.
Ask some follow-up questions that generate an updated .html file, go to (3)
This works best for learning technical topics though I’ll often still attempt it for current events (e.g. an interactive map/digital museum) and non-technical books (e.g. re-formatted as structured, progressively disclosed chapters of the verbatim content). I would go as far as saying that most of the lectures I sat in during college could have been more effective (personally) as a well-crafted interactive .html file.
Recently I was curious about GPU memory allocation for batch inference and had Claude build this explainer. I find making predictions about what the knobs will do and then playing with the knobs to see what actually happens to be a very sticky learning strategy. With practice I feel like I can knob-ify any arbitrary topic I’m interested in learning. As more of the rapidly evolving AI community sits on X, I also use the Grok X Search API via a custom CLI (used with dynamic workflows) quite a bit for deep researching prior art on some topic or for high-signal folks to follow (fun fact: it’s 10x cheaper via Grok than the X API directly).
Personal Background Assistants
While the hype around OpenClaw has died down a bit, autonomous personal assistants are better and cheaper than ever.
I’m mostly vanilla here as well. Using my existing Claude subscription, I ssh into my Mac Mini, open a tmux session , and just launch Claude Code like this:
$ claude --dangerously-skip-permissions “/start-ops-team”
Where “/start-ops-team” is a custom skill.
“/start-ops-team” teaches the agent some operating principles and a local markdown directory layout for it to use along with the subagents that it might want to spawn. It makes heavy use of Claude Code’s “/loop” built-in for keeping it running continuously for weeks.
I use brw for efficient parallel browser automation. Most of the things I want it to do don’t have an MCP and traditional browser use is pretty costly or sketchy so I built this for my agents to use.
I use a custom WhatsApp plugin to let me chat directly from WhatsApp. My assistant has its own real phone number set up as well. This uses a niche but very powerful “channels” MCP feature.
I don’t believe in personal “command centers” or Jarvis-like assistants. Instead I’m extremely background agent-pilled and focus my assistant on tasks it can do without me in the loop. I’ve literally prompted it to contact me at most once a week unless there’s an urgent exception. I also just get notification fatigue super easily. Tasks include:
Paying recurring bills without auto-pay and forwarding the receipts for expenses.
Responding to social media inbounds. Particularly sussing out LinkedIn DMs by researching and triaging strangers into scheduled coffee chats and other direct channels. The assistant pulls from a running runbook for how to respond and escalates in the weekly message when it hits edge cases. It’s important to me that folks aren’t having drawn-out conversations with the assistant not knowing it’s not really me so it’s steered heavily towards triaging to the right channel.
Signing me up for stuff and syncing my Google Calendar as my source of truth (e.g. I get invited to an event → It decides with enough certainty I’d want to go → signs me up + updates my calendar with a hold). These are often events from folks I have met up with in the past and the assistant knows that. Also like haircuts and other similar-shaped recurring appointments.
An AI-driven LinkedIn exchange. All I actually saw was the final Friday Google Calendar event with context on who this person was and what might be useful for me to chat on. The assistant ignores ~90% of messages after screening with most of the 10% getting served my calendar link. AI-generated replies are limited by a runbook of succinct pre-approved responses. Costs
Weirdly enough, I spend less now than I did a year ago per month ($800 → $500). That’s completely driven by me consolidating into just the Anthropic and OpenAI subscriptions and the incredible amount of usage you can get out of them. A lot of my historical costs came from API token billing which I also now tactically route through these subscriptions. My napkin math indicates my actual usage cost would be around $6,000/mo at this point without them.
Google AI Pro ($20/mo) — a handy AI family plan with GSuite benefits
Modal, Railway, Netlify ($20-500+/mo) — for hosting or running experiments
Dropped: Elevenlabs, Suno, Cursor, Vast.ai, Perplexity, Gemini Ultimate
Despite Fable/Sol ultra mode maxxing, I typically still have a bit of wiggle room in the max plans each month. I’ve never hit my ChatGPT Pro limit while I do regularly run out of Fable on idea-heavy weeks.
I’ll end with my latest recommendations for getting the most out of AI:
Wean off of using AI like a chat-based assistant. Shift-left so that most of the work is done in your first prompt and think of yourself as more of a manager than a co-pilot. Review results, not intermediate chat messages. In pair-prompting sessions I’ve done, the most common mistake I see is folks trickling narrow tasks into the chat session to accomplish a larger goal rather than just shifting left the full goal into a document and just letting the agent cook (without interruption!) from that.
Use frontier models as a proxy for scoring your own AI ambition and skill. I know it’s very popular to claim “AI has plateaued” or that the

[truncated]
