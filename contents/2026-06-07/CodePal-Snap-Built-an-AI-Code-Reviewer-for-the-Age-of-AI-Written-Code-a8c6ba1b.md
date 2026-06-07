---
source: "https://eng.snap.com/codepal"
hn_url: "https://news.ycombinator.com/item?id=48437093"
title: "CodePal: Snap Built an AI Code Reviewer for the Age of AI-Written Code"
article_title: "CodePal: How Snap Built an AI Code Reviewer for the Age of AI-Written Code"
author: "Kaedon"
captured_at: "2026-06-07T18:33:10Z"
capture_tool: "hn-digest"
hn_id: 48437093
score: 1
comments: 0
posted_at: "2026-06-07T17:50:10Z"
tags:
  - hacker-news
  - translated
---

# CodePal: Snap Built an AI Code Reviewer for the Age of AI-Written Code

- HN: [48437093](https://news.ycombinator.com/item?id=48437093)
- Source: [eng.snap.com](https://eng.snap.com/codepal)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T17:50:10Z

## Translation

タイトル: CodePal: Snap が AI 記述コードの時代に向けた AI コード レビューアーを構築
記事のタイトル: CodePal: Snap が AI 記述コードの時代に向けた AI コードレビューアをどのように構築したか

記事本文:
エンジニアリング ブログ ブログ 私たちの価値観 採用情報 CodePal: Snap が AI 記述コードの時代に向けた AI コード レビューアーをどのように構築したか
AI コーディング ツールは、ソフトウェアの作成方法を根本的に変えました。 Snap では、エンジニアが Cursor や Claude Code などのツールを毎日使用しており、速度への影響は現実的です。
年初から現在まで、マージされた (プル リクエスト) PR 率は 60% 増加しています。
しかし、落とし穴があります。誰もがコードをより速く作成できるようになると、ボトルネックは消えるのではなく、変化します。そしてSnapでは、コードレビュープロセスに真っ向から取り組みました。
これらすべての追加レビューでは、より多くの PR が増え、PR がより大きくなり、レビュー担当者が薄くなっていきました。その結果、レビューキューが滞り、マージ時間が長くなりました。コードレビューを迅速化せずにコードを迅速に作成すると、技術的負債が発生します。
そこで私たちは、Snap の内部 AI を活用したコード レビュー アシスタントである CodePal を構築しました。私たちの目標は、Snap での作業方法を正確に理解し、チームメイトの PR レビューの負担を軽減するために、PR 作成者に自分の仕事に関する実際的で価値のある迅速なフィードバックを提供する、高度にインテリジェントな AI コード レビュー システムを構築することでした。私たちのビジョンは、ソフトウェア開発ライフサイクルの一環としてエンジニアが一貫して協力するコードレビュー担当者です。
現在、CodePal は人間がレビューする前に、Snap ですべての PR の 90% をレビューしています。
社内で構築することを決定する前に、いくつかのベンダーのソリューションを評価しました。 AI コード レビュー ツールの状況はますます混雑しており、いずれも魅力的な機能を提供しています。しかし、Snap のニーズに対しては、2 つの重要な領域で不十分でした。
まず、統合の深さ: Snap のエンジニアリング インフラストラクチャは広範囲にわたり、深く相互接続されています。当社のビルド システム、デプロイ パイプライン、内部ツールにはすべて、既製のツールでは簡単に対応できない特定の要件があります。 CodePal は S を理解する必要がありました

ナップ固有のプロト定義、内部構成システム、そして何百ものサービスにわたる作業の微妙な違い。非常に具体的なチームレベルのカスタマイズが必要でしたが、ほとんどのツールではその一部は提供されていましたが、必要なすべては提供されませんでした。いくつかの設定ノブしか備えていない汎用ツールは、私たちにとっては機能しませんでした。
2 番目にスピード: 動作するエンドツーエンドのデモを 2 週間で出荷しました。ベンダーの調達サイクルもまだ終わっていませんでした。そして第 1 週目から、CodePal は独自のプル リクエストをレビューし、人間によるレビューよりも早く問題を明らかにしました。
CodePal の核心は、他のコード レビューアと同様にコード diff から始まりますが、リポジトリ全体にわたる深いシンボリック コンテキストを持ち、リポジトリ全体でレビューのシェアが増加しています。それは表面レベルのチェックをはるかに超えています。パート 2 では、クロスリポジトリ機能について詳しく説明します。
リポジトリ間の依存関係追跡 ますます多くのレビューに対応するため、CodePal は、Snap の完全なコードベースをカバーする内部セマンティック検索システムであるコード検索を通じて、リポジトリの境界を越えてアクセスすることもできます。関数シグネチャの変更がマージされる前に、CodePal は、PR がアクセスするリポジトリとは異なるリポジトリに存在する場合でも、壊れる可能性がある下流の呼び出し元を特定できます。これは、マージ前に実際のリポジトリ間の破損をすでに検出しており、差分だけを見てレビュー担当者がフラグを立てることができなかったバグです。パート 2 では、Code Search を CodePal にどのように統合したかなど、Code Search について詳しく説明します。
CodePal は、コンパイル スイートやテスト スイートでは通常検出されない問題 (論理エラー、null ポインターのリスク、競合状態、リソース リーク、エラー処理のギャップ、型の不一致、エッジ ケース、状態管理の問題など) を探します。一部のリポジトリでは、レビューを exa に合わせて調整できるように、そのリポジトリに固有の追加チェックが適用されます。

必要なものを正確に。
エンジニアは CodePal の調査結果を一貫して検証しており、CodePal によって発見された受け入れられたバグ (PR レビューで +1 票を獲得したバグ) の大部分は重大度または高にランク付けされます。これらの問題は、CodePal がなければ、本番環境に到達しないように人間のレビュー担当者が発見する必要があったでしょう。また、開発中、リリース前にパフォーマンスに関連する潜在的な発見を特定することもできました。
CodePal は、行ごとの分析を超えて、PR が何を行うのか、なぜ行うのかを説明するセマンティック diff 概要を生成し、レビュー担当者がコードに入る前に概要を理解できるようにします。また、PR の説明、タイトル、リリース ノートや移行手順などの構造化されたセクションを自動生成することもでき、これらはすべてコード変更から直接推測されます。
これらの機能により、著者と査読者の両方の時間が大幅に節約されます。レビュー担当者は、レビューの最初の 10 分を PR が何をしているのかを理解することに費やす代わりに、アプローチが正しいかどうかの評価にすぐに取り掛かることができます。これらの出力、レビュー結果、セマンティック概要、生成された説明はすべて、以下で説明する同じコンテキスト ビルドから流れます。
単独の差分は、ほとんどの場合、推論するには間違った単位です。 return user.Profile.Email を追加する行は、このコード パスで Profile が null 可能であることがわかるまでは問題ないように見えます。これは、diff が決して触れない 3 つのファイルを読み取ることによってのみわかります。ほとんどの AI レビュー ツールは、これを無視して実際のバグを見落とすか、リポジトリ全体をプロンプトに詰め込んで役に立たないトークンの料金を支払うかのどちらかです。 CodePal はどちらも行いません。
CodePal は、2 パスのアプローチを使用してこれを解決します。最初のパスでは、tree-sitter を使用してリポジトリを解析し、シンボルからファイルへのインデックスを構築します。 2 番目のパスでは diff を分析し、参照されているシンボルを抽出し、t 内でそれらを検索します。

ファイルにインデックスを付け、シンボルの重複によってスコアを付けます。このスコアリングにより、トークン バジェット内で最も関連性の高い上位 N 個のファイルが選択されます。 CodePal は何が変更されたかを確認するだけではありません。まずコードを理解します。
クローンなし、作業コピーなし CodePal がコードで何を行うかを説明する前に、リポジトリのクローンを作成するという、CodePal が実行しないことに注意してください。すべてのレビューは、GitHub Enterprise (GHE) API に対して完全にメモリ内で実行されます。作業コピーをチェックアウトしたり、ソースをディスクに書き込んだり、リポジトリの長期ミラーを保持したりすることはありません。これは意図的な制約であり、制限ではありません。 Snap の最大のリポジトリのクローン作成は時間がかかり、時間とストレージの両方にコストがかかります。 CodePal がレビューを実行している速度では、レビューごとにクローンを作成するアーキテクチャでは、インフラストラクチャか GHE インスタンスのどちらかが溶けてしまうでしょう。代わりに、git ツリーの差分を利用して変更内容を特定し、必要な特定の BLOB のみをフェッチし、2 パス ファイル ピッカーによってワーキング セットを小さく保ちます。一般的なレビューでは、リポジトリのサイズに関係なく、少数の API 呼び出しと数百 KB のソースが扱われます。それがなければ、残りのどれも拡張できません。
このアプローチにより、アーキテクチャ上の利点が得られます。 1 つの親ワークフローは、高価なシンボルのインデックス作成とファイル選択のプロセスを 1 回だけ実行し、結果のコンテキストを共有ストアに書き込みます。コード レビュー、概要生成、説明生成を処理する 3 つの子ワークフローは、すべてこの単一の共有コンテキストから読み取ります。製品の概要と説明は事実上無料です。彼らはコンテキストではなくストレージに対してお金を払っているのです。コード検索が同じシステムに接続されるため、この共有ストアの設計はパート 2 でも重要になります。
AI コード レビューを生成する素朴な方法は、差分とコンテキストをモデルに送信し、結果を求めることです。私たちはそれを試してみました。それはうまくいきます、まあ

f.このモデルは実際のバグを見つけますが、でっちあげたり、実行間で矛盾したり、ランダム シードに応じて大きく異なる結果を生成したりすることもあります。単一のモデル呼び出しよりも徹底的で一貫性のあるレビュー ループが必要でした。
この解決策は、CodePal を斬新なものにする中心となる、レビュー ループと呼ばれる多段階の反復プロセスです。
ブートストラップ同時パス 2 つのパスが最初から並行して実行されます。これらは同じモデルを使用しますが、異なるサンプリング パラメーターを使用します。モデルの主な役割はバグを見つけることです。 2 回実行して結果を比較すると、モデルが実際にどの結果を信じているかがわかります。
1 つのブートストラップ パスが終了した瞬間に、3 番目のパスがキャンセル可能なコンテキストに起動され、作業が破棄される可能性があります。両方のブートストラップ パスが完了すると、スーパーバイザは結果を比較します。彼らが同意した場合、パス 3 の作業は直ちに破棄されます。彼らの意見が一致しない場合は、これまでに行った作業が最終的な調査結果に反映されます。
パス 3 以降では、スーパーバイザーがこれまでに見たことのないものをパスで見つけた瞬間、現在のパスが終了するのを待たずに次のパスが開始されます。全く新しい発見をもたらさないパスには、単に後継者が得られません。このメカニズムにより、ループが経済的に収束し、検索対象が多い場合にはより多くのパスをトリガーし、差分がすでに枯渇している場合にはより少ないパスをトリガーします。
このレビュー ループは、 Verifier によって強化されます。 Verifier は、結果をマージするときにそれらを消費し、最終的なフィルターとして機能する、別個の長期実行モデルの会話です。検証者は、初期のシングルパスアプローチを悩ませていた自己矛盾と幻覚を捉えます。たとえば、調査結果で引用されているすべてのシンボルが提供されたコンテキストに実際に存在することをチェックし、調査結果が以下に基づいて投稿されるリスクを大幅に軽減します。

模範的な幻覚。
全体として、このループは、単一パスよりも徹底的で、単一のモデル呼び出しよりも実行全体で一貫性のあるレビューを生成します。より多くのコンピューティングコストがかかりますが、合意ゲートにより、利益のない投機的な作業のほとんどが無効になり、収束ルールにより、すでに使い果たされた差分で後続のパスが実行されなくなります。
1 四半期でゼロから 90% へ
CodePal はオプトイン実験として開始され、潜在的なバグやセキュリティの脆弱性について PR を分析しました。導入は PR の 9% で慎重に始まりましたが、一貫した徹底したレビューが不可欠である少数のリポジトリでレビューを必須としたため、急速に増加しました。リポジトリ所有者がオプトインを開始すると、CodePal は 300 のリポジトリ全体で自発的に使用され、70% 以上の肯定的な感情が得られました。 CodePal で検出できるバグがすり抜けているのを見て、エンジニアはこれをデフォルトにするよう求め始めました。私たちがチームの自動オプトインを開始したとき、重要な点は次のとおりでした。人々は私たちにオフに戻すよう求めているのではなく、実際にオンにしたいと思っていたのです。 1 四半期以内に、AI レビューを受けた PR が事実上ゼロから、Snap の全 PR の 90% 以上が CodePal レビューを受けるようになりました。重要なのは、品質が導入に合わせて維持されていることです。同じ四半期に、再現率は 30% から 80% に上昇し、バグ発見に対するエンジニアのセンチメントは 80% に達しました。
導入率 90% に到達することが最初のマイルストーンでしたが、導入だけではツールが信頼できるものになるわけではありません。工具が見えるようになり、亀裂が露出します。
3 つの欠陥が目立っていました。レビューは実行間で一貫性がなかったこと、CodePal はコードが行うべきことについて間違ったことを仮定することで意図を幻覚することがあったこと、そして変更をロールアウトする前に検証するプロセスがなかったことです。私たちはユーザーからのフィードバックを頼りにしました。

遅すぎる。
私たちは、CodePal のすべての新しい変更を AB テストするプロセスだけでなく、どこがうまくいっているのか、どこがうまくいっていないのかを明確に理解できるように評価フレームワークを構築しました。実際のエンジニアからのフィードバックからグラウンド トゥルース データセットを作成し、真陽性 (レビュー全体で一貫して発見できる実際のバグ) の再現率を向上させ、誤検知 (表面上の実際のバグのみ) を最小限に抑えるという目標を設定しました。スピードとコストはガードレールとして残りましたが、主な目的は信頼でした。
4 つの投資によりリコールの改善が推進されました。
より深いコード理解: CodePal はクロスファイル シンボル解決と拡張言語解析を使用するようになり、推測するのではなく意図を理解します。
マルチパス レビュー ループ (上記の「レビュー ループ」セクションで説明)。
検出範囲の拡大: 範囲が 8 つのバグ カテゴリから 12 に拡張されました。
増分レビュー: 新しいコミットごとに、差分に残されたファイルの検出結果を自動解決する集中的な再レビューがトリガーされます。
この取り組みの影響は非常に期待できます。真陽性の再現率は 30% から 80% に増加し、ゴールデン データセット内の偽陽性率は 0% に低下しました (ライブトラフィックではなく、保持されたゴールデン データセットで測定)

[切り捨てられた]

## Original Extract

Engineering Blog Blog Our Values Careers CodePal: How Snap Built an AI Code Reviewer for the Age of AI-Written Code
AI coding tools have fundamentally changed how software gets written. At Snap, engineers use tools like Cursor and Claude Code daily, and the impact on velocity has been real.
Year-to-date, our merged (Pull Request) PR rate is up 60%.
But there's a catch. When everyone can produce code faster, the bottleneck doesn't disappear, it shifts. And at Snap, it moved squarely onto the code review process.
We were seeing more PRs, larger PRs, and reviewers stretched thinner across all of these additional reviews. As a result, review queues backed up, and merge times grew. Faster code creation without faster code review is a recipe for technical debt.
So we built CodePal , Snap's internal AI-powered code review assistant. Our goal was to build a highly intelligent, AI code review system which understands exactly how we work at Snap, giving PR authors real, valuable, quick feedback on their work, in order to reduce the PR review burden on their teammates. Our vision is a code reviewer that engineers consistently collaborate with as part of the software development lifecycle.
Today, CodePal reviews 90% of all PRs at Snap, even before a human reviews them.
We evaluated several vendor solutions before deciding to build in-house. The landscape of AI code review tools is increasingly crowded and all offer compelling capabilities. But for Snap's needs, they fell short in two critical areas.
First, integration depth: Snap's engineering infrastructure is extensive and deeply interconnected. Our build systems, deployment pipelines and internal tooling all have specific requirements that off-the-shelf tools can't easily accommodate. CodePal needed to understand Snap-specific proto definitions, our internal configuration systems, and the nuances of working across hundreds of services. We needed very specific team-level customizations that most tools offered a portion of, but didn’t offer everything we needed. A generic tool with only a few configuration knobs wasn't going to work for us.
Second, speed: We shipped a working end-to-end demo in two weeks. A vendor procurement cycle hadn't even finished. And from week one, CodePal reviewed its own pull requests, which surfaced issues faster than any human review could.
At its core, CodePal starts with the code diff like any other code reviewer but now it has deep symbolic context across the repository, and for a growing share of reviews, across repositories. It goes well beyond surface-level checks. Part 2 covers the cross-repo capability in depth.
Cross-Repository Dependency Tracking For a growing portion of reviews, CodePal also reaches across repository boundaries via Code Search, an internal semantic search system covering Snap's full codebase. Before a function signature change merges, CodePal can identify the downstream callers that would break, even when they live in a different repo than the one the PR touches. This has already caught real cross-repo breakages before they merged, bugs no reviewer looking only at the diff could have flagged. Part 2 covers Code Search in depth, including how we integrated it into CodePal.
CodePal looks for issues that aren't typically caught by compilation or test suites: logic errors, null pointer risks, race conditions, resource leaks, error handling gaps, type mismatches, edge cases, and state management problems. For some repositories, it applies additional checks that are specific to that repo so that reviews can be tailored towards exactly what it needs.
Engineers consistently validate CodePal’s findings, and the majority of accepted bugs (bugs with a +1 vote in the PR review) found by CodePal rank as Critical or High severity. These are issues that, without CodePal, would have needed to have been caught by the human reviewer to ensure they don’t reach production. It has also allowed us to identify potential performance-related findings during development, prior to release.
Beyond line-by-line analysis, CodePal generates semantic diff summaries that explain what a PR does and why , giving reviewers a high-level understanding before they dive into the code. It can also auto-generate PR descriptions, titles, and structured sections like release notes or migration steps, all inferred directly from the code changes.
These features save significant time for both authors and reviewers. Instead of spending the first ten minutes of a review figuring out what a PR is doing, reviewers can jump straight to evaluating whether the approach is correct. These outputs, review findings, semantic summary, generated description, all flow from the same context build, which we'll cover below.
A diff in isolation is almost always the wrong unit to reason about. A line that adds return user.Profile.Email looks fine until you know that Profile is nullable on this code path, which you can only learn by reading three files the diff never touches. Most AI review tools either ignore this, and miss real bugs, or stuff the entire repository into the prompt and pay for tokens that don't help. CodePal does neither.
CodePal solves this using a two-pass approach. The first pass parses the repository with tree-sitter to build a symbol-to-file index. The second pass analyzes the diff, extracts referenced symbols, looks them up in the index, and scores files by symbol overlap. This scoring selects the top N most relevant files within the token budget. CodePal doesn't just see what changed; it understands the code first.
No clone, no working copy Before describing what CodePal does with the code, it's worth noting what it doesn't do: clone the repository. Every review runs entirely in memory, against the GitHub Enterprise (GHE) API. We never check out a working copy, never write source to disk, and never hold a long-lived mirror of any repository. This was a deliberate constraint, not a limitation. Cloning Snap's largest repos is slow and expensive in both time and storage. At the rate CodePal runs reviews, a clone-per-review architecture would have melted either our infrastructure or our GHE instance. Instead, we lean on git tree diffs to identify what changed, fetch only the specific blobs we need, and let the two-pass file picker keep the working set small. A typical review touches a handful of API calls and a few hundred KB of source, regardless of repo size. Without that, none of the rest of this would scale.
This approach provides an architectural payoff. One parent workflow executes the expensive symbol indexing and file selection process just once, writing the resulting context to a shared store. Three child workflows, which handle code review, summary generation, and description generation, then all read from this single, shared context. The summary and description products effectively come for free; they're paying for storage, not for context. This shared store design will also be critical for Part 2, as Code Search will plug into the same system.
The naive way to generate an AI code review is to send the diff and context to a model and ask it for findings. We tried that. It works, sort of. The model finds real bugs, but it also makes things up, contradicts itself between runs, and produces wildly different findings depending on the random seed. We needed a review loop that was both more thorough and more consistent than a single model call.
The solution is a multi-stage, iterative process we call the Review Loop, which is central to what makes CodePal novel.
Bootstrap concurrent passes Two passes run in parallel from the start. They use the same model but different sampling parameters. The model's primary role is to find bugs. Running it twice and comparing the results tells us which findings the model actually believes.
The moment one bootstrap pass finishes, a third pass launches into a cancellable context, work it might throw away. When both bootstrap passes complete, the supervisor compares their findings. If they agree, pass three's work is immediately discarded. If they disagree, the work it has already done counts toward the final findings.
From pass three onward, the moment a pass finds something the supervisor has not seen before, the next pass launches without waiting for the current one to finish. A pass that produces no net-new findings simply does not get a successor. This mechanism allows the loop to converge economically, triggering more passes when there is more to find and fewer when the diff is already exhausted.
This review loop is augmented by the Verifier , a separate long-running model conversation that consumes findings as they merge and acts as a final filter. The verifier catches the self-contradiction and hallucination that plagued our early single-pass approach. For instance, it checks that all symbols cited in a finding are actually present in the provided context, greatly reducing the risk of a finding being posted based on a model hallucination.
In aggregate, this loop produces reviews that are more thorough than a single pass and more consistent across runs than any single model call could be. It costs more compute, but the agreement gate kills most of the speculative work that doesn't pay off, and the convergence rule keeps successor passes from running on diffs that are already exhausted.
From Zero to 90% in One Quarter
CodePal launched as an opt-in experiment, analyzing PRs for potential bugs and security vulnerabilities. Adoption began cautiously at 9% of PRs, but quickly grew as we made reviews mandatory in a handful of repositories, where consistent, thorough reviews were essential. As repository owners started opting in, CodePal was being used voluntarily across 300 repositories with >70% positive sentiment. After seeing bugs slip through that CodePal would have caught, engineers began asking for it to be the default. When we started to auto opt-in teams, the key takeaway was this: people weren’t asking us to turn it back off, they actually wanted it on. Within a single quarter, we went from virtually no AI-reviewed PRs to over 90% of all PRs at Snap having a CodePal review . Importantly, quality kept pace with adoption. During that same quarter, our recall rate climbed from 30% to 80%, and engineer sentiment on bug findings reached 80% positive.
Getting to 90% adoption was the first milestone, but adoption alone doesn't make a tool trustworthy; it makes the tool visible, and visibility exposes the cracks.
Three flaws stood out: reviews were inconsistent across runs, CodePal would sometimes hallucinate intent by assuming the wrong thing about what the code was supposed to do, and we had no process to validate changes before they rolled out. We relied on user feedback, which was too slow.
We built an evaluation framework so we could clearly understand where we were doing well, where we were doing poorly, as well as a process to AB test all new CodePal changes. We formed a ground truth dataset from real engineer feedback and set goals to improve the recall of true positives (real bugs that are able to be found consistently across reviews) and minimizing the false positives (only surface real bugs). Speed and cost stayed as guardrails, but the primary objective was trust.
Four investments drove the recall improvements:
Deeper code understanding: CodePal now uses cross-file symbol resolution and expanded language parsing, so it understands intent rather than guesses it.
The multi-pass review loop, described in The Review Loop section above.
Broader detection scope: The scope expanded from 8 bug categories to 12.
Incremental reviews: Each new commit triggers a focused re-review with auto-resolution of findings whose files have left the diff.
The impact of this work has been very promising. Our recall rate of true positives has increased from 30% to 80%, the rate of false positives in our golden dataset has dropped to 0% (measured on the held-out golden dataset, not on live traffic

[truncated]
