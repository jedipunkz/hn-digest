---
source: "https://www.nofire.ai/blog/who-audits-the-ai-agent"
hn_url: "https://news.ycombinator.com/item?id=49194708"
title: "AI Agents Governance: Who guards the guardrails?"
article_title: "Quis custodiet ipsos custodes? | NOFire"
author: "spirosoik"
captured_at: "2026-08-06T10:28:24Z"
capture_tool: "hn-digest"
hn_id: 49194708
score: 3
comments: 0
posted_at: "2026-08-06T10:07:04Z"
tags:
  - hacker-news
  - translated
---

# AI Agents Governance: Who guards the guardrails?

- HN: [49194708](https://news.ycombinator.com/item?id=49194708)
- Source: [www.nofire.ai](https://www.nofire.ai/blog/who-audits-the-ai-agent)
- Score: 3
- Comments: 0
- Posted: 2026-08-06T10:07:04Z

## Translation

タイトル: AI エージェントのガバナンス: ガードレールを守るのは誰ですか?
記事のタイトル: 管理は必要ですか? |火災なし
説明: すべてのエージェント プラットフォームにはガードレールが付属しています。誰が見ているか尋ねる人はほとんどいません。サンドボックスの外側からエージェントを観察する際の信頼境界、エージェントが到達できない有利な地点から最強の監査証跡が書き込まれる理由、生の境界イベントをどのように運用に持ち込むのかについて説明します。
[切り捨てられた]

記事本文:
管理者はイプソス管理者ですか? | NOFire 製品概要 運用インシデントのコンテキストと制御モデル 根本原因を数分で記憶し、次回思い出す Prevent 危険な変更を出荷前にキャッチ AI エージェント ガバナンス ポリシーに対するすべてのエージェント アクションをゲートする サービス カタログ 自動保守カタログの各ページが Always-On ワークフローで実行される 定期的な運用作業が自動的に実行される データ ガバナンス リネージ、スキーマ変更、遅延、今日の Kafka 顧客 リソース リソース ベンチマーク、ガイド、参考資料 ブログ 運用イベントのフィールド ノートトーク、デモ、ポッドキャスト ドキュメント API リファレンス & セットアップ ガイド ↗ 統合 プロダクションを接続する 会社について ミッションとチーム セキュリティ 信頼、コンプライアンスとアーキテクチャ キャリア 価値観と自由な役割 お問い合わせ チームに連絡する デモを予約する 今すぐ試してみる ブログ / エンジニアリング Engineering 04 Aug 2026 5 min read Quis custodiet ipsos custodes?
すべてのエージェント プラットフォームにはガードレールが付属しています。誰が見ているか尋ねる人はほとんどいません。サンドボックスの外側からエージェントを観察する信頼境界、エージェントが到達できない有利な地点から最強の監査証跡が書き込まれる理由、生の境界イベントを人間が実際に読み取ることができる操作にどのように組み込むかについて説明します。
ガードレールの信頼性は、それが走行する有利な地点によって決まります。オブザーバー自身を含め、エージェントの爆発範囲内で実行されるものはすべて、裁判官ではなく証人です。
ハードウェア VM 境界は、エージェントが自身の履歴を書き換えることができない唯一の場所です。ゲストが協力するかどうかに関係なく、すべてのパケット、すべての出口、すべての I/O リクエストが境界を越えます。
境界証拠はそれ自体では意味的に不十分であるため、それをエージェントが実際に実行する操作の分類に組み入れます。監査証跡は物語のように読めますが、すべての行はエージェントが実行できる証拠によって裏付けられています。

鍛造していない。
管理者はイプソス管理者ですか?誰が警備員自身を守るのでしょうか？
問題は古いです。プラトンは共和国でこの問題を提起しました。都市を守るために守護者を任命したら、誰がその守護者を牽制するのでしょうか？私たちがそれを覚えているラテン語のタグは、後にローマの詩人ユウェナリスから来たものです。
本番環境の AI エージェントにとって、この質問はもはや修辞的なものではありません。現在、すべてのプラットフォームに、ポリシー プロンプト、モデレーション レイヤー、ホワイトリスト、「安全」サイドカーなどのガードレールが搭載されています。それらはすべてソフトウェアです。ソフトウェアはどこかで実行されます。そして、ガードレールがどこにあるのかによって、そのメッセージを信じられるかどうかが決まります。
エージェントがデフォルトで信頼できない理由と、それが雰囲気ではなく実際のコントロールであるためにサンドボックスが何を意味するのかについては、以前に書きました。これらの投稿は主に、ハードウェアで分離された microVM (タスクごとに 1 つ、 urunc 経由) の強制に関するものでした。この投稿は問題の残り半分、あまり注目されていない部分に関するものです。
エージェントを停止しません。それが何をしたのかを知ること。
ガードレールはソフトウェアであり、ソフトウェアはどこかで実行されます
ここで私たちが注目する故障モードを示します。エージェントは実稼働環境への資格情報を保持します。何か問題が発生した場合、または監査人が単に「このエージェントは先週の火曜日に正確に何をしましたか?」と尋ねるだけです。
明らかな答えは、エージェント自身のトランスクリプト、つまりツールの呼び出し、ログ、セッションのナレーションを読むことです。その記録は内容が豊富で読みやすいが、証拠としての価値はまったくない。これは、攻撃者が注入できるのと同じコンテキスト内で、監査しているものによって書き込まれます。エージェントが侵害された場合、または単に間違っていた場合、記録が最初の犠牲者となります。
次の答えは、同じコンテナまたは同じカーネル上に、エージェントの隣にモニターを配置することです。改善されましたが、モニターは監視対象のワークロードと運命を共にすることになりました。爆発範囲内に住む警備員 c

守るものによって無効化されたり、騙されたり、競争されたりする。私たちはマニフェストでそれを述べましたが、それはここでも当てはまります。監視対象のエージェントと同じカーネルで実行されるガードレールは提案であり、境界ではありません。
このことから、必然的に次のような疑問が生じます。エージェントが到達できない有利な地点はあるのでしょうか?
3 つの有利なポイント、3 つの信頼レベル
私たちの考え方では、エージェントを観察できる場所は基本的に 3 つあり、それらははしごを形成しています。
エージェントのコンテキスト内。自己報告: トランスクリプト、ツールのログ、エージェント自身のセッションのアカウント。最も豊富なセマンティクス、ゼロトラスト。エージェントが読み書きできるものはすべて、攻撃者も読み書きできます。
ゲストの中。私たちの実行層は意図的にゲストを最小限に抑えています。ゲストからホストに戻る制御チャネルはなく、タスクが要求していないものは何も実行されていません。より豊富なテレメトリーを実現するために microVM 内にオブザーバーを配置する場合、オブザーバーは厳格なルールに基づいて操作されます。つまり、オブザーバーは外部に向けて一方向に話し、判断者ではなく証人として扱われます。 OS レベルの真実 (プロセス、ファイル、接続) を適切なセマンティクスで認識しますが、爆発範囲内に存在します。完全に感染したゲストは、原則として、嘘をついたり、それについて嘘をついたりすることができます。
VM 境界で。これが興味深いところです。エージェントが送信するすべてのパケット、エージェントが書き込むすべてのブロック、エージェントが出るすべての出口は、CPU が強制する境界を越えており、ゲストが触れることができるすべての外部のホスト側から記録されます。ゲストはこの記録を抑制したり、書き換えたりすることはできず、（これが私たちが気に入っている部分です）それがどれほど注意深く監視されているかさえ知ることができません。観測は受動的で帯域外です。ゲスト内で検出できるプローブはありません。
この取引に注目してください。はしごを下るにつれて、信頼は高まり、意味の豊かさは下がります。トランスクリプトはストーリーを伝えます。境界はyに伝える

真実を、誰も読みたくない言語で。
つまり、デザインは「レイヤーを選ぶ」ものではありません。それは、すべてのレイヤーを使用し、その有利な点で正当な範囲でのみ各レイヤーを信頼し、下位を裏付けることです。エージェントの話は仮説です。ゲスト内ビューはそれをサポートしているかサポートしていないかのどちらかです。境界記録によって解決されます。
この重層的な問題には名前が付けられ始めています。 Uber は最近、Agentic Detection and Response (ADR) を専門分野として正式に制定しました。これは、エージェント自身のローカル店舗からのプロンプト、推論チェーン、およびツール呼び出しをキャプチャするセンサーベースのテレメトリです。それは真の進歩であり、このカテゴリーは成長するでしょう。ただし、センサーがはしごのどこにあるのかに注目してください。センサーはエージェントのキャッシュを解析します。つまり、エージェントが書き込むことを選択したものを読み取ります。それがレベル 1 です。私たちが関心を持っている有利な地点、つまりエージェントがそこから抜け出す方法を説明できない地点は、まだ下にあります。
それがユベナルに対する正直な答えだ。ゲスト内オブザーバーがエージェントを監視します。ハードウェア境界はオブザーバーを保護します。最後のガードは誰でも話せるソフトウェアではありません。それは分離メカニズムそのものです。
最も有用なシグナルは意見の相違です
同じセッションを複数の有利な地点から観察すると、発散検出という素晴らしいものが無料で提供されます。
エージェントがテスト スイートを実行したと主張する場合、境界にはテスト スイートが生成するファイルとネットワーク アクティビティが表示されるはずです。トランスクリプトには「構成を読んだだけです」と書かれているが、境界には誰も言及していないアウトバウンド接続が示されている場合、エージェントが嘘をついているのか混乱しているのかを判断する必要はありません。いずれにせよ、エージェントが偽造できない証拠によってフラグが立てられた、信頼すべきではないセッションが存在することになります。
裏付けられた監査証跡では、ストーリーと境界現実との間のギャップはノイズではありません。それはあなたが持っている最も強い信号です。
もちろん誰も望んでいない

生の境界イベントを読み取ることでエージェントを監査します。あるアドレスへの TCP 接続とブロック I/O のバーストは真実ですが、「エージェントが何をしたのか?」に対する答えではありません。
これは VM イントロスペクションの古典的なセマンティック ギャップであり、ほとんどの帯域外アプローチが歴史的に諦めていた場所です。私たちの見解は、エージェントは実際にそれを解決するのに友好的なケースです。任意のソフトウェアとは異なり、エージェントは驚くほど規則的な一連の操作 (依存関係のフェッチ、リポジトリのクローン作成、ビルドの実行、API の呼び出しなど) を実行し、それらのそれぞれが境界に特徴的な形状を残します。
そこで私たちは、実際のエージェントのセッションから学習した、実際のエージェントが実際に実行する操作の分類を構築し、そこに境界証拠を取り込みました。私たちはこれをインテントリフティングと呼びます。つまり、操作ごとに追跡される信頼度を使用して、「接続と I/O のこのパターン」を「エージェントがこのレジストリからパッケージをインストールした」に変換します。その結果、すべての行が境界証拠によって裏付けられ、裏付けられない行がそのようにマークされることを除いて、トランスクリプトと同様の監査証跡が作成されます。
境界のみの証拠はどの程度意図を復元できるのでしょうか?正直言って、始めたときに予想していたよりも良かったです。私たちは方法論と数値を適切にまとめており、それらはそれぞれの記事に値するものであるため、今はそのままにしておきます。
これらすべてに対する通常の反対意見は頭上にありますが、有利な地点が私たちの代わりに仕事をしてくれるので、これは私たちが答えるのが最も楽しい反対意見です。
境界での観察は受動的です。エージェントはオブザーバーを待ちません。ホット パスに介入はなく、ワークロードのコード変更も、ゲストが使用する手段もありません。 microVM 自体はミリ秒単位でコールドスタートします (その点については以前に説明しました)。また、トップにかかるイントロスペクションの負担は十分に小さいため、

デフォルトでは、本番環境のすべてのタスクで実行するだけです。完全な測定値は上記の記事の一部です。簡単に言えば、サンドボックスを購入する余裕があれば、サンドボックスを視聴する余裕があるということです。
最初に監査、次に執行
私たちはここでは強制ではなく観察可能性を意図的に主導しており、それは回避ではありません。注文です。
強制はエージェントが何を実行できるかを決定します。私たちは、分離モデルからそれを強制するランタイムに至るまで、これらのコントロールがどこに存在する必要があるかについて書きました。しかし、施行はそれが実行される根拠となる真実と同程度に優れており、事後は監査証跡が根拠となる真実となります。監査人、インシデントレビュー、またはあなた自身の事後調査がエージェントが何をしたか尋ねたとき、「エージェントが私たちに伝えたことは次のとおりです」では答えになりません。 「これが、ハードウェアの境界を越えて運用に持ち込まれ、私たちに伝えられた内容と照合されたものです」です。
基本的に、エージェント セッションの事後分析は、障害の事後分析と同じ基準、つまり証拠が第一、物語が二番目に保たれることを望んでいます。
このすべての部分 (分類法、リフティングの精度、実際にレイヤーが一致しない場合に何が起こるか) についてはさらに説明すべきことがありますが、それについては今後数週間以内に説明する予定です。今すぐ全体像 (脅威モデル、隔離スコアカード、各主張の背後にある証拠) が必要な場合は、以下のホワイトペーパーにすべて記載されています。
管理者はイプソス管理者ですか?境界はそうなります。それはむしろそれを持つことのポイントです。
7つの主張。スタック上でそれらを確認してください。
完全なホワイト ペーパーには、脅威モデル、隔離スコアカード、および各主張の背後にある証拠が含まれています。または、創設者との電話を予約して、境界を実際に確認してください。
⌐ 続きを読む サンドボックス マニフェスト エンジニアリング · 5 分 → バグは本物だがパスが違う場合: ITScape (CVE-2026-46316) と最小限のケース

ランタイム エンジニアリング · 5 分 → コンテナ速度での VM 分離: urunc と Kubernetes エージェント サンドボックス CRD エンジニアリング · 5 分 → 実稼働用のコンテキストおよび制御モデル。
製品概要 サービス カタログ インシデント 常時稼働ワークフロー データ ガバナンス Prevent AI Agent ガバナンス統合 顧客 企業 セキュリティ ブログについて お問い合わせ リソース ライブラリ イベントとポッドキャスト AI SRE ベンチマーク AI 信頼性ガイド 用語集 ドキュメント ↗ © 2026 NOFire AI プライバシー ポリシー 利用規約 v.2026.05 デモを予約する デモを予約する
スタック上で確認してください。 NOFire は生産を継続します。
創業者による 30 分間のウォークスルー。ライブでスタックをコンテキスト & コントロール モデルにマッピングします。

## Original Extract

Every agent platform ships guardrails. Almost nobody asks who watches them. We walk through the trust boundaries of observing an agent from outside its sandbox, why the strongest audit trail is written from a vantage point the agent cannot reach, and how we lift raw boundary events into operations a
[truncated]

Quis custodiet ipsos custodes? | NOFire Product Overview The Context & Control Model for production Incidents Root cause in minutes, remembered next time Prevent Catch risky changes before they ship AI Agent Governance Gate every agent action against policy Service Catalog The self-maintaining catalog every page runs on Always-On Workflows The recurring ops work, running for you Data Governance Lineage, schema changes, and lag, Kafka today Customers Resources Resources Benchmarks, guides & reference material Blog Field notes from production Events Talks, demos & podcasts Docs API reference & setup guides ↗ Integrations Connect your production Company About Mission & team Security Trust, compliance & architecture Careers Values & open roles Contact us Get in touch with our team Book a demo Try now Blog / Engineering Engineering 04 Aug 2026 5 min read Quis custodiet ipsos custodes?
Every agent platform ships guardrails. Almost nobody asks who watches them. We walk through the trust boundaries of observing an agent from outside its sandbox, why the strongest audit trail is written from a vantage point the agent cannot reach, and how we lift raw boundary events into operations a human can actually read.
A guardrail is only as trustworthy as the vantage point it runs from. Anything executing inside the agent's blast radius, including the observer itself, is a witness, not a judge.
The hardware VM boundary is the one place the agent cannot rewrite its own history: every packet, every exit, every I/O request crosses it whether the guest cooperates or not.
Boundary evidence is semantically poor on its own, so we lift it into a taxonomy of operations agents actually perform. The audit trail reads like a story, but every line is backed by evidence the agent could not have forged.
Quis custodiet ipsos custodes? Who guards the guards themselves?
The problem is old. Plato raised it in the Republic : once you appoint guardians to protect the city, who keeps the guardians in check? The Latin tag we remember it by came later, from the Roman poet Juvenal.
For AI agents in production, the question is not rhetorical anymore. Every platform ships guardrails now: policy prompts, moderation layers, allowlists, "safety" sidecars. All of them are software. Software runs somewhere. And where a guardrail runs decides whether you can believe what it tells you.
We have written before about why the agent is untrusted by default and what a sandbox has to mean for that to be a real control rather than a vibe. Those posts were mostly about enforcement: hardware-isolated microVMs, one per task, via urunc . This post is about the other half of the problem, the one that gets far less attention.
Not stopping the agent. Knowing what it did.
Guardrails are software, and software runs somewhere
Here is the failure mode we care about. An agent holds credentials to your production environment. Something goes wrong, or an auditor simply asks: what exactly did this agent do last Tuesday?
The obvious answer is to read the agent's own transcript: its tool calls, its logs, its narration of the session. That record is rich and readable, and it is worth exactly nothing as evidence. It is written by the thing you are auditing, in the same context an attacker can inject into. If the agent was compromised, or just wrong, the transcript is the first casualty.
The next answer is to put a monitor next to the agent, in the same container or on the same kernel. Better, but the monitor now shares fate with the workload it polices. A guard that lives inside the blast radius can be disabled, deceived, or raced by the thing it guards. We said it in the manifesto and it holds here: a guardrail running in the same kernel as the agent it polices is a suggestion, not a boundary.
Which inevitably leads us to the question: is there a vantage point the agent cannot reach?
Three vantage points, three levels of trust
The way we think about it, there are essentially three places you can observe an agent from, and they form a ladder.
Inside the agent's context. Self-reporting: transcripts, tool logs, the agent's own account of the session. Richest semantics, zero trust. Anything the agent can read or write, an attacker can too.
Inside the guest. Our execution layer deliberately keeps the guest minimal: no control channel from the guest back into the host, nothing running that the task did not ask for. When we do place an observer inside the microVM for richer telemetry, it goes in under a strict rule: it talks one way, outward, and we treat it as a witness rather than a judge. It sees the OS-level truth (processes, files, connections) with good semantics, but it lives inside the blast radius. A fully compromised guest can, in principle, lie to it or about it.
At the VM boundary. This is the interesting one. Every packet the agent sends, every block it writes, every exit it takes crosses a boundary the CPU enforces, and we record it from the host side, outside anything the guest can touch. The guest cannot suppress this record, cannot rewrite it, and (this is the part we like) cannot even tell how closely it is being watched. Observation is passive and out-of-band; there is no probe inside the guest to find.
Notice the trade: as you move down the ladder, trust goes up and semantic richness goes down. The transcript tells you a story. The boundary tells you the truth, in a language nobody wants to read.
So the design is not "pick a layer". It is: use every layer, trust each one only as far as its vantage point warrants, and corroborate downward. The agent's story is a hypothesis. The in-guest view either supports it or it does not. The boundary record settles it.
This layered problem is starting to get a name. Uber recently formalized Agentic Detection and Response (ADR) as a discipline: sensor-based telemetry that captures prompts, reasoning chains, and tool calls from the agent's own local stores. It is genuine progress, and the category will grow. But notice where their sensor sits on the ladder: it parses the agent's caches, which means it reads what the agent chose to write. That is Level 1. The vantage point we care about, the one the agent cannot narrate its way out of, is still below.
And that is the honest answer to Juvenal. The in-guest observer guards the agent. The hardware boundary guards the observer. The last guard is not software anyone can talk to; it is the isolation mechanism itself.
The most useful signal is disagreement
Once you have the same session observed from multiple vantage points, something nice falls out for free: divergence detection.
If the agent claims it ran a test suite, the boundary should show the file and network activity a test suite produces. If the transcript says "I only read the config", but the boundary shows an outbound connection nobody mentioned, you do not need to decide whether the agent is lying or confused. Either way, you have a session that should not be trusted, flagged by evidence the agent had no way to forge.
In a corroborated audit trail, a gap between the story and the boundary reality is not noise. It is the strongest signal you have.
Of course, nobody wants to audit an agent by reading raw boundary events. A TCP connection to some address and a burst of block I/O is the truth, but it is not an answer to "what did the agent do?".
This is the classic semantic gap of VM introspection, and it is where most out-of-band approaches historically gave up. Our take: agents are actually a friendly case for closing it. Unlike arbitrary software, an agent works through a surprisingly regular set of operations (fetch a dependency, clone a repo, run a build, call an API, and so on), and each of those leaves a characteristic shape at the boundary.
So we built a taxonomy of the operations real agents actually perform, learned from real agent sessions, and we lift boundary evidence into it. We call this intent lifting : turning "this pattern of connections and I/O" into "the agent installed a package from this registry", with the confidence tracked per operation. The result is an audit trail that reads like the transcript, except every line is backed by boundary evidence, and the lines that cannot be backed are marked as such.
How well does boundary-only evidence recover intent? Better than we expected when we started, honestly. We are writing up the methodology and numbers properly, and they deserve their own post, so we will leave it at that for now.
The usual objection to all of this is overhead, and it is the objection we have the most fun answering, because the vantage point does the work for us.
Observation at the boundary is passive. The agent does not wait for the observer; there is no interposition on the hot path, no code changes in the workload, no instrumentation for the guest to fight with. The microVM itself cold-starts in milliseconds (we have covered that ground before), and the introspection tax on top is small enough that we simply run it on every task, by default, in production. Full measurements are part of the write-up above; the short version is that if you can afford the sandbox, you can afford to watch it.
Audit first, enforcement second
We are deliberately leading with observability here, not enforcement, and that is not a dodge. It is an ordering.
Enforcement decides what an agent may do, and we have written about where those controls have to live, from the isolation model down to the runtime that enforces it . But enforcement is only as good as the ground truth it acts on, and after the fact, the audit trail is your ground truth. When the auditor, the incident review, or your own postmortem asks what the agent did, "here is what it told us" is not an answer. "Here is what crossed the hardware boundary, lifted into operations, cross-checked against what it told us" is.
Essentially, we want the postmortem of an agent session to be held to the same standard as the postmortem of an outage: evidence first, narrative second.
There is more to say about every piece of this (the taxonomy, the lifting accuracy, what happens when the layers disagree in the wild), and we will say it in the coming weeks. If you want the full picture now (the threat model, the isolation scorecard, and the evidence behind each claim), it is all in the whitepaper below.
Quis custodiet ipsos custodes? The boundary does. That is rather the point of having one.
Seven claims. Check them on your stack.
The full white paper has the threat model, the isolation scorecard, and the evidence behind each claim. Or book a call with a founder to see the boundary live.
⌐ Continue reading The Sandboxing Manifesto Engineering · 5 min → When the bug is real but the path isn't: ITScape (CVE-2026-46316) and the case for minimal runtimes Engineering · 5 min → VM isolation at container speed: urunc and the Kubernetes agent-sandbox CRD Engineering · 5 min → The Context & Control Model for Production.
Product Overview Service Catalog Incidents Always-On Workflows Data Governance Prevent AI Agent Governance Integrations Customers Company About Security Blog Contact us Resources Library Events & podcasts AI SRE benchmark AI reliability guide Glossary Documentation ↗ © 2026 NOFire AI Privacy Policy Terms of Service v.2026.05 Book a demo Book a demo
See it on your stack. NOFire keeps production up.
A 30-minute walkthrough with a founder. We map your stack to the Context & Control Model, live.
