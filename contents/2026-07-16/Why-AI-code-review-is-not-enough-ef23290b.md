---
source: "https://blog.codacy.com/ai-code-review-is-not-enough-how-engineering-leaders-should-gate-ai-generated-code"
hn_url: "https://news.ycombinator.com/item?id=48932183"
title: "Why AI code review is not enough"
article_title: "AI Code Review Is Not Enough: How Engineering Leaders Should Gate AI-Generated Code"
author: "ARayOutOfBounds"
captured_at: "2026-07-16T09:58:00Z"
capture_tool: "hn-digest"
hn_id: 48932183
score: 1
comments: 1
posted_at: "2026-07-16T09:23:41Z"
tags:
  - hacker-news
  - translated
---

# Why AI code review is not enough

- HN: [48932183](https://news.ycombinator.com/item?id=48932183)
- Source: [blog.codacy.com](https://blog.codacy.com/ai-code-review-is-not-enough-how-engineering-leaders-should-gate-ai-generated-code)
- Score: 1
- Comments: 1
- Posted: 2026-07-16T09:23:41Z

## Translation

タイトル: AI コードレビューだけでは不十分な理由
記事のタイトル: AI コード レビューだけでは十分ではない: エンジニアリング リーダーは AI で生成されたコードをどのようにゲートすべきか
説明: AI によって生成されたコードのリスクと、ソフトウェア開発のセキュリティと品質を確保するための堅牢なレビュー プロセスの必要性を探ります。

記事本文:
AI コード レビューだけでは十分ではない: エンジニアリング リーダーは AI によって生成されたコードをどのようにゲートすべきか
承認するのが最も簡単なプル リクエストは、多くの場合、綿密な検査に値します。 AI によって生成されたコードは、適切にフォーマットされ、十分に文書化され、何が変更されたのかについての説得力のある説明が伴って到着する傾向があります。
これらの情報は、実装が安全かどうか、新しい依存関係によってリスクがもたらされるかどうか、動作が実際に要求されたものと一致するかどうかを示しません。これらの質問には独自のゲートが必要です。残りは、それらのゲートがどこに属するか、何を強制する必要があるか、そしてなぜ AI レビューだけでは十分ではないのかを判断することです。
AI によって生成されたコードが、実際に運用される前に本番環境に対応しているように見える理由
ユーザーが表示名とプロフィールを更新できるように、プロフィール編集エンドポイントを追加するようにアシスタントに依頼する開発者を想像してください。アシスタントは、新しいルート、移行、バイオをレンダリングするための React コンポーネント、およびいくつかの合格した単体テストを返します。
差分には何も異常はなく、ビルドは緑色で、PR の説明は注意深くエンジニアが書いたもののように読めます。ただし、同じ変更の中に、DOM に書き込む前にユーザー入力を決してエスケープしないレンダリング パスや、誰もチェックしていないオープン アドバイザリとのマークダウン解析の依存関係が埋め込まれている可能性があります。
これは、直接名前を付ける価値のあるパターンです。最新の LLM は、多くの場合、機能的に正しいコードを生成できますが、安全なコード生成は依然として非常に難しい問題です。
実世界のソフトウェア リポジトリで LLM を評価する最近の独立したベンチマークでは、モデルが機能的に正しい実装を生成できる一方でセキュリティ上の脆弱性が依然として存在すること、および機能の正確性を向上させる技術がセキュリティの結果を確実に向上させるわけではないことが一貫して判明しています。
つまり、変更は本番環境に対応しているように見える可能性があります。

テストを行っていますが、安全にマージする前に決定的なセキュリティ チェックが必要です。
このばらつきは、ヘッドラインの合格率よりも重要です。ギャップを埋めるために一般的なモデルの改善を信頼するのではなく、エンジニアリング リーダーに決定的なゲートをどこに置くかを指示するからです。新しいモデルがデフォルトでより安全であると想定しているチームは、データが裏付けていない賭けをしていることになります。
実際的な対応策は、AI 支援のプル リクエストを個別のレビュー パスとして扱うことです。ラベルを付け、頑固な障害クラスを捕捉するセキュリティ関連のチェックを要求し、グリーン ビルドを準備の十分な証拠として扱うのではなく、マージ前にブランチ保護が実際に結果を評価することを確認します。
AI コードのリスクについて、2026 年の最近の証拠が示していること
セキュリティ合格率の差が全体像の一部を説明していますが、レビュー能力をどこに投資するかを決定しようとしているチームにとって、運用面も同様に重要です。
4,000 チームを超える 22,000 人の開発者からのデータに基づく Faros AI のテレメトリ調査では、組織が AI 導入を低レベルから高レベルに移行するにつれて、インシデントとプル リクエストの比率が 242.7% 増加したことがわかりました。
言い換えれば、マージされたプル リクエストに対する本番インシデントの数は、各組織の AI 導入が低いベースラインのときと比べて 3 倍以上高かったということです。
同じレポートでは、レビュー (人間または AI) を行わずにマージされたプル リクエストが 31.3% 増加したことがわかり、レビューの実践がコード スループットの増加に追いついていないことを示唆しています。
このデータを素直に読み取ると、人間のペースで出力するように構築されたレビュー プロセスは、コードを処理するために進化するレビュー実践よりも早くコードがレビューに入ると圧倒され、何かが与えられなければならないということになります。
AI 導入を純粋な生産性の話として扱い、何が起こるかを再検討しなかったチーム

PR ボタンの下流にあるものは、品質メトリクスが低下してこのデータに表示されます。
これらの数値はいずれも、特に、そのストーリーに関心のあるベンダー提供のベンチマークなど、福音ではなく方向性を示すものとして扱う価値があります。エンジニアリング リーダーにとって有益な演習は、パーセンテージを記憶することではなく、自分自身の数値のベースラインを再設定することです。
チームの AI ロールアウト前の脆弱性密度、レビュー レイテンシー、回避された欠陥率を取得し、それらを現在の同じ指標と比較します。傾向線が業界データの説明と一致する場合、人間のペースの世界で機能していたゲートを、機械のペースの世界に合わせて再構築する必要があります。
AI が生成した 1 つのプルリクエストがどのようにして安全対策を回避できるのか
別の例を考えてみましょう。開発者は、ユーザーが公開プロフィールを更新できるようにアシスタントに依頼します。このアシスタントは、新しい API エンドポイント、データベース移行、プロフィール ページに自己紹介を表示するコンポーネント、基本的な書式設定をサポートする新しいマークダウン解析ライブラリ、および有効な更新が正しく保存されることを確認する少数の単体テストを生成します。
PR は十分に小さいため、レビュー担当者は数分でざっと読むことができ、テストは合格し、生成された概要は明確に読めます。
ここで、スキムと AI が生成した解説を基に作業する査読者が重大なリスクを見逃す可能性があります。マークダウン レンダラーは、サニタイズされていない HTML をページに渡し、安全でないレンダリング パターンを特にチェックしていないユーザーに対して、通常の React コードのように見える XSS パスを開く可能性があります。
新しいマークダウン依存関係には、オープン アドバイザリが含まれたり、依存関係スキャナーの外部の誰もソースを読んでも気づかない危険な推移的なパッケージが取り込まれたりする可能性があります。テスト フィクスチャには、プレースホルダのように見えるトークンが含まれている可能性がありますが、アシスタントは現実的な外観を頻繁に生成するため、プレースホルダではありません。

g 認証情報は、レビュー担当者が開く可能性が最も低いファイルに保存されます。
また、テスト自体はハッピー パスのみを検証し、ユーザーが自分の経歴を更新できることを確認する一方、同じエンドポイントでユーザーが他の人の経歴を更新できるかどうかは確認しません。
これらの問題はいずれも、PR ボリュームが低く、差分が小さいときに機能した一種のリードスルー レビューからは見えません。
それらを捕らえるのは、コードがどれほどきれいに見えるかに関係なく実行される一連のゲートです。フィクスチャを含むすべてのファイルをスキャンするシークレット検出、新しい依存関係のアドバイザリにフラグを立てるソフトウェア構成分析、エスケープされていないレンダリングをキャッチする静的分析、およびテストされていない認可パスを通知するカバレッジ チェックです。
PR は依然として人間の目から恩恵を受けていますが、マージの決定は、自信に満ちた AI の要約や時間のないレビュー担当者によって放棄できないゲートに依存する必要があります。
AIがAIを再検討する理由は十分ではないのか
レビューのボトルネックに対する魅力的な解決策は、AI 作成者の上に AI レビュー担当者を追加することです (現在、GitHub 上のコード レビューの 5 件に 1 件以上にエージェントが関与しています)。境界が明確な場合、その設定には真の価値があります。
多くの場合、同様のデータでトレーニングされ、同様のコンテキストで動作する 2 番目のモデルが最終的な単語になると、問題が発生します。
Faros AI のデータは、ここで有用なデータ ポイントを提供します。調査した組織のプル リクエストの約 4 分の 1 が AI エージェントによってレビューされましたが、完全に自律的にオープンされたのは 1% 未満でした。つまり、このシステムのリスクのほとんどは、たとえレビュー自体が自動化されていたとしても、名目上は人間が責任を負うコードを通じて依然として実行されているということです。
さらに深刻な問題は、AI 支援開発に関する研究で一貫して現れる認識のギャップです。 METRのランダム化比較試験では、

自分のリポジトリで AI ツールを使用している経験豊富な開発者は、実際のタスクを完了するまでにかなりの時間がかかりましたが、より速く進んだと信じており、認識されている生産性と実際の生産性の間に大きなギャップが生じています。
この分離はレビューにおいて特に重要です。なぜなら、AI によって変更がシンプルかつ安全に行われたと確信している開発者やレビュー担当者は、基礎となる変更が手書きの変更とまったく同じくらい複雑でリスクがある場合でも、その変更を綿密に精査する可能性が低いからです。
スピード感に基づいて築かれた自信は、検証に基づいて築かれた自信とは同じではありません。人間もモデルも同じようにループ内の全員が、何が正しく見えるかという同じ感覚に基づいて構築されている場合、この 2 つは簡単に混同されます。
これはいずれも、AI レビューがワークフローに存在しないことを意味するものではありません。人間が完全な変更を読み取るよりも早く、差分を要約したり、明らかな間違いにフラグを立てたり、欠落しているエッジケース テストを提案したりできます。
重要な違いは、コードをマージするかどうかを決定する権限ではなく、その出力を複数の入力のうちの 1 つとして扱うことです。セキュリティ関連のマージ ゲートは、最初にコードを作成したモデルではなく、独立した品質ゲートと責任あるレビュー担当者に依存する必要があります。
AI審査を執行とみなさずに活用する方法
AI レビューは、リリースされる内容の記録システムとしてではなく、生産性の補助としての地位を獲得しています。
うまく使用すると、大規模な diff の平易な要約を生成できるため、人間のレビュー担当者は 1 行ずつ読む前にどこに焦点を当てるべきかを知ることができます。また、特に認可とエラー処理に関して、開発者が書くことを考えていなかったかもしれないテスト ケースを提案することもできます。
また、合理的なコーチング ツールでもあり、あるパターンが他のパターンよりも安全である理由を、不合格品よりも定着する方法で若手エンジニアに説明できます。

d PR コメントのみ。
境界は仮定ではなく明示する必要があります。 AI によって生成された承認が、必要なレビュー ルールを黙って満たすことは決してあってはならないし、機密ファイル、認証ロジック、支払いフロー、顧客データに関わるあらゆるものに対するコード所有者のサインオフは、AI レビュー担当者がどのような結論を出したとしても、人間の責任であり続ける必要があります。
AI アシスタントがフラグが立てられた問題の修正を提案する場合、その修正は問題を解決したというアシスタントの言葉で受け入れられるのではなく、元の問題を捕捉したのと同じスキャナーを通過する必要があります。そして、AI の解説とセキュリティまたは取材ゲートの意見が一致しない場合は、ゲートが勝ちます。
AI レビューをアドバイザリーとして構成し、ブランチ保護がマージ条件として決定論的なチェックと責任あるレビュー担当者のみを認識することで、モデルに独自の作業をグレーディングさせることで生じる盲点を引き継ぐことなく、生産性の利点を維持できます。
これは、AI 支援開発に対する Codacy のアプローチの背後にある哲学でもあります。AI は、開発者が変更をレビューして理解するのに役立ちますが、独立した品質、セキュリティ、依存関係、およびポリシー ゲートは、何が本番環境に到達するかを決定する責任を引き続き負います。
実践的なレビューと施行のワークフローを構築する方法
これらのアイデアをチームが日常的に実行できるものに変えるには、大規模なプログラムではなく、一連の短い意思決定が必要になります。
AI を利用したプル リクエストを特定します。ラベル、コミット規約、または開発者の宣言だけで十分です。目標は、生成されたコードがコードベースのどこに入っているかを可視化することであり、ツールの使用に責任を負わせることではありません。
例外なくすべての PR にベースライン ゲートを適用します。シークレットの検出、依存関係のスキャン、静的分析、テスト カバレッジの移動は、作成者に関係なく、すべての変更に対して実行する必要があります。
ハードルを上げる

r リスクの高い変更。認証、承認、支払いフロー、顧客データ、インフラストラクチャの変更は、AI がそれらの作成に関与したかどうかに関係なく、人間による必須のレビューとより厳格なしきい値をトリガーする必要があります。
AI レビューのアドバイスを維持します。要約と提案を行いますが、モデルの出力ではなく、必要なチェックと名前付きレビュー担当者とのマージ権限を保持します。
必要なチェックをオーバーライドできるユーザーを制限し、すべての例外をログに記録します。誰もレビューしないバイパスは、後で事件ではなくパターンになります。
レビュー システムがスループットに追随しているかどうかを追跡します。 AI 支援のコード量が増加する一方で、レビューのレイテンシ、回避された欠陥、または未レビューのマージも増加している場合、ギャップがさらに拡大する前にゲートを再構築する必要があります。
AI コーディング ツールはスループットを真に向上させますが、スループットはエンジニアリング リスクを排除するのではなく、その形を変えます。
これらのツールから最大限の価値を得ている組織は、AI 支援開発を、コードを作成したモデルとは独立して検証され、変更点で一貫して発生する施行と、その施行の状態を埋もれたものではなくエンジニアリングのリーダーに可視化するメトリクスを結び付けている組織です。

[切り捨てられた]

## Original Extract

Explore the risks of AI-generated code and the necessity for robust review processes to ensure security and quality in software development.

AI Code Review Is Not Enough: How Engineering Leaders Should Gate AI-Generated Code
The easiest pull requests to approve are often the ones that deserve the closest inspection. AI-generated code tends to arrive well formatted, well documented, and accompanied by a convincing explanation of what changed.
None of those things tell you whether the implementation is secure, whether a new dependency introduces risk, or whether the behavior matches what was actually requested. Those questions need their own gates. The rest is deciding where those gates belong, what they should enforce, and why AI review alone is not enough.
Why AI-generated code looks production-ready before it is
Picture a developer asking an assistant to add a profile-edit endpoint so users can update their display name and bio. The assistant returns a new route, a migration, a React component to render the bio, and a handful of passing unit tests.
Nothing in the diff looks unusual, the build is green, and the PR description reads like something a careful engineer would have written. Buried in that same change, though, might be a rendering path that never escapes user input before writing it to the DOM, or a markdown parsing dependency with an open advisory nobody checked.
This is the pattern worth naming directly: Modern LLMs can often generate functionally correct code, but secure code generation remains a significantly harder problem.
Recent independent benchmarks evaluating LLMs on real-world software repositories consistently find that models can generate functionally correct implementations while still introducing security vulnerabilities , and that techniques which improve functional correctness do not reliably improve security outcomes.
That means a change can look production-ready, pass its tests, and still require deterministic security checks before it is safe to merge.
That unevenness matters more than the headline pass rate, because it tells engineering leaders where to put deterministic gates rather than trusting general model improvement to close the gap. A team that assumes newer models are safer by default is making a bet the data does not support.
The practical response is to treat AI-assisted pull requests as a distinct review path: label them, require the security-relevant checks that catch the stubborn failure classes, and make sure branch protection actually evaluates the result before merge rather than treating a green build as sufficient proof of readiness.
What recent 2026 evidence says about AI code risk
The security pass rate gap explains part of the picture, but the operational side matters just as much for a team trying to decide where to invest review capacity.
Faros AI’s telemetry study , based on data from 22,000 developers across more than 4,000 teams, found that as organizations moved from low to high AI adoption, the incidents-to-pull-request ratio increased by 242.7%.
In other words, the number of production incidents relative to merged pull requests was more than three times higher than during each organization’s low-AI-adoption baseline.
The same report found that pull requests merged without any review (human or AI) increased by 31.3%, suggesting that review practices did not keep pace with the increase in code throughput.
The honest reading of this data is that a review process built for human-paced output gets overwhelmed when code enters review faster than review practices evolve to handle it, and something has to give .
Teams that treated AI adoption as a pure productivity story, without re-examining what happens downstream of the PR button, are the ones showing up in this data with degraded quality metrics.
Any of these figures are worth treating as directional rather than gospel, particularly vendor-sponsored benchmarks that have an interest in the story they tell. The useful exercise for an engineering leader is not memorizing a percentage but re-baselining your own numbers.
Pull your vulnerability density , review latency, and escaped defect rate from before your team's AI rollout and compare them to the same metrics today. If the trend lines match what the industry data describes, the gates that worked in a human-paced world need to be rebuilt for a machine-paced one.
How one AI-generated pull request can bypass safeguards
Consider another example. A developer asks an assistant to let users update their public bio. The assistant produces a new API endpoint, a database migration, a component to render the bio on the profile page, a new markdown parsing library to support basic formatting, and a small set of unit tests confirming that a valid update saves correctly.
The PR is small enough that a reviewer can skim it in a few minutes, the tests pass, and the generated summary reads clearly.
Here is where a reviewer working from a skim plus AI-generated commentary can miss substantial risk. The markdown renderer may pass unsanitized HTML into the page, opening an XSS path that looks like ordinary React code to someone who is not specifically checking for unsafe rendering patterns.
The new markdown dependency might carry an open advisory or pull in a risky transitive package that nobody outside a dependency scanner would notice by reading source. A test fixture might contain a token that looks like a placeholder but is not, since assistants frequently generate realistic-looking credentials in exactly the files reviewers are least likely to open.
And the tests themselves might validate only the happy path, confirming that a user can update their own bio while never checking whether that same endpoint lets a user update someone else's.
None of these problems are visible from the kind of read-through review that worked when PR volume was lower and diffs were smaller.
What catches them is a set of gates that run regardless of how clean the code looks: secrets detection scanning every file including fixtures, software composition analysis flagging the new dependency's advisory, static analysis catching the unescaped render, and a coverage check that notices authorization paths went untested.
The PR still benefits from human eyes, but the merge decision should depend on gates that cannot be waived by a confident-sounding AI summary or a reviewer running short on time.
Why AI re viewing AI is not enough
A tempting fix for the review bottleneck is to add an AI reviewer on top of the AI author ( more than one in five code reviews on GitHub now involve an agent) and there is real value in that setup when the boundaries are clear.
The problems emerge when a second model, often trained on similar data and operating with similar context, becomes the final word .
Faros AI's data offers a useful data point here: roughly a quarter of pull requests in the organizations it studied were reviewed by AI agents, while fewer than one percent were opened fully autonomously, meaning most of the risk in this system still runs through code a human is nominally accountable for even when the review itself was automated.
The deeper issue is a perception gap that shows up consistently in research on AI-assisted development. A randomized controlled trial from METR found that experienced developers using AI tools on their own repositories took measurably longer to complete real tasks, yet believed they had gone faster, producing a large gap between perceived and actual productivity.
That separation matters for review specifically, because a developer or reviewer who feels confident that AI made a change simple and safe is less likely to scrutinize it closely, even when the underlying change is exactly as complex and risky as one written by hand.
Confidence built on a feeling of speed is not the same as confidence built on verification , and the two get conflated easily when everyone in the loop, human and model alike, is drawing from a similar sense of what looks correct.
None of this means AI review has no place in the workflow. It can summarize a diff, flag an obvious mistake, or suggest a missing edge-case test faster than a human can read the full change.
The distinction that matters is treating that output as one input among several rather than the authority that decides whether code merges. Merge gates for anything security-relevant should depend on independent quality gates and accountable reviewers, not on the model that wrote the code in the first place.
How to use AI review without treating it as enforcement
AI review earns its place as a productivity aid, not as the system of record for what gets released.
Used well, it can generate a plain-language summary of a large diff so a human reviewer knows where to focus before reading line by line, and it can propose test cases a developer might not have thought to write, particularly around authorization and error handling.
It is also a reasonable coaching tool, capable of explaining to a junior engineer why one pattern is safer than another in a way that sticks better than a rejected PR comment alone.
The boundaries need to be explicit rather than assumed. AI-generated approval should never silently satisfy a required review rule, and code-owner sign-off on sensitive files, authentication logic, payment flows, anything touching customer data, should stay a human responsibility regardless of what an AI reviewer concludes.
When an AI assistant proposes a fix for a flagged issue, that fix needs to go back through the same scanners that caught the original problem rather than being accepted on the assistant's word that it resolved the issue. And when AI commentary and a security or coverage gate disagree, the gate wins.
Configuring AI review as advisory, with branch protection recognizing only deterministic checks and accountable reviewers as merge conditions, keeps the productivity benefit without inheriting the blind spots that come from letting a model grade its own work.
This is also the philosophy behind Codacy ’s approach to AI-assisted development: AI can help developers review and understand changes, while independent quality, security, dependency, and policy gates remain responsible for deciding what reaches production.
How to build a practical review and enforcement workflow
Turning these ideas into something a team can run day to day comes down to a short sequence of decisions rather than a large program.
Identify AI-assisted pull requests . A label, a commit convention, or a developer declaration is enough. The goal is visibility into where generated code is entering the codebase, not assigning blame for using the tools.
Apply baseline gates to every PR, without exception. Secrets detection, dependency scanning , static analysis, and test coverage movement should run on every change regardless of who or what authored it.
Raise the bar for high-risk changes. Authentication, authorization, payment flows, customer data, and infrastructure changes should trigger mandatory human review and stricter thresholds, whether or not AI was involved in writing them.
Keep AI review advisory. Let it summarize and suggest, but keep merge authority with required checks and named reviewers, not model output.
Restrict who can override required checks, and log every exception. A bypass that nobody reviews later becomes a pattern rather than an incident.
Track whether your review system is keeping pace with throughput. If AI-assisted code volume is climbing while review latency, escaped defects, or unreviewed merges are also climbing, the gates need to be rebuilt before the gap widens further.
AI coding tools genuinely increase throughput, but throughput changes the shape of engineering risk rather than eliminating it.
The organizations getting the most value value from these tools are the ones connecting AI-assisted development to enforcement that happens consistently at the point of change , verified independently of the model that wrote the code, with metrics that make the state of that enforcement visible to engineering leadership rather than buried i

[truncated]
