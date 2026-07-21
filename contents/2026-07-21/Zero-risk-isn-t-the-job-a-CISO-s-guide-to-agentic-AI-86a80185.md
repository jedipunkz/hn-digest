---
source: "https://claude.com/blog/ciso-guide-to-agentic-ai"
hn_url: "https://news.ycombinator.com/item?id=48989995"
title: "Zero risk isn't the job: a CISO's guide to agentic AI"
article_title: "CISO's guide to agentic AI | Claude by Anthropic"
author: "talboren"
captured_at: "2026-07-21T10:15:48Z"
capture_tool: "hn-digest"
hn_id: 48989995
score: 1
comments: 0
posted_at: "2026-07-21T09:30:47Z"
tags:
  - hacker-news
  - translated
---

# Zero risk isn't the job: a CISO's guide to agentic AI

- HN: [48989995](https://news.ycombinator.com/item?id=48989995)
- Source: [claude.com](https://claude.com/blog/ciso-guide-to-agentic-ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T09:30:47Z

## Translation

タイトル: リスクゼロが仕事ではない: CISO によるエージェント AI ガイド
記事のタイトル: CISO のエージェント AI ガイド |クロード by Anthropic
説明: Anthropic の副 CISO は、エージェント AI のリスクを評価するための 4 つの質問フレームワークを共有し、エージェントの導入を制限して監査可能に保つ制御について説明します。

記事本文:
CISO 向けエージェント AI ガイド |クロード by Anthropic
クロード製品のご紹介 クロード
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者に連絡する 営業担当者に連絡する 営業担当者に問い合わせる
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
ゼロリスクが仕事ではない: CISO によるエージェント AI ガイド
ゼロリスクが仕事ではない: CISO によるエージェント AI ガイド
Anthropic の副 CISO である Jason Clinton は、エージェント AI の導入でチームが得た教訓と、エージェントを安全に構築および展開するためにチームが開発したリスク評価フレームワークについて共有します。
共有 リンクをコピー https://claude.com/blog/ciso-guide-to-agentic-ai
セキュリティ リーダーは、数か月前には存在すらしなかったエージェント AI のユースケースを承認するよう求められています。取締役会は、そのいずれかが管理されているかどうかを知りたいと考えていますが、組織のどこかで、従業員がすでに何も言わずにエージェントを何かに接続していることがあります。
これらのリクエストに「ノー」と答えると、テレメトリがゼロになり、通常はオフ スイッチが存在しないシャドウ導入が生成されます。コントロールなしで「はい」と言うとインシデントが発生し、会社でエージェントによる最初の重大なインシデントが発生すると、AI プログラムが後戻りしてしまいます。
エージェント AI の時代における CISO の責任は、ゼロリスクを達成することではありません。代わりに、私たちの仕事は、エージェントのリスクを読みやすく、制限したものにすることです。こうすることで、私たちが管理できることを意図的に受け入れることができるため、ビジネスは私たちに合わせてではなく、私たちの都合に合わせて進むようになります。
この記事では、eva のフレームワークを共有します。

セキュリティ リスクについてエージェントを評価し、「制限付き」が実際に何を意味するのかを説明し、私たちの仕事がどこに向かっているのかをプレビューします。
AI による外部リスクとミトス以降の内部リスク
以前のブログ投稿 で、同僚と私は、AI によって脆弱性が存在してから悪用されるまでの時間がどのように短縮されているかを共有し、組織がこれらのリスクを軽減する方法を強調しました。今後数か月以内に、場合によっては何年もコード内に気付かれずに眠っていた膨大な数のバグが AI モデルによって発見され、実用的なエクスプロイトに連鎖されると予想されます。 Claude Mythos Preview や Claude Mythos 5 などのフロンティア モデルでは、OpenBSD、Linux カーネル、Mozilla Firefox など、何年にもわたる人間によるレビューで見逃された重大な脆弱性がすでに発見されています。
これらは、GRC プログラムにとって重大なリスクです。脆弱性のギャップを緩和して埋めること、そして来るべきエクスプロイトの波に備えることは最優先事項である必要があります。このトピックについては、別のドキュメント「AI 加速攻撃に備えたセキュリティ プログラムの準備」を用意しました。このガイドでは内部リスクに焦点を当てます。
多くの組織にとって、エージェント システムに対する最も可能性の高い脅威ベクトルは、監視が不十分な個人エージェントを介して異種システムを接続することによって引き起こされるデータ漏洩です。もう 1 つの懸念は、プロンプト インジェクションです。攻撃者はエージェントが読み取るコンテンツ内に指示を隠し、エージェントはユーザーの代わりに攻撃者に従います。モデルの防御がどれほど堅牢かによっては、信頼できないコンテンツに触れるエージェントが暴露される可能性があります。モデルの能力が高まるにつれて、注入に対する耐性も大幅に向上しています。攻撃の成功率は低下し続けていますが、ゼロではありません。これら 2 つの例以外にも多くの懸念があり、新たな種類の懸念が氾濫しています。

圧倒されるように思えるかもしれません。
エージェントのユースケースが当社のレビュープロセスに到達すると、次の 4 つの質問をしてそのリスクを評価します。
どのような信頼できないコンテンツが取り込まれますか?信頼できないとは、外部の電子メール、オープン Web、サードパーティのドキュメント、パブリック リポジトリなど、攻撃者がおそらく書き込みまたは変更できるあらゆるものを意味します。答えが「何もない」の場合、エージェント固有のリスクはほぼゼロであるため、すぐに行動する必要があります。
誰に代わってどのようなアクションを実行できますか?読み取り専用は、読み取り/書き込みとは異なる懸念事項です。ツールの呼び出し、コードの実行、ネットワークの出力によって、それぞれの可能性が広がります。すべてのアクションは何らかのアイデンティティの下で行われるため、誰が誰であるかを知る必要があります。
位置がずれている場合の爆発半径はどれくらいですか?スコープ X の重大度は簡単に計算できます。不正行為者または調整インシデントは 1 つのファイルまたは組織全体にアクセスできましたか?それは異常でしょうか、迷惑でしょうか、データ漏えいでしょうか、それとも実際の事件でしょうか?
私はどのような可観測性を持っていますか?エージェントのアクションとユーザーのアクションを区別できますか?あなたの SIEM に到達しますか?
これらの質問に対する 4 つの回答からリスクの全体像がわかりますが、最小代理店の原則により、その対処方法がわかります。つまり、タスクを完了できる最も狭い機能を付与することです (詳細については、AI エージェント向けゼロトラスト ホワイト ペーパーを参照してください)。 Anthropic のデフォルトの姿勢は、管理者のペースで展開することです。つまり、少人数のグループを有効にして、テレメトリを監視し、アクセスを拡張します。これらの質問を、危険なエージェント システムについて考えるための新しいパラダイムに当てはめてください。
ユーザーの意図とずれたエージェントは、内部関係者による攻撃と区別できません。セキュリティ業界は 2019 年から 2022 年にかけて、インサイダー リスクを境界防御とは異なる分野として正式に定式化しました。システム内で最も危険な外部攻撃の媒介者は、多くの場合、既に攻撃を行っている誰かを危険にさらすものであるという認識でした。

正当なアクセスとして。
運用上の違いは応答時間です。Ponemon Institute の 2026 年のインサイダー リスクのコストに関するレポートでは、専用のインサイダー リスク プログラムに何年も投資した後でも、組織がインサイダー インシデントを封じ込めるのに平均 67 日かかったことがわかりました。エージェントの実行速度では、数日で測定される応答は長すぎます。
私たちが展開するものはすべて、アイデンティティ アクセス モデルの範囲の両端の一方に位置します。
一方の端には、システム サービス アカウントがあります。これは、人間の ID が付加されず、ビジネスのために 1 つのことだけを実行する自己完結型、単一目的、最小限の特権の ID です。インシデント対応エージェント (以下を参照)、チケット優先順位付けエージェント、または自律的なコードレビュー担当者がその例です。もう 1 つの例は、Claude Tag です。これは、Claude でタグ付けすることで、人間のチームが Slack などの共有ワークスペースでエージェントと共同作業できるようにする、新しい共有ワークスペース エージェントです。
もう一方の端には人間の資格情報があります。従業員がラップトップでチャット インターフェイスや Claude Cowork のようなパーソナル エージェント ハーネスを使用する場合、キーボードを操作する人は、自分の資格情報を使用して行われた他の操作に対して責任を負うのと同じように、結果に対して責任を負います。
スペクトルの中間、つまりエージェントが個人の委任されたアイデンティティを、その人が監視していないシステムに持ち込む部分では、責任が曖昧になります。責任の所在が曖昧であると、事件は説明不能になります。
ユーザーの意図とずれたエージェントは、内部関係者による攻撃と区別できません。セキュリティ業界は 2019 年から 2022 年にかけて、システム内で最も危険な外部攻撃ベクトルは、すでに正当なアクセス権を持つユーザーを侵害するものであることが多いと認識し、インサイダー リスクを境界防御とは異なる分野として正式に定式化しました。
ポネモン研究所の 2026 年のインサイダー リスクのコスト

ポートの調査によると、専用のインサイダー リスク プログラムに何年も投資してきたにもかかわらず、組織はインサイダー インシデントを封じ込めるまでに平均 67 日かかったことが判明しました。エージェントの実行速度では、67 日という測定単位は完全に間違っています。
ケーススタディ: インシデント対応エージェント
1 年以上前、私たちはインシデント対応プロセスについてクロードに指摘しました。本番アプリケーションのオンコールを担当したことのある人なら誰でも問題を知っています。午前 2 時にセキュリティ インシデントについての電話があり、インシデント対応チャネルを立ち上げ、適切な人材を集めて仕事に取り掛かります。このプロセスは退屈で、大量の文書が必要で、時間がかかります。ただし、運用環境のコードベースに関する適切なコンテキストがあれば、その大部分は自動化できます。
そこで、それを行うためのエージェントを構築しました。私たちはエージェントに 3 つのツールへのアクセス権を与えました。1 つは PII を含まない運用ログへの読み取り専用アクセスです。 Slack にアクセスしてインシデント チャネルを開き、プロセスを実行します。インシデントの解決後に事後分析用に Google ドキュメントを作成する機能。
次の 4 つの質問について答えていきました。
信頼できないコンテンツ: なし。入力は私たち自身のログと私たち自身の内部 Slack であり、どちらも信頼境界内にあるため、注入には匿名の攻撃者ではなく、内部関係者または侵害されたアカウントが必要になります。
アクション: どこでも読み取り、書き込みは新しいドキュメントと Slack メッセージに限定されます。編集や削除、権限の変更、外部エンドポイントはありません。
爆発範囲: 私たちが構築できた最悪の結果は、すでにロックダウンされているインシデントチャネルに投稿された、軽度の機密ログ行でした。
可観測性: すべてのアクションが SIEM に到達するため、予期せぬものは数週間ではなく数分で表面化します。
エージェントにはリスクがなかったわけではありませんが、完全な監査カバレッジを備えた制限された書き込みサーフェス上で動作していたので、リスクの面で有利でした。

ファイルは使いやすかったです。
ただし、この話には興味深い補足があります。モデルがリリースされるたびに、エージェントはより賢くなっています。 2025 年 11 月に、このエージェントを Claude Opus 4 から Claude Opus 4.5 に移行しましたが、他には何も変更しませんでした。新しいツール、権限、プロンプトは何も変更しませんでした。この直後、初めて、インシデントの最中にエージェントは、インシデントの最中に、スタック トレースで根本原因をすでに見つけており、まだ到着していない人間がいない場合は、コード変更を生成するための適切なコード アクセスを持つ別のエージェントに連絡することで、独自に運用を修正できることを認識するのに十分でした。
事後、私たちはログを確認しました。思考追跡でこれが機能するのを観察しました。私は頼まれたことを実行しました。人間はここにはいません。問題を解決したらどうなるでしょうか? Anthropic の内部には、コードの変更を記述し、人間によるレビューのためにアップロードできる、Claude Tag のようなテクノロジーの内部バリアントがあります。独自に、Slack 経由でこの Claude Tag のようなインスタンスに連絡し、修正を作成するよう依頼しました。修正はプルリクエストに適用され、本番環境にプッシュする前に人間がレビューしました。
この緊急のエージェント間の通信によって拡大された爆発範囲自体は、私たちの原則によって管理されていました。起こり得る最悪の事態は、運用ログ行を含むコード変更がアップロードされることです。このエージェント間のコミュニケーションは現在、インシデント対応の根本原因と修復実践の定期的な一部となっています。すべて人間による監視が行われます。
この新たな行動は私たちに 2 つのことを教えてくれました。 1 つ目: 新しい機能は、エージェント展開の境界内で現れる可能性があります。現在のモデルの制限を超えて制限するのではなく、アクセスとアクションを制限することが重要です。 2番目: コントロール

ols はこのような確率的エージェントでも効果的です。新しい動作は Slack チャネルで発生したため、人間によるループであり、唯一の書き込みに似たアクションは引き続き人間によるレビューが必要でした。
現在、インシデント対応以外では、チャット チャネル内で人間がオンザループで作業するエージェント間のコミュニケーションが標準となっています。
インシデント対応エージェントは、制限されたサービス アカウントで 1 つのジョブを実行するサービス アカウントです。クロード・コワークは、人間のオペレーター側の立場にあります。キーボードを操作する従業員が結果に対して責任を負い、エージェントが従業員に代わって、従業員が承認した (ますます) クラウドで実行されるシステム内で動作します。
Claude Cowork の脅威モデルは単純です。エージェントは基本的に、ローカルまたはホストされたインターフェイス内で実行される Claude Code であるためです。デスクトップ アプリは、ローカル ファイル アクセス、ブラウザーの使用、およびコンピューターの使用に引き続き必要です。これらの機能はローカル マシンに直接到達するため、アプリがそうする必要があります。したがって、システム全体の表面は、オーケストレーション、MCP 呼び出し、送信ネットワーク要求を処理する (おそらくリモート) 実行環境と、ファイルと画面にアクセスするためのローカル ブリッジの 2 つの部分で構成されます。
上で概説した 4 つの質問は、クロード カウごとに異なる答えを導き出します。

[切り捨てられた]

## Original Extract

Anthropic's Deputy CISO shares a four-question framework for assessing agentic AI risk, and walks through controls that keep agent deployments bounded and auditable.

CISO's guide to agentic AI | Claude by Anthropic
Meet Claude Products Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Zero risk isn't the job: a CISO's guide to agentic AI
Zero risk isn't the job: a CISO's guide to agentic AI
Anthropic's Deputy CISO, Jason Clinton, shares his team's lessons learned adopting agentic AI, and the risk assessment framework they've developed for building and deploying agents securely.
Share Copy link https://claude.com/blog/ciso-guide-to-agentic-ai
Security leaders are being asked to approve agentic AI use cases that did not even exist a few months ago. Boards want to know whether any of it is governed, and somewhere in your organization, an employee has already connected an agent to something without telling you.
Saying “no” to these requests produces shadow adoption, which has zero telemetry and generally no off switch. Saying “yes” without controls produces incidents, and the first serious agent incident at your company will set your AI program back.
A CISO’s responsibility in the age of agentic AI is not to achieve zero risk. Instead, our jobs are to make agentic risk legible and bounded. This way, we can deliberately accept what we can manage, so the business moves on our terms instead of around us.
In this article, I share our framework for evaluating agents for security risk, explain what “bounded” means in practice, and preview where our work is headed.
External risk from AI versus internal risk in the post-Mythos era
In an earlier blog post , my colleagues and I shared how AI is collapsing the time between a vulnerability existing and a working exploit, highlighting how organizations can mitigate these risks. In the coming months, we expect that vast numbers of bugs that have sat unnoticed in code, sometimes for years, will be found by AI models and chained into working exploits. Frontier models like Claude Mythos Preview and Claude Mythos 5 are already finding serious vulnerabilities that years of human review missed, including in OpenBSD, the Linux Kernel and Mozilla Firefox .
These are serious risks to any GRC program. Mitigating and closing vulnerability gaps, as well as for preparing for the coming wave of exploits, should be a top priority. For this topic, we have prepared a separate doc: Preparing your security program for AI-accelerated offense . We’ll focus on internal risks for this guide.
For many organizations, the most likely threat vector for agentic systems is a data leak enabled by connecting disparate systems through personal agents with insufficient oversight. Another concern is prompt injection : an attacker hides instructions inside content the agent reads, and the agent follows the attacker instead of the user. Any agent that touches untrusted content could then be exposed, depending on how robust the defenses of the model are. As models grow increasingly capable, they’re getting meaningfully better at resisting injection. While attack success rates keep falling , they’re not zero. There are many concerns outside of these two examples, and the deluge of new classes of concern can seem overwhelming.
When an agentic use case reaches our review process, we assess its risk by asking four questions:
What untrusted content does it ingest? Untrusted means anything an attacker could plausibly write or alter, including outside email, the open web, third-party documents, or public repositories. If the answer is "nothing," the agent-specific risk is near zero and you should move quickly.
What actions can it take, and on whose behalf? Read-only is a different concern from read/write. Tool calls, code execution, and network egress each widen the aperture. Every action happens under some identity, and you need to know whose.
What is the blast radius if it is misaligned? Scope X severity is the quick calculation: did the bad actor or alignment incident have access to one file or the whole org? Would it be an anomaly, an annoyance, a data exposure, or a true incident?
What observability do I have? Can you tell agent actions from user actions? Does it land in your SIEM?
The four answers to these questions give you a picture of your risk, but the principle of least agency tells you what to do with it: grant the narrowest capability that still completes the task (see our Zero Trust for AI Agents white paper to learn more). Our default posture at Anthropic is admin-paced rollout: enable a small group, watch the telemetry, and then expand access. Apply these questions to a new paradigm for thinking about risky agentic systems.
An agent that drifts out of alignment with your intent is indistinguishable from an insider attack. The security industry spent 2019-2022 formalizing insider risk as a discipline distinct from perimeter defense—recognizing that the most dangerous external attack vectoractor in a system is often one that compromises someone who already has legitimate access.
The operational difference is response time: Ponemon Institute's 2026 Cost of Insider Risks report found organizations took an average of 67 days to contain an insider incident—even after years of investment in dedicated insider risk programs. At agent execution speeds, responses measured in days are too long.
Everything we deploy sits at one of two ends of an identity access model spectrum.
At one end is the system service account : a self-contained, single-purpose, least-privilege identity that does exactly one thing for the business, with no human identity attached. The incident-response agent (see below), a ticket triage agent, or an autonomous code reviewer are examples of these. Another example is Claude Tag , our new shared workspace agent that lets human teams collaborate with agents in shared workspaces like Slack by tagging in Claude.
At the other end is the human credential . When an employee uses a chat interface or a personal agent harness like Claude Cowork on their laptop, the person at the keyboard is accountable for the outcome, the same way they are accountable for anything else done with their credentials.
The middle of the spectrum, where an agent carries a person's delegated identity into systems that person is not watching, is where accountability gets ambiguous. Ambiguous accountability is how incidents become unexplainable.
An agent that drifts out of alignment with your intent is indistinguishable from an insider attack. The security industry spent 2019-2022 formalizing insider risk as a discipline distinct from perimeter defense—recognizing that the most dangerous external attack vector in a system is often one that compromises someone who already has legitimate access.
Ponemon Institute's 2026 Cost of Insider Risks report found organizations took an average of 67 days to contain an insider incident—even after years of investment in dedicated insider risk programs. At agent execution speeds, 67 days is the wrong unit of measurement entirely.
Case study: an incident response agent
More than a year ago, we pointed Claude at our incident response process. Anyone who has been on-call for a production application knows the problem: you’re paged at 2 a.m. about a security incident, you spin up an incident response channel, you pull in the right people, and get to work. This process is tedious, documentation-heavy, and fast-moving. But, with the right context about your production environment codebase, the majority of it can be automated.
So we built an agent to do it. We gave the agent access to three tools: read-only access to our production logs, which contain no PII; access to Slack, to open the incident channel and run the process; and the ability to draft a Google Doc for the postmortem after the incident is resolved.
We ran it through the four questions:
Untrusted content: none. The inputs were our own logs and our own internal Slack, both inside the trust boundary, so an injection would require an insider or a compromised account rather than an anonymous attacker.
Actions: reads everywhere, writes limited to new documents and Slack messages. No edits or deletes, no permission changes, no external endpoints.
Blast radius: the worst outcome we could construct was some mildly sensitive log lines posted into an incident channel that was already locked down.
Observability: every action landed in our SIEM, so anything unexpected would surface in minutes, not weeks.
While the agent wasn’t risk-free, it operated on a bounded write surface with full audit coverage, which was a risk profile we were comfortable with.
However, there’s an interesting addendum to this story: with each model release, the agent got smarter. In November 2025, we moved this agent from Claude Opus 4 to Claude Opus 4.5 and changed nothing else—no new tools, permissions, or prompts. Immediately after this, for the first time, the intelligence uplift alone was enough for the agent to notice, mid-incident, that it had already found the root cause in a stack trace and that, in the absence of the human who hadn't arrived yet, it could try to fix production on its own by reaching out to another agent that had the appropriate code access to produce the code change.
Post hoc, we reviewed logs: we watched it work through this in the thinking traces: I have done what I was asked to do. The human is not here. What if I fixed the problem? Inside of Anthropic we have an internal variant of Claude Tag-like technology which can write code changes and upload them for human review. On its own, it reached out over Slack to this Claude Tag-like instance and asked it to write the fix. The fix went to a pull request that a human reviewed before pushing it to production.
The expanded blast radius that came from this emergent agent-to-agent communication was itself governed by our principles: the worst that could happen would be that a code change would be uploaded which contained a production log line. This agent-to-agent communication is now a regular part of our incidence response root cause and remediation practices; all with human-on-the-loop monitoring.
This emergent behavior taught us two things. First: new capabilities can show up within the boundaries of an agent deployment. It’s important to limit access and actions, not around what you believed today's model limits are. Second: controls are effective even with stochastic agents like this. The new behavior was human-on-the-loop because it happened in a Slack channel, and the only write-like action still required a human review.
Today, outside of incidence response, agent-to-agent communication within chat channels, with human on-the-loop where people work, is the norm.
The incident response agent is a service account doing one job, in a bounded service account. Claude Cowork is at the human operator end of the spectrum: an employee at a keyboard is accountable for the outcome, and the agent then acts on their behalf, in systems they authorized—increasingly—running in the cloud.
Claude Cowork's threat model is straightforward, because the agent is essentially Claude Code running either locally or inside a hosted interface. The desktop app remains required for local file access, browser use, and computer use; those capabilities reach the local machine directly and need the app to do so. The full system surface is therefore two-part: a (possibly remote) execution environment handling orchestration, MCP calls, and outbound network requests, and a local bridge for file and screen access.
The four questions outlined above produce different answers for every Claude Cow

[truncated]
