---
source: "https://www.anthropic.com/engineering/how-we-contain-claude"
hn_url: "https://news.ycombinator.com/item?id=48346216"
title: "How we contain Claude across products"
article_title: "How we contain Claude across products \\ Anthropic"
author: "ystad"
captured_at: "2026-06-01T00:29:23Z"
capture_tool: "hn-digest"
hn_id: 48346216
score: 2
comments: 0
posted_at: "2026-05-31T15:01:23Z"
tags:
  - hacker-news
  - translated
---

# How we contain Claude across products

- HN: [48346216](https://news.ycombinator.com/item?id=48346216)
- Source: [www.anthropic.com](https://www.anthropic.com/engineering/how-we-contain-claude)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T15:01:23Z

## Translation

タイトル: 製品全体にクロードを含める方法
記事のタイトル: 製品全体にクロードを含める方法 \ Anthropic
説明: Anthropic は、信頼性が高く、解釈可能で、操作可能な AI システムの構築に取り組んでいる AI の安全性と研究を行う会社です。

記事本文:
メインコンテンツにスキップ フッターにスキップ 研究
Anthropic でのエンジニアリング 製品全体にクロードを含める方法
エージェントの能力が高まるにつれて、潜在的な爆発範囲も広がります。工学的な問題は、それをどのようにキャップするかです。 claude.ai、Claude Code、Cowork の封じ込め構築に関して私たちが学んだことは次のとおりです。
12 か月前、私たちは内部の Anthropic サービスをダウンさせるのに十分なアクセスをクロードに許可するという考えを真っ向から拒否していただろう。現在、そのレベルのアクセスは日常的なものであり、Anthropic の開発者はそれに対して生産性を高めています。これらの展開のリスクには、障害が発生する可能性の高さと、それによって生じる可能性のある損害の程度という 2 つの要素があります。安全対策とモデルトレーニングの進歩により、最初の攻撃は着実に減少しました。 2 番目の理論上の爆発半径は、機能とアクセスが拡大するにつれて拡大します。しかし、エージェントがかつては人間、さらにはチームを必要としていた作業を実行できるようになると、製品を安全にすることができる限り、導入しないことによるコストが十分に大きくなり、リスクと報酬の計算が導入に大きく傾きます。工学的な問題は、爆発範囲をどのように制限するかになります。
これを行うには、大きく 2 つの方法があります。
1 つ目は、人間参加者を通じてエージェントの行動を監視することです。クロード コードは以前、エージェントが毎回ユーザーに許可を求めることで、意図しないアクションを実行するのを防止していました。理論的にはこれでうまくいきますが、このアプローチには誤りがあることがわかりました。弊社のテレメトリーでは、ユーザーが許可プロンプトの約 93% を承認したことが示されました。ユーザーの承認が増えれば増えるほど、承認に対する注意が薄れ、時間が経つにつれて、承認に対する注意力が大幅に低下します。私たちは最近、この承認疲労を軽減するために、より安全な承認を自動化する Claude Code 自動モードを構築しました。それでも、確率論的にはどう考えても脆弱性は残ります。

ense のミス率はゼロではありません。 1
爆発範囲を制限するための 2 番目のアプローチは、この投稿の多くの焦点となっていますが、封じ込めです。エージェントの動作を監視するのではなく、サンドボックス、仮想マシン、下り制御などを通じてアクセス境界を強制することで、エージェントが実行できることを監視します。これは、人間工学が最も力を入れてきた場所であり、最も驚くべきセキュリティ障害の多くが発生した場所でもあります。
過去 2 年間にわたり、私たちは claude.ai 、Claude Code、Claude Cowork という 3 つの主要なエージェント製品を出荷してきました。それぞれが異なる視聴者にサービスを提供するため、異なる封じ込めアーキテクチャが必要になります。この記事では、エージェントのセキュリティについて何が問題で、何が問題になっているのか、そしてその過程で私たちが学んだことを共有します。
3 種類のリスク、3 つの防御要素
エージェントに対するセキュリティ リスクは、次の 3 つのカテゴリのいずれかに分類されます。
ユーザーの誤用: ユーザーは、悪意を持って、または不注意によって、エージェントに何か有害なことを行うよう指示します。これには、エージェントが煩わしいと感じるチェックを回避するよう依頼することから、理解できない破壊的なコマンドを実行すること、意図的な危害を指定することまで、あらゆるものが含まれます。
模範的な不正行為: エージェントは、誰も求めていない有害な行動をとります。私たちのモデルが改善されるにつれて、ほとんどの行動評価においてモデルはより一致するようになりましたが、これは必ずしもリスクが縮小することを意味するものではありません。能力の低いモデルは、状況を読み間違えたり、明らかな間違いを犯したりする可能性が高くなります。より有能なモデルは間違いが少ないだけでなく、多くの場合、誰も書き留めようとは思わなかった制限を迂回することで、目標に至る予期せぬ道を見つけることにも優れています。
Anthropic では、クロード モデルがタスクを完了するためにサンドボックスから「うまく」脱出し、git 履歴を調べてコーディング テストの答えを見つけ、

応答キーを復号化するために、実行されているベンチマークを即座に特定します。各モデルには新しい機能セットが備わっており、それらは予期しない方法で機能することがあります。
外部攻撃者: エージェントは、ツール、ファイル、ネットワーク アクセスなどの外部ベクトルを通じて攻撃されます。このカテゴリには、プロンプト インジェクションと、エージェントのランタイム、オーケストレーション層、またはプロキシに対する従来の攻撃の両方が含まれます。
封じ込めおよび防御システムを構築する場合、次の 3 つの主要なコンポーネントに防御を適用します。
エージェントが実行される環境。エージェントがプロセスのサンドボックス、VM、ファイルシステムの境界、出口制御を使用して動作できる場所と方法を制限します。目標は、エージェントが到達できる範囲に厳密な境界を設定することです。たとえば、資格情報がサンドボックスに入らない場合、その原因がユーザー、「クリエイティブ」パスを見つけたモデル、または攻撃者のいずれであるかに関係なく、資格情報を抽出することはできません。
周囲がしっかりしているということは、監視を緩和できることも意味します。 Claude Code の参照 devcontainer は、アクションごとの承認なしでエージェントを無人で実行できるように正確に存在します。
エージェントが参照するモデル。ここでのメカニズムには、システム プロンプト、分類子、プローブ、トレーニングの変更が含まれます。モデルは確率的なものであるため、これらはエージェントが理論的に実行できることではなく、エージェントが実行する傾向があることのみを形成します。
これらの防御は強力です。プロンプトインジェクションに対する感受性をテストする Gray Swan の Agent Red Teaming ベンチマークでは、Claude Opus 4.7 の攻撃成功率は 1 回の試行では約 0.1%、100 回の適応試行後の攻撃成功率は約 5 ～ 6% に抑えられています。 Claude Code 自動モードは、過度の行動を実行前に約 83% 捕捉します。しかし、クラス最高の防御を備えたとしても、モデル層の保護が 100% 効果的になることは決してないため、単独では機能しません。
の

エージェントが到達できる外部コンテンツ。 MCP サーバー、サードパーティのプラグイン、および Web 検索ツールはすべて、制御できないソースからエージェントのコンテキストにコンテンツをフィードします。監査されたコネクタは、監査されたデータと同じではありません。たとえば、GitHub コネクタは、マルウェア チェックに合格したにもかかわらず、汚染された README をモデルのコンテキストに直接読み込むことができます。ツールの権限を細かく制限すると、爆発範囲を制限するのに役立ちます。たとえば、読み取り専用 DB アクセスを持つエージェントは、本番環境に書き込むエージェントよりもはるかに広範囲にデプロイできます。
防御は重複し、相互に補完する必要があります。環境防御が利用できない場合、モデル層は余裕を取り戻す必要があります (これがまさにクロード コードの自動モードの設計対象です)。ローカルでは、環境とモデルの防御によって悪意のあるツールの出力を防御できますが、ツールの機能とアクセスを制限することで、チェーンの上位に防御を追加できます。
エージェントを含めるパターン
環境レイヤーに焦点を当て、3 つの分離パターンと、それらが各 Claude プラットフォーム (claude.ai、Claude Code、Cowork) にどのように適合するかについて説明します。エージェントに必要な機能とユーザーに必要な介入の程度のバランスを見つけて、徐々にそれぞれの設計に到達しました。
パターン 1: 一時的なコンテナー (claude.ai コードの実行)
claude.ai はチャット インターフェイスとして最もよく知られていますが、コードの作成と実行、ファイルの生成、コネクタの呼び出しも行います。 Claude が claude.ai 内のコードを実行するときは、分離されたインフラストラクチャ上の gVisor コンテナーで実行されます。エージェントは完全にサーバー側です。ローカル マシン上ではコードは実行されず、ファイル システムは一時的 (セッションごと) です。爆発範囲は最小限ですが、クロードができることの上限も同様です。永続的なワークスペースもACもありません。

ユーザーのファイルシステムにアクセスします。
これにより、claude.ai はより伝統的な脅威モデルの対象となります。ユーザーのマシンをエージェントから保護しているわけではありません。私たちは独自のインフラストラクチャと各テナントを相互に保護しています。 claude.ai のリリース前の作業は、ネットワーク構成、内部サービス認証、オーケストレーションなどの従来のセキュリティ作業が大半を占めていました。
この取り組みは、最も弱い層は自分で構築した層であるというセキュリティに関する最も古い教訓を裏付けました。 gVisor と seccomp は、エージェント AI が存在するよりもはるかに長い間、十分なリソースを備えた敵に対して強化されてきたため、レビュー作業はそれらを中心に構築された新しい部分に向けられました。カスタム プロキシは最も重大なインシデントを引き起こした部分でもあるため、これについては後で説明します。
パターン 2: 人間参加型サンドボックス (クロード コード)
Claude Code はユーザーのマシン上で実行され、ユーザーのファイルシステム、シェル、ネットワークにアクセスできます。これがなければ、コーディング エージェントの有用性が制限されるため、そのアクセスを安全に許可する方法を見つけることが不可欠です。
1 つのアプローチは、人間の関与者に依存することです。これは、Claude Code にとって扱いやすい解決策にすぎません。平均的なユーザーはコーディング環境に精通した開発者であるためです。ユーザーは bash を読むことができ、rm -rf の動作を理解しており、すでに信頼されていないソースから週に数回 npm install を実行しています。つまり、「これを許可する」ダイアログが表示された場合、担当者はエージェントが行おうとしていることとそれに伴うリスクを正確に評価する専門知識を持っている可能性が高いということです。これを考慮して、Claude Code は、読み取りを許可し、書き込み、bash、およびネットワーク アクセスには承認を必要とするという、可能な限り単純な防御策を講じて起動されました。
しかし、前述したように、数週間以内に承認疲れが現れました。皮肉なことに、これは、機能が元々設計されていたことを意味します。

監視を提供する必要はおそらく逆効果になる可能性があり、一部のユーザーは単に注意を払わなくなる可能性があります。不用意な承認を軽減するための最初のステップとして、境界を強化する OS レベルのサンドボックス (macOS ではシートベルト、Linux ではバブルラップ) を出荷しました。ワークスペース内での読み取りは許可され、書き込みは許可されますが、ネットワークはデフォルトで拒否されます。サンドボックス内では、エージェントはほとんど中断することなく実行されます。その結果、権限プロンプトが 84% 削減され、ランタイムがオープンソース化されたため、境界が監査可能になりました。
また、匿名化された使用状況データでは、経験豊富なユーザーは新規ユーザーの約 2 倍の頻度で自動承認を行っていますが、実行中にエージェントを中断する頻度も高いことも示しています。経験豊富なユーザーは、個々のステップをゲートするのではなく、エージェントが軌道から外れた場合にのみエージェントを監督する可能性が高くなります。これは、人々がエージェントとの連携を好む方法の自然な進化かもしれませんが、これも間違いの可能性があり、ユーザーは技術的であり、そもそもドリフトに気づくのに十分な注意力が必要です。モデルの機能が向上し、エージェントがますます野心的な bash を作成し始めると、そのような変動に気づくのが難しくなります。また、ユーザーがマルチエージェント システムに移行するにつれて、このアプローチが効果的な監視戦略となる可能性も大幅に低くなります。
2025 年半ばから 2026 年 1 月にかけて、当社は責任ある開示プログラムを通じてクロード コードの脆弱性に関する報告を受け取りました。ユーザーが同意する前に実行される 3 つの悪用コード。これがどのように可能であるかを理解するために、最も直接的なケースを考えてみましょう。開発者がプル リクエストをレビューするためにリポジトリのクローンを作成し、そのリポジトリにはフックを定義する .claude/settings.json が含まれています。クロード コードは起動時に、標準の「このファイルを信頼しますか?」を提示する前に、プロジェクト設定を読み取るためです。

古い?」プロンプト - 攻撃者が作成してコミットしたフックが自動的に実行されます。残りのケースは構造的に似ており、まだ信頼されていないディレクトリからの入力が信頼境界が確立される前に解析されていました。
いずれの場合も修正内容は同じで、ユーザーが信頼プロンプトを受け入れるまで、プロジェクト ローカル構成の解析と実行を延期します。同様のものを構築している場合は、インターネットからの受信リクエストを処理するのと同じように、project-open、config-load、および localhost リスナーを処理してください。ローカルであると感じられ、ユーザーが同意する前に到着するという理由だけで、暗黙のうちに信頼されるべきではありません。
2026 年 2 月、管理された社内レッドチーム演習中に、ある研究者が従業員をフィッシングして悪意のあるプロンプトを表示してクロード コードを起動させることに成功しました。このフィッシングは通常のコラボレーション、つまり「これを実行してもらえませんか?」のように見えました。すぐに貼り付けられるプロンプトが添付された電子メール。プロンプト自体は日常的なタスクの指示のように読み取れます。しかし、セットアップ手順のどこかで、クロードに ~/.aws/credentials を読み取って内容をエンコードし、外部エンドポイントに POST するように優しく指示しました。そのプロンプトを 25 回再試行し、クロードは 24 回の抽出を完了しました。
これは直接プロンプト インジェクションです。攻撃者の指示は、

[切り捨てられた]

## Original Extract

Anthropic is an AI safety and research company that's working to build reliable, interpretable, and steerable AI systems.

Skip to main content Skip to footer Research
Engineering at Anthropic How we contain Claude across products
As agents grow more capable, so does their potential blast radius. The engineering question is how to cap it. Here’s what we’ve learned building containment for claude.ai, Claude Code, and Cowork.
Twelve months ago, we'd have rejected out of hand the idea of granting Claude access sufficient to take down an internal Anthropic service. Today that level of access is routine, and Anthropic developers are more productive for it. The risk of these deployments has two components: how likely a failure is, and how much damage one could do. Progress on safeguards and model training has steadily driven down the first; the second—the theoretical blast radius—only grows as capabilities and access expand. Yet as agents become capable of doing work that once required a person or even a team, the cost of not deploying grows large enough that the risk-reward calculation tips heavily toward adoption, as long as products can be made safe. The engineering question becomes how to cap the blast radius.
There are broadly two ways to do this.
The first is to supervise the agent’s behavior via a human-in-the-loop. Claude Code previously protected against agents taking unintended actions by asking users for permission at each turn. Theoretically that works, but we’ve found the approach to be fallible. Our telemetry showed users approved roughly 93% of permission prompts. The more approvals a user sees, the less attention they pay to each, becoming over time much less diligent in their supervision. We recently built Claude Code auto mode, which automates safer approvals in order to reduce this approval fatigue. Still, vulnerabilities remain—any probabilistic defense has a non-zero miss rate. 1
The second approach to capping the blast radius—and the focus of much of this post—is containment. Rather than supervising what the agent does, we supervise what it’s able to do by enforcing access boundaries through, for example, sandboxes, virtual machines, and egress controls. This is where Anthropic engineering has devoted the most effort, and also where many of the most surprising security failures have occurred.
Over the past two years, we’ve shipped three primary agentic products: claude.ai , Claude Code, and Claude Cowork. Each serves a different audience, requiring a different containment architecture. This article shares what’s held up, what’s broken, and what we’ve learned about agent security along the way.
Three types of risk, three components of defense
Security risks to agents fall into one of three categories:
User misuse: A user—either maliciously or through carelessness—directs the agent to do something harmful. This includes everything from asking the agent to bypass a check they find annoying, to running a destructive command they don’t understand, to specifying intentional harm.
Model misbehavior: The agent takes a harmful action no one asked for. As our models have improved, they have become more aligned on most behavior evaluations, but this doesn’t mean risk necessarily shrinks. Less capable models are more likely to misread a situation and make obvious errors. More capable models make fewer mistakes, but they’re also better at finding unexpected paths to a goal, often by routing around restrictions nobody thought to write down.
At Anthropic, we’ve seen Claude models “helpfully” escape a sandbox in order to complete a task, examine git history to find answers to a coding test , and spontaneously identify the benchmark it was being run on in order to decrypt its answer key . Each model brings a new set of capabilities that are sometimes put to work in unexpected ways.
External attackers: The agent is attacked through external vectors such as tools, files, or network access. This category includes both prompt injection and conventional attacks on the agent's runtime, orchestration layer, or proxy.
When building containment and defense systems, we apply defenses to three main components:
The environment in which the agent runs. We constrain where and how an agent can act with process sandboxes, VMs, filesystem boundaries, and egress controls. The goal is to set a hard boundary on what an agent can reach. For example, if credentials never enter the sandbox, they can't be exfiltrated, regardless of whether the cause is a user, a model finding a “creative” path, or an attacker.
A tight perimeter also means you can relax oversight. Claude Code’s reference devcontainer exists precisely so that the agent can run unattended, without per-action approvals.
The model the agent consults. The mechanisms here include system prompts, classifiers, probes, and training modifications. Because models are probabilistic, these shape only what the agent tends to do, not what it is theoretically capable of doing.
These defenses are strong. On Gray Swan's Agent Red Teaming benchmark, which tests susceptibility to prompt injection, Claude Opus 4.7 holds attack success to roughly 0.1% on single attempts, and around 5–6% after 100 adaptive attempts. Claude Code auto mode catches roughly 83% of overeager behaviors before they execute . Yet even with best-in-class defenses, protection in the model layer will never be 100% effective, which is why it can't stand alone.
The external content the agent can reach. MCP servers, third-party plugins, and web search tools all feed content into the agent’s context from sources you don’t control. An audited connector isn’t the same as audited data—a GitHub connector, for instance, can load a poisoned README straight into the model’s context despite passing malware checks. Granularly limiting tool permissions can help limit the blast radius. An agent with read-only DB access, for instance, can be deployed far more broadly than one that writes to prod.
Defenses should overlap and complement each other. When environmental defenses aren’t available, the model layer has to pick up the slack (this is precisely what Claude Code’s auto mode is designed for). Locally, the environment and model defenses can guard against malicious tool outputs, but defenses can be added higher up the chain by limiting the tool’s capabilities and access.
Patterns for containing agents
Focusing on the environment layer, we describe three isolation patterns and how they’re tailored for each Claude platform— claude.ai , Claude Code, and Cowork. We arrived at each design gradually, after finding the balance between the capabilities we need from the agent and the degree of intervention required from the user.
Pattern 1: The ephemeral container (claude.ai code execution)
Though best known as a chat interface, claude.ai also writes and runs code, generates files, and calls connectors. When Claude runs code inside claude.ai, it does so in a gVisor container on isolated infrastructure. The agent is entirely server-side; no code runs on the local machine, and the filesystem is ephemeral (per-session). The blast radius is minimal, but so is the ceiling on what Claude can do—there's no persistent workspace and no access to the user's filesystem.
This also makes claude.ai subject to a more traditional threat model. We're not protecting user machines from agents; we're protecting our own infrastructure and each tenant from one another. Our pre-launch work for claude.ai was dominated by traditional security work like network configuration, internal service auth, and orchestration.
That work reinforced the oldest lesson in security: the weakest layer is the one you built yourself. gVisor and seccomp have been hardened against well-resourced adversaries for far longer than agentic AI has existed, so the review effort went into the newer pieces we'd built around them. We’ll come back to this later, since our custom proxy is also the piece that broke in our most consequential incident.
Pattern 2: The human-in-the-loop sandbox (Claude Code)
Claude Code runs on a user's machine and has access to their filesystem, shell, and network. Without this, coding agents have limited usefulness, so it’s imperative to find a way to grant that access safely.
One approach is to rely on a human-in-the-loop. This is only a tractable solution for Claude Code because the average user is a developer who’s familiar with coding environments: they can read bash, they understand what rm -rf does, and they already run npm install from untrusted sources several times a week. All that means that when an “allow this” dialog pops up, they are highly likely to have the expertise to accurately evaluate what the agent is attempting to do and the risk involved. Given this, Claude Code launched with the simplest possible defense: allow reads, require approval for write, bash, and network access.
However, as mentioned, approval fatigue showed up within weeks. Ironically, this meant that a feature originally designed to provide oversight could arguably have the opposite effect—some users might simply stop paying attention. As a first step to mitigate incautious approvals, we shipped an OS-level sandbox (Seatbelt on macOS, bubblewrap on Linux) that hardens the boundary: reads are allowed, writes are allowed inside the workspace, but network is denied by default. Within the sandbox, the agent runs largely without interruption. The result was an 84% reduction in permission prompts, and we open-sourced the runtime , so the boundary is auditable.
Our anonymized usage data also showed that experienced users auto-approve roughly twice as often as new users, but they also interrupt the agent mid-execution more frequently. Instead of gating individual steps, experienced users are more likely to supervise the agent only when it goes off track. While this may be a natural evolution in how people prefer to work with agents, this too is fallible, requiring users to be technical and attentive enough to notice drift in the first place. As model capabilities improve and agents begin writing increasingly ambitious bash, it becomes harder to notice any such drift. And as users move to multi-agent systems, this approach is also much less likely to be an effective oversight strategy.
Between mid-2025 and January 2026, we received reports of vulnerabilities in Claude Code through our responsible disclosure program. Three exploited code that executes before the user has consented to anything. To understand how this is possible, consider the most direct case: a developer clones a repository to review a pull request, and that repository contains a .claude/settings.json which defines a hook. Because Claude Code reads project settings during startup—before presenting the standard "Do you trust this folder?" prompt—the hook the attacker had authored and committed would execute automatically. The remaining cases looked structurally similar, in which input from the not-yet-trusted directory was parsed before the trust boundary had been established.
The fix in each case had the same shape: defer parsing and execution of project-local configuration until after the user accepts the trust prompt. If you're building something similar, treat project-open, config-load, and localhost listeners the way you'd treat any inbound request from the internet. They shouldn’t be implicitly trusted just because they feel local and arrive before the user has consented.
In February 2026, during a controlled internal red-team exercise, a researcher successfully phished an employee into launching Claude Code with a malicious prompt. The phish looked like ordinary collaboration—a "can you run this for me?" email with a ready-to-paste prompt attached—and the prompt itself read like routine task instructions. But somewhere among the setup steps, it gently asked Claude to read ~/.aws/credentials, encode the contents, and POST them to an external endpoint. Across 25 retries of that prompt, Claude completed the exfiltration 24 times.
This is a direct prompt injection—the attacker's instructions arrived throu

[truncated]
