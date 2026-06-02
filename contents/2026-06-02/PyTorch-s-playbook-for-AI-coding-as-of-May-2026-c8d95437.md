---
source: "https://docs.pytorch.org/devlogs/ai-agents/2026-05-30-ai-coding-playbook/"
hn_url: "https://news.ycombinator.com/item?id=48361727"
title: "PyTorch's playbook for AI coding, as of May 2026"
article_title: "PyTorch's playbook for AI coding, as of May 2026 — PyTorch DevLog"
author: "matt_d"
captured_at: "2026-06-02T00:32:53Z"
capture_tool: "hn-digest"
hn_id: 48361727
score: 3
comments: 0
posted_at: "2026-06-01T19:50:05Z"
tags:
  - hacker-news
  - translated
---

# PyTorch's playbook for AI coding, as of May 2026

- HN: [48361727](https://news.ycombinator.com/item?id=48361727)
- Source: [docs.pytorch.org](https://docs.pytorch.org/devlogs/ai-agents/2026-05-30-ai-coding-playbook/)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T19:50:05Z

## Translation

タイトル: AI コーディング用の PyTorch プレイブック (2026 年 5 月時点)
記事のタイトル: AI コーディングのための PyTorch のプレイブック、2026 年 5 月時点 — PyTorch DevLog

記事本文:
PyTorch DevLog GitHub ホーム / AI エージェント AI コーディングのための PyTorch のプレイブック (2026 年 5 月時点)
PyTorch チーム間で議論されている重要なトピックの 1 つは、
PyTorch コードベースは AI コーディング エージェントと連携する必要があります。今日はたくさんのPRの方に
PyTorch は AI によって作成されており、私たちがこれまで取り組んできたように、明らかな成長痛が生じています。
物事を理解した。最新の PyTorch コンパイラ オフサイトでの議論に基づく
(2026 年 5 月)、私は PyTorch で AI コーディングを行うためのこのプレイブックを組み立てました。半分です
説明的、半分規範的: 実践を成文化しようとしています。
チームの一部のメンバーの間で使用され、他の全員を連れて行きます。
この投稿が、次のことについての現在進行中の会話の始まりにすぎないことを願っています。
AI コーディング エージェントとどのように連携するか。
AI が生成したコードはスペクトルの中に存在すると考えることができます。
人間のコードとほぼ同じコードがありますが、
AI によって入力され、一方では完全にバイブコーディングされたソフトウェアです。
人間によって読まれたことはありません。
PyTorch は、多くの人に使用され、信頼されている運用ソフトウェアです。私たちは持っています
私たちが出荷するコードが正しく、理解できるものであることを保証するというユーザーに対する義務
そしてメンテナンス可能。 SOTA コーディング エージェントはより良い構築に役立つと考えています
今日私たちが純粋に手作業で構築できるソフトウェアよりも優れたソフトウェアが私たちに提供されます。
古いルールを適応させる必要がある新しい状況に直面しています。私たちは違うことを考えています
コードがスペクトル上のどこに存在するかに応じて、規範が必要となります。
人間が書いたコードの代替として
最も保守的な側では、AI コーディングを追加していますが、現状を維持しようとしています。
プロセスの他の多くの側面が修正されました。人間はすべての行を読む必要があります
コード。コードのすべての行に対して責任を負います。
ただし、すべてが同じままであるわけではありません。私たちは次のような新しい規範を提案します。
チアの時代に

p コード、人間によるレビューがボトルネックになっています。著者は次のようにすべきです
コードレビューを容易にするために懸命に努力してください。自分がどのような情報を持っているかを考えてみましょう
意図したレビュー担当者が必要とするものを書き留めます (AI が作成したコミット メッセージは
完全性には優れていますが、LLM はレビュー担当者の内容を知る可能性が低いです
知っているし知らない）。 PR が大きい場合は、一貫した内容であることを確認してください
変更に取り組み、「読み取り順序」を書き留めるためです。混ぜないでください
意味上の変更を伴う無関係な変更または表面的な変更。 LLMに分離を依頼する
これらはアウトです。
AI エージェントにコードに直接応答するよう依頼するのは非常に誘惑的です。
コメントをレビューします。私たちはこの点で人間の理解を信じているからです
私たちは、人間が対話することが重要であると考えています。
レビューコメントで。人間が質問を書くのに時間を費やしたなら、彼らはそれに値する
それに対する人間の反応。
反対に、コードレビューで尋ねられる可能性のある多くの伝統的な質問
AI エージェントが簡単に回答できます。コードレビュー担当者はAIに相談する必要があります
まず、未解決の質問のみを人間にエスカレートします。使用できます
GitHub の @claude またはローカルで PR をチェックアウトし、コーディング エージェントを使用してください
質問についてはローカルで質問してください。
コードレビューのコメントを自律的に修正するためにコーディングエージェントを使用しても問題ありません
(特に nits)、ただし、すべてを読んで所有する責任は依然としてあなたにあります。
修正。これには特に、コメントが実際にあったかどうかの確認が含まれます。
修正されました！
他の人のコードを直接編集するよう依頼することを検討してください。尋ねるべきです
たとえば、コミットをプッシュする前に最初に作成者を作成しますが、AI エージェントの場合、これは非常に困難です。
与えられるであろう細かいフィードバックを送信するコンパクトな方法
とにかくAIエージェントに直行します。よりコミュニケーションを図る良い方法でもあります
文章で説明するともっと時間がかかるような劇的な変化 – AIエージェント
テキストをコードに展開して役立ちます

テキストの意図を確認する
は明らかです。元の作成者は依然としてこれらを読んで所有権を取得する必要があります
変化します。
大量 AI PR は、エージェントを使用して多数の PR を並行して生成する場合です。例：
エージェントを使用して、問題トラッカーで問題を焼き尽くす。多くのバグはそうではありません
人間が数日かけて修正できるほど個別に重要なものですが、
全体として、バグ修正は重要であり、AI コーディング エージェントは大きな役割を果たします。
簡単に実現できる成果を殺す機会 (AI エージェントが実際にそうであるのと同じように)
セキュリティの脆弱性を発見するのが得意です。)
ここでの一般的な質問は、これらのことについてハイレベルの合意を得る必要があるということです。
修正には、全体として人が費やした時間を正当化する ROI が得られます。
エージェント群のオペレーターは初期設定を行う責任がありますが、
フィードバックに基づいてレビューし、指導し、改善することで、大勢の AI PR が
レビュー担当者の負担が増える。ポイントは、このレビューの負担に同意したことです
価値があります!
適切にカプセル化されたレビューされていないコード
現在のところ、未レビューの AI 生成コード (別名スロップ) は受け入れられていません。
メインの pytorch/pytorch リポジトリ。ただし、SOTA モデルの機能は
今日では、他の方法では存在できなかったシステムの作成が可能になります (例: 山登りによる)。
未レビューの AI 生成コード。今のところ、これらはすべて木の外に住んでいます
リポジトリ。これにより、このパッケージの実験的な性質が明らかになります。それ
コードの間違いも軽減されます (より迅速に出荷できるため)
リリース）。
未レビューのコードであっても、いくつかの標準に従う必要があります。
レビューされていないコードは、所有されていないコードを意味するわけではありません。責任を負う人間は、
全体的な生成プロセスの実行は、依然としてこのコードの責任を負っています。
(たとえば、信頼できないソースからの未レビューのコードを受け入れることはできません)。
このコードの所有者には依然として責任があります

e の設計をレビューするため
コード全体。すべての行を読まなくても、事実を尋ねることによって、
システムに関する質問があっても、強力なメンタルモデルを形成することは可能です
コードの全体的な設計について。今日の SOTA モデルでは、次のことがわかります。
設計ガイダンスにより全体的な成果が向上します。後から機能を要求しないでください
特徴。
スロップは、適切にカプセル化されたコンポーネント境界内に存在する必要があります。
スロップをシステムの残りの部分から分離する人間が設計した API 境界。
AI のスロップは、BC の影響を伴う公開 UI では受け入れられません。それ
すべてのコードを破棄して最初から書き直すことができるはずです
傾斜のない方法。スロップの出力が検証可能であればさらに良いのですが、
したがって、スロップコードが
バグっぽい。現時点ではこのような前例はありませんが、
未レビューのコードとレビュー済みの検証ツールをマージすることを検討する意思がある
pytorch/pytorch 。
未レビューのコードは AI コードレビューに合格する必要があります: セキュリティに関するガイダンス、テストなど
統合とグローバル不変条件。これは推論時間の単純な形式です
生成されたコードの下限を上げるスケーリング。
PyTorch の私たちの多くは、これらを使用せずにバイブコーディングされた便利な個人用ツールを持っています。
規格。これを阻止するつもりはありません。ただし、現時点での私たちの意見は、
このレベルの品質は、本来の PyTorch 機能には適切ではありません (たとえ
実験リポジトリ内)。
私たちは、次のツールが AI コーディング エージェントの世界に役立つと考えています。
近い将来に実装する予定です。
リスク認識差分自動レビュー (RADAR)。 OSSを作ります
Meta の RADAR システムの再実装。
このシステムは、一連の基準に基づいて PR を自動レビューし、自動的に承認します。
差分の危険性の評価について。社内でこれを使用しています
メットで

レーダーでも良好と認められています。目的はリラックスすることではない
PyTorch の PR ランディング標準。代わりに、周囲の摩擦を取り除くだけです
すでにレビューが簡単になっている PR。 RADAR は信頼を前提としています。
著者であるため、RADAR が承認するのは、そうでない場合は PyTorch メンテナに限定されます。
変更されるファイルに対する承認権限をすでに持っています。
自動 AI リンティング。を評価する AI リンターを実装します。
明確に定義された基準のリストとの差分。
サブシステムごとに。私たちはこれを AI コードレビューとは呼びません。
人間によるコードレビューに代わるものではありません。代わりに、目的は補完することです
AIは飽きることなく、熱心に学習するという特性を活かして、人間をサポートします。
指示に従ってください。適切なプロンプトが与えられると、AI は確実にチェックできます。
新しいパブリック API が導入されるかどうかに関係なく、所有権の間違いについては、
デバイスとホスト間の同期が導入されたかどうかなど。人間のレビュー担当者
これらのチェックを忘れる可能性があるため、AI は、
人間の査読者は何かをチェックすることを忘れていません。
FBcode から OSS テスト ケースへの自動生成。にとって大きな迷惑
非メタ貢献者は、PR が壊れたために元に戻されるときです
メタ内部の何か。以前は、最小限のエラーが発生していました
トレースは共有可能であり、障害をリバース エンジニアリングできました。代わりにしたいのは
常にエージェントのドラフトを作成し、合格する最小限の再現者を検証します。
差分なしでは失敗し、差分ありでは失敗します。多少のゲートはあるだろう
機密情報が漏洩しないようにするためのメカニズムですが、これは
信頼性の高いテスト ケースのパイプラインにより、OSS テスト カバレッジが向上します。
コード所有者。私たちは常にモジュールメンテナ（人）という概念を持っていました。
コードの特定の領域を担当する人）、そして私たちはさらに多くのことをしたいと考えています。
オーナーとしてこれを明確に定義してください

p は非常に貴重なリソースです。
それを育て、最大限に活用したいと考えています。一般的に、所有者は、
すべてのコードがどのように機能するかの全体像を把握している人、したがって
すべての変更を考慮する機会を得る特権を得る必要があります
この機能。いくつかの正確なメカニズムを解明する必要がありますが、
一般的な考え方は、コードベースの領域に影響を与える機能には、
自動クールダウン。PR は、
コード所有者は何らかの形でそれを支持しています。 AWOL コード所有者が次のことを行わないようにするには
すべての変更をブロックし、一定の SLA の後、PR は通常の方法で着陸できます。
レビュー – クールダウンは、誰かが突進してくる状況を避けるためにあります。
コード所有者が確認する前に変更してください。規範の側では、
PR 作成者は、デフォルトでコード所有者をある程度尊重する必要があります。例:
まずコード所有者に、あなたの変更が良いアイデアであることを説得する必要があります。
CODEOWNERS ファイルは、おそらくこのメカニズムの最良の実装ではありません。
それについては考えていきます (特に、非ファイルベースのルールがあると便利です)
所有権に関して）。
草案、変更のリクエスト。私たちはいくつかの新しいワークフローを広く適用しています
すべてのコードレビュー担当者が受信トレイを簡単に管理できるようにする PR のルール
PRレビューリクエストはゼロ。具体的には、 (1) PR を掲載する必要がある場合
しかし、レビューされることは望ましくありません。マークするのはあなたの責任です。
(2) あなたが PR の査読者であり、著者が
行動を起こす必要がある場合、変更をリクエストするのはあなたの責任です。
逆に、あなたがリクエストの変更がある PR の作成者である場合、それは次のとおりです。
レビューが必要な場合は、そのリクエストをクリアするのはあなたの責任です。の
これらの変更の目的は、誰が責任を負うのかを常に明確にすることです。
PRを前進させます。これらのツールは新しいものではありません

しかし、それらは使用されました
チームによって一貫性がありません – 私たちは現在、これを正式化し、使用を義務付けています
UI。

## Original Extract

PyTorch DevLog GitHub Home / AI Agents PyTorch's playbook for AI coding, as of May 2026
One of the important topics being discussed among the PyTorch team is how the
PyTorch codebase should engage with AI coding agents. Today, many PRs to
PyTorch are AI-authored, and there have been obvious growing pains as we’ve
figured things out. Based on discussions at the most recent PyTorch compiler offsite
(May 2026), I’ve assembled this playbook for AI coding in PyTorch. It is half
descriptive, half prescriptive: it is trying to codify practices that are
being used among some members of the team, and bring everyone else along.
Hopefully, this post is just the beginning of our ongoing conversation about
how to engage with AI coding agents.
We can think of AI generated code as living in a spectrum, where on one hand
we have code that is almost exactly the same as human code, except that it was
typed by an AI, and on the other hand completely vibe-coded software which has
never been read by a human.
PyTorch is production software, used and relied upon by many people. We have
a duty to our users to ensure that the code we ship is correct, understandable
and maintainable. We think that SOTA coding agents can help us build better
software than we could have built purely by hand today, but they present us
with novel situations that require adapting our old rules. We think different
norms are required depending on where code lives on the spectrum.
As a substitution for human written code
On the most conservative end, we are adding AI coding but trying to keep as
many other aspects of the process fixed. The human should read every line of
code. You are responsible for every line of code.
Not everything stays the same though. We propose these new norms:
In an age of cheap code, we are human review bottlenecked. Authors should
work hard to make code review easy. Think about what information your
intended reviewer needs and write it down (an AI written commit message is
good for completeness, but the LLM is unlikely to know what your reviewer
knows and doesn’t know). If a PR is big, make sure that there is a coherent
order to engage with the change and write down the “read order.” Don’t mix
unrelated or cosmetic changes with semantic changes; ask an LLM to separate
these out.
It is extremely tempting to ask an AI agent to directly respond to code
review comments. Because we believe in human understanding on this end
of the spectrum, we think it is important for humans to engage in dialog
in review comments. If a human spent time to write a question, they deserve
a human response in return.
On the flip side, many traditional questions one might ask in code review
can be easily answered by an AI agent. Code reviewers should consult AI
first, only escalating unresolved questions to humans. You can use
@claude on GitHub, or check out a PR locally and use your coding agent
locally on it for questions.
It is OK to use a coding agent to autonomously fix code review comments
(especially nits), but you are still responsible for reading and owning all
the fixes. This especially includes checking that the comment was actually
fixed!
Consider asking to directly edit someone else’s code. You should ask the
author first before, e.g., pushing a commit, but with AI agents this is a very
compact way to transmit small nitty feedback that would have been fed
straight into an AI agent anyway. It is also a good way to communicate more
dramatic changes that would take more time to explain in text–an AI agent
can expand text into code and help you verify that the intent of your text
is clear. The original author still has to read and take ownership of these
changes.
Mass AI PRs are when we use agents to generate many PRs in parallel; e.g.,
using agents to burn down issues on an issue tracker. Many bugs are not
individually important enough for a human to dedicate a few days fixing, but
in aggregate, fixing bugs is important, and AI coding agents are a big
opportunity to kill low hanging fruit (in the same way AI agents are really
good at discovering security vulnerabilities.)
The general ask here is that we should have high-level agreement that these
fixes, in aggregate have an ROI that justifies the human time spent on it.
While the operator of the agent swarm is responsible for doing initial
reviews, guiding it and improving it based on feedback, a mass of AI PRs will
increase reviewer burden. The point is to have agreed that this review burden
is worth it!
Well-encapsulated unreviewed code
As of today, we do not accept unreviewed AI generated code (aka slop) to the
main pytorch/pytorch repo. However, we think the capability of SOTA models
today enables the creation of systems that otherwise could not have existed (e.g., via hill climbing.) We have several live experiments in
unreviewed AI generated code; for now, these all live in out-of-tree
repositories. This makes clear the experimental nature of the package; it
also makes mistakes in the code lower stakes (as we can more rapidly ship
releases).
Even unreviewed code still needs to follow some standards:
Unreviewed code does not mean unowned code. The human responsible for
running the overall generation process is still accountable for this code
(for example, we cannot accept unreviewed code from untrusted sources).
The owner of this code still is responsible for reviewing the design of the
overall code. Even without having read every line, by asking factual
questions about systems, it’s still possible to form a strong mental model
about the overall design of the code. With SOTA models today, we find
design guidance improves overall outcomes. Don’t just ask for feature after
feature.
Slop should live in well-encapsulated component boundaries, where there is a
human-designed API boundary that separates slop from the rest of the system.
AI slop is not acceptable for public facing UI with BC implications. It
should be possible to throw out all the code and rewrite it from scratch in
a non-slop way. It is even better if the output of the slop is verifiable,
so it doesn’t matter for overall system correctness if the slop code is
buggy. Although we don’t currently have a precedent for this, we are
willing to consider merging unreviewed code with a reviewed verifier to
pytorch/pytorch .
Unreviewed code should pass AI code review: e.g., guidance on security, test
integration and global invariants. This is a simple form of inference-time
scaling that raises the floor for generated code.
Many of us at PyTorch have vibe-coded useful personal tools without these
standards. We don’t mean to discourage this! However, our current opinion is
that this level of quality is not appropriate for PyTorch features proper (even
in experimental repositories).
We think the following tools will be helpful for a world of AI coding agents, and
we plan to implement them in the near future:
Risk-Aware Diff Auto Review (RADAR). We will make an OSS
reimplementation of Meta’s RADAR system.
This system auto-reviews and auto-accepts PRs based on a set of criteria
around evaluating the riskiness of a diff. We’ve been using this internally
at Meta and the RADAR approves are good. The intention is not to relax
PyTorch’s PR landing standard; instead, it is to just remove friction around
PRs that are already easy to review. RADAR is predicated on trust of the
author, so RADAR approves will be limited to PyTorch maintainers who otherwise
already have approval rights on the files being changed.
Automatic AI Linting. We will implement an AI linter that will assess a
diff against a well defined list of criteria, which can be defined on a
subsystem-by-subsystem basis. We don’t call this AI code review, as it does
not replace human code review. Instead, the intention is to complement
humans by leveraging the fact that AIs never get bored and will studiously
follow instructions. Given appropriate prompting, an AI can reliably check
for ownership mistakes, whether or not a new public API is introduced,
whether or not a device-to-host sync was introduced, etc. A human reviewer
could forget to check these things, and an AI can help avoid relying on a
human reviewer remembering to check something.
Automated fbcode to OSS test case generation. A big annoyance for
non-Meta contributors is when your PR gets reverted because it broke
something Meta internal. Previously, you get whatever minimal error
trace was shareable and reverse engineer the failure. We want to instead
always have an agent draft and verify a minimal reproducer that passes
without the diff and fails with the diff. There will be some gating
mechanism to ensure secret information doesn’t leak, but this should setup a
reliable pipeline of test cases to improve our OSS test coverage.
Codeowners. We have always had a notion of module maintainers (people
who are responsible for a certain area of code), and we would like to more
sharply define this, as ownership is a very valuable resource and we would
like to foster it and make the best use of it. In general, the owner is
someone who has the overall picture of how all the code works, and therefore
should have the privilege of getting a chance to weigh in on all changes to
this feature. Some of the precise mechanics should be worked out, but the
general idea is that features affecting areas of the codebase should have an
automatic cooldown, where the PR shouldn’t be mergeable without the
codeowner endorsing it in some way. To ensure an AWOL codeowner doesn’t
block all changes, after some SLA, a PR can be landed with normal
review–the cooldown is there to avoid situations where someone rushes in a
change before the codeowner manages to take a look. On the norms side,
PR authors should give some deference by default to the codeowner–e.g.,
you should convince the codeowner first that your change is a good idea!
The CODEOWNERS file is probably not the best implementation of this mechanism;
we’ll figure it out (in particular it would be helpful to have non-file-based rules
on ownership).
Draft, Request Changes. We are universally applying some new workflow
rules for PRs to ensure that all code reviewers can easily maintain inbox
zero on PR review requests. Specifically, (1) if you need to put up a PR
but do not wish for it to be reviewed, it is your responsibility to mark it
as draft, (2) if you are a reviewer on a PR and you feel that the author
needs to take action, it is your responsibility to request changes.
Conversely, if you are an author of a PR that has request changes, it is
your responsibility to clear that request when you want review. The
intention of these changes is to always make it clear who is responsible for
moving a PR forward. These tools are not new but they were used
inconsistently by the team–we are now formalizing and requiring use of this
UI.
