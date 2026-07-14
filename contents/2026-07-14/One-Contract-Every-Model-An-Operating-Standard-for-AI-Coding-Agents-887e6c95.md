---
source: "https://manazir.dev/blog/operating-standard-harness"
hn_url: "https://news.ycombinator.com/item?id=48902072"
title: "One Contract, Every Model: An Operating Standard for AI Coding Agents"
article_title: "One Contract, Every Model: Engineering an Operating Standard for AI Coding Agents"
author: "mnzralee"
captured_at: "2026-07-14T04:42:02Z"
capture_tool: "hn-digest"
hn_id: 48902072
score: 2
comments: 0
posted_at: "2026-07-14T03:47:50Z"
tags:
  - hacker-news
  - translated
---

# One Contract, Every Model: An Operating Standard for AI Coding Agents

- HN: [48902072](https://news.ycombinator.com/item?id=48902072)
- Source: [manazir.dev](https://manazir.dev/blog/operating-standard-harness)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T03:47:50Z

## Translation

タイトル: 1 つの契約、すべてのモデル: AI コーディング エージェントの運用標準
記事のタイトル: 1 つの契約、すべてのモデル: AI コーディング エージェントの運用標準のエンジニアリング
説明: 安価なモデルを最先端のモデルと同じくらいスマートにすることはできません。それは重みです。ただし、同じ契約で動作さ​​せることはできます。このようにして、ポータブルなオペレーティング標準と、それをサーバー上のすべてのクロード コード セッションに適用するハーネスを構築しました。

記事本文:
1 つの契約、すべてのモデル: AI コーディング エージェントの運用標準のエンジニアリング コンテンツへスキップ スリランカ、コロンボ 仕事について ブログ ギャラリー © 2026 · Manazir Ali サイトマップ RSS ブログ 1 つの契約、すべてのモデル: AI コーディング エージェントの運用標準のエンジニアリング
最近行ったエンジニアリングの一部を共有したいと思います。これは、AI コーディング エージェントの使用についての私の考え方を変えました。それは、私が大声で尋ねた素朴な質問から始まりました。「Sonnet と Opus をフロンティア モデルのように動作させることはできるでしょうか?」それは質問に値するものよりも有益なところで終わりました。
簡単に言うと、小型のモデルを大型のモデルと同じように機能させることはできません。それはウェイトであり、ウェイトを移動するプロンプトはありません。しかし、優秀なエージェントと平凡なエージェントを分けるのは能力だけではありません。残りの半分は *教義* です。つまり、エージェントがどのように通信するか、いつ停止するか、どのように作業を証明するか、決定する前にどの程度深く分析するかです。そして教義は完全に移植可能です。それを一度書き留めて、システム プロンプト レイヤーとして適用すると、高価なものも安価なものも同様に、すべてのモデル層が同じ契約に従って動作し始めます。
この投稿は論文と設計図です。ここにあるものはすべて、独自のスタックに適応できます。再利用可能なパターンがどこに存在するのか、私の詳細が単なる説明にすぎないのかを指摘します。
私は、ブロックチェーンベースの金融プロトコル、20 を超えるバックエンド マイクロサービス、いくつかの Next.js フロントエンド、スマート コントラクト、および Kubernetes クラスターなど、かなり重いコードベースを実行しています。私は、専門のサブエージェントを使用して、クロード コードを通じてほとんどの作業を推進しています。どの日でも、アーキテクチャ エージェントは最も性能の高いモデルで実行され、実装エージェントは中間層のモデルで実行され、高速リサーチ エージェントは小規模なモデルで実行されます。 3 つの異なるインテリジェンス、1 つのコードベース、1 つの標準セット

ds。
摩擦は、安価なモデルが愚かであるということでは決してありませんでした。それは、安価なモデルの動作が異なるということでした。テストを実行せずにタスクが「完了」と宣言されます。最終メッセージは、答えの代わりに実装の詳細の壁で始まります。爆発範囲をトレースする代わりに、単一の grep から変更を加えます。タスクの途中で停止し、自動的に解決できるかもしれない質問をします。高価なモデルはデフォルトでこれらのことを正しく行っていました。安い方を伝える必要がありました。
長い間、私はそれを、より安価なモデルを使用することによる避けられない税金として扱いました。それは間違ったフレームでした。私が求めていた行動は知性ではありませんでした。彼らは行為でした。そして行為を特定することができます。
核となる考え方: ドクトリンは能力ではない
この区別が論文全体なので、正確に言っておきます。
能力とは、モデルが理解できるもの、つまり厳密な推論、斬新な統合、大きな問題を念頭に置く能力のことです。それは重みの中に生きています。プロンプトでは追加できません。 「GPT-3.5 を GPT-5 と同じくらいスマートにするシステム プロンプト」を販売する人は、何も販売しません。
原則とは、通信契約、何かが完了したと判断するための基準、コミットする前に変更をエンドツーエンドで追跡する規律、まだ実行できる作業がある間は停止しないというルールなど、モデルが持つあらゆる機能に基づいてモデルがどのように動作するかを表します。それはどれも知性ではありません。すべてテキストで指定可能です。そして重要なことは、有能なモデルがすでにこれらのことを行っている場合、それはそのモデルに向けてトレーニングされているためであり、つまり、潜在的な知性のギャップが残っているとしても、同じ動作をより有能なモデルに指示することができ、目に見える品質のギャップのほとんどを埋めることができます。
つまり、目標は「小型モデルをスマートにする」ことではありませんでした。それは、どんな関係であっても、あらゆるモデルを作るというものでした。

r、1 つの運営契約を遵守します。結果でリードし、すべての完了主張を証拠で裏付け、決定する前に追跡し、早期に停止しない中間層のターンは、外部から観察可能なあらゆる尺度で見て、フロンティア標準のターンです。標準化できる基本的な職業上の行動ではなく、タスクの正確性の上限が追加の能力を必要とする場合、真に大きなモデルに手が伸びます。
これらすべての中心となる成果物は、私が「オペレーティングスタンダード」と呼ぶ単一の文書です。これは約 12 のセクションで構成されており、契約を雰囲気ではなく強制可能な行動としてエンコードしています。耐荷重柱:
結果でリードする。最後のメッセージは、読者が作品を初めて見ることです。答えで始まり、それをサポートします。読みやすさは簡潔さに勝ります。断片や矢印の鎖に圧縮するのではなく、読者の次の動きを変えない詳細を切り取ることで短くします。読者が必要とするすべてが最後のメッセージの中に生きています。
成果物で完成を証明します。 「うまくいくはず」も「見栄えも良い」もありません。すべての完了/修正/合格のクレームには、コミット ハッシュ、逐語的なテスト出力、終了コード、またはファイル リストが含まれます。オーケストレーターは、完了の要求を信頼する前に、受け入れコマンド自体を再実行します。エージェントの言葉は受け入れられません。
決して直感からではなく、深さ優先で決定してください。決定する前に、スキーマ、データ アクセス、ドメイン、アプリケーション、API コントラクト、すべてのフロントエンド、非同期パイプライン、チェーン、移行、テストなど、接触するすべてのレイヤーを通じて横断的な変更を追跡します。緑色の型チェックは、実行時や契約のリップルについては何も証明しません。
早くやめないでください。リクエストから続く元に戻せるアクションについては、続行してください。ターンを終了する前に、最後の段落をもう一度読んでください。それが計画、質問、または「次は X をする」という約束であれば、今すぐその作業を行ってください。

タスクが完了するか、ユーザーが完全にブロックされた場合にのみ停止します。
うまくいく最も単純なことを実行してください。要求されていない機能、リファクタリング、または抽象化はありません。不可能な状態に対する防御的な検証は行われず、一方で、お金やユーザー向けのディスプレイなど、重要なパスでの入力エラーが発生してフェイルクローズされます。
正直な発見をすべて無制限に開示します。きれいに見せるために、警告の長い尾をトリミングしないでください。常に「約 2 つ」の調査結果を明らかにする報告書は、解釈により開示が不十分である。
これを採用する人に強調しておきたいのは、私がこれらの断片を好みからでっち上げたわけではないということです。行動ナッジ、つまりモデルをこの行動に導く正確な文言は、モデル ベンダー自身が公開した移行およびプロンプト ガイダンスから抽出されます。フロンティア モデルはデフォルトでこれらの動作を示します。ベンダーは、そうでないモデルにそれらを浸透させる正確な表現を文書化しています。私はその文言を取り上げて契約書にまとめました。この根拠があるからこそ、ただ単によく読むだけではなく、実際に効果があるのです。独自のスニペットを作成する場合は、ブログ投稿 (この記事を含む) の意見ではなく、プロバイダーのガイダンスに基づいてスニペットを入手してください。
2チャンネルなのでスキップ不可
誰もロードしない標準は装飾です。私は 2 つの方法を適用して、誰かがオプトインを忘れていないかどうかに関係なく表示されるようにします。
起動時に、追加されたシステム プロンプト ファイルとして。薄いラッパー スクリプトは、システム プロンプトにすでに含まれている完全なドクトリンを使用してエージェントを起動し、フォークするすべてのサブエージェントに継承されます。
セッション内 (常にロードされるルールとして)。ハーネスは、交渉不可能なコアをインラインで伝達し、完全なドクトリンを示す小さなルール ファイルをセッションごとに自動ロードします。したがって、セッションが単純な方法で開始された場合でも、契約までは引き続き実行されます。
セッション中の小さなルールは意図的に小さくされています: 完全なドクトリン

起動レイヤーに存在し、常に読み込まれるコンテキスト バジェットの外にあります。これにより、コアの存在が保証されながら、永続的なコストが低く抑えられます。
私を謙虚にさせた教訓: ハーネスに実際に負荷がかかることを確認する
これは、私が発見できて最もうれしかった発見です。それは、システム全体を静かにプラシーボにしてしまうようなものだからです。
私には 20 個の美しく書かれたサブエージェントの役割定義がありました。彼らは何か月も存在していました。それらは私の構成で参照されました。私は彼らが派遣されていると信じていました。実際に実行中のエージェントの利用可能なサブエージェント タイプのリストを調べたところ、それらはどれも存在しませんでした。ランタイムには、組み込みの汎用エージェントのみが表示されました。
役割ファイルは、エージェント フレームワークがエージェント定義として解析しない特注のメタデータ形式を使用していました。それらは機能的にはドキュメントでした。私のオーケストレーションは、ずっと静かに汎用エージェントにフォールバックし、ロールコンテキストをプロンプトに貼り付けることで補っていました。それは十分にうまく機能したので、宣言されたスペシャリストが本物ではないことにまったく気づきませんでした。
この修正は機械的でした。すべてのロール ファイルをフレームワークの実際のネイティブ形式、名前、説明、ツール許可リスト、およびモデル層を含む有効なフロントマターに正規化することで、フレームワークはそれぞれをファーストクラスのディスパッチ可能なエージェントとして検出します。私は、20 個すべてでこれを均一に実行する決定論的スクリプトを作成し、すべてのファイルが解析されることを検証し、ランタイムに対してエージェントが表示されることを確認しました。
伝えられる教訓はバグよりも重要です。ハーネスのグラウンド トゥルースは、構成が主張するものではなく、ランタイムが実際にロードするものです。宣言されたエージェント、有線フック、スコープ付きルール。ランタイムがそれを確認するまでは、どれも本物ではありません。 「構成に含まれている」を仮説として扱い、ランタイムがロードした内容に関するランタイム自身のレポートを唯一の evi として扱います。

それを閉じるデンス。私は現在、テストの合格を検証するのと同じ方法で、意図ではなく出力を見て検出を検証しています。
安全な実施: セッションを妨げない完了ゲート
プロンプトの教義はガイダンスです。メカニズムが必要な場合もあります。私が構築した最も強力なものは完了ゲートです。エージェントが完了しようとしたときに起動し、実際にターンが終了する前に実際の証拠に照らして「完了」の主張を再チェックするよう強制するフックです。
完了をブロックする可能性のあるフックを作成するのは本当に危険です。単純なバージョンでは無限ループが作成されるためです。フックは停止をブロックし、モデルは続行し、再び停止しようとし、フックは再びブロックし、永遠に続きます。門の造りが悪くても気にならない。セッションが中断されます。したがって、ここでのエンジニアリングはほぼ完全に安全性に関するものであり、3 つの不変条件は直接盗む価値があります。
デフォルトでは不活性です。環境フラグがゲートをオンにしない限り、ゲートは何も行いません。有線で出荷されますが、電源はオフになっているため、既存のワークフローを驚かせることはありません。早期完了が本当のリスクとなる長時間の自律実行に対してこれを有効にします。
ループセーフ。フレームワークはストップフックに、このターンですでに起動したかどうかを伝えます。ゲートはそのフラグを読み取り、2 回目のブロックを拒否します。それは最大でも 1 回介入し、その後ターンを終了させます。この 1 つのチェックが、有用なゲートと無限ループの違いとなります。
フェールオープン。例外、入力の不正、予期せぬ事態が発生しても、ゲートは正常に終了し、停止が続行されます。ゲートのバグは、何も操作を行わない状態にまで悪化し、セッションがスタックすることはありません。毎ターンのクリティカル パスに位置するツールでは、フェールオープンはオプションではありません。
シェルに接続する前に、シェルの 3 つのプロパティをすべてテストしました。フラグが設定されていない場合は不活性、有効な場合は 1 回だけブロック、すでに起動されたフラグの下では再ブロックしない、終了します。

ガベージ入力時にクリーンします。敵対的テストを行っていない安全メカニズムは、安全装置ではなく、責任です。
段階的構成: グローバルな原則、ローカルな専門家
この標準が 1 つのプロジェクトで機能すると、当然の疑問は次のとおりです。それは私が働いているどこにでも適用されるのか、それともここだけで適用されるのか?スコーピングを正しく行うことは、詳細ではなく、実際のアーキテクチャ上の決定であることが判明しました。
通常、エージェント フレームワークには (少なくとも) 2 つの構成層があります。1 つはプロジェクト内で作業している場合にのみ読み込まれるプロジェクト層、もう 1 つはあらゆるセッションで、どこでも読み込まれるユーザー層です。間違いは、すべてを 1 つに押し込むことです。正しい動きは、リーチごとに分割することです。
この原則は普遍的なものであるため、ユーザー層に属します。オペレーティング標準はあらゆるコードベースに適用されます。行為はプロジェクト固有のものではありません。これをユーザー レベルに配置したので、マシン上のあらゆるディレクトリ内のすべてのセッションがコントラクトに対して実行されます。グローバル コピーをメイン リポジトリ内のバージョン管理されたコピーにリンクして戻すことで、信頼できる単一のソースを維持しました。そのため、ソースが逸脱することはありません。
スペシャリスト エージェントはドメイン固有であるため、プロジェクト層に留まります。私の 20 人のエージェントは、このプロトコルのサービス、チェーン、名前空間について知っています。グローバル化すると表面化する

[切り捨てられた]

## Original Extract

You can't make a cheaper model as smart as a frontier one: that's weights. But you can make it operate to the same contract. This is how I built a portable operating standard and the harness that enforces it across every Claude Code session on my server.

One Contract, Every Model: Engineering an Operating Standard for AI Coding Agents Skip to content Colombo, Sri Lanka About Work Blog Gallery © 2026 · Manazir Ali Sitemap RSS Blog One Contract, Every Model: Engineering an Operating Standard for AI Coding Agents
I want to share a piece of engineering I did recently that changed how I think about working with AI coding agents. It started from a naive question I asked out loud: "can I make Sonnet and Opus behave like the frontier model?" It ended somewhere more useful than the question deserved.
The short version: you cannot make a smaller model as capable as a larger one. That's weights, and no prompt moves weights. But capability is not the only thing that separates a good agent turn from a mediocre one. The other half is *doctrine*: how the agent communicates, when it stops, how it proves its work, how deeply it analyses before deciding. And doctrine is completely portable. You can write it down once, apply it as a system-prompt layer, and every model tier, the expensive one and the cheap one alike, starts operating to the same contract.
This post is the thesis and the blueprint. Everything here is adaptable to your own stack; I'll point out where the reusable pattern lives versus where my specifics are just illustration.
I run a fairly heavy codebase: a blockchain-based financial protocol, twenty-plus backend microservices, several Next.js frontends, smart contracts, and a Kubernetes cluster. I drive most of the work through Claude Code with a fleet of specialised sub-agents. On any given day, an architecture agent runs on the most capable model, an implementation agent runs on a mid-tier model, and a fast research agent runs on a small one. Three different intelligences, one codebase, one set of standards.
The friction was never that the cheap model was dumb . It was that the cheap model behaved differently . It would declare a task "done" without running the test. It would open its final message with a wall of implementation detail instead of the answer. It would make a change from a single grep instead of tracing the blast radius. It would stop mid-task and ask a question it could have resolved itself. The expensive model did these things right by default; the cheaper one needed to be told.
For a long time I treated that as an unavoidable tax of using cheaper models. That was the wrong frame. The behaviours I wanted were not intelligence. They were conduct . And conduct can be specified.
The core idea: doctrine is not capability
This distinction is the whole thesis, so let me be precise about it.
Capability is what the model can figure out: the hard reasoning, the novel synthesis, the ability to hold a large problem in mind. It lives in the weights. A prompt cannot add it. Anyone selling you a "system prompt that makes GPT-3.5 as smart as GPT-5" is selling you nothing.
Doctrine is how the model conducts itself around whatever capability it has: the communication contract, the bar for calling something done, the discipline of tracing a change end-to-end before committing to it, the rule about not stopping while there's still work you can do. None of that is intelligence. All of it is specifiable in text. And critically, when a capable model already does these things, it's because it was trained toward them, which means the same behaviours can be instructed into a less capable model, closing most of the visible quality gap even though the underlying intelligence gap remains.
So the goal was never "make the small model smart." It was: make every model, whatever its tier, honour one operating contract. A mid-tier turn that leads with the outcome, backs every completion claim with evidence, traces before it decides, and doesn't stop early is , by every externally observable measure, a frontier-standard turn. You reach for the genuinely bigger model when the task's correctness ceiling demands the extra capability, not for basic professional conduct, which you can standardise.
The artifact at the centre of all this is a single document I call the Operating Standard. It's about a dozen sections, and it encodes the contract as enforceable behaviour rather than vibes. The load-bearing pillars:
Lead with the outcome. The final message is the reader's first look at the work: it opens with the answer, then supports it. Readability beats brevity: you shorten by cutting detail that doesn't change the reader's next move, never by compressing into fragments and arrow-chains. Everything the reader needs lives in that last message.
Prove completion with artifacts. No "should work," no "looks good." Every done / fixed / passing claim carries a commit hash, verbatim test output, an exit code, or a file listing. The orchestrator re-runs the acceptance command itself before it trusts a completion claim. It does not take the agent's word.
Decide depth-first, never from a hunch. Trace a cross-cutting change through every layer it touches before deciding: schema, data access, domain, application, API contract, every frontend, the async pipeline, the chain, the migration, the tests. A green type-check proves nothing about runtime or contract ripples.
Do not stop early. For reversible actions that follow from the request, proceed. Before ending a turn, re-read the last paragraph: if it's a plan, a question, or an "I'll do X next" promise, do that work now . Stop only when the task is complete or genuinely blocked on the user.
Do the simplest thing that works well. No unrequested features, refactors, or abstractions; no defensive validation for impossible states, while still failing closed with typed errors on the paths that matter, like money and user-facing displays.
Disclose every honest finding, uncapped. Never trim the long tail of caveats to look tidy. A report that always surfaces "about two" findings is under-disclosing by construction.
Here's the part I want to stress for anyone adapting this: I did not invent these snippets from taste. The behavioural nudges, the exact wording that pulls a model toward this conduct, are drawn from the model vendor's own published migration and prompting guidance. The frontier model exhibits these behaviours by default; the vendor documents the precise phrasing that instils them in models that don't. I lifted that phrasing and organised it into a contract. That grounding is why it actually works rather than just reading well. When you build your own, source your snippets from your provider's guidance, not from a blog post's opinion (including this one).
Two channels, so it can't be skipped
A standard nobody loads is decoration. I apply mine two ways so it's present whether or not someone remembers to opt in:
At launch , as an appended system-prompt file. A thin wrapper script starts the agent with the full doctrine already in the system prompt, inherited by every sub-agent it forks.
In-session , as an always-loaded rule. The harness auto-loads a small rule file every session that carries the non-negotiable core inline and points to the full doctrine. So even a session started the plain way still runs to the contract.
The small in-session rule is deliberately tiny: the full doctrine lives in the launch layer, which sits outside the always-loaded context budget. That keeps the permanent cost low while still guaranteeing the core is present.
The lesson that humbled me: verify the harness actually loads
Here is the finding I'm most glad I caught, because it's the kind of thing that quietly makes a whole system a placebo.
I had twenty beautifully written sub-agent role definitions. They'd existed for months. They were referenced in my config. I believed they were being dispatched. When I actually inspected the running agent's list of available sub-agent types, none of them were there. The runtime showed only the built-in generic agents.
The role files were using a bespoke metadata format the agent framework never parsed as agent definitions. They were, functionally, documentation. My orchestration had been quietly falling back to generic agents the whole time and compensating by pasting role context into prompts. It worked well enough that I never noticed the declared specialists weren't real.
The fix was mechanical: normalise every role file to the framework's actual native format, valid frontmatter with a name, a description, a tool allow-list, and a model tier, so the framework discovers each one as a first-class dispatchable agent. I wrote a deterministic script to do it uniformly across all twenty, validated that every file parsed, and confirmed against the runtime that the agents now appear.
The transferable lesson is bigger than the bug: the ground truth of a harness is what the runtime actually loads, not what your config claims. A declared agent, a wired hook, a scoped rule: none of it is real until you've watched the runtime confirm it. Treat "it's in the config" as a hypothesis, and the runtime's own report of what it loaded as the only evidence that closes it. I now verify discovery the same way I verify a passing test: by looking at the output, not the intention.
Safe enforcement: a completion gate that can't wedge your session
Doctrine in the prompt is guidance. Sometimes you want a mechanism . The strongest one I built is a completion gate: a hook that fires when the agent tries to finish and forces it to re-check its "done" claims against real evidence before the turn actually ends.
Hooks that can block completion are genuinely dangerous to write, because the naive version creates an infinite loop: the hook blocks the stop, the model continues, tries to stop again, the hook blocks again, forever. A badly built gate doesn't annoy you. It wedges the session. So the engineering here is almost entirely about safety , and the three invariants are worth stealing directly:
Inert by default. The gate does nothing unless an environment flag turns it on. It ships wired but off, so it can never surprise an existing workflow. You enable it for the long autonomous runs where premature-completion is the real risk.
Loop-safe. The framework tells a stop-hook whether it already fired this turn . The gate reads that flag and refuses to block a second time. It intervenes at most once, then lets the turn end. This single check is the difference between a useful gate and an infinite loop.
Fail-open. Any exception, any malformed input, any surprise: the gate exits cleanly and lets the stop proceed. A bug in the gate degrades to a no-op, never to a stuck session. On a tool that sits in the critical path of every turn, fail-open is not optional.
I tested all three properties from the shell before wiring it in: inert when the flag is unset, blocks exactly once when enabled, never re-blocks under the already-fired flag, exits clean on garbage input. A safety mechanism you haven't adversarially tested is a liability, not a safeguard.
Tiered configuration: global doctrine, local specialists
Once the standard worked in one project, the obvious question was: does it apply everywhere I work, or just here? Getting the scoping right turned out to be a real architectural decision, not a detail.
Agent frameworks typically have (at least) two configuration tiers: a project tier that loads only when you're working inside that project, and a user tier that loads in every session, everywhere. The mistake would be to shove everything into one. The right move is to split by reach :
The doctrine is universal , so it belongs in the user tier. The Operating Standard applies to any codebase; conduct is not project-specific. I placed it at the user level so every session on the machine, in any directory, runs to the contract. I kept a single source of truth by linking the global copy back to the versioned one in my main repo, so it can never drift.
The specialist agents are domain-specific , so they stay in the project tier. My twenty agents know about this protocol's services, its chain, its namespaces. Making them global would surface

[truncated]
