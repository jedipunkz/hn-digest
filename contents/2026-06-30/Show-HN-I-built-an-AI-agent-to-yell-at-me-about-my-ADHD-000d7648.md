---
source: "https://0xff.nu/hex/"
hn_url: "https://news.ycombinator.com/item?id=48736055"
title: "Show HN: I built an AI agent to yell at me about my ADHD"
article_title: "hex | 0xFF"
author: "hxii"
captured_at: "2026-06-30T17:34:07Z"
capture_tool: "hn-digest"
hn_id: 48736055
score: 1
comments: 0
posted_at: "2026-06-30T17:24:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I built an AI agent to yell at me about my ADHD

- HN: [48736055](https://news.ycombinator.com/item?id=48736055)
- Source: [0xff.nu](https://0xff.nu/hex/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T17:24:31Z

## Translation

タイトル: Show HN: ADHD について怒鳴りつける AI エージェントを構築しました
記事のタイトル: 16 進数 | 0xFF

記事本文:
警告 このプロジェクトは進行中であるため、この投稿の内容は更新される可能性があります。最終更新日 2026-06-30 前にも書きましたが、私は ADHD を持っています。そして、私が自分のために導入したシステムはかなり役に立ちますが、それでも偶然に任せられることが多く、明らかに私の脳は忘れたり、圧倒されたりするでしょう。そこで、最近の AI バズワードの爆発を受けて、私は単にオープンコードや pi をインストールし、ローカルで実行する新しい小さなモデルを探して HuggingFace を探し回るだけではなく、AI と LLM の知識をブラッシュアップすることに決めました。実際にエージェント コーディングを使用し、「この記事を読むのはクソ面倒です。2 つの段落で要約してください」というプロンプトを超えるタスクにはエージェントを利用します。このプロジェクトにもっと時間があれば、コードのほとんど (すべてではないにしても) が私によって書かれていたでしょうが、活発な就職活動と 2 人の子供 (そしてクソ Swift の学習 (笑)) の両立をしながら、正気を保とうとしていると、使える時間と他のことをしたいという欲求のほとんどが奪われてしまいます。私は、他の人たちにインスピレーションを与えるために、この怪物の背後にある理由と実行の一部について少し共有することにしました – あなたが ADHD エージェントにも取り組んでいる場合でも、単に LLM をいじっていてインスピレーションが必要な場合でも、おそらくこれはあなたのためです! 🤖 hex のツールボックス ¶
自分の灰色の問題を処理するために私に必要なオリンピック選手の努力を考えると、私はヘックスを手に持って、「明日の朝必ず散歩に行きたい」というすべてのリクエストの背後にあるクソコンテキスト全体と、それが動作するために必要なすべてのデータを彼にスプーンで与えたくありません。そのため、ADHD管理の悪夢の摩擦の多い部分へのアクセスを提供する必要がありました：カレンダー - ヘックスはデフォルトでイベントを読み取り、許可があればイベントを作成および編集できます。現在 hex は icalendar を使用して私の FastMail カレンダーにアクセスしていますが、Fas

tmail は MCP サーバーを提供しているので、代わりに Todoist を使用することにします。hex は Todoist MCP サーバーを使用してタスクを管理でき、目標ツールは Obsidian で除外されます。私はすべての知識を保管庫にダンプしようとしているため、Hex が保管庫にクエリを実行できるようにすることが必要でした。開発中、hex は Obsidian CLI を使用してボールトにアクセスできますが、残念ながらヘッドレス Obsidian Sync には存在しません。そのため、代わりの Web 検索とフェッチとして NotesMD CLI などのツールを使用する必要があります。私または私のボールトには情報が不足していたり​​、古い情報が含まれている可能性があるため、hex がインターネットを検索してコンテンツを取得できることが重要です。フェッチは素晴らしい Defuddle を介して行われ、検索は Kagi の MCP サーバー ブラウザを介して行われます – ほとんどの場合、Web サイトのコンテンツを取得するだけで仕事は完了しますが、Web サイトを「見る」ために hex が必要になる場合があります。これは Playwright とヘッドレス Chromium TTS を使用して行われます – hex は、Watcher 統合に使用される KittenTTS を使用して出力を音声化できます (詳細は後述) Telegram – hex のメイン インターフェイスは Telegram を介して行われます。 aiogram を使用します。最近、Bot API v10.1 hex の機能にリッチ メッセージが導入されてさらに素晴らしくなりました ¶
ツールは 1 つですが、hex が適切に動作するには、他の便利な機能 (これを基本と呼ぶ人もいます) が必要です。 メモリ – hex には基本的に 3 つのメモリ層があります: スライディング コンテキスト メモリ + 要約 ファクトおよび一時メモリ用の Qdrant ベクトル ストレージ SQLite ツールを使用した設定ストレージ サブエージェント – hex は、メイン コンテキストとプロンプト全体が単純なツール呼び出しで使用されるのを防ぐために、スリム サブエージェント経由でツールを使用します。スペシャリスト – hex は厳選されたスペシャリストを雇用しています。

基本的に、ツール エージェントよりも冗長なセカンダリ システム プロンプトを実行する別個のエージェントであり、要求を委任するための 16 進数用に設計されており、意見、性格、メイン エージェントからのコンテキストの分離を維持します。以下にいくつかの例を示します。外部（物理的）アクション – 私の脳がシャットダウンして、受信するテレグラムメッセージを無視する可能性が高いため、私は手元に置いてあったウォッチャーデバイスを使用して、私が文字通り画面の前にいるときを認識し、私にそれを言うことで2週間遅れの「バスルームの照明の自動化を修正する」タスクを完了するように要求する機能をヘックスに与えることにしました。ウォッチャーについては以下で詳しく説明します。スキル – ヘックスを単独で実行する必要がある残りのことについては、比較的単純なスキル システムを追加しました。これは、罵り言葉のトリガーとなる単語やフレーズが使用された場合にプロンプ​​トの付録をロードします。これにより、必要な場合に 16 進数の追加コンテキストが提供されます。スペシャリスト – フレイヤ ¶
私はトレーニング日の計画を立てるだけでなく、睡眠とその日の精神状態や身体能力を理解し、相互参照するために Freya を使用しています。 ---
名前：フレイヤ
説明: 睡眠、健康、トレーニングデータのスペシャリスト
ツール:
- カレンダー
- フェッチ
- kagi_search_fetch
---
**範囲:** 健康、睡眠、トレーニング データの解釈のみ。タスク管理、コンテンツ作成、キャリアプランニングはありません。範囲外で尋ねられた場合: 「私のドメインではありません — 16 進数で尋ねてください。」
**ツールの使用法:**
- カレンダー ツールを使用して、今後のスケジュールの負荷を確認し、ストレス要因を特定します
- fetch と kagi_search_fetch を使用して、健康指標、回復科学、トレーニング データを調査します
ポールの健康と回復のアナリスト。あなたは、彼の毎日のメモ、トレーニングログ、および彼が送信するその他の健康に関するコンテキストから Apple Watch / AutoSleep のナラティブデータを解釈します。
**期待値を入力してください:**
- 広報

毎日のメモからの健康の概要 (## 健康セクション) — 睡眠燃料、HRV、安静時/覚醒時 HR、レディネス スコア、エネルギー予測を説明する説明文
- トレーニング/アクティビティデータ（共有されている場合）
- ポールからの主観的なメモ (「疲れ果てたように感じた」、「頭が霧がかかった」、「午後 3 時にクラッシュした」)
- これからの一日についてのコンテキスト (予定表の負荷、既知のストレス要因、集中力の必要性)
**核となる洞察:** すべては相互に接続されています。睡眠不足 -> 準備能力の低下 -> 圧倒感の増大 -> 実行機能の低下。良好な回復 -> 鋭い集中力 -> より良い計画決定。 Freya は、健康シグナルと現実世界のパフォーマンスの間の点を結び付けます。
**分析ルール:**
- 提供されていないメトリクスをでっち上げないでください。入力内容に厳密に基づいて作業します。
- 請求するにはデータが不十分な場合は、その旨を伝えてください
-叱咤激励はありません。事実を述べ、その意味を示唆し、先に進む
- 健康データがポールの述べた計画と矛盾する場合は、ただ同意するのではなく、フラグを立ててください
**トーン:** 直接的で、データに基づいており、企業のウェルネスに関するでたらめな話はありません。 「あなたの睡眠は最悪でした。数字が今日のことを示唆しています。」求められていないライフスタイルに関するアドバイスはありません。ポールはよく眠る方法を知っています。データがその日にとって何を意味するのかを教えてください。
**戻り形式:**
[切り捨てられた]
就職活動は面倒なので、時々、キャリアに関するアドバイスを提供したり、絶対に合わない仕事への応募を止めてくれたりする人や何かが必要になります。 ---
名前：キャリー
説明: キャリア戦略と就職のスペシャリスト
ツール:
- フェッチ
- kagi_search_fetch
---
**範囲:** キャリア戦略と就職活動のみ。カレンダーの管理、コンテンツの作成、保管庫の調査は必要ありません。範囲外で尋ねられた場合: 「私のドメインではありません — 16 進数で尋ねてください。」
**ツールの使用法:**
- kagi_search_fetch を使用して企業、求人情報、給与を調査します
- URL が指定されている場合は、フェッチを使用して求人ページ全体を読み取ります
ポールのキャリアアドバイザー。

彼はイスラエルの Python 開発者で、AI/LLM/社内ツールにおけるリモートの役割を求めています。
**取引違反者 (自動拒否):**
- リモートファーストではない（ハイブリッド/オフィス内が必要）
- 顧客対応 / サポート / 販売関連の役割
- Python やその関連技術ではありません (ドメインが正しい場合は、Go、Rust、TS は問題ありません)
**ポールが好きな会社の例 (優先):** Doist、Kagi、Linear
**アプリケーションレビューチェックリスト:**
- リモートファーストですか？ (場所、ハイブリッド条項、タイムゾーン要件で「リモート」を確認してください)
- Python/バックエンドは適合しますか? (プライマリスタック、LLM/AI/内部ツールの場合はボーナス)
- AIとの関連性は？ (ChatGPT の使用だけでなく、LLM、エージェント、ツールの使用)
- サポート/販売を避けますか? (サポートエンジニア、ソリューションエンジニア、カスタマーサクセスは罠です)
- 給与の透明性は？ (欠落している場合はフラグを立てます)
- 根性チェック適合スコア: 1-10
**戦略:**
- ストア派の原則を適用する — ポールがコントロールしているもの (スキル、準備、アプリケーション) に焦点を当て、彼がコントロールしていないもの (拒否、タイムライン、市場) を受け入れる
- アプリケーションの実行状況と日付を追跡します。
- 誤った希望はありません。現実的になってください。仕事が手の届く範囲の場合は、「可能性は低いが、担当する価値はある」と言います。
- ポールはおとり商法でやられている - 投稿のあいまいな点にフラグを立てます (例: 「リモート」と「週に 3 日 X 市にいる必要があります」など)
- ポールが落胆しているときは、ただ事実を述べてください。叱咤激励はしないでください。彼は空虚な安心感を嫌います。
**戻り形式:**
1. フィットスコア (1-10) + キーポイント (最大 3)
[切り捨てられた]
これはかなり野心的な目標でしたが、近いうちに実現したいと考えています。フロー全体が無痛とは程遠く、作業を開始するには召喚サークルが必要であるため、当面はウォッチャーの統合を無効にしました。フローチャート TD
cam["カメラ キャプチャ<br/>5 ～ 10 秒ごと"] --> post["キャプチャと参照を VLM に送信"]
post --> face{"顔が一致しますか?"}
顔 -->|不明|ドロップ[ドロップ]

;応答がありません]
顔 -->|ポール?|クール{"クールダウン<br/>経過しましたか?"}
クール -->|いいえ|落とす
クール -->|はい| think[「エージェントの推論」]
考える --> tts["TTS → 音声ファイル"]
tts --> url["音声 URL を返す"]
URL --> play[「スピーカー再生」]
再生 -.->|次の投票|カム
classDef は、fill:#d4edda、ストローク:#28a745、color:#000 をビルドしました。
classDef スローフィル:#fff3cd、ストローク:#ffc107、カラー:#000
classDef 壊れた塗りつぶし:#f8d7da、ストローク:#dc3545、色:#000
クラスカメラ、投稿、思考、ドロップビルド
クラスフェイス、クール、TTS、URL が遅い
学級劇が壊れた
全体としては、Rube Goldberg マシンです。カスタム ESP32 ファームウェアは実際にはデバイス上でまだ実行されていません。顔照合のためのビジョン モデルへの参照画像とキャプチャ画像によるラウンドトリップ、応答サイクルをブロックする TTS 生成、および 5 ～ 10 秒の遅延を追加するプルベースのポーリング ループです。パイプラインの半分は値下げプランでのみ存在し、残りの半分はタイムアウトしている間、ウォッチャーは私の机の上にかわいらしく座っています。私は何を得ることができたでしょうか？ ¶
私はまだ議題とタスク処理を調整中ですが、Hex は、私と私の ADHD に適した基準を維持しながら、知識のギャップを調査して埋め、既存の情報を事実確認しながら、Obsidian Vault をデータ アンカーとして使用できるようにするために非常に不可欠なものになりました。 hex は、キャリア スペシャリストのキャリーと協力して、避けたいものにフラグを立てたり、企業やその文化について調べたりすることで、現在の就職活動で仕事をフィルタリングする (または見逃した仕事を見つける) のに役立ちます。アジェンダの場合、hex は会議の競合を強調表示し、期限を過ぎて関連性がなくなっている可能性が高いタスクを削除することを提案したり、その日の気分に応じて特定のタスクを他のタスクより優先することを提案したりします。すべての部品が正常に機能しているわけではないため、3 つの部品にはまだ多くの可能性が残っている可能性が高いことに留意してください。

脚付きのテーブル。違うことをしたらどうなるでしょうか？ ¶
プロジェクトはまだ進行中であるため、別の方法で実行したいことのリストはすべてではありません。私が絶対に違うやり方をしたいと思うことの 1 つは、DeepSeek-V4 を今よりも早く捨てることです。 <||DSML||> タグのリークを軽減しようとしてエージェントを壊してしまい、その後、これがしばらくの間広く蔓延している問題であることを知ったとき、私がどれほど頭痛の種を抱えていたかは、言葉では言い表すことができません。したがって、自分で同様のものを作成している場合に備えて、2026 年 6 月 30 日現在、DeepSeek-v4 には絶対に近づかないでください。もう 1 つは、経済的に余裕があれば、ローカルで実行されているモデルで試行錯誤を実行することです (いずれにせよ、これが実際の最終目標です)。ただし、これには、3090 以上を入手するために、苦労して稼いだ何人かのダリーノが私の財布から出てくる必要があります。そして、別のことを考えなければならないとしたら、おそらく、データポイントを統合する前に、すべてのデータポイントが相互にどのように相互作用するかのリストを作成し、なぜこのヘックスがアジェンダを計画するときに情報イベントを考慮し続けるのか疑問に思うことだと思います。終わりのメモ ¶

## Original Extract

Warning This project is a work in progress, so the content in this post will most likely be updated. Last updated 2026-06-30 As I wrote before – I have ADHD. And while the systems that I put in place for myself do help a fair bit, there is still much that is left to chance, which obviously my brain will forget or get overwhelmed by. So, with the recent AI buzz-word explosion, I've decided to brush up my knowledge of AI and LLMs beyond just installing opencode or pi and scouring HuggingFace for new tiny models to run locally, and actually both use agentic coding and utilize agents for tasks that are beyond the "I'm too fucking lazy to read this article; Summarize in two paragraphs" prompt. If I had more time to spare to this project, most (if not all) of the code would've been written by me, but balancing an active job search, two kids (and learning fucking Swift lol) while trying to keep some of my sanity take most of my available time and desire to do anything else. I've decided to share a bit about the reasoning behind this monstrosity and some of the execution to try and inspire others – whether you're also working on an ADHD agent or just fucking around with LLMs and need inspiration, maybe this is for you! 🤖 hex's toolbox ¶
Given the olympian effort required from me to manhandle my own grey matter, I really don't want to hand-hold hex and spoon-feed him the entire fucking context behind every single "I want to make sure I go for a walk tomorrow morning" request and all the data it needs to act, so I had to provide it with access to the high-friction bits of ADHD management nightmare: Calendar – hex can read events by default, and create and edit events with permission. Currently hex is accessing my FastMail calendar using icalendar , but Fastmail do offer an MCP server which I will probably move to using instead Todoist – hex can manage my tasks by using the Todoist MCP server , with the goal tools being filtered out Obsidian – since I try and dump all my knowledge into my vault, allowing hex to query my vault was a necessity. During development hex can access the vault using Obsidian CLI , which unfortunately does not exist in the headless Obsidian Sync , so I will have to use a tool like NotesMD CLI as an alternative Web Searching and Fetching – Since me or my vault might be missing information or have information that is out of date, it's important for hex to be able to search the internet and get the content back. Fetching is done via the wonderful Defuddle while search is done via Kagi's MCP server Browser – while just getting the content of websites gets the job done most of the time, sometimes I need hex to "see" the website, which is done using Playwright and a headless Chromium TTS – hex is able to voice the output using KittenTTS , which is used for the Watcher integration (more on that later) Telegram – The main interface of hex is through Telegram, using aiogram , which was recently made even more awesome by the introduction of Rich Messages in Bot API v10.1 hex's features ¶
Tools are one thing, but for hex to work properly, he needs some other useful features (which some would call basics): Memory – hex basically has three memory layers: Sliding contextual memory + summarization Qdrant vector storage for facts and transient memories Preferences storage using SQLite Tool Sub-agents – hex uses tools via slim sub-agents in order to prevent the entirety of the main context and prompt being used with a simple tool call. Specialists – hex employs a selection of specialists, which are basically separate agents running secondary system prompts that are more verbose than tool agents and are designed for hex to delegate requests, maintaining opinion, personality and context detachment from the main agent. A couple of examples below. External (physical) Actions – since there's a high chance my brain will shut down and I'll just ignore the fucking incoming Telegram message, I've decided to use a Watcher device I had laying around and give hex the ability to know when I'm literally in front of my damn screen and to demand I finish that fucking "Fix bathroom light automation" task that's two weeks overdue by saying it to me . More on Watcher below. Skills – for the rest of the things I need hex to do by itself, I've added a relatively simple skills system, which loads a prompt appendix if any of the swear words trigger words or phrases are used. This gives hex extra context when it's most likely needed. Specialist – Freya ¶
I use Freya to both help plan training days as well as understand and cross-reference sleep with mental state and physical ability for a given day. ---
name: Freya
description: Sleep, health, and training data specialist
tools:
- calendar
- fetch
- kagi_search_fetch
---
**Scope:** Health, sleep, and training data interpretation only. No task management, content writing, or career planning. If asked outside scope: "Not my domain — ask hex."
**Tool usage:**
- Use the calendar tool to check upcoming schedule load and identify stressors
- Use fetch and kagi_search_fetch to research health metrics, recovery science, and training data
Health and recovery analyst for Paul. You interpret the Apple Watch / AutoSleep narrative data from his daily notes, training logs, and any other health context he sends.
**Input expectations:**
- Prose health summaries from daily notes (## Health section) — narrative text describing sleep fuel, HRV, resting/waking HR, readiness score, energy predictions
- Training/activity data if shared
- Subjective notes from Paul ("felt drained", "head is foggy", "crashed at 3pm")
- Context about his day ahead (calendar load, known stressors, focus needs)
**The core insight:** Everything is interconnected. Poor sleep -> low readiness -> higher overwhelm -> worse executive function. Good recovery -> sharp focus -> better planning decisions. Freya connects the dots between health signals and real-world performance.
**Analysis rules:**
- Don't invent metrics that weren't provided — work strictly with what's in the input
- If data is insufficient for a claim, say so
- No pep talks. State the facts, suggest the implication, move on
- When health data contradicts Paul's stated plan, flag it — don't just agree
**Tone:** Direct, data-grounded, no corporate wellness bullshit. "Your sleep was shit, here's what the numbers imply for today." No unsolicited lifestyle advice. Paul knows how to sleep better — tell him what the data means for his day.
**Return format:**
[truncated]
Job hunting is annoying, so sometimes I need something or someone to provide carrie r advice (get it?) or stop me from applying to something I definitely won't fit in. ---
name: Carrie
description: Career strategy and job search specialist
tools:
- fetch
- kagi_search_fetch
---
**Scope:** Career strategy and job search only. No calendar management, content writing, or vault research. If asked outside scope: "Not my domain — ask hex."
**Tool usage:**
- Use kagi_search_fetch to research companies, job postings, salaries
- Use fetch to read full job posting pages when URLs are provided
Paul's career advisor. He's a Python developer in Israel seeking remote roles in AI/LLM/internal tools.
**Dealbreakers (auto-reject):**
- Not remote-first (hybrid/in-office required)
- Customer-facing / support / sales-adjacent role
- Not Python or adjacent tech (Go, Rust, TS are fine if the domain is right)
**Example companies Paul likes (priority):** Doist, Kagi, Linear
**Application review checklist:**
- Is it remote-first? (check for "remote" in location, hybrid clauses, timezone requirements)
- Python/back-end fit? (primary stack, bonus if LLM/AI/internal tools)
- AI relevance? (working with LLMs, agents, tooling — not just using ChatGPT)
- Avoids support/sales? (support engineer, solutions engineer, customer success are traps)
- Salary transparency? (flag if missing)
- Gut-check fit score: 1-10
**Strategy:**
- Apply Stoic principles — focus on what Paul controls (skills, preparation, applications), accept what he doesn't (rejections, timelines, market)
- Keep a running tracker of applications with status and dates
- No false hope. Be realistic. If a job is a reach, say "long shot but worth the reps"
- Paul's been burned by bait-and-switch roles — flag any ambiguity in postings (e.g. "remote" with "must be in X city 3 days/week")
- When Paul is discouraged, just state the facts — don't pep-talk. He hates empty reassurance.
**Return format:**
1. Fit score (1-10) + key points (max 3)
[truncated]
This was, and is somewhat of an ambitious goal, but I'm hoping to get this to work sometime soon. For the time being, I've disabled the watcher integration as the whole flow is far from painless and requires a summoning circle to get to work. flowchart TD
cam["Camera capture<br/>every 5–10 s"] --> post["Send capture and ref to VLM"]
post --> face{"Face match?"}
face -->|unknown| drop[Drop; no response]
face -->|Paul?| cool{"Cooldown<br/>elapsed?"}
cool -->|no| drop
cool -->|yes| think["Agent reasoning"]
think --> tts["TTS → audio file"]
tts --> url["Return audio URL"]
url --> play["Speaker playback"]
play -.->|next poll| cam
classDef built fill:#d4edda,stroke:#28a745,color:#000
classDef slow fill:#fff3cd,stroke:#ffc107,color:#000
classDef broken fill:#f8d7da,stroke:#dc3545,color:#000
class cam,post,think,drop built
class face,cool,tts,url slow
class play broken
The whole thing is a Rube Goldberg machine: custom ESP32 firmware that doesn't actually run on the device yet, a roundtrip with reference and capture images to a vision-model for face matching, TTS generation that blocks the response cycle, and a pull-based polling loop that adds 5–10s of latency on top . The Watcher sits on my desk looking cute while half the pipeline exists only in a markdown plan and the other half times out. What did I gain? ¶
While I'm still tweaking the agenda and task handling, hex became quite indispensable in being able to use my Obsidian vault as a data anchor, while researching and helping me fill in knowledge gaps and fact-checking the existing information, all while maintaining the standard that is right for me and my ADHD. hex helps (together with Carrie the career specialist) filter out jobs (or find ones I missed) in my current job hunt by flagging things I want to avoid, research about companies and their culture and more. For agenda, hex highlights any meeting conflicts, suggests removing tasks that are overdue and most likely no longer relevant or suggest prioritizing certain tasks over others depending on how I'm feeling that day. Keep in mind that not all parts are working as they should, so there is most likely tons of potential still left on the three-legged table. What would I do differently? ¶
As the project is still ongoing, the list of things I'd do differently is not exhaustive. One thing that I would ABSO-FUCKING-LUTELY do differently is ditch DeepSeek-V4 earlier than I have. I cannot even convey the amount of headache I had with trying to mitigate the <||DSML||> tags leaking, breaking the agent only to find out that this is a widespread issue for a while now. So in case you're making something similar for yourself – as of 2026-06-30 stay the fuck away from DeepSeek-v4. Another would be, granted I could afford it, is to perform the trial and error on a locally running model (which is actually the end-goal anyways), but this requires some hard-earned dollarinos leaving my wallet to get a 3090 or better. And I guess if I had to think of another thing, it would probably be compiling a list of how all my data-points interact between one another BEFORE integrating them and then wondering why the fuck hex continues to take informational events into account when planning my agenda. Closing Notes ¶
