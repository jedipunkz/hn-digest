---
source: "https://karavox.org/devlog/agents-as-subcontractors.html"
hn_url: "https://news.ycombinator.com/item?id=49361077"
title: "I treat my AI coding agents as subcontractors"
article_title: "I treat my AI coding agents as subcontractors — karavox devlog"
image: ""
author: "lukasznowak"
captured_at: "2026-08-19T13:37:25Z"
capture_tool: "hn-digest"
hn_id: 49361077
score: 2
comments: 0
posted_at: "2026-08-19T13:02:43Z"
tags:
  - hacker-news
  - translated
---

# I treat my AI coding agents as subcontractors

- HN: [49361077](https://news.ycombinator.com/item?id=49361077)
- Source: [karavox.org](https://karavox.org/devlog/agents-as-subcontractors.html)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T13:02:43Z

## Translation

タイトル: 私は AI コーディング エージェントを下請け業者として扱います
記事のタイトル: 私は AI コーディング エージェントを下請け業者として扱います — karavox devlog
説明: 個人開発者

記事本文:
私は AI コーディング エージェントを扱います
下請け業者として
2026-08-14 · ソフトウェア開発を行う方法
2026 年から、主に仕事のために AI エージェントを使用したソフトウェア開発を始めました。私はから始めました
1 つのエージェントが、その後ゆっくりと複数のエージェントに進化しました。さまざまなアプローチを試しましたが、どれもうまくいきませんでした
私の仕事のワークフロー。
私は不信感を抱いています（率直に言って、これを表す正しい英語の単語はわかりません。完全な不信感です）
しかし、どのエージェントも好奇心旺盛で、良い仕事をしてくれるという一種の期待を持っています。彼らは間違いを犯します、
トークン、ユーザー、VM でそれらをガードレールする必要があります。それでも、それらは私の仕事に役立つことが証明されました。
私のエージェントは YOLO モードで実行されます。ファイルの編集、コミット、プッシュについては、そのままの状態で質問されません。
自分たちの部分に責任がある。
職場でも活躍し始めました。私たちのチームは実用的なソリューションを提供しました。降ろすことができました
ワークフローの一部をツールに切り替えて、重要なことに集中します。本当にうまくいきました…
…それで、私は自分の得意なプロジェクト、つまりカラオケ周辺のプロジェクトについて考えるようになりました。次のようにコード化できます
しかし、エージェントの群れを使用しないのはなぜでしょうか。何度も行ったり来たり、私の気持ちが落ち込んだりして、
結局、私は彼らを下請け業者として扱うことになりました。自分のプロジェクトではもっとリスクを取ることができます。
私は個人開発者で、オープンソースの小さなエコシステム全体で複数のエージェントが働いています。
フォーマットとツールキット、それに関連するクローズドソース製品、およびそれを実行するインフラストラクチャ
それらすべて。すべてを形作る制約は単純です。私がボトルネックであり、私は
判断力を持つ唯一の人でもあります。エージェントは多くのことを行うことができますが、それについてのコンテキストはまったくありません。
インシデント履歴、エッジケース、または環境に存在しない運用上の制約
リポジトリ。したがって、設計目標は次のとおりです。エージェントは私がいなくてもできる限りのことを行うことができますが、彼らは
見たことのないものには決して触れることができません。
このモデルは最初から始まったものではありません

自分。それは、役割に名前を付けた投稿から始まりました。
サイモン・ウィリソンのバイブ・エンジニアリング
(2025-10-07) — AI 支援開発の規律ある終焉、そこでは専門家が
バイブコーディングの簡単で緩い終わりに対して、ソフトウェアに対して責任を負い続けます。
Willison 自身の 2026 年の最新情報では、これに勝った用語は次のとおりであると述べています。
エージェントエンジニアリング。の
その後の測定結果が残りの部分を形成しました。
パラレルコーディングエージェントのライフスタイルを受け入れる
(Willison、2025-10-05) — ボトルネックとしてレビュー帯域幅を持つ並列エージェント。
研究/PoC タスクおよび慎重に指定された作業が安全なカテゴリとして分類されます。
2025 年 9 月におけるコーディング エージェントの使用状況
(Jesse Vincent、2025-10-05) — アーキテクトと実装者が分離された git に分割される
ワークツリーとワークツリーの間に PM をプレイする人間がいます。
本番ワークフローで GitHub AI コーディング エージェントを使用するためのベスト プラクティスは何ですか?
(GitHub コミュニティ、2025 年 12 月 17 日) — 「AI エージェントは自律的なものではなく、強力なチームメイトです」
コミッター: エージェントはコードを提案しますが、それを所有することはありません。ドラフト PR のみ。人間関係者
合併契約。
レイヤ 1 — トークン: エージェントは運用環境の近くに書き込むことができません
すべてのエージェントは 2 つのトークンを取得します。本番リポジトリ上の読み取り専用トークン、および
別の -staging リポジトリにトークンを書き込みます。タスクのブランチが切断される
本番環境のメイン ブランチから直接 (読み取りだけで十分です)、ステージングにプッシュされます
リポジトリは、純粋に書き込みトークンが到達できる場所として存在します。
ステージング リポジトリのデフォルト ブランチは、文字通り名前が付けられた意図的な墓石です。
no-main 、README のみが含まれています: 「オリジナルの main ブランチを使用してください」
リポジトリ。」何もそこに溶け込むことはありません。何も同期することはありません。それには歴史がありません、いいえ
ミラー、エージェントのメールボックス以上の意味はありません。
なぜ標準ツールではないのでしょうか?私の計画ではそれらは存在しないから

: GitHub の
ドキュメント
保護されたブランチを無料プランのパブリック リポジトリと
プライベートリポジトリはPro upからのみ。プライベートリポジトリをフォークして
組織
また、無料ではなく GitHub チームが必要です。トークンのスコープ設定は唯一のものです
エージェントが本番環境に触れるのを物理的に防ぐメカニズム - したがって、設計
設定ではなくトークンから保証を構築します。
レイヤ 2 — 統合: 私はマージ ボットです
ブランチの準備ができたら、エージェントが私に告げます。それを取得し、diff を確認して組み込みます
ただし、チェリーピック、リベースマージ、または手動で適用することは可能です。プルリクエスト機構はありません。
エージェントによってマージ コミットが書き込まれることはなく、キューのバックアップ中に読み取られずに残る PR もありません。
新しい服を着た古いパターンです。 Git 自体のドキュメントでは、次のように説明されています。
統合マネージャーのワークフロー: 書き込みアクセス権を持たない貢献者はパッチを送信し、
メンテナがそれらを適用します。それがまさに私がやっていることです。私のエージェントはパッチの提供者であり、
ステージング リポジトリはメールボックスです。 Linux カーネルが使用してきたモデルです
20 年間、差分メールの代わりにブランチのみを使用していました。
1 つのルールにより、これが正直に保たれます。ブランチは、独立して検証されるまで決して削除されません。
本番環境内にある (本番環境に対する git merge-base --is-ancestor
デフォルトのブランチ、またはコミットが潰されたかどうかの同等のチェック）。チェック可能なビート
私自身の言葉も含めて、鵜呑みにした言葉です。
レイヤー 3 — PR ポリシー: 独断ではなく判断
パブリック リポジトリ karavox は PR 専用です。それは交渉の余地がありません。オープンです
情報源では、未知の貢献者に直面しており、そこでは PR が貢献の標準となっています。
プライベートリポジトリは私の判断です。なぜそれが擁護できるのでしょうか？レビューなので、
どちらの場合でも起こります。問題は、それがどの層で起こるかだけです。 PR モデルでは、
レビューする

w は GitHub によって強制される儀式です。私のモデルでは、レビューは統合そのものです。
QA も務める単独のインテグレーターにとって、プル リクエストはオーバーヘッドです。レビューはそうではありません。
私は審査を決して省略しません。式典も省略します。
ベンダーは同じ原則に基づいて集まっています。クロード・コードのセキュリティ
docs : 手動モードでは読み取り専用で開始します
許可、そして「提案されたコードをレビューする責任はあなたにあります。」 OpenAI のコーデックス
docs : デフォルトではサンドボックス化されています。
承認ポリシーを使用する — Codex はアクションを実行する前に質問する必要があります。 GitHub 独自の
エージェント ワークフロー ツール: エージェント ジョブは読み取り専用であり、
デフォルトでサンドボックス化されており、書き込みは検証された安全な出力ジョブを通じて適用されます。
範囲指定された権限。 AI コーディングに関する GitHub 独自のコミュニティ ガイダンス
エージェントは率直にこう言います。「AI
エージェントはコードを提案できますが、コードを所有することはありません。」業界は私の側に集結しつつある
引数 —
ほとんどの場合、ステージング ミラーの削除までは進んでいません。
戦争物語：亡くなったモデル
墓石は最初のデザインではありませんでした。当初、ステージング リポジトリには完全な
文字通り staging という名前のブランチ上の本番環境のミラー: タスク ブランチは
ミラーから切り取られ、ステージングにマージされ、本番環境にプロモートされます。モデル
慣習に従って足並みを揃えるためには、ミラーとステージングという 2 つのことが必要でした
支店自体。 2026 年 8 月 14 日、実際に漂流しました。2 本の枝がまっすぐに着地しました。
ステージング中は本番環境のメインは 2 つのコミットの後ろにありました。
この修正は同期を強化するものではありませんでした。修正はミラーを削除することでした。タスクブランチは現在
プロダクション自体の履歴に直接基づいているため、足並みを揃える必要はありません。
同じ午後にステージング リポジトリが廃止され、ワークフローがよりシンプルになりました
それ以来ずっと。本番環境で機能しなくなるガバナンス モデルはダメだ

d ガバナンスモデル — それ
パッチを適用する代わりに再設計できることが証明されました。
エージェントは判断が始まるところまですべてを行います。
私は重要なところだけに注意を払います。すべての統合は定義上レビューです。
トリアージするための PR キューも、マシンによって書き込まれるマージ コミットも、セレモニーもありません。
パブリック リポジトリは貢献基準を維持します。プライベート リポジトリは速度を維持します。
本番環境のメイン自体は、トークンのスコープと私の規律によってのみ保護されます — ブランチ
保護手段はベルトとブレースですが、私が加入しているプランでは利用できません。そして、
業界のシグナルは明らかです。現在、GitHub のコード レビューの 5 件中 1 件以上に、
代理店。
判断力がボトルネックであり、このモデルは、判断力ではなく、その事実に基づいて構築されています。
ボトルネックが存在しないふりをする。
ソース: GitHub ドキュメント — 保護されたブランチについて
そしてフォーク・
Git Pro ブック — プロジェクトへの貢献 ·
クロード コードのセキュリティ ドキュメント ·
OpenAI Codex — エージェントの承認とセキュリティ ·
GitHub エージェントティック ワークフロー (gh-aw) ·
GitHub コミュニティ: AI コーディング エージェントのベスト プラクティス ·
GitHub ブログ: エージェントのプル リクエスト 。

## Original Extract

A solo developer

I treat my AI coding agents
as subcontractors
2026-08-14 · how I run software development
Since 2026 I started to develop software with AI agents, mostly for work. I started with
one agent, then slowly evolved to more. I tried various approaches, but none really fit
my work workflow.
I distrust (frankly, I do not know the correct English word for this – it’s total distrust
but with a full of curiosity and kind of expectation of good work) every agent. They make mistakes,
I have to guardrail them with tokens, users and VMs. Nevertheless they proven to be useful at my work.
My agents run in YOLO mode – no questions asked about editing files, committing, pushing, as they are
responsible for their part.
It started to work at work. Our team delivered working solutions. I was able to offload
part of my workflow to a tool and focus on what’s important. It really worked…
…and that led me to thinking about my pet project – something around Karaoke. I could code it by
hand, but why not to use swarm of agents. With a lot of back and forth and with my low
trust I ended up to treat them as subcontractors . I could take more risks with my own project.
I’m a solo developer with several agents working across a small ecosystem: an open-source
format and toolkit, closed-source products around it, and the infrastructure that runs
them all. The constraint that shapes everything is simple: I’m the bottleneck, and I’m
also the only one with judgment. Agents can do a lot, but they have zero context about
my incident history, my edge cases, or the operational constraints that don’t live in the
repository. So the design goal is: agents may do as much as possible without me — but they
can never touch anything I haven’t seen.
This model didn’t start with me. It started with a post that gave the role a name:
Simon Willison’s vibe engineering
(2025-10-07) — the disciplined end of AI-assisted development, where a professional
stays accountable for the software, against the fast-and-loose end of vibe coding.
Willison’s own 2026 update notes the term that won out for this is
Agentic Engineering . The
readings that followed shaped the rest:
Embracing the parallel coding agent lifestyle
(Willison, 2025-10-05) — parallel agents with review bandwidth as the bottleneck;
research/PoC tasks and carefully-specified work as the safe categories.
How I’m using coding agents in September, 2025
(Jesse Vincent, 2025-10-05) — an architect/implementer split across isolated git
worktrees, with a human playing PM between them.
Best practices for using GitHub AI coding agents in production workflows?
(GitHub Community, 2025-12-17) — “AI agents are powerful teammates, not autonomous
committers”: agents propose code, never own it; draft PRs only; a human-in-the-loop
merge contract.
Layer 1 — the tokens: agents can’t write near production
Every agent gets two tokens. A read-only token on the production repository, and a
write token on a separate -staging repository. Task branches are cut
directly from production’s main branch (read is enough for that) and pushed to the staging
repository, which exists purely as a place the write token can reach.
The staging repository’s default branch is a deliberate tombstone, literally named
no-main , containing only a README: “please use main branch of the original
repository.” Nothing ever merges into it. Nothing ever syncs it. It has no history, no
mirror, no meaning beyond being the agents’ mailbox.
Why not the standard tools? Because on the plan I’m on, they don’t exist: GitHub’s
docs
make protected branches available in public repositories on the free plan and in
private repositories only from Pro up; forking a private repository into an
organization
also requires GitHub Team, not Free. Token scoping is the only
mechanism that physically prevents an agent from touching production — so the design
builds the guarantee out of tokens instead of settings.
Layer 2 — integration: I am the merge bot
When a branch is ready, the agent tells me. I fetch it, review the diff, and incorporate
it however fits: cherry-pick, rebase-merge, or apply by hand. No pull request machinery,
no merge commits written by agents, no PRs that sit unread while the queue backs up.
This is an old pattern wearing new clothes. Git’s own documentation describes it as the
integration-manager workflow : contributors without write access submit patches, and
a maintainer applies them. That’s exactly what I do — my agents are patch contributors and
the staging repository is their mailbox. It’s the model the Linux kernel has used for
twenty years, just with branches instead of emailed diffs.
One rule keeps this honest: a branch is never deleted until it’s independently verified to
be inside production ( git merge-base --is-ancestor against the production
default branch, or the equivalent check if the commits were squashed). Checkable beats
taken-on-word — including my own.
Layer 3 — the PR policy: judgment, not dogma
The public repository, karavox , is PR-only. That’s non-negotiable: it’s open
source, it faces unknown contributors, and PRs are the contribution norm there.
Private repositories are my judgment call. Why is that defensible? Because the review
happens either way — the question is only which layer it happens at. In the PR model the
review is ceremony enforced by GitHub; in my model the review is the integration itself.
For a solo integrator who is also the QA, the pull request is overhead; the review is not.
I never skip the review — I skip the ceremony.
The vendors converge on the same principles. Claude Code’s security
docs : in manual mode it starts with read-only
permissions, and “you’re responsible for reviewing proposed code.” OpenAI’s Codex
docs : sandboxed by default,
with an approval policy — Codex must ask before it executes actions. GitHub’s own
agentic workflow tool : agent jobs are read-only and
sandboxed by default, and writes are applied through validated safe-outputs jobs with
scoped permissions. GitHub’s own community guidance for AI coding
agents puts it bluntly: “AI
agents can propose code, never own it.” The industry is converging on my side of this
argument —
most of it just hasn’t gone as far as deleting the staging mirror.
The war story: the model that died
The tombstone wasn’t the first design. Originally the staging repository carried a full
mirror of production on a branch literally named staging : task branches were
cut from the mirror, merged into staging, and then promoted to production. The model
required two things to stay in lockstep by convention — the mirror, and the staging
branch itself. On 2026-08-14 it drifted for real: two branches landed straight on
production main while staging sat two commits behind.
The fix wasn’t hardening the sync. The fix was deleting the mirror. Task branches are now
based directly on production’s own history, so there is nothing left to stay in lockstep.
The staging repo became a tombstone the same afternoon, and the workflow has been simpler
ever since. A governance model that dies in production is a good governance model — it
proved it could be redesigned instead of patched.
Agents do everything up to the point where judgment starts.
I only spend attention where it matters — every integration is a review by definition.
No PR queue to triage, no merge commits written by machines, no ceremony.
The public repo keeps the contribution norm; the private repos keep the speed.
Production main itself is protected only by token scoping plus my discipline — branch
protection would be belt-and-braces, but it’s not available on the plan I’m on. And the
industry signal is clear: more than one in five code reviews on GitHub now involves an
agent .
Judgment is the bottleneck, and this model is built around that fact rather than
pretending the bottleneck doesn’t exist.
Sources: GitHub docs — about protected branches
and forks ·
Git Pro book — contributing to a project ·
Claude Code security docs ·
OpenAI Codex — agent approvals & security ·
GitHub Agentic Workflows (gh-aw) ·
GitHub Community: best practices for AI coding agents ·
GitHub blog: agent pull requests .
