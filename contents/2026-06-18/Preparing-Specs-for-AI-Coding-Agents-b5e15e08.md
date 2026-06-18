---
source: "https://forkline.dev/blog/assignment-layer-team-visible-specs/"
hn_url: "https://news.ycombinator.com/item?id=48587036"
title: "Preparing Specs for AI Coding Agents"
article_title: "Preparing Specs for AI Coding Agents - Forkline Blog"
author: "pando85"
captured_at: "2026-06-18T16:13:19Z"
capture_tool: "hn-digest"
hn_id: 48587036
score: 1
comments: 0
posted_at: "2026-06-18T15:36:21Z"
tags:
  - hacker-news
  - translated
---

# Preparing Specs for AI Coding Agents

- HN: [48587036](https://news.ycombinator.com/item?id=48587036)
- Source: [forkline.dev](https://forkline.dev/blog/assignment-layer-team-visible-specs/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T15:36:21Z

## Translation

タイトル: AI コーディング エージェントの仕様の準備
記事のタイトル: AI コーディング エージェントの仕様の準備 - Forkline ブログ
説明: AI コーディング エージェントがリポジトリの変更を開始する前に、適切なコンテキスト、境界、例、検証基準を提供する方法。

記事本文:
AI コーディング エージェントの仕様の準備 - Forkline ブログ
製品
機能 別のプロセス層を導入せずにランナーの作業を開始 コネクタ GitHub、GitLab、Gitea、Forgejo など 開発者 AI ランナー実行用の 1 つの操作面 ドキュメント ブログ 料金 FAQ
フォークラインを試す
ブログに戻る
フォークライン エンジニアリング AI コーディング エージェントの仕様の準備
AI コーディング エージェントがリポジトリの変更を開始する前に、適切なコンテキスト、境界、例、検証基準を提供する方法。
AI コーディング エージェントは、リポジトリを編集し、コマンドを実行し、ブランチを生成するようになりました。そのため、作業前の仕様がより重要になります。仕様には、エージェントが必要とするコンテキスト、境界、成功基準が含まれています。
優れたコーディング エージェント仕様には何が含まれますか
AI コーディング エージェントはもはや質問に答えるだけではないため、仕様の重要性が増しています。彼らは
リポジトリの読み取り、ファイルの編集、コマンドの実行、ブランチの作成、人間によるレビューの依頼などを行っています。
結果。それにより、プロンプトがどうなる必要があるかが変わります。
アシスタントが質問に答えるだけの場合は、プライベート プロンプトで十分です。エージェントが共有を変更するとき
コードベースでは、プロンプトは割り当てになります。そして、課題には適切な言葉遣い以上のものが必要です。それには、
正しい文脈、境界、例、そして作品が本来の意図と一致しているかどうかを判断する方法。
これが、コーディング エージェントをリポジトリに送信する前に仕様を準備する実際的な理由です。仕様
長くする必要はありません。どのような問題を解決しているのか、どのような動作をすべきなのかをエージェントに伝える必要があります。
何を変えるのか、何を変えてはいけないのか、その結果はどのように評価されるのか。
適切なコーディング エージェント仕様では、少なくとも次の 5 つのことがエージェントに提供される必要があります。
変えるべき行動
エージェントが保持すべき制約
正しさを定義する例またはシナリオ
検証

査読者が検査すべき証拠
これは、仕様駆動型開発、動作シナリオ、問題テンプレート、軽量の背後にある有用なアイデアです。
設計ドキュメント、OpenSpec、GitHub Spec Kit、および多くの社内エンジニアリング提案フォーマット。具体的な
フレームワークは仕様の形状よりも重要です。エージェントは動作するために十分なコンテキストを受け取る必要があります。
チームは結果をレビューするための十分な体制を整える必要があります。
この仕様は、これほど優れたプロンプトではありません。これは、人間の意図と機械の実行の間であらかじめ用意された割り当てです。
プロンプトは仕事を始めるのに適しています。スペック的には持ち運びに優れています。
プライベート プロンプトは即時性を重視して最適化されています。それはチャットセッションに住んでいます。省略表現、省略表現が含まれる場合があります
著者は理解しているが他の人は見ていない文脈や仮定。
これは、ローカルの説明や使い捨てのスクリプトとして機能します。チームエンジニアリング作業には弱いです。
問題は、プロンプトが非公式であるということではありません。非公式なものは役に立つことがよくあります。問題はプライベートなことだ
通常、プロンプトはエージェントの開始後にワークフローから消えます。自然にレビューにはならない
基準。プル リクエストと比較するのは困難です。次の人がその理由を理解するのに役立ちません
変化は存在します。
仕様は別の問題を解決します。これらにより、チームが継続的に検査できるように、割り当てに目に見える形が与えられます。
その仕様はさまざまな場所に生息できます。それは、リポジトリローカルの仕様、受け入れ基準の問題、または
BDD シナリオ、小さな設計ノート、変更提案、または動作に名前を付けるプル リクエストの説明
変化している。 OpenSpec はこのパターンの便利な実装の 1 つですが、それだけではありません。 GitHub
仕様キット、ガーキン スタイルのシナリオ、チーム RFC テンプレート、および通常の問題テンプレートはすべて同じものを保持できます。
コンテキストとレビュー基準を明確にする際には、規律を保つ必要があります。
それが、

シフトチームは気を配る必要があります。優れた仕様は、単にエージェントに指示するだけではありません。それは人間に与えます
そしてエージェントは、実装前、実装中、実装後に検査するために何かを共有しました。
思考を開始するのに便利ですが、共有レビューの対象としては弱いです。
実装前に表示され、レビュー中に役立ち、セッション終了後も永続的です。
割り当てレイヤーは意図と実行を分離します
最も強力なスペックは、小さな動作契約のように動作します。要件はシステムが何をすべきかを示します。
シナリオでは、多くの場合、Given/When/Then 形式で具体的な例が示されます。デザインノートとタスクリストは、
技術的なアプローチと実装チェックリストについて説明しますが、それらは同じものではありません。
要件。
この分離は、AI 支援エンジニアリングにとって最も役立つ分野の 1 つです。
インテントと実装が早期に混合されると、エージェントが間違った最適化を行う可能性があります。
事。提案された実装の詳細に忠実に従っているものの、チームの動作が欠落している可能性があります。
実際に必要です。あるいは、成功基準が満たされていないため、レビューが難しい、もっともらしい設計が生成される可能性があります。
決して明示されることはなかった。
課題レイヤーは 3 つの質問を分けて保持します。
正確さを定義する制約や例は何ですか?
現時点ではどの実装パスが適切だと思われますか?
これらの質問はつながっていますが、1 つの指示の塊にまとめてはいけません。の
実装は、エージェントがコードベースを読み取るにつれて進化する可能性があります。要件は十分に安定している必要があります
査読者は「作品はこれを満たしましたか?」と尋ねます。
これが、デルタ指向形式が既存のコードベースにとって興味深い理由でもあります。エンジニアリングの仕事のほとんどは、
グリーンフィールドではありません。チームは既存の行動を変えています。優れた仕様には次のように書かれています: 現在の仕様は次のとおりです
その契約に対する変更案がここにあります。査読者は必要ありません

全体的に精神的に違う
製品ドキュメント。彼らは行動デルタを調べることができます。
次のようなプライベート プロンプトを考えてみましょう。
不安定なログイン テストを修正し、変更が必要なものをすべて更新します。
開発者が一人で作業する場合はこれで十分かもしれません。チーム課題としては弱い。それは何とは言いません
障害が観察されたか、どの動作を安定させるべきか、どのチェックが重要か、またはどのような修正が行われたか
範囲の。
仕様が優れていると、作業の範囲が狭くなります。
確認された問題: コールバック要求がログイン前に到着すると、ログイン テストが断続的に失敗します。
セッション レコードはテスト アサーションに表示されます。
予期される動作: ログインではセッションを 1 つだけ作成し、ユーザーを元のセッションにリダイレクトする必要があります。
目的地。
制約: 認証チェックを弱めない、テストにスリープを追加しない、修正をローカルに保つ
コードベースでより広範な問題が示されていない限り、コールバック/セッション パス。
検証 : 影響を受ける認証テスト ファイルと、関連するバックエンドまたはフロントエンドのチェックを実行します。
オブジェクトの確認: 結果として得られる PR を、指定された動作、制約、および検証と比較します。
だからといって作品から判断力がなくなるわけではありません。エージェントに境界を与え、レビュー担当者に境界を与えます。
検査する関係。仕様がチームに表示されると、レビュー担当者はプル リクエストを比較できます。
エージェントが受け取ったのと同じコンテキストに対して。
仕様が小さくて修正可能なものであれば、ウォーターフォールではありません
仕様主導の作業は、多くの場合、合理的な反対意見を引き起こします。「これは新しい名前が付いた単なるウォーターフォールではないのか?」
チームが仕様書を儀式に変えるなら、それは可能だ。実装の数か月前に作成された巨大な文書は、
AI エージェントがそれを読み取るため、突然良くなります。
有用なカウンターパターンはより軽量で、流動的で反復的で修正が容易で、ブラウンフィールドファーストです。違う
フレームワークはそれを異なる方法で表現します。プロポーザルとデルタ仕様を使用するものもあります。ある程度の用途は、

訴訟チェックリストと
受け入れ基準。 BDD シナリオを使用するものもあります。重要なのは、これらは変化を伴う行動であるということです。
学習を遅らせるロックされたフェーズではありません。
その区別は重要です。割り当てレイヤーは、学習をフリーズさせるのではなく、曖昧さを軽減する必要があります。
実装がチームに何かを教えると、優れた仕様が変わる可能性があります。探索の結果、
最初のアプローチが間違っているので、設計を変更する必要があります。要件が広すぎる場合は、範囲を狭くする必要があります。
狭い。シナリオが誰も考慮していないエッジケースを明らかにした場合、仕様はそのシナリオを取得する必要があります。
この規律は「コードを書く前に完璧な計画を書く」というものではありません。規律は「目に見える意図を保つこと」です。
そして実装された現実は共に動きます。」
目に見える割り当てがない場合、レビュー担当者は主に差分をレビューします。
課題レイヤーを使用すると、レビュー担当者は次の 4 つの関係をレビューできます。
実装計画またはタスクの内訳
エージェントによって生成されたコードとテスト
CI、ローカルチェック、または手動レビューからの検証結果
この関係により、AI 支援による作業がより管理しやすくなります。査読者は次のことを求められていません
エージェントの信頼を信じてください。彼らは仕様、実装、証拠を比較しています。
フレームワークが異なれば、これをさまざまな方法で明示します。 OpenSpec には検証に関する概念があります
完全性、正確さ、一貫性。 GitHub の Spec Kit は、より厳格な仕様優先の立場をとっています。
BDD ワークフローは、実行可能または半実行可能な動作の期待値として例を使用します。課題主導型のチームは、
代わりに、受け入れ基準、ラベル、レビュー担当者、および CI 要件を使用します。
これらは同じ哲学ではありません。共通の教訓は範囲が狭く、より有益です。つまり、より強力です。
コーディングエージェントが増えれば増えるほど、彼らがサポートした割り当てを保存することがより重要になります

満足することができました。
チームの場合、これによりレビューの質問が次のように変更されます。
「この差分は、証拠とともに、私たちが指定した制約の下で、私たちが同意した動作の変更を満たしていますか?
検査してもらえますか？」
適切な仕様がチームのコンテキストとなる
適切な仕様はエージェントの起動に役立ちます。仕様を共有することで、チームの連携を保つことができます。
ランナーは、漠然とした個人的な意図を受け取り、孤立した実行環境に消えてはいけません。
レビュー担当者が最初からデコードする必要がある diff を返します。作業はチームが可能なコンテキストから開始する必要があります
チームがすでに理解している証拠 (ブランチ、コミット、プル リクエスト、CI 結果、
ランナーの概要、モデルの監査、人間によるレビュー。
どのツールも、チームごとに 1 つの仕様フレームワークを規定する必要はありません。一部のチームは課題を使用します。使う人もいるだろう
リポジトリローカルの仕様。軽量の設計ドキュメントや動作シナリオを使用する場合もあります。重要な境界線は、
コーディング エージェントの作業は、目に見える仕様とレビュー可能な結果に関連付けられたままである必要があります。
それが、有用な AI ランナーのワークフローをブラックボックスの自律的な出力から分離するものです。フォークラインはこれに続きます
同じ原則ですが、この原則は 1 つの製品よりも広範囲です。エージェントが共有コードに基づいて動作する場合、仕様は
そしてその結果は両方ともチームによって検査可能でなければなりません。
AI コーディング セッションは一時的なものです。リポジトリはそうではありません。
リポジトリローカル仕様のあまり知られていない利点の 1 つは、コードと共存できることです。チェックイン可能です
リポジトリ。機能または変更ごとに編成され、ワークランドとして更新されます。それにより、それらは次のような用途に役立ちます。
人々もエージェントも後で。
エージェントのコンテキストは脆弱であるため、これは重要です。チャット履歴が消去されます。コンテキスト ウィンドウがいっぱいになります。あ
別のモデルまたはツールが次のタスクを処理する可能性があります。元のプロンプトを書いた人はそうではない可能性があります
利用可能です。意図の唯一の記録がプライベートチャットだった場合、チームは負けます

チャット直後のコンテキスト
視界から消えてしまいます。
仕様はそのコンテキストに耐久性のある家を与えます。コード、テスト、ドキュメントを置き換えるものではありません。彼らはつながる
彼ら。新しいエージェントは現在の動作を読み取ることができます。新しい開発者はシステムが何であるかを理解できる
することが期待されています。レビュー担当者は、アーカイブされた変更を振り返って、何が変更されたかだけでなく、なぜ変更されたのかを確認できます。
変更が提案されました。
これは主に監査に関する議論ではありません。それは調整論です。チームには生き残る記憶が必要です
個人セッション。
限定された作業が適切な単位です
仕様によって、あらゆる種類の作業を安全に委任できるわけではありません。
これらは、動作の変更、バグの修正、互換性の更新、
小さな機能、移行ステップ、CI 修復、または明確な制約のある狭いリファクタリングなどです。そういった場合には、
チームは希望する変更を記述し、結果がそれに一致するかどうかを検査できます。
製品の方向性を選択したり、アーキテクチャを再設計したりするなど、判断が主なタスクである場合には、それらは弱くなります。
第一原則から、コードベース全体を改善し、UI を改善し、ユーザーが何を望むべきかを決定します。
この境界は健全です。スペックの要点は人間の判断を失わせることではない。ポイントは動くこと
適切な仕事を形に

[切り捨てられた]

## Original Extract

How to give AI coding agents the right context, boundaries, examples, and validation criteria before they start changing a repository.

Preparing Specs for AI Coding Agents - Forkline Blog
Product
Features Start runner work without introducing another layer of process Connectors GitHub, GitLab, Gitea, Forgejo and more Developers One operating surface for AI runner execution Docs Blog Pricing FAQ
Try Forkline
Back to blog
Forkline Engineering Preparing Specs for AI Coding Agents
How to give AI coding agents the right context, boundaries, examples, and validation criteria before they start changing a repository.
AI coding agents now edit repositories, run commands, and produce branches. That makes the spec before the work more important: it carries the context, boundaries, and success criteria the agent needs.
What a good coding-agent spec includes
Specs are becoming more important because AI coding agents are no longer only answering questions. They
are reading repositories, editing files, running commands, producing branches, and asking humans to review
the result. That changes what a prompt needs to become.
When an assistant only answers a question, a private prompt can be enough. When an agent changes a shared
codebase, the prompt becomes an assignment. And an assignment needs more than good wording. It needs the
right context, boundaries, examples, and a way to judge whether the work matched the original intent.
That is the practical reason to prepare a spec before sending a coding agent into a repository. The spec
does not need to be long. It does need to tell the agent what problem it is solving, what behavior should
change, what must not change, and how the result will be reviewed.
At minimum, a good coding-agent spec should give the agent five things:
the behavior that should change
the constraints the agent should preserve
examples or scenarios that define correctness
the validation evidence a reviewer should inspect
This is the useful idea behind spec-driven development, behavior scenarios, issue templates, lightweight
design docs, OpenSpec, GitHub Spec Kit, and many internal engineering proposal formats. The specific
framework matters less than the shape of the spec: the agent should receive enough context to act, and the
team should receive enough structure to review the result.
The spec is not a nicer prompt. It is the prepared assignment between human intent and machine execution.
Prompts are good at starting work. Specs are better at carrying it.
A private prompt is optimized for immediacy. It lives in a chat session. It can include shorthand, missing
context, and assumptions the author understands but nobody else sees.
That can work for a local explanation or a throwaway script. It is weaker for team engineering work.
The problem is not that prompts are informal. Informality is often useful. The problem is that private
prompts usually disappear from the workflow after the agent starts. They do not naturally become review
criteria. They are hard to compare against a pull request. They do not help the next person understand why
the change exists.
Specs solve a different problem. They give the assignment a visible shape the team can keep inspecting.
That spec can live in different places. It can be a repo-local spec, an issue with acceptance criteria, a
BDD scenario, a small design note, a change proposal, or a pull request description that names the behavior
being changed. OpenSpec is one useful implementation of this pattern, but it is not the only one. GitHub
Spec Kit, Gherkin-style scenarios, team RFC templates, and ordinary issue templates can all carry the same
discipline when they make context and review criteria explicit.
That is the shift teams should care about. A good spec does not merely instruct the agent. It gives humans
and agents something shared to inspect before, during, and after implementation.
Useful for starting thought, weak as a shared review object.
Visible before implementation, useful during review, durable after the session ends.
The assignment layer separates intent from execution
The strongest specs behave like small behavior contracts. Requirements say what the system should do.
Scenarios give concrete examples, often in a Given/When/Then style. Design notes and task lists can
describe the technical approach and implementation checklist, but those are not the same thing as the
requirement.
This separation is one of the most useful disciplines for AI-assisted engineering.
If the intent and implementation are mixed together too early, the agent can optimize for the wrong
thing. It may faithfully follow a suggested implementation detail while missing the behavior the team
actually needed. Or it may produce a plausible design that is hard to review because the success criteria
were never made explicit.
An assignment layer keeps three questions apart:
What constraints or examples define correctness?
What implementation path seems appropriate right now?
Those questions are connected, but they should not collapse into one blob of instructions. The
implementation can evolve as the agent reads the codebase. The requirement should remain stable enough
for a reviewer to ask: did the work satisfy this?
That is also why delta-oriented formats are interesting for existing codebases. Most engineering work is
not greenfield. Teams are changing behavior that already exists. A good spec says: here is the current
contract, and here is the proposed change to that contract. Reviewers do not need to mentally diff a whole
product document. They can look at the behavior delta.
Consider a private prompt like this:
Fix the flaky login test and update whatever needs changing.
That might be enough for a developer working alone. It is weak as a team assignment. It does not say what
failure is observed, which behavior should remain stable, which checks matter, or what kind of fix is out
of scope.
A better spec would make the work narrower:
Observed problem : the login test intermittently fails when the callback request arrives before the
session record is visible to the test assertion.
Expected behavior : login should create exactly one session and redirect the user to the original
destination.
Constraints : do not weaken auth checks, do not add sleeps to the test, and keep the fix local to
the callback/session path unless the codebase shows a wider issue.
Validation : run the affected auth test file and the relevant backend or frontend checks.
Review object : compare the resulting PR against the stated behavior, constraints, and validation.
That does not remove judgment from the work. It gives the agent a boundary and gives the reviewer a
relationship to inspect. When the spec is visible to the team, the reviewer can compare the pull request
against the same context the agent received.
Specs are not waterfall if they stay small and revisable
Spec-driven work often triggers a reasonable objection: is this just waterfall with a new name?
It can be, if the team turns specs into ceremony. A giant document, months before implementation, is not
suddenly better because an AI agent reads it.
The useful counter-pattern is lighter: fluid, iterative, easy to revise, and brownfield-first. Different
frameworks express that differently. Some use proposals and delta specs. Some use issue checklists and
acceptance criteria. Some use BDD scenarios. The important part is that these are actions around a change,
not locked phases that delay learning.
That distinction is important. The assignment layer should reduce ambiguity, not freeze learning.
A good spec can change when implementation teaches the team something. If exploration reveals that the
initial approach is wrong, the design should change. If a requirement was too broad, the scope should
narrow. If a scenario exposed an edge case nobody considered, the spec should gain that scenario.
The discipline is not “write the perfect plan before code.” The discipline is “keep the visible intent
and the implemented reality moving together.”
Without a visible assignment, reviewers mostly review the diff.
With an assignment layer, reviewers can review the relationship between four things:
the implementation plan or task breakdown
the code and tests produced by the agent
the validation result from CI, local checks, or manual review
That relationship is where AI-assisted work becomes more manageable. The reviewer is not being asked to
trust the agent’s confidence. They are comparing the spec, the implementation, and the evidence.
Different frameworks make this explicit in different ways. OpenSpec has verification concepts around
completeness, correctness, and coherence. GitHub’s Spec Kit takes a stricter specification-first position.
BDD workflows use examples as executable or semi-executable behavior expectations. Issue-driven teams may
use acceptance criteria, labels, reviewers, and CI requirements instead.
Those are not identical philosophies. The common lesson is narrower and more useful: the more powerful
coding agents become, the more important it is to preserve the assignment they were supposed to satisfy.
For teams, this changes the review question from:
“Does this diff satisfy the behavior change we agreed to, under the constraints we named, with evidence
we can inspect?”
The right spec becomes team context
A good spec helps the agent start. A shared spec helps the team stay aligned.
A runner should not receive vague private intent, disappear into an isolated execution environment, and
return a diff that reviewers have to decode from scratch. The work should start from context the team can
see, then return through evidence the team already understands: branch, commits, pull request, CI result,
runner summary, model audit, and human review.
No tool needs to prescribe one spec framework for every team. Some teams will use issues. Some will use
repo-local specs. Some will use lightweight design docs or behavior scenarios. The important boundary is
that coding-agent work should remain tied to a visible spec and a reviewable result.
That is what separates useful AI runner workflow from black-box autonomous output. Forkline follows this
same principle, but the principle is broader than any one product: if agents act on shared code, the spec
and the result should both be inspectable by the team.
AI coding sessions are temporary. Repositories are not.
One understated benefit of repo-local specs is that they can live with the code. They can be checked into
the repository, organized by capability or change, and updated as work lands. That makes them useful to
both people and agents later.
This matters because agent context is fragile. Chat history gets cleared. Context windows fill up. A
different model or tool may handle the next task. The person who wrote the original prompt may not be
available. If the only record of intent was a private chat, the team loses context as soon as the chat
falls out of view.
Specs give that context a durable home. They do not replace code, tests, or documentation. They connect
them. A new agent can read the current behavior. A new developer can understand what the system is
expected to do. A reviewer can look back at an archived change and see not only what changed, but why the
change was proposed.
This is not primarily an audit argument. It is a coordination argument. Teams need memory that survives
individual sessions.
Bounded work is the right unit
Specs do not make every kind of work safe to delegate.
They are strongest when the work is bounded: a behavior change, a bug fix, a compatibility update, a
small feature, a migration step, a CI repair, or a narrow refactor with clear constraints. In those cases,
the team can describe the desired change and inspect whether the result matches it.
They are weaker when the task is mostly judgment: choose the product direction, redesign the architecture
from first principles, improve the whole codebase, make the UI better, or decide what users should want.
This boundary is healthy. The point of specs is not to make human judgment disappear. The point is to move
appropriate work into a form

[truncated]
