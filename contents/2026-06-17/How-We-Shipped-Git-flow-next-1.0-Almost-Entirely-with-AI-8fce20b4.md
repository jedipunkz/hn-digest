---
source: "https://git-flow.sh/blog/posts/how-we-shipped-git-flow-next-1-0-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=48571907"
title: "How We Shipped Git-flow-next 1.0 Almost Entirely with AI"
article_title: "How We Shipped git-flow-next 1.0 Almost Entirely with AI | git-flow-next"
author: "speter"
captured_at: "2026-06-17T16:21:17Z"
capture_tool: "hn-digest"
hn_id: 48571907
score: 2
comments: 0
posted_at: "2026-06-17T15:33:31Z"
tags:
  - hacker-news
  - translated
---

# How We Shipped Git-flow-next 1.0 Almost Entirely with AI

- HN: [48571907](https://news.ycombinator.com/item?id=48571907)
- Source: [git-flow.sh](https://git-flow.sh/blog/posts/how-we-shipped-git-flow-next-1-0-with-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T15:33:31Z

## Translation

タイトル: Git-flow-next 1.0 をほぼ完全に AI で出荷した方法
記事のタイトル: git-flow-next 1.0 をほぼ完全に AI で出荷した方法 | git-flow-next
説明: AI を使用して git-flow-next を構築する 13 か月にわたる正直な旅 - 何がうまくいき、何が失敗し、なぜガイドラインが実際の製品になったのか。

記事本文:
git-flow-next 1.0 をほぼ完全に AI で出荷した方法 | git-flow-next
ドキュメント
フェーズ 1: カーソルを使用した Vibe コーディング (2025 年 3 月)
フェーズ 2: システムの構築 (2025 年 5 月～8 月)
フェーズ 3: 転換点 (2025 年 9 月)
フェーズ 4: モデルの追いつき (2025 年 10 月～11 月)
フェーズ 5: ワークフローの記入 (2025 年 12 月～2026 年 1 月)
フェーズ 6: v1.0.0 とスキル システム (2026 年 2 月)
フェーズ 7: AI コード レビューの実際の様子 (2026 年 2 月～4 月)
git-flow-next 1.0 をほぼ完全に AI で出荷した方法
執筆者: Alexander Rinass 発行日: 2026 年 6 月 9 日
1.0 の発表 で、git-flow-next は「この範囲のオープンソース プロジェクトで今日の AI 機能をどこまで拡張できるかを確認する」実験でもあったと述べました。
この旅については、今後の投稿でさらに詳しくお話しすることをお約束します。これです。
約 29,000 行の Go コード
私たちはこれらすべてを通常の仕事と並行して副業で行いました。これは、タイムラインを深読みする前に覚えておく価値があります。
コードのほとんどはAIによって書かれました。私たちは経験豊富な開発者ではありますが、Go には熟練していません。そこで私たちはアーキテクチャ、仕様、ガイダンスを提供しました。 AI はキーボードを握る手でした。
プロジェクトの動機は次の 2 つでした。
私たちは、継続的なメンテナンスをほとんど必要としないオープンソース ツールを望んでいました。
私たちは、技術スタックに詳しくない私たちにとって多かれ少なかれブラック ボックスに近いものを実験して、具体的に、そしてより一般的な質問として、AI を使用した構築がそのような条件下でどのように機能するかを確認したいと考えていました。
急いでいて要点だけを探している人のために、次のプロジェクトに引き継ぐ教訓を 1 つ挙げておきます。AI 開発は「AI に何を構築するかを指示する」ものではありません。それは「AIに何を伝えるシステムを構築するか」です。

そしてその構築方法についても説明します。」
コードは出力です。ガイドライン、スキル、レビュー基準、およびアーキテクチャに関するドキュメントが実際の製品です。
フェーズ 1: カーソルを使用した Vibe コーディング (2025 年 3 月)
最初のフェーズは、AI が何を達成できるかを探るためだけに、最小限の命令をすべて 1 つのファイルにまとめて、実際に取り組んでいくことでした。
そのペースは驚くべきものだった。約 2 週間で、Cursor は git-flow-next のすべてのコア コマンド ( start 、finish 、 list 、overview 、 update 、 delete 、 rename 、 checkout ) をほぼ 1 ～ 2 日ごとに 1 つの新しいコマンドで生成しました。
よく見ると、このプロジェクトは実際にはプロジェクトではありませんでした。これは、偶然ディレクトリを共有した、個別に生成されたファイルのコレクションでした。 CLI レイヤー、ビジネス ロジック、Git 操作の間に明確な分離はなく、すべてのコマンドが独自の方法で処理を実行しました。パターンはわずかに異なる形でファイル間で複製されており、多くの場合、AI に重複の検索を依頼すると、1 つが表面化して他のものは静かに見逃されるほど、微妙に複製されていました。 Git 設定は起動時に 1 回ではなく、操作ごとに読み込まれていました。これは偶然ではなく仕様によるパフォーマンスの低下です。
私たちはそれから抜け出す方法をリファクタリングしようとしました。カーソル付き、クロード付き、詳細な説明書付き。うまくいきませんでした。関数は重複し、コードはコンパイルされず、提案されたリファクタリングは間違った抽象化でオーバーエンジニアリングされました。
ゼロから完全に書き直すことも試みました。それらも失敗しました。
最終的にうまくいったのは、速度とは逆の、小さな段階的なステップでした。最初に空のコマンド ファイルを作成します。ロジックを 1 つ追加します。テストを実行します。先に進んでください。そしてさらに重要なことは、AI で何かをクリーンアップできるようになる前に、開発哲学を自分たちで明示的に書面で定義する必要があることに気づきました。
自分が何を望んでいるのかが分かっていなければ、AI の提案を評価することはできません。 T

私たちはこのとき初めて、本当の問題の形を理解しました。
フェーズ 2: システムの構築 (2025 年 5 月～8 月)
Claude Code を使用した初日に、プロジェクトを分析して CLAUDE.md ファイルを作成しました。一貫性のない設定キーなど、いくつかの間違いがありましたが、正確さは実際には重要ではありませんでした。
重要なのは、コードを記述する前にプロジェクトの理解を開始する場所を AI に提供することでした。
それがパターンを生み出しました。 AI が認識できる形で失敗するたびに、その答えは新しいガイドラインでした。数か月の間に、そのうちの 8 つが蓄積されました。
CLAUDE.md — このコードベースで動作する AI エージェントのエントリ ポイント。プロジェクトの概要、ビルドとテストのコマンド、物事が存在する場所。それがなければ、すべてのセッションはゼロから始まります。
CODING_GUIDELINES.md — 実用的でオーバーエンジニアリングに反対する哲学に加え、コマンドまたはその動作が変更されるたびに docs/ 内のマンページを更新する必要があるという厳格なルールを追加します。ガイドラインにそうしなければならないと記載されているため、ドキュメントは最新の状態に保たれます。
TESTING_GUIDELINES.md — テストの命名規則と必須のコメント パターン (説明、番号付きの手順、期待される結果) により、すべてのテストが一目でわかるようになります。
GIT_TEST_SCENARIOS.md — 一時的な Git リポジトリの起動、オンデマンドでのマージとリベースの競合の生成、ローカル リモートの接続、および Git の状態の検証のためのパターン。ついにテストスキャフォールディングを退屈なものにしたドキュメント。
COMMIT_GUIDELINES.md — 範囲、50 文字の件名、命令文、および 72 でハードラップされた本文を備えた従来のコミット形式。「何を」と「なぜ」を説明し、「方法」は決して説明しません。
ARCHITECTURE.md — プロジェクトのマップ: cmd/ が所有するものと external/ が所有するもの、3 層のコマンド パターン、および各ビジネス ロジックが存在する場所。
CONFIGURATION.md — 完全な構成リファレンス。

3 レベルの優先順位ルール (ブランチ タイプ定義 → コマンド固有の設定 → CLI フラグ) と、どのキーがどの層に属するか。
DEV_WORKFLOW.md — すべてを結び付ける、構造化された課題から PR ワークフロー: 課題→ .ai/ 内のアーティファクトの計画→ 実装→ レビュー→ PR。スキルシステムが最終的にエンコードするドキュメント。
もちろん、これらのドキュメントはすべて、オープンソース リポジトリにアクセスすることで完全に分析できます。
最も重要なものは GIT_TEST_SCENARIOS.md であることが判明しました。 Git テスト シナリオのセットアップ (マージの競合、リモート、複数ステップの操作) は、まさに AI が常に微妙に間違える類のものです。明示的なガイダンスがなければ、すべてのテスト作成セッションは壊れたセットアップのデバッグに発展しました。
ガイドラインが整備されたことで、テスト足場は退屈なものになりました。それがテスト足場のあるべき姿です。
最も重要なアーキテクチャ上のルールは、AI が違反し続けたものでした。
3 層のコマンド パターン: Cobra ハンドラー → コマンド ラッパー → 関数実行
すべての Git 操作は、internal/git/repo.go を経由します。exec.Command を直接実行することはありません。
設定の優先順位: ブランチタイプのデフォルト → git config → CLI フラグ
特定の終了コードとコンテキスト メッセージを含むカスタム エラー タイプ
この時期のいくつかの逸話が私たちの心に残っています。
AI に構成リゾルバーの作成を依頼しました。ガイドラインでは設定を一度ロードしてパススルーすると言っているにもかかわらず、個々のオプションごとに Git コマンドを呼び出していました。これを修正した後、重複した TagOptions 構造体が作成され、ヘルパー関数が間違ったファイルに残されました。各セッションでは、同じ問題の新しい変種が導入されました。
私たちはかつて、Opus と Sonnet の両方に同じバグを検出するよう依頼しました。それを見つけたのはオーパスだけだった。
繰り返し発生する障害モード: コマンドが正しいディレクトリで実行されません。一度修正してみますが、

そして、後で別のコードで再び現れることになります。
Sonnet 3.7 がリリースされたとき、このコードベースでは以前のバージョンよりも動作が悪かった。モデルのアップグレードは厳密には単調ではありません。
それらはどれも個別に驚くべきものではありません。彼らが形成するパターンは次のとおりです。AI はセッション内で学習せず、セッションをまたいで学習することもありません。 「学習」とはドキュメンテーションエンジニアリングです。
フェーズ 3: 転換点 (2025 年 9 月)
1 週間で 40 以上のコミット。その月のログを見ると、新機能 (finish の包括的なマージ戦略制御)、新しいテスト (finish コマンドだけで約 8,000 行)、および新しいガイドライン (テスト、Git 操作、AI にコミット メッセージに AI 属性を追加しないよう明確に指示するものさえ) が混在しています。ドキュメントとコードは一緒に進化していました。新しいガイドラインはそれぞれ、将来の一連の間違いを防止しました。最後に v0.1.0 を出荷しました。
見落としがちな部分は、ガイドラインの複合です。あなたが作成したすべてのガイドラインは、プロジェクトの存続期間中、将来のすべてのセッションで効果を発揮します。バイブコーディングフェーズには何の困難もありませんでした。Cursor が作成したすべてのコマンドは新たなスタートでした。 9 月までに、各セッションは 3 か月分の蓄積されたコンテキストがすでに読み込まれた状態で始まりました。
フェーズ 4: モデルの追いつき (2025 年 10 月～11 月)
11 月のモデル アップデートは 2 番目のターニングポイントでしたが、予想どおりの理由ではありませんでした。モデルのコード生成が必ずしも大幅に向上したわけではありません。彼らは、既存のコードの理解とエージェントの動作、つまりコードベースを探索し、関連するコンテキストを見つけ、ファイル間で部分を接続することが劇的に向上しました。
ガイドラインが蓄積されているプロジェクトにとって、これは非常に重要でした。エージェントは、コードを生成するのではなく、コードを書く前に実際に規約を読み取って内部化できるようになりました。

暗闇の中でいます。オーパスのような思考モデルは、複雑なタスクを分解するのがはるかに優れていました。システムを構築するために私たちが行った作業は、異なる規模で成果を上げ始めました。
2025 年初頭の多くの摩擦 (複数回の試行が必要なリファクタリングや実行されないテスト) は、おそらく今日の最初の試行で機能するでしょう。当時の難しさの一部は、非常に複雑だったということです。モデルがまだ完成していないということもありました。
フェーズ 5: ワークフローの記入 (2025 年 12 月～2026 年 1 月)
システムが安定したことで、機能の作業が容易になりました。フック システム、パブリッシュおよびトラック コマンド、スカッシュ メッセージを追加し、 v0.2.0 および v0.3.0 を出荷しました。
新機能はより安定して導入され、ガイドラインはその役割を果たしました。
フェーズ 6: v1.0.0 とスキル システム (2026 年 2 月)
2 月に 1.0 を出荷しました。それに加えて、開発ライフサイクル全体をカバーする 14 のクロード コード スキルを構築しました。
このうち、/resolve-issue は最も重要な機能です。ワークフローの各フェーズに対して順次サブエージェントを生成します。つまり、GitHub の問題が入力され、すぐにレビューできる PR が出力されます。ここで、ガイドライン システムは一連の文書であることをやめ、実行可能なパイプラインになります。
それは、論文が最終的に現金化された瞬間でもあります。 1年間蓄積されたガイドラインがスキルとしてコード化され、単一のコマンドになりました。
フェーズ 7: AI コード レビューの実際の様子 (2026 年 2 月～4 月)
クロードにプル リクエストを適切にレビューしてもらうのは、驚くほど困難でした。多くのチームは AI に PR をスパムさせるだけです。私たちは、明確で安定した概要形式を備えた、徹底的で的を絞ったものを求めていました。
いくつかの問題がすぐに明らかになりました。
矛盾。まったく同じ差分でまったく同じレビューを再実行すると、3 つのまったく異なるレビューが得られました。異なる所見、異なる重症度、異なる結論。
フォーマットドリフト。の

モデルが指定された出力形式に固執しない可能性があります。追加のセクションが表示され、見出しスタイルが変更されました。
ノイズ。初期のバージョンでは、すべてにフラグが付けられていました。実際のバグと並行して、細かい問題も報告されていました。
特定のガイドライン ファイルを参照するレビュー基準ドキュメント
明示的な出力形式仕様
重要な問題とあれば便利な問題を区別するための重大度レベル
6 つの明確に定義されたレビュー領域: テスト カバレッジ、コーディング ガイドライン、コード品質、セキュリティ、ドキュメント、コミット メッセージ
レビューを権威ある評決ではなく、有益なシグナルとして扱う
また、.github/instructions/ — ファイル パターン (コマンド、テスト、内部パッケージ) を対象とした Copilot レビュー手順も追加しました。重要なのは、AI レビュー担当者に単一の汎用プロンプトではなく、ファイル タイプごとに特定のコンテキストを提供することです。
これらの改良を加えた v1.1.0 を 4 月に出荷しました。
一般的な機能または修正のワークフローは次のとおりです。
問題は GitHub 上で作成されます (メモによると AI によって作成される場合もあります)。
/resolve-issue 42 によりパイプラインが開始されます。
サブエージェントは問題を分析し、コードベースを調査し、メモを .ai/ に書き込みます。
もう 1 つは、git-flow-next 自体を使用して機能ブランチを作成します。
もう 1 つは、特定のファイル パスを使用して実装計画を生成します。
もう 1 つは、ガイドラインに従って実装、テストを実行し、コミットします。
PR は構造化された概要とともに開きます。
自動レビューにより、プロジェクトのガイドラインに照らして変更がチェックされます。
結果を確認し、必要に応じてコメントを残します。
/アドレス

[切り捨てられた]

## Original Extract

The honest 13-month journey of building git-flow-next with AI — what worked, what failed, and why the guidelines turned out to be the real product.

How We Shipped git-flow-next 1.0 Almost Entirely with AI | git-flow-next
Docs
Phase 1: Vibe Coding with Cursor (March 2025)
Phase 2: Building the System (May–August 2025)
Phase 3: The Turning Point (September 2025)
Phase 4: The Models Catch Up (October–November 2025)
Phase 5: Filling Out the Workflow (December 2025–January 2026)
Phase 6: v1.0.0 and the Skill System (February 2026)
Phase 7: What AI Code Review Actually Looks Like (February–April 2026)
How We Shipped git-flow-next 1.0 Almost Entirely with AI
Written by: Alexander Rinass Published on June 9, 2026
In our 1.0 announcement , we mentioned that git-flow-next has also been an experiment “to see how far we could push today’s AI capabilities on an open-source project of this scope.”
We promised more on that journey in a future post. This is it.
Approximately 29,000 lines of Go code
We did all of this on the side, alongside our regular work — which is worth keeping in mind before reading too much into the timeline.
Most of the code was written by AI. While we’re experienced developers, we’re not proficient in Go. So we provided the architecture, specifications, and guidance; the AI was the hands on the keyboard.
Two things motivated the project:
We wanted an open-source tool that would need little ongoing maintenance from us.
We wanted to experiment with something that’s more or less a black box to us — we’re not familiar with the tech stack — to see how building with AI holds up under those conditions, both for us specifically and as a more general question.
If you’re in a hurry and are just looking for the tl;dr, here is the one lesson we’d carry to the next project: AI development isn’t “telling AI what to build.” It’s “building the system that tells AI what and how to build.”
The code is the output. The guidelines, skills, review criteria, and architectural docs are the actual product.
Phase 1: Vibe Coding with Cursor (March 2025)
The first phase was all about getting our feet wet: minimal instructions, all in one file, just to explore what AI could accomplish.
The pace was remarkable. In about two weeks, Cursor produced all of git-flow-next’s core commands — start , finish , list , overview , update , delete , rename , checkout — at roughly one new command every day or two.
Up close, the project wasn’t really a project. It was a collection of individually-generated files that happened to share a directory. There was no clear separation between the CLI layer, business logic, and Git operations — every command did things its own way. Patterns were duplicated across files in slightly different shapes, often subtly enough that asking AI to find duplicates would surface one and quietly miss the others. Git configuration was loaded on every operation instead of once at startup — performance regressions by design, not by accident.
We tried to refactor our way out of it. With Cursor, with Claude, with detailed instructions. It didn’t work. Functions were duplicated, code didn’t compile, proposed refactorings were over-engineered with the wrong abstractions.
We even tried full rewrites from scratch. Those failed, too.
What eventually worked was the opposite of speed: tiny incremental steps. Create empty command files first. Add one piece of logic. Run the tests. Move on. And more importantly, we realized we had to define a development philosophy ourselves — explicitly, in writing — before AI could help us clean anything up.
You can’t evaluate an AI’s proposal if you don’t already know what you want. That was the first time we understood the shape of the real problem.
Phase 2: Building the System (May–August 2025)
On day one with Claude Code, we let it analyze the project and write a CLAUDE.md file. It got some things wrong — inconsistent config keys, for one — but accuracy wasn’t really the point.
The point was to give AI a place to start understanding the project before writing code.
That set off a pattern. Every time AI failed in a recognizable way, the answer was a new guideline. Over a few months, eight of them accumulated:
CLAUDE.md — The entry point for any AI agent working on this codebase. Project overview, build and test commands, where things live. Without it, every session starts from zero.
CODING_GUIDELINES.md — A pragmatic, anti-over-engineering philosophy, plus a hard rule that the manpages in docs/ must be updated whenever a command or its behavior changes. Documentation stays current because the guideline says it must.
TESTING_GUIDELINES.md — Test naming conventions and a mandatory comment pattern — description, numbered steps, expected outcomes — so every test is self-explanatory at a glance.
GIT_TEST_SCENARIOS.md — Patterns for spinning up temporary Git repos, producing merge and rebase conflicts on demand, wiring up local remotes, and verifying Git state. The doc that finally made test scaffolding boring.
COMMIT_GUIDELINES.md — Conventional commit format with scopes, 50-character subject lines, imperative mood, and bodies hard-wrapped at 72 — explaining the “what” and the “why,” never the “how.”
ARCHITECTURE.md — The map of the project: what cmd/ owns versus internal/ , the three-layer command pattern, and where each piece of business logic lives.
CONFIGURATION.md — The complete config reference, including the three-level precedence rule (branch type definitions → command-specific config → CLI flags) and which keys belong at which layer.
DEV_WORKFLOW.md — The structured issue-to-PR workflow that ties everything together: issue → planning artifacts in .ai/ → implementation → review → PR. The doc the skill system ultimately encodes.
All of these documents can be analyzed in full by visiting our open-source repository , of course!
The single most important one turned out to be GIT_TEST_SCENARIOS.md . Setting up Git test scenarios — merge conflicts, remotes, multi-step operations — is exactly the kind of thing AI consistently gets subtly wrong. Without explicit guidance, every test-writing session devolved into debugging broken setups.
With the guideline in place, test scaffolding became boring, which is what test scaffolding should be.
The architectural rules that mattered most were the ones AI kept violating:
A three-layer command pattern: Cobra handler → command wrapper → execute function
All Git operations go through internal/git/repo.go — never direct exec.Command
Configuration precedence: branch type defaults → git config → CLI flags
Custom error types with specific exit codes and contextual messages
A few anecdotes from this period stick with us:
We asked AI to create a config resolver. Despite the guideline saying load config once and pass it through, it called a Git command for every individual option. After we corrected that, it created a duplicate TagOptions struct, then left a helper function in the wrong file. Each session introduced a new variant of the same problem.
We once asked both Opus and Sonnet to detect the same bug. Only Opus found it.
A recurring failure mode: commands not executed in the right directory. We’d fix it once, and it would crop up again later in different code.
Sonnet 3.7, when it came out, worked worse than its predecessor on this codebase. Model upgrades aren’t strictly monotonic.
None of those are individually surprising. The pattern they form is: AI doesn’t learn within a session, and it certainly doesn’t learn across sessions. “Learning” is documentation engineering.
Phase 3: The Turning Point (September 2025)
40+ commits in a single week. Looking at the log from that month, it’s a mix of new features (comprehensive merge strategy control for finish ), new tests (around 8,000 lines for the finish command alone), and new guidelines (testing, Git operations, even one specifically telling AI not to add AI attribution in commit messages). Documentation and code were evolving together. Each new guideline prevented a class of future mistakes. We shipped v0.1.0 at the end of it.
Here’s the part that’s easy to miss: guidelines compound. Every guideline you write pays off in every future session for the lifetime of the project. The vibe-coding phase had no compounding — every command Cursor wrote was a fresh start. By September, every session began with three months of accumulated context already loaded.
Phase 4: The Models Catch Up (October–November 2025)
The November model updates were the second turning point, and not quite for the reason you’d expect. The models didn’t necessarily get much better at producing code. They got dramatically better at understanding existing code and at agentic behavior — exploring a codebase, finding relevant context, connecting pieces across files.
For a project with accumulated guidelines, this mattered enormously. The agent could now actually read and internalize the conventions before writing code, rather than generating in the dark. Thinking models like Opus were much better at breaking down complex tasks. The work we’d done building the system started paying out at a different scale.
A lot of the friction from early 2025 — refactors that needed multiple attempts, tests that wouldn’t run — would probably work first try today. Some of the difficulty back then was real complexity. Some of it was just that the models weren’t quite there yet.
Phase 5: Filling Out the Workflow (December 2025–January 2026)
With the system stable, feature work got easier. We added the hooks system, the publish and track commands, squash messages, and shipped v0.2.0 and v0.3.0 .
New features landed more consistently — the guidelines were doing their job.
Phase 6: v1.0.0 and the Skill System (February 2026)
We shipped 1.0 in February. Alongside it, we built 14 Claude Code skills covering the full development lifecycle:
Of these, /resolve-issue is the crown jewel. It spawns sequential subagents for each phase of the workflow: a GitHub issue goes in, and a ready-to-review PR comes out. This is where the guideline system stops being a set of documents and becomes an executable pipeline.
It’s also the moment the thesis finally cashed out. A year of accumulated guidelines, encoded as skills, became a single command.
Phase 7: What AI Code Review Actually Looks Like (February–April 2026)
Getting Claude to review pull requests well was surprisingly hard. A lot of teams just let AI spam the PR. We wanted something thorough and well-targeted, with a clean, stable summary format.
A few problems showed up immediately:
Inconsistency. Re-running the exact same review on the exact same diff gave three very different reviews. Different findings, different severity, different conclusions.
Format drift. The model wouldn’t stick to the specified output format. Extra sections appeared, heading styles changed.
Noise. Early versions flagged everything — style nits alongside actual bugs.
A review criteria doc referencing specific guideline files
An explicit output format spec
Severity levels to separate critical issues from nice-to-haves
Six clearly defined review areas: test coverage, coding guidelines, code quality, security, documentation, commit messages
Treating the review as a useful signal, not an authoritative verdict
We also added .github/instructions/ — Copilot review instructions scoped to file patterns (commands, tests, internal packages). The point is to give AI reviewers specific context per file type, rather than a single generic prompt.
We shipped v1.1.0 in April with those refinements.
For a typical feature or fix, the workflow is:
An issue gets created on GitHub (sometimes by AI, from our notes).
/resolve-issue 42 kicks off the pipeline.
A subagent analyzes the issue, explores the codebase, writes notes to .ai/ .
Another creates a feature branch — using git-flow-next itself.
Another generates an implementation plan with specific file paths.
Another implements, runs tests, and commits per guidelines.
A PR is opened with a structured summary.
Automated review checks the changes against the project guidelines.
We review the result and leave comments if needed.
/addres

[truncated]
