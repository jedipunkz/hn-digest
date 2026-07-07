---
source: "https://runme.dev/blog/runme-eval"
hn_url: "https://news.ycombinator.com/item?id=48820522"
title: "Sure, Your Claude Is Doing Things. Prove It"
article_title: "Sure, Your Claude Is Doing Amazing Things. Prove It. • RUNME"
author: "sourishkrout"
captured_at: "2026-07-07T17:06:23Z"
capture_tool: "hn-digest"
hn_id: 48820522
score: 1
comments: 0
posted_at: "2026-07-07T17:01:27Z"
tags:
  - hacker-news
  - translated
---

# Sure, Your Claude Is Doing Things. Prove It

- HN: [48820522](https://news.ycombinator.com/item?id=48820522)
- Source: [runme.dev](https://runme.dev/blog/runme-eval)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T17:01:27Z

## Translation

タイトル: 確かに、あなたのクロードは物事をやっています。証明してみよう
記事のタイトル: 確かに、あなたのクロードは素晴らしいことをしています。それを証明してください。 •ランミー
説明: クロードのことを自慢するのは簡単です。エージェントに導入した動作が引き続き機能することを証明するのはさらに困難です。 Runme v3.17 では、ローカル タスクの評価がリポジトリに取り込まれるため、AI ワークフローを再実行、比較、証拠とともに促進できます。

記事本文:
確かに、あなたのクロードは素晴らしいことをしています。それを証明してください。 • RUNME RUNME Stars Docs コミュニティ ブログ
ai セバスチャン・ハックルベリー 確かに、あなたのクロードは素晴らしいことをしていますね。それを証明してください。
あなたのスキル (そして私のスキルも) をクソにしてください [1]。それが最も真実に近いセリフだった。 「スキルは役に立たない」ではありません。 「エージェントは偽物」ではありません。さらに不快なこと: 再利用可能なエージェントの動作として、指示、スクリプト、参照、ツールの期待値、およびワークフローの習慣のフォルダーを渡す場合、エージェントのスキルと作成者のスキルの両方をコードのように検査する必要があります。
プロンプトとしてではありません。魔法ほどではありません。コード。
スモーク テスト、統合テスト、単体テストなど、コードのようにテストしてみませんか?
なぜなら、エージェント スキルは便利に見えても、他のエージェント ワークフローと同様に失敗する可能性があるからです。たとえば、間違ったアクティベーション、間違ったツール、古いコンテキスト、欠落したソース、漏洩した境界、またはエージェントが一度運よく成功しただけで成功した洗練されたアーティファクトなどです。
「私も」という部分が重要です。これは、これらのワークフローを作成している人々に冗談を戻すことになります。スキルがレジストリ、チームメイト、またはユーザーに何かを証明する前に、スキルは作成者に何かを証明する必要があります。スキルは、エージェントが消費する単なるマークダウンではありません。それは、仮定、来歴、保守習慣、判断によって構成される小さなソフトウェア オブジェクトです。それをインストールすると、作者の操作上の伝承を継承することができます。
それが runme eval ( v3.17 ) が退屈なものにしようとしている問題です。繰り返しの雰囲気チェックを、民間伝承が去った後もワークフローが機能し続けることの証明に置き換えます。スキルは氷山の一角にすぎません。同じタスク評価パターンは、命令、ツール、コンテキスト、メモリ、および最終成果物が連携して動作し続ける必要があるあらゆる場所に適用されます。 runme eval は、Harbor の eval モデルを使用して、ワークフローがすでに存在するローカル リポジトリにその規律をもたらします。

[2]の下にあります。
これがスキルを超えて広がるのは、「あなたのクロード」が単なる一枚岩のモデルではなくなったからです。またはあなたのコーデックス。またはChatGPT。またはカーソル。またはオープンコード。これらのツールがワークフローの一部になると、重要になるのは、基本モデルと作業自体の間の導入層です。
repo instructions such as CLAUDE.md or AGENTS.md
repo-local and user-wide skills
ツール、検索システム、および MCP サーバー構成
secrets and redaction boundaries
the task trajectory the agent chooses
the final artifact it produces
この投稿は、LLM またはエージェント SDK で構築されたオーダーメイド エージェントの評価に関するものではなく、サプライ チェーン攻撃、悪意のあるスキル、機密漏洩などのセキュリティ問題に関するものでもありません。 Those are separate problems.これは、実際のリポジトリで人々がすでに使用している人気のある AI 活用に関するものです。スキル、命令、ツール、ワークフローの習慣を Claude、Codex、Cursor、ChatGPT、OpenCode などのエージェントにデプロイし、その動作が引き続き機能することを証明します。
私たちはそのレイヤーのようなインフラストラクチャに依存していますが、スクリーンショット、デモ、ラッキーラン、個人的な儀式など、コピー＆ペーストのプロンプトのようにそれを検証します。 Hamel Husain のスキル懐疑論に関する Show Us Your Agent Skills セグメントでも同様の点が指摘されています。公開スキルはコードのように読み取られ、再利用可能なインフラストラクチャとして扱われる前に来歴、保守、制約がチェックされるべきです [3]。同じ規律がワークフロー全体に適用されます。つまり、スモーク テスト、部分の集中テスト、ツールとコンテキストの統合テスト、以前は機能していた軌道の回帰履歴です。
runme eval は内部ループから始まります。タスクを記録し、ワークフローを実行し、証拠を保管し、Claude がすでに構成されて作業を行っているリポジトリの改善履歴の一部を作成します。 That lowers the barrier for p

外側のループへのパスを開いたままにし、CI またはスケールアウトされたベンチマーク ジョブで同じタスクの評価を実行します。これは、Harbor [2] や Terminal-Bench 2.0 [4]、SWE-Bench Verified などのベンチマークの背後にある同じ洞察に基づいて構築されています。エージェントの作業をタスクとしてキャプチャし、ハーネスを介して実行し、雰囲気ではなく再現可能な証拠によって判断できます。そのため、Harbor は本格的なベンチマーク インフラストラクチャとリポジトリ ローカル ワークフロー回帰テストの間の便利な橋渡しとなります。
実際の障害モードを含む小さなワークフローを使用して、これを具体的にしてみましょう。
スキルは検査するには十分な範囲があるものの、大規模なワークフローのように失敗する可能性があるほど乱雑であるため、スキルは依然として有用な出発点です。 Skills.sh にはすでに 100 万近くの公開スキルがリストされていますが、その多くは単一コミットのアーティファクトのように見えます。つまり、一度生成され、めったにインストールされず、積極的にメンテナンスされていません。私もスキルクリエイターのダンピングをしてきました。
ソフトウェア テストと同様に、作成者はテスト対象の境界を選択します。その境界は、1 つのスキル、エージェントのワークフロー全体、ツールの統合、またはチームが依存する最終成果物である可能性があります。 runme eval のデフォルトは ./evals/tasks です。これは、一般的なケースがスタンドアロンのスキル アーティファクトだけでなく、リポジトリに属する​​タスクまたはワークフロー統合であると想定しているためです。スキル固有の eval ディレクトリを指定することもできます。境界は、結果を再実行、比較、改善できるように十分に明確である必要があります。
1 つのタスクの評価だけでは、ワークフロー全体を証明するには十分ではありません。退行圧力を強め始めるだけで十分です。このようにして、私たちはすでに小規模で焦点を絞ったスモーク テストと統合テストを、徹底的な証明としてではなく、重要な障害モードの永続的なチェックとして使用しています。 Guillermo Rauch の古いテストに関するアドバイスは今でも当てはまります: テストを作成し、あまりにもテストをしすぎないでください

ニューヨーク州、主に統合[5]。リポジトリローカルの評価スイートは、実用化する前に何百、何千ものタスクを必要としません。実際の障害モードをキャプチャするタスクから始めて、ワークフローが中断、ドリフトする、または昇格の決定が必要なケースを追加します。
評価に値する小規模なワールドカップのワークフロー
ワールドカップのピックレポートスキル [6] は、そのための優れたショーケースとなります。それは、2026 年ワールドカップという実際の主題に加えて、スキルが有効化され、Web を検索し、その主張に近い適切な情報源を使用し、ターゲットのスレートとラインナップのステータスを網羅し、ワークフローを順番に守り、不確実性のガードレールを尊重し、厳密な形式で最終的な成果物を作成したという証拠に依存します。検査するには十分小さいですが、大規模なエージェント ワークフローと同じ可動部分を実行します。
実際の様子は次のとおりです。私の OpenClaw、「Sawyer」は、スキルを使用してスコアライン ピックを生成しています。
原則として、この評価は OpenClaw を直接ターゲットにすることができます。この実行では、Harbor の ATIF 変換のバグにより、OpenClaw の完全な会話が 1 つのステップにまとめられ、ハーネス間での軌道の移植性が低下します。私の OpenClaw はすでに内部で Codex を使用しているため、ショーケースでは代わりに Codex が実行されます。
重要な部分は、ドメイン固有のルーブリックです。フレームワークはタスクを実行し、証拠を保存し、結果を比較できますが、作成者はワークフローにとって「良い」とは何か、つまりどのソースが重要か、どの仮定が許容されるか、どの省略が重要か、有用なレポートに何を含める必要があるかを定義する必要があります。この例では、その判断はワールドカップレポートタスクの小さな採点ルーブリックに反映されています[7]。
報酬は最終成果物で終わってはいけません。出力のみをスコアリングすると、評価が測定する予定の作業が失われます。また、エージェントの軌跡、つまりエージェントがアクティブ化したかどうかもスコアリングする必要があります。

d 適切なツールと呼ばれる意図されたスキルは、期待されたシーケンスに従い、不確実性を処理し、報酬を説明するスコアリング出力を生成しました [8]。 ATIF (Agent Trajectory Interchange Format) は、すべてのハーネスに独自のプライベート ログ形式を残すのではなく、エージェント全体でこれらのステップを共有する形式を提供するため、ここで重要になります [9]。
それがRunmeとHarborを併用するポイントです。評価インフラストラクチャの変動部分は退屈になるはずなので、作成者は、テスト対象のユニット、ルーブリック、報酬、スコアなど、重要な部分に判断を費やすことができます。これは、新しいエージェント スタックを最初から構築するのではなく、再利用可能な動作を既存のエージェントにデプロイする人にとっては特に重要です。ベスト プラクティスを適用するには用語の共有が必要です。共有ハーネスにより、何かを測定する前にすべてのリポジトリに独自の用語集を作成するよう強制するのではなく、タスク、トライアル、仕事、軌跡、昇進に関する共通の語彙がチームに提供されます。
言い換えれば、これはモデルがもっともらしいスポーツの散文を書けるかどうかをテストしているわけではありません。これは、ワークフローが指示、ツール、コンテキスト、軌道、モデルの選択、作業量の設定、成果物全体にわたって依然としてまとまっているかどうかをテストしています。
$ runme 評価ビュー
ハーバービューアーを起動する
ジョブフォルダー: ~/sourishkrout-skills/.runme/evals/jobs
モード: ジョブ
サーバー: http://127.0.0.1:8080
これにより、ローカルの評価履歴が自動的に開きます。ショーケース リポジトリ [6] では、プロモートされたジョブも公開されている [10] ため、レビュー担当者は現在の実行を信頼する前に以前の実行を検査できます。現在、runme eval ビューは Harbor View へのショートカットです。これは機能しますが、レビュー担当者が eval モデルを理解していることが前提となります。時間が経つにつれて、Runme UX はレビュー担当者により明確な概要を提供し、レビューとプロモーションの決定をより迅速に行えるようになります。
リポジトリのルートから回帰を「強制」しましょう

少ない労力で:
$ runme eval skill/world-cup-picks-report/evals/regression \
--エージェント コーデックス \
--ak 推論_努力=低
回帰 • runme-codex
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━┓
┃ ルーブリックディメンション ┃ スコア ┃
┡━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━┩
│ アーティファクト_書き込み │ 1.000 │
│ 引用_近接性 │ 1.000 │
│ ガードレール │ 1.000 │
│ ラインナップ_ステータス │ 0.000 │
│ レポート_完全性 │ 1.000 │
│ スキル発動 │ 1.000 │
│ ソース品質 │ 0.750 │
│ ターゲット_スレート │ 0.867 │
│ ワークフロー_シーケンス │ 0.750 │
│ 総合報酬 │ 0.864 │
━━━━━━━━┴───────┘
求人情報
合計実行時間: 2 分 23 秒
結果は .runme/evals/jobs/2026-07-05__11-44-20/result.json に書き込まれます。
「harbor view .runme/evals/jobs」を実行して結果を検査します。
「harbor Upload .runme/evals/jobs/2026-07-05__11-44-20」を実行して結果を共有します。
これは、適用される作業単位です。パッケージ化されたエンドツーエンド タスク [11] に対して Codex を実行し、アーティファクトを収集し、スキルが実際にアクティブ化されたかどうかを判断するために必要な軌跡の証拠 [12] を保持し、期待されるツールを呼び出し、ワークフローに従い、適切なソースを使用し、ターゲット スレートをカバーし、ガードレールを尊重し、賭けのヒントの散文に陥ることなく期待されるレポートを作成します。
ローカル評価を実行すると、作業ツリーを変更できます。これは意図的なものです。タスクがエージェントにファイルの編集、アーティファクトの生成、または状態の更新を要求する場合、評価では、エージェント ハーネスを自分で操作することで得られるのと同じリポジトリの変更が確認される必要があります。
通常の使用では、モデルとエフォートのノブをスキップして、選択したエージェントを直接操作した場合と同様に動作させます。ドゥ

リング反復開発では、より低い推論労力のようなオーバーライドは、弱い構成の調査、障害の再現、または意図的な回帰の強制に依然として役立ちます。
$ runme eval スキル比較/world-cup-picks-report/evals/regression
ベース: HEAD で追跡される .runme/evals/jobs/2026-06-30__11-40-20
最新: .runme/evals/jobs/2026-07-05__11-44-20 local
データセット: スキル/ワールドカップピック-レポート/評価/回帰
エージェント: runme-codex
モデル: gpt-5.5
環境: runme_harbor.environment:RunmeEnvironment
仕事:
完了: 1 -> 1 +0
エラー: 0 -> 0 +0
evals: 1 -> 1 +0
結果:
回帰: 報酬 0.979 -> 0.864 -0.115
推奨事項: 候補者は後退しました。昇格する前にジョブ/タスクの詳細を再実行または検査します。
この場合、最新の実行は開発中に意図的に弱い構成に強制されているため、比較によって低下が検出されるはずです。ローカル ジョブを Git で追跡されたベースラインと比較し、弱い結果を別の逸話的な記録ではなく明示的な回帰に変換します。
結果を改善するには、スキルを編集するか、 CLAUDE.md や AGENTS.md などのコンテキスト ファイルを調整するか、MCP サーバー構成を変更するか、 --env docker とタスクローカル Dockerfile を使用してサンドボックス環境で評価を再実行してから、何を昇格させるかを決定します。
結果を宣伝する前に、証拠をプレビューすると、ゲート Runme が適用されます。
$ runme eval promote --latest --artifacts --dry-run skill/world-cup-picks-report/evals/regression
選択した評価ジョブ: .runme/evals/jobs/

[切り捨てられた]

## Original Extract

Bragging about your Claude is easy. Proving the behavior you deploy into agents keeps working is harder. Runme v3.17 brings local task evals into the repo, so AI workflows can be rerun, compared, and promoted with evidence.

Sure, Your Claude Is Doing Amazing Things. Prove It. • RUNME RUNME Stars Docs Community Blog
ai Sebastian Huckleberry Sure, Your Claude Is Doing Amazing Things. Prove It.
F*** Your Skills (And Mine Too) [1]. That was the line that got closest to the truth. Not "skills are useless." Not "agents are fake." Something more uncomfortable: if we are going to pass around folders of instructions, scripts, references, tool expectations, and workflow habits as reusable agent behavior, then both the agent's skill and the author's skill need to be inspected like code.
Not as prompts. Not as magic. Code.
Why don't we test them like code: smoke tests, integration tests, unit tests?
Because an agent skill can look useful and still fail like any agent workflow: wrong activation, wrong tool, stale context, missing sources, leaky boundaries, or a polished artifact that only passed because the agent got lucky once.
The "mine, too" part matters. It turns the joke back on the people writing these workflows. Before a skill proves anything to a registry, a teammate, or a user, it has to prove something to its author. A skill is not just markdown an agent consumes; it is a little software object made of assumptions, provenance, maintenance habits, and judgment. Install it, and you inherit the author's operational folklore.
That is the problem runme eval ( v3.17 ) is trying to make boring: replace repeated vibe checks with proof that the workflow continues to work after the folklore leaves home. Skills are just the tip of the iceberg. The same task eval pattern applies anywhere instructions, tools, context, memory, and final artifacts have to keep working together. runme eval brings that discipline into the local repo, where the workflow already lives, using Harbor's eval model underneath [2].
The reason this extends beyond skills is that "your Claude" is no longer just a monolithic model. Or your Codex. Or ChatGPT. Or Cursor. Or OpenCode. Once these tools become part of your workflow, the thing that matters is the adoption layer between the base model and the work itself:
repo instructions such as CLAUDE.md or AGENTS.md
repo-local and user-wide skills
tools, retrieval systems, and MCP server configuration
secrets and redaction boundaries
the task trajectory the agent chooses
the final artifact it produces
This post is not about evaluating bespoke agents built with LLM or agent SDKs, nor about security problems such as supply-chain attacks, malicious skills, or secret exfiltration. Those are separate problems. This is about the popular AI harnesses people already use in real repos: deploying skills, instructions, tools, and workflow habits into agents such as Claude, Codex, Cursor, ChatGPT, or OpenCode, then proving that behavior continues to work.
We depend on that layer like infrastructure, but still validate it like copy-paste prompts: screenshots, demos, lucky runs, and personal rituals. The Show Us Your Agent Skills segment on Hamel Husain's skill scepticism makes a similar point: public skills should be read like code, with provenance, maintenance, and constraints checked before anyone treats them as reusable infrastructure [3]. The same discipline applies to the full workflow: smoke tests, focused tests for the pieces, integration tests for tools and context, and regression history for trajectories that used to work.
runme eval starts with the inner loop: record the task, run the workflow, keep the evidence, and make improvement history part of the repo where Claude is already configured and doing the work. That lowers the barrier for people getting started, while keeping the path open to the outer loop: running the same task evals in CI or scaled-out benchmark jobs. It builds on the same insight behind Harbor [2] and benchmarks like Terminal-Bench 2.0 [4], SWE-Bench Verified, and others: agent work can be captured as tasks, run through a harness, and judged by repeatable evidence instead of vibes. That makes Harbor a useful bridge between serious benchmark infrastructure and repo-local workflow regression tests.
Let's make that concrete with a small workflow that has real-world failure modes.
Skills are still a useful starting point because they are bounded enough to inspect, but messy enough to fail like larger workflows. Skills.sh already lists close to a million published skills, and many of them look like single-commit artifacts: generated once, rarely installed, and not actively maintained. I have done my share of skill-creator dumping, too.
As with software tests, the author chooses the boundary of what is under test. That boundary might be one skill, a full agent workflow, a tool integration, or the final artifact a team depends on. runme eval defaults to ./evals/tasks because it assumes the common case is a task or workflow integration that belongs with the repo, not only a standalone skill artifact. You can still point it at a skill-specific eval directory; the boundary just needs to be explicit enough that the result can be rerun, compared, and improved.
One task eval is not enough to prove the whole workflow. It is enough to start building regression pressure. That is how we already use small, sharply focused smoke and integration tests: not as exhaustive proof, but as durable checks for the failure modes that matter. Guillermo Rauch's old testing advice still fits: write tests, not too many, mostly integration [5]. A repo-local eval suite does not need hundreds or thousands of tasks before it is useful. Start with the task that captures a real failure mode, then add cases where the workflow breaks, drifts, or needs a promotion decision.
A Small World Cup Workflow Worth Evaluating
The world-cup-picks-report skill [6] makes a good showcase for that reason. It depends on a live subject, the 2026 World Cup, plus evidence that the skill activated, searched the web, used good sources close to its claims, covered the target slate and lineup status, followed the workflow in order, respected uncertainty guardrails, and produced a final artifact with a strict format. It is small enough to inspect, but it exercises the same moving parts as larger agent workflows.
Here is what that looks like in practice: my OpenClaw, "Sawyer," using the skill to produce a scoreline pick.
In principle, this eval could target OpenClaw directly. In this run, a bug in Harbor's ATIF conversion collapses the full OpenClaw conversation into a single step, which makes the trajectory less portable across harnesses. The showcase runs Codex instead, because my OpenClaw already uses Codex under the hood.
The important part is the domain-specific rubric. The framework can run the task, preserve the evidence, and compare results, but the author still has to define what "good" means for the workflow: which sources count, which assumptions are allowed, which omissions matter, and what a useful report must contain. In this example, that judgment lives in a small scoring rubric for the World Cup report task [7].
The reward should not stop at the final artifact. Scoring only the output misses the work the eval is meant to measure. It should also score the agent trajectory: whether the agent activated the intended skill, called the right tools, followed the expected sequence, handled uncertainty, and produced scoring output that explains the reward [8]. ATIF, the Agent Trajectory Interchange Format, matters here because it gives those steps a shared shape across agents instead of leaving every harness with its own private log format [9].
That is the point of using Runme and Harbor together. The moving parts of eval infrastructure should get boring, so authors can spend their judgment where it matters: the unit under test, the rubric, the reward, and the score. That matters especially for people deploying reusable behavior into existing agents, not building new agent stacks from scratch. Applying best practices requires shared terms, and a shared harness gives teams a common vocabulary for tasks, trials, jobs, trajectories, and promotions instead of forcing every repo to invent its own glossary before it can measure anything.
In other words: this is not testing whether the model can write plausible sports prose. It is testing whether the workflow still holds together across instructions, tools, context, trajectory, model choice, effort settings, and artifact.
$ runme eval view
Starting Harbor Viewer
Jobs folder: ~/sourishkrout-skills/.runme/evals/jobs
Mode: jobs
Server: http://127.0.0.1:8080
This auto-opens the local eval history. In the showcase repo [6], promoted jobs are also published publicly [10], so a reviewer can inspect previous runs before trusting the current one. Today, runme eval view is a shortcut into Harbor View: it works, but it still assumes the reviewer understands the eval model. Over time, the Runme UX should give reviewers a clearer overview, so they can make review and promotion decisions faster.
From the repo root, let's "force" a regression using low effort:
$ runme eval skills/world-cup-picks-report/evals/regression \
--agent codex \
--ak reasoning_effort=low
regression • runme-codex
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━┓
┃ Rubric Dimension ┃ Score ┃
┡━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━┩
│ Artifact_Written │ 1.000 │
│ Citation_Proximity │ 1.000 │
│ Guardrails │ 1.000 │
│ Lineup_Status │ 0.000 │
│ Report_Completeness │ 1.000 │
│ Skill_Activation │ 1.000 │
│ Source_Quality │ 0.750 │
│ Target_Slate │ 0.867 │
│ Workflow_Sequence │ 0.750 │
│ Overall Reward │ 0.864 │
└─────────────────────┴───────┘
Job Info
Total runtime: 2m 23s
Results written to .runme/evals/jobs/2026-07-05__11-44-20/result.json
Inspect results by running `harbor view .runme/evals/jobs`
Share results by running `harbor upload .runme/evals/jobs/2026-07-05__11-44-20`
This is the applied unit of work: run Codex against the packaged end-to-end task [11], collect the artifact , and keep the trajectory evidence [12] needed to judge whether the skill actually activated, called the expected tools, followed the workflow, used good sources, covered the target slate, respected guardrails, and produced the expected report without drifting into betting-tip prose.
A local eval run can modify the working tree. That is intentional: if the task asks the agent to edit files, generate artifacts, or update state, the eval should see the same repo changes you would get by driving the agent harness yourself.
In normal use, skip the model and effort knobs and let the selected agent behave as it would if you drove it directly. During iterative development, overrides like lower reasoning effort are still useful for probing weaker configurations, reproducing failures, or deliberately forcing a regression.
$ runme eval compare skills/world-cup-picks-report/evals/regression
Base: .runme/evals/jobs/2026-06-30__11-40-20 tracked in HEAD
Latest: .runme/evals/jobs/2026-07-05__11-44-20 local
Dataset: skills/world-cup-picks-report/evals/regression
Agent: runme-codex
Model: gpt-5.5
Environment: runme_harbor.environment:RunmeEnvironment
Job:
completed: 1 -> 1 +0
errors: 0 -> 0 +0
evals: 1 -> 1 +0
Results:
regression: reward 0.979 -> 0.864 -0.115
Recommendation: candidate regressed; rerun or inspect job/task details before promotion.
In this case, the latest run was intentionally forced into a weaker configuration during development, so the comparison should catch the drop. It compares the local job against the Git-tracked baseline and turns the weaker result into an explicit regression instead of another anecdotal transcript.
To improve the result, edit the skill, adjust context files such as CLAUDE.md or AGENTS.md , change MCP server configuration, or rerun the eval in a sandboxed environment with --env docker and a task-local Dockerfile before deciding what to promote.
Before promoting a result, preview the evidence and the gate Runme would apply:
$ runme eval promote --latest --artifacts --dry-run skills/world-cup-picks-report/evals/regression
Selected eval job: .runme/evals/jobs/

[truncated]
