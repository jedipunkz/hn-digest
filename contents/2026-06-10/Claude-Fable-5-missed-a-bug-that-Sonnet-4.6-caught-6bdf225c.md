---
source: "https://alikhallad.com/i-tested-claude-fable-5-against-private-benchmarks-it-has-never-seen-here-is-where-it-failed/"
hn_url: "https://news.ycombinator.com/item?id=48480057"
title: "Claude Fable 5 missed a bug that Sonnet 4.6 caught"
article_title: "I Tested Claude Fable 5 Against Private Benchmarks It Has Never Seen, Here Is Where It Failed – Ali Khallad"
author: "startages"
captured_at: "2026-06-10T19:00:17Z"
capture_tool: "hn-digest"
hn_id: 48480057
score: 3
comments: 0
posted_at: "2026-06-10T17:51:41Z"
tags:
  - hacker-news
  - translated
---

# Claude Fable 5 missed a bug that Sonnet 4.6 caught

- HN: [48480057](https://news.ycombinator.com/item?id=48480057)
- Source: [alikhallad.com](https://alikhallad.com/i-tested-claude-fable-5-against-private-benchmarks-it-has-never-seen-here-is-where-it-failed/)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T17:51:41Z

## Translation

タイトル: Claude Fable 5 は Sonnet 4.6 が発見したバグを見逃しました
記事のタイトル: これまで見たことのないプライベート ベンチマークに対してクロード フェイブル 5 をテストしましたが、失敗したのはここです – Ali Khallad

記事本文:
サービス メニューを開く
プラグイン開発
これまでに見たことのないプライベート ベンチマークに対してクロード フェイブル 5 をテストしましたが、失敗したのはここです
今週 Anthropic が Claude Fable 5 をリリースすると、私のフィードは数時間以内に同じベンチマーク チャートでいっぱいになりました。 SWE ベンチ スコア、エージェント コーディング番号、Stripe 移行のストーリー。すべて印象的な数字であり、すべての数字はベンダー自身によって公開されています。これらのチャートでは分からないことを知りたかったのです。このモデルは、これまでに見たことのないバグや、読み取れないテストによって評価されたバグに対してどのように動作するのでしょうか?
私はまさにこの目的のために、小規模なプライベート ベンチマーク スイートを維持しています。これらのインスタンスは、私が取り組んできたプロジェクトの実世界のバグから来ており、自己完結型の壊れたアプリケーションとして再構築されたもので、どこにも公開されていません。公開ベンチマークがトレーニング データに漏洩し、それが発生すると、デバッグ能力を測定するのではなく、リコールを測定することになります。私の場合は非公開のままなので、この投稿ではタスクについて説明しますが、コードは決して示しません。
結果の前に、設定の概要を簡単に説明します。評価の規律が数値を読む価値のあるものであるためです。
各ベンチマークは、壊れたプロジェクトのクリーンなサンドボックス コピーをステージングします。エージェントは、QA が作成するような方法で作成されたバグ レポートを受け取ります。症状や制約はなく、ファイル名や原因に関するヒントは含まれていません。
実際の受け入れテストはサンドボックスの外にあります。これらは、エージェントの終了後にのみ、ネットワーク アクセスが無効になって実行されている Docker コンテナ内にコピーされます。エージェントがそれらを見ることはありません。
評価は行動に基づいて行われます。テストは実際のブラウザを駆動し、コードがどのように記述されたかではなく、必要に応じてフレームごとにレンダリング内容をアサートします。私が考えていたものだけでなく、正しい修正はすべて合格します。
構成ファイル、マニフェスト、およびロックファイルは、エージェントが実行される前にハッシュされます。どちらかといえば良い

変更が影響した場合、実行は無効になります。
実行は、すべての非表示ベリファイアが合格し、すべての既存の動作ガードが緑色のままである場合にのみ、解決済みとしてカウントされます。部分クレジットは存在しません。
各インスタンスは他のフロンティア エージェントでも事前に調整されていたため、基準点がありました。Claude Opus 4.7、Claude Sonnet 4.6、および Codex を介した GPT-5.5 はそれぞれ、同一の条件下でこれらのタスクの一部をすでに試みていました。
4 つのインスタンス (すべて React と TypeScript、中程度から極端な範囲)。スイートはプライベートであるため、説明は意図的に曖昧にしておきます。
クライアント側のナビゲーション中に内部ページのクロムが一時的にリセットされるダッシュボード SPA (中)。
Webhook コンソールでは、配信ログは成功を主張しますが、受信エンドポイントは同意しません (中)。
Webhook コンソールでは、個別のイベントが短期間バーストされると、それらの一部が静かに失われます (ハード)。
列をまたがるドラッグ アンド ドロップ (極端な) 中に 2 つの別々の視覚的不具合が発生するカンバン ボード。
クロード コードを介してクロード ファブル 5 を 4 つすべてに対して実行しました。それぞれ 1 回ずつ試行し、再試行はしませんでした。
4 つのうち 3 つには、GPT-5.5 がアプリ シェル全体を書き換え、その過程で 12 の動作ガードを破って失敗したというハード インスタンスが含まれます。 Fable は、約 40 行の差分で同じ問題を約 4 分の 1 の時間で修正しました。
スイート内のすべてのベンチマークには、インスタンスを構築したときに自分で書いた修正パッチであるリファレンス パッチが付属しています。これが私が最も驚いたパターンです。3 つのパスすべてで、Fable は私の参照パッチとは異なるシステムのレイヤーを修復しました。私の解決策に収束したことは一度もありませんでした。
ダッシュボードタスクではさらに一歩進んだ。私のリファレンス修正では、遷移中に以前にロードされたデータが無条件で画面上に保持されます。 Fable は、古いデータがエンティティ スイッチ全体に引き継がれないようにガードを追加しました。

h は、私自身のパッチが技術的に許可している短い間違ったデータのフラッシュを閉じます。私の隠しテストではそのフラッシュを検出することさえできません。このモデルは、私の回答キーが無視するエッジケースから防御しました。
バーストロスの課題は私が常に考えている課題です。ベンチマークを構築するときは、意図した修正の一部ではないバグの近くに、囮と思われるメカニズムを仕掛けます。これは、特に動作を追跡するのではなく、パターン一致するエージェントを捕捉するためです。このインスタンスに関する私の管理者のメモでは、そのような 2 つのメカニズムをおとりとして明示的にマークしています。 Fable はそれらの両方を取り上げ、まさにその 2 つの部分から実用的な修正を作成しました。その後、穴が見つかることを期待しながら、メカニズムを 1 行ずつトレースしました。穴はありません。すべての失われたイベントは、実際の配信経路を通って宛先に確実に届きます。トラップは実際の機械から作られており、実際の機械は耐荷重性があります。私が作成する次のインスタンスのために学んだ教訓。
カンバン インスタンスは、1 つのバグ レポートで 2 つの視覚的な問題を説明します。ドロップされたカードは、落ち着く前にフレームのソース列にフラッシュバックする可能性があり、空の列にドロップされたカードは、列ヘッダーを目に見えて通過する可能性があります。 2 つの症状、まったく無関係な 2 つの根本原因。 1 つはデータのタイミングの問題です。もう 1 つは純粋なレイアウト ジオメトリです。どちらかを修正しても、もう一方には何も影響しません。
Sonnet 4.6 はキャリブレーション中にジオメトリのバグを修正し、タイミングのバグを見逃しました。失敗した。
Fable 5 では、真に洗練された同期レンダリング アプローチを使用して、2 つのうちのはるかに難しいタイミング バグを修正しました。その後、ジオメトリのバグを完全に見逃していました。こちらも失敗。前世代のモデルは、新しいフラッグシップが見落としていた正確なバグを捕らえ、新しいフラッグシップは、旧モデルが解決できなかった正確なバグを解決しました。同じレポートの反対側の半分。
ファブルはなぜそうなったのか

古いモデルが何か見つけましたか？独自の終了概要がそれに答えています。タイミングの問題を解決した後、2 番目の症状は最初の症状の副作用であると説明しました。その理論では、プレースホルダー テキストを一時的に表示する空の列は「ヘッダーが消える」レポートでした。根本原因が 1 つあり、すべての症状が説明され、調査は終了しました。この理論はエレガントでしたが、半分は間違っていました。最初の原因ですべてが説明されたため、2 番目の原因を探すことはありませんでした。
その要約のもう 1 行は、これらのエージェントがどのように成功を宣言するかについて何かを述べているため、そのまま引用する価値があります。「影響を受けるコード パスを通じてレンダリング/コミットのタイミングをトレースすることによって検証される」ということです。修正を実行するのではなく、修正について検討しました。サンドボックスでは依存関係のインストールが許可されていないため、アプリを実行できませんでしたが、そのことを公表したのは名誉なことです。しかし、完全に自信を持って提示された半分検証済みの理論は、まさに査読者が注意すべき失敗モードです。
注目に値するのは、失敗しても何も壊れていないということです。そのインスタンスの 10 個の動作ガードはすべて緑色のままでした。失敗は診断が不完全であり、巻き添え被害ではありませんでした。
今週の Fable の目に見えない安全対策と、モデルが伝えることを信頼できるかどうかについての議論を踏まえて、私は逆の方向、つまりモデルが提供するものを信頼できるかどうかを確認しました。実行後は毎回、通常のトリック、環境スニッフィング、テストのみのブランチ、ハードコードされた期待値、偽の出力、脆弱なガードがないかどうかの差分を監査します。また、配送に隣接するコードが変更された 2 つの実行については、特に「領収書を偽造する」ショートカットが存在する場所であるため、詳細に監査しました。
4回のランはすべてクリーンに戻った。保護されたファイルの改ざん、テスト検出、ショートカットはありません。失敗した修正も含め、すべての修正は本物のエンジニアリングでした。
4

インスタンスはケーススタディであり、リーダーボードではありません。私は「Fable はバグの 75% を解決する」と主張しているわけではありません。
4 つのタスクはすべて React と TypeScript です。単一のスタック ファミリだけでは、Python、Go、PHP については何もわかりません。
タスクごとに 1 回の試行。エージェントの実行にはばらつきがあり、2 回目の試行では結果がどちらかの方向に反転する可能性があります。
トークンとコストのテレメトリは CLI 印刷モードでは存続できなかったので、タイミングはありますが、消費することはできません。
Claude Fable 5 は、私がこのスイートで実行した中で最も強力なデバッグ エージェントであり、制約の下でギャップが最も顕著に現れています。最小の差分、最速のパス、4 回の実行すべてで壊れたガードはゼロ、そして 1 つのタスクでは、私自身の解答キーよりも慎重に修正されました。ベンダーのチャートは能力に関して嘘をついていません。
しかし、失敗はそれをどのように対処するかを教えてくれます。バグが難しすぎたためにモデルが失敗したわけではありません。それが失敗したのは、あらゆる症状を説明する 1 つの美しい理論を見つけて停止したからです。エージェントが作成した修正をレビューする場合は、速度を緩める瞬間です。単一の根本原因がレポート内のすべてを説明している場合は、症状が実際に独立している場合はどうなるかを考えてください。古くて弱いモデルは、まさに証拠を統合しなかったため、同じタスクで部分的に成功しました。
モデルのバグ修正はますます向上しています。独立した原因を列挙し、1 つの洗練された説明の誘惑に抵抗し、それについて推論するのではなく実際に検証を実行することは、依然として人間が関与する必要があるところです。今のところ。
Plesk で WordPress サーバー エラーを推測せずにデバッグする方法

## Original Extract

Services Open menu
Plugin development
I Tested Claude Fable 5 Against Private Benchmarks It Has Never Seen, Here Is Where It Failed
When Anthropic released Claude Fable 5 this week, my feed filled up with the same benchmark charts within hours. SWE-bench scores, agentic coding numbers, the Stripe migration story. All impressive, and all numbers published by the vendor itself. I wanted to know something those charts cannot tell me: how does this model behave on bugs it has never seen, graded by tests it cannot read?
I maintain a small private benchmark suite for exactly this purpose. The instances come from real-world bugs in projects I have worked on, rebuilt as self-contained broken applications, and they have never been published anywhere. Public benchmarks leak into training data, and once that happens you are no longer measuring debugging ability, you are measuring recall. Mine stay private, which is also why this post describes the tasks but never shows their code.
Before the results, a quick summary of the setup, because the grading discipline is what makes the numbers worth reading.
Each benchmark stages a clean sandbox copy of a broken project. The agent gets a bug report written the way QA would write it – symptoms and constraints, no file names, no hints about the cause.
The real acceptance tests live outside the sandbox. They are copied in only after the agent finishes, inside a Docker container running with network access disabled. The agent never sees them.
Grading is behavioral. The tests drive a real browser and assert on what renders, frame by frame where needed, not on how the code was written. Any correct fix passes, not just the one I had in mind.
Config files, manifests, and lockfiles are hashed before the agent runs. If anything protected changes, the run is voided.
A run only counts as resolved when every hidden verifier passes and every existing behavior guard stays green. Partial credit does not exist.
Each instance was also calibrated beforehand with other frontier agents, so I had reference points: Claude Opus 4.7, Claude Sonnet 4.6, and GPT-5.5 through Codex had each already attempted some of these tasks under identical conditions.
Four instances, all React and TypeScript, ranging from medium to extreme. I am keeping descriptions deliberately vague since the suite is private.
A dashboard SPA where the inner page chrome briefly resets during client-side navigation (medium).
A webhook console where the delivery log claims success while the receiving endpoint disagrees (medium).
A webhook console where short bursts of distinct events silently lose some of them (hard).
A kanban board with two separate visual glitches during cross-column drag and drop (extreme).
I ran Claude Fable 5 through Claude Code against all four, one attempt each, no retries.
Three out of four, including the hard instance that GPT-5.5 had failed by rewriting an entire app shell and breaking a dozen behavior guards in the process. Fable fixed the same problem with a diff of around forty lines, in roughly a quarter of the time.
Every benchmark in the suite ships with a reference patch, the fix I wrote myself when I built the instance. Here is the pattern that surprised me most: on all three passes, Fable repaired a different layer of the system than my reference patch did. Not once did it converge on my solution.
On the dashboard task it went a step further. My reference fix keeps previously loaded data on screen during transitions, unconditionally. Fable added a guard so stale data can never carry over across an entity switch, which closes a brief wrong-data flash that my own patch technically allows. My hidden tests cannot even detect that flash. The model defended against an edge case my answer key ignores.
The burst-loss task is the one I keep thinking about. When I build a benchmark I plant decoys, plausible-looking mechanisms near the bug that are not part of the intended fix, specifically to catch agents that pattern-match instead of tracing behavior. My maintainer notes for this instance explicitly mark two such mechanisms as decoys. Fable picked up both of them and composed a working fix out of exactly those two pieces. I traced the mechanics line by line afterwards, fully expecting to find a hole. There is no hole. Every lost event genuinely reaches the destination through the real delivery path. The trap was built from real machinery, and real machinery can be load-bearing. Lesson learned for the next instance I author.
The kanban instance describes two visual problems in one bug report. A dropped card can flash back into its source column for a frame before settling, and a card dropped into an empty column can visibly pass through the column header. Two symptoms, two completely unrelated root causes. One is a data-timing problem. The other is pure layout geometry. Fixing either one does nothing for the other.
During calibration, Sonnet 4.6 fixed the geometry bug and missed the timing bug. Failed.
Fable 5 fixed the timing bug, which is by far the harder of the two, with a genuinely sophisticated synchronous rendering approach. Then it missed the geometry bug entirely. Also failed. The previous-generation model caught the exact bug the new flagship overlooked, and the new flagship cracked the exact bug the older model could not. Opposite halves of the same report.
Why did Fable miss something an older model found? Its own exit summary answers that. After solving the timing problem, it explained the second symptom as a side effect of the first: in its theory, the empty column briefly showing its placeholder text was the “header disappears” report. One root cause, every symptom accounted for, investigation closed. The theory was elegant and half wrong. It never went looking for a second cause because the first one explained everything.
One more line from its summary is worth quoting verbatim, because it tells you something about how these agents declare success: “verified by tracing the render/commit timing through the affected code paths.” It reasoned about the fix instead of executing it. The sandbox does not allow dependency installation, so it could not run the app, and to its credit it disclosed that. But a half-verified theory presented with full confidence is exactly the failure mode a reviewer needs to watch for.
Worth noting: even while failing, it broke nothing. All ten behavior guards on that instance stayed green. The failure was incomplete diagnosis, not collateral damage.
Given this week’s discussion about Fable’s invisible safeguards and whether you can trust what the model tells you, I checked the opposite direction: whether you can trust what it ships. After every run I audit the diff for the usual tricks, environment sniffing, test-only branches, hardcoded expected values, faked outputs, weakened guards. I also specifically deep-audited the two runs where it modified delivery-adjacent code, since that is exactly where a “fake the receipts” shortcut would live.
All four runs came back clean. No protected-file tampering, no test detection, no shortcuts. Every fix was real engineering, including the failed one.
Four instances is a case study, not a leaderboard. I am not claiming “Fable resolves 75% of bugs.”
All four tasks are React and TypeScript. A single stack family tells you nothing about Python or Go or PHP.
One attempt per task. Agent runs have variance, and a second attempt could flip results in either direction.
Token and cost telemetry did not survive the CLI print mode, so I have timing but not spend.
Claude Fable 5 is the strongest debugging agent I have run through this suite, and the gap shows most under constraint: smallest diffs, fastest passes, zero broken guards across all four runs, and on one task a fix more careful than my own answer key. The vendor charts are not lying about capability.
But the failure tells you how to work with it. The model did not fail because the bug was too hard. It failed because it found one beautiful theory that explained every symptom, and stopped. If you review agent-written fixes, that is the moment to slow down: when a single root cause accounts for everything in the report, ask what happens if the symptoms are actually independent. An older, weaker model partially succeeded on the same task precisely because it never unified the evidence.
Models are getting better at fixing bugs. Enumerating independent causes, resisting the pull of one elegant explanation, and actually running the verification instead of reasoning about it – that is still where they need a human in the loop. For now.
How to Debug WordPress Server Errors on Plesk Without Guessing
