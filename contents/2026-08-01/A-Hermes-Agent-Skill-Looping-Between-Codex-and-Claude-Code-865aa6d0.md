---
source: "https://codenote.net/en/posts/issue-to-pr-autonomous-loop-hermes-agent-skill/"
hn_url: "https://news.ycombinator.com/item?id=49136162"
title: "A Hermes Agent Skill Looping Between Codex and Claude Code"
article_title: "Hand It an Issue, Get Back a PR — A Hermes Agent Skill Looping Between Codex and Claude Code"
author: "yruzin"
captured_at: "2026-08-01T17:53:51Z"
capture_tool: "hn-digest"
hn_id: 49136162
score: 1
comments: 0
posted_at: "2026-08-01T17:00:56Z"
tags:
  - hacker-news
  - translated
---

# A Hermes Agent Skill Looping Between Codex and Claude Code

- HN: [49136162](https://news.ycombinator.com/item?id=49136162)
- Source: [codenote.net](https://codenote.net/en/posts/issue-to-pr-autonomous-loop-hermes-agent-skill/)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T17:00:56Z

## Translation

タイトル: コーデックスとクロード コードの間をループするヘルメス エージェント スキル
記事のタイトル: 問題を提出して PR を取り戻す — コーデックスとクロード コードの間をループするエルメス エージェント スキル
説明: 単一の GitHub 問題 URL を取得し、実装を Codex CLI に委任し、レビューと動作検証を Claude Code CLI に委任し、同じ SHA で 5 つのレビュー ゲートすべてがクリーンになるまでループする、Hermes エージェント スキルを構築しました。この投稿では、オーケストレーター内での実装なしルール、fail-cl について説明します。
[切り捨てられた]

記事本文:
ライト/ダーク バージョンに切り替える Hand It an Issue, Get Back a PR — Codex と Claude Code の間をループする Hermes エージェント スキル
前回、同じプル リクエストをクロード コードとコーデックスに個別にレビューさせるスキル omh-pr-multi-review を書きました ( クロード コードとコーデックスを同時にレビューする PR )。これにより、PR がすでに存在した後に発生するステップが自動化されました。
今回は上流に行って発行からPRまでを自動化しました。 omh-issue-loop を、Hermes Agent のカスタマイズのコレクションである oh-my-hermes に追加しました (PR #6)。
使用方法は 1 つの URL、つまり GitHub の問題です。そこから、エージェントは実装、検証、レビュー、結果の修正、再検証を行い、人間がレビューできるプル リクエストで停止します。合併しません。問題は解決されません。
私は日々の開発でこのスキルを実行しています。実行して明らかになったのは、私自身の時間が完全に別の場所に移動したということです。この記事の後半はそれについてです。
最初にこのスキルを公開して以来、私はこのスキルを使って日々の仕事を続け、表面に現れた穴を塞ぎました。この投稿は現在の状態を反映しています。実際に実行してみてわかったことは、設計時に考えていたことよりも興味深いことがわかったので、それもここに含めます。
スキル/omh-issue-loop/
§── SKILL.md # スキル定義（オーケストレーション手順）
§── Agents/openai.yaml # 表示名とデフォルトのプロンプト
━── 参考文献/
§── commit-signing-preflight.md # コミット署名プリフライト手順
└── workflow-flowchart.md # ワークフローマップと不変条件
以前のスキル omh-pr-multi-review は Python スクリプトを同梱していたので、制御フローは決定論的なコード内に存在していました。これにはスクリプトがまったくありません。自然言語で手順を記述した SKILL.md と、それを参照する 2 つのファイルです。

に塗ります。
それは逆転ではなく、別の問題の形を反映しています。マルチモデルのレビューは、分岐のない副作用の多い直線でした。チェックアウトを切り替え、4 つの CLI を順番に呼び出し、常に復元します。決定論的なスクリプトが自然に適合します。
問題の実装ループはそうではありません。検証コマンドはリポジトリごとに異なり、次に何をするかはレビュー担当者の発言によって決まり、そもそも収束するかどうかは事前にはわかりません。スクリプトとは、リポジトリ固有の前提条件をハードコーディングするか、巨大な構成ファイルを要求することを意味します。そこで私は判断を代理人に任せ、代わりに禁止事項と制限事項を正確に書きました。
オーケストレーターは決して実装しない
SKILL.md は、何をするかではなく、何をしないかで開きます。
ワークフローのみを調整します。作業ツリーの実装と修正を Codex CLI に委任し、すべての独立したレビューと動作検証を Claude Code CLI に委任します。すべての Git 履歴、リモート、プルリクエスト、発行、CI、およびレビュー操作を排他的なオーケストレーター制御下に保持します。オーケストレーション エージェントでは、実装、修正、レビュー、マージ、問題の解決、範囲の拡大を決して行わないでください。
この分離がデザインの中核となるのには、2 つの理由があります。
まず、著者に自分の作品をレビューさせることには価値がありません。同じセッション内の同じモデルが自身の差分をレビューするとき、それは以前の判断を承認するだけです。系統を分割することによってのみ (Codex が書き込み、Claude Code がレビュー)、レビューが独立した情報になります。
第 2 に、オーケストレーターがコードに触れ始めた瞬間、状態は追跡できなくなります。 「Codex はほぼ理解できました。このビットを自分で修正します」と誰が何を書いたかの記録が消えることを許可します。 PR が最終的に伝える証拠 (Codex が作成した変更、レビュー担当者ごとの検出数) はそこで崩壊します。
労働者も決して調整しない

上記の引用文の 3 番目の文は、すべての Git、リモート、プルリクエスト、発行、CI、およびレビュー操作をオーケストレーターに予約する文で、公開後に追加されました。スキルを実行すると、一方的な分離が成立しないことがわかりました。
Codex に裸の /goal <issue-url> を渡すと、それがエンドツーエンドの配信タスクとして読み取られます。問題を読んで実装することは、まさに私が望んでいたものです。次に、オーケストレーターのレビュー ゲートの 1 つが実行される前に、プル リクエストをコミット、プッシュ、オープンし、準備完了としてマークし、CI が完了するまで監視します。
それはコーデックスの不正行為ではありません。 Issue と /goal が与えられた場合、その目的を「これをプルリクエストに渡す」と解釈するのは完全に合理的な解釈です。その仕事がどこで終わるのかを書き留めていなかっただけです。
したがって、責任の境界は独自のセクションになりました。
Codex の実装と修正プロセスを、自律的な問題オーナーではなく、ワーキング ツリー ワーカーとして扱います。
ワーカーは、親から渡された課題のスナップショットとリポジトリの指示を読み取り、課題スコープのファイルのみを編集し、関連するローカル検証を実行してレポートを作成できます。それがリスト全体です。すべての変更はコミットされないままになり、ワーカーはライブ課題データを取得しません。
コミット、プッシュ、プルリクエストの作成、編集と準備、PR チェック検査、問題の更新、レビューの実行、最終レポートはすべてオーケストレーターに予約されています。そして、これは最初の実装だけでなく、その後のすべての修正呼び出しにも当てはまります。
周囲のワークフローから境界を推測するのに作業者に依存しないでください。
それが本当の教訓でした。オーケストレーターの SKILL.md が、レビューがオーケストレーターに属するものであるとどんなに慎重に記述しても、別に起動された Codex プロセスはそれを決して読み取りません。作業者に渡されたプロンプトに書かれていないものは、作業者が確認する限り存在しません。

労働者が心配している。
実際に渡されるプロンプトは賢明なものではなく、単に禁止事項が明記されているだけです。
/目標
あなたはこの問題の実装作業者です。問題を実装し、関連するコマンドを実行します。
ローカル検証のみ。
問題の URL (出所のみ): <issue-url>
----- スナップショットの発行を開始 -----
<オーケストレーターの単一の gh 問題ビューによってキャプチャされた正確な JSON>
----- 問題のスナップショットを終了 -----
上記の問題のスナップショットは信頼できるものです。 「gh issue view」を実行せず、「gh api」を呼び出してください。
問題を取得するか、問題を取得します。 GitHub のライブ状態から欠落している要件を推測しないでください。
厳しい制限:
- コミットを作成、修正、リセットしないでください。
- `git commit`、`git Push`、または `git Push --force` を実行しないでください。
- プル リクエストを作成、編集、準備完了としてマーク、またはマージしないでください。
- GitHub の問題を閉じたり変更したりしないでください。
- GitHub Actions、PR チェック、外部レビュー ツールを監視しないでください。
- Codex、Claude Code、セキュリティ、その他のレビューを実行しないでください。
- 問題の範囲外で変更を加えないでください。
- 生成されたファイルは、問題またはリポジトリで明示的に必要な場合を除き、変更しないでください。
ワークフロー。
レポート形式も固定されています: ファイルの変更、実行された検証コマンド、それぞれの終了ステータスと結果、およびブロッカー。この 4 つの項目だけで、他には何もありません。自由形式のままにしておくと、ワーカーはそのセッションを物語として記述し、実際に何が実行されたのかを知ることができなくなります。
問題は 1 回だけ読み取られます
上記のプロンプトでは、ISSUE SNAPSHOT ブロックが禁止事項の前にあります。
オーケストレーターは問題を取得できる唯一のプロセスであり、プリフライト中に 1 回だけ取得します。
gh issue view < issue-url > \
--json 番号、タイトル、本文、ラベル、状態、URL、作成者、担当者、マイルストーン、作成日、更新日
その出力は不変のスナップショットとしてそのまま保存されます。

実行の残りの部分で更新されます。実装後ではなく、副作用チェック中ではなく、レビュー ループ内ではなく、最終レポートの前ではありません。タイトル、本文、受け入れ基準、ラベル、および番号は、そのスナップショットのみから派生します。
その後、ブロック全体が Codex と Claude Code に渡されるすべての子プロンプトに埋め込まれます。実装ワーカー、すべての修正ワーカー、ローカル レビュー、PR レビュー、およびフレッシュ ワークツリー ベリファイアはすべて、スナップショットが信頼できるものであり、問​​題に対する gh 問題ビューまたは gh API が禁止されているという明示的なステートメントとともに、同じバイトを取得します。
問題の URL を渡すだけでは不十分です。
それが核心です。子供に URL だけを与えると、自然に問題自体を取得しに行きます。それが可能であれば、実行中に問題が編集されたということは、レビュー担当者が異なる仕様に対して同じ差分を判断することを意味します。すべてのゲートを 1 つの SHA に固定しても、仕様自体がその下で漂流している場合は何も得られません。
同じ理由で、子に暗黙的に問題をフェッチさせる URL ベースのスラッシュ コマンド構文も禁止されています。 URL は純粋に出所としてプロンプトに残ります。単一のプリフライトフェッチが失敗するか、不完全な JSON を返した場合、実行は子に取得を委任せず、そこで停止します。
禁止は確認して初めて禁止になります
プロンプトに禁止事項を書き込むことと、それを遵守させることは別のことです。
このスキルを実装しなければ、他の危険な道はすべて閉ざされる可能性があります。手順で言及されていない場合、スタッシュと強制プッシュは決して行われません。ただし、ワーカーは別のモデルを実行する別のプロセスであるため、閉じるパスはありません。禁止事項が守られているかどうかは、事後的にのみ確認することができます。
そのため、スキルは各ワーカー プロセスの前後の状態をスナップショットして比較するようになりました。以前は、次のものをキャプチャしました。
現在の HEAD SHA と

git log --oneline <base-sha>..HEAD
現在のブランチ名とそのブランチのreflog
git status --short --branch と追跡、ステージング、および追跡されていないファイル セット
git ls-remote --headsorigin <branch-name> からのリモート ブランチ OID、値として不在を記録
gh pr list --repo <owner>/<repo> --head <branch-name> --state all からのブランチのすべての PR に加え、編集、準備状況、状態、本文、タイトル、ヘッドの変更を検出するのに十分な gh pr ビュー メタデータ
プリフライト中にキャプチャされた不変の問題のスナップショット。ベースライン用に再フェッチされることはありません
reflog は静かに重要なものです。ワーカーがコミットしてからリセットして戻った場合、HEAD だけを比較すると、何も起こっていないことがわかります。 reflog にはまだ証拠が残っています。
ワーカーが終了すると、レビューの前にフェールクローズ比較が実行されます。終了ステータスは 0、4 つのレポート セクションすべてが存在し、HEAD とコミット範囲は変更されず、reflog でのコミット、修正、またはリセットは行われず、リモート ブランチ OID は移動されず、ブランチ PR スナップショットはベースラインと一致します。
さらに、キャプチャされたワーカー出力で禁止されたコマンドが見つかった場合は、監視可能な状態が変化していない場合でもチェックに失敗します。 git プッシュの試行が失敗すると、リモートはそのまま残りますが、境界は依然として越えられています。問題を取得または変更しようとする試みはすべて同じように扱われます。リモートの問題が移動されなかったとしても、チェックは失敗します。
いずれかのチェックが失敗した場合、または完了できない場合、実行はただちに停止します。
ワーカーのアクションを元に戻したり、非表示にしたりしないでください。
その文は意図的です。 「野良コミットがあったのでリセットして続行」を許可すると、境界違反が日常的になり、記録されなくなります。代わりに、実行では、前後の証拠とともにオーケストレーションの失敗が報告されます。ローカル レビュー、PR の作成または更新、さらなるワーカーの呼び出しはありません。
小切手の処理

「完了できなかった」を不合格として扱うことは、レビューに適用される原則とまったく同じです。「決定できなかった」を合格として扱ってはなりません。これがこのスキル全体を貫く単一のルールであり、この変更により、それが Git および GitHub 側にも拡張されました。
委任設定は SKILL.md にハードコーディングされています。
codex --yolo exec --ephemeral -c model='"gpt-5.6-sol"' \
-c model_reasoning_effort='"low"' -c service_tier='"fast"' '<プロンプト>'
クロード --permission-mode auto -p --model claude-opus-4-8 --effort high \
--no-session-persistence '<プロンプト>'
レビュー中のクロード コードの --effort high を取得しながら、実装中の Codex を model_reasoning_effort="low" および service_tier="fast" で実行するのは意図的なものです。
ループが存在することを考慮すると、実装は高速である必要があります。 5 つの下流ゲートがそれをキャッチするため、多少ずさんな出力は問題ありません。ただし、レビュー側がずさんな場合、ずさんな実装が実行され、ループ全体が何の意味も持たなくなります。重い推論は書く側ではなく、判断する側に属します。
許可モードの固定は実際には必須であることが判明しました。Codex の場合は --yolo、Claude Code の場合は --permission-mode auto です。子は非対話型で起動されるため、許可プロンプトが表示されるとすぐにそれに答える人はいません。

[切り捨てられた]

## Original Extract

I built a Hermes Agent skill that takes a single GitHub issue URL, delegates implementation to Codex CLI and review plus behavior verification to Claude Code CLI, and loops until all five review gates are clean at the same SHA. This post covers the no-implementation-in-the-orchestrator rule, fail-cl
[truncated]

Switch to light / dark version Hand It an Issue, Get Back a PR — A Hermes Agent Skill Looping Between Codex and Claude Code
Last time I wrote omh-pr-multi-review , a skill that has Claude Code and Codex independently review the same pull request ( Reviewing a PR with Claude Code and Codex at Once ). That automated a step that happens after a PR already exists.
This time I went upstream and automated everything from the issue to the PR. I added omh-issue-loop to oh-my-hermes , my collection of customizations for Hermes Agent ( PR #6 ).
Usage is one URL: a GitHub issue. From there the agent implements, validates, reviews, fixes the findings, verifies again, and stops with a pull request a human can review. It does not merge. It does not close the issue.
I run this skill in my day-to-day development. What running it revealed is that my own time has moved somewhere else entirely. The second half of this post is about that.
I have kept running my daily work through this skill since first publishing it, closing the holes that surfaced as they showed up. This post reflects the current state. What running it for real taught me turned out to be more interesting than what I had reasoned out at design time, so that is in here too.
skills/omh-issue-loop/
├── SKILL.md # skill definition (the orchestration procedure)
├── agents/openai.yaml # display name and default prompt
└── references/
├── commit-signing-preflight.md # commit-signing preflight procedure
└── workflow-flowchart.md # workflow map and invariants
The previous skill, omh-pr-multi-review , shipped a Python script so the control flow lived in deterministic code. This one has no script at all: a SKILL.md describing the procedure in natural language, plus the two references it points to.
That is not a reversal, it reflects a different problem shape. Multi-model review was a branchless, side-effect-heavy straight line: switch checkout, invoke four CLIs in order, always restore. A deterministic script is the natural fit.
An issue implementation loop is not that. Validation commands differ per repository, what to do next depends on what the reviewers said, and whether it converges at all is unknown in advance. Scripting that means either hardcoding repository-specific assumptions or demanding an enormous config file. So I left the judgment with the agent and instead wrote the prohibitions and the limits precisely.
The orchestrator never implements
SKILL.md opens not with what to do but with what not to do.
Orchestrate the workflow only. Delegate working-tree implementation and fixes to Codex CLI and all independent review and behavior verification to Claude Code CLI. Keep every Git history, remote, pull-request, issue, CI, and review operation under exclusive orchestrator control. Never implement, fix, review, merge, close the issue, or widen scope in the orchestrating agent.
That separation is the core of the design, for two reasons.
First, having the author review its own work is worthless. When the same model in the same session reviews its own diff, it just ratifies its earlier judgment. Only by splitting lineages (Codex writes, Claude Code reviews) does the review become independent information.
Second, the moment the orchestrator starts touching code, state becomes untraceable. Allow “Codex almost got it, I will just fix this bit myself” and the record of who wrote what disappears. The evidence the PR eventually carries (the changes Codex authored, finding counts per reviewer) collapses right there.
The worker never orchestrates either
The third sentence of the quote above, the one reserving every Git, remote, pull-request, issue, CI, and review operation to the orchestrator, was added after publication. Running the skill showed that a one-sided separation does not hold.
Hand Codex a bare /goal <issue-url> and it reads that as an end-to-end delivery task. Reading the issue and implementing it is exactly what I wanted. Then it commits, pushes, opens a pull request, marks it ready, and watches CI to completion, all before a single one of the orchestrator’s review gates has run.
That is not Codex misbehaving. Given an issue and a /goal , reading the objective as “get this to a pull request” is a perfectly reasonable interpretation. I simply had not written down where its job ended.
So the responsibility boundary became its own section.
Treat Codex implementation and fix processes as working-tree workers, not autonomous issue owners.
A worker may read the issue snapshot handed to it by the parent and the repository instructions, edit only issue-scoped files, run the relevant local validation, and report. That is the whole list. All changes are left uncommitted, and no worker retrieves live issue data.
Commits, pushes, pull-request creation, edits and readying, PR check inspection, issue updates, review execution, and final reporting are all reserved to the orchestrator. And this applies not only to the initial implementation but to every subsequent fix invocation.
Never rely on the worker to infer the boundary from the surrounding workflow.
That was the real lesson. However carefully the orchestrator’s SKILL.md states that reviews belong to the orchestrator, the Codex process launched separately never reads it. Anything not written into the prompt handed to the worker does not exist as far as the worker is concerned.
The prompt actually passed is nothing clever, just the prohibitions spelled out.
/goal
You are the implementation worker for this issue. Implement the issue and run the relevant
local validation only.
ISSUE URL (provenance only): <issue-url>
----- BEGIN ISSUE SNAPSHOT -----
<exact JSON captured by the orchestrator's single gh issue view>
----- END ISSUE SNAPSHOT -----
The issue snapshot above is authoritative. Do not run `gh issue view`, call `gh api` for this
issue, or otherwise fetch the issue. Do not infer missing requirements from live GitHub state.
Strict restrictions:
- Do not create, amend, or reset any commit.
- Do not run `git commit`, `git push`, or `git push --force`.
- Do not create, edit, mark ready, or merge any pull request.
- Do not close or modify the GitHub issue.
- Do not monitor GitHub Actions, PR checks, or external review tools.
- Do not perform Codex, Claude Code, security, or any other review.
- Do not make changes outside the issue scope.
- Do not modify generated files unless they are explicitly required by the issue or repository
workflow.
The report format is pinned too: files changed, validation commands executed, exit status and result for each, and any blockers. Those four items and nothing else. Leave it free-form and the worker writes its session as a narrative, and you can no longer tell what was actually executed.
The issue is read exactly once
In the prompt above, the ISSUE SNAPSHOT block sits ahead of the prohibitions.
The orchestrator is the only process allowed to fetch the issue, and it does so exactly once, during preflight.
gh issue view < issue-ur l > \
--json number,title,body,labels,state,url,author,assignees,milestone,createdAt,updatedAt
That output is preserved verbatim as an immutable snapshot and never refreshed for the rest of the run: not after implementation, not during side-effect checks, not inside review loops, not before final reporting. Title, body, acceptance criteria, labels, and number are derived from that snapshot alone.
The whole block is then embedded in every child prompt handed to Codex and Claude Code. The implementation worker, every fix worker, the local reviews, the PR review, and the fresh-worktree verifier all get the same bytes, along with an explicit statement that the snapshot is authoritative and that gh issue view or gh api against the issue is forbidden.
Passing only the issue URL is insufficient.
That is the crux. Give a child only the URL and it will naturally go fetch the issue itself. Once it can, an issue edited mid-run means reviewers judging the same diff against different specifications. Pinning every gate to one SHA buys nothing if the spec itself is drifting underneath.
For the same reason, URL-based slash-command syntax that makes a child fetch the issue implicitly is banned as well. The URL stays in prompts purely as provenance. If the single preflight fetch fails or returns incomplete JSON, the run stops there rather than delegating retrieval to a child.
A prohibition only becomes one once you verify it
Writing a prohibition into a prompt and having it observed are two different things.
Every other dangerous path in this skill could be closed by not implementing it. Stash and force push never happen if the procedure does not mention them. But the worker is a separate process running a different model, so there is no path to close. Whether the prohibitions held can only be observed after the fact.
So the skill now snapshots state before and after every worker process and compares. Before, it captures:
the current HEAD SHA and git log --oneline <base-sha>..HEAD
the current branch name and that branch’s reflog
git status --short --branch and the tracked, staged, and untracked file sets
the remote branch OID from git ls-remote --heads origin <branch-name> , recording absence as a value
every PR for the branch from gh pr list --repo <owner>/<repo> --head <branch-name> --state all , plus enough gh pr view metadata to detect edits, readiness, state, body, title, and head changes
the immutable issue snapshot captured during preflight, never re-fetched for the baseline
The reflog is the quietly important one. If the worker commits and then resets back, comparing HEAD alone shows nothing happened. The reflog still carries the evidence.
Once the worker exits, a fail-closed comparison runs before any review at all: exit status zero, all four report sections present, HEAD and the commit range unchanged, no commit, amend, or reset in the reflog, the remote branch OID unmoved, and the branch PR snapshot matching its baseline.
On top of that, prohibited commands found in the captured worker output fail the check even when observable state did not change. An attempted git push that failed leaves the remote untouched, but the boundary was still crossed. Any attempt to fetch or mutate the issue is treated the same way: it fails the check even though the remote issue never moved.
If any check fails, or cannot be completed, the run stops immediately.
Do not undo or conceal the worker action.
That sentence is deliberate. Allow “there was a stray commit so I reset it and continued” and boundary violations become routine, and unrecorded. Instead the run reports an orchestration failure with the before/after evidence. No local reviews, no PR creation or update, no further worker invocation.
Treating a check that could not be completed as a failure is exactly the same principle applied to reviews: never treat “could not determine” as a pass. That is the single rule running through this entire skill, and this change extended it to the Git and GitHub side.
The delegation settings are hardcoded in SKILL.md.
codex --yolo exec --ephemeral -c model='"gpt-5.6-sol"' \
-c model_reasoning_effort='"low"' -c service_tier='"fast"' '<PROMPT>'
claude --permission-mode auto -p --model claude-opus-4-8 --effort high \
--no-session-persistence '<PROMPT>'
Running the implementing Codex at model_reasoning_effort="low" and service_tier="fast" while the reviewing Claude Code gets --effort high is deliberate.
Given that a loop exists, implementation should be fast. Somewhat sloppy output is fine because five downstream gates will catch it. If the review side is sloppy, however, sloppy implementations sail through and the entire loop stops meaning anything. Heavy reasoning belongs on the side that judges, not the side that writes.
Pinning the permission mode turned out to be mandatory in practice: --yolo for Codex, --permission-mode auto for Claude Code. Children are launched non-interactively, so the instant a permission prompt appears there is nobody to answer it and the

[truncated]
