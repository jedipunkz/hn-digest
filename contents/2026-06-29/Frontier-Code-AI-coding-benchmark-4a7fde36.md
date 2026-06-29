---
source: "https://cognition.com/blog/frontier-code"
hn_url: "https://news.ycombinator.com/item?id=48717839"
title: "Frontier Code (AI coding benchmark)"
article_title: "Introducing FrontierCode | Cognition"
author: "nicoburns"
captured_at: "2026-06-29T11:38:26Z"
capture_tool: "hn-digest"
hn_id: 48717839
score: 1
comments: 0
posted_at: "2026-06-29T11:26:18Z"
tags:
  - hacker-news
  - translated
---

# Frontier Code (AI coding benchmark)

- HN: [48717839](https://news.ycombinator.com/item?id=48717839)
- Source: [cognition.com](https://cognition.com/blog/frontier-code)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T11:26:18Z

## Translation

タイトル: Frontier Code (AIコーディングベンチマーク)
記事のタイトル: FrontierCode の紹介 |認知
説明: 今日のコーディング ベンチマークは、モデルが正しいコードを記述できることを証明していますが、私たちが本当に問うべき質問は、「モデルは実際に適切なコードを書けるのか?」ということです。

記事本文:
FrontierCode のご紹介 |認知メニュー ホームを閉じる
正確さから品質へのハードルを上げる #
今日のコーディング ベンチマークでは、モデルが正しいコードを記述できることが確立されています。しかし、AI が生成したコードが本番環境への主要なパスになるにつれて、正確さが命取りになっています。私たちが問うべき質問は、モデルは実際に良いコードを書けるのかということです。
モデルが高品質な実稼働コードベースの基準をどの程度満たしているかを測定するベンチマークである FrontierCode を導入できることを嬉しく思います。当社の特徴:
メンテナは実際にこの PR をマージしますか?私たちはコードのマージ可能性を測定する最初のベンチマークです。私たちの基準は、正確性、テスト品質、スコープ規律、スタイル、コードベース標準への準拠など、エンドツーエンドのコード品質を評価します。これには、単体テスト、ルーブリック、新しいタイプの検証ツールなど、一連の新しい採点手法が採用されています。
オープンソースのメンテナーによって作成されました。 20 人以上の世界クラスのオープンソース開発者が、維持しているリポジトリから現実的で多様かつやりがいのあるコーディング タスクを構築し、タスクごとに 40 時間以上を費やしました。彼らはリポジトリで「マージ可能」が何を意味するかを定義しています。
厳格な品質管理。ルーブリックの評価は主観的なものであるため、敵対的テスト、キャリブレーション、多段階レビューを備えた広範な QC パイプラインを構築し、すべてのタスクが認知研究者によって手動でレビューされます。 SWE-Bench Pro と比較して 81% 低い誤検知率を達成します。
私たちのベンチマークは、高品質で保守可能なコードを作成するモデルの能力を示す最も強力なシグナルを提供します。今日の最も高性能なモデルでさえ、この新しい標準では苦戦していることがわかりました。
20 人以上の世界クラスのオープンソース メンテナー
認知研究者による手動レビュー
コード品質を測定する史上初のベンチマーク
3 つのネストされたサブセットを示します。

FrontierCode の難易度を上げていく:Extended、Main、Diamond。ダイヤモンドは最も難しい 50 のタスクで構成され、メインは最も難しい 100 のタスク (ダイヤモンドを含む)、および拡張は 150 のフルセットで構成されます。
合格率とスコアという 2 つの指標をレポートします。
ソリューションは、すべてのブロッカー基準 (つまり、コードレビュー中に保守者がハードストップとみなす基準) をクリアした場合に合格し、そうでない場合は不合格となります。
ソリューションのスコアは、ルーブリック項目の加重集計です。ブロック基準を満たさないソリューションには 0 が与えられます。
各モデルは、利用可能な推論努力ごとに 5 回実行されます。各取り組みについて、5 つのトライアルにわたるメトリクスを平均し、最もパフォーマンスの高い推論レベルでの各モデルのスコアを報告します。
FrontierCode Diamond はまだ飽和していません。最もパフォーマンスの高いモデルである Claude Opus 4.8 のスコアはわずか 13.4% です。他のモデルのスコアは大幅に低く、GPT-5.5 は 6.3%、Gemini 3.1 Pro は 4.7%、その他のモデルはさらに低くなります。ただし、GPT 5.5 は一貫して Opus 4.8 よりも最大 4 倍少ないトークンを使用し、より優れたコスト インテリジェンスのトレードオフを実現します。
FrontierCode Main と Extended では、Opus 4.8 がそれぞれ 34.3% と 51.8% で依然として明確なリードを維持しています。また、オープンソース モデルとフロンティアの間には大きなギャップがあることも観察されています。最もパフォーマンスの高いオープンソース モデルである Kim K2.6 は、Diamond でわずか 3.8%、Main で 16%、Extended で 37% を達成しています。
この投稿の残りの部分では、FrontierCode を構築した理由と方法について詳しく説明します。
SWE-Bench Verified や Pro などの第 1 世代のコーディング ベンチマークは、能力の低いモデル向けに設計されました。これらは現実性と堅牢性の多くの基準に達していません。
基本的に、機能の正確性のみをテストし、品質はテストしません。さらに、これらのベンチマークでは誤分類エラーが発生する傾向があります。 METR の実験では、これらのモデルで高得点が得られることがわかりました。

ベンチマークでは、人間のメンテナが受け入れられないパッチが生成されることがよくあります。
誤分類をどのように定義すればよいでしょうか?これらは次の 2 つのカテゴリに分類されます。
誤検知: 検証者は、間違った解決策に報酬を与えるべきではありません。テスト カバレッジが不完全である可能性があり、モデルがまだ受け入れられている間違ったソリューションを作成する可能性があります。
偽陰性: 検証者は、正しいソリューションにペナルティを課すべきではありません。テストは具体的すぎる可能性があります。正確なエラー文字列や関数名をチェックするか、 unsolvable をチェックして、命令やコードベースにない動作をテストします。
エージェントの軌跡の分析を通じて、FrontierCode が他の主要なベンチマークよりも誤分類エラーを 81% 削減していることを示しています。これは、FrontierCode スコアが現在利用可能な最も正確なランキングであることを意味します。
既存のベンチマークには、いくつかの点で多様性の欠如もあります。
他のベンチマークはプログラムによるスクレイピングを介して単一の PR から問題を生成しましたが、FrontierCode は複数の PR チェーンと自由形式のリクエストからリポジトリ管理者によって手動で選択されます。また、SWE-Bench Pro で表現できる言語の数も 3 倍に増加しました。
また、既存のベンチマークは、過度に指定された詳細なプロンプトの形で過剰なガイダンスを提供することも知られています。今日のフロンティア モデルでは、手で持つ必要がはるかに少なくなります。 FrontierCode は、人間の貢献者と同じコンテキストを前提として、エージェントがメンテナの意図を推測することを期待します。
プロンプトには 2 つの部分が含まれています。まずはタスクの説明です。 2 つ目は、AGENTS.md にあるような、汎用テスト、lint、およびスタイルの実践に関するコードベース ガイドラインです。タスクの説明は人間らしく、意図的に簡潔にまとめられています。長さは SWE-Bench Pro の 3 分の 1 です。
SWE-Bench Pro、DeepSWE、FrontierCode のプロンプト例を比較する プロンプト例を表示

各ベンチマークの s を同じスケールで示します。各列内をスクロールして、構造、長さ、特異性を比較します。さらに、パッチのサイズを単に増やすのではなく、高品質のルーブリックを使用してタスクの難易度を調整することを選択しました。 DeepSWE のようなベンチマークよりもパッチが小さいにもかかわらず、FrontierCode はエージェントが解決するのが困難です。
FrontierCode と同じくらい野心的なコード品質の評価を行うには、ベンチマーク作成プロセスのすべてのステップに品質を組み込む必要がありました。
オープンソース メンテナのチーム
FrontierCode は、モデルが実稼働コードベースにマージされるコードを生成できるかどうかを測定することを目的としています。これを確実にするために、私たちは 36 の主力オープンソース リポジトリのメンテナと直接協力しました。このオールスターの専門家チームは、コードベースへの数千のコミットを共同でレビューし、マージしました。彼らは、文体やデザインに関する深い知識を、目にするすべての PR に適用することができます。
各メンテナはタスクごとに 40 時間以上を投資し、他の評価エンジニアや認知研究者と複数ラウンドの反復を繰り返しました。彼らはその判断を具体的な評価基準に絞り込み、これらの基準を満たすあらゆる PR が実際に承認されることになります。
FrontierCode について彼らは次のように語っています。
「FrontierCode のチームと協力できたのは光栄でした。AI 評価の問題に取り組むことは、まさに芸術であると感じました…他の人が CI のように採点するのに対し、FrontierCode は技術リーダーのように採点します。」 Tomer Nosrati 氏、Celery (28,600 個のスター) CEO 兼技術責任者
「FrontierCode を際立たせているのは、細部へのこだわりです。各タスクは、LLM ベンチマークではこれまでに見たことのない深さまで調整されています。私たちはゲーム可能なベンチマークから離れ、代わりに FrontierCode のようなベンチマークを使用して、本物のモデルを実証する必要があります。」

知性と創造性。」 Martin McKeaveney 氏、Budibase の共同創設者兼 CTO (28,000 つ星)
「オープンソース コミュニティの主要な専門家と協力できたことに感謝しています。私たちは、正確性と品質、そしてリポジトリのコンテキストにおけるマージ可能性が何を意味するかについて、深く議論しました。 FrontierCode は、現実世界の主観的な品質を尊重する AI モデルのマイルストーンです。」 Merlijn Vos、uppy のコアメンテナー (30.8k スター)
「FrontierCode の独自の価値は、その eval にエンコードされた人間の経験、つまりコードを高品質にし、マージに値するものにするものについての長年の判断から生まれています。あらゆる基準に強迫的なまでの配慮が払われているからこそ、このベンチマークが SWE の評価に新たな基準を設けると私は信じています。」 Claudio Costa、Mattermost のコアメンテナー (37,000 つ星)
FrontierCode は、次の軸に沿ってコードを評価することでマージ可能性を測定します。
動作の正しさ: パッチは問題を正常に解決しますか?
回帰安全性: 既存のコードベースで何かが壊れるか?
機械的な清潔さ: プロジェクトのビルド、糸くず、スタイルのチェックに合格していますか?
テストの正確性: エージェントのテストは実際に望ましい動作を捉えていますか?
範囲: パッチは必要な部分のみに適用されますか?
コードの品質: コードはコードベースの規約に準拠しており、健全な設計パターンに従っており、共同作業者が読みやすい状態に保たれていますか?
次の表は、これらの基準を評価するために、古典的な単体テストと、適応型古典的グレーディング、スコープ、逆古典的テスト (これらの方法の詳細は後述) などの新しい方法の両方を使用する方法を説明しています。
各基準は、ブロッカーまたは非ブロッカーのいずれかです。
ブロッカーは、マージ可能性の要件、つまり、コードレビュー中に保守者がハードストップを考慮する基準を表します。これらには、正しさのチェックと不正さのチェックが含まれます。

パフォーマンスや範囲の制限などの問題。
非ブロッカーは、コード スタイル、型安全性、可読性などの品質シグナルを表しますが、必ずしもマージをブロックするわけではありません。
ソリューションがすべてのブロッカーを満たしている場合、そのソリューションは合格とみなされ、そのスコアは、合格したすべてのルーブリック項目の加重集計になります。それ以外の場合は、スコア 0 を受け取ります。
誤分類に対する基準を強化しながら、複数の有効な解決策を考慮する余地を与えるために、次の 3 つの主要な手法を導入しました。
逆古典的: 逆古典的基準は、エージェントが作成したテストが意味のあるものであることを確認する方法です。つまり、元の壊れたコードベースでテストを実行すると、テストは失敗する必要があります。これにより、エージェントが問題を効果的なテストを作成できるほど十分に理解しているかどうかを、自動化された確定的なチェックで確認できるようになります。
コードスコープ: 優れた PR は自制心を働かせる必要があります。無関係なファイルに触れたり、不必要なリファクタリングを導入したりせず、必要なものだけを変更します。スコープ基準は、これらの境界を強制する自動チェックです。次の 3 種類の制約を組み合わせます。
files : どのファイルが許可されるか、拒否されるか、または削除する必要があるのかを迅速かつ決定的にチェックします。
size : 変更された行数、正味行数の増加、または変更されたファイルの合計に制限を適用します。
semantic : ファイルの特定の部分 (単一関数内など) 内での変更の局所性または性質を検証する LLM ベースのチェック用。
アダプティブ クラシック グレーディング: オープンエンドのコーディング タスクには、多くの有効な解決策が存在します。静的単体テストは厳格すぎます。優れたソリューションは、関数名やエラー文言などの表面的な違いによって失敗する可能性があります。この競合を mutagent で解決します。このツールは、LLM を使用してテスト環境 (またはアプリケーション コード) に外科的にパッチを適用し、

エージェントの実装の詳細を確認できるため、制限のないソリューションに対して厳密で決定的なテストを実行できます。
eval の実行 タスクの説明 次のように、すべての警告ログを新しい auto LOG_WARNING() -> std::ostream & src/logger.h のメソッドにカプセル化します。
警告は常に標準エラーに出力されます
警告は、--verbose とは関係なく、常に出力されます。
ヘルパーは自動的に警告を出力します: プレフィックス
この新しい関数は、コードベース全体の warning: <message> メッセージのすべてのインスタンスで使用します。
make を実行して、コードの変更が残っていないことを確認します。さらに多くのコード変更がある場合は、コードが適切にフォーマットされていないことを意味します。
コードの変更が既存のテスト ケースですでにカバーされていることが確実でない限り、常に関連するテストを (./test ディレクトリ内で) 編集または作成して、変更が機能することを確認し、リグレッションを防止してください。
テストは GoogleTest および POSIX シェル スクリプト (bash ではない) を使用して記述されており、実行するには test/CMakeLists.txt ビルド定義に登録する必要があります。
makeconfigurecompile を実行して、コードをインプレースでコンパイルし、フォーマットします。コンパイル手順には、大量のリンターのようなチェックが含まれます。
すでに正しいベースコミットを行っています。このコミットからブランチを作成します。リベースしたり、マスター、メイン、その他のブランチから開始したりしないでください。
「Run eval」を押して生成します

[切り捨てられた]

## Original Extract

Today’s coding benchmarks have established that models can write correct code, but the question we should really be asking is: can models actually write good code?

Introducing FrontierCode | Cognition Menu Close Home
Raising the bar from correctness to quality #
Today’s coding benchmarks have established that models can write correct code. But as AI-generated code becomes the dominant path to production, correctness is now table stakes. The question that we should be asking is: can models actually write good code?
We’re excited to introduce FrontierCode, a benchmark that measures how well models can truly meet the standards of high-quality production codebases. What sets us apart:
Would the maintainer actually merge this PR? We’re the first benchmark to measure code mergeability. Our criteria assess end-to-end code quality — correctness, test quality, scope discipline, style, and adherence to codebase standards. This employs a novel ensemble of grading techniques, including unit tests, rubrics, and new types of verifiers.
Crafted by open-source maintainers. 20+ world-class open-source developers built realistic, diverse, and challenging coding tasks from the repos they maintain, spending more than 40 hours per task. They define what “mergeable” means in their repo.
Rigorous quality control. Rubric grading is subjective, so we built an extensive QC pipeline with adversarial testing, calibration, and multi-stage review, where every task is manually reviewed by a Cognition researcher. We achieve an 81% lower false positive rate compared to SWE-Bench Pro.
Our benchmark provides the strongest available signal of a model’s ability to write high-quality, maintainable code. We find that even today’s most capable models struggle on this new standard.
20+ world-class open-source maintainers
Manually reviewed by Cognition researchers
First-ever benchmark measuring code quality
We present three nested subsets of FrontierCode at increasing difficulty: Extended, Main, and Diamond. Diamond comprises the 50 hardest tasks, Main the 100 hardest (including Diamond), and Extended the full set of 150.
We report two metrics, pass rate and score :
A solution passes if it clears all blocker criteria, i.e., criteria that a maintainer would consider hard stops during code review, and fails otherwise.
A solution’s score is a weighted aggregate of the rubric items. Solutions that do not pass blocking criteria receive 0.
Each model is run 5 times at every available reasoning effort. For each effort, we average the metric across the 5 trials, then report each model’s score at its best performing reasoning level.
FrontierCode Diamond remains unsaturated: the best performing model, Claude Opus 4.8, achieves a score of only 13.4%. Other models score significantly lower: GPT-5.5 receives 6.3%, Gemini 3.1 Pro 4.7%, and others even less. However, GPT 5.5 consistently uses up to 4x fewer tokens than Opus 4.8, achieving a better cost-intelligence tradeoff.
On FrontierCode Main and Extended, Opus 4.8 still maintains a clear lead, at 34.3% and 51.8%, respectively. We also observe a large gap between open-source models and the frontier. Kimi K2.6, the best-performing open-source model, achieves just 3.8% on Diamond, 16% on Main and 37% on Extended.
The rest of this post will be a deep dive into why and how we built FrontierCode.
The first generation of coding benchmarks, such as SWE-Bench Verified and Pro, were designed for less capable models. They fall short on many measures of realism and robustness.
Fundamentally, they only test functional correctness , not quality. Moreover, these benchmarks are prone to misclassification errors . Experiments from METR have found that high-scoring models on these benchmarks often produce patches that wouldn’t be accepted by human maintainers.
How do we define misclassifications? These fall under two categories:
False Positives: The verifier should not reward solutions that are wrong. Test coverage may be incomplete , allowing the model to write an incorrect solution that’s still accepted.
False Negatives: The verifier should not penalize solutions that are correct. Tests can be either too specific , e.g. checking for exact error strings or function names, or unsolvable , testing for a behavior not in the instruction or in the codebase.
We show through analysis of agent trajectories that FrontierCode produces 81% less misclassification errors than other leading benchmarks. This means that FrontierCode scores are the most accurate ranking currently available.
Existing benchmarks also suffer from lack of diversity in several ways.
While other benchmarks generated issues from single PRs via programmatic scraping, FrontierCode is hand-selected by repo maintainers from multi-PR chains and freeform requests. We also triple the number of represented languages from SWE-Bench Pro.
It’s also known that existing benchmarks provide too much guidance in the form of overly specified and detailed prompts. Today’s frontier models need far less hand-holding. FrontierCode expects the agent to infer the maintainer’s intent, given the same context as a human contributor.
Our prompts contain two parts. First is the task description. Second, the codebase guidelines for generic testing, lint, and style practices, just like those found in AGENTS.md. The task descriptions are humanlike and deliberately concise — a third the length of SWE-Bench Pro’s.
Compare example prompts across SWE-Bench Pro, DeepSWE, and FrontierCode Show Example prompts from each benchmark, shown at the same scale. Scroll within each column to compare structure, length, and specificity. Furthermore, we’ve chosen to scale the difficulty of tasks using quality rubrics , rather than simply increasing patch size. Despite having smaller patches than benchmarks like DeepSWE, FrontierCode is harder for agents to solve.
To produce an evaluation for code quality as ambitious as FrontierCode, we had to embed quality into every step of the benchmark creation process.
A Team of Open Source Maintainers
FrontierCode aims to measure whether models can produce code that would be merged into production codebases. To ensure this, we collaborated directly with the maintainers of 36 flagship open-source repositories. This team of all-star experts has collectively reviewed and merged thousands of commits to their codebases. They can apply deep stylistic and design knowledge to every PR they see.
Each maintainer invested more than 40 hours per task, undergoing multiple rounds of iteration with other eval engineers and Cognition researchers. They’ve distilled their judgment into concrete evaluation criteria: any PR that satisfies these standards would actually be approved.
Here’s what they say about FrontierCode:
“Working with the team behind FrontierCode was a privilege. Taking on the AI evaluation problem felt like nothing less than an art… Where others grade like a CI, FrontierCode grades like a tech lead.” Tomer Nosrati, CEO and Tech Lead of Celery (28.6k stars)
“What sets FrontierCode apart is the attention to detail. Each task is calibrated to a depth that simply hasn’t been seen before in LLM benchmarking. We should be moving away from benchmarks that can be gamed and instead using ones like FrontierCode to demonstrate genuine model intelligence and creativity.” Martin McKeaveney, Co-Founder and CTO of Budibase (28k stars)
“I’m grateful to have worked with leading experts in the Open Source community. We had deep discussions on correctness versus quality and what mergeability means in the context of their repository. FrontierCode is a milestone for AI models respecting subjective quality in the real world.” Merlijn Vos, Core Maintainer of uppy (30.8k stars)
“FrontierCode’s unique value comes from the human experience encoded in its evals: years of judgment about what makes code high-quality and worthy of merging. The almost obsessive care brought to every criterion is why I believe this benchmark sets a new bar for SWE evaluation.” Claudio Costa, Core Maintainer of Mattermost (37k stars)
FrontierCode measures mergeability by evaluating code along the following axes:
Behavioral correctness: Does the patch successfully solve the problem?
Regression safety: Does it break anything in the existing codebase?
Mechanical cleanliness: Does it pass the project’s build, lint, and style checks?
Test correctness: Do the agent’s tests actually capture the desired behavior?
Scope: Does the patch touch only what it needs to?
Code quality: Does the code conform to codebase conventions, follow sound design patterns, and remain readable to collaborators?
The following table describes how we use both classical unit tests and novel methods, such as adaptive classical grading, scope, and reverse-classical tests (more on these methods below) to evaluate these criteria.
Each criterion is either a blocker or a non-blocker :
Blockers represent mergeability requirements, i.e., criteria that a maintainer would consider hard stops during code review. These include correctness checks, as well as non-correctness concerns like performance or scope restrictions.
Non-blockers represent quality signals such as code style, type safety, and readability, which would not necessarily block a merge.
If a solution satisfies all the blockers, it is considered passing, and its score is the weighted aggregate of all the rubric items it passes. Otherwise it receives a score of zero.
We’ve introduced three main techniques to strengthen criteria against misclassifications, while allowing space for multiple valid solutions:
Reverse-Classical: The reverse-classical criterion is a way to ensure that agent-written tests are meaningful: when we run them on the original, broken codebase, they must fail . This gives us an automated, deterministic check that the agent understood the problem well enough to write an effective test for it.
Code Scope: A good PR should exercise restraint : it modifies only what it needs to, without touching unrelated files or introducing unnecessary refactors. The scope criterion is an automated check that enforces these boundaries. It combines three types of constraints:
files : For fast, deterministic checks on which files can be allowed , denied , or must be deleted .
size : To enforce limits on the number of changed lines , net line growth , or total files modified.
semantic : For LLM-based checks that verify the locality or nature of a change within a specific part of a file (e.g., inside a single function).
Adaptive Classical Grading: Open-ended coding tasks can have many valid solutions. Static unit tests are too rigid; good solutions can fail for superficial differences like function names or error wording. We resolve this conflict with mutagent , a tool we built that uses an LLM to surgically patch the test environment (or the application code) and align with the agent’s implementation details, allowing us to run rigorous, deterministic tests on open-ended solutions.
Run eval Task description Encapsulate all warning logs in a new auto LOG_WARNING() -> std::ostream & method in src/logger.h such that:
Warnings are always printed to standard error
Warnings are always printed, independently of --verbose
The helper automatically prints the warning: prefix
Use this new function in every instance of warning: <message> messages throughout the codebase.
Run make and ensure no code changes remain. If there are more code changes, then it means that the code was not formatted properly.
Unless you are sure that the code change is already covered by an existing test case, always edit or create relevant tests (in the ./test directory) to confirm the changes work and prevent regressions.
The tests are written using GoogleTest and POSIX shell scripts (not bash) and must be registered in the test/CMakeLists.txt build definition to run.
Run make configure compile to compile and format the code in-place. The compile step comes with a large amount of linter-like checks.
You are already on the correct base commit. Create your branch from this commit. Do not rebase or start from master, main, or any other branch.
Press "Run eval" to generate

[truncated]
